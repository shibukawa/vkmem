package p1

import (
	base "github.com/shibukawa/vkmem/internal/aot/vkaot/base"
	"unsafe"
)

func F_zrangeCommand(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v13 int64
	_ = v13
	var v35 int32
	_ = v35
	v2 = int32(0)
	v3 = m.G0
	v5 = v3 - int32(48)
	m.G0 = v5
	*(*int32)(unsafe.Add(mBase, uint32(v5+int32(24)))) = v2
	v13 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v5+int32(16)))) = v13
	*(*int64)(unsafe.Add(mBase, uint32(v5+int32(8)))) = v13
	*(*int64)(unsafe.Add(mBase, uint32(v5))) = v13
	*(*int32)(unsafe.Add(mBase, uint32(v5)+40)) = int32(1101)
	*(*int32)(unsafe.Add(mBase, uint32(v5)+36)) = int32(1102)
	*(*int32)(unsafe.Add(mBase, uint32(v5)+32)) = int32(1103)
	*(*int32)(unsafe.Add(mBase, uint32(v5)+28)) = int32(1104)
	*(*int32)(unsafe.Add(mBase, uint32(v5)+4)) = l0
	F_zrangeGenericCommand(m, v5, int32(1), v2, v2, v2)
	mBase = m.M
	v35 = m.ExcPending
	if v35 != 0 {
		return
	} else {
		m.G0 = v5 + int32(48)
		return
	}
}
func F_zrangeGenericCommand(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v24 int32
	_ = v24
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v59 int32
	_ = v59
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v96 int32
	_ = v96
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
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v115 int32
	_ = v115
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
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
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
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v170 int32
	_ = v170
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v185 int32
	_ = v185
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v191 int32
	_ = v191
	var v195 int32
	_ = v195
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
	var v204 int32
	_ = v204
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v222 int32
	_ = v222
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v229 int32
	_ = v229
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v235 int32
	_ = v235
	var v239 int32
	_ = v239
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
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
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v257 int32
	_ = v257
	var v258 int32
	_ = v258
	var v261 int32
	_ = v261
	var v262 int32
	_ = v262
	var v263 int32
	_ = v263
	var v266 int32
	_ = v266
	var v267 int32
	_ = v267
	var v271 int32
	_ = v271
	var v272 int32
	_ = v272
	var v273 int32
	_ = v273
	var v276 int32
	_ = v276
	var v278 int32
	_ = v278
	var v279 int32
	_ = v279
	var v280 int32
	_ = v280
	var v282 int32
	_ = v282
	var v286 int32
	_ = v286
	var v288 int32
	_ = v288
	var v289 int32
	_ = v289
	var v290 int32
	_ = v290
	var v292 int32
	_ = v292
	var v293 int32
	_ = v293
	var v294 int32
	_ = v294
	var v295 int32
	_ = v295
	var v300 int32
	_ = v300
	var v301 int32
	_ = v301
	var v304 int32
	_ = v304
	var v305 int32
	_ = v305
	var v308 int32
	_ = v308
	var v309 int32
	_ = v309
	var v310 int32
	_ = v310
	var v316 int32
	_ = v316
	var v318 int32
	_ = v318
	var v319 int32
	_ = v319
	var v320 int32
	_ = v320
	var v321 int32
	_ = v321
	var v322 int32
	_ = v322
	var v326 int32
	_ = v326
	var v327 int32
	_ = v327
	var v329 int32
	_ = v329
	var v336 int32
	_ = v336
	var v337 int32
	_ = v337
	var v340 int32
	_ = v340
	var v342 int32
	_ = v342
	var v349 int32
	_ = v349
	var v353 int32
	_ = v353
	var v354 int32
	_ = v354
	var v358 int32
	_ = v358
	var v359 int32
	_ = v359
	var v361 int32
	_ = v361
	var v365 int32
	_ = v365
	var v366 int32
	_ = v366
	var v367 int32
	_ = v367
	var v374 int32
	_ = v374
	var v375 int32
	_ = v375
	var v376 int32
	_ = v376
	var v379 int32
	_ = v379
	var v383 int32
	_ = v383
	var v387 int32
	_ = v387
	var v388 int32
	_ = v388
	var v389 int32
	_ = v389
	var v393 int32
	_ = v393
	var v397 int32
	_ = v397
	var v398 int32
	_ = v398
	var v399 int32
	_ = v399
	var v400 int32
	_ = v400
	var v403 int32
	_ = v403
	var v407 int32
	_ = v407
	var v410 int32
	_ = v410
	var v415 int32
	_ = v415
	var v416 int32
	_ = v416
	var v417 int32
	_ = v417
	var v420 int32
	_ = v420
	var v424 int32
	_ = v424
	var v427 int32
	_ = v427
	var v428 int32
	_ = v428
	var v433 int32
	_ = v433
	var v435 int32
	_ = v435
	var v440 int32
	_ = v440
	var v441 int32
	_ = v441
	var v445 int32
	_ = v445
	var v446 int32
	_ = v446
	var v447 int32
	_ = v447
	var v451 int32
	_ = v451
	var v453 int32
	_ = v453
	var v455 int32
	_ = v455
	var v457 int32
	_ = v457
	var v459 int32
	_ = v459
	var v461 int32
	_ = v461
	var v463 int32
	_ = v463
	var v464 int32
	_ = v464
	var v467 int32
	_ = v467
	var v468 int32
	_ = v468
	var v474 int32
	_ = v474
	var v477 int32
	_ = v477
	var v478 int32
	_ = v478
	var v482 int32
	_ = v482
	var v487 int32
	_ = v487
	var v488 int32
	_ = v488
	var v492 int32
	_ = v492
	var v495 int32
	_ = v495
	var v496 int32
	_ = v496
	var v497 int32
	_ = v497
	var v499 int32
	_ = v499
	var v503 int32
	_ = v503
	var v504 int32
	_ = v504
	var v505 int32
	_ = v505
	var v507 int32
	_ = v507
	var v508 int32
	_ = v508
	var v509 int32
	_ = v509
	var v510 int32
	_ = v510
	var v514 int32
	_ = v514
	v6 = int32(0)
	v15 = m.G0
	v17 = v15 - int32(64)
	m.G0 = v17
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v20 = *(*int32)(unsafe.Add(mBase, uint32(v19)+20))
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v20+l1<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v17)+20)) = v6
	*(*int32)(unsafe.Add(mBase, uint32(v17)+16)) = v6
	*(*int32)(unsafe.Add(mBase, uint32(v17)+12)) = v6
	*(*int32)(unsafe.Add(mBase, uint32(v17)+8)) = int32(-1)
	v35 = l1 + int32(3)
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v19)+24))
	if v35 < v36 {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	m.G0 = v17 + int32(64)
	return
L2:
	;
	v365 = l1 + int32(1)
	v366 = int32(2)
	v367 = l1 + v366
	v374 = base.B2i32(v354 == v366) & base.B2i32(v353&int32(-2) == v366)
	if v374 != 0 {
		goto L103
	} else {
		goto L104
	}
L3:
	;
	v48 = l3
	v49 = l4
	v53 = int32(0)
	v54 = v35
	v55 = v36
	goto L8
L4:
	;
	v38 = int32(1)
	if base.Ui32(v38) < base.Ui32(l3) {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	v41 = l3
	goto L7
L6:
	;
	v41 = v38
	goto L7
L7:
	;
	v353 = l3
	v354 = l4
	v358 = v6
	v359 = v41
	v361 = base.B2i32(l3 == int32(3))
	goto L2
L8:
	;
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v19)+20))
	if l2 != 0 {
		v106 = v59
		goto L11
	} else {
		goto L12
	}
