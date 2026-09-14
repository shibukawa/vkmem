package p0

import (
	base "github.com/shibukawa/vkmem/internal/aot/vkaot/base"
	"unsafe"
)

func F___ofl_add(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	F___lock(m, int32(9116552))
	mBase = m.M
	v6 = int32(9116556)
	v7 = *(*int32)(unsafe.Add(mBase, _consts[1231]))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+56)) = v7
	if v7 == int32(0) {
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(v7)+52)) = l0
	}
	*(*int32)(unsafe.Add(mBase, _consts[1231])) = l0
	F___unlock(m, int32(9116552))
	mBase = m.M
	return l0
}
func F___openlog(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v1 int32
	_ = v1
	var v6 int32
	_ = v6
	var v12 int32
	_ = v12
	v1 = int32(0)
	v6 = F_socket(m, int32(1), int32(524290), v1)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, _consts[1263])) = v6
	if v6 < v1 {
	} else {
		v12 = F_connect(m, v6, int32(_a2377), int32(12))
		mBase = m.M
	}
	return
}
func F___overflow(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v18 int32
	_ = v18
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v37 int32
	_ = v37
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	v2 = l1
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	*(*uint8)(unsafe.Add(mBase, uint32(v8)+15)) = uint8(v2)
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v11 != 0 {
		v42 = v11
		v43 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
		if v43 == v42 {
			v57 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
			v58 = m.T0[v57].(func(*base.Module, int32, int32, int32) int32)(m, l0, v8+int32(15), int32(1))
			mBase = m.M
			v61 = m.ExcPending
			if v61 != 0 {
				return int32(0)
			} else {
				if v58 == int32(1) {
					v65 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+15)))
					v66 = v65
				} else {
					v66 = int32(-1)
				}
				m.G0 = v8 + int32(16)
				return v66
			}
		} else {
			v45 = *(*int32)(unsafe.Add(mBase, uint32(l0)+80))
			v47 = v2 & int32(255)
			if v45 == v47 {
				v57 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
				v58 = m.T0[v57].(func(*base.Module, int32, int32, int32) int32)(m, l0, v8+int32(15), int32(1))
				mBase = m.M
				v61 = m.ExcPending
				if v61 != 0 {
					return int32(0)
				} else {
					if v58 == int32(1) {
						v65 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+15)))
						v66 = v65
					} else {
						v66 = int32(-1)
					}
					m.G0 = v8 + int32(16)
					return v66
				}
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v43 + int32(1)
				*(*uint8)(unsafe.Add(mBase, uint32(v43))) = uint8(v2)
				v66 = v47
				m.G0 = v8 + int32(16)
				return v66
			}
		}
	} else {
		v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+72))
		*(*int32)(unsafe.Add(mBase, uint32(l0)+72)) = v13 + int32(-1) | v13
		v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		if v18&int32(8) == int32(0) {
			*(*int64)(unsafe.Add(mBase, uint32(l0)+4)) = int64(0)
			v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
			*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v29
			*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v29
			v32 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
			*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v29 + v32
			v37 = int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(l0))) = v18 | int32(32)
			v37 = int32(-1)
		}
		if v37 == int32(0) {
			v41 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
			v42 = v41
			v43 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
			if v43 == v42 {
				v57 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
				v58 = m.T0[v57].(func(*base.Module, int32, int32, int32) int32)(m, l0, v8+int32(15), int32(1))
				mBase = m.M
				v61 = m.ExcPending
				if v61 != 0 {
					return int32(0)
				} else {
					if v58 == int32(1) {
						v65 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+15)))
						v66 = v65
					} else {
						v66 = int32(-1)
					}
					m.G0 = v8 + int32(16)
					return v66
				}
			} else {
				v45 = *(*int32)(unsafe.Add(mBase, uint32(l0)+80))
				v47 = v2 & int32(255)
				if v45 == v47 {
					v57 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
					v58 = m.T0[v57].(func(*base.Module, int32, int32, int32) int32)(m, l0, v8+int32(15), int32(1))
					mBase = m.M
					v61 = m.ExcPending
					if v61 != 0 {
						return int32(0)
					} else {
						if v58 == int32(1) {
							v65 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+15)))
							v66 = v65
						} else {
							v66 = int32(-1)
						}
						m.G0 = v8 + int32(16)
						return v66
					}
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v43 + int32(1)
					*(*uint8)(unsafe.Add(mBase, uint32(v43))) = uint8(v2)
					v66 = v47
					m.G0 = v8 + int32(16)
					return v66
				}
			}
		} else {
			v66 = int32(-1)
			m.G0 = v8 + int32(16)
			return v66
		}
	}
}
func F_onMaxBatchSizeChange(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	v2 = int32(0)
	v4 = *(*int32)(unsafe.Add(mBase, _consts[338]))
	if v4 == v2 {
		F_freePrefetchCommandsBatch(m)
		mBase = m.M
		v11 = m.ExcPending
		if v11 != 0 {
			return int32(0)
		} else {
			F_prefetchCommandsBatchInit(m)
			mBase = m.M
			v13 = m.ExcPending
			if v13 != 0 {
				return int32(0)
			} else {
				return int32(1)
			}
		}
	} else {
		v7 = *(*int32)(unsafe.Add(mBase, uint32(v4)+12))
		if v7 != 0 {
			return int32(1)
		} else {
			F_freePrefetchCommandsBatch(m)
			mBase = m.M
			v11 = m.ExcPending
			if v11 != 0 {
				return int32(0)
			} else {
				F_prefetchCommandsBatchInit(m)
				mBase = m.M
				v13 = m.ExcPending
				if v13 != 0 {
					return int32(0)
				} else {
					return int32(1)
				}
			}
		}
	}
}
func F_openNewIncrAofForAppend(m *base.Module) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v27 int32
	_ = v27
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v46 int64
	_ = v46
	var v48 int64
	_ = v48
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v66 int64
	_ = v66
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
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
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v96 int32
	_ = v96
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v107 int32
	_ = v107
	var v110 int32
	_ = v110
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v118 int32
	_ = v118
	var v124 int32
	_ = v124
	var v133 int32
	_ = v133
	var v135 int32
	_ = v135
	var v137 int32
	_ = v137
	var v141 int64
	_ = v141
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v147 int64
	_ = v147
	var v149 int32
	_ = v149
	var v150 int64
	_ = v150
	var v159 int32
	_ = v159
	var v163 int32
	_ = v163
	var v166 int32
	_ = v166
	var v169 int32
	_ = v169
	var v173 int32
	_ = v173
	var v175 int32
	_ = v175
	var v177 int32
	_ = v177
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v186 int32
	_ = v186
	var v188 int32
	_ = v188
	var v197 int32
	_ = v197
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v208 int32
	_ = v208
	var v211 int32
	_ = v211
	var v214 int32
	_ = v214
	var v218 int32
	_ = v218
	var v220 int32
	_ = v220
	var v222 int32
	_ = v222
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v231 int32
	_ = v231
	var v233 int32
	_ = v233
	var v235 int32
	_ = v235
	v8 = m.G0
	v10 = v8 - int32(96)
	m.G0 = v10
	v13 = *(*int32)(unsafe.Add(mBase, _consts[12]))
	if v13 == int32(0) {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	m.G0 = v10 + int32(96)
	return v235
L2:
	;
	if v77 == int32(0) {
		goto L56
	} else {
		goto L57
	}
L3:
	;
	F__serverAssert(m, int32(_a89), int32(_a68), int32(775))
	mBase = m.M
	v197 = m.ExcPending
	if v197 != 0 {
		goto L8
	} else {
		goto L55
	}
L4:
	;
	v17 = *(*int32)(unsafe.Add(mBase, _consts[27]))
	switch v17 {
	case 0:
		v235 = v17
		goto L1
	default:
		goto L6
	case 2:
		goto L7
	}
L5:
	;
	v82 = *(*int32)(unsafe.Add(mBase, _consts[13]))
	v83 = F_makePath(m, v82, v77)
	mBase = m.M
	v84 = m.ExcPending
	if v84 != 0 {
		goto L8
	} else {
		goto L17
	}
L6:
	;
	v35 = F_aofManifestDup(m, v13)
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
		goto L8
	} else {
		goto L11
	}
