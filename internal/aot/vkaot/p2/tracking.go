package p2

import (
	base "github.com/shibukawa/valkeymem/internal/aot/vkaot/base"
	"unsafe"
)

func F_enableTracking(m *base.Module, l0 int32, l1 int64, l2 int32, l3 int32, l4 int32) {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
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
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v51 int32
	_ = v51
	var v63 int32
	_ = v63
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
	var v88 int32
	_ = v88
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v101 int32
	_ = v101
	var v111 int32
	_ = v111
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+204))
	if v10&int32(4) != 0 {
		v20 = v10
		goto L1
	} else {
		goto L2
	}
L1:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+204)) = v20&int32(-381) | int32(4)
	F_initClientPubSubData(m, l0)
	mBase = m.M
	v27 = m.ExcPending
	if v27 != 0 {
		goto L3
	} else {
		goto L4
	}
L2:
	;
	v13 = int32(_a44)
	v15 = *(*int32)(unsafe.Add(mBase, _consts[850]))
	*(*int32)(unsafe.Add(mBase, _consts[850])) = v15 + int32(1)
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+204))
	v20 = v19
	goto L1
L3:
	;
	return
L4:
	;
	v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+100))
	*(*int64)(unsafe.Add(mBase, uint32(v28)+16)) = l1
	v31 = *(*int32)(unsafe.Add(mBase, _consts[951]))
	if v31 != 0 {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	v46 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if v46&int32(16) == int32(0) {
		goto L10
	} else {
		goto L11
	}
L6:
	;
	v33 = F_raxNew(m)
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
		goto L3
	} else {
		goto L7
	}
L7:
	;
	*(*int32)(unsafe.Add(mBase, _consts[951])) = v33
	v37 = F_raxNew(m)
	mBase = m.M
	v38 = m.ExcPending
	if v38 != 0 {
		goto L3
	} else {
		goto L8
	}
L8:
	;
	*(*int32)(unsafe.Add(mBase, _consts[950])) = v37
	v43 = F_createStringObject_1(m, int32(_a1751), int32(20))
	mBase = m.M
	v44 = m.ExcPending
	if v44 != 0 {
		goto L3
	} else {
		goto L9
	}
L9:
	;
	*(*int32)(unsafe.Add(mBase, _consts[952])) = v43
	goto L5
L10:
	;
	v111 = *(*int32)(unsafe.Add(mBase, uint32(l0)+204))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+204)) = v111&int32(-353) | v46&int32(352)
	return
L11:
	;
	v51 = *(*int32)(unsafe.Add(mBase, uint32(l0)+204))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+204)) = v51 | int32(16)
	if l4 == int32(0) {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	F_enableBcastTrackingForPrefix(m, l0, int32(_a139), int32(0))
	mBase = m.M
	v101 = m.ExcPending
	if v101 != 0 {
		goto L3
	} else {
		goto L24
	}
L13:
	;
	v63 = int32(0)
	goto L14
L14:
	;
	v71 = *(*int32)(unsafe.Add(mBase, uint32(l3+v63<<(uint(int32(2))%32))))
	v72 = F_objectGetVal(m, v71)
	mBase = m.M
	v75 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v72+int32(-1)))))
	switch v75 & int32(7) {
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
		v92 = int32(0)
		goto L16
	}
L16:
	;
	F_enableBcastTrackingForPrefix(m, l0, v72, v92)
	mBase = m.M
	v94 = m.ExcPending
	if v94 != 0 {
		goto L3
	} else {
		goto L22
	}
L17:
	;
	v91 = *(*int32)(unsafe.Add(mBase, uint32(v72+int32(-17))))
	v92 = v91
	goto L16
L18:
	;
	v88 = *(*int32)(unsafe.Add(mBase, uint32(v72+int32(-9))))
	v92 = v88
	goto L16
L19:
	;
	v85 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v72+int32(-5)))))
	v92 = v85
	goto L16
L20:
	;
	v82 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v72+int32(-3)))))
	v92 = v82
	goto L16
L21:
	;
	v92 = int32(base.Ui32(v75) >> (uint(int32(3)) % 32))
	goto L16
L22:
	;
	v96 = v63 + int32(1)
	if v96 != l4 {
		v63 = v96
		goto L14
	} else {
		goto L23
	}
L23:
	;
	goto L10
