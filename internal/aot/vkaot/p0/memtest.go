package p0

import (
	base "github.com/shibukawa/vkmem/internal/aot/vkaot/base"
	"unsafe"
)

func F_memtest_compare(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v19 int32
	_ = v19
	var v24 int32
	_ = v24
	var v25 int64
	_ = v25
	var v27 int32
	_ = v27
	var v31 int32
	_ = v31
	var v34 int64
	_ = v34
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v61 int64
	_ = v61
	var v63 int64
	_ = v63
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v106 int32
	_ = v106
	var v111 int64
	_ = v111
	var v124 int32
	_ = v124
	var v140 int32
	_ = v140
	v10 = m.G0
	v12 = v10 - int32(16)
	m.G0 = v12
	if l1&int32(4095) != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	F__serverAssert(m, int32(_a_F_memtest_compare_0), int32(_a_F_memtest_compare_1), int32(199))
	mBase = m.M
	v140 = m.ExcPending
	if v140 != 0 {
		goto L12
	} else {
		goto L25
	}
L2:
	;
	if base.Ui32(l1) < base.Ui32(int32(8)) {
		goto L4
	} else {
		goto L5
	}
L3:
	;
	m.G0 = v12 + int32(16)
	return v124
L4:
	;
	v124 = int32(0)
	goto L3
L5:
	;
	v19 = int32(base.Ui32(l1) >> (uint(int32(3)) % 32))
	v24 = *(*int32)(unsafe.Add(mBase, _c_F_memtest_compare[0]))
	v25 = base.I64_extend_i32_u(v19)
	v27 = l0
	v31 = l0 + v19<<(uint(int32(2))%32)
	v34 = int64(0)
	goto L6
L6:
	;
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v27)))
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v31)))
	if v36 == v37 {
		goto L8
	} else {
		goto L9
	}
L7:
	;
	goto L4
L8:
	;
	if l2 == int32(0) {
		goto L14
	} else {
		goto L15
	}
L9:
	;
	if l2 != 0 {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+12)) = v37
	*(*int32)(unsafe.Add(mBase, uint32(v12)+8)) = v36
	*(*int32)(unsafe.Add(mBase, uint32(v12)+4)) = v31
	*(*int32)(unsafe.Add(mBase, uint32(v12))) = v27
	v45 = F_iprintf(m, int32(_a_F_memtest_compare_2), v12)
	mBase = m.M
	v48 = m.ExcPending
	if v48 != 0 {
		goto L12
	} else {
		goto L13
	}
L11:
	;
	v124 = int32(1)
	goto L3
L12:
	;
	return int32(0)
L13:
	;
	m.Env.Exit(m, int32(1))
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L14:
	;
	v106 = int32(4)
	v111 = v34 + int64(1)
	if v111 != v25 {
		v27 = v27 + v106
		v31 = v31 + v106
		v34 = v111
		goto L6
	} else {
		goto L24
	}
L15:
	;
	if v34&int64(65535) != int64(0) {
		goto L14
	} else {
		goto L16
	}
L16:
	;
	v57 = int32(0)
	v59 = *(*int32)(unsafe.Add(mBase, _c_F_memtest_compare[1]))
	v61 = int64(*(*uint32)(unsafe.Add(mBase, _c_F_memtest_compare[2])))
	v63 = base.I64_div_u_s(v34*v61, v25)
	v64 = base.I32_wrap_i64(v63)
	if v59 == v64 {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_memtest_compare[1])) = v64
	v95 = F_fflush(m, v24)
	mBase = m.M
	v96 = m.ExcPending
	if v96 != 0 {
		goto L12
	} else {
		goto L23
	}
L18:
	;
	v67 = v57
	goto L19
L19:
	;
	v76 = F_putchar(m, int32(61))
	mBase = m.M
	v77 = m.ExcPending
	if v77 != 0 {
		goto L12
	} else {
		goto L21
	}
L20:
	;
	goto L17
L21:
	;
	v79 = v67 + int32(1)
	v81 = *(*int32)(unsafe.Add(mBase, _c_F_memtest_compare[1]))
	if base.Ui32(v79) < base.Ui32(v64-v81) {
		v67 = v79
		goto L19
	} else {
		goto L22
	}
L22:
	;
	goto L20
L23:
	;
	goto L14
L24:
	;
	goto L7
L25:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_memtest_fill_value(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) {
	mBase := m.M
	_ = mBase
	var v23 int64
	_ = v23
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v43 int64
	_ = v43
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v81 int64
	_ = v81
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v96 int64
	_ = v96
	var v98 int64
	_ = v98
	var v99 int32
	_ = v99
	var v102 int32
	_ = v102
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v122 int32
	_ = v122
	var v124 int32
	_ = v124
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v167 int32
	_ = v167
	var v172 int64
	_ = v172
	var v193 int64
	_ = v193
	var v200 int32
	_ = v200
	if l1&int32(4095) != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	F__serverAssert(m, int32(_a_F_memtest_fill_value_0), int32(_a_F_memtest_fill_value_1), int32(176))
	mBase = m.M
	v200 = m.ExcPending
	if v200 != 0 {
		goto L19
	} else {
		goto L25
	}
L2:
	;
	v23 = base.I64_extend_i32_u(int32(base.Ui32(l1) >> (uint(int32(13)) % 32)))
	v25 = *(*int32)(unsafe.Add(mBase, _c_F_memtest_fill_value[0]))
	v27 = int32(base.Ui32(l1) >> (uint(int32(3)) % 32))
	v43 = int64(0)
	goto L3
L3:
	;
	if base.Ui32(l1) < base.Ui32(int32(8192)) {
		goto L5
	} else {
		goto L6
	}
L4:
	;
	return
L5:
	;
	v193 = v43 + int64(1)
	if v193 != int64(1024) {
		v43 = v193
		goto L3
	} else {
		goto L24
	}
L6:
	;
	v50 = base.I32_wrap_i64(v43)
	v51 = int32(2)
	v53 = l0 + v50<<(uint(v51)%32)
	if v50&int32(1) != 0 {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	v59 = l3
	goto L9
L8:
	;
	v59 = l2
	goto L9
L9:
	;
	v62 = v59<<(uint(int32(16))%32) | v59
	v77 = v53
	v78 = v53 + v27<<(uint(v51)%32)
	v81 = int64(0)
	goto L10
L10:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v78))) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v77))) = v62
	if l5 == int32(0) {
		goto L12
	} else {
		goto L13
	}
