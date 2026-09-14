package p2

import (
	base "github.com/shibukawa/vkmem/internal/aot/vkaot/base"
	"unsafe"
)

func F_intsetAdd(m *base.Module, l0 int32, l1 int64, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int64
	_ = v2
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v43 int64
	_ = v43
	var v44 int64
	_ = v44
	var v45 int64
	_ = v45
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v61 int32
	_ = v61
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v82 int64
	_ = v82
	var v86 int64
	_ = v86
	var v90 int64
	_ = v90
	var v91 int64
	_ = v91
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v121 int32
	_ = v121
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v150 int32
	_ = v150
	var v155 int32
	_ = v155
	var v161 int32
	_ = v161
	var v165 int32
	_ = v165
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v176 int64
	_ = v176
	var v180 int64
	_ = v180
	var v184 int64
	_ = v184
	var v185 int64
	_ = v185
	var v191 int64
	_ = v191
	var v192 int64
	_ = v192
	var v193 int64
	_ = v193
	var v194 int64
	_ = v194
	var v196 int32
	_ = v196
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v215 int32
	_ = v215
	var v219 int64
	_ = v219
	var v223 int64
	_ = v223
	var v227 int64
	_ = v227
	var v228 int64
	_ = v228
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
	var v238 int32
	_ = v238
	var v241 int32
	_ = v241
	var v245 int64
	_ = v245
	var v246 int32
	_ = v246
	var v249 int32
	_ = v249
	var v255 int32
	_ = v255
	var v260 int32
	_ = v260
	var v266 int32
	_ = v266
	var v267 int32
	_ = v267
	var v276 int32
	_ = v276
	var v284 int32
	_ = v284
	var v286 int32
	_ = v286
	var v291 int64
	_ = v291
	var v297 int32
	_ = v297
	var v298 int32
	_ = v298
	var v299 int32
	_ = v299
	var v300 int32
	_ = v300
	var v301 int32
	_ = v301
	var v304 int32
	_ = v304
	var v305 int32
	_ = v305
	var v307 int32
	_ = v307
	var v310 int32
	_ = v310
	var v317 int32
	_ = v317
	var v324 int32
	_ = v324
	var v330 int32
	_ = v330
	var v331 int32
	_ = v331
	var v332 int32
	_ = v332
	var v334 int32
	_ = v334
	var v338 int32
	_ = v338
	var v345 int32
	_ = v345
	var v348 int32
	_ = v348
	var v355 int32
	_ = v355
	var v356 int32
	_ = v356
	var v357 int32
	_ = v357
	var v361 int32
	_ = v361
	var v363 int32
	_ = v363
	var v364 int32
	_ = v364
	var v366 int32
	_ = v366
	var v368 int32
	_ = v368
	var v379 int32
	_ = v379
	var v385 int32
	_ = v385
	var v386 int32
	_ = v386
	var v388 int32
	_ = v388
	var v394 int32
	_ = v394
	var v401 int32
	_ = v401
	var v405 int32
	_ = v405
	var v408 int32
	_ = v408
	var v414 int32
	_ = v414
	var v421 int32
	_ = v421
	var v425 int32
	_ = v425
	var v428 int32
	_ = v428
	var v431 int32
	_ = v431
	var v432 int32
	_ = v432
	var v433 int32
	_ = v433
	var v438 int32
	_ = v438
	var v439 int32
	_ = v439
	var v440 int32
	_ = v440
	var v442 int32
	_ = v442
	var v444 int32
	_ = v444
	var v445 int32
	_ = v445
	var v447 int32
	_ = v447
	var v449 int32
	_ = v449
	var v453 int32
	_ = v453
	var v454 int32
	_ = v454
	var v455 int32
	_ = v455
	var v460 int32
	_ = v460
	var v461 int32
	_ = v461
	var v462 int32
	_ = v462
	var v464 int32
	_ = v464
	var v466 int32
	_ = v466
	var v471 int32
	_ = v471
	var v483 int32
	_ = v483
	var v484 int32
	_ = v484
	var v502 int32
	_ = v502
	var v509 int32
	_ = v509
	var v513 int32
	_ = v513
	var v532 int32
	_ = v532
	var v538 int32
	_ = v538
	v2 = l1
	v12 = m.G0
	v14 = v12 - int32(16)
	m.G0 = v14
	if l2 == int32(0) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v29 = base.B2i32(base.Ui64(v2+int64(-32768)) < base.Ui64(int64(-65536)))
	if base.Ui64(v2+int64(-32768)) < base.Ui64(int64(-65536)) {
		goto L7
	} else {
		goto L8
	}
L2:
	;
	v22 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l2))) = uint8(v22)
	goto L1