L24:
	;
	goto L10
}
func F_freeTrackingRadixTree(m *base.Module, l0 int32) {
	var v4 int32
	_ = v4
	F_raxFreeWithCallback(m, l0, int32(1105))
	v4 = m.ExcPending
	if v4 != 0 {
		return
	} else {
		return
	}
}
func F_freeTrackingRadixTreeAsync(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v7 int64
	_ = v7
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	v3 = m.G0
	v5 = v3 - int32(16)
	m.G0 = v5
	v7 = *(*int64)(unsafe.Add(mBase, uint32(l0)+16))
	if base.Ui64(v7) < base.Ui64(int64(65)) {
		F_freeTrackingRadixTree(m, l0)
		mBase = m.M
		v22 = m.ExcPending
		if v22 != 0 {
			return
		} else {
			m.G0 = v5 + int32(16)
			return
		}
	} else {
		v10 = int32(0)
		v12 = *(*int32)(unsafe.Add(mBase, _consts[503]))
		v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		*(*int32)(unsafe.Add(mBase, _consts[503])) = v12 + v13
		*(*int32)(unsafe.Add(mBase, uint32(v5))) = l0
		F_bioCreateLazyFreeJob(m, int32(553), int32(1), v5)
		mBase = m.M
		v20 = m.ExcPending
		if v20 != 0 {
			return
		} else {
			m.G0 = v5 + int32(16)
			return
		}
	}
}
func F_trackingBuildBroadcastReply(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int64
	_ = v10
	var v12 int32
	_ = v12
	var v18 int64
	_ = v18
	var v30 int32
	_ = v30
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
	var v52 int64
	_ = v52
	var v54 int32
	_ = v54
	var v57 int64
	_ = v57
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v65 int32
	_ = v65
	var v74 int64
	_ = v74
	var v85 int32
	_ = v85
	var v89 int32
	_ = v89
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v96 int64
	_ = v96
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v136 int32
	_ = v136
	var v142 int64
	_ = v142
	var v157 int32
	_ = v157
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v165 int32
	_ = v165
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v174 int32
	_ = v174
	var v177 int64
	_ = v177
	var v186 int32
	_ = v186
	var v190 int32
	_ = v190
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v197 int64
	_ = v197
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v218 int32
	_ = v218
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
	var v240 int32
	_ = v240
	var v245 int32
	_ = v245
	v6 = m.G0
	v8 = v6 - int32(336)
	m.G0 = v8
	if l0 != 0 {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	m.G0 = v8 + int32(336)
	return v245
L2:
	;
	if v74 <= int64(-1) {
		goto L23
	} else {
		goto L24
	}
L3:
	;
	v12 = v8 + int32(32)
	*(*int32)(unsafe.Add(mBase, uint32(v12)+4)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v12))) = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v12)+20)) = int32(128)
	v18 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v12)+12)) = v18
	*(*int64)(unsafe.Add(mBase, uint32(v12)+296)) = v18
	*(*int64)(unsafe.Add(mBase, uint32(v12)+160)) = int64(137438953472)
	*(*int32)(unsafe.Add(mBase, uint32(v12)+8)) = v8 + int32(56)
	*(*int32)(unsafe.Add(mBase, uint32(v12)+156)) = v8 + int32(200)
	goto L6
L4:
	;
	v10 = *(*int64)(unsafe.Add(mBase, uint32(l1)+8))
	goto L5
L5:
	;
	v74 = v10
	goto L2
L6:
	;
	v30 = int32(0)
	v36 = F_raxSeek(m, v8+int32(32), int32(_a263), v30, v30)
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	return int32(0)
L8:
	;
	v42 = F_raxNext(m, v8+int32(32))
	mBase = m.M
	v43 = m.ExcPending
	if v43 != 0 {
		goto L7
	} else {
		goto L10
	}
L9:
	;
	v52 = int64(0)
	goto L13
L10:
	;
	if v42 != 0 {
		goto L9
	} else {
		goto L11
	}
L11:
	;
	F_raxStop(m, v8+int32(32))
	mBase = m.M
	v47 = m.ExcPending
	if v47 != 0 {
		goto L7
	} else {
		goto L12
	}
L12:
	;
	v245 = v30
	goto L1
L13:
	;
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v8)+44))
	v57 = v52 + base.I64_extend_i32_u(base.B2i32(v54 != l0))
	v60 = F_raxNext(m, v8+int32(32))
	mBase = m.M
	v61 = m.ExcPending
	if v61 != 0 {
		goto L7
	} else {
		goto L15
	}
L14:
	;
	F_raxStop(m, v8+int32(32))
	mBase = m.M
	v65 = m.ExcPending
	if v65 != 0 {
		goto L7
	} else {
		goto L17
	}
L15:
	;
	if v60 != 0 {
		v52 = v57
		goto L13
	} else {
		goto L16
	}
L16:
	;
	goto L14
L17:
	;
	if base.B2i32(v57 == int64(0)) == int32(0) {
		v74 = v57
		goto L2
	} else {
		goto L18
	}
L18:
	;
	v245 = int32(0)
	goto L1
L19:
	;
	v118 = F_sdsempty(m)
	mBase = m.M
	v119 = m.ExcPending
	if v119 != 0 {
		goto L7
	} else {
		goto L28
	}
L20:
	;
	v117 = int32(0)
	goto L19
L22:
	;
	v98 = F_ull2string(m, v94, v95, v96)
	mBase = m.M
	if v98 == int32(0) {
		goto L20
	} else {
		goto L26
	}
L23:
	;
	goto L25
L24:
	;
	v94 = v8
	v95 = int32(32)
	v96 = v74
	v97 = int32(0)
	goto L22
L25:
	;
	v85 = int32(45)
	*(*uint8)(unsafe.Add(mBase, uint32(v8))) = uint8(v85)
	v89 = int32(1)
	v94 = v8 + v89
	v95 = int32(31)
	v96 = int64(0) - v74
	v97 = v89
	goto L22
L26:
	;
	v117 = v98 + v97
	goto L19
L28:
	;
	v123 = F_sdsMakeRoomFor(m, v118, base.I32_wrap_i64(v74)*int32(15))
	mBase = m.M
	v124 = m.ExcPending
	if v124 != 0 {
		goto L7
	} else {
		goto L29
	}
L29:
	;
	v127 = F_sdscatlen(m, v123, int32(_a1754), int32(1))
	mBase = m.M
	v128 = m.ExcPending
	if v128 != 0 {
		goto L7
	} else {
		goto L30
	}
L30:
	;
	v129 = F_sdscatlen(m, v127, v8, v117)
	mBase = m.M
	v130 = m.ExcPending
	if v130 != 0 {
		goto L7
	} else {
		goto L31
	}
L31:
	;
	v133 = F_sdscatlen(m, v129, int32(_a132), int32(2))
	mBase = m.M
	v134 = m.ExcPending
	if v134 != 0 {
		goto L7
	} else {
		goto L32
	}
