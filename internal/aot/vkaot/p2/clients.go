package p2

import (
	base "github.com/shibukawa/valkeymem/internal/aot/vkaot/base"
	"unsafe"
)

func F_clientsTimeProc(m *base.Module, l0 int32, l1 int64, l2 int32) int64 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v18 int32
	_ = v18
	var v19 int64
	_ = v19
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
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
	var v36 float32
	_ = v36
	var v39 float32
	_ = v39
	var v42 float32
	_ = v42
	var v44 float64
	_ = v44
	var v50 int32
	_ = v50
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v62 int64
	_ = v62
	var v63 int32
	_ = v63
	var v64 int64
	_ = v64
	var v66 int64
	_ = v66
	var v67 int32
	_ = v67
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v89 int32
	_ = v89
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v116 int32
	_ = v116
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v122 int32
	_ = v122
	var v124 int32
	_ = v124
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
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
	var v144 int32
	_ = v144
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v162 int32
	_ = v162
	var v169 int32
	_ = v169
	var v173 int32
	_ = v173
	var v177 int32
	_ = v177
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v193 int32
	_ = v193
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v202 int32
	_ = v202
	var v215 int32
	_ = v215
	var v218 int32
	_ = v218
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v226 int64
	_ = v226
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v246 int64
	_ = v246
	var v247 int64
	_ = v247
	var v249 int32
	_ = v249
	var v250 int64
	_ = v250
	var v253 int64
	_ = v253
	var v254 int64
	_ = v254
	var v257 int32
	_ = v257
	var v265 int32
	_ = v265
	var v267 int32
	_ = v267
	var v268 int32
	_ = v268
	var v271 int32
	_ = v271
	var v274 int32
	_ = v274
	var v275 int32
	_ = v275
	var v277 int32
	_ = v277
	var v279 int32
	_ = v279
	var v280 int32
	_ = v280
	var v288 int32
	_ = v288
	var v289 int32
	_ = v289
	var v306 int32
	_ = v306
	var v308 int32
	_ = v308
	var v312 int32
	_ = v312
	var v313 int64
	_ = v313
	var v316 int64
	_ = v316
	v5 = int32(0)
	v18 = *(*int32)(unsafe.Add(mBase, _consts[23]))
	v19 = m.T0[v18].(func(*base.Module) int64)(m)
	mBase = m.M
	v21 = *(*int32)(unsafe.Add(mBase, _consts[3]))
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v21)+20))
	v23 = int32(5)
	if v22 < v23 {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	v61 = *(*int32)(unsafe.Add(mBase, _consts[394]))
	if v61 != 0 {
		goto L15
	} else {
		goto L16
	}
L2:
	;
	v55 = base.I32_div_s(int32(1000), v28)
	v56 = v32
	v59 = v55
	goto L1
L3:
	;
	v26 = v22
	goto L5
L4:
	;
	v26 = v23
	goto L5
L5:
	;
	v28 = *(*int32)(unsafe.Add(mBase, _consts[219]))
	v29 = base.I32_div_s(v22, v28)
	if v29 < int32(5) {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	v32 = v26
	goto L8
L7:
	;
	v32 = v29
	goto L8
L8:
	;
	if v32 < int32(201) {
		goto L2
	} else {
		goto L9
	}
L9:
	;
	v36 = float32(500)
	v39 = base.F32_div(base.F32_convert_i32_s(v22), float32(200))
	if base.F32_gt(v39, v36) != 0 {
		goto L11
	} else {
		goto L12
	}
L10:
	;
	v56 = int32(200)
	v59 = int32(-2147483648)
	goto L1
L11:
	;
	v42 = v36
	goto L13
L12:
	;
	v42 = v39
	goto L13
L13:
	;
	v44 = base.F64_div(float64(1000), base.F64_promote_f32(v42))
	if base.F64_lt(base.F64_abs(v44), float64(2.147483648e+09)) == int32(0) {
		goto L10
	} else {
		goto L14
	}
L14:
	;
	v50 = base.I32_trunc_f64_s(v44)
	v56 = int32(200)
	v59 = v50
	goto L1
L15:
	;
	v306 = int32(0)
	v308 = base.I32_div_s(int32(1000), v59)
	*(*int32)(unsafe.Add(mBase, _consts[803])) = v308
	v312 = *(*int32)(unsafe.Add(mBase, _consts[23]))
	v313 = m.T0[v312].(func(*base.Module) int64)(m)
	mBase = m.M
	v316 = *(*int64)(unsafe.Add(mBase, _consts[804]))
	*(*int64)(unsafe.Add(mBase, _consts[804])) = v313 - v19 + v316
	return base.I64_extend_i32_s(v59)
L16:
	;
	v62 = F_mstime(m)
	mBase = m.M
	v63 = int32(0)
	v64 = *(*int64)(unsafe.Add(mBase, _consts[47]))
	v66 = base.I64_rem_s(v64, int64(8))
	v67 = base.I32_wrap_i64(v66)
	v72 = base.I32_rem_s(base.I32_extend8_s(v67+int32(1)), int32(8))
	v74 = v72 << (uint(int32(2)) % 32)
	*(*int32)(unsafe.Add(mBase, uint32(v74)+uint32(_consts[805]))) = v63
	*(*int32)(unsafe.Add(mBase, uint32(v74)+uint32(_consts[806]))) = v63
	v84 = *(*int32)(unsafe.Add(mBase, _consts[3]))
	v85 = *(*int32)(unsafe.Add(mBase, uint32(v84)+20))
	if v85 == v63 {
		goto L15
	} else {
		goto L17
	}
L17:
	;
	v89 = v67 << (uint(int32(2)) % 32)
	v98 = v84
	v100 = v56
	goto L18
L18:
	;
	if v100 == int32(0) {
		goto L15
	} else {
		goto L20
	}
L19:
	;
	goto L15
L20:
	;
	v112 = *(*int32)(unsafe.Add(mBase, uint32(v98)))
	v113 = *(*int32)(unsafe.Add(mBase, uint32(v112)+8))
	v116 = *(*int32)(unsafe.Add(mBase, uint32(v98)+20))
	if base.Ui32(v116) < base.Ui32(int32(2)) {
		goto L22
	} else {
		goto L23
	}
L21:
	;
	v132 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v113)+222)))
	if v132 != 0 {
		goto L24
	} else {
		goto L25
	}
L22:
	;
	goto L21
L23:
	;
	v119 = *(*int32)(unsafe.Add(mBase, uint32(v98)))
	v120 = *(*int32)(unsafe.Add(mBase, uint32(v119)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v98))) = v120
	v122 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v120))) = v122
	v124 = *(*int32)(unsafe.Add(mBase, uint32(v98)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v124)+4)) = v119
	*(*int32)(unsafe.Add(mBase, uint32(v119))) = v124
	*(*int32)(unsafe.Add(mBase, uint32(v119)+4)) = v122
	*(*int32)(unsafe.Add(mBase, uint32(v98)+4)) = v119
	goto L22
L24:
	;
	v288 = *(*int32)(unsafe.Add(mBase, _consts[3]))
	v289 = *(*int32)(unsafe.Add(mBase, uint32(v288)+20))
	if v289 != 0 {
		v98 = v288
		v100 = v100 + int32(-1)
		goto L18
	} else {
		goto L72
	}
