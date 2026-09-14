package p1

import (
	base "github.com/shibukawa/vkmem/internal/aot/vkaot/base"
	"math"
	"unsafe"
)

func F_ValkeyModuleCommandDispatcher(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v5 int64
	_ = v5
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
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v53 int64
	_ = v53
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v63 int64
	_ = v63
	var v67 int64
	_ = v67
	var v72 int32
	_ = v72
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
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v91 int32
	_ = v91
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v105 int32
	_ = v105
	v2 = int32(0)
	v5 = int64(0)
	v7 = m.G0
	v9 = v7 - int32(80)
	m.G0 = v9
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(v11)+208))
	v13 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
	*(*int64)(unsafe.Add(mBase, uint32(v9+int32(32)))) = v5
	*(*int64)(unsafe.Add(mBase, uint32(v9+int32(72)))) = v5
	*(*int64)(unsafe.Add(mBase, uint32(v9+int32(64)))) = v5
	*(*int64)(unsafe.Add(mBase, uint32(v9+int32(56)))) = v5
	*(*int64)(unsafe.Add(mBase, uint32(v9+int32(48)))) = v5
	*(*int64)(unsafe.Add(mBase, uint32(v9+int32(40)))) = v5
	*(*int64)(unsafe.Add(mBase, uint32(v9+int32(24)))) = v5
	*(*int64)(unsafe.Add(mBase, uint32(v9)+16)) = v5
	*(*int32)(unsafe.Add(mBase, uint32(v9)+12)) = v13
	*(*int32)(unsafe.Add(mBase, uint32(v9)+8)) = int32(561)
	*(*int32)(unsafe.Add(mBase, uint32(v9)+36)) = int32(512)
	v50 = *(*int32)(unsafe.Add(mBase, _c_F_ValkeyModuleCommandDispatcher[0]))
	v52 = *(*int32)(unsafe.Add(mBase, _c_F_ValkeyModuleCommandDispatcher[1]))
	v53 = m.T0[v52].(func(*base.Module) int64)(m)
	mBase = m.M
	if v50 == v2 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v9)+64)) = v67
	*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = l0
	v72 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v73 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v74 = *(*int32)(unsafe.Add(mBase, uint32(v12)+4))
	v75 = m.T0[v74].(func(*base.Module, int32, int32, int32) int32)(m, v9+int32(8), v72, v73)
	mBase = m.M
	v76 = m.ExcPending
	if v76 != 0 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	v63 = *(*int64)(unsafe.Add(mBase, _c_F_ValkeyModuleCommandDispatcher[2]))
	v67 = v63*int64(1000) + v53
	goto L1
L3:
	;
	v58 = *(*int32)(unsafe.Add(mBase, _c_F_ValkeyModuleCommandDispatcher[3]))
	v59 = base.I32_div_s(int32(1000000), v58)
	v67 = v53 + base.I64_extend_i32_s(v59)
	goto L1
L4:
	;
	return
L5:
	;
	F_moduleFreeContext(m, v9+int32(8))
	mBase = m.M
	v80 = m.ExcPending
	if v80 != 0 {
		goto L4
	} else {
		goto L6
	}
L6:
	;
	v81 = int32(0)
	v82 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v82 <= v81 {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	m.G0 = v9 + int32(80)
	return
L8:
	;
	v87 = v81
	v88 = v82
	goto L9
L9:
	;
	v91 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v95 = *(*int32)(unsafe.Add(mBase, uint32(v91+v87<<(uint(int32(2))%32))))
	v96 = *(*int32)(unsafe.Add(mBase, uint32(v95)+4))
	if base.Ui32(v96) < base.Ui32(int32(16)) {
		v103 = v88
		goto L11
	} else {
		goto L12
	}
L10:
	;
	goto L7
L11:
	;
	v105 = v87 + int32(1)
	if v105 < v103 {
		v87 = v105
		v88 = v103
		goto L9
	} else {
		goto L14
	}
L12:
	;
	F_trimStringObjectIfNeeded(m, v95, int32(0))
	mBase = m.M
	v101 = m.ExcPending
	if v101 != 0 {
		goto L4
	} else {
		goto L13
	}
L13:
	;
	v102 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v103 = v102
	goto L11
L14:
	;
	goto L10
}
func F_genValkeyInfoStringCommandStats(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v37 int32
	_ = v37
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v58 int32
	_ = v58
	var v73 int32
	_ = v73
	var v74 int64
	_ = v74
	var v77 int64
	_ = v77
	var v80 int64
	_ = v80
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v88 int32
	_ = v88
	var v95 int32
	_ = v95
	var v98 int32
	_ = v98
	var v101 int32
	_ = v101
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v119 int32
	_ = v119
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v135 int32
	_ = v135
	var v139 int32
	_ = v139
	var v142 int32
	_ = v142
	var v153 int32
	_ = v153
	var v170 int32
	_ = v170
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v183 int32
	_ = v183
	var v200 int32
	_ = v200
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v217 int32
	_ = v217
	var v221 int32
	_ = v221
	var v224 int32
	_ = v224
	var v227 int32
	_ = v227
	var v239 int32
	_ = v239
	var v250 int32
	_ = v250
	var v251 int32
	_ = v251
	var v252 int64
	_ = v252
	var v253 int64
	_ = v253
	var v263 float64
	_ = v263
	var v264 int64
	_ = v264
	var v265 int64
	_ = v265
	var v273 int32
	_ = v273
	var v274 int32
	_ = v274
	var v278 int32
	_ = v278
	var v279 int32
	_ = v279
	var v288 int32
	_ = v288
	var v291 int32
	_ = v291
	var v292 int32
	_ = v292
	var v293 int32
	_ = v293
	var v298 int32
	_ = v298
	var v299 int32
	_ = v299
	var v300 int32
	_ = v300
	var v318 int32
	_ = v318
	v3 = int32(0)
	v16 = m.G0
	v18 = v16 - int32(112)
	m.G0 = v18
	v21 = v18 + int32(64)
	v22 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v21)+14)) = uint8(v3)
	*(*int32)(unsafe.Add(mBase, uint32(v21))) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v21)+24)) = v3
	*(*uint8)(unsafe.Add(mBase, uint32(v21)+15)) = uint8(v22)
	*(*int32)(unsafe.Add(mBase, uint32(v21)+8)) = int32(-1)
	if l1 == v3 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	v44 = F_hashtableNext(m, v18+int32(64), v18+int32(60))
	mBase = m.M
	v47 = m.ExcPending
	if v47 != 0 {
		goto L6
	} else {
		goto L7
	}
L2:
	;
	goto L1
L3:
	;
	goto L4
L4:
	;
	v37 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v21)+24)) = v37
	*(*int32)(unsafe.Add(mBase, uint32(l1)+40)) = v21
	goto L2
L5:
	;
	F_hashtableCleanupIterator(m, v18+int32(64))
	mBase = m.M
	v318 = m.ExcPending
	if v318 != 0 {
		goto L6
	} else {
		goto L65
	}
L6:
	;
	return int32(0)
L7:
	;
	if v44 == int32(0) {
		v300 = l0
		goto L5
	} else {
		goto L8
	}
L8:
	;
	v58 = l0
	goto L9
L9:
	;
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v18)+60))
	v74 = *(*int64)(unsafe.Add(mBase, uint32(v73)+112))
	if v74 != int64(0) {
		goto L12
	} else {
		goto L13
	}
L10:
	;
	v300 = v293
	goto L5
L11:
	;
	v288 = *(*int32)(unsafe.Add(mBase, uint32(v73)+200))
	if v288 == int32(0) {
		v293 = v279
		goto L60
	} else {
		goto L61
	}
L12:
	;
	v83 = int32(0)
	v85 = *(*int32)(unsafe.Add(mBase, uint32(v73)+140))
	v88 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v85+int32(-1)))))
	switch v88 & int32(7) {
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
		v105 = v83
		goto L16
	}
L13:
	;
	v77 = *(*int64)(unsafe.Add(mBase, uint32(v73)+128))
	if v77 != int64(0) {
		goto L12
	} else {
		goto L14
	}
L14:
	;
	v80 = *(*int64)(unsafe.Add(mBase, uint32(v73)+120))
	if v80 == int64(0) {
		v279 = v58
		goto L11
	} else {
		goto L15
	}
L15:
	;
	goto L12
L16:
	;
	if v105 == int32(0) {
		goto L25
	} else {
		goto L26
	}
L17:
	;
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v85+int32(-17))))
	v105 = v104
	goto L16
L18:
	;
	v101 = *(*int32)(unsafe.Add(mBase, uint32(v85+int32(-9))))
	v105 = v101
	goto L16
L19:
	;
	v98 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v85+int32(-5)))))
	v105 = v98
	goto L16
L20:
	;
	v95 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v85+int32(-3)))))
	v105 = v95
	goto L16
L21:
	;
	v105 = int32(base.Ui32(v88) >> (uint(int32(3)) % 32))
	goto L16
L22:
	;
	v252 = *(*int64)(unsafe.Add(mBase, uint32(v73)+104))
	v253 = *(*int64)(unsafe.Add(mBase, uint32(v73)+112))
	if base.B2i32(v253 == int64(0)) == int32(0) {
		goto L55
	} else {
		goto L56
	}
L23:
	;
	if v170 == int32(0) {
		v250 = v83
		v251 = v85
		goto L22
	} else {
		goto L36
	}
L24:
	;
	goto L23
L25:
	;
	v170 = int32(0)
	goto L24
L26:
	;
	v119 = int32(0)
	goto L27
L27:
	;
	goto L30
L28:
	;
	goto L25
L29:
	;
	v153 = v119 + int32(1)
	if v153 != v105 {
		v119 = v153
		goto L27
	} else {
		goto L35
	}
L30:
	;
	v126 = v85 + v119
	v127 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v126))))
	v135 = int32(0)
	goto L31
L31:
	;
	v139 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v135)+uint32(_c_F_genValkeyInfoStringCommandStats[0]))))
	if v127&int32(255) == v139 {
		v170 = v126
		goto L24
	} else {
		goto L33
	}
L32:
	;
	goto L29
L33:
	;
	v142 = v135 + int32(1)
	if v142 != int32(4) {
		v135 = v142
		goto L31
	} else {
		goto L34
	}
L34:
	;
	goto L32
L35:
	;
	goto L28
L36:
	;
	v176 = F_valkey_malloc(m, v105+int32(1))
	mBase = m.M
	v177 = m.ExcPending
	if v177 != 0 {
		goto L6
	} else {
		goto L37
	}
L37:
	;
	if v105 == int32(0) {
		v181 = v176
		goto L39
	} else {
		goto L40
	}
L38:
	;
	v183 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v181+v105))) = uint8(v183)
	if v105 == v183 {
		goto L42
	} else {
		goto L43
	}
L39:
	;
	goto L38
L40:
	;
	v180 = F__emscripten_memcpy_bulkmem(m, v176, v85, v105)
	mBase = m.M
	v181 = v180
	goto L39
L41:
	;
	v250 = v176
	v251 = v181
	goto L22
L42:
	;
	goto L41
L43:
	;
	v200 = int32(0)
	goto L44
L44:
	;
	goto L47
L45:
	;
	goto L42
L46:
	;
	v239 = v200 + int32(1)
	if v239 != v105 {
		v200 = v239
		goto L44
	} else {
		goto L53
	}
L47:
	;
	v207 = v181 + v200
	v208 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v207))))
	v217 = int32(0)
	goto L48
L48:
	;
	v221 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v217)+uint32(_c_F_genValkeyInfoStringCommandStats[0]))))
	if v208&int32(255) != v221 {
		goto L50
	} else {
		goto L51
	}
L49:
	;
	goto L46
L50:
	;
	v227 = v217 + int32(1)
	if v227 != int32(4) {
		v217 = v227
		goto L48
	} else {
		goto L52
	}
L51:
	;
	v224 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v217)+uint32(_c_F_genValkeyInfoStringCommandStats[1]))))
	*(*uint8)(unsafe.Add(mBase, uint32(v207))) = uint8(v224)
	goto L46
L52:
	;
	goto L49
L53:
	;
	goto L45
L54:
	;
	v264 = *(*int64)(unsafe.Add(mBase, uint32(v73)+120))
	v265 = *(*int64)(unsafe.Add(mBase, uint32(v73)+128))
	*(*int64)(unsafe.Add(mBase, uint32(v18+int32(40)))) = v265
	*(*int64)(unsafe.Add(mBase, uint32(v18+int32(32)))) = v264
	*(*float64)(unsafe.Add(mBase, uint32(v18+int32(24)))) = v263
	*(*int64)(unsafe.Add(mBase, uint32(v18+int32(16)))) = v252
	*(*int64)(unsafe.Add(mBase, uint32(v18)+8)) = v253
	*(*int32)(unsafe.Add(mBase, uint32(v18))) = v251
	v273 = F_sdscatprintf(m, v58, int32(_a_F_genValkeyInfoStringCommandStats_0), v18)
	mBase = m.M
	v274 = m.ExcPending
	if v274 != 0 {
		goto L6
	} else {
		goto L57
	}
L55:
	;
	v263 = base.F64_promote_f32(base.F32_div(base.F32_convert_i64_s(v252), base.F32_convert_i64_s(v253)))
	goto L54
L56:
	;
	v263 = float64(0)
	goto L54
L57:
	;
	if v250 == int32(0) {
		v279 = v273
		goto L11
	} else {
		goto L58
	}
L58:
	;
	F_valkey_free(m, v250)
	mBase = m.M
	v278 = m.ExcPending
	if v278 != 0 {
		goto L6
	} else {
		goto L59
	}
L59:
	;
	v279 = v273
	goto L11
L60:
	;
	v298 = F_hashtableNext(m, v18+int32(64), v18+int32(60))
	mBase = m.M
	v299 = m.ExcPending
	if v299 != 0 {
		goto L6
	} else {
		goto L63
	}
L61:
	;
	v291 = F_genValkeyInfoStringCommandStats(m, v279, v288)
	mBase = m.M
	v292 = m.ExcPending
	if v292 != 0 {
		goto L6
	} else {
		goto L62
	}
L62:
	;
	v293 = v291
	goto L60
L63:
	;
	if v298 != 0 {
		v58 = v293
		goto L9
	} else {
		goto L64
	}
L64:
	;
	goto L10
L65:
	;
	m.G0 = v18 + int32(112)
	return v300
}
func F_valkeyAeAddRead(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v2 != 0 {
		return
	} else {
		v3 = int32(1)
		*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v3
		v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		v9 = F_aeCreateFileEvent(m, v5, v6, v3, int32(1016), l0)
		mBase = m.M
		v10 = m.ExcPending
		if v10 != 0 {
			return
		} else {
			return
		}
	}
}
func F_valkeyAeAddWrite(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v2 != 0 {
		return
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = int32(1)
		v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		v9 = F_aeCreateFileEvent(m, v5, v6, int32(2), int32(1015), l0)
		mBase = m.M
		v10 = m.ExcPending
		if v10 != 0 {
			return
		} else {
			return
		}
	}
}
func F_valkeyAeDelWrite(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v2 == int32(0) {
		return
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = int32(0)
		v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		F_aeDeleteFileEvent(m, v7, v8, int32(2))
		mBase = m.M
		v11 = m.ExcPending
		if v11 != 0 {
			return
		} else {
			return
		}
	}
}
func F_valkeyAsyncAppendCmdLen(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v49 int32
	_ = v49
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v67 int32
	_ = v67
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v107 int32
	_ = v107
	var v114 int32
	_ = v114
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v121 int32
	_ = v121
	var v125 int32
	_ = v125
	var v129 int32
	_ = v129
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v143 int32
	_ = v143
	var v145 int32
	_ = v145
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v165 int32
	_ = v165
	var v172 int32
	_ = v172
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v179 int32
	_ = v179
	var v183 int32
	_ = v183
	var v187 int32
	_ = v187
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v201 int32
	_ = v201
	var v203 int32
	_ = v203
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v223 int32
	_ = v223
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v229 int32
	_ = v229
	var v234 int32
	_ = v234
	var v236 int32
	_ = v236
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
	var v248 int32
	_ = v248
	var v252 int32
	_ = v252
	var v256 int32
	_ = v256
	var v258 int32
	_ = v258
	var v259 int32
	_ = v259
	var v260 int32
	_ = v260
	var v262 int32
	_ = v262
	var v263 int32
	_ = v263
	var v264 int32
	_ = v264
	var v265 int32
	_ = v265
	var v270 int32
	_ = v270
	var v272 int32
	_ = v272
	var v275 int32
	_ = v275
	var v276 int32
	_ = v276
	var v280 int32
	_ = v280
	var v281 int32
	_ = v281
	var v282 int32
	_ = v282
	var v290 int32
	_ = v290
	var v291 int32
	_ = v291
	var v294 int32
	_ = v294
	var v300 int32
	_ = v300
	var v301 int32
	_ = v301
	var v308 int32
	_ = v308
	var v311 int32
	_ = v311
	var v323 int32
	_ = v323
	var v324 int32
	_ = v324
	var v327 int32
	_ = v327
	var v331 int32
	_ = v331
	var v332 int32
	_ = v332
	var v333 int32
	_ = v333
	var v334 int32
	_ = v334
	var v337 int32
	_ = v337
	var v340 int32
	_ = v340
	var v341 int32
	_ = v341
	var v342 int32
	_ = v342
	var v344 int32
	_ = v344
	var v345 int32
	_ = v345
	var v346 int32
	_ = v346
	var v349 int32
	_ = v349
	var v350 int32
	_ = v350
	var v352 int32
	_ = v352
	var v353 int32
	_ = v353
	var v354 int32
	_ = v354
	var v356 int32
	_ = v356
	var v357 int32
	_ = v357
	var v358 int32
	_ = v358
	var v359 int32
	_ = v359
	var v368 int32
	_ = v368
	var v369 int32
	_ = v369
	var v371 int32
	_ = v371
	var v377 int32
	_ = v377
	var v378 int32
	_ = v378
	var v401 int32
	_ = v401
	var v402 int32
	_ = v402
	var v403 int32
	_ = v403
	var v404 int32
	_ = v404
	var v407 int32
	_ = v407
	var v408 int32
	_ = v408
	var v409 int32
	_ = v409
	var v410 int32
	_ = v410
	var v416 int32
	_ = v416
	var v422 int32
	_ = v422
	var v423 int32
	_ = v423
	var v424 int32
	_ = v424
	var v425 int32
	_ = v425
	var v439 int32
	_ = v439
	var v443 int32
	_ = v443
	var v445 int32
	_ = v445
	var v459 int32
	_ = v459
	var v463 int32
	_ = v463
	var v465 int32
	_ = v465
	var v470 int32
	_ = v470
	var v472 int32
	_ = v472
	var v477 int32
	_ = v477
	var v479 int32
	_ = v479
	var v480 int32
	_ = v480
	var v481 int32
	_ = v481
	var v482 int32
	_ = v482
	var v484 int32
	_ = v484
	var v488 int32
	_ = v488
	var v492 int32
	_ = v492
	var v494 int32
	_ = v494
	var v495 int32
	_ = v495
	var v496 int32
	_ = v496
	var v498 int32
	_ = v498
	var v499 int32
	_ = v499
	var v500 int32
	_ = v500
	var v501 int32
	_ = v501
	var v506 int32
	_ = v506
	var v508 int32
	_ = v508
	var v511 int32
	_ = v511
	var v512 int32
	_ = v512
	var v516 int32
	_ = v516
	var v517 int32
	_ = v517
	var v518 int32
	_ = v518
	var v526 int32
	_ = v526
	var v534 int32
	_ = v534
	var v537 int32
	_ = v537
	var v539 int32
	_ = v539
	var v542 int32
	_ = v542
	var v548 int32
	_ = v548
	var v549 int32
	_ = v549
	var v555 int32
	_ = v555
	var v561 int32
	_ = v561
	var v571 int32
	_ = v571
	var v572 int32
	_ = v572
	var v573 int32
	_ = v573
	var v576 int32
	_ = v576
	var v577 int32
	_ = v577
	var v580 int32
	_ = v580
	var v581 int32
	_ = v581
	var v584 int32
	_ = v584
	var v588 int32
	_ = v588
	var v594 int32
	_ = v594
	var v600 int32
	_ = v600
	var v601 int32
	_ = v601
	var v604 int32
	_ = v604
	var v613 int32
	_ = v613
	var v621 int32
	_ = v621
	var v622 int32
	_ = v622
	var v625 int32
	_ = v625
	var v628 int32
	_ = v628
	var v629 int32
	_ = v629
	var v633 int32
	_ = v633
	var v637 int32
	_ = v637
	var v638 int32
	_ = v638
	var v639 int32
	_ = v639
	var v642 int32
	_ = v642
	var v644 int32
	_ = v644
	var v646 int64
	_ = v646
	var v647 int64
	_ = v647
	var v648 int64
	_ = v648
	var v649 int64
	_ = v649
	var v650 int64
	_ = v650
	var v651 int64
	_ = v651
	var v652 int64
	_ = v652
	var v654 int64
	_ = v654
	var v656 int64
	_ = v656
	var v658 int64
	_ = v658
	var v660 int64
	_ = v660
	var v662 int64
	_ = v662
	var v664 int32
	_ = v664
	var v665 int32
	_ = v665
	var v666 int32
	_ = v666
	var v669 int32
	_ = v669
	var v672 int32
	_ = v672
	var v673 int32
	_ = v673
	var v674 int32
	_ = v674
	var v676 int32
	_ = v676
	var v679 int32
	_ = v679
	var v680 int32
	_ = v680
	var v684 int32
	_ = v684
	var v688 int32
	_ = v688
	var v690 int32
	_ = v690
	var v697 int32
	_ = v697
	var v698 int32
	_ = v698
	var v699 int32
	_ = v699
	var v704 int32
	_ = v704
	var v709 int32
	_ = v709
	var v713 int32
	_ = v713
	var v717 int32
	_ = v717
	var v730 int32
	_ = v730
	var v732 int32
	_ = v732
	var v742 int32
	_ = v742
	var v743 int32
	_ = v743
	var v747 int32
	_ = v747
	var v755 int32
	_ = v755
	var v756 int32
	_ = v756
	var v759 int32
	_ = v759
	var v762 int32
	_ = v762
	var v763 int32
	_ = v763
	var v767 int32
	_ = v767
	var v771 int32
	_ = v771
	var v772 int32
	_ = v772
	var v773 int32
	_ = v773
	var v776 int32
	_ = v776
	var v778 int32
	_ = v778
	var v780 int64
	_ = v780
	var v781 int64
	_ = v781
	var v782 int64
	_ = v782
	var v783 int64
	_ = v783
	var v784 int64
	_ = v784
	var v785 int64
	_ = v785
	var v786 int64
	_ = v786
	var v788 int64
	_ = v788
	var v790 int64
	_ = v790
	var v792 int64
	_ = v792
	var v794 int64
	_ = v794
	var v796 int64
	_ = v796
	var v798 int32
	_ = v798
	var v799 int32
	_ = v799
	var v800 int32
	_ = v800
	var v803 int32
	_ = v803
	var v806 int32
	_ = v806
	var v807 int32
	_ = v807
	var v808 int32
	_ = v808
	var v810 int32
	_ = v810
	var v813 int32
	_ = v813
	var v814 int32
	_ = v814
	var v818 int32
	_ = v818
	var v822 int32
	_ = v822
	var v824 int32
	_ = v824
	var v831 int32
	_ = v831
	var v832 int32
	_ = v832
	var v833 int32
	_ = v833
	var v838 int32
	_ = v838
	var v843 int32
	_ = v843
	var v847 int32
	_ = v847
	var v851 int32
	_ = v851
	var v863 int32
	_ = v863
	var v871 int32
	_ = v871
	var v872 int32
	_ = v872
	var v875 int32
	_ = v875
	var v878 int32
	_ = v878
	var v879 int32
	_ = v879
	var v883 int32
	_ = v883
	var v887 int32
	_ = v887
	var v888 int32
	_ = v888
	var v889 int32
	_ = v889
	var v892 int32
	_ = v892
	var v894 int32
	_ = v894
	var v896 int64
	_ = v896
	var v897 int64
	_ = v897
	var v898 int64
	_ = v898
	var v899 int64
	_ = v899
	var v900 int64
	_ = v900
	var v901 int64
	_ = v901
	var v902 int64
	_ = v902
	var v904 int64
	_ = v904
	var v906 int64
	_ = v906
	var v908 int64
	_ = v908
	var v910 int64
	_ = v910
	var v912 int64
	_ = v912
	var v914 int32
	_ = v914
	var v915 int32
	_ = v915
	var v916 int32
	_ = v916
	var v919 int32
	_ = v919
	var v922 int32
	_ = v922
	var v923 int32
	_ = v923
	var v924 int32
	_ = v924
	var v926 int32
	_ = v926
	var v929 int32
	_ = v929
	var v930 int32
	_ = v930
	var v934 int32
	_ = v934
	var v938 int32
	_ = v938
	var v940 int32
	_ = v940
	var v947 int32
	_ = v947
	var v948 int32
	_ = v948
	var v949 int32
	_ = v949
	var v954 int32
	_ = v954
	var v959 int32
	_ = v959
	var v963 int32
	_ = v963
	var v967 int32
	_ = v967
	var v972 int32
	_ = v972
	var v976 int32
	_ = v976
	var v977 int32
	_ = v977
	var v979 int32
	_ = v979
	var v984 int32
	_ = v984
	var v986 int32
	_ = v986
	var v987 int32
	_ = v987
	var v988 int32
	_ = v988
	var v989 int32
	_ = v989
	var v991 int32
	_ = v991
	var v995 int32
	_ = v995
	var v999 int32
	_ = v999
	var v1001 int32
	_ = v1001
	var v1002 int32
	_ = v1002
	var v1003 int32
	_ = v1003
	var v1005 int32
	_ = v1005
	var v1006 int32
	_ = v1006
	var v1007 int32
	_ = v1007
	var v1008 int32
	_ = v1008
	var v1013 int32
	_ = v1013
	var v1015 int32
	_ = v1015
	var v1018 int32
	_ = v1018
	var v1019 int32
	_ = v1019
	var v1023 int32
	_ = v1023
	var v1024 int32
	_ = v1024
	var v1025 int32
	_ = v1025
	var v1037 int32
	_ = v1037
	var v1038 int32
	_ = v1038
	var v1039 int32
	_ = v1039
	var v1040 int32
	_ = v1040
	var v1043 int32
	_ = v1043
	var v1051 int32
	_ = v1051
	var v1053 int32
	_ = v1053
	var v1059 int32
	_ = v1059
	var v1060 int32
	_ = v1060
	var v1061 int32
	_ = v1061
	var v1062 int32
	_ = v1062
	var v1069 int32
	_ = v1069
	var v1077 int32
	_ = v1077
	var v1079 int32
	_ = v1079
	var v1086 int32
	_ = v1086
	var v1094 int32
	_ = v1094
	var v1096 int32
	_ = v1096
	var v1120 int32
	_ = v1120
	var v1145 int32
	_ = v1145
	var v1165 int32
	_ = v1165
	var v1187 int32
	_ = v1187
	var v1188 int32
	_ = v1188
	var v1189 int32
	_ = v1189
	var v1190 int32
	_ = v1190
	var v1197 int32
	_ = v1197
	var v1200 int64
	_ = v1200
	var v1203 int32
	_ = v1203
	var v1206 int32
	_ = v1206
	var v1209 int64
	_ = v1209
	var v1212 int32
	_ = v1212
	var v1215 int32
	_ = v1215
	var v1216 int32
	_ = v1216
	var v1217 int32
	_ = v1217
	var v1221 int64
	_ = v1221
	var v1223 int64
	_ = v1223
	var v1226 int32
	_ = v1226
	var v1229 int32
	_ = v1229
	var v1230 int32
	_ = v1230
	var v1233 int32
	_ = v1233
	var v1235 int32
	_ = v1235
	var v1263 int32
	_ = v1263
	var v1274 int32
	_ = v1274
	var v1281 int32
	_ = v1281
	var v1285 int32
	_ = v1285
	var v1286 int32
	_ = v1286
	var v1287 int32
	_ = v1287
	var v1289 int32
	_ = v1289
	var v1290 int32
	_ = v1290
	var v1292 int32
	_ = v1292
	var v1299 int32
	_ = v1299
	v19 = m.G0
	v21 = v19 - int32(64)
	m.G0 = v21
	v23 = int32(-1)
	v24 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+140)))
	if v24&int32(12) != 0 {
		v1299 = v23
		goto L1
	} else {
		goto L2
	}
L1:
	;
	m.G0 = v21 + int32(64)
	return v1299
L2:
	;
	v31 = F_nextArgument(m, l3, l4, v21+int32(28), v21+int32(24))
	mBase = m.M
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v21)+28))
	if v32 == int32(0) {
		v1299 = v23
		goto L1
	} else {
		goto L3
	}
L3:
	;
	if v31 != 0 {
		goto L5
	} else {
		goto L6
	}
L4:
	;
	v40 = m.G3
	v41 = int32(1)
	v42 = int32(*(*int8)(unsafe.Add(mBase, uint32(v32))))
	if base.Ui32(v42+int32(-65)) < base.Ui32(int32(26)) {
		goto L8
	} else {
		goto L9
	}
L5:
	;
	v36 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v31))))
	v39 = base.B2i32(v36 == int32(36))
	goto L4
L6:
	;
	v39 = int32(0)
	goto L4
L7:
	;
	goto L13
L8:
	;
	v49 = v42 | int32(32)
	goto L10
L9:
	;
	v49 = v42
	goto L10
L10:
	;
	goto L7
L11:
	;
	v226 = v223 | base.B2i32(v49 == int32(112))
	v227 = v32 + v226
	*(*int32)(unsafe.Add(mBase, uint32(v21)+28)) = v227
	v229 = *(*int32)(unsafe.Add(mBase, uint32(v21)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v21)+24)) = v229 - v226
	if v39 == int32(0) {
		goto L65
	} else {
		goto L66
	}
L12:
	;
	if v95-v97 == int32(0) {
		v223 = v41
		goto L11
	} else {
		goto L27
	}
L13:
	;
	v56 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v40)+uint32(_c_F_valkeyAsyncAppendCmdLen[0]))))
	if v56 != 0 {
		goto L16
	} else {
		goto L17
	}
L15:
	;
	v95 = F_tolower(m, v90)
	mBase = m.M
	v96 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v91))))
	v97 = F_tolower(m, v96)
	mBase = m.M
	goto L12
L16:
	;
	v58 = v40 + int32(_a_F_valkeyAsyncAppendCmdLen_0)
	v59 = v32
	v60 = int32(2)
	v61 = v56
	goto L19
L17:
	;
	v90 = int32(0)
	v91 = v32
	goto L15
L18:
	;
	v90 = v87 & int32(255)
	v91 = v85
	goto L15
L19:
	;
	v63 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v59))))
	if v63 == int32(0) {
		v85 = v59
		v87 = v61
		goto L18
	} else {
		goto L21
	}
L20:
	;
	v85 = v79
	v87 = int32(0)
	goto L18
L21:
	;
	v67 = v60 + int32(-1)
	if v67 == int32(0) {
		v85 = v59
		v87 = v61
		goto L18
	} else {
		goto L22
	}
L22:
	;
	v71 = v61 & int32(255)
	if v71 == v63 {
		goto L23
	} else {
		goto L24
	}
L23:
	;
	v78 = int32(1)
	v79 = v59 + v78
	v80 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v58)+1)))
	if v80 != 0 {
		v58 = v58 + v78
		v59 = v79
		v60 = v67
		v61 = v80
		goto L19
	} else {
		goto L26
	}
L24:
	;
	v73 = F_tolower(m, v71)
	mBase = m.M
	v74 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v59))))
	v75 = F_tolower(m, v74)
	mBase = m.M
	if v73 == v75 {
		goto L23
	} else {
		goto L25
	}
L25:
	;
	v77 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v58))))
	v85 = v59
	v87 = v77
	goto L18
L26:
	;
	goto L20
L27:
	;
	v107 = m.G3
	goto L29
L28:
	;
	if v153-v155 == int32(0) {
		v223 = v41
		goto L11
	} else {
		goto L43
	}
L29:
	;
	v114 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v107)+uint32(_c_F_valkeyAsyncAppendCmdLen[1]))))
	if v114 != 0 {
		goto L32
	} else {
		goto L33
	}
L31:
	;
	v153 = F_tolower(m, v148)
	mBase = m.M
	v154 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v149))))
	v155 = F_tolower(m, v154)
	mBase = m.M
	goto L28
L32:
	;
	v116 = v107 + int32(_a_F_valkeyAsyncAppendCmdLen_1)
	v117 = v32
	v118 = int32(2)
	v119 = v114
	goto L35
L33:
	;
	v148 = int32(0)
	v149 = v32
	goto L31
L34:
	;
	v148 = v145 & int32(255)
	v149 = v143
	goto L31
L35:
	;
	v121 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v117))))
	if v121 == int32(0) {
		v143 = v117
		v145 = v119
		goto L34
	} else {
		goto L37
	}
L36:
	;
	v143 = v137
	v145 = int32(0)
	goto L34
L37:
	;
	v125 = v118 + int32(-1)
	if v125 == int32(0) {
		v143 = v117
		v145 = v119
		goto L34
	} else {
		goto L38
	}
L38:
	;
	v129 = v119 & int32(255)
	if v129 == v121 {
		goto L39
	} else {
		goto L40
	}
L39:
	;
	v136 = int32(1)
	v137 = v117 + v136
	v138 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v116)+1)))
	if v138 != 0 {
		v116 = v116 + v136
		v117 = v137
		v118 = v125
		v119 = v138
		goto L35
	} else {
		goto L42
	}
L40:
	;
	v131 = F_tolower(m, v129)
	mBase = m.M
	v132 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v117))))
	v133 = F_tolower(m, v132)
	mBase = m.M
	if v131 == v133 {
		goto L39
	} else {
		goto L41
	}
L41:
	;
	v135 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v116))))
	v143 = v117
	v145 = v135
	goto L34
L42:
	;
	goto L36
L43:
	;
	v165 = m.G3
	goto L45
L44:
	;
	v223 = base.B2i32(v211-v213 == int32(0))
	goto L11
L45:
	;
	v172 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v165)+uint32(_c_F_valkeyAsyncAppendCmdLen[2]))))
	if v172 != 0 {
		goto L48
	} else {
		goto L49
	}
L47:
	;
	v211 = F_tolower(m, v206)
	mBase = m.M
	v212 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v207))))
	v213 = F_tolower(m, v212)
	mBase = m.M
	goto L44
L48:
	;
	v174 = v165 + int32(_a_F_valkeyAsyncAppendCmdLen_2)
	v175 = v32
	v176 = int32(3)
	v177 = v172
	goto L51
L49:
	;
	v206 = int32(0)
	v207 = v32
	goto L47
L50:
	;
	v206 = v203 & int32(255)
	v207 = v201
	goto L47
L51:
	;
	v179 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v175))))
	if v179 == int32(0) {
		v201 = v175
		v203 = v177
		goto L50
	} else {
		goto L53
	}
L52:
	;
	v201 = v195
	v203 = int32(0)
	goto L50
L53:
	;
	v183 = v176 + int32(-1)
	if v183 == int32(0) {
		v201 = v175
		v203 = v177
		goto L50
	} else {
		goto L54
	}
L54:
	;
	v187 = v177 & int32(255)
	if v187 == v179 {
		goto L55
	} else {
		goto L56
	}
L55:
	;
	v194 = int32(1)
	v195 = v175 + v194
	v196 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v174)+1)))
	if v196 != 0 {
		v174 = v174 + v194
		v175 = v195
		v176 = v183
		v177 = v196
		goto L51
	} else {
		goto L58
	}
L56:
	;
	v189 = F_tolower(m, v187)
	mBase = m.M
	v190 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v175))))
	v191 = F_tolower(m, v190)
	mBase = m.M
	if v189 == v191 {
		goto L55
	} else {
		goto L57
	}
L57:
	;
	v193 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v174))))
	v201 = v175
	v203 = v193
	goto L50
L58:
	;
	goto L52
L59:
	;
	v1274 = m.G3
	F_valkeySetError(m, l0, int32(5), v1274+int32(_a_F_valkeyAsyncAppendCmdLen_3))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(l0)+208)) = l0 + int32(8)
	v1281 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+204)) = v1281
	if v1263 == int32(0) {
		v1299 = v23
		goto L1
	} else {
		goto L314
	}
L60:
	;
	v1263 = int32(0)
	goto L59
L61:
	;
	v1187 = F_valkeyAppendCmdLen(m, l0, l3, l4)
	mBase = m.M
	v1188 = m.ExcPending
	if v1188 != 0 {
		goto L93
	} else {
		goto L298
	}
L62:
	;
	v1165 = *(*int32)(unsafe.Add(mBase, uint32(l0)+292))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+292)) = v1165 + int32(1)
	goto L61
L63:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+208)) = l0 + int32(8)
	v1145 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+204)) = v1145
	v1299 = v23
	goto L1
L64:
	;
	v1120 = m.G3
	F_valkeySetError(m, l0, int32(5), v1120+int32(_a_F_valkeyAsyncAppendCmdLen_3))
	mBase = m.M
	goto L63
L65:
	;
	v470 = m.G3
	v472 = v470 + int32(_a_F_valkeyAsyncAppendCmdLen_4)
	goto L134
L66:
	;
	v234 = m.G3
	v236 = v234 + int32(_a_F_valkeyAsyncAppendCmdLen_5)
	goto L68
L67:
	;
	if v280-v282 != 0 {
		goto L65
	} else {
		goto L82
	}
L68:
	;
	v241 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v227))))
	if v241 != 0 {
		goto L71
	} else {
		goto L72
	}
L70:
	;
	v280 = F_tolower(m, v275)
	mBase = m.M
	v281 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v276))))
	v282 = F_tolower(m, v281)
	mBase = m.M
	goto L67
L71:
	;
	v243 = v227
	v244 = v236
	v245 = int32(11)
	v246 = v241
	goto L74
L72:
	;
	v275 = int32(0)
	v276 = v236
	goto L70
L73:
	;
	v275 = v272 & int32(255)
	v276 = v270
	goto L70
L74:
	;
	v248 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v244))))
	if v248 == int32(0) {
		v270 = v244
		v272 = v246
		goto L73
	} else {
		goto L76
	}
L75:
	;
	v270 = v264
	v272 = int32(0)
	goto L73
L76:
	;
	v252 = v245 + int32(-1)
	if v252 == int32(0) {
		v270 = v244
		v272 = v246
		goto L73
	} else {
		goto L77
	}
L77:
	;
	v256 = v246 & int32(255)
	if v256 == v248 {
		goto L78
	} else {
		goto L79
	}
L78:
	;
	v263 = int32(1)
	v264 = v244 + v263
	v265 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v243)+1)))
	if v265 != 0 {
		v243 = v243 + v263
		v244 = v264
		v245 = v252
		v246 = v265
		goto L74
	} else {
		goto L81
	}
L79:
	;
	v258 = F_tolower(m, v256)
	mBase = m.M
	v259 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v244))))
	v260 = F_tolower(m, v259)
	mBase = m.M
	if v258 == v260 {
		goto L78
	} else {
		goto L80
	}
L80:
	;
	v262 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v243))))
	v270 = v244
	v272 = v262
	goto L73
L81:
	;
	goto L75
L82:
	;
	v290 = *(*int32)(unsafe.Add(mBase, uint32(l0)+140))
	v291 = int32(32)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+140)) = v290 | v291
	v294 = l4 + l3
	v300 = F_nextArgument(m, v31, v294-v31, v21+v291, v21+int32(20))
	mBase = m.M
	v301 = *(*int32)(unsafe.Add(mBase, uint32(v21)+32))
	if v300|v301 == int32(0) {
		goto L83
	} else {
		goto L84
	}
L83:
	;
	if v223 == int32(0) {
		goto L61
	} else {
		goto L111
	}
L84:
	;
	if v223 != 0 {
		goto L85
	} else {
		goto L86
	}
L85:
	;
	v308 = int32(288)
	goto L87
L86:
	;
	v308 = int32(280)
	goto L87
L87:
	;
	if v49 == int32(112) {
		goto L88
	} else {
		goto L89
	}
L88:
	;
	v311 = int32(284)
	goto L90
L89:
	;
	v311 = v308
	goto L90
L90:
	;
	v323 = v301
	v324 = v300
	v327 = l1
	v331 = int32(1)
	v332 = int32(0)
	goto L91
L91:
	;
	v333 = *(*int32)(unsafe.Add(mBase, uint32(v21)+20))
	v334 = F_sdsnewlen(m, v323, v333)
	mBase = m.M
	v337 = m.ExcPending
	if v337 != 0 {
		goto L93
	} else {
		goto L94
	}
L92:
	;
	goto L83
L93:
	;
	return int32(0)
L94:
	;
	if v334 == int32(0) {
		goto L64
	} else {
		goto L95
	}
L95:
	;
	v340 = *(*int32)(unsafe.Add(mBase, uint32(l0+v311)))
	v341 = F_dictFind(m, v340, v334)
	mBase = m.M
	v342 = m.ExcPending
	if v342 != 0 {
		goto L93
	} else {
		goto L98
	}
L96:
	;
	v356 = m.G4
	v357 = *(*int32)(unsafe.Add(mBase, uint32(v356)))
	v358 = m.T0[v357].(func(*base.Module, int32) int32)(m, int32(24))
	mBase = m.M
	v359 = m.ExcPending
	if v359 != 0 {
		goto L93
	} else {
		goto L104
	}
L97:
	;
	v345 = *(*int32)(unsafe.Add(mBase, uint32(v341)+8))
	goto L103
L98:
	;
	if v341 != 0 {
		goto L97
	} else {
		goto L99
	}
L99:
	;
	if v223 != 0 {
		goto L100
	} else {
		goto L101
	}
L100:
	;
	v344 = int32(0)
	goto L102
L101:
	;
	v344 = v327
	goto L102
L102:
	;
	v352 = v344
	v353 = v331
	v354 = v332
	goto L96
L103:
	;
	v346 = *(*int32)(unsafe.Add(mBase, uint32(v345)+8))
	v349 = *(*int32)(unsafe.Add(mBase, uint32(v345)+4))
	v350 = *(*int32)(unsafe.Add(mBase, uint32(v345)+20))
	v352 = v349
	v353 = v346 + int32(1)
	v354 = v350
	goto L96
L104:
	;
	if v358 == int32(0) {
		goto L60
	} else {
		goto L105
	}
L105:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v358)+20)) = v354
	*(*int32)(unsafe.Add(mBase, uint32(v358)+16)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(v358)+12)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v358)+8)) = v353
	*(*int32)(unsafe.Add(mBase, uint32(v358)+4)) = v352
	v368 = F_dictReplace(m, v340, v334, v358)
	mBase = m.M
	v369 = m.ExcPending
	if v369 != 0 {
		goto L93
	} else {
		goto L107
	}
L106:
	;
	v377 = F_nextArgument(m, v324, v294-v324, v21+int32(32), v21+int32(20))
	mBase = m.M
	v378 = *(*int32)(unsafe.Add(mBase, uint32(v21)+32))
	if v377|v378 != 0 {
		v323 = v378
		v324 = v377
		v327 = v352
		v331 = v353
		v332 = v354
		goto L91
	} else {
		goto L110
	}
L107:
	;
	if v368 != 0 {
		goto L106
	} else {
		goto L108
	}
L108:
	;
	F_sdsfree(m, v334)
	mBase = m.M
	v371 = m.ExcPending
	if v371 != 0 {
		goto L93
	} else {
		goto L109
	}
L109:
	;
	goto L106
L110:
	;
	goto L92
L111:
	;
	v401 = m.G4
	v402 = *(*int32)(unsafe.Add(mBase, uint32(v401)))
	v403 = m.T0[v402].(func(*base.Module, int32) int32)(m, int32(16))
	mBase = m.M
	v404 = m.ExcPending
	if v404 != 0 {
		goto L93
	} else {
		goto L112
	}
L112:
	;
	if v403 == int32(0) {
		goto L60
	} else {
		goto L113
	}
L113:
	;
	v407 = m.G4
	v408 = *(*int32)(unsafe.Add(mBase, uint32(v407)))
	v409 = m.T0[v408].(func(*base.Module, int32) int32)(m, l4)
	mBase = m.M
	v410 = m.ExcPending
	if v410 != 0 {
		goto L93
	} else {
		goto L114
	}
L114:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v403))) = v409
	if v409 == int32(0) {
		v1263 = v403
		goto L59
	} else {
		goto L115
	}
L115:
	;
	if l4 == int32(0) {
		goto L117
	} else {
		goto L118
	}
L116:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v403)+12)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(v403)+8)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v403)+4)) = l4
	v422 = m.G4
	v423 = *(*int32)(unsafe.Add(mBase, uint32(v422)))
	v424 = m.T0[v423].(func(*base.Module, int32) int32)(m, int32(24))
	mBase = m.M
	v425 = m.ExcPending
	if v425 != 0 {
		goto L93
	} else {
		goto L119
	}
L117:
	;
	goto L116
L118:
	;
	v416 = F__emscripten_memcpy_bulkmem(m, v409, l3, l4)
	mBase = m.M
	goto L117
L119:
	;
	if v290&int32(32) == int32(0) {
		goto L120
	} else {
		goto L121
	}
L120:
	;
	if v424 == int32(0) {
		v1263 = v403
		goto L59
	} else {
		goto L127
	}
L121:
	;
	if v424 == int32(0) {
		v1263 = v403
		goto L59
	} else {
		goto L122
	}
L122:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v424)+20)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v424)+16)) = v403
	*(*int64)(unsafe.Add(mBase, uint32(v424)+8)) = int64(1)
	*(*int32)(unsafe.Add(mBase, uint32(v424))) = int32(0)
	v439 = m.G5
	*(*int32)(unsafe.Add(mBase, uint32(v424)+4)) = v439 + int32(1136)
	v443 = *(*int32)(unsafe.Add(mBase, uint32(l0)+272))
	if v443 != 0 {
		goto L123
	} else {
		goto L124
	}
L123:
	;
	v445 = *(*int32)(unsafe.Add(mBase, uint32(l0)+276))
	if v445 == int32(0) {
		goto L125
	} else {
		goto L126
	}
L124:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+272)) = v424
	goto L123
L125:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+276)) = v424
	goto L61
L126:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v445))) = v424
	goto L125
L127:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v424)+20)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v424)+16)) = v403
	*(*int64)(unsafe.Add(mBase, uint32(v424)+8)) = int64(1)
	*(*int32)(unsafe.Add(mBase, uint32(v424))) = int32(0)
	v459 = m.G5
	*(*int32)(unsafe.Add(mBase, uint32(v424)+4)) = v459 + int32(1136)
	v463 = *(*int32)(unsafe.Add(mBase, uint32(l0)+256))
	if v463 != 0 {
		goto L128
	} else {
		goto L129
	}
L128:
	;
	v465 = *(*int32)(unsafe.Add(mBase, uint32(l0)+260))
	if v465 == int32(0) {
		goto L130
	} else {
		goto L131
	}
L129:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+256)) = v424
	goto L128
L130:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+260)) = v424
	goto L61
L131:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v465))) = v424
	goto L130
L132:
	;
	v976 = m.G3
	v977 = *(*int32)(unsafe.Add(mBase, uint32(l0)+140))
	v979 = v976 + int32(_a_F_valkeyAsyncAppendCmdLen_6)
	goto L264
L133:
	;
	if v516-v518 != 0 {
		goto L132
	} else {
		goto L148
	}
L134:
	;
	v477 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v227))))
	if v477 != 0 {
		goto L137
	} else {
		goto L138
	}
L136:
	;
	v516 = F_tolower(m, v511)
	mBase = m.M
	v517 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v512))))
	v518 = F_tolower(m, v517)
	mBase = m.M
	goto L133
L137:
	;
	v479 = v227
	v480 = v472
	v481 = int32(13)
	v482 = v477
	goto L140
L138:
	;
	v511 = int32(0)
	v512 = v472
	goto L136
L139:
	;
	v511 = v508 & int32(255)
	v512 = v506
	goto L136
L140:
	;
	v484 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v480))))
	if v484 == int32(0) {
		v506 = v480
		v508 = v482
		goto L139
	} else {
		goto L142
	}
L141:
	;
	v506 = v500
	v508 = int32(0)
	goto L139
L142:
	;
	v488 = v481 + int32(-1)
	if v488 == int32(0) {
		v506 = v480
		v508 = v482
		goto L139
	} else {
		goto L143
	}
L143:
	;
	v492 = v482 & int32(255)
	if v492 == v484 {
		goto L144
	} else {
		goto L145
	}
L144:
	;
	v499 = int32(1)
	v500 = v480 + v499
	v501 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v479)+1)))
	if v501 != 0 {
		v479 = v479 + v499
		v480 = v500
		v481 = v488
		v482 = v501
		goto L140
	} else {
		goto L147
	}
L145:
	;
	v494 = F_tolower(m, v492)
	mBase = m.M
	v495 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v480))))
	v496 = F_tolower(m, v495)
	mBase = m.M
	if v494 == v496 {
		goto L144
	} else {
		goto L146
	}
L146:
	;
	v498 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v479))))
	v506 = v480
	v508 = v498
	goto L139
L147:
	;
	goto L141
L148:
	;
	v526 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+140)))
	if v526&int32(32) == int32(0) {
		v1299 = v23
		goto L1
	} else {
		goto L149
	}
L149:
	;
	if v223 != 0 {
		goto L150
	} else {
		goto L151
	}
L150:
	;
	v534 = int32(288)
	goto L152
L151:
	;
	v534 = int32(280)
	goto L152
L152:
	;
	if v49 == int32(112) {
		goto L153
	} else {
		goto L154
	}
L153:
	;
	v537 = int32(284)
	goto L155
L154:
	;
	v537 = v534
	goto L155
L155:
	;
	v539 = *(*int32)(unsafe.Add(mBase, uint32(l0+v537)))
	if v39 == int32(0) {
		goto L157
	} else {
		goto L158
	}
L156:
	;
	v972 = m.G3
	F_valkeySetError(m, l0, int32(5), v972+int32(_a_F_valkeyAsyncAppendCmdLen_3))
	mBase = m.M
	if l0 != 0 {
		goto L63
	} else {
		goto L261
	}
L157:
	;
	v604 = v21 + int32(32)
	*(*int32)(unsafe.Add(mBase, uint32(v604))) = v539
	*(*int32)(unsafe.Add(mBase, uint32(v604)+20)) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v604)+12)) = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v604)+4)) = int64(4294967295)
	goto L173
L158:
	;
	v542 = l4 + l3
	v548 = F_nextArgument(m, v31, v542-v31, v21+int32(32), v21+int32(20))
	mBase = m.M
	v549 = *(*int32)(unsafe.Add(mBase, uint32(v21)+32))
	if v548|v549 == int32(0) {
		goto L61
	} else {
		goto L159
	}
L159:
	;
	v555 = v549
	v561 = v548
	goto L160
L160:
	;
	v571 = *(*int32)(unsafe.Add(mBase, uint32(v21)+20))
	v572 = F_sdsnewlen(m, v555, v571)
	mBase = m.M
	v573 = m.ExcPending
	if v573 != 0 {
		goto L93
	} else {
		goto L162
	}
L162:
	;
	if v572 == int32(0) {
		goto L156
	} else {
		goto L163
	}
L163:
	;
	v576 = F_dictFind(m, v539, v572)
	mBase = m.M
	v577 = m.ExcPending
	if v577 != 0 {
		goto L93
	} else {
		goto L166
	}
L164:
	;
	F_sdsfree(m, v572)
	mBase = m.M
	v594 = m.ExcPending
	if v594 != 0 {
		goto L93
	} else {
		goto L171
	}
L165:
	;
	v588 = *(*int32)(unsafe.Add(mBase, uint32(l0)+292))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+292)) = v588 + int32(1)
	goto L164
L166:
	;
	if v576 == int32(0) {
		goto L165
	} else {
		goto L167
	}
L167:
	;
	v580 = *(*int32)(unsafe.Add(mBase, uint32(v576)+8))
	goto L169
L168:
	;
	v584 = *(*int32)(unsafe.Add(mBase, uint32(l0)+292))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+292)) = v584 + int32(1)
	goto L164
L169:
	;
	v581 = *(*int32)(unsafe.Add(mBase, uint32(v580)+12))
	if v581 != 0 {
		goto L168
	} else {
		goto L170
	}
L170:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v580)+12)) = int32(1)
	goto L164
L171:
	;
	v600 = F_nextArgument(m, v561, v542-v561, v21+int32(32), v21+int32(20))
	mBase = m.M
	v601 = *(*int32)(unsafe.Add(mBase, uint32(v21)+32))
	if v600|v601 != 0 {
		v555 = v601
		v561 = v600
		goto L160
	} else {
		goto L172
	}
L172:
	;
	goto L61
L173:
	;
	v613 = v21 + int32(32)
	v621 = v21 + int32(52)
	v622 = *(*int32)(unsafe.Add(mBase, uint32(v613)+16))
	if v622 != 0 {
		goto L176
	} else {
		goto L177
	}
L174:
	;
	if v717 == int32(0) {
		goto L62
	} else {
		goto L200
	}
L175:
	;
	v628 = v621
	v629 = v625
	goto L178
L176:
	;
	v625 = int32(1)
	goto L175
L177:
	;
	v625 = int32(0)
	goto L175
L178:
	;
	switch v629 {
	case 0:
		goto L183
	default:
		goto L182
	}
L180:
	;
	v629 = int32(0)
	goto L178
L181:
	;
	goto L174
L182:
	;
	v709 = *(*int32)(unsafe.Add(mBase, uint32(v628)))
	*(*int32)(unsafe.Add(mBase, uint32(v613)+16)) = v709
	if v709 == int32(0) {
		goto L180
	} else {
		goto L199
	}
L183:
	;
	v633 = *(*int32)(unsafe.Add(mBase, uint32(v613)+4))
	if v633 != int32(-1) {
		v672 = v633
		goto L184
	} else {
		goto L185
	}
L184:
	;
	v673 = int32(1)
	v674 = v672 + v673
	*(*int32)(unsafe.Add(mBase, uint32(v613)+4)) = v674
	v676 = int32(0)
	v679 = *(*int32)(unsafe.Add(mBase, uint32(v613)))
	v680 = *(*int32)(unsafe.Add(mBase, uint32(v613)+8))
	v684 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v679+v680+int32(26)))))
	if v684 == int32(255) {
		goto L193
	} else {
		goto L194
	}
L185:
	;
	v637 = *(*int32)(unsafe.Add(mBase, uint32(v613)+8))
	if v637 != 0 {
		v672 = int32(-1)
		goto L184
	} else {
		goto L186
	}
L186:
	;
	v638 = *(*int32)(unsafe.Add(mBase, uint32(v613)))
	v639 = *(*int32)(unsafe.Add(mBase, uint32(v613)+12))
	if v639 == int32(0) {
		goto L188
	} else {
		goto L189
	}
L187:
	;
	v666 = *(*int32)(unsafe.Add(mBase, uint32(v665)+20))
	if v666 != int32(-1) {
		goto L190
	} else {
		goto L191
	}
L188:
	;
	v646 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v638)+16)))
	v647 = int64(*(*int8)(unsafe.Add(mBase, uint32(v638)+27)))
	v648 = int64(*(*int32)(unsafe.Add(mBase, uint32(v638)+8)))
	v649 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v638)+12)))
	v650 = int64(*(*int8)(unsafe.Add(mBase, uint32(v638)+26)))
	v651 = int64(*(*int32)(unsafe.Add(mBase, uint32(v638)+4)))
	v652 = F_wangHash64(m, v651)
	mBase = m.M
	v654 = F_wangHash64(m, v650+v652)
	mBase = m.M
	v656 = F_wangHash64(m, v649+v654)
	mBase = m.M
	v658 = F_wangHash64(m, v648+v656)
	mBase = m.M
	v660 = F_wangHash64(m, v647+v658)
	mBase = m.M
	v662 = F_wangHash64(m, v646+v660)
	mBase = m.M
	*(*int64)(unsafe.Add(mBase, uint32(v613)+24)) = v662
	v664 = *(*int32)(unsafe.Add(mBase, uint32(v613)))
	v665 = v664
	goto L187
L189:
	;
	v642 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v638)+24)))
	v644 = v642 + int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v638)+24)) = uint16(v644)
	v665 = v638
	goto L187
L190:
	;
	v672 = v666 + int32(-1)
	goto L184
L191:
	;
	v669 = *(*int32)(unsafe.Add(mBase, uint32(v613)+4))
	v672 = v669
	goto L184
L192:
	;
	v699 = int32(2)
	v704 = *(*int32)(unsafe.Add(mBase, uint32(v679+v697<<(uint(v699)%32)+int32(4))))
	v628 = v704 + v698<<(uint(v699)%32)
	v629 = int32(1)
	goto L178
L193:
	;
	v688 = v676
	goto L195
L194:
	;
	v688 = v673 << (uint(v684) % 32)
	goto L195
L195:
	;
	if v674 < v688 {
		v697 = v680
		v698 = v674
		goto L192
	} else {
		goto L196
	}
L196:
	;
	if v680 != 0 {
		v717 = v676
		goto L181
	} else {
		goto L197
	}
L197:
	;
	v690 = *(*int32)(unsafe.Add(mBase, uint32(v679)+20))
	if v690 == int32(-1) {
		v717 = v676
		goto L181
	} else {
		goto L198
	}
L198:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v613)+4)) = int64(4294967296)
	v697 = int32(1)
	v698 = int32(0)
	goto L192
L199:
	;
	v713 = *(*int32)(unsafe.Add(mBase, uint32(v709)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v621))) = v713
	v717 = v709
	goto L181
L200:
	;
	v730 = v717
	v732 = int32(1)
	goto L201
L201:
	;
	v742 = *(*int32)(unsafe.Add(mBase, uint32(v730)+8))
	goto L204
L203:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v742)+12)) = int32(1)
	v863 = v21 + int32(32)
	v871 = v21 + int32(52)
	v872 = *(*int32)(unsafe.Add(mBase, uint32(v863)+16))
	if v872 != 0 {
		goto L236
	} else {
		goto L237
	}
L204:
	;
	v743 = *(*int32)(unsafe.Add(mBase, uint32(v742)+12))
	if v743 == int32(0) {
		goto L203
	} else {
		goto L205
	}
L205:
	;
	v747 = v21 + int32(32)
	v755 = v21 + int32(52)
	v756 = *(*int32)(unsafe.Add(mBase, uint32(v747)+16))
	if v756 != 0 {
		goto L208
	} else {
		goto L209
	}
L206:
	;
	if v851 != 0 {
		v730 = v851
		goto L201
	} else {
		goto L232
	}
L207:
	;
	v762 = v755
	v763 = v759
	goto L210
L208:
	;
	v759 = int32(1)
	goto L207
L209:
	;
	v759 = int32(0)
	goto L207
L210:
	;
	switch v763 {
	case 0:
		goto L215
	default:
		goto L214
	}
L212:
	;
	v763 = int32(0)
	goto L210
L213:
	;
	goto L206
L214:
	;
	v843 = *(*int32)(unsafe.Add(mBase, uint32(v762)))
	*(*int32)(unsafe.Add(mBase, uint32(v747)+16)) = v843
	if v843 == int32(0) {
		goto L212
	} else {
		goto L231
	}
L215:
	;
	v767 = *(*int32)(unsafe.Add(mBase, uint32(v747)+4))
	if v767 != int32(-1) {
		v806 = v767
		goto L216
	} else {
		goto L217
	}
L216:
	;
	v807 = int32(1)
	v808 = v806 + v807
	*(*int32)(unsafe.Add(mBase, uint32(v747)+4)) = v808
	v810 = int32(0)
	v813 = *(*int32)(unsafe.Add(mBase, uint32(v747)))
	v814 = *(*int32)(unsafe.Add(mBase, uint32(v747)+8))
	v818 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v813+v814+int32(26)))))
	if v818 == int32(255) {
		goto L225
	} else {
		goto L226
	}
L217:
	;
	v771 = *(*int32)(unsafe.Add(mBase, uint32(v747)+8))
	if v771 != 0 {
		v806 = int32(-1)
		goto L216
	} else {
		goto L218
	}
L218:
	;
	v772 = *(*int32)(unsafe.Add(mBase, uint32(v747)))
	v773 = *(*int32)(unsafe.Add(mBase, uint32(v747)+12))
	if v773 == int32(0) {
		goto L220
	} else {
		goto L221
	}
L219:
	;
	v800 = *(*int32)(unsafe.Add(mBase, uint32(v799)+20))
	if v800 != int32(-1) {
		goto L222
	} else {
		goto L223
	}
L220:
	;
	v780 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v772)+16)))
	v781 = int64(*(*int8)(unsafe.Add(mBase, uint32(v772)+27)))
	v782 = int64(*(*int32)(unsafe.Add(mBase, uint32(v772)+8)))
	v783 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v772)+12)))
	v784 = int64(*(*int8)(unsafe.Add(mBase, uint32(v772)+26)))
	v785 = int64(*(*int32)(unsafe.Add(mBase, uint32(v772)+4)))
	v786 = F_wangHash64(m, v785)
	mBase = m.M
	v788 = F_wangHash64(m, v784+v786)
	mBase = m.M
	v790 = F_wangHash64(m, v783+v788)
	mBase = m.M
	v792 = F_wangHash64(m, v782+v790)
	mBase = m.M
	v794 = F_wangHash64(m, v781+v792)
	mBase = m.M
	v796 = F_wangHash64(m, v780+v794)
	mBase = m.M
	*(*int64)(unsafe.Add(mBase, uint32(v747)+24)) = v796
	v798 = *(*int32)(unsafe.Add(mBase, uint32(v747)))
	v799 = v798
	goto L219
L221:
	;
	v776 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v772)+24)))
	v778 = v776 + int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v772)+24)) = uint16(v778)
	v799 = v772
	goto L219
L222:
	;
	v806 = v800 + int32(-1)
	goto L216
L223:
	;
	v803 = *(*int32)(unsafe.Add(mBase, uint32(v747)+4))
	v806 = v803
	goto L216
L224:
	;
	v833 = int32(2)
	v838 = *(*int32)(unsafe.Add(mBase, uint32(v813+v831<<(uint(v833)%32)+int32(4))))
	v762 = v838 + v832<<(uint(v833)%32)
	v763 = int32(1)
	goto L210
L225:
	;
	v822 = v810
	goto L227
L226:
	;
	v822 = v807 << (uint(v818) % 32)
	goto L227
L227:
	;
	if v808 < v822 {
		v831 = v814
		v832 = v808
		goto L224
	} else {
		goto L228
	}
L228:
	;
	if v814 != 0 {
		v851 = v810
		goto L213
	} else {
		goto L229
	}
L229:
	;
	v824 = *(*int32)(unsafe.Add(mBase, uint32(v813)+20))
	if v824 == int32(-1) {
		v851 = v810
		goto L213
	} else {
		goto L230
	}
L230:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v747)+4)) = int64(4294967296)
	v831 = int32(1)
	v832 = int32(0)
	goto L224
L231:
	;
	v847 = *(*int32)(unsafe.Add(mBase, uint32(v843)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v755))) = v847
	v851 = v843
	goto L213
L232:
	;
	if v732&int32(1) == int32(0) {
		goto L61
	} else {
		goto L233
	}
L233:
	;
	goto L62
L234:
	;
	if v967 != 0 {
		v730 = v967
		v732 = int32(0)
		goto L201
	} else {
		goto L260
	}
L235:
	;
	v878 = v871
	v879 = v875
	goto L238
L236:
	;
	v875 = int32(1)
	goto L235
L237:
	;
	v875 = int32(0)
	goto L235
L238:
	;
	switch v879 {
	case 0:
		goto L243
	default:
		goto L242
	}
L240:
	;
	v879 = int32(0)
	goto L238
L241:
	;
	goto L234
L242:
	;
	v959 = *(*int32)(unsafe.Add(mBase, uint32(v878)))
	*(*int32)(unsafe.Add(mBase, uint32(v863)+16)) = v959
	if v959 == int32(0) {
		goto L240
	} else {
		goto L259
	}
L243:
	;
	v883 = *(*int32)(unsafe.Add(mBase, uint32(v863)+4))
	if v883 != int32(-1) {
		v922 = v883
		goto L244
	} else {
		goto L245
	}
L244:
	;
	v923 = int32(1)
	v924 = v922 + v923
	*(*int32)(unsafe.Add(mBase, uint32(v863)+4)) = v924
	v926 = int32(0)
	v929 = *(*int32)(unsafe.Add(mBase, uint32(v863)))
	v930 = *(*int32)(unsafe.Add(mBase, uint32(v863)+8))
	v934 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v929+v930+int32(26)))))
	if v934 == int32(255) {
		goto L253
	} else {
		goto L254
	}
L245:
	;
	v887 = *(*int32)(unsafe.Add(mBase, uint32(v863)+8))
	if v887 != 0 {
		v922 = int32(-1)
		goto L244
	} else {
		goto L246
	}
L246:
	;
	v888 = *(*int32)(unsafe.Add(mBase, uint32(v863)))
	v889 = *(*int32)(unsafe.Add(mBase, uint32(v863)+12))
	if v889 == int32(0) {
		goto L248
	} else {
		goto L249
	}
L247:
	;
	v916 = *(*int32)(unsafe.Add(mBase, uint32(v915)+20))
	if v916 != int32(-1) {
		goto L250
	} else {
		goto L251
	}
L248:
	;
	v896 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v888)+16)))
	v897 = int64(*(*int8)(unsafe.Add(mBase, uint32(v888)+27)))
	v898 = int64(*(*int32)(unsafe.Add(mBase, uint32(v888)+8)))
	v899 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v888)+12)))
	v900 = int64(*(*int8)(unsafe.Add(mBase, uint32(v888)+26)))
	v901 = int64(*(*int32)(unsafe.Add(mBase, uint32(v888)+4)))
	v902 = F_wangHash64(m, v901)
	mBase = m.M
	v904 = F_wangHash64(m, v900+v902)
	mBase = m.M
	v906 = F_wangHash64(m, v899+v904)
	mBase = m.M
	v908 = F_wangHash64(m, v898+v906)
	mBase = m.M
	v910 = F_wangHash64(m, v897+v908)
	mBase = m.M
	v912 = F_wangHash64(m, v896+v910)
	mBase = m.M
	*(*int64)(unsafe.Add(mBase, uint32(v863)+24)) = v912
	v914 = *(*int32)(unsafe.Add(mBase, uint32(v863)))
	v915 = v914
	goto L247
L249:
	;
	v892 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v888)+24)))
	v894 = v892 + int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v888)+24)) = uint16(v894)
	v915 = v888
	goto L247
L250:
	;
	v922 = v916 + int32(-1)
	goto L244
L251:
	;
	v919 = *(*int32)(unsafe.Add(mBase, uint32(v863)+4))
	v922 = v919
	goto L244
L252:
	;
	v949 = int32(2)
	v954 = *(*int32)(unsafe.Add(mBase, uint32(v929+v947<<(uint(v949)%32)+int32(4))))
	v878 = v954 + v948<<(uint(v949)%32)
	v879 = int32(1)
	goto L238
L253:
	;
	v938 = v926
	goto L255
L254:
	;
	v938 = v923 << (uint(v934) % 32)
	goto L255
L255:
	;
	if v924 < v938 {
		v947 = v930
		v948 = v924
		goto L252
	} else {
		goto L256
	}
L256:
	;
	if v930 != 0 {
		v967 = v926
		goto L241
	} else {
		goto L257
	}
L257:
	;
	v940 = *(*int32)(unsafe.Add(mBase, uint32(v929)+20))
	if v940 == int32(-1) {
		v967 = v926
		goto L241
	} else {
		goto L258
	}
L258:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v863)+4)) = int64(4294967296)
	v947 = int32(1)
	v948 = int32(0)
	goto L252
L259:
	;
	v963 = *(*int32)(unsafe.Add(mBase, uint32(v959)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v871))) = v963
	v967 = v959
	goto L241
L260:
	;
	goto L61
L261:
	;
	v1299 = v23
	goto L1
L262:
	;
	v1059 = m.G4
	v1060 = *(*int32)(unsafe.Add(mBase, uint32(v1059)))
	v1061 = m.T0[v1060].(func(*base.Module, int32) int32)(m, int32(24))
	mBase = m.M
	v1062 = m.ExcPending
	if v1062 != 0 {
		goto L93
	} else {
		goto L285
	}
L263:
	;
	if v1023-v1025 != 0 {
		goto L262
	} else {
		goto L278
	}
L264:
	;
	v984 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v227))))
	if v984 != 0 {
		goto L267
	} else {
		goto L268
	}
L266:
	;
	v1023 = F_tolower(m, v1018)
	mBase = m.M
	v1024 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1019))))
	v1025 = F_tolower(m, v1024)
	mBase = m.M
	goto L263
L267:
	;
	v986 = v227
	v987 = v979
	v988 = int32(9)
	v989 = v984
	goto L270
L268:
	;
	v1018 = int32(0)
	v1019 = v979
	goto L266
L269:
	;
	v1018 = v1015 & int32(255)
	v1019 = v1013
	goto L266
L270:
	;
	v991 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v987))))
	if v991 == int32(0) {
		v1013 = v987
		v1015 = v989
		goto L269
	} else {
		goto L272
	}
L271:
	;
	v1013 = v1007
	v1015 = int32(0)
	goto L269
L272:
	;
	v995 = v988 + int32(-1)
	if v995 == int32(0) {
		v1013 = v987
		v1015 = v989
		goto L269
	} else {
		goto L273
	}
L273:
	;
	v999 = v989 & int32(255)
	if v999 == v991 {
		goto L274
	} else {
		goto L275
	}
L274:
	;
	v1006 = int32(1)
	v1007 = v987 + v1006
	v1008 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v986)+1)))
	if v1008 != 0 {
		v986 = v986 + v1006
		v987 = v1007
		v988 = v995
		v989 = v1008
		goto L270
	} else {
		goto L277
	}
L275:
	;
	v1001 = F_tolower(m, v999)
	mBase = m.M
	v1002 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v987))))
	v1003 = F_tolower(m, v1002)
	mBase = m.M
	if v1001 == v1003 {
		goto L274
	} else {
		goto L276
	}
L276:
	;
	v1005 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v986))))
	v1013 = v987
	v1015 = v1005
	goto L269
L277:
	;
	goto L271
L278:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+140)) = v977 | int32(64)
	v1037 = m.G4
	v1038 = *(*int32)(unsafe.Add(mBase, uint32(v1037)))
	v1039 = m.T0[v1038].(func(*base.Module, int32) int32)(m, int32(24))
	mBase = m.M
	v1040 = m.ExcPending
	if v1040 != 0 {
		goto L93
	} else {
		goto L279
	}
L279:
	;
	if v1039 == int32(0) {
		goto L64
	} else {
		goto L280
	}
L280:
	;
	v1043 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v1039)+20)) = v1043
	*(*int32)(unsafe.Add(mBase, uint32(v1039)+16)) = l2
	*(*int64)(unsafe.Add(mBase, uint32(v1039)+8)) = int64(1)
	*(*int32)(unsafe.Add(mBase, uint32(v1039)+4)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v1039))) = v1043
	v1051 = *(*int32)(unsafe.Add(mBase, uint32(l0)+256))
	if v1051 != 0 {
		goto L281
	} else {
		goto L282
	}
L281:
	;
	v1053 = *(*int32)(unsafe.Add(mBase, uint32(l0)+260))
	if v1053 == int32(0) {
		goto L283
	} else {
		goto L284
	}
L282:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+256)) = v1039
	goto L281
L283:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+260)) = v1039
	goto L61
L284:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1053))) = v1039
	goto L283
L285:
	;
	if v977&int32(32) == int32(0) {
		goto L286
	} else {
		goto L287
	}
L286:
	;
	if v1061 == int32(0) {
		goto L64
	} else {
		goto L293
	}
L287:
	;
	if v1061 == int32(0) {
		goto L64
	} else {
		goto L288
	}
L288:
	;
	v1069 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v1061)+20)) = v1069
	*(*int32)(unsafe.Add(mBase, uint32(v1061)+16)) = l2
	*(*int64)(unsafe.Add(mBase, uint32(v1061)+8)) = int64(1)
	*(*int32)(unsafe.Add(mBase, uint32(v1061)+4)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v1061))) = v1069
	v1077 = *(*int32)(unsafe.Add(mBase, uint32(l0)+272))
	if v1077 != 0 {
		goto L289
	} else {
		goto L290
	}
L289:
	;
	v1079 = *(*int32)(unsafe.Add(mBase, uint32(l0)+276))
	if v1079 == int32(0) {
		goto L291
	} else {
		goto L292
	}
L290:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+272)) = v1061
	goto L289
L291:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+276)) = v1061
	goto L61
L292:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1079))) = v1061
	goto L291
L293:
	;
	v1086 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v1061)+20)) = v1086
	*(*int32)(unsafe.Add(mBase, uint32(v1061)+16)) = l2
	*(*int64)(unsafe.Add(mBase, uint32(v1061)+8)) = int64(1)
	*(*int32)(unsafe.Add(mBase, uint32(v1061)+4)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v1061))) = v1086
	v1094 = *(*int32)(unsafe.Add(mBase, uint32(l0)+256))
	if v1094 != 0 {
		goto L294
	} else {
		goto L295
	}
L294:
	;
	v1096 = *(*int32)(unsafe.Add(mBase, uint32(l0)+260))
	if v1096 == int32(0) {
		goto L296
	} else {
		goto L297
	}
L295:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+256)) = v1061
	goto L294
L296:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+260)) = v1061
	goto L61
L297:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1096))) = v1061
	goto L296
L298:
	;
	v1189 = *(*int32)(unsafe.Add(mBase, uint32(l0)+244))
	v1190 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+140)))
	if v1190&int32(2) == int32(0) {
		goto L301
	} else {
		goto L302
	}
L299:
	;
	v1229 = int32(0)
	v1230 = *(*int32)(unsafe.Add(mBase, uint32(l0)+232))
	if v1230 == v1229 {
		v1299 = v1229
		goto L1
	} else {
		goto L312
	}
L300:
	;
	v1216 = *(*int32)(unsafe.Add(mBase, uint32(l0)+220))
	v1217 = int32(8)
	v1221 = *(*int64)(unsafe.Add(mBase, uint32(v1215+v1217)))
	*(*int64)(unsafe.Add(mBase, uint32(v21+v1217))) = v1221
	v1223 = *(*int64)(unsafe.Add(mBase, uint32(v1215)))
	*(*int64)(unsafe.Add(mBase, uint32(v21))) = v1223
	m.T0[v1189].(func(*base.Module, int32, int32))(m, v1216, v21)
	mBase = m.M
	v1226 = m.ExcPending
	if v1226 != 0 {
		goto L93
	} else {
		goto L311
	}
L301:
	;
	if v1189 == int32(0) {
		goto L299
	} else {
		goto L307
	}
L302:
	;
	if v1189 == int32(0) {
		goto L299
	} else {
		goto L303
	}
L303:
	;
	v1197 = *(*int32)(unsafe.Add(mBase, uint32(l0)+160))
	if v1197 == int32(0) {
		goto L299
	} else {
		goto L304
	}
L304:
	;
	v1200 = *(*int64)(unsafe.Add(mBase, uint32(v1197)))
	if v1200 != int64(0) {
		v1215 = v1197
		goto L300
	} else {
		goto L305
	}
L305:
	;
	v1203 = *(*int32)(unsafe.Add(mBase, uint32(v1197)+8))
	if v1203 != 0 {
		v1215 = v1197
		goto L300
	} else {
		goto L306
	}
L306:
	;
	goto L299
L307:
	;
	v1206 = *(*int32)(unsafe.Add(mBase, uint32(l0)+156))
	if v1206 == int32(0) {
		goto L299
	} else {
		goto L308
	}
L308:
	;
	v1209 = *(*int64)(unsafe.Add(mBase, uint32(v1206)))
	if v1209 != int64(0) {
		v1215 = v1206
		goto L300
	} else {
		goto L309
	}
L309:
	;
	v1212 = *(*int32)(unsafe.Add(mBase, uint32(v1206)+8))
	if v1212 == int32(0) {
		goto L299
	} else {
		goto L310
	}
L310:
	;
	v1215 = v1206
	goto L300
L311:
	;
	goto L299
L312:
	;
	v1233 = *(*int32)(unsafe.Add(mBase, uint32(l0)+220))
	m.T0[v1230].(func(*base.Module, int32))(m, v1233)
	mBase = m.M
	v1235 = m.ExcPending
	if v1235 != 0 {
		goto L93
	} else {
		goto L313
	}
L313:
	;
	v1299 = v1229
	goto L1
L314:
	;
	v1285 = m.G4
	v1286 = *(*int32)(unsafe.Add(mBase, uint32(v1263)))
	v1287 = *(*int32)(unsafe.Add(mBase, uint32(v1285)+16))
	m.T0[v1287].(func(*base.Module, int32))(m, v1286)
	mBase = m.M
	v1289 = m.ExcPending
	if v1289 != 0 {
		goto L93
	} else {
		goto L315
	}
L315:
	;
	v1290 = *(*int32)(unsafe.Add(mBase, uint32(v1285)+16))
	m.T0[v1290].(func(*base.Module, int32))(m, v1263)
	mBase = m.M
	v1292 = m.ExcPending
	if v1292 != 0 {
		goto L93
	} else {
		goto L316
	}
L316:
	;
	v1299 = v23
	goto L1
}
func F_valkeyAsyncCommand(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v21 int32
	_ = v21
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
	var v30 int32
	_ = v30
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	*(*int32)(unsafe.Add(mBase, uint32(v9)+8)) = l4
	v14 = F_valkeyvFormatCommand(m, v9+int32(12), l3, l4)
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		return int32(0)
	} else {
		if int32(0) <= v14 {
			v21 = m.G4
			v22 = *(*int32)(unsafe.Add(mBase, uint32(v9)+12))
			v23 = F_valkeyAsyncAppendCmdLen(m, l0, l1, l2, v22, v14)
			mBase = m.M
			v24 = m.ExcPending
			if v24 != 0 {
				return int32(0)
			} else {
				v25 = *(*int32)(unsafe.Add(mBase, uint32(v9)+12))
				v26 = *(*int32)(unsafe.Add(mBase, uint32(v21)+16))
				m.T0[v26].(func(*base.Module, int32))(m, v25)
				mBase = m.M
				v28 = m.ExcPending
				if v28 != 0 {
					return int32(0)
				} else {
					v30 = v23
					m.G0 = v9 + int32(16)
					return v30
				}
			}
		} else {
			v30 = int32(-1)
			m.G0 = v9 + int32(16)
			return v30
		}
	}
}
func F_valkeyAsyncConnectBind(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v13 int64
	_ = v13
	var v16 int32
	_ = v16
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	v7 = m.G0
	v9 = v7 - int32(48)
	m.G0 = v9
	v12 = v9 + int32(24)
	v13 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v12))) = v13
	v16 = v9 + int32(16)
	*(*int64)(unsafe.Add(mBase, uint32(v16))) = v13
	*(*int32)(unsafe.Add(mBase, uint32(v9+int32(40)))) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v9+int32(32)))) = v13
	*(*int64)(unsafe.Add(mBase, uint32(v9+int32(8)))) = v13
	*(*int32)(unsafe.Add(mBase, uint32(v12))) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v16))) = l2
	*(*int64)(unsafe.Add(mBase, uint32(v9))) = v13
	*(*int32)(unsafe.Add(mBase, uint32(v9)+20)) = l0
	v36 = F_valkeyAsyncConnectWithOptions(m, v9)
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		return int32(0)
	} else {
		m.G0 = v9 + int32(48)
		return v36
	}
}
func F_valkeyAsyncFree(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v12 int32
	_ = v12
	if l0 == int32(0) {
		return
	} else {
		v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+140))
		*(*int32)(unsafe.Add(mBase, uint32(l0)+140)) = v5 | int32(8)
		if v5&int32(16) != 0 {
			return
		} else {
			F_valkeyAsyncFreeInternal(m, l0)
			mBase = m.M
			v12 = m.ExcPending
			if v12 != 0 {
				return
			} else {
				return
			}
		}
	}
}
func F_valkeyAsyncFreeInternal(m *base.Module, l0 int32) {
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
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v49 int32
	_ = v49
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v82 int32
	_ = v82
	var v88 int32
	_ = v88
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v109 int32
	_ = v109
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v117 int32
	_ = v117
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v126 int32
	_ = v126
	var v128 int32
	_ = v128
	var v130 int64
	_ = v130
	var v131 int64
	_ = v131
	var v132 int64
	_ = v132
	var v133 int64
	_ = v133
	var v134 int64
	_ = v134
	var v135 int64
	_ = v135
	var v136 int64
	_ = v136
	var v138 int64
	_ = v138
	var v140 int64
	_ = v140
	var v142 int64
	_ = v142
	var v144 int64
	_ = v144
	var v146 int64
	_ = v146
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v153 int32
	_ = v153
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v160 int32
	_ = v160
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v168 int32
	_ = v168
	var v172 int32
	_ = v172
	var v174 int32
	_ = v174
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v188 int32
	_ = v188
	var v193 int32
	_ = v193
	var v197 int32
	_ = v197
	var v201 int32
	_ = v201
	var v209 int32
	_ = v209
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v216 int32
	_ = v216
	var v221 int32
	_ = v221
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
	var v239 int32
	_ = v239
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v247 int32
	_ = v247
	var v251 int32
	_ = v251
	var v252 int32
	_ = v252
	var v253 int32
	_ = v253
	var v256 int32
	_ = v256
	var v258 int32
	_ = v258
	var v260 int64
	_ = v260
	var v261 int64
	_ = v261
	var v262 int64
	_ = v262
	var v263 int64
	_ = v263
	var v264 int64
	_ = v264
	var v265 int64
	_ = v265
	var v266 int64
	_ = v266
	var v268 int64
	_ = v268
	var v270 int64
	_ = v270
	var v272 int64
	_ = v272
	var v274 int64
	_ = v274
	var v276 int64
	_ = v276
	var v278 int32
	_ = v278
	var v279 int32
	_ = v279
	var v280 int32
	_ = v280
	var v283 int32
	_ = v283
	var v286 int32
	_ = v286
	var v287 int32
	_ = v287
	var v288 int32
	_ = v288
	var v290 int32
	_ = v290
	var v293 int32
	_ = v293
	var v294 int32
	_ = v294
	var v298 int32
	_ = v298
	var v302 int32
	_ = v302
	var v304 int32
	_ = v304
	var v311 int32
	_ = v311
	var v312 int32
	_ = v312
	var v313 int32
	_ = v313
	var v318 int32
	_ = v318
	var v323 int32
	_ = v323
	var v327 int32
	_ = v327
	var v331 int32
	_ = v331
	var v340 int32
	_ = v340
	var v342 int32
	_ = v342
	var v348 int32
	_ = v348
	var v365 int32
	_ = v365
	var v366 int32
	_ = v366
	var v369 int32
	_ = v369
	var v372 int32
	_ = v372
	var v373 int32
	_ = v373
	var v377 int32
	_ = v377
	var v381 int32
	_ = v381
	var v382 int32
	_ = v382
	var v383 int32
	_ = v383
	var v386 int32
	_ = v386
	var v388 int32
	_ = v388
	var v390 int64
	_ = v390
	var v391 int64
	_ = v391
	var v392 int64
	_ = v392
	var v393 int64
	_ = v393
	var v394 int64
	_ = v394
	var v395 int64
	_ = v395
	var v396 int64
	_ = v396
	var v398 int64
	_ = v398
	var v400 int64
	_ = v400
	var v402 int64
	_ = v402
	var v404 int64
	_ = v404
	var v406 int64
	_ = v406
	var v408 int32
	_ = v408
	var v409 int32
	_ = v409
	var v410 int32
	_ = v410
	var v413 int32
	_ = v413
	var v416 int32
	_ = v416
	var v417 int32
	_ = v417
	var v418 int32
	_ = v418
	var v420 int32
	_ = v420
	var v423 int32
	_ = v423
	var v424 int32
	_ = v424
	var v428 int32
	_ = v428
	var v432 int32
	_ = v432
	var v434 int32
	_ = v434
	var v441 int32
	_ = v441
	var v442 int32
	_ = v442
	var v443 int32
	_ = v443
	var v448 int32
	_ = v448
	var v453 int32
	_ = v453
	var v457 int32
	_ = v457
	var v461 int32
	_ = v461
	var v469 int32
	_ = v469
	var v472 int32
	_ = v472
	var v473 int32
	_ = v473
	var v476 int32
	_ = v476
	var v481 int32
	_ = v481
	var v483 int32
	_ = v483
	var v484 int32
	_ = v484
	var v495 int32
	_ = v495
	var v496 int32
	_ = v496
	var v499 int32
	_ = v499
	var v502 int32
	_ = v502
	var v503 int32
	_ = v503
	var v507 int32
	_ = v507
	var v511 int32
	_ = v511
	var v512 int32
	_ = v512
	var v513 int32
	_ = v513
	var v516 int32
	_ = v516
	var v518 int32
	_ = v518
	var v520 int64
	_ = v520
	var v521 int64
	_ = v521
	var v522 int64
	_ = v522
	var v523 int64
	_ = v523
	var v524 int64
	_ = v524
	var v525 int64
	_ = v525
	var v526 int64
	_ = v526
	var v528 int64
	_ = v528
	var v530 int64
	_ = v530
	var v532 int64
	_ = v532
	var v534 int64
	_ = v534
	var v536 int64
	_ = v536
	var v538 int32
	_ = v538
	var v539 int32
	_ = v539
	var v540 int32
	_ = v540
	var v543 int32
	_ = v543
	var v546 int32
	_ = v546
	var v547 int32
	_ = v547
	var v548 int32
	_ = v548
	var v550 int32
	_ = v550
	var v553 int32
	_ = v553
	var v554 int32
	_ = v554
	var v558 int32
	_ = v558
	var v562 int32
	_ = v562
	var v564 int32
	_ = v564
	var v571 int32
	_ = v571
	var v572 int32
	_ = v572
	var v573 int32
	_ = v573
	var v578 int32
	_ = v578
	var v583 int32
	_ = v583
	var v587 int32
	_ = v587
	var v591 int32
	_ = v591
	var v600 int32
	_ = v600
	var v602 int32
	_ = v602
	var v608 int32
	_ = v608
	var v625 int32
	_ = v625
	var v626 int32
	_ = v626
	var v629 int32
	_ = v629
	var v632 int32
	_ = v632
	var v633 int32
	_ = v633
	var v637 int32
	_ = v637
	var v641 int32
	_ = v641
	var v642 int32
	_ = v642
	var v643 int32
	_ = v643
	var v646 int32
	_ = v646
	var v648 int32
	_ = v648
	var v650 int64
	_ = v650
	var v651 int64
	_ = v651
	var v652 int64
	_ = v652
	var v653 int64
	_ = v653
	var v654 int64
	_ = v654
	var v655 int64
	_ = v655
	var v656 int64
	_ = v656
	var v658 int64
	_ = v658
	var v660 int64
	_ = v660
	var v662 int64
	_ = v662
	var v664 int64
	_ = v664
	var v666 int64
	_ = v666
	var v668 int32
	_ = v668
	var v669 int32
	_ = v669
	var v670 int32
	_ = v670
	var v673 int32
	_ = v673
	var v676 int32
	_ = v676
	var v677 int32
	_ = v677
	var v678 int32
	_ = v678
	var v680 int32
	_ = v680
	var v683 int32
	_ = v683
	var v684 int32
	_ = v684
	var v688 int32
	_ = v688
	var v692 int32
	_ = v692
	var v694 int32
	_ = v694
	var v701 int32
	_ = v701
	var v702 int32
	_ = v702
	var v703 int32
	_ = v703
	var v708 int32
	_ = v708
	var v713 int32
	_ = v713
	var v717 int32
	_ = v717
	var v721 int32
	_ = v721
	var v729 int32
	_ = v729
	var v732 int32
	_ = v732
	var v733 int32
	_ = v733
	var v736 int32
	_ = v736
	var v741 int32
	_ = v741
	var v743 int32
	_ = v743
	var v744 int32
	_ = v744
	var v755 int32
	_ = v755
	var v756 int32
	_ = v756
	var v759 int32
	_ = v759
	var v762 int32
	_ = v762
	var v763 int32
	_ = v763
	var v767 int32
	_ = v767
	var v771 int32
	_ = v771
	var v772 int32
	_ = v772
	var v773 int32
	_ = v773
	var v776 int32
	_ = v776
	var v778 int32
	_ = v778
	var v780 int64
	_ = v780
	var v781 int64
	_ = v781
	var v782 int64
	_ = v782
	var v783 int64
	_ = v783
	var v784 int64
	_ = v784
	var v785 int64
	_ = v785
	var v786 int64
	_ = v786
	var v788 int64
	_ = v788
	var v790 int64
	_ = v790
	var v792 int64
	_ = v792
	var v794 int64
	_ = v794
	var v796 int64
	_ = v796
	var v798 int32
	_ = v798
	var v799 int32
	_ = v799
	var v800 int32
	_ = v800
	var v803 int32
	_ = v803
	var v806 int32
	_ = v806
	var v807 int32
	_ = v807
	var v808 int32
	_ = v808
	var v810 int32
	_ = v810
	var v813 int32
	_ = v813
	var v814 int32
	_ = v814
	var v818 int32
	_ = v818
	var v822 int32
	_ = v822
	var v824 int32
	_ = v824
	var v831 int32
	_ = v831
	var v832 int32
	_ = v832
	var v833 int32
	_ = v833
	var v838 int32
	_ = v838
	var v843 int32
	_ = v843
	var v847 int32
	_ = v847
	var v851 int32
	_ = v851
	var v860 int32
	_ = v860
	var v862 int32
	_ = v862
	var v868 int32
	_ = v868
	var v871 int32
	_ = v871
	var v873 int32
	_ = v873
	var v874 int32
	_ = v874
	var v876 int32
	_ = v876
	var v881 int32
	_ = v881
	var v884 int32
	_ = v884
	var v889 int32
	_ = v889
	var v893 int32
	_ = v893
	var v900 int32
	_ = v900
	var v901 int32
	_ = v901
	var v906 int32
	_ = v906
	var v909 int32
	_ = v909
	var v912 int32
	_ = v912
	var v914 int32
	_ = v914
	var v916 int32
	_ = v916
	v6 = m.G0
	v8 = v6 - int32(32)
	m.G0 = v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+256))
	if v10 == int32(0) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v49 = *(*int32)(unsafe.Add(mBase, uint32(l0)+272))
	if v49 == int32(0) {
		goto L13
	} else {
		goto L14
	}
L2:
	;
	v15 = v10
	goto L3
L3:
	;
	v18 = *(*int32)(unsafe.Add(mBase, uint32(v15)))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+256)) = v18
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+260))
	if v15 != v20 {
		goto L5
	} else {
		goto L6
	}
L4:
	;
	goto L1
L5:
	;
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v15)+16))
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v15)+4))
	v26 = m.G4
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v26)+16))
	m.T0[v27].(func(*base.Module, int32))(m, v15)
	mBase = m.M
	v29 = m.ExcPending
	if v29 != 0 {
		goto L7
	} else {
		goto L8
	}
L6:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+260)) = int32(0)
	goto L5
L7:
	;
	return
L8:
	;
	if v25 == int32(0) {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(l0)+256))
	if v43 != 0 {
		v15 = v43
		goto L3
	} else {
		goto L12
	}
L10:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(l0)+140))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+140)) = v32 | int32(16)
	m.T0[v25].(func(*base.Module, int32, int32, int32))(m, l0, int32(0), v24)
	mBase = m.M
	v38 = m.ExcPending
	if v38 != 0 {
		goto L7
	} else {
		goto L11
	}
L11:
	;
	v39 = *(*int32)(unsafe.Add(mBase, uint32(l0)+140))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+140)) = v39 & int32(-17)
	goto L9
L12:
	;
	goto L4
L13:
	;
	v88 = *(*int32)(unsafe.Add(mBase, uint32(l0)+280))
	if v88 == int32(0) {
		goto L24
	} else {
		goto L25
	}
L14:
	;
	v54 = v49
	goto L15
L15:
	;
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v54)))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+272)) = v57
	v59 = *(*int32)(unsafe.Add(mBase, uint32(l0)+276))
	if v54 != v59 {
		goto L17
	} else {
		goto L18
	}
L16:
	;
	goto L13
L17:
	;
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v54)+16))
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v54)+4))
	v65 = m.G4
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v65)+16))
	m.T0[v66].(func(*base.Module, int32))(m, v54)
	mBase = m.M
	v68 = m.ExcPending
	if v68 != 0 {
		goto L7
	} else {
		goto L19
	}
L18:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+276)) = int32(0)
	goto L17
L19:
	;
	if v64 == int32(0) {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	v82 = *(*int32)(unsafe.Add(mBase, uint32(l0)+272))
	if v82 != 0 {
		v54 = v82
		goto L15
	} else {
		goto L23
	}
L21:
	;
	v71 = *(*int32)(unsafe.Add(mBase, uint32(l0)+140))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+140)) = v71 | int32(16)
	m.T0[v64].(func(*base.Module, int32, int32, int32))(m, l0, int32(0), v63)
	mBase = m.M
	v77 = m.ExcPending
	if v77 != 0 {
		goto L7
	} else {
		goto L22
	}
L22:
	;
	v78 = *(*int32)(unsafe.Add(mBase, uint32(l0)+140))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+140)) = v78 & int32(-17)
	goto L20
L23:
	;
	goto L16
L24:
	;
	v348 = *(*int32)(unsafe.Add(mBase, uint32(l0)+284))
	if v348 == int32(0) {
		goto L89
	} else {
		goto L90
	}
L25:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8))) = v88
	*(*int32)(unsafe.Add(mBase, uint32(v8)+20)) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v8)+12)) = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v8)+4)) = int64(4294967295)
	goto L26
L26:
	;
	v105 = v8 + int32(20)
	v106 = *(*int32)(unsafe.Add(mBase, uint32(v8)+16))
	if v106 != 0 {
		goto L30
	} else {
		goto L31
	}
L27:
	;
	v340 = *(*int32)(unsafe.Add(mBase, uint32(l0)+280))
	F_dictRelease(m, v340)
	mBase = m.M
	v342 = m.ExcPending
	if v342 != 0 {
		goto L7
	} else {
		goto L88
	}
L28:
	;
	if v201 == int32(0) {
		goto L27
	} else {
		goto L54
	}
L29:
	;
	v112 = v105
	v113 = v109
	goto L32
L30:
	;
	v109 = int32(1)
	goto L29
L31:
	;
	v109 = int32(0)
	goto L29
L32:
	;
	switch v113 {
	case 0:
		goto L37
	default:
		goto L36
	}
L34:
	;
	v113 = int32(0)
	goto L32
L35:
	;
	goto L28
L36:
	;
	v193 = *(*int32)(unsafe.Add(mBase, uint32(v112)))
	*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = v193
	if v193 == int32(0) {
		goto L34
	} else {
		goto L53
	}
L37:
	;
	v117 = *(*int32)(unsafe.Add(mBase, uint32(v8)+4))
	if v117 != int32(-1) {
		v156 = v117
		goto L38
	} else {
		goto L39
	}
L38:
	;
	v157 = int32(1)
	v158 = v156 + v157
	*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = v158
	v160 = int32(0)
	v163 = *(*int32)(unsafe.Add(mBase, uint32(v8)))
	v164 = *(*int32)(unsafe.Add(mBase, uint32(v8)+8))
	v168 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v163+v164+int32(26)))))
	if v168 == int32(255) {
		goto L47
	} else {
		goto L48
	}
L39:
	;
	v121 = *(*int32)(unsafe.Add(mBase, uint32(v8)+8))
	if v121 != 0 {
		v156 = int32(-1)
		goto L38
	} else {
		goto L40
	}
L40:
	;
	v122 = *(*int32)(unsafe.Add(mBase, uint32(v8)))
	v123 = *(*int32)(unsafe.Add(mBase, uint32(v8)+12))
	if v123 == int32(0) {
		goto L42
	} else {
		goto L43
	}
L41:
	;
	v150 = *(*int32)(unsafe.Add(mBase, uint32(v149)+20))
	if v150 != int32(-1) {
		goto L44
	} else {
		goto L45
	}
L42:
	;
	v130 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v122)+16)))
	v131 = int64(*(*int8)(unsafe.Add(mBase, uint32(v122)+27)))
	v132 = int64(*(*int32)(unsafe.Add(mBase, uint32(v122)+8)))
	v133 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v122)+12)))
	v134 = int64(*(*int8)(unsafe.Add(mBase, uint32(v122)+26)))
	v135 = int64(*(*int32)(unsafe.Add(mBase, uint32(v122)+4)))
	v136 = F_wangHash64(m, v135)
	mBase = m.M
	v138 = F_wangHash64(m, v134+v136)
	mBase = m.M
	v140 = F_wangHash64(m, v133+v138)
	mBase = m.M
	v142 = F_wangHash64(m, v132+v140)
	mBase = m.M
	v144 = F_wangHash64(m, v131+v142)
	mBase = m.M
	v146 = F_wangHash64(m, v130+v144)
	mBase = m.M
	*(*int64)(unsafe.Add(mBase, uint32(v8)+24)) = v146
	v148 = *(*int32)(unsafe.Add(mBase, uint32(v8)))
	v149 = v148
	goto L41
L43:
	;
	v126 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v122)+24)))
	v128 = v126 + int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v122)+24)) = uint16(v128)
	v149 = v122
	goto L41
L44:
	;
	v156 = v150 + int32(-1)
	goto L38
L45:
	;
	v153 = *(*int32)(unsafe.Add(mBase, uint32(v8)+4))
	v156 = v153
	goto L38
L46:
	;
	v183 = int32(2)
	v188 = *(*int32)(unsafe.Add(mBase, uint32(v163+v181<<(uint(v183)%32)+int32(4))))
	v112 = v188 + v182<<(uint(v183)%32)
	v113 = int32(1)
	goto L32
L47:
	;
	v172 = v160
	goto L49
L48:
	;
	v172 = v157 << (uint(v168) % 32)
	goto L49
L49:
	;
	if v158 < v172 {
		v181 = v164
		v182 = v158
		goto L46
	} else {
		goto L50
	}
L50:
	;
	if v164 != 0 {
		v201 = v160
		goto L35
	} else {
		goto L51
	}
L51:
	;
	v174 = *(*int32)(unsafe.Add(mBase, uint32(v163)+20))
	if v174 == int32(-1) {
		v201 = v160
		goto L35
	} else {
		goto L52
	}
L52:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v8)+4)) = int64(4294967296)
	v181 = int32(1)
	v182 = int32(0)
	goto L46
L53:
	;
	v197 = *(*int32)(unsafe.Add(mBase, uint32(v193)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v105))) = v197
	v201 = v193
	goto L35
L54:
	;
	v209 = v201
	goto L55
L55:
	;
	v212 = *(*int32)(unsafe.Add(mBase, uint32(v209)+8))
	goto L58
L56:
	;
	goto L27
L57:
	;
	v235 = v8 + int32(20)
	v236 = *(*int32)(unsafe.Add(mBase, uint32(v8)+16))
	if v236 != 0 {
		goto L63
	} else {
		goto L64
	}
L58:
	;
	v213 = *(*int32)(unsafe.Add(mBase, uint32(v212)+4))
	if v213 == int32(0) {
		goto L57
	} else {
		goto L59
	}
L59:
	;
	v216 = *(*int32)(unsafe.Add(mBase, uint32(l0)+140))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+140)) = v216 | int32(16)
	v221 = *(*int32)(unsafe.Add(mBase, uint32(v212)+16))
	m.T0[v213].(func(*base.Module, int32, int32, int32))(m, l0, int32(0), v221)
	mBase = m.M
	v223 = m.ExcPending
	if v223 != 0 {
		goto L7
	} else {
		goto L60
	}
L60:
	;
	v224 = *(*int32)(unsafe.Add(mBase, uint32(l0)+140))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+140)) = v224 & int32(-17)
	goto L57
L61:
	;
	if v331 != 0 {
		v209 = v331
		goto L55
	} else {
		goto L87
	}
L62:
	;
	v242 = v235
	v243 = v239
	goto L65
L63:
	;
	v239 = int32(1)
	goto L62
L64:
	;
	v239 = int32(0)
	goto L62
L65:
	;
	switch v243 {
	case 0:
		goto L70
	default:
		goto L69
	}
L67:
	;
	v243 = int32(0)
	goto L65
L68:
	;
	goto L61
L69:
	;
	v323 = *(*int32)(unsafe.Add(mBase, uint32(v242)))
	*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = v323
	if v323 == int32(0) {
		goto L67
	} else {
		goto L86
	}
L70:
	;
	v247 = *(*int32)(unsafe.Add(mBase, uint32(v8)+4))
	if v247 != int32(-1) {
		v286 = v247
		goto L71
	} else {
		goto L72
	}
L71:
	;
	v287 = int32(1)
	v288 = v286 + v287
	*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = v288
	v290 = int32(0)
	v293 = *(*int32)(unsafe.Add(mBase, uint32(v8)))
	v294 = *(*int32)(unsafe.Add(mBase, uint32(v8)+8))
	v298 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v293+v294+int32(26)))))
	if v298 == int32(255) {
		goto L80
	} else {
		goto L81
	}
L72:
	;
	v251 = *(*int32)(unsafe.Add(mBase, uint32(v8)+8))
	if v251 != 0 {
		v286 = int32(-1)
		goto L71
	} else {
		goto L73
	}
L73:
	;
	v252 = *(*int32)(unsafe.Add(mBase, uint32(v8)))
	v253 = *(*int32)(unsafe.Add(mBase, uint32(v8)+12))
	if v253 == int32(0) {
		goto L75
	} else {
		goto L76
	}
L74:
	;
	v280 = *(*int32)(unsafe.Add(mBase, uint32(v279)+20))
	if v280 != int32(-1) {
		goto L77
	} else {
		goto L78
	}
L75:
	;
	v260 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v252)+16)))
	v261 = int64(*(*int8)(unsafe.Add(mBase, uint32(v252)+27)))
	v262 = int64(*(*int32)(unsafe.Add(mBase, uint32(v252)+8)))
	v263 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v252)+12)))
	v264 = int64(*(*int8)(unsafe.Add(mBase, uint32(v252)+26)))
	v265 = int64(*(*int32)(unsafe.Add(mBase, uint32(v252)+4)))
	v266 = F_wangHash64(m, v265)
	mBase = m.M
	v268 = F_wangHash64(m, v264+v266)
	mBase = m.M
	v270 = F_wangHash64(m, v263+v268)
	mBase = m.M
	v272 = F_wangHash64(m, v262+v270)
	mBase = m.M
	v274 = F_wangHash64(m, v261+v272)
	mBase = m.M
	v276 = F_wangHash64(m, v260+v274)
	mBase = m.M
	*(*int64)(unsafe.Add(mBase, uint32(v8)+24)) = v276
	v278 = *(*int32)(unsafe.Add(mBase, uint32(v8)))
	v279 = v278
	goto L74
L76:
	;
	v256 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v252)+24)))
	v258 = v256 + int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v252)+24)) = uint16(v258)
	v279 = v252
	goto L74
L77:
	;
	v286 = v280 + int32(-1)
	goto L71
L78:
	;
	v283 = *(*int32)(unsafe.Add(mBase, uint32(v8)+4))
	v286 = v283
	goto L71
L79:
	;
	v313 = int32(2)
	v318 = *(*int32)(unsafe.Add(mBase, uint32(v293+v311<<(uint(v313)%32)+int32(4))))
	v242 = v318 + v312<<(uint(v313)%32)
	v243 = int32(1)
	goto L65
L80:
	;
	v302 = v290
	goto L82
L81:
	;
	v302 = v287 << (uint(v298) % 32)
	goto L82
L82:
	;
	if v288 < v302 {
		v311 = v294
		v312 = v288
		goto L79
	} else {
		goto L83
	}
L83:
	;
	if v294 != 0 {
		v331 = v290
		goto L68
	} else {
		goto L84
	}
L84:
	;
	v304 = *(*int32)(unsafe.Add(mBase, uint32(v293)+20))
	if v304 == int32(-1) {
		v331 = v290
		goto L68
	} else {
		goto L85
	}
L85:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v8)+4)) = int64(4294967296)
	v311 = int32(1)
	v312 = int32(0)
	goto L79
L86:
	;
	v327 = *(*int32)(unsafe.Add(mBase, uint32(v323)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v235))) = v327
	v331 = v323
	goto L68
L87:
	;
	goto L56
L88:
	;
	goto L24
L89:
	;
	v608 = *(*int32)(unsafe.Add(mBase, uint32(l0)+288))
	if v608 == int32(0) {
		goto L154
	} else {
		goto L155
	}
L90:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8))) = v348
	*(*int32)(unsafe.Add(mBase, uint32(v8)+20)) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v8)+12)) = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v8)+4)) = int64(4294967295)
	goto L91
L91:
	;
	v365 = v8 + int32(20)
	v366 = *(*int32)(unsafe.Add(mBase, uint32(v8)+16))
	if v366 != 0 {
		goto L95
	} else {
		goto L96
	}
L92:
	;
	v600 = *(*int32)(unsafe.Add(mBase, uint32(l0)+284))
	F_dictRelease(m, v600)
	mBase = m.M
	v602 = m.ExcPending
	if v602 != 0 {
		goto L7
	} else {
		goto L153
	}
L93:
	;
	if v461 == int32(0) {
		goto L92
	} else {
		goto L119
	}
L94:
	;
	v372 = v365
	v373 = v369
	goto L97
L95:
	;
	v369 = int32(1)
	goto L94
L96:
	;
	v369 = int32(0)
	goto L94
L97:
	;
	switch v373 {
	case 0:
		goto L102
	default:
		goto L101
	}
L99:
	;
	v373 = int32(0)
	goto L97
L100:
	;
	goto L93
L101:
	;
	v453 = *(*int32)(unsafe.Add(mBase, uint32(v372)))
	*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = v453
	if v453 == int32(0) {
		goto L99
	} else {
		goto L118
	}
L102:
	;
	v377 = *(*int32)(unsafe.Add(mBase, uint32(v8)+4))
	if v377 != int32(-1) {
		v416 = v377
		goto L103
	} else {
		goto L104
	}
L103:
	;
	v417 = int32(1)
	v418 = v416 + v417
	*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = v418
	v420 = int32(0)
	v423 = *(*int32)(unsafe.Add(mBase, uint32(v8)))
	v424 = *(*int32)(unsafe.Add(mBase, uint32(v8)+8))
	v428 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v423+v424+int32(26)))))
	if v428 == int32(255) {
		goto L112
	} else {
		goto L113
	}
L104:
	;
	v381 = *(*int32)(unsafe.Add(mBase, uint32(v8)+8))
	if v381 != 0 {
		v416 = int32(-1)
		goto L103
	} else {
		goto L105
	}
L105:
	;
	v382 = *(*int32)(unsafe.Add(mBase, uint32(v8)))
	v383 = *(*int32)(unsafe.Add(mBase, uint32(v8)+12))
	if v383 == int32(0) {
		goto L107
	} else {
		goto L108
	}
L106:
	;
	v410 = *(*int32)(unsafe.Add(mBase, uint32(v409)+20))
	if v410 != int32(-1) {
		goto L109
	} else {
		goto L110
	}
L107:
	;
	v390 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v382)+16)))
	v391 = int64(*(*int8)(unsafe.Add(mBase, uint32(v382)+27)))
	v392 = int64(*(*int32)(unsafe.Add(mBase, uint32(v382)+8)))
	v393 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v382)+12)))
	v394 = int64(*(*int8)(unsafe.Add(mBase, uint32(v382)+26)))
	v395 = int64(*(*int32)(unsafe.Add(mBase, uint32(v382)+4)))
	v396 = F_wangHash64(m, v395)
	mBase = m.M
	v398 = F_wangHash64(m, v394+v396)
	mBase = m.M
	v400 = F_wangHash64(m, v393+v398)
	mBase = m.M
	v402 = F_wangHash64(m, v392+v400)
	mBase = m.M
	v404 = F_wangHash64(m, v391+v402)
	mBase = m.M
	v406 = F_wangHash64(m, v390+v404)
	mBase = m.M
	*(*int64)(unsafe.Add(mBase, uint32(v8)+24)) = v406
	v408 = *(*int32)(unsafe.Add(mBase, uint32(v8)))
	v409 = v408
	goto L106
L108:
	;
	v386 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v382)+24)))
	v388 = v386 + int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v382)+24)) = uint16(v388)
	v409 = v382
	goto L106
L109:
	;
	v416 = v410 + int32(-1)
	goto L103
L110:
	;
	v413 = *(*int32)(unsafe.Add(mBase, uint32(v8)+4))
	v416 = v413
	goto L103
L111:
	;
	v443 = int32(2)
	v448 = *(*int32)(unsafe.Add(mBase, uint32(v423+v441<<(uint(v443)%32)+int32(4))))
	v372 = v448 + v442<<(uint(v443)%32)
	v373 = int32(1)
	goto L97
L112:
	;
	v432 = v420
	goto L114
L113:
	;
	v432 = v417 << (uint(v428) % 32)
	goto L114
L114:
	;
	if v418 < v432 {
		v441 = v424
		v442 = v418
		goto L111
	} else {
		goto L115
	}
L115:
	;
	if v424 != 0 {
		v461 = v420
		goto L100
	} else {
		goto L116
	}
L116:
	;
	v434 = *(*int32)(unsafe.Add(mBase, uint32(v423)+20))
	if v434 == int32(-1) {
		v461 = v420
		goto L100
	} else {
		goto L117
	}
L117:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v8)+4)) = int64(4294967296)
	v441 = int32(1)
	v442 = int32(0)
	goto L111
L118:
	;
	v457 = *(*int32)(unsafe.Add(mBase, uint32(v453)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v365))) = v457
	v461 = v453
	goto L100
L119:
	;
	v469 = v461
	goto L120
L120:
	;
	v472 = *(*int32)(unsafe.Add(mBase, uint32(v469)+8))
	goto L123
L121:
	;
	goto L92
L122:
	;
	v495 = v8 + int32(20)
	v496 = *(*int32)(unsafe.Add(mBase, uint32(v8)+16))
	if v496 != 0 {
		goto L128
	} else {
		goto L129
	}
L123:
	;
	v473 = *(*int32)(unsafe.Add(mBase, uint32(v472)+4))
	if v473 == int32(0) {
		goto L122
	} else {
		goto L124
	}
L124:
	;
	v476 = *(*int32)(unsafe.Add(mBase, uint32(l0)+140))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+140)) = v476 | int32(16)
	v481 = *(*int32)(unsafe.Add(mBase, uint32(v472)+16))
	m.T0[v473].(func(*base.Module, int32, int32, int32))(m, l0, int32(0), v481)
	mBase = m.M
	v483 = m.ExcPending
	if v483 != 0 {
		goto L7
	} else {
		goto L125
	}
L125:
	;
	v484 = *(*int32)(unsafe.Add(mBase, uint32(l0)+140))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+140)) = v484 & int32(-17)
	goto L122
L126:
	;
	if v591 != 0 {
		v469 = v591
		goto L120
	} else {
		goto L152
	}
L127:
	;
	v502 = v495
	v503 = v499
	goto L130
L128:
	;
	v499 = int32(1)
	goto L127
L129:
	;
	v499 = int32(0)
	goto L127
L130:
	;
	switch v503 {
	case 0:
		goto L135
	default:
		goto L134
	}
L132:
	;
	v503 = int32(0)
	goto L130
L133:
	;
	goto L126
L134:
	;
	v583 = *(*int32)(unsafe.Add(mBase, uint32(v502)))
	*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = v583
	if v583 == int32(0) {
		goto L132
	} else {
		goto L151
	}
L135:
	;
	v507 = *(*int32)(unsafe.Add(mBase, uint32(v8)+4))
	if v507 != int32(-1) {
		v546 = v507
		goto L136
	} else {
		goto L137
	}
L136:
	;
	v547 = int32(1)
	v548 = v546 + v547
	*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = v548
	v550 = int32(0)
	v553 = *(*int32)(unsafe.Add(mBase, uint32(v8)))
	v554 = *(*int32)(unsafe.Add(mBase, uint32(v8)+8))
	v558 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v553+v554+int32(26)))))
	if v558 == int32(255) {
		goto L145
	} else {
		goto L146
	}
L137:
	;
	v511 = *(*int32)(unsafe.Add(mBase, uint32(v8)+8))
	if v511 != 0 {
		v546 = int32(-1)
		goto L136
	} else {
		goto L138
	}
L138:
	;
	v512 = *(*int32)(unsafe.Add(mBase, uint32(v8)))
	v513 = *(*int32)(unsafe.Add(mBase, uint32(v8)+12))
	if v513 == int32(0) {
		goto L140
	} else {
		goto L141
	}
L139:
	;
	v540 = *(*int32)(unsafe.Add(mBase, uint32(v539)+20))
	if v540 != int32(-1) {
		goto L142
	} else {
		goto L143
	}
L140:
	;
	v520 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v512)+16)))
	v521 = int64(*(*int8)(unsafe.Add(mBase, uint32(v512)+27)))
	v522 = int64(*(*int32)(unsafe.Add(mBase, uint32(v512)+8)))
	v523 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v512)+12)))
	v524 = int64(*(*int8)(unsafe.Add(mBase, uint32(v512)+26)))
	v525 = int64(*(*int32)(unsafe.Add(mBase, uint32(v512)+4)))
	v526 = F_wangHash64(m, v525)
	mBase = m.M
	v528 = F_wangHash64(m, v524+v526)
	mBase = m.M
	v530 = F_wangHash64(m, v523+v528)
	mBase = m.M
	v532 = F_wangHash64(m, v522+v530)
	mBase = m.M
	v534 = F_wangHash64(m, v521+v532)
	mBase = m.M
	v536 = F_wangHash64(m, v520+v534)
	mBase = m.M
	*(*int64)(unsafe.Add(mBase, uint32(v8)+24)) = v536
	v538 = *(*int32)(unsafe.Add(mBase, uint32(v8)))
	v539 = v538
	goto L139
L141:
	;
	v516 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v512)+24)))
	v518 = v516 + int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v512)+24)) = uint16(v518)
	v539 = v512
	goto L139
L142:
	;
	v546 = v540 + int32(-1)
	goto L136
L143:
	;
	v543 = *(*int32)(unsafe.Add(mBase, uint32(v8)+4))
	v546 = v543
	goto L136
L144:
	;
	v573 = int32(2)
	v578 = *(*int32)(unsafe.Add(mBase, uint32(v553+v571<<(uint(v573)%32)+int32(4))))
	v502 = v578 + v572<<(uint(v573)%32)
	v503 = int32(1)
	goto L130
L145:
	;
	v562 = v550
	goto L147
L146:
	;
	v562 = v547 << (uint(v558) % 32)
	goto L147
L147:
	;
	if v548 < v562 {
		v571 = v554
		v572 = v548
		goto L144
	} else {
		goto L148
	}
L148:
	;
	if v554 != 0 {
		v591 = v550
		goto L133
	} else {
		goto L149
	}
L149:
	;
	v564 = *(*int32)(unsafe.Add(mBase, uint32(v553)+20))
	if v564 == int32(-1) {
		v591 = v550
		goto L133
	} else {
		goto L150
	}
L150:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v8)+4)) = int64(4294967296)
	v571 = int32(1)
	v572 = int32(0)
	goto L144
L151:
	;
	v587 = *(*int32)(unsafe.Add(mBase, uint32(v583)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v495))) = v587
	v591 = v583
	goto L133
L152:
	;
	goto L121
L153:
	;
	goto L89
L154:
	;
	v868 = *(*int32)(unsafe.Add(mBase, uint32(l0)+240))
	if v868 == int32(0) {
		goto L219
	} else {
		goto L220
	}
L155:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8))) = v608
	*(*int32)(unsafe.Add(mBase, uint32(v8)+20)) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v8)+12)) = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v8)+4)) = int64(4294967295)
	goto L156
L156:
	;
	v625 = v8 + int32(20)
	v626 = *(*int32)(unsafe.Add(mBase, uint32(v8)+16))
	if v626 != 0 {
		goto L160
	} else {
		goto L161
	}
L157:
	;
	v860 = *(*int32)(unsafe.Add(mBase, uint32(l0)+288))
	F_dictRelease(m, v860)
	mBase = m.M
	v862 = m.ExcPending
	if v862 != 0 {
		goto L7
	} else {
		goto L218
	}
L158:
	;
	if v721 == int32(0) {
		goto L157
	} else {
		goto L184
	}
L159:
	;
	v632 = v625
	v633 = v629
	goto L162
L160:
	;
	v629 = int32(1)
	goto L159
L161:
	;
	v629 = int32(0)
	goto L159
L162:
	;
	switch v633 {
	case 0:
		goto L167
	default:
		goto L166
	}
L164:
	;
	v633 = int32(0)
	goto L162
L165:
	;
	goto L158
L166:
	;
	v713 = *(*int32)(unsafe.Add(mBase, uint32(v632)))
	*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = v713
	if v713 == int32(0) {
		goto L164
	} else {
		goto L183
	}
L167:
	;
	v637 = *(*int32)(unsafe.Add(mBase, uint32(v8)+4))
	if v637 != int32(-1) {
		v676 = v637
		goto L168
	} else {
		goto L169
	}
L168:
	;
	v677 = int32(1)
	v678 = v676 + v677
	*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = v678
	v680 = int32(0)
	v683 = *(*int32)(unsafe.Add(mBase, uint32(v8)))
	v684 = *(*int32)(unsafe.Add(mBase, uint32(v8)+8))
	v688 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v683+v684+int32(26)))))
	if v688 == int32(255) {
		goto L177
	} else {
		goto L178
	}
L169:
	;
	v641 = *(*int32)(unsafe.Add(mBase, uint32(v8)+8))
	if v641 != 0 {
		v676 = int32(-1)
		goto L168
	} else {
		goto L170
	}
L170:
	;
	v642 = *(*int32)(unsafe.Add(mBase, uint32(v8)))
	v643 = *(*int32)(unsafe.Add(mBase, uint32(v8)+12))
	if v643 == int32(0) {
		goto L172
	} else {
		goto L173
	}
L171:
	;
	v670 = *(*int32)(unsafe.Add(mBase, uint32(v669)+20))
	if v670 != int32(-1) {
		goto L174
	} else {
		goto L175
	}
L172:
	;
	v650 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v642)+16)))
	v651 = int64(*(*int8)(unsafe.Add(mBase, uint32(v642)+27)))
	v652 = int64(*(*int32)(unsafe.Add(mBase, uint32(v642)+8)))
	v653 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v642)+12)))
	v654 = int64(*(*int8)(unsafe.Add(mBase, uint32(v642)+26)))
	v655 = int64(*(*int32)(unsafe.Add(mBase, uint32(v642)+4)))
	v656 = F_wangHash64(m, v655)
	mBase = m.M
	v658 = F_wangHash64(m, v654+v656)
	mBase = m.M
	v660 = F_wangHash64(m, v653+v658)
	mBase = m.M
	v662 = F_wangHash64(m, v652+v660)
	mBase = m.M
	v664 = F_wangHash64(m, v651+v662)
	mBase = m.M
	v666 = F_wangHash64(m, v650+v664)
	mBase = m.M
	*(*int64)(unsafe.Add(mBase, uint32(v8)+24)) = v666
	v668 = *(*int32)(unsafe.Add(mBase, uint32(v8)))
	v669 = v668
	goto L171
L173:
	;
	v646 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v642)+24)))
	v648 = v646 + int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v642)+24)) = uint16(v648)
	v669 = v642
	goto L171
L174:
	;
	v676 = v670 + int32(-1)
	goto L168
L175:
	;
	v673 = *(*int32)(unsafe.Add(mBase, uint32(v8)+4))
	v676 = v673
	goto L168
L176:
	;
	v703 = int32(2)
	v708 = *(*int32)(unsafe.Add(mBase, uint32(v683+v701<<(uint(v703)%32)+int32(4))))
	v632 = v708 + v702<<(uint(v703)%32)
	v633 = int32(1)
	goto L162
L177:
	;
	v692 = v680
	goto L179
L178:
	;
	v692 = v677 << (uint(v688) % 32)
	goto L179
L179:
	;
	if v678 < v692 {
		v701 = v684
		v702 = v678
		goto L176
	} else {
		goto L180
	}
L180:
	;
	if v684 != 0 {
		v721 = v680
		goto L165
	} else {
		goto L181
	}
L181:
	;
	v694 = *(*int32)(unsafe.Add(mBase, uint32(v683)+20))
	if v694 == int32(-1) {
		v721 = v680
		goto L165
	} else {
		goto L182
	}
L182:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v8)+4)) = int64(4294967296)
	v701 = int32(1)
	v702 = int32(0)
	goto L176
L183:
	;
	v717 = *(*int32)(unsafe.Add(mBase, uint32(v713)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v625))) = v717
	v721 = v713
	goto L165
L184:
	;
	v729 = v721
	goto L185
L185:
	;
	v732 = *(*int32)(unsafe.Add(mBase, uint32(v729)+8))
	goto L188
L186:
	;
	goto L157
L187:
	;
	v755 = v8 + int32(20)
	v756 = *(*int32)(unsafe.Add(mBase, uint32(v8)+16))
	if v756 != 0 {
		goto L193
	} else {
		goto L194
	}
L188:
	;
	v733 = *(*int32)(unsafe.Add(mBase, uint32(v732)+4))
	if v733 == int32(0) {
		goto L187
	} else {
		goto L189
	}
L189:
	;
	v736 = *(*int32)(unsafe.Add(mBase, uint32(l0)+140))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+140)) = v736 | int32(16)
	v741 = *(*int32)(unsafe.Add(mBase, uint32(v732)+16))
	m.T0[v733].(func(*base.Module, int32, int32, int32))(m, l0, int32(0), v741)
	mBase = m.M
	v743 = m.ExcPending
	if v743 != 0 {
		goto L7
	} else {
		goto L190
	}
L190:
	;
	v744 = *(*int32)(unsafe.Add(mBase, uint32(l0)+140))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+140)) = v744 & int32(-17)
	goto L187
L191:
	;
	if v851 != 0 {
		v729 = v851
		goto L185
	} else {
		goto L217
	}
L192:
	;
	v762 = v755
	v763 = v759
	goto L195
L193:
	;
	v759 = int32(1)
	goto L192
L194:
	;
	v759 = int32(0)
	goto L192
L195:
	;
	switch v763 {
	case 0:
		goto L200
	default:
		goto L199
	}
L197:
	;
	v763 = int32(0)
	goto L195
L198:
	;
	goto L191
L199:
	;
	v843 = *(*int32)(unsafe.Add(mBase, uint32(v762)))
	*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = v843
	if v843 == int32(0) {
		goto L197
	} else {
		goto L216
	}
L200:
	;
	v767 = *(*int32)(unsafe.Add(mBase, uint32(v8)+4))
	if v767 != int32(-1) {
		v806 = v767
		goto L201
	} else {
		goto L202
	}
L201:
	;
	v807 = int32(1)
	v808 = v806 + v807
	*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = v808
	v810 = int32(0)
	v813 = *(*int32)(unsafe.Add(mBase, uint32(v8)))
	v814 = *(*int32)(unsafe.Add(mBase, uint32(v8)+8))
	v818 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v813+v814+int32(26)))))
	if v818 == int32(255) {
		goto L210
	} else {
		goto L211
	}
L202:
	;
	v771 = *(*int32)(unsafe.Add(mBase, uint32(v8)+8))
	if v771 != 0 {
		v806 = int32(-1)
		goto L201
	} else {
		goto L203
	}
L203:
	;
	v772 = *(*int32)(unsafe.Add(mBase, uint32(v8)))
	v773 = *(*int32)(unsafe.Add(mBase, uint32(v8)+12))
	if v773 == int32(0) {
		goto L205
	} else {
		goto L206
	}
L204:
	;
	v800 = *(*int32)(unsafe.Add(mBase, uint32(v799)+20))
	if v800 != int32(-1) {
		goto L207
	} else {
		goto L208
	}
L205:
	;
	v780 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v772)+16)))
	v781 = int64(*(*int8)(unsafe.Add(mBase, uint32(v772)+27)))
	v782 = int64(*(*int32)(unsafe.Add(mBase, uint32(v772)+8)))
	v783 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v772)+12)))
	v784 = int64(*(*int8)(unsafe.Add(mBase, uint32(v772)+26)))
	v785 = int64(*(*int32)(unsafe.Add(mBase, uint32(v772)+4)))
	v786 = F_wangHash64(m, v785)
	mBase = m.M
	v788 = F_wangHash64(m, v784+v786)
	mBase = m.M
	v790 = F_wangHash64(m, v783+v788)
	mBase = m.M
	v792 = F_wangHash64(m, v782+v790)
	mBase = m.M
	v794 = F_wangHash64(m, v781+v792)
	mBase = m.M
	v796 = F_wangHash64(m, v780+v794)
	mBase = m.M
	*(*int64)(unsafe.Add(mBase, uint32(v8)+24)) = v796
	v798 = *(*int32)(unsafe.Add(mBase, uint32(v8)))
	v799 = v798
	goto L204
L206:
	;
	v776 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v772)+24)))
	v778 = v776 + int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v772)+24)) = uint16(v778)
	v799 = v772
	goto L204
L207:
	;
	v806 = v800 + int32(-1)
	goto L201
L208:
	;
	v803 = *(*int32)(unsafe.Add(mBase, uint32(v8)+4))
	v806 = v803
	goto L201
L209:
	;
	v833 = int32(2)
	v838 = *(*int32)(unsafe.Add(mBase, uint32(v813+v831<<(uint(v833)%32)+int32(4))))
	v762 = v838 + v832<<(uint(v833)%32)
	v763 = int32(1)
	goto L195
L210:
	;
	v822 = v810
	goto L212
L211:
	;
	v822 = v807 << (uint(v818) % 32)
	goto L212
L212:
	;
	if v808 < v822 {
		v831 = v814
		v832 = v808
		goto L209
	} else {
		goto L213
	}
L213:
	;
	if v814 != 0 {
		v851 = v810
		goto L198
	} else {
		goto L214
	}
L214:
	;
	v824 = *(*int32)(unsafe.Add(mBase, uint32(v813)+20))
	if v824 == int32(-1) {
		v851 = v810
		goto L198
	} else {
		goto L215
	}
L215:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v8)+4)) = int64(4294967296)
	v831 = int32(1)
	v832 = int32(0)
	goto L209
L216:
	;
	v847 = *(*int32)(unsafe.Add(mBase, uint32(v843)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v755))) = v847
	v851 = v843
	goto L198
L217:
	;
	goto L186
L218:
	;
	goto L154
L219:
	;
	v874 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+240)) = v874
	v876 = *(*int32)(unsafe.Add(mBase, uint32(l0)+140))
	if v876&int32(2) == v874 {
		goto L222
	} else {
		goto L223
	}
L220:
	;
	v871 = *(*int32)(unsafe.Add(mBase, uint32(l0)+220))
	m.T0[v868].(func(*base.Module, int32))(m, v871)
	mBase = m.M
	v873 = m.ExcPending
	if v873 != 0 {
		goto L7
	} else {
		goto L221
	}
L221:
	;
	goto L219
L222:
	;
	v909 = *(*int32)(unsafe.Add(mBase, uint32(l0)+216))
	if v909 == int32(0) {
		goto L229
	} else {
		goto L230
	}
L223:
	;
	v881 = *(*int32)(unsafe.Add(mBase, uint32(l0)+248))
	if v881 == int32(0) {
		goto L222
	} else {
		goto L224
	}
L224:
	;
	v884 = int32(0)
	v889 = *(*int32)(unsafe.Add(mBase, uint32(l0)+204))
	v893 = v884 - base.B2i32(v876&int32(8) == v884)&base.B2i32(v889 != v884)
	if v876&int32(16) != 0 {
		goto L225
	} else {
		goto L226
	}
L225:
	;
	m.T0[v881].(func(*base.Module, int32, int32))(m, l0, v893)
	mBase = m.M
	v906 = m.ExcPending
	if v906 != 0 {
		goto L7
	} else {
		goto L228
	}
L226:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+140)) = v876 | int32(16)
	m.T0[v881].(func(*base.Module, int32, int32))(m, l0, v893)
	mBase = m.M
	v900 = m.ExcPending
	if v900 != 0 {
		goto L7
	} else {
		goto L227
	}
L227:
	;
	v901 = *(*int32)(unsafe.Add(mBase, uint32(l0)+140))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+140)) = v901 & int32(-17)
	goto L222
L228:
	;
	goto L222
L229:
	;
	F_valkeyFree(m, l0)
	mBase = m.M
	v916 = m.ExcPending
	if v916 != 0 {
		goto L7
	} else {
		goto L232
	}
L230:
	;
	v912 = *(*int32)(unsafe.Add(mBase, uint32(l0)+212))
	m.T0[v909].(func(*base.Module, int32))(m, v912)
	mBase = m.M
	v914 = m.ExcPending
	if v914 != 0 {
		goto L7
	} else {
		goto L231
	}
L231:
	;
	goto L229
L232:
	;
	m.G0 = v8 + int32(32)
	return
}
func F_valkeyAsyncSetConnectCallback(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v21 int32
	_ = v21
	var v24 int64
	_ = v24
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v33 int64
	_ = v33
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v45 int64
	_ = v45
	var v47 int64
	_ = v47
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+252))
	if v11 != 0 {
		v63 = int32(-1)
		m.G0 = v8 + int32(16)
		return v63
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(l0)+252)) = l1
		v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+244))
		v14 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+140)))
		if v14&int32(2) == int32(0) {
			if v13 == int32(0) {
				v55 = int32(0)
				v56 = *(*int32)(unsafe.Add(mBase, uint32(l0)+232))
				if v56 == v55 {
					v63 = v55
					m.G0 = v8 + int32(16)
					return v63
				} else {
					v59 = *(*int32)(unsafe.Add(mBase, uint32(l0)+220))
					m.T0[v56].(func(*base.Module, int32))(m, v59)
					mBase = m.M
					v61 = m.ExcPending
					if v61 != 0 {
						return int32(0)
					} else {
						v63 = v55
						m.G0 = v8 + int32(16)
						return v63
					}
				}
			} else {
				v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)+156))
				if v30 == int32(0) {
					v55 = int32(0)
					v56 = *(*int32)(unsafe.Add(mBase, uint32(l0)+232))
					if v56 == v55 {
						v63 = v55
						m.G0 = v8 + int32(16)
						return v63
					} else {
						v59 = *(*int32)(unsafe.Add(mBase, uint32(l0)+220))
						m.T0[v56].(func(*base.Module, int32))(m, v59)
						mBase = m.M
						v61 = m.ExcPending
						if v61 != 0 {
							return int32(0)
						} else {
							v63 = v55
							m.G0 = v8 + int32(16)
							return v63
						}
					}
				} else {
					v33 = *(*int64)(unsafe.Add(mBase, uint32(v30)))
					if v33 != int64(0) {
						v39 = v30
						v40 = *(*int32)(unsafe.Add(mBase, uint32(l0)+220))
						v41 = int32(8)
						v45 = *(*int64)(unsafe.Add(mBase, uint32(v39+v41)))
						*(*int64)(unsafe.Add(mBase, uint32(v8+v41))) = v45
						v47 = *(*int64)(unsafe.Add(mBase, uint32(v39)))
						*(*int64)(unsafe.Add(mBase, uint32(v8))) = v47
						m.T0[v13].(func(*base.Module, int32, int32))(m, v40, v8)
						mBase = m.M
						v52 = m.ExcPending
						if v52 != 0 {
							return int32(0)
						} else {
							v55 = int32(0)
							v56 = *(*int32)(unsafe.Add(mBase, uint32(l0)+232))
							if v56 == v55 {
								v63 = v55
								m.G0 = v8 + int32(16)
								return v63
							} else {
								v59 = *(*int32)(unsafe.Add(mBase, uint32(l0)+220))
								m.T0[v56].(func(*base.Module, int32))(m, v59)
								mBase = m.M
								v61 = m.ExcPending
								if v61 != 0 {
									return int32(0)
								} else {
									v63 = v55
									m.G0 = v8 + int32(16)
									return v63
								}
							}
						}
					} else {
						v36 = *(*int32)(unsafe.Add(mBase, uint32(v30)+8))
						if v36 == int32(0) {
							v55 = int32(0)
							v56 = *(*int32)(unsafe.Add(mBase, uint32(l0)+232))
							if v56 == v55 {
								v63 = v55
								m.G0 = v8 + int32(16)
								return v63
							} else {
								v59 = *(*int32)(unsafe.Add(mBase, uint32(l0)+220))
								m.T0[v56].(func(*base.Module, int32))(m, v59)
								mBase = m.M
								v61 = m.ExcPending
								if v61 != 0 {
									return int32(0)
								} else {
									v63 = v55
									m.G0 = v8 + int32(16)
									return v63
								}
							}
						} else {
							v39 = v30
							v40 = *(*int32)(unsafe.Add(mBase, uint32(l0)+220))
							v41 = int32(8)
							v45 = *(*int64)(unsafe.Add(mBase, uint32(v39+v41)))
							*(*int64)(unsafe.Add(mBase, uint32(v8+v41))) = v45
							v47 = *(*int64)(unsafe.Add(mBase, uint32(v39)))
							*(*int64)(unsafe.Add(mBase, uint32(v8))) = v47
							m.T0[v13].(func(*base.Module, int32, int32))(m, v40, v8)
							mBase = m.M
							v52 = m.ExcPending
							if v52 != 0 {
								return int32(0)
							} else {
								v55 = int32(0)
								v56 = *(*int32)(unsafe.Add(mBase, uint32(l0)+232))
								if v56 == v55 {
									v63 = v55
									m.G0 = v8 + int32(16)
									return v63
								} else {
									v59 = *(*int32)(unsafe.Add(mBase, uint32(l0)+220))
									m.T0[v56].(func(*base.Module, int32))(m, v59)
									mBase = m.M
									v61 = m.ExcPending
									if v61 != 0 {
										return int32(0)
									} else {
										v63 = v55
										m.G0 = v8 + int32(16)
										return v63
									}
								}
							}
						}
					}
				}
			}
		} else {
			if v13 == int32(0) {
				v55 = int32(0)
				v56 = *(*int32)(unsafe.Add(mBase, uint32(l0)+232))
				if v56 == v55 {
					v63 = v55
					m.G0 = v8 + int32(16)
					return v63
				} else {
					v59 = *(*int32)(unsafe.Add(mBase, uint32(l0)+220))
					m.T0[v56].(func(*base.Module, int32))(m, v59)
					mBase = m.M
					v61 = m.ExcPending
					if v61 != 0 {
						return int32(0)
					} else {
						v63 = v55
						m.G0 = v8 + int32(16)
						return v63
					}
				}
			} else {
				v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+160))
				if v21 == int32(0) {
					v55 = int32(0)
					v56 = *(*int32)(unsafe.Add(mBase, uint32(l0)+232))
					if v56 == v55 {
						v63 = v55
						m.G0 = v8 + int32(16)
						return v63
					} else {
						v59 = *(*int32)(unsafe.Add(mBase, uint32(l0)+220))
						m.T0[v56].(func(*base.Module, int32))(m, v59)
						mBase = m.M
						v61 = m.ExcPending
						if v61 != 0 {
							return int32(0)
						} else {
							v63 = v55
							m.G0 = v8 + int32(16)
							return v63
						}
					}
				} else {
					v24 = *(*int64)(unsafe.Add(mBase, uint32(v21)))
					if v24 != int64(0) {
						v39 = v21
						v40 = *(*int32)(unsafe.Add(mBase, uint32(l0)+220))
						v41 = int32(8)
						v45 = *(*int64)(unsafe.Add(mBase, uint32(v39+v41)))
						*(*int64)(unsafe.Add(mBase, uint32(v8+v41))) = v45
						v47 = *(*int64)(unsafe.Add(mBase, uint32(v39)))
						*(*int64)(unsafe.Add(mBase, uint32(v8))) = v47
						m.T0[v13].(func(*base.Module, int32, int32))(m, v40, v8)
						mBase = m.M
						v52 = m.ExcPending
						if v52 != 0 {
							return int32(0)
						} else {
							v55 = int32(0)
							v56 = *(*int32)(unsafe.Add(mBase, uint32(l0)+232))
							if v56 == v55 {
								v63 = v55
								m.G0 = v8 + int32(16)
								return v63
							} else {
								v59 = *(*int32)(unsafe.Add(mBase, uint32(l0)+220))
								m.T0[v56].(func(*base.Module, int32))(m, v59)
								mBase = m.M
								v61 = m.ExcPending
								if v61 != 0 {
									return int32(0)
								} else {
									v63 = v55
									m.G0 = v8 + int32(16)
									return v63
								}
							}
						}
					} else {
						v27 = *(*int32)(unsafe.Add(mBase, uint32(v21)+8))
						if v27 != 0 {
							v39 = v21
							v40 = *(*int32)(unsafe.Add(mBase, uint32(l0)+220))
							v41 = int32(8)
							v45 = *(*int64)(unsafe.Add(mBase, uint32(v39+v41)))
							*(*int64)(unsafe.Add(mBase, uint32(v8+v41))) = v45
							v47 = *(*int64)(unsafe.Add(mBase, uint32(v39)))
							*(*int64)(unsafe.Add(mBase, uint32(v8))) = v47
							m.T0[v13].(func(*base.Module, int32, int32))(m, v40, v8)
							mBase = m.M
							v52 = m.ExcPending
							if v52 != 0 {
								return int32(0)
							} else {
								v55 = int32(0)
								v56 = *(*int32)(unsafe.Add(mBase, uint32(l0)+232))
								if v56 == v55 {
									v63 = v55
									m.G0 = v8 + int32(16)
									return v63
								} else {
									v59 = *(*int32)(unsafe.Add(mBase, uint32(l0)+220))
									m.T0[v56].(func(*base.Module, int32))(m, v59)
									mBase = m.M
									v61 = m.ExcPending
									if v61 != 0 {
										return int32(0)
									} else {
										v63 = v55
										m.G0 = v8 + int32(16)
										return v63
									}
								}
							}
						} else {
							v55 = int32(0)
							v56 = *(*int32)(unsafe.Add(mBase, uint32(l0)+232))
							if v56 == v55 {
								v63 = v55
								m.G0 = v8 + int32(16)
								return v63
							} else {
								v59 = *(*int32)(unsafe.Add(mBase, uint32(l0)+220))
								m.T0[v56].(func(*base.Module, int32))(m, v59)
								mBase = m.M
								v61 = m.ExcPending
								if v61 != 0 {
									return int32(0)
								} else {
									v63 = v55
									m.G0 = v8 + int32(16)
									return v63
								}
							}
						}
					}
				}
			}
		}
	}
}
func F_valkeyAsyncSetDisconnectCallback(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+248))
	if v5 != 0 {
		v8 = int32(-1)
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(l0)+248)) = l1
		v8 = int32(0)
	}
	return v8
}
func F_valkeyBufferWrite(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
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
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
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
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v80 int32
	_ = v80
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v96 int32
	_ = v96
	var v99 int32
	_ = v99
	var v102 int32
	_ = v102
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v112 int32
	_ = v112
	var v116 int32
	_ = v116
	var v119 int32
	_ = v119
	var v124 int32
	_ = v124
	var v127 int32
	_ = v127
	var v130 int32
	_ = v130
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v141 int32
	_ = v141
	var v143 int32
	_ = v143
	var v145 int32
	_ = v145
	var v149 int32
	_ = v149
	var v151 int32
	_ = v151
	var v154 int32
	_ = v154
	var v188 int32
	_ = v188
	var v191 int32
	_ = v191
	var v198 int32
	_ = v198
	var v201 int32
	_ = v201
	var v204 int32
	_ = v204
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v218 int32
	_ = v218
	var v220 int32
	_ = v220
	var v223 int64
	_ = v223
	var v229 int64
	_ = v229
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v7 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(-1)
L2:
	;
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+144))
	v11 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8+int32(-1)))))
	switch v11 & int32(7) {
	case 0:
		goto L10
	case 1:
		goto L9
	case 2:
		goto L8
	case 3:
		goto L7
	case 4:
		goto L6
	default:
		goto L4
	}
L3:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = int32(5)
	v218 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+21)) = uint8(v218)
	v220 = m.G3
	v223 = *(*int64)(unsafe.Add(mBase, uint32(v220)+uint32(_c_F_valkeyBufferWrite[0])))
	*(*int64)(unsafe.Add(mBase, uint32(l0)+8)) = v223
	v229 = *(*int64)(unsafe.Add(mBase, uint32(v220)+uint32(_c_F_valkeyBufferWrite[1])))
	*(*int64)(unsafe.Add(mBase, uint32(l0+int32(13)))) = v229
	goto L1
L4:
	;
	if l1 != 0 {
		goto L66
	} else {
		goto L67
	}
L5:
	;
	if v28 == int32(0) {
		goto L4
	} else {
		goto L11
	}
L6:
	;
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v8+int32(-17))))
	v28 = v27
	goto L5
L7:
	;
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v8+int32(-9))))
	v28 = v24
	goto L5
L8:
	;
	v21 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v8+int32(-5)))))
	v28 = v21
	goto L5
L9:
	;
	v18 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8+int32(-3)))))
	v28 = v18
	goto L5
L10:
	;
	v28 = int32(base.Ui32(v11) >> (uint(int32(3)) % 32))
	goto L5
L11:
	;
	v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v31)+32))
	v33 = m.T0[v32].(func(*base.Module, int32) int32)(m, l0)
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	return int32(0)
L13:
	;
	if v33 < int32(0) {
		goto L1
	} else {
		goto L14
	}
L14:
	;
	if v33 == int32(0) {
		goto L4
	} else {
		goto L15
	}
L15:
	;
	v41 = *(*int32)(unsafe.Add(mBase, uint32(l0)+144))
	v44 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v41+int32(-1)))))
	v46 = v44 & int32(7)
	switch v46 {
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
		goto L16
	}
L16:
	;
	v80 = int32(-1)
	v88 = v41 + v80
	v89 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v88))))
	v91 = v89 & int32(7)
	switch v91 {
	case 0:
		goto L39
	case 1:
		goto L38
	case 2:
		goto L37
	case 3:
		goto L36
	case 4:
		goto L35
	default:
		goto L33
	}
L17:
	;
	if v33 != v61 {
		goto L23
	} else {
		goto L24
	}
L18:
	;
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v41+int32(-17))))
	v61 = v60
	goto L17
L19:
	;
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v41+int32(-9))))
	v61 = v57
	goto L17
L20:
	;
	v54 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v41+int32(-5)))))
	v61 = v54
	goto L17
L21:
	;
	v51 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v41+int32(-3)))))
	v61 = v51
	goto L17
L22:
	;
	v61 = int32(base.Ui32(v44) >> (uint(int32(3)) % 32))
	goto L17
L23:
	;
	switch v46 + int32(-3) {
	case 0:
		goto L30
	case 1:
		goto L29
	default:
		goto L16
	}
L24:
	;
	F_sdsfree(m, v41)
	mBase = m.M
	v64 = m.ExcPending
	if v64 != 0 {
		goto L12
	} else {
		goto L25
	}
L25:
	;
	v65 = F_sdsempty(m)
	mBase = m.M
	v66 = m.ExcPending
	if v66 != 0 {
		goto L12
	} else {
		goto L26
	}
L26:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+144)) = v65
	if v65 != 0 {
		goto L4
	} else {
		goto L27
	}
L27:
	;
	goto L3
L28:
	;
	if v76 < int32(0) {
		goto L3
	} else {
		goto L31
	}
L29:
	;
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v41+int32(-17))))
	v76 = v75
	goto L28
L30:
	;
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v41+int32(-9))))
	v76 = v72
	goto L28
L31:
	;
	goto L16
L32:
	;
	goto L4
L33:
	;
	goto L32
L34:
	;
	if v106 == int32(0) {
		goto L33
	} else {
		goto L40
	}
L35:
	;
	v105 = *(*int32)(unsafe.Add(mBase, uint32(v41+int32(-17))))
	v106 = v105
	goto L34
L36:
	;
	v102 = *(*int32)(unsafe.Add(mBase, uint32(v41+int32(-9))))
	v106 = v102
	goto L34
L37:
	;
	v99 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v41+int32(-5)))))
	v106 = v99
	goto L34
L38:
	;
	v96 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v41+int32(-3)))))
	v106 = v96
	goto L34
L39:
	;
	v106 = int32(base.Ui32(v89) >> (uint(int32(3)) % 32))
	goto L34
L40:
	;
	v112 = int32(-1)&v106 + v80
	v116 = v33>>(uint(int32(31))%32)&v106 + v33
	v119 = v112 - v116 + int32(1)
	switch v91 {
	default:
		goto L46
	case 1:
		goto L45
	case 2:
		goto L44
	case 3:
		goto L43
	case 4:
		goto L42
	}
L41:
	;
	v135 = int32(0)
	v137 = base.B2i32(base.Ui32(v116) < base.Ui32(v134))
	if base.Ui32(v116) < base.Ui32(v134) {
		goto L48
	} else {
		goto L49
	}
L42:
	;
	v133 = *(*int32)(unsafe.Add(mBase, uint32(v41+int32(-17))))
	v134 = v133
	goto L41
L43:
	;
	v130 = *(*int32)(unsafe.Add(mBase, uint32(v41+int32(-9))))
	v134 = v130
	goto L41
L44:
	;
	v127 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v41+int32(-5)))))
	v134 = v127
	goto L41
L45:
	;
	v124 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v41+int32(-3)))))
	v134 = v124
	goto L41
L46:
	;
	v134 = int32(base.Ui32(v89) >> (uint(int32(3)) % 32))
	goto L41
L47:
	;
	v151 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v41+v145))) = uint8(v151)
	switch v91 {
	default:
		goto L65
	case 1:
		goto L64
	case 2:
		goto L63
	case 3:
		goto L62
	case 4:
		goto L61
	}
L48:
	;
	v138 = v116
	goto L50
L49:
	;
	v138 = v135
	goto L50
L50:
	;
	v139 = v134 - v138
	if base.Ui32(v119) < base.Ui32(v139) {
		goto L51
	} else {
		goto L52
	}
L51:
	;
	v141 = v119
	goto L53
L52:
	;
	v141 = v139
	goto L53
L53:
	;
	if v112 < v116 {
		goto L54
	} else {
		goto L55
	}
L54:
	;
	v143 = v135
	goto L56
L55:
	;
	v143 = v141
	goto L56
L56:
	;
	if base.Ui32(v116) < base.Ui32(v134) {
		goto L57
	} else {
		goto L58
	}
L57:
	;
	v145 = v143
	goto L59
L58:
	;
	v145 = int32(0)
	goto L59
L59:
	;
	if v145 == int32(0) {
		goto L47
	} else {
		goto L60
	}
L60:
	;
	v149 = F_memmove(m, v41, v41+v138, v145)
	mBase = m.M
	goto L47
L61:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v41+int32(-17)))) = base.I64_extend_i32_u(v145)
	goto L33
L62:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v41+int32(-9)))) = v145
	goto L32
L63:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v41+int32(-5)))) = uint16(v145)
	goto L32
L64:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v41+int32(-3)))) = uint8(v145)
	goto L32
L65:
	;
	v154 = v145 << (uint(int32(3)) % 32)
	*(*uint8)(unsafe.Add(mBase, uint32(v88))) = uint8(v154)
	goto L32
L66:
	;
	v188 = *(*int32)(unsafe.Add(mBase, uint32(l0)+144))
	v191 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v188+int32(-1)))))
	switch v191 & int32(7) {
	case 0:
		goto L73
	case 1:
		goto L72
	case 2:
		goto L71
	case 3:
		goto L70
	case 4:
		goto L69
	default:
		v208 = int32(0)
		goto L68
	}
L67:
	;
	return int32(0)
L68:
	;
	v209 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = base.B2i32(v208 == v209)
	return v209
L69:
	;
	v207 = *(*int32)(unsafe.Add(mBase, uint32(v188+int32(-17))))
	v208 = v207
	goto L68
L70:
	;
	v204 = *(*int32)(unsafe.Add(mBase, uint32(v188+int32(-9))))
	v208 = v204
	goto L68
L71:
	;
	v201 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v188+int32(-5)))))
	v208 = v201
	goto L68
L72:
	;
	v198 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v188+int32(-3)))))
	v208 = v198
	goto L68
L73:
	;
	v208 = int32(base.Ui32(v191) >> (uint(int32(3)) % 32))
	goto L68
}
func F_valkeyContextConnectUnix(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
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
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v47 int32
	_ = v47
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v77 int32
	_ = v77
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v89 int32
	_ = v89
	var v90 int64
	_ = v90
	var v92 int32
	_ = v92
	var v96 int64
	_ = v96
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v102 int32
	_ = v102
	var v105 int64
	_ = v105
	var v109 int32
	_ = v109
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v119 int32
	_ = v119
	var v122 int32
	_ = v122
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v129 int32
	_ = v129
	var v134 int32
	_ = v134
	var v137 int32
	_ = v137
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v143 int32
	_ = v143
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v164 int32
	_ = v164
	var v168 int32
	_ = v168
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v186 int32
	_ = v186
	var v196 int32
	_ = v196
	var v200 int32
	_ = v200
	v10 = m.G0
	v12 = v10 - int32(128)
	m.G0 = v12
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+140))
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v17 = int32(-1)
	v18 = int32(1)
	v21 = F_socket(m, v18, v18, int32(0))
	mBase = m.M
	if v21 != v17 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v12 + int32(128)
	return v200
L2:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+136)) = v21
	v57 = F_valkeySetBlocking(m, l0, int32(0))
	mBase = m.M
	v60 = m.ExcPending
	if v60 != 0 {
		goto L11
	} else {
		goto L12
	}
L3:
	;
	goto L4
L4:
	;
	v25 = *(*int32)(unsafe.Add(mBase, _c_F_valkeyContextConnectUnix[0]))
	v29 = F__emscripten_memset_bulkmem(m, v12, base.I32_extend8_s(int32(0)), int32(128))
	mBase = m.M
	goto L5
L5:
	;
	v32 = F_strerror(m, v25)
	mBase = m.M
	v33 = F_strlen(m, v32)
	mBase = m.M
	if base.Ui32(v33) < base.Ui32(int32(128)) {
		goto L8
	} else {
		goto L9
	}
L6:
	;
	F_valkeySetError(m, l0, int32(1), v29)
	mBase = m.M
	v200 = v17
	goto L1
L7:
	;
	goto L6
L8:
	;
	v47 = F___memcpy(m, v29, v32, v33+int32(1))
	mBase = m.M
	goto L7
L9:
	;
	goto L10
L10:
	;
	v39 = int32(127)
	v40 = F___memcpy(m, v29, v32, v39)
	mBase = m.M
	v42 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v29+v39))) = uint8(v42)
	goto L6
L11:
	;
	return int32(0)
L12:
	;
	if v57 != 0 {
		v200 = v17
		goto L1
	} else {
		goto L13
	}
L13:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+152)) = int32(1)
	v63 = *(*int32)(unsafe.Add(mBase, uint32(l0)+176))
	if v63 == v15 {
		goto L15
	} else {
		goto L16
	}
L14:
	;
	v196 = m.G3
	F_valkeySetError(m, l0, int32(5), v196+int32(_a_F_valkeyContextConnectUnix_0))
	mBase = m.M
	v200 = v17
	goto L1
L15:
	;
	v77 = *(*int32)(unsafe.Add(mBase, uint32(l0)+156))
	if v16 == int32(0) {
		goto L22
	} else {
		goto L23
	}
L16:
	;
	v65 = m.G4
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v65)+16))
	m.T0[v66].(func(*base.Module, int32))(m, v63)
	mBase = m.M
	v68 = m.ExcPending
	if v68 != 0 {
		goto L11
	} else {
		goto L17
	}
L17:
	;
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v65)+12))
	v70 = m.T0[v69].(func(*base.Module, int32) int32)(m, v15)
	mBase = m.M
	v71 = m.ExcPending
	if v71 != 0 {
		goto L11
	} else {
		goto L18
	}
L18:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+176)) = v70
	if v70 == int32(0) {
		goto L14
	} else {
		goto L19
	}
L19:
	;
	goto L15
L20:
	;
	v137 = *(*int32)(unsafe.Add(mBase, uint32(l0)+180))
	if v137 == int32(0) {
		goto L38
	} else {
		goto L39
	}
L21:
	;
	v134 = int32(2147483647)
	goto L20
L22:
	;
	v126 = m.G4
	v127 = *(*int32)(unsafe.Add(mBase, uint32(v126)+16))
	m.T0[v127].(func(*base.Module, int32))(m, v77)
	mBase = m.M
	v129 = m.ExcPending
	if v129 != 0 {
		goto L11
	} else {
		goto L37
	}
L23:
	;
	if v77 == v16 {
		v99 = v77
		goto L24
	} else {
		goto L25
	}
L24:
	;
	if v99 == int32(0) {
		goto L21
	} else {
		goto L30
	}
L25:
	;
	if v77 != 0 {
		v89 = v77
		goto L26
	} else {
		goto L27
	}
L26:
	;
	v90 = *(*int64)(unsafe.Add(mBase, uint32(v16)))
	*(*int64)(unsafe.Add(mBase, uint32(v89))) = v90
	v92 = int32(8)
	v96 = *(*int64)(unsafe.Add(mBase, uint32(v16+v92)))
	*(*int64)(unsafe.Add(mBase, uint32(v89+v92))) = v96
	v98 = *(*int32)(unsafe.Add(mBase, uint32(l0)+156))
	v99 = v98
	goto L24
L27:
	;
	v82 = m.G4
	v83 = *(*int32)(unsafe.Add(mBase, uint32(v82)))
	v84 = m.T0[v83].(func(*base.Module, int32) int32)(m, int32(16))
	mBase = m.M
	v85 = m.ExcPending
	if v85 != 0 {
		goto L11
	} else {
		goto L28
	}
L28:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+156)) = v84
	if v84 == int32(0) {
		goto L14
	} else {
		goto L29
	}
L29:
	;
	v89 = v84
	goto L26
L30:
	;
	v102 = *(*int32)(unsafe.Add(mBase, uint32(v99)+8))
	if int32(1000000) < v102 {
		goto L31
	} else {
		goto L32
	}
L31:
	;
	v122 = m.G3
	F_valkeySetError(m, l0, int32(1), v122+int32(_a_F_valkeyContextConnectUnix_1))
	mBase = m.M
	v200 = v17
	goto L1
L32:
	;
	v105 = *(*int64)(unsafe.Add(mBase, uint32(v99)))
	if int64(2147482) < v105 {
		goto L31
	} else {
		goto L33
	}
L33:
	;
	v109 = int32(1000)
	v114 = base.I32_div_s(v102+int32(999), v109)
	v115 = base.I32_wrap_i64(v105)*v109 + v114
	v116 = int32(2147483647)
	if base.Ui32(v115) < base.Ui32(v116) {
		goto L34
	} else {
		goto L35
	}
L34:
	;
	v119 = v115
	goto L36
L35:
	;
	v119 = v116
	goto L36
L36:
	;
	v134 = v119
	goto L20
L37:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+156)) = int32(0)
	goto L21
L38:
	;
	v145 = m.G4
	v146 = *(*int32)(unsafe.Add(mBase, uint32(v145)))
	v147 = m.T0[v146].(func(*base.Module, int32) int32)(m, int32(110))
	mBase = m.M
	v148 = m.ExcPending
	if v148 != 0 {
		goto L11
	} else {
		goto L41
	}
L39:
	;
	v140 = m.G4
	v141 = *(*int32)(unsafe.Add(mBase, uint32(v140)+16))
	m.T0[v141].(func(*base.Module, int32))(m, v137)
	mBase = m.M
	v143 = m.ExcPending
	if v143 != 0 {
		goto L11
	} else {
		goto L40
	}
L40:
	;
	goto L38
L41:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+180)) = v147
	if v147 == int32(0) {
		goto L14
	} else {
		goto L42
	}
L42:
	;
	v152 = int32(1)
	v153 = v14 & v152
	*(*int32)(unsafe.Add(mBase, uint32(l0)+184)) = int32(110)
	*(*uint16)(unsafe.Add(mBase, uint32(v147))) = uint16(v152)
	v161 = F___stpncpy(m, v147+int32(2), v15, int32(107))
	mBase = m.M
	goto L43
L43:
	;
	v162 = *(*int32)(unsafe.Add(mBase, uint32(l0)+136))
	v164 = F_connect(m, v162, v147, int32(110))
	mBase = m.M
	if v164 != int32(-1) {
		goto L45
	} else {
		goto L46
	}
L44:
	;
	v186 = *(*int32)(unsafe.Add(mBase, uint32(l0)+140))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+140)) = v186 | int32(2)
	v200 = int32(0)
	goto L1
L45:
	;
	if v153 == int32(0) {
		goto L44
	} else {
		goto L53
	}
L46:
	;
	goto L48
L47:
	;
	v173 = F_valkeyContextWaitReady(m, l0, v134)
	mBase = m.M
	v174 = m.ExcPending
	if v174 != 0 {
		goto L11
	} else {
		goto L51
	}
L48:
	;
	v168 = *(*int32)(unsafe.Add(mBase, _c_F_valkeyContextConnectUnix[0]))
	if v168 != int32(26) {
		goto L47
	} else {
		goto L49
	}
L49:
	;
	if v153 == int32(0) {
		goto L44
	} else {
		goto L50
	}
L50:
	;
	goto L47
L51:
	;
	if v173 == int32(0) {
		goto L45
	} else {
		goto L52
	}
L52:
	;
	v200 = int32(-1)
	goto L1
L53:
	;
	v181 = F_valkeySetBlocking(m, l0, int32(1))
	mBase = m.M
	v182 = m.ExcPending
	if v182 != 0 {
		goto L11
	} else {
		goto L54
	}
L54:
	;
	if v181 == int32(0) {
		goto L44
	} else {
		goto L55
	}
L55:
	;
	v200 = int32(-1)
	goto L1
}
func F_valkeyContextRegisterTcpFuncs(m *base.Module) {
	var v1 int32
	_ = v1
	var v5 int32
	_ = v5
	v1 = m.G3
	v5 = F_valkeyContextRegisterFuncs(m, v1+int32(_a_F_valkeyContextRegisterTcpFuncs_0), int32(0))
	return
}
func F_valkeyContextSetFuncs(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v37 int32
	_ = v37
	var v41 int32
	_ = v41
	var v50 int32
	_ = v50
	var v59 int32
	_ = v59
	v3 = m.G3
	v6 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3)+uint32(_c_F_valkeyContextSetFuncs[0]))))
	if v6 != 0 {
	} else {
		v7 = m.G3
		v10 = int32(1)
		*(*uint8)(unsafe.Add(mBase, uint32(v7)+uint32(_c_F_valkeyContextSetFuncs[0]))) = uint8(v10)
		v12 = m.G3
		v16 = F_valkeyContextRegisterFuncs(m, v12+int32(_a_F_valkeyContextSetFuncs_0), int32(0))
		mBase = m.M
		v17 = m.G3
		v21 = F_valkeyContextRegisterFuncs(m, v17+int32(_a_F_valkeyContextSetFuncs_1), int32(1))
		mBase = m.M
		v22 = m.G3
		v26 = F_valkeyContextRegisterFuncs(m, v22+int32(_a_F_valkeyContextSetFuncs_2), int32(2))
		mBase = m.M
	}
	v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)+152))
	if base.Ui32(int32(4)) <= base.Ui32(v27) {
		v41 = m.G3
		m.Env.X__assert_fail(m, v41+int32(_a_F_valkeyContextSetFuncs_3), v41+int32(_a_F_valkeyContextSetFuncs_4), int32(104), v41+int32(_a_F_valkeyContextSetFuncs_5))
		mBase = m.M
		base.Wasm_trap_unreachable()
		for {
		}
	} else {
		v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		if v30 != 0 {
			v50 = m.G3
			m.Env.X__assert_fail(m, v50+int32(_a_F_valkeyContextSetFuncs_6), v50+int32(_a_F_valkeyContextSetFuncs_4), int32(105), v50+int32(_a_F_valkeyContextSetFuncs_5))
			mBase = m.M
			base.Wasm_trap_unreachable()
			for {
			}
		} else {
			v31 = m.G3
			v37 = *(*int32)(unsafe.Add(mBase, uint32(v31+int32(_a_F_valkeyContextSetFuncs_7)+v27<<(uint(int32(2))%32))))
			*(*int32)(unsafe.Add(mBase, uint32(l0))) = v37
			if v37 == int32(0) {
				v59 = m.G3
				m.Env.X__assert_fail(m, v59+int32(_a_F_valkeyContextSetFuncs_8), v59+int32(_a_F_valkeyContextSetFuncs_4), int32(107), v59+int32(_a_F_valkeyContextSetFuncs_5))
				mBase = m.M
				base.Wasm_trap_unreachable()
				for {
				}
			} else {
				return
			}
		}
	}
}
func F_valkeyFree(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
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
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
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
	var v49 int32
	_ = v49
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	if l0 == int32(0) {
		return
	} else {
		v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		if v6 == int32(0) {
			v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+144))
			F_sdsfree(m, v15)
			mBase = m.M
			v17 = m.ExcPending
			if v17 != 0 {
				return
			} else {
				v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+148))
				F_valkeyReaderFree(m, v18)
				mBase = m.M
				v20 = m.ExcPending
				if v20 != 0 {
					return
				} else {
					v21 = m.G4
					v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+164))
					v23 = *(*int32)(unsafe.Add(mBase, uint32(v21)+16))
					m.T0[v23].(func(*base.Module, int32))(m, v22)
					mBase = m.M
					v25 = m.ExcPending
					if v25 != 0 {
						return
					} else {
						v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+168))
						v27 = *(*int32)(unsafe.Add(mBase, uint32(v21)+16))
						m.T0[v27].(func(*base.Module, int32))(m, v26)
						mBase = m.M
						v29 = m.ExcPending
						if v29 != 0 {
							return
						} else {
							v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)+176))
							v31 = *(*int32)(unsafe.Add(mBase, uint32(v21)+16))
							m.T0[v31].(func(*base.Module, int32))(m, v30)
							mBase = m.M
							v33 = m.ExcPending
							if v33 != 0 {
								return
							} else {
								v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)+156))
								v35 = *(*int32)(unsafe.Add(mBase, uint32(v21)+16))
								m.T0[v35].(func(*base.Module, int32))(m, v34)
								mBase = m.M
								v37 = m.ExcPending
								if v37 != 0 {
									return
								} else {
									v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)+160))
									v39 = *(*int32)(unsafe.Add(mBase, uint32(v21)+16))
									m.T0[v39].(func(*base.Module, int32))(m, v38)
									mBase = m.M
									v41 = m.ExcPending
									if v41 != 0 {
										return
									} else {
										v42 = *(*int32)(unsafe.Add(mBase, uint32(l0)+180))
										v43 = *(*int32)(unsafe.Add(mBase, uint32(v21)+16))
										m.T0[v43].(func(*base.Module, int32))(m, v42)
										mBase = m.M
										v45 = m.ExcPending
										if v45 != 0 {
											return
										} else {
											v46 = *(*int32)(unsafe.Add(mBase, uint32(l0)+188))
											if v46 == int32(0) {
												v55 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
												if v55 == int32(0) {
													v68 = F__emscripten_memset_bulkmem(m, l0, base.I32_extend8_s(int32(255)), int32(204))
													mBase = m.M
													v69 = m.G4
													v70 = *(*int32)(unsafe.Add(mBase, uint32(v69)+16))
													m.T0[v70].(func(*base.Module, int32))(m, v68)
													mBase = m.M
													v72 = m.ExcPending
													if v72 != 0 {
														return
													} else {
														return
													}
												} else {
													v58 = *(*int32)(unsafe.Add(mBase, uint32(v55)+8))
													if v58 == int32(0) {
														v68 = F__emscripten_memset_bulkmem(m, l0, base.I32_extend8_s(int32(255)), int32(204))
														mBase = m.M
														v69 = m.G4
														v70 = *(*int32)(unsafe.Add(mBase, uint32(v69)+16))
														m.T0[v70].(func(*base.Module, int32))(m, v68)
														mBase = m.M
														v72 = m.ExcPending
														if v72 != 0 {
															return
														} else {
															return
														}
													} else {
														v61 = *(*int32)(unsafe.Add(mBase, uint32(l0)+196))
														m.T0[v58].(func(*base.Module, int32))(m, v61)
														mBase = m.M
														v63 = m.ExcPending
														if v63 != 0 {
															return
														} else {
															v68 = F__emscripten_memset_bulkmem(m, l0, base.I32_extend8_s(int32(255)), int32(204))
															mBase = m.M
															v69 = m.G4
															v70 = *(*int32)(unsafe.Add(mBase, uint32(v69)+16))
															m.T0[v70].(func(*base.Module, int32))(m, v68)
															mBase = m.M
															v72 = m.ExcPending
															if v72 != 0 {
																return
															} else {
																return
															}
														}
													}
												}
											} else {
												v49 = *(*int32)(unsafe.Add(mBase, uint32(l0)+192))
												if v49 == int32(0) {
													v55 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
													if v55 == int32(0) {
														v68 = F__emscripten_memset_bulkmem(m, l0, base.I32_extend8_s(int32(255)), int32(204))
														mBase = m.M
														v69 = m.G4
														v70 = *(*int32)(unsafe.Add(mBase, uint32(v69)+16))
														m.T0[v70].(func(*base.Module, int32))(m, v68)
														mBase = m.M
														v72 = m.ExcPending
														if v72 != 0 {
															return
														} else {
															return
														}
													} else {
														v58 = *(*int32)(unsafe.Add(mBase, uint32(v55)+8))
														if v58 == int32(0) {
															v68 = F__emscripten_memset_bulkmem(m, l0, base.I32_extend8_s(int32(255)), int32(204))
															mBase = m.M
															v69 = m.G4
															v70 = *(*int32)(unsafe.Add(mBase, uint32(v69)+16))
															m.T0[v70].(func(*base.Module, int32))(m, v68)
															mBase = m.M
															v72 = m.ExcPending
															if v72 != 0 {
																return
															} else {
																return
															}
														} else {
															v61 = *(*int32)(unsafe.Add(mBase, uint32(l0)+196))
															m.T0[v58].(func(*base.Module, int32))(m, v61)
															mBase = m.M
															v63 = m.ExcPending
															if v63 != 0 {
																return
															} else {
																v68 = F__emscripten_memset_bulkmem(m, l0, base.I32_extend8_s(int32(255)), int32(204))
																mBase = m.M
																v69 = m.G4
																v70 = *(*int32)(unsafe.Add(mBase, uint32(v69)+16))
																m.T0[v70].(func(*base.Module, int32))(m, v68)
																mBase = m.M
																v72 = m.ExcPending
																if v72 != 0 {
																	return
																} else {
																	return
																}
															}
														}
													}
												} else {
													m.T0[v49].(func(*base.Module, int32))(m, v46)
													mBase = m.M
													v53 = m.ExcPending
													if v53 != 0 {
														return
													} else {
														v55 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
														if v55 == int32(0) {
															v68 = F__emscripten_memset_bulkmem(m, l0, base.I32_extend8_s(int32(255)), int32(204))
															mBase = m.M
															v69 = m.G4
															v70 = *(*int32)(unsafe.Add(mBase, uint32(v69)+16))
															m.T0[v70].(func(*base.Module, int32))(m, v68)
															mBase = m.M
															v72 = m.ExcPending
															if v72 != 0 {
																return
															} else {
																return
															}
														} else {
															v58 = *(*int32)(unsafe.Add(mBase, uint32(v55)+8))
															if v58 == int32(0) {
																v68 = F__emscripten_memset_bulkmem(m, l0, base.I32_extend8_s(int32(255)), int32(204))
																mBase = m.M
																v69 = m.G4
																v70 = *(*int32)(unsafe.Add(mBase, uint32(v69)+16))
																m.T0[v70].(func(*base.Module, int32))(m, v68)
																mBase = m.M
																v72 = m.ExcPending
																if v72 != 0 {
																	return
																} else {
																	return
																}
															} else {
																v61 = *(*int32)(unsafe.Add(mBase, uint32(l0)+196))
																m.T0[v58].(func(*base.Module, int32))(m, v61)
																mBase = m.M
																v63 = m.ExcPending
																if v63 != 0 {
																	return
																} else {
																	v68 = F__emscripten_memset_bulkmem(m, l0, base.I32_extend8_s(int32(255)), int32(204))
																	mBase = m.M
																	v69 = m.G4
																	v70 = *(*int32)(unsafe.Add(mBase, uint32(v69)+16))
																	m.T0[v70].(func(*base.Module, int32))(m, v68)
																	mBase = m.M
																	v72 = m.ExcPending
																	if v72 != 0 {
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
									}
								}
							}
						}
					}
				}
			}
		} else {
			v9 = *(*int32)(unsafe.Add(mBase, uint32(v6)+4))
			if v9 == int32(0) {
				v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+144))
				F_sdsfree(m, v15)
				mBase = m.M
				v17 = m.ExcPending
				if v17 != 0 {
					return
				} else {
					v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+148))
					F_valkeyReaderFree(m, v18)
					mBase = m.M
					v20 = m.ExcPending
					if v20 != 0 {
						return
					} else {
						v21 = m.G4
						v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+164))
						v23 = *(*int32)(unsafe.Add(mBase, uint32(v21)+16))
						m.T0[v23].(func(*base.Module, int32))(m, v22)
						mBase = m.M
						v25 = m.ExcPending
						if v25 != 0 {
							return
						} else {
							v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+168))
							v27 = *(*int32)(unsafe.Add(mBase, uint32(v21)+16))
							m.T0[v27].(func(*base.Module, int32))(m, v26)
							mBase = m.M
							v29 = m.ExcPending
							if v29 != 0 {
								return
							} else {
								v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)+176))
								v31 = *(*int32)(unsafe.Add(mBase, uint32(v21)+16))
								m.T0[v31].(func(*base.Module, int32))(m, v30)
								mBase = m.M
								v33 = m.ExcPending
								if v33 != 0 {
									return
								} else {
									v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)+156))
									v35 = *(*int32)(unsafe.Add(mBase, uint32(v21)+16))
									m.T0[v35].(func(*base.Module, int32))(m, v34)
									mBase = m.M
									v37 = m.ExcPending
									if v37 != 0 {
										return
									} else {
										v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)+160))
										v39 = *(*int32)(unsafe.Add(mBase, uint32(v21)+16))
										m.T0[v39].(func(*base.Module, int32))(m, v38)
										mBase = m.M
										v41 = m.ExcPending
										if v41 != 0 {
											return
										} else {
											v42 = *(*int32)(unsafe.Add(mBase, uint32(l0)+180))
											v43 = *(*int32)(unsafe.Add(mBase, uint32(v21)+16))
											m.T0[v43].(func(*base.Module, int32))(m, v42)
											mBase = m.M
											v45 = m.ExcPending
											if v45 != 0 {
												return
											} else {
												v46 = *(*int32)(unsafe.Add(mBase, uint32(l0)+188))
												if v46 == int32(0) {
													v55 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
													if v55 == int32(0) {
														v68 = F__emscripten_memset_bulkmem(m, l0, base.I32_extend8_s(int32(255)), int32(204))
														mBase = m.M
														v69 = m.G4
														v70 = *(*int32)(unsafe.Add(mBase, uint32(v69)+16))
														m.T0[v70].(func(*base.Module, int32))(m, v68)
														mBase = m.M
														v72 = m.ExcPending
														if v72 != 0 {
															return
														} else {
															return
														}
													} else {
														v58 = *(*int32)(unsafe.Add(mBase, uint32(v55)+8))
														if v58 == int32(0) {
															v68 = F__emscripten_memset_bulkmem(m, l0, base.I32_extend8_s(int32(255)), int32(204))
															mBase = m.M
															v69 = m.G4
															v70 = *(*int32)(unsafe.Add(mBase, uint32(v69)+16))
															m.T0[v70].(func(*base.Module, int32))(m, v68)
															mBase = m.M
															v72 = m.ExcPending
															if v72 != 0 {
																return
															} else {
																return
															}
														} else {
															v61 = *(*int32)(unsafe.Add(mBase, uint32(l0)+196))
															m.T0[v58].(func(*base.Module, int32))(m, v61)
															mBase = m.M
															v63 = m.ExcPending
															if v63 != 0 {
																return
															} else {
																v68 = F__emscripten_memset_bulkmem(m, l0, base.I32_extend8_s(int32(255)), int32(204))
																mBase = m.M
																v69 = m.G4
																v70 = *(*int32)(unsafe.Add(mBase, uint32(v69)+16))
																m.T0[v70].(func(*base.Module, int32))(m, v68)
																mBase = m.M
																v72 = m.ExcPending
																if v72 != 0 {
																	return
																} else {
																	return
																}
															}
														}
													}
												} else {
													v49 = *(*int32)(unsafe.Add(mBase, uint32(l0)+192))
													if v49 == int32(0) {
														v55 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
														if v55 == int32(0) {
															v68 = F__emscripten_memset_bulkmem(m, l0, base.I32_extend8_s(int32(255)), int32(204))
															mBase = m.M
															v69 = m.G4
															v70 = *(*int32)(unsafe.Add(mBase, uint32(v69)+16))
															m.T0[v70].(func(*base.Module, int32))(m, v68)
															mBase = m.M
															v72 = m.ExcPending
															if v72 != 0 {
																return
															} else {
																return
															}
														} else {
															v58 = *(*int32)(unsafe.Add(mBase, uint32(v55)+8))
															if v58 == int32(0) {
																v68 = F__emscripten_memset_bulkmem(m, l0, base.I32_extend8_s(int32(255)), int32(204))
																mBase = m.M
																v69 = m.G4
																v70 = *(*int32)(unsafe.Add(mBase, uint32(v69)+16))
																m.T0[v70].(func(*base.Module, int32))(m, v68)
																mBase = m.M
																v72 = m.ExcPending
																if v72 != 0 {
																	return
																} else {
																	return
																}
															} else {
																v61 = *(*int32)(unsafe.Add(mBase, uint32(l0)+196))
																m.T0[v58].(func(*base.Module, int32))(m, v61)
																mBase = m.M
																v63 = m.ExcPending
																if v63 != 0 {
																	return
																} else {
																	v68 = F__emscripten_memset_bulkmem(m, l0, base.I32_extend8_s(int32(255)), int32(204))
																	mBase = m.M
																	v69 = m.G4
																	v70 = *(*int32)(unsafe.Add(mBase, uint32(v69)+16))
																	m.T0[v70].(func(*base.Module, int32))(m, v68)
																	mBase = m.M
																	v72 = m.ExcPending
																	if v72 != 0 {
																		return
																	} else {
																		return
																	}
																}
															}
														}
													} else {
														m.T0[v49].(func(*base.Module, int32))(m, v46)
														mBase = m.M
														v53 = m.ExcPending
														if v53 != 0 {
															return
														} else {
															v55 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
															if v55 == int32(0) {
																v68 = F__emscripten_memset_bulkmem(m, l0, base.I32_extend8_s(int32(255)), int32(204))
																mBase = m.M
																v69 = m.G4
																v70 = *(*int32)(unsafe.Add(mBase, uint32(v69)+16))
																m.T0[v70].(func(*base.Module, int32))(m, v68)
																mBase = m.M
																v72 = m.ExcPending
																if v72 != 0 {
																	return
																} else {
																	return
																}
															} else {
																v58 = *(*int32)(unsafe.Add(mBase, uint32(v55)+8))
																if v58 == int32(0) {
																	v68 = F__emscripten_memset_bulkmem(m, l0, base.I32_extend8_s(int32(255)), int32(204))
																	mBase = m.M
																	v69 = m.G4
																	v70 = *(*int32)(unsafe.Add(mBase, uint32(v69)+16))
																	m.T0[v70].(func(*base.Module, int32))(m, v68)
																	mBase = m.M
																	v72 = m.ExcPending
																	if v72 != 0 {
																		return
																	} else {
																		return
																	}
																} else {
																	v61 = *(*int32)(unsafe.Add(mBase, uint32(l0)+196))
																	m.T0[v58].(func(*base.Module, int32))(m, v61)
																	mBase = m.M
																	v63 = m.ExcPending
																	if v63 != 0 {
																		return
																	} else {
																		v68 = F__emscripten_memset_bulkmem(m, l0, base.I32_extend8_s(int32(255)), int32(204))
																		mBase = m.M
																		v69 = m.G4
																		v70 = *(*int32)(unsafe.Add(mBase, uint32(v69)+16))
																		m.T0[v70].(func(*base.Module, int32))(m, v68)
																		mBase = m.M
																		v72 = m.ExcPending
																		if v72 != 0 {
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
										}
									}
								}
							}
						}
					}
				}
			} else {
				m.T0[v9].(func(*base.Module, int32))(m, l0)
				mBase = m.M
				v13 = m.ExcPending
				if v13 != 0 {
					return
				} else {
					v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+144))
					F_sdsfree(m, v15)
					mBase = m.M
					v17 = m.ExcPending
					if v17 != 0 {
						return
					} else {
						v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+148))
						F_valkeyReaderFree(m, v18)
						mBase = m.M
						v20 = m.ExcPending
						if v20 != 0 {
							return
						} else {
							v21 = m.G4
							v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+164))
							v23 = *(*int32)(unsafe.Add(mBase, uint32(v21)+16))
							m.T0[v23].(func(*base.Module, int32))(m, v22)
							mBase = m.M
							v25 = m.ExcPending
							if v25 != 0 {
								return
							} else {
								v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+168))
								v27 = *(*int32)(unsafe.Add(mBase, uint32(v21)+16))
								m.T0[v27].(func(*base.Module, int32))(m, v26)
								mBase = m.M
								v29 = m.ExcPending
								if v29 != 0 {
									return
								} else {
									v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)+176))
									v31 = *(*int32)(unsafe.Add(mBase, uint32(v21)+16))
									m.T0[v31].(func(*base.Module, int32))(m, v30)
									mBase = m.M
									v33 = m.ExcPending
									if v33 != 0 {
										return
									} else {
										v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)+156))
										v35 = *(*int32)(unsafe.Add(mBase, uint32(v21)+16))
										m.T0[v35].(func(*base.Module, int32))(m, v34)
										mBase = m.M
										v37 = m.ExcPending
										if v37 != 0 {
											return
										} else {
											v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)+160))
											v39 = *(*int32)(unsafe.Add(mBase, uint32(v21)+16))
											m.T0[v39].(func(*base.Module, int32))(m, v38)
											mBase = m.M
											v41 = m.ExcPending
											if v41 != 0 {
												return
											} else {
												v42 = *(*int32)(unsafe.Add(mBase, uint32(l0)+180))
												v43 = *(*int32)(unsafe.Add(mBase, uint32(v21)+16))
												m.T0[v43].(func(*base.Module, int32))(m, v42)
												mBase = m.M
												v45 = m.ExcPending
												if v45 != 0 {
													return
												} else {
													v46 = *(*int32)(unsafe.Add(mBase, uint32(l0)+188))
													if v46 == int32(0) {
														v55 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
														if v55 == int32(0) {
															v68 = F__emscripten_memset_bulkmem(m, l0, base.I32_extend8_s(int32(255)), int32(204))
															mBase = m.M
															v69 = m.G4
															v70 = *(*int32)(unsafe.Add(mBase, uint32(v69)+16))
															m.T0[v70].(func(*base.Module, int32))(m, v68)
															mBase = m.M
															v72 = m.ExcPending
															if v72 != 0 {
																return
															} else {
																return
															}
														} else {
															v58 = *(*int32)(unsafe.Add(mBase, uint32(v55)+8))
															if v58 == int32(0) {
																v68 = F__emscripten_memset_bulkmem(m, l0, base.I32_extend8_s(int32(255)), int32(204))
																mBase = m.M
																v69 = m.G4
																v70 = *(*int32)(unsafe.Add(mBase, uint32(v69)+16))
																m.T0[v70].(func(*base.Module, int32))(m, v68)
																mBase = m.M
																v72 = m.ExcPending
																if v72 != 0 {
																	return
																} else {
																	return
																}
															} else {
																v61 = *(*int32)(unsafe.Add(mBase, uint32(l0)+196))
																m.T0[v58].(func(*base.Module, int32))(m, v61)
																mBase = m.M
																v63 = m.ExcPending
																if v63 != 0 {
																	return
																} else {
																	v68 = F__emscripten_memset_bulkmem(m, l0, base.I32_extend8_s(int32(255)), int32(204))
																	mBase = m.M
																	v69 = m.G4
																	v70 = *(*int32)(unsafe.Add(mBase, uint32(v69)+16))
																	m.T0[v70].(func(*base.Module, int32))(m, v68)
																	mBase = m.M
																	v72 = m.ExcPending
																	if v72 != 0 {
																		return
																	} else {
																		return
																	}
																}
															}
														}
													} else {
														v49 = *(*int32)(unsafe.Add(mBase, uint32(l0)+192))
														if v49 == int32(0) {
															v55 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
															if v55 == int32(0) {
																v68 = F__emscripten_memset_bulkmem(m, l0, base.I32_extend8_s(int32(255)), int32(204))
																mBase = m.M
																v69 = m.G4
																v70 = *(*int32)(unsafe.Add(mBase, uint32(v69)+16))
																m.T0[v70].(func(*base.Module, int32))(m, v68)
																mBase = m.M
																v72 = m.ExcPending
																if v72 != 0 {
																	return
																} else {
																	return
																}
															} else {
																v58 = *(*int32)(unsafe.Add(mBase, uint32(v55)+8))
																if v58 == int32(0) {
																	v68 = F__emscripten_memset_bulkmem(m, l0, base.I32_extend8_s(int32(255)), int32(204))
																	mBase = m.M
																	v69 = m.G4
																	v70 = *(*int32)(unsafe.Add(mBase, uint32(v69)+16))
																	m.T0[v70].(func(*base.Module, int32))(m, v68)
																	mBase = m.M
																	v72 = m.ExcPending
																	if v72 != 0 {
																		return
																	} else {
																		return
																	}
																} else {
																	v61 = *(*int32)(unsafe.Add(mBase, uint32(l0)+196))
																	m.T0[v58].(func(*base.Module, int32))(m, v61)
																	mBase = m.M
																	v63 = m.ExcPending
																	if v63 != 0 {
																		return
																	} else {
																		v68 = F__emscripten_memset_bulkmem(m, l0, base.I32_extend8_s(int32(255)), int32(204))
																		mBase = m.M
																		v69 = m.G4
																		v70 = *(*int32)(unsafe.Add(mBase, uint32(v69)+16))
																		m.T0[v70].(func(*base.Module, int32))(m, v68)
																		mBase = m.M
																		v72 = m.ExcPending
																		if v72 != 0 {
																			return
																		} else {
																			return
																		}
																	}
																}
															}
														} else {
															m.T0[v49].(func(*base.Module, int32))(m, v46)
															mBase = m.M
															v53 = m.ExcPending
															if v53 != 0 {
																return
															} else {
																v55 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
																if v55 == int32(0) {
																	v68 = F__emscripten_memset_bulkmem(m, l0, base.I32_extend8_s(int32(255)), int32(204))
																	mBase = m.M
																	v69 = m.G4
																	v70 = *(*int32)(unsafe.Add(mBase, uint32(v69)+16))
																	m.T0[v70].(func(*base.Module, int32))(m, v68)
																	mBase = m.M
																	v72 = m.ExcPending
																	if v72 != 0 {
																		return
																	} else {
																		return
																	}
																} else {
																	v58 = *(*int32)(unsafe.Add(mBase, uint32(v55)+8))
																	if v58 == int32(0) {
																		v68 = F__emscripten_memset_bulkmem(m, l0, base.I32_extend8_s(int32(255)), int32(204))
																		mBase = m.M
																		v69 = m.G4
																		v70 = *(*int32)(unsafe.Add(mBase, uint32(v69)+16))
																		m.T0[v70].(func(*base.Module, int32))(m, v68)
																		mBase = m.M
																		v72 = m.ExcPending
																		if v72 != 0 {
																			return
																		} else {
																			return
																		}
																	} else {
																		v61 = *(*int32)(unsafe.Add(mBase, uint32(l0)+196))
																		m.T0[v58].(func(*base.Module, int32))(m, v61)
																		mBase = m.M
																		v63 = m.ExcPending
																		if v63 != 0 {
																			return
																		} else {
																			v68 = F__emscripten_memset_bulkmem(m, l0, base.I32_extend8_s(int32(255)), int32(204))
																			mBase = m.M
																			v69 = m.G4
																			v70 = *(*int32)(unsafe.Add(mBase, uint32(v69)+16))
																			m.T0[v70].(func(*base.Module, int32))(m, v68)
																			mBase = m.M
																			v72 = m.ExcPending
																			if v72 != 0 {
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
}
func F_valkeyGetReply(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v77 int32
	_ = v77
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v88 int32
	_ = v88
	var v94 int32
	_ = v94
	var v97 int32
	_ = v97
	var v103 int32
	_ = v103
	var v107 int32
	_ = v107
	var v109 int32
	_ = v109
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v121 int32
	_ = v121
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v127 int32
	_ = v127
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v148 int32
	_ = v148
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v162 int32
	_ = v162
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v175 int32
	_ = v175
	var v178 int32
	_ = v178
	var v181 int32
	_ = v181
	var v184 int32
	_ = v184
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v204 int32
	_ = v204
	var v206 int32
	_ = v206
	var v214 int32
	_ = v214
	var v217 int32
	_ = v217
	var v221 int32
	_ = v221
	var v226 int32
	_ = v226
	var v228 int32
	_ = v228
	var v232 int32
	_ = v232
	var v238 int32
	_ = v238
	var v241 int32
	_ = v241
	var v247 int32
	_ = v247
	var v251 int32
	_ = v251
	var v253 int32
	_ = v253
	var v261 int32
	_ = v261
	var v262 int32
	_ = v262
	var v265 int32
	_ = v265
	var v268 int32
	_ = v268
	var v269 int32
	_ = v269
	var v271 int32
	_ = v271
	var v275 int32
	_ = v275
	var v276 int32
	_ = v276
	var v283 int32
	_ = v283
	var v289 int32
	_ = v289
	var v294 int32
	_ = v294
	v3 = int32(0)
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	*(*int32)(unsafe.Add(mBase, uint32(v9)+12)) = v3
	*(*int32)(unsafe.Add(mBase, uint32(v9)+8)) = v3
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+148))
	v19 = F_valkeyReaderGetReply(m, v16, v9+int32(8))
	mBase = m.M
	v22 = m.ExcPending
	if v22 != 0 {
		goto L5
	} else {
		goto L6
	}
L1:
	;
	m.G0 = v9 + int32(16)
	return v294
L2:
	;
	if l1 == int32(0) {
		goto L85
	} else {
		goto L86
	}
L3:
	;
	v129 = int32(0)
	v130 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+140)))
	if v130&int32(1) == v129 {
		v283 = v129
		goto L2
	} else {
		goto L38
	}
L4:
	;
	v56 = *(*int32)(unsafe.Add(mBase, uint32(l0)+148))
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v56)))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v57
	v60 = v56 + int32(4)
	v62 = l0 + int32(8)
	if v60&int32(3) == int32(0) {
		v84 = v60
		goto L18
	} else {
		goto L19
	}
L5:
	;
	return int32(0)
L6:
	;
	if v19 == int32(-1) {
		goto L4
	} else {
		goto L7
	}
L7:
	;
	goto L8
L8:
	;
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v9)+8))
	if v31 == int32(0) {
		goto L3
	} else {
		goto L10
	}
L9:
	;
	goto L4
L10:
	;
	v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)+200))
	if v34 == int32(0) {
		v283 = v31
		goto L2
	} else {
		goto L11
	}
L11:
	;
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v31)))
	if v37 != int32(12) {
		v283 = v31
		goto L2
	} else {
		goto L12
	}
L12:
	;
	v40 = *(*int32)(unsafe.Add(mBase, uint32(l0)+188))
	m.T0[v34].(func(*base.Module, int32, int32))(m, v40, v31)
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L5
	} else {
		goto L13
	}
L13:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(l0)+148))
	v46 = F_valkeyReaderGetReply(m, v43, v9+int32(8))
	mBase = m.M
	v47 = m.ExcPending
	if v47 != 0 {
		goto L5
	} else {
		goto L14
	}
L14:
	;
	if v46 != int32(-1) {
		goto L8
	} else {
		goto L15
	}
L15:
	;
	goto L9
L16:
	;
	v118 = int32(127)
	if base.Ui32(v117) < base.Ui32(v118) {
		goto L32
	} else {
		goto L33
	}
L17:
	;
	v117 = v109 - v60
	goto L16
L18:
	;
	v88 = v84
	goto L26
L19:
	;
	v70 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v60))))
	if v70 != 0 {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	v73 = v60
	goto L22
L21:
	;
	v117 = v60 - v60
	goto L16
L22:
	;
	v77 = v73 + int32(1)
	if v77&int32(3) == int32(0) {
		v84 = v77
		goto L18
	} else {
		goto L24
	}
L24:
	;
	v82 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v77))))
	if v82 != 0 {
		v73 = v77
		goto L22
	} else {
		goto L25
	}
L25:
	;
	v109 = v77
	goto L17
L26:
	;
	v94 = *(*int32)(unsafe.Add(mBase, uint32(v88)))
	v97 = int32(-2139062144)
	if (int32(16843008)-v94|v94)&v97 == v97 {
		v88 = v88 + int32(4)
		goto L26
	} else {
		goto L28
	}
L27:
	;
	v103 = v88
	goto L29
L28:
	;
	goto L27
L29:
	;
	v107 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v103))))
	if v107 != 0 {
		v103 = v103 + int32(1)
		goto L29
	} else {
		goto L31
	}
L30:
	;
	v109 = v103
	goto L17
L31:
	;
	goto L30
L32:
	;
	v121 = v117
	goto L34
L33:
	;
	v121 = v118
	goto L34
L34:
	;
	if v121 == int32(0) {
		v125 = v62
		goto L36
	} else {
		goto L37
	}
L35:
	;
	v127 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v125+v121))) = uint8(v127)
	v294 = int32(-1)
	goto L1
L36:
	;
	goto L35
L37:
	;
	v124 = F__emscripten_memcpy_bulkmem(m, v62, v60, v121)
	mBase = m.M
	v125 = v124
	goto L36
L38:
	;
	goto L39
L39:
	;
	v144 = F_valkeyBufferWrite(m, l0, v9+int32(12))
	mBase = m.M
	v145 = m.ExcPending
	if v145 != 0 {
		goto L5
	} else {
		goto L41
	}
L40:
	;
	v152 = F_valkeyBufferRead(m, l0)
	mBase = m.M
	v153 = m.ExcPending
	if v153 != 0 {
		goto L5
	} else {
		goto L44
	}
L41:
	;
	if v144 == int32(-1) {
		v294 = int32(-1)
		goto L1
	} else {
		goto L42
	}
L42:
	;
	v148 = *(*int32)(unsafe.Add(mBase, uint32(v9)+12))
	if v148 == int32(0) {
		goto L39
	} else {
		goto L43
	}
L43:
	;
	goto L40
L44:
	;
	if v152 == int32(-1) {
		v294 = int32(-1)
		goto L1
	} else {
		goto L45
	}
L45:
	;
	goto L46
L46:
	;
	v162 = *(*int32)(unsafe.Add(mBase, uint32(l0)+148))
	v165 = F_valkeyReaderGetReply(m, v162, v9+int32(8))
	mBase = m.M
	v166 = m.ExcPending
	if v166 != 0 {
		goto L5
	} else {
		goto L50
	}
L48:
	;
	v275 = F_valkeyBufferRead(m, l0)
	mBase = m.M
	v276 = m.ExcPending
	if v276 != 0 {
		goto L5
	} else {
		goto L82
	}
L49:
	;
	v200 = *(*int32)(unsafe.Add(mBase, uint32(l0)+148))
	v201 = *(*int32)(unsafe.Add(mBase, uint32(v200)))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v201
	v204 = v200 + int32(4)
	v206 = l0 + int32(8)
	if v204&int32(3) == int32(0) {
		v228 = v204
		goto L62
	} else {
		goto L63
	}
L50:
	;
	if v165 == int32(-1) {
		goto L49
	} else {
		goto L51
	}
L51:
	;
	goto L52
L52:
	;
	v175 = *(*int32)(unsafe.Add(mBase, uint32(v9)+8))
	if v175 == int32(0) {
		goto L48
	} else {
		goto L54
	}
L53:
	;
	goto L49
L54:
	;
	v178 = *(*int32)(unsafe.Add(mBase, uint32(l0)+200))
	if v178 == int32(0) {
		v283 = v175
		goto L2
	} else {
		goto L55
	}
L55:
	;
	v181 = *(*int32)(unsafe.Add(mBase, uint32(v175)))
	if v181 != int32(12) {
		v283 = v175
		goto L2
	} else {
		goto L56
	}
L56:
	;
	v184 = *(*int32)(unsafe.Add(mBase, uint32(l0)+188))
	m.T0[v178].(func(*base.Module, int32, int32))(m, v184, v175)
	mBase = m.M
	v186 = m.ExcPending
	if v186 != 0 {
		goto L5
	} else {
		goto L57
	}
L57:
	;
	v187 = *(*int32)(unsafe.Add(mBase, uint32(l0)+148))
	v190 = F_valkeyReaderGetReply(m, v187, v9+int32(8))
	mBase = m.M
	v191 = m.ExcPending
	if v191 != 0 {
		goto L5
	} else {
		goto L58
	}
L58:
	;
	if v190 != int32(-1) {
		goto L52
	} else {
		goto L59
	}
L59:
	;
	goto L53
L60:
	;
	v262 = int32(127)
	if base.Ui32(v261) < base.Ui32(v262) {
		goto L76
	} else {
		goto L77
	}
L61:
	;
	v261 = v253 - v204
	goto L60
L62:
	;
	v232 = v228
	goto L70
L63:
	;
	v214 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v204))))
	if v214 != 0 {
		goto L64
	} else {
		goto L65
	}
L64:
	;
	v217 = v204
	goto L66
L65:
	;
	v261 = v204 - v204
	goto L60
L66:
	;
	v221 = v217 + int32(1)
	if v221&int32(3) == int32(0) {
		v228 = v221
		goto L62
	} else {
		goto L68
	}
L68:
	;
	v226 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v221))))
	if v226 != 0 {
		v217 = v221
		goto L66
	} else {
		goto L69
	}
L69:
	;
	v253 = v221
	goto L61
L70:
	;
	v238 = *(*int32)(unsafe.Add(mBase, uint32(v232)))
	v241 = int32(-2139062144)
	if (int32(16843008)-v238|v238)&v241 == v241 {
		v232 = v232 + int32(4)
		goto L70
	} else {
		goto L72
	}
L71:
	;
	v247 = v232
	goto L73
L72:
	;
	goto L71
L73:
	;
	v251 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v247))))
	if v251 != 0 {
		v247 = v247 + int32(1)
		goto L73
	} else {
		goto L75
	}
L74:
	;
	v253 = v247
	goto L61
L75:
	;
	goto L74
L76:
	;
	v265 = v261
	goto L78
L77:
	;
	v265 = v262
	goto L78
L78:
	;
	if v265 == int32(0) {
		v269 = v206
		goto L80
	} else {
		goto L81
	}
L79:
	;
	v271 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v269+v265))) = uint8(v271)
	v294 = int32(-1)
	goto L1
L80:
	;
	goto L79
L81:
	;
	v268 = F__emscripten_memcpy_bulkmem(m, v206, v204, v265)
	mBase = m.M
	v269 = v268
	goto L80
L82:
	;
	if v275 != int32(-1) {
		goto L46
	} else {
		goto L83
	}
L83:
	;
	v294 = int32(-1)
	goto L1
L84:
	;
	v294 = int32(0)
	goto L1
L85:
	;
	F_freeReplyObject(m, v283)
	mBase = m.M
	v289 = m.ExcPending
	if v289 != 0 {
		goto L5
	} else {
		goto L87
	}
L86:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v283
	goto L84
L87:
	;
	goto L84
}
func F_valkeyNetWrite(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
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
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v42 int32
	_ = v42
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+136))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+144))
	v11 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8+int32(-1)))))
	switch v11 & int32(7) {
	case 0:
		v28 = int32(base.Ui32(v11) >> (uint(int32(3)) % 32))
	case 1:
		v18 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8+int32(-3)))))
		v28 = v18
	case 2:
		v21 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v8+int32(-5)))))
		v28 = v21
	case 3:
		v24 = *(*int32)(unsafe.Add(mBase, uint32(v8+int32(-9))))
		v28 = v24
	case 4:
		v27 = *(*int32)(unsafe.Add(mBase, uint32(v8+int32(-17))))
		v28 = v27
	default:
		v28 = int32(0)
	}
	v29 = int32(0)
	v32 = F_sendto(m, v6, v8, v28, v29, v29, v29)
	mBase = m.M
	if int32(-1) < v32 {
		v51 = v32
	} else {
		v35 = int32(0)
		v37 = *(*int32)(unsafe.Add(mBase, _c_F_valkeyNetWrite[0]))
		if v37 == int32(27) {
			v51 = v35
		} else {
			if v37 != int32(6) {
				v48 = F___strerror_l(m, v37, v37)
				mBase = m.M
				F_valkeySetError(m, l0, int32(1), v48)
				mBase = m.M
				v51 = int32(-1)
			} else {
				v42 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+140)))
				if v42&int32(1) == int32(0) {
					v51 = v35
				} else {
					v48 = F___strerror_l(m, v37, v37)
					mBase = m.M
					F_valkeySetError(m, l0, int32(1), v48)
					mBase = m.M
					v51 = int32(-1)
				}
			}
		}
	}
	return v51
}
func F_valkeyProcessCallbacks(m *base.Module, l0 int32) {
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
	var v28 int32
	_ = v28
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v71 int32
	_ = v71
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	var v97 int32
	_ = v97
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v105 int32
	_ = v105
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v117 int32
	_ = v117
	var v120 int32
	_ = v120
	var v127 int32
	_ = v127
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
	var v138 int32
	_ = v138
	var v142 int32
	_ = v142
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v156 int32
	_ = v156
	var v158 int32
	_ = v158
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
	var v178 int32
	_ = v178
	var v185 int32
	_ = v185
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v192 int32
	_ = v192
	var v196 int32
	_ = v196
	var v200 int32
	_ = v200
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
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
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v236 int32
	_ = v236
	var v243 int32
	_ = v243
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v248 int32
	_ = v248
	var v250 int32
	_ = v250
	var v254 int32
	_ = v254
	var v258 int32
	_ = v258
	var v260 int32
	_ = v260
	var v261 int32
	_ = v261
	var v262 int32
	_ = v262
	var v264 int32
	_ = v264
	var v265 int32
	_ = v265
	var v266 int32
	_ = v266
	var v267 int32
	_ = v267
	var v272 int32
	_ = v272
	var v274 int32
	_ = v274
	var v277 int32
	_ = v277
	var v278 int32
	_ = v278
	var v282 int32
	_ = v282
	var v283 int32
	_ = v283
	var v284 int32
	_ = v284
	var v294 int32
	_ = v294
	var v295 int32
	_ = v295
	var v296 int32
	_ = v296
	var v298 int32
	_ = v298
	var v299 int32
	_ = v299
	var v303 int32
	_ = v303
	var v305 int32
	_ = v305
	var v306 int32
	_ = v306
	var v307 int32
	_ = v307
	var v308 int32
	_ = v308
	var v310 int32
	_ = v310
	var v314 int32
	_ = v314
	var v318 int32
	_ = v318
	var v320 int32
	_ = v320
	var v321 int32
	_ = v321
	var v322 int32
	_ = v322
	var v324 int32
	_ = v324
	var v325 int32
	_ = v325
	var v326 int32
	_ = v326
	var v327 int32
	_ = v327
	var v332 int32
	_ = v332
	var v334 int32
	_ = v334
	var v337 int32
	_ = v337
	var v338 int32
	_ = v338
	var v342 int32
	_ = v342
	var v343 int32
	_ = v343
	var v344 int32
	_ = v344
	var v351 int32
	_ = v351
	var v354 int32
	_ = v354
	var v356 int32
	_ = v356
	var v360 int32
	_ = v360
	var v362 int32
	_ = v362
	var v363 int32
	_ = v363
	var v364 int32
	_ = v364
	var v365 int32
	_ = v365
	var v367 int32
	_ = v367
	var v371 int32
	_ = v371
	var v375 int32
	_ = v375
	var v377 int32
	_ = v377
	var v378 int32
	_ = v378
	var v379 int32
	_ = v379
	var v381 int32
	_ = v381
	var v382 int32
	_ = v382
	var v383 int32
	_ = v383
	var v384 int32
	_ = v384
	var v389 int32
	_ = v389
	var v391 int32
	_ = v391
	var v394 int32
	_ = v394
	var v395 int32
	_ = v395
	var v399 int32
	_ = v399
	var v400 int32
	_ = v400
	var v401 int32
	_ = v401
	var v408 int32
	_ = v408
	var v411 int32
	_ = v411
	var v413 int32
	_ = v413
	var v417 int32
	_ = v417
	var v419 int32
	_ = v419
	var v420 int32
	_ = v420
	var v421 int32
	_ = v421
	var v422 int32
	_ = v422
	var v424 int32
	_ = v424
	var v428 int32
	_ = v428
	var v432 int32
	_ = v432
	var v434 int32
	_ = v434
	var v435 int32
	_ = v435
	var v436 int32
	_ = v436
	var v438 int32
	_ = v438
	var v439 int32
	_ = v439
	var v440 int32
	_ = v440
	var v441 int32
	_ = v441
	var v446 int32
	_ = v446
	var v448 int32
	_ = v448
	var v451 int32
	_ = v451
	var v452 int32
	_ = v452
	var v456 int32
	_ = v456
	var v457 int32
	_ = v457
	var v458 int32
	_ = v458
	var v465 int32
	_ = v465
	var v471 int32
	_ = v471
	var v478 int32
	_ = v478
	var v479 int32
	_ = v479
	var v483 int32
	_ = v483
	var v484 int32
	_ = v484
	var v485 int32
	_ = v485
	var v486 int32
	_ = v486
	var v487 int32
	_ = v487
	var v489 int32
	_ = v489
	var v494 int64
	_ = v494
	var v500 int32
	_ = v500
	var v503 int32
	_ = v503
	var v505 int32
	_ = v505
	var v511 int64
	_ = v511
	var v515 int64
	_ = v515
	var v517 int64
	_ = v517
	var v519 int32
	_ = v519
	var v520 int32
	_ = v520
	var v522 int32
	_ = v522
	var v523 int32
	_ = v523
	var v525 int32
	_ = v525
	var v533 int32
	_ = v533
	var v534 int32
	_ = v534
	var v537 int32
	_ = v537
	var v539 int32
	_ = v539
	var v544 int32
	_ = v544
	var v546 int32
	_ = v546
	var v547 int32
	_ = v547
	var v548 int32
	_ = v548
	var v549 int32
	_ = v549
	var v553 int32
	_ = v553
	var v558 int32
	_ = v558
	var v561 int32
	_ = v561
	var v562 int32
	_ = v562
	var v563 int32
	_ = v563
	var v568 int32
	_ = v568
	var v570 int32
	_ = v570
	var v573 int32
	_ = v573
	var v574 int32
	_ = v574
	var v578 int32
	_ = v578
	var v588 int32
	_ = v588
	var v590 int32
	_ = v590
	var v595 int32
	_ = v595
	var v597 int32
	_ = v597
	var v598 int32
	_ = v598
	var v599 int32
	_ = v599
	var v600 int32
	_ = v600
	var v604 int32
	_ = v604
	var v609 int32
	_ = v609
	var v612 int32
	_ = v612
	var v613 int32
	_ = v613
	var v614 int32
	_ = v614
	var v619 int32
	_ = v619
	var v621 int32
	_ = v621
	var v624 int32
	_ = v624
	var v625 int32
	_ = v625
	var v629 int32
	_ = v629
	var v637 int32
	_ = v637
	var v640 int32
	_ = v640
	var v642 int32
	_ = v642
	var v648 int64
	_ = v648
	var v652 int64
	_ = v652
	var v654 int64
	_ = v654
	var v656 int32
	_ = v656
	var v657 int32
	_ = v657
	var v659 int32
	_ = v659
	var v665 int32
	_ = v665
	var v667 int32
	_ = v667
	var v670 int32
	_ = v670
	var v671 int32
	_ = v671
	var v672 int32
	_ = v672
	var v673 int32
	_ = v673
	var v674 int32
	_ = v674
	var v675 int32
	_ = v675
	var v677 int32
	_ = v677
	var v679 int32
	_ = v679
	var v681 int32
	_ = v681
	var v684 int32
	_ = v684
	var v686 int32
	_ = v686
	var v690 int32
	_ = v690
	var v691 int32
	_ = v691
	var v693 int32
	_ = v693
	var v694 int32
	_ = v694
	var v703 int32
	_ = v703
	var v708 int32
	_ = v708
	var v711 int32
	_ = v711
	var v713 int32
	_ = v713
	var v714 int32
	_ = v714
	var v716 int32
	_ = v716
	var v721 int32
	_ = v721
	var v733 int32
	_ = v733
	var v736 int32
	_ = v736
	var v737 int32
	_ = v737
	var v741 int32
	_ = v741
	var v743 int32
	_ = v743
	var v744 int32
	_ = v744
	var v746 int32
	_ = v746
	var v750 int32
	_ = v750
	var v751 int32
	_ = v751
	var v752 int32
	_ = v752
	var v753 int32
	_ = v753
	var v755 int32
	_ = v755
	var v756 int32
	_ = v756
	var v757 int32
	_ = v757
	var v762 int32
	_ = v762
	var v763 int32
	_ = v763
	var v764 int32
	_ = v764
	var v765 int32
	_ = v765
	var v767 int32
	_ = v767
	var v768 int32
	_ = v768
	var v769 int32
	_ = v769
	var v776 int32
	_ = v776
	var v777 int32
	_ = v777
	var v778 int32
	_ = v778
	var v779 int32
	_ = v779
	var v782 int64
	_ = v782
	var v786 int64
	_ = v786
	var v790 int64
	_ = v790
	var v794 int32
	_ = v794
	var v796 int32
	_ = v796
	var v809 int32
	_ = v809
	var v810 int32
	_ = v810
	var v814 int32
	_ = v814
	var v828 int32
	_ = v828
	var v834 int32
	_ = v834
	var v840 int32
	_ = v840
	var v843 int32
	_ = v843
	var v845 int32
	_ = v845
	var v849 int32
	_ = v849
	var v850 int32
	_ = v850
	var v852 int32
	_ = v852
	var v853 int32
	_ = v853
	var v862 int32
	_ = v862
	var v867 int32
	_ = v867
	var v870 int32
	_ = v870
	var v872 int32
	_ = v872
	var v875 int32
	_ = v875
	var v889 int32
	_ = v889
	v11 = m.G0
	v13 = v11 - int32(48)
	m.G0 = v13
	*(*int32)(unsafe.Add(mBase, uint32(v13)+44)) = int32(0)
	v19 = F_valkeyGetReply(m, l0, v13+int32(44))
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		goto L6
	} else {
		goto L7
	}
L1:
	;
	m.G0 = v13 + int32(48)
	return
L2:
	;
	F_valkeyAsyncFreeInternal(m, l0)
	mBase = m.M
	v889 = m.ExcPending
	if v889 != 0 {
		goto L6
	} else {
		goto L230
	}
L3:
	;
	if v834 != 0 {
		goto L220
	} else {
		goto L221
	}
L4:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+208)) = l0 + int32(8)
	v828 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+204)) = v828
	v834 = v828
	goto L3
L5:
	;
	if l0 != 0 {
		goto L4
	} else {
		goto L218
	}
L6:
	;
	return
L7:
	;
	if v19 != 0 {
		goto L5
	} else {
		goto L8
	}
L8:
	;
	v24 = v13 + int32(32)
	v28 = v13 + int32(24)
	goto L9
L9:
	;
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v13)+44))
	if v39 != 0 {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v90 = *(*int32)(unsafe.Add(mBase, uint32(v39)))
	if v90 != int32(12) {
		goto L30
	} else {
		goto L31
	}
L12:
	;
	v40 = *(*int32)(unsafe.Add(mBase, uint32(l0)+140))
	if v40&int32(4) == int32(0) {
		goto L1
	} else {
		goto L13
	}
L13:
	;
	v45 = *(*int32)(unsafe.Add(mBase, uint32(l0)+144))
	v48 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v45+int32(-1)))))
	switch v48 & int32(7) {
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
		goto L14
	}
L14:
	;
	v67 = *(*int32)(unsafe.Add(mBase, uint32(l0)+256))
	if v67 != 0 {
		goto L1
	} else {
		goto L22
	}
L15:
	;
	if v65 != 0 {
		goto L1
	} else {
		goto L21
	}
L16:
	;
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v45+int32(-17))))
	v65 = v64
	goto L15
L17:
	;
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v45+int32(-9))))
	v65 = v61
	goto L15
L18:
	;
	v58 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v45+int32(-5)))))
	v65 = v58
	goto L15
L19:
	;
	v55 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v45+int32(-3)))))
	v65 = v55
	goto L15
L20:
	;
	v65 = int32(base.Ui32(v48) >> (uint(int32(3)) % 32))
	goto L15
L21:
	;
	goto L14
L22:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+208)) = l0 + int32(8)
	v71 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+204)) = v71
	if v71 == int32(0) {
		goto L23
	} else {
		goto L24
	}
L23:
	;
	v76 = *(*int32)(unsafe.Add(mBase, uint32(l0)+240))
	if v76 == int32(0) {
		v83 = v40
		goto L25
	} else {
		goto L26
	}
L24:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+140)) = v40
	goto L23
L25:
	;
	v84 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+240)) = v84
	if v83&int32(512) == v84 {
		goto L2
	} else {
		goto L28
	}
L26:
	;
	v79 = *(*int32)(unsafe.Add(mBase, uint32(l0)+220))
	m.T0[v76].(func(*base.Module, int32))(m, v79)
	mBase = m.M
	v81 = m.ExcPending
	if v81 != 0 {
		goto L6
	} else {
		goto L27
	}
L27:
	;
	v82 = *(*int32)(unsafe.Add(mBase, uint32(l0)+140))
	v83 = v82
	goto L25
L28:
	;
	goto L1
L29:
	;
	v809 = F_valkeyGetReply(m, l0, v13+int32(44))
	mBase = m.M
	v810 = m.ExcPending
	if v810 != 0 {
		goto L6
	} else {
		goto L216
	}
L30:
	;
	v494 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v24))) = v494
	*(*int64)(unsafe.Add(mBase, uint32(v28))) = v494
	*(*int64)(unsafe.Add(mBase, uint32(v13)+16)) = v494
	v500 = *(*int32)(unsafe.Add(mBase, uint32(l0)+256))
	if v500 == int32(0) {
		goto L142
	} else {
		goto L143
	}
L31:
	;
	v93 = *(*int32)(unsafe.Add(mBase, uint32(l0)+140))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+140)) = v93 | int32(256)
	v97 = *(*int32)(unsafe.Add(mBase, uint32(v39)+36))
	if v97 == int32(0) {
		goto L32
	} else {
		goto L33
	}
L32:
	;
	v471 = *(*int32)(unsafe.Add(mBase, uint32(l0)+296))
	if v471 == int32(0) {
		v484 = v39
		goto L137
	} else {
		goto L138
	}
L33:
	;
	v100 = *(*int32)(unsafe.Add(mBase, uint32(v39)+40))
	v101 = *(*int32)(unsafe.Add(mBase, uint32(v100)))
	v102 = *(*int32)(unsafe.Add(mBase, uint32(v101)))
	if v102 != int32(1) {
		goto L32
	} else {
		goto L34
	}
L34:
	;
	v105 = *(*int32)(unsafe.Add(mBase, uint32(v101)+24))
	if base.Ui32(v105) < base.Ui32(int32(7)) {
		goto L32
	} else {
		goto L35
	}
L35:
	;
	v108 = int32(1)
	v109 = *(*int32)(unsafe.Add(mBase, uint32(v101)+28))
	v110 = int32(*(*int8)(unsafe.Add(mBase, uint32(v109))))
	if base.Ui32(v110+int32(-65)) < base.Ui32(int32(26)) {
		goto L38
	} else {
		goto L39
	}
L36:
	;
	v295 = v109 + v294
	v296 = m.G3
	v298 = v296 + int32(_a_F_valkeyProcessCallbacks_0)
	v299 = v105 - v294
	if v299 != 0 {
		goto L90
	} else {
		goto L91
	}
L37:
	;
	if v117 == int32(112) {
		v294 = v108
		goto L36
	} else {
		goto L41
	}
L38:
	;
	v117 = v110 | int32(32)
	goto L40
L39:
	;
	v117 = v110
	goto L40
L40:
	;
	goto L37
L41:
	;
	v120 = m.G3
	goto L43
L42:
	;
	if v166-v168 == int32(0) {
		v294 = v108
		goto L36
	} else {
		goto L57
	}
L43:
	;
	v127 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v120)+uint32(_c_F_valkeyProcessCallbacks[0]))))
	if v127 != 0 {
		goto L46
	} else {
		goto L47
	}
L45:
	;
	v166 = F_tolower(m, v161)
	mBase = m.M
	v167 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v162))))
	v168 = F_tolower(m, v167)
	mBase = m.M
	goto L42
L46:
	;
	v129 = v120 + int32(_a_F_valkeyProcessCallbacks_1)
	v130 = v109
	v131 = int32(2)
	v132 = v127
	goto L49
L47:
	;
	v161 = int32(0)
	v162 = v109
	goto L45
L48:
	;
	v161 = v158 & int32(255)
	v162 = v156
	goto L45
L49:
	;
	v134 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v130))))
	if v134 == int32(0) {
		v156 = v130
		v158 = v132
		goto L48
	} else {
		goto L51
	}
L50:
	;
	v156 = v150
	v158 = int32(0)
	goto L48
L51:
	;
	v138 = v131 + int32(-1)
	if v138 == int32(0) {
		v156 = v130
		v158 = v132
		goto L48
	} else {
		goto L52
	}
L52:
	;
	v142 = v132 & int32(255)
	if v142 == v134 {
		goto L53
	} else {
		goto L54
	}
L53:
	;
	v149 = int32(1)
	v150 = v130 + v149
	v151 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v129)+1)))
	if v151 != 0 {
		v129 = v129 + v149
		v130 = v150
		v131 = v138
		v132 = v151
		goto L49
	} else {
		goto L56
	}
L54:
	;
	v144 = F_tolower(m, v142)
	mBase = m.M
	v145 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v130))))
	v146 = F_tolower(m, v145)
	mBase = m.M
	if v144 == v146 {
		goto L53
	} else {
		goto L55
	}
L55:
	;
	v148 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v129))))
	v156 = v130
	v158 = v148
	goto L48
L56:
	;
	goto L50
L57:
	;
	v178 = m.G3
	goto L59
L58:
	;
	if v224-v226 == int32(0) {
		v294 = v108
		goto L36
	} else {
		goto L73
	}
L59:
	;
	v185 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v178)+uint32(_c_F_valkeyProcessCallbacks[1]))))
	if v185 != 0 {
		goto L62
	} else {
		goto L63
	}
L61:
	;
	v224 = F_tolower(m, v219)
	mBase = m.M
	v225 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v220))))
	v226 = F_tolower(m, v225)
	mBase = m.M
	goto L58
L62:
	;
	v187 = v178 + int32(_a_F_valkeyProcessCallbacks_2)
	v188 = v109
	v189 = int32(2)
	v190 = v185
	goto L65
L63:
	;
	v219 = int32(0)
	v220 = v109
	goto L61
L64:
	;
	v219 = v216 & int32(255)
	v220 = v214
	goto L61
L65:
	;
	v192 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v188))))
	if v192 == int32(0) {
		v214 = v188
		v216 = v190
		goto L64
	} else {
		goto L67
	}
L66:
	;
	v214 = v208
	v216 = int32(0)
	goto L64
L67:
	;
	v196 = v189 + int32(-1)
	if v196 == int32(0) {
		v214 = v188
		v216 = v190
		goto L64
	} else {
		goto L68
	}
L68:
	;
	v200 = v190 & int32(255)
	if v200 == v192 {
		goto L69
	} else {
		goto L70
	}
L69:
	;
	v207 = int32(1)
	v208 = v188 + v207
	v209 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v187)+1)))
	if v209 != 0 {
		v187 = v187 + v207
		v188 = v208
		v189 = v196
		v190 = v209
		goto L65
	} else {
		goto L72
	}
L70:
	;
	v202 = F_tolower(m, v200)
	mBase = m.M
	v203 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v188))))
	v204 = F_tolower(m, v203)
	mBase = m.M
	if v202 == v204 {
		goto L69
	} else {
		goto L71
	}
L71:
	;
	v206 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v187))))
	v214 = v188
	v216 = v206
	goto L64
L72:
	;
	goto L66
L73:
	;
	v236 = m.G3
	goto L75
L74:
	;
	v294 = base.B2i32(v282-v284 == int32(0))
	goto L36
L75:
	;
	v243 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v236)+uint32(_c_F_valkeyProcessCallbacks[2]))))
	if v243 != 0 {
		goto L78
	} else {
		goto L79
	}
L77:
	;
	v282 = F_tolower(m, v277)
	mBase = m.M
	v283 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v278))))
	v284 = F_tolower(m, v283)
	mBase = m.M
	goto L74
L78:
	;
	v245 = v236 + int32(_a_F_valkeyProcessCallbacks_3)
	v246 = v109
	v247 = int32(3)
	v248 = v243
	goto L81
L79:
	;
	v277 = int32(0)
	v278 = v109
	goto L77
L80:
	;
	v277 = v274 & int32(255)
	v278 = v272
	goto L77
L81:
	;
	v250 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v246))))
	if v250 == int32(0) {
		v272 = v246
		v274 = v248
		goto L80
	} else {
		goto L83
	}
L82:
	;
	v272 = v266
	v274 = int32(0)
	goto L80
L83:
	;
	v254 = v247 + int32(-1)
	if v254 == int32(0) {
		v272 = v246
		v274 = v248
		goto L80
	} else {
		goto L84
	}
L84:
	;
	v258 = v248 & int32(255)
	if v258 == v250 {
		goto L85
	} else {
		goto L86
	}
L85:
	;
	v265 = int32(1)
	v266 = v246 + v265
	v267 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v245)+1)))
	if v267 != 0 {
		v245 = v245 + v265
		v246 = v266
		v247 = v254
		v248 = v267
		goto L81
	} else {
		goto L88
	}
L86:
	;
	v260 = F_tolower(m, v258)
	mBase = m.M
	v261 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v246))))
	v262 = F_tolower(m, v261)
	mBase = m.M
	if v260 == v262 {
		goto L85
	} else {
		goto L87
	}
L87:
	;
	v264 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v245))))
	v272 = v246
	v274 = v264
	goto L80
L88:
	;
	goto L82
L89:
	;
	if v351 == int32(0) {
		goto L30
	} else {
		goto L104
	}
L90:
	;
	v303 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v295))))
	if v303 != 0 {
		goto L93
	} else {
		goto L94
	}
L91:
	;
	v351 = int32(0)
	goto L89
L92:
	;
	v342 = F_tolower(m, v337)
	mBase = m.M
	v343 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v338))))
	v344 = F_tolower(m, v343)
	mBase = m.M
	v351 = v342 - v344
	goto L89
L93:
	;
	v305 = v295
	v306 = v298
	v307 = v299
	v308 = v303
	goto L96
L94:
	;
	v337 = int32(0)
	v338 = v298
	goto L92
L95:
	;
	v337 = v334 & int32(255)
	v338 = v332
	goto L92
L96:
	;
	v310 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v306))))
	if v310 == int32(0) {
		v332 = v306
		v334 = v308
		goto L95
	} else {
		goto L98
	}
L97:
	;
	v332 = v326
	v334 = int32(0)
	goto L95
L98:
	;
	v314 = v307 + int32(-1)
	if v314 == int32(0) {
		v332 = v306
		v334 = v308
		goto L95
	} else {
		goto L99
	}
L99:
	;
	v318 = v308 & int32(255)
	if v318 == v310 {
		goto L100
	} else {
		goto L101
	}
L100:
	;
	v325 = int32(1)
	v326 = v306 + v325
	v327 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v305)+1)))
	if v327 != 0 {
		v305 = v305 + v325
		v306 = v326
		v307 = v314
		v308 = v327
		goto L96
	} else {
		goto L103
	}
L101:
	;
	v320 = F_tolower(m, v318)
	mBase = m.M
	v321 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v306))))
	v322 = F_tolower(m, v321)
	mBase = m.M
	if v320 == v322 {
		goto L100
	} else {
		goto L102
	}
L102:
	;
	v324 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v305))))
	v332 = v306
	v334 = v324
	goto L95
L103:
	;
	goto L97
L104:
	;
	v354 = m.G3
	v356 = v354 + int32(_a_F_valkeyProcessCallbacks_4)
	if v299 != 0 {
		goto L106
	} else {
		goto L107
	}
L105:
	;
	if v408 == int32(0) {
		goto L30
	} else {
		goto L120
	}
L106:
	;
	v360 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v295))))
	if v360 != 0 {
		goto L109
	} else {
		goto L110
	}
L107:
	;
	v408 = int32(0)
	goto L105
L108:
	;
	v399 = F_tolower(m, v394)
	mBase = m.M
	v400 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v395))))
	v401 = F_tolower(m, v400)
	mBase = m.M
	v408 = v399 - v401
	goto L105
L109:
	;
	v362 = v295
	v363 = v356
	v364 = v299
	v365 = v360
	goto L112
L110:
	;
	v394 = int32(0)
	v395 = v356
	goto L108
L111:
	;
	v394 = v391 & int32(255)
	v395 = v389
	goto L108
L112:
	;
	v367 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v363))))
	if v367 == int32(0) {
		v389 = v363
		v391 = v365
		goto L111
	} else {
		goto L114
	}
L113:
	;
	v389 = v383
	v391 = int32(0)
	goto L111
L114:
	;
	v371 = v364 + int32(-1)
	if v371 == int32(0) {
		v389 = v363
		v391 = v365
		goto L111
	} else {
		goto L115
	}
L115:
	;
	v375 = v365 & int32(255)
	if v375 == v367 {
		goto L116
	} else {
		goto L117
	}
L116:
	;
	v382 = int32(1)
	v383 = v363 + v382
	v384 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v362)+1)))
	if v384 != 0 {
		v362 = v362 + v382
		v363 = v383
		v364 = v371
		v365 = v384
		goto L112
	} else {
		goto L119
	}
L117:
	;
	v377 = F_tolower(m, v375)
	mBase = m.M
	v378 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v363))))
	v379 = F_tolower(m, v378)
	mBase = m.M
	if v377 == v379 {
		goto L116
	} else {
		goto L118
	}
L118:
	;
	v381 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v362))))
	v389 = v363
	v391 = v381
	goto L111
L119:
	;
	goto L113
L120:
	;
	v411 = m.G3
	v413 = v411 + int32(_a_F_valkeyProcessCallbacks_5)
	if v299 != 0 {
		goto L122
	} else {
		goto L123
	}
L121:
	;
	if v465 == int32(0) {
		goto L30
	} else {
		goto L136
	}
L122:
	;
	v417 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v295))))
	if v417 != 0 {
		goto L125
	} else {
		goto L126
	}
L123:
	;
	v465 = int32(0)
	goto L121
L124:
	;
	v456 = F_tolower(m, v451)
	mBase = m.M
	v457 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v452))))
	v458 = F_tolower(m, v457)
	mBase = m.M
	v465 = v456 - v458
	goto L121
L125:
	;
	v419 = v295
	v420 = v413
	v421 = v299
	v422 = v417
	goto L128
L126:
	;
	v451 = int32(0)
	v452 = v413
	goto L124
L127:
	;
	v451 = v448 & int32(255)
	v452 = v446
	goto L124
L128:
	;
	v424 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v420))))
	if v424 == int32(0) {
		v446 = v420
		v448 = v422
		goto L127
	} else {
		goto L130
	}
L129:
	;
	v446 = v440
	v448 = int32(0)
	goto L127
L130:
	;
	v428 = v421 + int32(-1)
	if v428 == int32(0) {
		v446 = v420
		v448 = v422
		goto L127
	} else {
		goto L131
	}
L131:
	;
	v432 = v422 & int32(255)
	if v432 == v424 {
		goto L132
	} else {
		goto L133
	}
L132:
	;
	v439 = int32(1)
	v440 = v420 + v439
	v441 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v419)+1)))
	if v441 != 0 {
		v419 = v419 + v439
		v420 = v440
		v421 = v428
		v422 = v441
		goto L128
	} else {
		goto L135
	}
L133:
	;
	v434 = F_tolower(m, v432)
	mBase = m.M
	v435 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v420))))
	v436 = F_tolower(m, v435)
	mBase = m.M
	if v434 == v436 {
		goto L132
	} else {
		goto L134
	}
L134:
	;
	v438 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v419))))
	v446 = v420
	v448 = v438
	goto L127
L135:
	;
	goto L129
L136:
	;
	goto L32
L137:
	;
	v485 = *(*int32)(unsafe.Add(mBase, uint32(l0)+148))
	v486 = *(*int32)(unsafe.Add(mBase, uint32(v485)+176))
	v487 = *(*int32)(unsafe.Add(mBase, uint32(v486)+24))
	m.T0[v487].(func(*base.Module, int32))(m, v484)
	mBase = m.M
	v489 = m.ExcPending
	if v489 != 0 {
		goto L6
	} else {
		goto L140
	}
L138:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+140)) = v93 | int32(272)
	m.T0[v471].(func(*base.Module, int32, int32))(m, l0, v39)
	mBase = m.M
	v478 = m.ExcPending
	if v478 != 0 {
		goto L6
	} else {
		goto L139
	}
L139:
	;
	v479 = *(*int32)(unsafe.Add(mBase, uint32(l0)+140))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+140)) = v479 & int32(-17)
	v483 = *(*int32)(unsafe.Add(mBase, uint32(v13)+44))
	v484 = v483
	goto L137
L140:
	;
	goto L29
L141:
	;
	v733 = *(*int32)(unsafe.Add(mBase, uint32(v13)+20))
	if v733 == int32(0) {
		goto L201
	} else {
		goto L202
	}
L142:
	;
	v523 = *(*int32)(unsafe.Add(mBase, uint32(l0)+140))
	v525 = v523 & int32(32)
	if v90 == int32(6) {
		goto L148
	} else {
		goto L149
	}
L143:
	;
	v503 = *(*int32)(unsafe.Add(mBase, uint32(v500)))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+256)) = v503
	v505 = *(*int32)(unsafe.Add(mBase, uint32(l0)+260))
	if v500 != v505 {
		goto L144
	} else {
		goto L145
	}
L144:
	;
	v511 = *(*int64)(unsafe.Add(mBase, uint32(v500+int32(16))))
	*(*int64)(unsafe.Add(mBase, uint32(v24))) = v511
	v515 = *(*int64)(unsafe.Add(mBase, uint32(v500+int32(8))))
	*(*int64)(unsafe.Add(mBase, uint32(v28))) = v515
	v517 = *(*int64)(unsafe.Add(mBase, uint32(v500)))
	*(*int64)(unsafe.Add(mBase, uint32(v13)+16)) = v517
	v519 = m.G4
	v520 = *(*int32)(unsafe.Add(mBase, uint32(v519)+16))
	m.T0[v520].(func(*base.Module, int32))(m, v500)
	mBase = m.M
	v522 = m.ExcPending
	if v522 != 0 {
		goto L6
	} else {
		goto L146
	}
L145:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+260)) = int32(0)
	goto L144
L146:
	;
	goto L141
L147:
	;
	v721 = m.G3
	m.Env.X__assert_fail(m, v721+int32(_a_F_valkeyProcessCallbacks_6), v721+int32(_a_F_valkeyProcessCallbacks_7), int32(617), v721+int32(_a_F_valkeyProcessCallbacks_8))
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L148:
	;
	v534 = *(*int32)(unsafe.Add(mBase, uint32(v39)+28))
	if v525 == int32(0) {
		goto L152
	} else {
		goto L153
	}
L149:
	;
	if v525 == int32(0) {
		goto L147
	} else {
		goto L150
	}
L150:
	;
	F_valkeyGetSubscribeCallback(m, l0, v39, v13+int32(16))
	mBase = m.M
	v533 = m.ExcPending
	if v533 != 0 {
		goto L6
	} else {
		goto L151
	}
L151:
	;
	goto L141
L152:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v13))) = v534
	v665 = l0 + int32(8)
	v667 = m.G3
	v670 = F_snprintf(m, v665, int32(128), v667+int32(_a_F_valkeyProcessCallbacks_9), v13)
	mBase = m.M
	v671 = m.ExcPending
	if v671 != 0 {
		goto L6
	} else {
		goto L187
	}
L153:
	;
	v537 = m.G3
	v539 = v537 + int32(_a_F_valkeyProcessCallbacks_10)
	goto L156
L154:
	;
	v637 = *(*int32)(unsafe.Add(mBase, uint32(l0)+272))
	if v637 == int32(0) {
		goto L152
	} else {
		goto L183
	}
L155:
	;
	if v573-v578 == int32(0) {
		goto L154
	} else {
		goto L168
	}
L156:
	;
	v544 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v534))))
	if v544 != 0 {
		goto L159
	} else {
		goto L160
	}
L158:
	;
	v578 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v574))))
	goto L155
L159:
	;
	v546 = v534
	v547 = v539
	v548 = int32(5)
	v549 = v544
	goto L162
L160:
	;
	v573 = int32(0)
	v574 = v539
	goto L158
L161:
	;
	v573 = v570 & int32(255)
	v574 = v568
	goto L158
L162:
	;
	v553 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v547))))
	if v549&int32(255) != v553 {
		v568 = v547
		v570 = v549
		goto L161
	} else {
		goto L164
	}
L163:
	;
	v568 = v562
	v570 = int32(0)
	goto L161
L164:
	;
	if v553 == int32(0) {
		v568 = v547
		v570 = v549
		goto L161
	} else {
		goto L165
	}
L165:
	;
	v558 = v548 + int32(-1)
	if v558 == int32(0) {
		v568 = v547
		v570 = v549
		goto L161
	} else {
		goto L166
	}
L166:
	;
	v561 = int32(1)
	v562 = v547 + v561
	v563 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v546)+1)))
	if v563 != 0 {
		v546 = v546 + v561
		v547 = v562
		v548 = v558
		v549 = v563
		goto L162
	} else {
		goto L167
	}
L167:
	;
	goto L163
L168:
	;
	v588 = m.G3
	v590 = v588 + int32(_a_F_valkeyProcessCallbacks_11)
	goto L170
L169:
	;
	if v624-v629 != 0 {
		goto L152
	} else {
		goto L182
	}
L170:
	;
	v595 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v534))))
	if v595 != 0 {
		goto L173
	} else {
		goto L174
	}
L172:
	;
	v629 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v625))))
	goto L169
L173:
	;
	v597 = v534
	v598 = v590
	v599 = int32(9)
	v600 = v595
	goto L176
L174:
	;
	v624 = int32(0)
	v625 = v590
	goto L172
L175:
	;
	v624 = v621 & int32(255)
	v625 = v619
	goto L172
L176:
	;
	v604 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v598))))
	if v600&int32(255) != v604 {
		v619 = v598
		v621 = v600
		goto L175
	} else {
		goto L178
	}
L177:
	;
	v619 = v613
	v621 = int32(0)
	goto L175
L178:
	;
	if v604 == int32(0) {
		v619 = v598
		v621 = v600
		goto L175
	} else {
		goto L179
	}
L179:
	;
	v609 = v599 + int32(-1)
	if v609 == int32(0) {
		v619 = v598
		v621 = v600
		goto L175
	} else {
		goto L180
	}
L180:
	;
	v612 = int32(1)
	v613 = v598 + v612
	v614 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v597)+1)))
	if v614 != 0 {
		v597 = v597 + v612
		v598 = v613
		v599 = v609
		v600 = v614
		goto L176
	} else {
		goto L181
	}
L181:
	;
	goto L177
L182:
	;
	goto L154
L183:
	;
	v640 = *(*int32)(unsafe.Add(mBase, uint32(v637)))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+272)) = v640
	v642 = *(*int32)(unsafe.Add(mBase, uint32(l0)+276))
	if v637 != v642 {
		goto L184
	} else {
		goto L185
	}
L184:
	;
	v648 = *(*int64)(unsafe.Add(mBase, uint32(v637+int32(16))))
	*(*int64)(unsafe.Add(mBase, uint32(v24))) = v648
	v652 = *(*int64)(unsafe.Add(mBase, uint32(v637+int32(8))))
	*(*int64)(unsafe.Add(mBase, uint32(v28))) = v652
	v654 = *(*int64)(unsafe.Add(mBase, uint32(v637)))
	*(*int64)(unsafe.Add(mBase, uint32(v13)+16)) = v654
	v656 = m.G4
	v657 = *(*int32)(unsafe.Add(mBase, uint32(v656)+16))
	m.T0[v657].(func(*base.Module, int32))(m, v637)
	mBase = m.M
	v659 = m.ExcPending
	if v659 != 0 {
		goto L6
	} else {
		goto L186
	}
L185:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+276)) = int32(0)
	goto L184
L186:
	;
	goto L141
L187:
	;
	v672 = *(*int32)(unsafe.Add(mBase, uint32(v13)+44))
	v673 = *(*int32)(unsafe.Add(mBase, uint32(l0)+148))
	v674 = *(*int32)(unsafe.Add(mBase, uint32(v673)+176))
	v675 = *(*int32)(unsafe.Add(mBase, uint32(v674)+24))
	m.T0[v675].(func(*base.Module, int32))(m, v672)
	mBase = m.M
	v677 = m.ExcPending
	if v677 != 0 {
		goto L6
	} else {
		goto L188
	}
L188:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+208)) = v665
	v679 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+204)) = v679
	if v679 != 0 {
		goto L190
	} else {
		goto L191
	}
L189:
	;
	v708 = *(*int32)(unsafe.Add(mBase, uint32(l0)+240))
	if v708 == int32(0) {
		goto L196
	} else {
		goto L197
	}
L190:
	;
	v703 = *(*int32)(unsafe.Add(mBase, uint32(l0)+140))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+140)) = v703 | int32(4)
	goto L189
L191:
	;
	v681 = *(*int32)(unsafe.Add(mBase, uint32(l0)+256))
	if v681 == int32(0) {
		goto L189
	} else {
		goto L192
	}
L192:
	;
	v684 = *(*int32)(unsafe.Add(mBase, uint32(v681)))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+256)) = v684
	v686 = *(*int32)(unsafe.Add(mBase, uint32(l0)+260))
	if v681 != v686 {
		goto L193
	} else {
		goto L194
	}
L193:
	;
	v690 = m.G4
	v691 = *(*int32)(unsafe.Add(mBase, uint32(v690)+16))
	m.T0[v691].(func(*base.Module, int32))(m, v681)
	mBase = m.M
	v693 = m.ExcPending
	if v693 != 0 {
		goto L6
	} else {
		goto L195
	}
L194:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+260)) = int32(0)
	goto L193
L195:
	;
	v694 = m.G3
	m.Env.X__assert_fail(m, v694+int32(_a_F_valkeyProcessCallbacks_12), v694+int32(_a_F_valkeyProcessCallbacks_7), int32(423), v694+int32(_a_F_valkeyProcessCallbacks_13))
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L196:
	;
	v714 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+240)) = v714
	v716 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+141)))
	if v716&int32(2) == v714 {
		goto L2
	} else {
		goto L199
	}
L197:
	;
	v711 = *(*int32)(unsafe.Add(mBase, uint32(l0)+220))
	m.T0[v708].(func(*base.Module, int32))(m, v711)
	mBase = m.M
	v713 = m.ExcPending
	if v713 != 0 {
		goto L6
	} else {
		goto L198
	}
L198:
	;
	goto L196
L199:
	;
	goto L1
L200:
	;
	if v769&int32(64) == int32(0) {
		goto L29
	} else {
		goto L209
	}
L201:
	;
	v762 = *(*int32)(unsafe.Add(mBase, uint32(v13)+44))
	v763 = *(*int32)(unsafe.Add(mBase, uint32(l0)+148))
	v764 = *(*int32)(unsafe.Add(mBase, uint32(v763)+176))
	v765 = *(*int32)(unsafe.Add(mBase, uint32(v764)+24))
	m.T0[v765].(func(*base.Module, int32))(m, v762)
	mBase = m.M
	v767 = m.ExcPending
	if v767 != 0 {
		goto L6
	} else {
		goto L208
	}
L202:
	;
	v736 = *(*int32)(unsafe.Add(mBase, uint32(v13)+44))
	v737 = *(*int32)(unsafe.Add(mBase, uint32(l0)+140))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+140)) = v737 | int32(16)
	v741 = *(*int32)(unsafe.Add(mBase, uint32(v13)+32))
	m.T0[v733].(func(*base.Module, int32, int32, int32))(m, l0, v736, v741)
	mBase = m.M
	v743 = m.ExcPending
	if v743 != 0 {
		goto L6
	} else {
		goto L203
	}
L203:
	;
	v744 = *(*int32)(unsafe.Add(mBase, uint32(l0)+140))
	v746 = v744 & int32(-17)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+140)) = v746
	if v744&int32(1024) != 0 {
		v757 = v746
		goto L204
	} else {
		goto L205
	}
L204:
	;
	if v757&int32(8) == int32(0) {
		v769 = v757
		goto L200
	} else {
		goto L207
	}
L205:
	;
	v750 = *(*int32)(unsafe.Add(mBase, uint32(v13)+44))
	v751 = *(*int32)(unsafe.Add(mBase, uint32(l0)+148))
	v752 = *(*int32)(unsafe.Add(mBase, uint32(v751)+176))
	v753 = *(*int32)(unsafe.Add(mBase, uint32(v752)+24))
	m.T0[v753].(func(*base.Module, int32))(m, v750)
	mBase = m.M
	v755 = m.ExcPending
	if v755 != 0 {
		goto L6
	} else {
		goto L206
	}
L206:
	;
	v756 = *(*int32)(unsafe.Add(mBase, uint32(l0)+140))
	v757 = v756
	goto L204
L207:
	;
	goto L2
L208:
	;
	v768 = *(*int32)(unsafe.Add(mBase, uint32(l0)+140))
	v769 = v768
	goto L200
L209:
	;
	v776 = m.G4
	v777 = *(*int32)(unsafe.Add(mBase, uint32(v776)))
	v778 = m.T0[v777].(func(*base.Module, int32) int32)(m, int32(24))
	mBase = m.M
	v779 = m.ExcPending
	if v779 != 0 {
		goto L6
	} else {
		goto L210
	}
L210:
	;
	if v778 == int32(0) {
		goto L29
	} else {
		goto L211
	}
L211:
	;
	v782 = *(*int64)(unsafe.Add(mBase, uint32(v13)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v778))) = v782
	v786 = *(*int64)(unsafe.Add(mBase, uint32(v24)))
	*(*int64)(unsafe.Add(mBase, uint32(v778+int32(16)))) = v786
	v790 = *(*int64)(unsafe.Add(mBase, uint32(v28)))
	*(*int64)(unsafe.Add(mBase, uint32(v778+int32(8)))) = v790
	*(*int32)(unsafe.Add(mBase, uint32(v778))) = int32(0)
	v794 = *(*int32)(unsafe.Add(mBase, uint32(l0)+256))
	if v794 != 0 {
		goto L212
	} else {
		goto L213
	}
L212:
	;
	v796 = *(*int32)(unsafe.Add(mBase, uint32(l0)+260))
	if v796 == int32(0) {
		goto L214
	} else {
		goto L215
	}
L213:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+256)) = v778
	goto L212
L214:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+260)) = v778
	goto L29
L215:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v796))) = v778
	goto L214
L216:
	;
	if v809 == int32(0) {
		goto L9
	} else {
		goto L217
	}
L217:
	;
	goto L4
L218:
	;
	v814 = *(*int32)(unsafe.Add(mBase, 204))
	v834 = v814
	goto L3
L219:
	;
	v867 = *(*int32)(unsafe.Add(mBase, uint32(l0)+240))
	if v867 == int32(0) {
		goto L226
	} else {
		goto L227
	}
L220:
	;
	v862 = *(*int32)(unsafe.Add(mBase, uint32(l0)+140))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+140)) = v862 | int32(4)
	goto L219
L221:
	;
	v840 = *(*int32)(unsafe.Add(mBase, uint32(l0)+256))
	if v840 == int32(0) {
		goto L219
	} else {
		goto L222
	}
L222:
	;
	v843 = *(*int32)(unsafe.Add(mBase, uint32(v840)))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+256)) = v843
	v845 = *(*int32)(unsafe.Add(mBase, uint32(l0)+260))
	if v840 != v845 {
		goto L223
	} else {
		goto L224
	}
L223:
	;
	v849 = m.G4
	v850 = *(*int32)(unsafe.Add(mBase, uint32(v849)+16))
	m.T0[v850].(func(*base.Module, int32))(m, v840)
	mBase = m.M
	v852 = m.ExcPending
	if v852 != 0 {
		goto L6
	} else {
		goto L225
	}
L224:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+260)) = int32(0)
	goto L223
L225:
	;
	v853 = m.G3
	m.Env.X__assert_fail(m, v853+int32(_a_F_valkeyProcessCallbacks_12), v853+int32(_a_F_valkeyProcessCallbacks_7), int32(423), v853+int32(_a_F_valkeyProcessCallbacks_13))
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L226:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+240)) = int32(0)
	v875 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+141)))
	if v875&int32(2) != 0 {
		goto L1
	} else {
		goto L229
	}
L227:
	;
	v870 = *(*int32)(unsafe.Add(mBase, uint32(l0)+220))
	m.T0[v867].(func(*base.Module, int32))(m, v870)
	mBase = m.M
	v872 = m.ExcPending
	if v872 != 0 {
		goto L6
	} else {
		goto L228
	}
L228:
	;
	goto L226
L229:
	;
	goto L2
L230:
	;
	goto L1
}
func F_valkeyReaderFree(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
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
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	if l0 == int32(0) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+172))
	if v6 == int32(0) {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+160))
	if v18 == int32(0) {
		goto L9
	} else {
		goto L10
	}
L4:
	;
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+176))
	if v9 == int32(0) {
		goto L3
	} else {
		goto L5
	}
L5:
	;
	v12 = *(*int32)(unsafe.Add(mBase, uint32(v9)+24))
	if v12 == int32(0) {
		goto L3
	} else {
		goto L6
	}
L6:
	;
	m.T0[v12].(func(*base.Module, int32))(m, v6)
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	return
L8:
	;
	goto L3
L9:
	;
	v52 = *(*int32)(unsafe.Add(mBase, uint32(l0)+132))
	F_sdsfree(m, v52)
	mBase = m.M
	v54 = m.ExcPending
	if v54 != 0 {
		goto L7
	} else {
		goto L18
	}
L10:
	;
	v21 = int32(0)
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+164))
	if v22 <= v21 {
		v44 = v18
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v45 = m.G4
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v45)+16))
	m.T0[v46].(func(*base.Module, int32))(m, v44)
	mBase = m.M
	v48 = m.ExcPending
	if v48 != 0 {
		goto L7
	} else {
		goto L17
	}
L12:
	;
	v26 = v21
	goto L13
L13:
	;
	v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+160))
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v28+v26<<(uint(int32(2))%32))))
	v33 = m.G4
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v33)+16))
	m.T0[v34].(func(*base.Module, int32))(m, v32)
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
		goto L7
	} else {
		goto L15
	}
L14:
	;
	v41 = *(*int32)(unsafe.Add(mBase, uint32(l0)+160))
	v44 = v41
	goto L11
L15:
	;
	v38 = v26 + int32(1)
	v39 = *(*int32)(unsafe.Add(mBase, uint32(l0)+164))
	if v38 < v39 {
		v26 = v38
		goto L13
	} else {
		goto L16
	}
L16:
	;
	goto L14
L17:
	;
	goto L9
L18:
	;
	v55 = m.G4
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v55)+16))
	m.T0[v56].(func(*base.Module, int32))(m, l0)
	mBase = m.M
	v58 = m.ExcPending
	if v58 != 0 {
		goto L7
	} else {
		goto L19
	}
L19:
	;
	goto L1
}
func F_valkeyReaderGetReply(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v29 int32
	_ = v29
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v58 int32
	_ = v58
	var v65 int32
	_ = v65
	var v81 int32
	_ = v81
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v95 int32
	_ = v95
	var v102 int32
	_ = v102
	var v104 int32
	_ = v104
	var v123 int32
	_ = v123
	var v128 int32
	_ = v128
	var v131 int32
	_ = v131
	var v136 int32
	_ = v136
	var v141 int32
	_ = v141
	var v143 int32
	_ = v143
	var v149 int32
	_ = v149
	var v154 int32
	_ = v154
	var v156 int32
	_ = v156
	var v162 int32
	_ = v162
	var v167 int32
	_ = v167
	var v169 int32
	_ = v169
	var v175 int32
	_ = v175
	var v180 int32
	_ = v180
	var v182 int32
	_ = v182
	var v188 int32
	_ = v188
	var v193 int32
	_ = v193
	var v195 int32
	_ = v195
	var v205 int32
	_ = v205
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v216 int32
	_ = v216
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v224 int32
	_ = v224
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v238 int32
	_ = v238
	var v241 int32
	_ = v241
	var v245 int32
	_ = v245
	var v249 int32
	_ = v249
	var v251 int32
	_ = v251
	var v252 int32
	_ = v252
	var v258 int32
	_ = v258
	var v261 int32
	_ = v261
	var v263 int32
	_ = v263
	var v273 int32
	_ = v273
	var v276 int32
	_ = v276
	var v280 int32
	_ = v280
	var v285 int32
	_ = v285
	var v287 int32
	_ = v287
	var v291 int32
	_ = v291
	var v297 int32
	_ = v297
	var v300 int32
	_ = v300
	var v306 int32
	_ = v306
	var v310 int32
	_ = v310
	var v312 int32
	_ = v312
	var v320 int32
	_ = v320
	var v321 int32
	_ = v321
	var v324 int32
	_ = v324
	var v327 int32
	_ = v327
	var v328 int32
	_ = v328
	var v330 int32
	_ = v330
	var v335 int32
	_ = v335
	var v345 int32
	_ = v345
	var v355 int32
	_ = v355
	var v361 int32
	_ = v361
	var v362 int32
	_ = v362
	var v363 int32
	_ = v363
	var v364 int32
	_ = v364
	var v367 int32
	_ = v367
	var v368 int32
	_ = v368
	var v371 int32
	_ = v371
	var v372 int32
	_ = v372
	var v375 int32
	_ = v375
	var v384 int32
	_ = v384
	var v386 int32
	_ = v386
	var v389 int32
	_ = v389
	var v392 int32
	_ = v392
	var v393 int32
	_ = v393
	var v394 int32
	_ = v394
	var v396 int32
	_ = v396
	var v401 int32
	_ = v401
	var v403 int32
	_ = v403
	var v404 int32
	_ = v404
	var v408 int32
	_ = v408
	var v418 int32
	_ = v418
	var v420 int32
	_ = v420
	var v424 int32
	_ = v424
	var v425 int32
	_ = v425
	var v428 int32
	_ = v428
	var v433 int32
	_ = v433
	var v435 int32
	_ = v435
	var v438 int32
	_ = v438
	var v440 int32
	_ = v440
	var v445 int32
	_ = v445
	var v447 int32
	_ = v447
	var v452 int32
	_ = v452
	var v454 int32
	_ = v454
	var v457 int32
	_ = v457
	var v462 int32
	_ = v462
	var v474 int32
	_ = v474
	var v482 int32
	_ = v482
	var v483 int32
	_ = v483
	var v486 int32
	_ = v486
	var v497 int32
	_ = v497
	var v501 int32
	_ = v501
	var v503 int32
	_ = v503
	var v505 int32
	_ = v505
	var v508 int32
	_ = v508
	var v517 int32
	_ = v517
	var v519 int32
	_ = v519
	var v522 int32
	_ = v522
	var v525 int32
	_ = v525
	var v526 int32
	_ = v526
	var v527 int32
	_ = v527
	var v529 int32
	_ = v529
	var v534 int32
	_ = v534
	var v536 int32
	_ = v536
	var v537 int32
	_ = v537
	var v541 int32
	_ = v541
	var v551 int32
	_ = v551
	var v553 int32
	_ = v553
	var v557 int32
	_ = v557
	var v558 int32
	_ = v558
	var v561 int32
	_ = v561
	var v566 int32
	_ = v566
	var v568 int32
	_ = v568
	var v571 int32
	_ = v571
	var v573 int32
	_ = v573
	var v578 int32
	_ = v578
	var v580 int32
	_ = v580
	var v585 int32
	_ = v585
	var v587 int32
	_ = v587
	var v590 int32
	_ = v590
	var v595 int32
	_ = v595
	var v607 int32
	_ = v607
	var v608 int32
	_ = v608
	var v619 int32
	_ = v619
	var v628 int32
	_ = v628
	var v637 int32
	_ = v637
	var v652 int32
	_ = v652
	var v660 int32
	_ = v660
	var v682 int32
	_ = v682
	var v683 int32
	_ = v683
	var v686 int32
	_ = v686
	var v689 int32
	_ = v689
	var v690 int32
	_ = v690
	var v691 int32
	_ = v691
	var v700 int32
	_ = v700
	var v715 int32
	_ = v715
	var v718 int32
	_ = v718
	var v721 int32
	_ = v721
	var v724 int32
	_ = v724
	var v728 int32
	_ = v728
	var v731 int32
	_ = v731
	var v733 int32
	_ = v733
	var v735 int32
	_ = v735
	var v736 int32
	_ = v736
	var v740 int32
	_ = v740
	var v747 int32
	_ = v747
	var v750 int64
	_ = v750
	var v756 int64
	_ = v756
	var v762 int64
	_ = v762
	var v765 int32
	_ = v765
	var v787 int32
	_ = v787
	var v790 int32
	_ = v790
	var v793 int32
	_ = v793
	var v794 int32
	_ = v794
	var v796 int32
	_ = v796
	var v799 int32
	_ = v799
	var v803 int32
	_ = v803
	var v804 int32
	_ = v804
	var v805 int32
	_ = v805
	var v808 int32
	_ = v808
	var v809 int32
	_ = v809
	var v812 int32
	_ = v812
	var v813 int32
	_ = v813
	var v816 int32
	_ = v816
	var v825 int32
	_ = v825
	var v827 int32
	_ = v827
	var v830 int32
	_ = v830
	var v833 int32
	_ = v833
	var v834 int32
	_ = v834
	var v835 int32
	_ = v835
	var v837 int32
	_ = v837
	var v842 int32
	_ = v842
	var v844 int32
	_ = v844
	var v845 int32
	_ = v845
	var v849 int32
	_ = v849
	var v859 int32
	_ = v859
	var v861 int32
	_ = v861
	var v865 int32
	_ = v865
	var v866 int32
	_ = v866
	var v869 int32
	_ = v869
	var v874 int32
	_ = v874
	var v876 int32
	_ = v876
	var v879 int32
	_ = v879
	var v881 int32
	_ = v881
	var v886 int32
	_ = v886
	var v888 int32
	_ = v888
	var v893 int32
	_ = v893
	var v895 int32
	_ = v895
	var v898 int32
	_ = v898
	var v903 int32
	_ = v903
	var v915 int32
	_ = v915
	var v923 int32
	_ = v923
	var v924 int32
	_ = v924
	var v927 int32
	_ = v927
	var v938 int32
	_ = v938
	var v942 int32
	_ = v942
	var v944 int32
	_ = v944
	var v946 int32
	_ = v946
	var v949 int32
	_ = v949
	var v958 int32
	_ = v958
	var v960 int32
	_ = v960
	var v963 int32
	_ = v963
	var v966 int32
	_ = v966
	var v967 int32
	_ = v967
	var v968 int32
	_ = v968
	var v970 int32
	_ = v970
	var v975 int32
	_ = v975
	var v977 int32
	_ = v977
	var v978 int32
	_ = v978
	var v982 int32
	_ = v982
	var v992 int32
	_ = v992
	var v994 int32
	_ = v994
	var v998 int32
	_ = v998
	var v999 int32
	_ = v999
	var v1002 int32
	_ = v1002
	var v1007 int32
	_ = v1007
	var v1009 int32
	_ = v1009
	var v1012 int32
	_ = v1012
	var v1014 int32
	_ = v1014
	var v1019 int32
	_ = v1019
	var v1021 int32
	_ = v1021
	var v1026 int32
	_ = v1026
	var v1028 int32
	_ = v1028
	var v1031 int32
	_ = v1031
	var v1036 int32
	_ = v1036
	var v1048 int32
	_ = v1048
	var v1050 int32
	_ = v1050
	var v1051 int32
	_ = v1051
	var v1054 int64
	_ = v1054
	var v1059 int32
	_ = v1059
	var v1064 int32
	_ = v1064
	var v1066 int32
	_ = v1066
	var v1067 int32
	_ = v1067
	var v1068 int32
	_ = v1068
	var v1079 int64
	_ = v1079
	var v1097 int32
	_ = v1097
	var v1101 int32
	_ = v1101
	var v1108 int64
	_ = v1108
	var v1112 int32
	_ = v1112
	var v1122 int64
	_ = v1122
	var v1127 int64
	_ = v1127
	var v1131 int32
	_ = v1131
	var v1133 int64
	_ = v1133
	var v1135 int32
	_ = v1135
	var v1140 int32
	_ = v1140
	var v1147 int32
	_ = v1147
	var v1152 int32
	_ = v1152
	var v1155 int32
	_ = v1155
	var v1156 int32
	_ = v1156
	var v1157 int32
	_ = v1157
	var v1158 int32
	_ = v1158
	var v1162 int32
	_ = v1162
	var v1186 int32
	_ = v1186
	var v1187 int32
	_ = v1187
	var v1188 int32
	_ = v1188
	var v1189 int32
	_ = v1189
	var v1190 int32
	_ = v1190
	var v1191 int32
	_ = v1191
	var v1199 int32
	_ = v1199
	var v1222 int32
	_ = v1222
	var v1223 int32
	_ = v1223
	var v1224 int32
	_ = v1224
	var v1227 int32
	_ = v1227
	var v1228 int32
	_ = v1228
	var v1231 int32
	_ = v1231
	var v1232 int32
	_ = v1232
	var v1235 int32
	_ = v1235
	var v1244 int32
	_ = v1244
	var v1246 int32
	_ = v1246
	var v1249 int32
	_ = v1249
	var v1252 int32
	_ = v1252
	var v1253 int32
	_ = v1253
	var v1254 int32
	_ = v1254
	var v1256 int32
	_ = v1256
	var v1261 int32
	_ = v1261
	var v1263 int32
	_ = v1263
	var v1264 int32
	_ = v1264
	var v1268 int32
	_ = v1268
	var v1278 int32
	_ = v1278
	var v1280 int32
	_ = v1280
	var v1284 int32
	_ = v1284
	var v1285 int32
	_ = v1285
	var v1288 int32
	_ = v1288
	var v1293 int32
	_ = v1293
	var v1295 int32
	_ = v1295
	var v1298 int32
	_ = v1298
	var v1300 int32
	_ = v1300
	var v1305 int32
	_ = v1305
	var v1307 int32
	_ = v1307
	var v1312 int32
	_ = v1312
	var v1314 int32
	_ = v1314
	var v1317 int32
	_ = v1317
	var v1322 int32
	_ = v1322
	var v1334 int32
	_ = v1334
	var v1342 int32
	_ = v1342
	var v1343 int32
	_ = v1343
	var v1346 int32
	_ = v1346
	var v1357 int32
	_ = v1357
	var v1361 int32
	_ = v1361
	var v1363 int32
	_ = v1363
	var v1365 int32
	_ = v1365
	var v1368 int32
	_ = v1368
	var v1377 int32
	_ = v1377
	var v1379 int32
	_ = v1379
	var v1382 int32
	_ = v1382
	var v1385 int32
	_ = v1385
	var v1386 int32
	_ = v1386
	var v1387 int32
	_ = v1387
	var v1389 int32
	_ = v1389
	var v1394 int32
	_ = v1394
	var v1396 int32
	_ = v1396
	var v1397 int32
	_ = v1397
	var v1401 int32
	_ = v1401
	var v1411 int32
	_ = v1411
	var v1413 int32
	_ = v1413
	var v1417 int32
	_ = v1417
	var v1418 int32
	_ = v1418
	var v1421 int32
	_ = v1421
	var v1426 int32
	_ = v1426
	var v1428 int32
	_ = v1428
	var v1431 int32
	_ = v1431
	var v1433 int32
	_ = v1433
	var v1438 int32
	_ = v1438
	var v1440 int32
	_ = v1440
	var v1445 int32
	_ = v1445
	var v1447 int32
	_ = v1447
	var v1450 int32
	_ = v1450
	var v1455 int32
	_ = v1455
	var v1467 int32
	_ = v1467
	var v1488 int32
	_ = v1488
	var v1491 int32
	_ = v1491
	var v1494 int32
	_ = v1494
	var v1498 int32
	_ = v1498
	var v1502 int32
	_ = v1502
	var v1504 int32
	_ = v1504
	var v1505 int32
	_ = v1505
	var v1509 int32
	_ = v1509
	var v1516 int32
	_ = v1516
	var v1519 int64
	_ = v1519
	var v1525 int64
	_ = v1525
	var v1527 int32
	_ = v1527
	var v1535 int32
	_ = v1535
	var v1542 int32
	_ = v1542
	var v1547 int32
	_ = v1547
	var v1549 int32
	_ = v1549
	var v1550 int32
	_ = v1550
	var v1551 int32
	_ = v1551
	var v1562 int64
	_ = v1562
	var v1580 int32
	_ = v1580
	var v1581 int32
	_ = v1581
	var v1591 int64
	_ = v1591
	var v1595 int32
	_ = v1595
	var v1605 int64
	_ = v1605
	var v1610 int64
	_ = v1610
	var v1614 int32
	_ = v1614
	var v1616 int64
	_ = v1616
	var v1618 int32
	_ = v1618
	var v1622 int32
	_ = v1622
	var v1642 int64
	_ = v1642
	var v1672 int32
	_ = v1672
	var v1675 int32
	_ = v1675
	var v1678 int32
	_ = v1678
	var v1682 int32
	_ = v1682
	var v1685 int32
	_ = v1685
	var v1687 int32
	_ = v1687
	var v1689 int32
	_ = v1689
	var v1690 int32
	_ = v1690
	var v1694 int32
	_ = v1694
	var v1701 int32
	_ = v1701
	var v1704 int64
	_ = v1704
	var v1710 int64
	_ = v1710
	var v1716 int64
	_ = v1716
	var v1734 int64
	_ = v1734
	var v1742 int32
	_ = v1742
	var v1743 int64
	_ = v1743
	var v1749 int32
	_ = v1749
	var v1752 int32
	_ = v1752
	var v1755 int32
	_ = v1755
	var v1759 int32
	_ = v1759
	var v1762 int32
	_ = v1762
	var v1764 int32
	_ = v1764
	var v1766 int32
	_ = v1766
	var v1767 int32
	_ = v1767
	var v1771 int32
	_ = v1771
	var v1778 int32
	_ = v1778
	var v1781 int64
	_ = v1781
	var v1787 int64
	_ = v1787
	var v1793 int64
	_ = v1793
	var v1799 int64
	_ = v1799
	var v1805 int32
	_ = v1805
	var v1806 int32
	_ = v1806
	var v1809 int32
	_ = v1809
	var v1812 int32
	_ = v1812
	var v1813 int32
	_ = v1813
	var v1815 int32
	_ = v1815
	var v1816 int32
	_ = v1816
	var v1830 int32
	_ = v1830
	var v1837 int64
	_ = v1837
	var v1841 int32
	_ = v1841
	var v1846 int64
	_ = v1846
	var v1847 int32
	_ = v1847
	var v1850 int32
	_ = v1850
	var v1854 int32
	_ = v1854
	var v1855 int32
	_ = v1855
	var v1857 int32
	_ = v1857
	var v1859 int32
	_ = v1859
	var v1864 int32
	_ = v1864
	var v1866 int32
	_ = v1866
	var v1868 int32
	_ = v1868
	var v1871 int32
	_ = v1871
	var v1872 int32
	_ = v1872
	var v1880 int32
	_ = v1880
	var v1881 int32
	_ = v1881
	var v1892 int32
	_ = v1892
	var v1923 int64
	_ = v1923
	var v1953 int32
	_ = v1953
	var v1956 int32
	_ = v1956
	var v1959 int32
	_ = v1959
	var v1963 int32
	_ = v1963
	var v1966 int32
	_ = v1966
	var v1968 int32
	_ = v1968
	var v1970 int32
	_ = v1970
	var v1971 int32
	_ = v1971
	var v1975 int32
	_ = v1975
	var v1982 int32
	_ = v1982
	var v1985 int64
	_ = v1985
	var v1991 int64
	_ = v1991
	var v1997 int64
	_ = v1997
	var v2015 int64
	_ = v2015
	var v2023 int32
	_ = v2023
	var v2026 int32
	_ = v2026
	var v2029 int32
	_ = v2029
	var v2033 int32
	_ = v2033
	var v2036 int32
	_ = v2036
	var v2038 int32
	_ = v2038
	var v2040 int32
	_ = v2040
	var v2041 int32
	_ = v2041
	var v2045 int32
	_ = v2045
	var v2052 int32
	_ = v2052
	var v2055 int64
	_ = v2055
	var v2061 int64
	_ = v2061
	var v2067 int64
	_ = v2067
	var v2073 int64
	_ = v2073
	var v2078 int32
	_ = v2078
	var v2079 int32
	_ = v2079
	var v2080 int32
	_ = v2080
	var v2083 int32
	_ = v2083
	var v2086 int32
	_ = v2086
	var v2087 int32
	_ = v2087
	var v2104 int64
	_ = v2104
	var v2108 int32
	_ = v2108
	var v2111 int32
	_ = v2111
	var v2115 int32
	_ = v2115
	var v2118 int32
	_ = v2118
	var v2121 int32
	_ = v2121
	var v2124 int32
	_ = v2124
	var v2127 int32
	_ = v2127
	var v2131 int32
	_ = v2131
	var v2134 int32
	_ = v2134
	var v2136 int32
	_ = v2136
	var v2138 int32
	_ = v2138
	var v2143 int32
	_ = v2143
	var v2146 int32
	_ = v2146
	var v2150 int32
	_ = v2150
	var v2156 int32
	_ = v2156
	var v2158 int32
	_ = v2158
	var v2160 int32
	_ = v2160
	var v2163 int32
	_ = v2163
	var v2168 int32
	_ = v2168
	var v2169 int32
	_ = v2169
	var v2175 int32
	_ = v2175
	var v2178 int32
	_ = v2178
	var v2192 int32
	_ = v2192
	var v2193 int32
	_ = v2193
	var v2198 int32
	_ = v2198
	var v2199 int32
	_ = v2199
	var v2202 int32
	_ = v2202
	var v2204 int32
	_ = v2204
	var v2217 int32
	_ = v2217
	var v2220 int32
	_ = v2220
	var v2223 int32
	_ = v2223
	var v2227 int32
	_ = v2227
	var v2231 int32
	_ = v2231
	var v2233 int32
	_ = v2233
	var v2234 int32
	_ = v2234
	var v2238 int32
	_ = v2238
	var v2245 int32
	_ = v2245
	var v2248 int64
	_ = v2248
	var v2254 int64
	_ = v2254
	var v2258 int32
	_ = v2258
	var v2263 int32
	_ = v2263
	var v2281 int32
	_ = v2281
	var v2284 int32
	_ = v2284
	var v2285 int32
	_ = v2285
	var v2292 int64
	_ = v2292
	var v2295 int32
	_ = v2295
	var v2296 int32
	_ = v2296
	var v2297 int64
	_ = v2297
	var v2300 int32
	_ = v2300
	var v2332 int32
	_ = v2332
	var v2341 int32
	_ = v2341
	var v2352 int32
	_ = v2352
	var v2354 int32
	_ = v2354
	var v2355 int32
	_ = v2355
	var v2356 int32
	_ = v2356
	var v2369 int32
	_ = v2369
	var v2371 int32
	_ = v2371
	var v2374 int32
	_ = v2374
	var v2377 int32
	_ = v2377
	var v2378 int32
	_ = v2378
	var v2379 int32
	_ = v2379
	var v2381 int32
	_ = v2381
	var v2386 int32
	_ = v2386
	var v2388 int32
	_ = v2388
	var v2389 int32
	_ = v2389
	var v2393 int32
	_ = v2393
	var v2403 int32
	_ = v2403
	var v2405 int32
	_ = v2405
	var v2409 int32
	_ = v2409
	var v2410 int32
	_ = v2410
	var v2413 int32
	_ = v2413
	var v2418 int32
	_ = v2418
	var v2420 int32
	_ = v2420
	var v2423 int32
	_ = v2423
	var v2425 int32
	_ = v2425
	var v2430 int32
	_ = v2430
	var v2432 int32
	_ = v2432
	var v2437 int32
	_ = v2437
	var v2439 int32
	_ = v2439
	var v2442 int32
	_ = v2442
	var v2447 int32
	_ = v2447
	var v2459 int32
	_ = v2459
	var v2461 int32
	_ = v2461
	var v2464 int32
	_ = v2464
	var v2467 int32
	_ = v2467
	var v2471 int32
	_ = v2471
	var v2474 int32
	_ = v2474
	var v2476 int32
	_ = v2476
	var v2478 int32
	_ = v2478
	var v2479 int32
	_ = v2479
	var v2483 int32
	_ = v2483
	var v2490 int32
	_ = v2490
	var v2493 int64
	_ = v2493
	var v2499 int64
	_ = v2499
	var v2501 int32
	_ = v2501
	var v2502 int32
	_ = v2502
	var v2505 int32
	_ = v2505
	var v2512 int32
	_ = v2512
	var v2513 int32
	_ = v2513
	var v2515 int32
	_ = v2515
	var v2518 int32
	_ = v2518
	var v2521 int32
	_ = v2521
	var v2525 int32
	_ = v2525
	var v2528 int32
	_ = v2528
	var v2530 int32
	_ = v2530
	var v2532 int32
	_ = v2532
	var v2533 int32
	_ = v2533
	var v2537 int32
	_ = v2537
	var v2544 int32
	_ = v2544
	var v2547 int64
	_ = v2547
	var v2553 int64
	_ = v2553
	var v2555 int32
	_ = v2555
	var v2556 int32
	_ = v2556
	var v2559 int32
	_ = v2559
	var v2562 int32
	_ = v2562
	var v2563 int32
	_ = v2563
	var v2566 int32
	_ = v2566
	var v2569 int32
	_ = v2569
	var v2572 int32
	_ = v2572
	var v2576 int32
	_ = v2576
	var v2579 int32
	_ = v2579
	var v2581 int32
	_ = v2581
	var v2583 int32
	_ = v2583
	var v2592 int32
	_ = v2592
	var v2595 int64
	_ = v2595
	var v2601 int64
	_ = v2601
	var v2607 int64
	_ = v2607
	var v2613 int32
	_ = v2613
	var v2620 int32
	_ = v2620
	var v2624 int32
	_ = v2624
	var v2625 int32
	_ = v2625
	var v2630 int32
	_ = v2630
	var v2631 int32
	_ = v2631
	var v2633 int32
	_ = v2633
	var v2636 int32
	_ = v2636
	var v2638 int32
	_ = v2638
	var v2639 int32
	_ = v2639
	var v2640 int32
	_ = v2640
	var v2642 int32
	_ = v2642
	var v2646 int32
	_ = v2646
	var v2648 int32
	_ = v2648
	var v2649 int32
	_ = v2649
	var v2650 int32
	_ = v2650
	var v2652 int32
	_ = v2652
	var v2653 int32
	_ = v2653
	var v2654 int32
	_ = v2654
	var v2655 int32
	_ = v2655
	var v2660 int32
	_ = v2660
	var v2661 int32
	_ = v2661
	var v2664 int32
	_ = v2664
	var v2665 int32
	_ = v2665
	var v2668 int32
	_ = v2668
	var v2669 int32
	_ = v2669
	var v2670 int32
	_ = v2670
	var v2674 int32
	_ = v2674
	var v2675 int32
	_ = v2675
	var v2677 int32
	_ = v2677
	var v2680 int32
	_ = v2680
	var v2682 int32
	_ = v2682
	var v2683 int32
	_ = v2683
	var v2684 int32
	_ = v2684
	var v2686 int32
	_ = v2686
	var v2690 int32
	_ = v2690
	var v2692 int32
	_ = v2692
	var v2693 int32
	_ = v2693
	var v2694 int32
	_ = v2694
	var v2696 int32
	_ = v2696
	var v2697 int32
	_ = v2697
	var v2698 int32
	_ = v2698
	var v2699 int32
	_ = v2699
	var v2704 int32
	_ = v2704
	var v2705 int32
	_ = v2705
	var v2708 int32
	_ = v2708
	var v2709 int32
	_ = v2709
	var v2712 int32
	_ = v2712
	var v2713 int32
	_ = v2713
	var v2714 int32
	_ = v2714
	var v2718 int32
	_ = v2718
	var v2719 int32
	_ = v2719
	var v2721 int32
	_ = v2721
	var v2724 int32
	_ = v2724
	var v2726 int32
	_ = v2726
	var v2727 int32
	_ = v2727
	var v2728 int32
	_ = v2728
	var v2730 int32
	_ = v2730
	var v2734 int32
	_ = v2734
	var v2736 int32
	_ = v2736
	var v2737 int32
	_ = v2737
	var v2738 int32
	_ = v2738
	var v2740 int32
	_ = v2740
	var v2741 int32
	_ = v2741
	var v2742 int32
	_ = v2742
	var v2743 int32
	_ = v2743
	var v2748 int32
	_ = v2748
	var v2749 int32
	_ = v2749
	var v2752 int32
	_ = v2752
	var v2753 int32
	_ = v2753
	var v2756 int32
	_ = v2756
	var v2757 int32
	_ = v2757
	var v2758 int32
	_ = v2758
	var v2762 int32
	_ = v2762
	var v2763 int32
	_ = v2763
	var v2765 int32
	_ = v2765
	var v2768 int32
	_ = v2768
	var v2770 int32
	_ = v2770
	var v2771 int32
	_ = v2771
	var v2772 int32
	_ = v2772
	var v2774 int32
	_ = v2774
	var v2778 int32
	_ = v2778
	var v2780 int32
	_ = v2780
	var v2781 int32
	_ = v2781
	var v2782 int32
	_ = v2782
	var v2784 int32
	_ = v2784
	var v2785 int32
	_ = v2785
	var v2786 int32
	_ = v2786
	var v2787 int32
	_ = v2787
	var v2792 int32
	_ = v2792
	var v2793 int32
	_ = v2793
	var v2796 int32
	_ = v2796
	var v2797 int32
	_ = v2797
	var v2800 int32
	_ = v2800
	var v2801 int32
	_ = v2801
	var v2802 int32
	_ = v2802
	var v2809 float64
	_ = v2809
	var v2810 int32
	_ = v2810
	var v2811 int32
	_ = v2811
	var v2814 int32
	_ = v2814
	var v2821 int32
	_ = v2821
	var v2824 int32
	_ = v2824
	var v2827 int32
	_ = v2827
	var v2831 int32
	_ = v2831
	var v2835 int32
	_ = v2835
	var v2837 int32
	_ = v2837
	var v2846 int32
	_ = v2846
	var v2849 int64
	_ = v2849
	var v2855 int64
	_ = v2855
	var v2859 int32
	_ = v2859
	var v2863 int32
	_ = v2863
	var v2865 float64
	_ = v2865
	var v2866 int32
	_ = v2866
	var v2867 int32
	_ = v2867
	var v2870 int32
	_ = v2870
	var v2875 int32
	_ = v2875
	var v2876 int32
	_ = v2876
	var v2878 int32
	_ = v2878
	var v2881 int64
	_ = v2881
	var v2886 int32
	_ = v2886
	var v2891 int32
	_ = v2891
	var v2893 int32
	_ = v2893
	var v2894 int32
	_ = v2894
	var v2895 int32
	_ = v2895
	var v2906 int64
	_ = v2906
	var v2925 int32
	_ = v2925
	var v2928 int32
	_ = v2928
	var v2935 int64
	_ = v2935
	var v2939 int32
	_ = v2939
	var v2949 int64
	_ = v2949
	var v2954 int64
	_ = v2954
	var v2958 int32
	_ = v2958
	var v2960 int64
	_ = v2960
	var v2962 int32
	_ = v2962
	var v2982 int64
	_ = v2982
	var v3012 int32
	_ = v3012
	var v3015 int32
	_ = v3015
	var v3018 int32
	_ = v3018
	var v3022 int32
	_ = v3022
	var v3025 int32
	_ = v3025
	var v3027 int32
	_ = v3027
	var v3029 int32
	_ = v3029
	var v3030 int32
	_ = v3030
	var v3034 int32
	_ = v3034
	var v3041 int32
	_ = v3041
	var v3044 int64
	_ = v3044
	var v3050 int64
	_ = v3050
	var v3056 int32
	_ = v3056
	var v3074 int64
	_ = v3074
	var v3078 int32
	_ = v3078
	var v3079 int32
	_ = v3079
	var v3082 int32
	_ = v3082
	var v3085 int32
	_ = v3085
	var v3086 int32
	_ = v3086
	var v3098 int32
	_ = v3098
	var v3109 int32
	_ = v3109
	var v3114 int32
	_ = v3114
	var v3121 int32
	_ = v3121
	var v3131 int32
	_ = v3131
	var v3134 int32
	_ = v3134
	var v3137 int32
	_ = v3137
	var v3141 int32
	_ = v3141
	var v3145 int32
	_ = v3145
	var v3147 int32
	_ = v3147
	var v3148 int32
	_ = v3148
	var v3152 int32
	_ = v3152
	var v3159 int32
	_ = v3159
	var v3162 int64
	_ = v3162
	var v3168 int64
	_ = v3168
	var v3172 int32
	_ = v3172
	var v3177 int32
	_ = v3177
	var v3195 int32
	_ = v3195
	var v3198 int32
	_ = v3198
	var v3199 int32
	_ = v3199
	var v3206 int64
	_ = v3206
	var v3209 int32
	_ = v3209
	var v3210 int32
	_ = v3210
	var v3211 int64
	_ = v3211
	var v3214 int32
	_ = v3214
	var v3266 int32
	_ = v3266
	var v3271 int32
	_ = v3271
	var v3272 int32
	_ = v3272
	var v3289 int32
	_ = v3289
	var v3293 int32
	_ = v3293
	var v3296 int32
	_ = v3296
	var v3299 int32
	_ = v3299
	var v3306 int32
	_ = v3306
	var v3309 int32
	_ = v3309
	var v3310 int32
	_ = v3310
	var v3315 int32
	_ = v3315
	var v3323 int32
	_ = v3323
	var v3324 int32
	_ = v3324
	var v3326 int32
	_ = v3326
	var v3331 int32
	_ = v3331
	var v3334 int32
	_ = v3334
	var v3337 int32
	_ = v3337
	var v3340 int32
	_ = v3340
	var v3341 int32
	_ = v3341
	var v3347 int32
	_ = v3347
	var v3351 int32
	_ = v3351
	var v3354 int32
	_ = v3354
	var v3359 int32
	_ = v3359
	var v3362 int32
	_ = v3362
	var v3365 int32
	_ = v3365
	var v3368 int32
	_ = v3368
	var v3369 int32
	_ = v3369
	var v3370 int32
	_ = v3370
	var v3372 int32
	_ = v3372
	var v3373 int32
	_ = v3373
	var v3374 int32
	_ = v3374
	var v3376 int32
	_ = v3376
	var v3378 int32
	_ = v3378
	var v3380 int32
	_ = v3380
	var v3384 int32
	_ = v3384
	var v3386 int32
	_ = v3386
	var v3389 int32
	_ = v3389
	var v3416 int32
	_ = v3416
	var v3419 int32
	_ = v3419
	var v3422 int32
	_ = v3422
	var v3429 int32
	_ = v3429
	var v3432 int32
	_ = v3432
	var v3435 int32
	_ = v3435
	var v3438 int32
	_ = v3438
	var v3439 int32
	_ = v3439
	var v3444 int32
	_ = v3444
	var v3447 int32
	_ = v3447
	var v3453 int32
	_ = v3453
	var v3456 int32
	_ = v3456
	var v3460 int32
	_ = v3460
	var v3462 int32
	_ = v3462
	var v3465 int32
	_ = v3465
	var v3474 int32
	_ = v3474
	var v3503 int32
	_ = v3503
	var v3506 int32
	_ = v3506
	var v3509 int32
	_ = v3509
	var v3513 int32
	_ = v3513
	var v3516 int32
	_ = v3516
	var v3518 int32
	_ = v3518
	var v3520 int32
	_ = v3520
	var v3521 int32
	_ = v3521
	var v3525 int32
	_ = v3525
	var v3532 int32
	_ = v3532
	var v3535 int64
	_ = v3535
	var v3541 int64
	_ = v3541
	var v3545 int32
	_ = v3545
	var v3546 int32
	_ = v3546
	v21 = m.G0
	v23 = v21 - int32(400)
	m.G0 = v23
	if l1 == int32(0) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v29 == int32(0) {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = int32(0)
	goto L1
L3:
	;
	m.G0 = v3545 + int32(400)
	return v3546
L4:
	;
	v33 = int32(0)
	v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)+140))
	if v34 == v33 {
		v3545 = v23
		v3546 = v33
		goto L3
	} else {
		goto L6
	}
L5:
	;
	v3545 = v23
	v3546 = int32(-1)
	goto L3
L6:
	;
	v37 = *(*int32)(unsafe.Add(mBase, uint32(l0)+168))
	if v37 != int32(-1) {
		goto L12
	} else {
		goto L13
	}
L7:
	;
	v3503 = *(*int32)(unsafe.Add(mBase, uint32(l0)+172))
	if v3503 == int32(0) {
		v3518 = v367
		goto L723
	} else {
		goto L724
	}
L8:
	;
	v3474 = m.G3
	m.Env.X__assert_fail(m, v3474+int32(_a_F_valkeyReaderGetReply_0), v3474+int32(_a_F_valkeyReaderGetReply_1), int32(280), v3474+int32(_a_F_valkeyReaderGetReply_2))
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L9:
	;
	v3465 = m.G3
	m.Env.X__assert_fail(m, v3465+int32(_a_F_valkeyReaderGetReply_3), v3465+int32(_a_F_valkeyReaderGetReply_1), int32(275), v3465+int32(_a_F_valkeyReaderGetReply_2))
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L10:
	;
	v3289 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v3289 == int32(0) {
		goto L666
	} else {
		goto L667
	}
L11:
	;
	v65 = v58
	goto L15
L12:
	;
	if v37 < int32(0) {
		v3271 = v23
		v3272 = v33
		goto L10
	} else {
		goto L14
	}
L13:
	;
	v40 = int32(0)
	v41 = *(*int32)(unsafe.Add(mBase, uint32(l0)+160))
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v41)))
	*(*int32)(unsafe.Add(mBase, uint32(v42)+24)) = v40
	*(*int64)(unsafe.Add(mBase, uint32(v42)+16)) = int64(4294967295)
	*(*int64)(unsafe.Add(mBase, uint32(v42)+8)) = int64(-1)
	*(*int32)(unsafe.Add(mBase, uint32(v42))) = int32(-1)
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v41)))
	v52 = *(*int32)(unsafe.Add(mBase, uint32(l0)+180))
	*(*int32)(unsafe.Add(mBase, uint32(v51)+28)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(l0)+168)) = v40
	v58 = v40
	goto L11
L14:
	;
	v58 = v37
	goto L11
L15:
	;
	v81 = *(*int32)(unsafe.Add(mBase, uint32(l0)+160))
	v85 = *(*int32)(unsafe.Add(mBase, uint32(v81+v65<<(uint(int32(2))%32))))
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v85)))
	if int32(-1) < v86 {
		goto L49
	} else {
		goto L50
	}
L16:
	;
	v3271 = v23
	v3272 = v33
	goto L10
L17:
	;
	v3266 = *(*int32)(unsafe.Add(mBase, uint32(l0)+168))
	if int32(-1) < v3266 {
		v65 = v3266
		goto L15
	} else {
		goto L665
	}
L18:
	;
	if v3206 <= v3211 {
		goto L8
	} else {
		goto L664
	}
L19:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+168)) = int32(-1)
	v3271 = v23
	v3272 = v33
	goto L10
L20:
	;
	if v3114 < int32(0) {
		goto L17
	} else {
		goto L656
	}
L21:
	;
	v3131 = *(*int32)(unsafe.Add(mBase, uint32(l0)+172))
	if v3131 == int32(0) {
		goto L650
	} else {
		goto L651
	}
L22:
	;
	if v3114 != 0 {
		goto L20
	} else {
		goto L649
	}
L23:
	;
	if v3098 == int32(0) {
		goto L21
	} else {
		goto L648
	}
L24:
	;
	if v483 == v368 {
		goto L614
	} else {
		goto L615
	}
L25:
	;
	if base.Ui32(v608) < base.Ui32(int32(326)) {
		goto L531
	} else {
		goto L532
	}
L26:
	;
	if v483 == v368 {
		goto L518
	} else {
		goto L519
	}
L27:
	;
	if v608 != int32(1) {
		goto L481
	} else {
		goto L482
	}
L28:
	;
	v2341 = m.G3
	m.Env.X__assert_fail(m, v2341+int32(_a_F_valkeyReaderGetReply_0), v2341+int32(_a_F_valkeyReaderGetReply_1), int32(280), v2341+int32(_a_F_valkeyReaderGetReply_2))
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L29:
	;
	v2332 = m.G3
	m.Env.X__assert_fail(m, v2332+int32(_a_F_valkeyReaderGetReply_3), v2332+int32(_a_F_valkeyReaderGetReply_1), int32(275), v2332+int32(_a_F_valkeyReaderGetReply_2))
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L30:
	;
	if v2292 <= v2297 {
		goto L28
	} else {
		goto L479
	}
L31:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+168)) = int32(-1)
	v3271 = v23
	v3272 = v33
	goto L10
L32:
	;
	if v2198 < int32(0) {
		goto L17
	} else {
		goto L471
	}
L33:
	;
	v2217 = *(*int32)(unsafe.Add(mBase, uint32(l0)+172))
	if v2217 == int32(0) {
		goto L465
	} else {
		goto L466
	}
L34:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+136)) = v2199 + v2204
	if v2198 != 0 {
		goto L32
	} else {
		goto L464
	}
L35:
	;
	if v2178 == int32(0) {
		goto L33
	} else {
		goto L463
	}
L36:
	;
	v2108 = base.I32_wrap_i64(v2104)
	v2111 = v1051 + v2108 + int32(4)
	if base.Ui32(v803) < base.Ui32(v2111+v804) {
		v3271 = v23
		v3272 = v33
		goto L10
	} else {
		goto L443
	}
L37:
	;
	if base.Ui64(int64(-4294967298)) < base.Ui64(v2015+int64(-4294967296)) {
		goto L431
	} else {
		goto L432
	}
L38:
	;
	v1953 = *(*int32)(unsafe.Add(mBase, uint32(l0)+172))
	if v1953 == int32(0) {
		v1968 = v808
		goto L425
	} else {
		goto L426
	}
L39:
	;
	if v1050 != int32(45) {
		goto L421
	} else {
		goto L422
	}
L40:
	;
	v1147 = *(*int32)(unsafe.Add(mBase, uint32(l0)+164))
	if v65 != v1147+int32(-1) {
		goto L280
	} else {
		goto L281
	}
L41:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v85))) = v1140
	goto L40
L42:
	;
	v1140 = int32(9)
	goto L41
L43:
	;
	v803 = *(*int32)(unsafe.Add(mBase, uint32(l0)+140))
	v804 = *(*int32)(unsafe.Add(mBase, uint32(l0)+136))
	v805 = v803 - v804
	if base.Ui32(v805) < base.Ui32(int32(2)) {
		v3271 = v23
		v3272 = v33
		goto L10
	} else {
		goto L202
	}
L44:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v85))) = v796
	v799 = v796
	goto L43
L45:
	;
	v796 = int32(14)
	goto L44
L46:
	;
	v362 = *(*int32)(unsafe.Add(mBase, uint32(l0)+140))
	v363 = *(*int32)(unsafe.Add(mBase, uint32(l0)+136))
	v364 = v362 - v363
	if base.Ui32(v364) < base.Ui32(int32(2)) {
		v3271 = v23
		v3272 = v33
		goto L10
	} else {
		goto L113
	}
L47:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v85))) = v355
	v361 = v355
	goto L46
L48:
	;
	v355 = int32(5)
	goto L47
L49:
	;
	if base.Ui32(int32(14)) < base.Ui32(v86) {
		goto L107
	} else {
		goto L108
	}
L50:
	;
	v89 = *(*int32)(unsafe.Add(mBase, uint32(l0)+140))
	v90 = *(*int32)(unsafe.Add(mBase, uint32(l0)+136))
	if v89 == v90 {
		v3271 = v23
		v3272 = v33
		goto L10
	} else {
		goto L51
	}
L51:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+136)) = v90 + int32(1)
	v95 = *(*int32)(unsafe.Add(mBase, uint32(l0)+132))
	if v95 == int32(0) {
		v3271 = v23
		v3272 = v33
		goto L10
	} else {
		goto L52
	}
L52:
	;
	v102 = int32(*(*int8)(unsafe.Add(mBase, uint32(v95+v90))))
	v104 = v102 & int32(255)
	switch v104 + int32(-35) {
	case 0:
		goto L56
	case 1:
		v796 = int32(1)
		goto L44
	case 2:
		goto L42
	default:
		goto L53
	case 5:
		goto L54
	case 7:
		v1140 = int32(2)
		goto L41
	case 8:
		goto L48
	case 9:
		goto L60
	case 10:
		v355 = int32(6)
		goto L47
	case 23:
		goto L61
	case 26:
		goto L45
	case 27:
		goto L55
	case 60:
		goto L59
	case 89:
		goto L58
	case 91:
		goto L57
	}
L53:
	;
	switch v104 + int32(-7) {
	case 0:
		goto L65
	case 1:
		goto L64
	case 2:
		goto L66
	case 3:
		goto L68
	case 4, 5, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26:
		goto L63
	case 6:
		goto L67
	case 27:
		goto L69
	default:
		goto L70
	}
L54:
	;
	v355 = int32(13)
	goto L47
L55:
	;
	v1140 = int32(12)
	goto L41
L56:
	;
	v355 = int32(8)
	goto L47
L57:
	;
	v1140 = int32(10)
	goto L41
L58:
	;
	v1140 = int32(11)
	goto L41
L59:
	;
	v355 = int32(4)
	goto L47
L60:
	;
	v355 = int32(7)
	goto L47
L61:
	;
	v355 = int32(3)
	goto L47
L62:
	;
	v224 = m.G3
	*(*int32)(unsafe.Add(mBase, uint32(v23))) = v23 + int32(392)
	v233 = F_snprintf(m, v23+int32(64), int32(128), v224+int32(_a_F_valkeyReaderGetReply_4), v23)
	mBase = m.M
	v234 = m.ExcPending
	if v234 != 0 {
		goto L72
	} else {
		goto L78
	}
L63:
	;
	if base.Ui32(int32(94)) < base.Ui32(v102+int32(-32)) {
		goto L74
	} else {
		goto L75
	}
L64:
	;
	v188 = m.G3
	v193 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v188)+uint32(_c_F_valkeyReaderGetReply[0]))))
	*(*uint8)(unsafe.Add(mBase, uint32(v23+int32(396)))) = uint8(v193)
	v195 = *(*int32)(unsafe.Add(mBase, uint32(v188)+uint32(_c_F_valkeyReaderGetReply[1])))
	*(*int32)(unsafe.Add(mBase, uint32(v23)+392)) = v195
	goto L62
L65:
	;
	v175 = m.G3
	v180 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v175)+uint32(_c_F_valkeyReaderGetReply[2]))))
	*(*uint8)(unsafe.Add(mBase, uint32(v23+int32(396)))) = uint8(v180)
	v182 = *(*int32)(unsafe.Add(mBase, uint32(v175)+uint32(_c_F_valkeyReaderGetReply[3])))
	*(*int32)(unsafe.Add(mBase, uint32(v23)+392)) = v182
	goto L62
L66:
	;
	v162 = m.G3
	v167 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v162)+uint32(_c_F_valkeyReaderGetReply[4]))))
	*(*uint8)(unsafe.Add(mBase, uint32(v23+int32(396)))) = uint8(v167)
	v169 = *(*int32)(unsafe.Add(mBase, uint32(v162)+uint32(_c_F_valkeyReaderGetReply[5])))
	*(*int32)(unsafe.Add(mBase, uint32(v23)+392)) = v169
	goto L62
L67:
	;
	v149 = m.G3
	v154 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v149)+uint32(_c_F_valkeyReaderGetReply[6]))))
	*(*uint8)(unsafe.Add(mBase, uint32(v23+int32(396)))) = uint8(v154)
	v156 = *(*int32)(unsafe.Add(mBase, uint32(v149)+uint32(_c_F_valkeyReaderGetReply[7])))
	*(*int32)(unsafe.Add(mBase, uint32(v23)+392)) = v156
	goto L62
L68:
	;
	v136 = m.G3
	v141 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v136)+uint32(_c_F_valkeyReaderGetReply[8]))))
	*(*uint8)(unsafe.Add(mBase, uint32(v23+int32(396)))) = uint8(v141)
	v143 = *(*int32)(unsafe.Add(mBase, uint32(v136)+uint32(_c_F_valkeyReaderGetReply[9])))
	*(*int32)(unsafe.Add(mBase, uint32(v23)+392)) = v143
	goto L62
L69:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+48)) = v102
	v123 = m.G3
	v128 = F_snprintf(m, v23+int32(392), int32(8), v123+int32(_a_F_valkeyReaderGetReply_5), v23+int32(48))
	mBase = m.M
	v131 = m.ExcPending
	if v131 != 0 {
		goto L72
	} else {
		goto L73
	}
L70:
	;
	if v104 != int32(92) {
		goto L63
	} else {
		goto L71
	}
L71:
	;
	goto L69
L72:
	;
	return int32(0)
L73:
	;
	goto L62
L74:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+32)) = v104
	v216 = m.G3
	v221 = F_snprintf(m, v23+int32(392), int32(8), v216+int32(_a_F_valkeyReaderGetReply_6), v23+int32(32))
	mBase = m.M
	v222 = m.ExcPending
	if v222 != 0 {
		goto L72
	} else {
		goto L77
	}
L75:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+16)) = v102
	v205 = m.G3
	v210 = F_snprintf(m, v23+int32(392), int32(8), v205+int32(_a_F_valkeyReaderGetReply_7), v23+int32(16))
	mBase = m.M
	v211 = m.ExcPending
	if v211 != 0 {
		goto L72
	} else {
		goto L76
	}
L76:
	;
	goto L62
L77:
	;
	goto L62
L78:
	;
	v235 = *(*int32)(unsafe.Add(mBase, uint32(l0)+172))
	if v235 == int32(0) {
		goto L79
	} else {
		goto L80
	}
L79:
	;
	v249 = *(*int32)(unsafe.Add(mBase, uint32(l0)+132))
	F_sdsfree(m, v249)
	mBase = m.M
	v251 = m.ExcPending
	if v251 != 0 {
		goto L72
	} else {
		goto L84
	}
L80:
	;
	v238 = *(*int32)(unsafe.Add(mBase, uint32(l0)+176))
	if v238 == int32(0) {
		goto L79
	} else {
		goto L81
	}
L81:
	;
	v241 = *(*int32)(unsafe.Add(mBase, uint32(v238)+24))
	if v241 == int32(0) {
		goto L79
	} else {
		goto L82
	}
L82:
	;
	m.T0[v241].(func(*base.Module, int32))(m, v235)
	mBase = m.M
	v245 = m.ExcPending
	if v245 != 0 {
		goto L72
	} else {
		goto L83
	}
L83:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+172)) = int32(0)
	goto L79
L84:
	;
	v252 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+140)) = v252
	*(*int64)(unsafe.Add(mBase, uint32(l0)+132)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+168)) = int32(-1)
	v258 = int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v258
	v261 = l0 + v258
	v263 = v23 + int32(64)
	if v263&int32(3) == v252 {
		v287 = v263
		goto L87
	} else {
		goto L88
	}
L85:
	;
	v321 = int32(127)
	if base.Ui32(v320) < base.Ui32(v321) {
		goto L101
	} else {
		goto L102
	}
L86:
	;
	v320 = v312 - v263
	goto L85
L87:
	;
	v291 = v287
	goto L95
L88:
	;
	v273 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v263))))
	if v273 != 0 {
		goto L89
	} else {
		goto L90
	}
L89:
	;
	v276 = v263
	goto L91
L90:
	;
	v320 = v263 - v263
	goto L85
L91:
	;
	v280 = v276 + int32(1)
	if v280&int32(3) == int32(0) {
		v287 = v280
		goto L87
	} else {
		goto L93
	}
L93:
	;
	v285 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v280))))
	if v285 != 0 {
		v276 = v280
		goto L91
	} else {
		goto L94
	}
L94:
	;
	v312 = v280
	goto L86
L95:
	;
	v297 = *(*int32)(unsafe.Add(mBase, uint32(v291)))
	v300 = int32(-2139062144)
	if (int32(16843008)-v297|v297)&v300 == v300 {
		v291 = v291 + int32(4)
		goto L95
	} else {
		goto L97
	}
L96:
	;
	v306 = v291
	goto L98
L97:
	;
	goto L96
L98:
	;
	v310 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v306))))
	if v310 != 0 {
		v306 = v306 + int32(1)
		goto L98
	} else {
		goto L100
	}
L99:
	;
	v312 = v306
	goto L86
L100:
	;
	goto L99
L101:
	;
	v324 = v320
	goto L103
L102:
	;
	v324 = v321
	goto L103
L103:
	;
	if v324 == int32(0) {
		v328 = v261
		goto L105
	} else {
		goto L106
	}
L104:
	;
	v330 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v328+v324))) = uint8(v330)
	v3271 = v23
	v3272 = v33
	goto L10
L105:
	;
	goto L104
L106:
	;
	v327 = F__emscripten_memcpy_bulkmem(m, v261, v263, v324)
	mBase = m.M
	v328 = v327
	goto L105
L107:
	;
	v345 = m.G3
	m.Env.X__assert_fail(m, v345+int32(_a_F_valkeyReaderGetReply_8), v345+int32(_a_F_valkeyReaderGetReply_1), int32(677), v345+int32(_a_F_valkeyReaderGetReply_9))
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L108:
	;
	v335 = int32(1) << (uint(v86) % 32)
	if v335&int32(8696) == int32(0) {
		goto L109
	} else {
		goto L110
	}
L109:
	;
	if v335&int32(7684) != 0 {
		goto L40
	} else {
		goto L111
	}
L110:
	;
	v361 = v86
	goto L46
L111:
	;
	if v335&int32(16386) != 0 {
		v799 = v86
		goto L43
	} else {
		goto L112
	}
L112:
	;
	goto L107
L113:
	;
	v367 = *(*int32)(unsafe.Add(mBase, uint32(l0)+132))
	v368 = v367 + v363
	v371 = v364 + int32(-1)
	v372 = int32(0)
	v375 = base.B2i32(v371 != v372)
	if v368&int32(3) == v372 {
		v401 = v368
		v403 = v371
		v404 = v375
		goto L117
	} else {
		goto L118
	}
L114:
	;
	if v474 == int32(0) {
		v3271 = v23
		v3272 = v33
		goto L10
	} else {
		goto L139
	}
L115:
	;
	v474 = int32(0)
	goto L114
L116:
	;
	v452 = v445
	v454 = v447
	goto L134
L117:
	;
	if v404 == int32(0) {
		goto L115
	} else {
		goto L125
	}
L118:
	;
	if v371 == int32(0) {
		v401 = v368
		v403 = v371
		v404 = v375
		goto L117
	} else {
		goto L119
	}
L119:
	;
	v384 = v368
	v386 = v371
	goto L120
L120:
	;
	v389 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v384))))
	if v389 == int32(13) {
		v445 = v384
		v447 = v386
		goto L116
	} else {
		goto L122
	}
L121:
	;
	v401 = v396
	v403 = v392
	v404 = v394
	goto L117
L122:
	;
	v392 = v386 + int32(-1)
	v393 = int32(0)
	v394 = base.B2i32(v392 != v393)
	v396 = v384 + int32(1)
	if v396&int32(3) == v393 {
		v401 = v396
		v403 = v392
		v404 = v394
		goto L117
	} else {
		goto L123
	}
L123:
	;
	if v392 != 0 {
		v384 = v396
		v386 = v392
		goto L120
	} else {
		goto L124
	}
L124:
	;
	goto L121
L125:
	;
	v408 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v401))))
	if v408 == int32(13) {
		v438 = v401
		v440 = v403
		goto L126
	} else {
		goto L127
	}
L126:
	;
	if v440 == int32(0) {
		goto L115
	} else {
		goto L133
	}
L127:
	;
	if base.Ui32(v403) < base.Ui32(int32(4)) {
		v438 = v401
		v440 = v403
		goto L126
	} else {
		goto L128
	}
L128:
	;
	v418 = v401
	v420 = v403
	goto L129
L129:
	;
	v424 = *(*int32)(unsafe.Add(mBase, uint32(v418)))
	v425 = v424 ^ int32(218959117)
	v428 = int32(-2139062144)
	if (int32(16843008)-v425|v425)&v428 != v428 {
		v445 = v418
		v447 = v420
		goto L116
	} else {
		goto L131
	}
L130:
	;
	v438 = v433
	v440 = v435
	goto L126
L131:
	;
	v433 = v418 + int32(4)
	v435 = v420 + int32(-4)
	if base.Ui32(int32(3)) < base.Ui32(v435) {
		v418 = v433
		v420 = v435
		goto L129
	} else {
		goto L132
	}
L132:
	;
	goto L130
L133:
	;
	v445 = v438
	v447 = v440
	goto L116
L134:
	;
	v457 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v452))))
	if v457 != int32(13) {
		goto L136
	} else {
		goto L137
	}
L135:
	;
	goto L115
L136:
	;
	v462 = v454 + int32(-1)
	if v462 != 0 {
		v452 = v452 + int32(1)
		v454 = v462
		goto L134
	} else {
		goto L138
	}
L137:
	;
	v474 = v452
	goto L114
L138:
	;
	goto L135
L139:
	;
	v482 = v371
	v483 = v474
	v486 = v368
	goto L141
L140:
	;
	v608 = v483 - v368
	*(*int32)(unsafe.Add(mBase, uint32(l0)+136)) = v363 + v608 + int32(2)
	if v367 == int32(0) {
		v3271 = v23
		v3272 = v33
		goto L10
	} else {
		goto L170
	}
L141:
	;
	v497 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v483)+1)))
	if v497 == int32(10) {
		goto L140
	} else {
		goto L143
	}
L143:
	;
	v501 = v483 + int32(1)
	v503 = v482 - v501 + v486
	v505 = int32(0)
	v508 = base.B2i32(v503 != v505)
	if v501&int32(3) == v505 {
		v534 = v501
		v536 = v503
		v537 = v508
		goto L147
	} else {
		goto L148
	}
L144:
	;
	if v607 != 0 {
		v482 = v503
		v483 = v607
		v486 = v501
		goto L141
	} else {
		goto L169
	}
L145:
	;
	v607 = int32(0)
	goto L144
L146:
	;
	v585 = v578
	v587 = v580
	goto L164
L147:
	;
	if v537 == int32(0) {
		goto L145
	} else {
		goto L155
	}
L148:
	;
	if v503 == int32(0) {
		v534 = v501
		v536 = v503
		v537 = v508
		goto L147
	} else {
		goto L149
	}
L149:
	;
	v517 = v501
	v519 = v503
	goto L150
L150:
	;
	v522 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v517))))
	if v522 == int32(13) {
		v578 = v517
		v580 = v519
		goto L146
	} else {
		goto L152
	}
L151:
	;
	v534 = v529
	v536 = v525
	v537 = v527
	goto L147
L152:
	;
	v525 = v519 + int32(-1)
	v526 = int32(0)
	v527 = base.B2i32(v525 != v526)
	v529 = v517 + int32(1)
	if v529&int32(3) == v526 {
		v534 = v529
		v536 = v525
		v537 = v527
		goto L147
	} else {
		goto L153
	}
L153:
	;
	if v525 != 0 {
		v517 = v529
		v519 = v525
		goto L150
	} else {
		goto L154
	}
L154:
	;
	goto L151
L155:
	;
	v541 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v534))))
	if v541 == int32(13) {
		v571 = v534
		v573 = v536
		goto L156
	} else {
		goto L157
	}
L156:
	;
	if v573 == int32(0) {
		goto L145
	} else {
		goto L163
	}
L157:
	;
	if base.Ui32(v536) < base.Ui32(int32(4)) {
		v571 = v534
		v573 = v536
		goto L156
	} else {
		goto L158
	}
L158:
	;
	v551 = v534
	v553 = v536
	goto L159
L159:
	;
	v557 = *(*int32)(unsafe.Add(mBase, uint32(v551)))
	v558 = v557 ^ int32(218959117)
	v561 = int32(-2139062144)
	if (int32(16843008)-v558|v558)&v561 != v561 {
		v578 = v551
		v580 = v553
		goto L146
	} else {
		goto L161
	}
L160:
	;
	v571 = v566
	v573 = v568
	goto L156
L161:
	;
	v566 = v551 + int32(4)
	v568 = v553 + int32(-4)
	if base.Ui32(int32(3)) < base.Ui32(v568) {
		v551 = v566
		v553 = v568
		goto L159
	} else {
		goto L162
	}
L162:
	;
	goto L160
L163:
	;
	v578 = v571
	v580 = v573
	goto L146
L164:
	;
	v590 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v585))))
	if v590 != int32(13) {
		goto L166
	} else {
		goto L167
	}
L165:
	;
	goto L145
L166:
	;
	v595 = v587 + int32(-1)
	if v595 != 0 {
		v585 = v585 + int32(1)
		v587 = v595
		goto L164
	} else {
		goto L168
	}
L167:
	;
	v607 = v585
	goto L144
L168:
	;
	goto L165
L169:
	;
	v3271 = v23
	v3272 = v33
	goto L10
L170:
	;
	switch v361 + int32(-3) {
	case 0:
		goto L24
	case 1:
		goto L26
	default:
		goto L171
	case 4:
		goto L25
	case 5:
		goto L27
	case 10:
		goto L172
	}
L171:
	;
	v691 = int32(0)
	if v608 <= v691 {
		goto L186
	} else {
		goto L187
	}
L172:
	;
	if v608 < int32(1) {
		goto L173
	} else {
		goto L174
	}
L173:
	;
	v682 = int32(13)
	v683 = *(*int32)(unsafe.Add(mBase, uint32(l0)+176))
	if v683 == int32(0) {
		v3114 = v65
		v3121 = v682
		goto L22
	} else {
		goto L183
	}
L174:
	;
	v619 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v368))))
	if v619 == int32(45) {
		goto L175
	} else {
		goto L176
	}
L175:
	;
	v628 = int32(1)
	if v608 == v628 {
		goto L173
	} else {
		goto L178
	}
L176:
	;
	if base.Ui32((v619+int32(-58))&int32(255)) < base.Ui32(int32(246)) {
		goto L7
	} else {
		goto L177
	}
L177:
	;
	goto L175
L178:
	;
	v637 = v628
	goto L179
L179:
	;
	v652 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v368+v637))))
	if base.Ui32((v652+int32(-58))&int32(255)) <= base.Ui32(int32(245)) {
		goto L7
	} else {
		goto L181
	}
L180:
	;
	goto L173
L181:
	;
	v660 = v637 + int32(1)
	if v660 != v608 {
		v637 = v660
		goto L179
	} else {
		goto L182
	}
L182:
	;
	goto L180
L183:
	;
	v686 = *(*int32)(unsafe.Add(mBase, uint32(v683)))
	if v686 == int32(0) {
		v3114 = v65
		v3121 = v682
		goto L22
	} else {
		goto L184
	}
L184:
	;
	v689 = m.T0[v686].(func(*base.Module, int32, int32, int32) int32)(m, v85, v368, v608)
	mBase = m.M
	v690 = m.ExcPending
	if v690 != 0 {
		goto L72
	} else {
		goto L185
	}
L185:
	;
	v3098 = v689
	goto L23
L186:
	;
	v787 = *(*int32)(unsafe.Add(mBase, uint32(l0)+176))
	if v787 == int32(0) {
		v3098 = v361
		goto L23
	} else {
		goto L199
	}
L187:
	;
	v700 = v691
	goto L188
L188:
	;
	v715 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v368+v700))))
	switch v715 + int32(-10) {
	case 0, 3:
		goto L191
	default:
		goto L190
	}
L189:
	;
	goto L186
L190:
	;
	v765 = v700 + int32(1)
	if v765 != v608 {
		v700 = v765
		goto L188
	} else {
		goto L198
	}
L191:
	;
	v718 = *(*int32)(unsafe.Add(mBase, uint32(l0)+172))
	if v718 == int32(0) {
		v733 = v367
		goto L192
	} else {
		goto L193
	}
L192:
	;
	F_sdsfree(m, v733)
	mBase = m.M
	v735 = m.ExcPending
	if v735 != 0 {
		goto L72
	} else {
		goto L197
	}
L193:
	;
	v721 = *(*int32)(unsafe.Add(mBase, uint32(l0)+176))
	if v721 == int32(0) {
		v733 = v367
		goto L192
	} else {
		goto L194
	}
L194:
	;
	v724 = *(*int32)(unsafe.Add(mBase, uint32(v721)+24))
	if v724 == int32(0) {
		v733 = v367
		goto L192
	} else {
		goto L195
	}
L195:
	;
	m.T0[v724].(func(*base.Module, int32))(m, v718)
	mBase = m.M
	v728 = m.ExcPending
	if v728 != 0 {
		goto L72
	} else {
		goto L196
	}
L196:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+172)) = int32(0)
	v731 = *(*int32)(unsafe.Add(mBase, uint32(l0)+132))
	v733 = v731
	goto L192
L197:
	;
	v736 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+140)) = v736
	*(*int64)(unsafe.Add(mBase, uint32(l0)+132)) = int64(0)
	v740 = int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+168)) = v740
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(4)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+27)) = uint8(v736)
	v747 = m.G3
	v750 = *(*int64)(unsafe.Add(mBase, uint32(v747)+uint32(_c_F_valkeyReaderGetReply[10])))
	*(*int64)(unsafe.Add(mBase, uint32(l0)+4)) = v750
	v756 = *(*int64)(unsafe.Add(mBase, uint32(v747)+uint32(_c_F_valkeyReaderGetReply[11])))
	*(*int64)(unsafe.Add(mBase, uint32(l0+int32(12)))) = v756
	v762 = *(*int64)(unsafe.Add(mBase, uint32(v747)+uint32(_c_F_valkeyReaderGetReply[12])))
	*(*int64)(unsafe.Add(mBase, uint32(l0+int32(19)))) = v762
	v3545 = v23
	v3546 = v740
	goto L3
L198:
	;
	goto L189
L199:
	;
	v790 = *(*int32)(unsafe.Add(mBase, uint32(v787)))
	if v790 == int32(0) {
		v3098 = v361
		goto L23
	} else {
		goto L200
	}
L200:
	;
	v793 = m.T0[v790].(func(*base.Module, int32, int32, int32) int32)(m, v85, v368, v608)
	mBase = m.M
	v794 = m.ExcPending
	if v794 != 0 {
		goto L72
	} else {
		goto L201
	}
L201:
	;
	v3098 = v793
	goto L23
L202:
	;
	v808 = *(*int32)(unsafe.Add(mBase, uint32(l0)+132))
	v809 = v808 + v804
	v812 = v805 + int32(-1)
	v813 = int32(0)
	v816 = base.B2i32(v812 != v813)
	if v809&int32(3) == v813 {
		v842 = v809
		v844 = v812
		v845 = v816
		goto L206
	} else {
		goto L207
	}
L203:
	;
	if v915 == int32(0) {
		v3271 = v23
		v3272 = v33
		goto L10
	} else {
		goto L228
	}
L204:
	;
	v915 = int32(0)
	goto L203
L205:
	;
	v893 = v886
	v895 = v888
	goto L223
L206:
	;
	if v845 == int32(0) {
		goto L204
	} else {
		goto L214
	}
L207:
	;
	if v812 == int32(0) {
		v842 = v809
		v844 = v812
		v845 = v816
		goto L206
	} else {
		goto L208
	}
L208:
	;
	v825 = v809
	v827 = v812
	goto L209
L209:
	;
	v830 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v825))))
	if v830 == int32(13) {
		v886 = v825
		v888 = v827
		goto L205
	} else {
		goto L211
	}
L210:
	;
	v842 = v837
	v844 = v833
	v845 = v835
	goto L206
L211:
	;
	v833 = v827 + int32(-1)
	v834 = int32(0)
	v835 = base.B2i32(v833 != v834)
	v837 = v825 + int32(1)
	if v837&int32(3) == v834 {
		v842 = v837
		v844 = v833
		v845 = v835
		goto L206
	} else {
		goto L212
	}
L212:
	;
	if v833 != 0 {
		v825 = v837
		v827 = v833
		goto L209
	} else {
		goto L213
	}
L213:
	;
	goto L210
L214:
	;
	v849 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v842))))
	if v849 == int32(13) {
		v879 = v842
		v881 = v844
		goto L215
	} else {
		goto L216
	}
L215:
	;
	if v881 == int32(0) {
		goto L204
	} else {
		goto L222
	}
L216:
	;
	if base.Ui32(v844) < base.Ui32(int32(4)) {
		v879 = v842
		v881 = v844
		goto L215
	} else {
		goto L217
	}
L217:
	;
	v859 = v842
	v861 = v844
	goto L218
L218:
	;
	v865 = *(*int32)(unsafe.Add(mBase, uint32(v859)))
	v866 = v865 ^ int32(218959117)
	v869 = int32(-2139062144)
	if (int32(16843008)-v866|v866)&v869 != v869 {
		v886 = v859
		v888 = v861
		goto L205
	} else {
		goto L220
	}
L219:
	;
	v879 = v874
	v881 = v876
	goto L215
L220:
	;
	v874 = v859 + int32(4)
	v876 = v861 + int32(-4)
	if base.Ui32(int32(3)) < base.Ui32(v876) {
		v859 = v874
		v861 = v876
		goto L218
	} else {
		goto L221
	}
L221:
	;
	goto L219
L222:
	;
	v886 = v879
	v888 = v881
	goto L205
L223:
	;
	v898 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v893))))
	if v898 != int32(13) {
		goto L225
	} else {
		goto L226
	}
L224:
	;
	goto L204
L225:
	;
	v903 = v895 + int32(-1)
	if v903 != 0 {
		v893 = v893 + int32(1)
		v895 = v903
		goto L223
	} else {
		goto L227
	}
L226:
	;
	v915 = v893
	goto L203
L227:
	;
	goto L224
L228:
	;
	v923 = v812
	v924 = v915
	v927 = v809
	goto L230
L229:
	;
	if v924 == v809 {
		goto L38
	} else {
		goto L259
	}
L230:
	;
	v938 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v924)+1)))
	if v938 == int32(10) {
		goto L229
	} else {
		goto L232
	}
L232:
	;
	v942 = v924 + int32(1)
	v944 = v923 - v942 + v927
	v946 = int32(0)
	v949 = base.B2i32(v944 != v946)
	if v942&int32(3) == v946 {
		v975 = v942
		v977 = v944
		v978 = v949
		goto L236
	} else {
		goto L237
	}
L233:
	;
	if v1048 != 0 {
		v923 = v944
		v924 = v1048
		v927 = v942
		goto L230
	} else {
		goto L258
	}
L234:
	;
	v1048 = int32(0)
	goto L233
L235:
	;
	v1026 = v1019
	v1028 = v1021
	goto L253
L236:
	;
	if v978 == int32(0) {
		goto L234
	} else {
		goto L244
	}
L237:
	;
	if v944 == int32(0) {
		v975 = v942
		v977 = v944
		v978 = v949
		goto L236
	} else {
		goto L238
	}
L238:
	;
	v958 = v942
	v960 = v944
	goto L239
L239:
	;
	v963 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v958))))
	if v963 == int32(13) {
		v1019 = v958
		v1021 = v960
		goto L235
	} else {
		goto L241
	}
L240:
	;
	v975 = v970
	v977 = v966
	v978 = v968
	goto L236
L241:
	;
	v966 = v960 + int32(-1)
	v967 = int32(0)
	v968 = base.B2i32(v966 != v967)
	v970 = v958 + int32(1)
	if v970&int32(3) == v967 {
		v975 = v970
		v977 = v966
		v978 = v968
		goto L236
	} else {
		goto L242
	}
L242:
	;
	if v966 != 0 {
		v958 = v970
		v960 = v966
		goto L239
	} else {
		goto L243
	}
L243:
	;
	goto L240
L244:
	;
	v982 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v975))))
	if v982 == int32(13) {
		v1012 = v975
		v1014 = v977
		goto L245
	} else {
		goto L246
	}
L245:
	;
	if v1014 == int32(0) {
		goto L234
	} else {
		goto L252
	}
L246:
	;
	if base.Ui32(v977) < base.Ui32(int32(4)) {
		v1012 = v975
		v1014 = v977
		goto L245
	} else {
		goto L247
	}
L247:
	;
	v992 = v975
	v994 = v977
	goto L248
L248:
	;
	v998 = *(*int32)(unsafe.Add(mBase, uint32(v992)))
	v999 = v998 ^ int32(218959117)
	v1002 = int32(-2139062144)
	if (int32(16843008)-v999|v999)&v1002 != v1002 {
		v1019 = v992
		v1021 = v994
		goto L235
	} else {
		goto L250
	}
L249:
	;
	v1012 = v1007
	v1014 = v1009
	goto L245
L250:
	;
	v1007 = v992 + int32(4)
	v1009 = v994 + int32(-4)
	if base.Ui32(int32(3)) < base.Ui32(v1009) {
		v992 = v1007
		v994 = v1009
		goto L248
	} else {
		goto L251
	}
L251:
	;
	goto L249
L252:
	;
	v1019 = v1012
	v1021 = v1014
	goto L235
L253:
	;
	v1031 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1026))))
	if v1031 != int32(13) {
		goto L255
	} else {
		goto L256
	}
L254:
	;
	goto L234
L255:
	;
	v1036 = v1028 + int32(-1)
	if v1036 != 0 {
		v1026 = v1026 + int32(1)
		v1028 = v1036
		goto L253
	} else {
		goto L257
	}
L256:
	;
	v1048 = v1026
	goto L233
L257:
	;
	goto L254
L258:
	;
	v3271 = v23
	v3272 = v33
	goto L10
L259:
	;
	v1050 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v809))))
	v1051 = v924 - v809
	if v1051 != int32(1) {
		goto L263
	} else {
		goto L264
	}
L260:
	;
	if v1050 != int32(48) {
		goto L38
	} else {
		goto L277
	}
L261:
	;
	v1097 = v1066
	v1101 = v1067
	v1108 = v1079
	goto L271
L262:
	;
	if base.Ui32(int32(8)) < base.Ui32((v1050+int32(-49))&int32(255)) {
		goto L260
	} else {
		goto L270
	}
L263:
	;
	v1059 = base.B2i32(v1050 != int32(45))
	if v1059 == int32(0) {
		goto L266
	} else {
		goto L267
	}
L264:
	;
	v1054 = int64(0)
	switch v1050 + int32(-45) {
	case 0:
		goto L38
	default:
		goto L262
	case 3:
		v2104 = v1054
		goto L36
	}
L265:
	;
	if base.Ui32(int32(9)) <= base.Ui32((v1068+int32(-49))&int32(255)) {
		goto L38
	} else {
		goto L268
	}
L266:
	;
	v1064 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v809)+1)))
	v1066 = v809 + int32(1)
	v1067 = int32(2)
	v1068 = v1064
	goto L265
L267:
	;
	v1066 = v809
	v1067 = int32(1)
	v1068 = v1050
	goto L265
L268:
	;
	v1079 = base.I64_extend_i32_u(v1068+int32(-48)) & int64(255)
	if base.Ui32(v1067) < base.Ui32(v1051) {
		goto L261
	} else {
		goto L269
	}
L269:
	;
	v1923 = v1079
	goto L39
L270:
	;
	v2015 = base.I64_extend_i32_u(v1050+int32(-48)) & int64(255)
	goto L37
L271:
	;
	v1112 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1097)+1)))
	if base.Ui32((v1112+int32(-58))&int32(255)) < base.Ui32(int32(246)) {
		goto L38
	} else {
		goto L273
	}
L273:
	;
	if base.Ui64(int64(1844674407370955161)) < base.Ui64(v1108) {
		goto L38
	} else {
		goto L274
	}
L274:
	;
	v1122 = v1108 * int64(10)
	v1127 = base.I64_extend_i32_u(v1112+int32(-48)) & int64(255)
	if base.Ui64(v1127^int64(-1)) < base.Ui64(v1122) {
		goto L38
	} else {
		goto L275
	}
L275:
	;
	v1131 = int32(1)
	v1133 = v1122 + v1127
	v1135 = v1101 + v1131
	if v1135 == v1051 {
		v1923 = v1133
		goto L39
	} else {
		goto L276
	}
L276:
	;
	v1097 = v1097 + v1131
	v1101 = v1135
	v1108 = v1133
	goto L271
L277:
	;
	v2104 = v1054
	goto L36
L278:
	;
	v1527 = v1343 - v1228
	*(*int32)(unsafe.Add(mBase, uint32(l0)+136)) = v1223 + v1527 + int32(2)
	if v1227 == int32(0) {
		v3271 = v23
		v3272 = v33
		goto L10
	} else {
		goto L352
	}
L279:
	;
	v1488 = *(*int32)(unsafe.Add(mBase, uint32(l0)+172))
	if v1488 == int32(0) {
		goto L346
	} else {
		goto L347
	}
L280:
	;
	v1222 = *(*int32)(unsafe.Add(mBase, uint32(l0)+140))
	v1223 = *(*int32)(unsafe.Add(mBase, uint32(l0)+136))
	v1224 = v1222 - v1223
	if base.Ui32(v1224) < base.Ui32(int32(2)) {
		v3271 = v23
		v3272 = v33
		goto L10
	} else {
		goto L290
	}
L281:
	;
	v1152 = v1147 + int32(9)
	v1155 = m.G4
	v1156 = *(*int32)(unsafe.Add(mBase, uint32(v1155)+8))
	v1157 = m.T0[v1156].(func(*base.Module, int32, int32) int32)(m, v81, v1152<<(uint(int32(2))%32))
	mBase = m.M
	v1158 = m.ExcPending
	if v1158 != 0 {
		goto L72
	} else {
		goto L282
	}
L282:
	;
	if v1157 == int32(0) {
		goto L279
	} else {
		goto L283
	}
L283:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+160)) = v1157
	v1162 = *(*int32)(unsafe.Add(mBase, uint32(l0)+164))
	if v1152 <= v1162 {
		goto L280
	} else {
		goto L284
	}
L284:
	;
	goto L285
L285:
	;
	v1186 = m.G4
	v1187 = *(*int32)(unsafe.Add(mBase, uint32(v1186)+4))
	v1188 = m.T0[v1187].(func(*base.Module, int32, int32) int32)(m, int32(1), int32(32))
	mBase = m.M
	v1189 = m.ExcPending
	if v1189 != 0 {
		goto L72
	} else {
		goto L287
	}
L286:
	;
	goto L280
L287:
	;
	v1190 = *(*int32)(unsafe.Add(mBase, uint32(l0)+160))
	v1191 = *(*int32)(unsafe.Add(mBase, uint32(l0)+164))
	*(*int32)(unsafe.Add(mBase, uint32(v1190+v1191<<(uint(int32(2))%32)))) = v1188
	if v1188 == int32(0) {
		goto L279
	} else {
		goto L288
	}
L288:
	;
	v1199 = v1191 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+164)) = v1199
	if v1199 < v1152 {
		goto L285
	} else {
		goto L289
	}
L289:
	;
	goto L286
L290:
	;
	v1227 = *(*int32)(unsafe.Add(mBase, uint32(l0)+132))
	v1228 = v1227 + v1223
	v1231 = v1224 + int32(-1)
	v1232 = int32(0)
	v1235 = base.B2i32(v1231 != v1232)
	if v1228&int32(3) == v1232 {
		v1261 = v1228
		v1263 = v1231
		v1264 = v1235
		goto L294
	} else {
		goto L295
	}
L291:
	;
	if v1334 == int32(0) {
		v3271 = v23
		v3272 = v33
		goto L10
	} else {
		goto L316
	}
L292:
	;
	v1334 = int32(0)
	goto L291
L293:
	;
	v1312 = v1305
	v1314 = v1307
	goto L311
L294:
	;
	if v1264 == int32(0) {
		goto L292
	} else {
		goto L302
	}
L295:
	;
	if v1231 == int32(0) {
		v1261 = v1228
		v1263 = v1231
		v1264 = v1235
		goto L294
	} else {
		goto L296
	}
L296:
	;
	v1244 = v1228
	v1246 = v1231
	goto L297
L297:
	;
	v1249 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1244))))
	if v1249 == int32(13) {
		v1305 = v1244
		v1307 = v1246
		goto L293
	} else {
		goto L299
	}
L298:
	;
	v1261 = v1256
	v1263 = v1252
	v1264 = v1254
	goto L294
L299:
	;
	v1252 = v1246 + int32(-1)
	v1253 = int32(0)
	v1254 = base.B2i32(v1252 != v1253)
	v1256 = v1244 + int32(1)
	if v1256&int32(3) == v1253 {
		v1261 = v1256
		v1263 = v1252
		v1264 = v1254
		goto L294
	} else {
		goto L300
	}
L300:
	;
	if v1252 != 0 {
		v1244 = v1256
		v1246 = v1252
		goto L297
	} else {
		goto L301
	}
L301:
	;
	goto L298
L302:
	;
	v1268 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1261))))
	if v1268 == int32(13) {
		v1298 = v1261
		v1300 = v1263
		goto L303
	} else {
		goto L304
	}
L303:
	;
	if v1300 == int32(0) {
		goto L292
	} else {
		goto L310
	}
L304:
	;
	if base.Ui32(v1263) < base.Ui32(int32(4)) {
		v1298 = v1261
		v1300 = v1263
		goto L303
	} else {
		goto L305
	}
L305:
	;
	v1278 = v1261
	v1280 = v1263
	goto L306
L306:
	;
	v1284 = *(*int32)(unsafe.Add(mBase, uint32(v1278)))
	v1285 = v1284 ^ int32(218959117)
	v1288 = int32(-2139062144)
	if (int32(16843008)-v1285|v1285)&v1288 != v1288 {
		v1305 = v1278
		v1307 = v1280
		goto L293
	} else {
		goto L308
	}
L307:
	;
	v1298 = v1293
	v1300 = v1295
	goto L303
L308:
	;
	v1293 = v1278 + int32(4)
	v1295 = v1280 + int32(-4)
	if base.Ui32(int32(3)) < base.Ui32(v1295) {
		v1278 = v1293
		v1280 = v1295
		goto L306
	} else {
		goto L309
	}
L309:
	;
	goto L307
L310:
	;
	v1305 = v1298
	v1307 = v1300
	goto L293
L311:
	;
	v1317 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1312))))
	if v1317 != int32(13) {
		goto L313
	} else {
		goto L314
	}
L312:
	;
	goto L292
L313:
	;
	v1322 = v1314 + int32(-1)
	if v1322 != 0 {
		v1312 = v1312 + int32(1)
		v1314 = v1322
		goto L311
	} else {
		goto L315
	}
L314:
	;
	v1334 = v1312
	goto L291
L315:
	;
	goto L312
L316:
	;
	v1342 = v1231
	v1343 = v1334
	v1346 = v1228
	goto L317
L317:
	;
	v1357 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1343)+1)))
	if v1357 == int32(10) {
		goto L278
	} else {
		goto L319
	}
L319:
	;
	v1361 = v1343 + int32(1)
	v1363 = v1342 - v1361 + v1346
	v1365 = int32(0)
	v1368 = base.B2i32(v1363 != v1365)
	if v1361&int32(3) == v1365 {
		v1394 = v1361
		v1396 = v1363
		v1397 = v1368
		goto L323
	} else {
		goto L324
	}
L320:
	;
	if v1467 != 0 {
		v1342 = v1363
		v1343 = v1467
		v1346 = v1361
		goto L317
	} else {
		goto L345
	}
L321:
	;
	v1467 = int32(0)
	goto L320
L322:
	;
	v1445 = v1438
	v1447 = v1440
	goto L340
L323:
	;
	if v1397 == int32(0) {
		goto L321
	} else {
		goto L331
	}
L324:
	;
	if v1363 == int32(0) {
		v1394 = v1361
		v1396 = v1363
		v1397 = v1368
		goto L323
	} else {
		goto L325
	}
L325:
	;
	v1377 = v1361
	v1379 = v1363
	goto L326
L326:
	;
	v1382 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1377))))
	if v1382 == int32(13) {
		v1438 = v1377
		v1440 = v1379
		goto L322
	} else {
		goto L328
	}
L327:
	;
	v1394 = v1389
	v1396 = v1385
	v1397 = v1387
	goto L323
L328:
	;
	v1385 = v1379 + int32(-1)
	v1386 = int32(0)
	v1387 = base.B2i32(v1385 != v1386)
	v1389 = v1377 + int32(1)
	if v1389&int32(3) == v1386 {
		v1394 = v1389
		v1396 = v1385
		v1397 = v1387
		goto L323
	} else {
		goto L329
	}
L329:
	;
	if v1385 != 0 {
		v1377 = v1389
		v1379 = v1385
		goto L326
	} else {
		goto L330
	}
L330:
	;
	goto L327
L331:
	;
	v1401 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1394))))
	if v1401 == int32(13) {
		v1431 = v1394
		v1433 = v1396
		goto L332
	} else {
		goto L333
	}
L332:
	;
	if v1433 == int32(0) {
		goto L321
	} else {
		goto L339
	}
L333:
	;
	if base.Ui32(v1396) < base.Ui32(int32(4)) {
		v1431 = v1394
		v1433 = v1396
		goto L332
	} else {
		goto L334
	}
L334:
	;
	v1411 = v1394
	v1413 = v1396
	goto L335
L335:
	;
	v1417 = *(*int32)(unsafe.Add(mBase, uint32(v1411)))
	v1418 = v1417 ^ int32(218959117)
	v1421 = int32(-2139062144)
	if (int32(16843008)-v1418|v1418)&v1421 != v1421 {
		v1438 = v1411
		v1440 = v1413
		goto L322
	} else {
		goto L337
	}
L336:
	;
	v1431 = v1426
	v1433 = v1428
	goto L332
L337:
	;
	v1426 = v1411 + int32(4)
	v1428 = v1413 + int32(-4)
	if base.Ui32(int32(3)) < base.Ui32(v1428) {
		v1411 = v1426
		v1413 = v1428
		goto L335
	} else {
		goto L338
	}
L338:
	;
	goto L336
L339:
	;
	v1438 = v1431
	v1440 = v1433
	goto L322
L340:
	;
	v1450 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1445))))
	if v1450 != int32(13) {
		goto L342
	} else {
		goto L343
	}
L341:
	;
	goto L321
L342:
	;
	v1455 = v1447 + int32(-1)
	if v1455 != 0 {
		v1445 = v1445 + int32(1)
		v1447 = v1455
		goto L340
	} else {
		goto L344
	}
L343:
	;
	v1467 = v1445
	goto L320
L344:
	;
	goto L341
L345:
	;
	v3271 = v23
	v3272 = v33
	goto L10
L346:
	;
	v1502 = *(*int32)(unsafe.Add(mBase, uint32(l0)+132))
	F_sdsfree(m, v1502)
	mBase = m.M
	v1504 = m.ExcPending
	if v1504 != 0 {
		goto L72
	} else {
		goto L351
	}
L347:
	;
	v1491 = *(*int32)(unsafe.Add(mBase, uint32(l0)+176))
	if v1491 == int32(0) {
		goto L346
	} else {
		goto L348
	}
L348:
	;
	v1494 = *(*int32)(unsafe.Add(mBase, uint32(v1491)+24))
	if v1494 == int32(0) {
		goto L346
	} else {
		goto L349
	}
L349:
	;
	m.T0[v1494].(func(*base.Module, int32))(m, v1488)
	mBase = m.M
	v1498 = m.ExcPending
	if v1498 != 0 {
		goto L72
	} else {
		goto L350
	}
L350:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+172)) = int32(0)
	goto L346
L351:
	;
	v1505 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+140)) = v1505
	*(*int64)(unsafe.Add(mBase, uint32(l0)+132)) = int64(0)
	v1509 = int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+168)) = v1509
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(5)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+17)) = uint8(v1505)
	v1516 = m.G3
	v1519 = *(*int64)(unsafe.Add(mBase, uint32(v1516)+uint32(_c_F_valkeyReaderGetReply[13])))
	*(*int64)(unsafe.Add(mBase, uint32(l0)+4)) = v1519
	v1525 = *(*int64)(unsafe.Add(mBase, uint32(v1516)+uint32(_c_F_valkeyReaderGetReply[14])))
	*(*int64)(unsafe.Add(mBase, uint32(l0+int32(9)))) = v1525
	v3545 = v23
	v3546 = v1509
	goto L3
L352:
	;
	if v1343 == v1228 {
		goto L356
	} else {
		goto L357
	}
L353:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+172)) = v1892
	goto L17
L354:
	;
	v1841 = *(*int32)(unsafe.Add(mBase, uint32(v85)))
	switch v1841 + int32(-9) {
	case 0, 2:
		goto L409
	default:
		v1846 = v1837
		goto L408
	}
L355:
	;
	if base.Ui64(v1734+int64(-4294967296)) < base.Ui64(int64(-4294967297)) {
		goto L389
	} else {
		goto L390
	}
L356:
	;
	v1672 = *(*int32)(unsafe.Add(mBase, uint32(l0)+172))
	if v1672 == int32(0) {
		v1687 = v1227
		goto L382
	} else {
		goto L383
	}
L357:
	;
	v1535 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1228))))
	if v1527 != int32(1) {
		goto L363
	} else {
		goto L364
	}
L358:
	;
	if v1535 != int32(45) {
		goto L378
	} else {
		goto L379
	}
L359:
	;
	v1622 = *(*int32)(unsafe.Add(mBase, uint32(l0)+168))
	v1830 = base.B2i32(v1622 == int32(0))
	v1837 = int64(0)
	goto L354
L360:
	;
	if v1535 != int32(48) {
		goto L356
	} else {
		goto L377
	}
L361:
	;
	v1580 = v1549
	v1581 = v1550
	v1591 = v1562
	goto L371
L362:
	;
	if base.Ui32(int32(8)) < base.Ui32((v1535+int32(-49))&int32(255)) {
		goto L360
	} else {
		goto L370
	}
L363:
	;
	v1542 = base.B2i32(v1535 != int32(45))
	if v1542 == int32(0) {
		goto L366
	} else {
		goto L367
	}
L364:
	;
	switch v1535 + int32(-45) {
	case 0:
		goto L356
	default:
		goto L362
	case 3:
		goto L359
	}
L365:
	;
	if base.Ui32(int32(9)) <= base.Ui32((v1551+int32(-49))&int32(255)) {
		goto L356
	} else {
		goto L368
	}
L366:
	;
	v1547 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1228)+1)))
	v1549 = int32(2)
	v1550 = v1228 + int32(1)
	v1551 = v1547
	goto L365
L367:
	;
	v1549 = int32(1)
	v1550 = v1228
	v1551 = v1535
	goto L365
L368:
	;
	v1562 = base.I64_extend_i32_u(v1551+int32(-48)) & int64(255)
	if base.Ui32(v1549) < base.Ui32(v1527) {
		goto L361
	} else {
		goto L369
	}
L369:
	;
	v1642 = v1562
	goto L358
L370:
	;
	v1734 = base.I64_extend_i32_u(v1535+int32(-48)) & int64(255)
	goto L355
L371:
	;
	v1595 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1581)+1)))
	if base.Ui32((v1595+int32(-58))&int32(255)) < base.Ui32(int32(246)) {
		goto L356
	} else {
		goto L373
	}
L373:
	;
	if base.Ui64(int64(1844674407370955161)) < base.Ui64(v1591) {
		goto L356
	} else {
		goto L374
	}
L374:
	;
	v1605 = v1591 * int64(10)
	v1610 = base.I64_extend_i32_u(v1595+int32(-48)) & int64(255)
	if base.Ui64(v1610^int64(-1)) < base.Ui64(v1605) {
		goto L356
	} else {
		goto L375
	}
L375:
	;
	v1614 = int32(1)
	v1616 = v1605 + v1610
	v1618 = v1580 + v1614
	if v1618 == v1527 {
		v1642 = v1616
		goto L358
	} else {
		goto L376
	}
L376:
	;
	v1580 = v1618
	v1581 = v1581 + v1614
	v1591 = v1616
	goto L371
L377:
	;
	goto L359
L378:
	;
	if int64(-1) < v1642 {
		v1734 = v1642
		goto L355
	} else {
		goto L381
	}
L379:
	;
	if base.Ui64(int64(-9223372036854775807-1)) < base.Ui64(v1642) {
		goto L356
	} else {
		goto L380
	}
L380:
	;
	v1734 = int64(0) - v1642
	goto L355
L381:
	;
	goto L356
L382:
	;
	F_sdsfree(m, v1687)
	mBase = m.M
	v1689 = m.ExcPending
	if v1689 != 0 {
		goto L72
	} else {
		goto L387
	}
L383:
	;
	v1675 = *(*int32)(unsafe.Add(mBase, uint32(l0)+176))
	if v1675 == int32(0) {
		v1687 = v1227
		goto L382
	} else {
		goto L384
	}
L384:
	;
	v1678 = *(*int32)(unsafe.Add(mBase, uint32(v1675)+24))
	if v1678 == int32(0) {
		v1687 = v1227
		goto L382
	} else {
		goto L385
	}
L385:
	;
	m.T0[v1678].(func(*base.Module, int32))(m, v1672)
	mBase = m.M
	v1682 = m.ExcPending
	if v1682 != 0 {
		goto L72
	} else {
		goto L386
	}
L386:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+172)) = int32(0)
	v1685 = *(*int32)(unsafe.Add(mBase, uint32(l0)+132))
	v1687 = v1685
	goto L382
L387:
	;
	v1690 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+140)) = v1690
	*(*int64)(unsafe.Add(mBase, uint32(l0)+132)) = int64(0)
	v1694 = int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+168)) = v1694
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(4)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+25)) = uint8(v1690)
	v1701 = m.G3
	v1704 = *(*int64)(unsafe.Add(mBase, uint32(v1701)+uint32(_c_F_valkeyReaderGetReply[15])))
	*(*int64)(unsafe.Add(mBase, uint32(l0)+4)) = v1704
	v1710 = *(*int64)(unsafe.Add(mBase, uint32(v1701)+uint32(_c_F_valkeyReaderGetReply[16])))
	*(*int64)(unsafe.Add(mBase, uint32(l0+int32(12)))) = v1710
	v1716 = *(*int64)(unsafe.Add(mBase, uint32(v1701)+uint32(_c_F_valkeyReaderGetReply[17])))
	*(*int64)(unsafe.Add(mBase, uint32(l0+int32(17)))) = v1716
	v3545 = v23
	v3546 = v1694
	goto L3
L388:
	;
	if v1734 == int64(-1) {
		goto L399
	} else {
		goto L400
	}
L389:
	;
	v1749 = *(*int32)(unsafe.Add(mBase, uint32(l0)+172))
	if v1749 == int32(0) {
		v1764 = v1227
		goto L393
	} else {
		goto L394
	}
L390:
	;
	v1742 = *(*int32)(unsafe.Add(mBase, uint32(l0)+168))
	v1743 = *(*int64)(unsafe.Add(mBase, uint32(l0)+152))
	if v1743 < int64(1) {
		goto L388
	} else {
		goto L391
	}
L391:
	;
	if v1734 <= v1743 {
		goto L388
	} else {
		goto L392
	}
L392:
	;
	goto L389
L393:
	;
	F_sdsfree(m, v1764)
	mBase = m.M
	v1766 = m.ExcPending
	if v1766 != 0 {
		goto L72
	} else {
		goto L398
	}
L394:
	;
	v1752 = *(*int32)(unsafe.Add(mBase, uint32(l0)+176))
	if v1752 == int32(0) {
		v1764 = v1227
		goto L393
	} else {
		goto L395
	}
L395:
	;
	v1755 = *(*int32)(unsafe.Add(mBase, uint32(v1752)+24))
	if v1755 == int32(0) {
		v1764 = v1227
		goto L393
	} else {
		goto L396
	}
L396:
	;
	m.T0[v1755].(func(*base.Module, int32))(m, v1749)
	mBase = m.M
	v1759 = m.ExcPending
	if v1759 != 0 {
		goto L72
	} else {
		goto L397
	}
L397:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+172)) = int32(0)
	v1762 = *(*int32)(unsafe.Add(mBase, uint32(l0)+132))
	v1764 = v1762
	goto L393
L398:
	;
	v1767 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+140)) = v1767
	*(*int64)(unsafe.Add(mBase, uint32(l0)+132)) = int64(0)
	v1771 = int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+168)) = v1771
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(4)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+34)) = uint8(v1767)
	v1778 = m.G3
	v1781 = *(*int64)(unsafe.Add(mBase, uint32(v1778)+uint32(_c_F_valkeyReaderGetReply[18])))
	*(*int64)(unsafe.Add(mBase, uint32(l0)+4)) = v1781
	v1787 = *(*int64)(unsafe.Add(mBase, uint32(v1778)+uint32(_c_F_valkeyReaderGetReply[19])))
	*(*int64)(unsafe.Add(mBase, uint32(l0+int32(12)))) = v1787
	v1793 = *(*int64)(unsafe.Add(mBase, uint32(v1778)+uint32(_c_F_valkeyReaderGetReply[20])))
	*(*int64)(unsafe.Add(mBase, uint32(l0+int32(20)))) = v1793
	v1799 = *(*int64)(unsafe.Add(mBase, uint32(v1778)+uint32(_c_F_valkeyReaderGetReply[21])))
	*(*int64)(unsafe.Add(mBase, uint32(l0+int32(26)))) = v1799
	v3545 = v23
	v3546 = v1771
	goto L3
L399:
	;
	v1805 = int32(4)
	v1806 = *(*int32)(unsafe.Add(mBase, uint32(l0)+176))
	if v1806 == int32(0) {
		v1816 = v1805
		goto L401
	} else {
		goto L402
	}
L400:
	;
	v1830 = base.B2i32(v1742 == int32(0))
	v1837 = v1734
	goto L354
L401:
	;
	F_moveToNextTask(m, l0)
	mBase = m.M
	if v1742 == int32(0) {
		v1892 = v1816
		goto L353
	} else {
		goto L407
	}
L402:
	;
	v1809 = *(*int32)(unsafe.Add(mBase, uint32(v1806)+16))
	if v1809 == int32(0) {
		v1816 = v1805
		goto L401
	} else {
		goto L403
	}
L403:
	;
	v1812 = m.T0[v1809].(func(*base.Module, int32) int32)(m, v85)
	mBase = m.M
	v1813 = m.ExcPending
	if v1813 != 0 {
		goto L72
	} else {
		goto L404
	}
L404:
	;
	if v1812 != 0 {
		v1816 = v1812
		goto L401
	} else {
		goto L405
	}
L405:
	;
	F_valkeyReaderSetErrorOOM(m, l0)
	mBase = m.M
	v1815 = m.ExcPending
	if v1815 != 0 {
		goto L72
	} else {
		goto L406
	}
L406:
	;
	v3271 = v23
	v3272 = v33
	goto L10
L407:
	;
	goto L17
L408:
	;
	v1847 = *(*int32)(unsafe.Add(mBase, uint32(l0)+176))
	if v1847 == int32(0) {
		v1857 = v1841
		goto L410
	} else {
		goto L411
	}
L409:
	;
	v1846 = v1837 << (uint(int64(1)) % 64)
	goto L408
L410:
	;
	if v1857 != 0 {
		goto L414
	} else {
		goto L415
	}
L411:
	;
	v1850 = *(*int32)(unsafe.Add(mBase, uint32(v1847)+4))
	if v1850 == int32(0) {
		v1857 = v1841
		goto L410
	} else {
		goto L412
	}
L412:
	;
	v1854 = m.T0[v1850].(func(*base.Module, int32, int32) int32)(m, v85, base.I32_wrap_i64(v1846))
	mBase = m.M
	v1855 = m.ExcPending
	if v1855 != 0 {
		goto L72
	} else {
		goto L413
	}
L413:
	;
	v1857 = v1854
	goto L410
L414:
	;
	if v1846 == int64(0) {
		goto L417
	} else {
		goto L418
	}
L415:
	;
	F_valkeyReaderSetErrorOOM(m, l0)
	mBase = m.M
	v1859 = m.ExcPending
	if v1859 != 0 {
		goto L72
	} else {
		goto L416
	}
L416:
	;
	v3271 = v23
	v3272 = v33
	goto L10
L417:
	;
	F_moveToNextTask(m, l0)
	mBase = m.M
	if v1830 == int32(0) {
		goto L17
	} else {
		goto L420
	}
L418:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v85)+20)) = v1857
	*(*int64)(unsafe.Add(mBase, uint32(v85)+8)) = v1846
	v1864 = *(*int32)(unsafe.Add(mBase, uint32(l0)+168))
	v1866 = v1864 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+168)) = v1866
	v1868 = *(*int32)(unsafe.Add(mBase, uint32(l0)+160))
	v1871 = v1868 + v1866<<(uint(int32(2))%32)
	v1872 = *(*int32)(unsafe.Add(mBase, uint32(v1871)))
	*(*int32)(unsafe.Add(mBase, uint32(v1872)+24)) = v85
	*(*int64)(unsafe.Add(mBase, uint32(v1872)+16)) = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v1872)+8)) = int64(-1)
	*(*int32)(unsafe.Add(mBase, uint32(v1872))) = int32(-1)
	v1880 = *(*int32)(unsafe.Add(mBase, uint32(v1871)))
	v1881 = *(*int32)(unsafe.Add(mBase, uint32(l0)+180))
	*(*int32)(unsafe.Add(mBase, uint32(v1880)+28)) = v1881
	if v1830 != 0 {
		v1892 = v1857
		goto L353
	} else {
		goto L419
	}
L419:
	;
	goto L17
L420:
	;
	v1892 = v1857
	goto L353
L421:
	;
	if int64(-1) < v1923 {
		v2015 = v1923
		goto L37
	} else {
		goto L424
	}
L422:
	;
	if base.Ui64(int64(-9223372036854775807-1)) < base.Ui64(v1923) {
		goto L38
	} else {
		goto L423
	}
L423:
	;
	v2015 = int64(0) - v1923
	goto L37
L424:
	;
	goto L38
L425:
	;
	F_sdsfree(m, v1968)
	mBase = m.M
	v1970 = m.ExcPending
	if v1970 != 0 {
		goto L72
	} else {
		goto L430
	}
L426:
	;
	v1956 = *(*int32)(unsafe.Add(mBase, uint32(l0)+176))
	if v1956 == int32(0) {
		v1968 = v808
		goto L425
	} else {
		goto L427
	}
L427:
	;
	v1959 = *(*int32)(unsafe.Add(mBase, uint32(v1956)+24))
	if v1959 == int32(0) {
		v1968 = v808
		goto L425
	} else {
		goto L428
	}
L428:
	;
	m.T0[v1959].(func(*base.Module, int32))(m, v1953)
	mBase = m.M
	v1963 = m.ExcPending
	if v1963 != 0 {
		goto L72
	} else {
		goto L429
	}
L429:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+172)) = int32(0)
	v1966 = *(*int32)(unsafe.Add(mBase, uint32(l0)+132))
	v1968 = v1966
	goto L425
L430:
	;
	v1971 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+140)) = v1971
	*(*int64)(unsafe.Add(mBase, uint32(l0)+132)) = int64(0)
	v1975 = int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+168)) = v1975
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(4)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+26)) = uint8(v1971)
	v1982 = m.G3
	v1985 = *(*int64)(unsafe.Add(mBase, uint32(v1982)+uint32(_c_F_valkeyReaderGetReply[22])))
	*(*int64)(unsafe.Add(mBase, uint32(l0)+4)) = v1985
	v1991 = *(*int64)(unsafe.Add(mBase, uint32(v1982)+uint32(_c_F_valkeyReaderGetReply[23])))
	*(*int64)(unsafe.Add(mBase, uint32(l0+int32(12)))) = v1991
	v1997 = *(*int64)(unsafe.Add(mBase, uint32(v1982)+uint32(_c_F_valkeyReaderGetReply[24])))
	*(*int64)(unsafe.Add(mBase, uint32(l0+int32(18)))) = v1997
	v3545 = v23
	v3546 = v1975
	goto L3
L431:
	;
	if v2015 != int64(-1) {
		v2104 = v2015
		goto L36
	} else {
		goto L439
	}
L432:
	;
	v2023 = *(*int32)(unsafe.Add(mBase, uint32(l0)+172))
	if v2023 == int32(0) {
		v2038 = v808
		goto L433
	} else {
		goto L434
	}
L433:
	;
	F_sdsfree(m, v2038)
	mBase = m.M
	v2040 = m.ExcPending
	if v2040 != 0 {
		goto L72
	} else {
		goto L438
	}
L434:
	;
	v2026 = *(*int32)(unsafe.Add(mBase, uint32(l0)+176))
	if v2026 == int32(0) {
		v2038 = v808
		goto L433
	} else {
		goto L435
	}
L435:
	;
	v2029 = *(*int32)(unsafe.Add(mBase, uint32(v2026)+24))
	if v2029 == int32(0) {
		v2038 = v808
		goto L433
	} else {
		goto L436
	}
L436:
	;
	m.T0[v2029].(func(*base.Module, int32))(m, v2023)
	mBase = m.M
	v2033 = m.ExcPending
	if v2033 != 0 {
		goto L72
	} else {
		goto L437
	}
L437:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+172)) = int32(0)
	v2036 = *(*int32)(unsafe.Add(mBase, uint32(l0)+132))
	v2038 = v2036
	goto L433
L438:
	;
	v2041 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+140)) = v2041
	*(*int64)(unsafe.Add(mBase, uint32(l0)+132)) = int64(0)
	v2045 = int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+168)) = v2045
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(4)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+35)) = uint8(v2041)
	v2052 = m.G3
	v2055 = *(*int64)(unsafe.Add(mBase, uint32(v2052)+uint32(_c_F_valkeyReaderGetReply[25])))
	*(*int64)(unsafe.Add(mBase, uint32(l0)+4)) = v2055
	v2061 = *(*int64)(unsafe.Add(mBase, uint32(v2052)+uint32(_c_F_valkeyReaderGetReply[26])))
	*(*int64)(unsafe.Add(mBase, uint32(l0+int32(12)))) = v2061
	v2067 = *(*int64)(unsafe.Add(mBase, uint32(v2052)+uint32(_c_F_valkeyReaderGetReply[27])))
	*(*int64)(unsafe.Add(mBase, uint32(l0+int32(20)))) = v2067
	v2073 = *(*int64)(unsafe.Add(mBase, uint32(v2052)+uint32(_c_F_valkeyReaderGetReply[28])))
	*(*int64)(unsafe.Add(mBase, uint32(l0+int32(27)))) = v2073
	v3545 = v23
	v3546 = v2045
	goto L3
L439:
	;
	v2078 = v1051 + int32(2)
	v2079 = int32(4)
	v2080 = *(*int32)(unsafe.Add(mBase, uint32(l0)+176))
	if v2080 == int32(0) {
		v2198 = v65
		v2199 = v2078
		v2202 = v2079
		v2204 = v804
		goto L34
	} else {
		goto L440
	}
L440:
	;
	v2083 = *(*int32)(unsafe.Add(mBase, uint32(v2080)+16))
	if v2083 == int32(0) {
		v2198 = v65
		v2199 = v2078
		v2202 = v2079
		v2204 = v804
		goto L34
	} else {
		goto L441
	}
L441:
	;
	v2086 = m.T0[v2083].(func(*base.Module, int32) int32)(m, v85)
	mBase = m.M
	v2087 = m.ExcPending
	if v2087 != 0 {
		goto L72
	} else {
		goto L442
	}
L442:
	;
	v2175 = v2078
	v2178 = v2086
	goto L35
L443:
	;
	v2115 = base.B2i32(v799 != int32(14))
	if v799 != int32(14) {
		goto L446
	} else {
		goto L447
	}
L444:
	;
	v2160 = *(*int32)(unsafe.Add(mBase, uint32(l0)+176))
	if v2160 == int32(0) {
		v2175 = v2111
		v2178 = v799
		goto L35
	} else {
		goto L460
	}
L445:
	;
	v2121 = *(*int32)(unsafe.Add(mBase, uint32(l0)+172))
	if v2121 == int32(0) {
		v2136 = v808
		goto L451
	} else {
		goto L452
	}
L446:
	;
	if v799 != int32(14) {
		goto L444
	} else {
		goto L449
	}
L447:
	;
	if base.Ui64(v2104) < base.Ui64(int64(4)) {
		goto L445
	} else {
		goto L448
	}
L448:
	;
	goto L446
L449:
	;
	v2118 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v924)+5)))
	if v2118 == int32(58) {
		goto L444
	} else {
		goto L450
	}
L450:
	;
	goto L445
L451:
	;
	F_sdsfree(m, v2136)
	mBase = m.M
	v2138 = m.ExcPending
	if v2138 != 0 {
		goto L72
	} else {
		goto L456
	}
L452:
	;
	v2124 = *(*int32)(unsafe.Add(mBase, uint32(l0)+176))
	if v2124 == int32(0) {
		v2136 = v808
		goto L451
	} else {
		goto L453
	}
L453:
	;
	v2127 = *(*int32)(unsafe.Add(mBase, uint32(v2124)+24))
	if v2127 == int32(0) {
		v2136 = v808
		goto L451
	} else {
		goto L454
	}
L454:
	;
	m.T0[v2127].(func(*base.Module, int32))(m, v2121)
	mBase = m.M
	v2131 = m.ExcPending
	if v2131 != 0 {
		goto L72
	} else {
		goto L455
	}
L455:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+172)) = int32(0)
	v2134 = *(*int32)(unsafe.Add(mBase, uint32(l0)+132))
	v2136 = v2134
	goto L451
L456:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+140)) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(l0)+132)) = int64(0)
	v2143 = int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+168)) = v2143
	v2146 = int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v2146
	v2150 = m.G3
	goto L459
L457:
	;
	v2158 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+79)) = uint8(v2158)
	v3545 = v23
	v3546 = v2143
	goto L3
L458:
	;
	goto L457
L459:
	;
	v2156 = F__emscripten_memcpy_bulkmem(m, l0+v2146, v2150+int32(_a_F_valkeyReaderGetReply_10), int32(75))
	mBase = m.M
	goto L458
L460:
	;
	v2163 = *(*int32)(unsafe.Add(mBase, uint32(v2160)))
	if v2163 == int32(0) {
		v2175 = v2111
		v2178 = v799
		goto L35
	} else {
		goto L461
	}
L461:
	;
	v2168 = m.T0[v2163].(func(*base.Module, int32, int32, int32) int32)(m, v85, v924+int32(2), v2108)
	mBase = m.M
	v2169 = m.ExcPending
	if v2169 != 0 {
		goto L72
	} else {
		goto L462
	}
L462:
	;
	v2175 = v2111
	v2178 = v2168
	goto L35
L463:
	;
	v2192 = *(*int32)(unsafe.Add(mBase, uint32(l0)+168))
	v2193 = *(*int32)(unsafe.Add(mBase, uint32(l0)+136))
	v2198 = v2192
	v2199 = v2175
	v2202 = v2178
	v2204 = v2193
	goto L34
L464:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+172)) = v2202
	goto L31
L465:
	;
	v2231 = *(*int32)(unsafe.Add(mBase, uint32(l0)+132))
	F_sdsfree(m, v2231)
	mBase = m.M
	v2233 = m.ExcPending
	if v2233 != 0 {
		goto L72
	} else {
		goto L470
	}
L466:
	;
	v2220 = *(*int32)(unsafe.Add(mBase, uint32(l0)+176))
	if v2220 == int32(0) {
		goto L465
	} else {
		goto L467
	}
L467:
	;
	v2223 = *(*int32)(unsafe.Add(mBase, uint32(v2220)+24))
	if v2223 == int32(0) {
		goto L465
	} else {
		goto L468
	}
L468:
	;
	m.T0[v2223].(func(*base.Module, int32))(m, v2217)
	mBase = m.M
	v2227 = m.ExcPending
	if v2227 != 0 {
		goto L72
	} else {
		goto L469
	}
L469:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+172)) = int32(0)
	goto L465
L470:
	;
	v2234 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+140)) = v2234
	*(*int64)(unsafe.Add(mBase, uint32(l0)+132)) = int64(0)
	v2238 = int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+168)) = v2238
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(5)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+17)) = uint8(v2234)
	v2245 = m.G3
	v2248 = *(*int64)(unsafe.Add(mBase, uint32(v2245)+uint32(_c_F_valkeyReaderGetReply[13])))
	*(*int64)(unsafe.Add(mBase, uint32(l0)+4)) = v2248
	v2254 = *(*int64)(unsafe.Add(mBase, uint32(v2245)+uint32(_c_F_valkeyReaderGetReply[14])))
	*(*int64)(unsafe.Add(mBase, uint32(l0+int32(9)))) = v2254
	v3545 = v23
	v3546 = v2238
	goto L3
L471:
	;
	v2258 = *(*int32)(unsafe.Add(mBase, uint32(l0)+160))
	v2263 = v2198
	goto L472
L472:
	;
	v2281 = v2258 + v2263<<(uint(int32(2))%32)
	v2284 = *(*int32)(unsafe.Add(mBase, uint32(v2281+int32(-4))))
	v2285 = *(*int32)(unsafe.Add(mBase, uint32(v2284)))
	if base.Ui32(v2285+int32(-9)) < base.Ui32(int32(4)) {
		goto L474
	} else {
		goto L475
	}
L473:
	;
	goto L31
L474:
	;
	v2292 = *(*int64)(unsafe.Add(mBase, uint32(v2284)+8))
	v2295 = *(*int32)(unsafe.Add(mBase, uint32(v2281)))
	v2296 = *(*int32)(unsafe.Add(mBase, uint32(v2295)+16))
	v2297 = base.I64_extend_i32_s(v2296)
	if v2292+int64(-1) != v2297 {
		goto L30
	} else {
		goto L477
	}
L475:
	;
	if v2285 != int32(2) {
		goto L29
	} else {
		goto L476
	}
L476:
	;
	goto L474
L477:
	;
	v2300 = v2263 + int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+168)) = v2300
	if v2300 != 0 {
		v2263 = v2300
		goto L472
	} else {
		goto L478
	}
L478:
	;
	goto L473
L479:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v2295)+8)) = int64(-1)
	*(*int32)(unsafe.Add(mBase, uint32(v2295))) = int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(v2295)+16)) = v2296 + int32(1)
	goto L17
L480:
	;
	v2501 = int32(8)
	v2502 = *(*int32)(unsafe.Add(mBase, uint32(l0)+176))
	if v2502 == int32(0) {
		v3114 = v65
		v3121 = v2501
		goto L22
	} else {
		goto L515
	}
L481:
	;
	v2461 = *(*int32)(unsafe.Add(mBase, uint32(l0)+172))
	if v2461 == int32(0) {
		v2476 = v367
		goto L509
	} else {
		goto L510
	}
L482:
	;
	v2352 = m.G3
	v2354 = v2352 + int32(_a_F_valkeyReaderGetReply_11)
	v2355 = int32(*(*int8)(unsafe.Add(mBase, uint32(v368))))
	v2356 = int32(5)
	if v2354&int32(3) == int32(0) {
		v2386 = v2354
		v2388 = v2356
		v2389 = int32(1)
		goto L486
	} else {
		goto L487
	}
L483:
	;
	if v2459 != 0 {
		goto L480
	} else {
		goto L508
	}
L484:
	;
	v2459 = int32(0)
	goto L483
L485:
	;
	v2437 = v2430
	v2439 = v2432
	goto L503
L486:
	;
	if v2389 == int32(0) {
		goto L484
	} else {
		goto L494
	}
L487:
	;
	goto L488
L488:
	;
	v2369 = v2354
	v2371 = v2356
	goto L489
L489:
	;
	v2374 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2369))))
	if v2374 == v2355&int32(255) {
		v2430 = v2369
		v2432 = v2371
		goto L485
	} else {
		goto L491
	}
L490:
	;
	v2386 = v2381
	v2388 = v2377
	v2389 = v2379
	goto L486
L491:
	;
	v2377 = v2371 + int32(-1)
	v2378 = int32(0)
	v2379 = base.B2i32(v2377 != v2378)
	v2381 = v2369 + int32(1)
	if v2381&int32(3) == v2378 {
		v2386 = v2381
		v2388 = v2377
		v2389 = v2379
		goto L486
	} else {
		goto L492
	}
L492:
	;
	if v2377 != 0 {
		v2369 = v2381
		v2371 = v2377
		goto L489
	} else {
		goto L493
	}
L493:
	;
	goto L490
L494:
	;
	v2393 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2386))))
	if v2393 == v2355&int32(255) {
		v2423 = v2386
		v2425 = v2388
		goto L495
	} else {
		goto L496
	}
L495:
	;
	if v2425 == int32(0) {
		goto L484
	} else {
		goto L502
	}
L496:
	;
	if base.Ui32(v2388) < base.Ui32(int32(4)) {
		v2423 = v2386
		v2425 = v2388
		goto L495
	} else {
		goto L497
	}
L497:
	;
	v2403 = v2386
	v2405 = v2388
	goto L498
L498:
	;
	v2409 = *(*int32)(unsafe.Add(mBase, uint32(v2403)))
	v2410 = v2409 ^ v2355&int32(255)*int32(16843009)
	v2413 = int32(-2139062144)
	if (int32(16843008)-v2410|v2410)&v2413 != v2413 {
		v2430 = v2403
		v2432 = v2405
		goto L485
	} else {
		goto L500
	}
L499:
	;
	v2423 = v2418
	v2425 = v2420
	goto L495
L500:
	;
	v2418 = v2403 + int32(4)
	v2420 = v2405 + int32(-4)
	if base.Ui32(int32(3)) < base.Ui32(v2420) {
		v2403 = v2418
		v2405 = v2420
		goto L498
	} else {
		goto L501
	}
L501:
	;
	goto L499
L502:
	;
	v2430 = v2423
	v2432 = v2425
	goto L485
L503:
	;
	v2442 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2437))))
	if v2442 != v2355&int32(255) {
		goto L505
	} else {
		goto L506
	}
L504:
	;
	goto L484
L505:
	;
	v2447 = v2439 + int32(-1)
	if v2447 != 0 {
		v2437 = v2437 + int32(1)
		v2439 = v2447
		goto L503
	} else {
		goto L507
	}
L506:
	;
	v2459 = v2437
	goto L483
L507:
	;
	goto L504
L508:
	;
	goto L481
L509:
	;
	F_sdsfree(m, v2476)
	mBase = m.M
	v2478 = m.ExcPending
	if v2478 != 0 {
		goto L72
	} else {
		goto L514
	}
L510:
	;
	v2464 = *(*int32)(unsafe.Add(mBase, uint32(l0)+176))
	if v2464 == int32(0) {
		v2476 = v367
		goto L509
	} else {
		goto L511
	}
L511:
	;
	v2467 = *(*int32)(unsafe.Add(mBase, uint32(v2464)+24))
	if v2467 == int32(0) {
		v2476 = v367
		goto L509
	} else {
		goto L512
	}
L512:
	;
	m.T0[v2467].(func(*base.Module, int32))(m, v2461)
	mBase = m.M
	v2471 = m.ExcPending
	if v2471 != 0 {
		goto L72
	} else {
		goto L513
	}
L513:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+172)) = int32(0)
	v2474 = *(*int32)(unsafe.Add(mBase, uint32(l0)+132))
	v2476 = v2474
	goto L509
L514:
	;
	v2479 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+140)) = v2479
	*(*int64)(unsafe.Add(mBase, uint32(l0)+132)) = int64(0)
	v2483 = int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+168)) = v2483
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(4)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+18)) = uint8(v2479)
	v2490 = m.G3
	v2493 = *(*int64)(unsafe.Add(mBase, uint32(v2490)+uint32(_c_F_valkeyReaderGetReply[29])))
	*(*int64)(unsafe.Add(mBase, uint32(l0)+4)) = v2493
	v2499 = *(*int64)(unsafe.Add(mBase, uint32(v2490)+uint32(_c_F_valkeyReaderGetReply[30])))
	*(*int64)(unsafe.Add(mBase, uint32(l0+int32(10)))) = v2499
	v3545 = v23
	v3546 = v2483
	goto L3
L515:
	;
	v2505 = *(*int32)(unsafe.Add(mBase, uint32(v2502)+20))
	if v2505 == int32(0) {
		v3114 = v65
		v3121 = v2501
		goto L22
	} else {
		goto L516
	}
L516:
	;
	v2512 = m.T0[v2505].(func(*base.Module, int32, int32) int32)(m, v85, base.B2i32(v2355&int32(-33) == int32(84)))
	mBase = m.M
	v2513 = m.ExcPending
	if v2513 != 0 {
		goto L72
	} else {
		goto L517
	}
L517:
	;
	v3098 = v2512
	goto L23
L518:
	;
	v2555 = int32(4)
	v2556 = *(*int32)(unsafe.Add(mBase, uint32(l0)+176))
	if v2556 == int32(0) {
		v3114 = v65
		v3121 = v2555
		goto L22
	} else {
		goto L526
	}
L519:
	;
	v2515 = *(*int32)(unsafe.Add(mBase, uint32(l0)+172))
	if v2515 == int32(0) {
		v2530 = v367
		goto L520
	} else {
		goto L521
	}
L520:
	;
	F_sdsfree(m, v2530)
	mBase = m.M
	v2532 = m.ExcPending
	if v2532 != 0 {
		goto L72
	} else {
		goto L525
	}
L521:
	;
	v2518 = *(*int32)(unsafe.Add(mBase, uint32(l0)+176))
	if v2518 == int32(0) {
		v2530 = v367
		goto L520
	} else {
		goto L522
	}
L522:
	;
	v2521 = *(*int32)(unsafe.Add(mBase, uint32(v2518)+24))
	if v2521 == int32(0) {
		v2530 = v367
		goto L520
	} else {
		goto L523
	}
L523:
	;
	m.T0[v2521].(func(*base.Module, int32))(m, v2515)
	mBase = m.M
	v2525 = m.ExcPending
	if v2525 != 0 {
		goto L72
	} else {
		goto L524
	}
L524:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+172)) = int32(0)
	v2528 = *(*int32)(unsafe.Add(mBase, uint32(l0)+132))
	v2530 = v2528
	goto L520
L525:
	;
	v2533 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+140)) = v2533
	*(*int64)(unsafe.Add(mBase, uint32(l0)+132)) = int64(0)
	v2537 = int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+168)) = v2537
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(4)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+17)) = uint8(v2533)
	v2544 = m.G3
	v2547 = *(*int64)(unsafe.Add(mBase, uint32(v2544)+uint32(_c_F_valkeyReaderGetReply[31])))
	*(*int64)(unsafe.Add(mBase, uint32(l0)+4)) = v2547
	v2553 = *(*int64)(unsafe.Add(mBase, uint32(v2544)+uint32(_c_F_valkeyReaderGetReply[32])))
	*(*int64)(unsafe.Add(mBase, uint32(l0+int32(9)))) = v2553
	v3545 = v23
	v3546 = v2537
	goto L3
L526:
	;
	v2559 = *(*int32)(unsafe.Add(mBase, uint32(v2556)+16))
	if v2559 == int32(0) {
		v3114 = v65
		v3121 = v2555
		goto L22
	} else {
		goto L527
	}
L527:
	;
	v2562 = m.T0[v2559].(func(*base.Module, int32) int32)(m, v85)
	mBase = m.M
	v2563 = m.ExcPending
	if v2563 != 0 {
		goto L72
	} else {
		goto L528
	}
L528:
	;
	v3098 = v2562
	goto L23
L529:
	;
	v2866 = int32(7)
	v2867 = *(*int32)(unsafe.Add(mBase, uint32(l0)+176))
	if v2867 == int32(0) {
		v3098 = v2866
		goto L23
	} else {
		goto L610
	}
L530:
	;
	v2863 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l0+v2859))) = uint8(v2863)
	v3271 = v23
	v3272 = v33
	goto L10
L531:
	;
	if v608 == int32(0) {
		goto L540
	} else {
		goto L541
	}
L532:
	;
	v2566 = *(*int32)(unsafe.Add(mBase, uint32(l0)+172))
	if v2566 == int32(0) {
		v2581 = v367
		goto L533
	} else {
		goto L534
	}
L533:
	;
	F_sdsfree(m, v2581)
	mBase = m.M
	v2583 = m.ExcPending
	if v2583 != 0 {
		goto L72
	} else {
		goto L538
	}
L534:
	;
	v2569 = *(*int32)(unsafe.Add(mBase, uint32(l0)+176))
	if v2569 == int32(0) {
		v2581 = v367
		goto L533
	} else {
		goto L535
	}
L535:
	;
	v2572 = *(*int32)(unsafe.Add(mBase, uint32(v2569)+24))
	if v2572 == int32(0) {
		v2581 = v367
		goto L533
	} else {
		goto L536
	}
L536:
	;
	m.T0[v2572].(func(*base.Module, int32))(m, v2566)
	mBase = m.M
	v2576 = m.ExcPending
	if v2576 != 0 {
		goto L72
	} else {
		goto L537
	}
L537:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+172)) = int32(0)
	v2579 = *(*int32)(unsafe.Add(mBase, uint32(l0)+132))
	v2581 = v2579
	goto L533
L538:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+140)) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(l0)+132)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+168)) = int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(4)
	v2592 = m.G3
	v2595 = *(*int64)(unsafe.Add(mBase, uint32(v2592)+uint32(_c_F_valkeyReaderGetReply[33])))
	*(*int64)(unsafe.Add(mBase, uint32(l0)+4)) = v2595
	v2601 = *(*int64)(unsafe.Add(mBase, uint32(v2592)+uint32(_c_F_valkeyReaderGetReply[34])))
	*(*int64)(unsafe.Add(mBase, uint32(l0+int32(12)))) = v2601
	v2607 = *(*int64)(unsafe.Add(mBase, uint32(v2592)+uint32(_c_F_valkeyReaderGetReply[35])))
	*(*int64)(unsafe.Add(mBase, uint32(l0+int32(20)))) = v2607
	v2613 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2592)+uint32(_c_F_valkeyReaderGetReply[36]))))
	*(*uint8)(unsafe.Add(mBase, uint32(l0+int32(28)))) = uint8(v2613)
	v2859 = int32(29)
	goto L530
L539:
	;
	v2624 = v23 + int32(64) + v608
	v2625 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v2624))) = uint8(v2625)
	switch v608 + int32(-3) {
	case 0:
		goto L546
	case 1:
		goto L545
	default:
		goto L542
	}
L540:
	;
	goto L539
L541:
	;
	v2620 = F__emscripten_memcpy_bulkmem(m, v23+int32(64), v368, v608)
	mBase = m.M
	goto L540
L542:
	;
	v2809 = F_strtod(m, v23+int32(64), v23+int32(392))
	mBase = m.M
	v2810 = m.ExcPending
	if v2810 != 0 {
		goto L72
	} else {
		goto L599
	}
L543:
	;
	v2762 = v23 + int32(64)
	v2763 = m.G3
	v2765 = v2763 + int32(_a_F_valkeyReaderGetReply_12)
	v2768 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2762))))
	if v2768 != 0 {
		goto L588
	} else {
		goto L589
	}
L544:
	;
	v2718 = v23 + int32(64)
	v2719 = m.G3
	v2721 = v2719 + int32(_a_F_valkeyReaderGetReply_13)
	v2724 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2718))))
	if v2724 != 0 {
		goto L575
	} else {
		goto L576
	}
L545:
	;
	v2674 = v23 + int32(64)
	v2675 = m.G3
	v2677 = v2675 + int32(_a_F_valkeyReaderGetReply_14)
	v2680 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2674))))
	if v2680 != 0 {
		goto L562
	} else {
		goto L563
	}
L546:
	;
	v2630 = v23 + int32(64)
	v2631 = m.G3
	v2633 = v2631 + int32(_a_F_valkeyReaderGetReply_15)
	v2636 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2630))))
	if v2636 != 0 {
		goto L549
	} else {
		goto L550
	}
L547:
	;
	if v2668-v2670 != 0 {
		goto L544
	} else {
		goto L559
	}
L548:
	;
	v2668 = F_tolower(m, v2664)
	mBase = m.M
	v2669 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2665))))
	v2670 = F_tolower(m, v2669)
	mBase = m.M
	goto L547
L549:
	;
	v2638 = v2630
	v2639 = v2633
	v2640 = v2636
	goto L552
L550:
	;
	v2664 = int32(0)
	v2665 = v2633
	goto L548
L551:
	;
	v2664 = v2661 & int32(255)
	v2665 = v2660
	goto L548
L552:
	;
	v2642 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2639))))
	if v2642 == int32(0) {
		v2660 = v2639
		v2661 = v2640
		goto L551
	} else {
		goto L554
	}
L553:
	;
	v2660 = v2654
	v2661 = int32(0)
	goto L551
L554:
	;
	v2646 = v2640 & int32(255)
	if v2646 == v2642 {
		goto L555
	} else {
		goto L556
	}
L555:
	;
	v2653 = int32(1)
	v2654 = v2639 + v2653
	v2655 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2638)+1)))
	if v2655 != 0 {
		v2638 = v2638 + v2653
		v2639 = v2654
		v2640 = v2655
		goto L552
	} else {
		goto L558
	}
L556:
	;
	v2648 = F_tolower(m, v2646)
	mBase = m.M
	v2649 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2639))))
	v2650 = F_tolower(m, v2649)
	mBase = m.M
	if v2648 == v2650 {
		goto L555
	} else {
		goto L557
	}
L557:
	;
	v2652 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2638))))
	v2660 = v2639
	v2661 = v2652
	goto L551
L558:
	;
	goto L553
L559:
	;
	v2865 = math.Float64frombits(uint64(0x7ff0000000000000))
	goto L529
L560:
	;
	if v2712-v2714 != 0 {
		goto L543
	} else {
		goto L572
	}
L561:
	;
	v2712 = F_tolower(m, v2708)
	mBase = m.M
	v2713 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2709))))
	v2714 = F_tolower(m, v2713)
	mBase = m.M
	goto L560
L562:
	;
	v2682 = v2674
	v2683 = v2677
	v2684 = v2680
	goto L565
L563:
	;
	v2708 = int32(0)
	v2709 = v2677
	goto L561
L564:
	;
	v2708 = v2705 & int32(255)
	v2709 = v2704
	goto L561
L565:
	;
	v2686 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2683))))
	if v2686 == int32(0) {
		v2704 = v2683
		v2705 = v2684
		goto L564
	} else {
		goto L567
	}
L566:
	;
	v2704 = v2698
	v2705 = int32(0)
	goto L564
L567:
	;
	v2690 = v2684 & int32(255)
	if v2690 == v2686 {
		goto L568
	} else {
		goto L569
	}
L568:
	;
	v2697 = int32(1)
	v2698 = v2683 + v2697
	v2699 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2682)+1)))
	if v2699 != 0 {
		v2682 = v2682 + v2697
		v2683 = v2698
		v2684 = v2699
		goto L565
	} else {
		goto L571
	}
L569:
	;
	v2692 = F_tolower(m, v2690)
	mBase = m.M
	v2693 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2683))))
	v2694 = F_tolower(m, v2693)
	mBase = m.M
	if v2692 == v2694 {
		goto L568
	} else {
		goto L570
	}
L570:
	;
	v2696 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2682))))
	v2704 = v2683
	v2705 = v2696
	goto L564
L571:
	;
	goto L566
L572:
	;
	v2865 = math.Float64frombits(uint64(0xfff0000000000000))
	goto L529
L573:
	;
	if v2756-v2758 != 0 {
		goto L542
	} else {
		goto L585
	}
L574:
	;
	v2756 = F_tolower(m, v2752)
	mBase = m.M
	v2757 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2753))))
	v2758 = F_tolower(m, v2757)
	mBase = m.M
	goto L573
L575:
	;
	v2726 = v2718
	v2727 = v2721
	v2728 = v2724
	goto L578
L576:
	;
	v2752 = int32(0)
	v2753 = v2721
	goto L574
L577:
	;
	v2752 = v2749 & int32(255)
	v2753 = v2748
	goto L574
L578:
	;
	v2730 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2727))))
	if v2730 == int32(0) {
		v2748 = v2727
		v2749 = v2728
		goto L577
	} else {
		goto L580
	}
L579:
	;
	v2748 = v2742
	v2749 = int32(0)
	goto L577
L580:
	;
	v2734 = v2728 & int32(255)
	if v2734 == v2730 {
		goto L581
	} else {
		goto L582
	}
L581:
	;
	v2741 = int32(1)
	v2742 = v2727 + v2741
	v2743 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2726)+1)))
	if v2743 != 0 {
		v2726 = v2726 + v2741
		v2727 = v2742
		v2728 = v2743
		goto L578
	} else {
		goto L584
	}
L582:
	;
	v2736 = F_tolower(m, v2734)
	mBase = m.M
	v2737 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2727))))
	v2738 = F_tolower(m, v2737)
	mBase = m.M
	if v2736 == v2738 {
		goto L581
	} else {
		goto L583
	}
L583:
	;
	v2740 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2726))))
	v2748 = v2727
	v2749 = v2740
	goto L577
L584:
	;
	goto L579
L585:
	;
	v2865 = math.Float64frombits(uint64(0x7ff8000000000000))
	goto L529
L586:
	;
	if v2800-v2802 != 0 {
		goto L542
	} else {
		goto L598
	}
L587:
	;
	v2800 = F_tolower(m, v2796)
	mBase = m.M
	v2801 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2797))))
	v2802 = F_tolower(m, v2801)
	mBase = m.M
	goto L586
L588:
	;
	v2770 = v2762
	v2771 = v2765
	v2772 = v2768
	goto L591
L589:
	;
	v2796 = int32(0)
	v2797 = v2765
	goto L587
L590:
	;
	v2796 = v2793 & int32(255)
	v2797 = v2792
	goto L587
L591:
	;
	v2774 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2771))))
	if v2774 == int32(0) {
		v2792 = v2771
		v2793 = v2772
		goto L590
	} else {
		goto L593
	}
L592:
	;
	v2792 = v2786
	v2793 = int32(0)
	goto L590
L593:
	;
	v2778 = v2772 & int32(255)
	if v2778 == v2774 {
		goto L594
	} else {
		goto L595
	}
L594:
	;
	v2785 = int32(1)
	v2786 = v2771 + v2785
	v2787 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2770)+1)))
	if v2787 != 0 {
		v2770 = v2770 + v2785
		v2771 = v2786
		v2772 = v2787
		goto L591
	} else {
		goto L597
	}
L595:
	;
	v2780 = F_tolower(m, v2778)
	mBase = m.M
	v2781 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2771))))
	v2782 = F_tolower(m, v2781)
	mBase = m.M
	if v2780 == v2782 {
		goto L594
	} else {
		goto L596
	}
L596:
	;
	v2784 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2770))))
	v2792 = v2771
	v2793 = v2784
	goto L590
L597:
	;
	goto L592
L598:
	;
	v2865 = math.Float64frombits(uint64(0x7ff8000000000000))
	goto L529
L599:
	;
	v2811 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23)+64)))
	if v2811 == int32(0) {
		goto L600
	} else {
		goto L601
	}
L600:
	;
	v2821 = *(*int32)(unsafe.Add(mBase, uint32(l0)+172))
	if v2821 == int32(0) {
		goto L604
	} else {
		goto L605
	}
L601:
	;
	v2814 = *(*int32)(unsafe.Add(mBase, uint32(v23)+392))
	if v2814 != v2624 {
		goto L600
	} else {
		goto L602
	}
L602:
	;
	if base.I64_reinterpret_f64(v2809)&int64(9223372036854775807) <= int64(9218868437227405311) {
		v2865 = v2809
		goto L529
	} else {
		goto L603
	}
L603:
	;
	goto L600
L604:
	;
	v2835 = *(*int32)(unsafe.Add(mBase, uint32(l0)+132))
	F_sdsfree(m, v2835)
	mBase = m.M
	v2837 = m.ExcPending
	if v2837 != 0 {
		goto L72
	} else {
		goto L609
	}
L605:
	;
	v2824 = *(*int32)(unsafe.Add(mBase, uint32(l0)+176))
	if v2824 == int32(0) {
		goto L604
	} else {
		goto L606
	}
L606:
	;
	v2827 = *(*int32)(unsafe.Add(mBase, uint32(v2824)+24))
	if v2827 == int32(0) {
		goto L604
	} else {
		goto L607
	}
L607:
	;
	m.T0[v2827].(func(*base.Module, int32))(m, v2821)
	mBase = m.M
	v2831 = m.ExcPending
	if v2831 != 0 {
		goto L72
	} else {
		goto L608
	}
L608:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+172)) = int32(0)
	goto L604
L609:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+140)) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(l0)+132)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+168)) = int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(4)
	v2846 = m.G3
	v2849 = *(*int64)(unsafe.Add(mBase, uint32(v2846)+uint32(_c_F_valkeyReaderGetReply[37])))
	*(*int64)(unsafe.Add(mBase, uint32(l0)+4)) = v2849
	v2855 = *(*int64)(unsafe.Add(mBase, uint32(v2846)+uint32(_c_F_valkeyReaderGetReply[38])))
	*(*int64)(unsafe.Add(mBase, uint32(l0+int32(12)))) = v2855
	v2859 = int32(20)
	goto L530
L610:
	;
	v2870 = *(*int32)(unsafe.Add(mBase, uint32(v2867)+12))
	if v2870 == int32(0) {
		v3098 = v2866
		goto L23
	} else {
		goto L611
	}
L611:
	;
	v2875 = m.T0[v2870].(func(*base.Module, int32, float64, int32, int32) int32)(m, v85, v2865, v23+int32(64), v608)
	mBase = m.M
	v2876 = m.ExcPending
	if v2876 != 0 {
		goto L72
	} else {
		goto L612
	}
L612:
	;
	v3098 = v2875
	goto L23
L613:
	;
	v3078 = int32(3)
	v3079 = *(*int32)(unsafe.Add(mBase, uint32(l0)+176))
	if v3079 == int32(0) {
		v3114 = v65
		v3121 = v3078
		goto L22
	} else {
		goto L645
	}
L614:
	;
	v3012 = *(*int32)(unsafe.Add(mBase, uint32(l0)+172))
	if v3012 == int32(0) {
		v3027 = v367
		goto L639
	} else {
		goto L640
	}
L615:
	;
	v2878 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v368))))
	if v608 != int32(1) {
		goto L620
	} else {
		goto L621
	}
L616:
	;
	if v2878 != int32(45) {
		goto L635
	} else {
		goto L636
	}
L617:
	;
	if v2878 != int32(48) {
		goto L614
	} else {
		goto L634
	}
L618:
	;
	v2925 = v2893
	v2928 = v2894
	v2935 = v2906
	goto L628
L619:
	;
	if base.Ui32(int32(8)) < base.Ui32((v2878+int32(-49))&int32(255)) {
		goto L617
	} else {
		goto L627
	}
L620:
	;
	v2886 = base.B2i32(v2878 != int32(45))
	if v2886 == int32(0) {
		goto L623
	} else {
		goto L624
	}
L621:
	;
	v2881 = int64(0)
	switch v2878 + int32(-45) {
	case 0:
		goto L614
	default:
		goto L619
	case 3:
		v3074 = v2881
		goto L613
	}
L622:
	;
	if base.Ui32(int32(9)) <= base.Ui32((v2895+int32(-49))&int32(255)) {
		goto L614
	} else {
		goto L625
	}
L623:
	;
	v2891 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v368)+1)))
	v2893 = v368 + int32(1)
	v2894 = int32(2)
	v2895 = v2891
	goto L622
L624:
	;
	v2893 = v368
	v2894 = int32(1)
	v2895 = v2878
	goto L622
L625:
	;
	v2906 = base.I64_extend_i32_u(v2895+int32(-48)) & int64(255)
	if base.Ui32(v2894) < base.Ui32(v608) {
		goto L618
	} else {
		goto L626
	}
L626:
	;
	v2982 = v2906
	goto L616
L627:
	;
	v3074 = base.I64_extend_i32_u(v2878+int32(-48)) & int64(255)
	goto L613
L628:
	;
	v2939 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2925)+1)))
	if base.Ui32((v2939+int32(-58))&int32(255)) < base.Ui32(int32(246)) {
		goto L614
	} else {
		goto L630
	}
L630:
	;
	if base.Ui64(int64(1844674407370955161)) < base.Ui64(v2935) {
		goto L614
	} else {
		goto L631
	}
L631:
	;
	v2949 = v2935 * int64(10)
	v2954 = base.I64_extend_i32_u(v2939+int32(-48)) & int64(255)
	if base.Ui64(v2954^int64(-1)) < base.Ui64(v2949) {
		goto L614
	} else {
		goto L632
	}
L632:
	;
	v2958 = int32(1)
	v2960 = v2949 + v2954
	v2962 = v2928 + v2958
	if v2962 == v608 {
		v2982 = v2960
		goto L616
	} else {
		goto L633
	}
L633:
	;
	v2925 = v2925 + v2958
	v2928 = v2962
	v2935 = v2960
	goto L628
L634:
	;
	v3074 = v2881
	goto L613
L635:
	;
	if int64(-1) < v2982 {
		v3074 = v2982
		goto L613
	} else {
		goto L638
	}
L636:
	;
	if base.Ui64(int64(-9223372036854775807-1)) < base.Ui64(v2982) {
		goto L614
	} else {
		goto L637
	}
L637:
	;
	v3074 = int64(0) - v2982
	goto L613
L638:
	;
	goto L614
L639:
	;
	F_sdsfree(m, v3027)
	mBase = m.M
	v3029 = m.ExcPending
	if v3029 != 0 {
		goto L72
	} else {
		goto L644
	}
L640:
	;
	v3015 = *(*int32)(unsafe.Add(mBase, uint32(l0)+176))
	if v3015 == int32(0) {
		v3027 = v367
		goto L639
	} else {
		goto L641
	}
L641:
	;
	v3018 = *(*int32)(unsafe.Add(mBase, uint32(v3015)+24))
	if v3018 == int32(0) {
		v3027 = v367
		goto L639
	} else {
		goto L642
	}
L642:
	;
	m.T0[v3018].(func(*base.Module, int32))(m, v3012)
	mBase = m.M
	v3022 = m.ExcPending
	if v3022 != 0 {
		goto L72
	} else {
		goto L643
	}
L643:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+172)) = int32(0)
	v3025 = *(*int32)(unsafe.Add(mBase, uint32(l0)+132))
	v3027 = v3025
	goto L639
L644:
	;
	v3030 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+140)) = v3030
	*(*int64)(unsafe.Add(mBase, uint32(l0)+132)) = int64(0)
	v3034 = int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+168)) = v3034
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(4)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+21)) = uint8(v3030)
	v3041 = m.G3
	v3044 = *(*int64)(unsafe.Add(mBase, uint32(v3041)+uint32(_c_F_valkeyReaderGetReply[39])))
	*(*int64)(unsafe.Add(mBase, uint32(l0)+4)) = v3044
	v3050 = *(*int64)(unsafe.Add(mBase, uint32(v3041)+uint32(_c_F_valkeyReaderGetReply[40])))
	*(*int64)(unsafe.Add(mBase, uint32(l0+int32(12)))) = v3050
	v3056 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3041)+uint32(_c_F_valkeyReaderGetReply[41]))))
	*(*uint8)(unsafe.Add(mBase, uint32(l0+int32(20)))) = uint8(v3056)
	v3545 = v23
	v3546 = v3034
	goto L3
L645:
	;
	v3082 = *(*int32)(unsafe.Add(mBase, uint32(v3079)+8))
	if v3082 == int32(0) {
		v3114 = v65
		v3121 = v3078
		goto L22
	} else {
		goto L646
	}
L646:
	;
	v3085 = m.T0[v3082].(func(*base.Module, int32, int64) int32)(m, v85, v3074)
	mBase = m.M
	v3086 = m.ExcPending
	if v3086 != 0 {
		goto L72
	} else {
		goto L647
	}
L647:
	;
	v3098 = v3085
	goto L23
L648:
	;
	v3109 = *(*int32)(unsafe.Add(mBase, uint32(l0)+168))
	v3114 = v3109
	v3121 = v3098
	goto L22
L649:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+172)) = v3121
	goto L19
L650:
	;
	v3145 = *(*int32)(unsafe.Add(mBase, uint32(l0)+132))
	F_sdsfree(m, v3145)
	mBase = m.M
	v3147 = m.ExcPending
	if v3147 != 0 {
		goto L72
	} else {
		goto L655
	}
L651:
	;
	v3134 = *(*int32)(unsafe.Add(mBase, uint32(l0)+176))
	if v3134 == int32(0) {
		goto L650
	} else {
		goto L652
	}
L652:
	;
	v3137 = *(*int32)(unsafe.Add(mBase, uint32(v3134)+24))
	if v3137 == int32(0) {
		goto L650
	} else {
		goto L653
	}
L653:
	;
	m.T0[v3137].(func(*base.Module, int32))(m, v3131)
	mBase = m.M
	v3141 = m.ExcPending
	if v3141 != 0 {
		goto L72
	} else {
		goto L654
	}
L654:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+172)) = int32(0)
	goto L650
L655:
	;
	v3148 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+140)) = v3148
	*(*int64)(unsafe.Add(mBase, uint32(l0)+132)) = int64(0)
	v3152 = int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+168)) = v3152
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(5)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+17)) = uint8(v3148)
	v3159 = m.G3
	v3162 = *(*int64)(unsafe.Add(mBase, uint32(v3159)+uint32(_c_F_valkeyReaderGetReply[13])))
	*(*int64)(unsafe.Add(mBase, uint32(l0)+4)) = v3162
	v3168 = *(*int64)(unsafe.Add(mBase, uint32(v3159)+uint32(_c_F_valkeyReaderGetReply[14])))
	*(*int64)(unsafe.Add(mBase, uint32(l0+int32(9)))) = v3168
	v3545 = v23
	v3546 = v3152
	goto L3
L656:
	;
	v3172 = *(*int32)(unsafe.Add(mBase, uint32(l0)+160))
	v3177 = v3114
	goto L657
L657:
	;
	v3195 = v3172 + v3177<<(uint(int32(2))%32)
	v3198 = *(*int32)(unsafe.Add(mBase, uint32(v3195+int32(-4))))
	v3199 = *(*int32)(unsafe.Add(mBase, uint32(v3198)))
	if base.Ui32(v3199+int32(-9)) < base.Ui32(int32(4)) {
		goto L659
	} else {
		goto L660
	}
L658:
	;
	goto L19
L659:
	;
	v3206 = *(*int64)(unsafe.Add(mBase, uint32(v3198)+8))
	v3209 = *(*int32)(unsafe.Add(mBase, uint32(v3195)))
	v3210 = *(*int32)(unsafe.Add(mBase, uint32(v3209)+16))
	v3211 = base.I64_extend_i32_s(v3210)
	if v3206+int64(-1) != v3211 {
		goto L18
	} else {
		goto L662
	}
L660:
	;
	if v3199 != int32(2) {
		goto L9
	} else {
		goto L661
	}
L661:
	;
	goto L659
L662:
	;
	v3214 = v3177 + int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+168)) = v3214
	if v3214 != 0 {
		v3177 = v3214
		goto L657
	} else {
		goto L663
	}
L663:
	;
	goto L658
L664:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v3209)+8)) = int64(-1)
	*(*int32)(unsafe.Add(mBase, uint32(v3209))) = int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(v3209)+16)) = v3210 + int32(1)
	goto L17
L665:
	;
	goto L16
L666:
	;
	v3293 = *(*int32)(unsafe.Add(mBase, uint32(l0)+136))
	if base.Ui32(v3293) < base.Ui32(int32(1024)) {
		goto L668
	} else {
		goto L669
	}
L667:
	;
	v3545 = v3271
	v3546 = int32(-1)
	goto L3
L668:
	;
	v3444 = *(*int32)(unsafe.Add(mBase, uint32(l0)+168))
	if v3444 != int32(-1) {
		v3545 = v3271
		v3546 = v3272
		goto L3
	} else {
		goto L715
	}
L669:
	;
	v3296 = *(*int32)(unsafe.Add(mBase, uint32(l0)+132))
	v3299 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3296+int32(-1)))))
	switch v3299&int32(7) + int32(-3) {
	case 0:
		goto L673
	case 1:
		goto L672
	default:
		goto L670
	}
L670:
	;
	v3315 = int32(-1)
	v3323 = v3296 + v3315
	v3324 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3323))))
	v3326 = v3324 & int32(7)
	switch v3326 {
	case 0:
		goto L682
	case 1:
		goto L681
	case 2:
		goto L680
	case 3:
		goto L679
	case 4:
		goto L678
	default:
		goto L676
	}
L671:
	;
	if int32(0) <= v3310 {
		goto L670
	} else {
		goto L674
	}
L672:
	;
	v3309 = *(*int32)(unsafe.Add(mBase, uint32(v3296+int32(-17))))
	v3310 = v3309
	goto L671
L673:
	;
	v3306 = *(*int32)(unsafe.Add(mBase, uint32(v3296+int32(-9))))
	v3310 = v3306
	goto L671
L674:
	;
	v3545 = v3271
	v3546 = int32(-1)
	goto L3
L675:
	;
	v3416 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+136)) = v3416
	v3419 = *(*int32)(unsafe.Add(mBase, uint32(l0)+132))
	v3422 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3419+int32(-1)))))
	switch v3422 & int32(7) {
	case 0:
		goto L714
	case 1:
		goto L713
	case 2:
		goto L712
	case 3:
		goto L711
	case 4:
		goto L710
	default:
		v3439 = v3416
		goto L709
	}
L676:
	;
	goto L675
L677:
	;
	if v3341 == int32(0) {
		goto L676
	} else {
		goto L683
	}
L678:
	;
	v3340 = *(*int32)(unsafe.Add(mBase, uint32(v3296+int32(-17))))
	v3341 = v3340
	goto L677
L679:
	;
	v3337 = *(*int32)(unsafe.Add(mBase, uint32(v3296+int32(-9))))
	v3341 = v3337
	goto L677
L680:
	;
	v3334 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v3296+int32(-5)))))
	v3341 = v3334
	goto L677
L681:
	;
	v3331 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3296+int32(-3)))))
	v3341 = v3331
	goto L677
L682:
	;
	v3341 = int32(base.Ui32(v3324) >> (uint(int32(3)) % 32))
	goto L677
L683:
	;
	v3347 = int32(-1)&v3341 + v3315
	v3351 = v3293>>(uint(int32(31))%32)&v3341 + v3293
	v3354 = v3347 - v3351 + int32(1)
	switch v3326 {
	default:
		goto L689
	case 1:
		goto L688
	case 2:
		goto L687
	case 3:
		goto L686
	case 4:
		goto L685
	}
L684:
	;
	v3370 = int32(0)
	v3372 = base.B2i32(base.Ui32(v3351) < base.Ui32(v3369))
	if base.Ui32(v3351) < base.Ui32(v3369) {
		goto L691
	} else {
		goto L692
	}
L685:
	;
	v3368 = *(*int32)(unsafe.Add(mBase, uint32(v3296+int32(-17))))
	v3369 = v3368
	goto L684
L686:
	;
	v3365 = *(*int32)(unsafe.Add(mBase, uint32(v3296+int32(-9))))
	v3369 = v3365
	goto L684
L687:
	;
	v3362 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v3296+int32(-5)))))
	v3369 = v3362
	goto L684
L688:
	;
	v3359 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3296+int32(-3)))))
	v3369 = v3359
	goto L684
L689:
	;
	v3369 = int32(base.Ui32(v3324) >> (uint(int32(3)) % 32))
	goto L684
L690:
	;
	v3386 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v3296+v3380))) = uint8(v3386)
	switch v3326 {
	default:
		goto L708
	case 1:
		goto L707
	case 2:
		goto L706
	case 3:
		goto L705
	case 4:
		goto L704
	}
L691:
	;
	v3373 = v3351
	goto L693
L692:
	;
	v3373 = v3370
	goto L693
L693:
	;
	v3374 = v3369 - v3373
	if base.Ui32(v3354) < base.Ui32(v3374) {
		goto L694
	} else {
		goto L695
	}
L694:
	;
	v3376 = v3354
	goto L696
L695:
	;
	v3376 = v3374
	goto L696
L696:
	;
	if v3347 < v3351 {
		goto L697
	} else {
		goto L698
	}
L697:
	;
	v3378 = v3370
	goto L699
L698:
	;
	v3378 = v3376
	goto L699
L699:
	;
	if base.Ui32(v3351) < base.Ui32(v3369) {
		goto L700
	} else {
		goto L701
	}
L700:
	;
	v3380 = v3378
	goto L702
L701:
	;
	v3380 = int32(0)
	goto L702
L702:
	;
	if v3380 == int32(0) {
		goto L690
	} else {
		goto L703
	}
L703:
	;
	v3384 = F_memmove(m, v3296, v3296+v3373, v3380)
	mBase = m.M
	goto L690
L704:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v3296+int32(-17)))) = base.I64_extend_i32_u(v3380)
	goto L676
L705:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3296+int32(-9)))) = v3380
	goto L675
L706:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v3296+int32(-5)))) = uint16(v3380)
	goto L675
L707:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v3296+int32(-3)))) = uint8(v3380)
	goto L675
L708:
	;
	v3389 = v3380 << (uint(int32(3)) % 32)
	*(*uint8)(unsafe.Add(mBase, uint32(v3323))) = uint8(v3389)
	goto L675
L709:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+140)) = v3439
	goto L668
L710:
	;
	v3438 = *(*int32)(unsafe.Add(mBase, uint32(v3419+int32(-17))))
	v3439 = v3438
	goto L709
L711:
	;
	v3435 = *(*int32)(unsafe.Add(mBase, uint32(v3419+int32(-9))))
	v3439 = v3435
	goto L709
L712:
	;
	v3432 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v3419+int32(-5)))))
	v3439 = v3432
	goto L709
L713:
	;
	v3429 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3419+int32(-3)))))
	v3439 = v3429
	goto L709
L714:
	;
	v3439 = int32(base.Ui32(v3422) >> (uint(int32(3)) % 32))
	goto L709
L715:
	;
	v3447 = *(*int32)(unsafe.Add(mBase, uint32(l0)+172))
	if l1 == int32(0) {
		goto L717
	} else {
		goto L718
	}
L716:
	;
	v3462 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+172)) = v3462
	v3545 = v3271
	v3546 = v3462
	goto L3
L717:
	;
	if v3447 == int32(0) {
		goto L716
	} else {
		goto L719
	}
L718:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v3447
	goto L716
L719:
	;
	v3453 = *(*int32)(unsafe.Add(mBase, uint32(l0)+176))
	if v3453 == int32(0) {
		goto L716
	} else {
		goto L720
	}
L720:
	;
	v3456 = *(*int32)(unsafe.Add(mBase, uint32(v3453)+24))
	if v3456 == int32(0) {
		goto L716
	} else {
		goto L721
	}
L721:
	;
	m.T0[v3456].(func(*base.Module, int32))(m, v3447)
	mBase = m.M
	v3460 = m.ExcPending
	if v3460 != 0 {
		goto L72
	} else {
		goto L722
	}
L722:
	;
	goto L716
L723:
	;
	F_sdsfree(m, v3518)
	mBase = m.M
	v3520 = m.ExcPending
	if v3520 != 0 {
		goto L72
	} else {
		goto L728
	}
L724:
	;
	v3506 = *(*int32)(unsafe.Add(mBase, uint32(l0)+176))
	if v3506 == int32(0) {
		v3518 = v367
		goto L723
	} else {
		goto L725
	}
L725:
	;
	v3509 = *(*int32)(unsafe.Add(mBase, uint32(v3506)+24))
	if v3509 == int32(0) {
		v3518 = v367
		goto L723
	} else {
		goto L726
	}
L726:
	;
	m.T0[v3509].(func(*base.Module, int32))(m, v3503)
	mBase = m.M
	v3513 = m.ExcPending
	if v3513 != 0 {
		goto L72
	} else {
		goto L727
	}
L727:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+172)) = int32(0)
	v3516 = *(*int32)(unsafe.Add(mBase, uint32(l0)+132))
	v3518 = v3516
	goto L723
L728:
	;
	v3521 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+140)) = v3521
	*(*int64)(unsafe.Add(mBase, uint32(l0)+132)) = int64(0)
	v3525 = int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+168)) = v3525
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(4)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+20)) = uint8(v3521)
	v3532 = m.G3
	v3535 = *(*int64)(unsafe.Add(mBase, uint32(v3532)+uint32(_c_F_valkeyReaderGetReply[42])))
	*(*int64)(unsafe.Add(mBase, uint32(l0)+4)) = v3535
	v3541 = *(*int64)(unsafe.Add(mBase, uint32(v3532)+uint32(_c_F_valkeyReaderGetReply[43])))
	*(*int64)(unsafe.Add(mBase, uint32(l0+int32(12)))) = v3541
	v3545 = v23
	v3546 = v3525
	goto L3
}
func F_valkeyReaderSetErrorOOM(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v14 int32
	_ = v14
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v31 int32
	_ = v31
	var v34 int64
	_ = v34
	var v40 int64
	_ = v40
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+172))
	if v4 == int32(0) {
		v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+132))
		F_sdsfree(m, v18)
		mBase = m.M
		v20 = m.ExcPending
		if v20 != 0 {
			return
		} else {
			v21 = int32(0)
			*(*int32)(unsafe.Add(mBase, uint32(l0)+140)) = v21
			*(*int64)(unsafe.Add(mBase, uint32(l0)+132)) = int64(0)
			*(*int32)(unsafe.Add(mBase, uint32(l0)+168)) = int32(-1)
			*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(5)
			*(*uint8)(unsafe.Add(mBase, uint32(l0)+17)) = uint8(v21)
			v31 = m.G3
			v34 = *(*int64)(unsafe.Add(mBase, uint32(v31)+uint32(_c_F_valkeyReaderSetErrorOOM[0])))
			*(*int64)(unsafe.Add(mBase, uint32(l0)+4)) = v34
			v40 = *(*int64)(unsafe.Add(mBase, uint32(v31)+uint32(_c_F_valkeyReaderSetErrorOOM[1])))
			*(*int64)(unsafe.Add(mBase, uint32(l0+int32(9)))) = v40
			return
		}
	} else {
		v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+176))
		if v7 == int32(0) {
			v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+132))
			F_sdsfree(m, v18)
			mBase = m.M
			v20 = m.ExcPending
			if v20 != 0 {
				return
			} else {
				v21 = int32(0)
				*(*int32)(unsafe.Add(mBase, uint32(l0)+140)) = v21
				*(*int64)(unsafe.Add(mBase, uint32(l0)+132)) = int64(0)
				*(*int32)(unsafe.Add(mBase, uint32(l0)+168)) = int32(-1)
				*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(5)
				*(*uint8)(unsafe.Add(mBase, uint32(l0)+17)) = uint8(v21)
				v31 = m.G3
				v34 = *(*int64)(unsafe.Add(mBase, uint32(v31)+uint32(_c_F_valkeyReaderSetErrorOOM[0])))
				*(*int64)(unsafe.Add(mBase, uint32(l0)+4)) = v34
				v40 = *(*int64)(unsafe.Add(mBase, uint32(v31)+uint32(_c_F_valkeyReaderSetErrorOOM[1])))
				*(*int64)(unsafe.Add(mBase, uint32(l0+int32(9)))) = v40
				return
			}
		} else {
			v10 = *(*int32)(unsafe.Add(mBase, uint32(v7)+24))
			if v10 == int32(0) {
				v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+132))
				F_sdsfree(m, v18)
				mBase = m.M
				v20 = m.ExcPending
				if v20 != 0 {
					return
				} else {
					v21 = int32(0)
					*(*int32)(unsafe.Add(mBase, uint32(l0)+140)) = v21
					*(*int64)(unsafe.Add(mBase, uint32(l0)+132)) = int64(0)
					*(*int32)(unsafe.Add(mBase, uint32(l0)+168)) = int32(-1)
					*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(5)
					*(*uint8)(unsafe.Add(mBase, uint32(l0)+17)) = uint8(v21)
					v31 = m.G3
					v34 = *(*int64)(unsafe.Add(mBase, uint32(v31)+uint32(_c_F_valkeyReaderSetErrorOOM[0])))
					*(*int64)(unsafe.Add(mBase, uint32(l0)+4)) = v34
					v40 = *(*int64)(unsafe.Add(mBase, uint32(v31)+uint32(_c_F_valkeyReaderSetErrorOOM[1])))
					*(*int64)(unsafe.Add(mBase, uint32(l0+int32(9)))) = v40
					return
				}
			} else {
				m.T0[v10].(func(*base.Module, int32))(m, v4)
				mBase = m.M
				v14 = m.ExcPending
				if v14 != 0 {
					return
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(l0)+172)) = int32(0)
					v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+132))
					F_sdsfree(m, v18)
					mBase = m.M
					v20 = m.ExcPending
					if v20 != 0 {
						return
					} else {
						v21 = int32(0)
						*(*int32)(unsafe.Add(mBase, uint32(l0)+140)) = v21
						*(*int64)(unsafe.Add(mBase, uint32(l0)+132)) = int64(0)
						*(*int32)(unsafe.Add(mBase, uint32(l0)+168)) = int32(-1)
						*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(5)
						*(*uint8)(unsafe.Add(mBase, uint32(l0)+17)) = uint8(v21)
						v31 = m.G3
						v34 = *(*int64)(unsafe.Add(mBase, uint32(v31)+uint32(_c_F_valkeyReaderSetErrorOOM[0])))
						*(*int64)(unsafe.Add(mBase, uint32(l0)+4)) = v34
						v40 = *(*int64)(unsafe.Add(mBase, uint32(v31)+uint32(_c_F_valkeyReaderSetErrorOOM[1])))
						*(*int64)(unsafe.Add(mBase, uint32(l0+int32(9)))) = v40
						return
					}
				}
			}
		}
	}
}
func F_valkeySetBlocking(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v18 int32
	_ = v18
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v30 int32
	_ = v30
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v59 int32
	_ = v59
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v84 int32
	_ = v84
	var v89 int32
	_ = v89
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v101 int32
	_ = v101
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v113 int32
	_ = v113
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v125 int32
	_ = v125
	var v130 int32
	_ = v130
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v144 int32
	_ = v144
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v152 int32
	_ = v152
	v6 = m.G0
	v8 = v6 - int32(176)
	m.G0 = v8
	v10 = int32(-1)
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+136))
	v14 = F_fcntl(m, v11, int32(3), int32(0))
	mBase = m.M
	if v14 != v10 {
		v72 = *(*int32)(unsafe.Add(mBase, uint32(l0)+136))
		v75 = int32(0)
		*(*int32)(unsafe.Add(mBase, uint32(v8)+32)) = v14&int32(-2049) | base.B2i32(l1 == v75)<<(uint(int32(11))%32)
		v84 = F_fcntl(m, v72, int32(4), v8+int32(32))
		mBase = m.M
		if v84 != int32(-1) {
			v152 = v75
			m.G0 = v8 + int32(176)
			return v152
		} else {
			v89 = *(*int32)(unsafe.Add(mBase, _c_F_valkeySetBlocking[0]))
			v95 = F__emscripten_memset_bulkmem(m, v8+int32(48), base.I32_extend8_s(int32(0)), int32(128))
			mBase = m.M
			v96 = m.G3
			*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = v96 + int32(_a_F_valkeySetBlocking_0)
			v101 = v8 + int32(48)
			v109 = F_snprintf(m, v101, int32(128), v96+int32(_a_F_valkeySetBlocking_1), v8+int32(16))
			mBase = m.M
			v110 = m.ExcPending
			if v110 != 0 {
				return int32(0)
			} else {
				v111 = v101 + v109
				v113 = int32(128) - v109
				v115 = F_strerror(m, v89)
				mBase = m.M
				v116 = F_strlen(m, v115)
				mBase = m.M
				if base.Ui32(v116) < base.Ui32(v113) {
					v130 = F___memcpy(m, v111, v115, v116+int32(1))
					mBase = m.M
				} else {
					if v113 == int32(0) {
					} else {
						v122 = v113 + int32(-1)
						v123 = F___memcpy(m, v111, v115, v122)
						mBase = m.M
						v125 = int32(0)
						*(*uint8)(unsafe.Add(mBase, uint32(v111+v122))) = uint8(v125)
					}
				}
				F_valkeySetError(m, l0, int32(1), v8+int32(48))
				mBase = m.M
				v140 = int32(-1)
				v141 = *(*int32)(unsafe.Add(mBase, uint32(l0)+136))
				if v141 == v140 {
					v152 = v140
				} else {
					v144 = v141
					v147 = F_close(m, v144)
					mBase = m.M
					v148 = int32(-1)
					*(*int32)(unsafe.Add(mBase, uint32(l0)+136)) = v148
					v152 = v148
				}
				m.G0 = v8 + int32(176)
				return v152
			}
		}
	} else {
		v18 = *(*int32)(unsafe.Add(mBase, _c_F_valkeySetBlocking[0]))
		v24 = F__emscripten_memset_bulkmem(m, v8+int32(48), base.I32_extend8_s(int32(0)), int32(128))
		mBase = m.M
		v25 = m.G3
		*(*int32)(unsafe.Add(mBase, uint32(v8))) = v25 + int32(_a_F_valkeySetBlocking_2)
		v30 = v8 + int32(48)
		v36 = F_snprintf(m, v30, int32(128), v25+int32(_a_F_valkeySetBlocking_1), v8)
		mBase = m.M
		v39 = m.ExcPending
		if v39 != 0 {
			return int32(0)
		} else {
			v40 = v30 + v36
			v42 = int32(128) - v36
			v44 = F_strerror(m, v18)
			mBase = m.M
			v45 = F_strlen(m, v44)
			mBase = m.M
			if base.Ui32(v45) < base.Ui32(v42) {
				v59 = F___memcpy(m, v40, v44, v45+int32(1))
				mBase = m.M
			} else {
				if v42 == int32(0) {
				} else {
					v51 = v42 + int32(-1)
					v52 = F___memcpy(m, v40, v44, v51)
					mBase = m.M
					v54 = int32(0)
					*(*uint8)(unsafe.Add(mBase, uint32(v40+v51))) = uint8(v54)
				}
			}
			F_valkeySetError(m, l0, int32(1), v8+int32(48))
			mBase = m.M
			v69 = *(*int32)(unsafe.Add(mBase, uint32(l0)+136))
			if v69 != int32(-1) {
				v144 = v69
				v147 = F_close(m, v144)
				mBase = m.M
				v148 = int32(-1)
				*(*int32)(unsafe.Add(mBase, uint32(l0)+136)) = v148
				v152 = v148
			} else {
				v152 = v10
			}
			m.G0 = v8 + int32(176)
			return v152
		}
	}
}
func F_valkeySetTcpNoDelay(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v75 int32
	_ = v75
	var v84 int32
	_ = v84
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v96 int32
	_ = v96
	var v102 int32
	_ = v102
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v120 int32
	_ = v120
	var v125 int32
	_ = v125
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v143 int32
	_ = v143
	v2 = int32(0)
	v5 = m.G0
	v7 = v5 - int32(144)
	m.G0 = v7
	v9 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v7)+12)) = v9
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+136))
	v21 = m.G0
	v23 = v21 - int32(16)
	m.G0 = v23
	v26 = F___syscall_setsockopt(m, v12, int32(6), v9, v7+int32(12), int32(4), v2)
	mBase = m.M
	v75 = F___syscall_ret(m, v26)
	mBase = m.M
	m.G0 = v23 + int32(16)
	if v75 != int32(-1) {
		v143 = v2
		m.G0 = v7 + int32(144)
		return v143
	} else {
		v84 = *(*int32)(unsafe.Add(mBase, _c_F_valkeySetTcpNoDelay[0]))
		v90 = F__emscripten_memset_bulkmem(m, v7+int32(16), base.I32_extend8_s(int32(0)), int32(128))
		mBase = m.M
		v91 = m.G3
		*(*int32)(unsafe.Add(mBase, uint32(v7))) = v91 + int32(_a_F_valkeySetTcpNoDelay_0)
		v96 = v7 + int32(16)
		v102 = F_snprintf(m, v96, int32(128), v91+int32(_a_F_valkeySetTcpNoDelay_1), v7)
		mBase = m.M
		v105 = m.ExcPending
		if v105 != 0 {
			return int32(0)
		} else {
			v106 = v96 + v102
			v108 = int32(128) - v102
			v110 = F_strerror(m, v84)
			mBase = m.M
			v111 = F_strlen(m, v110)
			mBase = m.M
			if base.Ui32(v111) < base.Ui32(v108) {
				v125 = F___memcpy(m, v106, v110, v111+int32(1))
				mBase = m.M
			} else {
				if v108 == int32(0) {
				} else {
					v117 = v108 + int32(-1)
					v118 = F___memcpy(m, v106, v110, v117)
					mBase = m.M
					v120 = int32(0)
					*(*uint8)(unsafe.Add(mBase, uint32(v106+v117))) = uint8(v120)
				}
			}
			F_valkeySetError(m, l0, int32(1), v7+int32(16))
			mBase = m.M
			v135 = int32(-1)
			v136 = *(*int32)(unsafe.Add(mBase, uint32(l0)+136))
			if v136 == v135 {
				v143 = v135
			} else {
				v139 = F_close(m, v136)
				mBase = m.M
				v140 = int32(-1)
				*(*int32)(unsafe.Add(mBase, uint32(l0)+136)) = v140
				v143 = v140
			}
			m.G0 = v7 + int32(144)
			return v143
		}
	}
}
func F_valkeySsubscribeCallback(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v15 int32
	_ = v15
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
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
	var v77 int32
	_ = v77
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
	var v116 int32
	_ = v116
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v123 int32
	_ = v123
	var v127 int32
	_ = v127
	var v131 int32
	_ = v131
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v145 int32
	_ = v145
	var v147 int32
	_ = v147
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v165 int32
	_ = v165
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v173 int32
	_ = v173
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v202 int32
	_ = v202
	var v207 int32
	_ = v207
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v214 int32
	_ = v214
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v235 int64
	_ = v235
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v250 int32
	_ = v250
	var v254 int32
	_ = v254
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v261 int32
	_ = v261
	var v262 int32
	_ = v262
	var v263 int32
	_ = v263
	var v265 int32
	_ = v265
	var v266 int32
	_ = v266
	var v268 int32
	_ = v268
	var v279 int32
	_ = v279
	var v280 int32
	_ = v280
	var v281 int32
	_ = v281
	var v282 int32
	_ = v282
	var v284 int32
	_ = v284
	var v285 int32
	_ = v285
	var v287 int32
	_ = v287
	var v299 int32
	_ = v299
	var v308 int32
	_ = v308
	var v317 int32
	_ = v317
	v9 = m.G0
	v11 = v9 - int32(48)
	m.G0 = v11
	if l2 == int32(0) {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	v317 = m.G3
	m.Env.X__assert_fail(m, v317+int32(_a_F_valkeySsubscribeCallback_0), v317+int32(_a_F_valkeySsubscribeCallback_1), int32(905), v317+int32(_a_F_valkeySsubscribeCallback_2))
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L2:
	;
	v308 = m.G3
	m.Env.X__assert_fail(m, v308+int32(_a_F_valkeySsubscribeCallback_3), v308+int32(_a_F_valkeySsubscribeCallback_1), int32(904), v308+int32(_a_F_valkeySsubscribeCallback_2))
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L3:
	;
	v299 = m.G3
	m.Env.X__assert_fail(m, v299+int32(_a_F_valkeySsubscribeCallback_4), v299+int32(_a_F_valkeySsubscribeCallback_1), int32(903), v299+int32(_a_F_valkeySsubscribeCallback_2))
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L4:
	;
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	if v15 == int32(0) {
		goto L2
	} else {
		goto L5
	}
L5:
	;
	if l1 == int32(0) {
		goto L1
	} else {
		goto L6
	}
L6:
	;
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	switch v20 + int32(-2) {
	case 0, 10:
		goto L10
	default:
		goto L9
	case 4:
		goto L11
	}
L7:
	;
	m.G0 = v11 + int32(48)
	return
L8:
	;
	F_sdsfree(m, int32(0))
	mBase = m.M
	v279 = m.ExcPending
	if v279 != 0 {
		goto L16
	} else {
		goto L66
	}
L9:
	;
	v235 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v11+int32(24)))) = v235
	*(*int64)(unsafe.Add(mBase, uint32(v11+int32(16)))) = v235
	*(*int64)(unsafe.Add(mBase, uint32(v11)+8)) = v235
	F_valkeyGetSubscribeCallback(m, l0, l1, v11+int32(8))
	mBase = m.M
	v246 = m.ExcPending
	if v246 != 0 {
		goto L16
	} else {
		goto L60
	}
L10:
	;
	v106 = m.G3
	v107 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
	v108 = *(*int32)(unsafe.Add(mBase, uint32(v107)))
	v109 = *(*int32)(unsafe.Add(mBase, uint32(v108)+28))
	v111 = v106 + int32(_a_F_valkeySsubscribeCallback_5)
	goto L33
L11:
	;
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	v28 = F_nextArgument(m, v15, v23, v11+int32(36), v11+int32(44))
	mBase = m.M
	v29 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	v31 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	v37 = F_nextArgument(m, v28, v29-v28+v31, v11+int32(32), v11+int32(40))
	mBase = m.M
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v11)+32))
	if v37|v38 == int32(0) {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	v94 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	v95 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	m.T0[v95].(func(*base.Module, int32, int32, int32))(m, l0, l1, v94)
	mBase = m.M
	v97 = m.ExcPending
	if v97 != 0 {
		goto L16
	} else {
		goto L29
	}
L13:
	;
	v46 = v37
	v47 = v38
	goto L14
L14:
	;
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v11)+40))
	v51 = F_sdsnewlen(m, v47, v50)
	mBase = m.M
	v52 = m.ExcPending
	if v52 != 0 {
		goto L16
	} else {
		goto L17
	}
L15:
	;
	goto L12
L16:
	;
	return
L17:
	;
	if v51 == int32(0) {
		goto L8
	} else {
		goto L18
	}
L18:
	;
	v55 = *(*int32)(unsafe.Add(mBase, uint32(l0)+288))
	v56 = F_dictFind(m, v55, v51)
	mBase = m.M
	v57 = m.ExcPending
	if v57 != 0 {
		goto L16
	} else {
		goto L20
	}
L19:
	;
	F_sdsfree(m, v51)
	mBase = m.M
	v74 = m.ExcPending
	if v74 != 0 {
		goto L16
	} else {
		goto L27
	}
L20:
	;
	if v56 == int32(0) {
		goto L19
	} else {
		goto L21
	}
L21:
	;
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v56)+8))
	goto L22
L22:
	;
	if v60 == int32(0) {
		goto L19
	} else {
		goto L23
	}
L23:
	;
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v60)+8))
	v65 = v63 + int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(v60)+8)) = v65
	if v65 != 0 {
		goto L19
	} else {
		goto L24
	}
L24:
	;
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v60)+20))
	if v67 != 0 {
		goto L19
	} else {
		goto L25
	}
L25:
	;
	v68 = *(*int32)(unsafe.Add(mBase, uint32(l0)+288))
	v69 = F_dictDelete(m, v68, v51)
	mBase = m.M
	v70 = m.ExcPending
	if v70 != 0 {
		goto L16
	} else {
		goto L26
	}
L26:
	;
	goto L19
L27:
	;
	v75 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	v77 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	v83 = F_nextArgument(m, v46, v75-v46+v77, v11+int32(32), v11+int32(40))
	mBase = m.M
	v84 = *(*int32)(unsafe.Add(mBase, uint32(v11)+32))
	if v83|v84 != 0 {
		v46 = v83
		v47 = v84
		goto L14
	} else {
		goto L28
	}
L28:
	;
	goto L15
L29:
	;
	v98 = m.G4
	v99 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	v100 = *(*int32)(unsafe.Add(mBase, uint32(v98)+16))
	m.T0[v100].(func(*base.Module, int32))(m, v99)
	mBase = m.M
	v102 = m.ExcPending
	if v102 != 0 {
		goto L16
	} else {
		goto L30
	}
L30:
	;
	v103 = *(*int32)(unsafe.Add(mBase, uint32(v98)+16))
	m.T0[v103].(func(*base.Module, int32))(m, l2)
	mBase = m.M
	v105 = m.ExcPending
	if v105 != 0 {
		goto L16
	} else {
		goto L31
	}
L31:
	;
	goto L7
L32:
	;
	if v155-v157 != 0 {
		goto L9
	} else {
		goto L47
	}
L33:
	;
	v116 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v109))))
	if v116 != 0 {
		goto L36
	} else {
		goto L37
	}
L35:
	;
	v155 = F_tolower(m, v150)
	mBase = m.M
	v156 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v151))))
	v157 = F_tolower(m, v156)
	mBase = m.M
	goto L32
L36:
	;
	v118 = v109
	v119 = v111
	v120 = int32(10)
	v121 = v116
	goto L39
L37:
	;
	v150 = int32(0)
	v151 = v111
	goto L35
L38:
	;
	v150 = v147 & int32(255)
	v151 = v145
	goto L35
L39:
	;
	v123 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v119))))
	if v123 == int32(0) {
		v145 = v119
		v147 = v121
		goto L38
	} else {
		goto L41
	}
L40:
	;
	v145 = v139
	v147 = int32(0)
	goto L38
L41:
	;
	v127 = v120 + int32(-1)
	if v127 == int32(0) {
		v145 = v119
		v147 = v121
		goto L38
	} else {
		goto L42
	}
L42:
	;
	v131 = v121 & int32(255)
	if v131 == v123 {
		goto L43
	} else {
		goto L44
	}
L43:
	;
	v138 = int32(1)
	v139 = v119 + v138
	v140 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v118)+1)))
	if v140 != 0 {
		v118 = v118 + v138
		v119 = v139
		v120 = v127
		v121 = v140
		goto L39
	} else {
		goto L46
	}
L44:
	;
	v133 = F_tolower(m, v131)
	mBase = m.M
	v134 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v119))))
	v135 = F_tolower(m, v134)
	mBase = m.M
	if v133 == v135 {
		goto L43
	} else {
		goto L45
	}
L45:
	;
	v137 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v118))))
	v145 = v119
	v147 = v137
	goto L38
L46:
	;
	goto L40
L47:
	;
	v165 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	v170 = F_nextArgument(m, v15, v165, v11+int32(36), v11+int32(44))
	mBase = m.M
	v171 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	v173 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	v179 = F_nextArgument(m, v170, v171-v170+v173, v11+int32(32), v11+int32(40))
	mBase = m.M
	v180 = *(*int32)(unsafe.Add(mBase, uint32(v11)+32))
	if v179|v180 == int32(0) {
		goto L9
	} else {
		goto L48
	}
L48:
	;
	v188 = v179
	v189 = v180
	goto L49
L49:
	;
	v192 = *(*int32)(unsafe.Add(mBase, uint32(v11)+40))
	v193 = F_sdsnewlen(m, v189, v192)
	mBase = m.M
	v194 = m.ExcPending
	if v194 != 0 {
		goto L16
	} else {
		goto L51
	}
L50:
	;
	goto L9
L51:
	;
	if v193 == int32(0) {
		goto L8
	} else {
		goto L52
	}
L52:
	;
	v197 = *(*int32)(unsafe.Add(mBase, uint32(l0)+288))
	v198 = F_dictFind(m, v197, v193)
	mBase = m.M
	v199 = m.ExcPending
	if v199 != 0 {
		goto L16
	} else {
		goto L54
	}
L53:
	;
	F_sdsfree(m, v193)
	mBase = m.M
	v211 = m.ExcPending
	if v211 != 0 {
		goto L16
	} else {
		goto L58
	}
L54:
	;
	if v198 == int32(0) {
		goto L53
	} else {
		goto L55
	}
L55:
	;
	v202 = *(*int32)(unsafe.Add(mBase, uint32(v198)+8))
	goto L56
L56:
	;
	if v202 == int32(0) {
		goto L53
	} else {
		goto L57
	}
L57:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v202)+20)) = int32(1)
	v207 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v202)+4)) = v207
	goto L53
L58:
	;
	v212 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	v214 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	v220 = F_nextArgument(m, v188, v212-v188+v214, v11+int32(32), v11+int32(40))
	mBase = m.M
	v221 = *(*int32)(unsafe.Add(mBase, uint32(v11)+32))
	if v220|v221 != 0 {
		v188 = v220
		v189 = v221
		goto L49
	} else {
		goto L59
	}
L59:
	;
	goto L50
L60:
	;
	v247 = *(*int32)(unsafe.Add(mBase, uint32(v11)+12))
	if v247 == int32(0) {
		goto L61
	} else {
		goto L62
	}
L61:
	;
	v261 = m.G4
	v262 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	v263 = *(*int32)(unsafe.Add(mBase, uint32(v261)+16))
	m.T0[v263].(func(*base.Module, int32))(m, v262)
	mBase = m.M
	v265 = m.ExcPending
	if v265 != 0 {
		goto L16
	} else {
		goto L64
	}
L62:
	;
	v250 = *(*int32)(unsafe.Add(mBase, uint32(l0)+140))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+140)) = v250 | int32(16)
	v254 = *(*int32)(unsafe.Add(mBase, uint32(v11)+24))
	m.T0[v247].(func(*base.Module, int32, int32, int32))(m, l0, l1, v254)
	mBase = m.M
	v256 = m.ExcPending
	if v256 != 0 {
		goto L16
	} else {
		goto L63
	}
L63:
	;
	v257 = *(*int32)(unsafe.Add(mBase, uint32(l0)+140))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+140)) = v257 & int32(-17)
	goto L61
L64:
	;
	v266 = *(*int32)(unsafe.Add(mBase, uint32(v261)+16))
	m.T0[v266].(func(*base.Module, int32))(m, l2)
	mBase = m.M
	v268 = m.ExcPending
	if v268 != 0 {
		goto L16
	} else {
		goto L65
	}
L65:
	;
	goto L7
L66:
	;
	v280 = m.G4
	v281 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	v282 = *(*int32)(unsafe.Add(mBase, uint32(v280)+16))
	m.T0[v282].(func(*base.Module, int32))(m, v281)
	mBase = m.M
	v284 = m.ExcPending
	if v284 != 0 {
		goto L16
	} else {
		goto L67
	}
L67:
	;
	v285 = *(*int32)(unsafe.Add(mBase, uint32(v280)+16))
	m.T0[v285].(func(*base.Module, int32))(m, l2)
	mBase = m.M
	v287 = m.ExcPending
	if v287 != 0 {
		goto L16
	} else {
		goto L68
	}
L68:
	;
	goto L7
}
func F_valkey_strlcat(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v60 int32
	_ = v60
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v75 int32
	_ = v75
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v89 int32
	_ = v89
	var v95 int32
	_ = v95
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v120 int32
	_ = v120
	if l2 == int32(0) {
		v28 = l0
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v32 = v28 - l0
	if l2 == v32 {
		goto L9
	} else {
		goto L10
	}
L2:
	;
	v14 = l0
	v16 = l2
	goto L3
L3:
	;
	v18 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14))))
	if v18 == int32(0) {
		v28 = v14
		goto L1
	} else {
		goto L5
	}
L4:
	;
	v28 = l0 + l2
	goto L1
L5:
	;
	v24 = v16 + int32(-1)
	if v24 != 0 {
		v14 = v14 + int32(1)
		v16 = v24
		goto L3
	} else {
		goto L6
	}
L6:
	;
	goto L4
L7:
	;
	v120 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v116))) = uint8(v120)
	return v117 - l1 + v32
L8:
	;
	v95 = v34
	v97 = l2 + (v32 ^ int32(-1))
	v98 = v28
	v100 = l1
	goto L28
L9:
	;
	if l1&int32(3) == int32(0) {
		v56 = l1
		goto L14
	} else {
		goto L15
	}
L10:
	;
	v34 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	if v34 != 0 {
		goto L8
	} else {
		goto L11
	}
L11:
	;
	v116 = v28
	v117 = l1
	goto L7
L12:
	;
	return v89 + v32
L13:
	;
	v89 = v81 - l1
	goto L12
L14:
	;
	v60 = v56
	goto L22
L15:
	;
	v42 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	if v42 != 0 {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	v45 = l1
	goto L18
L17:
	;
	v89 = l1 - l1
	goto L12
L18:
	;
	v49 = v45 + int32(1)
	if v49&int32(3) == int32(0) {
		v56 = v49
		goto L14
	} else {
		goto L20
	}
L20:
	;
	v54 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v49))))
	if v54 != 0 {
		v45 = v49
		goto L18
	} else {
		goto L21
	}
L21:
	;
	v81 = v49
	goto L13
L22:
	;
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v60)))
	v69 = int32(-2139062144)
	if (int32(16843008)-v66|v66)&v69 == v69 {
		v60 = v60 + int32(4)
		goto L22
	} else {
		goto L24
	}
L23:
	;
	v75 = v60
	goto L25
L24:
	;
	goto L23
L25:
	;
	v79 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v75))))
	if v79 != 0 {
		v75 = v75 + int32(1)
		goto L25
	} else {
		goto L27
	}
L26:
	;
	v81 = v75
	goto L13
L27:
	;
	goto L26
L28:
	;
	if v97 != 0 {
		goto L31
	} else {
		goto L32
	}
L29:
	;
	v116 = v109
	v117 = v112
	goto L7
L30:
	;
	v110 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v100)+1)))
	v112 = v100 + int32(1)
	if v110 != 0 {
		v95 = v110
		v97 = v108
		v98 = v109
		v100 = v112
		goto L28
	} else {
		goto L33
	}
L31:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v98))) = uint8(v95)
	v108 = v97 + int32(-1)
	v109 = v98 + int32(1)
	goto L30
L32:
	;
	v108 = int32(0)
	v109 = v98
	goto L30
L33:
	;
	goto L29
}
func F_valkey_strlcpy(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v28 int32
	_ = v28
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v42 int32
	_ = v42
	if l2 == int32(0) {
		v28 = l1
		goto L2
	} else {
		goto L3
	}
L1:
	;
	return v42 + (l1 ^ int32(-1))
L2:
	;
	v33 = v28
	goto L9
L3:
	;
	v8 = l0
	v10 = l2
	v12 = l1
	goto L5
L4:
	;
	v23 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v8))) = uint8(v23)
	v28 = v12
	goto L2
L5:
	;
	v14 = v10 + int32(-1)
	if v14 == int32(0) {
		goto L4
	} else {
		goto L7
	}
L7:
	;
	v17 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12))))
	*(*uint8)(unsafe.Add(mBase, uint32(v8))) = uint8(v17)
	v19 = int32(1)
	v22 = v12 + v19
	if v17 != 0 {
		v8 = v8 + v19
		v10 = v14
		v12 = v22
		goto L5
	} else {
		goto L8
	}
L8:
	;
	v42 = v22
	goto L1
L9:
	;
	v35 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v33))))
	v37 = v33 + int32(1)
	if v35 != 0 {
		v33 = v37
		goto L9
	} else {
		goto L11
	}
L10:
	;
	v42 = v37
	goto L1
L11:
	;
	goto L10
}
