package p0

import (
	base "github.com/shibukawa/vkmem/internal/aot/vkaot/base"
	"unsafe"
)

func F___quicklistCompress(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v24 int32
	_ = v24
	var v32 int32
	_ = v32
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
	var v43 int32
	_ = v43
	var v51 int32
	_ = v51
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v83 int32
	_ = v83
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v105 int32
	_ = v105
	var v114 int32
	_ = v114
	var v119 int32
	_ = v119
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v126 int32
	_ = v126
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v135 int32
	_ = v135
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v144 int32
	_ = v144
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v163 int32
	_ = v163
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v9 == int32(0) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	F__serverAssert(m, int32(_a_F___quicklistCompress_0), int32(_a_F___quicklistCompress_1), int32(303))
	mBase = m.M
	v163 = m.ExcPending
	if v163 != 0 {
		goto L15
	} else {
		goto L49
	}
L2:
	;
	return
L3:
	;
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v13 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+18)))
	if v13&int32(16) != 0 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v17 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16)+18)))
	if v17&int32(16) != 0 {
		goto L1
	} else {
		goto L5
	}
L5:
	;
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v24 = int32(base.Ui32(v20)>>(uint(int32(14))%32)) & int32(16383)
	if v24 == int32(0) {
		goto L2
	} else {
		goto L6
	}
L6:
	;
	if base.Ui32(v9) < base.Ui32(v24<<(uint(int32(1))%32)) {
		goto L2
	} else {
		goto L7
	}
L7:
	;
	v32 = int32(0)
	v36 = l0 + int32(4)
	v37 = v12
	v38 = v32
	v39 = v32
	goto L9
L8:
	;
	if l1 == int32(0) {
		goto L37
	} else {
		goto L38
	}
L9:
	;
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v36)))
	v43 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if base.Ui32(int32(base.Ui32(v43)>>(uint(int32(14))%32))&int32(16383)) <= base.Ui32(v38) {
		goto L8
	} else {
		goto L11
	}
L11:
	;
	if v37 == int32(0) {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	if v42 == int32(0) {
		goto L21
	} else {
		goto L22
	}
L13:
	;
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v37)+16))
	if v51&int32(196608) != int32(131072) {
		goto L12
	} else {
		goto L14
	}
L14:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v37)+16)) = v51 & int32(-3211265)
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v37)+12))
	v60 = F_valkey_malloc(m, v59)
	mBase = m.M
	v61 = m.ExcPending
	if v61 != 0 {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	return
L16:
	;
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v37)+8))
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v62)))
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v37)+12))
	v67 = F_lzf_decompress(m, v62+int32(4), v65, v60, v66)
	mBase = m.M
	if v67 != 0 {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	F_valkey_free(m, v62)
	mBase = m.M
	v71 = m.ExcPending
	if v71 != 0 {
		goto L15
	} else {
		goto L20
	}
L18:
	;
	F_valkey_free(m, v60)
	mBase = m.M
	v69 = m.ExcPending
	if v69 != 0 {
		goto L15
	} else {
		goto L19
	}
L19:
	;
	goto L12
L20:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v37)+8)) = v60
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v37)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v37)+16)) = v73&int32(-196609) | int32(65536)
	goto L12
L21:
	;
	if v37 == v42 {
		goto L2
	} else {
		goto L29
	}
L22:
	;
	v83 = *(*int32)(unsafe.Add(mBase, uint32(v42)+16))
	if v83&int32(196608) != int32(131072) {
		goto L21
	} else {
		goto L23
	}
L23:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v42)+16)) = v83 & int32(-3211265)
	v91 = *(*int32)(unsafe.Add(mBase, uint32(v42)+12))
	v92 = F_valkey_malloc(m, v91)
	mBase = m.M
	v93 = m.ExcPending
	if v93 != 0 {
		goto L15
	} else {
		goto L24
	}
L24:
	;
	v94 = *(*int32)(unsafe.Add(mBase, uint32(v42)+8))
	v97 = *(*int32)(unsafe.Add(mBase, uint32(v94)))
	v98 = *(*int32)(unsafe.Add(mBase, uint32(v42)+12))
	v99 = F_lzf_decompress(m, v94+int32(4), v97, v92, v98)
	mBase = m.M
	if v99 != 0 {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	F_valkey_free(m, v94)
	mBase = m.M
	v103 = m.ExcPending
	if v103 != 0 {
		goto L15
	} else {
		goto L28
	}
L26:
	;
	F_valkey_free(m, v92)
	mBase = m.M
	v101 = m.ExcPending
	if v101 != 0 {
		goto L15
	} else {
		goto L27
	}
L27:
	;
	goto L21
L28:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v42)+8)) = v92
	v105 = *(*int32)(unsafe.Add(mBase, uint32(v42)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v42)+16)) = v105&int32(-196609) | int32(65536)
	goto L21
L29:
	;
	v114 = int32(1)
	if v42 == l1 {
		goto L30
	} else {
		goto L31
	}
L30:
	;
	v119 = v114
	goto L32
L31:
	;
	v119 = v39
	goto L32
L32:
	;
	if v37 == l1 {
		goto L33
	} else {
		goto L34
	}
