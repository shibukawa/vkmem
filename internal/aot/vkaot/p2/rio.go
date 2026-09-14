package p2

import (
	base "github.com/shibukawa/valkeymem/internal/aot/vkaot/base"
	"unsafe"
)

func F_rioBufferWrite(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v10 int64
	_ = v10
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v5 = F_sdscatlen(m, v4, l1, l2)
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		return int32(0)
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(l0)+48)) = v5
		v10 = *(*int64)(unsafe.Add(mBase, uint32(l0)+56))
		*(*int64)(unsafe.Add(mBase, uint32(l0)+56)) = v10 + base.I64_extend_i32_u(l2)
		return int32(1)
	}
}
func F_rioConnRead(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v57 int32
	_ = v57
	var v61 int64
	_ = v61
	var v64 int64
	_ = v64
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
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
	var v85 int32
	_ = v85
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
	var v106 int32
	_ = v106
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
	var v115 int32
	_ = v115
	var v117 int32
	_ = v117
	var v124 int32
	_ = v124
	var v127 int32
	_ = v127
	var v131 int32
	_ = v131
	var v134 int32
	_ = v134
	var v138 int32
	_ = v138
	var v141 int32
	_ = v141
	var v145 int64
	_ = v145
	var v148 int64
	_ = v148
	var v151 int32
	_ = v151
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v166 int32
	_ = v166
	var v171 int32
	_ = v171
	var v174 int32
	_ = v174
	var v177 int32
	_ = v177
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v187 int32
	_ = v187
	var v191 int32
	_ = v191
	var v194 int32
	_ = v194
	var v199 int32
	_ = v199
	var v202 int32
	_ = v202
	var v205 int32
	_ = v205
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v216 int32
	_ = v216
	var v218 int32
	_ = v218
	var v220 int32
	_ = v220
	var v224 int32
	_ = v224
	var v226 int32
	_ = v226
	var v229 int32
	_ = v229
	var v259 int32
	_ = v259
	var v262 int32
	_ = v262
	var v270 int64
	_ = v270
	var v285 int32
	_ = v285
	var v288 int32
	_ = v288
	var v290 int32
	_ = v290
	var v295 int32
	_ = v295
	var v298 int32
	_ = v298
	var v301 int32
	_ = v301
	var v304 int32
	_ = v304
	var v305 int32
	_ = v305
	var v307 int64
	_ = v307
	var v317 int32
	_ = v317
	var v319 int32
	_ = v319
	var v320 int32
	_ = v320
	var v321 int32
	_ = v321
	var v324 int32
	_ = v324
	var v327 int32
	_ = v327
	var v331 int32
	_ = v331
	var v333 int32
	_ = v333
	var v334 int32
	_ = v334
	var v335 int32
	_ = v335
	var v338 int32
	_ = v338
	var v341 int32
	_ = v341
	var v345 int64
	_ = v345
	var v348 int32
	_ = v348
	var v349 int32
	_ = v349
	var v350 int32
	_ = v350
	var v353 int32
	_ = v353
	var v356 int64
	_ = v356
	var v364 int32
	_ = v364
	var v366 int32
	_ = v366
	var v367 int32
	_ = v367
	var v368 int32
	_ = v368
	var v371 int32
	_ = v371
	var v374 int32
	_ = v374
	var v376 int32
	_ = v376
	var v377 int32
	_ = v377
	var v378 int32
	_ = v378
	var v381 int32
	_ = v381
	var v390 int32
	_ = v390
	var v393 int32
	_ = v393
	var v397 int32
	_ = v397
	var v400 int32
	_ = v400
	var v404 int32
	_ = v404
	var v407 int32
	_ = v407
	var v411 int64
	_ = v411
	var v414 int64
	_ = v414
	var v417 int32
	_ = v417
	var v418 int32
	_ = v418
	var v421 int32
	_ = v421
	var v424 int32
	_ = v424
	var v425 int32
	_ = v425
	var v429 int32
	_ = v429
	var v431 int32
	_ = v431
	var v432 int32
	_ = v432
	var v438 int32
	_ = v438
	var v441 int32
	_ = v441
	var v444 int32
	_ = v444
	var v447 int32
	_ = v447
	var v448 int32
	_ = v448
	var v449 int32
	_ = v449
	var v451 int32
	_ = v451
	var v452 int32
	_ = v452
	var v453 int32
	_ = v453
	var v454 int32
	_ = v454
	var v459 int32
	_ = v459
	var v460 int32
	_ = v460
	var v463 int32
	_ = v463
	var v464 int32
	_ = v464
	var v471 int32
	_ = v471
	var v473 int32
	_ = v473
	var v478 int32
	_ = v478
	var v480 int32
	_ = v480
	var v483 int64
	_ = v483
	var v487 int32
	_ = v487
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
	v17 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14+int32(-1)))))
	v19 = v17 & int32(7)
	switch v19 {
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
		goto L3
	}
L1:
	;
	if base.Ui32(l2) <= base.Ui32(v85+v83) {
		v111 = v14
		goto L9
	} else {
		goto L10
	}
L2:
	;
	v82 = v78
	v83 = v79
	v84 = v80
	v85 = int32(0)
	goto L1
L3:
	;
	v74 = int32(0)
	v76 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	v78 = l0 + int32(56)
	v79 = v74
	v80 = v74 - v76
	goto L2
L4:
	;
	v61 = *(*int64)(unsafe.Add(mBase, uint32(v14+int32(-9))))
	v64 = *(*int64)(unsafe.Add(mBase, uint32(v14+int32(-17))))
	v69 = base.I32_wrap_i64(v64)
	v70 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	v82 = l0 + int32(56)
	v83 = v69
	v84 = v69 - v70
	v85 = base.I32_wrap_i64(v61 - v64)
	goto L1
L5:
	;
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v14+int32(-5))))
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v14+int32(-9))))
	v57 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	v82 = l0 + int32(56)
	v83 = v53
	v84 = v53 - v57
	v85 = v50 - v53
	goto L1