L7:
	;
	v18 = F_sdsempty(m)
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	return int32(0)
L9:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+80)) = int32(_a78)
	*(*int32)(unsafe.Add(mBase, uint32(v10)+88)) = int32(_a76)
	v27 = *(*int32)(unsafe.Add(mBase, _consts[14]))
	*(*int32)(unsafe.Add(mBase, uint32(v10)+84)) = v27
	v33 = F_sdscatprintf(m, v18, int32(_a79), v10+int32(80))
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
		goto L8
	} else {
		goto L10
	}
L10:
	;
	v76 = int32(0)
	v77 = v33
	goto L5
L11:
	;
	v38 = F_valkey_calloc(m, int32(24))
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L8
	} else {
		goto L12
	}
L12:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v38)+16)) = int32(105)
	v42 = F_sdsempty(m)
	mBase = m.M
	v43 = m.ExcPending
	if v43 != 0 {
		goto L8
	} else {
		goto L13
	}
L13:
	;
	v45 = *(*int32)(unsafe.Add(mBase, _consts[14]))
	v46 = *(*int64)(unsafe.Add(mBase, uint32(v35)+24))
	v48 = v46 + int64(1)
	*(*int64)(unsafe.Add(mBase, uint32(v35)+24)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v10+int32(68)))) = int32(_a75)
	*(*int32)(unsafe.Add(mBase, uint32(v10+int32(64)))) = int32(_a76)
	*(*int64)(unsafe.Add(mBase, uint32(v10)+56)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v10)+48)) = v45
	v63 = F_sdscatprintf(m, v42, int32(_a77), v10+int32(48))
	mBase = m.M
	v64 = m.ExcPending
	if v64 != 0 {
		goto L8
	} else {
		goto L14
	}