L33:
	;
	v121 = v114
	goto L35
L34:
	;
	v121 = v119
	goto L35
L35:
	;
	v122 = *(*int32)(unsafe.Add(mBase, uint32(v37)+4))
	if v122 != v42 {
		v36 = v42
		v37 = v122
		v38 = v38 + v114
		v39 = v121
		goto L9
	} else {
		goto L36
	}
L36:
	;
	goto L2
L37:
	;
	if v37 == int32(0) {
		goto L42
	} else {
		goto L43
	}
L38:
	;
	if v39 != 0 {
		goto L37
	} else {
		goto L39
	}
L39:
	;
	v126 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	if v126&int32(196608) != int32(65536) {
		goto L37
	} else {
		goto L40
	}
L40:
	;
	v131 = F___quicklistCompressNode(m, l1)
	mBase = m.M
	v132 = m.ExcPending
	if v132 != 0 {
		goto L15
	} else {
		goto L41
	}
L41:
	;
	goto L37
L42:
	;
	if v42 == int32(0) {
		goto L2
	} else {
		goto L46
	}
L43:
	;
	v135 = *(*int32)(unsafe.Add(mBase, uint32(v37)+16))
	if v135&int32(196608) != int32(65536) {
		goto L42
	} else {
		goto L44
	}
L44:
	;
	v140 = F___quicklistCompressNode(m, v37)
	mBase = m.M
	v141 = m.ExcPending
	if v141 != 0 {
		goto L15
	} else {
		goto L45
	}
L45:
	;
	goto L42
L46:
	;
	v144 = *(*int32)(unsafe.Add(mBase, uint32(v42)+16))
	if v144&int32(196608) != int32(65536) {
		goto L2
	} else {
		goto L47
	}
L47:
	;
	v149 = F___quicklistCompressNode(m, v42)
	mBase = m.M
	v150 = m.ExcPending
	if v150 != 0 {
		goto L15
	} else {
		goto L48
	}
L48:
	;
	goto L2