L9:
	;
	v329 = *(*int32)(unsafe.Add(mBase, uint32(v17)+8))
	if v329 == int32(-1) {
		goto L93
	} else {
		goto L94
	}
L10:
	;
	v326 = v322 + int32(1)
	v327 = *(*int32)(unsafe.Add(mBase, uint32(v19)+24))
	if v326 < v327 {
		v48 = v319
		v49 = v320
		v53 = v321
		v54 = v326
		v55 = v327
		goto L8
	} else {
		goto L92
	}
L11:
	;
	v108 = v54 << (uint(int32(2)) % 32)
	v110 = *(*int32)(unsafe.Add(mBase, uint32(v106+v108)))
	v111 = F_objectGetVal(m, v110)
	mBase = m.M
	v112 = int32(_a_F_zrangeGenericCommand_0)
	v115 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v111))))
	if v115 != 0 {
		goto L29
	} else {
		goto L30
	}
L12:
	;
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v59+v54<<(uint(int32(2))%32))))
	v64 = F_objectGetVal(m, v63)
	mBase = m.M
	v65 = int32(_a_F_zrangeGenericCommand_1)
	v68 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v64))))
	if v68 != 0 {
		goto L16
	} else {
		goto L17
	}
L13:
	;
	v105 = *(*int32)(unsafe.Add(mBase, uint32(v19)+20))
	v106 = v105
	goto L11
