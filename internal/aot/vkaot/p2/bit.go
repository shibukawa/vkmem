package p2

import (
	base "github.com/shibukawa/vkmem/internal/aot/vkaot/base"
	"unsafe"
)

func F_bit_band(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 float64
	_ = v7
	var v10 int32
	_ = v10
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
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
	var v28 int32
	_ = v28
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v36 float64
	_ = v36
	var v37 int32
	_ = v37
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v55 int32
	_ = v55
	var v61 int32
	_ = v61
	var v65 int32
	_ = v65
	v7 = F_lua_tonumber(m, l0, int32(1))
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v28 = (v24 - v25) >> (uint(int32(4)) % 32)
	goto L9
L2:
	;
	return int32(0)
L3:
	;
	v14 = base.I32_wrap_i64(base.I64_reinterpret_f64(base.F64_add(v7, float64(6.755399441055744e+15))))
	if v14 != 0 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v16 = F_lua_isnumber(m, l0, int32(1))
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		goto L2
	} else {
		goto L5
	}
L5:
	;
	if v16 != 0 {
		goto L1
	} else {
		goto L6
	}
L6:
	;
	v19 = m.G3
	v22 = F_luaL_typerror(m, l0, int32(1), v19+int32(_a_F_bit_band_0))
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		goto L2
	} else {
		goto L7
	}
L7:
	;
	goto L1
L8:
	;
	v61 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v61)+8)) = int32(3)
	*(*float64)(unsafe.Add(mBase, uint32(v61))) = base.F64_convert_i32_s(v55)
	v65 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v65 + int32(16)
	goto L20
L9:
	;
	if v28 < int32(2) {
		v55 = v14
		goto L8
	} else {
		goto L10
	}
L10:
	;
	v32 = v14
	v33 = v28
	goto L11
L11:
	;
	v36 = F_lua_tonumber(m, l0, v33)
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		goto L2
	} else {
		goto L14
	}
L12:
	;
	v55 = v51
	goto L8
L13:
	;
	v51 = v32 & v41
	if int32(2) < v33 {
		v32 = v51
		v33 = v33 + int32(-1)
		goto L11
	} else {
		goto L19
	}
L14:
	;
	v41 = base.I32_wrap_i64(base.I64_reinterpret_f64(base.F64_add(v36, float64(6.755399441055744e+15))))
	if v41 != 0 {
		goto L13
	} else {
		goto L15
	}
L15:
	;
	v42 = F_lua_isnumber(m, l0, v33)
	mBase = m.M
	v43 = m.ExcPending
	if v43 != 0 {
		goto L2
	} else {
		goto L16
	}
L16:
	;
	if v42 != 0 {
		goto L13
	} else {
		goto L17
	}
L17:
	;
	v44 = m.G3
	v47 = F_luaL_typerror(m, l0, v33, v44+int32(_a_F_bit_band_0))
	mBase = m.M
	v48 = m.ExcPending
	if v48 != 0 {
		goto L2
	} else {
		goto L18
	}
L18:
	;
	goto L13
L19:
	;
	goto L12
L20:
	;
	return int32(1)
}
func F_bit_bor(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 float64
	_ = v7
	var v10 int32
	_ = v10
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
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
	var v28 int32
	_ = v28
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v36 float64
	_ = v36
	var v37 int32
	_ = v37
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v55 int32
	_ = v55
	var v61 int32
	_ = v61
	var v65 int32
	_ = v65
	v7 = F_lua_tonumber(m, l0, int32(1))
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v28 = (v24 - v25) >> (uint(int32(4)) % 32)
	goto L9
L2:
	;
	return int32(0)
L3:
	;
	v14 = base.I32_wrap_i64(base.I64_reinterpret_f64(base.F64_add(v7, float64(6.755399441055744e+15))))
	if v14 != 0 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v16 = F_lua_isnumber(m, l0, int32(1))
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		goto L2
	} else {
		goto L5
	}
L5:
	;
	if v16 != 0 {
		goto L1
	} else {
		goto L6
	}