L49:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
}
func F___quicklistDecompressNode(m *base.Module, l0 int32) {
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
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v4 & int32(-3145729)
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v9 = F_valkey_malloc(m, v8)
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		return
	} else {
		v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		v14 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
		v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
		v16 = F_lzf_decompress(m, v11+int32(4), v14, v9, v15)
		mBase = m.M
		if v16 != 0 {
			F_valkey_free(m, v11)
			mBase = m.M
			v20 = m.ExcPending
			if v20 != 0 {
				return
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v9
				v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
				*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v22&int32(-196609) | int32(65536)
				return
			}
		} else {
			F_valkey_free(m, v9)
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
func F___quicklistInsertPlainNode(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	v9 = F_valkey_malloc(m, int32(20))
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		return
	} else {
		*(*int64)(unsafe.Add(mBase, uint32(v9))) = int64(0)
		v13 = *(*int32)(unsafe.Add(mBase, uint32(v9)+16))
		v14 = F_valkey_malloc(m, l3)
		mBase = m.M
		v15 = m.ExcPending
		if v15 != 0 {
			return
		} else {
			if l3 == int32(0) {
				v19 = v14
			} else {
				v18 = F__emscripten_memcpy_bulkmem(m, v14, l2, l3)
				mBase = m.M
				v19 = v18
			}
			*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = v13&int32(-6291456) | int32(327681)
			*(*int32)(unsafe.Add(mBase, uint32(v9)+12)) = l3
			*(*int32)(unsafe.Add(mBase, uint32(v9)+8)) = v19
			F___quicklistInsertNode(m, l0, l1, v9, l4)
			mBase = m.M
			v28 = m.ExcPending
			if v28 != 0 {
				return
			} else {
				v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v29 + int32(1)
				return
			}
		}
	}
}
func F_quicklistCount(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	return v2
}
func F_quicklistCreateNode(m *base.Module) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v7 int64
	_ = v7
	var v11 int32
	_ = v11
	v3 = F_valkey_malloc(m, int32(20))
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		v7 = int64(0)
		*(*int64)(unsafe.Add(mBase, uint32(v3)+8)) = v7
		*(*int64)(unsafe.Add(mBase, uint32(v3))) = v7
		v11 = *(*int32)(unsafe.Add(mBase, uint32(v3)+16))
		*(*int32)(unsafe.Add(mBase, uint32(v3)+16)) = v11&int32(-6291456) | int32(589824)
		return v3
	}
}
func F_quicklistDelEntry(m *base.Module, l0 int32, l1 int32) {
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
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v52 int32
	_ = v52
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(v9)+4))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(v9)))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(v9)+16))
	if v12&int32(786432) != int32(262144) {
		v19 = *(*int32)(unsafe.Add(mBase, uint32(v9)+8))
		v20 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
		v23 = F_lpDelete(m, v19, v20, l1+int32(8))
		mBase = m.M
		v24 = m.ExcPending
		if v24 != 0 {
			return
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v9)+8)) = v23
			v26 = *(*int32)(unsafe.Add(mBase, uint32(v9)+16))
			v28 = v26 + int32(-1)
			*(*uint16)(unsafe.Add(mBase, uint32(v9)+16)) = uint16(v28)
			if v26&int32(65535) == int32(1) {
				F___quicklistDelNode(m, v8, v9)
				mBase = m.M
				v43 = m.ExcPending
				if v43 != 0 {
					return
				} else {
					v44 = *(*int32)(unsafe.Add(mBase, uint32(v8)+8))
					*(*int32)(unsafe.Add(mBase, uint32(v8)+8)) = v44 + int32(-1)
					*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = int32(0)
					v52 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
					switch v52 {
					case 0:
						*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = int32(0)
						*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v10
						return
					case 1:
						*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = int32(-1)
						*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v11
						return
					default:
						return
					}
				}
			} else {
				v34 = *(*int32)(unsafe.Add(mBase, uint32(v23)))
				*(*int32)(unsafe.Add(mBase, uint32(v9)+12)) = v34
				v36 = *(*int32)(unsafe.Add(mBase, uint32(v8)+8))
				*(*int32)(unsafe.Add(mBase, uint32(v8)+8)) = v36 + int32(-1)
				*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = int32(0)
				return
			}
		}
	} else {
		F___quicklistDelNode(m, v8, v9)
		mBase = m.M
		v18 = m.ExcPending
		if v18 != 0 {
			return
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = int32(0)
			v52 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
			switch v52 {
			case 0:
				*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = int32(0)
				*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v10
				return
			case 1:
				*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = int32(-1)
				*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v11
				return
			default:
				return
			}
		}
	}
}
func F_quicklistGetIteratorAtIdx(m *base.Module, l0 int32, l1 int32, l2 int64) int32 {
	mBase := m.M
	_ = mBase
	var v16 int64
	_ = v16
	var v17 int32
	_ = v17
	var v22 int32
	_ = v22
	var v31 int32
	_ = v31
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v44 int64
	_ = v44
	var v47 int32
	_ = v47
	var v58 int32
	_ = v58
	var v60 int64
	_ = v60
	var v62 int32
	_ = v62
	var v64 int64
	_ = v64
	var v67 int32
	_ = v67
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v83 int64
	_ = v83
	var v85 int64
	_ = v85
	var v90 int64
	_ = v90
	var v97 int32
	_ = v97
	v16 = l2 ^ l2>>(uint(int64(63))%64)
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if base.Ui64(v16) < base.Ui64(base.I64_extend_i32_u(v17)) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v22 = int32(0)
	v31 = v17 + int32(-1)
	v35 = base.B2i32(base.Ui64(base.I64_extend_i32_u(int32(base.Ui32(v31)>>(uint(int32(1))%32)))) < base.Ui64(v16))
	if base.Ui64(base.I64_extend_i32_u(int32(base.Ui32(v31)>>(uint(int32(1))%32)))) < base.Ui64(v16) {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	return int32(0)
L3:
	;
	return v97
L4:
	;
	v36 = base.I32_wrap_i64(int64(base.Ui64(l2) >> (uint(int64(63)) % 64)))
	goto L6
L5:
	;
	v36 = base.B2i32(int64(-1) < l2)
	goto L6
L6:
	;
	if v36 != 0 {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	v37 = v22
	goto L9
L8:
	;
	v37 = int32(4)
	goto L9
L9:
	;
	v39 = *(*int32)(unsafe.Add(mBase, uint32(l0+v37)))
	if v39 == int32(0) {
		v97 = v22
		goto L3
	} else {
		goto L10
	}
L10:
	;
	if base.Ui64(base.I64_extend_i32_u(int32(base.Ui32(v31)>>(uint(int32(1))%32)))) < base.Ui64(v16) {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v44 = base.I64_extend_i32_u(v31) - v16
	goto L13
L12:
	;
	v44 = v16
	goto L13
L13:
	;
	if v36 != 0 {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v47 = int32(4)
	goto L16
L15:
	;
	v47 = int32(0)
	goto L16
L16:
	;
	v58 = v39
	v60 = int64(0)
	goto L18
L17:
	;
	v71 = F_valkey_malloc(m, int32(20))
	mBase = m.M
	v74 = m.ExcPending
	if v74 != 0 {
		goto L22
	} else {
		goto L23
	}
L18:
	;
	v62 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v58)+16)))
	v64 = v60 + base.I64_extend_i32_u(v62)
	if base.Ui64(v44) < base.Ui64(v64) {
		goto L17
	} else {
		goto L20
	}
L19:
	;
	return int32(0)
L20:
	;
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v58+v47)))
	if v67 != 0 {
		v58 = v67
		v60 = v64
		goto L18
	} else {
		goto L21
	}
L21:
	;
	goto L19
L22:
	;
	return int32(0)
