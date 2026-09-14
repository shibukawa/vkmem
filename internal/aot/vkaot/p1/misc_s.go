package p1

import (
	base "github.com/shibukawa/vkmem/internal/aot/vkaot/base"
	"unsafe"
)

func F_SHA1Init(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	v3 = m.Env.Vkmem_hash_create(m, int32(1))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v3
	return
}
func F_SHA1Update(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	m.Env.Vkmem_hash_update(m, v4, l1, l2)
	mBase = m.M
	return
}
func F___shlim(m *base.Module, l0 int32, l1 int64) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v11 int32
	_ = v11
	var v19 int32
	_ = v19
	*(*int64)(unsafe.Add(mBase, uint32(l0)+112)) = l1
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int64)(unsafe.Add(mBase, uint32(l0)+120)) = base.I64_extend_i32_s(v6 - v7)
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if l1 == int64(0) {
		v19 = v11
	} else {
		if base.I64_extend_i32_s(v11-v7) <= l1 {
			v19 = v11
		} else {
			v19 = v7 + base.I32_wrap_i64(l1)
		}
	}
	*(*int32)(unsafe.Add(mBase, uint32(l0)+104)) = v19
	return
}
func F___sin(m *base.Module, l0 float64, l1 float64, l2 int32) float64 {
	var v7 float64
	_ = v7
	var v22 float64
	_ = v22
	var v23 float64
	_ = v23
	v7 = base.F64_mul(l0, l0)
	v22 = base.F64_add(base.F64_mul(base.F64_mul(v7, base.F64_mul(v7, v7)), base.F64_add(base.F64_mul(v7, float64(1.58969099521155e-10)), float64(-2.5050760253406863e-08))), base.F64_add(base.F64_mul(v7, base.F64_add(base.F64_mul(v7, float64(2.7557313707070068e-06)), float64(-0.0001984126982985795))), float64(0.00833333333332249)))
	v23 = base.F64_mul(l0, v7)
	if l2 != 0 {
		return base.F64_sub(l0, base.F64_add(base.F64_sub(base.F64_mul(v7, base.F64_sub(base.F64_mul(l1, float64(0.5)), base.F64_mul(v23, v22))), l1), base.F64_mul(v23, float64(0.16666666666666632))))
	} else {
		return base.F64_add(base.F64_mul(v23, base.F64_add(base.F64_mul(v7, v22), float64(-0.16666666666666632))), l0)
	}
}
func F___small_vfprintf(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	v6 = F___vfprintf_internal(m, l0, l1, l2, int32(1387), int32(0))
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		return v6
	}
}
func F___small_vsnprintf(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v17 int32
	_ = v17
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v55 int32
	_ = v55
	v7 = m.G0
	v9 = v7 - int32(160)
	m.G0 = v9
	v17 = F__emscripten_memcpy_bulkmem(m, v9+int32(8), int32(_a_F___small_vsnprintf_0), int32(144))
	mBase = m.M
	if int32(0) < l1 {
		v24 = l0
		v25 = l1
		*(*int32)(unsafe.Add(mBase, uint32(v9)+52)) = v24
		*(*int32)(unsafe.Add(mBase, uint32(v9)+28)) = v24
		v29 = int32(-2) - v24
		if base.Ui32(v25) < base.Ui32(v29) {
			v31 = v25
		} else {
			v31 = v29
		}
		*(*int32)(unsafe.Add(mBase, uint32(v9)+56)) = v31
		v33 = v24 + v31
		*(*int32)(unsafe.Add(mBase, uint32(v9)+36)) = v33
		*(*int32)(unsafe.Add(mBase, uint32(v9)+24)) = v33
		v38 = F___small_vfprintf(m, v9+int32(8), l2, l3)
		mBase = m.M
		v41 = m.ExcPending
		if v41 != 0 {
			return int32(0)
		} else {
			if v24 == int32(-2) {
				v55 = v38
			} else {
				v44 = *(*int32)(unsafe.Add(mBase, uint32(v9)+28))
				v45 = *(*int32)(unsafe.Add(mBase, uint32(v9)+24))
				v48 = int32(0)
				*(*uint8)(unsafe.Add(mBase, uint32(v44-base.B2i32(v44 == v45)))) = uint8(v48)
				v55 = v38
			}
			m.G0 = v9 + int32(160)
			return v55
		}
	} else {
		if l1 != 0 {
			*(*int32)(unsafe.Add(mBase, _c_F___small_vsnprintf[0])) = int32(61)
			v55 = int32(-1)
			m.G0 = v9 + int32(160)
			return v55
		} else {
			v24 = v9 + int32(159)
			v25 = int32(1)
			*(*int32)(unsafe.Add(mBase, uint32(v9)+52)) = v24
			*(*int32)(unsafe.Add(mBase, uint32(v9)+28)) = v24
			v29 = int32(-2) - v24
			if base.Ui32(v25) < base.Ui32(v29) {
				v31 = v25
			} else {
				v31 = v29
			}
			*(*int32)(unsafe.Add(mBase, uint32(v9)+56)) = v31
			v33 = v24 + v31
			*(*int32)(unsafe.Add(mBase, uint32(v9)+36)) = v33
			*(*int32)(unsafe.Add(mBase, uint32(v9)+24)) = v33
			v38 = F___small_vfprintf(m, v9+int32(8), l2, l3)
			mBase = m.M
			v41 = m.ExcPending
			if v41 != 0 {
				return int32(0)
			} else {
				if v24 == int32(-2) {
					v55 = v38
				} else {
					v44 = *(*int32)(unsafe.Add(mBase, uint32(v9)+28))
					v45 = *(*int32)(unsafe.Add(mBase, uint32(v9)+24))
					v48 = int32(0)
					*(*uint8)(unsafe.Add(mBase, uint32(v44-base.B2i32(v44 == v45)))) = uint8(v48)
					v55 = v38
				}
				m.G0 = v9 + int32(160)
				return v55
			}
		}
	}
}
func F___stdio_read(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v14 int32
	_ = v14
	var v19 int32
	_ = v19
	var v23 int32
	_ = v23
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v35 int32
	_ = v35
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v52 int32
	_ = v52
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	v4 = int32(0)
	v8 = m.G0
	v9 = int32(32)
	v10 = v8 - v9
	m.G0 = v10
	*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = l1
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v10)+20)) = l2 - base.B2i32(v14 != v4)
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	*(*int32)(unsafe.Add(mBase, uint32(v10)+28)) = v14
	*(*int32)(unsafe.Add(mBase, uint32(v10)+24)) = v19
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	v29 = m.Wasi_snapshot_preview1.Fd_read(m, v23, v10+int32(16), int32(2), v10+int32(12))
	mBase = m.M
	if v29 != 0 {
		v31 = F___errno_location(m)
		mBase = m.M
		*(*int32)(unsafe.Add(mBase, uint32(v31))) = v29
	} else {
	}
	if v29 != 0 {
		v41 = v9
		v42 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		*(*int32)(unsafe.Add(mBase, uint32(l0))) = v42 | v41
		v63 = v4
	} else {
		v35 = *(*int32)(unsafe.Add(mBase, uint32(v10)+12))
		if int32(0) < v35 {
			v45 = *(*int32)(unsafe.Add(mBase, uint32(v10)+20))
			if base.Ui32(v35) <= base.Ui32(v45) {
				v63 = v35
			} else {
				v47 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
				*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v47
				*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v47 + (v35 - v45)
				v52 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
				if v52 == int32(0) {
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v47 + int32(1)
					v61 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v47))))
					*(*uint8)(unsafe.Add(mBase, uint32(l1+l2+int32(-1)))) = uint8(v61)
				}
				v63 = l2
			}
		} else {
			if v35 != 0 {
				v40 = int32(32)
			} else {
				v40 = int32(16)
			}
			v41 = v40
			v42 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			*(*int32)(unsafe.Add(mBase, uint32(l0))) = v42 | v41
			v63 = v4
		}
	}
	m.G0 = v10 + int32(32)
	return v63
}
func F___stdio_seek(m *base.Module, l0 int32, l1 int64, l2 int32) int64 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int64
	_ = v5
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	v5 = F___lseek(m, v4, l1, l2)
	mBase = m.M
	return v5
}
func F___stpncpy(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v103 int32
	_ = v103
	if (l1^l0)&int32(3) != 0 {
		v68 = l0
		v69 = l1
		v70 = l2
		goto L4
	} else {
		goto L5
	}
L1:
	;
	v103 = F__emscripten_memset_bulkmem(m, v97, base.I32_extend8_s(int32(0)), v99)
	mBase = m.M
	goto L26
L2:
	;
	v97 = v92
	v99 = int32(0)
	goto L1
L3:
	;
	v78 = v74
	v79 = v75
	v80 = v76
	goto L22
L4:
	;
	if v70 == int32(0) {
		v92 = v68
		goto L2
	} else {
		goto L21
	}
L5:
	;
	v8 = int32(0)
	v9 = base.B2i32(l2 != v8)
	if l1&int32(3) == v8 {
		v36 = l0
		v37 = l1
		v38 = l2
		v39 = v9
		goto L6
	} else {
		goto L7
	}
L6:
	;
	if v39 == int32(0) {
		v92 = v36
		goto L2
	} else {
		goto L14
	}
L7:
	;
	if l2 == int32(0) {
		v36 = l0
		v37 = l1
		v38 = l2
		v39 = v9
		goto L6
	} else {
		goto L8
	}
L8:
	;
	v16 = l0
	v17 = l1
	v18 = l2
	goto L9
L9:
	;
	v20 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17))))
	*(*uint8)(unsafe.Add(mBase, uint32(v16))) = uint8(v20)
	if v20 == int32(0) {
		v97 = v16
		v99 = v18
		goto L1
	} else {
		goto L11
	}
L10:
	;
	v36 = v25
	v37 = v31
	v38 = v27
	v39 = v29
	goto L6
L11:
	;
	v24 = int32(1)
	v25 = v16 + v24
	v27 = v18 + int32(-1)
	v28 = int32(0)
	v29 = base.B2i32(v27 != v28)
	v31 = v17 + v24
	if v31&int32(3) == v28 {
		v36 = v25
		v37 = v31
		v38 = v27
		v39 = v29
		goto L6
	} else {
		goto L12
	}
L12:
	;
	if v27 != 0 {
		v16 = v25
		v17 = v31
		v18 = v27
		goto L9
	} else {
		goto L13
	}
L13:
	;
	goto L10
L14:
	;
	v42 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v37))))
	if v42 == int32(0) {
		v97 = v36
		v99 = v38
		goto L1
	} else {
		goto L15
	}
L15:
	;
	if base.Ui32(v38) < base.Ui32(int32(4)) {
		v68 = v36
		v69 = v37
		v70 = v38
		goto L4
	} else {
		goto L16
	}
L16:
	;
	v47 = v36
	v48 = v37
	v49 = v38
	goto L17
L17:
	;
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v48)))
	v55 = int32(-2139062144)
	if (int32(16843008)-v52|v52)&v55 != v55 {
		v74 = v47
		v75 = v48
		v76 = v49
		goto L3
	} else {
		goto L19
	}
L18:
	;
	v68 = v61
	v69 = v63
	v70 = v65
	goto L4
L19:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v47))) = v52
	v60 = int32(4)
	v61 = v47 + v60
	v63 = v48 + v60
	v65 = v49 + int32(-4)
	if base.Ui32(int32(3)) < base.Ui32(v65) {
		v47 = v61
		v48 = v63
		v49 = v65
		goto L17
	} else {
		goto L20
	}
L20:
	;
	goto L18
L21:
	;
	v74 = v68
	v75 = v69
	v76 = v70
	goto L3
L22:
	;
	v82 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v79))))
	*(*uint8)(unsafe.Add(mBase, uint32(v78))) = uint8(v82)
	if v82 == int32(0) {
		v97 = v78
		v99 = v80
		goto L1
	} else {
		goto L24
	}
L23:
	;
	v92 = v87
	goto L2
L24:
	;
	v86 = int32(1)
	v87 = v78 + v86
	v91 = v80 + int32(-1)
	if v91 != 0 {
		v78 = v87
		v79 = v79 + v86
		v80 = v91
		goto L22
	} else {
		goto L25
	}
L25:
	;
	goto L23
L26:
	;
	return v97
}
func F___strchrnul(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v16 int32
	_ = v16
	var v21 int32
	_ = v21
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v63 int32
	_ = v63
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v81 int32
	_ = v81
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v92 int32
	_ = v92
	var v98 int32
	_ = v98
	var v101 int32
	_ = v101
	var v107 int32
	_ = v107
	var v111 int32
	_ = v111
	var v113 int32
	_ = v113
	var v121 int32
	_ = v121
	var v124 int32
	_ = v124
	var v133 int32
	_ = v133
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v147 int32
	_ = v147
	v7 = l1 & int32(255)
	if v7 == int32(0) {
		goto L4
	} else {
		goto L5
	}
L1:
	;
	return v147
L2:
	;
	v138 = v133
	goto L34
L3:
	;
	v133 = v124
	goto L2
L4:
	;
	if l0&int32(3) == int32(0) {
		v88 = l0
		goto L20
	} else {
		goto L21
	}
L5:
	;
	if l0&int32(3) == int32(0) {
		v29 = l0
		goto L6
	} else {
		goto L7
	}
L6:
	;
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v29)))
	v38 = int32(-2139062144)
	if (int32(16843008)-v35|v35)&v38 != v38 {
		v124 = v29
		goto L3
	} else {
		goto L13
	}
L7:
	;
	v16 = l0
	goto L8
L8:
	;
	v21 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16))))
	if v21 == int32(0) {
		v147 = v16
		goto L1
	} else {
		goto L10
	}
L9:
	;
	v29 = v26
	goto L6
L10:
	;
	if v21 == l1&int32(255) {
		v147 = v16
		goto L1
	} else {
		goto L11
	}
L11:
	;
	v26 = v16 + int32(1)
	if v26&int32(3) != 0 {
		v16 = v26
		goto L8
	} else {
		goto L12
	}
L12:
	;
	goto L9
L13:
	;
	v44 = v29
	v47 = v35
	goto L14
L14:
	;
	v50 = v47 ^ v7*int32(16843009)
	v53 = int32(-2139062144)
	if (int32(16843008)-v50|v50)&v53 != v53 {
		v124 = v44
		goto L3
	} else {
		goto L16
	}
L16:
	;
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v44)+4))
	v59 = v44 + int32(4)
	v63 = int32(-2139062144)
	if (v57|(int32(16843008)-v57))&v63 == v63 {
		v44 = v59
		v47 = v57
		goto L14
	} else {
		goto L17
	}
L17:
	;
	v133 = v59
	goto L2
L18:
	;
	return l0 + v121
L19:
	;
	v121 = v113 - l0
	goto L18
L20:
	;
	v92 = v88
	goto L28
L21:
	;
	v74 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v74 != 0 {
		goto L22
	} else {
		goto L23
	}
L22:
	;
	v77 = l0
	goto L24
L23:
	;
	v121 = l0 - l0
	goto L18
L24:
	;
	v81 = v77 + int32(1)
	if v81&int32(3) == int32(0) {
		v88 = v81
		goto L20
	} else {
		goto L26
	}
L26:
	;
	v86 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v81))))
	if v86 != 0 {
		v77 = v81
		goto L24
	} else {
		goto L27
	}
L27:
	;
	v113 = v81
	goto L19
L28:
	;
	v98 = *(*int32)(unsafe.Add(mBase, uint32(v92)))
	v101 = int32(-2139062144)
	if (int32(16843008)-v98|v98)&v101 == v101 {
		v92 = v92 + int32(4)
		goto L28
	} else {
		goto L30
	}
L29:
	;
	v107 = v92
	goto L31
L30:
	;
	goto L29
L31:
	;
	v111 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v107))))
	if v111 != 0 {
		v107 = v107 + int32(1)
		goto L31
	} else {
		goto L33
	}
L32:
	;
	v113 = v107
	goto L19
L33:
	;
	goto L32
L34:
	;
	v139 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v138))))
	if v139 == int32(0) {
		v147 = v138
		goto L1
	} else {
		goto L36
	}
L35:
	;
	v147 = v138
	goto L1
L36:
	;
	if v139 != l1&int32(255) {
		v138 = v138 + int32(1)
		goto L34
	} else {
		goto L37
	}
L37:
	;
	goto L35
}
func F___strcoll_l(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	v6 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	v7 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v7 == int32(0) {
		v30 = v6
		v31 = v7
		goto L2
	} else {
		goto L3
	}
L1:
	;
	return v31 - v30&int32(255)
L2:
	;
	goto L1
L3:
	;
	if v7 != v6&int32(255) {
		v30 = v6
		v31 = v7
		goto L2
	} else {
		goto L4
	}
L4:
	;
	v13 = l0
	v14 = l1
	goto L5
L5:
	;
	v17 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+1)))
	v18 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+1)))
	if v18 == int32(0) {
		v30 = v17
		v31 = v18
		goto L2
	} else {
		goto L7
	}
L6:
	;
	v30 = v17
	v31 = v18
	goto L2
L7:
	;
	v21 = int32(1)
	if v18 == v17&int32(255) {
		v13 = v13 + v21
		v14 = v14 + v21
		goto L5
	} else {
		goto L8
	}
L8:
	;
	goto L6
}
func F___strerror_l(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v11 int32
	_ = v11
	if base.Ui32(int32(153)) < base.Ui32(l0) {
		v6 = int32(0)
	} else {
		v6 = l0
	}
	v11 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v6<<(uint(int32(1))%32))+uint32(_c_F___strerror_l[0]))))
	return v11 + int32(_a_F___strerror_l_0)
}
func F_sbrk(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v5 int32
	_ = v5
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v18 int32
	_ = v18
	v2 = int32(0)
	v5 = *(*int32)(unsafe.Add(mBase, _c_F_sbrk[0]))
	v9 = (l0 + int32(7)) & int32(-8)
	v10 = v5 + v9
	if v9 == v2 {
		if base.Ui32(v10) <= base.Ui32(base.MemorySize(m)<<(uint(int32(16))%32)) {
			*(*int32)(unsafe.Add(mBase, _c_F_sbrk[0])) = v10
			return v5
		} else {
			v18 = m.Env.Emscripten_resize_heap(m, v10)
			mBase = m.M
			if v18 != 0 {
				*(*int32)(unsafe.Add(mBase, _c_F_sbrk[0])) = v10
				return v5
			} else {
				*(*int32)(unsafe.Add(mBase, _c_F_sbrk[1])) = int32(48)
				return int32(-1)
			}
		}
	} else {
		if base.Ui32(v10) <= base.Ui32(v5) {
			*(*int32)(unsafe.Add(mBase, _c_F_sbrk[1])) = int32(48)
			return int32(-1)
		} else {
			if base.Ui32(v10) <= base.Ui32(base.MemorySize(m)<<(uint(int32(16))%32)) {
				*(*int32)(unsafe.Add(mBase, _c_F_sbrk[0])) = v10
				return v5
			} else {
				v18 = m.Env.Emscripten_resize_heap(m, v10)
				mBase = m.M
				if v18 != 0 {
					*(*int32)(unsafe.Add(mBase, _c_F_sbrk[0])) = v10
					return v5
				} else {
					*(*int32)(unsafe.Add(mBase, _c_F_sbrk[1])) = int32(48)
					return int32(-1)
				}
			}
		}
	}
}
func F_scalbnl(m *base.Module, l0 int32, l1 int64, l2 int64, l3 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v21 int64
	_ = v21
	var v22 int64
	_ = v22
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v42 int64
	_ = v42
	var v43 int64
	_ = v43
	var v55 int64
	_ = v55
	var v56 int64
	_ = v56
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v76 int64
	_ = v76
	var v77 int64
	_ = v77
	var v78 int64
	_ = v78
	var v79 int64
	_ = v79
	var v80 int32
	_ = v80
	var v90 int64
	_ = v90
	var v92 int64
	_ = v92
	v6 = m.G0
	v8 = v6 - int32(80)
	m.G0 = v8
	if l3 < int32(16384) {
		if int32(-16383) < l3 {
			v78 = l1
			v79 = l2
			v80 = l3
		} else {
			F___multf3(m, v8+int32(64), l1, l2, int64(0), int64(32088147345014784))
			mBase = m.M
			v55 = *(*int64)(unsafe.Add(mBase, uint32(v8+int32(72))))
			v56 = *(*int64)(unsafe.Add(mBase, uint32(v8)+64))
			if base.Ui32(l3) <= base.Ui32(int32(-32652)) {
				F___multf3(m, v8+int32(48), v56, v55, int64(0), int64(32088147345014784))
				mBase = m.M
				v66 = int32(-48920)
				if base.Ui32(v66) < base.Ui32(l3) {
					v69 = l3
				} else {
					v69 = v66
				}
				v76 = *(*int64)(unsafe.Add(mBase, uint32(v8+int32(56))))
				v77 = *(*int64)(unsafe.Add(mBase, uint32(v8)+48))
				v78 = v77
				v79 = v76
				v80 = v69 + int32(32538)
			} else {
				v78 = v56
				v79 = v55
				v80 = l3 + int32(16269)
			}
		}
	} else {
		F___multf3(m, v8+int32(32), l1, l2, int64(0), int64(9222809086901354496))
		mBase = m.M
		v21 = *(*int64)(unsafe.Add(mBase, uint32(v8+int32(40))))
		v22 = *(*int64)(unsafe.Add(mBase, uint32(v8)+32))
		if base.Ui32(int32(32767)) <= base.Ui32(l3) {
			F___multf3(m, v8+int32(16), v22, v21, int64(0), int64(9222809086901354496))
			mBase = m.M
			v32 = int32(49149)
			if base.Ui32(l3) < base.Ui32(v32) {
				v35 = l3
			} else {
				v35 = v32
			}
			v42 = *(*int64)(unsafe.Add(mBase, uint32(v8+int32(24))))
			v43 = *(*int64)(unsafe.Add(mBase, uint32(v8)+16))
			v78 = v43
			v79 = v42
			v80 = v35 + int32(-32766)
		} else {
			v78 = v22
			v79 = v21
			v80 = l3 + int32(-16383)
		}
	}
	F___multf3(m, v8, v78, v79, int64(0), base.I64_extend_i32_u(v80+int32(16383))<<(uint(int64(48))%64))
	mBase = m.M
	v90 = *(*int64)(unsafe.Add(mBase, uint32(v8+int32(8))))
	*(*int64)(unsafe.Add(mBase, uint32(l0)+8)) = v90
	v92 = *(*int64)(unsafe.Add(mBase, uint32(v8)))
	*(*int64)(unsafe.Add(mBase, uint32(l0))) = v92
	m.G0 = v8 + int32(80)
	return
}
func F_scanDatabaseForReadyKeys(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v33 int32
	_ = v33
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v46 int64
	_ = v46
	var v47 int64
	_ = v47
	var v48 int64
	_ = v48
	var v49 int64
	_ = v49
	var v50 int64
	_ = v50
	var v51 int64
	_ = v51
	var v52 int64
	_ = v52
	var v54 int64
	_ = v54
	var v56 int64
	_ = v56
	var v58 int64
	_ = v58
	var v60 int64
	_ = v60
	var v62 int64
	_ = v62
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v84 int32
	_ = v84
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v104 int32
	_ = v104
	var v109 int32
	_ = v109
	var v113 int32
	_ = v113
	var v117 int32
	_ = v117
	var v126 int32
	_ = v126
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v133 int32
	_ = v133
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v141 int32
	_ = v141
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v149 int32
	_ = v149
	var v153 int32
	_ = v153
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v165 int32
	_ = v165
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v173 int32
	_ = v173
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v182 int32
	_ = v182
	var v184 int32
	_ = v184
	var v186 int64
	_ = v186
	var v187 int64
	_ = v187
	var v188 int64
	_ = v188
	var v189 int64
	_ = v189
	var v190 int64
	_ = v190
	var v191 int64
	_ = v191
	var v192 int64
	_ = v192
	var v194 int64
	_ = v194
	var v196 int64
	_ = v196
	var v198 int64
	_ = v198
	var v200 int64
	_ = v200
	var v202 int64
	_ = v202
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v209 int32
	_ = v209
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v216 int32
	_ = v216
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v224 int32
	_ = v224
	var v228 int32
	_ = v228
	var v230 int32
	_ = v230
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v244 int32
	_ = v244
	var v249 int32
	_ = v249
	var v253 int32
	_ = v253
	var v257 int32
	_ = v257
	var v268 int32
	_ = v268
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v12 = F_dictGetSafeIterator(m, v11)
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	F_dictReleaseIterator(m, v12)
	mBase = m.M
	v268 = m.ExcPending
	if v268 != 0 {
		goto L2
	} else {
		goto L68
	}
L2:
	;
	return
L3:
	;
	v21 = v12 + int32(20)
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v12)+16))
	if v22 != 0 {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	if v117 == int32(0) {
		goto L1
	} else {
		goto L30
	}
L5:
	;
	v28 = v21
	v29 = v25
	goto L8
L6:
	;
	v25 = int32(1)
	goto L5
L7:
	;
	v25 = int32(0)
	goto L5
L8:
	;
	switch v29 {
	case 0:
		goto L13
	default:
		goto L12
	}
L10:
	;
	v29 = int32(0)
	goto L8
L11:
	;
	goto L4
L12:
	;
	v109 = *(*int32)(unsafe.Add(mBase, uint32(v28)))
	*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = v109
	if v109 == int32(0) {
		goto L10
	} else {
		goto L29
	}
L13:
	;
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v12)+4))
	if v33 != int32(-1) {
		v72 = v33
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v73 = int32(1)
	v74 = v72 + v73
	*(*int32)(unsafe.Add(mBase, uint32(v12)+4)) = v74
	v76 = int32(0)
	v79 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
	v80 = *(*int32)(unsafe.Add(mBase, uint32(v12)+8))
	v84 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v79+v80+int32(26)))))
	if v84 == int32(255) {
		goto L23
	} else {
		goto L24
	}
L15:
	;
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v12)+8))
	if v37 != 0 {
		v72 = int32(-1)
		goto L14
	} else {
		goto L16
	}
L16:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v12)+12))
	if v39 == int32(0) {
		goto L18
	} else {
		goto L19
	}
L17:
	;
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v65)+20))
	if v66 != int32(-1) {
		goto L20
	} else {
		goto L21
	}
L18:
	;
	v46 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v38)+16)))
	v47 = int64(*(*int8)(unsafe.Add(mBase, uint32(v38)+27)))
	v48 = int64(*(*int32)(unsafe.Add(mBase, uint32(v38)+8)))
	v49 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v38)+12)))
	v50 = int64(*(*int8)(unsafe.Add(mBase, uint32(v38)+26)))
	v51 = int64(*(*int32)(unsafe.Add(mBase, uint32(v38)+4)))
	v52 = F_wangHash64(m, v51)
	mBase = m.M
	v54 = F_wangHash64(m, v50+v52)
	mBase = m.M
	v56 = F_wangHash64(m, v49+v54)
	mBase = m.M
	v58 = F_wangHash64(m, v48+v56)
	mBase = m.M
	v60 = F_wangHash64(m, v47+v58)
	mBase = m.M
	v62 = F_wangHash64(m, v46+v60)
	mBase = m.M
	*(*int64)(unsafe.Add(mBase, uint32(v12)+24)) = v62
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
	v65 = v64
	goto L17
L19:
	;
	v42 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v38)+24)))
	v44 = v42 + int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v38)+24)) = uint16(v44)
	v65 = v38
	goto L17
L20:
	;
	v72 = v66 + int32(-1)
	goto L14
L21:
	;
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v12)+4))
	v72 = v69
	goto L14
L22:
	;
	v99 = int32(2)
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v79+v97<<(uint(v99)%32)+int32(4))))
	v28 = v104 + v98<<(uint(v99)%32)
	v29 = int32(1)
	goto L8
L23:
	;
	v88 = v76
	goto L25
L24:
	;
	v88 = v73 << (uint(v84) % 32)
	goto L25
L25:
	;
	if v74 < v88 {
		v97 = v80
		v98 = v74
		goto L22
	} else {
		goto L26
	}
L26:
	;
	if v80 != 0 {
		v117 = v76
		goto L11
	} else {
		goto L27
	}
L27:
	;
	v90 = *(*int32)(unsafe.Add(mBase, uint32(v79)+20))
	if v90 == int32(-1) {
		v117 = v76
		goto L11
	} else {
		goto L28
	}
L28:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v12)+4)) = int64(4294967296)
	v97 = int32(1)
	v98 = int32(0)
	goto L22
L29:
	;
	v113 = *(*int32)(unsafe.Add(mBase, uint32(v109)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v21))) = v113
	v117 = v109
	goto L11
L30:
	;
	v126 = v117
	goto L31
L31:
	;
	v129 = *(*int32)(unsafe.Add(mBase, uint32(v126)))
	goto L33
L32:
	;
	goto L1
L33:
	;
	v130 = F_objectGetVal(m, v129)
	mBase = m.M
	v131 = int32(0)
	v133 = *(*int32)(unsafe.Add(mBase, _c_F_scanDatabaseForReadyKeys[0]))
	if v133 == v131 {
		v138 = v131
		goto L34
	} else {
		goto L35
	}
L34:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9)+12)) = int32(0)
	v141 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v144 = F_kvstoreHashtableFind(m, v141, v138, v130, v9+int32(12))
	mBase = m.M
	v145 = m.ExcPending
	if v145 != 0 {
		goto L2
	} else {
		goto L37
	}
L35:
	;
	v136 = F_getKeySlot(m, v130)
	mBase = m.M
	v137 = m.ExcPending
	if v137 != 0 {
		goto L2
	} else {
		goto L36
	}
L36:
	;
	v138 = v136
	goto L34
L37:
	;
	v146 = *(*int32)(unsafe.Add(mBase, uint32(v9)+12))
	if v146 == int32(0) {
		goto L38
	} else {
		goto L39
	}
L38:
	;
	v161 = v12 + int32(20)
	v162 = *(*int32)(unsafe.Add(mBase, uint32(v12)+16))
	if v162 != 0 {
		goto L43
	} else {
		goto L44
	}
L39:
	;
	v149 = *(*int32)(unsafe.Add(mBase, uint32(v146)))
	F_signalKeyAsReady(m, l0, v129, v149&int32(15))
	mBase = m.M
	v153 = m.ExcPending
	if v153 != 0 {
		goto L2
	} else {
		goto L40
	}
L40:
	;
	goto L38
L41:
	;
	if v257 != 0 {
		v126 = v257
		goto L31
	} else {
		goto L67
	}
L42:
	;
	v168 = v161
	v169 = v165
	goto L45
L43:
	;
	v165 = int32(1)
	goto L42
L44:
	;
	v165 = int32(0)
	goto L42
L45:
	;
	switch v169 {
	case 0:
		goto L50
	default:
		goto L49
	}
L47:
	;
	v169 = int32(0)
	goto L45
L48:
	;
	goto L41
L49:
	;
	v249 = *(*int32)(unsafe.Add(mBase, uint32(v168)))
	*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = v249
	if v249 == int32(0) {
		goto L47
	} else {
		goto L66
	}
L50:
	;
	v173 = *(*int32)(unsafe.Add(mBase, uint32(v12)+4))
	if v173 != int32(-1) {
		v212 = v173
		goto L51
	} else {
		goto L52
	}
L51:
	;
	v213 = int32(1)
	v214 = v212 + v213
	*(*int32)(unsafe.Add(mBase, uint32(v12)+4)) = v214
	v216 = int32(0)
	v219 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
	v220 = *(*int32)(unsafe.Add(mBase, uint32(v12)+8))
	v224 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v219+v220+int32(26)))))
	if v224 == int32(255) {
		goto L60
	} else {
		goto L61
	}
L52:
	;
	v177 = *(*int32)(unsafe.Add(mBase, uint32(v12)+8))
	if v177 != 0 {
		v212 = int32(-1)
		goto L51
	} else {
		goto L53
	}
L53:
	;
	v178 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
	v179 = *(*int32)(unsafe.Add(mBase, uint32(v12)+12))
	if v179 == int32(0) {
		goto L55
	} else {
		goto L56
	}
L54:
	;
	v206 = *(*int32)(unsafe.Add(mBase, uint32(v205)+20))
	if v206 != int32(-1) {
		goto L57
	} else {
		goto L58
	}
L55:
	;
	v186 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v178)+16)))
	v187 = int64(*(*int8)(unsafe.Add(mBase, uint32(v178)+27)))
	v188 = int64(*(*int32)(unsafe.Add(mBase, uint32(v178)+8)))
	v189 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v178)+12)))
	v190 = int64(*(*int8)(unsafe.Add(mBase, uint32(v178)+26)))
	v191 = int64(*(*int32)(unsafe.Add(mBase, uint32(v178)+4)))
	v192 = F_wangHash64(m, v191)
	mBase = m.M
	v194 = F_wangHash64(m, v190+v192)
	mBase = m.M
	v196 = F_wangHash64(m, v189+v194)
	mBase = m.M
	v198 = F_wangHash64(m, v188+v196)
	mBase = m.M
	v200 = F_wangHash64(m, v187+v198)
	mBase = m.M
	v202 = F_wangHash64(m, v186+v200)
	mBase = m.M
	*(*int64)(unsafe.Add(mBase, uint32(v12)+24)) = v202
	v204 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
	v205 = v204
	goto L54
L56:
	;
	v182 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v178)+24)))
	v184 = v182 + int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v178)+24)) = uint16(v184)
	v205 = v178
	goto L54
L57:
	;
	v212 = v206 + int32(-1)
	goto L51
L58:
	;
	v209 = *(*int32)(unsafe.Add(mBase, uint32(v12)+4))
	v212 = v209
	goto L51
L59:
	;
	v239 = int32(2)
	v244 = *(*int32)(unsafe.Add(mBase, uint32(v219+v237<<(uint(v239)%32)+int32(4))))
	v168 = v244 + v238<<(uint(v239)%32)
	v169 = int32(1)
	goto L45
L60:
	;
	v228 = v216
	goto L62
L61:
	;
	v228 = v213 << (uint(v224) % 32)
	goto L62
L62:
	;
	if v214 < v228 {
		v237 = v220
		v238 = v214
		goto L59
	} else {
		goto L63
	}
L63:
	;
	if v220 != 0 {
		v257 = v216
		goto L48
	} else {
		goto L64
	}
L64:
	;
	v230 = *(*int32)(unsafe.Add(mBase, uint32(v219)+20))
	if v230 == int32(-1) {
		v257 = v216
		goto L48
	} else {
		goto L65
	}
L65:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v12)+4)) = int64(4294967296)
	v237 = int32(1)
	v238 = int32(0)
	goto L59
L66:
	;
	v253 = *(*int32)(unsafe.Add(mBase, uint32(v249)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v161))) = v253
	v257 = v249
	goto L48
L67:
	;
	goto L32
L68:
	;
	m.G0 = v9 + int32(16)
	return
}
func F_scardCommand(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v4 = *(*int32)(unsafe.Add(mBase, uint32(v3)+4))
	v6 = *(*int32)(unsafe.Add(mBase, _c_F_scardCommand[0]))
	v7 = F_lookupKeyReadOrReply(m, l0, v4, v6)
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		return
	} else {
		if v7 == int32(0) {
			return
		} else {
			v12 = F_checkType(m, l0, v7, int32(2))
			mBase = m.M
			v13 = m.ExcPending
			if v13 != 0 {
				return
			} else {
				if v12 != 0 {
					return
				} else {
					v14 = F_setTypeSize(m, v7)
					mBase = m.M
					v15 = m.ExcPending
					if v15 != 0 {
						return
					} else {
						F_addReplyLongLong(m, l0, base.I64_extend_i32_u(v14))
						mBase = m.M
						v18 = m.ExcPending
						if v18 != 0 {
							return
						} else {
							return
						}
					}
				}
			}
		}
	}
}
func F_sched_yield(m *base.Module) int32 {
	var v1 float64
	_ = v1
	var v5 int32
	_ = v5
	v1 = m.Env.Emscripten_get_now(m)
	F__emscripten_yield(m, v1)
	v5 = m.ExcPending
	if v5 != 0 {
		return int32(0)
	} else {
		return int32(0)
	}
}
func F_sdiffstoreCommand(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v9 = *(*int32)(unsafe.Add(mBase, uint32(v3)+4))
	F_sunionDiffGenericCommand(m, l0, v3+int32(8), v6+int32(-2), v9, int32(1))
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		return
	} else {
		return
	}
}
func F_sdigits10(m *base.Module, l0 int64) int32 {
	var v6 int32
	_ = v6
	var v8 int64
	_ = v8
	var v11 int64
	_ = v11
	var v12 int32
	_ = v12
	var v33 int32
	_ = v33
	var v38 int32
	_ = v38
	var v45 int32
	_ = v45
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v56 int64
	_ = v56
	var v58 int32
	_ = v58
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v68 int32
	_ = v68
	var v71 int64
	_ = v71
	var v72 int32
	_ = v72
	var v99 int32
	_ = v99
	var v106 int32
	_ = v106
	var v115 int32
	_ = v115
	var v122 int32
	_ = v122
	var v126 int32
	_ = v126
	var v130 int64
	_ = v130
	var v132 int32
	_ = v132
	if int64(-1) < l0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v68 = int32(0)
	if base.Ui64(l0) < base.Ui64(int64(10)) {
		v132 = v68
		goto L35
	} else {
		goto L36
	}
L2:
	;
	v6 = int32(0)
	v8 = int64(0) - l0
	if base.Ui64(v8) < base.Ui64(int64(10)) {
		v58 = v6
		goto L4
	} else {
		goto L5
	}
L3:
	;
	return v62 + v63 + int32(1)
L4:
	;
	v62 = v58
	v63 = int32(1)
	goto L3
L5:
	;
	v11 = v8
	v12 = v6
	goto L6
L6:
	;
	if base.Ui64(int64(100)) <= base.Ui64(v11) {
		goto L8
	} else {
		goto L9
	}
L7:
	;
	v58 = v52
	goto L4
L8:
	;
	if base.Ui64(int64(1000)) <= base.Ui64(v11) {
		goto L10
	} else {
		goto L11
	}
L9:
	;
	v62 = v12
	v63 = int32(2)
	goto L3
L10:
	;
	if base.Ui64(int64(999999999999)) < base.Ui64(v11) {
		goto L12
	} else {
		goto L13
	}
L11:
	;
	v62 = v12
	v63 = int32(3)
	goto L3
L12:
	;
	v52 = v12 + int32(12)
	v56 = base.I64_div_u_s(v11, int64(1000000000000))
	if base.Ui64(int64(9999999999999)) < base.Ui64(v11) {
		v11 = v56
		v12 = v52
		goto L6
	} else {
		goto L34
	}
L13:
	;
	if base.Ui64(int64(99999999)) < base.Ui64(v11) {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	if base.Ui64(int64(9999999999)) < base.Ui64(v11) {
		goto L26
	} else {
		goto L27
	}
L15:
	;
	if base.Ui64(int64(999999)) < base.Ui64(v11) {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	if base.Ui64(int64(9999999)) < base.Ui64(v11) {
		goto L23
	} else {
		goto L24
	}
L17:
	;
	if base.Ui64(int64(10000)) <= base.Ui64(v11) {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	if base.Ui64(int64(99999)) < base.Ui64(v11) {
		goto L20
	} else {
		goto L21
	}
L19:
	;
	v62 = v12
	v63 = int32(4)
	goto L3
L20:
	;
	v33 = int32(6)
	goto L22
L21:
	;
	v33 = int32(5)
	goto L22
L22:
	;
	v62 = v12
	v63 = v33
	goto L3
L23:
	;
	v38 = int32(8)
	goto L25
L24:
	;
	v38 = int32(7)
	goto L25
L25:
	;
	v62 = v12
	v63 = v38
	goto L3
L26:
	;
	if base.Ui64(int64(99999999999)) < base.Ui64(v11) {
		goto L31
	} else {
		goto L32
	}
L27:
	;
	if base.Ui64(int64(999999999)) < base.Ui64(v11) {
		goto L28
	} else {
		goto L29
	}
L28:
	;
	v45 = int32(10)
	goto L30
L29:
	;
	v45 = int32(9)
	goto L30
L30:
	;
	v62 = v12
	v63 = v45
	goto L3
L31:
	;
	v50 = int32(12)
	goto L33
L32:
	;
	v50 = int32(11)
	goto L33
L33:
	;
	v62 = v12
	v63 = v50
	goto L3
L34:
	;
	goto L7
L35:
	;
	return int32(1) + v132
L36:
	;
	v71 = l0
	v72 = v68
	goto L37
L37:
	;
	if base.Ui64(int64(100)) <= base.Ui64(v71) {
		goto L39
	} else {
		goto L40
	}
L38:
	;
	v132 = v126
	goto L35
L39:
	;
	if base.Ui64(int64(1000)) <= base.Ui64(v71) {
		goto L41
	} else {
		goto L42
	}
L40:
	;
	return int32(2) + v72
L41:
	;
	if base.Ui64(int64(999999999999)) < base.Ui64(v71) {
		goto L43
	} else {
		goto L44
	}
L42:
	;
	return int32(3) + v72
L43:
	;
	v126 = v72 + int32(12)
	v130 = base.I64_div_u_s(v71, int64(1000000000000))
	if base.Ui64(int64(9999999999999)) < base.Ui64(v71) {
		v71 = v130
		v72 = v126
		goto L37
	} else {
		goto L65
	}
L44:
	;
	if base.Ui64(int64(99999999)) < base.Ui64(v71) {
		goto L45
	} else {
		goto L46
	}
L45:
	;
	if base.Ui64(int64(9999999999)) < base.Ui64(v71) {
		goto L57
	} else {
		goto L58
	}
L46:
	;
	if base.Ui64(int64(999999)) < base.Ui64(v71) {
		goto L47
	} else {
		goto L48
	}
L47:
	;
	if base.Ui64(int64(9999999)) < base.Ui64(v71) {
		goto L54
	} else {
		goto L55
	}
L48:
	;
	if base.Ui64(int64(10000)) <= base.Ui64(v71) {
		goto L49
	} else {
		goto L50
	}
L49:
	;
	if base.Ui64(int64(99999)) < base.Ui64(v71) {
		goto L51
	} else {
		goto L52
	}
L50:
	;
	return int32(4) + v72
L51:
	;
	v99 = int32(6)
	goto L53
L52:
	;
	v99 = int32(5)
	goto L53
L53:
	;
	return v99 + v72
L54:
	;
	v106 = int32(8)
	goto L56
L55:
	;
	v106 = int32(7)
	goto L56
L56:
	;
	return v106 + v72
L57:
	;
	if base.Ui64(int64(99999999999)) < base.Ui64(v71) {
		goto L62
	} else {
		goto L63
	}
L58:
	;
	if base.Ui64(int64(999999999)) < base.Ui64(v71) {
		goto L59
	} else {
		goto L60
	}
L59:
	;
	v115 = int32(10)
	goto L61
L60:
	;
	v115 = int32(9)
	goto L61
L61:
	;
	return v115 + v72
L62:
	;
	v122 = int32(12)
	goto L64
L63:
	;
	v122 = int32(11)
	goto L64
L64:
	;
	return v122 + v72
L65:
	;
	goto L38
}
func F_sdsavail_1(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v37 int64
	_ = v37
	var v40 int64
	_ = v40
	var v43 int32
	_ = v43
	v4 = int32(-1)
	v6 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0+v4))))
	switch v6&int32(7) + v4 {
	case 0:
		v13 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0+int32(-2)))))
		v16 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0+int32(-3)))))
		return v13 - v16
	case 1:
		v21 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0+int32(-3)))))
		v24 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0+int32(-5)))))
		return v21 - v24
	case 2:
		v29 = *(*int32)(unsafe.Add(mBase, uint32(l0+int32(-5))))
		v32 = *(*int32)(unsafe.Add(mBase, uint32(l0+int32(-9))))
		return v29 - v32
	case 3:
		v37 = *(*int64)(unsafe.Add(mBase, uint32(l0+int32(-9))))
		v40 = *(*int64)(unsafe.Add(mBase, uint32(l0+int32(-17))))
		v43 = base.I32_wrap_i64(v37 - v40)
		return v43
	default:
		v43 = int32(0)
		return v43
	}
}
func F_sdsavail_2(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v37 int64
	_ = v37
	var v40 int64
	_ = v40
	var v43 int32
	_ = v43
	v4 = int32(-1)
	v6 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0+v4))))
	switch v6&int32(7) + v4 {
	case 0:
		v13 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0+int32(-2)))))
		v16 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0+int32(-3)))))
		return v13 - v16
	case 1:
		v21 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0+int32(-3)))))
		v24 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0+int32(-5)))))
		return v21 - v24
	case 2:
		v29 = *(*int32)(unsafe.Add(mBase, uint32(l0+int32(-5))))
		v32 = *(*int32)(unsafe.Add(mBase, uint32(l0+int32(-9))))
		return v29 - v32
	case 3:
		v37 = *(*int64)(unsafe.Add(mBase, uint32(l0+int32(-9))))
		v40 = *(*int64)(unsafe.Add(mBase, uint32(l0+int32(-17))))
		v43 = base.I32_wrap_i64(v37 - v40)
		return v43
	default:
		v43 = int32(0)
		return v43
	}
}
func F_sdscat(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v20 int32
	_ = v20
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v31 int32
	_ = v31
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v46 int32
	_ = v46
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v60 int32
	_ = v60
	var v64 int32
	_ = v64
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v86 int32
	_ = v86
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v101 int32
	_ = v101
	var v117 int32
	_ = v117
	if l1&int32(3) == int32(0) {
		v27 = l1
		goto L3
	} else {
		goto L4
	}
L1:
	;
	v64 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0+int32(-1)))))
	switch v64 & int32(7) {
	case 0:
		goto L22
	case 1:
		goto L21
	case 2:
		goto L20
	case 3:
		goto L19
	case 4:
		goto L18
	default:
		v81 = int32(0)
		goto L17
	}
L2:
	;
	v60 = v52 - l1
	goto L1
L3:
	;
	v31 = v27
	goto L11
L4:
	;
	v13 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	if v13 != 0 {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	v16 = l1
	goto L7
L6:
	;
	v60 = l1 - l1
	goto L1
L7:
	;
	v20 = v16 + int32(1)
	if v20&int32(3) == int32(0) {
		v27 = v20
		goto L3
	} else {
		goto L9
	}
L9:
	;
	v25 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20))))
	if v25 != 0 {
		v16 = v20
		goto L7
	} else {
		goto L10
	}
L10:
	;
	v52 = v20
	goto L2
L11:
	;
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v31)))
	v40 = int32(-2139062144)
	if (int32(16843008)-v37|v37)&v40 == v40 {
		v31 = v31 + int32(4)
		goto L11
	} else {
		goto L13
	}
L12:
	;
	v46 = v31
	goto L14
L13:
	;
	goto L12
L14:
	;
	v50 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v46))))
	if v50 != 0 {
		v46 = v46 + int32(1)
		goto L14
	} else {
		goto L16
	}
L15:
	;
	v52 = v46
	goto L2
L16:
	;
	goto L15
L17:
	;
	v83 = F__sdsMakeRoomFor(m, l0, v60, int32(1))
	mBase = m.M
	v86 = m.ExcPending
	if v86 != 0 {
		goto L24
	} else {
		goto L25
	}
L18:
	;
	v80 = *(*int32)(unsafe.Add(mBase, uint32(l0+int32(-17))))
	v81 = v80
	goto L17
L19:
	;
	v77 = *(*int32)(unsafe.Add(mBase, uint32(l0+int32(-9))))
	v81 = v77
	goto L17
L20:
	;
	v74 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0+int32(-5)))))
	v81 = v74
	goto L17
L21:
	;
	v71 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0+int32(-3)))))
	v81 = v71
	goto L17
L22:
	;
	v81 = int32(base.Ui32(v64) >> (uint(int32(3)) % 32))
	goto L17
L23:
	;
	return v83
L24:
	;
	return int32(0)
L25:
	;
	if v83 == int32(0) {
		goto L23
	} else {
		goto L26
	}
L26:
	;
	if v60 == int32(0) {
		goto L28
	} else {
		goto L29
	}
L27:
	;
	v94 = v81 + v60
	v96 = v83 + int32(-1)
	v97 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v96))))
	switch v97 & int32(7) {
	case 0:
		goto L35
	case 1:
		goto L34
	case 2:
		goto L33
	case 3:
		goto L32
	case 4:
		goto L31
	default:
		goto L30
	}
L28:
	;
	goto L27
L29:
	;
	v92 = F__emscripten_memcpy_bulkmem(m, v83+v81, l1, v60)
	mBase = m.M
	goto L28
L30:
	;
	v117 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v83+v94))) = uint8(v117)
	goto L23
L31:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v83+int32(-17)))) = base.I64_extend_i32_u(v94)
	goto L30
L32:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v83+int32(-9)))) = v94
	goto L30
L33:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v83+int32(-5)))) = uint16(v94)
	goto L30
L34:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v83+int32(-3)))) = uint8(v94)
	goto L30
L35:
	;
	v101 = v94 << (uint(int32(3)) % 32)
	*(*uint8)(unsafe.Add(mBase, uint32(v96))) = uint8(v101)
	goto L30
}
func F_sdscatprintf(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
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
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	*(*int32)(unsafe.Add(mBase, uint32(v7)+12)) = l2
	v10 = F_sdscatvprintf(m, l0, l1, l2)
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return int32(0)
	} else {
		m.G0 = v7 + int32(16)
		return v10
	}
}
func F_sdscatvprintf(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v37 int32
	_ = v37
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v52 int32
	_ = v52
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v91 int32
	_ = v91
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v105 int32
	_ = v105
	var v112 int32
	_ = v112
	var v115 int32
	_ = v115
	var v118 int32
	_ = v118
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v131 int32
	_ = v131
	var v133 int32
	_ = v133
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v140 int32
	_ = v140
	var v156 int32
	_ = v156
	var v162 int32
	_ = v162
	var v165 int32
	_ = v165
	var v170 int32
	_ = v170
	v7 = m.G0
	v9 = v7 - int32(1040)
	m.G0 = v9
	*(*int32)(unsafe.Add(mBase, uint32(v9)+1036)) = l2
	if l1&int32(3) == int32(0) {
		v33 = l1
		goto L6
	} else {
		goto L7
	}
L1:
	;
	m.G0 = v9 + int32(1040)
	return v170
L2:
	;
	v81 = v77
	v83 = v78
	goto L26
L3:
	;
	v72 = F_valkey_malloc(m, v68)
	mBase = m.M
	v75 = m.ExcPending
	if v75 != 0 {
		goto L21
	} else {
		goto L22
	}
L4:
	;
	v68 = v66 << (uint(int32(1)) % 32)
	if base.Ui32(int32(1025)) <= base.Ui32(v68) {
		goto L3
	} else {
		goto L20
	}
L5:
	;
	v66 = v58 - l1
	goto L4
L6:
	;
	v37 = v33
	goto L14
L7:
	;
	v19 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	if v19 != 0 {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v22 = l1
	goto L10
L9:
	;
	v66 = l1 - l1
	goto L4
L10:
	;
	v26 = v22 + int32(1)
	if v26&int32(3) == int32(0) {
		v33 = v26
		goto L6
	} else {
		goto L12
	}
L12:
	;
	v31 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v26))))
	if v31 != 0 {
		v22 = v26
		goto L10
	} else {
		goto L13
	}
L13:
	;
	v58 = v26
	goto L5
L14:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v37)))
	v46 = int32(-2139062144)
	if (int32(16843008)-v43|v43)&v46 == v46 {
		v37 = v37 + int32(4)
		goto L14
	} else {
		goto L16
	}
L15:
	;
	v52 = v37
	goto L17
L16:
	;
	goto L15
L17:
	;
	v56 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v52))))
	if v56 != 0 {
		v52 = v52 + int32(1)
		goto L17
	} else {
		goto L19
	}
L18:
	;
	v58 = v52
	goto L5
L19:
	;
	goto L18
L20:
	;
	v77 = v9
	v78 = int32(1024)
	goto L2
L21:
	;
	return int32(0)
L22:
	;
	if v72 != 0 {
		v77 = v72
		v78 = v68
		goto L2
	} else {
		goto L23
	}
L23:
	;
	v170 = int32(0)
	goto L1
L24:
	;
	F_valkey_free(m, v81)
	mBase = m.M
	v165 = m.ExcPending
	if v165 != 0 {
		goto L21
	} else {
		goto L57
	}
L25:
	;
	v105 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0+int32(-1)))))
	switch v105 & int32(7) {
	case 0:
		goto L43
	case 1:
		goto L42
	case 2:
		goto L41
	case 3:
		goto L40
	case 4:
		goto L39
	default:
		v122 = int32(0)
		goto L38
	}
L26:
	;
	v85 = *(*int32)(unsafe.Add(mBase, uint32(v9)+1036))
	*(*int32)(unsafe.Add(mBase, uint32(v9)+1032)) = v85
	v87 = F_vsnprintf(m, v81, v83, l1, v85)
	mBase = m.M
	v88 = m.ExcPending
	if v88 != 0 {
		goto L21
	} else {
		goto L29
	}
L27:
	;
	v170 = int32(0)
	goto L1
L28:
	;
	if base.Ui32(v87) < base.Ui32(v83) {
		goto L25
	} else {
		goto L32
	}
L29:
	;
	if int32(-1) < v87 {
		goto L28
	} else {
		goto L30
	}
L30:
	;
	v91 = int32(0)
	if v81 != v9 {
		v162 = v91
		goto L24
	} else {
		goto L31
	}
L31:
	;
	v170 = v91
	goto L1
L32:
	;
	if v81 == v9 {
		goto L33
	} else {
		goto L34
	}
L33:
	;
	v98 = v87 + int32(1)
	v99 = F_valkey_malloc(m, v98)
	mBase = m.M
	v100 = m.ExcPending
	if v100 != 0 {
		goto L21
	} else {
		goto L36
	}
L34:
	;
	F_valkey_free(m, v81)
	mBase = m.M
	v96 = m.ExcPending
	if v96 != 0 {
		goto L21
	} else {
		goto L35
	}
L35:
	;
	goto L33
L36:
	;
	if v99 != 0 {
		v81 = v99
		v83 = v98
		goto L26
	} else {
		goto L37
	}
L37:
	;
	goto L27
L38:
	;
	v124 = F__sdsMakeRoomFor(m, l0, v87, int32(1))
	mBase = m.M
	v125 = m.ExcPending
	if v125 != 0 {
		goto L21
	} else {
		goto L45
	}
L39:
	;
	v121 = *(*int32)(unsafe.Add(mBase, uint32(l0+int32(-17))))
	v122 = v121
	goto L38
L40:
	;
	v118 = *(*int32)(unsafe.Add(mBase, uint32(l0+int32(-9))))
	v122 = v118
	goto L38
L41:
	;
	v115 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0+int32(-5)))))
	v122 = v115
	goto L38
L42:
	;
	v112 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0+int32(-3)))))
	v122 = v112
	goto L38
L43:
	;
	v122 = int32(base.Ui32(v105) >> (uint(int32(3)) % 32))
	goto L38
L44:
	;
	if v81 == v9 {
		v170 = v124
		goto L1
	} else {
		goto L56
	}
L45:
	;
	if v124 == int32(0) {
		goto L44
	} else {
		goto L46
	}
L46:
	;
	if v87 == int32(0) {
		goto L48
	} else {
		goto L49
	}
L47:
	;
	v133 = v122 + v87
	v135 = v124 + int32(-1)
	v136 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v135))))
	switch v136 & int32(7) {
	case 0:
		goto L55
	case 1:
		goto L54
	case 2:
		goto L53
	case 3:
		goto L52
	case 4:
		goto L51
	default:
		goto L50
	}
L48:
	;
	goto L47
L49:
	;
	v131 = F__emscripten_memcpy_bulkmem(m, v124+v122, v81, v87)
	mBase = m.M
	goto L48
L50:
	;
	v156 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v124+v133))) = uint8(v156)
	goto L44
L51:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v124+int32(-17)))) = base.I64_extend_i32_u(v133)
	goto L50
L52:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v124+int32(-9)))) = v133
	goto L50
L53:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v124+int32(-5)))) = uint16(v133)
	goto L50
L54:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v124+int32(-3)))) = uint8(v133)
	goto L50
L55:
	;
	v140 = v133 << (uint(int32(3)) % 32)
	*(*uint8)(unsafe.Add(mBase, uint32(v135))) = uint8(v140)
	goto L50
L56:
	;
	v162 = v124
	goto L24
L57:
	;
	v170 = v162
	goto L1
}
func F_sdscpylen(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v10 int32
	_ = v10
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v75 int32
	_ = v75
	var v95 int32
	_ = v95
	v3 = l2
	v10 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0+int32(-1)))))
	switch v10 & int32(7) {
	case 0:
		v27 = int32(base.Ui32(v10) >> (uint(int32(3)) % 32))
	case 1:
		v17 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0+int32(-2)))))
		v27 = v17
	case 2:
		v20 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0+int32(-3)))))
		v27 = v20
	case 3:
		v23 = *(*int32)(unsafe.Add(mBase, uint32(l0+int32(-5))))
		v27 = v23
	case 4:
		v26 = *(*int32)(unsafe.Add(mBase, uint32(l0+int32(-9))))
		v27 = v26
	default:
		v27 = int32(0)
	}
	if base.Ui32(v27) < base.Ui32(v3) {
		v29 = int32(0)
		switch v10 & int32(7) {
		case 0:
			v47 = int32(base.Ui32(v10) >> (uint(int32(3)) % 32))
		case 1:
			v37 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0+int32(-3)))))
			v47 = v37
		case 2:
			v40 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0+int32(-5)))))
			v47 = v40
		case 3:
			v43 = *(*int32)(unsafe.Add(mBase, uint32(l0+int32(-9))))
			v47 = v43
		case 4:
			v46 = *(*int32)(unsafe.Add(mBase, uint32(l0+int32(-17))))
			v47 = v46
		default:
			v47 = v29
		}
		v50 = F__sdsMakeRoomFor(m, l0, v3-v47, int32(1))
		mBase = m.M
		v53 = m.ExcPending
		if v53 != 0 {
			return int32(0)
		} else {
			if v50 == int32(0) {
				v95 = v29
				return v95
			} else {
				v58 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v50+int32(-1)))))
				v60 = v50
				v61 = v58
				if v3 == int32(0) {
					v66 = v60
				} else {
					v65 = F__emscripten_memcpy_bulkmem(m, v60, l1, v3)
					mBase = m.M
					v66 = v65
				}
				v68 = int32(0)
				*(*uint8)(unsafe.Add(mBase, uint32(v66+v3))) = uint8(v68)
				switch v61 & int32(7) {
				case 0:
					v75 = v3 << (uint(int32(3)) % 32)
					*(*uint8)(unsafe.Add(mBase, uint32(v66+int32(-1)))) = uint8(v75)
					return v66
				case 1:
					*(*uint8)(unsafe.Add(mBase, uint32(v66+int32(-3)))) = uint8(v3)
					return v66
				case 2:
					*(*uint16)(unsafe.Add(mBase, uint32(v66+int32(-5)))) = uint16(v3)
					return v66
				case 3:
					*(*int32)(unsafe.Add(mBase, uint32(v66+int32(-9)))) = v3
					return v66
				case 4:
					*(*int64)(unsafe.Add(mBase, uint32(v66+int32(-17)))) = base.I64_extend_i32_u(v3)
					v95 = v60
					return v95
				default:
					v95 = v60
					return v95
				}
			}
		}
	} else {
		v60 = l0
		v61 = v10
		if v3 == int32(0) {
			v66 = v60
		} else {
			v65 = F__emscripten_memcpy_bulkmem(m, v60, l1, v3)
			mBase = m.M
			v66 = v65
		}
		v68 = int32(0)
		*(*uint8)(unsafe.Add(mBase, uint32(v66+v3))) = uint8(v68)
		switch v61 & int32(7) {
		case 0:
			v75 = v3 << (uint(int32(3)) % 32)
			*(*uint8)(unsafe.Add(mBase, uint32(v66+int32(-1)))) = uint8(v75)
			return v66
		case 1:
			*(*uint8)(unsafe.Add(mBase, uint32(v66+int32(-3)))) = uint8(v3)
			return v66
		case 2:
			*(*uint16)(unsafe.Add(mBase, uint32(v66+int32(-5)))) = uint16(v3)
			return v66
		case 3:
			*(*int32)(unsafe.Add(mBase, uint32(v66+int32(-9)))) = v3
			return v66
		case 4:
			*(*int64)(unsafe.Add(mBase, uint32(v66+int32(-17)))) = base.I64_extend_i32_u(v3)
			v95 = v60
			return v95
		default:
			v95 = v60
			return v95
		}
	}
}
func F_sdsdup(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	v7 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0+int32(-1)))))
	switch v7 & int32(7) {
	case 0:
		v13 = F__sdsnewlen(m, l0, int32(base.Ui32(v7)>>(uint(int32(3))%32)), int32(0))
		mBase = m.M
		v16 = m.ExcPending
		if v16 != 0 {
			return int32(0)
		} else {
			return v13
		}
	case 1:
		v20 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0+int32(-3)))))
		v22 = F__sdsnewlen(m, l0, v20, int32(0))
		mBase = m.M
		v23 = m.ExcPending
		if v23 != 0 {
			return int32(0)
		} else {
			return v22
		}
	case 2:
		v27 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0+int32(-5)))))
		v29 = F__sdsnewlen(m, l0, v27, int32(0))
		mBase = m.M
		v30 = m.ExcPending
		if v30 != 0 {
			return int32(0)
		} else {
			return v29
		}
	case 3:
		v34 = *(*int32)(unsafe.Add(mBase, uint32(l0+int32(-9))))
		v36 = F__sdsnewlen(m, l0, v34, int32(0))
		mBase = m.M
		v37 = m.ExcPending
		if v37 != 0 {
			return int32(0)
		} else {
			return v36
		}
	case 4:
		v41 = *(*int32)(unsafe.Add(mBase, uint32(l0+int32(-17))))
		v42 = v41
		v44 = F__sdsnewlen(m, l0, v42, int32(0))
		mBase = m.M
		v45 = m.ExcPending
		if v45 != 0 {
			return int32(0)
		} else {
			return v44
		}
	default:
		v42 = int32(0)
		v44 = F__sdsnewlen(m, l0, v42, int32(0))
		mBase = m.M
		v45 = m.ExcPending
		if v45 != 0 {
			return int32(0)
		} else {
			return v44
		}
	}
}
func F_sdsempty(m *base.Module) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	v11 = F_zmalloc_usable(m, int32(4), v6+int32(12))
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		return int32(0)
	} else {
		if v11 != 0 {
			v16 = *(*int32)(unsafe.Add(mBase, uint32(v6)+12))
			v20 = v16 + int32(-4)
			if base.Ui32(v20) < base.Ui32(int32(65531)) {
				v23 = int32(2)
			} else {
				v23 = int32(3)
			}
			if base.Ui32(int32(255)) < base.Ui32(v20) {
				v27 = v23
			} else {
				v27 = int32(1)
			}
			v30 = F_sdswrite(m, v11, v16, v27, int32(_a_F_sdsempty_0), int32(0))
			mBase = m.M
			v31 = m.ExcPending
			if v31 != 0 {
				return int32(0)
			} else {
				v32 = v30
				m.G0 = v6 + int32(16)
				return v32
			}
		} else {
			v32 = int32(0)
			m.G0 = v6 + int32(16)
			return v32
		}
	}
}
func F_sdsneedsrepr(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v9 int32
	_ = v9
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v55 int32
	_ = v55
	var v60 int32
	_ = v60
	v2 = int32(0)
	v9 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0+int32(-1)))))
	switch v9 & int32(7) {
	case 0:
		goto L7
	case 1:
		goto L6
	case 2:
		goto L5
	case 3:
		goto L4
	case 4:
		goto L3
	default:
		v60 = v2
		goto L1
	}
L1:
	;
	return v60
L2:
	;
	if v26 == int32(0) {
		v60 = v2
		goto L1
	} else {
		goto L8
	}
L3:
	;
	v25 = *(*int32)(unsafe.Add(mBase, uint32(l0+int32(-17))))
	v26 = v25
	goto L2
L4:
	;
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l0+int32(-9))))
	v26 = v22
	goto L2
L5:
	;
	v19 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0+int32(-5)))))
	v26 = v19
	goto L2
L6:
	;
	v16 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0+int32(-3)))))
	v26 = v16
	goto L2
L7:
	;
	v26 = int32(base.Ui32(v9) >> (uint(int32(3)) % 32))
	goto L2
L8:
	;
	v29 = l0
	v32 = v26
	goto L10
L9:
	;
	v60 = int32(1)
	goto L1
L10:
	;
	v34 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v29))))
	v36 = v34 + int32(-7)
	if base.Ui32(int32(27)) < base.Ui32(v36) {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	if v34 == int32(92) {
		goto L9
	} else {
		goto L15
	}
L13:
	;
	if int32(1)<<(uint(v36)%32)&int32(134217807) != 0 {
		goto L9
	} else {
		goto L14
	}
L14:
	;
	goto L12
L15:
	;
	if v34 == int32(32) {
		goto L9
	} else {
		goto L16
	}
L16:
	;
	if base.Ui32(base.I32_extend8_s(v34)+int32(-127)) <= base.Ui32(int32(-96)) {
		goto L9
	} else {
		goto L17
	}
L17:
	;
	v55 = v32 + int32(-1)
	if v55 == int32(0) {
		v60 = v2
		goto L1
	} else {
		goto L18
	}
L18:
	;
	v29 = v29 + int32(1)
	v32 = v55
	goto L10
}
func F_sdsnewlen(m *base.Module, l0 int32, l1 int32) int32 {
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	v4 = F__sdsnewlen(m, l0, l1, int32(0))
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		return v4
	}
}
func F_sdsnsplitargs_internal(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v35 int32
	_ = v35
	var v42 int32
	_ = v42
	var v48 int32
	_ = v48
	var v67 int32
	_ = v67
	var v84 int32
	_ = v84
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v106 int32
	_ = v106
	var v116 int32
	_ = v116
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v124 int32
	_ = v124
	var v127 int32
	_ = v127
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v149 int32
	_ = v149
	var v152 int32
	_ = v152
	var v155 int32
	_ = v155
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v175 int32
	_ = v175
	var v191 int32
	_ = v191
	var v194 int32
	_ = v194
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v204 int32
	_ = v204
	var v209 int32
	_ = v209
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v226 int32
	_ = v226
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v245 int32
	_ = v245
	var v261 int32
	_ = v261
	v4 = int32(0)
	v14 = m.G0
	v16 = v14 - int32(16)
	m.G0 = v16
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v4
	v22 = l0
	v26 = v4
	v27 = v4
	goto L3
L1:
	;
	F__serverAssert(m, int32(_a_F_sdsnsplitargs_internal_0), int32(_a_F_sdsnsplitargs_internal_1), int32(1165))
	mBase = m.M
	v261 = m.ExcPending
	if v261 != 0 {
		goto L19
	} else {
		goto L56
	}
L2:
	;
	m.G0 = v16 + int32(16)
	return v245
L3:
	;
	v35 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22))))
	if v35 == int32(0) {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	v245 = int32(0)
	goto L2
L5:
	;
	v88 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v16)+12)) = v88
	v93 = F_sdsparsearg(m, v42, l1, v16+int32(12), v88)
	mBase = m.M
	if v93 != 0 {
		goto L22
	} else {
		goto L23
	}
L6:
	;
	if v27 != 0 {
		v245 = v27
		goto L2
	} else {
		goto L18
	}
L7:
	;
	if base.B2i32(l1 != int32(0))&base.B2i32(base.Ui32(l1) <= base.Ui32(v22)) != 0 {
		goto L6
	} else {
		goto L8
	}
L8:
	;
	v42 = v22
	v48 = v35
	goto L9
L9:
	;
	if l1 == int32(0) {
		goto L11
	} else {
		goto L12
	}
L10:
	;
	goto L6
L11:
	;
	if v48&int32(255) == int32(32) {
		goto L14
	} else {
		goto L15
	}
L12:
	;
	if base.Ui32(l1) <= base.Ui32(v42) {
		goto L6
	} else {
		goto L13
	}
L13:
	;
	goto L11
L14:
	;
	v67 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v42)+1)))
	if v67 != 0 {
		v42 = v42 + int32(1)
		v48 = v67
		goto L9
	} else {
		goto L17
	}
L15:
	;
	if base.Ui32(base.I32_extend8_s(v48)+int32(-14)) < base.Ui32(int32(-5)) {
		goto L5
	} else {
		goto L16
	}
L16:
	;
	goto L14
L17:
	;
	goto L10
L18:
	;
	v84 = F_valkey_malloc(m, int32(4))
	mBase = m.M
	v87 = m.ExcPending
	if v87 != 0 {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	return int32(0)
L20:
	;
	v245 = v84
	goto L2
L21:
	;
	if v93 != 0 {
		v22 = v226
		v26 = v230
		v27 = v231
		goto L3
	} else {
		goto L55
	}
L22:
	;
	v194 = int32(0)
	v196 = *(*int32)(unsafe.Add(mBase, _c_F_sdsnsplitargs_internal[0]))
	v197 = *(*int32)(unsafe.Add(mBase, uint32(v16)+12))
	v199 = F__sdsnewlen(m, v196, v197, v194)
	mBase = m.M
	v200 = m.ExcPending
	if v200 != 0 {
		goto L19
	} else {
		goto L47
	}
L23:
	;
	v94 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	v96 = v94 + int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v96
	if v94 == int32(0) {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	F_valkey_free(m, v27)
	mBase = m.M
	v191 = m.ExcPending
	if v191 != 0 {
		goto L19
	} else {
		goto L46
	}
L25:
	;
	v106 = v96
	goto L26
L26:
	;
	v116 = *(*int32)(unsafe.Add(mBase, uint32(v27+v106<<(uint(int32(2))%32))))
	if v116 == int32(0) {
		v168 = v106
		goto L28
	} else {
		goto L29
	}
L27:
	;
	goto L24
L28:
	;
	v175 = v168 + int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v175
	if v168 != 0 {
		v106 = v175
		goto L26
	} else {
		goto L45
	}
L29:
	;
	v121 = v116 + int32(-1)
	v122 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v121))))
	v124 = v122 & int32(7)
	switch v124 {
	case 0:
		goto L36
	case 1:
		goto L35
	case 2:
		goto L34
	case 3:
		goto L33
	case 4:
		goto L32
	default:
		v142 = v116
		v143 = int32(1)
		goto L31
	}
L30:
	;
	F_zfree_with_size(m, v161, v162)
	mBase = m.M
	v166 = m.ExcPending
	if v166 != 0 {
		goto L19
	} else {
		goto L44
	}
L31:
	;
	switch v124 {
	case 0:
		goto L43
	case 1:
		goto L42
	case 2:
		goto L41
	case 3:
		goto L40
	case 4:
		goto L39
	default:
		v159 = int32(0)
		goto L38
	}
L32:
	;
	v142 = v116 + int32(-17)
	v143 = int32(18)
	goto L31
L33:
	;
	v142 = v116 + int32(-9)
	v143 = int32(10)
	goto L31
L34:
	;
	v142 = v116 + int32(-5)
	v143 = int32(6)
	goto L31
L35:
	;
	v142 = v116 + int32(-3)
	v143 = int32(4)
	goto L31
L36:
	;
	v127 = *(*int32)(unsafe.Add(mBase, uint32(v116+int32(-9))))
	goto L37
L37:
	;
	v161 = v121
	v162 = v127 & int32(2147483647)
	goto L30
L38:
	;
	v161 = v142
	v162 = v159 + v143
	goto L30
L39:
	;
	v158 = *(*int32)(unsafe.Add(mBase, uint32(v116+int32(-9))))
	v159 = v158
	goto L38
L40:
	;
	v155 = *(*int32)(unsafe.Add(mBase, uint32(v116+int32(-5))))
	v159 = v155
	goto L38
L41:
	;
	v152 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v116+int32(-3)))))
	v159 = v152
	goto L38
L42:
	;
	v149 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v116+int32(-2)))))
	v159 = v149
	goto L38
L43:
	;
	v159 = int32(base.Ui32(v122) >> (uint(int32(3)) % 32))
	goto L38
L44:
	;
	v167 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	v168 = v167
	goto L28
L45:
	;
	goto L27
L46:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = int32(0)
	v226 = v42
	v230 = v26
	v231 = v27
	goto L21
L47:
	;
	v201 = F_sdsparsearg(m, v42, l1, v194, v199)
	mBase = m.M
	if v201 <= int32(0) {
		goto L1
	} else {
		goto L48
	}
L48:
	;
	v204 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	if v204 != v26 {
		v215 = v26
		v216 = v27
		v217 = v204
		goto L49
	} else {
		goto L50
	}
L49:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v217 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v216+v217<<(uint(int32(2))%32)))) = v199
	v226 = v42 + v201
	v230 = v215
	v231 = v216
	goto L21
L50:
	;
	if v26 != 0 {
		goto L51
	} else {
		goto L52
	}
L51:
	;
	v209 = v26 << (uint(int32(1)) % 32)
	goto L53
L52:
	;
	v209 = int32(8)
	goto L53
L53:
	;
	v212 = F_valkey_realloc(m, v27, v209<<(uint(int32(2))%32))
	mBase = m.M
	v213 = m.ExcPending
	if v213 != 0 {
		goto L19
	} else {
		goto L54
	}
L54:
	;
	v214 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	v215 = v209
	v216 = v212
	v217 = v214
	goto L49
L55:
	;
	goto L4
L56:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_sdssplitargs(m *base.Module, l0 int32, l1 int32) int32 {
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	v4 = F_sdsnsplitargs_internal(m, l0, int32(0), l1)
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		return v4
	}
}
func F_sdssubstr(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v84 int32
	_ = v84
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v99 int32
	_ = v99
	var v106 int32
	_ = v106
	var v110 int32
	_ = v110
	var v113 int32
	_ = v113
	var v119 int32
	_ = v119
	var v126 int32
	_ = v126
	var v130 int32
	_ = v130
	var v133 int32
	_ = v133
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v147 int32
	_ = v147
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v152 int32
	_ = v152
	var v154 int32
	_ = v154
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v169 int32
	_ = v169
	var v171 int32
	_ = v171
	var v176 int32
	_ = v176
	var v189 int32
	_ = v189
	var v192 int32
	_ = v192
	v10 = l0 + int32(-1)
	v11 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10))))
	v13 = v11 & int32(7)
	switch v13 {
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
		v28 = int32(0)
		goto L1
	}
L1:
	;
	v30 = base.B2i32(base.Ui32(l1) < base.Ui32(v28))
	if base.Ui32(l1) < base.Ui32(v28) {
		goto L8
	} else {
		goto L9
	}
L2:
	;
	v27 = *(*int32)(unsafe.Add(mBase, uint32(l0+int32(-17))))
	v28 = v27
	goto L1
L3:
	;
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l0+int32(-9))))
	v28 = v24
	goto L1
L4:
	;
	v21 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0+int32(-5)))))
	v28 = v21
	goto L1
L5:
	;
	v18 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0+int32(-3)))))
	v28 = v18
	goto L1
L6:
	;
	v28 = int32(base.Ui32(v11) >> (uint(int32(3)) % 32))
	goto L1
L7:
	;
	v189 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l0+v36))) = uint8(v189)
	switch v13 {
	case 0:
		goto L64
	case 1:
		goto L63
	case 2:
		goto L62
	case 3:
		goto L61
	case 4:
		goto L60
	default:
		goto L59
	}
L8:
	;
	v31 = l1
	goto L10
L9:
	;
	v31 = int32(0)
	goto L10
L10:
	;
	v32 = v28 - v31
	if base.Ui32(l2) < base.Ui32(v32) {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v34 = l2
	goto L13
L12:
	;
	v34 = v32
	goto L13
L13:
	;
	if base.Ui32(l1) < base.Ui32(v28) {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v36 = v34
	goto L16
L15:
	;
	v36 = int32(0)
	goto L16
L16:
	;
	if v36 == int32(0) {
		goto L7
	} else {
		goto L17
	}
L17:
	;
	v39 = l0 + v31
	if l0 == v39 {
		goto L19
	} else {
		goto L20
	}
L18:
	;
	goto L7
L19:
	;
	goto L18
L20:
	;
	v43 = v36 + l0
	if base.Ui32(int32(0)-v36<<(uint(int32(1))%32)) < base.Ui32(v39-v43) {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	v53 = (v39 ^ l0) & int32(3)
	if base.Ui32(v39) <= base.Ui32(l0) {
		goto L25
	} else {
		goto L26
	}
L22:
	;
	v50 = F___memcpy(m, l0, v39, v36)
	mBase = m.M
	goto L18
L23:
	;
	if v159 == int32(0) {
		goto L19
	} else {
		goto L55
	}
L24:
	;
	if base.Ui32(v137) <= base.Ui32(int32(3)) {
		v158 = v136
		v159 = v137
		v160 = v138
		goto L23
	} else {
		goto L51
	}
L25:
	;
	if v53 != 0 {
		v119 = v36
		goto L35
	} else {
		goto L36
	}
L26:
	;
	if v53 == int32(0) {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	if l0&int32(3) != 0 {
		goto L29
	} else {
		goto L30
	}
L28:
	;
	v158 = v39
	v159 = v36
	v160 = l0
	goto L23
L29:
	;
	v60 = v39
	v61 = v36
	v62 = l0
	goto L31
L30:
	;
	v136 = v39
	v137 = v36
	v138 = l0
	goto L24
L31:
	;
	if v61 == int32(0) {
		goto L19
	} else {
		goto L33
	}
L33:
	;
	v66 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v60))))
	*(*uint8)(unsafe.Add(mBase, uint32(v62))) = uint8(v66)
	v68 = int32(1)
	v69 = v60 + v68
	v71 = v61 + int32(-1)
	v73 = v62 + v68
	if v73&int32(3) == int32(0) {
		v136 = v69
		v137 = v71
		v138 = v73
		goto L24
	} else {
		goto L34
	}
L34:
	;
	v60 = v69
	v61 = v71
	v62 = v73
	goto L31
L35:
	;
	if v119 == int32(0) {
		goto L19
	} else {
		goto L47
	}
L36:
	;
	if v43&int32(3) == int32(0) {
		v99 = v36
		goto L37
	} else {
		goto L38
	}
L37:
	;
	if base.Ui32(v99) <= base.Ui32(int32(3)) {
		v119 = v99
		goto L35
	} else {
		goto L43
	}
L38:
	;
	v84 = v36
	goto L39
L39:
	;
	if v84 == int32(0) {
		goto L19
	} else {
		goto L41
	}
L40:
	;
	v99 = v90
	goto L37
L41:
	;
	v90 = v84 + int32(-1)
	v91 = l0 + v90
	v93 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v39+v90))))
	*(*uint8)(unsafe.Add(mBase, uint32(v91))) = uint8(v93)
	if v91&int32(3) != 0 {
		v84 = v90
		goto L39
	} else {
		goto L42
	}
L42:
	;
	goto L40
L43:
	;
	v106 = v99
	goto L44
L44:
	;
	v110 = v106 + int32(-4)
	v113 = *(*int32)(unsafe.Add(mBase, uint32(v39+v110)))
	*(*int32)(unsafe.Add(mBase, uint32(l0+v110))) = v113
	if base.Ui32(int32(3)) < base.Ui32(v110) {
		v106 = v110
		goto L44
	} else {
		goto L46
	}
L45:
	;
	v119 = v110
	goto L35
L46:
	;
	goto L45
L47:
	;
	v126 = v119
	goto L48
L48:
	;
	v130 = v126 + int32(-1)
	v133 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v39+v130))))
	*(*uint8)(unsafe.Add(mBase, uint32(l0+v130))) = uint8(v133)
	if v130 != 0 {
		v126 = v130
		goto L48
	} else {
		goto L50
	}
L50:
	;
	goto L19
L51:
	;
	v143 = v136
	v144 = v137
	v145 = v138
	goto L52
L52:
	;
	v147 = *(*int32)(unsafe.Add(mBase, uint32(v143)))
	*(*int32)(unsafe.Add(mBase, uint32(v145))) = v147
	v149 = int32(4)
	v150 = v143 + v149
	v152 = v145 + v149
	v154 = v144 + int32(-4)
	if base.Ui32(int32(3)) < base.Ui32(v154) {
		v143 = v150
		v144 = v154
		v145 = v152
		goto L52
	} else {
		goto L54
	}
L53:
	;
	v158 = v150
	v159 = v154
	v160 = v152
	goto L23
L54:
	;
	goto L53
L55:
	;
	v165 = v158
	v166 = v159
	v167 = v160
	goto L56
L56:
	;
	v169 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v165))))
	*(*uint8)(unsafe.Add(mBase, uint32(v167))) = uint8(v169)
	v171 = int32(1)
	v176 = v166 + int32(-1)
	if v176 != 0 {
		v165 = v165 + v171
		v166 = v176
		v167 = v167 + v171
		goto L56
	} else {
		goto L58
	}
L57:
	;
	goto L19
L58:
	;
	goto L57
L59:
	;
	return
L60:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l0+int32(-17)))) = base.I64_extend_i32_u(v36)
	goto L59
L61:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0+int32(-9)))) = v36
	return
L62:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(l0+int32(-5)))) = uint16(v36)
	return
L63:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(l0+int32(-3)))) = uint8(v36)
	return
L64:
	;
	v192 = v36 << (uint(int32(3)) % 32)
	*(*uint8)(unsafe.Add(mBase, uint32(v10))) = uint8(v192)
	return
}
func F_select_(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v16 int32
	_ = v16
	var v17 int64
	_ = v17
	var v18 int32
	_ = v18
	var v19 int64
	_ = v19
	var v27 int32
	_ = v27
	var v34 int32
	_ = v34
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v49 int32
	_ = v49
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	v10 = m.G0
	v12 = v10 - int32(16)
	m.G0 = v12
	if l4 != 0 {
		v16 = *(*int32)(unsafe.Add(mBase, uint32(l4)+8))
		v17 = *(*int64)(unsafe.Add(mBase, uint32(l4)))
		v18 = v16
		v19 = v17
	} else {
		v18 = int32(0)
		v19 = int64(0)
	}
	if v19 < int64(0) {
		v27 = F___errno_location(m)
		mBase = m.M
		*(*int32)(unsafe.Add(mBase, uint32(v27))) = int32(28)
		v63 = int32(-1)
	} else {
		if int32(-1) < v18 {
			v34 = base.I32_div_u_s(v18, int32(1000000))
			if l4 != 0 {
				v43 = base.B2i32(base.Ui64(v19^int64(9223372036854775807)) < base.Ui64(base.I64_extend_i32_u(v34)))
				if base.Ui64(v19^int64(9223372036854775807)) < base.Ui64(base.I64_extend_i32_u(v34)) {
					v44 = int32(999999)
				} else {
					v44 = v18 - v34*int32(1000000)
				}
				*(*int32)(unsafe.Add(mBase, uint32(v12)+12)) = v44
				if base.Ui64(v19^int64(9223372036854775807)) < base.Ui64(base.I64_extend_i32_u(v34)) {
					v49 = int32(-1)
				} else {
					v49 = v34 + base.I32_wrap_i64(v19)
				}
				*(*int32)(unsafe.Add(mBase, uint32(v12)+8)) = v49
				v53 = v12 + int32(8)
			} else {
				v53 = int32(0)
			}
			v54 = m.Env.X__syscall__newselect(m, l0, l1, l2, l3, v53)
			mBase = m.M
			if base.Ui32(v54) < base.Ui32(int32(-4095)) {
				v62 = v54
			} else {
				v57 = F___errno_location(m)
				mBase = m.M
				*(*int32)(unsafe.Add(mBase, uint32(v57))) = int32(0) - v54
				v62 = int32(-1)
			}
			v63 = v62
		} else {
			v27 = F___errno_location(m)
			mBase = m.M
			*(*int32)(unsafe.Add(mBase, uint32(v27))) = int32(28)
			v63 = int32(-1)
		}
	}
	m.G0 = v12 + int32(16)
	return v63
}
func F_send(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	v5 = int32(0)
	v7 = F_sendto(m, l0, l1, l2, l3, v5, v5)
	return v7
}
func F_sendBulkToReplica(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v21 int32
	_ = v21
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v64 int64
	_ = v64
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v86 int32
	_ = v86
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v102 int32
	_ = v102
	var v106 int32
	_ = v106
	var v109 int32
	_ = v109
	var v114 int32
	_ = v114
	var v117 int32
	_ = v117
	var v120 int32
	_ = v120
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v131 int32
	_ = v131
	var v133 int32
	_ = v133
	var v135 int32
	_ = v135
	var v139 int32
	_ = v139
	var v141 int32
	_ = v141
	var v144 int32
	_ = v144
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v175 int32
	_ = v175
	var v182 int32
	_ = v182
	var v185 int32
	_ = v185
	var v188 int32
	_ = v188
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v199 int32
	_ = v199
	var v202 int32
	_ = v202
	var v203 int64
	_ = v203
	var v204 int32
	_ = v204
	var v205 int64
	_ = v205
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v211 int32
	_ = v211
	var v215 int32
	_ = v215
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v227 int32
	_ = v227
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
	var v239 int32
	_ = v239
	var v243 int32
	_ = v243
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v248 int32
	_ = v248
	var v249 int32
	_ = v249
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v258 int32
	_ = v258
	var v259 int32
	_ = v259
	var v260 int64
	_ = v260
	var v261 int64
	_ = v261
	var v262 int64
	_ = v262
	var v264 int32
	_ = v264
	var v266 int64
	_ = v266
	var v269 int64
	_ = v269
	var v272 int32
	_ = v272
	var v273 int32
	_ = v273
	var v274 int32
	_ = v274
	var v276 int32
	_ = v276
	var v277 int32
	_ = v277
	var v278 int32
	_ = v278
	var v279 int32
	_ = v279
	var v280 int32
	_ = v280
	var v281 int32
	_ = v281
	var v282 int32
	_ = v282
	var v283 int32
	_ = v283
	var v284 int32
	_ = v284
	var v287 int32
	_ = v287
	var v291 int64
	_ = v291
	var v295 int32
	_ = v295
	var v302 int32
	_ = v302
	var v305 int32
	_ = v305
	var v308 int32
	_ = v308
	var v311 int32
	_ = v311
	var v314 int32
	_ = v314
	var v315 int32
	_ = v315
	var v318 int32
	_ = v318
	var v323 int32
	_ = v323
	var v341 int32
	_ = v341
	v9 = m.G0
	v11 = v9 - int32(16448)
	m.G0 = v11
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(v13)+104))
	v15 = *(*int32)(unsafe.Add(mBase, uint32(v14)+32))
	if v15 == int32(0) {
		v199 = v14
		v202 = *(*int32)(unsafe.Add(mBase, uint32(v199)+8))
		v203 = *(*int64)(unsafe.Add(mBase, uint32(v199)+16))
		v204 = int32(0)
		v205 = F___lseek(m, v202, v203, v204)
		mBase = m.M
		v206 = *(*int32)(unsafe.Add(mBase, uint32(v13)+104))
		v207 = *(*int32)(unsafe.Add(mBase, uint32(v206)+8))
		v211 = F_read(m, v207, v11+int32(48), int32(16384))
		mBase = m.M
		if v204 < v211 {
			v233 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			v234 = *(*int32)(unsafe.Add(mBase, uint32(v233)+68))
			v235 = m.T0[v234].(func(*base.Module, int32, int32, int32) int32)(m, l0, v11+int32(48), v211)
			mBase = m.M
			v236 = m.ExcPending
			if v236 != 0 {
				return
			} else {
				if v235 != int32(-1) {
					v259 = *(*int32)(unsafe.Add(mBase, uint32(v13)+104))
					v260 = *(*int64)(unsafe.Add(mBase, uint32(v259)+16))
					v261 = base.I64_extend_i32_s(v235)
					v262 = v260 + v261
					*(*int64)(unsafe.Add(mBase, uint32(v259)+16)) = v262
					v264 = int32(_a_F_sendBulkToReplica_0)
					v266 = *(*int64)(unsafe.Add(mBase, _c_F_sendBulkToReplica[0]))
					*(*int64)(unsafe.Add(mBase, _c_F_sendBulkToReplica[0])) = v266 + v261
					v269 = *(*int64)(unsafe.Add(mBase, uint32(v259)+24))
					if v262 != v269 {
						m.G0 = v11 + int32(16448)
						return
					} else {
						F_closeRepldbfd(m, v13)
						mBase = m.M
						v272 = m.ExcPending
						if v272 != 0 {
							return
						} else {
							v273 = *(*int32)(unsafe.Add(mBase, uint32(v13)+8))
							v274 = int32(0)
							v276 = *(*int32)(unsafe.Add(mBase, uint32(v273)))
							v277 = *(*int32)(unsafe.Add(mBase, uint32(v276)+80))
							v278 = m.T0[v277].(func(*base.Module, int32, int32, int32) int32)(m, v273, v274, v274)
							mBase = m.M
							v279 = m.ExcPending
							if v279 != 0 {
								return
							} else {
								v280 = F_replicaPutOnline(m, v13)
								mBase = m.M
								v281 = m.ExcPending
								if v281 != 0 {
									return
								} else {
									if v280 != 0 {
										v284 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+205)))
										if v284&int32(32) != 0 {
											F__serverAssert(m, int32(_a_F_sendBulkToReplica_1), int32(_a_F_sendBulkToReplica_2), int32(1617))
											mBase = m.M
											v341 = m.ExcPending
											if v341 != 0 {
												return
											} else {
												F_abort(m)
												mBase = m.M
												base.Wasm_trap_unreachable()
												for {
												}
											}
										} else {
											v287 = *(*int32)(unsafe.Add(mBase, uint32(v13)+104))
											*(*int32)(unsafe.Add(mBase, uint32(v287)+4)) = int32(0)
											v291 = *(*int64)(unsafe.Add(mBase, _c_F_sendBulkToReplica[1]))
											if v291 != int64(0) {
												v305 = *(*int32)(unsafe.Add(mBase, uint32(v13)+200))
												if v305&int32(4194304) != 0 {
												} else {
													v308 = *(*int32)(unsafe.Add(mBase, uint32(v13)+104))
													if v308 == int32(0) {
														v315 = F_clusterSlotMigrationShouldInstallWriteHandler(m, v13)
														mBase = m.M
														if v315 == int32(0) {
														} else {
															v318 = *(*int32)(unsafe.Add(mBase, uint32(v13)+200))
															*(*int32)(unsafe.Add(mBase, uint32(v13)+200)) = v318 | int32(4194304)
															v323 = *(*int32)(unsafe.Add(mBase, _c_F_sendBulkToReplica[2]))
															F_listLinkNodeHead(m, v323, v13+int32(168))
															mBase = m.M
														}
													} else {
														v311 = *(*int32)(unsafe.Add(mBase, uint32(v308)))
														switch v311 {
														case 0:
															v315 = F_clusterSlotMigrationShouldInstallWriteHandler(m, v13)
															mBase = m.M
															if v315 == int32(0) {
															} else {
																v318 = *(*int32)(unsafe.Add(mBase, uint32(v13)+200))
																*(*int32)(unsafe.Add(mBase, uint32(v13)+200)) = v318 | int32(4194304)
																v323 = *(*int32)(unsafe.Add(mBase, _c_F_sendBulkToReplica[2]))
																F_listLinkNodeHead(m, v323, v13+int32(168))
																mBase = m.M
															}
														default:
														case 9, 11:
															if v305&int32(1024) != 0 {
															} else {
																v314 = *(*int32)(unsafe.Add(mBase, uint32(v308)+4))
																if v314 != 0 {
																} else {
																	v315 = F_clusterSlotMigrationShouldInstallWriteHandler(m, v13)
																	mBase = m.M
																	if v315 == int32(0) {
																	} else {
																		v318 = *(*int32)(unsafe.Add(mBase, uint32(v13)+200))
																		*(*int32)(unsafe.Add(mBase, uint32(v13)+200)) = v318 | int32(4194304)
																		v323 = *(*int32)(unsafe.Add(mBase, _c_F_sendBulkToReplica[2]))
																		F_listLinkNodeHead(m, v323, v13+int32(168))
																		mBase = m.M
																	}
																}
															}
														}
													}
												}
												m.G0 = v11 + int32(16448)
												return
											} else {
												v295 = *(*int32)(unsafe.Add(mBase, _c_F_sendBulkToReplica[3]))
												*(*int32)(unsafe.Add(mBase, uint32(v11)+uint32(_c_F_sendBulkToReplica[4]))) = v295
												F_replicationFeedReplicas(m, int32(-1), v11+int32(16444), int32(1))
												mBase = m.M
												v302 = m.ExcPending
												if v302 != 0 {
													return
												} else {
													v305 = *(*int32)(unsafe.Add(mBase, uint32(v13)+200))
													if v305&int32(4194304) != 0 {
													} else {
														v308 = *(*int32)(unsafe.Add(mBase, uint32(v13)+104))
														if v308 == int32(0) {
															v315 = F_clusterSlotMigrationShouldInstallWriteHandler(m, v13)
															mBase = m.M
															if v315 == int32(0) {
															} else {
																v318 = *(*int32)(unsafe.Add(mBase, uint32(v13)+200))
																*(*int32)(unsafe.Add(mBase, uint32(v13)+200)) = v318 | int32(4194304)
																v323 = *(*int32)(unsafe.Add(mBase, _c_F_sendBulkToReplica[2]))
																F_listLinkNodeHead(m, v323, v13+int32(168))
																mBase = m.M
															}
														} else {
															v311 = *(*int32)(unsafe.Add(mBase, uint32(v308)))
															switch v311 {
															case 0:
																v315 = F_clusterSlotMigrationShouldInstallWriteHandler(m, v13)
																mBase = m.M
																if v315 == int32(0) {
																} else {
																	v318 = *(*int32)(unsafe.Add(mBase, uint32(v13)+200))
																	*(*int32)(unsafe.Add(mBase, uint32(v13)+200)) = v318 | int32(4194304)
																	v323 = *(*int32)(unsafe.Add(mBase, _c_F_sendBulkToReplica[2]))
																	F_listLinkNodeHead(m, v323, v13+int32(168))
																	mBase = m.M
																}
															default:
															case 9, 11:
																if v305&int32(1024) != 0 {
																} else {
																	v314 = *(*int32)(unsafe.Add(mBase, uint32(v308)+4))
																	if v314 != 0 {
																	} else {
																		v315 = F_clusterSlotMigrationShouldInstallWriteHandler(m, v13)
																		mBase = m.M
																		if v315 == int32(0) {
																		} else {
																			v318 = *(*int32)(unsafe.Add(mBase, uint32(v13)+200))
																			*(*int32)(unsafe.Add(mBase, uint32(v13)+200)) = v318 | int32(4194304)
																			v323 = *(*int32)(unsafe.Add(mBase, _c_F_sendBulkToReplica[2]))
																			F_listLinkNodeHead(m, v323, v13+int32(168))
																			mBase = m.M
																		}
																	}
																}
															}
														}
													}
													m.G0 = v11 + int32(16448)
													return
												}
											}
										}
									} else {
										v282 = F_freeClient(m, v13)
										mBase = m.M
										v283 = m.ExcPending
										if v283 != 0 {
											return
										} else {
											m.G0 = v11 + int32(16448)
											return
										}
									}
								}
							}
						}
					}
				} else {
					v239 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
					if v239 == int32(3) {
						m.G0 = v11 + int32(16448)
						return
					} else {
						v243 = *(*int32)(unsafe.Add(mBase, _c_F_sendBulkToReplica[5]))
						if int32(3) < v243 {
							v257 = F_freeClient(m, v13)
							mBase = m.M
							v258 = m.ExcPending
							if v258 != 0 {
								return
							} else {
								m.G0 = v11 + int32(16448)
								return
							}
						} else {
							v246 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
							v247 = *(*int32)(unsafe.Add(mBase, uint32(v246)+88))
							v248 = m.T0[v247].(func(*base.Module, int32) int32)(m, l0)
							mBase = m.M
							v249 = m.ExcPending
							if v249 != 0 {
								return
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = v248
								F__serverLog(m, int32(3), int32(_a_F_sendBulkToReplica_3), v11+int32(16))
								mBase = m.M
								v256 = m.ExcPending
								if v256 != 0 {
									return
								} else {
									v257 = F_freeClient(m, v13)
									mBase = m.M
									v258 = m.ExcPending
									if v258 != 0 {
										return
									} else {
										m.G0 = v11 + int32(16448)
										return
									}
								}
							}
						}
					}
				}
			}
		} else {
			v215 = *(*int32)(unsafe.Add(mBase, _c_F_sendBulkToReplica[5]))
			if int32(3) < v215 {
				v229 = F_freeClient(m, v13)
				mBase = m.M
				v230 = m.ExcPending
				if v230 != 0 {
					return
				} else {
					m.G0 = v11 + int32(16448)
					return
				}
			} else {
				if v211 != 0 {
					v220 = *(*int32)(unsafe.Add(mBase, _c_F_sendBulkToReplica[6]))
					v221 = F___strerror_l(m, v220, v220)
					mBase = m.M
					v222 = v221
				} else {
					v222 = int32(_a_F_sendBulkToReplica_4)
				}
				*(*int32)(unsafe.Add(mBase, uint32(v11))) = v222
				F__serverLog(m, int32(3), int32(_a_F_sendBulkToReplica_5), v11)
				mBase = m.M
				v227 = m.ExcPending
				if v227 != 0 {
					return
				} else {
					v229 = F_freeClient(m, v13)
					mBase = m.M
					v230 = m.ExcPending
					if v230 != 0 {
						return
					} else {
						m.G0 = v11 + int32(16448)
						return
					}
				}
			}
		}
	} else {
		v21 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15+int32(-1)))))
		switch v21 & int32(7) {
		case 0:
			v38 = int32(base.Ui32(v21) >> (uint(int32(3)) % 32))
		case 1:
			v28 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15+int32(-3)))))
			v38 = v28
		case 2:
			v31 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v15+int32(-5)))))
			v38 = v31
		case 3:
			v34 = *(*int32)(unsafe.Add(mBase, uint32(v15+int32(-9))))
			v38 = v34
		case 4:
			v37 = *(*int32)(unsafe.Add(mBase, uint32(v15+int32(-17))))
			v38 = v37
		default:
			v38 = int32(0)
		}
		v39 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		v40 = *(*int32)(unsafe.Add(mBase, uint32(v39)+68))
		v41 = m.T0[v40].(func(*base.Module, int32, int32, int32) int32)(m, l0, v15, v38)
		mBase = m.M
		v42 = m.ExcPending
		if v42 != 0 {
			return
		} else {
			if v41 != int32(-1) {
				v62 = int32(_a_F_sendBulkToReplica_0)
				v64 = *(*int64)(unsafe.Add(mBase, _c_F_sendBulkToReplica[0]))
				*(*int64)(unsafe.Add(mBase, _c_F_sendBulkToReplica[0])) = v64 + base.I64_extend_i32_s(v41)
				v68 = *(*int32)(unsafe.Add(mBase, uint32(v13)+104))
				v69 = *(*int32)(unsafe.Add(mBase, uint32(v68)+32))
				v70 = int32(-1)
				v78 = v69 + v70
				v79 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v78))))
				v81 = v79 & int32(7)
				switch v81 {
				case 0:
					v96 = int32(base.Ui32(v79) >> (uint(int32(3)) % 32))
					if v96 == int32(0) {
					} else {
						v102 = int32(-1)&v96 + v70
						v106 = v41>>(uint(int32(31))%32)&v96 + v41
						v109 = v102 - v106 + int32(1)
						switch v81 {
						default:
							v124 = int32(base.Ui32(v79) >> (uint(int32(3)) % 32))
						case 1:
							v114 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v69+int32(-3)))))
							v124 = v114
						case 2:
							v117 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v69+int32(-5)))))
							v124 = v117
						case 3:
							v120 = *(*int32)(unsafe.Add(mBase, uint32(v69+int32(-9))))
							v124 = v120
						case 4:
							v123 = *(*int32)(unsafe.Add(mBase, uint32(v69+int32(-17))))
							v124 = v123
						}
						v125 = int32(0)
						v127 = base.B2i32(base.Ui32(v106) < base.Ui32(v124))
						if base.Ui32(v106) < base.Ui32(v124) {
							v128 = v106
						} else {
							v128 = v125
						}
						v129 = v124 - v128
						if base.Ui32(v109) < base.Ui32(v129) {
							v131 = v109
						} else {
							v131 = v129
						}
						if v102 < v106 {
							v133 = v125
						} else {
							v133 = v131
						}
						if base.Ui32(v106) < base.Ui32(v124) {
							v135 = v133
						} else {
							v135 = int32(0)
						}
						if v135 == int32(0) {
						} else {
							v139 = F_memmove(m, v69, v69+v128, v135)
							mBase = m.M
						}
						v141 = int32(0)
						*(*uint8)(unsafe.Add(mBase, uint32(v69+v135))) = uint8(v141)
						switch v81 {
						default:
							v144 = v135 << (uint(int32(3)) % 32)
							*(*uint8)(unsafe.Add(mBase, uint32(v78))) = uint8(v144)
						case 1:
							*(*uint8)(unsafe.Add(mBase, uint32(v69+int32(-3)))) = uint8(v135)
						case 2:
							*(*uint16)(unsafe.Add(mBase, uint32(v69+int32(-5)))) = uint16(v135)
						case 3:
							*(*int32)(unsafe.Add(mBase, uint32(v69+int32(-9)))) = v135
						case 4:
							*(*int64)(unsafe.Add(mBase, uint32(v69+int32(-17)))) = base.I64_extend_i32_u(v135)
						}
					}
				case 1:
					v86 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v69+int32(-3)))))
					v96 = v86
					if v96 == int32(0) {
					} else {
						v102 = int32(-1)&v96 + v70
						v106 = v41>>(uint(int32(31))%32)&v96 + v41
						v109 = v102 - v106 + int32(1)
						switch v81 {
						default:
							v124 = int32(base.Ui32(v79) >> (uint(int32(3)) % 32))
						case 1:
							v114 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v69+int32(-3)))))
							v124 = v114
						case 2:
							v117 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v69+int32(-5)))))
							v124 = v117
						case 3:
							v120 = *(*int32)(unsafe.Add(mBase, uint32(v69+int32(-9))))
							v124 = v120
						case 4:
							v123 = *(*int32)(unsafe.Add(mBase, uint32(v69+int32(-17))))
							v124 = v123
						}
						v125 = int32(0)
						v127 = base.B2i32(base.Ui32(v106) < base.Ui32(v124))
						if base.Ui32(v106) < base.Ui32(v124) {
							v128 = v106
						} else {
							v128 = v125
						}
						v129 = v124 - v128
						if base.Ui32(v109) < base.Ui32(v129) {
							v131 = v109
						} else {
							v131 = v129
						}
						if v102 < v106 {
							v133 = v125
						} else {
							v133 = v131
						}
						if base.Ui32(v106) < base.Ui32(v124) {
							v135 = v133
						} else {
							v135 = int32(0)
						}
						if v135 == int32(0) {
						} else {
							v139 = F_memmove(m, v69, v69+v128, v135)
							mBase = m.M
						}
						v141 = int32(0)
						*(*uint8)(unsafe.Add(mBase, uint32(v69+v135))) = uint8(v141)
						switch v81 {
						default:
							v144 = v135 << (uint(int32(3)) % 32)
							*(*uint8)(unsafe.Add(mBase, uint32(v78))) = uint8(v144)
						case 1:
							*(*uint8)(unsafe.Add(mBase, uint32(v69+int32(-3)))) = uint8(v135)
						case 2:
							*(*uint16)(unsafe.Add(mBase, uint32(v69+int32(-5)))) = uint16(v135)
						case 3:
							*(*int32)(unsafe.Add(mBase, uint32(v69+int32(-9)))) = v135
						case 4:
							*(*int64)(unsafe.Add(mBase, uint32(v69+int32(-17)))) = base.I64_extend_i32_u(v135)
						}
					}
				case 2:
					v89 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v69+int32(-5)))))
					v96 = v89
					if v96 == int32(0) {
					} else {
						v102 = int32(-1)&v96 + v70
						v106 = v41>>(uint(int32(31))%32)&v96 + v41
						v109 = v102 - v106 + int32(1)
						switch v81 {
						default:
							v124 = int32(base.Ui32(v79) >> (uint(int32(3)) % 32))
						case 1:
							v114 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v69+int32(-3)))))
							v124 = v114
						case 2:
							v117 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v69+int32(-5)))))
							v124 = v117
						case 3:
							v120 = *(*int32)(unsafe.Add(mBase, uint32(v69+int32(-9))))
							v124 = v120
						case 4:
							v123 = *(*int32)(unsafe.Add(mBase, uint32(v69+int32(-17))))
							v124 = v123
						}
						v125 = int32(0)
						v127 = base.B2i32(base.Ui32(v106) < base.Ui32(v124))
						if base.Ui32(v106) < base.Ui32(v124) {
							v128 = v106
						} else {
							v128 = v125
						}
						v129 = v124 - v128
						if base.Ui32(v109) < base.Ui32(v129) {
							v131 = v109
						} else {
							v131 = v129
						}
						if v102 < v106 {
							v133 = v125
						} else {
							v133 = v131
						}
						if base.Ui32(v106) < base.Ui32(v124) {
							v135 = v133
						} else {
							v135 = int32(0)
						}
						if v135 == int32(0) {
						} else {
							v139 = F_memmove(m, v69, v69+v128, v135)
							mBase = m.M
						}
						v141 = int32(0)
						*(*uint8)(unsafe.Add(mBase, uint32(v69+v135))) = uint8(v141)
						switch v81 {
						default:
							v144 = v135 << (uint(int32(3)) % 32)
							*(*uint8)(unsafe.Add(mBase, uint32(v78))) = uint8(v144)
						case 1:
							*(*uint8)(unsafe.Add(mBase, uint32(v69+int32(-3)))) = uint8(v135)
						case 2:
							*(*uint16)(unsafe.Add(mBase, uint32(v69+int32(-5)))) = uint16(v135)
						case 3:
							*(*int32)(unsafe.Add(mBase, uint32(v69+int32(-9)))) = v135
						case 4:
							*(*int64)(unsafe.Add(mBase, uint32(v69+int32(-17)))) = base.I64_extend_i32_u(v135)
						}
					}
				case 3:
					v92 = *(*int32)(unsafe.Add(mBase, uint32(v69+int32(-9))))
					v96 = v92
					if v96 == int32(0) {
					} else {
						v102 = int32(-1)&v96 + v70
						v106 = v41>>(uint(int32(31))%32)&v96 + v41
						v109 = v102 - v106 + int32(1)
						switch v81 {
						default:
							v124 = int32(base.Ui32(v79) >> (uint(int32(3)) % 32))
						case 1:
							v114 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v69+int32(-3)))))
							v124 = v114
						case 2:
							v117 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v69+int32(-5)))))
							v124 = v117
						case 3:
							v120 = *(*int32)(unsafe.Add(mBase, uint32(v69+int32(-9))))
							v124 = v120
						case 4:
							v123 = *(*int32)(unsafe.Add(mBase, uint32(v69+int32(-17))))
							v124 = v123
						}
						v125 = int32(0)
						v127 = base.B2i32(base.Ui32(v106) < base.Ui32(v124))
						if base.Ui32(v106) < base.Ui32(v124) {
							v128 = v106
						} else {
							v128 = v125
						}
						v129 = v124 - v128
						if base.Ui32(v109) < base.Ui32(v129) {
							v131 = v109
						} else {
							v131 = v129
						}
						if v102 < v106 {
							v133 = v125
						} else {
							v133 = v131
						}
						if base.Ui32(v106) < base.Ui32(v124) {
							v135 = v133
						} else {
							v135 = int32(0)
						}
						if v135 == int32(0) {
						} else {
							v139 = F_memmove(m, v69, v69+v128, v135)
							mBase = m.M
						}
						v141 = int32(0)
						*(*uint8)(unsafe.Add(mBase, uint32(v69+v135))) = uint8(v141)
						switch v81 {
						default:
							v144 = v135 << (uint(int32(3)) % 32)
							*(*uint8)(unsafe.Add(mBase, uint32(v78))) = uint8(v144)
						case 1:
							*(*uint8)(unsafe.Add(mBase, uint32(v69+int32(-3)))) = uint8(v135)
						case 2:
							*(*uint16)(unsafe.Add(mBase, uint32(v69+int32(-5)))) = uint16(v135)
						case 3:
							*(*int32)(unsafe.Add(mBase, uint32(v69+int32(-9)))) = v135
						case 4:
							*(*int64)(unsafe.Add(mBase, uint32(v69+int32(-17)))) = base.I64_extend_i32_u(v135)
						}
					}
				case 4:
					v95 = *(*int32)(unsafe.Add(mBase, uint32(v69+int32(-17))))
					v96 = v95
					if v96 == int32(0) {
					} else {
						v102 = int32(-1)&v96 + v70
						v106 = v41>>(uint(int32(31))%32)&v96 + v41
						v109 = v102 - v106 + int32(1)
						switch v81 {
						default:
							v124 = int32(base.Ui32(v79) >> (uint(int32(3)) % 32))
						case 1:
							v114 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v69+int32(-3)))))
							v124 = v114
						case 2:
							v117 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v69+int32(-5)))))
							v124 = v117
						case 3:
							v120 = *(*int32)(unsafe.Add(mBase, uint32(v69+int32(-9))))
							v124 = v120
						case 4:
							v123 = *(*int32)(unsafe.Add(mBase, uint32(v69+int32(-17))))
							v124 = v123
						}
						v125 = int32(0)
						v127 = base.B2i32(base.Ui32(v106) < base.Ui32(v124))
						if base.Ui32(v106) < base.Ui32(v124) {
							v128 = v106
						} else {
							v128 = v125
						}
						v129 = v124 - v128
						if base.Ui32(v109) < base.Ui32(v129) {
							v131 = v109
						} else {
							v131 = v129
						}
						if v102 < v106 {
							v133 = v125
						} else {
							v133 = v131
						}
						if base.Ui32(v106) < base.Ui32(v124) {
							v135 = v133
						} else {
							v135 = int32(0)
						}
						if v135 == int32(0) {
						} else {
							v139 = F_memmove(m, v69, v69+v128, v135)
							mBase = m.M
						}
						v141 = int32(0)
						*(*uint8)(unsafe.Add(mBase, uint32(v69+v135))) = uint8(v141)
						switch v81 {
						default:
							v144 = v135 << (uint(int32(3)) % 32)
							*(*uint8)(unsafe.Add(mBase, uint32(v78))) = uint8(v144)
						case 1:
							*(*uint8)(unsafe.Add(mBase, uint32(v69+int32(-3)))) = uint8(v135)
						case 2:
							*(*uint16)(unsafe.Add(mBase, uint32(v69+int32(-5)))) = uint16(v135)
						case 3:
							*(*int32)(unsafe.Add(mBase, uint32(v69+int32(-9)))) = v135
						case 4:
							*(*int64)(unsafe.Add(mBase, uint32(v69+int32(-17)))) = base.I64_extend_i32_u(v135)
						}
					}
				default:
				}
				v171 = *(*int32)(unsafe.Add(mBase, uint32(v13)+104))
				v172 = *(*int32)(unsafe.Add(mBase, uint32(v171)+32))
				v175 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v172+int32(-1)))))
				switch v175 & int32(7) {
				case 0:
					v192 = int32(base.Ui32(v175) >> (uint(int32(3)) % 32))
					if v192 != 0 {
						m.G0 = v11 + int32(16448)
						return
					} else {
						F_sdsfree(m, v172)
						mBase = m.M
						v195 = m.ExcPending
						if v195 != 0 {
							return
						} else {
							v196 = *(*int32)(unsafe.Add(mBase, uint32(v13)+104))
							*(*int32)(unsafe.Add(mBase, uint32(v196)+32)) = int32(0)
							v199 = v196
							v202 = *(*int32)(unsafe.Add(mBase, uint32(v199)+8))
							v203 = *(*int64)(unsafe.Add(mBase, uint32(v199)+16))
							v204 = int32(0)
							v205 = F___lseek(m, v202, v203, v204)
							mBase = m.M
							v206 = *(*int32)(unsafe.Add(mBase, uint32(v13)+104))
							v207 = *(*int32)(unsafe.Add(mBase, uint32(v206)+8))
							v211 = F_read(m, v207, v11+int32(48), int32(16384))
							mBase = m.M
							if v204 < v211 {
								v233 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
								v234 = *(*int32)(unsafe.Add(mBase, uint32(v233)+68))
								v235 = m.T0[v234].(func(*base.Module, int32, int32, int32) int32)(m, l0, v11+int32(48), v211)
								mBase = m.M
								v236 = m.ExcPending
								if v236 != 0 {
									return
								} else {
									if v235 != int32(-1) {
										v259 = *(*int32)(unsafe.Add(mBase, uint32(v13)+104))
										v260 = *(*int64)(unsafe.Add(mBase, uint32(v259)+16))
										v261 = base.I64_extend_i32_s(v235)
										v262 = v260 + v261
										*(*int64)(unsafe.Add(mBase, uint32(v259)+16)) = v262
										v264 = int32(_a_F_sendBulkToReplica_0)
										v266 = *(*int64)(unsafe.Add(mBase, _c_F_sendBulkToReplica[0]))
										*(*int64)(unsafe.Add(mBase, _c_F_sendBulkToReplica[0])) = v266 + v261
										v269 = *(*int64)(unsafe.Add(mBase, uint32(v259)+24))
										if v262 != v269 {
											m.G0 = v11 + int32(16448)
											return
										} else {
											F_closeRepldbfd(m, v13)
											mBase = m.M
											v272 = m.ExcPending
											if v272 != 0 {
												return
											} else {
												v273 = *(*int32)(unsafe.Add(mBase, uint32(v13)+8))
												v274 = int32(0)
												v276 = *(*int32)(unsafe.Add(mBase, uint32(v273)))
												v277 = *(*int32)(unsafe.Add(mBase, uint32(v276)+80))
												v278 = m.T0[v277].(func(*base.Module, int32, int32, int32) int32)(m, v273, v274, v274)
												mBase = m.M
												v279 = m.ExcPending
												if v279 != 0 {
													return
												} else {
													v280 = F_replicaPutOnline(m, v13)
													mBase = m.M
													v281 = m.ExcPending
													if v281 != 0 {
														return
													} else {
														if v280 != 0 {
															v284 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+205)))
															if v284&int32(32) != 0 {
																F__serverAssert(m, int32(_a_F_sendBulkToReplica_1), int32(_a_F_sendBulkToReplica_2), int32(1617))
																mBase = m.M
																v341 = m.ExcPending
																if v341 != 0 {
																	return
																} else {
																	F_abort(m)
																	mBase = m.M
																	base.Wasm_trap_unreachable()
																	for {
																	}
																}
															} else {
																v287 = *(*int32)(unsafe.Add(mBase, uint32(v13)+104))
																*(*int32)(unsafe.Add(mBase, uint32(v287)+4)) = int32(0)
																v291 = *(*int64)(unsafe.Add(mBase, _c_F_sendBulkToReplica[1]))
																if v291 != int64(0) {
																	v305 = *(*int32)(unsafe.Add(mBase, uint32(v13)+200))
																	if v305&int32(4194304) != 0 {
																	} else {
																		v308 = *(*int32)(unsafe.Add(mBase, uint32(v13)+104))
																		if v308 == int32(0) {
																			v315 = F_clusterSlotMigrationShouldInstallWriteHandler(m, v13)
																			mBase = m.M
																			if v315 == int32(0) {
																			} else {
																				v318 = *(*int32)(unsafe.Add(mBase, uint32(v13)+200))
																				*(*int32)(unsafe.Add(mBase, uint32(v13)+200)) = v318 | int32(4194304)
																				v323 = *(*int32)(unsafe.Add(mBase, _c_F_sendBulkToReplica[2]))
																				F_listLinkNodeHead(m, v323, v13+int32(168))
																				mBase = m.M
																			}
																		} else {
																			v311 = *(*int32)(unsafe.Add(mBase, uint32(v308)))
																			switch v311 {
																			case 0:
																				v315 = F_clusterSlotMigrationShouldInstallWriteHandler(m, v13)
																				mBase = m.M
																				if v315 == int32(0) {
																				} else {
																					v318 = *(*int32)(unsafe.Add(mBase, uint32(v13)+200))
																					*(*int32)(unsafe.Add(mBase, uint32(v13)+200)) = v318 | int32(4194304)
																					v323 = *(*int32)(unsafe.Add(mBase, _c_F_sendBulkToReplica[2]))
																					F_listLinkNodeHead(m, v323, v13+int32(168))
																					mBase = m.M
																				}
																			default:
																			case 9, 11:
																				if v305&int32(1024) != 0 {
																				} else {
																					v314 = *(*int32)(unsafe.Add(mBase, uint32(v308)+4))
																					if v314 != 0 {
																					} else {
																						v315 = F_clusterSlotMigrationShouldInstallWriteHandler(m, v13)
																						mBase = m.M
																						if v315 == int32(0) {
																						} else {
																							v318 = *(*int32)(unsafe.Add(mBase, uint32(v13)+200))
																							*(*int32)(unsafe.Add(mBase, uint32(v13)+200)) = v318 | int32(4194304)
																							v323 = *(*int32)(unsafe.Add(mBase, _c_F_sendBulkToReplica[2]))
																							F_listLinkNodeHead(m, v323, v13+int32(168))
																							mBase = m.M
																						}
																					}
																				}
																			}
																		}
																	}
																	m.G0 = v11 + int32(16448)
																	return
																} else {
																	v295 = *(*int32)(unsafe.Add(mBase, _c_F_sendBulkToReplica[3]))
																	*(*int32)(unsafe.Add(mBase, uint32(v11)+uint32(_c_F_sendBulkToReplica[4]))) = v295
																	F_replicationFeedReplicas(m, int32(-1), v11+int32(16444), int32(1))
																	mBase = m.M
																	v302 = m.ExcPending
																	if v302 != 0 {
																		return
																	} else {
																		v305 = *(*int32)(unsafe.Add(mBase, uint32(v13)+200))
																		if v305&int32(4194304) != 0 {
																		} else {
																			v308 = *(*int32)(unsafe.Add(mBase, uint32(v13)+104))
																			if v308 == int32(0) {
																				v315 = F_clusterSlotMigrationShouldInstallWriteHandler(m, v13)
																				mBase = m.M
																				if v315 == int32(0) {
																				} else {
																					v318 = *(*int32)(unsafe.Add(mBase, uint32(v13)+200))
																					*(*int32)(unsafe.Add(mBase, uint32(v13)+200)) = v318 | int32(4194304)
																					v323 = *(*int32)(unsafe.Add(mBase, _c_F_sendBulkToReplica[2]))
																					F_listLinkNodeHead(m, v323, v13+int32(168))
																					mBase = m.M
																				}
																			} else {
																				v311 = *(*int32)(unsafe.Add(mBase, uint32(v308)))
																				switch v311 {
																				case 0:
																					v315 = F_clusterSlotMigrationShouldInstallWriteHandler(m, v13)
																					mBase = m.M
																					if v315 == int32(0) {
																					} else {
																						v318 = *(*int32)(unsafe.Add(mBase, uint32(v13)+200))
																						*(*int32)(unsafe.Add(mBase, uint32(v13)+200)) = v318 | int32(4194304)
																						v323 = *(*int32)(unsafe.Add(mBase, _c_F_sendBulkToReplica[2]))
																						F_listLinkNodeHead(m, v323, v13+int32(168))
																						mBase = m.M
																					}
																				default:
																				case 9, 11:
																					if v305&int32(1024) != 0 {
																					} else {
																						v314 = *(*int32)(unsafe.Add(mBase, uint32(v308)+4))
																						if v314 != 0 {
																						} else {
																							v315 = F_clusterSlotMigrationShouldInstallWriteHandler(m, v13)
																							mBase = m.M
																							if v315 == int32(0) {
																							} else {
																								v318 = *(*int32)(unsafe.Add(mBase, uint32(v13)+200))
																								*(*int32)(unsafe.Add(mBase, uint32(v13)+200)) = v318 | int32(4194304)
																								v323 = *(*int32)(unsafe.Add(mBase, _c_F_sendBulkToReplica[2]))
																								F_listLinkNodeHead(m, v323, v13+int32(168))
																								mBase = m.M
																							}
																						}
																					}
																				}
																			}
																		}
																		m.G0 = v11 + int32(16448)
																		return
																	}
																}
															}
														} else {
															v282 = F_freeClient(m, v13)
															mBase = m.M
															v283 = m.ExcPending
															if v283 != 0 {
																return
															} else {
																m.G0 = v11 + int32(16448)
																return
															}
														}
													}
												}
											}
										}
									} else {
										v239 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
										if v239 == int32(3) {
											m.G0 = v11 + int32(16448)
											return
										} else {
											v243 = *(*int32)(unsafe.Add(mBase, _c_F_sendBulkToReplica[5]))
											if int32(3) < v243 {
												v257 = F_freeClient(m, v13)
												mBase = m.M
												v258 = m.ExcPending
												if v258 != 0 {
													return
												} else {
													m.G0 = v11 + int32(16448)
													return
												}
											} else {
												v246 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
												v247 = *(*int32)(unsafe.Add(mBase, uint32(v246)+88))
												v248 = m.T0[v247].(func(*base.Module, int32) int32)(m, l0)
												mBase = m.M
												v249 = m.ExcPending
												if v249 != 0 {
													return
												} else {
													*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = v248
													F__serverLog(m, int32(3), int32(_a_F_sendBulkToReplica_3), v11+int32(16))
													mBase = m.M
													v256 = m.ExcPending
													if v256 != 0 {
														return
													} else {
														v257 = F_freeClient(m, v13)
														mBase = m.M
														v258 = m.ExcPending
														if v258 != 0 {
															return
														} else {
															m.G0 = v11 + int32(16448)
															return
														}
													}
												}
											}
										}
									}
								}
							} else {
								v215 = *(*int32)(unsafe.Add(mBase, _c_F_sendBulkToReplica[5]))
								if int32(3) < v215 {
									v229 = F_freeClient(m, v13)
									mBase = m.M
									v230 = m.ExcPending
									if v230 != 0 {
										return
									} else {
										m.G0 = v11 + int32(16448)
										return
									}
								} else {
									if v211 != 0 {
										v220 = *(*int32)(unsafe.Add(mBase, _c_F_sendBulkToReplica[6]))
										v221 = F___strerror_l(m, v220, v220)
										mBase = m.M
										v222 = v221
									} else {
										v222 = int32(_a_F_sendBulkToReplica_4)
									}
									*(*int32)(unsafe.Add(mBase, uint32(v11))) = v222
									F__serverLog(m, int32(3), int32(_a_F_sendBulkToReplica_5), v11)
									mBase = m.M
									v227 = m.ExcPending
									if v227 != 0 {
										return
									} else {
										v229 = F_freeClient(m, v13)
										mBase = m.M
										v230 = m.ExcPending
										if v230 != 0 {
											return
										} else {
											m.G0 = v11 + int32(16448)
											return
										}
									}
								}
							}
						}
					}
				case 1:
					v182 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v172+int32(-3)))))
					v192 = v182
					if v192 != 0 {
						m.G0 = v11 + int32(16448)
						return
					} else {
						F_sdsfree(m, v172)
						mBase = m.M
						v195 = m.ExcPending
						if v195 != 0 {
							return
						} else {
							v196 = *(*int32)(unsafe.Add(mBase, uint32(v13)+104))
							*(*int32)(unsafe.Add(mBase, uint32(v196)+32)) = int32(0)
							v199 = v196
							v202 = *(*int32)(unsafe.Add(mBase, uint32(v199)+8))
							v203 = *(*int64)(unsafe.Add(mBase, uint32(v199)+16))
							v204 = int32(0)
							v205 = F___lseek(m, v202, v203, v204)
							mBase = m.M
							v206 = *(*int32)(unsafe.Add(mBase, uint32(v13)+104))
							v207 = *(*int32)(unsafe.Add(mBase, uint32(v206)+8))
							v211 = F_read(m, v207, v11+int32(48), int32(16384))
							mBase = m.M
							if v204 < v211 {
								v233 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
								v234 = *(*int32)(unsafe.Add(mBase, uint32(v233)+68))
								v235 = m.T0[v234].(func(*base.Module, int32, int32, int32) int32)(m, l0, v11+int32(48), v211)
								mBase = m.M
								v236 = m.ExcPending
								if v236 != 0 {
									return
								} else {
									if v235 != int32(-1) {
										v259 = *(*int32)(unsafe.Add(mBase, uint32(v13)+104))
										v260 = *(*int64)(unsafe.Add(mBase, uint32(v259)+16))
										v261 = base.I64_extend_i32_s(v235)
										v262 = v260 + v261
										*(*int64)(unsafe.Add(mBase, uint32(v259)+16)) = v262
										v264 = int32(_a_F_sendBulkToReplica_0)
										v266 = *(*int64)(unsafe.Add(mBase, _c_F_sendBulkToReplica[0]))
										*(*int64)(unsafe.Add(mBase, _c_F_sendBulkToReplica[0])) = v266 + v261
										v269 = *(*int64)(unsafe.Add(mBase, uint32(v259)+24))
										if v262 != v269 {
											m.G0 = v11 + int32(16448)
											return
										} else {
											F_closeRepldbfd(m, v13)
											mBase = m.M
											v272 = m.ExcPending
											if v272 != 0 {
												return
											} else {
												v273 = *(*int32)(unsafe.Add(mBase, uint32(v13)+8))
												v274 = int32(0)
												v276 = *(*int32)(unsafe.Add(mBase, uint32(v273)))
												v277 = *(*int32)(unsafe.Add(mBase, uint32(v276)+80))
												v278 = m.T0[v277].(func(*base.Module, int32, int32, int32) int32)(m, v273, v274, v274)
												mBase = m.M
												v279 = m.ExcPending
												if v279 != 0 {
													return
												} else {
													v280 = F_replicaPutOnline(m, v13)
													mBase = m.M
													v281 = m.ExcPending
													if v281 != 0 {
														return
													} else {
														if v280 != 0 {
															v284 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+205)))
															if v284&int32(32) != 0 {
																F__serverAssert(m, int32(_a_F_sendBulkToReplica_1), int32(_a_F_sendBulkToReplica_2), int32(1617))
																mBase = m.M
																v341 = m.ExcPending
																if v341 != 0 {
																	return
																} else {
																	F_abort(m)
																	mBase = m.M
																	base.Wasm_trap_unreachable()
																	for {
																	}
																}
															} else {
																v287 = *(*int32)(unsafe.Add(mBase, uint32(v13)+104))
																*(*int32)(unsafe.Add(mBase, uint32(v287)+4)) = int32(0)
																v291 = *(*int64)(unsafe.Add(mBase, _c_F_sendBulkToReplica[1]))
																if v291 != int64(0) {
																	v305 = *(*int32)(unsafe.Add(mBase, uint32(v13)+200))
																	if v305&int32(4194304) != 0 {
																	} else {
																		v308 = *(*int32)(unsafe.Add(mBase, uint32(v13)+104))
																		if v308 == int32(0) {
																			v315 = F_clusterSlotMigrationShouldInstallWriteHandler(m, v13)
																			mBase = m.M
																			if v315 == int32(0) {
																			} else {
																				v318 = *(*int32)(unsafe.Add(mBase, uint32(v13)+200))
																				*(*int32)(unsafe.Add(mBase, uint32(v13)+200)) = v318 | int32(4194304)
																				v323 = *(*int32)(unsafe.Add(mBase, _c_F_sendBulkToReplica[2]))
																				F_listLinkNodeHead(m, v323, v13+int32(168))
																				mBase = m.M
																			}
																		} else {
																			v311 = *(*int32)(unsafe.Add(mBase, uint32(v308)))
																			switch v311 {
																			case 0:
																				v315 = F_clusterSlotMigrationShouldInstallWriteHandler(m, v13)
																				mBase = m.M
																				if v315 == int32(0) {
																				} else {
																					v318 = *(*int32)(unsafe.Add(mBase, uint32(v13)+200))
																					*(*int32)(unsafe.Add(mBase, uint32(v13)+200)) = v318 | int32(4194304)
																					v323 = *(*int32)(unsafe.Add(mBase, _c_F_sendBulkToReplica[2]))
																					F_listLinkNodeHead(m, v323, v13+int32(168))
																					mBase = m.M
																				}
																			default:
																			case 9, 11:
																				if v305&int32(1024) != 0 {
																				} else {
																					v314 = *(*int32)(unsafe.Add(mBase, uint32(v308)+4))
																					if v314 != 0 {
																					} else {
																						v315 = F_clusterSlotMigrationShouldInstallWriteHandler(m, v13)
																						mBase = m.M
																						if v315 == int32(0) {
																						} else {
																							v318 = *(*int32)(unsafe.Add(mBase, uint32(v13)+200))
																							*(*int32)(unsafe.Add(mBase, uint32(v13)+200)) = v318 | int32(4194304)
																							v323 = *(*int32)(unsafe.Add(mBase, _c_F_sendBulkToReplica[2]))
																							F_listLinkNodeHead(m, v323, v13+int32(168))
																							mBase = m.M
																						}
																					}
																				}
																			}
																		}
																	}
																	m.G0 = v11 + int32(16448)
																	return
																} else {
																	v295 = *(*int32)(unsafe.Add(mBase, _c_F_sendBulkToReplica[3]))
																	*(*int32)(unsafe.Add(mBase, uint32(v11)+uint32(_c_F_sendBulkToReplica[4]))) = v295
																	F_replicationFeedReplicas(m, int32(-1), v11+int32(16444), int32(1))
																	mBase = m.M
																	v302 = m.ExcPending
																	if v302 != 0 {
																		return
																	} else {
																		v305 = *(*int32)(unsafe.Add(mBase, uint32(v13)+200))
																		if v305&int32(4194304) != 0 {
																		} else {
																			v308 = *(*int32)(unsafe.Add(mBase, uint32(v13)+104))
																			if v308 == int32(0) {
																				v315 = F_clusterSlotMigrationShouldInstallWriteHandler(m, v13)
																				mBase = m.M
																				if v315 == int32(0) {
																				} else {
																					v318 = *(*int32)(unsafe.Add(mBase, uint32(v13)+200))
																					*(*int32)(unsafe.Add(mBase, uint32(v13)+200)) = v318 | int32(4194304)
																					v323 = *(*int32)(unsafe.Add(mBase, _c_F_sendBulkToReplica[2]))
																					F_listLinkNodeHead(m, v323, v13+int32(168))
																					mBase = m.M
																				}
																			} else {
																				v311 = *(*int32)(unsafe.Add(mBase, uint32(v308)))
																				switch v311 {
																				case 0:
																					v315 = F_clusterSlotMigrationShouldInstallWriteHandler(m, v13)
																					mBase = m.M
																					if v315 == int32(0) {
																					} else {
																						v318 = *(*int32)(unsafe.Add(mBase, uint32(v13)+200))
																						*(*int32)(unsafe.Add(mBase, uint32(v13)+200)) = v318 | int32(4194304)
																						v323 = *(*int32)(unsafe.Add(mBase, _c_F_sendBulkToReplica[2]))
																						F_listLinkNodeHead(m, v323, v13+int32(168))
																						mBase = m.M
																					}
																				default:
																				case 9, 11:
																					if v305&int32(1024) != 0 {
																					} else {
																						v314 = *(*int32)(unsafe.Add(mBase, uint32(v308)+4))
																						if v314 != 0 {
																						} else {
																							v315 = F_clusterSlotMigrationShouldInstallWriteHandler(m, v13)
																							mBase = m.M
																							if v315 == int32(0) {
																							} else {
																								v318 = *(*int32)(unsafe.Add(mBase, uint32(v13)+200))
																								*(*int32)(unsafe.Add(mBase, uint32(v13)+200)) = v318 | int32(4194304)
																								v323 = *(*int32)(unsafe.Add(mBase, _c_F_sendBulkToReplica[2]))
																								F_listLinkNodeHead(m, v323, v13+int32(168))
																								mBase = m.M
																							}
																						}
																					}
																				}
																			}
																		}
																		m.G0 = v11 + int32(16448)
																		return
																	}
																}
															}
														} else {
															v282 = F_freeClient(m, v13)
															mBase = m.M
															v283 = m.ExcPending
															if v283 != 0 {
																return
															} else {
																m.G0 = v11 + int32(16448)
																return
															}
														}
													}
												}
											}
										}
									} else {
										v239 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
										if v239 == int32(3) {
											m.G0 = v11 + int32(16448)
											return
										} else {
											v243 = *(*int32)(unsafe.Add(mBase, _c_F_sendBulkToReplica[5]))
											if int32(3) < v243 {
												v257 = F_freeClient(m, v13)
												mBase = m.M
												v258 = m.ExcPending
												if v258 != 0 {
													return
												} else {
													m.G0 = v11 + int32(16448)
													return
												}
											} else {
												v246 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
												v247 = *(*int32)(unsafe.Add(mBase, uint32(v246)+88))
												v248 = m.T0[v247].(func(*base.Module, int32) int32)(m, l0)
												mBase = m.M
												v249 = m.ExcPending
												if v249 != 0 {
													return
												} else {
													*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = v248
													F__serverLog(m, int32(3), int32(_a_F_sendBulkToReplica_3), v11+int32(16))
													mBase = m.M
													v256 = m.ExcPending
													if v256 != 0 {
														return
													} else {
														v257 = F_freeClient(m, v13)
														mBase = m.M
														v258 = m.ExcPending
														if v258 != 0 {
															return
														} else {
															m.G0 = v11 + int32(16448)
															return
														}
													}
												}
											}
										}
									}
								}
							} else {
								v215 = *(*int32)(unsafe.Add(mBase, _c_F_sendBulkToReplica[5]))
								if int32(3) < v215 {
									v229 = F_freeClient(m, v13)
									mBase = m.M
									v230 = m.ExcPending
									if v230 != 0 {
										return
									} else {
										m.G0 = v11 + int32(16448)
										return
									}
								} else {
									if v211 != 0 {
										v220 = *(*int32)(unsafe.Add(mBase, _c_F_sendBulkToReplica[6]))
										v221 = F___strerror_l(m, v220, v220)
										mBase = m.M
										v222 = v221
									} else {
										v222 = int32(_a_F_sendBulkToReplica_4)
									}
									*(*int32)(unsafe.Add(mBase, uint32(v11))) = v222
									F__serverLog(m, int32(3), int32(_a_F_sendBulkToReplica_5), v11)
									mBase = m.M
									v227 = m.ExcPending
									if v227 != 0 {
										return
									} else {
										v229 = F_freeClient(m, v13)
										mBase = m.M
										v230 = m.ExcPending
										if v230 != 0 {
											return
										} else {
											m.G0 = v11 + int32(16448)
											return
										}
									}
								}
							}
						}
					}
				case 2:
					v185 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v172+int32(-5)))))
					v192 = v185
					if v192 != 0 {
						m.G0 = v11 + int32(16448)
						return
					} else {
						F_sdsfree(m, v172)
						mBase = m.M
						v195 = m.ExcPending
						if v195 != 0 {
							return
						} else {
							v196 = *(*int32)(unsafe.Add(mBase, uint32(v13)+104))
							*(*int32)(unsafe.Add(mBase, uint32(v196)+32)) = int32(0)
							v199 = v196
							v202 = *(*int32)(unsafe.Add(mBase, uint32(v199)+8))
							v203 = *(*int64)(unsafe.Add(mBase, uint32(v199)+16))
							v204 = int32(0)
							v205 = F___lseek(m, v202, v203, v204)
							mBase = m.M
							v206 = *(*int32)(unsafe.Add(mBase, uint32(v13)+104))
							v207 = *(*int32)(unsafe.Add(mBase, uint32(v206)+8))
							v211 = F_read(m, v207, v11+int32(48), int32(16384))
							mBase = m.M
							if v204 < v211 {
								v233 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
								v234 = *(*int32)(unsafe.Add(mBase, uint32(v233)+68))
								v235 = m.T0[v234].(func(*base.Module, int32, int32, int32) int32)(m, l0, v11+int32(48), v211)
								mBase = m.M
								v236 = m.ExcPending
								if v236 != 0 {
									return
								} else {
									if v235 != int32(-1) {
										v259 = *(*int32)(unsafe.Add(mBase, uint32(v13)+104))
										v260 = *(*int64)(unsafe.Add(mBase, uint32(v259)+16))
										v261 = base.I64_extend_i32_s(v235)
										v262 = v260 + v261
										*(*int64)(unsafe.Add(mBase, uint32(v259)+16)) = v262
										v264 = int32(_a_F_sendBulkToReplica_0)
										v266 = *(*int64)(unsafe.Add(mBase, _c_F_sendBulkToReplica[0]))
										*(*int64)(unsafe.Add(mBase, _c_F_sendBulkToReplica[0])) = v266 + v261
										v269 = *(*int64)(unsafe.Add(mBase, uint32(v259)+24))
										if v262 != v269 {
											m.G0 = v11 + int32(16448)
											return
										} else {
											F_closeRepldbfd(m, v13)
											mBase = m.M
											v272 = m.ExcPending
											if v272 != 0 {
												return
											} else {
												v273 = *(*int32)(unsafe.Add(mBase, uint32(v13)+8))
												v274 = int32(0)
												v276 = *(*int32)(unsafe.Add(mBase, uint32(v273)))
												v277 = *(*int32)(unsafe.Add(mBase, uint32(v276)+80))
												v278 = m.T0[v277].(func(*base.Module, int32, int32, int32) int32)(m, v273, v274, v274)
												mBase = m.M
												v279 = m.ExcPending
												if v279 != 0 {
													return
												} else {
													v280 = F_replicaPutOnline(m, v13)
													mBase = m.M
													v281 = m.ExcPending
													if v281 != 0 {
														return
													} else {
														if v280 != 0 {
															v284 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+205)))
															if v284&int32(32) != 0 {
																F__serverAssert(m, int32(_a_F_sendBulkToReplica_1), int32(_a_F_sendBulkToReplica_2), int32(1617))
																mBase = m.M
																v341 = m.ExcPending
																if v341 != 0 {
																	return
																} else {
																	F_abort(m)
																	mBase = m.M
																	base.Wasm_trap_unreachable()
																	for {
																	}
																}
															} else {
																v287 = *(*int32)(unsafe.Add(mBase, uint32(v13)+104))
																*(*int32)(unsafe.Add(mBase, uint32(v287)+4)) = int32(0)
																v291 = *(*int64)(unsafe.Add(mBase, _c_F_sendBulkToReplica[1]))
																if v291 != int64(0) {
																	v305 = *(*int32)(unsafe.Add(mBase, uint32(v13)+200))
																	if v305&int32(4194304) != 0 {
																	} else {
																		v308 = *(*int32)(unsafe.Add(mBase, uint32(v13)+104))
																		if v308 == int32(0) {
																			v315 = F_clusterSlotMigrationShouldInstallWriteHandler(m, v13)
																			mBase = m.M
																			if v315 == int32(0) {
																			} else {
																				v318 = *(*int32)(unsafe.Add(mBase, uint32(v13)+200))
																				*(*int32)(unsafe.Add(mBase, uint32(v13)+200)) = v318 | int32(4194304)
																				v323 = *(*int32)(unsafe.Add(mBase, _c_F_sendBulkToReplica[2]))
																				F_listLinkNodeHead(m, v323, v13+int32(168))
																				mBase = m.M
																			}
																		} else {
																			v311 = *(*int32)(unsafe.Add(mBase, uint32(v308)))
																			switch v311 {
																			case 0:
																				v315 = F_clusterSlotMigrationShouldInstallWriteHandler(m, v13)
																				mBase = m.M
																				if v315 == int32(0) {
																				} else {
																					v318 = *(*int32)(unsafe.Add(mBase, uint32(v13)+200))
																					*(*int32)(unsafe.Add(mBase, uint32(v13)+200)) = v318 | int32(4194304)
																					v323 = *(*int32)(unsafe.Add(mBase, _c_F_sendBulkToReplica[2]))
																					F_listLinkNodeHead(m, v323, v13+int32(168))
																					mBase = m.M
																				}
																			default:
																			case 9, 11:
																				if v305&int32(1024) != 0 {
																				} else {
																					v314 = *(*int32)(unsafe.Add(mBase, uint32(v308)+4))
																					if v314 != 0 {
																					} else {
																						v315 = F_clusterSlotMigrationShouldInstallWriteHandler(m, v13)
																						mBase = m.M
																						if v315 == int32(0) {
																						} else {
																							v318 = *(*int32)(unsafe.Add(mBase, uint32(v13)+200))
																							*(*int32)(unsafe.Add(mBase, uint32(v13)+200)) = v318 | int32(4194304)
																							v323 = *(*int32)(unsafe.Add(mBase, _c_F_sendBulkToReplica[2]))
																							F_listLinkNodeHead(m, v323, v13+int32(168))
																							mBase = m.M
																						}
																					}
																				}
																			}
																		}
																	}
																	m.G0 = v11 + int32(16448)
																	return
																} else {
																	v295 = *(*int32)(unsafe.Add(mBase, _c_F_sendBulkToReplica[3]))
																	*(*int32)(unsafe.Add(mBase, uint32(v11)+uint32(_c_F_sendBulkToReplica[4]))) = v295
																	F_replicationFeedReplicas(m, int32(-1), v11+int32(16444), int32(1))
																	mBase = m.M
																	v302 = m.ExcPending
																	if v302 != 0 {
																		return
																	} else {
																		v305 = *(*int32)(unsafe.Add(mBase, uint32(v13)+200))
																		if v305&int32(4194304) != 0 {
																		} else {
																			v308 = *(*int32)(unsafe.Add(mBase, uint32(v13)+104))
																			if v308 == int32(0) {
																				v315 = F_clusterSlotMigrationShouldInstallWriteHandler(m, v13)
																				mBase = m.M
																				if v315 == int32(0) {
																				} else {
																					v318 = *(*int32)(unsafe.Add(mBase, uint32(v13)+200))
																					*(*int32)(unsafe.Add(mBase, uint32(v13)+200)) = v318 | int32(4194304)
																					v323 = *(*int32)(unsafe.Add(mBase, _c_F_sendBulkToReplica[2]))
																					F_listLinkNodeHead(m, v323, v13+int32(168))
																					mBase = m.M
																				}
																			} else {
																				v311 = *(*int32)(unsafe.Add(mBase, uint32(v308)))
																				switch v311 {
																				case 0:
																					v315 = F_clusterSlotMigrationShouldInstallWriteHandler(m, v13)
																					mBase = m.M
																					if v315 == int32(0) {
																					} else {
																						v318 = *(*int32)(unsafe.Add(mBase, uint32(v13)+200))
																						*(*int32)(unsafe.Add(mBase, uint32(v13)+200)) = v318 | int32(4194304)
																						v323 = *(*int32)(unsafe.Add(mBase, _c_F_sendBulkToReplica[2]))
																						F_listLinkNodeHead(m, v323, v13+int32(168))
																						mBase = m.M
																					}
																				default:
																				case 9, 11:
																					if v305&int32(1024) != 0 {
																					} else {
																						v314 = *(*int32)(unsafe.Add(mBase, uint32(v308)+4))
																						if v314 != 0 {
																						} else {
																							v315 = F_clusterSlotMigrationShouldInstallWriteHandler(m, v13)
																							mBase = m.M
																							if v315 == int32(0) {
																							} else {
																								v318 = *(*int32)(unsafe.Add(mBase, uint32(v13)+200))
																								*(*int32)(unsafe.Add(mBase, uint32(v13)+200)) = v318 | int32(4194304)
																								v323 = *(*int32)(unsafe.Add(mBase, _c_F_sendBulkToReplica[2]))
																								F_listLinkNodeHead(m, v323, v13+int32(168))
																								mBase = m.M
																							}
																						}
																					}
																				}
																			}
																		}
																		m.G0 = v11 + int32(16448)
																		return
																	}
																}
															}
														} else {
															v282 = F_freeClient(m, v13)
															mBase = m.M
															v283 = m.ExcPending
															if v283 != 0 {
																return
															} else {
																m.G0 = v11 + int32(16448)
																return
															}
														}
													}
												}
											}
										}
									} else {
										v239 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
										if v239 == int32(3) {
											m.G0 = v11 + int32(16448)
											return
										} else {
											v243 = *(*int32)(unsafe.Add(mBase, _c_F_sendBulkToReplica[5]))
											if int32(3) < v243 {
												v257 = F_freeClient(m, v13)
												mBase = m.M
												v258 = m.ExcPending
												if v258 != 0 {
													return
												} else {
													m.G0 = v11 + int32(16448)
													return
												}
											} else {
												v246 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
												v247 = *(*int32)(unsafe.Add(mBase, uint32(v246)+88))
												v248 = m.T0[v247].(func(*base.Module, int32) int32)(m, l0)
												mBase = m.M
												v249 = m.ExcPending
												if v249 != 0 {
													return
												} else {
													*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = v248
													F__serverLog(m, int32(3), int32(_a_F_sendBulkToReplica_3), v11+int32(16))
													mBase = m.M
													v256 = m.ExcPending
													if v256 != 0 {
														return
													} else {
														v257 = F_freeClient(m, v13)
														mBase = m.M
														v258 = m.ExcPending
														if v258 != 0 {
															return
														} else {
															m.G0 = v11 + int32(16448)
															return
														}
													}
												}
											}
										}
									}
								}
							} else {
								v215 = *(*int32)(unsafe.Add(mBase, _c_F_sendBulkToReplica[5]))
								if int32(3) < v215 {
									v229 = F_freeClient(m, v13)
									mBase = m.M
									v230 = m.ExcPending
									if v230 != 0 {
										return
									} else {
										m.G0 = v11 + int32(16448)
										return
									}
								} else {
									if v211 != 0 {
										v220 = *(*int32)(unsafe.Add(mBase, _c_F_sendBulkToReplica[6]))
										v221 = F___strerror_l(m, v220, v220)
										mBase = m.M
										v222 = v221
									} else {
										v222 = int32(_a_F_sendBulkToReplica_4)
									}
									*(*int32)(unsafe.Add(mBase, uint32(v11))) = v222
									F__serverLog(m, int32(3), int32(_a_F_sendBulkToReplica_5), v11)
									mBase = m.M
									v227 = m.ExcPending
									if v227 != 0 {
										return
									} else {
										v229 = F_freeClient(m, v13)
										mBase = m.M
										v230 = m.ExcPending
										if v230 != 0 {
											return
										} else {
											m.G0 = v11 + int32(16448)
											return
										}
									}
								}
							}
						}
					}
				case 3:
					v188 = *(*int32)(unsafe.Add(mBase, uint32(v172+int32(-9))))
					v192 = v188
					if v192 != 0 {
						m.G0 = v11 + int32(16448)
						return
					} else {
						F_sdsfree(m, v172)
						mBase = m.M
						v195 = m.ExcPending
						if v195 != 0 {
							return
						} else {
							v196 = *(*int32)(unsafe.Add(mBase, uint32(v13)+104))
							*(*int32)(unsafe.Add(mBase, uint32(v196)+32)) = int32(0)
							v199 = v196
							v202 = *(*int32)(unsafe.Add(mBase, uint32(v199)+8))
							v203 = *(*int64)(unsafe.Add(mBase, uint32(v199)+16))
							v204 = int32(0)
							v205 = F___lseek(m, v202, v203, v204)
							mBase = m.M
							v206 = *(*int32)(unsafe.Add(mBase, uint32(v13)+104))
							v207 = *(*int32)(unsafe.Add(mBase, uint32(v206)+8))
							v211 = F_read(m, v207, v11+int32(48), int32(16384))
							mBase = m.M
							if v204 < v211 {
								v233 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
								v234 = *(*int32)(unsafe.Add(mBase, uint32(v233)+68))
								v235 = m.T0[v234].(func(*base.Module, int32, int32, int32) int32)(m, l0, v11+int32(48), v211)
								mBase = m.M
								v236 = m.ExcPending
								if v236 != 0 {
									return
								} else {
									if v235 != int32(-1) {
										v259 = *(*int32)(unsafe.Add(mBase, uint32(v13)+104))
										v260 = *(*int64)(unsafe.Add(mBase, uint32(v259)+16))
										v261 = base.I64_extend_i32_s(v235)
										v262 = v260 + v261
										*(*int64)(unsafe.Add(mBase, uint32(v259)+16)) = v262
										v264 = int32(_a_F_sendBulkToReplica_0)
										v266 = *(*int64)(unsafe.Add(mBase, _c_F_sendBulkToReplica[0]))
										*(*int64)(unsafe.Add(mBase, _c_F_sendBulkToReplica[0])) = v266 + v261
										v269 = *(*int64)(unsafe.Add(mBase, uint32(v259)+24))
										if v262 != v269 {
											m.G0 = v11 + int32(16448)
											return
										} else {
											F_closeRepldbfd(m, v13)
											mBase = m.M
											v272 = m.ExcPending
											if v272 != 0 {
												return
											} else {
												v273 = *(*int32)(unsafe.Add(mBase, uint32(v13)+8))
												v274 = int32(0)
												v276 = *(*int32)(unsafe.Add(mBase, uint32(v273)))
												v277 = *(*int32)(unsafe.Add(mBase, uint32(v276)+80))
												v278 = m.T0[v277].(func(*base.Module, int32, int32, int32) int32)(m, v273, v274, v274)
												mBase = m.M
												v279 = m.ExcPending
												if v279 != 0 {
													return
												} else {
													v280 = F_replicaPutOnline(m, v13)
													mBase = m.M
													v281 = m.ExcPending
													if v281 != 0 {
														return
													} else {
														if v280 != 0 {
															v284 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+205)))
															if v284&int32(32) != 0 {
																F__serverAssert(m, int32(_a_F_sendBulkToReplica_1), int32(_a_F_sendBulkToReplica_2), int32(1617))
																mBase = m.M
																v341 = m.ExcPending
																if v341 != 0 {
																	return
																} else {
																	F_abort(m)
																	mBase = m.M
																	base.Wasm_trap_unreachable()
																	for {
																	}
																}
															} else {
																v287 = *(*int32)(unsafe.Add(mBase, uint32(v13)+104))
																*(*int32)(unsafe.Add(mBase, uint32(v287)+4)) = int32(0)
																v291 = *(*int64)(unsafe.Add(mBase, _c_F_sendBulkToReplica[1]))
																if v291 != int64(0) {
																	v305 = *(*int32)(unsafe.Add(mBase, uint32(v13)+200))
																	if v305&int32(4194304) != 0 {
																	} else {
																		v308 = *(*int32)(unsafe.Add(mBase, uint32(v13)+104))
																		if v308 == int32(0) {
																			v315 = F_clusterSlotMigrationShouldInstallWriteHandler(m, v13)
																			mBase = m.M
																			if v315 == int32(0) {
																			} else {
																				v318 = *(*int32)(unsafe.Add(mBase, uint32(v13)+200))
																				*(*int32)(unsafe.Add(mBase, uint32(v13)+200)) = v318 | int32(4194304)
																				v323 = *(*int32)(unsafe.Add(mBase, _c_F_sendBulkToReplica[2]))
																				F_listLinkNodeHead(m, v323, v13+int32(168))
																				mBase = m.M
																			}
																		} else {
																			v311 = *(*int32)(unsafe.Add(mBase, uint32(v308)))
																			switch v311 {
																			case 0:
																				v315 = F_clusterSlotMigrationShouldInstallWriteHandler(m, v13)
																				mBase = m.M
																				if v315 == int32(0) {
																				} else {
																					v318 = *(*int32)(unsafe.Add(mBase, uint32(v13)+200))
																					*(*int32)(unsafe.Add(mBase, uint32(v13)+200)) = v318 | int32(4194304)
																					v323 = *(*int32)(unsafe.Add(mBase, _c_F_sendBulkToReplica[2]))
																					F_listLinkNodeHead(m, v323, v13+int32(168))
																					mBase = m.M
																				}
																			default:
																			case 9, 11:
																				if v305&int32(1024) != 0 {
																				} else {
																					v314 = *(*int32)(unsafe.Add(mBase, uint32(v308)+4))
																					if v314 != 0 {
																					} else {
																						v315 = F_clusterSlotMigrationShouldInstallWriteHandler(m, v13)
																						mBase = m.M
																						if v315 == int32(0) {
																						} else {
																							v318 = *(*int32)(unsafe.Add(mBase, uint32(v13)+200))
																							*(*int32)(unsafe.Add(mBase, uint32(v13)+200)) = v318 | int32(4194304)
																							v323 = *(*int32)(unsafe.Add(mBase, _c_F_sendBulkToReplica[2]))
																							F_listLinkNodeHead(m, v323, v13+int32(168))
																							mBase = m.M
																						}
																					}
																				}
																			}
																		}
																	}
																	m.G0 = v11 + int32(16448)
																	return
																} else {
																	v295 = *(*int32)(unsafe.Add(mBase, _c_F_sendBulkToReplica[3]))
																	*(*int32)(unsafe.Add(mBase, uint32(v11)+uint32(_c_F_sendBulkToReplica[4]))) = v295
																	F_replicationFeedReplicas(m, int32(-1), v11+int32(16444), int32(1))
																	mBase = m.M
																	v302 = m.ExcPending
																	if v302 != 0 {
																		return
																	} else {
																		v305 = *(*int32)(unsafe.Add(mBase, uint32(v13)+200))
																		if v305&int32(4194304) != 0 {
																		} else {
																			v308 = *(*int32)(unsafe.Add(mBase, uint32(v13)+104))
																			if v308 == int32(0) {
																				v315 = F_clusterSlotMigrationShouldInstallWriteHandler(m, v13)
																				mBase = m.M
																				if v315 == int32(0) {
																				} else {
																					v318 = *(*int32)(unsafe.Add(mBase, uint32(v13)+200))
																					*(*int32)(unsafe.Add(mBase, uint32(v13)+200)) = v318 | int32(4194304)
																					v323 = *(*int32)(unsafe.Add(mBase, _c_F_sendBulkToReplica[2]))
																					F_listLinkNodeHead(m, v323, v13+int32(168))
																					mBase = m.M
																				}
																			} else {
																				v311 = *(*int32)(unsafe.Add(mBase, uint32(v308)))
																				switch v311 {
																				case 0:
																					v315 = F_clusterSlotMigrationShouldInstallWriteHandler(m, v13)
																					mBase = m.M
																					if v315 == int32(0) {
																					} else {
																						v318 = *(*int32)(unsafe.Add(mBase, uint32(v13)+200))
																						*(*int32)(unsafe.Add(mBase, uint32(v13)+200)) = v318 | int32(4194304)
																						v323 = *(*int32)(unsafe.Add(mBase, _c_F_sendBulkToReplica[2]))
																						F_listLinkNodeHead(m, v323, v13+int32(168))
																						mBase = m.M
																					}
																				default:
																				case 9, 11:
																					if v305&int32(1024) != 0 {
																					} else {
																						v314 = *(*int32)(unsafe.Add(mBase, uint32(v308)+4))
																						if v314 != 0 {
																						} else {
																							v315 = F_clusterSlotMigrationShouldInstallWriteHandler(m, v13)
																							mBase = m.M
																							if v315 == int32(0) {
																							} else {
																								v318 = *(*int32)(unsafe.Add(mBase, uint32(v13)+200))
																								*(*int32)(unsafe.Add(mBase, uint32(v13)+200)) = v318 | int32(4194304)
																								v323 = *(*int32)(unsafe.Add(mBase, _c_F_sendBulkToReplica[2]))
																								F_listLinkNodeHead(m, v323, v13+int32(168))
																								mBase = m.M
																							}
																						}
																					}
																				}
																			}
																		}
																		m.G0 = v11 + int32(16448)
																		return
																	}
																}
															}
														} else {
															v282 = F_freeClient(m, v13)
															mBase = m.M
															v283 = m.ExcPending
															if v283 != 0 {
																return
															} else {
																m.G0 = v11 + int32(16448)
																return
															}
														}
													}
												}
											}
										}
									} else {
										v239 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
										if v239 == int32(3) {
											m.G0 = v11 + int32(16448)
											return
										} else {
											v243 = *(*int32)(unsafe.Add(mBase, _c_F_sendBulkToReplica[5]))
											if int32(3) < v243 {
												v257 = F_freeClient(m, v13)
												mBase = m.M
												v258 = m.ExcPending
												if v258 != 0 {
													return
												} else {
													m.G0 = v11 + int32(16448)
													return
												}
											} else {
												v246 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
												v247 = *(*int32)(unsafe.Add(mBase, uint32(v246)+88))
												v248 = m.T0[v247].(func(*base.Module, int32) int32)(m, l0)
												mBase = m.M
												v249 = m.ExcPending
												if v249 != 0 {
													return
												} else {
													*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = v248
													F__serverLog(m, int32(3), int32(_a_F_sendBulkToReplica_3), v11+int32(16))
													mBase = m.M
													v256 = m.ExcPending
													if v256 != 0 {
														return
													} else {
														v257 = F_freeClient(m, v13)
														mBase = m.M
														v258 = m.ExcPending
														if v258 != 0 {
															return
														} else {
															m.G0 = v11 + int32(16448)
															return
														}
													}
												}
											}
										}
									}
								}
							} else {
								v215 = *(*int32)(unsafe.Add(mBase, _c_F_sendBulkToReplica[5]))
								if int32(3) < v215 {
									v229 = F_freeClient(m, v13)
									mBase = m.M
									v230 = m.ExcPending
									if v230 != 0 {
										return
									} else {
										m.G0 = v11 + int32(16448)
										return
									}
								} else {
									if v211 != 0 {
										v220 = *(*int32)(unsafe.Add(mBase, _c_F_sendBulkToReplica[6]))
										v221 = F___strerror_l(m, v220, v220)
										mBase = m.M
										v222 = v221
									} else {
										v222 = int32(_a_F_sendBulkToReplica_4)
									}
									*(*int32)(unsafe.Add(mBase, uint32(v11))) = v222
									F__serverLog(m, int32(3), int32(_a_F_sendBulkToReplica_5), v11)
									mBase = m.M
									v227 = m.ExcPending
									if v227 != 0 {
										return
									} else {
										v229 = F_freeClient(m, v13)
										mBase = m.M
										v230 = m.ExcPending
										if v230 != 0 {
											return
										} else {
											m.G0 = v11 + int32(16448)
											return
										}
									}
								}
							}
						}
					}
				case 4:
					v191 = *(*int32)(unsafe.Add(mBase, uint32(v172+int32(-17))))
					v192 = v191
					if v192 != 0 {
						m.G0 = v11 + int32(16448)
						return
					} else {
						F_sdsfree(m, v172)
						mBase = m.M
						v195 = m.ExcPending
						if v195 != 0 {
							return
						} else {
							v196 = *(*int32)(unsafe.Add(mBase, uint32(v13)+104))
							*(*int32)(unsafe.Add(mBase, uint32(v196)+32)) = int32(0)
							v199 = v196
							v202 = *(*int32)(unsafe.Add(mBase, uint32(v199)+8))
							v203 = *(*int64)(unsafe.Add(mBase, uint32(v199)+16))
							v204 = int32(0)
							v205 = F___lseek(m, v202, v203, v204)
							mBase = m.M
							v206 = *(*int32)(unsafe.Add(mBase, uint32(v13)+104))
							v207 = *(*int32)(unsafe.Add(mBase, uint32(v206)+8))
							v211 = F_read(m, v207, v11+int32(48), int32(16384))
							mBase = m.M
							if v204 < v211 {
								v233 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
								v234 = *(*int32)(unsafe.Add(mBase, uint32(v233)+68))
								v235 = m.T0[v234].(func(*base.Module, int32, int32, int32) int32)(m, l0, v11+int32(48), v211)
								mBase = m.M
								v236 = m.ExcPending
								if v236 != 0 {
									return
								} else {
									if v235 != int32(-1) {
										v259 = *(*int32)(unsafe.Add(mBase, uint32(v13)+104))
										v260 = *(*int64)(unsafe.Add(mBase, uint32(v259)+16))
										v261 = base.I64_extend_i32_s(v235)
										v262 = v260 + v261
										*(*int64)(unsafe.Add(mBase, uint32(v259)+16)) = v262
										v264 = int32(_a_F_sendBulkToReplica_0)
										v266 = *(*int64)(unsafe.Add(mBase, _c_F_sendBulkToReplica[0]))
										*(*int64)(unsafe.Add(mBase, _c_F_sendBulkToReplica[0])) = v266 + v261
										v269 = *(*int64)(unsafe.Add(mBase, uint32(v259)+24))
										if v262 != v269 {
											m.G0 = v11 + int32(16448)
											return
										} else {
											F_closeRepldbfd(m, v13)
											mBase = m.M
											v272 = m.ExcPending
											if v272 != 0 {
												return
											} else {
												v273 = *(*int32)(unsafe.Add(mBase, uint32(v13)+8))
												v274 = int32(0)
												v276 = *(*int32)(unsafe.Add(mBase, uint32(v273)))
												v277 = *(*int32)(unsafe.Add(mBase, uint32(v276)+80))
												v278 = m.T0[v277].(func(*base.Module, int32, int32, int32) int32)(m, v273, v274, v274)
												mBase = m.M
												v279 = m.ExcPending
												if v279 != 0 {
													return
												} else {
													v280 = F_replicaPutOnline(m, v13)
													mBase = m.M
													v281 = m.ExcPending
													if v281 != 0 {
														return
													} else {
														if v280 != 0 {
															v284 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+205)))
															if v284&int32(32) != 0 {
																F__serverAssert(m, int32(_a_F_sendBulkToReplica_1), int32(_a_F_sendBulkToReplica_2), int32(1617))
																mBase = m.M
																v341 = m.ExcPending
																if v341 != 0 {
																	return
																} else {
																	F_abort(m)
																	mBase = m.M
																	base.Wasm_trap_unreachable()
																	for {
																	}
																}
															} else {
																v287 = *(*int32)(unsafe.Add(mBase, uint32(v13)+104))
																*(*int32)(unsafe.Add(mBase, uint32(v287)+4)) = int32(0)
																v291 = *(*int64)(unsafe.Add(mBase, _c_F_sendBulkToReplica[1]))
																if v291 != int64(0) {
																	v305 = *(*int32)(unsafe.Add(mBase, uint32(v13)+200))
																	if v305&int32(4194304) != 0 {
																	} else {
																		v308 = *(*int32)(unsafe.Add(mBase, uint32(v13)+104))
																		if v308 == int32(0) {
																			v315 = F_clusterSlotMigrationShouldInstallWriteHandler(m, v13)
																			mBase = m.M
																			if v315 == int32(0) {
																			} else {
																				v318 = *(*int32)(unsafe.Add(mBase, uint32(v13)+200))
																				*(*int32)(unsafe.Add(mBase, uint32(v13)+200)) = v318 | int32(4194304)
																				v323 = *(*int32)(unsafe.Add(mBase, _c_F_sendBulkToReplica[2]))
																				F_listLinkNodeHead(m, v323, v13+int32(168))
																				mBase = m.M
																			}
																		} else {
																			v311 = *(*int32)(unsafe.Add(mBase, uint32(v308)))
																			switch v311 {
																			case 0:
																				v315 = F_clusterSlotMigrationShouldInstallWriteHandler(m, v13)
																				mBase = m.M
																				if v315 == int32(0) {
																				} else {
																					v318 = *(*int32)(unsafe.Add(mBase, uint32(v13)+200))
																					*(*int32)(unsafe.Add(mBase, uint32(v13)+200)) = v318 | int32(4194304)
																					v323 = *(*int32)(unsafe.Add(mBase, _c_F_sendBulkToReplica[2]))
																					F_listLinkNodeHead(m, v323, v13+int32(168))
																					mBase = m.M
																				}
																			default:
																			case 9, 11:
																				if v305&int32(1024) != 0 {
																				} else {
																					v314 = *(*int32)(unsafe.Add(mBase, uint32(v308)+4))
																					if v314 != 0 {
																					} else {
																						v315 = F_clusterSlotMigrationShouldInstallWriteHandler(m, v13)
																						mBase = m.M
																						if v315 == int32(0) {
																						} else {
																							v318 = *(*int32)(unsafe.Add(mBase, uint32(v13)+200))
																							*(*int32)(unsafe.Add(mBase, uint32(v13)+200)) = v318 | int32(4194304)
																							v323 = *(*int32)(unsafe.Add(mBase, _c_F_sendBulkToReplica[2]))
																							F_listLinkNodeHead(m, v323, v13+int32(168))
																							mBase = m.M
																						}
																					}
																				}
																			}
																		}
																	}
																	m.G0 = v11 + int32(16448)
																	return
																} else {
																	v295 = *(*int32)(unsafe.Add(mBase, _c_F_sendBulkToReplica[3]))
																	*(*int32)(unsafe.Add(mBase, uint32(v11)+uint32(_c_F_sendBulkToReplica[4]))) = v295
																	F_replicationFeedReplicas(m, int32(-1), v11+int32(16444), int32(1))
																	mBase = m.M
																	v302 = m.ExcPending
																	if v302 != 0 {
																		return
																	} else {
																		v305 = *(*int32)(unsafe.Add(mBase, uint32(v13)+200))
																		if v305&int32(4194304) != 0 {
																		} else {
																			v308 = *(*int32)(unsafe.Add(mBase, uint32(v13)+104))
																			if v308 == int32(0) {
																				v315 = F_clusterSlotMigrationShouldInstallWriteHandler(m, v13)
																				mBase = m.M
																				if v315 == int32(0) {
																				} else {
																					v318 = *(*int32)(unsafe.Add(mBase, uint32(v13)+200))
																					*(*int32)(unsafe.Add(mBase, uint32(v13)+200)) = v318 | int32(4194304)
																					v323 = *(*int32)(unsafe.Add(mBase, _c_F_sendBulkToReplica[2]))
																					F_listLinkNodeHead(m, v323, v13+int32(168))
																					mBase = m.M
																				}
																			} else {
																				v311 = *(*int32)(unsafe.Add(mBase, uint32(v308)))
																				switch v311 {
																				case 0:
																					v315 = F_clusterSlotMigrationShouldInstallWriteHandler(m, v13)
																					mBase = m.M
																					if v315 == int32(0) {
																					} else {
																						v318 = *(*int32)(unsafe.Add(mBase, uint32(v13)+200))
																						*(*int32)(unsafe.Add(mBase, uint32(v13)+200)) = v318 | int32(4194304)
																						v323 = *(*int32)(unsafe.Add(mBase, _c_F_sendBulkToReplica[2]))
																						F_listLinkNodeHead(m, v323, v13+int32(168))
																						mBase = m.M
																					}
																				default:
																				case 9, 11:
																					if v305&int32(1024) != 0 {
																					} else {
																						v314 = *(*int32)(unsafe.Add(mBase, uint32(v308)+4))
																						if v314 != 0 {
																						} else {
																							v315 = F_clusterSlotMigrationShouldInstallWriteHandler(m, v13)
																							mBase = m.M
																							if v315 == int32(0) {
																							} else {
																								v318 = *(*int32)(unsafe.Add(mBase, uint32(v13)+200))
																								*(*int32)(unsafe.Add(mBase, uint32(v13)+200)) = v318 | int32(4194304)
																								v323 = *(*int32)(unsafe.Add(mBase, _c_F_sendBulkToReplica[2]))
																								F_listLinkNodeHead(m, v323, v13+int32(168))
																								mBase = m.M
																							}
																						}
																					}
																				}
																			}
																		}
																		m.G0 = v11 + int32(16448)
																		return
																	}
																}
															}
														} else {
															v282 = F_freeClient(m, v13)
															mBase = m.M
															v283 = m.ExcPending
															if v283 != 0 {
																return
															} else {
																m.G0 = v11 + int32(16448)
																return
															}
														}
													}
												}
											}
										}
									} else {
										v239 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
										if v239 == int32(3) {
											m.G0 = v11 + int32(16448)
											return
										} else {
											v243 = *(*int32)(unsafe.Add(mBase, _c_F_sendBulkToReplica[5]))
											if int32(3) < v243 {
												v257 = F_freeClient(m, v13)
												mBase = m.M
												v258 = m.ExcPending
												if v258 != 0 {
													return
												} else {
													m.G0 = v11 + int32(16448)
													return
												}
											} else {
												v246 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
												v247 = *(*int32)(unsafe.Add(mBase, uint32(v246)+88))
												v248 = m.T0[v247].(func(*base.Module, int32) int32)(m, l0)
												mBase = m.M
												v249 = m.ExcPending
												if v249 != 0 {
													return
												} else {
													*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = v248
													F__serverLog(m, int32(3), int32(_a_F_sendBulkToReplica_3), v11+int32(16))
													mBase = m.M
													v256 = m.ExcPending
													if v256 != 0 {
														return
													} else {
														v257 = F_freeClient(m, v13)
														mBase = m.M
														v258 = m.ExcPending
														if v258 != 0 {
															return
														} else {
															m.G0 = v11 + int32(16448)
															return
														}
													}
												}
											}
										}
									}
								}
							} else {
								v215 = *(*int32)(unsafe.Add(mBase, _c_F_sendBulkToReplica[5]))
								if int32(3) < v215 {
									v229 = F_freeClient(m, v13)
									mBase = m.M
									v230 = m.ExcPending
									if v230 != 0 {
										return
									} else {
										m.G0 = v11 + int32(16448)
										return
									}
								} else {
									if v211 != 0 {
										v220 = *(*int32)(unsafe.Add(mBase, _c_F_sendBulkToReplica[6]))
										v221 = F___strerror_l(m, v220, v220)
										mBase = m.M
										v222 = v221
									} else {
										v222 = int32(_a_F_sendBulkToReplica_4)
									}
									*(*int32)(unsafe.Add(mBase, uint32(v11))) = v222
									F__serverLog(m, int32(3), int32(_a_F_sendBulkToReplica_5), v11)
									mBase = m.M
									v227 = m.ExcPending
									if v227 != 0 {
										return
									} else {
										v229 = F_freeClient(m, v13)
										mBase = m.M
										v230 = m.ExcPending
										if v230 != 0 {
											return
										} else {
											m.G0 = v11 + int32(16448)
											return
										}
									}
								}
							}
						}
					}
				default:
					F_sdsfree(m, v172)
					mBase = m.M
					v195 = m.ExcPending
					if v195 != 0 {
						return
					} else {
						v196 = *(*int32)(unsafe.Add(mBase, uint32(v13)+104))
						*(*int32)(unsafe.Add(mBase, uint32(v196)+32)) = int32(0)
						v199 = v196
						v202 = *(*int32)(unsafe.Add(mBase, uint32(v199)+8))
						v203 = *(*int64)(unsafe.Add(mBase, uint32(v199)+16))
						v204 = int32(0)
						v205 = F___lseek(m, v202, v203, v204)
						mBase = m.M
						v206 = *(*int32)(unsafe.Add(mBase, uint32(v13)+104))
						v207 = *(*int32)(unsafe.Add(mBase, uint32(v206)+8))
						v211 = F_read(m, v207, v11+int32(48), int32(16384))
						mBase = m.M
						if v204 < v211 {
							v233 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
							v234 = *(*int32)(unsafe.Add(mBase, uint32(v233)+68))
							v235 = m.T0[v234].(func(*base.Module, int32, int32, int32) int32)(m, l0, v11+int32(48), v211)
							mBase = m.M
							v236 = m.ExcPending
							if v236 != 0 {
								return
							} else {
								if v235 != int32(-1) {
									v259 = *(*int32)(unsafe.Add(mBase, uint32(v13)+104))
									v260 = *(*int64)(unsafe.Add(mBase, uint32(v259)+16))
									v261 = base.I64_extend_i32_s(v235)
									v262 = v260 + v261
									*(*int64)(unsafe.Add(mBase, uint32(v259)+16)) = v262
									v264 = int32(_a_F_sendBulkToReplica_0)
									v266 = *(*int64)(unsafe.Add(mBase, _c_F_sendBulkToReplica[0]))
									*(*int64)(unsafe.Add(mBase, _c_F_sendBulkToReplica[0])) = v266 + v261
									v269 = *(*int64)(unsafe.Add(mBase, uint32(v259)+24))
									if v262 != v269 {
										m.G0 = v11 + int32(16448)
										return
									} else {
										F_closeRepldbfd(m, v13)
										mBase = m.M
										v272 = m.ExcPending
										if v272 != 0 {
											return
										} else {
											v273 = *(*int32)(unsafe.Add(mBase, uint32(v13)+8))
											v274 = int32(0)
											v276 = *(*int32)(unsafe.Add(mBase, uint32(v273)))
											v277 = *(*int32)(unsafe.Add(mBase, uint32(v276)+80))
											v278 = m.T0[v277].(func(*base.Module, int32, int32, int32) int32)(m, v273, v274, v274)
											mBase = m.M
											v279 = m.ExcPending
											if v279 != 0 {
												return
											} else {
												v280 = F_replicaPutOnline(m, v13)
												mBase = m.M
												v281 = m.ExcPending
												if v281 != 0 {
													return
												} else {
													if v280 != 0 {
														v284 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+205)))
														if v284&int32(32) != 0 {
															F__serverAssert(m, int32(_a_F_sendBulkToReplica_1), int32(_a_F_sendBulkToReplica_2), int32(1617))
															mBase = m.M
															v341 = m.ExcPending
															if v341 != 0 {
																return
															} else {
																F_abort(m)
																mBase = m.M
																base.Wasm_trap_unreachable()
																for {
																}
															}
														} else {
															v287 = *(*int32)(unsafe.Add(mBase, uint32(v13)+104))
															*(*int32)(unsafe.Add(mBase, uint32(v287)+4)) = int32(0)
															v291 = *(*int64)(unsafe.Add(mBase, _c_F_sendBulkToReplica[1]))
															if v291 != int64(0) {
																v305 = *(*int32)(unsafe.Add(mBase, uint32(v13)+200))
																if v305&int32(4194304) != 0 {
																} else {
																	v308 = *(*int32)(unsafe.Add(mBase, uint32(v13)+104))
																	if v308 == int32(0) {
																		v315 = F_clusterSlotMigrationShouldInstallWriteHandler(m, v13)
																		mBase = m.M
																		if v315 == int32(0) {
																		} else {
																			v318 = *(*int32)(unsafe.Add(mBase, uint32(v13)+200))
																			*(*int32)(unsafe.Add(mBase, uint32(v13)+200)) = v318 | int32(4194304)
																			v323 = *(*int32)(unsafe.Add(mBase, _c_F_sendBulkToReplica[2]))
																			F_listLinkNodeHead(m, v323, v13+int32(168))
																			mBase = m.M
																		}
																	} else {
																		v311 = *(*int32)(unsafe.Add(mBase, uint32(v308)))
																		switch v311 {
																		case 0:
																			v315 = F_clusterSlotMigrationShouldInstallWriteHandler(m, v13)
																			mBase = m.M
																			if v315 == int32(0) {
																			} else {
																				v318 = *(*int32)(unsafe.Add(mBase, uint32(v13)+200))
																				*(*int32)(unsafe.Add(mBase, uint32(v13)+200)) = v318 | int32(4194304)
																				v323 = *(*int32)(unsafe.Add(mBase, _c_F_sendBulkToReplica[2]))
																				F_listLinkNodeHead(m, v323, v13+int32(168))
																				mBase = m.M
																			}
																		default:
																		case 9, 11:
																			if v305&int32(1024) != 0 {
																			} else {
																				v314 = *(*int32)(unsafe.Add(mBase, uint32(v308)+4))
																				if v314 != 0 {
																				} else {
																					v315 = F_clusterSlotMigrationShouldInstallWriteHandler(m, v13)
																					mBase = m.M
																					if v315 == int32(0) {
																					} else {
																						v318 = *(*int32)(unsafe.Add(mBase, uint32(v13)+200))
																						*(*int32)(unsafe.Add(mBase, uint32(v13)+200)) = v318 | int32(4194304)
																						v323 = *(*int32)(unsafe.Add(mBase, _c_F_sendBulkToReplica[2]))
																						F_listLinkNodeHead(m, v323, v13+int32(168))
																						mBase = m.M
																					}
																				}
																			}
																		}
																	}
																}
																m.G0 = v11 + int32(16448)
																return
															} else {
																v295 = *(*int32)(unsafe.Add(mBase, _c_F_sendBulkToReplica[3]))
																*(*int32)(unsafe.Add(mBase, uint32(v11)+uint32(_c_F_sendBulkToReplica[4]))) = v295
																F_replicationFeedReplicas(m, int32(-1), v11+int32(16444), int32(1))
																mBase = m.M
																v302 = m.ExcPending
																if v302 != 0 {
																	return
																} else {
																	v305 = *(*int32)(unsafe.Add(mBase, uint32(v13)+200))
																	if v305&int32(4194304) != 0 {
																	} else {
																		v308 = *(*int32)(unsafe.Add(mBase, uint32(v13)+104))
																		if v308 == int32(0) {
																			v315 = F_clusterSlotMigrationShouldInstallWriteHandler(m, v13)
																			mBase = m.M
																			if v315 == int32(0) {
																			} else {
																				v318 = *(*int32)(unsafe.Add(mBase, uint32(v13)+200))
																				*(*int32)(unsafe.Add(mBase, uint32(v13)+200)) = v318 | int32(4194304)
																				v323 = *(*int32)(unsafe.Add(mBase, _c_F_sendBulkToReplica[2]))
																				F_listLinkNodeHead(m, v323, v13+int32(168))
																				mBase = m.M
																			}
																		} else {
																			v311 = *(*int32)(unsafe.Add(mBase, uint32(v308)))
																			switch v311 {
																			case 0:
																				v315 = F_clusterSlotMigrationShouldInstallWriteHandler(m, v13)
																				mBase = m.M
																				if v315 == int32(0) {
																				} else {
																					v318 = *(*int32)(unsafe.Add(mBase, uint32(v13)+200))
																					*(*int32)(unsafe.Add(mBase, uint32(v13)+200)) = v318 | int32(4194304)
																					v323 = *(*int32)(unsafe.Add(mBase, _c_F_sendBulkToReplica[2]))
																					F_listLinkNodeHead(m, v323, v13+int32(168))
																					mBase = m.M
																				}
																			default:
																			case 9, 11:
																				if v305&int32(1024) != 0 {
																				} else {
																					v314 = *(*int32)(unsafe.Add(mBase, uint32(v308)+4))
																					if v314 != 0 {
																					} else {
																						v315 = F_clusterSlotMigrationShouldInstallWriteHandler(m, v13)
																						mBase = m.M
																						if v315 == int32(0) {
																						} else {
																							v318 = *(*int32)(unsafe.Add(mBase, uint32(v13)+200))
																							*(*int32)(unsafe.Add(mBase, uint32(v13)+200)) = v318 | int32(4194304)
																							v323 = *(*int32)(unsafe.Add(mBase, _c_F_sendBulkToReplica[2]))
																							F_listLinkNodeHead(m, v323, v13+int32(168))
																							mBase = m.M
																						}
																					}
																				}
																			}
																		}
																	}
																	m.G0 = v11 + int32(16448)
																	return
																}
															}
														}
													} else {
														v282 = F_freeClient(m, v13)
														mBase = m.M
														v283 = m.ExcPending
														if v283 != 0 {
															return
														} else {
															m.G0 = v11 + int32(16448)
															return
														}
													}
												}
											}
										}
									}
								} else {
									v239 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
									if v239 == int32(3) {
										m.G0 = v11 + int32(16448)
										return
									} else {
										v243 = *(*int32)(unsafe.Add(mBase, _c_F_sendBulkToReplica[5]))
										if int32(3) < v243 {
											v257 = F_freeClient(m, v13)
											mBase = m.M
											v258 = m.ExcPending
											if v258 != 0 {
												return
											} else {
												m.G0 = v11 + int32(16448)
												return
											}
										} else {
											v246 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
											v247 = *(*int32)(unsafe.Add(mBase, uint32(v246)+88))
											v248 = m.T0[v247].(func(*base.Module, int32) int32)(m, l0)
											mBase = m.M
											v249 = m.ExcPending
											if v249 != 0 {
												return
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = v248
												F__serverLog(m, int32(3), int32(_a_F_sendBulkToReplica_3), v11+int32(16))
												mBase = m.M
												v256 = m.ExcPending
												if v256 != 0 {
													return
												} else {
													v257 = F_freeClient(m, v13)
													mBase = m.M
													v258 = m.ExcPending
													if v258 != 0 {
														return
													} else {
														m.G0 = v11 + int32(16448)
														return
													}
												}
											}
										}
									}
								}
							}
						} else {
							v215 = *(*int32)(unsafe.Add(mBase, _c_F_sendBulkToReplica[5]))
							if int32(3) < v215 {
								v229 = F_freeClient(m, v13)
								mBase = m.M
								v230 = m.ExcPending
								if v230 != 0 {
									return
								} else {
									m.G0 = v11 + int32(16448)
									return
								}
							} else {
								if v211 != 0 {
									v220 = *(*int32)(unsafe.Add(mBase, _c_F_sendBulkToReplica[6]))
									v221 = F___strerror_l(m, v220, v220)
									mBase = m.M
									v222 = v221
								} else {
									v222 = int32(_a_F_sendBulkToReplica_4)
								}
								*(*int32)(unsafe.Add(mBase, uint32(v11))) = v222
								F__serverLog(m, int32(3), int32(_a_F_sendBulkToReplica_5), v11)
								mBase = m.M
								v227 = m.ExcPending
								if v227 != 0 {
									return
								} else {
									v229 = F_freeClient(m, v13)
									mBase = m.M
									v230 = m.ExcPending
									if v230 != 0 {
										return
									} else {
										m.G0 = v11 + int32(16448)
										return
									}
								}
							}
						}
					}
				}
			} else {
				v46 = *(*int32)(unsafe.Add(mBase, _c_F_sendBulkToReplica[5]))
				if int32(3) < v46 {
					v60 = F_freeClient(m, v13)
					mBase = m.M
					v61 = m.ExcPending
					if v61 != 0 {
						return
					} else {
						m.G0 = v11 + int32(16448)
						return
					}
				} else {
					v49 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
					v50 = *(*int32)(unsafe.Add(mBase, uint32(v49)+88))
					v51 = m.T0[v50].(func(*base.Module, int32) int32)(m, l0)
					mBase = m.M
					v52 = m.ExcPending
					if v52 != 0 {
						return
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v11)+32)) = v51
						F__serverLog(m, int32(3), int32(_a_F_sendBulkToReplica_6), v11+int32(32))
						mBase = m.M
						v59 = m.ExcPending
						if v59 != 0 {
							return
						} else {
							v60 = F_freeClient(m, v13)
							mBase = m.M
							v61 = m.ExcPending
							if v61 != 0 {
								return
							} else {
								m.G0 = v11 + int32(16448)
								return
							}
						}
					}
				}
			}
		}
	}
}
func F_sendCurrentOffsetToReplica(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v14 int64
	_ = v14
	var v16 int32
	_ = v16
	var v17 int64
	_ = v17
	var v24 int32
	_ = v24
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v44 int64
	_ = v44
	var v48 int64
	_ = v48
	var v53 int32
	_ = v53
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
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	v6 = m.G0
	v8 = v6 - int32(192)
	m.G0 = v8
	v14 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	*(*int64)(unsafe.Add(mBase, uint32(v8+int32(48)))) = v14
	v16 = int32(_a_F_sendCurrentOffsetToReplica_0)
	v17 = *(*int64)(unsafe.Add(mBase, _c_F_sendCurrentOffsetToReplica[0]))
	*(*int64)(unsafe.Add(mBase, uint32(v8)+32)) = v17
	*(*int32)(unsafe.Add(mBase, uint32(v8)+40)) = int32(_a_F_sendCurrentOffsetToReplica_1)
	v24 = *(*int32)(unsafe.Add(mBase, _c_F_sendCurrentOffsetToReplica[1]))
	*(*int32)(unsafe.Add(mBase, uint32(v8)+44)) = v24
	v32 = F_snprintf(m, v8+int32(64), int32(128), int32(_a_F_sendCurrentOffsetToReplica_2), v8+int32(32))
	mBase = m.M
	v35 = m.ExcPending
	if v35 != 0 {
		return int32(0)
	} else {
		v37 = *(*int32)(unsafe.Add(mBase, _c_F_sendCurrentOffsetToReplica[2]))
		if int32(2) < v37 {
			v56 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			v60 = *(*int32)(unsafe.Add(mBase, _c_F_sendCurrentOffsetToReplica[3]))
			v64 = *(*int32)(unsafe.Add(mBase, uint32(v56)))
			v65 = *(*int32)(unsafe.Add(mBase, uint32(v64)+92))
			v66 = m.T0[v65].(func(*base.Module, int32, int32, int32, int64) int32)(m, v56, v8+int32(64), v32, base.I64_extend_i32_s(v60*int32(1000)))
			mBase = m.M
			v67 = m.ExcPending
			if v67 != 0 {
				return int32(0)
			} else {
				if v66 == v32 {
					v72 = int32(0)
					m.G0 = v8 + int32(192)
					return v72
				} else {
					F_freeClientAsync(m, l0)
					mBase = m.M
					v70 = m.ExcPending
					if v70 != 0 {
						return int32(0)
					} else {
						v72 = int32(-1)
						m.G0 = v8 + int32(192)
						return v72
					}
				}
			}
		} else {
			v40 = F_replicationGetReplicaName(m, l0)
			mBase = m.M
			v41 = m.ExcPending
			if v41 != 0 {
				return int32(0)
			} else {
				v44 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
				*(*int64)(unsafe.Add(mBase, uint32(v8+int32(16)))) = v44
				*(*int32)(unsafe.Add(mBase, uint32(v8))) = v40
				v48 = *(*int64)(unsafe.Add(mBase, _c_F_sendCurrentOffsetToReplica[0]))
				*(*int64)(unsafe.Add(mBase, uint32(v8)+8)) = v48
				F__serverLog(m, int32(2), int32(_a_F_sendCurrentOffsetToReplica_3), v8)
				mBase = m.M
				v53 = m.ExcPending
				if v53 != 0 {
					return int32(0)
				} else {
					v56 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
					v60 = *(*int32)(unsafe.Add(mBase, _c_F_sendCurrentOffsetToReplica[3]))
					v64 = *(*int32)(unsafe.Add(mBase, uint32(v56)))
					v65 = *(*int32)(unsafe.Add(mBase, uint32(v64)+92))
					v66 = m.T0[v65].(func(*base.Module, int32, int32, int32, int64) int32)(m, v56, v8+int32(64), v32, base.I64_extend_i32_s(v60*int32(1000)))
					mBase = m.M
					v67 = m.ExcPending
					if v67 != 0 {
						return int32(0)
					} else {
						if v66 == v32 {
							v72 = int32(0)
							m.G0 = v8 + int32(192)
							return v72
						} else {
							F_freeClientAsync(m, l0)
							mBase = m.M
							v70 = m.ExcPending
							if v70 != 0 {
								return int32(0)
							} else {
								v72 = int32(-1)
								m.G0 = v8 + int32(192)
								return v72
							}
						}
					}
				}
			}
		}
	}
}
func F_setImportingSlotSource(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	v6 = *(*int32)(unsafe.Add(mBase, _c_F_setImportingSlotSource[0]))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(v6)+48))
	v8 = F_dictFind(m, v7, l0)
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return
	} else {
		if l1 != 0 {
			v18 = *(*int32)(unsafe.Add(mBase, _c_F_setImportingSlotSource[0]))
			v19 = *(*int32)(unsafe.Add(mBase, uint32(v18)+48))
			if v8 == int32(0) {
				v23 = F_dictAdd(m, v19, l0, l1)
				mBase = m.M
				v24 = m.ExcPending
				if v24 != 0 {
					return
				} else {
					return
				}
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v8)+8)) = l1
				return
			}
		} else {
			if v8 == int32(0) {
				return
			} else {
				v13 = *(*int32)(unsafe.Add(mBase, _c_F_setImportingSlotSource[0]))
				v14 = *(*int32)(unsafe.Add(mBase, uint32(v13)+48))
				v15 = F_dictDelete(m, v14, l0)
				mBase = m.M
				v16 = m.ExcPending
				if v16 != 0 {
					return
				} else {
					return
				}
			}
		}
	}
}
func F_set_jemalloc_bg_thread(m *base.Module, l0 int32) {
	return
}
func F_setn(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v70 int32
	_ = v70
	var v71 int64
	_ = v71
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	F_luaL_checktype(m, l0, int32(1), int32(5))
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		v8 = m.G3
		v12 = F_luaL_error(m, l0, v8+int32(_a_F_setn_0), int32(0))
		mBase = m.M
		v13 = m.ExcPending
		if v13 != 0 {
			return int32(0)
		} else {
			v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
			v25 = v20 + int32(0)
			v26 = m.G398
			if base.Ui32(v25) < base.Ui32(v19) {
				v28 = v25
			} else {
				v28 = v26
			}
			v70 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			v71 = *(*int64)(unsafe.Add(mBase, uint32(v28)))
			*(*int64)(unsafe.Add(mBase, uint32(v70))) = v71
			v73 = *(*int32)(unsafe.Add(mBase, uint32(v28)+8))
			*(*int32)(unsafe.Add(mBase, uint32(v70)+8)) = v73
			v75 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v75 + int32(16)
			return int32(1)
		}
	}
}
func F_setsockopt(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	v9 = m.G0
	v11 = v9 - int32(16)
	m.G0 = v11
	if l1 != int32(1) {
	} else {
	}
	m.G0 = v11 + int32(16)
	return int32(0)
}
func F_setupSigSegvHandler(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v31 int32
	_ = v31
	var v36 int32
	_ = v36
	var v59 int32
	_ = v59
	var v64 int32
	_ = v64
	var v87 int32
	_ = v87
	var v92 int32
	_ = v92
	var v115 int32
	_ = v115
	var v120 int32
	_ = v120
	var v143 int32
	_ = v143
	var v148 int32
	_ = v148
	var v171 int32
	_ = v171
	v2 = m.G0
	v4 = v2 - int32(144)
	m.G0 = v4
	v7 = *(*int32)(unsafe.Add(mBase, _c_F_setupSigSegvHandler[0]))
	if v7 != 0 {
	} else {
		*(*int32)(unsafe.Add(mBase, _c_F_setupSigSegvHandler[0])) = int32(1)
	}
	*(*int64)(unsafe.Add(mBase, uint32(v4+int32(8)))) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v4)+4)) = int32(519)
	*(*int32)(unsafe.Add(mBase, uint32(v4)+136)) = int32(1073741828)
	v31 = *(*int32)(unsafe.Add(mBase, _c_F_setupSigSegvHandler[1]))
	if v31 == int32(0) {
	} else {
		v36 = v4 + int32(4)
		if v36 == int32(0) {
		} else {
			v59 = F___memcpy(m, int32(9118724), v36, int32(140))
			mBase = m.M
		}
		v64 = v4 + int32(4)
		if v64 == int32(0) {
		} else {
			v87 = F___memcpy(m, int32(9118164), v64, int32(140))
			mBase = m.M
		}
		v92 = v4 + int32(4)
		if v92 == int32(0) {
		} else {
			v115 = F___memcpy(m, int32(9118304), v92, int32(140))
			mBase = m.M
		}
		v120 = v4 + int32(4)
		if v120 == int32(0) {
		} else {
			v143 = F___memcpy(m, int32(9117744), v120, int32(140))
			mBase = m.M
		}
		v148 = v4 + int32(4)
		if v148 == int32(0) {
		} else {
			v171 = F___memcpy(m, int32(9118024), v148, int32(140))
			mBase = m.M
		}
	}
	m.G0 = v4 + int32(144)
	return
}
func F_sha1hex(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v31 int32
	_ = v31
	var v36 int32
	_ = v36
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	v6 = m.G0
	v8 = v6 - int32(112)
	m.G0 = v8
	v11 = v8 + int32(20)
	F_SHA1Init(m, v11)
	mBase = m.M
	F_SHA1Update(m, v11, l1, l2)
	mBase = m.M
	F_SHA1Final(m, v8, v11)
	mBase = m.M
	v22 = int32(0)
	for {
		v25 = int32(1)
		v27 = l0 + v22<<(uint(v25)%32)
		v31 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8+v22))))
		v36 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v31&int32(15))+uint32(_c_F_sha1hex[0]))))
		*(*uint8)(unsafe.Add(mBase, uint32(v27+v25))) = uint8(v36)
		v42 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v31)>>(uint(int32(4))%32)))+uint32(_c_F_sha1hex[0]))))
		*(*uint8)(unsafe.Add(mBase, uint32(v27))) = uint8(v42)
		v45 = v22 + v25
		if v45 != int32(20) {
			v22 = v45
			continue
		} else {
			break
		}
		break
	}
	v48 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+40)) = uint8(v48)
	m.G0 = v8 + int32(112)
	return
}
func F_sha256_final(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+80))
	m.Env.Vkmem_hash_final(m, v3, l1, int32(32))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(l0)+80)) = int32(0)
	return
}
func F_sha256_init(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	v3 = m.Env.Vkmem_hash_create(m, int32(2))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(l0)+80)) = v3
	return
}
func F_shl(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	if base.Ui32(int32(31)) < base.Ui32(l1) {
		v13 = l1 + int32(-32)
		v14 = l0
		v15 = int32(0)
	} else {
		v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		v13 = l1
		v14 = l0 + int32(4)
		v15 = v9
	}
	v16 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v15 << (uint(v13) % 32)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = int32(base.Ui32(v15)>>(uint(int32(32)-v13)%32)) | v16<<(uint(v13)%32)
	return
}
func F_shouldFilterFromCommandList(m *base.Module, l0 int32, l1 int32) int32 {
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
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v42 int64
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v50 int32
	_ = v50
	var v51 int64
	_ = v51
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v62 int32
	_ = v62
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v71 int64
	_ = v71
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v82 int64
	_ = v82
	var v89 int64
	_ = v89
	var v93 int64
	_ = v93
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	var v104 int32
	_ = v104
	var v111 int32
	_ = v111
	var v114 int32
	_ = v114
	var v117 int32
	_ = v117
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v125 int32
	_ = v125
	var v132 int32
	_ = v132
	var v135 int32
	_ = v135
	var v138 int32
	_ = v138
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v154 int32
	_ = v154
	var v165 int32
	_ = v165
	var v167 int64
	_ = v167
	var v172 int32
	_ = v172
	v8 = m.G0
	v10 = v8 - int32(16)
	m.G0 = v10
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	switch v12 {
	case 0:
		goto L6
	case 1:
		goto L5
	case 2:
		goto L4
	default:
		goto L3
	}
L1:
	;
	m.G0 = v10 + int32(16)
	return v172
L2:
	;
	v167 = *(*int64)(unsafe.Add(mBase, uint32(l0)+64))
	v172 = base.B2i32(v167&v93 == int64(0))
	goto L1
L3:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10))) = v12
	F__serverPanic_1(m, int32(_a_F_shouldFilterFromCommandList_0), int32(5634), int32(_a_F_shouldFilterFromCommandList_1), v10)
	mBase = m.M
	v165 = m.ExcPending
	if v165 != 0 {
		goto L10
	} else {
		goto L43
	}
L4:
	;
	v99 = int32(0)
	v101 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v104 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v101+int32(-1)))))
	switch v104 & int32(7) {
	case 0:
		goto L35
	case 1:
		goto L34
	case 2:
		goto L33
	case 3:
		goto L32
	case 4:
		goto L31
	default:
		v121 = v99
		goto L30
	}
L5:
	;
	v39 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	if v39 == int32(0) {
		goto L17
	} else {
		goto L18
	}
L6:
	;
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	if v13 == int32(0) {
		goto L8
	} else {
		goto L9
	}
L7:
	;
	v26 = int32(0)
	if v25 == v26 {
		v36 = v26
		goto L13
	} else {
		goto L14
	}
L8:
	;
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v18 = F_moduleGetHandleByName(m, v17)
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		goto L10
	} else {
		goto L11
	}
L9:
	;
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	v25 = v16
	goto L7
L10:
	;
	return int32(0)
L11:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+8)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l1)+16)) = v18
	v25 = v18
	goto L7
L12:
	;
	v172 = base.B2i32(v36 == int32(0))
	goto L1
L13:
	;
	goto L12
L14:
	;
	v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	if v30 != int32(562) {
		v36 = v26
		goto L13
	} else {
		goto L15
	}
L15:
	;
	v33 = *(*int32)(unsafe.Add(mBase, uint32(l0)+208))
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v33)))
	v36 = base.B2i32(v34 == v25)
	goto L13
L16:
	;
	if base.B2i32(v93 == int64(0)) == int32(0) {
		goto L2
	} else {
		goto L29
	}
L17:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v44 = int32(0)
	v50 = *(*int32)(unsafe.Add(mBase, _c_F_shouldFilterFromCommandList[0]))
	v51 = *(*int64)(unsafe.Add(mBase, uint32(v50)+8))
	if base.B2i32(v51 == int64(0)) == v44 {
		goto L20
	} else {
		goto L21
	}
L18:
	;
	v42 = *(*int64)(unsafe.Add(mBase, uint32(l1)+16))
	v93 = v42
	goto L16
L19:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+8)) = int32(1)
	*(*int64)(unsafe.Add(mBase, uint32(l1)+16)) = v89
	v93 = v89
	goto L16
L20:
	;
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v50)))
	v58 = F_strcasecmp(m, v43, v57)
	mBase = m.M
	if v58 == int32(0) {
		v82 = v51
		goto L22
	} else {
		goto L23
	}
L21:
	;
	v89 = int64(0)
	goto L19
L22:
	;
	v89 = v82
	goto L19
L23:
	;
	v62 = v44
	goto L24
L24:
	;
	v67 = v62 + int32(1)
	v70 = v50 + v67<<(uint(int32(4))%32)
	v71 = *(*int64)(unsafe.Add(mBase, uint32(v70)+8))
	if base.B2i32(v71 == int64(0)) == int32(0) {
		goto L26
	} else {
		goto L27
	}
L25:
	;
	v82 = v71
	goto L22
L26:
	;
	v77 = *(*int32)(unsafe.Add(mBase, uint32(v70)))
	v78 = F_strcasecmp(m, v43, v77)
	mBase = m.M
	if v78 != 0 {
		v62 = v67
		goto L24
	} else {
		goto L28
	}
L27:
	;
	v89 = int64(0)
	goto L19
L28:
	;
	goto L25
L29:
	;
	v172 = int32(1)
	goto L1
L30:
	;
	v122 = *(*int32)(unsafe.Add(mBase, uint32(l0)+140))
	v125 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v122+int32(-1)))))
	switch v125 & int32(7) {
	case 0:
		goto L41
	case 1:
		goto L40
	case 2:
		goto L39
	case 3:
		goto L38
	case 4:
		goto L37
	default:
		v142 = v99
		goto L36
	}
L31:
	;
	v120 = *(*int32)(unsafe.Add(mBase, uint32(v101+int32(-17))))
	v121 = v120
	goto L30
L32:
	;
	v117 = *(*int32)(unsafe.Add(mBase, uint32(v101+int32(-9))))
	v121 = v117
	goto L30
L33:
	;
	v114 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v101+int32(-5)))))
	v121 = v114
	goto L30
L34:
	;
	v111 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v101+int32(-3)))))
	v121 = v111
	goto L30
L35:
	;
	v121 = int32(base.Ui32(v104) >> (uint(int32(3)) % 32))
	goto L30
L36:
	;
	v144 = int32(0)
	v145 = m.G0
	v146 = int32(16)
	v147 = v145 - v146
	m.G0 = v147
	*(*int32)(unsafe.Add(mBase, uint32(v147)+12)) = v144
	v154 = F_stringmatchlen_impl(m, v101, v121, v122, v142, int32(1), v147+int32(12), v144)
	mBase = m.M
	m.G0 = v147 + v146
	goto L42
L37:
	;
	v141 = *(*int32)(unsafe.Add(mBase, uint32(v122+int32(-17))))
	v142 = v141
	goto L36
L38:
	;
	v138 = *(*int32)(unsafe.Add(mBase, uint32(v122+int32(-9))))
	v142 = v138
	goto L36
L39:
	;
	v135 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v122+int32(-5)))))
	v142 = v135
	goto L36
L40:
	;
	v132 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v122+int32(-3)))))
	v142 = v132
	goto L36
L41:
	;
	v142 = int32(base.Ui32(v125) >> (uint(int32(3)) % 32))
	goto L36
L42:
	;
	v172 = base.B2i32(v154 == int32(0))
	goto L1
L43:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_shr(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if base.Ui32(int32(31)) < base.Ui32(l1) {
		v13 = l1 + int32(-32)
		v14 = v6
		v15 = int32(0)
	} else {
		v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		v13 = l1
		v14 = v9
		v15 = v6
	}
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = int32(base.Ui32(v15) >> (uint(v13) % 32))
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v15<<(uint(int32(32)-v13)%32) | int32(base.Ui32(v14)>>(uint(v13)%32))
	return
}
func F_shutdownIOThread(m *base.Module, l0 int32) {
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
	var v20 int32
	_ = v20
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v36 int32
	_ = v36
	var v44 int32
	_ = v44
	var v62 int32
	_ = v62
	v6 = m.G0
	v8 = v6 - int32(32)
	m.G0 = v8
	v10 = F___get_tp(m)
	mBase = m.M
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0<<(uint(int32(2))%32))+uint32(_c_F_shutdownIOThread[0])))
	if v15 == int32(0) {
		m.G0 = v8 + int32(32)
		return
	} else {
		if v15 == v10 {
			m.G0 = v8 + int32(32)
			return
		} else {
			v20 = *(*int32)(unsafe.Add(mBase, _c_F_shutdownIOThread[1]))
			if l0 < v20 {
			} else {
			}
			v29 = int32(28)
			v31 = *(*int32)(unsafe.Add(mBase, _c_F_shutdownIOThread[2]))
			if int32(3) < v31 {
				F_spscFree(m, l0*int32(192)+int32(_a_F_shutdownIOThread_0))
				mBase = m.M
				v62 = m.ExcPending
				if v62 != 0 {
					return
				} else {
					m.G0 = v8 + int32(32)
					return
				}
			} else {
				v36 = F___strerror_l(m, v29, v29)
				mBase = m.M
				*(*int32)(unsafe.Add(mBase, uint32(v8)+20)) = v36
				*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = v15
				F__serverLog(m, int32(3), int32(_a_F_shutdownIOThread_1), v8+int32(16))
				mBase = m.M
				v44 = m.ExcPending
				if v44 != 0 {
					return
				} else {
					F_spscFree(m, l0*int32(192)+int32(_a_F_shutdownIOThread_0))
					mBase = m.M
					v62 = m.ExcPending
					if v62 != 0 {
						return
					} else {
						m.G0 = v8 + int32(32)
						return
					}
				}
			}
		}
	}
}
func F_sigKillChildHandler(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v61 int32
	_ = v61
	var v65 int32
	_ = v65
	var v68 int64
	_ = v68
	var v69 int32
	_ = v69
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v89 int32
	_ = v89
	var v94 int32
	_ = v94
	v5 = *(*int32)(unsafe.Add(mBase, _c_F_sigKillChildHandler[0]))
	if v5 == int32(4) {
		v8 = int32(1)
	} else {
		v8 = int32(3)
	}
	v9 = int32(_a_F_sigKillChildHandler_0)
	v13 = m.G0
	v15 = v13 - int32(80)
	m.G0 = v15
	v20 = *(*int32)(unsafe.Add(mBase, _c_F_sigKillChildHandler[1]))
	if v8&int32(255) < v20 {
	} else {
		v23 = *(*int32)(unsafe.Add(mBase, _c_F_sigKillChildHandler[2]))
		v24 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23))))
		if v24&int32(255) != 0 {
			if v24&int32(255) != 0 {
				*(*int32)(unsafe.Add(mBase, uint32(v15))) = int32(420)
				v35 = F_open(m, v23, int32(1089), v15)
				mBase = m.M
				if v35 == int32(-1) {
				} else {
					v38 = v35
					if v8&int32(1024) == int32(0) {
						v46 = v15 + int32(16)
						v48 = F_getpid(m)
						mBase = m.M
						v50 = F_ll2string(m, v46, int32(64), base.I64_extend_i32_s(v48))
						mBase = m.M
						v55 = F_strlen(m, v46)
						mBase = m.M
						v56 = F_write(m, v38, v46, v55)
						mBase = m.M
						if v56 == int32(-1) {
						} else {
							v61 = F_write(m, v38, int32(_a_F_sigKillChildHandler_1), int32(17))
							mBase = m.M
							if v61 == int32(-1) {
							} else {
								v65 = v15 + int32(16)
								v68 = F___time(m, int32(0))
								mBase = m.M
								v69 = F_ll2string(m, v65, int32(64), v68)
								mBase = m.M
								v74 = F_strlen(m, v65)
								mBase = m.M
								v75 = F_write(m, v38, v65, v74)
								mBase = m.M
								if v75 == int32(-1) {
								} else {
									v80 = F_write(m, v38, int32(_a_F_sigKillChildHandler_2), int32(2))
									mBase = m.M
									if v80 == int32(-1) {
									} else {
										v83 = F_strlen(m, v9)
										mBase = m.M
										v84 = F_write(m, v38, v9, v83)
										mBase = m.M
										if v84 == int32(-1) {
										} else {
											v89 = F_write(m, v38, int32(_a_F_sigKillChildHandler_3), int32(1))
											mBase = m.M
										}
									}
								}
							}
						}
					} else {
						v43 = F_strlen(m, v9)
						mBase = m.M
						v44 = F_write(m, v38, v9, v43)
						mBase = m.M
					}
					if v24&int32(255) == int32(0) {
					} else {
						v94 = F_close(m, v38)
						mBase = m.M
					}
				}
			} else {
				v38 = int32(1)
				if v8&int32(1024) == int32(0) {
					v46 = v15 + int32(16)
					v48 = F_getpid(m)
					mBase = m.M
					v50 = F_ll2string(m, v46, int32(64), base.I64_extend_i32_s(v48))
					mBase = m.M
					v55 = F_strlen(m, v46)
					mBase = m.M
					v56 = F_write(m, v38, v46, v55)
					mBase = m.M
					if v56 == int32(-1) {
					} else {
						v61 = F_write(m, v38, int32(_a_F_sigKillChildHandler_1), int32(17))
						mBase = m.M
						if v61 == int32(-1) {
						} else {
							v65 = v15 + int32(16)
							v68 = F___time(m, int32(0))
							mBase = m.M
							v69 = F_ll2string(m, v65, int32(64), v68)
							mBase = m.M
							v74 = F_strlen(m, v65)
							mBase = m.M
							v75 = F_write(m, v38, v65, v74)
							mBase = m.M
							if v75 == int32(-1) {
							} else {
								v80 = F_write(m, v38, int32(_a_F_sigKillChildHandler_2), int32(2))
								mBase = m.M
								if v80 == int32(-1) {
								} else {
									v83 = F_strlen(m, v9)
									mBase = m.M
									v84 = F_write(m, v38, v9, v83)
									mBase = m.M
									if v84 == int32(-1) {
									} else {
										v89 = F_write(m, v38, int32(_a_F_sigKillChildHandler_3), int32(1))
										mBase = m.M
									}
								}
							}
						}
					}
				} else {
					v43 = F_strlen(m, v9)
					mBase = m.M
					v44 = F_write(m, v38, v9, v43)
					mBase = m.M
				}
				if v24&int32(255) == int32(0) {
				} else {
					v94 = F_close(m, v38)
					mBase = m.M
				}
			}
		} else {
			v28 = *(*int32)(unsafe.Add(mBase, _c_F_sigKillChildHandler[3]))
			if v28 != 0 {
			} else {
				if v24&int32(255) != 0 {
					*(*int32)(unsafe.Add(mBase, uint32(v15))) = int32(420)
					v35 = F_open(m, v23, int32(1089), v15)
					mBase = m.M
					if v35 == int32(-1) {
					} else {
						v38 = v35
						if v8&int32(1024) == int32(0) {
							v46 = v15 + int32(16)
							v48 = F_getpid(m)
							mBase = m.M
							v50 = F_ll2string(m, v46, int32(64), base.I64_extend_i32_s(v48))
							mBase = m.M
							v55 = F_strlen(m, v46)
							mBase = m.M
							v56 = F_write(m, v38, v46, v55)
							mBase = m.M
							if v56 == int32(-1) {
							} else {
								v61 = F_write(m, v38, int32(_a_F_sigKillChildHandler_1), int32(17))
								mBase = m.M
								if v61 == int32(-1) {
								} else {
									v65 = v15 + int32(16)
									v68 = F___time(m, int32(0))
									mBase = m.M
									v69 = F_ll2string(m, v65, int32(64), v68)
									mBase = m.M
									v74 = F_strlen(m, v65)
									mBase = m.M
									v75 = F_write(m, v38, v65, v74)
									mBase = m.M
									if v75 == int32(-1) {
									} else {
										v80 = F_write(m, v38, int32(_a_F_sigKillChildHandler_2), int32(2))
										mBase = m.M
										if v80 == int32(-1) {
										} else {
											v83 = F_strlen(m, v9)
											mBase = m.M
											v84 = F_write(m, v38, v9, v83)
											mBase = m.M
											if v84 == int32(-1) {
											} else {
												v89 = F_write(m, v38, int32(_a_F_sigKillChildHandler_3), int32(1))
												mBase = m.M
											}
										}
									}
								}
							}
						} else {
							v43 = F_strlen(m, v9)
							mBase = m.M
							v44 = F_write(m, v38, v9, v43)
							mBase = m.M
						}
						if v24&int32(255) == int32(0) {
						} else {
							v94 = F_close(m, v38)
							mBase = m.M
						}
					}
				} else {
					v38 = int32(1)
					if v8&int32(1024) == int32(0) {
						v46 = v15 + int32(16)
						v48 = F_getpid(m)
						mBase = m.M
						v50 = F_ll2string(m, v46, int32(64), base.I64_extend_i32_s(v48))
						mBase = m.M
						v55 = F_strlen(m, v46)
						mBase = m.M
						v56 = F_write(m, v38, v46, v55)
						mBase = m.M
						if v56 == int32(-1) {
						} else {
							v61 = F_write(m, v38, int32(_a_F_sigKillChildHandler_1), int32(17))
							mBase = m.M
							if v61 == int32(-1) {
							} else {
								v65 = v15 + int32(16)
								v68 = F___time(m, int32(0))
								mBase = m.M
								v69 = F_ll2string(m, v65, int32(64), v68)
								mBase = m.M
								v74 = F_strlen(m, v65)
								mBase = m.M
								v75 = F_write(m, v38, v65, v74)
								mBase = m.M
								if v75 == int32(-1) {
								} else {
									v80 = F_write(m, v38, int32(_a_F_sigKillChildHandler_2), int32(2))
									mBase = m.M
									if v80 == int32(-1) {
									} else {
										v83 = F_strlen(m, v9)
										mBase = m.M
										v84 = F_write(m, v38, v9, v83)
										mBase = m.M
										if v84 == int32(-1) {
										} else {
											v89 = F_write(m, v38, int32(_a_F_sigKillChildHandler_3), int32(1))
											mBase = m.M
										}
									}
								}
							}
						}
					} else {
						v43 = F_strlen(m, v9)
						mBase = m.M
						v44 = F_write(m, v38, v9, v43)
						mBase = m.M
					}
					if v24&int32(255) == int32(0) {
					} else {
						v94 = F_close(m, v38)
						mBase = m.M
					}
				}
			}
		}
	}
	m.G0 = v15 + int32(80)
	F__Exit(m, int32(255))
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_sigemptyset(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	*(*int64)(unsafe.Add(mBase, uint32(l0))) = int64(0)
	return int32(0)
}
func F_signal(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v13 int32
	_ = v13
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v29 int32
	_ = v29
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v52 int32
	_ = v52
	v4 = m.G0
	v6 = v4 - int32(288)
	m.G0 = v6
	v13 = F__emscripten_memset_bulkmem(m, v6+int32(12), base.I32_extend8_s(int32(0)), int32(136))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v6)+140)) = int32(268435456)
	*(*int32)(unsafe.Add(mBase, uint32(v6)+8)) = l1
	v18 = v6 + int32(8)
	v20 = v6 + int32(148)
	if base.Ui32(l0) < base.Ui32(int32(65)) {
		if v20 == int32(0) {
		} else {
			v29 = int32(140)
			v34 = F___memcpy(m, v20, l0*v29+int32(9117184), v29)
			mBase = m.M
		}
		if v18 == int32(0) {
		} else {
			v37 = int32(140)
			v42 = F___memcpy(m, l0*v37+int32(9117184), v18, v37)
			mBase = m.M
		}
		v44 = int32(0)
	} else {
		v23 = F___errno_location(m)
		mBase = m.M
		*(*int32)(unsafe.Add(mBase, uint32(v23))) = int32(28)
		v44 = int32(-1)
	}
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v6)+148))
	m.G0 = v6 + int32(288)
	if v44 < int32(0) {
		v52 = int32(-1)
	} else {
		v52 = v45
	}
	return v52
}
func F_signalDeletedKeyAsReady(m *base.Module, l0 int32, l1 int32, l2 int32) {
	var v6 int32
	_ = v6
	F_signalKeyAsReadyLogic(m, l0, l1, l2, int32(1))
	v6 = m.ExcPending
	if v6 != 0 {
		return
	} else {
		return
	}
}
func F_simpleStringCallback(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v8 < int32(0) {
		v32 = v7
	} else {
		v13 = l0 + v8<<(uint(int32(2))%32)
		v14 = *(*int32)(unsafe.Add(mBase, uint32(v13)+16))
		v15 = int32(1)
		v16 = v14 + v15
		*(*int32)(unsafe.Add(mBase, uint32(v13)+16)) = v16
		v18 = *(*int32)(unsafe.Add(mBase, uint32(v13)+1040))
		if v18 != v15 {
			v32 = v7
		} else {
			v23 = *(*int32)(unsafe.Add(mBase, uint32(v7)+8))
			*(*int32)(unsafe.Add(mBase, uint32(v23)+8)) = int32(3)
			*(*float64)(unsafe.Add(mBase, uint32(v23))) = base.F64_convert_i32_u(v16)
			v27 = *(*int32)(unsafe.Add(mBase, uint32(v7)+8))
			*(*int32)(unsafe.Add(mBase, uint32(v7)+8)) = v27 + int32(16)
			v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			v32 = v31
		}
	}
	v36 = F_lua_checkstack(m, v32, int32(3))
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		return
	} else {
		if v36 != 0 {
			v41 = int32(0)
			F_lua_createtable(m, v32, v41, v41)
			mBase = m.M
			v44 = m.ExcPending
			if v44 != 0 {
				return
			} else {
				v45 = m.G3
				F_lua_pushstring(m, v32, v45+int32(_a_F_simpleStringCallback_0))
				mBase = m.M
				v49 = m.ExcPending
				if v49 != 0 {
					return
				} else {
					F_lua_pushlstring(m, v32, l1, l2)
					mBase = m.M
					v51 = m.ExcPending
					if v51 != 0 {
						return
					} else {
						F_lua_settable(m, v32, int32(-3))
						mBase = m.M
						v54 = m.ExcPending
						if v54 != 0 {
							return
						} else {
							F_processCollectionElementEnd(m, l0)
							mBase = m.M
							v56 = m.ExcPending
							if v56 != 0 {
								return
							} else {
								return
							}
						}
					}
				}
			}
		} else {
			F__serverPanic_2(m, int32(927))
			mBase = m.M
			v40 = m.ExcPending
			if v40 != 0 {
				return
			} else {
				base.Wasm_trap_unreachable()
				for {
				}
			}
		}
	}
}
func F_singlevar(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
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
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v9 == int32(285) {
		v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
		F_luaX_next(m, l0)
		mBase = m.M
		v27 = m.ExcPending
		if v27 != 0 {
			return
		} else {
			v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
			v30 = F_singlevaraux(m, v28, v25, l1, int32(1))
			mBase = m.M
			v31 = m.ExcPending
			if v31 != 0 {
				return
			} else {
				if v30 != int32(8) {
					m.G0 = v7 + int32(16)
					return
				} else {
					v34 = F_luaK_stringK(m, v28, v25)
					mBase = m.M
					v35 = m.ExcPending
					if v35 != 0 {
						return
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(l1)+8)) = v34
						m.G0 = v7 + int32(16)
						return
					}
				}
			}
		}
	} else {
		v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
		v14 = F_luaX_token2str(m, l0, int32(285))
		mBase = m.M
		v15 = m.ExcPending
		if v15 != 0 {
			return
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v7))) = v14
			v17 = m.G3
			v20 = F_luaO_pushfstring(m, v12, v17+int32(_a_F_singlevar_0), v7)
			mBase = m.M
			v21 = m.ExcPending
			if v21 != 0 {
				return
			} else {
				F_luaX_syntaxerror(m, l0, v20)
				mBase = m.M
				v23 = m.ExcPending
				if v23 != 0 {
					return
				} else {
					v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
					F_luaX_next(m, l0)
					mBase = m.M
					v27 = m.ExcPending
					if v27 != 0 {
						return
					} else {
						v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
						v30 = F_singlevaraux(m, v28, v25, l1, int32(1))
						mBase = m.M
						v31 = m.ExcPending
						if v31 != 0 {
							return
						} else {
							if v30 != int32(8) {
								m.G0 = v7 + int32(16)
								return
							} else {
								v34 = F_luaK_stringK(m, v28, v25)
								mBase = m.M
								v35 = m.ExcPending
								if v35 != 0 {
									return
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(l1)+8)) = v34
									m.G0 = v7 + int32(16)
									return
								}
							}
						}
					}
				}
			}
		}
	}
}
func F_sinterCardCommand(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v31 int32
	_ = v31
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
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
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v97 int32
	_ = v97
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v107 int32
	_ = v107
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
	var v115 int32
	_ = v115
	v2 = int32(0)
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	*(*int32)(unsafe.Add(mBase, uint32(v9)+12)) = v2
	*(*int32)(unsafe.Add(mBase, uint32(v9)+8)) = v2
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v16 = *(*int32)(unsafe.Add(mBase, uint32(v15)+4))
	v22 = F_getRangeLongFromObjectOrReply(m, l0, v16, int32(1), int32(2147483647), v9+int32(12), int32(_a_F_sinterCardCommand_0))
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v9 + int32(16)
	return
L2:
	;
	return
L3:
	;
	if v22 != 0 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v9)+12))
	v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v25+int32(-2) < v24 {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	F_addReplyError(m, l0, int32(_a_F_sinterCardCommand_1))
	mBase = m.M
	v115 = m.ExcPending
	if v115 != 0 {
		goto L2
	} else {
		goto L31
	}
L6:
	;
	v31 = v24
	goto L8
L7:
	;
	v107 = *(*int32)(unsafe.Add(mBase, uint32(v9)+12))
	v110 = *(*int32)(unsafe.Add(mBase, uint32(v9)+8))
	F_sinterGenericCommand(m, l0, v35+int32(8), v107, int32(0), int32(1), v110)
	mBase = m.M
	v112 = m.ExcPending
	if v112 != 0 {
		goto L2
	} else {
		goto L30
	}
L8:
	;
	v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v37 = v31 + int32(2)
	v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v38 <= v37 {
		goto L7
	} else {
		goto L10
	}
L10:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v35+v37<<(uint(int32(2))%32))))
	v44 = F_objectGetVal(m, v43)
	mBase = m.M
	v45 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v46 = int32(_a_F_sinterCardCommand_2)
	v49 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v44))))
	if v49 != 0 {
		goto L15
	} else {
		goto L16
	}
L11:
	;
	v93 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v97 = *(*int32)(unsafe.Add(mBase, uint32(v93+v86<<(uint(int32(2))%32))))
	v101 = F_getPositiveLongFromObjectOrReply(m, l0, v97, v9+int32(8), int32(_a_F_sinterCardCommand_3))
	mBase = m.M
	v102 = m.ExcPending
	if v102 != 0 {
		goto L2
	} else {
		goto L28
	}
L12:
	;
	v90 = *(*int32)(unsafe.Add(mBase, _c_F_sinterCardCommand[0]))
	F_addReplyErrorObject(m, l0, v90)
	mBase = m.M
	v92 = m.ExcPending
	if v92 != 0 {
		goto L2
	} else {
		goto L27
	}
L13:
	;
	if v81-v83 != 0 {
		goto L12
	} else {
		goto L25
	}
L14:
	;
	v81 = F_tolower(m, v77)
	mBase = m.M
	v82 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v78))))
	v83 = F_tolower(m, v82)
	mBase = m.M
	goto L13
L15:
	;
	v51 = v44
	v52 = v46
	v53 = v49
	goto L18
L16:
	;
	v77 = int32(0)
	v78 = v46
	goto L14
L17:
	;
	v77 = v74 & int32(255)
	v78 = v73
	goto L14
L18:
	;
	v55 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v52))))
	if v55 == int32(0) {
		v73 = v52
		v74 = v53
		goto L17
	} else {
		goto L20
	}
L19:
	;
	v73 = v67
	v74 = int32(0)
	goto L17
L20:
	;
	v59 = v53 & int32(255)
	if v59 == v55 {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	v66 = int32(1)
	v67 = v52 + v66
	v68 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v51)+1)))
	if v68 != 0 {
		v51 = v51 + v66
		v52 = v67
		v53 = v68
		goto L18
	} else {
		goto L24
	}
L22:
	;
	v61 = F_tolower(m, v59)
	mBase = m.M
	v62 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v52))))
	v63 = F_tolower(m, v62)
	mBase = m.M
	if v61 == v63 {
		goto L21
	} else {
		goto L23
	}
L23:
	;
	v65 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v51))))
	v73 = v52
	v74 = v65
	goto L17
L24:
	;
	goto L19
L25:
	;
	v86 = v31 + int32(3)
	if v45 != v86 {
		goto L11
	} else {
		goto L26
	}
L26:
	;
	goto L12
L27:
	;
	goto L1
L28:
	;
	if v101 == int32(0) {
		v31 = v37
		goto L8
	} else {
		goto L29
	}
L29:
	;
	goto L1
L30:
	;
	goto L1
L31:
	;
	goto L1
}
func F_sinterstoreCommand(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v13 int32
	_ = v13
	v2 = int32(0)
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v9 = *(*int32)(unsafe.Add(mBase, uint32(v3)+4))
	F_sinterGenericCommand(m, l0, v3+int32(8), v6+int32(-2), v9, v2, v2)
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return
	} else {
		return
	}
}
func F_siphash(m *base.Module, l0 int32, l1 int32, l2 int32) int64 {
	mBase := m.M
	_ = mBase
	var v11 int64
	_ = v11
	var v13 int64
	_ = v13
	var v14 int64
	_ = v14
	var v16 int64
	_ = v16
	var v18 int64
	_ = v18
	var v20 int64
	_ = v20
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v30 int64
	_ = v30
	var v31 int64
	_ = v31
	var v32 int64
	_ = v32
	var v33 int64
	_ = v33
	var v37 int64
	_ = v37
	var v38 int64
	_ = v38
	var v39 int64
	_ = v39
	var v40 int64
	_ = v40
	var v43 int64
	_ = v43
	var v44 int64
	_ = v44
	var v47 int64
	_ = v47
	var v50 int64
	_ = v50
	var v53 int64
	_ = v53
	var v55 int64
	_ = v55
	var v56 int64
	_ = v56
	var v58 int64
	_ = v58
	var v59 int64
	_ = v59
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v66 int64
	_ = v66
	var v67 int64
	_ = v67
	var v68 int64
	_ = v68
	var v69 int64
	_ = v69
	var v74 int64
	_ = v74
	var v75 int64
	_ = v75
	var v79 int64
	_ = v79
	var v80 int64
	_ = v80
	var v84 int64
	_ = v84
	var v85 int64
	_ = v85
	var v89 int64
	_ = v89
	var v90 int64
	_ = v90
	var v94 int64
	_ = v94
	var v95 int64
	_ = v95
	var v99 int64
	_ = v99
	var v100 int64
	_ = v100
	var v104 int64
	_ = v104
	var v105 int64
	_ = v105
	var v107 int64
	_ = v107
	var v108 int64
	_ = v108
	var v109 int64
	_ = v109
	var v111 int64
	_ = v111
	var v112 int64
	_ = v112
	var v113 int64
	_ = v113
	var v115 int64
	_ = v115
	var v116 int64
	_ = v116
	var v118 int64
	_ = v118
	var v119 int64
	_ = v119
	var v122 int64
	_ = v122
	var v124 int64
	_ = v124
	var v125 int64
	_ = v125
	var v130 int64
	_ = v130
	var v131 int64
	_ = v131
	var v135 int64
	_ = v135
	var v137 int64
	_ = v137
	var v138 int64
	_ = v138
	var v141 int64
	_ = v141
	var v142 int64
	_ = v142
	var v147 int64
	_ = v147
	var v148 int64
	_ = v148
	var v151 int64
	_ = v151
	var v157 int64
	_ = v157
	var v161 int64
	_ = v161
	var v165 int64
	_ = v165
	v11 = *(*int64)(unsafe.Add(mBase, uint32(l2)))
	v13 = v11 ^ int64(8317987319222330741)
	v14 = *(*int64)(unsafe.Add(mBase, uint32(l2)+8))
	v16 = v14 ^ int64(7237128888997146477)
	v18 = v11 ^ int64(7816392313619706465)
	v20 = v14 ^ int64(8387220255154660723)
	v24 = l1 & int32(7)
	v25 = l0 + l1 - v24
	if l0 == v25 {
		v63 = l0
		v66 = v18
		v67 = v13
		v68 = v20
		v69 = v16
	} else {
		v27 = l0
		v30 = v18
		v31 = v13
		v32 = v20
		v33 = v16
		for {
			v37 = *(*int64)(unsafe.Add(mBase, uint32(v27)))
			v38 = v37 ^ v32
			v39 = v38 + v30
			v40 = v31 + v33
			v43 = v40 ^ base.I64_rotl(v33, int64(13))
			v44 = v39 + v43
			v47 = v44 ^ base.I64_rotl(v43, int64(17))
			v50 = base.I64_rotl(v38, int64(16)) ^ v39
			v53 = int64(32)
			v55 = v50 + base.I64_rotl(v40, v53)
			v56 = base.I64_rotl(v50, int64(21)) ^ v55
			v58 = base.I64_rotl(v44, v53)
			v59 = v55 ^ v37
			v61 = v27 + int32(8)
			if v61 != v25 {
				v27 = v61
				v30 = v58
				v31 = v59
				v32 = v56
				v33 = v47
				continue
			} else {
				break
			}
			break
		}
		v63 = v25
		v66 = v58
		v67 = v59
		v68 = v56
		v69 = v47
	}
	v74 = base.I64_extend_i32_u(l1) << (uint(int64(56)) % 64)
	switch v24 {
	default:
		v107 = v74
	case 1:
		v104 = v74
		v105 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v63))))
		v107 = v104 | v105
	case 2:
		v99 = v74
		v100 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v63)+1)))
		v104 = v100<<(uint(int64(8))%64) | v99
		v105 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v63))))
		v107 = v104 | v105
	case 3:
		v94 = v74
		v95 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v63)+2)))
		v99 = v95<<(uint(int64(16))%64) | v94
		v100 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v63)+1)))
		v104 = v100<<(uint(int64(8))%64) | v99
		v105 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v63))))
		v107 = v104 | v105
	case 4:
		v89 = v74
		v90 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v63)+3)))
		v94 = v90<<(uint(int64(24))%64) | v89
		v95 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v63)+2)))
		v99 = v95<<(uint(int64(16))%64) | v94
		v100 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v63)+1)))
		v104 = v100<<(uint(int64(8))%64) | v99
		v105 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v63))))
		v107 = v104 | v105
	case 5:
		v84 = v74
		v85 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v63)+4)))
		v89 = v85<<(uint(int64(32))%64) | v84
		v90 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v63)+3)))
		v94 = v90<<(uint(int64(24))%64) | v89
		v95 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v63)+2)))
		v99 = v95<<(uint(int64(16))%64) | v94
		v100 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v63)+1)))
		v104 = v100<<(uint(int64(8))%64) | v99
		v105 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v63))))
		v107 = v104 | v105
	case 6:
		v79 = v74
		v80 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v63)+5)))
		v84 = v80<<(uint(int64(40))%64) | v79
		v85 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v63)+4)))
		v89 = v85<<(uint(int64(32))%64) | v84
		v90 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v63)+3)))
		v94 = v90<<(uint(int64(24))%64) | v89
		v95 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v63)+2)))
		v99 = v95<<(uint(int64(16))%64) | v94
		v100 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v63)+1)))
		v104 = v100<<(uint(int64(8))%64) | v99
		v105 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v63))))
		v107 = v104 | v105
	case 7:
		v75 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v63)+6)))
		v79 = v75<<(uint(int64(48))%64) | v74
		v80 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v63)+5)))
		v84 = v80<<(uint(int64(40))%64) | v79
		v85 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v63)+4)))
		v89 = v85<<(uint(int64(32))%64) | v84
		v90 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v63)+3)))
		v94 = v90<<(uint(int64(24))%64) | v89
		v95 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v63)+2)))
		v99 = v95<<(uint(int64(16))%64) | v94
		v100 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v63)+1)))
		v104 = v100<<(uint(int64(8))%64) | v99
		v105 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v63))))
		v107 = v104 | v105
	}
	v108 = v107 ^ v68
	v109 = int64(16)
	v111 = v108 + v66
	v112 = base.I64_rotl(v108, v109) ^ v111
	v113 = int64(21)
	v115 = v67 + v69
	v116 = int64(32)
	v118 = v112 + base.I64_rotl(v115, v116)
	v119 = base.I64_rotl(v112, v113) ^ v118
	v122 = int64(13)
	v124 = v115 ^ base.I64_rotl(v69, v122)
	v125 = v111 + v124
	v130 = base.I64_rotl(v125, v116) ^ int64(255) + v119
	v131 = base.I64_rotl(v119, v109) ^ v130
	v135 = int64(17)
	v137 = v125 ^ base.I64_rotl(v124, v135)
	v138 = v118 ^ v107 + v137
	v141 = base.I64_rotl(v138, v116) + v131
	v142 = base.I64_rotl(v131, v113) ^ v141
	v147 = v138 ^ base.I64_rotl(v137, v122)
	v148 = v147 + v130
	v151 = base.I64_rotl(v148, v116) + v142
	v157 = base.I64_rotl(v147, v135) ^ v148
	v161 = base.I64_rotl(v157, v122) ^ (v157 + v141)
	v165 = v161 + v151
	return base.I64_rotl(base.I64_rotl(v142, v109)^v151, v113) ^ base.I64_rotl(v161, v135) ^ base.I64_rotl(v165, v116) ^ v165
}
func F_siprintf(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
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
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	*(*int32)(unsafe.Add(mBase, uint32(v7)+12)) = l2
	v10 = F_vsiprintf(m, l0, l1, l2)
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return int32(0)
	} else {
		m.G0 = v7 + int32(16)
		return v10
	}
}
func F_sismemberCommand(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(v6)+4))
	v9 = *(*int32)(unsafe.Add(mBase, _c_F_sismemberCommand[0]))
	v10 = F_lookupKeyReadOrReply(m, l0, v7, v9)
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return
	} else {
		if v10 == int32(0) {
			return
		} else {
			v15 = F_checkType(m, l0, v10, int32(2))
			mBase = m.M
			v16 = m.ExcPending
			if v16 != 0 {
				return
			} else {
				if v15 != 0 {
					return
				} else {
					v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
					v19 = *(*int32)(unsafe.Add(mBase, uint32(v18)+8))
					v20 = F_objectGetVal(m, v19)
					mBase = m.M
					v23 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20+int32(-1)))))
					switch v23 & int32(7) {
					case 0:
						v40 = int32(base.Ui32(v23) >> (uint(int32(3)) % 32))
					case 1:
						v30 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20+int32(-3)))))
						v40 = v30
					case 2:
						v33 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v20+int32(-5)))))
						v40 = v33
					case 3:
						v36 = *(*int32)(unsafe.Add(mBase, uint32(v20+int32(-9))))
						v40 = v36
					case 4:
						v39 = *(*int32)(unsafe.Add(mBase, uint32(v20+int32(-17))))
						v40 = v39
					default:
						v40 = int32(0)
					}
					v46 = F_setTypeIsMemberAux(m, v10, v20, v40, int64(0), int32(1))
					mBase = m.M
					v47 = m.ExcPending
					if v47 != 0 {
						return
					} else {
						if v46 != 0 {
							v48 = int32(16)
						} else {
							v48 = int32(12)
						}
						v50 = *(*int32)(unsafe.Add(mBase, uint32(v48)+uint32(_c_F_sismemberCommand[1])))
						F_addReply(m, l0, v50)
						mBase = m.M
						v52 = m.ExcPending
						if v52 != 0 {
							return
						} else {
							return
						}
					}
				}
			}
		}
	}
}
func F_skip_sep(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v18 int32
	_ = v18
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
	var v34 int32
	_ = v34
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v46 int32
	_ = v46
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v57 int32
	_ = v57
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v70 int32
	_ = v70
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	F_save(m, l0, v6)
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
	*(*int32)(unsafe.Add(mBase, uint32(v11))) = v12 + int32(-1)
	if v12 == int32(0) {
		goto L4
	} else {
		goto L5
	}
L3:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v25
	v28 = int32(0)
	if v25 != int32(61) {
		v62 = v25
		v63 = v28
		goto L7
	} else {
		goto L8
	}
L4:
	;
	v23 = F_luaZ_fill(m, v11)
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		goto L1
	} else {
		goto L6
	}
L5:
	;
	v18 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+4)) = v18 + int32(1)
	v22 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18))))
	v25 = v22
	goto L3
L6:
	;
	v25 = v23
	goto L3
L7:
	;
	if v62 == v6 {
		goto L17
	} else {
		goto L18
	}
L8:
	;
	v34 = v28
	goto L9
L9:
	;
	F_save(m, l0, int32(61))
	mBase = m.M
	v38 = m.ExcPending
	if v38 != 0 {
		goto L1
	} else {
		goto L11
	}
L10:
	;
	v62 = v53
	v63 = v57
	goto L7
L11:
	;
	v39 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v39)))
	*(*int32)(unsafe.Add(mBase, uint32(v39))) = v40 + int32(-1)
	if v40 == int32(0) {
		goto L13
	} else {
		goto L14
	}
L12:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v53
	v57 = v34 + int32(1)
	if v53 == int32(61) {
		v34 = v57
		goto L9
	} else {
		goto L16
	}
L13:
	;
	v51 = F_luaZ_fill(m, v39)
	mBase = m.M
	v52 = m.ExcPending
	if v52 != 0 {
		goto L1
	} else {
		goto L15
	}
L14:
	;
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v39)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v39)+4)) = v46 + int32(1)
	v50 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v46))))
	v53 = v50
	goto L12
L15:
	;
	v53 = v51
	goto L12
L16:
	;
	goto L10
L17:
	;
	v70 = v63 + int32(2)
	goto L19
L18:
	;
	v70 = base.B2i32(v63 == int32(0))
	goto L19
L19:
	;
	return v70
}
func F_slowlogCommand(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v55 int32
	_ = v55
	var v56 int64
	_ = v56
	var v61 int64
	_ = v61
	var v66 int64
	_ = v66
	var v69 int64
	_ = v69
	var v72 int64
	_ = v72
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v95 int32
	_ = v95
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v113 int32
	_ = v113
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
	var v129 int32
	_ = v129
	var v131 int32
	_ = v131
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v139 int32
	_ = v139
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v151 int32
	_ = v151
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v157 int32
	_ = v157
	var v161 int32
	_ = v161
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v201 int32
	_ = v201
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v207 int32
	_ = v207
	var v211 int32
	_ = v211
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v237 int32
	_ = v237
	var v240 int32
	_ = v240
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v248 int32
	_ = v248
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v254 int32
	_ = v254
	var v255 int32
	_ = v255
	var v257 int32
	_ = v257
	var v260 int32
	_ = v260
	var v262 int32
	_ = v262
	var v264 int32
	_ = v264
	var v265 int64
	_ = v265
	var v267 int32
	_ = v267
	v4 = m.G0
	v6 = v4 - int32(48)
	m.G0 = v6
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v8 != int32(2) {
		v190 = v8
		goto L3
	} else {
		goto L4
	}
L1:
	;
	m.G0 = v6 + int32(48)
	return
L2:
	;
	v264 = *(*int32)(unsafe.Add(mBase, _c_F_slowlogCommand[0]))
	v265 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v264)+20)))
	F_addReplyLongLong(m, l0, v265)
	mBase = m.M
	v267 = m.ExcPending
	if v267 != 0 {
		goto L19
	} else {
		goto L79
	}
L3:
	;
	if v190&int32(-2) != int32(2) {
		goto L57
	} else {
		goto L58
	}
L4:
	;
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
	v13 = F_objectGetVal(m, v12)
	mBase = m.M
	v14 = int32(_a_F_slowlogCommand_0)
	v17 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13))))
	if v17 != 0 {
		goto L8
	} else {
		goto L9
	}
L5:
	;
	v76 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v76 != int32(2) {
		v190 = v76
		goto L3
	} else {
		goto L21
	}
L6:
	;
	if v49-v51 != 0 {
		goto L5
	} else {
		goto L18
	}
L7:
	;
	v49 = F_tolower(m, v45)
	mBase = m.M
	v50 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v46))))
	v51 = F_tolower(m, v50)
	mBase = m.M
	goto L6
L8:
	;
	v19 = v13
	v20 = v14
	v21 = v17
	goto L11
L9:
	;
	v45 = int32(0)
	v46 = v14
	goto L7
L10:
	;
	v45 = v42 & int32(255)
	v46 = v41
	goto L7
L11:
	;
	v23 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20))))
	if v23 == int32(0) {
		v41 = v20
		v42 = v21
		goto L10
	} else {
		goto L13
	}
L12:
	;
	v41 = v35
	v42 = int32(0)
	goto L10
L13:
	;
	v27 = v21 & int32(255)
	if v27 == v23 {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v34 = int32(1)
	v35 = v20 + v34
	v36 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19)+1)))
	if v36 != 0 {
		v19 = v19 + v34
		v20 = v35
		v21 = v36
		goto L11
	} else {
		goto L17
	}
L15:
	;
	v29 = F_tolower(m, v27)
	mBase = m.M
	v30 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20))))
	v31 = F_tolower(m, v30)
	mBase = m.M
	if v29 == v31 {
		goto L14
	} else {
		goto L16
	}
L16:
	;
	v33 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19))))
	v41 = v20
	v42 = v33
	goto L10
L17:
	;
	goto L12
L18:
	;
	v55 = int32(0)
	v56 = *(*int64)(unsafe.Add(mBase, _c_F_slowlogCommand[1]))
	*(*int64)(unsafe.Add(mBase, uint32(v6+int32(32)))) = v56
	v61 = *(*int64)(unsafe.Add(mBase, _c_F_slowlogCommand[2]))
	*(*int64)(unsafe.Add(mBase, uint32(v6+int32(24)))) = v61
	v66 = *(*int64)(unsafe.Add(mBase, _c_F_slowlogCommand[3]))
	*(*int64)(unsafe.Add(mBase, uint32(v6+int32(16)))) = v66
	v69 = *(*int64)(unsafe.Add(mBase, _c_F_slowlogCommand[4]))
	*(*int64)(unsafe.Add(mBase, uint32(v6)+8)) = v69
	v72 = *(*int64)(unsafe.Add(mBase, _c_F_slowlogCommand[5]))
	*(*int64)(unsafe.Add(mBase, uint32(v6))) = v72
	F_addReplyHelp(m, l0, v6)
	mBase = m.M
	v75 = m.ExcPending
	if v75 != 0 {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	return
L20:
	;
	goto L1
L21:
	;
	v79 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v80 = *(*int32)(unsafe.Add(mBase, uint32(v79)+4))
	v81 = F_objectGetVal(m, v80)
	mBase = m.M
	v82 = int32(_a_F_slowlogCommand_1)
	v85 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v81))))
	if v85 != 0 {
		goto L25
	} else {
		goto L26
	}
L22:
	;
	v142 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v142 != int32(2) {
		v190 = v142
		goto L3
	} else {
		goto L43
	}
L23:
	;
	if v117-v119 != 0 {
		goto L22
	} else {
		goto L35
	}
L24:
	;
	v117 = F_tolower(m, v113)
	mBase = m.M
	v118 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v114))))
	v119 = F_tolower(m, v118)
	mBase = m.M
	goto L23
L25:
	;
	v87 = v81
	v88 = v82
	v89 = v85
	goto L28
L26:
	;
	v113 = int32(0)
	v114 = v82
	goto L24
L27:
	;
	v113 = v110 & int32(255)
	v114 = v109
	goto L24
L28:
	;
	v91 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v88))))
	if v91 == int32(0) {
		v109 = v88
		v110 = v89
		goto L27
	} else {
		goto L30
	}
L29:
	;
	v109 = v103
	v110 = int32(0)
	goto L27
L30:
	;
	v95 = v89 & int32(255)
	if v95 == v91 {
		goto L31
	} else {
		goto L32
	}
L31:
	;
	v102 = int32(1)
	v103 = v88 + v102
	v104 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v87)+1)))
	if v104 != 0 {
		v87 = v87 + v102
		v88 = v103
		v89 = v104
		goto L28
	} else {
		goto L34
	}
L32:
	;
	v97 = F_tolower(m, v95)
	mBase = m.M
	v98 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v88))))
	v99 = F_tolower(m, v98)
	mBase = m.M
	if v97 == v99 {
		goto L31
	} else {
		goto L33
	}
L33:
	;
	v101 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v87))))
	v109 = v88
	v110 = v101
	goto L27
L34:
	;
	goto L29
L35:
	;
	v122 = *(*int32)(unsafe.Add(mBase, _c_F_slowlogCommand[0]))
	v123 = *(*int32)(unsafe.Add(mBase, uint32(v122)+20))
	if v123 == int32(0) {
		goto L36
	} else {
		goto L37
	}
L36:
	;
	v139 = *(*int32)(unsafe.Add(mBase, _c_F_slowlogCommand[6]))
	F_addReply(m, l0, v139)
	mBase = m.M
	v141 = m.ExcPending
	if v141 != 0 {
		goto L19
	} else {
		goto L42
	}
L37:
	;
	v128 = v122
	goto L38
L38:
	;
	v129 = *(*int32)(unsafe.Add(mBase, uint32(v128)+4))
	F_listDelNode(m, v128, v129)
	mBase = m.M
	v131 = m.ExcPending
	if v131 != 0 {
		goto L19
	} else {
		goto L40
	}
L39:
	;
	goto L36
L40:
	;
	v133 = *(*int32)(unsafe.Add(mBase, _c_F_slowlogCommand[0]))
	v134 = *(*int32)(unsafe.Add(mBase, uint32(v133)+20))
	if v134 != 0 {
		v128 = v133
		goto L38
	} else {
		goto L41
	}
L41:
	;
	goto L39
L42:
	;
	goto L1
L43:
	;
	v145 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v146 = *(*int32)(unsafe.Add(mBase, uint32(v145)+4))
	v147 = F_objectGetVal(m, v146)
	mBase = m.M
	v148 = int32(_a_F_slowlogCommand_2)
	v151 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v147))))
	if v151 != 0 {
		goto L46
	} else {
		goto L47
	}
L44:
	;
	if v183-v185 == int32(0) {
		goto L2
	} else {
		goto L56
	}
L45:
	;
	v183 = F_tolower(m, v179)
	mBase = m.M
	v184 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v180))))
	v185 = F_tolower(m, v184)
	mBase = m.M
	goto L44
L46:
	;
	v153 = v147
	v154 = v148
	v155 = v151
	goto L49
L47:
	;
	v179 = int32(0)
	v180 = v148
	goto L45
L48:
	;
	v179 = v176 & int32(255)
	v180 = v175
	goto L45
L49:
	;
	v157 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v154))))
	if v157 == int32(0) {
		v175 = v154
		v176 = v155
		goto L48
	} else {
		goto L51
	}
L50:
	;
	v175 = v169
	v176 = int32(0)
	goto L48
L51:
	;
	v161 = v155 & int32(255)
	if v161 == v157 {
		goto L52
	} else {
		goto L53
	}
L52:
	;
	v168 = int32(1)
	v169 = v154 + v168
	v170 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v153)+1)))
	if v170 != 0 {
		v153 = v153 + v168
		v154 = v169
		v155 = v170
		goto L49
	} else {
		goto L55
	}
L53:
	;
	v163 = F_tolower(m, v161)
	mBase = m.M
	v164 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v154))))
	v165 = F_tolower(m, v164)
	mBase = m.M
	if v163 == v165 {
		goto L52
	} else {
		goto L54
	}
L54:
	;
	v167 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v153))))
	v175 = v154
	v176 = v167
	goto L48
L55:
	;
	goto L50
L56:
	;
	v189 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v190 = v189
	goto L3
L57:
	;
	F_addReplySubcommandSyntaxError(m, l0)
	mBase = m.M
	v262 = m.ExcPending
	if v262 != 0 {
		goto L19
	} else {
		goto L78
	}
L58:
	;
	v195 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v196 = *(*int32)(unsafe.Add(mBase, uint32(v195)+4))
	v197 = F_objectGetVal(m, v196)
	mBase = m.M
	v198 = int32(_a_F_slowlogCommand_3)
	v201 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v197))))
	if v201 != 0 {
		goto L61
	} else {
		goto L62
	}
L59:
	;
	if v233-v235 != 0 {
		goto L57
	} else {
		goto L71
	}
L60:
	;
	v233 = F_tolower(m, v229)
	mBase = m.M
	v234 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v230))))
	v235 = F_tolower(m, v234)
	mBase = m.M
	goto L59
L61:
	;
	v203 = v197
	v204 = v198
	v205 = v201
	goto L64
L62:
	;
	v229 = int32(0)
	v230 = v198
	goto L60
L63:
	;
	v229 = v226 & int32(255)
	v230 = v225
	goto L60
L64:
	;
	v207 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v204))))
	if v207 == int32(0) {
		v225 = v204
		v226 = v205
		goto L63
	} else {
		goto L66
	}
L65:
	;
	v225 = v219
	v226 = int32(0)
	goto L63
L66:
	;
	v211 = v205 & int32(255)
	if v211 == v207 {
		goto L67
	} else {
		goto L68
	}
L67:
	;
	v218 = int32(1)
	v219 = v204 + v218
	v220 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v203)+1)))
	if v220 != 0 {
		v203 = v203 + v218
		v204 = v219
		v205 = v220
		goto L64
	} else {
		goto L70
	}
L68:
	;
	v213 = F_tolower(m, v211)
	mBase = m.M
	v214 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v204))))
	v215 = F_tolower(m, v214)
	mBase = m.M
	if v213 == v215 {
		goto L67
	} else {
		goto L69
	}
L69:
	;
	v217 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v203))))
	v225 = v204
	v226 = v217
	goto L63
L70:
	;
	goto L65
L71:
	;
	v237 = int32(10)
	*(*int32)(unsafe.Add(mBase, uint32(v6))) = v237
	v240 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v240 != int32(3) {
		v257 = v237
		goto L72
	} else {
		goto L73
	}
L72:
	;
	F_commandlogGetReply(m, l0, int32(0), v257)
	mBase = m.M
	v260 = m.ExcPending
	if v260 != 0 {
		goto L19
	} else {
		goto L77
	}
L73:
	;
	v243 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v244 = *(*int32)(unsafe.Add(mBase, uint32(v243)+8))
	v248 = F_getRangeLongFromObjectOrReply(m, l0, v244, int32(-1), int32(2147483647), v6, int32(_a_F_slowlogCommand_4))
	mBase = m.M
	v249 = m.ExcPending
	if v249 != 0 {
		goto L19
	} else {
		goto L74
	}
L74:
	;
	if v248 != 0 {
		goto L1
	} else {
		goto L75
	}
L75:
	;
	v250 = *(*int32)(unsafe.Add(mBase, uint32(v6)))
	if v250 != int32(-1) {
		v257 = v250
		goto L72
	} else {
		goto L76
	}
L76:
	;
	v254 = *(*int32)(unsafe.Add(mBase, _c_F_slowlogCommand[0]))
	v255 = *(*int32)(unsafe.Add(mBase, uint32(v254)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v6))) = v255
	v257 = v255
	goto L72
L77:
	;
	goto L1
L78:
	;
	goto L1
L79:
	;
	goto L1
}
func F_socket(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v13 int32
	_ = v13
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v39 int32
	_ = v39
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v56 int32
	_ = v56
	v4 = int32(0)
	v6 = m.G0
	v8 = v6 - int32(32)
	m.G0 = v8
	v13 = m.Env.X__syscall_socket(m, l0, l1, l2, v4, v4, v4)
	mBase = m.M
	if v13 == int32(-28) {
		if l1&int32(526336) == int32(0) {
			v48 = v13
		} else {
			v24 = int32(0)
			v27 = m.Env.X__syscall_socket(m, l0, l1&int32(-526337), l2, v24, v24, v24)
			mBase = m.M
			if v27 < v24 {
				v48 = v27
			} else {
				if l1&int32(524288) == int32(0) {
				} else {
					*(*int64)(unsafe.Add(mBase, uint32(v8)+16)) = int64(1)
					v39 = m.Env.X__syscall_fcntl64(m, v27, int32(2), v8+int32(16))
					mBase = m.M
				}
				if l1&int32(2048) == int32(0) {
					v48 = v27
				} else {
					*(*int64)(unsafe.Add(mBase, uint32(v8))) = int64(2048)
					v47 = m.Env.X__syscall_fcntl64(m, v27, int32(4), v8)
					mBase = m.M
					v48 = v27
				}
			}
		}
	} else {
		if v13 != int32(-66) {
			v48 = v13
		} else {
			if l1&int32(526336) == int32(0) {
				v48 = v13
			} else {
				v24 = int32(0)
				v27 = m.Env.X__syscall_socket(m, l0, l1&int32(-526337), l2, v24, v24, v24)
				mBase = m.M
				if v27 < v24 {
					v48 = v27
				} else {
					if l1&int32(524288) == int32(0) {
					} else {
						*(*int64)(unsafe.Add(mBase, uint32(v8)+16)) = int64(1)
						v39 = m.Env.X__syscall_fcntl64(m, v27, int32(2), v8+int32(16))
						mBase = m.M
					}
					if l1&int32(2048) == int32(0) {
						v48 = v27
					} else {
						*(*int64)(unsafe.Add(mBase, uint32(v8))) = int64(2048)
						v47 = m.Env.X__syscall_fcntl64(m, v27, int32(4), v8)
						mBase = m.M
						v48 = v27
					}
				}
			}
		}
	}
	if base.Ui32(v48) < base.Ui32(int32(-4095)) {
		v56 = v48
	} else {
		v51 = F___errno_location(m)
		mBase = m.M
		*(*int32)(unsafe.Add(mBase, uint32(v51))) = int32(0) - v48
		v56 = int32(-1)
	}
	m.G0 = v8 + int32(32)
	return v56
}
func F_sortCompare(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v8 float64
	_ = v8
	var v9 float64
	_ = v9
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
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
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	v7 = *(*int32)(unsafe.Add(mBase, _c_F_sortCompare[0]))
	if v7 != 0 {
		v25 = *(*int32)(unsafe.Add(mBase, _c_F_sortCompare[1]))
		if v25 == int32(0) {
			v49 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
			v50 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			v52 = *(*int32)(unsafe.Add(mBase, _c_F_sortCompare[2]))
			if v52 == int32(0) {
				v57 = F_collateStringObjects(m, v50, v49)
				mBase = m.M
				v58 = m.ExcPending
				if v58 != 0 {
					return int32(0)
				} else {
					v60 = v57
					v67 = *(*int32)(unsafe.Add(mBase, _c_F_sortCompare[3]))
					if v67 != 0 {
						v68 = int32(0) - v60
					} else {
						v68 = v60
					}
					return v68
				}
			} else {
				v55 = F_compareStringObjects(m, v50, v49)
				mBase = m.M
				v56 = m.ExcPending
				if v56 != 0 {
					return int32(0)
				} else {
					v60 = v55
					v67 = *(*int32)(unsafe.Add(mBase, _c_F_sortCompare[3]))
					if v67 != 0 {
						v68 = int32(0) - v60
					} else {
						v68 = v60
					}
					return v68
				}
			}
		} else {
			v28 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
			v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			if v29 != 0 {
				if v28 != 0 {
					v39 = *(*int32)(unsafe.Add(mBase, _c_F_sortCompare[2]))
					if v39 == int32(0) {
						v44 = F_objectGetVal(m, v29)
						mBase = m.M
						v45 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
						v46 = F_objectGetVal(m, v45)
						mBase = m.M
						v47 = F___get_tp(m)
						mBase = m.M
						v48 = F___strcoll_l(m, v44, v46, v46)
						mBase = m.M
						v60 = v48
						v67 = *(*int32)(unsafe.Add(mBase, _c_F_sortCompare[3]))
						if v67 != 0 {
							v68 = int32(0) - v60
						} else {
							v68 = v60
						}
						return v68
					} else {
						v42 = F_compareStringObjects(m, v29, v28)
						mBase = m.M
						v43 = m.ExcPending
						if v43 != 0 {
							return int32(0)
						} else {
							v60 = v42
							v67 = *(*int32)(unsafe.Add(mBase, _c_F_sortCompare[3]))
							if v67 != 0 {
								v68 = int32(0) - v60
							} else {
								v68 = v60
							}
							return v68
						}
					}
				} else {
					v33 = int32(1)
					v34 = int32(0)
					if v29 == v34 {
						v37 = int32(0)
					} else {
						v37 = v33
					}
					v60 = v37
					v67 = *(*int32)(unsafe.Add(mBase, _c_F_sortCompare[3]))
					if v67 != 0 {
						v68 = int32(0) - v60
					} else {
						v68 = v60
					}
					return v68
				}
			} else {
				v33 = int32(-1)
				v34 = v28
				if v29 == v34 {
					v37 = int32(0)
				} else {
					v37 = v33
				}
				v60 = v37
				v67 = *(*int32)(unsafe.Add(mBase, _c_F_sortCompare[3]))
				if v67 != 0 {
					v68 = int32(0) - v60
				} else {
					v68 = v60
				}
				return v68
			}
		}
	} else {
		v8 = *(*float64)(unsafe.Add(mBase, uint32(l0)+8))
		v9 = *(*float64)(unsafe.Add(mBase, uint32(l1)+8))
		if base.F64_gt(v8, v9) == int32(0) {
			if base.F64_lt(v8, v9) == int32(0) {
				v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
				v19 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
				v20 = F_compareStringObjects(m, v18, v19)
				mBase = m.M
				v23 = m.ExcPending
				if v23 != 0 {
					return int32(0)
				} else {
					v60 = v20
					v67 = *(*int32)(unsafe.Add(mBase, _c_F_sortCompare[3]))
					if v67 != 0 {
						v68 = int32(0) - v60
					} else {
						v68 = v60
					}
					return v68
				}
			} else {
				v60 = int32(-1)
				v67 = *(*int32)(unsafe.Add(mBase, _c_F_sortCompare[3]))
				if v67 != 0 {
					v68 = int32(0) - v60
				} else {
					v68 = v60
				}
				return v68
			}
		} else {
			v60 = int32(1)
			v67 = *(*int32)(unsafe.Add(mBase, _c_F_sortCompare[3]))
			if v67 != 0 {
				v68 = int32(0) - v60
			} else {
				v68 = v60
			}
			return v68
		}
	}
}
func F_sort_1(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
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
	var v82 int32
	_ = v82
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v91 int32
	_ = v91
	var v95 int32
	_ = v95
	var v99 int32
	_ = v99
	var v113 int32
	_ = v113
	F_luaL_checktype(m, l0, int32(1), int32(5))
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		v10 = m.G3
		v12 = F_lua_objlen(m, l0, int32(1))
		mBase = m.M
		v13 = m.ExcPending
		if v13 != 0 {
			return int32(0)
		} else {
			F_luaL_checkstack(m, l0, int32(40), v10+int32(_a_F_sort_1_0))
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
				if v76 < int32(1) {
					v87 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
					v88 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
					v91 = v88 + int32(32)
					if base.Ui32(v91) <= base.Ui32(v87) {
					} else {
						v95 = v87
						for {
							*(*int32)(unsafe.Add(mBase, uint32(v95)+8)) = int32(0)
							v99 = v95 + int32(16)
							if base.Ui32(v99) < base.Ui32(v91) {
								v95 = v99
								continue
							} else {
								break
							}
							break
						}
					}
					*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v91
					F_auxsort(m, l0, int32(1), v12)
					mBase = m.M
					v113 = m.ExcPending
					if v113 != 0 {
						return int32(0)
					} else {
						return int32(0)
					}
				} else {
					F_luaL_checktype(m, l0, int32(2), int32(6))
					mBase = m.M
					v82 = m.ExcPending
					if v82 != 0 {
						return int32(0)
					} else {
						v87 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
						v88 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
						v91 = v88 + int32(32)
						if base.Ui32(v91) <= base.Ui32(v87) {
						} else {
							v95 = v87
							for {
								*(*int32)(unsafe.Add(mBase, uint32(v95)+8)) = int32(0)
								v99 = v95 + int32(16)
								if base.Ui32(v99) < base.Ui32(v91) {
									v95 = v99
									continue
								} else {
									break
								}
								break
							}
						}
						*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v91
						F_auxsort(m, l0, int32(1), v12)
						mBase = m.M
						v113 = m.ExcPending
						if v113 != 0 {
							return int32(0)
						} else {
							return int32(0)
						}
					}
				}
			}
		}
	}
}
func F_sort_gp_desc(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 float64
	_ = v6
	var v7 float64
	_ = v7
	var v10 int32
	_ = v10
	v6 = *(*float64)(unsafe.Add(mBase, uint32(l0)+16))
	v7 = *(*float64)(unsafe.Add(mBase, uint32(l1)+16))
	if base.F64_gt(v6, v7) != 0 {
		v10 = int32(-1)
	} else {
		v10 = base.F64_ne(v6, v7)
	}
	return v10
}
func F_sparklineRenderRange(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32 {
	mBase := m.M
	_ = mBase
	var v25 float64
	_ = v25
	var v26 float64
	_ = v26
	var v27 float64
	_ = v27
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v38 float64
	_ = v38
	var v42 float64
	_ = v42
	var v45 int32
	_ = v45
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v63 int32
	_ = v63
	var v79 int32
	_ = v79
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v95 float64
	_ = v95
	var v96 int32
	_ = v96
	var v103 int32
	_ = v103
	var v119 int32
	_ = v119
	var v124 int32
	_ = v124
	var v125 float64
	_ = v125
	var v126 float64
	_ = v126
	var v131 float64
	_ = v131
	var v132 float64
	_ = v132
	var v134 float64
	_ = v134
	var v140 int32
	_ = v140
	var v142 int32
	_ = v142
	var v144 float64
	_ = v144
	var v150 int32
	_ = v150
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v156 int32
	_ = v156
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v170 int32
	_ = v170
	var v174 int32
	_ = v174
	var v184 int32
	_ = v184
	var v187 int32
	_ = v187
	var v191 int32
	_ = v191
	var v196 int32
	_ = v196
	var v198 int32
	_ = v198
	var v202 int32
	_ = v202
	var v208 int32
	_ = v208
	var v211 int32
	_ = v211
	var v217 int32
	_ = v217
	var v221 int32
	_ = v221
	var v223 int32
	_ = v223
	var v231 int32
	_ = v231
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v243 int32
	_ = v243
	var v247 int32
	_ = v247
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v260 int32
	_ = v260
	var v261 int32
	_ = v261
	var v264 int32
	_ = v264
	var v267 int32
	_ = v267
	var v292 int32
	_ = v292
	v25 = *(*float64)(unsafe.Add(mBase, uint32(l1)+24))
	v26 = *(*float64)(unsafe.Add(mBase, uint32(l1)+16))
	v27 = base.F64_sub(v25, v26)
	v28 = F_valkey_malloc(m, l4)
	mBase = m.M
	v31 = m.ExcPending
	if v31 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v33 = l5 & int32(2)
	if v33 == int32(0) {
		goto L4
	} else {
		goto L5
	}
L3:
	;
	v45 = F__emscripten_memset_bulkmem(m, v28, base.I32_extend8_s(int32(32)), l4)
	mBase = m.M
	goto L7
L4:
	;
	if base.F64_ne(v27, float64(0)) != 0 {
		v42 = v27
		goto L3
	} else {
		goto L6
	}
L5:
	;
	v38 = F_log(m, base.F64_add(v27, float64(1)))
	mBase = m.M
	v42 = v38
	goto L3
L6:
	;
	v42 = float64(1)
	goto L3
L7:
	;
	if l4 < int32(1) {
		v267 = l0
		goto L8
	} else {
		goto L9
	}
L8:
	;
	F_valkey_free(m, v45)
	mBase = m.M
	v292 = m.ExcPending
	if v292 != 0 {
		goto L1
	} else {
		goto L66
	}
L9:
	;
	v51 = l5 & int32(1)
	if v51 != 0 {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	v52 = int32(_a_F_sparklineRenderRange_0)
	goto L12
L11:
	;
	v52 = int32(_a_F_sparklineRenderRange_1)
	goto L12
L12:
	;
	v54 = l2 * int32(3)
	v55 = int32(-1)
	v63 = l0
	v79 = int32(0)
	goto L13
L13:
	;
	v87 = v79 + (l2 ^ v55)
	v88 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v95 = *(*float64)(unsafe.Add(mBase, uint32(l1)+16))
	v96 = int32(0)
	v103 = v96
	v119 = v96
	goto L16
L15:
	;
	v256 = F_sdscatlen(m, v63, v45, l4)
	mBase = m.M
	v257 = m.ExcPending
	if v257 != 0 {
		goto L1
	} else {
		goto L63
	}
L16:
	;
	v124 = v88 + l3<<(uint(int32(4))%32) + v103<<(uint(int32(4))%32)
	v125 = *(*float64)(unsafe.Add(mBase, uint32(v124)))
	v126 = base.F64_sub(v125, v95)
	if v33 == int32(0) {
		v132 = v126
		goto L18
	} else {
		goto L19
	}
L17:
	;
	if v243 == int32(0) {
		v267 = v63
		goto L8
	} else {
		goto L62
	}
L18:
	;
	if l2 <= v79 {
		goto L22
	} else {
		goto L23
	}
L19:
	;
	v131 = F_log(m, base.F64_add(v126, float64(1)))
	mBase = m.M
	v132 = v131
	goto L18
L20:
	;
	v247 = v103 + int32(1)
	if v247 != l4 {
		v103 = v247
		v119 = v243
		goto L16
	} else {
		goto L61
	}
L21:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v45+v103))) = uint8(v235)
	v243 = int32(1)
	goto L20
L22:
	;
	v170 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v170 == int32(0) {
		goto L40
	} else {
		goto L41
	}
L23:
	;
	v134 = base.F64_mul(v132, base.F64_convert_i32_s(v54))
	if base.F64_lt(base.F64_abs(v134), float64(2.147483648e+09)) == int32(0) {
		goto L25
	} else {
		goto L26
	}
L24:
	;
	v144 = base.F64_div(base.F64_convert_i32_s(v142), v42)
	if base.F64_lt(base.F64_abs(v144), float64(2.147483648e+09)) == int32(0) {
		goto L28
	} else {
		goto L29
	}
L25:
	;
	v142 = int32(-2147483648)
	goto L24
L26:
	;
	v140 = base.I32_trunc_f64_s(v134)
	v142 = v140
	goto L24
L27:
	;
	v153 = int32(0)
	if v153 < v152 {
		goto L31
	} else {
		goto L32
	}
L28:
	;
	v152 = int32(-2147483648)
	goto L27
L29:
	;
	v150 = base.I32_trunc_f64_s(v144)
	v152 = v150
	goto L27
L30:
	;
	v164 = int32(1)
	if v51 == int32(0) {
		v243 = v164
		goto L20
	} else {
		goto L38
	}
L31:
	;
	v156 = v152
	goto L33
L32:
	;
	v156 = v153
	goto L33
L33:
	;
	if v156 < v54 {
		goto L34
	} else {
		goto L35
	}
L34:
	;
	v158 = v156
	goto L36
L35:
	;
	v158 = v54 + v55
	goto L36
L36:
	;
	v159 = v158 + (l2+(v79^int32(-1)))*int32(-3)
	if base.Ui32(int32(2)) < base.Ui32(v159) {
		goto L30
	} else {
		goto L37
	}
L37:
	;
	v163 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v52+v159))))
	v235 = v163
	goto L21
L38:
	;
	if int32(2) < v159 {
		v235 = int32(124)
		goto L21
	} else {
		goto L39
	}
L39:
	;
	v243 = v164
	goto L20
L40:
	;
	v174 = *(*int32)(unsafe.Add(mBase, uint32(v124)+8))
	if v174 == int32(0) {
		v243 = v119
		goto L20
	} else {
		goto L43
	}
L41:
	;
	if v79 <= l2 {
		goto L15
	} else {
		goto L42
	}
L42:
	;
	goto L40
L43:
	;
	if v174&int32(3) == int32(0) {
		v198 = v174
		goto L46
	} else {
		goto L47
	}
L44:
	;
	if v231 <= v87 {
		v243 = v119
		goto L20
	} else {
		goto L60
	}
L45:
	;
	v231 = v223 - v174
	goto L44
L46:
	;
	v202 = v198
	goto L54
L47:
	;
	v184 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v174))))
	if v184 != 0 {
		goto L48
	} else {
		goto L49
	}
L48:
	;
	v187 = v174
	goto L50
L49:
	;
	v231 = v174 - v174
	goto L44
L50:
	;
	v191 = v187 + int32(1)
	if v191&int32(3) == int32(0) {
		v198 = v191
		goto L46
	} else {
		goto L52
	}
L52:
	;
	v196 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v191))))
	if v196 != 0 {
		v187 = v191
		goto L50
	} else {
		goto L53
	}
L53:
	;
	v223 = v191
	goto L45
L54:
	;
	v208 = *(*int32)(unsafe.Add(mBase, uint32(v202)))
	v211 = int32(-2139062144)
	if (int32(16843008)-v208|v208)&v211 == v211 {
		v202 = v202 + int32(4)
		goto L54
	} else {
		goto L56
	}
L55:
	;
	v217 = v202
	goto L57
L56:
	;
	goto L55
L57:
	;
	v221 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v217))))
	if v221 != 0 {
		v217 = v217 + int32(1)
		goto L57
	} else {
		goto L59
	}
L58:
	;
	v223 = v217
	goto L45
L59:
	;
	goto L58
L60:
	;
	v234 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v174+v87))))
	v235 = v234
	goto L21
L61:
	;
	goto L17
L62:
	;
	goto L15
L63:
	;
	v260 = F_sdscatlen(m, v256, int32(_a_F_sparklineRenderRange_2), int32(1))
	mBase = m.M
	v261 = m.ExcPending
	if v261 != 0 {
		goto L1
	} else {
		goto L64
	}
L64:
	;
	v264 = F__emscripten_memset_bulkmem(m, v45, base.I32_extend8_s(int32(32)), l4)
	mBase = m.M
	goto L65
L65:
	;
	v63 = v260
	v79 = v79 + int32(1)
	goto L13
L66:
	;
	return v267
}
func F_sparklineSequenceAddSample(m *base.Module, l0 int32, l1 float64, l2 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v17 float64
	_ = v17
	var v22 float64
	_ = v22
	var v27 int32
	_ = v27
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v46 int32
	_ = v46
	v4 = int32(0)
	if l2 == v4 {
		v14 = v4
		v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		if v15 != 0 {
			v17 = *(*float64)(unsafe.Add(mBase, uint32(l0)+16))
			if base.F64_lt(l1, v17) == int32(0) {
				v22 = *(*float64)(unsafe.Add(mBase, uint32(l0)+24))
				if base.F64_gt(l1, v22) == int32(0) {
				} else {
					*(*float64)(unsafe.Add(mBase, uint32(l0)+24)) = l1
				}
			} else {
				*(*float64)(unsafe.Add(mBase, uint32(l0)+16)) = l1
			}
		} else {
			*(*float64)(unsafe.Add(mBase, uint32(l0)+16)) = l1
			*(*float64)(unsafe.Add(mBase, uint32(l0)+24)) = l1
		}
		v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		v32 = F_valkey_realloc(m, v27, v15<<(uint(int32(4))%32)+int32(16))
		mBase = m.M
		v33 = m.ExcPending
		if v33 != 0 {
			return
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v32
			v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			v38 = v32 + v35<<(uint(int32(4))%32)
			*(*int32)(unsafe.Add(mBase, uint32(v38)+8)) = v14
			*(*float64)(unsafe.Add(mBase, uint32(v38))) = l1
			*(*int32)(unsafe.Add(mBase, uint32(l0))) = v35 + int32(1)
			if v14 == int32(0) {
			} else {
				v46 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
				*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v46 + int32(1)
			}
			return
		}
	} else {
		v9 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2))))
		if v9 == int32(0) {
			v14 = v4
			v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			if v15 != 0 {
				v17 = *(*float64)(unsafe.Add(mBase, uint32(l0)+16))
				if base.F64_lt(l1, v17) == int32(0) {
					v22 = *(*float64)(unsafe.Add(mBase, uint32(l0)+24))
					if base.F64_gt(l1, v22) == int32(0) {
					} else {
						*(*float64)(unsafe.Add(mBase, uint32(l0)+24)) = l1
					}
				} else {
					*(*float64)(unsafe.Add(mBase, uint32(l0)+16)) = l1
				}
			} else {
				*(*float64)(unsafe.Add(mBase, uint32(l0)+16)) = l1
				*(*float64)(unsafe.Add(mBase, uint32(l0)+24)) = l1
			}
			v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			v32 = F_valkey_realloc(m, v27, v15<<(uint(int32(4))%32)+int32(16))
			mBase = m.M
			v33 = m.ExcPending
			if v33 != 0 {
				return
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v32
				v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
				v38 = v32 + v35<<(uint(int32(4))%32)
				*(*int32)(unsafe.Add(mBase, uint32(v38)+8)) = v14
				*(*float64)(unsafe.Add(mBase, uint32(v38))) = l1
				*(*int32)(unsafe.Add(mBase, uint32(l0))) = v35 + int32(1)
				if v14 == int32(0) {
				} else {
					v46 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
					*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v46 + int32(1)
				}
				return
			}
		} else {
			v12 = F_zstrdup(m, l2)
			mBase = m.M
			v13 = m.ExcPending
			if v13 != 0 {
				return
			} else {
				v14 = v12
				v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
				if v15 != 0 {
					v17 = *(*float64)(unsafe.Add(mBase, uint32(l0)+16))
					if base.F64_lt(l1, v17) == int32(0) {
						v22 = *(*float64)(unsafe.Add(mBase, uint32(l0)+24))
						if base.F64_gt(l1, v22) == int32(0) {
						} else {
							*(*float64)(unsafe.Add(mBase, uint32(l0)+24)) = l1
						}
					} else {
						*(*float64)(unsafe.Add(mBase, uint32(l0)+16)) = l1
					}
				} else {
					*(*float64)(unsafe.Add(mBase, uint32(l0)+16)) = l1
					*(*float64)(unsafe.Add(mBase, uint32(l0)+24)) = l1
				}
				v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				v32 = F_valkey_realloc(m, v27, v15<<(uint(int32(4))%32)+int32(16))
				mBase = m.M
				v33 = m.ExcPending
				if v33 != 0 {
					return
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v32
					v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
					v38 = v32 + v35<<(uint(int32(4))%32)
					*(*int32)(unsafe.Add(mBase, uint32(v38)+8)) = v14
					*(*float64)(unsafe.Add(mBase, uint32(v38))) = l1
					*(*int32)(unsafe.Add(mBase, uint32(l0))) = v35 + int32(1)
					if v14 == int32(0) {
					} else {
						v46 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
						*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v46 + int32(1)
					}
					return
				}
			}
		}
	}
}
func F_spopWithCountCommand(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v60 int64
	_ = v60
	var v62 int32
	_ = v62
	var v67 int32
	_ = v67
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v119 int32
	_ = v119
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v130 int32
	_ = v130
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v156 int32
	_ = v156
	var v158 int32
	_ = v158
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v163 int32
	_ = v163
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v172 int32
	_ = v172
	var v183 int32
	_ = v183
	var v185 int32
	_ = v185
	var v187 int32
	_ = v187
	var v197 int32
	_ = v197
	var v202 int32
	_ = v202
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v219 int32
	_ = v219
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v241 int32
	_ = v241
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v247 int64
	_ = v247
	var v249 int32
	_ = v249
	var v250 int64
	_ = v250
	var v251 int32
	_ = v251
	var v252 int32
	_ = v252
	var v253 int32
	_ = v253
	var v259 int32
	_ = v259
	var v261 int32
	_ = v261
	var v262 int32
	_ = v262
	var v264 int32
	_ = v264
	var v266 int32
	_ = v266
	var v267 int32
	_ = v267
	var v273 int32
	_ = v273
	var v284 int32
	_ = v284
	var v286 int32
	_ = v286
	var v288 int32
	_ = v288
	var v298 int32
	_ = v298
	var v306 int32
	_ = v306
	var v307 int32
	_ = v307
	var v308 int32
	_ = v308
	var v309 int32
	_ = v309
	var v313 int32
	_ = v313
	var v315 int32
	_ = v315
	var v316 int32
	_ = v316
	var v318 int32
	_ = v318
	var v320 int32
	_ = v320
	var v321 int32
	_ = v321
	var v324 int32
	_ = v324
	var v331 int32
	_ = v331
	var v333 int32
	_ = v333
	var v347 int32
	_ = v347
	var v348 int32
	_ = v348
	var v349 int32
	_ = v349
	var v352 int32
	_ = v352
	var v353 int32
	_ = v353
	var v354 int32
	_ = v354
	var v355 int32
	_ = v355
	var v356 int32
	_ = v356
	var v358 int32
	_ = v358
	var v359 int32
	_ = v359
	var v360 int64
	_ = v360
	var v362 int32
	_ = v362
	var v363 int32
	_ = v363
	var v364 int32
	_ = v364
	var v365 int64
	_ = v365
	var v366 int32
	_ = v366
	var v367 int32
	_ = v367
	var v369 int32
	_ = v369
	var v370 int32
	_ = v370
	var v371 int32
	_ = v371
	var v373 int32
	_ = v373
	var v374 int32
	_ = v374
	var v375 int32
	_ = v375
	var v376 int32
	_ = v376
	var v381 int32
	_ = v381
	var v382 int32
	_ = v382
	var v385 int32
	_ = v385
	var v391 int32
	_ = v391
	var v399 int32
	_ = v399
	var v400 int32
	_ = v400
	var v405 int32
	_ = v405
	var v406 int32
	_ = v406
	var v407 int32
	_ = v407
	var v408 int64
	_ = v408
	var v410 int32
	_ = v410
	var v411 int32
	_ = v411
	var v416 int32
	_ = v416
	var v417 int32
	_ = v417
	var v418 int32
	_ = v418
	var v419 int32
	_ = v419
	var v423 int32
	_ = v423
	var v434 int32
	_ = v434
	var v438 int32
	_ = v438
	var v439 int32
	_ = v439
	var v441 int32
	_ = v441
	var v443 int32
	_ = v443
	var v456 int32
	_ = v456
	var v457 int32
	_ = v457
	var v464 int32
	_ = v464
	var v465 int32
	_ = v465
	var v478 int32
	_ = v478
	var v482 int32
	_ = v482
	var v483 int64
	_ = v483
	var v485 int32
	_ = v485
	var v486 int64
	_ = v486
	var v487 int32
	_ = v487
	var v488 int32
	_ = v488
	var v489 int32
	_ = v489
	var v491 int32
	_ = v491
	var v492 int32
	_ = v492
	var v493 int32
	_ = v493
	var v495 int32
	_ = v495
	var v501 int32
	_ = v501
	var v503 int32
	_ = v503
	var v504 int32
	_ = v504
	var v506 int32
	_ = v506
	var v508 int32
	_ = v508
	var v509 int32
	_ = v509
	var v515 int32
	_ = v515
	var v526 int32
	_ = v526
	var v528 int32
	_ = v528
	var v530 int32
	_ = v530
	var v540 int32
	_ = v540
	var v550 int32
	_ = v550
	var v551 int32
	_ = v551
	var v562 int32
	_ = v562
	var v567 int32
	_ = v567
	var v568 int32
	_ = v568
	var v569 int32
	_ = v569
	var v570 int32
	_ = v570
	var v574 int32
	_ = v574
	var v583 int32
	_ = v583
	var v587 int32
	_ = v587
	var v590 int32
	_ = v590
	var v591 int32
	_ = v591
	var v593 int32
	_ = v593
	var v595 int32
	_ = v595
	var v602 int32
	_ = v602
	var v613 int32
	_ = v613
	var v615 int32
	_ = v615
	var v617 int32
	_ = v617
	var v632 int32
	_ = v632
	var v633 int32
	_ = v633
	var v637 int32
	_ = v637
	var v638 int32
	_ = v638
	var v639 int32
	_ = v639
	var v641 int32
	_ = v641
	v2 = int32(0)
	v13 = m.G0
	v15 = v13 - int32(48)
	m.G0 = v15
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v18 = *(*int32)(unsafe.Add(mBase, uint32(v17)+8))
	v22 = F_getPositiveLongFromObjectOrReply(m, l0, v18, v15+int32(44), v2)
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v15 + int32(48)
	return
L2:
	;
	return
L3:
	;
	if v22 != 0 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v15)+44))
	v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v26 = *(*int32)(unsafe.Add(mBase, uint32(v25)+4))
	v28 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+224)))
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v28<<(uint(int32(2))%32))+uint32(_c_F_spopWithCountCommand[0])))
	v33 = F_lookupKeyWriteOrReply(m, l0, v26, v32)
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
		goto L2
	} else {
		goto L5
	}
L5:
	;
	if v33 == int32(0) {
		goto L1
	} else {
		goto L6
	}
L6:
	;
	v38 = F_checkType(m, l0, v33, int32(2))
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L2
	} else {
		goto L7
	}
L7:
	;
	if v38 != 0 {
		goto L1
	} else {
		goto L8
	}
L8:
	;
	if v24 != 0 {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	v48 = F_setTypeSize(m, v33)
	mBase = m.M
	v49 = m.ExcPending
	if v49 != 0 {
		goto L2
	} else {
		goto L12
	}
L10:
	;
	v41 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+224)))
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v41<<(uint(int32(2))%32))+uint32(_c_F_spopWithCountCommand[0])))
	F_addReply(m, l0, v45)
	mBase = m.M
	v47 = m.ExcPending
	if v47 != 0 {
		goto L2
	} else {
		goto L11
	}
L11:
	;
	goto L1
L12:
	;
	v52 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v52)+4))
	v54 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v54)+28))
	F_notifyKeyspaceEvent(m, int32(32), int32(_a_F_spopWithCountCommand_0), v53, v55)
	mBase = m.M
	v57 = m.ExcPending
	if v57 != 0 {
		goto L2
	} else {
		goto L13
	}
L13:
	;
	v58 = int32(_a_F_spopWithCountCommand_1)
	v60 = *(*int64)(unsafe.Add(mBase, _c_F_spopWithCountCommand[1]))
	if base.Ui32(v24) < base.Ui32(v48) {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v62 = v24
	goto L16
L15:
	;
	v62 = v48
	goto L16
L16:
	;
	*(*int64)(unsafe.Add(mBase, _c_F_spopWithCountCommand[1])) = v60 + base.I64_extend_i32_u(v62)
	if base.Ui32(v24) < base.Ui32(v48) {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	v108 = int32(1024)
	if base.Ui32(v24) < base.Ui32(v108) {
		goto L27
	} else {
		goto L28
	}
L18:
	;
	v67 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v71 = int32(0)
	F_sunionDiffGenericCommand(m, l0, v67+int32(4), int32(1), v71, v71)
	mBase = m.M
	v74 = m.ExcPending
	if v74 != 0 {
		goto L2
	} else {
		goto L19
	}
L19:
	;
	v75 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v76 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v77 = *(*int32)(unsafe.Add(mBase, uint32(v76)+4))
	v78 = F_dbDelete(m, v75, v77)
	mBase = m.M
	v79 = m.ExcPending
	if v79 != 0 {
		goto L2
	} else {
		goto L20
	}
L20:
	;
	v82 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v83 = *(*int32)(unsafe.Add(mBase, uint32(v82)+4))
	v84 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v85 = *(*int32)(unsafe.Add(mBase, uint32(v84)+28))
	F_notifyKeyspaceEvent(m, int32(4), int32(_a_F_spopWithCountCommand_2), v83, v85)
	mBase = m.M
	v87 = m.ExcPending
	if v87 != 0 {
		goto L2
	} else {
		goto L21
	}
L21:
	;
	v88 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v89 = *(*int32)(unsafe.Add(mBase, uint32(v88)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+4)) = v89
	v95 = *(*int32)(unsafe.Add(mBase, _c_F_spopWithCountCommand[2]))
	if v95 != 0 {
		goto L22
	} else {
		goto L23
	}
L22:
	;
	v96 = int32(244)
	goto L24
L23:
	;
	v96 = int32(240)
	goto L24
L24:
	;
	v98 = *(*int32)(unsafe.Add(mBase, uint32(v96)+uint32(_c_F_spopWithCountCommand[3])))
	*(*int32)(unsafe.Add(mBase, uint32(v15))) = v98
	F_rewriteClientCommandVector(m, l0, int32(2), v15)
	mBase = m.M
	v102 = m.ExcPending
	if v102 != 0 {
		goto L2
	} else {
		goto L25
	}
L25:
	;
	v103 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v104 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v105 = *(*int32)(unsafe.Add(mBase, uint32(v104)+4))
	F_signalModifiedKey(m, l0, v103, v105)
	mBase = m.M
	v107 = m.ExcPending
	if v107 != 0 {
		goto L2
	} else {
		goto L26
	}
L26:
	;
	goto L1
L27:
	;
	v111 = v24
	goto L29
L28:
	;
	v111 = v108
	goto L29
L29:
	;
	v112 = int32(2)
	v113 = v111 + v112
	v116 = F_valkey_malloc(m, v113<<(uint(v112)%32))
	mBase = m.M
	v117 = m.ExcPending
	if v117 != 0 {
		goto L2
	} else {
		goto L30
	}
L30:
	;
	v119 = *(*int32)(unsafe.Add(mBase, _c_F_spopWithCountCommand[4]))
	*(*int32)(unsafe.Add(mBase, uint32(v116))) = v119
	v121 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v122 = *(*int32)(unsafe.Add(mBase, uint32(v121)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v116)+4)) = v122
	F_addReplySetLen(m, l0, v24)
	mBase = m.M
	v125 = m.ExcPending
	if v125 != 0 {
		goto L2
	} else {
		goto L31
	}
L31:
	;
	v126 = v48 - v24
	if base.Ui32(v126*int32(5)) <= base.Ui32(v24) {
		goto L33
	} else {
		goto L34
	}
L32:
	;
	v587 = int32(2)
	if v583 == v587 {
		goto L131
	} else {
		goto L132
	}
L33:
	;
	v321 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v15)+20)) = v321
	v324 = *(*int32)(unsafe.Add(mBase, uint32(v33)))
	if v324&int32(240) == int32(176) {
		goto L75
	} else {
		goto L76
	}
L34:
	;
	v130 = *(*int32)(unsafe.Add(mBase, uint32(v33)))
	if v130&int32(240) == int32(176) {
		goto L35
	} else {
		goto L36
	}
L35:
	;
	v204 = F_objectGetVal(m, v33)
	mBase = m.M
	v205 = F_lpFirst(m, v204)
	mBase = m.M
	v206 = m.ExcPending
	if v206 != 0 {
		goto L2
	} else {
		goto L49
	}
L36:
	;
	v144 = int32(0)
	v145 = int32(2)
	goto L37
L37:
	;
	v152 = F_setTypePopRandom(m, v33)
	mBase = m.M
	v153 = m.ExcPending
	if v153 != 0 {
		goto L2
	} else {
		goto L39
	}
L39:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v116+v145<<(uint(int32(2))%32)))) = v152
	F_addReplyBulk(m, l0, v152)
	mBase = m.M
	v156 = m.ExcPending
	if v156 != 0 {
		goto L2
	} else {
		goto L40
	}
L40:
	;
	v158 = v145 + int32(1)
	if v158 != v113 {
		v197 = v158
		goto L41
	} else {
		goto L42
	}
L41:
	;
	v202 = v144 + int32(1)
	if v202 != v24 {
		v144 = v202
		v145 = v197
		goto L37
	} else {
		goto L48
	}
L42:
	;
	v160 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v161 = *(*int32)(unsafe.Add(mBase, uint32(v160)+28))
	v163 = *(*int32)(unsafe.Add(mBase, uint32(l0)+292))
	F_alsoPropagate(m, v161, v116, v113, int32(3), v163)
	mBase = m.M
	v165 = m.ExcPending
	if v165 != 0 {
		goto L2
	} else {
		goto L43
	}
L43:
	;
	v166 = int32(2)
	v172 = v166
	goto L44
L44:
	;
	v183 = *(*int32)(unsafe.Add(mBase, uint32(v116+v172<<(uint(int32(2))%32))))
	F_decrRefCount(m, v183)
	mBase = m.M
	v185 = m.ExcPending
	if v185 != 0 {
		goto L2
	} else {
		goto L46
	}
L45:
	;
	v197 = v166
	goto L41
L46:
	;
	v187 = v172 + int32(1)
	if v187 != v113 {
		v172 = v187
		goto L44
	} else {
		goto L47
	}
L47:
	;
	goto L45
L48:
	;
	v583 = v197
	goto L32
L49:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+40)) = int32(0)
	v211 = F_valkey_malloc(m, v24<<(uint(int32(2))%32))
	mBase = m.M
	v212 = m.ExcPending
	if v212 != 0 {
		goto L2
	} else {
		goto L50
	}
L50:
	;
	v219 = v205
	v222 = int32(0)
	v223 = int32(2)
	goto L51
L51:
	;
	v231 = F_lpNextRandom(m, v204, v219, v15+int32(40), v24-v222, int32(0))
	mBase = m.M
	v232 = m.ExcPending
	if v232 != 0 {
		goto L2
	} else {
		goto L55
	}
L52:
	;
	v315 = F_lpBatchDelete(m, v204, v211, v24)
	mBase = m.M
	v316 = m.ExcPending
	if v316 != 0 {
		goto L2
	} else {
		goto L71
	}
L53:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v116+v223<<(uint(int32(2))%32)))) = v253
	v259 = v223 + int32(1)
	if v259 != v113 {
		v298 = v259
		goto L62
	} else {
		goto L63
	}
L54:
	;
	v247 = *(*int64)(unsafe.Add(mBase, uint32(v15)+24))
	F_addReplyBulkLongLong(m, l0, v247)
	mBase = m.M
	v249 = m.ExcPending
	if v249 != 0 {
		goto L2
	} else {
		goto L60
	}
L55:
	;
	v237 = F_lpGetValue(m, v231, v15+int32(36), v15+int32(24))
	mBase = m.M
	v238 = m.ExcPending
	if v238 != 0 {
		goto L2
	} else {
		goto L56
	}
L56:
	;
	if v237 == int32(0) {
		goto L54
	} else {
		goto L57
	}
L57:
	;
	v241 = *(*int32)(unsafe.Add(mBase, uint32(v15)+36))
	F_addReplyBulkCBuffer(m, l0, v237, v241)
	mBase = m.M
	v243 = m.ExcPending
	if v243 != 0 {
		goto L2
	} else {
		goto L58
	}
L58:
	;
	v244 = *(*int32)(unsafe.Add(mBase, uint32(v15)+36))
	v245 = F_createStringObject_1(m, v237, v244)
	mBase = m.M
	v246 = m.ExcPending
	if v246 != 0 {
		goto L2
	} else {
		goto L59
	}
L59:
	;
	v253 = v245
	goto L53
L60:
	;
	v250 = *(*int64)(unsafe.Add(mBase, uint32(v15)+24))
	v251 = F_createStringObjectFromLongLong(m, v250)
	mBase = m.M
	v252 = m.ExcPending
	if v252 != 0 {
		goto L2
	} else {
		goto L61
	}
L61:
	;
	v253 = v251
	goto L53
L62:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v211+v222<<(uint(int32(2))%32)))) = v231
	v306 = F_lpNext(m, v204, v231)
	mBase = m.M
	v307 = m.ExcPending
	if v307 != 0 {
		goto L2
	} else {
		goto L69
	}
L63:
	;
	v261 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v262 = *(*int32)(unsafe.Add(mBase, uint32(v261)+28))
	v264 = *(*int32)(unsafe.Add(mBase, uint32(l0)+292))
	F_alsoPropagate(m, v262, v116, v113, int32(3), v264)
	mBase = m.M
	v266 = m.ExcPending
	if v266 != 0 {
		goto L2
	} else {
		goto L64
	}
L64:
	;
	v267 = int32(2)
	v273 = v267
	goto L65
L65:
	;
	v284 = *(*int32)(unsafe.Add(mBase, uint32(v116+v273<<(uint(int32(2))%32))))
	F_decrRefCount(m, v284)
	mBase = m.M
	v286 = m.ExcPending
	if v286 != 0 {
		goto L2
	} else {
		goto L67
	}
L66:
	;
	v298 = v267
	goto L62
L67:
	;
	v288 = v273 + int32(1)
	if v288 != v113 {
		v273 = v288
		goto L65
	} else {
		goto L68
	}
L68:
	;
	goto L66
L69:
	;
	v308 = *(*int32)(unsafe.Add(mBase, uint32(v15)+40))
	v309 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v15)+40)) = v308 + v309
	v313 = v222 + v309
	if v313 != v24 {
		v219 = v306
		v222 = v313
		v223 = v298
		goto L51
	} else {
		goto L70
	}
L70:
	;
	goto L52
L71:
	;
	F_valkey_free(m, v211)
	mBase = m.M
	v318 = m.ExcPending
	if v318 != 0 {
		goto L2
	} else {
		goto L72
	}
L72:
	;
	F_objectSetVal(m, v33, v315)
	mBase = m.M
	v320 = m.ExcPending
	if v320 != 0 {
		goto L2
	} else {
		goto L73
	}
L73:
	;
	v583 = v298
	goto L32
L74:
	;
	v456 = F_setTypeInitIterator(m, v33)
	mBase = m.M
	v457 = m.ExcPending
	if v457 != 0 {
		goto L2
	} else {
		goto L108
	}
L75:
	;
	v370 = F_createSetListpackObject(m)
	mBase = m.M
	v371 = m.ExcPending
	if v371 != 0 {
		goto L2
	} else {
		goto L90
	}
L76:
	;
	v331 = v321
	v333 = v126
	goto L77
L77:
	;
	v347 = F_setTypeRandomElement(m, v33, v15+int32(40), v15+int32(36), v15+int32(24))
	mBase = m.M
	v348 = m.ExcPending
	if v348 != 0 {
		goto L2
	} else {
		goto L79
	}
L79:
	;
	v349 = *(*int32)(unsafe.Add(mBase, uint32(v15)+40))
	if v331 != 0 {
		v358 = v331
		goto L80
	} else {
		goto L81
	}
L80:
	;
	v359 = *(*int32)(unsafe.Add(mBase, uint32(v15)+36))
	v360 = *(*int64)(unsafe.Add(mBase, uint32(v15)+24))
	v362 = base.B2i32(v347 == int32(2))
	v363 = F_setTypeAddAux(m, v358, v349, v359, v360, v362)
	mBase = m.M
	v364 = m.ExcPending
	if v364 != 0 {
		goto L2
	} else {
		goto L87
	}
L81:
	;
	if v349 == int32(0) {
		goto L83
	} else {
		goto L84
	}
L82:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+20)) = v356
	v358 = v356
	goto L80
L83:
	;
	v354 = F_createIntsetObject(m)
	mBase = m.M
	v355 = m.ExcPending
	if v355 != 0 {
		goto L2
	} else {
		goto L86
	}
L84:
	;
	v352 = F_createSetListpackObject(m)
	mBase = m.M
	v353 = m.ExcPending
	if v353 != 0 {
		goto L2
	} else {
		goto L85
	}
L85:
	;
	v356 = v352
	goto L82
L86:
	;
	v356 = v354
	goto L82
L87:
	;
	v365 = *(*int64)(unsafe.Add(mBase, uint32(v15)+24))
	v366 = F_setTypeRemoveAux(m, v33, v349, v359, v365, v362)
	mBase = m.M
	v367 = m.ExcPending
	if v367 != 0 {
		goto L2
	} else {
		goto L88
	}
L88:
	;
	v369 = v333 + int32(-1)
	if v369 != 0 {
		v331 = v358
		v333 = v369
		goto L77
	} else {
		goto L89
	}
L89:
	;
	goto L74
L90:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+20)) = v370
	v373 = F_objectGetVal(m, v33)
	mBase = m.M
	v374 = F_lpFirst(m, v373)
	mBase = m.M
	v375 = m.ExcPending
	if v375 != 0 {
		goto L2
	} else {
		goto L91
	}
L91:
	;
	v376 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v15)+36)) = v376
	v381 = F_valkey_malloc(m, v126<<(uint(int32(2))%32))
	mBase = m.M
	v382 = m.ExcPending
	if v382 != 0 {
		goto L2
	} else {
		goto L92
	}
L92:
	;
	if v126 != 0 {
		goto L94
	} else {
		goto L95
	}
L93:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+40)) = v434
	v438 = F_lpBatchDelete(m, v373, v381, v126)
	mBase = m.M
	v439 = m.ExcPending
	if v439 != 0 {
		goto L2
	} else {
		goto L103
	}
L94:
	;
	v385 = v374
	v391 = v376
	goto L96
L95:
	;
	v434 = v2
	goto L93
L96:
	;
	v399 = F_lpNextRandom(m, v373, v385, v15+int32(36), v126-v391, int32(0))
	mBase = m.M
	v400 = m.ExcPending
	if v400 != 0 {
		goto L2
	} else {
		goto L98
	}
L97:
	;
	v434 = v405
	goto L93
L98:
	;
	v405 = F_lpGetValue(m, v399, v15+int32(40), v15+int32(24))
	mBase = m.M
	v406 = m.ExcPending
	if v406 != 0 {
		goto L2
	} else {
		goto L99
	}
L99:
	;
	v407 = *(*int32)(unsafe.Add(mBase, uint32(v15)+40))
	v408 = *(*int64)(unsafe.Add(mBase, uint32(v15)+24))
	v410 = F_setTypeAddAux(m, v370, v405, v407, v408, int32(0))
	mBase = m.M
	v411 = m.ExcPending
	if v411 != 0 {
		goto L2
	} else {
		goto L100
	}
L100:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v381+v391<<(uint(int32(2))%32)))) = v399
	v416 = F_lpNext(m, v373, v399)
	mBase = m.M
	v417 = m.ExcPending
	if v417 != 0 {
		goto L2
	} else {
		goto L101
	}
L101:
	;
	v418 = *(*int32)(unsafe.Add(mBase, uint32(v15)+36))
	v419 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v15)+36)) = v418 + v419
	v423 = v391 + v419
	if v423 != v126 {
		v385 = v416
		v391 = v423
		goto L96
	} else {
		goto L102
	}
L102:
	;
	goto L97
L103:
	;
	F_valkey_free(m, v381)
	mBase = m.M
	v441 = m.ExcPending
	if v441 != 0 {
		goto L2
	} else {
		goto L104
	}
L104:
	;
	F_objectSetVal(m, v33, v438)
	mBase = m.M
	v443 = m.ExcPending
	if v443 != 0 {
		goto L2
	} else {
		goto L105
	}
L105:
	;
	goto L74
L106:
	;
	F_setTypeReleaseIterator(m, v456)
	mBase = m.M
	v567 = m.ExcPending
	if v567 != 0 {
		goto L2
	} else {
		goto L129
	}
L107:
	;
	v478 = int32(2)
	goto L111
L108:
	;
	v464 = F_setTypeNext(m, v456, v15+int32(40), v15+int32(36), v15+int32(24))
	mBase = m.M
	v465 = m.ExcPending
	if v465 != 0 {
		goto L2
	} else {
		goto L109
	}
L109:
	;
	if v464 != int32(-1) {
		goto L107
	} else {
		goto L110
	}
L110:
	;
	v562 = int32(2)
	goto L106
L111:
	;
	v482 = *(*int32)(unsafe.Add(mBase, uint32(v15)+40))
	if v482 != 0 {
		goto L114
	} else {
		goto L115
	}
L112:
	;
	v562 = v540
	goto L106
L113:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v116+v478<<(uint(int32(2))%32)))) = v495
	v501 = v478 + int32(1)
	if v501 != v113 {
		v540 = v501
		goto L120
	} else {
		goto L121
	}
L114:
	;
	v489 = *(*int32)(unsafe.Add(mBase, uint32(v15)+36))
	F_addReplyBulkCBuffer(m, l0, v482, v489)
	mBase = m.M
	v491 = m.ExcPending
	if v491 != 0 {
		goto L2
	} else {
		goto L118
	}
L115:
	;
	v483 = *(*int64)(unsafe.Add(mBase, uint32(v15)+24))
	F_addReplyBulkLongLong(m, l0, v483)
	mBase = m.M
	v485 = m.ExcPending
	if v485 != 0 {
		goto L2
	} else {
		goto L116
	}
L116:
	;
	v486 = *(*int64)(unsafe.Add(mBase, uint32(v15)+24))
	v487 = F_createStringObjectFromLongLong(m, v486)
	mBase = m.M
	v488 = m.ExcPending
	if v488 != 0 {
		goto L2
	} else {
		goto L117
	}
L117:
	;
	v495 = v487
	goto L113
L118:
	;
	v492 = F_createStringObject_1(m, v482, v489)
	mBase = m.M
	v493 = m.ExcPending
	if v493 != 0 {
		goto L2
	} else {
		goto L119
	}
L119:
	;
	v495 = v492
	goto L113
L120:
	;
	v550 = F_setTypeNext(m, v456, v15+int32(40), v15+int32(36), v15+int32(24))
	mBase = m.M
	v551 = m.ExcPending
	if v551 != 0 {
		goto L2
	} else {
		goto L127
	}
L121:
	;
	v503 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v504 = *(*int32)(unsafe.Add(mBase, uint32(v503)+28))
	v506 = *(*int32)(unsafe.Add(mBase, uint32(l0)+292))
	F_alsoPropagate(m, v504, v116, v113, int32(3), v506)
	mBase = m.M
	v508 = m.ExcPending
	if v508 != 0 {
		goto L2
	} else {
		goto L122
	}
L122:
	;
	v509 = int32(2)
	v515 = v509
	goto L123
L123:
	;
	v526 = *(*int32)(unsafe.Add(mBase, uint32(v116+v515<<(uint(int32(2))%32))))
	F_decrRefCount(m, v526)
	mBase = m.M
	v528 = m.ExcPending
	if v528 != 0 {
		goto L2
	} else {
		goto L125
	}
L124:
	;
	v540 = v509
	goto L120
L125:
	;
	v530 = v515 + int32(1)
	if v530 != v113 {
		v515 = v530
		goto L123
	} else {
		goto L126
	}
L126:
	;
	goto L124
L127:
	;
	if v550 != int32(-1) {
		v478 = v540
		goto L111
	} else {
		goto L128
	}
L128:
	;
	goto L112
L129:
	;
	v568 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v569 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v570 = *(*int32)(unsafe.Add(mBase, uint32(v569)+4))
	F_dbReplaceValue(m, v568, v570, v15+int32(20))
	mBase = m.M
	v574 = m.ExcPending
	if v574 != 0 {
		goto L2
	} else {
		goto L130
	}
L130:
	;
	v583 = v562
	goto L32
L131:
	;
	F_valkey_free(m, v116)
	mBase = m.M
	v632 = m.ExcPending
	if v632 != 0 {
		goto L2
	} else {
		goto L139
	}
L132:
	;
	v590 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v591 = *(*int32)(unsafe.Add(mBase, uint32(v590)+28))
	v593 = *(*int32)(unsafe.Add(mBase, uint32(l0)+292))
	F_alsoPropagate(m, v591, v116, v583, int32(3), v593)
	mBase = m.M
	v595 = m.ExcPending
	if v595 != 0 {
		goto L2
	} else {
		goto L133
	}
L133:
	;
	if base.Ui32(v583) < base.Ui32(int32(3)) {
		goto L131
	} else {
		goto L134
	}
L134:
	;
	v602 = v587
	goto L135
L135:
	;
	v613 = *(*int32)(unsafe.Add(mBase, uint32(v116+v602<<(uint(int32(2))%32))))
	F_decrRefCount(m, v613)
	mBase = m.M
	v615 = m.ExcPending
	if v615 != 0 {
		goto L2
	} else {
		goto L137
	}
L136:
	;
	goto L131
L137:
	;
	v617 = v602 + int32(1)
	if v617 != v583 {
		v602 = v617
		goto L135
	} else {
		goto L138
	}
L138:
	;
	goto L136
L139:
	;
	v633 = *(*int32)(unsafe.Add(mBase, uint32(l0)+200))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+200)) = v633 | int32(2097152)
	goto L140
L140:
	;
	v637 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v638 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v639 = *(*int32)(unsafe.Add(mBase, uint32(v638)+4))
	F_signalModifiedKey(m, l0, v637, v639)
	mBase = m.M
	v641 = m.ExcPending
	if v641 != 0 {
		goto L2
	} else {
		goto L141
	}
L141:
	;
	goto L1
}
func F_spscCommit(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
	if v3 == v4 {
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(l0)+64)) = v3
	}
	return
}
func F_spscDequeueBatch(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
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
	var v39 int32
	_ = v39
	var v45 int32
	_ = v45
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v85 int32
	_ = v85
	var v88 int32
	_ = v88
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v102 int32
	_ = v102
	var v108 int32
	_ = v108
	var v122 int32
	_ = v122
	var v126 int32
	_ = v126
	var v130 int32
	_ = v130
	var v138 int32
	_ = v138
	var v140 int32
	_ = v140
	var v143 int32
	_ = v143
	var v164 int32
	_ = v164
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v13 != v14 {
		v20 = v14
		v22 = v20 - v13
		if base.Ui32(l2) < base.Ui32(v22) {
			v24 = l2
		} else {
			v24 = v22
		}
		if v24 == int32(0) {
		} else {
			v28 = v24 & int32(3)
			v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)+132))
			v31 = v29 + int32(-1)
			v32 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
			v33 = int32(0)
			if base.Ui32(v24) < base.Ui32(int32(4)) {
				v108 = v33
			} else {
				v39 = int32(0)
				v45 = v39
				v51 = v39
				for {
					v53 = int32(2)
					v61 = *(*int32)(unsafe.Add(mBase, uint32(v32+v31&(v45+v13)<<(uint(v53)%32))))
					*(*int32)(unsafe.Add(mBase, uint32(l1+v45<<(uint(v53)%32)))) = v61
					v64 = v45 | int32(1)
					v73 = *(*int32)(unsafe.Add(mBase, uint32(v32+v31&(v64+v13)<<(uint(v53)%32))))
					*(*int32)(unsafe.Add(mBase, uint32(l1+v64<<(uint(v53)%32)))) = v73
					v76 = v45 | v53
					v85 = *(*int32)(unsafe.Add(mBase, uint32(v32+v31&(v76+v13)<<(uint(v53)%32))))
					*(*int32)(unsafe.Add(mBase, uint32(l1+v76<<(uint(v53)%32)))) = v85
					v88 = v45 | int32(3)
					v97 = *(*int32)(unsafe.Add(mBase, uint32(v32+v31&(v88+v13)<<(uint(v53)%32))))
					*(*int32)(unsafe.Add(mBase, uint32(l1+v88<<(uint(v53)%32)))) = v97
					v99 = int32(4)
					v100 = v45 + v99
					v102 = v51 + v99
					if v102 != v24&int32(-4) {
						v45 = v100
						v51 = v102
						continue
					} else {
						break
					}
					break
				}
				v108 = v100
			}
			if v28 == int32(0) {
			} else {
				v122 = v108
				v126 = v33
				for {
					v130 = int32(2)
					v138 = *(*int32)(unsafe.Add(mBase, uint32(v32+v31&(v122+v13)<<(uint(v130)%32))))
					*(*int32)(unsafe.Add(mBase, uint32(l1+v122<<(uint(v130)%32)))) = v138
					v140 = int32(1)
					v143 = v126 + v140
					if v143 != v28 {
						v122 = v122 + v140
						v126 = v143
						continue
					} else {
						break
					}
					break
				}
			}
		}
		*(*int32)(unsafe.Add(mBase, uint32(l0))) = v24 + v13
		v164 = v24
	} else {
		v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
		*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v16
		if v13 == v16 {
			v164 = int32(0)
		} else {
			v20 = v16
			v22 = v20 - v13
			if base.Ui32(l2) < base.Ui32(v22) {
				v24 = l2
			} else {
				v24 = v22
			}
			if v24 == int32(0) {
			} else {
				v28 = v24 & int32(3)
				v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)+132))
				v31 = v29 + int32(-1)
				v32 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
				v33 = int32(0)
				if base.Ui32(v24) < base.Ui32(int32(4)) {
					v108 = v33
				} else {
					v39 = int32(0)
					v45 = v39
					v51 = v39
					for {
						v53 = int32(2)
						v61 = *(*int32)(unsafe.Add(mBase, uint32(v32+v31&(v45+v13)<<(uint(v53)%32))))
						*(*int32)(unsafe.Add(mBase, uint32(l1+v45<<(uint(v53)%32)))) = v61
						v64 = v45 | int32(1)
						v73 = *(*int32)(unsafe.Add(mBase, uint32(v32+v31&(v64+v13)<<(uint(v53)%32))))
						*(*int32)(unsafe.Add(mBase, uint32(l1+v64<<(uint(v53)%32)))) = v73
						v76 = v45 | v53
						v85 = *(*int32)(unsafe.Add(mBase, uint32(v32+v31&(v76+v13)<<(uint(v53)%32))))
						*(*int32)(unsafe.Add(mBase, uint32(l1+v76<<(uint(v53)%32)))) = v85
						v88 = v45 | int32(3)
						v97 = *(*int32)(unsafe.Add(mBase, uint32(v32+v31&(v88+v13)<<(uint(v53)%32))))
						*(*int32)(unsafe.Add(mBase, uint32(l1+v88<<(uint(v53)%32)))) = v97
						v99 = int32(4)
						v100 = v45 + v99
						v102 = v51 + v99
						if v102 != v24&int32(-4) {
							v45 = v100
							v51 = v102
							continue
						} else {
							break
						}
						break
					}
					v108 = v100
				}
				if v28 == int32(0) {
				} else {
					v122 = v108
					v126 = v33
					for {
						v130 = int32(2)
						v138 = *(*int32)(unsafe.Add(mBase, uint32(v32+v31&(v122+v13)<<(uint(v130)%32))))
						*(*int32)(unsafe.Add(mBase, uint32(l1+v122<<(uint(v130)%32)))) = v138
						v140 = int32(1)
						v143 = v126 + v140
						if v143 != v28 {
							v122 = v122 + v140
							v126 = v143
							continue
						} else {
							break
						}
						break
					}
				}
			}
			*(*int32)(unsafe.Add(mBase, uint32(l0))) = v24 + v13
			v164 = v24
		}
	}
	return v164
}
func F_spscEnqueue(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v16 int32
	_ = v16
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+132))
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
	*(*int32)(unsafe.Add(mBase, uint32(v5+(v6+int32(-1))&v9<<(uint(int32(2))%32)))) = l1
	v16 = v9 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+68)) = v16
	if l2 == int32(0) {
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(l0)+64)) = v16
	}
	return
}
func F_spscFree(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v7 int32
	_ = v7
	var v12 int64
	_ = v12
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
	if v3 == int32(0) {
		*(*int32)(unsafe.Add(mBase, uint32(l0)+72)) = int32(0)
		v12 = int64(0)
		*(*int64)(unsafe.Add(mBase, uint32(l0))) = v12
		*(*int64)(unsafe.Add(mBase, uint32(l0)+64)) = v12
		return
	} else {
		F_valkey_free(m, v3)
		mBase = m.M
		v7 = m.ExcPending
		if v7 != 0 {
			return
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(l0)+128)) = int32(0)
			*(*int32)(unsafe.Add(mBase, uint32(l0)+72)) = int32(0)
			v12 = int64(0)
			*(*int64)(unsafe.Add(mBase, uint32(l0))) = v12
			*(*int64)(unsafe.Add(mBase, uint32(l0)+64)) = v12
			return
		}
	}
}
func F_spublishCommand(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
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
	var v16 int32
	_ = v16
	var v21 int64
	_ = v21
	var v26 int64
	_ = v26
	var v29 int64
	_ = v29
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	v2 = int32(0)
	v6 = m.G0
	v8 = v6 - int32(32)
	m.G0 = v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(v10)+8))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(v10)+4))
	v16 = *(*int32)(unsafe.Add(mBase, _c_F_spublishCommand[0]))
	*(*int32)(unsafe.Add(mBase, uint32(v8+int32(24)))) = v16
	v21 = *(*int64)(unsafe.Add(mBase, _c_F_spublishCommand[1]))
	*(*int64)(unsafe.Add(mBase, uint32(v8+int32(16)))) = v21
	v26 = *(*int64)(unsafe.Add(mBase, _c_F_spublishCommand[2]))
	*(*int64)(unsafe.Add(mBase, uint32(v8+int32(8)))) = v26
	v29 = *(*int64)(unsafe.Add(mBase, _c_F_spublishCommand[3]))
	*(*int64)(unsafe.Add(mBase, uint32(v8))) = v29
	v31 = F_pubsubPublishMessageInternal(m, v12, v11, v8)
	mBase = m.M
	v32 = m.ExcPending
	if v32 != 0 {
		return
	} else {
		v34 = *(*int32)(unsafe.Add(mBase, _c_F_spublishCommand[4]))
		if v34 == int32(0) {
			F_forceCommandPropagation(m, l0, int32(2))
			mBase = m.M
			v44 = m.ExcPending
			if v44 != 0 {
				return
			} else {
				F_addReplyLongLong(m, l0, base.I64_extend_i32_s(v31))
				mBase = m.M
				v47 = m.ExcPending
				if v47 != 0 {
					return
				} else {
					m.G0 = v8 + int32(32)
					return
				}
			}
		} else {
			F_clusterPropagatePublish(m, v12, v11, int32(1))
			mBase = m.M
			v39 = m.ExcPending
			if v39 != 0 {
				return
			} else {
				v41 = *(*int32)(unsafe.Add(mBase, _c_F_spublishCommand[4]))
				if v41 != 0 {
					F_addReplyLongLong(m, l0, base.I64_extend_i32_s(v31))
					mBase = m.M
					v47 = m.ExcPending
					if v47 != 0 {
						return
					} else {
						m.G0 = v8 + int32(32)
						return
					}
				} else {
					F_forceCommandPropagation(m, l0, int32(2))
					mBase = m.M
					v44 = m.ExcPending
					if v44 != 0 {
						return
					} else {
						F_addReplyLongLong(m, l0, base.I64_extend_i32_s(v31))
						mBase = m.M
						v47 = m.ExcPending
						if v47 != 0 {
							return
						} else {
							m.G0 = v8 + int32(32)
							return
						}
					}
				}
			}
		}
	}
}
func F_sqrt(m *base.Module, l0 float64) float64 {
	return base.F64_sqrt(l0)
}
func F_srandmemberCommand(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v12 int32
	_ = v12
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int64
	_ = v41
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v8 != int32(3) {
		if v8 < int32(4) {
			v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
			v20 = *(*int32)(unsafe.Add(mBase, uint32(v19)+4))
			v22 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+224)))
			v26 = *(*int32)(unsafe.Add(mBase, uint32(v22<<(uint(int32(2))%32))+uint32(_c_F_srandmemberCommand[0])))
			v27 = F_lookupKeyReadOrReply(m, l0, v20, v26)
			mBase = m.M
			v28 = m.ExcPending
			if v28 != 0 {
				return
			} else {
				if v27 == int32(0) {
					m.G0 = v6 + int32(16)
					return
				} else {
					v32 = F_checkType(m, l0, v27, int32(2))
					mBase = m.M
					v33 = m.ExcPending
					if v33 != 0 {
						return
					} else {
						if v32 != 0 {
							m.G0 = v6 + int32(16)
							return
						} else {
							v38 = F_setTypeRandomElement(m, v27, v6+int32(12), v6+int32(8), v6)
							mBase = m.M
							v39 = m.ExcPending
							if v39 != 0 {
								return
							} else {
								v40 = *(*int32)(unsafe.Add(mBase, uint32(v6)+12))
								if v40 != 0 {
									v44 = *(*int32)(unsafe.Add(mBase, uint32(v6)+8))
									F_addReplyBulkCBuffer(m, l0, v40, v44)
									mBase = m.M
									v46 = m.ExcPending
									if v46 != 0 {
										return
									} else {
										m.G0 = v6 + int32(16)
										return
									}
								} else {
									v41 = *(*int64)(unsafe.Add(mBase, uint32(v6)))
									F_addReplyBulkLongLong(m, l0, v41)
									mBase = m.M
									v43 = m.ExcPending
									if v43 != 0 {
										return
									} else {
										m.G0 = v6 + int32(16)
										return
									}
								}
							}
						}
					}
				}
			}
		} else {
			v16 = *(*int32)(unsafe.Add(mBase, _c_F_srandmemberCommand[1]))
			F_addReplyErrorObject(m, l0, v16)
			mBase = m.M
			v18 = m.ExcPending
			if v18 != 0 {
				return
			} else {
				m.G0 = v6 + int32(16)
				return
			}
		}
	} else {
		F_srandmemberWithCountCommand(m, l0)
		mBase = m.M
		v12 = m.ExcPending
		if v12 != 0 {
			return
		} else {
			m.G0 = v6 + int32(16)
			return
		}
	}
}
func F_srandom(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
	var v31 int32
	_ = v31
	var v34 int64
	_ = v34
	var v38 int64
	_ = v38
	var v40 int64
	_ = v40
	var v43 int32
	_ = v43
	var v49 int32
	_ = v49
	v7 = *(*int32)(unsafe.Add(mBase, _c_F_srandom[0]))
	if v7 != 0 {
		v12 = int32(3)
		if v7 == int32(7) {
			v17 = v12
		} else {
			v17 = int32(1)
		}
		if v7 == int32(31) {
			v20 = v12
		} else {
			v20 = v17
		}
		*(*int32)(unsafe.Add(mBase, _c_F_srandom[1])) = v20
		v22 = int32(0)
		*(*int32)(unsafe.Add(mBase, _c_F_srandom[2])) = v22
		v26 = *(*int32)(unsafe.Add(mBase, _c_F_srandom[3]))
		if v7 < int32(1) {
		} else {
			v31 = int32(0)
			v34 = base.I64_extend_i32_u(l0)
			for {
				v38 = F_lcg64(m, v34)
				mBase = m.M
				v40 = int64(base.Ui64(v38) >> (uint(int64(32)) % 64))
				*(*uint32)(unsafe.Add(mBase, uint32(v26+v31<<(uint(int32(2))%32)))) = uint32(v40)
				v43 = v31 + int32(1)
				if v43 != v7 {
					v31 = v43
					v34 = v38
					continue
				} else {
					break
				}
				break
			}
		}
		v49 = *(*int32)(unsafe.Add(mBase, uint32(v26)))
		*(*int32)(unsafe.Add(mBase, uint32(v26))) = v49 | int32(1)
	} else {
		v9 = *(*int32)(unsafe.Add(mBase, _c_F_srandom[3]))
		*(*int32)(unsafe.Add(mBase, uint32(v9))) = l0
	}
	return
}
func F_sremCommand(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v48 int32
	_ = v48
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
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
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v98 int32
	_ = v98
	var v104 int32
	_ = v104
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v126 int32
	_ = v126
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v134 int32
	_ = v134
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v147 int64
	_ = v147
	var v148 int64
	_ = v148
	var v151 int64
	_ = v151
	var v153 int32
	_ = v153
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(v10)+4))
	v13 = *(*int32)(unsafe.Add(mBase, _c_F_sremCommand[0]))
	v14 = F_lookupKeyWriteOrReply(m, l0, v11, v13)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	return
L2:
	;
	return
L3:
	;
	if v14 == int32(0) {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v19 = F_checkType(m, l0, v14, int32(2))
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		goto L2
	} else {
		goto L5
	}
L5:
	;
	if v19 != 0 {
		goto L1
	} else {
		goto L6
	}
L6:
	;
	v21 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
	if v21&int32(240) != int32(32) {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	v31 = int32(0)
	v33 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v33 < int32(3) {
		v98 = v31
		goto L11
	} else {
		goto L12
	}
L8:
	;
	v26 = F_objectGetVal(m, v14)
	mBase = m.M
	v27 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v26)+28)))
	v29 = v27 + int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v26)+28)) = uint16(v29)
	goto L9
L9:
	;
	goto L7
L10:
	;
	if v115 != 0 {
		goto L31
	} else {
		goto L32
	}
L11:
	;
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
	if v104&int32(240) != int32(32) {
		v114 = v31
		v115 = v98
		goto L10
	} else {
		goto L28
	}
L12:
	;
	v41 = int32(0)
	v42 = int32(2)
	goto L13
L13:
	;
	v48 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v48+v42<<(uint(int32(2))%32))))
	v53 = F_objectGetVal(m, v52)
	mBase = m.M
	v56 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v53+int32(-1)))))
	switch v56 & int32(7) {
	case 0:
		goto L20
	case 1:
		goto L19
	case 2:
		goto L18
	case 3:
		goto L17
	case 4:
		goto L16
	default:
		v73 = int32(0)
		goto L15
	}
L14:
	;
	v98 = v90
	goto L11
L15:
	;
	v76 = F_setTypeRemoveAux(m, v14, v53, v73, int64(0), int32(1))
	mBase = m.M
	v77 = m.ExcPending
	if v77 != 0 {
		goto L2
	} else {
		goto L22
	}
L16:
	;
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v53+int32(-17))))
	v73 = v72
	goto L15
L17:
	;
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v53+int32(-9))))
	v73 = v69
	goto L15
L18:
	;
	v66 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v53+int32(-5)))))
	v73 = v66
	goto L15
L19:
	;
	v63 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v53+int32(-3)))))
	v73 = v63
	goto L15
L20:
	;
	v73 = int32(base.Ui32(v56) >> (uint(int32(3)) % 32))
	goto L15
L21:
	;
	v92 = v42 + int32(1)
	v93 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v92 < v93 {
		v41 = v90
		v42 = v92
		goto L13
	} else {
		goto L27
	}
L22:
	;
	if v76 == int32(0) {
		v90 = v41
		goto L21
	} else {
		goto L23
	}
L23:
	;
	v81 = v41 + int32(1)
	v82 = F_setTypeSize(m, v14)
	mBase = m.M
	v83 = m.ExcPending
	if v83 != 0 {
		goto L2
	} else {
		goto L24
	}
L24:
	;
	if v82 != 0 {
		v90 = v81
		goto L21
	} else {
		goto L25
	}
L25:
	;
	v84 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v85 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v85)+4))
	v87 = F_dbDelete(m, v84, v86)
	mBase = m.M
	v88 = m.ExcPending
	if v88 != 0 {
		goto L2
	} else {
		goto L26
	}
L26:
	;
	v114 = int32(1)
	v115 = v81
	goto L10
L27:
	;
	goto L14
L28:
	;
	v109 = F_objectGetVal(m, v14)
	mBase = m.M
	F_hashtableResumeAutoShrink(m, v109)
	mBase = m.M
	v111 = m.ExcPending
	if v111 != 0 {
		goto L2
	} else {
		goto L29
	}
L29:
	;
	v114 = v31
	v115 = v98
	goto L10
L30:
	;
	F_addReplyLongLong(m, l0, v151)
	mBase = m.M
	v153 = m.ExcPending
	if v153 != 0 {
		goto L2
	} else {
		goto L38
	}
L31:
	;
	v122 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v123 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v124 = *(*int32)(unsafe.Add(mBase, uint32(v123)+4))
	F_signalModifiedKey(m, l0, v122, v124)
	mBase = m.M
	v126 = m.ExcPending
	if v126 != 0 {
		goto L2
	} else {
		goto L33
	}
L32:
	;
	v151 = int64(0)
	goto L30
L33:
	;
	v129 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v130 = *(*int32)(unsafe.Add(mBase, uint32(v129)+4))
	v131 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v132 = *(*int32)(unsafe.Add(mBase, uint32(v131)+28))
	F_notifyKeyspaceEvent(m, int32(32), int32(_a_F_sremCommand_0), v130, v132)
	mBase = m.M
	v134 = m.ExcPending
	if v134 != 0 {
		goto L2
	} else {
		goto L34
	}
L34:
	;
	if v114 == int32(0) {
		goto L35
	} else {
		goto L36
	}
L35:
	;
	v145 = int32(_a_F_sremCommand_1)
	v147 = *(*int64)(unsafe.Add(mBase, _c_F_sremCommand[1]))
	v148 = base.I64_extend_i32_s(v115)
	*(*int64)(unsafe.Add(mBase, _c_F_sremCommand[1])) = v147 + v148
	v151 = v148
	goto L30
L36:
	;
	v139 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v140 = *(*int32)(unsafe.Add(mBase, uint32(v139)+4))
	v141 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v142 = *(*int32)(unsafe.Add(mBase, uint32(v141)+28))
	F_notifyKeyspaceEvent(m, int32(4), int32(_a_F_sremCommand_2), v140, v142)
	mBase = m.M
	v144 = m.ExcPending
	if v144 != 0 {
		goto L2
	} else {
		goto L37
	}
L37:
	;
	goto L35
L38:
	;
	goto L1
}
func F_stopLoading(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v18 int32
	_ = v18
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	v2 = int32(_a_F_stopLoading_0)
	v3 = int32(0)
	*(*int32)(unsafe.Add(mBase, _c_F_stopLoading[0])) = v3
	*(*int32)(unsafe.Add(mBase, _c_F_stopLoading[1])) = v3
	v11 = *(*int32)(unsafe.Add(mBase, _c_F_stopLoading[2]))
	v13 = v11 + int32(-1)
	*(*int32)(unsafe.Add(mBase, _c_F_stopLoading[2])) = v13
	if v13 != 0 {
	} else {
		*(*int64)(unsafe.Add(mBase, _c_F_stopLoading[3])) = int64(0)
	}
	v18 = int32(0)
	*(*int32)(unsafe.Add(mBase, _c_F_stopLoading[4])) = v18
	if l0 != 0 {
		v24 = int32(3)
	} else {
		v24 = int32(4)
	}
	F_moduleFireServerEvent(m, int64(3), v24, int32(0))
	mBase = m.M
	v27 = m.ExcPending
	if v27 != 0 {
		return
	} else {
		return
	}
}
func F_strbuf_append_string(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	v7 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	if v7 == int32(0) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v18 = v7
	v19 = v11
	v20 = v10 + (v11 ^ int32(-1))
	v21 = int32(0)
	goto L3
L3:
	;
	if v20 != 0 {
		v33 = v18
		v34 = v19
		v35 = v20
		goto L5
	} else {
		goto L6
	}
L4:
	;
	goto L1
L5:
	;
	v36 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*uint8)(unsafe.Add(mBase, uint32(v36+v34))) = uint8(v33)
	v39 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v40 = int32(1)
	v41 = v39 + v40
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v41
	v46 = v21 + v40
	v48 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1+v46))))
	if v48 != 0 {
		v18 = v48
		v19 = v41
		v20 = v35 + int32(-1)
		v21 = v46
		goto L3
	} else {
		goto L9
	}
L6:
	;
	F_strbuf_resize(m, l0, v19+int32(1))
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	return
L8:
	;
	v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v32 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1+v21))))
	v33 = v32
	v34 = v27
	v35 = v26 + (v27 ^ int32(-1))
	goto L5
L9:
	;
	goto L4
}
func F_strbuf_free(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v12 int64
	_ = v12
	var v13 int32
	_ = v13
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	var v88 int32
	_ = v88
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v99 int32
	_ = v99
	var __phi99 int32
	_ = __phi99
	var v100 int32
	_ = v100
	var __phi100 int32
	_ = __phi100
	var v106 int32
	_ = v106
	var v109 int32
	_ = v109
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v131 int32
	_ = v131
	var v140 int32
	_ = v140
	var v142 int32
	_ = v142
	var v145 int32
	_ = v145
	var v148 int32
	_ = v148
	var v150 int32
	_ = v150
	var v155 int32
	_ = v155
	var v162 int32
	_ = v162
	var v167 int32
	_ = v167
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v182 int32
	_ = v182
	var v190 int32
	_ = v190
	var v192 int32
	_ = v192
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v203 int32
	_ = v203
	var v205 int32
	_ = v205
	var v212 int32
	_ = v212
	var v214 int32
	_ = v214
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v232 int32
	_ = v232
	var v234 int32
	_ = v234
	var v236 int32
	_ = v236
	var v245 int32
	_ = v245
	var v247 int32
	_ = v247
	var v250 int32
	_ = v250
	var v255 int32
	_ = v255
	var v260 int32
	_ = v260
	var v261 int32
	_ = v261
	var v266 int32
	_ = v266
	var __phi266 int32
	_ = __phi266
	var v267 int32
	_ = v267
	var __phi267 int32
	_ = __phi267
	var v273 int32
	_ = v273
	var v276 int32
	_ = v276
	var v290 int32
	_ = v290
	var v299 int32
	_ = v299
	var v301 int32
	_ = v301
	var v304 int32
	_ = v304
	var v307 int32
	_ = v307
	var v309 int32
	_ = v309
	var v314 int32
	_ = v314
	var v321 int32
	_ = v321
	var v326 int32
	_ = v326
	var v348 int32
	_ = v348
	var v360 int32
	_ = v360
	var v362 int32
	_ = v362
	var v364 int32
	_ = v364
	var v368 int32
	_ = v368
	var v373 int32
	_ = v373
	var v374 int32
	_ = v374
	var v385 int32
	_ = v385
	var v388 int32
	_ = v388
	var v395 int32
	_ = v395
	var v400 int32
	_ = v400
	var v404 int32
	_ = v404
	var v406 int32
	_ = v406
	var v421 int32
	_ = v421
	var v423 int32
	_ = v423
	var v426 int32
	_ = v426
	var v429 int32
	_ = v429
	var v433 int32
	_ = v433
	var v443 int32
	_ = v443
	var v444 int32
	_ = v444
	var v450 int32
	_ = v450
	var v452 int32
	_ = v452
	var v455 int32
	_ = v455
	var v459 int32
	_ = v459
	var v465 int32
	_ = v465
	var v467 int32
	_ = v467
	var v468 int32
	_ = v468
	var v469 int32
	_ = v469
	var v470 int32
	_ = v470
	var v479 int32
	_ = v479
	var v481 int32
	_ = v481
	var v482 int32
	_ = v482
	var v483 int32
	_ = v483
	var v485 int32
	_ = v485
	var v507 int32
	_ = v507
	var v520 int32
	_ = v520
	var v521 int32
	_ = v521
	var v524 int32
	_ = v524
	var v526 int32
	_ = v526
	var v527 int32
	_ = v527
	var v534 int32
	_ = v534
	var v535 int32
	_ = v535
	var v537 int32
	_ = v537
	var v539 int32
	_ = v539
	var v541 int32
	_ = v541
	var v543 int32
	_ = v543
	var v546 int32
	_ = v546
	var v548 int32
	_ = v548
	var v550 int32
	_ = v550
	var v557 int32
	_ = v557
	var v559 int32
	_ = v559
	var v562 int32
	_ = v562
	var v567 int32
	_ = v567
	var v572 int32
	_ = v572
	var v573 int32
	_ = v573
	var v578 int32
	_ = v578
	var __phi578 int32
	_ = __phi578
	var v579 int32
	_ = v579
	var __phi579 int32
	_ = __phi579
	var v585 int32
	_ = v585
	var v588 int32
	_ = v588
	var v591 int32
	_ = v591
	var v592 int32
	_ = v592
	var v610 int32
	_ = v610
	var v619 int32
	_ = v619
	var v621 int32
	_ = v621
	var v624 int32
	_ = v624
	var v627 int32
	_ = v627
	var v629 int32
	_ = v629
	var v634 int32
	_ = v634
	var v641 int32
	_ = v641
	var v646 int32
	_ = v646
	var v651 int32
	_ = v651
	var v652 int32
	_ = v652
	var v661 int32
	_ = v661
	var v669 int32
	_ = v669
	var v671 int32
	_ = v671
	var v675 int32
	_ = v675
	var v676 int32
	_ = v676
	var v682 int32
	_ = v682
	var v684 int32
	_ = v684
	var v691 int32
	_ = v691
	var v693 int32
	_ = v693
	var v697 int32
	_ = v697
	var v698 int32
	_ = v698
	var v707 int32
	_ = v707
	var v708 int32
	_ = v708
	var v711 int32
	_ = v711
	var v713 int32
	_ = v713
	var v715 int32
	_ = v715
	var v724 int32
	_ = v724
	var v726 int32
	_ = v726
	var v729 int32
	_ = v729
	var v734 int32
	_ = v734
	var v739 int32
	_ = v739
	var v740 int32
	_ = v740
	var v745 int32
	_ = v745
	var __phi745 int32
	_ = __phi745
	var v746 int32
	_ = v746
	var __phi746 int32
	_ = __phi746
	var v752 int32
	_ = v752
	var v755 int32
	_ = v755
	var v769 int32
	_ = v769
	var v778 int32
	_ = v778
	var v780 int32
	_ = v780
	var v783 int32
	_ = v783
	var v786 int32
	_ = v786
	var v788 int32
	_ = v788
	var v793 int32
	_ = v793
	var v800 int32
	_ = v800
	var v805 int32
	_ = v805
	var v827 int32
	_ = v827
	var v839 int32
	_ = v839
	var v841 int32
	_ = v841
	var v843 int32
	_ = v843
	var v847 int32
	_ = v847
	var v852 int32
	_ = v852
	var v853 int32
	_ = v853
	var v864 int32
	_ = v864
	var v867 int32
	_ = v867
	var v874 int32
	_ = v874
	var v879 int32
	_ = v879
	var v883 int32
	_ = v883
	var v885 int32
	_ = v885
	var v900 int32
	_ = v900
	var v902 int32
	_ = v902
	var v905 int32
	_ = v905
	var v908 int32
	_ = v908
	var v912 int32
	_ = v912
	var v922 int32
	_ = v922
	var v923 int32
	_ = v923
	var v929 int32
	_ = v929
	var v931 int32
	_ = v931
	var v934 int32
	_ = v934
	var v938 int32
	_ = v938
	var v944 int32
	_ = v944
	var v946 int32
	_ = v946
	var v947 int32
	_ = v947
	var v948 int32
	_ = v948
	var v949 int32
	_ = v949
	var v958 int32
	_ = v958
	var v960 int32
	_ = v960
	var v961 int32
	_ = v961
	var v962 int32
	_ = v962
	var v964 int32
	_ = v964
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v9 == int32(0) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v28 == int32(0) {
		goto L5
	} else {
		goto L6
	}
L2:
	;
	v12 = *(*int64)(unsafe.Add(mBase, uint32(l0)+4))
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v7)+4)) = v13
	*(*int32)(unsafe.Add(mBase, uint32(v7))) = l0
	*(*int64)(unsafe.Add(mBase, uint32(v7)+8)) = base.I64_rotl(v12, int64(32))
	v19 = m.G3
	v20 = m.G397
	v21 = *(*int32)(unsafe.Add(mBase, uint32(v20)))
	v24 = F_fiprintf(m, v21, v19+int32(_a_F_strbuf_free_0), v7)
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	return
L4:
	;
	goto L1
L5:
	;
	v507 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v507 == int32(0) {
		goto L107
	} else {
		goto L108
	}
L6:
	;
	if v28 == int32(0) {
		goto L8
	} else {
		goto L9
	}
L7:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(0)
	goto L5
L8:
	;
	goto L7
L9:
	;
	v41 = int32(-8)
	v42 = v28 + v41
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v28+int32(-4))))
	v47 = v45 & v41
	v48 = v42 + v47
	if v45&int32(1) != 0 {
		v172 = v47
		v173 = v42
		goto L10
	} else {
		goto L11
	}
L10:
	;
	if base.Ui32(v48) <= base.Ui32(v173) {
		goto L8
	} else {
		goto L45
	}
L11:
	;
	if v45&int32(2) == int32(0) {
		goto L8
	} else {
		goto L12
	}
L12:
	;
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v42)))
	v56 = v42 - v55
	v58 = *(*int32)(unsafe.Add(mBase, _c_F_strbuf_free[0]))
	if base.Ui32(v56) < base.Ui32(v58) {
		goto L8
	} else {
		goto L13
	}
L13:
	;
	v60 = v55 + v47
	v62 = *(*int32)(unsafe.Add(mBase, _c_F_strbuf_free[1]))
	if v56 == v62 {
		goto L17
	} else {
		goto L18
	}
L14:
	;
	if v78 == int32(0) {
		v172 = v60
		v173 = v56
		goto L10
	} else {
		goto L33
	}
L15:
	;
	v131 = int32(0)
	goto L14
L16:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v67)+12)) = v64
	*(*int32)(unsafe.Add(mBase, uint32(v64)+8)) = v67
	v172 = v60
	v173 = v56
	goto L10
L17:
	;
	v112 = *(*int32)(unsafe.Add(mBase, uint32(v48)+4))
	v113 = int32(3)
	if v112&v113 != v113 {
		v172 = v60
		v173 = v56
		goto L10
	} else {
		goto L32
	}
L18:
	;
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v56)+12))
	if base.Ui32(int32(255)) < base.Ui32(v55) {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v56)+24))
	if v64 == v56 {
		goto L22
	} else {
		goto L23
	}
L20:
	;
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v56)+8))
	if v64 != v67 {
		goto L16
	} else {
		goto L21
	}
L21:
	;
	v69 = int32(0)
	v71 = *(*int32)(unsafe.Add(mBase, _c_F_strbuf_free[2]))
	*(*int32)(unsafe.Add(mBase, _c_F_strbuf_free[2])) = v71 & base.I32_rotl(int32(-2), int32(base.Ui32(v55)>>(uint(int32(3))%32)))
	v172 = v60
	v173 = v56
	goto L10
L22:
	;
	v83 = *(*int32)(unsafe.Add(mBase, uint32(v56)+20))
	if v83 == int32(0) {
		goto L25
	} else {
		goto L26
	}
L23:
	;
	v80 = *(*int32)(unsafe.Add(mBase, uint32(v56)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v80)+12)) = v64
	*(*int32)(unsafe.Add(mBase, uint32(v64)+8)) = v80
	v131 = v64
	goto L14
L24:
	;
	__phi99 = v93
	__phi100 = v94
	v99 = __phi99
	v100 = __phi100
	goto L28
L25:
	;
	v88 = *(*int32)(unsafe.Add(mBase, uint32(v56)+16))
	if v88 == int32(0) {
		goto L15
	} else {
		goto L27
	}
L26:
	;
	v93 = v83
	v94 = v56 + int32(20)
	goto L24
L27:
	;
	v93 = v88
	v94 = v56 + int32(16)
	goto L24
L28:
	;
	v106 = *(*int32)(unsafe.Add(mBase, uint32(v99)+20))
	if v106 != 0 {
		__phi99 = v106
		__phi100 = v99 + int32(20)
		v99 = __phi99
		v100 = __phi100
		goto L28
	} else {
		goto L30
	}
L29:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v100))) = int32(0)
	v131 = v99
	goto L14
L30:
	;
	v109 = *(*int32)(unsafe.Add(mBase, uint32(v99)+16))
	if v109 != 0 {
		__phi99 = v109
		__phi100 = v99 + int32(16)
		v99 = __phi99
		v100 = __phi100
		goto L28
	} else {
		goto L31
	}
L31:
	;
	goto L29
L32:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_strbuf_free[3])) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v48)+4)) = v112 & int32(-2)
	*(*int32)(unsafe.Add(mBase, uint32(v56)+4)) = v60 | int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v48))) = v60
	goto L7
L33:
	;
	v140 = *(*int32)(unsafe.Add(mBase, uint32(v56)+28))
	v142 = v140 << (uint(int32(2)) % 32)
	v145 = *(*int32)(unsafe.Add(mBase, uint32(v142)+uint32(_c_F_strbuf_free[4])))
	if v56 != v145 {
		goto L35
	} else {
		goto L36
	}
L34:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v131)+24)) = v78
	v162 = *(*int32)(unsafe.Add(mBase, uint32(v56)+16))
	if v162 == int32(0) {
		goto L42
	} else {
		goto L43
	}
L35:
	;
	v155 = *(*int32)(unsafe.Add(mBase, uint32(v78)+16))
	if v155 != v56 {
		goto L39
	} else {
		goto L40
	}
L36:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v142)+uint32(_c_F_strbuf_free[4]))) = v131
	if v131 != 0 {
		goto L34
	} else {
		goto L37
	}
L37:
	;
	v148 = int32(0)
	v150 = *(*int32)(unsafe.Add(mBase, _c_F_strbuf_free[5]))
	*(*int32)(unsafe.Add(mBase, _c_F_strbuf_free[5])) = v150 & base.I32_rotl(int32(-2), v140)
	v172 = v60
	v173 = v56
	goto L10
L38:
	;
	if v131 == int32(0) {
		v172 = v60
		v173 = v56
		goto L10
	} else {
		goto L41
	}
L39:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v78)+20)) = v131
	goto L38
L40:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v78)+16)) = v131
	goto L38
L41:
	;
	goto L34
L42:
	;
	v167 = *(*int32)(unsafe.Add(mBase, uint32(v56)+20))
	if v167 == int32(0) {
		v172 = v60
		v173 = v56
		goto L10
	} else {
		goto L44
	}
L43:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v131)+16)) = v162
	*(*int32)(unsafe.Add(mBase, uint32(v162)+24)) = v131
	goto L42
L44:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v131)+20)) = v167
	*(*int32)(unsafe.Add(mBase, uint32(v167)+24)) = v131
	v172 = v60
	v173 = v56
	goto L10
L45:
	;
	v182 = *(*int32)(unsafe.Add(mBase, uint32(v48)+4))
	if v182&int32(1) == int32(0) {
		goto L8
	} else {
		goto L46
	}
L46:
	;
	if v182&int32(2) != 0 {
		goto L51
	} else {
		goto L52
	}
L47:
	;
	if base.Ui32(int32(255)) < base.Ui32(v348) {
		goto L85
	} else {
		goto L86
	}
L48:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v173)+4)) = v228 | int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v173+v228))) = v228
	if v173 != v212 {
		v348 = v228
		goto L47
	} else {
		goto L84
	}
L49:
	;
	if v245 == int32(0) {
		goto L48
	} else {
		goto L72
	}
L50:
	;
	v290 = int32(0)
	goto L49
L51:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v48)+4)) = v182 & int32(-2)
	*(*int32)(unsafe.Add(mBase, uint32(v173)+4)) = v172 | int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v173+v172))) = v172
	v348 = v172
	goto L47
L52:
	;
	v190 = *(*int32)(unsafe.Add(mBase, _c_F_strbuf_free[6]))
	if v48 != v190 {
		goto L53
	} else {
		goto L54
	}
L53:
	;
	v212 = *(*int32)(unsafe.Add(mBase, _c_F_strbuf_free[1]))
	if v48 != v212 {
		goto L56
	} else {
		goto L57
	}
L54:
	;
	v192 = int32(0)
	*(*int32)(unsafe.Add(mBase, _c_F_strbuf_free[6])) = v173
	v196 = *(*int32)(unsafe.Add(mBase, _c_F_strbuf_free[7]))
	v197 = v196 + v172
	*(*int32)(unsafe.Add(mBase, _c_F_strbuf_free[7])) = v197
	*(*int32)(unsafe.Add(mBase, uint32(v173)+4)) = v197 | int32(1)
	v203 = *(*int32)(unsafe.Add(mBase, _c_F_strbuf_free[1]))
	if v173 != v203 {
		goto L8
	} else {
		goto L55
	}
L55:
	;
	v205 = int32(0)
	*(*int32)(unsafe.Add(mBase, _c_F_strbuf_free[3])) = v205
	*(*int32)(unsafe.Add(mBase, _c_F_strbuf_free[1])) = v205
	goto L7
L56:
	;
	v228 = v182&int32(-8) + v172
	v229 = *(*int32)(unsafe.Add(mBase, uint32(v48)+12))
	if base.Ui32(int32(255)) < base.Ui32(v182) {
		goto L58
	} else {
		goto L59
	}
L57:
	;
	v214 = int32(0)
	*(*int32)(unsafe.Add(mBase, _c_F_strbuf_free[1])) = v173
	v218 = *(*int32)(unsafe.Add(mBase, _c_F_strbuf_free[3]))
	v219 = v218 + v172
	*(*int32)(unsafe.Add(mBase, _c_F_strbuf_free[3])) = v219
	*(*int32)(unsafe.Add(mBase, uint32(v173)+4)) = v219 | int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v173+v219))) = v219
	goto L7
L58:
	;
	v245 = *(*int32)(unsafe.Add(mBase, uint32(v48)+24))
	if v229 == v48 {
		goto L62
	} else {
		goto L63
	}
L59:
	;
	v232 = *(*int32)(unsafe.Add(mBase, uint32(v48)+8))
	if v229 != v232 {
		goto L60
	} else {
		goto L61
	}
L60:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v232)+12)) = v229
	*(*int32)(unsafe.Add(mBase, uint32(v229)+8)) = v232
	goto L48
L61:
	;
	v234 = int32(0)
	v236 = *(*int32)(unsafe.Add(mBase, _c_F_strbuf_free[2]))
	*(*int32)(unsafe.Add(mBase, _c_F_strbuf_free[2])) = v236 & base.I32_rotl(int32(-2), int32(base.Ui32(v182)>>(uint(int32(3))%32)))
	goto L48
L62:
	;
	v250 = *(*int32)(unsafe.Add(mBase, uint32(v48)+20))
	if v250 == int32(0) {
		goto L65
	} else {
		goto L66
	}
L63:
	;
	v247 = *(*int32)(unsafe.Add(mBase, uint32(v48)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v247)+12)) = v229
	*(*int32)(unsafe.Add(mBase, uint32(v229)+8)) = v247
	v290 = v229
	goto L49
L64:
	;
	__phi266 = v260
	__phi267 = v261
	v266 = __phi266
	v267 = __phi267
	goto L68
L65:
	;
	v255 = *(*int32)(unsafe.Add(mBase, uint32(v48)+16))
	if v255 == int32(0) {
		goto L50
	} else {
		goto L67
	}
L66:
	;
	v260 = v250
	v261 = v48 + int32(20)
	goto L64
L67:
	;
	v260 = v255
	v261 = v48 + int32(16)
	goto L64
L68:
	;
	v273 = *(*int32)(unsafe.Add(mBase, uint32(v266)+20))
	if v273 != 0 {
		__phi266 = v273
		__phi267 = v266 + int32(20)
		v266 = __phi266
		v267 = __phi267
		goto L68
	} else {
		goto L70
	}
L69:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v267))) = int32(0)
	v290 = v266
	goto L49
L70:
	;
	v276 = *(*int32)(unsafe.Add(mBase, uint32(v266)+16))
	if v276 != 0 {
		__phi266 = v276
		__phi267 = v266 + int32(16)
		v266 = __phi266
		v267 = __phi267
		goto L68
	} else {
		goto L71
	}
L71:
	;
	goto L69
L72:
	;
	v299 = *(*int32)(unsafe.Add(mBase, uint32(v48)+28))
	v301 = v299 << (uint(int32(2)) % 32)
	v304 = *(*int32)(unsafe.Add(mBase, uint32(v301)+uint32(_c_F_strbuf_free[4])))
	if v48 != v304 {
		goto L74
	} else {
		goto L75
	}
L73:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v290)+24)) = v245
	v321 = *(*int32)(unsafe.Add(mBase, uint32(v48)+16))
	if v321 == int32(0) {
		goto L81
	} else {
		goto L82
	}
L74:
	;
	v314 = *(*int32)(unsafe.Add(mBase, uint32(v245)+16))
	if v314 != v48 {
		goto L78
	} else {
		goto L79
	}
L75:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v301)+uint32(_c_F_strbuf_free[4]))) = v290
	if v290 != 0 {
		goto L73
	} else {
		goto L76
	}
L76:
	;
	v307 = int32(0)
	v309 = *(*int32)(unsafe.Add(mBase, _c_F_strbuf_free[5]))
	*(*int32)(unsafe.Add(mBase, _c_F_strbuf_free[5])) = v309 & base.I32_rotl(int32(-2), v299)
	goto L48
L77:
	;
	if v290 == int32(0) {
		goto L48
	} else {
		goto L80
	}
L78:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v245)+20)) = v290
	goto L77
L79:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v245)+16)) = v290
	goto L77
L80:
	;
	goto L73
L81:
	;
	v326 = *(*int32)(unsafe.Add(mBase, uint32(v48)+20))
	if v326 == int32(0) {
		goto L48
	} else {
		goto L83
	}
L82:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v290)+16)) = v321
	*(*int32)(unsafe.Add(mBase, uint32(v321)+24)) = v290
	goto L81
L83:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v290)+20)) = v326
	*(*int32)(unsafe.Add(mBase, uint32(v326)+24)) = v290
	goto L48
L84:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_strbuf_free[3])) = v228
	goto L7
L85:
	;
	if base.Ui32(int32(16777215)) < base.Ui32(v348) {
		v395 = int32(31)
		goto L90
	} else {
		goto L91
	}
L86:
	;
	v360 = v348 & int32(-8)
	v362 = v360 + int32(9128464)
	v364 = *(*int32)(unsafe.Add(mBase, _c_F_strbuf_free[2]))
	v368 = int32(1) << (uint(int32(base.Ui32(v348)>>(uint(int32(3))%32))) % 32)
	if v364&v368 != 0 {
		goto L88
	} else {
		goto L89
	}
L87:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v360)+uint32(_c_F_strbuf_free[8]))) = v173
	*(*int32)(unsafe.Add(mBase, uint32(v374)+12)) = v173
	*(*int32)(unsafe.Add(mBase, uint32(v173)+12)) = v362
	*(*int32)(unsafe.Add(mBase, uint32(v173)+8)) = v374
	goto L7
L88:
	;
	v373 = *(*int32)(unsafe.Add(mBase, uint32(v360)+uint32(_c_F_strbuf_free[8])))
	v374 = v373
	goto L87
L89:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_strbuf_free[2])) = v364 | v368
	v374 = v362
	goto L87
L90:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v173)+28)) = v395
	*(*int64)(unsafe.Add(mBase, uint32(v173)+16)) = int64(0)
	v400 = v395 << (uint(int32(2)) % 32)
	v404 = *(*int32)(unsafe.Add(mBase, _c_F_strbuf_free[5]))
	v406 = int32(1) << (uint(v395) % 32)
	if v404&v406 != 0 {
		goto L95
	} else {
		goto L96
	}
L91:
	;
	v385 = base.I32_clz(int32(base.Ui32(v348) >> (uint(int32(8)) % 32)))
	v388 = int32(1)
	v395 = int32(base.Ui32(v348)>>(uint(int32(38)-v385)%32))&v388 - v385<<(uint(v388)%32) + int32(62)
	goto L90
L92:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v173+v467))) = v470
	*(*int32)(unsafe.Add(mBase, uint32(v173)+12)) = v469
	*(*int32)(unsafe.Add(mBase, uint32(v173+v465))) = v468
	v479 = int32(0)
	v481 = *(*int32)(unsafe.Add(mBase, _c_F_strbuf_free[9]))
	v482 = int32(-1)
	v483 = v481 + v482
	if v483 != 0 {
		goto L104
	} else {
		goto L105
	}
L93:
	;
	v459 = *(*int32)(unsafe.Add(mBase, uint32(v429)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v459)+12)) = v173
	*(*int32)(unsafe.Add(mBase, uint32(v429)+8)) = v173
	v465 = int32(24)
	v467 = int32(8)
	v468 = int32(0)
	v469 = v429
	v470 = v459
	goto L92
L94:
	;
	v465 = v450
	v467 = v452
	v468 = v173
	v469 = v173
	v470 = v455
	goto L92
L95:
	;
	if v395 == int32(31) {
		goto L97
	} else {
		goto L98
	}
L96:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_strbuf_free[5])) = v404 | v406
	*(*int32)(unsafe.Add(mBase, uint32(v400)+uint32(_c_F_strbuf_free[4]))) = v173
	v450 = int32(8)
	v452 = int32(24)
	v455 = v400 + int32(9128728)
	goto L94
L97:
	;
	v421 = int32(0)
	goto L99
L98:
	;
	v421 = int32(25) - int32(base.Ui32(v395)>>(uint(int32(1))%32))
	goto L99
L99:
	;
	v423 = *(*int32)(unsafe.Add(mBase, uint32(v400)+uint32(_c_F_strbuf_free[4])))
	v426 = v348 << (uint(v421) % 32)
	v429 = v423
	goto L100
L100:
	;
	v433 = *(*int32)(unsafe.Add(mBase, uint32(v429)+4))
	if v433&int32(-8) == v348 {
		goto L93
	} else {
		goto L102
	}
L101:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v443+int32(16)))) = v173
	v450 = int32(8)
	v452 = int32(24)
	v455 = v429
	goto L94
L102:
	;
	v443 = v429 + int32(base.Ui32(v426)>>(uint(int32(29))%32))&int32(4)
	v444 = *(*int32)(unsafe.Add(mBase, uint32(v443)+16))
	if v444 != 0 {
		v426 = v426 << (uint(int32(1)) % 32)
		v429 = v444
		goto L100
	} else {
		goto L103
	}
L103:
	;
	goto L101
L104:
	;
	v485 = v483
	goto L106
L105:
	;
	v485 = v482
	goto L106
L106:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_strbuf_free[9])) = v485
	goto L8
L107:
	;
	m.G0 = v7 + int32(16)
	return
L108:
	;
	if l0 == int32(0) {
		goto L110
	} else {
		goto L111
	}
L109:
	;
	goto L107
L110:
	;
	goto L109
L111:
	;
	v520 = int32(-8)
	v521 = l0 + v520
	v524 = *(*int32)(unsafe.Add(mBase, uint32(l0+int32(-4))))
	v526 = v524 & v520
	v527 = v521 + v526
	if v524&int32(1) != 0 {
		v651 = v526
		v652 = v521
		goto L112
	} else {
		goto L113
	}
L112:
	;
	if base.Ui32(v527) <= base.Ui32(v652) {
		goto L110
	} else {
		goto L147
	}
L113:
	;
	if v524&int32(2) == int32(0) {
		goto L110
	} else {
		goto L114
	}
L114:
	;
	v534 = *(*int32)(unsafe.Add(mBase, uint32(v521)))
	v535 = v521 - v534
	v537 = *(*int32)(unsafe.Add(mBase, _c_F_strbuf_free[0]))
	if base.Ui32(v535) < base.Ui32(v537) {
		goto L110
	} else {
		goto L115
	}
L115:
	;
	v539 = v534 + v526
	v541 = *(*int32)(unsafe.Add(mBase, _c_F_strbuf_free[1]))
	if v535 == v541 {
		goto L119
	} else {
		goto L120
	}
L116:
	;
	if v557 == int32(0) {
		v651 = v539
		v652 = v535
		goto L112
	} else {
		goto L135
	}
L117:
	;
	v610 = int32(0)
	goto L116
L118:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v546)+12)) = v543
	*(*int32)(unsafe.Add(mBase, uint32(v543)+8)) = v546
	v651 = v539
	v652 = v535
	goto L112
L119:
	;
	v591 = *(*int32)(unsafe.Add(mBase, uint32(v527)+4))
	v592 = int32(3)
	if v591&v592 != v592 {
		v651 = v539
		v652 = v535
		goto L112
	} else {
		goto L134
	}
L120:
	;
	v543 = *(*int32)(unsafe.Add(mBase, uint32(v535)+12))
	if base.Ui32(int32(255)) < base.Ui32(v534) {
		goto L121
	} else {
		goto L122
	}
L121:
	;
	v557 = *(*int32)(unsafe.Add(mBase, uint32(v535)+24))
	if v543 == v535 {
		goto L124
	} else {
		goto L125
	}
L122:
	;
	v546 = *(*int32)(unsafe.Add(mBase, uint32(v535)+8))
	if v543 != v546 {
		goto L118
	} else {
		goto L123
	}
L123:
	;
	v548 = int32(0)
	v550 = *(*int32)(unsafe.Add(mBase, _c_F_strbuf_free[2]))
	*(*int32)(unsafe.Add(mBase, _c_F_strbuf_free[2])) = v550 & base.I32_rotl(int32(-2), int32(base.Ui32(v534)>>(uint(int32(3))%32)))
	v651 = v539
	v652 = v535
	goto L112
L124:
	;
	v562 = *(*int32)(unsafe.Add(mBase, uint32(v535)+20))
	if v562 == int32(0) {
		goto L127
	} else {
		goto L128
	}
L125:
	;
	v559 = *(*int32)(unsafe.Add(mBase, uint32(v535)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v559)+12)) = v543
	*(*int32)(unsafe.Add(mBase, uint32(v543)+8)) = v559
	v610 = v543
	goto L116
L126:
	;
	__phi578 = v572
	__phi579 = v573
	v578 = __phi578
	v579 = __phi579
	goto L130
L127:
	;
	v567 = *(*int32)(unsafe.Add(mBase, uint32(v535)+16))
	if v567 == int32(0) {
		goto L117
	} else {
		goto L129
	}
L128:
	;
	v572 = v562
	v573 = v535 + int32(20)
	goto L126
L129:
	;
	v572 = v567
	v573 = v535 + int32(16)
	goto L126
L130:
	;
	v585 = *(*int32)(unsafe.Add(mBase, uint32(v578)+20))
	if v585 != 0 {
		__phi578 = v585
		__phi579 = v578 + int32(20)
		v578 = __phi578
		v579 = __phi579
		goto L130
	} else {
		goto L132
	}
L131:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v579))) = int32(0)
	v610 = v578
	goto L116
L132:
	;
	v588 = *(*int32)(unsafe.Add(mBase, uint32(v578)+16))
	if v588 != 0 {
		__phi578 = v588
		__phi579 = v578 + int32(16)
		v578 = __phi578
		v579 = __phi579
		goto L130
	} else {
		goto L133
	}
L133:
	;
	goto L131
L134:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_strbuf_free[3])) = v539
	*(*int32)(unsafe.Add(mBase, uint32(v527)+4)) = v591 & int32(-2)
	*(*int32)(unsafe.Add(mBase, uint32(v535)+4)) = v539 | int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v527))) = v539
	goto L109
L135:
	;
	v619 = *(*int32)(unsafe.Add(mBase, uint32(v535)+28))
	v621 = v619 << (uint(int32(2)) % 32)
	v624 = *(*int32)(unsafe.Add(mBase, uint32(v621)+uint32(_c_F_strbuf_free[4])))
	if v535 != v624 {
		goto L137
	} else {
		goto L138
	}
L136:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v610)+24)) = v557
	v641 = *(*int32)(unsafe.Add(mBase, uint32(v535)+16))
	if v641 == int32(0) {
		goto L144
	} else {
		goto L145
	}
L137:
	;
	v634 = *(*int32)(unsafe.Add(mBase, uint32(v557)+16))
	if v634 != v535 {
		goto L141
	} else {
		goto L142
	}
L138:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v621)+uint32(_c_F_strbuf_free[4]))) = v610
	if v610 != 0 {
		goto L136
	} else {
		goto L139
	}
L139:
	;
	v627 = int32(0)
	v629 = *(*int32)(unsafe.Add(mBase, _c_F_strbuf_free[5]))
	*(*int32)(unsafe.Add(mBase, _c_F_strbuf_free[5])) = v629 & base.I32_rotl(int32(-2), v619)
	v651 = v539
	v652 = v535
	goto L112
L140:
	;
	if v610 == int32(0) {
		v651 = v539
		v652 = v535
		goto L112
	} else {
		goto L143
	}
L141:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v557)+20)) = v610
	goto L140
L142:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v557)+16)) = v610
	goto L140
L143:
	;
	goto L136
L144:
	;
	v646 = *(*int32)(unsafe.Add(mBase, uint32(v535)+20))
	if v646 == int32(0) {
		v651 = v539
		v652 = v535
		goto L112
	} else {
		goto L146
	}
L145:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v610)+16)) = v641
	*(*int32)(unsafe.Add(mBase, uint32(v641)+24)) = v610
	goto L144
L146:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v610)+20)) = v646
	*(*int32)(unsafe.Add(mBase, uint32(v646)+24)) = v610
	v651 = v539
	v652 = v535
	goto L112
L147:
	;
	v661 = *(*int32)(unsafe.Add(mBase, uint32(v527)+4))
	if v661&int32(1) == int32(0) {
		goto L110
	} else {
		goto L148
	}
L148:
	;
	if v661&int32(2) != 0 {
		goto L153
	} else {
		goto L154
	}
L149:
	;
	if base.Ui32(int32(255)) < base.Ui32(v827) {
		goto L187
	} else {
		goto L188
	}
L150:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v652)+4)) = v707 | int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v652+v707))) = v707
	if v652 != v691 {
		v827 = v707
		goto L149
	} else {
		goto L186
	}
L151:
	;
	if v724 == int32(0) {
		goto L150
	} else {
		goto L174
	}
L152:
	;
	v769 = int32(0)
	goto L151
L153:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v527)+4)) = v661 & int32(-2)
	*(*int32)(unsafe.Add(mBase, uint32(v652)+4)) = v651 | int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v652+v651))) = v651
	v827 = v651
	goto L149
L154:
	;
	v669 = *(*int32)(unsafe.Add(mBase, _c_F_strbuf_free[6]))
	if v527 != v669 {
		goto L155
	} else {
		goto L156
	}
L155:
	;
	v691 = *(*int32)(unsafe.Add(mBase, _c_F_strbuf_free[1]))
	if v527 != v691 {
		goto L158
	} else {
		goto L159
	}
L156:
	;
	v671 = int32(0)
	*(*int32)(unsafe.Add(mBase, _c_F_strbuf_free[6])) = v652
	v675 = *(*int32)(unsafe.Add(mBase, _c_F_strbuf_free[7]))
	v676 = v675 + v651
	*(*int32)(unsafe.Add(mBase, _c_F_strbuf_free[7])) = v676
	*(*int32)(unsafe.Add(mBase, uint32(v652)+4)) = v676 | int32(1)
	v682 = *(*int32)(unsafe.Add(mBase, _c_F_strbuf_free[1]))
	if v652 != v682 {
		goto L110
	} else {
		goto L157
	}
L157:
	;
	v684 = int32(0)
	*(*int32)(unsafe.Add(mBase, _c_F_strbuf_free[3])) = v684
	*(*int32)(unsafe.Add(mBase, _c_F_strbuf_free[1])) = v684
	goto L109
L158:
	;
	v707 = v661&int32(-8) + v651
	v708 = *(*int32)(unsafe.Add(mBase, uint32(v527)+12))
	if base.Ui32(int32(255)) < base.Ui32(v661) {
		goto L160
	} else {
		goto L161
	}
L159:
	;
	v693 = int32(0)
	*(*int32)(unsafe.Add(mBase, _c_F_strbuf_free[1])) = v652
	v697 = *(*int32)(unsafe.Add(mBase, _c_F_strbuf_free[3]))
	v698 = v697 + v651
	*(*int32)(unsafe.Add(mBase, _c_F_strbuf_free[3])) = v698
	*(*int32)(unsafe.Add(mBase, uint32(v652)+4)) = v698 | int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v652+v698))) = v698
	goto L109
L160:
	;
	v724 = *(*int32)(unsafe.Add(mBase, uint32(v527)+24))
	if v708 == v527 {
		goto L164
	} else {
		goto L165
	}
L161:
	;
	v711 = *(*int32)(unsafe.Add(mBase, uint32(v527)+8))
	if v708 != v711 {
		goto L162
	} else {
		goto L163
	}
L162:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v711)+12)) = v708
	*(*int32)(unsafe.Add(mBase, uint32(v708)+8)) = v711
	goto L150
L163:
	;
	v713 = int32(0)
	v715 = *(*int32)(unsafe.Add(mBase, _c_F_strbuf_free[2]))
	*(*int32)(unsafe.Add(mBase, _c_F_strbuf_free[2])) = v715 & base.I32_rotl(int32(-2), int32(base.Ui32(v661)>>(uint(int32(3))%32)))
	goto L150
L164:
	;
	v729 = *(*int32)(unsafe.Add(mBase, uint32(v527)+20))
	if v729 == int32(0) {
		goto L167
	} else {
		goto L168
	}
L165:
	;
	v726 = *(*int32)(unsafe.Add(mBase, uint32(v527)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v726)+12)) = v708
	*(*int32)(unsafe.Add(mBase, uint32(v708)+8)) = v726
	v769 = v708
	goto L151
L166:
	;
	__phi745 = v739
	__phi746 = v740
	v745 = __phi745
	v746 = __phi746
	goto L170
L167:
	;
	v734 = *(*int32)(unsafe.Add(mBase, uint32(v527)+16))
	if v734 == int32(0) {
		goto L152
	} else {
		goto L169
	}
L168:
	;
	v739 = v729
	v740 = v527 + int32(20)
	goto L166
L169:
	;
	v739 = v734
	v740 = v527 + int32(16)
	goto L166
L170:
	;
	v752 = *(*int32)(unsafe.Add(mBase, uint32(v745)+20))
	if v752 != 0 {
		__phi745 = v752
		__phi746 = v745 + int32(20)
		v745 = __phi745
		v746 = __phi746
		goto L170
	} else {
		goto L172
	}
L171:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v746))) = int32(0)
	v769 = v745
	goto L151
L172:
	;
	v755 = *(*int32)(unsafe.Add(mBase, uint32(v745)+16))
	if v755 != 0 {
		__phi745 = v755
		__phi746 = v745 + int32(16)
		v745 = __phi745
		v746 = __phi746
		goto L170
	} else {
		goto L173
	}
L173:
	;
	goto L171
L174:
	;
	v778 = *(*int32)(unsafe.Add(mBase, uint32(v527)+28))
	v780 = v778 << (uint(int32(2)) % 32)
	v783 = *(*int32)(unsafe.Add(mBase, uint32(v780)+uint32(_c_F_strbuf_free[4])))
	if v527 != v783 {
		goto L176
	} else {
		goto L177
	}
L175:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v769)+24)) = v724
	v800 = *(*int32)(unsafe.Add(mBase, uint32(v527)+16))
	if v800 == int32(0) {
		goto L183
	} else {
		goto L184
	}
L176:
	;
	v793 = *(*int32)(unsafe.Add(mBase, uint32(v724)+16))
	if v793 != v527 {
		goto L180
	} else {
		goto L181
	}
L177:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v780)+uint32(_c_F_strbuf_free[4]))) = v769
	if v769 != 0 {
		goto L175
	} else {
		goto L178
	}
L178:
	;
	v786 = int32(0)
	v788 = *(*int32)(unsafe.Add(mBase, _c_F_strbuf_free[5]))
	*(*int32)(unsafe.Add(mBase, _c_F_strbuf_free[5])) = v788 & base.I32_rotl(int32(-2), v778)
	goto L150
L179:
	;
	if v769 == int32(0) {
		goto L150
	} else {
		goto L182
	}
L180:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v724)+20)) = v769
	goto L179
L181:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v724)+16)) = v769
	goto L179
L182:
	;
	goto L175
L183:
	;
	v805 = *(*int32)(unsafe.Add(mBase, uint32(v527)+20))
	if v805 == int32(0) {
		goto L150
	} else {
		goto L185
	}
L184:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v769)+16)) = v800
	*(*int32)(unsafe.Add(mBase, uint32(v800)+24)) = v769
	goto L183
L185:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v769)+20)) = v805
	*(*int32)(unsafe.Add(mBase, uint32(v805)+24)) = v769
	goto L150
L186:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_strbuf_free[3])) = v707
	goto L109
L187:
	;
	if base.Ui32(int32(16777215)) < base.Ui32(v827) {
		v874 = int32(31)
		goto L192
	} else {
		goto L193
	}
L188:
	;
	v839 = v827 & int32(-8)
	v841 = v839 + int32(9128464)
	v843 = *(*int32)(unsafe.Add(mBase, _c_F_strbuf_free[2]))
	v847 = int32(1) << (uint(int32(base.Ui32(v827)>>(uint(int32(3))%32))) % 32)
	if v843&v847 != 0 {
		goto L190
	} else {
		goto L191
	}
L189:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v839)+uint32(_c_F_strbuf_free[8]))) = v652
	*(*int32)(unsafe.Add(mBase, uint32(v853)+12)) = v652
	*(*int32)(unsafe.Add(mBase, uint32(v652)+12)) = v841
	*(*int32)(unsafe.Add(mBase, uint32(v652)+8)) = v853
	goto L109
L190:
	;
	v852 = *(*int32)(unsafe.Add(mBase, uint32(v839)+uint32(_c_F_strbuf_free[8])))
	v853 = v852
	goto L189
L191:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_strbuf_free[2])) = v843 | v847
	v853 = v841
	goto L189
L192:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v652)+28)) = v874
	*(*int64)(unsafe.Add(mBase, uint32(v652)+16)) = int64(0)
	v879 = v874 << (uint(int32(2)) % 32)
	v883 = *(*int32)(unsafe.Add(mBase, _c_F_strbuf_free[5]))
	v885 = int32(1) << (uint(v874) % 32)
	if v883&v885 != 0 {
		goto L197
	} else {
		goto L198
	}
L193:
	;
	v864 = base.I32_clz(int32(base.Ui32(v827) >> (uint(int32(8)) % 32)))
	v867 = int32(1)
	v874 = int32(base.Ui32(v827)>>(uint(int32(38)-v864)%32))&v867 - v864<<(uint(v867)%32) + int32(62)
	goto L192
L194:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v652+v946))) = v949
	*(*int32)(unsafe.Add(mBase, uint32(v652)+12)) = v948
	*(*int32)(unsafe.Add(mBase, uint32(v652+v944))) = v947
	v958 = int32(0)
	v960 = *(*int32)(unsafe.Add(mBase, _c_F_strbuf_free[9]))
	v961 = int32(-1)
	v962 = v960 + v961
	if v962 != 0 {
		goto L206
	} else {
		goto L207
	}
L195:
	;
	v938 = *(*int32)(unsafe.Add(mBase, uint32(v908)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v938)+12)) = v652
	*(*int32)(unsafe.Add(mBase, uint32(v908)+8)) = v652
	v944 = int32(24)
	v946 = int32(8)
	v947 = int32(0)
	v948 = v908
	v949 = v938
	goto L194
L196:
	;
	v944 = v929
	v946 = v931
	v947 = v652
	v948 = v652
	v949 = v934
	goto L194
L197:
	;
	if v874 == int32(31) {
		goto L199
	} else {
		goto L200
	}
L198:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_strbuf_free[5])) = v883 | v885
	*(*int32)(unsafe.Add(mBase, uint32(v879)+uint32(_c_F_strbuf_free[4]))) = v652
	v929 = int32(8)
	v931 = int32(24)
	v934 = v879 + int32(9128728)
	goto L196
L199:
	;
	v900 = int32(0)
	goto L201
L200:
	;
	v900 = int32(25) - int32(base.Ui32(v874)>>(uint(int32(1))%32))
	goto L201
L201:
	;
	v902 = *(*int32)(unsafe.Add(mBase, uint32(v879)+uint32(_c_F_strbuf_free[4])))
	v905 = v827 << (uint(v900) % 32)
	v908 = v902
	goto L202
L202:
	;
	v912 = *(*int32)(unsafe.Add(mBase, uint32(v908)+4))
	if v912&int32(-8) == v827 {
		goto L195
	} else {
		goto L204
	}
L203:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v922+int32(16)))) = v652
	v929 = int32(8)
	v931 = int32(24)
	v934 = v908
	goto L196
L204:
	;
	v922 = v908 + int32(base.Ui32(v905)>>(uint(int32(29))%32))&int32(4)
	v923 = *(*int32)(unsafe.Add(mBase, uint32(v922)+16))
	if v923 != 0 {
		v905 = v905 << (uint(int32(1)) % 32)
		v908 = v923
		goto L202
	} else {
		goto L205
	}
L205:
	;
	goto L203
L206:
	;
	v964 = v962
	goto L208
L207:
	;
	v964 = v961
	goto L208
L208:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_strbuf_free[9])) = v964
	goto L110
}
func F_strcasecmp(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v15 int32
	_ = v15
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v63 int32
	_ = v63
	v5 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v5 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	if base.Ui32(v45+int32(-65)) < base.Ui32(int32(26)) {
		goto L21
	} else {
		goto L22
	}
L2:
	;
	v7 = l0
	v8 = l1
	v9 = v5
	goto L5
L3:
	;
	v45 = int32(0)
	v46 = l1
	goto L1
L4:
	;
	v45 = v42 & int32(255)
	v46 = v41
	goto L1
L5:
	;
	v11 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8))))
	if v11 == int32(0) {
		v41 = v8
		v42 = v9
		goto L4
	} else {
		goto L7
	}
L6:
	;
	v41 = v35
	v42 = int32(0)
	goto L4
L7:
	;
	v15 = v9 & int32(255)
	if v15 == v11 {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v34 = int32(1)
	v35 = v8 + v34
	v36 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7)+1)))
	if v36 != 0 {
		v7 = v7 + v34
		v8 = v35
		v9 = v36
		goto L5
	} else {
		goto L19
	}
L9:
	;
	if base.Ui32(v15+int32(-65)) < base.Ui32(int32(26)) {
		goto L11
	} else {
		goto L12
	}
L10:
	;
	v24 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8))))
	if base.Ui32(v24+int32(-65)) < base.Ui32(int32(26)) {
		goto L15
	} else {
		goto L16
	}
L11:
	;
	v23 = v15 | int32(32)
	goto L13
L12:
	;
	v23 = v15
	goto L13
L13:
	;
	goto L10
L14:
	;
	if v23 == v31 {
		goto L8
	} else {
		goto L18
	}
L15:
	;
	v31 = v24 | int32(32)
	goto L17
L16:
	;
	v31 = v24
	goto L17
L17:
	;
	goto L14
L18:
	;
	v33 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7))))
	v41 = v8
	v42 = v33
	goto L4
L19:
	;
	goto L6
L20:
	;
	v56 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v46))))
	if base.Ui32(v56+int32(-65)) < base.Ui32(int32(26)) {
		goto L25
	} else {
		goto L26
	}
L21:
	;
	v55 = v45 | int32(32)
	goto L23
L22:
	;
	v55 = v45
	goto L23
L23:
	;
	goto L20
L24:
	;
	return v55 - v63
L25:
	;
	v63 = v56 | int32(32)
	goto L27
L26:
	;
	v63 = v56
	goto L27
L27:
	;
	goto L24
}
func F_strcspn(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v18 int32
	_ = v18
	var v27 int32
	_ = v27
	var v32 int32
	_ = v32
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v74 int32
	_ = v74
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v89 int32
	_ = v89
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v103 int32
	_ = v103
	var v113 int32
	_ = v113
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v122 int32
	_ = v122
	var v124 int32
	_ = v124
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v136 int32
	_ = v136
	var v144 int32
	_ = v144
	var v148 int32
	_ = v148
	var v150 int32
	_ = v150
	var v157 int32
	_ = v157
	var v163 int32
	_ = v163
	var v165 int32
	_ = v165
	var v170 int32
	_ = v170
	v6 = m.G0
	v8 = v6 - int32(32)
	m.G0 = v8
	v10 = int32(*(*int8)(unsafe.Add(mBase, uint32(l1))))
	if v10 == int32(0) {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	m.G0 = v8 + int32(32)
	return v170 - l0
L2:
	;
	v117 = F__emscripten_memset_bulkmem(m, v8, base.I32_extend8_s(int32(0)), int32(32))
	mBase = m.M
	goto L28
L3:
	;
	v18 = v10 & int32(255)
	if v18 == int32(0) {
		goto L10
	} else {
		goto L11
	}
L4:
	;
	v13 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+1)))
	if v13 != 0 {
		goto L2
	} else {
		goto L5
	}
L5:
	;
	goto L3
L6:
	;
	v170 = v113
	goto L1
L7:
	;
	v113 = v103
	goto L6
L8:
	;
	v94 = v89
	goto L24
L9:
	;
	v89 = v80
	goto L8
L10:
	;
	v78 = F_strlen(m, l0)
	mBase = m.M
	v113 = l0 + v78
	goto L6
L11:
	;
	if l0&int32(3) == int32(0) {
		v40 = l0
		goto L12
	} else {
		goto L13
	}
L12:
	;
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v40)))
	v49 = int32(-2139062144)
	if (int32(16843008)-v46|v46)&v49 != v49 {
		v80 = v40
		goto L9
	} else {
		goto L19
	}
L13:
	;
	v27 = l0
	goto L14
L14:
	;
	v32 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v27))))
	if v32 == int32(0) {
		v103 = v27
		goto L7
	} else {
		goto L16
	}
L15:
	;
	v40 = v37
	goto L12
L16:
	;
	if v32 == v10&int32(255) {
		v103 = v27
		goto L7
	} else {
		goto L17
	}
L17:
	;
	v37 = v27 + int32(1)
	if v37&int32(3) != 0 {
		v27 = v37
		goto L14
	} else {
		goto L18
	}
L18:
	;
	goto L15
L19:
	;
	v55 = v40
	v58 = v46
	goto L20
L20:
	;
	v61 = v58 ^ v18*int32(16843009)
	v64 = int32(-2139062144)
	if (int32(16843008)-v61|v61)&v64 != v64 {
		v80 = v55
		goto L9
	} else {
		goto L22
	}
L22:
	;
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v55)+4))
	v70 = v55 + int32(4)
	v74 = int32(-2139062144)
	if (v68|(int32(16843008)-v68))&v74 == v74 {
		v55 = v70
		v58 = v68
		goto L20
	} else {
		goto L23
	}
L23:
	;
	v89 = v70
	goto L8
L24:
	;
	v95 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v94))))
	if v95 == int32(0) {
		v103 = v94
		goto L7
	} else {
		goto L26
	}
L25:
	;
	v103 = v94
	goto L7
L26:
	;
	if v95 != v10&int32(255) {
		v94 = v94 + int32(1)
		goto L24
	} else {
		goto L27
	}
L27:
	;
	goto L25
L28:
	;
	v118 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	if v118 == int32(0) {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	v144 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v144 == int32(0) {
		v170 = l0
		goto L1
	} else {
		goto L34
	}
L30:
	;
	v122 = l1
	v124 = v118
	goto L31
L31:
	;
	v130 = v8 + int32(base.Ui32(v124)>>(uint(int32(3))%32))&int32(28)
	v131 = *(*int32)(unsafe.Add(mBase, uint32(v130)))
	v132 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v130))) = v131 | v132<<(uint(v124)%32)
	v136 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v122)+1)))
	if v136 != 0 {
		v122 = v122 + v132
		v124 = v136
		goto L31
	} else {
		goto L33
	}
L32:
	;
	goto L29
L33:
	;
	goto L32
L34:
	;
	v148 = l0
	v150 = v144
	goto L35
L35:
	;
	v157 = *(*int32)(unsafe.Add(mBase, uint32(v8+int32(base.Ui32(v150)>>(uint(int32(3))%32))&int32(28))))
	if int32(base.Ui32(v157)>>(uint(v150)%32))&int32(1) == int32(0) {
		goto L37
	} else {
		goto L38
	}
L36:
	;
	v170 = v165
	goto L1
L37:
	;
	v163 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v148)+1)))
	v165 = v148 + int32(1)
	if v163 != 0 {
		v148 = v165
		v150 = v163
		goto L35
	} else {
		goto L39
	}
L38:
	;
	v170 = v148
	goto L1
L39:
	;
	goto L36
}
func F_string2l(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v21 int32
	_ = v21
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v50 int64
	_ = v50
	var v56 int32
	_ = v56
	var v58 int64
	_ = v58
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v72 int64
	_ = v72
	var v77 int64
	_ = v77
	var v81 int32
	_ = v81
	var v83 int64
	_ = v83
	var v85 int32
	_ = v85
	var v93 int64
	_ = v93
	var v113 int64
	_ = v113
	var v126 int32
	_ = v126
	v4 = int32(0)
	if base.Ui32(l1+int32(-21)) < base.Ui32(int32(-20)) {
		v126 = v4
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return v126
L2:
	;
	v16 = int32(1)
	v17 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if l1 != v16 {
		goto L4
	} else {
		goto L5
	}
L3:
	;
	if base.Ui64(v113+int64(-2147483648)) < base.Ui64(int64(-4294967296)) {
		v126 = v4
		goto L1
	} else {
		goto L23
	}
L4:
	;
	if v17&int32(255) == int32(45) {
		goto L8
	} else {
		goto L9
	}
L5:
	;
	v21 = v17 + int32(-48)
	if base.Ui32(int32(9)) < base.Ui32(v21&int32(255)) {
		v126 = v4
		goto L1
	} else {
		goto L6
	}
L6:
	;
	v113 = base.I64_extend_i32_u(v21) & int64(255)
	goto L3
L7:
	;
	if base.Ui32(int32(8)) < base.Ui32((v38+int32(-49))&int32(255)) {
		v126 = v4
		goto L1
	} else {
		goto L10
	}
L8:
	;
	v35 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+1)))
	v37 = int32(2)
	v38 = v35
	v39 = l0 + int32(1)
	goto L7
L9:
	;
	v37 = v16
	v38 = v17
	v39 = l0
	goto L7
L10:
	;
	v50 = base.I64_extend_i32_u(v38+int32(-48)) & int64(255)
	if base.Ui32(l1) <= base.Ui32(v37) {
		v93 = v50
		goto L11
	} else {
		goto L12
	}
L11:
	;
	if v17&int32(255) != int32(45) {
		goto L19
	} else {
		goto L20
	}
L12:
	;
	v56 = v37
	v58 = v50
	v60 = v39
	goto L13
L13:
	;
	v62 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v60)+1)))
	if base.Ui32((v62+int32(-58))&int32(255)) < base.Ui32(int32(246)) {
		v126 = v4
		goto L1
	} else {
		goto L15
	}
L14:
	;
	v93 = v83
	goto L11
L15:
	;
	if base.Ui64(int64(1844674407370955161)) < base.Ui64(v58) {
		v126 = v4
		goto L1
	} else {
		goto L16
	}
L16:
	;
	v72 = v58 * int64(10)
	v77 = base.I64_extend_i32_u(v62+int32(-48)) & int64(255)
	if base.Ui64(v77^int64(-1)) < base.Ui64(v72) {
		v126 = v4
		goto L1
	} else {
		goto L17
	}
L17:
	;
	v81 = int32(1)
	v83 = v72 + v77
	v85 = v56 + v81
	if v85 != l1 {
		v56 = v85
		v58 = v83
		v60 = v60 + v81
		goto L13
	} else {
		goto L18
	}
L18:
	;
	goto L14
L19:
	;
	if v93 < int64(0) {
		v126 = v4
		goto L1
	} else {
		goto L22
	}
L20:
	;
	if base.Ui64(int64(-9223372036854775807-1)) < base.Ui64(v93) {
		v126 = v4
		goto L1
	} else {
		goto L21
	}
L21:
	;
	v113 = int64(0) - v93
	goto L3
L22:
	;
	v113 = v93
	goto L3
L23:
	;
	*(*uint32)(unsafe.Add(mBase, uint32(l2))) = uint32(v113)
	v126 = int32(1)
	goto L1
}
func F_string2ull(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v24 int32
	_ = v24
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v53 int64
	_ = v53
	var v59 int32
	_ = v59
	var v61 int64
	_ = v61
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v75 int64
	_ = v75
	var v80 int64
	_ = v80
	var v84 int32
	_ = v84
	var v86 int64
	_ = v86
	var v88 int32
	_ = v88
	var v96 int64
	_ = v96
	var v107 int64
	_ = v107
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v134 int64
	_ = v134
	var v136 int32
	_ = v136
	var v141 int32
	_ = v141
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v154 int64
	_ = v154
	var v161 int32
	_ = v161
	v11 = m.G0
	v13 = v11 - int32(16)
	m.G0 = v13
	if base.Ui32(l1+int32(-21)) < base.Ui32(int32(-20)) {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	m.G0 = v13 + int32(16)
	return v161
L2:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l2))) = v154
	v161 = int32(1)
	goto L1
L3:
	;
	v124 = int32(0)
	v125 = int32(9116376)
	goto L24
L4:
	;
	v19 = int32(1)
	v20 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if l1 != v19 {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	if v20&int32(255) != int32(45) {
		v40 = v19
		v41 = v20
		v42 = l0
		goto L8
	} else {
		goto L9
	}
L6:
	;
	v24 = v20 + int32(-48)
	if base.Ui32(int32(9)) < base.Ui32(v24&int32(255)) {
		goto L3
	} else {
		goto L7
	}
L7:
	;
	v154 = base.I64_extend_i32_u(v24) & int64(255)
	goto L2
L8:
	;
	if base.Ui32(int32(8)) < base.Ui32((v41+int32(-49))&int32(255)) {
		goto L3
	} else {
		goto L10
	}
L9:
	;
	v38 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+1)))
	v40 = int32(2)
	v41 = v38
	v42 = l0 + int32(1)
	goto L8
L10:
	;
	v53 = base.I64_extend_i32_u(v41+int32(-48)) & int64(255)
	if base.Ui32(l1) <= base.Ui32(v40) {
		v96 = v53
		goto L11
	} else {
		goto L12
	}
L11:
	;
	if v20&int32(255) != int32(45) {
		goto L19
	} else {
		goto L20
	}
L12:
	;
	v59 = v40
	v61 = v53
	v63 = v42
	goto L13
L13:
	;
	v65 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v63)+1)))
	if base.Ui32((v65+int32(-58))&int32(255)) < base.Ui32(int32(246)) {
		goto L3
	} else {
		goto L15
	}
L14:
	;
	v96 = v86
	goto L11
L15:
	;
	if base.Ui64(int64(1844674407370955161)) < base.Ui64(v61) {
		goto L3
	} else {
		goto L16
	}
L16:
	;
	v75 = v61 * int64(10)
	v80 = base.I64_extend_i32_u(v65+int32(-48)) & int64(255)
	if base.Ui64(v80^int64(-1)) < base.Ui64(v75) {
		goto L3
	} else {
		goto L17
	}
L17:
	;
	v84 = int32(1)
	v86 = v75 + v80
	v88 = v59 + v84
	if v88 != l1 {
		v59 = v88
		v61 = v86
		v63 = v63 + v84
		goto L13
	} else {
		goto L18
	}
L18:
	;
	goto L14
L19:
	;
	if int64(0) <= v96 {
		v154 = v96
		goto L2
	} else {
		goto L23
	}
L20:
	;
	if base.Ui64(int64(-9223372036854775807-1)) < base.Ui64(v96) {
		goto L3
	} else {
		goto L21
	}
L21:
	;
	v107 = int64(-1)
	if v107 < v96+v107 {
		v161 = int32(0)
		goto L1
	} else {
		goto L22
	}
L22:
	;
	v154 = int64(0)
	goto L2
L23:
	;
	goto L3
L24:
	;
	v126 = int32(0)
	*(*int32)(unsafe.Add(mBase, _c_F_string2ull[0])) = v126
	*(*int32)(unsafe.Add(mBase, uint32(v13)+12)) = v126
	v134 = F_strtox_2(m, l0, v13+int32(12), int32(10), int64(-1))
	mBase = m.M
	goto L25
L25:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l2))) = v134
	v136 = *(*int32)(unsafe.Add(mBase, _c_F_string2ull[0]))
	if v136 == int32(28) {
		v161 = v124
		goto L1
	} else {
		goto L26
	}
L26:
	;
	if v136 == int32(68) {
		v161 = v124
		goto L1
	} else {
		goto L27
	}
L27:
	;
	v141 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v141 == int32(0) {
		v161 = v124
		goto L1
	} else {
		goto L28
	}
L28:
	;
	v144 = *(*int32)(unsafe.Add(mBase, uint32(v13)+12))
	v145 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v144))))
	v161 = base.B2i32(v145 == int32(0))
	goto L1
}
func F_stringmatch(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v25 int32
	_ = v25
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v51 int32
	_ = v51
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v65 int32
	_ = v65
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v80 int32
	_ = v80
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v91 int32
	_ = v91
	var v97 int32
	_ = v97
	var v100 int32
	_ = v100
	var v106 int32
	_ = v106
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v126 int32
	_ = v126
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	if l0&int32(3) == int32(0) {
		v32 = l0
		goto L3
	} else {
		goto L4
	}
L1:
	;
	if l1&int32(3) == int32(0) {
		v87 = l1
		goto L19
	} else {
		goto L20
	}
L2:
	;
	v65 = v57 - l0
	goto L1
L3:
	;
	v36 = v32
	goto L11
L4:
	;
	v18 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v18 != 0 {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	v21 = l0
	goto L7
L6:
	;
	v65 = l0 - l0
	goto L1
L7:
	;
	v25 = v21 + int32(1)
	if v25&int32(3) == int32(0) {
		v32 = v25
		goto L3
	} else {
		goto L9
	}
L9:
	;
	v30 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v25))))
	if v30 != 0 {
		v21 = v25
		goto L7
	} else {
		goto L10
	}
L10:
	;
	v57 = v25
	goto L2
L11:
	;
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v36)))
	v45 = int32(-2139062144)
	if (int32(16843008)-v42|v42)&v45 == v45 {
		v36 = v36 + int32(4)
		goto L11
	} else {
		goto L13
	}
L12:
	;
	v51 = v36
	goto L14
L13:
	;
	goto L12
L14:
	;
	v55 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v51))))
	if v55 != 0 {
		v51 = v51 + int32(1)
		goto L14
	} else {
		goto L16
	}
L15:
	;
	v57 = v51
	goto L2
L16:
	;
	goto L15
L17:
	;
	v121 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v9)+12)) = v121
	v126 = F_stringmatchlen_impl(m, l0, v65, l1, v120, l2, v9+int32(12), v121)
	mBase = m.M
	m.G0 = v9 + int32(16)
	return v126
L18:
	;
	v120 = v112 - l1
	goto L17
L19:
	;
	v91 = v87
	goto L27
L20:
	;
	v73 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	if v73 != 0 {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	v76 = l1
	goto L23
L22:
	;
	v120 = l1 - l1
	goto L17
L23:
	;
	v80 = v76 + int32(1)
	if v80&int32(3) == int32(0) {
		v87 = v80
		goto L19
	} else {
		goto L25
	}
L25:
	;
	v85 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v80))))
	if v85 != 0 {
		v76 = v80
		goto L23
	} else {
		goto L26
	}
L26:
	;
	v112 = v80
	goto L18
L27:
	;
	v97 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
	v100 = int32(-2139062144)
	if (int32(16843008)-v97|v97)&v100 == v100 {
		v91 = v91 + int32(4)
		goto L27
	} else {
		goto L29
	}
L28:
	;
	v106 = v91
	goto L30
L29:
	;
	goto L28
L30:
	;
	v110 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v106))))
	if v110 != 0 {
		v106 = v106 + int32(1)
		goto L30
	} else {
		goto L32
	}
L31:
	;
	v112 = v106
	goto L18
L32:
	;
	goto L31
}
func F_strlen(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v13 int32
	_ = v13
	var v17 int32
	_ = v17
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	if l0&int32(3) == int32(0) {
		v24 = l0
		goto L2
	} else {
		goto L3
	}
L1:
	;
	return v49 - l0
L2:
	;
	v28 = v24
	goto L10
L3:
	;
	v9 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v9 != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v13 = l0
	goto L6
L5:
	;
	return l0 - l0
L6:
	;
	v17 = v13 + int32(1)
	if v17&int32(3) == int32(0) {
		v24 = v17
		goto L2
	} else {
		goto L8
	}
L8:
	;
	v22 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17))))
	if v22 != 0 {
		v13 = v17
		goto L6
	} else {
		goto L9
	}
L9:
	;
	v49 = v17
	goto L1
L10:
	;
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v28)))
	v37 = int32(-2139062144)
	if (int32(16843008)-v34|v34)&v37 == v37 {
		v28 = v28 + int32(4)
		goto L10
	} else {
		goto L12
	}
L11:
	;
	v43 = v28
	goto L13
L12:
	;
	goto L11
L13:
	;
	v47 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v43))))
	if v47 != 0 {
		v43 = v43 + int32(1)
		goto L13
	} else {
		goto L15
	}
L14:
	;
	v49 = v43
	goto L1
L15:
	;
	goto L14
}
func F_strncasecmp(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
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
	var v15 int32
	_ = v15
	var v19 int32
	_ = v19
	var v23 int32
	_ = v23
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v73 int32
	_ = v73
	if l2 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v8 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v8 != 0 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	return int32(0)
L3:
	;
	if base.Ui32(v54+int32(-65)) < base.Ui32(int32(26)) {
		goto L24
	} else {
		goto L25
	}
L4:
	;
	v10 = l0
	v11 = l1
	v12 = l2
	v13 = v8
	goto L7
L5:
	;
	v54 = int32(0)
	v55 = l1
	goto L3
L6:
	;
	v54 = v51 & int32(255)
	v55 = v49
	goto L3
L7:
	;
	v15 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11))))
	if v15 == int32(0) {
		v49 = v11
		v51 = v13
		goto L6
	} else {
		goto L9
	}
L8:
	;
	v49 = v43
	v51 = int32(0)
	goto L6
L9:
	;
	v19 = v12 + int32(-1)
	if v19 == int32(0) {
		v49 = v11
		v51 = v13
		goto L6
	} else {
		goto L10
	}
L10:
	;
	v23 = v13 & int32(255)
	if v23 == v15 {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v42 = int32(1)
	v43 = v11 + v42
	v44 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+1)))
	if v44 != 0 {
		v10 = v10 + v42
		v11 = v43
		v12 = v19
		v13 = v44
		goto L7
	} else {
		goto L22
	}
L12:
	;
	if base.Ui32(v23+int32(-65)) < base.Ui32(int32(26)) {
		goto L14
	} else {
		goto L15
	}
L13:
	;
	v32 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11))))
	if base.Ui32(v32+int32(-65)) < base.Ui32(int32(26)) {
		goto L18
	} else {
		goto L19
	}
L14:
	;
	v31 = v23 | int32(32)
	goto L16
L15:
	;
	v31 = v23
	goto L16
L16:
	;
	goto L13
L17:
	;
	if v31 == v39 {
		goto L11
	} else {
		goto L21
	}
L18:
	;
	v39 = v32 | int32(32)
	goto L20
L19:
	;
	v39 = v32
	goto L20
L20:
	;
	goto L17
L21:
	;
	v41 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10))))
	v49 = v11
	v51 = v41
	goto L6
L22:
	;
	goto L8
L23:
	;
	v66 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v55))))
	if base.Ui32(v66+int32(-65)) < base.Ui32(int32(26)) {
		goto L28
	} else {
		goto L29
	}
L24:
	;
	v65 = v54 | int32(32)
	goto L26
L25:
	;
	v65 = v54
	goto L26
L26:
	;
	goto L23
L27:
	;
	return v65 - v73
L28:
	;
	v73 = v66 | int32(32)
	goto L30
L29:
	;
	v73 = v66
	goto L30
L30:
	;
	goto L27
}
func F_strnlen(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v8 int32
	_ = v8
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v41 int32
	_ = v41
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v95 int32
	_ = v95
	var v107 int32
	_ = v107
	var v109 int32
	_ = v109
	v3 = int32(0)
	v8 = base.B2i32(l1 != v3)
	if l0&int32(3) == v3 {
		v34 = l0
		v36 = l1
		v37 = v8
		goto L4
	} else {
		goto L5
	}
L1:
	;
	if v107 != 0 {
		goto L26
	} else {
		goto L27
	}
L2:
	;
	v107 = int32(0)
	goto L1
L3:
	;
	v85 = v78
	v87 = v80
	goto L21
L4:
	;
	if v37 == int32(0) {
		goto L2
	} else {
		goto L12
	}
L5:
	;
	if l1 == int32(0) {
		v34 = l0
		v36 = l1
		v37 = v8
		goto L4
	} else {
		goto L6
	}
L6:
	;
	v17 = l0
	v19 = l1
	goto L7
L7:
	;
	v22 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17))))
	if v22 == int32(0) {
		v78 = v17
		v80 = v19
		goto L3
	} else {
		goto L9
	}
L8:
	;
	v34 = v29
	v36 = v25
	v37 = v27
	goto L4
L9:
	;
	v25 = v19 + int32(-1)
	v26 = int32(0)
	v27 = base.B2i32(v25 != v26)
	v29 = v17 + int32(1)
	if v29&int32(3) == v26 {
		v34 = v29
		v36 = v25
		v37 = v27
		goto L4
	} else {
		goto L10
	}
L10:
	;
	if v25 != 0 {
		v17 = v29
		v19 = v25
		goto L7
	} else {
		goto L11
	}
L11:
	;
	goto L8
L12:
	;
	v41 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v34))))
	if v41 == int32(0) {
		v71 = v34
		v73 = v36
		goto L13
	} else {
		goto L14
	}
L13:
	;
	if v73 == int32(0) {
		goto L2
	} else {
		goto L20
	}
L14:
	;
	if base.Ui32(v36) < base.Ui32(int32(4)) {
		v71 = v34
		v73 = v36
		goto L13
	} else {
		goto L15
	}
L15:
	;
	v51 = v34
	v53 = v36
	goto L16
L16:
	;
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v51)))
	v58 = v57 ^ int32(0)
	v61 = int32(-2139062144)
	if (int32(16843008)-v58|v58)&v61 != v61 {
		v78 = v51
		v80 = v53
		goto L3
	} else {
		goto L18
	}
L17:
	;
	v71 = v66
	v73 = v68
	goto L13
L18:
	;
	v66 = v51 + int32(4)
	v68 = v53 + int32(-4)
	if base.Ui32(int32(3)) < base.Ui32(v68) {
		v51 = v66
		v53 = v68
		goto L16
	} else {
		goto L19
	}
L19:
	;
	goto L17
L20:
	;
	v78 = v71
	v80 = v73
	goto L3
L21:
	;
	v90 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v85))))
	if v90 != int32(0) {
		goto L23
	} else {
		goto L24
	}
L22:
	;
	goto L2
L23:
	;
	v95 = v87 + int32(-1)
	if v95 != 0 {
		v85 = v85 + int32(1)
		v87 = v95
		goto L21
	} else {
		goto L25
	}
L24:
	;
	v107 = v85
	goto L1
L25:
	;
	goto L22
L26:
	;
	v109 = v107 - l0
	goto L28
L27:
	;
	v109 = l1
	goto L28
L28:
	;
	return v109
}
func F_strpbrk(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v57 int32
	_ = v57
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v70 int32
	_ = v70
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	v6 = m.G0
	v8 = v6 - int32(32)
	m.G0 = v8
	v10 = int32(*(*int8)(unsafe.Add(mBase, uint32(l1))))
	if v10 == int32(0) {
		goto L4
	} else {
		goto L5
	}
L1:
	;
	v75 = l0 + (v70 - l0)
	v77 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v75))))
	if v77 != 0 {
		goto L18
	} else {
		goto L19
	}
L2:
	;
	m.G0 = v8 + int32(32)
	goto L1
L3:
	;
	v15 = int32(0)
	v17 = F___memset(m, v8, v15, int32(32))
	mBase = m.M
	v18 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	if v18 == v15 {
		goto L7
	} else {
		goto L8
	}
L4:
	;
	v14 = F___strchrnul(m, l0, v10)
	mBase = m.M
	v70 = v14
	goto L2
L5:
	;
	v13 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+1)))
	if v13 != 0 {
		goto L3
	} else {
		goto L6
	}
L6:
	;
	goto L4
L7:
	;
	v44 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v44 == int32(0) {
		v70 = l0
		goto L2
	} else {
		goto L12
	}
L8:
	;
	v22 = l1
	v24 = v18
	goto L9
L9:
	;
	v30 = v8 + int32(base.Ui32(v24)>>(uint(int32(3))%32))&int32(28)
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v30)))
	v32 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v30))) = v31 | v32<<(uint(v24)%32)
	v36 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22)+1)))
	if v36 != 0 {
		v22 = v22 + v32
		v24 = v36
		goto L9
	} else {
		goto L11
	}
L10:
	;
	goto L7
L11:
	;
	goto L10
L12:
	;
	v48 = l0
	v50 = v44
	goto L13
L13:
	;
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v8+int32(base.Ui32(v50)>>(uint(int32(3))%32))&int32(28))))
	if int32(base.Ui32(v57)>>(uint(v50)%32))&int32(1) == int32(0) {
		goto L15
	} else {
		goto L16
	}
L14:
	;
	v70 = v65
	goto L2
L15:
	;
	v63 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v48)+1)))
	v65 = v48 + int32(1)
	if v63 != 0 {
		v48 = v65
		v50 = v63
		goto L13
	} else {
		goto L17
	}
L16:
	;
	v70 = v48
	goto L2
L17:
	;
	goto L14
L18:
	;
	v78 = v75
	goto L20
L19:
	;
	v78 = int32(0)
	goto L20
L20:
	;
	return v78
}
func F_strtod(m *base.Module, l0 int32, l1 int32) float64 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v13 int32
	_ = v13
	var v14 int64
	_ = v14
	var v17 int64
	_ = v17
	var v18 float64
	_ = v18
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	F_strtox_1(m, v7, l0, l1, int32(1))
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return float64(0)
	} else {
		v14 = *(*int64)(unsafe.Add(mBase, uint32(v7)))
		v17 = *(*int64)(unsafe.Add(mBase, uint32(v7+int32(8))))
		v18 = F___trunctfdf2(m, v14, v17)
		mBase = m.M
		m.G0 = v7 + int32(16)
		return v18
	}
}
func F_strtoll(m *base.Module, l0 int32, l1 int32, l2 int32) int64 {
	var v5 int64
	_ = v5
	v5 = F_strtox_2(m, l0, l1, l2, int64(-9223372036854775807-1))
	return v5
}
func F_strtox_1(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v17 int32
	_ = v17
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	var v41 int32
	_ = v41
	var v44 int64
	_ = v44
	var v45 int64
	_ = v45
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	v8 = m.G0
	v10 = v8 - int32(160)
	m.G0 = v10
	*(*int32)(unsafe.Add(mBase, uint32(v10)+60)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v10)+20)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v10)+24)) = int32(-1)
	v17 = v10 + int32(16)
	*(*int64)(unsafe.Add(mBase, uint32(v17)+112)) = int64(0)
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v17)+44))
	v23 = *(*int32)(unsafe.Add(mBase, uint32(v17)+4))
	*(*int64)(unsafe.Add(mBase, uint32(v17)+120)) = base.I64_extend_i32_s(v22 - v23)
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v17)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v17)+104)) = v27
	F___floatscan(m, v10, v10+int32(16), l3, int32(1))
	mBase = m.M
	v41 = m.ExcPending
	if v41 != 0 {
		return
	} else {
		v44 = *(*int64)(unsafe.Add(mBase, uint32(v10+int32(8))))
		v45 = *(*int64)(unsafe.Add(mBase, uint32(v10)))
		if l2 == int32(0) {
		} else {
			v48 = *(*int32)(unsafe.Add(mBase, uint32(v10)+20))
			v49 = *(*int32)(unsafe.Add(mBase, uint32(v10)+60))
			v52 = *(*int32)(unsafe.Add(mBase, uint32(v10)+136))
			*(*int32)(unsafe.Add(mBase, uint32(l2))) = l1 + (v48 - v49) + v52
		}
		*(*int64)(unsafe.Add(mBase, uint32(l0)+8)) = v44
		*(*int64)(unsafe.Add(mBase, uint32(l0))) = v45
		m.G0 = v10 + int32(160)
		return
	}
}
func F_strtox_2(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int64) int64 {
	mBase := m.M
	_ = mBase
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v43 int32
	_ = v43
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v64 int32
	_ = v64
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v82 int32
	_ = v82
	var v86 int32
	_ = v86
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v108 int64
	_ = v108
	var v113 int32
	_ = v113
	var v118 int32
	_ = v118
	var v120 int32
	_ = v120
	var v123 int64
	_ = v123
	var v126 int32
	_ = v126
	var v128 int32
	_ = v128
	var v149 int32
	_ = v149
	var v153 int64
	_ = v153
	var v159 int64
	_ = v159
	var v160 int64
	_ = v160
	var v162 int64
	_ = v162
	var v165 int64
	_ = v165
	var v166 int64
	_ = v166
	var v168 int64
	_ = v168
	var v169 int64
	_ = v169
	var v173 int64
	_ = v173
	var v180 int64
	_ = v180
	var v191 int32
	_ = v191
	var v192 int64
	_ = v192
	var v195 int64
	_ = v195
	var v198 int64
	_ = v198
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v206 int64
	_ = v206
	var v214 int32
	_ = v214
	var v223 int64
	_ = v223
	var v226 int32
	_ = v226
	var v230 int32
	_ = v230
	var v231 int64
	_ = v231
	var v232 int64
	_ = v232
	var v243 int32
	_ = v243
	var v245 int64
	_ = v245
	var v246 int64
	_ = v246
	var v252 int64
	_ = v252
	v16 = m.G0
	v18 = v16 - int32(16)
	m.G0 = v18
	if int32(36) < l2 {
		goto L4
	} else {
		goto L5
	}
L1:
	;
	m.G0 = v18 + int32(16)
	return v252
L2:
	;
	v82 = int32(16)
	if l2|v82 != v82 {
		goto L19
	} else {
		goto L20
	}
L3:
	;
	v34 = v23
	v35 = l0
	goto L9
L4:
	;
	goto L7
L5:
	;
	v22 = int32(0)
	v23 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v23 != 0 {
		goto L3
	} else {
		goto L6
	}
L6:
	;
	v72 = v22
	v74 = l0
	goto L2
L7:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_strtox_2[0])) = int32(28)
	v252 = int64(0)
	goto L1
L8:
	;
	v57 = v34 & int32(255)
	switch v57 + int32(-43) {
	case 0, 2:
		goto L14
	default:
		v72 = v22
		v74 = v35
		goto L2
	}
L9:
	;
	v43 = base.I32_extend8_s(v34)
	goto L11
L10:
	;
	v72 = v22
	v74 = v55
	goto L2
L11:
	;
	if base.B2i32(v43 == int32(32))|base.B2i32(base.Ui32(v43+int32(-9)) < base.Ui32(int32(5))) == int32(0) {
		goto L8
	} else {
		goto L12
	}
L12:
	;
	v53 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v35)+1)))
	v55 = v35 + int32(1)
	if v53 != 0 {
		v34 = v53
		v35 = v55
		goto L9
	} else {
		goto L13
	}
L13:
	;
	goto L10
L14:
	;
	if v57 == int32(45) {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	v64 = int32(-1)
	goto L17
L16:
	;
	v64 = int32(0)
	goto L17
L17:
	;
	v72 = v64
	v74 = v35 + int32(1)
	goto L2
L18:
	;
	v108 = base.I64_extend_i32_u(v107)
	v113 = int32(0)
	v118 = v105
	v120 = v106
	v123 = int64(0)
	goto L31
L19:
	;
	if l2 != 0 {
		goto L27
	} else {
		goto L28
	}
L20:
	;
	v86 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v74))))
	if v86 != int32(48) {
		goto L19
	} else {
		goto L21
	}
L21:
	;
	v89 = int32(1)
	v90 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v74)+1)))
	if v90&int32(223) != int32(88) {
		goto L22
	} else {
		goto L23
	}
L22:
	;
	if l2 != 0 {
		goto L24
	} else {
		goto L25
	}
L23:
	;
	v105 = v74 + int32(2)
	v106 = v89
	v107 = int32(16)
	goto L18
L24:
	;
	v101 = l2
	goto L26
L25:
	;
	v101 = int32(8)
	goto L26
L26:
	;
	v105 = v74 + int32(1)
	v106 = v89
	v107 = v101
	goto L18
L27:
	;
	v103 = l2
	goto L29
L28:
	;
	v103 = int32(10)
	goto L29
L29:
	;
	v105 = v74
	v106 = int32(0)
	v107 = v103
	goto L18
L30:
	;
	if l1 == int32(0) {
		goto L43
	} else {
		goto L44
	}
L31:
	;
	v126 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v118))))
	v128 = v126 + int32(-48)
	if base.Ui32(v128&int32(255)) < base.Ui32(int32(10)) {
		v149 = v128
		goto L33
	} else {
		goto L34
	}
L33:
	;
	if v107 <= v149&int32(255) {
		goto L30
	} else {
		goto L38
	}
L34:
	;
	if base.Ui32(int32(25)) < base.Ui32((v126+int32(-97))&int32(255)) {
		goto L35
	} else {
		goto L36
	}
L35:
	;
	if base.Ui32(int32(25)) < base.Ui32((v126+int32(-65))&int32(255)) {
		goto L30
	} else {
		goto L37
	}
L36:
	;
	v149 = v126 + int32(-87)
	goto L33
L37:
	;
	v149 = v126 + int32(-55)
	goto L33
L38:
	;
	v153 = int64(0)
	v159 = int64(32)
	v160 = int64(base.Ui64(v123) >> (uint(v159) % 64))
	v162 = int64(base.Ui64(v108) >> (uint(v159) % 64))
	v165 = int64(4294967295)
	v166 = v123 & v165
	v168 = v108 & v165
	v169 = v166 * v168
	v173 = int64(base.Ui64(v169)>>(uint(v159)%64)) + v166*v162
	v180 = v173&v165 + v160*v168
	*(*int64)(unsafe.Add(mBase, uint32(v18)+8)) = v153*v108 + v153*v123 + v160*v162 + int64(base.Ui64(v173)>>(uint(v159)%64)) + int64(base.Ui64(v180)>>(uint(v159)%64))
	*(*int64)(unsafe.Add(mBase, uint32(v18))) = v180<<(uint(v159)%64) | v169&v165
	goto L39
L39:
	;
	v191 = int32(1)
	v192 = *(*int64)(unsafe.Add(mBase, uint32(v18)+8))
	if v192 != int64(0) {
		v204 = v191
		v205 = v120
		v206 = v123
		goto L40
	} else {
		goto L41
	}
L40:
	;
	v113 = v204
	v118 = v118 + int32(1)
	v120 = v205
	v123 = v206
	goto L31
L41:
	;
	v195 = v123 * v108
	v198 = base.I64_extend_i32_u(v149) & int64(255)
	if base.Ui64(v198^int64(-1)) < base.Ui64(v195) {
		v204 = v191
		v205 = v120
		v206 = v123
		goto L40
	} else {
		goto L42
	}
L42:
	;
	v204 = v113
	v205 = int32(1)
	v206 = v195 + v198
	goto L40
L43:
	;
	if v113 == int32(0) {
		goto L50
	} else {
		goto L51
	}
L44:
	;
	if v120 != 0 {
		goto L45
	} else {
		goto L46
	}
L45:
	;
	v214 = v118
	goto L47
L46:
	;
	v214 = l0
	goto L47
L47:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v214
	goto L43
L48:
	;
	v246 = base.I64_extend_i32_s(v243)
	v252 = v245 ^ v246 - v246
	goto L1
L49:
	;
	if base.I32_wrap_i64(v231) != 0 {
		goto L57
	} else {
		goto L58
	}
L50:
	;
	if base.Ui64(v123) < base.Ui64(l3) {
		v243 = v72
		v245 = v123
		goto L48
	} else {
		goto L56
	}
L51:
	;
	goto L52
L52:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_strtox_2[0])) = int32(68)
	v223 = l3 & int64(1)
	if v223 == int64(0) {
		goto L53
	} else {
		goto L54
	}
L53:
	;
	v226 = v72
	goto L55
L54:
	;
	v226 = int32(0)
	goto L55
L55:
	;
	v230 = v226
	v231 = v223
	v232 = l3
	goto L49
L56:
	;
	v230 = v72
	v231 = l3 & int64(1)
	v232 = v123
	goto L49
L57:
	;
	if base.Ui64(v232) <= base.Ui64(l3) {
		v243 = v230
		v245 = v232
		goto L48
	} else {
		goto L61
	}
L58:
	;
	if v230 != 0 {
		goto L57
	} else {
		goto L59
	}
L59:
	;
	goto L60
L60:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_strtox_2[0])) = int32(68)
	v252 = l3 + int64(-1)
	goto L1
L61:
	;
	goto L62
L62:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_strtox_2[0])) = int32(68)
	v252 = l3
	goto L1
}
func F_subexpr(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
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
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v22 int32
	_ = v22
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
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
	var v49 int32
	_ = v49
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v100 int32
	_ = v100
	var v102 int32
	_ = v102
	var v104 int32
	_ = v104
	var v108 int32
	_ = v108
	var v115 float64
	_ = v115
	var v121 int32
	_ = v121
	var v123 int32
	_ = v123
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v148 int32
	_ = v148
	var v156 int32
	_ = v156
	var v158 int32
	_ = v158
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v170 int32
	_ = v170
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v188 int32
	_ = v188
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v194 int32
	_ = v194
	v8 = m.G0
	v10 = v8 - int32(32)
	m.G0 = v10
	v12 = int32(1)
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v14 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v13)+52)))
	v16 = v14 + v12
	*(*uint16)(unsafe.Add(mBase, uint32(v13)+52)) = uint16(v16)
	if base.Ui32(v16&int32(65535)) < base.Ui32(int32(201)) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	switch v30 + int32(-263) {
	case 0:
		goto L13
	case 1, 3, 4, 5, 8, 9, 10, 11, 13, 14, 15, 17, 18, 19, 20, 22:
		goto L9
	case 2:
		goto L11
	case 6:
		goto L15
	case 7:
		v37 = v12
		goto L17
	case 12:
		goto L14
	case 16:
		goto L12
	case 21:
		goto L8
	case 23:
		goto L16
	default:
		goto L18
	}
L2:
	;
	v22 = m.G3
	F_luaX_lexerror(m, l0, v22+int32(_a_F_subexpr_0), int32(0))
	mBase = m.M
	v29 = m.ExcPending
	if v29 != 0 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	return int32(0)
L4:
	;
	goto L1
L5:
	;
	v128 = int32(15)
	v129 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	switch v129 + int32(-37) {
	case 0:
		goto L47
	case 1, 2, 3, 4, 7, 9, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 24, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46, 47, 48, 49, 50, 51, 52, 53, 54, 55, 56:
		v188 = v128
		goto L35
	case 5:
		goto L49
	case 6:
		v148 = int32(0)
		goto L36
	case 8:
		goto L50
	case 10:
		goto L48
	case 23:
		goto L42
	case 25:
		goto L40
	case 57:
		goto L46
	default:
		goto L51
	}
L6:
	;
	F_constructor(m, l0, l1)
	mBase = m.M
	v123 = m.ExcPending
	if v123 != 0 {
		goto L3
	} else {
		goto L34
	}
L7:
	;
	F_luaX_next(m, l0)
	mBase = m.M
	v121 = m.ExcPending
	if v121 != 0 {
		goto L3
	} else {
		goto L33
	}
L8:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+8)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = int32(5)
	*(*int64)(unsafe.Add(mBase, uint32(l1)+16)) = int64(-1)
	v115 = *(*float64)(unsafe.Add(mBase, uint32(l0)+24))
	*(*float64)(unsafe.Add(mBase, uint32(l1)+8)) = v115
	goto L7
L9:
	;
	F_primaryexp(m, l0, l1)
	mBase = m.M
	v108 = m.ExcPending
	if v108 != 0 {
		goto L3
	} else {
		goto L32
	}
L10:
	;
	if v30 == int32(123) {
		goto L6
	} else {
		goto L31
	}
L11:
	;
	F_luaX_next(m, l0)
	mBase = m.M
	v100 = m.ExcPending
	if v100 != 0 {
		goto L3
	} else {
		goto L29
	}
L12:
	;
	v73 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v74 = *(*int32)(unsafe.Add(mBase, uint32(v73)))
	v75 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v74)+74)))
	if v75 != 0 {
		v83 = v74
		v84 = v75
		goto L25
	} else {
		goto L26
	}
L13:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+8)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = int32(3)
	*(*int64)(unsafe.Add(mBase, uint32(l1)+16)) = int64(-1)
	goto L7
L14:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+8)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = int32(2)
	*(*int64)(unsafe.Add(mBase, uint32(l1)+16)) = int64(-1)
	goto L7
L15:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+8)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = int32(1)
	*(*int64)(unsafe.Add(mBase, uint32(l1)+16)) = int64(-1)
	goto L7
L16:
	;
	v46 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v47 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v48 = F_luaK_stringK(m, v46, v47)
	mBase = m.M
	v49 = m.ExcPending
	if v49 != 0 {
		goto L3
	} else {
		goto L24
	}
L17:
	;
	F_luaX_next(m, l0)
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L3
	} else {
		goto L21
	}
L18:
	;
	switch v30 + int32(-35) {
	case 0:
		goto L19
	case 1, 2, 3, 4, 5, 6, 7, 8, 9:
		goto L9
	case 10:
		goto L20
	default:
		goto L10
	}
L19:
	;
	v37 = int32(2)
	goto L17
L20:
	;
	v37 = int32(0)
	goto L17
L21:
	;
	v41 = F_subexpr(m, l0, l1, int32(8))
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L3
	} else {
		goto L22
	}
L22:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	F_luaK_prefix(m, v43, v37, l1)
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L3
	} else {
		goto L23
	}
L23:
	;
	goto L5
L24:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+8)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = int32(4)
	*(*int64)(unsafe.Add(mBase, uint32(l1)+16)) = int64(-1)
	goto L7
L25:
	;
	v86 = v84 & int32(251)
	*(*uint8)(unsafe.Add(mBase, uint32(v83)+74)) = uint8(v86)
	v89 = int32(0)
	v92 = F_luaK_codeABC(m, v73, int32(37), v89, int32(1), v89)
	mBase = m.M
	v93 = m.ExcPending
	if v93 != 0 {
		goto L3
	} else {
		goto L28
	}
L26:
	;
	v76 = m.G3
	F_luaX_syntaxerror(m, l0, v76+int32(_a_F_subexpr_1))
	mBase = m.M
	v80 = m.ExcPending
	if v80 != 0 {
		goto L3
	} else {
		goto L27
	}
L27:
	;
	v81 = *(*int32)(unsafe.Add(mBase, uint32(v73)))
	v82 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v81)+74)))
	v83 = v81
	v84 = v82
	goto L25
L28:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+8)) = v92
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = int32(14)
	*(*int64)(unsafe.Add(mBase, uint32(l1)+16)) = int64(-1)
	goto L7
L29:
	;
	v102 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	F_body(m, l0, l1, int32(0), v102)
	mBase = m.M
	v104 = m.ExcPending
	if v104 != 0 {
		goto L3
	} else {
		goto L30
	}
L30:
	;
	goto L5
L31:
	;
	goto L9
L32:
	;
	goto L5
L33:
	;
	goto L5
L34:
	;
	goto L5
L35:
	;
	v191 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v192 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v191)+52)))
	v194 = v192 + int32(-1)
	*(*uint16)(unsafe.Add(mBase, uint32(v191)+52)) = uint16(v194)
	m.G0 = v10 + int32(32)
	return v188
L36:
	;
	v156 = v148
	goto L52
L37:
	;
	v148 = int32(14)
	goto L36
L38:
	;
	v148 = int32(13)
	goto L36
L39:
	;
	v148 = int32(12)
	goto L36
L40:
	;
	v148 = int32(11)
	goto L36
L41:
	;
	v148 = int32(10)
	goto L36
L42:
	;
	v148 = int32(9)
	goto L36
L43:
	;
	v148 = int32(8)
	goto L36
L44:
	;
	v148 = int32(7)
	goto L36
L45:
	;
	v148 = int32(6)
	goto L36
L46:
	;
	v148 = int32(5)
	goto L36
L47:
	;
	v148 = int32(4)
	goto L36
L48:
	;
	v148 = int32(3)
	goto L36
L49:
	;
	v148 = int32(2)
	goto L36
L50:
	;
	v148 = int32(1)
	goto L36
L51:
	;
	switch v129 + int32(-257) {
	case 0:
		goto L38
	default:
		v188 = v128
		goto L35
	case 14:
		goto L37
	case 21:
		goto L45
	case 23:
		goto L43
	case 24:
		goto L39
	case 25:
		goto L41
	case 26:
		goto L44
	}
L52:
	;
	v158 = m.G3
	v163 = v158 + int32(_a_F_subexpr_2) + v156<<(uint(int32(1))%32)
	v164 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v163))))
	if base.Ui32(l2&int32(255)) < base.Ui32(v164) {
		goto L54
	} else {
		goto L55
	}
L53:
	;
	v188 = v181
	goto L35
L54:
	;
	F_luaX_next(m, l0)
	mBase = m.M
	v167 = m.ExcPending
	if v167 != 0 {
		goto L3
	} else {
		goto L56
	}
L55:
	;
	v188 = v156
	goto L35
L56:
	;
	v168 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	F_luaK_infix(m, v168, v156, l1)
	mBase = m.M
	v170 = m.ExcPending
	if v170 != 0 {
		goto L3
	} else {
		goto L57
	}
L57:
	;
	v173 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v163)+1)))
	v174 = F_subexpr(m, l0, v10+int32(8), v173)
	mBase = m.M
	v175 = m.ExcPending
	if v175 != 0 {
		goto L3
	} else {
		goto L58
	}
L58:
	;
	v176 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	F_luaK_posfix(m, v176, v156, l1, v10+int32(8))
	mBase = m.M
	v180 = m.ExcPending
	if v180 != 0 {
		goto L3
	} else {
		goto L59
	}
L59:
	;
	v181 = int32(15)
	if v174 != v181 {
		v156 = v174
		goto L52
	} else {
		goto L60
	}
L60:
	;
	goto L53
}
func F_sunionDiffGenericCommand(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) {
	mBase := m.M
	_ = mBase
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v96 int32
	_ = v96
	var v99 int32
	_ = v99
	var v103 int32
	_ = v103
	var v106 int32
	_ = v106
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v114 int32
	_ = v114
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v120 int32
	_ = v120
	var v132 int32
	_ = v132
	var v134 int32
	_ = v134
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v174 int64
	_ = v174
	var v187 int32
	_ = v187
	var v189 int64
	_ = v189
	var v190 int64
	_ = v190
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v210 int64
	_ = v210
	var v211 int64
	_ = v211
	var v213 int32
	_ = v213
	var v218 int64
	_ = v218
	var v220 int32
	_ = v220
	var v224 int32
	_ = v224
	var v231 int32
	_ = v231
	var v238 int32
	_ = v238
	var v252 int32
	_ = v252
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v255 int32
	_ = v255
	var v256 int32
	_ = v256
	var v258 int32
	_ = v258
	var v263 int32
	_ = v263
	var v273 int32
	_ = v273
	var v283 int32
	_ = v283
	var v286 int32
	_ = v286
	var v287 int32
	_ = v287
	var v294 int32
	_ = v294
	var v295 int32
	_ = v295
	var v309 int32
	_ = v309
	var v310 int32
	_ = v310
	var v316 int32
	_ = v316
	var v317 int32
	_ = v317
	var v318 int64
	_ = v318
	var v321 int32
	_ = v321
	var v322 int32
	_ = v322
	var v323 int32
	_ = v323
	var v330 int32
	_ = v330
	var v331 int32
	_ = v331
	var v345 int32
	_ = v345
	var v352 int32
	_ = v352
	var v355 int32
	_ = v355
	var v357 int32
	_ = v357
	var v359 int32
	_ = v359
	var v371 int32
	_ = v371
	var v379 int32
	_ = v379
	var v382 int32
	_ = v382
	var v387 int32
	_ = v387
	var v392 int32
	_ = v392
	var v393 int32
	_ = v393
	var v394 int32
	_ = v394
	var v401 int32
	_ = v401
	var v402 int32
	_ = v402
	var v418 int32
	_ = v418
	var v423 int32
	_ = v423
	var v429 int32
	_ = v429
	var v430 int32
	_ = v430
	var v443 int32
	_ = v443
	var v452 int32
	_ = v452
	var v455 int32
	_ = v455
	var v457 int64
	_ = v457
	var v458 int32
	_ = v458
	var v459 int32
	_ = v459
	var v461 int32
	_ = v461
	var v475 int32
	_ = v475
	var v500 int32
	_ = v500
	var v501 int32
	_ = v501
	var v502 int64
	_ = v502
	var v505 int32
	_ = v505
	var v506 int32
	_ = v506
	var v519 int32
	_ = v519
	var v532 int32
	_ = v532
	var v533 int32
	_ = v533
	var v547 int32
	_ = v547
	var v554 int32
	_ = v554
	var v557 int32
	_ = v557
	var v559 int32
	_ = v559
	var v561 int32
	_ = v561
	var v569 int32
	_ = v569
	var v581 int32
	_ = v581
	var v582 int32
	_ = v582
	var v592 int32
	_ = v592
	var v595 int32
	_ = v595
	var v596 int32
	_ = v596
	var v603 int32
	_ = v603
	var v604 int32
	_ = v604
	var v618 int32
	_ = v618
	var v619 int32
	_ = v619
	var v626 int32
	_ = v626
	var v627 int64
	_ = v627
	var v628 int32
	_ = v628
	var v629 int32
	_ = v629
	var v630 int32
	_ = v630
	var v631 int32
	_ = v631
	var v633 int32
	_ = v633
	var v634 int32
	_ = v634
	var v636 int32
	_ = v636
	var v643 int32
	_ = v643
	var v644 int32
	_ = v644
	var v658 int32
	_ = v658
	var v665 int32
	_ = v665
	var v668 int32
	_ = v668
	var v670 int32
	_ = v670
	var v672 int32
	_ = v672
	var v685 int32
	_ = v685
	var v693 int32
	_ = v693
	var v706 int32
	_ = v706
	var v714 int32
	_ = v714
	var v715 int32
	_ = v715
	var v716 int32
	_ = v716
	var v723 int32
	_ = v723
	var v724 int32
	_ = v724
	var v745 int32
	_ = v745
	var v748 int32
	_ = v748
	var v750 int32
	_ = v750
	var v751 int64
	_ = v751
	var v753 int32
	_ = v753
	var v760 int32
	_ = v760
	var v761 int32
	_ = v761
	var v782 int32
	_ = v782
	var v785 int32
	_ = v785
	var v787 int32
	_ = v787
	var v789 int32
	_ = v789
	var v791 int32
	_ = v791
	var v797 int32
	_ = v797
	var v798 int32
	_ = v798
	var v799 int32
	_ = v799
	var v800 int32
	_ = v800
	var v807 int32
	_ = v807
	var v811 int32
	_ = v811
	var v812 int32
	_ = v812
	var v813 int32
	_ = v813
	var v815 int32
	_ = v815
	var v816 int32
	_ = v816
	var v818 int64
	_ = v818
	var v822 int32
	_ = v822
	var v823 int32
	_ = v823
	var v824 int32
	_ = v824
	var v827 int32
	_ = v827
	var v828 int32
	_ = v828
	var v829 int32
	_ = v829
	var v832 int32
	_ = v832
	var v834 int64
	_ = v834
	var v838 int32
	_ = v838
	var v840 int32
	_ = v840
	var v843 int32
	_ = v843
	var v844 int32
	_ = v844
	var v846 int32
	_ = v846
	var v848 int32
	_ = v848
	var v850 int32
	_ = v850
	var v870 int32
	_ = v870
	var v890 int32
	_ = v890
	v19 = m.G0
	v21 = v19 - int32(32)
	m.G0 = v21
	v26 = F_valkey_malloc(m, l2<<(uint(int32(2))%32))
	mBase = m.M
	v27 = m.ExcPending
	if v27 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	if int32(1) <= l2 {
		goto L5
	} else {
		goto L6
	}
L3:
	;
	F_valkey_free(m, v26)
	mBase = m.M
	v890 = m.ExcPending
	if v890 != 0 {
		goto L1
	} else {
		goto L185
	}
L4:
	;
	if l4 != int32(1) {
		v238 = int32(1)
		goto L42
	} else {
		goto L43
	}
L5:
	;
	v32 = int32(6)
	v33 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v34 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v35 = F_lookupKeyRead(m, v33, v34)
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
		goto L1
	} else {
		goto L8
	}
L6:
	;
	v152 = int32(0)
	v153 = int32(1)
	goto L4
L7:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26))) = v35
	v59 = int32(1)
	v60 = int32(0)
	if l2 == v59 {
		v132 = v60
		v134 = v56
		goto L19
	} else {
		goto L20
	}
L8:
	;
	if v35 == int32(0) {
		v56 = v32
		goto L7
	} else {
		goto L9
	}
L9:
	;
	v40 = F_checkType(m, l0, v35, int32(2))
	mBase = m.M
	v41 = m.ExcPending
	if v41 != 0 {
		goto L1
	} else {
		goto L10
	}
L10:
	;
	if v40 != 0 {
		goto L3
	} else {
		goto L11
	}
L11:
	;
	if l3 != 0 {
		v56 = v32
		goto L7
	} else {
		goto L12
	}
L12:
	;
	v42 = int32(2)
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v35)))
	v49 = int32(base.Ui32(v45)>>(uint(int32(4))%32)) & int32(15)
	if v49 == v42 {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	v52 = v42
	goto L15
L14:
	;
	v52 = int32(6)
	goto L15
L15:
	;
	if v49 == int32(11) {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	v55 = v42
	goto L18
L17:
	;
	v55 = v52
	goto L18
L18:
	;
	v56 = v55
	goto L7
L19:
	;
	v152 = v132
	v153 = base.B2i32(v134 == int32(6))
	goto L4
L20:
	;
	v71 = v60
	v73 = v56
	v75 = v59
	goto L21
L21:
	;
	v81 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v83 = v75 << (uint(int32(2)) % 32)
	v85 = *(*int32)(unsafe.Add(mBase, uint32(l1+v83)))
	v86 = F_lookupKeyRead(m, v81, v85)
	mBase = m.M
	v87 = m.ExcPending
	if v87 != 0 {
		goto L1
	} else {
		goto L25
	}
L22:
	;
	v132 = base.B2i32(v117 != int32(0))
	v134 = v118
	goto L19
L23:
	;
	v120 = v75 + int32(1)
	if v120 != l2 {
		v71 = v117
		v73 = v118
		v75 = v120
		goto L21
	} else {
		goto L41
	}
L24:
	;
	v92 = F_checkType(m, l0, v86, int32(2))
	mBase = m.M
	v93 = m.ExcPending
	if v93 != 0 {
		goto L1
	} else {
		goto L27
	}
L25:
	;
	if v86 != 0 {
		goto L24
	} else {
		goto L26
	}
L26:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26+v83))) = int32(0)
	v117 = v71
	v118 = v73
	goto L23
L27:
	;
	if v92 != 0 {
		goto L3
	} else {
		goto L28
	}
L28:
	;
	if l3 != 0 {
		v110 = v73
		goto L29
	} else {
		goto L30
	}
L29:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26+v83))) = v86
	v114 = *(*int32)(unsafe.Add(mBase, uint32(v26)))
	if v114 == v86 {
		goto L38
	} else {
		goto L39
	}
L30:
	;
	if v73 != int32(6) {
		v110 = v73
		goto L29
	} else {
		goto L31
	}
L31:
	;
	v96 = int32(2)
	v99 = *(*int32)(unsafe.Add(mBase, uint32(v86)))
	v103 = int32(base.Ui32(v99)>>(uint(int32(4))%32)) & int32(15)
	if v103 == v96 {
		goto L32
	} else {
		goto L33
	}
L32:
	;
	v106 = v96
	goto L34
L33:
	;
	v106 = int32(6)
	goto L34
L34:
	;
	if v103 == int32(11) {
		goto L35
	} else {
		goto L36
	}
L35:
	;
	v109 = v96
	goto L37
L36:
	;
	v109 = v106
	goto L37
L37:
	;
	v110 = v109
	goto L29
L38:
	;
	v116 = int32(1)
	goto L40
L39:
	;
	v116 = v71
	goto L40
L40:
	;
	v117 = v116
	v118 = v110
	goto L23
L41:
	;
	goto L22
L42:
	;
	if v153 == int32(0) {
		goto L60
	} else {
		goto L61
	}
L43:
	;
	v164 = int32(1)
	v165 = *(*int32)(unsafe.Add(mBase, uint32(v26)))
	if (base.B2i32(v165 == int32(0))|v152)&v164 != 0 {
		v238 = v164
		goto L42
	} else {
		goto L44
	}
L44:
	;
	if l2 < int32(1) {
		v238 = v164
		goto L42
	} else {
		goto L45
	}
L45:
	;
	v174 = int64(0)
	v187 = int32(0)
	v189 = v174
	v190 = v174
	goto L46
L46:
	;
	v196 = v26 + v187<<(uint(int32(2))%32)
	v197 = *(*int32)(unsafe.Add(mBase, uint32(v196)))
	if v197 == int32(0) {
		v210 = v189
		v211 = v190
		goto L48
	} else {
		goto L49
	}
L47:
	;
	v218 = base.I64_div_s(v211, int64(2))
	if v218 <= v210 {
		goto L53
	} else {
		goto L54
	}
L48:
	;
	v213 = v187 + int32(1)
	if v213 != l2 {
		v187 = v213
		v189 = v210
		v190 = v211
		goto L46
	} else {
		goto L52
	}
L49:
	;
	v200 = *(*int32)(unsafe.Add(mBase, uint32(v26)))
	v201 = F_setTypeSize(m, v200)
	mBase = m.M
	v202 = m.ExcPending
	if v202 != 0 {
		goto L1
	} else {
		goto L50
	}
L50:
	;
	v205 = *(*int32)(unsafe.Add(mBase, uint32(v196)))
	v206 = F_setTypeSize(m, v205)
	mBase = m.M
	v207 = m.ExcPending
	if v207 != 0 {
		goto L1
	} else {
		goto L51
	}
L51:
	;
	v210 = v189 + base.I64_extend_i32_u(v206)
	v211 = v190 + base.I64_extend_i32_u(v201)
	goto L48
L52:
	;
	goto L47
L53:
	;
	v220 = int32(1)
	goto L55
L54:
	;
	v220 = int32(2)
	goto L55
L55:
	;
	if l2 < int32(2) {
		v238 = v220
		goto L42
	} else {
		goto L56
	}
L56:
	;
	if v210 < v218 {
		v238 = v220
		goto L42
	} else {
		goto L57
	}
L57:
	;
	v224 = int32(4)
	F_qsort(m, v26+v224, l2+int32(-1), v224, int32(1090))
	mBase = m.M
	v231 = m.ExcPending
	if v231 != 0 {
		goto L1
	} else {
		goto L58
	}
L58:
	;
	v238 = v220
	goto L42
L59:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+28)) = v256
	if l4 != 0 {
		goto L65
	} else {
		goto L66
	}
L60:
	;
	v254 = F_createSetObject(m)
	mBase = m.M
	v255 = m.ExcPending
	if v255 != 0 {
		goto L1
	} else {
		goto L63
	}
L61:
	;
	v252 = F_createIntsetObject(m)
	mBase = m.M
	v253 = m.ExcPending
	if v253 != 0 {
		goto L1
	} else {
		goto L62
	}
L62:
	;
	v256 = v252
	goto L59
L63:
	;
	v256 = v254
	goto L59
L64:
	;
	if l3 != 0 {
		goto L146
	} else {
		goto L147
	}
L65:
	;
	v382 = int32(1)
	if (base.B2i32(l4 != v382)|v152)&v382 != 0 {
		v706 = int32(0)
		goto L64
	} else {
		goto L86
	}
L66:
	;
	v258 = int32(0)
	if l2 < int32(1) {
		v706 = v258
		goto L64
	} else {
		goto L67
	}
L67:
	;
	v263 = int32(0)
	v273 = v258
	goto L68
L68:
	;
	v283 = *(*int32)(unsafe.Add(mBase, uint32(v26+v263<<(uint(int32(2))%32))))
	if v283 == int32(0) {
		v371 = v273
		goto L70
	} else {
		goto L71
	}
L70:
	;
	v379 = v263 + int32(1)
	if v379 != l2 {
		v263 = v379
		v273 = v371
		goto L68
	} else {
		goto L85
	}
L71:
	;
	v286 = F_setTypeInitIterator(m, v283)
	mBase = m.M
	v287 = m.ExcPending
	if v287 != 0 {
		goto L1
	} else {
		goto L73
	}
L72:
	;
	v352 = *(*int32)(unsafe.Add(mBase, uint32(v286)+4))
	if v352 != int32(2) {
		goto L81
	} else {
		goto L82
	}
L73:
	;
	v294 = F_setTypeNext(m, v286, v21+int32(24), v21+int32(20), v21+int32(8))
	mBase = m.M
	v295 = m.ExcPending
	if v295 != 0 {
		goto L1
	} else {
		goto L74
	}
L74:
	;
	if v294 == int32(-1) {
		v345 = v273
		goto L72
	} else {
		goto L75
	}
L75:
	;
	v309 = v273
	v310 = v294
	goto L76
L76:
	;
	v316 = *(*int32)(unsafe.Add(mBase, uint32(v21)+24))
	v317 = *(*int32)(unsafe.Add(mBase, uint32(v21)+20))
	v318 = *(*int64)(unsafe.Add(mBase, uint32(v21)+8))
	v321 = F_setTypeAddAux(m, v256, v316, v317, v318, base.B2i32(v310 == int32(2)))
	mBase = m.M
	v322 = m.ExcPending
	if v322 != 0 {
		goto L1
	} else {
		goto L78
	}
L77:
	;
	v345 = v323
	goto L72
L78:
	;
	v323 = v321 + v309
	v330 = F_setTypeNext(m, v286, v21+int32(24), v21+int32(20), v21+int32(8))
	mBase = m.M
	v331 = m.ExcPending
	if v331 != 0 {
		goto L1
	} else {
		goto L79
	}
L79:
	;
	if v330 != int32(-1) {
		v309 = v323
		v310 = v330
		goto L76
	} else {
		goto L80
	}
L80:
	;
	goto L77
L81:
	;
	F_valkey_free(m, v286)
	mBase = m.M
	v359 = m.ExcPending
	if v359 != 0 {
		goto L1
	} else {
		goto L84
	}
L82:
	;
	v355 = *(*int32)(unsafe.Add(mBase, uint32(v286)+12))
	F_hashtableReleaseIterator(m, v355)
	mBase = m.M
	v357 = m.ExcPending
	if v357 != 0 {
		goto L1
	} else {
		goto L83
	}
L83:
	;
	goto L81
L84:
	;
	v371 = v345
	goto L70
L85:
	;
	v706 = v371
	goto L64
L86:
	;
	v387 = *(*int32)(unsafe.Add(mBase, uint32(v26)))
	if v238 != int32(1) {
		goto L87
	} else {
		goto L88
	}
L87:
	;
	if v238 == int32(2) {
		goto L116
	} else {
		goto L117
	}
L88:
	;
	if v387 == int32(0) {
		goto L87
	} else {
		goto L89
	}
L89:
	;
	v392 = int32(0)
	v393 = F_setTypeInitIterator(m, v387)
	mBase = m.M
	v394 = m.ExcPending
	if v394 != 0 {
		goto L1
	} else {
		goto L91
	}
L90:
	;
	v554 = *(*int32)(unsafe.Add(mBase, uint32(v393)+4))
	if v554 != int32(2) {
		goto L112
	} else {
		goto L113
	}
L91:
	;
	v401 = F_setTypeNext(m, v393, v21+int32(24), v21+int32(20), v21+int32(8))
	mBase = m.M
	v402 = m.ExcPending
	if v402 != 0 {
		goto L1
	} else {
		goto L92
	}
L92:
	;
	if v401 == int32(-1) {
		v547 = v392
		goto L90
	} else {
		goto L93
	}
L93:
	;
	v418 = v392
	v423 = v401
	goto L94
L94:
	;
	if l2 < int32(2) {
		v475 = int32(1)
		goto L98
	} else {
		goto L99
	}
L95:
	;
	v547 = v519
	goto L90
L96:
	;
	v532 = F_setTypeNext(m, v393, v21+int32(24), v21+int32(20), v21+int32(8))
	mBase = m.M
	v533 = m.ExcPending
	if v533 != 0 {
		goto L1
	} else {
		goto L110
	}
L97:
	;
	v500 = *(*int32)(unsafe.Add(mBase, uint32(v21)+24))
	v501 = *(*int32)(unsafe.Add(mBase, uint32(v21)+20))
	v502 = *(*int64)(unsafe.Add(mBase, uint32(v21)+8))
	v505 = F_setTypeAddAux(m, v256, v500, v501, v502, base.B2i32(v423 == int32(2)))
	mBase = m.M
	v506 = m.ExcPending
	if v506 != 0 {
		goto L1
	} else {
		goto L109
	}
L98:
	;
	if v475 != l2 {
		v519 = v418
		goto L96
	} else {
		goto L108
	}
L99:
	;
	v429 = *(*int32)(unsafe.Add(mBase, uint32(v21)+20))
	v430 = *(*int32)(unsafe.Add(mBase, uint32(v21)+24))
	v443 = int32(1)
	goto L100
L100:
	;
	v452 = *(*int32)(unsafe.Add(mBase, uint32(v26+v443<<(uint(int32(2))%32))))
	if v452 == int32(0) {
		goto L102
	} else {
		goto L103
	}
L102:
	;
	v461 = v443 + int32(1)
	if v461 != l2 {
		v443 = v461
		goto L100
	} else {
		goto L107
	}
L103:
	;
	v455 = *(*int32)(unsafe.Add(mBase, uint32(v26)))
	if v452 == v455 {
		v475 = v443
		goto L98
	} else {
		goto L104
	}
L104:
	;
	v457 = *(*int64)(unsafe.Add(mBase, uint32(v21)+8))
	v458 = F_setTypeIsMemberAux(m, v452, v430, v429, v457, base.B2i32(v423 == int32(2)))
	mBase = m.M
	v459 = m.ExcPending
	if v459 != 0 {
		goto L1
	} else {
		goto L105
	}
L105:
	;
	if v458 != 0 {
		v475 = v443
		goto L98
	} else {
		goto L106
	}
L106:
	;
	goto L102
L107:
	;
	goto L97
L108:
	;
	goto L97
L109:
	;
	v519 = v505 + v418
	goto L96
L110:
	;
	if v532 != int32(-1) {
		v418 = v519
		v423 = v532
		goto L94
	} else {
		goto L111
	}
L111:
	;
	goto L95
L112:
	;
	F_valkey_free(m, v393)
	mBase = m.M
	v561 = m.ExcPending
	if v561 != 0 {
		goto L1
	} else {
		goto L115
	}
L113:
	;
	v557 = *(*int32)(unsafe.Add(mBase, uint32(v393)+12))
	F_hashtableReleaseIterator(m, v557)
	mBase = m.M
	v559 = m.ExcPending
	if v559 != 0 {
		goto L1
	} else {
		goto L114
	}
L114:
	;
	goto L112
L115:
	;
	v706 = v547
	goto L64
L116:
	;
	if int32(1) <= l2 {
		goto L118
	} else {
		goto L119
	}
L117:
	;
	v706 = int32(0)
	goto L64
L118:
	;
	if v387 != 0 {
		goto L120
	} else {
		goto L121
	}
L119:
	;
	v706 = int32(0)
	goto L64
L120:
	;
	v569 = int32(0)
	v581 = v569
	v582 = v569
	goto L122
L121:
	;
	v706 = int32(0)
	goto L64
L122:
	;
	v592 = *(*int32)(unsafe.Add(mBase, uint32(v26+v581<<(uint(int32(2))%32))))
	if v592 == int32(0) {
		v685 = v582
		goto L124
	} else {
		goto L125
	}
L123:
	;
	v706 = v685
	goto L64
L124:
	;
	v693 = v581 + int32(1)
	if v693 != l2 {
		v581 = v693
		v582 = v685
		goto L122
	} else {
		goto L144
	}
L125:
	;
	v595 = F_setTypeInitIterator(m, v592)
	mBase = m.M
	v596 = m.ExcPending
	if v596 != 0 {
		goto L1
	} else {
		goto L127
	}
L126:
	;
	v665 = *(*int32)(unsafe.Add(mBase, uint32(v595)+4))
	if v665 != int32(2) {
		goto L139
	} else {
		goto L140
	}
L127:
	;
	v603 = F_setTypeNext(m, v595, v21+int32(24), v21+int32(20), v21+int32(8))
	mBase = m.M
	v604 = m.ExcPending
	if v604 != 0 {
		goto L1
	} else {
		goto L128
	}
L128:
	;
	if v603 == int32(-1) {
		v658 = v582
		goto L126
	} else {
		goto L129
	}
L129:
	;
	v618 = v582
	v619 = v603
	goto L130
L130:
	;
	v626 = base.B2i32(v619 == int32(2))
	v627 = *(*int64)(unsafe.Add(mBase, uint32(v21)+8))
	v628 = *(*int32)(unsafe.Add(mBase, uint32(v21)+20))
	v629 = *(*int32)(unsafe.Add(mBase, uint32(v21)+24))
	if v581 != 0 {
		goto L133
	} else {
		goto L134
	}
L131:
	;
	v658 = v636
	goto L126
L132:
	;
	v643 = F_setTypeNext(m, v595, v21+int32(24), v21+int32(20), v21+int32(8))
	mBase = m.M
	v644 = m.ExcPending
	if v644 != 0 {
		goto L1
	} else {
		goto L137
	}
L133:
	;
	v633 = F_setTypeRemoveAux(m, v256, v629, v628, v627, v626)
	mBase = m.M
	v634 = m.ExcPending
	if v634 != 0 {
		goto L1
	} else {
		goto L136
	}
L134:
	;
	v630 = F_setTypeAddAux(m, v256, v629, v628, v627, v626)
	mBase = m.M
	v631 = m.ExcPending
	if v631 != 0 {
		goto L1
	} else {
		goto L135
	}
L135:
	;
	v636 = v630 + v618
	goto L132
L136:
	;
	v636 = v618 - v633
	goto L132
L137:
	;
	if v643 != int32(-1) {
		v618 = v636
		v619 = v643
		goto L130
	} else {
		goto L138
	}
L138:
	;
	goto L131
L139:
	;
	F_valkey_free(m, v595)
	mBase = m.M
	v672 = m.ExcPending
	if v672 != 0 {
		goto L1
	} else {
		goto L142
	}
L140:
	;
	v668 = *(*int32)(unsafe.Add(mBase, uint32(v595)+12))
	F_hashtableReleaseIterator(m, v668)
	mBase = m.M
	v670 = m.ExcPending
	if v670 != 0 {
		goto L1
	} else {
		goto L141
	}
L141:
	;
	goto L139
L142:
	;
	if v658 != 0 {
		v685 = v658
		goto L124
	} else {
		goto L143
	}
L143:
	;
	v706 = int32(0)
	goto L64
L144:
	;
	goto L123
L145:
	;
	F_decrRefCount(m, v256)
	mBase = m.M
	v870 = m.ExcPending
	if v870 != 0 {
		goto L1
	} else {
		goto L184
	}
L146:
	;
	v798 = F_setTypeSize(m, v256)
	mBase = m.M
	v799 = m.ExcPending
	if v799 != 0 {
		goto L1
	} else {
		goto L168
	}
L147:
	;
	F_addReplySetLen(m, l0, v706)
	mBase = m.M
	v714 = m.ExcPending
	if v714 != 0 {
		goto L1
	} else {
		goto L148
	}
L148:
	;
	v715 = F_setTypeInitIterator(m, v256)
	mBase = m.M
	v716 = m.ExcPending
	if v716 != 0 {
		goto L1
	} else {
		goto L150
	}
L149:
	;
	v782 = *(*int32)(unsafe.Add(mBase, uint32(v715)+4))
	if v782 != int32(2) {
		goto L162
	} else {
		goto L163
	}
L150:
	;
	v723 = F_setTypeNext(m, v715, v21+int32(24), v21+int32(20), v21+int32(8))
	mBase = m.M
	v724 = m.ExcPending
	if v724 != 0 {
		goto L1
	} else {
		goto L151
	}
L151:
	;
	if v723 == int32(-1) {
		goto L149
	} else {
		goto L152
	}
L152:
	;
	goto L153
L153:
	;
	v745 = *(*int32)(unsafe.Add(mBase, uint32(v21)+24))
	if v745 == int32(0) {
		goto L156
	} else {
		goto L157
	}
L154:
	;
	goto L149
L155:
	;
	v760 = F_setTypeNext(m, v715, v21+int32(24), v21+int32(20), v21+int32(8))
	mBase = m.M
	v761 = m.ExcPending
	if v761 != 0 {
		goto L1
	} else {
		goto L160
	}
L156:
	;
	v751 = *(*int64)(unsafe.Add(mBase, uint32(v21)+8))
	F_addReplyBulkLongLong(m, l0, v751)
	mBase = m.M
	v753 = m.ExcPending
	if v753 != 0 {
		goto L1
	} else {
		goto L159
	}
L157:
	;
	v748 = *(*int32)(unsafe.Add(mBase, uint32(v21)+20))
	F_addReplyBulkCBuffer(m, l0, v745, v748)
	mBase = m.M
	v750 = m.ExcPending
	if v750 != 0 {
		goto L1
	} else {
		goto L158
	}
L158:
	;
	goto L155
L159:
	;
	goto L155
L160:
	;
	if v760 != int32(-1) {
		goto L153
	} else {
		goto L161
	}
L161:
	;
	goto L154
L162:
	;
	F_valkey_free(m, v715)
	mBase = m.M
	v789 = m.ExcPending
	if v789 != 0 {
		goto L1
	} else {
		goto L165
	}
L163:
	;
	v785 = *(*int32)(unsafe.Add(mBase, uint32(v715)+12))
	F_hashtableReleaseIterator(m, v785)
	mBase = m.M
	v787 = m.ExcPending
	if v787 != 0 {
		goto L1
	} else {
		goto L164
	}
L164:
	;
	goto L162
L165:
	;
	v791 = *(*int32)(unsafe.Add(mBase, _c_F_sunionDiffGenericCommand[0]))
	if v791 == int32(0) {
		goto L145
	} else {
		goto L166
	}
L166:
	;
	F_freeObjAsync(m, int32(0), v256, int32(-1))
	mBase = m.M
	v797 = m.ExcPending
	if v797 != 0 {
		goto L1
	} else {
		goto L167
	}
L167:
	;
	goto L3
L168:
	;
	v800 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	if v798 == int32(0) {
		goto L169
	} else {
		goto L170
	}
L169:
	;
	v828 = F_dbDelete(m, v800, l3)
	mBase = m.M
	v829 = m.ExcPending
	if v829 != 0 {
		goto L1
	} else {
		goto L179
	}
L170:
	;
	F_setKey(m, l0, v800, l3, v21+int32(28), int32(0))
	mBase = m.M
	v807 = m.ExcPending
	if v807 != 0 {
		goto L1
	} else {
		goto L171
	}
L171:
	;
	if l4 != 0 {
		goto L172
	} else {
		goto L173
	}
L172:
	;
	v811 = int32(_a_F_sunionDiffGenericCommand_0)
	goto L174
L173:
	;
	v811 = int32(_a_F_sunionDiffGenericCommand_1)
	goto L174
L174:
	;
	v812 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v813 = *(*int32)(unsafe.Add(mBase, uint32(v812)+28))
	F_notifyKeyspaceEvent(m, int32(32), v811, l3, v813)
	mBase = m.M
	v815 = m.ExcPending
	if v815 != 0 {
		goto L1
	} else {
		goto L175
	}
L175:
	;
	v816 = int32(_a_F_sunionDiffGenericCommand_2)
	v818 = *(*int64)(unsafe.Add(mBase, _c_F_sunionDiffGenericCommand[1]))
	*(*int64)(unsafe.Add(mBase, _c_F_sunionDiffGenericCommand[1])) = v818 + int64(1)
	v822 = *(*int32)(unsafe.Add(mBase, uint32(v21)+28))
	v823 = F_setTypeSize(m, v822)
	mBase = m.M
	v824 = m.ExcPending
	if v824 != 0 {
		goto L1
	} else {
		goto L176
	}
L176:
	;
	F_addReplyLongLong(m, l0, base.I64_extend_i32_u(v823))
	mBase = m.M
	v827 = m.ExcPending
	if v827 != 0 {
		goto L1
	} else {
		goto L177
	}
L177:
	;
	goto L3
L178:
	;
	v848 = *(*int32)(unsafe.Add(mBase, _c_F_sunionDiffGenericCommand[2]))
	F_addReply(m, l0, v848)
	mBase = m.M
	v850 = m.ExcPending
	if v850 != 0 {
		goto L1
	} else {
		goto L183
	}
L179:
	;
	if v828 == int32(0) {
		goto L178
	} else {
		goto L180
	}
L180:
	;
	v832 = int32(_a_F_sunionDiffGenericCommand_2)
	v834 = *(*int64)(unsafe.Add(mBase, _c_F_sunionDiffGenericCommand[1]))
	*(*int64)(unsafe.Add(mBase, _c_F_sunionDiffGenericCommand[1])) = v834 + int64(1)
	v838 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	F_signalModifiedKey(m, l0, v838, l3)
	mBase = m.M
	v840 = m.ExcPending
	if v840 != 0 {
		goto L1
	} else {
		goto L181
	}
L181:
	;
	v843 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v844 = *(*int32)(unsafe.Add(mBase, uint32(v843)+28))
	F_notifyKeyspaceEvent(m, int32(4), int32(_a_F_sunionDiffGenericCommand_3), l3, v844)
	mBase = m.M
	v846 = m.ExcPending
	if v846 != 0 {
		goto L1
	} else {
		goto L182
	}
L182:
	;
	goto L178
L183:
	;
	goto L145
L184:
	;
	goto L3
L185:
	;
	m.G0 = v21 + int32(32)
	return
}
func F_sunionstoreCommand(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v9 = *(*int32)(unsafe.Add(mBase, uint32(v3)+4))
	F_sunionDiffGenericCommand(m, l0, v3+int32(8), v6+int32(-2), v9, int32(0))
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		return
	} else {
		return
	}
}
func F_swapdbDbIdArgs(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
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
	var v26 int64
	_ = v26
	var v30 int64
	_ = v30
	var v32 int64
	_ = v32
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v43 int32
	_ = v43
	v8 = m.G0
	v10 = v8 - int32(16)
	m.G0 = v10
	if l1 < int32(3) {
		v43 = int32(0)
		m.G0 = v10 + int32(16)
		return v43
	} else {
		v15 = int32(0)
		v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		v19 = F_getLongLongFromObject(m, v16, v10+int32(8))
		mBase = m.M
		v22 = m.ExcPending
		if v22 != 0 {
			return int32(0)
		} else {
			if v19 != 0 {
				v43 = v15
				m.G0 = v10 + int32(16)
				return v43
			} else {
				v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				v24 = F_getLongLongFromObject(m, v23, v10)
				mBase = m.M
				v25 = m.ExcPending
				if v25 != 0 {
					return int32(0)
				} else {
					if v24 != 0 {
						v43 = v15
						m.G0 = v10 + int32(16)
						return v43
					} else {
						v26 = *(*int64)(unsafe.Add(mBase, uint32(v10)+8))
						if v26 < int64(0) {
							v43 = v15
							m.G0 = v10 + int32(16)
							return v43
						} else {
							v30 = int64(*(*int32)(unsafe.Add(mBase, _c_F_swapdbDbIdArgs[0])))
							if v30 <= v26 {
								v43 = v15
								m.G0 = v10 + int32(16)
								return v43
							} else {
								v32 = *(*int64)(unsafe.Add(mBase, uint32(v10)))
								if v32 < int64(0) {
									v43 = v15
									m.G0 = v10 + int32(16)
									return v43
								} else {
									if v30 <= v32 {
										v43 = v15
										m.G0 = v10 + int32(16)
										return v43
									} else {
										v37 = F_valkey_malloc(m, int32(8))
										mBase = m.M
										v38 = m.ExcPending
										if v38 != 0 {
											return int32(0)
										} else {
											*(*int64)(unsafe.Add(mBase, uint32(v37))) = int64(8589934593)
											*(*int32)(unsafe.Add(mBase, uint32(l2))) = int32(2)
											v43 = v37
											m.G0 = v10 + int32(16)
											return v43
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
func F_syscheck(m *base.Module) int32 {
	mBase := m.M
	_ = mBase
	var v1 int32
	_ = v1
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
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
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v53 int32
	_ = v53
	v1 = int32(0)
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	*(*int32)(unsafe.Add(mBase, uint32(v7)+12)) = v1
	v11 = int32(1)
	v13 = *(*int32)(unsafe.Add(mBase, _c_F_syscheck[0]))
	if v13 == v1 {
		v53 = v11
		goto L1
	} else {
		goto L2
	}
L1:
	;
	m.G0 = v7 + int32(16)
	return v53
L2:
	;
	v18 = v11
	v19 = v13
	v20 = int32(_a_F_syscheck_0)
	goto L3
L3:
	;
	v23 = m.T0[v19].(func(*base.Module, int32) int32)(m, v7+int32(12))
	mBase = m.M
	v26 = m.ExcPending
	if v26 != 0 {
		goto L5
	} else {
		goto L6
	}
L4:
	;
	v53 = v48
	goto L1
L5:
	;
	return int32(0)
L6:
	;
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v20)))
	*(*int32)(unsafe.Add(mBase, uint32(v7))) = v27
	v30 = F_iprintf(m, int32(_a_F_syscheck_1), v7)
	mBase = m.M
	v31 = m.ExcPending
	if v31 != 0 {
		goto L5
	} else {
		goto L7
	}
L7:
	;
	switch v23 {
	case 0:
		goto L11
	case 1:
		goto L10
	default:
		goto L9
	}
L8:
	;
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v20)+12))
	if v49 != 0 {
		v18 = v48
		v19 = v49
		v20 = v20 + int32(8)
		goto L3
	} else {
		goto L17
	}
L9:
	;
	v39 = F_puts(m, int32(_a_F_syscheck_2))
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
		goto L5
	} else {
		goto L14
	}
L10:
	;
	v36 = F_puts(m, int32(_a_F_syscheck_3))
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		goto L5
	} else {
		goto L13
	}
L11:
	;
	v33 = F_puts(m, int32(_a_F_syscheck_4))
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
		goto L5
	} else {
		goto L12
	}
L12:
	;
	v48 = v18
	goto L8
L13:
	;
	v48 = v18
	goto L8
L14:
	;
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v7)+12))
	v42 = F_puts(m, v41)
	mBase = m.M
	v43 = m.ExcPending
	if v43 != 0 {
		goto L5
	} else {
		goto L15
	}
L15:
	;
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v7)+12))
	F_sdsfree(m, v44)
	mBase = m.M
	v46 = m.ExcPending
	if v46 != 0 {
		goto L5
	} else {
		goto L16
	}
L16:
	;
	v48 = int32(0)
	goto L8
L17:
	;
	goto L4
}
