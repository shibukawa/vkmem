package p2

import (
	base "github.com/shibukawa/valkeymem/internal/aot/vkaot/base"
	"unsafe"
)

func F_addNodeDetailsToShardReply(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
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
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	var v96 int32
	_ = v96
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	var v110 int32
	_ = v110
	var v117 int32
	_ = v117
	var v120 int32
	_ = v120
	var v123 int32
	_ = v123
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v129 int32
	_ = v129
	var v132 int32
	_ = v132
	var v135 int32
	_ = v135
	var v148 int32
	_ = v148
	var v152 int32
	_ = v152
	var v155 int32
	_ = v155
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v161 int64
	_ = v161
	var v162 int64
	_ = v162
	var v164 int64
	_ = v164
	var v167 int64
	_ = v167
	var v169 int64
	_ = v169
	var v170 int64
	_ = v170
	var v171 int64
	_ = v171
	var v174 int32
	_ = v174
	var v177 int32
	_ = v177
	var v180 int32
	_ = v180
	var v182 int32
	_ = v182
	var v185 int32
	_ = v185
	var v187 int32
	_ = v187
	var v190 int32
	_ = v190
	var v192 int32
	_ = v192
	var v199 int32
	_ = v199
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v208 int32
	_ = v208
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v214 int32
	_ = v214
	var v221 int32
	_ = v221
	var v224 int32
	_ = v224
	var v227 int32
	_ = v227
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v236 int32
	_ = v236
	var v238 int32
	_ = v238
	var v241 int32
	_ = v241
	var v248 int32
	_ = v248
	var v251 int32
	_ = v251
	var v254 int32
	_ = v254
	var v257 int32
	_ = v257
	var v258 int32
	_ = v258
	var v260 int32
	_ = v260
	var v264 int32
	_ = v264
	var v267 int32
	_ = v267
	v9 = F_addReplyDeferredLen(m, l0)
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	F_addReplyBulkCString(m, l0, int32(_a344))
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	F_addReplyBulkCBuffer(m, l0, l1+int32(8), int32(40))
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l1)+2324))
	if v19 != 0 {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	v35 = *(*int32)(unsafe.Add(mBase, uint32(l1)+2328))
	if v35 == int32(0) {
		v52 = v34
		goto L14
	} else {
		goto L15
	}
L6:
	;
	F_addReplyBulkCString(m, l0, int32(_a345))
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		goto L1
	} else {
		goto L8
	}
L7:
	;
	v34 = int32(1)
	goto L5
L8:
	;
	if l0 == int32(0) {
		goto L10
	} else {
		goto L11
	}
L9:
	;
	F_addReplyLongLong(m, l0, base.I64_extend_i32_s(v29))
	mBase = m.M
	v32 = m.ExcPending
	if v32 != 0 {
		goto L1
	} else {
		goto L13
	}
L10:
	;
	v28 = *(*int32)(unsafe.Add(mBase, uint32(l1)+2324))
	v29 = v28
	goto L9
L11:
	;
	v26 = *(*int32)(unsafe.Add(mBase, uint32(l1)+2336))
	if v26 != 0 {
		v29 = v26
		goto L9
	} else {
		goto L12
	}
L12:
	;
	goto L10
L13:
	;
	v34 = int32(2)
	goto L5
L14:
	;
	F_addReplyBulkCString(m, l0, int32(_a346))
	mBase = m.M
	v56 = m.ExcPending
	if v56 != 0 {
		goto L1
	} else {
		goto L22
	}
L15:
	;
	F_addReplyBulkCString(m, l0, int32(_a356))
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
		goto L1
	} else {
		goto L16
	}
L16:
	;
	if l0 == int32(0) {
		goto L18
	} else {
		goto L19
	}
L17:
	;
	F_addReplyLongLong(m, l0, base.I64_extend_i32_s(v46))
	mBase = m.M
	v49 = m.ExcPending
	if v49 != 0 {
		goto L1
	} else {
		goto L21
	}
L18:
	;
	v45 = *(*int32)(unsafe.Add(mBase, uint32(l1)+2328))
	v46 = v45
	goto L17
L19:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(l1)+2340))
	if v43 != 0 {
		v46 = v43
		goto L17
	} else {
		goto L20
	}
L20:
	;
	goto L18
L21:
	;
	v52 = v34 + int32(1)
	goto L14
L22:
	;
	v57 = F_clusterNodeIp(m, l1, l0)
	mBase = m.M
	v58 = m.ExcPending
	if v58 != 0 {
		goto L1
	} else {
		goto L23
	}
L23:
	;
	F_addReplyBulkCString(m, l0, v57)
	mBase = m.M
	v60 = m.ExcPending
	if v60 != 0 {
		goto L1
	} else {
		goto L24
	}
L24:
	;
	F_addReplyBulkCString(m, l0, int32(_a347))
	mBase = m.M
	v63 = m.ExcPending
	if v63 != 0 {
		goto L1
	} else {
		goto L25
	}
L25:
	;
	v66 = *(*int32)(unsafe.Add(mBase, _consts[207]))
	switch v66 {
	case 0:
		goto L29
	case 1:
		goto L28
	case 2:
		v75 = int32(_a139)
		goto L26
	default:
		goto L27
	}
L26:
	;
	F_addReplyBulkCString(m, l0, v75)
	mBase = m.M
	v77 = m.ExcPending
	if v77 != 0 {
		goto L1
	} else {
		goto L36
	}
L27:
	;
	v75 = int32(_a288)
	goto L26
L28:
	;
	v69 = *(*int32)(unsafe.Add(mBase, uint32(l1)+2312))
	if v69 != 0 {
		goto L31
	} else {
		goto L32
	}
L29:
	;
	v67 = F_clusterNodeIp(m, l1, l0)
	mBase = m.M
	v68 = m.ExcPending
	if v68 != 0 {
		goto L1
	} else {
		goto L30
	}
L30:
	;
	v75 = v67
	goto L26
L31:
	;
	v72 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v69))))
	if v72 != 0 {
		goto L33
	} else {
		goto L34
	}
L32:
	;
	v75 = int32(_a357)
	goto L26
L33:
	;
	v73 = v69
	goto L35
L34:
	;
	v73 = int32(_a357)
	goto L35
L35:
	;
	v75 = v73
	goto L26
L36:
	;
	v79 = v52 + int32(2)
	v80 = *(*int32)(unsafe.Add(mBase, uint32(l1)+2312))
	v83 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v80+int32(-1)))))
	switch v83 & int32(7) {
	case 0:
		goto L43
	case 1:
		goto L42
	case 2:
		goto L41
	case 3:
		goto L40
	case 4:
		goto L39
	default:
		v132 = v79
		goto L37
	}
L37:
	;
	v135 = *(*int32)(unsafe.Add(mBase, uint32(l1)+88))
	if v135&int32(16) == int32(0) {
		goto L54
	} else {
		goto L55
	}
L38:
	;
	if v100 == int32(0) {
		v132 = v79
		goto L37
	} else {
		goto L44
	}
L39:
	;
	v99 = *(*int32)(unsafe.Add(mBase, uint32(v80+int32(-17))))
	v100 = v99
	goto L38
L40:
	;
	v96 = *(*int32)(unsafe.Add(mBase, uint32(v80+int32(-9))))
	v100 = v96
	goto L38
L41:
	;
	v93 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v80+int32(-5)))))
	v100 = v93
	goto L38
L42:
	;
	v90 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v80+int32(-3)))))
	v100 = v90
	goto L38
L43:
	;
	v100 = int32(base.Ui32(v83) >> (uint(int32(3)) % 32))
	goto L38
L44:
	;
	F_addReplyBulkCString(m, l0, int32(_a355))
	mBase = m.M
	v105 = m.ExcPending
	if v105 != 0 {
		goto L1
	} else {
		goto L45
	}
