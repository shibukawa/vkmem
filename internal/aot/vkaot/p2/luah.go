package p2

import (
	base "github.com/shibukawa/vkmem/internal/aot/vkaot/base"
	"unsafe"
)

func F_luaH_get(m *base.Module, l0 int32, l1 int32) int32 {
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
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v22 int32
	_ = v22
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 float64
	_ = v37
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v46 float64
	_ = v46
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v65 int64
	_ = v65
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v77 int32
	_ = v77
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v101 int32
	_ = v101
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v114 int32
	_ = v114
	var v118 int32
	_ = v118
	var v123 int32
	_ = v123
	var v131 int32
	_ = v131
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v139 float64
	_ = v139
	var v140 float64
	_ = v140
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v151 int32
	_ = v151
	var v153 int32
	_ = v153
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v159 float64
	_ = v159
	var v160 int32
	_ = v160
	var v161 int64
	_ = v161
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v173 int32
	_ = v173
	var v177 int32
	_ = v177
	var v178 float64
	_ = v178
	var v180 int32
	_ = v180
	var v187 int32
	_ = v187
	var v190 float64
	_ = v190
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v198 int32
	_ = v198
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v9 = m.G398
	switch v8 {
	case 0:
		v198 = v9
		goto L1
	case 1:
		goto L7
	case 2:
		goto L6
	case 3:
		goto L8
	case 4:
		goto L9
	default:
		goto L5
	}
L1:
	;
	return v198
L2:
	;
	v180 = v177
	goto L41
L3:
	;
	v160 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v161 = base.I64_reinterpret_f64(v159)
	v166 = int32(-1)
	v167 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+12)))
	v173 = base.I32_rem_u_s(base.I32_wrap_i64(int64(base.Ui64(v161)>>(uint(int64(32))%64))+v161), v166<<(uint(v167)%32)^v166|int32(1))
	v177 = v160 + v173<<(uint(int32(5))%32)
	v178 = v159
	goto L2
L4:
	;
	v123 = v118
	goto L28
L5:
	;
	v105 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v106 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v107 = int32(-1)
	v108 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+12)))
	v114 = base.I32_rem_u_s(v106, v107<<(uint(v108)%32)^v107|int32(1))
	v118 = v105 + v114<<(uint(int32(5))%32)
	goto L4
L6:
	;
	v92 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v93 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v94 = int32(-1)
	v95 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+12)))
	v101 = base.I32_rem_u_s(v93, v94<<(uint(v95)%32)^v94|int32(1))
	v118 = v92 + v101<<(uint(int32(5))%32)
	goto L4
L7:
	;
	v81 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v82 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v83 = int32(-1)
	v84 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+12)))
	v118 = v81 + v82&(v83<<(uint(v84)%32)^v83)<<(uint(int32(5))%32)
	goto L4
L8:
	;
	v37 = *(*float64)(unsafe.Add(mBase, uint32(l1)))
	if base.F64_lt(base.F64_abs(v37), float64(2.147483648e+09)) == int32(0) {
		goto L17
	} else {
		goto L18
	}
L9:
	;
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(v11)+8))
	v13 = int32(-1)
	v14 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+12)))
	v22 = v10 + v12&(v13<<(uint(v14)%32)^v13)<<(uint(int32(5))%32)
	goto L10
L10:
	;
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v22)+24))
	if v29 != int32(4) {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v22)+28))
	v36 = m.G398
	if v35 != 0 {
		v22 = v35
		goto L10
	} else {
		goto L15
	}
L13:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v22)+16))
	if v32 != v11 {
		goto L12
	} else {
		goto L14
	}
L14:
	;
	return v22
L15:
	;
	v198 = v36
	goto L1
L16:
	;
	v46 = base.F64_convert_i32_s(v45)
	if base.F64_ne(v37, v46) != 0 {
		goto L19
	} else {
		goto L20
	}
L17:
	;
	v45 = int32(-2147483648)
	goto L16
L18:
	;
	v43 = base.I32_trunc_f64_s(v37)
	v45 = v43
	goto L16