L3:
	;
	F__serverAssert(m, int32(_a_F_intsetAdd_0), int32(_a_F_intsetAdd_1), int32(108))
	mBase = m.M
	v538 = m.ExcPending
	if v538 != 0 {
		goto L21
	} else {
		goto L137
	}
L4:
	;
	F__serverAssert(m, int32(_a_F_intsetAdd_0), int32(_a_F_intsetAdd_1), int32(108))
	mBase = m.M
	v532 = m.ExcPending
	if v532 != 0 {
		goto L21
	} else {
		goto L136
	}
L5:
	;
	m.G0 = v14 + int32(16)
	return v513
L6:
	;
	v155 = v14 + int32(12)
	v161 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v161 == int32(0) {
		goto L50
	} else {
		goto L51
	}
L7:
	;
	v30 = int32(4)
	goto L9
L8:
	;
	v30 = int32(2)
	goto L9
L9:
	;
	v32 = base.B2i32(base.Ui64(v2+int64(-2147483648)) < base.Ui64(int64(-4294967296)))
	if base.Ui64(v2+int64(-2147483648)) < base.Ui64(int64(-4294967296)) {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	v33 = int32(8)
	goto L12
L11:
	;
	v33 = v30
	goto L12
L12:
	;
	if base.Ui32(v33) <= base.Ui32(v24) {
		goto L6
	} else {
		goto L13
	}
L13:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v33
	v36 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if base.Ui64(v2+int64(-32768)) < base.Ui64(int64(-65536)) {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v43 = int64(2)
	goto L16
L15:
	;
	v43 = int64(1)
	goto L16
L16:
	;
	if base.Ui64(v2+int64(-2147483648)) < base.Ui64(int64(-4294967296)) {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	v44 = int64(3)
	goto L19
L18:
	;
	v44 = v43
	goto L19
L19:
	;
	v45 = base.I64_extend_i32_u(v36+int32(1)) << (uint(v44) % 64)
	if base.Ui64(int64(4294967288)) <= base.Ui64(v45) {
		goto L4
	} else {
		goto L20
	}
L20:
	;
	v51 = F_valkey_realloc(m, l0, base.I32_wrap_i64(v45)+int32(8))
	mBase = m.M
	v54 = m.ExcPending
	if v54 != 0 {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	return int32(0)
L22:
	;
	if v36 == int32(0) {
		goto L23
	} else {
		goto L24
	}
L23:
	;
	if int64(-1) < v2 {
		goto L37
	} else {
		goto L38
	}
L24:
	;
	v61 = v51 + int32(8)
	v75 = v36
	goto L25
L25:
	;
	v78 = v75 + int32(-1)
	switch v24&int32(255) + int32(-4) {
	case 0:
		goto L29
	default:
		goto L28
	case 4:
		goto L30
	}
L26:
	;
	goto L23
L27:
	;
	v92 = v78 + base.I32_wrap_i64(int64(base.Ui64(v2)>>(uint(int64(63))%64)))
	v93 = *(*int32)(unsafe.Add(mBase, uint32(v51)))
	switch v93 + int32(-4) {
	case 0:
		goto L33
	default:
		goto L32
	case 4:
		goto L34
	}
L28:
	;
	v90 = int64(*(*int16)(unsafe.Add(mBase, uint32(v61+v78<<(uint(int32(1))%32)))))
	v91 = v90
	goto L27
L29:
	;
	v86 = int64(*(*int32)(unsafe.Add(mBase, uint32(v61+v78<<(uint(int32(2))%32)))))
	v91 = v86
	goto L27
L30:
	;
	v82 = *(*int64)(unsafe.Add(mBase, uint32(v51+v75<<(uint(int32(3))%32))))
	v91 = v82
	goto L27
L31:
	;
	if v78 != 0 {
		v75 = v78
		goto L25
	} else {
		goto L35
	}
L32:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v61+v92<<(uint(int32(1))%32)))) = uint16(v91)
	goto L31
L33:
	;
	*(*uint32)(unsafe.Add(mBase, uint32(v61+v92<<(uint(int32(2))%32)))) = uint32(v91)
	goto L31
L34:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v61+v92<<(uint(int32(3))%32)))) = v91
	goto L31