L45:
	;
	v107 = *(*int32)(unsafe.Add(mBase, uint32(l1)+2312))
	v110 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v107+int32(-1)))))
	switch v110 & int32(7) {
	case 0:
		goto L51
	case 1:
		goto L50
	case 2:
		goto L49
	case 3:
		goto L48
	case 4:
		goto L47
	default:
		v127 = int32(0)
		goto L46
	}
L46:
	;
	F_addReplyBulkCBuffer(m, l0, v107, v127)
	mBase = m.M
	v129 = m.ExcPending
	if v129 != 0 {
		goto L1
	} else {
		goto L52
	}
L47:
	;
	v126 = *(*int32)(unsafe.Add(mBase, uint32(v107+int32(-17))))
	v127 = v126
	goto L46
L48:
	;
	v123 = *(*int32)(unsafe.Add(mBase, uint32(v107+int32(-9))))
	v127 = v123
	goto L46
L49:
	;
	v120 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v107+int32(-5)))))
	v127 = v120
	goto L46
L50:
	;
	v117 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v107+int32(-3)))))
	v127 = v117
	goto L46
L51:
	;
	v127 = int32(base.Ui32(v110) >> (uint(int32(3)) % 32))
	goto L46
L52:
	;
	v132 = v52 + int32(3)
	goto L37
L53:
	;
	F_addReplyBulkCString(m, l0, int32(_a348))
	mBase = m.M
	v174 = m.ExcPending
	if v174 != 0 {
		goto L1
	} else {
		goto L67
	}
L54:
	;
	v170 = *(*int64)(unsafe.Add(mBase, uint32(l1)+2248))
	v171 = v170
	goto L53
L55:
	;
	if v135&int32(2) == int32(0) {
		goto L56
	} else {
		goto L57
	}
L56:
	;
	v169 = *(*int64)(unsafe.Add(mBase, _consts[40]))
	v171 = v169
	goto L53
L57:
	;
	v148 = *(*int32)(unsafe.Add(mBase, _consts[130]))
	if v148 == int32(0) {
		v162 = int64(0)
		goto L59
	} else {
		goto L60
	}
L58:
	;
	v171 = v167
	goto L53
L59:
	;
	v164 = int64(0)
	if v164 < v162 {
		goto L64
	} else {
		goto L65
	}
L60:
	;
	v152 = *(*int32)(unsafe.Add(mBase, _consts[188]))
	if v152 != 0 {
		v159 = v152
		goto L61
	} else {
		goto L62
	}
L61:
	;
	v160 = *(*int32)(unsafe.Add(mBase, uint32(v159)+104))
	v161 = *(*int64)(unsafe.Add(mBase, uint32(v160)+48))
	v162 = v161
	goto L59
L62:
	;
	v155 = *(*int32)(unsafe.Add(mBase, _consts[189]))
	if v155 == int32(0) {
		v162 = int64(0)
		goto L59
	} else {
		goto L63
	}
L63:
	;
	v159 = v155
	goto L61
L64:
	;
	v167 = v162
	goto L66
L65:
	;
	v167 = v164
	goto L66
L66:
	;
	goto L58
L67:
	;
	v177 = *(*int32)(unsafe.Add(mBase, uint32(l1)+88))
	if v177&int32(2) != 0 {
		goto L68
	} else {
		goto L69
	}
L68:
	;
	v180 = int32(_a268)
	goto L70
L69:
	;
	v180 = int32(_a349)
	goto L70
L70:
	;
	F_addReplyBulkCString(m, l0, v180)
	mBase = m.M
	v182 = m.ExcPending
	if v182 != 0 {
		goto L1
	} else {
		goto L71
	}
L71:
	;
	F_addReplyBulkCString(m, l0, int32(_a350))
	mBase = m.M
	v185 = m.ExcPending
	if v185 != 0 {
		goto L1
	} else {
		goto L72
	}
L72:
	;
	F_addReplyLongLong(m, l0, v171)
	mBase = m.M
	v187 = m.ExcPending
	if v187 != 0 {
		goto L1
	} else {
		goto L73
	}
L73:
	;
	F_addReplyBulkCString(m, l0, int32(_a351))
	mBase = m.M
	v190 = m.ExcPending
	if v190 != 0 {
		goto L1
	} else {
		goto L74
	}
L74:
	;
	v192 = *(*int32)(unsafe.Add(mBase, uint32(l1)+88))
	if v192&int32(8) != 0 {
		v206 = int32(_a192)
		goto L75
	} else {
		goto L76
	}
L75:
	;
	F_addReplyBulkCString(m, l0, v206)
	mBase = m.M
	v208 = m.ExcPending
	if v208 != 0 {
		goto L1
	} else {
		goto L83
	}
L76:
	;
	if v171 == int64(0) {
		goto L77
	} else {
		goto L78
	}
L77:
	;
	v199 = int32(_a352)
	goto L79
L78:
	;
	v199 = int32(_a353)
	goto L79
L79:
	;
	if int32(base.Ui32(v192&int32(2))>>(uint(int32(1))%32)) != 0 {
		goto L80
	} else {
		goto L81
	}
L80:
	;
	v205 = v199
	goto L82
L81:
	;
	v205 = int32(_a353)
	goto L82
L82:
	;
	v206 = v205
	goto L75
L83:
	;
	v210 = v132 + int32(3)
	v211 = *(*int32)(unsafe.Add(mBase, uint32(l1)+2320))
	v214 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v211+int32(-1)))))
	switch v214 & int32(7) {
	case 0:
		goto L90
	case 1:
		goto L89
	case 2:
		goto L88
	case 3:
		goto L87
	case 4:
		goto L86
	default:
		v264 = v210
		goto L84
	}
L84:
	;
	F_setDeferredMapLen(m, l0, v9, v264)
	mBase = m.M
	v267 = m.ExcPending
	if v267 != 0 {
		goto L1
	} else {
		goto L100
	}
L85:
	;
	if v231 == int32(0) {
		v264 = v210
		goto L84
	} else {
		goto L91
	}
L86:
	;
	v230 = *(*int32)(unsafe.Add(mBase, uint32(v211+int32(-17))))
	v231 = v230
	goto L85
L87:
	;
	v227 = *(*int32)(unsafe.Add(mBase, uint32(v211+int32(-9))))
	v231 = v227
	goto L85
L88:
	;
	v224 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v211+int32(-5)))))
	v231 = v224
	goto L85
L89:
	;
	v221 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v211+int32(-3)))))
	v231 = v221
	goto L85
L90:
	;
	v231 = int32(base.Ui32(v214) >> (uint(int32(3)) % 32))
	goto L85
L91:
	;
	F_addReplyBulkCString(m, l0, int32(_a354))
	mBase = m.M
	v236 = m.ExcPending
	if v236 != 0 {
		goto L1
	} else {
		goto L92
	}
L92:
	;
	v238 = *(*int32)(unsafe.Add(mBase, uint32(l1)+2320))
	v241 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v238+int32(-1)))))
	switch v241 & int32(7) {
	case 0:
		goto L98
	case 1:
		goto L97
	case 2:
		goto L96
	case 3:
		goto L95
	case 4:
		goto L94
	default:
		v258 = int32(0)
		goto L93
	}
L93:
	;
	F_addReplyBulkCBuffer(m, l0, v238, v258)
	mBase = m.M
	v260 = m.ExcPending
	if v260 != 0 {
		goto L1
	} else {
		goto L99
	}
L94:
	;
	v257 = *(*int32)(unsafe.Add(mBase, uint32(v238+int32(-17))))
	v258 = v257
	goto L93
L95:
	;
	v254 = *(*int32)(unsafe.Add(mBase, uint32(v238+int32(-9))))
	v258 = v254
	goto L93
L96:
	;
	v251 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v238+int32(-5)))))
	v258 = v251
	goto L93
L97:
	;
	v248 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v238+int32(-3)))))
	v258 = v248
	goto L93
L98:
	;
	v258 = int32(base.Ui32(v241) >> (uint(int32(3)) % 32))
	goto L93
L99:
	;
	v264 = v132 + int32(4)
	goto L84