L32:
	;
	v136 = v8 + int32(32)
	*(*int32)(unsafe.Add(mBase, uint32(v136)+4)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v136))) = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v136)+20)) = int32(128)
	v142 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v136)+12)) = v142
	*(*int64)(unsafe.Add(mBase, uint32(v136)+296)) = v142
	*(*int64)(unsafe.Add(mBase, uint32(v136)+160)) = int64(137438953472)
	*(*int32)(unsafe.Add(mBase, uint32(v136)+8)) = v8 + int32(56)
	*(*int32)(unsafe.Add(mBase, uint32(v136)+156)) = v8 + int32(200)
	goto L33
L33:
	;
	v157 = int32(0)
	v159 = F_raxSeek(m, v8+int32(32), int32(_a263), v157, v157)
	mBase = m.M
	v160 = m.ExcPending
	if v160 != 0 {
		goto L7
	} else {
		goto L34
	}
L34:
	;
	v165 = v133
	goto L36
L35:
	;
	F_raxStop(m, v8+int32(32))
	mBase = m.M
	v240 = m.ExcPending
	if v240 != 0 {
		goto L7
	} else {
		goto L57
	}
L36:
	;
	v168 = F_raxNext(m, v8+int32(32))
	mBase = m.M
	v169 = m.ExcPending
	if v169 != 0 {
		goto L7
	} else {
		goto L38
	}
L38:
	;
	if v168 == int32(0) {
		goto L35
	} else {
		goto L39
	}
L39:
	;
	if l0 == int32(0) {
		goto L40
	} else {
		goto L41
	}
L40:
	;
	v177 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v8)+48)))
	if v177 <= int64(-1) {
		goto L47
	} else {
		goto L48
	}
L41:
	;
	v174 = *(*int32)(unsafe.Add(mBase, uint32(v8)+44))
	if v174 == l0 {
		goto L36
	} else {
		goto L42
	}
L42:
	;
	goto L40
L43:
	;
	v221 = F_sdscatlen(m, v165, int32(_a1695), int32(1))
	mBase = m.M
	v222 = m.ExcPending
	if v222 != 0 {
		goto L7
	} else {
		goto L52
	}
L44:
	;
	v218 = int32(0)
	goto L43
L46:
	;
	v199 = F_ull2string(m, v195, v196, v197)
	mBase = m.M
	if v199 == int32(0) {
		goto L44
	} else {
		goto L50
	}
L47:
	;
	goto L49
L48:
	;
	v195 = v8
	v196 = int32(32)
	v197 = v177
	v198 = int32(0)
	goto L46
L49:
	;
	v186 = int32(45)
	*(*uint8)(unsafe.Add(mBase, uint32(v8))) = uint8(v186)
	v190 = int32(1)
	v195 = v8 + v190
	v196 = int32(31)
	v197 = int64(0) - v177
	v198 = v190
	goto L46
L50:
	;
	v218 = v199 + v198
	goto L43
L52:
	;
	v223 = F_sdscatlen(m, v221, v8, v218)
	mBase = m.M
	v224 = m.ExcPending
	if v224 != 0 {
		goto L7
	} else {
		goto L53
	}
L53:
	;
	v227 = F_sdscatlen(m, v223, int32(_a132), int32(2))
	mBase = m.M
	v228 = m.ExcPending
	if v228 != 0 {
		goto L7
	} else {
		goto L54
	}
L54:
	;
	v229 = *(*int32)(unsafe.Add(mBase, uint32(v8)+40))
	v230 = *(*int32)(unsafe.Add(mBase, uint32(v8)+48))
	v231 = F_sdscatlen(m, v227, v229, v230)
	mBase = m.M
	v232 = m.ExcPending
	if v232 != 0 {
		goto L7
	} else {
		goto L55
	}
L55:
	;
	v235 = F_sdscatlen(m, v231, int32(_a132), int32(2))
	mBase = m.M
	v236 = m.ExcPending
	if v236 != 0 {
		goto L7
	} else {
		goto L56
	}
L56:
	;
	v165 = v235
	goto L36