L6:
	;
	v39 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v14+int32(-3)))))
	v42 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v14+int32(-5)))))
	v46 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	v82 = l0 + int32(56)
	v83 = v42
	v84 = v42 - v46
	v85 = v39 - v42
	goto L1
L7:
	;
	v28 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14+int32(-2)))))
	v31 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14+int32(-3)))))
	v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	v82 = l0 + int32(56)
	v83 = v31
	v84 = v31 - v35
	v85 = v28 - v31
	goto L1
L8:
	;
	v23 = int32(base.Ui32(v17) >> (uint(int32(3)) % 32))
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	v78 = l0 + int32(56)
	v79 = v23
	v80 = v23 - v24
	goto L2
L9:
	;
	if base.Ui32(l2) <= base.Ui32(v84) {
		goto L19
	} else {
		goto L20
	}
L10:
	;
	switch v19 {
	case 0:
		goto L16
	case 1:
		goto L15
	case 2:
		goto L14
	case 3:
		goto L13
	case 4:
		goto L12
	default:
		v104 = int32(0)
		goto L11
	}
L11:
	;
	v106 = F_sdsMakeRoomFor(m, v14, l2-v104)
	mBase = m.M
	v109 = m.ExcPending
	if v109 != 0 {
		goto L17
	} else {
		goto L18
	}
L12:
	;
	v103 = *(*int32)(unsafe.Add(mBase, uint32(v14+int32(-17))))
	v104 = v103
	goto L11
L13:
	;
	v100 = *(*int32)(unsafe.Add(mBase, uint32(v14+int32(-9))))
	v104 = v100
	goto L11
L14:
	;
	v97 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v14+int32(-5)))))
	v104 = v97
	goto L11
L15:
	;
	v94 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14+int32(-3)))))
	v104 = v94
	goto L11
L16:
	;
	v104 = int32(base.Ui32(v17) >> (uint(int32(3)) % 32))
	goto L11
L17:
	;
	return int32(0)
L18:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+64)) = v106
	v111 = v106
	goto L9
L19:
	;
	v259 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
	if v259 == int32(0) {
		goto L61
	} else {
		goto L62
	}
L20:
	;
	v115 = int32(-1)
	v117 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v111+v115))))
	switch v117&int32(7) + v115 {
	case 0:
		goto L25
	case 1:
		goto L24
	case 2:
		goto L23
	case 3:
		goto L22
	default:
		v151 = int32(0)
		goto L21
	}
L21:
	;
	if base.Ui32(l2-v84) <= base.Ui32(v151) {
		goto L19
	} else {
		goto L26
	}
L22:
	;
	v145 = *(*int64)(unsafe.Add(mBase, uint32(v111+int32(-9))))
	v148 = *(*int64)(unsafe.Add(mBase, uint32(v111+int32(-17))))
	v151 = base.I32_wrap_i64(v145 - v148)
	goto L21
L23:
	;
	v138 = *(*int32)(unsafe.Add(mBase, uint32(v111+int32(-5))))
	v141 = *(*int32)(unsafe.Add(mBase, uint32(v111+int32(-9))))
	v151 = v138 - v141
	goto L21
L24:
	;
	v131 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v111+int32(-3)))))
	v134 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v111+int32(-5)))))
	v151 = v131 - v134
	goto L21
L25:
	;
	v124 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v111+int32(-2)))))
	v127 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v111+int32(-3)))))
	v151 = v124 - v127
	goto L21
L26:
	;
	v154 = *(*int32)(unsafe.Add(mBase, uint32(v82)))
	v155 = int32(-1)
	v163 = v111 + v155
	v164 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v163))))
	v166 = v164 & int32(7)
	switch v166 {
	case 0:
		goto L34
	case 1:
		goto L33
	case 2:
		goto L32
	case 3:
		goto L31
	case 4:
		goto L30
	default:
		goto L28
	}
L27:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v82))) = int64(0)
	goto L19
L28:
	;
	goto L27
L29:
	;
	if v181 == int32(0) {
		goto L28
	} else {
		goto L35
	}
L30:
	;
	v180 = *(*int32)(unsafe.Add(mBase, uint32(v111+int32(-17))))
	v181 = v180
	goto L29
L31:
	;
	v177 = *(*int32)(unsafe.Add(mBase, uint32(v111+int32(-9))))
	v181 = v177
	goto L29
L32:
	;
	v174 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v111+int32(-5)))))
	v181 = v174
	goto L29
L33:
	;
	v171 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v111+int32(-3)))))
	v181 = v171
	goto L29
L34:
	;
	v181 = int32(base.Ui32(v164) >> (uint(int32(3)) % 32))
	goto L29
L35:
	;
	v187 = int32(-1)&v181 + v155
	v191 = v154>>(uint(int32(31))%32)&v181 + v154
	v194 = v187 - v191 + int32(1)
	switch v166 {
	default:
		goto L41
	case 1:
		goto L40
	case 2:
		goto L39
	case 3:
		goto L38
	case 4:
		goto L37
	}
L36:
	;
	v210 = int32(0)
	v212 = base.B2i32(base.Ui32(v191) < base.Ui32(v209))
	if base.Ui32(v191) < base.Ui32(v209) {
		goto L43
	} else {
		goto L44
	}
L37:
	;
	v208 = *(*int32)(unsafe.Add(mBase, uint32(v111+int32(-17))))
	v209 = v208
	goto L36
L38:
	;
	v205 = *(*int32)(unsafe.Add(mBase, uint32(v111+int32(-9))))
	v209 = v205
	goto L36
L39:
	;
	v202 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v111+int32(-5)))))
	v209 = v202
	goto L36
L40:
	;
	v199 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v111+int32(-3)))))
	v209 = v199
	goto L36
L41:
	;
	v209 = int32(base.Ui32(v164) >> (uint(int32(3)) % 32))
	goto L36
L42:
	;
	v226 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v111+v220))) = uint8(v226)
	switch v166 {
	default:
		goto L60
	case 1:
		goto L59
	case 2:
		goto L58
	case 3:
		goto L57
	case 4:
		goto L56
	}
L43:
	;
	v213 = v191
	goto L45
