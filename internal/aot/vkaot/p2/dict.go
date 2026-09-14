package p2

import (
	base "github.com/shibukawa/valkeymem/internal/aot/vkaot/base"
	"unsafe"
)

func F_dictCStrCaseHash(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v16 int32
	_ = v16
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v56 int32
	_ = v56
	var v58 int64
	_ = v58
	if l0&int32(3) == int32(0) {
		v23 = l0
		goto L3
	} else {
		goto L4
	}
L1:
	;
	v58 = F_siphash_nocase(m, l0, v56, int32(_a765))
	mBase = m.M
	goto L17
L2:
	;
	v56 = v48 - l0
	goto L1
L3:
	;
	v27 = v23
	goto L11
L4:
	;
	v9 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v9 != 0 {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	v12 = l0
	goto L7
L6:
	;
	v56 = l0 - l0
	goto L1
L7:
	;
	v16 = v12 + int32(1)
	if v16&int32(3) == int32(0) {
		v23 = v16
		goto L3
	} else {
		goto L9
	}
L9:
	;
	v21 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16))))
	if v21 != 0 {
		v12 = v16
		goto L7
	} else {
		goto L10
	}
L10:
	;
	v48 = v16
	goto L2
L11:
	;
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v27)))
	v36 = int32(-2139062144)
	if (int32(16843008)-v33|v33)&v36 == v36 {
		v27 = v27 + int32(4)
		goto L11
	} else {
		goto L13
	}
L12:
	;
	v42 = v27
	goto L14
L13:
	;
	goto L12
L14:
	;
	v46 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v42))))
	if v46 != 0 {
		v42 = v42 + int32(1)
		goto L14
	} else {
		goto L16
	}
L15:
	;
	v48 = v42
	goto L2
L16:
	;
	goto L15
L17:
	;
	return v58
}
func F_dictCreate(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v4 != 0 {
		v11 = m.T0[v4].(func(*base.Module, int32) int32)(m, int32(0))
		mBase = m.M
		v12 = m.ExcPending
		if v12 != 0 {
			return int32(0)
		} else {
			v15 = F_valkey_malloc(m, v11+int32(32))
			mBase = m.M
			v16 = m.ExcPending
			if v16 != 0 {
				return int32(0)
			} else {
				if v11 == int32(0) {
					v24 = v15
				} else {
					v23 = F__emscripten_memset_bulkmem(m, v15+int32(32), base.I32_extend8_s(int32(0)), v11)
					mBase = m.M
					v24 = v15
				}
				*(*int64)(unsafe.Add(mBase, uint32(v24)+4)) = int64(0)
				v28 = int32(0)
				*(*int32)(unsafe.Add(mBase, uint32(v24)+12)) = v28
				*(*uint16)(unsafe.Add(mBase, uint32(v24)+28)) = uint16(v28)
				*(*int32)(unsafe.Add(mBase, uint32(v24)+24)) = int32(-65536)
				*(*int64)(unsafe.Add(mBase, uint32(v24)+16)) = int64(-4294967296)
				*(*int32)(unsafe.Add(mBase, uint32(v24))) = l0
				return v24
			}
		}
	} else {
		v6 = F_valkey_malloc(m, int32(32))
		mBase = m.M
		v9 = m.ExcPending
		if v9 != 0 {
			return int32(0)
		} else {
			v24 = v6
			*(*int64)(unsafe.Add(mBase, uint32(v24)+4)) = int64(0)
			v28 = int32(0)
			*(*int32)(unsafe.Add(mBase, uint32(v24)+12)) = v28
			*(*uint16)(unsafe.Add(mBase, uint32(v24)+28)) = uint16(v28)
			*(*int32)(unsafe.Add(mBase, uint32(v24)+24)) = int32(-65536)
			*(*int64)(unsafe.Add(mBase, uint32(v24)+16)) = int64(-4294967296)
			*(*int32)(unsafe.Add(mBase, uint32(v24))) = l0
			return v24
		}
	}
}
func F_dictDelete(m *base.Module, l0 int32, l1 int32) int32 {
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	v4 = F_dictGenericDelete(m, l0, l1, int32(0))
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		return base.B2i32(v4 == int32(0))
	}
}
func F_dictFind(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int64
	_ = v15
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v28 int64
	_ = v28
	var v29 int64
	_ = v29
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v62 int64
	_ = v62
	var v68 int32
	_ = v68
	var v74 int32
	_ = v74
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v90 int32
	_ = v90
	var v98 int32
	_ = v98
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v118 int64
	_ = v118
	var v124 int32
	_ = v124
	var v130 int32
	_ = v130
	var v134 int32
	_ = v134
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v146 int32
	_ = v146
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v8 == int32(0)-v10 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
	v15 = m.T0[v14].(func(*base.Module, int32) int64)(m, l1)
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	return int32(0)
L4:
	;
	v20 = int32(-1)
	v21 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+26)))
	if v21 == int32(255) {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	v28 = int64(0)
	goto L7
L6:
	;
	v28 = base.I64_extend_i32_u(v20<<(uint(v21)%32) ^ v20)
	goto L7
L7:
	;
	v29 = v15 & v28
	v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v30 == int32(-1) {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v49 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if base.I32_wrap_i64(v29) < v49 {
		goto L16
	} else {
		goto L17
	}
L9:
	;
	v33 = base.I32_wrap_i64(v29)
	if v33 < v30 {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	v44 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+24)))
	if v44 != 0 {
		goto L8
	} else {
		goto L14
	}
L11:
	;
	v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v35+v33<<(uint(int32(2))%32))))
	if v39 == int32(0) {
		goto L10
	} else {
		goto L12
	}
L12:
	;
	F_dictBucketRehash(m, l0, v29)
	mBase = m.M
	v43 = m.ExcPending
	if v43 != 0 {
		goto L3
	} else {
		goto L13
	}
L13:
	;
	goto L8
L14:
	;
	v46 = F_dictRehash(m, l0, int32(1))
	mBase = m.M
	v47 = m.ExcPending
	if v47 != 0 {
		goto L3
	} else {
		goto L15
	}
L15:
	;
	goto L8
L16:
	;
	v108 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v110 = int32(-1)
	v111 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+27)))
	if v111 == int32(255) {
		goto L33
	} else {
		goto L34
	}
