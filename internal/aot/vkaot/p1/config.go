package p1

import (
	base "github.com/shibukawa/valkeymem/internal/aot/vkaot/base"
	"math"
	"unsafe"
)

func F_getConfigBindOption(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	v6 = *(*int32)(unsafe.Add(mBase, _consts[135]))
	v8 = F_sdsjoin(m, int32(_a261), v6, int32(_a6))
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return int32(0)
	} else {
		return v8
	}
}
func F_getConfigDirOption(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	v4 = m.G0
	v5 = int32(1024)
	v6 = v4 - v5
	m.G0 = v6
	v9 = F_getcwd(m, v6, v5)
	mBase = m.M
	if v9 != 0 {
	} else {
		v10 = int32(0)
		*(*uint8)(unsafe.Add(mBase, uint32(v6))) = uint8(v10)
	}
	v12 = F_sdsnew(m, v6)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		return int32(0)
	} else {
		m.G0 = v6 + int32(1024)
		return v12
	}
}
func F_getConfigLatencyTrackingInfoPercentilesOutputOption(m *base.Module, l0 int32) int32 {
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
	var v16 int32
	_ = v16
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v29 float64
	_ = v29
	var v32 int32
	_ = v32
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v64 int32
	_ = v64
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v90 int32
	_ = v90
	v6 = m.G0
	v8 = v6 - int32(144)
	m.G0 = v8
	v10 = F_sdsempty(m)
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v14 = int32(0)
	v16 = *(*int32)(unsafe.Add(mBase, _consts[276]))
	if v16 <= v14 {
		v90 = v10
		goto L3
	} else {
		goto L4
	}
L3:
	;
	m.G0 = v8 + int32(144)
	return v90
L4:
	;
	v21 = v10
	v22 = v14
	goto L5
L5:
	;
	v25 = *(*int32)(unsafe.Add(mBase, _consts[275]))
	v29 = *(*float64)(unsafe.Add(mBase, uint32(v25+v22<<(uint(int32(3))%32))))
	*(*float64)(unsafe.Add(mBase, uint32(v8))) = v29
	v32 = v8 + int32(16)
	v39 = F_snprintf(m, v32, int32(128), int32(_a546), v8)
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
		goto L1
	} else {
		goto L7
	}
L6:
	;
	v90 = v83
	goto L3
L7:
	;
	v44 = F_strchr(m, v32, int32(46))
	mBase = m.M
	if v44 == int32(0) {
		v64 = v39
		goto L9
	} else {
		goto L10
	}
L8:
	;
	v70 = F_sdscatlen(m, v21, v32, v64)
	mBase = m.M
	v71 = m.ExcPending
	if v71 != 0 {
		goto L1
	} else {
		goto L16
	}
L9:
	;
	v68 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v32+v64))) = uint8(v68)
	goto L8
L10:
	;
	v49 = v39
	v50 = v32 + v39
	goto L11
L11:
	;
	v53 = v50 + int32(-1)
	v54 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v53))))
	if v54 == int32(48) {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	v49 = v49 + int32(-1)
	v50 = v53
	goto L11
L14:
	;
	if v54 != int32(46) {
		v64 = v49
		goto L9
	} else {
		goto L15
	}
L15:
	;
	v64 = v49 + int32(-1)
	goto L9
L16:
	;
	v73 = *(*int32)(unsafe.Add(mBase, _consts[276]))
	if v22 == v73+int32(-1) {
		v83 = v70
		v84 = v73
		goto L17
	} else {
		goto L18
	}
L17:
	;
	v86 = v22 + int32(1)
	if v86 < v84 {
		v21 = v83
		v22 = v86
		goto L5
	} else {
		goto L20
	}
L18:
	;
	v79 = F_sdscatlen(m, v70, int32(_a6), int32(1))
	mBase = m.M
	v80 = m.ExcPending
	if v80 != 0 {
		goto L1
	} else {
		goto L19
	}
L19:
	;
	v82 = *(*int32)(unsafe.Add(mBase, _consts[276]))
	v83 = v79
	v84 = v82
	goto L17
L20:
	;
	goto L6
}
func F_getConfigReplicaOfOption(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v14 int32
	_ = v14
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	v4 = m.G0
	v6 = v4 - int32(272)
	m.G0 = v6
	v9 = *(*int32)(unsafe.Add(mBase, _consts[166]))
	if v9 == int32(0) {
		v24 = int32(0)
		*(*uint8)(unsafe.Add(mBase, uint32(v6)+16)) = uint8(v24)
		v28 = F_sdsnew(m, v6+int32(16))
		mBase = m.M
		v29 = m.ExcPending
		if v29 != 0 {
			return int32(0)
		} else {
			m.G0 = v6 + int32(272)
			return v28
		}
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(v6))) = v9
		v14 = *(*int32)(unsafe.Add(mBase, _consts[274]))
		*(*int32)(unsafe.Add(mBase, uint32(v6)+4)) = v14
		v20 = F_snprintf(m, v6+int32(16), int32(256), int32(_a543), v6)
		mBase = m.M
		v23 = m.ExcPending
		if v23 != 0 {
			return int32(0)
		} else {
			v28 = F_sdsnew(m, v6+int32(16))
			mBase = m.M
			v29 = m.ExcPending
			if v29 != 0 {
				return int32(0)
			} else {
				m.G0 = v6 + int32(272)
				return v28
			}
		}
	}
}
func F_initConfigValues(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v33 int32
	_ = v33
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
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
	var v58 int32
	_ = v58
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v89 int32
	_ = v89
	var v95 int32
	_ = v95
	var v101 int32
	_ = v101
	var v107 int32
	_ = v107
	v6 = F_dictCreate(m, int32(_a547))
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
	*(*int32)(unsafe.Add(mBase, _consts[247])) = v6
	v10 = F_dictExpand(m, v6, int32(226))
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v12 = int32(0)
	v13 = *(*int32)(unsafe.Add(mBase, _consts[277]))
	if v13 == v12 {
		goto L8
	} else {
		goto L9
	}
L4:
	;
	F__serverAssert(m, int32(_a548), int32(_a473), int32(3571))
	mBase = m.M
	v107 = m.ExcPending
	if v107 != 0 {
		goto L1
	} else {
		goto L39
	}
L5:
	;
	F__serverAssert(m, int32(_a548), int32(_a473), int32(3565))
	mBase = m.M
	v101 = m.ExcPending
	if v101 != 0 {
		goto L1
	} else {
		goto L38
	}
L6:
	;
	F__serverAssert(m, int32(_a549), int32(_a473), int32(3560))
	mBase = m.M
	v95 = m.ExcPending
	if v95 != 0 {
		goto L1
	} else {
		goto L37
	}
L7:
	;
	F__serverAssert(m, int32(_a550), int32(_a473), int32(3555))
	mBase = m.M
	v89 = m.ExcPending
	if v89 != 0 {
		goto L1
	} else {
		goto L36
	}
L8:
	;
	return
L9:
	;
	v17 = int32(_a551)
	goto L10
L10:
	;
	v20 = *(*int32)(unsafe.Add(mBase, uint32(v17)+12))
	if v20 == int32(0) {
		goto L12
	} else {
		goto L13
	}
L11:
	;
	goto L8
L12:
	;
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v17)+80))
	if v25 != int32(1) {
		goto L15
	} else {
		goto L16
	}
L13:
	;
	m.T0[v20].(func(*base.Module, int32))(m, v17)
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		goto L1
	} else {
		goto L14
	}
L14:
	;
	goto L12
L15:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v17)))
	v40 = F_valkey_malloc(m, int32(88))
	mBase = m.M
	v41 = m.ExcPending
	if v41 != 0 {
		goto L1
	} else {
		goto L19
	}
L16:
	;
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v17)+36))
	if v28&int32(17) == int32(16) {
		goto L7
	} else {
		goto L17
	}
L17:
	;
	v33 = int32(18)
	if v28&v33 == v33 {
		goto L6
	} else {
		goto L18
	}
L18:
	;
	goto L15
L19:
	;
	goto L22
L20:
	;
	v48 = *(*int32)(unsafe.Add(mBase, _consts[247]))
	v49 = F_sdsnew(m, v38)
	mBase = m.M
	v50 = m.ExcPending
	if v50 != 0 {
		goto L1
	} else {
		goto L23
	}
L21:
	;
	goto L20
L22:
	;
	v45 = F__emscripten_memcpy_bulkmem(m, v40, v17, int32(88))
	mBase = m.M
	goto L21
L23:
	;
	v51 = F_dictAdd(m, v48, v49, v45)
	mBase = m.M
	v52 = m.ExcPending
	if v52 != 0 {
		goto L1
	} else {
		goto L24
	}
L24:
	;
	if v51 != 0 {
		goto L5
	} else {
		goto L25
	}
L25:
	;
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v17)+4))
	if v53 == int32(0) {
		goto L26
	} else {
		goto L27
	}
L26:
	;
	v79 = *(*int32)(unsafe.Add(mBase, uint32(v17)+88))
	if v79 != 0 {
		v17 = v17 + int32(88)
		goto L10
	} else {
		goto L35
	}