L25:
	;
	v133 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v113)+223)))
	if v133 != 0 {
		goto L24
	} else {
		goto L26
	}
L26:
	;
	v134 = F_clientsCronHandleTimeout(m, v113, v62)
	mBase = m.M
	v137 = m.ExcPending
	if v137 != 0 {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	return int64(0)
L28:
	;
	if v134 != 0 {
		goto L24
	} else {
		goto L29
	}
L29:
	;
	v138 = F_clientsCronResizeQueryBuffer(m, v113)
	mBase = m.M
	v139 = m.ExcPending
	if v139 != 0 {
		goto L27
	} else {
		goto L30
	}
L30:
	;
	v140 = F_clientsCronResizeOutputBuffer(m, v113, v62)
	mBase = m.M
	v141 = m.ExcPending
	if v141 != 0 {
		goto L27
	} else {
		goto L31
	}
L31:
	;
	v142 = int32(0)
	v144 = *(*int32)(unsafe.Add(mBase, uint32(v113)+12))
	if v144 == v142 {
		v187 = v142
		goto L32
	} else {
		goto L33
	}
L32:
	;
	v188 = *(*int32)(unsafe.Add(mBase, uint32(v113)+20))
	if v188 == int32(0) {
		v198 = v142
		goto L47
	} else {
		goto L48
	}
L33:
	;
	v153 = v144 + int32(-1)
	v154 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v153))))
	v156 = v154 & int32(7)
	switch v156 {
	case 0:
		goto L40
	case 1:
		v162 = int32(4)
		goto L35
	case 2:
		goto L39
	case 3:
		goto L38
	case 4:
		goto L37
	default:
		goto L36
	}
L34:
	;
	v187 = v186
	goto L32
L35:
	;
	switch v156 {
	case 0:
		goto L46
	case 1:
		goto L45
	case 2:
		goto L44
	case 3:
		goto L43
	case 4:
		goto L42
	default:
		v182 = int32(0)
		goto L41
	}
L36:
	;
	v162 = int32(1)
	goto L35
L37:
	;
	v162 = int32(18)
	goto L35
L38:
	;
	v162 = int32(10)
	goto L35
L39:
	;
	v162 = int32(6)
	goto L35
L40:
	;
	v157 = F_zmalloc_usable_size(m, v153)
	mBase = m.M
	v186 = v157
	goto L34
L41:
	;
	v186 = v162 + v182
	goto L34
L42:
	;
	v181 = *(*int32)(unsafe.Add(mBase, uint32(v144+int32(-9))))
	v182 = v181
	goto L41
L43:
	;
	v177 = *(*int32)(unsafe.Add(mBase, uint32(v144+int32(-5))))
	v186 = v162 + v177
	goto L34
L44:
	;
	v173 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v144+int32(-3)))))
	v186 = v162 + v173
	goto L34
L45:
	;
	v169 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v144+int32(-2)))))
	v186 = v162 + v169
	goto L34
L46:
	;
	v186 = v162 + int32(base.Ui32(v154)>>(uint(int32(3))%32))
	goto L34
L47:
	;
	v199 = *(*int32)(unsafe.Add(mBase, uint32(v113)+32))
	v202 = *(*int32)(unsafe.Add(mBase, uint32(v113)+200))
	if v202&int32(1) != 0 {
		goto L52
	} else {
		goto L53
	}
L48:
	;
	v193 = *(*int32)(unsafe.Add(mBase, uint32(v188+int32(-8))))
	goto L49
L49:
	;
	v198 = v193&int32(2147483647) + int32(8)
	goto L47
L50:
	;
	v267 = v199 + (v198 + v187)
	v268 = *(*int32)(unsafe.Add(mBase, uint32(v89)+uint32(_consts[805])))
	if base.Ui32(v267) <= base.Ui32(v268) {
		goto L63
	} else {
		goto L64
	}
L51:
	;
	v239 = *(*int32)(unsafe.Add(mBase, uint32(v113)+104))
	v240 = *(*int32)(unsafe.Add(mBase, uint32(v239)+184))
	if v240 != 0 {
		goto L61
	} else {
		goto L62
	}
L52:
	;
	v220 = *(*int32)(unsafe.Add(mBase, uint32(v113)+132))
	v221 = *(*int32)(unsafe.Add(mBase, uint32(v220)+20))
	v224 = *(*int32)(unsafe.Add(mBase, uint32(v113)+160))
	v225 = v221*int32(28) + v224
	v226 = *(*int64)(unsafe.Add(mBase, uint32(v113)+384))
	if v226 == int64(-1) {
		v236 = v225
		goto L59
	} else {
		goto L60
	}
L53:
	;
	if v202&int32(2) == int32(0) {
		goto L54
	} else {
		goto L55
	}
L54:
	;
	if v202&int32(262144) != 0 {
		goto L52
	} else {
		goto L57
	}
L55:
	;
	if v202&int32(4) == int32(0) {
		goto L51
	} else {
		goto L56
	}
L56:
	;
	goto L54
L57:
	;
	v215 = *(*int32)(unsafe.Add(mBase, uint32(v113)+216))
	if v215 == int32(0) {
		goto L52
	} else {
		goto L58
	}
L58:
	;
	v218 = F_isImportSlotMigrationJob(m, v215)
	mBase = m.M
	goto L52
L59:
	;
	v237 = *(*int32)(unsafe.Add(mBase, uint32(v113)+272))
	v265 = v237 + v236
	goto L50
L60:
	;
	v231 = *(*int32)(unsafe.Add(mBase, uint32(v113)+376))
	v232 = *(*int32)(unsafe.Add(mBase, uint32(v231)+20))
	v236 = v225 + base.I32_wrap_i64(v226) + v232*int32(28)
	goto L59
L61:
	;
	v243 = *(*int32)(unsafe.Add(mBase, _consts[217]))
	v244 = *(*int32)(unsafe.Add(mBase, uint32(v243)+4))
	v245 = *(*int32)(unsafe.Add(mBase, uint32(v244)+8))
	v246 = *(*int64)(unsafe.Add(mBase, uint32(v245)+16))
	v247 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v245)+24)))
	v249 = *(*int32)(unsafe.Add(mBase, uint32(v240)+8))
	v250 = *(*int64)(unsafe.Add(mBase, uint32(v249)+16))
	v253 = *(*int64)(unsafe.Add(mBase, uint32(v245)+8))
	v254 = *(*int64)(unsafe.Add(mBase, uint32(v249)+8))
	v257 = int32(44)
	v265 = base.I32_wrap_i64(v246+v247-v250) + base.I32_wrap_i64(v253-v254)*v257 + v257
	goto L50
L62:
	;
	v265 = int32(0)
	goto L50
L63:
	;
	v271 = *(*int32)(unsafe.Add(mBase, uint32(v89)+uint32(_consts[806])))
	if base.Ui32(v265) <= base.Ui32(v271) {
		goto L65
	} else {
		goto L66
	}
L64:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v89)+uint32(_consts[805]))) = v267
	goto L63
L65:
	;
	v274 = F_updateClientMemUsageAndBucket(m, v113)
	mBase = m.M
	v275 = m.ExcPending
	if v275 != 0 {
		goto L27
	} else {
		goto L68
	}
L66:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v89)+uint32(_consts[806]))) = v265
	goto L65
