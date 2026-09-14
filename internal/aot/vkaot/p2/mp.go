package p2

import (
	base "github.com/shibukawa/vkmem/internal/aot/vkaot/base"
	"unsafe"
)

func F_mp_decode_to_lua_array(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v14 int32
	_ = v14
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v41 int32
	_ = v41
	v4 = int32(0)
	F_lua_createtable(m, l0, v4, v4)
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v10 = m.G3
	F_luaL_checkstack(m, l0, int32(1), v10+int32(_a_F_mp_decode_to_lua_array_0))
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	if l2 == int32(0) {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	return
L5:
	;
	v20 = l2
	v21 = int32(1)
	goto L6
L6:
	;
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v24)+8)) = int32(3)
	*(*float64)(unsafe.Add(mBase, uint32(v24))) = base.F64_convert_i32_u(v21)
	v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v28 + int32(16)
	goto L8
L7:
	;
	goto L4
L8:
	;
	F_mp_decode_to_lua_type(m, l0, l1)
	mBase = m.M
	v33 = m.ExcPending
	if v33 != 0 {
		goto L1
	} else {
		goto L9
	}
L9:
	;
	v34 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	if v34 != 0 {
		goto L4
	} else {
		goto L10
	}
L10:
	;
	F_lua_settable(m, l0, int32(-3))
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		goto L1
	} else {
		goto L11
	}
L11:
	;
	v41 = v20 + int32(-1)
	if v41 != 0 {
		v20 = v41
		v21 = v21 + int32(1)
		goto L6
	} else {
		goto L12
	}
L12:
	;
	goto L7
}
func F_mp_decode_to_lua_hash(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	v4 = int32(0)
	F_lua_createtable(m, l0, v4, v4)
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	if l2 == int32(0) {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	return
L4:
	;
	v12 = l2
	goto L5
L5:
	;
	F_mp_decode_to_lua_type(m, l0, l1)
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		goto L1
	} else {
		goto L7
	}
L6:
	;
	goto L3
L7:
	;
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	if v15 != 0 {
		goto L3
	} else {
		goto L8
	}
L8:
	;
	F_mp_decode_to_lua_type(m, l0, l1)
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		goto L1
	} else {
		goto L9
	}
L9:
	;
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	if v18 != 0 {
		goto L3
	} else {
		goto L10
	}
L10:
	;
	F_lua_settable(m, l0, int32(-3))
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		goto L1
	} else {
		goto L11
	}
L11:
	;
	v23 = v12 + int32(-1)
	if v23 != 0 {
		v12 = v23
		goto L5
	} else {
		goto L12
	}