L14:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v38))) = v63
	v66 = *(*int64)(unsafe.Add(mBase, uint32(v35)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v38)+8)) = v66
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v35)+4))
	v69 = F_listAddNodeTail(m, v68, v38)
	mBase = m.M
	v70 = m.ExcPending
	if v70 != 0 {
		goto L8
	} else {
		goto L15
	}
L15:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v35)+32)) = int32(1)
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v38)))
	v74 = F_sdsdup(m, v73)
	mBase = m.M
	v75 = m.ExcPending
	if v75 != 0 {
		goto L8
	} else {
		goto L16
	}
L16:
	;
	v76 = v35
	v77 = v74
	goto L5
L17:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+32)) = int32(438)
	v90 = F_open(m, v83, int32(577), v10+int32(32))
	mBase = m.M
	F_sdsfree(m, v83)
	mBase = m.M
	v92 = m.ExcPending
	if v92 != 0 {
		goto L8
	} else {
		goto L18
	}
L18:
	;
	if v90 != int32(-1) {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	if v76 == int32(0) {
		goto L25
	} else {
		goto L26
	}
L20:
	;
	v96 = *(*int32)(unsafe.Add(mBase, _consts[15]))
	if int32(3) < v96 {
		goto L2
	} else {
		goto L21
	}
L21:
	;
	goto L22
L22:
	;
	v100 = *(*int32)(unsafe.Add(mBase, _consts[9]))
	v101 = F___strerror_l(m, v100, v100)
	mBase = m.M
	goto L23
L23:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+4)) = v101
	*(*int32)(unsafe.Add(mBase, uint32(v10))) = v77
	F__serverLog(m, int32(3), int32(_a91), v10)
	mBase = m.M
	v107 = m.ExcPending
	if v107 != 0 {
		goto L8
	} else {
		goto L24
	}
L24:
	;
	goto L2