L17:
	;
	v52 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v54 = int32(-1)
	v55 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+26)))
	if v55 == int32(255) {
		goto L19
	} else {
		goto L20
	}
L18:
	;
	v98 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v98 == int32(-1) {
		goto L1
	} else {
		goto L32
	}
L19:
	;
	v62 = int64(0)
	goto L21
L20:
	;
	v62 = base.I64_extend_i32_u(v54<<(uint(v55)%32) ^ v54)
	goto L21
L21:
	;
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v52+base.I32_wrap_i64(v62&v15)<<(uint(int32(2))%32))))
	if v68 == int32(0) {
		goto L18
	} else {
		goto L22
	}
L22:
	;
	v74 = v68
	goto L23
L23:
	;
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v74)))
	if l1 != v78 {
		goto L25
	} else {
		goto L26
	}
L24:
	;
	goto L18
L25:
	;
	v81 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v82 = *(*int32)(unsafe.Add(mBase, uint32(v81)+8))
	if v82 == int32(0) {
		goto L27
	} else {
		goto L28
	}
L26:
	;
	return v74
L27:
	;
	v90 = *(*int32)(unsafe.Add(mBase, uint32(v74)+16))
	if v90 != 0 {
		v74 = v90
		goto L23
	} else {
		goto L31
	}
L28:
	;
	v85 = m.T0[v82].(func(*base.Module, int32, int32) int32)(m, l1, v78)
	mBase = m.M
	v86 = m.ExcPending
	if v86 != 0 {
		goto L3
	} else {
		goto L29
	}
L29:
	;
	if v85 == int32(0) {
		goto L27
	} else {
		goto L30
	}
L30:
	;
	return v74
L31:
	;
	goto L24
L32:
	;
	goto L16
L33:
	;
	v118 = int64(0)
	goto L35
L34:
	;
	v118 = base.I64_extend_i32_u(v110<<(uint(v111)%32) ^ v110)
	goto L35
L35:
	;
	v124 = *(*int32)(unsafe.Add(mBase, uint32(v108+base.I32_wrap_i64(v118&v15)<<(uint(int32(2))%32))))
	if v124 == int32(0) {
		goto L1
	} else {
		goto L36
	}
L36:
	;
	v130 = v124
	goto L37
L37:
	;
	v134 = *(*int32)(unsafe.Add(mBase, uint32(v130)))
	if l1 != v134 {
		goto L39
	} else {
		goto L40
	}
L38:
	;
	goto L1
L39:
	;
	v137 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v138 = *(*int32)(unsafe.Add(mBase, uint32(v137)+8))
	if v138 == int32(0) {
		goto L41
	} else {
		goto L42
	}
L40:
	;
	return v130
L41:
	;
	v146 = *(*int32)(unsafe.Add(mBase, uint32(v130)+16))
	if v146 != 0 {
		v130 = v146
		goto L37
	} else {
		goto L45
	}
L42:
	;
	v141 = m.T0[v138].(func(*base.Module, int32, int32) int32)(m, l1, v134)
	mBase = m.M
	v142 = m.ExcPending
	if v142 != 0 {
		goto L3
	} else {
		goto L43
	}
L43:
	;
	if v141 == int32(0) {
		goto L41
	} else {
		goto L44
	}
L44:
	;
	return v130
L45:
	;
	goto L38
}
func F_dictGetUnsignedIntegerVal(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int64
	_ = v2
	v2 = *(*int64)(unsafe.Add(mBase, uint32(l0)+8))
	return v2
}
func F_dictGetVal(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	return v2
}
func F_dictInstancesValDestructor(m *base.Module, l0 int32) {
	var v3 int32
	_ = v3
	F_releaseSentinelValkeyInstance(m, l0)
	v3 = m.ExcPending
	if v3 != 0 {
		return
	} else {
		return
	}
}
func F_dictListDestructor(m *base.Module, l0 int32) {
	var v3 int32
	_ = v3
	F_listRelease(m, l0)
	v3 = m.ExcPending
	if v3 != 0 {
		return
	} else {
		return
	}
}
func F_dictObjKeyCompare(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v12 int32
	_ = v12
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	var v100 int32
	_ = v100
	var v115 int32
	_ = v115
	var v118 int32
	_ = v118
	v3 = int32(0)
	v6 = F_objectGetVal(m, l0)
	mBase = m.M
	v7 = F_objectGetVal(m, l1)
	mBase = m.M
	v12 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6+int32(-1)))))
	switch v12 & int32(7) {
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
		v29 = v3
		goto L1
	}
L1:
	;
	v32 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7+int32(-1)))))
	switch v32 & int32(7) {
	case 0:
		goto L12
	case 1:
		goto L11
	case 2:
		goto L10
	case 3:
		goto L9
	case 4:
		goto L8
	default:
		v49 = v3
		goto L7
	}
L2:
	;
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v6+int32(-17))))
	v29 = v28
	goto L1
L3:
	;
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v6+int32(-9))))
	v29 = v25
	goto L1
L4:
	;
	v22 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v6+int32(-5)))))
	v29 = v22
	goto L1
L5:
	;
	v19 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6+int32(-3)))))
	v29 = v19
	goto L1
L6:
	;
	v29 = int32(base.Ui32(v12) >> (uint(int32(3)) % 32))
	goto L1
L7:
	;
	if v29 != v49 {
		v118 = int32(0)
		goto L13
	} else {
		goto L14
	}
L8:
	;
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v7+int32(-17))))
	v49 = v48
	goto L7
L9:
	;
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v7+int32(-9))))
	v49 = v45
	goto L7
L10:
	;
	v42 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v7+int32(-5)))))
	v49 = v42
	goto L7
L11:
	;
	v39 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7+int32(-3)))))
	v49 = v39
	goto L7
L12:
	;
	v49 = int32(base.Ui32(v32) >> (uint(int32(3)) % 32))
	goto L7
L13:
	;
	return v118
L14:
	;
	if base.Ui32(v29) < base.Ui32(int32(4)) {
		v75 = v6
		v76 = v7
		v77 = v29
		goto L18
	} else {
		goto L19
	}
L15:
	;
	v118 = base.B2i32(v115 == int32(0))
	goto L13
L16:
	;
	v115 = int32(0)
	goto L15
L17:
	;
	v87 = v82
	v88 = v83
	v89 = v84
	goto L27