L35:
	;
	goto L26
L36:
	;
	v150 = *(*int32)(unsafe.Add(mBase, uint32(v51)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v51)+4)) = v150 + int32(1)
	v513 = v51
	goto L5
L37:
	;
	v127 = *(*int32)(unsafe.Add(mBase, uint32(v51)+4))
	v128 = *(*int32)(unsafe.Add(mBase, uint32(v51)))
	switch v128 + int32(-4) {
	case 0:
		goto L43
	default:
		goto L42
	case 4:
		goto L44
	}
L38:
	;
	v121 = *(*int32)(unsafe.Add(mBase, uint32(v51)))
	switch v121 + int32(-4) {
	case 0:
		goto L40
	default:
		goto L39
	case 4:
		goto L41
	}
L39:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v51)+8)) = uint16(v2)
	goto L36
L40:
	;
	*(*uint32)(unsafe.Add(mBase, uint32(v51)+8)) = uint32(v2)
	goto L36
L41:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v51)+8)) = v2
	goto L36
L42:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v51+v127<<(uint(int32(1))%32)+int32(8)))) = uint16(v2)
	goto L36
L43:
	;
	*(*uint32)(unsafe.Add(mBase, uint32(v51+v127<<(uint(int32(2))%32)+int32(8)))) = uint32(v2)
	goto L36
L44:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v51+v127<<(uint(int32(3))%32)+int32(8)))) = v2
	goto L36
L45:
	;
	v286 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v291 = base.I64_extend_i32_u(v286+int32(1)) * base.I64_extend_i32_u(v24)
	if base.Ui64(int64(4294967288)) <= base.Ui64(v291) {
		goto L3
	} else {
		goto L83
	}
L46:
	;
	if v276 == int32(0) {
		goto L45
	} else {
		goto L81
	}
L47:
	;
	goto L46
L48:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v155))) = v266
	v276 = v267
	goto L47
L49:
	;
	v260 = int32(0)
	if v155 == v260 {
		v276 = v260
		goto L47
	} else {
		goto L80
	}
L50:
	;
	v255 = int32(0)
	goto L49
L51:
	;
	v165 = v161 + int32(-1)
	v167 = l0 + int32(8)
	v168 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	switch v168&int32(255) + int32(-4) {
	case 0:
		goto L54
	default:
		goto L53
	case 4:
		goto L55
	}
L52:
	;
	if v185 < v2 {
		v255 = v161
		goto L49
	} else {
		goto L56
	}
L53:
	;
	v184 = int64(*(*int16)(unsafe.Add(mBase, uint32(v167+v165<<(uint(int32(1))%32)))))
	v185 = v184
	goto L52
L54:
	;
	v180 = int64(*(*int32)(unsafe.Add(mBase, uint32(v167+v165<<(uint(int32(2))%32)))))
	v185 = v180
	goto L52
L55:
	;
	v176 = *(*int64)(unsafe.Add(mBase, uint32(l0+v161<<(uint(int32(3))%32))))
	v185 = v176
	goto L52
L56:
	;
	switch v168&int32(255) + int32(-4) {
	case 0:
		goto L59
	default:
		goto L58
	case 4:
		goto L60
	}
L57:
	;
	if v2 < v194 {
		goto L50
	} else {
		goto L61
	}
L58:
	;
	v193 = int64(*(*int16)(unsafe.Add(mBase, uint32(v167))))
	v194 = v193
	goto L57
L59:
	;
	v192 = int64(*(*int32)(unsafe.Add(mBase, uint32(v167))))
	v194 = v192
	goto L57
L60:
	;
	v191 = *(*int64)(unsafe.Add(mBase, uint32(v167)))
	v194 = v191
	goto L57
L61:
	;
	v196 = int32(0)
	if v196 <= v165 {
		goto L63
	} else {
		goto L64
	}
