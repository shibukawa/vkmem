package p0

import (
	base "github.com/shibukawa/vkmem/internal/aot/vkaot/base"
	"unsafe"
)

func F__serverAssertPrintClientInfo(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v38 int32
	_ = v38
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v50 int64
	_ = v50
	var v51 int64
	_ = v51
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v97 int32
	_ = v97
	var v100 int32
	_ = v100
	var v107 int32
	_ = v107
	var v114 int32
	_ = v114
	var v118 int32
	_ = v118
	var v122 int32
	_ = v122
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v130 int32
	_ = v130
	var v137 int32
	_ = v137
	var v140 int32
	_ = v140
	var v143 int32
	_ = v143
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v158 int32
	_ = v158
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v164 int32
	_ = v164
	var v167 int32
	_ = v167
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v173 int32
	_ = v173
	var v179 int32
	_ = v179
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v190 int32
	_ = v190
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
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
	var v215 int32
	_ = v215
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v228 int32
	_ = v228
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v235 int32
	_ = v235
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v241 int32
	_ = v241
	var v245 int32
	_ = v245
	var v247 int32
	_ = v247
	var v248 int32
	_ = v248
	var v249 int32
	_ = v249
	var v251 int32
	_ = v251
	var v252 int32
	_ = v252
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v259 int32
	_ = v259
	var v260 int32
	_ = v260
	var v263 int32
	_ = v263
	var v264 int32
	_ = v264
	var v267 int32
	_ = v267
	var v268 int32
	_ = v268
	var v269 int32
	_ = v269
	var v277 int32
	_ = v277
	var v278 int32
	_ = v278
	v8 = m.G0
	v10 = v8 - int32(144)
	m.G0 = v10
	goto L1
L1:
	;
	v15 = int32(*(*uint8)(unsafe.Add(mBase, _consts[238])))
	if v15 != 0 {
		goto L2
	} else {
		goto L3
	}
L2:
	;
	goto L11
L3:
	;
	v17 = *(*int32)(unsafe.Add(mBase, _consts[15]))
	if int32(3) < v17 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v33 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _consts[238])) = uint8(v33)
	goto L2
L5:
	;
	v23 = *(*int32)(unsafe.Add(mBase, _consts[237]))
	if v23 != 0 {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	v24 = int32(_a518)
	goto L8
L7:
	;
	v24 = int32(_a519)
	goto L8
L8:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+96)) = v24
	F__serverLog(m, int32(1027), int32(_a520), v10+int32(96))
	mBase = m.M
	v31 = m.ExcPending
	if v31 != 0 {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	return
L10:
	;
	goto L4
L11:
	;
	v38 = *(*int32)(unsafe.Add(mBase, _consts[15]))
	if int32(3) < v38 {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	v100 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v100 < int32(1) {
		goto L25
	} else {
		goto L26
	}
L13:
	;
	F__serverLog(m, int32(3), int32(_a521), int32(0))
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L9
	} else {
		goto L14
	}
L14:
	;
	v47 = *(*int32)(unsafe.Add(mBase, _consts[15]))
	if int32(3) < v47 {
		goto L12
	} else {
		goto L15
	}
L15:
	;
	v50 = *(*int64)(unsafe.Add(mBase, uint32(l0)+200))
	v51 = *(*int64)(unsafe.Add(mBase, uint32(l0)+208))
	*(*int64)(unsafe.Add(mBase, uint32(v10)+88)) = v51
	*(*int64)(unsafe.Add(mBase, uint32(v10)+80)) = v50
	F__serverLog(m, int32(3), int32(_a522), v10+int32(80))
	mBase = m.M
	v59 = m.ExcPending
	if v59 != 0 {
		goto L9
	} else {
		goto L16
	}
L16:
	;
	v61 = *(*int32)(unsafe.Add(mBase, _consts[15]))
	if int32(3) < v61 {
		goto L12
	} else {
		goto L17
	}
L17:
	;
	v64 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v64 != 0 {
		goto L19
	} else {
		goto L20
	}
L18:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+64)) = v67
	v75 = F_snprintf(m, v10+int32(112), int32(31), int32(_a523), v10+int32(64))
	mBase = m.M
	v76 = m.ExcPending
	if v76 != 0 {
		goto L9
	} else {
		goto L21
	}
L19:
	;
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v64)+12))
	v67 = v66
	goto L18
L20:
	;
	v67 = int32(-1)
	goto L18
L21:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+48)) = v10 + int32(112)
	F__serverLog(m, int32(3), int32(_a524), v10+int32(48))
	mBase = m.M
	v85 = m.ExcPending
	if v85 != 0 {
		goto L9
	} else {
		goto L22
	}
L22:
	;
	v87 = *(*int32)(unsafe.Add(mBase, _consts[15]))
	if int32(3) < v87 {
		goto L12
	} else {
		goto L23
	}
L23:
	;
	v90 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v10)+32)) = v90
	F__serverLog(m, int32(3), int32(_a525), v10+int32(32))
	mBase = m.M
	v97 = m.ExcPending
	if v97 != 0 {
		goto L9
	} else {
		goto L24
	}
L24:
	;
	goto L12
L25:
	;
	m.G0 = v10 + int32(144)
	return
L26:
	;
	v107 = int32(0)
	goto L27
L27:
	;
	if v107 == int32(0) {
		goto L30
	} else {
		goto L31
	}
L28:
	;
	goto L25
L29:
	;
	v277 = v107 + int32(1)
	v278 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v277 < v278 {
		v107 = v277
		goto L27
	} else {
		goto L72
	}
L30:
	;
	v156 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v158 = v107 << (uint(int32(2)) % 32)
	v160 = *(*int32)(unsafe.Add(mBase, uint32(v156+v158)))
	v161 = F_getArgvReprString(m, v160)
	mBase = m.M
	v162 = m.ExcPending
	if v162 != 0 {
		goto L9
	} else {
		goto L41
	}
L31:
	;
	v114 = *(*int32)(unsafe.Add(mBase, _consts[151]))
	if v114 == int32(0) {
		goto L30
	} else {
		goto L32
	}
L32:
	;
	v118 = *(*int32)(unsafe.Add(mBase, _consts[15]))
	if int32(3) < v118 {
		goto L29
	} else {
		goto L33
	}
L33:
	;
	v122 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v126 = *(*int32)(unsafe.Add(mBase, uint32(v122+v107<<(uint(int32(2))%32))))
	v127 = F_objectGetVal(m, v126)
	mBase = m.M
	v130 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v127+int32(-1)))))
	switch v130 & int32(7) {
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
		v147 = int32(0)
		goto L34
	}
L34:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+20)) = v147
	*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = v107
	F__serverLog(m, int32(3), int32(_a526), v10+int32(16))
	mBase = m.M
	v155 = m.ExcPending
	if v155 != 0 {
		goto L9
	} else {
		goto L40
	}
L35:
	;
	v146 = *(*int32)(unsafe.Add(mBase, uint32(v127+int32(-17))))
	v147 = v146
	goto L34
L36:
	;
	v143 = *(*int32)(unsafe.Add(mBase, uint32(v127+int32(-9))))
	v147 = v143
	goto L34
L37:
	;
	v140 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v127+int32(-5)))))
	v147 = v140
	goto L34
L38:
	;
	v137 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v127+int32(-3)))))
	v147 = v137
	goto L34
L39:
	;
	v147 = int32(base.Ui32(v130) >> (uint(int32(3)) % 32))
	goto L34
L40:
	;
	goto L29
L41:
	;
	v164 = *(*int32)(unsafe.Add(mBase, _consts[15]))
	if int32(3) < v164 {
		goto L42
	} else {
		goto L43
	}
L42:
	;
	F_sdsfree(m, v161)
	mBase = m.M
	v182 = m.ExcPending
	if v182 != 0 {
		goto L9
	} else {
		goto L45
	}
L43:
	;
	v167 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v169 = *(*int32)(unsafe.Add(mBase, uint32(v167+v158)))
	v170 = *(*int32)(unsafe.Add(mBase, uint32(v169)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v10)+4)) = v161
	*(*int32)(unsafe.Add(mBase, uint32(v10))) = v107
	v173 = int32(3)
	*(*int32)(unsafe.Add(mBase, uint32(v10)+8)) = int32(base.Ui32(v170) >> (uint(v173) % 32))
	F__serverLog(m, v173, int32(_a527), v10)
	mBase = m.M
	v179 = m.ExcPending
	if v179 != 0 {
		goto L9
	} else {
		goto L44
	}
L44:
	;
	goto L42
L45:
	;
	v183 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v185 = *(*int32)(unsafe.Add(mBase, uint32(v183+v158)))
	v186 = F_objectGetVal(m, v185)
	mBase = m.M
	v187 = int32(_a139)
	v190 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v186))))
	if v190 != 0 {
		goto L48
	} else {
		goto L49
	}
L46:
	;
	if v222-v224 == int32(0) {
		goto L25
	} else {
		goto L58
	}
L47:
	;
	v222 = F_tolower(m, v218)
	mBase = m.M
	v223 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v219))))
	v224 = F_tolower(m, v223)
	mBase = m.M
	goto L46
L48:
	;
	v192 = v186
	v193 = v187
	v194 = v190
	goto L51
L49:
	;
	v218 = int32(0)
	v219 = v187
	goto L47
L50:
	;
	v218 = v215 & int32(255)
	v219 = v214
	goto L47
L51:
	;
	v196 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v193))))
	if v196 == int32(0) {
		v214 = v193
		v215 = v194
		goto L50
	} else {
		goto L53
	}
L52:
	;
	v214 = v208
	v215 = int32(0)
	goto L50
L53:
	;
	v200 = v194 & int32(255)
	if v200 == v196 {
		goto L54
	} else {
		goto L55
	}
L54:
	;
	v207 = int32(1)
	v208 = v193 + v207
	v209 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v192)+1)))
	if v209 != 0 {
		v192 = v192 + v207
		v193 = v208
		v194 = v209
		goto L51
	} else {
		goto L57
	}
L55:
	;
	v202 = F_tolower(m, v200)
	mBase = m.M
	v203 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v193))))
	v204 = F_tolower(m, v203)
	mBase = m.M
	if v202 == v204 {
		goto L54
	} else {
		goto L56
	}
L56:
	;
	v206 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v192))))
	v214 = v193
	v215 = v206
	goto L50
L57:
	;
	goto L52
L58:
	;
	v228 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v230 = *(*int32)(unsafe.Add(mBase, uint32(v228+v158)))
	v231 = F_objectGetVal(m, v230)
	mBase = m.M
	v232 = int32(_a140)
	v235 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v231))))
	if v235 != 0 {
		goto L61
	} else {
		goto L62
	}
L59:
	;
	if v267-v269 == int32(0) {
		goto L25
	} else {
		goto L71
	}
L60:
	;
	v267 = F_tolower(m, v263)
	mBase = m.M
	v268 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v264))))
	v269 = F_tolower(m, v268)
	mBase = m.M
	goto L59
L61:
	;
	v237 = v231
	v238 = v232
	v239 = v235
	goto L64
L62:
	;
	v263 = int32(0)
	v264 = v232
	goto L60
L63:
	;
	v263 = v260 & int32(255)
	v264 = v259
	goto L60
L64:
	;
	v241 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v238))))
	if v241 == int32(0) {
		v259 = v238
		v260 = v239
		goto L63
	} else {
		goto L66
	}
L65:
	;
	v259 = v253
	v260 = int32(0)
	goto L63
L66:
	;
	v245 = v239 & int32(255)
	if v245 == v241 {
		goto L67
	} else {
		goto L68
	}
L67:
	;
	v252 = int32(1)
	v253 = v238 + v252
	v254 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v237)+1)))
	if v254 != 0 {
		v237 = v237 + v252
		v238 = v253
		v239 = v254
		goto L64
	} else {
		goto L70
	}
L68:
	;
	v247 = F_tolower(m, v245)
	mBase = m.M
	v248 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v238))))
	v249 = F_tolower(m, v248)
	mBase = m.M
	if v247 == v249 {
		goto L67
	} else {
		goto L69
	}
L69:
	;
	v251 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v237))))
	v259 = v238
	v260 = v251
	goto L63
L70:
	;
	goto L65
L71:
	;
	goto L29
