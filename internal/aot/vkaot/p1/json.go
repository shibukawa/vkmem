package p1

import (
	base "github.com/shibukawa/vkmem/internal/aot/vkaot/base"
	"unsafe"
)

func F_json_append_string(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
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
	var v23 int32
	_ = v23
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v46 int32
	_ = v46
	var v52 int32
	_ = v52
	var v56 int32
	_ = v56
	var v60 int32
	_ = v60
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
	v9 = m.G0
	v11 = v9 - int32(16)
	m.G0 = v11
	v15 = F_lua_tolstring(m, l0, l2, v11+int32(12))
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v17 = *(*int32)(unsafe.Add(mBase, uint32(v11)+12))
	if base.Ui32(int32(-3)) <= base.Ui32(v17) {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L4:
	;
	v21 = v17 + int32(2)
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	if base.Ui32(v21) <= base.Ui32(v22+(v23^int32(-1))) {
		v32 = v23
		goto L5
	} else {
		goto L6
	}
L5:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+8)) = v32 + int32(1)
	v36 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v38 = int32(34)
	*(*uint8)(unsafe.Add(mBase, uint32(v36+v32))) = uint8(v38)
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v11)+12))
	if v40 == int32(0) {
		goto L8
	} else {
		goto L9
	}
L6:
	;
	F_strbuf_resize(m, l1, v23+v21)
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		goto L1
	} else {
		goto L7
	}
L7:
	;
	v31 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v32 = v31
	goto L5
L8:
	;
	v96 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v97 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v99 = v97 + int32(1)
	if v96 != v99 {
		v106 = v99
		v107 = v97
		goto L20
	} else {
		goto L21
	}
L9:
	;
	v46 = int32(0)
	goto L10
L10:
	;
	v52 = m.G3
	v56 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15+v46))))
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v52+int32(_a_F_json_append_string_0)+v56<<(uint(int32(2))%32))))
	if v60 == int32(0) {
		goto L13
	} else {
		goto L14
	}
L11:
	;
	goto L8
L12:
	;
	v85 = v46 + int32(1)
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v11)+12))
	if base.Ui32(v85) < base.Ui32(v86) {
		v46 = v85
		goto L10
	} else {
		goto L19
	}
L13:
	;
	v65 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v66 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v68 = v66 + int32(1)
	if v65 != v68 {
		v75 = v66
		v76 = v68
		goto L16
	} else {
		goto L17
	}
L14:
	;
	F_strbuf_append_string(m, l1, v60)
	mBase = m.M
	v64 = m.ExcPending
	if v64 != 0 {
		goto L1
	} else {
		goto L15
	}
L15:
	;
	goto L12
L16:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+8)) = v76
	v78 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	*(*uint8)(unsafe.Add(mBase, uint32(v78+v75))) = uint8(v56)
	goto L12
L17:
	;
	F_strbuf_resize(m, l1, v65)
	mBase = m.M
	v71 = m.ExcPending
	if v71 != 0 {
		goto L1
	} else {
		goto L18
	}
L18:
	;
	v72 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v75 = v72
	v76 = v72 + int32(1)
	goto L16
L19:
	;
	goto L11
L20:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+8)) = v106
	v109 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v111 = int32(34)
	*(*uint8)(unsafe.Add(mBase, uint32(v109+v107))) = uint8(v111)
	m.G0 = v11 + int32(16)
	return
L21:
	;
	F_strbuf_resize(m, l1, v96)
	mBase = m.M
	v102 = m.ExcPending
	if v102 != 0 {
		goto L1
	} else {
		goto L22
	}