L11:
	;
	goto L5
L12:
	;
	v167 = int32(4096)
	v172 = v81 + int64(1)
	if v172 != v23 {
		v77 = v77 + v167
		v78 = v78 + v167
		v81 = v172
		goto L10
	} else {
		goto L23
	}
L13:
	;
	if v81&int64(65535) != int64(0) {
		goto L12
	} else {
		goto L14
	}
L14:
	;
	v91 = int32(0)
	v93 = *(*int32)(unsafe.Add(mBase, _c_F_memtest_fill_value[1]))
	v96 = int64(*(*uint32)(unsafe.Add(mBase, _c_F_memtest_fill_value[2])))
	v98 = base.I64_div_u_s((v81+v43*v23)*v96, base.I64_extend_i32_u(v27))
	v99 = base.I32_wrap_i64(v98)
	if v93 == v99 {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_memtest_fill_value[1])) = v99
	v147 = F_fflush(m, v25)
	mBase = m.M
	v148 = m.ExcPending
	if v148 != 0 {
		goto L19
	} else {
		goto L22
	}
L16:
	;
	v102 = v91
	goto L17
L17:
	;
	v119 = F_putchar(m, l4)
	mBase = m.M
	v120 = m.ExcPending
	if v120 != 0 {
		goto L19
	} else {
		goto L20
	}
L18:
	;
	goto L15
L19:
	;
	return
L20:
	;
	v122 = v102 + int32(1)
	v124 = *(*int32)(unsafe.Add(mBase, _c_F_memtest_fill_value[1]))
	if base.Ui32(v122) < base.Ui32(v99-v124) {
		v102 = v122
		goto L17
	} else {
		goto L21
	}
L21:
	;
	goto L18
L22:
	;
	goto L12
L23:
	;
	goto L11
L24:
	;
	goto L4
L25:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_memtest_progress_start(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v19 int32
	_ = v19
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v37 int32
	_ = v37
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v66 int32
	_ = v66
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	v3 = int32(0)
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v12 = F_iprintf(m, int32(_a_F_memtest_progress_start_0), v3)
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v14 = int32(0)
	v15 = int32(*(*uint16)(unsafe.Add(mBase, _c_F_memtest_progress_start[0])))
	v19 = int32(*(*uint16)(unsafe.Add(mBase, _c_F_memtest_progress_start[1])))
	if (v15+int32(-2))*v19 < int32(1) {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v45 = F_puts(m, int32(_a_F_memtest_progress_start_1))
	mBase = m.M
	v46 = m.ExcPending
	if v46 != 0 {
		goto L1
	} else {
		goto L9
	}
L4:
	;
	v26 = v3
	goto L5
L5:
	;
	v28 = F_putchar(m, int32(46))
	mBase = m.M
	v29 = m.ExcPending
	if v29 != 0 {
		goto L1
	} else {
		goto L7
	}
L6:
	;
	goto L3
L7:
	;
	v31 = v26 + int32(1)
	v32 = int32(0)
	v33 = int32(*(*uint16)(unsafe.Add(mBase, _c_F_memtest_progress_start[0])))
	v37 = int32(*(*uint16)(unsafe.Add(mBase, _c_F_memtest_progress_start[1])))
	if v31 < (v33+int32(-2))*v37 {
		v26 = v31
		goto L5
	} else {
		goto L8
	}
L8:
	;
	goto L6
L9:
	;
	v49 = F_iprintf(m, int32(_a_F_memtest_progress_start_2), int32(0))
	mBase = m.M
	v50 = m.ExcPending
	if v50 != 0 {
		goto L1
	} else {
		goto L10
	}
L10:
	;
	v53 = F_iprintf(m, int32(_a_F_memtest_progress_start_3), int32(0))
	mBase = m.M
	v54 = m.ExcPending
	if v54 != 0 {
		goto L1
	} else {
		goto L11
	}
L11:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v7)+4)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v7))) = l0
	v58 = F_iprintf(m, int32(_a_F_memtest_progress_start_4), v7)
	mBase = m.M
	v59 = m.ExcPending
	if v59 != 0 {
		goto L1
	} else {
		goto L12
	}
L12:
	;
	v60 = int32(0)
	v62 = int32(*(*uint16)(unsafe.Add(mBase, _c_F_memtest_progress_start[0])))
	v66 = int32(*(*uint16)(unsafe.Add(mBase, _c_F_memtest_progress_start[1])))
	*(*int32)(unsafe.Add(mBase, _c_F_memtest_progress_start[2])) = (v62 + int32(-3)) * v66
	*(*int32)(unsafe.Add(mBase, _c_F_memtest_progress_start[3])) = v60
	v73 = *(*int32)(unsafe.Add(mBase, _c_F_memtest_progress_start[4]))
	v74 = F_fflush(m, v73)
	mBase = m.M
	v75 = m.ExcPending
	if v75 != 0 {
		goto L1
	} else {
		goto L13
	}
L13:
	;
	m.G0 = v7 + int32(16)
	return
}