L57:
	;
	v245 = v165
	goto L1
}
func F_trackingGetTotalItems(m *base.Module) int64 {
	mBase := m.M
	_ = mBase
	var v2 int64
	_ = v2
	v2 = *(*int64)(unsafe.Add(mBase, _consts[953]))
	return v2
}
func F_trackingGetTotalKeys(m *base.Module) int64 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v6 int64
	_ = v6
	v3 = *(*int32)(unsafe.Add(mBase, _consts[951]))
	if v3 != 0 {
		v6 = *(*int64)(unsafe.Add(mBase, uint32(v3)+8))
		return v6
	} else {
		return int64(0)
	}
}
func F_trackingInvalidateKey(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v47 int64
	_ = v47
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v92 int32
	_ = v92
	var v102 int32
	_ = v102
	var v105 int32
	_ = v105
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v115 int32
	_ = v115
	var v125 int32
	_ = v125
	var v128 int32
	_ = v128
	var v133 int32
	_ = v133
	var v136 int32
	_ = v136
	var v149 int32
	_ = v149
	var v151 int32
	_ = v151
	var v153 int32
	_ = v153
	var v158 int32
	_ = v158
	var v165 int32
	_ = v165
	var v170 int32
	_ = v170
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v192 int32
	_ = v192
	var v195 int32
	_ = v195
	var v199 int32
	_ = v199
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v208 int32
	_ = v208
	var v212 int32
	_ = v212
	var v214 int32
	_ = v214
	var v220 int32
	_ = v220
	var v224 int32
	_ = v224
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v241 int32
	_ = v241
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v249 int32
	_ = v249
	var v254 int32
	_ = v254
	var v255 int32
	_ = v255
	var v261 int64
	_ = v261
	var v276 int32
	_ = v276
	var v278 int32
	_ = v278
	var v279 int32
	_ = v279
	var v282 int32
	_ = v282
	var v283 int32
	_ = v283
	var v297 int32
	_ = v297
	var v298 int64
	_ = v298
	var v301 int32
	_ = v301
	var v302 int32
	_ = v302
	var v303 int32
	_ = v303
	var v305 int64
	_ = v305
	var v307 int64
	_ = v307
	var v309 int64
	_ = v309
	var v312 int64
	_ = v312
	var v314 int64
	_ = v314
	var v316 int64
	_ = v316
	var v318 int64
	_ = v318
	var v344 int32
	_ = v344
	var v345 int32
	_ = v345
	var v350 int32
	_ = v350
	var v351 int32
	_ = v351
	var v357 int32
	_ = v357
	var v363 int32
	_ = v363
	var v370 int32
	_ = v370
	var v376 int32
	_ = v376
	var v378 int32
	_ = v378
	var v379 int32
	_ = v379
	var v380 int32
	_ = v380
	var v381 int32
	_ = v381
	var v383 int32
	_ = v383
	var v386 int32
	_ = v386
	var v393 int32
	_ = v393
	var v396 int32
	_ = v396
	var v399 int32
	_ = v399
	var v402 int32
	_ = v402
	var v403 int32
	_ = v403
	var v406 int32
	_ = v406
	var v413 int32
	_ = v413
	var v414 int32
	_ = v414
	var v429 int32
	_ = v429
	var v430 int64
	_ = v430
	var v431 int32
	_ = v431
	var v433 int64
	_ = v433
	var v437 int32
	_ = v437
	var v438 int32
	_ = v438
	var v439 int32
	_ = v439
	var v441 int32
	_ = v441
	var v442 int32
	_ = v442
	v4 = int32(0)
	v12 = m.G0
	v14 = v12 - int32(320)
	m.G0 = v14
	v18 = *(*int32)(unsafe.Add(mBase, _consts[951]))
	if v18 == v4 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	m.G0 = v14 + int32(320)
	return
L2:
	;
	v21 = F_objectGetVal(m, l1)
	mBase = m.M
	v22 = F_objectGetVal(m, l1)
	mBase = m.M
	v25 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22+int32(-1)))))
	switch v25 & int32(7) {
	case 0:
		goto L8
	case 1:
		goto L7
	case 2:
		goto L6
	case 3:
		goto L5
	case 4:
		goto L4
	default:
		v42 = v4
		goto L3
	}
L3:
	;
	if l2 == int32(0) {
		goto L9
	} else {
		goto L10
	}
L4:
	;
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v22+int32(-17))))
	v42 = v41
	goto L3
L5:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v22+int32(-9))))
	v42 = v38
	goto L3
L6:
	;
	v35 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v22+int32(-5)))))
	v42 = v35
	goto L3
L7:
	;
	v32 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22+int32(-3)))))
	v42 = v32
	goto L3
L8:
	;
	v42 = int32(base.Ui32(v25) >> (uint(int32(3)) % 32))
	goto L3
L9:
	;
	v52 = int32(0)
	v53 = *(*int32)(unsafe.Add(mBase, _consts[951]))
	v55 = v14 + int32(316)
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v53)))
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v64)))
	if v42 == v52 {
		goto L17
	} else {
		goto L18
	}
L10:
	;
	v46 = *(*int32)(unsafe.Add(mBase, _consts[950]))
	v47 = *(*int64)(unsafe.Add(mBase, uint32(v46)+8))
	goto L11
L11:
	;
	if v47 == int64(0) {
		goto L9
	} else {
		goto L12
	}
L12:
	;
	F_trackingRememberKeyToBroadcast(m, l0, v21, v42)
	mBase = m.M
	v51 = m.ExcPending
	if v51 != 0 {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	return
L14:
	;
	goto L9
L15:
	;
	if v249 == int32(0) {
		goto L1
	} else {
		goto L53
	}
L16:
	;
	if v208 != v42 {
		v249 = v52
		goto L43
	} else {
		goto L44
	}
L17:
	;
	v199 = int32(0)
	v205 = v64
	v206 = v65
	v208 = v199
	v212 = v199
	goto L16
L18:
	;
	if base.Ui32(v65) < base.Ui32(int32(8)) {
		goto L17
	} else {
		goto L19
	}
L19:
	;
	v76 = v64
	v77 = v65
	v79 = int32(0)
	goto L21
L20:
	;
	v205 = v189
	v206 = v190
	v208 = v192
	v212 = base.B2i32(v195 != int32(0))
	goto L16
L21:
	;
	v85 = int32(base.Ui32(v77) >> (uint(int32(3)) % 32))
	v86 = int32(4)
	v87 = v76 + v86
	if v77&v86 == int32(0) {
		goto L24
	} else {
		goto L25
	}
L22:
	;
	v189 = v180
	v190 = v181
	v192 = v165
	v195 = v170
	goto L20
L23:
	;
	v170 = int32(0)
	v180 = *(*int32)(unsafe.Add(mBase, uint32(v87+v85+(v170-v85)&int32(3)+v158<<(uint(int32(2))%32))))
	v181 = *(*int32)(unsafe.Add(mBase, uint32(v180)))
	if base.Ui32(v181) < base.Ui32(int32(8)) {
		v189 = v180
		v190 = v181
		v192 = v165
		v195 = v170
		goto L20
	} else {
		goto L41
	}
L24:
	;
	v133 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21+v79))))
	v136 = int32(0)
	goto L35
L25:
	;
	v92 = int32(0)
	if base.Ui32(v42) <= base.Ui32(v79) {
		v125 = v79
		v128 = v92
		goto L26
	} else {
		goto L27
	}
L26:
	;
	if v128 == v85 {
		v158 = v92
		v165 = v125
		goto L23
	} else {
		goto L33
	}
L27:
	;
	v102 = v79
	v105 = v92
	goto L28