L27:
	;
	v57 = F_valkey_malloc(m, int32(88))
	mBase = m.M
	v58 = m.ExcPending
	if v58 != 0 {
		goto L1
	} else {
		goto L28
	}
L28:
	;
	goto L31
L29:
	;
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v62)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v62)+8)) = v64 | int32(128)
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v17)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v62))) = v68
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v17)))
	*(*int32)(unsafe.Add(mBase, uint32(v62)+4)) = v70
	v73 = *(*int32)(unsafe.Add(mBase, _consts[247]))
	v74 = F_sdsnew(m, v53)
	mBase = m.M
	v75 = m.ExcPending
	if v75 != 0 {
		goto L1
	} else {
		goto L32
	}
L30:
	;
	goto L29
L31:
	;
	v62 = F__emscripten_memcpy_bulkmem(m, v57, v17, int32(88))
	mBase = m.M
	goto L30
L32:
	;
	v76 = F_dictAdd(m, v73, v74, v62)
	mBase = m.M
	v77 = m.ExcPending
	if v77 != 0 {
		goto L1
	} else {
		goto L33
	}
L33:
	;
	if v76 != 0 {
		goto L4
	} else {
		goto L34
	}
L34:
	;
	goto L26
L35:
	;
	goto L11
L36:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L37:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L38:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L39:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_removeConfig(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
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
	var v30 int32
	_ = v30
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
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	v6 = *(*int32)(unsafe.Add(mBase, _consts[247]))
	v7 = F_dictFind(m, v6, l0)
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
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
	if v7 == int32(0) {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v11 = *(*int32)(unsafe.Add(mBase, uint32(v7)+8))
	goto L5
L5:
	;
	if v11 == int32(0) {
		goto L1
	} else {
		goto L6
	}
L6:
	;
	v14 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+9)))
	if v14&int32(1) == int32(0) {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	v55 = *(*int32)(unsafe.Add(mBase, _consts[247]))
	v56 = F_dictDelete(m, v55, l0)
	mBase = m.M
	v57 = m.ExcPending
	if v57 != 0 {
		goto L2
	} else {
		goto L21
	}
L8:
	;
	v19 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
	F_sdsfree(m, v19)
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		goto L2
	} else {
		goto L9
	}
L9:
	;
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v11)+80))
	switch v22 + int32(-3) {
	case 0:
		goto L10
	case 1:
		goto L11
	default:
		goto L7
	}
L10:
	;
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v11)+36))
	if v45 == int32(0) {
		goto L7
	} else {
		goto L19
	}
L11:
	;
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v11)+36))
	v26 = *(*int32)(unsafe.Add(mBase, uint32(v25)))
	if v26 == int32(0) {
		v40 = v25
		goto L12
	} else {
		goto L13
	}
L12:
	;
	F_valkey_free(m, v40)
	mBase = m.M
	v44 = m.ExcPending
	if v44 != 0 {
		goto L2
	} else {
		goto L18
	}
L13:
	;
	v30 = v25
	v32 = v26
	goto L14
L14:
	;
	F_valkey_free(m, v32)
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
		goto L2
	} else {
		goto L16
	}
L15:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v11)+36))
	v40 = v38
	goto L12
L16:
	;
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v30)+8))
	if v35 != 0 {
		v30 = v30 + int32(8)
		v32 = v35
		goto L14
	} else {
		goto L17
	}
L17:
	;
	goto L15
L18:
	;
	goto L7
L19:
	;
	F_sdsfree(m, v45)
	mBase = m.M
	v49 = m.ExcPending
	if v49 != 0 {
		goto L2
	} else {
		goto L20
	}
L20:
	;
	goto L7
L21:
	;
	goto L1
}
func F_rewriteConfigBindOption(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
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
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	if l3 < int32(1) {
		v11 = F_sdsnew(m, int32(_a542))
		v12 = m.ExcPending
		if v12 != 0 {
			return
		} else {
			v13 = v11
			v14 = F_sdsnew(m, l0)
			v15 = m.ExcPending
			if v15 != 0 {
				return
			} else {
				v18 = F_sdscatlen(m, v14, int32(_a6), int32(1))
				v19 = m.ExcPending
				if v19 != 0 {
					return
				} else {
					v20 = F_sdscatsds(m, v18, v13)
					v21 = m.ExcPending
					if v21 != 0 {
						return
					} else {
						F_sdsfree(m, v13)
						v23 = m.ExcPending
						if v23 != 0 {
							return
						} else {
							v25 = F_rewriteConfigRewriteLine(m, l1, l0, v20, int32(1))
							v26 = m.ExcPending
							if v26 != 0 {
								return
							} else {
								return
							}
						}
					}
				}
			}
		}
	} else {
		v8 = F_sdsjoin(m, l2, l3, int32(_a6))
		v9 = m.ExcPending
		if v9 != 0 {
			return
		} else {
			v13 = v8
			v14 = F_sdsnew(m, l0)
			v15 = m.ExcPending
			if v15 != 0 {
				return
			} else {
				v18 = F_sdscatlen(m, v14, int32(_a6), int32(1))
				v19 = m.ExcPending
				if v19 != 0 {
					return
				} else {
					v20 = F_sdscatsds(m, v18, v13)
					v21 = m.ExcPending
					if v21 != 0 {
						return
					} else {
						F_sdsfree(m, v13)
						v23 = m.ExcPending
						if v23 != 0 {
							return
						} else {
							v25 = F_rewriteConfigRewriteLine(m, l1, l0, v20, int32(1))
							v26 = m.ExcPending
							if v26 != 0 {
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
}
func F_rewriteConfigGetContentFromState(m *base.Module, l0 int32) int32 {
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
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	v9 = F_sdsempty(m)
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v13 < int32(1) {
		v76 = v9
		goto L3
	} else {
		goto L4
	}
L3:
	;
	return v76
L4:
	;
	v16 = int32(0)
	v19 = v9
	v20 = v13
	v21 = v16
	v22 = v16
	goto L5
L5:
	;
	v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v26+v22<<(uint(int32(2))%32))))
	v33 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v30+int32(-1)))))
	switch v33 & int32(7) {
	case 0:
		goto L15
	case 1:
		goto L14
	case 2:
		goto L13
	case 3:
		goto L12
	case 4:
		goto L11
	default:
		goto L9
	}
L6:
	;
	v76 = v67
	goto L3
L7:
	;
	v73 = v22 + int32(1)
	if v73 < v68 {
		v19 = v67
		v20 = v68
		v21 = v69
		v22 = v73
		goto L5
	} else {
		goto L20
	}
L8:
	;
	v60 = F_sdscatsds(m, v19, v30)
	mBase = m.M
	v61 = m.ExcPending
	if v61 != 0 {
		goto L1
	} else {
		goto L18
	}
L9:
	;
	if v21 == int32(0) {
		v58 = int32(1)
		goto L8
	} else {
		goto L17
	}
L10:
	;
	if v50 != 0 {
		v58 = int32(0)
		goto L8
	} else {
		goto L16
	}
L11:
	;
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v30+int32(-17))))
	v50 = v49
	goto L10
L12:
	;
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v30+int32(-9))))
	v50 = v46
	goto L10
L13:
	;
	v43 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v30+int32(-5)))))
	v50 = v43
	goto L10
L14:
	;
	v40 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v30+int32(-3)))))
	v50 = v40
	goto L10
L15:
	;
	v50 = int32(base.Ui32(v33) >> (uint(int32(3)) % 32))
	goto L10
L16:
	;
	goto L9
L17:
	;
	v67 = v19
	v68 = v20
	v69 = int32(1)
	goto L7
L18:
	;
	v64 = F_sdscatlen(m, v60, int32(_a26), int32(1))
	mBase = m.M
	v65 = m.ExcPending
	if v65 != 0 {
		goto L1
	} else {
		goto L19
	}
L19:
	;
	v66 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v67 = v64
	v68 = v66
	v69 = v58
	goto L7
