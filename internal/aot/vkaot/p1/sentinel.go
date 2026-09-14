package p1

import (
	base "github.com/shibukawa/vkmem/internal/aot/vkaot/base"
	"unsafe"
)

func F_getSentinelValkeyInstanceByAddrAndRunID(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
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
	var v67 int64
	_ = v67
	var v69 int64
	_ = v69
	var v71 int64
	_ = v71
	var v73 int64
	_ = v73
	var v75 int64
	_ = v75
	var v77 int64
	_ = v77
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v99 int32
	_ = v99
	var v103 int32
	_ = v103
	var v105 int32
	_ = v105
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v119 int32
	_ = v119
	var v124 int32
	_ = v124
	var v128 int32
	_ = v128
	var v132 int32
	_ = v132
	var v138 int32
	_ = v138
	var v141 int32
	_ = v141
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v161 int32
	_ = v161
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v201 int32
	_ = v201
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v221 int32
	_ = v221
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v227 int32
	_ = v227
	var v231 int32
	_ = v231
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
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
	var v260 int32
	_ = v260
	var v263 int32
	_ = v263
	var v266 int32
	_ = v266
	var v268 int32
	_ = v268
	var v270 int32
	_ = v270
	var v271 int32
	_ = v271
	var v273 int32
	_ = v273
	var v275 int32
	_ = v275
	var v276 int32
	_ = v276
	var v283 int32
	_ = v283
	if l1|l3 == int32(0) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	F__serverAssert(m, int32(_a_F_getSentinelValkeyInstanceByAddrAndRunID_0), int32(_a_F_getSentinelValkeyInstanceByAddrAndRunID_1), int32(1481))
	mBase = m.M
	v283 = m.ExcPending
	if v283 != 0 {
		goto L5
	} else {
		goto L86
	}
L2:
	;
	v10 = int32(0)
	if l1 == v10 {
		v20 = v10
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v21 = F_dictGetIterator(m, l0)
	mBase = m.M
	v22 = m.ExcPending
	if v22 != 0 {
		goto L5
	} else {
		goto L8
	}
L4:
	;
	v14 = F_createSentinelAddr(m, l1, l2, int32(1))
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	return int32(0)
L6:
	;
	if v14 != 0 {
		v20 = v14
		goto L3
	} else {
		goto L7
	}
L7:
	;
	return int32(0)
L8:
	;
	goto L13
L9:
	;
	return v276
L10:
	;
	v268 = *(*int32)(unsafe.Add(mBase, uint32(v20)))
	F_sdsfree(m, v268)
	mBase = m.M
	v270 = m.ExcPending
	if v270 != 0 {
		goto L5
	} else {
		goto L83
	}
L11:
	;
	F_dictReleaseIterator(m, v21)
	mBase = m.M
	v263 = m.ExcPending
	if v263 != 0 {
		goto L5
	} else {
		goto L81
	}
L12:
	;
	v260 = int32(0)
	goto L11
L13:
	;
	v36 = v21 + int32(20)
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v21)+16))
	if v37 != 0 {
		goto L17
	} else {
		goto L18
	}
L14:
	;
	F_dictReleaseIterator(m, v21)
	mBase = m.M
	v258 = m.ExcPending
	if v258 != 0 {
		goto L5
	} else {
		goto L80
	}
L15:
	;
	if v132 == int32(0) {
		goto L12
	} else {
		goto L41
	}
L16:
	;
	v43 = v36
	v44 = v40
	goto L19
L17:
	;
	v40 = int32(1)
	goto L16
L18:
	;
	v40 = int32(0)
	goto L16
L19:
	;
	switch v44 {
	case 0:
		goto L24
	default:
		goto L23
	}
L21:
	;
	v44 = int32(0)
	goto L19
L22:
	;
	goto L15
L23:
	;
	v124 = *(*int32)(unsafe.Add(mBase, uint32(v43)))
	*(*int32)(unsafe.Add(mBase, uint32(v21)+16)) = v124
	if v124 == int32(0) {
		goto L21
	} else {
		goto L40
	}
L24:
	;
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v21)+4))
	if v48 != int32(-1) {
		v87 = v48
		goto L25
	} else {
		goto L26
	}
L25:
	;
	v88 = int32(1)
	v89 = v87 + v88
	*(*int32)(unsafe.Add(mBase, uint32(v21)+4)) = v89
	v91 = int32(0)
	v94 = *(*int32)(unsafe.Add(mBase, uint32(v21)))
	v95 = *(*int32)(unsafe.Add(mBase, uint32(v21)+8))
	v99 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v94+v95+int32(26)))))
	if v99 == int32(255) {
		goto L34
	} else {
		goto L35
	}
L26:
	;
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v21)+8))
	if v52 != 0 {
		v87 = int32(-1)
		goto L25
	} else {
		goto L27
	}
L27:
	;
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v21)))
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v21)+12))
	if v54 == int32(0) {
		goto L29
	} else {
		goto L30
	}
L28:
	;
	v81 = *(*int32)(unsafe.Add(mBase, uint32(v80)+20))
	if v81 != int32(-1) {
		goto L31
	} else {
		goto L32
	}
L29:
	;
	v61 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v53)+16)))
	v62 = int64(*(*int8)(unsafe.Add(mBase, uint32(v53)+27)))
	v63 = int64(*(*int32)(unsafe.Add(mBase, uint32(v53)+8)))
	v64 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v53)+12)))
	v65 = int64(*(*int8)(unsafe.Add(mBase, uint32(v53)+26)))
	v66 = int64(*(*int32)(unsafe.Add(mBase, uint32(v53)+4)))
	v67 = F_wangHash64(m, v66)
	mBase = m.M
	v69 = F_wangHash64(m, v65+v67)
	mBase = m.M
	v71 = F_wangHash64(m, v64+v69)
	mBase = m.M
	v73 = F_wangHash64(m, v63+v71)
	mBase = m.M
	v75 = F_wangHash64(m, v62+v73)
	mBase = m.M
	v77 = F_wangHash64(m, v61+v75)
	mBase = m.M
	*(*int64)(unsafe.Add(mBase, uint32(v21)+24)) = v77
	v79 = *(*int32)(unsafe.Add(mBase, uint32(v21)))
	v80 = v79
	goto L28
L30:
	;
	v57 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v53)+24)))
	v59 = v57 + int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v53)+24)) = uint16(v59)
	v80 = v53
	goto L28
L31:
	;
	v87 = v81 + int32(-1)
	goto L25
L32:
	;
	v84 = *(*int32)(unsafe.Add(mBase, uint32(v21)+4))
	v87 = v84
	goto L25
L33:
	;
	v114 = int32(2)
	v119 = *(*int32)(unsafe.Add(mBase, uint32(v94+v112<<(uint(v114)%32)+int32(4))))
	v43 = v119 + v113<<(uint(v114)%32)
	v44 = int32(1)
	goto L19
L34:
	;
	v103 = v91
	goto L36
L35:
	;
	v103 = v88 << (uint(v99) % 32)
	goto L36
L36:
	;
	if v89 < v103 {
		v112 = v95
		v113 = v89
		goto L33
	} else {
		goto L37
	}
L37:
	;
	if v95 != 0 {
		v132 = v91
		goto L22
	} else {
		goto L38
	}
L38:
	;
	v105 = *(*int32)(unsafe.Add(mBase, uint32(v94)+20))
	if v105 == int32(-1) {
		v132 = v91
		goto L22
	} else {
		goto L39
	}
L39:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v21)+4)) = int64(4294967296)
	v112 = int32(1)
	v113 = int32(0)
	goto L33
L40:
	;
	v128 = *(*int32)(unsafe.Add(mBase, uint32(v124)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v36))) = v128
	v132 = v124
	goto L22
L41:
	;
	v138 = *(*int32)(unsafe.Add(mBase, uint32(v132)+8))
	goto L42
L42:
	;
	if l3 == int32(0) {
		goto L43
	} else {
		goto L44
	}
L43:
	;
	if l1 == int32(0) {
		v260 = v138
		goto L11
	} else {
		goto L55
	}
L44:
	;
	v141 = *(*int32)(unsafe.Add(mBase, uint32(v138)+8))
	if v141 == int32(0) {
		goto L13
	} else {
		goto L45
	}
L45:
	;
	v146 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3))))
	v147 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v141))))
	if v147 == int32(0) {
		v170 = v146
		v171 = v147
		goto L47
	} else {
		goto L48
	}
L46:
	;
	if v171-v170&int32(255) != 0 {
		goto L13
	} else {
		goto L54
	}
L47:
	;
	goto L46
L48:
	;
	if v147 != v146&int32(255) {
		v170 = v146
		v171 = v147
		goto L47
	} else {
		goto L49
	}
L49:
	;
	v153 = v141
	v154 = l3
	goto L50
L50:
	;
	v157 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v154)+1)))
	v158 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v153)+1)))
	if v158 == int32(0) {
		v170 = v157
		v171 = v158
		goto L47
	} else {
		goto L52
	}
L51:
	;
	v170 = v157
	v171 = v158
	goto L47
L52:
	;
	v161 = int32(1)
	if v158 == v157&int32(255) {
		v153 = v153 + v161
		v154 = v154 + v161
		goto L50
	} else {
		goto L53
	}
L53:
	;
	goto L51
L54:
	;
	goto L43
L55:
	;
	v178 = *(*int32)(unsafe.Add(mBase, uint32(v138)+24))
	v179 = *(*int32)(unsafe.Add(mBase, uint32(v178)+8))
	v180 = *(*int32)(unsafe.Add(mBase, uint32(v20)+8))
	if v179 != v180 {
		goto L13
	} else {
		goto L56
	}
L56:
	;
	v182 = *(*int32)(unsafe.Add(mBase, uint32(v178)+4))
	v183 = *(*int32)(unsafe.Add(mBase, uint32(v20)+4))
	v186 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v183))))
	v187 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v182))))
	if v187 == int32(0) {
		v210 = v186
		v211 = v187
		goto L59
	} else {
		goto L60
	}
L57:
	;
	goto L14
L58:
	;
	if v211-v210&int32(255) == int32(0) {
		goto L57
	} else {
		goto L66
	}
L59:
	;
	goto L58
L60:
	;
	if v187 != v186&int32(255) {
		v210 = v186
		v211 = v187
		goto L59
	} else {
		goto L61
	}
L61:
	;
	v193 = v182
	v194 = v183
	goto L62
L62:
	;
	v197 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v194)+1)))
	v198 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v193)+1)))
	if v198 == int32(0) {
		v210 = v197
		v211 = v198
		goto L59
	} else {
		goto L64
	}
L63:
	;
	v210 = v197
	v211 = v198
	goto L59
L64:
	;
	v201 = int32(1)
	if v198 == v197&int32(255) {
		v193 = v193 + v201
		v194 = v194 + v201
		goto L62
	} else {
		goto L65
	}
L65:
	;
	goto L63
L66:
	;
	v217 = *(*int32)(unsafe.Add(mBase, uint32(v178)))
	v218 = *(*int32)(unsafe.Add(mBase, uint32(v20)))
	v221 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v217))))
	if v221 != 0 {
		goto L69
	} else {
		goto L70
	}
L67:
	;
	if v253-v255 != 0 {
		goto L13
	} else {
		goto L79
	}
L68:
	;
	v253 = F_tolower(m, v249)
	mBase = m.M
	v254 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v250))))
	v255 = F_tolower(m, v254)
	mBase = m.M
	goto L67
L69:
	;
	v223 = v217
	v224 = v218
	v225 = v221
	goto L72
L70:
	;
	v249 = int32(0)
	v250 = v218
	goto L68
L71:
	;
	v249 = v246 & int32(255)
	v250 = v245
	goto L68
L72:
	;
	v227 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v224))))
	if v227 == int32(0) {
		v245 = v224
		v246 = v225
		goto L71
	} else {
		goto L74
	}
L73:
	;
	v245 = v239
	v246 = int32(0)
	goto L71
L74:
	;
	v231 = v225 & int32(255)
	if v231 == v227 {
		goto L75
	} else {
		goto L76
	}
L75:
	;
	v238 = int32(1)
	v239 = v224 + v238
	v240 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v223)+1)))
	if v240 != 0 {
		v223 = v223 + v238
		v224 = v239
		v225 = v240
		goto L72
	} else {
		goto L78
	}
L76:
	;
	v233 = F_tolower(m, v231)
	mBase = m.M
	v234 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v224))))
	v235 = F_tolower(m, v234)
	mBase = m.M
	if v233 == v235 {
		goto L75
	} else {
		goto L77
	}
L77:
	;
	v237 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v223))))
	v245 = v224
	v246 = v237
	goto L71
L78:
	;
	goto L73
L79:
	;
	goto L57
L80:
	;
	v266 = v138
	goto L10
L81:
	;
	if v20 == int32(0) {
		v276 = v260
		goto L9
	} else {
		goto L82
	}
L82:
	;
	v266 = v260
	goto L10
L83:
	;
	v271 = *(*int32)(unsafe.Add(mBase, uint32(v20)+4))
	F_sdsfree(m, v271)
	mBase = m.M
	v273 = m.ExcPending
	if v273 != 0 {
		goto L5
	} else {
		goto L84
	}
L84:
	;
	F_valkey_free(m, v20)
	mBase = m.M
	v275 = m.ExcPending
	if v275 != 0 {
		goto L5
	} else {
		goto L85
	}
L85:
	;
	v276 = v266
	goto L9
L86:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_initSentinel(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v9 int64
	_ = v9
	var v17 int64
	_ = v17
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v28 int64
	_ = v28
	*(*int64)(unsafe.Add(mBase, _c_F_initSentinel[0])) = int64(0)
	v6 = F_dictCreate(m, int32(_a_F_initSentinel_0))
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return
	} else {
		v8 = int32(0)
		v9 = int64(0)
		*(*int64)(unsafe.Add(mBase, _c_F_initSentinel[1])) = v9
		*(*int32)(unsafe.Add(mBase, _c_F_initSentinel[2])) = v6
		*(*int64)(unsafe.Add(mBase, _c_F_initSentinel[3])) = v9
		v17 = F_mstime(m)
		mBase = m.M
		*(*int64)(unsafe.Add(mBase, _c_F_initSentinel[4])) = v17
		*(*int32)(unsafe.Add(mBase, _c_F_initSentinel[5])) = v8
		v22 = F_listCreate(m)
		mBase = m.M
		v23 = m.ExcPending
		if v23 != 0 {
			return
		} else {
			v24 = int32(0)
			*(*int64)(unsafe.Add(mBase, _c_F_initSentinel[6])) = int64(4294967296)
			v28 = int64(0)
			*(*int64)(unsafe.Add(mBase, _c_F_initSentinel[7])) = v28
			*(*int32)(unsafe.Add(mBase, _c_F_initSentinel[8])) = v22
			*(*int64)(unsafe.Add(mBase, _c_F_initSentinel[9])) = v28
			*(*int64)(unsafe.Add(mBase, _c_F_initSentinel[10])) = v28
			*(*int64)(unsafe.Add(mBase, _c_F_initSentinel[11])) = v28
			*(*int64)(unsafe.Add(mBase, _c_F_initSentinel[12])) = v28
			*(*int64)(unsafe.Add(mBase, _c_F_initSentinel[13])) = v28
			*(*int64)(unsafe.Add(mBase, _c_F_initSentinel[14])) = v28
			*(*int64)(unsafe.Add(mBase, _c_F_initSentinel[15])) = v28
			*(*int32)(unsafe.Add(mBase, _c_F_initSentinel[16])) = v24
			*(*uint8)(unsafe.Add(mBase, _c_F_initSentinel[17])) = uint8(v24)
			return
		}
	}
}
func F_initSentinelConfig(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v1 int32
	_ = v1
	v1 = int32(_a_F_initSentinelConfig_0)
	*(*int32)(unsafe.Add(mBase, _c_F_initSentinelConfig[0])) = int32(0)
	*(*int32)(unsafe.Add(mBase, _c_F_initSentinelConfig[1])) = int32(26379)
	return
}
func F_sentinelAbortFailover(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v16 int64
	_ = v16
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v31 int32
	_ = v31
	var v37 int32
	_ = v37
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v3&int32(64) == int32(0) {
		F__serverAssert(m, int32(_a_F_sentinelAbortFailover_0), int32(_a_F_sentinelAbortFailover_1), int32(5359))
		mBase = m.M
		v31 = m.ExcPending
		if v31 != 0 {
			return
		} else {
			F_abort(m)
			mBase = m.M
			base.Wasm_trap_unreachable()
			for {
			}
		}
	} else {
		v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+240))
		if int32(5) <= v8 {
			F__serverAssert(m, int32(_a_F_sentinelAbortFailover_2), int32(_a_F_sentinelAbortFailover_1), int32(5360))
			mBase = m.M
			v37 = m.ExcPending
			if v37 != 0 {
				return
			} else {
				F_abort(m)
				mBase = m.M
				base.Wasm_trap_unreachable()
				for {
				}
			}
		} else {
			v11 = int32(0)
			*(*int32)(unsafe.Add(mBase, uint32(l0)+240)) = v11
			*(*int32)(unsafe.Add(mBase, uint32(l0))) = v3 & int32(-18497)
			v16 = F_mstime(m)
			mBase = m.M
			*(*int64)(unsafe.Add(mBase, uint32(l0)+248)) = v16
			v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+280))
			if v18 == v11 {
			} else {
				v21 = *(*int32)(unsafe.Add(mBase, uint32(v18)))
				*(*int32)(unsafe.Add(mBase, uint32(v18))) = v21 & int32(-129)
				*(*int32)(unsafe.Add(mBase, uint32(l0)+280)) = int32(0)
			}
			return
		}
	}
}
func F_sentinelAskPrimaryStateToOtherSentinels(m *base.Module, l0 int32, l1 int32) {
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
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v25 int64
	_ = v25
	var v34 int32
	_ = v34
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int64
	_ = v45
	var v47 int32
	_ = v47
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v89 int32
	_ = v89
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v102 int64
	_ = v102
	var v103 int64
	_ = v103
	var v104 int64
	_ = v104
	var v105 int64
	_ = v105
	var v106 int64
	_ = v106
	var v107 int64
	_ = v107
	var v108 int64
	_ = v108
	var v110 int64
	_ = v110
	var v112 int64
	_ = v112
	var v114 int64
	_ = v114
	var v116 int64
	_ = v116
	var v118 int64
	_ = v118
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v125 int32
	_ = v125
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v132 int32
	_ = v132
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v140 int32
	_ = v140
	var v144 int32
	_ = v144
	var v146 int32
	_ = v146
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v160 int32
	_ = v160
	var v165 int32
	_ = v165
	var v169 int32
	_ = v169
	var v173 int32
	_ = v173
	var v189 int32
	_ = v189
	var v196 int32
	_ = v196
	var v197 int64
	_ = v197
	var v198 int64
	_ = v198
	var v201 int64
	_ = v201
	var v205 int32
	_ = v205
	var v209 int32
	_ = v209
	var v211 int32
	_ = v211
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v216 int64
	_ = v216
	var v217 int64
	_ = v217
	var v220 int64
	_ = v220
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v226 int32
	_ = v226
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
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v246 int64
	_ = v246
	var v252 int32
	_ = v252
	var v255 int32
	_ = v255
	var v263 int32
	_ = v263
	var v264 int32
	_ = v264
	var v265 int32
	_ = v265
	var v266 int32
	_ = v266
	var v282 int32
	_ = v282
	var v283 int32
	_ = v283
	var v286 int32
	_ = v286
	var v289 int32
	_ = v289
	var v290 int32
	_ = v290
	var v294 int32
	_ = v294
	var v298 int32
	_ = v298
	var v299 int32
	_ = v299
	var v300 int32
	_ = v300
	var v303 int32
	_ = v303
	var v305 int32
	_ = v305
	var v307 int64
	_ = v307
	var v308 int64
	_ = v308
	var v309 int64
	_ = v309
	var v310 int64
	_ = v310
	var v311 int64
	_ = v311
	var v312 int64
	_ = v312
	var v313 int64
	_ = v313
	var v315 int64
	_ = v315
	var v317 int64
	_ = v317
	var v319 int64
	_ = v319
	var v321 int64
	_ = v321
	var v323 int64
	_ = v323
	var v325 int32
	_ = v325
	var v326 int32
	_ = v326
	var v327 int32
	_ = v327
	var v330 int32
	_ = v330
	var v333 int32
	_ = v333
	var v334 int32
	_ = v334
	var v335 int32
	_ = v335
	var v337 int32
	_ = v337
	var v340 int32
	_ = v340
	var v341 int32
	_ = v341
	var v345 int32
	_ = v345
	var v349 int32
	_ = v349
	var v351 int32
	_ = v351
	var v358 int32
	_ = v358
	var v359 int32
	_ = v359
	var v360 int32
	_ = v360
	var v365 int32
	_ = v365
	var v370 int32
	_ = v370
	var v374 int32
	_ = v374
	var v378 int32
	_ = v378
	var v394 int32
	_ = v394
	v12 = m.G0
	v14 = v12 - int32(64)
	m.G0 = v14
	v16 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v16&int32(8) == int32(0) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	m.G0 = v14 + int32(64)
	return
L2:
	;
	v21 = int32(32)
	v22 = v14 + v21
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v25 = int64(*(*int32)(unsafe.Add(mBase, uint32(v24)+8)))
	if v25 <= int64(-1) {
		goto L7
	} else {
		goto L8
	}
L3:
	;
	v67 = *(*int32)(unsafe.Add(mBase, uint32(l0)+144))
	v68 = F_dictGetIterator(m, v67)
	mBase = m.M
	v69 = m.ExcPending
	if v69 != 0 {
		goto L13
	} else {
		goto L14
	}
L4:
	;
	goto L3
L6:
	;
	v47 = F_ull2string(m, v43, v44, v45)
	mBase = m.M
	if v47 == int32(0) {
		goto L4
	} else {
		goto L10
	}
L7:
	;
	goto L9
L8:
	;
	v43 = v22
	v44 = v21
	v45 = v25
	goto L6
L9:
	;
	v34 = int32(45)
	*(*uint8)(unsafe.Add(mBase, uint32(v22))) = uint8(v34)
	v43 = v14 + int32(33)
	v44 = int32(31)
	v45 = int64(0) - v25
	goto L6
L10:
	;
	goto L3
L12:
	;
	F_dictReleaseIterator(m, v68)
	mBase = m.M
	v394 = m.ExcPending
	if v394 != 0 {
		goto L13
	} else {
		goto L94
	}
L13:
	;
	return
L14:
	;
	v77 = v68 + int32(20)
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v68)+16))
	if v78 != 0 {
		goto L17
	} else {
		goto L18
	}
L15:
	;
	if v173 == int32(0) {
		goto L12
	} else {
		goto L41
	}
L16:
	;
	v84 = v77
	v85 = v81
	goto L19
L17:
	;
	v81 = int32(1)
	goto L16
L18:
	;
	v81 = int32(0)
	goto L16
L19:
	;
	switch v85 {
	case 0:
		goto L24
	default:
		goto L23
	}
L21:
	;
	v85 = int32(0)
	goto L19
L22:
	;
	goto L15
L23:
	;
	v165 = *(*int32)(unsafe.Add(mBase, uint32(v84)))
	*(*int32)(unsafe.Add(mBase, uint32(v68)+16)) = v165
	if v165 == int32(0) {
		goto L21
	} else {
		goto L40
	}
L24:
	;
	v89 = *(*int32)(unsafe.Add(mBase, uint32(v68)+4))
	if v89 != int32(-1) {
		v128 = v89
		goto L25
	} else {
		goto L26
	}
L25:
	;
	v129 = int32(1)
	v130 = v128 + v129
	*(*int32)(unsafe.Add(mBase, uint32(v68)+4)) = v130
	v132 = int32(0)
	v135 = *(*int32)(unsafe.Add(mBase, uint32(v68)))
	v136 = *(*int32)(unsafe.Add(mBase, uint32(v68)+8))
	v140 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v135+v136+int32(26)))))
	if v140 == int32(255) {
		goto L34
	} else {
		goto L35
	}
L26:
	;
	v93 = *(*int32)(unsafe.Add(mBase, uint32(v68)+8))
	if v93 != 0 {
		v128 = int32(-1)
		goto L25
	} else {
		goto L27
	}
L27:
	;
	v94 = *(*int32)(unsafe.Add(mBase, uint32(v68)))
	v95 = *(*int32)(unsafe.Add(mBase, uint32(v68)+12))
	if v95 == int32(0) {
		goto L29
	} else {
		goto L30
	}
L28:
	;
	v122 = *(*int32)(unsafe.Add(mBase, uint32(v121)+20))
	if v122 != int32(-1) {
		goto L31
	} else {
		goto L32
	}
L29:
	;
	v102 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v94)+16)))
	v103 = int64(*(*int8)(unsafe.Add(mBase, uint32(v94)+27)))
	v104 = int64(*(*int32)(unsafe.Add(mBase, uint32(v94)+8)))
	v105 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v94)+12)))
	v106 = int64(*(*int8)(unsafe.Add(mBase, uint32(v94)+26)))
	v107 = int64(*(*int32)(unsafe.Add(mBase, uint32(v94)+4)))
	v108 = F_wangHash64(m, v107)
	mBase = m.M
	v110 = F_wangHash64(m, v106+v108)
	mBase = m.M
	v112 = F_wangHash64(m, v105+v110)
	mBase = m.M
	v114 = F_wangHash64(m, v104+v112)
	mBase = m.M
	v116 = F_wangHash64(m, v103+v114)
	mBase = m.M
	v118 = F_wangHash64(m, v102+v116)
	mBase = m.M
	*(*int64)(unsafe.Add(mBase, uint32(v68)+24)) = v118
	v120 = *(*int32)(unsafe.Add(mBase, uint32(v68)))
	v121 = v120
	goto L28
L30:
	;
	v98 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v94)+24)))
	v100 = v98 + int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v94)+24)) = uint16(v100)
	v121 = v94
	goto L28
L31:
	;
	v128 = v122 + int32(-1)
	goto L25
L32:
	;
	v125 = *(*int32)(unsafe.Add(mBase, uint32(v68)+4))
	v128 = v125
	goto L25
L33:
	;
	v155 = int32(2)
	v160 = *(*int32)(unsafe.Add(mBase, uint32(v135+v153<<(uint(v155)%32)+int32(4))))
	v84 = v160 + v154<<(uint(v155)%32)
	v85 = int32(1)
	goto L19
L34:
	;
	v144 = v132
	goto L36
L35:
	;
	v144 = v129 << (uint(v140) % 32)
	goto L36
L36:
	;
	if v130 < v144 {
		v153 = v136
		v154 = v130
		goto L33
	} else {
		goto L37
	}
L37:
	;
	if v136 != 0 {
		v173 = v132
		goto L22
	} else {
		goto L38
	}
L38:
	;
	v146 = *(*int32)(unsafe.Add(mBase, uint32(v135)+20))
	if v146 == int32(-1) {
		v173 = v132
		goto L22
	} else {
		goto L39
	}
L39:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v68)+4)) = int64(4294967296)
	v153 = int32(1)
	v154 = int32(0)
	goto L33
L40:
	;
	v169 = *(*int32)(unsafe.Add(mBase, uint32(v165)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v77))) = v169
	v173 = v165
	goto L22
L41:
	;
	v189 = v173
	goto L42
L42:
	;
	v196 = *(*int32)(unsafe.Add(mBase, uint32(v189)+8))
	goto L44
L43:
	;
	goto L12
L44:
	;
	v197 = F_mstime(m)
	mBase = m.M
	v198 = *(*int64)(unsafe.Add(mBase, uint32(v196)+48))
	v201 = *(*int64)(unsafe.Add(mBase, _c_F_sentinelAskPrimaryStateToOtherSentinels[0]))
	if v197-v198 <= v201*int64(5) {
		goto L45
	} else {
		goto L46
	}
L45:
	;
	v214 = *(*int32)(unsafe.Add(mBase, uint32(v196)+28))
	v215 = *(*int32)(unsafe.Add(mBase, uint32(v214)+4))
	if v215 != 0 {
		goto L48
	} else {
		goto L49
	}
L46:
	;
	v205 = *(*int32)(unsafe.Add(mBase, uint32(v196)))
	*(*int32)(unsafe.Add(mBase, uint32(v196))) = v205 & int32(-33)
	v209 = *(*int32)(unsafe.Add(mBase, uint32(v196)+216))
	F_sdsfree(m, v209)
	mBase = m.M
	v211 = m.ExcPending
	if v211 != 0 {
		goto L13
	} else {
		goto L47
	}
L47:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v196)+216)) = int32(0)
	goto L45
L48:
	;
	v282 = v68 + int32(20)
	v283 = *(*int32)(unsafe.Add(mBase, uint32(v68)+16))
	if v283 != 0 {
		goto L69
	} else {
		goto L70
	}
L49:
	;
	if l1&int32(1) != 0 {
		v223 = v214
		goto L50
	} else {
		goto L51
	}
L50:
	;
	v224 = *(*int32)(unsafe.Add(mBase, uint32(v223)+12))
	v226 = F_sdsnew(m, int32(_a_F_sentinelAskPrimaryStateToOtherSentinels_0))
	mBase = m.M
	v227 = m.ExcPending
	if v227 != 0 {
		goto L13
	} else {
		goto L53
	}
L51:
	;
	v216 = F_mstime(m)
	mBase = m.M
	v217 = *(*int64)(unsafe.Add(mBase, uint32(v196)+48))
	v220 = *(*int64)(unsafe.Add(mBase, _c_F_sentinelAskPrimaryStateToOtherSentinels[0]))
	if v216-v217 < v220 {
		goto L48
	} else {
		goto L52
	}
L52:
	;
	v222 = *(*int32)(unsafe.Add(mBase, uint32(v196)+28))
	v223 = v222
	goto L50
L53:
	;
	v228 = *(*int32)(unsafe.Add(mBase, uint32(v196)+192))
	if v228 != 0 {
		goto L54
	} else {
		goto L55
	}
L54:
	;
	v229 = v228
	goto L56
L55:
	;
	v229 = v196
	goto L56
L56:
	;
	v230 = *(*int32)(unsafe.Add(mBase, uint32(v229)+104))
	v231 = F_dictFetchValue(m, v230, v226)
	mBase = m.M
	v232 = m.ExcPending
	if v232 != 0 {
		goto L13
	} else {
		goto L57
	}
L57:
	;
	F_sdsfree(m, v226)
	mBase = m.M
	v234 = m.ExcPending
	if v234 != 0 {
		goto L13
	} else {
		goto L58
	}
L58:
	;
	v235 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v236 = int32(0)
	v237 = *(*int32)(unsafe.Add(mBase, _c_F_sentinelAskPrimaryStateToOtherSentinels[1]))
	v243 = *(*int32)(unsafe.Add(mBase, uint32(v235+base.B2i32(v237 == v236)<<(uint(int32(2))%32))))
	v244 = *(*int32)(unsafe.Add(mBase, uint32(l0)+240))
	v246 = *(*int64)(unsafe.Add(mBase, _c_F_sentinelAskPrimaryStateToOtherSentinels[2]))
	*(*int64)(unsafe.Add(mBase, uint32(v14+int32(16)))) = v246
	if v236 < v244 {
		goto L59
	} else {
		goto L60
	}
L59:
	;
	v252 = int32(_a_F_sentinelAskPrimaryStateToOtherSentinels_1)
	goto L61
L60:
	;
	v252 = int32(_a_F_sentinelAskPrimaryStateToOtherSentinels_2)
	goto L61
L61:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14+int32(24)))) = v252
	if v231 != 0 {
		goto L62
	} else {
		goto L63
	}
L62:
	;
	v255 = v231
	goto L64
L63:
	;
	v255 = int32(_a_F_sentinelAskPrimaryStateToOtherSentinels_0)
	goto L64
L64:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14))) = v255
	*(*int32)(unsafe.Add(mBase, uint32(v14)+4)) = v243
	*(*int32)(unsafe.Add(mBase, uint32(v14)+8)) = v14 + int32(32)
	v263 = F_valkeyAsyncCommand(m, v224, int32(1020), v196, int32(_a_F_sentinelAskPrimaryStateToOtherSentinels_3), v14)
	mBase = m.M
	v264 = m.ExcPending
	if v264 != 0 {
		goto L13
	} else {
		goto L65
	}
L65:
	;
	if v263 != 0 {
		goto L48
	} else {
		goto L66
	}
L66:
	;
	v265 = *(*int32)(unsafe.Add(mBase, uint32(v196)+28))
	v266 = *(*int32)(unsafe.Add(mBase, uint32(v265)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v265)+8)) = v266 + int32(1)
	goto L48
L67:
	;
	if v378 != 0 {
		v189 = v378
		goto L42
	} else {
		goto L93
	}
L68:
	;
	v289 = v282
	v290 = v286
	goto L71
L69:
	;
	v286 = int32(1)
	goto L68
L70:
	;
	v286 = int32(0)
	goto L68
L71:
	;
	switch v290 {
	case 0:
		goto L76
	default:
		goto L75
	}
L73:
	;
	v290 = int32(0)
	goto L71
L74:
	;
	goto L67
L75:
	;
	v370 = *(*int32)(unsafe.Add(mBase, uint32(v289)))
	*(*int32)(unsafe.Add(mBase, uint32(v68)+16)) = v370
	if v370 == int32(0) {
		goto L73
	} else {
		goto L92
	}
L76:
	;
	v294 = *(*int32)(unsafe.Add(mBase, uint32(v68)+4))
	if v294 != int32(-1) {
		v333 = v294
		goto L77
	} else {
		goto L78
	}
L77:
	;
	v334 = int32(1)
	v335 = v333 + v334
	*(*int32)(unsafe.Add(mBase, uint32(v68)+4)) = v335
	v337 = int32(0)
	v340 = *(*int32)(unsafe.Add(mBase, uint32(v68)))
	v341 = *(*int32)(unsafe.Add(mBase, uint32(v68)+8))
	v345 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v340+v341+int32(26)))))
	if v345 == int32(255) {
		goto L86
	} else {
		goto L87
	}
L78:
	;
	v298 = *(*int32)(unsafe.Add(mBase, uint32(v68)+8))
	if v298 != 0 {
		v333 = int32(-1)
		goto L77
	} else {
		goto L79
	}
L79:
	;
	v299 = *(*int32)(unsafe.Add(mBase, uint32(v68)))
	v300 = *(*int32)(unsafe.Add(mBase, uint32(v68)+12))
	if v300 == int32(0) {
		goto L81
	} else {
		goto L82
	}
L80:
	;
	v327 = *(*int32)(unsafe.Add(mBase, uint32(v326)+20))
	if v327 != int32(-1) {
		goto L83
	} else {
		goto L84
	}
L81:
	;
	v307 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v299)+16)))
	v308 = int64(*(*int8)(unsafe.Add(mBase, uint32(v299)+27)))
	v309 = int64(*(*int32)(unsafe.Add(mBase, uint32(v299)+8)))
	v310 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v299)+12)))
	v311 = int64(*(*int8)(unsafe.Add(mBase, uint32(v299)+26)))
	v312 = int64(*(*int32)(unsafe.Add(mBase, uint32(v299)+4)))
	v313 = F_wangHash64(m, v312)
	mBase = m.M
	v315 = F_wangHash64(m, v311+v313)
	mBase = m.M
	v317 = F_wangHash64(m, v310+v315)
	mBase = m.M
	v319 = F_wangHash64(m, v309+v317)
	mBase = m.M
	v321 = F_wangHash64(m, v308+v319)
	mBase = m.M
	v323 = F_wangHash64(m, v307+v321)
	mBase = m.M
	*(*int64)(unsafe.Add(mBase, uint32(v68)+24)) = v323
	v325 = *(*int32)(unsafe.Add(mBase, uint32(v68)))
	v326 = v325
	goto L80
L82:
	;
	v303 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v299)+24)))
	v305 = v303 + int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v299)+24)) = uint16(v305)
	v326 = v299
	goto L80
L83:
	;
	v333 = v327 + int32(-1)
	goto L77
L84:
	;
	v330 = *(*int32)(unsafe.Add(mBase, uint32(v68)+4))
	v333 = v330
	goto L77
L85:
	;
	v360 = int32(2)
	v365 = *(*int32)(unsafe.Add(mBase, uint32(v340+v358<<(uint(v360)%32)+int32(4))))
	v289 = v365 + v359<<(uint(v360)%32)
	v290 = int32(1)
	goto L71
L86:
	;
	v349 = v337
	goto L88
L87:
	;
	v349 = v334 << (uint(v345) % 32)
	goto L88
L88:
	;
	if v335 < v349 {
		v358 = v341
		v359 = v335
		goto L85
	} else {
		goto L89
	}
L89:
	;
	if v341 != 0 {
		v378 = v337
		goto L74
	} else {
		goto L90
	}
L90:
	;
	v351 = *(*int32)(unsafe.Add(mBase, uint32(v340)+20))
	if v351 == int32(-1) {
		v378 = v337
		goto L74
	} else {
		goto L91
	}
L91:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v68)+4)) = int64(4294967296)
	v358 = int32(1)
	v359 = int32(0)
	goto L85
L92:
	;
	v374 = *(*int32)(unsafe.Add(mBase, uint32(v370)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v282))) = v374
	v378 = v370
	goto L74
L93:
	;
	goto L43
L94:
	;
	goto L1
}
func F_sentinelCheckObjectivelyDown(m *base.Module, l0 int32) {
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
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v51 int64
	_ = v51
	var v52 int64
	_ = v52
	var v53 int64
	_ = v53
	var v54 int64
	_ = v54
	var v55 int64
	_ = v55
	var v56 int64
	_ = v56
	var v57 int64
	_ = v57
	var v59 int64
	_ = v59
	var v61 int64
	_ = v61
	var v63 int64
	_ = v63
	var v65 int64
	_ = v65
	var v67 int64
	_ = v67
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v89 int32
	_ = v89
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v109 int32
	_ = v109
	var v114 int32
	_ = v114
	var v118 int32
	_ = v118
	var v122 int32
	_ = v122
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v139 int32
	_ = v139
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v151 int32
	_ = v151
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v159 int32
	_ = v159
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v168 int32
	_ = v168
	var v170 int32
	_ = v170
	var v172 int64
	_ = v172
	var v173 int64
	_ = v173
	var v174 int64
	_ = v174
	var v175 int64
	_ = v175
	var v176 int64
	_ = v176
	var v177 int64
	_ = v177
	var v178 int64
	_ = v178
	var v180 int64
	_ = v180
	var v182 int64
	_ = v182
	var v184 int64
	_ = v184
	var v186 int64
	_ = v186
	var v188 int64
	_ = v188
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v195 int32
	_ = v195
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v202 int32
	_ = v202
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v210 int32
	_ = v210
	var v214 int32
	_ = v214
	var v216 int32
	_ = v216
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v230 int32
	_ = v230
	var v235 int32
	_ = v235
	var v239 int32
	_ = v239
	var v243 int32
	_ = v243
	var v250 int32
	_ = v250
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v255 int32
	_ = v255
	var v265 int32
	_ = v265
	var v266 int32
	_ = v266
	var v270 int64
	_ = v270
	var v274 int32
	_ = v274
	var v286 int32
	_ = v286
	var v287 int32
	_ = v287
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v10&int32(8) == int32(0) {
		v274 = v10
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v8 + int32(16)
	return
L2:
	;
	if v274&int32(16) == int32(0) {
		goto L1
	} else {
		goto L68
	}
L3:
	;
	v15 = int32(1)
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+144))
	v17 = F_dictGetIterator(m, v16)
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		goto L5
	} else {
		goto L6
	}
L4:
	;
	F_dictReleaseIterator(m, v17)
	mBase = m.M
	v253 = m.ExcPending
	if v253 != 0 {
		goto L5
	} else {
		goto L64
	}
L5:
	;
	return
L6:
	;
	v26 = v17 + int32(20)
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v17)+16))
	if v27 != 0 {
		goto L9
	} else {
		goto L10
	}
L7:
	;
	if v122 == int32(0) {
		v250 = v15
		goto L4
	} else {
		goto L33
	}
L8:
	;
	v33 = v26
	v34 = v30
	goto L11
L9:
	;
	v30 = int32(1)
	goto L8
L10:
	;
	v30 = int32(0)
	goto L8
L11:
	;
	switch v34 {
	case 0:
		goto L16
	default:
		goto L15
	}
L13:
	;
	v34 = int32(0)
	goto L11
L14:
	;
	goto L7
L15:
	;
	v114 = *(*int32)(unsafe.Add(mBase, uint32(v33)))
	*(*int32)(unsafe.Add(mBase, uint32(v17)+16)) = v114
	if v114 == int32(0) {
		goto L13
	} else {
		goto L32
	}
L16:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v17)+4))
	if v38 != int32(-1) {
		v77 = v38
		goto L17
	} else {
		goto L18
	}
L17:
	;
	v78 = int32(1)
	v79 = v77 + v78
	*(*int32)(unsafe.Add(mBase, uint32(v17)+4)) = v79
	v81 = int32(0)
	v84 = *(*int32)(unsafe.Add(mBase, uint32(v17)))
	v85 = *(*int32)(unsafe.Add(mBase, uint32(v17)+8))
	v89 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v84+v85+int32(26)))))
	if v89 == int32(255) {
		goto L26
	} else {
		goto L27
	}
L18:
	;
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v17)+8))
	if v42 != 0 {
		v77 = int32(-1)
		goto L17
	} else {
		goto L19
	}
L19:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v17)))
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v17)+12))
	if v44 == int32(0) {
		goto L21
	} else {
		goto L22
	}
L20:
	;
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v70)+20))
	if v71 != int32(-1) {
		goto L23
	} else {
		goto L24
	}
L21:
	;
	v51 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v43)+16)))
	v52 = int64(*(*int8)(unsafe.Add(mBase, uint32(v43)+27)))
	v53 = int64(*(*int32)(unsafe.Add(mBase, uint32(v43)+8)))
	v54 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v43)+12)))
	v55 = int64(*(*int8)(unsafe.Add(mBase, uint32(v43)+26)))
	v56 = int64(*(*int32)(unsafe.Add(mBase, uint32(v43)+4)))
	v57 = F_wangHash64(m, v56)
	mBase = m.M
	v59 = F_wangHash64(m, v55+v57)
	mBase = m.M
	v61 = F_wangHash64(m, v54+v59)
	mBase = m.M
	v63 = F_wangHash64(m, v53+v61)
	mBase = m.M
	v65 = F_wangHash64(m, v52+v63)
	mBase = m.M
	v67 = F_wangHash64(m, v51+v65)
	mBase = m.M
	*(*int64)(unsafe.Add(mBase, uint32(v17)+24)) = v67
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v17)))
	v70 = v69
	goto L20
L22:
	;
	v47 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v43)+24)))
	v49 = v47 + int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v43)+24)) = uint16(v49)
	v70 = v43
	goto L20
L23:
	;
	v77 = v71 + int32(-1)
	goto L17
L24:
	;
	v74 = *(*int32)(unsafe.Add(mBase, uint32(v17)+4))
	v77 = v74
	goto L17
L25:
	;
	v104 = int32(2)
	v109 = *(*int32)(unsafe.Add(mBase, uint32(v84+v102<<(uint(v104)%32)+int32(4))))
	v33 = v109 + v103<<(uint(v104)%32)
	v34 = int32(1)
	goto L11
L26:
	;
	v93 = v81
	goto L28
L27:
	;
	v93 = v78 << (uint(v89) % 32)
	goto L28
L28:
	;
	if v79 < v93 {
		v102 = v85
		v103 = v79
		goto L25
	} else {
		goto L29
	}
L29:
	;
	if v85 != 0 {
		v122 = v81
		goto L14
	} else {
		goto L30
	}
L30:
	;
	v95 = *(*int32)(unsafe.Add(mBase, uint32(v84)+20))
	if v95 == int32(-1) {
		v122 = v81
		goto L14
	} else {
		goto L31
	}
L31:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v17)+4)) = int64(4294967296)
	v102 = int32(1)
	v103 = int32(0)
	goto L25
L32:
	;
	v118 = *(*int32)(unsafe.Add(mBase, uint32(v114)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v26))) = v118
	v122 = v114
	goto L14
L33:
	;
	v130 = v122
	v131 = v15
	goto L34
L34:
	;
	v133 = *(*int32)(unsafe.Add(mBase, uint32(v130)+8))
	goto L36
L35:
	;
	v250 = v139
	goto L4
L36:
	;
	v134 = *(*int32)(unsafe.Add(mBase, uint32(v133)))
	v139 = int32(base.Ui32(v134)>>(uint(int32(5))%32))&int32(1) + v131
	v147 = v17 + int32(20)
	v148 = *(*int32)(unsafe.Add(mBase, uint32(v17)+16))
	if v148 != 0 {
		goto L39
	} else {
		goto L40
	}
L37:
	;
	if v243 != 0 {
		v130 = v243
		v131 = v139
		goto L34
	} else {
		goto L63
	}
L38:
	;
	v154 = v147
	v155 = v151
	goto L41
L39:
	;
	v151 = int32(1)
	goto L38
L40:
	;
	v151 = int32(0)
	goto L38
L41:
	;
	switch v155 {
	case 0:
		goto L46
	default:
		goto L45
	}
L43:
	;
	v155 = int32(0)
	goto L41
L44:
	;
	goto L37
L45:
	;
	v235 = *(*int32)(unsafe.Add(mBase, uint32(v154)))
	*(*int32)(unsafe.Add(mBase, uint32(v17)+16)) = v235
	if v235 == int32(0) {
		goto L43
	} else {
		goto L62
	}
L46:
	;
	v159 = *(*int32)(unsafe.Add(mBase, uint32(v17)+4))
	if v159 != int32(-1) {
		v198 = v159
		goto L47
	} else {
		goto L48
	}
L47:
	;
	v199 = int32(1)
	v200 = v198 + v199
	*(*int32)(unsafe.Add(mBase, uint32(v17)+4)) = v200
	v202 = int32(0)
	v205 = *(*int32)(unsafe.Add(mBase, uint32(v17)))
	v206 = *(*int32)(unsafe.Add(mBase, uint32(v17)+8))
	v210 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v205+v206+int32(26)))))
	if v210 == int32(255) {
		goto L56
	} else {
		goto L57
	}
L48:
	;
	v163 = *(*int32)(unsafe.Add(mBase, uint32(v17)+8))
	if v163 != 0 {
		v198 = int32(-1)
		goto L47
	} else {
		goto L49
	}
L49:
	;
	v164 = *(*int32)(unsafe.Add(mBase, uint32(v17)))
	v165 = *(*int32)(unsafe.Add(mBase, uint32(v17)+12))
	if v165 == int32(0) {
		goto L51
	} else {
		goto L52
	}
L50:
	;
	v192 = *(*int32)(unsafe.Add(mBase, uint32(v191)+20))
	if v192 != int32(-1) {
		goto L53
	} else {
		goto L54
	}
L51:
	;
	v172 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v164)+16)))
	v173 = int64(*(*int8)(unsafe.Add(mBase, uint32(v164)+27)))
	v174 = int64(*(*int32)(unsafe.Add(mBase, uint32(v164)+8)))
	v175 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v164)+12)))
	v176 = int64(*(*int8)(unsafe.Add(mBase, uint32(v164)+26)))
	v177 = int64(*(*int32)(unsafe.Add(mBase, uint32(v164)+4)))
	v178 = F_wangHash64(m, v177)
	mBase = m.M
	v180 = F_wangHash64(m, v176+v178)
	mBase = m.M
	v182 = F_wangHash64(m, v175+v180)
	mBase = m.M
	v184 = F_wangHash64(m, v174+v182)
	mBase = m.M
	v186 = F_wangHash64(m, v173+v184)
	mBase = m.M
	v188 = F_wangHash64(m, v172+v186)
	mBase = m.M
	*(*int64)(unsafe.Add(mBase, uint32(v17)+24)) = v188
	v190 = *(*int32)(unsafe.Add(mBase, uint32(v17)))
	v191 = v190
	goto L50
L52:
	;
	v168 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v164)+24)))
	v170 = v168 + int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v164)+24)) = uint16(v170)
	v191 = v164
	goto L50
L53:
	;
	v198 = v192 + int32(-1)
	goto L47
L54:
	;
	v195 = *(*int32)(unsafe.Add(mBase, uint32(v17)+4))
	v198 = v195
	goto L47
L55:
	;
	v225 = int32(2)
	v230 = *(*int32)(unsafe.Add(mBase, uint32(v205+v223<<(uint(v225)%32)+int32(4))))
	v154 = v230 + v224<<(uint(v225)%32)
	v155 = int32(1)
	goto L41
L56:
	;
	v214 = v202
	goto L58
L57:
	;
	v214 = v199 << (uint(v210) % 32)
	goto L58
L58:
	;
	if v200 < v214 {
		v223 = v206
		v224 = v200
		goto L55
	} else {
		goto L59
	}
L59:
	;
	if v206 != 0 {
		v243 = v202
		goto L44
	} else {
		goto L60
	}
L60:
	;
	v216 = *(*int32)(unsafe.Add(mBase, uint32(v205)+20))
	if v216 == int32(-1) {
		v243 = v202
		goto L44
	} else {
		goto L61
	}
L61:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v17)+4)) = int64(4294967296)
	v223 = int32(1)
	v224 = int32(0)
	goto L55
L62:
	;
	v239 = *(*int32)(unsafe.Add(mBase, uint32(v235)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v147))) = v239
	v243 = v235
	goto L44
L63:
	;
	goto L35
L64:
	;
	v254 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v255 = *(*int32)(unsafe.Add(mBase, uint32(l0)+152))
	if base.Ui32(v250) < base.Ui32(v255) {
		v274 = v254
		goto L2
	} else {
		goto L65
	}
L65:
	;
	if v254&int32(16) != 0 {
		goto L1
	} else {
		goto L66
	}
L66:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = v255
	*(*int32)(unsafe.Add(mBase, uint32(v8))) = v250
	F_sentinelEvent(m, int32(3), int32(_a_F_sentinelCheckObjectivelyDown_0), l0, int32(_a_F_sentinelCheckObjectivelyDown_1), v8)
	mBase = m.M
	v265 = m.ExcPending
	if v265 != 0 {
		goto L5
	} else {
		goto L67
	}
L67:
	;
	v266 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v266 | int32(16)
	v270 = F_mstime(m)
	mBase = m.M
	*(*int64)(unsafe.Add(mBase, uint32(l0)+64)) = v270
	goto L1
L68:
	;
	F_sentinelEvent(m, int32(3), int32(_a_F_sentinelCheckObjectivelyDown_2), l0, int32(_a_F_sentinelCheckObjectivelyDown_3), int32(0))
	mBase = m.M
	v286 = m.ExcPending
	if v286 != 0 {
		goto L5
	} else {
		goto L69
	}
L69:
	;
	v287 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v287 & int32(-17)
	goto L1
}
func F_sentinelCollectTerminatedScripts(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v23 int32
	_ = v23
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v76 int32
	_ = v76
	var v82 int32
	_ = v82
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v94 int32
	_ = v94
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v102 int32
	_ = v102
	var v108 int32
	_ = v108
	var v119 int32
	_ = v119
	var v126 int32
	_ = v126
	var v131 int32
	_ = v131
	var v134 int32
	_ = v134
	var v136 int32
	_ = v136
	var v140 int64
	_ = v140
	var v142 int64
	_ = v142
	var v143 int32
	_ = v143
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v164 int64
	_ = v164
	var v166 int64
	_ = v166
	var v168 int32
	_ = v168
	var v170 int32
	_ = v170
	var v174 int32
	_ = v174
	var v179 int64
	_ = v179
	var v184 int32
	_ = v184
	var v189 int64
	_ = v189
	var v191 int64
	_ = v191
	var v205 int64
	_ = v205
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v223 int32
	_ = v223
	var v225 int32
	_ = v225
	var v227 int32
	_ = v227
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v246 int32
	_ = v246
	var v250 int32
	_ = v250
	var v254 int32
	_ = v254
	var v260 int32
	_ = v260
	var v262 int32
	_ = v262
	var v271 int32
	_ = v271
	var v273 int32
	_ = v273
	var v290 int32
	_ = v290
	var v291 int32
	_ = v291
	v9 = m.G0
	v11 = v9 - int32(64)
	m.G0 = v11
	v18 = F___syscall_wait4(m, int32(-1), v11+int32(52), int32(1), int32(0))
	mBase = m.M
	v19 = F___syscall_ret(m, v18)
	mBase = m.M
	goto L2
L1:
	;
	m.G0 = v11 + int32(64)
	return
L2:
	;
	if v19 < int32(1) {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v23 = v19
	goto L4
L4:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+32)) = v23
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v11)+52))
	v34 = int32(255)
	v35 = int32(base.Ui32(v31)>>(uint(int32(8))%32)) & v34
	*(*int32)(unsafe.Add(mBase, uint32(v11)+36)) = v35
	if base.Ui32(v31&int32(65535)+int32(-1)) < base.Ui32(v34) {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	goto L1
L6:
	;
	v46 = v31 & int32(127)
	goto L8
L7:
	;
	v46 = int32(0)
	goto L8
L8:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+40)) = v46
	v48 = int32(0)
	F_sentinelEvent(m, v48, int32(_a_F_sentinelCollectTerminatedScripts_0), v48, int32(_a_F_sentinelCollectTerminatedScripts_1), v11+int32(32))
	mBase = m.M
	v55 = m.ExcPending
	if v55 != 0 {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	return
L10:
	;
	v56 = int32(0)
	v57 = *(*int32)(unsafe.Add(mBase, _c_F_sentinelCollectTerminatedScripts[0]))
	v59 = v11 + int32(56)
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v57)))
	*(*int32)(unsafe.Add(mBase, uint32(v59)+4)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v59))) = v60
	goto L11
L11:
	;
	v65 = v11 + int32(56)
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v65)))
	if v67 == int32(0) {
		goto L16
	} else {
		goto L17
	}
L12:
	;
	v290 = F___syscall_wait4(m, int32(-1), v11+int32(52), int32(1), int32(0))
	mBase = m.M
	v291 = F___syscall_ret(m, v290)
	mBase = m.M
	goto L59
L13:
	;
	if v35 == int32(1) {
		goto L32
	} else {
		goto L33
	}
L14:
	;
	v119 = *(*int32)(unsafe.Add(mBase, _c_F_sentinelCollectTerminatedScripts[1]))
	if int32(3) < v119 {
		goto L12
	} else {
		goto L28
	}
L15:
	;
	if v67 == int32(0) {
		goto L14
	} else {
		goto L18
	}
L16:
	;
	goto L15
L17:
	;
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v65)+4))
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v67+base.B2i32(v70 == int32(0))<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v65))) = v76
	goto L16
L18:
	;
	v82 = v67
	goto L19
L19:
	;
	v88 = *(*int32)(unsafe.Add(mBase, uint32(v82)+8))
	v89 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v88))))
	if v89&int32(1) == int32(0) {
		goto L21
	} else {
		goto L22
	}
L20:
	;
	goto L14
L21:
	;
	v97 = v11 + int32(56)
	v99 = *(*int32)(unsafe.Add(mBase, uint32(v97)))
	if v99 == int32(0) {
		goto L25
	} else {
		goto L26
	}
L22:
	;
	v94 = *(*int32)(unsafe.Add(mBase, uint32(v88)+24))
	if v94 == v23 {
		goto L13
	} else {
		goto L23
	}
L23:
	;
	goto L21
L24:
	;
	if v99 != 0 {
		v82 = v99
		goto L19
	} else {
		goto L27
	}
L25:
	;
	goto L24
L26:
	;
	v102 = *(*int32)(unsafe.Add(mBase, uint32(v97)+4))
	v108 = *(*int32)(unsafe.Add(mBase, uint32(v99+base.B2i32(v102 == int32(0))<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v97))) = v108
	goto L25
L27:
	;
	goto L20
L28:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11))) = v23
	F__serverLog(m, int32(3), int32(_a_F_sentinelCollectTerminatedScripts_2), v11)
	mBase = m.M
	v126 = m.ExcPending
	if v126 != 0 {
		goto L9
	} else {
		goto L29
	}
L29:
	;
	goto L12
L30:
	;
	v271 = int32(0)
	v273 = *(*int32)(unsafe.Add(mBase, _c_F_sentinelCollectTerminatedScripts[2]))
	*(*int32)(unsafe.Add(mBase, _c_F_sentinelCollectTerminatedScripts[2])) = v273 + int32(-1)
	goto L12
L31:
	;
	if v46|v35 == int32(0) {
		goto L47
	} else {
		goto L48
	}
L32:
	;
	v131 = *(*int32)(unsafe.Add(mBase, uint32(v88)+4))
	if v131 == int32(10) {
		goto L31
	} else {
		goto L35
	}
L33:
	;
	if v46 == int32(0) {
		goto L31
	} else {
		goto L34
	}
L34:
	;
	goto L32
L35:
	;
	v134 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v88)+24)) = v134
	v136 = *(*int32)(unsafe.Add(mBase, uint32(v88)))
	*(*int32)(unsafe.Add(mBase, uint32(v88))) = v136 & int32(-2)
	v140 = F_mstime(m)
	mBase = m.M
	v142 = *(*int64)(unsafe.Add(mBase, _c_F_sentinelCollectTerminatedScripts[3]))
	v143 = *(*int32)(unsafe.Add(mBase, uint32(v88)+4))
	if v143 < int32(2) {
		v205 = v142
		goto L36
	} else {
		goto L37
	}
L36:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v88)+16)) = v205 + v140
	goto L30
L37:
	;
	if v143&int32(7) == int32(1) {
		v174 = v143
		v179 = v142
		goto L38
	} else {
		goto L39
	}
L38:
	;
	if base.Ui32(v143+int32(-2)) < base.Ui32(int32(7)) {
		v205 = v179
		goto L36
	} else {
		goto L43
	}
L39:
	;
	v158 = int32(0)
	v159 = v143
	v164 = v142
	goto L40
L40:
	;
	v166 = v164 << (uint(int64(1)) % 64)
	v168 = v159 + int32(-1)
	v170 = v158 + int32(1)
	if v170 != (v143+int32(-1))&int32(7) {
		v158 = v170
		v159 = v168
		v164 = v166
		goto L40
	} else {
		goto L42
	}
L41:
	;
	v174 = v168
	v179 = v166
	goto L38
L42:
	;
	goto L41
L43:
	;
	v184 = v174
	v189 = v179
	goto L44
L44:
	;
	v191 = v189 << (uint(int64(8)) % 64)
	if base.Ui32(v184+int32(-10)) < base.Ui32(int32(-3)) {
		v184 = v184 + int32(-8)
		v189 = v191
		goto L44
	} else {
		goto L46
	}
L45:
	;
	v205 = v191
	goto L36
L46:
	;
	goto L45
L47:
	;
	v225 = int32(0)
	v227 = *(*int32)(unsafe.Add(mBase, _c_F_sentinelCollectTerminatedScripts[0]))
	F_listDelNode(m, v227, v82)
	mBase = m.M
	v229 = m.ExcPending
	if v229 != 0 {
		goto L9
	} else {
		goto L50
	}
L48:
	;
	v211 = *(*int32)(unsafe.Add(mBase, uint32(v88)+8))
	v212 = *(*int32)(unsafe.Add(mBase, uint32(v211)))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+24)) = v35
	*(*int32)(unsafe.Add(mBase, uint32(v11)+20)) = v46
	*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = v212
	F_sentinelEvent(m, int32(3), int32(_a_F_sentinelCollectTerminatedScripts_3), int32(0), int32(_a_F_sentinelCollectTerminatedScripts_4), v11+int32(16))
	mBase = m.M
	v223 = m.ExcPending
	if v223 != 0 {
		goto L9
	} else {
		goto L49
	}
L49:
	;
	goto L47
L50:
	;
	v230 = *(*int32)(unsafe.Add(mBase, uint32(v88)+8))
	v231 = *(*int32)(unsafe.Add(mBase, uint32(v230)))
	if v231 == int32(0) {
		v254 = v230
		goto L51
	} else {
		goto L52
	}
L51:
	;
	F_valkey_free(m, v254)
	mBase = m.M
	v260 = m.ExcPending
	if v260 != 0 {
		goto L9
	} else {
		goto L57
	}
L52:
	;
	v235 = v225
	v236 = v231
	goto L53
L53:
	;
	F_sdsfree(m, v236)
	mBase = m.M
	v243 = m.ExcPending
	if v243 != 0 {
		goto L9
	} else {
		goto L55
	}
L54:
	;
	v254 = v244
	goto L51
L55:
	;
	v244 = *(*int32)(unsafe.Add(mBase, uint32(v88)+8))
	v246 = v235 + int32(1)
	v250 = *(*int32)(unsafe.Add(mBase, uint32(v244+v246<<(uint(int32(2))%32))))
	if v250 != 0 {
		v235 = v246
		v236 = v250
		goto L53
	} else {
		goto L56
	}
L56:
	;
	goto L54
L57:
	;
	F_valkey_free(m, v88)
	mBase = m.M
	v262 = m.ExcPending
	if v262 != 0 {
		goto L9
	} else {
		goto L58
	}
L58:
	;
	goto L30
L59:
	;
	if int32(0) < v291 {
		v23 = v291
		goto L4
	} else {
		goto L60
	}
L60:
	;
	goto L5
}
func F_sentinelDiscardReplyCallback(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+212))
	if v4 == int32(0) {
	} else {
		v7 = *(*int32)(unsafe.Add(mBase, uint32(v4)+8))
		*(*int32)(unsafe.Add(mBase, uint32(v4)+8)) = v7 + int32(-1)
	}
	return
}
func F_sentinelDropConnections(m *base.Module) int32 {
	mBase := m.M
	_ = mBase
	var v1 int32
	_ = v1
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v33 int32
	_ = v33
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
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
	var v62 int64
	_ = v62
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v84 int32
	_ = v84
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v104 int32
	_ = v104
	var v109 int32
	_ = v109
	var v113 int32
	_ = v113
	var v117 int32
	_ = v117
	var v123 int32
	_ = v123
	var v125 int32
	_ = v125
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v150 int32
	_ = v150
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v158 int32
	_ = v158
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v167 int32
	_ = v167
	var v169 int32
	_ = v169
	var v171 int64
	_ = v171
	var v172 int64
	_ = v172
	var v173 int64
	_ = v173
	var v174 int64
	_ = v174
	var v175 int64
	_ = v175
	var v176 int64
	_ = v176
	var v177 int64
	_ = v177
	var v179 int64
	_ = v179
	var v181 int64
	_ = v181
	var v183 int64
	_ = v183
	var v185 int64
	_ = v185
	var v187 int64
	_ = v187
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v194 int32
	_ = v194
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v201 int32
	_ = v201
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v209 int32
	_ = v209
	var v213 int32
	_ = v213
	var v215 int32
	_ = v215
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v229 int32
	_ = v229
	var v234 int32
	_ = v234
	var v238 int32
	_ = v238
	var v242 int32
	_ = v242
	var v248 int32
	_ = v248
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v251 int32
	_ = v251
	var v254 int32
	_ = v254
	var v258 int32
	_ = v258
	var v265 int32
	_ = v265
	var v266 int32
	_ = v266
	var v267 int32
	_ = v267
	var v268 int32
	_ = v268
	var v273 int32
	_ = v273
	var v282 int32
	_ = v282
	var v286 int32
	_ = v286
	var v294 int32
	_ = v294
	var v295 int32
	_ = v295
	var v298 int32
	_ = v298
	var v301 int32
	_ = v301
	var v302 int32
	_ = v302
	var v306 int32
	_ = v306
	var v310 int32
	_ = v310
	var v311 int32
	_ = v311
	var v312 int32
	_ = v312
	var v315 int32
	_ = v315
	var v317 int32
	_ = v317
	var v319 int64
	_ = v319
	var v320 int64
	_ = v320
	var v321 int64
	_ = v321
	var v322 int64
	_ = v322
	var v323 int64
	_ = v323
	var v324 int64
	_ = v324
	var v325 int64
	_ = v325
	var v327 int64
	_ = v327
	var v329 int64
	_ = v329
	var v331 int64
	_ = v331
	var v333 int64
	_ = v333
	var v335 int64
	_ = v335
	var v337 int32
	_ = v337
	var v338 int32
	_ = v338
	var v339 int32
	_ = v339
	var v342 int32
	_ = v342
	var v345 int32
	_ = v345
	var v346 int32
	_ = v346
	var v347 int32
	_ = v347
	var v349 int32
	_ = v349
	var v352 int32
	_ = v352
	var v353 int32
	_ = v353
	var v357 int32
	_ = v357
	var v361 int32
	_ = v361
	var v363 int32
	_ = v363
	var v370 int32
	_ = v370
	var v371 int32
	_ = v371
	var v372 int32
	_ = v372
	var v377 int32
	_ = v377
	var v382 int32
	_ = v382
	var v386 int32
	_ = v386
	var v390 int32
	_ = v390
	var v394 int32
	_ = v394
	var v401 int32
	_ = v401
	v1 = int32(0)
	v9 = *(*int32)(unsafe.Add(mBase, _c_F_sentinelDropConnections[0]))
	v10 = F_dictGetIterator(m, v9)
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	F_dictReleaseIterator(m, v10)
	mBase = m.M
	v401 = m.ExcPending
	if v401 != 0 {
		goto L2
	} else {
		goto L105
	}
L2:
	;
	return int32(0)
L3:
	;
	v21 = v10 + int32(20)
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v10)+16))
	if v22 != 0 {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	if v117 == int32(0) {
		v394 = v1
		goto L1
	} else {
		goto L30
	}
L5:
	;
	v28 = v21
	v29 = v25
	goto L8
L6:
	;
	v25 = int32(1)
	goto L5
L7:
	;
	v25 = int32(0)
	goto L5
L8:
	;
	switch v29 {
	case 0:
		goto L13
	default:
		goto L12
	}
L10:
	;
	v29 = int32(0)
	goto L8
L11:
	;
	goto L4
L12:
	;
	v109 = *(*int32)(unsafe.Add(mBase, uint32(v28)))
	*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = v109
	if v109 == int32(0) {
		goto L10
	} else {
		goto L29
	}
L13:
	;
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v10)+4))
	if v33 != int32(-1) {
		v72 = v33
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v73 = int32(1)
	v74 = v72 + v73
	*(*int32)(unsafe.Add(mBase, uint32(v10)+4)) = v74
	v76 = int32(0)
	v79 = *(*int32)(unsafe.Add(mBase, uint32(v10)))
	v80 = *(*int32)(unsafe.Add(mBase, uint32(v10)+8))
	v84 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v79+v80+int32(26)))))
	if v84 == int32(255) {
		goto L23
	} else {
		goto L24
	}
L15:
	;
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v10)+8))
	if v37 != 0 {
		v72 = int32(-1)
		goto L14
	} else {
		goto L16
	}
L16:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v10)))
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v10)+12))
	if v39 == int32(0) {
		goto L18
	} else {
		goto L19
	}
L17:
	;
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v65)+20))
	if v66 != int32(-1) {
		goto L20
	} else {
		goto L21
	}
L18:
	;
	v46 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v38)+16)))
	v47 = int64(*(*int8)(unsafe.Add(mBase, uint32(v38)+27)))
	v48 = int64(*(*int32)(unsafe.Add(mBase, uint32(v38)+8)))
	v49 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v38)+12)))
	v50 = int64(*(*int8)(unsafe.Add(mBase, uint32(v38)+26)))
	v51 = int64(*(*int32)(unsafe.Add(mBase, uint32(v38)+4)))
	v52 = F_wangHash64(m, v51)
	mBase = m.M
	v54 = F_wangHash64(m, v50+v52)
	mBase = m.M
	v56 = F_wangHash64(m, v49+v54)
	mBase = m.M
	v58 = F_wangHash64(m, v48+v56)
	mBase = m.M
	v60 = F_wangHash64(m, v47+v58)
	mBase = m.M
	v62 = F_wangHash64(m, v46+v60)
	mBase = m.M
	*(*int64)(unsafe.Add(mBase, uint32(v10)+24)) = v62
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v10)))
	v65 = v64
	goto L17
L19:
	;
	v42 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v38)+24)))
	v44 = v42 + int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v38)+24)) = uint16(v44)
	v65 = v38
	goto L17
L20:
	;
	v72 = v66 + int32(-1)
	goto L14
L21:
	;
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v10)+4))
	v72 = v69
	goto L14
L22:
	;
	v99 = int32(2)
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v79+v97<<(uint(v99)%32)+int32(4))))
	v28 = v104 + v98<<(uint(v99)%32)
	v29 = int32(1)
	goto L8
L23:
	;
	v88 = v76
	goto L25
L24:
	;
	v88 = v73 << (uint(v84) % 32)
	goto L25
L25:
	;
	if v74 < v88 {
		v97 = v80
		v98 = v74
		goto L22
	} else {
		goto L26
	}
L26:
	;
	if v80 != 0 {
		v117 = v76
		goto L11
	} else {
		goto L27
	}
L27:
	;
	v90 = *(*int32)(unsafe.Add(mBase, uint32(v79)+20))
	if v90 == int32(-1) {
		v117 = v76
		goto L11
	} else {
		goto L28
	}
L28:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v10)+4)) = int64(4294967296)
	v97 = int32(1)
	v98 = int32(0)
	goto L22
L29:
	;
	v113 = *(*int32)(unsafe.Add(mBase, uint32(v109)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v21))) = v113
	v117 = v109
	goto L11
L30:
	;
	v123 = v1
	v125 = v117
	goto L31
L31:
	;
	v129 = *(*int32)(unsafe.Add(mBase, uint32(v125)+8))
	goto L33
L32:
	;
	v394 = v133
	goto L1
L33:
	;
	v130 = *(*int32)(unsafe.Add(mBase, uint32(v129)+144))
	v131 = F_dictGetIterator(m, v130)
	mBase = m.M
	v132 = m.ExcPending
	if v132 != 0 {
		goto L2
	} else {
		goto L34
	}
L34:
	;
	v133 = v123
	goto L36
L35:
	;
	F_dictReleaseIterator(m, v131)
	mBase = m.M
	v286 = m.ExcPending
	if v286 != 0 {
		goto L2
	} else {
		goto L77
	}
L36:
	;
	v146 = v131 + int32(20)
	v147 = *(*int32)(unsafe.Add(mBase, uint32(v131)+16))
	if v147 != 0 {
		goto L40
	} else {
		goto L41
	}
L38:
	;
	if v242 == int32(0) {
		goto L35
	} else {
		goto L64
	}
L39:
	;
	v153 = v146
	v154 = v150
	goto L42
L40:
	;
	v150 = int32(1)
	goto L39
L41:
	;
	v150 = int32(0)
	goto L39
L42:
	;
	switch v154 {
	case 0:
		goto L47
	default:
		goto L46
	}
L44:
	;
	v154 = int32(0)
	goto L42
L45:
	;
	goto L38
L46:
	;
	v234 = *(*int32)(unsafe.Add(mBase, uint32(v153)))
	*(*int32)(unsafe.Add(mBase, uint32(v131)+16)) = v234
	if v234 == int32(0) {
		goto L44
	} else {
		goto L63
	}
L47:
	;
	v158 = *(*int32)(unsafe.Add(mBase, uint32(v131)+4))
	if v158 != int32(-1) {
		v197 = v158
		goto L48
	} else {
		goto L49
	}
L48:
	;
	v198 = int32(1)
	v199 = v197 + v198
	*(*int32)(unsafe.Add(mBase, uint32(v131)+4)) = v199
	v201 = int32(0)
	v204 = *(*int32)(unsafe.Add(mBase, uint32(v131)))
	v205 = *(*int32)(unsafe.Add(mBase, uint32(v131)+8))
	v209 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v204+v205+int32(26)))))
	if v209 == int32(255) {
		goto L57
	} else {
		goto L58
	}
L49:
	;
	v162 = *(*int32)(unsafe.Add(mBase, uint32(v131)+8))
	if v162 != 0 {
		v197 = int32(-1)
		goto L48
	} else {
		goto L50
	}
L50:
	;
	v163 = *(*int32)(unsafe.Add(mBase, uint32(v131)))
	v164 = *(*int32)(unsafe.Add(mBase, uint32(v131)+12))
	if v164 == int32(0) {
		goto L52
	} else {
		goto L53
	}
L51:
	;
	v191 = *(*int32)(unsafe.Add(mBase, uint32(v190)+20))
	if v191 != int32(-1) {
		goto L54
	} else {
		goto L55
	}
L52:
	;
	v171 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v163)+16)))
	v172 = int64(*(*int8)(unsafe.Add(mBase, uint32(v163)+27)))
	v173 = int64(*(*int32)(unsafe.Add(mBase, uint32(v163)+8)))
	v174 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v163)+12)))
	v175 = int64(*(*int8)(unsafe.Add(mBase, uint32(v163)+26)))
	v176 = int64(*(*int32)(unsafe.Add(mBase, uint32(v163)+4)))
	v177 = F_wangHash64(m, v176)
	mBase = m.M
	v179 = F_wangHash64(m, v175+v177)
	mBase = m.M
	v181 = F_wangHash64(m, v174+v179)
	mBase = m.M
	v183 = F_wangHash64(m, v173+v181)
	mBase = m.M
	v185 = F_wangHash64(m, v172+v183)
	mBase = m.M
	v187 = F_wangHash64(m, v171+v185)
	mBase = m.M
	*(*int64)(unsafe.Add(mBase, uint32(v131)+24)) = v187
	v189 = *(*int32)(unsafe.Add(mBase, uint32(v131)))
	v190 = v189
	goto L51
L53:
	;
	v167 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v163)+24)))
	v169 = v167 + int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v163)+24)) = uint16(v169)
	v190 = v163
	goto L51
L54:
	;
	v197 = v191 + int32(-1)
	goto L48
L55:
	;
	v194 = *(*int32)(unsafe.Add(mBase, uint32(v131)+4))
	v197 = v194
	goto L48
L56:
	;
	v224 = int32(2)
	v229 = *(*int32)(unsafe.Add(mBase, uint32(v204+v222<<(uint(v224)%32)+int32(4))))
	v153 = v229 + v223<<(uint(v224)%32)
	v154 = int32(1)
	goto L42
L57:
	;
	v213 = v201
	goto L59
L58:
	;
	v213 = v198 << (uint(v209) % 32)
	goto L59
L59:
	;
	if v199 < v213 {
		v222 = v205
		v223 = v199
		goto L56
	} else {
		goto L60
	}
L60:
	;
	if v205 != 0 {
		v242 = v201
		goto L45
	} else {
		goto L61
	}
L61:
	;
	v215 = *(*int32)(unsafe.Add(mBase, uint32(v204)+20))
	if v215 == int32(-1) {
		v242 = v201
		goto L45
	} else {
		goto L62
	}
L62:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v131)+4)) = int64(4294967296)
	v222 = int32(1)
	v223 = int32(0)
	goto L56
L63:
	;
	v238 = *(*int32)(unsafe.Add(mBase, uint32(v234)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v146))) = v238
	v242 = v234
	goto L45
L64:
	;
	v248 = *(*int32)(unsafe.Add(mBase, uint32(v242)+8))
	goto L65
L65:
	;
	v249 = *(*int32)(unsafe.Add(mBase, uint32(v248)+28))
	v250 = *(*int32)(unsafe.Add(mBase, uint32(v249)+4))
	if v250 != 0 {
		goto L36
	} else {
		goto L66
	}
L66:
	;
	v251 = *(*int32)(unsafe.Add(mBase, uint32(v249)+16))
	if v251 == int32(0) {
		v267 = v249
		goto L67
	} else {
		goto L68
	}
L67:
	;
	v268 = *(*int32)(unsafe.Add(mBase, uint32(v267)+12))
	if v268 == int32(0) {
		goto L72
	} else {
		goto L73
	}
L68:
	;
	v254 = *(*int32)(unsafe.Add(mBase, uint32(v249)+12))
	if v254 != v251 {
		goto L69
	} else {
		goto L70
	}
L69:
	;
	v258 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v249)+16)) = v258
	*(*int32)(unsafe.Add(mBase, uint32(v251)+212)) = v258
	*(*int32)(unsafe.Add(mBase, uint32(v249)+4)) = int32(1)
	F_valkeyAsyncFree(m, v251)
	mBase = m.M
	v265 = m.ExcPending
	if v265 != 0 {
		goto L2
	} else {
		goto L71
	}
L70:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v249)+8)) = int64(0)
	goto L69
L71:
	;
	v266 = *(*int32)(unsafe.Add(mBase, uint32(v248)+28))
	v267 = v266
	goto L67
L72:
	;
	v133 = v133 + int32(1)
	goto L36
L73:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v267)+8)) = int64(0)
	v273 = *(*int32)(unsafe.Add(mBase, uint32(v267)+16))
	if v273 != v268 {
		goto L74
	} else {
		goto L75
	}
L74:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v268)+212)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v267)+4)) = int32(1)
	F_valkeyAsyncFree(m, v268)
	mBase = m.M
	v282 = m.ExcPending
	if v282 != 0 {
		goto L2
	} else {
		goto L76
	}
L75:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v267)+16)) = int32(0)
	goto L74
L76:
	;
	goto L72
L77:
	;
	v294 = v10 + int32(20)
	v295 = *(*int32)(unsafe.Add(mBase, uint32(v10)+16))
	if v295 != 0 {
		goto L80
	} else {
		goto L81
	}
L78:
	;
	if v390 != 0 {
		v123 = v133
		v125 = v390
		goto L31
	} else {
		goto L104
	}
L79:
	;
	v301 = v294
	v302 = v298
	goto L82
L80:
	;
	v298 = int32(1)
	goto L79
L81:
	;
	v298 = int32(0)
	goto L79
L82:
	;
	switch v302 {
	case 0:
		goto L87
	default:
		goto L86
	}
L84:
	;
	v302 = int32(0)
	goto L82
L85:
	;
	goto L78
L86:
	;
	v382 = *(*int32)(unsafe.Add(mBase, uint32(v301)))
	*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = v382
	if v382 == int32(0) {
		goto L84
	} else {
		goto L103
	}
L87:
	;
	v306 = *(*int32)(unsafe.Add(mBase, uint32(v10)+4))
	if v306 != int32(-1) {
		v345 = v306
		goto L88
	} else {
		goto L89
	}
L88:
	;
	v346 = int32(1)
	v347 = v345 + v346
	*(*int32)(unsafe.Add(mBase, uint32(v10)+4)) = v347
	v349 = int32(0)
	v352 = *(*int32)(unsafe.Add(mBase, uint32(v10)))
	v353 = *(*int32)(unsafe.Add(mBase, uint32(v10)+8))
	v357 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v352+v353+int32(26)))))
	if v357 == int32(255) {
		goto L97
	} else {
		goto L98
	}
L89:
	;
	v310 = *(*int32)(unsafe.Add(mBase, uint32(v10)+8))
	if v310 != 0 {
		v345 = int32(-1)
		goto L88
	} else {
		goto L90
	}
L90:
	;
	v311 = *(*int32)(unsafe.Add(mBase, uint32(v10)))
	v312 = *(*int32)(unsafe.Add(mBase, uint32(v10)+12))
	if v312 == int32(0) {
		goto L92
	} else {
		goto L93
	}
L91:
	;
	v339 = *(*int32)(unsafe.Add(mBase, uint32(v338)+20))
	if v339 != int32(-1) {
		goto L94
	} else {
		goto L95
	}
L92:
	;
	v319 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v311)+16)))
	v320 = int64(*(*int8)(unsafe.Add(mBase, uint32(v311)+27)))
	v321 = int64(*(*int32)(unsafe.Add(mBase, uint32(v311)+8)))
	v322 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v311)+12)))
	v323 = int64(*(*int8)(unsafe.Add(mBase, uint32(v311)+26)))
	v324 = int64(*(*int32)(unsafe.Add(mBase, uint32(v311)+4)))
	v325 = F_wangHash64(m, v324)
	mBase = m.M
	v327 = F_wangHash64(m, v323+v325)
	mBase = m.M
	v329 = F_wangHash64(m, v322+v327)
	mBase = m.M
	v331 = F_wangHash64(m, v321+v329)
	mBase = m.M
	v333 = F_wangHash64(m, v320+v331)
	mBase = m.M
	v335 = F_wangHash64(m, v319+v333)
	mBase = m.M
	*(*int64)(unsafe.Add(mBase, uint32(v10)+24)) = v335
	v337 = *(*int32)(unsafe.Add(mBase, uint32(v10)))
	v338 = v337
	goto L91
L93:
	;
	v315 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v311)+24)))
	v317 = v315 + int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v311)+24)) = uint16(v317)
	v338 = v311
	goto L91
L94:
	;
	v345 = v339 + int32(-1)
	goto L88
L95:
	;
	v342 = *(*int32)(unsafe.Add(mBase, uint32(v10)+4))
	v345 = v342
	goto L88
L96:
	;
	v372 = int32(2)
	v377 = *(*int32)(unsafe.Add(mBase, uint32(v352+v370<<(uint(v372)%32)+int32(4))))
	v301 = v377 + v371<<(uint(v372)%32)
	v302 = int32(1)
	goto L82
L97:
	;
	v361 = v349
	goto L99
L98:
	;
	v361 = v346 << (uint(v357) % 32)
	goto L99
L99:
	;
	if v347 < v361 {
		v370 = v353
		v371 = v347
		goto L96
	} else {
		goto L100
	}
L100:
	;
	if v353 != 0 {
		v390 = v349
		goto L85
	} else {
		goto L101
	}
L101:
	;
	v363 = *(*int32)(unsafe.Add(mBase, uint32(v352)+20))
	if v363 == int32(-1) {
		v390 = v349
		goto L85
	} else {
		goto L102
	}
L102:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v10)+4)) = int64(4294967296)
	v370 = int32(1)
	v371 = int32(0)
	goto L96
L103:
	;
	v386 = *(*int32)(unsafe.Add(mBase, uint32(v382)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v294))) = v386
	v390 = v382
	goto L85
L104:
	;
	goto L32
L105:
	;
	return v394
}
func F_sentinelFailoverDetectEnd(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v6 int64
	_ = v6
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v13 int64
	_ = v13
	var v14 int64
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v37 int32
	_ = v37
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v50 int64
	_ = v50
	var v51 int64
	_ = v51
	var v52 int64
	_ = v52
	var v53 int64
	_ = v53
	var v54 int64
	_ = v54
	var v55 int64
	_ = v55
	var v56 int64
	_ = v56
	var v58 int64
	_ = v58
	var v60 int64
	_ = v60
	var v62 int64
	_ = v62
	var v64 int64
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
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v88 int32
	_ = v88
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v108 int32
	_ = v108
	var v113 int32
	_ = v113
	var v117 int32
	_ = v117
	var v121 int32
	_ = v121
	var v130 int32
	_ = v130
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v139 int32
	_ = v139
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v151 int32
	_ = v151
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v159 int32
	_ = v159
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v168 int32
	_ = v168
	var v170 int32
	_ = v170
	var v172 int64
	_ = v172
	var v173 int64
	_ = v173
	var v174 int64
	_ = v174
	var v175 int64
	_ = v175
	var v176 int64
	_ = v176
	var v177 int64
	_ = v177
	var v178 int64
	_ = v178
	var v180 int64
	_ = v180
	var v182 int64
	_ = v182
	var v184 int64
	_ = v184
	var v186 int64
	_ = v186
	var v188 int64
	_ = v188
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v195 int32
	_ = v195
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v202 int32
	_ = v202
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v210 int32
	_ = v210
	var v214 int32
	_ = v214
	var v216 int32
	_ = v216
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v230 int32
	_ = v230
	var v235 int32
	_ = v235
	var v239 int32
	_ = v239
	var v243 int32
	_ = v243
	var v248 int32
	_ = v248
	var v249 int64
	_ = v249
	var v252 int32
	_ = v252
	var v253 int64
	_ = v253
	var v265 int32
	_ = v265
	var v271 int32
	_ = v271
	var v274 int64
	_ = v274
	var v276 int32
	_ = v276
	var v277 int32
	_ = v277
	var v278 int32
	_ = v278
	var v286 int32
	_ = v286
	var v287 int32
	_ = v287
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
	var v303 int32
	_ = v303
	var v304 int32
	_ = v304
	var v307 int32
	_ = v307
	var v309 int32
	_ = v309
	var v311 int64
	_ = v311
	var v312 int64
	_ = v312
	var v313 int64
	_ = v313
	var v314 int64
	_ = v314
	var v315 int64
	_ = v315
	var v316 int64
	_ = v316
	var v317 int64
	_ = v317
	var v319 int64
	_ = v319
	var v321 int64
	_ = v321
	var v323 int64
	_ = v323
	var v325 int64
	_ = v325
	var v327 int64
	_ = v327
	var v329 int32
	_ = v329
	var v330 int32
	_ = v330
	var v331 int32
	_ = v331
	var v334 int32
	_ = v334
	var v337 int32
	_ = v337
	var v338 int32
	_ = v338
	var v339 int32
	_ = v339
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
	var v355 int32
	_ = v355
	var v362 int32
	_ = v362
	var v363 int32
	_ = v363
	var v364 int32
	_ = v364
	var v369 int32
	_ = v369
	var v374 int32
	_ = v374
	var v378 int32
	_ = v378
	var v382 int32
	_ = v382
	var v390 int32
	_ = v390
	var v393 int32
	_ = v393
	var v394 int32
	_ = v394
	var v397 int32
	_ = v397
	var v398 int32
	_ = v398
	var v399 int32
	_ = v399
	var v400 int32
	_ = v400
	var v401 int32
	_ = v401
	var v402 int32
	_ = v402
	var v408 int32
	_ = v408
	var v409 int32
	_ = v409
	var v420 int32
	_ = v420
	var v421 int32
	_ = v421
	var v424 int32
	_ = v424
	var v427 int32
	_ = v427
	var v428 int32
	_ = v428
	var v432 int32
	_ = v432
	var v436 int32
	_ = v436
	var v437 int32
	_ = v437
	var v438 int32
	_ = v438
	var v441 int32
	_ = v441
	var v443 int32
	_ = v443
	var v445 int64
	_ = v445
	var v446 int64
	_ = v446
	var v447 int64
	_ = v447
	var v448 int64
	_ = v448
	var v449 int64
	_ = v449
	var v450 int64
	_ = v450
	var v451 int64
	_ = v451
	var v453 int64
	_ = v453
	var v455 int64
	_ = v455
	var v457 int64
	_ = v457
	var v459 int64
	_ = v459
	var v461 int64
	_ = v461
	var v463 int32
	_ = v463
	var v464 int32
	_ = v464
	var v465 int32
	_ = v465
	var v468 int32
	_ = v468
	var v471 int32
	_ = v471
	var v472 int32
	_ = v472
	var v473 int32
	_ = v473
	var v475 int32
	_ = v475
	var v478 int32
	_ = v478
	var v479 int32
	_ = v479
	var v483 int32
	_ = v483
	var v487 int32
	_ = v487
	var v489 int32
	_ = v489
	var v496 int32
	_ = v496
	var v497 int32
	_ = v497
	var v498 int32
	_ = v498
	var v503 int32
	_ = v503
	var v508 int32
	_ = v508
	var v512 int32
	_ = v512
	var v516 int32
	_ = v516
	var v526 int32
	_ = v526
	var v542 int32
	_ = v542
	var v545 int64
	_ = v545
	v6 = F_mstime(m)
	mBase = m.M
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+280))
	if v7 == int32(0) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	F_sentinelEvent(m, int32(3), int32(_a_F_sentinelFailoverDetectEnd_0), l0, int32(_a_F_sentinelFailoverDetectEnd_1), int32(0))
	mBase = m.M
	v542 = m.ExcPending
	if v542 != 0 {
		goto L7
	} else {
		goto L139
	}
L2:
	;
	return
L3:
	;
	v10 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7))))
	if v10&int32(8) != 0 {
		goto L2
	} else {
		goto L4
	}
L4:
	;
	v13 = *(*int64)(unsafe.Add(mBase, uint32(l0)+248))
	v14 = v6 - v13
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+148))
	v16 = F_dictGetIterator(m, v15)
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		goto L7
	} else {
		goto L8
	}
L5:
	;
	F_sentinelEvent(m, int32(3), int32(_a_F_sentinelFailoverDetectEnd_2), l0, int32(_a_F_sentinelFailoverDetectEnd_1), int32(0))
	mBase = m.M
	v265 = m.ExcPending
	if v265 != 0 {
		goto L7
	} else {
		goto L71
	}
L6:
	;
	F_dictReleaseIterator(m, v16)
	mBase = m.M
	v252 = m.ExcPending
	if v252 != 0 {
		goto L7
	} else {
		goto L69
	}
L7:
	;
	return
L8:
	;
	v25 = v16 + int32(20)
	v26 = *(*int32)(unsafe.Add(mBase, uint32(v16)+16))
	if v26 != 0 {
		goto L11
	} else {
		goto L12
	}
L9:
	;
	if v121 == int32(0) {
		goto L6
	} else {
		goto L35
	}
L10:
	;
	v32 = v25
	v33 = v29
	goto L13
L11:
	;
	v29 = int32(1)
	goto L10
L12:
	;
	v29 = int32(0)
	goto L10
L13:
	;
	switch v33 {
	case 0:
		goto L18
	default:
		goto L17
	}
L15:
	;
	v33 = int32(0)
	goto L13
L16:
	;
	goto L9
L17:
	;
	v113 = *(*int32)(unsafe.Add(mBase, uint32(v32)))
	*(*int32)(unsafe.Add(mBase, uint32(v16)+16)) = v113
	if v113 == int32(0) {
		goto L15
	} else {
		goto L34
	}
L18:
	;
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v16)+4))
	if v37 != int32(-1) {
		v76 = v37
		goto L19
	} else {
		goto L20
	}
L19:
	;
	v77 = int32(1)
	v78 = v76 + v77
	*(*int32)(unsafe.Add(mBase, uint32(v16)+4)) = v78
	v80 = int32(0)
	v83 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
	v84 = *(*int32)(unsafe.Add(mBase, uint32(v16)+8))
	v88 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v83+v84+int32(26)))))
	if v88 == int32(255) {
		goto L28
	} else {
		goto L29
	}
L20:
	;
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v16)+8))
	if v41 != 0 {
		v76 = int32(-1)
		goto L19
	} else {
		goto L21
	}
L21:
	;
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v16)+12))
	if v43 == int32(0) {
		goto L23
	} else {
		goto L24
	}
L22:
	;
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v69)+20))
	if v70 != int32(-1) {
		goto L25
	} else {
		goto L26
	}
L23:
	;
	v50 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v42)+16)))
	v51 = int64(*(*int8)(unsafe.Add(mBase, uint32(v42)+27)))
	v52 = int64(*(*int32)(unsafe.Add(mBase, uint32(v42)+8)))
	v53 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v42)+12)))
	v54 = int64(*(*int8)(unsafe.Add(mBase, uint32(v42)+26)))
	v55 = int64(*(*int32)(unsafe.Add(mBase, uint32(v42)+4)))
	v56 = F_wangHash64(m, v55)
	mBase = m.M
	v58 = F_wangHash64(m, v54+v56)
	mBase = m.M
	v60 = F_wangHash64(m, v53+v58)
	mBase = m.M
	v62 = F_wangHash64(m, v52+v60)
	mBase = m.M
	v64 = F_wangHash64(m, v51+v62)
	mBase = m.M
	v66 = F_wangHash64(m, v50+v64)
	mBase = m.M
	*(*int64)(unsafe.Add(mBase, uint32(v16)+24)) = v66
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
	v69 = v68
	goto L22
L24:
	;
	v46 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v42)+24)))
	v48 = v46 + int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v42)+24)) = uint16(v48)
	v69 = v42
	goto L22
L25:
	;
	v76 = v70 + int32(-1)
	goto L19
L26:
	;
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v16)+4))
	v76 = v73
	goto L19
L27:
	;
	v103 = int32(2)
	v108 = *(*int32)(unsafe.Add(mBase, uint32(v83+v101<<(uint(v103)%32)+int32(4))))
	v32 = v108 + v102<<(uint(v103)%32)
	v33 = int32(1)
	goto L13
L28:
	;
	v92 = v80
	goto L30
L29:
	;
	v92 = v77 << (uint(v88) % 32)
	goto L30
L30:
	;
	if v78 < v92 {
		v101 = v84
		v102 = v78
		goto L27
	} else {
		goto L31
	}
L31:
	;
	if v84 != 0 {
		v121 = v80
		goto L16
	} else {
		goto L32
	}
L32:
	;
	v94 = *(*int32)(unsafe.Add(mBase, uint32(v83)+20))
	if v94 == int32(-1) {
		v121 = v80
		goto L16
	} else {
		goto L33
	}
L33:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v16)+4)) = int64(4294967296)
	v101 = int32(1)
	v102 = int32(0)
	goto L27
L34:
	;
	v117 = *(*int32)(unsafe.Add(mBase, uint32(v113)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v25))) = v117
	v121 = v113
	goto L16
L35:
	;
	v130 = v121
	v132 = int32(0)
	goto L36
L36:
	;
	v133 = *(*int32)(unsafe.Add(mBase, uint32(v130)+8))
	goto L38
L37:
	;
	F_dictReleaseIterator(m, v16)
	mBase = m.M
	v248 = m.ExcPending
	if v248 != 0 {
		goto L7
	} else {
		goto L66
	}
L38:
	;
	v134 = *(*int32)(unsafe.Add(mBase, uint32(v133)))
	v139 = v132 + base.B2i32(v134&int32(1160) == int32(0))
	v147 = v16 + int32(20)
	v148 = *(*int32)(unsafe.Add(mBase, uint32(v16)+16))
	if v148 != 0 {
		goto L41
	} else {
		goto L42
	}
L39:
	;
	if v243 != 0 {
		v130 = v243
		v132 = v139
		goto L36
	} else {
		goto L65
	}
L40:
	;
	v154 = v147
	v155 = v151
	goto L43
L41:
	;
	v151 = int32(1)
	goto L40
L42:
	;
	v151 = int32(0)
	goto L40
L43:
	;
	switch v155 {
	case 0:
		goto L48
	default:
		goto L47
	}
L45:
	;
	v155 = int32(0)
	goto L43
L46:
	;
	goto L39
L47:
	;
	v235 = *(*int32)(unsafe.Add(mBase, uint32(v154)))
	*(*int32)(unsafe.Add(mBase, uint32(v16)+16)) = v235
	if v235 == int32(0) {
		goto L45
	} else {
		goto L64
	}
L48:
	;
	v159 = *(*int32)(unsafe.Add(mBase, uint32(v16)+4))
	if v159 != int32(-1) {
		v198 = v159
		goto L49
	} else {
		goto L50
	}
L49:
	;
	v199 = int32(1)
	v200 = v198 + v199
	*(*int32)(unsafe.Add(mBase, uint32(v16)+4)) = v200
	v202 = int32(0)
	v205 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
	v206 = *(*int32)(unsafe.Add(mBase, uint32(v16)+8))
	v210 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v205+v206+int32(26)))))
	if v210 == int32(255) {
		goto L58
	} else {
		goto L59
	}
L50:
	;
	v163 = *(*int32)(unsafe.Add(mBase, uint32(v16)+8))
	if v163 != 0 {
		v198 = int32(-1)
		goto L49
	} else {
		goto L51
	}
L51:
	;
	v164 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
	v165 = *(*int32)(unsafe.Add(mBase, uint32(v16)+12))
	if v165 == int32(0) {
		goto L53
	} else {
		goto L54
	}
L52:
	;
	v192 = *(*int32)(unsafe.Add(mBase, uint32(v191)+20))
	if v192 != int32(-1) {
		goto L55
	} else {
		goto L56
	}
L53:
	;
	v172 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v164)+16)))
	v173 = int64(*(*int8)(unsafe.Add(mBase, uint32(v164)+27)))
	v174 = int64(*(*int32)(unsafe.Add(mBase, uint32(v164)+8)))
	v175 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v164)+12)))
	v176 = int64(*(*int8)(unsafe.Add(mBase, uint32(v164)+26)))
	v177 = int64(*(*int32)(unsafe.Add(mBase, uint32(v164)+4)))
	v178 = F_wangHash64(m, v177)
	mBase = m.M
	v180 = F_wangHash64(m, v176+v178)
	mBase = m.M
	v182 = F_wangHash64(m, v175+v180)
	mBase = m.M
	v184 = F_wangHash64(m, v174+v182)
	mBase = m.M
	v186 = F_wangHash64(m, v173+v184)
	mBase = m.M
	v188 = F_wangHash64(m, v172+v186)
	mBase = m.M
	*(*int64)(unsafe.Add(mBase, uint32(v16)+24)) = v188
	v190 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
	v191 = v190
	goto L52
L54:
	;
	v168 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v164)+24)))
	v170 = v168 + int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v164)+24)) = uint16(v170)
	v191 = v164
	goto L52
L55:
	;
	v198 = v192 + int32(-1)
	goto L49
L56:
	;
	v195 = *(*int32)(unsafe.Add(mBase, uint32(v16)+4))
	v198 = v195
	goto L49
L57:
	;
	v225 = int32(2)
	v230 = *(*int32)(unsafe.Add(mBase, uint32(v205+v223<<(uint(v225)%32)+int32(4))))
	v154 = v230 + v224<<(uint(v225)%32)
	v155 = int32(1)
	goto L43
L58:
	;
	v214 = v202
	goto L60
L59:
	;
	v214 = v199 << (uint(v210) % 32)
	goto L60
L60:
	;
	if v200 < v214 {
		v223 = v206
		v224 = v200
		goto L57
	} else {
		goto L61
	}
L61:
	;
	if v206 != 0 {
		v243 = v202
		goto L46
	} else {
		goto L62
	}
L62:
	;
	v216 = *(*int32)(unsafe.Add(mBase, uint32(v205)+20))
	if v216 == int32(-1) {
		v243 = v202
		goto L46
	} else {
		goto L63
	}
L63:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v16)+4)) = int64(4294967296)
	v223 = int32(1)
	v224 = int32(0)
	goto L57
L64:
	;
	v239 = *(*int32)(unsafe.Add(mBase, uint32(v235)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v147))) = v239
	v243 = v235
	goto L46
L65:
	;
	goto L37
L66:
	;
	v249 = *(*int64)(unsafe.Add(mBase, uint32(l0)+264))
	if v249 < v14 {
		goto L5
	} else {
		goto L67
	}
L67:
	;
	if v139 != 0 {
		goto L2
	} else {
		goto L68
	}
L68:
	;
	goto L1
L69:
	;
	v253 = *(*int64)(unsafe.Add(mBase, uint32(l0)+264))
	if v14 <= v253 {
		goto L1
	} else {
		goto L70
	}
L70:
	;
	goto L5
L71:
	;
	F_sentinelEvent(m, int32(3), int32(_a_F_sentinelFailoverDetectEnd_0), l0, int32(_a_F_sentinelFailoverDetectEnd_1), int32(0))
	mBase = m.M
	v271 = m.ExcPending
	if v271 != 0 {
		goto L7
	} else {
		goto L72
	}
L72:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+240)) = int32(6)
	v274 = F_mstime(m)
	mBase = m.M
	*(*int64)(unsafe.Add(mBase, uint32(l0)+248)) = v274
	v276 = *(*int32)(unsafe.Add(mBase, uint32(l0)+148))
	v277 = F_dictGetIterator(m, v276)
	mBase = m.M
	v278 = m.ExcPending
	if v278 != 0 {
		goto L7
	} else {
		goto L74
	}
L73:
	;
	F_dictReleaseIterator(m, v277)
	mBase = m.M
	v526 = m.ExcPending
	if v526 != 0 {
		goto L7
	} else {
		goto L138
	}
L74:
	;
	v286 = v277 + int32(20)
	v287 = *(*int32)(unsafe.Add(mBase, uint32(v277)+16))
	if v287 != 0 {
		goto L77
	} else {
		goto L78
	}
L75:
	;
	if v382 == int32(0) {
		goto L73
	} else {
		goto L101
	}
L76:
	;
	v293 = v286
	v294 = v290
	goto L79
L77:
	;
	v290 = int32(1)
	goto L76
L78:
	;
	v290 = int32(0)
	goto L76
L79:
	;
	switch v294 {
	case 0:
		goto L84
	default:
		goto L83
	}
L81:
	;
	v294 = int32(0)
	goto L79
L82:
	;
	goto L75
L83:
	;
	v374 = *(*int32)(unsafe.Add(mBase, uint32(v293)))
	*(*int32)(unsafe.Add(mBase, uint32(v277)+16)) = v374
	if v374 == int32(0) {
		goto L81
	} else {
		goto L100
	}
L84:
	;
	v298 = *(*int32)(unsafe.Add(mBase, uint32(v277)+4))
	if v298 != int32(-1) {
		v337 = v298
		goto L85
	} else {
		goto L86
	}
L85:
	;
	v338 = int32(1)
	v339 = v337 + v338
	*(*int32)(unsafe.Add(mBase, uint32(v277)+4)) = v339
	v341 = int32(0)
	v344 = *(*int32)(unsafe.Add(mBase, uint32(v277)))
	v345 = *(*int32)(unsafe.Add(mBase, uint32(v277)+8))
	v349 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v344+v345+int32(26)))))
	if v349 == int32(255) {
		goto L94
	} else {
		goto L95
	}
L86:
	;
	v302 = *(*int32)(unsafe.Add(mBase, uint32(v277)+8))
	if v302 != 0 {
		v337 = int32(-1)
		goto L85
	} else {
		goto L87
	}
L87:
	;
	v303 = *(*int32)(unsafe.Add(mBase, uint32(v277)))
	v304 = *(*int32)(unsafe.Add(mBase, uint32(v277)+12))
	if v304 == int32(0) {
		goto L89
	} else {
		goto L90
	}
L88:
	;
	v331 = *(*int32)(unsafe.Add(mBase, uint32(v330)+20))
	if v331 != int32(-1) {
		goto L91
	} else {
		goto L92
	}
L89:
	;
	v311 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v303)+16)))
	v312 = int64(*(*int8)(unsafe.Add(mBase, uint32(v303)+27)))
	v313 = int64(*(*int32)(unsafe.Add(mBase, uint32(v303)+8)))
	v314 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v303)+12)))
	v315 = int64(*(*int8)(unsafe.Add(mBase, uint32(v303)+26)))
	v316 = int64(*(*int32)(unsafe.Add(mBase, uint32(v303)+4)))
	v317 = F_wangHash64(m, v316)
	mBase = m.M
	v319 = F_wangHash64(m, v315+v317)
	mBase = m.M
	v321 = F_wangHash64(m, v314+v319)
	mBase = m.M
	v323 = F_wangHash64(m, v313+v321)
	mBase = m.M
	v325 = F_wangHash64(m, v312+v323)
	mBase = m.M
	v327 = F_wangHash64(m, v311+v325)
	mBase = m.M
	*(*int64)(unsafe.Add(mBase, uint32(v277)+24)) = v327
	v329 = *(*int32)(unsafe.Add(mBase, uint32(v277)))
	v330 = v329
	goto L88
L90:
	;
	v307 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v303)+24)))
	v309 = v307 + int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v303)+24)) = uint16(v309)
	v330 = v303
	goto L88
L91:
	;
	v337 = v331 + int32(-1)
	goto L85
L92:
	;
	v334 = *(*int32)(unsafe.Add(mBase, uint32(v277)+4))
	v337 = v334
	goto L85
L93:
	;
	v364 = int32(2)
	v369 = *(*int32)(unsafe.Add(mBase, uint32(v344+v362<<(uint(v364)%32)+int32(4))))
	v293 = v369 + v363<<(uint(v364)%32)
	v294 = int32(1)
	goto L79
L94:
	;
	v353 = v341
	goto L96
L95:
	;
	v353 = v338 << (uint(v349) % 32)
	goto L96
L96:
	;
	if v339 < v353 {
		v362 = v345
		v363 = v339
		goto L93
	} else {
		goto L97
	}
L97:
	;
	if v345 != 0 {
		v382 = v341
		goto L82
	} else {
		goto L98
	}
L98:
	;
	v355 = *(*int32)(unsafe.Add(mBase, uint32(v344)+20))
	if v355 == int32(-1) {
		v382 = v341
		goto L82
	} else {
		goto L99
	}
L99:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v277)+4)) = int64(4294967296)
	v362 = int32(1)
	v363 = int32(0)
	goto L93
L100:
	;
	v378 = *(*int32)(unsafe.Add(mBase, uint32(v374)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v286))) = v378
	v382 = v374
	goto L82
L101:
	;
	v390 = v382
	goto L102
L102:
	;
	v393 = *(*int32)(unsafe.Add(mBase, uint32(v390)+8))
	goto L105
L103:
	;
	goto L73
L104:
	;
	v420 = v277 + int32(20)
	v421 = *(*int32)(unsafe.Add(mBase, uint32(v277)+16))
	if v421 != 0 {
		goto L113
	} else {
		goto L114
	}
L105:
	;
	v394 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v393))))
	if v394&int32(1408) != 0 {
		goto L104
	} else {
		goto L106
	}
L106:
	;
	v397 = *(*int32)(unsafe.Add(mBase, uint32(v393)+28))
	v398 = *(*int32)(unsafe.Add(mBase, uint32(v397)+4))
	if v398 != 0 {
		goto L104
	} else {
		goto L107
	}
L107:
	;
	v399 = *(*int32)(unsafe.Add(mBase, uint32(l0)+280))
	v400 = *(*int32)(unsafe.Add(mBase, uint32(v399)+24))
	v401 = F_sentinelSendReplicaOf(m, v393, v400)
	mBase = m.M
	v402 = m.ExcPending
	if v402 != 0 {
		goto L7
	} else {
		goto L108
	}
L108:
	;
	if v401 != 0 {
		goto L104
	} else {
		goto L109
	}
L109:
	;
	F_sentinelEvent(m, int32(2), int32(_a_F_sentinelFailoverDetectEnd_3), v393, int32(_a_F_sentinelFailoverDetectEnd_1), int32(0))
	mBase = m.M
	v408 = m.ExcPending
	if v408 != 0 {
		goto L7
	} else {
		goto L110
	}
L110:
	;
	v409 = *(*int32)(unsafe.Add(mBase, uint32(v393)))
	*(*int32)(unsafe.Add(mBase, uint32(v393))) = v409 | int32(256)
	goto L104
L111:
	;
	if v516 != 0 {
		v390 = v516
		goto L102
	} else {
		goto L137
	}
L112:
	;
	v427 = v420
	v428 = v424
	goto L115
L113:
	;
	v424 = int32(1)
	goto L112
L114:
	;
	v424 = int32(0)
	goto L112
L115:
	;
	switch v428 {
	case 0:
		goto L120
	default:
		goto L119
	}
L117:
	;
	v428 = int32(0)
	goto L115
L118:
	;
	goto L111
L119:
	;
	v508 = *(*int32)(unsafe.Add(mBase, uint32(v427)))
	*(*int32)(unsafe.Add(mBase, uint32(v277)+16)) = v508
	if v508 == int32(0) {
		goto L117
	} else {
		goto L136
	}
L120:
	;
	v432 = *(*int32)(unsafe.Add(mBase, uint32(v277)+4))
	if v432 != int32(-1) {
		v471 = v432
		goto L121
	} else {
		goto L122
	}
L121:
	;
	v472 = int32(1)
	v473 = v471 + v472
	*(*int32)(unsafe.Add(mBase, uint32(v277)+4)) = v473
	v475 = int32(0)
	v478 = *(*int32)(unsafe.Add(mBase, uint32(v277)))
	v479 = *(*int32)(unsafe.Add(mBase, uint32(v277)+8))
	v483 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v478+v479+int32(26)))))
	if v483 == int32(255) {
		goto L130
	} else {
		goto L131
	}
L122:
	;
	v436 = *(*int32)(unsafe.Add(mBase, uint32(v277)+8))
	if v436 != 0 {
		v471 = int32(-1)
		goto L121
	} else {
		goto L123
	}
L123:
	;
	v437 = *(*int32)(unsafe.Add(mBase, uint32(v277)))
	v438 = *(*int32)(unsafe.Add(mBase, uint32(v277)+12))
	if v438 == int32(0) {
		goto L125
	} else {
		goto L126
	}
L124:
	;
	v465 = *(*int32)(unsafe.Add(mBase, uint32(v464)+20))
	if v465 != int32(-1) {
		goto L127
	} else {
		goto L128
	}
L125:
	;
	v445 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v437)+16)))
	v446 = int64(*(*int8)(unsafe.Add(mBase, uint32(v437)+27)))
	v447 = int64(*(*int32)(unsafe.Add(mBase, uint32(v437)+8)))
	v448 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v437)+12)))
	v449 = int64(*(*int8)(unsafe.Add(mBase, uint32(v437)+26)))
	v450 = int64(*(*int32)(unsafe.Add(mBase, uint32(v437)+4)))
	v451 = F_wangHash64(m, v450)
	mBase = m.M
	v453 = F_wangHash64(m, v449+v451)
	mBase = m.M
	v455 = F_wangHash64(m, v448+v453)
	mBase = m.M
	v457 = F_wangHash64(m, v447+v455)
	mBase = m.M
	v459 = F_wangHash64(m, v446+v457)
	mBase = m.M
	v461 = F_wangHash64(m, v445+v459)
	mBase = m.M
	*(*int64)(unsafe.Add(mBase, uint32(v277)+24)) = v461
	v463 = *(*int32)(unsafe.Add(mBase, uint32(v277)))
	v464 = v463
	goto L124
L126:
	;
	v441 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v437)+24)))
	v443 = v441 + int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v437)+24)) = uint16(v443)
	v464 = v437
	goto L124
L127:
	;
	v471 = v465 + int32(-1)
	goto L121
L128:
	;
	v468 = *(*int32)(unsafe.Add(mBase, uint32(v277)+4))
	v471 = v468
	goto L121
L129:
	;
	v498 = int32(2)
	v503 = *(*int32)(unsafe.Add(mBase, uint32(v478+v496<<(uint(v498)%32)+int32(4))))
	v427 = v503 + v497<<(uint(v498)%32)
	v428 = int32(1)
	goto L115
L130:
	;
	v487 = v475
	goto L132
L131:
	;
	v487 = v472 << (uint(v483) % 32)
	goto L132
L132:
	;
	if v473 < v487 {
		v496 = v479
		v497 = v473
		goto L129
	} else {
		goto L133
	}
L133:
	;
	if v479 != 0 {
		v516 = v475
		goto L118
	} else {
		goto L134
	}
L134:
	;
	v489 = *(*int32)(unsafe.Add(mBase, uint32(v478)+20))
	if v489 == int32(-1) {
		v516 = v475
		goto L118
	} else {
		goto L135
	}
L135:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v277)+4)) = int64(4294967296)
	v496 = int32(1)
	v497 = int32(0)
	goto L129
L136:
	;
	v512 = *(*int32)(unsafe.Add(mBase, uint32(v508)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v420))) = v512
	v516 = v508
	goto L118
L137:
	;
	goto L103
L138:
	;
	goto L2
L139:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+240)) = int32(6)
	v545 = F_mstime(m)
	mBase = m.M
	*(*int64)(unsafe.Add(mBase, uint32(l0)+248)) = v545
	return
}
func F_sentinelFailoverSendFailover(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v3 int64
	_ = v3
	var v4 int64
	_ = v4
	var v6 int64
	_ = v6
	var v7 int64
	_ = v7
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v30 int32
	_ = v30
	var v33 int64
	_ = v33
	v3 = F_mstime(m)
	mBase = m.M
	v4 = *(*int64)(unsafe.Add(mBase, uint32(l0)+248))
	v6 = *(*int64)(unsafe.Add(mBase, uint32(l0)+264))
	v7 = v4 - v3 + v6
	if int64(999) < v7 {
		v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v19 = *(*int32)(unsafe.Add(mBase, uint32(v18)+4))
		if v19 != 0 {
			return
		} else {
			v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+280))
			v21 = *(*int32)(unsafe.Add(mBase, uint32(v20)+24))
			v22 = F_sentinelFailoverTo(m, l0, v21, v7)
			mBase = m.M
			v23 = m.ExcPending
			if v23 != 0 {
				return
			} else {
				if v22 != 0 {
					return
				} else {
					v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+280))
					F_sentinelEvent(m, int32(2), int32(_a_F_sentinelFailoverSendFailover_0), v26, int32(_a_F_sentinelFailoverSendFailover_1), int32(0))
					mBase = m.M
					v30 = m.ExcPending
					if v30 != 0 {
						return
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(l0)+240)) = int32(4)
						v33 = F_mstime(m)
						mBase = m.M
						*(*int64)(unsafe.Add(mBase, uint32(l0)+248)) = v33
						return
					}
				}
			}
		}
	} else {
		F_sentinelEvent(m, int32(3), int32(_a_F_sentinelFailoverSendFailover_2), l0, int32(_a_F_sentinelFailoverSendFailover_1), int32(0))
		mBase = m.M
		v15 = m.ExcPending
		if v15 != 0 {
			return
		} else {
			F_sentinelAbortFailover(m, l0)
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
func F_sentinelFailoverWaitStart(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v6 int64
	_ = v6
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
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
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v59 int64
	_ = v59
	var v61 int64
	_ = v61
	var v62 int64
	_ = v62
	var v63 int64
	_ = v63
	var v66 int64
	_ = v66
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v90 int32
	_ = v90
	var v97 int32
	_ = v97
	var v102 int64
	_ = v102
	var v109 int32
	_ = v109
	v6 = *(*int64)(unsafe.Add(mBase, uint32(l0)+232))
	v7 = F_sentinelGetLeader(m, l0, v6)
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		goto L5
	} else {
		goto L6
	}
L1:
	;
	return
L2:
	;
	F_sentinelEvent(m, int32(3), int32(_a_F_sentinelFailoverWaitStart_0), l0, int32(_a_F_sentinelFailoverWaitStart_1), int32(0))
	mBase = m.M
	v82 = m.ExcPending
	if v82 != 0 {
		goto L5
	} else {
		goto L30
	}
L3:
	;
	v56 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+1)))
	if v56&int32(8) != 0 {
		goto L2
	} else {
		goto L23
	}
L4:
	;
	F_sdsfree(m, int32(0))
	mBase = m.M
	v54 = m.ExcPending
	if v54 != 0 {
		goto L5
	} else {
		goto L22
	}
L5:
	;
	return
L6:
	;
	if v7 == int32(0) {
		goto L4
	} else {
		goto L7
	}
L7:
	;
	v11 = int32(_a_F_sentinelFailoverWaitStart_2)
	v14 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7))))
	if v14 != 0 {
		goto L10
	} else {
		goto L11
	}
L8:
	;
	F_sdsfree(m, v7)
	mBase = m.M
	v51 = m.ExcPending
	if v51 != 0 {
		goto L5
	} else {
		goto L20
	}
L9:
	;
	v46 = F_tolower(m, v42)
	mBase = m.M
	v47 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v43))))
	v48 = F_tolower(m, v47)
	mBase = m.M
	goto L8
L10:
	;
	v16 = v7
	v17 = v11
	v18 = v14
	goto L13
L11:
	;
	v42 = int32(0)
	v43 = v11
	goto L9
L12:
	;
	v42 = v39 & int32(255)
	v43 = v38
	goto L9
L13:
	;
	v20 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17))))
	if v20 == int32(0) {
		v38 = v17
		v39 = v18
		goto L12
	} else {
		goto L15
	}
L14:
	;
	v38 = v32
	v39 = int32(0)
	goto L12
L15:
	;
	v24 = v18 & int32(255)
	if v24 == v20 {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	v31 = int32(1)
	v32 = v17 + v31
	v33 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16)+1)))
	if v33 != 0 {
		v16 = v16 + v31
		v17 = v32
		v18 = v33
		goto L13
	} else {
		goto L19
	}
L17:
	;
	v26 = F_tolower(m, v24)
	mBase = m.M
	v27 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17))))
	v28 = F_tolower(m, v27)
	mBase = m.M
	if v26 == v28 {
		goto L16
	} else {
		goto L18
	}
L18:
	;
	v30 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16))))
	v38 = v17
	v39 = v30
	goto L12
L19:
	;
	goto L14
L20:
	;
	if v46-v48 != 0 {
		goto L3
	} else {
		goto L21
	}
L21:
	;
	goto L2
L22:
	;
	goto L3
L23:
	;
	v59 = *(*int64)(unsafe.Add(mBase, uint32(l0)+264))
	v61 = *(*int64)(unsafe.Add(mBase, _c_F_sentinelFailoverWaitStart[0]))
	v62 = F_mstime(m)
	mBase = m.M
	v63 = *(*int64)(unsafe.Add(mBase, uint32(l0)+256))
	if v61 < v59 {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	v66 = v61
	goto L26
L25:
	;
	v66 = v59
	goto L26
L26:
	;
	if v62-v63 <= v66 {
		goto L1
	} else {
		goto L27
	}
L27:
	;
	F_sentinelEvent(m, int32(3), int32(_a_F_sentinelFailoverWaitStart_3), l0, int32(_a_F_sentinelFailoverWaitStart_1), int32(0))
	mBase = m.M
	v73 = m.ExcPending
	if v73 != 0 {
		goto L5
	} else {
		goto L28
	}
L28:
	;
	F_sentinelAbortFailover(m, l0)
	mBase = m.M
	v75 = m.ExcPending
	if v75 != 0 {
		goto L5
	} else {
		goto L29
	}
L29:
	;
	return
L30:
	;
	v83 = int32(0)
	v84 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_sentinelFailoverWaitStart[1])))
	if v84&int32(1) == v83 {
		goto L31
	} else {
		goto L32
	}
L31:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+240)) = int32(2)
	v102 = F_mstime(m)
	mBase = m.M
	*(*int64)(unsafe.Add(mBase, uint32(l0)+248)) = v102
	F_sentinelEvent(m, int32(3), int32(_a_F_sentinelFailoverWaitStart_4), l0, int32(_a_F_sentinelFailoverWaitStart_1), int32(0))
	mBase = m.M
	v109 = m.ExcPending
	if v109 != 0 {
		goto L5
	} else {
		goto L36
	}
L32:
	;
	v90 = *(*int32)(unsafe.Add(mBase, _c_F_sentinelFailoverWaitStart[2]))
	if int32(3) < v90 {
		goto L33
	} else {
		goto L34
	}
L33:
	;
	m.Env.Exit(m, int32(99))
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L34:
	;
	F__serverLog(m, int32(3), int32(_a_F_sentinelFailoverWaitStart_5), int32(0))
	mBase = m.M
	v97 = m.ExcPending
	if v97 != 0 {
		goto L5
	} else {
		goto L35
	}
L35:
	;
	goto L33
L36:
	;
	goto L1
}
func F_sentinelForceHelloUpdateForPrimary(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v12 int64
	_ = v12
	var v14 int64
	_ = v14
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
	var v135 int32
	_ = v135
	var v139 int32
	_ = v139
	var v140 int64
	_ = v140
	var v142 int64
	_ = v142
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v159 int32
	_ = v159
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v167 int32
	_ = v167
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
	var v180 int64
	_ = v180
	var v181 int64
	_ = v181
	var v182 int64
	_ = v182
	var v183 int64
	_ = v183
	var v184 int64
	_ = v184
	var v185 int64
	_ = v185
	var v186 int64
	_ = v186
	var v188 int64
	_ = v188
	var v190 int64
	_ = v190
	var v192 int64
	_ = v192
	var v194 int64
	_ = v194
	var v196 int64
	_ = v196
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v203 int32
	_ = v203
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v210 int32
	_ = v210
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v218 int32
	_ = v218
	var v222 int32
	_ = v222
	var v224 int32
	_ = v224
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v238 int32
	_ = v238
	var v243 int32
	_ = v243
	var v247 int32
	_ = v247
	var v251 int32
	_ = v251
	var v261 int32
	_ = v261
	var v262 int32
	_ = v262
	var v263 int32
	_ = v263
	var v264 int32
	_ = v264
	var v272 int32
	_ = v272
	var v273 int32
	_ = v273
	var v276 int32
	_ = v276
	var v279 int32
	_ = v279
	var v280 int32
	_ = v280
	var v284 int32
	_ = v284
	var v288 int32
	_ = v288
	var v289 int32
	_ = v289
	var v290 int32
	_ = v290
	var v293 int32
	_ = v293
	var v295 int32
	_ = v295
	var v297 int64
	_ = v297
	var v298 int64
	_ = v298
	var v299 int64
	_ = v299
	var v300 int64
	_ = v300
	var v301 int64
	_ = v301
	var v302 int64
	_ = v302
	var v303 int64
	_ = v303
	var v305 int64
	_ = v305
	var v307 int64
	_ = v307
	var v309 int64
	_ = v309
	var v311 int64
	_ = v311
	var v313 int64
	_ = v313
	var v315 int32
	_ = v315
	var v316 int32
	_ = v316
	var v317 int32
	_ = v317
	var v320 int32
	_ = v320
	var v323 int32
	_ = v323
	var v324 int32
	_ = v324
	var v325 int32
	_ = v325
	var v327 int32
	_ = v327
	var v330 int32
	_ = v330
	var v331 int32
	_ = v331
	var v335 int32
	_ = v335
	var v339 int32
	_ = v339
	var v341 int32
	_ = v341
	var v348 int32
	_ = v348
	var v349 int32
	_ = v349
	var v350 int32
	_ = v350
	var v355 int32
	_ = v355
	var v360 int32
	_ = v360
	var v364 int32
	_ = v364
	var v368 int32
	_ = v368
	var v375 int32
	_ = v375
	var v379 int32
	_ = v379
	var v380 int64
	_ = v380
	var v382 int64
	_ = v382
	var v395 int32
	_ = v395
	var v396 int32
	_ = v396
	var v399 int32
	_ = v399
	var v402 int32
	_ = v402
	var v403 int32
	_ = v403
	var v407 int32
	_ = v407
	var v411 int32
	_ = v411
	var v412 int32
	_ = v412
	var v413 int32
	_ = v413
	var v416 int32
	_ = v416
	var v418 int32
	_ = v418
	var v420 int64
	_ = v420
	var v421 int64
	_ = v421
	var v422 int64
	_ = v422
	var v423 int64
	_ = v423
	var v424 int64
	_ = v424
	var v425 int64
	_ = v425
	var v426 int64
	_ = v426
	var v428 int64
	_ = v428
	var v430 int64
	_ = v430
	var v432 int64
	_ = v432
	var v434 int64
	_ = v434
	var v436 int64
	_ = v436
	var v438 int32
	_ = v438
	var v439 int32
	_ = v439
	var v440 int32
	_ = v440
	var v443 int32
	_ = v443
	var v446 int32
	_ = v446
	var v447 int32
	_ = v447
	var v448 int32
	_ = v448
	var v450 int32
	_ = v450
	var v453 int32
	_ = v453
	var v454 int32
	_ = v454
	var v458 int32
	_ = v458
	var v462 int32
	_ = v462
	var v464 int32
	_ = v464
	var v471 int32
	_ = v471
	var v472 int32
	_ = v472
	var v473 int32
	_ = v473
	var v478 int32
	_ = v478
	var v483 int32
	_ = v483
	var v487 int32
	_ = v487
	var v491 int32
	_ = v491
	var v501 int32
	_ = v501
	var v504 int32
	_ = v504
	v7 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v7&int32(1) == int32(0) {
		v504 = int32(-1)
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return v504
L2:
	;
	v12 = *(*int64)(unsafe.Add(mBase, uint32(l0)+32))
	v14 = *(*int64)(unsafe.Add(mBase, _c_F_sentinelForceHelloUpdateForPrimary[0]))
	if v12 <= v14 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+144))
	v21 = F_dictGetSafeIterator(m, v20)
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l0)+32)) = v12 + (v14 ^ int64(-1))
	goto L3
L5:
	;
	F_dictReleaseIterator(m, v21)
	mBase = m.M
	v261 = m.ExcPending
	if v261 != 0 {
		goto L6
	} else {
		goto L67
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
	v135 = v128
	goto L35
L35:
	;
	v139 = *(*int32)(unsafe.Add(mBase, uint32(v135)+8))
	goto L38
L36:
	;
	goto L5
L37:
	;
	v155 = v21 + int32(20)
	v156 = *(*int32)(unsafe.Add(mBase, uint32(v21)+16))
	if v156 != 0 {
		goto L42
	} else {
		goto L43
	}
L38:
	;
	v140 = *(*int64)(unsafe.Add(mBase, uint32(v139)+32))
	v142 = *(*int64)(unsafe.Add(mBase, _c_F_sentinelForceHelloUpdateForPrimary[0]))
	if v140 <= v142 {
		goto L37
	} else {
		goto L39
	}
L39:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v139)+32)) = v140 + (v142 ^ int64(-1))
	goto L37
L40:
	;
	if v251 != 0 {
		v135 = v251
		goto L35
	} else {
		goto L66
	}
L41:
	;
	v162 = v155
	v163 = v159
	goto L44
L42:
	;
	v159 = int32(1)
	goto L41
L43:
	;
	v159 = int32(0)
	goto L41
L44:
	;
	switch v163 {
	case 0:
		goto L49
	default:
		goto L48
	}
L46:
	;
	v163 = int32(0)
	goto L44
L47:
	;
	goto L40
L48:
	;
	v243 = *(*int32)(unsafe.Add(mBase, uint32(v162)))
	*(*int32)(unsafe.Add(mBase, uint32(v21)+16)) = v243
	if v243 == int32(0) {
		goto L46
	} else {
		goto L65
	}
L49:
	;
	v167 = *(*int32)(unsafe.Add(mBase, uint32(v21)+4))
	if v167 != int32(-1) {
		v206 = v167
		goto L50
	} else {
		goto L51
	}
L50:
	;
	v207 = int32(1)
	v208 = v206 + v207
	*(*int32)(unsafe.Add(mBase, uint32(v21)+4)) = v208
	v210 = int32(0)
	v213 = *(*int32)(unsafe.Add(mBase, uint32(v21)))
	v214 = *(*int32)(unsafe.Add(mBase, uint32(v21)+8))
	v218 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v213+v214+int32(26)))))
	if v218 == int32(255) {
		goto L59
	} else {
		goto L60
	}
L51:
	;
	v171 = *(*int32)(unsafe.Add(mBase, uint32(v21)+8))
	if v171 != 0 {
		v206 = int32(-1)
		goto L50
	} else {
		goto L52
	}
L52:
	;
	v172 = *(*int32)(unsafe.Add(mBase, uint32(v21)))
	v173 = *(*int32)(unsafe.Add(mBase, uint32(v21)+12))
	if v173 == int32(0) {
		goto L54
	} else {
		goto L55
	}
L53:
	;
	v200 = *(*int32)(unsafe.Add(mBase, uint32(v199)+20))
	if v200 != int32(-1) {
		goto L56
	} else {
		goto L57
	}
L54:
	;
	v180 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v172)+16)))
	v181 = int64(*(*int8)(unsafe.Add(mBase, uint32(v172)+27)))
	v182 = int64(*(*int32)(unsafe.Add(mBase, uint32(v172)+8)))
	v183 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v172)+12)))
	v184 = int64(*(*int8)(unsafe.Add(mBase, uint32(v172)+26)))
	v185 = int64(*(*int32)(unsafe.Add(mBase, uint32(v172)+4)))
	v186 = F_wangHash64(m, v185)
	mBase = m.M
	v188 = F_wangHash64(m, v184+v186)
	mBase = m.M
	v190 = F_wangHash64(m, v183+v188)
	mBase = m.M
	v192 = F_wangHash64(m, v182+v190)
	mBase = m.M
	v194 = F_wangHash64(m, v181+v192)
	mBase = m.M
	v196 = F_wangHash64(m, v180+v194)
	mBase = m.M
	*(*int64)(unsafe.Add(mBase, uint32(v21)+24)) = v196
	v198 = *(*int32)(unsafe.Add(mBase, uint32(v21)))
	v199 = v198
	goto L53
L55:
	;
	v176 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v172)+24)))
	v178 = v176 + int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v172)+24)) = uint16(v178)
	v199 = v172
	goto L53
L56:
	;
	v206 = v200 + int32(-1)
	goto L50
L57:
	;
	v203 = *(*int32)(unsafe.Add(mBase, uint32(v21)+4))
	v206 = v203
	goto L50
L58:
	;
	v233 = int32(2)
	v238 = *(*int32)(unsafe.Add(mBase, uint32(v213+v231<<(uint(v233)%32)+int32(4))))
	v162 = v238 + v232<<(uint(v233)%32)
	v163 = int32(1)
	goto L44
L59:
	;
	v222 = v210
	goto L61
L60:
	;
	v222 = v207 << (uint(v218) % 32)
	goto L61
L61:
	;
	if v208 < v222 {
		v231 = v214
		v232 = v208
		goto L58
	} else {
		goto L62
	}
L62:
	;
	if v214 != 0 {
		v251 = v210
		goto L47
	} else {
		goto L63
	}
L63:
	;
	v224 = *(*int32)(unsafe.Add(mBase, uint32(v213)+20))
	if v224 == int32(-1) {
		v251 = v210
		goto L47
	} else {
		goto L64
	}
L64:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v21)+4)) = int64(4294967296)
	v231 = int32(1)
	v232 = int32(0)
	goto L58
L65:
	;
	v247 = *(*int32)(unsafe.Add(mBase, uint32(v243)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v155))) = v247
	v251 = v243
	goto L47
L66:
	;
	goto L36
L67:
	;
	v262 = *(*int32)(unsafe.Add(mBase, uint32(l0)+148))
	v263 = F_dictGetSafeIterator(m, v262)
	mBase = m.M
	v264 = m.ExcPending
	if v264 != 0 {
		goto L6
	} else {
		goto L69
	}
L68:
	;
	F_dictReleaseIterator(m, v263)
	mBase = m.M
	v501 = m.ExcPending
	if v501 != 0 {
		goto L6
	} else {
		goto L129
	}
L69:
	;
	v272 = v263 + int32(20)
	v273 = *(*int32)(unsafe.Add(mBase, uint32(v263)+16))
	if v273 != 0 {
		goto L72
	} else {
		goto L73
	}
L70:
	;
	if v368 == int32(0) {
		goto L68
	} else {
		goto L96
	}
L71:
	;
	v279 = v272
	v280 = v276
	goto L74
L72:
	;
	v276 = int32(1)
	goto L71
L73:
	;
	v276 = int32(0)
	goto L71
L74:
	;
	switch v280 {
	case 0:
		goto L79
	default:
		goto L78
	}
L76:
	;
	v280 = int32(0)
	goto L74
L77:
	;
	goto L70
L78:
	;
	v360 = *(*int32)(unsafe.Add(mBase, uint32(v279)))
	*(*int32)(unsafe.Add(mBase, uint32(v263)+16)) = v360
	if v360 == int32(0) {
		goto L76
	} else {
		goto L95
	}
L79:
	;
	v284 = *(*int32)(unsafe.Add(mBase, uint32(v263)+4))
	if v284 != int32(-1) {
		v323 = v284
		goto L80
	} else {
		goto L81
	}
L80:
	;
	v324 = int32(1)
	v325 = v323 + v324
	*(*int32)(unsafe.Add(mBase, uint32(v263)+4)) = v325
	v327 = int32(0)
	v330 = *(*int32)(unsafe.Add(mBase, uint32(v263)))
	v331 = *(*int32)(unsafe.Add(mBase, uint32(v263)+8))
	v335 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v330+v331+int32(26)))))
	if v335 == int32(255) {
		goto L89
	} else {
		goto L90
	}
L81:
	;
	v288 = *(*int32)(unsafe.Add(mBase, uint32(v263)+8))
	if v288 != 0 {
		v323 = int32(-1)
		goto L80
	} else {
		goto L82
	}
L82:
	;
	v289 = *(*int32)(unsafe.Add(mBase, uint32(v263)))
	v290 = *(*int32)(unsafe.Add(mBase, uint32(v263)+12))
	if v290 == int32(0) {
		goto L84
	} else {
		goto L85
	}
L83:
	;
	v317 = *(*int32)(unsafe.Add(mBase, uint32(v316)+20))
	if v317 != int32(-1) {
		goto L86
	} else {
		goto L87
	}
L84:
	;
	v297 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v289)+16)))
	v298 = int64(*(*int8)(unsafe.Add(mBase, uint32(v289)+27)))
	v299 = int64(*(*int32)(unsafe.Add(mBase, uint32(v289)+8)))
	v300 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v289)+12)))
	v301 = int64(*(*int8)(unsafe.Add(mBase, uint32(v289)+26)))
	v302 = int64(*(*int32)(unsafe.Add(mBase, uint32(v289)+4)))
	v303 = F_wangHash64(m, v302)
	mBase = m.M
	v305 = F_wangHash64(m, v301+v303)
	mBase = m.M
	v307 = F_wangHash64(m, v300+v305)
	mBase = m.M
	v309 = F_wangHash64(m, v299+v307)
	mBase = m.M
	v311 = F_wangHash64(m, v298+v309)
	mBase = m.M
	v313 = F_wangHash64(m, v297+v311)
	mBase = m.M
	*(*int64)(unsafe.Add(mBase, uint32(v263)+24)) = v313
	v315 = *(*int32)(unsafe.Add(mBase, uint32(v263)))
	v316 = v315
	goto L83
L85:
	;
	v293 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v289)+24)))
	v295 = v293 + int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v289)+24)) = uint16(v295)
	v316 = v289
	goto L83
L86:
	;
	v323 = v317 + int32(-1)
	goto L80
L87:
	;
	v320 = *(*int32)(unsafe.Add(mBase, uint32(v263)+4))
	v323 = v320
	goto L80
L88:
	;
	v350 = int32(2)
	v355 = *(*int32)(unsafe.Add(mBase, uint32(v330+v348<<(uint(v350)%32)+int32(4))))
	v279 = v355 + v349<<(uint(v350)%32)
	v280 = int32(1)
	goto L74
L89:
	;
	v339 = v327
	goto L91
L90:
	;
	v339 = v324 << (uint(v335) % 32)
	goto L91
L91:
	;
	if v325 < v339 {
		v348 = v331
		v349 = v325
		goto L88
	} else {
		goto L92
	}
L92:
	;
	if v331 != 0 {
		v368 = v327
		goto L77
	} else {
		goto L93
	}
L93:
	;
	v341 = *(*int32)(unsafe.Add(mBase, uint32(v330)+20))
	if v341 == int32(-1) {
		v368 = v327
		goto L77
	} else {
		goto L94
	}
L94:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v263)+4)) = int64(4294967296)
	v348 = int32(1)
	v349 = int32(0)
	goto L88
L95:
	;
	v364 = *(*int32)(unsafe.Add(mBase, uint32(v360)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v272))) = v364
	v368 = v360
	goto L77
L96:
	;
	v375 = v368
	goto L97
L97:
	;
	v379 = *(*int32)(unsafe.Add(mBase, uint32(v375)+8))
	goto L100
L98:
	;
	goto L68
L99:
	;
	v395 = v263 + int32(20)
	v396 = *(*int32)(unsafe.Add(mBase, uint32(v263)+16))
	if v396 != 0 {
		goto L104
	} else {
		goto L105
	}
L100:
	;
	v380 = *(*int64)(unsafe.Add(mBase, uint32(v379)+32))
	v382 = *(*int64)(unsafe.Add(mBase, _c_F_sentinelForceHelloUpdateForPrimary[0]))
	if v380 <= v382 {
		goto L99
	} else {
		goto L101
	}
L101:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v379)+32)) = v380 + (v382 ^ int64(-1))
	goto L99
L102:
	;
	if v491 != 0 {
		v375 = v491
		goto L97
	} else {
		goto L128
	}
L103:
	;
	v402 = v395
	v403 = v399
	goto L106
L104:
	;
	v399 = int32(1)
	goto L103
L105:
	;
	v399 = int32(0)
	goto L103
L106:
	;
	switch v403 {
	case 0:
		goto L111
	default:
		goto L110
	}
L108:
	;
	v403 = int32(0)
	goto L106
L109:
	;
	goto L102
L110:
	;
	v483 = *(*int32)(unsafe.Add(mBase, uint32(v402)))
	*(*int32)(unsafe.Add(mBase, uint32(v263)+16)) = v483
	if v483 == int32(0) {
		goto L108
	} else {
		goto L127
	}
L111:
	;
	v407 = *(*int32)(unsafe.Add(mBase, uint32(v263)+4))
	if v407 != int32(-1) {
		v446 = v407
		goto L112
	} else {
		goto L113
	}
L112:
	;
	v447 = int32(1)
	v448 = v446 + v447
	*(*int32)(unsafe.Add(mBase, uint32(v263)+4)) = v448
	v450 = int32(0)
	v453 = *(*int32)(unsafe.Add(mBase, uint32(v263)))
	v454 = *(*int32)(unsafe.Add(mBase, uint32(v263)+8))
	v458 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v453+v454+int32(26)))))
	if v458 == int32(255) {
		goto L121
	} else {
		goto L122
	}
L113:
	;
	v411 = *(*int32)(unsafe.Add(mBase, uint32(v263)+8))
	if v411 != 0 {
		v446 = int32(-1)
		goto L112
	} else {
		goto L114
	}
L114:
	;
	v412 = *(*int32)(unsafe.Add(mBase, uint32(v263)))
	v413 = *(*int32)(unsafe.Add(mBase, uint32(v263)+12))
	if v413 == int32(0) {
		goto L116
	} else {
		goto L117
	}
L115:
	;
	v440 = *(*int32)(unsafe.Add(mBase, uint32(v439)+20))
	if v440 != int32(-1) {
		goto L118
	} else {
		goto L119
	}
L116:
	;
	v420 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v412)+16)))
	v421 = int64(*(*int8)(unsafe.Add(mBase, uint32(v412)+27)))
	v422 = int64(*(*int32)(unsafe.Add(mBase, uint32(v412)+8)))
	v423 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v412)+12)))
	v424 = int64(*(*int8)(unsafe.Add(mBase, uint32(v412)+26)))
	v425 = int64(*(*int32)(unsafe.Add(mBase, uint32(v412)+4)))
	v426 = F_wangHash64(m, v425)
	mBase = m.M
	v428 = F_wangHash64(m, v424+v426)
	mBase = m.M
	v430 = F_wangHash64(m, v423+v428)
	mBase = m.M
	v432 = F_wangHash64(m, v422+v430)
	mBase = m.M
	v434 = F_wangHash64(m, v421+v432)
	mBase = m.M
	v436 = F_wangHash64(m, v420+v434)
	mBase = m.M
	*(*int64)(unsafe.Add(mBase, uint32(v263)+24)) = v436
	v438 = *(*int32)(unsafe.Add(mBase, uint32(v263)))
	v439 = v438
	goto L115
L117:
	;
	v416 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v412)+24)))
	v418 = v416 + int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v412)+24)) = uint16(v418)
	v439 = v412
	goto L115
L118:
	;
	v446 = v440 + int32(-1)
	goto L112
L119:
	;
	v443 = *(*int32)(unsafe.Add(mBase, uint32(v263)+4))
	v446 = v443
	goto L112
L120:
	;
	v473 = int32(2)
	v478 = *(*int32)(unsafe.Add(mBase, uint32(v453+v471<<(uint(v473)%32)+int32(4))))
	v402 = v478 + v472<<(uint(v473)%32)
	v403 = int32(1)
	goto L106
L121:
	;
	v462 = v450
	goto L123
L122:
	;
	v462 = v447 << (uint(v458) % 32)
	goto L123
L123:
	;
	if v448 < v462 {
		v471 = v454
		v472 = v448
		goto L120
	} else {
		goto L124
	}
L124:
	;
	if v454 != 0 {
		v491 = v450
		goto L109
	} else {
		goto L125
	}
L125:
	;
	v464 = *(*int32)(unsafe.Add(mBase, uint32(v453)+20))
	if v464 == int32(-1) {
		v491 = v450
		goto L109
	} else {
		goto L126
	}
L126:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v263)+4)) = int64(4294967296)
	v471 = int32(1)
	v472 = int32(0)
	goto L120
L127:
	;
	v487 = *(*int32)(unsafe.Add(mBase, uint32(v483)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v395))) = v487
	v491 = v483
	goto L109
L128:
	;
	goto L98
L129:
	;
	v504 = int32(0)
	goto L1
}
func F_sentinelGenerateInitialMonitorEvents(m *base.Module) {
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
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v131 int32
	_ = v131
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v143 int32
	_ = v143
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v151 int32
	_ = v151
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v160 int32
	_ = v160
	var v162 int32
	_ = v162
	var v164 int64
	_ = v164
	var v165 int64
	_ = v165
	var v166 int64
	_ = v166
	var v167 int64
	_ = v167
	var v168 int64
	_ = v168
	var v169 int64
	_ = v169
	var v170 int64
	_ = v170
	var v172 int64
	_ = v172
	var v174 int64
	_ = v174
	var v176 int64
	_ = v176
	var v178 int64
	_ = v178
	var v180 int64
	_ = v180
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v187 int32
	_ = v187
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v194 int32
	_ = v194
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v202 int32
	_ = v202
	var v206 int32
	_ = v206
	var v208 int32
	_ = v208
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v222 int32
	_ = v222
	var v227 int32
	_ = v227
	var v231 int32
	_ = v231
	var v235 int32
	_ = v235
	var v243 int32
	_ = v243
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	v9 = *(*int32)(unsafe.Add(mBase, _c_F_sentinelGenerateInitialMonitorEvents[0]))
	v10 = F_dictGetIterator(m, v9)
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
	v243 = m.ExcPending
	if v243 != 0 {
		goto L2
	} else {
		goto L62
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
	v123 = v115
	goto L31
L31:
	;
	v124 = *(*int32)(unsafe.Add(mBase, uint32(v123)+8))
	goto L33
L32:
	;
	goto L1
L33:
	;
	v125 = *(*int32)(unsafe.Add(mBase, uint32(v124)+152))
	*(*int32)(unsafe.Add(mBase, uint32(v6))) = v125
	F_sentinelEvent(m, int32(3), int32(_a_F_sentinelGenerateInitialMonitorEvents_0), v124, int32(_a_F_sentinelGenerateInitialMonitorEvents_1), v6)
	mBase = m.M
	v131 = m.ExcPending
	if v131 != 0 {
		goto L2
	} else {
		goto L34
	}
L34:
	;
	v139 = v10 + int32(20)
	v140 = *(*int32)(unsafe.Add(mBase, uint32(v10)+16))
	if v140 != 0 {
		goto L37
	} else {
		goto L38
	}
L35:
	;
	if v235 != 0 {
		v123 = v235
		goto L31
	} else {
		goto L61
	}
L36:
	;
	v146 = v139
	v147 = v143
	goto L39
L37:
	;
	v143 = int32(1)
	goto L36
L38:
	;
	v143 = int32(0)
	goto L36
L39:
	;
	switch v147 {
	case 0:
		goto L44
	default:
		goto L43
	}
L41:
	;
	v147 = int32(0)
	goto L39
L42:
	;
	goto L35
L43:
	;
	v227 = *(*int32)(unsafe.Add(mBase, uint32(v146)))
	*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = v227
	if v227 == int32(0) {
		goto L41
	} else {
		goto L60
	}
L44:
	;
	v151 = *(*int32)(unsafe.Add(mBase, uint32(v10)+4))
	if v151 != int32(-1) {
		v190 = v151
		goto L45
	} else {
		goto L46
	}
L45:
	;
	v191 = int32(1)
	v192 = v190 + v191
	*(*int32)(unsafe.Add(mBase, uint32(v10)+4)) = v192
	v194 = int32(0)
	v197 = *(*int32)(unsafe.Add(mBase, uint32(v10)))
	v198 = *(*int32)(unsafe.Add(mBase, uint32(v10)+8))
	v202 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v197+v198+int32(26)))))
	if v202 == int32(255) {
		goto L54
	} else {
		goto L55
	}
L46:
	;
	v155 = *(*int32)(unsafe.Add(mBase, uint32(v10)+8))
	if v155 != 0 {
		v190 = int32(-1)
		goto L45
	} else {
		goto L47
	}
L47:
	;
	v156 = *(*int32)(unsafe.Add(mBase, uint32(v10)))
	v157 = *(*int32)(unsafe.Add(mBase, uint32(v10)+12))
	if v157 == int32(0) {
		goto L49
	} else {
		goto L50
	}
L48:
	;
	v184 = *(*int32)(unsafe.Add(mBase, uint32(v183)+20))
	if v184 != int32(-1) {
		goto L51
	} else {
		goto L52
	}
L49:
	;
	v164 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v156)+16)))
	v165 = int64(*(*int8)(unsafe.Add(mBase, uint32(v156)+27)))
	v166 = int64(*(*int32)(unsafe.Add(mBase, uint32(v156)+8)))
	v167 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v156)+12)))
	v168 = int64(*(*int8)(unsafe.Add(mBase, uint32(v156)+26)))
	v169 = int64(*(*int32)(unsafe.Add(mBase, uint32(v156)+4)))
	v170 = F_wangHash64(m, v169)
	mBase = m.M
	v172 = F_wangHash64(m, v168+v170)
	mBase = m.M
	v174 = F_wangHash64(m, v167+v172)
	mBase = m.M
	v176 = F_wangHash64(m, v166+v174)
	mBase = m.M
	v178 = F_wangHash64(m, v165+v176)
	mBase = m.M
	v180 = F_wangHash64(m, v164+v178)
	mBase = m.M
	*(*int64)(unsafe.Add(mBase, uint32(v10)+24)) = v180
	v182 = *(*int32)(unsafe.Add(mBase, uint32(v10)))
	v183 = v182
	goto L48
L50:
	;
	v160 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v156)+24)))
	v162 = v160 + int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v156)+24)) = uint16(v162)
	v183 = v156
	goto L48
L51:
	;
	v190 = v184 + int32(-1)
	goto L45
L52:
	;
	v187 = *(*int32)(unsafe.Add(mBase, uint32(v10)+4))
	v190 = v187
	goto L45
L53:
	;
	v217 = int32(2)
	v222 = *(*int32)(unsafe.Add(mBase, uint32(v197+v215<<(uint(v217)%32)+int32(4))))
	v146 = v222 + v216<<(uint(v217)%32)
	v147 = int32(1)
	goto L39
L54:
	;
	v206 = v194
	goto L56
L55:
	;
	v206 = v191 << (uint(v202) % 32)
	goto L56
L56:
	;
	if v192 < v206 {
		v215 = v198
		v216 = v192
		goto L53
	} else {
		goto L57
	}
L57:
	;
	if v198 != 0 {
		v235 = v194
		goto L42
	} else {
		goto L58
	}
L58:
	;
	v208 = *(*int32)(unsafe.Add(mBase, uint32(v197)+20))
	if v208 == int32(-1) {
		v235 = v194
		goto L42
	} else {
		goto L59
	}
L59:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v10)+4)) = int64(4294967296)
	v215 = int32(1)
	v216 = int32(0)
	goto L53
L60:
	;
	v231 = *(*int32)(unsafe.Add(mBase, uint32(v227)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v139))) = v231
	v235 = v227
	goto L42
L61:
	;
	goto L32
L62:
	;
	m.G0 = v6 + int32(16)
	return
}
func F_sentinelHandleConfiguration(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
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
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	var v53 int32
	_ = v53
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	var v98 int32
	_ = v98
	var v103 int32
	_ = v103
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v110 int32
	_ = v110
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v131 int32
	_ = v131
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v139 int32
	_ = v139
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v146 int32
	_ = v146
	var v149 int32
	_ = v149
	var v155 int32
	_ = v155
	var v157 int32
	_ = v157
	var v160 int32
	_ = v160
	var v163 int32
	_ = v163
	var v174 int32
	_ = v174
	var v177 int32
	_ = v177
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
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
	var v202 int32
	_ = v202
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v223 int32
	_ = v223
	var v226 int32
	_ = v226
	var v230 int32
	_ = v230
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v248 int32
	_ = v248
	var v251 int32
	_ = v251
	var v255 int32
	_ = v255
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v259 int32
	_ = v259
	var v261 int32
	_ = v261
	var v262 int32
	_ = v262
	var v266 int32
	_ = v266
	var v269 int32
	_ = v269
	var v275 int32
	_ = v275
	var v282 int32
	_ = v282
	var v285 int32
	_ = v285
	var v287 int32
	_ = v287
	var v288 int32
	_ = v288
	var v289 int32
	_ = v289
	var v291 int32
	_ = v291
	var v295 int32
	_ = v295
	var v297 int32
	_ = v297
	var v298 int32
	_ = v298
	var v299 int32
	_ = v299
	var v301 int32
	_ = v301
	var v302 int32
	_ = v302
	var v303 int32
	_ = v303
	var v304 int32
	_ = v304
	var v309 int32
	_ = v309
	var v310 int32
	_ = v310
	var v313 int32
	_ = v313
	var v314 int32
	_ = v314
	var v317 int32
	_ = v317
	var v318 int32
	_ = v318
	var v319 int32
	_ = v319
	var v321 int32
	_ = v321
	var v322 int32
	_ = v322
	var v323 int32
	_ = v323
	var v325 int32
	_ = v325
	var v326 int32
	_ = v326
	var v327 int32
	_ = v327
	var v329 int32
	_ = v329
	var v332 int32
	_ = v332
	var v336 int32
	_ = v336
	var v341 int32
	_ = v341
	var v342 int32
	_ = v342
	var v343 int32
	_ = v343
	var v344 int32
	_ = v344
	var v350 int32
	_ = v350
	var v351 int32
	_ = v351
	var v352 int32
	_ = v352
	var v353 int32
	_ = v353
	var v354 int32
	_ = v354
	var v357 int32
	_ = v357
	var v361 int32
	_ = v361
	var v362 int32
	_ = v362
	var v363 int32
	_ = v363
	var v365 int32
	_ = v365
	var v367 int32
	_ = v367
	var v368 int32
	_ = v368
	var v372 int32
	_ = v372
	var v375 int32
	_ = v375
	var v381 int32
	_ = v381
	var v388 int32
	_ = v388
	var v391 int32
	_ = v391
	var v393 int32
	_ = v393
	var v394 int32
	_ = v394
	var v395 int32
	_ = v395
	var v397 int32
	_ = v397
	var v401 int32
	_ = v401
	var v403 int32
	_ = v403
	var v404 int32
	_ = v404
	var v405 int32
	_ = v405
	var v407 int32
	_ = v407
	var v408 int32
	_ = v408
	var v409 int32
	_ = v409
	var v410 int32
	_ = v410
	var v415 int32
	_ = v415
	var v416 int32
	_ = v416
	var v419 int32
	_ = v419
	var v420 int32
	_ = v420
	var v423 int32
	_ = v423
	var v424 int32
	_ = v424
	var v425 int32
	_ = v425
	var v428 int32
	_ = v428
	var v429 int32
	_ = v429
	var v430 int32
	_ = v430
	var v431 int32
	_ = v431
	var v433 int32
	_ = v433
	var v434 int32
	_ = v434
	var v435 int32
	_ = v435
	var v437 int32
	_ = v437
	var v440 int32
	_ = v440
	var v444 int32
	_ = v444
	var v449 int32
	_ = v449
	var v450 int32
	_ = v450
	var v451 int32
	_ = v451
	var v452 int32
	_ = v452
	var v458 int32
	_ = v458
	var v459 int32
	_ = v459
	var v460 int32
	_ = v460
	var v461 int32
	_ = v461
	var v462 int32
	_ = v462
	var v465 int32
	_ = v465
	var v469 int32
	_ = v469
	var v470 int32
	_ = v470
	var v471 int32
	_ = v471
	var v473 int32
	_ = v473
	var v475 int32
	_ = v475
	var v476 int32
	_ = v476
	var v480 int32
	_ = v480
	var v483 int32
	_ = v483
	var v489 int32
	_ = v489
	var v491 int32
	_ = v491
	var v494 int32
	_ = v494
	var v496 int32
	_ = v496
	var v497 int32
	_ = v497
	var v498 int32
	_ = v498
	var v500 int32
	_ = v500
	var v504 int32
	_ = v504
	var v506 int32
	_ = v506
	var v507 int32
	_ = v507
	var v508 int32
	_ = v508
	var v510 int32
	_ = v510
	var v511 int32
	_ = v511
	var v512 int32
	_ = v512
	var v513 int32
	_ = v513
	var v518 int32
	_ = v518
	var v519 int32
	_ = v519
	var v522 int32
	_ = v522
	var v523 int32
	_ = v523
	var v526 int32
	_ = v526
	var v527 int32
	_ = v527
	var v528 int32
	_ = v528
	var v530 int32
	_ = v530
	var v531 int32
	_ = v531
	var v532 int32
	_ = v532
	var v534 int32
	_ = v534
	var v535 int32
	_ = v535
	var v536 int32
	_ = v536
	var v538 int32
	_ = v538
	var v541 int32
	_ = v541
	var v543 int32
	_ = v543
	var v548 int32
	_ = v548
	var v549 int32
	_ = v549
	var v550 int32
	_ = v550
	var v552 int32
	_ = v552
	var v555 int32
	_ = v555
	var v557 int32
	_ = v557
	var v558 int32
	_ = v558
	var v559 int32
	_ = v559
	var v561 int32
	_ = v561
	var v565 int32
	_ = v565
	var v567 int32
	_ = v567
	var v568 int32
	_ = v568
	var v569 int32
	_ = v569
	var v571 int32
	_ = v571
	var v572 int32
	_ = v572
	var v573 int32
	_ = v573
	var v574 int32
	_ = v574
	var v579 int32
	_ = v579
	var v580 int32
	_ = v580
	var v583 int32
	_ = v583
	var v584 int32
	_ = v584
	var v587 int32
	_ = v587
	var v588 int32
	_ = v588
	var v589 int32
	_ = v589
	var v592 int32
	_ = v592
	var v593 int32
	_ = v593
	var v594 int32
	_ = v594
	var v595 int32
	_ = v595
	var v598 int32
	_ = v598
	var v600 int32
	_ = v600
	var v605 int32
	_ = v605
	var v606 int32
	_ = v606
	var v607 int32
	_ = v607
	var v609 int32
	_ = v609
	var v612 int32
	_ = v612
	var v614 int32
	_ = v614
	var v615 int32
	_ = v615
	var v616 int32
	_ = v616
	var v618 int32
	_ = v618
	var v622 int32
	_ = v622
	var v624 int32
	_ = v624
	var v625 int32
	_ = v625
	var v626 int32
	_ = v626
	var v628 int32
	_ = v628
	var v629 int32
	_ = v629
	var v630 int32
	_ = v630
	var v631 int32
	_ = v631
	var v636 int32
	_ = v636
	var v637 int32
	_ = v637
	var v640 int32
	_ = v640
	var v641 int32
	_ = v641
	var v644 int32
	_ = v644
	var v645 int32
	_ = v645
	var v646 int32
	_ = v646
	var v648 int32
	_ = v648
	var v649 int32
	_ = v649
	var v650 int32
	_ = v650
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
	var v662 int32
	_ = v662
	var v663 int32
	_ = v663
	var v664 int32
	_ = v664
	var v666 int32
	_ = v666
	var v670 int32
	_ = v670
	var v672 int32
	_ = v672
	var v673 int32
	_ = v673
	var v674 int32
	_ = v674
	var v676 int32
	_ = v676
	var v677 int32
	_ = v677
	var v678 int32
	_ = v678
	var v679 int32
	_ = v679
	var v684 int32
	_ = v684
	var v685 int32
	_ = v685
	var v688 int32
	_ = v688
	var v689 int32
	_ = v689
	var v692 int32
	_ = v692
	var v693 int32
	_ = v693
	var v694 int32
	_ = v694
	var v698 int32
	_ = v698
	var v699 int32
	_ = v699
	var v700 int32
	_ = v700
	var v703 int32
	_ = v703
	var v704 int32
	_ = v704
	var v705 int32
	_ = v705
	var v707 int32
	_ = v707
	var v710 int32
	_ = v710
	var v712 int32
	_ = v712
	var v713 int32
	_ = v713
	var v714 int32
	_ = v714
	var v716 int32
	_ = v716
	var v720 int32
	_ = v720
	var v722 int32
	_ = v722
	var v723 int32
	_ = v723
	var v724 int32
	_ = v724
	var v726 int32
	_ = v726
	var v727 int32
	_ = v727
	var v728 int32
	_ = v728
	var v729 int32
	_ = v729
	var v734 int32
	_ = v734
	var v735 int32
	_ = v735
	var v738 int32
	_ = v738
	var v739 int32
	_ = v739
	var v742 int32
	_ = v742
	var v743 int32
	_ = v743
	var v744 int32
	_ = v744
	var v747 int32
	_ = v747
	var v748 int32
	_ = v748
	var v752 int64
	_ = v752
	var v754 int64
	_ = v754
	var v758 int32
	_ = v758
	var v761 int32
	_ = v761
	var v763 int32
	_ = v763
	var v764 int32
	_ = v764
	var v765 int32
	_ = v765
	var v767 int32
	_ = v767
	var v771 int32
	_ = v771
	var v773 int32
	_ = v773
	var v774 int32
	_ = v774
	var v775 int32
	_ = v775
	var v777 int32
	_ = v777
	var v778 int32
	_ = v778
	var v779 int32
	_ = v779
	var v780 int32
	_ = v780
	var v785 int32
	_ = v785
	var v786 int32
	_ = v786
	var v789 int32
	_ = v789
	var v790 int32
	_ = v790
	var v793 int32
	_ = v793
	var v794 int32
	_ = v794
	var v795 int32
	_ = v795
	var v797 int32
	_ = v797
	var v805 int32
	_ = v805
	var v808 int32
	_ = v808
	var v812 int32
	_ = v812
	var v817 int32
	_ = v817
	var v819 int32
	_ = v819
	var v823 int32
	_ = v823
	var v829 int32
	_ = v829
	var v832 int32
	_ = v832
	var v838 int32
	_ = v838
	var v842 int32
	_ = v842
	var v844 int32
	_ = v844
	var v852 int32
	_ = v852
	var v857 int32
	_ = v857
	var v858 int64
	_ = v858
	var v863 int64
	_ = v863
	var v868 int64
	_ = v868
	var v873 int64
	_ = v873
	var v878 int64
	_ = v878
	var v880 int32
	_ = v880
	var v883 int32
	_ = v883
	var v885 int32
	_ = v885
	var v886 int32
	_ = v886
	var v887 int32
	_ = v887
	var v889 int32
	_ = v889
	var v893 int32
	_ = v893
	var v895 int32
	_ = v895
	var v896 int32
	_ = v896
	var v897 int32
	_ = v897
	var v899 int32
	_ = v899
	var v900 int32
	_ = v900
	var v901 int32
	_ = v901
	var v902 int32
	_ = v902
	var v907 int32
	_ = v907
	var v908 int32
	_ = v908
	var v911 int32
	_ = v911
	var v912 int32
	_ = v912
	var v915 int32
	_ = v915
	var v916 int32
	_ = v916
	var v917 int32
	_ = v917
	var v920 int32
	_ = v920
	var v921 int32
	_ = v921
	var v922 int32
	_ = v922
	var v923 int32
	_ = v923
	var v926 int32
	_ = v926
	var v930 int64
	_ = v930
	var v933 int64
	_ = v933
	var v937 int32
	_ = v937
	var v940 int32
	_ = v940
	var v942 int32
	_ = v942
	var v943 int32
	_ = v943
	var v944 int32
	_ = v944
	var v946 int32
	_ = v946
	var v950 int32
	_ = v950
	var v952 int32
	_ = v952
	var v953 int32
	_ = v953
	var v954 int32
	_ = v954
	var v956 int32
	_ = v956
	var v957 int32
	_ = v957
	var v958 int32
	_ = v958
	var v959 int32
	_ = v959
	var v964 int32
	_ = v964
	var v965 int32
	_ = v965
	var v968 int32
	_ = v968
	var v969 int32
	_ = v969
	var v972 int32
	_ = v972
	var v973 int32
	_ = v973
	var v974 int32
	_ = v974
	var v976 int32
	_ = v976
	var v977 int32
	_ = v977
	var v978 int32
	_ = v978
	var v981 int32
	_ = v981
	var v985 int64
	_ = v985
	var v987 int32
	_ = v987
	var v990 int32
	_ = v990
	var v992 int32
	_ = v992
	var v993 int32
	_ = v993
	var v994 int32
	_ = v994
	var v996 int32
	_ = v996
	var v1000 int32
	_ = v1000
	var v1002 int32
	_ = v1002
	var v1003 int32
	_ = v1003
	var v1004 int32
	_ = v1004
	var v1006 int32
	_ = v1006
	var v1007 int32
	_ = v1007
	var v1008 int32
	_ = v1008
	var v1009 int32
	_ = v1009
	var v1014 int32
	_ = v1014
	var v1015 int32
	_ = v1015
	var v1018 int32
	_ = v1018
	var v1019 int32
	_ = v1019
	var v1022 int32
	_ = v1022
	var v1023 int32
	_ = v1023
	var v1024 int32
	_ = v1024
	var v1028 int32
	_ = v1028
	var v1031 int32
	_ = v1031
	var v1033 int32
	_ = v1033
	var v1034 int32
	_ = v1034
	var v1035 int32
	_ = v1035
	var v1037 int32
	_ = v1037
	var v1041 int32
	_ = v1041
	var v1043 int32
	_ = v1043
	var v1044 int32
	_ = v1044
	var v1045 int32
	_ = v1045
	var v1047 int32
	_ = v1047
	var v1048 int32
	_ = v1048
	var v1049 int32
	_ = v1049
	var v1050 int32
	_ = v1050
	var v1055 int32
	_ = v1055
	var v1056 int32
	_ = v1056
	var v1059 int32
	_ = v1059
	var v1060 int32
	_ = v1060
	var v1063 int32
	_ = v1063
	var v1064 int32
	_ = v1064
	var v1065 int32
	_ = v1065
	var v1074 int32
	_ = v1074
	var v1075 int32
	_ = v1075
	var v1076 int32
	_ = v1076
	var v1081 int32
	_ = v1081
	var v1082 int32
	_ = v1082
	var v1086 int32
	_ = v1086
	var v1091 int32
	_ = v1091
	var v1092 int32
	_ = v1092
	var v1093 int32
	_ = v1093
	var v1094 int32
	_ = v1094
	var v1100 int32
	_ = v1100
	var v1101 int32
	_ = v1101
	var v1102 int32
	_ = v1102
	var v1103 int32
	_ = v1103
	var v1104 int32
	_ = v1104
	var v1107 int32
	_ = v1107
	var v1111 int32
	_ = v1111
	var v1112 int32
	_ = v1112
	var v1113 int32
	_ = v1113
	var v1115 int32
	_ = v1115
	var v1117 int32
	_ = v1117
	var v1118 int32
	_ = v1118
	var v1122 int32
	_ = v1122
	var v1125 int32
	_ = v1125
	var v1131 int32
	_ = v1131
	var v1132 int32
	_ = v1132
	var v1133 int32
	_ = v1133
	var v1134 int32
	_ = v1134
	var v1137 int32
	_ = v1137
	var v1149 int32
	_ = v1149
	var v1151 int32
	_ = v1151
	var v1154 int32
	_ = v1154
	var v1156 int32
	_ = v1156
	var v1157 int32
	_ = v1157
	var v1158 int32
	_ = v1158
	var v1160 int32
	_ = v1160
	var v1164 int32
	_ = v1164
	var v1166 int32
	_ = v1166
	var v1167 int32
	_ = v1167
	var v1168 int32
	_ = v1168
	var v1170 int32
	_ = v1170
	var v1171 int32
	_ = v1171
	var v1172 int32
	_ = v1172
	var v1173 int32
	_ = v1173
	var v1178 int32
	_ = v1178
	var v1179 int32
	_ = v1179
	var v1182 int32
	_ = v1182
	var v1183 int32
	_ = v1183
	var v1186 int32
	_ = v1186
	var v1187 int32
	_ = v1187
	var v1188 int32
	_ = v1188
	var v1196 int32
	_ = v1196
	var v1197 int32
	_ = v1197
	var v1198 int32
	_ = v1198
	var v1201 int32
	_ = v1201
	var v1203 int32
	_ = v1203
	var v1204 int32
	_ = v1204
	var v1208 int32
	_ = v1208
	var v1213 int32
	_ = v1213
	var v1214 int32
	_ = v1214
	var v1215 int32
	_ = v1215
	var v1216 int32
	_ = v1216
	var v1222 int32
	_ = v1222
	var v1223 int32
	_ = v1223
	var v1224 int32
	_ = v1224
	var v1225 int32
	_ = v1225
	var v1226 int32
	_ = v1226
	var v1229 int32
	_ = v1229
	var v1233 int32
	_ = v1233
	var v1234 int32
	_ = v1234
	var v1235 int32
	_ = v1235
	var v1237 int32
	_ = v1237
	var v1239 int32
	_ = v1239
	var v1240 int32
	_ = v1240
	var v1244 int32
	_ = v1244
	var v1247 int32
	_ = v1247
	var v1253 int32
	_ = v1253
	var v1254 int32
	_ = v1254
	var v1255 int32
	_ = v1255
	var v1256 int32
	_ = v1256
	var v1259 int32
	_ = v1259
	var v1270 int32
	_ = v1270
	var v1271 int32
	_ = v1271
	var v1272 int32
	_ = v1272
	var v1274 int32
	_ = v1274
	var v1275 int32
	_ = v1275
	var v1276 int32
	_ = v1276
	var v1279 int32
	_ = v1279
	var v1281 int32
	_ = v1281
	var v1282 int32
	_ = v1282
	var v1283 int32
	_ = v1283
	var v1285 int32
	_ = v1285
	var v1289 int32
	_ = v1289
	var v1291 int32
	_ = v1291
	var v1292 int32
	_ = v1292
	var v1293 int32
	_ = v1293
	var v1295 int32
	_ = v1295
	var v1296 int32
	_ = v1296
	var v1297 int32
	_ = v1297
	var v1298 int32
	_ = v1298
	var v1303 int32
	_ = v1303
	var v1304 int32
	_ = v1304
	var v1307 int32
	_ = v1307
	var v1308 int32
	_ = v1308
	var v1311 int32
	_ = v1311
	var v1312 int32
	_ = v1312
	var v1313 int32
	_ = v1313
	var v1317 int32
	_ = v1317
	var v1318 int32
	_ = v1318
	var v1319 int32
	_ = v1319
	var v1322 int32
	_ = v1322
	var v1323 int32
	_ = v1323
	var v1324 int32
	_ = v1324
	var v1325 int32
	_ = v1325
	var v1326 int32
	_ = v1326
	var v1327 int32
	_ = v1327
	var v1328 int32
	_ = v1328
	var v1329 int32
	_ = v1329
	var v1330 int32
	_ = v1330
	var v1334 int32
	_ = v1334
	var v1336 int32
	_ = v1336
	var v1339 int32
	_ = v1339
	var v1342 int32
	_ = v1342
	var v1344 int32
	_ = v1344
	var v1345 int32
	_ = v1345
	var v1346 int32
	_ = v1346
	var v1348 int32
	_ = v1348
	var v1352 int32
	_ = v1352
	var v1354 int32
	_ = v1354
	var v1355 int32
	_ = v1355
	var v1356 int32
	_ = v1356
	var v1358 int32
	_ = v1358
	var v1359 int32
	_ = v1359
	var v1360 int32
	_ = v1360
	var v1361 int32
	_ = v1361
	var v1366 int32
	_ = v1366
	var v1367 int32
	_ = v1367
	var v1370 int32
	_ = v1370
	var v1371 int32
	_ = v1371
	var v1374 int32
	_ = v1374
	var v1375 int32
	_ = v1375
	var v1376 int32
	_ = v1376
	var v1379 int32
	_ = v1379
	var v1380 int32
	_ = v1380
	var v1381 int32
	_ = v1381
	var v1385 int32
	_ = v1385
	var v1386 int32
	_ = v1386
	var v1388 int32
	_ = v1388
	var v1391 int32
	_ = v1391
	var v1393 int32
	_ = v1393
	var v1394 int32
	_ = v1394
	var v1395 int32
	_ = v1395
	var v1397 int32
	_ = v1397
	var v1401 int32
	_ = v1401
	var v1403 int32
	_ = v1403
	var v1404 int32
	_ = v1404
	var v1405 int32
	_ = v1405
	var v1407 int32
	_ = v1407
	var v1408 int32
	_ = v1408
	var v1409 int32
	_ = v1409
	var v1410 int32
	_ = v1410
	var v1415 int32
	_ = v1415
	var v1416 int32
	_ = v1416
	var v1419 int32
	_ = v1419
	var v1420 int32
	_ = v1420
	var v1423 int32
	_ = v1423
	var v1424 int32
	_ = v1424
	var v1425 int32
	_ = v1425
	var v1428 int32
	_ = v1428
	var v1432 int32
	_ = v1432
	var v1437 int32
	_ = v1437
	var v1438 int32
	_ = v1438
	var v1439 int32
	_ = v1439
	var v1440 int32
	_ = v1440
	var v1446 int32
	_ = v1446
	var v1447 int32
	_ = v1447
	var v1448 int32
	_ = v1448
	var v1449 int32
	_ = v1449
	var v1450 int32
	_ = v1450
	var v1453 int32
	_ = v1453
	var v1457 int32
	_ = v1457
	var v1458 int32
	_ = v1458
	var v1459 int32
	_ = v1459
	var v1461 int32
	_ = v1461
	var v1463 int32
	_ = v1463
	var v1464 int32
	_ = v1464
	var v1468 int32
	_ = v1468
	var v1471 int32
	_ = v1471
	var v1477 int32
	_ = v1477
	var v1479 int32
	_ = v1479
	var v1482 int32
	_ = v1482
	var v1484 int32
	_ = v1484
	var v1485 int32
	_ = v1485
	var v1486 int32
	_ = v1486
	var v1488 int32
	_ = v1488
	var v1492 int32
	_ = v1492
	var v1494 int32
	_ = v1494
	var v1495 int32
	_ = v1495
	var v1496 int32
	_ = v1496
	var v1498 int32
	_ = v1498
	var v1499 int32
	_ = v1499
	var v1500 int32
	_ = v1500
	var v1501 int32
	_ = v1501
	var v1506 int32
	_ = v1506
	var v1507 int32
	_ = v1507
	var v1510 int32
	_ = v1510
	var v1511 int32
	_ = v1511
	var v1514 int32
	_ = v1514
	var v1515 int32
	_ = v1515
	var v1516 int32
	_ = v1516
	var v1519 int32
	_ = v1519
	var v1521 int32
	_ = v1521
	var v1523 int32
	_ = v1523
	var v1528 int32
	_ = v1528
	var v1529 int32
	_ = v1529
	var v1530 int32
	_ = v1530
	var v1536 int32
	_ = v1536
	var v1539 int32
	_ = v1539
	var v1541 int32
	_ = v1541
	var v1542 int32
	_ = v1542
	var v1543 int32
	_ = v1543
	var v1545 int32
	_ = v1545
	var v1549 int32
	_ = v1549
	var v1551 int32
	_ = v1551
	var v1552 int32
	_ = v1552
	var v1553 int32
	_ = v1553
	var v1555 int32
	_ = v1555
	var v1556 int32
	_ = v1556
	var v1557 int32
	_ = v1557
	var v1558 int32
	_ = v1558
	var v1563 int32
	_ = v1563
	var v1564 int32
	_ = v1564
	var v1567 int32
	_ = v1567
	var v1568 int32
	_ = v1568
	var v1571 int32
	_ = v1571
	var v1572 int32
	_ = v1572
	var v1573 int32
	_ = v1573
	var v1575 int32
	_ = v1575
	var v1576 int32
	_ = v1576
	var v1580 int32
	_ = v1580
	var v1581 int32
	_ = v1581
	var v1583 int32
	_ = v1583
	var v1586 int32
	_ = v1586
	var v1588 int32
	_ = v1588
	var v1589 int32
	_ = v1589
	var v1590 int32
	_ = v1590
	var v1592 int32
	_ = v1592
	var v1596 int32
	_ = v1596
	var v1598 int32
	_ = v1598
	var v1599 int32
	_ = v1599
	var v1600 int32
	_ = v1600
	var v1602 int32
	_ = v1602
	var v1603 int32
	_ = v1603
	var v1604 int32
	_ = v1604
	var v1605 int32
	_ = v1605
	var v1610 int32
	_ = v1610
	var v1611 int32
	_ = v1611
	var v1614 int32
	_ = v1614
	var v1615 int32
	_ = v1615
	var v1618 int32
	_ = v1618
	var v1619 int32
	_ = v1619
	var v1620 int32
	_ = v1620
	var v1623 int32
	_ = v1623
	var v1624 int32
	_ = v1624
	var v1625 int32
	_ = v1625
	var v1629 int32
	_ = v1629
	var v1630 int32
	_ = v1630
	var v1632 int32
	_ = v1632
	var v1635 int32
	_ = v1635
	var v1637 int32
	_ = v1637
	var v1638 int32
	_ = v1638
	var v1639 int32
	_ = v1639
	var v1641 int32
	_ = v1641
	var v1645 int32
	_ = v1645
	var v1647 int32
	_ = v1647
	var v1648 int32
	_ = v1648
	var v1649 int32
	_ = v1649
	var v1651 int32
	_ = v1651
	var v1652 int32
	_ = v1652
	var v1653 int32
	_ = v1653
	var v1654 int32
	_ = v1654
	var v1659 int32
	_ = v1659
	var v1660 int32
	_ = v1660
	var v1663 int32
	_ = v1663
	var v1664 int32
	_ = v1664
	var v1667 int32
	_ = v1667
	var v1668 int32
	_ = v1668
	var v1669 int32
	_ = v1669
	var v1672 int32
	_ = v1672
	var v1674 int32
	_ = v1674
	var v1679 int32
	_ = v1679
	var v1680 int32
	_ = v1680
	var v1681 int32
	_ = v1681
	var v1687 int32
	_ = v1687
	var v1690 int32
	_ = v1690
	var v1692 int32
	_ = v1692
	var v1693 int32
	_ = v1693
	var v1694 int32
	_ = v1694
	var v1696 int32
	_ = v1696
	var v1700 int32
	_ = v1700
	var v1702 int32
	_ = v1702
	var v1703 int32
	_ = v1703
	var v1704 int32
	_ = v1704
	var v1706 int32
	_ = v1706
	var v1707 int32
	_ = v1707
	var v1708 int32
	_ = v1708
	var v1709 int32
	_ = v1709
	var v1714 int32
	_ = v1714
	var v1715 int32
	_ = v1715
	var v1718 int32
	_ = v1718
	var v1719 int32
	_ = v1719
	var v1722 int32
	_ = v1722
	var v1723 int32
	_ = v1723
	var v1724 int32
	_ = v1724
	var v1729 int32
	_ = v1729
	var v1731 int32
	_ = v1731
	var v1736 int32
	_ = v1736
	var v1737 int32
	_ = v1737
	var v1738 int32
	_ = v1738
	var v1744 int32
	_ = v1744
	var v1747 int32
	_ = v1747
	var v1749 int32
	_ = v1749
	var v1750 int32
	_ = v1750
	var v1751 int32
	_ = v1751
	var v1753 int32
	_ = v1753
	var v1757 int32
	_ = v1757
	var v1759 int32
	_ = v1759
	var v1760 int32
	_ = v1760
	var v1761 int32
	_ = v1761
	var v1763 int32
	_ = v1763
	var v1764 int32
	_ = v1764
	var v1765 int32
	_ = v1765
	var v1766 int32
	_ = v1766
	var v1771 int32
	_ = v1771
	var v1772 int32
	_ = v1772
	var v1775 int32
	_ = v1775
	var v1776 int32
	_ = v1776
	var v1779 int32
	_ = v1779
	var v1780 int32
	_ = v1780
	var v1781 int32
	_ = v1781
	var v1785 int32
	_ = v1785
	var v1786 int32
	_ = v1786
	var v1789 int32
	_ = v1789
	var v1791 int32
	_ = v1791
	var v1792 int32
	_ = v1792
	var v1793 int32
	_ = v1793
	var v1795 int32
	_ = v1795
	var v1799 int32
	_ = v1799
	var v1801 int32
	_ = v1801
	var v1802 int32
	_ = v1802
	var v1803 int32
	_ = v1803
	var v1805 int32
	_ = v1805
	var v1806 int32
	_ = v1806
	var v1807 int32
	_ = v1807
	var v1808 int32
	_ = v1808
	var v1813 int32
	_ = v1813
	var v1814 int32
	_ = v1814
	var v1817 int32
	_ = v1817
	var v1818 int32
	_ = v1818
	var v1821 int32
	_ = v1821
	var v1822 int32
	_ = v1822
	var v1823 int32
	_ = v1823
	var v1833 int32
	_ = v1833
	var v1834 int32
	_ = v1834
	var v1835 int32
	_ = v1835
	var v1838 int32
	_ = v1838
	var v1842 int32
	_ = v1842
	var v1847 int32
	_ = v1847
	var v1848 int32
	_ = v1848
	var v1849 int32
	_ = v1849
	var v1850 int32
	_ = v1850
	var v1856 int32
	_ = v1856
	var v1857 int32
	_ = v1857
	var v1858 int32
	_ = v1858
	var v1859 int32
	_ = v1859
	var v1860 int32
	_ = v1860
	var v1863 int32
	_ = v1863
	var v1867 int32
	_ = v1867
	var v1868 int32
	_ = v1868
	var v1869 int32
	_ = v1869
	var v1871 int32
	_ = v1871
	var v1873 int32
	_ = v1873
	var v1874 int32
	_ = v1874
	var v1878 int32
	_ = v1878
	var v1881 int32
	_ = v1881
	var v1887 int32
	_ = v1887
	var v1904 int32
	_ = v1904
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v8 = int32(_a_F_sentinelHandleConfiguration_0)
	v11 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7))))
	if v11 != 0 {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	if l1 != int32(5) {
		goto L15
	} else {
		goto L16
	}
L2:
	;
	v43 = F_tolower(m, v39)
	mBase = m.M
	v44 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v40))))
	v45 = F_tolower(m, v44)
	mBase = m.M
	goto L1
L3:
	;
	v13 = v7
	v14 = v8
	v15 = v11
	goto L6
L4:
	;
	v39 = int32(0)
	v40 = v8
	goto L2
L5:
	;
	v39 = v36 & int32(255)
	v40 = v35
	goto L2
L6:
	;
	v17 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14))))
	if v17 == int32(0) {
		v35 = v14
		v36 = v15
		goto L5
	} else {
		goto L8
	}
L7:
	;
	v35 = v29
	v36 = int32(0)
	goto L5
L8:
	;
	v21 = v15 & int32(255)
	if v21 == v17 {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	v28 = int32(1)
	v29 = v14 + v28
	v30 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+1)))
	if v30 != 0 {
		v13 = v13 + v28
		v14 = v29
		v15 = v30
		goto L6
	} else {
		goto L12
	}
L10:
	;
	v23 = F_tolower(m, v21)
	mBase = m.M
	v24 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14))))
	v25 = F_tolower(m, v24)
	mBase = m.M
	if v23 == v25 {
		goto L9
	} else {
		goto L11
	}
L11:
	;
	v27 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13))))
	v35 = v14
	v36 = v27
	goto L5
L12:
	;
	goto L7
L13:
	;
	return v1904
L14:
	;
	v1904 = int32(0)
	goto L13
L15:
	;
	v174 = int32(_a_F_sentinelHandleConfiguration_1)
	v177 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7))))
	if v177 != 0 {
		goto L60
	} else {
		goto L61
	}
L16:
	;
	if v43-v45 != 0 {
		goto L15
	} else {
		goto L17
	}
L17:
	;
	v49 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v53 = v49
	goto L20
L18:
	;
	v103 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v105 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v106 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v110 = v106
	goto L36
L19:
	;
	if int32(1) <= v98 {
		goto L18
	} else {
		goto L34
	}
L20:
	;
	v58 = v53 + int32(1)
	v59 = int32(*(*int8)(unsafe.Add(mBase, uint32(v53))))
	v60 = F___isspace_1(m, v59)
	mBase = m.M
	if v60 != 0 {
		v53 = v58
		goto L20
	} else {
		goto L22
	}
L21:
	;
	v61 = int32(1)
	switch v59&int32(255) + int32(-43) {
	case 0:
		v67 = v61
		goto L24
	default:
		v69 = v53
		v70 = v59
		v71 = v61
		goto L23
	case 2:
		goto L25
	}
L22:
	;
	goto L21
L23:
	;
	v74 = v70 + int32(-48)
	if base.Ui32(int32(9)) < base.Ui32(v74) {
		v92 = int32(0)
		goto L26
	} else {
		goto L27
	}
L24:
	;
	v68 = int32(*(*int8)(unsafe.Add(mBase, uint32(v58))))
	v69 = v58
	v70 = v68
	v71 = v67
	goto L23
L25:
	;
	v67 = int32(0)
	goto L24
L26:
	;
	if v71 != 0 {
		goto L31
	} else {
		goto L32
	}
L27:
	;
	v78 = int32(0)
	v79 = v69
	v80 = v74
	goto L28
L28:
	;
	v82 = int32(10)
	v84 = v78*v82 - v80
	v85 = int32(*(*int8)(unsafe.Add(mBase, uint32(v79)+1)))
	v89 = v85 + int32(-48)
	if base.Ui32(v89) < base.Ui32(v82) {
		v78 = v84
		v79 = v79 + int32(1)
		v80 = v89
		goto L28
	} else {
		goto L30
	}
L29:
	;
	v92 = v84
	goto L26
L30:
	;
	goto L29
L31:
	;
	v98 = int32(0) - v92
	goto L33
L32:
	;
	v98 = v92
	goto L33
L33:
	;
	goto L19
L34:
	;
	return int32(_a_F_sentinelHandleConfiguration_2)
L35:
	;
	v157 = F_createSentinelValkeyInstance(m, v103, int32(1), v105, v155, v98, int32(0))
	mBase = m.M
	v160 = m.ExcPending
	if v160 != 0 {
		goto L50
	} else {
		goto L51
	}
L36:
	;
	v115 = v110 + int32(1)
	v116 = int32(*(*int8)(unsafe.Add(mBase, uint32(v110))))
	v117 = F___isspace_1(m, v116)
	mBase = m.M
	if v117 != 0 {
		v110 = v115
		goto L36
	} else {
		goto L38
	}
L37:
	;
	v118 = int32(1)
	switch v116&int32(255) + int32(-43) {
	case 0:
		v124 = v118
		goto L40
	default:
		v126 = v110
		v127 = v116
		v128 = v118
		goto L39
	case 2:
		goto L41
	}
L38:
	;
	goto L37
L39:
	;
	v131 = v127 + int32(-48)
	if base.Ui32(int32(9)) < base.Ui32(v131) {
		v149 = int32(0)
		goto L42
	} else {
		goto L43
	}
L40:
	;
	v125 = int32(*(*int8)(unsafe.Add(mBase, uint32(v115))))
	v126 = v115
	v127 = v125
	v128 = v124
	goto L39
L41:
	;
	v124 = int32(0)
	goto L40
L42:
	;
	if v128 != 0 {
		goto L47
	} else {
		goto L48
	}
L43:
	;
	v135 = int32(0)
	v136 = v126
	v137 = v131
	goto L44
L44:
	;
	v139 = int32(10)
	v141 = v135*v139 - v137
	v142 = int32(*(*int8)(unsafe.Add(mBase, uint32(v136)+1)))
	v146 = v142 + int32(-48)
	if base.Ui32(v146) < base.Ui32(v139) {
		v135 = v141
		v136 = v136 + int32(1)
		v137 = v146
		goto L44
	} else {
		goto L46
	}
L45:
	;
	v149 = v141
	goto L42
L46:
	;
	goto L45
L47:
	;
	v155 = int32(0) - v149
	goto L49
L48:
	;
	v155 = v149
	goto L49
L49:
	;
	goto L35
L50:
	;
	return int32(0)
L51:
	;
	if v157 != 0 {
		goto L14
	} else {
		goto L52
	}
L52:
	;
	goto L56
L53:
	;
	return int32(_a_F_sentinelHandleConfiguration_3)
L54:
	;
	return int32(_a_F_sentinelHandleConfiguration_4)
L55:
	;
	if v163 != int32(10) {
		goto L53
	} else {
		goto L57
	}
L56:
	;
	v163 = *(*int32)(unsafe.Add(mBase, _c_F_sentinelHandleConfiguration[0]))
	switch v163 + int32(-28) {
	case 0:
		goto L54
	case 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15:
		goto L53
	case 16:
		v1904 = int32(_a_F_sentinelHandleConfiguration_5)
		goto L13
	default:
		goto L55
	}
L57:
	;
	return int32(_a_F_sentinelHandleConfiguration_6)
L58:
	;
	v214 = base.B2i32(l1 != int32(3))
	if l1 != int32(3) {
		goto L72
	} else {
		goto L73
	}
L59:
	;
	v209 = F_tolower(m, v205)
	mBase = m.M
	v210 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v206))))
	v211 = F_tolower(m, v210)
	mBase = m.M
	goto L58
L60:
	;
	v179 = v7
	v180 = v174
	v181 = v177
	goto L63
L61:
	;
	v205 = int32(0)
	v206 = v174
	goto L59
L62:
	;
	v205 = v202 & int32(255)
	v206 = v201
	goto L59
L63:
	;
	v183 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v180))))
	if v183 == int32(0) {
		v201 = v180
		v202 = v181
		goto L62
	} else {
		goto L65
	}
L64:
	;
	v201 = v195
	v202 = int32(0)
	goto L62
L65:
	;
	v187 = v181 & int32(255)
	if v187 == v183 {
		goto L66
	} else {
		goto L67
	}
L66:
	;
	v194 = int32(1)
	v195 = v180 + v194
	v196 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v179)+1)))
	if v196 != 0 {
		v179 = v179 + v194
		v180 = v195
		v181 = v196
		goto L63
	} else {
		goto L69
	}
L67:
	;
	v189 = F_tolower(m, v187)
	mBase = m.M
	v190 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v180))))
	v191 = F_tolower(m, v190)
	mBase = m.M
	if v189 == v191 {
		goto L66
	} else {
		goto L68
	}
L68:
	;
	v193 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v179))))
	v201 = v180
	v202 = v193
	goto L62
L69:
	;
	goto L64
L70:
	;
	v1151 = int32(_a_F_sentinelHandleConfiguration_7)
	v1154 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7))))
	if v1154 != 0 {
		goto L396
	} else {
		goto L397
	}
L71:
	;
	F_sentinelPropagateDownAfterPeriod(m, v220)
	mBase = m.M
	v1149 = m.ExcPending
	if v1149 != 0 {
		goto L50
	} else {
		goto L393
	}
L72:
	;
	v282 = int32(_a_F_sentinelHandleConfiguration_8)
	v285 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7))))
	if v285 != 0 {
		goto L98
	} else {
		goto L99
	}
L73:
	;
	if v209-v211 != 0 {
		goto L72
	} else {
		goto L74
	}
L74:
	;
	v215 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v216 = F_sdsnew(m, v215)
	mBase = m.M
	v217 = m.ExcPending
	if v217 != 0 {
		goto L50
	} else {
		goto L75
	}
L75:
	;
	v219 = *(*int32)(unsafe.Add(mBase, _c_F_sentinelHandleConfiguration[1]))
	v220 = F_dictFetchValue(m, v219, v216)
	mBase = m.M
	v221 = m.ExcPending
	if v221 != 0 {
		goto L50
	} else {
		goto L76
	}
L76:
	;
	F_sdsfree(m, v216)
	mBase = m.M
	v223 = m.ExcPending
	if v223 != 0 {
		goto L50
	} else {
		goto L77
	}
L77:
	;
	if v220 != 0 {
		goto L78
	} else {
		goto L79
	}
L78:
	;
	v226 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v230 = v226
	goto L81
L79:
	;
	return int32(_a_F_sentinelHandleConfiguration_9)
L80:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v220)+72)) = base.I64_extend_i32_s(v275)
	if int32(1) <= v275 {
		goto L71
	} else {
		goto L95
	}
L81:
	;
	v235 = v230 + int32(1)
	v236 = int32(*(*int8)(unsafe.Add(mBase, uint32(v230))))
	v237 = F___isspace_1(m, v236)
	mBase = m.M
	if v237 != 0 {
		v230 = v235
		goto L81
	} else {
		goto L83
	}
L82:
	;
	v238 = int32(1)
	switch v236&int32(255) + int32(-43) {
	case 0:
		v244 = v238
		goto L85
	default:
		v246 = v230
		v247 = v236
		v248 = v238
		goto L84
	case 2:
		goto L86
	}
L83:
	;
	goto L82
L84:
	;
	v251 = v247 + int32(-48)
	if base.Ui32(int32(9)) < base.Ui32(v251) {
		v269 = int32(0)
		goto L87
	} else {
		goto L88
	}
L85:
	;
	v245 = int32(*(*int8)(unsafe.Add(mBase, uint32(v235))))
	v246 = v235
	v247 = v245
	v248 = v244
	goto L84
L86:
	;
	v244 = int32(0)
	goto L85
L87:
	;
	if v248 != 0 {
		goto L92
	} else {
		goto L93
	}
L88:
	;
	v255 = int32(0)
	v256 = v246
	v257 = v251
	goto L89
L89:
	;
	v259 = int32(10)
	v261 = v255*v259 - v257
	v262 = int32(*(*int8)(unsafe.Add(mBase, uint32(v256)+1)))
	v266 = v262 + int32(-48)
	if base.Ui32(v266) < base.Ui32(v259) {
		v255 = v261
		v256 = v256 + int32(1)
		v257 = v266
		goto L89
	} else {
		goto L91
	}
L90:
	;
	v269 = v261
	goto L87
L91:
	;
	goto L90
L92:
	;
	v275 = int32(0) - v269
	goto L94
L93:
	;
	v275 = v269
	goto L94
L94:
	;
	goto L80
L95:
	;
	return int32(_a_F_sentinelHandleConfiguration_10)
L96:
	;
	if l1 != int32(3) {
		goto L108
	} else {
		goto L109
	}
L97:
	;
	v317 = F_tolower(m, v313)
	mBase = m.M
	v318 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v314))))
	v319 = F_tolower(m, v318)
	mBase = m.M
	goto L96
L98:
	;
	v287 = v7
	v288 = v282
	v289 = v285
	goto L101
L99:
	;
	v313 = int32(0)
	v314 = v282
	goto L97
L100:
	;
	v313 = v310 & int32(255)
	v314 = v309
	goto L97
L101:
	;
	v291 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v288))))
	if v291 == int32(0) {
		v309 = v288
		v310 = v289
		goto L100
	} else {
		goto L103
	}
L102:
	;
	v309 = v303
	v310 = int32(0)
	goto L100
L103:
	;
	v295 = v289 & int32(255)
	if v295 == v291 {
		goto L104
	} else {
		goto L105
	}
L104:
	;
	v302 = int32(1)
	v303 = v288 + v302
	v304 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v287)+1)))
	if v304 != 0 {
		v287 = v287 + v302
		v288 = v303
		v289 = v304
		goto L101
	} else {
		goto L107
	}
L105:
	;
	v297 = F_tolower(m, v295)
	mBase = m.M
	v298 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v288))))
	v299 = F_tolower(m, v298)
	mBase = m.M
	if v297 == v299 {
		goto L104
	} else {
		goto L106
	}
L106:
	;
	v301 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v287))))
	v309 = v288
	v310 = v301
	goto L100
L107:
	;
	goto L102
L108:
	;
	v388 = int32(_a_F_sentinelHandleConfiguration_11)
	v391 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7))))
	if v391 != 0 {
		goto L134
	} else {
		goto L135
	}
L109:
	;
	if v317-v319 != 0 {
		goto L108
	} else {
		goto L110
	}
L110:
	;
	v321 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v322 = F_sdsnew(m, v321)
	mBase = m.M
	v323 = m.ExcPending
	if v323 != 0 {
		goto L50
	} else {
		goto L111
	}
L111:
	;
	v325 = *(*int32)(unsafe.Add(mBase, _c_F_sentinelHandleConfiguration[1]))
	v326 = F_dictFetchValue(m, v325, v322)
	mBase = m.M
	v327 = m.ExcPending
	if v327 != 0 {
		goto L50
	} else {
		goto L112
	}
L112:
	;
	F_sdsfree(m, v322)
	mBase = m.M
	v329 = m.ExcPending
	if v329 != 0 {
		goto L50
	} else {
		goto L113
	}
L113:
	;
	if v326 != 0 {
		goto L114
	} else {
		goto L115
	}
L114:
	;
	v332 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v336 = v332
	goto L117
L115:
	;
	return int32(_a_F_sentinelHandleConfiguration_9)
L116:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v326)+264)) = base.I64_extend_i32_s(v381)
	if int32(1) <= v381 {
		goto L14
	} else {
		goto L131
	}
L117:
	;
	v341 = v336 + int32(1)
	v342 = int32(*(*int8)(unsafe.Add(mBase, uint32(v336))))
	v343 = F___isspace_1(m, v342)
	mBase = m.M
	if v343 != 0 {
		v336 = v341
		goto L117
	} else {
		goto L119
	}
L118:
	;
	v344 = int32(1)
	switch v342&int32(255) + int32(-43) {
	case 0:
		v350 = v344
		goto L121
	default:
		v352 = v336
		v353 = v342
		v354 = v344
		goto L120
	case 2:
		goto L122
	}
L119:
	;
	goto L118
L120:
	;
	v357 = v353 + int32(-48)
	if base.Ui32(int32(9)) < base.Ui32(v357) {
		v375 = int32(0)
		goto L123
	} else {
		goto L124
	}
L121:
	;
	v351 = int32(*(*int8)(unsafe.Add(mBase, uint32(v341))))
	v352 = v341
	v353 = v351
	v354 = v350
	goto L120
L122:
	;
	v350 = int32(0)
	goto L121
L123:
	;
	if v354 != 0 {
		goto L128
	} else {
		goto L129
	}
L124:
	;
	v361 = int32(0)
	v362 = v352
	v363 = v357
	goto L125
L125:
	;
	v365 = int32(10)
	v367 = v361*v365 - v363
	v368 = int32(*(*int8)(unsafe.Add(mBase, uint32(v362)+1)))
	v372 = v368 + int32(-48)
	if base.Ui32(v372) < base.Ui32(v365) {
		v361 = v367
		v362 = v362 + int32(1)
		v363 = v372
		goto L125
	} else {
		goto L127
	}
L126:
	;
	v375 = v367
	goto L123
L127:
	;
	goto L126
L128:
	;
	v381 = int32(0) - v375
	goto L130
L129:
	;
	v381 = v375
	goto L130
L130:
	;
	goto L116
L131:
	;
	return int32(_a_F_sentinelHandleConfiguration_10)
L132:
	;
	v428 = base.B2i32(l1 != int32(3))
	if l1 != int32(3) {
		goto L144
	} else {
		goto L145
	}
L133:
	;
	v423 = F_tolower(m, v419)
	mBase = m.M
	v424 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v420))))
	v425 = F_tolower(m, v424)
	mBase = m.M
	goto L132
L134:
	;
	v393 = v7
	v394 = v388
	v395 = v391
	goto L137
L135:
	;
	v419 = int32(0)
	v420 = v388
	goto L133
L136:
	;
	v419 = v416 & int32(255)
	v420 = v415
	goto L133
L137:
	;
	v397 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v394))))
	if v397 == int32(0) {
		v415 = v394
		v416 = v395
		goto L136
	} else {
		goto L139
	}
L138:
	;
	v415 = v409
	v416 = int32(0)
	goto L136
L139:
	;
	v401 = v395 & int32(255)
	if v401 == v397 {
		goto L140
	} else {
		goto L141
	}
L140:
	;
	v408 = int32(1)
	v409 = v394 + v408
	v410 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v393)+1)))
	if v410 != 0 {
		v393 = v393 + v408
		v394 = v409
		v395 = v410
		goto L137
	} else {
		goto L143
	}
L141:
	;
	v403 = F_tolower(m, v401)
	mBase = m.M
	v404 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v394))))
	v405 = F_tolower(m, v404)
	mBase = m.M
	if v403 == v405 {
		goto L140
	} else {
		goto L142
	}
L142:
	;
	v407 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v393))))
	v415 = v394
	v416 = v407
	goto L136
L143:
	;
	goto L138
L144:
	;
	v491 = int32(_a_F_sentinelHandleConfiguration_12)
	v494 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7))))
	if v494 != 0 {
		goto L169
	} else {
		goto L170
	}
L145:
	;
	if v423-v425 != 0 {
		goto L144
	} else {
		goto L146
	}
L146:
	;
	v429 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v430 = F_sdsnew(m, v429)
	mBase = m.M
	v431 = m.ExcPending
	if v431 != 0 {
		goto L50
	} else {
		goto L147
	}
L147:
	;
	v433 = *(*int32)(unsafe.Add(mBase, _c_F_sentinelHandleConfiguration[1]))
	v434 = F_dictFetchValue(m, v433, v430)
	mBase = m.M
	v435 = m.ExcPending
	if v435 != 0 {
		goto L50
	} else {
		goto L148
	}
L148:
	;
	F_sdsfree(m, v430)
	mBase = m.M
	v437 = m.ExcPending
	if v437 != 0 {
		goto L50
	} else {
		goto L149
	}
L149:
	;
	if v434 != 0 {
		goto L150
	} else {
		goto L151
	}
L150:
	;
	v440 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v444 = v440
	goto L153
L151:
	;
	return int32(_a_F_sentinelHandleConfiguration_9)
L152:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v434)+156)) = v489
	goto L14
L153:
	;
	v449 = v444 + int32(1)
	v450 = int32(*(*int8)(unsafe.Add(mBase, uint32(v444))))
	v451 = F___isspace_1(m, v450)
	mBase = m.M
	if v451 != 0 {
		v444 = v449
		goto L153
	} else {
		goto L155
	}
L154:
	;
	v452 = int32(1)
	switch v450&int32(255) + int32(-43) {
	case 0:
		v458 = v452
		goto L157
	default:
		v460 = v444
		v461 = v450
		v462 = v452
		goto L156
	case 2:
		goto L158
	}
L155:
	;
	goto L154
L156:
	;
	v465 = v461 + int32(-48)
	if base.Ui32(int32(9)) < base.Ui32(v465) {
		v483 = int32(0)
		goto L159
	} else {
		goto L160
	}
L157:
	;
	v459 = int32(*(*int8)(unsafe.Add(mBase, uint32(v449))))
	v460 = v449
	v461 = v459
	v462 = v458
	goto L156
L158:
	;
	v458 = int32(0)
	goto L157
L159:
	;
	if v462 != 0 {
		goto L164
	} else {
		goto L165
	}
L160:
	;
	v469 = int32(0)
	v470 = v460
	v471 = v465
	goto L161
L161:
	;
	v473 = int32(10)
	v475 = v469*v473 - v471
	v476 = int32(*(*int8)(unsafe.Add(mBase, uint32(v470)+1)))
	v480 = v476 + int32(-48)
	if base.Ui32(v480) < base.Ui32(v473) {
		v469 = v475
		v470 = v470 + int32(1)
		v471 = v480
		goto L161
	} else {
		goto L163
	}
L162:
	;
	v483 = v475
	goto L159
L163:
	;
	goto L162
L164:
	;
	v489 = int32(0) - v483
	goto L166
L165:
	;
	v489 = v483
	goto L166
L166:
	;
	goto L152
L167:
	;
	if l1 != int32(3) {
		goto L179
	} else {
		goto L180
	}
L168:
	;
	v526 = F_tolower(m, v522)
	mBase = m.M
	v527 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v523))))
	v528 = F_tolower(m, v527)
	mBase = m.M
	goto L167
L169:
	;
	v496 = v7
	v497 = v491
	v498 = v494
	goto L172
L170:
	;
	v522 = int32(0)
	v523 = v491
	goto L168
L171:
	;
	v522 = v519 & int32(255)
	v523 = v518
	goto L168
L172:
	;
	v500 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v497))))
	if v500 == int32(0) {
		v518 = v497
		v519 = v498
		goto L171
	} else {
		goto L174
	}
L173:
	;
	v518 = v512
	v519 = int32(0)
	goto L171
L174:
	;
	v504 = v498 & int32(255)
	if v504 == v500 {
		goto L175
	} else {
		goto L176
	}
L175:
	;
	v511 = int32(1)
	v512 = v497 + v511
	v513 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v496)+1)))
	if v513 != 0 {
		v496 = v496 + v511
		v497 = v512
		v498 = v513
		goto L172
	} else {
		goto L178
	}
L176:
	;
	v506 = F_tolower(m, v504)
	mBase = m.M
	v507 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v497))))
	v508 = F_tolower(m, v507)
	mBase = m.M
	if v506 == v508 {
		goto L175
	} else {
		goto L177
	}
L177:
	;
	v510 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v496))))
	v518 = v497
	v519 = v510
	goto L171
L178:
	;
	goto L173
L179:
	;
	v552 = int32(_a_F_sentinelHandleConfiguration_13)
	v555 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7))))
	if v555 != 0 {
		goto L192
	} else {
		goto L193
	}
L180:
	;
	if v526-v528 != 0 {
		goto L179
	} else {
		goto L181
	}
L181:
	;
	v530 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v531 = F_sdsnew(m, v530)
	mBase = m.M
	v532 = m.ExcPending
	if v532 != 0 {
		goto L50
	} else {
		goto L182
	}
L182:
	;
	v534 = *(*int32)(unsafe.Add(mBase, _c_F_sentinelHandleConfiguration[1]))
	v535 = F_dictFetchValue(m, v534, v531)
	mBase = m.M
	v536 = m.ExcPending
	if v536 != 0 {
		goto L50
	} else {
		goto L183
	}
L183:
	;
	F_sdsfree(m, v531)
	mBase = m.M
	v538 = m.ExcPending
	if v538 != 0 {
		goto L50
	} else {
		goto L184
	}
L184:
	;
	if v535 != 0 {
		goto L185
	} else {
		goto L186
	}
L185:
	;
	v541 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v543 = F_access(m, v541, int32(1))
	mBase = m.M
	if v543 != int32(-1) {
		goto L187
	} else {
		goto L188
	}
L186:
	;
	return int32(_a_F_sentinelHandleConfiguration_9)
L187:
	;
	v548 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v549 = F_sdsnew(m, v548)
	mBase = m.M
	v550 = m.ExcPending
	if v550 != 0 {
		goto L50
	} else {
		goto L189
	}
L188:
	;
	return int32(_a_F_sentinelHandleConfiguration_14)
L189:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v535)+284)) = v549
	goto L14
L190:
	;
	v592 = base.B2i32(l1 != int32(3))
	if l1 != int32(3) {
		goto L202
	} else {
		goto L203
	}
L191:
	;
	v587 = F_tolower(m, v583)
	mBase = m.M
	v588 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v584))))
	v589 = F_tolower(m, v588)
	mBase = m.M
	goto L190
L192:
	;
	v557 = v7
	v558 = v552
	v559 = v555
	goto L195
L193:
	;
	v583 = int32(0)
	v584 = v552
	goto L191
L194:
	;
	v583 = v580 & int32(255)
	v584 = v579
	goto L191
L195:
	;
	v561 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v558))))
	if v561 == int32(0) {
		v579 = v558
		v580 = v559
		goto L194
	} else {
		goto L197
	}
L196:
	;
	v579 = v573
	v580 = int32(0)
	goto L194
L197:
	;
	v565 = v559 & int32(255)
	if v565 == v561 {
		goto L198
	} else {
		goto L199
	}
L198:
	;
	v572 = int32(1)
	v573 = v558 + v572
	v574 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v557)+1)))
	if v574 != 0 {
		v557 = v557 + v572
		v558 = v573
		v559 = v574
		goto L195
	} else {
		goto L201
	}
L199:
	;
	v567 = F_tolower(m, v565)
	mBase = m.M
	v568 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v558))))
	v569 = F_tolower(m, v568)
	mBase = m.M
	if v567 == v569 {
		goto L198
	} else {
		goto L200
	}
L200:
	;
	v571 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v557))))
	v579 = v558
	v580 = v571
	goto L194
L201:
	;
	goto L196
L202:
	;
	v609 = int32(_a_F_sentinelHandleConfiguration_15)
	v612 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7))))
	if v612 != 0 {
		goto L213
	} else {
		goto L214
	}
L203:
	;
	if v587-v589 != 0 {
		goto L202
	} else {
		goto L204
	}
L204:
	;
	v593 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v594 = F_sentinelGetPrimaryByName(m, v593)
	mBase = m.M
	v595 = m.ExcPending
	if v595 != 0 {
		goto L50
	} else {
		goto L206
	}
L205:
	;
	v598 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v600 = F_access(m, v598, int32(1))
	mBase = m.M
	if v600 != int32(-1) {
		goto L208
	} else {
		goto L209
	}
L206:
	;
	if v594 != 0 {
		goto L205
	} else {
		goto L207
	}
L207:
	;
	return int32(_a_F_sentinelHandleConfiguration_9)
L208:
	;
	v605 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v606 = F_sdsnew(m, v605)
	mBase = m.M
	v607 = m.ExcPending
	if v607 != 0 {
		goto L50
	} else {
		goto L210
	}
L209:
	;
	return int32(_a_F_sentinelHandleConfiguration_16)
L210:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v594)+288)) = v606
	goto L14
L211:
	;
	if l1 != int32(3) {
		goto L223
	} else {
		goto L224
	}
L212:
	;
	v644 = F_tolower(m, v640)
	mBase = m.M
	v645 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v641))))
	v646 = F_tolower(m, v645)
	mBase = m.M
	goto L211
L213:
	;
	v614 = v7
	v615 = v609
	v616 = v612
	goto L216
L214:
	;
	v640 = int32(0)
	v641 = v609
	goto L212
L215:
	;
	v640 = v637 & int32(255)
	v641 = v636
	goto L212
L216:
	;
	v618 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v615))))
	if v618 == int32(0) {
		v636 = v615
		v637 = v616
		goto L215
	} else {
		goto L218
	}
L217:
	;
	v636 = v630
	v637 = int32(0)
	goto L215
L218:
	;
	v622 = v616 & int32(255)
	if v622 == v618 {
		goto L219
	} else {
		goto L220
	}
L219:
	;
	v629 = int32(1)
	v630 = v615 + v629
	v631 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v614)+1)))
	if v631 != 0 {
		v614 = v614 + v629
		v615 = v630
		v616 = v631
		goto L216
	} else {
		goto L222
	}
L220:
	;
	v624 = F_tolower(m, v622)
	mBase = m.M
	v625 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v615))))
	v626 = F_tolower(m, v625)
	mBase = m.M
	if v624 == v626 {
		goto L219
	} else {
		goto L221
	}
L221:
	;
	v628 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v614))))
	v636 = v615
	v637 = v628
	goto L215
L222:
	;
	goto L217
L223:
	;
	v657 = int32(_a_F_sentinelHandleConfiguration_17)
	v660 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7))))
	if v660 != 0 {
		goto L232
	} else {
		goto L233
	}
L224:
	;
	if v644-v646 != 0 {
		goto L223
	} else {
		goto L225
	}
L225:
	;
	v648 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v649 = F_sentinelGetPrimaryByName(m, v648)
	mBase = m.M
	v650 = m.ExcPending
	if v650 != 0 {
		goto L50
	} else {
		goto L227
	}
L226:
	;
	v653 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v654 = F_sdsnew(m, v653)
	mBase = m.M
	v655 = m.ExcPending
	if v655 != 0 {
		goto L50
	} else {
		goto L229
	}
L227:
	;
	if v649 != 0 {
		goto L226
	} else {
		goto L228
	}
L228:
	;
	return int32(_a_F_sentinelHandleConfiguration_9)
L229:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v649)+160)) = v654
	goto L14
L230:
	;
	if l1 != int32(3) {
		goto L242
	} else {
		goto L243
	}
L231:
	;
	v692 = F_tolower(m, v688)
	mBase = m.M
	v693 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v689))))
	v694 = F_tolower(m, v693)
	mBase = m.M
	goto L230
L232:
	;
	v662 = v7
	v663 = v657
	v664 = v660
	goto L235
L233:
	;
	v688 = int32(0)
	v689 = v657
	goto L231
L234:
	;
	v688 = v685 & int32(255)
	v689 = v684
	goto L231
L235:
	;
	v666 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v663))))
	if v666 == int32(0) {
		v684 = v663
		v685 = v664
		goto L234
	} else {
		goto L237
	}
L236:
	;
	v684 = v678
	v685 = int32(0)
	goto L234
L237:
	;
	v670 = v664 & int32(255)
	if v670 == v666 {
		goto L238
	} else {
		goto L239
	}
L238:
	;
	v677 = int32(1)
	v678 = v663 + v677
	v679 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v662)+1)))
	if v679 != 0 {
		v662 = v662 + v677
		v663 = v678
		v664 = v679
		goto L235
	} else {
		goto L241
	}
L239:
	;
	v672 = F_tolower(m, v670)
	mBase = m.M
	v673 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v663))))
	v674 = F_tolower(m, v673)
	mBase = m.M
	if v672 == v674 {
		goto L238
	} else {
		goto L240
	}
L240:
	;
	v676 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v662))))
	v684 = v663
	v685 = v676
	goto L234
L241:
	;
	goto L236
L242:
	;
	v707 = int32(_a_F_sentinelHandleConfiguration_18)
	v710 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7))))
	if v710 != 0 {
		goto L251
	} else {
		goto L252
	}
L243:
	;
	if v692-v694 != 0 {
		goto L242
	} else {
		goto L244
	}
L244:
	;
	v698 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v699 = F_sentinelGetPrimaryByName(m, v698)
	mBase = m.M
	v700 = m.ExcPending
	if v700 != 0 {
		goto L50
	} else {
		goto L246
	}
L245:
	;
	v703 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v704 = F_sdsnew(m, v703)
	mBase = m.M
	v705 = m.ExcPending
	if v705 != 0 {
		goto L50
	} else {
		goto L248
	}
L246:
	;
	if v699 != 0 {
		goto L245
	} else {
		goto L247
	}
L247:
	;
	return int32(_a_F_sentinelHandleConfiguration_9)
L248:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v699)+164)) = v704
	goto L14
L249:
	;
	v747 = base.B2i32(l1 != int32(2))
	if l1 != int32(2) {
		goto L261
	} else {
		goto L262
	}
L250:
	;
	v742 = F_tolower(m, v738)
	mBase = m.M
	v743 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v739))))
	v744 = F_tolower(m, v743)
	mBase = m.M
	goto L249
L251:
	;
	v712 = v7
	v713 = v707
	v714 = v710
	goto L254
L252:
	;
	v738 = int32(0)
	v739 = v707
	goto L250
L253:
	;
	v738 = v735 & int32(255)
	v739 = v734
	goto L250
L254:
	;
	v716 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v713))))
	if v716 == int32(0) {
		v734 = v713
		v735 = v714
		goto L253
	} else {
		goto L256
	}
L255:
	;
	v734 = v728
	v735 = int32(0)
	goto L253
L256:
	;
	v720 = v714 & int32(255)
	if v720 == v716 {
		goto L257
	} else {
		goto L258
	}
L257:
	;
	v727 = int32(1)
	v728 = v713 + v727
	v729 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v712)+1)))
	if v729 != 0 {
		v712 = v712 + v727
		v713 = v728
		v714 = v729
		goto L254
	} else {
		goto L260
	}
L258:
	;
	v722 = F_tolower(m, v720)
	mBase = m.M
	v723 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v713))))
	v724 = F_tolower(m, v723)
	mBase = m.M
	if v722 == v724 {
		goto L257
	} else {
		goto L259
	}
L259:
	;
	v726 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v712))))
	v734 = v713
	v735 = v726
	goto L253
L260:
	;
	goto L255
L261:
	;
	v758 = int32(_a_F_sentinelHandleConfiguration_19)
	v761 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7))))
	if v761 != 0 {
		goto L268
	} else {
		goto L269
	}
L262:
	;
	if v742-v744 != 0 {
		goto L261
	} else {
		goto L263
	}
L263:
	;
	v748 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v752 = F_strtox_2(m, v748, int32(0), int32(10), int64(-1))
	mBase = m.M
	goto L264
L264:
	;
	v754 = *(*int64)(unsafe.Add(mBase, _c_F_sentinelHandleConfiguration[2]))
	if base.Ui64(v752) <= base.Ui64(v754) {
		goto L14
	} else {
		goto L265
	}
L265:
	;
	*(*int64)(unsafe.Add(mBase, _c_F_sentinelHandleConfiguration[2])) = v752
	goto L14
L266:
	;
	if l1 != int32(2) {
		goto L278
	} else {
		goto L279
	}
L267:
	;
	v793 = F_tolower(m, v789)
	mBase = m.M
	v794 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v790))))
	v795 = F_tolower(m, v794)
	mBase = m.M
	goto L266
L268:
	;
	v763 = v7
	v764 = v758
	v765 = v761
	goto L271
L269:
	;
	v789 = int32(0)
	v790 = v758
	goto L267
L270:
	;
	v789 = v786 & int32(255)
	v790 = v785
	goto L267
L271:
	;
	v767 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v764))))
	if v767 == int32(0) {
		v785 = v764
		v786 = v765
		goto L270
	} else {
		goto L273
	}
L272:
	;
	v785 = v779
	v786 = int32(0)
	goto L270
L273:
	;
	v771 = v765 & int32(255)
	if v771 == v767 {
		goto L274
	} else {
		goto L275
	}
L274:
	;
	v778 = int32(1)
	v779 = v764 + v778
	v780 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v763)+1)))
	if v780 != 0 {
		v763 = v763 + v778
		v764 = v779
		v765 = v780
		goto L271
	} else {
		goto L277
	}
L275:
	;
	v773 = F_tolower(m, v771)
	mBase = m.M
	v774 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v764))))
	v775 = F_tolower(m, v774)
	mBase = m.M
	if v773 == v775 {
		goto L274
	} else {
		goto L276
	}
L276:
	;
	v777 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v763))))
	v785 = v764
	v786 = v777
	goto L270
L277:
	;
	goto L272
L278:
	;
	v880 = int32(_a_F_sentinelHandleConfiguration_20)
	v883 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7))))
	if v883 != 0 {
		goto L301
	} else {
		goto L302
	}
L279:
	;
	if v793-v795 != 0 {
		goto L278
	} else {
		goto L280
	}
L280:
	;
	v797 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v797&int32(3) == int32(0) {
		v819 = v797
		goto L284
	} else {
		goto L285
	}
L281:
	;
	v857 = int32(0)
	v858 = *(*int64)(unsafe.Add(mBase, uint32(v797)))
	*(*int64)(unsafe.Add(mBase, _c_F_sentinelHandleConfiguration[3])) = v858
	v863 = *(*int64)(unsafe.Add(mBase, uint32(v797+int32(32))))
	*(*int64)(unsafe.Add(mBase, _c_F_sentinelHandleConfiguration[4])) = v863
	v868 = *(*int64)(unsafe.Add(mBase, uint32(v797+int32(24))))
	*(*int64)(unsafe.Add(mBase, _c_F_sentinelHandleConfiguration[5])) = v868
	v873 = *(*int64)(unsafe.Add(mBase, uint32(v797+int32(16))))
	*(*int64)(unsafe.Add(mBase, _c_F_sentinelHandleConfiguration[6])) = v873
	v878 = *(*int64)(unsafe.Add(mBase, uint32(v797+int32(8))))
	*(*int64)(unsafe.Add(mBase, _c_F_sentinelHandleConfiguration[7])) = v878
	goto L14
L282:
	;
	if v852 == int32(40) {
		goto L281
	} else {
		goto L298
	}
L283:
	;
	v852 = v844 - v797
	goto L282
L284:
	;
	v823 = v819
	goto L292
L285:
	;
	v805 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v797))))
	if v805 != 0 {
		goto L286
	} else {
		goto L287
	}
L286:
	;
	v808 = v797
	goto L288
L287:
	;
	v852 = v797 - v797
	goto L282
L288:
	;
	v812 = v808 + int32(1)
	if v812&int32(3) == int32(0) {
		v819 = v812
		goto L284
	} else {
		goto L290
	}
L290:
	;
	v817 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v812))))
	if v817 != 0 {
		v808 = v812
		goto L288
	} else {
		goto L291
	}
L291:
	;
	v844 = v812
	goto L283
L292:
	;
	v829 = *(*int32)(unsafe.Add(mBase, uint32(v823)))
	v832 = int32(-2139062144)
	if (int32(16843008)-v829|v829)&v832 == v832 {
		v823 = v823 + int32(4)
		goto L292
	} else {
		goto L294
	}
L293:
	;
	v838 = v823
	goto L295
L294:
	;
	goto L293
L295:
	;
	v842 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v838))))
	if v842 != 0 {
		v838 = v838 + int32(1)
		goto L295
	} else {
		goto L297
	}
L296:
	;
	v844 = v838
	goto L283
L297:
	;
	goto L296
L298:
	;
	return int32(_a_F_sentinelHandleConfiguration_21)
L299:
	;
	v920 = base.B2i32(l1 != int32(3))
	if l1 != int32(3) {
		goto L311
	} else {
		goto L312
	}
L300:
	;
	v915 = F_tolower(m, v911)
	mBase = m.M
	v916 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v912))))
	v917 = F_tolower(m, v916)
	mBase = m.M
	goto L299
L301:
	;
	v885 = v7
	v886 = v880
	v887 = v883
	goto L304
L302:
	;
	v911 = int32(0)
	v912 = v880
	goto L300
L303:
	;
	v911 = v908 & int32(255)
	v912 = v907
	goto L300
L304:
	;
	v889 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v886))))
	if v889 == int32(0) {
		v907 = v886
		v908 = v887
		goto L303
	} else {
		goto L306
	}
L305:
	;
	v907 = v901
	v908 = int32(0)
	goto L303
L306:
	;
	v893 = v887 & int32(255)
	if v893 == v889 {
		goto L307
	} else {
		goto L308
	}
L307:
	;
	v900 = int32(1)
	v901 = v886 + v900
	v902 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v885)+1)))
	if v902 != 0 {
		v885 = v885 + v900
		v886 = v901
		v887 = v902
		goto L304
	} else {
		goto L310
	}
L308:
	;
	v895 = F_tolower(m, v893)
	mBase = m.M
	v896 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v886))))
	v897 = F_tolower(m, v896)
	mBase = m.M
	if v895 == v897 {
		goto L307
	} else {
		goto L309
	}
L309:
	;
	v899 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v885))))
	v907 = v886
	v908 = v899
	goto L303
L310:
	;
	goto L305
L311:
	;
	v937 = int32(_a_F_sentinelHandleConfiguration_22)
	v940 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7))))
	if v940 != 0 {
		goto L321
	} else {
		goto L322
	}
L312:
	;
	if v915-v917 != 0 {
		goto L311
	} else {
		goto L313
	}
L313:
	;
	v921 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v922 = F_sentinelGetPrimaryByName(m, v921)
	mBase = m.M
	v923 = m.ExcPending
	if v923 != 0 {
		goto L50
	} else {
		goto L315
	}
L314:
	;
	v926 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v930 = F_strtox_2(m, v926, int32(0), int32(10), int64(-1))
	mBase = m.M
	goto L317
L315:
	;
	if v922 != 0 {
		goto L314
	} else {
		goto L316
	}
L316:
	;
	return int32(_a_F_sentinelHandleConfiguration_9)
L317:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v922)+16)) = v930
	v933 = *(*int64)(unsafe.Add(mBase, _c_F_sentinelHandleConfiguration[2]))
	if base.Ui64(v930) <= base.Ui64(v933) {
		goto L14
	} else {
		goto L318
	}
L318:
	;
	*(*int64)(unsafe.Add(mBase, _c_F_sentinelHandleConfiguration[2])) = v930
	goto L14
L319:
	;
	if l1 != int32(3) {
		goto L331
	} else {
		goto L332
	}
L320:
	;
	v972 = F_tolower(m, v968)
	mBase = m.M
	v973 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v969))))
	v974 = F_tolower(m, v973)
	mBase = m.M
	goto L319
L321:
	;
	v942 = v7
	v943 = v937
	v944 = v940
	goto L324
L322:
	;
	v968 = int32(0)
	v969 = v937
	goto L320
L323:
	;
	v968 = v965 & int32(255)
	v969 = v964
	goto L320
L324:
	;
	v946 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v943))))
	if v946 == int32(0) {
		v964 = v943
		v965 = v944
		goto L323
	} else {
		goto L326
	}
L325:
	;
	v964 = v958
	v965 = int32(0)
	goto L323
L326:
	;
	v950 = v944 & int32(255)
	if v950 == v946 {
		goto L327
	} else {
		goto L328
	}
L327:
	;
	v957 = int32(1)
	v958 = v943 + v957
	v959 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v942)+1)))
	if v959 != 0 {
		v942 = v942 + v957
		v943 = v958
		v944 = v959
		goto L324
	} else {
		goto L330
	}
L328:
	;
	v952 = F_tolower(m, v950)
	mBase = m.M
	v953 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v943))))
	v954 = F_tolower(m, v953)
	mBase = m.M
	if v952 == v954 {
		goto L327
	} else {
		goto L329
	}
L329:
	;
	v956 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v942))))
	v964 = v943
	v965 = v956
	goto L323
L330:
	;
	goto L325
L331:
	;
	v987 = int32(_a_F_sentinelHandleConfiguration_23)
	v990 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7))))
	if v990 != 0 {
		goto L342
	} else {
		goto L343
	}
L332:
	;
	if v972-v974 != 0 {
		goto L331
	} else {
		goto L333
	}
L333:
	;
	v976 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v977 = F_sentinelGetPrimaryByName(m, v976)
	mBase = m.M
	v978 = m.ExcPending
	if v978 != 0 {
		goto L50
	} else {
		goto L335
	}
L334:
	;
	v981 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v985 = F_strtox_2(m, v981, int32(0), int32(10), int64(-1))
	mBase = m.M
	goto L337
L335:
	;
	if v977 != 0 {
		goto L334
	} else {
		goto L336
	}
L336:
	;
	return int32(_a_F_sentinelHandleConfiguration_9)
L337:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v977)+224)) = v985
	goto L14
L338:
	;
	v1074 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v1075 = F_sentinelGetPrimaryByName(m, v1074)
	mBase = m.M
	v1076 = m.ExcPending
	if v1076 != 0 {
		goto L50
	} else {
		goto L369
	}
L339:
	;
	if l1 != int32(4) {
		goto L70
	} else {
		goto L367
	}
L340:
	;
	if v1022-v1024 == int32(0) {
		goto L339
	} else {
		goto L352
	}
L341:
	;
	v1022 = F_tolower(m, v1018)
	mBase = m.M
	v1023 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1019))))
	v1024 = F_tolower(m, v1023)
	mBase = m.M
	goto L340
L342:
	;
	v992 = v7
	v993 = v987
	v994 = v990
	goto L345
L343:
	;
	v1018 = int32(0)
	v1019 = v987
	goto L341
L344:
	;
	v1018 = v1015 & int32(255)
	v1019 = v1014
	goto L341
L345:
	;
	v996 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v993))))
	if v996 == int32(0) {
		v1014 = v993
		v1015 = v994
		goto L344
	} else {
		goto L347
	}
L346:
	;
	v1014 = v1008
	v1015 = int32(0)
	goto L344
L347:
	;
	v1000 = v994 & int32(255)
	if v1000 == v996 {
		goto L348
	} else {
		goto L349
	}
L348:
	;
	v1007 = int32(1)
	v1008 = v993 + v1007
	v1009 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v992)+1)))
	if v1009 != 0 {
		v992 = v992 + v1007
		v993 = v1008
		v994 = v1009
		goto L345
	} else {
		goto L351
	}
L349:
	;
	v1002 = F_tolower(m, v1000)
	mBase = m.M
	v1003 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v993))))
	v1004 = F_tolower(m, v1003)
	mBase = m.M
	if v1002 == v1004 {
		goto L348
	} else {
		goto L350
	}
L350:
	;
	v1006 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v992))))
	v1014 = v993
	v1015 = v1006
	goto L344
L351:
	;
	goto L346
L352:
	;
	v1028 = int32(_a_F_sentinelHandleConfiguration_24)
	v1031 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7))))
	if v1031 != 0 {
		goto L355
	} else {
		goto L356
	}
L353:
	;
	if l1 != int32(4) {
		goto L70
	} else {
		goto L365
	}
L354:
	;
	v1063 = F_tolower(m, v1059)
	mBase = m.M
	v1064 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1060))))
	v1065 = F_tolower(m, v1064)
	mBase = m.M
	goto L353
L355:
	;
	v1033 = v7
	v1034 = v1028
	v1035 = v1031
	goto L358
L356:
	;
	v1059 = int32(0)
	v1060 = v1028
	goto L354
L357:
	;
	v1059 = v1056 & int32(255)
	v1060 = v1055
	goto L354
L358:
	;
	v1037 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1034))))
	if v1037 == int32(0) {
		v1055 = v1034
		v1056 = v1035
		goto L357
	} else {
		goto L360
	}
L359:
	;
	v1055 = v1049
	v1056 = int32(0)
	goto L357
L360:
	;
	v1041 = v1035 & int32(255)
	if v1041 == v1037 {
		goto L361
	} else {
		goto L362
	}
L361:
	;
	v1048 = int32(1)
	v1049 = v1034 + v1048
	v1050 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1033)+1)))
	if v1050 != 0 {
		v1033 = v1033 + v1048
		v1034 = v1049
		v1035 = v1050
		goto L358
	} else {
		goto L364
	}
L362:
	;
	v1043 = F_tolower(m, v1041)
	mBase = m.M
	v1044 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1034))))
	v1045 = F_tolower(m, v1044)
	mBase = m.M
	if v1043 == v1045 {
		goto L361
	} else {
		goto L363
	}
L363:
	;
	v1047 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1033))))
	v1055 = v1034
	v1056 = v1047
	goto L357
L364:
	;
	goto L359
L365:
	;
	if v1063-v1065 == int32(0) {
		goto L338
	} else {
		goto L366
	}
L366:
	;
	goto L70
L367:
	;
	goto L338
L368:
	;
	v1081 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1082 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v1086 = v1082
	goto L372
L369:
	;
	if v1075 != 0 {
		goto L368
	} else {
		goto L370
	}
L370:
	;
	return int32(_a_F_sentinelHandleConfiguration_9)
L371:
	;
	v1132 = *(*int32)(unsafe.Add(mBase, uint32(v1075)+152))
	v1133 = F_createSentinelValkeyInstance(m, int32(0), int32(2), v1081, v1131, v1132, v1075)
	mBase = m.M
	v1134 = m.ExcPending
	if v1134 != 0 {
		goto L50
	} else {
		goto L386
	}
L372:
	;
	v1091 = v1086 + int32(1)
	v1092 = int32(*(*int8)(unsafe.Add(mBase, uint32(v1086))))
	v1093 = F___isspace_1(m, v1092)
	mBase = m.M
	if v1093 != 0 {
		v1086 = v1091
		goto L372
	} else {
		goto L374
	}
L373:
	;
	v1094 = int32(1)
	switch v1092&int32(255) + int32(-43) {
	case 0:
		v1100 = v1094
		goto L376
	default:
		v1102 = v1086
		v1103 = v1092
		v1104 = v1094
		goto L375
	case 2:
		goto L377
	}
L374:
	;
	goto L373
L375:
	;
	v1107 = v1103 + int32(-48)
	if base.Ui32(int32(9)) < base.Ui32(v1107) {
		v1125 = int32(0)
		goto L378
	} else {
		goto L379
	}
L376:
	;
	v1101 = int32(*(*int8)(unsafe.Add(mBase, uint32(v1091))))
	v1102 = v1091
	v1103 = v1101
	v1104 = v1100
	goto L375
L377:
	;
	v1100 = int32(0)
	goto L376
L378:
	;
	if v1104 != 0 {
		goto L383
	} else {
		goto L384
	}
L379:
	;
	v1111 = int32(0)
	v1112 = v1102
	v1113 = v1107
	goto L380
L380:
	;
	v1115 = int32(10)
	v1117 = v1111*v1115 - v1113
	v1118 = int32(*(*int8)(unsafe.Add(mBase, uint32(v1112)+1)))
	v1122 = v1118 + int32(-48)
	if base.Ui32(v1122) < base.Ui32(v1115) {
		v1111 = v1117
		v1112 = v1112 + int32(1)
		v1113 = v1122
		goto L380
	} else {
		goto L382
	}
L381:
	;
	v1125 = v1117
	goto L378
L382:
	;
	goto L381
L383:
	;
	v1131 = int32(0) - v1125
	goto L385
L384:
	;
	v1131 = v1125
	goto L385
L385:
	;
	goto L371
L386:
	;
	if v1133 != 0 {
		goto L14
	} else {
		goto L387
	}
L387:
	;
	goto L391
L388:
	;
	return int32(_a_F_sentinelHandleConfiguration_3)
L389:
	;
	return int32(_a_F_sentinelHandleConfiguration_4)
L390:
	;
	if v1137 != int32(10) {
		goto L388
	} else {
		goto L392
	}
L391:
	;
	v1137 = *(*int32)(unsafe.Add(mBase, _c_F_sentinelHandleConfiguration[0]))
	switch v1137 + int32(-28) {
	case 0:
		goto L389
	case 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15:
		goto L388
	case 16:
		v1904 = int32(_a_F_sentinelHandleConfiguration_5)
		goto L13
	default:
		goto L390
	}
L392:
	;
	return int32(_a_F_sentinelHandleConfiguration_25)
L393:
	;
	goto L14
L394:
	;
	if l1&int32(-2) != int32(4) {
		goto L406
	} else {
		goto L407
	}
L395:
	;
	v1186 = F_tolower(m, v1182)
	mBase = m.M
	v1187 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1183))))
	v1188 = F_tolower(m, v1187)
	mBase = m.M
	goto L394
L396:
	;
	v1156 = v7
	v1157 = v1151
	v1158 = v1154
	goto L399
L397:
	;
	v1182 = int32(0)
	v1183 = v1151
	goto L395
L398:
	;
	v1182 = v1179 & int32(255)
	v1183 = v1178
	goto L395
L399:
	;
	v1160 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1157))))
	if v1160 == int32(0) {
		v1178 = v1157
		v1179 = v1158
		goto L398
	} else {
		goto L401
	}
L400:
	;
	v1178 = v1172
	v1179 = int32(0)
	goto L398
L401:
	;
	v1164 = v1158 & int32(255)
	if v1164 == v1160 {
		goto L402
	} else {
		goto L403
	}
L402:
	;
	v1171 = int32(1)
	v1172 = v1157 + v1171
	v1173 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1156)+1)))
	if v1173 != 0 {
		v1156 = v1156 + v1171
		v1157 = v1172
		v1158 = v1173
		goto L399
	} else {
		goto L405
	}
L403:
	;
	v1166 = F_tolower(m, v1164)
	mBase = m.M
	v1167 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1157))))
	v1168 = F_tolower(m, v1167)
	mBase = m.M
	if v1166 == v1168 {
		goto L402
	} else {
		goto L404
	}
L404:
	;
	v1170 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1156))))
	v1178 = v1157
	v1179 = v1170
	goto L398
L405:
	;
	goto L400
L406:
	;
	v1276 = int32(_a_F_sentinelHandleConfiguration_26)
	v1279 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7))))
	if v1279 != 0 {
		goto L440
	} else {
		goto L441
	}
L407:
	;
	if v1186-v1188 != 0 {
		goto L406
	} else {
		goto L408
	}
L408:
	;
	if l1 != int32(5) {
		goto L14
	} else {
		goto L409
	}
L409:
	;
	v1196 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v1197 = F_sentinelGetPrimaryByName(m, v1196)
	mBase = m.M
	v1198 = m.ExcPending
	if v1198 != 0 {
		goto L50
	} else {
		goto L411
	}
L410:
	;
	v1201 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v1203 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1204 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v1208 = v1204
	goto L415
L411:
	;
	if v1197 != 0 {
		goto L410
	} else {
		goto L412
	}
L412:
	;
	return int32(_a_F_sentinelHandleConfiguration_9)
L413:
	;
	v1270 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v1271 = F_sdsnew(m, v1270)
	mBase = m.M
	v1272 = m.ExcPending
	if v1272 != 0 {
		goto L50
	} else {
		goto L436
	}
L414:
	;
	v1254 = *(*int32)(unsafe.Add(mBase, uint32(v1197)+152))
	v1255 = F_createSentinelValkeyInstance(m, v1201, int32(4), v1203, v1253, v1254, v1197)
	mBase = m.M
	v1256 = m.ExcPending
	if v1256 != 0 {
		goto L50
	} else {
		goto L429
	}
L415:
	;
	v1213 = v1208 + int32(1)
	v1214 = int32(*(*int8)(unsafe.Add(mBase, uint32(v1208))))
	v1215 = F___isspace_1(m, v1214)
	mBase = m.M
	if v1215 != 0 {
		v1208 = v1213
		goto L415
	} else {
		goto L417
	}
L416:
	;
	v1216 = int32(1)
	switch v1214&int32(255) + int32(-43) {
	case 0:
		v1222 = v1216
		goto L419
	default:
		v1224 = v1208
		v1225 = v1214
		v1226 = v1216
		goto L418
	case 2:
		goto L420
	}
L417:
	;
	goto L416
L418:
	;
	v1229 = v1225 + int32(-48)
	if base.Ui32(int32(9)) < base.Ui32(v1229) {
		v1247 = int32(0)
		goto L421
	} else {
		goto L422
	}
L419:
	;
	v1223 = int32(*(*int8)(unsafe.Add(mBase, uint32(v1213))))
	v1224 = v1213
	v1225 = v1223
	v1226 = v1222
	goto L418
L420:
	;
	v1222 = int32(0)
	goto L419
L421:
	;
	if v1226 != 0 {
		goto L426
	} else {
		goto L427
	}
L422:
	;
	v1233 = int32(0)
	v1234 = v1224
	v1235 = v1229
	goto L423
L423:
	;
	v1237 = int32(10)
	v1239 = v1233*v1237 - v1235
	v1240 = int32(*(*int8)(unsafe.Add(mBase, uint32(v1234)+1)))
	v1244 = v1240 + int32(-48)
	if base.Ui32(v1244) < base.Ui32(v1237) {
		v1233 = v1239
		v1234 = v1234 + int32(1)
		v1235 = v1244
		goto L423
	} else {
		goto L425
	}
L424:
	;
	v1247 = v1239
	goto L421
L425:
	;
	goto L424
L426:
	;
	v1253 = int32(0) - v1247
	goto L428
L427:
	;
	v1253 = v1247
	goto L428
L428:
	;
	goto L414
L429:
	;
	if v1255 != 0 {
		goto L413
	} else {
		goto L430
	}
L430:
	;
	goto L434
L431:
	;
	return int32(_a_F_sentinelHandleConfiguration_3)
L432:
	;
	return int32(_a_F_sentinelHandleConfiguration_4)
L433:
	;
	if v1259 != int32(10) {
		goto L431
	} else {
		goto L435
	}
L434:
	;
	v1259 = *(*int32)(unsafe.Add(mBase, _c_F_sentinelHandleConfiguration[0]))
	switch v1259 + int32(-28) {
	case 0:
		goto L432
	case 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15:
		goto L431
	case 16:
		v1904 = int32(_a_F_sentinelHandleConfiguration_5)
		goto L13
	default:
		goto L433
	}
L435:
	;
	return int32(_a_F_sentinelHandleConfiguration_27)
L436:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1255)+8)) = v1271
	v1274 = F_sentinelTryConnectionSharing(m, v1255)
	mBase = m.M
	v1275 = m.ExcPending
	if v1275 != 0 {
		goto L50
	} else {
		goto L437
	}
L437:
	;
	goto L14
L438:
	;
	if l1 != int32(4) {
		goto L450
	} else {
		goto L451
	}
L439:
	;
	v1311 = F_tolower(m, v1307)
	mBase = m.M
	v1312 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1308))))
	v1313 = F_tolower(m, v1312)
	mBase = m.M
	goto L438
L440:
	;
	v1281 = v7
	v1282 = v1276
	v1283 = v1279
	goto L443
L441:
	;
	v1307 = int32(0)
	v1308 = v1276
	goto L439
L442:
	;
	v1307 = v1304 & int32(255)
	v1308 = v1303
	goto L439
L443:
	;
	v1285 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1282))))
	if v1285 == int32(0) {
		v1303 = v1282
		v1304 = v1283
		goto L442
	} else {
		goto L445
	}
L444:
	;
	v1303 = v1297
	v1304 = int32(0)
	goto L442
L445:
	;
	v1289 = v1283 & int32(255)
	if v1289 == v1285 {
		goto L446
	} else {
		goto L447
	}
L446:
	;
	v1296 = int32(1)
	v1297 = v1282 + v1296
	v1298 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1281)+1)))
	if v1298 != 0 {
		v1281 = v1281 + v1296
		v1282 = v1297
		v1283 = v1298
		goto L443
	} else {
		goto L449
	}
L447:
	;
	v1291 = F_tolower(m, v1289)
	mBase = m.M
	v1292 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1282))))
	v1293 = F_tolower(m, v1292)
	mBase = m.M
	if v1291 == v1293 {
		goto L446
	} else {
		goto L448
	}
L448:
	;
	v1295 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1281))))
	v1303 = v1282
	v1304 = v1295
	goto L442
L449:
	;
	goto L444
L450:
	;
	v1339 = int32(_a_F_sentinelHandleConfiguration_28)
	v1342 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7))))
	if v1342 != 0 {
		goto L464
	} else {
		goto L465
	}
L451:
	;
	if v1311-v1313 != 0 {
		goto L450
	} else {
		goto L452
	}
L452:
	;
	v1317 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v1318 = F_sentinelGetPrimaryByName(m, v1317)
	mBase = m.M
	v1319 = m.ExcPending
	if v1319 != 0 {
		goto L50
	} else {
		goto L454
	}
L453:
	;
	v1322 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1323 = F_sdsnew(m, v1322)
	mBase = m.M
	v1324 = m.ExcPending
	if v1324 != 0 {
		goto L50
	} else {
		goto L456
	}
L454:
	;
	if v1318 != 0 {
		goto L453
	} else {
		goto L455
	}
L455:
	;
	return int32(_a_F_sentinelHandleConfiguration_9)
L456:
	;
	v1325 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v1326 = F_sdsnew(m, v1325)
	mBase = m.M
	v1327 = m.ExcPending
	if v1327 != 0 {
		goto L50
	} else {
		goto L457
	}
L457:
	;
	v1328 = *(*int32)(unsafe.Add(mBase, uint32(v1318)+104))
	v1329 = F_dictAdd(m, v1328, v1323, v1326)
	mBase = m.M
	v1330 = m.ExcPending
	if v1330 != 0 {
		goto L50
	} else {
		goto L458
	}
L458:
	;
	if v1329 == int32(0) {
		goto L14
	} else {
		goto L459
	}
L459:
	;
	F_sdsfree(m, v1323)
	mBase = m.M
	v1334 = m.ExcPending
	if v1334 != 0 {
		goto L50
	} else {
		goto L460
	}
L460:
	;
	F_sdsfree(m, v1326)
	mBase = m.M
	v1336 = m.ExcPending
	if v1336 != 0 {
		goto L50
	} else {
		goto L461
	}
L461:
	;
	return int32(_a_F_sentinelHandleConfiguration_29)
L462:
	;
	v1379 = base.B2i32(l1 != int32(2))
	if l1 != int32(2) {
		goto L474
	} else {
		goto L475
	}
L463:
	;
	v1374 = F_tolower(m, v1370)
	mBase = m.M
	v1375 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1371))))
	v1376 = F_tolower(m, v1375)
	mBase = m.M
	goto L462
L464:
	;
	v1344 = v7
	v1345 = v1339
	v1346 = v1342
	goto L467
L465:
	;
	v1370 = int32(0)
	v1371 = v1339
	goto L463
L466:
	;
	v1370 = v1367 & int32(255)
	v1371 = v1366
	goto L463
L467:
	;
	v1348 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1345))))
	if v1348 == int32(0) {
		v1366 = v1345
		v1367 = v1346
		goto L466
	} else {
		goto L469
	}
L468:
	;
	v1366 = v1360
	v1367 = int32(0)
	goto L466
L469:
	;
	v1352 = v1346 & int32(255)
	if v1352 == v1348 {
		goto L470
	} else {
		goto L471
	}
L470:
	;
	v1359 = int32(1)
	v1360 = v1345 + v1359
	v1361 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1344)+1)))
	if v1361 != 0 {
		v1344 = v1344 + v1359
		v1345 = v1360
		v1346 = v1361
		goto L467
	} else {
		goto L473
	}
L471:
	;
	v1354 = F_tolower(m, v1352)
	mBase = m.M
	v1355 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1345))))
	v1356 = F_tolower(m, v1355)
	mBase = m.M
	if v1354 == v1356 {
		goto L470
	} else {
		goto L472
	}
L472:
	;
	v1358 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1344))))
	v1366 = v1345
	v1367 = v1358
	goto L466
L473:
	;
	goto L468
L474:
	;
	v1388 = int32(_a_F_sentinelHandleConfiguration_30)
	v1391 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7))))
	if v1391 != 0 {
		goto L481
	} else {
		goto L482
	}
L475:
	;
	if v1374-v1376 != 0 {
		goto L474
	} else {
		goto L476
	}
L476:
	;
	v1380 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v1381 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1380))))
	if v1381 == int32(0) {
		goto L14
	} else {
		goto L477
	}
L477:
	;
	v1385 = F_sdsnew(m, v1380)
	mBase = m.M
	v1386 = m.ExcPending
	if v1386 != 0 {
		goto L50
	} else {
		goto L478
	}
L478:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_sentinelHandleConfiguration[8])) = v1385
	goto L14
L479:
	;
	if l1 != int32(2) {
		goto L491
	} else {
		goto L492
	}
L480:
	;
	v1423 = F_tolower(m, v1419)
	mBase = m.M
	v1424 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1420))))
	v1425 = F_tolower(m, v1424)
	mBase = m.M
	goto L479
L481:
	;
	v1393 = v7
	v1394 = v1388
	v1395 = v1391
	goto L484
L482:
	;
	v1419 = int32(0)
	v1420 = v1388
	goto L480
L483:
	;
	v1419 = v1416 & int32(255)
	v1420 = v1415
	goto L480
L484:
	;
	v1397 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1394))))
	if v1397 == int32(0) {
		v1415 = v1394
		v1416 = v1395
		goto L483
	} else {
		goto L486
	}
L485:
	;
	v1415 = v1409
	v1416 = int32(0)
	goto L483
L486:
	;
	v1401 = v1395 & int32(255)
	if v1401 == v1397 {
		goto L487
	} else {
		goto L488
	}
L487:
	;
	v1408 = int32(1)
	v1409 = v1394 + v1408
	v1410 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1393)+1)))
	if v1410 != 0 {
		v1393 = v1393 + v1408
		v1394 = v1409
		v1395 = v1410
		goto L484
	} else {
		goto L490
	}
L488:
	;
	v1403 = F_tolower(m, v1401)
	mBase = m.M
	v1404 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1394))))
	v1405 = F_tolower(m, v1404)
	mBase = m.M
	if v1403 == v1405 {
		goto L487
	} else {
		goto L489
	}
L489:
	;
	v1407 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1393))))
	v1415 = v1394
	v1416 = v1407
	goto L483
L490:
	;
	goto L485
L491:
	;
	v1479 = int32(_a_F_sentinelHandleConfiguration_31)
	v1482 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7))))
	if v1482 != 0 {
		goto L511
	} else {
		goto L512
	}
L492:
	;
	if v1423-v1425 != 0 {
		goto L491
	} else {
		goto L493
	}
L493:
	;
	v1428 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v1432 = v1428
	goto L495
L494:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_sentinelHandleConfiguration[9])) = v1477
	goto L14
L495:
	;
	v1437 = v1432 + int32(1)
	v1438 = int32(*(*int8)(unsafe.Add(mBase, uint32(v1432))))
	v1439 = F___isspace_1(m, v1438)
	mBase = m.M
	if v1439 != 0 {
		v1432 = v1437
		goto L495
	} else {
		goto L497
	}
L496:
	;
	v1440 = int32(1)
	switch v1438&int32(255) + int32(-43) {
	case 0:
		v1446 = v1440
		goto L499
	default:
		v1448 = v1432
		v1449 = v1438
		v1450 = v1440
		goto L498
	case 2:
		goto L500
	}
L497:
	;
	goto L496
L498:
	;
	v1453 = v1449 + int32(-48)
	if base.Ui32(int32(9)) < base.Ui32(v1453) {
		v1471 = int32(0)
		goto L501
	} else {
		goto L502
	}
L499:
	;
	v1447 = int32(*(*int8)(unsafe.Add(mBase, uint32(v1437))))
	v1448 = v1437
	v1449 = v1447
	v1450 = v1446
	goto L498
L500:
	;
	v1446 = int32(0)
	goto L499
L501:
	;
	if v1450 != 0 {
		goto L506
	} else {
		goto L507
	}
L502:
	;
	v1457 = int32(0)
	v1458 = v1448
	v1459 = v1453
	goto L503
L503:
	;
	v1461 = int32(10)
	v1463 = v1457*v1461 - v1459
	v1464 = int32(*(*int8)(unsafe.Add(mBase, uint32(v1458)+1)))
	v1468 = v1464 + int32(-48)
	if base.Ui32(v1468) < base.Ui32(v1461) {
		v1457 = v1463
		v1458 = v1458 + int32(1)
		v1459 = v1468
		goto L503
	} else {
		goto L505
	}
L504:
	;
	v1471 = v1463
	goto L501
L505:
	;
	goto L504
L506:
	;
	v1477 = int32(0) - v1471
	goto L508
L507:
	;
	v1477 = v1471
	goto L508
L508:
	;
	goto L494
L509:
	;
	v1519 = base.B2i32(l1 != int32(2))
	if l1 != int32(2) {
		goto L521
	} else {
		goto L522
	}
L510:
	;
	v1514 = F_tolower(m, v1510)
	mBase = m.M
	v1515 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1511))))
	v1516 = F_tolower(m, v1515)
	mBase = m.M
	goto L509
L511:
	;
	v1484 = v7
	v1485 = v1479
	v1486 = v1482
	goto L514
L512:
	;
	v1510 = int32(0)
	v1511 = v1479
	goto L510
L513:
	;
	v1510 = v1507 & int32(255)
	v1511 = v1506
	goto L510
L514:
	;
	v1488 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1485))))
	if v1488 == int32(0) {
		v1506 = v1485
		v1507 = v1486
		goto L513
	} else {
		goto L516
	}
L515:
	;
	v1506 = v1500
	v1507 = int32(0)
	goto L513
L516:
	;
	v1492 = v1486 & int32(255)
	if v1492 == v1488 {
		goto L517
	} else {
		goto L518
	}
L517:
	;
	v1499 = int32(1)
	v1500 = v1485 + v1499
	v1501 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1484)+1)))
	if v1501 != 0 {
		v1484 = v1484 + v1499
		v1485 = v1500
		v1486 = v1501
		goto L514
	} else {
		goto L520
	}
L518:
	;
	v1494 = F_tolower(m, v1492)
	mBase = m.M
	v1495 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1485))))
	v1496 = F_tolower(m, v1495)
	mBase = m.M
	if v1494 == v1496 {
		goto L517
	} else {
		goto L519
	}
L519:
	;
	v1498 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1484))))
	v1506 = v1485
	v1507 = v1498
	goto L513
L520:
	;
	goto L515
L521:
	;
	v1536 = int32(_a_F_sentinelHandleConfiguration_32)
	v1539 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7))))
	if v1539 != 0 {
		goto L533
	} else {
		goto L534
	}
L522:
	;
	if v1514-v1516 != 0 {
		goto L521
	} else {
		goto L523
	}
L523:
	;
	v1521 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v1523 = F_strcasecmp(m, v1521, int32(_a_F_sentinelHandleConfiguration_33))
	mBase = m.M
	if v1523 != 0 {
		goto L525
	} else {
		goto L526
	}
L524:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_sentinelHandleConfiguration[10])) = v1530
	if v1530 != int32(-1) {
		goto L14
	} else {
		goto L530
	}
L525:
	;
	v1528 = F_strcasecmp(m, v1521, int32(_a_F_sentinelHandleConfiguration_34))
	mBase = m.M
	if v1528 != 0 {
		goto L527
	} else {
		goto L528
	}
L526:
	;
	v1530 = int32(1)
	goto L524
L527:
	;
	v1529 = int32(-1)
	goto L529
L528:
	;
	v1529 = int32(0)
	goto L529
L529:
	;
	v1530 = v1529
	goto L524
L530:
	;
	return int32(_a_F_sentinelHandleConfiguration_35)
L531:
	;
	if l1 != int32(2) {
		goto L543
	} else {
		goto L544
	}
L532:
	;
	v1571 = F_tolower(m, v1567)
	mBase = m.M
	v1572 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1568))))
	v1573 = F_tolower(m, v1572)
	mBase = m.M
	goto L531
L533:
	;
	v1541 = v7
	v1542 = v1536
	v1543 = v1539
	goto L536
L534:
	;
	v1567 = int32(0)
	v1568 = v1536
	goto L532
L535:
	;
	v1567 = v1564 & int32(255)
	v1568 = v1563
	goto L532
L536:
	;
	v1545 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1542))))
	if v1545 == int32(0) {
		v1563 = v1542
		v1564 = v1543
		goto L535
	} else {
		goto L538
	}
L537:
	;
	v1563 = v1557
	v1564 = int32(0)
	goto L535
L538:
	;
	v1549 = v1543 & int32(255)
	if v1549 == v1545 {
		goto L539
	} else {
		goto L540
	}
L539:
	;
	v1556 = int32(1)
	v1557 = v1542 + v1556
	v1558 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1541)+1)))
	if v1558 != 0 {
		v1541 = v1541 + v1556
		v1542 = v1557
		v1543 = v1558
		goto L536
	} else {
		goto L542
	}
L540:
	;
	v1551 = F_tolower(m, v1549)
	mBase = m.M
	v1552 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1542))))
	v1553 = F_tolower(m, v1552)
	mBase = m.M
	if v1551 == v1553 {
		goto L539
	} else {
		goto L541
	}
L541:
	;
	v1555 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1541))))
	v1563 = v1542
	v1564 = v1555
	goto L535
L542:
	;
	goto L537
L543:
	;
	v1583 = int32(_a_F_sentinelHandleConfiguration_36)
	v1586 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7))))
	if v1586 != 0 {
		goto L550
	} else {
		goto L551
	}
L544:
	;
	if v1571-v1573 != 0 {
		goto L543
	} else {
		goto L545
	}
L545:
	;
	v1575 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v1576 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1575))))
	if v1576 == int32(0) {
		goto L14
	} else {
		goto L546
	}
L546:
	;
	v1580 = F_sdsnew(m, v1575)
	mBase = m.M
	v1581 = m.ExcPending
	if v1581 != 0 {
		goto L50
	} else {
		goto L547
	}
L547:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_sentinelHandleConfiguration[11])) = v1580
	goto L14
L548:
	;
	v1623 = base.B2i32(l1 != int32(2))
	if l1 != int32(2) {
		goto L560
	} else {
		goto L561
	}
L549:
	;
	v1618 = F_tolower(m, v1614)
	mBase = m.M
	v1619 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1615))))
	v1620 = F_tolower(m, v1619)
	mBase = m.M
	goto L548
L550:
	;
	v1588 = v7
	v1589 = v1583
	v1590 = v1586
	goto L553
L551:
	;
	v1614 = int32(0)
	v1615 = v1583
	goto L549
L552:
	;
	v1614 = v1611 & int32(255)
	v1615 = v1610
	goto L549
L553:
	;
	v1592 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1589))))
	if v1592 == int32(0) {
		v1610 = v1589
		v1611 = v1590
		goto L552
	} else {
		goto L555
	}
L554:
	;
	v1610 = v1604
	v1611 = int32(0)
	goto L552
L555:
	;
	v1596 = v1590 & int32(255)
	if v1596 == v1592 {
		goto L556
	} else {
		goto L557
	}
L556:
	;
	v1603 = int32(1)
	v1604 = v1589 + v1603
	v1605 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1588)+1)))
	if v1605 != 0 {
		v1588 = v1588 + v1603
		v1589 = v1604
		v1590 = v1605
		goto L553
	} else {
		goto L559
	}
L557:
	;
	v1598 = F_tolower(m, v1596)
	mBase = m.M
	v1599 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1589))))
	v1600 = F_tolower(m, v1599)
	mBase = m.M
	if v1598 == v1600 {
		goto L556
	} else {
		goto L558
	}
L558:
	;
	v1602 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1588))))
	v1610 = v1589
	v1611 = v1602
	goto L552
L559:
	;
	goto L554
L560:
	;
	v1632 = int32(_a_F_sentinelHandleConfiguration_37)
	v1635 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7))))
	if v1635 != 0 {
		goto L567
	} else {
		goto L568
	}
L561:
	;
	if v1618-v1620 != 0 {
		goto L560
	} else {
		goto L562
	}
L562:
	;
	v1624 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v1625 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1624))))
	if v1625 == int32(0) {
		goto L14
	} else {
		goto L563
	}
L563:
	;
	v1629 = F_sdsnew(m, v1624)
	mBase = m.M
	v1630 = m.ExcPending
	if v1630 != 0 {
		goto L50
	} else {
		goto L564
	}
L564:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_sentinelHandleConfiguration[12])) = v1629
	goto L14
L565:
	;
	if l1 != int32(2) {
		goto L577
	} else {
		goto L578
	}
L566:
	;
	v1667 = F_tolower(m, v1663)
	mBase = m.M
	v1668 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1664))))
	v1669 = F_tolower(m, v1668)
	mBase = m.M
	goto L565
L567:
	;
	v1637 = v7
	v1638 = v1632
	v1639 = v1635
	goto L570
L568:
	;
	v1663 = int32(0)
	v1664 = v1632
	goto L566
L569:
	;
	v1663 = v1660 & int32(255)
	v1664 = v1659
	goto L566
L570:
	;
	v1641 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1638))))
	if v1641 == int32(0) {
		v1659 = v1638
		v1660 = v1639
		goto L569
	} else {
		goto L572
	}
L571:
	;
	v1659 = v1653
	v1660 = int32(0)
	goto L569
L572:
	;
	v1645 = v1639 & int32(255)
	if v1645 == v1641 {
		goto L573
	} else {
		goto L574
	}
L573:
	;
	v1652 = int32(1)
	v1653 = v1638 + v1652
	v1654 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1637)+1)))
	if v1654 != 0 {
		v1637 = v1637 + v1652
		v1638 = v1653
		v1639 = v1654
		goto L570
	} else {
		goto L576
	}
L574:
	;
	v1647 = F_tolower(m, v1645)
	mBase = m.M
	v1648 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1638))))
	v1649 = F_tolower(m, v1648)
	mBase = m.M
	if v1647 == v1649 {
		goto L573
	} else {
		goto L575
	}
L575:
	;
	v1651 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1637))))
	v1659 = v1638
	v1660 = v1651
	goto L569
L576:
	;
	goto L571
L577:
	;
	v1687 = int32(_a_F_sentinelHandleConfiguration_38)
	v1690 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7))))
	if v1690 != 0 {
		goto L589
	} else {
		goto L590
	}
L578:
	;
	if v1667-v1669 != 0 {
		goto L577
	} else {
		goto L579
	}
L579:
	;
	v1672 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v1674 = F_strcasecmp(m, v1672, int32(_a_F_sentinelHandleConfiguration_33))
	mBase = m.M
	if v1674 != 0 {
		goto L581
	} else {
		goto L582
	}
L580:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_sentinelHandleConfiguration[13])) = v1681
	if v1681 != int32(-1) {
		goto L14
	} else {
		goto L586
	}
L581:
	;
	v1679 = F_strcasecmp(m, v1672, int32(_a_F_sentinelHandleConfiguration_34))
	mBase = m.M
	if v1679 != 0 {
		goto L583
	} else {
		goto L584
	}
L582:
	;
	v1681 = int32(1)
	goto L580
L583:
	;
	v1680 = int32(-1)
	goto L585
L584:
	;
	v1680 = int32(0)
	goto L585
L585:
	;
	v1681 = v1680
	goto L580
L586:
	;
	return int32(_a_F_sentinelHandleConfiguration_39)
L587:
	;
	if l1 != int32(2) {
		goto L599
	} else {
		goto L600
	}
L588:
	;
	v1722 = F_tolower(m, v1718)
	mBase = m.M
	v1723 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1719))))
	v1724 = F_tolower(m, v1723)
	mBase = m.M
	goto L587
L589:
	;
	v1692 = v7
	v1693 = v1687
	v1694 = v1690
	goto L592
L590:
	;
	v1718 = int32(0)
	v1719 = v1687
	goto L588
L591:
	;
	v1718 = v1715 & int32(255)
	v1719 = v1714
	goto L588
L592:
	;
	v1696 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1693))))
	if v1696 == int32(0) {
		v1714 = v1693
		v1715 = v1694
		goto L591
	} else {
		goto L594
	}
L593:
	;
	v1714 = v1708
	v1715 = int32(0)
	goto L591
L594:
	;
	v1700 = v1694 & int32(255)
	if v1700 == v1696 {
		goto L595
	} else {
		goto L596
	}
L595:
	;
	v1707 = int32(1)
	v1708 = v1693 + v1707
	v1709 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1692)+1)))
	if v1709 != 0 {
		v1692 = v1692 + v1707
		v1693 = v1708
		v1694 = v1709
		goto L592
	} else {
		goto L598
	}
L596:
	;
	v1702 = F_tolower(m, v1700)
	mBase = m.M
	v1703 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1693))))
	v1704 = F_tolower(m, v1703)
	mBase = m.M
	if v1702 == v1704 {
		goto L595
	} else {
		goto L597
	}
L597:
	;
	v1706 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1692))))
	v1714 = v1693
	v1715 = v1706
	goto L591
L598:
	;
	goto L593
L599:
	;
	v1744 = int32(_a_F_sentinelHandleConfiguration_40)
	v1747 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7))))
	if v1747 != 0 {
		goto L613
	} else {
		goto L614
	}
L600:
	;
	if v1722-v1724 != 0 {
		goto L599
	} else {
		goto L601
	}
L601:
	;
	v1729 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v1731 = F_strcasecmp(m, v1729, int32(_a_F_sentinelHandleConfiguration_33))
	mBase = m.M
	if v1731 != 0 {
		goto L603
	} else {
		goto L604
	}
L602:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_sentinelHandleConfiguration[14])) = v1738
	if v1738 != int32(-1) {
		goto L14
	} else {
		goto L608
	}
L603:
	;
	v1736 = F_strcasecmp(m, v1729, int32(_a_F_sentinelHandleConfiguration_34))
	mBase = m.M
	if v1736 != 0 {
		goto L605
	} else {
		goto L606
	}
L604:
	;
	v1738 = int32(1)
	goto L602
L605:
	;
	v1737 = int32(-1)
	goto L607
L606:
	;
	v1737 = int32(0)
	goto L607
L607:
	;
	v1738 = v1737
	goto L602
L608:
	;
	return int32(_a_F_sentinelHandleConfiguration_41)
L609:
	;
	v1833 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v1834 = F_sentinelGetPrimaryByName(m, v1833)
	mBase = m.M
	v1835 = m.ExcPending
	if v1835 != 0 {
		goto L50
	} else {
		goto L640
	}
L610:
	;
	if l1 == int32(3) {
		goto L609
	} else {
		goto L638
	}
L611:
	;
	if v1779-v1781 == int32(0) {
		goto L610
	} else {
		goto L623
	}
L612:
	;
	v1779 = F_tolower(m, v1775)
	mBase = m.M
	v1780 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1776))))
	v1781 = F_tolower(m, v1780)
	mBase = m.M
	goto L611
L613:
	;
	v1749 = v7
	v1750 = v1744
	v1751 = v1747
	goto L616
L614:
	;
	v1775 = int32(0)
	v1776 = v1744
	goto L612
L615:
	;
	v1775 = v1772 & int32(255)
	v1776 = v1771
	goto L612
L616:
	;
	v1753 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1750))))
	if v1753 == int32(0) {
		v1771 = v1750
		v1772 = v1751
		goto L615
	} else {
		goto L618
	}
L617:
	;
	v1771 = v1765
	v1772 = int32(0)
	goto L615
L618:
	;
	v1757 = v1751 & int32(255)
	if v1757 == v1753 {
		goto L619
	} else {
		goto L620
	}
L619:
	;
	v1764 = int32(1)
	v1765 = v1750 + v1764
	v1766 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1749)+1)))
	if v1766 != 0 {
		v1749 = v1749 + v1764
		v1750 = v1765
		v1751 = v1766
		goto L616
	} else {
		goto L622
	}
L620:
	;
	v1759 = F_tolower(m, v1757)
	mBase = m.M
	v1760 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1750))))
	v1761 = F_tolower(m, v1760)
	mBase = m.M
	if v1759 == v1761 {
		goto L619
	} else {
		goto L621
	}
L621:
	;
	v1763 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1749))))
	v1771 = v1750
	v1772 = v1763
	goto L615
L622:
	;
	goto L617
L623:
	;
	v1785 = int32(_a_F_sentinelHandleConfiguration_42)
	v1786 = int32(_a_F_sentinelHandleConfiguration_43)
	v1789 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7))))
	if v1789 != 0 {
		goto L626
	} else {
		goto L627
	}
L624:
	;
	if l1 != int32(3) {
		v1904 = v1785
		goto L13
	} else {
		goto L636
	}
L625:
	;
	v1821 = F_tolower(m, v1817)
	mBase = m.M
	v1822 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1818))))
	v1823 = F_tolower(m, v1822)
	mBase = m.M
	goto L624
L626:
	;
	v1791 = v7
	v1792 = v1786
	v1793 = v1789
	goto L629
L627:
	;
	v1817 = int32(0)
	v1818 = v1786
	goto L625
L628:
	;
	v1817 = v1814 & int32(255)
	v1818 = v1813
	goto L625
L629:
	;
	v1795 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1792))))
	if v1795 == int32(0) {
		v1813 = v1792
		v1814 = v1793
		goto L628
	} else {
		goto L631
	}
L630:
	;
	v1813 = v1807
	v1814 = int32(0)
	goto L628
L631:
	;
	v1799 = v1793 & int32(255)
	if v1799 == v1795 {
		goto L632
	} else {
		goto L633
	}
L632:
	;
	v1806 = int32(1)
	v1807 = v1792 + v1806
	v1808 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1791)+1)))
	if v1808 != 0 {
		v1791 = v1791 + v1806
		v1792 = v1807
		v1793 = v1808
		goto L629
	} else {
		goto L635
	}
L633:
	;
	v1801 = F_tolower(m, v1799)
	mBase = m.M
	v1802 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1792))))
	v1803 = F_tolower(m, v1802)
	mBase = m.M
	if v1801 == v1803 {
		goto L632
	} else {
		goto L634
	}
L634:
	;
	v1805 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1791))))
	v1813 = v1792
	v1814 = v1805
	goto L628
L635:
	;
	goto L630
L636:
	;
	if v1821-v1823 != 0 {
		v1904 = v1785
		goto L13
	} else {
		goto L637
	}
L637:
	;
	goto L609
L638:
	;
	return int32(_a_F_sentinelHandleConfiguration_42)
L639:
	;
	v1838 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1842 = v1838
	goto L643
L640:
	;
	if v1834 != 0 {
		goto L639
	} else {
		goto L641
	}
L641:
	;
	return int32(_a_F_sentinelHandleConfiguration_9)
L642:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v1834)+80)) = base.I64_extend_i32_s(v1887)
	if int32(0) <= v1887 {
		goto L14
	} else {
		goto L657
	}
L643:
	;
	v1847 = v1842 + int32(1)
	v1848 = int32(*(*int8)(unsafe.Add(mBase, uint32(v1842))))
	v1849 = F___isspace_1(m, v1848)
	mBase = m.M
	if v1849 != 0 {
		v1842 = v1847
		goto L643
	} else {
		goto L645
	}
L644:
	;
	v1850 = int32(1)
	switch v1848&int32(255) + int32(-43) {
	case 0:
		v1856 = v1850
		goto L647
	default:
		v1858 = v1842
		v1859 = v1848
		v1860 = v1850
		goto L646
	case 2:
		goto L648
	}
L645:
	;
	goto L644
L646:
	;
	v1863 = v1859 + int32(-48)
	if base.Ui32(int32(9)) < base.Ui32(v1863) {
		v1881 = int32(0)
		goto L649
	} else {
		goto L650
	}
L647:
	;
	v1857 = int32(*(*int8)(unsafe.Add(mBase, uint32(v1847))))
	v1858 = v1847
	v1859 = v1857
	v1860 = v1856
	goto L646
L648:
	;
	v1856 = int32(0)
	goto L647
L649:
	;
	if v1860 != 0 {
		goto L654
	} else {
		goto L655
	}
L650:
	;
	v1867 = int32(0)
	v1868 = v1858
	v1869 = v1863
	goto L651
L651:
	;
	v1871 = int32(10)
	v1873 = v1867*v1871 - v1869
	v1874 = int32(*(*int8)(unsafe.Add(mBase, uint32(v1868)+1)))
	v1878 = v1874 + int32(-48)
	if base.Ui32(v1878) < base.Ui32(v1871) {
		v1867 = v1873
		v1868 = v1868 + int32(1)
		v1869 = v1878
		goto L651
	} else {
		goto L653
	}
L652:
	;
	v1881 = v1873
	goto L649
L653:
	;
	goto L652
L654:
	;
	v1887 = int32(0) - v1881
	goto L656
L655:
	;
	v1887 = v1881
	goto L656
L656:
	;
	goto L642
L657:
	;
	return int32(_a_F_sentinelHandleConfiguration_44)
}
func F_sentinelHandleValkeyInstance(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v10 int64
	_ = v10
	var v11 int32
	_ = v11
	var v12 int64
	_ = v12
	var v15 int64
	_ = v15
	var v17 int32
	_ = v17
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
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
	F_sentinelReconnectInstance(m, l0)
	mBase = m.M
	v3 = m.ExcPending
	if v3 != 0 {
		return
	} else {
		F_sentinelSendPeriodicCommands(m, l0)
		mBase = m.M
		v5 = m.ExcPending
		if v5 != 0 {
			return
		} else {
			v6 = int32(0)
			v7 = *(*int32)(unsafe.Add(mBase, _c_F_sentinelHandleValkeyInstance[0]))
			if v7 == v6 {
				F_sentinelCheckSubjectivelyDown(m, l0)
				mBase = m.M
				v28 = m.ExcPending
				if v28 != 0 {
					return
				} else {
					v29 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
					if v29&int32(1) == int32(0) {
						return
					} else {
						F_sentinelCheckObjectivelyDown(m, l0)
						mBase = m.M
						v35 = m.ExcPending
						if v35 != 0 {
							return
						} else {
							v36 = F_sentinelStartFailoverIfNeeded(m, l0)
							mBase = m.M
							v37 = m.ExcPending
							if v37 != 0 {
								return
							} else {
								if v36 == int32(0) {
									F_sentinelFailoverStateMachine(m, l0)
									mBase = m.M
									v44 = m.ExcPending
									if v44 != 0 {
										return
									} else {
										F_sentinelAskPrimaryStateToOtherSentinels(m, l0, int32(0))
										mBase = m.M
										v47 = m.ExcPending
										if v47 != 0 {
											return
										} else {
											return
										}
									}
								} else {
									F_sentinelAskPrimaryStateToOtherSentinels(m, l0, int32(1))
									mBase = m.M
									v42 = m.ExcPending
									if v42 != 0 {
										return
									} else {
										F_sentinelFailoverStateMachine(m, l0)
										mBase = m.M
										v44 = m.ExcPending
										if v44 != 0 {
											return
										} else {
											F_sentinelAskPrimaryStateToOtherSentinels(m, l0, int32(0))
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
				}
			} else {
				v10 = F_mstime(m)
				mBase = m.M
				v11 = int32(0)
				v12 = *(*int64)(unsafe.Add(mBase, _c_F_sentinelHandleValkeyInstance[1]))
				v15 = *(*int64)(unsafe.Add(mBase, _c_F_sentinelHandleValkeyInstance[2]))
				if v10-v12 < v15 {
					return
				} else {
					v17 = int32(0)
					*(*int32)(unsafe.Add(mBase, _c_F_sentinelHandleValkeyInstance[0])) = v17
					F_sentinelEvent(m, int32(3), int32(_a_F_sentinelHandleValkeyInstance_0), v17, int32(_a_F_sentinelHandleValkeyInstance_1), v17)
					mBase = m.M
					v26 = m.ExcPending
					if v26 != 0 {
						return
					} else {
						F_sentinelCheckSubjectivelyDown(m, l0)
						mBase = m.M
						v28 = m.ExcPending
						if v28 != 0 {
							return
						} else {
							v29 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
							if v29&int32(1) == int32(0) {
								return
							} else {
								F_sentinelCheckObjectivelyDown(m, l0)
								mBase = m.M
								v35 = m.ExcPending
								if v35 != 0 {
									return
								} else {
									v36 = F_sentinelStartFailoverIfNeeded(m, l0)
									mBase = m.M
									v37 = m.ExcPending
									if v37 != 0 {
										return
									} else {
										if v36 == int32(0) {
											F_sentinelFailoverStateMachine(m, l0)
											mBase = m.M
											v44 = m.ExcPending
											if v44 != 0 {
												return
											} else {
												F_sentinelAskPrimaryStateToOtherSentinels(m, l0, int32(0))
												mBase = m.M
												v47 = m.ExcPending
												if v47 != 0 {
													return
												} else {
													return
												}
											}
										} else {
											F_sentinelAskPrimaryStateToOtherSentinels(m, l0, int32(1))
											mBase = m.M
											v42 = m.ExcPending
											if v42 != 0 {
												return
											} else {
												F_sentinelFailoverStateMachine(m, l0)
												mBase = m.M
												v44 = m.ExcPending
												if v44 != 0 {
													return
												} else {
													F_sentinelAskPrimaryStateToOtherSentinels(m, l0, int32(0))
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
						}
					}
				}
			}
		}
	}
}
func F_sentinelLeaderIncr(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v18 int64
	_ = v18
	var v19 int32
	_ = v19
	var v21 int64
	_ = v21
	var v29 int32
	_ = v29
	var v39 int32
	_ = v39
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v11 = F_dictAddRaw(m, l0, l1, v7+int32(12))
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		return int32(0)
	} else {
		v15 = *(*int32)(unsafe.Add(mBase, uint32(v7)+12))
		if v15 == int32(0) {
			if v11 == int32(0) {
				F__serverAssert(m, int32(_a_F_sentinelLeaderIncr_0), int32(_a_F_sentinelLeaderIncr_1), int32(4678))
				mBase = m.M
				v39 = m.ExcPending
				if v39 != 0 {
					return int32(0)
				} else {
					F_abort(m)
					mBase = m.M
					base.Wasm_trap_unreachable()
					for {
					}
				}
			} else {
				*(*int64)(unsafe.Add(mBase, uint32(v11)+8)) = int64(1)
				v29 = int32(1)
				m.G0 = v7 + int32(16)
				return v29
			}
		} else {
			v18 = *(*int64)(unsafe.Add(mBase, uint32(v15)+8))
			v19 = *(*int32)(unsafe.Add(mBase, uint32(v7)+12))
			v21 = v18 + int64(1)
			*(*int64)(unsafe.Add(mBase, uint32(v19)+8)) = v21
			v29 = base.I32_wrap_i64(v21)
			m.G0 = v7 + int32(16)
			return v29
		}
	}
}
func F_sentinelProcessHelloMessage(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v92 int32
	_ = v92
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
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
	var v113 int32
	_ = v113
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v121 int32
	_ = v121
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v128 int32
	_ = v128
	var v131 int32
	_ = v131
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v147 int64
	_ = v147
	var v148 int32
	_ = v148
	var v152 int64
	_ = v152
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v214 int32
	_ = v214
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v222 int32
	_ = v222
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v231 int32
	_ = v231
	var v233 int32
	_ = v233
	var v235 int64
	_ = v235
	var v236 int64
	_ = v236
	var v237 int64
	_ = v237
	var v238 int64
	_ = v238
	var v239 int64
	_ = v239
	var v240 int64
	_ = v240
	var v241 int64
	_ = v241
	var v243 int64
	_ = v243
	var v245 int64
	_ = v245
	var v247 int64
	_ = v247
	var v249 int64
	_ = v249
	var v251 int64
	_ = v251
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v255 int32
	_ = v255
	var v258 int32
	_ = v258
	var v261 int32
	_ = v261
	var v262 int32
	_ = v262
	var v263 int32
	_ = v263
	var v265 int32
	_ = v265
	var v268 int32
	_ = v268
	var v269 int32
	_ = v269
	var v273 int32
	_ = v273
	var v277 int32
	_ = v277
	var v279 int32
	_ = v279
	var v286 int32
	_ = v286
	var v287 int32
	_ = v287
	var v288 int32
	_ = v288
	var v293 int32
	_ = v293
	var v298 int32
	_ = v298
	var v302 int32
	_ = v302
	var v306 int32
	_ = v306
	var v312 int32
	_ = v312
	var v313 int32
	_ = v313
	var v314 int32
	_ = v314
	var v316 int32
	_ = v316
	var v318 int32
	_ = v318
	var v332 int32
	_ = v332
	var v334 int32
	_ = v334
	var v335 int32
	_ = v335
	var v336 int32
	_ = v336
	var v337 int32
	_ = v337
	var v344 int32
	_ = v344
	var v345 int32
	_ = v345
	var v346 int32
	_ = v346
	var v347 int32
	_ = v347
	var v349 int32
	_ = v349
	var v350 int32
	_ = v350
	var v351 int32
	_ = v351
	var v352 int32
	_ = v352
	var v353 int32
	_ = v353
	var v355 int32
	_ = v355
	var v356 int32
	_ = v356
	var v357 int32
	_ = v357
	var v358 int32
	_ = v358
	var v359 int32
	_ = v359
	var v360 int32
	_ = v360
	var v365 int32
	_ = v365
	var v367 int32
	_ = v367
	var v368 int32
	_ = v368
	var v369 int32
	_ = v369
	var v372 int32
	_ = v372
	var v378 int32
	_ = v378
	var v379 int32
	_ = v379
	var v386 int32
	_ = v386
	var v393 int32
	_ = v393
	var v397 int32
	_ = v397
	var v408 int64
	_ = v408
	var v410 int32
	_ = v410
	var v412 int32
	_ = v412
	var v413 int32
	_ = v413
	var v418 int32
	_ = v418
	var v420 int32
	_ = v420
	var v421 int32
	_ = v421
	var v422 int32
	_ = v422
	var v425 int32
	_ = v425
	var v431 int32
	_ = v431
	var v432 int32
	_ = v432
	var v439 int32
	_ = v439
	var v446 int32
	_ = v446
	var v448 int64
	_ = v448
	var v456 int32
	_ = v456
	var v461 int64
	_ = v461
	var v464 int32
	_ = v464
	var v465 int32
	_ = v465
	var v467 int32
	_ = v467
	var v468 int32
	_ = v468
	var v473 int32
	_ = v473
	var v476 int32
	_ = v476
	var v477 int32
	_ = v477
	var v478 int32
	_ = v478
	var v479 int32
	_ = v479
	var v486 int32
	_ = v486
	var v487 int32
	_ = v487
	var v489 int32
	_ = v489
	var v492 int32
	_ = v492
	var v495 int32
	_ = v495
	var v497 int32
	_ = v497
	var v498 int32
	_ = v498
	var v499 int32
	_ = v499
	var v501 int32
	_ = v501
	var v505 int32
	_ = v505
	var v507 int32
	_ = v507
	var v508 int32
	_ = v508
	var v509 int32
	_ = v509
	var v511 int32
	_ = v511
	var v512 int32
	_ = v512
	var v513 int32
	_ = v513
	var v514 int32
	_ = v514
	var v519 int32
	_ = v519
	var v520 int32
	_ = v520
	var v523 int32
	_ = v523
	var v524 int32
	_ = v524
	var v527 int32
	_ = v527
	var v528 int32
	_ = v528
	var v529 int32
	_ = v529
	var v540 int32
	_ = v540
	var v541 int32
	_ = v541
	var v542 int32
	_ = v542
	var v543 int32
	_ = v543
	var v549 int32
	_ = v549
	var v550 int32
	_ = v550
	var v551 int32
	_ = v551
	var v552 int32
	_ = v552
	var v568 int32
	_ = v568
	var v569 int32
	_ = v569
	var v571 int32
	_ = v571
	var v572 int32
	_ = v572
	var v573 int32
	_ = v573
	var v574 int32
	_ = v574
	var v575 int32
	_ = v575
	var v577 int32
	_ = v577
	var v578 int32
	_ = v578
	var v579 int32
	_ = v579
	var v581 int32
	_ = v581
	var v583 int32
	_ = v583
	var v584 int32
	_ = v584
	var v585 int32
	_ = v585
	var v586 int32
	_ = v586
	var v589 int32
	_ = v589
	var v591 int32
	_ = v591
	var v593 int64
	_ = v593
	var v602 int32
	_ = v602
	var v611 int32
	_ = v611
	var v612 int32
	_ = v612
	var v613 int64
	_ = v613
	var v615 int32
	_ = v615
	var v636 int32
	_ = v636
	var v638 int64
	_ = v638
	var v647 int32
	_ = v647
	var v656 int32
	_ = v656
	var v657 int32
	_ = v657
	var v658 int64
	_ = v658
	var v660 int32
	_ = v660
	var v680 int32
	_ = v680
	var v681 int32
	_ = v681
	var v685 int32
	_ = v685
	var v687 int32
	_ = v687
	var v688 int32
	_ = v688
	var v689 int32
	_ = v689
	var v693 int32
	_ = v693
	var v716 int32
	_ = v716
	var v723 int32
	_ = v723
	var v725 int32
	_ = v725
	var v727 int32
	_ = v727
	var v736 int64
	_ = v736
	var v751 int32
	_ = v751
	var v753 int32
	_ = v753
	v14 = m.G0
	v16 = v14 - int32(208)
	m.G0 = v16
	v22 = F_sdssplitlen(m, l0, l1, int32(_a_F_sentinelProcessHelloMessage_0), int32(1), v16+int32(124))
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v16)+124))
	if v24 != int32(8) {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v751 = *(*int32)(unsafe.Add(mBase, uint32(v16)+124))
	F_sdsfreesplitres(m, v22, v751)
	mBase = m.M
	v753 = m.ExcPending
	if v753 != 0 {
		goto L1
	} else {
		goto L177
	}
L4:
	;
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v22)+16))
	v28 = F_sdsnew(m, v27)
	mBase = m.M
	v29 = m.ExcPending
	if v29 != 0 {
		goto L1
	} else {
		goto L5
	}
L5:
	;
	v31 = *(*int32)(unsafe.Add(mBase, _c_F_sentinelProcessHelloMessage[0]))
	v32 = F_dictFetchValue(m, v31, v28)
	mBase = m.M
	v33 = m.ExcPending
	if v33 != 0 {
		goto L1
	} else {
		goto L6
	}
L6:
	;
	F_sdsfree(m, v28)
	mBase = m.M
	v35 = m.ExcPending
	if v35 != 0 {
		goto L1
	} else {
		goto L7
	}
L7:
	;
	if v32 == int32(0) {
		goto L3
	} else {
		goto L8
	}
L8:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v22)+4))
	v42 = v38
	goto L10
L9:
	;
	v88 = *(*int32)(unsafe.Add(mBase, uint32(v22)+24))
	v92 = v88
	goto L25
L10:
	;
	v47 = v42 + int32(1)
	v48 = int32(*(*int8)(unsafe.Add(mBase, uint32(v42))))
	v49 = F___isspace_1(m, v48)
	mBase = m.M
	if v49 != 0 {
		v42 = v47
		goto L10
	} else {
		goto L12
	}
L11:
	;
	v50 = int32(1)
	switch v48&int32(255) + int32(-43) {
	case 0:
		v56 = v50
		goto L14
	default:
		v58 = v42
		v59 = v48
		v60 = v50
		goto L13
	case 2:
		goto L15
	}
L12:
	;
	goto L11
L13:
	;
	v63 = v59 + int32(-48)
	if base.Ui32(int32(9)) < base.Ui32(v63) {
		v81 = int32(0)
		goto L16
	} else {
		goto L17
	}
L14:
	;
	v57 = int32(*(*int8)(unsafe.Add(mBase, uint32(v47))))
	v58 = v47
	v59 = v57
	v60 = v56
	goto L13
L15:
	;
	v56 = int32(0)
	goto L14
L16:
	;
	if v60 != 0 {
		goto L21
	} else {
		goto L22
	}
L17:
	;
	v67 = int32(0)
	v68 = v58
	v69 = v63
	goto L18
L18:
	;
	v71 = int32(10)
	v73 = v67*v71 - v69
	v74 = int32(*(*int8)(unsafe.Add(mBase, uint32(v68)+1)))
	v78 = v74 + int32(-48)
	if base.Ui32(v78) < base.Ui32(v71) {
		v67 = v73
		v68 = v68 + int32(1)
		v69 = v78
		goto L18
	} else {
		goto L20
	}
L19:
	;
	v81 = v73
	goto L16
L20:
	;
	goto L19
L21:
	;
	v87 = int32(0) - v81
	goto L23
L22:
	;
	v87 = v81
	goto L23
L23:
	;
	goto L9
L24:
	;
	v138 = *(*int32)(unsafe.Add(mBase, uint32(v32)+144))
	v139 = *(*int32)(unsafe.Add(mBase, uint32(v22)))
	v140 = *(*int32)(unsafe.Add(mBase, uint32(v22)+8))
	v141 = F_getSentinelValkeyInstanceByAddrAndRunID(m, v138, v139, v87, v140)
	mBase = m.M
	v142 = m.ExcPending
	if v142 != 0 {
		goto L1
	} else {
		goto L39
	}
L25:
	;
	v97 = v92 + int32(1)
	v98 = int32(*(*int8)(unsafe.Add(mBase, uint32(v92))))
	v99 = F___isspace_1(m, v98)
	mBase = m.M
	if v99 != 0 {
		v92 = v97
		goto L25
	} else {
		goto L27
	}
L26:
	;
	v100 = int32(1)
	switch v98&int32(255) + int32(-43) {
	case 0:
		v106 = v100
		goto L29
	default:
		v108 = v92
		v109 = v98
		v110 = v100
		goto L28
	case 2:
		goto L30
	}
L27:
	;
	goto L26
L28:
	;
	v113 = v109 + int32(-48)
	if base.Ui32(int32(9)) < base.Ui32(v113) {
		v131 = int32(0)
		goto L31
	} else {
		goto L32
	}
L29:
	;
	v107 = int32(*(*int8)(unsafe.Add(mBase, uint32(v97))))
	v108 = v97
	v109 = v107
	v110 = v106
	goto L28
L30:
	;
	v106 = int32(0)
	goto L29
L31:
	;
	if v110 != 0 {
		goto L36
	} else {
		goto L37
	}
L32:
	;
	v117 = int32(0)
	v118 = v108
	v119 = v113
	goto L33
L33:
	;
	v121 = int32(10)
	v123 = v117*v121 - v119
	v124 = int32(*(*int8)(unsafe.Add(mBase, uint32(v118)+1)))
	v128 = v124 + int32(-48)
	if base.Ui32(v128) < base.Ui32(v121) {
		v117 = v123
		v118 = v118 + int32(1)
		v119 = v128
		goto L33
	} else {
		goto L35
	}
L34:
	;
	v131 = v123
	goto L31
L35:
	;
	goto L34
L36:
	;
	v137 = int32(0) - v131
	goto L38
L37:
	;
	v137 = v131
	goto L38
L38:
	;
	goto L24
L39:
	;
	v143 = *(*int32)(unsafe.Add(mBase, uint32(v22)+12))
	v147 = F_strtox_2(m, v143, int32(0), int32(10), int64(-1))
	mBase = m.M
	goto L40
L40:
	;
	v148 = *(*int32)(unsafe.Add(mBase, uint32(v22)+28))
	v152 = F_strtox_2(m, v148, int32(0), int32(10), int64(-1))
	mBase = m.M
	goto L41
L41:
	;
	if v141 != 0 {
		v397 = v141
		goto L42
	} else {
		goto L43
	}
L42:
	;
	v408 = *(*int64)(unsafe.Add(mBase, _c_F_sentinelProcessHelloMessage[1]))
	if base.Ui64(v147) <= base.Ui64(v408) {
		goto L109
	} else {
		goto L110
	}
L43:
	;
	v153 = *(*int32)(unsafe.Add(mBase, uint32(v22)+8))
	v154 = F_removeMatchingSentinelFromPrimary(m, v32, v153)
	mBase = m.M
	v155 = m.ExcPending
	if v155 != 0 {
		goto L1
	} else {
		goto L46
	}
L44:
	;
	v332 = *(*int32)(unsafe.Add(mBase, uint32(v22)+8))
	v334 = *(*int32)(unsafe.Add(mBase, uint32(v22)))
	v335 = *(*int32)(unsafe.Add(mBase, uint32(v32)+152))
	v336 = F_createSentinelValkeyInstance(m, v332, int32(4), v334, v87, v335, v32)
	mBase = m.M
	v337 = m.ExcPending
	if v337 != 0 {
		goto L1
	} else {
		goto L89
	}
L45:
	;
	v170 = *(*int32)(unsafe.Add(mBase, uint32(v32)+144))
	v171 = *(*int32)(unsafe.Add(mBase, uint32(v22)))
	v173 = F_getSentinelValkeyInstanceByAddrAndRunID(m, v170, v171, v87, int32(0))
	mBase = m.M
	v174 = m.ExcPending
	if v174 != 0 {
		goto L1
	} else {
		goto L49
	}
L46:
	;
	if v154 == int32(0) {
		goto L45
	} else {
		goto L47
	}
L47:
	;
	v158 = *(*int32)(unsafe.Add(mBase, uint32(v22)))
	v159 = *(*int32)(unsafe.Add(mBase, uint32(v22)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v16)+120)) = v159
	*(*int32)(unsafe.Add(mBase, uint32(v16)+116)) = v87
	*(*int32)(unsafe.Add(mBase, uint32(v16)+112)) = v158
	F_sentinelEvent(m, int32(2), int32(_a_F_sentinelProcessHelloMessage_1), v32, int32(_a_F_sentinelProcessHelloMessage_2), v16+int32(112))
	mBase = m.M
	v169 = m.ExcPending
	if v169 != 0 {
		goto L1
	} else {
		goto L48
	}
L48:
	;
	goto L44
L49:
	;
	if v173 == int32(0) {
		goto L44
	} else {
		goto L50
	}
L50:
	;
	F_sentinelEvent(m, int32(2), int32(_a_F_sentinelProcessHelloMessage_3), v173, int32(_a_F_sentinelProcessHelloMessage_4), int32(0))
	mBase = m.M
	v182 = m.ExcPending
	if v182 != 0 {
		goto L1
	} else {
		goto L51
	}
L51:
	;
	v183 = *(*int32)(unsafe.Add(mBase, uint32(v173)+8))
	v184 = F_sdsnew(m, v183)
	mBase = m.M
	v185 = m.ExcPending
	if v185 != 0 {
		goto L1
	} else {
		goto L52
	}
L52:
	;
	v187 = *(*int32)(unsafe.Add(mBase, _c_F_sentinelProcessHelloMessage[0]))
	v188 = F_dictGetIterator(m, v187)
	mBase = m.M
	v189 = m.ExcPending
	if v189 != 0 {
		goto L1
	} else {
		goto L53
	}
L53:
	;
	goto L55
L54:
	;
	F_dictReleaseIterator(m, v188)
	mBase = m.M
	v316 = m.ExcPending
	if v316 != 0 {
		goto L1
	} else {
		goto L86
	}
L55:
	;
	v210 = v188 + int32(20)
	v211 = *(*int32)(unsafe.Add(mBase, uint32(v188)+16))
	if v211 != 0 {
		goto L59
	} else {
		goto L60
	}
L57:
	;
	if v306 == int32(0) {
		goto L54
	} else {
		goto L83
	}
L58:
	;
	v217 = v210
	v218 = v214
	goto L61
L59:
	;
	v214 = int32(1)
	goto L58
L60:
	;
	v214 = int32(0)
	goto L58
L61:
	;
	switch v218 {
	case 0:
		goto L66
	default:
		goto L65
	}
L63:
	;
	v218 = int32(0)
	goto L61
L64:
	;
	goto L57
L65:
	;
	v298 = *(*int32)(unsafe.Add(mBase, uint32(v217)))
	*(*int32)(unsafe.Add(mBase, uint32(v188)+16)) = v298
	if v298 == int32(0) {
		goto L63
	} else {
		goto L82
	}
L66:
	;
	v222 = *(*int32)(unsafe.Add(mBase, uint32(v188)+4))
	if v222 != int32(-1) {
		v261 = v222
		goto L67
	} else {
		goto L68
	}
L67:
	;
	v262 = int32(1)
	v263 = v261 + v262
	*(*int32)(unsafe.Add(mBase, uint32(v188)+4)) = v263
	v265 = int32(0)
	v268 = *(*int32)(unsafe.Add(mBase, uint32(v188)))
	v269 = *(*int32)(unsafe.Add(mBase, uint32(v188)+8))
	v273 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v268+v269+int32(26)))))
	if v273 == int32(255) {
		goto L76
	} else {
		goto L77
	}
L68:
	;
	v226 = *(*int32)(unsafe.Add(mBase, uint32(v188)+8))
	if v226 != 0 {
		v261 = int32(-1)
		goto L67
	} else {
		goto L69
	}
L69:
	;
	v227 = *(*int32)(unsafe.Add(mBase, uint32(v188)))
	v228 = *(*int32)(unsafe.Add(mBase, uint32(v188)+12))
	if v228 == int32(0) {
		goto L71
	} else {
		goto L72
	}
L70:
	;
	v255 = *(*int32)(unsafe.Add(mBase, uint32(v254)+20))
	if v255 != int32(-1) {
		goto L73
	} else {
		goto L74
	}
L71:
	;
	v235 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v227)+16)))
	v236 = int64(*(*int8)(unsafe.Add(mBase, uint32(v227)+27)))
	v237 = int64(*(*int32)(unsafe.Add(mBase, uint32(v227)+8)))
	v238 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v227)+12)))
	v239 = int64(*(*int8)(unsafe.Add(mBase, uint32(v227)+26)))
	v240 = int64(*(*int32)(unsafe.Add(mBase, uint32(v227)+4)))
	v241 = F_wangHash64(m, v240)
	mBase = m.M
	v243 = F_wangHash64(m, v239+v241)
	mBase = m.M
	v245 = F_wangHash64(m, v238+v243)
	mBase = m.M
	v247 = F_wangHash64(m, v237+v245)
	mBase = m.M
	v249 = F_wangHash64(m, v236+v247)
	mBase = m.M
	v251 = F_wangHash64(m, v235+v249)
	mBase = m.M
	*(*int64)(unsafe.Add(mBase, uint32(v188)+24)) = v251
	v253 = *(*int32)(unsafe.Add(mBase, uint32(v188)))
	v254 = v253
	goto L70
L72:
	;
	v231 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v227)+24)))
	v233 = v231 + int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v227)+24)) = uint16(v233)
	v254 = v227
	goto L70
L73:
	;
	v261 = v255 + int32(-1)
	goto L67
L74:
	;
	v258 = *(*int32)(unsafe.Add(mBase, uint32(v188)+4))
	v261 = v258
	goto L67
L75:
	;
	v288 = int32(2)
	v293 = *(*int32)(unsafe.Add(mBase, uint32(v268+v286<<(uint(v288)%32)+int32(4))))
	v217 = v293 + v287<<(uint(v288)%32)
	v218 = int32(1)
	goto L61
L76:
	;
	v277 = v265
	goto L78
L77:
	;
	v277 = v262 << (uint(v273) % 32)
	goto L78
L78:
	;
	if v263 < v277 {
		v286 = v269
		v287 = v263
		goto L75
	} else {
		goto L79
	}
L79:
	;
	if v269 != 0 {
		v306 = v265
		goto L64
	} else {
		goto L80
	}
L80:
	;
	v279 = *(*int32)(unsafe.Add(mBase, uint32(v268)+20))
	if v279 == int32(-1) {
		v306 = v265
		goto L64
	} else {
		goto L81
	}
L81:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v188)+4)) = int64(4294967296)
	v286 = int32(1)
	v287 = int32(0)
	goto L75
L82:
	;
	v302 = *(*int32)(unsafe.Add(mBase, uint32(v298)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v210))) = v302
	v306 = v298
	goto L64
L83:
	;
	v312 = *(*int32)(unsafe.Add(mBase, uint32(v306)+8))
	goto L84
L84:
	;
	v313 = F_removeMatchingSentinelFromPrimary(m, v312, v184)
	mBase = m.M
	v314 = m.ExcPending
	if v314 != 0 {
		goto L1
	} else {
		goto L85
	}
L85:
	;
	goto L55
L86:
	;
	F_sdsfree(m, v184)
	mBase = m.M
	v318 = m.ExcPending
	if v318 != 0 {
		goto L1
	} else {
		goto L87
	}
L87:
	;
	goto L44
L88:
	;
	if v154 != 0 {
		goto L92
	} else {
		goto L93
	}
L89:
	;
	if v336 != 0 {
		goto L88
	} else {
		goto L90
	}
L90:
	;
	v397 = int32(0)
	goto L42
L91:
	;
	v359 = int32(_a_F_sentinelProcessHelloMessage_5)
	v360 = *(*int32)(unsafe.Add(mBase, _c_F_sentinelProcessHelloMessage[2]))
	*(*int32)(unsafe.Add(mBase, _c_F_sentinelProcessHelloMessage[2])) = int32(10)
	v365 = *(*int32)(unsafe.Add(mBase, _c_F_sentinelProcessHelloMessage[3]))
	v367 = F_rewriteConfig(m, v365, int32(0))
	mBase = m.M
	v368 = m.ExcPending
	if v368 != 0 {
		goto L1
	} else {
		goto L100
	}
L92:
	;
	v351 = *(*int32)(unsafe.Add(mBase, uint32(v22)+8))
	v352 = F_sdsnew(m, v351)
	mBase = m.M
	v353 = m.ExcPending
	if v353 != 0 {
		goto L1
	} else {
		goto L97
	}
L93:
	;
	F_sentinelEvent(m, int32(2), int32(_a_F_sentinelProcessHelloMessage_6), v336, int32(_a_F_sentinelProcessHelloMessage_4), int32(0))
	mBase = m.M
	v344 = m.ExcPending
	if v344 != 0 {
		goto L1
	} else {
		goto L94
	}
L94:
	;
	v345 = *(*int32)(unsafe.Add(mBase, uint32(v22)+8))
	v346 = F_sdsnew(m, v345)
	mBase = m.M
	v347 = m.ExcPending
	if v347 != 0 {
		goto L1
	} else {
		goto L95
	}
L95:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v336)+8)) = v346
	v349 = F_sentinelTryConnectionSharing(m, v336)
	mBase = m.M
	v350 = m.ExcPending
	if v350 != 0 {
		goto L1
	} else {
		goto L96
	}
L96:
	;
	goto L91
L97:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v336)+8)) = v352
	v355 = F_sentinelTryConnectionSharing(m, v336)
	mBase = m.M
	v356 = m.ExcPending
	if v356 != 0 {
		goto L1
	} else {
		goto L98
	}
L98:
	;
	v357 = F_sentinelUpdateSentinelAddressInAllPrimaries(m, v336)
	mBase = m.M
	v358 = m.ExcPending
	if v358 != 0 {
		goto L1
	} else {
		goto L99
	}
L99:
	;
	goto L91
L100:
	;
	v369 = int32(_a_F_sentinelProcessHelloMessage_5)
	*(*int32)(unsafe.Add(mBase, _c_F_sentinelProcessHelloMessage[2])) = v360
	v372 = *(*int32)(unsafe.Add(mBase, _c_F_sentinelProcessHelloMessage[4]))
	if v367 != int32(-1) {
		goto L101
	} else {
		goto L102
	}
L101:
	;
	if int32(2) < v372 {
		v397 = v336
		goto L42
	} else {
		goto L107
	}
L102:
	;
	if int32(3) < v372 {
		v397 = v336
		goto L42
	} else {
		goto L103
	}
L103:
	;
	goto L104
L104:
	;
	v378 = *(*int32)(unsafe.Add(mBase, _c_F_sentinelProcessHelloMessage[5]))
	v379 = F___strerror_l(m, v378, v378)
	mBase = m.M
	goto L105
L105:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+96)) = v379
	F__serverLog(m, int32(3), int32(_a_F_sentinelProcessHelloMessage_7), v16+int32(96))
	mBase = m.M
	v386 = m.ExcPending
	if v386 != 0 {
		goto L1
	} else {
		goto L106
	}
L106:
	;
	v397 = v336
	goto L42
L107:
	;
	F__serverLog(m, int32(2), int32(_a_F_sentinelProcessHelloMessage_8), int32(0))
	mBase = m.M
	v393 = m.ExcPending
	if v393 != 0 {
		goto L1
	} else {
		goto L108
	}
L108:
	;
	v397 = v336
	goto L42
L109:
	;
	if v397 == int32(0) {
		goto L3
	} else {
		goto L122
	}
L110:
	;
	v410 = int32(0)
	*(*int64)(unsafe.Add(mBase, _c_F_sentinelProcessHelloMessage[1])) = v147
	v412 = int32(_a_F_sentinelProcessHelloMessage_5)
	v413 = *(*int32)(unsafe.Add(mBase, _c_F_sentinelProcessHelloMessage[2]))
	*(*int32)(unsafe.Add(mBase, _c_F_sentinelProcessHelloMessage[2])) = int32(10)
	v418 = *(*int32)(unsafe.Add(mBase, _c_F_sentinelProcessHelloMessage[3]))
	v420 = F_rewriteConfig(m, v418, v410)
	mBase = m.M
	v421 = m.ExcPending
	if v421 != 0 {
		goto L1
	} else {
		goto L111
	}
L111:
	;
	v422 = int32(_a_F_sentinelProcessHelloMessage_5)
	*(*int32)(unsafe.Add(mBase, _c_F_sentinelProcessHelloMessage[2])) = v413
	v425 = *(*int32)(unsafe.Add(mBase, _c_F_sentinelProcessHelloMessage[4]))
	if v420 != int32(-1) {
		goto L113
	} else {
		goto L114
	}
L112:
	;
	v448 = *(*int64)(unsafe.Add(mBase, _c_F_sentinelProcessHelloMessage[1]))
	*(*int64)(unsafe.Add(mBase, uint32(v16)+64)) = v448
	F_sentinelEvent(m, int32(3), int32(_a_F_sentinelProcessHelloMessage_9), v32, int32(_a_F_sentinelProcessHelloMessage_10), v16+int32(64))
	mBase = m.M
	v456 = m.ExcPending
	if v456 != 0 {
		goto L1
	} else {
		goto L121
	}
L113:
	;
	if int32(2) < v425 {
		goto L112
	} else {
		goto L119
	}
L114:
	;
	if int32(3) < v425 {
		goto L112
	} else {
		goto L115
	}
L115:
	;
	goto L116
L116:
	;
	v431 = *(*int32)(unsafe.Add(mBase, _c_F_sentinelProcessHelloMessage[5]))
	v432 = F___strerror_l(m, v431, v431)
	mBase = m.M
	goto L117
L117:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+80)) = v432
	F__serverLog(m, int32(3), int32(_a_F_sentinelProcessHelloMessage_7), v16+int32(80))
	mBase = m.M
	v439 = m.ExcPending
	if v439 != 0 {
		goto L1
	} else {
		goto L118
	}
L118:
	;
	goto L112
L119:
	;
	F__serverLog(m, int32(2), int32(_a_F_sentinelProcessHelloMessage_8), int32(0))
	mBase = m.M
	v446 = m.ExcPending
	if v446 != 0 {
		goto L1
	} else {
		goto L120
	}
L120:
	;
	goto L112
L121:
	;
	goto L109
L122:
	;
	v461 = *(*int64)(unsafe.Add(mBase, uint32(v32)+16))
	if base.Ui64(v152) <= base.Ui64(v461) {
		goto L123
	} else {
		goto L124
	}
L123:
	;
	v736 = F_mstime(m)
	mBase = m.M
	*(*int64)(unsafe.Add(mBase, uint32(v397)+40)) = v736
	goto L3
L124:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v32)+16)) = v152
	v464 = *(*int32)(unsafe.Add(mBase, uint32(v32)+24))
	v465 = *(*int32)(unsafe.Add(mBase, uint32(v464)+8))
	if v137 != v465 {
		goto L125
	} else {
		goto L126
	}
L125:
	;
	F_sentinelEvent(m, int32(3), int32(_a_F_sentinelProcessHelloMessage_11), v397, int32(_a_F_sentinelProcessHelloMessage_4), int32(0))
	mBase = m.M
	v540 = m.ExcPending
	if v540 != 0 {
		goto L1
	} else {
		goto L147
	}
L126:
	;
	v467 = int32(0)
	v468 = *(*int32)(unsafe.Add(mBase, uint32(v22)+20))
	v473 = *(*int32)(unsafe.Add(mBase, _c_F_sentinelProcessHelloMessage[6]))
	v476 = F_anetResolve(m, v467, v468, v16+int32(128), int32(46), base.B2i32(v473 == v467))
	mBase = m.M
	v477 = m.ExcPending
	if v477 != 0 {
		goto L1
	} else {
		goto L127
	}
L127:
	;
	v478 = int32(0)
	v479 = *(*int32)(unsafe.Add(mBase, _c_F_sentinelProcessHelloMessage[6]))
	v486 = base.B2i32(v476 == int32(-1))
	if v476 == int32(-1) {
		goto L128
	} else {
		goto L129
	}
L128:
	;
	v487 = base.B2i32(v479 == v478) << (uint(int32(2)) % 32)
	goto L130
L129:
	;
	v487 = int32(4)
	goto L130
L130:
	;
	v489 = *(*int32)(unsafe.Add(mBase, uint32(v464+v487)))
	if v476 == int32(-1) {
		goto L131
	} else {
		goto L132
	}
L131:
	;
	v492 = v468
	goto L133
L132:
	;
	v492 = v16 + int32(128)
	goto L133
L133:
	;
	v495 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v489))))
	if v495 != 0 {
		goto L136
	} else {
		goto L137
	}
L134:
	;
	if v527-v529 == int32(0) {
		goto L123
	} else {
		goto L146
	}
L135:
	;
	v527 = F_tolower(m, v523)
	mBase = m.M
	v528 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v524))))
	v529 = F_tolower(m, v528)
	mBase = m.M
	goto L134
L136:
	;
	v497 = v489
	v498 = v492
	v499 = v495
	goto L139
L137:
	;
	v523 = int32(0)
	v524 = v492
	goto L135
L138:
	;
	v523 = v520 & int32(255)
	v524 = v519
	goto L135
L139:
	;
	v501 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v498))))
	if v501 == int32(0) {
		v519 = v498
		v520 = v499
		goto L138
	} else {
		goto L141
	}
L140:
	;
	v519 = v513
	v520 = int32(0)
	goto L138
L141:
	;
	v505 = v499 & int32(255)
	if v505 == v501 {
		goto L142
	} else {
		goto L143
	}
L142:
	;
	v512 = int32(1)
	v513 = v498 + v512
	v514 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v497)+1)))
	if v514 != 0 {
		v497 = v497 + v512
		v498 = v513
		v499 = v514
		goto L139
	} else {
		goto L145
	}
L143:
	;
	v507 = F_tolower(m, v505)
	mBase = m.M
	v508 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v498))))
	v509 = F_tolower(m, v508)
	mBase = m.M
	if v507 == v509 {
		goto L142
	} else {
		goto L144
	}
L144:
	;
	v511 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v497))))
	v519 = v498
	v520 = v511
	goto L138
L145:
	;
	goto L140
L146:
	;
	goto L125
L147:
	;
	v541 = *(*int32)(unsafe.Add(mBase, uint32(v32)+24))
	v542 = int32(0)
	v543 = *(*int32)(unsafe.Add(mBase, _c_F_sentinelProcessHelloMessage[7]))
	v549 = *(*int32)(unsafe.Add(mBase, uint32(v541+base.B2i32(v543 == v542)<<(uint(int32(2))%32))))
	v550 = *(*int32)(unsafe.Add(mBase, uint32(v22)+20))
	v551 = *(*int32)(unsafe.Add(mBase, uint32(v541)+8))
	v552 = *(*int32)(unsafe.Add(mBase, uint32(v32)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v16+int32(48)))) = v137
	*(*int32)(unsafe.Add(mBase, uint32(v16)+32)) = v552
	*(*int32)(unsafe.Add(mBase, uint32(v16)+36)) = v549
	*(*int32)(unsafe.Add(mBase, uint32(v16)+40)) = v551
	*(*int32)(unsafe.Add(mBase, uint32(v16)+44)) = v550
	F_sentinelEvent(m, int32(3), int32(_a_F_sentinelProcessHelloMessage_12), v32, int32(_a_F_sentinelProcessHelloMessage_13), v16+int32(32))
	mBase = m.M
	v568 = m.ExcPending
	if v568 != 0 {
		goto L1
	} else {
		goto L148
	}
L148:
	;
	v569 = *(*int32)(unsafe.Add(mBase, uint32(v32)+24))
	v571 = F_valkey_malloc(m, int32(12))
	mBase = m.M
	v572 = m.ExcPending
	if v572 != 0 {
		goto L1
	} else {
		goto L149
	}
L149:
	;
	v573 = *(*int32)(unsafe.Add(mBase, uint32(v569)))
	v574 = F_sdsnew(m, v573)
	mBase = m.M
	v575 = m.ExcPending
	if v575 != 0 {
		goto L1
	} else {
		goto L150
	}
L150:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v571))) = v574
	v577 = *(*int32)(unsafe.Add(mBase, uint32(v569)+4))
	v578 = F_sdsnew(m, v577)
	mBase = m.M
	v579 = m.ExcPending
	if v579 != 0 {
		goto L1
	} else {
		goto L151
	}
L151:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v571)+4)) = v578
	v581 = *(*int32)(unsafe.Add(mBase, uint32(v569)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v571)+8)) = v581
	v583 = *(*int32)(unsafe.Add(mBase, uint32(v22)+20))
	v584 = F_sentinelResetPrimaryAndChangeAddress(m, v32, v583, v137)
	mBase = m.M
	v585 = m.ExcPending
	if v585 != 0 {
		goto L1
	} else {
		goto L152
	}
L152:
	;
	v586 = *(*int32)(unsafe.Add(mBase, uint32(v32)+288))
	if v586 == int32(0) {
		goto L153
	} else {
		goto L154
	}
L153:
	;
	F_sdsfree(m, v574)
	mBase = m.M
	v723 = m.ExcPending
	if v723 != 0 {
		goto L1
	} else {
		goto L174
	}
L154:
	;
	v589 = *(*int32)(unsafe.Add(mBase, uint32(v32)+24))
	v591 = v16 + int32(128)
	v593 = base.I64_extend_i32_s(v581)
	if v593 <= int64(-1) {
		goto L159
	} else {
		goto L160
	}
L155:
	;
	v636 = v16 + int32(176)
	v638 = int64(*(*int32)(unsafe.Add(mBase, uint32(v589)+8)))
	if v638 <= int64(-1) {
		goto L168
	} else {
		goto L169
	}
L156:
	;
	goto L155
L158:
	;
	v615 = F_ull2string(m, v611, v612, v613)
	mBase = m.M
	if v615 == int32(0) {
		goto L156
	} else {
		goto L162
	}
L159:
	;
	goto L161
L160:
	;
	v611 = v591
	v612 = int32(32)
	v613 = v593
	goto L158
L161:
	;
	v602 = int32(45)
	*(*uint8)(unsafe.Add(mBase, uint32(v591))) = uint8(v602)
	v611 = v16 + int32(129)
	v612 = int32(31)
	v613 = int64(0) - v593
	goto L158
L162:
	;
	goto L155
L164:
	;
	v680 = int32(0)
	v681 = *(*int32)(unsafe.Add(mBase, _c_F_sentinelProcessHelloMessage[7]))
	v685 = base.B2i32(v681 == v680) << (uint(int32(2)) % 32)
	v687 = *(*int32)(unsafe.Add(mBase, uint32(v571+v685)))
	v688 = *(*int32)(unsafe.Add(mBase, uint32(v32)+288))
	v689 = *(*int32)(unsafe.Add(mBase, uint32(v32)+4))
	v693 = *(*int32)(unsafe.Add(mBase, uint32(v589+v685)))
	*(*int32)(unsafe.Add(mBase, uint32(v16+int32(20)))) = v693
	*(*int32)(unsafe.Add(mBase, uint32(v16+int32(28)))) = v680
	*(*int32)(unsafe.Add(mBase, uint32(v16))) = v689
	*(*int32)(unsafe.Add(mBase, uint32(v16)+4)) = int32(_a_F_sentinelProcessHelloMessage_14)
	*(*int32)(unsafe.Add(mBase, uint32(v16)+8)) = int32(_a_F_sentinelProcessHelloMessage_15)
	*(*int32)(unsafe.Add(mBase, uint32(v16)+12)) = v687
	*(*int32)(unsafe.Add(mBase, uint32(v16+int32(16)))) = v16 + int32(128)
	*(*int32)(unsafe.Add(mBase, uint32(v16+int32(24)))) = v16 + int32(176)
	F_sentinelScheduleScriptExecution(m, v688, v16)
	mBase = m.M
	v716 = m.ExcPending
	if v716 != 0 {
		goto L1
	} else {
		goto L173
	}
L165:
	;
	goto L164
L167:
	;
	v660 = F_ull2string(m, v656, v657, v658)
	mBase = m.M
	if v660 == int32(0) {
		goto L165
	} else {
		goto L171
	}
L168:
	;
	goto L170
L169:
	;
	v656 = v636
	v657 = int32(32)
	v658 = v638
	goto L167
L170:
	;
	v647 = int32(45)
	*(*uint8)(unsafe.Add(mBase, uint32(v636))) = uint8(v647)
	v656 = v16 + int32(177)
	v657 = int32(31)
	v658 = int64(0) - v638
	goto L167
L171:
	;
	goto L164
L173:
	;
	goto L153
L174:
	;
	F_sdsfree(m, v578)
	mBase = m.M
	v725 = m.ExcPending
	if v725 != 0 {
		goto L1
	} else {
		goto L175
	}
L175:
	;
	F_valkey_free(m, v571)
	mBase = m.M
	v727 = m.ExcPending
	if v727 != 0 {
		goto L1
	} else {
		goto L176
	}
L176:
	;
	goto L123
L177:
	;
	m.G0 = v16 + int32(208)
	return
}
func F_sentinelPropagateDownAfterPeriod(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
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
	var v67 int64
	_ = v67
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
	var v78 int64
	_ = v78
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v100 int32
	_ = v100
	var v104 int32
	_ = v104
	var v106 int32
	_ = v106
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v120 int32
	_ = v120
	var v125 int32
	_ = v125
	var v129 int32
	_ = v129
	var v133 int32
	_ = v133
	var v139 int32
	_ = v139
	var v140 int64
	_ = v140
	var v143 int32
	_ = v143
	var v147 int32
	_ = v147
	var v151 int32
	_ = v151
	v2 = int32(0)
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+148))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+144))
	*(*int32)(unsafe.Add(mBase, uint32(v8)+12)) = v2
	*(*int32)(unsafe.Add(mBase, uint32(v8)+8)) = v11
	if v10 == v2 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	m.G0 = v8 + int32(16)
	return
L2:
	;
	v20 = v10
	v22 = v2
	goto L3
L3:
	;
	v23 = F_dictGetIterator(m, v20)
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
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
	goto L8
L7:
	;
	F_dictReleaseIterator(m, v23)
	mBase = m.M
	v143 = m.ExcPending
	if v143 != 0 {
		goto L5
	} else {
		goto L38
	}
L8:
	;
	v37 = v23 + int32(20)
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v23)+16))
	if v38 != 0 {
		goto L12
	} else {
		goto L13
	}
L10:
	;
	if v133 == int32(0) {
		goto L7
	} else {
		goto L36
	}
L11:
	;
	v44 = v37
	v45 = v41
	goto L14
L12:
	;
	v41 = int32(1)
	goto L11
L13:
	;
	v41 = int32(0)
	goto L11
L14:
	;
	switch v45 {
	case 0:
		goto L19
	default:
		goto L18
	}
L16:
	;
	v45 = int32(0)
	goto L14
L17:
	;
	goto L10
L18:
	;
	v125 = *(*int32)(unsafe.Add(mBase, uint32(v44)))
	*(*int32)(unsafe.Add(mBase, uint32(v23)+16)) = v125
	if v125 == int32(0) {
		goto L16
	} else {
		goto L35
	}
L19:
	;
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v23)+4))
	if v49 != int32(-1) {
		v88 = v49
		goto L20
	} else {
		goto L21
	}
L20:
	;
	v89 = int32(1)
	v90 = v88 + v89
	*(*int32)(unsafe.Add(mBase, uint32(v23)+4)) = v90
	v92 = int32(0)
	v95 = *(*int32)(unsafe.Add(mBase, uint32(v23)))
	v96 = *(*int32)(unsafe.Add(mBase, uint32(v23)+8))
	v100 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v95+v96+int32(26)))))
	if v100 == int32(255) {
		goto L29
	} else {
		goto L30
	}
L21:
	;
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v23)+8))
	if v53 != 0 {
		v88 = int32(-1)
		goto L20
	} else {
		goto L22
	}
L22:
	;
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v23)))
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v23)+12))
	if v55 == int32(0) {
		goto L24
	} else {
		goto L25
	}
L23:
	;
	v82 = *(*int32)(unsafe.Add(mBase, uint32(v81)+20))
	if v82 != int32(-1) {
		goto L26
	} else {
		goto L27
	}
L24:
	;
	v62 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v54)+16)))
	v63 = int64(*(*int8)(unsafe.Add(mBase, uint32(v54)+27)))
	v64 = int64(*(*int32)(unsafe.Add(mBase, uint32(v54)+8)))
	v65 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v54)+12)))
	v66 = int64(*(*int8)(unsafe.Add(mBase, uint32(v54)+26)))
	v67 = int64(*(*int32)(unsafe.Add(mBase, uint32(v54)+4)))
	v68 = F_wangHash64(m, v67)
	mBase = m.M
	v70 = F_wangHash64(m, v66+v68)
	mBase = m.M
	v72 = F_wangHash64(m, v65+v70)
	mBase = m.M
	v74 = F_wangHash64(m, v64+v72)
	mBase = m.M
	v76 = F_wangHash64(m, v63+v74)
	mBase = m.M
	v78 = F_wangHash64(m, v62+v76)
	mBase = m.M
	*(*int64)(unsafe.Add(mBase, uint32(v23)+24)) = v78
	v80 = *(*int32)(unsafe.Add(mBase, uint32(v23)))
	v81 = v80
	goto L23
L25:
	;
	v58 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v54)+24)))
	v60 = v58 + int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v54)+24)) = uint16(v60)
	v81 = v54
	goto L23
L26:
	;
	v88 = v82 + int32(-1)
	goto L20
L27:
	;
	v85 = *(*int32)(unsafe.Add(mBase, uint32(v23)+4))
	v88 = v85
	goto L20
L28:
	;
	v115 = int32(2)
	v120 = *(*int32)(unsafe.Add(mBase, uint32(v95+v113<<(uint(v115)%32)+int32(4))))
	v44 = v120 + v114<<(uint(v115)%32)
	v45 = int32(1)
	goto L14
L29:
	;
	v104 = v92
	goto L31
L30:
	;
	v104 = v89 << (uint(v100) % 32)
	goto L31
L31:
	;
	if v90 < v104 {
		v113 = v96
		v114 = v90
		goto L28
	} else {
		goto L32
	}
L32:
	;
	if v96 != 0 {
		v133 = v92
		goto L17
	} else {
		goto L33
	}
L33:
	;
	v106 = *(*int32)(unsafe.Add(mBase, uint32(v95)+20))
	if v106 == int32(-1) {
		v133 = v92
		goto L17
	} else {
		goto L34
	}
L34:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v23)+4)) = int64(4294967296)
	v113 = int32(1)
	v114 = int32(0)
	goto L28
L35:
	;
	v129 = *(*int32)(unsafe.Add(mBase, uint32(v125)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v37))) = v129
	v133 = v125
	goto L17
L36:
	;
	v139 = *(*int32)(unsafe.Add(mBase, uint32(v133)+8))
	goto L37
L37:
	;
	v140 = *(*int64)(unsafe.Add(mBase, uint32(l0)+72))
	*(*int64)(unsafe.Add(mBase, uint32(v139)+72)) = v140
	goto L8
L38:
	;
	v147 = v22 + int32(1)
	v151 = *(*int32)(unsafe.Add(mBase, uint32(v8+int32(4)+v147<<(uint(int32(2))%32))))
	if v151 != 0 {
		v20 = v151
		v22 = v147
		goto L3
	} else {
		goto L39
	}
L39:
	;
	goto L4
}
func F_sentinelPublishReplyCallback(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v13 int32
	_ = v13
	var v16 int64
	_ = v16
	if l1 == int32(0) {
	} else {
		v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+212))
		if v6 == int32(0) {
		} else {
			v9 = *(*int32)(unsafe.Add(mBase, uint32(v6)+8))
			*(*int32)(unsafe.Add(mBase, uint32(v6)+8)) = v9 + int32(-1)
			v13 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
			if v13 == int32(6) {
			} else {
				v16 = F_mstime(m)
				mBase = m.M
				*(*int64)(unsafe.Add(mBase, uint32(l2)+32)) = v16
			}
		}
	}
	return
}
func F_sentinelReceiveHelloMessages(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v9 int64
	_ = v9
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
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
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
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
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
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
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	if l1 == int32(0) {
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
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v9 = F_mstime(m)
	mBase = m.M
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l2)+28))
	*(*int64)(unsafe.Add(mBase, uint32(v10)+40)) = v9
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	switch v12 + int32(-2) {
	case 0, 10:
		goto L4
	default:
		goto L1
	}
L4:
	;
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l1)+36))
	if v15 != int32(3) {
		goto L1
	} else {
		goto L5
	}
L5:
	;
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
	v19 = *(*int32)(unsafe.Add(mBase, uint32(v18)))
	v20 = *(*int32)(unsafe.Add(mBase, uint32(v19)))
	if v20 != int32(1) {
		goto L1
	} else {
		goto L6
	}
L6:
	;
	v23 = *(*int32)(unsafe.Add(mBase, uint32(v18)+4))
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v23)))
	if v24 != int32(1) {
		goto L1
	} else {
		goto L7
	}
L7:
	;
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v18)+8))
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v27)))
	if v28 != int32(1) {
		goto L1
	} else {
		goto L8
	}
L8:
	;
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v19)+28))
	v32 = int32(_a_F_sentinelReceiveHelloMessages_0)
	v35 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_sentinelReceiveHelloMessages[0])))
	v36 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v31))))
	if v36 == int32(0) {
		v59 = v35
		v60 = v36
		goto L10
	} else {
		goto L11
	}
L9:
	;
	if v60-v59&int32(255) != 0 {
		goto L1
	} else {
		goto L17
	}
L10:
	;
	goto L9
L11:
	;
	if v36 != v35&int32(255) {
		v59 = v35
		v60 = v36
		goto L10
	} else {
		goto L12
	}
L12:
	;
	v42 = v31
	v43 = v32
	goto L13
L13:
	;
	v46 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v43)+1)))
	v47 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v42)+1)))
	if v47 == int32(0) {
		v59 = v46
		v60 = v47
		goto L10
	} else {
		goto L15
	}
L14:
	;
	v59 = v46
	v60 = v47
	goto L10
L15:
	;
	v50 = int32(1)
	if v47 == v46&int32(255) {
		v42 = v42 + v50
		v43 = v43 + v50
		goto L13
	} else {
		goto L16
	}
L16:
	;
	goto L14
L17:
	;
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v27)+28))
	v65 = int32(_a_F_sentinelReceiveHelloMessages_1)
	v68 = int32(*(*int8)(unsafe.Add(mBase, _c_F_sentinelReceiveHelloMessages[1])))
	if v68 != 0 {
		goto L19
	} else {
		goto L20
	}
L18:
	;
	if v93 != 0 {
		goto L1
	} else {
		goto L34
	}
L19:
	;
	v69 = int32(0)
	v70 = F_strchr(m, v64, v68)
	mBase = m.M
	if v70 == v69 {
		v90 = v69
		goto L21
	} else {
		goto L22
	}
L20:
	;
	v93 = v64
	goto L18
L21:
	;
	v93 = v90
	goto L18
L22:
	;
	v73 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_sentinelReceiveHelloMessages[2])))
	if v73 != 0 {
		goto L23
	} else {
		goto L24
	}
L23:
	;
	v74 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v70)+1)))
	if v74 == int32(0) {
		v90 = v69
		goto L21
	} else {
		goto L25
	}
L24:
	;
	v93 = v70
	goto L18
L25:
	;
	v77 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_sentinelReceiveHelloMessages[3])))
	if v77 != 0 {
		goto L26
	} else {
		goto L27
	}
L26:
	;
	v79 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v70)+2)))
	if v79 == int32(0) {
		v90 = v69
		goto L21
	} else {
		goto L28
	}
L27:
	;
	v78 = F_twobyte_strstr(m, v70, v65)
	mBase = m.M
	v93 = v78
	goto L18
L28:
	;
	v82 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_sentinelReceiveHelloMessages[4])))
	if v82 != 0 {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	v84 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v70)+3)))
	if v84 == int32(0) {
		v90 = v69
		goto L21
	} else {
		goto L31
	}
L30:
	;
	v83 = F_threebyte_strstr(m, v70, v65)
	mBase = m.M
	v93 = v83
	goto L18
L31:
	;
	v87 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_sentinelReceiveHelloMessages[5])))
	if v87 != 0 {
		goto L32
	} else {
		goto L33
	}
L32:
	;
	v89 = F_twoway_strstr(m, v70, v65)
	mBase = m.M
	v90 = v89
	goto L21
L33:
	;
	v88 = F_fourbyte_strstr(m, v70, v65)
	mBase = m.M
	v93 = v88
	goto L18
L34:
	;
	v94 = *(*int32)(unsafe.Add(mBase, uint32(v27)+24))
	F_sentinelProcessHelloMessage(m, v64, v94)
	mBase = m.M
	v96 = m.ExcPending
	if v96 != 0 {
		goto L35
	} else {
		goto L36
	}
L35:
	;
	return
L36:
	;
	goto L1
}
func F_sentinelReceiveIsPrimaryDownReply(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v39 int64
	_ = v39
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int64
	_ = v46
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v64 int64
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v67 int64
	_ = v67
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v93 int64
	_ = v93
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	if l1 == int32(0) {
		m.G0 = v9 + int32(16)
		return
	} else {
		v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+212))
		if v13 == int32(0) {
			m.G0 = v9 + int32(16)
			return
		} else {
			v16 = *(*int32)(unsafe.Add(mBase, uint32(v13)+8))
			*(*int32)(unsafe.Add(mBase, uint32(v13)+8)) = v16 + int32(-1)
			v20 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
			if v20 != int32(2) {
				m.G0 = v9 + int32(16)
				return
			} else {
				v23 = *(*int32)(unsafe.Add(mBase, uint32(l1)+36))
				if v23 != int32(3) {
					m.G0 = v9 + int32(16)
					return
				} else {
					v26 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
					v27 = *(*int32)(unsafe.Add(mBase, uint32(v26)))
					v28 = *(*int32)(unsafe.Add(mBase, uint32(v27)))
					if v28 != int32(3) {
						m.G0 = v9 + int32(16)
						return
					} else {
						v31 = *(*int32)(unsafe.Add(mBase, uint32(v26)+4))
						v32 = *(*int32)(unsafe.Add(mBase, uint32(v31)))
						if v32 != int32(1) {
							m.G0 = v9 + int32(16)
							return
						} else {
							v35 = *(*int32)(unsafe.Add(mBase, uint32(v26)+8))
							v36 = *(*int32)(unsafe.Add(mBase, uint32(v35)))
							if v36 != int32(3) {
								m.G0 = v9 + int32(16)
								return
							} else {
								v39 = F_mstime(m)
								mBase = m.M
								*(*int64)(unsafe.Add(mBase, uint32(l2)+48)) = v39
								v41 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
								v44 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
								v45 = *(*int32)(unsafe.Add(mBase, uint32(v44)))
								v46 = *(*int64)(unsafe.Add(mBase, uint32(v45)+8))
								*(*int32)(unsafe.Add(mBase, uint32(l2))) = v41&int32(-33) | base.B2i32(v46 == int64(1))<<(uint(int32(5))%32)
								v53 = *(*int32)(unsafe.Add(mBase, uint32(v44)+4))
								v54 = *(*int32)(unsafe.Add(mBase, uint32(v53)+28))
								v55 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v54))))
								if v55 != int32(42) {
									v61 = *(*int32)(unsafe.Add(mBase, uint32(l2)+216))
									F_sdsfree(m, v61)
									mBase = m.M
									v63 = m.ExcPending
									if v63 != 0 {
										return
									} else {
										v64 = *(*int64)(unsafe.Add(mBase, uint32(l2)+224))
										v65 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
										v66 = *(*int32)(unsafe.Add(mBase, uint32(v65)+8))
										v67 = *(*int64)(unsafe.Add(mBase, uint32(v66)+8))
										if v64 == v67 {
											v84 = v65
											v86 = *(*int32)(unsafe.Add(mBase, uint32(v84)+4))
											v87 = *(*int32)(unsafe.Add(mBase, uint32(v86)+28))
											v88 = F_sdsnew(m, v87)
											mBase = m.M
											v89 = m.ExcPending
											if v89 != 0 {
												return
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(l2)+216)) = v88
												v91 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
												v92 = *(*int32)(unsafe.Add(mBase, uint32(v91)+8))
												v93 = *(*int64)(unsafe.Add(mBase, uint32(v92)+8))
												*(*int64)(unsafe.Add(mBase, uint32(l2)+224)) = v93
												m.G0 = v9 + int32(16)
												return
											}
										} else {
											v70 = *(*int32)(unsafe.Add(mBase, _c_F_sentinelReceiveIsPrimaryDownReply[0]))
											if int32(2) < v70 {
												v84 = v65
												v86 = *(*int32)(unsafe.Add(mBase, uint32(v84)+4))
												v87 = *(*int32)(unsafe.Add(mBase, uint32(v86)+28))
												v88 = F_sdsnew(m, v87)
												mBase = m.M
												v89 = m.ExcPending
												if v89 != 0 {
													return
												} else {
													*(*int32)(unsafe.Add(mBase, uint32(l2)+216)) = v88
													v91 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
													v92 = *(*int32)(unsafe.Add(mBase, uint32(v91)+8))
													v93 = *(*int64)(unsafe.Add(mBase, uint32(v92)+8))
													*(*int64)(unsafe.Add(mBase, uint32(l2)+224)) = v93
													m.G0 = v9 + int32(16)
													return
												}
											} else {
												v73 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
												v74 = *(*int32)(unsafe.Add(mBase, uint32(v65)+4))
												v75 = *(*int32)(unsafe.Add(mBase, uint32(v74)+28))
												*(*int64)(unsafe.Add(mBase, uint32(v9)+8)) = v67
												*(*int32)(unsafe.Add(mBase, uint32(v9)+4)) = v75
												*(*int32)(unsafe.Add(mBase, uint32(v9))) = v73
												F__serverLog(m, int32(2), int32(_a_F_sentinelReceiveIsPrimaryDownReply_0), v9)
												mBase = m.M
												v82 = m.ExcPending
												if v82 != 0 {
													return
												} else {
													v83 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
													v84 = v83
													v86 = *(*int32)(unsafe.Add(mBase, uint32(v84)+4))
													v87 = *(*int32)(unsafe.Add(mBase, uint32(v86)+28))
													v88 = F_sdsnew(m, v87)
													mBase = m.M
													v89 = m.ExcPending
													if v89 != 0 {
														return
													} else {
														*(*int32)(unsafe.Add(mBase, uint32(l2)+216)) = v88
														v91 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
														v92 = *(*int32)(unsafe.Add(mBase, uint32(v91)+8))
														v93 = *(*int64)(unsafe.Add(mBase, uint32(v92)+8))
														*(*int64)(unsafe.Add(mBase, uint32(l2)+224)) = v93
														m.G0 = v9 + int32(16)
														return
													}
												}
											}
										}
									}
								} else {
									v58 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v54)+1)))
									if v58 == int32(0) {
										m.G0 = v9 + int32(16)
										return
									} else {
										v61 = *(*int32)(unsafe.Add(mBase, uint32(l2)+216))
										F_sdsfree(m, v61)
										mBase = m.M
										v63 = m.ExcPending
										if v63 != 0 {
											return
										} else {
											v64 = *(*int64)(unsafe.Add(mBase, uint32(l2)+224))
											v65 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
											v66 = *(*int32)(unsafe.Add(mBase, uint32(v65)+8))
											v67 = *(*int64)(unsafe.Add(mBase, uint32(v66)+8))
											if v64 == v67 {
												v84 = v65
												v86 = *(*int32)(unsafe.Add(mBase, uint32(v84)+4))
												v87 = *(*int32)(unsafe.Add(mBase, uint32(v86)+28))
												v88 = F_sdsnew(m, v87)
												mBase = m.M
												v89 = m.ExcPending
												if v89 != 0 {
													return
												} else {
													*(*int32)(unsafe.Add(mBase, uint32(l2)+216)) = v88
													v91 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
													v92 = *(*int32)(unsafe.Add(mBase, uint32(v91)+8))
													v93 = *(*int64)(unsafe.Add(mBase, uint32(v92)+8))
													*(*int64)(unsafe.Add(mBase, uint32(l2)+224)) = v93
													m.G0 = v9 + int32(16)
													return
												}
											} else {
												v70 = *(*int32)(unsafe.Add(mBase, _c_F_sentinelReceiveIsPrimaryDownReply[0]))
												if int32(2) < v70 {
													v84 = v65
													v86 = *(*int32)(unsafe.Add(mBase, uint32(v84)+4))
													v87 = *(*int32)(unsafe.Add(mBase, uint32(v86)+28))
													v88 = F_sdsnew(m, v87)
													mBase = m.M
													v89 = m.ExcPending
													if v89 != 0 {
														return
													} else {
														*(*int32)(unsafe.Add(mBase, uint32(l2)+216)) = v88
														v91 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
														v92 = *(*int32)(unsafe.Add(mBase, uint32(v91)+8))
														v93 = *(*int64)(unsafe.Add(mBase, uint32(v92)+8))
														*(*int64)(unsafe.Add(mBase, uint32(l2)+224)) = v93
														m.G0 = v9 + int32(16)
														return
													}
												} else {
													v73 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
													v74 = *(*int32)(unsafe.Add(mBase, uint32(v65)+4))
													v75 = *(*int32)(unsafe.Add(mBase, uint32(v74)+28))
													*(*int64)(unsafe.Add(mBase, uint32(v9)+8)) = v67
													*(*int32)(unsafe.Add(mBase, uint32(v9)+4)) = v75
													*(*int32)(unsafe.Add(mBase, uint32(v9))) = v73
													F__serverLog(m, int32(2), int32(_a_F_sentinelReceiveIsPrimaryDownReply_0), v9)
													mBase = m.M
													v82 = m.ExcPending
													if v82 != 0 {
														return
													} else {
														v83 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
														v84 = v83
														v86 = *(*int32)(unsafe.Add(mBase, uint32(v84)+4))
														v87 = *(*int32)(unsafe.Add(mBase, uint32(v86)+28))
														v88 = F_sdsnew(m, v87)
														mBase = m.M
														v89 = m.ExcPending
														if v89 != 0 {
															return
														} else {
															*(*int32)(unsafe.Add(mBase, uint32(l2)+216)) = v88
															v91 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
															v92 = *(*int32)(unsafe.Add(mBase, uint32(v91)+8))
															v93 = *(*int64)(unsafe.Add(mBase, uint32(v92)+8))
															*(*int64)(unsafe.Add(mBase, uint32(l2)+224)) = v93
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
		}
	}
}
func F_sentinelRoleCommand(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
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
	var v135 int32
	_ = v135
	var v137 int32
	_ = v137
	var v139 int32
	_ = v139
	F_addReplyArrayLen(m, l0, int32(2))
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
	F_addReplyBulkCBuffer(m, l0, int32(_a_F_sentinelRoleCommand_0), int32(8))
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v12 = *(*int32)(unsafe.Add(mBase, _c_F_sentinelRoleCommand[0]))
	v13 = *(*int32)(unsafe.Add(mBase, uint32(v12)+16))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(v12)+12))
	F_addReplyArrayLen(m, l0, v13+v14)
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v19 = *(*int32)(unsafe.Add(mBase, _c_F_sentinelRoleCommand[0]))
	v20 = F_dictGetIterator(m, v19)
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		goto L1
	} else {
		goto L5
	}
L5:
	;
	goto L7
L6:
	;
	F_dictReleaseIterator(m, v20)
	mBase = m.M
	v139 = m.ExcPending
	if v139 != 0 {
		goto L1
	} else {
		goto L38
	}
L7:
	;
	v32 = v20 + int32(20)
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v20)+16))
	if v33 != 0 {
		goto L11
	} else {
		goto L12
	}
L9:
	;
	if v128 == int32(0) {
		goto L6
	} else {
		goto L35
	}
L10:
	;
	v39 = v32
	v40 = v36
	goto L13
L11:
	;
	v36 = int32(1)
	goto L10
L12:
	;
	v36 = int32(0)
	goto L10
L13:
	;
	switch v40 {
	case 0:
		goto L18
	default:
		goto L17
	}
L15:
	;
	v40 = int32(0)
	goto L13
L16:
	;
	goto L9
L17:
	;
	v120 = *(*int32)(unsafe.Add(mBase, uint32(v39)))
	*(*int32)(unsafe.Add(mBase, uint32(v20)+16)) = v120
	if v120 == int32(0) {
		goto L15
	} else {
		goto L34
	}
L18:
	;
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v20)+4))
	if v44 != int32(-1) {
		v83 = v44
		goto L19
	} else {
		goto L20
	}
L19:
	;
	v84 = int32(1)
	v85 = v83 + v84
	*(*int32)(unsafe.Add(mBase, uint32(v20)+4)) = v85
	v87 = int32(0)
	v90 = *(*int32)(unsafe.Add(mBase, uint32(v20)))
	v91 = *(*int32)(unsafe.Add(mBase, uint32(v20)+8))
	v95 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v90+v91+int32(26)))))
	if v95 == int32(255) {
		goto L28
	} else {
		goto L29
	}
L20:
	;
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v20)+8))
	if v48 != 0 {
		v83 = int32(-1)
		goto L19
	} else {
		goto L21
	}
L21:
	;
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v20)))
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v20)+12))
	if v50 == int32(0) {
		goto L23
	} else {
		goto L24
	}
L22:
	;
	v77 = *(*int32)(unsafe.Add(mBase, uint32(v76)+20))
	if v77 != int32(-1) {
		goto L25
	} else {
		goto L26
	}
L23:
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
	*(*int64)(unsafe.Add(mBase, uint32(v20)+24)) = v73
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v20)))
	v76 = v75
	goto L22
L24:
	;
	v53 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v49)+24)))
	v55 = v53 + int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v49)+24)) = uint16(v55)
	v76 = v49
	goto L22
L25:
	;
	v83 = v77 + int32(-1)
	goto L19
L26:
	;
	v80 = *(*int32)(unsafe.Add(mBase, uint32(v20)+4))
	v83 = v80
	goto L19
L27:
	;
	v110 = int32(2)
	v115 = *(*int32)(unsafe.Add(mBase, uint32(v90+v108<<(uint(v110)%32)+int32(4))))
	v39 = v115 + v109<<(uint(v110)%32)
	v40 = int32(1)
	goto L13
L28:
	;
	v99 = v87
	goto L30
L29:
	;
	v99 = v84 << (uint(v95) % 32)
	goto L30
L30:
	;
	if v85 < v99 {
		v108 = v91
		v109 = v85
		goto L27
	} else {
		goto L31
	}
L31:
	;
	if v91 != 0 {
		v128 = v87
		goto L16
	} else {
		goto L32
	}
L32:
	;
	v101 = *(*int32)(unsafe.Add(mBase, uint32(v90)+20))
	if v101 == int32(-1) {
		v128 = v87
		goto L16
	} else {
		goto L33
	}
L33:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v20)+4)) = int64(4294967296)
	v108 = int32(1)
	v109 = int32(0)
	goto L27
L34:
	;
	v124 = *(*int32)(unsafe.Add(mBase, uint32(v120)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v32))) = v124
	v128 = v120
	goto L16
L35:
	;
	v134 = *(*int32)(unsafe.Add(mBase, uint32(v128)+8))
	goto L36
L36:
	;
	v135 = *(*int32)(unsafe.Add(mBase, uint32(v134)+4))
	F_addReplyBulkCString(m, l0, v135)
	mBase = m.M
	v137 = m.ExcPending
	if v137 != 0 {
		goto L1
	} else {
		goto L37
	}
L37:
	;
	goto L7
L38:
	;
	return
}
func F_sentinelScheduleScriptExecution(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v13 int32
	_ = v13
	var v17 int32
	_ = v17
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
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
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v89 int32
	_ = v89
	var v95 int32
	_ = v95
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v103 int32
	_ = v103
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v113 int32
	_ = v113
	var v116 int32
	_ = v116
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v121 int32
	_ = v121
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v132 int32
	_ = v132
	var v134 int32
	_ = v134
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v157 int32
	_ = v157
	v6 = m.G0
	v8 = v6 - int32(96)
	m.G0 = v8
	*(*int32)(unsafe.Add(mBase, uint32(v8)+92)) = l1
	v13 = int32(1)
	goto L2
L1:
	;
	v38 = F_sdsnew(m, l0)
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L6
	} else {
		goto L9
	}
L2:
	;
	v17 = *(*int32)(unsafe.Add(mBase, uint32(v8)+92))
	*(*int32)(unsafe.Add(mBase, uint32(v8)+92)) = v17 + int32(4)
	v25 = v8 + int32(16) + v13<<(uint(int32(2))%32)
	v26 = *(*int32)(unsafe.Add(mBase, uint32(v17)))
	*(*int32)(unsafe.Add(mBase, uint32(v25))) = v26
	if v26 != 0 {
		goto L4
	} else {
		goto L5
	}
L3:
	;
	v37 = v31
	goto L1
L4:
	;
	v28 = F_sdsnew(m, v26)
	mBase = m.M
	v29 = m.ExcPending
	if v29 != 0 {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	v37 = v13
	goto L1
L6:
	;
	return
L7:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25))) = v28
	v31 = int32(16)
	v33 = v13 + int32(1)
	if v33 != v31 {
		v13 = v33
		goto L2
	} else {
		goto L8
	}
L8:
	;
	goto L3
L9:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = v38
	v42 = F_valkey_malloc(m, int32(32))
	mBase = m.M
	v43 = m.ExcPending
	if v43 != 0 {
		goto L6
	} else {
		goto L10
	}
L10:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v42))) = int64(0)
	v49 = v37<<(uint(int32(2))%32) + int32(4)
	v50 = F_valkey_malloc(m, v49)
	mBase = m.M
	v51 = m.ExcPending
	if v51 != 0 {
		goto L6
	} else {
		goto L11
	}
L11:
	;
	v52 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v42)+24)) = v52
	*(*int64)(unsafe.Add(mBase, uint32(v42)+16)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v42)+8)) = v50
	if v49 == v52 {
		goto L13
	} else {
		goto L14
	}
L12:
	;
	v64 = *(*int32)(unsafe.Add(mBase, _c_F_sentinelScheduleScriptExecution[0]))
	v65 = F_listAddNodeTail(m, v64, v42)
	mBase = m.M
	v66 = m.ExcPending
	if v66 != 0 {
		goto L6
	} else {
		goto L15
	}
L13:
	;
	goto L12
L14:
	;
	v61 = F__emscripten_memcpy_bulkmem(m, v50, v8+int32(16), v49)
	mBase = m.M
	goto L13
L15:
	;
	v68 = *(*int32)(unsafe.Add(mBase, _c_F_sentinelScheduleScriptExecution[0]))
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v68)+20))
	if base.Ui32(v69) < base.Ui32(int32(257)) {
		goto L17
	} else {
		goto L18
	}
L16:
	;
	F__serverAssert(m, int32(_a_F_sentinelScheduleScriptExecution_0), int32(_a_F_sentinelScheduleScriptExecution_1), int32(797))
	mBase = m.M
	v157 = m.ExcPending
	if v157 != 0 {
		goto L6
	} else {
		goto L38
	}
L17:
	;
	m.G0 = v8 + int32(96)
	return
L18:
	;
	v73 = v8 + int32(8)
	v74 = *(*int32)(unsafe.Add(mBase, uint32(v68)))
	*(*int32)(unsafe.Add(mBase, uint32(v73)+4)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v73))) = v74
	goto L19
L19:
	;
	goto L21
L20:
	;
	v141 = *(*int32)(unsafe.Add(mBase, _c_F_sentinelScheduleScriptExecution[0]))
	v142 = *(*int32)(unsafe.Add(mBase, uint32(v141)+20))
	if base.Ui32(int32(257)) <= base.Ui32(v142) {
		goto L16
	} else {
		goto L37
	}
L21:
	;
	v84 = v8 + int32(8)
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v84)))
	if v86 == int32(0) {
		goto L24
	} else {
		goto L25
	}
L22:
	;
	v103 = int32(0)
	v105 = *(*int32)(unsafe.Add(mBase, _c_F_sentinelScheduleScriptExecution[0]))
	F_listDelNode(m, v105, v86)
	mBase = m.M
	v107 = m.ExcPending
	if v107 != 0 {
		goto L6
	} else {
		goto L28
	}
L23:
	;
	if v86 == int32(0) {
		goto L20
	} else {
		goto L26
	}
L24:
	;
	goto L23
L25:
	;
	v89 = *(*int32)(unsafe.Add(mBase, uint32(v84)+4))
	v95 = *(*int32)(unsafe.Add(mBase, uint32(v86+base.B2i32(v89 == int32(0))<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v84))) = v95
	goto L24
L26:
	;
	v99 = *(*int32)(unsafe.Add(mBase, uint32(v86)+8))
	v100 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v99))))
	if v100&int32(1) != 0 {
		goto L21
	} else {
		goto L27
	}
L27:
	;
	goto L22
L28:
	;
	v108 = *(*int32)(unsafe.Add(mBase, uint32(v99)+8))
	v109 = *(*int32)(unsafe.Add(mBase, uint32(v108)))
	if v109 == int32(0) {
		v126 = v108
		goto L29
	} else {
		goto L30
	}
L29:
	;
	F_valkey_free(m, v126)
	mBase = m.M
	v132 = m.ExcPending
	if v132 != 0 {
		goto L6
	} else {
		goto L35
	}
L30:
	;
	v113 = v109
	v116 = v103
	goto L31
L31:
	;
	F_sdsfree(m, v113)
	mBase = m.M
	v118 = m.ExcPending
	if v118 != 0 {
		goto L6
	} else {
		goto L33
	}
L32:
	;
	v126 = v119
	goto L29
L33:
	;
	v119 = *(*int32)(unsafe.Add(mBase, uint32(v99)+8))
	v121 = v116 + int32(1)
	v125 = *(*int32)(unsafe.Add(mBase, uint32(v119+v121<<(uint(int32(2))%32))))
	if v125 != 0 {
		v113 = v125
		v116 = v121
		goto L31
	} else {
		goto L34
	}
L34:
	;
	goto L32
L35:
	;
	F_valkey_free(m, v99)
	mBase = m.M
	v134 = m.ExcPending
	if v134 != 0 {
		goto L6
	} else {
		goto L36
	}
L36:
	;
	goto L20
L37:
	;
	goto L17
L38:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_sentinelSendHello(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v24 int32
	_ = v24
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
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v75 int32
	_ = v75
	var v76 int64
	_ = v76
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v82 int64
	_ = v82
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
	var v112 int32
	_ = v112
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
	var v119 int32
	_ = v119
	var v123 int32
	_ = v123
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v141 int32
	_ = v141
	v9 = m.G0
	v11 = v9 - int32(1184)
	m.G0 = v11
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v13&int32(1) != 0 {
		v18 = l0
		v19 = v13
	} else {
		v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+192))
		v17 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
		v18 = v16
		v19 = v17
	}
	if v19&int32(64) == int32(0) {
		v31 = v18
	} else {
		v24 = *(*int32)(unsafe.Add(mBase, uint32(v18)+280))
		if v24 == int32(0) {
			v31 = v18
		} else {
			v27 = *(*int32)(unsafe.Add(mBase, uint32(v18)+240))
			if int32(4) < v27 {
				v31 = v24
			} else {
				v31 = v18
			}
		}
	}
	v33 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v33)+4))
	if v34 != 0 {
		v141 = int32(-1)
		m.G0 = v11 + int32(1184)
		return v141
	} else {
		v35 = *(*int32)(unsafe.Add(mBase, uint32(v31)+24))
		v37 = *(*int32)(unsafe.Add(mBase, _c_F_sentinelSendHello[0]))
		if v37 != 0 {
			v55 = v37
			v57 = *(*int32)(unsafe.Add(mBase, _c_F_sentinelSendHello[1]))
			if v57 != 0 {
				v66 = v57
			} else {
				v58 = int32(_a_F_sentinelSendHello_0)
				v59 = *(*int32)(unsafe.Add(mBase, _c_F_sentinelSendHello[2]))
				v61 = *(*int32)(unsafe.Add(mBase, _c_F_sentinelSendHello[3]))
				if v59 != 0 {
					v62 = v59
				} else {
					v62 = v61
				}
				v64 = *(*int32)(unsafe.Add(mBase, _c_F_sentinelSendHello[4]))
				if v64 != 0 {
					v65 = v62
				} else {
					v65 = v61
				}
				v66 = v65
			}
			v68 = int32(0)
			v69 = *(*int32)(unsafe.Add(mBase, _c_F_sentinelSendHello[5]))
			v75 = *(*int32)(unsafe.Add(mBase, uint32(v35+base.B2i32(v69 == v68)<<(uint(int32(2))%32))))
			v76 = *(*int64)(unsafe.Add(mBase, uint32(v18)+16))
			v77 = *(*int32)(unsafe.Add(mBase, uint32(v35)+8))
			v78 = *(*int32)(unsafe.Add(mBase, uint32(v18)+4))
			v82 = *(*int64)(unsafe.Add(mBase, _c_F_sentinelSendHello[6]))
			*(*int64)(unsafe.Add(mBase, uint32(v11+int32(32)))) = v82
			*(*int32)(unsafe.Add(mBase, uint32(v11+int32(40)))) = v78
			*(*int32)(unsafe.Add(mBase, uint32(v11+int32(44)))) = v75
			*(*int32)(unsafe.Add(mBase, uint32(v11+int32(48)))) = v77
			*(*int64)(unsafe.Add(mBase, uint32(v11+int32(56)))) = v76
			*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = v55
			*(*int32)(unsafe.Add(mBase, uint32(v11)+20)) = v66
			*(*int32)(unsafe.Add(mBase, uint32(v11)+24)) = int32(_a_F_sentinelSendHello_1)
			v106 = F_snprintf(m, v11+int32(64), int32(1070), int32(_a_F_sentinelSendHello_2), v11+int32(16))
			mBase = m.M
			v107 = m.ExcPending
			if v107 != 0 {
				return int32(0)
			} else {
				v108 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
				v109 = *(*int32)(unsafe.Add(mBase, uint32(v108)+12))
				v111 = F_sdsnew(m, int32(_a_F_sentinelSendHello_3))
				mBase = m.M
				v112 = m.ExcPending
				if v112 != 0 {
					return int32(0)
				} else {
					v113 = *(*int32)(unsafe.Add(mBase, uint32(l0)+192))
					if v113 != 0 {
						v114 = v113
					} else {
						v114 = l0
					}
					v115 = *(*int32)(unsafe.Add(mBase, uint32(v114)+104))
					v116 = F_dictFetchValue(m, v115, v111)
					mBase = m.M
					v117 = m.ExcPending
					if v117 != 0 {
						return int32(0)
					} else {
						F_sdsfree(m, v111)
						mBase = m.M
						v119 = m.ExcPending
						if v119 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v11)+4)) = int32(_a_F_sentinelSendHello_4)
							if v116 != 0 {
								v123 = v116
							} else {
								v123 = int32(_a_F_sentinelSendHello_3)
							}
							*(*int32)(unsafe.Add(mBase, uint32(v11))) = v123
							*(*int32)(unsafe.Add(mBase, uint32(v11)+8)) = v11 + int32(64)
							v131 = F_valkeyAsyncCommand(m, v109, int32(1017), l0, int32(_a_F_sentinelSendHello_5), v11)
							mBase = m.M
							v132 = m.ExcPending
							if v132 != 0 {
								return int32(0)
							} else {
								if v131 != 0 {
									v141 = int32(-1)
								} else {
									v133 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
									v134 = *(*int32)(unsafe.Add(mBase, uint32(v133)+8))
									*(*int32)(unsafe.Add(mBase, uint32(v133)+8)) = v134 + int32(1)
									v141 = int32(0)
								}
								m.G0 = v11 + int32(1184)
								return v141
							}
						}
					}
				}
			}
		} else {
			v40 = v11 + int32(1136)
			v41 = *(*int32)(unsafe.Add(mBase, uint32(v33)+12))
			v42 = *(*int32)(unsafe.Add(mBase, uint32(v41)+136))
			v46 = int32(0)
			v48 = F_anetFdToString(m, v42, v40, int32(46), v46, v46)
			mBase = m.M
			v51 = m.ExcPending
			if v51 != 0 {
				return int32(0)
			} else {
				if v48 == int32(-1) {
					v141 = int32(-1)
					m.G0 = v11 + int32(1184)
					return v141
				} else {
					v55 = v40
					v57 = *(*int32)(unsafe.Add(mBase, _c_F_sentinelSendHello[1]))
					if v57 != 0 {
						v66 = v57
					} else {
						v58 = int32(_a_F_sentinelSendHello_0)
						v59 = *(*int32)(unsafe.Add(mBase, _c_F_sentinelSendHello[2]))
						v61 = *(*int32)(unsafe.Add(mBase, _c_F_sentinelSendHello[3]))
						if v59 != 0 {
							v62 = v59
						} else {
							v62 = v61
						}
						v64 = *(*int32)(unsafe.Add(mBase, _c_F_sentinelSendHello[4]))
						if v64 != 0 {
							v65 = v62
						} else {
							v65 = v61
						}
						v66 = v65
					}
					v68 = int32(0)
					v69 = *(*int32)(unsafe.Add(mBase, _c_F_sentinelSendHello[5]))
					v75 = *(*int32)(unsafe.Add(mBase, uint32(v35+base.B2i32(v69 == v68)<<(uint(int32(2))%32))))
					v76 = *(*int64)(unsafe.Add(mBase, uint32(v18)+16))
					v77 = *(*int32)(unsafe.Add(mBase, uint32(v35)+8))
					v78 = *(*int32)(unsafe.Add(mBase, uint32(v18)+4))
					v82 = *(*int64)(unsafe.Add(mBase, _c_F_sentinelSendHello[6]))
					*(*int64)(unsafe.Add(mBase, uint32(v11+int32(32)))) = v82
					*(*int32)(unsafe.Add(mBase, uint32(v11+int32(40)))) = v78
					*(*int32)(unsafe.Add(mBase, uint32(v11+int32(44)))) = v75
					*(*int32)(unsafe.Add(mBase, uint32(v11+int32(48)))) = v77
					*(*int64)(unsafe.Add(mBase, uint32(v11+int32(56)))) = v76
					*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = v55
					*(*int32)(unsafe.Add(mBase, uint32(v11)+20)) = v66
					*(*int32)(unsafe.Add(mBase, uint32(v11)+24)) = int32(_a_F_sentinelSendHello_1)
					v106 = F_snprintf(m, v11+int32(64), int32(1070), int32(_a_F_sentinelSendHello_2), v11+int32(16))
					mBase = m.M
					v107 = m.ExcPending
					if v107 != 0 {
						return int32(0)
					} else {
						v108 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
						v109 = *(*int32)(unsafe.Add(mBase, uint32(v108)+12))
						v111 = F_sdsnew(m, int32(_a_F_sentinelSendHello_3))
						mBase = m.M
						v112 = m.ExcPending
						if v112 != 0 {
							return int32(0)
						} else {
							v113 = *(*int32)(unsafe.Add(mBase, uint32(l0)+192))
							if v113 != 0 {
								v114 = v113
							} else {
								v114 = l0
							}
							v115 = *(*int32)(unsafe.Add(mBase, uint32(v114)+104))
							v116 = F_dictFetchValue(m, v115, v111)
							mBase = m.M
							v117 = m.ExcPending
							if v117 != 0 {
								return int32(0)
							} else {
								F_sdsfree(m, v111)
								mBase = m.M
								v119 = m.ExcPending
								if v119 != 0 {
									return int32(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v11)+4)) = int32(_a_F_sentinelSendHello_4)
									if v116 != 0 {
										v123 = v116
									} else {
										v123 = int32(_a_F_sentinelSendHello_3)
									}
									*(*int32)(unsafe.Add(mBase, uint32(v11))) = v123
									*(*int32)(unsafe.Add(mBase, uint32(v11)+8)) = v11 + int32(64)
									v131 = F_valkeyAsyncCommand(m, v109, int32(1017), l0, int32(_a_F_sentinelSendHello_5), v11)
									mBase = m.M
									v132 = m.ExcPending
									if v132 != 0 {
										return int32(0)
									} else {
										if v131 != 0 {
											v141 = int32(-1)
										} else {
											v133 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
											v134 = *(*int32)(unsafe.Add(mBase, uint32(v133)+8))
											*(*int32)(unsafe.Add(mBase, uint32(v133)+8)) = v134 + int32(1)
											v141 = int32(0)
										}
										m.G0 = v11 + int32(1184)
										return v141
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
func F_sentinelSendPeriodicCommands(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v14 int64
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v22 int32
	_ = v22
	var v27 int64
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v32 int64
	_ = v32
	var v37 int64
	_ = v37
	var v38 int64
	_ = v38
	var v39 int64
	_ = v39
	var v41 int64
	_ = v41
	var v43 int64
	_ = v43
	var v46 int64
	_ = v46
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
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
	var v78 int64
	_ = v78
	var v81 int64
	_ = v81
	var v84 int64
	_ = v84
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v88 int64
	_ = v88
	var v91 int64
	_ = v91
	var v93 int32
	_ = v93
	var v98 int32
	_ = v98
	var v101 int32
	_ = v101
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	v10 = m.G0
	v12 = v10 - int32(16)
	m.G0 = v12
	v14 = F_mstime(m)
	mBase = m.M
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v16 = *(*int32)(unsafe.Add(mBase, uint32(v15)+4))
	if v16 != 0 {
		m.G0 = v12 + int32(16)
		return
	} else {
		v17 = *(*int32)(unsafe.Add(mBase, uint32(v15)+8))
		v18 = *(*int32)(unsafe.Add(mBase, uint32(v15)))
		if v18*int32(100) <= v17 {
			m.G0 = v12 + int32(16)
			return
		} else {
			v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			if v22&int32(2) == int32(0) {
				v37 = *(*int64)(unsafe.Add(mBase, _c_F_sentinelSendPeriodicCommands[0]))
				v38 = v37
			} else {
				v27 = int64(1000)
				v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+192))
				v29 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28))))
				if v29&int32(80) != 0 {
					v38 = v27
				} else {
					v32 = *(*int64)(unsafe.Add(mBase, uint32(l0)+168))
					if v32 != int64(0) {
						v38 = v27
					} else {
						v37 = *(*int64)(unsafe.Add(mBase, _c_F_sentinelSendPeriodicCommands[0]))
						v38 = v37
					}
				}
			}
			v39 = *(*int64)(unsafe.Add(mBase, uint32(l0)+72))
			v41 = *(*int64)(unsafe.Add(mBase, _c_F_sentinelSendPeriodicCommands[1]))
			if v39 < v41 {
				v43 = v39
			} else {
				v43 = v41
			}
			if v22&int32(4) != 0 {
				v74 = v15
				v78 = *(*int64)(unsafe.Add(mBase, uint32(v74)+72))
				if v14-v78 <= v43 {
					v88 = *(*int64)(unsafe.Add(mBase, uint32(l0)+32))
					v91 = *(*int64)(unsafe.Add(mBase, _c_F_sentinelSendPeriodicCommands[2]))
					if v14-v88 <= v91 {
						m.G0 = v12 + int32(16)
						return
					} else {
						v93 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+1)))
						if v93&int32(64) == int32(0) {
							v101 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
							if v101 == int32(2) {
								m.G0 = v12 + int32(16)
								return
							} else {
								v104 = F_sentinelSendHello(m, l0)
								mBase = m.M
								v105 = m.ExcPending
								if v105 != 0 {
									return
								} else {
									m.G0 = v12 + int32(16)
									return
								}
							}
						} else {
							v98 = *(*int32)(unsafe.Add(mBase, uint32(l0)+240))
							if v98 < int32(5) {
								m.G0 = v12 + int32(16)
								return
							} else {
								v101 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
								if v101 == int32(2) {
									m.G0 = v12 + int32(16)
									return
								} else {
									v104 = F_sentinelSendHello(m, l0)
									mBase = m.M
									v105 = m.ExcPending
									if v105 != 0 {
										return
									} else {
										m.G0 = v12 + int32(16)
										return
									}
								}
							}
						}
					}
				} else {
					v81 = *(*int64)(unsafe.Add(mBase, uint32(v74)+64))
					v84 = base.I64_div_s(v43, int64(2))
					if v14-v81 <= v84 {
						v88 = *(*int64)(unsafe.Add(mBase, uint32(l0)+32))
						v91 = *(*int64)(unsafe.Add(mBase, _c_F_sentinelSendPeriodicCommands[2]))
						if v14-v88 <= v91 {
							m.G0 = v12 + int32(16)
							return
						} else {
							v93 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+1)))
							if v93&int32(64) == int32(0) {
								v101 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
								if v101 == int32(2) {
									m.G0 = v12 + int32(16)
									return
								} else {
									v104 = F_sentinelSendHello(m, l0)
									mBase = m.M
									v105 = m.ExcPending
									if v105 != 0 {
										return
									} else {
										m.G0 = v12 + int32(16)
										return
									}
								}
							} else {
								v98 = *(*int32)(unsafe.Add(mBase, uint32(l0)+240))
								if v98 < int32(5) {
									m.G0 = v12 + int32(16)
									return
								} else {
									v101 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
									if v101 == int32(2) {
										m.G0 = v12 + int32(16)
										return
									} else {
										v104 = F_sentinelSendHello(m, l0)
										mBase = m.M
										v105 = m.ExcPending
										if v105 != 0 {
											return
										} else {
											m.G0 = v12 + int32(16)
											return
										}
									}
								}
							}
						}
					} else {
						v86 = F_sentinelSendPing(m, l0)
						mBase = m.M
						v87 = m.ExcPending
						if v87 != 0 {
							return
						} else {
							v88 = *(*int64)(unsafe.Add(mBase, uint32(l0)+32))
							v91 = *(*int64)(unsafe.Add(mBase, _c_F_sentinelSendPeriodicCommands[2]))
							if v14-v88 <= v91 {
								m.G0 = v12 + int32(16)
								return
							} else {
								v93 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+1)))
								if v93&int32(64) == int32(0) {
									v101 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
									if v101 == int32(2) {
										m.G0 = v12 + int32(16)
										return
									} else {
										v104 = F_sentinelSendHello(m, l0)
										mBase = m.M
										v105 = m.ExcPending
										if v105 != 0 {
											return
										} else {
											m.G0 = v12 + int32(16)
											return
										}
									}
								} else {
									v98 = *(*int32)(unsafe.Add(mBase, uint32(l0)+240))
									if v98 < int32(5) {
										m.G0 = v12 + int32(16)
										return
									} else {
										v101 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
										if v101 == int32(2) {
											m.G0 = v12 + int32(16)
											return
										} else {
											v104 = F_sentinelSendHello(m, l0)
											mBase = m.M
											v105 = m.ExcPending
											if v105 != 0 {
												return
											} else {
												m.G0 = v12 + int32(16)
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
				v46 = *(*int64)(unsafe.Add(mBase, uint32(l0)+96))
				if v46 == int64(0) {
					v51 = *(*int32)(unsafe.Add(mBase, uint32(v15)+12))
					v53 = F_sdsnew(m, int32(_a_F_sentinelSendPeriodicCommands_0))
					mBase = m.M
					v54 = m.ExcPending
					if v54 != 0 {
						return
					} else {
						v55 = *(*int32)(unsafe.Add(mBase, uint32(l0)+192))
						if v55 != 0 {
							v56 = v55
						} else {
							v56 = l0
						}
						v57 = *(*int32)(unsafe.Add(mBase, uint32(v56)+104))
						v58 = F_dictFetchValue(m, v57, v53)
						mBase = m.M
						v59 = m.ExcPending
						if v59 != 0 {
							return
						} else {
							F_sdsfree(m, v53)
							mBase = m.M
							v61 = m.ExcPending
							if v61 != 0 {
								return
							} else {
								if v58 != 0 {
									v63 = v58
								} else {
									v63 = int32(_a_F_sentinelSendPeriodicCommands_0)
								}
								*(*int32)(unsafe.Add(mBase, uint32(v12))) = v63
								v67 = F_valkeyAsyncCommand(m, v51, int32(1018), l0, int32(_a_F_sentinelSendPeriodicCommands_1), v12)
								mBase = m.M
								v68 = m.ExcPending
								if v68 != 0 {
									return
								} else {
									v69 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
									if v67 != 0 {
										v74 = v69
									} else {
										v70 = *(*int32)(unsafe.Add(mBase, uint32(v69)+8))
										*(*int32)(unsafe.Add(mBase, uint32(v69)+8)) = v70 + int32(1)
										v74 = v69
									}
									v78 = *(*int64)(unsafe.Add(mBase, uint32(v74)+72))
									if v14-v78 <= v43 {
										v88 = *(*int64)(unsafe.Add(mBase, uint32(l0)+32))
										v91 = *(*int64)(unsafe.Add(mBase, _c_F_sentinelSendPeriodicCommands[2]))
										if v14-v88 <= v91 {
											m.G0 = v12 + int32(16)
											return
										} else {
											v93 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+1)))
											if v93&int32(64) == int32(0) {
												v101 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
												if v101 == int32(2) {
													m.G0 = v12 + int32(16)
													return
												} else {
													v104 = F_sentinelSendHello(m, l0)
													mBase = m.M
													v105 = m.ExcPending
													if v105 != 0 {
														return
													} else {
														m.G0 = v12 + int32(16)
														return
													}
												}
											} else {
												v98 = *(*int32)(unsafe.Add(mBase, uint32(l0)+240))
												if v98 < int32(5) {
													m.G0 = v12 + int32(16)
													return
												} else {
													v101 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
													if v101 == int32(2) {
														m.G0 = v12 + int32(16)
														return
													} else {
														v104 = F_sentinelSendHello(m, l0)
														mBase = m.M
														v105 = m.ExcPending
														if v105 != 0 {
															return
														} else {
															m.G0 = v12 + int32(16)
															return
														}
													}
												}
											}
										}
									} else {
										v81 = *(*int64)(unsafe.Add(mBase, uint32(v74)+64))
										v84 = base.I64_div_s(v43, int64(2))
										if v14-v81 <= v84 {
											v88 = *(*int64)(unsafe.Add(mBase, uint32(l0)+32))
											v91 = *(*int64)(unsafe.Add(mBase, _c_F_sentinelSendPeriodicCommands[2]))
											if v14-v88 <= v91 {
												m.G0 = v12 + int32(16)
												return
											} else {
												v93 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+1)))
												if v93&int32(64) == int32(0) {
													v101 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
													if v101 == int32(2) {
														m.G0 = v12 + int32(16)
														return
													} else {
														v104 = F_sentinelSendHello(m, l0)
														mBase = m.M
														v105 = m.ExcPending
														if v105 != 0 {
															return
														} else {
															m.G0 = v12 + int32(16)
															return
														}
													}
												} else {
													v98 = *(*int32)(unsafe.Add(mBase, uint32(l0)+240))
													if v98 < int32(5) {
														m.G0 = v12 + int32(16)
														return
													} else {
														v101 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
														if v101 == int32(2) {
															m.G0 = v12 + int32(16)
															return
														} else {
															v104 = F_sentinelSendHello(m, l0)
															mBase = m.M
															v105 = m.ExcPending
															if v105 != 0 {
																return
															} else {
																m.G0 = v12 + int32(16)
																return
															}
														}
													}
												}
											}
										} else {
											v86 = F_sentinelSendPing(m, l0)
											mBase = m.M
											v87 = m.ExcPending
											if v87 != 0 {
												return
											} else {
												v88 = *(*int64)(unsafe.Add(mBase, uint32(l0)+32))
												v91 = *(*int64)(unsafe.Add(mBase, _c_F_sentinelSendPeriodicCommands[2]))
												if v14-v88 <= v91 {
													m.G0 = v12 + int32(16)
													return
												} else {
													v93 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+1)))
													if v93&int32(64) == int32(0) {
														v101 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
														if v101 == int32(2) {
															m.G0 = v12 + int32(16)
															return
														} else {
															v104 = F_sentinelSendHello(m, l0)
															mBase = m.M
															v105 = m.ExcPending
															if v105 != 0 {
																return
															} else {
																m.G0 = v12 + int32(16)
																return
															}
														}
													} else {
														v98 = *(*int32)(unsafe.Add(mBase, uint32(l0)+240))
														if v98 < int32(5) {
															m.G0 = v12 + int32(16)
															return
														} else {
															v101 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
															if v101 == int32(2) {
																m.G0 = v12 + int32(16)
																return
															} else {
																v104 = F_sentinelSendHello(m, l0)
																mBase = m.M
																v105 = m.ExcPending
																if v105 != 0 {
																	return
																} else {
																	m.G0 = v12 + int32(16)
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
				} else {
					if v14-v46 <= v38 {
						v74 = v15
						v78 = *(*int64)(unsafe.Add(mBase, uint32(v74)+72))
						if v14-v78 <= v43 {
							v88 = *(*int64)(unsafe.Add(mBase, uint32(l0)+32))
							v91 = *(*int64)(unsafe.Add(mBase, _c_F_sentinelSendPeriodicCommands[2]))
							if v14-v88 <= v91 {
								m.G0 = v12 + int32(16)
								return
							} else {
								v93 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+1)))
								if v93&int32(64) == int32(0) {
									v101 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
									if v101 == int32(2) {
										m.G0 = v12 + int32(16)
										return
									} else {
										v104 = F_sentinelSendHello(m, l0)
										mBase = m.M
										v105 = m.ExcPending
										if v105 != 0 {
											return
										} else {
											m.G0 = v12 + int32(16)
											return
										}
									}
								} else {
									v98 = *(*int32)(unsafe.Add(mBase, uint32(l0)+240))
									if v98 < int32(5) {
										m.G0 = v12 + int32(16)
										return
									} else {
										v101 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
										if v101 == int32(2) {
											m.G0 = v12 + int32(16)
											return
										} else {
											v104 = F_sentinelSendHello(m, l0)
											mBase = m.M
											v105 = m.ExcPending
											if v105 != 0 {
												return
											} else {
												m.G0 = v12 + int32(16)
												return
											}
										}
									}
								}
							}
						} else {
							v81 = *(*int64)(unsafe.Add(mBase, uint32(v74)+64))
							v84 = base.I64_div_s(v43, int64(2))
							if v14-v81 <= v84 {
								v88 = *(*int64)(unsafe.Add(mBase, uint32(l0)+32))
								v91 = *(*int64)(unsafe.Add(mBase, _c_F_sentinelSendPeriodicCommands[2]))
								if v14-v88 <= v91 {
									m.G0 = v12 + int32(16)
									return
								} else {
									v93 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+1)))
									if v93&int32(64) == int32(0) {
										v101 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
										if v101 == int32(2) {
											m.G0 = v12 + int32(16)
											return
										} else {
											v104 = F_sentinelSendHello(m, l0)
											mBase = m.M
											v105 = m.ExcPending
											if v105 != 0 {
												return
											} else {
												m.G0 = v12 + int32(16)
												return
											}
										}
									} else {
										v98 = *(*int32)(unsafe.Add(mBase, uint32(l0)+240))
										if v98 < int32(5) {
											m.G0 = v12 + int32(16)
											return
										} else {
											v101 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
											if v101 == int32(2) {
												m.G0 = v12 + int32(16)
												return
											} else {
												v104 = F_sentinelSendHello(m, l0)
												mBase = m.M
												v105 = m.ExcPending
												if v105 != 0 {
													return
												} else {
													m.G0 = v12 + int32(16)
													return
												}
											}
										}
									}
								}
							} else {
								v86 = F_sentinelSendPing(m, l0)
								mBase = m.M
								v87 = m.ExcPending
								if v87 != 0 {
									return
								} else {
									v88 = *(*int64)(unsafe.Add(mBase, uint32(l0)+32))
									v91 = *(*int64)(unsafe.Add(mBase, _c_F_sentinelSendPeriodicCommands[2]))
									if v14-v88 <= v91 {
										m.G0 = v12 + int32(16)
										return
									} else {
										v93 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+1)))
										if v93&int32(64) == int32(0) {
											v101 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
											if v101 == int32(2) {
												m.G0 = v12 + int32(16)
												return
											} else {
												v104 = F_sentinelSendHello(m, l0)
												mBase = m.M
												v105 = m.ExcPending
												if v105 != 0 {
													return
												} else {
													m.G0 = v12 + int32(16)
													return
												}
											}
										} else {
											v98 = *(*int32)(unsafe.Add(mBase, uint32(l0)+240))
											if v98 < int32(5) {
												m.G0 = v12 + int32(16)
												return
											} else {
												v101 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
												if v101 == int32(2) {
													m.G0 = v12 + int32(16)
													return
												} else {
													v104 = F_sentinelSendHello(m, l0)
													mBase = m.M
													v105 = m.ExcPending
													if v105 != 0 {
														return
													} else {
														m.G0 = v12 + int32(16)
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
						v51 = *(*int32)(unsafe.Add(mBase, uint32(v15)+12))
						v53 = F_sdsnew(m, int32(_a_F_sentinelSendPeriodicCommands_0))
						mBase = m.M
						v54 = m.ExcPending
						if v54 != 0 {
							return
						} else {
							v55 = *(*int32)(unsafe.Add(mBase, uint32(l0)+192))
							if v55 != 0 {
								v56 = v55
							} else {
								v56 = l0
							}
							v57 = *(*int32)(unsafe.Add(mBase, uint32(v56)+104))
							v58 = F_dictFetchValue(m, v57, v53)
							mBase = m.M
							v59 = m.ExcPending
							if v59 != 0 {
								return
							} else {
								F_sdsfree(m, v53)
								mBase = m.M
								v61 = m.ExcPending
								if v61 != 0 {
									return
								} else {
									if v58 != 0 {
										v63 = v58
									} else {
										v63 = int32(_a_F_sentinelSendPeriodicCommands_0)
									}
									*(*int32)(unsafe.Add(mBase, uint32(v12))) = v63
									v67 = F_valkeyAsyncCommand(m, v51, int32(1018), l0, int32(_a_F_sentinelSendPeriodicCommands_1), v12)
									mBase = m.M
									v68 = m.ExcPending
									if v68 != 0 {
										return
									} else {
										v69 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
										if v67 != 0 {
											v74 = v69
										} else {
											v70 = *(*int32)(unsafe.Add(mBase, uint32(v69)+8))
											*(*int32)(unsafe.Add(mBase, uint32(v69)+8)) = v70 + int32(1)
											v74 = v69
										}
										v78 = *(*int64)(unsafe.Add(mBase, uint32(v74)+72))
										if v14-v78 <= v43 {
											v88 = *(*int64)(unsafe.Add(mBase, uint32(l0)+32))
											v91 = *(*int64)(unsafe.Add(mBase, _c_F_sentinelSendPeriodicCommands[2]))
											if v14-v88 <= v91 {
												m.G0 = v12 + int32(16)
												return
											} else {
												v93 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+1)))
												if v93&int32(64) == int32(0) {
													v101 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
													if v101 == int32(2) {
														m.G0 = v12 + int32(16)
														return
													} else {
														v104 = F_sentinelSendHello(m, l0)
														mBase = m.M
														v105 = m.ExcPending
														if v105 != 0 {
															return
														} else {
															m.G0 = v12 + int32(16)
															return
														}
													}
												} else {
													v98 = *(*int32)(unsafe.Add(mBase, uint32(l0)+240))
													if v98 < int32(5) {
														m.G0 = v12 + int32(16)
														return
													} else {
														v101 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
														if v101 == int32(2) {
															m.G0 = v12 + int32(16)
															return
														} else {
															v104 = F_sentinelSendHello(m, l0)
															mBase = m.M
															v105 = m.ExcPending
															if v105 != 0 {
																return
															} else {
																m.G0 = v12 + int32(16)
																return
															}
														}
													}
												}
											}
										} else {
											v81 = *(*int64)(unsafe.Add(mBase, uint32(v74)+64))
											v84 = base.I64_div_s(v43, int64(2))
											if v14-v81 <= v84 {
												v88 = *(*int64)(unsafe.Add(mBase, uint32(l0)+32))
												v91 = *(*int64)(unsafe.Add(mBase, _c_F_sentinelSendPeriodicCommands[2]))
												if v14-v88 <= v91 {
													m.G0 = v12 + int32(16)
													return
												} else {
													v93 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+1)))
													if v93&int32(64) == int32(0) {
														v101 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
														if v101 == int32(2) {
															m.G0 = v12 + int32(16)
															return
														} else {
															v104 = F_sentinelSendHello(m, l0)
															mBase = m.M
															v105 = m.ExcPending
															if v105 != 0 {
																return
															} else {
																m.G0 = v12 + int32(16)
																return
															}
														}
													} else {
														v98 = *(*int32)(unsafe.Add(mBase, uint32(l0)+240))
														if v98 < int32(5) {
															m.G0 = v12 + int32(16)
															return
														} else {
															v101 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
															if v101 == int32(2) {
																m.G0 = v12 + int32(16)
																return
															} else {
																v104 = F_sentinelSendHello(m, l0)
																mBase = m.M
																v105 = m.ExcPending
																if v105 != 0 {
																	return
																} else {
																	m.G0 = v12 + int32(16)
																	return
																}
															}
														}
													}
												}
											} else {
												v86 = F_sentinelSendPing(m, l0)
												mBase = m.M
												v87 = m.ExcPending
												if v87 != 0 {
													return
												} else {
													v88 = *(*int64)(unsafe.Add(mBase, uint32(l0)+32))
													v91 = *(*int64)(unsafe.Add(mBase, _c_F_sentinelSendPeriodicCommands[2]))
													if v14-v88 <= v91 {
														m.G0 = v12 + int32(16)
														return
													} else {
														v93 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+1)))
														if v93&int32(64) == int32(0) {
															v101 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
															if v101 == int32(2) {
																m.G0 = v12 + int32(16)
																return
															} else {
																v104 = F_sentinelSendHello(m, l0)
																mBase = m.M
																v105 = m.ExcPending
																if v105 != 0 {
																	return
																} else {
																	m.G0 = v12 + int32(16)
																	return
																}
															}
														} else {
															v98 = *(*int32)(unsafe.Add(mBase, uint32(l0)+240))
															if v98 < int32(5) {
																m.G0 = v12 + int32(16)
																return
															} else {
																v101 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
																if v101 == int32(2) {
																	m.G0 = v12 + int32(16)
																	return
																} else {
																	v104 = F_sentinelSendHello(m, l0)
																	mBase = m.M
																	v105 = m.ExcPending
																	if v105 != 0 {
																		return
																	} else {
																		m.G0 = v12 + int32(16)
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
func F_sentinelSendPing(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v39 int64
	_ = v39
	var v40 int32
	_ = v40
	var v42 int64
	_ = v42
	var v47 int32
	_ = v47
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(v11)+12))
	v14 = F_sdsnew(m, int32(_a_F_sentinelSendPing_0))
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		return int32(0)
	} else {
		v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+192))
		if v18 != 0 {
			v19 = v18
		} else {
			v19 = l0
		}
		v20 = *(*int32)(unsafe.Add(mBase, uint32(v19)+104))
		v21 = F_dictFetchValue(m, v20, v14)
		mBase = m.M
		v22 = m.ExcPending
		if v22 != 0 {
			return int32(0)
		} else {
			F_sdsfree(m, v14)
			mBase = m.M
			v24 = m.ExcPending
			if v24 != 0 {
				return int32(0)
			} else {
				if v21 != 0 {
					v26 = v21
				} else {
					v26 = int32(_a_F_sentinelSendPing_0)
				}
				*(*int32)(unsafe.Add(mBase, uint32(v9))) = v26
				v31 = F_valkeyAsyncCommand(m, v12, int32(1014), l0, int32(_a_F_sentinelSendPing_1), v9)
				mBase = m.M
				v32 = m.ExcPending
				if v32 != 0 {
					return int32(0)
				} else {
					if v31 != 0 {
						v47 = int32(0)
					} else {
						v33 = int32(1)
						v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
						v35 = *(*int32)(unsafe.Add(mBase, uint32(v34)+8))
						*(*int32)(unsafe.Add(mBase, uint32(v34)+8)) = v35 + v33
						v39 = F_mstime(m)
						mBase = m.M
						v40 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
						*(*int64)(unsafe.Add(mBase, uint32(v40)+64)) = v39
						v42 = *(*int64)(unsafe.Add(mBase, uint32(v40)+56))
						if v42 != int64(0) {
							v47 = v33
						} else {
							*(*int64)(unsafe.Add(mBase, uint32(v40)+56)) = v39
							v47 = v33
						}
					}
					m.G0 = v9 + int32(16)
					return v47
				}
			}
		}
	}
}
func F_sentinelUpdateSentinelAddressInAllPrimaries(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v55 int64
	_ = v55
	var v56 int64
	_ = v56
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
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v93 int32
	_ = v93
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v113 int32
	_ = v113
	var v118 int32
	_ = v118
	var v122 int32
	_ = v122
	var v126 int32
	_ = v126
	var v131 int32
	_ = v131
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v156 int32
	_ = v156
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v171 int32
	_ = v171
	var v175 int32
	_ = v175
	var v182 int32
	_ = v182
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v190 int32
	_ = v190
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
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v205 int32
	_ = v205
	var v211 int32
	_ = v211
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v224 int32
	_ = v224
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v232 int32
	_ = v232
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v241 int32
	_ = v241
	var v243 int32
	_ = v243
	var v245 int64
	_ = v245
	var v246 int64
	_ = v246
	var v247 int64
	_ = v247
	var v248 int64
	_ = v248
	var v249 int64
	_ = v249
	var v250 int64
	_ = v250
	var v251 int64
	_ = v251
	var v253 int64
	_ = v253
	var v255 int64
	_ = v255
	var v257 int64
	_ = v257
	var v259 int64
	_ = v259
	var v261 int64
	_ = v261
	var v263 int32
	_ = v263
	var v264 int32
	_ = v264
	var v265 int32
	_ = v265
	var v268 int32
	_ = v268
	var v271 int32
	_ = v271
	var v272 int32
	_ = v272
	var v273 int32
	_ = v273
	var v275 int32
	_ = v275
	var v278 int32
	_ = v278
	var v279 int32
	_ = v279
	var v283 int32
	_ = v283
	var v287 int32
	_ = v287
	var v289 int32
	_ = v289
	var v296 int32
	_ = v296
	var v297 int32
	_ = v297
	var v298 int32
	_ = v298
	var v303 int32
	_ = v303
	var v308 int32
	_ = v308
	var v312 int32
	_ = v312
	var v316 int32
	_ = v316
	var v321 int32
	_ = v321
	var v328 int32
	_ = v328
	var v333 int32
	_ = v333
	var v344 int32
	_ = v344
	v8 = m.G0
	v10 = v8 - int32(16)
	m.G0 = v10
	v12 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v12&int32(4) == int32(0) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	F__serverAssert(m, int32(_a_F_sentinelUpdateSentinelAddressInAllPrimaries_0), int32(_a_F_sentinelUpdateSentinelAddressInAllPrimaries_1), int32(1184))
	mBase = m.M
	v344 = m.ExcPending
	if v344 != 0 {
		goto L5
	} else {
		goto L89
	}
L2:
	;
	v18 = *(*int32)(unsafe.Add(mBase, _c_F_sentinelUpdateSentinelAddressInAllPrimaries[0]))
	v19 = F_dictGetIterator(m, v18)
	mBase = m.M
	v22 = m.ExcPending
	if v22 != 0 {
		goto L5
	} else {
		goto L6
	}
L3:
	;
	m.G0 = v10 + int32(16)
	return v333
L4:
	;
	v137 = v126
	v138 = int32(0)
	goto L35
L5:
	;
	return int32(0)
L6:
	;
	v30 = v19 + int32(20)
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v19)+16))
	if v31 != 0 {
		goto L9
	} else {
		goto L10
	}
L7:
	;
	if v126 != 0 {
		goto L4
	} else {
		goto L33
	}
L8:
	;
	v37 = v30
	v38 = v34
	goto L11
L9:
	;
	v34 = int32(1)
	goto L8
L10:
	;
	v34 = int32(0)
	goto L8
L11:
	;
	switch v38 {
	case 0:
		goto L16
	default:
		goto L15
	}
L13:
	;
	v38 = int32(0)
	goto L11
L14:
	;
	goto L7
L15:
	;
	v118 = *(*int32)(unsafe.Add(mBase, uint32(v37)))
	*(*int32)(unsafe.Add(mBase, uint32(v19)+16)) = v118
	if v118 == int32(0) {
		goto L13
	} else {
		goto L32
	}
L16:
	;
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v19)+4))
	if v42 != int32(-1) {
		v81 = v42
		goto L17
	} else {
		goto L18
	}
L17:
	;
	v82 = int32(1)
	v83 = v81 + v82
	*(*int32)(unsafe.Add(mBase, uint32(v19)+4)) = v83
	v85 = int32(0)
	v88 = *(*int32)(unsafe.Add(mBase, uint32(v19)))
	v89 = *(*int32)(unsafe.Add(mBase, uint32(v19)+8))
	v93 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v88+v89+int32(26)))))
	if v93 == int32(255) {
		goto L26
	} else {
		goto L27
	}
L18:
	;
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v19)+8))
	if v46 != 0 {
		v81 = int32(-1)
		goto L17
	} else {
		goto L19
	}
L19:
	;
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v19)))
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v19)+12))
	if v48 == int32(0) {
		goto L21
	} else {
		goto L22
	}
L20:
	;
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v74)+20))
	if v75 != int32(-1) {
		goto L23
	} else {
		goto L24
	}
L21:
	;
	v55 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v47)+16)))
	v56 = int64(*(*int8)(unsafe.Add(mBase, uint32(v47)+27)))
	v57 = int64(*(*int32)(unsafe.Add(mBase, uint32(v47)+8)))
	v58 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v47)+12)))
	v59 = int64(*(*int8)(unsafe.Add(mBase, uint32(v47)+26)))
	v60 = int64(*(*int32)(unsafe.Add(mBase, uint32(v47)+4)))
	v61 = F_wangHash64(m, v60)
	mBase = m.M
	v63 = F_wangHash64(m, v59+v61)
	mBase = m.M
	v65 = F_wangHash64(m, v58+v63)
	mBase = m.M
	v67 = F_wangHash64(m, v57+v65)
	mBase = m.M
	v69 = F_wangHash64(m, v56+v67)
	mBase = m.M
	v71 = F_wangHash64(m, v55+v69)
	mBase = m.M
	*(*int64)(unsafe.Add(mBase, uint32(v19)+24)) = v71
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v19)))
	v74 = v73
	goto L20
L22:
	;
	v51 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v47)+24)))
	v53 = v51 + int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v47)+24)) = uint16(v53)
	v74 = v47
	goto L20
L23:
	;
	v81 = v75 + int32(-1)
	goto L17
L24:
	;
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v19)+4))
	v81 = v78
	goto L17
L25:
	;
	v108 = int32(2)
	v113 = *(*int32)(unsafe.Add(mBase, uint32(v88+v106<<(uint(v108)%32)+int32(4))))
	v37 = v113 + v107<<(uint(v108)%32)
	v38 = int32(1)
	goto L11
L26:
	;
	v97 = v85
	goto L28
L27:
	;
	v97 = v82 << (uint(v93) % 32)
	goto L28
L28:
	;
	if v83 < v97 {
		v106 = v89
		v107 = v83
		goto L25
	} else {
		goto L29
	}
L29:
	;
	if v89 != 0 {
		v126 = v85
		goto L14
	} else {
		goto L30
	}
L30:
	;
	v99 = *(*int32)(unsafe.Add(mBase, uint32(v88)+20))
	if v99 == int32(-1) {
		v126 = v85
		goto L14
	} else {
		goto L31
	}
L31:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v19)+4)) = int64(4294967296)
	v106 = int32(1)
	v107 = int32(0)
	goto L25
L32:
	;
	v122 = *(*int32)(unsafe.Add(mBase, uint32(v118)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v30))) = v122
	v126 = v118
	goto L14
L33:
	;
	F_dictReleaseIterator(m, v19)
	mBase = m.M
	v131 = m.ExcPending
	if v131 != 0 {
		goto L5
	} else {
		goto L34
	}
L34:
	;
	v333 = int32(0)
	goto L3
L35:
	;
	v141 = *(*int32)(unsafe.Add(mBase, uint32(v137)+8))
	goto L38
L36:
	;
	F_dictReleaseIterator(m, v19)
	mBase = m.M
	v321 = m.ExcPending
	if v321 != 0 {
		goto L5
	} else {
		goto L85
	}
L37:
	;
	v220 = v19 + int32(20)
	v221 = *(*int32)(unsafe.Add(mBase, uint32(v19)+16))
	if v221 != 0 {
		goto L60
	} else {
		goto L61
	}
L38:
	;
	v142 = *(*int32)(unsafe.Add(mBase, uint32(v141)+144))
	v143 = int32(0)
	v145 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v146 = F_getSentinelValkeyInstanceByAddrAndRunID(m, v142, v143, v143, v145)
	mBase = m.M
	v147 = m.ExcPending
	if v147 != 0 {
		goto L5
	} else {
		goto L39
	}
L39:
	;
	if v146 == int32(0) {
		v211 = v138
		goto L37
	} else {
		goto L40
	}
L40:
	;
	v150 = *(*int32)(unsafe.Add(mBase, uint32(v146)+28))
	v151 = *(*int32)(unsafe.Add(mBase, uint32(v150)+12))
	if v151 == int32(0) {
		v167 = v150
		goto L41
	} else {
		goto L42
	}
L41:
	;
	v168 = *(*int32)(unsafe.Add(mBase, uint32(v167)+16))
	if v168 == int32(0) {
		goto L46
	} else {
		goto L47
	}
L42:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v150)+8)) = int64(0)
	v156 = *(*int32)(unsafe.Add(mBase, uint32(v150)+16))
	if v156 != v151 {
		goto L43
	} else {
		goto L44
	}
L43:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v151)+212)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v150)+4)) = int32(1)
	F_valkeyAsyncFree(m, v151)
	mBase = m.M
	v165 = m.ExcPending
	if v165 != 0 {
		goto L5
	} else {
		goto L45
	}
L44:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v150)+16)) = int32(0)
	goto L43
L45:
	;
	v166 = *(*int32)(unsafe.Add(mBase, uint32(v146)+28))
	v167 = v166
	goto L41
L46:
	;
	if v146 == l0 {
		v211 = v138
		goto L37
	} else {
		goto L51
	}
L47:
	;
	v171 = *(*int32)(unsafe.Add(mBase, uint32(v167)+12))
	if v171 != v168 {
		goto L48
	} else {
		goto L49
	}
L48:
	;
	v175 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v167)+16)) = v175
	*(*int32)(unsafe.Add(mBase, uint32(v168)+212)) = v175
	*(*int32)(unsafe.Add(mBase, uint32(v167)+4)) = int32(1)
	F_valkeyAsyncFree(m, v168)
	mBase = m.M
	v182 = m.ExcPending
	if v182 != 0 {
		goto L5
	} else {
		goto L50
	}
L49:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v167)+8)) = int64(0)
	goto L48
L50:
	;
	goto L46
L51:
	;
	v184 = *(*int32)(unsafe.Add(mBase, uint32(v146)+24))
	v185 = *(*int32)(unsafe.Add(mBase, uint32(v184)))
	F_sdsfree(m, v185)
	mBase = m.M
	v187 = m.ExcPending
	if v187 != 0 {
		goto L5
	} else {
		goto L52
	}
L52:
	;
	v188 = *(*int32)(unsafe.Add(mBase, uint32(v184)+4))
	F_sdsfree(m, v188)
	mBase = m.M
	v190 = m.ExcPending
	if v190 != 0 {
		goto L5
	} else {
		goto L53
	}
L53:
	;
	F_valkey_free(m, v184)
	mBase = m.M
	v192 = m.ExcPending
	if v192 != 0 {
		goto L5
	} else {
		goto L54
	}
L54:
	;
	v193 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v195 = F_valkey_malloc(m, int32(12))
	mBase = m.M
	v196 = m.ExcPending
	if v196 != 0 {
		goto L5
	} else {
		goto L55
	}
L55:
	;
	v197 = *(*int32)(unsafe.Add(mBase, uint32(v193)))
	v198 = F_sdsnew(m, v197)
	mBase = m.M
	v199 = m.ExcPending
	if v199 != 0 {
		goto L5
	} else {
		goto L56
	}
L56:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v195))) = v198
	v201 = *(*int32)(unsafe.Add(mBase, uint32(v193)+4))
	v202 = F_sdsnew(m, v201)
	mBase = m.M
	v203 = m.ExcPending
	if v203 != 0 {
		goto L5
	} else {
		goto L57
	}
L57:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v195)+4)) = v202
	v205 = *(*int32)(unsafe.Add(mBase, uint32(v193)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v195)+8)) = v205
	*(*int32)(unsafe.Add(mBase, uint32(v146)+24)) = v195
	v211 = v138 + int32(1)
	goto L37
L58:
	;
	if v316 != 0 {
		v137 = v316
		v138 = v211
		goto L35
	} else {
		goto L84
	}
L59:
	;
	v227 = v220
	v228 = v224
	goto L62
L60:
	;
	v224 = int32(1)
	goto L59
L61:
	;
	v224 = int32(0)
	goto L59
L62:
	;
	switch v228 {
	case 0:
		goto L67
	default:
		goto L66
	}
L64:
	;
	v228 = int32(0)
	goto L62
L65:
	;
	goto L58
L66:
	;
	v308 = *(*int32)(unsafe.Add(mBase, uint32(v227)))
	*(*int32)(unsafe.Add(mBase, uint32(v19)+16)) = v308
	if v308 == int32(0) {
		goto L64
	} else {
		goto L83
	}
L67:
	;
	v232 = *(*int32)(unsafe.Add(mBase, uint32(v19)+4))
	if v232 != int32(-1) {
		v271 = v232
		goto L68
	} else {
		goto L69
	}
L68:
	;
	v272 = int32(1)
	v273 = v271 + v272
	*(*int32)(unsafe.Add(mBase, uint32(v19)+4)) = v273
	v275 = int32(0)
	v278 = *(*int32)(unsafe.Add(mBase, uint32(v19)))
	v279 = *(*int32)(unsafe.Add(mBase, uint32(v19)+8))
	v283 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v278+v279+int32(26)))))
	if v283 == int32(255) {
		goto L77
	} else {
		goto L78
	}
L69:
	;
	v236 = *(*int32)(unsafe.Add(mBase, uint32(v19)+8))
	if v236 != 0 {
		v271 = int32(-1)
		goto L68
	} else {
		goto L70
	}
L70:
	;
	v237 = *(*int32)(unsafe.Add(mBase, uint32(v19)))
	v238 = *(*int32)(unsafe.Add(mBase, uint32(v19)+12))
	if v238 == int32(0) {
		goto L72
	} else {
		goto L73
	}
L71:
	;
	v265 = *(*int32)(unsafe.Add(mBase, uint32(v264)+20))
	if v265 != int32(-1) {
		goto L74
	} else {
		goto L75
	}
L72:
	;
	v245 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v237)+16)))
	v246 = int64(*(*int8)(unsafe.Add(mBase, uint32(v237)+27)))
	v247 = int64(*(*int32)(unsafe.Add(mBase, uint32(v237)+8)))
	v248 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v237)+12)))
	v249 = int64(*(*int8)(unsafe.Add(mBase, uint32(v237)+26)))
	v250 = int64(*(*int32)(unsafe.Add(mBase, uint32(v237)+4)))
	v251 = F_wangHash64(m, v250)
	mBase = m.M
	v253 = F_wangHash64(m, v249+v251)
	mBase = m.M
	v255 = F_wangHash64(m, v248+v253)
	mBase = m.M
	v257 = F_wangHash64(m, v247+v255)
	mBase = m.M
	v259 = F_wangHash64(m, v246+v257)
	mBase = m.M
	v261 = F_wangHash64(m, v245+v259)
	mBase = m.M
	*(*int64)(unsafe.Add(mBase, uint32(v19)+24)) = v261
	v263 = *(*int32)(unsafe.Add(mBase, uint32(v19)))
	v264 = v263
	goto L71
L73:
	;
	v241 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v237)+24)))
	v243 = v241 + int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v237)+24)) = uint16(v243)
	v264 = v237
	goto L71
L74:
	;
	v271 = v265 + int32(-1)
	goto L68
L75:
	;
	v268 = *(*int32)(unsafe.Add(mBase, uint32(v19)+4))
	v271 = v268
	goto L68
L76:
	;
	v298 = int32(2)
	v303 = *(*int32)(unsafe.Add(mBase, uint32(v278+v296<<(uint(v298)%32)+int32(4))))
	v227 = v303 + v297<<(uint(v298)%32)
	v228 = int32(1)
	goto L62
L77:
	;
	v287 = v275
	goto L79
L78:
	;
	v287 = v272 << (uint(v283) % 32)
	goto L79
L79:
	;
	if v273 < v287 {
		v296 = v279
		v297 = v273
		goto L76
	} else {
		goto L80
	}
L80:
	;
	if v279 != 0 {
		v316 = v275
		goto L65
	} else {
		goto L81
	}
L81:
	;
	v289 = *(*int32)(unsafe.Add(mBase, uint32(v278)+20))
	if v289 == int32(-1) {
		v316 = v275
		goto L65
	} else {
		goto L82
	}
L82:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v19)+4)) = int64(4294967296)
	v296 = int32(1)
	v297 = int32(0)
	goto L76
L83:
	;
	v312 = *(*int32)(unsafe.Add(mBase, uint32(v308)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v220))) = v312
	v316 = v308
	goto L65
L84:
	;
	goto L36
L85:
	;
	if v211 != 0 {
		goto L86
	} else {
		goto L87
	}
L86:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10))) = v211
	F_sentinelEvent(m, int32(2), int32(_a_F_sentinelUpdateSentinelAddressInAllPrimaries_2), l0, int32(_a_F_sentinelUpdateSentinelAddressInAllPrimaries_3), v10)
	mBase = m.M
	v328 = m.ExcPending
	if v328 != 0 {
		goto L5
	} else {
		goto L88
	}
L87:
	;
	v333 = int32(0)
	goto L3
L88:
	;
	v333 = v211
	goto L3
L89:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_sentinelValidateArgs(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v26 int32
	_ = v26
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v61 int32
	_ = v61
	var v65 int32
	_ = v65
	var v71 int32
	_ = v71
	var v79 int32
	_ = v79
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v98 int32
	_ = v98
	v4 = int32(0)
	v10 = m.G0
	v12 = v10 - int32(16)
	m.G0 = v12
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v15 <= l1 {
		v98 = v4
		goto L1
	} else {
		goto L2
	}
L1:
	;
	m.G0 = v12 + int32(16)
	return v98
L2:
	;
	v18 = l1
	goto L3
L3:
	;
	v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v26+v18<<(uint(int32(2))%32))))
	v31 = F_objectGetVal(m, v30)
	mBase = m.M
	v34 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v31+int32(-1)))))
	switch v34 & int32(7) {
	case 0:
		goto L11
	case 1:
		goto L10
	case 2:
		goto L9
	case 3:
		goto L8
	case 4:
		goto L7
	default:
		goto L5
	}
L4:
	;
	v98 = v4
	goto L1
L5:
	;
	v91 = v18 + int32(1)
	v92 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v91 < v92 {
		v18 = v91
		goto L3
	} else {
		goto L21
	}
L6:
	;
	v52 = int32(0)
	if v51 == v52 {
		goto L5
	} else {
		goto L12
	}
L7:
	;
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v31+int32(-17))))
	v51 = v50
	goto L6
L8:
	;
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v31+int32(-9))))
	v51 = v47
	goto L6
L9:
	;
	v44 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v31+int32(-5)))))
	v51 = v44
	goto L6
L10:
	;
	v41 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v31+int32(-3)))))
	v51 = v41
	goto L6
L11:
	;
	v51 = int32(base.Ui32(v34) >> (uint(int32(3)) % 32))
	goto L6
L12:
	;
	v61 = v52
	goto L14
L13:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+4)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(v12))) = v18
	F_addReplyErrorFormat(m, l0, int32(_a_F_sentinelValidateArgs_0), v12)
	mBase = m.M
	v79 = m.ExcPending
	if v79 != 0 {
		goto L19
	} else {
		goto L20
	}
L14:
	;
	v65 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v31+v61))))
	if base.Ui32(v65) < base.Ui32(int32(32)) {
		goto L13
	} else {
		goto L16
	}
L16:
	;
	if v65 == int32(127) {
		goto L13
	} else {
		goto L17
	}
L17:
	;
	v71 = v61 + int32(1)
	if v71 == v51 {
		goto L5
	} else {
		goto L18
	}
L18:
	;
	v61 = v71
	goto L14
L19:
	;
	return int32(0)
L20:
	;
	v98 = int32(-1)
	goto L1
L21:
	;
	goto L4
}
