package p0

import (
	base "github.com/shibukawa/vkmem/internal/aot/vkaot/base"
	"unsafe"
)

func F_freeFunctionsAsync(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v9 = *(*int32)(unsafe.Add(mBase, uint32(v8)+16))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(v8)+12))
	if base.Ui32(v9+v10) < base.Ui32(int32(65)) {
		F_functionsLibCtxFree(m, l0, int32(0), l1)
		mBase = m.M
		v31 = m.ExcPending
		if v31 != 0 {
			return
		} else {
			m.G0 = v6 + int32(16)
			return
		}
	} else {
		v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		v16 = *(*int32)(unsafe.Add(mBase, uint32(v15)+16))
		v17 = *(*int32)(unsafe.Add(mBase, uint32(v15)+12))
		v20 = *(*int32)(unsafe.Add(mBase, _consts[323]))
		*(*int32)(unsafe.Add(mBase, _consts[323])) = v16 + v17 + v20
		*(*int32)(unsafe.Add(mBase, uint32(v6))) = l0
		*(*int32)(unsafe.Add(mBase, uint32(v6)+4)) = l1
		F_bioCreateLazyFreeJob(m, int32(556), int32(2), v6)
		mBase = m.M
		v28 = m.ExcPending
		if v28 != 0 {
			return
		} else {
			m.G0 = v6 + int32(16)
			return
		}
	}
}
func F_functionsLibCtxClear(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
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
	var v12 int32
	_ = v12
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
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
	var v51 int64
	_ = v51
	var v53 int64
	_ = v53
	var v55 int64
	_ = v55
	var v57 int64
	_ = v57
	var v59 int64
	_ = v59
	var v61 int64
	_ = v61
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v83 int32
	_ = v83
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v103 int32
	_ = v103
	var v108 int32
	_ = v108
	var v112 int32
	_ = v112
	var v116 int32
	_ = v116
	var v123 int32
	_ = v123
	var v125 int32
	_ = v125
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v139 int32
	_ = v139
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v147 int32
	_ = v147
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v156 int32
	_ = v156
	var v158 int32
	_ = v158
	var v160 int64
	_ = v160
	var v161 int64
	_ = v161
	var v162 int64
	_ = v162
	var v163 int64
	_ = v163
	var v164 int64
	_ = v164
	var v165 int64
	_ = v165
	var v166 int64
	_ = v166
	var v168 int64
	_ = v168
	var v170 int64
	_ = v170
	var v172 int64
	_ = v172
	var v174 int64
	_ = v174
	var v176 int64
	_ = v176
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v183 int32
	_ = v183
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v190 int32
	_ = v190
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v198 int32
	_ = v198
	var v202 int32
	_ = v202
	var v204 int32
	_ = v204
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v218 int32
	_ = v218
	var v223 int32
	_ = v223
	var v227 int32
	_ = v227
	var v231 int32
	_ = v231
	var v239 int32
	_ = v239
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	F_dictEmpty(m, v4, l1)
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	F_dictEmpty(m, v7, l1)
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v11 = F_dictGetIterator(m, v10)
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		goto L1
	} else {
		goto L5
	}
L4:
	;
	F_dictReleaseIterator(m, v11)
	mBase = m.M
	v239 = m.ExcPending
	if v239 != 0 {
		goto L1
	} else {
		goto L63
	}
L5:
	;
	v20 = v11 + int32(20)
	v21 = *(*int32)(unsafe.Add(mBase, uint32(v11)+16))
	if v21 != 0 {
		goto L8
	} else {
		goto L9
	}
L6:
	;
	if v116 == int32(0) {
		goto L4
	} else {
		goto L32
	}
L7:
	;
	v27 = v20
	v28 = v24
	goto L10
L8:
	;
	v24 = int32(1)
	goto L7
L9:
	;
	v24 = int32(0)
	goto L7
L10:
	;
	switch v28 {
	case 0:
		goto L15
	default:
		goto L14
	}
L12:
	;
	v28 = int32(0)
	goto L10
L13:
	;
	goto L6
L14:
	;
	v108 = *(*int32)(unsafe.Add(mBase, uint32(v27)))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = v108
	if v108 == int32(0) {
		goto L12
	} else {
		goto L31
	}
L15:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
	if v32 != int32(-1) {
		v71 = v32
		goto L16
	} else {
		goto L17
	}
L16:
	;
	v72 = int32(1)
	v73 = v71 + v72
	*(*int32)(unsafe.Add(mBase, uint32(v11)+4)) = v73
	v75 = int32(0)
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
	v79 = *(*int32)(unsafe.Add(mBase, uint32(v11)+8))
	v83 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v78+v79+int32(26)))))
	if v83 == int32(255) {
		goto L25
	} else {
		goto L26
	}
L17:
	;
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v11)+8))
	if v36 != 0 {
		v71 = int32(-1)
		goto L16
	} else {
		goto L18
	}