L6:
	;
	v19 = m.G3
	v22 = F_luaL_typerror(m, l0, int32(1), v19+int32(_a_F_bit_bor_0))
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		goto L2
	} else {
		goto L7
	}
L7:
	;
	goto L1
L8:
	;
	v61 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v61)+8)) = int32(3)
	*(*float64)(unsafe.Add(mBase, uint32(v61))) = base.F64_convert_i32_s(v55)
	v65 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v65 + int32(16)
	goto L20
L9:
	;
	if v28 < int32(2) {
		v55 = v14
		goto L8
	} else {
		goto L10
	}
L10:
	;
	v32 = v14
	v33 = v28
	goto L11
L11:
	;
	v36 = F_lua_tonumber(m, l0, v33)
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		goto L2
	} else {
		goto L14
	}
L12:
	;
	v55 = v51
	goto L8
L13:
	;
	v51 = v32 | v41
	if int32(2) < v33 {
		v32 = v51
		v33 = v33 + int32(-1)
		goto L11
	} else {
		goto L19
	}
L14:
	;
	v41 = base.I32_wrap_i64(base.I64_reinterpret_f64(base.F64_add(v36, float64(6.755399441055744e+15))))
	if v41 != 0 {
		goto L13
	} else {
		goto L15
	}
L15:
	;
	v42 = F_lua_isnumber(m, l0, v33)
	mBase = m.M
	v43 = m.ExcPending
	if v43 != 0 {
		goto L2
	} else {
		goto L16
	}
L16:
	;
	if v42 != 0 {
		goto L13
	} else {
		goto L17
	}
L17:
	;
	v44 = m.G3
	v47 = F_luaL_typerror(m, l0, v33, v44+int32(_a_F_bit_bor_0))
	mBase = m.M
	v48 = m.ExcPending
	if v48 != 0 {
		goto L2
	} else {
		goto L18
	}
L18:
	;
	goto L13
L19:
	;
	goto L12
L20:
	;
	return int32(1)
}
func F_bit_bxor(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 float64
	_ = v7
	var v10 int32
	_ = v10
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
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
	var v28 int32
	_ = v28
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v36 float64
	_ = v36
	var v37 int32
	_ = v37
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v55 int32
	_ = v55
	var v61 int32
	_ = v61
	var v65 int32
	_ = v65
	v7 = F_lua_tonumber(m, l0, int32(1))
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v28 = (v24 - v25) >> (uint(int32(4)) % 32)
	goto L9
L2:
	;
	return int32(0)
L3:
	;
	v14 = base.I32_wrap_i64(base.I64_reinterpret_f64(base.F64_add(v7, float64(6.755399441055744e+15))))
	if v14 != 0 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v16 = F_lua_isnumber(m, l0, int32(1))
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		goto L2
	} else {
		goto L5
	}
L5:
	;
	if v16 != 0 {
		goto L1
	} else {
		goto L6
	}
L6:
	;
	v19 = m.G3
	v22 = F_luaL_typerror(m, l0, int32(1), v19+int32(_a_F_bit_bxor_0))
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		goto L2
	} else {
		goto L7
	}
L7:
	;
	goto L1
L8:
	;
	v61 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v61)+8)) = int32(3)
	*(*float64)(unsafe.Add(mBase, uint32(v61))) = base.F64_convert_i32_s(v55)
	v65 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v65 + int32(16)
	goto L20
L9:
	;
	if v28 < int32(2) {
		v55 = v14
		goto L8
	} else {
		goto L10
	}
L10:
	;
	v32 = v14
	v33 = v28
	goto L11
L11:
	;
	v36 = F_lua_tonumber(m, l0, v33)
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		goto L2
	} else {
		goto L14
	}
L12:
	;
	v55 = v51
	goto L8
L13:
	;
	v51 = v32 ^ v41
	if int32(2) < v33 {
		v32 = v51
		v33 = v33 + int32(-1)
		goto L11
	} else {
		goto L19
	}
L14:
	;
	v41 = base.I32_wrap_i64(base.I64_reinterpret_f64(base.F64_add(v36, float64(6.755399441055744e+15))))
	if v41 != 0 {
		goto L13
	} else {
		goto L15
	}