L100:
	;
	return
}
func F_clearNodeFailureIfNeeded(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int64
	_ = v10
	var v11 int32
	_ = v11
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v36 int32
	_ = v36
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v57 int32
	_ = v57
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v70 int32
	_ = v70
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v80 int64
	_ = v80
	var v83 int64
	_ = v83
	var v88 int32
	_ = v88
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v113 int32
	_ = v113
	var v119 int32
	_ = v119
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v136 int32
	_ = v136
	v6 = m.G0
	v8 = v6 - int32(32)
	m.G0 = v8
	v10 = F_mstime(m)
	mBase = m.M
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+88))
	if v11&int32(8) == int32(0) {
		F__serverAssert(m, int32(_a266), int32(_a253), int32(2613))
		mBase = m.M
		v136 = m.ExcPending
		if v136 != 0 {
			return
		} else {
			F_abort(m)
			mBase = m.M
			base.Wasm_trap_unreachable()
			for {
			}
		}
	} else {
		if v11&int32(1) == int32(0) {
			v22 = *(*int32)(unsafe.Add(mBase, _consts[6]))
			if int32(2) < v22 {
				v45 = v11
				*(*int32)(unsafe.Add(mBase, uint32(l0)+88)) = v45 & int32(-9)
				v50 = int32(0)
				v51 = *(*int32)(unsafe.Add(mBase, _consts[154]))
				v52 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v51)+88)))
				if v52&int32(2) == v50 {
				} else {
					v57 = *(*int32)(unsafe.Add(mBase, uint32(v51)+2172))
					if v57 != l0 {
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(l0)+88)) = v45 & int32(-8201)
					}
				}
				F_clearCachedClusterSlotsResponse(m)
				mBase = m.M
				v63 = m.ExcPending
				if v63 != 0 {
					return
				} else {
					v65 = *(*int32)(unsafe.Add(mBase, _consts[136]))
					v66 = *(*int32)(unsafe.Add(mBase, uint32(v65)+uint32(_consts[150])))
					*(*int32)(unsafe.Add(mBase, uint32(v65)+uint32(_consts[150]))) = v66 | int32(6)
					v70 = *(*int32)(unsafe.Add(mBase, uint32(l0)+88))
					if v70&int32(1) == int32(0) {
						m.G0 = v8 + int32(32)
						return
					} else {
						v75 = v70
						v77 = *(*int32)(unsafe.Add(mBase, uint32(l0)+2160))
						if v77 == int32(0) {
							m.G0 = v8 + int32(32)
							return
						} else {
							v80 = *(*int64)(unsafe.Add(mBase, uint32(l0)+2216))
							v83 = *(*int64)(unsafe.Add(mBase, _consts[164]))
							if v10-v80 <= v83<<(uint(int64(1))%64) {
								m.G0 = v8 + int32(32)
								return
							} else {
								v88 = *(*int32)(unsafe.Add(mBase, _consts[6]))
								if int32(2) < v88 {
									v102 = v75
									*(*int32)(unsafe.Add(mBase, uint32(l0)+88)) = v102 & int32(-9)
									v106 = int32(0)
									v107 = *(*int32)(unsafe.Add(mBase, _consts[154]))
									v108 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v107)+88)))
									if v108&int32(2) == v106 {
									} else {
										v113 = *(*int32)(unsafe.Add(mBase, uint32(v107)+2172))
										if v113 != l0 {
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(l0)+88)) = v102 & int32(-8201)
										}
									}
									F_clearCachedClusterSlotsResponse(m)
									mBase = m.M
									v119 = m.ExcPending
									if v119 != 0 {
										return
									} else {
										v121 = *(*int32)(unsafe.Add(mBase, _consts[136]))
										v122 = *(*int32)(unsafe.Add(mBase, uint32(v121)+uint32(_consts[150])))
										*(*int32)(unsafe.Add(mBase, uint32(v121)+uint32(_consts[150]))) = v122 | int32(6)
										m.G0 = v8 + int32(32)
										return
									}
								} else {
									v91 = F_humanNodename(m, l0)
									mBase = m.M
									v92 = m.ExcPending
									if v92 != 0 {
										return
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = v91
										*(*int32)(unsafe.Add(mBase, uint32(v8))) = l0 + int32(8)
										F__serverLog(m, int32(2), int32(_a267), v8)
										mBase = m.M
										v100 = m.ExcPending
										if v100 != 0 {
											return
										} else {
											v101 = *(*int32)(unsafe.Add(mBase, uint32(l0)+88))
											v102 = v101
											*(*int32)(unsafe.Add(mBase, uint32(l0)+88)) = v102 & int32(-9)
											v106 = int32(0)
											v107 = *(*int32)(unsafe.Add(mBase, _consts[154]))
											v108 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v107)+88)))
											if v108&int32(2) == v106 {
											} else {
												v113 = *(*int32)(unsafe.Add(mBase, uint32(v107)+2172))
												if v113 != l0 {
												} else {
													*(*int32)(unsafe.Add(mBase, uint32(l0)+88)) = v102 & int32(-8201)
												}
											}
											F_clearCachedClusterSlotsResponse(m)
											mBase = m.M
											v119 = m.ExcPending
											if v119 != 0 {
												return
											} else {
												v121 = *(*int32)(unsafe.Add(mBase, _consts[136]))
												v122 = *(*int32)(unsafe.Add(mBase, uint32(v121)+uint32(_consts[150])))
												*(*int32)(unsafe.Add(mBase, uint32(v121)+uint32(_consts[150]))) = v122 | int32(6)
												m.G0 = v8 + int32(32)
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
				v25 = F_humanNodename(m, l0)
				mBase = m.M
				v26 = m.ExcPending
				if v26 != 0 {
					return
				} else {
					v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)+88))
					*(*int32)(unsafe.Add(mBase, uint32(v8)+20)) = v25
					*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = l0 + int32(8)
					if v27&int32(2) != 0 {
						v36 = int32(_a268)
					} else {
						v36 = int32(_a269)
					}
					*(*int32)(unsafe.Add(mBase, uint32(v8)+24)) = v36
					F__serverLog(m, int32(2), int32(_a270), v8+int32(16))
					mBase = m.M
					v43 = m.ExcPending
					if v43 != 0 {
						return
					} else {
						v44 = *(*int32)(unsafe.Add(mBase, uint32(l0)+88))
						v45 = v44
						*(*int32)(unsafe.Add(mBase, uint32(l0)+88)) = v45 & int32(-9)
						v50 = int32(0)
						v51 = *(*int32)(unsafe.Add(mBase, _consts[154]))
						v52 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v51)+88)))
						if v52&int32(2) == v50 {
						} else {
							v57 = *(*int32)(unsafe.Add(mBase, uint32(v51)+2172))
							if v57 != l0 {
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(l0)+88)) = v45 & int32(-8201)
							}
						}
						F_clearCachedClusterSlotsResponse(m)
						mBase = m.M
						v63 = m.ExcPending
						if v63 != 0 {
							return
						} else {
							v65 = *(*int32)(unsafe.Add(mBase, _consts[136]))
							v66 = *(*int32)(unsafe.Add(mBase, uint32(v65)+uint32(_consts[150])))
							*(*int32)(unsafe.Add(mBase, uint32(v65)+uint32(_consts[150]))) = v66 | int32(6)
							v70 = *(*int32)(unsafe.Add(mBase, uint32(l0)+88))
							if v70&int32(1) == int32(0) {
								m.G0 = v8 + int32(32)
								return
							} else {
								v75 = v70
								v77 = *(*int32)(unsafe.Add(mBase, uint32(l0)+2160))
								if v77 == int32(0) {
									m.G0 = v8 + int32(32)
									return
								} else {
									v80 = *(*int64)(unsafe.Add(mBase, uint32(l0)+2216))
									v83 = *(*int64)(unsafe.Add(mBase, _consts[164]))
									if v10-v80 <= v83<<(uint(int64(1))%64) {
										m.G0 = v8 + int32(32)
										return
									} else {
										v88 = *(*int32)(unsafe.Add(mBase, _consts[6]))
										if int32(2) < v88 {
											v102 = v75
											*(*int32)(unsafe.Add(mBase, uint32(l0)+88)) = v102 & int32(-9)
											v106 = int32(0)
											v107 = *(*int32)(unsafe.Add(mBase, _consts[154]))
											v108 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v107)+88)))
											if v108&int32(2) == v106 {
											} else {
												v113 = *(*int32)(unsafe.Add(mBase, uint32(v107)+2172))
												if v113 != l0 {
												} else {
													*(*int32)(unsafe.Add(mBase, uint32(l0)+88)) = v102 & int32(-8201)
												}
											}
											F_clearCachedClusterSlotsResponse(m)
											mBase = m.M
											v119 = m.ExcPending
											if v119 != 0 {
												return
											} else {
												v121 = *(*int32)(unsafe.Add(mBase, _consts[136]))
												v122 = *(*int32)(unsafe.Add(mBase, uint32(v121)+uint32(_consts[150])))
												*(*int32)(unsafe.Add(mBase, uint32(v121)+uint32(_consts[150]))) = v122 | int32(6)
												m.G0 = v8 + int32(32)
												return
											}
										} else {
											v91 = F_humanNodename(m, l0)
											mBase = m.M
											v92 = m.ExcPending
											if v92 != 0 {
												return
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = v91
												*(*int32)(unsafe.Add(mBase, uint32(v8))) = l0 + int32(8)
												F__serverLog(m, int32(2), int32(_a267), v8)
												mBase = m.M
												v100 = m.ExcPending
												if v100 != 0 {
													return
												} else {
													v101 = *(*int32)(unsafe.Add(mBase, uint32(l0)+88))
													v102 = v101
													*(*int32)(unsafe.Add(mBase, uint32(l0)+88)) = v102 & int32(-9)
													v106 = int32(0)
													v107 = *(*int32)(unsafe.Add(mBase, _consts[154]))
													v108 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v107)+88)))
													if v108&int32(2) == v106 {
													} else {
														v113 = *(*int32)(unsafe.Add(mBase, uint32(v107)+2172))
														if v113 != l0 {
														} else {
															*(*int32)(unsafe.Add(mBase, uint32(l0)+88)) = v102 & int32(-8201)
														}
													}
													F_clearCachedClusterSlotsResponse(m)
													mBase = m.M
													v119 = m.ExcPending
													if v119 != 0 {
														return
													} else {
														v121 = *(*int32)(unsafe.Add(mBase, _consts[136]))
														v122 = *(*int32)(unsafe.Add(mBase, uint32(v121)+uint32(_consts[150])))
														*(*int32)(unsafe.Add(mBase, uint32(v121)+uint32(_consts[150]))) = v122 | int32(6)
														m.G0 = v8 + int32(32)
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
		} else {
			v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+2160))
			if v20 != 0 {
				v75 = v11
				v77 = *(*int32)(unsafe.Add(mBase, uint32(l0)+2160))
				if v77 == int32(0) {
					m.G0 = v8 + int32(32)
					return
				} else {
					v80 = *(*int64)(unsafe.Add(mBase, uint32(l0)+2216))
					v83 = *(*int64)(unsafe.Add(mBase, _consts[164]))
					if v10-v80 <= v83<<(uint(int64(1))%64) {
						m.G0 = v8 + int32(32)
						return
					} else {
						v88 = *(*int32)(unsafe.Add(mBase, _consts[6]))
						if int32(2) < v88 {
							v102 = v75
							*(*int32)(unsafe.Add(mBase, uint32(l0)+88)) = v102 & int32(-9)
							v106 = int32(0)
							v107 = *(*int32)(unsafe.Add(mBase, _consts[154]))
							v108 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v107)+88)))
							if v108&int32(2) == v106 {
							} else {
								v113 = *(*int32)(unsafe.Add(mBase, uint32(v107)+2172))
								if v113 != l0 {
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(l0)+88)) = v102 & int32(-8201)
								}
							}
							F_clearCachedClusterSlotsResponse(m)
							mBase = m.M
							v119 = m.ExcPending
							if v119 != 0 {
								return
							} else {
								v121 = *(*int32)(unsafe.Add(mBase, _consts[136]))
								v122 = *(*int32)(unsafe.Add(mBase, uint32(v121)+uint32(_consts[150])))
								*(*int32)(unsafe.Add(mBase, uint32(v121)+uint32(_consts[150]))) = v122 | int32(6)
								m.G0 = v8 + int32(32)
								return
							}
						} else {
							v91 = F_humanNodename(m, l0)
							mBase = m.M
							v92 = m.ExcPending
							if v92 != 0 {
								return
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = v91
								*(*int32)(unsafe.Add(mBase, uint32(v8))) = l0 + int32(8)
								F__serverLog(m, int32(2), int32(_a267), v8)
								mBase = m.M
								v100 = m.ExcPending
								if v100 != 0 {
									return
								} else {
									v101 = *(*int32)(unsafe.Add(mBase, uint32(l0)+88))
									v102 = v101
									*(*int32)(unsafe.Add(mBase, uint32(l0)+88)) = v102 & int32(-9)
									v106 = int32(0)
									v107 = *(*int32)(unsafe.Add(mBase, _consts[154]))
									v108 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v107)+88)))
									if v108&int32(2) == v106 {
									} else {
										v113 = *(*int32)(unsafe.Add(mBase, uint32(v107)+2172))
										if v113 != l0 {
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(l0)+88)) = v102 & int32(-8201)
										}
									}
									F_clearCachedClusterSlotsResponse(m)
									mBase = m.M
									v119 = m.ExcPending
									if v119 != 0 {
										return
									} else {
										v121 = *(*int32)(unsafe.Add(mBase, _consts[136]))
										v122 = *(*int32)(unsafe.Add(mBase, uint32(v121)+uint32(_consts[150])))
										*(*int32)(unsafe.Add(mBase, uint32(v121)+uint32(_consts[150]))) = v122 | int32(6)
										m.G0 = v8 + int32(32)
										return
									}
								}
							}
						}
					}
				}
			} else {
				v22 = *(*int32)(unsafe.Add(mBase, _consts[6]))
				if int32(2) < v22 {
					v45 = v11
					*(*int32)(unsafe.Add(mBase, uint32(l0)+88)) = v45 & int32(-9)
					v50 = int32(0)
					v51 = *(*int32)(unsafe.Add(mBase, _consts[154]))
					v52 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v51)+88)))
					if v52&int32(2) == v50 {
					} else {
						v57 = *(*int32)(unsafe.Add(mBase, uint32(v51)+2172))
						if v57 != l0 {
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(l0)+88)) = v45 & int32(-8201)
						}
					}
					F_clearCachedClusterSlotsResponse(m)
					mBase = m.M
					v63 = m.ExcPending
					if v63 != 0 {
						return
					} else {
						v65 = *(*int32)(unsafe.Add(mBase, _consts[136]))
						v66 = *(*int32)(unsafe.Add(mBase, uint32(v65)+uint32(_consts[150])))
						*(*int32)(unsafe.Add(mBase, uint32(v65)+uint32(_consts[150]))) = v66 | int32(6)
						v70 = *(*int32)(unsafe.Add(mBase, uint32(l0)+88))
						if v70&int32(1) == int32(0) {
							m.G0 = v8 + int32(32)
							return
						} else {
							v75 = v70
							v77 = *(*int32)(unsafe.Add(mBase, uint32(l0)+2160))
							if v77 == int32(0) {
								m.G0 = v8 + int32(32)
								return
							} else {
								v80 = *(*int64)(unsafe.Add(mBase, uint32(l0)+2216))
								v83 = *(*int64)(unsafe.Add(mBase, _consts[164]))
								if v10-v80 <= v83<<(uint(int64(1))%64) {
									m.G0 = v8 + int32(32)
									return
								} else {
									v88 = *(*int32)(unsafe.Add(mBase, _consts[6]))
									if int32(2) < v88 {
										v102 = v75
										*(*int32)(unsafe.Add(mBase, uint32(l0)+88)) = v102 & int32(-9)
										v106 = int32(0)
										v107 = *(*int32)(unsafe.Add(mBase, _consts[154]))
										v108 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v107)+88)))
										if v108&int32(2) == v106 {
										} else {
											v113 = *(*int32)(unsafe.Add(mBase, uint32(v107)+2172))
											if v113 != l0 {
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(l0)+88)) = v102 & int32(-8201)
											}
										}
										F_clearCachedClusterSlotsResponse(m)
										mBase = m.M
										v119 = m.ExcPending
										if v119 != 0 {
											return
										} else {
											v121 = *(*int32)(unsafe.Add(mBase, _consts[136]))
											v122 = *(*int32)(unsafe.Add(mBase, uint32(v121)+uint32(_consts[150])))
											*(*int32)(unsafe.Add(mBase, uint32(v121)+uint32(_consts[150]))) = v122 | int32(6)
											m.G0 = v8 + int32(32)
											return
										}
									} else {
										v91 = F_humanNodename(m, l0)
										mBase = m.M
										v92 = m.ExcPending
										if v92 != 0 {
											return
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = v91
											*(*int32)(unsafe.Add(mBase, uint32(v8))) = l0 + int32(8)
											F__serverLog(m, int32(2), int32(_a267), v8)
											mBase = m.M
											v100 = m.ExcPending
											if v100 != 0 {
												return
											} else {
												v101 = *(*int32)(unsafe.Add(mBase, uint32(l0)+88))
												v102 = v101
												*(*int32)(unsafe.Add(mBase, uint32(l0)+88)) = v102 & int32(-9)
												v106 = int32(0)
												v107 = *(*int32)(unsafe.Add(mBase, _consts[154]))
												v108 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v107)+88)))
												if v108&int32(2) == v106 {
												} else {
													v113 = *(*int32)(unsafe.Add(mBase, uint32(v107)+2172))
													if v113 != l0 {
													} else {
														*(*int32)(unsafe.Add(mBase, uint32(l0)+88)) = v102 & int32(-8201)
													}
												}
												F_clearCachedClusterSlotsResponse(m)
												mBase = m.M
												v119 = m.ExcPending
												if v119 != 0 {
													return
												} else {
													v121 = *(*int32)(unsafe.Add(mBase, _consts[136]))
													v122 = *(*int32)(unsafe.Add(mBase, uint32(v121)+uint32(_consts[150])))
													*(*int32)(unsafe.Add(mBase, uint32(v121)+uint32(_consts[150]))) = v122 | int32(6)
													m.G0 = v8 + int32(32)
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
					v25 = F_humanNodename(m, l0)
					mBase = m.M
					v26 = m.ExcPending
					if v26 != 0 {
						return
					} else {
						v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)+88))
						*(*int32)(unsafe.Add(mBase, uint32(v8)+20)) = v25
						*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = l0 + int32(8)
						if v27&int32(2) != 0 {
							v36 = int32(_a268)
						} else {
							v36 = int32(_a269)
						}
						*(*int32)(unsafe.Add(mBase, uint32(v8)+24)) = v36
						F__serverLog(m, int32(2), int32(_a270), v8+int32(16))
						mBase = m.M
						v43 = m.ExcPending
						if v43 != 0 {
							return
						} else {
							v44 = *(*int32)(unsafe.Add(mBase, uint32(l0)+88))
							v45 = v44
							*(*int32)(unsafe.Add(mBase, uint32(l0)+88)) = v45 & int32(-9)
							v50 = int32(0)
							v51 = *(*int32)(unsafe.Add(mBase, _consts[154]))
							v52 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v51)+88)))
							if v52&int32(2) == v50 {
							} else {
								v57 = *(*int32)(unsafe.Add(mBase, uint32(v51)+2172))
								if v57 != l0 {
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(l0)+88)) = v45 & int32(-8201)
								}
							}
							F_clearCachedClusterSlotsResponse(m)
							mBase = m.M
							v63 = m.ExcPending
							if v63 != 0 {
								return
							} else {
								v65 = *(*int32)(unsafe.Add(mBase, _consts[136]))
								v66 = *(*int32)(unsafe.Add(mBase, uint32(v65)+uint32(_consts[150])))
								*(*int32)(unsafe.Add(mBase, uint32(v65)+uint32(_consts[150]))) = v66 | int32(6)
								v70 = *(*int32)(unsafe.Add(mBase, uint32(l0)+88))
								if v70&int32(1) == int32(0) {
									m.G0 = v8 + int32(32)
									return
								} else {
									v75 = v70
									v77 = *(*int32)(unsafe.Add(mBase, uint32(l0)+2160))
									if v77 == int32(0) {
										m.G0 = v8 + int32(32)
										return
									} else {
										v80 = *(*int64)(unsafe.Add(mBase, uint32(l0)+2216))
										v83 = *(*int64)(unsafe.Add(mBase, _consts[164]))
										if v10-v80 <= v83<<(uint(int64(1))%64) {
											m.G0 = v8 + int32(32)
											return
										} else {
											v88 = *(*int32)(unsafe.Add(mBase, _consts[6]))
											if int32(2) < v88 {
												v102 = v75
												*(*int32)(unsafe.Add(mBase, uint32(l0)+88)) = v102 & int32(-9)
												v106 = int32(0)
												v107 = *(*int32)(unsafe.Add(mBase, _consts[154]))
												v108 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v107)+88)))
												if v108&int32(2) == v106 {
												} else {
													v113 = *(*int32)(unsafe.Add(mBase, uint32(v107)+2172))
													if v113 != l0 {
													} else {
														*(*int32)(unsafe.Add(mBase, uint32(l0)+88)) = v102 & int32(-8201)
													}
												}
												F_clearCachedClusterSlotsResponse(m)
												mBase = m.M
												v119 = m.ExcPending
												if v119 != 0 {
													return
												} else {
													v121 = *(*int32)(unsafe.Add(mBase, _consts[136]))
													v122 = *(*int32)(unsafe.Add(mBase, uint32(v121)+uint32(_consts[150])))
													*(*int32)(unsafe.Add(mBase, uint32(v121)+uint32(_consts[150]))) = v122 | int32(6)
													m.G0 = v8 + int32(32)
													return
												}
											} else {
												v91 = F_humanNodename(m, l0)
												mBase = m.M
												v92 = m.ExcPending
												if v92 != 0 {
													return
												} else {
													*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = v91
													*(*int32)(unsafe.Add(mBase, uint32(v8))) = l0 + int32(8)
													F__serverLog(m, int32(2), int32(_a267), v8)
													mBase = m.M
													v100 = m.ExcPending
													if v100 != 0 {
														return
													} else {
														v101 = *(*int32)(unsafe.Add(mBase, uint32(l0)+88))
														v102 = v101
														*(*int32)(unsafe.Add(mBase, uint32(l0)+88)) = v102 & int32(-9)
														v106 = int32(0)
														v107 = *(*int32)(unsafe.Add(mBase, _consts[154]))
														v108 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v107)+88)))
														if v108&int32(2) == v106 {
														} else {
															v113 = *(*int32)(unsafe.Add(mBase, uint32(v107)+2172))
															if v113 != l0 {
															} else {
																*(*int32)(unsafe.Add(mBase, uint32(l0)+88)) = v102 & int32(-8201)
															}
														}
														F_clearCachedClusterSlotsResponse(m)
														mBase = m.M
														v119 = m.ExcPending
														if v119 != 0 {
															return
														} else {
															v121 = *(*int32)(unsafe.Add(mBase, _consts[136]))
															v122 = *(*int32)(unsafe.Add(mBase, uint32(v121)+uint32(_consts[150])))
															*(*int32)(unsafe.Add(mBase, uint32(v121)+uint32(_consts[150]))) = v122 | int32(6)
															m.G0 = v8 + int32(32)
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
func F_getNodeBySlot(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v9 int32
	_ = v9
	v3 = *(*int32)(unsafe.Add(mBase, _consts[136]))
	v9 = *(*int32)(unsafe.Add(mBase, uint32(v3+l0<<(uint(int32(2))%32)+int32(52))))
	return v9
}
func F_getNodeDefaultReplicationPort(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	v5 = *(*int32)(unsafe.Add(mBase, _consts[149]))
	if v5 != 0 {
		v6 = int32(2328)
	} else {
		v6 = int32(2324)
	}
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0+v6)))
	return v8
}
func F_markNodeAsFailingIfNeeded(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int64
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v32 int32
	_ = v32
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v57 int64
	_ = v57
	var v59 int32
	_ = v59
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v70 int32
	_ = v70
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v98 int64
	_ = v98
	var v104 int64
	_ = v104
	var v110 int64
	_ = v110
	var v116 int64
	_ = v116
	var v118 int64
	_ = v118
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v124 int32
	_ = v124
	var v128 int32
	_ = v128
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v135 int32
	_ = v135
	var v147 int32
	_ = v147
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v11 = *(*int32)(unsafe.Add(mBase, _consts[136]))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(v11)+28))
	v14 = base.I32_div_s(v12, int32(2))
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+88))
	if v15&int32(12) != int32(4) {
		m.G0 = v8 + int32(16)
		return
	} else {
		F_clusterNodeCleanupFailureReports(m, l0)
		mBase = m.M
		v21 = m.ExcPending
		if v21 != 0 {
			return
		} else {
			v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+2352))
			v23 = *(*int64)(unsafe.Add(mBase, uint32(v22)+8))
			v24 = base.I32_wrap_i64(v23)
			v25 = int32(0)
			v26 = *(*int32)(unsafe.Add(mBase, _consts[154]))
			v27 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v26)+88)))
			if v27&int32(1) == v25 {
				v37 = v24
			} else {
				v32 = *(*int32)(unsafe.Add(mBase, uint32(v26)+2160))
				if v32 == int32(0) {
					v37 = v24
				} else {
					v37 = v24 + int32(1)
				}
			}
			if v37 <= v14 {
				m.G0 = v8 + int32(16)
				return
			} else {
				v40 = *(*int32)(unsafe.Add(mBase, _consts[6]))
				if int32(2) < v40 {
					v53 = *(*int32)(unsafe.Add(mBase, uint32(l0)+88))
					*(*int32)(unsafe.Add(mBase, uint32(l0)+88)) = v53 | int32(8)
					v57 = F_mstime(m)
					mBase = m.M
					*(*int64)(unsafe.Add(mBase, uint32(l0)+2216)) = v57
					v59 = *(*int32)(unsafe.Add(mBase, uint32(l0)+88))
					*(*int32)(unsafe.Add(mBase, uint32(l0)+88)) = v59 & int32(-5)
					v63 = int32(0)
					v64 = *(*int32)(unsafe.Add(mBase, _consts[154]))
					v65 = *(*int32)(unsafe.Add(mBase, uint32(v64)+88))
					if v65&int32(2) == v63 {
					} else {
						v70 = *(*int32)(unsafe.Add(mBase, uint32(v64)+2172))
						if v70 != l0 {
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v64)+88)) = v65 | int32(8192)
							v76 = *(*int32)(unsafe.Add(mBase, _consts[136]))
							v77 = *(*int32)(unsafe.Add(mBase, uint32(v76)+uint32(_consts[150])))
							*(*int32)(unsafe.Add(mBase, uint32(v76)+uint32(_consts[150]))) = v77 | int32(1)
						}
					}
					F_clearCachedClusterSlotsResponse(m)
					mBase = m.M
					v83 = m.ExcPending
					if v83 != 0 {
						return
					} else {
						v85 = *(*int32)(unsafe.Add(mBase, _consts[136]))
						v86 = *(*int32)(unsafe.Add(mBase, uint32(v85)+uint32(_consts[150])))
						*(*int32)(unsafe.Add(mBase, uint32(v85)+uint32(_consts[150]))) = v86 | int32(6)
						v92 = F_createClusterMsgSendBlock(m, int32(3), int32(2296))
						mBase = m.M
						v93 = m.ExcPending
						if v93 != 0 {
							return
						} else {
							v98 = *(*int64)(unsafe.Add(mBase, uint32(l0+int32(40))))
							*(*int64)(unsafe.Add(mBase, uint32(v92+int32(2296)))) = v98
							v104 = *(*int64)(unsafe.Add(mBase, uint32(l0+int32(32))))
							*(*int64)(unsafe.Add(mBase, uint32(v92+int32(2288)))) = v104
							v110 = *(*int64)(unsafe.Add(mBase, uint32(l0+int32(24))))
							*(*int64)(unsafe.Add(mBase, uint32(v92+int32(2280)))) = v110
							v116 = *(*int64)(unsafe.Add(mBase, uint32(l0+int32(16))))
							*(*int64)(unsafe.Add(mBase, uint32(v92+int32(2272)))) = v116
							v118 = *(*int64)(unsafe.Add(mBase, uint32(l0)+8))
							*(*int64)(unsafe.Add(mBase, uint32(v92)+2264)) = v118
							F_clusterBroadcastMessage(m, v92)
							mBase = m.M
							v121 = m.ExcPending
							if v121 != 0 {
								return
							} else {
								v122 = *(*int32)(unsafe.Add(mBase, uint32(v92)+4))
								v124 = v122 + int32(-1)
								*(*int32)(unsafe.Add(mBase, uint32(v92)+4)) = v124
								if v122 <= int32(0) {
									F__serverAssert(m, int32(_a264), int32(_a253), int32(1756))
									mBase = m.M
									v147 = m.ExcPending
									if v147 != 0 {
										return
									} else {
										F_abort(m)
										mBase = m.M
										base.Wasm_trap_unreachable()
										for {
										}
									}
								} else {
									if v124 != 0 {
										m.G0 = v8 + int32(16)
										return
									} else {
										v128 = int32(_a44)
										v130 = *(*int32)(unsafe.Add(mBase, _consts[163]))
										v131 = *(*int32)(unsafe.Add(mBase, uint32(v92)))
										*(*int32)(unsafe.Add(mBase, _consts[163])) = v130 - v131
										F_valkey_free(m, v92)
										mBase = m.M
										v135 = m.ExcPending
										if v135 != 0 {
											return
										} else {
											m.G0 = v8 + int32(16)
											return
										}
									}
								}
							}
						}
					}
				} else {
					v43 = F_humanNodename(m, l0)
					mBase = m.M
					v44 = m.ExcPending
					if v44 != 0 {
						return
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = v43
						*(*int32)(unsafe.Add(mBase, uint32(v8))) = l0 + int32(8)
						F__serverLog(m, int32(2), int32(_a265), v8)
						mBase = m.M
						v52 = m.ExcPending
						if v52 != 0 {
							return
						} else {
							v53 = *(*int32)(unsafe.Add(mBase, uint32(l0)+88))
							*(*int32)(unsafe.Add(mBase, uint32(l0)+88)) = v53 | int32(8)
							v57 = F_mstime(m)
							mBase = m.M
							*(*int64)(unsafe.Add(mBase, uint32(l0)+2216)) = v57
							v59 = *(*int32)(unsafe.Add(mBase, uint32(l0)+88))
							*(*int32)(unsafe.Add(mBase, uint32(l0)+88)) = v59 & int32(-5)
							v63 = int32(0)
							v64 = *(*int32)(unsafe.Add(mBase, _consts[154]))
							v65 = *(*int32)(unsafe.Add(mBase, uint32(v64)+88))
							if v65&int32(2) == v63 {
							} else {
								v70 = *(*int32)(unsafe.Add(mBase, uint32(v64)+2172))
								if v70 != l0 {
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v64)+88)) = v65 | int32(8192)
									v76 = *(*int32)(unsafe.Add(mBase, _consts[136]))
									v77 = *(*int32)(unsafe.Add(mBase, uint32(v76)+uint32(_consts[150])))
									*(*int32)(unsafe.Add(mBase, uint32(v76)+uint32(_consts[150]))) = v77 | int32(1)
								}
							}
							F_clearCachedClusterSlotsResponse(m)
							mBase = m.M
							v83 = m.ExcPending
							if v83 != 0 {
								return
							} else {
								v85 = *(*int32)(unsafe.Add(mBase, _consts[136]))
								v86 = *(*int32)(unsafe.Add(mBase, uint32(v85)+uint32(_consts[150])))
								*(*int32)(unsafe.Add(mBase, uint32(v85)+uint32(_consts[150]))) = v86 | int32(6)
								v92 = F_createClusterMsgSendBlock(m, int32(3), int32(2296))
								mBase = m.M
								v93 = m.ExcPending
								if v93 != 0 {
									return
								} else {
									v98 = *(*int64)(unsafe.Add(mBase, uint32(l0+int32(40))))
									*(*int64)(unsafe.Add(mBase, uint32(v92+int32(2296)))) = v98
									v104 = *(*int64)(unsafe.Add(mBase, uint32(l0+int32(32))))
									*(*int64)(unsafe.Add(mBase, uint32(v92+int32(2288)))) = v104
									v110 = *(*int64)(unsafe.Add(mBase, uint32(l0+int32(24))))
									*(*int64)(unsafe.Add(mBase, uint32(v92+int32(2280)))) = v110
									v116 = *(*int64)(unsafe.Add(mBase, uint32(l0+int32(16))))
									*(*int64)(unsafe.Add(mBase, uint32(v92+int32(2272)))) = v116
									v118 = *(*int64)(unsafe.Add(mBase, uint32(l0)+8))
									*(*int64)(unsafe.Add(mBase, uint32(v92)+2264)) = v118
									F_clusterBroadcastMessage(m, v92)
									mBase = m.M
									v121 = m.ExcPending
									if v121 != 0 {
										return
									} else {
										v122 = *(*int32)(unsafe.Add(mBase, uint32(v92)+4))
										v124 = v122 + int32(-1)
										*(*int32)(unsafe.Add(mBase, uint32(v92)+4)) = v124
										if v122 <= int32(0) {
											F__serverAssert(m, int32(_a264), int32(_a253), int32(1756))
											mBase = m.M
											v147 = m.ExcPending
											if v147 != 0 {
												return
											} else {
												F_abort(m)
												mBase = m.M
												base.Wasm_trap_unreachable()
												for {
												}
											}
										} else {
											if v124 != 0 {
												m.G0 = v8 + int32(16)
												return
											} else {
												v128 = int32(_a44)
												v130 = *(*int32)(unsafe.Add(mBase, _consts[163]))
												v131 = *(*int32)(unsafe.Add(mBase, uint32(v92)))
												*(*int32)(unsafe.Add(mBase, _consts[163])) = v130 - v131
												F_valkey_free(m, v92)
												mBase = m.M
												v135 = m.ExcPending
												if v135 != 0 {
													return
												} else {
													m.G0 = v8 + int32(16)
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
func F_nodeUpdateAddressIfNeeded(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v16 int64
	_ = v16
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
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
	var v64 int64
	_ = v64
	var v72 int64
	_ = v72
	var v80 int64
	_ = v80
	var v88 int64
	_ = v88
	var v90 int64
	_ = v90
	var v94 int64
	_ = v94
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v110 int32
	_ = v110
	var v113 int32
	_ = v113
	var v118 int32
	_ = v118
	var v121 int32
	_ = v121
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v132 int32
	_ = v132
	var v135 int32
	_ = v135
	var v137 int32
	_ = v137
	var v139 int32
	_ = v139
	var v142 int32
	_ = v142
	var v144 int32
	_ = v144
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v162 int32
	_ = v162
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v178 int64
	_ = v178
	var v182 int64
	_ = v182
	var v188 int64
	_ = v188
	var v194 int64
	_ = v194
	var v200 int64
	_ = v200
	var v206 int64
	_ = v206
	var v211 int32
	_ = v211
	var v215 int32
	_ = v215
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v223 int32
	_ = v223
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v234 int32
	_ = v234
	var v246 int32
	_ = v246
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v251 int32
	_ = v251
	var v252 int32
	_ = v252
	var v257 int32
	_ = v257
	var v262 int32
	_ = v262
	var v263 int32
	_ = v263
	var v265 int32
	_ = v265
	var v266 int32
	_ = v266
	var v269 int32
	_ = v269
	var v273 int32
	_ = v273
	v10 = m.G0
	v12 = v10 - int32(80)
	m.G0 = v12
	v16 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v12+int32(70)))) = v16
	*(*int64)(unsafe.Add(mBase, uint32(v12+int32(64)))) = v16
	*(*int64)(unsafe.Add(mBase, uint32(v12+int32(56)))) = v16
	*(*int64)(unsafe.Add(mBase, uint32(v12+int32(48)))) = v16
	*(*int64)(unsafe.Add(mBase, uint32(v12)+40)) = v16
	*(*int64)(unsafe.Add(mBase, uint32(v12)+32)) = v16
	v34 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l2)+2248)))
	v35 = F___bswap_16_2(m, v34)
	mBase = m.M
	goto L1
