package p1

import (
	base "github.com/shibukawa/vkmem/internal/aot/vkaot/base"
	"unsafe"
)

func F_rioBufferFlush(m *base.Module, l0 int32) int32 {
	return int32(1)
}
func F_rioBufferRead(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v34 int64
	_ = v34
	var v36 int64
	_ = v36
	var v42 int32
	_ = v42
	var v44 int64
	_ = v44
	var v48 int32
	_ = v48
	v4 = int32(0)
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v15 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12+int32(-1)))))
	switch v15 & int32(7) {
	case 0:
		v32 = int32(base.Ui32(v15) >> (uint(int32(3)) % 32))
	case 1:
		v22 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12+int32(-3)))))
		v32 = v22
	case 2:
		v25 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v12+int32(-5)))))
		v32 = v25
	case 3:
		v28 = *(*int32)(unsafe.Add(mBase, uint32(v12+int32(-9))))
		v32 = v28
	case 4:
		v31 = *(*int32)(unsafe.Add(mBase, uint32(v12+int32(-17))))
		v32 = v31
	default:
		v32 = v4
	}
	v34 = *(*int64)(unsafe.Add(mBase, uint32(l0)+56))
	v36 = base.I64_extend_i32_u(l2)
	if base.I64_extend_i32_u(v32)-v34 < v36 {
		v48 = v4
	} else {
		if l2 == int32(0) {
		} else {
			v42 = F__emscripten_memcpy_bulkmem(m, l1, v12+base.I32_wrap_i64(v34), l2)
			mBase = m.M
		}
		v44 = *(*int64)(unsafe.Add(mBase, uint32(l0)+56))
		*(*int64)(unsafe.Add(mBase, uint32(l0)+56)) = v44 + v36
		v48 = int32(1)
	}
	return v48
}
func F_rioBufferTell(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int64
	_ = v2
	v2 = *(*int64)(unsafe.Add(mBase, uint32(l0)+56))
	return v2
}
func F_rioConnTell(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int64
	_ = v2
	v2 = int64(*(*uint32)(unsafe.Add(mBase, uint32(l0)+72)))
	return v2
}
func F_rioConnWrite(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	return int32(0)
}
func F_rioFdFlush(m *base.Module, l0 int32) int32 {
	var v2 int32
	_ = v2
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	v2 = int32(0)
	v4 = F_rioFdWrite(m, l0, v2, v2)
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		return v4
	}
}
func F_rioFileFlush(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v3 = F_fflush(m, v2)
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		return base.B2i32(v3 == int32(0))
	}
}
func F_rioFileWrite(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v10 int64
	_ = v10
	var v15 int64
	_ = v15
	var v20 int64
	_ = v20
	var v21 int32
	_ = v21
	var v26 int64
	_ = v26
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v43 int64
	_ = v43
	var v45 int64
	_ = v45
	var v47 int32
	_ = v47
	var v48 int64
	_ = v48
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v56 int64
	_ = v56
	var v57 int64
	_ = v57
	var v62 int64
	_ = v62
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v76 int32
	_ = v76
	var v80 int32
	_ = v80
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v88 int32
	_ = v88
	var v93 int32
	_ = v93
	var v96 int32
	_ = v96
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v105 int32
	_ = v105
	var v109 int32
	_ = v109
	var v117 int64
	_ = v117
	var v120 int64
	_ = v120
	var v138 int32
	_ = v138
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v151 int32
	_ = v151
	var v157 int32
	_ = v157
	var v163 int32
	_ = v163
	v10 = *(*int64)(unsafe.Add(mBase, uint32(l0)+64))
	if v10 == int64(0) {
		goto L4
	} else {
		goto L5
	}
L1:
	;
	F__serverAssert(m, int32(_a2015), int32(_a2016), int32(132))
	mBase = m.M
	v163 = m.ExcPending
	if v163 != 0 {
		goto L15
	} else {
		goto L45
	}
L2:
	;
	F__serverAssert(m, int32(_a2017), int32(_a2016), int32(131))
	mBase = m.M
	v157 = m.ExcPending
	if v157 != 0 {
		goto L15
	} else {
		goto L44
	}
L3:
	;
	F__serverAssert(m, int32(_a2018), int32(_a2016), int32(119))
	mBase = m.M
	v151 = m.ExcPending
	if v151 != 0 {
		goto L15
	} else {
		goto L43
	}
L4:
	;
	v143 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v144 = F_fwrite(m, l1, l2, int32(1), v143)
	mBase = m.M
	v145 = m.ExcPending
	if v145 != 0 {
		goto L15
	} else {
		goto L42
	}
L5:
	;
	if l2 == int32(0) {
		goto L7
	} else {
		goto L8
	}
L6:
	;
	return v138
L7:
	;
	v138 = int32(1)
	goto L6
L8:
	;
	v15 = *(*int64)(unsafe.Add(mBase, uint32(l0)+56))
	v20 = v15
	v21 = int32(0)
	goto L9
L9:
	;
	v26 = *(*int64)(unsafe.Add(mBase, uint32(l0)+64))
	if v26 <= v20 {
		goto L3
	} else {
		goto L11
	}
L10:
	;
	goto L7
L11:
	;
	v28 = int32(0)
	v30 = l2 - v21
	v32 = base.I32_wrap_i64(v26 - v20)
	if base.Ui32(v30) < base.Ui32(v32) {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	v34 = v30
	goto L14
L13:
	;
	v34 = v32
	goto L14
L14:
	;
	v36 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v37 = F_fwrite(m, l1+v21, v34, int32(1), v36)
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	return int32(0)
L16:
	;
	if v37 == int32(0) {
		v138 = v28
		goto L6
	} else {
		goto L17
	}
L17:
	;
	v43 = *(*int64)(unsafe.Add(mBase, uint32(l0)+56))
	v45 = v43 + base.I64_extend_i32_u(v34)
	*(*int64)(unsafe.Add(mBase, uint32(l0)+56)) = v45
	v47 = v34 + v21
	v48 = *(*int64)(unsafe.Add(mBase, uint32(l0)+64))
	if v45 < v48 {
		v120 = v45
		goto L18
	} else {
		goto L19
	}
L18:
	;
	if l2 != v47 {
		v20 = v120
		v21 = v47
		goto L9
	} else {
		goto L41
	}
L19:
	;
	v50 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v51 = F_fflush(m, v50)
	mBase = m.M
	v52 = m.ExcPending
	if v52 != 0 {
		goto L15
	} else {
		goto L20
	}
L20:
	;
	v53 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v56 = *(*int64)(unsafe.Add(mBase, uint32(l0)+64))
	v57 = base.I64_rem_s(base.I64_extend_i32_u(v53+v47), v56)
	if base.B2i32(v57 == int64(0)) == int32(0) {
		goto L2
	} else {
		goto L21
	}
L21:
	;
	v62 = *(*int64)(unsafe.Add(mBase, uint32(l0)+56))
	if v62 != v56 {
		goto L1
	} else {
		goto L22
	}
L22:
	;
	v64 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v64)+76))
	if int32(-1) < v67 {
		goto L25
	} else {
		goto L26
	}