L15:
	;
	v42 = F_lua_isnumber(m, l0, v33)
	mBase = m.M
	v43 = m.ExcPending
	if v43 != 0 {
		goto L2
	} else {
		goto L16
	}
L16:
	;
	if v42 != 0 {
		goto L13
	} else {
		goto L17
	}
L17:
	;
	v44 = m.G3
	v47 = F_luaL_typerror(m, l0, v33, v44+int32(_a_F_bit_bxor_0))
	mBase = m.M
	v48 = m.ExcPending
	if v48 != 0 {
		goto L2
	} else {
		goto L18
	}
L18:
	;
	goto L13
L19:
	;
	goto L12
L20:
	;
	return int32(1)
}
func F_bit_rol(m *base.Module, l0 int32) int32 {
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
					*(*float64)(unsafe.Add(mBase, uint32(v41))) = base.F64_convert_i32_s(base.I32_rotl(v12, v28))
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
							*(*float64)(unsafe.Add(mBase, uint32(v41))) = base.F64_convert_i32_s(base.I32_rotl(v12, v28))
							v45 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
							*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v45 + int32(16)
							return int32(1)
						} else {
							v33 = m.G3
							v36 = F_luaL_typerror(m, l0, int32(2), v33+int32(_a_F_bit_rol_0))
							mBase = m.M
							v37 = m.ExcPending
							if v37 != 0 {
								return int32(0)
							} else {
								v41 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
								*(*int32)(unsafe.Add(mBase, uint32(v41)+8)) = int32(3)
								*(*float64)(unsafe.Add(mBase, uint32(v41))) = base.F64_convert_i32_s(base.I32_rotl(v12, v28))
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
							*(*float64)(unsafe.Add(mBase, uint32(v41))) = base.F64_convert_i32_s(base.I32_rotl(v12, v28))
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
									*(*float64)(unsafe.Add(mBase, uint32(v41))) = base.F64_convert_i32_s(base.I32_rotl(v12, v28))
									v45 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
									*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v45 + int32(16)
									return int32(1)
								} else {
									v33 = m.G3
									v36 = F_luaL_typerror(m, l0, int32(2), v33+int32(_a_F_bit_rol_0))
									mBase = m.M
									v37 = m.ExcPending
									if v37 != 0 {
										return int32(0)
									} else {
										v41 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
										*(*int32)(unsafe.Add(mBase, uint32(v41)+8)) = int32(3)
										*(*float64)(unsafe.Add(mBase, uint32(v41))) = base.F64_convert_i32_s(base.I32_rotl(v12, v28))
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
					v20 = F_luaL_typerror(m, l0, int32(1), v17+int32(_a_F_bit_rol_0))
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
								*(*float64)(unsafe.Add(mBase, uint32(v41))) = base.F64_convert_i32_s(base.I32_rotl(v12, v28))
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
										*(*float64)(unsafe.Add(mBase, uint32(v41))) = base.F64_convert_i32_s(base.I32_rotl(v12, v28))
										v45 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
										*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v45 + int32(16)
										return int32(1)
									} else {
										v33 = m.G3
										v36 = F_luaL_typerror(m, l0, int32(2), v33+int32(_a_F_bit_rol_0))
										mBase = m.M
										v37 = m.ExcPending
										if v37 != 0 {
											return int32(0)
										} else {
											v41 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
											*(*int32)(unsafe.Add(mBase, uint32(v41)+8)) = int32(3)
											*(*float64)(unsafe.Add(mBase, uint32(v41))) = base.F64_convert_i32_s(base.I32_rotl(v12, v28))
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
func F_getBitOffsetFromArgument(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v66 int32
	_ = v66
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v97 int64
	_ = v97
	var v103 int32
	_ = v103
	var v105 int64
	_ = v105
	var v107 int32
	_ = v107
	var v109 int32
	_ = v109
	var v119 int64
	_ = v119
	var v124 int64
	_ = v124
	var v128 int32
	_ = v128
	var v130 int64
	_ = v130
	var v132 int32
	_ = v132
	var v140 int64
	_ = v140
	var v164 int64
	_ = v164
	var v183 int32
	_ = v183
	var v194 int32
	_ = v194
	var v198 int64
	_ = v198
	var v200 int64
	_ = v200
	var v202 int64
	_ = v202
	var v203 int64
	_ = v203
	var v207 int64
	_ = v207
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v215 int32
	_ = v215
	var v217 int32
	_ = v217
	var v219 int32
	_ = v219
	var v222 int32
	_ = v222
	var v223 int64
	_ = v223
	var v227 int64
	_ = v227
	var v233 int32
	_ = v233
	var v237 int32
	_ = v237
	v10 = m.G0
	v12 = v10 - int32(16)
	m.G0 = v12
	v15 = F_objectGetVal(m, l1)
	mBase = m.M
	v18 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15+int32(-1)))))
	switch v18 & int32(7) {
	case 0:
		goto L6
	case 1:
		goto L5
	case 2:
		goto L4
	case 3:
		goto L3
	case 4:
		goto L2
	default:
		v35 = int32(0)
		goto L1
	}
L1:
	;
	v36 = int32(0)
	v40 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15))))
	v44 = base.B2i32(v36 < l4) & (base.B2i32(l3 != v36) & base.B2i32(v40 == int32(35)))
	v45 = v15 + v44
	v46 = v35 - v44
	v48 = v12 + int32(8)
	if base.Ui32(v46+int32(-21)) < base.Ui32(int32(-20)) {
		v183 = v36
		goto L10
	} else {
		goto L11
	}