L72:
	;
	goto L28
}
func F__serverPanic_2(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v9 = m.G3
	v14 = m.G397
	v15 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
	v16 = F_fwrite(m, v9+int32(_a2231), int32(48), int32(1), v15)
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		return
	} else {
		v22 = F_fwrite(m, v9+int32(_a2232), int32(57), int32(1), v15)
		mBase = m.M
		v23 = m.ExcPending
		if v23 != 0 {
			return
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v7)+8)) = l0
			*(*int32)(unsafe.Add(mBase, uint32(v7)+4)) = v9 + int32(_a2211)
			*(*int32)(unsafe.Add(mBase, uint32(v7))) = v9 + int32(_a2233)
			v33 = F_fiprintf(m, v15, v9+int32(_a2234), v7)
			mBase = m.M
			v34 = m.ExcPending
			if v34 != 0 {
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
func F_initServerClientMemUsageBuckets(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
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
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
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
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
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
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
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
	var v120 int32
	_ = v120
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v141 int32
	_ = v141
	v4 = *(*int32)(unsafe.Add(mBase, _consts[191]))
	if v4 != 0 {
		return
	} else {
		v7 = F_valkey_malloc(m, int32(152))
		mBase = m.M
		v8 = m.ExcPending
		if v8 != 0 {
			return
		} else {
			*(*int32)(unsafe.Add(mBase, _consts[191])) = v7
			*(*int32)(unsafe.Add(mBase, uint32(v7)+4)) = int32(0)
			v12 = F_listCreate(m)
			mBase = m.M
			v13 = m.ExcPending
			if v13 != 0 {
				return
			} else {
				v14 = int32(0)
				v15 = *(*int32)(unsafe.Add(mBase, _consts[191]))
				*(*int32)(unsafe.Add(mBase, uint32(v15)+12)) = v14
				*(*int32)(unsafe.Add(mBase, uint32(v15))) = v12
				v19 = F_listCreate(m)
				mBase = m.M
				v20 = m.ExcPending
				if v20 != 0 {
					return
				} else {
					v21 = int32(0)
					v22 = *(*int32)(unsafe.Add(mBase, _consts[191]))
					*(*int32)(unsafe.Add(mBase, uint32(v22)+20)) = v21
					*(*int32)(unsafe.Add(mBase, uint32(v22)+8)) = v19
					v26 = F_listCreate(m)
					mBase = m.M
					v27 = m.ExcPending
					if v27 != 0 {
						return
					} else {
						v28 = int32(0)
						v29 = *(*int32)(unsafe.Add(mBase, _consts[191]))
						*(*int32)(unsafe.Add(mBase, uint32(v29)+28)) = v28
						*(*int32)(unsafe.Add(mBase, uint32(v29)+16)) = v26
						v33 = F_listCreate(m)
						mBase = m.M
						v34 = m.ExcPending
						if v34 != 0 {
							return
						} else {
							v35 = int32(0)
							v36 = *(*int32)(unsafe.Add(mBase, _consts[191]))
							*(*int32)(unsafe.Add(mBase, uint32(v36)+36)) = v35
							*(*int32)(unsafe.Add(mBase, uint32(v36)+24)) = v33
							v40 = F_listCreate(m)
							mBase = m.M
							v41 = m.ExcPending
							if v41 != 0 {
								return
							} else {
								v42 = int32(0)
								v43 = *(*int32)(unsafe.Add(mBase, _consts[191]))
								*(*int32)(unsafe.Add(mBase, uint32(v43)+44)) = v42
								*(*int32)(unsafe.Add(mBase, uint32(v43)+32)) = v40
								v47 = F_listCreate(m)
								mBase = m.M
								v48 = m.ExcPending
								if v48 != 0 {
									return
								} else {
									v49 = int32(0)
									v50 = *(*int32)(unsafe.Add(mBase, _consts[191]))
									*(*int32)(unsafe.Add(mBase, uint32(v50)+52)) = v49
									*(*int32)(unsafe.Add(mBase, uint32(v50)+40)) = v47
									v54 = F_listCreate(m)
									mBase = m.M
									v55 = m.ExcPending
									if v55 != 0 {
										return
									} else {
										v56 = int32(0)
										v57 = *(*int32)(unsafe.Add(mBase, _consts[191]))
										*(*int32)(unsafe.Add(mBase, uint32(v57)+60)) = v56
										*(*int32)(unsafe.Add(mBase, uint32(v57)+48)) = v54
										v61 = F_listCreate(m)
										mBase = m.M
										v62 = m.ExcPending
										if v62 != 0 {
											return
										} else {
											v63 = int32(0)
											v64 = *(*int32)(unsafe.Add(mBase, _consts[191]))
											*(*int32)(unsafe.Add(mBase, uint32(v64)+68)) = v63
											*(*int32)(unsafe.Add(mBase, uint32(v64)+56)) = v61
											v68 = F_listCreate(m)
											mBase = m.M
											v69 = m.ExcPending
											if v69 != 0 {
												return
											} else {
												v70 = int32(0)
												v71 = *(*int32)(unsafe.Add(mBase, _consts[191]))
												*(*int32)(unsafe.Add(mBase, uint32(v71)+76)) = v70
												*(*int32)(unsafe.Add(mBase, uint32(v71)+64)) = v68
												v75 = F_listCreate(m)
												mBase = m.M
												v76 = m.ExcPending
												if v76 != 0 {
													return
												} else {
													v77 = int32(0)
													v78 = *(*int32)(unsafe.Add(mBase, _consts[191]))
													*(*int32)(unsafe.Add(mBase, uint32(v78)+84)) = v77
													*(*int32)(unsafe.Add(mBase, uint32(v78)+72)) = v75
													v82 = F_listCreate(m)
													mBase = m.M
													v83 = m.ExcPending
													if v83 != 0 {
														return
													} else {
														v84 = int32(0)
														v85 = *(*int32)(unsafe.Add(mBase, _consts[191]))
														*(*int32)(unsafe.Add(mBase, uint32(v85)+92)) = v84
														*(*int32)(unsafe.Add(mBase, uint32(v85)+80)) = v82
														v89 = F_listCreate(m)
														mBase = m.M
														v90 = m.ExcPending
														if v90 != 0 {
															return
														} else {
															v91 = int32(0)
															v92 = *(*int32)(unsafe.Add(mBase, _consts[191]))
															*(*int32)(unsafe.Add(mBase, uint32(v92)+100)) = v91
															*(*int32)(unsafe.Add(mBase, uint32(v92)+88)) = v89
															v96 = F_listCreate(m)
															mBase = m.M
															v97 = m.ExcPending
															if v97 != 0 {
																return
															} else {
																v98 = int32(0)
																v99 = *(*int32)(unsafe.Add(mBase, _consts[191]))
																*(*int32)(unsafe.Add(mBase, uint32(v99)+108)) = v98
																*(*int32)(unsafe.Add(mBase, uint32(v99)+96)) = v96
																v103 = F_listCreate(m)
																mBase = m.M
																v104 = m.ExcPending
																if v104 != 0 {
																	return
																} else {
																	v105 = int32(0)
																	v106 = *(*int32)(unsafe.Add(mBase, _consts[191]))
																	*(*int32)(unsafe.Add(mBase, uint32(v106)+116)) = v105
																	*(*int32)(unsafe.Add(mBase, uint32(v106)+104)) = v103
																	v110 = F_listCreate(m)
																	mBase = m.M
																	v111 = m.ExcPending
																	if v111 != 0 {
																		return
																	} else {
																		v112 = int32(0)
																		v113 = *(*int32)(unsafe.Add(mBase, _consts[191]))
																		*(*int32)(unsafe.Add(mBase, uint32(v113)+124)) = v112
																		*(*int32)(unsafe.Add(mBase, uint32(v113)+112)) = v110
																		v117 = F_listCreate(m)
																		mBase = m.M
																		v118 = m.ExcPending
																		if v118 != 0 {
																			return
																		} else {
																			v119 = int32(0)
																			v120 = *(*int32)(unsafe.Add(mBase, _consts[191]))
																			*(*int32)(unsafe.Add(mBase, uint32(v120)+132)) = v119
																			*(*int32)(unsafe.Add(mBase, uint32(v120)+120)) = v117
																			v124 = F_listCreate(m)
																			mBase = m.M
																			v125 = m.ExcPending
																			if v125 != 0 {
																				return
																			} else {
																				v126 = int32(0)
																				v127 = *(*int32)(unsafe.Add(mBase, _consts[191]))
																				*(*int32)(unsafe.Add(mBase, uint32(v127)+140)) = v126
																				*(*int32)(unsafe.Add(mBase, uint32(v127)+128)) = v124
																				v131 = F_listCreate(m)
																				mBase = m.M
																				v132 = m.ExcPending
																				if v132 != 0 {
																					return
																				} else {
																					v133 = int32(0)
																					v134 = *(*int32)(unsafe.Add(mBase, _consts[191]))
																					*(*int32)(unsafe.Add(mBase, uint32(v134)+148)) = v133
																					*(*int32)(unsafe.Add(mBase, uint32(v134)+136)) = v131
																					v138 = F_listCreate(m)
																					mBase = m.M
																					v139 = m.ExcPending
																					if v139 != 0 {
																						return
																					} else {
																						v141 = *(*int32)(unsafe.Add(mBase, _consts[191]))
																						*(*int32)(unsafe.Add(mBase, uint32(v141)+144)) = v138
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
	}
}
func F_initServerConfig(m *base.Module) {
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
	var v12 int64
	_ = v12
	var v15 int64
	_ = v15
	var v16 int64
	_ = v16
	var v20 int64
	_ = v20
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	var v32 int64
	_ = v32
	var v36 int64
	_ = v36
	var v41 int32
	_ = v41
	var v42 int64
	_ = v42
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v54 int64
	_ = v54
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int64
	_ = v66
	var v72 int32
	_ = v72
	var v94 int32
	_ = v94
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v130 int64
	_ = v130
	var v154 int64
	_ = v154
	var v159 int64
	_ = v159
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v231 int32
	_ = v231
	var v251 int32
	_ = v251
	var v252 int32
	_ = v252
	var v260 int32
	_ = v260
	var v261 int64
	_ = v261
	var v276 int32
	_ = v276
	var v280 int32
	_ = v280
	var v284 int32
	_ = v284
	var v288 int32
	_ = v288
	var v289 int32
	_ = v289
	var v290 int64
	_ = v290
	var v296 int64
	_ = v296
	var v332 int64
	_ = v332
	var v349 int32
	_ = v349
	var v351 int32
	_ = v351
	var v367 int32
	_ = v367
	var v371 int64
	_ = v371
	var v375 int32
	_ = v375
	var v376 int32
	_ = v376
	var v379 int32
	_ = v379
	var v380 int32
	_ = v380
	var v381 int32
	_ = v381
	var v387 int32
	_ = v387
	var v388 int32
	_ = v388
	v5 = m.G0
	v7 = v5 - int32(64)
	m.G0 = v7
	F_initConfigValues(m)
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		return
	} else {
		v11 = int32(0)
		v12 = F_ustime(m)
		mBase = m.M
		*(*int64)(unsafe.Add(mBase, _consts[277])) = v12
		v15 = int64(1000)
		v16 = base.I64_div_s(v12, v15)
		*(*int64)(unsafe.Add(mBase, _consts[32])) = v16
		v20 = base.I64_div_s(v12, int64(1000000))
		*(*int64)(unsafe.Add(mBase, _consts[37])) = v20
		v23 = *(*int32)(unsafe.Add(mBase, _consts[167]))
		v27 = int32(base.Ui32(v23&int32(2)) >> (uint(int32(1)) % 32))
		*(*uint8)(unsafe.Add(mBase, _consts[334])) = uint8(v27)
		v32 = base.I64_div_s(v16, int64(60000))
		*(*uint16)(unsafe.Add(mBase, _consts[335])) = uint16(v32)
		v36 = base.I64_div_s(v16, v15)
		*(*int32)(unsafe.Add(mBase, _consts[333])) = base.I32_wrap_i64(v36) & int32(16777215)
		v41 = int32(0)
		v42 = *(*int64)(unsafe.Add(mBase, _consts[37]))
		*(*int64)(unsafe.Add(mBase, uint32(v7)+8)) = v42
		v48 = F___localtime_r(m, v7+int32(8), v7+int32(20))
		mBase = m.M
		v50 = *(*int32)(unsafe.Add(mBase, uint32(v7)+52))
		*(*int32)(unsafe.Add(mBase, _consts[685])) = v50
		v54 = *(*int64)(unsafe.Add(mBase, _consts[32]))
		*(*int64)(unsafe.Add(mBase, _consts[78])) = v54
		F_getRandomHexChars(m, int32(_a1335), int32(40))
		mBase = m.M
		v59 = m.ExcPending
		if v59 != 0 {
			return
		} else {
			v60 = int32(0)
			*(*uint8)(unsafe.Add(mBase, _consts[781])) = uint8(v60)
			F_changeReplicationId(m)
			mBase = m.M
			v64 = m.ExcPending
			if v64 != 0 {
				return
			} else {
				v65 = int32(_a69)
				v66 = int64(3472328296227680304)
				*(*int64)(unsafe.Add(mBase, _consts[512])) = v66
				*(*int64)(unsafe.Add(mBase, _consts[508])) = int64(-1)
				v72 = int32(0)
				*(*uint8)(unsafe.Add(mBase, _consts[513])) = uint8(v72)
				*(*int64)(unsafe.Add(mBase, _consts[514])) = v66
				*(*int64)(unsafe.Add(mBase, _consts[515])) = v66
				*(*int64)(unsafe.Add(mBase, _consts[516])) = v66
				*(*int64)(unsafe.Add(mBase, _consts[517])) = v66
				v94 = int32(0)
				*(*int32)(unsafe.Add(mBase, _consts[149])) = int32(10)
				v98 = *(*int32)(unsafe.Add(mBase, _consts[782]))
				v99 = int32(0)
				*(*int32)(unsafe.Add(mBase, _consts[192])) = int32(2)
				*(*int32)(unsafe.Add(mBase, _consts[783])) = int32(32)
				*(*int64)(unsafe.Add(mBase, _consts[178])) = int64(0)
				*(*int64)(unsafe.Add(mBase, _consts[784])) = base.I64_extend_i32_s(v98)
				v113 = F_zstrdup(m, int32(_a1180))
				mBase = m.M
				v114 = m.ExcPending
				if v114 != 0 {
					return
				} else {
					*(*int32)(unsafe.Add(mBase, _consts[785])) = v113
					v118 = F_zstrdup(m, int32(_a1336))
					mBase = m.M
					v119 = m.ExcPending
					if v119 != 0 {
						return
					} else {
						*(*int32)(unsafe.Add(mBase, _consts[786])) = v118
						v125 = F__emscripten_memset_bulkmem(m, int32(_a1337), base.I32_extend8_s(int32(0)), int32(352))
						mBase = m.M
						v126 = int32(0)
						*(*int32)(unsafe.Add(mBase, _consts[787])) = int32(1)
						v130 = int64(0)
						*(*int64)(unsafe.Add(mBase, _consts[788])) = v130
						*(*int64)(unsafe.Add(mBase, _consts[485])) = v130
						*(*int64)(unsafe.Add(mBase, _consts[789])) = v130
						*(*int32)(unsafe.Add(mBase, _consts[116])) = v126
						*(*int32)(unsafe.Add(mBase, _consts[790])) = v126
						*(*int32)(unsafe.Add(mBase, _consts[217])) = v126
						*(*int32)(unsafe.Add(mBase, _consts[27])) = v126
						*(*int32)(unsafe.Add(mBase, _consts[482])) = v126
						v154 = F___time(m, v126)
						mBase = m.M
						*(*int64)(unsafe.Add(mBase, _consts[791])) = v130
						v159 = int64(-1)
						*(*int64)(unsafe.Add(mBase, _consts[792])) = v159
						*(*int64)(unsafe.Add(mBase, _consts[40])) = v159
						*(*int64)(unsafe.Add(mBase, _consts[793])) = v130
						*(*int64)(unsafe.Add(mBase, _consts[794])) = v130
						*(*int64)(unsafe.Add(mBase, _consts[28])) = v159
						*(*int64)(unsafe.Add(mBase, _consts[29])) = v130
						*(*int64)(unsafe.Add(mBase, _consts[30])) = v130
						*(*int64)(unsafe.Add(mBase, _consts[795])) = v130
						*(*int64)(unsafe.Add(mBase, _consts[796])) = v130
						*(*int64)(unsafe.Add(mBase, _consts[67])) = v130
						*(*int64)(unsafe.Add(mBase, _consts[797])) = v130
						*(*int64)(unsafe.Add(mBase, _consts[33])) = v154 * int64(1000)
						*(*int32)(unsafe.Add(mBase, _consts[43])) = v126
						*(*int32)(unsafe.Add(mBase, _consts[439])) = v126
						*(*int32)(unsafe.Add(mBase, _consts[798])) = v126
						*(*int32)(unsafe.Add(mBase, _consts[799])) = v126
						*(*int32)(unsafe.Add(mBase, _consts[800])) = v126
						*(*int32)(unsafe.Add(mBase, _consts[686])) = v126
						*(*int64)(unsafe.Add(mBase, _consts[801])) = v130
						*(*int32)(unsafe.Add(mBase, _consts[802])) = v126
						*(*int32)(unsafe.Add(mBase, _consts[141])) = v126
						v223 = F_dictCreate(m, int32(_a1338))
						mBase = m.M
						v224 = m.ExcPending
						if v224 != 0 {
							return
						} else {
							v225 = int32(0)
							*(*int64)(unsafe.Add(mBase, _consts[803])) = int64(1)
							*(*int32)(unsafe.Add(mBase, _consts[804])) = v223
							v231 = F_sysconf(m, int32(30))
							mBase = m.M
							*(*int64)(unsafe.Add(mBase, _consts[237])) = int64(0)
							*(*int32)(unsafe.Add(mBase, _consts[805])) = v231
							*(*int32)(unsafe.Add(mBase, _consts[806])) = int32(1)
							*(*int32)(unsafe.Add(mBase, _consts[180])) = int32(3)
							*(*int32)(unsafe.Add(mBase, _consts[807])) = v225
							*(*int32)(unsafe.Add(mBase, _consts[218])) = v225
							v251 = F_valkey_malloc(m, int32(24))
							mBase = m.M
							v252 = m.ExcPending
							if v252 != 0 {
								return
							} else {
								*(*int32)(unsafe.Add(mBase, _consts[181])) = v251
								*(*int64)(unsafe.Add(mBase, uint32(v251)+16)) = int64(4636730254480218522)
								*(*int64)(unsafe.Add(mBase, uint32(v251)+8)) = int64(4636666922610458624)
								*(*int64)(unsafe.Add(mBase, uint32(v251))) = int64(4632233691727265792)
								v260 = int32(0)
								v261 = int64(0)
								*(*int64)(unsafe.Add(mBase, _consts[808])) = v261
								*(*int64)(unsafe.Add(mBase, _consts[809])) = v261
								*(*int64)(unsafe.Add(mBase, _consts[810])) = v261
								*(*int64)(unsafe.Add(mBase, _consts[811])) = v261
								*(*int32)(unsafe.Add(mBase, _consts[812])) = v260
								F_resetServerSaveParams(m)
								mBase = m.M
								v276 = m.ExcPending
								if v276 != 0 {
									return
								} else {
									F_appendServerSaveParams(m, int64(3600), int32(1))
									mBase = m.M
									v280 = m.ExcPending
									if v280 != 0 {
										return
									} else {
										F_appendServerSaveParams(m, int64(300), int32(100))
										mBase = m.M
										v284 = m.ExcPending
										if v284 != 0 {
											return
										} else {
											F_appendServerSaveParams(m, int64(60), int32(10000))
											mBase = m.M
											v288 = m.ExcPending
											if v288 != 0 {
												return
											} else {
												v289 = int32(0)
												v290 = int64(-1)
												*(*int64)(unsafe.Add(mBase, _consts[574])) = v290
												*(*int64)(unsafe.Add(mBase, _consts[64])) = int64(27397596381184)
												v296 = int64(0)
												*(*int64)(unsafe.Add(mBase, _consts[518])) = v296
												*(*int64)(unsafe.Add(mBase, _consts[539])) = int64(4294967295)
												*(*int64)(unsafe.Add(mBase, _consts[519])) = v296
												*(*int64)(unsafe.Add(mBase, _consts[586])) = v290
												*(*int64)(unsafe.Add(mBase, _consts[813])) = v296
												*(*int64)(unsafe.Add(mBase, _consts[31])) = v296
												*(*int64)(unsafe.Add(mBase, _consts[569])) = int64(21474836480)
												*(*int64)(unsafe.Add(mBase, _consts[814])) = int64(100)
												*(*int32)(unsafe.Add(mBase, _consts[134])) = v289
												*(*int32)(unsafe.Add(mBase, _consts[133])) = v289
												*(*int32)(unsafe.Add(mBase, _consts[521])) = v289
												*(*int32)(unsafe.Add(mBase, _consts[448])) = v289
												v332 = F___time(m, v289)
												mBase = m.M
												*(*int64)(unsafe.Add(mBase, _consts[600])) = v296
												*(*int64)(unsafe.Add(mBase, _consts[603])) = v332
												*(*int64)(unsafe.Add(mBase, _consts[815])) = v296
												*(*int64)(unsafe.Add(mBase, _consts[816])) = v296
												v349 = F__emscripten_memcpy_bulkmem(m, int32(_a1339), int32(_a1340), int32(72))
												mBase = m.M
												v351 = int32(0)
												*(*int64)(unsafe.Add(mBase, _consts[817])) = int64(0)
												*(*int64)(unsafe.Add(mBase, _consts[479])) = int64(9218868437227405312)
												*(*int64)(unsafe.Add(mBase, _consts[480])) = int64(-4503599627370496)
												*(*int64)(unsafe.Add(mBase, _consts[478])) = int64(9221120237041090560)
												v367 = *(*int32)(unsafe.Add(mBase, _consts[818]))
												*(*int32)(unsafe.Add(mBase, _consts[819])) = v367
												v371 = *(*int64)(unsafe.Add(mBase, _consts[820]))
												*(*int64)(unsafe.Add(mBase, _consts[821])) = v371
												v375 = F_hashtableCreate(m, int32(_a1341))
												mBase = m.M
												v376 = m.ExcPending
												if v376 != 0 {
													return
												} else {
													*(*int32)(unsafe.Add(mBase, _consts[822])) = v375
													v379 = F_hashtableCreate(m, int32(_a1342))
													mBase = m.M
													v380 = m.ExcPending
													if v380 != 0 {
														return
													} else {
														v381 = int32(0)
														*(*int64)(unsafe.Add(mBase, _consts[823])) = int64(0)
														*(*int32)(unsafe.Add(mBase, _consts[8])) = v379
														F_populateCommandTable(m)
														mBase = m.M
														v387 = m.ExcPending
														if v387 != 0 {
															return
														} else {
															v388 = int32(0)
															*(*int32)(unsafe.Add(mBase, _consts[241])) = v388
															m.G0 = v7 + int32(64)
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
func F_loadServerConfig(m *base.Module, l0 int32, l1 int32, l2 int32) {
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
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v77 int32
	_ = v77
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v104 int32
	_ = v104
	var v107 int32
	_ = v107
	var v109 int32
	_ = v109
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v123 int32
	_ = v123
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v133 int32
	_ = v133
	var v138 int32
	_ = v138
	var v141 int32
	_ = v141
	var v145 int32
	_ = v145
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v156 int32
	_ = v156
	var v161 int32
	_ = v161
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v175 int32
	_ = v175
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v190 int32
	_ = v190
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v195 int32
	_ = v195
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v206 int32
	_ = v206
	var v213 int32
	_ = v213
	var v218 int32
	_ = v218
	var v225 int32
	_ = v225
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v239 int32
	_ = v239
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v248 int32
	_ = v248
	var v249 int32
	_ = v249
	var v254 int32
	_ = v254
	var v259 int32
	_ = v259
	var v260 int32
	_ = v260
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
	v7 = m.G0
	v9 = v7 - int32(1104)
	m.G0 = v9
	v11 = F_sdsempty(m)
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	if l0 == int32(0) {
		v213 = v11
		goto L3
	} else {
		goto L4
	}
L3:
	;
	if l1 == int32(0) {
		v254 = v213
		goto L67
	} else {
		goto L68
	}
L4:
	;
	v15 = int32(42)
	v16 = F___strchrnul(m, l0, v15)
	mBase = m.M
	v18 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16))))
	if v18 == v15 {
		goto L8
	} else {
		goto L9
	}
L5:
	;
	v161 = F_fopen(m, l0, int32(_a178))
	mBase = m.M
	if v161 == int32(0) {
		goto L51
	} else {
		goto L52
	}
L6:
	;
	v41 = int32(0)
	v45 = F_glob(m, l0, v41, v41, v9+int32(28))
	mBase = m.M
	v46 = m.ExcPending
	if v46 != 0 {
		goto L1
	} else {
		goto L22
	}
L7:
	;
	if v22 != 0 {
		goto L6
	} else {
		goto L11
	}
L8:
	;
	v22 = v16
	goto L10
L9:
	;
	v22 = int32(0)
	goto L10
L10:
	;
	goto L7
L11:
	;
	v23 = int32(63)
	v24 = F___strchrnul(m, l0, v23)
	mBase = m.M
	v26 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24))))
	if v26 == v23 {
		goto L13
	} else {
		goto L14
	}
L12:
	;
	if v30 != 0 {
		goto L6
	} else {
		goto L16
	}
L13:
	;
	v30 = v24
	goto L15
L14:
	;
	v30 = int32(0)
	goto L15
L15:
	;
	goto L12
L16:
	;
	v31 = int32(91)
	v32 = F___strchrnul(m, l0, v31)
	mBase = m.M
	v34 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v32))))
	if v34 == v31 {
		goto L18
	} else {
		goto L19
	}
L17:
	;
	if v38 == int32(0) {
		goto L5
	} else {
		goto L21
	}
L18:
	;
	v38 = v32
	goto L20
L19:
	;
	v38 = int32(0)
	goto L20
L20:
	;
	goto L17
L21:
	;
	goto L6
L22:
	;
	if v45 != 0 {
		v213 = v11
		goto L3
	} else {
		goto L23
	}
L23:
	;
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v9)+28))
	if v47 == int32(0) {
		v104 = v11
		goto L25
	} else {
		goto L26
	}
L24:
	;
	v138 = *(*int32)(unsafe.Add(mBase, _consts[15]))
	if int32(3) < v138 {
		goto L46
	} else {
		goto L47
	}
L25:
	;
	v107 = v9 + int32(28)
	v109 = *(*int32)(unsafe.Add(mBase, uint32(v107)))
	if v109 == int32(0) {
		goto L41
	} else {
		goto L42
	}
L26:
	;
	v55 = v11
	v56 = int32(0)
	goto L27
L27:
	;
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v9)+32))
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v57+v56<<(uint(int32(2))%32))))
	v63 = F_fopen(m, v61, int32(_a178))
	mBase = m.M
	if v63 == int32(0) {
		goto L24
	} else {
		goto L29
	}
L28:
	;
	v104 = v92
	goto L25
L29:
	;
	v69 = F_fgets(m, v9+int32(64), int32(1025), v63)
	mBase = m.M
	v70 = m.ExcPending
	if v70 != 0 {
		goto L1
	} else {
		goto L31
	}
L30:
	;
	v94 = F_fclose(m, v63)
	mBase = m.M
	v95 = m.ExcPending
	if v95 != 0 {
		goto L1
	} else {
		goto L38
	}
L31:
	;
	if v69 == int32(0) {
		v92 = v55
		goto L30
	} else {
		goto L32
	}
L32:
	;
	v77 = v55
	goto L33
L33:
	;
	v81 = F_sdscat(m, v77, v9+int32(64))
	mBase = m.M
	v82 = m.ExcPending
	if v82 != 0 {
		goto L1
	} else {
		goto L35
	}
L34:
	;
	v92 = v81
	goto L30
L35:
	;
	v86 = F_fgets(m, v9+int32(64), int32(1025), v63)
	mBase = m.M
	v87 = m.ExcPending
	if v87 != 0 {
		goto L1
	} else {
		goto L36
	}
L36:
	;
	if v86 != 0 {
		v77 = v81
		goto L33
	} else {
		goto L37
	}
L37:
	;
	goto L34
L38:
	;
	v97 = v56 + int32(1)
	v98 = *(*int32)(unsafe.Add(mBase, uint32(v9)+28))
	if base.Ui32(v97) < base.Ui32(v98) {
		v55 = v92
		v56 = v97
		goto L27
	} else {
		goto L39
	}
L39:
	;
	goto L28
L40:
	;
	v213 = v104
	goto L3
L41:
	;
	v133 = *(*int32)(unsafe.Add(mBase, uint32(v107)+4))
	F_emscripten_builtin_free(m, v133)
	mBase = m.M
	*(*int64)(unsafe.Add(mBase, uint32(v107))) = int64(0)
	goto L40
L42:
	;
	v114 = int32(0)
	goto L43
L43:
	;
	v115 = *(*int32)(unsafe.Add(mBase, uint32(v107)+4))
	v116 = *(*int32)(unsafe.Add(mBase, uint32(v107)+8))
	v117 = int32(2)
	v123 = *(*int32)(unsafe.Add(mBase, uint32(v115+v116<<(uint(v117)%32)+v114<<(uint(v117)%32))))
	F_emscripten_builtin_free(m, v123+int32(-4))
	mBase = m.M
	v128 = v114 + int32(1)
	v129 = *(*int32)(unsafe.Add(mBase, uint32(v107)))
	if base.Ui32(v128) < base.Ui32(v129) {
		v114 = v128
		goto L43
	} else {
		goto L45
	}
L44:
	;
	goto L41
L45:
	;
	goto L44
L46:
	;
	m.Env.Exit(m, int32(1))
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L47:
	;
	v141 = *(*int32)(unsafe.Add(mBase, uint32(v9)+32))
	v145 = *(*int32)(unsafe.Add(mBase, uint32(v141+v56<<(uint(int32(2))%32))))
	goto L48
L48:
	;
	v147 = *(*int32)(unsafe.Add(mBase, _consts[9]))
	v148 = F___strerror_l(m, v147, v147)
	mBase = m.M
	goto L49
L49:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9)+20)) = v148
	*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = v145
	F__serverLog(m, int32(3), int32(_a395), v9+int32(16))
	mBase = m.M
	v156 = m.ExcPending
	if v156 != 0 {
		goto L1
	} else {
		goto L50
	}
L50:
	;
	goto L46
L51:
	;
	v195 = *(*int32)(unsafe.Add(mBase, _consts[15]))
	if int32(3) < v195 {
		goto L62
	} else {
		goto L63
	}
L52:
	;
	v167 = F_fgets(m, v9+int32(64), int32(1025), v161)
	mBase = m.M
	v168 = m.ExcPending
	if v168 != 0 {
		goto L1
	} else {
		goto L54
	}
L53:
	;
	v192 = F_fclose(m, v161)
	mBase = m.M
	v193 = m.ExcPending
	if v193 != 0 {
		goto L1
	} else {
		goto L61
	}
L54:
	;
	if v167 == int32(0) {
		v190 = v11
		goto L53
	} else {
		goto L55
	}
L55:
	;
	v175 = v11
	goto L56
L56:
	;
	v179 = F_sdscat(m, v175, v9+int32(64))
	mBase = m.M
	v180 = m.ExcPending
	if v180 != 0 {
		goto L1
	} else {
		goto L58
	}
L57:
	;
	v190 = v179
	goto L53
L58:
	;
	v184 = F_fgets(m, v9+int32(64), int32(1025), v161)
	mBase = m.M
	v185 = m.ExcPending
	if v185 != 0 {
		goto L1
	} else {
		goto L59
	}
L59:
	;
	if v184 != 0 {
		v175 = v179
		goto L56
	} else {
		goto L60
	}
L60:
	;
	goto L57
L61:
	;
	v213 = v190
	goto L3
L62:
	;
	m.Env.Exit(m, int32(1))
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L63:
	;
	goto L64
L64:
	;
	v199 = *(*int32)(unsafe.Add(mBase, _consts[9]))
	v200 = F___strerror_l(m, v199, v199)
	mBase = m.M
	goto L65
L65:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9)+4)) = v200
	*(*int32)(unsafe.Add(mBase, uint32(v9))) = l0
	F__serverLog(m, int32(3), int32(_a395), v9)
	mBase = m.M
	v206 = m.ExcPending
	if v206 != 0 {
		goto L1
	} else {
		goto L66
	}
L66:
	;
	goto L62
L67:
	;
	if l2 == int32(0) {
		v263 = v254
		goto L79
	} else {
		goto L80
	}
L68:
	;
	v218 = *(*int32)(unsafe.Add(mBase, _consts[15]))
	if int32(2) < v218 {
		goto L69
	} else {
		goto L70
	}
L69:
	;
	v230 = *(*int32)(unsafe.Add(mBase, _consts[169]))
	v231 = F_fgets(m, v9+int32(64), int32(1025), v230)
	mBase = m.M
	v232 = m.ExcPending
	if v232 != 0 {
		goto L1
	} else {
		goto L72
	}
L70:
	;
	F__serverLog(m, int32(2), int32(_a396), int32(0))
	mBase = m.M
	v225 = m.ExcPending
	if v225 != 0 {
		goto L1
	} else {
		goto L71
	}
L71:
	;
	goto L69
L72:
	;
	if v231 == int32(0) {
		v254 = v213
		goto L67
	} else {
		goto L73
	}
L73:
	;
	v239 = v213
	goto L74
L74:
	;
	v243 = F_sdscat(m, v239, v9+int32(64))
	mBase = m.M
	v244 = m.ExcPending
	if v244 != 0 {
		goto L1
	} else {
		goto L76
	}
L75:
	;
	v254 = v243
	goto L67
L76:
	;
	v248 = F_fgets(m, v9+int32(64), int32(1025), v230)
	mBase = m.M
	v249 = m.ExcPending
	if v249 != 0 {
		goto L1
	} else {
		goto L77
	}
L77:
	;
	if v248 != 0 {
		v239 = v243
		goto L74
	} else {
		goto L78
	}
L78:
	;
	goto L75
L79:
	;
	F_loadServerConfigFromString(m, v263)
	mBase = m.M
	v265 = m.ExcPending
	if v265 != 0 {
		goto L1
	} else {
		goto L83
	}
L80:
	;
	v259 = F_sdscat(m, v254, int32(_a397))
	mBase = m.M
	v260 = m.ExcPending
	if v260 != 0 {
		goto L1
	} else {
		goto L81
	}
L81:
	;
	v261 = F_sdscat(m, v259, l2)
	mBase = m.M
	v262 = m.ExcPending
	if v262 != 0 {
		goto L1
	} else {
		goto L82
	}
L82:
	;
	v263 = v261
	goto L79
L83:
	;
	F_sdsfree(m, v263)
	mBase = m.M
	v267 = m.ExcPending
	if v267 != 0 {
		goto L1
	} else {
		goto L84
	}
L84:
	;
	m.G0 = v9 + int32(1104)
	return
}
func F_serverCommunicateSystemd(m *base.Module, l0 int32) int32 {
	return int32(0)
}
func F_serverFork(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v33 int32
	_ = v33
	var v37 int32
	_ = v37
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v57 int64
	_ = v57
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v78 int32
	_ = v78
	var v101 int32
	_ = v101
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
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v130 int32
	_ = v130
	var v133 int32
	_ = v133
	var v136 int32
	_ = v136
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v143 int32
	_ = v143
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v159 int32
	_ = v159
	var v163 int32
	_ = v163
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v178 int32
	_ = v178
	var v180 int64
	_ = v180
	var v185 int64
	_ = v185
	var v198 int32
	_ = v198
	var v202 int32
	_ = v202
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v209 int32
	_ = v209
	var v216 int32
	_ = v216
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v223 int32
	_ = v223
	var v226 int32
	_ = v226
	var v229 int32
	_ = v229
	var v232 int32
	_ = v232
	var v235 int32
	_ = v235
	var v238 int32
	_ = v238
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v246 int32
	_ = v246
	var v248 int32
	_ = v248
	var v249 int32
	_ = v249
	var v257 int32
	_ = v257
	var v258 int32
	_ = v258
	var v260 int32
	_ = v260
	var v268 int32
	_ = v268
	var v269 int32
	_ = v269
	var v270 int32
	_ = v270
	var v273 int32
	_ = v273
	var v275 int32
	_ = v275
	var v285 int32
	_ = v285
	var v286 int64
	_ = v286
	var v293 int64
	_ = v293
	var v301 int32
	_ = v301
	var v310 int32
	_ = v310
	var v311 int64
	_ = v311
	var v331 int32
	_ = v331
	var v335 int64
	_ = v335
	var v336 int32
	_ = v336
	var v337 int32
	_ = v337
	var v339 int32
	_ = v339
	var v343 int32
	_ = v343
	var v344 int32
	_ = v344
	var v345 int64
	_ = v345
	var v347 int64
	_ = v347
	var v349 int32
	_ = v349
	var v351 int32
	_ = v351
	var v353 int64
	_ = v353
	var v356 int32
	_ = v356
	var v357 int32
	_ = v357
	var v358 int32
	_ = v358
	var v363 int32
	_ = v363
	var v364 int32
	_ = v364
	var v366 int32
	_ = v366
	var v367 int32
	_ = v367
	var v373 int32
	_ = v373
	var v376 int32
	_ = v376
	var v378 int32
	_ = v378
	v7 = m.G0
	v9 = v7 - int32(144)
	m.G0 = v9
	if base.Ui32(int32(5)) < base.Ui32(l0) {
		v57 = F_ustime(m)
		mBase = m.M
		v58 = F___errno_location(m)
		mBase = m.M
		*(*int32)(unsafe.Add(mBase, uint32(v58))) = int32(52)
		v61 = int32(-1)
		switch int32(0) {
		case 0:
			v147 = int32(9116376)
			v148 = *(*int32)(unsafe.Add(mBase, _consts[9]))
			if base.Ui32(int32(5)) < base.Ui32(l0) {
			} else {
				if int32(1)<<(uint(l0)%32)&int32(54) == int32(0) {
				} else {
					v159 = *(*int32)(unsafe.Add(mBase, _consts[1054]))
					if v159 != int32(-1) {
						v166 = F_close(m, v159)
						mBase = m.M
						v167 = int32(_a69)
						v168 = *(*int32)(unsafe.Add(mBase, _consts[1055]))
						v169 = F_close(m, v168)
						mBase = m.M
						*(*int32)(unsafe.Add(mBase, _consts[1056])) = int32(0)
						*(*int64)(unsafe.Add(mBase, _consts[1054])) = int64(-1)
					} else {
						v163 = *(*int32)(unsafe.Add(mBase, _consts[1055]))
						if v163 == int32(-1) {
						} else {
							v166 = F_close(m, v159)
							mBase = m.M
							v167 = int32(_a69)
							v168 = *(*int32)(unsafe.Add(mBase, _consts[1055]))
							v169 = F_close(m, v168)
							mBase = m.M
							*(*int32)(unsafe.Add(mBase, _consts[1056])) = int32(0)
							*(*int64)(unsafe.Add(mBase, _consts[1054])) = int64(-1)
						}
					}
				}
			}
			*(*int32)(unsafe.Add(mBase, _consts[9])) = v148
			v378 = int32(-1)
			m.G0 = v9 + int32(144)
			return v378
		case 1:
			v64 = int32(0)
			*(*int32)(unsafe.Add(mBase, _consts[46])) = l0
			*(*int64)(unsafe.Add(mBase, uint32(v9+int32(8)))) = int64(0)
			*(*int32)(unsafe.Add(mBase, uint32(v9)+4)) = int32(1032)
			*(*int32)(unsafe.Add(mBase, uint32(v9)+136)) = int32(0)
			v78 = v9 + int32(4)
			if v78 == int32(0) {
			} else {
				v101 = F___memcpy(m, int32(9118584), v78, int32(140))
				mBase = m.M
			}
			v104 = int32(2)
			v105 = int32(0)
			v106 = *(*int32)(unsafe.Add(mBase, _consts[39]))
			v111 = *(*int32)(unsafe.Add(mBase, _consts[806]))
			if v111 != 0 {
				v112 = base.B2i32(v106 != int32(-1))
			} else {
				v112 = v104
			}
			v114 = *(*int32)(unsafe.Add(mBase, _consts[46]))
			if v114 != 0 {
				v115 = v104
			} else {
				v115 = v112
			}
			*(*int32)(unsafe.Add(mBase, _consts[244])) = v115
			*(*int32)(unsafe.Add(mBase, _consts[295])) = v115
			F_closeListeningSockets(m, int32(0))
			mBase = m.M
			v124 = m.ExcPending
			if v124 != 0 {
				return int32(0)
			} else {
				v125 = int32(0)
				v126 = *(*int32)(unsafe.Add(mBase, _consts[63]))
				if v126 == v125 {
				} else {
					v130 = *(*int32)(unsafe.Add(mBase, _consts[1057]))
					if v130 == int32(-1) {
					} else {
						v133 = F_close(m, v130)
						mBase = m.M
					}
				}
				v136 = *(*int32)(unsafe.Add(mBase, _consts[1053]))
				F_valkey_free(m, v136)
				mBase = m.M
				v138 = m.ExcPending
				if v138 != 0 {
					return int32(0)
				} else {
					v139 = int32(0)
					*(*int32)(unsafe.Add(mBase, _consts[1053])) = v139
					v143 = *(*int32)(unsafe.Add(mBase, _consts[1054]))
					if v143 == int32(-1) {
						v378 = v64
					} else {
						v146 = F_close(m, v143)
						mBase = m.M
						v378 = v64
					}
					m.G0 = v9 + int32(144)
					return v378
				}
			}
		default:
			v178 = int32(0)
			v180 = *(*int64)(unsafe.Add(mBase, _consts[979]))
			*(*int64)(unsafe.Add(mBase, _consts[979])) = v180 + int64(1)
			v185 = F_ustime(m)
			mBase = m.M
			*(*int64)(unsafe.Add(mBase, _consts[978])) = v185 - v57
			v198 = *(*int32)(unsafe.Add(mBase, _consts[276]))
			if v198 < int32(261) {
				if v198 < int32(1) {
					v275 = v178
				} else {
					v206 = v178
					v207 = v198
					v209 = v207 & int32(3)
					if base.Ui32(int32(4)) <= base.Ui32(v207) {
						v216 = int32(0)
						v218 = v206
						v219 = v216
						v223 = v216
						for {
							v226 = v219 << (uint(int32(2)) % 32)
							v229 = *(*int32)(unsafe.Add(mBase, uint32(v226)+uint32(_consts[279])))
							v232 = *(*int32)(unsafe.Add(mBase, uint32(v226)+uint32(_consts[280])))
							v235 = *(*int32)(unsafe.Add(mBase, uint32(v226)+uint32(_consts[281])))
							v238 = *(*int32)(unsafe.Add(mBase, uint32(v226)+uint32(_consts[282])))
							v242 = v229 + (v232 + (v235 + (v238 + v218)))
							v243 = int32(4)
							v244 = v219 + v243
							v246 = v223 + v243
							if v246 != v207&int32(2147483644) {
								v218 = v242
								v219 = v244
								v223 = v246
								continue
							} else {
								break
							}
							break
						}
						v248 = v242
						v249 = v244
					} else {
						v248 = v206
						v249 = int32(0)
					}
					if v209 == int32(0) {
						v275 = v248
					} else {
						v257 = v248
						v258 = v249
						v260 = int32(0)
						for {
							v268 = *(*int32)(unsafe.Add(mBase, uint32(v258<<(uint(int32(2))%32))+uint32(_consts[282])))
							v269 = v268 + v257
							v270 = int32(1)
							v273 = v260 + v270
							if v273 != v209 {
								v257 = v269
								v258 = v258 + v270
								v260 = v273
								continue
							} else {
								break
							}
							break
						}
						v275 = v269
					}
				}
			} else {
				v202 = *(*int32)(unsafe.Add(mBase, _consts[278]))
				v206 = v202
				v207 = int32(260)
				v209 = v207 & int32(3)
				if base.Ui32(int32(4)) <= base.Ui32(v207) {
					v216 = int32(0)
					v218 = v206
					v219 = v216
					v223 = v216
					for {
						v226 = v219 << (uint(int32(2)) % 32)
						v229 = *(*int32)(unsafe.Add(mBase, uint32(v226)+uint32(_consts[279])))
						v232 = *(*int32)(unsafe.Add(mBase, uint32(v226)+uint32(_consts[280])))
						v235 = *(*int32)(unsafe.Add(mBase, uint32(v226)+uint32(_consts[281])))
						v238 = *(*int32)(unsafe.Add(mBase, uint32(v226)+uint32(_consts[282])))
						v242 = v229 + (v232 + (v235 + (v238 + v218)))
						v243 = int32(4)
						v244 = v219 + v243
						v246 = v223 + v243
						if v246 != v207&int32(2147483644) {
							v218 = v242
							v219 = v244
							v223 = v246
							continue
						} else {
							break
						}
						break
					}
					v248 = v242
					v249 = v244
				} else {
					v248 = v206
					v249 = int32(0)
				}
				if v209 == int32(0) {
					v275 = v248
				} else {
					v257 = v248
					v258 = v249
					v260 = int32(0)
					for {
						v268 = *(*int32)(unsafe.Add(mBase, uint32(v258<<(uint(int32(2))%32))+uint32(_consts[282])))
						v269 = v268 + v257
						v270 = int32(1)
						v273 = v260 + v270
						if v273 != v209 {
							v257 = v269
							v258 = v258 + v270
							v260 = v273
							continue
						} else {
							break
						}
						break
					}
					v275 = v269
				}
			}
			v285 = int32(0)
			v286 = *(*int64)(unsafe.Add(mBase, _consts[978]))
			*(*float64)(unsafe.Add(mBase, _consts[1058])) = base.F64_mul(base.F64_div(base.F64_mul(base.F64_convert_i32_u(v275), float64(1e+06)), base.F64_convert_i64_s(v286)), float64(9.313225746154785e-10))
			v293 = *(*int64)(unsafe.Add(mBase, _consts[270]))
			if v293 == int64(0) {
				if base.Ui32(int32(5)) < base.Ui32(l0) {
				} else {
					if int32(1)<<(uint(l0)%32)&int32(54) == int32(0) {
					} else {
						v310 = int32(0)
						v311 = int64(0)
						*(*int64)(unsafe.Add(mBase, _consts[851])) = v311
						*(*int32)(unsafe.Add(mBase, _consts[38])) = l0
						*(*int32)(unsafe.Add(mBase, _consts[39])) = v61
						*(*int64)(unsafe.Add(mBase, _consts[888])) = v311
						*(*int64)(unsafe.Add(mBase, _consts[889])) = v311
						*(*int32)(unsafe.Add(mBase, _consts[886])) = v310
						v331 = *(*int32)(unsafe.Add(mBase, _consts[10]))
						if v331 < int32(1) {
							v353 = v311
						} else {
							v335 = v311
							v336 = int32(0)
							for {
								v337 = F_dbHasNoKeys(m, v336)
								mBase = m.M
								if v337 != 0 {
									v347 = v335
								} else {
									v339 = *(*int32)(unsafe.Add(mBase, _consts[23]))
									v343 = *(*int32)(unsafe.Add(mBase, uint32(v339+v336<<(uint(int32(2))%32))))
									v344 = *(*int32)(unsafe.Add(mBase, uint32(v343)))
									v345 = F_kvstoreSize(m, v344)
									mBase = m.M
									v347 = v345 + v335
								}
								v349 = v336 + int32(1)
								v351 = *(*int32)(unsafe.Add(mBase, _consts[10]))
								if v349 < v351 {
									v335 = v347
									v336 = v349
									continue
								} else {
									break
								}
								break
							}
							v353 = v347
						}
						*(*uint32)(unsafe.Add(mBase, _consts[885])) = uint32(v353)
					}
				}
				v356 = int32(2)
				v357 = int32(0)
				v358 = *(*int32)(unsafe.Add(mBase, _consts[39]))
				v363 = *(*int32)(unsafe.Add(mBase, _consts[806]))
				if v363 != 0 {
					v364 = base.B2i32(v358 != int32(-1))
				} else {
					v364 = v356
				}
				v366 = *(*int32)(unsafe.Add(mBase, _consts[46]))
				if v366 != 0 {
					v367 = v356
				} else {
					v367 = v364
				}
				*(*int32)(unsafe.Add(mBase, _consts[244])) = v367
				*(*int32)(unsafe.Add(mBase, _consts[295])) = v367
				v373 = int32(0)
				F_moduleFireServerEvent(m, int64(13), v373, v373)
				mBase = m.M
				v376 = m.ExcPending
				if v376 != 0 {
					return int32(0)
				} else {
					v378 = v61
					m.G0 = v9 + int32(144)
					return v378
				}
			} else {
				if v286 < v293*int64(1000) {
					if base.Ui32(int32(5)) < base.Ui32(l0) {
					} else {
						if int32(1)<<(uint(l0)%32)&int32(54) == int32(0) {
						} else {
							v310 = int32(0)
							v311 = int64(0)
							*(*int64)(unsafe.Add(mBase, _consts[851])) = v311
							*(*int32)(unsafe.Add(mBase, _consts[38])) = l0
							*(*int32)(unsafe.Add(mBase, _consts[39])) = v61
							*(*int64)(unsafe.Add(mBase, _consts[888])) = v311
							*(*int64)(unsafe.Add(mBase, _consts[889])) = v311
							*(*int32)(unsafe.Add(mBase, _consts[886])) = v310
							v331 = *(*int32)(unsafe.Add(mBase, _consts[10]))
							if v331 < int32(1) {
								v353 = v311
							} else {
								v335 = v311
								v336 = int32(0)
								for {
									v337 = F_dbHasNoKeys(m, v336)
									mBase = m.M
									if v337 != 0 {
										v347 = v335
									} else {
										v339 = *(*int32)(unsafe.Add(mBase, _consts[23]))
										v343 = *(*int32)(unsafe.Add(mBase, uint32(v339+v336<<(uint(int32(2))%32))))
										v344 = *(*int32)(unsafe.Add(mBase, uint32(v343)))
										v345 = F_kvstoreSize(m, v344)
										mBase = m.M
										v347 = v345 + v335
									}
									v349 = v336 + int32(1)
									v351 = *(*int32)(unsafe.Add(mBase, _consts[10]))
									if v349 < v351 {
										v335 = v347
										v336 = v349
										continue
									} else {
										break
									}
									break
								}
								v353 = v347
							}
							*(*uint32)(unsafe.Add(mBase, _consts[885])) = uint32(v353)
						}
					}
					v356 = int32(2)
					v357 = int32(0)
					v358 = *(*int32)(unsafe.Add(mBase, _consts[39]))
					v363 = *(*int32)(unsafe.Add(mBase, _consts[806]))
					if v363 != 0 {
						v364 = base.B2i32(v358 != int32(-1))
					} else {
						v364 = v356
					}
					v366 = *(*int32)(unsafe.Add(mBase, _consts[46]))
					if v366 != 0 {
						v367 = v356
					} else {
						v367 = v364
					}
					*(*int32)(unsafe.Add(mBase, _consts[244])) = v367
					*(*int32)(unsafe.Add(mBase, _consts[295])) = v367
					v373 = int32(0)
					F_moduleFireServerEvent(m, int64(13), v373, v373)
					mBase = m.M
					v376 = m.ExcPending
					if v376 != 0 {
						return int32(0)
					} else {
						v378 = v61
						m.G0 = v9 + int32(144)
						return v378
					}
				} else {
					F_latencyAddSample(m, int32(_a1456), v286)
					mBase = m.M
					v301 = m.ExcPending
					if v301 != 0 {
						return int32(0)
					} else {
						if base.Ui32(int32(5)) < base.Ui32(l0) {
						} else {
							if int32(1)<<(uint(l0)%32)&int32(54) == int32(0) {
							} else {
								v310 = int32(0)
								v311 = int64(0)
								*(*int64)(unsafe.Add(mBase, _consts[851])) = v311
								*(*int32)(unsafe.Add(mBase, _consts[38])) = l0
								*(*int32)(unsafe.Add(mBase, _consts[39])) = v61
								*(*int64)(unsafe.Add(mBase, _consts[888])) = v311
								*(*int64)(unsafe.Add(mBase, _consts[889])) = v311
								*(*int32)(unsafe.Add(mBase, _consts[886])) = v310
								v331 = *(*int32)(unsafe.Add(mBase, _consts[10]))
								if v331 < int32(1) {
									v353 = v311
								} else {
									v335 = v311
									v336 = int32(0)
									for {
										v337 = F_dbHasNoKeys(m, v336)
										mBase = m.M
										if v337 != 0 {
											v347 = v335
										} else {
											v339 = *(*int32)(unsafe.Add(mBase, _consts[23]))
											v343 = *(*int32)(unsafe.Add(mBase, uint32(v339+v336<<(uint(int32(2))%32))))
											v344 = *(*int32)(unsafe.Add(mBase, uint32(v343)))
											v345 = F_kvstoreSize(m, v344)
											mBase = m.M
											v347 = v345 + v335
										}
										v349 = v336 + int32(1)
										v351 = *(*int32)(unsafe.Add(mBase, _consts[10]))
										if v349 < v351 {
											v335 = v347
											v336 = v349
											continue
										} else {
											break
										}
										break
									}
									v353 = v347
								}
								*(*uint32)(unsafe.Add(mBase, _consts[885])) = uint32(v353)
							}
						}
						v356 = int32(2)
						v357 = int32(0)
						v358 = *(*int32)(unsafe.Add(mBase, _consts[39]))
						v363 = *(*int32)(unsafe.Add(mBase, _consts[806]))
						if v363 != 0 {
							v364 = base.B2i32(v358 != int32(-1))
						} else {
							v364 = v356
						}
						v366 = *(*int32)(unsafe.Add(mBase, _consts[46]))
						if v366 != 0 {
							v367 = v356
						} else {
							v367 = v364
						}
						*(*int32)(unsafe.Add(mBase, _consts[244])) = v367
						*(*int32)(unsafe.Add(mBase, _consts[295])) = v367
						v373 = int32(0)
						F_moduleFireServerEvent(m, int64(13), v373, v373)
						mBase = m.M
						v376 = m.ExcPending
						if v376 != 0 {
							return int32(0)
						} else {
							v378 = v61
							m.G0 = v9 + int32(144)
							return v378
						}
					}
				}
			}
		}
	} else {
		if int32(1)<<(uint(l0)%32)&int32(54) == int32(0) {
			v57 = F_ustime(m)
			mBase = m.M
			v58 = F___errno_location(m)
			mBase = m.M
			*(*int32)(unsafe.Add(mBase, uint32(v58))) = int32(52)
			v61 = int32(-1)
			switch int32(0) {
			case 0:
				v147 = int32(9116376)
				v148 = *(*int32)(unsafe.Add(mBase, _consts[9]))
				if base.Ui32(int32(5)) < base.Ui32(l0) {
				} else {
					if int32(1)<<(uint(l0)%32)&int32(54) == int32(0) {
					} else {
						v159 = *(*int32)(unsafe.Add(mBase, _consts[1054]))
						if v159 != int32(-1) {
							v166 = F_close(m, v159)
							mBase = m.M
							v167 = int32(_a69)
							v168 = *(*int32)(unsafe.Add(mBase, _consts[1055]))
							v169 = F_close(m, v168)
							mBase = m.M
							*(*int32)(unsafe.Add(mBase, _consts[1056])) = int32(0)
							*(*int64)(unsafe.Add(mBase, _consts[1054])) = int64(-1)
						} else {
							v163 = *(*int32)(unsafe.Add(mBase, _consts[1055]))
							if v163 == int32(-1) {
							} else {
								v166 = F_close(m, v159)
								mBase = m.M
								v167 = int32(_a69)
								v168 = *(*int32)(unsafe.Add(mBase, _consts[1055]))
								v169 = F_close(m, v168)
								mBase = m.M
								*(*int32)(unsafe.Add(mBase, _consts[1056])) = int32(0)
								*(*int64)(unsafe.Add(mBase, _consts[1054])) = int64(-1)
							}
						}
					}
				}
				*(*int32)(unsafe.Add(mBase, _consts[9])) = v148
				v378 = int32(-1)
				m.G0 = v9 + int32(144)
				return v378
			case 1:
				v64 = int32(0)
				*(*int32)(unsafe.Add(mBase, _consts[46])) = l0
				*(*int64)(unsafe.Add(mBase, uint32(v9+int32(8)))) = int64(0)
				*(*int32)(unsafe.Add(mBase, uint32(v9)+4)) = int32(1032)
				*(*int32)(unsafe.Add(mBase, uint32(v9)+136)) = int32(0)
				v78 = v9 + int32(4)
				if v78 == int32(0) {
				} else {
					v101 = F___memcpy(m, int32(9118584), v78, int32(140))
					mBase = m.M
				}
				v104 = int32(2)
				v105 = int32(0)
				v106 = *(*int32)(unsafe.Add(mBase, _consts[39]))
				v111 = *(*int32)(unsafe.Add(mBase, _consts[806]))
				if v111 != 0 {
					v112 = base.B2i32(v106 != int32(-1))
				} else {
					v112 = v104
				}
				v114 = *(*int32)(unsafe.Add(mBase, _consts[46]))
				if v114 != 0 {
					v115 = v104
				} else {
					v115 = v112
				}
				*(*int32)(unsafe.Add(mBase, _consts[244])) = v115
				*(*int32)(unsafe.Add(mBase, _consts[295])) = v115
				F_closeListeningSockets(m, int32(0))
				mBase = m.M
				v124 = m.ExcPending
				if v124 != 0 {
					return int32(0)
				} else {
					v125 = int32(0)
					v126 = *(*int32)(unsafe.Add(mBase, _consts[63]))
					if v126 == v125 {
					} else {
						v130 = *(*int32)(unsafe.Add(mBase, _consts[1057]))
						if v130 == int32(-1) {
						} else {
							v133 = F_close(m, v130)
							mBase = m.M
						}
					}
					v136 = *(*int32)(unsafe.Add(mBase, _consts[1053]))
					F_valkey_free(m, v136)
					mBase = m.M
					v138 = m.ExcPending
					if v138 != 0 {
						return int32(0)
					} else {
						v139 = int32(0)
						*(*int32)(unsafe.Add(mBase, _consts[1053])) = v139
						v143 = *(*int32)(unsafe.Add(mBase, _consts[1054]))
						if v143 == int32(-1) {
							v378 = v64
						} else {
							v146 = F_close(m, v143)
							mBase = m.M
							v378 = v64
						}
						m.G0 = v9 + int32(144)
						return v378
					}
				}
			default:
				v178 = int32(0)
				v180 = *(*int64)(unsafe.Add(mBase, _consts[979]))
				*(*int64)(unsafe.Add(mBase, _consts[979])) = v180 + int64(1)
				v185 = F_ustime(m)
				mBase = m.M
				*(*int64)(unsafe.Add(mBase, _consts[978])) = v185 - v57
				v198 = *(*int32)(unsafe.Add(mBase, _consts[276]))
				if v198 < int32(261) {
					if v198 < int32(1) {
						v275 = v178
					} else {
						v206 = v178
						v207 = v198
						v209 = v207 & int32(3)
						if base.Ui32(int32(4)) <= base.Ui32(v207) {
							v216 = int32(0)
							v218 = v206
							v219 = v216
							v223 = v216
							for {
								v226 = v219 << (uint(int32(2)) % 32)
								v229 = *(*int32)(unsafe.Add(mBase, uint32(v226)+uint32(_consts[279])))
								v232 = *(*int32)(unsafe.Add(mBase, uint32(v226)+uint32(_consts[280])))
								v235 = *(*int32)(unsafe.Add(mBase, uint32(v226)+uint32(_consts[281])))
								v238 = *(*int32)(unsafe.Add(mBase, uint32(v226)+uint32(_consts[282])))
								v242 = v229 + (v232 + (v235 + (v238 + v218)))
								v243 = int32(4)
								v244 = v219 + v243
								v246 = v223 + v243
								if v246 != v207&int32(2147483644) {
									v218 = v242
									v219 = v244
									v223 = v246
									continue
								} else {
									break
								}
								break
							}
							v248 = v242
							v249 = v244
						} else {
							v248 = v206
							v249 = int32(0)
						}
						if v209 == int32(0) {
							v275 = v248
						} else {
							v257 = v248
							v258 = v249
							v260 = int32(0)
							for {
								v268 = *(*int32)(unsafe.Add(mBase, uint32(v258<<(uint(int32(2))%32))+uint32(_consts[282])))
								v269 = v268 + v257
								v270 = int32(1)
								v273 = v260 + v270
								if v273 != v209 {
									v257 = v269
									v258 = v258 + v270
									v260 = v273
									continue
								} else {
									break
								}
								break
							}
							v275 = v269
						}
					}
				} else {
					v202 = *(*int32)(unsafe.Add(mBase, _consts[278]))
					v206 = v202
					v207 = int32(260)
					v209 = v207 & int32(3)
					if base.Ui32(int32(4)) <= base.Ui32(v207) {
						v216 = int32(0)
						v218 = v206
						v219 = v216
						v223 = v216
						for {
							v226 = v219 << (uint(int32(2)) % 32)
							v229 = *(*int32)(unsafe.Add(mBase, uint32(v226)+uint32(_consts[279])))
							v232 = *(*int32)(unsafe.Add(mBase, uint32(v226)+uint32(_consts[280])))
							v235 = *(*int32)(unsafe.Add(mBase, uint32(v226)+uint32(_consts[281])))
							v238 = *(*int32)(unsafe.Add(mBase, uint32(v226)+uint32(_consts[282])))
							v242 = v229 + (v232 + (v235 + (v238 + v218)))
							v243 = int32(4)
							v244 = v219 + v243
							v246 = v223 + v243
							if v246 != v207&int32(2147483644) {
								v218 = v242
								v219 = v244
								v223 = v246
								continue
							} else {
								break
							}
							break
						}
						v248 = v242
						v249 = v244
					} else {
						v248 = v206
						v249 = int32(0)
					}
					if v209 == int32(0) {
						v275 = v248
					} else {
						v257 = v248
						v258 = v249
						v260 = int32(0)
						for {
							v268 = *(*int32)(unsafe.Add(mBase, uint32(v258<<(uint(int32(2))%32))+uint32(_consts[282])))
							v269 = v268 + v257
							v270 = int32(1)
							v273 = v260 + v270
							if v273 != v209 {
								v257 = v269
								v258 = v258 + v270
								v260 = v273
								continue
							} else {
								break
							}
							break
						}
						v275 = v269
					}
				}
				v285 = int32(0)
				v286 = *(*int64)(unsafe.Add(mBase, _consts[978]))
				*(*float64)(unsafe.Add(mBase, _consts[1058])) = base.F64_mul(base.F64_div(base.F64_mul(base.F64_convert_i32_u(v275), float64(1e+06)), base.F64_convert_i64_s(v286)), float64(9.313225746154785e-10))
				v293 = *(*int64)(unsafe.Add(mBase, _consts[270]))
				if v293 == int64(0) {
					if base.Ui32(int32(5)) < base.Ui32(l0) {
					} else {
						if int32(1)<<(uint(l0)%32)&int32(54) == int32(0) {
						} else {
							v310 = int32(0)
							v311 = int64(0)
							*(*int64)(unsafe.Add(mBase, _consts[851])) = v311
							*(*int32)(unsafe.Add(mBase, _consts[38])) = l0
							*(*int32)(unsafe.Add(mBase, _consts[39])) = v61
							*(*int64)(unsafe.Add(mBase, _consts[888])) = v311
							*(*int64)(unsafe.Add(mBase, _consts[889])) = v311
							*(*int32)(unsafe.Add(mBase, _consts[886])) = v310
							v331 = *(*int32)(unsafe.Add(mBase, _consts[10]))
							if v331 < int32(1) {
								v353 = v311
							} else {
								v335 = v311
								v336 = int32(0)
								for {
									v337 = F_dbHasNoKeys(m, v336)
									mBase = m.M
									if v337 != 0 {
										v347 = v335
									} else {
										v339 = *(*int32)(unsafe.Add(mBase, _consts[23]))
										v343 = *(*int32)(unsafe.Add(mBase, uint32(v339+v336<<(uint(int32(2))%32))))
										v344 = *(*int32)(unsafe.Add(mBase, uint32(v343)))
										v345 = F_kvstoreSize(m, v344)
										mBase = m.M
										v347 = v345 + v335
									}
									v349 = v336 + int32(1)
									v351 = *(*int32)(unsafe.Add(mBase, _consts[10]))
									if v349 < v351 {
										v335 = v347
										v336 = v349
										continue
									} else {
										break
									}
									break
								}
								v353 = v347
							}
							*(*uint32)(unsafe.Add(mBase, _consts[885])) = uint32(v353)
						}
					}
					v356 = int32(2)
					v357 = int32(0)
					v358 = *(*int32)(unsafe.Add(mBase, _consts[39]))
					v363 = *(*int32)(unsafe.Add(mBase, _consts[806]))
					if v363 != 0 {
						v364 = base.B2i32(v358 != int32(-1))
					} else {
						v364 = v356
					}
					v366 = *(*int32)(unsafe.Add(mBase, _consts[46]))
					if v366 != 0 {
						v367 = v356
					} else {
						v367 = v364
					}
					*(*int32)(unsafe.Add(mBase, _consts[244])) = v367
					*(*int32)(unsafe.Add(mBase, _consts[295])) = v367
					v373 = int32(0)
					F_moduleFireServerEvent(m, int64(13), v373, v373)
					mBase = m.M
					v376 = m.ExcPending
					if v376 != 0 {
						return int32(0)
					} else {
						v378 = v61
						m.G0 = v9 + int32(144)
						return v378
					}
				} else {
					if v286 < v293*int64(1000) {
						if base.Ui32(int32(5)) < base.Ui32(l0) {
						} else {
							if int32(1)<<(uint(l0)%32)&int32(54) == int32(0) {
							} else {
								v310 = int32(0)
								v311 = int64(0)
								*(*int64)(unsafe.Add(mBase, _consts[851])) = v311
								*(*int32)(unsafe.Add(mBase, _consts[38])) = l0
								*(*int32)(unsafe.Add(mBase, _consts[39])) = v61
								*(*int64)(unsafe.Add(mBase, _consts[888])) = v311
								*(*int64)(unsafe.Add(mBase, _consts[889])) = v311
								*(*int32)(unsafe.Add(mBase, _consts[886])) = v310
								v331 = *(*int32)(unsafe.Add(mBase, _consts[10]))
								if v331 < int32(1) {
									v353 = v311
								} else {
									v335 = v311
									v336 = int32(0)
									for {
										v337 = F_dbHasNoKeys(m, v336)
										mBase = m.M
										if v337 != 0 {
											v347 = v335
										} else {
											v339 = *(*int32)(unsafe.Add(mBase, _consts[23]))
											v343 = *(*int32)(unsafe.Add(mBase, uint32(v339+v336<<(uint(int32(2))%32))))
											v344 = *(*int32)(unsafe.Add(mBase, uint32(v343)))
											v345 = F_kvstoreSize(m, v344)
											mBase = m.M
											v347 = v345 + v335
										}
										v349 = v336 + int32(1)
										v351 = *(*int32)(unsafe.Add(mBase, _consts[10]))
										if v349 < v351 {
											v335 = v347
											v336 = v349
											continue
										} else {
											break
										}
										break
									}
									v353 = v347
								}
								*(*uint32)(unsafe.Add(mBase, _consts[885])) = uint32(v353)
							}
						}
						v356 = int32(2)
						v357 = int32(0)
						v358 = *(*int32)(unsafe.Add(mBase, _consts[39]))
						v363 = *(*int32)(unsafe.Add(mBase, _consts[806]))
						if v363 != 0 {
							v364 = base.B2i32(v358 != int32(-1))
						} else {
							v364 = v356
						}
						v366 = *(*int32)(unsafe.Add(mBase, _consts[46]))
						if v366 != 0 {
							v367 = v356
						} else {
							v367 = v364
						}
						*(*int32)(unsafe.Add(mBase, _consts[244])) = v367
						*(*int32)(unsafe.Add(mBase, _consts[295])) = v367
						v373 = int32(0)
						F_moduleFireServerEvent(m, int64(13), v373, v373)
						mBase = m.M
						v376 = m.ExcPending
						if v376 != 0 {
							return int32(0)
						} else {
							v378 = v61
							m.G0 = v9 + int32(144)
							return v378
						}
					} else {
						F_latencyAddSample(m, int32(_a1456), v286)
						mBase = m.M
						v301 = m.ExcPending
						if v301 != 0 {
							return int32(0)
						} else {
							if base.Ui32(int32(5)) < base.Ui32(l0) {
							} else {
								if int32(1)<<(uint(l0)%32)&int32(54) == int32(0) {
								} else {
									v310 = int32(0)
									v311 = int64(0)
									*(*int64)(unsafe.Add(mBase, _consts[851])) = v311
									*(*int32)(unsafe.Add(mBase, _consts[38])) = l0
									*(*int32)(unsafe.Add(mBase, _consts[39])) = v61
									*(*int64)(unsafe.Add(mBase, _consts[888])) = v311
									*(*int64)(unsafe.Add(mBase, _consts[889])) = v311
									*(*int32)(unsafe.Add(mBase, _consts[886])) = v310
									v331 = *(*int32)(unsafe.Add(mBase, _consts[10]))
									if v331 < int32(1) {
										v353 = v311
									} else {
										v335 = v311
										v336 = int32(0)
										for {
											v337 = F_dbHasNoKeys(m, v336)
											mBase = m.M
											if v337 != 0 {
												v347 = v335
											} else {
												v339 = *(*int32)(unsafe.Add(mBase, _consts[23]))
												v343 = *(*int32)(unsafe.Add(mBase, uint32(v339+v336<<(uint(int32(2))%32))))
												v344 = *(*int32)(unsafe.Add(mBase, uint32(v343)))
												v345 = F_kvstoreSize(m, v344)
												mBase = m.M
												v347 = v345 + v335
											}
											v349 = v336 + int32(1)
											v351 = *(*int32)(unsafe.Add(mBase, _consts[10]))
											if v349 < v351 {
												v335 = v347
												v336 = v349
												continue
											} else {
												break
											}
											break
										}
										v353 = v347
									}
									*(*uint32)(unsafe.Add(mBase, _consts[885])) = uint32(v353)
								}
							}
							v356 = int32(2)
							v357 = int32(0)
							v358 = *(*int32)(unsafe.Add(mBase, _consts[39]))
							v363 = *(*int32)(unsafe.Add(mBase, _consts[806]))
							if v363 != 0 {
								v364 = base.B2i32(v358 != int32(-1))
							} else {
								v364 = v356
							}
							v366 = *(*int32)(unsafe.Add(mBase, _consts[46]))
							if v366 != 0 {
								v367 = v356
							} else {
								v367 = v364
							}
							*(*int32)(unsafe.Add(mBase, _consts[244])) = v367
							*(*int32)(unsafe.Add(mBase, _consts[295])) = v367
							v373 = int32(0)
							F_moduleFireServerEvent(m, int64(13), v373, v373)
							mBase = m.M
							v376 = m.ExcPending
							if v376 != 0 {
								return int32(0)
							} else {
								v378 = v61
								m.G0 = v9 + int32(144)
								return v378
							}
						}
					}
				}
			}
		} else {
			v19 = int32(-1)
			v21 = *(*int32)(unsafe.Add(mBase, _consts[39]))
			if v21 == v19 {
				v33 = F_anetPipe(m, int32(_a1457), int32(2048), int32(0))
				mBase = m.M
				if v33 != int32(-1) {
					*(*int32)(unsafe.Add(mBase, _consts[1056])) = int32(0)
				} else {
					v37 = *(*int32)(unsafe.Add(mBase, _consts[1054]))
					if v37 != int32(-1) {
						v44 = F_close(m, v37)
						mBase = m.M
						v45 = int32(_a69)
						v46 = *(*int32)(unsafe.Add(mBase, _consts[1055]))
						v47 = F_close(m, v46)
						mBase = m.M
						*(*int64)(unsafe.Add(mBase, _consts[1054])) = int64(-1)
						*(*int32)(unsafe.Add(mBase, _consts[1056])) = int32(0)
					} else {
						v41 = *(*int32)(unsafe.Add(mBase, _consts[1055]))
						if v41 == int32(-1) {
						} else {
							v44 = F_close(m, v37)
							mBase = m.M
							v45 = int32(_a69)
							v46 = *(*int32)(unsafe.Add(mBase, _consts[1055]))
							v47 = F_close(m, v46)
							mBase = m.M
							*(*int64)(unsafe.Add(mBase, _consts[1054])) = int64(-1)
							*(*int32)(unsafe.Add(mBase, _consts[1056])) = int32(0)
						}
					}
				}
				v57 = F_ustime(m)
				mBase = m.M
				v58 = F___errno_location(m)
				mBase = m.M
				*(*int32)(unsafe.Add(mBase, uint32(v58))) = int32(52)
				v61 = int32(-1)
				switch int32(0) {
				case 0:
					v147 = int32(9116376)
					v148 = *(*int32)(unsafe.Add(mBase, _consts[9]))
					if base.Ui32(int32(5)) < base.Ui32(l0) {
					} else {
						if int32(1)<<(uint(l0)%32)&int32(54) == int32(0) {
						} else {
							v159 = *(*int32)(unsafe.Add(mBase, _consts[1054]))
							if v159 != int32(-1) {
								v166 = F_close(m, v159)
								mBase = m.M
								v167 = int32(_a69)
								v168 = *(*int32)(unsafe.Add(mBase, _consts[1055]))
								v169 = F_close(m, v168)
								mBase = m.M
								*(*int32)(unsafe.Add(mBase, _consts[1056])) = int32(0)
								*(*int64)(unsafe.Add(mBase, _consts[1054])) = int64(-1)
							} else {
								v163 = *(*int32)(unsafe.Add(mBase, _consts[1055]))
								if v163 == int32(-1) {
								} else {
									v166 = F_close(m, v159)
									mBase = m.M
									v167 = int32(_a69)
									v168 = *(*int32)(unsafe.Add(mBase, _consts[1055]))
									v169 = F_close(m, v168)
									mBase = m.M
									*(*int32)(unsafe.Add(mBase, _consts[1056])) = int32(0)
									*(*int64)(unsafe.Add(mBase, _consts[1054])) = int64(-1)
								}
							}
						}
					}
					*(*int32)(unsafe.Add(mBase, _consts[9])) = v148
					v378 = int32(-1)
					m.G0 = v9 + int32(144)
					return v378
				case 1:
					v64 = int32(0)
					*(*int32)(unsafe.Add(mBase, _consts[46])) = l0
					*(*int64)(unsafe.Add(mBase, uint32(v9+int32(8)))) = int64(0)
					*(*int32)(unsafe.Add(mBase, uint32(v9)+4)) = int32(1032)
					*(*int32)(unsafe.Add(mBase, uint32(v9)+136)) = int32(0)
					v78 = v9 + int32(4)
					if v78 == int32(0) {
					} else {
						v101 = F___memcpy(m, int32(9118584), v78, int32(140))
						mBase = m.M
					}
					v104 = int32(2)
					v105 = int32(0)
					v106 = *(*int32)(unsafe.Add(mBase, _consts[39]))
					v111 = *(*int32)(unsafe.Add(mBase, _consts[806]))
					if v111 != 0 {
						v112 = base.B2i32(v106 != int32(-1))
					} else {
						v112 = v104
					}
					v114 = *(*int32)(unsafe.Add(mBase, _consts[46]))
					if v114 != 0 {
						v115 = v104
					} else {
						v115 = v112
					}
					*(*int32)(unsafe.Add(mBase, _consts[244])) = v115
					*(*int32)(unsafe.Add(mBase, _consts[295])) = v115
					F_closeListeningSockets(m, int32(0))
					mBase = m.M
					v124 = m.ExcPending
					if v124 != 0 {
						return int32(0)
					} else {
						v125 = int32(0)
						v126 = *(*int32)(unsafe.Add(mBase, _consts[63]))
						if v126 == v125 {
						} else {
							v130 = *(*int32)(unsafe.Add(mBase, _consts[1057]))
							if v130 == int32(-1) {
							} else {
								v133 = F_close(m, v130)
								mBase = m.M
							}
						}
						v136 = *(*int32)(unsafe.Add(mBase, _consts[1053]))
						F_valkey_free(m, v136)
						mBase = m.M
						v138 = m.ExcPending
						if v138 != 0 {
							return int32(0)
						} else {
							v139 = int32(0)
							*(*int32)(unsafe.Add(mBase, _consts[1053])) = v139
							v143 = *(*int32)(unsafe.Add(mBase, _consts[1054]))
							if v143 == int32(-1) {
								v378 = v64
							} else {
								v146 = F_close(m, v143)
								mBase = m.M
								v378 = v64
							}
							m.G0 = v9 + int32(144)
							return v378
						}
					}
				default:
					v178 = int32(0)
					v180 = *(*int64)(unsafe.Add(mBase, _consts[979]))
					*(*int64)(unsafe.Add(mBase, _consts[979])) = v180 + int64(1)
					v185 = F_ustime(m)
					mBase = m.M
					*(*int64)(unsafe.Add(mBase, _consts[978])) = v185 - v57
					v198 = *(*int32)(unsafe.Add(mBase, _consts[276]))
					if v198 < int32(261) {
						if v198 < int32(1) {
							v275 = v178
						} else {
							v206 = v178
							v207 = v198
							v209 = v207 & int32(3)
							if base.Ui32(int32(4)) <= base.Ui32(v207) {
								v216 = int32(0)
								v218 = v206
								v219 = v216
								v223 = v216
								for {
									v226 = v219 << (uint(int32(2)) % 32)
									v229 = *(*int32)(unsafe.Add(mBase, uint32(v226)+uint32(_consts[279])))
									v232 = *(*int32)(unsafe.Add(mBase, uint32(v226)+uint32(_consts[280])))
									v235 = *(*int32)(unsafe.Add(mBase, uint32(v226)+uint32(_consts[281])))
									v238 = *(*int32)(unsafe.Add(mBase, uint32(v226)+uint32(_consts[282])))
									v242 = v229 + (v232 + (v235 + (v238 + v218)))
									v243 = int32(4)
									v244 = v219 + v243
									v246 = v223 + v243
									if v246 != v207&int32(2147483644) {
										v218 = v242
										v219 = v244
										v223 = v246
										continue
									} else {
										break
									}
									break
								}
								v248 = v242
								v249 = v244
							} else {
								v248 = v206
								v249 = int32(0)
							}
							if v209 == int32(0) {
								v275 = v248
							} else {
								v257 = v248
								v258 = v249
								v260 = int32(0)
								for {
									v268 = *(*int32)(unsafe.Add(mBase, uint32(v258<<(uint(int32(2))%32))+uint32(_consts[282])))
									v269 = v268 + v257
									v270 = int32(1)
									v273 = v260 + v270
									if v273 != v209 {
										v257 = v269
										v258 = v258 + v270
										v260 = v273
										continue
									} else {
										break
									}
									break
								}
								v275 = v269
							}
						}
					} else {
						v202 = *(*int32)(unsafe.Add(mBase, _consts[278]))
						v206 = v202
						v207 = int32(260)
						v209 = v207 & int32(3)
						if base.Ui32(int32(4)) <= base.Ui32(v207) {
							v216 = int32(0)
							v218 = v206
							v219 = v216
							v223 = v216
							for {
								v226 = v219 << (uint(int32(2)) % 32)
								v229 = *(*int32)(unsafe.Add(mBase, uint32(v226)+uint32(_consts[279])))
								v232 = *(*int32)(unsafe.Add(mBase, uint32(v226)+uint32(_consts[280])))
								v235 = *(*int32)(unsafe.Add(mBase, uint32(v226)+uint32(_consts[281])))
								v238 = *(*int32)(unsafe.Add(mBase, uint32(v226)+uint32(_consts[282])))
								v242 = v229 + (v232 + (v235 + (v238 + v218)))
								v243 = int32(4)
								v244 = v219 + v243
								v246 = v223 + v243
								if v246 != v207&int32(2147483644) {
									v218 = v242
									v219 = v244
									v223 = v246
									continue
								} else {
									break
								}
								break
							}
							v248 = v242
							v249 = v244
						} else {
							v248 = v206
							v249 = int32(0)
						}
						if v209 == int32(0) {
							v275 = v248
						} else {
							v257 = v248
							v258 = v249
							v260 = int32(0)
							for {
								v268 = *(*int32)(unsafe.Add(mBase, uint32(v258<<(uint(int32(2))%32))+uint32(_consts[282])))
								v269 = v268 + v257
								v270 = int32(1)
								v273 = v260 + v270
								if v273 != v209 {
									v257 = v269
									v258 = v258 + v270
									v260 = v273
									continue
								} else {
									break
								}
								break
							}
							v275 = v269
						}
					}
					v285 = int32(0)
					v286 = *(*int64)(unsafe.Add(mBase, _consts[978]))
					*(*float64)(unsafe.Add(mBase, _consts[1058])) = base.F64_mul(base.F64_div(base.F64_mul(base.F64_convert_i32_u(v275), float64(1e+06)), base.F64_convert_i64_s(v286)), float64(9.313225746154785e-10))
					v293 = *(*int64)(unsafe.Add(mBase, _consts[270]))
					if v293 == int64(0) {
						if base.Ui32(int32(5)) < base.Ui32(l0) {
						} else {
							if int32(1)<<(uint(l0)%32)&int32(54) == int32(0) {
							} else {
								v310 = int32(0)
								v311 = int64(0)
								*(*int64)(unsafe.Add(mBase, _consts[851])) = v311
								*(*int32)(unsafe.Add(mBase, _consts[38])) = l0
								*(*int32)(unsafe.Add(mBase, _consts[39])) = v61
								*(*int64)(unsafe.Add(mBase, _consts[888])) = v311
								*(*int64)(unsafe.Add(mBase, _consts[889])) = v311
								*(*int32)(unsafe.Add(mBase, _consts[886])) = v310
								v331 = *(*int32)(unsafe.Add(mBase, _consts[10]))
								if v331 < int32(1) {
									v353 = v311
								} else {
									v335 = v311
									v336 = int32(0)
									for {
										v337 = F_dbHasNoKeys(m, v336)
										mBase = m.M
										if v337 != 0 {
											v347 = v335
										} else {
											v339 = *(*int32)(unsafe.Add(mBase, _consts[23]))
											v343 = *(*int32)(unsafe.Add(mBase, uint32(v339+v336<<(uint(int32(2))%32))))
											v344 = *(*int32)(unsafe.Add(mBase, uint32(v343)))
											v345 = F_kvstoreSize(m, v344)
											mBase = m.M
											v347 = v345 + v335
										}
										v349 = v336 + int32(1)
										v351 = *(*int32)(unsafe.Add(mBase, _consts[10]))
										if v349 < v351 {
											v335 = v347
											v336 = v349
											continue
										} else {
											break
										}
										break
									}
									v353 = v347
								}
								*(*uint32)(unsafe.Add(mBase, _consts[885])) = uint32(v353)
							}
						}
						v356 = int32(2)
						v357 = int32(0)
						v358 = *(*int32)(unsafe.Add(mBase, _consts[39]))
						v363 = *(*int32)(unsafe.Add(mBase, _consts[806]))
						if v363 != 0 {
							v364 = base.B2i32(v358 != int32(-1))
						} else {
							v364 = v356
						}
						v366 = *(*int32)(unsafe.Add(mBase, _consts[46]))
						if v366 != 0 {
							v367 = v356
						} else {
							v367 = v364
						}
						*(*int32)(unsafe.Add(mBase, _consts[244])) = v367
						*(*int32)(unsafe.Add(mBase, _consts[295])) = v367
						v373 = int32(0)
						F_moduleFireServerEvent(m, int64(13), v373, v373)
						mBase = m.M
						v376 = m.ExcPending
						if v376 != 0 {
							return int32(0)
						} else {
							v378 = v61
							m.G0 = v9 + int32(144)
							return v378
						}
					} else {
						if v286 < v293*int64(1000) {
							if base.Ui32(int32(5)) < base.Ui32(l0) {
							} else {
								if int32(1)<<(uint(l0)%32)&int32(54) == int32(0) {
								} else {
									v310 = int32(0)
									v311 = int64(0)
									*(*int64)(unsafe.Add(mBase, _consts[851])) = v311
									*(*int32)(unsafe.Add(mBase, _consts[38])) = l0
									*(*int32)(unsafe.Add(mBase, _consts[39])) = v61
									*(*int64)(unsafe.Add(mBase, _consts[888])) = v311
									*(*int64)(unsafe.Add(mBase, _consts[889])) = v311
									*(*int32)(unsafe.Add(mBase, _consts[886])) = v310
									v331 = *(*int32)(unsafe.Add(mBase, _consts[10]))
									if v331 < int32(1) {
										v353 = v311
									} else {
										v335 = v311
										v336 = int32(0)
										for {
											v337 = F_dbHasNoKeys(m, v336)
											mBase = m.M
											if v337 != 0 {
												v347 = v335
											} else {
												v339 = *(*int32)(unsafe.Add(mBase, _consts[23]))
												v343 = *(*int32)(unsafe.Add(mBase, uint32(v339+v336<<(uint(int32(2))%32))))
												v344 = *(*int32)(unsafe.Add(mBase, uint32(v343)))
												v345 = F_kvstoreSize(m, v344)
												mBase = m.M
												v347 = v345 + v335
											}
											v349 = v336 + int32(1)
											v351 = *(*int32)(unsafe.Add(mBase, _consts[10]))
											if v349 < v351 {
												v335 = v347
												v336 = v349
												continue
											} else {
												break
											}
											break
										}
										v353 = v347
									}
									*(*uint32)(unsafe.Add(mBase, _consts[885])) = uint32(v353)
								}
							}
							v356 = int32(2)
							v357 = int32(0)
							v358 = *(*int32)(unsafe.Add(mBase, _consts[39]))
							v363 = *(*int32)(unsafe.Add(mBase, _consts[806]))
							if v363 != 0 {
								v364 = base.B2i32(v358 != int32(-1))
							} else {
								v364 = v356
							}
							v366 = *(*int32)(unsafe.Add(mBase, _consts[46]))
							if v366 != 0 {
								v367 = v356
							} else {
								v367 = v364
							}
							*(*int32)(unsafe.Add(mBase, _consts[244])) = v367
							*(*int32)(unsafe.Add(mBase, _consts[295])) = v367
							v373 = int32(0)
							F_moduleFireServerEvent(m, int64(13), v373, v373)
							mBase = m.M
							v376 = m.ExcPending
							if v376 != 0 {
								return int32(0)
							} else {
								v378 = v61
								m.G0 = v9 + int32(144)
								return v378
							}
						} else {
							F_latencyAddSample(m, int32(_a1456), v286)
							mBase = m.M
							v301 = m.ExcPending
							if v301 != 0 {
								return int32(0)
							} else {
								if base.Ui32(int32(5)) < base.Ui32(l0) {
								} else {
									if int32(1)<<(uint(l0)%32)&int32(54) == int32(0) {
									} else {
										v310 = int32(0)
										v311 = int64(0)
										*(*int64)(unsafe.Add(mBase, _consts[851])) = v311
										*(*int32)(unsafe.Add(mBase, _consts[38])) = l0
										*(*int32)(unsafe.Add(mBase, _consts[39])) = v61
										*(*int64)(unsafe.Add(mBase, _consts[888])) = v311
										*(*int64)(unsafe.Add(mBase, _consts[889])) = v311
										*(*int32)(unsafe.Add(mBase, _consts[886])) = v310
										v331 = *(*int32)(unsafe.Add(mBase, _consts[10]))
										if v331 < int32(1) {
											v353 = v311
										} else {
											v335 = v311
											v336 = int32(0)
											for {
												v337 = F_dbHasNoKeys(m, v336)
												mBase = m.M
												if v337 != 0 {
													v347 = v335
												} else {
													v339 = *(*int32)(unsafe.Add(mBase, _consts[23]))
													v343 = *(*int32)(unsafe.Add(mBase, uint32(v339+v336<<(uint(int32(2))%32))))
													v344 = *(*int32)(unsafe.Add(mBase, uint32(v343)))
													v345 = F_kvstoreSize(m, v344)
													mBase = m.M
													v347 = v345 + v335
												}
												v349 = v336 + int32(1)
												v351 = *(*int32)(unsafe.Add(mBase, _consts[10]))
												if v349 < v351 {
													v335 = v347
													v336 = v349
													continue
												} else {
													break
												}
												break
											}
											v353 = v347
										}
										*(*uint32)(unsafe.Add(mBase, _consts[885])) = uint32(v353)
									}
								}
								v356 = int32(2)
								v357 = int32(0)
								v358 = *(*int32)(unsafe.Add(mBase, _consts[39]))
								v363 = *(*int32)(unsafe.Add(mBase, _consts[806]))
								if v363 != 0 {
									v364 = base.B2i32(v358 != int32(-1))
								} else {
									v364 = v356
								}
								v366 = *(*int32)(unsafe.Add(mBase, _consts[46]))
								if v366 != 0 {
									v367 = v356
								} else {
									v367 = v364
								}
								*(*int32)(unsafe.Add(mBase, _consts[244])) = v367
								*(*int32)(unsafe.Add(mBase, _consts[295])) = v367
								v373 = int32(0)
								F_moduleFireServerEvent(m, int64(13), v373, v373)
								mBase = m.M
								v376 = m.ExcPending
								if v376 != 0 {
									return int32(0)
								} else {
									v378 = v61
									m.G0 = v9 + int32(144)
									return v378
								}
							}
						}
					}
				}
			} else {
				*(*int32)(unsafe.Add(mBase, _consts[9])) = int32(7)
				v378 = v19
				m.G0 = v9 + int32(144)
				return v378
			}
		}
	}
}
func F_serverGitSHA1(m *base.Module) int32 {
	return int32(_a904)
}
func F_serverInitThreadAttribute(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
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
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
	v2 = int32(0)
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v10 = m.G0
	v11 = int32(48)
	v12 = v10 - v11
	m.G0 = v12
	v15 = v12 + int32(4)
	v17 = int32(44)
	v18 = F___memset(m, v15, v2, v17)
	mBase = m.M
	v22 = F___memcpy(m, l0, v15, v17)
	mBase = m.M
	F___acquire_ptc(m)
	mBase = m.M
	v25 = *(*int32)(unsafe.Add(mBase, _consts[1069]))
	*(*int32)(unsafe.Add(mBase, uint32(v22))) = v25
	v28 = *(*int32)(unsafe.Add(mBase, _consts[1070]))
	*(*int32)(unsafe.Add(mBase, uint32(v22)+4)) = v28
	F___release_ptc(m)
	mBase = m.M
	m.G0 = v12 + v11
	v37 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(v7+int32(12)))) = v37
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v7)+12))
	v41 = int32(1)
	if base.Ui32(v41) < base.Ui32(v40) {
		v44 = v40
	} else {
		v44 = v41
	}
	v48 = v44
	for {
		if base.Ui32(v48) < base.Ui32(int32(4194304)) {
			v48 = v48 << (uint(int32(1)) % 32)
			continue
		} else {
			break
		}
		break
	}
	*(*int32)(unsafe.Add(mBase, uint32(v7)+12)) = v48
	if base.Ui32(v48+int32(-1073743872)) < base.Ui32(int32(-1073741824)) {
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(l0))) = v48
		*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = int32(0)
	}
	m.G0 = v7 + int32(16)
	return
}
func F_serverIsSupervised(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v6 int32
	_ = v6
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
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v38 int32
	_ = v38
	var v45 int32
	_ = v45
	var v53 int32
	_ = v53
	var v57 int32
	_ = v57
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v88 int32
	_ = v88
	var v92 int32
	_ = v92
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
	var v103 int32
	_ = v103
	var v110 int32
	_ = v110
	var v118 int32
	_ = v118
	var v125 int32
	_ = v125
	var v128 int32
	_ = v128
	var v134 int32
	_ = v134
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v145 int32
	_ = v145
	var v149 int32
	_ = v149
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v160 int32
	_ = v160
	var v167 int32
	_ = v167
	var v175 int32
	_ = v175
	var v179 int32
	_ = v179
	var v186 int32
	_ = v186
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v197 int32
	_ = v197
	var v206 int32
	_ = v206
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v217 int32
	_ = v217
	var v219 int32
	_ = v219
	var v221 int32
	_ = v221
	var v225 int32
	_ = v225
	var v229 int32
	_ = v229
	var v231 int32
	_ = v231
	var v234 int32
	_ = v234
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v246 int32
	_ = v246
	var v253 int32
	_ = v253
	var v257 int32
	_ = v257
	var v260 int32
	_ = v260
	var v261 int32
	_ = v261
	var v284 int32
	_ = v284
	v2 = int32(0)
	switch l0 + int32(-1) {
	case 0:
		goto L8
	case 1:
		goto L4
	case 2:
		goto L7
	default:
		v284 = v2
		goto L1
	}
L1:
	;
	return v284
L2:
	;
	if int32(2) < v70 {
		goto L63
	} else {
		goto L64
	}
L3:
	;
	if int32(3) < v190 {
		v284 = v2
		goto L1
	} else {
		goto L61
	}
L4:
	;
	v189 = *(*int32)(unsafe.Add(mBase, _consts[15]))
	v190 = v189
	goto L3
L5:
	;
	v128 = int32(_a1470)
	v134 = F___strchrnul(m, v128, int32(61))
	mBase = m.M
	if v134 != v128 {
		goto L45
	} else {
		goto L46
	}
L6:
	;
	v71 = int32(_a1471)
	v77 = F___strchrnul(m, v71, int32(61))
	mBase = m.M
	if v77 != v71 {
		goto L28
	} else {
		goto L29
	}
L7:
	;
	v69 = *(*int32)(unsafe.Add(mBase, _consts[15]))
	v70 = v69
	goto L6
L8:
	;
	v6 = int32(_a1471)
	v12 = F___strchrnul(m, v6, int32(61))
	mBase = m.M
	if v12 != v6 {
		goto L10
	} else {
		goto L11
	}
L9:
	;
	if v53 == int32(0) {
		goto L5
	} else {
		goto L23
	}
L10:
	;
	v15 = int32(0)
	v16 = v12 - v6
	v18 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16)+uint32(_consts[1066]))))
	if v18 != 0 {
		v45 = v15
		goto L12
	} else {
		goto L13
	}