L19:
	;
	if base.F64_ne(v37, float64(0)) != 0 {
		goto L26
	} else {
		goto L27
	}
L20:
	;
	if v45 < int32(1) {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	if v45 != 0 {
		v159 = v46
		goto L3
	} else {
		goto L25
	}
L22:
	;
	v50 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	if v45 <= v50 {
		goto L23
	} else {
		goto L24
	}
L23:
	;
	v53 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	return v53 + v45<<(uint(int32(4))%32) + int32(-16)
L24:
	;
	v159 = base.F64_convert_i32_u(v45)
	goto L3
L25:
	;
	v60 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v177 = v60
	v178 = v46
	goto L2
L26:
	;
	v64 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v65 = base.I64_reinterpret_f64(v37)
	v70 = int32(-1)
	v71 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+12)))
	v77 = base.I32_rem_u_s(base.I32_wrap_i64(int64(base.Ui64(v65)>>(uint(int64(32))%64))+v65), v70<<(uint(v71)%32)^v70|int32(1))
	v118 = v64 + v77<<(uint(int32(5))%32)
	goto L4
L27:
	;
	v63 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v118 = v63
	goto L4
L28:
	;
	v131 = v123 + int32(16)
	v134 = *(*int32)(unsafe.Add(mBase, uint32(v131)+8))
	v135 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	if v134 == v135 {
		goto L32
	} else {
		goto L33
	}
L30:
	;
	v157 = *(*int32)(unsafe.Add(mBase, uint32(v123)+28))
	v158 = m.G398
	if v157 != 0 {
		v123 = v157
		goto L28
	} else {
		goto L40
	}
L31:
	;
	if v153 == int32(0) {
		goto L30
	} else {
		goto L39
	}
L32:
	;
	switch v134 {
	case 0:
		v151 = int32(1)
		goto L34
	case 1:
		goto L37
	case 2:
		goto L36
	case 3:
		goto L38
	default:
		goto L35
	}
L33:
	;
	v153 = int32(0)
	goto L31
L34:
	;
	v153 = v151
	goto L31
L35:
	;
	v148 = *(*int32)(unsafe.Add(mBase, uint32(v131)))
	v149 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v151 = base.B2i32(v148 == v149)
	goto L34
L36:
	;
	v145 = *(*int32)(unsafe.Add(mBase, uint32(v131)))
	v146 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v153 = base.B2i32(v145 == v146)
	goto L31
L37:
	;
	v142 = *(*int32)(unsafe.Add(mBase, uint32(v131)))
	v143 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v153 = base.B2i32(v142 == v143)
	goto L31
L38:
	;
	v139 = *(*float64)(unsafe.Add(mBase, uint32(v131)))
	v140 = *(*float64)(unsafe.Add(mBase, uint32(l1)))
	v153 = base.F64_eq(v139, v140)
	goto L31
L39:
	;
	return v123
L40:
	;
	v198 = v158
	goto L1
L41:
	;
	v187 = *(*int32)(unsafe.Add(mBase, uint32(v180)+24))
	if v187 != int32(3) {
		goto L43
	} else {
		goto L44
	}
L42:
	;
	v198 = v194
	goto L1
L43:
	;
	v193 = *(*int32)(unsafe.Add(mBase, uint32(v180)+28))
	v194 = m.G398
	if v193 != 0 {
		v180 = v193
		goto L41
	} else {
		goto L46
	}
L44:
	;
	v190 = *(*float64)(unsafe.Add(mBase, uint32(v180)+16))
	if base.F64_ne(v190, v178) != 0 {
		goto L43
	} else {
		goto L45
	}
L45:
	;
	return v180
L46:
	;
	goto L42
}
func F_luaH_getnum(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v17 float64
	_ = v17
	var v18 int32
	_ = v18
	var v19 float64
	_ = v19
	var v20 int32
	_ = v20
	var v21 int64
	_ = v21
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v33 int32
	_ = v33
	var v37 int32
	_ = v37
	var v38 float64
	_ = v38
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v47 float64
	_ = v47
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	if l1 < int32(1) {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	v41 = v37
	goto L8
L2:
	;
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v21 = base.I64_reinterpret_f64(v19)
	v26 = int32(-1)
	v27 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+12)))
	v33 = base.I32_rem_u_s(base.I32_wrap_i64(int64(base.Ui64(v21)>>(uint(int64(32))%64))+v21), v26<<(uint(v27)%32)^v26|int32(1))
	v37 = v20 + v33<<(uint(int32(5))%32)
	v38 = v19
	goto L1