L22:
	;
	v103 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v106 = v103 + int32(1)
	v107 = v103
	goto L20
}
func F_json_cfg_decode_invalid_numbers(m *base.Module, l0 int32) int32 {
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v11 int32
	_ = v11
	v3 = F_json_arg_init(m, l0, int32(1))
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		F_json_enum_option(m, l0, v3+int32(1332), int32(0))
		v11 = m.ExcPending
		if v11 != 0 {
			return int32(0)
		} else {
			return int32(1)
		}
	}
}
func F_json_cfg_encode_max_depth(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v17 int32
	_ = v17
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v101 int32
	_ = v101
	v5 = m.G0
	v7 = v5 - int32(80)
	m.G0 = v7
	v10 = F_json_arg_init(m, l0, int32(1))
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return int32(0)
	} else {
		v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
		v22 = v17 + int32(0)
		v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		if base.Ui32(v22) < base.Ui32(v23) {
			v65 = m.G398
			if v22 != v65 {
				v68 = *(*int32)(unsafe.Add(mBase, uint32(v22)+8))
				v71 = v68
			} else {
				v71 = int32(-1)
			}
		} else {
			v71 = int32(-1)
		}
		if v71 != 0 {
			v74 = F_luaL_checkinteger(m, l0, int32(1))
			mBase = m.M
			v75 = m.ExcPending
			if v75 != 0 {
				return int32(0)
			} else {
				*(*int64)(unsafe.Add(mBase, uint32(v7))) = int64(9223372032559808513)
				v81 = m.G3
				v84 = F_snprintf(m, v7+int32(16), int32(64), v81+int32(_a_F_json_cfg_encode_max_depth_0), v7)
				mBase = m.M
				v85 = m.ExcPending
				if v85 != 0 {
					return int32(0)
				} else {
					if int32(0) < v74 {
						*(*int32)(unsafe.Add(mBase, uint32(v10)+1316)) = v74
						v94 = v74
						v96 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
						*(*int32)(unsafe.Add(mBase, uint32(v96)+8)) = int32(3)
						*(*float64)(unsafe.Add(mBase, uint32(v96))) = base.F64_convert_i32_s(v94)
						v101 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
						*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v101 + int32(16)
						m.G0 = v7 + int32(80)
						return int32(1)
					} else {
						v91 = F_luaL_argerror(m, l0, int32(1), v7+int32(16))
						mBase = m.M
						v92 = m.ExcPending
						if v92 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v10)+1316)) = v74
							v94 = v74
							v96 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
							*(*int32)(unsafe.Add(mBase, uint32(v96)+8)) = int32(3)
							*(*float64)(unsafe.Add(mBase, uint32(v96))) = base.F64_convert_i32_s(v94)
							v101 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
							*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v101 + int32(16)
							m.G0 = v7 + int32(80)
							return int32(1)
						}
					}
				}
			}
		} else {
			v72 = *(*int32)(unsafe.Add(mBase, uint32(v10)+1316))
			v94 = v72
			v96 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			*(*int32)(unsafe.Add(mBase, uint32(v96)+8)) = int32(3)
			*(*float64)(unsafe.Add(mBase, uint32(v96))) = base.F64_convert_i32_s(v94)
			v101 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v101 + int32(16)
			m.G0 = v7 + int32(80)
			return int32(1)
		}
	}
}
func F_json_cfg_encode_number_precision(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v17 int32
	_ = v17
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v103 int32
	_ = v103
	v5 = m.G0
	v7 = v5 - int32(80)
	m.G0 = v7
	v10 = F_json_arg_init(m, l0, int32(1))
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return int32(0)
	} else {
		v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
		v22 = v17 + int32(0)
		v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		if base.Ui32(v22) < base.Ui32(v23) {
			v65 = m.G398
			if v22 != v65 {
				v68 = *(*int32)(unsafe.Add(mBase, uint32(v22)+8))
				v71 = v68
			} else {
				v71 = int32(-1)
			}
		} else {
			v71 = int32(-1)
		}
		if v71 != 0 {
			v74 = F_luaL_checkinteger(m, l0, int32(1))
			mBase = m.M
			v75 = m.ExcPending
			if v75 != 0 {
				return int32(0)
			} else {
				*(*int64)(unsafe.Add(mBase, uint32(v7))) = int64(60129542145)
				v81 = m.G3
				v84 = F_snprintf(m, v7+int32(16), int32(64), v81+int32(_a_F_json_cfg_encode_number_precision_0), v7)
				mBase = m.M
				v85 = m.ExcPending
				if v85 != 0 {
					return int32(0)
				} else {
					v86 = int32(-15)
					if base.Ui32(v86) < base.Ui32(v74+v86) {
						*(*int32)(unsafe.Add(mBase, uint32(v10)+1324)) = v74
						v96 = v74
						v98 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
						*(*int32)(unsafe.Add(mBase, uint32(v98)+8)) = int32(3)
						*(*float64)(unsafe.Add(mBase, uint32(v98))) = base.F64_convert_i32_s(v96)
						v103 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
						*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v103 + int32(16)
						m.G0 = v7 + int32(80)
						return int32(1)
					} else {
						v93 = F_luaL_argerror(m, l0, int32(1), v7+int32(16))
						mBase = m.M
						v94 = m.ExcPending
						if v94 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v10)+1324)) = v74
							v96 = v74
							v98 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
							*(*int32)(unsafe.Add(mBase, uint32(v98)+8)) = int32(3)
							*(*float64)(unsafe.Add(mBase, uint32(v98))) = base.F64_convert_i32_s(v96)
							v103 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
							*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v103 + int32(16)
							m.G0 = v7 + int32(80)
							return int32(1)
						}
					}
				}
			}
		} else {
			v72 = *(*int32)(unsafe.Add(mBase, uint32(v10)+1324))
			v96 = v72
			v98 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			*(*int32)(unsafe.Add(mBase, uint32(v98)+8)) = int32(3)
			*(*float64)(unsafe.Add(mBase, uint32(v98))) = base.F64_convert_i32_s(v96)
			v103 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v103 + int32(16)
			m.G0 = v7 + int32(80)
			return int32(1)
		}
	}
}
func F_json_cfg_encode_sparse_array(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v18 int32
	_ = v18
	var v22 int32
	_ = v22
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v86 int32
	_ = v86
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v108 int32
	_ = v108
	var v115 int32
	_ = v115
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v163 int32
	_ = v163
	var v166 int32
	_ = v166
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v179 int32
	_ = v179
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v192 int32
	_ = v192
	var v194 int32
	_ = v194
	var v199 int32
	_ = v199
	v5 = m.G0
	v7 = v5 - int32(96)
	m.G0 = v7
	v10 = F_json_arg_init(m, l0, int32(3))
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return int32(0)
	} else {
		F_json_enum_option(m, l0, v10+int32(1304), int32(0))
		mBase = m.M
		v18 = m.ExcPending
		if v18 != 0 {
			return int32(0)
		} else {
			v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
			v27 = v22 + int32(16)
			v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			if base.Ui32(v27) < base.Ui32(v28) {
				v70 = m.G398
				if v27 != v70 {
					v73 = *(*int32)(unsafe.Add(mBase, uint32(v27)+8))
					v76 = v73
				} else {
					v76 = int32(-1)
				}
			} else {
				v76 = int32(-1)
			}
			if v76 != 0 {
				v79 = F_luaL_checkinteger(m, l0, int32(2))
				mBase = m.M
				v80 = m.ExcPending
				if v80 != 0 {
					return int32(0)
				} else {
					*(*int64)(unsafe.Add(mBase, uint32(v7)+16)) = int64(9223372032559808512)
					v86 = m.G3
					v91 = F_snprintf(m, v7+int32(32), int32(64), v86+int32(_a_F_json_cfg_encode_sparse_array_0), v7+int32(16))
					mBase = m.M
					v92 = m.ExcPending
					if v92 != 0 {
						return int32(0)
					} else {
						if int32(-1) < v79 {
							*(*int32)(unsafe.Add(mBase, uint32(v10)+1308)) = v79
							v101 = v79
							v103 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
							*(*int32)(unsafe.Add(mBase, uint32(v103)+8)) = int32(3)
							*(*float64)(unsafe.Add(mBase, uint32(v103))) = base.F64_convert_i32_s(v101)
							v108 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
							*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v108 + int32(16)
							v115 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
							v120 = v115 + int32(32)
							v121 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
							if base.Ui32(v120) < base.Ui32(v121) {
								v163 = m.G398
								if v120 != v163 {
									v166 = *(*int32)(unsafe.Add(mBase, uint32(v120)+8))
									v169 = v166
								} else {
									v169 = int32(-1)
								}
							} else {
								v169 = int32(-1)
							}
							if v169 != 0 {
								v172 = F_luaL_checkinteger(m, l0, int32(3))
								mBase = m.M
								v173 = m.ExcPending
								if v173 != 0 {
									return int32(0)
								} else {
									*(*int64)(unsafe.Add(mBase, uint32(v7))) = int64(9223372032559808512)
									v179 = m.G3
									v182 = F_snprintf(m, v7+int32(32), int32(64), v179+int32(_a_F_json_cfg_encode_sparse_array_0), v7)
									mBase = m.M
									v183 = m.ExcPending
									if v183 != 0 {
										return int32(0)
									} else {
										if int32(-1) < v172 {
											*(*int32)(unsafe.Add(mBase, uint32(v10)+1312)) = v172
											v192 = v172
											v194 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
											*(*int32)(unsafe.Add(mBase, uint32(v194)+8)) = int32(3)
											*(*float64)(unsafe.Add(mBase, uint32(v194))) = base.F64_convert_i32_s(v192)
											v199 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
											*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v199 + int32(16)
											m.G0 = v7 + int32(96)
											return int32(3)
										} else {
											v189 = F_luaL_argerror(m, l0, int32(1), v7+int32(32))
											mBase = m.M
											v190 = m.ExcPending
											if v190 != 0 {
												return int32(0)
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(v10)+1312)) = v172
												v192 = v172
												v194 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
												*(*int32)(unsafe.Add(mBase, uint32(v194)+8)) = int32(3)
												*(*float64)(unsafe.Add(mBase, uint32(v194))) = base.F64_convert_i32_s(v192)
												v199 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
												*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v199 + int32(16)
												m.G0 = v7 + int32(96)
												return int32(3)
											}
										}
									}
								}
							} else {
								v170 = *(*int32)(unsafe.Add(mBase, uint32(v10)+1312))
								v192 = v170
								v194 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
								*(*int32)(unsafe.Add(mBase, uint32(v194)+8)) = int32(3)
								*(*float64)(unsafe.Add(mBase, uint32(v194))) = base.F64_convert_i32_s(v192)
								v199 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
								*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v199 + int32(16)
								m.G0 = v7 + int32(96)
								return int32(3)
							}
						} else {
							v98 = F_luaL_argerror(m, l0, int32(1), v7+int32(32))
							mBase = m.M
							v99 = m.ExcPending
							if v99 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v10)+1308)) = v79
								v101 = v79
								v103 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
								*(*int32)(unsafe.Add(mBase, uint32(v103)+8)) = int32(3)
								*(*float64)(unsafe.Add(mBase, uint32(v103))) = base.F64_convert_i32_s(v101)
								v108 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
								*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v108 + int32(16)
								v115 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
								v120 = v115 + int32(32)
								v121 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
								if base.Ui32(v120) < base.Ui32(v121) {
									v163 = m.G398
									if v120 != v163 {
										v166 = *(*int32)(unsafe.Add(mBase, uint32(v120)+8))
										v169 = v166
									} else {
										v169 = int32(-1)
									}
								} else {
									v169 = int32(-1)
								}
								if v169 != 0 {
									v172 = F_luaL_checkinteger(m, l0, int32(3))
									mBase = m.M
									v173 = m.ExcPending
									if v173 != 0 {
										return int32(0)
									} else {
										*(*int64)(unsafe.Add(mBase, uint32(v7))) = int64(9223372032559808512)
										v179 = m.G3
										v182 = F_snprintf(m, v7+int32(32), int32(64), v179+int32(_a_F_json_cfg_encode_sparse_array_0), v7)
										mBase = m.M
										v183 = m.ExcPending
										if v183 != 0 {
											return int32(0)
										} else {
											if int32(-1) < v172 {
												*(*int32)(unsafe.Add(mBase, uint32(v10)+1312)) = v172
												v192 = v172
												v194 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
												*(*int32)(unsafe.Add(mBase, uint32(v194)+8)) = int32(3)
												*(*float64)(unsafe.Add(mBase, uint32(v194))) = base.F64_convert_i32_s(v192)
												v199 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
												*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v199 + int32(16)
												m.G0 = v7 + int32(96)
												return int32(3)
											} else {
												v189 = F_luaL_argerror(m, l0, int32(1), v7+int32(32))
												mBase = m.M
												v190 = m.ExcPending
												if v190 != 0 {
													return int32(0)
												} else {
													*(*int32)(unsafe.Add(mBase, uint32(v10)+1312)) = v172
													v192 = v172
													v194 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
													*(*int32)(unsafe.Add(mBase, uint32(v194)+8)) = int32(3)
													*(*float64)(unsafe.Add(mBase, uint32(v194))) = base.F64_convert_i32_s(v192)
													v199 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
													*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v199 + int32(16)
													m.G0 = v7 + int32(96)
													return int32(3)
												}
											}
										}
									}
								} else {
									v170 = *(*int32)(unsafe.Add(mBase, uint32(v10)+1312))
									v192 = v170
									v194 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
									*(*int32)(unsafe.Add(mBase, uint32(v194)+8)) = int32(3)
									*(*float64)(unsafe.Add(mBase, uint32(v194))) = base.F64_convert_i32_s(v192)
									v199 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
									*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v199 + int32(16)
									m.G0 = v7 + int32(96)
									return int32(3)
								}
							}
						}
					}
				}
			} else {
				v77 = *(*int32)(unsafe.Add(mBase, uint32(v10)+1308))
				v101 = v77
				v103 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				*(*int32)(unsafe.Add(mBase, uint32(v103)+8)) = int32(3)
				*(*float64)(unsafe.Add(mBase, uint32(v103))) = base.F64_convert_i32_s(v101)
				v108 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v108 + int32(16)
				v115 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
				v120 = v115 + int32(32)
				v121 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				if base.Ui32(v120) < base.Ui32(v121) {
					v163 = m.G398
					if v120 != v163 {
						v166 = *(*int32)(unsafe.Add(mBase, uint32(v120)+8))
						v169 = v166
					} else {
						v169 = int32(-1)
					}
				} else {
					v169 = int32(-1)
				}
				if v169 != 0 {
					v172 = F_luaL_checkinteger(m, l0, int32(3))
					mBase = m.M
					v173 = m.ExcPending
					if v173 != 0 {
						return int32(0)
					} else {
						*(*int64)(unsafe.Add(mBase, uint32(v7))) = int64(9223372032559808512)
						v179 = m.G3
						v182 = F_snprintf(m, v7+int32(32), int32(64), v179+int32(_a_F_json_cfg_encode_sparse_array_0), v7)
						mBase = m.M
						v183 = m.ExcPending
						if v183 != 0 {
							return int32(0)
						} else {
							if int32(-1) < v172 {
								*(*int32)(unsafe.Add(mBase, uint32(v10)+1312)) = v172
								v192 = v172
								v194 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
								*(*int32)(unsafe.Add(mBase, uint32(v194)+8)) = int32(3)
								*(*float64)(unsafe.Add(mBase, uint32(v194))) = base.F64_convert_i32_s(v192)
								v199 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
								*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v199 + int32(16)
								m.G0 = v7 + int32(96)
								return int32(3)
							} else {
								v189 = F_luaL_argerror(m, l0, int32(1), v7+int32(32))
								mBase = m.M
								v190 = m.ExcPending
								if v190 != 0 {
									return int32(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v10)+1312)) = v172
									v192 = v172
									v194 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
									*(*int32)(unsafe.Add(mBase, uint32(v194)+8)) = int32(3)
									*(*float64)(unsafe.Add(mBase, uint32(v194))) = base.F64_convert_i32_s(v192)
									v199 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
									*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v199 + int32(16)
									m.G0 = v7 + int32(96)
									return int32(3)
								}
							}
						}
					}
				} else {
					v170 = *(*int32)(unsafe.Add(mBase, uint32(v10)+1312))
					v192 = v170
					v194 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
					*(*int32)(unsafe.Add(mBase, uint32(v194)+8)) = int32(3)
					*(*float64)(unsafe.Add(mBase, uint32(v194))) = base.F64_convert_i32_s(v192)
					v199 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
					*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v199 + int32(16)
					m.G0 = v7 + int32(96)
					return int32(3)
				}
			}
		}
	}
}
func F_json_is_invalid_number(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v18 int32
	_ = v18
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v54 int32
	_ = v54
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v112 int32
	_ = v112
	var v116 int32
	_ = v116
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v130 int32
	_ = v130
	var v132 int32
	_ = v132
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v153 int32
	_ = v153
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v6 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5))))
	switch v6 + int32(-43) {
	case 0:
		v153 = int32(1)
		goto L1
	default:
		v12 = v5
		v13 = v6
		goto L2
	case 2:
		goto L3
	}