L11:
	;
	v53 = int32(0)
	goto L9
L12:
	;
	v53 = v45
	goto L9
L13:
	;
	v19 = int32(0)
	v20 = *(*int32)(unsafe.Add(mBase, _consts[1067]))
	if v20 == v19 {
		v45 = v15
		goto L12
	} else {
		goto L14
	}
L14:
	;
	v23 = *(*int32)(unsafe.Add(mBase, uint32(v20)))
	if v23 == int32(0) {
		v45 = v15
		goto L12
	} else {
		goto L15
	}
L15:
	;
	v27 = v20
	v30 = v23
	goto L17
L16:
	;
	v45 = v33 + int32(1)
	goto L12
L17:
	;
	v31 = F_strncmp(m, v6, v30, v16)
	mBase = m.M
	if v31 != 0 {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v27)+4))
	if v38 != 0 {
		v27 = v27 + int32(4)
		v30 = v38
		goto L17
	} else {
		goto L22
	}
L20:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v27)))
	v33 = v32 + v16
	v34 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v33))))
	if v34 == int32(61) {
		goto L16
	} else {
		goto L21
	}
L21:
	;
	goto L19
L22:
	;
	v45 = v15
	goto L12
L23:
	;
	v57 = *(*int32)(unsafe.Add(mBase, _consts[15]))
	if int32(1) < v57 {
		v70 = v57
		goto L6
	} else {
		goto L24
	}