L1:
	;
	v39 = *(*int32)(unsafe.Add(mBase, _consts[165]))
	if v39 != 0 {
		goto L2
	} else {
		goto L3
	}
L2:
	;
	v40 = int32(10)
	goto L4
L3:
	;
	v40 = int32(2246)
	goto L4
L4:
	;
	v42 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l2+v40))))
	v43 = F___bswap_16_2(m, v42)
	mBase = m.M
	goto L5
L5:
	;
	if v39 != 0 {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	v46 = int32(2246)
	goto L8
L7:
	;
	v46 = int32(10)
	goto L8
L8:
	;
	v48 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l2+v46))))
	v49 = F___bswap_16_2(m, v48)
	mBase = m.M
	goto L9
L9:
	;
	v50 = int32(0)
	v51 = *(*int32)(unsafe.Add(mBase, uint32(l0)+2344))
	if l1 == v51 {
		v273 = v50
		goto L10
	} else {
		goto L11
	}
L10:
	;
	m.G0 = v12 + int32(80)
	return v273
L11:
	;
	v53 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+2168)))
	if v53 == int32(0) {
		goto L13
	} else {
		goto L14
	}
L12:
	;
	v135 = *(*int32)(unsafe.Add(mBase, uint32(l0)+2324))
	if v135 != v49 {
		goto L27
	} else {
		goto L28
	}