L44:
	;
	v213 = v210
	goto L45
L45:
	;
	v214 = v209 - v213
	if base.Ui32(v194) < base.Ui32(v214) {
		goto L46
	} else {
		goto L47
	}
L46:
	;
	v216 = v194
	goto L48
L47:
	;
	v216 = v214
	goto L48
L48:
	;
	if v187 < v191 {
		goto L49
	} else {
		goto L50
	}
L49:
	;
	v218 = v210
	goto L51
L50:
	;
	v218 = v216
	goto L51
L51:
	;
	if base.Ui32(v191) < base.Ui32(v209) {
		goto L52
	} else {
		goto L53
	}
L52:
	;
	v220 = v218
	goto L54
L53:
	;
	v220 = int32(0)
	goto L54
L54:
	;
	if v220 == int32(0) {
		goto L42
	} else {
		goto L55
	}
L55:
	;
	v224 = F_memmove(m, v111, v111+v213, v220)
	mBase = m.M
	goto L42
L56:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v111+int32(-17)))) = base.I64_extend_i32_u(v220)
	goto L28
L57:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v111+int32(-9)))) = v220
	goto L27
L58:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v111+int32(-5)))) = uint16(v220)
	goto L27
L59:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v111+int32(-3)))) = uint8(v220)
	goto L27
L60:
	;
	v229 = v220 << (uint(int32(3)) % 32)
	*(*uint8)(unsafe.Add(mBase, uint32(v163))) = uint8(v229)
	goto L27
L61:
	;
	v270 = base.I64_extend_i32_u(l2)
	goto L66
L62:
	;
	v262 = *(*int32)(unsafe.Add(mBase, uint32(l0)+72))
	if base.Ui32(v262+l2) <= base.Ui32(v259) {
		goto L61
	} else {
		goto L63
	}
L63:
	;
	goto L64
L64:
	;
	*(*int32)(unsafe.Add(mBase, _consts[5])) = int32(61)
	return int32(0)
L65:
	;
	return v487
L66:
	;
	v285 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
	v288 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v285+int32(-1)))))
	v290 = v288 & int32(7)
	switch v290 {
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
		v305 = int32(0)
		goto L68
	}
L67:
	;
	if l2 == int32(0) {
		goto L122
	} else {
		goto L123
	}
L68:
	;
	v307 = *(*int64)(unsafe.Add(mBase, uint32(v82)))
	if v270 <= base.I64_extend_i32_u(v305)-v307 {
		goto L74
	} else {
		goto L75
	}
L69:
	;
	v304 = *(*int32)(unsafe.Add(mBase, uint32(v285+int32(-17))))
	v305 = v304
	goto L68
L70:
	;
	v301 = *(*int32)(unsafe.Add(mBase, uint32(v285+int32(-9))))
	v305 = v301
	goto L68
L71:
	;
	v298 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v285+int32(-5)))))
	v305 = v298
	goto L68
L72:
	;
	v295 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v285+int32(-3)))))
	v305 = v295
	goto L68
L73:
	;
	v305 = int32(base.Ui32(v288) >> (uint(int32(3)) % 32))
	goto L68
L74:
	;
	goto L67
L75:
	;
	switch v290 {
	case 0:
		goto L84
	case 1:
		goto L79
	case 2:
		goto L83
	case 3:
		goto L82
	case 4:
		goto L81
	default:
		goto L80
	}
L76:
	;
	v421 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
	if v421 == int32(0) {
		v431 = v418
		goto L102
	} else {
		goto L103
	}
L77:
	;
	switch v290 + int32(-1) {
	case 0:
		goto L101
	case 1:
		goto L100
	case 2:
		goto L99
	case 3:
		goto L98
	default:
		v417 = v381
		v418 = int32(0)
		goto L76
	}
L78:
	;
	if base.Ui32(v377) <= base.Ui32(v378) {
		v417 = v376
		v418 = v377
		goto L76
	} else {
		goto L97
	}
L79:
	;
	v364 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v285+int32(-3)))))
	v366 = v364 - base.I32_wrap_i64(v307)
	v367 = l2 - v366
	v368 = int32(16384)
	if base.Ui32(v368) < base.Ui32(v367) {
		goto L94
	} else {
		goto L95
	}
L80:
	;
	v381 = int32(0) - base.I32_wrap_i64(v307)
	goto L77
L81:
	;
	v345 = *(*int64)(unsafe.Add(mBase, uint32(v285+int32(-17))))
	v348 = base.I32_wrap_i64(v345) - base.I32_wrap_i64(v307)
	v349 = l2 - v348
	v350 = int32(16384)
	if base.Ui32(v350) < base.Ui32(v349) {
		goto L91
	} else {
		goto L92
	}
L82:
	;
	v331 = *(*int32)(unsafe.Add(mBase, uint32(v285+int32(-9))))
	v333 = v331 - base.I32_wrap_i64(v307)
	v334 = l2 - v333
	v335 = int32(16384)
	if base.Ui32(v335) < base.Ui32(v334) {
		goto L88
	} else {
		goto L89
	}
L83:
	;
	v317 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v285+int32(-5)))))
	v319 = v317 - base.I32_wrap_i64(v307)
	v320 = l2 - v319
	v321 = int32(16384)
	if base.Ui32(v321) < base.Ui32(v320) {
		goto L85
	} else {
		goto L86
	}
L84:
	;
	v417 = int32(base.Ui32(v288)>>(uint(int32(3))%32)) - base.I32_wrap_i64(v307)
	v418 = int32(0)
	goto L76
L85:
	;
	v324 = v320
	goto L87
L86:
	;
	v324 = v321
	goto L87
L87:
	;
	v327 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v285+int32(-3)))))
	v376 = v319
	v377 = v324
	v378 = v327 - v317
	goto L78
L88:
	;
	v338 = v334
	goto L90
L89:
	;
	v338 = v335
	goto L90
L90:
	;
	v341 = *(*int32)(unsafe.Add(mBase, uint32(v285+int32(-5))))
	v376 = v333
	v377 = v338
	v378 = v341 - v331
	goto L78