L62:
	;
	v246 = base.B2i32(v2 == v245)
	if v155 == int32(0) {
		v276 = v246
		goto L47
	} else {
		goto L76
	}
L63:
	;
	v208 = v196
	v209 = v165
	goto L65
L64:
	;
	v238 = int32(-1)
	v241 = v196
	v245 = int64(-1)
	goto L62
L65:
	;
	v215 = int32(base.Ui32(v209+v208) >> (uint(int32(1)) % 32))
	switch v168&int32(255) + int32(-4) {
	case 0:
		goto L69
	default:
		goto L68
	case 4:
		goto L70
	}
L66:
	;
	v238 = v215
	v241 = v235
	v245 = v228
	goto L62
L67:
	;
	if v2 <= v228 {
		goto L72
	} else {
		goto L73
	}
L68:
	;
	v227 = int64(*(*int16)(unsafe.Add(mBase, uint32(v167+v215<<(uint(int32(1))%32)))))
	v228 = v227
	goto L67
L69:
	;
	v223 = int64(*(*int32)(unsafe.Add(mBase, uint32(v167+v215<<(uint(int32(2))%32)))))
	v228 = v223
	goto L67
L70:
	;
	v219 = *(*int64)(unsafe.Add(mBase, uint32(v167+v215<<(uint(int32(3))%32))))
	v228 = v219
	goto L67
L71:
	;
	if v235 <= v236 {
		v208 = v235
		v209 = v236
		goto L65
	} else {
		goto L75
	}
L72:
	;
	if v228 <= v2 {
		v238 = v215
		v241 = v208
		v245 = v228
		goto L62
	} else {
		goto L74
	}
L73:
	;
	v235 = v215 + int32(1)
	v236 = v209
	goto L71
L74:
	;
	v235 = v208
	v236 = v215 + int32(-1)
	goto L71
L75:
	;
	goto L66
L76:
	;
	if v2 == v245 {
		goto L77
	} else {
		goto L78
	}
L77:
	;
	v249 = v238
	goto L79
L78:
	;
	v249 = v241
	goto L79
L79:
	;
	v266 = v249
	v267 = v246
	goto L48
L80:
	;
	v266 = v255
	v267 = v260
	goto L48
L81:
	;
	if l2 == int32(0) {
		v513 = l0
		goto L5
	} else {
		goto L82
	}
L82:
	;
	v284 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l2))) = uint8(v284)
	v513 = l0
	goto L5
L83:
	;
	v297 = F_valkey_realloc(m, l0, base.I32_wrap_i64(v291)+int32(8))
	mBase = m.M
	v298 = m.ExcPending
	if v298 != 0 {
		goto L21
	} else {
		goto L84
	}
L84:
	;
	v299 = *(*int32)(unsafe.Add(mBase, uint32(v297)))
	v300 = *(*int32)(unsafe.Add(mBase, uint32(v297)+4))
	v301 = *(*int32)(unsafe.Add(mBase, uint32(v14)+12))
	if base.Ui32(v300) <= base.Ui32(v301) {
		v484 = v301
		goto L85
	} else {
		goto L86
	}
L85:
	;
	switch v299 + int32(-4) {
	case 0:
		goto L134
	default:
		goto L133
	case 4:
		goto L135
	}
L86:
	;
	v304 = v297 + int32(8)
	v305 = int32(1)
	v307 = v301 + v305
	switch v299 + int32(-4) {
	case 0:
		goto L89
	default:
		goto L88
	case 4:
		goto L90
	}
L87:
	;
	v334 = (v300 - v301) << (uint(v330) % 32)
	if v332 == v331 {
		goto L92
	} else {
		goto L93
	}
L88:
	;
	v324 = int32(1)
	v330 = v305
	v331 = v304 + v301<<(uint(v324)%32)
	v332 = v304 + v307<<(uint(v324)%32)
	goto L87
L89:
	;
	v317 = int32(2)
	v330 = v317
	v331 = v304 + v301<<(uint(v317)%32)
	v332 = v304 + v307<<(uint(v317)%32)
	goto L87
L90:
	;
	v310 = int32(3)
	v330 = v310
	v331 = v304 + v301<<(uint(v310)%32)
	v332 = v304 + v307<<(uint(v310)%32)
	goto L87