L13:
	;
	v98 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	if v98 == int32(0) {
		goto L15
	} else {
		goto L16
	}
L14:
	;
	v64 = *(*int64)(unsafe.Add(mBase, uint32(l2+int32(2205))))
	*(*int64)(unsafe.Add(mBase, uint32(v12+int32(69)))) = v64
	v72 = *(*int64)(unsafe.Add(mBase, uint32(l2+int32(2200))))
	*(*int64)(unsafe.Add(mBase, uint32(v12+int32(64)))) = v72
	v80 = *(*int64)(unsafe.Add(mBase, uint32(l2+int32(2192))))
	*(*int64)(unsafe.Add(mBase, uint32(v12+int32(56)))) = v80
	v88 = *(*int64)(unsafe.Add(mBase, uint32(l2+int32(2184))))
	*(*int64)(unsafe.Add(mBase, uint32(v12+int32(48)))) = v88
	v90 = *(*int64)(unsafe.Add(mBase, uint32(l2+int32(2168))))
	*(*int64)(unsafe.Add(mBase, uint32(v12)+32)) = v90
	v94 = *(*int64)(unsafe.Add(mBase, uint32(l2+int32(2176))))
	*(*int64)(unsafe.Add(mBase, uint32(v12)+40)) = v94
	v96 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v12)+77)) = uint8(v96)
	goto L12