L1:
	;
	return v153
L2:
	;
	if v13&int32(255) != int32(48) {
		goto L4
	} else {
		goto L5
	}
L3:
	;
	v9 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5)+1)))
	v12 = v5 + int32(1)
	v13 = v9
	goto L2
L4:
	;
	if base.I32_extend8_s(v13) < int32(58) {
		v153 = int32(0)
		goto L1
	} else {
		goto L6
	}
L5:
	;
	v18 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+1)))
	return base.B2i32(v18&int32(223) == int32(88)) | base.B2i32(base.Ui32((v18+int32(-48))&int32(255)) < base.Ui32(int32(10)))
L6:
	;
	v36 = m.G3
	v38 = v36 + int32(_a_F_json_is_invalid_number_0)
	goto L8
L7:
	;
	if v82-v84 == int32(0) {
		v153 = int32(1)
		goto L1
	} else {
		goto L22
	}
L8:
	;
	v43 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12))))
	if v43 != 0 {
		goto L11
	} else {
		goto L12
	}
L10:
	;
	v82 = F_tolower(m, v77)
	mBase = m.M
	v83 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v78))))
	v84 = F_tolower(m, v83)
	mBase = m.M
	goto L7
L11:
	;
	v45 = v12
	v46 = v38
	v47 = int32(3)
	v48 = v43
	goto L14
L12:
	;
	v77 = int32(0)
	v78 = v38
	goto L10
L13:
	;
	v77 = v74 & int32(255)
	v78 = v72
	goto L10
L14:
	;
	v50 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v46))))
	if v50 == int32(0) {
		v72 = v46
		v74 = v48
		goto L13
	} else {
		goto L16
	}
L15:
	;
	v72 = v66
	v74 = int32(0)
	goto L13