L14:
	;
	if v100-v102 != 0 {
		goto L13
	} else {
		goto L26
	}
L15:
	;
	v100 = F_tolower(m, v96)
	mBase = m.M
	v101 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v97))))
	v102 = F_tolower(m, v101)
	mBase = m.M
	goto L14
L16:
	;
	v70 = v64
	v71 = v65
	v72 = v68
	goto L19
L17:
	;
	v96 = int32(0)
	v97 = v65
	goto L15
L18:
	;
	v96 = v93 & int32(255)
	v97 = v92
	goto L15
L19:
	;
	v74 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v71))))
	if v74 == int32(0) {
		v92 = v71
		v93 = v72
		goto L18
	} else {
		goto L21
	}
L20:
	;
	v92 = v86
	v93 = int32(0)
	goto L18
L21:
	;
	v78 = v72 & int32(255)
	if v78 == v74 {
		goto L22
	} else {
		goto L23
	}
L22:
	;
	v85 = int32(1)
	v86 = v71 + v85
	v87 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v70)+1)))
	if v87 != 0 {
		v70 = v70 + v85
		v71 = v86
		v72 = v87
		goto L19
	} else {
		goto L25
	}
L23:
	;
	v80 = F_tolower(m, v78)
	mBase = m.M
	v81 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v71))))
	v82 = F_tolower(m, v81)
	mBase = m.M
	if v80 == v82 {
		goto L22
	} else {
		goto L24
	}
L24:
	;
	v84 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v70))))
	v92 = v71
	v93 = v84
	goto L18
L25:
	;
	goto L20
L26:
	;
	v319 = v48
	v320 = v49
	v321 = int32(1)
	v322 = v54
	goto L10
L27:
	;
	if int32(-3) < v54-v55 {
		goto L39
	} else {
		goto L40
	}
L28:
	;
	v147 = F_tolower(m, v143)
	mBase = m.M
	v148 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v144))))
	v149 = F_tolower(m, v148)
	mBase = m.M
	goto L27
L29:
	;
	v117 = v111
	v118 = v112
	v119 = v115
	goto L32
L30:
	;
	v143 = int32(0)
	v144 = v112
	goto L28
L31:
	;
	v143 = v140 & int32(255)
	v144 = v139
	goto L28
L32:
	;
	v121 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v118))))
	if v121 == int32(0) {
		v139 = v118
		v140 = v119
		goto L31
	} else {
		goto L34
	}
L33:
	;
	v139 = v133
	v140 = int32(0)
	goto L31
L34:
	;
	v125 = v119 & int32(255)
	if v125 == v121 {
		goto L35
	} else {
		goto L36
	}
L35:
	;
	v132 = int32(1)
	v133 = v118 + v132
	v134 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v117)+1)))
	if v134 != 0 {
		v117 = v117 + v132
		v118 = v133
		v119 = v134
		goto L32
	} else {
		goto L38
	}
L36:
	;
	v127 = F_tolower(m, v125)
	mBase = m.M
	v128 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v118))))
	v129 = F_tolower(m, v128)
	mBase = m.M
	if v127 == v129 {
		goto L35
	} else {
		goto L37
	}
L37:
	;
	v131 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v117))))
	v139 = v118
	v140 = v131
	goto L31
L38:
	;
	goto L33
