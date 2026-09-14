package p2

import (
	base "github.com/shibukawa/valkeymem/internal/aot/vkaot/base"
	"unsafe"
)

func F_zzlFirstInRange(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v29 int32
	_ = v29
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v58 int64
	_ = v58
	var v63 int64
	_ = v63
	var v71 int32
	_ = v71
	var v78 int32
	_ = v78
	var v84 float64
	_ = v84
	var v88 int64
	_ = v88
	var v90 float64
	_ = v90
	var v91 float64
	_ = v91
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v99 float64
	_ = v99
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v111 int32
	_ = v111
	var v124 int32
	_ = v124
	v3 = int32(0)
	v9 = m.G0
	v11 = v9 - int32(16)
	m.G0 = v11
	v15 = F_lpSeek(m, l0, v3)
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v19 = F_zzlIsInRange(m, l0, l1)
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		goto L1
	} else {
		goto L5
	}
L3:
	;
	F__serverAssert(m, int32(_a169), int32(_a1723), int32(1000))
	mBase = m.M
	v124 = m.ExcPending
	if v124 != 0 {
		goto L1
	} else {
		goto L37
	}
L4:
	;
	m.G0 = v11 + int32(16)
	return v111
L5:
	;
	if v19 == int32(0) {
		v111 = v3
		goto L4
	} else {
		goto L6
	}
L6:
	;
	if v15 == int32(0) {
		v111 = v3
		goto L4
	} else {
		goto L7
	}
L7:
	;
	v29 = v15
	goto L8
L8:
	;
	v33 = F_lpNext(m, l0, v29)
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
		goto L1
	} else {
		goto L10
	}
L9:
	;
	v111 = int32(0)
	goto L4
L10:
	;
	if v33 == int32(0) {
		goto L3
	} else {
		goto L11
	}
L11:
	;
	v39 = F_lpGetValue(m, v33, v11+int32(12), v11)
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
		goto L1
	} else {
		goto L14
	}
L12:
	;
	v91 = *(*float64)(unsafe.Add(mBase, uint32(l1)))
	v94 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	if v94 != 0 {
		goto L25
	} else {
		goto L26
	}
L13:
	;
	v88 = *(*int64)(unsafe.Add(mBase, uint32(v11)))
	v90 = base.F64_convert_i64_s(v88)
	goto L12
L14:
	;
	if v39 == int32(0) {
		goto L13
	} else {
		goto L15
	}
L15:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v11)+12))
	v44 = int32(0)
	v48 = m.G0
	v50 = v48 - int32(32)
	m.G0 = v50
	v52 = F___errno_location(m)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v52))) = v44
	v58 = *(*int64)(unsafe.Add(mBase, _consts[378]))
	*(*int64)(unsafe.Add(mBase, uint32(v50+int32(8)))) = v58
	*(*int64)(unsafe.Add(mBase, uint32(v50)+24)) = int64(0)
	v63 = *(*int64)(unsafe.Add(mBase, _consts[379]))
	*(*int64)(unsafe.Add(mBase, uint32(v50))) = v63
	F_ffc_from_chars_double_options(m, v50+int32(16), v39, v39+v43, v50+int32(24), v50)
	mBase = m.M
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v50)+20))
	if v71 == v44 {
		goto L17
	} else {
		goto L18
	}
L16:
	;
	v90 = v84
	goto L12
L17:
	;
	goto L22
L18:
	;
	if v71 == int32(2) {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	v78 = int32(68)
	goto L21
L20:
	;
	v78 = int32(28)
	goto L21
L21:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v52))) = v78
	goto L17
L22:
	;
	v84 = *(*float64)(unsafe.Add(mBase, uint32(v50)+24))
	m.G0 = v50 + int32(32)
	goto L16
L24:
	;
	v105 = F_lpNext(m, l0, v33)
	mBase = m.M
	v106 = m.ExcPending
	if v106 != 0 {
		goto L1
	} else {
		goto L35
	}
L25:
	;
	v95 = base.F64_gt(v90, v91)
	goto L27
L26:
	;
	v95 = base.F64_ge(v90, v91)
	goto L27
L27:
	;
	if v95 != int32(1) {
		goto L24
	} else {
		goto L28
	}