L25:
	;
	v124 = *(*int32)(unsafe.Add(mBase, _consts[15]))
	if int32(2) < v124 {
		goto L32
	} else {
		goto L33
	}
L26:
	;
	v110 = *(*int32)(unsafe.Add(mBase, uint32(v76)+32))
	if v110 == int32(0) {
		goto L25
	} else {
		goto L27
	}
L27:
	;
	v113 = F_getAofManifestAsString(m, v76)
	mBase = m.M
	v114 = m.ExcPending
	if v114 != 0 {
		goto L8
	} else {
		goto L28
	}
L28:
	;
	v115 = F_writeAofManifestFile(m, v113)
	mBase = m.M
	v116 = m.ExcPending
	if v116 != 0 {
		goto L8
	} else {
		goto L29
	}
L29:
	;
	F_sdsfree(m, v113)
	mBase = m.M
	v118 = m.ExcPending
	if v118 != 0 {
		goto L8
	} else {
		goto L30
	}
L30:
	;
	if v115 != 0 {
		goto L2
	} else {
		goto L31
	}
L31:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v76)+32)) = int32(0)
	goto L25
L32:
	;
	F_sdsfree(m, v77)
	mBase = m.M
	v135 = m.ExcPending
	if v135 != 0 {
		goto L8
	} else {
		goto L35
	}
L33:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = v77
	F__serverLog(m, int32(2), int32(_a90), v10+int32(16))
	mBase = m.M
	v133 = m.ExcPending
	if v133 != 0 {
		goto L8
	} else {
		goto L34
	}
L34:
	;
	goto L32
L35:
	;
	v137 = *(*int32)(unsafe.Add(mBase, _consts[28]))
	if v137 == int32(-1) {
		goto L36
	} else {
		goto L37
	}
L36:
	;
	v149 = int32(_a69)
	v150 = int64(0)
	*(*int64)(unsafe.Add(mBase, _consts[29])) = v150
	*(*int32)(unsafe.Add(mBase, _consts[28])) = v90
	*(*int64)(unsafe.Add(mBase, _consts[30])) = v150
	v159 = int32(0)
	if v76 == v159 {
		v235 = v159
		goto L1
	} else {
		goto L39
	}
L37:
	;
	v141 = *(*int64)(unsafe.Add(mBase, _consts[31]))
	F_bioCreateCloseAofJob(m, v137, v141, int32(1))
	mBase = m.M
	v144 = m.ExcPending
	if v144 != 0 {
		goto L8
	} else {
		goto L38
	}
L38:
	;
	v145 = int32(_a69)
	v147 = *(*int64)(unsafe.Add(mBase, _consts[32]))
	*(*int64)(unsafe.Add(mBase, _consts[33])) = v147
	goto L36
L39:
	;
	v163 = *(*int32)(unsafe.Add(mBase, _consts[12]))
	if v163 == int32(0) {
		goto L40
	} else {
		goto L41
	}
L40:
	;
	*(*int32)(unsafe.Add(mBase, _consts[12])) = v76
	v235 = v159
	goto L1
L41:
	;
	v166 = *(*int32)(unsafe.Add(mBase, uint32(v163)))
	if v166 == int32(0) {
		goto L42
	} else {
		goto L43
	}
L42:
	;
	v177 = *(*int32)(unsafe.Add(mBase, uint32(v163)+4))
	if v177 == int32(0) {
		goto L48
	} else {
		goto L49
	}
L43:
	;
	v169 = *(*int32)(unsafe.Add(mBase, uint32(v166)))
	if v169 == int32(0) {
		goto L44
	} else {
		goto L45
	}
L44:
	;
	F_valkey_free(m, v166)
	mBase = m.M
	v175 = m.ExcPending
	if v175 != 0 {
		goto L8
	} else {
		goto L47
	}
L45:
	;
	F_sdsfree(m, v169)
	mBase = m.M
	v173 = m.ExcPending
	if v173 != 0 {
		goto L8
	} else {
		goto L46
	}
L46:
	;
	goto L44
L47:
	;
	goto L42