L18:
	;
	if v77 == int32(0) {
		goto L16
	} else {
		goto L25
	}
L19:
	;
	if (v7|v6)&int32(3) != 0 {
		v82 = v6
		v83 = v7
		v84 = v29
		goto L17
	} else {
		goto L20
	}
L20:
	;
	v59 = v6
	v60 = v7
	v61 = v29
	goto L21
L21:
	;
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v59)))
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v60)))
	if v64 != v65 {
		v82 = v59
		v83 = v60
		v84 = v61
		goto L17
	} else {
		goto L23
	}
L22:
	;
	v75 = v70
	v76 = v68
	v77 = v72
	goto L18
L23:
	;
	v67 = int32(4)
	v68 = v60 + v67
	v70 = v59 + v67
	v72 = v61 + int32(-4)
	if base.Ui32(int32(3)) < base.Ui32(v72) {
		v59 = v70
		v60 = v68
		v61 = v72
		goto L21
	} else {
		goto L24
	}
L24:
	;
	goto L22
L25:
	;
	v82 = v75
	v83 = v76
	v84 = v77
	goto L17
L26:
	;
	v115 = v92 - v93
	goto L15
L27:
	;
	v92 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v87))))
	v93 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v88))))
	if v92 != v93 {
		goto L26
	} else {
		goto L29
	}
L29:
	;
	v95 = int32(1)
	v100 = v89 + int32(-1)
	if v100 == int32(0) {
		goto L16
	} else {
		goto L30
	}