L20:
	;
	goto L6
}
func F_rewriteConfigNotifyKeyspaceEventsOption(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
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
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v22 int32
	_ = v22
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
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
	v10 = *(*int32)(unsafe.Add(mBase, _consts[252]))
	v11 = F_keyspaceEventsFlagsToString(m, v10)
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		return
	} else {
		v13 = F_sdsnew(m, l1)
		mBase = m.M
		v14 = m.ExcPending
		if v14 != 0 {
			return
		} else {
			v17 = F_sdscatlen(m, v13, int32(_a6), int32(1))
			mBase = m.M
			v18 = m.ExcPending
			if v18 != 0 {
				return
			} else {
				v22 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11+int32(-1)))))
				switch v22 & int32(7) {
				case 0:
					v39 = int32(base.Ui32(v22) >> (uint(int32(3)) % 32))
				case 1:
					v29 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11+int32(-3)))))
					v39 = v29
				case 2:
					v32 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v11+int32(-5)))))
					v39 = v32
				case 3:
					v35 = *(*int32)(unsafe.Add(mBase, uint32(v11+int32(-9))))
					v39 = v35
				case 4:
					v38 = *(*int32)(unsafe.Add(mBase, uint32(v11+int32(-17))))
					v39 = v38
				default:
					v39 = int32(0)
				}
				v40 = F_sdscatrepr(m, v17, v11, v39)
				mBase = m.M
				v41 = m.ExcPending
				if v41 != 0 {
					return
				} else {
					F_sdsfree(m, v11)
					mBase = m.M
					v43 = m.ExcPending
					if v43 != 0 {
						return
					} else {
						v46 = F_rewriteConfigRewriteLine(m, l2, l1, v40, base.B2i32(v10 != int32(0)))
						mBase = m.M
						v47 = m.ExcPending
						if v47 != 0 {
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
func F_rewriteConfigOverwriteFile(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v37 int32
	_ = v37
	var v44 int32
	_ = v44
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
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
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v105 int32
	_ = v105
	var v109 int32
	_ = v109
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v115 int32
	_ = v115
	var v118 int32
	_ = v118
	var v126 int32
	_ = v126
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
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v150 int32
	_ = v150
	var v155 int32
	_ = v155
	var v158 int32
	_ = v158
	var v160 int32
	_ = v160
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v172 int32
	_ = v172
	var v175 int32
	_ = v175
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
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v198 int32
	_ = v198
	var v200 int32
	_ = v200
	var v202 int32
	_ = v202
	var v205 int32
	_ = v205
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v215 int32
	_ = v215
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v234 int32
	_ = v234
	var v239 int32
	_ = v239
	var v241 int32
	_ = v241
	var v248 int32
	_ = v248
	var v251 int32
	_ = v251
	var v255 int32
	_ = v255
	var v256 int32
	_ = v256
	var v263 int32
	_ = v263
	var v264 int32
	_ = v264
	var v265 int32
	_ = v265
	var v269 int32
	_ = v269
	var v275 int32
	_ = v275
	var v276 int32
	_ = v276
	var v279 int32
	_ = v279
	var v280 int32
	_ = v280
	var v281 int32
	_ = v281
	var v282 int32
	_ = v282
	var v285 int32
	_ = v285
	var v288 int32
	_ = v288
	var v292 int32
	_ = v292
	var v293 int32
	_ = v293
	var v298 int32
	_ = v298
	var v302 int32
	_ = v302
	v15 = m.G0
	v17 = v15 - int32(4224)
	m.G0 = v17
	*(*int32)(unsafe.Add(mBase, uint32(v17)+112)) = l0
	*(*int32)(unsafe.Add(mBase, uint32(v17)+116)) = int32(_a500)
	v28 = F_snprintf(m, v17+int32(128), int32(4096), int32(_a501), v17+int32(112))
	mBase = m.M
	v31 = m.ExcPending
	if v31 != 0 {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	m.G0 = v17 + int32(4224)
	return v302
L2:
	;
	v49 = int32(-1)
	v52 = int32(0)
	v54 = F___mkostemps(m, v17+int32(128), v52, v52)
	mBase = m.M
	goto L11
L3:
	;
	return int32(0)
L4:
	;
	v32 = int32(-4096)
	if base.Ui32(v32) < base.Ui32(v28+v32) {
		goto L2
	} else {
		goto L5
	}
L5:
	;
	v37 = *(*int32)(unsafe.Add(mBase, _consts[28]))
	if int32(3) < v37 {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	goto L9
L7:
	;
	F__serverLog(m, int32(3), int32(_a502), int32(0))
	mBase = m.M
	v44 = m.ExcPending
	if v44 != 0 {
		goto L3
	} else {
		goto L8
	}
L8:
	;
	goto L6
L9:
	;
	*(*int32)(unsafe.Add(mBase, _consts[18])) = int32(37)
	v302 = int32(-1)
	goto L1
L10:
	;
	v288 = *(*int32)(unsafe.Add(mBase, _consts[28]))
	if int32(3) < v288 {
		v302 = v49
		goto L1
	} else {
		goto L79
	}
L11:
	;
	if v54 == int32(-1) {
		goto L10
	} else {
		goto L12
	}
L12:
	;
	v58 = l1 + int32(-3)
	v60 = l1 + int32(-5)
	v62 = l1 + int32(-9)
	v64 = l1 + int32(-17)
	v67 = int32(0)
	v80 = l1 + v67
	v81 = v67
	goto L14
L13:
	;
	v280 = int32(9116376)
	goto L76
L14:
	;
	v85 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1+int32(-1)))))
	v87 = v85 & int32(7)
	switch v87 {
	case 0:
		goto L21
	case 1:
		goto L20
	case 2:
		goto L19
	case 3:
		goto L18
	case 4:
		goto L17
	default:
		v94 = int32(0)
		goto L16
	}
L15:
	;
	v129 = F_fsync(m, v54)
	mBase = m.M
	if v129 == int32(0) {
		goto L37
	} else {
		goto L38
	}
L16:
	;
	if base.Ui32(v94) <= base.Ui32(v81) {
		goto L22
	} else {
		goto L23
	}
L17:
	;
	v93 = *(*int32)(unsafe.Add(mBase, uint32(v64)))
	v94 = v93
	goto L16
L18:
	;
	v92 = *(*int32)(unsafe.Add(mBase, uint32(v62)))
	v94 = v92
	goto L16
L19:
	;
	v91 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v60))))
	v94 = v91
	goto L16
L20:
	;
	v90 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v58))))
	v94 = v90
	goto L16
L21:
	;
	v94 = int32(base.Ui32(v85) >> (uint(int32(3)) % 32))
	goto L16
L22:
	;
	goto L15
L23:
	;
	switch v87 {
	case 0:
		goto L29
	case 1:
		goto L28
	case 2:
		goto L27
	case 3:
		goto L26
	case 4:
		goto L25
	default:
		v103 = int32(0)
		goto L24
	}
L24:
	;
	v105 = F_write(m, v54, v80, v103-v81)
	mBase = m.M
	if int32(0) < v105 {
		goto L30
	} else {
		goto L31
	}
L25:
	;
	v102 = *(*int32)(unsafe.Add(mBase, uint32(v64)))
	v103 = v102
	goto L24
L26:
	;
	v101 = *(*int32)(unsafe.Add(mBase, uint32(v62)))
	v103 = v101
	goto L24
L27:
	;
	v100 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v60))))
	v103 = v100
	goto L24
L28:
	;
	v99 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v58))))
	v103 = v99
	goto L24
L29:
	;
	v103 = int32(base.Ui32(v85) >> (uint(int32(3)) % 32))
	goto L24
L30:
	;
	v127 = v105 + v81
	v80 = l1 + v127
	v81 = v127
	goto L14
L31:
	;
	goto L32
L32:
	;
	v109 = *(*int32)(unsafe.Add(mBase, _consts[18]))
	if v109 == int32(27) {
		goto L14
	} else {
		goto L33
	}
L33:
	;
	v112 = int32(-1)
	v113 = int32(0)
	v115 = *(*int32)(unsafe.Add(mBase, _consts[28]))
	if int32(3) < v115 {
		v276 = v112
		v279 = v113
		goto L13
	} else {
		goto L34
	}
L34:
	;
	v118 = F___strerror_l(m, v109, v109)
	mBase = m.M
	goto L35
L35:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+20)) = v118
	*(*int32)(unsafe.Add(mBase, uint32(v17)+16)) = v81
	F__serverLog(m, int32(3), int32(_a503), v17+int32(16))
	mBase = m.M
	v126 = m.ExcPending
	if v126 != 0 {
		goto L3
	} else {
		goto L36
	}
L36:
	;
	v276 = v112
	v279 = v113
	goto L13
L37:
	;
	v148 = int32(-1)
	v150 = *(*int32)(unsafe.Add(mBase, _consts[253]))
	v155 = F_fchmod(m, v54, (v150^v148)&int32(420))
	mBase = m.M
	if v155 != v148 {
		goto L43
	} else {
		goto L44
	}
L38:
	;
	v132 = int32(-1)
	v133 = int32(0)
	v135 = *(*int32)(unsafe.Add(mBase, _consts[28]))
	if int32(3) < v135 {
		v276 = v132
		v279 = v133
		goto L13
	} else {
		goto L39
	}
L39:
	;
	goto L40
L40:
	;
	v139 = *(*int32)(unsafe.Add(mBase, _consts[18]))
	v140 = F___strerror_l(m, v139, v139)
	mBase = m.M
	goto L41
L41:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+96)) = v140
	F__serverLog(m, int32(3), int32(_a504), v17+int32(96))
	mBase = m.M
	v147 = m.ExcPending
	if v147 != 0 {
		goto L3
	} else {
		goto L42
	}
L42:
	;
	v276 = v132
	v279 = v133
	goto L13
L43:
	;
	v175 = F_rename(m, v17+int32(128), l0)
	mBase = m.M
	if v175 != int32(-1) {
		goto L49
	} else {
		goto L50
	}
L44:
	;
	v158 = int32(0)
	v160 = *(*int32)(unsafe.Add(mBase, _consts[28]))
	if int32(3) < v160 {
		v276 = v148
		v279 = v158
		goto L13
	} else {
		goto L45
	}
L45:
	;
	goto L46
L46:
	;
	v164 = *(*int32)(unsafe.Add(mBase, _consts[18]))
	v165 = F___strerror_l(m, v164, v164)
	mBase = m.M
	goto L47