L23:
	;
	v85 = F_fsync(m, v84)
	mBase = m.M
	if v85 == int32(-1) {
		v138 = v28
		goto L6
	} else {
		goto L30
	}
L24:
	;
	if int32(-1) < v76 {
		v84 = v76
		goto L28
	} else {
		goto L29
	}
L25:
	;
	v71 = F___lockfile(m, v64)
	mBase = m.M
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v64)+60))
	if v71 == int32(0) {
		v76 = v72
		goto L24
	} else {
		goto L27
	}
L26:
	;
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v64)+60))
	v76 = v70
	goto L24
L27:
	;
	F___unlockfile(m, v64)
	mBase = m.M
	v76 = v72
	goto L24
L28:
	;
	goto L23
L29:
	;
	v80 = F___errno_location(m)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v80))) = int32(8)
	v84 = int32(-1)
	goto L28
L30:
	;
	v88 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+72)))
	if v88&int32(1) == int32(0) {
		goto L31
	} else {
		goto L32
	}
L31:
	;
	v117 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(l0)+56)) = v117
	v120 = v117
	goto L18
L32:
	;
	v93 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v96 = *(*int32)(unsafe.Add(mBase, uint32(v93)+76))
	if int32(-1) < v96 {
		goto L35
	} else {
		goto L36
	}
L33:
	;
	goto L40
L34:
	;
	if int32(-1) < v105 {
		goto L38
	} else {
		goto L39
	}
L35:
	;
	v100 = F___lockfile(m, v93)
	mBase = m.M
	v101 = *(*int32)(unsafe.Add(mBase, uint32(v93)+60))
	if v100 == int32(0) {
		v105 = v101
		goto L34
	} else {
		goto L37
	}