L3:
	;
	v17 = base.F64_convert_i32_s(l1)
	if l1 != 0 {
		v19 = v17
		goto L2
	} else {
		goto L7
	}
L4:
	;
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	if l1 <= v7 {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	return v10 + l1<<(uint(int32(4))%32) + int32(-16)
L6:
	;
	v19 = base.F64_convert_i32_u(l1)
	goto L2
L7:
	;
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v37 = v18
	v38 = v17
	goto L1
L8:
	;
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v41)+24))
	if v44 != int32(3) {
		goto L10
	} else {
		goto L11
	}
L9:
	;
	return v51
L10:
	;
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v41)+28))
	v51 = m.G398
	if v50 != 0 {
		v41 = v50
		goto L8
	} else {
		goto L13
	}
L11:
	;
	v47 = *(*float64)(unsafe.Add(mBase, uint32(v41)+16))
	if base.F64_ne(v47, v38) != 0 {
		goto L10
	} else {
		goto L12
	}
L12:
	;
	return v41
L13:
	;
	goto L9
}
func F_luaH_getstr(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v6 = int32(-1)
	v7 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+12)))
	v15 = v4 + v5&(v6<<(uint(v7)%32)^v6)<<(uint(int32(5))%32)
	goto L1
L1:
	;
	v18 = *(*int32)(unsafe.Add(mBase, uint32(v15)+24))
	if v18 != int32(4) {
		goto L3
	} else {
		goto L4
	}
L2:
	;
	return v25
L3:
	;
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v15)+28))
	v25 = m.G398
	if v24 != 0 {
		v15 = v24
		goto L1
	} else {
		goto L6
	}
L4:
	;
	v21 = *(*int32)(unsafe.Add(mBase, uint32(v15)+16))
	if v21 != l1 {
		goto L3
	} else {
		goto L5
	}
L5:
	;
	return v15
L6:
	;
	goto L2
}
func F_luaH_resizearray(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	v5 = m.G3
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
	if v6 != v5+int32(_a_F_luaH_resizearray_0) {
		v14 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+12)))
		F_resize_2(m, l0, l1, l2, int32(1)<<(uint(v14)%32))
		mBase = m.M
		v17 = m.ExcPending
		if v17 != 0 {
			return
		} else {
			return
		}
	} else {
		F_resize_2(m, l0, l1, l2, int32(0))
		mBase = m.M
		v12 = m.ExcPending
		if v12 != 0 {
			return
		} else {
			return
		}
	}
}
func F_luaH_setstr(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	v12 = int32(-1)
	v13 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+12)))
	v25 = v10 + v11&(v12<<(uint(v13)%32)^v12)<<(uint(int32(5))%32)
	goto L4
L1:
	;
	m.G0 = v8 + int32(16)
	return v42
L2:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8)+8)) = int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v8))) = l2
	v38 = F_newkey(m, l0, l1, v8)
	mBase = m.M
	v41 = m.ExcPending
	if v41 != 0 {
		goto L11
	} else {
		goto L12
	}
L3:
	;
	v32 = m.G398
	if v25 != v32 {
		v42 = v25
		goto L1
	} else {
		goto L10
	}
L4:
	;
	v26 = *(*int32)(unsafe.Add(mBase, uint32(v25)+24))
	if v26 != int32(4) {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v25)+28))
	if v31 != 0 {
		v25 = v31
		goto L4
	} else {
		goto L9
	}
L7:
	;
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v25)+16))
	if v29 == l2 {
		goto L3
	} else {
		goto L8
	}
L8:
	;
	goto L6
L9:
	;
	goto L2
L10:
	;
	goto L2
L11:
	;
	return int32(0)
L12:
	;
	v42 = v38
	goto L1
}