L23:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v71)+8)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v71))) = l0
	*(*int32)(unsafe.Add(mBase, uint32(v71)+16)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v71)+4)) = v58
	if base.Ui64(base.I64_extend_i32_u(int32(base.Ui32(v31)>>(uint(int32(1))%32)))) < base.Ui64(v16) {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	v83 = base.I64_extend_i32_u(v17-v62) - v60
	goto L26
L25:
	;
	v83 = v60
	goto L26
L26:
	;
	v85 = int64(-1)
	if v85 < l2 {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	v90 = v16 - v83
	goto L29
L28:
	;
	v90 = v83 + (v16 ^ v85)
	goto L29
L29:
	;
	*(*uint32)(unsafe.Add(mBase, uint32(v71)+12)) = uint32(v90)
	v97 = v71
	goto L3
}
func F_quicklistInsertBefore(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	var v7 int32
	_ = v7
	F__quicklistInsert(m, l0, l1, l2, l3, int32(0))
	v7 = m.ExcPending
	if v7 != 0 {
		return
	} else {
		return
	}
}
func F_quicklistPushHead(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v33 int32
	_ = v33
	var v39 int32
	_ = v39
	var v44 int32
	_ = v44
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v68 int32
	_ = v68
	var v73 int32
	_ = v73
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
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
	var v127 int32
	_ = v127
	var v129 int32
	_ = v129
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v135 int32
	_ = v135
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v142 int32
	_ = v142
	var v144 int32
	_ = v144
	var v146 int32
	_ = v146
	v4 = int32(0)
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v9 = int32(18)
	v12 = v8 << (uint(v9) % 32) >> (uint(v9) % 32)
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v15 = *(*int32)(unsafe.Add(mBase, _c_F_quicklistPushHead[0]))
	if v15 == v4 {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	if v13 == int32(0) {
		goto L16
	} else {
		goto L17
	}
L2:
	;
	F___quicklistInsertPlainNode(m, l0, v13, l1, l2, int32(0))
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L13
	} else {
		goto L14
	}
L3:
	;
	if v12 < int32(0) {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	if base.Ui32(v15) <= base.Ui32(l2) {
		goto L2
	} else {
		goto L5
	}
L5:
	;
	goto L1
L6:
	;
	v23 = int32(-5)
	if base.Ui32(v23) < base.Ui32(v12) {
		goto L9
	} else {
		goto L10
	}
L7:
	;
	if base.Ui32(l2) <= base.Ui32(int32(8192)) {
		goto L1
	} else {
		goto L8
	}
L8:
	;
	goto L2
L9:
	;
	v26 = v12
	goto L11
L10:
	;
	v26 = v23
	goto L11
L11:
	;
	v33 = *(*int32)(unsafe.Add(mBase, uint32((v26^int32(-1))<<(uint(int32(2))%32))+uint32(_c_F_quicklistPushHead[1])))
	if base.Ui32(l2) <= base.Ui32(v33) {
		goto L1
	} else {
		goto L12
	}
L12:
	;
	goto L2
L13:
	;
	return int32(0)
L14:
	;
	return int32(1)
L15:
	;
	v138 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v139 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v138 + v139
	v142 = *(*int32)(unsafe.Add(mBase, uint32(v135)+16))
	v144 = v142 + v139
	*(*uint16)(unsafe.Add(mBase, uint32(v135)+16)) = uint16(v144)
	v146 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	return base.B2i32(v13 != v146)
L16:
	;
	v119 = F_quicklistCreateNode(m)
	mBase = m.M
	v120 = m.ExcPending
	if v120 != 0 {
		goto L13
	} else {
		goto L44
	}
L17:
	;
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v13)+16))
	if v44&int32(786432) == int32(262144) {
		goto L16
	} else {
		goto L18
	}
L18:
	;
	if v15 == int32(0) {
		goto L23
	} else {
		goto L24
	}
L19:
	;
	v108 = *(*int32)(unsafe.Add(mBase, uint32(v13)+8))
	v109 = F_lpPrepend(m, v108, l1, l2)
	mBase = m.M
	v110 = m.ExcPending
	if v110 != 0 {
		goto L13
	} else {
		goto L42
	}
L20:
	;
	if base.Ui32(v103) < base.Ui32(v104) {
		goto L16
	} else {
		goto L41
	}
L21:
	;
	v87 = int32(-5)
	if base.Ui32(v87) < base.Ui32(v12) {
		goto L37
	} else {
		goto L38
	}
L22:
	;
	v80 = int32(1)
	if base.Ui32(v80) < base.Ui32(v12) {
		goto L32
	} else {
		goto L33
	}
L23:
	;
	if v12 < int32(0) {
		goto L21
	} else {
		goto L30
	}
L24:
	;
	if base.Ui32(v15) <= base.Ui32(l2) {
		goto L16
	} else {
		goto L25
	}
L25:
	;
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v13)+12))
	v55 = l2 + v52 + int32(8)
	if int32(-1) < v12 {
		v77 = v55
		goto L22
	} else {
		goto L26
	}
L26:
	;
	v58 = int32(-5)
	if base.Ui32(v58) < base.Ui32(v12) {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	v61 = v12
	goto L29
L28:
	;
	v61 = v58
	goto L29
L29:
	;
	v68 = *(*int32)(unsafe.Add(mBase, uint32((v61^int32(-1))<<(uint(int32(2))%32))+uint32(_c_F_quicklistPushHead[1])))
	v103 = v68
	v104 = v55
	goto L20
L30:
	;
	if base.Ui32(int32(8192)) < base.Ui32(l2) {
		goto L16
	} else {
		goto L31
	}
L31:
	;
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v13)+12))
	v77 = l2 + v73 + int32(8)
	goto L22
L32:
	;
	v83 = v12
	goto L34