L2:
	;
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v15+int32(-17))))
	v35 = v34
	goto L1
L3:
	;
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v15+int32(-9))))
	v35 = v31
	goto L1
L4:
	;
	v28 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v15+int32(-5)))))
	v35 = v28
	goto L1
L5:
	;
	v25 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15+int32(-3)))))
	v35 = v25
	goto L1
L6:
	;
	v35 = int32(base.Ui32(v18) >> (uint(int32(3)) % 32))
	goto L1
L7:
	;
	m.G0 = v12 + int32(16)
	return v237
L8:
	;
	if v44 == int32(0) {
		goto L40
	} else {
		goto L41
	}
L9:
	;
	if v183 != 0 {
		goto L8
	} else {
		goto L36
	}
L10:
	;
	goto L9
L11:
	;
	v61 = int32(1)
	v62 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v45))))
	if v46 != v61 {
		goto L14
	} else {
		goto L15
	}
L12:
	;
	v183 = int32(1)
	goto L10
L13:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v48))) = v164
	goto L12
L14:
	;
	if v62&int32(255) == int32(45) {
		goto L19
	} else {
		goto L20
	}
L15:
	;
	v66 = v62 + int32(-48)
	if base.Ui32(int32(9)) < base.Ui32(v66&int32(255)) {
		v183 = v36
		goto L10
	} else {
		goto L16
	}
L16:
	;
	if v48 == int32(0) {
		goto L12
	} else {
		goto L17
	}
L17:
	;
	v164 = base.I64_extend_i32_u(v66) & int64(255)
	goto L13
L18:
	;
	if base.Ui32(int32(8)) < base.Ui32((v85+int32(-49))&int32(255)) {
		v183 = v36
		goto L10
	} else {
		goto L21
	}
L19:
	;
	v82 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v45)+1)))
	v84 = int32(2)
	v85 = v82
	v86 = v45 + int32(1)
	goto L18
L20:
	;
	v84 = v61
	v85 = v62
	v86 = v45
	goto L18
L21:
	;
	v97 = base.I64_extend_i32_u(v85+int32(-48)) & int64(255)
	if base.Ui32(v46) <= base.Ui32(v84) {
		v140 = v97
		goto L22
	} else {
		goto L23
	}
L22:
	;
	if v62&int32(255) != int32(45) {
		goto L30
	} else {
		goto L31
	}
L23:
	;
	v103 = v84
	v105 = v97
	v107 = v86
	goto L24