L18:
	;
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v11)+12))
	if v38 == int32(0) {
		goto L20
	} else {
		goto L21
	}
L19:
	;
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v64)+20))
	if v65 != int32(-1) {
		goto L22
	} else {
		goto L23
	}
L20:
	;
	v45 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v37)+16)))
	v46 = int64(*(*int8)(unsafe.Add(mBase, uint32(v37)+27)))
	v47 = int64(*(*int32)(unsafe.Add(mBase, uint32(v37)+8)))
	v48 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v37)+12)))
	v49 = int64(*(*int8)(unsafe.Add(mBase, uint32(v37)+26)))
	v50 = int64(*(*int32)(unsafe.Add(mBase, uint32(v37)+4)))
	v51 = F_wangHash64(m, v50)
	mBase = m.M
	v53 = F_wangHash64(m, v49+v51)
	mBase = m.M
	v55 = F_wangHash64(m, v48+v53)
	mBase = m.M
	v57 = F_wangHash64(m, v47+v55)
	mBase = m.M
	v59 = F_wangHash64(m, v46+v57)
	mBase = m.M
	v61 = F_wangHash64(m, v45+v59)
	mBase = m.M
	*(*int64)(unsafe.Add(mBase, uint32(v11)+24)) = v61
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
	v64 = v63
	goto L19
L21:
	;
	v41 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v37)+24)))
	v43 = v41 + int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v37)+24)) = uint16(v43)
	v64 = v37
	goto L19
L22:
	;
	v71 = v65 + int32(-1)
	goto L16
L23:
	;
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
	v71 = v68
	goto L16
L24:
	;
	v98 = int32(2)
	v103 = *(*int32)(unsafe.Add(mBase, uint32(v78+v96<<(uint(v98)%32)+int32(4))))
	v27 = v103 + v97<<(uint(v98)%32)
	v28 = int32(1)
	goto L10
L25:
	;
	v87 = v75
	goto L27
L26:
	;
	v87 = v72 << (uint(v83) % 32)
	goto L27
L27:
	;
	if v73 < v87 {
		v96 = v79
		v97 = v73
		goto L24
	} else {
		goto L28
	}
L28:
	;
	if v79 != 0 {
		v116 = v75
		goto L13
	} else {
		goto L29
	}
L29:
	;
	v89 = *(*int32)(unsafe.Add(mBase, uint32(v78)+20))
	if v89 == int32(-1) {
		v116 = v75
		goto L13
	} else {
		goto L30
	}
L30:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v11)+4)) = int64(4294967296)
	v96 = int32(1)
	v97 = int32(0)
	goto L24
L31:
	;
	v112 = *(*int32)(unsafe.Add(mBase, uint32(v108)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v20))) = v112
	v116 = v108
	goto L13
L32:
	;
	v123 = v116
	goto L33
L33:
	;
	v125 = *(*int32)(unsafe.Add(mBase, uint32(v123)+8))
	goto L35
L34:
	;
	goto L4
L35:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v125))) = int64(0)
	v135 = v11 + int32(20)
	v136 = *(*int32)(unsafe.Add(mBase, uint32(v11)+16))
	if v136 != 0 {
		goto L38
	} else {
		goto L39
	}
L36:
	;
	if v231 != 0 {
		v123 = v231
		goto L33
	} else {
		goto L62
	}
L37:
	;
	v142 = v135
	v143 = v139
	goto L40
L38:
	;
	v139 = int32(1)
	goto L37
L39:
	;
	v139 = int32(0)
	goto L37
L40:
	;
	switch v143 {
	case 0:
		goto L45
	default:
		goto L44
	}
L42:
	;
	v143 = int32(0)
	goto L40
L43:
	;
	goto L36
L44:
	;
	v223 = *(*int32)(unsafe.Add(mBase, uint32(v142)))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = v223
	if v223 == int32(0) {
		goto L42
	} else {
		goto L61
	}
L45:
	;
	v147 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
	if v147 != int32(-1) {
		v186 = v147
		goto L46
	} else {
		goto L47
	}
L46:
	;
	v187 = int32(1)
	v188 = v186 + v187
	*(*int32)(unsafe.Add(mBase, uint32(v11)+4)) = v188
	v190 = int32(0)
	v193 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
	v194 = *(*int32)(unsafe.Add(mBase, uint32(v11)+8))
	v198 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v193+v194+int32(26)))))
	if v198 == int32(255) {
		goto L55
	} else {
		goto L56
	}
L47:
	;
	v151 = *(*int32)(unsafe.Add(mBase, uint32(v11)+8))
	if v151 != 0 {
		v186 = int32(-1)
		goto L46
	} else {
		goto L48
	}
L48:
	;
	v152 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
	v153 = *(*int32)(unsafe.Add(mBase, uint32(v11)+12))
	if v153 == int32(0) {
		goto L50
	} else {
		goto L51
	}