L28:
	;
	v108 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v87+v105))))
	v110 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21+v102))))
	if v108 != v110 {
		v125 = v102
		v128 = v105
		goto L26
	} else {
		goto L30
	}
L29:
	;
	v125 = v113
	v128 = v115
	goto L26
L30:
	;
	v112 = int32(1)
	v113 = v102 + v112
	v115 = v105 + v112
	if base.Ui32(v85) <= base.Ui32(v115) {
		v125 = v113
		v128 = v115
		goto L26
	} else {
		goto L31
	}
L31:
	;
	if base.Ui32(v113) < base.Ui32(v42) {
		v102 = v113
		v105 = v115
		goto L28
	} else {
		goto L32
	}
L32:
	;
	goto L29
L33:
	;
	v189 = v76
	v190 = v77
	v192 = v125
	v195 = v128
	goto L20
L34:
	;
	if v136 != v85 {
		goto L39
	} else {
		goto L40
	}
L35:
	;
	v149 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v87+v136))))
	if v149 == v133&int32(255) {
		goto L34
	} else {
		goto L37
	}
L37:
	;
	v151 = int32(1)
	v153 = v136 + v151
	if v153 != v85 {
		v136 = v153
		goto L35
	} else {
		goto L38
	}
L38:
	;
	v205 = v76
	v206 = v77
	v208 = v79
	v212 = v151
	goto L16
L39:
	;
	v158 = v136
	v165 = v79 + int32(1)
	goto L23
L40:
	;
	v189 = v76
	v190 = v77
	v192 = v79
	v195 = v85
	goto L20
L41:
	;
	if base.Ui32(v165) < base.Ui32(v42) {
		v76 = v180
		v77 = v181
		v79 = v165
		goto L21
	} else {
		goto L42
	}
L42:
	;
	goto L22
L43:
	;
	goto L15
L44:
	;
	v214 = int32(0)
	if v206&int32(1) == v214 {
		v249 = v214
		goto L43
	} else {
		goto L45
	}
L45:
	;
	v220 = v206 & int32(4)
	if v212&base.B2i32(v220 != int32(0)) != 0 {
		v249 = v214
		goto L43
	} else {
		goto L46
	}
L46:
	;
	v224 = int32(1)
	if v55 == int32(0) {
		v249 = v224
		goto L43
	} else {
		goto L47
	}
L47:
	;
	if v206&int32(2) != 0 {
		v246 = int32(0)
		goto L48
	} else {
		goto L49
	}
L48:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v55))) = v246
	v249 = v224
	goto L43
L49:
	;
	v230 = int32(3)
	v231 = int32(base.Ui32(v206) >> (uint(v230) % 32))
	if v220 != 0 {
		goto L50
	} else {
		goto L51
	}
L50:
	;
	v241 = int32(4)
	goto L52
L51:
	;
	v241 = v231 << (uint(int32(2)) % 32)
	goto L52
L52:
	;
	v245 = *(*int32)(unsafe.Add(mBase, uint32(v205+v231+(int32(0)-v231)&v230+v241+int32(4))))
	v246 = v245
	goto L48
L53:
	;
	v254 = v14 + int32(12)
	v255 = *(*int32)(unsafe.Add(mBase, uint32(v14)+316))
	*(*int32)(unsafe.Add(mBase, uint32(v254)+4)) = v255
	*(*int32)(unsafe.Add(mBase, uint32(v254))) = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v254)+20)) = int32(128)
	v261 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v254)+12)) = v261
	*(*int64)(unsafe.Add(mBase, uint32(v254)+296)) = v261
	*(*int64)(unsafe.Add(mBase, uint32(v254)+160)) = int64(137438953472)
	*(*int32)(unsafe.Add(mBase, uint32(v254)+8)) = v14 + int32(36)
	*(*int32)(unsafe.Add(mBase, uint32(v254)+156)) = v14 + int32(180)
	goto L54
L54:
	;
	v276 = int32(0)
	v278 = F_raxSeek(m, v14+int32(12), int32(_a263), v276, v276)
	mBase = m.M
	v279 = m.ExcPending
	if v279 != 0 {
		goto L13
	} else {
		goto L55
	}
L55:
	;
	v282 = F_raxNext(m, v14+int32(12))
	mBase = m.M
	v283 = m.ExcPending
	if v283 != 0 {
		goto L13
	} else {
		goto L57
	}
L56:
	;
	F_raxStop(m, v14+int32(12))
	mBase = m.M
	v429 = m.ExcPending
	if v429 != 0 {
		goto L13
	} else {
		goto L82
	}
L57:
	;
	if v282 == int32(0) {
		goto L56
	} else {
		goto L58
	}
L58:
	;
	goto L59
L59:
	;
	v297 = *(*int32)(unsafe.Add(mBase, uint32(v14)+20))
	v298 = *(*int64)(unsafe.Add(mBase, uint32(v297)))
	v301 = m.G0
	v302 = int32(16)
	v303 = v301 - v302
	m.G0 = v303
	v305 = int64(56)
	v307 = int64(65280)
	v309 = int64(40)
	v312 = int64(16711680)
	v314 = int64(24)
	v316 = int64(4278190080)
	v318 = int64(8)
	*(*int64)(unsafe.Add(mBase, uint32(v303)+8)) = v298<<(uint(v305)%64) | v298&v307<<(uint(v309)%64) | (v298&v312<<(uint(v314)%64) | v298&v316<<(uint(v318)%64)) | (int64(base.Ui64(v298)>>(uint(v318)%64))&v316 | int64(base.Ui64(v298)>>(uint(v314)%64))&v312 | (int64(base.Ui64(v298)>>(uint(v309)%64))&v307 | int64(base.Ui64(v298)>>(uint(v305)%64))))
	*(*int32)(unsafe.Add(mBase, uint32(v303)+4)) = int32(0)
	v344 = *(*int32)(unsafe.Add(mBase, _consts[404]))
	v345 = int32(8)
	v350 = F_raxFind(m, v344, v303+v345, v345, v303+int32(4))
	mBase = m.M
	v351 = *(*int32)(unsafe.Add(mBase, uint32(v303)+4))
	m.G0 = v303 + v302
	goto L62