L12:
	;
	goto L6
}
func F_mp_encode_double(m *base.Module, l0 int32, l1 int32, l2 float64) {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v14 float32
	_ = v14
	var v17 int32
	_ = v17
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
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v74 int32
	_ = v74
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v100 int32
	_ = v100
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
	var v111 int32
	_ = v111
	var v112 int64
	_ = v112
	var v113 int64
	_ = v113
	var v115 int64
	_ = v115
	var v117 int64
	_ = v117
	var v120 int64
	_ = v120
	var v122 int64
	_ = v122
	var v124 int64
	_ = v124
	var v126 int64
	_ = v126
	var v149 int32
	_ = v149
	var v151 int32
	_ = v151
	var v155 int32
	_ = v155
	v10 = m.G0
	v12 = v10 - int32(16)
	m.G0 = v12
	v14 = base.F32_demote_f64(l2)
	if base.F64_ne(l2, base.F64_promote_f32(v14)) != 0 {
		v78 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
		if base.Ui32(v78) < base.Ui32(int32(9)) {
			v83 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
			if base.Ui32(int32(2147483638)) <= base.Ui32(v83) {
				F_abort(m)
				mBase = m.M
				base.Wasm_trap_unreachable()
				for {
				}
			} else {
				v86 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
				v88 = v12 + int32(12)
				v89 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
				if v88 == int32(0) {
				} else {
					v92 = *(*int32)(unsafe.Add(mBase, uint32(v89)+16))
					*(*int32)(unsafe.Add(mBase, uint32(v88))) = v92
				}
				v94 = *(*int32)(unsafe.Add(mBase, uint32(v89)+12))
				v95 = *(*int32)(unsafe.Add(mBase, uint32(v12)+12))
				v100 = v83<<(uint(int32(1))%32) + int32(18)
				v101 = m.T0[v94].(func(*base.Module, int32, int32, int32, int32) int32)(m, v95, v86, v83+v78, v100)
				mBase = m.M
				v102 = m.ExcPending
				if v102 != 0 {
					return
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(l1))) = v101
					v104 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
					*(*int32)(unsafe.Add(mBase, uint32(l1)+8)) = v100 - v104
					v107 = v104
					v108 = v101
					v111 = v108 + v107
					v112 = base.I64_reinterpret_f64(l2)
					v113 = int64(56)
					v115 = int64(65280)
					v117 = int64(40)
					v120 = int64(16711680)
					v122 = int64(24)
					v124 = int64(4278190080)
					v126 = int64(8)
					*(*int64)(unsafe.Add(mBase, uint32(v111)+1)) = v112<<(uint(v113)%64) | v112&v115<<(uint(v117)%64) | (v112&v120<<(uint(v122)%64) | v112&v124<<(uint(v126)%64)) | (int64(base.Ui64(v112)>>(uint(v126)%64))&v124 | int64(base.Ui64(v112)>>(uint(v122)%64))&v120 | (int64(base.Ui64(v112)>>(uint(v117)%64))&v115 | int64(base.Ui64(v112)>>(uint(v113)%64))))
					v149 = int32(203)
					*(*uint8)(unsafe.Add(mBase, uint32(v111))) = uint8(v149)
					v151 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
					*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v151 + int32(9)
					v155 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
					*(*int32)(unsafe.Add(mBase, uint32(l1)+8)) = v155 + int32(-9)
					m.G0 = v12 + int32(16)
					return
				}
			}
		} else {
			v81 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
			v82 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
			v107 = v81
			v108 = v82
			v111 = v108 + v107
			v112 = base.I64_reinterpret_f64(l2)
			v113 = int64(56)
			v115 = int64(65280)
			v117 = int64(40)
			v120 = int64(16711680)
			v122 = int64(24)
			v124 = int64(4278190080)
			v126 = int64(8)
			*(*int64)(unsafe.Add(mBase, uint32(v111)+1)) = v112<<(uint(v113)%64) | v112&v115<<(uint(v117)%64) | (v112&v120<<(uint(v122)%64) | v112&v124<<(uint(v126)%64)) | (int64(base.Ui64(v112)>>(uint(v126)%64))&v124 | int64(base.Ui64(v112)>>(uint(v122)%64))&v120 | (int64(base.Ui64(v112)>>(uint(v117)%64))&v115 | int64(base.Ui64(v112)>>(uint(v113)%64))))
			v149 = int32(203)
			*(*uint8)(unsafe.Add(mBase, uint32(v111))) = uint8(v149)
			v151 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
			*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v151 + int32(9)
			v155 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
			*(*int32)(unsafe.Add(mBase, uint32(l1)+8)) = v155 + int32(-9)
			m.G0 = v12 + int32(16)
			return
		}
	} else {
		v17 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
		if base.Ui32(v17) < base.Ui32(int32(5)) {
			v22 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
			if base.Ui32(int32(2147483642)) <= base.Ui32(v22) {
				F_abort(m)
				mBase = m.M
				base.Wasm_trap_unreachable()
				for {
				}
			} else {
				v25 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
				v27 = v12 + int32(8)
				v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
				if v27 == int32(0) {
				} else {
					v31 = *(*int32)(unsafe.Add(mBase, uint32(v28)+16))
					*(*int32)(unsafe.Add(mBase, uint32(v27))) = v31
				}
				v33 = *(*int32)(unsafe.Add(mBase, uint32(v28)+12))
				v34 = *(*int32)(unsafe.Add(mBase, uint32(v12)+8))
				v39 = v22<<(uint(int32(1))%32) + int32(10)
				v40 = m.T0[v33].(func(*base.Module, int32, int32, int32, int32) int32)(m, v34, v25, v22+v17, v39)
				mBase = m.M
				v41 = m.ExcPending
				if v41 != 0 {
					return
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(l1))) = v40
					v43 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
					*(*int32)(unsafe.Add(mBase, uint32(l1)+8)) = v39 - v43
					v46 = v43
					v47 = v40
					v50 = v47 + v46
					v51 = base.I32_reinterpret_f32(v14)
					v52 = int32(24)
					v54 = int32(65280)
					v56 = int32(8)
					*(*int32)(unsafe.Add(mBase, uint32(v50)+1)) = v51<<(uint(v52)%32) | v51&v54<<(uint(v56)%32) | (int32(base.Ui32(v51)>>(uint(v56)%32))&v54 | int32(base.Ui32(v51)>>(uint(v52)%32)))
					v68 = int32(202)
					*(*uint8)(unsafe.Add(mBase, uint32(v50))) = uint8(v68)
					v70 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
					*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v70 + int32(5)
					v74 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
					*(*int32)(unsafe.Add(mBase, uint32(l1)+8)) = v74 + int32(-5)
					m.G0 = v12 + int32(16)
					return
				}
			}
		} else {
			v20 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
			v21 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
			v46 = v20
			v47 = v21
			v50 = v47 + v46
			v51 = base.I32_reinterpret_f32(v14)
			v52 = int32(24)
			v54 = int32(65280)
			v56 = int32(8)
			*(*int32)(unsafe.Add(mBase, uint32(v50)+1)) = v51<<(uint(v52)%32) | v51&v54<<(uint(v56)%32) | (int32(base.Ui32(v51)>>(uint(v56)%32))&v54 | int32(base.Ui32(v51)>>(uint(v52)%32)))
			v68 = int32(202)
			*(*uint8)(unsafe.Add(mBase, uint32(v50))) = uint8(v68)
			v70 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
			*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v70 + int32(5)
			v74 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
			*(*int32)(unsafe.Add(mBase, uint32(l1)+8)) = v74 + int32(-5)
			m.G0 = v12 + int32(16)
			return
		}
	}
}
func F_mp_encode_lua_bool(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
	var v116 int32
	_ = v116
	v8 = m.G0
	v10 = v8 - int32(16)
	m.G0 = v10
	v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v31 = v28 + int32(-16)
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v31)+8))
	switch v67 {
	case 0:
		v72 = v67
		v74 = v72
	case 1:
		v68 = *(*int32)(unsafe.Add(mBase, uint32(v31)))
		v74 = base.B2i32(v68 != int32(0))
	default:
		v72 = int32(1)
		v74 = v72
	}
	v75 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	if v75 == int32(0) {
		v80 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
		if base.Ui32(int32(2147483646)) <= base.Ui32(v80) {
			F_abort(m)
			mBase = m.M
			base.Wasm_trap_unreachable()
			for {
			}
		} else {
			v83 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
			v85 = v10 + int32(12)
			v86 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
			if v85 == int32(0) {
			} else {
				v89 = *(*int32)(unsafe.Add(mBase, uint32(v86)+16))
				*(*int32)(unsafe.Add(mBase, uint32(v85))) = v89
			}
			v91 = *(*int32)(unsafe.Add(mBase, uint32(v86)+12))
			v92 = *(*int32)(unsafe.Add(mBase, uint32(v10)+12))
			v96 = v80<<(uint(int32(1))%32) + int32(2)
			v97 = m.T0[v91].(func(*base.Module, int32, int32, int32, int32) int32)(m, v92, v83, v80, v96)
			mBase = m.M
			v98 = m.ExcPending
			if v98 != 0 {
				return
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(l1))) = v97
				v100 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
				*(*int32)(unsafe.Add(mBase, uint32(l1)+8)) = v96 - v100
				v103 = v100
				v104 = v97
				if v74 != 0 {
					v110 = int32(-61)
				} else {
					v110 = int32(-62)
				}
				*(*uint8)(unsafe.Add(mBase, uint32(v104+v103))) = uint8(v110)
				v112 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
				*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v112 + int32(1)
				v116 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
				*(*int32)(unsafe.Add(mBase, uint32(l1)+8)) = v116 + int32(-1)
				m.G0 = v10 + int32(16)
				return
			}
		}
	} else {
		v78 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
		v79 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
		v103 = v78
		v104 = v79
		if v74 != 0 {
			v110 = int32(-61)
		} else {
			v110 = int32(-62)
		}
		*(*uint8)(unsafe.Add(mBase, uint32(v104+v103))) = uint8(v110)
		v112 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
		*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v112 + int32(1)
		v116 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
		*(*int32)(unsafe.Add(mBase, uint32(l1)+8)) = v116 + int32(-1)
		m.G0 = v10 + int32(16)
		return
	}
}
func F_mp_encode_lua_table_as_array(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v25 int32
	_ = v25
	var v30 int32
	_ = v30
	var v34 int32
	_ = v34
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	v7 = F_lua_objlen(m, l0, int32(-1))
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	F_mp_encode_array(m, l0, l1, base.I64_extend_i32_u(v7))
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v13 = m.G3
	F_luaL_checkstack(m, l0, int32(1), v13+int32(_a_F_mp_encode_lua_table_as_array_0))
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	if v7 == int32(0) {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	return
L6:
	;
	v20 = int32(1)
	v25 = v20
	goto L7
L7:
	;
	v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v30)+8)) = int32(3)
	*(*float64)(unsafe.Add(mBase, uint32(v30))) = base.F64_convert_i32_u(v25)
	v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v34 + int32(16)
	goto L9