L39:
	;
	if v49 != 0 {
		goto L47
	} else {
		goto L48
	}
L40:
	;
	if v147-v149 != 0 {
		goto L39
	} else {
		goto L41
	}
L41:
	;
	v154 = *(*int32)(unsafe.Add(mBase, uint32(v19)+20))
	v158 = *(*int32)(unsafe.Add(mBase, uint32(v154+v108+int32(4))))
	v162 = F_getLongFromObjectOrReply(m, v19, v158, v17+int32(12), int32(0))
	mBase = m.M
	v163 = m.ExcPending
	if v163 != 0 {
		goto L42
	} else {
		goto L43
	}
L42:
	;
	return
L43:
	;
	if v162 != 0 {
		goto L1
	} else {
		goto L44
	}
L44:
	;
	v164 = *(*int32)(unsafe.Add(mBase, uint32(v19)+20))
	v165 = int32(2)
	v166 = v54 + v165
	v170 = *(*int32)(unsafe.Add(mBase, uint32(v164+v166<<(uint(v165)%32))))
	v174 = F_getLongFromObjectOrReply(m, v19, v170, v17+int32(8), int32(0))
	mBase = m.M
	v175 = m.ExcPending
	if v175 != 0 {
		goto L42
	} else {
		goto L45
	}
L45:
	;
	if v174 != 0 {
		goto L1
	} else {
		goto L46
	}
L46:
	;
	v319 = v48
	v320 = v49
	v321 = v53
	v322 = v166
	goto L10
L47:
	;
	if v48 != 0 {
		goto L62
	} else {
		goto L63
	}
L48:
	;
	v176 = *(*int32)(unsafe.Add(mBase, uint32(v19)+20))
	v180 = *(*int32)(unsafe.Add(mBase, uint32(v176+v54<<(uint(int32(2))%32))))
	v181 = F_objectGetVal(m, v180)
	mBase = m.M
	v182 = int32(_a_F_zrangeGenericCommand_2)
	v185 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v181))))
	if v185 != 0 {
		goto L51
	} else {
		goto L52
	}
L49:
	;
	if v217-v219 != 0 {
		goto L47
	} else {
		goto L61
	}
L50:
	;
	v217 = F_tolower(m, v213)
	mBase = m.M
	v218 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v214))))
	v219 = F_tolower(m, v218)
	mBase = m.M
	goto L49
L51:
	;
	v187 = v181
	v188 = v182
	v189 = v185
	goto L54
L52:
	;
	v213 = int32(0)
	v214 = v182
	goto L50
L53:
	;
	v213 = v210 & int32(255)
	v214 = v209
	goto L50
L54:
	;
	v191 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v188))))
	if v191 == int32(0) {
		v209 = v188
		v210 = v189
		goto L53
	} else {
		goto L56
	}
L55:
	;
	v209 = v203
	v210 = int32(0)
	goto L53
L56:
	;
	v195 = v189 & int32(255)
	if v195 == v191 {
		goto L57
	} else {
		goto L58
	}
L57:
	;
	v202 = int32(1)
	v203 = v188 + v202
	v204 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v187)+1)))
	if v204 != 0 {
		v187 = v187 + v202
		v188 = v203
		v189 = v204
		goto L54
	} else {
		goto L60
	}
L58:
	;
	v197 = F_tolower(m, v195)
	mBase = m.M
	v198 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v188))))
	v199 = F_tolower(m, v198)
	mBase = m.M
	if v197 == v199 {
		goto L57
	} else {
		goto L59
	}
L59:
	;
	v201 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v187))))
	v209 = v188
	v210 = v201
	goto L53
L60:
	;
	goto L55
L61:
	;
	v319 = v48
	v320 = int32(2)
	v321 = v53
	v322 = v54
	goto L10
L62:
	;
	v316 = *(*int32)(unsafe.Add(mBase, _c_F_zrangeGenericCommand[0]))
	F_addReplyErrorObject(m, v19, v316)
	mBase = m.M
	v318 = m.ExcPending
	if v318 != 0 {
		goto L42
	} else {
		goto L91
	}