L24:
	;
	F__serverLog(m, int32(1), int32(_a1472), int32(0))
	mBase = m.M
	v66 = m.ExcPending
	if v66 != 0 {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	return int32(0)
L26:
	;
	goto L7
L27:
	;
	if v118 != 0 {
		goto L2
	} else {
		goto L41
	}
L28:
	;
	v80 = int32(0)
	v81 = v77 - v71
	v83 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v81)+uint32(_consts[1066]))))
	if v83 != 0 {
		v110 = v80
		goto L30
	} else {
		goto L31
	}
L29:
	;
	v118 = int32(0)
	goto L27
L30:
	;
	v118 = v110
	goto L27
L31:
	;
	v84 = int32(0)
	v85 = *(*int32)(unsafe.Add(mBase, _consts[1067]))
	if v85 == v84 {
		v110 = v80
		goto L30
	} else {
		goto L32
	}
L32:
	;
	v88 = *(*int32)(unsafe.Add(mBase, uint32(v85)))
	if v88 == int32(0) {
		v110 = v80
		goto L30
	} else {
		goto L33
	}
L33:
	;
	v92 = v85
	v95 = v88
	goto L35
L34:
	;
	v110 = v98 + int32(1)
	goto L30
L35:
	;
	v96 = F_strncmp(m, v71, v95, v81)
	mBase = m.M
	if v96 != 0 {
		goto L37
	} else {
		goto L38
	}