L30:
	;
	v87 = v87 + v95
	v88 = v88 + v95
	v89 = v100
	goto L27
}
func F_dictRelease(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v3 == int32(-1) {
		v13 = int32(0)
		F_dictClear(m, l0, v13, v13)
		mBase = m.M
		v16 = m.ExcPending
		if v16 != 0 {
			return
		} else {
			F_dictClear(m, l0, int32(1), int32(0))
			mBase = m.M
			v20 = m.ExcPending
			if v20 != 0 {
				return
			} else {
				F_valkey_free(m, l0)
				mBase = m.M
				v22 = m.ExcPending
				if v22 != 0 {
					return
				} else {
					return
				}
			}
		}
	} else {
		v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		v7 = *(*int32)(unsafe.Add(mBase, uint32(v6)+28))
		if v7 == int32(0) {
			v13 = int32(0)
			F_dictClear(m, l0, v13, v13)
			mBase = m.M
			v16 = m.ExcPending
			if v16 != 0 {
				return
			} else {
				F_dictClear(m, l0, int32(1), int32(0))
				mBase = m.M
				v20 = m.ExcPending
				if v20 != 0 {
					return
				} else {
					F_valkey_free(m, l0)
					mBase = m.M
					v22 = m.ExcPending
					if v22 != 0 {
						return
					} else {
						return
					}
				}
			}
		} else {
			m.T0[v7].(func(*base.Module, int32))(m, l0)
			mBase = m.M
			v11 = m.ExcPending
			if v11 != 0 {
				return
			} else {
				v13 = int32(0)
				F_dictClear(m, l0, v13, v13)
				mBase = m.M
				v16 = m.ExcPending
				if v16 != 0 {
					return
				} else {
					F_dictClear(m, l0, int32(1), int32(0))
					mBase = m.M
					v20 = m.ExcPending
					if v20 != 0 {
						return
					} else {
						F_valkey_free(m, l0)
						mBase = m.M
						v22 = m.ExcPending
						if v22 != 0 {
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
func F_dictReleaseIterator(m *base.Module, l0 int32) {
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	F_dictResetIterator(m, l0)
	v3 = m.ExcPending
	if v3 != 0 {
		return
	} else {
		F_valkey_free(m, l0)
		v5 = m.ExcPending
		if v5 != 0 {
			return
		} else {
			return
		}
	}
}
func F_dictResizeWithOptionalCheck(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v85 int32
	_ = v85
	var v91 int32
	_ = v91
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v123 int32
	_ = v123
	var v125 int32
	_ = v125
	var v135 int32
	_ = v135
	var v141 int32
	_ = v141
	if l2 == int32(0) {
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(l2))) = int32(0)
	}
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v11 != int32(-1) {
		F__serverAssert(m, int32(_a763), int32(_a764), int32(198))
		mBase = m.M
		v141 = m.ExcPending
		if v141 != 0 {
			return int32(0)
		} else {
			F_abort(m)
			mBase = m.M
			base.Wasm_trap_unreachable()
			for {
			}
		}
	} else {
		if base.Ui32(int32(5)) <= base.Ui32(l1) {
			if base.Ui32(l1) <= base.Ui32(int32(2147483646)) {
				v25 = int32(32) - base.I32_clz(l1+int32(-1))
			} else {
				v25 = int32(31)
			}
		} else {
			v25 = int32(2)
		}
		v26 = int32(1)
		v28 = v26 << (uint(v25) % 32)
		if base.Ui32(v28) < base.Ui32(l1) {
			v135 = v26
			return v135
		} else {
			v31 = int32(4) << (uint(v25) % 32)
			if base.Ui32(v31) < base.Ui32(v28) {
				v135 = v26
				return v135
			} else {
				v35 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+26)))
				if v25&int32(255) == v35 {
					v135 = v26
					return v135
				} else {
					if l2 == int32(0) {
						v91 = F_valkey_calloc(m, v31)
						mBase = m.M
						v94 = m.ExcPending
						if v94 != 0 {
							return int32(0)
						} else {
							v95 = v91
							*(*uint8)(unsafe.Add(mBase, uint32(l0)+27)) = uint8(v25)
							*(*int64)(unsafe.Add(mBase, uint32(l0)+16)) = int64(0)
							*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v95
							v100 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
							v101 = *(*int32)(unsafe.Add(mBase, uint32(v100)+24))
							if v101 == int32(0) {
								v106 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
								if v106 == int32(0) {
									v112 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
									v113 = *(*int32)(unsafe.Add(mBase, uint32(v112)+28))
									if v113 == int32(0) {
										v119 = v106
										if v119 == int32(0) {
											*(*uint8)(unsafe.Add(mBase, uint32(l0)+26)) = uint8(v25)
											v125 = int32(255)
											*(*uint8)(unsafe.Add(mBase, uint32(l0)+27)) = uint8(v125)
											*(*int64)(unsafe.Add(mBase, uint32(l0)+8)) = int64(0)
											*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v95
											*(*int64)(unsafe.Add(mBase, uint32(l0)+16)) = int64(-4294967296)
											v135 = int32(0)
											return v135
										} else {
											F_valkey_free(m, v119)
											mBase = m.M
											v123 = m.ExcPending
											if v123 != 0 {
												return int32(0)
											} else {
												*(*uint8)(unsafe.Add(mBase, uint32(l0)+26)) = uint8(v25)
												v125 = int32(255)
												*(*uint8)(unsafe.Add(mBase, uint32(l0)+27)) = uint8(v125)
												*(*int64)(unsafe.Add(mBase, uint32(l0)+8)) = int64(0)
												*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v95
												*(*int64)(unsafe.Add(mBase, uint32(l0)+16)) = int64(-4294967296)
												v135 = int32(0)
												return v135
											}
										}
									} else {
										m.T0[v113].(func(*base.Module, int32))(m, l0)
										mBase = m.M
										v117 = m.ExcPending
										if v117 != 0 {
											return int32(0)
										} else {
											v118 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
											v119 = v118
											if v119 == int32(0) {
												*(*uint8)(unsafe.Add(mBase, uint32(l0)+26)) = uint8(v25)
												v125 = int32(255)
												*(*uint8)(unsafe.Add(mBase, uint32(l0)+27)) = uint8(v125)
												*(*int64)(unsafe.Add(mBase, uint32(l0)+8)) = int64(0)
												*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v95
												*(*int64)(unsafe.Add(mBase, uint32(l0)+16)) = int64(-4294967296)
												v135 = int32(0)
												return v135
											} else {
												F_valkey_free(m, v119)
												mBase = m.M
												v123 = m.ExcPending
												if v123 != 0 {
													return int32(0)
												} else {
													*(*uint8)(unsafe.Add(mBase, uint32(l0)+26)) = uint8(v25)
													v125 = int32(255)
													*(*uint8)(unsafe.Add(mBase, uint32(l0)+27)) = uint8(v125)
													*(*int64)(unsafe.Add(mBase, uint32(l0)+8)) = int64(0)
													*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v95
													*(*int64)(unsafe.Add(mBase, uint32(l0)+16)) = int64(-4294967296)
													v135 = int32(0)
													return v135
												}
											}
										}
									}
								} else {
									v110 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
									if v110 != 0 {
										v135 = int32(0)
										return v135
									} else {
										v112 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
										v113 = *(*int32)(unsafe.Add(mBase, uint32(v112)+28))
										if v113 == int32(0) {
											v119 = v106
											if v119 == int32(0) {
												*(*uint8)(unsafe.Add(mBase, uint32(l0)+26)) = uint8(v25)
												v125 = int32(255)
												*(*uint8)(unsafe.Add(mBase, uint32(l0)+27)) = uint8(v125)
												*(*int64)(unsafe.Add(mBase, uint32(l0)+8)) = int64(0)
												*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v95
												*(*int64)(unsafe.Add(mBase, uint32(l0)+16)) = int64(-4294967296)
												v135 = int32(0)
												return v135
											} else {
												F_valkey_free(m, v119)
												mBase = m.M
												v123 = m.ExcPending
												if v123 != 0 {
													return int32(0)
												} else {
													*(*uint8)(unsafe.Add(mBase, uint32(l0)+26)) = uint8(v25)
													v125 = int32(255)
													*(*uint8)(unsafe.Add(mBase, uint32(l0)+27)) = uint8(v125)
													*(*int64)(unsafe.Add(mBase, uint32(l0)+8)) = int64(0)
													*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v95
													*(*int64)(unsafe.Add(mBase, uint32(l0)+16)) = int64(-4294967296)
													v135 = int32(0)
													return v135
												}
											}
										} else {
											m.T0[v113].(func(*base.Module, int32))(m, l0)
											mBase = m.M
											v117 = m.ExcPending
											if v117 != 0 {
												return int32(0)
											} else {
												v118 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
												v119 = v118
												if v119 == int32(0) {
													*(*uint8)(unsafe.Add(mBase, uint32(l0)+26)) = uint8(v25)
													v125 = int32(255)
													*(*uint8)(unsafe.Add(mBase, uint32(l0)+27)) = uint8(v125)
													*(*int64)(unsafe.Add(mBase, uint32(l0)+8)) = int64(0)
													*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v95
													*(*int64)(unsafe.Add(mBase, uint32(l0)+16)) = int64(-4294967296)
													v135 = int32(0)
													return v135
												} else {
													F_valkey_free(m, v119)
													mBase = m.M
													v123 = m.ExcPending
													if v123 != 0 {
														return int32(0)
													} else {
														*(*uint8)(unsafe.Add(mBase, uint32(l0)+26)) = uint8(v25)
														v125 = int32(255)
														*(*uint8)(unsafe.Add(mBase, uint32(l0)+27)) = uint8(v125)
														*(*int64)(unsafe.Add(mBase, uint32(l0)+8)) = int64(0)
														*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v95
														*(*int64)(unsafe.Add(mBase, uint32(l0)+16)) = int64(-4294967296)
														v135 = int32(0)
														return v135
													}
												}
											}
										}
									}
								}
							} else {
								m.T0[v101].(func(*base.Module, int32))(m, l0)
								mBase = m.M
								v105 = m.ExcPending
								if v105 != 0 {
									return int32(0)
								} else {
									v106 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
									if v106 == int32(0) {
										v112 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
										v113 = *(*int32)(unsafe.Add(mBase, uint32(v112)+28))
										if v113 == int32(0) {
											v119 = v106
											if v119 == int32(0) {
												*(*uint8)(unsafe.Add(mBase, uint32(l0)+26)) = uint8(v25)
												v125 = int32(255)
												*(*uint8)(unsafe.Add(mBase, uint32(l0)+27)) = uint8(v125)
												*(*int64)(unsafe.Add(mBase, uint32(l0)+8)) = int64(0)
												*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v95
												*(*int64)(unsafe.Add(mBase, uint32(l0)+16)) = int64(-4294967296)
												v135 = int32(0)
												return v135
											} else {
												F_valkey_free(m, v119)
												mBase = m.M
												v123 = m.ExcPending
												if v123 != 0 {
													return int32(0)
												} else {
													*(*uint8)(unsafe.Add(mBase, uint32(l0)+26)) = uint8(v25)
													v125 = int32(255)
													*(*uint8)(unsafe.Add(mBase, uint32(l0)+27)) = uint8(v125)
													*(*int64)(unsafe.Add(mBase, uint32(l0)+8)) = int64(0)
													*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v95
													*(*int64)(unsafe.Add(mBase, uint32(l0)+16)) = int64(-4294967296)
													v135 = int32(0)
													return v135
												}
											}
										} else {
											m.T0[v113].(func(*base.Module, int32))(m, l0)
											mBase = m.M
											v117 = m.ExcPending
											if v117 != 0 {
												return int32(0)
											} else {
												v118 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
												v119 = v118
												if v119 == int32(0) {
													*(*uint8)(unsafe.Add(mBase, uint32(l0)+26)) = uint8(v25)
													v125 = int32(255)
													*(*uint8)(unsafe.Add(mBase, uint32(l0)+27)) = uint8(v125)
													*(*int64)(unsafe.Add(mBase, uint32(l0)+8)) = int64(0)
													*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v95
													*(*int64)(unsafe.Add(mBase, uint32(l0)+16)) = int64(-4294967296)
													v135 = int32(0)
													return v135
												} else {
													F_valkey_free(m, v119)
													mBase = m.M
													v123 = m.ExcPending
													if v123 != 0 {
														return int32(0)
													} else {
														*(*uint8)(unsafe.Add(mBase, uint32(l0)+26)) = uint8(v25)
														v125 = int32(255)
														*(*uint8)(unsafe.Add(mBase, uint32(l0)+27)) = uint8(v125)
														*(*int64)(unsafe.Add(mBase, uint32(l0)+8)) = int64(0)
														*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v95
														*(*int64)(unsafe.Add(mBase, uint32(l0)+16)) = int64(-4294967296)
														v135 = int32(0)
														return v135
													}
												}
											}
										}
									} else {
										v110 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
										if v110 != 0 {
											v135 = int32(0)
											return v135
										} else {
											v112 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
											v113 = *(*int32)(unsafe.Add(mBase, uint32(v112)+28))
											if v113 == int32(0) {
												v119 = v106
												if v119 == int32(0) {
													*(*uint8)(unsafe.Add(mBase, uint32(l0)+26)) = uint8(v25)
													v125 = int32(255)
													*(*uint8)(unsafe.Add(mBase, uint32(l0)+27)) = uint8(v125)
													*(*int64)(unsafe.Add(mBase, uint32(l0)+8)) = int64(0)
													*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v95
													*(*int64)(unsafe.Add(mBase, uint32(l0)+16)) = int64(-4294967296)
													v135 = int32(0)
													return v135
												} else {
													F_valkey_free(m, v119)
													mBase = m.M
													v123 = m.ExcPending
													if v123 != 0 {
														return int32(0)
													} else {
														*(*uint8)(unsafe.Add(mBase, uint32(l0)+26)) = uint8(v25)
														v125 = int32(255)
														*(*uint8)(unsafe.Add(mBase, uint32(l0)+27)) = uint8(v125)
														*(*int64)(unsafe.Add(mBase, uint32(l0)+8)) = int64(0)
														*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v95
														*(*int64)(unsafe.Add(mBase, uint32(l0)+16)) = int64(-4294967296)
														v135 = int32(0)
														return v135
													}
												}
											} else {
												m.T0[v113].(func(*base.Module, int32))(m, l0)
												mBase = m.M
												v117 = m.ExcPending
												if v117 != 0 {
													return int32(0)
												} else {
													v118 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
													v119 = v118
													if v119 == int32(0) {
														*(*uint8)(unsafe.Add(mBase, uint32(l0)+26)) = uint8(v25)
														v125 = int32(255)
														*(*uint8)(unsafe.Add(mBase, uint32(l0)+27)) = uint8(v125)
														*(*int64)(unsafe.Add(mBase, uint32(l0)+8)) = int64(0)
														*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v95
														*(*int64)(unsafe.Add(mBase, uint32(l0)+16)) = int64(-4294967296)
														v135 = int32(0)
														return v135
													} else {
														F_valkey_free(m, v119)
														mBase = m.M
														v123 = m.ExcPending
														if v123 != 0 {
															return int32(0)
														} else {
															*(*uint8)(unsafe.Add(mBase, uint32(l0)+26)) = uint8(v25)
															v125 = int32(255)
															*(*uint8)(unsafe.Add(mBase, uint32(l0)+27)) = uint8(v125)
															*(*int64)(unsafe.Add(mBase, uint32(l0)+8)) = int64(0)
															*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v95
															*(*int64)(unsafe.Add(mBase, uint32(l0)+16)) = int64(-4294967296)
															v135 = int32(0)
															return v135
														}
													}
												}
											}
										}
									}
								}
							}
						}
					} else {
						v39 = int32(0)
						if base.Ui32(int32(2147483646)) < base.Ui32(v31) {
							v85 = v39
						} else {
							if v31 != 0 {
								v47 = v31
							} else {
								v47 = int32(4)
							}
							v49 = v47 + int32(8)
							v50 = F_emscripten_builtin_calloc(m, int32(1), v49)
							mBase = m.M
							if v50 == int32(0) {
								v85 = v39
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v50))) = v47
								v55 = *(*int32)(unsafe.Add(mBase, _consts[411]))
								if v55 != int32(-1) {
									v66 = v55
								} else {
									v58 = int32(0)
									v60 = *(*int32)(unsafe.Add(mBase, _consts[281]))
									*(*int32)(unsafe.Add(mBase, _consts[411])) = v60
									*(*int32)(unsafe.Add(mBase, _consts[281])) = v60 + int32(1)
									v66 = v60
								}
								if v66 < int32(260) {
									v75 = v66 << (uint(int32(2)) % 32)
									v78 = *(*int32)(unsafe.Add(mBase, uint32(v75)+uint32(_consts[285])))
									*(*int32)(unsafe.Add(mBase, uint32(v75)+uint32(_consts[285]))) = v78 + v49
								} else {
									v69 = int32(0)
									v71 = *(*int32)(unsafe.Add(mBase, _consts[286]))
									*(*int32)(unsafe.Add(mBase, _consts[286])) = v71 + v49
								}
								v85 = v50 + int32(8)
							}
						}
						*(*int32)(unsafe.Add(mBase, uint32(l2))) = base.B2i32(v85 == int32(0))
						if v85 != 0 {
							v95 = v85
							*(*uint8)(unsafe.Add(mBase, uint32(l0)+27)) = uint8(v25)
							*(*int64)(unsafe.Add(mBase, uint32(l0)+16)) = int64(0)
							*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v95
							v100 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
							v101 = *(*int32)(unsafe.Add(mBase, uint32(v100)+24))
							if v101 == int32(0) {
								v106 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
								if v106 == int32(0) {
									v112 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
									v113 = *(*int32)(unsafe.Add(mBase, uint32(v112)+28))
									if v113 == int32(0) {
										v119 = v106
										if v119 == int32(0) {
											*(*uint8)(unsafe.Add(mBase, uint32(l0)+26)) = uint8(v25)
											v125 = int32(255)
											*(*uint8)(unsafe.Add(mBase, uint32(l0)+27)) = uint8(v125)
											*(*int64)(unsafe.Add(mBase, uint32(l0)+8)) = int64(0)
											*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v95
											*(*int64)(unsafe.Add(mBase, uint32(l0)+16)) = int64(-4294967296)
											v135 = int32(0)
											return v135
										} else {
											F_valkey_free(m, v119)
											mBase = m.M
											v123 = m.ExcPending
											if v123 != 0 {
												return int32(0)
											} else {
												*(*uint8)(unsafe.Add(mBase, uint32(l0)+26)) = uint8(v25)
												v125 = int32(255)
												*(*uint8)(unsafe.Add(mBase, uint32(l0)+27)) = uint8(v125)
												*(*int64)(unsafe.Add(mBase, uint32(l0)+8)) = int64(0)
												*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v95
												*(*int64)(unsafe.Add(mBase, uint32(l0)+16)) = int64(-4294967296)
												v135 = int32(0)
												return v135
											}
										}
									} else {
										m.T0[v113].(func(*base.Module, int32))(m, l0)
										mBase = m.M
										v117 = m.ExcPending
										if v117 != 0 {
											return int32(0)
										} else {
											v118 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
											v119 = v118
											if v119 == int32(0) {
												*(*uint8)(unsafe.Add(mBase, uint32(l0)+26)) = uint8(v25)
												v125 = int32(255)
												*(*uint8)(unsafe.Add(mBase, uint32(l0)+27)) = uint8(v125)
												*(*int64)(unsafe.Add(mBase, uint32(l0)+8)) = int64(0)
												*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v95
												*(*int64)(unsafe.Add(mBase, uint32(l0)+16)) = int64(-4294967296)
												v135 = int32(0)
												return v135
											} else {
												F_valkey_free(m, v119)
												mBase = m.M
												v123 = m.ExcPending
												if v123 != 0 {
													return int32(0)
												} else {
													*(*uint8)(unsafe.Add(mBase, uint32(l0)+26)) = uint8(v25)
													v125 = int32(255)
													*(*uint8)(unsafe.Add(mBase, uint32(l0)+27)) = uint8(v125)
													*(*int64)(unsafe.Add(mBase, uint32(l0)+8)) = int64(0)
													*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v95
													*(*int64)(unsafe.Add(mBase, uint32(l0)+16)) = int64(-4294967296)
													v135 = int32(0)
													return v135
												}
											}
										}
									}
								} else {
									v110 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
									if v110 != 0 {
										v135 = int32(0)
										return v135
									} else {
										v112 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
										v113 = *(*int32)(unsafe.Add(mBase, uint32(v112)+28))
										if v113 == int32(0) {
											v119 = v106
											if v119 == int32(0) {
												*(*uint8)(unsafe.Add(mBase, uint32(l0)+26)) = uint8(v25)
												v125 = int32(255)
												*(*uint8)(unsafe.Add(mBase, uint32(l0)+27)) = uint8(v125)
												*(*int64)(unsafe.Add(mBase, uint32(l0)+8)) = int64(0)
												*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v95
												*(*int64)(unsafe.Add(mBase, uint32(l0)+16)) = int64(-4294967296)
												v135 = int32(0)
												return v135
											} else {
												F_valkey_free(m, v119)
												mBase = m.M
												v123 = m.ExcPending
												if v123 != 0 {
													return int32(0)
												} else {
													*(*uint8)(unsafe.Add(mBase, uint32(l0)+26)) = uint8(v25)
													v125 = int32(255)
													*(*uint8)(unsafe.Add(mBase, uint32(l0)+27)) = uint8(v125)
													*(*int64)(unsafe.Add(mBase, uint32(l0)+8)) = int64(0)
													*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v95
													*(*int64)(unsafe.Add(mBase, uint32(l0)+16)) = int64(-4294967296)
													v135 = int32(0)
													return v135
												}
											}
										} else {
											m.T0[v113].(func(*base.Module, int32))(m, l0)
											mBase = m.M
											v117 = m.ExcPending
											if v117 != 0 {
												return int32(0)
											} else {
												v118 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
												v119 = v118
												if v119 == int32(0) {
													*(*uint8)(unsafe.Add(mBase, uint32(l0)+26)) = uint8(v25)
													v125 = int32(255)
													*(*uint8)(unsafe.Add(mBase, uint32(l0)+27)) = uint8(v125)
													*(*int64)(unsafe.Add(mBase, uint32(l0)+8)) = int64(0)
													*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v95
													*(*int64)(unsafe.Add(mBase, uint32(l0)+16)) = int64(-4294967296)
													v135 = int32(0)
													return v135
												} else {
													F_valkey_free(m, v119)
													mBase = m.M
													v123 = m.ExcPending
													if v123 != 0 {
														return int32(0)
													} else {
														*(*uint8)(unsafe.Add(mBase, uint32(l0)+26)) = uint8(v25)
														v125 = int32(255)
														*(*uint8)(unsafe.Add(mBase, uint32(l0)+27)) = uint8(v125)
														*(*int64)(unsafe.Add(mBase, uint32(l0)+8)) = int64(0)
														*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v95
														*(*int64)(unsafe.Add(mBase, uint32(l0)+16)) = int64(-4294967296)
														v135 = int32(0)
														return v135
													}
												}
											}
										}
									}
								}
							} else {
								m.T0[v101].(func(*base.Module, int32))(m, l0)
								mBase = m.M
								v105 = m.ExcPending
								if v105 != 0 {
									return int32(0)
								} else {
									v106 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
									if v106 == int32(0) {
										v112 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
										v113 = *(*int32)(unsafe.Add(mBase, uint32(v112)+28))
										if v113 == int32(0) {
											v119 = v106
											if v119 == int32(0) {
												*(*uint8)(unsafe.Add(mBase, uint32(l0)+26)) = uint8(v25)
												v125 = int32(255)
												*(*uint8)(unsafe.Add(mBase, uint32(l0)+27)) = uint8(v125)
												*(*int64)(unsafe.Add(mBase, uint32(l0)+8)) = int64(0)
												*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v95
												*(*int64)(unsafe.Add(mBase, uint32(l0)+16)) = int64(-4294967296)
												v135 = int32(0)
												return v135
											} else {
												F_valkey_free(m, v119)
												mBase = m.M
												v123 = m.ExcPending
												if v123 != 0 {
													return int32(0)
												} else {
													*(*uint8)(unsafe.Add(mBase, uint32(l0)+26)) = uint8(v25)
													v125 = int32(255)
													*(*uint8)(unsafe.Add(mBase, uint32(l0)+27)) = uint8(v125)
													*(*int64)(unsafe.Add(mBase, uint32(l0)+8)) = int64(0)
													*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v95
													*(*int64)(unsafe.Add(mBase, uint32(l0)+16)) = int64(-4294967296)
													v135 = int32(0)
													return v135
												}
											}
										} else {
											m.T0[v113].(func(*base.Module, int32))(m, l0)
											mBase = m.M
											v117 = m.ExcPending
											if v117 != 0 {
												return int32(0)
											} else {
												v118 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
												v119 = v118
												if v119 == int32(0) {
													*(*uint8)(unsafe.Add(mBase, uint32(l0)+26)) = uint8(v25)
													v125 = int32(255)
													*(*uint8)(unsafe.Add(mBase, uint32(l0)+27)) = uint8(v125)
													*(*int64)(unsafe.Add(mBase, uint32(l0)+8)) = int64(0)
													*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v95
													*(*int64)(unsafe.Add(mBase, uint32(l0)+16)) = int64(-4294967296)
													v135 = int32(0)
													return v135
												} else {
													F_valkey_free(m, v119)
													mBase = m.M
													v123 = m.ExcPending
													if v123 != 0 {
														return int32(0)
													} else {
														*(*uint8)(unsafe.Add(mBase, uint32(l0)+26)) = uint8(v25)
														v125 = int32(255)
														*(*uint8)(unsafe.Add(mBase, uint32(l0)+27)) = uint8(v125)
														*(*int64)(unsafe.Add(mBase, uint32(l0)+8)) = int64(0)
														*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v95
														*(*int64)(unsafe.Add(mBase, uint32(l0)+16)) = int64(-4294967296)
														v135 = int32(0)
														return v135
													}
												}
											}
										}
									} else {
										v110 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
										if v110 != 0 {
											v135 = int32(0)
											return v135
										} else {
											v112 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
											v113 = *(*int32)(unsafe.Add(mBase, uint32(v112)+28))
											if v113 == int32(0) {
												v119 = v106
												if v119 == int32(0) {
													*(*uint8)(unsafe.Add(mBase, uint32(l0)+26)) = uint8(v25)
													v125 = int32(255)
													*(*uint8)(unsafe.Add(mBase, uint32(l0)+27)) = uint8(v125)
													*(*int64)(unsafe.Add(mBase, uint32(l0)+8)) = int64(0)
													*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v95
													*(*int64)(unsafe.Add(mBase, uint32(l0)+16)) = int64(-4294967296)
													v135 = int32(0)
													return v135
												} else {
													F_valkey_free(m, v119)
													mBase = m.M
													v123 = m.ExcPending
													if v123 != 0 {
														return int32(0)
													} else {
														*(*uint8)(unsafe.Add(mBase, uint32(l0)+26)) = uint8(v25)
														v125 = int32(255)
														*(*uint8)(unsafe.Add(mBase, uint32(l0)+27)) = uint8(v125)
														*(*int64)(unsafe.Add(mBase, uint32(l0)+8)) = int64(0)
														*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v95
														*(*int64)(unsafe.Add(mBase, uint32(l0)+16)) = int64(-4294967296)
														v135 = int32(0)
														return v135
													}
												}
											} else {
												m.T0[v113].(func(*base.Module, int32))(m, l0)
												mBase = m.M
												v117 = m.ExcPending
												if v117 != 0 {
													return int32(0)
												} else {
													v118 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
													v119 = v118
													if v119 == int32(0) {
														*(*uint8)(unsafe.Add(mBase, uint32(l0)+26)) = uint8(v25)
														v125 = int32(255)
														*(*uint8)(unsafe.Add(mBase, uint32(l0)+27)) = uint8(v125)
														*(*int64)(unsafe.Add(mBase, uint32(l0)+8)) = int64(0)
														*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v95
														*(*int64)(unsafe.Add(mBase, uint32(l0)+16)) = int64(-4294967296)
														v135 = int32(0)
														return v135
													} else {
														F_valkey_free(m, v119)
														mBase = m.M
														v123 = m.ExcPending
														if v123 != 0 {
															return int32(0)
														} else {
															*(*uint8)(unsafe.Add(mBase, uint32(l0)+26)) = uint8(v25)
															v125 = int32(255)
															*(*uint8)(unsafe.Add(mBase, uint32(l0)+27)) = uint8(v125)
															*(*int64)(unsafe.Add(mBase, uint32(l0)+8)) = int64(0)
															*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v95
															*(*int64)(unsafe.Add(mBase, uint32(l0)+16)) = int64(-4294967296)
															v135 = int32(0)
															return v135
														}
													}
												}
											}
										}
									}
								}
							}
						} else {
							v135 = v26
							return v135
						}
					}
				}
			}
		}
	}
}
func F_dictSdsDup(m *base.Module, l0 int32) int32 {
	var v2 int32
	_ = v2
	var v5 int32
	_ = v5
	v2 = F_sdsdup(m, l0)
	v5 = m.ExcPending
	if v5 != 0 {
		return int32(0)
	} else {
		return v2
	}
}
func F_dictSdsHash(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v13 int64
	_ = v13
	var v17 int32
	_ = v17
	var v19 int64
	_ = v19
	var v23 int32
	_ = v23
	var v25 int64
	_ = v25
	var v29 int32
	_ = v29
	var v31 int64
	_ = v31
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v38 int64
	_ = v38
	v7 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0+int32(-1)))))
	switch v7 & int32(7) {
	case 0:
		v13 = F_siphash(m, l0, int32(base.Ui32(v7)>>(uint(int32(3))%32)), int32(_a765))
		mBase = m.M
		return v13
	case 1:
		v17 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0+int32(-3)))))
		v19 = F_siphash(m, l0, v17, int32(_a765))
		mBase = m.M
		return v19
	case 2:
		v23 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0+int32(-5)))))
		v25 = F_siphash(m, l0, v23, int32(_a765))
		mBase = m.M
		return v25
	case 3:
		v29 = *(*int32)(unsafe.Add(mBase, uint32(l0+int32(-9))))
		v31 = F_siphash(m, l0, v29, int32(_a765))
		mBase = m.M
		return v31
	case 4:
		v35 = *(*int32)(unsafe.Add(mBase, uint32(l0+int32(-17))))
		v36 = v35
		v38 = F_siphash(m, l0, v36, int32(_a765))
		mBase = m.M
		return v38
	default:
		v36 = int32(0)
		v38 = F_siphash(m, l0, v36, int32(_a765))
		mBase = m.M
		return v38
	}
}
func F_dictSetHashFunctionSeed(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int64
	_ = v3
	var v8 int64
	_ = v8
	v2 = int32(0)
	v3 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	*(*int64)(unsafe.Add(mBase, _consts[409])) = v3
	v8 = *(*int64)(unsafe.Add(mBase, uint32(l0+int32(8))))
	*(*int64)(unsafe.Add(mBase, _consts[410])) = v8
	return
}
func F_dictSetKey(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(v4)+4))
	if v5 == int32(0) {
		v10 = l2
		*(*int32)(unsafe.Add(mBase, uint32(l1))) = v10
		return
	} else {
		v8 = m.T0[v5].(func(*base.Module, int32) int32)(m, l2)
		mBase = m.M
		v9 = m.ExcPending
		if v9 != 0 {
			return
		} else {
			v10 = v8
			*(*int32)(unsafe.Add(mBase, uint32(l1))) = v10
			return
		}
	}
}
func F_dictStrCaseHash_1(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v16 int32
	_ = v16
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v56 int32
	_ = v56
	var v58 int64
	_ = v58
	if l0&int32(3) == int32(0) {
		v23 = l0
		goto L3
	} else {
		goto L4
	}
L1:
	;
	v58 = F_siphash_nocase(m, l0, v56, int32(_a765))
	mBase = m.M
	goto L17
L2:
	;
	v56 = v48 - l0
	goto L1
L3:
	;
	v27 = v23
	goto L11
L4:
	;
	v9 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v9 != 0 {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	v12 = l0
	goto L7
L6:
	;
	v56 = l0 - l0
	goto L1
L7:
	;
	v16 = v12 + int32(1)
	if v16&int32(3) == int32(0) {
		v23 = v16
		goto L3
	} else {
		goto L9
	}
L9:
	;
	v21 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16))))
	if v21 != 0 {
		v12 = v16
		goto L7
	} else {
		goto L10
	}
