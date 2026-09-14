package p2

import (
	base "github.com/shibukawa/vkmem/internal/aot/vkaot/base"
	"unsafe"
)

func F_functionsLibCtxSwapWithCurrent(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	v3 = int32(0)
	v5 = *(*int32)(unsafe.Add(mBase, _c_F_functionsLibCtxSwapWithCurrent[0]))
	if l1 == v3 {
		F_functionsLibCtxClear(m, v5, int32(0))
		mBase = m.M
		v13 = m.ExcPending
		if v13 != 0 {
			return
		} else {
			v14 = *(*int32)(unsafe.Add(mBase, uint32(v5)+4))
			F_dictRelease(m, v14)
			mBase = m.M
			v16 = m.ExcPending
			if v16 != 0 {
				return
			} else {
				v17 = *(*int32)(unsafe.Add(mBase, uint32(v5)))
				F_dictRelease(m, v17)
				mBase = m.M
				v19 = m.ExcPending
				if v19 != 0 {
					return
				} else {
					v20 = *(*int32)(unsafe.Add(mBase, uint32(v5)+12))
					F_dictRelease(m, v20)
					mBase = m.M
					v22 = m.ExcPending
					if v22 != 0 {
						return
					} else {
						F_valkey_free(m, v5)
						mBase = m.M
						v24 = m.ExcPending
						if v24 != 0 {
							return
						} else {
							*(*int32)(unsafe.Add(mBase, _c_F_functionsLibCtxSwapWithCurrent[0])) = l0
							return
						}
					}
				}
			}
		}
	} else {
		F_freeFunctionsAsync(m, v5, int32(0))
		mBase = m.M
		v10 = m.ExcPending
		if v10 != 0 {
			return
		} else {
			*(*int32)(unsafe.Add(mBase, _c_F_functionsLibCtxSwapWithCurrent[0])) = l0
			return
		}
	}
}
func F_functionsMemory(m *base.Module) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	v3 = m.G0
	v5 = v3 - int32(16)
	m.G0 = v5
	*(*int32)(unsafe.Add(mBase, uint32(v5)+12)) = int32(0)
	F_scriptingEngineManagerForEachEngine(m, int32(532), v5+int32(12))
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		return int32(0)
	} else {
		v16 = *(*int32)(unsafe.Add(mBase, uint32(v5)+12))
		m.G0 = v5 + int32(16)
		return v16
	}
}
func F_functionsMemoryOverhead(m *base.Module) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v12 int32
	_ = v12
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	v2 = *(*int32)(unsafe.Add(mBase, _c_F_functionsMemoryOverhead[0]))
	v3 = F_dictMemUsage(m, v2)
	mBase = m.M
	v6 = int32(0)
	v7 = *(*int32)(unsafe.Add(mBase, _c_F_functionsMemoryOverhead[1]))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(v7)+4))
	v12 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+27)))
	if v12 == int32(255) {
		v16 = v6
	} else {
		v16 = int32(1) << (uint(v12) % 32)
	}
	v19 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+26)))
	if v19 == int32(255) {
		v23 = int32(0)
	} else {
		v23 = int32(1) << (uint(v19) % 32)
	}
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v8)+16))
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v8)+12))
	v34 = int32(0)
	v35 = *(*int32)(unsafe.Add(mBase, _c_F_functionsMemoryOverhead[1]))
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v35)+8))
	v39 = *(*int32)(unsafe.Add(mBase, _c_F_functionsMemoryOverhead[2]))
	return v3 + int32(8) + ((v16+v23)<<(uint(int32(2))%32) + (v27+v28)*int32(24)) + v36 + v39 + int32(16)
}
func F_functionsRemoveLibFromEngine(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
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
	var v31 int32
	_ = v31
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v44 int64
	_ = v44
	var v45 int64
	_ = v45
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
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
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
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v82 int32
	_ = v82
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v102 int32
	_ = v102
	var v107 int32
	_ = v107
	var v111 int32
	_ = v111
	var v115 int32
	_ = v115
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v125 int32
	_ = v125
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v136 int32
	_ = v136
	var v138 int32
	_ = v138
	var v140 int32
	_ = v140
	v5 = *(*int32)(unsafe.Add(mBase, _c_F_functionsRemoveLibFromEngine[0]))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(v5)))
	v7 = F_dictGetSafeIterator(m, v6)
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
	goto L4
L3:
	;
	F_dictReleaseIterator(m, v7)
	mBase = m.M
	v140 = m.ExcPending
	if v140 != 0 {
		goto L1
	} else {
		goto L40
	}
L4:
	;
	v19 = v7 + int32(20)
	v20 = *(*int32)(unsafe.Add(mBase, uint32(v7)+16))
	if v20 != 0 {
		goto L8
	} else {
		goto L9
	}
L6:
	;
	if v115 == int32(0) {
		goto L3
	} else {
		goto L32
	}
L7:
	;
	v26 = v19
	v27 = v23
	goto L10
L8:
	;
	v23 = int32(1)
	goto L7
L9:
	;
	v23 = int32(0)
	goto L7
L10:
	;
	switch v27 {
	case 0:
		goto L15
	default:
		goto L14
	}
L12:
	;
	v27 = int32(0)
	goto L10
L13:
	;
	goto L6