L60:
	;
	goto L56
L61:
	;
	v413 = F_raxNext(m, v14+int32(12))
	mBase = m.M
	v414 = m.ExcPending
	if v414 != 0 {
		goto L13
	} else {
		goto L80
	}
L62:
	;
	if v351 == int32(0) {
		goto L61
	} else {
		goto L63
	}
L63:
	;
	v357 = *(*int32)(unsafe.Add(mBase, uint32(v351)+204))
	if v357&int32(20) != int32(4) {
		goto L61
	} else {
		goto L64
	}
L64:
	;
	v363 = *(*int32)(unsafe.Add(mBase, _consts[67]))
	if v357&int32(256) == int32(0) {
		goto L65
	} else {
		goto L66
	}
L65:
	;
	if v351 != v363 {
		goto L68
	} else {
		goto L69
	}
L66:
	;
	if v351 == v363 {
		goto L61
	} else {
		goto L67
	}
L67:
	;
	goto L65
L68:
	;
	v381 = F_objectGetVal(m, l1)
	mBase = m.M
	v383 = F_objectGetVal(m, l1)
	mBase = m.M
	v386 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v383+int32(-1)))))
	switch v386 & int32(7) {
	case 0:
		goto L78
	case 1:
		goto L77
	case 2:
		goto L76
	case 3:
		goto L75
	case 4:
		goto L74
	default:
		v403 = int32(0)
		goto L73
	}
L69:
	;
	v370 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v363)+204)))
	if v370&int32(1) == int32(0) {
		goto L68
	} else {
		goto L70
	}
L70:
	;
	F_incrRefCount(m, l1)
	mBase = m.M
	v376 = m.ExcPending
	if v376 != 0 {
		goto L13
	} else {
		goto L71
	}
L71:
	;
	v378 = *(*int32)(unsafe.Add(mBase, _consts[878]))
	v379 = F_listAddNodeTail(m, v378, l1)
	mBase = m.M
	v380 = m.ExcPending
	if v380 != 0 {
		goto L13
	} else {
		goto L72
	}
L72:
	;
	goto L61
L73:
	;
	F_sendTrackingMessage(m, v351, v381, v403, int32(0))
	mBase = m.M
	v406 = m.ExcPending
	if v406 != 0 {
		goto L13
	} else {
		goto L79
	}
L74:
	;
	v402 = *(*int32)(unsafe.Add(mBase, uint32(v383+int32(-17))))
	v403 = v402
	goto L73
L75:
	;
	v399 = *(*int32)(unsafe.Add(mBase, uint32(v383+int32(-9))))
	v403 = v399
	goto L73
L76:
	;
	v396 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v383+int32(-5)))))
	v403 = v396
	goto L73
L77:
	;
	v393 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v383+int32(-3)))))
	v403 = v393
	goto L73
L78:
	;
	v403 = int32(base.Ui32(v386) >> (uint(int32(3)) % 32))
	goto L73
L79:
	;
	goto L61
L80:
	;
	if v413 != 0 {
		goto L59
	} else {
		goto L81
	}
L81:
	;
	goto L60
L82:
	;
	v430 = *(*int64)(unsafe.Add(mBase, uint32(v255)+8))
	goto L83
L83:
	;
	v431 = int32(0)
	v433 = *(*int64)(unsafe.Add(mBase, _consts[953]))
	*(*int64)(unsafe.Add(mBase, _consts[953])) = v433 - v430
	F_raxFree(m, v255)
	mBase = m.M
	v437 = m.ExcPending
	if v437 != 0 {
		goto L13
	} else {
		goto L84
	}
L84:
	;
	v438 = int32(0)
	v439 = *(*int32)(unsafe.Add(mBase, _consts[951]))
	v441 = F_raxRemove(m, v439, v21, v42, v438)
	mBase = m.M
	v442 = m.ExcPending
	if v442 != 0 {
		goto L13
	} else {
		goto L85
	}
L85:
	;
	goto L1
}
func F_trackingRememberKeys(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v21 int32
	_ = v21
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v45 int32
	_ = v45
	var v51 int32
	_ = v51
	var v58 int32
	_ = v58
	var v62 int32
	_ = v62
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v77 int32
	_ = v77
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
	var v89 int32
	_ = v89
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v113 int32
	_ = v113
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v126 int32
	_ = v126
	var v136 int32
	_ = v136
	var v139 int32
	_ = v139
	var v142 int32
	_ = v142
	var v144 int32
	_ = v144
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v149 int32
	_ = v149
	var v159 int32
	_ = v159
	var v162 int32
	_ = v162
	var v167 int32
	_ = v167
	var v170 int32
	_ = v170
	var v183 int32
	_ = v183
	var v185 int32
	_ = v185
	var v187 int32
	_ = v187
	var v192 int32
	_ = v192
	var v199 int32
	_ = v199
	var v204 int32
	_ = v204
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v226 int32
	_ = v226
	var v229 int32
	_ = v229
	var v233 int32
	_ = v233
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v242 int32
	_ = v242
	var v246 int32
	_ = v246
	var v248 int32
	_ = v248
	var v254 int32
	_ = v254
	var v258 int32
	_ = v258
	var v264 int32
	_ = v264
	var v265 int32
	_ = v265
	var v275 int32
	_ = v275
	var v279 int32
	_ = v279
	var v280 int32
	_ = v280
	var v283 int32
	_ = v283
	var v285 int32
	_ = v285
	var v286 int32
	_ = v286
	var v287 int32
	_ = v287
	var v289 int32
	_ = v289
	var v290 int32
	_ = v290
	var v297 int32
	_ = v297
	var v300 int32
	_ = v300
	var v303 int32
	_ = v303
	var v306 int32
	_ = v306
	var v307 int32
	_ = v307
	var v309 int32
	_ = v309
	var v310 int32
	_ = v310
	var v317 int32
	_ = v317
	var v319 int32
	_ = v319
	var v320 int32
	_ = v320
	var v325 int32
	_ = v325
	var v327 int32
	_ = v327
	var v328 int32
	_ = v328
	var v331 int32
	_ = v331
	var v333 int64
	_ = v333
	var v338 int32
	_ = v338
	var v354 int32
	_ = v354
	v12 = m.G0
	v14 = v12 - int32(2064)
	m.G0 = v14
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+204))
	if v16&int32(160) == int32(32) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	m.G0 = v14 + int32(2064)
	return