L47:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+32)) = v165
	F__serverLog(m, int32(3), int32(_a505), v17+int32(32))
	mBase = m.M
	v172 = m.ExcPending
	if v172 != 0 {
		goto L3
	} else {
		goto L48
	}
L48:
	;
	v276 = v148
	v279 = v158
	goto L13
L49:
	;
	v194 = int32(-1)
	v198 = m.G0
	v200 = v198 - int32(4112)
	m.G0 = v200
	v202 = F_strlen(m, l0)
	mBase = m.M
	if base.Ui32(v202) < base.Ui32(int32(4097)) {
		goto L57
	} else {
		goto L58
	}
L50:
	;
	v178 = int32(-1)
	v179 = int32(0)
	v181 = *(*int32)(unsafe.Add(mBase, _consts[28]))
	if int32(3) < v181 {
		v276 = v178
		v279 = v179
		goto L13
	} else {
		goto L51
	}
L51:
	;
	goto L52
L52:
	;
	v185 = *(*int32)(unsafe.Add(mBase, _consts[18]))
	v186 = F___strerror_l(m, v185, v185)
	mBase = m.M
	goto L53
L53:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+48)) = v186
	F__serverLog(m, int32(3), int32(_a506), v17+int32(48))
	mBase = m.M
	v193 = m.ExcPending
	if v193 != 0 {
		goto L3
	} else {
		goto L54
	}
L54:
	;
	v276 = v178
	v279 = v179
	goto L13
L55:
	;
	v248 = *(*int32)(unsafe.Add(mBase, _consts[28]))
	if v241 != int32(-1) {
		goto L68
	} else {
		goto L69
	}
L56:
	;
	m.G0 = v200 + int32(4112)
	goto L55
L57:
	;
	v211 = F___memcpy(m, v200, l0, v202+int32(1))
	mBase = m.M
	v212 = F_dirname(m, v211)
	mBase = m.M
	v213 = int32(0)
	v215 = F_open(m, v212, v213, v213)
	mBase = m.M
	if v215 != int32(-1) {
		goto L59
	} else {
		goto L60
	}
L58:
	;
	v205 = F___errno_location(m)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v205))) = int32(37)
	v241 = int32(-1)
	goto L56
L59:
	;
	v225 = F_fsync(m, v215)
	mBase = m.M
	if v225 != int32(-1) {
		goto L64
	} else {
		goto L65
	}
L60:
	;
	v220 = F___errno_location(m)
	mBase = m.M
	v221 = *(*int32)(unsafe.Add(mBase, uint32(v220)))
	if v221 != int32(31) {
		goto L61
	} else {
		goto L62
	}
L61:
	;
	v224 = int32(-1)
	goto L63
L62:
	;
	v224 = int32(0)
	goto L63
L63:
	;
	v241 = v224
	goto L56
L64:
	;
	v239 = F_close(m, v215)
	mBase = m.M
	v241 = int32(0)
	goto L56
L65:
	;
	v228 = F___errno_location(m)
	mBase = m.M
	v229 = *(*int32)(unsafe.Add(mBase, uint32(v228)))
	if v229 == int32(8) {
		goto L64
	} else {
		goto L66
	}
L66:
	;
	if v229 == int32(28) {
		goto L64
	} else {
		goto L67
	}
L67:
	;
	v234 = F_close(m, v215)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v228))) = v229
	v241 = int32(-1)
	goto L56
L68:
	;
	v264 = int32(0)
	v265 = int32(1)
	if v264 < v248 {
		v276 = v264
		v279 = v265
		goto L13
	} else {
		goto L74
	}
L69:
	;
	v251 = int32(0)
	if int32(3) < v248 {
		v276 = v194
		v279 = v251
		goto L13
	} else {
		goto L70
	}
L70:
	;
	goto L71
L71:
	;
	v255 = *(*int32)(unsafe.Add(mBase, _consts[18]))
	v256 = F___strerror_l(m, v255, v255)
	mBase = m.M
	goto L72
L72:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+64)) = v256
	F__serverLog(m, int32(3), int32(_a507), v17+int32(64))
	mBase = m.M
	v263 = m.ExcPending
	if v263 != 0 {
		goto L3
	} else {
		goto L73
	}
L73:
	;
	v276 = v194
	v279 = v251
	goto L13
L74:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+80)) = l0
	v269 = int32(0)
	F__serverLog(m, v269, int32(_a508), v17+int32(80))
	mBase = m.M
	v275 = m.ExcPending
	if v275 != 0 {
		goto L3
	} else {
		goto L75
	}
L75:
	;
	v276 = v269
	v279 = v265
	goto L13
L76:
	;
	v281 = *(*int32)(unsafe.Add(mBase, _consts[18]))
	v282 = F_close(m, v54)
	mBase = m.M
	if v279 != 0 {
		goto L77
	} else {
		goto L78
	}
L77:
	;
	*(*int32)(unsafe.Add(mBase, _consts[18])) = v281
	v302 = v276
	goto L1
L78:
	;
	v285 = F_unlink(m, v17+int32(128))
	mBase = m.M
	goto L77
L79:
	;
	goto L80
L80:
	;
	v292 = *(*int32)(unsafe.Add(mBase, _consts[18]))
	v293 = F___strerror_l(m, v292, v292)
	mBase = m.M
	goto L81
L81:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17))) = v293
	F__serverLog(m, int32(3), int32(_a509), v17)
	mBase = m.M
	v298 = m.ExcPending
	if v298 != 0 {
		goto L3
	} else {
		goto L82
	}
L82:
	;
	v302 = v49
	goto L1
}
func F_rewriteConfigSocketBindOption(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v86 int32
	_ = v86
	var v91 int32
	_ = v91
	v6 = *(*int32)(unsafe.Add(mBase, _consts[135]))
	if v6 != int32(2) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	return
L2:
	;
	F_rewriteConfigBindOption(m, l1, l2, int32(_a261), v6)
	mBase = m.M
	v91 = m.ExcPending
	if v91 != 0 {
		goto L22
	} else {
		goto L27
	}
L3:
	;
	v10 = *(*int32)(unsafe.Add(mBase, _consts[270]))
	v11 = int32(_a540)
	v14 = int32(*(*uint8)(unsafe.Add(mBase, _consts[271])))
	v15 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10))))
	if v15 == int32(0) {
		v38 = v14
		v39 = v15
		goto L5
	} else {
		goto L6
	}
L4:
	;
	if v39-v38&int32(255) != 0 {
		goto L2
	} else {
		goto L12
	}
L5:
	;
	goto L4
L6:
	;
	if v15 != v14&int32(255) {
		v38 = v14
		v39 = v15
		goto L5
	} else {
		goto L7
	}
L7:
	;
	v21 = v10
	v22 = v11
	goto L8
L8:
	;
	v25 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22)+1)))
	v26 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21)+1)))
	if v26 == int32(0) {
		v38 = v25
		v39 = v26
		goto L5
	} else {
		goto L10
	}
L9:
	;
	v38 = v25
	v39 = v26
	goto L5
L10:
	;
	v29 = int32(1)
	if v26 == v25&int32(255) {
		v21 = v21 + v29
		v22 = v22 + v29
		goto L8
	} else {
		goto L11
	}
L11:
	;
	goto L9
L12:
	;
	v44 = *(*int32)(unsafe.Add(mBase, _consts[272]))
	v45 = int32(_a541)
	v48 = int32(*(*uint8)(unsafe.Add(mBase, _consts[273])))
	v49 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v44))))
	if v49 == int32(0) {
		v72 = v48
		v73 = v49
		goto L14
	} else {
		goto L15
	}
L13:
	;
	if v73-v72&int32(255) != 0 {
		goto L2
	} else {
		goto L21
	}
L14:
	;
	goto L13
L15:
	;
	if v49 != v48&int32(255) {
		v72 = v48
		v73 = v49
		goto L14
	} else {
		goto L16
	}
L16:
	;
	v55 = v44
	v56 = v45
	goto L17
L17:
	;
	v59 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v56)+1)))
	v60 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v55)+1)))
	if v60 == int32(0) {
		v72 = v59
		v73 = v60
		goto L14
	} else {
		goto L19
	}
L18:
	;
	v72 = v59
	v73 = v60
	goto L14
L19:
	;
	v63 = int32(1)
	if v60 == v59&int32(255) {
		v55 = v55 + v63
		v56 = v56 + v63
		goto L17
	} else {
		goto L20
	}
L20:
	;
	goto L18
L21:
	;
	v77 = F_sdsnew(m, l1)
	mBase = m.M
	v78 = m.ExcPending
	if v78 != 0 {
		goto L22
	} else {
		goto L23
	}
L22:
	;
	return
L23:
	;
	v79 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	v81 = F_dictAdd(m, v79, v77, int32(0))
	mBase = m.M
	v82 = m.ExcPending
	if v82 != 0 {
		goto L22
	} else {
		goto L24
	}
L24:
	;
	if v81 == int32(0) {
		goto L1
	} else {
		goto L25
	}
L25:
	;
	F_sdsfree(m, v77)
	mBase = m.M
	v86 = m.ExcPending
	if v86 != 0 {
		goto L22
	} else {
		goto L26
	}