L37:
	;
	v103 = *(*int32)(unsafe.Add(mBase, uint32(v92)+4))
	if v103 != 0 {
		v92 = v92 + int32(4)
		v95 = v103
		goto L35
	} else {
		goto L40
	}
L38:
	;
	v97 = *(*int32)(unsafe.Add(mBase, uint32(v92)))
	v98 = v97 + v81
	v99 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v98))))
	if v99 == int32(61) {
		goto L34
	} else {
		goto L39
	}
L39:
	;
	goto L37
L40:
	;
	v110 = v80
	goto L30
L41:
	;
	if int32(3) < v70 {
		v284 = v2
		goto L1
	} else {
		goto L42
	}
L42:
	;
	F__serverLog(m, int32(3), int32(_a1473), int32(0))
	mBase = m.M
	v125 = m.ExcPending
	if v125 != 0 {
		goto L25
	} else {
		goto L43
	}
L43:
	;
	return int32(0)
L44:
	;
	if v175 == int32(0) {
		v284 = v2
		goto L1
	} else {
		goto L58
	}
L45:
	;
	v137 = int32(0)
	v138 = v134 - v128
	v140 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v138)+uint32(_consts[1068]))))
	if v140 != 0 {
		v167 = v137
		goto L47
	} else {
		goto L48
	}