L15:
	;
	v118 = *(*int32)(unsafe.Add(mBase, _consts[6]))
	if int32(2) < v118 {
		v273 = v50
		goto L10
	} else {
		goto L21
	}
L16:
	;
	v101 = *(*int32)(unsafe.Add(mBase, uint32(v98)))
	v102 = *(*int32)(unsafe.Add(mBase, uint32(v101)+24))
	if v102 == int32(0) {
		goto L15
	} else {
		goto L17
	}
L17:
	;
	v110 = m.T0[v102].(func(*base.Module, int32, int32, int32, int32, int32) int32)(m, v98, v12+int32(32), int32(46), int32(0), int32(1))
	mBase = m.M
	v113 = m.ExcPending
	if v113 != 0 {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	return int32(0)
L19:
	;
	if v110 != int32(-1) {
		goto L12
	} else {
		goto L20
	}
L20:
	;
	goto L15
L21:
	;
	v121 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	if v121 != 0 {
		goto L23
	} else {
		goto L24
	}
L22:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12))) = v127
	F__serverLog(m, int32(2), int32(_a282), v12)
	mBase = m.M
	v132 = m.ExcPending
	if v132 != 0 {
		goto L18
	} else {
		goto L26
	}
L23:
	;
	v123 = *(*int32)(unsafe.Add(mBase, uint32(v121)))
	v124 = *(*int32)(unsafe.Add(mBase, uint32(v123)+88))
	v125 = m.T0[v124].(func(*base.Module, int32) int32)(m, v121)
	mBase = m.M
	v126 = m.ExcPending
	if v126 != 0 {
		goto L18
	} else {
		goto L25
	}