L26:
	;
	return
L27:
	;
	goto L1
}
func F_rewriteConfigStringOption(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v15 int32
	_ = v15
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v73 int32
	_ = v73
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v84 int32
	_ = v84
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	var v99 int32
	_ = v99
	var v103 int32
	_ = v103
	var v105 int32
	_ = v105
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	if l2 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	return
L2:
	;
	if l3 == int32(0) {
		v52 = int32(1)
		goto L9
	} else {
		goto L10
	}
L3:
	;
	v6 = F_sdsnew(m, l1)
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	return
L5:
	;
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v10 = F_dictAdd(m, v8, v6, int32(0))
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		goto L4
	} else {
		goto L6
	}
L6:
	;
	if v10 == int32(0) {
		goto L1
	} else {
		goto L7
	}
L7:
	;
	F_sdsfree(m, v6)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		goto L4
	} else {
		goto L8
	}
L8:
	;
	return
L9:
	;
	v53 = F_sdsnew(m, l1)
	mBase = m.M
	v54 = m.ExcPending
	if v54 != 0 {
		goto L4
	} else {
		goto L19
	}
L10:
	;
	v21 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3))))
	v22 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2))))
	if v22 == int32(0) {
		v45 = v21
		v46 = v22
		goto L12
	} else {
		goto L13
	}
L11:
	;
	v52 = base.B2i32(v46-v45&int32(255) != int32(0))
	goto L9
L12:
	;
	goto L11
L13:
	;
	if v22 != v21&int32(255) {
		v45 = v21
		v46 = v22
		goto L12
	} else {
		goto L14
	}
L14:
	;
	v28 = l2
	v29 = l3
	goto L15
L15:
	;
	v32 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v29)+1)))
	v33 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28)+1)))
	if v33 == int32(0) {
		v45 = v32
		v46 = v33
		goto L12
	} else {
		goto L17
	}
L16:
	;
	v45 = v32
	v46 = v33
	goto L12
L17:
	;
	v36 = int32(1)
	if v33 == v32&int32(255) {
		v28 = v28 + v36
		v29 = v29 + v36
		goto L15
	} else {
		goto L18
	}
L18:
	;
	goto L16
L19:
	;
	v57 = F_sdscatlen(m, v53, int32(_a6), int32(1))
	mBase = m.M
	v58 = m.ExcPending
	if v58 != 0 {
		goto L4
	} else {
		goto L20
	}
L20:
	;
	if l2&int32(3) == int32(0) {
		v80 = l2
		goto L23
	} else {
		goto L24
	}
L21:
	;
	v114 = F_sdscatrepr(m, v57, l2, v113)
	mBase = m.M
	v115 = m.ExcPending
	if v115 != 0 {
		goto L4
	} else {
		goto L37
	}
L22:
	;
	v113 = v105 - l2
	goto L21
L23:
	;
	v84 = v80
	goto L31
L24:
	;
	v66 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2))))
	if v66 != 0 {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	v69 = l2
	goto L27
L26:
	;
	v113 = l2 - l2
	goto L21
L27:
	;
	v73 = v69 + int32(1)
	if v73&int32(3) == int32(0) {
		v80 = v73
		goto L23
	} else {
		goto L29
	}
L29:
	;
	v78 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v73))))
	if v78 != 0 {
		v69 = v73
		goto L27
	} else {
		goto L30
	}
L30:
	;
	v105 = v73
	goto L22
L31:
	;
	v90 = *(*int32)(unsafe.Add(mBase, uint32(v84)))
	v93 = int32(-2139062144)
	if (int32(16843008)-v90|v90)&v93 == v93 {
		v84 = v84 + int32(4)
		goto L31
	} else {
		goto L33
	}
L32:
	;
	v99 = v84
	goto L34
L33:
	;
	goto L32
L34:
	;
	v103 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v99))))
	if v103 != 0 {
		v99 = v99 + int32(1)
		goto L34
	} else {
		goto L36
	}
L35:
	;
	v105 = v99
	goto L22
L36:
	;
	goto L35
L37:
	;
	v116 = F_rewriteConfigRewriteLine(m, l0, l1, v114, v52)
	mBase = m.M
	v117 = m.ExcPending
	if v117 != 0 {
		goto L4
	} else {
		goto L38
	}
L38:
	;
	goto L1
}
func F_setConfigClientOutputBufferLimitOption(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v25 int32
	_ = v25
	var v34 int32
	_ = v34
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v53 int32
	_ = v53
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v79 int32
	_ = v79
	var v82 int64
	_ = v82
	var v85 int32
	_ = v85
	var v88 int64
	_ = v88
	var v91 int32
	_ = v91
	var v96 int64
	_ = v96
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v107 int32
	_ = v107
	var v116 int32
	_ = v116
	var v128 int32
	_ = v128
	var v130 int32
	_ = v130
	var v134 int64
	_ = v134
	var v139 int64
	_ = v139
	var v146 int64
	_ = v146
	var v159 int32
	_ = v159
	var v163 int64
	_ = v163
	var v170 int64
	_ = v170
	var v177 int64
	_ = v177
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v184 int64
	_ = v184
	var v191 int64
	_ = v191
	var v198 int64
	_ = v198
	var v209 int32
	_ = v209
	v5 = int32(0)
	v12 = m.G0
	v14 = v12 - int32(112)
	m.G0 = v14
	*(*int32)(unsafe.Add(mBase, uint32(v14+int32(8)))) = v5
	*(*int64)(unsafe.Add(mBase, uint32(v14))) = int64(0)
	if l2&int32(3) != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v14 + int32(112)
	return v209
L2:
	;
	if l3 == int32(0) {
		v209 = v5
		goto L1
	} else {
		goto L37
	}
L3:
	;
	v25 = int32(0)
	if l2 <= v25 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v159 = *(*int32)(unsafe.Add(mBase, uint32(v14)+4))
	if v159 == int32(0) {
		goto L34
	} else {
		goto L35
	}
L5:
	;
	v34 = v25
	goto L6
L6:
	;
	v41 = l1 + v34<<(uint(int32(2))%32)
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v41)))
	v45 = F_strcasecmp(m, v42, int32(_a529))
	mBase = m.M
	if v45 != 0 {
		goto L11
	} else {
		goto L12
	}
L7:
	;
	v130 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
	if v130 == int32(0) {
		goto L4
	} else {
		goto L33
	}
L8:
	;
	v79 = *(*int32)(unsafe.Add(mBase, uint32(v41+int32(4))))
	v82 = F_memtoull(m, v79, v14+int32(108))
	mBase = m.M
	v85 = *(*int32)(unsafe.Add(mBase, uint32(v41+int32(8))))
	v88 = F_memtoull(m, v85, v14+int32(104))
	mBase = m.M
	v91 = *(*int32)(unsafe.Add(mBase, uint32(v41+int32(12))))
	v96 = F_strtox_2(m, v91, v14+int32(100), int32(10), int64(-9223372036854775807-1))
	mBase = m.M
	goto L24
L9:
	;
	v72 = int32(0)
	if l3 == v72 {
		v209 = v72
		goto L1
	} else {
		goto L23
	}
L10:
	;
	switch v69 + int32(1) {
	case 0, 4:
		goto L9
	default:
		goto L8
	}
L11:
	;
	v47 = int32(1)
	v49 = F_strcasecmp(m, v42, int32(_a530))
	mBase = m.M
	if v49 == int32(0) {
		v67 = v47
		goto L13
	} else {
		goto L14
	}
L12:
	;
	v69 = int32(0)
	goto L10
L13:
	;
	v69 = v67
	goto L10
L14:
	;
	v53 = F_strcasecmp(m, v42, int32(_a531))
	mBase = m.M
	if v53 == int32(0) {
		v67 = v47
		goto L13
	} else {
		goto L15
	}
L15:
	;
	v57 = F_strcasecmp(m, v42, int32(_a532))
	mBase = m.M
	if v57 != 0 {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	v60 = F_strcasecmp(m, v42, int32(_a533))
	mBase = m.M
	if v60 != 0 {
		goto L18
	} else {
		goto L19
	}
L17:
	;
	v69 = int32(2)
	goto L10
L18:
	;
	v65 = F_strcasecmp(m, v42, int32(_a534))
	mBase = m.M
	if v65 != 0 {
		goto L20
	} else {
		goto L21
	}
L19:
	;
	v69 = int32(3)
	goto L10
L20:
	;
	v66 = int32(-1)
	goto L22
L21:
	;
	v66 = int32(3)
	goto L22
L22:
	;
	v67 = v66
	goto L13
L23:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = int32(_a535)
	v209 = v72
	goto L1
L24:
	;
	v97 = *(*int32)(unsafe.Add(mBase, uint32(v14)+108))
	if v97 != 0 {
		goto L26
	} else {
		goto L27
	}
L25:
	;
	v116 = v14 + int32(16) + v69*int32(24)
	*(*int64)(unsafe.Add(mBase, uint32(v116)+16)) = v96 & int64(2147483647)
	*(*int64)(unsafe.Add(mBase, uint32(v116)+8)) = v88
	*(*int64)(unsafe.Add(mBase, uint32(v116))) = v82
	*(*int32)(unsafe.Add(mBase, uint32(v14+v69<<(uint(int32(2))%32)))) = int32(1)
	v128 = v34 + int32(4)
	if v128 < l2 {
		v34 = v128
		goto L6
	} else {
		goto L32
	}
L26:
	;
	v107 = int32(0)
	if l3 == v107 {
		v209 = v107
		goto L1
	} else {
		goto L31
	}
L27:
	;
	v98 = *(*int32)(unsafe.Add(mBase, uint32(v14)+104))
	if v98 != 0 {
		goto L26
	} else {
		goto L28
	}
L28:
	;
	if v96&int64(2147483648) != int64(0) {
		goto L26
	} else {
		goto L29
	}
L29:
	;
	v103 = *(*int32)(unsafe.Add(mBase, uint32(v14)+100))
	v104 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v103))))
	if v104 == int32(0) {
		goto L25
	} else {
		goto L30
	}