L24:
	;
	v109 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v107)+1)))
	if base.Ui32((v109+int32(-58))&int32(255)) < base.Ui32(int32(246)) {
		v183 = v36
		goto L10
	} else {
		goto L26
	}
L25:
	;
	v140 = v130
	goto L22
L26:
	;
	if base.Ui64(int64(1844674407370955161)) < base.Ui64(v105) {
		v183 = v36
		goto L10
	} else {
		goto L27
	}
L27:
	;
	v119 = v105 * int64(10)
	v124 = base.I64_extend_i32_u(v109+int32(-48)) & int64(255)
	if base.Ui64(v124^int64(-1)) < base.Ui64(v119) {
		v183 = v36
		goto L10
	} else {
		goto L28
	}
L28:
	;
	v128 = int32(1)
	v130 = v119 + v124
	v132 = v103 + v128
	if v132 != v46 {
		v103 = v132
		v105 = v130
		v107 = v107 + v128
		goto L24
	} else {
		goto L29
	}
L29:
	;
	goto L25
L30:
	;
	if v140 < int64(0) {
		v183 = v36
		goto L10
	} else {
		goto L34
	}
L31:
	;
	if base.Ui64(int64(-9223372036854775807-1)) < base.Ui64(v140) {
		v183 = v36
		goto L10
	} else {
		goto L32
	}
L32:
	;
	if v48 == int32(0) {
		goto L12
	} else {
		goto L33
	}
L33:
	;
	v164 = int64(0) - v140
	goto L13
L34:
	;
	if v48 == int32(0) {
		goto L12
	} else {
		goto L35
	}
L35:
	;
	v164 = v140
	goto L13
L36:
	;
	F_addReplyError(m, l0, int32(_a_F_getBitOffsetFromArgument_0))
	mBase = m.M
	v194 = m.ExcPending
	if v194 != 0 {
		goto L37
	} else {
		goto L38
	}
L37:
	;
	return int32(0)
L38:
	;
	v237 = int32(-1)
	goto L7
L39:
	;
	if v203 < int64(0) {
		goto L43
	} else {
		goto L44
	}
L40:
	;
	v202 = *(*int64)(unsafe.Add(mBase, uint32(v12)+8))
	v203 = v202
	goto L39
L41:
	;
	v198 = *(*int64)(unsafe.Add(mBase, uint32(v12)+8))
	v200 = v198 * base.I64_extend_i32_u(l4)
	*(*int64)(unsafe.Add(mBase, uint32(v12)+8)) = v200
	v203 = v200
	goto L39
L42:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l2))) = v223
	v237 = int32(0)
	goto L7
L43:
	;
	F_addReplyError(m, l0, int32(_a_F_getBitOffsetFromArgument_0))
	mBase = m.M
	v233 = m.ExcPending
	if v233 != 0 {
		goto L37
	} else {
		goto L54
	}
L44:
	;
	v207 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	if v207 != int64(-1) {
		goto L46
	} else {
		goto L47
	}
L45:
	;
	v223 = *(*int64)(unsafe.Add(mBase, uint32(v12)+8))
	if v222 != 0 {
		goto L42
	} else {
		goto L52
	}
L46:
	;
	v211 = int32(1)
	v212 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+200)))
	if v212&v211 != 0 {
		v219 = v211
		goto L48
	} else {
		goto L49
	}
L47:
	;
	v222 = int32(1)
	goto L45
L48:
	;
	v222 = v219
	goto L45
L49:
	;
	v215 = *(*int32)(unsafe.Add(mBase, uint32(l0)+216))
	if v215 != 0 {
		goto L50
	} else {
		goto L51
	}
L50:
	;
	v217 = F_isImportSlotMigrationJob(m, v215)
	mBase = m.M
	v219 = v217
	goto L48
L51:
	;
	v222 = int32(0)
	goto L45
L52:
	;
	v227 = *(*int64)(unsafe.Add(mBase, _c_F_getBitOffsetFromArgument[0]))
	if v223>>(uint(int64(3))%64) < v227 {
		goto L42
	} else {
		goto L53
	}
L53:
	;
	goto L43
L54:
	;
	v237 = int32(-1)
	goto L7
}