L2:
	;
	v21 = int32(192)
	if v16&v21 == v21 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+12)) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v14)+4)) = int64(1099511627776)
	v29 = *(*int32)(unsafe.Add(mBase, uint32(l1)+68))
	v30 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	v31 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
	v34 = F_getKeysFromCommand(m, v29, v30, v31, v14+int32(4))
	mBase = m.M
	v35 = m.ExcPending
	if v35 != 0 {
		goto L5
	} else {
		goto L6
	}
L4:
	;
	F_getKeysFreeResult(m, v14+int32(4))
	mBase = m.M
	v354 = m.ExcPending
	if v354 != 0 {
		goto L5
	} else {
		goto L73
	}
L5:
	;
	return
L6:
	;
	if v34 == int32(0) {
		goto L4
	} else {
		goto L7
	}
L7:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(l1)+68))
	v39 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v38)+56)))
	if v39&int32(32) != 0 {
		goto L1
	} else {
		goto L8
	}
L8:
	;
	if v34 < int32(1) {
		goto L4
	} else {
		goto L9
	}
L9:
	;
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v14)+12))
	v51 = int32(0)
	goto L10
L10:
	;
	v58 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v45+v51<<(uint(int32(3))%32))))
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v58+v62<<(uint(int32(2))%32))))
	v67 = F_objectGetVal(m, v66)
	mBase = m.M
	v69 = v67 + int32(-1)
	v70 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v69))))
	switch v70 & int32(7) {
	case 0:
		goto L17
	case 1:
		goto L16
	case 2:
		goto L15
	case 3:
		goto L14
	case 4:
		goto L13
	default:
		v87 = int32(0)
		goto L12
	}
L11:
	;
	goto L4
L12:
	;
	v88 = int32(0)
	v89 = *(*int32)(unsafe.Add(mBase, _consts[951]))
	v98 = *(*int32)(unsafe.Add(mBase, uint32(v89)))
	v99 = *(*int32)(unsafe.Add(mBase, uint32(v98)))
	if v87 == v88 {
		goto L22
	} else {
		goto L23
	}
L13:
	;
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v67+int32(-17))))
	v87 = v86
	goto L12
L14:
	;
	v83 = *(*int32)(unsafe.Add(mBase, uint32(v67+int32(-9))))
	v87 = v83
	goto L12
L15:
	;
	v80 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v67+int32(-5)))))
	v87 = v80
	goto L12
L16:
	;
	v77 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v67+int32(-3)))))
	v87 = v77
	goto L12
L17:
	;
	v87 = int32(base.Ui32(v70) >> (uint(int32(3)) % 32))
	goto L12
L18:
	;
	v325 = int32(0)
	v327 = F_raxTryInsert(m, v320, l0, int32(8), v325, v325)
	mBase = m.M
	v328 = m.ExcPending
	if v328 != 0 {
		goto L5
	} else {
		goto L70
	}
L19:
	;
	v319 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
	v320 = v319
	goto L18
L20:
	;
	if v283 != 0 {
		goto L19
	} else {
		goto L58
	}
L21:
	;
	if v242 != v87 {
		v283 = v88
		goto L48
	} else {
		goto L49
	}
L22:
	;
	v233 = int32(0)
	v239 = v98
	v240 = v99
	v242 = v233
	v246 = v233
	goto L21
L23:
	;
	if base.Ui32(v99) < base.Ui32(int32(8)) {
		goto L22
	} else {
		goto L24
	}
L24:
	;
	v110 = v98
	v111 = v99
	v113 = int32(0)
	goto L26
L25:
	;
	v239 = v223
	v240 = v224
	v242 = v226
	v246 = base.B2i32(v229 != int32(0))
	goto L21
L26:
	;
	v119 = int32(base.Ui32(v111) >> (uint(int32(3)) % 32))
	v120 = int32(4)
	v121 = v110 + v120
	if v111&v120 == int32(0) {
		goto L29
	} else {
		goto L30
	}
L27:
	;
	v223 = v214
	v224 = v215
	v226 = v199
	v229 = v204
	goto L25
L28:
	;
	v204 = int32(0)
	v214 = *(*int32)(unsafe.Add(mBase, uint32(v121+v119+(v204-v119)&int32(3)+v192<<(uint(int32(2))%32))))
	v215 = *(*int32)(unsafe.Add(mBase, uint32(v214)))
	if base.Ui32(v215) < base.Ui32(int32(8)) {
		v223 = v214
		v224 = v215
		v226 = v199
		v229 = v204
		goto L25
	} else {
		goto L46
	}
L29:
	;
	v167 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v67+v113))))
	v170 = int32(0)
	goto L40