L33:
	;
	v83 = v80
	goto L34
L34:
	;
	if base.Ui32(v83) <= base.Ui32(v44&int32(65535)) {
		goto L16
	} else {
		goto L35
	}
L35:
	;
	if base.Ui32(int32(8192)) < base.Ui32(v77) {
		goto L16
	} else {
		goto L36
	}
L36:
	;
	goto L19
L37:
	;
	v90 = v12
	goto L39
L38:
	;
	v90 = v87
	goto L39
L39:
	;
	v97 = *(*int32)(unsafe.Add(mBase, uint32((v90^int32(-1))<<(uint(int32(2))%32))+uint32(_c_F_quicklistPushHead[1])))
	if base.Ui32(v97) < base.Ui32(l2) {
		goto L16
	} else {
		goto L40
	}
L40:
	;
	v99 = *(*int32)(unsafe.Add(mBase, uint32(v13)+12))
	v103 = v97
	v104 = l2 + v99 + int32(8)
	goto L20
L41:
	;
	goto L19
L42:
	;
	v111 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(v111)+8)) = v109
	v113 = *(*int32)(unsafe.Add(mBase, uint32(v109)))
	goto L43
L43:
	;
	v114 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(v114)+12)) = v113
	v135 = v114
	goto L15
L44:
	;
	v122 = F_lpNew(m, int32(0))
	mBase = m.M
	v123 = m.ExcPending
	if v123 != 0 {
		goto L13
	} else {
		goto L45
	}
L45:
	;
	v124 = F_lpPrepend(m, v122, l1, l2)
	mBase = m.M
	v125 = m.ExcPending
	if v125 != 0 {
		goto L13
	} else {
		goto L46
	}
L46:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v119)+8)) = v124
	v127 = *(*int32)(unsafe.Add(mBase, uint32(v124)))
	goto L47
L47:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v119)+12)) = v127
	v129 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	F___quicklistInsertNode(m, l0, v129, v119, int32(0))
	mBase = m.M
	v132 = m.ExcPending
	if v132 != 0 {
		goto L13
	} else {
		goto L48
	}
L48:
	;
	v133 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v135 = v133
	goto L15
}
func F_quicklistRelease(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
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
	var v28 int32
	_ = v28
	var v33 int32
	_ = v33
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v68 int32
	_ = v68
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v5 == int32(0) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v33 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if base.Ui32(v33) < base.Ui32(int32(268435456)) {
		goto L9
	} else {
		goto L10
	}
L2:
	;
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v10 = v5
	v11 = v8
	goto L3
L3:
	;
	v13 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(v11)+8))
	F_valkey_free(m, v14)
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		goto L5
	} else {
		goto L6
	}
L4:
	;
	goto L1
L5:
	;
	return
L6:
	;
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v18 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v11)+16)))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v17 - v18
	F_valkey_free(m, v11)
	mBase = m.M
	v22 = m.ExcPending
	if v22 != 0 {
		goto L5
	} else {
		goto L7
	}
L7:
	;
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v24 = int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v23 + v24
	v28 = v10 + v24
	if v28 != 0 {
		v10 = v28
		v11 = v13
		goto L3
	} else {
		goto L8
	}
L8:
	;
	goto L4
L9:
	;
	F_valkey_free(m, l0)
	mBase = m.M
	v68 = m.ExcPending
	if v68 != 0 {
		goto L5
	} else {
		goto L15
	}
L10:
	;
	v40 = v33
	goto L11
L11:
	;
	v42 = int32(28)
	v45 = int32(base.Ui32(v40)>>(uint(v42)%32)) + int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v45<<(uint(v42)%32) | v40&int32(268435455)
	v57 = *(*int32)(unsafe.Add(mBase, uint32(l0+int32(24)+v45&int32(15)<<(uint(int32(3))%32))))
	F_valkey_free(m, v57)
	mBase = m.M
	v59 = m.ExcPending
	if v59 != 0 {
		goto L5
	} else {
		goto L13
	}
L12:
	;
	goto L9
L13:
	;
	v60 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if base.Ui32(int32(268435455)) < base.Ui32(v60) {
		v40 = v60
		goto L11
	} else {
		goto L14
	}
L14:
	;
	goto L12