L24:
	;
	v127 = int32(_a283)
	goto L22
L25:
	;
	v127 = v125
	goto L22
L26:
	;
	v273 = v50
	goto L10
L27:
	;
	v178 = *(*int64)(unsafe.Add(mBase, uint32(v12)+32))
	*(*int64)(unsafe.Add(mBase, uint32(l0)+2256)) = v178
	v182 = *(*int64)(unsafe.Add(mBase, uint32(v12)+40))
	*(*int64)(unsafe.Add(mBase, uint32(l0+int32(2264)))) = v182
	v188 = *(*int64)(unsafe.Add(mBase, uint32(v12+int32(70))))
	*(*int64)(unsafe.Add(mBase, uint32(l0+int32(2294)))) = v188
	v194 = *(*int64)(unsafe.Add(mBase, uint32(v12+int32(64))))
	*(*int64)(unsafe.Add(mBase, uint32(l0+int32(2288)))) = v194
	v200 = *(*int64)(unsafe.Add(mBase, uint32(v12+int32(56))))
	*(*int64)(unsafe.Add(mBase, uint32(l0+int32(2280)))) = v200
	v206 = *(*int64)(unsafe.Add(mBase, uint32(v12+int32(48))))
	*(*int64)(unsafe.Add(mBase, uint32(l0+int32(2272)))) = v206
	*(*int32)(unsafe.Add(mBase, uint32(l0)+2332)) = v35
	*(*int32)(unsafe.Add(mBase, uint32(l0)+2328)) = v43
	*(*int32)(unsafe.Add(mBase, uint32(l0)+2324)) = v49
	v211 = *(*int32)(unsafe.Add(mBase, uint32(l0)+2344))
	if v211 == int32(0) {
		goto L40
	} else {
		goto L41
	}