L63:
	;
	v222 = *(*int32)(unsafe.Add(mBase, uint32(v19)+20))
	v224 = *(*int32)(unsafe.Add(mBase, uint32(v222+v108)))
	v225 = F_objectGetVal(m, v224)
	mBase = m.M
	v226 = int32(_a_F_zrangeGenericCommand_3)
	v229 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v225))))
	if v229 != 0 {
		goto L67
	} else {
		goto L68
	}
L64:
	;
	v266 = int32(2)
	v267 = *(*int32)(unsafe.Add(mBase, uint32(v19)+20))
	v271 = *(*int32)(unsafe.Add(mBase, uint32(v267+v54<<(uint(v266)%32))))
	v272 = F_objectGetVal(m, v271)
	mBase = m.M
	v273 = int32(_a_F_zrangeGenericCommand_4)
	v276 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v272))))
	if v276 != 0 {
		goto L80
	} else {
		goto L81
	}
L65:
	;
	if v261-v263 != 0 {
		goto L64
	} else {
		goto L77
	}
L66:
	;
	v261 = F_tolower(m, v257)
	mBase = m.M
	v262 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v258))))
	v263 = F_tolower(m, v262)
	mBase = m.M
	goto L65
L67:
	;
	v231 = v225
	v232 = v226
	v233 = v229
	goto L70
L68:
	;
	v257 = int32(0)
	v258 = v226
	goto L66
L69:
	;
	v257 = v254 & int32(255)
	v258 = v253
	goto L66
L70:
	;
	v235 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v232))))
	if v235 == int32(0) {
		v253 = v232
		v254 = v233
		goto L69
	} else {
		goto L72
	}
L71:
	;
	v253 = v247
	v254 = int32(0)
	goto L69
L72:
	;
	v239 = v233 & int32(255)
	if v239 == v235 {
		goto L73
	} else {
		goto L74
	}
L73:
	;
	v246 = int32(1)
	v247 = v232 + v246
	v248 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v231)+1)))
	if v248 != 0 {
		v231 = v231 + v246
		v232 = v247
		v233 = v248
		goto L70
	} else {
		goto L76
	}
L74:
	;
	v241 = F_tolower(m, v239)
	mBase = m.M
	v242 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v232))))
	v243 = F_tolower(m, v242)
	mBase = m.M
	if v241 == v243 {
		goto L73
	} else {
		goto L75
	}
L75:
	;
	v245 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v231))))
	v253 = v232
	v254 = v245
	goto L69
L76:
	;
	goto L71
L77:
	;
	v319 = int32(3)
	v320 = v49
	v321 = v53
	v322 = v54
	goto L10
L78:
	;
	if v308-v310 == int32(0) {
		v319 = v266
		v320 = v49
		v321 = v53
		v322 = v54
		goto L10
	} else {
		goto L90
	}
L79:
	;
	v308 = F_tolower(m, v304)
	mBase = m.M
	v309 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v305))))
	v310 = F_tolower(m, v309)
	mBase = m.M
	goto L78
L80:
	;
	v278 = v272
	v279 = v273
	v280 = v276
	goto L83
L81:
	;
	v304 = int32(0)
	v305 = v273
	goto L79
L82:
	;
	v304 = v301 & int32(255)
	v305 = v300
	goto L79
L83:
	;
	v282 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v279))))
	if v282 == int32(0) {
		v300 = v279
		v301 = v280
		goto L82
	} else {
		goto L85
	}
L84:
	;
	v300 = v294
	v301 = int32(0)
	goto L82
L85:
	;
	v286 = v280 & int32(255)
	if v286 == v282 {
		goto L86
	} else {
		goto L87
	}
L86:
	;
	v293 = int32(1)
	v294 = v279 + v293
	v295 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v278)+1)))
	if v295 != 0 {
		v278 = v278 + v293
		v279 = v294
		v280 = v295
		goto L83
	} else {
		goto L89
	}
L87:
	;
	v288 = F_tolower(m, v286)
	mBase = m.M
	v289 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v279))))
	v290 = F_tolower(m, v289)
	mBase = m.M
	if v288 == v290 {
		goto L86
	} else {
		goto L88
	}
L88:
	;
	v292 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v278))))
	v300 = v279
	v301 = v292
	goto L82