L15:
	;
	return
}
func F_quicklistReleaseIterator(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	if l0 == int32(0) {
		return
	} else {
		v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		if v6 == int32(0) {
			F_valkey_free(m, l0)
			mBase = m.M
			v25 = m.ExcPending
			if v25 != 0 {
				return
			} else {
				return
			}
		} else {
			v9 = *(*int32)(unsafe.Add(mBase, uint32(v6)+16))
			if v9&int32(1048576) == int32(0) {
				v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
				F___quicklistCompress(m, v20, v6)
				mBase = m.M
				v22 = m.ExcPending
				if v22 != 0 {
					return
				} else {
					F_valkey_free(m, l0)
					mBase = m.M
					v25 = m.ExcPending
					if v25 != 0 {
						return
					} else {
						return
					}
				}
			} else {
				if v9&int32(196608) != int32(65536) {
					F_valkey_free(m, l0)
					mBase = m.M
					v25 = m.ExcPending
					if v25 != 0 {
						return
					} else {
						return
					}
				} else {
					v18 = F___quicklistCompressNode(m, v6)
					mBase = m.M
					v19 = m.ExcPending
					if v19 != 0 {
						return
					} else {
						F_valkey_free(m, l0)
						mBase = m.M
						v25 = m.ExcPending
						if v25 != 0 {
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
func F_quicklistReplaceEntry(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v46 int32
	_ = v46
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
	var v60 int32
	_ = v60
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v89 int32
	_ = v89
	var v94 int32
	_ = v94
	var v97 int32
	_ = v97
	var v104 int32
	_ = v104
	var v107 int32
	_ = v107
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v117 int32
	_ = v117
	var v119 int32
	_ = v119
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v137 int32
	_ = v137
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v159 int32
	_ = v159
	var v164 int32
	_ = v164
	var v167 int32
	_ = v167
	var v174 int32
	_ = v174
	var v176 int32
	_ = v176
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v184 int32
	_ = v184
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v205 int32
	_ = v205
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v247 int32
	_ = v247
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v259 int32
	_ = v259
	var v263 int32
	_ = v263
	var v264 int32
	_ = v264
	var v266 int32
	_ = v266
	v11 = m.G0
	v13 = v11 - int32(16)
	m.G0 = v13
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v17 = *(*int32)(unsafe.Add(mBase, uint32(v16)+16))
	if v17&int32(786432) != int32(262144) {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l0)+4)) = int64(0)
	m.G0 = v13 + int32(16)
	return
L2:
	;
	if v73&int32(786432) != int32(262144) {
		goto L27
	} else {
		goto L28
	}
L3:
	;
	v22 = int32(0)
	v23 = *(*int32)(unsafe.Add(mBase, _c_F_quicklistReplaceEntry[0]))
	if v23 == v22 {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	v73 = v17
	v74 = v16
	goto L2
L5:
	;
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v16)+8))
	v52 = F_lpReplace(m, v49, l1+int32(8), l2, l3)
	mBase = m.M
	v53 = m.ExcPending
	if v53 != 0 {
		goto L16
	} else {
		goto L17
	}
L6:
	;
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v15)+16))
	v28 = int32(18)
	v31 = v27 << (uint(v28) % 32) >> (uint(v28) % 32)
	if v31 < int32(0) {
		goto L9
	} else {
		goto L10
	}
L7:
	;
	if base.Ui32(l3) < base.Ui32(v23) {
		goto L5
	} else {
		goto L8
	}
L8:
	;
	v73 = v17
	v74 = v16
	goto L2
L9:
	;
	v36 = int32(-5)
	if base.Ui32(v36) < base.Ui32(v31) {
		goto L12
	} else {
		goto L13
	}
L10:
	;
	if base.Ui32(int32(8192)) < base.Ui32(l3) {
		v73 = v17
		v74 = v16
		goto L2
	} else {
		goto L11
	}
L11:
	;
	goto L5
L12:
	;
	v39 = v31
	goto L14
L13:
	;
	v39 = v36
	goto L14
L14:
	;
	v46 = *(*int32)(unsafe.Add(mBase, uint32((v39^int32(-1))<<(uint(int32(2))%32))+uint32(_c_F_quicklistReplaceEntry[1])))
	if base.Ui32(v46) < base.Ui32(l3) {
		v73 = v17
		v74 = v16
		goto L2
	} else {
		goto L15
	}
L15:
	;
	goto L5
L16:
	;
	return
L17:
	;
	v54 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v52 != 0 {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v54)+8)) = v52
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v52)))
	goto L20
L19:
	;
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v54)+16))
	v73 = v55
	v74 = v54
	goto L2
L20:
	;
	v58 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v58)+12)) = v57
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v58)+16))
	if v60&int32(1048576) == int32(0) {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	F___quicklistCompress(m, v15, v58)
	mBase = m.M
	v72 = m.ExcPending
	if v72 != 0 {
		goto L16
	} else {
		goto L25
	}
L22:
	;
	if v60&int32(196608) != int32(65536) {
		goto L1
	} else {
		goto L23
	}
L23:
	;
	v69 = F___quicklistCompressNode(m, v58)
	mBase = m.M
	v70 = m.ExcPending
	if v70 != 0 {
		goto L16
	} else {
		goto L24
	}
L24:
	;
	goto L1
L25:
	;
	goto L1
L26:
	;
	F__quicklistInsert(m, l0, l1, l2, l3, int32(1))
	mBase = m.M
	v263 = m.ExcPending
	if v263 != 0 {
		goto L16
	} else {
		goto L94
	}
L27:
	;
	v132 = *(*int32)(unsafe.Add(mBase, uint32(v16)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v16)+16)) = v132 | int32(4194304)
	v137 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
	if v137 == v132&int32(65535)+int32(-1) {
		v149 = int32(0)
		goto L50
	} else {
		goto L51
	}
L28:
	;
	v80 = int32(0)
	v81 = *(*int32)(unsafe.Add(mBase, _c_F_quicklistReplaceEntry[0]))
	if v81 == v80 {
		goto L30
	} else {
		goto L31
	}