L91:
	;
	v483 = *(*int32)(unsafe.Add(mBase, uint32(v14)+12))
	v484 = v483
	goto L85
L92:
	;
	goto L91
L93:
	;
	v338 = v334 + v332
	if base.Ui32(int32(0)-v334<<(uint(int32(1))%32)) < base.Ui32(v331-v338) {
		goto L94
	} else {
		goto L95
	}
L94:
	;
	v348 = (v331 ^ v332) & int32(3)
	if base.Ui32(v331) <= base.Ui32(v332) {
		goto L98
	} else {
		goto L99
	}
L95:
	;
	v345 = F___memcpy(m, v332, v331, v334)
	mBase = m.M
	goto L91
L96:
	;
	if v454 == int32(0) {
		goto L92
	} else {
		goto L128
	}
L97:
	;
	if base.Ui32(v432) <= base.Ui32(int32(3)) {
		v453 = v431
		v454 = v432
		v455 = v433
		goto L96
	} else {
		goto L124
	}
L98:
	;
	if v348 != 0 {
		v414 = v334
		goto L108
	} else {
		goto L109
	}
L99:
	;
	if v348 == int32(0) {
		goto L100
	} else {
		goto L101
	}
L100:
	;
	if v332&int32(3) != 0 {
		goto L102
	} else {
		goto L103
	}
L101:
	;
	v453 = v331
	v454 = v334
	v455 = v332
	goto L96
L102:
	;
	v355 = v331
	v356 = v334
	v357 = v332
	goto L104
L103:
	;
	v431 = v331
	v432 = v334
	v433 = v332
	goto L97
L104:
	;
	if v356 == int32(0) {
		goto L92
	} else {
		goto L106
	}
L106:
	;
	v361 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v355))))
	*(*uint8)(unsafe.Add(mBase, uint32(v357))) = uint8(v361)
	v363 = int32(1)
	v364 = v355 + v363
	v366 = v356 + int32(-1)
	v368 = v357 + v363
	if v368&int32(3) == int32(0) {
		v431 = v364
		v432 = v366
		v433 = v368
		goto L97
	} else {
		goto L107
	}
L107:
	;
	v355 = v364
	v356 = v366
	v357 = v368
	goto L104
L108:
	;
	if v414 == int32(0) {
		goto L92
	} else {
		goto L120
	}
L109:
	;
	if v338&int32(3) == int32(0) {
		v394 = v334
		goto L110
	} else {
		goto L111
	}
L110:
	;
	if base.Ui32(v394) <= base.Ui32(int32(3)) {
		v414 = v394
		goto L108
	} else {
		goto L116
	}
L111:
	;
	v379 = v334
	goto L112
L112:
	;
	if v379 == int32(0) {
		goto L92
	} else {
		goto L114
	}
L113:
	;
	v394 = v385
	goto L110
L114:
	;
	v385 = v379 + int32(-1)
	v386 = v332 + v385
	v388 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v331+v385))))
	*(*uint8)(unsafe.Add(mBase, uint32(v386))) = uint8(v388)
	if v386&int32(3) != 0 {
		v379 = v385
		goto L112
	} else {
		goto L115
	}
L115:
	;
	goto L113
L116:
	;
	v401 = v394
	goto L117
L117:
	;
	v405 = v401 + int32(-4)
	v408 = *(*int32)(unsafe.Add(mBase, uint32(v331+v405)))
	*(*int32)(unsafe.Add(mBase, uint32(v332+v405))) = v408
	if base.Ui32(int32(3)) < base.Ui32(v405) {
		v401 = v405
		goto L117
	} else {
		goto L119
	}
L118:
	;
	v414 = v405
	goto L108
L119:
	;
	goto L118
L120:
	;
	v421 = v414
	goto L121
L121:
	;
	v425 = v421 + int32(-1)
	v428 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v331+v425))))
	*(*uint8)(unsafe.Add(mBase, uint32(v332+v425))) = uint8(v428)
	if v425 != 0 {
		v421 = v425
		goto L121
	} else {
		goto L123
	}
L123:
	;
	goto L92
L124:
	;
	v438 = v431
	v439 = v432
	v440 = v433
	goto L125