L16:
	;
	v54 = v47 + int32(-1)
	if v54 == int32(0) {
		v72 = v46
		v74 = v48
		goto L13
	} else {
		goto L17
	}
L17:
	;
	v58 = v48 & int32(255)
	if v58 == v50 {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	v65 = int32(1)
	v66 = v46 + v65
	v67 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v45)+1)))
	if v67 != 0 {
		v45 = v45 + v65
		v46 = v66
		v47 = v54
		v48 = v67
		goto L14
	} else {
		goto L21
	}
L19:
	;
	v60 = F_tolower(m, v58)
	mBase = m.M
	v61 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v46))))
	v62 = F_tolower(m, v61)
	mBase = m.M
	if v60 == v62 {
		goto L18
	} else {
		goto L20
	}
L20:
	;
	v64 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v45))))
	v72 = v46
	v74 = v64
	goto L13
L21:
	;
	goto L15
L22:
	;
	v94 = m.G3
	v96 = v94 + int32(_a_F_json_is_invalid_number_1)
	goto L24
L23:
	;
	v153 = base.B2i32(v140-v142 == int32(0))
	goto L1
L24:
	;
	v101 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12))))
	if v101 != 0 {
		goto L27
	} else {
		goto L28
	}
L26:
	;
	v140 = F_tolower(m, v135)
	mBase = m.M
	v141 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v136))))
	v142 = F_tolower(m, v141)
	mBase = m.M
	goto L23
L27:
	;
	v103 = v12
	v104 = v96
	v105 = int32(3)
	v106 = v101
	goto L30
L28:
	;
	v135 = int32(0)
	v136 = v96
	goto L26
L29:
	;
	v135 = v132 & int32(255)
	v136 = v130
	goto L26
L30:
	;
	v108 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v104))))
	if v108 == int32(0) {
		v130 = v104
		v132 = v106
		goto L29
	} else {
		goto L32
	}
L31:
	;
	v130 = v124
	v132 = int32(0)
	goto L29
L32:
	;
	v112 = v105 + int32(-1)
	if v112 == int32(0) {
		v130 = v104
		v132 = v106
		goto L29
	} else {
		goto L33
	}
L33:
	;
	v116 = v106 & int32(255)
	if v116 == v108 {
		goto L34
	} else {
		goto L35
	}
L34:
	;
	v123 = int32(1)
	v124 = v104 + v123
	v125 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v103)+1)))
	if v125 != 0 {
		v103 = v103 + v123
		v104 = v124
		v105 = v112
		v106 = v125
		goto L30
	} else {
		goto L37
	}
L35:
	;
	v118 = F_tolower(m, v116)
	mBase = m.M
	v119 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v104))))
	v120 = F_tolower(m, v119)
	mBase = m.M
	if v118 == v120 {
		goto L34
	} else {
		goto L36
	}
L36:
	;
	v122 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v103))))
	v130 = v104
	v132 = v122
	goto L29
L37:
	;
	goto L31
}
func F_json_next_token(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v25 int32
	_ = v25
	var v33 int32
	_ = v33
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v50 int32
	_ = v50
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v76 int32
	_ = v76
	var v85 int32
	_ = v85
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v100 int32
	_ = v100
	var v110 int32
	_ = v110
	var v117 int32
	_ = v117
	var v120 int32
	_ = v120
	var v124 int32
	_ = v124
	var v126 int32
	_ = v126
	var v131 int32
	_ = v131
	var v134 int32
	_ = v134
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v149 int32
	_ = v149
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v167 int32
	_ = v167
	var v170 int32
	_ = v170
	var v172 int32
	_ = v172
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v181 int32
	_ = v181
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v188 int32
	_ = v188
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v201 int32
	_ = v201
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v223 int32
	_ = v223
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
	var v243 int32
	_ = v243
	var v252 int32
	_ = v252
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v274 int32
	_ = v274
	var v283 int32
	_ = v283
	var v286 int32
	_ = v286
	var v289 int32
	_ = v289
	var v290 int32
	_ = v290
	var v300 int32
	_ = v300
	var v319 int32
	_ = v319
	var v321 int32
	_ = v321
	var v325 int32
	_ = v325
	var v331 int32
	_ = v331
	var v333 int32
	_ = v333
	var v334 int32
	_ = v334
	var v339 int32
	_ = v339
	var v341 int32
	_ = v341
	var v347 int32
	_ = v347
	var v350 int32
	_ = v350
	var v351 int32
	_ = v351
	var v352 int32
	_ = v352
	var v361 int32
	_ = v361
	var v370 int32
	_ = v370
	var v371 int32
	_ = v371
	var v372 int32
	_ = v372
	var v373 int32
	_ = v373
	var v383 int32
	_ = v383
	var v392 int32
	_ = v392
	var v393 int32
	_ = v393
	var v394 int32
	_ = v394
	var v403 int32
	_ = v403
	var v412 int32
	_ = v412
	var v413 int32
	_ = v413
	var v414 int32
	_ = v414
	var v434 int32
	_ = v434
	var v451 int32
	_ = v451
	var v452 int32
	_ = v452
	var v458 int32
	_ = v458
	var v460 int32
	_ = v460
	var v461 int32
	_ = v461
	var v466 int32
	_ = v466
	var v473 int32
	_ = v473
	var v480 int32
	_ = v480
	var v485 int32
	_ = v485
	var v486 int32
	_ = v486
	var v492 int32
	_ = v492
	var v493 int32
	_ = v493
	var v494 int32
	_ = v494
	var v498 int32
	_ = v498
	var v500 int32
	_ = v500
	var v503 int32
	_ = v503
	var v504 int32
	_ = v504
	var v506 int32
	_ = v506
	var v521 int32
	_ = v521
	var v522 int32
	_ = v522
	var v527 int32
	_ = v527
	var v528 int32
	_ = v528
	var v529 int32
	_ = v529
	var v534 int32
	_ = v534
	var v548 int32
	_ = v548
	var v550 int32
	_ = v550
	var v555 int32
	_ = v555
	var v557 int32
	_ = v557
	var v558 int32
	_ = v558
	var v559 int32
	_ = v559
	var v560 int32
	_ = v560
	var v562 int32
	_ = v562
	var v566 int32
	_ = v566
	var v570 int32
	_ = v570
	var v572 int32
	_ = v572
	var v573 int32
	_ = v573
	var v574 int32
	_ = v574
	var v576 int32
	_ = v576
	var v577 int32
	_ = v577
	var v578 int32
	_ = v578
	var v579 int32
	_ = v579
	var v584 int32
	_ = v584
	var v586 int32
	_ = v586
	var v589 int32
	_ = v589
	var v590 int32
	_ = v590
	var v594 int32
	_ = v594
	var v595 int32
	_ = v595
	var v596 int32
	_ = v596
	var v606 int32
	_ = v606
	var v608 int32
	_ = v608
	var v613 int32
	_ = v613
	var v615 int32
	_ = v615
	var v616 int32
	_ = v616
	var v617 int32
	_ = v617
	var v618 int32
	_ = v618
	var v620 int32
	_ = v620
	var v624 int32
	_ = v624
	var v628 int32
	_ = v628
	var v630 int32
	_ = v630
	var v631 int32
	_ = v631
	var v632 int32
	_ = v632
	var v634 int32
	_ = v634
	var v635 int32
	_ = v635
	var v636 int32
	_ = v636
	var v637 int32
	_ = v637
	var v642 int32
	_ = v642
	var v644 int32
	_ = v644
	var v647 int32
	_ = v647
	var v648 int32
	_ = v648
	var v652 int32
	_ = v652
	var v653 int32
	_ = v653
	var v654 int32
	_ = v654
	var v667 int32
	_ = v667
	var v671 int32
	_ = v671
	var v673 int32
	_ = v673
	var v678 int32
	_ = v678
	var v680 int32
	_ = v680
	var v681 int32
	_ = v681
	var v682 int32
	_ = v682
	var v683 int32
	_ = v683
	var v687 int32
	_ = v687
	var v692 int32
	_ = v692
	var v695 int32
	_ = v695
	var v696 int32
	_ = v696
	var v697 int32
	_ = v697
	var v702 int32
	_ = v702
	var v704 int32
	_ = v704
	var v707 int32
	_ = v707
	var v708 int32
	_ = v708
	var v712 int32
	_ = v712
	var v724 int32
	_ = v724
	var v728 int32
	_ = v728
	var v730 int32
	_ = v730
	var v735 int32
	_ = v735
	var v737 int32
	_ = v737
	var v738 int32
	_ = v738
	var v739 int32
	_ = v739
	var v740 int32
	_ = v740
	var v744 int32
	_ = v744
	var v749 int32
	_ = v749
	var v752 int32
	_ = v752
	var v753 int32
	_ = v753
	var v754 int32
	_ = v754
	var v759 int32
	_ = v759
	var v761 int32
	_ = v761
	var v764 int32
	_ = v764
	var v765 int32
	_ = v765
	var v769 int32
	_ = v769
	var v781 int32
	_ = v781
	var v785 int32
	_ = v785
	var v787 int32
	_ = v787
	var v792 int32
	_ = v792
	var v794 int32
	_ = v794
	var v795 int32
	_ = v795
	var v796 int32
	_ = v796
	var v797 int32
	_ = v797
	var v801 int32
	_ = v801
	var v806 int32
	_ = v806
	var v809 int32
	_ = v809
	var v810 int32
	_ = v810
	var v811 int32
	_ = v811
	var v816 int32
	_ = v816
	var v818 int32
	_ = v818
	var v821 int32
	_ = v821
	var v822 int32
	_ = v822
	var v826 int32
	_ = v826
	var v839 int32
	_ = v839
	var v845 int32
	_ = v845
	var v846 int32
	_ = v846
	var v849 int32
	_ = v849
	var v852 int32
	_ = v852
	var v853 int32
	_ = v853
	var v858 int32
	_ = v858
	var v875 int32
	_ = v875
	var v879 int32
	_ = v879
	var v882 int32
	_ = v882
	var v886 int32
	_ = v886
	var v890 int32
	_ = v890
	var v895 int32
	_ = v895
	var v899 int32
	_ = v899
	var v903 int32
	_ = v903
	var v913 float64
	_ = v913
	var v914 int32
	_ = v914
	var v916 int32
	_ = v916
	var v917 int32
	_ = v917
	var v921 int32
	_ = v921
	var v922 int32
	_ = v922
	var v931 int32
	_ = v931
	var v942 int32
	_ = v942
	var v961 int32
	_ = v961
	var v964 int32
	_ = v964
	v15 = m.G0
	v17 = v15 - int32(32)
	m.G0 = v17
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v21 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20))))
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v19+v21<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v25
	if v25 == int32(11) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	v68 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v69 = v61 - v68
	*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v69
	switch v60 + int32(-10) {
	case 0:
		goto L7
	default:
		goto L9
	case 2:
		goto L10
	case 3:
		goto L8
	}
