package p0

import (
	base "github.com/shibukawa/valkeymem/internal/aot/vkaot/base"
	"unsafe"
)

func F_clusterAcceptHandler(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v37 int32
	_ = v37
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
	var v57 int32
	_ = v57
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
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
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v98 int64
	_ = v98
	var v100 int64
	_ = v100
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	var v110 int32
	_ = v110
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
	var v128 int32
	_ = v128
	var v132 int32
	_ = v132
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v149 int32
	_ = v149
	var v152 int32
	_ = v152
	v9 = m.G0
	v11 = v9 - int32(128)
	m.G0 = v11
	v17 = *(*int32)(unsafe.Add(mBase, _consts[89]))
	if v17 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v18 = int32(3688)
	goto L3
L2:
	;
	v18 = int32(3692)
	goto L3
L3:
	;
	v20 = *(*int32)(unsafe.Add(mBase, uint32(v18)+uint32(_consts[115])))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+60)) = int32(1)
	v24 = *(*int32)(unsafe.Add(mBase, _consts[64]))
	if v24 != 0 {
		goto L5
	} else {
		goto L6
	}
L4:
	;
	m.G0 = v11 + int32(128)
	return
L5:
	;
	if v20 == int32(0) {
		goto L4
	} else {
		goto L8
	}
L6:
	;
	v26 = *(*int32)(unsafe.Add(mBase, _consts[116]))
	if v26 != 0 {
		goto L4
	} else {
		goto L7
	}
L7:
	;
	goto L5
L8:
	;
	v37 = v20
	goto L9
L9:
	;
	v45 = F_anetTcpAccept(m, int32(_a227), l1, v11+int32(64), int32(46), v11+int32(124))
	mBase = m.M
	v46 = m.ExcPending
	if v46 != 0 {
		goto L13
	} else {
		goto L14
	}
L10:
	;
	goto L4
L11:
	;
	v152 = v37 + int32(-1)
	if v152 != 0 {
		v37 = v152
		goto L9
	} else {
		goto L43
	}
L12:
	;
	v70 = F_connTypeOfCluster(m)
	mBase = m.M
	v71 = m.ExcPending
	if v71 != 0 {
		goto L13
	} else {
		goto L23
	}
L13:
	;
	return
L14:
	;
	if v45 != int32(-1) {
		goto L12
	} else {
		goto L15
	}
L15:
	;
	v49 = int32(9116376)
	goto L16
L16:
	;
	v50 = *(*int32)(unsafe.Add(mBase, _consts[9]))
	goto L17
L17:
	;
	if v50 == int32(13) {
		goto L11
	} else {
		goto L18
	}
L18:
	;
	v53 = *(*int32)(unsafe.Add(mBase, _consts[9]))
	if v53 == int32(6) {
		goto L4
	} else {
		goto L19
	}
L19:
	;
	v57 = *(*int32)(unsafe.Add(mBase, _consts[15]))
	if int32(1) < v57 {
		goto L4
	} else {
		goto L20
	}
L20:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11))) = int32(_a227)
	F__serverLog(m, int32(1), int32(_a228), v11)
	mBase = m.M
	v67 = m.ExcPending
	if v67 != 0 {
		goto L13
	} else {
		goto L21
	}
L21:
	;
	goto L4
L22:
	;
	v98 = *(*int64)(unsafe.Add(mBase, _consts[117]))
	v100 = base.I64_div_s(v98, int64(1000))
	v104 = F_connKeepAlive(m, v73, base.I32_wrap_i64(v100)<<(uint(int32(1))%32))
	mBase = m.M
	v105 = m.ExcPending
	if v105 != 0 {
		goto L13
	} else {
		goto L31
	}
L23:
	;
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v70)+44))
	v73 = m.T0[v72].(func(*base.Module, int32, int32) int32)(m, v45, v11+int32(60))
	mBase = m.M
	v74 = m.ExcPending
	if v74 != 0 {
		goto L13
	} else {
		goto L24
	}
L24:
	;
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v73)+4))
	if v75 == int32(2) {
		goto L22
	} else {
		goto L25
	}
L25:
	;
	v79 = *(*int32)(unsafe.Add(mBase, _consts[15]))
	if int32(1) < v79 {
		goto L26
	} else {
		goto L27
	}
L26:
	;
	v93 = *(*int32)(unsafe.Add(mBase, uint32(v73)))
	v94 = *(*int32)(unsafe.Add(mBase, uint32(v93)+52))
	m.T0[v94].(func(*base.Module, int32))(m, v73)
	mBase = m.M
	v96 = m.ExcPending
	if v96 != 0 {
		goto L13
	} else {
		goto L30
	}
L27:
	;
	v82 = *(*int32)(unsafe.Add(mBase, uint32(v73)))
	v83 = *(*int32)(unsafe.Add(mBase, uint32(v82)+88))
	v84 = m.T0[v83].(func(*base.Module, int32) int32)(m, v73)
	mBase = m.M
	v85 = m.ExcPending
	if v85 != 0 {
		goto L13
	} else {
		goto L28
	}
L28:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+48)) = v84
	F__serverLog(m, int32(1), int32(_a229), v11+int32(48))
	mBase = m.M
	v92 = m.ExcPending
	if v92 != 0 {
		goto L13
	} else {
		goto L29
	}
L29:
	;
	goto L26
L30:
	;
	goto L4
L31:
	;
	v107 = *(*int32)(unsafe.Add(mBase, _consts[15]))
	if int32(1) < v107 {
		goto L32
	} else {
		goto L33
	}
L32:
	;
	v122 = *(*int32)(unsafe.Add(mBase, uint32(v73)))
	v123 = *(*int32)(unsafe.Add(mBase, uint32(v122)+64))
	v124 = m.T0[v123].(func(*base.Module, int32, int32) int32)(m, v73, int32(62))
	mBase = m.M
	v125 = m.ExcPending
	if v125 != 0 {
		goto L13
	} else {
		goto L35
	}
L33:
	;
	v110 = *(*int32)(unsafe.Add(mBase, uint32(v11)+124))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+36)) = v110
	*(*int32)(unsafe.Add(mBase, uint32(v11)+32)) = v11 + int32(64)
	F__serverLog(m, int32(1), int32(_a230), v11+int32(32))
	mBase = m.M
	v120 = m.ExcPending
	if v120 != 0 {
		goto L13
	} else {
		goto L34
	}
L34:
	;
	goto L32
L35:
	;
	if v124 != int32(-1) {
		goto L11
	} else {
		goto L36
	}
L36:
	;
	v128 = *(*int32)(unsafe.Add(mBase, uint32(v73)+4))
	if v128 != int32(5) {
		goto L37
	} else {
		goto L38
	}
L37:
	;
	v146 = *(*int32)(unsafe.Add(mBase, uint32(v73)))
	v147 = *(*int32)(unsafe.Add(mBase, uint32(v146)+52))
	m.T0[v147].(func(*base.Module, int32))(m, v73)
	mBase = m.M
	v149 = m.ExcPending
	if v149 != 0 {
		goto L13
	} else {
		goto L42
	}
L38:
	;
	v132 = *(*int32)(unsafe.Add(mBase, _consts[15]))
	if int32(1) < v132 {
		goto L37
	} else {
		goto L39
	}
L39:
	;
	v135 = *(*int32)(unsafe.Add(mBase, uint32(v73)))
	v136 = *(*int32)(unsafe.Add(mBase, uint32(v135)+88))
	v137 = m.T0[v136].(func(*base.Module, int32) int32)(m, v73)
	mBase = m.M
	v138 = m.ExcPending
	if v138 != 0 {
		goto L13
	} else {
		goto L40
	}
L40:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = v137
	F__serverLog(m, int32(1), int32(_a231), v11+int32(16))
	mBase = m.M
	v145 = m.ExcPending
	if v145 != 0 {
		goto L13
	} else {
		goto L41
	}
L41:
	;
	goto L37
L42:
	;
	goto L4
L43:
	;
	goto L10
}
func F_clusterAddNode(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
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
	var v18 int32
	_ = v18
	v3 = *(*int32)(unsafe.Add(mBase, _consts[90]))
	v4 = *(*int32)(unsafe.Add(mBase, uint32(v3)+32))
	v8 = F_sdsnewlen(m, l0+int32(8), int32(40))
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return
	} else {
		v10 = F_dictAdd(m, v4, v8, l0)
		mBase = m.M
		v11 = m.ExcPending
		if v11 != 0 {
			return
		} else {
			if v10 == int32(0) {
				return
			} else {
				F__serverAssert(m, int32(_a182), int32(_a179), int32(2220))
				mBase = m.M
				v18 = m.ExcPending
				if v18 != 0 {
					return
				} else {
					F_abort(m)
					mBase = m.M
					base.Wasm_trap_unreachable()
					for {
					}
				}
			}
		}
	}
}
func F_clusterAddSlot(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v13 int32
	_ = v13
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v84 int32
	_ = v84
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v93 int64
	_ = v93
	var v104 int32
	_ = v104
	v7 = *(*int32)(unsafe.Add(mBase, _consts[90]))
	v9 = l1 << (uint(int32(2)) % 32)
	v13 = *(*int32)(unsafe.Add(mBase, uint32(v7+v9+int32(52))))
	if v13 != 0 {
		v104 = int32(-1)
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return v104
L2:
	;
	v17 = m.G0
	v19 = v17 - int32(32)
	m.G0 = v19
	v24 = int32(1) << (uint(l1&int32(7)) % 32)
	v26 = base.I32_div_s(l1, int32(8))
	v29 = l0 + v26 + int32(104)
	v30 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v29))))
	if v24&v30 != 0 {
		goto L4
	} else {
		goto L5
	}
L3:
	;
	v68 = int32(_a69)
	v69 = *(*int32)(unsafe.Add(mBase, _consts[90]))
	*(*int32)(unsafe.Add(mBase, uint32(v69+v9+int32(52)))) = l0
	v75 = base.I32_div_s(l1, int32(8))
	v76 = v69 + v75
	v79 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v76)+uint32(_consts[96]))))
	v84 = v79 & base.I32_rotl(int32(-2), l1&int32(7))
	*(*uint8)(unsafe.Add(mBase, uint32(v76)+uint32(_consts[96]))) = uint8(v84)
	v87 = *(*int32)(unsafe.Add(mBase, _consts[90]))
	v90 = v87 + l1*int32(24)
	v93 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v90)+uint32(_consts[97]))) = v93
	*(*int64)(unsafe.Add(mBase, uint32(v90)+uint32(_consts[98]))) = v93
	*(*int64)(unsafe.Add(mBase, uint32(v90)+uint32(_consts[99]))) = v93
	goto L15
L4:
	;
	m.G0 = v19 + int32(32)
	goto L3
L5:
	;
	v32 = v30 | v24
	*(*uint8)(unsafe.Add(mBase, uint32(v29))) = uint8(v32)
	v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)+2160))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+2160)) = v34 + int32(1)
	if v34 != 0 {
		goto L4
	} else {
		goto L6
	}
L6:
	;
	v39 = *(*int32)(unsafe.Add(mBase, _consts[90]))
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v39)+32))
	F_dictInitIterator(m, v19, v40)
	mBase = m.M
	v42 = F_dictNext(m, v19)
	mBase = m.M
	if v42 == int32(0) {
		goto L4
	} else {
		goto L7
	}
L7:
	;
	v46 = v42
	goto L9
L8:
	;
	v56 = *(*int32)(unsafe.Add(mBase, uint32(l0)+88))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+88)) = v56 | int32(256)
	goto L4
L9:
	;
	v50 = F_dictGetVal(m, v46)
	mBase = m.M
	v51 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v50)+88)))
	if v51&int32(2) != 0 {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v55 = F_dictNext(m, v19)
	mBase = m.M
	if v55 != 0 {
		v46 = v55
		goto L9
	} else {
		goto L14
	}
L12:
	;
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v50)+2164))
	if v54 != 0 {
		goto L8
	} else {
		goto L13
	}
L13:
	;
	goto L11
L14:
	;
	goto L4
L15:
	;
	v104 = int32(0)
	goto L1
}
func F_clusterBroadcastPong(m *base.Module, l0 int32) {
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
	var v125 int32
	_ = v125
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v134 int32
	_ = v134
	var v136 int32
	_ = v136
	var v143 int32
	_ = v143
	var v147 int32
	_ = v147
	var v152 int32
	_ = v152
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v167 int32
	_ = v167
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v175 int32
	_ = v175
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v184 int32
	_ = v184
	var v186 int32
	_ = v186
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
	var v193 int64
	_ = v193
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
	var v204 int64
	_ = v204
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v211 int32
	_ = v211
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v218 int32
	_ = v218
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v226 int32
	_ = v226
	var v230 int32
	_ = v230
	var v232 int32
	_ = v232
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v246 int32
	_ = v246
	var v251 int32
	_ = v251
	var v255 int32
	_ = v255
	var v259 int32
	_ = v259
	var v270 int32
	_ = v270
	v8 = *(*int32)(unsafe.Add(mBase, _consts[90]))
	v9 = *(*int32)(unsafe.Add(mBase, uint32(v8)+32))
	v10 = F_dictGetSafeIterator(m, v9)
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	F_dictReleaseIterator(m, v10)
	mBase = m.M
	v270 = m.ExcPending
	if v270 != 0 {
		goto L2
	} else {
		goto L72
	}
L2:
	;
	return
L3:
	;
	v19 = v10 + int32(20)
	v20 = *(*int32)(unsafe.Add(mBase, uint32(v10)+16))
	if v20 != 0 {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	if v115 == int32(0) {
		goto L1
	} else {
		goto L30
	}
L5:
	;
	v26 = v19
	v27 = v23
	goto L8
L6:
	;
	v23 = int32(1)
	goto L5
L7:
	;
	v23 = int32(0)
	goto L5
L8:
	;
	switch v27 {
	case 0:
		goto L13
	default:
		goto L12
	}
L10:
	;
	v27 = int32(0)
	goto L8
L11:
	;
	goto L4
L12:
	;
	v107 = *(*int32)(unsafe.Add(mBase, uint32(v26)))
	*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = v107
	if v107 == int32(0) {
		goto L10
	} else {
		goto L29
	}
L13:
	;
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v10)+4))
	if v31 != int32(-1) {
		v70 = v31
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v71 = int32(1)
	v72 = v70 + v71
	*(*int32)(unsafe.Add(mBase, uint32(v10)+4)) = v72
	v74 = int32(0)
	v77 = *(*int32)(unsafe.Add(mBase, uint32(v10)))
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v10)+8))
	v82 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v77+v78+int32(26)))))
	if v82 == int32(255) {
		goto L23
	} else {
		goto L24
	}
L15:
	;
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v10)+8))
	if v35 != 0 {
		v70 = int32(-1)
		goto L14
	} else {
		goto L16
	}
L16:
	;
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v10)))
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v10)+12))
	if v37 == int32(0) {
		goto L18
	} else {
		goto L19
	}
L17:
	;
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v63)+20))
	if v64 != int32(-1) {
		goto L20
	} else {
		goto L21
	}
L18:
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
	*(*int64)(unsafe.Add(mBase, uint32(v10)+24)) = v60
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v10)))
	v63 = v62
	goto L17
L19:
	;
	v40 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v36)+24)))
	v42 = v40 + int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v36)+24)) = uint16(v42)
	v63 = v36
	goto L17
L20:
	;
	v70 = v64 + int32(-1)
	goto L14
L21:
	;
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v10)+4))
	v70 = v67
	goto L14
L22:
	;
	v97 = int32(2)
	v102 = *(*int32)(unsafe.Add(mBase, uint32(v77+v95<<(uint(v97)%32)+int32(4))))
	v26 = v102 + v96<<(uint(v97)%32)
	v27 = int32(1)
	goto L8
L23:
	;
	v86 = v74
	goto L25
L24:
	;
	v86 = v71 << (uint(v82) % 32)
	goto L25
L25:
	;
	if v72 < v86 {
		v95 = v78
		v96 = v72
		goto L22
	} else {
		goto L26
	}
L26:
	;
	if v78 != 0 {
		v115 = v74
		goto L11
	} else {
		goto L27
	}
L27:
	;
	v88 = *(*int32)(unsafe.Add(mBase, uint32(v77)+20))
	if v88 == int32(-1) {
		v115 = v74
		goto L11
	} else {
		goto L28
	}
L28:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v10)+4)) = int64(4294967296)
	v95 = int32(1)
	v96 = int32(0)
	goto L22
L29:
	;
	v111 = *(*int32)(unsafe.Add(mBase, uint32(v107)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v19))) = v111
	v115 = v107
	goto L11
L30:
	;
	v125 = v115
	goto L31
L31:
	;
	v129 = *(*int32)(unsafe.Add(mBase, uint32(v125)+8))
	goto L34
L32:
	;
	goto L1
L33:
	;
	v163 = v10 + int32(20)
	v164 = *(*int32)(unsafe.Add(mBase, uint32(v10)+16))
	if v164 != 0 {
		goto L47
	} else {
		goto L48
	}
L34:
	;
	v130 = *(*int32)(unsafe.Add(mBase, uint32(v129)+2344))
	if v130 == int32(0) {
		goto L33
	} else {
		goto L35
	}
L35:
	;
	v134 = *(*int32)(unsafe.Add(mBase, _consts[94]))
	if v129 == v134 {
		goto L33
	} else {
		goto L36
	}
L36:
	;
	v136 = *(*int32)(unsafe.Add(mBase, uint32(v129)+88))
	if v136&int32(32) != 0 {
		goto L33
	} else {
		goto L37
	}
L37:
	;
	if l0 != int32(1) {
		goto L38
	} else {
		goto L39
	}
L38:
	;
	F_clusterSendPing(m, v130, int32(1))
	mBase = m.M
	v152 = m.ExcPending
	if v152 != 0 {
		goto L2
	} else {
		goto L44
	}
L39:
	;
	if v136&int32(2) == int32(0) {
		goto L33
	} else {
		goto L40
	}
L40:
	;
	v143 = *(*int32)(unsafe.Add(mBase, uint32(v129)+2172))
	if v143 == int32(0) {
		goto L33
	} else {
		goto L41
	}
L41:
	;
	if v143 == v134 {
		goto L38
	} else {
		goto L42
	}
L42:
	;
	v147 = *(*int32)(unsafe.Add(mBase, uint32(v134)+2172))
	if v143 != v147 {
		goto L33
	} else {
		goto L43
	}
L43:
	;
	goto L38
L44:
	;
	goto L33
L45:
	;
	if v259 != 0 {
		v125 = v259
		goto L31
	} else {
		goto L71
	}
L46:
	;
	v170 = v163
	v171 = v167
	goto L49
L47:
	;
	v167 = int32(1)
	goto L46
L48:
	;
	v167 = int32(0)
	goto L46
L49:
	;
	switch v171 {
	case 0:
		goto L54
	default:
		goto L53
	}
L51:
	;
	v171 = int32(0)
	goto L49
L52:
	;
	goto L45
L53:
	;
	v251 = *(*int32)(unsafe.Add(mBase, uint32(v170)))
	*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = v251
	if v251 == int32(0) {
		goto L51
	} else {
		goto L70
	}
L54:
	;
	v175 = *(*int32)(unsafe.Add(mBase, uint32(v10)+4))
	if v175 != int32(-1) {
		v214 = v175
		goto L55
	} else {
		goto L56
	}
L55:
	;
	v215 = int32(1)
	v216 = v214 + v215
	*(*int32)(unsafe.Add(mBase, uint32(v10)+4)) = v216
	v218 = int32(0)
	v221 = *(*int32)(unsafe.Add(mBase, uint32(v10)))
	v222 = *(*int32)(unsafe.Add(mBase, uint32(v10)+8))
	v226 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v221+v222+int32(26)))))
	if v226 == int32(255) {
		goto L64
	} else {
		goto L65
	}
L56:
	;
	v179 = *(*int32)(unsafe.Add(mBase, uint32(v10)+8))
	if v179 != 0 {
		v214 = int32(-1)
		goto L55
	} else {
		goto L57
	}
L57:
	;
	v180 = *(*int32)(unsafe.Add(mBase, uint32(v10)))
	v181 = *(*int32)(unsafe.Add(mBase, uint32(v10)+12))
	if v181 == int32(0) {
		goto L59
	} else {
		goto L60
	}
L58:
	;
	v208 = *(*int32)(unsafe.Add(mBase, uint32(v207)+20))
	if v208 != int32(-1) {
		goto L61
	} else {
		goto L62
	}
L59:
	;
	v188 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v180)+16)))
	v189 = int64(*(*int8)(unsafe.Add(mBase, uint32(v180)+27)))
	v190 = int64(*(*int32)(unsafe.Add(mBase, uint32(v180)+8)))
	v191 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v180)+12)))
	v192 = int64(*(*int8)(unsafe.Add(mBase, uint32(v180)+26)))
	v193 = int64(*(*int32)(unsafe.Add(mBase, uint32(v180)+4)))
	v194 = F_wangHash64(m, v193)
	mBase = m.M
	v196 = F_wangHash64(m, v192+v194)
	mBase = m.M
	v198 = F_wangHash64(m, v191+v196)
	mBase = m.M
	v200 = F_wangHash64(m, v190+v198)
	mBase = m.M
	v202 = F_wangHash64(m, v189+v200)
	mBase = m.M
	v204 = F_wangHash64(m, v188+v202)
	mBase = m.M
	*(*int64)(unsafe.Add(mBase, uint32(v10)+24)) = v204
	v206 = *(*int32)(unsafe.Add(mBase, uint32(v10)))
	v207 = v206
	goto L58
L60:
	;
	v184 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v180)+24)))
	v186 = v184 + int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v180)+24)) = uint16(v186)
	v207 = v180
	goto L58
L61:
	;
	v214 = v208 + int32(-1)
	goto L55
L62:
	;
	v211 = *(*int32)(unsafe.Add(mBase, uint32(v10)+4))
	v214 = v211
	goto L55
L63:
	;
	v241 = int32(2)
	v246 = *(*int32)(unsafe.Add(mBase, uint32(v221+v239<<(uint(v241)%32)+int32(4))))
	v170 = v246 + v240<<(uint(v241)%32)
	v171 = int32(1)
	goto L49
L64:
	;
	v230 = v218
	goto L66
L65:
	;
	v230 = v215 << (uint(v226) % 32)
	goto L66
L66:
	;
	if v216 < v230 {
		v239 = v222
		v240 = v216
		goto L63
	} else {
		goto L67
	}
L67:
	;
	if v222 != 0 {
		v259 = v218
		goto L52
	} else {
		goto L68
	}
L68:
	;
	v232 = *(*int32)(unsafe.Add(mBase, uint32(v221)+20))
	if v232 == int32(-1) {
		v259 = v218
		goto L52
	} else {
		goto L69
	}
L69:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v10)+4)) = int64(4294967296)
	v239 = int32(1)
	v240 = int32(0)
	goto L63
L70:
	;
	v255 = *(*int32)(unsafe.Add(mBase, uint32(v251)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v163))) = v255
	v259 = v251
	goto L52
L71:
	;
	goto L32
L72:
	;
	return
}
func F_clusterCleanSlotImportsAfterLoad(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v17 int32
	_ = v17
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v43 int32
	_ = v43
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v78 int32
	_ = v78
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	v9 = *(*int32)(unsafe.Add(mBase, _consts[90]))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(v9)))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(v10)+88))
	goto L2
L1:
	;
	m.G0 = v6 + int32(16)
	return
L2:
	;
	if v11&int32(1) == int32(0) {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v17 = *(*int32)(unsafe.Add(mBase, _consts[63]))
	if v17 == int32(0) {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v21 = *(*int32)(unsafe.Add(mBase, _consts[90]))
	if v21 == int32(0) {
		goto L1
	} else {
		goto L5
	}
L5:
	;
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v21)+uint32(_consts[143])))
	v26 = v6 + int32(8)
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v24)))
	*(*int32)(unsafe.Add(mBase, uint32(v26)+4)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v26))) = v27
	goto L6
L6:
	;
	v32 = v6 + int32(8)
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v32)))
	if v34 == int32(0) {
		goto L8
	} else {
		goto L9
	}
L7:
	;
	if v34 == int32(0) {
		goto L1
	} else {
		goto L10
	}
L8:
	;
	goto L7
L9:
	;
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v32)+4))
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v34+base.B2i32(v37 == int32(0))<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v32))) = v43
	goto L8
L10:
	;
	v48 = v34
	goto L11
L11:
	;
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v48)+8))
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v50)))
	if v51 != int32(1) {
		goto L13
	} else {
		goto L14
	}
L12:
	;
	goto L1
L13:
	;
	v67 = v6 + int32(8)
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v67)))
	if v69 == int32(0) {
		goto L21
	} else {
		goto L22
	}
L14:
	;
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v50)+156))
	if base.Ui32(int32(20)) < base.Ui32(v54) {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	F_finishSlotMigrationJob(m, v50, int32(18), int32(_a346))
	mBase = m.M
	v64 = m.ExcPending
	if v64 != 0 {
		goto L18
	} else {
		goto L19
	}
L16:
	;
	if int32(1)<<(uint(v54)%32)&int32(1835040) != 0 {
		goto L13
	} else {
		goto L17
	}
L17:
	;
	goto L15
L18:
	;
	return
L19:
	;
	goto L13
L20:
	;
	if v69 != 0 {
		v48 = v69
		goto L11
	} else {
		goto L23
	}
L21:
	;
	goto L20
L22:
	;
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v67)+4))
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v69+base.B2i32(v72 == int32(0))<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v67))) = v78
	goto L21
L23:
	;
	goto L12
}
func F_clusterCommandCancelSlotMigrations(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v36 int32
	_ = v36
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v63 int32
	_ = v63
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v87 int32
	_ = v87
	var v93 int32
	_ = v93
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v109 int32
	_ = v109
	var v115 int32
	_ = v115
	var v117 int32
	_ = v117
	var v119 int32
	_ = v119
	var v122 int32
	_ = v122
	var v128 int32
	_ = v128
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
	var v148 int32
	_ = v148
	var v150 int32
	_ = v150
	var v152 int32
	_ = v152
	var v154 int32
	_ = v154
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v10 = *(*int32)(unsafe.Add(mBase, _consts[63]))
	if v10 == int32(0) {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	m.G0 = v7 + int32(16)
	return
L2:
	;
	v73 = *(*int32)(unsafe.Add(mBase, _consts[90]))
	v74 = *(*int32)(unsafe.Add(mBase, uint32(v73)+uint32(_consts[143])))
	v76 = v7 + int32(8)
	v77 = *(*int32)(unsafe.Add(mBase, uint32(v74)))
	*(*int32)(unsafe.Add(mBase, uint32(v76)+4)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v76))) = v77
	goto L22
L3:
	;
	F_addReplyError(m, l0, int32(_a371))
	mBase = m.M
	v71 = m.ExcPending
	if v71 != 0 {
		goto L20
	} else {
		goto L21
	}
L4:
	;
	v14 = *(*int32)(unsafe.Add(mBase, _consts[90]))
	if v14 == int32(0) {
		goto L3
	} else {
		goto L5
	}
L5:
	;
	v17 = *(*int32)(unsafe.Add(mBase, uint32(v14)+uint32(_consts[143])))
	v19 = v7 + int32(8)
	v20 = *(*int32)(unsafe.Add(mBase, uint32(v17)))
	*(*int32)(unsafe.Add(mBase, uint32(v19)+4)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v19))) = v20
	goto L6
L6:
	;
	v25 = v7 + int32(8)
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v25)))
	if v27 == int32(0) {
		goto L8
	} else {
		goto L9
	}
L7:
	;
	if v27 == int32(0) {
		goto L3
	} else {
		goto L10
	}
L8:
	;
	goto L7
L9:
	;
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v25)+4))
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v27+base.B2i32(v30 == int32(0))<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v25))) = v36
	goto L8
L10:
	;
	v42 = v27
	goto L11
L11:
	;
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v42)+8))
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v44)))
	if v45 != 0 {
		goto L13
	} else {
		goto L14
	}
L12:
	;
	goto L3
L13:
	;
	v52 = v7 + int32(8)
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v52)))
	if v54 == int32(0) {
		goto L17
	} else {
		goto L18
	}
L14:
	;
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v44)+156))
	if base.Ui32(int32(2)) < base.Ui32(v46+int32(-18)) {
		goto L2
	} else {
		goto L15
	}
L15:
	;
	goto L13
L16:
	;
	if v54 != 0 {
		v42 = v54
		goto L11
	} else {
		goto L19
	}
L17:
	;
	goto L16
L18:
	;
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v52)+4))
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v54+base.B2i32(v57 == int32(0))<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v52))) = v63
	goto L17
L19:
	;
	goto L12
L20:
	;
	return
L21:
	;
	goto L1
L22:
	;
	v82 = v7 + int32(8)
	v84 = *(*int32)(unsafe.Add(mBase, uint32(v82)))
	if v84 == int32(0) {
		goto L25
	} else {
		goto L26
	}
L23:
	;
	v134 = F_sdsempty(m)
	mBase = m.M
	v135 = m.ExcPending
	if v135 != 0 {
		goto L20
	} else {
		goto L40
	}
L24:
	;
	if v84 == int32(0) {
		goto L23
	} else {
		goto L27
	}
L25:
	;
	goto L24
L26:
	;
	v87 = *(*int32)(unsafe.Add(mBase, uint32(v82)+4))
	v93 = *(*int32)(unsafe.Add(mBase, uint32(v84+base.B2i32(v87 == int32(0))<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v82))) = v93
	goto L25
L27:
	;
	v99 = v84
	goto L28
L28:
	;
	v101 = *(*int32)(unsafe.Add(mBase, uint32(v99)+8))
	v102 = *(*int32)(unsafe.Add(mBase, uint32(v101)+156))
	if base.Ui32(int32(20)) < base.Ui32(v102) {
		goto L31
	} else {
		goto L32
	}
L29:
	;
	goto L23
L30:
	;
	v117 = v7 + int32(8)
	v119 = *(*int32)(unsafe.Add(mBase, uint32(v117)))
	if v119 == int32(0) {
		goto L37
	} else {
		goto L38
	}
L31:
	;
	v109 = *(*int32)(unsafe.Add(mBase, uint32(v101)))
	if v109 == int32(1) {
		goto L30
	} else {
		goto L34
	}
L32:
	;
	if int32(1)<<(uint(v102)%32)&int32(1835040) != 0 {
		goto L30
	} else {
		goto L33
	}
L33:
	;
	goto L31
L34:
	;
	F_finishSlotMigrationJob(m, v101, int32(19), int32(0))
	mBase = m.M
	v115 = m.ExcPending
	if v115 != 0 {
		goto L20
	} else {
		goto L35
	}
L35:
	;
	goto L30
L36:
	;
	if v119 != 0 {
		v99 = v119
		goto L28
	} else {
		goto L39
	}
L37:
	;
	goto L36
L38:
	;
	v122 = *(*int32)(unsafe.Add(mBase, uint32(v117)+4))
	v128 = *(*int32)(unsafe.Add(mBase, uint32(v119+base.B2i32(v122 == int32(0))<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v117))) = v128
	goto L37
L39:
	;
	goto L29
L40:
	;
	v137 = *(*int32)(unsafe.Add(mBase, _consts[151]))
	v138 = F_catClientInfoShortString(m, v134, l0, v137)
	mBase = m.M
	v139 = m.ExcPending
	if v139 != 0 {
		goto L20
	} else {
		goto L41
	}
L41:
	;
	v141 = *(*int32)(unsafe.Add(mBase, _consts[15]))
	if int32(2) < v141 {
		goto L42
	} else {
		goto L43
	}
L42:
	;
	F_sdsfree(m, v138)
	mBase = m.M
	v150 = m.ExcPending
	if v150 != 0 {
		goto L20
	} else {
		goto L45
	}
L43:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v7))) = v138
	F__serverLog(m, int32(2), int32(_a372), v7)
	mBase = m.M
	v148 = m.ExcPending
	if v148 != 0 {
		goto L20
	} else {
		goto L44
	}
L44:
	;
	goto L42
L45:
	;
	v152 = *(*int32)(unsafe.Add(mBase, _consts[77]))
	F_addReply(m, l0, v152)
	mBase = m.M
	v154 = m.ExcPending
	if v154 != 0 {
		goto L20
	} else {
		goto L46
	}
L46:
	;
	goto L1
}
func F_clusterCommandShards(m *base.Module, l0 int32) {
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
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v60 int64
	_ = v60
	var v61 int64
	_ = v61
	var v62 int64
	_ = v62
	var v63 int64
	_ = v63
	var v64 int64
	_ = v64
	var v65 int64
	_ = v65
	var v66 int64
	_ = v66
	var v68 int64
	_ = v68
	var v70 int64
	_ = v70
	var v72 int64
	_ = v72
	var v74 int64
	_ = v74
	var v76 int64
	_ = v76
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v98 int32
	_ = v98
	var v102 int32
	_ = v102
	var v104 int32
	_ = v104
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v118 int32
	_ = v118
	var v123 int32
	_ = v123
	var v127 int32
	_ = v127
	var v131 int32
	_ = v131
	var v141 int32
	_ = v141
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v150 int32
	_ = v150
	var v153 int32
	_ = v153
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v161 int32
	_ = v161
	var v163 int32
	_ = v163
	var v166 int32
	_ = v166
	var v172 int32
	_ = v172
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v186 int32
	_ = v186
	var v188 int32
	_ = v188
	var v191 int32
	_ = v191
	var v197 int32
	_ = v197
	var v199 int32
	_ = v199
	var v203 int32
	_ = v203
	var v207 int32
	_ = v207
	var v209 int32
	_ = v209
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v223 int64
	_ = v223
	var v225 int32
	_ = v225
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v239 int32
	_ = v239
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v252 int32
	_ = v252
	var v254 int32
	_ = v254
	var v255 int32
	_ = v255
	var v260 int32
	_ = v260
	var v262 int32
	_ = v262
	var v265 int32
	_ = v265
	var v271 int32
	_ = v271
	var v277 int32
	_ = v277
	var v282 int32
	_ = v282
	var v284 int32
	_ = v284
	var v285 int32
	_ = v285
	var v287 int32
	_ = v287
	var v291 int32
	_ = v291
	var v293 int32
	_ = v293
	var v296 int32
	_ = v296
	var v302 int32
	_ = v302
	var v308 int32
	_ = v308
	var v314 int32
	_ = v314
	var v325 int32
	_ = v325
	var v326 int32
	_ = v326
	var v329 int32
	_ = v329
	var v337 int32
	_ = v337
	var v338 int32
	_ = v338
	var v341 int32
	_ = v341
	var v344 int32
	_ = v344
	var v345 int32
	_ = v345
	var v349 int32
	_ = v349
	var v353 int32
	_ = v353
	var v354 int32
	_ = v354
	var v355 int32
	_ = v355
	var v358 int32
	_ = v358
	var v360 int32
	_ = v360
	var v362 int64
	_ = v362
	var v363 int64
	_ = v363
	var v364 int64
	_ = v364
	var v365 int64
	_ = v365
	var v366 int64
	_ = v366
	var v367 int64
	_ = v367
	var v368 int64
	_ = v368
	var v370 int64
	_ = v370
	var v372 int64
	_ = v372
	var v374 int64
	_ = v374
	var v376 int64
	_ = v376
	var v378 int64
	_ = v378
	var v380 int32
	_ = v380
	var v381 int32
	_ = v381
	var v382 int32
	_ = v382
	var v385 int32
	_ = v385
	var v388 int32
	_ = v388
	var v389 int32
	_ = v389
	var v390 int32
	_ = v390
	var v392 int32
	_ = v392
	var v395 int32
	_ = v395
	var v396 int32
	_ = v396
	var v400 int32
	_ = v400
	var v404 int32
	_ = v404
	var v406 int32
	_ = v406
	var v413 int32
	_ = v413
	var v414 int32
	_ = v414
	var v415 int32
	_ = v415
	var v420 int32
	_ = v420
	var v425 int32
	_ = v425
	var v429 int32
	_ = v429
	var v433 int32
	_ = v433
	var v445 int32
	_ = v445
	v8 = m.G0
	v10 = v8 - int32(16)
	m.G0 = v10
	v13 = *(*int32)(unsafe.Add(mBase, _consts[90]))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(v13)+36))
	v15 = *(*int32)(unsafe.Add(mBase, uint32(v14)+16))
	v16 = *(*int32)(unsafe.Add(mBase, uint32(v14)+12))
	F_addReplyArrayLen(m, l0, v15+v16)
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	F_clusterGenNodesSlotsInfo(m, int32(0))
	mBase = m.M
	v22 = m.ExcPending
	if v22 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v24 = *(*int32)(unsafe.Add(mBase, _consts[90]))
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v24)+36))
	v26 = F_dictGetSafeIterator(m, v25)
	mBase = m.M
	v27 = m.ExcPending
	if v27 != 0 {
		goto L1
	} else {
		goto L5
	}
L4:
	;
	F_dictReleaseIterator(m, v26)
	mBase = m.M
	v445 = m.ExcPending
	if v445 != 0 {
		goto L1
	} else {
		goto L113
	}
L5:
	;
	v35 = v26 + int32(20)
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v26)+16))
	if v36 != 0 {
		goto L8
	} else {
		goto L9
	}
L6:
	;
	if v131 == int32(0) {
		goto L4
	} else {
		goto L32
	}
L7:
	;
	v42 = v35
	v43 = v39
	goto L10
L8:
	;
	v39 = int32(1)
	goto L7
L9:
	;
	v39 = int32(0)
	goto L7
L10:
	;
	switch v43 {
	case 0:
		goto L15
	default:
		goto L14
	}
L12:
	;
	v43 = int32(0)
	goto L10
L13:
	;
	goto L6
L14:
	;
	v123 = *(*int32)(unsafe.Add(mBase, uint32(v42)))
	*(*int32)(unsafe.Add(mBase, uint32(v26)+16)) = v123
	if v123 == int32(0) {
		goto L12
	} else {
		goto L31
	}
L15:
	;
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v26)+4))
	if v47 != int32(-1) {
		v86 = v47
		goto L16
	} else {
		goto L17
	}
L16:
	;
	v87 = int32(1)
	v88 = v86 + v87
	*(*int32)(unsafe.Add(mBase, uint32(v26)+4)) = v88
	v90 = int32(0)
	v93 = *(*int32)(unsafe.Add(mBase, uint32(v26)))
	v94 = *(*int32)(unsafe.Add(mBase, uint32(v26)+8))
	v98 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v93+v94+int32(26)))))
	if v98 == int32(255) {
		goto L25
	} else {
		goto L26
	}
L17:
	;
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v26)+8))
	if v51 != 0 {
		v86 = int32(-1)
		goto L16
	} else {
		goto L18
	}
L18:
	;
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v26)))
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v26)+12))
	if v53 == int32(0) {
		goto L20
	} else {
		goto L21
	}
L19:
	;
	v80 = *(*int32)(unsafe.Add(mBase, uint32(v79)+20))
	if v80 != int32(-1) {
		goto L22
	} else {
		goto L23
	}
L20:
	;
	v60 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v52)+16)))
	v61 = int64(*(*int8)(unsafe.Add(mBase, uint32(v52)+27)))
	v62 = int64(*(*int32)(unsafe.Add(mBase, uint32(v52)+8)))
	v63 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v52)+12)))
	v64 = int64(*(*int8)(unsafe.Add(mBase, uint32(v52)+26)))
	v65 = int64(*(*int32)(unsafe.Add(mBase, uint32(v52)+4)))
	v66 = F_wangHash64(m, v65)
	mBase = m.M
	v68 = F_wangHash64(m, v64+v66)
	mBase = m.M
	v70 = F_wangHash64(m, v63+v68)
	mBase = m.M
	v72 = F_wangHash64(m, v62+v70)
	mBase = m.M
	v74 = F_wangHash64(m, v61+v72)
	mBase = m.M
	v76 = F_wangHash64(m, v60+v74)
	mBase = m.M
	*(*int64)(unsafe.Add(mBase, uint32(v26)+24)) = v76
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v26)))
	v79 = v78
	goto L19
L21:
	;
	v56 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v52)+24)))
	v58 = v56 + int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v52)+24)) = uint16(v58)
	v79 = v52
	goto L19
L22:
	;
	v86 = v80 + int32(-1)
	goto L16
L23:
	;
	v83 = *(*int32)(unsafe.Add(mBase, uint32(v26)+4))
	v86 = v83
	goto L16
L24:
	;
	v113 = int32(2)
	v118 = *(*int32)(unsafe.Add(mBase, uint32(v93+v111<<(uint(v113)%32)+int32(4))))
	v42 = v118 + v112<<(uint(v113)%32)
	v43 = int32(1)
	goto L10
L25:
	;
	v102 = v90
	goto L27
L26:
	;
	v102 = v87 << (uint(v98) % 32)
	goto L27
L27:
	;
	if v88 < v102 {
		v111 = v94
		v112 = v88
		goto L24
	} else {
		goto L28
	}
L28:
	;
	if v94 != 0 {
		v131 = v90
		goto L13
	} else {
		goto L29
	}
L29:
	;
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v93)+20))
	if v104 == int32(-1) {
		v131 = v90
		goto L13
	} else {
		goto L30
	}
L30:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v26)+4)) = int64(4294967296)
	v111 = int32(1)
	v112 = int32(0)
	goto L24
L31:
	;
	v127 = *(*int32)(unsafe.Add(mBase, uint32(v123)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v35))) = v127
	v131 = v123
	goto L13
L32:
	;
	v141 = v131
	goto L33
L33:
	;
	v144 = *(*int32)(unsafe.Add(mBase, uint32(v141)+8))
	goto L38
L34:
	;
	goto L4
L35:
	;
	F_addReplyBulkCString(m, l0, int32(_a312))
	mBase = m.M
	v325 = m.ExcPending
	if v325 != 0 {
		goto L1
	} else {
		goto L83
	}
L36:
	;
	F__serverAssert(m, int32(_a313), int32(_a179), int32(7329))
	mBase = m.M
	v314 = m.ExcPending
	if v314 != 0 {
		goto L1
	} else {
		goto L82
	}
L37:
	;
	F__serverAssert(m, int32(_a314), int32(_a179), int32(7313))
	mBase = m.M
	v308 = m.ExcPending
	if v308 != 0 {
		goto L1
	} else {
		goto L81
	}
L38:
	;
	v145 = *(*int32)(unsafe.Add(mBase, uint32(v144)+20))
	if v145 == int32(0) {
		goto L37
	} else {
		goto L39
	}
L39:
	;
	F_addReplyMapLen(m, l0, int32(3))
	mBase = m.M
	v150 = m.ExcPending
	if v150 != 0 {
		goto L1
	} else {
		goto L40
	}
L40:
	;
	F_addReplyBulkCString(m, l0, int32(_a315))
	mBase = m.M
	v153 = m.ExcPending
	if v153 != 0 {
		goto L1
	} else {
		goto L41
	}
L41:
	;
	v155 = v10 + int32(8)
	v156 = *(*int32)(unsafe.Add(mBase, uint32(v144)))
	*(*int32)(unsafe.Add(mBase, uint32(v155)+4)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v155))) = v156
	goto L42
L42:
	;
	v161 = v10 + int32(8)
	v163 = *(*int32)(unsafe.Add(mBase, uint32(v161)))
	if v163 == int32(0) {
		goto L46
	} else {
		goto L47
	}
L43:
	;
	F_addReplyBulkCString(m, l0, int32(_a316))
	mBase = m.M
	v249 = m.ExcPending
	if v249 != 0 {
		goto L1
	} else {
		goto L66
	}
L44:
	;
	F_addReplyArrayLen(m, l0, int32(0))
	mBase = m.M
	v239 = m.ExcPending
	if v239 != 0 {
		goto L1
	} else {
		goto L65
	}
L45:
	;
	if v163 == int32(0) {
		goto L44
	} else {
		goto L48
	}
L46:
	;
	goto L45
L47:
	;
	v166 = *(*int32)(unsafe.Add(mBase, uint32(v161)+4))
	v172 = *(*int32)(unsafe.Add(mBase, uint32(v163+base.B2i32(v166 == int32(0))<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v161))) = v172
	goto L46
L48:
	;
	v182 = v163
	goto L50
L49:
	;
	v203 = *(*int32)(unsafe.Add(mBase, uint32(v183)+2156))
	if v203&int32(1) != 0 {
		goto L36
	} else {
		goto L58
	}
L50:
	;
	v183 = *(*int32)(unsafe.Add(mBase, uint32(v182)+8))
	v184 = *(*int32)(unsafe.Add(mBase, uint32(v183)+2152))
	if v184 != 0 {
		goto L49
	} else {
		goto L52
	}
L51:
	;
	v199 = *(*int32)(unsafe.Add(mBase, uint32(v183)+2152))
	if v199 == int32(0) {
		goto L44
	} else {
		goto L57
	}
L52:
	;
	v186 = v10 + int32(8)
	v188 = *(*int32)(unsafe.Add(mBase, uint32(v186)))
	if v188 == int32(0) {
		goto L54
	} else {
		goto L55
	}
L53:
	;
	if v188 != 0 {
		v182 = v188
		goto L50
	} else {
		goto L56
	}
L54:
	;
	goto L53
L55:
	;
	v191 = *(*int32)(unsafe.Add(mBase, uint32(v186)+4))
	v197 = *(*int32)(unsafe.Add(mBase, uint32(v188+base.B2i32(v191 == int32(0))<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v186))) = v197
	goto L54
L56:
	;
	goto L51
L57:
	;
	goto L49
L58:
	;
	F_addReplyArrayLen(m, l0, v203)
	mBase = m.M
	v207 = m.ExcPending
	if v207 != 0 {
		goto L1
	} else {
		goto L59
	}
L59:
	;
	v209 = *(*int32)(unsafe.Add(mBase, uint32(v183)+2156))
	if v209 < int32(1) {
		goto L43
	} else {
		goto L60
	}
L60:
	;
	v218 = int32(0)
	goto L61
L61:
	;
	v219 = *(*int32)(unsafe.Add(mBase, uint32(v183)+2152))
	v223 = int64(*(*uint16)(unsafe.Add(mBase, uint32(v219+v218<<(uint(int32(1))%32)))))
	F_addReplyLongLong(m, l0, v223)
	mBase = m.M
	v225 = m.ExcPending
	if v225 != 0 {
		goto L1
	} else {
		goto L63
	}
L63:
	;
	v227 = v218 + int32(1)
	v228 = *(*int32)(unsafe.Add(mBase, uint32(v183)+2156))
	if v227 < v228 {
		v218 = v227
		goto L61
	} else {
		goto L64
	}
L64:
	;
	goto L43
L65:
	;
	goto L43
L66:
	;
	v250 = *(*int32)(unsafe.Add(mBase, uint32(v144)+20))
	F_addReplyArrayLen(m, l0, v250)
	mBase = m.M
	v252 = m.ExcPending
	if v252 != 0 {
		goto L1
	} else {
		goto L67
	}
L67:
	;
	v254 = v10 + int32(8)
	v255 = *(*int32)(unsafe.Add(mBase, uint32(v144)))
	*(*int32)(unsafe.Add(mBase, uint32(v254)+4)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v254))) = v255
	goto L68
L68:
	;
	v260 = v10 + int32(8)
	v262 = *(*int32)(unsafe.Add(mBase, uint32(v260)))
	if v262 == int32(0) {
		goto L70
	} else {
		goto L71
	}
L69:
	;
	if v262 == int32(0) {
		goto L35
	} else {
		goto L72
	}
L70:
	;
	goto L69
L71:
	;
	v265 = *(*int32)(unsafe.Add(mBase, uint32(v260)+4))
	v271 = *(*int32)(unsafe.Add(mBase, uint32(v262+base.B2i32(v265 == int32(0))<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v260))) = v271
	goto L70
L72:
	;
	v277 = v262
	goto L73
L73:
	;
	v282 = *(*int32)(unsafe.Add(mBase, uint32(v277)+8))
	F_addNodeDetailsToShardReply(m, l0, v282)
	mBase = m.M
	v284 = m.ExcPending
	if v284 != 0 {
		goto L1
	} else {
		goto L75
	}
L75:
	;
	v285 = *(*int32)(unsafe.Add(mBase, uint32(v282)+2152))
	F_valkey_free(m, v285)
	mBase = m.M
	v287 = m.ExcPending
	if v287 != 0 {
		goto L1
	} else {
		goto L76
	}
L76:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v282)+2152)) = int64(0)
	v291 = v10 + int32(8)
	v293 = *(*int32)(unsafe.Add(mBase, uint32(v291)))
	if v293 == int32(0) {
		goto L78
	} else {
		goto L79
	}
L77:
	;
	if v293 != 0 {
		v277 = v293
		goto L73
	} else {
		goto L80
	}
L78:
	;
	goto L77
L79:
	;
	v296 = *(*int32)(unsafe.Add(mBase, uint32(v291)+4))
	v302 = *(*int32)(unsafe.Add(mBase, uint32(v293+base.B2i32(v296 == int32(0))<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v291))) = v302
	goto L78
L80:
	;
	goto L35
L81:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L82:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L83:
	;
	v326 = *(*int32)(unsafe.Add(mBase, uint32(v141)))
	goto L84
L84:
	;
	F_addReplyBulkCBuffer(m, l0, v326, int32(40))
	mBase = m.M
	v329 = m.ExcPending
	if v329 != 0 {
		goto L1
	} else {
		goto L85
	}
L85:
	;
	v337 = v26 + int32(20)
	v338 = *(*int32)(unsafe.Add(mBase, uint32(v26)+16))
	if v338 != 0 {
		goto L88
	} else {
		goto L89
	}
L86:
	;
	if v433 != 0 {
		v141 = v433
		goto L33
	} else {
		goto L112
	}
L87:
	;
	v344 = v337
	v345 = v341
	goto L90
L88:
	;
	v341 = int32(1)
	goto L87
L89:
	;
	v341 = int32(0)
	goto L87
L90:
	;
	switch v345 {
	case 0:
		goto L95
	default:
		goto L94
	}
L92:
	;
	v345 = int32(0)
	goto L90
L93:
	;
	goto L86
L94:
	;
	v425 = *(*int32)(unsafe.Add(mBase, uint32(v344)))
	*(*int32)(unsafe.Add(mBase, uint32(v26)+16)) = v425
	if v425 == int32(0) {
		goto L92
	} else {
		goto L111
	}
L95:
	;
	v349 = *(*int32)(unsafe.Add(mBase, uint32(v26)+4))
	if v349 != int32(-1) {
		v388 = v349
		goto L96
	} else {
		goto L97
	}
L96:
	;
	v389 = int32(1)
	v390 = v388 + v389
	*(*int32)(unsafe.Add(mBase, uint32(v26)+4)) = v390
	v392 = int32(0)
	v395 = *(*int32)(unsafe.Add(mBase, uint32(v26)))
	v396 = *(*int32)(unsafe.Add(mBase, uint32(v26)+8))
	v400 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v395+v396+int32(26)))))
	if v400 == int32(255) {
		goto L105
	} else {
		goto L106
	}
L97:
	;
	v353 = *(*int32)(unsafe.Add(mBase, uint32(v26)+8))
	if v353 != 0 {
		v388 = int32(-1)
		goto L96
	} else {
		goto L98
	}
L98:
	;
	v354 = *(*int32)(unsafe.Add(mBase, uint32(v26)))
	v355 = *(*int32)(unsafe.Add(mBase, uint32(v26)+12))
	if v355 == int32(0) {
		goto L100
	} else {
		goto L101
	}
L99:
	;
	v382 = *(*int32)(unsafe.Add(mBase, uint32(v381)+20))
	if v382 != int32(-1) {
		goto L102
	} else {
		goto L103
	}
L100:
	;
	v362 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v354)+16)))
	v363 = int64(*(*int8)(unsafe.Add(mBase, uint32(v354)+27)))
	v364 = int64(*(*int32)(unsafe.Add(mBase, uint32(v354)+8)))
	v365 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v354)+12)))
	v366 = int64(*(*int8)(unsafe.Add(mBase, uint32(v354)+26)))
	v367 = int64(*(*int32)(unsafe.Add(mBase, uint32(v354)+4)))
	v368 = F_wangHash64(m, v367)
	mBase = m.M
	v370 = F_wangHash64(m, v366+v368)
	mBase = m.M
	v372 = F_wangHash64(m, v365+v370)
	mBase = m.M
	v374 = F_wangHash64(m, v364+v372)
	mBase = m.M
	v376 = F_wangHash64(m, v363+v374)
	mBase = m.M
	v378 = F_wangHash64(m, v362+v376)
	mBase = m.M
	*(*int64)(unsafe.Add(mBase, uint32(v26)+24)) = v378
	v380 = *(*int32)(unsafe.Add(mBase, uint32(v26)))
	v381 = v380
	goto L99
L101:
	;
	v358 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v354)+24)))
	v360 = v358 + int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v354)+24)) = uint16(v360)
	v381 = v354
	goto L99
L102:
	;
	v388 = v382 + int32(-1)
	goto L96
L103:
	;
	v385 = *(*int32)(unsafe.Add(mBase, uint32(v26)+4))
	v388 = v385
	goto L96
L104:
	;
	v415 = int32(2)
	v420 = *(*int32)(unsafe.Add(mBase, uint32(v395+v413<<(uint(v415)%32)+int32(4))))
	v344 = v420 + v414<<(uint(v415)%32)
	v345 = int32(1)
	goto L90
L105:
	;
	v404 = v392
	goto L107
L106:
	;
	v404 = v389 << (uint(v400) % 32)
	goto L107
L107:
	;
	if v390 < v404 {
		v413 = v396
		v414 = v390
		goto L104
	} else {
		goto L108
	}
L108:
	;
	if v396 != 0 {
		v433 = v392
		goto L93
	} else {
		goto L109
	}
L109:
	;
	v406 = *(*int32)(unsafe.Add(mBase, uint32(v395)+20))
	if v406 == int32(-1) {
		v433 = v392
		goto L93
	} else {
		goto L110
	}
L110:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v26)+4)) = int64(4294967296)
	v413 = int32(1)
	v414 = int32(0)
	goto L104
L111:
	;
	v429 = *(*int32)(unsafe.Add(mBase, uint32(v425)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v337))) = v429
	v433 = v425
	goto L93
L112:
	;
	goto L34
L113:
	;
	m.G0 = v10 + int32(16)
	return
}
func F_clusterCommandSyncSlotsAck(m *base.Module, l0 int32) {
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
	var v13 int32
	_ = v13
	var v15 int64
	_ = v15
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v31 int32
	_ = v31
	var v34 int64
	_ = v34
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+216))
	if v9 != 0 {
		v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+216))
		v15 = *(*int64)(unsafe.Add(mBase, _consts[37]))
		*(*int64)(unsafe.Add(mBase, uint32(v13)+24)) = v15
		v17 = *(*int32)(unsafe.Add(mBase, uint32(v13)+156))
		if v17 != 0 {
			m.G0 = v7 + int32(16)
			return
		} else {
			v19 = *(*int32)(unsafe.Add(mBase, _consts[15]))
			if int32(2) < v19 {
				v34 = *(*int64)(unsafe.Add(mBase, _consts[37]))
				*(*int32)(unsafe.Add(mBase, uint32(v13)+156)) = int32(1)
				*(*int64)(unsafe.Add(mBase, uint32(v13)+16)) = v34
				m.G0 = v7 + int32(16)
				return
			} else {
				v22 = *(*int32)(unsafe.Add(mBase, uint32(v13)+188))
				*(*int32)(unsafe.Add(mBase, uint32(v7)+8)) = int32(_a385)
				*(*int32)(unsafe.Add(mBase, uint32(v7)+4)) = int32(_a386)
				*(*int32)(unsafe.Add(mBase, uint32(v7))) = v22
				F__serverLog(m, int32(2), int32(_a330), v7)
				mBase = m.M
				v31 = m.ExcPending
				if v31 != 0 {
					return
				} else {
					v34 = *(*int64)(unsafe.Add(mBase, _consts[37]))
					*(*int32)(unsafe.Add(mBase, uint32(v13)+156)) = int32(1)
					*(*int64)(unsafe.Add(mBase, uint32(v13)+16)) = v34
					m.G0 = v7 + int32(16)
					return
				}
			}
		}
	} else {
		F_addReplyError(m, l0, int32(_a387))
		mBase = m.M
		v12 = m.ExcPending
		if v12 != 0 {
			return
		} else {
			m.G0 = v7 + int32(16)
			return
		}
	}
}
func F_clusterCommandSyncSlotsFailoverGranted(m *base.Module, l0 int32) {
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
	var v13 int32
	_ = v13
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v53 int32
	_ = v53
	var v56 int64
	_ = v56
	var v62 int32
	_ = v62
	v5 = m.G0
	v7 = v5 - int32(32)
	m.G0 = v7
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+216))
	if v9 != 0 {
		v13 = *(*int32)(unsafe.Add(mBase, uint32(v9)+156))
		switch v13 + int32(-3) {
		case 0:
			v39 = *(*int32)(unsafe.Add(mBase, _consts[15]))
			if int32(2) < v39 {
				v56 = *(*int64)(unsafe.Add(mBase, _consts[37]))
				*(*int32)(unsafe.Add(mBase, uint32(v9)+156)) = int32(4)
				*(*int64)(unsafe.Add(mBase, uint32(v9)+16)) = v56
				F_clusterDoBeforeSleep(m, int32(64))
				mBase = m.M
				v62 = m.ExcPending
				if v62 != 0 {
					return
				} else {
					m.G0 = v7 + int32(32)
					return
				}
			} else {
				v42 = *(*int32)(unsafe.Add(mBase, uint32(v9)+188))
				*(*int32)(unsafe.Add(mBase, uint32(v7)+24)) = int32(_a328)
				*(*int32)(unsafe.Add(mBase, uint32(v7)+20)) = int32(_a329)
				*(*int32)(unsafe.Add(mBase, uint32(v7)+16)) = v42
				F__serverLog(m, int32(2), int32(_a330), v7+int32(16))
				mBase = m.M
				v53 = m.ExcPending
				if v53 != 0 {
					return
				} else {
					v56 = *(*int64)(unsafe.Add(mBase, _consts[37]))
					*(*int32)(unsafe.Add(mBase, uint32(v9)+156)) = int32(4)
					*(*int64)(unsafe.Add(mBase, uint32(v9)+16)) = v56
					F_clusterDoBeforeSleep(m, int32(64))
					mBase = m.M
					v62 = m.ExcPending
					if v62 != 0 {
						return
					} else {
						m.G0 = v7 + int32(32)
						return
					}
				}
			}
		default:
			v23 = *(*int32)(unsafe.Add(mBase, _consts[15]))
			if int32(3) < v23 {
				v33 = v9
				F_finishSlotMigrationJob(m, v33, int32(18), int32(_a331))
				mBase = m.M
				v37 = m.ExcPending
				if v37 != 0 {
					return
				} else {
					m.G0 = v7 + int32(32)
					return
				}
			} else {
				v26 = *(*int32)(unsafe.Add(mBase, uint32(v9)+188))
				*(*int32)(unsafe.Add(mBase, uint32(v7))) = v26
				F__serverLog(m, int32(3), int32(_a332), v7)
				mBase = m.M
				v31 = m.ExcPending
				if v31 != 0 {
					return
				} else {
					v32 = *(*int32)(unsafe.Add(mBase, uint32(l0)+216))
					v33 = v32
					F_finishSlotMigrationJob(m, v33, int32(18), int32(_a331))
					mBase = m.M
					v37 = m.ExcPending
					if v37 != 0 {
						return
					} else {
						m.G0 = v7 + int32(32)
						return
					}
				}
			}
		case 2, 15, 16, 17:
			F__serverAssert(m, int32(_a333), int32(_a325), int32(715))
			mBase = m.M
			v20 = m.ExcPending
			if v20 != 0 {
				return
			} else {
				F_abort(m)
				mBase = m.M
				base.Wasm_trap_unreachable()
				for {
				}
			}
		}
	} else {
		F_addReplyError(m, l0, int32(_a334))
		mBase = m.M
		v12 = m.ExcPending
		if v12 != 0 {
			return
		} else {
			m.G0 = v7 + int32(32)
			return
		}
	}
}
func F_clusterCommandSyncSlotsFinish(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v16 int32
	_ = v16
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
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
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v77 int32
	_ = v77
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
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
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
	var v112 int32
	_ = v112
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
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
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v135 int32
	_ = v135
	var v142 int32
	_ = v142
	var v145 int32
	_ = v145
	var v148 int32
	_ = v148
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v158 int32
	_ = v158
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v169 int32
	_ = v169
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v176 int32
	_ = v176
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v182 int32
	_ = v182
	var v186 int32
	_ = v186
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v218 int32
	_ = v218
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v232 int32
	_ = v232
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v248 int32
	_ = v248
	var v260 int32
	_ = v260
	var v262 int32
	_ = v262
	var v263 int32
	_ = v263
	var v265 int32
	_ = v265
	var v272 int32
	_ = v272
	var v275 int32
	_ = v275
	var v278 int32
	_ = v278
	var v281 int32
	_ = v281
	var v282 int32
	_ = v282
	var v286 int32
	_ = v286
	var v287 int32
	_ = v287
	var v289 int32
	_ = v289
	var v290 int32
	_ = v290
	var v302 int32
	_ = v302
	var v304 int32
	_ = v304
	var v307 int32
	_ = v307
	var v313 int32
	_ = v313
	var v317 int32
	_ = v317
	var v319 int32
	_ = v319
	var v320 int32
	_ = v320
	var v328 int32
	_ = v328
	var v329 int32
	_ = v329
	var v330 int32
	_ = v330
	var v333 int32
	_ = v333
	var v334 int32
	_ = v334
	var v336 int32
	_ = v336
	var v337 int32
	_ = v337
	var v339 int32
	_ = v339
	var v341 int32
	_ = v341
	var v351 int32
	_ = v351
	var v352 int32
	_ = v352
	var v353 int32
	_ = v353
	var v356 int32
	_ = v356
	var v357 int32
	_ = v357
	var v358 int32
	_ = v358
	var v361 int32
	_ = v361
	var v362 int32
	_ = v362
	var v364 int32
	_ = v364
	var v369 int32
	_ = v369
	var v384 int32
	_ = v384
	var v394 int32
	_ = v394
	var v395 int32
	_ = v395
	var v398 int32
	_ = v398
	var v400 int32
	_ = v400
	var v401 int32
	_ = v401
	var v402 int32
	_ = v402
	var v404 int32
	_ = v404
	var v408 int32
	_ = v408
	var v410 int32
	_ = v410
	var v411 int32
	_ = v411
	var v412 int32
	_ = v412
	var v414 int32
	_ = v414
	var v415 int32
	_ = v415
	var v416 int32
	_ = v416
	var v417 int32
	_ = v417
	var v422 int32
	_ = v422
	var v423 int32
	_ = v423
	var v426 int32
	_ = v426
	var v427 int32
	_ = v427
	var v430 int32
	_ = v430
	var v431 int32
	_ = v431
	var v432 int32
	_ = v432
	var v435 int32
	_ = v435
	var v438 int32
	_ = v438
	var v440 int32
	_ = v440
	var v441 int32
	_ = v441
	var v442 int32
	_ = v442
	var v444 int32
	_ = v444
	var v448 int32
	_ = v448
	var v450 int32
	_ = v450
	var v451 int32
	_ = v451
	var v452 int32
	_ = v452
	var v454 int32
	_ = v454
	var v455 int32
	_ = v455
	var v456 int32
	_ = v456
	var v457 int32
	_ = v457
	var v462 int32
	_ = v462
	var v463 int32
	_ = v463
	var v466 int32
	_ = v466
	var v467 int32
	_ = v467
	var v470 int32
	_ = v470
	var v471 int32
	_ = v471
	var v472 int32
	_ = v472
	var v475 int32
	_ = v475
	var v477 int32
	_ = v477
	var v479 int32
	_ = v479
	var v482 int32
	_ = v482
	var v484 int32
	_ = v484
	var v486 int32
	_ = v486
	var v488 int32
	_ = v488
	v8 = m.G0
	v10 = v8 - int32(16)
	m.G0 = v10
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v12 < int32(4) {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	m.G0 = v10 + int32(16)
	return
L2:
	;
	v263 = int32(-1)
	v265 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v246+v263))))
	switch v265&int32(7) + v263 {
	case 0:
		goto L85
	case 1:
		goto L84
	case 2:
		goto L83
	case 3:
		goto L82
	default:
		goto L80
	}
L3:
	;
	v260 = *(*int32)(unsafe.Add(mBase, _consts[59]))
	F_addReplyErrorObject(m, l0, v260)
	mBase = m.M
	v262 = m.ExcPending
	if v262 != 0 {
		goto L26
	} else {
		goto L78
	}
L4:
	;
	v16 = int32(0)
	v21 = int32(3)
	v22 = v16
	v23 = v16
	v24 = v16
	goto L6
L5:
	;
	if v248 == int32(0) {
		goto L3
	} else {
		goto L76
	}
L6:
	;
	v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v28 = v21 << (uint(int32(2)) % 32)
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v26+v28)))
	v31 = F_objectGetVal(m, v30)
	mBase = m.M
	v32 = int32(_a335)
	v35 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v31))))
	if v35 != 0 {
		goto L12
	} else {
		goto L13
	}
L7:
	;
	v246 = v22
	v247 = v23
	v248 = v240
	goto L5
L8:
	;
	v235 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v236 = int32(2)
	v239 = *(*int32)(unsafe.Add(mBase, uint32(v235+v72<<(uint(v236)%32))))
	v240 = F_objectGetVal(m, v239)
	mBase = m.M
	v242 = v21 + v236
	v243 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v242 < v243 {
		v21 = v242
		v24 = v240
		goto L6
	} else {
		goto L75
	}
L9:
	;
	v80 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v82 = *(*int32)(unsafe.Add(mBase, uint32(v80+v28)))
	v83 = F_objectGetVal(m, v82)
	mBase = m.M
	v84 = int32(_a336)
	v87 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v83))))
	if v87 != 0 {
		goto L31
	} else {
		goto L32
	}
L10:
	;
	if v67-v69 != 0 {
		goto L9
	} else {
		goto L22
	}
L11:
	;
	v67 = F_tolower(m, v63)
	mBase = m.M
	v68 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v64))))
	v69 = F_tolower(m, v68)
	mBase = m.M
	goto L10
L12:
	;
	v37 = v31
	v38 = v32
	v39 = v35
	goto L15
L13:
	;
	v63 = int32(0)
	v64 = v32
	goto L11
L14:
	;
	v63 = v60 & int32(255)
	v64 = v59
	goto L11
L15:
	;
	v41 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v38))))
	if v41 == int32(0) {
		v59 = v38
		v60 = v39
		goto L14
	} else {
		goto L17
	}
L16:
	;
	v59 = v53
	v60 = int32(0)
	goto L14
L17:
	;
	v45 = v39 & int32(255)
	if v45 == v41 {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	v52 = int32(1)
	v53 = v38 + v52
	v54 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v37)+1)))
	if v54 != 0 {
		v37 = v37 + v52
		v38 = v53
		v39 = v54
		goto L15
	} else {
		goto L21
	}
L19:
	;
	v47 = F_tolower(m, v45)
	mBase = m.M
	v48 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v38))))
	v49 = F_tolower(m, v48)
	mBase = m.M
	if v47 == v49 {
		goto L18
	} else {
		goto L20
	}
L20:
	;
	v51 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v37))))
	v59 = v38
	v60 = v51
	goto L14
L21:
	;
	goto L16
L22:
	;
	if v24 != 0 {
		goto L23
	} else {
		goto L24
	}
L23:
	;
	v77 = *(*int32)(unsafe.Add(mBase, _consts[59]))
	F_addReplyErrorObject(m, l0, v77)
	mBase = m.M
	v79 = m.ExcPending
	if v79 != 0 {
		goto L26
	} else {
		goto L27
	}
L24:
	;
	v72 = v21 + int32(1)
	v73 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v72 < v73 {
		goto L8
	} else {
		goto L25
	}
L25:
	;
	goto L23
L26:
	;
	return
L27:
	;
	goto L1
L28:
	;
	v169 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v171 = *(*int32)(unsafe.Add(mBase, uint32(v169+v28)))
	v172 = F_objectGetVal(m, v171)
	mBase = m.M
	v173 = int32(_a337)
	v176 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v172))))
	if v176 != 0 {
		goto L57
	} else {
		goto L58
	}
L29:
	;
	if v119-v121 != 0 {
		goto L28
	} else {
		goto L41
	}
L30:
	;
	v119 = F_tolower(m, v115)
	mBase = m.M
	v120 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v116))))
	v121 = F_tolower(m, v120)
	mBase = m.M
	goto L29
L31:
	;
	v89 = v83
	v90 = v84
	v91 = v87
	goto L34
L32:
	;
	v115 = int32(0)
	v116 = v84
	goto L30
L33:
	;
	v115 = v112 & int32(255)
	v116 = v111
	goto L30
L34:
	;
	v93 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v90))))
	if v93 == int32(0) {
		v111 = v90
		v112 = v91
		goto L33
	} else {
		goto L36
	}
L35:
	;
	v111 = v105
	v112 = int32(0)
	goto L33
L36:
	;
	v97 = v91 & int32(255)
	if v97 == v93 {
		goto L37
	} else {
		goto L38
	}
L37:
	;
	v104 = int32(1)
	v105 = v90 + v104
	v106 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v89)+1)))
	if v106 != 0 {
		v89 = v89 + v104
		v90 = v105
		v91 = v106
		goto L34
	} else {
		goto L40
	}
L38:
	;
	v99 = F_tolower(m, v97)
	mBase = m.M
	v100 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v90))))
	v101 = F_tolower(m, v100)
	mBase = m.M
	if v99 == v101 {
		goto L37
	} else {
		goto L39
	}
L39:
	;
	v103 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v89))))
	v111 = v90
	v112 = v103
	goto L33
L40:
	;
	goto L35
L41:
	;
	if v22 != 0 {
		goto L43
	} else {
		goto L44
	}
L42:
	;
	v161 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v163 = *(*int32)(unsafe.Add(mBase, uint32(v161+v129)))
	v164 = F_objectGetVal(m, v163)
	mBase = m.M
	v166 = v21 + int32(2)
	v167 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v166 < v167 {
		v21 = v166
		v22 = v164
		goto L6
	} else {
		goto L53
	}
L43:
	;
	v158 = *(*int32)(unsafe.Add(mBase, _consts[59]))
	F_addReplyErrorObject(m, l0, v158)
	mBase = m.M
	v160 = m.ExcPending
	if v160 != 0 {
		goto L26
	} else {
		goto L52
	}
L44:
	;
	v124 = v21 + int32(1)
	v125 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v125 <= v124 {
		goto L43
	} else {
		goto L45
	}
L45:
	;
	v127 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v129 = v124 << (uint(int32(2)) % 32)
	v131 = *(*int32)(unsafe.Add(mBase, uint32(v127+v129)))
	v132 = F_objectGetVal(m, v131)
	mBase = m.M
	v133 = int32(-1)
	v135 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v132+v133))))
	switch v135&int32(7) + v133 {
	case 0:
		goto L50
	case 1:
		goto L49
	case 2:
		goto L48
	case 3:
		goto L47
	default:
		goto L43
	}
L46:
	;
	if v152 == int32(40) {
		goto L42
	} else {
		goto L51
	}
L47:
	;
	v151 = *(*int32)(unsafe.Add(mBase, uint32(v132+int32(-17))))
	v152 = v151
	goto L46
L48:
	;
	v148 = *(*int32)(unsafe.Add(mBase, uint32(v132+int32(-9))))
	v152 = v148
	goto L46
L49:
	;
	v145 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v132+int32(-5)))))
	v152 = v145
	goto L46
L50:
	;
	v142 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v132+int32(-3)))))
	v152 = v142
	goto L46
L51:
	;
	goto L43
L52:
	;
	goto L1
L53:
	;
	v246 = v164
	v247 = v23
	v248 = v24
	goto L5
L54:
	;
	v232 = *(*int32)(unsafe.Add(mBase, _consts[59]))
	F_addReplyErrorObject(m, l0, v232)
	mBase = m.M
	v234 = m.ExcPending
	if v234 != 0 {
		goto L26
	} else {
		goto L74
	}
L55:
	;
	if v208-v210 != 0 {
		goto L54
	} else {
		goto L67
	}
L56:
	;
	v208 = F_tolower(m, v204)
	mBase = m.M
	v209 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v205))))
	v210 = F_tolower(m, v209)
	mBase = m.M
	goto L55
L57:
	;
	v178 = v172
	v179 = v173
	v180 = v176
	goto L60
L58:
	;
	v204 = int32(0)
	v205 = v173
	goto L56
L59:
	;
	v204 = v201 & int32(255)
	v205 = v200
	goto L56
L60:
	;
	v182 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v179))))
	if v182 == int32(0) {
		v200 = v179
		v201 = v180
		goto L59
	} else {
		goto L62
	}
L61:
	;
	v200 = v194
	v201 = int32(0)
	goto L59
L62:
	;
	v186 = v180 & int32(255)
	if v186 == v182 {
		goto L63
	} else {
		goto L64
	}
L63:
	;
	v193 = int32(1)
	v194 = v179 + v193
	v195 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v178)+1)))
	if v195 != 0 {
		v178 = v178 + v193
		v179 = v194
		v180 = v195
		goto L60
	} else {
		goto L66
	}
L64:
	;
	v188 = F_tolower(m, v186)
	mBase = m.M
	v189 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v179))))
	v190 = F_tolower(m, v189)
	mBase = m.M
	if v188 == v190 {
		goto L63
	} else {
		goto L65
	}
L65:
	;
	v192 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v178))))
	v200 = v179
	v201 = v192
	goto L59
L66:
	;
	goto L61
L67:
	;
	if v23 != 0 {
		goto L69
	} else {
		goto L70
	}
L68:
	;
	v221 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v222 = int32(2)
	v225 = *(*int32)(unsafe.Add(mBase, uint32(v221+v213<<(uint(v222)%32))))
	v226 = F_objectGetVal(m, v225)
	mBase = m.M
	v228 = v21 + v222
	v229 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v228 < v229 {
		v21 = v228
		v23 = v226
		goto L6
	} else {
		goto L73
	}
L69:
	;
	v218 = *(*int32)(unsafe.Add(mBase, _consts[59]))
	F_addReplyErrorObject(m, l0, v218)
	mBase = m.M
	v220 = m.ExcPending
	if v220 != 0 {
		goto L26
	} else {
		goto L72
	}
L70:
	;
	v213 = v21 + int32(1)
	v214 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v213 < v214 {
		goto L68
	} else {
		goto L71
	}
L71:
	;
	goto L69
L72:
	;
	goto L1
L73:
	;
	v246 = v22
	v247 = v226
	v248 = v24
	goto L5
L74:
	;
	goto L1
L75:
	;
	goto L7
L76:
	;
	if v246 != 0 {
		goto L2
	} else {
		goto L77
	}
L77:
	;
	goto L3
L78:
	;
	goto L1
L79:
	;
	v395 = int32(_a338)
	v398 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v248))))
	if v398 != 0 {
		goto L118
	} else {
		goto L119
	}
L80:
	;
	F_addReplyError(m, l0, int32(_a339))
	mBase = m.M
	v394 = m.ExcPending
	if v394 != 0 {
		goto L26
	} else {
		goto L112
	}
L81:
	;
	if v282 != int32(40) {
		goto L80
	} else {
		goto L86
	}
L82:
	;
	v281 = *(*int32)(unsafe.Add(mBase, uint32(v246+int32(-17))))
	v282 = v281
	goto L81
L83:
	;
	v278 = *(*int32)(unsafe.Add(mBase, uint32(v246+int32(-9))))
	v282 = v278
	goto L81
L84:
	;
	v275 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v246+int32(-5)))))
	v282 = v275
	goto L81
L85:
	;
	v272 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v246+int32(-3)))))
	v282 = v272
	goto L81
L86:
	;
	v286 = *(*int32)(unsafe.Add(mBase, _consts[90]))
	v287 = *(*int32)(unsafe.Add(mBase, uint32(v286)+uint32(_consts[143])))
	v289 = v10 + int32(8)
	v290 = *(*int32)(unsafe.Add(mBase, uint32(v287)))
	*(*int32)(unsafe.Add(mBase, uint32(v289)+4)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v289))) = v290
	goto L87
L87:
	;
	goto L88
L88:
	;
	v302 = v10 + int32(8)
	v304 = *(*int32)(unsafe.Add(mBase, uint32(v302)))
	if v304 == int32(0) {
		goto L91
	} else {
		goto L92
	}
L89:
	;
	if v317 != 0 {
		goto L79
	} else {
		goto L111
	}
L90:
	;
	if v304 == int32(0) {
		goto L80
	} else {
		goto L93
	}
L91:
	;
	goto L90
L92:
	;
	v307 = *(*int32)(unsafe.Add(mBase, uint32(v302)+4))
	v313 = *(*int32)(unsafe.Add(mBase, uint32(v304+base.B2i32(v307 == int32(0))<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v302))) = v313
	goto L91
L93:
	;
	v317 = *(*int32)(unsafe.Add(mBase, uint32(v304)+8))
	v319 = v317 + int32(112)
	v320 = int32(40)
	goto L98
L94:
	;
	if v384 != 0 {
		goto L88
	} else {
		goto L110
	}
L95:
	;
	v384 = int32(0)
	goto L94
L96:
	;
	v356 = v351
	v357 = v352
	v358 = v353
	goto L106
L97:
	;
	if v341 == int32(0) {
		goto L95
	} else {
		goto L104
	}
L98:
	;
	if (v319|v246)&int32(3) != 0 {
		v351 = v246
		v352 = v319
		v353 = v320
		goto L96
	} else {
		goto L99
	}
L99:
	;
	v328 = v246
	v329 = v319
	v330 = v320
	goto L100
L100:
	;
	v333 = *(*int32)(unsafe.Add(mBase, uint32(v328)))
	v334 = *(*int32)(unsafe.Add(mBase, uint32(v329)))
	if v333 != v334 {
		v351 = v328
		v352 = v329
		v353 = v330
		goto L96
	} else {
		goto L102
	}
L101:
	;
	goto L97
L102:
	;
	v336 = int32(4)
	v337 = v329 + v336
	v339 = v328 + v336
	v341 = v330 + int32(-4)
	if base.Ui32(int32(3)) < base.Ui32(v341) {
		v328 = v339
		v329 = v337
		v330 = v341
		goto L100
	} else {
		goto L103
	}
L103:
	;
	goto L101
L104:
	;
	v351 = v339
	v352 = v337
	v353 = v341
	goto L96
L105:
	;
	v384 = v361 - v362
	goto L94
L106:
	;
	v361 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v356))))
	v362 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v357))))
	if v361 != v362 {
		goto L105
	} else {
		goto L108
	}
L108:
	;
	v364 = int32(1)
	v369 = v358 + int32(-1)
	if v369 == int32(0) {
		goto L95
	} else {
		goto L109
	}
L109:
	;
	v356 = v356 + v364
	v357 = v357 + v364
	v358 = v369
	goto L106
L110:
	;
	goto L89
L111:
	;
	goto L80
L112:
	;
	goto L1
L113:
	;
	v486 = *(*int32)(unsafe.Add(mBase, _consts[59]))
	F_addReplyErrorObject(m, l0, v486)
	mBase = m.M
	v488 = m.ExcPending
	if v488 != 0 {
		goto L26
	} else {
		goto L145
	}
L114:
	;
	v477 = *(*int32)(unsafe.Add(mBase, _consts[77]))
	F_addReply(m, l0, v477)
	mBase = m.M
	v479 = m.ExcPending
	if v479 != 0 {
		goto L26
	} else {
		goto L142
	}
L115:
	;
	v435 = int32(_a340)
	v438 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v248))))
	if v438 != 0 {
		goto L131
	} else {
		goto L132
	}
L116:
	;
	if v430-v432 != 0 {
		goto L115
	} else {
		goto L128
	}
L117:
	;
	v430 = F_tolower(m, v426)
	mBase = m.M
	v431 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v427))))
	v432 = F_tolower(m, v431)
	mBase = m.M
	goto L116
L118:
	;
	v400 = v248
	v401 = v395
	v402 = v398
	goto L121
L119:
	;
	v426 = int32(0)
	v427 = v395
	goto L117
L120:
	;
	v426 = v423 & int32(255)
	v427 = v422
	goto L117
L121:
	;
	v404 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v401))))
	if v404 == int32(0) {
		v422 = v401
		v423 = v402
		goto L120
	} else {
		goto L123
	}
L122:
	;
	v422 = v416
	v423 = int32(0)
	goto L120
L123:
	;
	v408 = v402 & int32(255)
	if v408 == v404 {
		goto L124
	} else {
		goto L125
	}
L124:
	;
	v415 = int32(1)
	v416 = v401 + v415
	v417 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v400)+1)))
	if v417 != 0 {
		v400 = v400 + v415
		v401 = v416
		v402 = v417
		goto L121
	} else {
		goto L127
	}
L125:
	;
	v410 = F_tolower(m, v408)
	mBase = m.M
	v411 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v401))))
	v412 = F_tolower(m, v411)
	mBase = m.M
	if v410 == v412 {
		goto L124
	} else {
		goto L126
	}
L126:
	;
	v414 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v400))))
	v422 = v401
	v423 = v414
	goto L120
L127:
	;
	goto L122
L128:
	;
	v475 = int32(20)
	goto L114
L129:
	;
	if v470-v472 != 0 {
		goto L113
	} else {
		goto L141
	}
L130:
	;
	v470 = F_tolower(m, v466)
	mBase = m.M
	v471 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v467))))
	v472 = F_tolower(m, v471)
	mBase = m.M
	goto L129
L131:
	;
	v440 = v248
	v441 = v435
	v442 = v438
	goto L134
L132:
	;
	v466 = int32(0)
	v467 = v435
	goto L130
L133:
	;
	v466 = v463 & int32(255)
	v467 = v462
	goto L130
L134:
	;
	v444 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v441))))
	if v444 == int32(0) {
		v462 = v441
		v463 = v442
		goto L133
	} else {
		goto L136
	}
L135:
	;
	v462 = v456
	v463 = int32(0)
	goto L133
L136:
	;
	v448 = v442 & int32(255)
	if v448 == v444 {
		goto L137
	} else {
		goto L138
	}
L137:
	;
	v455 = int32(1)
	v456 = v441 + v455
	v457 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v440)+1)))
	if v457 != 0 {
		v440 = v440 + v455
		v441 = v456
		v442 = v457
		goto L134
	} else {
		goto L140
	}
L138:
	;
	v450 = F_tolower(m, v448)
	mBase = m.M
	v451 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v441))))
	v452 = F_tolower(m, v451)
	mBase = m.M
	if v450 == v452 {
		goto L137
	} else {
		goto L139
	}
L139:
	;
	v454 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v440))))
	v462 = v441
	v463 = v454
	goto L133
L140:
	;
	goto L135
L141:
	;
	v475 = int32(18)
	goto L114
L142:
	;
	F_forceCommandPropagation(m, l0, int32(3))
	mBase = m.M
	v482 = m.ExcPending
	if v482 != 0 {
		goto L26
	} else {
		goto L143
	}
L143:
	;
	F_finishSlotMigrationJob(m, v317, v475, v247)
	mBase = m.M
	v484 = m.ExcPending
	if v484 != 0 {
		goto L26
	} else {
		goto L144
	}
L144:
	;
	goto L1
L145:
	;
	goto L1
}
func F_clusterCommandSyncSlotsRequestPause(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v77 int32
	_ = v77
	var v82 int64
	_ = v82
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v108 int32
	_ = v108
	var v113 int64
	_ = v113
	var v129 int32
	_ = v129
	v7 = m.G0
	v9 = v7 - int32(48)
	m.G0 = v9
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+216))
	if v11 != 0 {
		v15 = *(*int32)(unsafe.Add(mBase, uint32(v11)+156))
		if base.Ui32(int32(20)) < base.Ui32(v15) {
			if base.Ui32(v15+int32(-13)) < base.Ui32(int32(2)) {
				if v15 == int32(14) {
					v49 = v11
					v50 = F_slotExportTryDoPause(m, v49)
					mBase = m.M
					v51 = m.ExcPending
					if v51 != 0 {
						return
					} else {
						v53 = *(*int32)(unsafe.Add(mBase, _consts[15]))
						v54 = *(*int32)(unsafe.Add(mBase, uint32(l0)+216))
						if v50 != int32(-1) {
							if int32(2) < v53 {
								v113 = *(*int64)(unsafe.Add(mBase, _consts[37]))
								*(*int32)(unsafe.Add(mBase, uint32(v54)+156)) = int32(16)
								*(*int64)(unsafe.Add(mBase, uint32(v54)+16)) = v113
								m.G0 = v9 + int32(48)
								return
							} else {
								v88 = *(*int32)(unsafe.Add(mBase, uint32(v54)+188))
								v90 = *(*int32)(unsafe.Add(mBase, uint32(v54)+156))
								if base.Ui32(int32(20)) < base.Ui32(v90) {
									v98 = int32(_a242)
								} else {
									v97 = *(*int32)(unsafe.Add(mBase, uint32(v90<<(uint(int32(2))%32))+uint32(_consts[145])))
									v98 = v97
								}
								*(*int32)(unsafe.Add(mBase, uint32(v9)+24)) = int32(_a373)
								*(*int32)(unsafe.Add(mBase, uint32(v9)+20)) = v98
								*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = v88
								F__serverLog(m, int32(2), int32(_a330), v9+int32(16))
								mBase = m.M
								v108 = m.ExcPending
								if v108 != 0 {
									return
								} else {
									v113 = *(*int64)(unsafe.Add(mBase, _consts[37]))
									*(*int32)(unsafe.Add(mBase, uint32(v54)+156)) = int32(16)
									*(*int64)(unsafe.Add(mBase, uint32(v54)+16)) = v113
									m.G0 = v9 + int32(48)
									return
								}
							}
						} else {
							if int32(2) < v53 {
								v82 = *(*int64)(unsafe.Add(mBase, _consts[37]))
								*(*int32)(unsafe.Add(mBase, uint32(v54)+156)) = int32(15)
								*(*int64)(unsafe.Add(mBase, uint32(v54)+16)) = v82
								m.G0 = v9 + int32(48)
								return
							} else {
								v59 = *(*int32)(unsafe.Add(mBase, uint32(v54)+188))
								v61 = *(*int32)(unsafe.Add(mBase, uint32(v54)+156))
								if base.Ui32(int32(20)) < base.Ui32(v61) {
									v69 = int32(_a242)
								} else {
									v68 = *(*int32)(unsafe.Add(mBase, uint32(v61<<(uint(int32(2))%32))+uint32(_consts[145])))
									v69 = v68
								}
								*(*int32)(unsafe.Add(mBase, uint32(v9)+8)) = int32(_a374)
								*(*int32)(unsafe.Add(mBase, uint32(v9)+4)) = v69
								*(*int32)(unsafe.Add(mBase, uint32(v9))) = v59
								F__serverLog(m, int32(2), int32(_a330), v9)
								mBase = m.M
								v77 = m.ExcPending
								if v77 != 0 {
									return
								} else {
									v82 = *(*int64)(unsafe.Add(mBase, _consts[37]))
									*(*int32)(unsafe.Add(mBase, uint32(v54)+156)) = int32(15)
									*(*int64)(unsafe.Add(mBase, uint32(v54)+16)) = v82
									m.G0 = v9 + int32(48)
									return
								}
							}
						}
					}
				} else {
					F_slotExportBeginStreaming(m, v11)
					mBase = m.M
					v47 = m.ExcPending
					if v47 != 0 {
						return
					} else {
						v48 = *(*int32)(unsafe.Add(mBase, uint32(l0)+216))
						v49 = v48
						v50 = F_slotExportTryDoPause(m, v49)
						mBase = m.M
						v51 = m.ExcPending
						if v51 != 0 {
							return
						} else {
							v53 = *(*int32)(unsafe.Add(mBase, _consts[15]))
							v54 = *(*int32)(unsafe.Add(mBase, uint32(l0)+216))
							if v50 != int32(-1) {
								if int32(2) < v53 {
									v113 = *(*int64)(unsafe.Add(mBase, _consts[37]))
									*(*int32)(unsafe.Add(mBase, uint32(v54)+156)) = int32(16)
									*(*int64)(unsafe.Add(mBase, uint32(v54)+16)) = v113
									m.G0 = v9 + int32(48)
									return
								} else {
									v88 = *(*int32)(unsafe.Add(mBase, uint32(v54)+188))
									v90 = *(*int32)(unsafe.Add(mBase, uint32(v54)+156))
									if base.Ui32(int32(20)) < base.Ui32(v90) {
										v98 = int32(_a242)
									} else {
										v97 = *(*int32)(unsafe.Add(mBase, uint32(v90<<(uint(int32(2))%32))+uint32(_consts[145])))
										v98 = v97
									}
									*(*int32)(unsafe.Add(mBase, uint32(v9)+24)) = int32(_a373)
									*(*int32)(unsafe.Add(mBase, uint32(v9)+20)) = v98
									*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = v88
									F__serverLog(m, int32(2), int32(_a330), v9+int32(16))
									mBase = m.M
									v108 = m.ExcPending
									if v108 != 0 {
										return
									} else {
										v113 = *(*int64)(unsafe.Add(mBase, _consts[37]))
										*(*int32)(unsafe.Add(mBase, uint32(v54)+156)) = int32(16)
										*(*int64)(unsafe.Add(mBase, uint32(v54)+16)) = v113
										m.G0 = v9 + int32(48)
										return
									}
								}
							} else {
								if int32(2) < v53 {
									v82 = *(*int64)(unsafe.Add(mBase, _consts[37]))
									*(*int32)(unsafe.Add(mBase, uint32(v54)+156)) = int32(15)
									*(*int64)(unsafe.Add(mBase, uint32(v54)+16)) = v82
									m.G0 = v9 + int32(48)
									return
								} else {
									v59 = *(*int32)(unsafe.Add(mBase, uint32(v54)+188))
									v61 = *(*int32)(unsafe.Add(mBase, uint32(v54)+156))
									if base.Ui32(int32(20)) < base.Ui32(v61) {
										v69 = int32(_a242)
									} else {
										v68 = *(*int32)(unsafe.Add(mBase, uint32(v61<<(uint(int32(2))%32))+uint32(_consts[145])))
										v69 = v68
									}
									*(*int32)(unsafe.Add(mBase, uint32(v9)+8)) = int32(_a374)
									*(*int32)(unsafe.Add(mBase, uint32(v9)+4)) = v69
									*(*int32)(unsafe.Add(mBase, uint32(v9))) = v59
									F__serverLog(m, int32(2), int32(_a330), v9)
									mBase = m.M
									v77 = m.ExcPending
									if v77 != 0 {
										return
									} else {
										v82 = *(*int64)(unsafe.Add(mBase, _consts[37]))
										*(*int32)(unsafe.Add(mBase, uint32(v54)+156)) = int32(15)
										*(*int64)(unsafe.Add(mBase, uint32(v54)+16)) = v82
										m.G0 = v9 + int32(48)
										return
									}
								}
							}
						}
					}
				}
			} else {
				v27 = *(*int32)(unsafe.Add(mBase, _consts[15]))
				if int32(3) < v27 {
					v39 = v11
					F_finishSlotMigrationJob(m, v39, int32(18), int32(_a331))
					mBase = m.M
					v43 = m.ExcPending
					if v43 != 0 {
						return
					} else {
						m.G0 = v9 + int32(48)
						return
					}
				} else {
					v30 = *(*int32)(unsafe.Add(mBase, uint32(v11)+188))
					*(*int32)(unsafe.Add(mBase, uint32(v9)+32)) = v30
					F__serverLog(m, int32(3), int32(_a375), v9+int32(32))
					mBase = m.M
					v37 = m.ExcPending
					if v37 != 0 {
						return
					} else {
						v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)+216))
						v39 = v38
						F_finishSlotMigrationJob(m, v39, int32(18), int32(_a331))
						mBase = m.M
						v43 = m.ExcPending
						if v43 != 0 {
							return
						} else {
							m.G0 = v9 + int32(48)
							return
						}
					}
				}
			}
		} else {
			if int32(1)<<(uint(v15)%32)&int32(1835040) != 0 {
				F__serverAssert(m, int32(_a333), int32(_a325), int32(1486))
				mBase = m.M
				v129 = m.ExcPending
				if v129 != 0 {
					return
				} else {
					F_abort(m)
					mBase = m.M
					base.Wasm_trap_unreachable()
					for {
					}
				}
			} else {
				if base.Ui32(v15+int32(-13)) < base.Ui32(int32(2)) {
					if v15 == int32(14) {
						v49 = v11
						v50 = F_slotExportTryDoPause(m, v49)
						mBase = m.M
						v51 = m.ExcPending
						if v51 != 0 {
							return
						} else {
							v53 = *(*int32)(unsafe.Add(mBase, _consts[15]))
							v54 = *(*int32)(unsafe.Add(mBase, uint32(l0)+216))
							if v50 != int32(-1) {
								if int32(2) < v53 {
									v113 = *(*int64)(unsafe.Add(mBase, _consts[37]))
									*(*int32)(unsafe.Add(mBase, uint32(v54)+156)) = int32(16)
									*(*int64)(unsafe.Add(mBase, uint32(v54)+16)) = v113
									m.G0 = v9 + int32(48)
									return
								} else {
									v88 = *(*int32)(unsafe.Add(mBase, uint32(v54)+188))
									v90 = *(*int32)(unsafe.Add(mBase, uint32(v54)+156))
									if base.Ui32(int32(20)) < base.Ui32(v90) {
										v98 = int32(_a242)
									} else {
										v97 = *(*int32)(unsafe.Add(mBase, uint32(v90<<(uint(int32(2))%32))+uint32(_consts[145])))
										v98 = v97
									}
									*(*int32)(unsafe.Add(mBase, uint32(v9)+24)) = int32(_a373)
									*(*int32)(unsafe.Add(mBase, uint32(v9)+20)) = v98
									*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = v88
									F__serverLog(m, int32(2), int32(_a330), v9+int32(16))
									mBase = m.M
									v108 = m.ExcPending
									if v108 != 0 {
										return
									} else {
										v113 = *(*int64)(unsafe.Add(mBase, _consts[37]))
										*(*int32)(unsafe.Add(mBase, uint32(v54)+156)) = int32(16)
										*(*int64)(unsafe.Add(mBase, uint32(v54)+16)) = v113
										m.G0 = v9 + int32(48)
										return
									}
								}
							} else {
								if int32(2) < v53 {
									v82 = *(*int64)(unsafe.Add(mBase, _consts[37]))
									*(*int32)(unsafe.Add(mBase, uint32(v54)+156)) = int32(15)
									*(*int64)(unsafe.Add(mBase, uint32(v54)+16)) = v82
									m.G0 = v9 + int32(48)
									return
								} else {
									v59 = *(*int32)(unsafe.Add(mBase, uint32(v54)+188))
									v61 = *(*int32)(unsafe.Add(mBase, uint32(v54)+156))
									if base.Ui32(int32(20)) < base.Ui32(v61) {
										v69 = int32(_a242)
									} else {
										v68 = *(*int32)(unsafe.Add(mBase, uint32(v61<<(uint(int32(2))%32))+uint32(_consts[145])))
										v69 = v68
									}
									*(*int32)(unsafe.Add(mBase, uint32(v9)+8)) = int32(_a374)
									*(*int32)(unsafe.Add(mBase, uint32(v9)+4)) = v69
									*(*int32)(unsafe.Add(mBase, uint32(v9))) = v59
									F__serverLog(m, int32(2), int32(_a330), v9)
									mBase = m.M
									v77 = m.ExcPending
									if v77 != 0 {
										return
									} else {
										v82 = *(*int64)(unsafe.Add(mBase, _consts[37]))
										*(*int32)(unsafe.Add(mBase, uint32(v54)+156)) = int32(15)
										*(*int64)(unsafe.Add(mBase, uint32(v54)+16)) = v82
										m.G0 = v9 + int32(48)
										return
									}
								}
							}
						}
					} else {
						F_slotExportBeginStreaming(m, v11)
						mBase = m.M
						v47 = m.ExcPending
						if v47 != 0 {
							return
						} else {
							v48 = *(*int32)(unsafe.Add(mBase, uint32(l0)+216))
							v49 = v48
							v50 = F_slotExportTryDoPause(m, v49)
							mBase = m.M
							v51 = m.ExcPending
							if v51 != 0 {
								return
							} else {
								v53 = *(*int32)(unsafe.Add(mBase, _consts[15]))
								v54 = *(*int32)(unsafe.Add(mBase, uint32(l0)+216))
								if v50 != int32(-1) {
									if int32(2) < v53 {
										v113 = *(*int64)(unsafe.Add(mBase, _consts[37]))
										*(*int32)(unsafe.Add(mBase, uint32(v54)+156)) = int32(16)
										*(*int64)(unsafe.Add(mBase, uint32(v54)+16)) = v113
										m.G0 = v9 + int32(48)
										return
									} else {
										v88 = *(*int32)(unsafe.Add(mBase, uint32(v54)+188))
										v90 = *(*int32)(unsafe.Add(mBase, uint32(v54)+156))
										if base.Ui32(int32(20)) < base.Ui32(v90) {
											v98 = int32(_a242)
										} else {
											v97 = *(*int32)(unsafe.Add(mBase, uint32(v90<<(uint(int32(2))%32))+uint32(_consts[145])))
											v98 = v97
										}
										*(*int32)(unsafe.Add(mBase, uint32(v9)+24)) = int32(_a373)
										*(*int32)(unsafe.Add(mBase, uint32(v9)+20)) = v98
										*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = v88
										F__serverLog(m, int32(2), int32(_a330), v9+int32(16))
										mBase = m.M
										v108 = m.ExcPending
										if v108 != 0 {
											return
										} else {
											v113 = *(*int64)(unsafe.Add(mBase, _consts[37]))
											*(*int32)(unsafe.Add(mBase, uint32(v54)+156)) = int32(16)
											*(*int64)(unsafe.Add(mBase, uint32(v54)+16)) = v113
											m.G0 = v9 + int32(48)
											return
										}
									}
								} else {
									if int32(2) < v53 {
										v82 = *(*int64)(unsafe.Add(mBase, _consts[37]))
										*(*int32)(unsafe.Add(mBase, uint32(v54)+156)) = int32(15)
										*(*int64)(unsafe.Add(mBase, uint32(v54)+16)) = v82
										m.G0 = v9 + int32(48)
										return
									} else {
										v59 = *(*int32)(unsafe.Add(mBase, uint32(v54)+188))
										v61 = *(*int32)(unsafe.Add(mBase, uint32(v54)+156))
										if base.Ui32(int32(20)) < base.Ui32(v61) {
											v69 = int32(_a242)
										} else {
											v68 = *(*int32)(unsafe.Add(mBase, uint32(v61<<(uint(int32(2))%32))+uint32(_consts[145])))
											v69 = v68
										}
										*(*int32)(unsafe.Add(mBase, uint32(v9)+8)) = int32(_a374)
										*(*int32)(unsafe.Add(mBase, uint32(v9)+4)) = v69
										*(*int32)(unsafe.Add(mBase, uint32(v9))) = v59
										F__serverLog(m, int32(2), int32(_a330), v9)
										mBase = m.M
										v77 = m.ExcPending
										if v77 != 0 {
											return
										} else {
											v82 = *(*int64)(unsafe.Add(mBase, _consts[37]))
											*(*int32)(unsafe.Add(mBase, uint32(v54)+156)) = int32(15)
											*(*int64)(unsafe.Add(mBase, uint32(v54)+16)) = v82
											m.G0 = v9 + int32(48)
											return
										}
									}
								}
							}
						}
					}
				} else {
					v27 = *(*int32)(unsafe.Add(mBase, _consts[15]))
					if int32(3) < v27 {
						v39 = v11
						F_finishSlotMigrationJob(m, v39, int32(18), int32(_a331))
						mBase = m.M
						v43 = m.ExcPending
						if v43 != 0 {
							return
						} else {
							m.G0 = v9 + int32(48)
							return
						}
					} else {
						v30 = *(*int32)(unsafe.Add(mBase, uint32(v11)+188))
						*(*int32)(unsafe.Add(mBase, uint32(v9)+32)) = v30
						F__serverLog(m, int32(3), int32(_a375), v9+int32(32))
						mBase = m.M
						v37 = m.ExcPending
						if v37 != 0 {
							return
						} else {
							v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)+216))
							v39 = v38
							F_finishSlotMigrationJob(m, v39, int32(18), int32(_a331))
							mBase = m.M
							v43 = m.ExcPending
							if v43 != 0 {
								return
							} else {
								m.G0 = v9 + int32(48)
								return
							}
						}
					}
				}
			}
		}
	} else {
		F_addReplyError(m, l0, int32(_a376))
		mBase = m.M
		v14 = m.ExcPending
		if v14 != 0 {
			return
		} else {
			m.G0 = v9 + int32(48)
			return
		}
	}
}
func F_clusterDelNodeSlots(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
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
	var v36 int32
	_ = v36
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
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+2160))
	if int32(1) <= v7 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v14 = int32(0)
	v17 = v7
	v19 = v14
	v20 = v14
	goto L3
L2:
	;
	return int32(0)
L3:
	;
	v23 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0+int32(104)+v20))))
	if v23 == int32(0) {
		v50 = v17
		v52 = v19
		goto L5
	} else {
		goto L6
	}
L4:
	;
	return v52
L5:
	;
	if base.Ui32(int32(2046)) < base.Ui32(v20) {
		goto L12
	} else {
		goto L13
	}
L6:
	;
	v28 = v23
	v29 = v17
	v31 = v19
	goto L7
L7:
	;
	v36 = F_clusterDelSlot(m, v20<<(uint(int32(3))%32)|base.I32_ctz(v28))
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L9
	} else {
		goto L10
	}
L8:
	;
	v50 = v41
	v52 = v43
	goto L5
L9:
	;
	return int32(0)
L10:
	;
	v40 = int32(-1)
	v41 = v29 + v40
	v43 = v31 + int32(1)
	v46 = (v28 + v40) & v28
	if v46&int32(255) != 0 {
		v28 = v46
		v29 = v41
		v31 = v43
		goto L7
	} else {
		goto L11
	}
L11:
	;
	goto L8
L12:
	;
	goto L4
L13:
	;
	if int32(0) < v50 {
		v17 = v50
		v19 = v52
		v20 = v20 + int32(1)
		goto L3
	} else {
		goto L14
	}
L14:
	;
	goto L12
}
func F_clusterDoBeforeSleep(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	if l0&int32(4) == int32(0) {
		v10 = *(*int32)(unsafe.Add(mBase, _consts[90]))
		v11 = *(*int32)(unsafe.Add(mBase, uint32(v10)+uint32(_consts[101])))
		*(*int32)(unsafe.Add(mBase, uint32(v10)+uint32(_consts[101]))) = v11 | l0
		return
	} else {
		F_clearCachedClusterSlotsResponse(m)
		mBase = m.M
		v8 = m.ExcPending
		if v8 != 0 {
			return
		} else {
			v10 = *(*int32)(unsafe.Add(mBase, _consts[90]))
			v11 = *(*int32)(unsafe.Add(mBase, uint32(v10)+uint32(_consts[101])))
			*(*int32)(unsafe.Add(mBase, uint32(v10)+uint32(_consts[101]))) = v11 | l0
			return
		}
	}
}
func F_clusterFailAllSlotExportsWithMessage(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v30 int32
	_ = v30
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v63 int32
	_ = v63
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v10 = *(*int32)(unsafe.Add(mBase, _consts[90]))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(v10)+uint32(_consts[143])))
	v13 = v7 + int32(8)
	v14 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
	*(*int32)(unsafe.Add(mBase, uint32(v13)+4)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v13))) = v14
	goto L1
L1:
	;
	v19 = v7 + int32(8)
	v21 = *(*int32)(unsafe.Add(mBase, uint32(v19)))
	if v21 == int32(0) {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	m.G0 = v7 + int32(16)
	return
L3:
	;
	if v21 == int32(0) {
		goto L2
	} else {
		goto L6
	}
L4:
	;
	goto L3
L5:
	;
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v19)+4))
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v21+base.B2i32(v24 == int32(0))<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v19))) = v30
	goto L4
L6:
	;
	v36 = v21
	goto L7
L7:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v36)+8))
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v38)))
	if v39 != 0 {
		goto L9
	} else {
		goto L10
	}
L8:
	;
	goto L2
L9:
	;
	v52 = v7 + int32(8)
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v52)))
	if v54 == int32(0) {
		goto L17
	} else {
		goto L18
	}
L10:
	;
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v38)+156))
	if base.Ui32(int32(20)) < base.Ui32(v40) {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	F_finishSlotMigrationJob(m, v38, int32(18), l0)
	mBase = m.M
	v49 = m.ExcPending
	if v49 != 0 {
		goto L14
	} else {
		goto L15
	}
L12:
	;
	if int32(1)<<(uint(v40)%32)&int32(1835040) != 0 {
		goto L9
	} else {
		goto L13
	}
L13:
	;
	goto L11
L14:
	;
	return
L15:
	;
	goto L9
L16:
	;
	if v54 != 0 {
		v36 = v54
		goto L7
	} else {
		goto L19
	}
L17:
	;
	goto L16
L18:
	;
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v52)+4))
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v54+base.B2i32(v57 == int32(0))<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v52))) = v63
	goto L17
L19:
	;
	goto L8
}
func F_clusterFailoverReplaceYourPrimary(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v64 int32
	_ = v64
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v95 int32
	_ = v95
	var v97 int32
	_ = v97
	var v102 int32
	_ = v102
	var v104 int32
	_ = v104
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v120 int32
	_ = v120
	var v124 int32
	_ = v124
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v154 int32
	_ = v154
	var v157 int32
	_ = v157
	var v160 int32
	_ = v160
	var v163 int64
	_ = v163
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v180 int32
	_ = v180
	var v186 int32
	_ = v186
	var v212 int32
	_ = v212
	var v214 int32
	_ = v214
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v221 int32
	_ = v221
	var v226 int32
	_ = v226
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v241 int32
	_ = v241
	v12 = m.G0
	v14 = v12 - int32(16)
	m.G0 = v14
	v17 = *(*int32)(unsafe.Add(mBase, _consts[94]))
	v18 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17)+88)))
	if v18&int32(1) != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	m.G0 = v14 + int32(16)
	return
L2:
	;
	v21 = *(*int32)(unsafe.Add(mBase, uint32(v17)+2172))
	if v21 == int32(0) {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v25 = *(*int32)(unsafe.Add(mBase, _consts[15]))
	if int32(2) < v25 {
		v43 = v17
		goto L4
	} else {
		goto L5
	}
L4:
	;
	F_clusterSetNodeAsPrimary(m, v43)
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L6
	} else {
		goto L9
	}
L5:
	;
	v28 = F_humanNodename(m, v21)
	mBase = m.M
	v29 = m.ExcPending
	if v29 != 0 {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	return
L7:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+8)) = v28
	*(*int32)(unsafe.Add(mBase, uint32(v14)+4)) = v21 + int32(8)
	*(*int32)(unsafe.Add(mBase, uint32(v14))) = v17 + int32(48)
	F__serverLog(m, int32(2), int32(_a293), v14)
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
		goto L6
	} else {
		goto L8
	}
L8:
	;
	v42 = *(*int32)(unsafe.Add(mBase, _consts[94]))
	v43 = v42
	goto L4
L9:
	;
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v21)+2160))
	if v46 < int32(1) {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	F_clusterUpdateState(m)
	mBase = m.M
	v212 = m.ExcPending
	if v212 != 0 {
		goto L6
	} else {
		goto L37
	}
L11:
	;
	v55 = v46
	v57 = int32(0)
	goto L12
L12:
	;
	v64 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21+int32(104)+v57))))
	if v64 == int32(0) {
		v186 = v55
		goto L14
	} else {
		goto L15
	}
L13:
	;
	goto L10
L14:
	;
	if base.Ui32(int32(2046)) < base.Ui32(v57) {
		goto L10
	} else {
		goto L35
	}
L15:
	;
	v70 = v64
	v72 = v55
	goto L16
L16:
	;
	v80 = base.I32_ctz(v70)
	v81 = v57<<(uint(int32(3))%32) | v80
	v82 = F_clusterDelSlot(m, v81)
	mBase = m.M
	v83 = m.ExcPending
	if v83 != 0 {
		goto L6
	} else {
		goto L18
	}
L17:
	;
	v186 = v177
	goto L14
L18:
	;
	v85 = *(*int32)(unsafe.Add(mBase, _consts[90]))
	v87 = v81 << (uint(int32(2)) % 32)
	v89 = *(*int32)(unsafe.Add(mBase, uint32(v85+v87)+52))
	if v89 != 0 {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	v176 = int32(-1)
	v177 = v72 + v176
	v180 = (v70 + v176) & v70
	if v180&int32(255) != 0 {
		v70 = v180
		v72 = v177
		goto L16
	} else {
		goto L34
	}
L20:
	;
	v91 = *(*int32)(unsafe.Add(mBase, _consts[94]))
	v95 = m.G0
	v97 = v95 - int32(32)
	m.G0 = v97
	v102 = int32(1) << (uint(v81&int32(7)) % 32)
	v104 = base.I32_div_s(v81, int32(8))
	v107 = v91 + v104 + int32(104)
	v108 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v107))))
	if v102&v108 != 0 {
		goto L22
	} else {
		goto L23
	}
L21:
	;
	v146 = int32(_a69)
	v147 = *(*int32)(unsafe.Add(mBase, _consts[90]))
	*(*int32)(unsafe.Add(mBase, uint32(v147+v87)+52)) = v91
	v150 = v147 + v57
	v151 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v150)+uint32(_consts[96]))))
	v154 = v151 & base.I32_rotl(int32(-2), v80)
	*(*uint8)(unsafe.Add(mBase, uint32(v150)+uint32(_consts[96]))) = uint8(v154)
	v157 = *(*int32)(unsafe.Add(mBase, _consts[90]))
	v160 = v157 + v81*int32(24)
	v163 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v160)+uint32(_consts[97]))) = v163
	*(*int64)(unsafe.Add(mBase, uint32(v160)+uint32(_consts[98]))) = v163
	*(*int64)(unsafe.Add(mBase, uint32(v160)+uint32(_consts[99]))) = v163
	goto L33
L22:
	;
	m.G0 = v97 + int32(32)
	goto L21
L23:
	;
	v110 = v108 | v102
	*(*uint8)(unsafe.Add(mBase, uint32(v107))) = uint8(v110)
	v112 = *(*int32)(unsafe.Add(mBase, uint32(v91)+2160))
	*(*int32)(unsafe.Add(mBase, uint32(v91)+2160)) = v112 + int32(1)
	if v112 != 0 {
		goto L22
	} else {
		goto L24
	}
L24:
	;
	v117 = *(*int32)(unsafe.Add(mBase, _consts[90]))
	v118 = *(*int32)(unsafe.Add(mBase, uint32(v117)+32))
	F_dictInitIterator(m, v97, v118)
	mBase = m.M
	v120 = F_dictNext(m, v97)
	mBase = m.M
	if v120 == int32(0) {
		goto L22
	} else {
		goto L25
	}
L25:
	;
	v124 = v120
	goto L27
L26:
	;
	v134 = *(*int32)(unsafe.Add(mBase, uint32(v91)+88))
	*(*int32)(unsafe.Add(mBase, uint32(v91)+88)) = v134 | int32(256)
	goto L22
L27:
	;
	v128 = F_dictGetVal(m, v124)
	mBase = m.M
	v129 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v128)+88)))
	if v129&int32(2) != 0 {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	v133 = F_dictNext(m, v97)
	mBase = m.M
	if v133 != 0 {
		v124 = v133
		goto L27
	} else {
		goto L32
	}
L30:
	;
	v132 = *(*int32)(unsafe.Add(mBase, uint32(v128)+2164))
	if v132 != 0 {
		goto L26
	} else {
		goto L31
	}
L31:
	;
	goto L29
L32:
	;
	goto L22
L33:
	;
	goto L19
L34:
	;
	goto L17
L35:
	;
	if int32(0) < v186 {
		v55 = v186
		v57 = v57 + int32(1)
		goto L12
	} else {
		goto L36
	}
L36:
	;
	goto L13
L37:
	;
	F_clearCachedClusterSlotsResponse(m)
	mBase = m.M
	v214 = m.ExcPending
	if v214 != 0 {
		goto L6
	} else {
		goto L38
	}
L38:
	;
	v216 = *(*int32)(unsafe.Add(mBase, _consts[90]))
	v217 = *(*int32)(unsafe.Add(mBase, uint32(v216)+uint32(_consts[101])))
	*(*int32)(unsafe.Add(mBase, uint32(v216)+uint32(_consts[101]))) = v217 | int32(44)
	v221 = *(*int32)(unsafe.Add(mBase, uint32(v216)+uint32(_consts[111])))
	if v221 == int32(0) {
		v229 = v216
		goto L39
	} else {
		goto L40
	}
L39:
	;
	v230 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v229)+uint32(_consts[112]))) = v230
	*(*int64)(unsafe.Add(mBase, uint32(v229)+uint32(_consts[113]))) = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v229)+uint32(_consts[114]))) = int64(-1)
	*(*int32)(unsafe.Add(mBase, uint32(v229)+uint32(_consts[111]))) = v230
	v238 = F_verifyClusterConfigWithData(m)
	mBase = m.M
	v239 = m.ExcPending
	if v239 != 0 {
		goto L6
	} else {
		goto L42
	}
L40:
	;
	F_unpauseActions(m, int32(2))
	mBase = m.M
	v226 = m.ExcPending
	if v226 != 0 {
		goto L6
	} else {
		goto L41
	}
L41:
	;
	v228 = *(*int32)(unsafe.Add(mBase, _consts[90]))
	v229 = v228
	goto L39
L42:
	;
	v241 = *(*int32)(unsafe.Add(mBase, _consts[90]))
	*(*int64)(unsafe.Add(mBase, uint32(v241)+uint32(_consts[121]))) = int64(0)
	goto L1
}
func F_clusterGenNodeDescription(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v20 int32
	_ = v20
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v51 int32
	_ = v51
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
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v87 int32
	_ = v87
	var v94 int32
	_ = v94
	var v97 int32
	_ = v97
	var v100 int32
	_ = v100
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v114 int32
	_ = v114
	var v122 int32
	_ = v122
	var v124 int32
	_ = v124
	var v138 int32
	_ = v138
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v148 int32
	_ = v148
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v159 int32
	_ = v159
	var v168 int32
	_ = v168
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v199 int64
	_ = v199
	var v200 int64
	_ = v200
	var v201 int64
	_ = v201
	var v203 int32
	_ = v203
	var v206 int32
	_ = v206
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v227 int32
	_ = v227
	var v231 int32
	_ = v231
	var v235 int32
	_ = v235
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v248 int32
	_ = v248
	var v254 int32
	_ = v254
	var v255 int32
	_ = v255
	var v261 int32
	_ = v261
	var v262 int32
	_ = v262
	var v263 int32
	_ = v263
	var v265 int32
	_ = v265
	var v267 int32
	_ = v267
	var v274 int32
	_ = v274
	var v278 int32
	_ = v278
	var v279 int32
	_ = v279
	var v286 int32
	_ = v286
	var v288 int32
	_ = v288
	var v293 int32
	_ = v293
	var v294 int32
	_ = v294
	var v297 int32
	_ = v297
	var v302 int32
	_ = v302
	var v306 int32
	_ = v306
	var v307 int32
	_ = v307
	var v309 int32
	_ = v309
	var v315 int32
	_ = v315
	var v316 int32
	_ = v316
	var v322 int32
	_ = v322
	var v323 int32
	_ = v323
	var v324 int32
	_ = v324
	var v327 int32
	_ = v327
	var v328 int32
	_ = v328
	var v329 int32
	_ = v329
	var v339 int32
	_ = v339
	var v346 int32
	_ = v346
	var v352 int32
	_ = v352
	var v356 int32
	_ = v356
	var v364 int32
	_ = v364
	var v365 int32
	_ = v365
	var v366 int32
	_ = v366
	var v367 int32
	_ = v367
	var v369 int32
	_ = v369
	var v370 int32
	_ = v370
	var v372 int32
	_ = v372
	var v373 int32
	_ = v373
	var v374 int32
	_ = v374
	var v375 int32
	_ = v375
	var v378 int32
	_ = v378
	var v388 int32
	_ = v388
	var v390 int32
	_ = v390
	var v392 int32
	_ = v392
	var v393 int32
	_ = v393
	var v397 int32
	_ = v397
	var v398 int32
	_ = v398
	var v400 int32
	_ = v400
	var v401 int32
	_ = v401
	var v404 int32
	_ = v404
	var v407 int32
	_ = v407
	var v414 int32
	_ = v414
	v12 = m.G0
	v14 = v12 - int32(176)
	m.G0 = v14
	if l0 == int32(0) {
		goto L4
	} else {
		goto L5
	}
L1:
	;
	v33 = F_clusterNodeIp(m, l1, l0)
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
		goto L12
	} else {
		goto L13
	}
L2:
	;
	if l0 == int32(0) {
		goto L9
	} else {
		goto L10
	}
L3:
	;
	v26 = *(*int32)(unsafe.Add(mBase, uint32(l1)+2328))
	v32 = v26
	goto L1
L4:
	;
	if l2 == int32(0) {
		goto L2
	} else {
		goto L8
	}
L5:
	;
	if l2 == int32(0) {
		goto L4
	} else {
		goto L6
	}
L6:
	;
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l1)+2340))
	if v20 == int32(0) {
		goto L3
	} else {
		goto L7
	}
L7:
	;
	v32 = v20
	goto L1
L8:
	;
	goto L3
L9:
	;
	v31 = *(*int32)(unsafe.Add(mBase, uint32(l1)+2324))
	v32 = v31
	goto L1
L10:
	;
	v29 = *(*int32)(unsafe.Add(mBase, uint32(l1)+2336))
	if v29 != 0 {
		v32 = v29
		goto L1
	} else {
		goto L11
	}
L11:
	;
	goto L9
L12:
	;
	return int32(0)
L13:
	;
	v37 = F_sdsempty(m)
	mBase = m.M
	v38 = m.ExcPending
	if v38 != 0 {
		goto L12
	} else {
		goto L14
	}
L14:
	;
	v42 = F_sdscatlen(m, v37, l1+int32(8), int32(40))
	mBase = m.M
	v43 = m.ExcPending
	if v43 != 0 {
		goto L12
	} else {
		goto L15
	}
L15:
	;
	v44 = *(*int32)(unsafe.Add(mBase, uint32(l1)+2332))
	*(*int32)(unsafe.Add(mBase, uint32(v14)+168)) = v44
	*(*int32)(unsafe.Add(mBase, uint32(v14)+164)) = v32
	*(*int32)(unsafe.Add(mBase, uint32(v14)+160)) = v33
	v51 = F_sdscatfmt(m, v42, int32(_a213), v14+int32(160))
	mBase = m.M
	v52 = m.ExcPending
	if v52 != 0 {
		goto L12
	} else {
		goto L16
	}
L16:
	;
	v53 = *(*int32)(unsafe.Add(mBase, uint32(l1)+2312))
	v56 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v53+int32(-1)))))
	switch v56 & int32(7) {
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
		v83 = v51
		goto L17
	}
L17:
	;
	if l0 != 0 {
		v168 = v83
		goto L26
	} else {
		goto L27
	}
L18:
	;
	if v73 == int32(0) {
		v83 = v51
		goto L17
	} else {
		goto L24
	}
L19:
	;
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v53+int32(-17))))
	v73 = v72
	goto L18
L20:
	;
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v53+int32(-9))))
	v73 = v69
	goto L18
L21:
	;
	v66 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v53+int32(-5)))))
	v73 = v66
	goto L18
L22:
	;
	v63 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v53+int32(-3)))))
	v73 = v63
	goto L18
L23:
	;
	v73 = int32(base.Ui32(v56) >> (uint(int32(3)) % 32))
	goto L18
L24:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+144)) = v53
	v80 = F_sdscatfmt(m, v51, int32(_a214), v14+int32(144))
	mBase = m.M
	v81 = m.ExcPending
	if v81 != 0 {
		goto L12
	} else {
		goto L25
	}
L25:
	;
	v83 = v80
	goto L17
L26:
	;
	v175 = F_sdscatlen(m, v168, int32(_a10), int32(1))
	mBase = m.M
	v176 = m.ExcPending
	if v176 != 0 {
		goto L12
	} else {
		goto L48
	}
L27:
	;
	v84 = *(*int32)(unsafe.Add(mBase, uint32(l1)+2312))
	v87 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v84+int32(-1)))))
	switch v87 & int32(7) {
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
		goto L29
	}
L28:
	;
	v122 = int32(8)
	v124 = v114
	goto L38
L29:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+128)) = int32(1)
	v111 = F_sdscatfmt(m, v83, int32(_a196), v14+int32(128))
	mBase = m.M
	v112 = m.ExcPending
	if v112 != 0 {
		goto L12
	} else {
		goto L37
	}
L30:
	;
	if v104 != 0 {
		v114 = v83
		goto L28
	} else {
		goto L36
	}
L31:
	;
	v103 = *(*int32)(unsafe.Add(mBase, uint32(v84+int32(-17))))
	v104 = v103
	goto L30
L32:
	;
	v100 = *(*int32)(unsafe.Add(mBase, uint32(v84+int32(-9))))
	v104 = v100
	goto L30
L33:
	;
	v97 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v84+int32(-5)))))
	v104 = v97
	goto L30
L34:
	;
	v94 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v84+int32(-3)))))
	v104 = v94
	goto L30
L35:
	;
	v104 = int32(base.Ui32(v87) >> (uint(int32(3)) % 32))
	goto L30
L36:
	;
	goto L29
L37:
	;
	v114 = v111
	goto L28
L38:
	;
	if base.B2i32(l2 != int32(0))&base.B2i32(v122 == int32(3)) != 0 {
		v159 = v124
		goto L40
	} else {
		goto L41
	}
L39:
	;
	v168 = v159
	goto L26
L40:
	;
	if v122 != 0 {
		v122 = v122 + int32(-1)
		v124 = v159
		goto L38
	} else {
		goto L47
	}
L41:
	;
	if base.B2i32(l2 == int32(0))&base.B2i32(v122 == int32(2)) != 0 {
		v159 = v124
		goto L40
	} else {
		goto L42
	}
L42:
	;
	v138 = v122 << (uint(int32(4)) % 32)
	v141 = *(*int32)(unsafe.Add(mBase, uint32(v138)+uint32(_consts[95])))
	v142 = m.T0[v141].(func(*base.Module, int32) int32)(m, l1)
	mBase = m.M
	v143 = m.ExcPending
	if v143 != 0 {
		goto L12
	} else {
		goto L43
	}
L43:
	;
	if v142 == int32(0) {
		v159 = v124
		goto L40
	} else {
		goto L44
	}
L44:
	;
	v148 = *(*int32)(unsafe.Add(mBase, uint32(v138)+uint32(_consts[92])))
	*(*int32)(unsafe.Add(mBase, uint32(v14)+112)) = v148
	v153 = F_sdscatfmt(m, v124, int32(_a215), v14+int32(112))
	mBase = m.M
	v154 = m.ExcPending
	if v154 != 0 {
		goto L12
	} else {
		goto L45
	}
L45:
	;
	v155 = *(*int32)(unsafe.Add(mBase, uint32(v138)+uint32(_consts[100])))
	v156 = m.T0[v155].(func(*base.Module, int32, int32) int32)(m, l1, v153)
	mBase = m.M
	v157 = m.ExcPending
	if v157 != 0 {
		goto L12
	} else {
		goto L46
	}
L46:
	;
	v159 = v156
	goto L40
L47:
	;
	goto L39
L48:
	;
	v177 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+88)))
	v178 = F_representClusterNodeFlags(m, v175, v177)
	mBase = m.M
	v179 = m.ExcPending
	if v179 != 0 {
		goto L12
	} else {
		goto L49
	}
L49:
	;
	v182 = F_sdscatlen(m, v178, int32(_a10), int32(1))
	mBase = m.M
	v183 = m.ExcPending
	if v183 != 0 {
		goto L12
	} else {
		goto L50
	}
L50:
	;
	v184 = *(*int32)(unsafe.Add(mBase, uint32(l1)+2172))
	if v184 == int32(0) {
		goto L52
	} else {
		goto L53
	}
L51:
	;
	v197 = *(*int32)(unsafe.Add(mBase, uint32(l1)+2172))
	if v197 != 0 {
		goto L56
	} else {
		goto L57
	}
L52:
	;
	v194 = F_sdscatlen(m, v182, int32(_a216), int32(1))
	mBase = m.M
	v195 = m.ExcPending
	if v195 != 0 {
		goto L12
	} else {
		goto L55
	}
L53:
	;
	v190 = F_sdscatlen(m, v182, v184+int32(8), int32(40))
	mBase = m.M
	v191 = m.ExcPending
	if v191 != 0 {
		goto L12
	} else {
		goto L54
	}
L54:
	;
	v196 = v190
	goto L51
L55:
	;
	v196 = v194
	goto L51
L56:
	;
	v198 = v197
	goto L58
L57:
	;
	v198 = l1
	goto L58
L58:
	;
	v199 = *(*int64)(unsafe.Add(mBase, uint32(v198)+96))
	v200 = *(*int64)(unsafe.Add(mBase, uint32(l1)+2192))
	v201 = *(*int64)(unsafe.Add(mBase, uint32(l1)+2184))
	v203 = *(*int32)(unsafe.Add(mBase, uint32(l1)+2344))
	if v203 != 0 {
		v210 = int32(_a217)
		goto L59
	} else {
		goto L60
	}
L59:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14+int32(104)))) = v210
	*(*int64)(unsafe.Add(mBase, uint32(v14+int32(96)))) = v199
	*(*int64)(unsafe.Add(mBase, uint32(v14)+88)) = v200
	*(*int64)(unsafe.Add(mBase, uint32(v14)+80)) = v201
	v222 = F_sdscatfmt(m, v196, int32(_a218), v14+int32(80))
	mBase = m.M
	v223 = m.ExcPending
	if v223 != 0 {
		goto L12
	} else {
		goto L64
	}
L60:
	;
	v206 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+88)))
	if v206&int32(16) != 0 {
		goto L61
	} else {
		goto L62
	}
L61:
	;
	v209 = int32(_a217)
	goto L63
L62:
	;
	v209 = int32(_a219)
	goto L63
L63:
	;
	v210 = v209
	goto L59
L64:
	;
	v224 = *(*int32)(unsafe.Add(mBase, uint32(l1)+2152))
	if v224 == int32(0) {
		goto L66
	} else {
		goto L67
	}
L65:
	;
	v346 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+88)))
	if v346&int32(16) == int32(0) {
		v414 = v339
		goto L104
	} else {
		goto L105
	}
L66:
	;
	v267 = *(*int32)(unsafe.Add(mBase, uint32(l1)+2160))
	if v267 < int32(1) {
		v339 = v222
		goto L65
	} else {
		goto L77
	}
L67:
	;
	v227 = *(*int32)(unsafe.Add(mBase, uint32(l1)+2156))
	if v227 < int32(1) {
		v339 = v222
		goto L65
	} else {
		goto L68
	}
L68:
	;
	v231 = int32(0)
	v235 = v222
	goto L69
L69:
	;
	v244 = v224 + v231<<(uint(int32(1))%32)
	v245 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v244))))
	v248 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v244+int32(2)))))
	if v245 != v248 {
		goto L72
	} else {
		goto L73
	}
L71:
	;
	v265 = v231 + int32(2)
	if v265 < v227 {
		v231 = v265
		v235 = v263
		goto L69
	} else {
		goto L76
	}
L72:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+68)) = v248
	*(*int32)(unsafe.Add(mBase, uint32(v14)+64)) = v245
	v261 = F_sdscatfmt(m, v235, int32(_a220), v14+int32(64))
	mBase = m.M
	v262 = m.ExcPending
	if v262 != 0 {
		goto L12
	} else {
		goto L75
	}
L73:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+48)) = v245
	v254 = F_sdscatfmt(m, v235, int32(_a221), v14+int32(48))
	mBase = m.M
	v255 = m.ExcPending
	if v255 != 0 {
		goto L12
	} else {
		goto L74
	}
L74:
	;
	v263 = v254
	goto L71
L75:
	;
	v263 = v261
	goto L71
L76:
	;
	v339 = v263
	goto L65
L77:
	;
	v274 = int32(0)
	v278 = v222
	v279 = int32(-1)
	goto L78
L78:
	;
	v286 = base.I32_div_s(v274, int32(8))
	v288 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1+int32(104)+v286))))
	v293 = int32(base.Ui32(v288)>>(uint(v274&int32(7))%32)) & int32(1)
	if v293 != 0 {
		goto L82
	} else {
		goto L83
	}
L79:
	;
	v339 = v328
	goto L65
L80:
	;
	if v327 < int32(16383) {
		v274 = v327 + int32(1)
		v278 = v328
		v279 = v329
		goto L78
	} else {
		goto L103
	}
L81:
	;
	v302 = base.B2i32(v274 == int32(16383))
	if v274 == int32(16383) {
		goto L89
	} else {
		goto L90
	}
L82:
	;
	v294 = v274
	goto L84
L83:
	;
	v294 = v279
	goto L84
L84:
	;
	if v279 == int32(-1) {
		goto L85
	} else {
		goto L86
	}
L85:
	;
	v297 = v294
	goto L87
L86:
	;
	v297 = v279
	goto L87
L87:
	;
	if v297 != int32(-1) {
		goto L81
	} else {
		goto L88
	}
L88:
	;
	v327 = v274
	v328 = v278
	v329 = int32(-1)
	goto L80
L89:
	;
	if v293 != 0 {
		goto L94
	} else {
		goto L95
	}
L90:
	;
	if v293 == int32(0) {
		goto L89
	} else {
		goto L91
	}
L91:
	;
	v327 = v274
	v328 = v278
	v329 = v297
	goto L80
L92:
	;
	v327 = v307
	v328 = v324
	v329 = int32(-1)
	goto L80
L93:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+36)) = v309
	*(*int32)(unsafe.Add(mBase, uint32(v14)+32)) = v297
	v322 = F_sdscatfmt(m, v278, int32(_a222), v14+int32(32))
	mBase = m.M
	v323 = m.ExcPending
	if v323 != 0 {
		goto L12
	} else {
		goto L102
	}
L94:
	;
	v306 = int32(16384)
	goto L96
L95:
	;
	v306 = v274
	goto L96
L96:
	;
	if v274 == int32(16383) {
		goto L97
	} else {
		goto L98
	}
L97:
	;
	v307 = v306
	goto L99
L98:
	;
	v307 = v274
	goto L99
L99:
	;
	v309 = v307 + int32(-1)
	if v297 != v309 {
		goto L93
	} else {
		goto L100
	}
L100:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+16)) = v297
	v315 = F_sdscatfmt(m, v278, int32(_a223), v14+int32(16))
	mBase = m.M
	v316 = m.ExcPending
	if v316 != 0 {
		goto L12
	} else {
		goto L101
	}
L101:
	;
	v324 = v315
	goto L92
L102:
	;
	v324 = v322
	goto L92
L103:
	;
	goto L79
L104:
	;
	m.G0 = v14 + int32(176)
	return v414
L105:
	;
	v352 = int32(0)
	v356 = v339
	goto L106
L106:
	;
	v364 = *(*int32)(unsafe.Add(mBase, _consts[90]))
	v365 = *(*int32)(unsafe.Add(mBase, uint32(v364)+44))
	v366 = F_dictFind(m, v365, v352)
	mBase = m.M
	v367 = m.ExcPending
	if v367 != 0 {
		goto L12
	} else {
		goto L110
	}
L107:
	;
	v414 = v404
	goto L104
L108:
	;
	v372 = *(*int32)(unsafe.Add(mBase, _consts[90]))
	v373 = *(*int32)(unsafe.Add(mBase, uint32(v372)+48))
	v374 = F_dictFind(m, v373, v352)
	mBase = m.M
	v375 = m.ExcPending
	if v375 != 0 {
		goto L12
	} else {
		goto L117
	}
L109:
	;
	v369 = *(*int32)(unsafe.Add(mBase, uint32(v366)+8))
	goto L112
L110:
	;
	if v366 != 0 {
		goto L109
	} else {
		goto L111
	}
L111:
	;
	v370 = int32(0)
	goto L108
L112:
	;
	v370 = v369
	goto L108
L113:
	;
	v407 = v352 + int32(1)
	if v407 != int32(16384) {
		v352 = v407
		v356 = v404
		goto L106
	} else {
		goto L126
	}
L114:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14))) = v352
	v392 = F_sdscatfmt(m, v356, v388, v14)
	mBase = m.M
	v393 = m.ExcPending
	if v393 != 0 {
		goto L12
	} else {
		goto L123
	}
L115:
	;
	if v378 == int32(0) {
		v404 = v356
		goto L113
	} else {
		goto L122
	}
L116:
	;
	if v370 == int32(0) {
		v404 = v356
		goto L113
	} else {
		goto L121
	}
L117:
	;
	if v374 == int32(0) {
		goto L116
	} else {
		goto L118
	}
L118:
	;
	v378 = *(*int32)(unsafe.Add(mBase, uint32(v374)+8))
	goto L119
L119:
	;
	if v370 == int32(0) {
		goto L115
	} else {
		goto L120
	}
L120:
	;
	v388 = int32(_a224)
	v390 = v370
	goto L114
L121:
	;
	v388 = int32(_a224)
	v390 = v370
	goto L114
L122:
	;
	v388 = int32(_a225)
	v390 = v378
	goto L114
L123:
	;
	v397 = F_sdscatlen(m, v392, v390+int32(8), int32(40))
	mBase = m.M
	v398 = m.ExcPending
	if v398 != 0 {
		goto L12
	} else {
		goto L124
	}
L124:
	;
	v400 = F_sdscat(m, v397, int32(_a226))
	mBase = m.M
	v401 = m.ExcPending
	if v401 != 0 {
		goto L12
	} else {
		goto L125
	}
L125:
	;
	v404 = v400
	goto L113
L126:
	;
	goto L107
}
func F_clusterGenNodesSlotsInfo(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
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
	var v48 int32
	_ = v48
	var v59 int32
	_ = v59
	var v69 int32
	_ = v69
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v86 int32
	_ = v86
	v2 = int32(0)
	v13 = v2
	v14 = int32(-1)
	v15 = v2
	goto L3
L1:
	;
	return
L2:
	;
	F__serverAssert(m, int32(_a212), int32(_a179), int32(7036))
	mBase = m.M
	v86 = m.ExcPending
	if v86 != 0 {
		goto L17
	} else {
		goto L21
	}
L3:
	;
	if v13 != 0 {
		goto L7
	} else {
		goto L8
	}
L5:
	;
	v13 = v74
	v14 = v75
	v15 = v15 + int32(1)
	goto L3
L6:
	;
	v69 = *(*int32)(unsafe.Add(mBase, _consts[90]))
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v69+v15<<(uint(int32(2))%32))+52))
	v74 = v73
	v75 = v15
	goto L5
L7:
	;
	v23 = base.B2i32(v15 == int32(16384))
	if v15 == int32(16384) {
		goto L10
	} else {
		goto L11
	}
L8:
	;
	if v15 != int32(16384) {
		goto L6
	} else {
		goto L9
	}
L9:
	;
	goto L1
L10:
	;
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v13)+88))
	if v31&l0 != 0 {
		goto L13
	} else {
		goto L14
	}
L11:
	;
	v25 = *(*int32)(unsafe.Add(mBase, _consts[90]))
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v25+v15<<(uint(int32(2))%32))+52))
	if v13 == v29 {
		v74 = v13
		v75 = v14
		goto L5
	} else {
		goto L12
	}
L12:
	;
	goto L10
L13:
	;
	if v15 == int32(16384) {
		goto L1
	} else {
		goto L20
	}
L14:
	;
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v13)+2152))
	if v33 != 0 {
		v40 = v33
		goto L15
	} else {
		goto L16
	}
L15:
	;
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v13)+2156))
	v42 = int32(1)
	v43 = v41 + v42
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v13)+2160))
	if v44<<(uint(v42)%32) <= v43 {
		goto L2
	} else {
		goto L19
	}
L16:
	;
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v13)+2160))
	v37 = F_valkey_malloc(m, v34<<(uint(int32(2))%32))
	mBase = m.M
	v38 = m.ExcPending
	if v38 != 0 {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	return
L18:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+2152)) = v37
	v40 = v37
	goto L15
L19:
	;
	v48 = int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v40+v41<<(uint(v48)%32)))) = uint16(v14)
	*(*int32)(unsafe.Add(mBase, uint32(v13)+2156)) = v41 + int32(2)
	v59 = v15 + int32(-1)
	*(*uint16)(unsafe.Add(mBase, uint32(v40+v43<<(uint(v48)%32)))) = uint16(v59)
	goto L13
L20:
	;
	goto L6
L21:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_clusterGetFailedPrimaryRank(m *base.Module) int32 {
	mBase := m.M
	_ = mBase
	var v1 int32
	_ = v1
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v13 int32
	_ = v13
	var v16 int64
	_ = v16
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v57 int64
	_ = v57
	var v58 int64
	_ = v58
	var v59 int64
	_ = v59
	var v60 int64
	_ = v60
	var v61 int64
	_ = v61
	var v62 int64
	_ = v62
	var v63 int64
	_ = v63
	var v65 int64
	_ = v65
	var v67 int64
	_ = v67
	var v69 int64
	_ = v69
	var v71 int64
	_ = v71
	var v73 int64
	_ = v73
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
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
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v115 int32
	_ = v115
	var v120 int32
	_ = v120
	var v124 int32
	_ = v124
	var v128 int32
	_ = v128
	var v134 int32
	_ = v134
	var v136 int32
	_ = v136
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v145 int32
	_ = v145
	var v148 int32
	_ = v148
	var v152 int32
	_ = v152
	var v155 int64
	_ = v155
	var v158 int64
	_ = v158
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v165 int32
	_ = v165
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v187 int32
	_ = v187
	var v189 int32
	_ = v189
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v212 int32
	_ = v212
	var v217 int32
	_ = v217
	var v232 int32
	_ = v232
	var v236 int32
	_ = v236
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v249 int32
	_ = v249
	var v252 int32
	_ = v252
	var v253 int32
	_ = v253
	var v257 int32
	_ = v257
	var v261 int32
	_ = v261
	var v262 int32
	_ = v262
	var v263 int32
	_ = v263
	var v266 int32
	_ = v266
	var v268 int32
	_ = v268
	var v270 int64
	_ = v270
	var v271 int64
	_ = v271
	var v272 int64
	_ = v272
	var v273 int64
	_ = v273
	var v274 int64
	_ = v274
	var v275 int64
	_ = v275
	var v276 int64
	_ = v276
	var v278 int64
	_ = v278
	var v280 int64
	_ = v280
	var v282 int64
	_ = v282
	var v284 int64
	_ = v284
	var v286 int64
	_ = v286
	var v288 int32
	_ = v288
	var v289 int32
	_ = v289
	var v290 int32
	_ = v290
	var v293 int32
	_ = v293
	var v296 int32
	_ = v296
	var v297 int32
	_ = v297
	var v298 int32
	_ = v298
	var v300 int32
	_ = v300
	var v303 int32
	_ = v303
	var v304 int32
	_ = v304
	var v308 int32
	_ = v308
	var v312 int32
	_ = v312
	var v314 int32
	_ = v314
	var v321 int32
	_ = v321
	var v322 int32
	_ = v322
	var v323 int32
	_ = v323
	var v328 int32
	_ = v328
	var v333 int32
	_ = v333
	var v337 int32
	_ = v337
	var v341 int32
	_ = v341
	var v347 int32
	_ = v347
	var v351 int32
	_ = v351
	var v357 int32
	_ = v357
	var v363 int32
	_ = v363
	v1 = int32(0)
	v7 = *(*int32)(unsafe.Add(mBase, _consts[94]))
	v8 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7)+88)))
	if v8&int32(2) == v1 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	F__serverAssert(m, int32(_a291), int32(_a179), int32(5467))
	mBase = m.M
	v363 = m.ExcPending
	if v363 != 0 {
		goto L6
	} else {
		goto L90
	}
L2:
	;
	F__serverAssert(m, int32(_a292), int32(_a179), int32(5466))
	mBase = m.M
	v357 = m.ExcPending
	if v357 != 0 {
		goto L6
	} else {
		goto L89
	}
L3:
	;
	v13 = *(*int32)(unsafe.Add(mBase, uint32(v7)+2172))
	if v13 == int32(0) {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v16 = F_mstime(m)
	mBase = m.M
	v17 = int32(0)
	v19 = *(*int32)(unsafe.Add(mBase, _consts[90]))
	v20 = *(*int32)(unsafe.Add(mBase, uint32(v19)+32))
	v21 = F_dictGetSafeIterator(m, v20)
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	F_dictReleaseIterator(m, v21)
	mBase = m.M
	v351 = m.ExcPending
	if v351 != 0 {
		goto L6
	} else {
		goto L88
	}
L6:
	;
	return int32(0)
L7:
	;
	v32 = v21 + int32(20)
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v21)+16))
	if v33 != 0 {
		goto L10
	} else {
		goto L11
	}
L8:
	;
	if v128 == int32(0) {
		v347 = v17
		goto L5
	} else {
		goto L34
	}
L9:
	;
	v39 = v32
	v40 = v36
	goto L12
L10:
	;
	v36 = int32(1)
	goto L9
L11:
	;
	v36 = int32(0)
	goto L9
L12:
	;
	switch v40 {
	case 0:
		goto L17
	default:
		goto L16
	}
L14:
	;
	v40 = int32(0)
	goto L12
L15:
	;
	goto L8
L16:
	;
	v120 = *(*int32)(unsafe.Add(mBase, uint32(v39)))
	*(*int32)(unsafe.Add(mBase, uint32(v21)+16)) = v120
	if v120 == int32(0) {
		goto L14
	} else {
		goto L33
	}
L17:
	;
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v21)+4))
	if v44 != int32(-1) {
		v83 = v44
		goto L18
	} else {
		goto L19
	}
L18:
	;
	v84 = int32(1)
	v85 = v83 + v84
	*(*int32)(unsafe.Add(mBase, uint32(v21)+4)) = v85
	v87 = int32(0)
	v90 = *(*int32)(unsafe.Add(mBase, uint32(v21)))
	v91 = *(*int32)(unsafe.Add(mBase, uint32(v21)+8))
	v95 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v90+v91+int32(26)))))
	if v95 == int32(255) {
		goto L27
	} else {
		goto L28
	}
L19:
	;
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v21)+8))
	if v48 != 0 {
		v83 = int32(-1)
		goto L18
	} else {
		goto L20
	}
L20:
	;
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v21)))
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v21)+12))
	if v50 == int32(0) {
		goto L22
	} else {
		goto L23
	}
L21:
	;
	v77 = *(*int32)(unsafe.Add(mBase, uint32(v76)+20))
	if v77 != int32(-1) {
		goto L24
	} else {
		goto L25
	}
L22:
	;
	v57 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v49)+16)))
	v58 = int64(*(*int8)(unsafe.Add(mBase, uint32(v49)+27)))
	v59 = int64(*(*int32)(unsafe.Add(mBase, uint32(v49)+8)))
	v60 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v49)+12)))
	v61 = int64(*(*int8)(unsafe.Add(mBase, uint32(v49)+26)))
	v62 = int64(*(*int32)(unsafe.Add(mBase, uint32(v49)+4)))
	v63 = F_wangHash64(m, v62)
	mBase = m.M
	v65 = F_wangHash64(m, v61+v63)
	mBase = m.M
	v67 = F_wangHash64(m, v60+v65)
	mBase = m.M
	v69 = F_wangHash64(m, v59+v67)
	mBase = m.M
	v71 = F_wangHash64(m, v58+v69)
	mBase = m.M
	v73 = F_wangHash64(m, v57+v71)
	mBase = m.M
	*(*int64)(unsafe.Add(mBase, uint32(v21)+24)) = v73
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v21)))
	v76 = v75
	goto L21
L23:
	;
	v53 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v49)+24)))
	v55 = v53 + int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v49)+24)) = uint16(v55)
	v76 = v49
	goto L21
L24:
	;
	v83 = v77 + int32(-1)
	goto L18
L25:
	;
	v80 = *(*int32)(unsafe.Add(mBase, uint32(v21)+4))
	v83 = v80
	goto L18
L26:
	;
	v110 = int32(2)
	v115 = *(*int32)(unsafe.Add(mBase, uint32(v90+v108<<(uint(v110)%32)+int32(4))))
	v39 = v115 + v109<<(uint(v110)%32)
	v40 = int32(1)
	goto L12
L27:
	;
	v99 = v87
	goto L29
L28:
	;
	v99 = v84 << (uint(v95) % 32)
	goto L29
L29:
	;
	if v85 < v99 {
		v108 = v91
		v109 = v85
		goto L26
	} else {
		goto L30
	}
L30:
	;
	if v91 != 0 {
		v128 = v87
		goto L15
	} else {
		goto L31
	}
L31:
	;
	v101 = *(*int32)(unsafe.Add(mBase, uint32(v90)+20))
	if v101 == int32(-1) {
		v128 = v87
		goto L15
	} else {
		goto L32
	}
L32:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v21)+4)) = int64(4294967296)
	v108 = int32(1)
	v109 = int32(0)
	goto L26
L33:
	;
	v124 = *(*int32)(unsafe.Add(mBase, uint32(v120)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v32))) = v124
	v128 = v120
	goto L15
L34:
	;
	v134 = v128
	v136 = v17
	goto L35
L35:
	;
	v139 = *(*int32)(unsafe.Add(mBase, uint32(v134)+8))
	goto L38
L36:
	;
	v347 = v236
	goto L5
L37:
	;
	v245 = v21 + int32(20)
	v246 = *(*int32)(unsafe.Add(mBase, uint32(v21)+16))
	if v246 != 0 {
		goto L63
	} else {
		goto L64
	}
L38:
	;
	v140 = *(*int32)(unsafe.Add(mBase, uint32(v139)+88))
	v141 = int32(9)
	if v140&v141 != v141 {
		v236 = v136
		goto L37
	} else {
		goto L39
	}
L39:
	;
	v145 = *(*int32)(unsafe.Add(mBase, uint32(v139)+2160))
	if v145 == int32(0) {
		v236 = v136
		goto L37
	} else {
		goto L40
	}
L40:
	;
	v148 = *(*int32)(unsafe.Add(mBase, uint32(v139)+2164))
	if v148 == int32(0) {
		v236 = v136
		goto L37
	} else {
		goto L41
	}
L41:
	;
	v152 = *(*int32)(unsafe.Add(mBase, _consts[137]))
	if v152 == int32(0) {
		goto L42
	} else {
		goto L43
	}
L42:
	;
	v162 = int32(48)
	v163 = v139 + v162
	v165 = *(*int32)(unsafe.Add(mBase, _consts[94]))
	v167 = v165 + v162
	v168 = int32(40)
	goto L49
L43:
	;
	v155 = *(*int64)(unsafe.Add(mBase, uint32(v139)+2216))
	v158 = *(*int64)(unsafe.Add(mBase, _consts[117]))
	if v158*base.I64_extend_i32_s(v152) < v16-v155 {
		v236 = v136
		goto L37
	} else {
		goto L44
	}
L44:
	;
	goto L42
L45:
	;
	v236 = int32(base.Ui32(v232)>>(uint(int32(31))%32)) + v136
	goto L37
L46:
	;
	v232 = int32(0)
	goto L45
L47:
	;
	v204 = v199
	v205 = v200
	v206 = v201
	goto L57
L48:
	;
	if v189 == int32(0) {
		goto L46
	} else {
		goto L55
	}
L49:
	;
	if (v167|v163)&int32(3) != 0 {
		v199 = v163
		v200 = v167
		v201 = v168
		goto L47
	} else {
		goto L50
	}
L50:
	;
	v176 = v163
	v177 = v167
	v178 = v168
	goto L51
L51:
	;
	v181 = *(*int32)(unsafe.Add(mBase, uint32(v176)))
	v182 = *(*int32)(unsafe.Add(mBase, uint32(v177)))
	if v181 != v182 {
		v199 = v176
		v200 = v177
		v201 = v178
		goto L47
	} else {
		goto L53
	}
L52:
	;
	goto L48
L53:
	;
	v184 = int32(4)
	v185 = v177 + v184
	v187 = v176 + v184
	v189 = v178 + int32(-4)
	if base.Ui32(int32(3)) < base.Ui32(v189) {
		v176 = v187
		v177 = v185
		v178 = v189
		goto L51
	} else {
		goto L54
	}
L54:
	;
	goto L52
L55:
	;
	v199 = v187
	v200 = v185
	v201 = v189
	goto L47
L56:
	;
	v232 = v209 - v210
	goto L45
L57:
	;
	v209 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v204))))
	v210 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v205))))
	if v209 != v210 {
		goto L56
	} else {
		goto L59
	}
L59:
	;
	v212 = int32(1)
	v217 = v206 + int32(-1)
	if v217 == int32(0) {
		goto L46
	} else {
		goto L60
	}
L60:
	;
	v204 = v204 + v212
	v205 = v205 + v212
	v206 = v217
	goto L57
L61:
	;
	if v341 != 0 {
		v134 = v341
		v136 = v236
		goto L35
	} else {
		goto L87
	}
L62:
	;
	v252 = v245
	v253 = v249
	goto L65
L63:
	;
	v249 = int32(1)
	goto L62
L64:
	;
	v249 = int32(0)
	goto L62
L65:
	;
	switch v253 {
	case 0:
		goto L70
	default:
		goto L69
	}
L67:
	;
	v253 = int32(0)
	goto L65
L68:
	;
	goto L61
L69:
	;
	v333 = *(*int32)(unsafe.Add(mBase, uint32(v252)))
	*(*int32)(unsafe.Add(mBase, uint32(v21)+16)) = v333
	if v333 == int32(0) {
		goto L67
	} else {
		goto L86
	}
L70:
	;
	v257 = *(*int32)(unsafe.Add(mBase, uint32(v21)+4))
	if v257 != int32(-1) {
		v296 = v257
		goto L71
	} else {
		goto L72
	}
L71:
	;
	v297 = int32(1)
	v298 = v296 + v297
	*(*int32)(unsafe.Add(mBase, uint32(v21)+4)) = v298
	v300 = int32(0)
	v303 = *(*int32)(unsafe.Add(mBase, uint32(v21)))
	v304 = *(*int32)(unsafe.Add(mBase, uint32(v21)+8))
	v308 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v303+v304+int32(26)))))
	if v308 == int32(255) {
		goto L80
	} else {
		goto L81
	}
L72:
	;
	v261 = *(*int32)(unsafe.Add(mBase, uint32(v21)+8))
	if v261 != 0 {
		v296 = int32(-1)
		goto L71
	} else {
		goto L73
	}
L73:
	;
	v262 = *(*int32)(unsafe.Add(mBase, uint32(v21)))
	v263 = *(*int32)(unsafe.Add(mBase, uint32(v21)+12))
	if v263 == int32(0) {
		goto L75
	} else {
		goto L76
	}
L74:
	;
	v290 = *(*int32)(unsafe.Add(mBase, uint32(v289)+20))
	if v290 != int32(-1) {
		goto L77
	} else {
		goto L78
	}
L75:
	;
	v270 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v262)+16)))
	v271 = int64(*(*int8)(unsafe.Add(mBase, uint32(v262)+27)))
	v272 = int64(*(*int32)(unsafe.Add(mBase, uint32(v262)+8)))
	v273 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v262)+12)))
	v274 = int64(*(*int8)(unsafe.Add(mBase, uint32(v262)+26)))
	v275 = int64(*(*int32)(unsafe.Add(mBase, uint32(v262)+4)))
	v276 = F_wangHash64(m, v275)
	mBase = m.M
	v278 = F_wangHash64(m, v274+v276)
	mBase = m.M
	v280 = F_wangHash64(m, v273+v278)
	mBase = m.M
	v282 = F_wangHash64(m, v272+v280)
	mBase = m.M
	v284 = F_wangHash64(m, v271+v282)
	mBase = m.M
	v286 = F_wangHash64(m, v270+v284)
	mBase = m.M
	*(*int64)(unsafe.Add(mBase, uint32(v21)+24)) = v286
	v288 = *(*int32)(unsafe.Add(mBase, uint32(v21)))
	v289 = v288
	goto L74
L76:
	;
	v266 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v262)+24)))
	v268 = v266 + int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v262)+24)) = uint16(v268)
	v289 = v262
	goto L74
L77:
	;
	v296 = v290 + int32(-1)
	goto L71
L78:
	;
	v293 = *(*int32)(unsafe.Add(mBase, uint32(v21)+4))
	v296 = v293
	goto L71
L79:
	;
	v323 = int32(2)
	v328 = *(*int32)(unsafe.Add(mBase, uint32(v303+v321<<(uint(v323)%32)+int32(4))))
	v252 = v328 + v322<<(uint(v323)%32)
	v253 = int32(1)
	goto L65
L80:
	;
	v312 = v300
	goto L82
L81:
	;
	v312 = v297 << (uint(v308) % 32)
	goto L82
L82:
	;
	if v298 < v312 {
		v321 = v304
		v322 = v298
		goto L79
	} else {
		goto L83
	}
L83:
	;
	if v304 != 0 {
		v341 = v300
		goto L68
	} else {
		goto L84
	}
L84:
	;
	v314 = *(*int32)(unsafe.Add(mBase, uint32(v303)+20))
	if v314 == int32(-1) {
		v341 = v300
		goto L68
	} else {
		goto L85
	}
L85:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v21)+4)) = int64(4294967296)
	v321 = int32(1)
	v322 = int32(0)
	goto L79
L86:
	;
	v337 = *(*int32)(unsafe.Add(mBase, uint32(v333)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v245))) = v337
	v341 = v333
	goto L68
L87:
	;
	goto L36
L88:
	;
	return v347
L89:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L90:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_clusterHandleFlushDuringSlotMigration(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
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
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v29 int32
	_ = v29
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v60 int32
	_ = v60
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	v9 = *(*int32)(unsafe.Add(mBase, _consts[90]))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(v9)+uint32(_consts[143])))
	v12 = v6 + int32(8)
	v13 = *(*int32)(unsafe.Add(mBase, uint32(v10)))
	*(*int32)(unsafe.Add(mBase, uint32(v12)+4)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v12))) = v13
	goto L1
L1:
	;
	v18 = v6 + int32(8)
	v20 = *(*int32)(unsafe.Add(mBase, uint32(v18)))
	if v20 == int32(0) {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	m.G0 = v6 + int32(16)
	return
L3:
	;
	if v20 == int32(0) {
		goto L2
	} else {
		goto L6
	}
L4:
	;
	goto L3
L5:
	;
	v23 = *(*int32)(unsafe.Add(mBase, uint32(v18)+4))
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v20+base.B2i32(v23 == int32(0))<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v18))) = v29
	goto L4
L6:
	;
	v34 = v20
	goto L7
L7:
	;
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v34)+8))
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v36)+156))
	if base.Ui32(int32(20)) < base.Ui32(v37) {
		goto L10
	} else {
		goto L11
	}
L8:
	;
	goto L2
L9:
	;
	v49 = v6 + int32(8)
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v49)))
	if v51 == int32(0) {
		goto L16
	} else {
		goto L17
	}
L10:
	;
	F_finishSlotMigrationJob(m, v36, int32(18), int32(_a345))
	mBase = m.M
	v47 = m.ExcPending
	if v47 != 0 {
		goto L13
	} else {
		goto L14
	}
L11:
	;
	if int32(1)<<(uint(v37)%32)&int32(1835040) != 0 {
		goto L9
	} else {
		goto L12
	}
L12:
	;
	goto L10
L13:
	;
	return
L14:
	;
	goto L9
L15:
	;
	if v51 != 0 {
		v34 = v51
		goto L7
	} else {
		goto L18
	}
L16:
	;
	goto L15
L17:
	;
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v49)+4))
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v51+base.B2i32(v54 == int32(0))<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v49))) = v60
	goto L16
L18:
	;
	goto L8
}
func F_clusterHandleSlotMigrationClientClose(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v10 int32
	_ = v10
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	*(*int32)(unsafe.Add(mBase, uint32(l0)+152)) = int32(0)
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+156))
	if base.Ui32(int32(20)) < base.Ui32(v10) {
		v17 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+196)))
		if v17 != 0 {
			m.G0 = v6 + int32(16)
			return
		} else {
			v19 = *(*int32)(unsafe.Add(mBase, _consts[15]))
			if int32(2) < v19 {
				v29 = v10
				if v29 == int32(17) {
					F_clusterDoBeforeSleep(m, int32(64))
					mBase = m.M
					v41 = m.ExcPending
					if v41 != 0 {
						return
					} else {
						m.G0 = v6 + int32(16)
						return
					}
				} else {
					v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
					if v35 != 0 {
						v36 = int32(_a382)
					} else {
						v36 = int32(_a383)
					}
					F_finishSlotMigrationJob(m, l0, int32(18), v36)
					mBase = m.M
					v38 = m.ExcPending
					if v38 != 0 {
						return
					} else {
						F_clusterDoBeforeSleep(m, int32(64))
						mBase = m.M
						v41 = m.ExcPending
						if v41 != 0 {
							return
						} else {
							m.G0 = v6 + int32(16)
							return
						}
					}
				}
			} else {
				v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+188))
				*(*int32)(unsafe.Add(mBase, uint32(v6))) = v22
				F__serverLog(m, int32(2), int32(_a384), v6)
				mBase = m.M
				v27 = m.ExcPending
				if v27 != 0 {
					return
				} else {
					v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+156))
					v29 = v28
					if v29 == int32(17) {
						F_clusterDoBeforeSleep(m, int32(64))
						mBase = m.M
						v41 = m.ExcPending
						if v41 != 0 {
							return
						} else {
							m.G0 = v6 + int32(16)
							return
						}
					} else {
						v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
						if v35 != 0 {
							v36 = int32(_a382)
						} else {
							v36 = int32(_a383)
						}
						F_finishSlotMigrationJob(m, l0, int32(18), v36)
						mBase = m.M
						v38 = m.ExcPending
						if v38 != 0 {
							return
						} else {
							F_clusterDoBeforeSleep(m, int32(64))
							mBase = m.M
							v41 = m.ExcPending
							if v41 != 0 {
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
	} else {
		if int32(1)<<(uint(v10)%32)&int32(1835040) != 0 {
			m.G0 = v6 + int32(16)
			return
		} else {
			v17 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+196)))
			if v17 != 0 {
				m.G0 = v6 + int32(16)
				return
			} else {
				v19 = *(*int32)(unsafe.Add(mBase, _consts[15]))
				if int32(2) < v19 {
					v29 = v10
					if v29 == int32(17) {
						F_clusterDoBeforeSleep(m, int32(64))
						mBase = m.M
						v41 = m.ExcPending
						if v41 != 0 {
							return
						} else {
							m.G0 = v6 + int32(16)
							return
						}
					} else {
						v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
						if v35 != 0 {
							v36 = int32(_a382)
						} else {
							v36 = int32(_a383)
						}
						F_finishSlotMigrationJob(m, l0, int32(18), v36)
						mBase = m.M
						v38 = m.ExcPending
						if v38 != 0 {
							return
						} else {
							F_clusterDoBeforeSleep(m, int32(64))
							mBase = m.M
							v41 = m.ExcPending
							if v41 != 0 {
								return
							} else {
								m.G0 = v6 + int32(16)
								return
							}
						}
					}
				} else {
					v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+188))
					*(*int32)(unsafe.Add(mBase, uint32(v6))) = v22
					F__serverLog(m, int32(2), int32(_a384), v6)
					mBase = m.M
					v27 = m.ExcPending
					if v27 != 0 {
						return
					} else {
						v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+156))
						v29 = v28
						if v29 == int32(17) {
							F_clusterDoBeforeSleep(m, int32(64))
							mBase = m.M
							v41 = m.ExcPending
							if v41 != 0 {
								return
							} else {
								m.G0 = v6 + int32(16)
								return
							}
						} else {
							v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
							if v35 != 0 {
								v36 = int32(_a382)
							} else {
								v36 = int32(_a383)
							}
							F_finishSlotMigrationJob(m, l0, int32(18), v36)
							mBase = m.M
							v38 = m.ExcPending
							if v38 != 0 {
								return
							} else {
								F_clusterDoBeforeSleep(m, int32(64))
								mBase = m.M
								v41 = m.ExcPending
								if v41 != 0 {
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
	}
}
func F_clusterIsAnySlotImporting(m *base.Module) int32 {
	mBase := m.M
	_ = mBase
	var v1 int32
	_ = v1
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v10 int32
	_ = v10
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v37 int32
	_ = v37
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	v1 = int32(0)
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	v10 = *(*int32)(unsafe.Add(mBase, _consts[63]))
	if v10 == v1 {
		v69 = v1
		goto L1
	} else {
		goto L2
	}
L1:
	;
	m.G0 = v6 + int32(16)
	return v69
L2:
	;
	v14 = *(*int32)(unsafe.Add(mBase, _consts[90]))
	if v14 == int32(0) {
		v69 = v1
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v17 = *(*int32)(unsafe.Add(mBase, uint32(v14)+uint32(_consts[143])))
	v19 = v6 + int32(8)
	v20 = *(*int32)(unsafe.Add(mBase, uint32(v17)))
	*(*int32)(unsafe.Add(mBase, uint32(v19)+4)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v19))) = v20
	goto L4
L4:
	;
	v24 = int32(0)
	v26 = v6 + int32(8)
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v26)))
	if v28 == v24 {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	if v28 == int32(0) {
		v69 = v24
		goto L1
	} else {
		goto L8
	}
L6:
	;
	goto L5
L7:
	;
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v26)+4))
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v28+base.B2i32(v31 == int32(0))<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v26))) = v37
	goto L6
L8:
	;
	v43 = v28
	goto L9
L9:
	;
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v43)+8))
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v44)))
	if v45 != int32(1) {
		goto L11
	} else {
		goto L12
	}
L10:
	;
	v69 = v24
	goto L1
L11:
	;
	v55 = v6 + int32(8)
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v55)))
	if v57 == int32(0) {
		goto L15
	} else {
		goto L16
	}
L12:
	;
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v44)+156))
	if base.Ui32(v48+int32(-18)) <= base.Ui32(int32(2)) {
		goto L11
	} else {
		goto L13
	}
L13:
	;
	v69 = int32(1)
	goto L1
L14:
	;
	if v57 != 0 {
		v43 = v57
		goto L9
	} else {
		goto L17
	}
L15:
	;
	goto L14
L16:
	;
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v55)+4))
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v57+base.B2i32(v60 == int32(0))<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v55))) = v66
	goto L15
L17:
	;
	goto L10
}
func F_clusterIsSlotExporting(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v27 int32
	_ = v27
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v61 int32
	_ = v61
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v86 int32
	_ = v86
	var v93 int32
	_ = v93
	var v96 int32
	_ = v96
	var v102 int32
	_ = v102
	var v107 int32
	_ = v107
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v10 = *(*int32)(unsafe.Add(mBase, _consts[90]))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(v10)+uint32(_consts[143])))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
	*(*int32)(unsafe.Add(mBase, uint32(v7)+4)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v7))) = v12
	goto L1
L1:
	;
	v16 = int32(0)
	v18 = *(*int32)(unsafe.Add(mBase, uint32(v7)))
	if v18 == v16 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	m.G0 = v7 + int32(16)
	return v107
L3:
	;
	if v18 == int32(0) {
		v107 = v16
		goto L2
	} else {
		goto L6
	}
L4:
	;
	goto L3
L5:
	;
	v21 = *(*int32)(unsafe.Add(mBase, uint32(v7)+4))
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v18+base.B2i32(v21 == int32(0))<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v7))) = v27
	goto L4
L6:
	;
	v34 = v18
	goto L8
L7:
	;
	v107 = int32(1)
	goto L2
L8:
	;
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v34)+8))
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v35)))
	if v36 != 0 {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	v93 = *(*int32)(unsafe.Add(mBase, uint32(v7)))
	if v93 == int32(0) {
		goto L28
	} else {
		goto L29
	}
L11:
	;
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v35)+156))
	if base.Ui32(v37+int32(-18)) < base.Ui32(int32(3)) {
		goto L10
	} else {
		goto L12
	}
L12:
	;
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v35)+164))
	v44 = v7 + int32(8)
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v42)))
	*(*int32)(unsafe.Add(mBase, uint32(v44)+4)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v44))) = v45
	goto L13
L13:
	;
	v50 = v7 + int32(8)
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v50)))
	if v52 == int32(0) {
		goto L15
	} else {
		goto L16
	}
L14:
	;
	if v52 == int32(0) {
		goto L10
	} else {
		goto L17
	}
L15:
	;
	goto L14
L16:
	;
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v50)+4))
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v52+base.B2i32(v55 == int32(0))<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v50))) = v61
	goto L15
L17:
	;
	v68 = v52
	goto L18
L18:
	;
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v68)+8))
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v69)))
	if l0 < v70 {
		goto L20
	} else {
		goto L21
	}
L19:
	;
	goto L10
L20:
	;
	v75 = v7 + int32(8)
	v77 = *(*int32)(unsafe.Add(mBase, uint32(v75)))
	if v77 == int32(0) {
		goto L24
	} else {
		goto L25
	}
L21:
	;
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v69)+4))
	if l0 <= v72 {
		goto L7
	} else {
		goto L22
	}
L22:
	;
	goto L20
L23:
	;
	if v77 != 0 {
		v68 = v77
		goto L18
	} else {
		goto L26
	}
L24:
	;
	goto L23
L25:
	;
	v80 = *(*int32)(unsafe.Add(mBase, uint32(v75)+4))
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v77+base.B2i32(v80 == int32(0))<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v75))) = v86
	goto L24
L26:
	;
	goto L19
L27:
	;
	if v93 != 0 {
		v34 = v93
		goto L8
	} else {
		goto L30
	}
L28:
	;
	goto L27
L29:
	;
	v96 = *(*int32)(unsafe.Add(mBase, uint32(v7)+4))
	v102 = *(*int32)(unsafe.Add(mBase, uint32(v93+base.B2i32(v96 == int32(0))<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v7))) = v102
	goto L28
L30:
	;
	v107 = v16
	goto L2
}
func F_clusterKeySlotCommand(m *base.Module, l0 int32) {
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
	var v37 int32
	_ = v37
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v72 int32
	_ = v72
	var v81 int32
	_ = v81
	var v88 int32
	_ = v88
	var v91 int32
	_ = v91
	var v98 int32
	_ = v98
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(v6)+8))
	v8 = F_objectGetVal(m, v7)
	mBase = m.M
	v11 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8+int32(-1)))))
	switch v11 & int32(7) {
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
	v29 = int32(0)
	if v28 < int32(1) {
		v49 = v29
		goto L11
	} else {
		goto L12
	}
L2:
	;
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v8+int32(-17))))
	v28 = v27
	goto L1
L3:
	;
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v8+int32(-9))))
	v28 = v24
	goto L1
L4:
	;
	v21 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v8+int32(-5)))))
	v28 = v21
	goto L1
L5:
	;
	v18 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8+int32(-3)))))
	v28 = v18
	goto L1
L6:
	;
	v28 = int32(base.Ui32(v11) >> (uint(int32(3)) % 32))
	goto L1
L7:
	;
	F_addReplyLongLong(m, l0, base.I64_extend_i32_u(v91&int32(16383)))
	mBase = m.M
	v98 = m.ExcPending
	if v98 != 0 {
		goto L28
	} else {
		goto L29
	}
L8:
	;
	goto L7
L9:
	;
	v60 = v49 + int32(1)
	if v28 <= v60 {
		goto L19
	} else {
		goto L20
	}
L10:
	;
	v58 = F_crc16(m, v8, v28)
	mBase = m.M
	v91 = v58
	goto L8
L11:
	;
	if v49 != v28 {
		goto L9
	} else {
		goto L17
	}
L12:
	;
	v37 = v29
	goto L13
L13:
	;
	v41 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8+v37))))
	if v41 == int32(123) {
		v49 = v37
		goto L11
	} else {
		goto L15
	}
L15:
	;
	v45 = v37 + int32(1)
	if v45 != v28 {
		v37 = v45
		goto L13
	} else {
		goto L16
	}
L16:
	;
	goto L10
L17:
	;
	goto L10
L18:
	;
	v88 = F_crc16(m, v8+v49+int32(1), v66+(v49^int32(-1)))
	mBase = m.M
	v91 = v88
	goto L8
L19:
	;
	v81 = F_crc16(m, v8, v28)
	mBase = m.M
	v91 = v81
	goto L8
L20:
	;
	v66 = v60
	goto L22
L21:
	;
	if v66 == v28 {
		goto L19
	} else {
		goto L26
	}
L22:
	;
	v68 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8+v66))))
	if v68 == int32(125) {
		goto L21
	} else {
		goto L24
	}
L24:
	;
	v72 = v66 + int32(1)
	if v72 != v28 {
		v66 = v72
		goto L22
	} else {
		goto L25
	}
L25:
	;
	goto L19
L26:
	;
	if v66 != v60 {
		goto L18
	} else {
		goto L27
	}
L27:
	;
	goto L19
L28:
	;
	return
L29:
	;
	return
}
func F_clusterLookupNode(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	v3 = int32(0)
	if l1 != int32(40) {
		v37 = int32(-1)
		goto L3
	} else {
		goto L4
	}
L1:
	;
	return v54
L2:
	;
	if v37 != 0 {
		v54 = v3
		goto L1
	} else {
		goto L10
	}
L3:
	;
	goto L2
L4:
	;
	v12 = int32(0)
	goto L6
L5:
	;
	v37 = int32(0) - v27
	goto L3
L6:
	;
	v14 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0+v12))))
	v17 = int32(255)
	v27 = base.B2i32(base.Ui32((v14+int32(-123))&v17) < base.Ui32(int32(230))) & base.B2i32(base.Ui32((v14+int32(-58))&v17) < base.Ui32(int32(246)))
	if v27 != 0 {
		goto L5
	} else {
		goto L8
	}
L7:
	;
	goto L5
L8:
	;
	v29 = v12 + int32(1)
	if v29 != int32(40) {
		v12 = v29
		goto L6
	} else {
		goto L9
	}
L9:
	;
	goto L7
L10:
	;
	v38 = F_sdsnewlen(m, l0, l1)
	mBase = m.M
	v41 = m.ExcPending
	if v41 != 0 {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	return int32(0)
L12:
	;
	v43 = *(*int32)(unsafe.Add(mBase, _consts[90]))
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v43)+32))
	v45 = F_dictFind(m, v44, v38)
	mBase = m.M
	v46 = m.ExcPending
	if v46 != 0 {
		goto L11
	} else {
		goto L13
	}
L13:
	;
	F_sdsfree(m, v38)
	mBase = m.M
	v48 = m.ExcPending
	if v48 != 0 {
		goto L11
	} else {
		goto L14
	}
L14:
	;
	if v45 == int32(0) {
		v54 = v3
		goto L1
	} else {
		goto L15
	}
L15:
	;
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v45)+8))
	goto L16
L16:
	;
	v54 = v51
	goto L1
}
func F_clusterMsgSendBlockDecrRefCount(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v17 int32
	_ = v17
	var v22 int32
	_ = v22
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v6 = v4 + int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v6
	if v4 <= int32(0) {
		F__serverAssert(m, int32(_a233), int32(_a179), int32(1756))
		mBase = m.M
		v22 = m.ExcPending
		if v22 != 0 {
			return
		} else {
			F_abort(m)
			mBase = m.M
			base.Wasm_trap_unreachable()
			for {
			}
		}
	} else {
		if v6 != 0 {
			return
		} else {
			v10 = int32(_a69)
			v12 = *(*int32)(unsafe.Add(mBase, _consts[118]))
			v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			*(*int32)(unsafe.Add(mBase, _consts[118])) = v12 - v13
			F_valkey_free(m, l0)
			mBase = m.M
			v17 = m.ExcPending
			if v17 != 0 {
				return
			} else {
				return
			}
		}
	}
}
func F_clusterNodeCleanupFailureReports(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v3 int64
	_ = v3
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int64
	_ = v10
	var v12 int64
	_ = v12
	var v13 int32
	_ = v13
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v48 int32
	_ = v48
	var v49 int64
	_ = v49
	var v50 int64
	_ = v50
	var v52 int64
	_ = v52
	var v54 int64
	_ = v54
	var v57 int64
	_ = v57
	var v59 int64
	_ = v59
	var v61 int64
	_ = v61
	var v63 int64
	_ = v63
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
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
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v106 int32
	_ = v106
	var v114 int32
	_ = v114
	v3 = int64(0)
	v6 = m.G0
	v8 = v6 - int32(304)
	m.G0 = v8
	v10 = F_mstime(m)
	mBase = m.M
	v12 = *(*int64)(unsafe.Add(mBase, _consts[117]))
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+2352))
	*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = v13
	*(*int32)(unsafe.Add(mBase, uint32(v8))) = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v8)+20)) = int32(128)
	*(*int64)(unsafe.Add(mBase, uint32(v8)+12)) = v3
	*(*int64)(unsafe.Add(mBase, uint32(v8)+296)) = v3
	*(*int64)(unsafe.Add(mBase, uint32(v8)+160)) = int64(137438953472)
	*(*int32)(unsafe.Add(mBase, uint32(v8)+8)) = v8 + int32(24)
	*(*int32)(unsafe.Add(mBase, uint32(v8)+156)) = v8 + int32(168)
	goto L1
L1:
	;
	v32 = int32(0)
	v34 = F_raxSeek(m, v8, int32(_a4), v32, v32)
	mBase = m.M
	v35 = m.ExcPending
	if v35 != 0 {
		goto L2
	} else {
		goto L3
	}
L2:
	;
	return
L3:
	;
	v36 = F_raxNext(m, v8)
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		goto L2
	} else {
		goto L6
	}
L4:
	;
	F__serverAssert(m, int32(_a234), int32(_a179), int32(2092))
	mBase = m.M
	v114 = m.ExcPending
	if v114 != 0 {
		goto L2
	} else {
		goto L17
	}
L5:
	;
	F_raxStop(m, v8)
	mBase = m.M
	v106 = m.ExcPending
	if v106 != 0 {
		goto L2
	} else {
		goto L16
	}
L6:
	;
	if v36 == int32(0) {
		goto L5
	} else {
		goto L7
	}
L7:
	;
	goto L8
L8:
	;
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v8)+8))
	v49 = *(*int64)(unsafe.Add(mBase, uint32(v48)))
	v50 = int64(56)
	v52 = int64(65280)
	v54 = int64(40)
	v57 = int64(16711680)
	v59 = int64(24)
	v61 = int64(4278190080)
	v63 = int64(8)
	if v10-v12<<(uint(int64(1))%64) < v49<<(uint(v50)%64)|v49&v52<<(uint(v54)%64)|(v49&v57<<(uint(v59)%64)|v49&v61<<(uint(v63)%64))|(int64(base.Ui64(v49)>>(uint(v63)%64))&v61|int64(base.Ui64(v49)>>(uint(v59)%64))&v57|(int64(base.Ui64(v49)>>(uint(v54)%64))&v52|int64(base.Ui64(v49)>>(uint(v50)%64)))) {
		goto L5
	} else {
		goto L10
	}
L9:
	;
	goto L5
L10:
	;
	v86 = *(*int32)(unsafe.Add(mBase, uint32(l0)+2352))
	v87 = *(*int32)(unsafe.Add(mBase, uint32(v8)+16))
	v89 = F_raxRemove(m, v86, v48, v87, int32(0))
	mBase = m.M
	v90 = m.ExcPending
	if v90 != 0 {
		goto L2
	} else {
		goto L11
	}
L11:
	;
	if v89 != int32(1) {
		goto L4
	} else {
		goto L12
	}
L12:
	;
	v94 = *(*int32)(unsafe.Add(mBase, uint32(v8)+8))
	v95 = *(*int32)(unsafe.Add(mBase, uint32(v8)+16))
	v96 = F_raxSeek(m, v8, int32(_a235), v94, v95)
	mBase = m.M
	v97 = m.ExcPending
	if v97 != 0 {
		goto L2
	} else {
		goto L13
	}
L13:
	;
	v98 = F_raxNext(m, v8)
	mBase = m.M
	v99 = m.ExcPending
	if v99 != 0 {
		goto L2
	} else {
		goto L14
	}
L14:
	;
	if v98 != 0 {
		goto L8
	} else {
		goto L15
	}
L15:
	;
	goto L9
L16:
	;
	m.G0 = v8 + int32(304)
	return
L17:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_clusterNodeCoversSlot(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v8 int32
	_ = v8
	v4 = base.I32_div_s(l1, int32(8))
	v8 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0+v4+int32(104)))))
	return int32(base.Ui32(v8)>>(uint(l1&int32(7))%32)) & int32(1)
}
func F_clusterNodeHostname(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+2312))
	return v2
}
func F_clusterNodeIsReplica(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+88))
	return v2 & int32(2)
}
func F_clusterNodeNameComparator(m *base.Module, l0 int32, l1 int32) int32 {
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
	var v8 int32
	_ = v8
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
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
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v4 = int32(8)
	v5 = v3 + v4
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v8 = v6 + v4
	goto L2
L1:
	;
	return v52 - v54
L2:
	;
	v13 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5))))
	if v13 != 0 {
		goto L5
	} else {
		goto L6
	}
L4:
	;
	v52 = F_tolower(m, v47)
	mBase = m.M
	v53 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v48))))
	v54 = F_tolower(m, v53)
	mBase = m.M
	goto L1
L5:
	;
	v15 = v5
	v16 = v8
	v17 = int32(40)
	v18 = v13
	goto L8
L6:
	;
	v47 = int32(0)
	v48 = v8
	goto L4
L7:
	;
	v47 = v44 & int32(255)
	v48 = v42
	goto L4
L8:
	;
	v20 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16))))
	if v20 == int32(0) {
		v42 = v16
		v44 = v18
		goto L7
	} else {
		goto L10
	}
L9:
	;
	v42 = v36
	v44 = int32(0)
	goto L7
L10:
	;
	v24 = v17 + int32(-1)
	if v24 == int32(0) {
		v42 = v16
		v44 = v18
		goto L7
	} else {
		goto L11
	}
L11:
	;
	v28 = v18 & int32(255)
	if v28 == v20 {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	v35 = int32(1)
	v36 = v16 + v35
	v37 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15)+1)))
	if v37 != 0 {
		v15 = v15 + v35
		v16 = v36
		v17 = v24
		v18 = v37
		goto L8
	} else {
		goto L15
	}
L13:
	;
	v30 = F_tolower(m, v28)
	mBase = m.M
	v31 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16))))
	v32 = F_tolower(m, v31)
	mBase = m.M
	if v30 == v32 {
		goto L12
	} else {
		goto L14
	}
L14:
	;
	v34 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15))))
	v42 = v16
	v44 = v34
	goto L7
L15:
	;
	goto L9
}
func F_clusterNodeSetSlotBit(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v56 int32
	_ = v56
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v69 int64
	_ = v69
	var v70 int64
	_ = v70
	var v71 int64
	_ = v71
	var v72 int64
	_ = v72
	var v73 int64
	_ = v73
	var v74 int64
	_ = v74
	var v75 int64
	_ = v75
	var v77 int64
	_ = v77
	var v79 int64
	_ = v79
	var v81 int64
	_ = v81
	var v83 int64
	_ = v83
	var v85 int64
	_ = v85
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v107 int32
	_ = v107
	var v111 int32
	_ = v111
	var v113 int32
	_ = v113
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v127 int32
	_ = v127
	var v132 int32
	_ = v132
	var v136 int32
	_ = v136
	var v140 int32
	_ = v140
	var v147 int32
	_ = v147
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v155 int32
	_ = v155
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v167 int32
	_ = v167
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v175 int32
	_ = v175
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v184 int32
	_ = v184
	var v186 int32
	_ = v186
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
	var v193 int64
	_ = v193
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
	var v204 int64
	_ = v204
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v211 int32
	_ = v211
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v218 int32
	_ = v218
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v226 int32
	_ = v226
	var v230 int32
	_ = v230
	var v232 int32
	_ = v232
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v246 int32
	_ = v246
	var v251 int32
	_ = v251
	var v255 int32
	_ = v255
	var v259 int32
	_ = v259
	var v263 int32
	_ = v263
	v6 = m.G0
	v8 = v6 - int32(32)
	m.G0 = v8
	v13 = int32(1) << (uint(l1&int32(7)) % 32)
	v15 = base.I32_div_s(l1, int32(8))
	v18 = l0 + v15 + int32(104)
	v19 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18))))
	if v13&v19 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	m.G0 = v8 + int32(32)
	return
L2:
	;
	v21 = v19 | v13
	*(*uint8)(unsafe.Add(mBase, uint32(v18))) = uint8(v21)
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+2160))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+2160)) = v23 + int32(1)
	if v23 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v28 = *(*int32)(unsafe.Add(mBase, _consts[90]))
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v28)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v8))) = v29
	*(*int32)(unsafe.Add(mBase, uint32(v8)+20)) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v8)+12)) = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v8)+4)) = int64(4294967295)
	goto L4
L4:
	;
	v44 = v8 + int32(20)
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v8)+16))
	if v45 != 0 {
		goto L7
	} else {
		goto L8
	}
L5:
	;
	if v140 == int32(0) {
		goto L1
	} else {
		goto L31
	}
L6:
	;
	v51 = v44
	v52 = v48
	goto L9
L7:
	;
	v48 = int32(1)
	goto L6
L8:
	;
	v48 = int32(0)
	goto L6
L9:
	;
	switch v52 {
	case 0:
		goto L14
	default:
		goto L13
	}
L11:
	;
	v52 = int32(0)
	goto L9
L12:
	;
	goto L5
L13:
	;
	v132 = *(*int32)(unsafe.Add(mBase, uint32(v51)))
	*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = v132
	if v132 == int32(0) {
		goto L11
	} else {
		goto L30
	}
L14:
	;
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v8)+4))
	if v56 != int32(-1) {
		v95 = v56
		goto L15
	} else {
		goto L16
	}
L15:
	;
	v96 = int32(1)
	v97 = v95 + v96
	*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = v97
	v99 = int32(0)
	v102 = *(*int32)(unsafe.Add(mBase, uint32(v8)))
	v103 = *(*int32)(unsafe.Add(mBase, uint32(v8)+8))
	v107 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v102+v103+int32(26)))))
	if v107 == int32(255) {
		goto L24
	} else {
		goto L25
	}
L16:
	;
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v8)+8))
	if v60 != 0 {
		v95 = int32(-1)
		goto L15
	} else {
		goto L17
	}
L17:
	;
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v8)))
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v8)+12))
	if v62 == int32(0) {
		goto L19
	} else {
		goto L20
	}
L18:
	;
	v89 = *(*int32)(unsafe.Add(mBase, uint32(v88)+20))
	if v89 != int32(-1) {
		goto L21
	} else {
		goto L22
	}
L19:
	;
	v69 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v61)+16)))
	v70 = int64(*(*int8)(unsafe.Add(mBase, uint32(v61)+27)))
	v71 = int64(*(*int32)(unsafe.Add(mBase, uint32(v61)+8)))
	v72 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v61)+12)))
	v73 = int64(*(*int8)(unsafe.Add(mBase, uint32(v61)+26)))
	v74 = int64(*(*int32)(unsafe.Add(mBase, uint32(v61)+4)))
	v75 = F_wangHash64(m, v74)
	mBase = m.M
	v77 = F_wangHash64(m, v73+v75)
	mBase = m.M
	v79 = F_wangHash64(m, v72+v77)
	mBase = m.M
	v81 = F_wangHash64(m, v71+v79)
	mBase = m.M
	v83 = F_wangHash64(m, v70+v81)
	mBase = m.M
	v85 = F_wangHash64(m, v69+v83)
	mBase = m.M
	*(*int64)(unsafe.Add(mBase, uint32(v8)+24)) = v85
	v87 = *(*int32)(unsafe.Add(mBase, uint32(v8)))
	v88 = v87
	goto L18
L20:
	;
	v65 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v61)+24)))
	v67 = v65 + int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v61)+24)) = uint16(v67)
	v88 = v61
	goto L18
L21:
	;
	v95 = v89 + int32(-1)
	goto L15
L22:
	;
	v92 = *(*int32)(unsafe.Add(mBase, uint32(v8)+4))
	v95 = v92
	goto L15
L23:
	;
	v122 = int32(2)
	v127 = *(*int32)(unsafe.Add(mBase, uint32(v102+v120<<(uint(v122)%32)+int32(4))))
	v51 = v127 + v121<<(uint(v122)%32)
	v52 = int32(1)
	goto L9
L24:
	;
	v111 = v99
	goto L26
L25:
	;
	v111 = v96 << (uint(v107) % 32)
	goto L26
L26:
	;
	if v97 < v111 {
		v120 = v103
		v121 = v97
		goto L23
	} else {
		goto L27
	}
L27:
	;
	if v103 != 0 {
		v140 = v99
		goto L12
	} else {
		goto L28
	}
L28:
	;
	v113 = *(*int32)(unsafe.Add(mBase, uint32(v102)+20))
	if v113 == int32(-1) {
		v140 = v99
		goto L12
	} else {
		goto L29
	}
L29:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v8)+4)) = int64(4294967296)
	v120 = int32(1)
	v121 = int32(0)
	goto L23
L30:
	;
	v136 = *(*int32)(unsafe.Add(mBase, uint32(v132)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v44))) = v136
	v140 = v132
	goto L12
L31:
	;
	v147 = v140
	goto L33
L32:
	;
	v263 = *(*int32)(unsafe.Add(mBase, uint32(l0)+88))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+88)) = v263 | int32(256)
	goto L1
L33:
	;
	v151 = *(*int32)(unsafe.Add(mBase, uint32(v147)+8))
	goto L36
L35:
	;
	v163 = v8 + int32(20)
	v164 = *(*int32)(unsafe.Add(mBase, uint32(v8)+16))
	if v164 != 0 {
		goto L41
	} else {
		goto L42
	}
L36:
	;
	v152 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v151)+88)))
	if v152&int32(2) != 0 {
		goto L35
	} else {
		goto L37
	}
L37:
	;
	v155 = *(*int32)(unsafe.Add(mBase, uint32(v151)+2164))
	if v155 != 0 {
		goto L32
	} else {
		goto L38
	}
L38:
	;
	goto L35
L39:
	;
	if v259 != 0 {
		v147 = v259
		goto L33
	} else {
		goto L65
	}
L40:
	;
	v170 = v163
	v171 = v167
	goto L43
L41:
	;
	v167 = int32(1)
	goto L40
L42:
	;
	v167 = int32(0)
	goto L40
L43:
	;
	switch v171 {
	case 0:
		goto L48
	default:
		goto L47
	}
L45:
	;
	v171 = int32(0)
	goto L43
L46:
	;
	goto L39
L47:
	;
	v251 = *(*int32)(unsafe.Add(mBase, uint32(v170)))
	*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = v251
	if v251 == int32(0) {
		goto L45
	} else {
		goto L64
	}
L48:
	;
	v175 = *(*int32)(unsafe.Add(mBase, uint32(v8)+4))
	if v175 != int32(-1) {
		v214 = v175
		goto L49
	} else {
		goto L50
	}
L49:
	;
	v215 = int32(1)
	v216 = v214 + v215
	*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = v216
	v218 = int32(0)
	v221 = *(*int32)(unsafe.Add(mBase, uint32(v8)))
	v222 = *(*int32)(unsafe.Add(mBase, uint32(v8)+8))
	v226 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v221+v222+int32(26)))))
	if v226 == int32(255) {
		goto L58
	} else {
		goto L59
	}
L50:
	;
	v179 = *(*int32)(unsafe.Add(mBase, uint32(v8)+8))
	if v179 != 0 {
		v214 = int32(-1)
		goto L49
	} else {
		goto L51
	}
L51:
	;
	v180 = *(*int32)(unsafe.Add(mBase, uint32(v8)))
	v181 = *(*int32)(unsafe.Add(mBase, uint32(v8)+12))
	if v181 == int32(0) {
		goto L53
	} else {
		goto L54
	}
L52:
	;
	v208 = *(*int32)(unsafe.Add(mBase, uint32(v207)+20))
	if v208 != int32(-1) {
		goto L55
	} else {
		goto L56
	}
L53:
	;
	v188 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v180)+16)))
	v189 = int64(*(*int8)(unsafe.Add(mBase, uint32(v180)+27)))
	v190 = int64(*(*int32)(unsafe.Add(mBase, uint32(v180)+8)))
	v191 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v180)+12)))
	v192 = int64(*(*int8)(unsafe.Add(mBase, uint32(v180)+26)))
	v193 = int64(*(*int32)(unsafe.Add(mBase, uint32(v180)+4)))
	v194 = F_wangHash64(m, v193)
	mBase = m.M
	v196 = F_wangHash64(m, v192+v194)
	mBase = m.M
	v198 = F_wangHash64(m, v191+v196)
	mBase = m.M
	v200 = F_wangHash64(m, v190+v198)
	mBase = m.M
	v202 = F_wangHash64(m, v189+v200)
	mBase = m.M
	v204 = F_wangHash64(m, v188+v202)
	mBase = m.M
	*(*int64)(unsafe.Add(mBase, uint32(v8)+24)) = v204
	v206 = *(*int32)(unsafe.Add(mBase, uint32(v8)))
	v207 = v206
	goto L52
L54:
	;
	v184 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v180)+24)))
	v186 = v184 + int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v180)+24)) = uint16(v186)
	v207 = v180
	goto L52
L55:
	;
	v214 = v208 + int32(-1)
	goto L49
L56:
	;
	v211 = *(*int32)(unsafe.Add(mBase, uint32(v8)+4))
	v214 = v211
	goto L49
L57:
	;
	v241 = int32(2)
	v246 = *(*int32)(unsafe.Add(mBase, uint32(v221+v239<<(uint(v241)%32)+int32(4))))
	v170 = v246 + v240<<(uint(v241)%32)
	v171 = int32(1)
	goto L43
L58:
	;
	v230 = v218
	goto L60
L59:
	;
	v230 = v215 << (uint(v226) % 32)
	goto L60
L60:
	;
	if v216 < v230 {
		v239 = v222
		v240 = v216
		goto L57
	} else {
		goto L61
	}
L61:
	;
	if v222 != 0 {
		v259 = v218
		goto L46
	} else {
		goto L62
	}
L62:
	;
	v232 = *(*int32)(unsafe.Add(mBase, uint32(v221)+20))
	if v232 == int32(-1) {
		v259 = v218
		goto L46
	} else {
		goto L63
	}
L63:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v8)+4)) = int64(4294967296)
	v239 = int32(1)
	v240 = int32(0)
	goto L57
L64:
	;
	v255 = *(*int32)(unsafe.Add(mBase, uint32(v251)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v163))) = v255
	v259 = v251
	goto L46
L65:
	;
	goto L1
}
func F_clusterProcessPacket(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v41 int32
	_ = v41
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v65 int32
	_ = v65
	var v66 int64
	_ = v66
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v83 int32
	_ = v83
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v111 int32
	_ = v111
	var v118 int32
	_ = v118
	var v120 int64
	_ = v120
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v132 int32
	_ = v132
	var v134 int32
	_ = v134
	var v140 int32
	_ = v140
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v153 int32
	_ = v153
	var v159 int32
	_ = v159
	var v168 int32
	_ = v168
	var v170 int32
	_ = v170
	var v173 int32
	_ = v173
	var v183 int32
	_ = v183
	var v185 int32
	_ = v185
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v203 int32
	_ = v203
	var v206 int32
	_ = v206
	var v209 int32
	_ = v209
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
	var v236 int32
	_ = v236
	var v238 int32
	_ = v238
	var v248 int32
	_ = v248
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v255 int32
	_ = v255
	var v258 int32
	_ = v258
	var v259 int32
	_ = v259
	var v261 int32
	_ = v261
	var v266 int32
	_ = v266
	var v281 int32
	_ = v281
	var v282 int32
	_ = v282
	var v283 int32
	_ = v283
	var v290 int32
	_ = v290
	var v293 int32
	_ = v293
	var v300 int32
	_ = v300
	var v306 int32
	_ = v306
	var v307 int32
	_ = v307
	var v308 int32
	_ = v308
	var v316 int32
	_ = v316
	var v317 int32
	_ = v317
	var v318 int32
	_ = v318
	var v321 int32
	_ = v321
	var v322 int32
	_ = v322
	var v324 int32
	_ = v324
	var v325 int32
	_ = v325
	var v327 int32
	_ = v327
	var v329 int32
	_ = v329
	var v339 int32
	_ = v339
	var v340 int32
	_ = v340
	var v341 int32
	_ = v341
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
	var v357 int32
	_ = v357
	var v372 int32
	_ = v372
	var v373 int32
	_ = v373
	var v376 int32
	_ = v376
	var v377 int32
	_ = v377
	var v378 int32
	_ = v378
	var v379 int32
	_ = v379
	var v380 int32
	_ = v380
	var v381 int32
	_ = v381
	var v382 int32
	_ = v382
	var v387 int32
	_ = v387
	var v391 int64
	_ = v391
	var v393 int32
	_ = v393
	var v398 int32
	_ = v398
	var v403 int32
	_ = v403
	var v413 int64
	_ = v413
	var v414 int64
	_ = v414
	var v416 int64
	_ = v416
	var v418 int64
	_ = v418
	var v420 int64
	_ = v420
	var v423 int64
	_ = v423
	var v425 int64
	_ = v425
	var v427 int64
	_ = v427
	var v448 int64
	_ = v448
	var v483 int64
	_ = v483
	var v485 int32
	_ = v485
	var v486 int64
	_ = v486
	var v489 int64
	_ = v489
	var v490 int64
	_ = v490
	var v494 int32
	_ = v494
	var v496 int32
	_ = v496
	var v497 int32
	_ = v497
	var v499 int32
	_ = v499
	var v501 int64
	_ = v501
	var v504 int32
	_ = v504
	var v507 int64
	_ = v507
	var v508 int64
	_ = v508
	var v513 int32
	_ = v513
	var v516 int32
	_ = v516
	var v517 int32
	_ = v517
	var v520 int64
	_ = v520
	var v532 int32
	_ = v532
	var v534 int32
	_ = v534
	var v535 int32
	_ = v535
	var v536 int32
	_ = v536
	var v537 int32
	_ = v537
	var v541 int32
	_ = v541
	var v544 int64
	_ = v544
	var v545 int64
	_ = v545
	var v547 int64
	_ = v547
	var v549 int64
	_ = v549
	var v552 int64
	_ = v552
	var v554 int64
	_ = v554
	var v556 int64
	_ = v556
	var v558 int64
	_ = v558
	var v579 int64
	_ = v579
	var v581 int64
	_ = v581
	var v584 int32
	_ = v584
	var v585 int32
	_ = v585
	var v586 int32
	_ = v586
	var v591 int32
	_ = v591
	var v593 int32
	_ = v593
	var v598 int64
	_ = v598
	var v602 int32
	_ = v602
	var v607 int32
	_ = v607
	var v616 int32
	_ = v616
	var v617 int64
	_ = v617
	var v618 int64
	_ = v618
	var v628 int32
	_ = v628
	var v629 int32
	_ = v629
	var v631 int32
	_ = v631
	var v635 int32
	_ = v635
	var v636 int32
	_ = v636
	var v639 int32
	_ = v639
	var v640 int32
	_ = v640
	var v646 int32
	_ = v646
	var v648 int32
	_ = v648
	var v649 int32
	_ = v649
	var v653 int32
	_ = v653
	var v654 int32
	_ = v654
	var v655 int32
	_ = v655
	var v657 int32
	_ = v657
	var v660 int32
	_ = v660
	var v661 int32
	_ = v661
	var v667 int32
	_ = v667
	var v668 int32
	_ = v668
	var v671 int32
	_ = v671
	var v672 int32
	_ = v672
	var v675 int32
	_ = v675
	var v684 int32
	_ = v684
	var v685 int32
	_ = v685
	var v691 int64
	_ = v691
	var v695 int64
	_ = v695
	var v703 int64
	_ = v703
	var v711 int64
	_ = v711
	var v719 int64
	_ = v719
	var v727 int64
	_ = v727
	var v730 int32
	_ = v730
	var v739 int32
	_ = v739
	var v741 int32
	_ = v741
	var v743 int32
	_ = v743
	var v744 int32
	_ = v744
	var v754 int32
	_ = v754
	var v757 int64
	_ = v757
	var v762 int32
	_ = v762
	var v768 int32
	_ = v768
	var v774 int32
	_ = v774
	var v781 int32
	_ = v781
	var v792 int64
	_ = v792
	var v796 int64
	_ = v796
	var v800 int64
	_ = v800
	var v804 int64
	_ = v804
	var v806 int64
	_ = v806
	var v810 int64
	_ = v810
	var v812 int32
	_ = v812
	var v814 int32
	_ = v814
	var v817 int32
	_ = v817
	var v818 int32
	_ = v818
	var v826 int32
	_ = v826
	var v827 int32
	_ = v827
	var v832 int32
	_ = v832
	var v835 int32
	_ = v835
	var v837 int32
	_ = v837
	var v838 int32
	_ = v838
	var v839 int32
	_ = v839
	var v840 int32
	_ = v840
	var v841 int32
	_ = v841
	var v848 int32
	_ = v848
	var v850 int32
	_ = v850
	var v857 int32
	_ = v857
	var v859 int32
	_ = v859
	var v862 int32
	_ = v862
	var v865 int32
	_ = v865
	var v871 int32
	_ = v871
	var v872 int32
	_ = v872
	var v875 int64
	_ = v875
	var v878 int64
	_ = v878
	var v879 int64
	_ = v879
	var v882 int64
	_ = v882
	var v886 int32
	_ = v886
	var v889 int32
	_ = v889
	var v890 int32
	_ = v890
	var v900 int32
	_ = v900
	var v901 int32
	_ = v901
	var v902 int32
	_ = v902
	var v904 int32
	_ = v904
	var v909 int32
	_ = v909
	var v915 int32
	_ = v915
	var v921 int32
	_ = v921
	var v922 int32
	_ = v922
	var v927 int64
	_ = v927
	var v935 int64
	_ = v935
	var v941 int64
	_ = v941
	var v947 int64
	_ = v947
	var v951 int64
	_ = v951
	var v953 int64
	_ = v953
	var v964 int32
	_ = v964
	var v965 int32
	_ = v965
	var v967 int32
	_ = v967
	var v970 int32
	_ = v970
	var v972 int32
	_ = v972
	var v973 int32
	_ = v973
	var v975 int32
	_ = v975
	var v977 int32
	_ = v977
	var v978 int32
	_ = v978
	var v980 int32
	_ = v980
	var v985 int32
	_ = v985
	var v990 int32
	_ = v990
	var v992 int32
	_ = v992
	var v994 int32
	_ = v994
	var v996 int32
	_ = v996
	var v997 int32
	_ = v997
	var v1002 int32
	_ = v1002
	var v1010 int32
	_ = v1010
	var v1017 int32
	_ = v1017
	var v1020 int32
	_ = v1020
	var v1023 int32
	_ = v1023
	var v1024 int32
	_ = v1024
	var v1033 int32
	_ = v1033
	var v1034 int32
	_ = v1034
	var v1036 int32
	_ = v1036
	var v1048 int32
	_ = v1048
	var v1049 int32
	_ = v1049
	var v1050 int32
	_ = v1050
	var v1051 int32
	_ = v1051
	var v1055 int32
	_ = v1055
	var v1058 int32
	_ = v1058
	var v1059 int32
	_ = v1059
	var v1060 int32
	_ = v1060
	var v1063 int32
	_ = v1063
	var v1076 int32
	_ = v1076
	var v1081 int32
	_ = v1081
	var v1082 int32
	_ = v1082
	var v1086 int32
	_ = v1086
	var v1089 int32
	_ = v1089
	var v1090 int32
	_ = v1090
	var v1091 int32
	_ = v1091
	var v1099 int32
	_ = v1099
	var v1108 int32
	_ = v1108
	var v1117 int32
	_ = v1117
	var v1120 int32
	_ = v1120
	var v1121 int32
	_ = v1121
	var v1131 int32
	_ = v1131
	var v1133 int32
	_ = v1133
	var v1134 int32
	_ = v1134
	var v1135 int32
	_ = v1135
	var v1143 int32
	_ = v1143
	var v1146 int32
	_ = v1146
	var v1147 int32
	_ = v1147
	var v1157 int32
	_ = v1157
	var v1158 int32
	_ = v1158
	var v1159 int32
	_ = v1159
	var v1163 int32
	_ = v1163
	var v1165 int32
	_ = v1165
	var v1166 int32
	_ = v1166
	var v1171 int32
	_ = v1171
	var v1173 int32
	_ = v1173
	var v1178 int32
	_ = v1178
	var v1180 int32
	_ = v1180
	var v1183 int32
	_ = v1183
	var v1184 int32
	_ = v1184
	var v1185 int32
	_ = v1185
	var v1195 int32
	_ = v1195
	var v1197 int32
	_ = v1197
	var v1198 int32
	_ = v1198
	var v1206 int32
	_ = v1206
	var v1208 int32
	_ = v1208
	var v1209 int32
	_ = v1209
	var v1214 int32
	_ = v1214
	var v1215 int32
	_ = v1215
	var v1216 int32
	_ = v1216
	var v1225 int32
	_ = v1225
	var v1226 int32
	_ = v1226
	var v1227 int32
	_ = v1227
	var v1230 int32
	_ = v1230
	var v1231 int32
	_ = v1231
	var v1233 int32
	_ = v1233
	var v1234 int32
	_ = v1234
	var v1236 int32
	_ = v1236
	var v1238 int32
	_ = v1238
	var v1248 int32
	_ = v1248
	var v1249 int32
	_ = v1249
	var v1250 int32
	_ = v1250
	var v1253 int32
	_ = v1253
	var v1254 int32
	_ = v1254
	var v1255 int32
	_ = v1255
	var v1258 int32
	_ = v1258
	var v1259 int32
	_ = v1259
	var v1261 int32
	_ = v1261
	var v1266 int32
	_ = v1266
	var v1281 int32
	_ = v1281
	var v1285 int32
	_ = v1285
	var v1288 int32
	_ = v1288
	var v1289 int32
	_ = v1289
	var v1290 int32
	_ = v1290
	var v1291 int64
	_ = v1291
	var v1294 int32
	_ = v1294
	var v1301 int64
	_ = v1301
	var v1308 int32
	_ = v1308
	var v1309 int32
	_ = v1309
	var v1310 int32
	_ = v1310
	var v1311 int32
	_ = v1311
	var v1314 int32
	_ = v1314
	var v1315 int32
	_ = v1315
	var v1321 int32
	_ = v1321
	var v1327 int32
	_ = v1327
	var v1328 int32
	_ = v1328
	var v1332 int32
	_ = v1332
	var v1334 int32
	_ = v1334
	var v1336 int32
	_ = v1336
	var v1338 int32
	_ = v1338
	var v1339 int32
	_ = v1339
	var v1348 int32
	_ = v1348
	var v1349 int32
	_ = v1349
	var v1352 int32
	_ = v1352
	var v1355 int32
	_ = v1355
	var v1357 int32
	_ = v1357
	var v1358 int32
	_ = v1358
	var v1362 int32
	_ = v1362
	var v1363 int32
	_ = v1363
	var v1370 int32
	_ = v1370
	var v1371 int32
	_ = v1371
	var v1375 int32
	_ = v1375
	var v1377 int32
	_ = v1377
	var v1378 int32
	_ = v1378
	var v1386 int32
	_ = v1386
	var v1388 int32
	_ = v1388
	var v1397 int32
	_ = v1397
	var v1399 int32
	_ = v1399
	var v1402 int32
	_ = v1402
	var v1412 int32
	_ = v1412
	var v1414 int32
	_ = v1414
	var v1424 int32
	_ = v1424
	var v1425 int32
	_ = v1425
	var v1427 int32
	_ = v1427
	var v1428 int32
	_ = v1428
	var v1429 int32
	_ = v1429
	var v1430 int32
	_ = v1430
	var v1432 int32
	_ = v1432
	var v1435 int32
	_ = v1435
	var v1438 int32
	_ = v1438
	var v1442 int32
	_ = v1442
	var v1445 int32
	_ = v1445
	var v1446 int32
	_ = v1446
	var v1447 int32
	_ = v1447
	var v1448 int32
	_ = v1448
	var v1460 int32
	_ = v1460
	var v1463 int32
	_ = v1463
	var v1464 int32
	_ = v1464
	var v1466 int32
	_ = v1466
	var v1480 int32
	_ = v1480
	var v1485 int32
	_ = v1485
	var v1491 int32
	_ = v1491
	var v1494 int32
	_ = v1494
	var v1498 int32
	_ = v1498
	var v1502 int32
	_ = v1502
	var v1503 int32
	_ = v1503
	var v1508 int32
	_ = v1508
	var v1512 int32
	_ = v1512
	var v1513 int64
	_ = v1513
	var v1515 int32
	_ = v1515
	var v1516 int32
	_ = v1516
	var v1520 int32
	_ = v1520
	var v1524 int32
	_ = v1524
	var v1527 int32
	_ = v1527
	var v1529 int32
	_ = v1529
	var v1532 int32
	_ = v1532
	var v1533 int32
	_ = v1533
	var v1534 int64
	_ = v1534
	var v1536 int32
	_ = v1536
	var v1546 int32
	_ = v1546
	var v1548 int32
	_ = v1548
	var v1551 int32
	_ = v1551
	var v1552 int32
	_ = v1552
	var v1562 int32
	_ = v1562
	var v1565 int32
	_ = v1565
	var v1566 int32
	_ = v1566
	var v1569 int64
	_ = v1569
	var v1573 int32
	_ = v1573
	var v1574 int32
	_ = v1574
	var v1577 int64
	_ = v1577
	var v1578 int64
	_ = v1578
	var v1580 int64
	_ = v1580
	var v1582 int64
	_ = v1582
	var v1585 int64
	_ = v1585
	var v1587 int64
	_ = v1587
	var v1589 int64
	_ = v1589
	var v1591 int64
	_ = v1591
	var v1612 int64
	_ = v1612
	var v1615 int32
	_ = v1615
	var v1618 int32
	_ = v1618
	var v1619 int32
	_ = v1619
	var v1620 int32
	_ = v1620
	var v1621 int32
	_ = v1621
	var v1622 int64
	_ = v1622
	var v1631 int32
	_ = v1631
	var v1637 int32
	_ = v1637
	var v1652 int32
	_ = v1652
	var v1656 int32
	_ = v1656
	var v1662 int32
	_ = v1662
	var v1665 int32
	_ = v1665
	var v1667 int32
	_ = v1667
	var v1668 int32
	_ = v1668
	var v1675 int32
	_ = v1675
	var v1679 int32
	_ = v1679
	var v1681 int32
	_ = v1681
	var v1683 int32
	_ = v1683
	var v1692 int32
	_ = v1692
	var v1698 int32
	_ = v1698
	var v1701 int32
	_ = v1701
	var v1702 int32
	_ = v1702
	var v1706 int32
	_ = v1706
	var v1715 int32
	_ = v1715
	var v1717 int32
	_ = v1717
	var v1718 int32
	_ = v1718
	var v1727 int32
	_ = v1727
	var v1741 int32
	_ = v1741
	var v1744 int32
	_ = v1744
	var v1745 int32
	_ = v1745
	var v1746 int32
	_ = v1746
	var v1760 int32
	_ = v1760
	var v1763 int32
	_ = v1763
	var v1764 int32
	_ = v1764
	var v1773 int32
	_ = v1773
	var v1775 int32
	_ = v1775
	var v1778 int32
	_ = v1778
	var v1788 int32
	_ = v1788
	var v1790 int32
	_ = v1790
	var v1800 int32
	_ = v1800
	var v1801 int32
	_ = v1801
	var v1803 int32
	_ = v1803
	var v1804 int32
	_ = v1804
	var v1805 int32
	_ = v1805
	var v1806 int32
	_ = v1806
	var v1808 int32
	_ = v1808
	var v1811 int32
	_ = v1811
	var v1812 int32
	_ = v1812
	var v1818 int32
	_ = v1818
	var v1821 int32
	_ = v1821
	var v1822 int32
	_ = v1822
	var v1837 int32
	_ = v1837
	var v1841 int32
	_ = v1841
	var v1842 int32
	_ = v1842
	var v1844 int32
	_ = v1844
	var v1845 int32
	_ = v1845
	var v1853 int32
	_ = v1853
	var v1854 int32
	_ = v1854
	var v1855 int32
	_ = v1855
	var v1858 int32
	_ = v1858
	var v1859 int32
	_ = v1859
	var v1861 int32
	_ = v1861
	var v1862 int32
	_ = v1862
	var v1864 int32
	_ = v1864
	var v1866 int32
	_ = v1866
	var v1876 int32
	_ = v1876
	var v1877 int32
	_ = v1877
	var v1878 int32
	_ = v1878
	var v1881 int32
	_ = v1881
	var v1882 int32
	_ = v1882
	var v1883 int32
	_ = v1883
	var v1886 int32
	_ = v1886
	var v1887 int32
	_ = v1887
	var v1889 int32
	_ = v1889
	var v1894 int32
	_ = v1894
	var v1909 int32
	_ = v1909
	var v1910 int32
	_ = v1910
	var v1911 int32
	_ = v1911
	var v1912 int64
	_ = v1912
	var v1914 int32
	_ = v1914
	var v1915 int32
	_ = v1915
	var v1916 int64
	_ = v1916
	var v1919 int32
	_ = v1919
	var v1921 int32
	_ = v1921
	var v1924 int32
	_ = v1924
	var v1925 int32
	_ = v1925
	var v1926 int32
	_ = v1926
	var v1927 int32
	_ = v1927
	var v1928 int64
	_ = v1928
	var v1929 int32
	_ = v1929
	var v1930 int32
	_ = v1930
	var v1931 int64
	_ = v1931
	var v1938 int64
	_ = v1938
	var v1950 int32
	_ = v1950
	var v1951 int32
	_ = v1951
	var v1956 int32
	_ = v1956
	var v1969 int32
	_ = v1969
	var v1971 int32
	_ = v1971
	var v1974 int32
	_ = v1974
	var v1975 int32
	_ = v1975
	var v1980 int32
	_ = v1980
	var v1981 int32
	_ = v1981
	var v1982 int32
	_ = v1982
	var v1983 int32
	_ = v1983
	var v1984 int32
	_ = v1984
	var v1987 int64
	_ = v1987
	var v1994 int32
	_ = v1994
	var v2008 int32
	_ = v2008
	var v2010 int32
	_ = v2010
	var v2011 int32
	_ = v2011
	var v2014 int32
	_ = v2014
	var v2019 int32
	_ = v2019
	var v2020 int32
	_ = v2020
	var v2032 int32
	_ = v2032
	var v2034 int32
	_ = v2034
	var v2035 int32
	_ = v2035
	var v2036 int32
	_ = v2036
	var v2041 int32
	_ = v2041
	var v2042 int32
	_ = v2042
	var v2054 int32
	_ = v2054
	var v2056 int32
	_ = v2056
	var v2063 int32
	_ = v2063
	var v2067 int32
	_ = v2067
	var v2068 int32
	_ = v2068
	var v2070 int32
	_ = v2070
	var v2073 int32
	_ = v2073
	var v2074 int32
	_ = v2074
	var v2088 int32
	_ = v2088
	var v2092 int32
	_ = v2092
	var v2095 int32
	_ = v2095
	var v2096 int32
	_ = v2096
	var v2109 int32
	_ = v2109
	var v2112 int32
	_ = v2112
	var v2118 int32
	_ = v2118
	var v2125 int32
	_ = v2125
	var v2127 int32
	_ = v2127
	var v2128 int32
	_ = v2128
	var v2139 int32
	_ = v2139
	var v2150 int32
	_ = v2150
	var v2153 int32
	_ = v2153
	var v2160 int32
	_ = v2160
	var v2164 int32
	_ = v2164
	var v2167 int32
	_ = v2167
	var v2168 int32
	_ = v2168
	var v2171 int32
	_ = v2171
	var v2179 int32
	_ = v2179
	var v2180 int32
	_ = v2180
	var v2181 int32
	_ = v2181
	var v2183 int32
	_ = v2183
	var v2186 int32
	_ = v2186
	var v2210 int32
	_ = v2210
	var v2213 int32
	_ = v2213
	var v2214 int32
	_ = v2214
	var v2215 int32
	_ = v2215
	var v2216 int32
	_ = v2216
	var v2223 int32
	_ = v2223
	var v2235 int32
	_ = v2235
	var v2238 int32
	_ = v2238
	var v2239 int32
	_ = v2239
	var v2241 int32
	_ = v2241
	var v2242 int32
	_ = v2242
	var v2243 int32
	_ = v2243
	var v2246 int32
	_ = v2246
	var v2251 int32
	_ = v2251
	var v2254 int32
	_ = v2254
	var v2265 int32
	_ = v2265
	var v2267 int32
	_ = v2267
	var v2268 int32
	_ = v2268
	var v2269 int32
	_ = v2269
	var v2270 int32
	_ = v2270
	var v2271 int32
	_ = v2271
	var v2273 int32
	_ = v2273
	var v2274 int32
	_ = v2274
	var v2276 int32
	_ = v2276
	var v2277 int32
	_ = v2277
	var v2285 int32
	_ = v2285
	var v2286 int32
	_ = v2286
	var v2287 int32
	_ = v2287
	var v2290 int32
	_ = v2290
	var v2291 int32
	_ = v2291
	var v2293 int32
	_ = v2293
	var v2294 int32
	_ = v2294
	var v2296 int32
	_ = v2296
	var v2298 int32
	_ = v2298
	var v2308 int32
	_ = v2308
	var v2309 int32
	_ = v2309
	var v2310 int32
	_ = v2310
	var v2313 int32
	_ = v2313
	var v2314 int32
	_ = v2314
	var v2315 int32
	_ = v2315
	var v2318 int32
	_ = v2318
	var v2319 int32
	_ = v2319
	var v2321 int32
	_ = v2321
	var v2326 int32
	_ = v2326
	var v2341 int32
	_ = v2341
	var v2345 int32
	_ = v2345
	var v2347 int32
	_ = v2347
	var v2349 int32
	_ = v2349
	var v2350 int32
	_ = v2350
	var v2359 int32
	_ = v2359
	var v2361 int32
	_ = v2361
	var v2363 int32
	_ = v2363
	var v2364 int32
	_ = v2364
	var v2369 int32
	_ = v2369
	var v2371 int32
	_ = v2371
	var v2372 int32
	_ = v2372
	var v2380 int32
	_ = v2380
	var v2381 int32
	_ = v2381
	var v2382 int32
	_ = v2382
	var v2385 int32
	_ = v2385
	var v2386 int32
	_ = v2386
	var v2388 int32
	_ = v2388
	var v2389 int32
	_ = v2389
	var v2391 int32
	_ = v2391
	var v2393 int32
	_ = v2393
	var v2403 int32
	_ = v2403
	var v2404 int32
	_ = v2404
	var v2405 int32
	_ = v2405
	var v2408 int32
	_ = v2408
	var v2409 int32
	_ = v2409
	var v2410 int32
	_ = v2410
	var v2413 int32
	_ = v2413
	var v2414 int32
	_ = v2414
	var v2416 int32
	_ = v2416
	var v2421 int32
	_ = v2421
	var v2436 int32
	_ = v2436
	var v2440 int32
	_ = v2440
	var v2446 int32
	_ = v2446
	var v2449 int32
	_ = v2449
	var v2450 int32
	_ = v2450
	var v2463 int32
	_ = v2463
	var v2466 int32
	_ = v2466
	var v2468 int32
	_ = v2468
	var v2472 int32
	_ = v2472
	var v2490 int32
	_ = v2490
	var v2502 int64
	_ = v2502
	var v2512 int64
	_ = v2512
	var v2529 int32
	_ = v2529
	var v2530 int32
	_ = v2530
	var v2534 int32
	_ = v2534
	var v2541 int32
	_ = v2541
	var v2547 int64
	_ = v2547
	var v2550 int32
	_ = v2550
	var v2553 int32
	_ = v2553
	var v2554 int32
	_ = v2554
	var v2555 int32
	_ = v2555
	var v2556 int32
	_ = v2556
	var v2568 int32
	_ = v2568
	var v2570 int32
	_ = v2570
	var v2572 int32
	_ = v2572
	var v2575 int64
	_ = v2575
	var v2624 int32
	_ = v2624
	var v2625 int32
	_ = v2625
	var v2626 int32
	_ = v2626
	var v2627 int32
	_ = v2627
	var v2632 int32
	_ = v2632
	var v2637 int64
	_ = v2637
	var v2640 int32
	_ = v2640
	var v2642 int32
	_ = v2642
	var v2644 int32
	_ = v2644
	var v2649 int32
	_ = v2649
	var v2655 int32
	_ = v2655
	var v2657 int32
	_ = v2657
	var v2658 int32
	_ = v2658
	var v2666 int32
	_ = v2666
	var v2667 int32
	_ = v2667
	var v2676 int32
	_ = v2676
	var v2696 int32
	_ = v2696
	var v2702 int32
	_ = v2702
	v22 = m.G0
	v24 = v22 - int32(672)
	m.G0 = v24
	v26 = F_clusterIsValidPacket(m, l0)
	mBase = m.M
	v29 = m.ExcPending
	if v29 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v26 != 0 {
		goto L4
	} else {
		goto L5
	}
L3:
	;
	m.G0 = v24 + int32(672)
	return v2702
L4:
	;
	v66 = F_mstime(m)
	mBase = m.M
	v67 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v30)+12)))
	v68 = F___bswap_16_2(m, v67)
	mBase = m.M
	goto L16
L5:
	;
	v31 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v30)+12)))
	v32 = F___bswap_16_2(m, v31)
	mBase = m.M
	goto L6
L6:
	;
	v33 = int32(1)
	v35 = int32(*(*uint8)(unsafe.Add(mBase, _consts[123])))
	if v35&v33 == int32(0) {
		v2702 = v33
		goto L3
	} else {
		goto L7
	}
L7:
	;
	v41 = *(*int32)(unsafe.Add(mBase, _consts[124]))
	if v41 == v32 {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	F_freeClusterLink(m, l0)
	mBase = m.M
	v46 = m.ExcPending
	if v46 != 0 {
		goto L1
	} else {
		goto L11
	}
L9:
	;
	if v41 != int32(-2) {
		v2702 = v33
		goto L3
	} else {
		goto L10
	}
L10:
	;
	goto L8
L11:
	;
	v47 = int32(0)
	v49 = *(*int32)(unsafe.Add(mBase, _consts[15]))
	if int32(3) < v49 {
		v2702 = v47
		goto L3
	} else {
		goto L12
	}
L12:
	;
	if base.Ui32(int32(10)) < base.Ui32(v32) {
		v60 = int32(_a242)
		goto L13
	} else {
		goto L14
	}
L13:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24))) = v60
	F__serverLog(m, int32(3), int32(_a243), v24)
	mBase = m.M
	v65 = m.ExcPending
	if v65 != 0 {
		goto L1
	} else {
		goto L15
	}
L14:
	;
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v32<<(uint(int32(2))%32))+uint32(_consts[125])))
	v60 = v59
	goto L13
L15:
	;
	v2702 = v47
	goto L3
L16:
	;
	v70 = v68 & int32(32767)
	v71 = base.I32_extend16_s(v68)
	if int32(-1) < v71 {
		goto L23
	} else {
		goto L24
	}
L17:
	;
	F_freeClusterLink(m, l0)
	mBase = m.M
	v2696 = m.ExcPending
	if v2696 != 0 {
		goto L1
	} else {
		goto L597
	}
L18:
	;
	if base.Ui32(int32(10)) < base.Ui32(v70) {
		v2667 = int32(_a242)
		goto L594
	} else {
		goto L595
	}
L19:
	;
	v382 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v142)+2253)))
	if v382&int32(4) == int32(0) {
		goto L102
	} else {
		goto L103
	}
L20:
	;
	v306 = v142 + int32(2128)
	v307 = int32(_a244)
	v308 = int32(40)
	goto L90
L21:
	;
	F__serverAssert(m, int32(_a245), int32(_a179), int32(145))
	mBase = m.M
	v300 = m.ExcPending
	if v300 != 0 {
		goto L1
	} else {
		goto L85
	}
L22:
	;
	v290 = int32(_a246)
	v293 = *(*int32)(unsafe.Add(mBase, _consts[15]))
	if v293 <= int32(2) {
		v2657 = v290
		v2658 = v290
		goto L18
	} else {
		goto L84
	}
L23:
	;
	v142 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	v143 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v142)+12)))
	v144 = F___bswap_16_2(m, v143)
	mBase = m.M
	goto L44
L24:
	;
	v74 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	if v74 == int32(0) {
		goto L22
	} else {
		goto L25
	}
L25:
	;
	v77 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v74)+88)))
	if v77&int32(32) == int32(0) {
		goto L26
	} else {
		goto L27
	}
L26:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v74)+2200)) = v66
	v91 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	v93 = *(*int32)(unsafe.Add(mBase, _consts[15]))
	if int32(0) < v93 {
		goto L30
	} else {
		goto L31
	}
L27:
	;
	v83 = *(*int32)(unsafe.Add(mBase, _consts[15]))
	if int32(2) < v83 {
		goto L17
	} else {
		goto L28
	}
L28:
	;
	v88 = F_humanNodename(m, v74)
	mBase = m.M
	v89 = m.ExcPending
	if v89 != 0 {
		goto L1
	} else {
		goto L29
	}
L29:
	;
	v2657 = v88
	v2658 = v74 + int32(8)
	goto L18
L30:
	;
	switch v70 + int32(-4) {
	case 0, 6:
		goto L37
	default:
		goto L35
	case 5:
		goto L36
	}
L31:
	;
	if base.Ui32(int32(10)) < base.Ui32(v70) {
		v104 = int32(_a242)
		goto L32
	} else {
		goto L33
	}
L32:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+608)) = v104
	F__serverLog(m, int32(0), int32(_a247), v24+int32(608))
	mBase = m.M
	v111 = m.ExcPending
	if v111 != 0 {
		goto L1
	} else {
		goto L34
	}
L33:
	;
	v103 = *(*int32)(unsafe.Add(mBase, uint32(v70<<(uint(int32(2))%32))+uint32(_consts[125])))
	v104 = v103
	goto L32
L34:
	;
	goto L30
L35:
	;
	F__serverAssert(m, int32(_a107), int32(_a179), int32(3641))
	mBase = m.M
	v140 = m.ExcPending
	if v140 != 0 {
		goto L1
	} else {
		goto L43
	}
L36:
	;
	v120 = *(*int64)(unsafe.Add(mBase, uint32(v91)+16))
	v121 = *(*int32)(unsafe.Add(mBase, uint32(v91)+24))
	v122 = F___bswap_32_2(m, v121)
	mBase = m.M
	goto L39
L37:
	;
	F_clusterProcessPublishPacket(m, v91+int32(16), v70)
	mBase = m.M
	v118 = m.ExcPending
	if v118 != 0 {
		goto L1
	} else {
		goto L38
	}
L38:
	;
	v2702 = int32(1)
	goto L3
L39:
	;
	v123 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v91)+28)))
	v127 = F_sdsnewlen(m, v74+int32(8), int32(40))
	mBase = m.M
	v128 = m.ExcPending
	if v128 != 0 {
		goto L1
	} else {
		goto L40
	}
L40:
	;
	F_moduleCallClusterReceivers(m, v127, v120, v123, v91+int32(29), v122)
	mBase = m.M
	v132 = m.ExcPending
	if v132 != 0 {
		goto L1
	} else {
		goto L41
	}
L41:
	;
	F_sdsfree(m, v127)
	mBase = m.M
	v134 = m.ExcPending
	if v134 != 0 {
		goto L1
	} else {
		goto L42
	}
L42:
	;
	v2702 = int32(1)
	goto L3
L43:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L44:
	;
	if base.I32_extend16_s(v144) <= int32(-1) {
		goto L21
	} else {
		goto L45
	}
L45:
	;
	v148 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v142)+2250)))
	v149 = F___bswap_16_2(m, v148)
	mBase = m.M
	goto L46
L46:
	;
	v150 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	if v150 == int32(0) {
		goto L48
	} else {
		goto L49
	}
L47:
	;
	v215 = v142 + int32(2128)
	v216 = int32(_a244)
	v217 = int32(40)
	goto L72
L48:
	;
	v159 = v142 + int32(40)
	goto L53
L49:
	;
	v153 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v150)+88)))
	if v153&int32(32) == int32(0) {
		v212 = v150
		goto L47
	} else {
		goto L50
	}
L50:
	;
	goto L48
L51:
	;
	if int32(0)-v183 != 0 {
		goto L20
	} else {
		goto L59
	}
L52:
	;
	goto L51
L53:
	;
	v168 = int32(0)
	goto L55
L54:
	;
	goto L52
L55:
	;
	v170 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v159+v168))))
	v173 = int32(255)
	v183 = base.B2i32(base.Ui32((v170+int32(-123))&v173) < base.Ui32(int32(230))) & base.B2i32(base.Ui32((v170+int32(-58))&v173) < base.Ui32(int32(246)))
	if v183 != 0 {
		goto L54
	} else {
		goto L57
	}
L56:
	;
	goto L54
L57:
	;
	v185 = v168 + int32(1)
	if v185 != int32(40) {
		v168 = v185
		goto L55
	} else {
		goto L58
	}
L58:
	;
	goto L56
L59:
	;
	v195 = F_sdsnewlen(m, v159, int32(40))
	mBase = m.M
	v196 = m.ExcPending
	if v196 != 0 {
		goto L1
	} else {
		goto L60
	}
L60:
	;
	v198 = *(*int32)(unsafe.Add(mBase, _consts[90]))
	v199 = *(*int32)(unsafe.Add(mBase, uint32(v198)+32))
	v200 = F_dictFind(m, v199, v195)
	mBase = m.M
	v201 = m.ExcPending
	if v201 != 0 {
		goto L1
	} else {
		goto L61
	}
L61:
	;
	F_sdsfree(m, v195)
	mBase = m.M
	v203 = m.ExcPending
	if v203 != 0 {
		goto L1
	} else {
		goto L62
	}
L62:
	;
	if v200 == int32(0) {
		goto L20
	} else {
		goto L63
	}
L63:
	;
	v206 = *(*int32)(unsafe.Add(mBase, uint32(v200)+8))
	goto L64
L64:
	;
	if v206 == int32(0) {
		goto L20
	} else {
		goto L65
	}
L65:
	;
	v209 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	if v209 != 0 {
		v212 = v206
		goto L47
	} else {
		goto L66
	}
L66:
	;
	F_setClusterNodeToInboundClusterLink(m, v206, l0)
	mBase = m.M
	v211 = m.ExcPending
	if v211 != 0 {
		goto L1
	} else {
		goto L67
	}
L67:
	;
	v212 = v206
	goto L47
L68:
	;
	v282 = int32(1)
	v283 = *(*int32)(unsafe.Add(mBase, uint32(v212)+88))
	v376 = v212
	v377 = v282
	v378 = v215
	v379 = v281
	v380 = int32(base.Ui32(v283&int32(2)) >> (uint(v282) % 32))
	v381 = v283 & v282
	goto L19
L69:
	;
	v281 = int32(0)
	goto L68
L70:
	;
	v253 = v248
	v254 = v249
	v255 = v250
	goto L80
L71:
	;
	if v238 == int32(0) {
		goto L69
	} else {
		goto L78
	}
L72:
	;
	if (v216|v215)&int32(3) != 0 {
		v248 = v215
		v249 = v216
		v250 = v217
		goto L70
	} else {
		goto L73
	}
L73:
	;
	v225 = v215
	v226 = v216
	v227 = v217
	goto L74
L74:
	;
	v230 = *(*int32)(unsafe.Add(mBase, uint32(v225)))
	v231 = *(*int32)(unsafe.Add(mBase, uint32(v226)))
	if v230 != v231 {
		v248 = v225
		v249 = v226
		v250 = v227
		goto L70
	} else {
		goto L76
	}
L75:
	;
	goto L71
L76:
	;
	v233 = int32(4)
	v234 = v226 + v233
	v236 = v225 + v233
	v238 = v227 + int32(-4)
	if base.Ui32(int32(3)) < base.Ui32(v238) {
		v225 = v236
		v226 = v234
		v227 = v238
		goto L74
	} else {
		goto L77
	}
L77:
	;
	goto L75
L78:
	;
	v248 = v236
	v249 = v234
	v250 = v238
	goto L70
L79:
	;
	v281 = v258 - v259
	goto L68
L80:
	;
	v258 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v253))))
	v259 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v254))))
	if v258 != v259 {
		goto L79
	} else {
		goto L82
	}
L82:
	;
	v261 = int32(1)
	v266 = v255 + int32(-1)
	if v266 == int32(0) {
		goto L69
	} else {
		goto L83
	}
L83:
	;
	v253 = v253 + v261
	v254 = v254 + v261
	v255 = v266
	goto L80
L84:
	;
	goto L17
L85:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L86:
	;
	v373 = int32(0)
	v376 = v373
	v377 = v373
	v378 = v306
	v379 = v372
	v380 = int32(0)
	v381 = v373
	goto L19
L87:
	;
	v372 = int32(0)
	goto L86
L88:
	;
	v344 = v339
	v345 = v340
	v346 = v341
	goto L98
L89:
	;
	if v329 == int32(0) {
		goto L87
	} else {
		goto L96
	}
L90:
	;
	if (v307|v306)&int32(3) != 0 {
		v339 = v306
		v340 = v307
		v341 = v308
		goto L88
	} else {
		goto L91
	}
L91:
	;
	v316 = v306
	v317 = v307
	v318 = v308
	goto L92
L92:
	;
	v321 = *(*int32)(unsafe.Add(mBase, uint32(v316)))
	v322 = *(*int32)(unsafe.Add(mBase, uint32(v317)))
	if v321 != v322 {
		v339 = v316
		v340 = v317
		v341 = v318
		goto L88
	} else {
		goto L94
	}
L93:
	;
	goto L89
L94:
	;
	v324 = int32(4)
	v325 = v317 + v324
	v327 = v316 + v324
	v329 = v318 + int32(-4)
	if base.Ui32(int32(3)) < base.Ui32(v329) {
		v316 = v327
		v317 = v325
		v318 = v329
		goto L92
	} else {
		goto L95
	}
L95:
	;
	goto L93
L96:
	;
	v339 = v327
	v340 = v325
	v341 = v329
	goto L88
L97:
	;
	v372 = v349 - v350
	goto L86
L98:
	;
	v349 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v344))))
	v350 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v345))))
	if v349 != v350 {
		goto L97
	} else {
		goto L100
	}
L100:
	;
	v352 = int32(1)
	v357 = v346 + int32(-1)
	if v357 == int32(0) {
		goto L87
	} else {
		goto L101
	}
L101:
	;
	v344 = v344 + v352
	v345 = v345 + v352
	v346 = v357
	goto L98
L102:
	;
	v391 = int64(0)
	if v377 != 0 {
		goto L105
	} else {
		goto L106
	}
L103:
	;
	v387 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+52)) = v387 | int32(1)
	goto L102
L104:
	;
	switch v71 & int32(65535) {
	case 0, 2:
		goto L131
	default:
		goto L130
	}
L105:
	;
	v393 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+52)))
	if v393&int32(1) == int32(0) {
		goto L107
	} else {
		goto L108
	}
L106:
	;
	v617 = v391
	v618 = int64(0)
	goto L104
L107:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v376)+2200)) = v66
	v403 = *(*int32)(unsafe.Add(mBase, uint32(v376)+88))
	*(*int32)(unsafe.Add(mBase, uint32(v376)+88)) = v403&int32(-14337) | v149&int32(14336)
	if v403&int32(32) != 0 {
		v617 = v391
		v618 = int64(0)
		goto L104
	} else {
		goto L109
	}
L108:
	;
	v398 = *(*int32)(unsafe.Add(mBase, uint32(v376)+88))
	*(*int32)(unsafe.Add(mBase, uint32(v376)+88)) = v398 | int32(1024)
	goto L107
L109:
	;
	v413 = *(*int64)(unsafe.Add(mBase, uint32(v142)+24))
	v414 = int64(8)
	v416 = int64(4278190080)
	v418 = int64(24)
	v420 = int64(16711680)
	v423 = int64(40)
	v425 = int64(65280)
	v427 = int64(56)
	v448 = *(*int64)(unsafe.Add(mBase, uint32(v142)+16))
	v483 = v448<<(uint(v427)%64) | v448&v425<<(uint(v423)%64) | (v448&v420<<(uint(v418)%64) | v448&v416<<(uint(v414)%64)) | (int64(base.Ui64(v448)>>(uint(v414)%64))&v416 | int64(base.Ui64(v448)>>(uint(v418)%64))&v420 | (int64(base.Ui64(v448)>>(uint(v423)%64))&v425 | int64(base.Ui64(v448)>>(uint(v427)%64))))
	v485 = *(*int32)(unsafe.Add(mBase, _consts[90]))
	v486 = *(*int64)(unsafe.Add(mBase, uint32(v485)+8))
	if base.Ui64(v483) <= base.Ui64(v486) {
		goto L110
	} else {
		goto L111
	}
L110:
	;
	v489 = v413<<(uint(v427)%64) | v413&v425<<(uint(v423)%64) | (v413&v420<<(uint(v418)%64) | v413&v416<<(uint(v414)%64)) | (int64(base.Ui64(v413)>>(uint(v414)%64))&v416 | int64(base.Ui64(v413)>>(uint(v418)%64))&v420 | (int64(base.Ui64(v413)>>(uint(v423)%64))&v425 | int64(base.Ui64(v413)>>(uint(v427)%64))))
	if v379 != 0 {
		v541 = v485
		goto L112
	} else {
		goto L113
	}
L111:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v485)+8)) = v483
	goto L110
L112:
	;
	v544 = *(*int64)(unsafe.Add(mBase, uint32(v142)+32))
	v545 = int64(56)
	v547 = int64(65280)
	v549 = int64(40)
	v552 = int64(16711680)
	v554 = int64(24)
	v556 = int64(4278190080)
	v558 = int64(8)
	v579 = v544<<(uint(v545)%64) | v544&v547<<(uint(v549)%64) | (v544&v552<<(uint(v554)%64) | v544&v556<<(uint(v558)%64)) | (int64(base.Ui64(v544)>>(uint(v558)%64))&v556 | int64(base.Ui64(v544)>>(uint(v554)%64))&v552 | (int64(base.Ui64(v544)>>(uint(v549)%64))&v547 | int64(base.Ui64(v544)>>(uint(v545)%64))))
	*(*int64)(unsafe.Add(mBase, uint32(v376)+2248)) = v579
	v581 = *(*int64)(unsafe.Add(mBase, uint32(v541)+uint32(_consts[113])))
	if v581 == int64(0) {
		v617 = v489
		v618 = v483
		goto L104
	} else {
		goto L123
	}
L113:
	;
	v490 = *(*int64)(unsafe.Add(mBase, uint32(v376)+96))
	if base.Ui64(v489) <= base.Ui64(v490) {
		v541 = v485
		goto L112
	} else {
		goto L114
	}
L114:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v376)+96)) = v489
	F_clearCachedClusterSlotsResponse(m)
	mBase = m.M
	v494 = m.ExcPending
	if v494 != 0 {
		goto L1
	} else {
		goto L115
	}
L115:
	;
	v496 = *(*int32)(unsafe.Add(mBase, _consts[90]))
	v497 = *(*int32)(unsafe.Add(mBase, uint32(v496)+uint32(_consts[101])))
	v499 = v497 | int32(12)
	*(*int32)(unsafe.Add(mBase, uint32(v496)+uint32(_consts[101]))) = v499
	v501 = *(*int64)(unsafe.Add(mBase, uint32(v496)+uint32(_consts[121])))
	if v501 == int64(0) {
		v541 = v496
		goto L112
	} else {
		goto L116
	}
L116:
	;
	v504 = *(*int32)(unsafe.Add(mBase, uint32(v496)+uint32(_consts[126])))
	if v504 == int32(0) {
		v541 = v496
		goto L112
	} else {
		goto L117
	}
L117:
	;
	v507 = *(*int64)(unsafe.Add(mBase, uint32(v376)+96))
	v508 = *(*int64)(unsafe.Add(mBase, uint32(v496)+uint32(_consts[127])))
	if base.Ui64(v507) < base.Ui64(v508) {
		v541 = v496
		goto L112
	} else {
		goto L118
	}
L118:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v496)+uint32(_consts[121]))) = int64(0)
	v513 = *(*int32)(unsafe.Add(mBase, _consts[15]))
	if int32(3) < v513 {
		v536 = v496
		v537 = v499
		goto L119
	} else {
		goto L120
	}
L119:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v536)+uint32(_consts[101]))) = v537 | int32(1)
	v541 = v536
	goto L112
L120:
	;
	v516 = F_humanNodename(m, v376)
	mBase = m.M
	v517 = m.ExcPending
	if v517 != 0 {
		goto L1
	} else {
		goto L121
	}
L121:
	;
	v520 = *(*int64)(unsafe.Add(mBase, uint32(v376)+96))
	*(*int64)(unsafe.Add(mBase, uint32(v24+int32(576)))) = v520
	*(*int32)(unsafe.Add(mBase, uint32(v24)+572)) = v516
	*(*int32)(unsafe.Add(mBase, uint32(v24)+568)) = v376 + int32(8)
	*(*int64)(unsafe.Add(mBase, uint32(v24)+560)) = v508
	F__serverLog(m, int32(3), int32(_a248), v24+int32(560))
	mBase = m.M
	v532 = m.ExcPending
	if v532 != 0 {
		goto L1
	} else {
		goto L122
	}
L122:
	;
	v534 = *(*int32)(unsafe.Add(mBase, _consts[90]))
	v535 = *(*int32)(unsafe.Add(mBase, uint32(v534)+uint32(_consts[101])))
	v536 = v534
	v537 = v535
	goto L119
L123:
	;
	v584 = int32(0)
	v585 = *(*int32)(unsafe.Add(mBase, _consts[94]))
	v586 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v585)+88)))
	if v586&int32(2) == v584 {
		v617 = v489
		v618 = v483
		goto L104
	} else {
		goto L124
	}
L124:
	;
	v591 = *(*int32)(unsafe.Add(mBase, uint32(v585)+2172))
	if v591 != v376 {
		v617 = v489
		v618 = v483
		goto L104
	} else {
		goto L125
	}
L125:
	;
	v593 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v142)+2253)))
	if v593&int32(1) == int32(0) {
		v617 = v489
		v618 = v483
		goto L104
	} else {
		goto L126
	}
L126:
	;
	v598 = *(*int64)(unsafe.Add(mBase, uint32(v541)+uint32(_consts[114])))
	if v598 != int64(-1) {
		v617 = v489
		v618 = v483
		goto L104
	} else {
		goto L127
	}
L127:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v541)+uint32(_consts[114]))) = v579
	v602 = *(*int32)(unsafe.Add(mBase, uint32(v541)+uint32(_consts[101])))
	*(*int32)(unsafe.Add(mBase, uint32(v541)+uint32(_consts[101]))) = v602 | int32(16)
	v607 = *(*int32)(unsafe.Add(mBase, _consts[15]))
	if int32(2) < v607 {
		v617 = v489
		v618 = v483
		goto L104
	} else {
		goto L128
	}
L128:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v24)+544)) = v579
	F__serverLog(m, int32(2), int32(_a249), v24+int32(544))
	mBase = m.M
	v616 = m.ExcPending
	if v616 != 0 {
		goto L1
	} else {
		goto L129
	}
L129:
	;
	v617 = v489
	v618 = v483
	goto L104
L130:
	;
	v1017 = *(*int32)(unsafe.Add(mBase, _consts[15]))
	if int32(0) < v1017 {
		goto L215
	} else {
		goto L216
	}
L131:
	;
	if v70 == int32(2) {
		goto L135
	} else {
		goto L136
	}
L132:
	;
	F_clusterSendPing(m, l0, int32(1))
	mBase = m.M
	v1010 = m.ExcPending
	if v1010 != 0 {
		goto L1
	} else {
		goto L214
	}
L133:
	;
	if v377 != 0 {
		goto L162
	} else {
		goto L163
	}
L134:
	;
	v636 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v636 == int32(0) {
		goto L140
	} else {
		goto L141
	}
L135:
	;
	v635 = *(*int32)(unsafe.Add(mBase, _consts[128]))
	if v635 != 0 {
		goto L133
	} else {
		goto L139
	}
L136:
	;
	v628 = *(*int32)(unsafe.Add(mBase, _consts[94]))
	v629 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v628)+2256)))
	if v629 != 0 {
		goto L132
	} else {
		goto L137
	}
L137:
	;
	v631 = *(*int32)(unsafe.Add(mBase, _consts[128]))
	if v631 == int32(0) {
		goto L134
	} else {
		goto L138
	}
L138:
	;
	goto L132
L139:
	;
	goto L134
L140:
	;
	if v70 != int32(2) {
		goto L132
	} else {
		goto L158
	}
L141:
	;
	v639 = *(*int32)(unsafe.Add(mBase, uint32(v636)))
	v640 = *(*int32)(unsafe.Add(mBase, uint32(v639)+24))
	if v640 == int32(0) {
		goto L140
	} else {
		goto L142
	}
L142:
	;
	v646 = int32(0)
	v648 = m.T0[v640].(func(*base.Module, int32, int32, int32, int32, int32) int32)(m, v636, v24+int32(624), int32(46), v646, v646)
	mBase = m.M
	v649 = m.ExcPending
	if v649 != 0 {
		goto L1
	} else {
		goto L143
	}
L143:
	;
	if v648 == int32(-1) {
		goto L140
	} else {
		goto L144
	}
L144:
	;
	v653 = v24 + int32(624)
	v654 = int32(0)
	v655 = *(*int32)(unsafe.Add(mBase, _consts[94]))
	v657 = v655 + int32(2256)
	v660 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v657))))
	v661 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v653))))
	if v661 == v654 {
		v684 = v660
		v685 = v661
		goto L146
	} else {
		goto L147
	}
L145:
	;
	if v685-v684&int32(255) == int32(0) {
		goto L140
	} else {
		goto L153
	}
L146:
	;
	goto L145
L147:
	;
	if v661 != v660&int32(255) {
		v684 = v660
		v685 = v661
		goto L146
	} else {
		goto L148
	}
L148:
	;
	v667 = v653
	v668 = v657
	goto L149
L149:
	;
	v671 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v668)+1)))
	v672 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v667)+1)))
	if v672 == int32(0) {
		v684 = v671
		v685 = v672
		goto L146
	} else {
		goto L151
	}
L150:
	;
	v684 = v671
	v685 = v672
	goto L146
L151:
	;
	v675 = int32(1)
	if v672 == v671&int32(255) {
		v667 = v667 + v675
		v668 = v668 + v675
		goto L149
	} else {
		goto L152
	}
L152:
	;
	goto L150
L153:
	;
	v691 = *(*int64)(unsafe.Add(mBase, uint32(v24)+624))
	*(*int64)(unsafe.Add(mBase, uint32(v657))) = v691
	v695 = *(*int64)(unsafe.Add(mBase, uint32(v24)+632))
	*(*int64)(unsafe.Add(mBase, uint32(v655+int32(2264)))) = v695
	v703 = *(*int64)(unsafe.Add(mBase, uint32(v24+int32(662))))
	*(*int64)(unsafe.Add(mBase, uint32(v655+int32(2294)))) = v703
	v711 = *(*int64)(unsafe.Add(mBase, uint32(v24+int32(656))))
	*(*int64)(unsafe.Add(mBase, uint32(v655+int32(2288)))) = v711
	v719 = *(*int64)(unsafe.Add(mBase, uint32(v24+int32(648))))
	*(*int64)(unsafe.Add(mBase, uint32(v655+int32(2280)))) = v719
	v727 = *(*int64)(unsafe.Add(mBase, uint32(v24+int32(640))))
	*(*int64)(unsafe.Add(mBase, uint32(v655+int32(2272)))) = v727
	v730 = *(*int32)(unsafe.Add(mBase, _consts[15]))
	if int32(2) < v730 {
		goto L154
	} else {
		goto L155
	}
L154:
	;
	F_clearCachedClusterSlotsResponse(m)
	mBase = m.M
	v741 = m.ExcPending
	if v741 != 0 {
		goto L1
	} else {
		goto L157
	}
L155:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+528)) = v657
	F__serverLog(m, int32(2), int32(_a250), v24+int32(528))
	mBase = m.M
	v739 = m.ExcPending
	if v739 != 0 {
		goto L1
	} else {
		goto L156
	}
L156:
	;
	goto L154
L157:
	;
	v743 = *(*int32)(unsafe.Add(mBase, _consts[90]))
	v744 = *(*int32)(unsafe.Add(mBase, uint32(v743)+uint32(_consts[101])))
	*(*int32)(unsafe.Add(mBase, uint32(v743)+uint32(_consts[101]))) = v744 | int32(4)
	goto L140
L158:
	;
	goto L133
L159:
	;
	v921 = F_createClusterNode(m, int32(0), int32(32))
	mBase = m.M
	v922 = m.ExcPending
	if v922 != 0 {
		goto L1
	} else {
		goto L199
	}
L160:
	;
	F__serverAssert(m, int32(_a251), int32(_a179), int32(4067))
	mBase = m.M
	v915 = m.ExcPending
	if v915 != 0 {
		goto L1
	} else {
		goto L198
	}
L161:
	;
	F__serverAssert(m, int32(_a252), int32(_a179), int32(4049))
	mBase = m.M
	v909 = m.ExcPending
	if v909 != 0 {
		goto L1
	} else {
		goto L197
	}
L162:
	;
	v872 = *(*int32)(unsafe.Add(mBase, uint32(v376)+2344))
	if v872 == int32(0) {
		goto L132
	} else {
		goto L186
	}
L163:
	;
	v754 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	if v754 != 0 {
		goto L164
	} else {
		goto L165
	}
L164:
	;
	v859 = *(*int32)(unsafe.Add(mBase, _consts[122]))
	if v859 == int32(0) {
		goto L181
	} else {
		goto L182
	}
L165:
	;
	v757 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v24+int32(662)))) = v757
	v762 = v24 + int32(656)
	*(*int64)(unsafe.Add(mBase, uint32(v762))) = v757
	v768 = v24 + int32(648)
	*(*int64)(unsafe.Add(mBase, uint32(v768))) = v757
	v774 = v24 + int32(640)
	*(*int64)(unsafe.Add(mBase, uint32(v774))) = v757
	*(*int64)(unsafe.Add(mBase, uint32(v24)+632)) = v757
	*(*int64)(unsafe.Add(mBase, uint32(v24)+624)) = v757
	v781 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v142)+2168)))
	if v781 == int32(0) {
		goto L166
	} else {
		goto L167
	}
L166:
	;
	v814 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v814 == int32(0) {
		goto L168
	} else {
		goto L169
	}
L167:
	;
	v792 = *(*int64)(unsafe.Add(mBase, uint32(v142+int32(2205))))
	*(*int64)(unsafe.Add(mBase, uint32(v24+int32(661)))) = v792
	v796 = *(*int64)(unsafe.Add(mBase, uint32(v142+int32(2200))))
	*(*int64)(unsafe.Add(mBase, uint32(v762))) = v796
	v800 = *(*int64)(unsafe.Add(mBase, uint32(v142+int32(2192))))
	*(*int64)(unsafe.Add(mBase, uint32(v768))) = v800
	v804 = *(*int64)(unsafe.Add(mBase, uint32(v142+int32(2184))))
	*(*int64)(unsafe.Add(mBase, uint32(v774))) = v804
	v806 = *(*int64)(unsafe.Add(mBase, uint32(v142+int32(2168))))
	*(*int64)(unsafe.Add(mBase, uint32(v24)+624)) = v806
	v810 = *(*int64)(unsafe.Add(mBase, uint32(v142+int32(2176))))
	*(*int64)(unsafe.Add(mBase, uint32(v24)+632)) = v810
	v812 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v24)+669)) = uint8(v812)
	goto L159
L168:
	;
	v832 = *(*int32)(unsafe.Add(mBase, _consts[15]))
	if int32(2) < v832 {
		goto L17
	} else {
		goto L173
	}
L169:
	;
	v817 = *(*int32)(unsafe.Add(mBase, uint32(v814)))
	v818 = *(*int32)(unsafe.Add(mBase, uint32(v817)+24))
	if v818 == int32(0) {
		goto L168
	} else {
		goto L170
	}
L170:
	;
	v826 = m.T0[v818].(func(*base.Module, int32, int32, int32, int32, int32) int32)(m, v814, v24+int32(624), int32(46), int32(0), int32(1))
	mBase = m.M
	v827 = m.ExcPending
	if v827 != 0 {
		goto L1
	} else {
		goto L171
	}
L171:
	;
	if v826 != int32(-1) {
		goto L159
	} else {
		goto L172
	}
L172:
	;
	goto L168
L173:
	;
	v835 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v835 != 0 {
		goto L175
	} else {
		goto L176
	}
L174:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+512)) = v841
	F__serverLog(m, int32(2), int32(_a253), v24+int32(512))
	mBase = m.M
	v848 = m.ExcPending
	if v848 != 0 {
		goto L1
	} else {
		goto L178
	}
L175:
	;
	v837 = *(*int32)(unsafe.Add(mBase, uint32(v835)))
	v838 = *(*int32)(unsafe.Add(mBase, uint32(v837)+88))
	v839 = m.T0[v838].(func(*base.Module, int32) int32)(m, v835)
	mBase = m.M
	v840 = m.ExcPending
	if v840 != 0 {
		goto L1
	} else {
		goto L177
	}
L176:
	;
	v841 = int32(_a254)
	goto L174
L177:
	;
	v841 = v839
	goto L174
L178:
	;
	v850 = *(*int32)(unsafe.Add(mBase, _consts[15]))
	if int32(2) < v850 {
		goto L17
	} else {
		goto L179
	}
L179:
	;
	F__serverLog(m, int32(2), int32(_a255), int32(0))
	mBase = m.M
	v857 = m.ExcPending
	if v857 != 0 {
		goto L1
	} else {
		goto L180
	}
L180:
	;
	goto L17
L181:
	;
	F_clusterProcessGossipSection(m, v142, l0)
	mBase = m.M
	v871 = m.ExcPending
	if v871 != 0 {
		goto L1
	} else {
		goto L185
	}
L182:
	;
	v862 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	if v862 == int32(0) {
		goto L161
	} else {
		goto L183
	}
L183:
	;
	v865 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v754)+88)))
	if v865&int32(32) == int32(0) {
		goto L161
	} else {
		goto L184
	}
L184:
	;
	goto L181
L185:
	;
	goto L132
L186:
	;
	v875 = *(*int64)(unsafe.Add(mBase, uint32(v376)))
	v878 = *(*int64)(unsafe.Add(mBase, _consts[117]))
	v879 = int64(1000)
	if v879 < v878 {
		goto L187
	} else {
		goto L188
	}
L187:
	;
	v882 = v878
	goto L189
L188:
	;
	v882 = v879
	goto L189
L189:
	;
	if v66-v875 <= v882 {
		goto L132
	} else {
		goto L190
	}
L190:
	;
	if l0 == v872 {
		goto L160
	} else {
		goto L191
	}
L191:
	;
	v886 = *(*int32)(unsafe.Add(mBase, _consts[15]))
	if int32(2) < v886 {
		v902 = v872
		goto L192
	} else {
		goto L193
	}
L192:
	;
	F_freeClusterLink(m, v902)
	mBase = m.M
	v904 = m.ExcPending
	if v904 != 0 {
		goto L1
	} else {
		goto L196
	}
L193:
	;
	v889 = F_humanNodename(m, v376)
	mBase = m.M
	v890 = m.ExcPending
	if v890 != 0 {
		goto L1
	} else {
		goto L194
	}
L194:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+500)) = v889
	*(*int32)(unsafe.Add(mBase, uint32(v24)+496)) = v376 + int32(8)
	F__serverLog(m, int32(2), int32(_a256), v24+int32(496))
	mBase = m.M
	v900 = m.ExcPending
	if v900 != 0 {
		goto L1
	} else {
		goto L195
	}
L195:
	;
	v901 = *(*int32)(unsafe.Add(mBase, uint32(v376)+2344))
	v902 = v901
	goto L192
L196:
	;
	goto L132
L197:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L198:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L199:
	;
	v927 = *(*int64)(unsafe.Add(mBase, uint32(v24+int32(662))))
	*(*int64)(unsafe.Add(mBase, uint32(v921+int32(2294)))) = v927
	v935 = *(*int64)(unsafe.Add(mBase, uint32(v24+int32(656))))
	*(*int64)(unsafe.Add(mBase, uint32(v921+int32(2288)))) = v935
	v941 = *(*int64)(unsafe.Add(mBase, uint32(v24+int32(648))))
	*(*int64)(unsafe.Add(mBase, uint32(v921+int32(2280)))) = v941
	v947 = *(*int64)(unsafe.Add(mBase, uint32(v24+int32(640))))
	*(*int64)(unsafe.Add(mBase, uint32(v921+int32(2272)))) = v947
	v951 = *(*int64)(unsafe.Add(mBase, uint32(v24)+632))
	*(*int64)(unsafe.Add(mBase, uint32(v921+int32(2264)))) = v951
	v953 = *(*int64)(unsafe.Add(mBase, uint32(v24)+624))
	*(*int64)(unsafe.Add(mBase, uint32(v921)+2256)) = v953
	v964 = *(*int32)(unsafe.Add(mBase, _consts[89]))
	if v964 != 0 {
		goto L201
	} else {
		goto L202
	}
L200:
	;
	v977 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v142)+2248)))
	v978 = F___bswap_16_2(m, v977)
	mBase = m.M
	goto L207
L201:
	;
	v965 = int32(2246)
	goto L203
L202:
	;
	v965 = int32(10)
	goto L203
L203:
	;
	v967 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v142+v965))))
	if v964 != 0 {
		goto L204
	} else {
		goto L205
	}
L204:
	;
	v970 = int32(10)
	goto L206
L205:
	;
	v970 = int32(2246)
	goto L206
L206:
	;
	v972 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v142+v970))))
	v973 = F_ntohs(m, v972)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v921+int32(2328)))) = v973
	v975 = F_ntohs(m, v967)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v921+int32(2324)))) = v975
	goto L200
L207:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v921)+2332)) = v978
	v980 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+52)))
	if v980&int32(1) == int32(0) {
		goto L208
	} else {
		goto L209
	}
L208:
	;
	F_setClusterNodeToInboundClusterLink(m, v921, l0)
	mBase = m.M
	v990 = m.ExcPending
	if v990 != 0 {
		goto L1
	} else {
		goto L210
	}
L209:
	;
	v985 = *(*int32)(unsafe.Add(mBase, uint32(v921)+88))
	*(*int32)(unsafe.Add(mBase, uint32(v921)+88)) = v985 | int32(1024)
	goto L208
L210:
	;
	F_clusterAddNode(m, v921)
	mBase = m.M
	v992 = m.ExcPending
	if v992 != 0 {
		goto L1
	} else {
		goto L211
	}
L211:
	;
	F_clearCachedClusterSlotsResponse(m)
	mBase = m.M
	v994 = m.ExcPending
	if v994 != 0 {
		goto L1
	} else {
		goto L212
	}
L212:
	;
	v996 = *(*int32)(unsafe.Add(mBase, _consts[90]))
	v997 = *(*int32)(unsafe.Add(mBase, uint32(v996)+uint32(_consts[101])))
	*(*int32)(unsafe.Add(mBase, uint32(v996)+uint32(_consts[101]))) = v997 | int32(4)
	F_clusterProcessGossipSection(m, v142, l0)
	mBase = m.M
	v1002 = m.ExcPending
	if v1002 != 0 {
		goto L1
	} else {
		goto L213
	}
L213:
	;
	goto L132
L214:
	;
	goto L130
L215:
	;
	if base.Ui32(int32(2)) < base.Ui32(v70) {
		goto L243
	} else {
		goto L244
	}
L216:
	;
	v1020 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v1020 == int32(0) {
		goto L217
	} else {
		goto L218
	}
L217:
	;
	v1081 = int32(_a246)
	v1082 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	if v1082 != 0 {
		goto L234
	} else {
		goto L235
	}
L218:
	;
	v1023 = *(*int32)(unsafe.Add(mBase, uint32(v1020)))
	v1024 = *(*int32)(unsafe.Add(mBase, uint32(v1023)+24))
	if v1024 == int32(0) {
		goto L217
	} else {
		goto L219
	}
L219:
	;
	v1033 = m.T0[v1024].(func(*base.Module, int32, int32, int32, int32, int32) int32)(m, v1020, v24+int32(624), int32(46), v24+int32(620), int32(1))
	mBase = m.M
	v1034 = m.ExcPending
	if v1034 != 0 {
		goto L1
	} else {
		goto L220
	}
L220:
	;
	v1036 = *(*int32)(unsafe.Add(mBase, _consts[15]))
	if v1033 == int32(-1) {
		goto L221
	} else {
		goto L222
	}
L221:
	;
	if int32(0) < v1036 {
		goto L215
	} else {
		goto L233
	}
L222:
	;
	if int32(0) < v1036 {
		goto L215
	} else {
		goto L223
	}
L223:
	;
	if base.Ui32(int32(10)) < base.Ui32(v70) {
		v1049 = int32(_a242)
		goto L224
	} else {
		goto L225
	}
L224:
	;
	v1050 = int32(_a246)
	v1051 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	if v1051 != 0 {
		goto L226
	} else {
		goto L227
	}
L225:
	;
	v1048 = *(*int32)(unsafe.Add(mBase, uint32(v70<<(uint(int32(2))%32))+uint32(_consts[125])))
	v1049 = v1048
	goto L224
L226:
	;
	v1055 = v1051 + int32(8)
	goto L228
L227:
	;
	v1055 = v1050
	goto L228
L228:
	;
	if v1051 == int32(0) {
		v1060 = v1050
		goto L229
	} else {
		goto L230
	}
L229:
	;
	v1063 = *(*int32)(unsafe.Add(mBase, uint32(v24)+620))
	*(*int32)(unsafe.Add(mBase, uint32(v24+int32(480)))) = v1063
	*(*int32)(unsafe.Add(mBase, uint32(v24)+464)) = v1049
	*(*int32)(unsafe.Add(mBase, uint32(v24)+468)) = v1055
	*(*int32)(unsafe.Add(mBase, uint32(v24)+472)) = v1060
	*(*int32)(unsafe.Add(mBase, uint32(v24)+476)) = v24 + int32(624)
	F__serverLog(m, int32(0), int32(_a257), v24+int32(464))
	mBase = m.M
	v1076 = m.ExcPending
	if v1076 != 0 {
		goto L1
	} else {
		goto L232
	}
L230:
	;
	v1058 = F_humanNodename(m, v1051)
	mBase = m.M
	v1059 = m.ExcPending
	if v1059 != 0 {
		goto L1
	} else {
		goto L231
	}
L231:
	;
	v1060 = v1058
	goto L229
L232:
	;
	goto L215
L233:
	;
	goto L217
L234:
	;
	v1086 = v1082 + int32(8)
	goto L236
L235:
	;
	v1086 = v1081
	goto L236
L236:
	;
	if v1082 == int32(0) {
		v1091 = v1081
		goto L237
	} else {
		goto L238
	}
L237:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+452)) = v1091
	*(*int32)(unsafe.Add(mBase, uint32(v24)+448)) = v1086
	F__serverLog(m, int32(0), int32(_a258), v24+int32(448))
	mBase = m.M
	v1099 = m.ExcPending
	if v1099 != 0 {
		goto L1
	} else {
		goto L240
	}
L238:
	;
	v1089 = F_humanNodename(m, v1082)
	mBase = m.M
	v1090 = m.ExcPending
	if v1090 != 0 {
		goto L1
	} else {
		goto L239
	}
L239:
	;
	v1091 = v1089
	goto L237
L240:
	;
	goto L215
L241:
	;
	if v379 != 0 {
		goto L392
	} else {
		goto L393
	}
L242:
	;
	v1698 = int32(1)
	if v70 != v1698 {
		goto L378
	} else {
		goto L379
	}
L243:
	;
	if v70 != int32(3) {
		goto L308
	} else {
		goto L309
	}
L244:
	;
	if v377 == int32(0) {
		goto L245
	} else {
		goto L246
	}
L245:
	;
	v1133 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	if v1133 != 0 {
		goto L251
	} else {
		goto L252
	}
L246:
	;
	v1108 = *(*int32)(unsafe.Add(mBase, uint32(v376)+88))
	if v1108&int32(128) == int32(0) {
		goto L245
	} else {
		goto L247
	}
L247:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v376)+88)) = v1108 & int32(-129)
	v1117 = *(*int32)(unsafe.Add(mBase, _consts[15]))
	if int32(2) < v1117 {
		goto L245
	} else {
		goto L248
	}
L248:
	;
	v1120 = F_humanNodename(m, v376)
	mBase = m.M
	v1121 = m.ExcPending
	if v1121 != 0 {
		goto L1
	} else {
		goto L249
	}
L249:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+324)) = v1120
	*(*int32)(unsafe.Add(mBase, uint32(v24)+320)) = v376 + int32(8)
	F__serverLog(m, int32(2), int32(_a259), v24+int32(320))
	mBase = m.M
	v1131 = m.ExcPending
	if v1131 != 0 {
		goto L1
	} else {
		goto L250
	}
L250:
	;
	goto L245
L251:
	;
	if v377 == int32(0) {
		goto L242
	} else {
		goto L299
	}
L252:
	;
	v1134 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	v1135 = *(*int32)(unsafe.Add(mBase, uint32(v1134)+88))
	if v1135&int32(32) == int32(0) {
		goto L253
	} else {
		goto L254
	}
L253:
	;
	v1214 = v1134 + int32(8)
	v1215 = int32(40)
	v1216 = v142 + v1215
	goto L276
L254:
	;
	if v377 == int32(0) {
		goto L255
	} else {
		goto L256
	}
L255:
	;
	F_clusterRenameNode(m, v1134, v142+int32(40))
	mBase = m.M
	v1178 = m.ExcPending
	if v1178 != 0 {
		goto L1
	} else {
		goto L266
	}
L256:
	;
	v1143 = *(*int32)(unsafe.Add(mBase, _consts[15]))
	if int32(1) < v1143 {
		goto L257
	} else {
		goto L258
	}
L257:
	;
	v1158 = F_nodeUpdateAddressIfNeeded(m, v376, l0, v142)
	mBase = m.M
	v1159 = m.ExcPending
	if v1159 != 0 {
		goto L1
	} else {
		goto L262
	}
L258:
	;
	v1146 = F_humanNodename(m, v376)
	mBase = m.M
	v1147 = m.ExcPending
	if v1147 != 0 {
		goto L1
	} else {
		goto L259
	}
L259:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+292)) = v1146
	*(*int32)(unsafe.Add(mBase, uint32(v24)+288)) = v376 + int32(8)
	F__serverLog(m, int32(1), int32(_a260), v24+int32(288))
	mBase = m.M
	v1157 = m.ExcPending
	if v1157 != 0 {
		goto L1
	} else {
		goto L260
	}
L260:
	;
	goto L257
L261:
	;
	v1171 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	F_clusterDelNode(m, v1171)
	mBase = m.M
	v1173 = m.ExcPending
	if v1173 != 0 {
		goto L1
	} else {
		goto L265
	}
L262:
	;
	if v1158 == int32(0) {
		goto L261
	} else {
		goto L263
	}
L263:
	;
	F_clearCachedClusterSlotsResponse(m)
	mBase = m.M
	v1163 = m.ExcPending
	if v1163 != 0 {
		goto L1
	} else {
		goto L264
	}
L264:
	;
	v1165 = *(*int32)(unsafe.Add(mBase, _consts[90]))
	v1166 = *(*int32)(unsafe.Add(mBase, uint32(v1165)+uint32(_consts[101])))
	*(*int32)(unsafe.Add(mBase, uint32(v1165)+uint32(_consts[101]))) = v1166 | int32(6)
	goto L261
L265:
	;
	v2702 = int32(0)
	goto L3
L266:
	;
	v1180 = *(*int32)(unsafe.Add(mBase, _consts[15]))
	if int32(0) < v1180 {
		goto L267
	} else {
		goto L268
	}
L267:
	;
	v1197 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	v1198 = *(*int32)(unsafe.Add(mBase, uint32(v1197)+88))
	*(*int32)(unsafe.Add(mBase, uint32(v1197)+88)) = v1198&int32(-33) | v149&int32(3)
	F_clearCachedClusterSlotsResponse(m)
	mBase = m.M
	v1206 = m.ExcPending
	if v1206 != 0 {
		goto L1
	} else {
		goto L271
	}
L268:
	;
	v1183 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	v1184 = F_humanNodename(m, v1183)
	mBase = m.M
	v1185 = m.ExcPending
	if v1185 != 0 {
		goto L1
	} else {
		goto L269
	}
L269:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+308)) = v1184
	*(*int32)(unsafe.Add(mBase, uint32(v24)+304)) = v1183 + int32(8)
	F__serverLog(m, int32(0), int32(_a261), v24+int32(304))
	mBase = m.M
	v1195 = m.ExcPending
	if v1195 != 0 {
		goto L1
	} else {
		goto L270
	}
L270:
	;
	goto L267
L271:
	;
	v1208 = *(*int32)(unsafe.Add(mBase, _consts[90]))
	v1209 = *(*int32)(unsafe.Add(mBase, uint32(v1208)+uint32(_consts[101])))
	*(*int32)(unsafe.Add(mBase, uint32(v1208)+uint32(_consts[101]))) = v1209 | int32(4)
	goto L242
L272:
	;
	if v1281 == int32(0) {
		goto L251
	} else {
		goto L288
	}
L273:
	;
	v1281 = int32(0)
	goto L272
L274:
	;
	v1253 = v1248
	v1254 = v1249
	v1255 = v1250
	goto L284
L275:
	;
	if v1238 == int32(0) {
		goto L273
	} else {
		goto L282
	}
L276:
	;
	if (v1216|v1214)&int32(3) != 0 {
		v1248 = v1214
		v1249 = v1216
		v1250 = v1215
		goto L274
	} else {
		goto L277
	}
L277:
	;
	v1225 = v1214
	v1226 = v1216
	v1227 = v1215
	goto L278
L278:
	;
	v1230 = *(*int32)(unsafe.Add(mBase, uint32(v1225)))
	v1231 = *(*int32)(unsafe.Add(mBase, uint32(v1226)))
	if v1230 != v1231 {
		v1248 = v1225
		v1249 = v1226
		v1250 = v1227
		goto L274
	} else {
		goto L280
	}
L279:
	;
	goto L275
L280:
	;
	v1233 = int32(4)
	v1234 = v1226 + v1233
	v1236 = v1225 + v1233
	v1238 = v1227 + int32(-4)
	if base.Ui32(int32(3)) < base.Ui32(v1238) {
		v1225 = v1236
		v1226 = v1234
		v1227 = v1238
		goto L278
	} else {
		goto L281
	}
L281:
	;
	goto L279
L282:
	;
	v1248 = v1236
	v1249 = v1234
	v1250 = v1238
	goto L274
L283:
	;
	v1281 = v1258 - v1259
	goto L272
L284:
	;
	v1258 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1253))))
	v1259 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1254))))
	if v1258 != v1259 {
		goto L283
	} else {
		goto L286
	}
L286:
	;
	v1261 = int32(1)
	v1266 = v1255 + int32(-1)
	if v1266 == int32(0) {
		goto L273
	} else {
		goto L287
	}
L287:
	;
	v1253 = v1253 + v1261
	v1254 = v1254 + v1261
	v1255 = v1266
	goto L284
L288:
	;
	v1285 = *(*int32)(unsafe.Add(mBase, _consts[15]))
	if int32(2) < v1285 {
		v1311 = v1134
		v1314 = v1135
		goto L289
	} else {
		goto L290
	}
L289:
	;
	v1315 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v1311)+2256)) = uint8(v1315)
	*(*int32)(unsafe.Add(mBase, uint32(v1311)+88)) = v1314 | int32(64)
	v1321 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	*(*int32)(unsafe.Add(mBase, uint32(v1321)+2332)) = v1315
	*(*int64)(unsafe.Add(mBase, uint32(v1321)+2324)) = int64(0)
	F_freeClusterLink(m, l0)
	mBase = m.M
	v1327 = m.ExcPending
	if v1327 != 0 {
		goto L1
	} else {
		goto L293
	}
L290:
	;
	v1288 = F_humanNodename(m, v1134)
	mBase = m.M
	v1289 = m.ExcPending
	if v1289 != 0 {
		goto L1
	} else {
		goto L291
	}
L291:
	;
	v1290 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	v1291 = *(*int64)(unsafe.Add(mBase, uint32(v1290)))
	v1294 = *(*int32)(unsafe.Add(mBase, uint32(v1290)+88))
	*(*int32)(unsafe.Add(mBase, uint32(v24+int32(272)))) = v1294
	*(*int32)(unsafe.Add(mBase, uint32(v24)+264)) = v1290 + int32(48)
	*(*int32)(unsafe.Add(mBase, uint32(v24)+260)) = v1288
	*(*int32)(unsafe.Add(mBase, uint32(v24)+256)) = v1214
	v1301 = v66 - v1291
	*(*uint32)(unsafe.Add(mBase, uint32(v24)+268)) = uint32(v1301)
	F__serverLog(m, int32(2), int32(_a262), v24+int32(256))
	mBase = m.M
	v1308 = m.ExcPending
	if v1308 != 0 {
		goto L1
	} else {
		goto L292
	}
L292:
	;
	v1309 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	v1310 = *(*int32)(unsafe.Add(mBase, uint32(v1309)+88))
	v1311 = v1309
	v1314 = v1310
	goto L289
L293:
	;
	v1328 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1134)+88)))
	if v1328&int32(8) != 0 {
		goto L294
	} else {
		goto L295
	}
L294:
	;
	F_clearCachedClusterSlotsResponse(m)
	mBase = m.M
	v1336 = m.ExcPending
	if v1336 != 0 {
		goto L1
	} else {
		goto L298
	}
L295:
	;
	F_markNodeAsFailing(m, v1134)
	mBase = m.M
	v1332 = m.ExcPending
	if v1332 != 0 {
		goto L1
	} else {
		goto L296
	}
L296:
	;
	F_clusterSendFail(m, v1214)
	mBase = m.M
	v1334 = m.ExcPending
	if v1334 != 0 {
		goto L1
	} else {
		goto L297
	}
L297:
	;
	goto L294
L298:
	;
	v1338 = *(*int32)(unsafe.Add(mBase, _consts[90]))
	v1339 = *(*int32)(unsafe.Add(mBase, uint32(v1338)+uint32(_consts[101])))
	*(*int32)(unsafe.Add(mBase, uint32(v1338)+uint32(_consts[101]))) = v1339 | int32(6)
	v2702 = v1315
	goto L3
L299:
	;
	v1348 = *(*int32)(unsafe.Add(mBase, uint32(v376)+88))
	v1349 = int32(512)
	v1352 = v149 & v1349
	if v1348&v1349 == v1352 {
		v1363 = v1348
		goto L300
	} else {
		goto L301
	}
L300:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v376)+88)) = v1363&int32(-513) | v1352
	if v70 != 0 {
		goto L242
	} else {
		goto L303
	}
L301:
	;
	F_clearCachedClusterSlotsResponse(m)
	mBase = m.M
	v1355 = m.ExcPending
	if v1355 != 0 {
		goto L1
	} else {
		goto L302
	}
L302:
	;
	v1357 = *(*int32)(unsafe.Add(mBase, _consts[90]))
	v1358 = *(*int32)(unsafe.Add(mBase, uint32(v1357)+uint32(_consts[101])))
	*(*int32)(unsafe.Add(mBase, uint32(v1357)+uint32(_consts[101]))) = v1358 | int32(4)
	v1362 = *(*int32)(unsafe.Add(mBase, uint32(v376)+88))
	v1363 = v1362
	goto L300
L303:
	;
	if v1363&int32(32) != 0 {
		goto L241
	} else {
		goto L304
	}
L304:
	;
	v1370 = F_nodeUpdateAddressIfNeeded(m, v376, l0, v142)
	mBase = m.M
	v1371 = m.ExcPending
	if v1371 != 0 {
		goto L1
	} else {
		goto L305
	}
L305:
	;
	if v1370 == int32(0) {
		goto L241
	} else {
		goto L306
	}
L306:
	;
	F_clearCachedClusterSlotsResponse(m)
	mBase = m.M
	v1375 = m.ExcPending
	if v1375 != 0 {
		goto L1
	} else {
		goto L307
	}
L307:
	;
	v1377 = *(*int32)(unsafe.Add(mBase, _consts[90]))
	v1378 = *(*int32)(unsafe.Add(mBase, uint32(v1377)+uint32(_consts[101])))
	*(*int32)(unsafe.Add(mBase, uint32(v1377)+uint32(_consts[101]))) = v1378 | int32(6)
	goto L241
L308:
	;
	switch v71&int32(65535) + int32(-4) {
	case 0, 6:
		goto L337
	default:
		goto L336
	}
L309:
	;
	if v377 == int32(0) {
		goto L310
	} else {
		goto L311
	}
L310:
	;
	v1464 = int32(1)
	v1466 = *(*int32)(unsafe.Add(mBase, _consts[15]))
	if int32(2) < v1466 {
		v2702 = v1464
		goto L3
	} else {
		goto L334
	}
L311:
	;
	v1386 = int32(1)
	v1388 = v142 + int32(2256)
	goto L314
L312:
	;
	if int32(0)-v1412 != 0 {
		v2702 = v1386
		goto L3
	} else {
		goto L320
	}
L313:
	;
	goto L312
L314:
	;
	v1397 = int32(0)
	goto L316
L315:
	;
	goto L313
L316:
	;
	v1399 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1388+v1397))))
	v1402 = int32(255)
	v1412 = base.B2i32(base.Ui32((v1399+int32(-123))&v1402) < base.Ui32(int32(230))) & base.B2i32(base.Ui32((v1399+int32(-58))&v1402) < base.Ui32(int32(246)))
	if v1412 != 0 {
		goto L315
	} else {
		goto L318
	}
L317:
	;
	goto L315
L318:
	;
	v1414 = v1397 + int32(1)
	if v1414 != int32(40) {
		v1397 = v1414
		goto L316
	} else {
		goto L319
	}
L319:
	;
	goto L317
L320:
	;
	v1424 = F_sdsnewlen(m, v1388, int32(40))
	mBase = m.M
	v1425 = m.ExcPending
	if v1425 != 0 {
		goto L1
	} else {
		goto L321
	}
L321:
	;
	v1427 = *(*int32)(unsafe.Add(mBase, _consts[90]))
	v1428 = *(*int32)(unsafe.Add(mBase, uint32(v1427)+32))
	v1429 = F_dictFind(m, v1428, v1424)
	mBase = m.M
	v1430 = m.ExcPending
	if v1430 != 0 {
		goto L1
	} else {
		goto L322
	}
L322:
	;
	F_sdsfree(m, v1424)
	mBase = m.M
	v1432 = m.ExcPending
	if v1432 != 0 {
		goto L1
	} else {
		goto L323
	}
L323:
	;
	if v1429 == int32(0) {
		v2702 = v1386
		goto L3
	} else {
		goto L324
	}
L324:
	;
	v1435 = *(*int32)(unsafe.Add(mBase, uint32(v1429)+8))
	goto L325
L325:
	;
	if v1435 == int32(0) {
		v2702 = v1386
		goto L3
	} else {
		goto L326
	}
L326:
	;
	v1438 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1435)+88)))
	if v1438&int32(24) != 0 {
		v2702 = v1386
		goto L3
	} else {
		goto L327
	}
L327:
	;
	v1442 = *(*int32)(unsafe.Add(mBase, _consts[15]))
	if int32(2) < v1442 {
		goto L328
	} else {
		goto L329
	}
L328:
	;
	F_markNodeAsFailing(m, v1435)
	mBase = m.M
	v1463 = m.ExcPending
	if v1463 != 0 {
		goto L1
	} else {
		goto L333
	}
L329:
	;
	v1445 = F_humanNodename(m, v376)
	mBase = m.M
	v1446 = m.ExcPending
	if v1446 != 0 {
		goto L1
	} else {
		goto L330
	}
L330:
	;
	v1447 = F_humanNodename(m, v1435)
	mBase = m.M
	v1448 = m.ExcPending
	if v1448 != 0 {
		goto L1
	} else {
		goto L331
	}
L331:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+348)) = v1447
	*(*int32)(unsafe.Add(mBase, uint32(v24)+344)) = v1388
	*(*int32)(unsafe.Add(mBase, uint32(v24)+340)) = v1445
	*(*int32)(unsafe.Add(mBase, uint32(v24)+336)) = v142 + int32(40)
	F__serverLog(m, int32(2), int32(_a263), v24+int32(336))
	mBase = m.M
	v1460 = m.ExcPending
	if v1460 != 0 {
		goto L1
	} else {
		goto L332
	}
L332:
	;
	goto L328
L333:
	;
	v2702 = v1386
	goto L3
L334:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+356)) = v142 + int32(2256)
	*(*int32)(unsafe.Add(mBase, uint32(v24)+352)) = v142 + int32(40)
	F__serverLog(m, int32(2), int32(_a264), v24+int32(352))
	mBase = m.M
	v1480 = m.ExcPending
	if v1480 != 0 {
		goto L1
	} else {
		goto L335
	}
L335:
	;
	v2702 = v1464
	goto L3
L336:
	;
	switch v70 + int32(-5) {
	case 0:
		goto L345
	case 1:
		goto L344
	case 2:
		goto L342
	case 3:
		goto L343
	case 4:
		goto L341
	default:
		goto L340
	}
L337:
	;
	v1485 = int32(1)
	if v377 == int32(0) {
		v2702 = v1485
		goto L3
	} else {
		goto L338
	}
L338:
	;
	F_clusterProcessPublishPacket(m, v142+int32(2256), v70)
	mBase = m.M
	v1491 = m.ExcPending
	if v1491 != 0 {
		goto L1
	} else {
		goto L339
	}
L339:
	;
	v2702 = v1485
	goto L3
L340:
	;
	v1681 = int32(1)
	v1683 = *(*int32)(unsafe.Add(mBase, _consts[15]))
	if int32(3) < v1683 {
		v2702 = v1681
		goto L3
	} else {
		goto L376
	}
L341:
	;
	F_clusterProcessModulePacket(m, v142+int32(2256), v376)
	mBase = m.M
	v1679 = m.ExcPending
	if v1679 != 0 {
		goto L1
	} else {
		goto L375
	}
L342:
	;
	v1566 = int32(1)
	if v377 == int32(0) {
		v2702 = v1566
		goto L3
	} else {
		goto L361
	}
L343:
	;
	v1524 = int32(1)
	if v377 == int32(0) {
		v2702 = v1524
		goto L3
	} else {
		goto L352
	}
L344:
	;
	if v377 == int32(0) {
		v2702 = int32(1)
		goto L3
	} else {
		goto L348
	}
L345:
	;
	v1494 = int32(1)
	if v377 == int32(0) {
		v2702 = v1494
		goto L3
	} else {
		goto L346
	}
L346:
	;
	F_clusterSendFailoverAuthIfNeeded(m, v376, v142)
	mBase = m.M
	v1498 = m.ExcPending
	if v1498 != 0 {
		goto L1
	} else {
		goto L347
	}
L347:
	;
	v2702 = v1494
	goto L3
L348:
	;
	v1502 = int32(1)
	v1503 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v376)+88)))
	if v1503&v1502 == int32(0) {
		v2702 = v1502
		goto L3
	} else {
		goto L349
	}
L349:
	;
	v1508 = *(*int32)(unsafe.Add(mBase, uint32(v376)+2160))
	if v1508 == int32(0) {
		v2702 = v1502
		goto L3
	} else {
		goto L350
	}
L350:
	;
	v1512 = *(*int32)(unsafe.Add(mBase, _consts[90]))
	v1513 = *(*int64)(unsafe.Add(mBase, uint32(v1512)+uint32(_consts[127])))
	if base.Ui64(v618) < base.Ui64(v1513) {
		v2702 = v1502
		goto L3
	} else {
		goto L351
	}
L351:
	;
	v1515 = int32(1)
	v1516 = *(*int32)(unsafe.Add(mBase, uint32(v1512)+uint32(_consts[129])))
	*(*int32)(unsafe.Add(mBase, uint32(v1512)+uint32(_consts[129]))) = v1516 + v1515
	v1520 = *(*int32)(unsafe.Add(mBase, uint32(v1512)+uint32(_consts[101])))
	*(*int32)(unsafe.Add(mBase, uint32(v1512)+uint32(_consts[101]))) = v1520 | v1515
	v2702 = v1515
	goto L3
L352:
	;
	v1527 = *(*int32)(unsafe.Add(mBase, uint32(v376)+2172))
	v1529 = *(*int32)(unsafe.Add(mBase, _consts[94]))
	if v1527 != v1529 {
		v2702 = v1524
		goto L3
	} else {
		goto L353
	}
L353:
	;
	F_resetManualFailover(m)
	mBase = m.M
	v1532 = m.ExcPending
	if v1532 != 0 {
		goto L1
	} else {
		goto L354
	}
L354:
	;
	v1533 = int32(_a69)
	v1534 = *(*int64)(unsafe.Add(mBase, _consts[130]))
	v1536 = *(*int32)(unsafe.Add(mBase, _consts[90]))
	*(*int32)(unsafe.Add(mBase, uint32(v1536)+uint32(_consts[111]))) = v376
	*(*int64)(unsafe.Add(mBase, uint32(v1536)+uint32(_consts[113]))) = v1534 + v66
	F_pauseActions(m, int32(2), v1534<<(uint(int64(1))%64)+v66, int32(29))
	mBase = m.M
	v1546 = m.ExcPending
	if v1546 != 0 {
		goto L1
	} else {
		goto L355
	}
L355:
	;
	v1548 = *(*int32)(unsafe.Add(mBase, _consts[15]))
	if int32(2) < v1548 {
		goto L356
	} else {
		goto L357
	}
L356:
	;
	F_clusterSendPing(m, l0, int32(0))
	mBase = m.M
	v1565 = m.ExcPending
	if v1565 != 0 {
		goto L1
	} else {
		goto L360
	}
L357:
	;
	v1551 = F_humanNodename(m, v376)
	mBase = m.M
	v1552 = m.ExcPending
	if v1552 != 0 {
		goto L1
	} else {
		goto L358
	}
L358:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+388)) = v1551
	*(*int32)(unsafe.Add(mBase, uint32(v24)+384)) = v376 + int32(8)
	F__serverLog(m, int32(2), int32(_a265), v24+int32(384))
	mBase = m.M
	v1562 = m.ExcPending
	if v1562 != 0 {
		goto L1
	} else {
		goto L359
	}
L359:
	;
	goto L356
L360:
	;
	v2702 = v1524
	goto L3
L361:
	;
	v1569 = *(*int64)(unsafe.Add(mBase, uint32(v142)+2256))
	v1573 = F_clusterLookupNode(m, v142+int32(2264), int32(40))
	mBase = m.M
	v1574 = m.ExcPending
	if v1574 != 0 {
		goto L1
	} else {
		goto L362
	}
L362:
	;
	if v1573 == int32(0) {
		v2702 = v1566
		goto L3
	} else {
		goto L363
	}
L363:
	;
	v1577 = *(*int64)(unsafe.Add(mBase, uint32(v1573)+96))
	v1578 = int64(56)
	v1580 = int64(65280)
	v1582 = int64(40)
	v1585 = int64(16711680)
	v1587 = int64(24)
	v1589 = int64(4278190080)
	v1591 = int64(8)
	v1612 = v1569<<(uint(v1578)%64) | v1569&v1580<<(uint(v1582)%64) | (v1569&v1585<<(uint(v1587)%64) | v1569&v1589<<(uint(v1591)%64)) | (int64(base.Ui64(v1569)>>(uint(v1591)%64))&v1589 | int64(base.Ui64(v1569)>>(uint(v1587)%64))&v1585 | (int64(base.Ui64(v1569)>>(uint(v1582)%64))&v1580 | int64(base.Ui64(v1569)>>(uint(v1578)%64))))
	if base.Ui64(v1612) <= base.Ui64(v1577) {
		v2702 = v1566
		goto L3
	} else {
		goto L364
	}
L364:
	;
	v1615 = *(*int32)(unsafe.Add(mBase, _consts[15]))
	if int32(2) < v1615 {
		goto L365
	} else {
		goto L366
	}
L365:
	;
	v1656 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1573)+88)))
	if v1656&int32(2) == int32(0) {
		goto L370
	} else {
		goto L371
	}
L366:
	;
	v1618 = F_humanNodename(m, v376)
	mBase = m.M
	v1619 = m.ExcPending
	if v1619 != 0 {
		goto L1
	} else {
		goto L367
	}
L367:
	;
	v1620 = F_humanNodename(m, v1573)
	mBase = m.M
	v1621 = m.ExcPending
	if v1621 != 0 {
		goto L1
	} else {
		goto L368
	}
L368:
	;
	v1622 = *(*int64)(unsafe.Add(mBase, uint32(v1573)+96))
	*(*int64)(unsafe.Add(mBase, uint32(v24+int32(432)))) = v1612
	*(*int64)(unsafe.Add(mBase, uint32(v24+int32(424)))) = v1622
	v1631 = int32(48)
	*(*int32)(unsafe.Add(mBase, uint32(v24+int32(420)))) = v1573 + v1631
	*(*int32)(unsafe.Add(mBase, uint32(v24+int32(416)))) = v1620
	v1637 = int32(8)
	*(*int32)(unsafe.Add(mBase, uint32(v24)+412)) = v1573 + v1637
	*(*int32)(unsafe.Add(mBase, uint32(v24)+408)) = v376 + v1631
	*(*int32)(unsafe.Add(mBase, uint32(v24)+404)) = v1618
	*(*int32)(unsafe.Add(mBase, uint32(v24)+400)) = v376 + v1637
	F__serverLog(m, int32(2), int32(_a266), v24+int32(400))
	mBase = m.M
	v1652 = m.ExcPending
	if v1652 != 0 {
		goto L1
	} else {
		goto L369
	}
L369:
	;
	goto L365
L370:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v1573)+96)) = v1612
	F_clearCachedClusterSlotsResponse(m)
	mBase = m.M
	v1665 = m.ExcPending
	if v1665 != 0 {
		goto L1
	} else {
		goto L373
	}
L371:
	;
	F_clusterSetNodeAsPrimary(m, v1573)
	mBase = m.M
	v1662 = m.ExcPending
	if v1662 != 0 {
		goto L1
	} else {
		goto L372
	}
L372:
	;
	goto L370
L373:
	;
	v1667 = *(*int32)(unsafe.Add(mBase, _consts[90]))
	v1668 = *(*int32)(unsafe.Add(mBase, uint32(v1667)+uint32(_consts[101])))
	*(*int32)(unsafe.Add(mBase, uint32(v1667)+uint32(_consts[101]))) = v1668 | int32(12)
	F_clusterUpdateSlotsConfigWith(m, v1573, v1612, v142+int32(2304))
	mBase = m.M
	v1675 = m.ExcPending
	if v1675 != 0 {
		goto L1
	} else {
		goto L374
	}
L374:
	;
	v2702 = v1566
	goto L3
L375:
	;
	v2702 = int32(1)
	goto L3
L376:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+368)) = v70
	F__serverLog(m, int32(3), int32(_a267), v24+int32(368))
	mBase = m.M
	v1692 = m.ExcPending
	if v1692 != 0 {
		goto L1
	} else {
		goto L377
	}
L377:
	;
	v2702 = v1681
	goto L3
L378:
	;
	if v377 == int32(0) {
		v2702 = v1698
		goto L3
	} else {
		goto L386
	}
L379:
	;
	v1701 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	if v1701 != 0 {
		goto L378
	} else {
		goto L380
	}
L380:
	;
	v1702 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	*(*int64)(unsafe.Add(mBase, uint32(v1702)+2184)) = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v1702)+2192)) = v66
	v1706 = *(*int32)(unsafe.Add(mBase, uint32(v1702)+88))
	if v1706&int32(4) == int32(0) {
		goto L381
	} else {
		goto L382
	}
L381:
	;
	if v1706&int32(8) == int32(0) {
		goto L378
	} else {
		goto L384
	}
L382:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1702)+88)) = v1706 & int32(-5)
	F_clearCachedClusterSlotsResponse(m)
	mBase = m.M
	v1715 = m.ExcPending
	if v1715 != 0 {
		goto L1
	} else {
		goto L383
	}
L383:
	;
	v1717 = *(*int32)(unsafe.Add(mBase, _consts[90]))
	v1718 = *(*int32)(unsafe.Add(mBase, uint32(v1717)+uint32(_consts[101])))
	*(*int32)(unsafe.Add(mBase, uint32(v1717)+uint32(_consts[101]))) = v1718 | int32(6)
	goto L378
L384:
	;
	F_clearNodeFailureIfNeeded(m, v1702)
	mBase = m.M
	v1727 = m.ExcPending
	if v1727 != 0 {
		goto L1
	} else {
		goto L385
	}
L385:
	;
	goto L378
L386:
	;
	goto L241
L387:
	;
	F__serverAssert(m, int32(_a268), int32(_a179), int32(4350))
	mBase = m.M
	v2655 = m.ExcPending
	if v2655 != 0 {
		goto L1
	} else {
		goto L593
	}
L388:
	;
	F__serverAssert(m, int32(_a269), int32(_a179), int32(4275))
	mBase = m.M
	v2649 = m.ExcPending
	if v2649 != 0 {
		goto L1
	} else {
		goto L592
	}
L389:
	;
	v2624 = int32(1)
	v2625 = int32(0)
	v2626 = *(*int32)(unsafe.Add(mBase, _consts[94]))
	v2627 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2626)+88)))
	if v2627&v2624 == v2625 {
		goto L585
	} else {
		goto L586
	}
L390:
	;
	v2440 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v376)+88)))
	if v2440&int32(1) == int32(0) {
		goto L387
	} else {
		goto L560
	}
L391:
	;
	v2369 = v376 + int32(104)
	v2371 = v142 + int32(80)
	v2372 = int32(2048)
	goto L547
L392:
	;
	v1764 = int32(0)
	goto L403
L393:
	;
	if v380 == int32(0) {
		goto L391
	} else {
		goto L394
	}
L394:
	;
	v1741 = *(*int32)(unsafe.Add(mBase, _consts[15]))
	if int32(0) < v1741 {
		goto L395
	} else {
		goto L396
	}
L395:
	;
	F_clusterSetNodeAsPrimary(m, v376)
	mBase = m.M
	v1763 = m.ExcPending
	if v1763 != 0 {
		goto L1
	} else {
		goto L399
	}
L396:
	;
	v1744 = F_humanNodename(m, v376)
	mBase = m.M
	v1745 = m.ExcPending
	if v1745 != 0 {
		goto L1
	} else {
		goto L397
	}
L397:
	;
	v1746 = int32(48)
	*(*int32)(unsafe.Add(mBase, uint32(v24)+60)) = v376 + v1746
	*(*int32)(unsafe.Add(mBase, uint32(v24)+56)) = int32(_a202)
	*(*int32)(unsafe.Add(mBase, uint32(v24)+52)) = v1744
	*(*int32)(unsafe.Add(mBase, uint32(v24)+48)) = v376 + int32(8)
	F__serverLog(m, int32(0), int32(_a270), v24+v1746)
	mBase = m.M
	v1760 = m.ExcPending
	if v1760 != 0 {
		goto L1
	} else {
		goto L398
	}
L398:
	;
	goto L395
L399:
	;
	goto L390
L400:
	;
	if v381 == int32(0) {
		goto L415
	} else {
		goto L416
	}
L401:
	;
	if int32(0)-v1788 != 0 {
		v1812 = v1764
		goto L400
	} else {
		goto L409
	}
L402:
	;
	goto L401
L403:
	;
	v1773 = int32(0)
	goto L405
L404:
	;
	goto L402
L405:
	;
	v1775 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v378+v1773))))
	v1778 = int32(255)
	v1788 = base.B2i32(base.Ui32((v1775+int32(-123))&v1778) < base.Ui32(int32(230))) & base.B2i32(base.Ui32((v1775+int32(-58))&v1778) < base.Ui32(int32(246)))
	if v1788 != 0 {
		goto L404
	} else {
		goto L407
	}
L406:
	;
	goto L404
L407:
	;
	v1790 = v1773 + int32(1)
	if v1790 != int32(40) {
		v1773 = v1790
		goto L405
	} else {
		goto L408
	}
L408:
	;
	goto L406
L409:
	;
	v1800 = F_sdsnewlen(m, v378, int32(40))
	mBase = m.M
	v1801 = m.ExcPending
	if v1801 != 0 {
		goto L1
	} else {
		goto L410
	}
L410:
	;
	v1803 = *(*int32)(unsafe.Add(mBase, _consts[90]))
	v1804 = *(*int32)(unsafe.Add(mBase, uint32(v1803)+32))
	v1805 = F_dictFind(m, v1804, v1800)
	mBase = m.M
	v1806 = m.ExcPending
	if v1806 != 0 {
		goto L1
	} else {
		goto L411
	}
L411:
	;
	F_sdsfree(m, v1800)
	mBase = m.M
	v1808 = m.ExcPending
	if v1808 != 0 {
		goto L1
	} else {
		goto L412
	}
L412:
	;
	if v1805 == int32(0) {
		v1812 = v1764
		goto L400
	} else {
		goto L413
	}
L413:
	;
	v1811 = *(*int32)(unsafe.Add(mBase, uint32(v1805)+8))
	goto L414
L414:
	;
	v1812 = v1811
	goto L400
L415:
	;
	if v1812 == int32(0) {
		goto L389
	} else {
		goto L495
	}
L416:
	;
	v1818 = *(*int32)(unsafe.Add(mBase, _consts[15]))
	if int32(0) < v1818 {
		goto L417
	} else {
		goto L418
	}
L417:
	;
	if v1812 == int32(0) {
		goto L422
	} else {
		goto L423
	}
L418:
	;
	v1821 = F_humanNodename(m, v376)
	mBase = m.M
	v1822 = m.ExcPending
	if v1822 != 0 {
		goto L1
	} else {
		goto L419
	}
L419:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+252)) = v376 + int32(48)
	*(*int32)(unsafe.Add(mBase, uint32(v24)+248)) = int32(_a204)
	*(*int32)(unsafe.Add(mBase, uint32(v24)+244)) = v1821
	*(*int32)(unsafe.Add(mBase, uint32(v24)+240)) = v376 + int32(8)
	F__serverLog(m, int32(0), int32(_a270), v24+int32(240))
	mBase = m.M
	v1837 = m.ExcPending
	if v1837 != 0 {
		goto L1
	} else {
		goto L420
	}
L420:
	;
	goto L417
L421:
	;
	v2118 = *(*int32)(unsafe.Add(mBase, uint32(v376)+88))
	*(*int32)(unsafe.Add(mBase, uint32(v376)+88)) = v2118&int32(-260) | int32(2)
	F_clearCachedClusterSlotsResponse(m)
	mBase = m.M
	v2125 = m.ExcPending
	if v2125 != 0 {
		goto L1
	} else {
		goto L494
	}
L422:
	;
	v2067 = F_clusterDelNodeSlots(m, v376)
	mBase = m.M
	v2068 = m.ExcPending
	if v2068 != 0 {
		goto L1
	} else {
		goto L484
	}
L423:
	;
	v1841 = int32(48)
	v1842 = v1812 + v1841
	v1844 = v376 + v1841
	v1845 = int32(40)
	goto L428
L424:
	;
	if v1909 != 0 {
		goto L422
	} else {
		goto L440
	}
L425:
	;
	v1909 = int32(0)
	goto L424
L426:
	;
	v1881 = v1876
	v1882 = v1877
	v1883 = v1878
	goto L436
L427:
	;
	if v1866 == int32(0) {
		goto L425
	} else {
		goto L434
	}
L428:
	;
	if (v1844|v1842)&int32(3) != 0 {
		v1876 = v1842
		v1877 = v1844
		v1878 = v1845
		goto L426
	} else {
		goto L429
	}
L429:
	;
	v1853 = v1842
	v1854 = v1844
	v1855 = v1845
	goto L430
L430:
	;
	v1858 = *(*int32)(unsafe.Add(mBase, uint32(v1853)))
	v1859 = *(*int32)(unsafe.Add(mBase, uint32(v1854)))
	if v1858 != v1859 {
		v1876 = v1853
		v1877 = v1854
		v1878 = v1855
		goto L426
	} else {
		goto L432
	}
L431:
	;
	goto L427
L432:
	;
	v1861 = int32(4)
	v1862 = v1854 + v1861
	v1864 = v1853 + v1861
	v1866 = v1855 + int32(-4)
	if base.Ui32(int32(3)) < base.Ui32(v1866) {
		v1853 = v1864
		v1854 = v1862
		v1855 = v1866
		goto L430
	} else {
		goto L433
	}
L433:
	;
	goto L431
L434:
	;
	v1876 = v1864
	v1877 = v1862
	v1878 = v1866
	goto L426
L435:
	;
	v1909 = v1886 - v1887
	goto L424
L436:
	;
	v1886 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1881))))
	v1887 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1882))))
	if v1886 != v1887 {
		goto L435
	} else {
		goto L438
	}
L438:
	;
	v1889 = int32(1)
	v1894 = v1883 + int32(-1)
	if v1894 == int32(0) {
		goto L425
	} else {
		goto L439
	}
L439:
	;
	v1881 = v1881 + v1889
	v1882 = v1882 + v1889
	v1883 = v1894
	goto L436
L440:
	;
	v1910 = *(*int32)(unsafe.Add(mBase, uint32(v376)+2172))
	if v1910 != 0 {
		goto L443
	} else {
		goto L444
	}
L441:
	;
	v1951 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1812)+88)))
	if v1951&int32(2) == int32(0) {
		goto L421
	} else {
		goto L463
	}
L442:
	;
	v1919 = int32(1)
	v1921 = *(*int32)(unsafe.Add(mBase, _consts[15]))
	if int32(2) < v1921 {
		v2702 = v1919
		goto L3
	} else {
		goto L451
	}
L443:
	;
	v1911 = v1910
	goto L445
L444:
	;
	v1911 = v376
	goto L445
L445:
	;
	v1912 = *(*int64)(unsafe.Add(mBase, uint32(v1911)+96))
	if base.Ui64(v617) < base.Ui64(v1912) {
		goto L442
	} else {
		goto L446
	}
L446:
	;
	v1914 = *(*int32)(unsafe.Add(mBase, uint32(v1812)+2172))
	if v1914 != 0 {
		goto L447
	} else {
		goto L448
	}
L447:
	;
	v1915 = v1914
	goto L449
L448:
	;
	v1915 = v1812
	goto L449
L449:
	;
	v1916 = *(*int64)(unsafe.Add(mBase, uint32(v1915)+96))
	if base.Ui64(v1916) <= base.Ui64(v617) {
		goto L441
	} else {
		goto L450
	}
L450:
	;
	goto L442
L451:
	;
	v1924 = F_humanNodename(m, v376)
	mBase = m.M
	v1925 = m.ExcPending
	if v1925 != 0 {
		goto L1
	} else {
		goto L452
	}
L452:
	;
	v1926 = *(*int32)(unsafe.Add(mBase, uint32(v1812)+2172))
	if v1926 != 0 {
		goto L453
	} else {
		goto L454
	}
L453:
	;
	v1927 = v1926
	goto L455
L454:
	;
	v1927 = v1812
	goto L455
L455:
	;
	v1928 = *(*int64)(unsafe.Add(mBase, uint32(v1927)+96))
	v1929 = *(*int32)(unsafe.Add(mBase, uint32(v376)+2172))
	if v1929 != 0 {
		goto L456
	} else {
		goto L457
	}
L456:
	;
	v1930 = v1929
	goto L458
L457:
	;
	v1930 = v376
	goto L458
L458:
	;
	v1931 = *(*int64)(unsafe.Add(mBase, uint32(v1930)+96))
	*(*int64)(unsafe.Add(mBase, uint32(v24+int32(160)))) = v617
	if base.Ui64(v1928) < base.Ui64(v1931) {
		goto L459
	} else {
		goto L460
	}
L459:
	;
	v1938 = v1931
	goto L461
L460:
	;
	v1938 = v1928
	goto L461
L461:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v24+int32(168)))) = v1938
	*(*int32)(unsafe.Add(mBase, uint32(v24)+152)) = v1844
	*(*int32)(unsafe.Add(mBase, uint32(v24)+148)) = v1924
	*(*int32)(unsafe.Add(mBase, uint32(v24)+144)) = v376 + int32(8)
	F__serverLog(m, int32(2), int32(_a271), v24+int32(144))
	mBase = m.M
	v1950 = m.ExcPending
	if v1950 != 0 {
		goto L1
	} else {
		goto L462
	}
L462:
	;
	v2702 = v1919
	goto L3
L463:
	;
	v1956 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v24)+624)) = v1956
	*(*int32)(unsafe.Add(mBase, uint32(v24)+620)) = v1956
	*(*int32)(unsafe.Add(mBase, uint32(v24)+616)) = v1956
	F_clusterMoveNodeSlots(m, v376, v1812, v24+int32(624), v24+int32(620), v24+int32(616))
	mBase = m.M
	v1969 = m.ExcPending
	if v1969 != 0 {
		goto L1
	} else {
		goto L464
	}
L464:
	;
	F_clusterSetNodeAsPrimary(m, v1812)
	mBase = m.M
	v1971 = m.ExcPending
	if v1971 != 0 {
		goto L1
	} else {
		goto L465
	}
L465:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v1812)+96)) = v617
	v1974 = *(*int32)(unsafe.Add(mBase, _consts[15]))
	v1975 = *(*int32)(unsafe.Add(mBase, uint32(v24)+624))
	if v1975 == int32(0) {
		v2011 = v1974
		goto L466
	} else {
		goto L467
	}
L466:
	;
	v2014 = *(*int32)(unsafe.Add(mBase, uint32(v24)+620))
	if v2014 == int32(0) {
		v2035 = v2011
		goto L472
	} else {
		goto L473
	}
L467:
	;
	if int32(2) < v1974 {
		v2011 = v1974
		goto L466
	} else {
		goto L468
	}
L468:
	;
	v1980 = F_humanNodename(m, v376)
	mBase = m.M
	v1981 = m.ExcPending
	if v1981 != 0 {
		goto L1
	} else {
		goto L469
	}
L469:
	;
	v1982 = *(*int32)(unsafe.Add(mBase, uint32(v24)+624))
	v1983 = F_humanNodename(m, v1812)
	mBase = m.M
	v1984 = m.ExcPending
	if v1984 != 0 {
		goto L1
	} else {
		goto L470
	}
L470:
	;
	v1987 = *(*int64)(unsafe.Add(mBase, uint32(v1812)+96))
	*(*int64)(unsafe.Add(mBase, uint32(v24+int32(232)))) = v1987
	*(*int32)(unsafe.Add(mBase, uint32(v24+int32(228)))) = v1983
	v1994 = int32(8)
	*(*int32)(unsafe.Add(mBase, uint32(v24+int32(224)))) = v1812 + v1994
	*(*int32)(unsafe.Add(mBase, uint32(v24)+220)) = v1982
	*(*int32)(unsafe.Add(mBase, uint32(v24)+216)) = v1980
	*(*int32)(unsafe.Add(mBase, uint32(v24)+212)) = v376 + v1994
	*(*int32)(unsafe.Add(mBase, uint32(v24)+208)) = v1844
	F__serverLog(m, int32(2), int32(_a272), v24+int32(208))
	mBase = m.M
	v2008 = m.ExcPending
	if v2008 != 0 {
		goto L1
	} else {
		goto L471
	}
L471:
	;
	v2010 = *(*int32)(unsafe.Add(mBase, _consts[15]))
	v2011 = v2010
	goto L466
L472:
	;
	v2036 = *(*int32)(unsafe.Add(mBase, uint32(v24)+616))
	if v2036 == int32(0) {
		goto L477
	} else {
		goto L478
	}
L473:
	;
	if int32(2) < v2011 {
		v2035 = v2011
		goto L472
	} else {
		goto L474
	}
L474:
	;
	v2019 = F_humanNodename(m, v1812)
	mBase = m.M
	v2020 = m.ExcPending
	if v2020 != 0 {
		goto L1
	} else {
		goto L475
	}
L475:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+204)) = v1842
	*(*int32)(unsafe.Add(mBase, uint32(v24)+200)) = v2019
	*(*int32)(unsafe.Add(mBase, uint32(v24)+196)) = v1812 + int32(8)
	*(*int32)(unsafe.Add(mBase, uint32(v24)+192)) = v2014
	F__serverLog(m, int32(2), int32(_a273), v24+int32(192))
	mBase = m.M
	v2032 = m.ExcPending
	if v2032 != 0 {
		goto L1
	} else {
		goto L476
	}
L476:
	;
	v2034 = *(*int32)(unsafe.Add(mBase, _consts[15]))
	v2035 = v2034
	goto L472
L477:
	;
	v2056 = *(*int32)(unsafe.Add(mBase, uint32(v376)+2160))
	if v2056 == int32(0) {
		goto L421
	} else {
		goto L482
	}
L478:
	;
	if int32(2) < v2035 {
		goto L477
	} else {
		goto L479
	}
L479:
	;
	v2041 = F_humanNodename(m, v1812)
	mBase = m.M
	v2042 = m.ExcPending
	if v2042 != 0 {
		goto L1
	} else {
		goto L480
	}
L480:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+188)) = v1842
	*(*int32)(unsafe.Add(mBase, uint32(v24)+184)) = v2041
	*(*int32)(unsafe.Add(mBase, uint32(v24)+180)) = v1812 + int32(8)
	*(*int32)(unsafe.Add(mBase, uint32(v24)+176)) = v2036
	F__serverLog(m, int32(2), int32(_a274), v24+int32(176))
	mBase = m.M
	v2054 = m.ExcPending
	if v2054 != 0 {
		goto L1
	} else {
		goto L481
	}
L481:
	;
	goto L477
L482:
	;
	F__serverAssert(m, int32(_a269), int32(_a179), int32(4262))
	mBase = m.M
	v2063 = m.ExcPending
	if v2063 != 0 {
		goto L1
	} else {
		goto L483
	}
L483:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L484:
	;
	v2070 = *(*int32)(unsafe.Add(mBase, _consts[15]))
	if int32(2) < v2070 {
		goto L485
	} else {
		goto L486
	}
L485:
	;
	v2112 = *(*int32)(unsafe.Add(mBase, uint32(v376)+2160))
	if v2112 != 0 {
		goto L388
	} else {
		goto L493
	}
L486:
	;
	v2073 = F_humanNodename(m, v376)
	mBase = m.M
	v2074 = m.ExcPending
	if v2074 != 0 {
		goto L1
	} else {
		goto L487
	}
L487:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+140)) = v2067
	*(*int32)(unsafe.Add(mBase, uint32(v24)+136)) = v376 + int32(48)
	*(*int32)(unsafe.Add(mBase, uint32(v24)+132)) = v2073
	*(*int32)(unsafe.Add(mBase, uint32(v24)+128)) = v376 + int32(8)
	F__serverLog(m, int32(2), int32(_a275), v24+int32(128))
	mBase = m.M
	v2088 = m.ExcPending
	if v2088 != 0 {
		goto L1
	} else {
		goto L488
	}
L488:
	;
	if v1812 == int32(0) {
		goto L485
	} else {
		goto L489
	}
L489:
	;
	v2092 = *(*int32)(unsafe.Add(mBase, _consts[15]))
	if int32(2) < v2092 {
		goto L485
	} else {
		goto L490
	}
L490:
	;
	v2095 = F_humanNodename(m, v376)
	mBase = m.M
	v2096 = m.ExcPending
	if v2096 != 0 {
		goto L1
	} else {
		goto L491
	}
L491:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+120)) = v1812 + int32(48)
	*(*int32)(unsafe.Add(mBase, uint32(v24)+116)) = v2095
	*(*int32)(unsafe.Add(mBase, uint32(v24)+112)) = v376 + int32(8)
	F__serverLog(m, int32(2), int32(_a276), v24+int32(112))
	mBase = m.M
	v2109 = m.ExcPending
	if v2109 != 0 {
		goto L1
	} else {
		goto L492
	}
L492:
	;
	goto L485
L493:
	;
	goto L421
L494:
	;
	v2127 = *(*int32)(unsafe.Add(mBase, _consts[90]))
	v2128 = *(*int32)(unsafe.Add(mBase, uint32(v2127)+uint32(_consts[101])))
	*(*int32)(unsafe.Add(mBase, uint32(v2127)+uint32(_consts[101]))) = v2128 | int32(14)
	goto L415
L495:
	;
	v2139 = *(*int32)(unsafe.Add(mBase, uint32(v376)+2172))
	if v2139 == v1812 {
		goto L389
	} else {
		goto L496
	}
L496:
	;
	if v2139 == int32(0) {
		goto L497
	} else {
		goto L498
	}
L497:
	;
	v2210 = *(*int32)(unsafe.Add(mBase, _consts[15]))
	if int32(2) < v2210 {
		goto L510
	} else {
		goto L511
	}
L498:
	;
	v2150 = *(*int32)(unsafe.Add(mBase, uint32(v2139)+2164))
	if v2150 < int32(1) {
		goto L500
	} else {
		goto L501
	}
L499:
	;
	goto L497
L500:
	;
	goto L499
L501:
	;
	v2153 = *(*int32)(unsafe.Add(mBase, uint32(v2139)+2168))
	v2160 = int32(0)
	goto L502
L502:
	;
	v2164 = v2160 + int32(1)
	v2167 = v2153 + v2160<<(uint(int32(2))%32)
	v2168 = *(*int32)(unsafe.Add(mBase, uint32(v2167)))
	if v2168 != v376 {
		goto L504
	} else {
		goto L505
	}
L503:
	;
	goto L500
L504:
	;
	if v2164 != v2150 {
		v2160 = v2164
		goto L502
	} else {
		goto L509
	}
L505:
	;
	if v2150 <= v2164 {
		v2181 = v2150
		goto L506
	} else {
		goto L507
	}
L506:
	;
	v2183 = v2181 + int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(v2139)+2164)) = v2183
	if v2183 != 0 {
		goto L500
	} else {
		goto L508
	}
L507:
	;
	v2171 = int32(2)
	v2179 = F_memmove(m, v2167, v2153+v2164<<(uint(v2171)%32), (v2150+(v2160^int32(-1)))<<(uint(v2171)%32))
	mBase = m.M
	v2180 = *(*int32)(unsafe.Add(mBase, uint32(v2139)+2164))
	v2181 = v2180
	goto L506
L508:
	;
	v2186 = *(*int32)(unsafe.Add(mBase, uint32(v2139)+88))
	*(*int32)(unsafe.Add(mBase, uint32(v2139)+88)) = v2186 & int32(-257)
	goto L499
L509:
	;
	goto L503
L510:
	;
	v2238 = F_clusterNodeAddReplica(m, v1812, v376)
	mBase = m.M
	v2239 = m.ExcPending
	if v2239 != 0 {
		goto L1
	} else {
		goto L515
	}
L511:
	;
	v2213 = F_humanNodename(m, v376)
	mBase = m.M
	v2214 = m.ExcPending
	if v2214 != 0 {
		goto L1
	} else {
		goto L512
	}
L512:
	;
	v2215 = F_humanNodename(m, v1812)
	mBase = m.M
	v2216 = m.ExcPending
	if v2216 != 0 {
		goto L1
	} else {
		goto L513
	}
L513:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24+int32(96)))) = v1812 + int32(48)
	*(*int32)(unsafe.Add(mBase, uint32(v24)+92)) = v2215
	v2223 = int32(8)
	*(*int32)(unsafe.Add(mBase, uint32(v24)+88)) = v1812 + v2223
	*(*int32)(unsafe.Add(mBase, uint32(v24)+84)) = v2213
	*(*int32)(unsafe.Add(mBase, uint32(v24)+80)) = v376 + v2223
	F__serverLog(m, int32(2), int32(_a277), v24+int32(80))
	mBase = m.M
	v2235 = m.ExcPending
	if v2235 != 0 {
		goto L1
	} else {
		goto L514
	}
L514:
	;
	goto L510
L515:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v376)+2172)) = v1812
	v2241 = int32(0)
	v2242 = *(*int32)(unsafe.Add(mBase, _consts[94]))
	v2243 = *(*int32)(unsafe.Add(mBase, uint32(v2242)+2172))
	if v2243 == v2241 {
		goto L516
	} else {
		goto L517
	}
L516:
	;
	F_updateShardId(m, v376, v1812+int32(48))
	mBase = m.M
	v2359 = m.ExcPending
	if v2359 != 0 {
		goto L1
	} else {
		goto L541
	}
L517:
	;
	v2246 = *(*int32)(unsafe.Add(mBase, uint32(v2243)+2172))
	if v2246 == int32(0) {
		goto L516
	} else {
		goto L518
	}
L518:
	;
	if v2246 == v2242 {
		goto L516
	} else {
		goto L519
	}
L519:
	;
	v2251 = *(*int32)(unsafe.Add(mBase, _consts[15]))
	if int32(2) < v2251 {
		v2270 = v2246
		v2271 = v2242
		goto L520
	} else {
		goto L521
	}
L520:
	;
	v2273 = int32(48)
	v2274 = v2270 + v2273
	v2276 = v2271 + v2273
	v2277 = int32(40)
	goto L527
L521:
	;
	v2254 = int32(8)
	*(*int32)(unsafe.Add(mBase, uint32(v24)+68)) = v2243 + v2254
	*(*int32)(unsafe.Add(mBase, uint32(v24)+64)) = v2246 + v2254
	F__serverLog(m, int32(2), int32(_a278), v24+int32(64))
	mBase = m.M
	v2265 = m.ExcPending
	if v2265 != 0 {
		goto L1
	} else {
		goto L522
	}
L522:
	;
	v2267 = *(*int32)(unsafe.Add(mBase, _consts[94]))
	v2268 = *(*int32)(unsafe.Add(mBase, uint32(v2267)+2172))
	v2269 = *(*int32)(unsafe.Add(mBase, uint32(v2268)+2172))
	v2270 = v2269
	v2271 = v2267
	goto L520
L523:
	;
	F_clusterSetPrimary(m, v2270, int32(1), base.B2i32(v2341 != int32(0)))
	mBase = m.M
	v2345 = m.ExcPending
	if v2345 != 0 {
		goto L1
	} else {
		goto L539
	}
L524:
	;
	v2341 = int32(0)
	goto L523
L525:
	;
	v2313 = v2308
	v2314 = v2309
	v2315 = v2310
	goto L535
L526:
	;
	if v2298 == int32(0) {
		goto L524
	} else {
		goto L533
	}
L527:
	;
	if (v2276|v2274)&int32(3) != 0 {
		v2308 = v2274
		v2309 = v2276
		v2310 = v2277
		goto L525
	} else {
		goto L528
	}
L528:
	;
	v2285 = v2274
	v2286 = v2276
	v2287 = v2277
	goto L529
L529:
	;
	v2290 = *(*int32)(unsafe.Add(mBase, uint32(v2285)))
	v2291 = *(*int32)(unsafe.Add(mBase, uint32(v2286)))
	if v2290 != v2291 {
		v2308 = v2285
		v2309 = v2286
		v2310 = v2287
		goto L525
	} else {
		goto L531
	}
L530:
	;
	goto L526
L531:
	;
	v2293 = int32(4)
	v2294 = v2286 + v2293
	v2296 = v2285 + v2293
	v2298 = v2287 + int32(-4)
	if base.Ui32(int32(3)) < base.Ui32(v2298) {
		v2285 = v2296
		v2286 = v2294
		v2287 = v2298
		goto L529
	} else {
		goto L532
	}
L532:
	;
	goto L530
L533:
	;
	v2308 = v2296
	v2309 = v2294
	v2310 = v2298
	goto L525
L534:
	;
	v2341 = v2318 - v2319
	goto L523
L535:
	;
	v2318 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2313))))
	v2319 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2314))))
	if v2318 != v2319 {
		goto L534
	} else {
		goto L537
	}
L537:
	;
	v2321 = int32(1)
	v2326 = v2315 + int32(-1)
	if v2326 == int32(0) {
		goto L524
	} else {
		goto L538
	}
L538:
	;
	v2313 = v2313 + v2321
	v2314 = v2314 + v2321
	v2315 = v2326
	goto L535
L539:
	;
	F_clearCachedClusterSlotsResponse(m)
	mBase = m.M
	v2347 = m.ExcPending
	if v2347 != 0 {
		goto L1
	} else {
		goto L540
	}
L540:
	;
	v2349 = *(*int32)(unsafe.Add(mBase, _consts[90]))
	v2350 = *(*int32)(unsafe.Add(mBase, uint32(v2349)+uint32(_consts[101])))
	*(*int32)(unsafe.Add(mBase, uint32(v2349)+uint32(_consts[101]))) = v2350 | int32(46)
	goto L516
L541:
	;
	F_clearCachedClusterSlotsResponse(m)
	mBase = m.M
	v2361 = m.ExcPending
	if v2361 != 0 {
		goto L1
	} else {
		goto L542
	}
L542:
	;
	v2363 = *(*int32)(unsafe.Add(mBase, _consts[90]))
	v2364 = *(*int32)(unsafe.Add(mBase, uint32(v2363)+uint32(_consts[101])))
	*(*int32)(unsafe.Add(mBase, uint32(v2363)+uint32(_consts[101]))) = v2364 | int32(4)
	goto L389
L543:
	;
	if v2436 == int32(0) {
		goto L389
	} else {
		goto L559
	}
L544:
	;
	v2436 = int32(0)
	goto L543
L545:
	;
	v2408 = v2403
	v2409 = v2404
	v2410 = v2405
	goto L555
L546:
	;
	if v2393 == int32(0) {
		goto L544
	} else {
		goto L553
	}
L547:
	;
	if (v2371|v2369)&int32(3) != 0 {
		v2403 = v2369
		v2404 = v2371
		v2405 = v2372
		goto L545
	} else {
		goto L548
	}
L548:
	;
	v2380 = v2369
	v2381 = v2371
	v2382 = v2372
	goto L549
L549:
	;
	v2385 = *(*int32)(unsafe.Add(mBase, uint32(v2380)))
	v2386 = *(*int32)(unsafe.Add(mBase, uint32(v2381)))
	if v2385 != v2386 {
		v2403 = v2380
		v2404 = v2381
		v2405 = v2382
		goto L545
	} else {
		goto L551
	}
L550:
	;
	goto L546
L551:
	;
	v2388 = int32(4)
	v2389 = v2381 + v2388
	v2391 = v2380 + v2388
	v2393 = v2382 + int32(-4)
	if base.Ui32(int32(3)) < base.Ui32(v2393) {
		v2380 = v2391
		v2381 = v2389
		v2382 = v2393
		goto L549
	} else {
		goto L552
	}
L552:
	;
	goto L550
L553:
	;
	v2403 = v2391
	v2404 = v2389
	v2405 = v2393
	goto L545
L554:
	;
	v2436 = v2413 - v2414
	goto L543
L555:
	;
	v2413 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2408))))
	v2414 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2409))))
	if v2413 != v2414 {
		goto L554
	} else {
		goto L557
	}
L557:
	;
	v2416 = int32(1)
	v2421 = v2410 + int32(-1)
	if v2421 == int32(0) {
		goto L544
	} else {
		goto L558
	}
L558:
	;
	v2408 = v2408 + v2416
	v2409 = v2409 + v2416
	v2410 = v2421
	goto L555
L559:
	;
	goto L390
L560:
	;
	v2446 = *(*int32)(unsafe.Add(mBase, _consts[15]))
	if int32(2) < v2446 {
		goto L561
	} else {
		goto L562
	}
L561:
	;
	v2466 = v142 + int32(80)
	F_clusterUpdateSlotsConfigWith(m, v376, v617, v2466)
	mBase = m.M
	v2468 = m.ExcPending
	if v2468 != 0 {
		goto L1
	} else {
		goto L565
	}
L562:
	;
	v2449 = F_humanNodename(m, v376)
	mBase = m.M
	v2450 = m.ExcPending
	if v2450 != 0 {
		goto L1
	} else {
		goto L563
	}
L563:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+40)) = v376 + int32(48)
	*(*int32)(unsafe.Add(mBase, uint32(v24)+36)) = v2449
	*(*int32)(unsafe.Add(mBase, uint32(v24)+32)) = v376 + int32(8)
	F__serverLog(m, int32(2), int32(_a279), v24+int32(32))
	mBase = m.M
	v2463 = m.ExcPending
	if v2463 != 0 {
		goto L1
	} else {
		goto L564
	}
L564:
	;
	goto L561
L565:
	;
	v2472 = *(*int32)(unsafe.Add(mBase, _consts[90]))
	v2490 = int32(0)
	goto L566
L566:
	;
	v2502 = *(*int64)(unsafe.Add(mBase, uint32(v2466+v2490<<(uint(int32(3))%32))))
	if v2502 == int64(0) {
		goto L568
	} else {
		goto L569
	}
L567:
	;
	goto L389
L568:
	;
	if base.Ui32(v2490) < base.Ui32(int32(255)) {
		v2490 = v2490 + int32(1)
		goto L566
	} else {
		goto L584
	}
L569:
	;
	v2512 = v2502
	goto L570
L570:
	;
	v2529 = base.I32_wrap_i64(base.I64_ctz(v2512))
	v2530 = v2490<<(uint(int32(6))%32) | v2529
	v2534 = *(*int32)(unsafe.Add(mBase, uint32(v2472+int32(52)+v2530<<(uint(int32(2))%32))))
	if v2534 == v376 {
		goto L572
	} else {
		goto L573
	}
L571:
	;
	goto L568
L572:
	;
	v2575 = (v2512 + int64(-1)) & v2512
	if v2575 != int64(0) {
		v2512 = v2575
		goto L570
	} else {
		goto L583
	}
L573:
	;
	if v2534 == int32(0) {
		goto L572
	} else {
		goto L574
	}
L574:
	;
	v2541 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2472+int32(65920)+int32(base.Ui32(v2530)>>(uint(int32(3))%32))))))
	if int32(base.Ui32(v2541)>>(uint(v2529&int32(7))%32))&int32(1) != 0 {
		goto L572
	} else {
		goto L575
	}
L575:
	;
	v2547 = *(*int64)(unsafe.Add(mBase, uint32(v2534)+96))
	if base.Ui64(v2547) <= base.Ui64(v617) {
		goto L572
	} else {
		goto L576
	}
L576:
	;
	v2550 = *(*int32)(unsafe.Add(mBase, _consts[15]))
	if int32(1) < v2550 {
		goto L577
	} else {
		goto L578
	}
L577:
	;
	v2570 = *(*int32)(unsafe.Add(mBase, uint32(v376)+2344))
	F_clusterSendUpdate(m, v2570, v2534)
	mBase = m.M
	v2572 = m.ExcPending
	if v2572 != 0 {
		goto L1
	} else {
		goto L582
	}
L578:
	;
	v2553 = F_humanNodename(m, v376)
	mBase = m.M
	v2554 = m.ExcPending
	if v2554 != 0 {
		goto L1
	} else {
		goto L579
	}
L579:
	;
	v2555 = F_humanNodename(m, v2534)
	mBase = m.M
	v2556 = m.ExcPending
	if v2556 != 0 {
		goto L1
	} else {
		goto L580
	}
L580:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+28)) = v2555
	*(*int32)(unsafe.Add(mBase, uint32(v24)+24)) = v2534 + int32(8)
	*(*int32)(unsafe.Add(mBase, uint32(v24)+20)) = v2553
	*(*int32)(unsafe.Add(mBase, uint32(v24)+16)) = v376 + int32(8)
	F__serverLog(m, int32(1), int32(_a280), v24+int32(16))
	mBase = m.M
	v2568 = m.ExcPending
	if v2568 != 0 {
		goto L1
	} else {
		goto L581
	}
L581:
	;
	goto L577
L582:
	;
	goto L389
L583:
	;
	goto L571
L584:
	;
	goto L567
L585:
	;
	F_clusterProcessGossipSection(m, v142, l0)
	mBase = m.M
	v2642 = m.ExcPending
	if v2642 != 0 {
		goto L1
	} else {
		goto L590
	}
L586:
	;
	v2632 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v376)+88)))
	if v2632&int32(1) == int32(0) {
		goto L585
	} else {
		goto L587
	}
L587:
	;
	v2637 = *(*int64)(unsafe.Add(mBase, uint32(v2626)+96))
	if v617 != v2637 {
		goto L585
	} else {
		goto L588
	}
L588:
	;
	F_clusterHandleConfigEpochCollision(m, v376)
	mBase = m.M
	v2640 = m.ExcPending
	if v2640 != 0 {
		goto L1
	} else {
		goto L589
	}
L589:
	;
	goto L585
L590:
	;
	F_clusterProcessPingExtensions(m, v142, l0)
	mBase = m.M
	v2644 = m.ExcPending
	if v2644 != 0 {
		goto L1
	} else {
		goto L591
	}
L591:
	;
	v2702 = v2624
	goto L3
L592:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L593:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L594:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+600)) = v2667
	*(*int32)(unsafe.Add(mBase, uint32(v24)+596)) = v2657
	*(*int32)(unsafe.Add(mBase, uint32(v24)+592)) = v2658
	F__serverLog(m, int32(2), int32(_a281), v24+int32(592))
	mBase = m.M
	v2676 = m.ExcPending
	if v2676 != 0 {
		goto L1
	} else {
		goto L596
	}
L595:
	;
	v2666 = *(*int32)(unsafe.Add(mBase, uint32(v70<<(uint(int32(2))%32))+uint32(_consts[125])))
	v2667 = v2666
	goto L594
L596:
	;
	goto L17
L597:
	;
	v2702 = int32(0)
	goto L3
}
func F_clusterProcessPingExtensions(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
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
	var v73 int32
	_ = v73
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
	var v92 int32
	_ = v92
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
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
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v131 int32
	_ = v131
	var v140 int32
	_ = v140
	var v142 int32
	_ = v142
	var v145 int32
	_ = v145
	var v155 int32
	_ = v155
	var v157 int32
	_ = v157
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v175 int32
	_ = v175
	var v178 int32
	_ = v178
	var v182 int32
	_ = v182
	var v184 int32
	_ = v184
	var v189 int32
	_ = v189
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v202 int32
	_ = v202
	var v203 int64
	_ = v203
	var v204 int64
	_ = v204
	var v206 int64
	_ = v206
	var v208 int64
	_ = v208
	var v211 int64
	_ = v211
	var v213 int64
	_ = v213
	var v215 int64
	_ = v215
	var v217 int64
	_ = v217
	var v240 int64
	_ = v240
	var v244 int32
	_ = v244
	var v246 int32
	_ = v246
	var v248 int32
	_ = v248
	var v249 int32
	_ = v249
	var v258 int32
	_ = v258
	var v265 int32
	_ = v265
	var v267 int32
	_ = v267
	var v268 int32
	_ = v268
	var v269 int32
	_ = v269
	var v270 int32
	_ = v270
	var v271 int32
	_ = v271
	var v272 int32
	_ = v272
	var v273 int32
	_ = v273
	var v274 int32
	_ = v274
	var v278 int32
	_ = v278
	var v279 int32
	_ = v279
	var v282 int32
	_ = v282
	var v290 int32
	_ = v290
	var v291 int32
	_ = v291
	var v292 int32
	_ = v292
	var v293 int32
	_ = v293
	var v294 int32
	_ = v294
	var v295 int32
	_ = v295
	var v296 int32
	_ = v296
	var v297 int32
	_ = v297
	var v304 int32
	_ = v304
	var v308 int32
	_ = v308
	var v312 int32
	_ = v312
	var v316 int32
	_ = v316
	var v317 int32
	_ = v317
	var v321 int32
	_ = v321
	var v323 int32
	_ = v323
	var v324 int32
	_ = v324
	var v329 int32
	_ = v329
	var v333 int32
	_ = v333
	var v335 int32
	_ = v335
	var v336 int32
	_ = v336
	var v344 int32
	_ = v344
	var v346 int32
	_ = v346
	var v361 int32
	_ = v361
	var v365 int32
	_ = v365
	var v367 int32
	_ = v367
	var v370 int32
	_ = v370
	var v378 int32
	_ = v378
	var v390 int32
	_ = v390
	var v398 int32
	_ = v398
	v17 = m.G0
	v19 = v17 - int32(16)
	m.G0 = v19
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l1)+44))
	if v21 != 0 {
		v73 = v21
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v75 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+2214)))
	v76 = F___bswap_16_2(m, v75)
	mBase = m.M
	goto L18
L2:
	;
	v22 = int32(0)
	v24 = l0 + int32(40)
	goto L5
L3:
	;
	if int32(0)-v48 != 0 {
		v73 = v22
		goto L1
	} else {
		goto L11
	}
L4:
	;
	goto L3
L5:
	;
	v33 = int32(0)
	goto L7
L6:
	;
	goto L4
L7:
	;
	v35 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24+v33))))
	v38 = int32(255)
	v48 = base.B2i32(base.Ui32((v35+int32(-123))&v38) < base.Ui32(int32(230))) & base.B2i32(base.Ui32((v35+int32(-58))&v38) < base.Ui32(int32(246)))
	if v48 != 0 {
		goto L6
	} else {
		goto L9
	}
L8:
	;
	goto L6
L9:
	;
	v50 = v33 + int32(1)
	if v50 != int32(40) {
		v33 = v50
		goto L7
	} else {
		goto L10
	}
L10:
	;
	goto L8
L11:
	;
	v60 = F_sdsnewlen(m, v24, int32(40))
	mBase = m.M
	v61 = m.ExcPending
	if v61 != 0 {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	return
L13:
	;
	v63 = *(*int32)(unsafe.Add(mBase, _consts[90]))
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v63)+32))
	v65 = F_dictFind(m, v64, v60)
	mBase = m.M
	v66 = m.ExcPending
	if v66 != 0 {
		goto L12
	} else {
		goto L14
	}
L14:
	;
	F_sdsfree(m, v60)
	mBase = m.M
	v68 = m.ExcPending
	if v68 != 0 {
		goto L12
	} else {
		goto L15
	}
L15:
	;
	if v65 == int32(0) {
		v73 = v22
		goto L1
	} else {
		goto L16
	}
L16:
	;
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v65)+8))
	goto L17
L17:
	;
	v73 = v71
	goto L1
L18:
	;
	v77 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+14)))
	v78 = F___bswap_16_2(m, v77)
	mBase = m.M
	goto L19
L19:
	;
	if v76 != 0 {
		goto L21
	} else {
		goto L22
	}
L20:
	;
	F_updateSdsExtensionField(m, v73+int32(2312), v297)
	mBase = m.M
	v304 = m.ExcPending
	if v304 != 0 {
		goto L12
	} else {
		goto L71
	}
L21:
	;
	v92 = int32(0)
	v100 = l0 + v78*int32(104) + int32(2256)
	v101 = v76
	v105 = v92
	v106 = v92
	v107 = v92
	v108 = v92
	v109 = v92
	v110 = v92
	v111 = v92
	v112 = v92
	goto L23
L22:
	;
	v79 = int32(0)
	v290 = v79
	v291 = v79
	v292 = v79
	v293 = v79
	v294 = v79
	v295 = v79
	v296 = v79
	v297 = v79
	goto L20
L23:
	;
	v116 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v100)+4)))
	v117 = F___bswap_16_2(m, v116)
	mBase = m.M
	goto L36
L24:
	;
	v290 = v267
	v291 = v268
	v292 = v269
	v293 = v270
	v294 = v271
	v295 = v272
	v296 = v273
	v297 = v274
	goto L20
L25:
	;
	v278 = *(*int32)(unsafe.Add(mBase, uint32(v100)))
	v279 = F___bswap_32_2(m, v278)
	mBase = m.M
	goto L69
L26:
	;
	v258 = *(*int32)(unsafe.Add(mBase, _consts[15]))
	if int32(3) < v258 {
		v267 = v105
		v268 = v106
		v269 = v107
		v270 = v108
		v271 = v109
		v272 = v110
		v273 = v111
		v274 = v112
		goto L25
	} else {
		goto L67
	}
L27:
	;
	v267 = v105
	v268 = v106
	v269 = v107
	v270 = v108
	v271 = v109
	v272 = v100 + int32(8)
	v273 = v111
	v274 = v112
	goto L25
L28:
	;
	v267 = v100 + int32(8)
	v268 = v106
	v269 = v107
	v270 = v108
	v271 = v109
	v272 = v110
	v273 = v111
	v274 = v112
	goto L25
L29:
	;
	v131 = v100 + int32(8)
	goto L41
L30:
	;
	v128 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v100)+8)))
	v129 = F___bswap_16_2(m, v128)
	mBase = m.M
	goto L38
L31:
	;
	v126 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v100)+8)))
	v127 = F___bswap_16_2(m, v126)
	mBase = m.M
	goto L37
L32:
	;
	v267 = v105
	v268 = v106
	v269 = v107
	v270 = v100 + int32(8)
	v271 = v109
	v272 = v110
	v273 = v111
	v274 = v112
	goto L25
L33:
	;
	v267 = v105
	v268 = v106
	v269 = v107
	v270 = v108
	v271 = v100 + int32(8)
	v272 = v110
	v273 = v111
	v274 = v112
	goto L25
L34:
	;
	v267 = v105
	v268 = v106
	v269 = v107
	v270 = v108
	v271 = v109
	v272 = v110
	v273 = v100 + int32(8)
	v274 = v112
	goto L25
L35:
	;
	v267 = v105
	v268 = v106
	v269 = v107
	v270 = v108
	v271 = v109
	v272 = v110
	v273 = v111
	v274 = v100 + int32(8)
	goto L25
L36:
	;
	switch v117 {
	case 0:
		goto L35
	case 1:
		goto L34
	case 2:
		goto L29
	case 3:
		goto L28
	case 4:
		goto L33
	case 5:
		goto L32
	case 6:
		goto L31
	case 7:
		goto L30
	case 8:
		goto L27
	default:
		goto L26
	}
L37:
	;
	v267 = v105
	v268 = v106
	v269 = v127
	v270 = v108
	v271 = v109
	v272 = v110
	v273 = v111
	v274 = v112
	goto L25
L38:
	;
	v267 = v105
	v268 = v129
	v269 = v107
	v270 = v108
	v271 = v109
	v272 = v110
	v273 = v111
	v274 = v112
	goto L25
L39:
	;
	if int32(0)-v155 != 0 {
		v267 = v105
		v268 = v106
		v269 = v107
		v270 = v108
		v271 = v109
		v272 = v110
		v273 = v111
		v274 = v112
		goto L25
	} else {
		goto L47
	}
L40:
	;
	goto L39
L41:
	;
	v140 = int32(0)
	goto L43
L42:
	;
	goto L40
L43:
	;
	v142 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v131+v140))))
	v145 = int32(255)
	v155 = base.B2i32(base.Ui32((v142+int32(-123))&v145) < base.Ui32(int32(230))) & base.B2i32(base.Ui32((v142+int32(-58))&v145) < base.Ui32(int32(246)))
	if v155 != 0 {
		goto L42
	} else {
		goto L45
	}
L44:
	;
	goto L42
L45:
	;
	v157 = v140 + int32(1)
	if v157 != int32(40) {
		v140 = v157
		goto L43
	} else {
		goto L46
	}
L46:
	;
	goto L44
L47:
	;
	v167 = F_sdsnewlen(m, v131, int32(40))
	mBase = m.M
	v168 = m.ExcPending
	if v168 != 0 {
		goto L12
	} else {
		goto L48
	}
L48:
	;
	v170 = *(*int32)(unsafe.Add(mBase, _consts[90]))
	v171 = *(*int32)(unsafe.Add(mBase, uint32(v170)+32))
	v172 = F_dictFind(m, v171, v167)
	mBase = m.M
	v173 = m.ExcPending
	if v173 != 0 {
		goto L12
	} else {
		goto L49
	}
L49:
	;
	F_sdsfree(m, v167)
	mBase = m.M
	v175 = m.ExcPending
	if v175 != 0 {
		goto L12
	} else {
		goto L50
	}
L50:
	;
	if v172 == int32(0) {
		v267 = v105
		v268 = v106
		v269 = v107
		v270 = v108
		v271 = v109
		v272 = v110
		v273 = v111
		v274 = v112
		goto L25
	} else {
		goto L51
	}
L51:
	;
	v178 = *(*int32)(unsafe.Add(mBase, uint32(v172)+8))
	goto L52
L52:
	;
	if v178 == int32(0) {
		v267 = v105
		v268 = v106
		v269 = v107
		v270 = v108
		v271 = v109
		v272 = v110
		v273 = v111
		v274 = v112
		goto L25
	} else {
		goto L53
	}
L53:
	;
	v182 = *(*int32)(unsafe.Add(mBase, _consts[94]))
	if v178 == v182 {
		v267 = v105
		v268 = v106
		v269 = v107
		v270 = v108
		v271 = v109
		v272 = v110
		v273 = v111
		v274 = v112
		goto L25
	} else {
		goto L54
	}
L54:
	;
	v184 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v182)+88)))
	if v184&int32(2) == int32(0) {
		goto L55
	} else {
		goto L56
	}
L55:
	;
	v192 = F_sdsnewlen(m, v131, int32(40))
	mBase = m.M
	v193 = m.ExcPending
	if v193 != 0 {
		goto L12
	} else {
		goto L58
	}
L56:
	;
	v189 = *(*int32)(unsafe.Add(mBase, uint32(v182)+2172))
	if v189 == v178 {
		v267 = v105
		v268 = v106
		v269 = v107
		v270 = v108
		v271 = v109
		v272 = v110
		v273 = v111
		v274 = v112
		goto L25
	} else {
		goto L57
	}
L57:
	;
	goto L55
L58:
	;
	v195 = *(*int32)(unsafe.Add(mBase, _consts[90]))
	v196 = *(*int32)(unsafe.Add(mBase, uint32(v195)+40))
	v197 = F_dictAddOrFind(m, v196, v192)
	mBase = m.M
	v198 = m.ExcPending
	if v198 != 0 {
		goto L12
	} else {
		goto L60
	}
L59:
	;
	v203 = *(*int64)(unsafe.Add(mBase, uint32(v100)+48))
	v204 = int64(56)
	v206 = int64(65280)
	v208 = int64(40)
	v211 = int64(16711680)
	v213 = int64(24)
	v215 = int64(4278190080)
	v217 = int64(8)
	v240 = *(*int64)(unsafe.Add(mBase, _consts[37]))
	*(*int64)(unsafe.Add(mBase, uint32(v197)+8)) = v203<<(uint(v204)%64) | v203&v206<<(uint(v208)%64) | (v203&v211<<(uint(v213)%64) | v203&v215<<(uint(v217)%64)) | (int64(base.Ui64(v203)>>(uint(v217)%64))&v215 | int64(base.Ui64(v203)>>(uint(v213)%64))&v211 | (int64(base.Ui64(v203)>>(uint(v208)%64))&v206 | int64(base.Ui64(v203)>>(uint(v204)%64)))) + v240
	goto L64
L60:
	;
	v199 = *(*int32)(unsafe.Add(mBase, uint32(v197)))
	goto L61
L61:
	;
	if v192 == v199 {
		goto L59
	} else {
		goto L62
	}
L62:
	;
	F_sdsfree(m, v192)
	mBase = m.M
	v202 = m.ExcPending
	if v202 != 0 {
		goto L12
	} else {
		goto L63
	}
L63:
	;
	goto L59
L64:
	;
	F_clusterDelNode(m, v178)
	mBase = m.M
	v244 = m.ExcPending
	if v244 != 0 {
		goto L12
	} else {
		goto L65
	}
L65:
	;
	F_clearCachedClusterSlotsResponse(m)
	mBase = m.M
	v246 = m.ExcPending
	if v246 != 0 {
		goto L12
	} else {
		goto L66
	}
L66:
	;
	v248 = *(*int32)(unsafe.Add(mBase, _consts[90]))
	v249 = *(*int32)(unsafe.Add(mBase, uint32(v248)+uint32(_consts[101])))
	*(*int32)(unsafe.Add(mBase, uint32(v248)+uint32(_consts[101]))) = v249 | int32(6)
	v267 = v105
	v268 = v106
	v269 = v107
	v270 = v108
	v271 = v109
	v272 = v110
	v273 = v111
	v274 = v112
	goto L25
L67:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19))) = v117
	F__serverLog(m, int32(3), int32(_a240), v19)
	mBase = m.M
	v265 = m.ExcPending
	if v265 != 0 {
		goto L12
	} else {
		goto L68
	}
L68:
	;
	v267 = v105
	v268 = v106
	v269 = v107
	v270 = v108
	v271 = v109
	v272 = v110
	v273 = v111
	v274 = v112
	goto L25
L69:
	;
	v282 = v101 + int32(-1)
	if v282&int32(65535) != 0 {
		v100 = v100 + v279
		v101 = v282
		v105 = v267
		v106 = v268
		v107 = v269
		v108 = v270
		v109 = v271
		v110 = v272
		v111 = v273
		v112 = v274
		goto L23
	} else {
		goto L70
	}
L70:
	;
	goto L24
L71:
	;
	F_updateSdsExtensionField(m, v73+int32(2316), v296)
	mBase = m.M
	v308 = m.ExcPending
	if v308 != 0 {
		goto L12
	} else {
		goto L72
	}
L72:
	;
	F_updateSdsExtensionField(m, v73+int32(2304), v294)
	mBase = m.M
	v312 = m.ExcPending
	if v312 != 0 {
		goto L12
	} else {
		goto L73
	}
L73:
	;
	F_updateSdsExtensionField(m, v73+int32(2308), v293)
	mBase = m.M
	v316 = m.ExcPending
	if v316 != 0 {
		goto L12
	} else {
		goto L74
	}
L74:
	;
	v317 = *(*int32)(unsafe.Add(mBase, uint32(v73)+2336))
	if v292 == v317 {
		goto L75
	} else {
		goto L76
	}
L75:
	;
	v329 = *(*int32)(unsafe.Add(mBase, uint32(v73)+2340))
	if v291 == v329 {
		goto L78
	} else {
		goto L79
	}
L76:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v73)+2336)) = v292
	F_clearCachedClusterSlotsResponse(m)
	mBase = m.M
	v321 = m.ExcPending
	if v321 != 0 {
		goto L12
	} else {
		goto L77
	}
L77:
	;
	v323 = *(*int32)(unsafe.Add(mBase, _consts[90]))
	v324 = *(*int32)(unsafe.Add(mBase, uint32(v323)+uint32(_consts[101])))
	*(*int32)(unsafe.Add(mBase, uint32(v323)+uint32(_consts[101]))) = v324 | int32(4)
	goto L75
L78:
	;
	F_updateSdsExtensionField(m, v73+int32(2320), v295)
	mBase = m.M
	v344 = m.ExcPending
	if v344 != 0 {
		goto L12
	} else {
		goto L81
	}
L79:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v73)+2340)) = v291
	F_clearCachedClusterSlotsResponse(m)
	mBase = m.M
	v333 = m.ExcPending
	if v333 != 0 {
		goto L12
	} else {
		goto L80
	}
L80:
	;
	v335 = *(*int32)(unsafe.Add(mBase, _consts[90]))
	v336 = *(*int32)(unsafe.Add(mBase, uint32(v335)+uint32(_consts[101])))
	*(*int32)(unsafe.Add(mBase, uint32(v335)+uint32(_consts[101]))) = v336 | int32(4)
	goto L78
L81:
	;
	if v290 != 0 {
		v378 = v290
		goto L83
	} else {
		goto L84
	}
L82:
	;
	F__serverAssert(m, int32(_a241), int32(_a179), int32(7623))
	mBase = m.M
	v398 = m.ExcPending
	if v398 != 0 {
		goto L12
	} else {
		goto L94
	}
L83:
	;
	F_updateShardId(m, v73, v378)
	mBase = m.M
	v390 = m.ExcPending
	if v390 != 0 {
		goto L12
	} else {
		goto L93
	}
L84:
	;
	v346 = v73
	goto L86
L85:
	;
	v367 = *(*int32)(unsafe.Add(mBase, _consts[122]))
	if v367 == int32(0) {
		goto L90
	} else {
		goto L91
	}
L86:
	;
	v361 = *(*int32)(unsafe.Add(mBase, uint32(v346)+2172))
	if v361 == int32(0) {
		v365 = v346
		goto L85
	} else {
		goto L88
	}
L87:
	;
	v365 = v361
	goto L85
L88:
	;
	if v361 != v73 {
		v346 = v361
		goto L86
	} else {
		goto L89
	}
L89:
	;
	goto L87
L90:
	;
	v378 = v365 + int32(48)
	goto L83
L91:
	;
	v370 = *(*int32)(unsafe.Add(mBase, uint32(v365)+2172))
	if v370 != 0 {
		goto L82
	} else {
		goto L92
	}
L92:
	;
	goto L90
L93:
	;
	m.G0 = v19 + int32(16)
	return
L94:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_clusterProcessPublishPacket(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int64
	_ = v12
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v23 int32
	_ = v23
	var v24 int64
	_ = v24
	var v28 int32
	_ = v28
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
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	switch l1 + int32(-4) {
	case 0:
		v10 = int32(_a69)
		v11 = *(*int32)(unsafe.Add(mBase, _consts[131]))
		v12 = F_kvstoreSize(m, v11)
		mBase = m.M
		v14 = *(*int32)(unsafe.Add(mBase, _consts[132]))
		v15 = *(*int32)(unsafe.Add(mBase, uint32(v14)+12))
		v18 = *(*int32)(unsafe.Add(mBase, uint32(v14)+16))
		if int32(0) < v15+base.I32_wrap_i64(v12)+v18 {
			v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			v29 = F___bswap_32_2(m, v28)
			mBase = m.M
			v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			v31 = F___bswap_32_2(m, v30)
			mBase = m.M
			v33 = l0 + int32(8)
			v34 = F_createStringObject_1(m, v33, v29)
			mBase = m.M
			v35 = m.ExcPending
			if v35 != 0 {
				return
			} else {
				v37 = F_createStringObject_1(m, v33+v29, v31)
				mBase = m.M
				v38 = m.ExcPending
				if v38 != 0 {
					return
				} else {
					v41 = F_pubsubPublishMessage(m, v34, v37, base.B2i32(l1 == int32(10)))
					mBase = m.M
					v42 = m.ExcPending
					if v42 != 0 {
						return
					} else {
						F_decrRefCount(m, v34)
						mBase = m.M
						v44 = m.ExcPending
						if v44 != 0 {
							return
						} else {
							F_decrRefCount(m, v37)
							mBase = m.M
							v46 = m.ExcPending
							if v46 != 0 {
								return
							} else {
								return
							}
						}
					}
				}
			}
		} else {
			return
		}
	default:
		return
	case 6:
		v23 = *(*int32)(unsafe.Add(mBase, _consts[120]))
		v24 = F_kvstoreSize(m, v23)
		mBase = m.M
		if base.I32_wrap_i64(v24) < int32(1) {
			return
		} else {
			v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			v29 = F___bswap_32_2(m, v28)
			mBase = m.M
			v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			v31 = F___bswap_32_2(m, v30)
			mBase = m.M
			v33 = l0 + int32(8)
			v34 = F_createStringObject_1(m, v33, v29)
			mBase = m.M
			v35 = m.ExcPending
			if v35 != 0 {
				return
			} else {
				v37 = F_createStringObject_1(m, v33+v29, v31)
				mBase = m.M
				v38 = m.ExcPending
				if v38 != 0 {
					return
				} else {
					v41 = F_pubsubPublishMessage(m, v34, v37, base.B2i32(l1 == int32(10)))
					mBase = m.M
					v42 = m.ExcPending
					if v42 != 0 {
						return
					} else {
						F_decrRefCount(m, v34)
						mBase = m.M
						v44 = m.ExcPending
						if v44 != 0 {
							return
						} else {
							F_decrRefCount(m, v37)
							mBase = m.M
							v46 = m.ExcPending
							if v46 != 0 {
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
func F_clusterPropagatePublish(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v71 int32
	_ = v71
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v89 int64
	_ = v89
	var v90 int64
	_ = v90
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
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
	var v106 int32
	_ = v106
	var v111 int32
	_ = v111
	var v113 int32
	_ = v113
	var v117 int32
	_ = v117
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v124 int32
	_ = v124
	var v127 int32
	_ = v127
	var v129 int32
	_ = v129
	var v133 int32
	_ = v133
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v140 int32
	_ = v140
	var v149 int32
	_ = v149
	var v155 int32
	_ = v155
	var v161 int32
	_ = v161
	v10 = m.G0
	v12 = v10 - int32(48)
	m.G0 = v12
	v15 = F_clusterCreatePublishMsgBlock(m, l0, l1, int32(1), l2)
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	if l2 == int32(0) {
		goto L7
	} else {
		goto L8
	}
L3:
	;
	F__serverAssert(m, int32(_a233), int32(_a179), int32(1756))
	mBase = m.M
	v161 = m.ExcPending
	if v161 != 0 {
		goto L1
	} else {
		goto L47
	}
L4:
	;
	F__serverAssert(m, int32(_a233), int32(_a179), int32(1756))
	mBase = m.M
	v155 = m.ExcPending
	if v155 != 0 {
		goto L1
	} else {
		goto L46
	}
L5:
	;
	F__serverAssert(m, int32(_a290), int32(_a179), int32(320))
	mBase = m.M
	v149 = m.ExcPending
	if v149 != 0 {
		goto L1
	} else {
		goto L45
	}
L6:
	;
	v71 = int32(0)
	goto L18
L7:
	;
	v47 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v12)+8)) = v47
	v50 = v12 + int32(16)
	v52 = *(*int32)(unsafe.Add(mBase, _consts[90]))
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v52)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v50))) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v50)+20)) = v47
	*(*int64)(unsafe.Add(mBase, uint32(v50)+4)) = int64(4294967295)
	*(*int64)(unsafe.Add(mBase, uint32(v50)+12)) = int64(1)
	goto L16
L8:
	;
	v20 = *(*int32)(unsafe.Add(mBase, _consts[90]))
	v21 = *(*int32)(unsafe.Add(mBase, uint32(v20)))
	v25 = F_sdsnewlen(m, v21+int32(48), int32(40))
	mBase = m.M
	v26 = m.ExcPending
	if v26 != 0 {
		goto L1
	} else {
		goto L9
	}
L9:
	;
	v28 = *(*int32)(unsafe.Add(mBase, _consts[90]))
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v28)+36))
	v30 = F_dictFind(m, v29, v25)
	mBase = m.M
	v31 = m.ExcPending
	if v31 != 0 {
		goto L1
	} else {
		goto L10
	}
L10:
	;
	F_sdsfree(m, v25)
	mBase = m.M
	v33 = m.ExcPending
	if v33 != 0 {
		goto L1
	} else {
		goto L11
	}
L11:
	;
	if v30 == int32(0) {
		goto L5
	} else {
		goto L12
	}
L12:
	;
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v30)+8))
	goto L13
L13:
	;
	if v36 == int32(0) {
		goto L5
	} else {
		goto L14
	}
L14:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+8)) = int32(1)
	v42 = v12 + int32(16)
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v36)))
	*(*int32)(unsafe.Add(mBase, uint32(v42)+4)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v42))) = v43
	goto L15
L15:
	;
	goto L6
L16:
	;
	goto L6
L17:
	;
	v102 = *(*int32)(unsafe.Add(mBase, uint32(v12)+8))
	switch v102 {
	case 0:
		goto L34
	default:
		goto L32
	case 2:
		goto L33
	}
L18:
	;
	v75 = F_clusterNodeIterNext(m, v12+int32(8))
	mBase = m.M
	v76 = m.ExcPending
	if v76 != 0 {
		goto L1
	} else {
		goto L20
	}
L20:
	;
	if v75 == int32(0) {
		goto L17
	} else {
		goto L21
	}
L21:
	;
	v79 = *(*int32)(unsafe.Add(mBase, uint32(v75)+88))
	if v79&int32(48) != 0 {
		goto L18
	} else {
		goto L22
	}
L22:
	;
	v82 = *(*int32)(unsafe.Add(mBase, uint32(v75)+2344))
	if v82 == int32(0) {
		goto L23
	} else {
		goto L24
	}
L23:
	;
	if v71 != 0 {
		v98 = v71
		v99 = v82
		goto L28
	} else {
		goto L29
	}
L24:
	;
	if v79&int32(2048) == int32(0) {
		goto L23
	} else {
		goto L25
	}
L25:
	;
	v89 = *(*int64)(unsafe.Add(mBase, uint32(v75)+2192))
	v90 = *(*int64)(unsafe.Add(mBase, uint32(v82)))
	if v89 < v90 {
		goto L23
	} else {
		goto L26
	}
L26:
	;
	F_clusterSendMessage(m, v82, v15)
	mBase = m.M
	v93 = m.ExcPending
	if v93 != 0 {
		goto L1
	} else {
		goto L27
	}
L27:
	;
	goto L18
L28:
	;
	F_clusterSendMessage(m, v99, v98)
	mBase = m.M
	v101 = m.ExcPending
	if v101 != 0 {
		goto L1
	} else {
		goto L31
	}
L29:
	;
	v95 = F_clusterCreatePublishMsgBlock(m, l0, l1, int32(0), l2)
	mBase = m.M
	v96 = m.ExcPending
	if v96 != 0 {
		goto L1
	} else {
		goto L30
	}
L30:
	;
	v97 = *(*int32)(unsafe.Add(mBase, uint32(v75)+2344))
	v98 = v95
	v99 = v97
	goto L28
L31:
	;
	v71 = v98
	goto L18
L32:
	;
	if v71 == int32(0) {
		goto L36
	} else {
		goto L37
	}
L33:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = int32(0)
	goto L32
L34:
	;
	F_dictResetIterator(m, v12+int32(16))
	mBase = m.M
	v106 = m.ExcPending
	if v106 != 0 {
		goto L1
	} else {
		goto L35
	}
L35:
	;
	goto L32
L36:
	;
	v127 = *(*int32)(unsafe.Add(mBase, uint32(v15)+4))
	v129 = v127 + int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(v15)+4)) = v129
	if v127 <= int32(0) {
		goto L3
	} else {
		goto L41
	}
L37:
	;
	v111 = *(*int32)(unsafe.Add(mBase, uint32(v71)+4))
	v113 = v111 + int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(v71)+4)) = v113
	if v111 <= int32(0) {
		goto L4
	} else {
		goto L38
	}
L38:
	;
	if v113 != 0 {
		goto L36
	} else {
		goto L39
	}
L39:
	;
	v117 = int32(_a69)
	v119 = *(*int32)(unsafe.Add(mBase, _consts[118]))
	v120 = *(*int32)(unsafe.Add(mBase, uint32(v71)))
	*(*int32)(unsafe.Add(mBase, _consts[118])) = v119 - v120
	F_valkey_free(m, v71)
	mBase = m.M
	v124 = m.ExcPending
	if v124 != 0 {
		goto L1
	} else {
		goto L40
	}
L40:
	;
	goto L36
L41:
	;
	if v129 != 0 {
		goto L42
	} else {
		goto L43
	}
L42:
	;
	m.G0 = v12 + int32(48)
	return
L43:
	;
	v133 = int32(_a69)
	v135 = *(*int32)(unsafe.Add(mBase, _consts[118]))
	v136 = *(*int32)(unsafe.Add(mBase, uint32(v15)))
	*(*int32)(unsafe.Add(mBase, _consts[118])) = v135 - v136
	F_valkey_free(m, v15)
	mBase = m.M
	v140 = m.ExcPending
	if v140 != 0 {
		goto L1
	} else {
		goto L44
	}
L44:
	;
	goto L42
L45:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L46:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L47:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_clusterRemoveNodeFromShard(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
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
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v32 int32
	_ = v32
	v7 = F_sdsnewlen(m, l0+int32(48), int32(40))
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		return
	} else {
		v10 = *(*int32)(unsafe.Add(mBase, _consts[90]))
		v11 = *(*int32)(unsafe.Add(mBase, uint32(v10)+36))
		v12 = F_dictFind(m, v11, v7)
		mBase = m.M
		v13 = m.ExcPending
		if v13 != 0 {
			return
		} else {
			if v12 == int32(0) {
				F_sdsfree(m, v7)
				mBase = m.M
				v32 = m.ExcPending
				if v32 != 0 {
					return
				} else {
					return
				}
			} else {
				v16 = *(*int32)(unsafe.Add(mBase, uint32(v12)+8))
				v17 = F_listSearchKey(m, v16, l0)
				mBase = m.M
				v18 = m.ExcPending
				if v18 != 0 {
					return
				} else {
					if v17 == int32(0) {
						v23 = *(*int32)(unsafe.Add(mBase, uint32(v16)+20))
						if v23 != 0 {
							F_sdsfree(m, v7)
							mBase = m.M
							v32 = m.ExcPending
							if v32 != 0 {
								return
							} else {
								return
							}
						} else {
							v25 = *(*int32)(unsafe.Add(mBase, _consts[90]))
							v26 = *(*int32)(unsafe.Add(mBase, uint32(v25)+36))
							v27 = F_dictDelete(m, v26, v7)
							mBase = m.M
							v28 = m.ExcPending
							if v28 != 0 {
								return
							} else {
								F_sdsfree(m, v7)
								mBase = m.M
								v32 = m.ExcPending
								if v32 != 0 {
									return
								} else {
									return
								}
							}
						}
					} else {
						F_listDelNode(m, v16, v17)
						mBase = m.M
						v22 = m.ExcPending
						if v22 != 0 {
							return
						} else {
							v23 = *(*int32)(unsafe.Add(mBase, uint32(v16)+20))
							if v23 != 0 {
								F_sdsfree(m, v7)
								mBase = m.M
								v32 = m.ExcPending
								if v32 != 0 {
									return
								} else {
									return
								}
							} else {
								v25 = *(*int32)(unsafe.Add(mBase, _consts[90]))
								v26 = *(*int32)(unsafe.Add(mBase, uint32(v25)+36))
								v27 = F_dictDelete(m, v26, v7)
								mBase = m.M
								v28 = m.ExcPending
								if v28 != 0 {
									return
								} else {
									F_sdsfree(m, v7)
									mBase = m.M
									v32 = m.ExcPending
									if v32 != 0 {
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
func F_clusterRenameNode(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
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
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v37 int64
	_ = v37
	var v39 int32
	_ = v39
	var v43 int64
	_ = v43
	var v45 int32
	_ = v45
	var v49 int64
	_ = v49
	var v51 int32
	_ = v51
	var v55 int64
	_ = v55
	var v61 int64
	_ = v61
	var v64 int32
	_ = v64
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
	var v82 int32
	_ = v82
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	v12 = l0 + int32(8)
	v14 = F_sdsnewlen(m, v12, int32(40))
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		return
	} else {
		v17 = *(*int32)(unsafe.Add(mBase, _consts[15]))
		if int32(0) < v17 {
			v31 = *(*int32)(unsafe.Add(mBase, _consts[90]))
			v32 = *(*int32)(unsafe.Add(mBase, uint32(v31)+32))
			v33 = F_dictDelete(m, v32, v14)
			mBase = m.M
			v34 = m.ExcPending
			if v34 != 0 {
				return
			} else {
				F_sdsfree(m, v14)
				mBase = m.M
				v36 = m.ExcPending
				if v36 != 0 {
					return
				} else {
					if v33 != 0 {
						F__serverAssert(m, int32(_a182), int32(_a179), int32(2301))
						mBase = m.M
						v82 = m.ExcPending
						if v82 != 0 {
							return
						} else {
							F_abort(m)
							mBase = m.M
							base.Wasm_trap_unreachable()
							for {
							}
						}
					} else {
						v37 = *(*int64)(unsafe.Add(mBase, uint32(l1)))
						*(*int64)(unsafe.Add(mBase, uint32(v12))) = v37
						v39 = int32(32)
						v90 = int32(40)
						v43 = *(*int64)(unsafe.Add(mBase, uint32(l1+v39)))
						*(*int64)(unsafe.Add(mBase, uint32(l0+v90))) = v43
						v45 = int32(24)
						v49 = *(*int64)(unsafe.Add(mBase, uint32(l1+v45)))
						*(*int64)(unsafe.Add(mBase, uint32(l0+v39))) = v49
						v51 = int32(16)
						v55 = *(*int64)(unsafe.Add(mBase, uint32(l1+v51)))
						*(*int64)(unsafe.Add(mBase, uint32(l0+v45))) = v55
						v61 = *(*int64)(unsafe.Add(mBase, uint32(l1+int32(8))))
						*(*int64)(unsafe.Add(mBase, uint32(l0+v51))) = v61
						v64 = *(*int32)(unsafe.Add(mBase, _consts[90]))
						v65 = *(*int32)(unsafe.Add(mBase, uint32(v64)+32))
						v67 = F_sdsnewlen(m, v12, v90)
						mBase = m.M
						v68 = m.ExcPending
						if v68 != 0 {
							return
						} else {
							v69 = F_dictAdd(m, v65, v67, l0)
							mBase = m.M
							v70 = m.ExcPending
							if v70 != 0 {
								return
							} else {
								if v69 != 0 {
									F__serverAssert(m, int32(_a182), int32(_a179), int32(2220))
									mBase = m.M
									v88 = m.ExcPending
									if v88 != 0 {
										return
									} else {
										F_abort(m)
										mBase = m.M
										base.Wasm_trap_unreachable()
										for {
										}
									}
								} else {
									F_clusterAddNodeToShard(m, l0+int32(48), l0)
									mBase = m.M
									v74 = m.ExcPending
									if v74 != 0 {
										return
									} else {
										m.G0 = v9 + int32(16)
										return
									}
								}
							}
						}
					}
				}
			}
		} else {
			v20 = F_humanNodename(m, l0)
			mBase = m.M
			v21 = m.ExcPending
			if v21 != 0 {
				return
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v9)+8)) = l1
				*(*int32)(unsafe.Add(mBase, uint32(v9)+4)) = v20
				*(*int32)(unsafe.Add(mBase, uint32(v9))) = v12
				F__serverLog(m, int32(0), int32(_a236), v9)
				mBase = m.M
				v28 = m.ExcPending
				if v28 != 0 {
					return
				} else {
					v31 = *(*int32)(unsafe.Add(mBase, _consts[90]))
					v32 = *(*int32)(unsafe.Add(mBase, uint32(v31)+32))
					v33 = F_dictDelete(m, v32, v14)
					mBase = m.M
					v34 = m.ExcPending
					if v34 != 0 {
						return
					} else {
						F_sdsfree(m, v14)
						mBase = m.M
						v36 = m.ExcPending
						if v36 != 0 {
							return
						} else {
							if v33 != 0 {
								F__serverAssert(m, int32(_a182), int32(_a179), int32(2301))
								mBase = m.M
								v82 = m.ExcPending
								if v82 != 0 {
									return
								} else {
									F_abort(m)
									mBase = m.M
									base.Wasm_trap_unreachable()
									for {
									}
								}
							} else {
								v37 = *(*int64)(unsafe.Add(mBase, uint32(l1)))
								*(*int64)(unsafe.Add(mBase, uint32(v12))) = v37
								v39 = int32(32)
								v90 = int32(40)
								v43 = *(*int64)(unsafe.Add(mBase, uint32(l1+v39)))
								*(*int64)(unsafe.Add(mBase, uint32(l0+v90))) = v43
								v45 = int32(24)
								v49 = *(*int64)(unsafe.Add(mBase, uint32(l1+v45)))
								*(*int64)(unsafe.Add(mBase, uint32(l0+v39))) = v49
								v51 = int32(16)
								v55 = *(*int64)(unsafe.Add(mBase, uint32(l1+v51)))
								*(*int64)(unsafe.Add(mBase, uint32(l0+v45))) = v55
								v61 = *(*int64)(unsafe.Add(mBase, uint32(l1+int32(8))))
								*(*int64)(unsafe.Add(mBase, uint32(l0+v51))) = v61
								v64 = *(*int32)(unsafe.Add(mBase, _consts[90]))
								v65 = *(*int32)(unsafe.Add(mBase, uint32(v64)+32))
								v67 = F_sdsnewlen(m, v12, v90)
								mBase = m.M
								v68 = m.ExcPending
								if v68 != 0 {
									return
								} else {
									v69 = F_dictAdd(m, v65, v67, l0)
									mBase = m.M
									v70 = m.ExcPending
									if v70 != 0 {
										return
									} else {
										if v69 != 0 {
											F__serverAssert(m, int32(_a182), int32(_a179), int32(2220))
											mBase = m.M
											v88 = m.ExcPending
											if v88 != 0 {
												return
											} else {
												F_abort(m)
												mBase = m.M
												base.Wasm_trap_unreachable()
												for {
												}
											}
										} else {
											F_clusterAddNodeToShard(m, l0+int32(48), l0)
											mBase = m.M
											v74 = m.ExcPending
											if v74 != 0 {
												return
											} else {
												m.G0 = v9 + int32(16)
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
func F_clusterRequestFailoverAuth(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v11 int64
	_ = v11
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v20 int64
	_ = v20
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v51 int64
	_ = v51
	var v52 int64
	_ = v52
	var v54 int64
	_ = v54
	var v57 int64
	_ = v57
	var v59 int64
	_ = v59
	var v60 int64
	_ = v60
	var v61 int64
	_ = v61
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
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v94 int32
	_ = v94
	var v99 int32
	_ = v99
	v7 = F_createClusterMsgSendBlock(m, int32(5), int32(2256))
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		return
	} else {
		v10 = *(*int32)(unsafe.Add(mBase, _consts[90]))
		v11 = *(*int64)(unsafe.Add(mBase, uint32(v10)+uint32(_consts[113])))
		if v11 == int64(0) {
			v23 = int32(0)
			v24 = *(*int32)(unsafe.Add(mBase, _consts[94]))
			v25 = *(*int32)(unsafe.Add(mBase, uint32(v24)+88))
			if v25&int32(16) == v23 {
				v60 = *(*int64)(unsafe.Add(mBase, uint32(v24)+2248))
				v61 = v60
			} else {
				if v25&int32(2) == int32(0) {
					v59 = *(*int64)(unsafe.Add(mBase, _consts[31]))
					v61 = v59
				} else {
					v38 = *(*int32)(unsafe.Add(mBase, _consts[64]))
					if v38 == int32(0) {
						v52 = int64(0)
					} else {
						v42 = *(*int32)(unsafe.Add(mBase, _consts[133]))
						if v42 != 0 {
							v49 = v42
							v50 = *(*int32)(unsafe.Add(mBase, uint32(v49)+104))
							v51 = *(*int64)(unsafe.Add(mBase, uint32(v50)+48))
							v52 = v51
						} else {
							v45 = *(*int32)(unsafe.Add(mBase, _consts[134]))
							if v45 == int32(0) {
								v52 = int64(0)
							} else {
								v49 = v45
								v50 = *(*int32)(unsafe.Add(mBase, uint32(v49)+104))
								v51 = *(*int64)(unsafe.Add(mBase, uint32(v50)+48))
								v52 = v51
							}
						}
					}
					v54 = int64(0)
					if v54 < v52 {
						v57 = v52
					} else {
						v57 = v54
					}
					v61 = v57
				}
			}
			if v61 == int64(0) {
				F_clusterBroadcastMessage(m, v7)
				mBase = m.M
				v80 = m.ExcPending
				if v80 != 0 {
					return
				} else {
					v81 = *(*int32)(unsafe.Add(mBase, uint32(v7)+4))
					v83 = v81 + int32(-1)
					*(*int32)(unsafe.Add(mBase, uint32(v7)+4)) = v83
					if v81 <= int32(0) {
						F__serverAssert(m, int32(_a233), int32(_a179), int32(1756))
						mBase = m.M
						v99 = m.ExcPending
						if v99 != 0 {
							return
						} else {
							F_abort(m)
							mBase = m.M
							base.Wasm_trap_unreachable()
							for {
							}
						}
					} else {
						if v83 != 0 {
							return
						} else {
							v87 = int32(_a69)
							v89 = *(*int32)(unsafe.Add(mBase, _consts[118]))
							v90 = *(*int32)(unsafe.Add(mBase, uint32(v7)))
							*(*int32)(unsafe.Add(mBase, _consts[118])) = v89 - v90
							F_valkey_free(m, v7)
							mBase = m.M
							v94 = m.ExcPending
							if v94 != 0 {
								return
							} else {
								return
							}
						}
					}
				}
			} else {
				v65 = *(*int32)(unsafe.Add(mBase, _consts[90]))
				v66 = *(*int32)(unsafe.Add(mBase, uint32(v65)+uint32(_consts[135])))
				if v66 != 0 {
					F_clusterBroadcastMessage(m, v7)
					mBase = m.M
					v80 = m.ExcPending
					if v80 != 0 {
						return
					} else {
						v81 = *(*int32)(unsafe.Add(mBase, uint32(v7)+4))
						v83 = v81 + int32(-1)
						*(*int32)(unsafe.Add(mBase, uint32(v7)+4)) = v83
						if v81 <= int32(0) {
							F__serverAssert(m, int32(_a233), int32(_a179), int32(1756))
							mBase = m.M
							v99 = m.ExcPending
							if v99 != 0 {
								return
							} else {
								F_abort(m)
								mBase = m.M
								base.Wasm_trap_unreachable()
								for {
								}
							}
						} else {
							if v83 != 0 {
								return
							} else {
								v87 = int32(_a69)
								v89 = *(*int32)(unsafe.Add(mBase, _consts[118]))
								v90 = *(*int32)(unsafe.Add(mBase, uint32(v7)))
								*(*int32)(unsafe.Add(mBase, _consts[118])) = v89 - v90
								F_valkey_free(m, v7)
								mBase = m.M
								v94 = m.ExcPending
								if v94 != 0 {
									return
								} else {
									return
								}
							}
						}
					}
				} else {
					v67 = *(*int32)(unsafe.Add(mBase, uint32(v65)+uint32(_consts[136])))
					if v67 != 0 {
						F_clusterBroadcastMessage(m, v7)
						mBase = m.M
						v80 = m.ExcPending
						if v80 != 0 {
							return
						} else {
							v81 = *(*int32)(unsafe.Add(mBase, uint32(v7)+4))
							v83 = v81 + int32(-1)
							*(*int32)(unsafe.Add(mBase, uint32(v7)+4)) = v83
							if v81 <= int32(0) {
								F__serverAssert(m, int32(_a233), int32(_a179), int32(1756))
								mBase = m.M
								v99 = m.ExcPending
								if v99 != 0 {
									return
								} else {
									F_abort(m)
									mBase = m.M
									base.Wasm_trap_unreachable()
									for {
									}
								}
							} else {
								if v83 != 0 {
									return
								} else {
									v87 = int32(_a69)
									v89 = *(*int32)(unsafe.Add(mBase, _consts[118]))
									v90 = *(*int32)(unsafe.Add(mBase, uint32(v7)))
									*(*int32)(unsafe.Add(mBase, _consts[118])) = v89 - v90
									F_valkey_free(m, v7)
									mBase = m.M
									v94 = m.ExcPending
									if v94 != 0 {
										return
									} else {
										return
									}
								}
							}
						}
					} else {
						v68 = F_clusterAllReplicasThinkPrimaryIsFail(m)
						mBase = m.M
						v69 = m.ExcPending
						if v69 != 0 {
							return
						} else {
							if v68 == int32(0) {
							} else {
								v72 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7)+2261)))
								v74 = v72 | int32(2)
								*(*uint8)(unsafe.Add(mBase, uint32(v7)+2261)) = uint8(v74)
							}
							F_clusterBroadcastMessage(m, v7)
							mBase = m.M
							v80 = m.ExcPending
							if v80 != 0 {
								return
							} else {
								v81 = *(*int32)(unsafe.Add(mBase, uint32(v7)+4))
								v83 = v81 + int32(-1)
								*(*int32)(unsafe.Add(mBase, uint32(v7)+4)) = v83
								if v81 <= int32(0) {
									F__serverAssert(m, int32(_a233), int32(_a179), int32(1756))
									mBase = m.M
									v99 = m.ExcPending
									if v99 != 0 {
										return
									} else {
										F_abort(m)
										mBase = m.M
										base.Wasm_trap_unreachable()
										for {
										}
									}
								} else {
									if v83 != 0 {
										return
									} else {
										v87 = int32(_a69)
										v89 = *(*int32)(unsafe.Add(mBase, _consts[118]))
										v90 = *(*int32)(unsafe.Add(mBase, uint32(v7)))
										*(*int32)(unsafe.Add(mBase, _consts[118])) = v89 - v90
										F_valkey_free(m, v7)
										mBase = m.M
										v94 = m.ExcPending
										if v94 != 0 {
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
		} else {
			v14 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7)+2261)))
			v16 = v14 | int32(2)
			*(*uint8)(unsafe.Add(mBase, uint32(v7)+2261)) = uint8(v16)
			v19 = *(*int32)(unsafe.Add(mBase, _consts[90]))
			v20 = *(*int64)(unsafe.Add(mBase, uint32(v19)+uint32(_consts[113])))
			if v20 != int64(0) {
				F_clusterBroadcastMessage(m, v7)
				mBase = m.M
				v80 = m.ExcPending
				if v80 != 0 {
					return
				} else {
					v81 = *(*int32)(unsafe.Add(mBase, uint32(v7)+4))
					v83 = v81 + int32(-1)
					*(*int32)(unsafe.Add(mBase, uint32(v7)+4)) = v83
					if v81 <= int32(0) {
						F__serverAssert(m, int32(_a233), int32(_a179), int32(1756))
						mBase = m.M
						v99 = m.ExcPending
						if v99 != 0 {
							return
						} else {
							F_abort(m)
							mBase = m.M
							base.Wasm_trap_unreachable()
							for {
							}
						}
					} else {
						if v83 != 0 {
							return
						} else {
							v87 = int32(_a69)
							v89 = *(*int32)(unsafe.Add(mBase, _consts[118]))
							v90 = *(*int32)(unsafe.Add(mBase, uint32(v7)))
							*(*int32)(unsafe.Add(mBase, _consts[118])) = v89 - v90
							F_valkey_free(m, v7)
							mBase = m.M
							v94 = m.ExcPending
							if v94 != 0 {
								return
							} else {
								return
							}
						}
					}
				}
			} else {
				v23 = int32(0)
				v24 = *(*int32)(unsafe.Add(mBase, _consts[94]))
				v25 = *(*int32)(unsafe.Add(mBase, uint32(v24)+88))
				if v25&int32(16) == v23 {
					v60 = *(*int64)(unsafe.Add(mBase, uint32(v24)+2248))
					v61 = v60
				} else {
					if v25&int32(2) == int32(0) {
						v59 = *(*int64)(unsafe.Add(mBase, _consts[31]))
						v61 = v59
					} else {
						v38 = *(*int32)(unsafe.Add(mBase, _consts[64]))
						if v38 == int32(0) {
							v52 = int64(0)
						} else {
							v42 = *(*int32)(unsafe.Add(mBase, _consts[133]))
							if v42 != 0 {
								v49 = v42
								v50 = *(*int32)(unsafe.Add(mBase, uint32(v49)+104))
								v51 = *(*int64)(unsafe.Add(mBase, uint32(v50)+48))
								v52 = v51
							} else {
								v45 = *(*int32)(unsafe.Add(mBase, _consts[134]))
								if v45 == int32(0) {
									v52 = int64(0)
								} else {
									v49 = v45
									v50 = *(*int32)(unsafe.Add(mBase, uint32(v49)+104))
									v51 = *(*int64)(unsafe.Add(mBase, uint32(v50)+48))
									v52 = v51
								}
							}
						}
						v54 = int64(0)
						if v54 < v52 {
							v57 = v52
						} else {
							v57 = v54
						}
						v61 = v57
					}
				}
				if v61 == int64(0) {
					F_clusterBroadcastMessage(m, v7)
					mBase = m.M
					v80 = m.ExcPending
					if v80 != 0 {
						return
					} else {
						v81 = *(*int32)(unsafe.Add(mBase, uint32(v7)+4))
						v83 = v81 + int32(-1)
						*(*int32)(unsafe.Add(mBase, uint32(v7)+4)) = v83
						if v81 <= int32(0) {
							F__serverAssert(m, int32(_a233), int32(_a179), int32(1756))
							mBase = m.M
							v99 = m.ExcPending
							if v99 != 0 {
								return
							} else {
								F_abort(m)
								mBase = m.M
								base.Wasm_trap_unreachable()
								for {
								}
							}
						} else {
							if v83 != 0 {
								return
							} else {
								v87 = int32(_a69)
								v89 = *(*int32)(unsafe.Add(mBase, _consts[118]))
								v90 = *(*int32)(unsafe.Add(mBase, uint32(v7)))
								*(*int32)(unsafe.Add(mBase, _consts[118])) = v89 - v90
								F_valkey_free(m, v7)
								mBase = m.M
								v94 = m.ExcPending
								if v94 != 0 {
									return
								} else {
									return
								}
							}
						}
					}
				} else {
					v65 = *(*int32)(unsafe.Add(mBase, _consts[90]))
					v66 = *(*int32)(unsafe.Add(mBase, uint32(v65)+uint32(_consts[135])))
					if v66 != 0 {
						F_clusterBroadcastMessage(m, v7)
						mBase = m.M
						v80 = m.ExcPending
						if v80 != 0 {
							return
						} else {
							v81 = *(*int32)(unsafe.Add(mBase, uint32(v7)+4))
							v83 = v81 + int32(-1)
							*(*int32)(unsafe.Add(mBase, uint32(v7)+4)) = v83
							if v81 <= int32(0) {
								F__serverAssert(m, int32(_a233), int32(_a179), int32(1756))
								mBase = m.M
								v99 = m.ExcPending
								if v99 != 0 {
									return
								} else {
									F_abort(m)
									mBase = m.M
									base.Wasm_trap_unreachable()
									for {
									}
								}
							} else {
								if v83 != 0 {
									return
								} else {
									v87 = int32(_a69)
									v89 = *(*int32)(unsafe.Add(mBase, _consts[118]))
									v90 = *(*int32)(unsafe.Add(mBase, uint32(v7)))
									*(*int32)(unsafe.Add(mBase, _consts[118])) = v89 - v90
									F_valkey_free(m, v7)
									mBase = m.M
									v94 = m.ExcPending
									if v94 != 0 {
										return
									} else {
										return
									}
								}
							}
						}
					} else {
						v67 = *(*int32)(unsafe.Add(mBase, uint32(v65)+uint32(_consts[136])))
						if v67 != 0 {
							F_clusterBroadcastMessage(m, v7)
							mBase = m.M
							v80 = m.ExcPending
							if v80 != 0 {
								return
							} else {
								v81 = *(*int32)(unsafe.Add(mBase, uint32(v7)+4))
								v83 = v81 + int32(-1)
								*(*int32)(unsafe.Add(mBase, uint32(v7)+4)) = v83
								if v81 <= int32(0) {
									F__serverAssert(m, int32(_a233), int32(_a179), int32(1756))
									mBase = m.M
									v99 = m.ExcPending
									if v99 != 0 {
										return
									} else {
										F_abort(m)
										mBase = m.M
										base.Wasm_trap_unreachable()
										for {
										}
									}
								} else {
									if v83 != 0 {
										return
									} else {
										v87 = int32(_a69)
										v89 = *(*int32)(unsafe.Add(mBase, _consts[118]))
										v90 = *(*int32)(unsafe.Add(mBase, uint32(v7)))
										*(*int32)(unsafe.Add(mBase, _consts[118])) = v89 - v90
										F_valkey_free(m, v7)
										mBase = m.M
										v94 = m.ExcPending
										if v94 != 0 {
											return
										} else {
											return
										}
									}
								}
							}
						} else {
							v68 = F_clusterAllReplicasThinkPrimaryIsFail(m)
							mBase = m.M
							v69 = m.ExcPending
							if v69 != 0 {
								return
							} else {
								if v68 == int32(0) {
								} else {
									v72 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7)+2261)))
									v74 = v72 | int32(2)
									*(*uint8)(unsafe.Add(mBase, uint32(v7)+2261)) = uint8(v74)
								}
								F_clusterBroadcastMessage(m, v7)
								mBase = m.M
								v80 = m.ExcPending
								if v80 != 0 {
									return
								} else {
									v81 = *(*int32)(unsafe.Add(mBase, uint32(v7)+4))
									v83 = v81 + int32(-1)
									*(*int32)(unsafe.Add(mBase, uint32(v7)+4)) = v83
									if v81 <= int32(0) {
										F__serverAssert(m, int32(_a233), int32(_a179), int32(1756))
										mBase = m.M
										v99 = m.ExcPending
										if v99 != 0 {
											return
										} else {
											F_abort(m)
											mBase = m.M
											base.Wasm_trap_unreachable()
											for {
											}
										}
									} else {
										if v83 != 0 {
											return
										} else {
											v87 = int32(_a69)
											v89 = *(*int32)(unsafe.Add(mBase, _consts[118]))
											v90 = *(*int32)(unsafe.Add(mBase, uint32(v7)))
											*(*int32)(unsafe.Add(mBase, _consts[118])) = v89 - v90
											F_valkey_free(m, v7)
											mBase = m.M
											v94 = m.ExcPending
											if v94 != 0 {
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
func F_clusterSendFailoverAuthIfNeeded(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v26 int32
	_ = v26
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v45 int32
	_ = v45
	var v46 int64
	_ = v46
	var v47 int64
	_ = v47
	var v49 int64
	_ = v49
	var v51 int64
	_ = v51
	var v54 int64
	_ = v54
	var v56 int64
	_ = v56
	var v58 int64
	_ = v58
	var v60 int64
	_ = v60
	var v81 int64
	_ = v81
	var v82 int64
	_ = v82
	var v85 int32
	_ = v85
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v93 int32
	_ = v93
	var v94 int64
	_ = v94
	var v106 int32
	_ = v106
	var v107 int64
	_ = v107
	var v110 int32
	_ = v110
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v116 int32
	_ = v116
	var v117 int64
	_ = v117
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v132 int32
	_ = v132
	var v135 int64
	_ = v135
	var v136 int32
	_ = v136
	var v139 int32
	_ = v139
	var v144 int64
	_ = v144
	var v146 int64
	_ = v146
	var v148 int64
	_ = v148
	var v151 int64
	_ = v151
	var v153 int64
	_ = v153
	var v155 int64
	_ = v155
	var v157 int64
	_ = v157
	var v178 int64
	_ = v178
	var v198 int32
	_ = v198
	var v203 int64
	_ = v203
	var v213 int64
	_ = v213
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v228 int32
	_ = v228
	var v234 int32
	_ = v234
	var v240 int64
	_ = v240
	var v244 int64
	_ = v244
	var v264 int32
	_ = v264
	var v269 int32
	_ = v269
	var v271 int32
	_ = v271
	var v272 int32
	_ = v272
	var v277 int32
	_ = v277
	var v279 int32
	_ = v279
	var v282 int32
	_ = v282
	var v283 int32
	_ = v283
	var v285 int32
	_ = v285
	var v286 int64
	_ = v286
	var v297 int32
	_ = v297
	var v301 int32
	_ = v301
	var v304 int32
	_ = v304
	var v305 int32
	_ = v305
	var v316 int32
	_ = v316
	var v319 int32
	_ = v319
	var v320 int32
	_ = v320
	var v325 int32
	_ = v325
	var v328 int32
	_ = v328
	var v329 int32
	_ = v329
	var v340 int32
	_ = v340
	var v342 int32
	_ = v342
	var v345 int32
	_ = v345
	var v346 int32
	_ = v346
	var v357 int32
	_ = v357
	var v359 int32
	_ = v359
	var v362 int32
	_ = v362
	var v363 int32
	_ = v363
	var v364 int64
	_ = v364
	var v374 int32
	_ = v374
	var v381 int32
	_ = v381
	var v383 int32
	_ = v383
	var v386 int32
	_ = v386
	var v387 int32
	_ = v387
	var v388 int32
	_ = v388
	var v389 int32
	_ = v389
	var v401 int32
	_ = v401
	var v405 int32
	_ = v405
	var v407 int32
	_ = v407
	v3 = int32(0)
	v15 = m.G0
	v17 = v15 - int32(176)
	m.G0 = v17
	v20 = *(*int32)(unsafe.Add(mBase, _consts[94]))
	v21 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20)+88)))
	if v21&int32(1) == v3 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	m.G0 = v17 + int32(176)
	return
L2:
	;
	v26 = *(*int32)(unsafe.Add(mBase, uint32(v20)+2160))
	if v26 == int32(0) {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v30 = *(*int32)(unsafe.Add(mBase, _consts[90]))
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v30)+24))
	if v31 != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v46 = *(*int64)(unsafe.Add(mBase, uint32(l1)+16))
	v47 = int64(56)
	v49 = int64(65280)
	v51 = int64(40)
	v54 = int64(16711680)
	v56 = int64(24)
	v58 = int64(4278190080)
	v60 = int64(8)
	v81 = v46<<(uint(v47)%64) | v46&v49<<(uint(v51)%64) | (v46&v54<<(uint(v56)%64) | v46&v58<<(uint(v60)%64)) | (int64(base.Ui64(v46)>>(uint(v60)%64))&v58 | int64(base.Ui64(v46)>>(uint(v56)%64))&v54 | (int64(base.Ui64(v46)>>(uint(v51)%64))&v49 | int64(base.Ui64(v46)>>(uint(v47)%64))))
	v82 = *(*int64)(unsafe.Add(mBase, uint32(v30)+8))
	if base.Ui64(v82) <= base.Ui64(v81) {
		goto L10
	} else {
		goto L11
	}
L5:
	;
	v33 = *(*int32)(unsafe.Add(mBase, _consts[15]))
	if int32(3) < v33 {
		goto L1
	} else {
		goto L6
	}
L6:
	;
	v36 = F_humanNodename(m, l0)
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	return
L8:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+4)) = v36
	*(*int32)(unsafe.Add(mBase, uint32(v17))) = l0 + int32(8)
	F__serverLog(m, int32(3), int32(_a282), v17)
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L7
	} else {
		goto L9
	}
L9:
	;
	goto L1
L10:
	;
	v107 = *(*int64)(unsafe.Add(mBase, uint32(v30)+uint32(_consts[91])))
	if v107 != v82 {
		goto L15
	} else {
		goto L16
	}
L11:
	;
	v85 = *(*int32)(unsafe.Add(mBase, _consts[15]))
	if int32(3) < v85 {
		goto L1
	} else {
		goto L12
	}
L12:
	;
	v88 = F_humanNodename(m, l0)
	mBase = m.M
	v89 = m.ExcPending
	if v89 != 0 {
		goto L7
	} else {
		goto L13
	}
L13:
	;
	v93 = *(*int32)(unsafe.Add(mBase, _consts[90]))
	v94 = *(*int64)(unsafe.Add(mBase, uint32(v93)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v17+int32(32)))) = v94
	*(*int32)(unsafe.Add(mBase, uint32(v17)+16)) = l0 + int32(8)
	*(*int32)(unsafe.Add(mBase, uint32(v17)+20)) = v88
	*(*int64)(unsafe.Add(mBase, uint32(v17)+24)) = v81
	F__serverLog(m, int32(3), int32(_a283), v17+int32(16))
	mBase = m.M
	v106 = m.ExcPending
	if v106 != 0 {
		goto L7
	} else {
		goto L14
	}
L14:
	;
	goto L1
L15:
	;
	v129 = *(*int32)(unsafe.Add(mBase, uint32(l0)+2172))
	v130 = *(*int32)(unsafe.Add(mBase, uint32(l0)+88))
	v132 = v130 & int32(1)
	if v132 != 0 {
		goto L23
	} else {
		goto L24
	}
L16:
	;
	v110 = *(*int32)(unsafe.Add(mBase, _consts[15]))
	if int32(3) < v110 {
		goto L1
	} else {
		goto L17
	}
L17:
	;
	v113 = F_humanNodename(m, l0)
	mBase = m.M
	v114 = m.ExcPending
	if v114 != 0 {
		goto L7
	} else {
		goto L18
	}
L18:
	;
	v116 = *(*int32)(unsafe.Add(mBase, _consts[90]))
	v117 = *(*int64)(unsafe.Add(mBase, uint32(v116)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v17)+48)) = l0 + int32(8)
	*(*int32)(unsafe.Add(mBase, uint32(v17)+52)) = v113
	*(*int64)(unsafe.Add(mBase, uint32(v17)+56)) = v117
	F__serverLog(m, int32(3), int32(_a284), v17+int32(48))
	mBase = m.M
	v128 = m.ExcPending
	if v128 != 0 {
		goto L7
	} else {
		goto L19
	}
L19:
	;
	goto L1
L20:
	;
	v359 = *(*int32)(unsafe.Add(mBase, _consts[15]))
	if int32(3) < v359 {
		goto L59
	} else {
		goto L60
	}
L21:
	;
	v342 = *(*int32)(unsafe.Add(mBase, _consts[15]))
	if int32(3) < v342 {
		goto L1
	} else {
		goto L56
	}
L22:
	;
	if v320&int32(8) != 0 {
		goto L1
	} else {
		goto L52
	}
L23:
	;
	if v132 == int32(0) {
		goto L46
	} else {
		goto L47
	}
L24:
	;
	if v129 == int32(0) {
		goto L23
	} else {
		goto L25
	}
L25:
	;
	v135 = *(*int64)(unsafe.Add(mBase, uint32(l1)+24))
	v136 = *(*int32)(unsafe.Add(mBase, uint32(v129)+88))
	if v136&int32(8) != 0 {
		goto L26
	} else {
		goto L27
	}
L26:
	;
	v144 = int64(56)
	v146 = int64(65280)
	v148 = int64(40)
	v151 = int64(16711680)
	v153 = int64(24)
	v155 = int64(4278190080)
	v157 = int64(8)
	v178 = v135<<(uint(v144)%64) | v135&v146<<(uint(v148)%64) | (v135&v151<<(uint(v153)%64) | v135&v155<<(uint(v157)%64)) | (int64(base.Ui64(v135)>>(uint(v157)%64))&v155 | int64(base.Ui64(v135)>>(uint(v153)%64))&v151 | (int64(base.Ui64(v135)>>(uint(v148)%64))&v146 | int64(base.Ui64(v135)>>(uint(v144)%64))))
	v198 = int32(0)
	goto L29
L27:
	;
	v139 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+2253)))
	if v139&int32(2) == int32(0) {
		v320 = v136
		goto L22
	} else {
		goto L28
	}
L28:
	;
	goto L26
L29:
	;
	v203 = *(*int64)(unsafe.Add(mBase, uint32(l1+int32(80)+v198<<(uint(int32(3))%32))))
	if v203 == int64(0) {
		goto L31
	} else {
		goto L32
	}
L30:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v30)+uint32(_consts[91]))) = v82
	F_clearCachedClusterSlotsResponse(m)
	mBase = m.M
	v269 = m.ExcPending
	if v269 != 0 {
		goto L7
	} else {
		goto L41
	}
L31:
	;
	v264 = v198 + int32(1)
	if v264 != int32(256) {
		v198 = v264
		goto L29
	} else {
		goto L40
	}
L32:
	;
	v213 = v203
	goto L33
L33:
	;
	v223 = base.I32_wrap_i64(base.I64_ctz(v213))
	v224 = v198<<(uint(int32(6))%32) | v223
	v228 = *(*int32)(unsafe.Add(mBase, uint32(v30+int32(52)+v224<<(uint(int32(2))%32))))
	if v228 == int32(0) {
		goto L35
	} else {
		goto L36
	}
L34:
	;
	goto L31
L35:
	;
	v244 = (v213 + int64(-1)) & v213
	if base.B2i32(v244 == int64(0)) == int32(0) {
		v213 = v244
		goto L33
	} else {
		goto L39
	}
L36:
	;
	v234 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v30+int32(65920)+int32(base.Ui32(v224)>>(uint(int32(3))%32))))))
	if int32(base.Ui32(v234)>>(uint(v223&int32(7))%32))&int32(1) != 0 {
		goto L35
	} else {
		goto L37
	}
L37:
	;
	v240 = *(*int64)(unsafe.Add(mBase, uint32(v228)+96))
	if base.Ui64(v178) < base.Ui64(v240) {
		goto L20
	} else {
		goto L38
	}
L38:
	;
	goto L35
L39:
	;
	goto L34
L40:
	;
	goto L30
L41:
	;
	v271 = *(*int32)(unsafe.Add(mBase, _consts[90]))
	v272 = *(*int32)(unsafe.Add(mBase, uint32(v271)+uint32(_consts[101])))
	*(*int32)(unsafe.Add(mBase, uint32(v271)+uint32(_consts[101]))) = v272 | int32(12)
	F_clusterSendFailoverAuth(m, l0)
	mBase = m.M
	v277 = m.ExcPending
	if v277 != 0 {
		goto L7
	} else {
		goto L42
	}
L42:
	;
	v279 = *(*int32)(unsafe.Add(mBase, _consts[15]))
	if int32(2) < v279 {
		goto L1
	} else {
		goto L43
	}
L43:
	;
	v282 = F_humanNodename(m, l0)
	mBase = m.M
	v283 = m.ExcPending
	if v283 != 0 {
		goto L7
	} else {
		goto L44
	}
L44:
	;
	v285 = *(*int32)(unsafe.Add(mBase, _consts[90]))
	v286 = *(*int64)(unsafe.Add(mBase, uint32(v285)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v17)+112)) = l0 + int32(8)
	*(*int32)(unsafe.Add(mBase, uint32(v17)+116)) = v282
	*(*int64)(unsafe.Add(mBase, uint32(v17)+120)) = v286
	F__serverLog(m, int32(2), int32(_a285), v17+int32(112))
	mBase = m.M
	v297 = m.ExcPending
	if v297 != 0 {
		goto L7
	} else {
		goto L45
	}
L45:
	;
	goto L1
L46:
	;
	if v129 == int32(0) {
		goto L21
	} else {
		goto L51
	}
L47:
	;
	v301 = *(*int32)(unsafe.Add(mBase, _consts[15]))
	if int32(3) < v301 {
		goto L1
	} else {
		goto L48
	}
L48:
	;
	v304 = F_humanNodename(m, l0)
	mBase = m.M
	v305 = m.ExcPending
	if v305 != 0 {
		goto L7
	} else {
		goto L49
	}
L49:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v17)+104)) = v81
	*(*int32)(unsafe.Add(mBase, uint32(v17)+100)) = v304
	*(*int32)(unsafe.Add(mBase, uint32(v17)+96)) = l0 + int32(8)
	F__serverLog(m, int32(3), int32(_a286), v17+int32(96))
	mBase = m.M
	v316 = m.ExcPending
	if v316 != 0 {
		goto L7
	} else {
		goto L50
	}
L50:
	;
	goto L1
L51:
	;
	v319 = *(*int32)(unsafe.Add(mBase, uint32(v129)+88))
	v320 = v319
	goto L22
L52:
	;
	v325 = *(*int32)(unsafe.Add(mBase, _consts[15]))
	if int32(3) < v325 {
		goto L1
	} else {
		goto L53
	}
L53:
	;
	v328 = F_humanNodename(m, l0)
	mBase = m.M
	v329 = m.ExcPending
	if v329 != 0 {
		goto L7
	} else {
		goto L54
	}
L54:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v17)+88)) = v81
	*(*int32)(unsafe.Add(mBase, uint32(v17)+84)) = v328
	*(*int32)(unsafe.Add(mBase, uint32(v17)+80)) = l0 + int32(8)
	F__serverLog(m, int32(3), int32(_a287), v17+int32(80))
	mBase = m.M
	v340 = m.ExcPending
	if v340 != 0 {
		goto L7
	} else {
		goto L55
	}
L55:
	;
	goto L1
L56:
	;
	v345 = F_humanNodename(m, l0)
	mBase = m.M
	v346 = m.ExcPending
	if v346 != 0 {
		goto L7
	} else {
		goto L57
	}
L57:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v17)+72)) = v81
	*(*int32)(unsafe.Add(mBase, uint32(v17)+68)) = v345
	*(*int32)(unsafe.Add(mBase, uint32(v17)+64)) = l0 + int32(8)
	F__serverLog(m, int32(3), int32(_a288), v17+int32(64))
	mBase = m.M
	v357 = m.ExcPending
	if v357 != 0 {
		goto L7
	} else {
		goto L58
	}
L58:
	;
	goto L1
L59:
	;
	v405 = *(*int32)(unsafe.Add(mBase, uint32(l0)+2344))
	F_clusterSendUpdate(m, v405, v228)
	mBase = m.M
	v407 = m.ExcPending
	if v407 != 0 {
		goto L7
	} else {
		goto L67
	}
L60:
	;
	v362 = F_humanNodename(m, l0)
	mBase = m.M
	v363 = m.ExcPending
	if v363 != 0 {
		goto L7
	} else {
		goto L61
	}
L61:
	;
	v364 = *(*int64)(unsafe.Add(mBase, uint32(v228)+96))
	*(*int64)(unsafe.Add(mBase, uint32(v17+int32(168)))) = v178
	*(*int64)(unsafe.Add(mBase, uint32(v17+int32(160)))) = v364
	*(*int32)(unsafe.Add(mBase, uint32(v17)+152)) = v224
	*(*int32)(unsafe.Add(mBase, uint32(v17)+148)) = v362
	v374 = l0 + int32(8)
	*(*int32)(unsafe.Add(mBase, uint32(v17)+144)) = v374
	F__serverLog(m, int32(3), int32(_a289), v17+int32(144))
	mBase = m.M
	v381 = m.ExcPending
	if v381 != 0 {
		goto L7
	} else {
		goto L62
	}
L62:
	;
	v383 = *(*int32)(unsafe.Add(mBase, _consts[15]))
	if int32(1) < v383 {
		goto L59
	} else {
		goto L63
	}
L63:
	;
	v386 = F_humanNodename(m, l0)
	mBase = m.M
	v387 = m.ExcPending
	if v387 != 0 {
		goto L7
	} else {
		goto L64
	}
L64:
	;
	v388 = F_humanNodename(m, v228)
	mBase = m.M
	v389 = m.ExcPending
	if v389 != 0 {
		goto L7
	} else {
		goto L65
	}
L65:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+140)) = v388
	*(*int32)(unsafe.Add(mBase, uint32(v17)+136)) = v228 + int32(8)
	*(*int32)(unsafe.Add(mBase, uint32(v17)+132)) = v386
	*(*int32)(unsafe.Add(mBase, uint32(v17)+128)) = v374
	F__serverLog(m, int32(1), int32(_a280), v17+int32(128))
	mBase = m.M
	v401 = m.ExcPending
	if v401 != 0 {
		goto L7
	} else {
		goto L66
	}
L66:
	;
	goto L59
L67:
	;
	goto L1
}
func F_clusterSendModuleMessageToTarget(m *base.Module, l0 int32, l1 int64, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v11 int32
	_ = v11
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
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
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
	var v110 int32
	_ = v110
	var v113 int32
	_ = v113
	var v116 int32
	_ = v116
	var v120 int32
	_ = v120
	var v123 int32
	_ = v123
	var v127 int32
	_ = v127
	v6 = int32(0)
	if l0 == v6 {
		v120 = v6
		goto L2
	} else {
		goto L3
	}
L1:
	;
	return v127
L2:
	;
	F_clusterSendModule(m, v120, l1, l2, l3, l4)
	mBase = m.M
	v123 = m.ExcPending
	if v123 != 0 {
		goto L29
	} else {
		goto L37
	}
L3:
	;
	v11 = int32(-1)
	if l0&int32(3) == int32(0) {
		v33 = l0
		goto L6
	} else {
		goto L7
	}
L4:
	;
	if v66 != int32(40) {
		v99 = int32(-1)
		goto L21
	} else {
		goto L22
	}
L5:
	;
	v66 = v58 - l0
	goto L4
L6:
	;
	v37 = v33
	goto L14
L7:
	;
	v19 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v19 != 0 {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v22 = l0
	goto L10
L9:
	;
	v66 = l0 - l0
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
	if v99 != 0 {
		v127 = v11
		goto L1
	} else {
		goto L28
	}
L21:
	;
	goto L20
L22:
	;
	v74 = int32(0)
	goto L24
L23:
	;
	v99 = int32(0) - v89
	goto L21
L24:
	;
	v76 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0+v74))))
	v79 = int32(255)
	v89 = base.B2i32(base.Ui32((v76+int32(-123))&v79) < base.Ui32(int32(230))) & base.B2i32(base.Ui32((v76+int32(-58))&v79) < base.Ui32(int32(246)))
	if v89 != 0 {
		goto L23
	} else {
		goto L26
	}
L25:
	;
	goto L23
L26:
	;
	v91 = v74 + int32(1)
	if v91 != int32(40) {
		v74 = v91
		goto L24
	} else {
		goto L27
	}
L27:
	;
	goto L25
L28:
	;
	v100 = F_sdsnewlen(m, l0, v66)
	mBase = m.M
	v103 = m.ExcPending
	if v103 != 0 {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	return int32(0)
L30:
	;
	v105 = *(*int32)(unsafe.Add(mBase, _consts[90]))
	v106 = *(*int32)(unsafe.Add(mBase, uint32(v105)+32))
	v107 = F_dictFind(m, v106, v100)
	mBase = m.M
	v108 = m.ExcPending
	if v108 != 0 {
		goto L29
	} else {
		goto L31
	}
L31:
	;
	F_sdsfree(m, v100)
	mBase = m.M
	v110 = m.ExcPending
	if v110 != 0 {
		goto L29
	} else {
		goto L32
	}
L32:
	;
	if v107 == int32(0) {
		v127 = v11
		goto L1
	} else {
		goto L33
	}
L33:
	;
	v113 = *(*int32)(unsafe.Add(mBase, uint32(v107)+8))
	goto L34
L34:
	;
	if v113 == int32(0) {
		v127 = v11
		goto L1
	} else {
		goto L35
	}
L35:
	;
	v116 = *(*int32)(unsafe.Add(mBase, uint32(v113)+2344))
	if v116 == int32(0) {
		v127 = v11
		goto L1
	} else {
		goto L36
	}
L36:
	;
	v120 = v116
	goto L2
L37:
	;
	v127 = int32(0)
	goto L1
}
func F_clusterSetPrimary(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v37 int32
	_ = v37
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v58 int32
	_ = v58
	var v62 int32
	_ = v62
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v103 int32
	_ = v103
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
	var v118 int32
	_ = v118
	var v125 int32
	_ = v125
	var v129 int32
	_ = v129
	var v132 int32
	_ = v132
	var v138 int32
	_ = v138
	var v145 int32
	_ = v145
	var v149 int32
	_ = v149
	var v152 int32
	_ = v152
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v166 int32
	_ = v166
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v171 int32
	_ = v171
	var v173 int32
	_ = v173
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v188 int32
	_ = v188
	var v190 int32
	_ = v190
	var v195 int32
	_ = v195
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v210 int32
	_ = v210
	var v212 int32
	_ = v212
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v234 int32
	_ = v234
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v243 int32
	_ = v243
	var v248 int32
	_ = v248
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v251 int32
	_ = v251
	var v252 int32
	_ = v252
	var v259 int32
	_ = v259
	var v268 int32
	_ = v268
	var v271 int32
	_ = v271
	var v287 int32
	_ = v287
	var v288 int32
	_ = v288
	var v290 int32
	_ = v290
	var v296 int32
	_ = v296
	var v301 int32
	_ = v301
	var v302 int32
	_ = v302
	var v321 int32
	_ = v321
	var v322 int32
	_ = v322
	var v324 int32
	_ = v324
	var v327 int32
	_ = v327
	var v329 int32
	_ = v329
	var v330 int32
	_ = v330
	var v333 int64
	_ = v333
	var v334 int32
	_ = v334
	var v335 int32
	_ = v335
	var v337 int32
	_ = v337
	var v340 int64
	_ = v340
	var v343 int32
	_ = v343
	var v346 int32
	_ = v346
	var v349 int32
	_ = v349
	var v352 int32
	_ = v352
	var v353 int32
	_ = v353
	var v358 int32
	_ = v358
	var v365 int32
	_ = v365
	var v369 int32
	_ = v369
	var v372 int32
	_ = v372
	var v373 int32
	_ = v373
	var v377 int32
	_ = v377
	var v379 int32
	_ = v379
	var v380 int32
	_ = v380
	var v384 int32
	_ = v384
	var v386 int32
	_ = v386
	var v400 int32
	_ = v400
	var v401 int32
	_ = v401
	var v406 int32
	_ = v406
	var v408 int32
	_ = v408
	var v409 int32
	_ = v409
	var v410 int32
	_ = v410
	var v419 int32
	_ = v419
	var v421 int32
	_ = v421
	var v423 int32
	_ = v423
	var v424 int64
	_ = v424
	var v433 int32
	_ = v433
	var v439 int32
	_ = v439
	v12 = *(*int32)(unsafe.Add(mBase, _consts[94]))
	if l0 == v12 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	F__serverAssert(m, int32(_a237), int32(_a179), int32(6837))
	mBase = m.M
	v439 = m.ExcPending
	if v439 != 0 {
		goto L61
	} else {
		goto L106
	}
L2:
	;
	F__serverAssert(m, int32(_a238), int32(_a179), int32(6836))
	mBase = m.M
	v433 = m.ExcPending
	if v433 != 0 {
		goto L61
	} else {
		goto L105
	}
L3:
	;
	v14 = *(*int32)(unsafe.Add(mBase, uint32(v12)+2160))
	if v14 != 0 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v15 = *(*int32)(unsafe.Add(mBase, uint32(v12)+88))
	if v15&int32(1) == int32(0) {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	if l1 == int32(0) {
		goto L59
	} else {
		goto L60
	}
L6:
	;
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v12)+2172))
	if v25 == int32(0) {
		goto L5
	} else {
		goto L8
	}
L7:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+88)) = v15&int32(-260) | int32(2)
	goto L5
L8:
	;
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v25)+2164))
	if v28 < int32(1) {
		goto L5
	} else {
		goto L9
	}
L9:
	;
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v25)+2168))
	v37 = int32(0)
	goto L10
L10:
	;
	v44 = v37 + int32(1)
	v47 = v31 + v37<<(uint(int32(2))%32)
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v47)))
	if v48 != v12 {
		goto L12
	} else {
		goto L13
	}
L11:
	;
	goto L5
L12:
	;
	if v44 != v28 {
		v37 = v44
		goto L10
	} else {
		goto L58
	}
L13:
	;
	if v28 <= v44 {
		v208 = v28
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v210 = v208 + int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(v25)+2164)) = v210
	if v210 != 0 {
		goto L5
	} else {
		goto L57
	}
L15:
	;
	v51 = int32(2)
	v53 = v31 + v44<<(uint(v51)%32)
	v58 = (v28 + (v37 ^ int32(-1))) << (uint(v51) % 32)
	if v47 == v53 {
		goto L17
	} else {
		goto L18
	}
L16:
	;
	v207 = *(*int32)(unsafe.Add(mBase, uint32(v25)+2164))
	v208 = v207
	goto L14
L17:
	;
	goto L16
L18:
	;
	v62 = v58 + v47
	if base.Ui32(int32(0)-v58<<(uint(int32(1))%32)) < base.Ui32(v53-v62) {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	v72 = (v53 ^ v47) & int32(3)
	if base.Ui32(v53) <= base.Ui32(v47) {
		goto L23
	} else {
		goto L24
	}
L20:
	;
	v69 = F___memcpy(m, v47, v53, v58)
	mBase = m.M
	goto L16
L21:
	;
	if v178 == int32(0) {
		goto L17
	} else {
		goto L53
	}
L22:
	;
	if base.Ui32(v156) <= base.Ui32(int32(3)) {
		v177 = v155
		v178 = v156
		v179 = v157
		goto L21
	} else {
		goto L49
	}
L23:
	;
	if v72 != 0 {
		v138 = v58
		goto L33
	} else {
		goto L34
	}
L24:
	;
	if v72 == int32(0) {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	if v47&int32(3) != 0 {
		goto L27
	} else {
		goto L28
	}
L26:
	;
	v177 = v53
	v178 = v58
	v179 = v47
	goto L21
L27:
	;
	v79 = v53
	v80 = v58
	v81 = v47
	goto L29
L28:
	;
	v155 = v53
	v156 = v58
	v157 = v47
	goto L22
L29:
	;
	if v80 == int32(0) {
		goto L17
	} else {
		goto L31
	}
L31:
	;
	v85 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v79))))
	*(*uint8)(unsafe.Add(mBase, uint32(v81))) = uint8(v85)
	v87 = int32(1)
	v88 = v79 + v87
	v90 = v80 + int32(-1)
	v92 = v81 + v87
	if v92&int32(3) == int32(0) {
		v155 = v88
		v156 = v90
		v157 = v92
		goto L22
	} else {
		goto L32
	}
L32:
	;
	v79 = v88
	v80 = v90
	v81 = v92
	goto L29
L33:
	;
	if v138 == int32(0) {
		goto L17
	} else {
		goto L45
	}
L34:
	;
	if v62&int32(3) == int32(0) {
		v118 = v58
		goto L35
	} else {
		goto L36
	}
L35:
	;
	if base.Ui32(v118) <= base.Ui32(int32(3)) {
		v138 = v118
		goto L33
	} else {
		goto L41
	}
L36:
	;
	v103 = v58
	goto L37
L37:
	;
	if v103 == int32(0) {
		goto L17
	} else {
		goto L39
	}
L38:
	;
	v118 = v109
	goto L35
L39:
	;
	v109 = v103 + int32(-1)
	v110 = v47 + v109
	v112 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v53+v109))))
	*(*uint8)(unsafe.Add(mBase, uint32(v110))) = uint8(v112)
	if v110&int32(3) != 0 {
		v103 = v109
		goto L37
	} else {
		goto L40
	}
L40:
	;
	goto L38
L41:
	;
	v125 = v118
	goto L42
L42:
	;
	v129 = v125 + int32(-4)
	v132 = *(*int32)(unsafe.Add(mBase, uint32(v53+v129)))
	*(*int32)(unsafe.Add(mBase, uint32(v47+v129))) = v132
	if base.Ui32(int32(3)) < base.Ui32(v129) {
		v125 = v129
		goto L42
	} else {
		goto L44
	}
L43:
	;
	v138 = v129
	goto L33
L44:
	;
	goto L43
L45:
	;
	v145 = v138
	goto L46
L46:
	;
	v149 = v145 + int32(-1)
	v152 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v53+v149))))
	*(*uint8)(unsafe.Add(mBase, uint32(v47+v149))) = uint8(v152)
	if v149 != 0 {
		v145 = v149
		goto L46
	} else {
		goto L48
	}
L48:
	;
	goto L17
L49:
	;
	v162 = v155
	v163 = v156
	v164 = v157
	goto L50
L50:
	;
	v166 = *(*int32)(unsafe.Add(mBase, uint32(v162)))
	*(*int32)(unsafe.Add(mBase, uint32(v164))) = v166
	v168 = int32(4)
	v169 = v162 + v168
	v171 = v164 + v168
	v173 = v163 + int32(-4)
	if base.Ui32(int32(3)) < base.Ui32(v173) {
		v162 = v169
		v163 = v173
		v164 = v171
		goto L50
	} else {
		goto L52
	}
L51:
	;
	v177 = v169
	v178 = v173
	v179 = v171
	goto L21
L52:
	;
	goto L51
L53:
	;
	v184 = v177
	v185 = v178
	v186 = v179
	goto L54
L54:
	;
	v188 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v184))))
	*(*uint8)(unsafe.Add(mBase, uint32(v186))) = uint8(v188)
	v190 = int32(1)
	v195 = v185 + int32(-1)
	if v195 != 0 {
		v184 = v184 + v190
		v185 = v195
		v186 = v186 + v190
		goto L54
	} else {
		goto L56
	}
L55:
	;
	goto L17
L56:
	;
	goto L55
L57:
	;
	v212 = *(*int32)(unsafe.Add(mBase, uint32(v25)+88))
	*(*int32)(unsafe.Add(mBase, uint32(v25)+88)) = v212 & int32(-257)
	goto L5
L58:
	;
	goto L11
L59:
	;
	v241 = int32(0)
	v243 = *(*int32)(unsafe.Add(mBase, _consts[94]))
	*(*int32)(unsafe.Add(mBase, uint32(v243)+2172)) = l0
	F_updateShardId(m, v243, l0+int32(48))
	mBase = m.M
	v248 = m.ExcPending
	if v248 != 0 {
		goto L61
	} else {
		goto L64
	}
L60:
	;
	v230 = *(*int32)(unsafe.Add(mBase, _consts[90]))
	v231 = *(*int32)(unsafe.Add(mBase, uint32(v230)+44))
	F_dictEmpty(m, v231, int32(0))
	mBase = m.M
	v234 = m.ExcPending
	if v234 != 0 {
		goto L61
	} else {
		goto L62
	}
L61:
	;
	return
L62:
	;
	v236 = *(*int32)(unsafe.Add(mBase, _consts[90]))
	v237 = *(*int32)(unsafe.Add(mBase, uint32(v236)+48))
	F_dictEmpty(m, v237, int32(0))
	mBase = m.M
	v240 = m.ExcPending
	if v240 != 0 {
		goto L61
	} else {
		goto L63
	}
L63:
	;
	goto L59
L64:
	;
	v249 = *(*int32)(unsafe.Add(mBase, uint32(l0)+2168))
	v250 = int32(0)
	v251 = *(*int32)(unsafe.Add(mBase, _consts[94]))
	v252 = *(*int32)(unsafe.Add(mBase, uint32(l0)+2164))
	if v252 <= v250 {
		goto L66
	} else {
		goto L67
	}
L65:
	;
	v321 = *(*int32)(unsafe.Add(mBase, _consts[119]))
	if v321 != 0 {
		goto L74
	} else {
		goto L75
	}
L66:
	;
	v287 = F_valkey_realloc(m, v249, v252<<(uint(int32(2))%32)+int32(4))
	mBase = m.M
	v288 = m.ExcPending
	if v288 != 0 {
		goto L61
	} else {
		goto L72
	}
L67:
	;
	v259 = v241
	goto L68
L68:
	;
	v268 = *(*int32)(unsafe.Add(mBase, uint32(v249+v259<<(uint(int32(2))%32))))
	if v268 == v251 {
		goto L65
	} else {
		goto L70
	}
L69:
	;
	goto L66
L70:
	;
	v271 = v259 + int32(1)
	if v271 != v252 {
		v259 = v271
		goto L68
	} else {
		goto L71
	}
L71:
	;
	goto L69
L72:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+2168)) = v287
	v290 = *(*int32)(unsafe.Add(mBase, uint32(l0)+2164))
	*(*int32)(unsafe.Add(mBase, uint32(v287+v290<<(uint(int32(2))%32)))) = v251
	v296 = v290 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+2164)) = v296
	F_qsort(m, v287, v296, int32(4), int32(58))
	mBase = m.M
	v301 = m.ExcPending
	if v301 != 0 {
		goto L61
	} else {
		goto L73
	}
L73:
	;
	v302 = *(*int32)(unsafe.Add(mBase, uint32(l0)+88))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+88)) = v302 | int32(256)
	goto L65
L74:
	;
	v322 = int32(2328)
	goto L76
L75:
	;
	v322 = int32(2324)
	goto L76
L76:
	;
	v324 = *(*int32)(unsafe.Add(mBase, uint32(l0+v322)))
	F_replicationSetPrimary(m, l0+int32(2256), v324, l2, int32(1))
	mBase = m.M
	v327 = m.ExcPending
	if v327 != 0 {
		goto L61
	} else {
		goto L77
	}
L77:
	;
	v329 = *(*int32)(unsafe.Add(mBase, _consts[120]))
	v330 = *(*int32)(unsafe.Add(mBase, uint32(v329)+12))
	if v330 == int32(1) {
		goto L80
	} else {
		goto L81
	}
L78:
	;
	v400 = *(*int32)(unsafe.Add(mBase, _consts[90]))
	v401 = *(*int32)(unsafe.Add(mBase, uint32(v400)+uint32(_consts[111])))
	if v401 == int32(0) {
		v409 = v400
		goto L98
	} else {
		goto L99
	}
L79:
	;
	if v340 == int64(0) {
		goto L78
	} else {
		goto L84
	}
L80:
	;
	v334 = *(*int32)(unsafe.Add(mBase, uint32(v329)+8))
	v335 = *(*int32)(unsafe.Add(mBase, uint32(v334)))
	if v335 != 0 {
		goto L82
	} else {
		goto L83
	}
L81:
	;
	v333 = *(*int64)(unsafe.Add(mBase, uint32(v329)+40))
	v340 = v333
	goto L79
L82:
	;
	v337 = F_hashtableSize(m, v335)
	mBase = m.M
	v340 = base.I64_extend_i32_u(v337)
	goto L79
L83:
	;
	v340 = int64(0)
	goto L79
L84:
	;
	v343 = int32(0)
	v346 = *(*int32)(unsafe.Add(mBase, _consts[94]))
	v349 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v346)+88)))
	if v349&int32(1) != 0 {
		goto L85
	} else {
		goto L86
	}
L85:
	;
	v352 = int32(_a239)
	goto L87
L86:
	;
	v352 = v346 + int32(2172)
	goto L87
L87:
	;
	v353 = *(*int32)(unsafe.Add(mBase, uint32(v352)))
	v358 = v343
	goto L88
L88:
	;
	v365 = *(*int32)(unsafe.Add(mBase, _consts[90]))
	v369 = *(*int32)(unsafe.Add(mBase, uint32(v365+v358<<(uint(int32(2))%32))+52))
	if v369 == v353 {
		goto L90
	} else {
		goto L91
	}
L89:
	;
	goto L78
L90:
	;
	v386 = v358 + int32(1)
	if v386 != int32(16384) {
		v358 = v386
		goto L88
	} else {
		goto L97
	}
L91:
	;
	v372 = *(*int32)(unsafe.Add(mBase, _consts[120]))
	v373 = *(*int32)(unsafe.Add(mBase, uint32(v372)+8))
	v377 = *(*int32)(unsafe.Add(mBase, uint32(v373+v358<<(uint(int32(2))%32))))
	if v377 != 0 {
		goto L93
	} else {
		goto L94
	}
L92:
	;
	if v380 == int32(0) {
		goto L90
	} else {
		goto L95
	}
L93:
	;
	v379 = F_hashtableSize(m, v377)
	mBase = m.M
	v380 = v379
	goto L92
L94:
	;
	v380 = int32(0)
	goto L92
L95:
	;
	F_pubsubShardUnsubscribeAllChannelsInSlot(m, v358)
	mBase = m.M
	v384 = m.ExcPending
	if v384 != 0 {
		goto L61
	} else {
		goto L96
	}
L96:
	;
	goto L90
L97:
	;
	goto L89
L98:
	;
	v410 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v409)+uint32(_consts[112]))) = v410
	*(*int64)(unsafe.Add(mBase, uint32(v409)+uint32(_consts[113]))) = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v409)+uint32(_consts[114]))) = int64(-1)
	*(*int32)(unsafe.Add(mBase, uint32(v409)+uint32(_consts[111]))) = v410
	F_clusterUpdateSlotExportsOnOwnershipChange(m)
	mBase = m.M
	v419 = m.ExcPending
	if v419 != 0 {
		goto L61
	} else {
		goto L101
	}
L99:
	;
	F_unpauseActions(m, int32(2))
	mBase = m.M
	v406 = m.ExcPending
	if v406 != 0 {
		goto L61
	} else {
		goto L100
	}
L100:
	;
	v408 = *(*int32)(unsafe.Add(mBase, _consts[90]))
	v409 = v408
	goto L98
L101:
	;
	F_clusterUpdateSlotImportsOnOwnershipChange(m)
	mBase = m.M
	v421 = m.ExcPending
	if v421 != 0 {
		goto L61
	} else {
		goto L102
	}
L102:
	;
	v423 = *(*int32)(unsafe.Add(mBase, _consts[90]))
	v424 = *(*int64)(unsafe.Add(mBase, uint32(v423)+uint32(_consts[121])))
	if v424 == int64(0) {
		goto L103
	} else {
		goto L104
	}
L103:
	;
	return
L104:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v423)+uint32(_consts[121]))) = int64(0)
	goto L103
L105:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L106:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_clusterSlotStatReset(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v9 int64
	_ = v9
	v3 = *(*int32)(unsafe.Add(mBase, _consts[90]))
	v6 = v3 + l0*int32(24)
	v9 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v6)+uint32(_consts[97]))) = v9
	*(*int64)(unsafe.Add(mBase, uint32(v6)+uint32(_consts[98]))) = v9
	*(*int64)(unsafe.Add(mBase, uint32(v6)+uint32(_consts[99]))) = v9
	return
}
func F_clusterSlotStatsAddNetworkBytesOutForShardedPubSubInternalPropagation(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v10 int32
	_ = v10
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v22 int64
	_ = v22
	var v23 int64
	_ = v23
	var v33 int32
	_ = v33
	if l1 == int32(-1) {
		return
	} else {
		v6 = *(*int32)(unsafe.Add(mBase, _consts[158]))
		if v6 == int32(0) {
			return
		} else {
			v10 = *(*int32)(unsafe.Add(mBase, _consts[63]))
			if v10 == int32(0) {
				return
			} else {
				if base.Ui32(int32(16384)) <= base.Ui32(l1) {
					F__serverAssert(m, int32(_a388), int32(_a389), int32(185))
					mBase = m.M
					v33 = m.ExcPending
					if v33 != 0 {
						return
					} else {
						F_abort(m)
						mBase = m.M
						base.Wasm_trap_unreachable()
						for {
						}
					}
				} else {
					v16 = *(*int32)(unsafe.Add(mBase, _consts[90]))
					v19 = v16 + l1*int32(24)
					v22 = *(*int64)(unsafe.Add(mBase, uint32(v19)+uint32(_consts[97])))
					v23 = *(*int64)(unsafe.Add(mBase, uint32(l0)+264))
					*(*int64)(unsafe.Add(mBase, uint32(v19)+uint32(_consts[97]))) = v22 + v23
					*(*int64)(unsafe.Add(mBase, uint32(l0)+264)) = int64(0)
					return
				}
			}
		}
	}
}
func F_clusterSlotStatsAddNetworkBytesOutForSlot(m *base.Module, l0 int32, l1 int64) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v10 int32
	_ = v10
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v22 int64
	_ = v22
	var v30 int32
	_ = v30
	if l0 == int32(-1) {
		return
	} else {
		v6 = *(*int32)(unsafe.Add(mBase, _consts[158]))
		if v6 == int32(0) {
			return
		} else {
			v10 = *(*int32)(unsafe.Add(mBase, _consts[63]))
			if v10 == int32(0) {
				return
			} else {
				if base.Ui32(int32(16384)) <= base.Ui32(l0) {
					F__serverAssert(m, int32(_a388), int32(_a389), int32(140))
					mBase = m.M
					v30 = m.ExcPending
					if v30 != 0 {
						return
					} else {
						F_abort(m)
						mBase = m.M
						base.Wasm_trap_unreachable()
						for {
						}
					}
				} else {
					v16 = *(*int32)(unsafe.Add(mBase, _consts[90]))
					v19 = v16 + l0*int32(24)
					v22 = *(*int64)(unsafe.Add(mBase, uint32(v19)+uint32(_consts[97])))
					*(*int64)(unsafe.Add(mBase, uint32(v19)+uint32(_consts[97]))) = v22 + l1
					return
				}
			}
		}
	}
}
func F_clusterSlotStatsAddNetworkBytesOutForUserClient(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v7 int32
	_ = v7
	var v11 int32
	_ = v11
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v23 int64
	_ = v23
	var v24 int64
	_ = v24
	var v32 int32
	_ = v32
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+292))
	if v3 == int32(-1) {
		return
	} else {
		v7 = *(*int32)(unsafe.Add(mBase, _consts[158]))
		if v7 == int32(0) {
			return
		} else {
			v11 = *(*int32)(unsafe.Add(mBase, _consts[63]))
			if v11 == int32(0) {
				return
			} else {
				if base.Ui32(int32(16384)) <= base.Ui32(v3) {
					F__serverAssert(m, int32(_a388), int32(_a389), int32(140))
					mBase = m.M
					v32 = m.ExcPending
					if v32 != 0 {
						return
					} else {
						F_abort(m)
						mBase = m.M
						base.Wasm_trap_unreachable()
						for {
						}
					}
				} else {
					v17 = *(*int32)(unsafe.Add(mBase, _consts[90]))
					v20 = v17 + v3*int32(24)
					v23 = *(*int64)(unsafe.Add(mBase, uint32(v20)+uint32(_consts[97])))
					v24 = *(*int64)(unsafe.Add(mBase, uint32(l0)+264))
					*(*int64)(unsafe.Add(mBase, uint32(v20)+uint32(_consts[97]))) = v23 + v24
					return
				}
			}
		}
	}
}
func F_clusterSlotStatsDecrNetworkBytesOutForReplication(m *base.Module, l0 int64) {
	var v5 int32
	_ = v5
	F_clusterSlotStatsUpdateNetworkBytesOutForReplication(m, int64(0)-l0)
	v5 = m.ExcPending
	if v5 != 0 {
		return
	} else {
		return
	}
}
func F_clusterSlotStatsIncrNetworkBytesOutForReplication(m *base.Module, l0 int64) {
	var v3 int32
	_ = v3
	F_clusterSlotStatsUpdateNetworkBytesOutForReplication(m, l0)
	v3 = m.ExcPending
	if v3 != 0 {
		return
	} else {
		return
	}
}
func F_clusterUpdateMyselfAnnouncedPorts(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v1 int32
	_ = v1
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	v1 = int32(0)
	v6 = *(*int32)(unsafe.Add(mBase, _consts[94]))
	if v6 == v1 {
		return
	} else {
		v9 = int32(_a69)
		v10 = *(*int32)(unsafe.Add(mBase, _consts[102]))
		v12 = *(*int32)(unsafe.Add(mBase, _consts[103]))
		if v10 != 0 {
			v13 = v10
		} else {
			v13 = v12
		}
		*(*int32)(unsafe.Add(mBase, uint32(v6)+2324)) = v13
		v15 = int32(_a69)
		v16 = *(*int32)(unsafe.Add(mBase, _consts[104]))
		v18 = *(*int32)(unsafe.Add(mBase, _consts[105]))
		if v16 != 0 {
			v19 = v16
		} else {
			v19 = v18
		}
		*(*int32)(unsafe.Add(mBase, uint32(v6)+2328)) = v19
		v22 = *(*int32)(unsafe.Add(mBase, _consts[106]))
		if v22 != 0 {
			v30 = v22
		} else {
			v24 = *(*int32)(unsafe.Add(mBase, _consts[107]))
			if v24 != 0 {
				v30 = v24
			} else {
				v26 = *(*int32)(unsafe.Add(mBase, _consts[89]))
				if v26 != 0 {
					v27 = v18
				} else {
					v27 = v12
				}
				v30 = v27 + int32(10000)
			}
		}
		*(*int32)(unsafe.Add(mBase, uint32(v6)+2332)) = v30
		v32 = int32(_a69)
		v33 = *(*int32)(unsafe.Add(mBase, _consts[108]))
		*(*int32)(unsafe.Add(mBase, uint32(v6)+2336)) = v33
		v36 = *(*int32)(unsafe.Add(mBase, _consts[109]))
		*(*int32)(unsafe.Add(mBase, uint32(v6)+2340)) = v36
		F_clearCachedClusterSlotsResponse(m)
		mBase = m.M
		v39 = m.ExcPending
		if v39 != 0 {
			return
		} else {
			v41 = *(*int32)(unsafe.Add(mBase, _consts[90]))
			v42 = *(*int32)(unsafe.Add(mBase, uint32(v41)+uint32(_consts[101])))
			*(*int32)(unsafe.Add(mBase, uint32(v41)+uint32(_consts[101]))) = v42 | int32(4)
			return
		}
	}
}
func F_clusterUpdateMyselfHostname(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v1 int32
	_ = v1
	var v3 int32
	_ = v3
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	v1 = int32(0)
	v3 = *(*int32)(unsafe.Add(mBase, _consts[94]))
	if v3 == v1 {
		return
	} else {
		v9 = *(*int32)(unsafe.Add(mBase, _consts[110]))
		F_updateSdsExtensionField(m, v3+int32(2312), v9)
		mBase = m.M
		v11 = m.ExcPending
		if v11 != 0 {
			return
		} else {
			return
		}
	}
}
func F_clusterUpdateSlotExportsOnOwnershipChange(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
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
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v29 int32
	_ = v29
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v55 int32
	_ = v55
	var v56 int64
	_ = v56
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v73 int32
	_ = v73
	var v83 int32
	_ = v83
	var v88 int32
	_ = v88
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	v9 = *(*int32)(unsafe.Add(mBase, _consts[90]))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(v9)+uint32(_consts[143])))
	v12 = v6 + int32(8)
	v13 = *(*int32)(unsafe.Add(mBase, uint32(v10)))
	*(*int32)(unsafe.Add(mBase, uint32(v12)+4)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v12))) = v13
	goto L1
L1:
	;
	v18 = v6 + int32(8)
	v20 = *(*int32)(unsafe.Add(mBase, uint32(v18)))
	if v20 == int32(0) {
		goto L5
	} else {
		goto L6
	}
L2:
	;
	m.G0 = v6 + int32(16)
	return
L3:
	;
	v83 = *(*int32)(unsafe.Add(mBase, _consts[150]))
	goto L25
L4:
	;
	if v20 == int32(0) {
		goto L3
	} else {
		goto L7
	}
L5:
	;
	goto L4
L6:
	;
	v23 = *(*int32)(unsafe.Add(mBase, uint32(v18)+4))
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v20+base.B2i32(v23 == int32(0))<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v18))) = v29
	goto L5
L7:
	;
	v35 = v20
	v36 = int32(0)
	goto L8
L8:
	;
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v35)+8))
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v37)))
	if v38 != 0 {
		v60 = v36
		goto L10
	} else {
		goto L11
	}
L9:
	;
	if v60 != 0 {
		goto L2
	} else {
		goto L24
	}
L10:
	;
	v62 = v6 + int32(8)
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v62)))
	if v64 == int32(0) {
		goto L21
	} else {
		goto L22
	}
L11:
	;
	v41 = F_checkSlotExportOwnership(m, v37, v6+int32(7))
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L14
	} else {
		goto L15
	}
L12:
	;
	v56 = *(*int64)(unsafe.Add(mBase, uint32(v37)+176))
	v60 = v36 + base.B2i32(v56 != int64(0))
	goto L10
L13:
	;
	v49 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6)+7)))
	if v49 != int32(1) {
		goto L12
	} else {
		goto L18
	}
L14:
	;
	return
L15:
	;
	if v41 != int32(-1) {
		goto L13
	} else {
		goto L16
	}
L16:
	;
	F_finishSlotMigrationJob(m, v37, int32(18), int32(_a370))
	mBase = m.M
	v48 = m.ExcPending
	if v48 != 0 {
		goto L14
	} else {
		goto L17
	}
L17:
	;
	goto L12
L18:
	;
	F_finishSlotMigrationJob(m, v37, int32(20), int32(0))
	mBase = m.M
	v55 = m.ExcPending
	if v55 != 0 {
		goto L14
	} else {
		goto L19
	}
L19:
	;
	goto L12
L20:
	;
	if v64 != 0 {
		v35 = v64
		v36 = v60
		goto L8
	} else {
		goto L23
	}
L21:
	;
	goto L20
L22:
	;
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v62)+4))
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v64+base.B2i32(v67 == int32(0))<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v62))) = v73
	goto L21
L23:
	;
	goto L9
L24:
	;
	goto L3
L25:
	;
	if v83 == int32(0) {
		goto L2
	} else {
		goto L26
	}
L26:
	;
	F_unpauseActions(m, int32(3))
	mBase = m.M
	v88 = m.ExcPending
	if v88 != 0 {
		goto L14
	} else {
		goto L27
	}
L27:
	;
	goto L2
}
func F_clusterUpdateSlotImportsOnOwnershipChange(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v35 int32
	_ = v35
	var v40 int32
	_ = v40
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v88 int32
	_ = v88
	var v92 int64
	_ = v92
	var v93 int32
	_ = v93
	var v98 int32
	_ = v98
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v111 int32
	_ = v111
	var v114 int32
	_ = v114
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v122 int32
	_ = v122
	var v124 int32
	_ = v124
	var v127 int32
	_ = v127
	var v133 int32
	_ = v133
	var v138 int32
	_ = v138
	var v142 int32
	_ = v142
	var v144 int32
	_ = v144
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v154 int32
	_ = v154
	var v159 int32
	_ = v159
	var v161 int32
	_ = v161
	var v169 int32
	_ = v169
	var v172 int32
	_ = v172
	var v183 int32
	_ = v183
	var v189 int32
	_ = v189
	var v191 int32
	_ = v191
	var v194 int32
	_ = v194
	var v200 int32
	_ = v200
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v224 int32
	_ = v224
	var v226 int32
	_ = v226
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v249 int32
	_ = v249
	var v254 int32
	_ = v254
	var v269 int32
	_ = v269
	var v276 int32
	_ = v276
	var v292 int32
	_ = v292
	var v298 int32
	_ = v298
	var v299 int32
	_ = v299
	var v304 int32
	_ = v304
	var v308 int32
	_ = v308
	var v319 int32
	_ = v319
	var v321 int32
	_ = v321
	var v324 int32
	_ = v324
	var v330 int32
	_ = v330
	v10 = m.G0
	v12 = v10 - int32(32)
	m.G0 = v12
	v15 = *(*int32)(unsafe.Add(mBase, _consts[90]))
	v16 = *(*int32)(unsafe.Add(mBase, uint32(v15)+uint32(_consts[143])))
	v18 = v12 + int32(16)
	v19 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
	*(*int32)(unsafe.Add(mBase, uint32(v18)+4)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v18))) = v19
	goto L1
L1:
	;
	v24 = v12 + int32(16)
	v26 = *(*int32)(unsafe.Add(mBase, uint32(v24)))
	if v26 == int32(0) {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	m.G0 = v12 + int32(32)
	return
L3:
	;
	if v26 == int32(0) {
		goto L2
	} else {
		goto L6
	}
L4:
	;
	goto L3
L5:
	;
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v24)+4))
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v26+base.B2i32(v29 == int32(0))<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v24))) = v35
	goto L4
L6:
	;
	v40 = v26
	goto L7
L7:
	;
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v40)+8))
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v48)))
	if v49 != int32(1) {
		goto L9
	} else {
		goto L10
	}
L8:
	;
	goto L2
L9:
	;
	v319 = v12 + int32(16)
	v321 = *(*int32)(unsafe.Add(mBase, uint32(v319)))
	if v321 == int32(0) {
		goto L76
	} else {
		goto L77
	}
L10:
	;
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v48)+156))
	if base.Ui32(int32(20)) < base.Ui32(v52) {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v60 = *(*int32)(unsafe.Add(mBase, _consts[90]))
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v60)))
	v62 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v61)+88)))
	if v62&int32(1) != 0 {
		goto L17
	} else {
		goto L18
	}
L12:
	;
	if int32(1)<<(uint(v52)%32)&int32(1835040) != 0 {
		goto L9
	} else {
		goto L13
	}
L13:
	;
	goto L11
L14:
	;
	v298 = *(*int32)(unsafe.Add(mBase, _consts[90]))
	v299 = *(*int32)(unsafe.Add(mBase, uint32(v298)))
	if v292 != v299 {
		goto L71
	} else {
		goto L72
	}
L15:
	;
	v292 = int32(0)
	goto L14
L16:
	;
	F__serverAssert(m, int32(_a341), int32(_a325), int32(2186))
	mBase = m.M
	v276 = m.ExcPending
	if v276 != 0 {
		goto L24
	} else {
		goto L70
	}
L17:
	;
	v111 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v48)+72)))
	if v111 == int32(0) {
		goto L9
	} else {
		goto L29
	}
L18:
	;
	if v52 == int32(6) {
		goto L9
	} else {
		goto L19
	}
L19:
	;
	v68 = *(*int32)(unsafe.Add(mBase, _consts[15]))
	if int32(2) < v68 {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	v92 = *(*int64)(unsafe.Add(mBase, _consts[37]))
	v93 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v48)+196)) = uint8(v93)
	*(*int32)(unsafe.Add(mBase, uint32(v48)+156)) = int32(6)
	*(*int64)(unsafe.Add(mBase, uint32(v48)+16)) = v92
	v98 = *(*int32)(unsafe.Add(mBase, uint32(v48)+152))
	*(*int32)(unsafe.Add(mBase, uint32(v98)+216)) = int32(0)
	v101 = *(*int32)(unsafe.Add(mBase, uint32(v48)+200))
	if v101 != 0 {
		goto L16
	} else {
		goto L26
	}
L21:
	;
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v48)+188))
	if base.Ui32(int32(17)) < base.Ui32(v52) {
		v80 = int32(_a242)
		goto L22
	} else {
		goto L23
	}
L22:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+8)) = int32(_a342)
	*(*int32)(unsafe.Add(mBase, uint32(v12)+4)) = v80
	*(*int32)(unsafe.Add(mBase, uint32(v12))) = v71
	F__serverLog(m, int32(2), int32(_a330), v12)
	mBase = m.M
	v88 = m.ExcPending
	if v88 != 0 {
		goto L24
	} else {
		goto L25
	}
L23:
	;
	v79 = *(*int32)(unsafe.Add(mBase, uint32(v52<<(uint(int32(2))%32))+uint32(_consts[144])))
	v80 = v79
	goto L22
L24:
	;
	return
L25:
	;
	goto L20
L26:
	;
	F_freeClientAsync(m, v98)
	mBase = m.M
	v103 = m.ExcPending
	if v103 != 0 {
		goto L24
	} else {
		goto L27
	}
L27:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v48)+152)) = int32(0)
	v106 = *(*int32)(unsafe.Add(mBase, uint32(v48)+204))
	F_sdsfree(m, v106)
	mBase = m.M
	v108 = m.ExcPending
	if v108 != 0 {
		goto L24
	} else {
		goto L28
	}
L28:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v48)+204)) = int32(0)
	goto L9
L29:
	;
	v114 = *(*int32)(unsafe.Add(mBase, uint32(v48)+164))
	v116 = v12 + int32(24)
	v117 = *(*int32)(unsafe.Add(mBase, uint32(v114)))
	*(*int32)(unsafe.Add(mBase, uint32(v116)+4)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v116))) = v117
	goto L30
L30:
	;
	v122 = v12 + int32(24)
	v124 = *(*int32)(unsafe.Add(mBase, uint32(v122)))
	if v124 == int32(0) {
		goto L32
	} else {
		goto L33
	}
L31:
	;
	if v124 == int32(0) {
		goto L15
	} else {
		goto L34
	}
L32:
	;
	goto L31
L33:
	;
	v127 = *(*int32)(unsafe.Add(mBase, uint32(v122)+4))
	v133 = *(*int32)(unsafe.Add(mBase, uint32(v124+base.B2i32(v127 == int32(0))<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v122))) = v133
	goto L32
L34:
	;
	v138 = v48 + int32(72)
	v142 = v124
	v144 = int32(0)
	goto L35
L35:
	;
	v149 = *(*int32)(unsafe.Add(mBase, uint32(v142)+8))
	v150 = *(*int32)(unsafe.Add(mBase, uint32(v149)))
	v151 = *(*int32)(unsafe.Add(mBase, uint32(v149)+4))
	if v151 < v150 {
		v183 = v144
		goto L37
	} else {
		goto L38
	}
L36:
	;
	if v183 != 0 {
		goto L51
	} else {
		goto L52
	}
L37:
	;
	v189 = v12 + int32(24)
	v191 = *(*int32)(unsafe.Add(mBase, uint32(v189)))
	if v191 == int32(0) {
		goto L48
	} else {
		goto L49
	}
L38:
	;
	v154 = *(*int32)(unsafe.Add(mBase, _consts[90]))
	v159 = v150
	v161 = v144
	goto L39
L39:
	;
	v169 = *(*int32)(unsafe.Add(mBase, uint32(v154+int32(52)+v159<<(uint(int32(2))%32))))
	if v169 == int32(0) {
		goto L15
	} else {
		goto L41
	}
L40:
	;
	v183 = v172
	goto L37
L41:
	;
	if v161 != 0 {
		goto L42
	} else {
		goto L43
	}
L42:
	;
	v172 = v161
	goto L44
L43:
	;
	v172 = v169
	goto L44
L44:
	;
	if v172 != v169 {
		goto L15
	} else {
		goto L45
	}
L45:
	;
	if base.B2i32(v159 == v151) == int32(0) {
		v159 = v159 + int32(1)
		v161 = v172
		goto L39
	} else {
		goto L46
	}
L46:
	;
	goto L40
L47:
	;
	if v191 != 0 {
		v142 = v191
		v144 = v183
		goto L35
	} else {
		goto L50
	}
L48:
	;
	goto L47
L49:
	;
	v194 = *(*int32)(unsafe.Add(mBase, uint32(v189)+4))
	v200 = *(*int32)(unsafe.Add(mBase, uint32(v191+base.B2i32(v194 == int32(0))<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v189))) = v200
	goto L48
L50:
	;
	goto L36
L51:
	;
	v204 = v183 + int32(8)
	v205 = int32(40)
	goto L57
L52:
	;
	v292 = int32(0)
	goto L14
L53:
	;
	if v269 == int32(0) {
		goto L9
	} else {
		goto L69
	}
L54:
	;
	v269 = int32(0)
	goto L53
L55:
	;
	v241 = v236
	v242 = v237
	v243 = v238
	goto L65
L56:
	;
	if v226 == int32(0) {
		goto L54
	} else {
		goto L63
	}
L57:
	;
	if (v138|v204)&int32(3) != 0 {
		v236 = v204
		v237 = v138
		v238 = v205
		goto L55
	} else {
		goto L58
	}
L58:
	;
	v213 = v204
	v214 = v138
	v215 = v205
	goto L59
L59:
	;
	v218 = *(*int32)(unsafe.Add(mBase, uint32(v213)))
	v219 = *(*int32)(unsafe.Add(mBase, uint32(v214)))
	if v218 != v219 {
		v236 = v213
		v237 = v214
		v238 = v215
		goto L55
	} else {
		goto L61
	}
L60:
	;
	goto L56
L61:
	;
	v221 = int32(4)
	v222 = v214 + v221
	v224 = v213 + v221
	v226 = v215 + int32(-4)
	if base.Ui32(int32(3)) < base.Ui32(v226) {
		v213 = v224
		v214 = v222
		v215 = v226
		goto L59
	} else {
		goto L62
	}
L62:
	;
	goto L60
L63:
	;
	v236 = v224
	v237 = v222
	v238 = v226
	goto L55
L64:
	;
	v269 = v246 - v247
	goto L53
L65:
	;
	v246 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v241))))
	v247 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v242))))
	if v246 != v247 {
		goto L64
	} else {
		goto L67
	}
L67:
	;
	v249 = int32(1)
	v254 = v243 + int32(-1)
	if v254 == int32(0) {
		goto L54
	} else {
		goto L68
	}
L68:
	;
	v241 = v241 + v249
	v242 = v242 + v249
	v243 = v254
	goto L65
L69:
	;
	v292 = v183
	goto L14
L70:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L71:
	;
	F_finishSlotMigrationJob(m, v48, int32(18), int32(_a343))
	mBase = m.M
	v308 = m.ExcPending
	if v308 != 0 {
		goto L24
	} else {
		goto L74
	}
L72:
	;
	F_finishSlotMigrationJob(m, v48, int32(18), int32(_a344))
	mBase = m.M
	v304 = m.ExcPending
	if v304 != 0 {
		goto L24
	} else {
		goto L73
	}
L73:
	;
	goto L9
L74:
	;
	goto L9
L75:
	;
	if v321 != 0 {
		v40 = v321
		goto L7
	} else {
		goto L78
	}
L76:
	;
	goto L75
L77:
	;
	v324 = *(*int32)(unsafe.Add(mBase, uint32(v319)+4))
	v330 = *(*int32)(unsafe.Add(mBase, uint32(v321+base.B2i32(v324 == int32(0))<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v319))) = v330
	goto L76
L78:
	;
	goto L8
}
func F_clusterUpdateState(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v21 int64
	_ = v21
	var v25 int64
	_ = v25
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v37 int32
	_ = v37
	var v40 int64
	_ = v40
	var v42 int32
	_ = v42
	var v44 int64
	_ = v44
	var v50 int32
	_ = v50
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v84 int32
	_ = v84
	var v87 int32
	_ = v87
	var v91 int32
	_ = v91
	var v99 int32
	_ = v99
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v120 int32
	_ = v120
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v128 int32
	_ = v128
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v137 int32
	_ = v137
	var v139 int32
	_ = v139
	var v141 int64
	_ = v141
	var v142 int64
	_ = v142
	var v143 int64
	_ = v143
	var v144 int64
	_ = v144
	var v145 int64
	_ = v145
	var v146 int64
	_ = v146
	var v147 int64
	_ = v147
	var v149 int64
	_ = v149
	var v151 int64
	_ = v151
	var v153 int64
	_ = v153
	var v155 int64
	_ = v155
	var v157 int64
	_ = v157
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v164 int32
	_ = v164
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v171 int32
	_ = v171
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v179 int32
	_ = v179
	var v183 int32
	_ = v183
	var v185 int32
	_ = v185
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v199 int32
	_ = v199
	var v204 int32
	_ = v204
	var v208 int32
	_ = v208
	var v212 int32
	_ = v212
	var v221 int32
	_ = v221
	var v224 int32
	_ = v224
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v233 int32
	_ = v233
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v247 int32
	_ = v247
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v260 int32
	_ = v260
	var v263 int32
	_ = v263
	var v264 int32
	_ = v264
	var v268 int32
	_ = v268
	var v272 int32
	_ = v272
	var v273 int32
	_ = v273
	var v274 int32
	_ = v274
	var v277 int32
	_ = v277
	var v279 int32
	_ = v279
	var v281 int64
	_ = v281
	var v282 int64
	_ = v282
	var v283 int64
	_ = v283
	var v284 int64
	_ = v284
	var v285 int64
	_ = v285
	var v286 int64
	_ = v286
	var v287 int64
	_ = v287
	var v289 int64
	_ = v289
	var v291 int64
	_ = v291
	var v293 int64
	_ = v293
	var v295 int64
	_ = v295
	var v297 int64
	_ = v297
	var v299 int32
	_ = v299
	var v300 int32
	_ = v300
	var v301 int32
	_ = v301
	var v304 int32
	_ = v304
	var v307 int32
	_ = v307
	var v308 int32
	_ = v308
	var v309 int32
	_ = v309
	var v311 int32
	_ = v311
	var v314 int32
	_ = v314
	var v315 int32
	_ = v315
	var v319 int32
	_ = v319
	var v323 int32
	_ = v323
	var v325 int32
	_ = v325
	var v332 int32
	_ = v332
	var v333 int32
	_ = v333
	var v334 int32
	_ = v334
	var v339 int32
	_ = v339
	var v344 int32
	_ = v344
	var v348 int32
	_ = v348
	var v352 int32
	_ = v352
	var v359 int32
	_ = v359
	var v366 int32
	_ = v366
	var v368 int32
	_ = v368
	var v369 int32
	_ = v369
	var v371 int32
	_ = v371
	var v374 int64
	_ = v374
	var v377 int32
	_ = v377
	var v380 int32
	_ = v380
	var v381 int32
	_ = v381
	var v382 int32
	_ = v382
	var v383 int32
	_ = v383
	var v386 int32
	_ = v386
	var v387 int32
	_ = v387
	var v388 int32
	_ = v388
	var v389 int32
	_ = v389
	var v395 int64
	_ = v395
	var v396 int64
	_ = v396
	var v398 int64
	_ = v398
	var v400 int64
	_ = v400
	var v403 int64
	_ = v403
	var v404 int64
	_ = v404
	var v407 int64
	_ = v407
	var v410 int32
	_ = v410
	var v412 int32
	_ = v412
	var v416 int32
	_ = v416
	var v422 int32
	_ = v422
	var v424 int32
	_ = v424
	var v430 int32
	_ = v430
	var v432 int32
	_ = v432
	var v441 int32
	_ = v441
	var v443 int32
	_ = v443
	var v445 int32
	_ = v445
	var v449 int32
	_ = v449
	var v451 int32
	_ = v451
	var v452 int32
	_ = v452
	var v456 int32
	_ = v456
	var v458 int32
	_ = v458
	var v465 int32
	_ = v465
	var v467 int32
	_ = v467
	var v469 int32
	_ = v469
	var v472 int32
	_ = v472
	v10 = m.G0
	v12 = v10 - int32(48)
	m.G0 = v12
	v15 = *(*int32)(unsafe.Add(mBase, _consts[90]))
	v16 = *(*int32)(unsafe.Add(mBase, uint32(v15)+uint32(_consts[101])))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+uint32(_consts[101]))) = v16 & int32(-3)
	v21 = *(*int64)(unsafe.Add(mBase, _consts[138]))
	if v21 != int64(0) {
		v29 = v15
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v30 = int32(0)
	v31 = *(*int32)(unsafe.Add(mBase, _consts[94]))
	v32 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v31)+88)))
	if v32&int32(1) == v30 {
		v50 = v29
		goto L4
	} else {
		goto L5
	}
L2:
	;
	v25 = F_mstime(m)
	mBase = m.M
	*(*int64)(unsafe.Add(mBase, _consts[138])) = v25
	v28 = *(*int32)(unsafe.Add(mBase, _consts[90]))
	v29 = v28
	goto L1
L3:
	;
	m.G0 = v12 + int32(48)
	return
L4:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v50)+24)) = int32(1)
	v54 = int32(0)
	v57 = *(*int32)(unsafe.Add(mBase, _consts[139]))
	if v57 == v54 {
		v99 = v54
		goto L8
	} else {
		goto L9
	}
L5:
	;
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v29)+16))
	if v37 != int32(1) {
		v50 = v29
		goto L4
	} else {
		goto L6
	}
L6:
	;
	v40 = F_mstime(m)
	mBase = m.M
	v42 = *(*int32)(unsafe.Add(mBase, _consts[90]))
	v44 = *(*int64)(unsafe.Add(mBase, _consts[138]))
	if int64(1999) < v40-v44 {
		v50 = v42
		goto L4
	} else {
		goto L7
	}
L7:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v42)+24)) = int32(0)
	goto L3
L8:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v50)+28)) = int32(0)
	v106 = *(*int32)(unsafe.Add(mBase, uint32(v50)+32))
	v107 = F_dictGetSafeIterator(m, v106)
	mBase = m.M
	v108 = m.ExcPending
	if v108 != 0 {
		goto L18
	} else {
		goto L19
	}
L9:
	;
	v69 = int32(0)
	goto L10
L10:
	;
	v72 = int32(1)
	v75 = v50 + int32(52) + v69<<(uint(int32(2))%32)
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v75)))
	if v76 == int32(0) {
		v99 = v72
		goto L8
	} else {
		goto L12
	}
L11:
	;
	v99 = int32(0)
	goto L8
L12:
	;
	v79 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v76)+88)))
	if v79&int32(8) != 0 {
		v99 = v72
		goto L8
	} else {
		goto L13
	}
L13:
	;
	v84 = *(*int32)(unsafe.Add(mBase, uint32(v75+int32(4))))
	if v84 == int32(0) {
		v99 = v72
		goto L8
	} else {
		goto L14
	}
L14:
	;
	v87 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v84)+88)))
	if v87&int32(8) != 0 {
		v99 = v72
		goto L8
	} else {
		goto L15
	}
L15:
	;
	v91 = v69 + int32(2)
	if v91 != int32(16384) {
		v69 = v91
		goto L10
	} else {
		goto L16
	}
L16:
	;
	goto L11
L17:
	;
	F_dictReleaseIterator(m, v107)
	mBase = m.M
	v366 = m.ExcPending
	if v366 != 0 {
		goto L18
	} else {
		goto L80
	}
L18:
	;
	return
L19:
	;
	v116 = v107 + int32(20)
	v117 = *(*int32)(unsafe.Add(mBase, uint32(v107)+16))
	if v117 != 0 {
		goto L22
	} else {
		goto L23
	}
L20:
	;
	if v212 == int32(0) {
		v359 = v54
		goto L17
	} else {
		goto L46
	}
L21:
	;
	v123 = v116
	v124 = v120
	goto L24
L22:
	;
	v120 = int32(1)
	goto L21
L23:
	;
	v120 = int32(0)
	goto L21
L24:
	;
	switch v124 {
	case 0:
		goto L29
	default:
		goto L28
	}
L26:
	;
	v124 = int32(0)
	goto L24
L27:
	;
	goto L20
L28:
	;
	v204 = *(*int32)(unsafe.Add(mBase, uint32(v123)))
	*(*int32)(unsafe.Add(mBase, uint32(v107)+16)) = v204
	if v204 == int32(0) {
		goto L26
	} else {
		goto L45
	}
L29:
	;
	v128 = *(*int32)(unsafe.Add(mBase, uint32(v107)+4))
	if v128 != int32(-1) {
		v167 = v128
		goto L30
	} else {
		goto L31
	}
L30:
	;
	v168 = int32(1)
	v169 = v167 + v168
	*(*int32)(unsafe.Add(mBase, uint32(v107)+4)) = v169
	v171 = int32(0)
	v174 = *(*int32)(unsafe.Add(mBase, uint32(v107)))
	v175 = *(*int32)(unsafe.Add(mBase, uint32(v107)+8))
	v179 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v174+v175+int32(26)))))
	if v179 == int32(255) {
		goto L39
	} else {
		goto L40
	}
L31:
	;
	v132 = *(*int32)(unsafe.Add(mBase, uint32(v107)+8))
	if v132 != 0 {
		v167 = int32(-1)
		goto L30
	} else {
		goto L32
	}
L32:
	;
	v133 = *(*int32)(unsafe.Add(mBase, uint32(v107)))
	v134 = *(*int32)(unsafe.Add(mBase, uint32(v107)+12))
	if v134 == int32(0) {
		goto L34
	} else {
		goto L35
	}
L33:
	;
	v161 = *(*int32)(unsafe.Add(mBase, uint32(v160)+20))
	if v161 != int32(-1) {
		goto L36
	} else {
		goto L37
	}
L34:
	;
	v141 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v133)+16)))
	v142 = int64(*(*int8)(unsafe.Add(mBase, uint32(v133)+27)))
	v143 = int64(*(*int32)(unsafe.Add(mBase, uint32(v133)+8)))
	v144 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v133)+12)))
	v145 = int64(*(*int8)(unsafe.Add(mBase, uint32(v133)+26)))
	v146 = int64(*(*int32)(unsafe.Add(mBase, uint32(v133)+4)))
	v147 = F_wangHash64(m, v146)
	mBase = m.M
	v149 = F_wangHash64(m, v145+v147)
	mBase = m.M
	v151 = F_wangHash64(m, v144+v149)
	mBase = m.M
	v153 = F_wangHash64(m, v143+v151)
	mBase = m.M
	v155 = F_wangHash64(m, v142+v153)
	mBase = m.M
	v157 = F_wangHash64(m, v141+v155)
	mBase = m.M
	*(*int64)(unsafe.Add(mBase, uint32(v107)+24)) = v157
	v159 = *(*int32)(unsafe.Add(mBase, uint32(v107)))
	v160 = v159
	goto L33
L35:
	;
	v137 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v133)+24)))
	v139 = v137 + int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v133)+24)) = uint16(v139)
	v160 = v133
	goto L33
L36:
	;
	v167 = v161 + int32(-1)
	goto L30
L37:
	;
	v164 = *(*int32)(unsafe.Add(mBase, uint32(v107)+4))
	v167 = v164
	goto L30
L38:
	;
	v194 = int32(2)
	v199 = *(*int32)(unsafe.Add(mBase, uint32(v174+v192<<(uint(v194)%32)+int32(4))))
	v123 = v199 + v193<<(uint(v194)%32)
	v124 = int32(1)
	goto L24
L39:
	;
	v183 = v171
	goto L41
L40:
	;
	v183 = v168 << (uint(v179) % 32)
	goto L41
L41:
	;
	if v169 < v183 {
		v192 = v175
		v193 = v169
		goto L38
	} else {
		goto L42
	}
L42:
	;
	if v175 != 0 {
		v212 = v171
		goto L27
	} else {
		goto L43
	}
L43:
	;
	v185 = *(*int32)(unsafe.Add(mBase, uint32(v174)+20))
	if v185 == int32(-1) {
		v212 = v171
		goto L27
	} else {
		goto L44
	}
L44:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v107)+4)) = int64(4294967296)
	v192 = int32(1)
	v193 = int32(0)
	goto L38
L45:
	;
	v208 = *(*int32)(unsafe.Add(mBase, uint32(v204)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v116))) = v208
	v212 = v204
	goto L27
L46:
	;
	v221 = v54
	v224 = v212
	goto L47
L47:
	;
	v227 = *(*int32)(unsafe.Add(mBase, uint32(v224)+8))
	goto L50
L48:
	;
	v359 = v247
	goto L17
L49:
	;
	v256 = v107 + int32(20)
	v257 = *(*int32)(unsafe.Add(mBase, uint32(v107)+16))
	if v257 != 0 {
		goto L55
	} else {
		goto L56
	}
L50:
	;
	v228 = *(*int32)(unsafe.Add(mBase, uint32(v227)+88))
	if v228&int32(1) == int32(0) {
		v247 = v221
		goto L49
	} else {
		goto L51
	}
L51:
	;
	v233 = *(*int32)(unsafe.Add(mBase, uint32(v227)+2160))
	if v233 == int32(0) {
		v247 = v221
		goto L49
	} else {
		goto L52
	}
L52:
	;
	v237 = *(*int32)(unsafe.Add(mBase, _consts[90]))
	v238 = *(*int32)(unsafe.Add(mBase, uint32(v237)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v237)+28)) = v238 + int32(1)
	v247 = v221 + base.B2i32(v228&int32(12) == int32(0))
	goto L49
L53:
	;
	if v352 != 0 {
		v221 = v247
		v224 = v352
		goto L47
	} else {
		goto L79
	}
L54:
	;
	v263 = v256
	v264 = v260
	goto L57
L55:
	;
	v260 = int32(1)
	goto L54
L56:
	;
	v260 = int32(0)
	goto L54
L57:
	;
	switch v264 {
	case 0:
		goto L62
	default:
		goto L61
	}
L59:
	;
	v264 = int32(0)
	goto L57
L60:
	;
	goto L53
L61:
	;
	v344 = *(*int32)(unsafe.Add(mBase, uint32(v263)))
	*(*int32)(unsafe.Add(mBase, uint32(v107)+16)) = v344
	if v344 == int32(0) {
		goto L59
	} else {
		goto L78
	}
L62:
	;
	v268 = *(*int32)(unsafe.Add(mBase, uint32(v107)+4))
	if v268 != int32(-1) {
		v307 = v268
		goto L63
	} else {
		goto L64
	}
L63:
	;
	v308 = int32(1)
	v309 = v307 + v308
	*(*int32)(unsafe.Add(mBase, uint32(v107)+4)) = v309
	v311 = int32(0)
	v314 = *(*int32)(unsafe.Add(mBase, uint32(v107)))
	v315 = *(*int32)(unsafe.Add(mBase, uint32(v107)+8))
	v319 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v314+v315+int32(26)))))
	if v319 == int32(255) {
		goto L72
	} else {
		goto L73
	}
L64:
	;
	v272 = *(*int32)(unsafe.Add(mBase, uint32(v107)+8))
	if v272 != 0 {
		v307 = int32(-1)
		goto L63
	} else {
		goto L65
	}
L65:
	;
	v273 = *(*int32)(unsafe.Add(mBase, uint32(v107)))
	v274 = *(*int32)(unsafe.Add(mBase, uint32(v107)+12))
	if v274 == int32(0) {
		goto L67
	} else {
		goto L68
	}
L66:
	;
	v301 = *(*int32)(unsafe.Add(mBase, uint32(v300)+20))
	if v301 != int32(-1) {
		goto L69
	} else {
		goto L70
	}
L67:
	;
	v281 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v273)+16)))
	v282 = int64(*(*int8)(unsafe.Add(mBase, uint32(v273)+27)))
	v283 = int64(*(*int32)(unsafe.Add(mBase, uint32(v273)+8)))
	v284 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v273)+12)))
	v285 = int64(*(*int8)(unsafe.Add(mBase, uint32(v273)+26)))
	v286 = int64(*(*int32)(unsafe.Add(mBase, uint32(v273)+4)))
	v287 = F_wangHash64(m, v286)
	mBase = m.M
	v289 = F_wangHash64(m, v285+v287)
	mBase = m.M
	v291 = F_wangHash64(m, v284+v289)
	mBase = m.M
	v293 = F_wangHash64(m, v283+v291)
	mBase = m.M
	v295 = F_wangHash64(m, v282+v293)
	mBase = m.M
	v297 = F_wangHash64(m, v281+v295)
	mBase = m.M
	*(*int64)(unsafe.Add(mBase, uint32(v107)+24)) = v297
	v299 = *(*int32)(unsafe.Add(mBase, uint32(v107)))
	v300 = v299
	goto L66
L68:
	;
	v277 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v273)+24)))
	v279 = v277 + int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v273)+24)) = uint16(v279)
	v300 = v273
	goto L66
L69:
	;
	v307 = v301 + int32(-1)
	goto L63
L70:
	;
	v304 = *(*int32)(unsafe.Add(mBase, uint32(v107)+4))
	v307 = v304
	goto L63
L71:
	;
	v334 = int32(2)
	v339 = *(*int32)(unsafe.Add(mBase, uint32(v314+v332<<(uint(v334)%32)+int32(4))))
	v263 = v339 + v333<<(uint(v334)%32)
	v264 = int32(1)
	goto L57
L72:
	;
	v323 = v311
	goto L74
L73:
	;
	v323 = v308 << (uint(v319) % 32)
	goto L74
L74:
	;
	if v309 < v323 {
		v332 = v315
		v333 = v309
		goto L71
	} else {
		goto L75
	}
L75:
	;
	if v315 != 0 {
		v352 = v311
		goto L60
	} else {
		goto L76
	}
L76:
	;
	v325 = *(*int32)(unsafe.Add(mBase, uint32(v314)+20))
	if v325 == int32(-1) {
		v352 = v311
		goto L60
	} else {
		goto L77
	}
L77:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v107)+4)) = int64(4294967296)
	v332 = int32(1)
	v333 = int32(0)
	goto L71
L78:
	;
	v348 = *(*int32)(unsafe.Add(mBase, uint32(v344)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v256))) = v348
	v352 = v344
	goto L60
L79:
	;
	goto L48
L80:
	;
	v368 = *(*int32)(unsafe.Add(mBase, _consts[90]))
	v369 = *(*int32)(unsafe.Add(mBase, uint32(v368)+28))
	v371 = base.I32_div_s(v369, int32(2))
	if v371 < v359 {
		v380 = v99
		v381 = v99
		v382 = v368
		goto L81
	} else {
		goto L82
	}
L81:
	;
	v383 = *(*int32)(unsafe.Add(mBase, uint32(v382)+16))
	if v381 == v383 {
		goto L85
	} else {
		goto L86
	}
L82:
	;
	v374 = F_mstime(m)
	mBase = m.M
	*(*int64)(unsafe.Add(mBase, _consts[140])) = v374
	v377 = *(*int32)(unsafe.Add(mBase, _consts[90]))
	v380 = int32(2)
	v381 = int32(1)
	v382 = v377
	goto L81
L83:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v472)+20)) = int32(0)
	goto L3
L84:
	;
	if v380 == v451 {
		goto L3
	} else {
		goto L110
	}
L85:
	;
	if v381 == int32(0) {
		v472 = v382
		goto L83
	} else {
		goto L109
	}
L86:
	;
	if v381 != 0 {
		v410 = int32(3)
		goto L87
	} else {
		goto L88
	}
L87:
	;
	v412 = *(*int32)(unsafe.Add(mBase, _consts[15]))
	if v410 < v412 {
		goto L97
	} else {
		goto L98
	}
L88:
	;
	v386 = int32(2)
	v387 = int32(0)
	v388 = *(*int32)(unsafe.Add(mBase, _consts[94]))
	v389 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v388)+88)))
	if v389&int32(1) == v387 {
		v410 = v386
		goto L87
	} else {
		goto L89
	}
L89:
	;
	v395 = *(*int64)(unsafe.Add(mBase, _consts[117]))
	v396 = F_mstime(m)
	mBase = m.M
	v398 = *(*int64)(unsafe.Add(mBase, _consts[140]))
	v400 = int64(5000)
	if v395 < v400 {
		goto L90
	} else {
		goto L91
	}
L90:
	;
	v403 = v395
	goto L92
L91:
	;
	v403 = v400
	goto L92
L92:
	;
	v404 = int64(500)
	if v404 < v403 {
		goto L93
	} else {
		goto L94
	}
L93:
	;
	v407 = v403
	goto L95
L94:
	;
	v407 = v404
	goto L95
L95:
	;
	if v396-v398 < v407 {
		goto L3
	} else {
		goto L96
	}
L96:
	;
	v410 = v386
	goto L87
L97:
	;
	v424 = *(*int32)(unsafe.Add(mBase, _consts[90]))
	*(*int32)(unsafe.Add(mBase, uint32(v424)+16)) = v381
	if v381 == int32(0) {
		v472 = v424
		goto L83
	} else {
		goto L103
	}
L98:
	;
	if v381 != 0 {
		goto L99
	} else {
		goto L100
	}
L99:
	;
	v416 = int32(_a205)
	goto L101
L100:
	;
	v416 = int32(_a294)
	goto L101
L101:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+32)) = v416
	F__serverLog(m, v410, int32(_a295), v12+int32(32))
	mBase = m.M
	v422 = m.ExcPending
	if v422 != 0 {
		goto L18
	} else {
		goto L102
	}
L102:
	;
	goto L97
L103:
	;
	switch v380 {
	default:
		v445 = v424
		goto L104
	case 1:
		v430 = int32(_a296)
		goto L105
	case 2:
		goto L106
	}
L104:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v445)+20)) = v380
	v451 = v380
	v452 = v445
	goto L84
L105:
	;
	v432 = *(*int32)(unsafe.Add(mBase, _consts[15]))
	if int32(3) < v432 {
		v445 = v424
		goto L104
	} else {
		goto L107
	}
L106:
	;
	v430 = int32(_a297)
	goto L105
L107:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = v430
	F__serverLog(m, int32(3), int32(_a298), v12+int32(16))
	mBase = m.M
	v441 = m.ExcPending
	if v441 != 0 {
		goto L18
	} else {
		goto L108
	}
L108:
	;
	v443 = *(*int32)(unsafe.Add(mBase, _consts[90]))
	v445 = v443
	goto L104
L109:
	;
	v449 = *(*int32)(unsafe.Add(mBase, uint32(v382)+20))
	v451 = v449
	v452 = v382
	goto L84
L110:
	;
	switch v380 {
	default:
		v469 = v452
		goto L111
	case 1:
		v456 = int32(_a296)
		goto L112
	case 2:
		goto L113
	}
L111:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v469)+20)) = v380
	goto L3
L112:
	;
	v458 = *(*int32)(unsafe.Add(mBase, _consts[15]))
	if int32(3) < v458 {
		v469 = v452
		goto L111
	} else {
		goto L114
	}
L113:
	;
	v456 = int32(_a297)
	goto L112
L114:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12))) = v456
	F__serverLog(m, int32(3), int32(_a298), v12)
	mBase = m.M
	v465 = m.ExcPending
	if v465 != 0 {
		goto L18
	} else {
		goto L115
	}
L115:
	;
	v467 = *(*int32)(unsafe.Add(mBase, _consts[90]))
	v469 = v467
	goto L111
}
func F_freeClusterNode(m *base.Module, l0 int32) {
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
	var v21 int32
	_ = v21
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v44 int32
	_ = v44
	var v49 int32
	_ = v49
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v62 int32
	_ = v62
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v94 int32
	_ = v94
	var v99 int32
	_ = v99
	var v102 int32
	_ = v102
	var v105 int32
	_ = v105
	var v112 int32
	_ = v112
	var v116 int32
	_ = v116
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v123 int32
	_ = v123
	var v125 int32
	_ = v125
	var v130 int32
	_ = v130
	var v134 int32
	_ = v134
	var v141 int32
	_ = v141
	var v144 int32
	_ = v144
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v157 int32
	_ = v157
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v162 int32
	_ = v162
	var v164 int32
	_ = v164
	var v175 int32
	_ = v175
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v184 int32
	_ = v184
	var v190 int32
	_ = v190
	var v197 int32
	_ = v197
	var v201 int32
	_ = v201
	var v204 int32
	_ = v204
	var v210 int32
	_ = v210
	var v217 int32
	_ = v217
	var v221 int32
	_ = v221
	var v224 int32
	_ = v224
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
	var v238 int32
	_ = v238
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v243 int32
	_ = v243
	var v245 int32
	_ = v245
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v251 int32
	_ = v251
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v258 int32
	_ = v258
	var v260 int32
	_ = v260
	var v262 int32
	_ = v262
	var v267 int32
	_ = v267
	var v279 int32
	_ = v279
	var v280 int32
	_ = v280
	var v282 int32
	_ = v282
	var v284 int32
	_ = v284
	var v300 int32
	_ = v300
	var v301 int32
	_ = v301
	var v303 int32
	_ = v303
	var v304 int32
	_ = v304
	var v305 int32
	_ = v305
	var v306 int32
	_ = v306
	var v308 int32
	_ = v308
	var v309 int32
	_ = v309
	var v313 int32
	_ = v313
	var v314 int32
	_ = v314
	var v318 int32
	_ = v318
	var v319 int32
	_ = v319
	var v321 int32
	_ = v321
	var v322 int32
	_ = v322
	var v324 int32
	_ = v324
	var v325 int32
	_ = v325
	var v327 int32
	_ = v327
	var v328 int32
	_ = v328
	var v330 int32
	_ = v330
	var v331 int32
	_ = v331
	var v333 int32
	_ = v333
	var v334 int32
	_ = v334
	var v336 int32
	_ = v336
	var v337 int32
	_ = v337
	var v339 int32
	_ = v339
	var v341 int32
	_ = v341
	var v346 int32
	_ = v346
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+2164))
	if v9 < int32(1) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v94 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+88)))
	if v94&int32(2) == int32(0) {
		goto L12
	} else {
		goto L13
	}
L2:
	;
	v13 = v9 & int32(3)
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+2168))
	v15 = int32(0)
	if base.Ui32(v9) < base.Ui32(int32(4)) {
		v62 = v15
		goto L3
	} else {
		goto L4
	}
L3:
	;
	if v13 == int32(0) {
		goto L1
	} else {
		goto L8
	}
L4:
	;
	v21 = int32(0)
	v28 = v21
	v30 = v21
	goto L5
L5:
	;
	v33 = v14 + v28<<(uint(int32(2))%32)
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v33)))
	v35 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v34)+2172)) = v35
	v37 = int32(4)
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v33+v37)))
	*(*int32)(unsafe.Add(mBase, uint32(v39)+2172)) = v35
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v33+int32(8))))
	*(*int32)(unsafe.Add(mBase, uint32(v44)+2172)) = v35
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v33+int32(12))))
	*(*int32)(unsafe.Add(mBase, uint32(v49)+2172)) = v35
	v53 = v28 + v37
	v55 = v30 + v37
	if v55 != v9&int32(2147483644) {
		v28 = v53
		v30 = v55
		goto L5
	} else {
		goto L7
	}
L6:
	;
	v62 = v53
	goto L3
L7:
	;
	goto L6
L8:
	;
	v71 = v15
	v72 = v62
	goto L9
L9:
	;
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v14+v72<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v78)+2172)) = int32(0)
	v81 = int32(1)
	v84 = v71 + v81
	if v84 != v13 {
		v71 = v84
		v72 = v72 + v81
		goto L9
	} else {
		goto L11
	}
L10:
	;
	goto L1
L11:
	;
	goto L10
L12:
	;
	v300 = F_sdsnewlen(m, l0+int32(8), int32(40))
	mBase = m.M
	v301 = m.ExcPending
	if v301 != 0 {
		goto L65
	} else {
		goto L66
	}
L13:
	;
	v99 = *(*int32)(unsafe.Add(mBase, uint32(l0)+2172))
	if v99 == int32(0) {
		goto L12
	} else {
		goto L14
	}
L14:
	;
	v102 = *(*int32)(unsafe.Add(mBase, uint32(v99)+2164))
	if v102 < int32(1) {
		goto L12
	} else {
		goto L15
	}
L15:
	;
	v105 = *(*int32)(unsafe.Add(mBase, uint32(v99)+2168))
	v112 = int32(0)
	goto L16
L16:
	;
	v116 = v112 + int32(1)
	v119 = v105 + v112<<(uint(int32(2))%32)
	v120 = *(*int32)(unsafe.Add(mBase, uint32(v119)))
	if v120 != l0 {
		goto L18
	} else {
		goto L19
	}
L17:
	;
	goto L12
L18:
	;
	if v116 != v102 {
		v112 = v116
		goto L16
	} else {
		goto L64
	}
L19:
	;
	if v102 <= v116 {
		v280 = v102
		goto L20
	} else {
		goto L21
	}
L20:
	;
	v282 = v280 + int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(v99)+2164)) = v282
	if v282 != 0 {
		goto L12
	} else {
		goto L63
	}
L21:
	;
	v123 = int32(2)
	v125 = v105 + v116<<(uint(v123)%32)
	v130 = (v102 + (v112 ^ int32(-1))) << (uint(v123) % 32)
	if v119 == v125 {
		goto L23
	} else {
		goto L24
	}
L22:
	;
	v279 = *(*int32)(unsafe.Add(mBase, uint32(v99)+2164))
	v280 = v279
	goto L20
L23:
	;
	goto L22
L24:
	;
	v134 = v130 + v119
	if base.Ui32(int32(0)-v130<<(uint(int32(1))%32)) < base.Ui32(v125-v134) {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	v144 = (v125 ^ v119) & int32(3)
	if base.Ui32(v125) <= base.Ui32(v119) {
		goto L29
	} else {
		goto L30
	}
L26:
	;
	v141 = F___memcpy(m, v119, v125, v130)
	mBase = m.M
	goto L22
L27:
	;
	if v250 == int32(0) {
		goto L23
	} else {
		goto L59
	}
L28:
	;
	if base.Ui32(v228) <= base.Ui32(int32(3)) {
		v249 = v227
		v250 = v228
		v251 = v229
		goto L27
	} else {
		goto L55
	}
L29:
	;
	if v144 != 0 {
		v210 = v130
		goto L39
	} else {
		goto L40
	}
L30:
	;
	if v144 == int32(0) {
		goto L31
	} else {
		goto L32
	}
L31:
	;
	if v119&int32(3) != 0 {
		goto L33
	} else {
		goto L34
	}
L32:
	;
	v249 = v125
	v250 = v130
	v251 = v119
	goto L27
L33:
	;
	v151 = v125
	v152 = v130
	v153 = v119
	goto L35
L34:
	;
	v227 = v125
	v228 = v130
	v229 = v119
	goto L28
L35:
	;
	if v152 == int32(0) {
		goto L23
	} else {
		goto L37
	}
L37:
	;
	v157 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v151))))
	*(*uint8)(unsafe.Add(mBase, uint32(v153))) = uint8(v157)
	v159 = int32(1)
	v160 = v151 + v159
	v162 = v152 + int32(-1)
	v164 = v153 + v159
	if v164&int32(3) == int32(0) {
		v227 = v160
		v228 = v162
		v229 = v164
		goto L28
	} else {
		goto L38
	}
L38:
	;
	v151 = v160
	v152 = v162
	v153 = v164
	goto L35
L39:
	;
	if v210 == int32(0) {
		goto L23
	} else {
		goto L51
	}
L40:
	;
	if v134&int32(3) == int32(0) {
		v190 = v130
		goto L41
	} else {
		goto L42
	}
L41:
	;
	if base.Ui32(v190) <= base.Ui32(int32(3)) {
		v210 = v190
		goto L39
	} else {
		goto L47
	}
L42:
	;
	v175 = v130
	goto L43
L43:
	;
	if v175 == int32(0) {
		goto L23
	} else {
		goto L45
	}
L44:
	;
	v190 = v181
	goto L41
L45:
	;
	v181 = v175 + int32(-1)
	v182 = v119 + v181
	v184 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v125+v181))))
	*(*uint8)(unsafe.Add(mBase, uint32(v182))) = uint8(v184)
	if v182&int32(3) != 0 {
		v175 = v181
		goto L43
	} else {
		goto L46
	}
L46:
	;
	goto L44
L47:
	;
	v197 = v190
	goto L48
L48:
	;
	v201 = v197 + int32(-4)
	v204 = *(*int32)(unsafe.Add(mBase, uint32(v125+v201)))
	*(*int32)(unsafe.Add(mBase, uint32(v119+v201))) = v204
	if base.Ui32(int32(3)) < base.Ui32(v201) {
		v197 = v201
		goto L48
	} else {
		goto L50
	}
L49:
	;
	v210 = v201
	goto L39
L50:
	;
	goto L49
L51:
	;
	v217 = v210
	goto L52
L52:
	;
	v221 = v217 + int32(-1)
	v224 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v125+v221))))
	*(*uint8)(unsafe.Add(mBase, uint32(v119+v221))) = uint8(v224)
	if v221 != 0 {
		v217 = v221
		goto L52
	} else {
		goto L54
	}
L54:
	;
	goto L23
L55:
	;
	v234 = v227
	v235 = v228
	v236 = v229
	goto L56
L56:
	;
	v238 = *(*int32)(unsafe.Add(mBase, uint32(v234)))
	*(*int32)(unsafe.Add(mBase, uint32(v236))) = v238
	v240 = int32(4)
	v241 = v234 + v240
	v243 = v236 + v240
	v245 = v235 + int32(-4)
	if base.Ui32(int32(3)) < base.Ui32(v245) {
		v234 = v241
		v235 = v245
		v236 = v243
		goto L56
	} else {
		goto L58
	}
L57:
	;
	v249 = v241
	v250 = v245
	v251 = v243
	goto L27
L58:
	;
	goto L57
L59:
	;
	v256 = v249
	v257 = v250
	v258 = v251
	goto L60
L60:
	;
	v260 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v256))))
	*(*uint8)(unsafe.Add(mBase, uint32(v258))) = uint8(v260)
	v262 = int32(1)
	v267 = v257 + int32(-1)
	if v267 != 0 {
		v256 = v256 + v262
		v257 = v267
		v258 = v258 + v262
		goto L60
	} else {
		goto L62
	}
L61:
	;
	goto L23
L62:
	;
	goto L61
L63:
	;
	v284 = *(*int32)(unsafe.Add(mBase, uint32(v99)+88))
	*(*int32)(unsafe.Add(mBase, uint32(v99)+88)) = v284 & int32(-257)
	goto L12
L64:
	;
	goto L17
L65:
	;
	return
L66:
	;
	v303 = *(*int32)(unsafe.Add(mBase, _consts[90]))
	v304 = *(*int32)(unsafe.Add(mBase, uint32(v303)+32))
	v305 = F_dictDelete(m, v304, v300)
	mBase = m.M
	v306 = m.ExcPending
	if v306 != 0 {
		goto L65
	} else {
		goto L68
	}
L67:
	;
	F__serverAssert(m, int32(_a232), int32(_a179), int32(2197))
	mBase = m.M
	v346 = m.ExcPending
	if v346 != 0 {
		goto L65
	} else {
		goto L85
	}
L68:
	;
	if v305 != 0 {
		goto L67
	} else {
		goto L69
	}
L69:
	;
	F_sdsfree(m, v300)
	mBase = m.M
	v308 = m.ExcPending
	if v308 != 0 {
		goto L65
	} else {
		goto L70
	}
L70:
	;
	v309 = *(*int32)(unsafe.Add(mBase, uint32(l0)+2344))
	if v309 == int32(0) {
		goto L71
	} else {
		goto L72
	}
L71:
	;
	v314 = *(*int32)(unsafe.Add(mBase, uint32(l0)+2348))
	if v314 == int32(0) {
		goto L74
	} else {
		goto L75
	}
L72:
	;
	F_freeClusterLink(m, v309)
	mBase = m.M
	v313 = m.ExcPending
	if v313 != 0 {
		goto L65
	} else {
		goto L73
	}
L73:
	;
	goto L71
L74:
	;
	v319 = *(*int32)(unsafe.Add(mBase, uint32(l0)+2312))
	F_sdsfree(m, v319)
	mBase = m.M
	v321 = m.ExcPending
	if v321 != 0 {
		goto L65
	} else {
		goto L77
	}
L75:
	;
	F_freeClusterLink(m, v314)
	mBase = m.M
	v318 = m.ExcPending
	if v318 != 0 {
		goto L65
	} else {
		goto L76
	}
L76:
	;
	goto L74
L77:
	;
	v322 = *(*int32)(unsafe.Add(mBase, uint32(l0)+2316))
	F_sdsfree(m, v322)
	mBase = m.M
	v324 = m.ExcPending
	if v324 != 0 {
		goto L65
	} else {
		goto L78
	}
L78:
	;
	v325 = *(*int32)(unsafe.Add(mBase, uint32(l0)+2320))
	F_sdsfree(m, v325)
	mBase = m.M
	v327 = m.ExcPending
	if v327 != 0 {
		goto L65
	} else {
		goto L79
	}
L79:
	;
	v328 = *(*int32)(unsafe.Add(mBase, uint32(l0)+2304))
	F_sdsfree(m, v328)
	mBase = m.M
	v330 = m.ExcPending
	if v330 != 0 {
		goto L65
	} else {
		goto L80
	}
L80:
	;
	v331 = *(*int32)(unsafe.Add(mBase, uint32(l0)+2308))
	F_sdsfree(m, v331)
	mBase = m.M
	v333 = m.ExcPending
	if v333 != 0 {
		goto L65
	} else {
		goto L81
	}
L81:
	;
	v334 = *(*int32)(unsafe.Add(mBase, uint32(l0)+2352))
	F_raxFree(m, v334)
	mBase = m.M
	v336 = m.ExcPending
	if v336 != 0 {
		goto L65
	} else {
		goto L82
	}
L82:
	;
	v337 = *(*int32)(unsafe.Add(mBase, uint32(l0)+2168))
	F_valkey_free(m, v337)
	mBase = m.M
	v339 = m.ExcPending
	if v339 != 0 {
		goto L65
	} else {
		goto L83
	}
L83:
	;
	F_valkey_free(m, l0)
	mBase = m.M
	v341 = m.ExcPending
	if v341 != 0 {
		goto L65
	} else {
		goto L84
	}
L84:
	;
	return
L85:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_getClusterSize(m *base.Module) int32 {
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
	v3 = *(*int32)(unsafe.Add(mBase, _consts[90]))
	v4 = *(*int32)(unsafe.Add(mBase, uint32(v3)+32))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(v4)+16))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(v4)+12))
	return v5 + v6
}
func F_updateClusterAvailabilityZone(m *base.Module, l0 int32) int32 {
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	F_clusterUpdateMyselfAvailabilityZone(m)
	v5 = m.ExcPending
	if v5 != 0 {
		return int32(0)
	} else {
		F_clearCachedClusterSlotsResponse(m)
		v7 = m.ExcPending
		if v7 != 0 {
			return int32(0)
		} else {
			return int32(1)
		}
	}
}
func F_updateClusterHostname(m *base.Module, l0 int32) int32 {
	var v5 int32
	_ = v5
	F_clusterUpdateMyselfHostname(m)
	v5 = m.ExcPending
	if v5 != 0 {
		return int32(0)
	} else {
		return int32(1)
	}
}
func F_verifyClusterConfigWithData(m *base.Module) int32 {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v27 int32
	_ = v27
	var v34 int32
	_ = v34
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v53 int32
	_ = v53
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
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
	var v90 int32
	_ = v90
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v102 int32
	_ = v102
	var v104 int32
	_ = v104
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v111 int32
	_ = v111
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
	var v119 int32
	_ = v119
	var v121 int32
	_ = v121
	var v123 int32
	_ = v123
	var v130 int32
	_ = v130
	var v132 int32
	_ = v132
	var v134 int32
	_ = v134
	var v136 int32
	_ = v136
	var v140 int32
	_ = v140
	var v142 int32
	_ = v142
	var v147 int32
	_ = v147
	var v149 int32
	_ = v149
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v155 int32
	_ = v155
	var v157 int32
	_ = v157
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v165 int32
	_ = v165
	var v169 int32
	_ = v169
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v203 int32
	_ = v203
	var v206 int32
	_ = v206
	var v209 int32
	_ = v209
	var v212 int64
	_ = v212
	var v224 int32
	_ = v224
	var v233 int32
	_ = v233
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v244 int32
	_ = v244
	var v260 int32
	_ = v260
	var v264 int32
	_ = v264
	var v267 int32
	_ = v267
	var v268 int32
	_ = v268
	var v269 int32
	_ = v269
	var v275 int32
	_ = v275
	var v281 int32
	_ = v281
	var v282 int32
	_ = v282
	var v286 int32
	_ = v286
	var v293 int32
	_ = v293
	v11 = m.G0
	v13 = v11 - int32(64)
	m.G0 = v13
	v16 = int32(*(*uint8)(unsafe.Add(mBase, _consts[141])))
	if v16&int32(4) != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	m.G0 = v13 + int32(64)
	return int32(0)
L2:
	;
	v20 = *(*int32)(unsafe.Add(mBase, _consts[94]))
	v21 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20)+88)))
	if v21&int32(2) != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	F_clusterCleanSlotImportsAfterLoad(m)
	mBase = m.M
	v27 = m.ExcPending
	if v27 != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	return int32(0)
L5:
	;
	v34 = int32(0)
	v40 = v34
	v41 = v34
	goto L6
L6:
	;
	v53 = *(*int32)(unsafe.Add(mBase, _consts[10]))
	if int32(1) <= v53 {
		goto L11
	} else {
		goto L12
	}
L7:
	;
	if v269 == int32(0) {
		goto L1
	} else {
		goto L61
	}
L8:
	;
	v275 = v40 + int32(1)
	if v275 != int32(16384) {
		v40 = v275
		v41 = v269
		goto L6
	} else {
		goto L60
	}
L9:
	;
	if v90 == int32(0) {
		v269 = v41
		goto L8
	} else {
		goto L18
	}
L10:
	;
	goto L9
L11:
	;
	v58 = *(*int32)(unsafe.Add(mBase, _consts[23]))
	v59 = int32(0)
	v62 = v53
	v63 = v59
	v64 = v58
	v65 = v59
	goto L13
L12:
	;
	v90 = int32(0)
	goto L10
L13:
	;
	v68 = int32(0)
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v64+v65<<(uint(int32(2))%32))))
	if v72 == v68 {
		v81 = v62
		v82 = v64
		v83 = v68
		goto L15
	} else {
		goto L16
	}
L14:
	;
	v90 = v84
	goto L10
L15:
	;
	v84 = v83 + v63
	v86 = v65 + int32(1)
	if v86 < v81 {
		v62 = v81
		v63 = v84
		v64 = v82
		v65 = v86
		goto L13
	} else {
		goto L17
	}
L16:
	;
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v72)))
	v76 = F_kvstoreHashtableSize(m, v75, v40)
	mBase = m.M
	v77 = int32(_a69)
	v78 = *(*int32)(unsafe.Add(mBase, _consts[10]))
	v80 = *(*int32)(unsafe.Add(mBase, _consts[23]))
	v81 = v78
	v82 = v80
	v83 = v76
	goto L15
L17:
	;
	goto L14
L18:
	;
	v98 = *(*int32)(unsafe.Add(mBase, _consts[90]))
	v100 = v40 << (uint(int32(2)) % 32)
	v102 = *(*int32)(unsafe.Add(mBase, uint32(v98+v100)+52))
	v104 = *(*int32)(unsafe.Add(mBase, _consts[94]))
	if v102 == v104 {
		v269 = v41
		goto L8
	} else {
		goto L19
	}
L19:
	;
	v106 = *(*int32)(unsafe.Add(mBase, uint32(v98)+48))
	v107 = F_dictFind(m, v106, v40)
	mBase = m.M
	v108 = m.ExcPending
	if v108 != 0 {
		goto L4
	} else {
		goto L21
	}
L20:
	;
	v113 = *(*int32)(unsafe.Add(mBase, _consts[90]))
	v114 = *(*int32)(unsafe.Add(mBase, uint32(v113)+48))
	v115 = F_dictFind(m, v114, v40)
	mBase = m.M
	v116 = m.ExcPending
	if v116 != 0 {
		goto L4
	} else {
		goto L27
	}
L21:
	;
	if v107 == int32(0) {
		goto L20
	} else {
		goto L22
	}
L22:
	;
	v111 = *(*int32)(unsafe.Add(mBase, uint32(v107)+8))
	goto L23
L23:
	;
	if v111 != 0 {
		v269 = v41
		goto L8
	} else {
		goto L24
	}
L24:
	;
	goto L20
L25:
	;
	v121 = v41 + int32(1)
	if v102 != 0 {
		goto L30
	} else {
		goto L31
	}
L26:
	;
	v118 = *(*int32)(unsafe.Add(mBase, uint32(v115)+8))
	goto L29
L27:
	;
	if v115 != 0 {
		goto L26
	} else {
		goto L28
	}
L28:
	;
	v119 = int32(0)
	goto L25
L29:
	;
	v119 = v118
	goto L25
L30:
	;
	if v119 == v102 {
		v269 = v121
		goto L8
	} else {
		goto L49
	}
L31:
	;
	v123 = *(*int32)(unsafe.Add(mBase, _consts[15]))
	if int32(2) < v123 {
		goto L32
	} else {
		goto L33
	}
L32:
	;
	v132 = *(*int32)(unsafe.Add(mBase, _consts[90]))
	v134 = *(*int32)(unsafe.Add(mBase, uint32(v132+v100)+52))
	if v134 != 0 {
		v269 = v121
		goto L8
	} else {
		goto L35
	}
L33:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13))) = v40
	F__serverLog(m, int32(2), int32(_a299), v13)
	mBase = m.M
	v130 = m.ExcPending
	if v130 != 0 {
		goto L4
	} else {
		goto L34
	}
L34:
	;
	goto L32
L35:
	;
	v136 = *(*int32)(unsafe.Add(mBase, _consts[94]))
	v140 = m.G0
	v142 = v140 - int32(32)
	m.G0 = v142
	v147 = int32(1) << (uint(v40&int32(7)) % 32)
	v149 = base.I32_div_s(v40, int32(8))
	v152 = v136 + v149 + int32(104)
	v153 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v152))))
	if v147&v153 != 0 {
		goto L37
	} else {
		goto L38
	}
L36:
	;
	v191 = int32(_a69)
	v192 = *(*int32)(unsafe.Add(mBase, _consts[90]))
	*(*int32)(unsafe.Add(mBase, uint32(v192+v100)+52)) = v136
	v197 = v192 + int32(base.Ui32(v40)>>(uint(int32(3))%32))
	v198 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v197)+uint32(_consts[96]))))
	v203 = v198 & base.I32_rotl(int32(-2), v40&int32(7))
	*(*uint8)(unsafe.Add(mBase, uint32(v197)+uint32(_consts[96]))) = uint8(v203)
	v206 = *(*int32)(unsafe.Add(mBase, _consts[90]))
	v209 = v206 + v40*int32(24)
	v212 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v209)+uint32(_consts[97]))) = v212
	*(*int64)(unsafe.Add(mBase, uint32(v209)+uint32(_consts[98]))) = v212
	*(*int64)(unsafe.Add(mBase, uint32(v209)+uint32(_consts[99]))) = v212
	goto L48
L37:
	;
	m.G0 = v142 + int32(32)
	goto L36
L38:
	;
	v155 = v153 | v147
	*(*uint8)(unsafe.Add(mBase, uint32(v152))) = uint8(v155)
	v157 = *(*int32)(unsafe.Add(mBase, uint32(v136)+2160))
	*(*int32)(unsafe.Add(mBase, uint32(v136)+2160)) = v157 + int32(1)
	if v157 != 0 {
		goto L37
	} else {
		goto L39
	}
L39:
	;
	v162 = *(*int32)(unsafe.Add(mBase, _consts[90]))
	v163 = *(*int32)(unsafe.Add(mBase, uint32(v162)+32))
	F_dictInitIterator(m, v142, v163)
	mBase = m.M
	v165 = F_dictNext(m, v142)
	mBase = m.M
	if v165 == int32(0) {
		goto L37
	} else {
		goto L40
	}
L40:
	;
	v169 = v165
	goto L42
L41:
	;
	v179 = *(*int32)(unsafe.Add(mBase, uint32(v136)+88))
	*(*int32)(unsafe.Add(mBase, uint32(v136)+88)) = v179 | int32(256)
	goto L37
L42:
	;
	v173 = F_dictGetVal(m, v169)
	mBase = m.M
	v174 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v173)+88)))
	if v174&int32(2) != 0 {
		goto L44
	} else {
		goto L45
	}
L44:
	;
	v178 = F_dictNext(m, v142)
	mBase = m.M
	if v178 != 0 {
		v169 = v178
		goto L42
	} else {
		goto L47
	}
L45:
	;
	v177 = *(*int32)(unsafe.Add(mBase, uint32(v173)+2164))
	if v177 != 0 {
		goto L41
	} else {
		goto L46
	}
L46:
	;
	goto L44
L47:
	;
	goto L37
L48:
	;
	v269 = v121
	goto L8
L49:
	;
	v224 = *(*int32)(unsafe.Add(mBase, _consts[15]))
	if v119 != 0 {
		goto L51
	} else {
		goto L52
	}
L50:
	;
	v264 = *(*int32)(unsafe.Add(mBase, _consts[142]))
	v267 = F_delKeysInSlot(m, v40, v264, int32(1), int32(0))
	mBase = m.M
	v268 = m.ExcPending
	if v268 != 0 {
		goto L4
	} else {
		goto L59
	}
L51:
	;
	if int32(2) < v224 {
		goto L50
	} else {
		goto L55
	}
L52:
	;
	if int32(2) < v224 {
		goto L50
	} else {
		goto L53
	}
L53:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+16)) = v40
	F__serverLog(m, int32(2), int32(_a300), v13+int32(16))
	mBase = m.M
	v233 = m.ExcPending
	if v233 != 0 {
		goto L4
	} else {
		goto L54
	}
L54:
	;
	goto L50
L55:
	;
	v236 = F_humanNodename(m, v119)
	mBase = m.M
	v237 = m.ExcPending
	if v237 != 0 {
		goto L4
	} else {
		goto L56
	}
L56:
	;
	v238 = F_humanNodename(m, v102)
	mBase = m.M
	v239 = m.ExcPending
	if v239 != 0 {
		goto L4
	} else {
		goto L57
	}
L57:
	;
	v240 = int32(48)
	*(*int32)(unsafe.Add(mBase, uint32(v13+int32(56)))) = v102 + v240
	*(*int32)(unsafe.Add(mBase, uint32(v13+int32(52)))) = v238
	v244 = int32(8)
	*(*int32)(unsafe.Add(mBase, uint32(v13+int32(48)))) = v102 + v244
	*(*int32)(unsafe.Add(mBase, uint32(v13)+44)) = v40
	*(*int32)(unsafe.Add(mBase, uint32(v13)+40)) = v119 + v240
	*(*int32)(unsafe.Add(mBase, uint32(v13)+36)) = v236
	*(*int32)(unsafe.Add(mBase, uint32(v13)+32)) = v119 + v244
	F__serverLog(m, int32(2), int32(_a301), v13+int32(32))
	mBase = m.M
	v260 = m.ExcPending
	if v260 != 0 {
		goto L4
	} else {
		goto L58
	}
L58:
	;
	goto L50
L59:
	;
	v269 = v121
	goto L8
L60:
	;
	goto L7
L61:
	;
	v281 = F_clusterSaveConfig(m, int32(1))
	mBase = m.M
	v282 = m.ExcPending
	if v282 != 0 {
		goto L4
	} else {
		goto L62
	}
L62:
	;
	if v281 != int32(-1) {
		goto L1
	} else {
		goto L63
	}
L63:
	;
	v286 = *(*int32)(unsafe.Add(mBase, _consts[15]))
	if int32(3) < v286 {
		goto L64
	} else {
		goto L65
	}
L64:
	;
	m.Env.Exit(m, int32(1))
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L65:
	;
	F__serverLog(m, int32(3), int32(_a302), int32(0))
	mBase = m.M
	v293 = m.ExcPending
	if v293 != 0 {
		goto L4
	} else {
		goto L66
	}
L66:
	;
	goto L64
}