L10:
	;
	v48 = v16
	goto L2
L11:
	;
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v27)))
	v36 = int32(-2139062144)
	if (int32(16843008)-v33|v33)&v36 == v36 {
		v27 = v27 + int32(4)
		goto L11
	} else {
		goto L13
	}
L12:
	;
	v42 = v27
	goto L14
L13:
	;
	goto L12
L14:
	;
	v46 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v42))))
	if v46 != 0 {
		v42 = v42 + int32(1)
		goto L14
	} else {
		goto L16
	}
L15:
	;
	v48 = v42
	goto L2
L16:
	;
	goto L15
L17:
	;
	return v58
}
func F_dictStrCaseHash_2(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v16 int32
	_ = v16
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v56 int32
	_ = v56
	var v58 int64
	_ = v58
	if l0&int32(3) == int32(0) {
		v23 = l0
		goto L3
	} else {
		goto L4
	}
L1:
	;
	v58 = F_siphash_nocase(m, l0, v56, int32(_a765))
	mBase = m.M
	goto L17
L2:
	;
	v56 = v48 - l0
	goto L1
L3:
	;
	v27 = v23
	goto L11
L4:
	;
	v9 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v9 != 0 {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	v12 = l0
	goto L7
L6:
	;
	v56 = l0 - l0
	goto L1
L7:
	;
	v16 = v12 + int32(1)
	if v16&int32(3) == int32(0) {
		v23 = v16
		goto L3
	} else {
		goto L9
	}
L9:
	;
	v21 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16))))
	if v21 != 0 {
		v12 = v16
		goto L7
	} else {
		goto L10
	}