L125:
	;
	v442 = *(*int32)(unsafe.Add(mBase, uint32(v438)))
	*(*int32)(unsafe.Add(mBase, uint32(v440))) = v442
	v444 = int32(4)
	v445 = v438 + v444
	v447 = v440 + v444
	v449 = v439 + int32(-4)
	if base.Ui32(int32(3)) < base.Ui32(v449) {
		v438 = v445
		v439 = v449
		v440 = v447
		goto L125
	} else {
		goto L127
	}
L126:
	;
	v453 = v445
	v454 = v449
	v455 = v447
	goto L96
L127:
	;
	goto L126
L128:
	;
	v460 = v453
	v461 = v454
	v462 = v455
	goto L129
L129:
	;
	v464 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v460))))
	*(*uint8)(unsafe.Add(mBase, uint32(v462))) = uint8(v464)
	v466 = int32(1)
	v471 = v461 + int32(-1)
	if v471 != 0 {
		v460 = v460 + v466
		v461 = v471
		v462 = v462 + v466
		goto L129
	} else {
		goto L131
	}
L130:
	;
	goto L92
L131:
	;
	goto L130
L132:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v297)+4)) = v509 + int32(1)
	v513 = v297
	goto L5
L133:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v297+v484<<(uint(int32(1))%32)+int32(8)))) = uint16(v2)
	v509 = v300
	goto L132
L134:
	;
	*(*uint32)(unsafe.Add(mBase, uint32(v297+v484<<(uint(int32(2))%32)+int32(8)))) = uint32(v2)
	v502 = *(*int32)(unsafe.Add(mBase, uint32(v297)+4))
	v509 = v502
	goto L132
L135:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v297+v484<<(uint(int32(3))%32)+int32(8)))) = v2
	v509 = v300
	goto L132
L136:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L137:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_intsetDup(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	if l0 != 0 {
		v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		v15 = v11*v12 + int32(8)
		v16 = F_valkey_malloc(m, v15)
		mBase = m.M
		v17 = m.ExcPending
		if v17 != 0 {
			return int32(0)
		} else {
			if v15 == int32(0) {
				v21 = v16
			} else {
				v20 = F__emscripten_memcpy_bulkmem(m, v16, l0, v15)
				mBase = m.M
				v21 = v20
			}
			return v21
		}
	} else {
		v4 = F_valkey_malloc(m, int32(8))
		mBase = m.M
		v7 = m.ExcPending
		if v7 != 0 {
			return int32(0)
		} else {
			*(*int64)(unsafe.Add(mBase, uint32(v4))) = int64(2)
			return v4
		}
	}
}
func F_intsetFind(m *base.Module, l0 int32, l1 int64) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v13 int32
	_ = v13
	var v18 int32
	_ = v18
	var v26 int32
	_ = v26
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v41 int64
	_ = v41
	var v45 int64
	_ = v45
	var v49 int64
	_ = v49
	var v50 int64
	_ = v50
	var v56 int64
	_ = v56
	var v57 int64
	_ = v57
	var v58 int64
	_ = v58
	var v59 int64
	_ = v59
	var v61 int32
	_ = v61
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v80 int32
	_ = v80
	var v84 int64
	_ = v84
	var v88 int64
	_ = v88
	var v92 int64
	_ = v92
	var v93 int64
	_ = v93
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v110 int64
	_ = v110
	var v141 int32
	_ = v141
	var v147 int32
	_ = v147
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if base.Ui64(l1+int64(-32768)) < base.Ui64(int64(-65536)) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	return v147
L2:
	;
	v13 = int32(4)
	goto L4
L3:
	;
	v13 = int32(2)
	goto L4
