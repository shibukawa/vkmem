package p1

import (
	base "github.com/shibukawa/valkeymem/internal/aot/vkaot/base"
	"unsafe"
)

func F_bit_lshift(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 float64
	_ = v5
	var v8 int32
	_ = v8
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v23 float64
	_ = v23
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	v5 = F_lua_tonumber(m, l0, int32(1))
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		return int32(0)
	} else {
		v12 = base.I32_wrap_i64(base.I64_reinterpret_f64(base.F64_add(v5, float64(6.755399441055744e+15))))
		if v12 != 0 {
			v23 = F_lua_tonumber(m, l0, int32(2))
			mBase = m.M
			v24 = m.ExcPending
			if v24 != 0 {
				return int32(0)
			} else {
				v28 = base.I32_wrap_i64(base.I64_reinterpret_f64(base.F64_add(v23, float64(6.755399441055744e+15))))
				if v28 != 0 {
					v41 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
					*(*int32)(unsafe.Add(mBase, uint32(v41)+8)) = int32(3)
					*(*float64)(unsafe.Add(mBase, uint32(v41))) = base.F64_convert_i32_s(v12 << (uint(v28) % 32))
					v45 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
					*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v45 + int32(16)
					return int32(1)
				} else {
					v30 = F_lua_isnumber(m, l0, int32(2))
					mBase = m.M
					v31 = m.ExcPending
					if v31 != 0 {
						return int32(0)
					} else {
						if v30 != 0 {
							v41 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
							*(*int32)(unsafe.Add(mBase, uint32(v41)+8)) = int32(3)
							*(*float64)(unsafe.Add(mBase, uint32(v41))) = base.F64_convert_i32_s(v12 << (uint(v28) % 32))
							v45 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
							*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v45 + int32(16)
							return int32(1)
						} else {
							v33 = m.G3
							v36 = F_luaL_typerror(m, l0, int32(2), v33+int32(_a2760))
							mBase = m.M
							v37 = m.ExcPending
							if v37 != 0 {
								return int32(0)
							} else {
								v41 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
								*(*int32)(unsafe.Add(mBase, uint32(v41)+8)) = int32(3)
								*(*float64)(unsafe.Add(mBase, uint32(v41))) = base.F64_convert_i32_s(v12 << (uint(v28) % 32))
								v45 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
								*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v45 + int32(16)
								return int32(1)
							}
						}
					}
				}
			}
		} else {
			v14 = F_lua_isnumber(m, l0, int32(1))
			mBase = m.M
			v15 = m.ExcPending
			if v15 != 0 {
				return int32(0)
			} else {
				if v14 != 0 {
					v23 = F_lua_tonumber(m, l0, int32(2))
					mBase = m.M
					v24 = m.ExcPending
					if v24 != 0 {
						return int32(0)
					} else {
						v28 = base.I32_wrap_i64(base.I64_reinterpret_f64(base.F64_add(v23, float64(6.755399441055744e+15))))
						if v28 != 0 {
							v41 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
							*(*int32)(unsafe.Add(mBase, uint32(v41)+8)) = int32(3)
							*(*float64)(unsafe.Add(mBase, uint32(v41))) = base.F64_convert_i32_s(v12 << (uint(v28) % 32))
							v45 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
							*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v45 + int32(16)
							return int32(1)
						} else {
							v30 = F_lua_isnumber(m, l0, int32(2))
							mBase = m.M
							v31 = m.ExcPending
							if v31 != 0 {
								return int32(0)
							} else {
								if v30 != 0 {
									v41 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
									*(*int32)(unsafe.Add(mBase, uint32(v41)+8)) = int32(3)
									*(*float64)(unsafe.Add(mBase, uint32(v41))) = base.F64_convert_i32_s(v12 << (uint(v28) % 32))
									v45 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
									*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v45 + int32(16)
									return int32(1)
								} else {
									v33 = m.G3
									v36 = F_luaL_typerror(m, l0, int32(2), v33+int32(_a2760))
									mBase = m.M
									v37 = m.ExcPending
									if v37 != 0 {
										return int32(0)
									} else {
										v41 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
										*(*int32)(unsafe.Add(mBase, uint32(v41)+8)) = int32(3)
										*(*float64)(unsafe.Add(mBase, uint32(v41))) = base.F64_convert_i32_s(v12 << (uint(v28) % 32))
										v45 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
										*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v45 + int32(16)
										return int32(1)
									}
								}
							}
						}
					}
				} else {
					v17 = m.G3
					v20 = F_luaL_typerror(m, l0, int32(1), v17+int32(_a2760))
					mBase = m.M
					v21 = m.ExcPending
					if v21 != 0 {
						return int32(0)
					} else {
						v23 = F_lua_tonumber(m, l0, int32(2))
						mBase = m.M
						v24 = m.ExcPending
						if v24 != 0 {
							return int32(0)
						} else {
							v28 = base.I32_wrap_i64(base.I64_reinterpret_f64(base.F64_add(v23, float64(6.755399441055744e+15))))
							if v28 != 0 {
								v41 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
								*(*int32)(unsafe.Add(mBase, uint32(v41)+8)) = int32(3)
								*(*float64)(unsafe.Add(mBase, uint32(v41))) = base.F64_convert_i32_s(v12 << (uint(v28) % 32))
								v45 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
								*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v45 + int32(16)
								return int32(1)
							} else {
								v30 = F_lua_isnumber(m, l0, int32(2))
								mBase = m.M
								v31 = m.ExcPending
								if v31 != 0 {
									return int32(0)
								} else {
									if v30 != 0 {
										v41 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
										*(*int32)(unsafe.Add(mBase, uint32(v41)+8)) = int32(3)
										*(*float64)(unsafe.Add(mBase, uint32(v41))) = base.F64_convert_i32_s(v12 << (uint(v28) % 32))
										v45 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
										*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v45 + int32(16)
										return int32(1)
									} else {
										v33 = m.G3
										v36 = F_luaL_typerror(m, l0, int32(2), v33+int32(_a2760))
										mBase = m.M
										v37 = m.ExcPending
										if v37 != 0 {
											return int32(0)
										} else {
											v41 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
											*(*int32)(unsafe.Add(mBase, uint32(v41)+8)) = int32(3)
											*(*float64)(unsafe.Add(mBase, uint32(v41))) = base.F64_convert_i32_s(v12 << (uint(v28) % 32))
											v45 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
											*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v45 + int32(16)
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
func F_bit_ror(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 float64
	_ = v5
	var v8 int32
	_ = v8
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v23 float64
	_ = v23
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	v5 = F_lua_tonumber(m, l0, int32(1))
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		return int32(0)
	} else {
		v12 = base.I32_wrap_i64(base.I64_reinterpret_f64(base.F64_add(v5, float64(6.755399441055744e+15))))
		if v12 != 0 {
			v23 = F_lua_tonumber(m, l0, int32(2))
			mBase = m.M
			v24 = m.ExcPending
			if v24 != 0 {
				return int32(0)
			} else {
				v28 = base.I32_wrap_i64(base.I64_reinterpret_f64(base.F64_add(v23, float64(6.755399441055744e+15))))
				if v28 != 0 {
					v41 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
					*(*int32)(unsafe.Add(mBase, uint32(v41)+8)) = int32(3)
					*(*float64)(unsafe.Add(mBase, uint32(v41))) = base.F64_convert_i32_s(base.I32_rotr(v12, v28))
					v45 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
					*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v45 + int32(16)
					return int32(1)
				} else {
					v30 = F_lua_isnumber(m, l0, int32(2))
					mBase = m.M
					v31 = m.ExcPending
					if v31 != 0 {
						return int32(0)
					} else {
						if v30 != 0 {
							v41 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
							*(*int32)(unsafe.Add(mBase, uint32(v41)+8)) = int32(3)
							*(*float64)(unsafe.Add(mBase, uint32(v41))) = base.F64_convert_i32_s(base.I32_rotr(v12, v28))
							v45 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
							*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v45 + int32(16)
							return int32(1)
						} else {
							v33 = m.G3
							v36 = F_luaL_typerror(m, l0, int32(2), v33+int32(_a2760))
							mBase = m.M
							v37 = m.ExcPending
							if v37 != 0 {
								return int32(0)
							} else {
								v41 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
								*(*int32)(unsafe.Add(mBase, uint32(v41)+8)) = int32(3)
								*(*float64)(unsafe.Add(mBase, uint32(v41))) = base.F64_convert_i32_s(base.I32_rotr(v12, v28))
								v45 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
								*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v45 + int32(16)
								return int32(1)
							}
						}
					}
				}
			}
		} else {
			v14 = F_lua_isnumber(m, l0, int32(1))
			mBase = m.M
			v15 = m.ExcPending
			if v15 != 0 {
				return int32(0)
			} else {
				if v14 != 0 {
					v23 = F_lua_tonumber(m, l0, int32(2))
					mBase = m.M
					v24 = m.ExcPending
					if v24 != 0 {
						return int32(0)
					} else {
						v28 = base.I32_wrap_i64(base.I64_reinterpret_f64(base.F64_add(v23, float64(6.755399441055744e+15))))
						if v28 != 0 {
							v41 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
							*(*int32)(unsafe.Add(mBase, uint32(v41)+8)) = int32(3)
							*(*float64)(unsafe.Add(mBase, uint32(v41))) = base.F64_convert_i32_s(base.I32_rotr(v12, v28))
							v45 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
							*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v45 + int32(16)
							return int32(1)
						} else {
							v30 = F_lua_isnumber(m, l0, int32(2))
							mBase = m.M
							v31 = m.ExcPending
							if v31 != 0 {
								return int32(0)
							} else {
								if v30 != 0 {
									v41 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
									*(*int32)(unsafe.Add(mBase, uint32(v41)+8)) = int32(3)
									*(*float64)(unsafe.Add(mBase, uint32(v41))) = base.F64_convert_i32_s(base.I32_rotr(v12, v28))
									v45 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
									*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v45 + int32(16)
									return int32(1)
								} else {
									v33 = m.G3
									v36 = F_luaL_typerror(m, l0, int32(2), v33+int32(_a2760))
									mBase = m.M
									v37 = m.ExcPending
									if v37 != 0 {
										return int32(0)
									} else {
										v41 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
										*(*int32)(unsafe.Add(mBase, uint32(v41)+8)) = int32(3)
										*(*float64)(unsafe.Add(mBase, uint32(v41))) = base.F64_convert_i32_s(base.I32_rotr(v12, v28))
										v45 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
										*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v45 + int32(16)
										return int32(1)
									}
								}
							}
						}
					}
				} else {
					v17 = m.G3
					v20 = F_luaL_typerror(m, l0, int32(1), v17+int32(_a2760))
					mBase = m.M
					v21 = m.ExcPending
					if v21 != 0 {
						return int32(0)
					} else {
						v23 = F_lua_tonumber(m, l0, int32(2))
						mBase = m.M
						v24 = m.ExcPending
						if v24 != 0 {
							return int32(0)
						} else {
							v28 = base.I32_wrap_i64(base.I64_reinterpret_f64(base.F64_add(v23, float64(6.755399441055744e+15))))
							if v28 != 0 {
								v41 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
								*(*int32)(unsafe.Add(mBase, uint32(v41)+8)) = int32(3)
								*(*float64)(unsafe.Add(mBase, uint32(v41))) = base.F64_convert_i32_s(base.I32_rotr(v12, v28))
								v45 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
								*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v45 + int32(16)
								return int32(1)
							} else {
								v30 = F_lua_isnumber(m, l0, int32(2))
								mBase = m.M
								v31 = m.ExcPending
								if v31 != 0 {
									return int32(0)
								} else {
									if v30 != 0 {
										v41 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
										*(*int32)(unsafe.Add(mBase, uint32(v41)+8)) = int32(3)
										*(*float64)(unsafe.Add(mBase, uint32(v41))) = base.F64_convert_i32_s(base.I32_rotr(v12, v28))
										v45 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
										*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v45 + int32(16)
										return int32(1)
									} else {
										v33 = m.G3
										v36 = F_luaL_typerror(m, l0, int32(2), v33+int32(_a2760))
										mBase = m.M
										v37 = m.ExcPending
										if v37 != 0 {
											return int32(0)
										} else {
											v41 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
											*(*int32)(unsafe.Add(mBase, uint32(v41)+8)) = int32(3)
											*(*float64)(unsafe.Add(mBase, uint32(v41))) = base.F64_convert_i32_s(base.I32_rotr(v12, v28))
											v45 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
											*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v45 + int32(16)
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
func F_bit_tobit(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 float64
	_ = v4
	var v7 int32
	_ = v7
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	v4 = F_lua_tonumber(m, l0, int32(1))
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		v11 = base.I32_wrap_i64(base.I64_reinterpret_f64(base.F64_add(v4, float64(6.755399441055744e+15))))
		if v11 != 0 {
			v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			*(*int32)(unsafe.Add(mBase, uint32(v23)+8)) = int32(3)
			*(*float64)(unsafe.Add(mBase, uint32(v23))) = base.F64_convert_i32_s(v11)
			v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v27 + int32(16)
			return int32(1)
		} else {
			v13 = F_lua_isnumber(m, l0, int32(1))
			mBase = m.M
			v14 = m.ExcPending
			if v14 != 0 {
				return int32(0)
			} else {
				if v13 != 0 {
					v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
					*(*int32)(unsafe.Add(mBase, uint32(v23)+8)) = int32(3)
					*(*float64)(unsafe.Add(mBase, uint32(v23))) = base.F64_convert_i32_s(v11)
					v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
					*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v27 + int32(16)
					return int32(1)
				} else {
					v16 = m.G3
					v19 = F_luaL_typerror(m, l0, int32(1), v16+int32(_a2760))
					mBase = m.M
					v20 = m.ExcPending
					if v20 != 0 {
						return int32(0)
					} else {
						v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
						*(*int32)(unsafe.Add(mBase, uint32(v23)+8)) = int32(3)
						*(*float64)(unsafe.Add(mBase, uint32(v23))) = base.F64_convert_i32_s(v11)
						v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
						*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v27 + int32(16)
						return int32(1)
					}
				}
			}
		}
	}
}
func F_bit_tohex(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v14 float64
	_ = v14
	var v17 int32
	_ = v17
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v35 int32
	_ = v35
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v83 int32
	_ = v83
	var v86 int32
	_ = v86
	var v89 int32
	_ = v89
	var v93 float64
	_ = v93
	var v94 int32
	_ = v94
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v104 int32
	_ = v104
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
	var v114 int32
	_ = v114
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v128 int32
	_ = v128
	var v135 int32
	_ = v135
	var v139 int32
	_ = v139
	var v150 int32
	_ = v150
	var v161 int32
	_ = v161
	var v172 int32
	_ = v172
	var v183 int32
	_ = v183
	var v194 int32
	_ = v194
	var v205 int32
	_ = v205
	var v214 int32
	_ = v214
	var v221 int32
	_ = v221
	v9 = m.G0
	v11 = v9 - int32(16)
	m.G0 = v11
	v14 = F_lua_tonumber(m, l0, int32(1))
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		return int32(0)
	} else {
		v21 = base.I32_wrap_i64(base.I64_reinterpret_f64(base.F64_add(v14, float64(6.755399441055744e+15))))
		if v21 != 0 {
			v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
			v40 = v35 + int32(16)
			v41 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			if base.Ui32(v40) < base.Ui32(v41) {
				v83 = m.G398
				if v40 != v83 {
					v86 = *(*int32)(unsafe.Add(mBase, uint32(v40)+8))
					v89 = v86
				} else {
					v89 = int32(-1)
				}
			} else {
				v89 = int32(-1)
			}
			if v89 == int32(-1) {
				v109 = int32(8)
				v111 = int32(-2147483647)
				if v111 < v109 {
					v114 = v109
				} else {
					v114 = v111
				}
				v117 = base.B2i32(v109 < int32(0))
				if v109 < int32(0) {
					v118 = int32(0) - v114
				} else {
					v118 = v114
				}
				v119 = int32(8)
				if v118 < v119 {
					v122 = v118
				} else {
					v122 = v119
				}
				v123 = m.G3
				if v118 < int32(1) {
				} else {
					v128 = v122 + (v11 + int32(8))
					if v109 < int32(0) {
						v135 = v123 + int32(_a2761)
					} else {
						v135 = v123 + int32(_a2762)
					}
					v139 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v135+v21&int32(15)))))
					*(*uint8)(unsafe.Add(mBase, uint32(v128+int32(-1)))) = uint8(v139)
					if v118 == int32(1) {
					} else {
						v150 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v135+int32(base.Ui32(v21)>>(uint(int32(4))%32))&int32(15)))))
						*(*uint8)(unsafe.Add(mBase, uint32(v128+int32(-2)))) = uint8(v150)
						if v118 < int32(3) {
						} else {
							v161 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v135+int32(base.Ui32(v21)>>(uint(int32(8))%32))&int32(15)))))
							*(*uint8)(unsafe.Add(mBase, uint32(v128+int32(-3)))) = uint8(v161)
							if v118 == int32(3) {
							} else {
								v172 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v135+int32(base.Ui32(v21)>>(uint(int32(12))%32))&int32(15)))))
								*(*uint8)(unsafe.Add(mBase, uint32(v128+int32(-4)))) = uint8(v172)
								if v118 < int32(5) {
								} else {
									v183 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v135+int32(base.Ui32(v21)>>(uint(int32(16))%32))&int32(15)))))
									*(*uint8)(unsafe.Add(mBase, uint32(v128+int32(-5)))) = uint8(v183)
									if v118 == int32(5) {
									} else {
										v194 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v135+int32(base.Ui32(v21)>>(uint(int32(20))%32))&int32(15)))))
										*(*uint8)(unsafe.Add(mBase, uint32(v128+int32(-6)))) = uint8(v194)
										if v118 < int32(7) {
										} else {
											v205 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v135+int32(base.Ui32(v21)>>(uint(int32(24))%32))&int32(15)))))
											*(*uint8)(unsafe.Add(mBase, uint32(v128+int32(-7)))) = uint8(v205)
											if v118 == int32(7) {
											} else {
												v214 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v135+int32(base.Ui32(v21)>>(uint(int32(28))%32))))))
												*(*uint8)(unsafe.Add(mBase, uint32(v128+int32(-8)))) = uint8(v214)
											}
										}
									}
								}
							}
						}
					}
				}
				F_lua_pushlstring(m, l0, v11+int32(8), v122)
				mBase = m.M
				v221 = m.ExcPending
				if v221 != 0 {
					return int32(0)
				} else {
					m.G0 = v11 + int32(16)
					return int32(1)
				}
			} else {
				v93 = F_lua_tonumber(m, l0, int32(2))
				mBase = m.M
				v94 = m.ExcPending
				if v94 != 0 {
					return int32(0)
				} else {
					v98 = base.I32_wrap_i64(base.I64_reinterpret_f64(base.F64_add(v93, float64(6.755399441055744e+15))))
					if v98 != 0 {
						v109 = v98
						v111 = int32(-2147483647)
						if v111 < v109 {
							v114 = v109
						} else {
							v114 = v111
						}
						v117 = base.B2i32(v109 < int32(0))
						if v109 < int32(0) {
							v118 = int32(0) - v114
						} else {
							v118 = v114
						}
						v119 = int32(8)
						if v118 < v119 {
							v122 = v118
						} else {
							v122 = v119
						}
						v123 = m.G3
						if v118 < int32(1) {
						} else {
							v128 = v122 + (v11 + int32(8))
							if v109 < int32(0) {
								v135 = v123 + int32(_a2761)
							} else {
								v135 = v123 + int32(_a2762)
							}
							v139 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v135+v21&int32(15)))))
							*(*uint8)(unsafe.Add(mBase, uint32(v128+int32(-1)))) = uint8(v139)
							if v118 == int32(1) {
							} else {
								v150 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v135+int32(base.Ui32(v21)>>(uint(int32(4))%32))&int32(15)))))
								*(*uint8)(unsafe.Add(mBase, uint32(v128+int32(-2)))) = uint8(v150)
								if v118 < int32(3) {
								} else {
									v161 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v135+int32(base.Ui32(v21)>>(uint(int32(8))%32))&int32(15)))))
									*(*uint8)(unsafe.Add(mBase, uint32(v128+int32(-3)))) = uint8(v161)
									if v118 == int32(3) {
									} else {
										v172 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v135+int32(base.Ui32(v21)>>(uint(int32(12))%32))&int32(15)))))
										*(*uint8)(unsafe.Add(mBase, uint32(v128+int32(-4)))) = uint8(v172)
										if v118 < int32(5) {
										} else {
											v183 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v135+int32(base.Ui32(v21)>>(uint(int32(16))%32))&int32(15)))))
											*(*uint8)(unsafe.Add(mBase, uint32(v128+int32(-5)))) = uint8(v183)
											if v118 == int32(5) {
											} else {
												v194 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v135+int32(base.Ui32(v21)>>(uint(int32(20))%32))&int32(15)))))
												*(*uint8)(unsafe.Add(mBase, uint32(v128+int32(-6)))) = uint8(v194)
												if v118 < int32(7) {
												} else {
													v205 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v135+int32(base.Ui32(v21)>>(uint(int32(24))%32))&int32(15)))))
													*(*uint8)(unsafe.Add(mBase, uint32(v128+int32(-7)))) = uint8(v205)
													if v118 == int32(7) {
													} else {
														v214 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v135+int32(base.Ui32(v21)>>(uint(int32(28))%32))))))
														*(*uint8)(unsafe.Add(mBase, uint32(v128+int32(-8)))) = uint8(v214)
													}
												}
											}
										}
									}
								}
							}
						}
						F_lua_pushlstring(m, l0, v11+int32(8), v122)
						mBase = m.M
						v221 = m.ExcPending
						if v221 != 0 {
							return int32(0)
						} else {
							m.G0 = v11 + int32(16)
							return int32(1)
						}
					} else {
						v99 = int32(0)
						v101 = F_lua_isnumber(m, l0, int32(2))
						mBase = m.M
						v102 = m.ExcPending
						if v102 != 0 {
							return int32(0)
						} else {
							if v101 != 0 {
								v109 = v99
								v111 = int32(-2147483647)
								if v111 < v109 {
									v114 = v109
								} else {
									v114 = v111
								}
								v117 = base.B2i32(v109 < int32(0))
								if v109 < int32(0) {
									v118 = int32(0) - v114
								} else {
									v118 = v114
								}
								v119 = int32(8)
								if v118 < v119 {
									v122 = v118
								} else {
									v122 = v119
								}
								v123 = m.G3
								if v118 < int32(1) {
								} else {
									v128 = v122 + (v11 + int32(8))
									if v109 < int32(0) {
										v135 = v123 + int32(_a2761)
									} else {
										v135 = v123 + int32(_a2762)
									}
									v139 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v135+v21&int32(15)))))
									*(*uint8)(unsafe.Add(mBase, uint32(v128+int32(-1)))) = uint8(v139)
									if v118 == int32(1) {
									} else {
										v150 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v135+int32(base.Ui32(v21)>>(uint(int32(4))%32))&int32(15)))))
										*(*uint8)(unsafe.Add(mBase, uint32(v128+int32(-2)))) = uint8(v150)
										if v118 < int32(3) {
										} else {
											v161 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v135+int32(base.Ui32(v21)>>(uint(int32(8))%32))&int32(15)))))
											*(*uint8)(unsafe.Add(mBase, uint32(v128+int32(-3)))) = uint8(v161)
											if v118 == int32(3) {
											} else {
												v172 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v135+int32(base.Ui32(v21)>>(uint(int32(12))%32))&int32(15)))))
												*(*uint8)(unsafe.Add(mBase, uint32(v128+int32(-4)))) = uint8(v172)
												if v118 < int32(5) {
												} else {
													v183 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v135+int32(base.Ui32(v21)>>(uint(int32(16))%32))&int32(15)))))
													*(*uint8)(unsafe.Add(mBase, uint32(v128+int32(-5)))) = uint8(v183)
													if v118 == int32(5) {
													} else {
														v194 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v135+int32(base.Ui32(v21)>>(uint(int32(20))%32))&int32(15)))))
														*(*uint8)(unsafe.Add(mBase, uint32(v128+int32(-6)))) = uint8(v194)
														if v118 < int32(7) {
														} else {
															v205 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v135+int32(base.Ui32(v21)>>(uint(int32(24))%32))&int32(15)))))
															*(*uint8)(unsafe.Add(mBase, uint32(v128+int32(-7)))) = uint8(v205)
															if v118 == int32(7) {
															} else {
																v214 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v135+int32(base.Ui32(v21)>>(uint(int32(28))%32))))))
																*(*uint8)(unsafe.Add(mBase, uint32(v128+int32(-8)))) = uint8(v214)
															}
														}
													}
												}
											}
										}
									}
								}
								F_lua_pushlstring(m, l0, v11+int32(8), v122)
								mBase = m.M
								v221 = m.ExcPending
								if v221 != 0 {
									return int32(0)
								} else {
									m.G0 = v11 + int32(16)
									return int32(1)
								}
							} else {
								v104 = m.G3
								v107 = F_luaL_typerror(m, l0, int32(2), v104+int32(_a2760))
								mBase = m.M
								v108 = m.ExcPending
								if v108 != 0 {
									return int32(0)
								} else {
									v109 = v99
									v111 = int32(-2147483647)
									if v111 < v109 {
										v114 = v109
									} else {
										v114 = v111
									}
									v117 = base.B2i32(v109 < int32(0))
									if v109 < int32(0) {
										v118 = int32(0) - v114
									} else {
										v118 = v114
									}
									v119 = int32(8)
									if v118 < v119 {
										v122 = v118
									} else {
										v122 = v119
									}
									v123 = m.G3
									if v118 < int32(1) {
									} else {
										v128 = v122 + (v11 + int32(8))
										if v109 < int32(0) {
											v135 = v123 + int32(_a2761)
										} else {
											v135 = v123 + int32(_a2762)
										}
										v139 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v135+v21&int32(15)))))
										*(*uint8)(unsafe.Add(mBase, uint32(v128+int32(-1)))) = uint8(v139)
										if v118 == int32(1) {
										} else {
											v150 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v135+int32(base.Ui32(v21)>>(uint(int32(4))%32))&int32(15)))))
											*(*uint8)(unsafe.Add(mBase, uint32(v128+int32(-2)))) = uint8(v150)
											if v118 < int32(3) {
											} else {
												v161 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v135+int32(base.Ui32(v21)>>(uint(int32(8))%32))&int32(15)))))
												*(*uint8)(unsafe.Add(mBase, uint32(v128+int32(-3)))) = uint8(v161)
												if v118 == int32(3) {
												} else {
													v172 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v135+int32(base.Ui32(v21)>>(uint(int32(12))%32))&int32(15)))))
													*(*uint8)(unsafe.Add(mBase, uint32(v128+int32(-4)))) = uint8(v172)
													if v118 < int32(5) {
													} else {
														v183 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v135+int32(base.Ui32(v21)>>(uint(int32(16))%32))&int32(15)))))
														*(*uint8)(unsafe.Add(mBase, uint32(v128+int32(-5)))) = uint8(v183)
														if v118 == int32(5) {
														} else {
															v194 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v135+int32(base.Ui32(v21)>>(uint(int32(20))%32))&int32(15)))))
															*(*uint8)(unsafe.Add(mBase, uint32(v128+int32(-6)))) = uint8(v194)
															if v118 < int32(7) {
															} else {
																v205 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v135+int32(base.Ui32(v21)>>(uint(int32(24))%32))&int32(15)))))
																*(*uint8)(unsafe.Add(mBase, uint32(v128+int32(-7)))) = uint8(v205)
																if v118 == int32(7) {
																} else {
																	v214 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v135+int32(base.Ui32(v21)>>(uint(int32(28))%32))))))
																	*(*uint8)(unsafe.Add(mBase, uint32(v128+int32(-8)))) = uint8(v214)
																}
															}
														}
													}
												}
											}
										}
									}
									F_lua_pushlstring(m, l0, v11+int32(8), v122)
									mBase = m.M
									v221 = m.ExcPending
									if v221 != 0 {
										return int32(0)
									} else {
										m.G0 = v11 + int32(16)
										return int32(1)
									}
								}
							}
						}
					}
				}
			}
		} else {
			v23 = F_lua_isnumber(m, l0, int32(1))
			mBase = m.M
			v24 = m.ExcPending
			if v24 != 0 {
				return int32(0)
			} else {
				if v23 != 0 {
					v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
					v40 = v35 + int32(16)
					v41 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
					if base.Ui32(v40) < base.Ui32(v41) {
						v83 = m.G398
						if v40 != v83 {
							v86 = *(*int32)(unsafe.Add(mBase, uint32(v40)+8))
							v89 = v86
						} else {
							v89 = int32(-1)
						}
					} else {
						v89 = int32(-1)
					}
					if v89 == int32(-1) {
						v109 = int32(8)
						v111 = int32(-2147483647)
						if v111 < v109 {
							v114 = v109
						} else {
							v114 = v111
						}
						v117 = base.B2i32(v109 < int32(0))
						if v109 < int32(0) {
							v118 = int32(0) - v114
						} else {
							v118 = v114
						}
						v119 = int32(8)
						if v118 < v119 {
							v122 = v118
						} else {
							v122 = v119
						}
						v123 = m.G3
						if v118 < int32(1) {
						} else {
							v128 = v122 + (v11 + int32(8))
							if v109 < int32(0) {
								v135 = v123 + int32(_a2761)
							} else {
								v135 = v123 + int32(_a2762)
							}
							v139 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v135+v21&int32(15)))))
							*(*uint8)(unsafe.Add(mBase, uint32(v128+int32(-1)))) = uint8(v139)
							if v118 == int32(1) {
							} else {
								v150 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v135+int32(base.Ui32(v21)>>(uint(int32(4))%32))&int32(15)))))
								*(*uint8)(unsafe.Add(mBase, uint32(v128+int32(-2)))) = uint8(v150)
								if v118 < int32(3) {
								} else {
									v161 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v135+int32(base.Ui32(v21)>>(uint(int32(8))%32))&int32(15)))))
									*(*uint8)(unsafe.Add(mBase, uint32(v128+int32(-3)))) = uint8(v161)
									if v118 == int32(3) {
									} else {
										v172 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v135+int32(base.Ui32(v21)>>(uint(int32(12))%32))&int32(15)))))
										*(*uint8)(unsafe.Add(mBase, uint32(v128+int32(-4)))) = uint8(v172)
										if v118 < int32(5) {
										} else {
											v183 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v135+int32(base.Ui32(v21)>>(uint(int32(16))%32))&int32(15)))))
											*(*uint8)(unsafe.Add(mBase, uint32(v128+int32(-5)))) = uint8(v183)
											if v118 == int32(5) {
											} else {
												v194 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v135+int32(base.Ui32(v21)>>(uint(int32(20))%32))&int32(15)))))
												*(*uint8)(unsafe.Add(mBase, uint32(v128+int32(-6)))) = uint8(v194)
												if v118 < int32(7) {
												} else {
													v205 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v135+int32(base.Ui32(v21)>>(uint(int32(24))%32))&int32(15)))))
													*(*uint8)(unsafe.Add(mBase, uint32(v128+int32(-7)))) = uint8(v205)
													if v118 == int32(7) {
													} else {
														v214 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v135+int32(base.Ui32(v21)>>(uint(int32(28))%32))))))
														*(*uint8)(unsafe.Add(mBase, uint32(v128+int32(-8)))) = uint8(v214)
													}
												}
											}
										}
									}
								}
							}
						}
						F_lua_pushlstring(m, l0, v11+int32(8), v122)
						mBase = m.M
						v221 = m.ExcPending
						if v221 != 0 {
							return int32(0)
						} else {
							m.G0 = v11 + int32(16)
							return int32(1)
						}
					} else {
						v93 = F_lua_tonumber(m, l0, int32(2))
						mBase = m.M
						v94 = m.ExcPending
						if v94 != 0 {
							return int32(0)
						} else {
							v98 = base.I32_wrap_i64(base.I64_reinterpret_f64(base.F64_add(v93, float64(6.755399441055744e+15))))
							if v98 != 0 {
								v109 = v98
								v111 = int32(-2147483647)
								if v111 < v109 {
									v114 = v109
								} else {
									v114 = v111
								}
								v117 = base.B2i32(v109 < int32(0))
								if v109 < int32(0) {
									v118 = int32(0) - v114
								} else {
									v118 = v114
								}
								v119 = int32(8)
								if v118 < v119 {
									v122 = v118
								} else {
									v122 = v119
								}
								v123 = m.G3
								if v118 < int32(1) {
								} else {
									v128 = v122 + (v11 + int32(8))
									if v109 < int32(0) {
										v135 = v123 + int32(_a2761)
									} else {
										v135 = v123 + int32(_a2762)
									}
									v139 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v135+v21&int32(15)))))
									*(*uint8)(unsafe.Add(mBase, uint32(v128+int32(-1)))) = uint8(v139)
									if v118 == int32(1) {
									} else {
										v150 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v135+int32(base.Ui32(v21)>>(uint(int32(4))%32))&int32(15)))))
										*(*uint8)(unsafe.Add(mBase, uint32(v128+int32(-2)))) = uint8(v150)
										if v118 < int32(3) {
										} else {
											v161 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v135+int32(base.Ui32(v21)>>(uint(int32(8))%32))&int32(15)))))
											*(*uint8)(unsafe.Add(mBase, uint32(v128+int32(-3)))) = uint8(v161)
											if v118 == int32(3) {
											} else {
												v172 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v135+int32(base.Ui32(v21)>>(uint(int32(12))%32))&int32(15)))))
												*(*uint8)(unsafe.Add(mBase, uint32(v128+int32(-4)))) = uint8(v172)
												if v118 < int32(5) {
												} else {
													v183 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v135+int32(base.Ui32(v21)>>(uint(int32(16))%32))&int32(15)))))
													*(*uint8)(unsafe.Add(mBase, uint32(v128+int32(-5)))) = uint8(v183)
													if v118 == int32(5) {
													} else {
														v194 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v135+int32(base.Ui32(v21)>>(uint(int32(20))%32))&int32(15)))))
														*(*uint8)(unsafe.Add(mBase, uint32(v128+int32(-6)))) = uint8(v194)
														if v118 < int32(7) {
														} else {
															v205 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v135+int32(base.Ui32(v21)>>(uint(int32(24))%32))&int32(15)))))
															*(*uint8)(unsafe.Add(mBase, uint32(v128+int32(-7)))) = uint8(v205)
															if v118 == int32(7) {
															} else {
																v214 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v135+int32(base.Ui32(v21)>>(uint(int32(28))%32))))))
																*(*uint8)(unsafe.Add(mBase, uint32(v128+int32(-8)))) = uint8(v214)
															}
														}
													}
												}
											}
										}
									}
								}
								F_lua_pushlstring(m, l0, v11+int32(8), v122)
								mBase = m.M
								v221 = m.ExcPending
								if v221 != 0 {
									return int32(0)
								} else {
									m.G0 = v11 + int32(16)
									return int32(1)
								}
							} else {
								v99 = int32(0)
								v101 = F_lua_isnumber(m, l0, int32(2))
								mBase = m.M
								v102 = m.ExcPending
								if v102 != 0 {
									return int32(0)
								} else {
									if v101 != 0 {
										v109 = v99
										v111 = int32(-2147483647)
										if v111 < v109 {
											v114 = v109
										} else {
											v114 = v111
										}
										v117 = base.B2i32(v109 < int32(0))
										if v109 < int32(0) {
											v118 = int32(0) - v114
										} else {
											v118 = v114
										}
										v119 = int32(8)
										if v118 < v119 {
											v122 = v118
										} else {
											v122 = v119
										}
										v123 = m.G3
										if v118 < int32(1) {
										} else {
											v128 = v122 + (v11 + int32(8))
											if v109 < int32(0) {
												v135 = v123 + int32(_a2761)
											} else {
												v135 = v123 + int32(_a2762)
											}
											v139 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v135+v21&int32(15)))))
											*(*uint8)(unsafe.Add(mBase, uint32(v128+int32(-1)))) = uint8(v139)
											if v118 == int32(1) {
											} else {
												v150 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v135+int32(base.Ui32(v21)>>(uint(int32(4))%32))&int32(15)))))
												*(*uint8)(unsafe.Add(mBase, uint32(v128+int32(-2)))) = uint8(v150)
												if v118 < int32(3) {
												} else {
													v161 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v135+int32(base.Ui32(v21)>>(uint(int32(8))%32))&int32(15)))))
													*(*uint8)(unsafe.Add(mBase, uint32(v128+int32(-3)))) = uint8(v161)
													if v118 == int32(3) {
													} else {
														v172 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v135+int32(base.Ui32(v21)>>(uint(int32(12))%32))&int32(15)))))
														*(*uint8)(unsafe.Add(mBase, uint32(v128+int32(-4)))) = uint8(v172)
														if v118 < int32(5) {
														} else {
															v183 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v135+int32(base.Ui32(v21)>>(uint(int32(16))%32))&int32(15)))))
															*(*uint8)(unsafe.Add(mBase, uint32(v128+int32(-5)))) = uint8(v183)
															if v118 == int32(5) {
															} else {
																v194 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v135+int32(base.Ui32(v21)>>(uint(int32(20))%32))&int32(15)))))
																*(*uint8)(unsafe.Add(mBase, uint32(v128+int32(-6)))) = uint8(v194)
																if v118 < int32(7) {
																} else {
																	v205 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v135+int32(base.Ui32(v21)>>(uint(int32(24))%32))&int32(15)))))
																	*(*uint8)(unsafe.Add(mBase, uint32(v128+int32(-7)))) = uint8(v205)
																	if v118 == int32(7) {
																	} else {
																		v214 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v135+int32(base.Ui32(v21)>>(uint(int32(28))%32))))))
																		*(*uint8)(unsafe.Add(mBase, uint32(v128+int32(-8)))) = uint8(v214)
																	}
																}
															}
														}
													}
												}
											}
										}
										F_lua_pushlstring(m, l0, v11+int32(8), v122)
										mBase = m.M
										v221 = m.ExcPending
										if v221 != 0 {
											return int32(0)
										} else {
											m.G0 = v11 + int32(16)
											return int32(1)
										}
									} else {
										v104 = m.G3
										v107 = F_luaL_typerror(m, l0, int32(2), v104+int32(_a2760))
										mBase = m.M
										v108 = m.ExcPending
										if v108 != 0 {
											return int32(0)
										} else {
											v109 = v99
											v111 = int32(-2147483647)
											if v111 < v109 {
												v114 = v109
											} else {
												v114 = v111
											}
											v117 = base.B2i32(v109 < int32(0))
											if v109 < int32(0) {
												v118 = int32(0) - v114
											} else {
												v118 = v114
											}
											v119 = int32(8)
											if v118 < v119 {
												v122 = v118
											} else {
												v122 = v119
											}
											v123 = m.G3
											if v118 < int32(1) {
											} else {
												v128 = v122 + (v11 + int32(8))
												if v109 < int32(0) {
													v135 = v123 + int32(_a2761)
												} else {
													v135 = v123 + int32(_a2762)
												}
												v139 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v135+v21&int32(15)))))
												*(*uint8)(unsafe.Add(mBase, uint32(v128+int32(-1)))) = uint8(v139)
												if v118 == int32(1) {
												} else {
													v150 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v135+int32(base.Ui32(v21)>>(uint(int32(4))%32))&int32(15)))))
													*(*uint8)(unsafe.Add(mBase, uint32(v128+int32(-2)))) = uint8(v150)
													if v118 < int32(3) {
													} else {
														v161 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v135+int32(base.Ui32(v21)>>(uint(int32(8))%32))&int32(15)))))
														*(*uint8)(unsafe.Add(mBase, uint32(v128+int32(-3)))) = uint8(v161)
														if v118 == int32(3) {
														} else {
															v172 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v135+int32(base.Ui32(v21)>>(uint(int32(12))%32))&int32(15)))))
															*(*uint8)(unsafe.Add(mBase, uint32(v128+int32(-4)))) = uint8(v172)
															if v118 < int32(5) {
															} else {
																v183 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v135+int32(base.Ui32(v21)>>(uint(int32(16))%32))&int32(15)))))
																*(*uint8)(unsafe.Add(mBase, uint32(v128+int32(-5)))) = uint8(v183)
																if v118 == int32(5) {
																} else {
																	v194 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v135+int32(base.Ui32(v21)>>(uint(int32(20))%32))&int32(15)))))
																	*(*uint8)(unsafe.Add(mBase, uint32(v128+int32(-6)))) = uint8(v194)
																	if v118 < int32(7) {
																	} else {
																		v205 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v135+int32(base.Ui32(v21)>>(uint(int32(24))%32))&int32(15)))))
																		*(*uint8)(unsafe.Add(mBase, uint32(v128+int32(-7)))) = uint8(v205)
																		if v118 == int32(7) {
																		} else {
																			v214 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v135+int32(base.Ui32(v21)>>(uint(int32(28))%32))))))
																			*(*uint8)(unsafe.Add(mBase, uint32(v128+int32(-8)))) = uint8(v214)
																		}
																	}
																}
															}
														}
													}
												}
											}
											F_lua_pushlstring(m, l0, v11+int32(8), v122)
											mBase = m.M
											v221 = m.ExcPending
											if v221 != 0 {
												return int32(0)
											} else {
												m.G0 = v11 + int32(16)
												return int32(1)
											}
										}
									}
								}
							}
						}
					}
				} else {
					v26 = m.G3
					v29 = F_luaL_typerror(m, l0, int32(1), v26+int32(_a2760))
					mBase = m.M
					v30 = m.ExcPending
					if v30 != 0 {
						return int32(0)
					} else {
						v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
						v40 = v35 + int32(16)
						v41 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
						if base.Ui32(v40) < base.Ui32(v41) {
							v83 = m.G398
							if v40 != v83 {
								v86 = *(*int32)(unsafe.Add(mBase, uint32(v40)+8))
								v89 = v86
							} else {
								v89 = int32(-1)
							}
						} else {
							v89 = int32(-1)
						}
						if v89 == int32(-1) {
							v109 = int32(8)
							v111 = int32(-2147483647)
							if v111 < v109 {
								v114 = v109
							} else {
								v114 = v111
							}
							v117 = base.B2i32(v109 < int32(0))
							if v109 < int32(0) {
								v118 = int32(0) - v114
							} else {
								v118 = v114
							}
							v119 = int32(8)
							if v118 < v119 {
								v122 = v118
							} else {
								v122 = v119
							}
							v123 = m.G3
							if v118 < int32(1) {
							} else {
								v128 = v122 + (v11 + int32(8))
								if v109 < int32(0) {
									v135 = v123 + int32(_a2761)
								} else {
									v135 = v123 + int32(_a2762)
								}
								v139 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v135+v21&int32(15)))))
								*(*uint8)(unsafe.Add(mBase, uint32(v128+int32(-1)))) = uint8(v139)
								if v118 == int32(1) {
								} else {
									v150 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v135+int32(base.Ui32(v21)>>(uint(int32(4))%32))&int32(15)))))
									*(*uint8)(unsafe.Add(mBase, uint32(v128+int32(-2)))) = uint8(v150)
									if v118 < int32(3) {
									} else {
										v161 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v135+int32(base.Ui32(v21)>>(uint(int32(8))%32))&int32(15)))))
										*(*uint8)(unsafe.Add(mBase, uint32(v128+int32(-3)))) = uint8(v161)
										if v118 == int32(3) {
										} else {
											v172 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v135+int32(base.Ui32(v21)>>(uint(int32(12))%32))&int32(15)))))
											*(*uint8)(unsafe.Add(mBase, uint32(v128+int32(-4)))) = uint8(v172)
											if v118 < int32(5) {
											} else {
												v183 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v135+int32(base.Ui32(v21)>>(uint(int32(16))%32))&int32(15)))))
												*(*uint8)(unsafe.Add(mBase, uint32(v128+int32(-5)))) = uint8(v183)
												if v118 == int32(5) {
												} else {
													v194 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v135+int32(base.Ui32(v21)>>(uint(int32(20))%32))&int32(15)))))
													*(*uint8)(unsafe.Add(mBase, uint32(v128+int32(-6)))) = uint8(v194)
													if v118 < int32(7) {
													} else {
														v205 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v135+int32(base.Ui32(v21)>>(uint(int32(24))%32))&int32(15)))))
														*(*uint8)(unsafe.Add(mBase, uint32(v128+int32(-7)))) = uint8(v205)
														if v118 == int32(7) {
														} else {
															v214 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v135+int32(base.Ui32(v21)>>(uint(int32(28))%32))))))
															*(*uint8)(unsafe.Add(mBase, uint32(v128+int32(-8)))) = uint8(v214)
														}
													}
												}
											}
										}
									}
								}
							}
							F_lua_pushlstring(m, l0, v11+int32(8), v122)
							mBase = m.M
							v221 = m.ExcPending
							if v221 != 0 {
								return int32(0)
							} else {
								m.G0 = v11 + int32(16)
								return int32(1)
							}
						} else {
							v93 = F_lua_tonumber(m, l0, int32(2))
							mBase = m.M
							v94 = m.ExcPending
							if v94 != 0 {
								return int32(0)
							} else {
								v98 = base.I32_wrap_i64(base.I64_reinterpret_f64(base.F64_add(v93, float64(6.755399441055744e+15))))
								if v98 != 0 {
									v109 = v98
									v111 = int32(-2147483647)
									if v111 < v109 {
										v114 = v109
									} else {
										v114 = v111
									}
									v117 = base.B2i32(v109 < int32(0))
									if v109 < int32(0) {
										v118 = int32(0) - v114
									} else {
										v118 = v114
									}
									v119 = int32(8)
									if v118 < v119 {
										v122 = v118
									} else {
										v122 = v119
									}
									v123 = m.G3
									if v118 < int32(1) {
									} else {
										v128 = v122 + (v11 + int32(8))
										if v109 < int32(0) {
											v135 = v123 + int32(_a2761)
										} else {
											v135 = v123 + int32(_a2762)
										}
										v139 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v135+v21&int32(15)))))
										*(*uint8)(unsafe.Add(mBase, uint32(v128+int32(-1)))) = uint8(v139)
										if v118 == int32(1) {
										} else {
											v150 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v135+int32(base.Ui32(v21)>>(uint(int32(4))%32))&int32(15)))))
											*(*uint8)(unsafe.Add(mBase, uint32(v128+int32(-2)))) = uint8(v150)
											if v118 < int32(3) {
											} else {
												v161 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v135+int32(base.Ui32(v21)>>(uint(int32(8))%32))&int32(15)))))
												*(*uint8)(unsafe.Add(mBase, uint32(v128+int32(-3)))) = uint8(v161)
												if v118 == int32(3) {
												} else {
													v172 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v135+int32(base.Ui32(v21)>>(uint(int32(12))%32))&int32(15)))))
													*(*uint8)(unsafe.Add(mBase, uint32(v128+int32(-4)))) = uint8(v172)
													if v118 < int32(5) {
													} else {
														v183 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v135+int32(base.Ui32(v21)>>(uint(int32(16))%32))&int32(15)))))
														*(*uint8)(unsafe.Add(mBase, uint32(v128+int32(-5)))) = uint8(v183)
														if v118 == int32(5) {
														} else {
															v194 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v135+int32(base.Ui32(v21)>>(uint(int32(20))%32))&int32(15)))))
															*(*uint8)(unsafe.Add(mBase, uint32(v128+int32(-6)))) = uint8(v194)
															if v118 < int32(7) {
															} else {
																v205 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v135+int32(base.Ui32(v21)>>(uint(int32(24))%32))&int32(15)))))
																*(*uint8)(unsafe.Add(mBase, uint32(v128+int32(-7)))) = uint8(v205)
																if v118 == int32(7) {
																} else {
																	v214 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v135+int32(base.Ui32(v21)>>(uint(int32(28))%32))))))
																	*(*uint8)(unsafe.Add(mBase, uint32(v128+int32(-8)))) = uint8(v214)
																}
															}
														}
													}
												}
											}
										}
									}
									F_lua_pushlstring(m, l0, v11+int32(8), v122)
									mBase = m.M
									v221 = m.ExcPending
									if v221 != 0 {
										return int32(0)
									} else {
										m.G0 = v11 + int32(16)
										return int32(1)
									}
								} else {
									v99 = int32(0)
									v101 = F_lua_isnumber(m, l0, int32(2))
									mBase = m.M
									v102 = m.ExcPending
									if v102 != 0 {
										return int32(0)
									} else {
										if v101 != 0 {
											v109 = v99
											v111 = int32(-2147483647)
											if v111 < v109 {
												v114 = v109
											} else {
												v114 = v111
											}
											v117 = base.B2i32(v109 < int32(0))
											if v109 < int32(0) {
												v118 = int32(0) - v114
											} else {
												v118 = v114
											}
											v119 = int32(8)
											if v118 < v119 {
												v122 = v118
											} else {
												v122 = v119
											}
											v123 = m.G3
											if v118 < int32(1) {
											} else {
												v128 = v122 + (v11 + int32(8))
												if v109 < int32(0) {
													v135 = v123 + int32(_a2761)
												} else {
													v135 = v123 + int32(_a2762)
												}
												v139 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v135+v21&int32(15)))))
												*(*uint8)(unsafe.Add(mBase, uint32(v128+int32(-1)))) = uint8(v139)
												if v118 == int32(1) {
												} else {
													v150 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v135+int32(base.Ui32(v21)>>(uint(int32(4))%32))&int32(15)))))
													*(*uint8)(unsafe.Add(mBase, uint32(v128+int32(-2)))) = uint8(v150)
													if v118 < int32(3) {
													} else {
														v161 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v135+int32(base.Ui32(v21)>>(uint(int32(8))%32))&int32(15)))))
														*(*uint8)(unsafe.Add(mBase, uint32(v128+int32(-3)))) = uint8(v161)
														if v118 == int32(3) {
														} else {
															v172 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v135+int32(base.Ui32(v21)>>(uint(int32(12))%32))&int32(15)))))
															*(*uint8)(unsafe.Add(mBase, uint32(v128+int32(-4)))) = uint8(v172)
															if v118 < int32(5) {
															} else {
																v183 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v135+int32(base.Ui32(v21)>>(uint(int32(16))%32))&int32(15)))))
																*(*uint8)(unsafe.Add(mBase, uint32(v128+int32(-5)))) = uint8(v183)
																if v118 == int32(5) {
																} else {
																	v194 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v135+int32(base.Ui32(v21)>>(uint(int32(20))%32))&int32(15)))))
																	*(*uint8)(unsafe.Add(mBase, uint32(v128+int32(-6)))) = uint8(v194)
																	if v118 < int32(7) {
																	} else {
																		v205 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v135+int32(base.Ui32(v21)>>(uint(int32(24))%32))&int32(15)))))
																		*(*uint8)(unsafe.Add(mBase, uint32(v128+int32(-7)))) = uint8(v205)
																		if v118 == int32(7) {
																		} else {
																			v214 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v135+int32(base.Ui32(v21)>>(uint(int32(28))%32))))))
																			*(*uint8)(unsafe.Add(mBase, uint32(v128+int32(-8)))) = uint8(v214)
																		}
																	}
																}
															}
														}
													}
												}
											}
											F_lua_pushlstring(m, l0, v11+int32(8), v122)
											mBase = m.M
											v221 = m.ExcPending
											if v221 != 0 {
												return int32(0)
											} else {
												m.G0 = v11 + int32(16)
												return int32(1)
											}
										} else {
											v104 = m.G3
											v107 = F_luaL_typerror(m, l0, int32(2), v104+int32(_a2760))
											mBase = m.M
											v108 = m.ExcPending
											if v108 != 0 {
												return int32(0)
											} else {
												v109 = v99
												v111 = int32(-2147483647)
												if v111 < v109 {
													v114 = v109
												} else {
													v114 = v111
												}
												v117 = base.B2i32(v109 < int32(0))
												if v109 < int32(0) {
													v118 = int32(0) - v114
												} else {
													v118 = v114
												}
												v119 = int32(8)
												if v118 < v119 {
													v122 = v118
												} else {
													v122 = v119
												}
												v123 = m.G3
												if v118 < int32(1) {
												} else {
													v128 = v122 + (v11 + int32(8))
													if v109 < int32(0) {
														v135 = v123 + int32(_a2761)
													} else {
														v135 = v123 + int32(_a2762)
													}
													v139 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v135+v21&int32(15)))))
													*(*uint8)(unsafe.Add(mBase, uint32(v128+int32(-1)))) = uint8(v139)
													if v118 == int32(1) {
													} else {
														v150 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v135+int32(base.Ui32(v21)>>(uint(int32(4))%32))&int32(15)))))
														*(*uint8)(unsafe.Add(mBase, uint32(v128+int32(-2)))) = uint8(v150)
														if v118 < int32(3) {
														} else {
															v161 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v135+int32(base.Ui32(v21)>>(uint(int32(8))%32))&int32(15)))))
															*(*uint8)(unsafe.Add(mBase, uint32(v128+int32(-3)))) = uint8(v161)
															if v118 == int32(3) {
															} else {
																v172 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v135+int32(base.Ui32(v21)>>(uint(int32(12))%32))&int32(15)))))
																*(*uint8)(unsafe.Add(mBase, uint32(v128+int32(-4)))) = uint8(v172)
																if v118 < int32(5) {
																} else {
																	v183 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v135+int32(base.Ui32(v21)>>(uint(int32(16))%32))&int32(15)))))
																	*(*uint8)(unsafe.Add(mBase, uint32(v128+int32(-5)))) = uint8(v183)
																	if v118 == int32(5) {
																	} else {
																		v194 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v135+int32(base.Ui32(v21)>>(uint(int32(20))%32))&int32(15)))))
																		*(*uint8)(unsafe.Add(mBase, uint32(v128+int32(-6)))) = uint8(v194)
																		if v118 < int32(7) {
																		} else {
																			v205 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v135+int32(base.Ui32(v21)>>(uint(int32(24))%32))&int32(15)))))
																			*(*uint8)(unsafe.Add(mBase, uint32(v128+int32(-7)))) = uint8(v205)
																			if v118 == int32(7) {
																			} else {
																				v214 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v135+int32(base.Ui32(v21)>>(uint(int32(28))%32))))))
																				*(*uint8)(unsafe.Add(mBase, uint32(v128+int32(-8)))) = uint8(v214)
																			}
																		}
																	}
																}
															}
														}
													}
												}
												F_lua_pushlstring(m, l0, v11+int32(8), v122)
												mBase = m.M
												v221 = m.ExcPending
												if v221 != 0 {
													return int32(0)
												} else {
													m.G0 = v11 + int32(16)
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