L46:
	;
	v175 = int32(0)
	goto L44
L47:
	;
	v175 = v167
	goto L44
L48:
	;
	v141 = int32(0)
	v142 = *(*int32)(unsafe.Add(mBase, _consts[1067]))
	if v142 == v141 {
		v167 = v137
		goto L47
	} else {
		goto L49
	}
L49:
	;
	v145 = *(*int32)(unsafe.Add(mBase, uint32(v142)))
	if v145 == int32(0) {
		v167 = v137
		goto L47
	} else {
		goto L50
	}
L50:
	;
	v149 = v142
	v152 = v145
	goto L52
L51:
	;
	v167 = v155 + int32(1)
	goto L47
L52:
	;
	v153 = F_strncmp(m, v128, v152, v138)
	mBase = m.M
	if v153 != 0 {
		goto L54
	} else {
		goto L55
	}
L54:
	;
	v160 = *(*int32)(unsafe.Add(mBase, uint32(v149)+4))
	if v160 != 0 {
		v149 = v149 + int32(4)
		v152 = v160
		goto L52
	} else {
		goto L57
	}
L55:
	;
	v154 = *(*int32)(unsafe.Add(mBase, uint32(v149)))
	v155 = v154 + v138
	v156 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v155))))
	if v156 == int32(61) {
		goto L51
	} else {
		goto L56
	}