L4:
	;
	if base.Ui64(l1+int64(-2147483648)) < base.Ui64(int64(-4294967296)) {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	v18 = int32(8)
	goto L7
L6:
	;
	v18 = v13
	goto L7
L7:
	;
	if base.Ui32(v5) < base.Ui32(v18) {
		v147 = int32(0)
		goto L1
	} else {
		goto L8
	}
L8:
	;
	v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v26 == int32(0) {
		goto L13
	} else {
		goto L14
	}
L9:
	;
	v147 = base.B2i32(v141 != int32(0))
	goto L1
L10:
	;
	goto L9
L12:
	;
	v141 = int32(0)
	goto L10
L13:
	;
	goto L12
L14:
	;
	v30 = v26 + int32(-1)
	v32 = l0 + int32(8)
	v33 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	switch v33&int32(255) + int32(-4) {
	case 0:
		goto L17
	default:
		goto L16
	case 4:
		goto L18
	}
L15:
	;
	if v50 < l1 {
		goto L12
	} else {
		goto L19
	}
L16:
	;
	v49 = int64(*(*int16)(unsafe.Add(mBase, uint32(v32+v30<<(uint(int32(1))%32)))))
	v50 = v49
	goto L15
L17:
	;
	v45 = int64(*(*int32)(unsafe.Add(mBase, uint32(v32+v30<<(uint(int32(2))%32)))))
	v50 = v45
	goto L15
L18:
	;
	v41 = *(*int64)(unsafe.Add(mBase, uint32(l0+v26<<(uint(int32(3))%32))))
	v50 = v41
	goto L15
L19:
	;
	switch v33&int32(255) + int32(-4) {
	case 0:
		goto L22
	default:
		goto L21
	case 4:
		goto L23
	}
L20:
	;
	if l1 < v59 {
		goto L13
	} else {
		goto L24
	}
L21:
	;
	v58 = int64(*(*int16)(unsafe.Add(mBase, uint32(v32))))
	v59 = v58
	goto L20
L22:
	;
	v57 = int64(*(*int32)(unsafe.Add(mBase, uint32(v32))))
	v59 = v57
	goto L20
L23:
	;
	v56 = *(*int64)(unsafe.Add(mBase, uint32(v32)))
	v59 = v56
	goto L20
L24:
	;
	v61 = int32(0)
	if v61 <= v30 {
		goto L26
	} else {
		goto L27
	}
L25:
	;
	v141 = base.B2i32(l1 == v110)
	goto L10
L26:
	;
	v73 = v61
	v74 = v30
	goto L28
L27:
	;
	v110 = int64(-1)
	goto L25
L28:
	;
	v80 = int32(base.Ui32(v74+v73) >> (uint(int32(1)) % 32))
	switch v33&int32(255) + int32(-4) {
	case 0:
		goto L32
	default:
		goto L31
	case 4:
		goto L33
	}
L29:
	;
	v110 = v93
	goto L25
L30:
	;
	if l1 <= v93 {
		goto L35
	} else {
		goto L36
	}
L31:
	;
	v92 = int64(*(*int16)(unsafe.Add(mBase, uint32(v32+v80<<(uint(int32(1))%32)))))
	v93 = v92
	goto L30
L32:
	;
	v88 = int64(*(*int32)(unsafe.Add(mBase, uint32(v32+v80<<(uint(int32(2))%32)))))
	v93 = v88
	goto L30
L33:
	;
	v84 = *(*int64)(unsafe.Add(mBase, uint32(v32+v80<<(uint(int32(3))%32))))
	v93 = v84
	goto L30
L34:
	;
	if v100 <= v101 {
		v73 = v100
		v74 = v101
		goto L28
	} else {
		goto L38
	}
L35:
	;
	if v93 <= l1 {
		v110 = v93
		goto L25
	} else {
		goto L37
	}
L36:
	;
	v100 = v80 + int32(1)
	v101 = v74
	goto L34
L37:
	;
	v100 = v73
	v101 = v80 + int32(-1)
	goto L34
L38:
	;
	goto L29
}
func F_intsetFree(m *base.Module, l0 int32) {
	var v5 int32
	_ = v5
	if l0 == int32(0) {
		return
	} else {
		F_valkey_free(m, l0)
		v5 = m.ExcPending
		if v5 != 0 {
			return
		} else {
			return
		}
	}
}
func F_intsetGet(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v17 int64
	_ = v17
	var v21 int64
	_ = v21
	var v25 int64
	_ = v25
	var v26 int64
	_ = v26
	var v29 int32
	_ = v29
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if base.Ui32(v7) <= base.Ui32(l1) {
		v29 = int32(0)
	} else {
		v10 = l0 + int32(8)
		v11 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
		switch v11 + int32(-4) {
		case 0:
			v21 = int64(*(*int32)(unsafe.Add(mBase, uint32(v10+l1<<(uint(int32(2))%32)))))
			v26 = v21
		default:
			v25 = int64(*(*int16)(unsafe.Add(mBase, uint32(v10+l1<<(uint(int32(1))%32)))))
			v26 = v25
		case 4:
			v17 = *(*int64)(unsafe.Add(mBase, uint32(v10+l1<<(uint(int32(3))%32))))
			v26 = v17
		}
		*(*int64)(unsafe.Add(mBase, uint32(l2))) = v26
		v29 = int32(1)
	}
	return v29
}
func F_intsetMax(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v16 int64
	_ = v16
	var v21 int64
	_ = v21
	var v26 int64
	_ = v26
	v6 = l0 + int32(8)
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v9 = v7 + int32(-1)
	v10 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	switch v10 + int32(-4) {
	case 0:
		v21 = int64(*(*int32)(unsafe.Add(mBase, uint32(v6+v9<<(uint(int32(2))%32)))))
		return v21
	default:
		v26 = int64(*(*int16)(unsafe.Add(mBase, uint32(v6+v9<<(uint(int32(1))%32)))))
		return v26
	case 4:
		v16 = *(*int64)(unsafe.Add(mBase, uint32(l0+v7<<(uint(int32(3))%32))))
		return v16
	}
}
func F_intsetValidateIntegrity(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v19 int32
	_ = v19
	var v29 int32
	_ = v29
	var v34 int64
	_ = v34
	var v35 int64
	_ = v35
	var v36 int64
	_ = v36
	var v37 int64
	_ = v37
	var v48 int32
	_ = v48
	var v52 int64
	_ = v52
	var v57 int64
	_ = v57
	var v61 int64
	_ = v61
	var v65 int64
	_ = v65
	var v66 int64
	_ = v66
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v80 int32
	_ = v80
	if base.Ui32(l1) < base.Ui32(int32(8)) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	return v80
L2:
	;
	v80 = int32(0)
	goto L1
L3:
	;
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if base.Ui32(int32(8)) < base.Ui32(v10) {
		goto L2
	} else {
		goto L4
	}
L4:
	;
	if int32(1)<<(uint(v10)%32)&int32(276) == int32(0) {
		goto L2
	} else {
		goto L5
	}
L5:
	;
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v19 == int32(0) {
		goto L2
	} else {
		goto L6
	}
L6:
	;
	if v19*v10+int32(8) != l1 {
		goto L2
	} else {
		goto L7
	}
L7:
	;
	if l2 != 0 {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v29 = l0 + int32(8)
	switch v10&int32(255) + int32(-4) {
	case 0:
		goto L12
	default:
		goto L11
	case 4:
		goto L13
	}
L9:
	;
	return int32(1)
L10:
	;
	if base.Ui32(int32(2)) <= base.Ui32(v19) {
		goto L14
	} else {
		goto L15
	}
L11:
	;
	v36 = int64(*(*int16)(unsafe.Add(mBase, uint32(v29))))
	v37 = v36
	goto L10
L12:
	;
	v35 = int64(*(*int32)(unsafe.Add(mBase, uint32(v29))))
	v37 = v35
	goto L10
L13:
	;
	v34 = *(*int64)(unsafe.Add(mBase, uint32(v29)))
	v37 = v34
	goto L10
L14:
	;
	v48 = int32(1)
	v52 = v37
	goto L16
L15:
	;
	return int32(1)
L16:
	;
	switch v10&int32(255) + int32(-4) {
	case 0:
		goto L20
	default:
		goto L19
	case 4:
		goto L21
	}
L18:
	;
	if v66 <= v52 {
		goto L2
	} else {
		goto L22
	}
L19:
	;
	v65 = int64(*(*int16)(unsafe.Add(mBase, uint32(v29+v48<<(uint(int32(1))%32)))))
	v66 = v65
	goto L18
L20:
	;
	v61 = int64(*(*int32)(unsafe.Add(mBase, uint32(v29+v48<<(uint(int32(2))%32)))))
	v66 = v61
	goto L18
L21:
	;
	v57 = *(*int64)(unsafe.Add(mBase, uint32(v29+v48<<(uint(int32(3))%32))))
	v66 = v57
	goto L18
L22:
	;
	v68 = int32(1)
	v70 = v48 + v68
	if v70 == v19 {
		v80 = v68
		goto L1
	} else {
		goto L23
	}
L23:
	;
	v48 = v70
	v52 = v66
	goto L16
}