L89:
	;
	goto L84
L90:
	;
	goto L62
L91:
	;
	goto L1
L92:
	;
	goto L9
L93:
	;
	v337 = int32(1)
	if base.Ui32(v337) < base.Ui32(v319) {
		goto L97
	} else {
		goto L98
	}
L94:
	;
	if base.Ui32(int32(1)) < base.Ui32(v319) {
		goto L93
	} else {
		goto L95
	}
L95:
	;
	F_addReplyError(m, v19, int32(_a_F_zrangeGenericCommand_5))
	mBase = m.M
	v336 = m.ExcPending
	if v336 != 0 {
		goto L42
	} else {
		goto L96
	}
L96:
	;
	goto L1
L97:
	;
	v340 = v319
	goto L99
L98:
	;
	v340 = v337
	goto L99
L99:
	;
	v342 = base.B2i32(v319 == int32(3))
	if v321 == int32(0) {
		v353 = v319
		v354 = v320
		v358 = v321
		v359 = v340
		v361 = v342
		goto L2
	} else {
		goto L100
	}
L100:
	;
	if v319 != int32(3) {
		v353 = v319
		v354 = v320
		v358 = v321
		v359 = v340
		v361 = v342
		goto L2
	} else {
		goto L101
	}
L101:
	;
	F_addReplyError(m, v19, int32(_a_F_zrangeGenericCommand_6))
	mBase = m.M
	v349 = m.ExcPending
	if v349 != 0 {
		goto L42
	} else {
		goto L102
	}
L102:
	;
	goto L1
L103:
	;
	v375 = v365
	goto L105
L104:
	;
	v375 = v367
	goto L105
L105:
	;
	if v374 != 0 {
		goto L106
	} else {
		goto L107
	}
L106:
	;
	v376 = v367
	goto L108
L107:
	;
	v376 = v365
	goto L108
L108:
	;
	switch v359 + int32(-1) {
	case 0:
		goto L112
	case 1:
		goto L111
	case 2:
		goto L110
	default:
		goto L109
	}
L109:
	;
	v435 = v358 | l2
	if v435 == int32(0) {
		goto L122
	} else {
		goto L123
	}
L110:
	;
	v416 = *(*int32)(unsafe.Add(mBase, uint32(v19)+20))
	v417 = int32(2)
	v420 = *(*int32)(unsafe.Add(mBase, uint32(v416+v376<<(uint(v417)%32))))
	v424 = *(*int32)(unsafe.Add(mBase, uint32(v416+v375<<(uint(v417)%32))))
	v427 = F_zsetParseLexRange(m, v420, v424, v17+int32(24))
	mBase = m.M
	v428 = m.ExcPending
	if v428 != 0 {
		goto L42
	} else {
		goto L119
	}
L111:
	;
	v399 = *(*int32)(unsafe.Add(mBase, uint32(v19)+20))
	v400 = int32(2)
	v403 = *(*int32)(unsafe.Add(mBase, uint32(v399+v376<<(uint(v400)%32))))
	v407 = *(*int32)(unsafe.Add(mBase, uint32(v399+v375<<(uint(v400)%32))))
	v410 = F_zslParseRange(m, v403, v407, v17+int32(40))
	mBase = m.M
	if v410 == int32(0) {
		goto L109
	} else {
		goto L117
	}
L112:
	;
	v379 = *(*int32)(unsafe.Add(mBase, uint32(v19)+20))
	v383 = *(*int32)(unsafe.Add(mBase, uint32(v379+v376<<(uint(int32(2))%32))))
	v387 = F_getLongFromObjectOrReply(m, v19, v383, v17+int32(20), int32(0))
	mBase = m.M
	v388 = m.ExcPending
	if v388 != 0 {
		goto L42
	} else {
		goto L113
	}
L113:
	;
	if v387 != 0 {
		goto L1
	} else {
		goto L114
	}
L114:
	;
	v389 = *(*int32)(unsafe.Add(mBase, uint32(v19)+20))
	v393 = *(*int32)(unsafe.Add(mBase, uint32(v389+v375<<(uint(int32(2))%32))))
	v397 = F_getLongFromObjectOrReply(m, v19, v393, v17+int32(16), int32(0))
	mBase = m.M
	v398 = m.ExcPending
	if v398 != 0 {
		goto L42
	} else {
		goto L115
	}