L91:
	;
	v353 = v349
	goto L93
L92:
	;
	v353 = v350
	goto L93
L93:
	;
	v356 = *(*int64)(unsafe.Add(mBase, uint32(v285+int32(-9))))
	v376 = v348
	v377 = v353
	v378 = base.I32_wrap_i64(v356 - v345)
	goto L78
L94:
	;
	v371 = v367
	goto L96
L95:
	;
	v371 = v368
	goto L96
L96:
	;
	v374 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v285+int32(-2)))))
	v376 = v366
	v377 = v371
	v378 = v374 - v364
	goto L78
L97:
	;
	v381 = v376
	goto L77
L98:
	;
	v411 = *(*int64)(unsafe.Add(mBase, uint32(v285+int32(-9))))
	v414 = *(*int64)(unsafe.Add(mBase, uint32(v285+int32(-17))))
	v417 = v381
	v418 = base.I32_wrap_i64(v411 - v414)
	goto L76
L99:
	;
	v404 = *(*int32)(unsafe.Add(mBase, uint32(v285+int32(-5))))
	v407 = *(*int32)(unsafe.Add(mBase, uint32(v285+int32(-9))))
	v417 = v381
	v418 = v404 - v407
	goto L76
L100:
	;
	v397 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v285+int32(-3)))))
	v400 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v285+int32(-5)))))
	v417 = v381
	v418 = v397 - v400
	goto L76
L101:
	;
	v390 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v285+int32(-2)))))
	v393 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v285+int32(-3)))))
	v417 = v381
	v418 = v390 - v393
	goto L76
L102:
	;
	v432 = int32(0)
	switch v290 {
	case 0:
		goto L112
	case 1:
		goto L111
	case 2:
		goto L110
	case 3:
		goto L109
	case 4:
		goto L108
	default:
		v448 = v432
		goto L107
	}
L103:
	;
	v424 = *(*int32)(unsafe.Add(mBase, uint32(l0)+72))
	v425 = v424 + v417
	if base.Ui32(v421) < base.Ui32(v425+v418) {
		goto L104
	} else {
		goto L105
	}
L104:
	;
	v429 = v421 - v425
	goto L106
L105:
	;
	v429 = v418
	goto L106
L106:
	;
	v431 = v429
	goto L102
L107:
	;
	v449 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v451 = *(*int32)(unsafe.Add(mBase, uint32(v449)))
	v452 = *(*int32)(unsafe.Add(mBase, uint32(v451)+76))
	v453 = m.T0[v452].(func(*base.Module, int32, int32, int32) int32)(m, v449, v285+v448, v431)
	mBase = m.M
	v454 = m.ExcPending
	if v454 != 0 {
		goto L17
	} else {
		goto L113
	}
L108:
	;
	v447 = *(*int32)(unsafe.Add(mBase, uint32(v285+int32(-17))))
	v448 = v447
	goto L107
L109:
	;
	v444 = *(*int32)(unsafe.Add(mBase, uint32(v285+int32(-9))))
	v448 = v444
	goto L107
L110:
	;
	v441 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v285+int32(-5)))))
	v448 = v441
	goto L107
L111:
	;
	v438 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v285+int32(-3)))))
	v448 = v438
	goto L107
L112:
	;
	v448 = int32(base.Ui32(v288) >> (uint(int32(3)) % 32))
	goto L107
L113:
	;
	if v453 == int32(0) {
		v487 = v432
		goto L65
	} else {
		goto L114
	}
L114:
	;
	if int32(-1) < v453 {
		goto L115
	} else {
		goto L116
	}
L115:
	;
	v471 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
	F_sdsIncrLen(m, v471, v453)
	mBase = m.M
	v473 = m.ExcPending
	if v473 != 0 {
		goto L17
	} else {
		goto L120
	}
L116:
	;
	v459 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v460 = *(*int32)(unsafe.Add(mBase, uint32(v459)+8))
	if v460 == int32(27) {
		goto L66
	} else {
		goto L117
	}
L117:
	;
	v463 = int32(9116376)
	goto L118
L118:
	;
	v464 = *(*int32)(unsafe.Add(mBase, _consts[5]))
	if v464 != int32(6) {
		v487 = v432
		goto L65
	} else {
		goto L119
	}
L119:
	;
	*(*int32)(unsafe.Add(mBase, _consts[5])) = int32(73)
	return int32(0)
L120:
	;
	goto L66
L121:
	;
	v480 = *(*int32)(unsafe.Add(mBase, uint32(l0)+72))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+72)) = v480 + l2
	v483 = *(*int64)(unsafe.Add(mBase, uint32(l0)+56))
	*(*int64)(unsafe.Add(mBase, uint32(l0)+56)) = v483 + v270
	v487 = l2
	goto L65
L122:
	;
	goto L121
L123:
	;
	v478 = F__emscripten_memcpy_bulkmem(m, l1, v285+base.I32_wrap_i64(v307), l2)
	mBase = m.M
	goto L122
}
func F_rioConnsetRead(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	return int32(0)
}
func F_rioConnsetWrite(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v24 int32
	_ = v24
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v85 int32
	_ = v85
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v102 int32
	_ = v102
	var v104 int32
	_ = v104
	var v106 int32
	_ = v106
	var v110 int32
	_ = v110
	var v117 int32
	_ = v117
	var v119 int32
	_ = v119
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
	var v131 int32
	_ = v131
	var v136 int32
	_ = v136
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
	var v155 int32
	_ = v155
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v164 int32
	_ = v164
	var v168 int32
	_ = v168
	var v174 int64
	_ = v174
	var v179 int32
	_ = v179
	var v189 int32
	_ = v189
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v196 int32
	_ = v196
	var v200 int32
	_ = v200
	var v204 int32
	_ = v204
	var v214 int32
	_ = v214
	v4 = int32(0)
	v12 = base.B2i32(l1|l2 == v4)
	if l2 == v4 {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	return int32(1)
L2:
	;
	switch v51 & int32(7) {
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
		goto L15
	}
L3:
	;
	if v12 == int32(0) {
		goto L1
	} else {
		goto L14
	}
L4:
	;
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+72))
	v16 = F_sdscatlen(m, v15, l1, l2)
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	return int32(0)
L6:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+72)) = v16
	v24 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16+int32(-1)))))
	switch v24 & int32(7) {
	case 0:
		goto L12
	case 1:
		goto L11
	case 2:
		goto L10
	case 3:
		goto L9
	case 4:
		goto L8
	default:
		v41 = int32(0)
		goto L7
	}