L30:
	;
	v126 = int32(0)
	if base.Ui32(v87) <= base.Ui32(v113) {
		v159 = v113
		v162 = v126
		goto L31
	} else {
		goto L32
	}
L31:
	;
	if v162 == v119 {
		v192 = v126
		v199 = v159
		goto L28
	} else {
		goto L38
	}
L32:
	;
	v136 = v113
	v139 = v126
	goto L33
L33:
	;
	v142 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v121+v139))))
	v144 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v67+v136))))
	if v142 != v144 {
		v159 = v136
		v162 = v139
		goto L31
	} else {
		goto L35
	}
L34:
	;
	v159 = v147
	v162 = v149
	goto L31
L35:
	;
	v146 = int32(1)
	v147 = v136 + v146
	v149 = v139 + v146
	if base.Ui32(v119) <= base.Ui32(v149) {
		v159 = v147
		v162 = v149
		goto L31
	} else {
		goto L36
	}
L36:
	;
	if base.Ui32(v147) < base.Ui32(v87) {
		v136 = v147
		v139 = v149
		goto L33
	} else {
		goto L37
	}
L37:
	;
	goto L34
L38:
	;
	v223 = v110
	v224 = v111
	v226 = v159
	v229 = v162
	goto L25
L39:
	;
	if v170 != v119 {
		goto L44
	} else {
		goto L45
	}
L40:
	;
	v183 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v121+v170))))
	if v183 == v167&int32(255) {
		goto L39
	} else {
		goto L42
	}
L42:
	;
	v185 = int32(1)
	v187 = v170 + v185
	if v187 != v119 {
		v170 = v187
		goto L40
	} else {
		goto L43
	}
L43:
	;
	v239 = v110
	v240 = v111
	v242 = v113
	v246 = v185
	goto L21
L44:
	;
	v192 = v170
	v199 = v113 + int32(1)
	goto L28
L45:
	;
	v223 = v110
	v224 = v111
	v226 = v113
	v229 = v119
	goto L25
L46:
	;
	if base.Ui32(v199) < base.Ui32(v87) {
		v110 = v214
		v111 = v215
		v113 = v199
		goto L26
	} else {
		goto L47
	}
L47:
	;
	goto L27
L48:
	;
	goto L20
L49:
	;
	v248 = int32(0)
	if v240&int32(1) == v248 {
		v283 = v248
		goto L48
	} else {
		goto L50
	}
L50:
	;
	v254 = v240 & int32(4)
	if v246&base.B2i32(v254 != int32(0)) != 0 {
		v283 = v248
		goto L48
	} else {
		goto L51
	}
L51:
	;
	v258 = int32(1)
	if v14 == int32(0) {
		v283 = v258
		goto L48
	} else {
		goto L52
	}
L52:
	;
	if v240&int32(2) != 0 {
		v280 = int32(0)
		goto L53
	} else {
		goto L54
	}
L53:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14))) = v280
	v283 = v258
	goto L48
L54:
	;
	v264 = int32(3)
	v265 = int32(base.Ui32(v240) >> (uint(v264) % 32))
	if v254 != 0 {
		goto L55
	} else {
		goto L56
	}
L55:
	;
	v275 = int32(4)
	goto L57
L56:
	;
	v275 = v265 << (uint(int32(2)) % 32)
	goto L57
L57:
	;
	v279 = *(*int32)(unsafe.Add(mBase, uint32(v239+v265+(int32(0)-v265)&v264+v275+int32(4))))
	v280 = v279
	goto L53
L58:
	;
	v285 = F_raxNew(m)
	mBase = m.M
	v286 = m.ExcPending
	if v286 != 0 {
		goto L5
	} else {
		goto L59
	}
L59:
	;
	v287 = int32(0)
	v289 = *(*int32)(unsafe.Add(mBase, _consts[951]))
	v290 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v69))))
	switch v290 & int32(7) {
	case 0:
		goto L65
	case 1:
		goto L64
	case 2:
		goto L63
	case 3:
		goto L62
	case 4:
		goto L61
	default:
		v307 = v287
		goto L60
	}
L60:
	;
	v309 = F_raxTryInsert(m, v289, v67, v307, v285, int32(0))
	mBase = m.M
	v310 = m.ExcPending
	if v310 != 0 {
		goto L5
	} else {
		goto L66
	}
L61:
	;
	v306 = *(*int32)(unsafe.Add(mBase, uint32(v67+int32(-17))))
	v307 = v306
	goto L60
L62:
	;
	v303 = *(*int32)(unsafe.Add(mBase, uint32(v67+int32(-9))))
	v307 = v303
	goto L60
L63:
	;
	v300 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v67+int32(-5)))))
	v307 = v300
	goto L60
L64:
	;
	v297 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v67+int32(-3)))))
	v307 = v297
	goto L60
L65:
	;
	v307 = int32(base.Ui32(v290) >> (uint(int32(3)) % 32))
	goto L60
L66:
	;
	if v309 == int32(1) {
		v320 = v285
		goto L18
	} else {
		goto L67
	}
L67:
	;
	F__serverAssert(m, int32(_a1752), int32(_a1753), int32(254))
	mBase = m.M
	v317 = m.ExcPending
	if v317 != 0 {
		goto L5
	} else {
		goto L68
	}
L68:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L69:
	;
	v338 = v51 + int32(1)
	if v338 != v34 {
		v51 = v338
		goto L10
	} else {
		goto L72
	}
L70:
	;
	if v327 == int32(0) {
		goto L69
	} else {
		goto L71
	}
L71:
	;
	v331 = int32(0)
	v333 = *(*int64)(unsafe.Add(mBase, _consts[953]))
	*(*int64)(unsafe.Add(mBase, _consts[953])) = v333 + int64(1)
	goto L69
L72:
	;
	goto L11
L73:
	;
	goto L1
}