L28:
	;
	v99 = *(*float64)(unsafe.Add(mBase, uint32(l1)+8))
	v102 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	if v102 != 0 {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	v103 = base.F64_lt(v90, v99)
	goto L31
L30:
	;
	v103 = base.F64_le(v90, v99)
	goto L31
L31:
	;
	if v103 != 0 {
		goto L32
	} else {
		goto L33
	}
L32:
	;
	v104 = v29
	goto L34
L33:
	;
	v104 = int32(0)
	goto L34
L34:
	;
	v111 = v104
	goto L4
L35:
	;
	if v105 != 0 {
		v29 = v105
		goto L8
	} else {
		goto L36
	}
L36:
	;
	goto L9
L37:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_zzlLastInRange(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v56 int64
	_ = v56
	var v61 int64
	_ = v61
	var v69 int32
	_ = v69
	var v76 int32
	_ = v76
	var v82 float64
	_ = v82
	var v86 int64
	_ = v86
	var v88 float64
	_ = v88
	var v89 float64
	_ = v89
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v97 float64
	_ = v97
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v112 int32
	_ = v112
	var v118 int32
	_ = v118
	var v123 int32
	_ = v123
	v3 = int32(0)
	v8 = m.G0
	v10 = v8 - int32(16)
	m.G0 = v10
	v14 = F_lpSeek(m, l0, int32(-2))
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v18 = F_zzlIsInRange(m, l0, l1)
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		goto L1
	} else {
		goto L4
	}
L3:
	;
	m.G0 = v10 + int32(16)
	return v123
L4:
	;
	if v18 == int32(0) {
		v123 = v3
		goto L3
	} else {
		goto L5
	}
L5:
	;
	if v14 == int32(0) {
		v123 = v3
		goto L3
	} else {
		goto L6
	}
L6:
	;
	v28 = v14
	goto L8
L7:
	;
	F__serverAssert(m, int32(_a169), int32(_a1723), int32(1027))
	mBase = m.M
	v118 = m.ExcPending
	if v118 != 0 {
		goto L1
	} else {
		goto L41
	}
L8:
	;
	v31 = F_lpNext(m, l0, v28)
	mBase = m.M
	v32 = m.ExcPending
	if v32 != 0 {
		goto L1
	} else {
		goto L10
	}
L9:
	;
	F__serverAssert(m, int32(_a1724), int32(_a1723), int32(1040))
	mBase = m.M
	v112 = m.ExcPending
	if v112 != 0 {
		goto L1
	} else {
		goto L40
	}
L10:
	;
	if v31 == int32(0) {
		goto L7
	} else {
		goto L11
	}
L11:
	;
	v37 = F_lpGetValue(m, v31, v10+int32(12), v10)
	mBase = m.M
	v38 = m.ExcPending
	if v38 != 0 {
		goto L1
	} else {
		goto L14
	}
L12:
	;
	v89 = *(*float64)(unsafe.Add(mBase, uint32(l1)+8))
	v92 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	if v92 != 0 {
		goto L25
	} else {
		goto L26
	}
L13:
	;
	v86 = *(*int64)(unsafe.Add(mBase, uint32(v10)))
	v88 = base.F64_convert_i64_s(v86)
	goto L12
L14:
	;
	if v37 == int32(0) {
		goto L13
	} else {
		goto L15
	}
L15:
	;
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v10)+12))
	v42 = int32(0)
	v46 = m.G0
	v48 = v46 - int32(32)
	m.G0 = v48
	v50 = F___errno_location(m)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v50))) = v42
	v56 = *(*int64)(unsafe.Add(mBase, _consts[378]))
	*(*int64)(unsafe.Add(mBase, uint32(v48+int32(8)))) = v56
	*(*int64)(unsafe.Add(mBase, uint32(v48)+24)) = int64(0)
	v61 = *(*int64)(unsafe.Add(mBase, _consts[379]))
	*(*int64)(unsafe.Add(mBase, uint32(v48))) = v61
	F_ffc_from_chars_double_options(m, v48+int32(16), v37, v37+v41, v48+int32(24), v48)
	mBase = m.M
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v48)+20))
	if v69 == v42 {
		goto L17
	} else {
		goto L18
	}
L16:
	;
	v88 = v82
	goto L12
L17:
	;
	goto L22
L18:
	;
	if v69 == int32(2) {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	v76 = int32(68)
	goto L21
L20:
	;
	v76 = int32(28)
	goto L21
L21:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v50))) = v76
	goto L17
L22:
	;
	v82 = *(*float64)(unsafe.Add(mBase, uint32(v48)+24))
	m.G0 = v48 + int32(32)
	goto L16
L24:
	;
	v103 = F_lpPrev(m, l0, v28)
	mBase = m.M
	v104 = m.ExcPending
	if v104 != 0 {
		goto L1
	} else {
		goto L36
	}
L25:
	;
	v93 = base.F64_lt(v88, v89)
	goto L27
L26:
	;
	v93 = base.F64_le(v88, v89)
	goto L27