L30:
	;
	goto L26
L31:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = int32(_a536)
	v209 = v107
	goto L1
L32:
	;
	goto L7
L33:
	;
	v134 = *(*int64)(unsafe.Add(mBase, uint32(v14)+16))
	*(*int64)(unsafe.Add(mBase, _consts[258])) = v134
	v139 = *(*int64)(unsafe.Add(mBase, uint32(v14)+24))
	*(*int64)(unsafe.Add(mBase, _consts[259])) = v139
	v146 = *(*int64)(unsafe.Add(mBase, uint32(v14+int32(32))))
	*(*int64)(unsafe.Add(mBase, _consts[260])) = v146
	goto L4
L34:
	;
	v179 = int32(1)
	v180 = *(*int32)(unsafe.Add(mBase, uint32(v14)+8))
	if v180 == int32(0) {
		v209 = v179
		goto L1
	} else {
		goto L36
	}
L35:
	;
	v163 = *(*int64)(unsafe.Add(mBase, uint32(v14)+40))
	*(*int64)(unsafe.Add(mBase, _consts[261])) = v163
	v170 = *(*int64)(unsafe.Add(mBase, uint32(v14+int32(56))))
	*(*int64)(unsafe.Add(mBase, _consts[262])) = v170
	v177 = *(*int64)(unsafe.Add(mBase, uint32(v14+int32(48))))
	*(*int64)(unsafe.Add(mBase, _consts[263])) = v177
	goto L34
L36:
	;
	v184 = *(*int64)(unsafe.Add(mBase, uint32(v14)+64))
	*(*int64)(unsafe.Add(mBase, _consts[264])) = v184
	v191 = *(*int64)(unsafe.Add(mBase, uint32(v14+int32(80))))
	*(*int64)(unsafe.Add(mBase, _consts[265])) = v191
	v198 = *(*int64)(unsafe.Add(mBase, uint32(v14+int32(72))))
	*(*int64)(unsafe.Add(mBase, _consts[266])) = v198
	v209 = v179
	goto L1
L37:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = int32(_a537)
	v209 = v5
	goto L1
}
func F_setConfigLatencyTrackingInfoPercentilesOutputOption(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v73 int32
	_ = v73
	var v83 int32
	_ = v83
	var v86 int32
	_ = v86
	var v93 int32
	_ = v93
	var v96 int32
	_ = v96
	var v99 int32
	_ = v99
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v111 int32
	_ = v111
	var v113 int32
	_ = v113
	var v115 int32
	_ = v115
	var v120 float64
	_ = v120
	var v124 int32
	_ = v124
	var v131 int32
	_ = v131
	var v134 int32
	_ = v134
	var v141 int32
	_ = v141
	var v144 int64
	_ = v144
	var v145 int64
	_ = v145
	var v151 int32
	_ = v151
	var v157 int32
	_ = v157
	var v160 int32
	_ = v160
	var v167 int32
	_ = v167
	var v168 float64
	_ = v168
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v179 int32
	_ = v179
	var v181 int32
	_ = v181
	var v183 int32
	_ = v183
	var v185 int32
	_ = v185
	var v189 int32
	_ = v189
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v203 int32
	_ = v203
	var v216 int32
	_ = v216
	v12 = m.G0
	v14 = v12 - int32(16)
	m.G0 = v14
	v17 = *(*int32)(unsafe.Add(mBase, _consts[275]))
	F_valkey_free(m, v17)
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v22 = int32(_a20)
	*(*int32)(unsafe.Add(mBase, _consts[276])) = l2
	v24 = int32(0)
	*(*int32)(unsafe.Add(mBase, _consts[275])) = v24
	v28 = int32(1)
	if l2 != v28 {
		goto L5
	} else {
		goto L6
	}
L3:
	;
	m.G0 = v14 + int32(16)
	return v216
L4:
	;
	*(*int32)(unsafe.Add(mBase, _consts[276])) = int32(0)
	v216 = v203
	goto L3
L5:
	;
	v61 = F_valkey_malloc(m, l2<<(uint(int32(3))%32))
	mBase = m.M
	v62 = m.ExcPending
	if v62 != 0 {
		goto L1
	} else {
		goto L14
	}
L6:
	;
	v31 = int32(1)
	v32 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v35 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v32+int32(-1)))))
	switch v35 & int32(7) {
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
		v203 = v31
		goto L4
	}
L7:
	;
	if v52 == int32(0) {
		v203 = v31
		goto L4
	} else {
		goto L13
	}
L8:
	;
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v32+int32(-17))))
	v52 = v51
	goto L7
L9:
	;
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v32+int32(-9))))
	v52 = v48
	goto L7
L10:
	;
	v45 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v32+int32(-5)))))
	v52 = v45
	goto L7
L11:
	;
	v42 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v32+int32(-3)))))
	v52 = v42
	goto L7
L12:
	;
	v52 = int32(base.Ui32(v35) >> (uint(int32(3)) % 32))
	goto L7
L13:
	;
	goto L5
L14:
	;
	*(*int32)(unsafe.Add(mBase, _consts[275])) = v61
	v65 = *(*int32)(unsafe.Add(mBase, _consts[276]))
	if v65 < int32(1) {
		v216 = v28
		goto L3
	} else {
		goto L15
	}
L15:
	;
	v73 = v24
	goto L16
L16:
	;
	v83 = *(*int32)(unsafe.Add(mBase, uint32(l1+v73<<(uint(int32(2))%32))))
	v86 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v83+int32(-1)))))
	switch v86 & int32(7) {
	case 0:
		goto L23
	case 1:
		goto L22
	case 2:
		goto L21
	case 3:
		goto L20
	case 4:
		goto L19
	default:
		v103 = int32(0)
		goto L18
	}
L17:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v185
	v189 = *(*int32)(unsafe.Add(mBase, _consts[275]))
	F_valkey_free(m, v189)
	mBase = m.M
	v191 = m.ExcPending
	if v191 != 0 {
		goto L1
	} else {
		goto L44
	}
L18:
	;
	v105 = v14 + int32(8)
	v106 = int32(0)
	v111 = m.G0
	v113 = v111 - int32(16)
	m.G0 = v113
	v115 = F___errno_location(m)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v115))) = v106
	v120 = F_valkey_strtod_n(m, v83, v103, v113+int32(12))
	mBase = m.M
	*(*float64)(unsafe.Add(mBase, uint32(v105))) = v120
	if v103 == v106 {
		goto L28
	} else {
		goto L29
	}
L19:
	;
	v102 = *(*int32)(unsafe.Add(mBase, uint32(v83+int32(-17))))
	v103 = v102
	goto L18
L20:
	;
	v99 = *(*int32)(unsafe.Add(mBase, uint32(v83+int32(-9))))
	v103 = v99
	goto L18
L21:
	;
	v96 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v83+int32(-5)))))
	v103 = v96
	goto L18
L22:
	;
	v93 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v83+int32(-3)))))
	v103 = v93
	goto L18
L23:
	;
	v103 = int32(base.Ui32(v86) >> (uint(int32(3)) % 32))
	goto L18
L24:
	;
	goto L17
L25:
	;
	v167 = int32(_a544)
	v168 = *(*float64)(unsafe.Add(mBase, uint32(v14)+8))
	if base.F64_gt(v168, float64(100)) != 0 {
		v185 = v167
		goto L24
	} else {
		goto L41
	}
L26:
	;
	if v160 != 0 {
		goto L25
	} else {
		goto L40
	}
L27:
	;
	m.G0 = v113 + int32(16)
	goto L26
L28:
	;
	v157 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v115))) = v157
	v160 = v157
	goto L27
L29:
	;
	v124 = int32(*(*int8)(unsafe.Add(mBase, uint32(v83))))
	if v124 == int32(32) {
		goto L28
	} else {
		goto L30
	}
L30:
	;
	if base.Ui32(int32(-6)) < base.Ui32(v124+int32(-14)) {
		goto L28
	} else {
		goto L31
	}
L31:
	;
	v131 = *(*int32)(unsafe.Add(mBase, uint32(v113)+12))
	if v131-v83 != v103 {
		goto L28
	} else {
		goto L32
	}
L32:
	;
	v134 = *(*int32)(unsafe.Add(mBase, uint32(v115)))
	if v134 == int32(68) {
		goto L34
	} else {
		goto L35
	}