L8:
	;
	goto L5
L9:
	;
	F_lua_gettable(m, l0, int32(-2))
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
		goto L1
	} else {
		goto L10
	}
L10:
	;
	F_mp_encode_lua_type(m, l0, l1, l2+v20)
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L1
	} else {
		goto L11
	}
L11:
	;
	v44 = v25 + int32(1)
	if base.Ui32(v44) <= base.Ui32(v7) {
		v25 = v44
		goto L7
	} else {
		goto L12
	}
L12:
	;
	goto L8
}
func F_mp_encode_lua_table_as_map(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v28 int32
	_ = v28
	var v47 int32
	_ = v47
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v66 int64
	_ = v66
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v83 int32
	_ = v83
	var v106 int32
	_ = v106
	var v109 int32
	_ = v109
	var v145 int32
	_ = v145
	var v146 int64
	_ = v146
	var v148 int32
	_ = v148
	var v150 int32
	_ = v150
	var v155 int32
	_ = v155
	var v157 int32
	_ = v157
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	v7 = m.G3
	F_luaL_checkstack(m, l0, int32(3), v7+int32(_a_F_mp_encode_lua_table_as_map_0))
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v13)+8)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v13 + int32(16)
	goto L3
L3:
	;
	v20 = F_lua_next(m, l0, int32(-2))
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		goto L1
	} else {
		goto L6
	}