L115:
	;
	if v397 != 0 {
		goto L1
	} else {
		goto L116
	}
L116:
	;
	goto L109
L117:
	;
	F_addReplyError(m, v19, int32(_a_F_zrangeGenericCommand_7))
	mBase = m.M
	v415 = m.ExcPending
	if v415 != 0 {
		goto L42
	} else {
		goto L118
	}
L118:
	;
	goto L1
L119:
	;
	if v427 == int32(0) {
		goto L109
	} else {
		goto L120
	}
L120:
	;
	F_addReplyError(m, v19, int32(_a_F_zrangeGenericCommand_8))
	mBase = m.M
	v433 = m.ExcPending
	if v433 != 0 {
		goto L42
	} else {
		goto L121
	}
L121:
	;
	goto L1
L122:
	;
	v445 = *(*int32)(unsafe.Add(mBase, uint32(v19)+96))
	v446 = F_lookupKeyRead(m, v445, v24)
	mBase = m.M
	v447 = m.ExcPending
	if v447 != 0 {
		goto L42
	} else {
		goto L126
	}
L123:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = int32(1)
	v440 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v441 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v440)+224)))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = base.B2i32(base.Ui32(int32(2)) < base.Ui32(v441))
	goto L122
L124:
	;
	if v361 == int32(0) {
		goto L1
	} else {
		goto L141
	}
L125:
	;
	v463 = F_checkType(m, v19, v446, int32(3))
	mBase = m.M
	v464 = m.ExcPending
	if v464 != 0 {
		goto L42
	} else {
		goto L133
	}
L126:
	;
	if v446 != 0 {
		goto L125
	} else {
		goto L127
	}
L127:
	;
	if l2 == int32(0) {
		goto L128
	} else {
		goto L129
	}
L128:
	;
	v459 = *(*int32)(unsafe.Add(mBase, _c_F_zrangeGenericCommand[1]))
	F_addReply(m, v19, v459)
	mBase = m.M
	v461 = m.ExcPending
	if v461 != 0 {
		goto L42
	} else {
		goto L132
	}
L129:
	;
	v451 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	m.T0[v451].(func(*base.Module, int32, int32))(m, l0, int32(-1))
	mBase = m.M
	v453 = m.ExcPending
	if v453 != 0 {
		goto L42
	} else {
		goto L130
	}
L130:
	;
	v455 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	m.T0[v455].(func(*base.Module, int32, int32))(m, l0, int32(0))
	mBase = m.M
	v457 = m.ExcPending
	if v457 != 0 {
		goto L42
	} else {
		goto L131
	}
L131:
	;
	goto L124
L132:
	;
	goto L124
L133:
	;
	if v463 != 0 {
		goto L124
	} else {
		goto L134
	}
L134:
	;
	switch v359 + int32(-1) {
	case 0:
		goto L137
	case 1:
		goto L136
	case 2:
		goto L135
	default:
		goto L124
	}
L135:
	;
	v487 = *(*int32)(unsafe.Add(mBase, uint32(v17)+12))
	v488 = *(*int32)(unsafe.Add(mBase, uint32(v17)+8))
	F_genericZrangebylexCommand(m, l0, v17+int32(24), v446, base.B2i32(v435 != int32(0)), v487, v488, base.B2i32(v354 == int32(2)))
	mBase = m.M
	v492 = m.ExcPending
	if v492 != 0 {
		goto L42
	} else {
		goto L140
	}
L136:
	;
	v477 = *(*int32)(unsafe.Add(mBase, uint32(v17)+12))
	v478 = *(*int32)(unsafe.Add(mBase, uint32(v17)+8))
	F_genericZrangebyscoreCommand(m, l0, v17+int32(40), v446, v477, v478, base.B2i32(v354 == int32(2)))
	mBase = m.M
	v482 = m.ExcPending
	if v482 != 0 {
		goto L42
	} else {
		goto L139
	}