L67:
	;
	v279 = F_closeClientOnOutputBufferLimitReached(m, v113, int32(0))
	mBase = m.M
	v280 = m.ExcPending
	if v280 != 0 {
		goto L27
	} else {
		goto L71
	}
L68:
	;
	if v274 != 0 {
		goto L67
	} else {
		goto L69
	}
L69:
	;
	F_updateClientMemoryUsage(m, v113)
	mBase = m.M
	v277 = m.ExcPending
	if v277 != 0 {
		goto L27
	} else {
		goto L70
	}
L70:
	;
	goto L67
L71:
	;
	goto L24
L72:
	;
	goto L19
}
func F_handleClientsBlockedOnKeys(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
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
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
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
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v96 int32
	_ = v96
	var v99 int32
	_ = v99
	var v105 int32
	_ = v105
	var v109 int32
	_ = v109
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v130 int32
	_ = v130
	var v131 int64
	_ = v131
	var v133 int32
	_ = v133
	var v135 int32
	_ = v135
	var v138 int32
	_ = v138
	var v144 int32
	_ = v144
	var v149 int32
	_ = v149
	var v157 int32
	_ = v157
	var v170 int64
	_ = v170
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v200 int32
	_ = v200
	var v202 int32
	_ = v202
	var v205 int32
	_ = v205
	var v211 int32
	_ = v211
	var v215 int32
	_ = v215
	var v216 int64
	_ = v216
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v226 int32
	_ = v226
	var v228 int32
	_ = v228
	var v232 int32
	_ = v232
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v247 int32
	_ = v247
	var v250 int32
	_ = v250
	var v251 int32
	_ = v251
	var v254 int32
	_ = v254
	var v257 int32
	_ = v257
	var v258 int32
	_ = v258
	var v259 int32
	_ = v259
	var v262 int32
	_ = v262
	var v263 int32
	_ = v263
	var v264 int32
	_ = v264
	var v273 int32
	_ = v273
	var v274 int32
	_ = v274
	var v284 int32
	_ = v284
	var v285 int32
	_ = v285
	var v290 int32
	_ = v290
	var v294 int32
	_ = v294
	var v302 int64
	_ = v302
	var v304 int32
	_ = v304
	var v308 int64
	_ = v308
	var v312 int64
	_ = v312
	var v315 int32
	_ = v315
	var v323 int64
	_ = v323
	var v327 int32
	_ = v327
	var v328 int32
	_ = v328
	var v331 int32
	_ = v331
	var v333 int32
	_ = v333
	var v339 int32
	_ = v339
	var v347 int32
	_ = v347
	var v354 int32
	_ = v354
	var v355 int32
	_ = v355
	var v356 int32
	_ = v356
	var v357 int32
	_ = v357
	var v359 int32
	_ = v359
	var v364 int32
	_ = v364
	var v365 int32
	_ = v365
	var v371 int32
	_ = v371
	var v372 int32
	_ = v372
	var v376 int64
	_ = v376
	var v378 int32
	_ = v378
	var v379 int64
	_ = v379
	var v380 int32
	_ = v380
	var v381 int32
	_ = v381
	var v384 int32
	_ = v384
	var v386 int32
	_ = v386
	var v387 int64
	_ = v387
	var v391 int64
	_ = v391
	var v396 int32
	_ = v396
	var v398 int32
	_ = v398
	var v400 int32
	_ = v400
	var v437 int32
	_ = v437
	var v453 int32
	_ = v453
	var v455 int32
	_ = v455
	var v457 int32
	_ = v457
	var v459 int32
	_ = v459
	var v460 int32
	_ = v460
	var v477 int32
	_ = v477
	var v479 int32
	_ = v479
	var v482 int32
	_ = v482
	var v483 int32
	_ = v483
	var v498 int32
	_ = v498
	var v501 int32
	_ = v501
	var v523 int32
	_ = v523
	var v529 int32
	_ = v529
	v16 = m.G0
	v18 = v16 - int32(16)
	m.G0 = v18
	v21 = int32(*(*uint8)(unsafe.Add(mBase, _consts[90])))
	if v21 != 0 {
		v501 = v18
		goto L3
	} else {
		goto L4
	}
L1:
	;
	F__serverAssert(m, int32(_a198), int32(_a196), int32(740))
	mBase = m.M
	v529 = m.ExcPending
	if v529 != 0 {
		goto L10
	} else {
		goto L102
	}
L2:
	;
	F__serverAssert(m, int32(_a199), int32(_a196), int32(392))
	mBase = m.M
	v523 = m.ExcPending
	if v523 != 0 {
		goto L10
	} else {
		goto L101
	}
L3:
	;
	m.G0 = v501 + int32(16)
	return
L4:
	;
	v23 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _consts[90])) = uint8(v23)
	v26 = *(*int32)(unsafe.Add(mBase, _consts[91]))
	if v26 != 0 {
		goto L2
	} else {
		goto L5
	}
L5:
	;
	v28 = *(*int32)(unsafe.Add(mBase, _consts[92]))
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v28)+20))
	if v29 == int32(0) {
		v483 = v18
		goto L6
	} else {
		goto L7
	}
L6:
	;
	v498 = int32(0)
	*(*uint8)(unsafe.Add(mBase, _consts[90])) = uint8(v498)
	v501 = v483
	goto L3
L7:
	;
	v35 = v28
	v36 = v28 + int32(20)
	goto L8
L8:
	;
	v50 = F_listCreate(m)
	mBase = m.M
	v51 = m.ExcPending
	if v51 != 0 {
		goto L10
	} else {
		goto L11
	}
L9:
	;
	v483 = v18
	goto L6
L10:
	;
	return
L11:
	;
	*(*int32)(unsafe.Add(mBase, _consts[92])) = v50
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v35)+20))
	if v53 == int32(0) {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	F_listRelease(m, v35)
	mBase = m.M
	v477 = m.ExcPending
	if v477 != 0 {
		goto L10
	} else {
		goto L99
	}
L13:
	;
	goto L14
L14:
	;
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v35)))
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v71)+8))
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v72)))
	v74 = *(*int32)(unsafe.Add(mBase, uint32(v73)+20))
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v72)+4))
	v76 = F_dictDelete(m, v74, v75)
	mBase = m.M
	v77 = m.ExcPending
	if v77 != 0 {
		goto L10
	} else {
		goto L16
	}
L15:
	;
	goto L12
L16:
	;
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v72)))
	v79 = *(*int32)(unsafe.Add(mBase, uint32(v78)+12))
	v80 = *(*int32)(unsafe.Add(mBase, uint32(v72)+4))
	v81 = F_dictFind(m, v79, v80)
	mBase = m.M
	v82 = m.ExcPending
	if v82 != 0 {
		goto L10
	} else {
		goto L18
	}
L17:
	;
	v453 = *(*int32)(unsafe.Add(mBase, uint32(v72)+4))
	F_decrRefCount(m, v453)
	mBase = m.M
	v455 = m.ExcPending
	if v455 != 0 {
		goto L10
	} else {
		goto L95
	}
L18:
	;
	if v81 == int32(0) {
		goto L17
	} else {
		goto L19
	}
L19:
	;
	v85 = *(*int32)(unsafe.Add(mBase, uint32(v81)+8))
	goto L20