L28:
	;
	v137 = *(*int32)(unsafe.Add(mBase, uint32(l0)+2332))
	if v137 != v35 {
		goto L27
	} else {
		goto L29
	}
L29:
	;
	v139 = *(*int32)(unsafe.Add(mBase, uint32(l0)+2328))
	if v139 != v43 {
		goto L27
	} else {
		goto L30
	}
L30:
	;
	v142 = v12 + int32(32)
	v144 = l0 + int32(2256)
	v147 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v144))))
	v148 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v142))))
	if v148 == int32(0) {
		v171 = v147
		v172 = v148
		goto L32
	} else {
		goto L33
	}
L31:
	;
	if v172-v171&int32(255) == int32(0) {
		v273 = v50
		goto L10
	} else {
		goto L39
	}
L32:
	;
	goto L31
L33:
	;
	if v148 != v147&int32(255) {
		v171 = v147
		v172 = v148
		goto L32
	} else {
		goto L34
	}
L34:
	;
	v154 = v142
	v155 = v144
	goto L35
L35:
	;
	v158 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v155)+1)))
	v159 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v154)+1)))
	if v159 == int32(0) {
		v171 = v158
		v172 = v159
		goto L32
	} else {
		goto L37
	}
L36:
	;
	v171 = v158
	v172 = v159
	goto L32
L37:
	;
	v162 = int32(1)
	if v159 == v158&int32(255) {
		v154 = v154 + v162
		v155 = v155 + v162
		goto L35
	} else {
		goto L38
	}
L38:
	;
	goto L36
L39:
	;
	goto L27
L40:
	;
	v217 = l0 + int32(2256)
	v218 = *(*int32)(unsafe.Add(mBase, uint32(l0)+88))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+88)) = v218 & int32(-65)
	v223 = *(*int32)(unsafe.Add(mBase, _consts[6]))
	if int32(2) < v223 {
		goto L43
	} else {
		goto L44
	}
L41:
	;
	F_freeClusterLink(m, v211)
	mBase = m.M
	v215 = m.ExcPending
	if v215 != 0 {
		goto L18
	} else {
		goto L42
	}
L42:
	;
	goto L40
L43:
	;
	v249 = int32(1)
	v250 = int32(0)
	v251 = *(*int32)(unsafe.Add(mBase, _consts[154]))
	v252 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v251)+88)))
	if v252&int32(2) == v250 {
		v273 = v249
		goto L10
	} else {
		goto L50
	}
L44:
	;
	v226 = F_humanNodename(m, l0)
	mBase = m.M
	v227 = m.ExcPending
	if v227 != 0 {
		goto L18
	} else {
		goto L45
	}
L45:
	;
	v231 = *(*int32)(unsafe.Add(mBase, _consts[165]))
	if v231 != 0 {
		goto L46
	} else {
		goto L47
	}
L46:
	;
	v232 = int32(2328)
	goto L48
L47:
	;
	v232 = int32(2324)
	goto L48
L48:
	;
	v234 = *(*int32)(unsafe.Add(mBase, uint32(l0+v232)))
	*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = l0 + int32(8)
	*(*int32)(unsafe.Add(mBase, uint32(v12)+20)) = v226
	*(*int32)(unsafe.Add(mBase, uint32(v12)+24)) = v217
	*(*int32)(unsafe.Add(mBase, uint32(v12)+28)) = v234
	F__serverLog(m, int32(2), int32(_a284), v12+int32(16))
	mBase = m.M
	v246 = m.ExcPending
	if v246 != 0 {
		goto L18
	} else {
		goto L49
	}
L49:
	;
	goto L43
L50:
	;
	v257 = *(*int32)(unsafe.Add(mBase, uint32(v251)+2172))
	if v257 != l0 {
		v273 = v249
		goto L10
	} else {
		goto L51
	}
L51:
	;
	v262 = *(*int32)(unsafe.Add(mBase, _consts[149]))
	if v262 != 0 {
		goto L52
	} else {
		goto L53
	}
L52:
	;
	v263 = int32(2328)
	goto L54
L53:
	;
	v263 = int32(2324)
	goto L54
L54:
	;
	v265 = *(*int32)(unsafe.Add(mBase, uint32(l0+v263)))
	v266 = int32(0)
	F_replicationSetPrimary(m, v217, v265, v266, v266)
	mBase = m.M
	v269 = m.ExcPending
	if v269 != 0 {
		goto L18
	} else {
		goto L55
	}
L55:
	;
	v273 = v249
	goto L10
}