L4:
	;
	F_mp_encode_map(m, l0, l1, v66)
	mBase = m.M
	v69 = m.ExcPending
	if v69 != 0 {
		goto L1
	} else {
		goto L20
	}
L5:
	;
	v28 = int32(0)
	goto L8
L6:
	;
	if v20 != 0 {
		goto L5
	} else {
		goto L7
	}
L7:
	;
	v66 = int64(0)
	goto L4
L8:
	;
	goto L12
L9:
	;
	v66 = base.I64_extend_i32_u(v58)
	goto L4
L10:
	;
	v58 = v28 + int32(1)
	v60 = F_lua_next(m, l0, int32(-2))
	mBase = m.M
	v61 = m.ExcPending
	if v61 != 0 {
		goto L1
	} else {
		goto L18
	}
L11:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v47 + int32(-16)
	goto L10
L12:
	;
	v47 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	goto L11
L18:
	;
	if v60 != 0 {
		v28 = v58
		goto L8
	} else {
		goto L19
	}
L19:
	;
	goto L9
L20:
	;
	v71 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v71)+8)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v71 + int32(16)
	goto L21
L21:
	;
	v78 = F_lua_next(m, l0, int32(-2))
	mBase = m.M
	v79 = m.ExcPending
	if v79 != 0 {
		goto L1
	} else {
		goto L23
	}
L22:
	;
	return
L23:
	;
	if v78 == int32(0) {
		goto L22
	} else {
		goto L24
	}
L24:
	;
	v83 = l2 + int32(1)
	goto L25
L25:
	;
	goto L29
L26:
	;
	goto L22
L27:
	;
	F_mp_encode_lua_type(m, l0, l1, v83)
	mBase = m.M
	v155 = m.ExcPending
	if v155 != 0 {
		goto L1
	} else {
		goto L43
	}
L28:
	;
	v145 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v146 = *(*int64)(unsafe.Add(mBase, uint32(v109)))
	*(*int64)(unsafe.Add(mBase, uint32(v145))) = v146
	v148 = *(*int32)(unsafe.Add(mBase, uint32(v109)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v145)+8)) = v148
	v150 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v150 + int32(16)
	goto L27
L29:
	;
	goto L35
L35:
	;
	v106 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v109 = v106 + int32(-32)
	goto L28
L43:
	;
	F_mp_encode_lua_type(m, l0, l1, v83)
	mBase = m.M
	v157 = m.ExcPending
	if v157 != 0 {
		goto L1
	} else {
		goto L44
	}