L20:
	;
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v85)+20))
	v89 = F_valkey_malloc(m, v86<<(uint(int32(3))%32))
	mBase = m.M
	v90 = m.ExcPending
	if v90 != 0 {
		goto L10
	} else {
		goto L21
	}
L21:
	;
	v91 = *(*int32)(unsafe.Add(mBase, uint32(v85)))
	*(*int32)(unsafe.Add(mBase, uint32(v18)+4)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v18))) = v91
	goto L22
L22:
	;
	v96 = *(*int32)(unsafe.Add(mBase, uint32(v18)))
	if v96 == int32(0) {
		goto L24
	} else {
		goto L25
	}
L23:
	;
	if v86 < int32(1) {
		goto L26
	} else {
		goto L27
	}
L24:
	;
	goto L23
L25:
	;
	v99 = *(*int32)(unsafe.Add(mBase, uint32(v18)+4))
	v105 = *(*int32)(unsafe.Add(mBase, uint32(v96+base.B2i32(v99 == int32(0))<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v18))) = v105
	goto L24
L26:
	;
	F_valkey_free(m, v89)
	mBase = m.M
	v437 = m.ExcPending
	if v437 != 0 {
		goto L10
	} else {
		goto L94
	}
L27:
	;
	v109 = int32(0)
	if v96 == v109 {
		goto L26
	} else {
		goto L28
	}
L28:
	;
	v120 = v96
	v121 = v109
	goto L30
L29:
	;
	v157 = v133
	goto L37
L30:
	;
	v130 = *(*int32)(unsafe.Add(mBase, uint32(v120)+8))
	v131 = *(*int64)(unsafe.Add(mBase, uint32(v130)))
	*(*int64)(unsafe.Add(mBase, uint32(v89+v121<<(uint(int32(3))%32)))) = v131
	v133 = int32(0)
	v135 = *(*int32)(unsafe.Add(mBase, uint32(v18)))
	if v135 == v133 {
		goto L33
	} else {
		goto L34
	}
L31:
	;
	goto L29
L32:
	;
	if v135 == int32(0) {
		goto L29
	} else {
		goto L35
	}
L33:
	;
	goto L32
L34:
	;
	v138 = *(*int32)(unsafe.Add(mBase, uint32(v18)+4))
	v144 = *(*int32)(unsafe.Add(mBase, uint32(v135+base.B2i32(v138 == int32(0))<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v18))) = v144
	goto L33
L35:
	;
	v149 = v121 + int32(1)
	if v149 < v86 {
		v120 = v135
		v121 = v149
		goto L30
	} else {
		goto L36
	}
L36:
	;
	goto L31
L37:
	;
	v170 = *(*int64)(unsafe.Add(mBase, uint32(v89+v157<<(uint(int32(3))%32))))
	v171 = *(*int32)(unsafe.Add(mBase, uint32(v72)))
	v172 = *(*int32)(unsafe.Add(mBase, uint32(v171)+12))
	v173 = *(*int32)(unsafe.Add(mBase, uint32(v72)+4))
	v174 = F_dictFetchValue(m, v172, v173)
	mBase = m.M
	v175 = m.ExcPending
	if v175 != 0 {
		goto L10
	} else {
		goto L40
	}
L38:
	;
	goto L26
L39:
	;
	if v157 != v121 {
		v157 = v157 + int32(1)
		goto L37
	} else {
		goto L93
	}
L40:
	;
	if v174 == int32(0) {
		goto L39
	} else {
		goto L41
	}
L41:
	;
	v179 = v18 + int32(8)
	v180 = *(*int32)(unsafe.Add(mBase, uint32(v174)))
	*(*int32)(unsafe.Add(mBase, uint32(v179)+4)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v179))) = v180
	goto L42
L42:
	;
	goto L43
L43:
	;
	v200 = v18 + int32(8)
	v202 = *(*int32)(unsafe.Add(mBase, uint32(v200)))
	if v202 == int32(0) {
		goto L46
	} else {
		goto L47
	}
L44:
	;
	v218 = *(*int32)(unsafe.Add(mBase, uint32(v72)))
	v219 = *(*int32)(unsafe.Add(mBase, uint32(v72)+4))
	v221 = F_lookupKeyReadWithFlags(m, v218, v219, int32(23))
	mBase = m.M
	v222 = m.ExcPending
	if v222 != 0 {
		goto L10
	} else {
		goto L50
	}
L45:
	;
	if v202 == int32(0) {
		goto L39
	} else {
		goto L48
	}
L46:
	;
	goto L45
L47:
	;
	v205 = *(*int32)(unsafe.Add(mBase, uint32(v200)+4))
	v211 = *(*int32)(unsafe.Add(mBase, uint32(v202+base.B2i32(v205 == int32(0))<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v200))) = v211
	goto L46
L48:
	;
	v215 = *(*int32)(unsafe.Add(mBase, uint32(v202)+8))
	v216 = *(*int64)(unsafe.Add(mBase, uint32(v215)))
	if v216 != v170 {
		goto L43
	} else {
		goto L49
	}
L49:
	;
	goto L44
L50:
	;
	v223 = *(*int32)(unsafe.Add(mBase, uint32(v215)+116))
	if v221 == int32(0) {
		goto L52
	} else {
		goto L53
	}
L51:
	;
	v254 = *(*int32)(unsafe.Add(mBase, uint32(v72)+4))
	if v251 == int32(3) {
		goto L59
	} else {
		goto L60
	}
L52:
	;
	v247 = *(*int32)(unsafe.Add(mBase, uint32(v223)+16))
	if v247 == int32(0) {
		goto L39
	} else {
		goto L58
	}
L53:
	;
	v226 = *(*int32)(unsafe.Add(mBase, uint32(v223)))
	v228 = *(*int32)(unsafe.Add(mBase, uint32(v221)))
	v232 = v228&int32(15) + int32(-1)
	if base.Ui32(int32(5)) < base.Ui32(v232) {
		v240 = int32(0)
		goto L54
	} else {
		goto L55
	}
L54:
	;
	if v226 == int32(3) {
		v251 = v226
		goto L51
	} else {
		goto L56
	}
L55:
	;
	v239 = *(*int32)(unsafe.Add(mBase, uint32(v232<<(uint(int32(2))%32))+uint32(_consts[93])))
	v240 = v239
	goto L54
L56:
	;
	if v226 == v240 {
		v251 = v226
		goto L51
	} else {
		goto L57
	}
L57:
	;
	goto L52
L58:
	;
	v250 = *(*int32)(unsafe.Add(mBase, uint32(v223)))
	v251 = v250
	goto L51
L59:
	;
	v371 = int32(_a44)
	v372 = *(*int32)(unsafe.Add(mBase, _consts[67]))
	*(*int32)(unsafe.Add(mBase, _consts[67])) = v215
	v376 = *(*int64)(unsafe.Add(mBase, _consts[94]))
	v378 = *(*int32)(unsafe.Add(mBase, _consts[23]))
	v379 = m.T0[v378].(func(*base.Module) int64)(m)
	mBase = m.M
	v380 = F_moduleTryServeClientBlockedOnKey(m, v215, v254)
	mBase = m.M
	v381 = m.ExcPending
	if v381 != 0 {
		goto L10
	} else {
		goto L88
	}