L29:
	;
	v107 = *(*int32)(unsafe.Add(mBase, uint32(v74)+8))
	F_valkey_free(m, v107)
	mBase = m.M
	v109 = m.ExcPending
	if v109 != 0 {
		goto L16
	} else {
		goto L40
	}
L30:
	;
	v85 = *(*int32)(unsafe.Add(mBase, uint32(v15)+16))
	v86 = int32(18)
	v89 = v85 << (uint(v86) % 32) >> (uint(v86) % 32)
	if v89 < int32(0) {
		goto L33
	} else {
		goto L34
	}
L31:
	;
	if base.Ui32(v81) <= base.Ui32(l3) {
		goto L29
	} else {
		goto L32
	}
L32:
	;
	goto L26
L33:
	;
	v94 = int32(-5)
	if base.Ui32(v94) < base.Ui32(v89) {
		goto L36
	} else {
		goto L37
	}
L34:
	;
	if base.Ui32(l3) <= base.Ui32(int32(8192)) {
		goto L26
	} else {
		goto L35
	}
L35:
	;
	goto L29
L36:
	;
	v97 = v89
	goto L38
L37:
	;
	v97 = v94
	goto L38
L38:
	;
	v104 = *(*int32)(unsafe.Add(mBase, uint32((v97^int32(-1))<<(uint(int32(2))%32))+uint32(_c_F_quicklistReplaceEntry[1])))
	if base.Ui32(l3) <= base.Ui32(v104) {
		goto L26
	} else {
		goto L39
	}
L39:
	;
	goto L29
L40:
	;
	v110 = F_valkey_malloc(m, l3)
	mBase = m.M
	v111 = m.ExcPending
	if v111 != 0 {
		goto L16
	} else {
		goto L41
	}
L41:
	;
	v112 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v112)+12)) = l3
	*(*int32)(unsafe.Add(mBase, uint32(v112)+8)) = v110
	if l3 == int32(0) {
		goto L43
	} else {
		goto L44
	}
L42:
	;
	v119 = *(*int32)(unsafe.Add(mBase, uint32(v112)+16))
	if v119&int32(1048576) == int32(0) {
		goto L45
	} else {
		goto L46
	}
L43:
	;
	goto L42
L44:
	;
	v117 = F__emscripten_memcpy_bulkmem(m, v110, l2, l3)
	mBase = m.M
	goto L43
L45:
	;
	F___quicklistCompress(m, v15, v112)
	mBase = m.M
	v131 = m.ExcPending
	if v131 != 0 {
		goto L16
	} else {
		goto L49
	}
L46:
	;
	if v119&int32(196608) != int32(65536) {
		goto L1
	} else {
		goto L47
	}
L47:
	;
	v128 = F___quicklistCompressNode(m, v112)
	mBase = m.M
	v129 = m.ExcPending
	if v129 != 0 {
		goto L16
	} else {
		goto L48
	}
L48:
	;
	goto L1
L49:
	;
	goto L1
L50:
	;
	v150 = int32(0)
	v151 = *(*int32)(unsafe.Add(mBase, _c_F_quicklistReplaceEntry[0]))
	if v151 == v150 {
		goto L55
	} else {
		goto L56
	}
L51:
	;
	if v137 == int32(-1) {
		v149 = int32(0)
		goto L50
	} else {
		goto L52
	}
L52:
	;
	v147 = F__quicklistSplitNode(m, v16, v137, int32(1))
	mBase = m.M
	v148 = m.ExcPending
	if v148 != 0 {
		goto L16
	} else {
		goto L53
	}
L53:
	;
	v149 = v147
	goto L50
L54:
	;
	if v176 != 0 {
		goto L62
	} else {
		goto L63
	}
L55:
	;
	v155 = *(*int32)(unsafe.Add(mBase, uint32(v15)+16))
	v156 = int32(18)
	v159 = v155 << (uint(v156) % 32) >> (uint(v156) % 32)
	if v159 < int32(0) {
		goto L57
	} else {
		goto L58
	}
L56:
	;
	v176 = base.B2i32(base.Ui32(v151) <= base.Ui32(l3))
	goto L54
L57:
	;
	v164 = int32(-5)
	if base.Ui32(v164) < base.Ui32(v159) {
		goto L59
	} else {
		goto L60
	}
L58:
	;
	v176 = base.B2i32(base.Ui32(int32(8192)) < base.Ui32(l3))
	goto L54
L59:
	;
	v167 = v159
	goto L61
L60:
	;
	v167 = v164
	goto L61
L61:
	;
	v174 = *(*int32)(unsafe.Add(mBase, uint32((v167^int32(-1))<<(uint(int32(2))%32))+uint32(_c_F_quicklistReplaceEntry[1])))
	v176 = base.B2i32(base.Ui32(v174) < base.Ui32(l3))
	goto L54
L62:
	;
	v179 = int32(1)
	goto L64
L63:
	;
	v179 = int32(2)
	goto L64
L64:
	;
	v180 = F___quicklistCreateNode(m, v179, l2, l3)
	mBase = m.M
	v181 = m.ExcPending
	if v181 != 0 {
		goto L16
	} else {
		goto L65
	}
L65:
	;
	F___quicklistInsertNode(m, v15, v16, v180, int32(1))
	mBase = m.M
	v184 = m.ExcPending
	if v184 != 0 {
		goto L16
	} else {
		goto L66
	}