L36:
	;
	v99 = *(*int32)(unsafe.Add(mBase, uint32(v93)+60))
	v105 = v99
	goto L34
L37:
	;
	F___unlockfile(m, v93)
	mBase = m.M
	v105 = v101
	goto L34
L38:
	;
	goto L33
L39:
	;
	v109 = F___errno_location(m)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v109))) = int32(8)
	goto L38
L40:
	;
	goto L31
L41:
	;
	goto L10
L42:
	;
	return v144
L43:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L44:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L45:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_rioFreeFd(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v4 int32
	_ = v4
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
	F_sdsfree(m, v2)
	mBase = m.M
	v4 = m.ExcPending
	if v4 != 0 {
		return
	} else {
		return
	}
}
func F_rioInitWithBuffer(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	v7 = F__emscripten_memcpy_bulkmem(m, l0, int32(_a201), int32(80))
	mBase = m.M
	*(*int64)(unsafe.Add(mBase, uint32(v7)+56)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v7)+48)) = l1
	return
}
func F_rioInitWithConnset(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v25 int32
	_ = v25
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
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v74 int32
	_ = v74
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	v13 = F__emscripten_memcpy_bulkmem(m, l0, int32(_a2019), int32(80))
	mBase = m.M
	v16 = l2 << (uint(int32(2)) % 32)
	v17 = F_valkey_malloc(m, v16)
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		return
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(v13)+48)) = v17
		v20 = F_valkey_malloc(m, v16)
		mBase = m.M
		v21 = m.ExcPending
		if v21 != 0 {
			return
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v13)+52)) = v20
			if l2 < int32(1) {
			} else {
				v25 = int32(1)
				if l2 == v25 {
					v74 = int32(0)
				} else {
					v32 = int32(0)
					v39 = v32
					v41 = v32
					for {
						v42 = *(*int32)(unsafe.Add(mBase, uint32(v13)+48))
						v43 = int32(2)
						v44 = v39 << (uint(v43) % 32)
						v47 = *(*int32)(unsafe.Add(mBase, uint32(l1+v44)))
						*(*int32)(unsafe.Add(mBase, uint32(v42+v44))) = v47
						v49 = *(*int32)(unsafe.Add(mBase, uint32(v13)+52))
						v51 = int32(0)
						*(*int32)(unsafe.Add(mBase, uint32(v49+v44))) = v51
						v53 = *(*int32)(unsafe.Add(mBase, uint32(v13)+48))
						v55 = v44 | int32(4)
						v58 = *(*int32)(unsafe.Add(mBase, uint32(l1+v55)))
						*(*int32)(unsafe.Add(mBase, uint32(v53+v55))) = v58
						v60 = *(*int32)(unsafe.Add(mBase, uint32(v13)+52))
						*(*int32)(unsafe.Add(mBase, uint32(v60+v55))) = v51
						v65 = v39 + v43
						v67 = v41 + v43
						if v67 != l2&int32(2147483646) {
							v39 = v65
							v41 = v67
							continue
						} else {
							break
						}
						break
					}
					v74 = v65
				}
				if l2&v25 == int32(0) {
				} else {
					v79 = *(*int32)(unsafe.Add(mBase, uint32(v13)+48))
					v81 = v74 << (uint(int32(2)) % 32)
					v84 = *(*int32)(unsafe.Add(mBase, uint32(l1+v81)))
					*(*int32)(unsafe.Add(mBase, uint32(v79+v81))) = v84
					v86 = *(*int32)(unsafe.Add(mBase, uint32(v13)+52))
					*(*int32)(unsafe.Add(mBase, uint32(v86+v81))) = int32(0)
				}
			}
			*(*int64)(unsafe.Add(mBase, uint32(v13)+64)) = int64(0)
			*(*int32)(unsafe.Add(mBase, uint32(v13)+56)) = l2
			v101 = F_sdsempty(m)
			mBase = m.M
			v102 = m.ExcPending
			if v102 != 0 {
				return
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v13)+72)) = v101
				return
			}
		}
	}
}
func F_rioRead_1(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
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
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v26 int64
	_ = v26
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v41 int32
	_ = v41
	var v51 int32
	_ = v51
	v7 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
	if v7&int32(5) != 0 {
		v51 = int32(0)
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return v51
L2:
	;
	if l2 == int32(0) {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v51 = int32(1)
	goto L1
L4:
	;
	v13 = l1
	v14 = l2
	goto L5
L5:
	;
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	if base.Ui32(v17) < base.Ui32(v14) {
		goto L8
	} else {
		goto L9
	}
L6:
	;
	goto L3
L7:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v32 == int32(0) {
		goto L17
	} else {
		goto L18
	}
L8:
	;
	v19 = v17
	goto L10
L9:
	;
	v19 = v14
	goto L10
L10:
	;
	if v17 != 0 {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v20 = v19
	goto L13
L12:
	;
	v20 = v14
	goto L13
L13:
	;
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v22 = m.T0[v21].(func(*base.Module, int32, int32, int32) int32)(m, l0, v13, v20)
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	return int32(0)
L15:
	;
	if v22 != 0 {
		goto L7
	} else {
		goto L16
	}
L16:
	;
	v26 = *(*int64)(unsafe.Add(mBase, uint32(l0)+32))
	*(*int64)(unsafe.Add(mBase, uint32(l0)+32)) = v26 | int64(1)
	return int32(0)
L17:
	;
	v37 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = v37 + v20
	v41 = v14 - v20
	if v41 != 0 {
		v13 = v13 + v20
		v14 = v41
		goto L5
	} else {
		goto L20
	}
L18:
	;
	m.T0[v32].(func(*base.Module, int32, int32, int32))(m, l0, v13, v20)
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
		goto L14
	} else {
		goto L19
	}
L19:
	;
	goto L17
L20:
	;
	goto L6
}
func F_rioRead_2(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v28 int64
	_ = v28
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v51 int32
	_ = v51
	v3 = int32(0)
	v7 = int32(*(*uint8)(unsafe.Add(mBase, _consts[916])))
	if v7&int32(5) != 0 {
		v51 = v3
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return v51
L2:
	;
	v10 = l0
	v11 = l1
	goto L3
L3:
	;
	v16 = *(*int32)(unsafe.Add(mBase, _consts[917]))
	if base.Ui32(v16) < base.Ui32(v11) {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	v51 = int32(1)
	goto L1
L5:
	;
	v34 = int32(0)
	v35 = *(*int32)(unsafe.Add(mBase, _consts[918]))
	if v35 == v34 {
		goto L15
	} else {
		goto L16
	}
L6:
	;
	v18 = v16
	goto L8
L7:
	;
	v18 = v11
	goto L8
L8:
	;
	if v16 != 0 {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	v19 = v18
	goto L11
L10:
	;
	v19 = v11
	goto L11
L11:
	;
	v21 = *(*int32)(unsafe.Add(mBase, _consts[919]))
	v22 = m.T0[v21].(func(*base.Module, int32, int32, int32) int32)(m, int32(_a2497), v10, v19)
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	return int32(0)
L13:
	;
	if v22 != 0 {
		goto L5
	} else {
		goto L14
	}
L14:
	;
	v26 = int32(0)
	v28 = *(*int64)(unsafe.Add(mBase, _consts[916]))
	*(*int64)(unsafe.Add(mBase, _consts[916])) = v28 | int64(1)
	return v26
L15:
	;
	v41 = int32(0)
	v43 = *(*int32)(unsafe.Add(mBase, _consts[920]))
	*(*int32)(unsafe.Add(mBase, _consts[920])) = v43 + v19
	v47 = v11 - v19
	if v47 != 0 {
		v10 = v10 + v19
		v11 = v47
		goto L3
	} else {
		goto L18
	}
L16:
	;
	m.T0[v35].(func(*base.Module, int32, int32, int32))(m, int32(_a2497), v10, v19)
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
		goto L12
	} else {
		goto L17
	}
L17:
	;
	goto L15
L18:
	;
	goto L4
}
func F_rioSetAutoSync(m *base.Module, l0 int32, l1 int64) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v3 != int32(980) {
	} else {
		*(*int64)(unsafe.Add(mBase, uint32(l0)+64)) = l1
	}
	return
}
func F_rioWriteBulkDouble(m *base.Module, l0 int32, l1 float64) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	v3 = int32(0)
	v5 = m.G0
	v7 = v5 - int32(128)
	m.G0 = v7
	v9 = F_fpconv_dtoa(m, l1, v7)
	mBase = m.M
	*(*uint8)(unsafe.Add(mBase, uint32(v7+v9))) = uint8(v3)
	v13 = F_rioWriteBulkString(m, l0, v7, v9)
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		return int32(0)
	} else {
		m.G0 = v7 + int32(128)
		return v13
	}
}