L2:
	;
	v33 = v20
	goto L4
L3:
	;
	v59 = v21
	v60 = v25
	v61 = v20
	goto L1
L4:
	;
	v44 = v33 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v44
	v46 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v33)+1)))
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v19+v46<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v50
	if v50 == int32(11) {
		v33 = v44
		goto L4
	} else {
		goto L6
	}
L5:
	;
	v59 = v46
	v60 = v50
	v61 = v44
	goto L1
L6:
	;
	goto L5
L7:
	;
	m.G0 = v17 + int32(32)
	return
L8:
	;
	switch v59 + int32(-34) {
	case 0:
		goto L18
	default:
		goto L17
	case 11:
		goto L16
	}
L9:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v61 + int32(1)
	goto L7
L10:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v69
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = int32(12)
	v76 = m.G3
	*(*int32)(unsafe.Add(mBase, uint32(l1)+8)) = v76 + int32(_a_F_json_next_token_0)
	goto L7
L11:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = int32(12)
	v961 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v100 - v961
	v964 = m.G3
	*(*int32)(unsafe.Add(mBase, uint32(l1)+8)) = v964 + int32(_a_F_json_next_token_1)
	goto L7
L12:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v942))) = int32(-1)
	goto L11
L13:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v931))) = int32(-1)
	goto L11
L14:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = int32(5)
	v913 = F_fpconv_strtod(m, v61, v17+int32(16))
	mBase = m.M
	v914 = m.ExcPending
	if v914 != 0 {
		goto L176
	} else {
		goto L178
	}
L15:
	;
	v671 = m.G3
	v673 = v671 + int32(_a_F_json_next_token_2)
	goto L122
L16:
	;
	v521 = *(*int32)(unsafe.Add(mBase, uint32(v19)+1332))
	if v521 != 0 {
		goto L14
	} else {
		goto L79
	}
L17:
	;
	if base.Ui32(int32(9)) < base.Ui32((v59+int32(-48))&int32(255)) {
		goto L15
	} else {
		goto L78
	}
L18:
	;
	v85 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v61))))
	if v85 != int32(34) {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	v506 = m.G3
	m.Env.X__assert_fail(m, v506+int32(_a_F_json_next_token_3), v506+int32(_a_F_json_next_token_4), int32(889), v506+int32(_a_F_json_next_token_5))
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L20:
	;
	v91 = v61 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v91
	v93 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v93)+8)) = int32(0)
	v100 = v91
	goto L21
L21:
	;
	v110 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v100))))
	if v110 == int32(92) {
		goto L26
	} else {
		goto L27
	}
L23:
	;
	v192 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v100)+3)))
	if base.Ui32((v192+int32(-48))&int32(255)) < base.Ui32(int32(10)) {
		v210 = v192
		v211 = v138
		goto L37
	} else {
		goto L38
	}
L24:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v100 + int32(1)
	v177 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v178 = *(*int32)(unsafe.Add(mBase, uint32(v177)))
	v179 = *(*int32)(unsafe.Add(mBase, uint32(v177)+8))
	v181 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v178+v179))) = uint8(v181)
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = int32(4)
	v185 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v186 = *(*int32)(unsafe.Add(mBase, uint32(v185)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l1)+16)) = v186
	v188 = *(*int32)(unsafe.Add(mBase, uint32(v185)))
	*(*int32)(unsafe.Add(mBase, uint32(l1)+8)) = v188
	goto L7
L25:
	;
	v162 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v163 = *(*int32)(unsafe.Add(mBase, uint32(v162)+8))
	v164 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v162)+8)) = v163 + v164
	v167 = *(*int32)(unsafe.Add(mBase, uint32(v162)))
	*(*uint8)(unsafe.Add(mBase, uint32(v163+v167))) = uint8(v161)
	v170 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v172 = v170 + v164
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v172
	v100 = v172
	goto L21
L26:
	;
	v124 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v100)+1)))
	v126 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19+int32(1024)+v124))))
	if v126 == int32(117) {
		goto L31
	} else {
		goto L32
	}
L27:
	;
	if v110 == int32(34) {
		goto L24
	} else {
		goto L28
	}