L33:
	;
	if base.Ui64(int64(9218868437227405312)) < base.Ui64(v145&int64(9223372036854775807)) {
		goto L28
	} else {
		goto L38
	}
L34:
	;
	if base.F64_eq(base.F64_abs(v120), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		goto L28
	} else {
		goto L36
	}
L35:
	;
	v145 = base.I64_reinterpret_f64(v120)
	goto L33
L36:
	;
	v141 = F___fpclassify(m, v120)
	mBase = m.M
	if v141 == int32(2) {
		goto L28
	} else {
		goto L37
	}
L37:
	;
	v144 = *(*int64)(unsafe.Add(mBase, uint32(v105)))
	v145 = v144
	goto L33
L38:
	;
	v151 = *(*int32)(unsafe.Add(mBase, uint32(v115)))
	if v151 != int32(28) {
		v160 = int32(1)
		goto L27
	} else {
		goto L39
	}
L39:
	;
	goto L28
L40:
	;
	v185 = int32(_a545)
	goto L24
L41:
	;
	if base.F64_lt(v168, float64(0)) != 0 {
		v185 = v167
		goto L24
	} else {
		goto L42
	}
L42:
	;
	v173 = int32(_a20)
	v174 = *(*int32)(unsafe.Add(mBase, _consts[275]))
	*(*float64)(unsafe.Add(mBase, uint32(v174+v73<<(uint(int32(3))%32)))) = v168
	v179 = int32(1)
	v181 = v73 + v179
	v183 = *(*int32)(unsafe.Add(mBase, _consts[276]))
	if v181 < v183 {
		v73 = v181
		goto L16
	} else {
		goto L43
	}
L43:
	;
	v216 = v179
	goto L3
L44:
	;
	v192 = int32(0)
	*(*int32)(unsafe.Add(mBase, _consts[275])) = v192
	v203 = v192
	goto L4
}
func F_setConfigOOMScoreAdjValuesOption(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v15 int32
	_ = v15
	var v20 int64
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v27 int32
	_ = v27
	var v32 int64
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v39 int32
	_ = v39
	var v44 int64
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v51 int32
	_ = v51
	var v55 int32
	_ = v55
	var v59 int32
	_ = v59
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v94 int32
	_ = v94
	var v107 int32
	_ = v107
	v9 = m.G0
	v11 = v9 - int32(16)
	m.G0 = v11
	if l2 != int32(3) {
		*(*int32)(unsafe.Add(mBase, uint32(l3))) = int32(_a483)
		v107 = int32(0)
		m.G0 = v11 + int32(16)
		return v107
	} else {
		v15 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
		v20 = F_strtox_2(m, v15, v11+int32(12), int32(10), int64(-9223372036854775807-1))
		mBase = m.M
		v21 = *(*int32)(unsafe.Add(mBase, uint32(v11)+12))
		v22 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21))))
		if v22 != 0 {
			if l3 == int32(0) {
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(l3))) = int32(_a538)
			}
			v107 = int32(0)
			m.G0 = v11 + int32(16)
			return v107
		} else {
			if base.Ui64(int64(4001)) <= base.Ui64(v20+int64(2000)) {
				if l3 == int32(0) {
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(l3))) = int32(_a538)
				}
				v107 = int32(0)
				m.G0 = v11 + int32(16)
				return v107
			} else {
				v27 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
				v32 = F_strtox_2(m, v27, v11+int32(12), int32(10), int64(-9223372036854775807-1))
				mBase = m.M
				v33 = *(*int32)(unsafe.Add(mBase, uint32(v11)+12))
				v34 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v33))))
				if v34 != 0 {
					if l3 == int32(0) {
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(l3))) = int32(_a538)
					}
					v107 = int32(0)
					m.G0 = v11 + int32(16)
					return v107
				} else {
					if base.Ui64(int64(4000)) < base.Ui64(v32+int64(2000)) {
						if l3 == int32(0) {
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(l3))) = int32(_a538)
						}
						v107 = int32(0)
						m.G0 = v11 + int32(16)
						return v107
					} else {
						v39 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
						v44 = F_strtox_2(m, v39, v11+int32(12), int32(10), int64(-9223372036854775807-1))
						mBase = m.M
						v45 = *(*int32)(unsafe.Add(mBase, uint32(v11)+12))
						v46 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v45))))
						if v46 != 0 {
							if l3 == int32(0) {
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(l3))) = int32(_a538)
							}
							v107 = int32(0)
							m.G0 = v11 + int32(16)
							return v107
						} else {
							if base.Ui64(int64(4000)) < base.Ui64(v44+int64(2000)) {
								if l3 == int32(0) {
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(l3))) = int32(_a538)
								}
								v107 = int32(0)
								m.G0 = v11 + int32(16)
								return v107
							} else {
								v51 = base.I32_wrap_i64(v20)
								if v32 < v20 {
									v59 = *(*int32)(unsafe.Add(mBase, _consts[28]))
									if int32(3) < v59 {
										v69 = base.I32_wrap_i64(v32)
										v72 = *(*int32)(unsafe.Add(mBase, _consts[267]))
										if v72 == v51 {
											v77 = int32(0)
										} else {
											*(*int32)(unsafe.Add(mBase, _consts[267])) = v51
											v77 = int32(1)
										}
										v78 = base.I32_wrap_i64(v44)
										v80 = *(*int32)(unsafe.Add(mBase, _consts[268]))
										if v80 == v69 {
											v85 = v77
										} else {
											*(*int32)(unsafe.Add(mBase, _consts[268])) = v69
											v85 = int32(1)
										}
										v87 = *(*int32)(unsafe.Add(mBase, _consts[269]))
										if v87 == v78 {
											if v85 != 0 {
												v94 = int32(1)
											} else {
												v94 = int32(2)
											}
											v107 = v94
										} else {
											*(*int32)(unsafe.Add(mBase, _consts[269])) = v78
											v107 = int32(1)
										}
										m.G0 = v11 + int32(16)
										return v107
									} else {
										F__serverLog(m, int32(3), int32(_a539), int32(0))
										mBase = m.M
										v68 = m.ExcPending
										if v68 != 0 {
											return int32(0)
										} else {
											v69 = base.I32_wrap_i64(v32)
											v72 = *(*int32)(unsafe.Add(mBase, _consts[267]))
											if v72 == v51 {
												v77 = int32(0)
											} else {
												*(*int32)(unsafe.Add(mBase, _consts[267])) = v51
												v77 = int32(1)
											}
											v78 = base.I32_wrap_i64(v44)
											v80 = *(*int32)(unsafe.Add(mBase, _consts[268]))
											if v80 == v69 {
												v85 = v77
											} else {
												*(*int32)(unsafe.Add(mBase, _consts[268])) = v69
												v85 = int32(1)
											}
											v87 = *(*int32)(unsafe.Add(mBase, _consts[269]))
											if v87 == v78 {
												if v85 != 0 {
													v94 = int32(1)
												} else {
													v94 = int32(2)
												}
												v107 = v94
											} else {
												*(*int32)(unsafe.Add(mBase, _consts[269])) = v78
												v107 = int32(1)
											}
											m.G0 = v11 + int32(16)
											return v107
										}
									}
								} else {
									if v32 <= v44 {
										v69 = base.I32_wrap_i64(v32)
										v72 = *(*int32)(unsafe.Add(mBase, _consts[267]))
										if v72 == v51 {
											v77 = int32(0)
										} else {
											*(*int32)(unsafe.Add(mBase, _consts[267])) = v51
											v77 = int32(1)
										}
										v78 = base.I32_wrap_i64(v44)
										v80 = *(*int32)(unsafe.Add(mBase, _consts[268]))
										if v80 == v69 {
											v85 = v77
										} else {
											*(*int32)(unsafe.Add(mBase, _consts[268])) = v69
											v85 = int32(1)
										}
										v87 = *(*int32)(unsafe.Add(mBase, _consts[269]))
										if v87 == v78 {
											if v85 != 0 {
												v94 = int32(1)
											} else {
												v94 = int32(2)
											}
											v107 = v94
										} else {
											*(*int32)(unsafe.Add(mBase, _consts[269])) = v78
											v107 = int32(1)
										}
										m.G0 = v11 + int32(16)
										return v107
									} else {
										v55 = *(*int32)(unsafe.Add(mBase, _consts[28]))
										if v55 <= int32(3) {
											F__serverLog(m, int32(3), int32(_a539), int32(0))
											mBase = m.M
											v68 = m.ExcPending
											if v68 != 0 {
												return int32(0)
											} else {
												v69 = base.I32_wrap_i64(v32)
												v72 = *(*int32)(unsafe.Add(mBase, _consts[267]))
												if v72 == v51 {
													v77 = int32(0)
												} else {
													*(*int32)(unsafe.Add(mBase, _consts[267])) = v51
													v77 = int32(1)
												}
												v78 = base.I32_wrap_i64(v44)
												v80 = *(*int32)(unsafe.Add(mBase, _consts[268]))
												if v80 == v69 {
													v85 = v77
												} else {
													*(*int32)(unsafe.Add(mBase, _consts[268])) = v69
													v85 = int32(1)
												}
												v87 = *(*int32)(unsafe.Add(mBase, _consts[269]))
												if v87 == v78 {
													if v85 != 0 {
														v94 = int32(1)
													} else {
														v94 = int32(2)
													}
													v107 = v94
												} else {
													*(*int32)(unsafe.Add(mBase, _consts[269])) = v78
													v107 = int32(1)
												}
												m.G0 = v11 + int32(16)
												return v107
											}
										} else {
											v69 = base.I32_wrap_i64(v32)
											v72 = *(*int32)(unsafe.Add(mBase, _consts[267]))
											if v72 == v51 {
												v77 = int32(0)
											} else {
												*(*int32)(unsafe.Add(mBase, _consts[267])) = v51
												v77 = int32(1)
											}
											v78 = base.I32_wrap_i64(v44)
											v80 = *(*int32)(unsafe.Add(mBase, _consts[268]))
											if v80 == v69 {
												v85 = v77
											} else {
												*(*int32)(unsafe.Add(mBase, _consts[268])) = v69
												v85 = int32(1)
											}
											v87 = *(*int32)(unsafe.Add(mBase, _consts[269]))
											if v87 == v78 {
												if v85 != 0 {
													v94 = int32(1)
												} else {
													v94 = int32(2)
												}
												v107 = v94
											} else {
												*(*int32)(unsafe.Add(mBase, _consts[269])) = v78
												v107 = int32(1)
											}
											m.G0 = v11 + int32(16)
											return v107
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
func F_setConfigSaveOption(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
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
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v62 int32
	_ = v62
	var v66 int32
	_ = v66
	var v70 int32
	_ = v70
	var v78 int32
	_ = v78
	var v86 int32
	_ = v86
	var v91 int64
	_ = v91
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v109 int32
	_ = v109
	var v114 int32
	_ = v114
	var v117 int32
	_ = v117
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v127 int32
	_ = v127
	var v129 int32
	_ = v129
	var v132 int32
	_ = v132
	var v134 int32
	_ = v134
	var v147 int32
	_ = v147
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v159 int64
	_ = v159
	var v162 int32
	_ = v162
	var v166 int64
	_ = v166
	var v167 int32
	_ = v167
	var v169 int32
	_ = v169
	var v171 int32
	_ = v171
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v183 int32
	_ = v183
	var v186 int32
	_ = v186
	var v192 int32
	_ = v192
	var v204 int32
	_ = v204
	v11 = m.G0
	v13 = v11 - int32(16)
	m.G0 = v13
	if l2 != int32(1) {
		goto L4
	} else {
		goto L5
	}
L1:
	;
	m.G0 = v13 + int32(16)
	return v204
L2:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = int32(_a528)
	v204 = int32(0)
	goto L1
L3:
	;
	v122 = int32(0)
	v123 = int32(*(*uint8)(unsafe.Add(mBase, _consts[244])))
	if v123 == v122 {
		goto L36
	} else {
		goto L37
	}
L4:
	;
	if l2&int32(1) != 0 {
		goto L2
	} else {
		goto L21
	}
L5:
	;
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v18 = int32(_a320)
	v21 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17))))
	if v21 != 0 {
		goto L8
	} else {
		goto L9
	}
L6:
	;
	if v53-v55 != 0 {
		goto L2
	} else {
		goto L18
	}
L7:
	;
	v53 = F_tolower(m, v49)
	mBase = m.M
	v54 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v50))))
	v55 = F_tolower(m, v54)
	mBase = m.M
	goto L6