L7:
	;
	if v12|base.B2i32(base.Ui32(int32(16384)) < base.Ui32(v41)) != 0 {
		v51 = v24
		v53 = v16
		goto L2
	} else {
		goto L13
	}
L8:
	;
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v16+int32(-17))))
	v41 = v40
	goto L7
L9:
	;
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v16+int32(-9))))
	v41 = v37
	goto L7
L10:
	;
	v34 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v16+int32(-5)))))
	v41 = v34
	goto L7
L11:
	;
	v31 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16+int32(-3)))))
	v41 = v31
	goto L7
L12:
	;
	v41 = int32(base.Ui32(v24) >> (uint(int32(3)) % 32))
	goto L7
L13:
	;
	goto L1
L14:
	;
	v47 = *(*int32)(unsafe.Add(mBase, uint32(l0)+72))
	v50 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v47+int32(-1)))))
	v51 = v50
	v53 = v47
	goto L2
L15:
	;
	v189 = *(*int32)(unsafe.Add(mBase, uint32(l0)+72))
	v192 = v189 + int32(-1)
	v193 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v192))))
	switch v193 & int32(7) {
	case 0:
		goto L56
	case 1:
		goto L55
	case 2:
		goto L54
	case 3:
		goto L53
	case 4:
		goto L52
	default:
		goto L51
	}
L16:
	;
	if v72 == int32(0) {
		goto L15
	} else {
		goto L22
	}
L17:
	;
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v53+int32(-17))))
	v72 = v71
	goto L16
L18:
	;
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v53+int32(-9))))
	v72 = v68
	goto L16
L19:
	;
	v65 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v53+int32(-5)))))
	v72 = v65
	goto L16
L20:
	;
	v62 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v53+int32(-3)))))
	v72 = v62
	goto L16
L21:
	;
	v72 = int32(base.Ui32(v51&int32(248)) >> (uint(int32(3)) % 32))
	goto L16
L22:
	;
	v75 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	v78 = v75
	v80 = v53
	v81 = v72
	goto L23
L23:
	;
	v85 = int32(16384)
	if base.Ui32(v81) < base.Ui32(v85) {
		goto L25
	} else {
		goto L26
	}
L24:
	;
	goto L15
L25:
	;
	v88 = v81
	goto L27
L26:
	;
	v88 = v85
	goto L27
L27:
	;
	v89 = int32(0)
	if v78 < int32(1) {
		v164 = v78
		v168 = v89
		goto L28
	} else {
		goto L29
	}
L28:
	;
	if v168 != v164 {
		goto L47
	} else {
		goto L48
	}
L29:
	;
	v99 = v89
	v100 = v89
	goto L30
L30:
	;
	v102 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v104 = v100 << (uint(int32(2)) % 32)
	v106 = *(*int32)(unsafe.Add(mBase, uint32(v102+v104)))
	if v106 != 0 {
		goto L33
	} else {
		goto L34
	}
L31:
	;
	v164 = v160
	v168 = v155
	goto L28
L32:
	;
	v159 = v100 + int32(1)
	v160 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	if v159 < v160 {
		v99 = v155
		v100 = v159
		goto L30
	} else {
		goto L46
	}
L33:
	;
	v155 = v99 + int32(1)
	goto L32
L34:
	;
	v110 = int32(0)
	goto L35
L35:
	;
	v117 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v119 = *(*int32)(unsafe.Add(mBase, uint32(v117+v104)))
	v122 = *(*int32)(unsafe.Add(mBase, uint32(v119)))
	v123 = *(*int32)(unsafe.Add(mBase, uint32(v122)+68))
	v124 = m.T0[v123].(func(*base.Module, int32, int32, int32) int32)(m, v119, v80+v110, v88-v110)
	mBase = m.M
	v125 = m.ExcPending
	if v125 != 0 {
		goto L5
	} else {
		goto L38
	}
L37:
	;
	v145 = v124 + v110
	if v145 != v88 {
		v110 = v145
		goto L35
	} else {
		goto L45
	}
L38:
	;
	if int32(0) < v124 {
		goto L37
	} else {
		goto L39
	}
L39:
	;
	v128 = int32(9116376)
	goto L40
L40:
	;
	if v124 != int32(-1) {
		goto L41
	} else {
		goto L42
	}
L41:
	;
	v136 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v138 = *(*int32)(unsafe.Add(mBase, _consts[5]))
	*(*int32)(unsafe.Add(mBase, uint32(v136+v104))) = v138
	v140 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v141 = v140 + v104
	v142 = *(*int32)(unsafe.Add(mBase, uint32(v141)))
	if v142 != 0 {
		v155 = v99
		goto L32
	} else {
		goto L44
	}
L42:
	;
	v131 = *(*int32)(unsafe.Add(mBase, _consts[5]))
	if v131 != int32(6) {
		goto L41
	} else {
		goto L43
	}
L43:
	;
	*(*int32)(unsafe.Add(mBase, _consts[5])) = int32(73)
	goto L41
L44:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v141))) = int32(29)
	v155 = v99
	goto L32
L45:
	;
	v155 = v99
	goto L32
L46:
	;
	goto L31
L47:
	;
	v174 = *(*int64)(unsafe.Add(mBase, uint32(l0)+64))
	*(*int64)(unsafe.Add(mBase, uint32(l0)+64)) = v174 + base.I64_extend_i32_u(v88)
	v179 = v81 - v88
	if v179 != 0 {
		v78 = v164
		v80 = v80 + v88
		v81 = v179
		goto L23
	} else {
		goto L49
	}
L48:
	;
	return int32(0)
L49:
	;
	goto L24