L48:
	;
	v182 = *(*int32)(unsafe.Add(mBase, uint32(v163)+8))
	if v182 == int32(0) {
		goto L51
	} else {
		goto L52
	}
L49:
	;
	F_listRelease(m, v177)
	mBase = m.M
	v181 = m.ExcPending
	if v181 != 0 {
		goto L8
	} else {
		goto L50
	}
L50:
	;
	goto L48
L51:
	;
	F_valkey_free(m, v163)
	mBase = m.M
	v188 = m.ExcPending
	if v188 != 0 {
		goto L8
	} else {
		goto L54
	}
L52:
	;
	F_listRelease(m, v182)
	mBase = m.M
	v186 = m.ExcPending
	if v186 != 0 {
		goto L8
	} else {
		goto L53
	}
L53:
	;
	goto L51
L54:
	;
	goto L40
L55:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L56:
	;
	v205 = int32(-1)
	if v90 == v205 {
		goto L59
	} else {
		goto L60
	}
L57:
	;
	F_sdsfree(m, v77)
	mBase = m.M
	v204 = m.ExcPending
	if v204 != 0 {
		goto L8
	} else {
		goto L58
	}
L58:
	;
	goto L56
L59:
	;
	if v76 == int32(0) {
		v235 = v205
		goto L1
	} else {
		goto L61
	}
L60:
	;
	v208 = F_close(m, v90)
	mBase = m.M
	goto L59
L61:
	;
	v211 = *(*int32)(unsafe.Add(mBase, uint32(v76)))
	if v211 == int32(0) {
		goto L62
	} else {
		goto L63
	}
L62:
	;
	v222 = *(*int32)(unsafe.Add(mBase, uint32(v76)+4))
	if v222 == int32(0) {
		goto L68
	} else {
		goto L69
	}
L63:
	;
	v214 = *(*int32)(unsafe.Add(mBase, uint32(v211)))
	if v214 == int32(0) {
		goto L64
	} else {
		goto L65
	}
L64:
	;
	F_valkey_free(m, v211)
	mBase = m.M
	v220 = m.ExcPending
	if v220 != 0 {
		goto L8
	} else {
		goto L67
	}
L65:
	;
	F_sdsfree(m, v214)
	mBase = m.M
	v218 = m.ExcPending
	if v218 != 0 {
		goto L8
	} else {
		goto L66
	}
L66:
	;
	goto L64
L67:
	;
	goto L62
L68:
	;
	v227 = *(*int32)(unsafe.Add(mBase, uint32(v76)+8))
	if v227 == int32(0) {
		goto L71
	} else {
		goto L72
	}
L69:
	;
	F_listRelease(m, v222)
	mBase = m.M
	v226 = m.ExcPending
	if v226 != 0 {
		goto L8
	} else {
		goto L70
	}
L70:
	;
	goto L68
L71:
	;
	F_valkey_free(m, v76)
	mBase = m.M
	v233 = m.ExcPending
	if v233 != 0 {
		goto L8
	} else {
		goto L74
	}
L72:
	;
	F_listRelease(m, v227)
	mBase = m.M
	v231 = m.ExcPending
	if v231 != 0 {
		goto L8
	} else {
		goto L73
	}
L73:
	;
	goto L71