L56:
	;
	goto L54
L57:
	;
	v167 = v137
	goto L47
L58:
	;
	v179 = *(*int32)(unsafe.Add(mBase, _consts[15]))
	if int32(1) < v179 {
		v190 = v179
		goto L3
	} else {
		goto L59
	}
L59:
	;
	F__serverLog(m, int32(1), int32(_a1474), int32(0))
	mBase = m.M
	v186 = m.ExcPending
	if v186 != 0 {
		goto L25
	} else {
		goto L60
	}
L60:
	;
	goto L4
L61:
	;
	F__serverLog(m, int32(3), int32(_a1475), int32(0))
	mBase = m.M
	v197 = m.ExcPending
	if v197 != 0 {
		goto L25
	} else {
		goto L62
	}
L62:
	;
	return int32(0)
L63:
	;
	v208 = F_raise(m, int32(19))
	mBase = m.M
	v209 = m.ExcPending
	if v209 != 0 {
		goto L25
	} else {
		goto L66
	}
L64:
	;
	F__serverLog(m, int32(2), int32(_a1476), int32(0))
	mBase = m.M
	v206 = m.ExcPending
	if v206 != 0 {
		goto L25
	} else {
		goto L65
	}
L65:
	;
	goto L63
L66:
	;
	v210 = int32(_a1471)
	v217 = F___strchrnul(m, v210, int32(61))
	mBase = m.M
	if v217 == v210 {
		goto L69
	} else {
		goto L70
	}
L67:
	;
	*(*int32)(unsafe.Add(mBase, _consts[533])) = int32(3)
	v284 = int32(1)
	goto L1
L68:
	;
	v229 = int32(0)
	v231 = *(*int32)(unsafe.Add(mBase, _consts[1067]))
	if v231 == v229 {
		goto L72
	} else {
		goto L73
	}
L69:
	;
	v225 = F___errno_location(m)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v225))) = int32(28)
	goto L67
L70:
	;
	v219 = v217 - v210
	v221 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v219)+uint32(_consts[1066]))))
	if v221 == int32(0) {
		goto L68
	} else {
		goto L71
	}
L71:
	;
	goto L69
L72:
	;
	goto L67
L73:
	;
	v234 = *(*int32)(unsafe.Add(mBase, uint32(v231)))
	if v234 == int32(0) {
		goto L72
	} else {
		goto L74
	}
L74:
	;
	v240 = v231
	v241 = v234
	v242 = v231
	goto L75
L75:
	;
	v243 = F_strncmp(m, v210, v241, v219)
	mBase = m.M
	if v243 != 0 {
		goto L78
	} else {
		goto L79
	}
L76:
	;
	if v257 == v260 {
		goto L72
	} else {
		goto L84
	}
L77:
	;
	v260 = v242 + int32(4)
	v261 = *(*int32)(unsafe.Add(mBase, uint32(v242)+4))
	if v261 != 0 {
		v240 = v257
		v241 = v261
		v242 = v260
		goto L75
	} else {
		goto L83
	}