L50:
	;
	goto L1
L51:
	;
	v214 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v189))) = uint8(v214)
	goto L50
L52:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v189+int32(-17)))) = int64(0)
	goto L51
L53:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v189+int32(-9)))) = int32(0)
	goto L51
L54:
	;
	v204 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v189+int32(-5)))) = uint16(v204)
	goto L51
L55:
	;
	v200 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v189+int32(-3)))) = uint8(v200)
	goto L51
L56:
	;
	v196 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v192))) = uint8(v196)
	goto L51
}
func F_rioFdTell(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int64
	_ = v2
	v2 = *(*int64)(unsafe.Add(mBase, uint32(l0)+56))
	return v2
}
func F_rioFdWrite(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v16 int32
	_ = v16
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v81 int64
	_ = v81
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
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v117 int32
	_ = v117
	var v121 int32
	_ = v121
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v137 int32
	_ = v137
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v150 int32
	_ = v150
	var v162 int64
	_ = v162
	var v164 int64
	_ = v164
	var v167 int32
	_ = v167
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v174 int32
	_ = v174
	var v178 int32
	_ = v178
	var v182 int32
	_ = v182
	var v192 int32
	_ = v192
	var v199 int32
	_ = v199
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
	if base.Ui32(l2) < base.Ui32(int32(16385)) {
		goto L4
	} else {
		goto L5
	}
L1:
	;
	return v199
L2:
	;
	v164 = *(*int64)(unsafe.Add(mBase, uint32(l0)+56))
	*(*int64)(unsafe.Add(mBase, uint32(l0)+56)) = v164 + v162
	v167 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
	v170 = v167 + int32(-1)
	v171 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v170))))
	switch v171 & int32(7) {
	case 0:
		goto L49
	case 1:
		goto L48
	case 2:
		goto L47
	case 3:
		goto L46
	case 4:
		goto L45
	default:
		goto L44
	}
L3:
	;
	v117 = v104
	v121 = int32(0)
	goto L32
L4:
	;
	if l2 != 0 {
		goto L17
	} else {
		goto L18
	}
L5:
	;
	v16 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11+int32(-1)))))
	switch v16 & int32(7) {
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
		v103 = l1
		v104 = l2
		goto L3
	}
L6:
	;
	if v33 == int32(0) {
		v103 = l1
		v104 = l2
		goto L3
	} else {
		goto L12
	}
L7:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v11+int32(-17))))
	v33 = v32
	goto L6
L8:
	;
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v11+int32(-9))))
	v33 = v29
	goto L6
L9:
	;
	v26 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v11+int32(-5)))))
	v33 = v26
	goto L6
L10:
	;
	v23 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11+int32(-3)))))
	v33 = v23
	goto L6
L11:
	;
	v33 = int32(base.Ui32(v16) >> (uint(int32(3)) % 32))
	goto L6
L12:
	;
	v36 = int32(0)
	v39 = F_rioFdWrite(m, l0, v36, v36)
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	return int32(0)
L14:
	;
	if v39 != 0 {
		v103 = l1
		v104 = l2
		goto L3
	} else {
		goto L15
	}
L15:
	;
	v199 = v36
	goto L1
L16:
	;
	v81 = int64(0)
	switch v78 & int32(7) {
	case 0:
		goto L30
	case 1:
		goto L29
	case 2:
		goto L28
	case 3:
		goto L27
	case 4:
		goto L26
	default:
		v162 = v81
		goto L2
	}
L17:
	;
	v46 = F_sdscatlen(m, v11, l1, l2)
	mBase = m.M
	v47 = m.ExcPending
	if v47 != 0 {
		goto L13
	} else {
		goto L19
	}
L18:
	;
	v45 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11+int32(-1)))))
	v75 = v11
	v78 = v45
	goto L16
L19:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+64)) = v46
	v49 = int32(1)
	v52 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v46+int32(-1)))))
	switch v52&int32(7) + int32(-2) {
	case 0:
		goto L23
	case 1:
		goto L22
	case 2:
		goto L21
	default:
		v199 = v49
		goto L1
	}
L20:
	;
	if (base.B2i32(l1|l2 == int32(0))^int32(-1))&base.B2i32(base.Ui32(v66) < base.Ui32(int32(16385))) != 0 {
		v199 = v49
		goto L1
	} else {
		goto L24
	}
L21:
	;
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v46+int32(-17))))
	v66 = v65
	goto L20
L22:
	;
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v46+int32(-9))))
	v66 = v62
	goto L20
L23:
	;
	v59 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v46+int32(-5)))))
	v66 = v59
	goto L20
L24:
	;
	v75 = v46
	v78 = v52
	goto L16
L25:
	;
	if v100 == int32(0) {
		v162 = v81
		goto L2
	} else {
		goto L31
	}
L26:
	;
	v99 = *(*int32)(unsafe.Add(mBase, uint32(v75+int32(-17))))
	v100 = v99
	goto L25
L27:
	;
	v96 = *(*int32)(unsafe.Add(mBase, uint32(v75+int32(-9))))
	v100 = v96
	goto L25
L28:
	;
	v93 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v75+int32(-5)))))
	v100 = v93
	goto L25
L29:
	;
	v90 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v75+int32(-3)))))
	v100 = v90
	goto L25
L30:
	;
	v100 = int32(base.Ui32(v78&int32(248)) >> (uint(int32(3)) % 32))
	goto L25
L31:
	;
	v103 = v75
	v104 = v100
	goto L3
L32:
	;
	goto L35
L33:
	;
	v162 = base.I64_extend_i32_u(v104)
	goto L2
L34:
	;
	v150 = v134 + v121
	if v104 != v150 {
		v117 = v104 - v150
		v121 = v150
		goto L32
	} else {
		goto L42
	}
L35:
	;
	v133 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v134 = F_write(m, v133, v103+v121, v117)
	mBase = m.M
	if int32(0) < v134 {
		goto L34
	} else {
		goto L37
	}
L36:
	;
	if v141 != int32(6) {
		v199 = v137
		goto L1
	} else {
		goto L41
	}