L10:
	;
	v48 = v16
	goto L2
L11:
	;
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v27)))
	v36 = int32(-2139062144)
	if (int32(16843008)-v33|v33)&v36 == v36 {
		v27 = v27 + int32(4)
		goto L11
	} else {
		goto L13
	}
L12:
	;
	v42 = v27
	goto L14
L13:
	;
	goto L12
L14:
	;
	v46 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v42))))
	if v46 != 0 {
		v42 = v42 + int32(1)
		goto L14
	} else {
		goto L16
	}
L15:
	;
	v48 = v42
	goto L2
L16:
	;
	goto L15
L17:
	;
	return v58
}
func F_updateDictResizePolicy(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v1 int32
	_ = v1
	var v2 int32
	_ = v2
	var v4 int32
	_ = v4
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	v1 = int32(0)
	v2 = int32(2)
	v4 = *(*int32)(unsafe.Add(mBase, _consts[45]))
	v9 = *(*int32)(unsafe.Add(mBase, _consts[399]))
	if v9 != 0 {
		v10 = base.B2i32(v4 != int32(-1))
	} else {
		v10 = v2
	}
	v12 = *(*int32)(unsafe.Add(mBase, _consts[400]))
	if v12 != 0 {
		v13 = v2
	} else {
		v13 = v10
	}
	*(*int32)(unsafe.Add(mBase, _consts[798])) = v13
	*(*int32)(unsafe.Add(mBase, _consts[427])) = v13
	return
}