L44:
	;
	v159 = F_lua_next(m, l0, int32(-2))
	mBase = m.M
	v160 = m.ExcPending
	if v160 != 0 {
		goto L1
	} else {
		goto L45
	}
L45:
	;
	if v159 != 0 {
		goto L25
	} else {
		goto L46
	}
L46:
	;
	goto L26
}
func F_mp_pack(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v60 int32
	_ = v60
	var v64 int32
	_ = v64
	var v68 int32
	_ = v68
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
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
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v124 int32
	_ = v124
	var v125 int64
	_ = v125
	var v127 int32
	_ = v127
	var v129 int32
	_ = v129
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v149 int32
	_ = v149
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v162 int32
	_ = v162
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v174 int32
	_ = v174
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v183 int32
	_ = v183
	var v188 int32
	_ = v188
	v8 = m.G0
	v10 = v8 - int32(16)
	m.G0 = v10
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v16 = (v12 - v13) >> (uint(int32(4)) % 32)
	goto L3
L1:
	;
	m.G0 = v10 + int32(16)
	return v188
L2:
	;
	v25 = F_lua_checkstack(m, l0, v16)
	mBase = m.M
	v26 = m.ExcPending
	if v26 != 0 {
		goto L5
	} else {
		goto L8
	}
L3:
	;
	if v16 != 0 {
		goto L2
	} else {
		goto L4
	}
L4:
	;
	v18 = m.G3
	v21 = F_luaL_argerror(m, l0, int32(0), v18+int32(_a_F_mp_pack_0))
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	return int32(0)
L6:
	;
	v188 = v21
	goto L1
L7:
	;
	v34 = v10 + int32(4)
	v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v34 == int32(0) {
		goto L12
	} else {
		goto L13
	}
L8:
	;
	if v25 != 0 {
		goto L7
	} else {
		goto L9
	}
L9:
	;
	v28 = m.G3
	v31 = F_luaL_argerror(m, l0, int32(0), v28+int32(_a_F_mp_pack_1))
	mBase = m.M
	v32 = m.ExcPending
	if v32 != 0 {
		goto L5
	} else {
		goto L10
	}
L10:
	;
	v188 = v31
	goto L1
L11:
	;
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v10)+4))
	v42 = int32(0)
	v45 = m.T0[v40].(func(*base.Module, int32, int32, int32, int32) int32)(m, v41, v42, v42, int32(12))
	mBase = m.M
	v46 = m.ExcPending
	if v46 != 0 {
		goto L5
	} else {
		goto L14
	}
L12:
	;
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v35)+12))
	goto L11
L13:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v35)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v34))) = v38
	goto L12
L14:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v45)+8)) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v45))) = int64(0)
	v51 = int32(1)
	if v51 <= v16 {
		goto L16
	} else {
		goto L17
	}
L15:
	;
	v158 = v10 + int32(8)
	v159 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v158 == int32(0) {
		goto L41
	} else {
		goto L42
	}
L16:
	;
	v60 = v51
	goto L18
L17:
	;
	v54 = int32(0)
	v154 = v54
	v155 = v54
	goto L15
L18:
	;
	v64 = m.G3
	F_luaL_checkstack(m, l0, int32(1), v64+int32(_a_F_mp_pack_2))
	mBase = m.M
	v68 = m.ExcPending
	if v68 != 0 {
		goto L5
	} else {
		goto L20
	}
L19:
	;
	v149 = *(*int32)(unsafe.Add(mBase, uint32(v45)))
	v154 = v149
	v155 = v144
	goto L15
L20:
	;
	if v60 < int32(1) {
		goto L23
	} else {
		goto L24
	}
L21:
	;
	F_mp_encode_lua_type(m, l0, v45, int32(0))
	mBase = m.M
	v135 = m.ExcPending
	if v135 != 0 {
		goto L5
	} else {
		goto L37
	}
L22:
	;
	v124 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v125 = *(*int64)(unsafe.Add(mBase, uint32(v121)))
	*(*int64)(unsafe.Add(mBase, uint32(v124))) = v125
	v127 = *(*int32)(unsafe.Add(mBase, uint32(v121)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v124)+8)) = v127
	v129 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v129 + int32(16)
	goto L21
L23:
	;
	if v60 < int32(-9999) {
		goto L28
	} else {
		goto L29
	}