L137:
	;
	v467 = *(*int32)(unsafe.Add(mBase, uint32(v17)+20))
	v468 = *(*int32)(unsafe.Add(mBase, uint32(v17)+16))
	F_genericZrangebyrankCommand(m, l0, v446, v467, v468, base.B2i32(v435 != int32(0)), base.B2i32(v354 == int32(2)))
	mBase = m.M
	v474 = m.ExcPending
	if v474 != 0 {
		goto L42
	} else {
		goto L138
	}
L138:
	;
	goto L124
L139:
	;
	goto L124
L140:
	;
	goto L124
L141:
	;
	v495 = int32(_a_F_zrangeGenericCommand_9)
	v496 = *(*int32)(unsafe.Add(mBase, _c_F_zrangeGenericCommand[2]))
	v497 = *(*int32)(unsafe.Add(mBase, uint32(v17)+24))
	v499 = *(*int32)(unsafe.Add(mBase, _c_F_zrangeGenericCommand[3]))
	if v497 == v499 {
		v508 = v499
		v509 = v496
		goto L142
	} else {
		goto L143
	}
L142:
	;
	v510 = *(*int32)(unsafe.Add(mBase, uint32(v17)+28))
	if v510 == v508 {
		goto L1
	} else {
		goto L146
	}
L143:
	;
	if v497 == v496 {
		v508 = v499
		v509 = v496
		goto L142
	} else {
		goto L144
	}
L144:
	;
	F_sdsfree(m, v497)
	mBase = m.M
	v503 = m.ExcPending
	if v503 != 0 {
		goto L42
	} else {
		goto L145
	}
L145:
	;
	v504 = int32(_a_F_zrangeGenericCommand_9)
	v505 = *(*int32)(unsafe.Add(mBase, _c_F_zrangeGenericCommand[2]))
	v507 = *(*int32)(unsafe.Add(mBase, _c_F_zrangeGenericCommand[3]))
	v508 = v507
	v509 = v505
	goto L142
L146:
	;
	if v510 == v509 {
		goto L1
	} else {
		goto L147
	}
L147:
	;
	F_sdsfree(m, v510)
	mBase = m.M
	v514 = m.ExcPending
	if v514 != 0 {
		goto L42
	} else {
		goto L148
	}
L148:
	;
	goto L1
}
func F_zrangeResultBeginClient(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	if l1 < int32(1) {
		v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		v20 = F_addReplyDeferredLen(m, v19)
		mBase = m.M
		v21 = m.ExcPending
		if v21 != 0 {
			return
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v20
			return
		}
	} else {
		v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
		if v7 == int32(0) {
			v14 = l1
		} else {
			v10 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6)+224)))
			v14 = l1 << (uint(base.B2i32(v10 == int32(2))) % 32)
		}
		F_addReplyArrayLen(m, v6, v14)
		mBase = m.M
		v16 = m.ExcPending
		if v16 != 0 {
			return
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = int32(0)
			return
		}
	}
}
func F_zrangeResultEmitCBufferToClient(m *base.Module, l0 int32, l1 int32, l2 int32, l3 float64) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
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
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v5 == int32(0) {
		v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		F_addReplyBulkCBuffer(m, v12, l1, l2)
		mBase = m.M
		v14 = m.ExcPending
		if v14 != 0 {
			return
		} else {
			v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
			if v15 == int32(0) {
				return
			} else {
				v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
				F_addReplyDouble(m, v18, l3)
				mBase = m.M
				v20 = m.ExcPending
				if v20 != 0 {
					return
				} else {
					return
				}
			}
		}
	} else {
		v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		F_addReplyArrayLen(m, v8, int32(2))
		mBase = m.M
		v11 = m.ExcPending
		if v11 != 0 {
			return
		} else {
			v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			F_addReplyBulkCBuffer(m, v12, l1, l2)
			mBase = m.M
			v14 = m.ExcPending
			if v14 != 0 {
				return
			} else {
				v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
				if v15 == int32(0) {
					return
				} else {
					v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
					F_addReplyDouble(m, v18, l3)
					mBase = m.M
					v20 = m.ExcPending
					if v20 != 0 {
						return
					} else {
						return
					}
				}
			}
		}
	}
}