L14:
	;
	v107 = *(*int32)(unsafe.Add(mBase, uint32(v26)))
	*(*int32)(unsafe.Add(mBase, uint32(v7)+16)) = v107
	if v107 == int32(0) {
		goto L12
	} else {
		goto L31
	}
L15:
	;
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v7)+4))
	if v31 != int32(-1) {
		v70 = v31
		goto L16
	} else {
		goto L17
	}
L16:
	;
	v71 = int32(1)
	v72 = v70 + v71
	*(*int32)(unsafe.Add(mBase, uint32(v7)+4)) = v72
	v74 = int32(0)
	v77 = *(*int32)(unsafe.Add(mBase, uint32(v7)))
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v7)+8))
	v82 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v77+v78+int32(26)))))
	if v82 == int32(255) {
		goto L25
	} else {
		goto L26
	}
L17:
	;
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v7)+8))
	if v35 != 0 {
		v70 = int32(-1)
		goto L16
	} else {
		goto L18
	}
L18:
	;
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v7)))
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v7)+12))
	if v37 == int32(0) {
		goto L20
	} else {
		goto L21
	}
L19:
	;
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v63)+20))
	if v64 != int32(-1) {
		goto L22
	} else {
		goto L23
	}
L20:
	;
	v44 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v36)+16)))
	v45 = int64(*(*int8)(unsafe.Add(mBase, uint32(v36)+27)))
	v46 = int64(*(*int32)(unsafe.Add(mBase, uint32(v36)+8)))
	v47 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v36)+12)))
	v48 = int64(*(*int8)(unsafe.Add(mBase, uint32(v36)+26)))
	v49 = int64(*(*int32)(unsafe.Add(mBase, uint32(v36)+4)))
	v50 = F_wangHash64(m, v49)
	mBase = m.M
	v52 = F_wangHash64(m, v48+v50)
	mBase = m.M
	v54 = F_wangHash64(m, v47+v52)
	mBase = m.M
	v56 = F_wangHash64(m, v46+v54)
	mBase = m.M
	v58 = F_wangHash64(m, v45+v56)
	mBase = m.M
	v60 = F_wangHash64(m, v44+v58)
	mBase = m.M
	*(*int64)(unsafe.Add(mBase, uint32(v7)+24)) = v60
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v7)))
	v63 = v62
	goto L19
L21:
	;
	v40 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v36)+24)))
	v42 = v40 + int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v36)+24)) = uint16(v42)
	v63 = v36
	goto L19
L22:
	;
	v70 = v64 + int32(-1)
	goto L16
L23:
	;
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v7)+4))
	v70 = v67
	goto L16
L24:
	;
	v97 = int32(2)
	v102 = *(*int32)(unsafe.Add(mBase, uint32(v77+v95<<(uint(v97)%32)+int32(4))))
	v26 = v102 + v96<<(uint(v97)%32)
	v27 = int32(1)
	goto L10
L25:
	;
	v86 = v74
	goto L27
L26:
	;
	v86 = v71 << (uint(v82) % 32)
	goto L27
L27:
	;
	if v72 < v86 {
		v95 = v78
		v96 = v72
		goto L24
	} else {
		goto L28
	}
L28:
	;
	if v78 != 0 {
		v115 = v74
		goto L13
	} else {
		goto L29
	}
L29:
	;
	v88 = *(*int32)(unsafe.Add(mBase, uint32(v77)+20))
	if v88 == int32(-1) {
		v115 = v74
		goto L13
	} else {
		goto L30
	}
L30:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v7)+4)) = int64(4294967296)
	v95 = int32(1)
	v96 = int32(0)
	goto L24
L31:
	;
	v111 = *(*int32)(unsafe.Add(mBase, uint32(v107)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v19))) = v111
	v115 = v107
	goto L13
L32:
	;
	v121 = *(*int32)(unsafe.Add(mBase, uint32(v115)+8))
	goto L33
L33:
	;
	v122 = *(*int32)(unsafe.Add(mBase, uint32(v121)+8))
	if v122 != l0 {
		goto L4
	} else {
		goto L34
	}
L34:
	;
	v125 = *(*int32)(unsafe.Add(mBase, _c_F_functionsRemoveLibFromEngine[0]))
	F_libraryUnlink(m, v125, v121)
	mBase = m.M
	v127 = m.ExcPending
	if v127 != 0 {
		goto L1
	} else {
		goto L35
	}
L35:
	;
	v128 = *(*int32)(unsafe.Add(mBase, uint32(v121)+4))
	F_dictRelease(m, v128)
	mBase = m.M
	v130 = m.ExcPending
	if v130 != 0 {
		goto L1
	} else {
		goto L36
	}
L36:
	;
	v131 = *(*int32)(unsafe.Add(mBase, uint32(v121)))
	F_sdsfree(m, v131)
	mBase = m.M
	v133 = m.ExcPending
	if v133 != 0 {
		goto L1
	} else {
		goto L37
	}
L37:
	;
	v134 = *(*int32)(unsafe.Add(mBase, uint32(v121)+12))
	F_sdsfree(m, v134)
	mBase = m.M
	v136 = m.ExcPending
	if v136 != 0 {
		goto L1
	} else {
		goto L38
	}
L38:
	;
	F_valkey_free(m, v121)
	mBase = m.M
	v138 = m.ExcPending
	if v138 != 0 {
		goto L1
	} else {
		goto L39
	}
L39:
	;
	goto L4
L40:
	;
	return
}