L74:
	;
	v235 = v205
	goto L1
}
func F_optsize(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v43 int32
	_ = v43
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
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v90 int32
	_ = v90
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v115 int32
	_ = v115
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v125 int32
	_ = v125
	v9 = m.G0
	v11 = v9 - int32(16)
	m.G0 = v11
	v13 = int32(1)
	switch l1 + int32(-66) {
	case 0, 32, 54:
		v125 = v13
		goto L1
	default:
		goto L2
	case 6, 38:
		goto L7
	case 7, 39:
		goto L3
	case 10, 18, 36, 42:
		goto L6
	case 33:
		goto L4
	case 34:
		goto L5
	}
L1:
	;
	m.G0 = v11 + int32(16)
	return v125
L2:
	;
	v125 = int32(0)
	goto L1
L3:
	;
	v66 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	v67 = int32(*(*int8)(unsafe.Add(mBase, uint32(v66))))
	if base.Ui32(int32(9)) < base.Ui32(v67+int32(-48)) {
		v125 = int32(4)
		goto L1
	} else {
		goto L18
	}
L4:
	;
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	v20 = int32(*(*int8)(unsafe.Add(mBase, uint32(v19))))
	if base.Ui32(int32(9)) < base.Ui32(v20+int32(-48)) {
		v125 = v13
		goto L1
	} else {
		goto L8
	}
L5:
	;
	v125 = int32(8)
	goto L1
L6:
	;
	v125 = int32(4)
	goto L1
L7:
	;
	v125 = int32(2)
	goto L1
L8:
	;
	v30 = int32(0)
	v31 = v19
	v32 = v20
	goto L9
L9:
	;
	v35 = v30 * int32(10)
	if int32(214748364) < v30 {
		goto L12
	} else {
		goto L13
	}
L11:
	;
	v54 = v52 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v54
	v56 = int32(*(*int8)(unsafe.Add(mBase, uint32(v52))))
	v58 = int32(-48)
	v59 = v35 + v56 + v58
	v60 = int32(*(*int8)(unsafe.Add(mBase, uint32(v52)+1)))
	if base.Ui32(v60+v58) < base.Ui32(int32(10)) {
		v30 = v59
		v31 = v54
		v32 = v60
		goto L9
	} else {
		goto L17
	}
L12:
	;
	v43 = m.G3
	v47 = F_luaL_error(m, l0, v43+int32(_a2337), int32(0))
	mBase = m.M
	v50 = m.ExcPending
	if v50 != 0 {
		goto L15
	} else {
		goto L16
	}
L13:
	;
	if v35 <= int32(-2147483601)-base.I32_extend8_s(v32) {
		v52 = v31
		goto L11
	} else {
		goto L14
	}
L14:
	;
	goto L12
L15:
	;
	return int32(0)
L16:
	;
	v51 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	v52 = v51
	goto L11
L17:
	;
	v125 = v59
	goto L1
L18:
	;
	v77 = int32(0)
	v78 = v66
	v79 = v67
	goto L19
L19:
	;
	v82 = v77 * int32(10)
	if int32(214748364) < v77 {
		goto L22
	} else {
		goto L23
	}
L20:
	;
	if v104 < int32(33) {
		v125 = v104
		goto L1
	} else {
		goto L27
	}
L21:
	;
	v99 = v97 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v99
	v101 = int32(*(*int8)(unsafe.Add(mBase, uint32(v97))))
	v103 = int32(-48)
	v104 = v82 + v101 + v103
	v105 = int32(*(*int8)(unsafe.Add(mBase, uint32(v97)+1)))
	if base.Ui32(v105+v103) < base.Ui32(int32(10)) {
		v77 = v104
		v78 = v99
		v79 = v105
		goto L19
	} else {
		goto L26
	}
L22:
	;
	v90 = m.G3
	v94 = F_luaL_error(m, l0, v90+int32(_a2337), int32(0))
	mBase = m.M
	v95 = m.ExcPending
	if v95 != 0 {
		goto L15
	} else {
		goto L25
	}
L23:
	;
	if v82 <= int32(-2147483601)-base.I32_extend8_s(v79) {
		v97 = v78
		goto L21
	} else {
		goto L24
	}
L24:
	;
	goto L22
L25:
	;
	v96 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	v97 = v96
	goto L21
L26:
	;
	goto L20
L27:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+4)) = int32(32)
	*(*int32)(unsafe.Add(mBase, uint32(v11))) = v104
	v115 = m.G3
	v118 = F_luaL_error(m, l0, v115+int32(_a2338), v11)
	mBase = m.M
	v119 = m.ExcPending
	if v119 != 0 {
		goto L15
	} else {
		goto L28
	}
L28:
	;
	v125 = v104
	goto L1
}