L28:
	;
	if v110 != 0 {
		v161 = v110
		goto L25
	} else {
		goto L29
	}
L29:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = int32(12)
	v117 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v100 - v117
	v120 = m.G3
	*(*int32)(unsafe.Add(mBase, uint32(l1)+8)) = v120 + int32(_a_F_json_next_token_6)
	goto L7
L30:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v100 + int32(1)
	v161 = v126
	goto L25
L31:
	;
	v138 = int32(-48)
	v139 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v100)+2)))
	if base.Ui32(int32(10)) <= base.Ui32((v139+v138)&int32(255)) {
		goto L34
	} else {
		goto L35
	}
L32:
	;
	if v126 != 0 {
		goto L30
	} else {
		goto L33
	}
L33:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = int32(12)
	v131 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v100 - v131
	v134 = m.G3
	*(*int32)(unsafe.Add(mBase, uint32(l1)+8)) = v134 + int32(_a_F_json_next_token_7)
	goto L7
L34:
	;
	v149 = v139 | int32(32)
	if base.Ui32((v149+int32(-97))&int32(255)) < base.Ui32(int32(6)) {
		v190 = v149
		v191 = int32(-87)
		goto L23
	} else {
		goto L36
	}
L35:
	;
	v190 = v139
	v191 = int32(-48)
	goto L23
L36:
	;
	v942 = v17 + int32(16)
	goto L12
L37:
	;
	v212 = int32(-48)
	v213 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v100)+4)))
	if base.Ui32(int32(10)) <= base.Ui32((v213+v212)&int32(255)) {
		goto L41
	} else {
		goto L42
	}
L38:
	;
	v201 = v192 | int32(32)
	if base.Ui32((v201+int32(-97))&int32(255)) <= base.Ui32(int32(5)) {
		v210 = v201
		v211 = int32(-87)
		goto L37
	} else {
		goto L39
	}
L39:
	;
	v942 = v17 + int32(12)
	goto L12
L40:
	;
	v234 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v100)+5)))
	if base.Ui32((v234+int32(-48))&int32(255)) < base.Ui32(int32(10)) {
		v252 = v234
		v253 = v212
		goto L44
	} else {
		goto L45
	}
L41:
	;
	v223 = v213 | int32(32)
	if base.Ui32((v223+int32(-97))&int32(255)) <= base.Ui32(int32(5)) {
		v232 = v223
		v233 = int32(-87)
		goto L40
	} else {
		goto L43
	}
L42:
	;
	v232 = v213
	v233 = int32(-48)
	goto L40
L43:
	;
	v942 = v17 + int32(8)
	goto L12
L44:
	;
	v254 = int32(255)
	v274 = (v211+v210&v254)<<(uint(int32(8))%32) + (v191+v190&v254)<<(uint(int32(12))%32) + (v233+v232&v254)<<(uint(int32(4))%32) + v252&v254 + v253
	if v274 < int32(0) {
		goto L11
	} else {
		goto L47
	}
L45:
	;
	v243 = v234 | int32(32)
	if base.Ui32((v243+int32(-97))&int32(255)) <= base.Ui32(int32(5)) {
		v252 = v243
		v253 = int32(-87)
		goto L44
	} else {
		goto L46
	}
L46:
	;
	v942 = v17 + int32(4)
	goto L12
L47:
	;
	if v274&int32(63488) != int32(55296) {
		goto L51
	} else {
		goto L52
	}
L48:
	;
	v492 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v493 = *(*int32)(unsafe.Add(mBase, uint32(v492)))
	v494 = *(*int32)(unsafe.Add(mBase, uint32(v492)+8))
	if v485 == int32(0) {
		goto L76
	} else {
		goto L77
	}
L49:
	;
	v458 = int32(63)
	v460 = int32(128)
	v461 = v451&v458 | v460
	*(*uint8)(unsafe.Add(mBase, uint32(v17)+3)) = uint8(v461)
	v466 = int32(base.Ui32(v451)>>(uint(int32(18))%32)) | int32(240)
	*(*uint8)(unsafe.Add(mBase, uint32(v17))) = uint8(v466)
	v473 = int32(base.Ui32(v451)>>(uint(int32(6))%32))&v458 | v460
	*(*uint8)(unsafe.Add(mBase, uint32(v17)+2)) = uint8(v473)
	v480 = int32(base.Ui32(v451)>>(uint(int32(12))%32))&v458 | v460
	*(*uint8)(unsafe.Add(mBase, uint32(v17)+1)) = uint8(v480)
	v485 = int32(4)
	v486 = v452
	goto L48
L50:
	;
	v352 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v100)+9)))
	if base.Ui32((v352+int32(-48))&int32(255)) < base.Ui32(int32(10)) {
		v370 = v352
		v371 = v289
		goto L64
	} else {
		goto L65
	}
L51:
	;
	if base.Ui32(int32(127)) < base.Ui32(v274) {
		goto L59
	} else {
		goto L60
	}
L52:
	;
	if v274&int32(1024) != 0 {
		goto L11
	} else {
		goto L53
	}
L53:
	;
	v283 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v100)+6)))
	if v283 != int32(92) {
		goto L11
	} else {
		goto L54
	}
L54:
	;
	v286 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v100)+7)))
	if v286 != int32(117) {
		goto L11
	} else {
		goto L55
	}
L55:
	;
	v289 = int32(-48)
	v290 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v100)+8)))
	if base.Ui32(int32(10)) <= base.Ui32((v290+v289)&int32(255)) {
		goto L56
	} else {
		goto L57
	}
L56:
	;
	v300 = v290 | int32(32)
	if base.Ui32((v300+int32(-97))&int32(255)) < base.Ui32(int32(6)) {
		v350 = v300
		v351 = int32(-87)
		goto L50
	} else {
		goto L58
	}
L57:
	;
	v350 = v290
	v351 = int32(-48)
	goto L50
L58:
	;
	v931 = v17 + int32(16)
	goto L13
L59:
	;
	if base.Ui32(int32(2047)) < base.Ui32(v274) {
		goto L61
	} else {
		goto L62
	}
L60:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v17))) = uint8(v274)
	v485 = int32(1)
	v486 = int32(6)
	goto L48
L61:
	;
	if base.Ui32(int32(65535)) < base.Ui32(v274) {
		v451 = v274
		v452 = int32(6)
		goto L49
	} else {
		goto L63
	}
L62:
	;
	v319 = v274&int32(63) | int32(128)
	*(*uint8)(unsafe.Add(mBase, uint32(v17)+1)) = uint8(v319)
	v321 = int32(6)
	v325 = int32(base.Ui32(v274)>>(uint(v321)%32)) | int32(192)
	*(*uint8)(unsafe.Add(mBase, uint32(v17))) = uint8(v325)
	v485 = int32(2)
	v486 = v321
	goto L48
L63:
	;
	v331 = int32(63)
	v333 = int32(128)
	v334 = v274&v331 | v333
	*(*uint8)(unsafe.Add(mBase, uint32(v17)+2)) = uint8(v334)
	v339 = int32(base.Ui32(v274)>>(uint(int32(12))%32)) | int32(224)
	*(*uint8)(unsafe.Add(mBase, uint32(v17))) = uint8(v339)
	v341 = int32(6)
	v347 = int32(base.Ui32(v274)>>(uint(v341)%32))&v331 | v333
	*(*uint8)(unsafe.Add(mBase, uint32(v17)+1)) = uint8(v347)
	v485 = int32(3)
	v486 = v341
	goto L48
L64:
	;
	v372 = int32(-48)
	v373 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v100)+10)))
	if base.Ui32(int32(10)) <= base.Ui32((v373+v372)&int32(255)) {
		goto L68
	} else {
		goto L69
	}