L27:
	;
	if v93 != int32(1) {
		goto L24
	} else {
		goto L28
	}
L28:
	;
	v97 = *(*float64)(unsafe.Add(mBase, uint32(l1)))
	v100 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	if v100 != 0 {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	v101 = base.F64_gt(v88, v97)
	goto L31
L30:
	;
	v101 = base.F64_ge(v88, v97)
	goto L31
L31:
	;
	if v101 != 0 {
		goto L32
	} else {
		goto L33
	}
L32:
	;
	v102 = v28
	goto L34
L33:
	;
	v102 = int32(0)
	goto L34
L34:
	;
	v123 = v102
	goto L3
L35:
	;
	v106 = F_lpPrev(m, l0, v103)
	mBase = m.M
	v107 = m.ExcPending
	if v107 != 0 {
		goto L1
	} else {
		goto L38
	}
L36:
	;
	if v103 != 0 {
		goto L35
	} else {
		goto L37
	}
L37:
	;
	v123 = int32(0)
	goto L3
L38:
	;
	if v106 != 0 {
		v28 = v106
		goto L8
	} else {
		goto L39
	}
L39:
	;
	goto L9
L40:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L41:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_zzlValidateScores(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v20 int32
	_ = v20
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v50 int64
	_ = v50
	var v55 int64
	_ = v55
	var v63 int32
	_ = v63
	var v70 int32
	_ = v70
	var v76 float64
	_ = v76
	var v80 int64
	_ = v80
	var v82 float64
	_ = v82
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v100 int32
	_ = v100
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	v12 = F_lpSeek(m, l0, int32(0))
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	m.G0 = v9 + int32(16)
	return v100
L2:
	;
	v100 = int32(1)
	goto L1
L3:
	;
	return int32(0)
L4:
	;
	if v12 == int32(0) {
		goto L2
	} else {
		goto L5
	}
L5:
	;
	v20 = v12
	goto L6
L6:
	;
	v24 = int32(0)
	v25 = F_lpNext(m, l0, v20)
	mBase = m.M
	v26 = m.ExcPending
	if v26 != 0 {
		goto L3
	} else {
		goto L8
	}
L7:
	;
	goto L2
L8:
	;
	if v25 == int32(0) {
		v100 = v24
		goto L1
	} else {
		goto L9
	}
L9:
	;
	v31 = F_lpGetValue(m, v25, v9+int32(12), v9)
	mBase = m.M
	v32 = m.ExcPending
	if v32 != 0 {
		goto L3
	} else {
		goto L12
	}
L10:
	;
	if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v82)&int64(9223372036854775807)) {
		v100 = v24
		goto L1
	} else {
		goto L22
	}
L11:
	;
	v80 = *(*int64)(unsafe.Add(mBase, uint32(v9)))
	v82 = base.F64_convert_i64_s(v80)
	goto L10
L12:
	;
	if v31 == int32(0) {
		goto L11
	} else {
		goto L13
	}
L13:
	;
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v9)+12))
	v36 = int32(0)
	v40 = m.G0
	v42 = v40 - int32(32)
	m.G0 = v42
	v44 = F___errno_location(m)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v44))) = v36
	v50 = *(*int64)(unsafe.Add(mBase, _consts[378]))
	*(*int64)(unsafe.Add(mBase, uint32(v42+int32(8)))) = v50
	*(*int64)(unsafe.Add(mBase, uint32(v42)+24)) = int64(0)
	v55 = *(*int64)(unsafe.Add(mBase, _consts[379]))
	*(*int64)(unsafe.Add(mBase, uint32(v42))) = v55
	F_ffc_from_chars_double_options(m, v42+int32(16), v31, v31+v35, v42+int32(24), v42)
	mBase = m.M
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v42)+20))
	if v63 == v36 {
		goto L15
	} else {
		goto L16
	}
L14:
	;
	v82 = v76
	goto L10
L15:
	;
	goto L20
L16:
	;
	if v63 == int32(2) {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	v70 = int32(68)
	goto L19
L18:
	;
	v70 = int32(28)
	goto L19
L19:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44))) = v70
	goto L15
L20:
	;
	v76 = *(*float64)(unsafe.Add(mBase, uint32(v42)+24))
	m.G0 = v42 + int32(32)
	goto L14
L22:
	;
	v88 = F_lpNext(m, l0, v25)
	mBase = m.M
	v89 = m.ExcPending
	if v89 != 0 {
		goto L3
	} else {
		goto L23
	}
L23:
	;
	if v88 != 0 {
		v20 = v88
		goto L6
	} else {
		goto L24
	}
L24:
	;
	goto L7
}