L24:
	;
	v73 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v74 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v79 = v74 + v60<<(uint(int32(4))%32) + int32(-16)
	v80 = m.G398
	if base.Ui32(v79) < base.Ui32(v73) {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	v82 = v79
	goto L27
L26:
	;
	v82 = v80
	goto L27
L27:
	;
	v121 = v82
	goto L22
L28:
	;
	switch v60 + int32(10002) {
	case 0:
		goto L31
	case 1:
		goto L32
	case 2:
		goto L33
	default:
		goto L30
	}
L29:
	;
	v85 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v121 = v85 + v60<<(uint(int32(4))%32)
	goto L22
L30:
	;
	v105 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v106 = *(*int32)(unsafe.Add(mBase, uint32(v105)+4))
	v107 = *(*int32)(unsafe.Add(mBase, uint32(v106)))
	v108 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v107)+7)))
	v109 = m.G398
	if base.Ui32(v108) < base.Ui32(int32(-10002)-v60) {
		goto L34
	} else {
		goto L35
	}
L31:
	;
	v121 = l0 + int32(72)
	goto L22
L32:
	;
	v94 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v95 = *(*int32)(unsafe.Add(mBase, uint32(v94)+4))
	v96 = *(*int32)(unsafe.Add(mBase, uint32(v95)))
	v97 = *(*int32)(unsafe.Add(mBase, uint32(v96)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+96)) = int32(5)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+88)) = v97
	v121 = l0 + int32(88)
	goto L22
L33:
	;
	v91 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v121 = v91 + int32(96)
	goto L22
L34:
	;
	v120 = v109
	goto L36
L35:
	;
	v120 = v107 + (int32(-10003)-v60)<<(uint(int32(4))%32) + int32(24)
	goto L36
L36:
	;
	v121 = v120
	goto L22
L37:
	;
	v136 = *(*int32)(unsafe.Add(mBase, uint32(v45)))
	v137 = *(*int32)(unsafe.Add(mBase, uint32(v45)+4))
	F_lua_pushlstring(m, l0, v136, v137)
	mBase = m.M
	v139 = m.ExcPending
	if v139 != 0 {
		goto L5
	} else {
		goto L38
	}
L38:
	;
	v140 = *(*int32)(unsafe.Add(mBase, uint32(v45)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v45)+4)) = int32(0)
	v143 = *(*int32)(unsafe.Add(mBase, uint32(v45)+8))
	v144 = v140 + v143
	*(*int32)(unsafe.Add(mBase, uint32(v45)+8)) = v144
	if v60 != v16 {
		v60 = v60 + int32(1)
		goto L18
	} else {
		goto L39
	}
L39:
	;
	goto L19
L40:
	;
	v165 = *(*int32)(unsafe.Add(mBase, uint32(v10)+8))
	v167 = m.T0[v164].(func(*base.Module, int32, int32, int32, int32) int32)(m, v165, v154, v155, int32(0))
	mBase = m.M
	v168 = m.ExcPending
	if v168 != 0 {
		goto L5
	} else {
		goto L43
	}
L41:
	;
	v164 = *(*int32)(unsafe.Add(mBase, uint32(v159)+12))
	goto L40
L42:
	;
	v162 = *(*int32)(unsafe.Add(mBase, uint32(v159)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v158))) = v162
	goto L41
L43:
	;
	v170 = v10 + int32(12)
	v171 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v170 == int32(0) {
		goto L45
	} else {
		goto L46
	}
L44:
	;
	v177 = *(*int32)(unsafe.Add(mBase, uint32(v10)+12))
	v180 = m.T0[v176].(func(*base.Module, int32, int32, int32, int32) int32)(m, v177, v45, int32(12), int32(0))
	mBase = m.M
	v181 = m.ExcPending
	if v181 != 0 {
		goto L5
	} else {
		goto L47
	}
L45:
	;
	v176 = *(*int32)(unsafe.Add(mBase, uint32(v171)+12))
	goto L44
L46:
	;
	v174 = *(*int32)(unsafe.Add(mBase, uint32(v171)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v170))) = v174
	goto L45
L47:
	;
	F_lua_concat(m, l0, v16)
	mBase = m.M
	v183 = m.ExcPending
	if v183 != 0 {
		goto L5
	} else {
		goto L48
	}
L48:
	;
	v188 = int32(1)
	goto L1
}
func F_mp_unpack(m *base.Module, l0 int32) int32 {
	var v2 int32
	_ = v2
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	v2 = int32(0)
	v4 = F_mp_unpack_full(m, l0, v2, v2)
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		return v4
	}
}