L65:
	;
	v361 = v352 | int32(32)
	if base.Ui32((v361+int32(-97))&int32(255)) <= base.Ui32(int32(5)) {
		v370 = v361
		v371 = int32(-87)
		goto L64
	} else {
		goto L66
	}
L66:
	;
	v931 = v17 + int32(12)
	goto L13
L67:
	;
	v394 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v100)+11)))
	if base.Ui32((v394+int32(-48))&int32(255)) < base.Ui32(int32(10)) {
		v412 = v394
		v413 = v372
		goto L71
	} else {
		goto L72
	}
L68:
	;
	v383 = v373 | int32(32)
	if base.Ui32((v383+int32(-97))&int32(255)) <= base.Ui32(int32(5)) {
		v392 = v383
		v393 = int32(-87)
		goto L67
	} else {
		goto L70
	}
L69:
	;
	v392 = v373
	v393 = int32(-48)
	goto L67
L70:
	;
	v931 = v17 + int32(8)
	goto L13
L71:
	;
	v414 = int32(255)
	v434 = (v371+v370&v414)<<(uint(int32(8))%32) + (v351+v350&v414)<<(uint(int32(12))%32) + (v393+v392&v414)<<(uint(int32(4))%32) + v412&v414 + v413
	if v434&int32(-2147419136) != int32(56320) {
		goto L11
	} else {
		goto L74
	}
L72:
	;
	v403 = v394 | int32(32)
	if base.Ui32((v403+int32(-97))&int32(255)) <= base.Ui32(int32(5)) {
		v412 = v403
		v413 = int32(-87)
		goto L71
	} else {
		goto L73
	}
L73:
	;
	v931 = v17 + int32(4)
	goto L13
L74:
	;
	v451 = v274<<(uint(int32(10))%32)&int32(1047552) | v434&int32(1023) + int32(65536)
	v452 = int32(12)
	goto L49
L75:
	;
	v500 = *(*int32)(unsafe.Add(mBase, uint32(v492)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v492)+8)) = v500 + v485
	v503 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v504 = v503 + v486
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v504
	v100 = v504
	goto L21
L76:
	;
	goto L75
L77:
	;
	v498 = F__emscripten_memcpy_bulkmem(m, v493+v494, v17, v485)
	mBase = m.M
	goto L76
L78:
	;
	goto L16
L79:
	;
	v522 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v61))))
	switch v522 + int32(-43) {
	case 0:
		goto L80
	default:
		v528 = v61
		v529 = v522
		goto L81
	case 2:
		goto L82
	}
L80:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v69
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = int32(12)
	v667 = m.G3
	*(*int32)(unsafe.Add(mBase, uint32(l1)+8)) = v667 + int32(_a_F_json_next_token_8)
	goto L7
L81:
	;
	if v529&int32(255) != int32(48) {
		goto L83
	} else {
		goto L84
	}
L82:
	;
	v527 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v61)+1)))
	v528 = v61 + int32(1)
	v529 = v527
	goto L81
L83:
	;
	if base.I32_extend8_s(v529) < int32(58) {
		goto L14
	} else {
		goto L87
	}
L84:
	;
	v534 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v528)+1)))
	if v534&int32(223) == int32(88) {
		goto L80
	} else {
		goto L85
	}
L85:
	;
	if base.Ui32((v534+int32(-48))&int32(255)) < base.Ui32(int32(10)) {
		goto L80
	} else {
		goto L86
	}
L86:
	;
	goto L14
L87:
	;
	v548 = m.G3
	v550 = v548 + int32(_a_F_json_next_token_9)
	goto L89
L88:
	;
	if v594-v596 == int32(0) {
		goto L80
	} else {
		goto L103
	}
L89:
	;
	v555 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v528))))
	if v555 != 0 {
		goto L92
	} else {
		goto L93
	}
L91:
	;
	v594 = F_tolower(m, v589)
	mBase = m.M
	v595 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v590))))
	v596 = F_tolower(m, v595)
	mBase = m.M
	goto L88
L92:
	;
	v557 = v528
	v558 = v550
	v559 = int32(3)
	v560 = v555
	goto L95
L93:
	;
	v589 = int32(0)
	v590 = v550
	goto L91
L94:
	;
	v589 = v586 & int32(255)
	v590 = v584
	goto L91
L95:
	;
	v562 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v558))))
	if v562 == int32(0) {
		v584 = v558
		v586 = v560
		goto L94
	} else {
		goto L97
	}
L96:
	;
	v584 = v578
	v586 = int32(0)
	goto L94
L97:
	;
	v566 = v559 + int32(-1)
	if v566 == int32(0) {
		v584 = v558
		v586 = v560
		goto L94
	} else {
		goto L98
	}
L98:
	;
	v570 = v560 & int32(255)
	if v570 == v562 {
		goto L99
	} else {
		goto L100
	}
L99:
	;
	v577 = int32(1)
	v578 = v558 + v577
	v579 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v557)+1)))
	if v579 != 0 {
		v557 = v557 + v577
		v558 = v578
		v559 = v566
		v560 = v579
		goto L95
	} else {
		goto L102
	}
L100:
	;
	v572 = F_tolower(m, v570)
	mBase = m.M
	v573 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v558))))
	v574 = F_tolower(m, v573)
	mBase = m.M
	if v572 == v574 {
		goto L99
	} else {
		goto L101
	}
L101:
	;
	v576 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v557))))
	v584 = v558
	v586 = v576
	goto L94
L102:
	;
	goto L96
L103:
	;
	v606 = m.G3
	v608 = v606 + int32(_a_F_json_next_token_10)
	goto L105
L104:
	;
	if v652-v654 != 0 {
		goto L14
	} else {
		goto L119
	}
L105:
	;
	v613 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v528))))
	if v613 != 0 {
		goto L108
	} else {
		goto L109
	}
L107:
	;
	v652 = F_tolower(m, v647)
	mBase = m.M
	v653 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v648))))
	v654 = F_tolower(m, v653)
	mBase = m.M
	goto L104
L108:
	;
	v615 = v528
	v616 = v608
	v617 = int32(3)
	v618 = v613
	goto L111
L109:
	;
	v647 = int32(0)
	v648 = v608
	goto L107
L110:
	;
	v647 = v644 & int32(255)
	v648 = v642
	goto L107
L111:
	;
	v620 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v616))))
	if v620 == int32(0) {
		v642 = v616
		v644 = v618
		goto L110
	} else {
		goto L113
	}
L112:
	;
	v642 = v636
	v644 = int32(0)
	goto L110
L113:
	;
	v624 = v617 + int32(-1)
	if v624 == int32(0) {
		v642 = v616
		v644 = v618
		goto L110
	} else {
		goto L114
	}
L114:
	;
	v628 = v618 & int32(255)
	if v628 == v620 {
		goto L115
	} else {
		goto L116
	}
L115:
	;
	v635 = int32(1)
	v636 = v616 + v635
	v637 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v615)+1)))
	if v637 != 0 {
		v615 = v615 + v635
		v616 = v636
		v617 = v624
		v618 = v637
		goto L111
	} else {
		goto L118
	}
L116:
	;
	v630 = F_tolower(m, v628)
	mBase = m.M
	v631 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v616))))
	v632 = F_tolower(m, v631)
	mBase = m.M
	if v630 == v632 {
		goto L115
	} else {
		goto L117
	}
L117:
	;
	v634 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v615))))
	v642 = v616
	v644 = v634
	goto L110
L118:
	;
	goto L112