L60:
	;
	v257 = *(*int32)(unsafe.Add(mBase, uint32(v223)+24))
	v258 = F_dictFind(m, v257, v254)
	mBase = m.M
	v259 = m.ExcPending
	if v259 != 0 {
		goto L10
	} else {
		goto L61
	}
L61:
	;
	F_releaseBlockedEntry(m, v215, v258, int32(1))
	mBase = m.M
	v262 = m.ExcPending
	if v262 != 0 {
		goto L10
	} else {
		goto L62
	}
L62:
	;
	v263 = *(*int32)(unsafe.Add(mBase, uint32(v215)+116))
	v264 = *(*int32)(unsafe.Add(mBase, uint32(v263)))
	if base.Ui32(v264+int32(-4)) < base.Ui32(int32(2)) {
		goto L63
	} else {
		goto L64
	}
L63:
	;
	F_unblockClient(m, v215, int32(0))
	mBase = m.M
	v273 = m.ExcPending
	if v273 != 0 {
		goto L10
	} else {
		goto L66
	}
L64:
	;
	if v264 != int32(1) {
		goto L1
	} else {
		goto L65
	}
L65:
	;
	goto L63
L66:
	;
	v274 = *(*int32)(unsafe.Add(mBase, uint32(v215)+204))
	if v274&int32(2) == int32(0) {
		goto L39
	} else {
		goto L67
	}
L67:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v215)+204)) = v274&int32(-2097155) | int32(2097152)
	v284 = int32(_a44)
	v285 = *(*int32)(unsafe.Add(mBase, _consts[67]))
	*(*int32)(unsafe.Add(mBase, _consts[67])) = v215
	v290 = int32(0)
	v294 = *(*int32)(unsafe.Add(mBase, _consts[95]))
	*(*int32)(unsafe.Add(mBase, _consts[95])) = v294 + int32(1)
	goto L70
L68:
	;
	v327 = F_processCommandAndResetClient(m, v215)
	mBase = m.M
	v328 = m.ExcPending
	if v328 != 0 {
		goto L10
	} else {
		goto L75
	}
L69:
	;
	goto L68
L70:
	;
	if v294 != 0 {
		goto L69
	} else {
		goto L71
	}
L71:
	;
	goto L73
L72:
	;
	v304 = int32(0)
	*(*int64)(unsafe.Add(mBase, _consts[96])) = v302
	v308 = base.I64_div_s(v302, int64(1000))
	*(*int64)(unsafe.Add(mBase, _consts[35])) = v308
	v312 = base.I64_div_s(v302, int64(1000000))
	*(*int64)(unsafe.Add(mBase, _consts[47])) = v312
	v315 = *(*int32)(unsafe.Add(mBase, _consts[97]))
	F_lrulfu_updateClockAndPolicy(m, v308, int32(base.Ui32(v315&int32(2))>>(uint(int32(1))%32)))
	mBase = m.M
	v323 = *(*int64)(unsafe.Add(mBase, _consts[35]))
	*(*int64)(unsafe.Add(mBase, _consts[98])) = v323
	goto L69
L73:
	;
	v302 = F_ustime(m)
	mBase = m.M
	goto L72
L74:
	;
	v339 = *(*int32)(unsafe.Add(mBase, uint32(v215)+200))
	if v339&int32(16) != 0 {
		goto L78
	} else {
		goto L79
	}
L75:
	;
	if v327 != int32(-1) {
		goto L74
	} else {
		goto L76
	}
L76:
	;
	v331 = int32(0)
	v333 = *(*int32)(unsafe.Add(mBase, _consts[95]))
	*(*int32)(unsafe.Add(mBase, _consts[95])) = v333 + int32(-1)
	goto L77
L77:
	;
	*(*int32)(unsafe.Add(mBase, _consts[67])) = v285
	goto L39
L78:
	;
	v357 = int32(0)
	v359 = *(*int32)(unsafe.Add(mBase, _consts[95]))
	*(*int32)(unsafe.Add(mBase, _consts[95])) = v359 + int32(-1)
	goto L85
L79:
	;
	if v339&int32(1073741824) == int32(0) {
		goto L80
	} else {
		goto L81
	}
L80:
	;
	if v339&int32(128) != 0 {
		goto L78
	} else {
		goto L83
	}
L81:
	;
	F_moduleCallCommandUnblockedHandler(m, v215)
	mBase = m.M
	v347 = m.ExcPending
	if v347 != 0 {
		goto L10
	} else {
		goto L82
	}
L82:
	;
	goto L78
L83:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v215)+200)) = v339 | int32(128)
	v354 = *(*int32)(unsafe.Add(mBase, _consts[99]))
	v355 = F_listAddNodeTail(m, v354, v215)
	mBase = m.M
	v356 = m.ExcPending
	if v356 != 0 {
		goto L10
	} else {
		goto L84
	}
L84:
	;
	goto L78
L85:
	;
	F_afterCommand(m, v215)
	mBase = m.M
	v364 = m.ExcPending
	if v364 != 0 {
		goto L10
	} else {
		goto L86
	}
L86:
	;
	v365 = *(*int32)(unsafe.Add(mBase, uint32(v215)+204))
	*(*int32)(unsafe.Add(mBase, uint32(v215)+204)) = v365 & int32(-2097153)
	*(*int32)(unsafe.Add(mBase, _consts[67])) = v285
	goto L39
L87:
	;
	F_afterCommand(m, v215)
	mBase = m.M
	v400 = m.ExcPending
	if v400 != 0 {
		goto L10
	} else {
		goto L92
	}
L88:
	;
	if v380 == int32(0) {
		goto L87
	} else {
		goto L89
	}
L89:
	;
	v384 = int32(0)
	v386 = *(*int32)(unsafe.Add(mBase, _consts[23]))
	v387 = m.T0[v386].(func(*base.Module) int64)(m)
	mBase = m.M
	v391 = *(*int64)(unsafe.Add(mBase, _consts[94]))
	F_updateStatsOnUnblock(m, v215, v384, base.I32_wrap_i64(v387-v379), base.B2i32(v391 != v376)<<(uint(int32(1))%32))
	mBase = m.M
	v396 = m.ExcPending
	if v396 != 0 {
		goto L10
	} else {
		goto L90
	}
L90:
	;
	F_moduleUnblockClient(m, v215)
	mBase = m.M
	v398 = m.ExcPending
	if v398 != 0 {
		goto L10
	} else {
		goto L91
	}
L91:
	;
	goto L87
L92:
	;
	*(*int32)(unsafe.Add(mBase, _consts[67])) = v372
	goto L39
L93:
	;
	goto L38
L94:
	;
	goto L17
L95:
	;
	F_valkey_free(m, v72)
	mBase = m.M
	v457 = m.ExcPending
	if v457 != 0 {
		goto L10
	} else {
		goto L96
	}
L96:
	;
	F_listDelNode(m, v35, v71)
	mBase = m.M
	v459 = m.ExcPending
	if v459 != 0 {
		goto L10
	} else {
		goto L97
	}
L97:
	;
	v460 = *(*int32)(unsafe.Add(mBase, uint32(v36)))
	if v460 != 0 {
		goto L14
	} else {
		goto L98
	}
L98:
	;
	goto L15
L99:
	;
	v479 = *(*int32)(unsafe.Add(mBase, _consts[92]))
	v482 = *(*int32)(unsafe.Add(mBase, uint32(v479)+20))
	if v482 != 0 {
		v35 = v479
		v36 = v479 + int32(20)
		goto L8
	} else {
		goto L100
	}