L49:
	;
	v180 = *(*int32)(unsafe.Add(mBase, uint32(v179)+20))
	if v180 != int32(-1) {
		goto L52
	} else {
		goto L53
	}
L50:
	;
	v160 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v152)+16)))
	v161 = int64(*(*int8)(unsafe.Add(mBase, uint32(v152)+27)))
	v162 = int64(*(*int32)(unsafe.Add(mBase, uint32(v152)+8)))
	v163 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v152)+12)))
	v164 = int64(*(*int8)(unsafe.Add(mBase, uint32(v152)+26)))
	v165 = int64(*(*int32)(unsafe.Add(mBase, uint32(v152)+4)))
	v166 = F_wangHash64(m, v165)
	mBase = m.M
	v168 = F_wangHash64(m, v164+v166)
	mBase = m.M
	v170 = F_wangHash64(m, v163+v168)
	mBase = m.M
	v172 = F_wangHash64(m, v162+v170)
	mBase = m.M
	v174 = F_wangHash64(m, v161+v172)
	mBase = m.M
	v176 = F_wangHash64(m, v160+v174)
	mBase = m.M
	*(*int64)(unsafe.Add(mBase, uint32(v11)+24)) = v176
	v178 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
	v179 = v178
	goto L49
L51:
	;
	v156 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v152)+24)))
	v158 = v156 + int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v152)+24)) = uint16(v158)
	v179 = v152
	goto L49
L52:
	;
	v186 = v180 + int32(-1)
	goto L46
L53:
	;
	v183 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
	v186 = v183
	goto L46
L54:
	;
	v213 = int32(2)
	v218 = *(*int32)(unsafe.Add(mBase, uint32(v193+v211<<(uint(v213)%32)+int32(4))))
	v142 = v218 + v212<<(uint(v213)%32)
	v143 = int32(1)
	goto L40
L55:
	;
	v202 = v190
	goto L57
L56:
	;
	v202 = v187 << (uint(v198) % 32)
	goto L57
L57:
	;
	if v188 < v202 {
		v211 = v194
		v212 = v188
		goto L54
	} else {
		goto L58
	}
L58:
	;
	if v194 != 0 {
		v231 = v190
		goto L43
	} else {
		goto L59
	}
L59:
	;
	v204 = *(*int32)(unsafe.Add(mBase, uint32(v193)+20))
	if v204 == int32(-1) {
		v231 = v190
		goto L43
	} else {
		goto L60
	}
L60:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v11)+4)) = int64(4294967296)
	v211 = int32(1)
	v212 = int32(0)
	goto L54
L61:
	;
	v227 = *(*int32)(unsafe.Add(mBase, uint32(v223)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v135))) = v227
	v231 = v223
	goto L43
L62:
	;
	goto L34
L63:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = int32(0)
	return
}
func F_functionsLibCtxCreate(m *base.Module) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
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
	var v21 int32
	_ = v21
	v3 = F_valkey_malloc(m, int32(16))
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		v8 = F_dictCreate(m, int32(_a577))
		mBase = m.M
		v9 = m.ExcPending
		if v9 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v3))) = v8
			v12 = F_dictCreate(m, int32(_a578))
			mBase = m.M
			v13 = m.ExcPending
			if v13 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v3)+4)) = v12
				v16 = F_dictCreate(m, int32(_a579))
				mBase = m.M
				v17 = m.ExcPending
				if v17 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v3)+12)) = v16
					F_scriptingEngineManagerForEachEngine(m, int32(529), v3)
					mBase = m.M
					v21 = m.ExcPending
					if v21 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v3)+8)) = int32(0)
						return v3
					}
				}
			}
		}
	}
}
func F_functionsLibCtxFunctionsLen(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v3 = *(*int32)(unsafe.Add(mBase, uint32(v2)+16))
	v4 = *(*int32)(unsafe.Add(mBase, uint32(v2)+12))
	return v3 + v4
}
func F_functionsLibCtxGetCurrent(m *base.Module) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	v2 = *(*int32)(unsafe.Add(mBase, _consts[18]))
	return v2
}
func F_functionsLibNum(m *base.Module) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	v3 = *(*int32)(unsafe.Add(mBase, _consts[18]))
	v4 = *(*int32)(unsafe.Add(mBase, uint32(v3)))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(v4)+16))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(v4)+12))
	return v5 + v6
}
func F_functionsNum(m *base.Module) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	v3 = *(*int32)(unsafe.Add(mBase, _consts[18]))
	v4 = *(*int32)(unsafe.Add(mBase, uint32(v3)+4))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(v4)+16))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(v4)+12))
	return v5 + v6
}
func F_initializeFunctionsLibEngineStats(m *base.Module, l0 int32, l1 int32) {
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
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	v5 = F_valkey_calloc(m, int32(8))
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return
	} else {
		v7 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
		v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		v9 = F_dictAdd(m, v7, v8, v5)
		mBase = m.M
		v10 = m.ExcPending
		if v10 != 0 {
			return
		} else {
			return
		}
	}
}