L119:
	;
	goto L80
L120:
	;
	v728 = m.G3
	v730 = v728 + int32(_a_F_json_next_token_11)
	goto L137
L121:
	;
	if v707-v712 != 0 {
		goto L120
	} else {
		goto L134
	}
L122:
	;
	v678 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v61))))
	if v678 != 0 {
		goto L125
	} else {
		goto L126
	}
L124:
	;
	v712 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v708))))
	goto L121
L125:
	;
	v680 = v61
	v681 = v673
	v682 = int32(4)
	v683 = v678
	goto L128
L126:
	;
	v707 = int32(0)
	v708 = v673
	goto L124
L127:
	;
	v707 = v704 & int32(255)
	v708 = v702
	goto L124
L128:
	;
	v687 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v681))))
	if v683&int32(255) != v687 {
		v702 = v681
		v704 = v683
		goto L127
	} else {
		goto L130
	}
L129:
	;
	v702 = v696
	v704 = int32(0)
	goto L127
L130:
	;
	if v687 == int32(0) {
		v702 = v681
		v704 = v683
		goto L127
	} else {
		goto L131
	}
L131:
	;
	v692 = v682 + int32(-1)
	if v692 == int32(0) {
		v702 = v681
		v704 = v683
		goto L127
	} else {
		goto L132
	}
L132:
	;
	v695 = int32(1)
	v696 = v681 + v695
	v697 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v680)+1)))
	if v697 != 0 {
		v680 = v680 + v695
		v681 = v696
		v682 = v692
		v683 = v697
		goto L128
	} else {
		goto L133
	}
L133:
	;
	goto L129
L134:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+8)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = int32(6)
	v724 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v724 + int32(4)
	goto L7
L135:
	;
	v785 = m.G3
	v787 = v785 + int32(_a_F_json_next_token_12)
	goto L152
L136:
	;
	if v764-v769 != 0 {
		goto L135
	} else {
		goto L149
	}
L137:
	;
	v735 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v61))))
	if v735 != 0 {
		goto L140
	} else {
		goto L141
	}
L139:
	;
	v769 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v765))))
	goto L136
L140:
	;
	v737 = v61
	v738 = v730
	v739 = int32(5)
	v740 = v735
	goto L143
L141:
	;
	v764 = int32(0)
	v765 = v730
	goto L139
L142:
	;
	v764 = v761 & int32(255)
	v765 = v759
	goto L139
L143:
	;
	v744 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v738))))
	if v740&int32(255) != v744 {
		v759 = v738
		v761 = v740
		goto L142
	} else {
		goto L145
	}
L144:
	;
	v759 = v753
	v761 = int32(0)
	goto L142
L145:
	;
	if v744 == int32(0) {
		v759 = v738
		v761 = v740
		goto L142
	} else {
		goto L146
	}
L146:
	;
	v749 = v739 + int32(-1)
	if v749 == int32(0) {
		v759 = v738
		v761 = v740
		goto L142
	} else {
		goto L147
	}
L147:
	;
	v752 = int32(1)
	v753 = v738 + v752
	v754 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v737)+1)))
	if v754 != 0 {
		v737 = v737 + v752
		v738 = v753
		v739 = v749
		v740 = v754
		goto L143
	} else {
		goto L148
	}
L148:
	;
	goto L144
L149:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+8)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = int32(6)
	v781 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v781 + int32(5)
	goto L7
L150:
	;
	v839 = *(*int32)(unsafe.Add(mBase, uint32(v19)+1332))
	if v839 == int32(0) {
		goto L165
	} else {
		goto L166
	}
L151:
	;
	if v821-v826 != 0 {
		goto L150
	} else {
		goto L164
	}
L152:
	;
	v792 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v61))))
	if v792 != 0 {
		goto L155
	} else {
		goto L156
	}
L154:
	;
	v826 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v822))))
	goto L151
L155:
	;
	v794 = v61
	v795 = v787
	v796 = int32(4)
	v797 = v792
	goto L158
L156:
	;
	v821 = int32(0)
	v822 = v787
	goto L154
L157:
	;
	v821 = v818 & int32(255)
	v822 = v816
	goto L154
L158:
	;
	v801 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v795))))
	if v797&int32(255) != v801 {
		v816 = v795
		v818 = v797
		goto L157
	} else {
		goto L160
	}
L159:
	;
	v816 = v810
	v818 = int32(0)
	goto L157
L160:
	;
	if v801 == int32(0) {
		v816 = v795
		v818 = v797
		goto L157
	} else {
		goto L161
	}
L161:
	;
	v806 = v796 + int32(-1)
	if v806 == int32(0) {
		v816 = v795
		v818 = v797
		goto L157
	} else {
		goto L162
	}
L162:
	;
	v809 = int32(1)
	v810 = v795 + v809
	v811 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v794)+1)))
	if v811 != 0 {
		v794 = v794 + v809
		v795 = v810
		v796 = v806
		v797 = v811
		goto L158
	} else {
		goto L163
	}
L163:
	;
	goto L159
L164:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = int32(7)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v61 + int32(4)
	goto L7
L165:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v69
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = int32(12)
	v903 = m.G3
	*(*int32)(unsafe.Add(mBase, uint32(l1)+8)) = v903 + int32(_a_F_json_next_token_0)
	goto L7
L166:
	;
	v845 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v846 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v845))))
	switch v846 + int32(-43) {
	case 0:
		v890 = int32(1)
		goto L168
	default:
		v852 = v845
		v853 = v846
		goto L169
	case 2:
		goto L170
	}
L167:
	;
	if v895 == int32(0) {
		goto L165
	} else {
		goto L175
	}
L168:
	;
	v895 = v890
	goto L167
L169:
	;
	if v853&int32(255) != int32(48) {
		goto L171
	} else {
		goto L172
	}
L170:
	;
	v849 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v845)+1)))
	v852 = v845 + int32(1)
	v853 = v849
	goto L169
L171:
	;
	if base.I32_extend8_s(v853) < int32(58) {
		v890 = int32(0)
		goto L168
	} else {
		goto L173
	}
L172:
	;
	v858 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v852)+1)))
	v895 = base.B2i32(v858&int32(223) == int32(88)) | base.B2i32(base.Ui32((v858+int32(-48))&int32(255)) < base.Ui32(int32(10)))
	goto L167
L173:
	;
	v875 = m.G3
	v879 = F_strncasecmp(m, v852, v875+int32(_a_F_json_next_token_9), int32(3))
	mBase = m.M
	if v879 == int32(0) {
		v890 = int32(1)
		goto L168
	} else {
		goto L174
	}
L174:
	;
	v882 = m.G3
	v886 = F_strncasecmp(m, v852, v882+int32(_a_F_json_next_token_10), int32(3))
	mBase = m.M
	v890 = base.B2i32(v886 == int32(0))
	goto L168
L175:
	;
	F_json_next_number_token(m, l0, l1)
	mBase = m.M
	v899 = m.ExcPending
	if v899 != 0 {
		goto L176
	} else {
		goto L177
	}
L176:
	;
	return
L177:
	;
	goto L7
L178:
	;
	*(*float64)(unsafe.Add(mBase, uint32(l1)+8)) = v913
	v916 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v917 = *(*int32)(unsafe.Add(mBase, uint32(v17)+16))
	if v916 != v917 {
		goto L179
	} else {
		goto L180
	}
L179:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v917
	goto L7
L180:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = int32(12)
	v921 = m.G3
	v922 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(l1)+8)) = v921 + int32(_a_F_json_next_token_8)
	*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v916 - v922
	goto L7
}