L100:
	;
	goto L9
L101:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L102:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_handleClientsWithPendingWrites(m *base.Module) int32 {
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
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
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
	var v50 int32
	_ = v50
	var v57 int32
	_ = v57
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v96 int32
	_ = v96
	var v100 int32
	_ = v100
	var v111 int32
	_ = v111
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v11 = *(*int32)(unsafe.Add(mBase, _consts[216]))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(v11)+20))
	if v12 != 0 {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	F__serverAssert(m, int32(_a1021), int32(_a977), int32(3358))
	mBase = m.M
	v111 = m.ExcPending
	if v111 != 0 {
		goto L16
	} else {
		goto L32
	}
L2:
	;
	m.G0 = v8 + int32(16)
	return v100
L3:
	;
	v15 = v8 + int32(8)
	v16 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+4)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v15))) = v16
	goto L5
L4:
	;
	v100 = int32(0)
	goto L2
L5:
	;
	v20 = int32(0)
	v22 = v8 + int32(8)
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v22)))
	if v24 == v20 {
		goto L7
	} else {
		goto L8
	}
L6:
	;
	if v24 == int32(0) {
		v100 = v20
		goto L2
	} else {
		goto L9
	}
L7:
	;
	goto L6
L8:
	;
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v22)+4))
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v24+base.B2i32(v27 == int32(0))<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v22))) = v33
	goto L7
L9:
	;
	v38 = v24
	v39 = v20
	goto L10
L10:
	;
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v38)+8))
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v42)+200))
	if v43&int32(4194304) == int32(0) {
		goto L1
	} else {
		goto L12
	}
L11:
	;
	v100 = v83
	goto L2
L12:
	;
	if v43&int32(-2147482624) != 0 {
		v83 = v39
		goto L13
	} else {
		goto L14
	}
L13:
	;
	v85 = v8 + int32(8)
	v87 = *(*int32)(unsafe.Add(mBase, uint32(v85)))
	if v87 == int32(0) {
		goto L29
	} else {
		goto L30
	}
L14:
	;
	v50 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v42)+222)))
	if v50 == int32(1) {
		v83 = v39
		goto L13
	} else {
		goto L15
	}
L15:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v42)+200)) = v43 & int32(2143288319)
	v57 = *(*int32)(unsafe.Add(mBase, _consts[216]))
	F_listUnlinkNode(m, v57, v38)
	mBase = m.M
	v61 = m.ExcPending
	if v61 != 0 {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	return int32(0)
L17:
	;
	v62 = F_clientHasPendingReplies(m, v42)
	mBase = m.M
	v63 = m.ExcPending
	if v63 != 0 {
		goto L16
	} else {
		goto L18
	}
L18:
	;
	if v62 == int32(0) {
		v83 = v39
		goto L13
	} else {
		goto L19
	}
L19:
	;
	v66 = F_trySendWriteToIOThreads(m, v42)
	mBase = m.M
	v67 = m.ExcPending
	if v67 != 0 {
		goto L16
	} else {
		goto L20
	}
L20:
	;
	if v66 == int32(0) {
		v83 = v39
		goto L13
	} else {
		goto L21
	}
L21:
	;
	v70 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v42)+223)))
	if v70 != 0 {
		v83 = v39
		goto L13
	} else {
		goto L22
	}
L22:
	;
	v72 = v39 + int32(1)
	v73 = F_writeToClient(m, v42)
	mBase = m.M
	v74 = m.ExcPending
	if v74 != 0 {
		goto L16
	} else {
		goto L23
	}
L23:
	;
	if v73 == int32(-1) {
		v83 = v72
		goto L13
	} else {
		goto L24
	}
L24:
	;
	v77 = F_clientHasPendingReplies(m, v42)
	mBase = m.M
	v78 = m.ExcPending
	if v78 != 0 {
		goto L16
	} else {
		goto L25
	}
L25:
	;
	if v77 == int32(0) {
		v83 = v72
		goto L13
	} else {
		goto L26
	}
L26:
	;
	F_installClientWriteHandler(m, v42)
	mBase = m.M
	v82 = m.ExcPending
	if v82 != 0 {
		goto L16
	} else {
		goto L27
	}
L27:
	;
	v83 = v72
	goto L13
L28:
	;
	if v87 != 0 {
		v38 = v87
		v39 = v83
		goto L10
	} else {
		goto L31
	}
L29:
	;
	goto L28
L30:
	;
	v90 = *(*int32)(unsafe.Add(mBase, uint32(v85)+4))
	v96 = *(*int32)(unsafe.Add(mBase, uint32(v87+base.B2i32(v90 == int32(0))<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v85))) = v96
	goto L29
L31:
	;
	goto L11
L32:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_processClientsCommandsBatch(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v1 int32
	_ = v1
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v28 int32
	_ = v28
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v61 int32
	_ = v61
	var v67 int32
	_ = v67
	var v78 int32
	_ = v78
	var v87 int32
	_ = v87
	var v93 int32
	_ = v93
	var v105 int32
	_ = v105
	var v109 int32
	_ = v109
	var v115 int32
	_ = v115
	var v117 int32
	_ = v117
	var v121 int32
	_ = v121
	var v124 int32
	_ = v124
	var v133 int32
	_ = v133
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
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
	var v154 int32
	_ = v154
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v166 int32
	_ = v166
	var v169 int32
	_ = v169
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v181 int32
	_ = v181
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v196 int32
	_ = v196
	var v198 int64
	_ = v198
	var v202 int32
	_ = v202
	var v205 int32
	_ = v205
	var v209 int32
	_ = v209
	var v213 int32
	_ = v213
	var v216 int32
	_ = v216
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v232 int32
	_ = v232
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v243 int32
	_ = v243
	var v245 int32
	_ = v245
	var v247 int32
	_ = v247
	var v248 int32
	_ = v248
	var v251 int32
	_ = v251
	var v252 int32
	_ = v252
	var v255 int32
	_ = v255
	var v263 int32
	_ = v263
	var v264 int32
	_ = v264
	var v270 int32
	_ = v270
	var v276 int32
	_ = v276
	var v277 int32
	_ = v277
	var v282 int32
	_ = v282
	var v283 int32
	_ = v283
	var v288 int32
	_ = v288
	var v289 int32
	_ = v289
	var v290 int32
	_ = v290
	var v291 int32
	_ = v291
	var v294 int32
	_ = v294
	var v297 int32
	_ = v297
	var v298 int32
	_ = v298
	var v300 int32
	_ = v300
	var v301 int32
	_ = v301
	var v303 int32
	_ = v303
	var v307 int32
	_ = v307
	var v309 int64
	_ = v309
	var v313 int32
	_ = v313
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
	var v333 int32
	_ = v333
	var v335 int64
	_ = v335
	var v340 int32
	_ = v340
	var v341 int32
	_ = v341
	var v350 int32
	_ = v350
	var v353 int32
	_ = v353
	var v361 int32
	_ = v361
	var v366 int32
	_ = v366
	var v370 int32
	_ = v370
	var v374 int32
	_ = v374
	var v377 int32
	_ = v377
	var v378 int32
	_ = v378
	var v383 int32
	_ = v383
	var v387 int32
	_ = v387
	var v388 int32
	_ = v388
	var v392 int32
	_ = v392
	var v394 int32
	_ = v394
	var v396 int32
	_ = v396
	var v397 int32
	_ = v397
	var v400 int32
	_ = v400
	var v408 int64
	_ = v408
	var v416 int32
	_ = v416
	var v418 int32
	_ = v418
	var v421 int32
	_ = v421
	var v423 int32
	_ = v423
	v1 = int32(0)
	v10 = m.G0
	v12 = v10 - int32(16)
	m.G0 = v12
	v15 = *(*int32)(unsafe.Add(mBase, _consts[507]))
	if v15 == v1 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	m.G0 = v12 + int32(16)
	return
L2:
	;
	v18 = *(*int32)(unsafe.Add(mBase, uint32(v15)+12))
	if v18 == int32(0) {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v21 = *(*int32)(unsafe.Add(mBase, uint32(v15)+20))
	if v21 != 0 {
		v353 = v15
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v361 = *(*int32)(unsafe.Add(mBase, uint32(v353)+12))
	if v361 == int32(0) {
		v400 = v353
		goto L68
	} else {
		goto L69
	}
L5:
	;
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v15)+32))
	v28 = int32(0)
	goto L6