L37:
	;
	v137 = int32(0)
	if v134 != int32(-1) {
		v199 = v137
		goto L1
	} else {
		goto L38
	}
L38:
	;
	v140 = int32(9116376)
	goto L39
L39:
	;
	v141 = *(*int32)(unsafe.Add(mBase, _consts[5]))
	if v141 == int32(27) {
		goto L35
	} else {
		goto L40
	}
L40:
	;
	goto L36
L41:
	;
	*(*int32)(unsafe.Add(mBase, _consts[5])) = int32(73)
	return int32(0)
L42:
	;
	goto L33
L43:
	;
	v199 = int32(1)
	goto L1
L44:
	;
	v192 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v167))) = uint8(v192)
	goto L43
L45:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v167+int32(-17)))) = int64(0)
	goto L44
L46:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v167+int32(-9)))) = int32(0)
	goto L44
L47:
	;
	v182 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v167+int32(-5)))) = uint16(v182)
	goto L44
L48:
	;
	v178 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v167+int32(-3)))) = uint8(v178)
	goto L44
L49:
	;
	v174 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v170))) = uint8(v174)
	goto L44
}
func F_rioFileTell(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int64
	_ = v3
	var v6 int32
	_ = v6
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v3 = F___ftello(m, v2)
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return int64(0)
	} else {
		return v3
	}
}
func F_rioInitWithFile(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int64
	_ = v9
	var v18 int32
	_ = v18
	v7 = F__emscripten_memcpy_bulkmem(m, l0, int32(_a144), int32(80))
	mBase = m.M
	v9 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v7)+56)) = v9
	*(*int32)(unsafe.Add(mBase, uint32(v7)+48)) = l1
	*(*int64)(unsafe.Add(mBase, uint32(v7+int32(64)))) = v9
	v18 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v7+int32(72)))) = uint8(v18)
	return
}
func F_rioWriteBulkCount(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v15 int32
	_ = v15
	var v17 int64
	_ = v17
	var v26 int32
	_ = v26
	var v30 int32
	_ = v30
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int64
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v58 int32
	_ = v58
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v69 int32
	_ = v69
	var v74 int32
	_ = v74
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v94 int64
	_ = v94
	var v98 int32
	_ = v98
	var v102 int32
	_ = v102
	var v107 int32
	_ = v107
	v2 = l1
	v9 = m.G0
	v11 = v9 - int32(128)
	m.G0 = v11
	*(*uint8)(unsafe.Add(mBase, uint32(v11))) = uint8(v2)
	v15 = v11 | int32(1)
	v17 = base.I64_extend_i32_s(l2)
	if v17 <= int64(-1) {
		goto L5
	} else {
		goto L6
	}
L1:
	;
	v62 = int32(2573)
	*(*uint16)(unsafe.Add(mBase, uint32(v58+v11+int32(1)))) = uint16(v62)
	v64 = int32(0)
	v65 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
	if v65&int32(6) != 0 {
		v107 = v64
		goto L10
	} else {
		goto L11
	}
L2:
	;
	v58 = int32(0)
	goto L1
L4:
	;
	v39 = F_ull2string(m, v35, v36, v37)
	mBase = m.M
	if v39 == int32(0) {
		goto L2
	} else {
		goto L8
	}
L5:
	;
	goto L7
L6:
	;
	v35 = v15
	v36 = int32(127)
	v37 = v17
	v38 = int32(0)
	goto L4
L7:
	;
	v26 = int32(45)
	*(*uint8)(unsafe.Add(mBase, uint32(v15))) = uint8(v26)
	v30 = int32(1)
	v35 = v15 + v30
	v36 = int32(126)
	v37 = int64(0) - v17
	v38 = v30
	goto L4
L8:
	;
	v58 = v39 + v38
	goto L1
L10:
	;
	m.G0 = v11 + int32(128)
	return v107
L11:
	;
	v69 = v58 + int32(3)
	if v69 == int32(0) {
		v107 = v64
		goto L10
	} else {
		goto L12
	}
L12:
	;
	v74 = v69
	v78 = v11
	goto L13
L13:
	;
	v80 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	if base.Ui32(v80) < base.Ui32(v74) {
		goto L15
	} else {
		goto L16
	}
L14:
	;
	v107 = v69
	goto L10
L15:
	;
	v82 = v80
	goto L17
L16:
	;
	v82 = v74
	goto L17
L17:
	;
	if v80 != 0 {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	v83 = v82
	goto L20
L19:
	;
	v83 = v74
	goto L20
L20:
	;
	v84 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v84 == int32(0) {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	v91 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v92 = m.T0[v91].(func(*base.Module, int32, int32, int32) int32)(m, l0, v78, v83)
	mBase = m.M
	v93 = m.ExcPending
	if v93 != 0 {
		goto L23
	} else {
		goto L26
	}
L22:
	;
	m.T0[v84].(func(*base.Module, int32, int32, int32))(m, l0, v78, v83)
	mBase = m.M
	v90 = m.ExcPending
	if v90 != 0 {
		goto L23
	} else {
		goto L24
	}
L23:
	;
	return int32(0)
L24:
	;
	goto L21
L25:
	;
	v98 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = v98 + v83
	v102 = v74 - v83
	if v102 != 0 {
		v74 = v102
		v78 = v78 + v83
		goto L13
	} else {
		goto L28
	}
L26:
	;
	if v92 != 0 {
		goto L25
	} else {
		goto L27
	}
L27:
	;
	v94 = *(*int64)(unsafe.Add(mBase, uint32(l0)+32))
	*(*int64)(unsafe.Add(mBase, uint32(l0)+32)) = v94 | int64(2)
	v107 = v64
	goto L10
L28:
	;
	goto L14
}
func F_rioWriteBulkLongLong(m *base.Module, l0 int32, l1 int64) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v17 int32
	_ = v17
	var v21 int32
	_ = v21
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int64
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	v4 = m.G0
	v5 = int32(32)
	v6 = v4 - v5
	m.G0 = v6
	if l1 <= int64(-1) {
		v17 = int32(45)
		*(*uint8)(unsafe.Add(mBase, uint32(v6))) = uint8(v17)
		v21 = int32(1)
		v26 = v6 + v21
		v27 = int32(31)
		v28 = int64(0) - l1
		v29 = v21
	} else {
		v26 = v6
		v27 = v5
		v28 = l1
		v29 = int32(0)
	}
	v30 = F_ull2string(m, v26, v27, v28)
	mBase = m.M
	if v30 == int32(0) {
		v49 = int32(0)
	} else {
		v49 = v30 + v29
	}
	v50 = F_rioWriteBulkString(m, l0, v6, v49)
	mBase = m.M
	v53 = m.ExcPending
	if v53 != 0 {
		return int32(0)
	} else {
		m.G0 = v6 + int32(32)
		return v50
	}
}
func F_rioWriteBulkObject(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
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
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	switch int32(base.Ui32(v6)>>(uint(int32(4))%32)) & int32(15) {
	case 0, 8:
		v11 = F_objectGetVal(m, l1)
		mBase = m.M
		v13 = F_objectGetVal(m, l1)
		mBase = m.M
		v16 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13+int32(-1)))))
		switch v16 & int32(7) {
		case 0:
			v21 = F_rioWriteBulkString(m, l0, v11, int32(base.Ui32(v16)>>(uint(int32(3))%32)))
			mBase = m.M
			v24 = m.ExcPending
			if v24 != 0 {
				return int32(0)
			} else {
				return v21
			}
		case 1:
			v28 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13+int32(-3)))))
			v29 = F_rioWriteBulkString(m, l0, v11, v28)
			mBase = m.M
			v30 = m.ExcPending
			if v30 != 0 {
				return int32(0)
			} else {
				return v29
			}
		case 2:
			v34 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v13+int32(-5)))))
			v35 = F_rioWriteBulkString(m, l0, v11, v34)
			mBase = m.M
			v36 = m.ExcPending
			if v36 != 0 {
				return int32(0)
			} else {
				return v35
			}
		case 3:
			v40 = *(*int32)(unsafe.Add(mBase, uint32(v13+int32(-9))))
			v41 = F_rioWriteBulkString(m, l0, v11, v40)
			mBase = m.M
			v42 = m.ExcPending
			if v42 != 0 {
				return int32(0)
			} else {
				return v41
			}
		case 4:
			v46 = *(*int32)(unsafe.Add(mBase, uint32(v13+int32(-17))))
			v47 = v46
			v48 = F_rioWriteBulkString(m, l0, v11, v47)
			mBase = m.M
			v49 = m.ExcPending
			if v49 != 0 {
				return int32(0)
			} else {
				return v48
			}
		default:
			v47 = int32(0)
			v48 = F_rioWriteBulkString(m, l0, v11, v47)
			mBase = m.M
			v49 = m.ExcPending
			if v49 != 0 {
				return int32(0)
			} else {
				return v48
			}
		}
	case 1:
		v58 = F_objectGetVal(m, l1)
		mBase = m.M
		v60 = F_rioWriteBulkLongLong(m, l0, base.I64_extend_i32_s(v58))
		mBase = m.M
		v61 = m.ExcPending
		if v61 != 0 {
			return int32(0)
		} else {
			return v60
		}
	default:
		F__serverPanic_1(m, int32(_a85), int32(1916), int32(_a166), int32(0))
		mBase = m.M
		v56 = m.ExcPending
		if v56 != 0 {
			return int32(0)
		} else {
			F_abort(m)
			mBase = m.M
			base.Wasm_trap_unreachable()
			for {
			}
		}
	}
}
func F_rioWriteHashIteratorCursor(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v25 int32
	_ = v25
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
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v42 int32
	_ = v42
	var v44 int64
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	switch v9 + int32(-2) {
	case 0:
		v32 = F_hashTypeCurrentFromHashTable(m, l1, l2, v7)
		mBase = m.M
		v33 = m.ExcPending
		if v33 != 0 {
			return int32(0)
		} else {
			v34 = *(*int32)(unsafe.Add(mBase, uint32(v7)))
			v35 = F_rioWriteBulkString(m, l0, v32, v34)
			mBase = m.M
			v36 = m.ExcPending
			if v36 != 0 {
				return int32(0)
			} else {
				v47 = v35
				m.G0 = v7 + int32(16)
				return v47
			}
		}
	default:
		F__serverPanic_1(m, int32(_a85), int32(2082), int32(_a176), int32(0))
		mBase = m.M
		v42 = m.ExcPending
		if v42 != 0 {
			return int32(0)
		} else {
			F_abort(m)
			mBase = m.M
			base.Wasm_trap_unreachable()
			for {
			}
		}
	case 9:
		*(*int32)(unsafe.Add(mBase, uint32(v7)+12)) = int32(0)
		*(*int32)(unsafe.Add(mBase, uint32(v7)+8)) = int32(-1)
		*(*int64)(unsafe.Add(mBase, uint32(v7))) = int64(9223372036854775807)
		F_hashTypeCurrentFromListpack(m, l1, l2, v7+int32(12), v7+int32(8), v7)
		mBase = m.M
		v25 = m.ExcPending
		if v25 != 0 {
			return int32(0)
		} else {
			v26 = *(*int32)(unsafe.Add(mBase, uint32(v7)+12))
			if v26 == int32(0) {
				v44 = *(*int64)(unsafe.Add(mBase, uint32(v7)))
				v45 = F_rioWriteBulkLongLong(m, l0, v44)
				mBase = m.M
				v46 = m.ExcPending
				if v46 != 0 {
					return int32(0)
				} else {
					v47 = v45
					m.G0 = v7 + int32(16)
					return v47
				}
			} else {
				v29 = *(*int32)(unsafe.Add(mBase, uint32(v7)+8))
				v30 = F_rioWriteBulkString(m, l0, v26, v29)
				mBase = m.M
				v31 = m.ExcPending
				if v31 != 0 {
					return int32(0)
				} else {
					v47 = v30
					m.G0 = v7 + int32(16)
					return v47
				}
			}
		}
	}
}