L78:
	;
	if v240 == v242 {
		goto L81
	} else {
		goto L82
	}
L79:
	;
	v244 = *(*int32)(unsafe.Add(mBase, uint32(v242)))
	v246 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v244+v219))))
	if v246 != int32(61) {
		goto L78
	} else {
		goto L80
	}
L80:
	;
	F_dummy_5(m, v244, int32(0))
	mBase = m.M
	v257 = v240
	goto L77
L81:
	;
	v257 = v240 + int32(4)
	goto L77
L82:
	;
	v253 = *(*int32)(unsafe.Add(mBase, uint32(v242)))
	*(*int32)(unsafe.Add(mBase, uint32(v240))) = v253
	goto L81
L83:
	;
	goto L76
L84:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v257))) = int32(0)
	goto L72
}
func F_serverProcTitleGetVariable(m *base.Module, l0 int32, l1 int32) int32 {
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
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v41 int32
	_ = v41
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
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v128 int32
	_ = v128
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v176 int32
	_ = v176
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
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
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v214 int32
	_ = v214
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v231 int32
	_ = v231
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
	var v248 int32
	_ = v248
	var v249 int32
	_ = v249
	var v252 int32
	_ = v252
	var v253 int32
	_ = v253
	var v256 int32
	_ = v256
	var v265 int32
	_ = v265
	var v266 int32
	_ = v266
	var v270 int32
	_ = v270
	var v271 int32
	_ = v271
	var v273 int32
	_ = v273
	var v278 int32
	_ = v278
	var v279 int32
	_ = v279
	var v280 int32
	_ = v280
	var v281 int32
	_ = v281
	var v284 int32
	_ = v284
	var v285 int32
	_ = v285
	var v291 int32
	_ = v291
	var v292 int32
	_ = v292
	var v295 int32
	_ = v295
	var v296 int32
	_ = v296
	var v299 int32
	_ = v299
	var v308 int32
	_ = v308
	var v309 int32
	_ = v309
	var v314 int32
	_ = v314
	var v315 int32
	_ = v315
	var v316 int32
	_ = v316
	var v318 int32
	_ = v318
	v5 = m.G0
	v7 = v5 - int32(64)
	m.G0 = v7
	v9 = int32(_a1458)
	v12 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1059])))
	v13 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v13 == int32(0) {
		v36 = v12
		v37 = v13
		goto L4
	} else {
		goto L5
	}
L1:
	;
	m.G0 = v7 + int32(64)
	return v318
L2:
	;
	v45 = int32(_a1459)
	v48 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1060])))
	v49 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v49 == int32(0) {
		v72 = v48
		v73 = v49
		goto L16
	} else {
		goto L17
	}
L3:
	;
	if v37-v36&int32(255) != 0 {
		goto L2
	} else {
		goto L11
	}
L4:
	;
	goto L3
L5:
	;
	if v13 != v12&int32(255) {
		v36 = v12
		v37 = v13
		goto L4
	} else {
		goto L6
	}
L6:
	;
	v19 = l0
	v20 = v9
	goto L7
L7:
	;
	v23 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20)+1)))
	v24 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19)+1)))
	if v24 == int32(0) {
		v36 = v23
		v37 = v24
		goto L4
	} else {
		goto L9
	}
L8:
	;
	v36 = v23
	v37 = v24
	goto L4
L9:
	;
	v27 = int32(1)
	if v24 == v23&int32(255) {
		v19 = v19 + v27
		v20 = v20 + v27
		goto L7
	} else {
		goto L10
	}
L10:
	;
	goto L8
L11:
	;
	v41 = F_sdsnew(m, l1)
	mBase = m.M
	v44 = m.ExcPending
	if v44 != 0 {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	return int32(0)
L13:
	;
	v318 = v41
	goto L1
L14:
	;
	v110 = int32(_a1460)
	v113 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1061])))
	v114 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v114 == int32(0) {
		v137 = v113
		v138 = v114
		goto L37
	} else {
		goto L38
	}
L15:
	;
	if v73-v72&int32(255) != 0 {
		goto L14
	} else {
		goto L23
	}
L16:
	;
	goto L15
L17:
	;
	if v49 != v48&int32(255) {
		v72 = v48
		v73 = v49
		goto L16
	} else {
		goto L18
	}
L18:
	;
	v55 = l0
	v56 = v45
	goto L19
L19:
	;
	v59 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v56)+1)))
	v60 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v55)+1)))
	if v60 == int32(0) {
		v72 = v59
		v73 = v60
		goto L16
	} else {
		goto L21
	}
L20:
	;
	v72 = v59
	v73 = v60
	goto L16
L21:
	;
	v63 = int32(1)
	if v60 == v59&int32(255) {
		v55 = v55 + v63
		v56 = v56 + v63
		goto L19
	} else {
		goto L22
	}
L22:
	;
	goto L20
L23:
	;
	v77 = int32(0)
	v78 = *(*int32)(unsafe.Add(mBase, _consts[103]))
	v80 = *(*int32)(unsafe.Add(mBase, _consts[105]))
	v81 = F_sdsempty(m)
	mBase = m.M
	v82 = m.ExcPending
	if v82 != 0 {
		goto L12
	} else {
		goto L24
	}
L24:
	;
	if v78|v80 == int32(0) {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	v103 = *(*int32)(unsafe.Add(mBase, _consts[507]))
	*(*int32)(unsafe.Add(mBase, uint32(v7)+16)) = v103
	v108 = F_sdscatprintf(m, v81, int32(_a1461), v7+int32(16))
	mBase = m.M
	v109 = m.ExcPending
	if v109 != 0 {
		goto L12
	} else {
		goto L34
	}
L26:
	;
	v86 = int32(0)
	v87 = *(*int32)(unsafe.Add(mBase, _consts[785]))
	v90 = *(*int32)(unsafe.Add(mBase, _consts[192]))
	if v90 != 0 {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	v91 = v87
	goto L29
L28:
	;
	v91 = int32(_a1180)
	goto L29
L29:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v7))) = v91
	v93 = int32(0)
	v94 = *(*int32)(unsafe.Add(mBase, _consts[103]))
	v96 = *(*int32)(unsafe.Add(mBase, _consts[105]))
	if v94 != 0 {
		goto L30
	} else {
		goto L31
	}
L30:
	;
	v97 = v94
	goto L32
L31:
	;
	v97 = v96
	goto L32
L32:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v7)+4)) = v97
	v100 = F_sdscatprintf(m, v81, int32(_a1462), v7)
	mBase = m.M
	v101 = m.ExcPending
	if v101 != 0 {
		goto L12
	} else {
		goto L33
	}
L33:
	;
	v318 = v100
	goto L1
L34:
	;
	v318 = v108
	goto L1
L35:
	;
	v158 = int32(_a1463)
	v161 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1062])))
	v162 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v162 == int32(0) {
		v185 = v161
		v186 = v162
		goto L54
	} else {
		goto L55
	}
L36:
	;
	if v138-v137&int32(255) != 0 {
		goto L35
	} else {
		goto L44
	}
L37:
	;
	goto L36
L38:
	;
	if v114 != v113&int32(255) {
		v137 = v113
		v138 = v114
		goto L37
	} else {
		goto L39
	}
L39:
	;
	v120 = l0
	v121 = v110
	goto L40
L40:
	;
	v124 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v121)+1)))
	v125 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v120)+1)))
	if v125 == int32(0) {
		v137 = v124
		v138 = v125
		goto L37
	} else {
		goto L42
	}
L41:
	;
	v137 = v124
	v138 = v125
	goto L37
L42:
	;
	v128 = int32(1)
	if v125 == v124&int32(255) {
		v120 = v120 + v128
		v121 = v121 + v128
		goto L40
	} else {
		goto L43
	}
L43:
	;
	goto L41
L44:
	;
	v142 = int32(0)
	v143 = *(*int32)(unsafe.Add(mBase, _consts[63]))
	if v143 == v142 {
		goto L45
	} else {
		goto L46
	}
L45:
	;
	v149 = int32(0)
	v150 = *(*int32)(unsafe.Add(mBase, _consts[173]))
	if v150 == v149 {
		goto L48
	} else {
		goto L49
	}
L46:
	;
	v147 = F_sdsnew(m, int32(_a1464))
	mBase = m.M
	v148 = m.ExcPending
	if v148 != 0 {
		goto L12
	} else {
		goto L47
	}
L47:
	;
	v318 = v147
	goto L1
L48:
	;
	v156 = F_sdsempty(m)
	mBase = m.M
	v157 = m.ExcPending
	if v157 != 0 {
		goto L12
	} else {
		goto L51
	}
L49:
	;
	v154 = F_sdsnew(m, int32(_a1465))
	mBase = m.M
	v155 = m.ExcPending
	if v155 != 0 {
		goto L12
	} else {
		goto L50
	}
L50:
	;
	v318 = v154
	goto L1
L51:
	;
	v318 = v156
	goto L1
L52:
	;
	v196 = int32(_a1466)
	v199 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1063])))
	v200 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v200 == int32(0) {
		v223 = v199
		v224 = v200
		goto L68
	} else {
		goto L69
	}
L53:
	;
	if v186-v185&int32(255) != 0 {
		goto L52
	} else {
		goto L61
	}
L54:
	;
	goto L53
L55:
	;
	if v162 != v161&int32(255) {
		v185 = v161
		v186 = v162
		goto L54
	} else {
		goto L56
	}
L56:
	;
	v168 = l0
	v169 = v158
	goto L57
L57:
	;
	v172 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v169)+1)))
	v173 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v168)+1)))
	if v173 == int32(0) {
		v185 = v172
		v186 = v173
		goto L54
	} else {
		goto L59
	}
L58:
	;
	v185 = v172
	v186 = v173
	goto L54
L59:
	;
	v176 = int32(1)
	if v173 == v172&int32(255) {
		v168 = v168 + v176
		v169 = v169 + v176
		goto L57
	} else {
		goto L60
	}
L60:
	;
	goto L58
L61:
	;
	v191 = *(*int32)(unsafe.Add(mBase, _consts[178]))
	if v191 != 0 {
		goto L62
	} else {
		goto L63
	}
L62:
	;
	v193 = v191
	goto L64
L63:
	;
	v193 = int32(_a216)
	goto L64
L64:
	;
	v194 = F_sdsnew(m, v193)
	mBase = m.M
	v195 = m.ExcPending
	if v195 != 0 {
		goto L12
	} else {
		goto L65
	}
L65:
	;
	v318 = v194
	goto L1
L66:
	;
	v238 = int32(_a1467)
	v241 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1064])))
	v242 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v242 == int32(0) {
		v265 = v241
		v266 = v242
		goto L80
	} else {
		goto L81
	}
L67:
	;
	if v224-v223&int32(255) != 0 {
		goto L66
	} else {
		goto L75
	}
L68:
	;
	goto L67
L69:
	;
	if v200 != v199&int32(255) {
		v223 = v199
		v224 = v200
		goto L68
	} else {
		goto L70
	}
L70:
	;
	v206 = l0
	v207 = v196
	goto L71
L71:
	;
	v210 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v207)+1)))
	v211 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v206)+1)))
	if v211 == int32(0) {
		v223 = v210
		v224 = v211
		goto L68
	} else {
		goto L73
	}
L72:
	;
	v223 = v210
	v224 = v211
	goto L68
L73:
	;
	v214 = int32(1)
	if v211 == v210&int32(255) {
		v206 = v206 + v214
		v207 = v207 + v214
		goto L71
	} else {
		goto L74
	}
L74:
	;
	goto L72
L75:
	;
	v228 = F_sdsempty(m)
	mBase = m.M
	v229 = m.ExcPending
	if v229 != 0 {
		goto L12
	} else {
		goto L76
	}
L76:
	;
	v231 = *(*int32)(unsafe.Add(mBase, _consts[103]))
	*(*int32)(unsafe.Add(mBase, uint32(v7)+32)) = v231
	v236 = F_sdscatprintf(m, v228, int32(_a1468), v7+int32(32))
	mBase = m.M
	v237 = m.ExcPending
	if v237 != 0 {
		goto L12
	} else {
		goto L77
	}
L77:
	;
	v318 = v236
	goto L1
L78:
	;
	v280 = int32(0)
	v281 = int32(_a1469)
	v284 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1065])))
	v285 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v285 == v280 {
		v308 = v284
		v309 = v285
		goto L91
	} else {
		goto L92
	}
L79:
	;
	if v266-v265&int32(255) != 0 {
		goto L78
	} else {
		goto L87
	}
L80:
	;
	goto L79
L81:
	;
	if v242 != v241&int32(255) {
		v265 = v241
		v266 = v242
		goto L80
	} else {
		goto L82
	}
L82:
	;
	v248 = l0
	v249 = v238
	goto L83
L83:
	;
	v252 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v249)+1)))
	v253 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v248)+1)))
	if v253 == int32(0) {
		v265 = v252
		v266 = v253
		goto L80
	} else {
		goto L85
	}
L84:
	;
	v265 = v252
	v266 = v253
	goto L80
L85:
	;
	v256 = int32(1)
	if v253 == v252&int32(255) {
		v248 = v248 + v256
		v249 = v249 + v256
		goto L83
	} else {
		goto L86
	}
L86:
	;
	goto L84
L87:
	;
	v270 = F_sdsempty(m)
	mBase = m.M
	v271 = m.ExcPending
	if v271 != 0 {
		goto L12
	} else {
		goto L88
	}
L88:
	;
	v273 = *(*int32)(unsafe.Add(mBase, _consts[105]))
	*(*int32)(unsafe.Add(mBase, uint32(v7)+48)) = v273
	v278 = F_sdscatprintf(m, v270, int32(_a1468), v7+int32(48))
	mBase = m.M
	v279 = m.ExcPending
	if v279 != 0 {
		goto L12
	} else {
		goto L89
	}
L89:
	;
	v318 = v278
	goto L1
L90:
	;
	if v309-v308&int32(255) != 0 {
		v318 = v280
		goto L1
	} else {
		goto L98
	}
L91:
	;
	goto L90
L92:
	;
	if v285 != v284&int32(255) {
		v308 = v284
		v309 = v285
		goto L91
	} else {
		goto L93
	}
L93:
	;
	v291 = l0
	v292 = v281
	goto L94
L94:
	;
	v295 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v292)+1)))
	v296 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v291)+1)))
	if v296 == int32(0) {
		v308 = v295
		v309 = v296
		goto L91
	} else {
		goto L96
	}
L95:
	;
	v308 = v295
	v309 = v296
	goto L91
L96:
	;
	v299 = int32(1)
	if v296 == v295&int32(255) {
		v291 = v291 + v299
		v292 = v292 + v299
		goto L94
	} else {
		goto L97
	}
L97:
	;
	goto L95
L98:
	;
	v314 = *(*int32)(unsafe.Add(mBase, _consts[507]))
	v315 = F_sdsnew(m, v314)
	mBase = m.M
	v316 = m.ExcPending
	if v316 != 0 {
		goto L12
	} else {
		goto L99
	}
L99:
	;
	v318 = v315
	goto L1
}
func F_serverPubsubSubscriptionCount(m *base.Module) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v8 int64
	_ = v8
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v15 int64
	_ = v15
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	v4 = *(*int32)(unsafe.Add(mBase, _consts[131]))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(v4)+12))
	if v5 == int32(1) {
		v9 = *(*int32)(unsafe.Add(mBase, uint32(v4)+8))
		v10 = *(*int32)(unsafe.Add(mBase, uint32(v9)))
		if v10 != 0 {
			v12 = F_hashtableSize(m, v10)
			mBase = m.M
			v15 = base.I64_extend_i32_u(v12)
		} else {
			v15 = int64(0)
		}
	} else {
		v8 = *(*int64)(unsafe.Add(mBase, uint32(v4)+40))
		v15 = v8
	}
	v17 = *(*int32)(unsafe.Add(mBase, _consts[132]))
	v18 = *(*int32)(unsafe.Add(mBase, uint32(v17)+12))
	v21 = *(*int32)(unsafe.Add(mBase, uint32(v17)+16))
	return v18 + base.I32_wrap_i64(v15) + v21
}
func F_server_math_randomseed(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	v3 = F_luaL_checkinteger(m, l0, int32(1))
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		v7 = int32(0)
		*(*int32)(unsafe.Add(mBase, _consts[1191])) = int32(13070)
		*(*int32)(unsafe.Add(mBase, _consts[1192])) = v3 & int32(65535)
		*(*int32)(unsafe.Add(mBase, _consts[1193])) = int32(base.Ui32(v3) >> (uint(int32(16)) % 32))
		return int32(0)
	}
}