L6:
	;
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v22+v28<<(uint(int32(2))%32))))
	if v36 == int32(0) {
		goto L8
	} else {
		goto L9
	}
L7:
	;
	v109 = v15
	v115 = int32(0)
	goto L21
L8:
	;
	v105 = v28 + int32(1)
	if v105 != v18 {
		v28 = v105
		goto L6
	} else {
		goto L20
	}
L9:
	;
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v36)+24))
	if v39 < int32(2) {
		goto L8
	} else {
		goto L10
	}
L10:
	;
	v43 = v39 + int32(-1)
	v44 = int32(3)
	v45 = v43 & v44
	if base.Ui32(v39+int32(-2)) < base.Ui32(v44) {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v78 = int32(0)
	if v45 == v78 {
		goto L8
	} else {
		goto L16
	}
L12:
	;
	v61 = int32(0)
	goto L13
L13:
	;
	v67 = v61 + int32(4)
	if v67 != v43&int32(-4) {
		v61 = v67
		goto L13
	} else {
		goto L15
	}
L14:
	;
	goto L11
L15:
	;
	goto L14
L16:
	;
	v87 = v78
	goto L17
L17:
	;
	v93 = v87 + int32(1)
	if v93 != v45 {
		v87 = v93
		goto L17
	} else {
		goto L19
	}
L18:
	;
	goto L8
L19:
	;
	goto L18
L20:
	;
	goto L7
L21:
	;
	v117 = *(*int32)(unsafe.Add(mBase, uint32(v109)+32))
	v121 = *(*int32)(unsafe.Add(mBase, uint32(v117+v115<<(uint(int32(2))%32))))
	if v121 == int32(0) {
		v154 = v109
		goto L23
	} else {
		goto L24
	}
L22:
	;
	v166 = *(*int32)(unsafe.Add(mBase, uint32(v154)+8))
	if v166 == int32(0) {
		v353 = v154
		goto L4
	} else {
		goto L32
	}
L23:
	;
	v163 = v115 + int32(1)
	v164 = *(*int32)(unsafe.Add(mBase, uint32(v154)+12))
	if base.Ui32(v163) < base.Ui32(v164) {
		v109 = v154
		v115 = v163
		goto L21
	} else {
		goto L31
	}
L24:
	;
	v124 = *(*int32)(unsafe.Add(mBase, uint32(v121)+24))
	if v124 < int32(2) {
		v154 = v109
		goto L23
	} else {
		goto L25
	}
L25:
	;
	v133 = int32(1)
	v136 = v124
	goto L26
L26:
	;
	v137 = *(*int32)(unsafe.Add(mBase, uint32(v121)+20))
	v141 = *(*int32)(unsafe.Add(mBase, uint32(v137+v133<<(uint(int32(2))%32))))
	v142 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v141))))
	if v142&int32(240) != 0 {
		v147 = v136
		goto L28
	} else {
		goto L29
	}
L27:
	;
	v152 = *(*int32)(unsafe.Add(mBase, _consts[507]))
	v154 = v152
	goto L23
L28:
	;
	v149 = v133 + int32(1)
	if v149 < v147 {
		v133 = v149
		v136 = v147
		goto L26
	} else {
		goto L30
	}
L29:
	;
	v145 = F_objectGetVal(m, v141)
	mBase = m.M
	v146 = *(*int32)(unsafe.Add(mBase, uint32(v121)+24))
	v147 = v146
	goto L28
L30:
	;
	goto L27
L31:
	;
	goto L22
L32:
	;
	v169 = *(*int32)(unsafe.Add(mBase, uint32(v154)+28))
	v176 = int32(0)
	v177 = v169
	goto L33
L33:
	;
	v181 = v176 << (uint(int32(2)) % 32)
	v183 = *(*int32)(unsafe.Add(mBase, uint32(v177+v181)))
	v184 = F_objectGetVal(m, v183)
	mBase = m.M
	v186 = *(*int32)(unsafe.Add(mBase, _consts[507]))
	v187 = *(*int32)(unsafe.Add(mBase, uint32(v186)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v187+v181))) = v184
	v191 = v176 + int32(1)
	v192 = *(*int32)(unsafe.Add(mBase, uint32(v186)+8))
	if base.Ui32(v191) < base.Ui32(v192) {
		v176 = v191
		v177 = v187
		goto L33
	} else {
		goto L35
	}
L34:
	;
	if base.Ui32(v192) <= base.Ui32(int32(1)) {
		v353 = v186
		goto L4
	} else {
		goto L36
	}
L35:
	;
	goto L34
L36:
	;
	v196 = int32(_a44)
	v198 = *(*int64)(unsafe.Add(mBase, _consts[508]))
	*(*int64)(unsafe.Add(mBase, _consts[508])) = v198 + int64(1)
	v202 = *(*int32)(unsafe.Add(mBase, uint32(v186)+36))
	v205 = v186
	v209 = int32(0)
	goto L37
L37:
	;
	v213 = *(*int32)(unsafe.Add(mBase, uint32(v205)+40))
	v216 = v213 + v209*int32(48)
	v218 = v209 << (uint(int32(2)) % 32)
	v219 = v202 + v218
	v220 = *(*int32)(unsafe.Add(mBase, uint32(v219)))
	if v220 == int32(0) {
		v228 = v205
		goto L41
	} else {
		goto L42
	}
L38:
	;
	v255 = v248
	goto L48
L39:
	;
	v251 = v209 + int32(1)
	v252 = *(*int32)(unsafe.Add(mBase, uint32(v248)+8))
	if base.Ui32(v251) < base.Ui32(v252) {
		v205 = v248
		v209 = v251
		goto L37
	} else {
		goto L47
	}
L40:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v216))) = int32(0)
	v240 = *(*int32)(unsafe.Add(mBase, uint32(v219)))
	v241 = *(*int32)(unsafe.Add(mBase, uint32(v227)+28))
	v243 = *(*int32)(unsafe.Add(mBase, uint32(v241+v218)))
	F_hashtableIncrementalFindInit(m, v216+int32(8), v240, v243)
	mBase = m.M
	v245 = m.ExcPending
	if v245 != 0 {
		goto L45
	} else {
		goto L46
	}