L66:
	;
	if v149 == int32(0) {
		goto L67
	} else {
		goto L68
	}
L67:
	;
	v190 = *(*int32)(unsafe.Add(mBase, uint32(v15)+8))
	v191 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v15)+8)) = v190 + v191
	v194 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v195 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v194)+16)))
	if v195 != v191 {
		goto L70
	} else {
		goto L71
	}
L68:
	;
	F___quicklistInsertNode(m, v15, v180, v149, int32(1))
	mBase = m.M
	v189 = m.ExcPending
	if v189 != 0 {
		goto L16
	} else {
		goto L69
	}
L69:
	;
	goto L67
L70:
	;
	v200 = *(*int32)(unsafe.Add(mBase, uint32(v194)+8))
	v202 = F_lpSeek(m, v200, int32(-1))
	mBase = m.M
	v203 = m.ExcPending
	if v203 != 0 {
		goto L16
	} else {
		goto L73
	}
L71:
	;
	F___quicklistDelNode(m, v15, v194)
	mBase = m.M
	v199 = m.ExcPending
	if v199 != 0 {
		goto L16
	} else {
		goto L72
	}
L72:
	;
	goto L1
L73:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+12)) = v202
	v205 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	F_quicklistDelIndex(m, v15, v205, v13+int32(12))
	mBase = m.M
	v209 = m.ExcPending
	if v209 != 0 {
		goto L16
	} else {
		goto L74
	}
L74:
	;
	v210 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v211 = *(*int32)(unsafe.Add(mBase, uint32(v210)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v210)+16)) = v211 & int32(-4194305)
	v215 = F__quicklistMergeNodes(m, v15, v180)
	mBase = m.M
	v216 = m.ExcPending
	if v216 != 0 {
		goto L16
	} else {
		goto L77
	}
L75:
	;
	v230 = *(*int32)(unsafe.Add(mBase, uint32(v215)))
	v231 = *(*int32)(unsafe.Add(mBase, uint32(v230)+16))
	if v231&int32(1048576) == int32(0) {
		goto L83
	} else {
		goto L84
	}
L76:
	;
	F___quicklistCompress(m, v15, v215)
	mBase = m.M
	v229 = m.ExcPending
	if v229 != 0 {
		goto L16
	} else {
		goto L81
	}
L77:
	;
	v217 = *(*int32)(unsafe.Add(mBase, uint32(v215)+16))
	if v217&int32(1048576) == int32(0) {
		goto L76
	} else {
		goto L78
	}
L78:
	;
	if v217&int32(196608) != int32(65536) {
		goto L75
	} else {
		goto L79
	}
L79:
	;
	v226 = F___quicklistCompressNode(m, v215)
	mBase = m.M
	v227 = m.ExcPending
	if v227 != 0 {
		goto L16
	} else {
		goto L80
	}
L80:
	;
	goto L75
L81:
	;
	goto L75
L82:
	;
	v244 = *(*int32)(unsafe.Add(mBase, uint32(v215)+4))
	if v244 == int32(0) {
		goto L1
	} else {
		goto L88
	}
L83:
	;
	F___quicklistCompress(m, v15, v230)
	mBase = m.M
	v243 = m.ExcPending
	if v243 != 0 {
		goto L16
	} else {
		goto L87
	}
L84:
	;
	if v231&int32(196608) != int32(65536) {
		goto L82
	} else {
		goto L85
	}
L85:
	;
	v240 = F___quicklistCompressNode(m, v230)
	mBase = m.M
	v241 = m.ExcPending
	if v241 != 0 {
		goto L16
	} else {
		goto L86
	}
L86:
	;
	goto L82
L87:
	;
	goto L82
L88:
	;
	v247 = *(*int32)(unsafe.Add(mBase, uint32(v244)+16))
	if v247&int32(1048576) == int32(0) {
		goto L89
	} else {
		goto L90
	}
L89:
	;
	F___quicklistCompress(m, v15, v244)
	mBase = m.M
	v259 = m.ExcPending
	if v259 != 0 {
		goto L16
	} else {
		goto L93
	}
L90:
	;
	if v247&int32(196608) != int32(65536) {
		goto L1
	} else {
		goto L91
	}
L91:
	;
	v256 = F___quicklistCompressNode(m, v244)
	mBase = m.M
	v257 = m.ExcPending
	if v257 != 0 {
		goto L16
	} else {
		goto L92
	}
L92:
	;
	goto L1
L93:
	;
	goto L1
L94:
	;
	v264 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	F___quicklistDelNode(m, v15, v264)
	mBase = m.M
	v266 = m.ExcPending
	if v266 != 0 {
		goto L16
	} else {
		goto L95
	}
L95:
	;
	goto L1
}
func F_quicklistSetDirection(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = l1
	return
}
func F_quicklistSetPackedThreshold(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	if base.Ui32(int32(-1048576)) < base.Ui32(l0) {
		v9 = int32(0)
	} else {
		*(*int32)(unsafe.Add(mBase, _c_F_quicklistSetPackedThreshold[0])) = l0
		v9 = int32(1)
	}
	return v9
}