L8:
	;
	v23 = v17
	v24 = v18
	v25 = v21
	goto L11
L9:
	;
	v49 = int32(0)
	v50 = v18
	goto L7
L10:
	;
	v49 = v46 & int32(255)
	v50 = v45
	goto L7
L11:
	;
	v27 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24))))
	if v27 == int32(0) {
		v45 = v24
		v46 = v25
		goto L10
	} else {
		goto L13
	}
L12:
	;
	v45 = v39
	v46 = int32(0)
	goto L10
L13:
	;
	v31 = v25 & int32(255)
	if v31 == v27 {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v38 = int32(1)
	v39 = v24 + v38
	v40 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23)+1)))
	if v40 != 0 {
		v23 = v23 + v38
		v24 = v39
		v25 = v40
		goto L11
	} else {
		goto L17
	}
L15:
	;
	v33 = F_tolower(m, v31)
	mBase = m.M
	v34 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24))))
	v35 = F_tolower(m, v34)
	mBase = m.M
	if v33 == v35 {
		goto L14
	} else {
		goto L16
	}
L16:
	;
	v37 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23))))
	v45 = v24
	v46 = v37
	goto L10
L17:
	;
	goto L12
L18:
	;
	v58 = *(*int32)(unsafe.Add(mBase, _consts[243]))
	F_valkey_free(m, v58)
	mBase = m.M
	v62 = m.ExcPending
	if v62 != 0 {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	return int32(0)
L20:
	;
	*(*int64)(unsafe.Add(mBase, _consts[243])) = int64(0)
	v66 = int32(0)
	v114 = v66
	v117 = v66
	goto L3
L21:
	;
	v70 = int32(0)
	if l2 <= v70 {
		v114 = l2
		v117 = v70
		goto L3
	} else {
		goto L22
	}
L22:
	;
	v78 = v70
	goto L23
L23:
	;
	v86 = *(*int32)(unsafe.Add(mBase, uint32(l1+v78<<(uint(int32(2))%32))))
	v91 = F_strtox_2(m, v86, v13+int32(12), int32(10), int64(-9223372036854775807-1))
	mBase = m.M
	goto L25
L24:
	;
	v114 = l2
	v117 = int32(1)
	goto L3
L25:
	;
	v92 = *(*int32)(unsafe.Add(mBase, uint32(v13)+12))
	v93 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v92))))
	if v93 != 0 {
		goto L27
	} else {
		goto L28
	}
L26:
	;
	v109 = v78 + int32(1)
	if v109 != l2 {
		v78 = v109
		goto L23
	} else {
		goto L34
	}
L27:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = int32(_a528)
	v204 = int32(0)
	goto L1
L28:
	;
	v94 = base.I32_wrap_i64(v91)
	v96 = v78 & int32(1)
	if v96 != 0 {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	if v96 == int32(0) {
		goto L26
	} else {
		goto L32
	}
L30:
	;
	if v94 < int32(1) {
		goto L27
	} else {
		goto L31
	}
L31:
	;
	goto L29
L32:
	;
	if int32(-1) < v94 {
		goto L26
	} else {
		goto L33
	}
L33:
	;
	goto L27
L34:
	;
	goto L24
L35:
	;
	if v117 == int32(0) {
		v204 = int32(1)
		goto L1
	} else {
		goto L40
	}
L36:
	;
	v132 = *(*int32)(unsafe.Add(mBase, _consts[243]))
	F_valkey_free(m, v132)
	mBase = m.M
	v134 = m.ExcPending
	if v134 != 0 {
		goto L19
	} else {
		goto L39
	}
L37:
	;
	v127 = int32(*(*uint8)(unsafe.Add(mBase, _consts[257])))
	if v127 != 0 {
		goto L35
	} else {
		goto L38
	}
L38:
	;
	v129 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _consts[257])) = uint8(v129)
	goto L36
L39:
	;
	*(*int64)(unsafe.Add(mBase, _consts[243])) = int64(0)
	goto L35
L40:
	;
	v147 = int32(0)
	goto L41
L41:
	;
	v154 = l1 + v147<<(uint(int32(2))%32)
	v155 = *(*int32)(unsafe.Add(mBase, uint32(v154)))
	v159 = F_strtox_2(m, v155, int32(0), int32(10), int64(-9223372036854775807-1))
	mBase = m.M
	goto L43
L43:
	;
	v162 = *(*int32)(unsafe.Add(mBase, uint32(v154+int32(4))))
	v166 = F_strtox_2(m, v162, int32(0), int32(10), int64(-9223372036854775807-1))
	mBase = m.M
	goto L44
L44:
	;
	v167 = int32(_a20)
	v169 = *(*int32)(unsafe.Add(mBase, _consts[243]))
	v171 = *(*int32)(unsafe.Add(mBase, _consts[59]))
	v176 = F_valkey_realloc(m, v169, v171<<(uint(int32(4))%32)+int32(16))
	mBase = m.M
	v177 = m.ExcPending
	if v177 != 0 {
		goto L19
	} else {
		goto L45
	}
L45:
	;
	*(*int32)(unsafe.Add(mBase, _consts[243])) = v176
	v179 = int32(_a20)
	v180 = *(*int32)(unsafe.Add(mBase, _consts[59]))
	v183 = v176 + v180<<(uint(int32(4))%32)
	*(*uint32)(unsafe.Add(mBase, uint32(v183)+8)) = uint32(v166)
	*(*int64)(unsafe.Add(mBase, uint32(v183))) = v159
	v186 = int32(1)
	*(*int32)(unsafe.Add(mBase, _consts[59])) = v180 + v186
	v192 = v147 + int32(2)
	if v192 < v114 {
		v147 = v192
		goto L41
	} else {
		goto L46
	}
L46:
	;
	v204 = v186
	goto L1
}