L41:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v216))) = int32(2)
	v232 = *(*int32)(unsafe.Add(mBase, uint32(v228)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v228)+4)) = v232 + int32(1)
	v248 = v228
	goto L39
L42:
	;
	v223 = *(*int32)(unsafe.Add(mBase, uint32(v220)+20))
	v224 = *(*int32)(unsafe.Add(mBase, uint32(v220)+16))
	goto L43
L43:
	;
	v227 = *(*int32)(unsafe.Add(mBase, _consts[507]))
	if v223+v224 != 0 {
		goto L40
	} else {
		goto L44
	}
L44:
	;
	v228 = v227
	goto L41
L45:
	;
	return
L46:
	;
	v247 = *(*int32)(unsafe.Add(mBase, _consts[507]))
	v248 = v247
	goto L39
L47:
	;
	goto L38
L48:
	;
	v263 = *(*int32)(unsafe.Add(mBase, uint32(v255)+40))
	v264 = *(*int32)(unsafe.Add(mBase, uint32(v255)))
	v270 = v264
	goto L53
L49:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12))) = v277
	F__serverPanic_1(m, int32(_a910), int32(165), int32(_a911), v12)
	mBase = m.M
	v350 = m.ExcPending
	if v350 != 0 {
		goto L45
	} else {
		goto L67
	}
L50:
	;
	goto L49
L51:
	;
	v323 = F_hashtableIncrementalFindGetResult(m, v276+int32(8), v12+int32(12))
	mBase = m.M
	v324 = m.ExcPending
	if v324 != 0 {
		goto L45
	} else {
		goto L64
	}
L52:
	;
	v288 = F_hashtableIncrementalFindStep(m, v276+int32(8))
	mBase = m.M
	v289 = m.ExcPending
	if v289 != 0 {
		goto L45
	} else {
		goto L58
	}
L53:
	;
	v276 = v263 + v270*int32(48)
	v277 = *(*int32)(unsafe.Add(mBase, uint32(v276)))
	if v277 == int32(2) {
		goto L55
	} else {
		goto L56
	}
L55:
	;
	v282 = *(*int32)(unsafe.Add(mBase, uint32(v255)+8))
	v283 = base.I32_rem_u_s(v270+int32(1), v282)
	*(*int32)(unsafe.Add(mBase, uint32(v255))) = v283
	if v283 != v264 {
		v270 = v283
		goto L53
	} else {
		goto L57
	}
L56:
	;
	switch v277 {
	case 0:
		goto L52
	case 1:
		goto L51
	default:
		goto L50
	}
L57:
	;
	v353 = v255
	goto L4
L58:
	;
	v290 = int32(0)
	v291 = *(*int32)(unsafe.Add(mBase, _consts[507]))
	if v288 == v290 {
		goto L59
	} else {
		goto L60
	}
L59:
	;
	v300 = int32(_a44)
	v301 = *(*int32)(unsafe.Add(mBase, _consts[436]))
	v303 = *(*int32)(unsafe.Add(mBase, _consts[509]))
	if v301 < v303 {
		goto L61
	} else {
		goto L62
	}
L60:
	;
	v294 = *(*int32)(unsafe.Add(mBase, uint32(v291)))
	v297 = *(*int32)(unsafe.Add(mBase, uint32(v291)+8))
	v298 = base.I32_rem_u_s(v294+int32(1), v297)
	*(*int32)(unsafe.Add(mBase, uint32(v291))) = v298
	v255 = v291
	goto L48
L61:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v276))) = int32(1)
	v255 = v291
	goto L48
L62:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v276))) = int32(2)
	v307 = int32(_a44)
	v309 = *(*int64)(unsafe.Add(mBase, _consts[510]))
	*(*int64)(unsafe.Add(mBase, _consts[510])) = v309 + int64(1)
	v313 = *(*int32)(unsafe.Add(mBase, uint32(v291)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v291)+4)) = v313 + int32(1)
	v255 = v291
	goto L48
L63:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v276))) = int32(2)
	v333 = int32(_a44)
	v335 = *(*int64)(unsafe.Add(mBase, _consts[510]))
	*(*int64)(unsafe.Add(mBase, _consts[510])) = v335 + int64(1)
	v340 = *(*int32)(unsafe.Add(mBase, _consts[507]))
	v341 = *(*int32)(unsafe.Add(mBase, uint32(v340)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v340)+4)) = v341 + int32(1)
	v255 = v340
	goto L48
L64:
	;
	if v323 == int32(0) {
		goto L63
	} else {
		goto L65
	}
L65:
	;
	v327 = *(*int32)(unsafe.Add(mBase, uint32(v12)+12))
	v328 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v327))))
	if v328 != 0 {
		goto L63
	} else {
		goto L66
	}
L66:
	;
	v329 = F_objectGetVal(m, v327)
	mBase = m.M
	goto L63
L67:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L68:
	;
	v408 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v400))) = v408
	*(*int32)(unsafe.Add(mBase, uint32(v400)+20)) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v400+int32(8)))) = v408
	v416 = *(*int32)(unsafe.Add(mBase, uint32(v400)+16))
	v418 = *(*int32)(unsafe.Add(mBase, _consts[511]))
	if v416 == v418 {
		goto L1
	} else {
		goto L78
	}
L69:
	;
	v366 = v353
	v370 = int32(0)
	goto L70
L70:
	;
	v374 = *(*int32)(unsafe.Add(mBase, uint32(v366)+32))
	v377 = v374 + v370<<(uint(int32(2))%32)
	v378 = *(*int32)(unsafe.Add(mBase, uint32(v377)))
	if v378 == int32(0) {
		goto L72
	} else {
		goto L73
	}
L71:
	;
	v400 = v396
	goto L68
L72:
	;
	v394 = v370 + int32(1)
	v396 = *(*int32)(unsafe.Add(mBase, _consts[507]))
	v397 = *(*int32)(unsafe.Add(mBase, uint32(v396)+12))
	if base.Ui32(v394) < base.Ui32(v397) {
		v366 = v396
		v370 = v394
		goto L70
	} else {
		goto L77
	}
L73:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v377))) = int32(0)
	v383 = *(*int32)(unsafe.Add(mBase, uint32(v366)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v366)+20)) = v383 + int32(1)
	v387 = F_processPendingCommandAndInputBuffer(m, v378)
	mBase = m.M
	v388 = m.ExcPending
	if v388 != 0 {
		goto L45
	} else {
		goto L74
	}
L74:
	;
	if v387 == int32(-1) {
		goto L72
	} else {
		goto L75
	}
L75:
	;
	F_beforeNextClient(m, v378)
	mBase = m.M
	v392 = m.ExcPending
	if v392 != 0 {
		goto L45
	} else {
		goto L76
	}
L76:
	;
	goto L72
L77:
	;
	goto L71
L78:
	;
	F_freePrefetchCommandsBatch(m)
	mBase = m.M
	v421 = m.ExcPending
	if v421 != 0 {
		goto L45
	} else {
		goto L79
	}
L79:
	;
	F_prefetchCommandsBatchInit(m)
	mBase = m.M
	v423 = m.ExcPending
	if v423 != 0 {
		goto L45
	} else {
		goto L80
	}
L80:
	;
	goto L1
}
