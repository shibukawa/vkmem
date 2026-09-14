package p1

import (
	base "github.com/shibukawa/vkmem/internal/aot/vkaot/base"
	"unsafe"
)

func F___bswap_32_2(m *base.Module, l0 int32) int32 {
	var v2 int32
	_ = v2
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	v2 = int32(24)
	v4 = int32(65280)
	v6 = int32(8)
	return l0<<(uint(v2)%32) | l0&v4<<(uint(v6)%32) | (int32(base.Ui32(l0)>>(uint(v6)%32))&v4 | int32(base.Ui32(l0)>>(uint(v2)%32)))
}
func F_b_pack(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	var v35 int32
	_ = v35
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v90 int32
	_ = v90
	var v102 int32
	_ = v102
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
	var v125 int32
	_ = v125
	var v135 int32
	_ = v135
	var v136 float64
	_ = v136
	var v137 int32
	_ = v137
	var v145 int32
	_ = v145
	var v147 int32
	_ = v147
	var v155 int32
	_ = v155
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v163 int32
	_ = v163
	var v165 int32
	_ = v165
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v177 int32
	_ = v177
	var v185 int32
	_ = v185
	var v189 int32
	_ = v189
	var v191 int32
	_ = v191
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v226 int32
	_ = v226
	var v228 int32
	_ = v228
	var v230 int32
	_ = v230
	var v232 int32
	_ = v232
	var v252 int32
	_ = v252
	var v259 int32
	_ = v259
	var v263 int32
	_ = v263
	var v264 int32
	_ = v264
	var v274 int32
	_ = v274
	var v277 int32
	_ = v277
	var v282 int32
	_ = v282
	var v287 int32
	_ = v287
	var v292 int32
	_ = v292
	var v294 int32
	_ = v294
	var v295 int32
	_ = v295
	var v298 int32
	_ = v298
	var v303 int32
	_ = v303
	var v304 int32
	_ = v304
	var v319 int32
	_ = v319
	var v320 int32
	_ = v320
	var v324 int32
	_ = v324
	var v335 int32
	_ = v335
	var v340 int32
	_ = v340
	var v363 int32
	_ = v363
	var v364 int32
	_ = v364
	var v368 int32
	_ = v368
	var v369 int32
	_ = v369
	var v370 int32
	_ = v370
	var v371 int32
	_ = v371
	var v375 int32
	_ = v375
	var v377 float64
	_ = v377
	var v378 int32
	_ = v378
	var v383 int32
	_ = v383
	var v391 int32
	_ = v391
	var v392 int32
	_ = v392
	var v404 int32
	_ = v404
	var v405 int32
	_ = v405
	var v406 int32
	_ = v406
	var v409 int32
	_ = v409
	var v410 int32
	_ = v410
	var v414 int32
	_ = v414
	var v416 int32
	_ = v416
	var v439 int32
	_ = v439
	var v440 float64
	_ = v440
	var v441 int32
	_ = v441
	var v445 int32
	_ = v445
	var v453 int32
	_ = v453
	var v454 int32
	_ = v454
	var v466 int32
	_ = v466
	var v467 int32
	_ = v467
	var v468 int32
	_ = v468
	var v471 int32
	_ = v471
	var v472 int32
	_ = v472
	var v476 int32
	_ = v476
	var v478 int32
	_ = v478
	var v501 int32
	_ = v501
	var v503 int32
	_ = v503
	var v506 int32
	_ = v506
	var v507 int32
	_ = v507
	var v508 int32
	_ = v508
	var v509 int32
	_ = v509
	var v511 int32
	_ = v511
	var v514 int32
	_ = v514
	var v515 int32
	_ = v515
	var v519 int32
	_ = v519
	var v522 int32
	_ = v522
	var v526 int32
	_ = v526
	var v527 int32
	_ = v527
	var v528 int32
	_ = v528
	var v529 int32
	_ = v529
	var v530 int32
	_ = v530
	var v533 int32
	_ = v533
	var v538 int32
	_ = v538
	var v544 int32
	_ = v544
	var v550 int32
	_ = v550
	var v553 int32
	_ = v553
	var v560 int32
	_ = v560
	var v561 int32
	_ = v561
	var v579 int32
	_ = v579
	v15 = m.G0
	v17 = v15 - int32(1088)
	m.G0 = v17
	v21 = F_luaL_checklstring(m, l0, int32(1), int32(0))
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v17)+8)) = int64(4294967297)
	v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v28)+8)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v28 + int32(16)
	goto L3
L3:
	;
	v35 = v17 + int32(20)
	*(*int32)(unsafe.Add(mBase, uint32(v35)+4)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v35)+8)) = l0
	*(*int32)(unsafe.Add(mBase, uint32(v35))) = v17 + int32(32)
	goto L4
L4:
	;
	v42 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21))))
	if v42 == int32(0) {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	F_luaL_pushresult(m, v17+int32(20))
	mBase = m.M
	v579 = m.ExcPending
	if v579 != 0 {
		goto L1
	} else {
		goto L99
	}
L6:
	;
	v46 = v17 + int32(1056)
	v51 = v21
	v52 = v42
	v54 = int32(2)
	v55 = int32(0)
	goto L7
L7:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+16)) = v51 + int32(1)
	v66 = base.I32_extend8_s(v52)
	v69 = F_optsize(m, l0, v66, v17+int32(16))
	mBase = m.M
	v70 = m.ExcPending
	if v70 != 0 {
		goto L1
	} else {
		goto L9
	}
L8:
	;
	goto L5
L9:
	;
	v72 = v52 & int32(255)
	if v72 == int32(99) {
		v125 = v55
		goto L10
	} else {
		goto L11
	}
L10:
	;
	switch v72 + int32(-66) {
	case 0, 6, 7, 10, 18, 32, 38, 39, 42:
		goto L29
	default:
		goto L24
	case 33, 49:
		goto L25
	case 34:
		goto L26
	case 36:
		goto L27
	case 54:
		goto L28
	}
L11:
	;
	if v69 == int32(0) {
		v125 = v55
		goto L10
	} else {
		goto L12
	}
L12:
	;
	v77 = *(*int32)(unsafe.Add(mBase, uint32(v17)+12))
	if base.Ui32(v69) < base.Ui32(v77) {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	v79 = v69
	goto L15
L14:
	;
	v79 = v77
	goto L15
L15:
	;
	v81 = v79 + int32(-1)
	v84 = (v79 - v81&v55) & v81
	v85 = v84 + v55
	if v84 < int32(1) {
		v125 = v85
		goto L10
	} else {
		goto L16
	}
L16:
	;
	v90 = v84
	goto L17
L17:
	;
	v102 = *(*int32)(unsafe.Add(mBase, uint32(v17)+20))
	if base.Ui32(v102) < base.Ui32(v46) {
		v109 = v102
		goto L19
	} else {
		goto L20
	}
L18:
	;
	v125 = v85
	goto L10
L19:
	;
	v110 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v17)+20)) = v109 + v110
	v113 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v109))) = uint8(v113)
	if base.Ui32(v110) < base.Ui32(v90) {
		v90 = v90 + int32(-1)
		goto L17
	} else {
		goto L22
	}
L20:
	;
	v106 = F_luaL_prepbuffer(m, v17+int32(20))
	mBase = m.M
	v107 = m.ExcPending
	if v107 != 0 {
		goto L1
	} else {
		goto L21
	}
L21:
	;
	v108 = *(*int32)(unsafe.Add(mBase, uint32(v17)+20))
	v109 = v108
	goto L19
L22:
	;
	goto L18
L23:
	;
	v560 = *(*int32)(unsafe.Add(mBase, uint32(v17)+16))
	v561 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v560))))
	if v561 != 0 {
		v51 = v560
		v52 = v561
		v54 = v550
		v55 = v125 + v553
		goto L7
	} else {
		goto L98
	}
L24:
	;
	F_controloptions(m, l0, v66, v17+int32(16), v17+int32(8))
	mBase = m.M
	v544 = m.ExcPending
	if v544 != 0 {
		goto L1
	} else {
		goto L97
	}
L25:
	;
	v503 = v54 + int32(1)
	v506 = F_luaL_checklstring(m, l0, v54, v17+int32(1056))
	mBase = m.M
	v507 = m.ExcPending
	if v507 != 0 {
		goto L1
	} else {
		goto L84
	}
L26:
	;
	v440 = F_luaL_checknumber(m, l0, v54)
	mBase = m.M
	v441 = m.ExcPending
	if v441 != 0 {
		goto L1
	} else {
		goto L76
	}
L27:
	;
	v377 = F_luaL_checknumber(m, l0, v54)
	mBase = m.M
	v378 = m.ExcPending
	if v378 != 0 {
		goto L1
	} else {
		goto L68
	}
L28:
	;
	v364 = *(*int32)(unsafe.Add(mBase, uint32(v17)+20))
	if base.Ui32(v364) < base.Ui32(v46) {
		v371 = v364
		goto L65
	} else {
		goto L66
	}
L29:
	;
	v135 = *(*int32)(unsafe.Add(mBase, uint32(v17)+8))
	v136 = F_luaL_checknumber(m, l0, v54)
	mBase = m.M
	v137 = m.ExcPending
	if v137 != 0 {
		goto L1
	} else {
		goto L32
	}
L30:
	;
	if base.F64_lt(base.F64_abs(v136), float64(2.147483648e+09)) == int32(0) {
		goto L35
	} else {
		goto L36
	}
L31:
	;
	v147 = int32(0)
	goto L30
L32:
	;
	if base.F64_lt(v136, float64(4.294967296e+09))&base.F64_ge(v136, float64(0)) == int32(0) {
		goto L31
	} else {
		goto L33
	}
L33:
	;
	v145 = base.I32_trunc_f64_u(v136)
	v147 = v145
	goto L30
L34:
	;
	if base.F64_lt(v136, float64(0)) != 0 {
		goto L37
	} else {
		goto L38
	}
L35:
	;
	v157 = int32(-2147483648)
	goto L34
L36:
	;
	v155 = base.I32_trunc_f64_s(v136)
	v157 = v155
	goto L34
L37:
	;
	v158 = v157
	goto L39
L38:
	;
	v158 = v147
	goto L39
L39:
	;
	if v135 == int32(1) {
		goto L41
	} else {
		goto L42
	}
L40:
	;
	F_luaL_addlstring(m, v17+int32(20), v17+int32(1056), v69)
	mBase = m.M
	v363 = m.ExcPending
	if v363 != 0 {
		goto L1
	} else {
		goto L64
	}
L41:
	;
	if v69 < int32(1) {
		goto L40
	} else {
		goto L53
	}
L42:
	;
	if v69 < int32(1) {
		goto L40
	} else {
		goto L43
	}
L43:
	;
	v163 = int32(0)
	v165 = v69 & int32(3)
	if v165 == v163 {
		v195 = v69
		v196 = v158
		goto L44
	} else {
		goto L45
	}
L44:
	;
	if base.Ui32(v69) < base.Ui32(int32(4)) {
		goto L40
	} else {
		goto L49
	}
L45:
	;
	v170 = v69
	v171 = v158
	v177 = v163
	goto L46
L46:
	;
	v185 = v170 + int32(-1)
	*(*uint8)(unsafe.Add(mBase, uint32(v17+int32(1056)+v185))) = uint8(v171)
	v189 = int32(base.Ui32(v171) >> (uint(int32(8)) % 32))
	v191 = v177 + int32(1)
	if v191 != v165 {
		v170 = v185
		v171 = v189
		v177 = v191
		goto L46
	} else {
		goto L48
	}
L47:
	;
	v195 = v185
	v196 = v189
	goto L44
L48:
	;
	goto L47
L49:
	;
	v211 = v195
	v212 = v196
	goto L50
L50:
	;
	v226 = v211 + int32(-4)
	v228 = int32(24)
	v230 = int32(65280)
	v232 = int32(8)
	*(*int32)(unsafe.Add(mBase, uint32(v17+int32(1056)+v226))) = v212<<(uint(v228)%32) | v212&v230<<(uint(v232)%32) | (int32(base.Ui32(v212)>>(uint(v232)%32))&v230 | int32(base.Ui32(v212)>>(uint(v228)%32)))
	if base.Ui32(int32(1)) < base.Ui32(v211+int32(-3)) {
		v211 = v226
		v212 = int32(0)
		goto L50
	} else {
		goto L52
	}
L52:
	;
	goto L40
L53:
	;
	v252 = v69 & int32(3)
	if base.Ui32(int32(4)) <= base.Ui32(v69) {
		goto L55
	} else {
		goto L56
	}
L54:
	;
	if v252 == int32(0) {
		goto L40
	} else {
		goto L60
	}
L55:
	;
	v259 = int32(0)
	v263 = v259
	v264 = v158
	v274 = v259
	goto L57
L56:
	;
	v303 = int32(0)
	v304 = v158
	goto L54
L57:
	;
	v277 = v17 + int32(1056) + v263
	*(*uint8)(unsafe.Add(mBase, uint32(v277))) = uint8(v264)
	v282 = int32(base.Ui32(v264) >> (uint(int32(8)) % 32))
	*(*uint8)(unsafe.Add(mBase, uint32(v277+int32(1)))) = uint8(v282)
	v287 = int32(base.Ui32(v264) >> (uint(int32(16)) % 32))
	*(*uint8)(unsafe.Add(mBase, uint32(v277+int32(2)))) = uint8(v287)
	v292 = int32(base.Ui32(v264) >> (uint(int32(24)) % 32))
	*(*uint8)(unsafe.Add(mBase, uint32(v277+int32(3)))) = uint8(v292)
	v294 = int32(4)
	v295 = v263 + v294
	v298 = v274 + v294
	if v298 != v69&int32(2147483644) {
		v263 = v295
		v264 = int32(0)
		v274 = v298
		goto L57
	} else {
		goto L59
	}
L58:
	;
	v303 = v295
	v304 = int32(0)
	goto L54
L59:
	;
	goto L58
L60:
	;
	v319 = v303
	v320 = v304
	v324 = int32(0)
	goto L61
L61:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v17+int32(1056)+v319))) = uint8(v320)
	v335 = int32(1)
	v340 = v324 + v335
	if v340 != v252 {
		v319 = v319 + v335
		v320 = int32(base.Ui32(v320) >> (uint(int32(8)) % 32))
		v324 = v340
		goto L61
	} else {
		goto L63
	}
L62:
	;
	goto L40
L63:
	;
	goto L62
L64:
	;
	v550 = v54 + int32(1)
	v553 = v69
	goto L23
L65:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+20)) = v371 + int32(1)
	v375 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v371))) = uint8(v375)
	v550 = v54
	v553 = v69
	goto L23
L66:
	;
	v368 = F_luaL_prepbuffer(m, v17+int32(20))
	mBase = m.M
	v369 = m.ExcPending
	if v369 != 0 {
		goto L1
	} else {
		goto L67
	}
L67:
	;
	v370 = *(*int32)(unsafe.Add(mBase, uint32(v17)+20))
	v371 = v370
	goto L65
L68:
	;
	*(*float32)(unsafe.Add(mBase, uint32(v17)+1056)) = base.F32_demote_f64(v377)
	if v69 < int32(2) {
		goto L69
	} else {
		goto L70
	}
L69:
	;
	F_luaL_addlstring(m, v17+int32(20), v17+int32(1056), v69)
	mBase = m.M
	v439 = m.ExcPending
	if v439 != 0 {
		goto L1
	} else {
		goto L75
	}
L70:
	;
	v383 = *(*int32)(unsafe.Add(mBase, uint32(v17)+8))
	if v383 == int32(1) {
		goto L69
	} else {
		goto L71
	}
L71:
	;
	v391 = int32(0)
	v392 = v69 + int32(-1)
	goto L72
L72:
	;
	v404 = v17 + int32(1056)
	v405 = v404 + v391
	v406 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v405))))
	v409 = v404 + v392
	v410 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v409))))
	*(*uint8)(unsafe.Add(mBase, uint32(v405))) = uint8(v410)
	*(*uint8)(unsafe.Add(mBase, uint32(v409))) = uint8(v406)
	v414 = v391 + int32(1)
	v416 = v392 + int32(-1)
	if v414 < v416 {
		v391 = v414
		v392 = v416
		goto L72
	} else {
		goto L74
	}
L73:
	;
	goto L69
L74:
	;
	goto L73
L75:
	;
	v550 = v54 + int32(1)
	v553 = v69
	goto L23
L76:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v17)+1056)) = v440
	if v69 < int32(2) {
		goto L77
	} else {
		goto L78
	}
L77:
	;
	F_luaL_addlstring(m, v17+int32(20), v17+int32(1056), v69)
	mBase = m.M
	v501 = m.ExcPending
	if v501 != 0 {
		goto L1
	} else {
		goto L83
	}
L78:
	;
	v445 = *(*int32)(unsafe.Add(mBase, uint32(v17)+8))
	if v445 == int32(1) {
		goto L77
	} else {
		goto L79
	}
L79:
	;
	v453 = int32(0)
	v454 = v69 + int32(-1)
	goto L80
L80:
	;
	v466 = v17 + int32(1056)
	v467 = v466 + v453
	v468 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v467))))
	v471 = v466 + v454
	v472 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v471))))
	*(*uint8)(unsafe.Add(mBase, uint32(v467))) = uint8(v472)
	*(*uint8)(unsafe.Add(mBase, uint32(v471))) = uint8(v468)
	v476 = v453 + int32(1)
	v478 = v454 + int32(-1)
	if v476 < v478 {
		v453 = v476
		v454 = v478
		goto L80
	} else {
		goto L82
	}
L81:
	;
	goto L77
L82:
	;
	goto L81
L83:
	;
	v550 = v54 + int32(1)
	v553 = v69
	goto L23
L84:
	;
	v508 = *(*int32)(unsafe.Add(mBase, uint32(v17)+1056))
	if v69 != 0 {
		goto L86
	} else {
		goto L87
	}
L85:
	;
	F_luaL_addlstring(m, v17+int32(20), v506, v509)
	mBase = m.M
	v519 = m.ExcPending
	if v519 != 0 {
		goto L1
	} else {
		goto L91
	}
L86:
	;
	v509 = v69
	goto L88
L87:
	;
	v509 = v508
	goto L88
L88:
	;
	if base.Ui32(v509) <= base.Ui32(v508) {
		goto L85
	} else {
		goto L89
	}
L89:
	;
	v511 = m.G3
	v514 = F_luaL_argerror(m, l0, v503, v511+int32(_a2753))
	mBase = m.M
	v515 = m.ExcPending
	if v515 != 0 {
		goto L1
	} else {
		goto L90
	}
L90:
	;
	goto L85
L91:
	;
	if v72 != int32(115) {
		v538 = v509
		goto L92
	} else {
		goto L93
	}
L92:
	;
	v550 = v503
	v553 = v538
	goto L23
L93:
	;
	v522 = *(*int32)(unsafe.Add(mBase, uint32(v17)+20))
	if base.Ui32(v522) < base.Ui32(v46) {
		v529 = v522
		goto L94
	} else {
		goto L95
	}
L94:
	;
	v530 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v17)+20)) = v529 + v530
	v533 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v529))) = uint8(v533)
	v538 = v509 + v530
	goto L92
L95:
	;
	v526 = F_luaL_prepbuffer(m, v17+int32(20))
	mBase = m.M
	v527 = m.ExcPending
	if v527 != 0 {
		goto L1
	} else {
		goto L96
	}
L96:
	;
	v528 = *(*int32)(unsafe.Add(mBase, uint32(v17)+20))
	v529 = v528
	goto L94
L97:
	;
	v550 = v54
	v553 = v69
	goto L23
L98:
	;
	goto L8
L99:
	;
	m.G0 = v17 + int32(1088)
	return int32(1)
}
func F_backfillRdbReplicasToPsyncWait(m *base.Module) {
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
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v27 int64
	_ = v27
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v66 int32
	_ = v66
	var v69 int64
	_ = v69
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v87 int32
	_ = v87
	v6 = m.G0
	v8 = v6 - int32(320)
	m.G0 = v8
	v11 = *(*int32)(unsafe.Add(mBase, _consts[475]))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
	if v12 == int32(0) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	m.G0 = v8 + int32(320)
	return
L2:
	;
	v15 = *(*int32)(unsafe.Add(mBase, uint32(v12)+8))
	if v15 == int32(0) {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v19 = v8 + int32(16)
	v21 = *(*int32)(unsafe.Add(mBase, _consts[569]))
	*(*int32)(unsafe.Add(mBase, uint32(v19)+4)) = v21
	*(*int32)(unsafe.Add(mBase, uint32(v19))) = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v19)+20)) = int32(128)
	v27 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v19)+12)) = v27
	*(*int64)(unsafe.Add(mBase, uint32(v19)+296)) = v27
	*(*int64)(unsafe.Add(mBase, uint32(v19)+160)) = int64(137438953472)
	*(*int32)(unsafe.Add(mBase, uint32(v19)+8)) = v8 + int32(40)
	*(*int32)(unsafe.Add(mBase, uint32(v19)+156)) = v8 + int32(184)
	goto L4
L4:
	;
	v42 = int32(0)
	v44 = F_raxSeek(m, v8+int32(16), int32(_a67), v42, v42)
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	return
L6:
	;
	v48 = F_raxNext(m, v8+int32(16))
	mBase = m.M
	v49 = m.ExcPending
	if v49 != 0 {
		goto L5
	} else {
		goto L8
	}
L7:
	;
	F_raxStop(m, v8+int32(16))
	mBase = m.M
	v87 = m.ExcPending
	if v87 != 0 {
		goto L5
	} else {
		goto L18
	}
L8:
	;
	if v48 == int32(0) {
		goto L7
	} else {
		goto L9
	}
L9:
	;
	goto L10
L10:
	;
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v8)+28))
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v57)+104))
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v58)+184))
	if v59 != 0 {
		goto L12
	} else {
		goto L13
	}
L11:
	;
	goto L7
L12:
	;
	v77 = F_raxNext(m, v8+int32(16))
	mBase = m.M
	v78 = m.ExcPending
	if v78 != 0 {
		goto L5
	} else {
		goto L16
	}
L13:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v58)+184)) = v12
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v15)))
	*(*int32)(unsafe.Add(mBase, uint32(v15))) = v61 + int32(1)
	v66 = *(*int32)(unsafe.Add(mBase, _consts[28]))
	if int32(0) < v66 {
		goto L12
	} else {
		goto L14
	}
L14:
	;
	v69 = *(*int64)(unsafe.Add(mBase, uint32(v57)))
	*(*int64)(unsafe.Add(mBase, uint32(v8))) = v69
	F__serverLog(m, int32(0), int32(_a1918), v8)
	mBase = m.M
	v74 = m.ExcPending
	if v74 != 0 {
		goto L5
	} else {
		goto L15
	}
L15:
	;
	goto L12
L16:
	;
	if v77 != 0 {
		goto L10
	} else {
		goto L17
	}
L17:
	;
	goto L11
L18:
	;
	goto L1
}
func F_backgroundRewriteDoneHandler(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v19 int64
	_ = v19
	var v21 int32
	_ = v21
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
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
	var v47 int32
	_ = v47
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
	var v57 int64
	_ = v57
	var v63 int64
	_ = v63
	var v64 int64
	_ = v64
	var v67 int32
	_ = v67
	var v71 int32
	_ = v71
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v89 int32
	_ = v89
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	var v97 int32
	_ = v97
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v116 int64
	_ = v116
	var v121 int64
	_ = v121
	var v124 int64
	_ = v124
	var v126 int64
	_ = v126
	var v129 int64
	_ = v129
	var v135 int32
	_ = v135
	var v140 int32
	_ = v140
	var v152 int32
	_ = v152
	var v155 int32
	_ = v155
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v165 int32
	_ = v165
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
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
	var v185 int64
	_ = v185
	var v187 int64
	_ = v187
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v205 int64
	_ = v205
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v218 int64
	_ = v218
	var v224 int64
	_ = v224
	var v225 int64
	_ = v225
	var v226 int32
	_ = v226
	var v230 int32
	_ = v230
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v248 int32
	_ = v248
	var v249 int32
	_ = v249
	var v252 int32
	_ = v252
	var v256 int32
	_ = v256
	var v258 int32
	_ = v258
	var v260 int32
	_ = v260
	var v264 int32
	_ = v264
	var v265 int32
	_ = v265
	var v269 int32
	_ = v269
	var v271 int32
	_ = v271
	var v273 int32
	_ = v273
	var v275 int32
	_ = v275
	var v277 int32
	_ = v277
	var v278 int32
	_ = v278
	var v283 int64
	_ = v283
	var v288 int64
	_ = v288
	var v291 int64
	_ = v291
	var v293 int64
	_ = v293
	var v296 int64
	_ = v296
	var v302 int32
	_ = v302
	var v307 int32
	_ = v307
	var v317 int32
	_ = v317
	var v319 int32
	_ = v319
	var v321 int32
	_ = v321
	var v323 int32
	_ = v323
	var v331 int32
	_ = v331
	var v332 int32
	_ = v332
	var v335 int32
	_ = v335
	var v336 int32
	_ = v336
	var v337 int32
	_ = v337
	var v338 int32
	_ = v338
	var v340 int32
	_ = v340
	var v346 int32
	_ = v346
	var v350 int32
	_ = v350
	var v352 int32
	_ = v352
	var v354 int32
	_ = v354
	var v359 int64
	_ = v359
	var v360 int32
	_ = v360
	var v361 int32
	_ = v361
	var v362 int64
	_ = v362
	var v363 int64
	_ = v363
	var v368 int32
	_ = v368
	var v369 int32
	_ = v369
	var v370 int32
	_ = v370
	var v377 int32
	_ = v377
	var v384 int32
	_ = v384
	var v386 int32
	_ = v386
	var v389 int32
	_ = v389
	var v394 int64
	_ = v394
	var v397 int32
	_ = v397
	var v400 int64
	_ = v400
	var v408 int32
	_ = v408
	var v409 int32
	_ = v409
	var v410 int32
	_ = v410
	var v411 int32
	_ = v411
	var v414 int32
	_ = v414
	var v418 int32
	_ = v418
	var v420 int32
	_ = v420
	var v422 int32
	_ = v422
	var v426 int32
	_ = v426
	var v427 int32
	_ = v427
	var v431 int32
	_ = v431
	var v433 int32
	_ = v433
	var v435 int32
	_ = v435
	var v438 int32
	_ = v438
	var v439 int32
	_ = v439
	var v441 int32
	_ = v441
	var v442 int32
	_ = v442
	var v447 int64
	_ = v447
	var v453 int32
	_ = v453
	var v458 int64
	_ = v458
	var v463 int32
	_ = v463
	var v470 int32
	_ = v470
	var v473 int32
	_ = v473
	var v478 int64
	_ = v478
	var v483 int32
	_ = v483
	var v492 int32
	_ = v492
	var v506 int32
	_ = v506
	var v509 int32
	_ = v509
	var v511 int32
	_ = v511
	var v515 int32
	_ = v515
	var v517 int32
	_ = v517
	var v519 int32
	_ = v519
	var v520 int32
	_ = v520
	var v523 int32
	_ = v523
	var v525 int64
	_ = v525
	var v526 int32
	_ = v526
	var v527 int64
	_ = v527
	var v535 int32
	_ = v535
	var v548 int32
	_ = v548
	var v554 int32
	_ = v554
	v14 = m.G0
	v16 = v14 - int32(416)
	m.G0 = v16
	if l1|l0 != 0 {
		goto L4
	} else {
		goto L5
	}
L1:
	;
	F__serverAssert(m, int32(_a153), int32(_a123), int32(2797))
	mBase = m.M
	v554 = m.ExcPending
	if v554 != 0 {
		goto L8
	} else {
		goto L157
	}
L2:
	;
	F__serverAssert(m, int32(_a137), int32(_a123), int32(2789))
	mBase = m.M
	v548 = m.ExcPending
	if v548 != 0 {
		goto L8
	} else {
		goto L156
	}
L3:
	;
	v506 = *(*int32)(unsafe.Add(mBase, _consts[61]))
	F_aofRemoveTempFile(m, v506, int32(0))
	mBase = m.M
	v509 = m.ExcPending
	if v509 != 0 {
		goto L8
	} else {
		goto L148
	}
L4:
	;
	if l0 == int32(0) {
		goto L139
	} else {
		goto L140
	}
L5:
	;
	v19 = F_ustime(m)
	mBase = m.M
	v21 = *(*int32)(unsafe.Add(mBase, _consts[28]))
	if int32(2) < v21 {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	v30 = *(*int32)(unsafe.Add(mBase, _consts[61]))
	*(*int32)(unsafe.Add(mBase, uint32(v16)+128)) = v30
	v38 = F_snprintf(m, v16+int32(160), int32(256), int32(_a168), v16+int32(128))
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L8
	} else {
		goto L10
	}
L7:
	;
	F__serverLog(m, int32(2), int32(_a171), int32(0))
	mBase = m.M
	v28 = m.ExcPending
	if v28 != 0 {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	return
L9:
	;
	goto L6
L10:
	;
	v41 = *(*int32)(unsafe.Add(mBase, _consts[39]))
	if v41 == int32(0) {
		goto L2
	} else {
		goto L11
	}
L11:
	;
	v44 = F_aofManifestDup(m, v41)
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L8
	} else {
		goto L12
	}
L12:
	;
	v47 = *(*int32)(unsafe.Add(mBase, _consts[67]))
	v48 = F_getNewBaseFileNameAndMarkPreAsHistory(m, v44, v47)
	mBase = m.M
	v49 = m.ExcPending
	if v49 != 0 {
		goto L8
	} else {
		goto L13
	}
L13:
	;
	if v48 == int32(0) {
		goto L1
	} else {
		goto L14
	}
L14:
	;
	v53 = *(*int32)(unsafe.Add(mBase, _consts[37]))
	v54 = F_makePath(m, v53, v48)
	mBase = m.M
	v55 = m.ExcPending
	if v55 != 0 {
		goto L8
	} else {
		goto L15
	}
L15:
	;
	v57 = *(*int64)(unsafe.Add(mBase, _consts[45]))
	if base.B2i32(v57 == int64(0)) == int32(0) {
		goto L17
	} else {
		goto L18
	}
L16:
	;
	v67 = F_rename(m, v16+int32(160), v54)
	mBase = m.M
	if v67 != int32(-1) {
		goto L19
	} else {
		goto L20
	}
L17:
	;
	v63 = F_ustime(m)
	mBase = m.M
	v64 = v63
	goto L16
L18:
	;
	v64 = int64(0)
	goto L16
L19:
	;
	v121 = *(*int64)(unsafe.Add(mBase, _consts[45]))
	if v121 == int64(0) {
		goto L40
	} else {
		goto L41
	}
L20:
	;
	v71 = *(*int32)(unsafe.Add(mBase, _consts[28]))
	if int32(3) < v71 {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v44)))
	if v86 == int32(0) {
		goto L26
	} else {
		goto L27
	}
L22:
	;
	goto L23
L23:
	;
	v75 = *(*int32)(unsafe.Add(mBase, _consts[18]))
	v76 = F___strerror_l(m, v75, v75)
	mBase = m.M
	goto L24
L24:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+8)) = v76
	*(*int32)(unsafe.Add(mBase, uint32(v16)+4)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v16))) = v16 + int32(160)
	F__serverLog(m, int32(3), int32(_a172), v16)
	mBase = m.M
	v85 = m.ExcPending
	if v85 != 0 {
		goto L8
	} else {
		goto L25
	}
L25:
	;
	goto L21
L26:
	;
	v97 = *(*int32)(unsafe.Add(mBase, uint32(v44)+4))
	if v97 == int32(0) {
		goto L32
	} else {
		goto L33
	}
L27:
	;
	v89 = *(*int32)(unsafe.Add(mBase, uint32(v86)))
	if v89 == int32(0) {
		goto L28
	} else {
		goto L29
	}
L28:
	;
	F_valkey_free(m, v86)
	mBase = m.M
	v95 = m.ExcPending
	if v95 != 0 {
		goto L8
	} else {
		goto L31
	}
L29:
	;
	F_sdsfree(m, v89)
	mBase = m.M
	v93 = m.ExcPending
	if v93 != 0 {
		goto L8
	} else {
		goto L30
	}
L30:
	;
	goto L28
L31:
	;
	goto L26
L32:
	;
	v102 = *(*int32)(unsafe.Add(mBase, uint32(v44)+8))
	if v102 == int32(0) {
		goto L35
	} else {
		goto L36
	}
L33:
	;
	F_listRelease(m, v97)
	mBase = m.M
	v101 = m.ExcPending
	if v101 != 0 {
		goto L8
	} else {
		goto L34
	}
L34:
	;
	goto L32
L35:
	;
	F_valkey_free(m, v44)
	mBase = m.M
	v108 = m.ExcPending
	if v108 != 0 {
		goto L8
	} else {
		goto L38
	}
L36:
	;
	F_listRelease(m, v102)
	mBase = m.M
	v106 = m.ExcPending
	if v106 != 0 {
		goto L8
	} else {
		goto L37
	}
L37:
	;
	goto L35
L38:
	;
	F_sdsfree(m, v54)
	mBase = m.M
	v110 = m.ExcPending
	if v110 != 0 {
		goto L8
	} else {
		goto L39
	}
L39:
	;
	v111 = int32(_a20)
	*(*int32)(unsafe.Add(mBase, _consts[53])) = int32(-1)
	v116 = *(*int64)(unsafe.Add(mBase, _consts[52]))
	*(*int64)(unsafe.Add(mBase, _consts[52])) = v116 + int64(1)
	goto L3
L40:
	;
	v140 = *(*int32)(unsafe.Add(mBase, _consts[28]))
	if int32(2) < v140 {
		goto L45
	} else {
		goto L46
	}
L41:
	;
	v124 = F_ustime(m)
	mBase = m.M
	v126 = *(*int64)(unsafe.Add(mBase, _consts[45]))
	if v126 == int64(0) {
		goto L40
	} else {
		goto L42
	}
L42:
	;
	v129 = v124 - v64
	if v129 < v126*int64(1000) {
		goto L40
	} else {
		goto L43
	}
L43:
	;
	F_latencyAddSample(m, int32(_a173), v129)
	mBase = m.M
	v135 = m.ExcPending
	if v135 != 0 {
		goto L8
	} else {
		goto L44
	}
L44:
	;
	goto L40
L45:
	;
	v155 = *(*int32)(unsafe.Add(mBase, _consts[38]))
	if v155 != int32(2) {
		v323 = int32(0)
		goto L48
	} else {
		goto L49
	}
L46:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+116)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v16)+112)) = v16 + int32(160)
	F__serverLog(m, int32(2), int32(_a174), v16+int32(112))
	mBase = m.M
	v152 = m.ExcPending
	if v152 != 0 {
		goto L8
	} else {
		goto L47
	}
L47:
	;
	goto L45
L48:
	;
	F_markRewrittenIncrAofAsHistory(m, v44)
	mBase = m.M
	v331 = m.ExcPending
	if v331 != 0 {
		goto L8
	} else {
		goto L96
	}
L49:
	;
	v158 = F_sdsempty(m)
	mBase = m.M
	v159 = m.ExcPending
	if v159 != 0 {
		goto L8
	} else {
		goto L50
	}
L50:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+96)) = int32(_a133)
	*(*int32)(unsafe.Add(mBase, uint32(v16)+104)) = int32(_a134)
	v165 = *(*int32)(unsafe.Add(mBase, _consts[36]))
	*(*int32)(unsafe.Add(mBase, uint32(v16)+100)) = v165
	v170 = F_sdscatprintf(m, v158, int32(_a135), v16+int32(96))
	mBase = m.M
	v171 = m.ExcPending
	if v171 != 0 {
		goto L8
	} else {
		goto L51
	}
L51:
	;
	v173 = *(*int32)(unsafe.Add(mBase, _consts[37]))
	v174 = F_makePath(m, v173, v170)
	mBase = m.M
	v175 = m.ExcPending
	if v175 != 0 {
		goto L8
	} else {
		goto L52
	}
L52:
	;
	v177 = F_valkey_calloc(m, int32(24))
	mBase = m.M
	v178 = m.ExcPending
	if v178 != 0 {
		goto L8
	} else {
		goto L53
	}
L53:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v177)+16)) = int32(105)
	v181 = F_sdsempty(m)
	mBase = m.M
	v182 = m.ExcPending
	if v182 != 0 {
		goto L8
	} else {
		goto L54
	}
L54:
	;
	v184 = *(*int32)(unsafe.Add(mBase, _consts[36]))
	v185 = *(*int64)(unsafe.Add(mBase, uint32(v44)+24))
	v187 = v185 + int64(1)
	*(*int64)(unsafe.Add(mBase, uint32(v44)+24)) = v187
	*(*int32)(unsafe.Add(mBase, uint32(v16+int32(84)))) = int32(_a129)
	*(*int32)(unsafe.Add(mBase, uint32(v16+int32(80)))) = int32(_a134)
	*(*int64)(unsafe.Add(mBase, uint32(v16)+72)) = v187
	*(*int32)(unsafe.Add(mBase, uint32(v16)+64)) = v184
	v202 = F_sdscatprintf(m, v181, int32(_a131), v16+int32(64))
	mBase = m.M
	v203 = m.ExcPending
	if v203 != 0 {
		goto L8
	} else {
		goto L55
	}
L55:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v177))) = v202
	v205 = *(*int64)(unsafe.Add(mBase, uint32(v44)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v177)+8)) = v205
	v207 = *(*int32)(unsafe.Add(mBase, uint32(v44)+4))
	v208 = F_listAddNodeTail(m, v207, v177)
	mBase = m.M
	v209 = m.ExcPending
	if v209 != 0 {
		goto L8
	} else {
		goto L56
	}
L56:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+32)) = int32(1)
	v213 = *(*int32)(unsafe.Add(mBase, _consts[37]))
	v214 = *(*int32)(unsafe.Add(mBase, uint32(v177)))
	v215 = F_makePath(m, v213, v214)
	mBase = m.M
	v216 = m.ExcPending
	if v216 != 0 {
		goto L8
	} else {
		goto L57
	}
L57:
	;
	v218 = *(*int64)(unsafe.Add(mBase, _consts[45]))
	if base.B2i32(v218 == int64(0)) == int32(0) {
		goto L59
	} else {
		goto L60
	}
L58:
	;
	v226 = F_rename(m, v174, v215)
	mBase = m.M
	if v226 != int32(-1) {
		goto L61
	} else {
		goto L62
	}
L59:
	;
	v224 = F_ustime(m)
	mBase = m.M
	v225 = v224
	goto L58
L60:
	;
	v225 = int64(0)
	goto L58
L61:
	;
	v288 = *(*int64)(unsafe.Add(mBase, _consts[45]))
	if v288 == int64(0) {
		goto L86
	} else {
		goto L87
	}
L62:
	;
	v230 = *(*int32)(unsafe.Add(mBase, _consts[28]))
	if int32(3) < v230 {
		goto L63
	} else {
		goto L64
	}
L63:
	;
	v245 = F_bg_unlink(m, v54)
	mBase = m.M
	v246 = m.ExcPending
	if v246 != 0 {
		goto L8
	} else {
		goto L68
	}
L64:
	;
	goto L65
L65:
	;
	v234 = *(*int32)(unsafe.Add(mBase, _consts[18]))
	v235 = F___strerror_l(m, v234, v234)
	mBase = m.M
	goto L66
L66:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+24)) = v235
	*(*int32)(unsafe.Add(mBase, uint32(v16)+20)) = v215
	*(*int32)(unsafe.Add(mBase, uint32(v16)+16)) = v174
	F__serverLog(m, int32(3), int32(_a175), v16+int32(16))
	mBase = m.M
	v244 = m.ExcPending
	if v244 != 0 {
		goto L8
	} else {
		goto L67
	}
L67:
	;
	goto L63
L68:
	;
	F_sdsfree(m, v54)
	mBase = m.M
	v248 = m.ExcPending
	if v248 != 0 {
		goto L8
	} else {
		goto L69
	}
L69:
	;
	v249 = *(*int32)(unsafe.Add(mBase, uint32(v44)))
	if v249 == int32(0) {
		goto L70
	} else {
		goto L71
	}
L70:
	;
	v260 = *(*int32)(unsafe.Add(mBase, uint32(v44)+4))
	if v260 == int32(0) {
		goto L76
	} else {
		goto L77
	}
L71:
	;
	v252 = *(*int32)(unsafe.Add(mBase, uint32(v249)))
	if v252 == int32(0) {
		goto L72
	} else {
		goto L73
	}
L72:
	;
	F_valkey_free(m, v249)
	mBase = m.M
	v258 = m.ExcPending
	if v258 != 0 {
		goto L8
	} else {
		goto L75
	}
L73:
	;
	F_sdsfree(m, v252)
	mBase = m.M
	v256 = m.ExcPending
	if v256 != 0 {
		goto L8
	} else {
		goto L74
	}
L74:
	;
	goto L72
L75:
	;
	goto L70
L76:
	;
	v265 = *(*int32)(unsafe.Add(mBase, uint32(v44)+8))
	if v265 == int32(0) {
		goto L79
	} else {
		goto L80
	}
L77:
	;
	F_listRelease(m, v260)
	mBase = m.M
	v264 = m.ExcPending
	if v264 != 0 {
		goto L8
	} else {
		goto L78
	}
L78:
	;
	goto L76
L79:
	;
	F_valkey_free(m, v44)
	mBase = m.M
	v271 = m.ExcPending
	if v271 != 0 {
		goto L8
	} else {
		goto L82
	}
L80:
	;
	F_listRelease(m, v265)
	mBase = m.M
	v269 = m.ExcPending
	if v269 != 0 {
		goto L8
	} else {
		goto L81
	}
L81:
	;
	goto L79
L82:
	;
	F_sdsfree(m, v174)
	mBase = m.M
	v273 = m.ExcPending
	if v273 != 0 {
		goto L8
	} else {
		goto L83
	}
L83:
	;
	F_sdsfree(m, v215)
	mBase = m.M
	v275 = m.ExcPending
	if v275 != 0 {
		goto L8
	} else {
		goto L84
	}
L84:
	;
	F_sdsfree(m, v170)
	mBase = m.M
	v277 = m.ExcPending
	if v277 != 0 {
		goto L8
	} else {
		goto L85
	}
L85:
	;
	v278 = int32(_a20)
	*(*int32)(unsafe.Add(mBase, _consts[53])) = int32(-1)
	v283 = *(*int64)(unsafe.Add(mBase, _consts[52]))
	*(*int64)(unsafe.Add(mBase, _consts[52])) = v283 + int64(1)
	goto L3
L86:
	;
	v307 = *(*int32)(unsafe.Add(mBase, _consts[28]))
	if int32(2) < v307 {
		goto L91
	} else {
		goto L92
	}
L87:
	;
	v291 = F_ustime(m)
	mBase = m.M
	v293 = *(*int64)(unsafe.Add(mBase, _consts[45]))
	if v293 == int64(0) {
		goto L86
	} else {
		goto L88
	}
L88:
	;
	v296 = v291 - v225
	if v296 < v293*int64(1000) {
		goto L86
	} else {
		goto L89
	}
L89:
	;
	F_latencyAddSample(m, int32(_a173), v296)
	mBase = m.M
	v302 = m.ExcPending
	if v302 != 0 {
		goto L8
	} else {
		goto L90
	}
L90:
	;
	goto L86
L91:
	;
	F_sdsfree(m, v174)
	mBase = m.M
	v319 = m.ExcPending
	if v319 != 0 {
		goto L8
	} else {
		goto L94
	}
L92:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+52)) = v214
	*(*int32)(unsafe.Add(mBase, uint32(v16)+48)) = v170
	F__serverLog(m, int32(2), int32(_a176), v16+int32(48))
	mBase = m.M
	v317 = m.ExcPending
	if v317 != 0 {
		goto L8
	} else {
		goto L93
	}
L93:
	;
	goto L91
L94:
	;
	F_sdsfree(m, v170)
	mBase = m.M
	v321 = m.ExcPending
	if v321 != 0 {
		goto L8
	} else {
		goto L95
	}
L95:
	;
	v323 = v215
	goto L48
L96:
	;
	v332 = *(*int32)(unsafe.Add(mBase, uint32(v44)+32))
	if v332 == int32(0) {
		goto L98
	} else {
		goto L99
	}
L97:
	;
	v409 = F_bg_unlink(m, v54)
	mBase = m.M
	v410 = m.ExcPending
	if v410 != 0 {
		goto L8
	} else {
		goto L120
	}
L98:
	;
	F_sdsfree(m, v54)
	mBase = m.M
	v346 = m.ExcPending
	if v346 != 0 {
		goto L8
	} else {
		goto L104
	}
L99:
	;
	v335 = F_getAofManifestAsString(m, v44)
	mBase = m.M
	v336 = m.ExcPending
	if v336 != 0 {
		goto L8
	} else {
		goto L100
	}
L100:
	;
	v337 = F_writeAofManifestFile(m, v335)
	mBase = m.M
	v338 = m.ExcPending
	if v338 != 0 {
		goto L8
	} else {
		goto L101
	}
L101:
	;
	F_sdsfree(m, v335)
	mBase = m.M
	v340 = m.ExcPending
	if v340 != 0 {
		goto L8
	} else {
		goto L102
	}
L102:
	;
	if v337 != 0 {
		goto L97
	} else {
		goto L103
	}
L103:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+32)) = int32(0)
	goto L98
L104:
	;
	if v323 == int32(0) {
		goto L105
	} else {
		goto L106
	}
L105:
	;
	F_aofManifestFreeAndUpdate(m, v44)
	mBase = m.M
	v352 = m.ExcPending
	if v352 != 0 {
		goto L8
	} else {
		goto L108
	}
L106:
	;
	F_sdsfree(m, v323)
	mBase = m.M
	v350 = m.ExcPending
	if v350 != 0 {
		goto L8
	} else {
		goto L107
	}
L107:
	;
	goto L105
L108:
	;
	v354 = *(*int32)(unsafe.Add(mBase, _consts[38]))
	if v354 == int32(0) {
		goto L109
	} else {
		goto L110
	}
L109:
	;
	v368 = F_aofDelHistoryFiles(m)
	mBase = m.M
	v369 = m.ExcPending
	if v369 != 0 {
		goto L8
	} else {
		goto L112
	}
L110:
	;
	v359 = F_getAppendOnlyFileSize(m, v48, int32(0))
	mBase = m.M
	v360 = m.ExcPending
	if v360 != 0 {
		goto L8
	} else {
		goto L111
	}
L111:
	;
	v361 = int32(_a20)
	v362 = *(*int64)(unsafe.Add(mBase, _consts[42]))
	v363 = v359 + v362
	*(*int64)(unsafe.Add(mBase, _consts[56])) = v363
	*(*int64)(unsafe.Add(mBase, _consts[57])) = v363
	goto L109
L112:
	;
	v370 = int32(_a20)
	*(*int64)(unsafe.Add(mBase, _consts[52])) = int64(0)
	*(*int32)(unsafe.Add(mBase, _consts[53])) = int32(0)
	v377 = *(*int32)(unsafe.Add(mBase, _consts[28]))
	if int32(2) < v377 {
		goto L113
	} else {
		goto L114
	}
L113:
	;
	v386 = *(*int32)(unsafe.Add(mBase, _consts[38]))
	if v386 != int32(2) {
		goto L116
	} else {
		goto L117
	}
L114:
	;
	F__serverLog(m, int32(2), int32(_a170), int32(0))
	mBase = m.M
	v384 = m.ExcPending
	if v384 != 0 {
		goto L8
	} else {
		goto L115
	}
L115:
	;
	goto L113
L116:
	;
	v397 = *(*int32)(unsafe.Add(mBase, _consts[28]))
	if int32(1) < v397 {
		goto L3
	} else {
		goto L118
	}
L117:
	;
	v389 = int32(_a20)
	*(*int32)(unsafe.Add(mBase, _consts[38])) = int32(1)
	v394 = *(*int64)(unsafe.Add(mBase, _consts[48]))
	*(*int64)(unsafe.Add(mBase, _consts[49])) = v394
	goto L116
L118:
	;
	v400 = F_ustime(m)
	mBase = m.M
	*(*int64)(unsafe.Add(mBase, uint32(v16)+32)) = v400 - v19
	F__serverLog(m, int32(1), int32(_a169), v16+int32(32))
	mBase = m.M
	v408 = m.ExcPending
	if v408 != 0 {
		goto L8
	} else {
		goto L119
	}
L119:
	;
	goto L3
L120:
	;
	v411 = *(*int32)(unsafe.Add(mBase, uint32(v44)))
	if v411 == int32(0) {
		goto L121
	} else {
		goto L122
	}
L121:
	;
	v422 = *(*int32)(unsafe.Add(mBase, uint32(v44)+4))
	if v422 == int32(0) {
		goto L127
	} else {
		goto L128
	}
L122:
	;
	v414 = *(*int32)(unsafe.Add(mBase, uint32(v411)))
	if v414 == int32(0) {
		goto L123
	} else {
		goto L124
	}
L123:
	;
	F_valkey_free(m, v411)
	mBase = m.M
	v420 = m.ExcPending
	if v420 != 0 {
		goto L8
	} else {
		goto L126
	}
L124:
	;
	F_sdsfree(m, v414)
	mBase = m.M
	v418 = m.ExcPending
	if v418 != 0 {
		goto L8
	} else {
		goto L125
	}
L125:
	;
	goto L123
L126:
	;
	goto L121
L127:
	;
	v427 = *(*int32)(unsafe.Add(mBase, uint32(v44)+8))
	if v427 == int32(0) {
		goto L130
	} else {
		goto L131
	}
L128:
	;
	F_listRelease(m, v422)
	mBase = m.M
	v426 = m.ExcPending
	if v426 != 0 {
		goto L8
	} else {
		goto L129
	}
L129:
	;
	goto L127
L130:
	;
	F_valkey_free(m, v44)
	mBase = m.M
	v433 = m.ExcPending
	if v433 != 0 {
		goto L8
	} else {
		goto L133
	}
L131:
	;
	F_listRelease(m, v427)
	mBase = m.M
	v431 = m.ExcPending
	if v431 != 0 {
		goto L8
	} else {
		goto L132
	}
L132:
	;
	goto L130
L133:
	;
	F_sdsfree(m, v54)
	mBase = m.M
	v435 = m.ExcPending
	if v435 != 0 {
		goto L8
	} else {
		goto L134
	}
L134:
	;
	if v323 == int32(0) {
		goto L135
	} else {
		goto L136
	}
L135:
	;
	v442 = int32(_a20)
	*(*int32)(unsafe.Add(mBase, _consts[53])) = int32(-1)
	v447 = *(*int64)(unsafe.Add(mBase, _consts[52]))
	*(*int64)(unsafe.Add(mBase, _consts[52])) = v447 + int64(1)
	goto L3
L136:
	;
	v438 = F_bg_unlink(m, v323)
	mBase = m.M
	v439 = m.ExcPending
	if v439 != 0 {
		goto L8
	} else {
		goto L137
	}
L137:
	;
	F_sdsfree(m, v323)
	mBase = m.M
	v441 = m.ExcPending
	if v441 != 0 {
		goto L8
	} else {
		goto L138
	}
L138:
	;
	goto L135
L139:
	;
	if l1 == int32(10) {
		goto L144
	} else {
		goto L145
	}
L140:
	;
	if l1 != 0 {
		goto L139
	} else {
		goto L141
	}
L141:
	;
	v453 = int32(_a20)
	*(*int32)(unsafe.Add(mBase, _consts[53])) = int32(-1)
	v458 = *(*int64)(unsafe.Add(mBase, _consts[52]))
	*(*int64)(unsafe.Add(mBase, _consts[52])) = v458 + int64(1)
	v463 = *(*int32)(unsafe.Add(mBase, _consts[28]))
	if int32(3) < v463 {
		goto L3
	} else {
		goto L142
	}
L142:
	;
	F__serverLog(m, int32(3), int32(_a167), int32(0))
	mBase = m.M
	v470 = m.ExcPending
	if v470 != 0 {
		goto L8
	} else {
		goto L143
	}
L143:
	;
	goto L3
L144:
	;
	v483 = *(*int32)(unsafe.Add(mBase, _consts[28]))
	if int32(3) < v483 {
		goto L3
	} else {
		goto L146
	}
L145:
	;
	v473 = int32(_a20)
	*(*int32)(unsafe.Add(mBase, _consts[53])) = int32(-1)
	v478 = *(*int64)(unsafe.Add(mBase, _consts[52]))
	*(*int64)(unsafe.Add(mBase, _consts[52])) = v478 + int64(1)
	goto L144
L146:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+144)) = l1
	F__serverLog(m, int32(3), int32(_a166), v16+int32(144))
	mBase = m.M
	v492 = m.ExcPending
	if v492 != 0 {
		goto L8
	} else {
		goto L147
	}
L147:
	;
	goto L3
L148:
	;
	v511 = *(*int32)(unsafe.Add(mBase, _consts[38]))
	if v511 != int32(2) {
		goto L149
	} else {
		goto L150
	}
L149:
	;
	v525 = F___time(m, int32(0))
	mBase = m.M
	v526 = int32(_a20)
	v527 = *(*int64)(unsafe.Add(mBase, _consts[64]))
	*(*int64)(unsafe.Add(mBase, _consts[64])) = int64(-1)
	*(*int64)(unsafe.Add(mBase, _consts[65])) = v525 - v527
	v535 = *(*int32)(unsafe.Add(mBase, _consts[38]))
	if v535 != int32(2) {
		goto L154
	} else {
		goto L155
	}
L150:
	;
	v515 = *(*int32)(unsafe.Add(mBase, _consts[66]))
	F_sdsfree(m, v515)
	mBase = m.M
	v517 = m.ExcPending
	if v517 != 0 {
		goto L8
	} else {
		goto L151
	}
L151:
	;
	v519 = F_sdsempty(m)
	mBase = m.M
	v520 = m.ExcPending
	if v520 != 0 {
		goto L8
	} else {
		goto L152
	}
L152:
	;
	*(*int32)(unsafe.Add(mBase, _consts[66])) = v519
	F_aofDelTempIncrAofFile(m)
	mBase = m.M
	v523 = m.ExcPending
	if v523 != 0 {
		goto L8
	} else {
		goto L153
	}
L153:
	;
	goto L149
L154:
	;
	m.G0 = v16 + int32(416)
	return
L155:
	;
	*(*int32)(unsafe.Add(mBase, _consts[63])) = int32(1)
	goto L154
L156:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L157:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_backgroundSaveDoneHandler(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v18 int64
	_ = v18
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v37 int64
	_ = v37
	var v39 int64
	_ = v39
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v54 int32
	_ = v54
	var v61 int32
	_ = v61
	var v63 int64
	_ = v63
	var v69 int64
	_ = v69
	var v70 int64
	_ = v70
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v75 int64
	_ = v75
	var v84 int32
	_ = v84
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v95 int64
	_ = v95
	var v97 int32
	_ = v97
	var v125 int32
	_ = v125
	var v127 int32
	_ = v127
	var v129 int32
	_ = v129
	var v131 int32
	_ = v131
	var v134 int32
	_ = v134
	var v136 int32
	_ = v136
	var v140 int32
	_ = v140
	var v150 int32
	_ = v150
	var v152 int32
	_ = v152
	var v164 int32
	_ = v164
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v178 int32
	_ = v178
	var v180 int32
	_ = v180
	var v182 int32
	_ = v182
	var v188 int32
	_ = v188
	var v192 int32
	_ = v192
	var v196 int32
	_ = v196
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v204 int32
	_ = v204
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v209 int32
	_ = v209
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v225 int32
	_ = v225
	var v229 int32
	_ = v229
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v255 int32
	_ = v255
	var v257 int32
	_ = v257
	var v259 int32
	_ = v259
	var v265 int32
	_ = v265
	var v269 int32
	_ = v269
	var v273 int32
	_ = v273
	var v275 int32
	_ = v275
	var v276 int32
	_ = v276
	var v281 int32
	_ = v281
	var v283 int32
	_ = v283
	var v284 int32
	_ = v284
	var v286 int32
	_ = v286
	var v294 int32
	_ = v294
	var v295 int32
	_ = v295
	var v296 int32
	_ = v296
	var v302 int32
	_ = v302
	var v306 int32
	_ = v306
	var v320 int32
	_ = v320
	var v321 int32
	_ = v321
	var v323 int64
	_ = v323
	var v326 int64
	_ = v326
	var v328 int64
	_ = v328
	var v331 int64
	_ = v331
	var v337 int32
	_ = v337
	var v345 int32
	_ = v345
	var v352 int32
	_ = v352
	var v354 int32
	_ = v354
	var v361 int32
	_ = v361
	var v363 int32
	_ = v363
	var v372 int32
	_ = v372
	var v381 int32
	_ = v381
	var v384 int32
	_ = v384
	var v387 int32
	_ = v387
	var v389 int32
	_ = v389
	var v393 int32
	_ = v393
	var v396 int32
	_ = v396
	var v398 int32
	_ = v398
	var v399 int32
	_ = v399
	var v400 int32
	_ = v400
	var v404 int32
	_ = v404
	var v406 int32
	_ = v406
	var v407 int32
	_ = v407
	var v414 int32
	_ = v414
	var v416 int32
	_ = v416
	var v423 int32
	_ = v423
	var v427 int32
	_ = v427
	var v428 int32
	_ = v428
	var v431 int64
	_ = v431
	var v440 int32
	_ = v440
	var v442 int32
	_ = v442
	var v451 int32
	_ = v451
	v11 = m.G0
	v13 = v11 - int32(320)
	m.G0 = v13
	v15 = int32(_a20)
	v16 = *(*int32)(unsafe.Add(mBase, _consts[537]))
	v18 = F___time(m, int32(0))
	mBase = m.M
	v20 = *(*int32)(unsafe.Add(mBase, _consts[537]))
	switch v20 + int32(-1) {
	case 0:
		goto L4
	case 1:
		goto L3
	default:
		goto L1
	}
L1:
	;
	F__serverPanic_1(m, int32(_a1848), int32(3738), int32(_a1901), int32(0))
	mBase = m.M
	v451 = m.ExcPending
	if v451 != 0 {
		goto L10
	} else {
		goto L110
	}
L2:
	;
	v427 = int32(_a20)
	v428 = int32(0)
	*(*int32)(unsafe.Add(mBase, _consts[537])) = v428
	v431 = *(*int64)(unsafe.Add(mBase, _consts[538]))
	*(*int64)(unsafe.Add(mBase, _consts[538])) = int64(-1)
	*(*int64)(unsafe.Add(mBase, _consts[563])) = v18 - v431
	if v423 != 0 {
		goto L106
	} else {
		goto L107
	}
L3:
	;
	v352 = l1 | l0
	if v352 != 0 {
		goto L88
	} else {
		goto L89
	}
L4:
	;
	v23 = l1 | l0
	if v23 != 0 {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	*(*int32)(unsafe.Add(mBase, _consts[431])) = v345
	v423 = v23
	goto L2
L6:
	;
	v44 = *(*int32)(unsafe.Add(mBase, _consts[28]))
	if l0 == int32(0) {
		goto L12
	} else {
		goto L13
	}
L7:
	;
	v25 = *(*int32)(unsafe.Add(mBase, _consts[28]))
	if int32(2) < v25 {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v33 = int32(_a20)
	*(*int64)(unsafe.Add(mBase, _consts[295])) = v18
	v37 = *(*int64)(unsafe.Add(mBase, _consts[180]))
	v39 = *(*int64)(unsafe.Add(mBase, _consts[528]))
	*(*int64)(unsafe.Add(mBase, _consts[180])) = v37 - v39
	v345 = int32(0)
	goto L5
L9:
	;
	F__serverLog(m, int32(2), int32(_a1902), int32(0))
	mBase = m.M
	v32 = m.ExcPending
	if v32 != 0 {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	return
L11:
	;
	goto L8
L12:
	;
	if int32(3) < v44 {
		goto L17
	} else {
		goto L18
	}
L13:
	;
	if l1 != 0 {
		goto L12
	} else {
		goto L14
	}
L14:
	;
	v47 = int32(-1)
	if int32(3) < v44 {
		v345 = v47
		goto L5
	} else {
		goto L15
	}
L15:
	;
	F__serverLog(m, int32(3), int32(_a1903), int32(0))
	mBase = m.M
	v54 = m.ExcPending
	if v54 != 0 {
		goto L10
	} else {
		goto L16
	}
L16:
	;
	v345 = v47
	goto L5
L17:
	;
	v63 = *(*int64)(unsafe.Add(mBase, _consts[45]))
	if base.B2i32(v63 == int64(0)) == int32(0) {
		goto L21
	} else {
		goto L22
	}
L18:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13))) = l1
	F__serverLog(m, int32(3), int32(_a1904), v13)
	mBase = m.M
	v61 = m.ExcPending
	if v61 != 0 {
		goto L10
	} else {
		goto L19
	}
L19:
	;
	goto L17
L20:
	;
	v71 = int32(32)
	v72 = v13 + v71
	v75 = int64(*(*int32)(unsafe.Add(mBase, _consts[61])))
	if v75 <= int64(-1) {
		goto L27
	} else {
		goto L28
	}
L21:
	;
	v69 = F_ustime(m)
	mBase = m.M
	v70 = v69
	goto L20
L22:
	;
	v70 = int64(0)
	goto L20
L23:
	;
	goto L35
L24:
	;
	goto L23
L26:
	;
	v97 = F_ull2string(m, v93, v94, v95)
	mBase = m.M
	if v97 == int32(0) {
		goto L24
	} else {
		goto L30
	}
L27:
	;
	goto L29
L28:
	;
	v93 = v72
	v94 = v71
	v95 = v75
	goto L26
L29:
	;
	v84 = int32(45)
	*(*uint8)(unsafe.Add(mBase, uint32(v72))) = uint8(v84)
	v93 = v13 + int32(33)
	v94 = int32(31)
	v95 = int64(0) - v75
	goto L26
L30:
	;
	goto L23
L32:
	;
	v164 = v13 + int32(64)
	v166 = v13 + int32(32)
	v167 = int32(256)
	goto L46
L33:
	;
	goto L32
L34:
	;
	v150 = v129
	goto L41
L35:
	;
	v125 = v13 + int32(64)
	v127 = int32(256)
	v129 = int32(_a133)
	goto L37
L36:
	;
	v140 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v125))) = uint8(v140)
	goto L34
L37:
	;
	v131 = v127 + int32(-1)
	if v131 == int32(0) {
		goto L36
	} else {
		goto L39
	}
L39:
	;
	v134 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v129))))
	*(*uint8)(unsafe.Add(mBase, uint32(v125))) = uint8(v134)
	v136 = int32(1)
	if v134 != 0 {
		v125 = v125 + v136
		v127 = v131
		v129 = v129 + v136
		goto L37
	} else {
		goto L40
	}
L40:
	;
	goto L33
L41:
	;
	v152 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v150))))
	if v152 != 0 {
		v150 = v150 + int32(1)
		goto L41
	} else {
		goto L43
	}
L42:
	;
	goto L33
L43:
	;
	goto L42
L44:
	;
	v242 = v13 + int32(64)
	v243 = int32(_a128)
	v244 = int32(256)
	goto L64
L45:
	;
	v196 = v192 - v164
	if v167 == v196 {
		goto L53
	} else {
		goto L54
	}
L46:
	;
	v178 = v164
	v180 = v167
	goto L47
L47:
	;
	v182 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v178))))
	if v182 == int32(0) {
		v192 = v178
		goto L45
	} else {
		goto L49
	}
L48:
	;
	v192 = v13 + int32(320)
	goto L45
L49:
	;
	v188 = v180 + int32(-1)
	if v188 != 0 {
		v178 = v178 + int32(1)
		v180 = v188
		goto L47
	} else {
		goto L50
	}
L50:
	;
	goto L48
L51:
	;
	v229 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v225))) = uint8(v229)
	goto L44
L52:
	;
	v204 = v198
	v206 = v167 + (v196 ^ int32(-1))
	v207 = v192
	v209 = v166
	goto L56
L53:
	;
	v199 = F_strlen(m, v166)
	mBase = m.M
	goto L44
L54:
	;
	v198 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v166))))
	if v198 != 0 {
		goto L52
	} else {
		goto L55
	}
L55:
	;
	v225 = v192
	goto L51
L56:
	;
	if v206 != 0 {
		goto L59
	} else {
		goto L60
	}
L57:
	;
	v225 = v218
	goto L51
L58:
	;
	v219 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v209)+1)))
	if v219 != 0 {
		v204 = v219
		v206 = v217
		v207 = v218
		v209 = v209 + int32(1)
		goto L56
	} else {
		goto L61
	}
L59:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v207))) = uint8(v204)
	v217 = v206 + int32(-1)
	v218 = v207 + int32(1)
	goto L58
L60:
	;
	v217 = int32(0)
	v218 = v207
	goto L58
L61:
	;
	goto L57
L62:
	;
	v320 = F_bg_unlink(m, v13+int32(64))
	mBase = m.M
	v321 = m.ExcPending
	if v321 != 0 {
		goto L10
	} else {
		goto L80
	}
L63:
	;
	v273 = v269 - v242
	if v244 == v273 {
		goto L71
	} else {
		goto L72
	}
L64:
	;
	v255 = v242
	v257 = v244
	goto L65
L65:
	;
	v259 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v255))))
	if v259 == int32(0) {
		v269 = v255
		goto L63
	} else {
		goto L67
	}
L66:
	;
	v269 = v13 + int32(320)
	goto L63
L67:
	;
	v265 = v257 + int32(-1)
	if v265 != 0 {
		v255 = v255 + int32(1)
		v257 = v265
		goto L65
	} else {
		goto L68
	}
L68:
	;
	goto L66
L69:
	;
	v306 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v302))) = uint8(v306)
	goto L62
L70:
	;
	v281 = v275
	v283 = v244 + (v273 ^ int32(-1))
	v284 = v269
	v286 = v243
	goto L74
L71:
	;
	v276 = F_strlen(m, v243)
	mBase = m.M
	goto L62
L72:
	;
	v275 = int32(*(*uint8)(unsafe.Add(mBase, _consts[539])))
	if v275 != 0 {
		goto L70
	} else {
		goto L73
	}
L73:
	;
	v302 = v269
	goto L69
L74:
	;
	if v283 != 0 {
		goto L77
	} else {
		goto L78
	}
L75:
	;
	v302 = v295
	goto L69
L76:
	;
	v296 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v286)+1)))
	if v296 != 0 {
		v281 = v296
		v283 = v294
		v284 = v295
		v286 = v286 + int32(1)
		goto L74
	} else {
		goto L79
	}
L77:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v284))) = uint8(v281)
	v294 = v283 + int32(-1)
	v295 = v284 + int32(1)
	goto L76
L78:
	;
	v294 = int32(0)
	v295 = v284
	goto L76
L79:
	;
	goto L75
L80:
	;
	v323 = *(*int64)(unsafe.Add(mBase, _consts[45]))
	if v323 == int64(0) {
		goto L81
	} else {
		goto L82
	}
L81:
	;
	if l1 == int32(10) {
		v423 = int32(1)
		goto L2
	} else {
		goto L86
	}
L82:
	;
	v326 = F_ustime(m)
	mBase = m.M
	v328 = *(*int64)(unsafe.Add(mBase, _consts[45]))
	if v328 == int64(0) {
		goto L81
	} else {
		goto L83
	}
L83:
	;
	v331 = v326 - v70
	if v331 < v328*int64(1000) {
		goto L81
	} else {
		goto L84
	}
L84:
	;
	F_latencyAddSample(m, int32(_a699), v331)
	mBase = m.M
	v337 = m.ExcPending
	if v337 != 0 {
		goto L10
	} else {
		goto L85
	}
L85:
	;
	goto L81
L86:
	;
	v345 = int32(-1)
	goto L5
L87:
	;
	v384 = *(*int32)(unsafe.Add(mBase, _consts[564]))
	if v384 == int32(-1) {
		goto L99
	} else {
		goto L100
	}
L88:
	;
	v363 = *(*int32)(unsafe.Add(mBase, _consts[28]))
	if l0 == int32(0) {
		goto L92
	} else {
		goto L93
	}
L89:
	;
	v354 = *(*int32)(unsafe.Add(mBase, _consts[28]))
	if int32(2) < v354 {
		goto L87
	} else {
		goto L90
	}
L90:
	;
	F__serverLog(m, int32(2), int32(_a1905), int32(0))
	mBase = m.M
	v361 = m.ExcPending
	if v361 != 0 {
		goto L10
	} else {
		goto L91
	}
L91:
	;
	goto L87
L92:
	;
	if int32(3) < v363 {
		goto L87
	} else {
		goto L97
	}
L93:
	;
	if l1 != 0 {
		goto L92
	} else {
		goto L94
	}
L94:
	;
	if int32(3) < v363 {
		goto L87
	} else {
		goto L95
	}
L95:
	;
	F__serverLog(m, int32(3), int32(_a1906), int32(0))
	mBase = m.M
	v372 = m.ExcPending
	if v372 != 0 {
		goto L10
	} else {
		goto L96
	}
L96:
	;
	goto L87
L97:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+16)) = l1
	F__serverLog(m, int32(3), int32(_a1907), v13+int32(16))
	mBase = m.M
	v381 = m.ExcPending
	if v381 != 0 {
		goto L10
	} else {
		goto L98
	}
L98:
	;
	goto L87
L99:
	;
	v389 = *(*int32)(unsafe.Add(mBase, _consts[565]))
	if v389 < int32(1) {
		goto L101
	} else {
		goto L102
	}
L100:
	;
	v387 = F_close(m, v384)
	mBase = m.M
	goto L99
L101:
	;
	v400 = int32(_a20)
	*(*int64)(unsafe.Add(mBase, _consts[565])) = int64(-1)
	v404 = *(*int32)(unsafe.Add(mBase, _consts[476]))
	F_valkey_free(m, v404)
	mBase = m.M
	v406 = m.ExcPending
	if v406 != 0 {
		goto L10
	} else {
		goto L104
	}
L102:
	;
	v393 = *(*int32)(unsafe.Add(mBase, _consts[223]))
	F_aeDeleteFileEvent(m, v393, v389, int32(1))
	mBase = m.M
	v396 = m.ExcPending
	if v396 != 0 {
		goto L10
	} else {
		goto L103
	}
L103:
	;
	v398 = *(*int32)(unsafe.Add(mBase, _consts[565]))
	v399 = F_close(m, v398)
	mBase = m.M
	goto L101
L104:
	;
	v407 = int32(_a20)
	*(*int32)(unsafe.Add(mBase, _consts[566])) = int32(0)
	*(*int64)(unsafe.Add(mBase, _consts[476])) = int64(0)
	v414 = *(*int32)(unsafe.Add(mBase, _consts[567]))
	F_valkey_free(m, v414)
	mBase = m.M
	v416 = m.ExcPending
	if v416 != 0 {
		goto L10
	} else {
		goto L105
	}
L105:
	;
	*(*int64)(unsafe.Add(mBase, _consts[567])) = int64(0)
	v423 = v352
	goto L2
L106:
	;
	v440 = int32(-1)
	goto L108
L107:
	;
	v440 = v428
	goto L108
L108:
	;
	F_updateReplicasWaitingBgsave(m, v440, v16)
	mBase = m.M
	v442 = m.ExcPending
	if v442 != 0 {
		goto L10
	} else {
		goto L109
	}
L109:
	;
	m.G0 = v13 + int32(320)
	return
L110:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_beforeNextClient(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
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
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
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
	var v57 int32
	_ = v57
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
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
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v99 int32
	_ = v99
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v134 int32
	_ = v134
	var v137 int32
	_ = v137
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v147 int32
	_ = v147
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	v5 = int32(1)
	v6 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+200)))
	if v6&v5 != 0 {
		v13 = v5
		v16 = v13
	} else {
		v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+216))
		if v9 != 0 {
			v11 = F_isImportSlotMigrationJob(m, v9)
			mBase = m.M
			v13 = v11
			v16 = v13
		} else {
			v16 = int32(0)
		}
	}
	if v16 == int32(0) {
		F_trimClientQueryBuffer(m, l0)
		mBase = m.M
		v134 = m.ExcPending
		if v134 != 0 {
			return
		} else {
			v137 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+201)))
			if v137&int32(4) == int32(0) {
				v144 = F_updateClientMemUsageAndBucket(m, l0)
				mBase = m.M
				v145 = m.ExcPending
				if v145 != 0 {
					return
				} else {
					v147 = *(*int32)(unsafe.Add(mBase, _consts[254]))
					if v147 == int32(1) {
						return
					} else {
						v150 = F_trySendWriteToIOThreads(m, l0)
						mBase = m.M
						v151 = m.ExcPending
						if v151 != 0 {
							return
						} else {
							return
						}
					}
				}
			} else {
				v142 = F_freeClient(m, l0)
				mBase = m.M
				v143 = m.ExcPending
				if v143 != 0 {
					return
				} else {
					return
				}
			}
		}
	} else {
		v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
		v20 = *(*int64)(unsafe.Add(mBase, uint32(v19)+56))
		if v20 == int64(0) {
		} else {
			v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
			v24 = base.I32_wrap_i64(v20)
			v25 = int32(-1)
			v33 = v23 + v25
			v34 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v33))))
			v36 = v34 & int32(7)
			switch v36 {
			case 0:
				v51 = int32(base.Ui32(v34) >> (uint(int32(3)) % 32))
				if v51 == int32(0) {
				} else {
					v57 = int32(-1)&v51 + v25
					v61 = v24>>(uint(int32(31))%32)&v51 + v24
					v64 = v57 - v61 + int32(1)
					switch v36 {
					default:
						v79 = int32(base.Ui32(v34) >> (uint(int32(3)) % 32))
					case 1:
						v69 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23+int32(-3)))))
						v79 = v69
					case 2:
						v72 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v23+int32(-5)))))
						v79 = v72
					case 3:
						v75 = *(*int32)(unsafe.Add(mBase, uint32(v23+int32(-9))))
						v79 = v75
					case 4:
						v78 = *(*int32)(unsafe.Add(mBase, uint32(v23+int32(-17))))
						v79 = v78
					}
					v80 = int32(0)
					v82 = base.B2i32(base.Ui32(v61) < base.Ui32(v79))
					if base.Ui32(v61) < base.Ui32(v79) {
						v83 = v61
					} else {
						v83 = v80
					}
					v84 = v79 - v83
					if base.Ui32(v64) < base.Ui32(v84) {
						v86 = v64
					} else {
						v86 = v84
					}
					if v57 < v61 {
						v88 = v80
					} else {
						v88 = v86
					}
					if base.Ui32(v61) < base.Ui32(v79) {
						v90 = v88
					} else {
						v90 = int32(0)
					}
					if v90 == int32(0) {
					} else {
						v94 = F_memmove(m, v23, v23+v83, v90)
						mBase = m.M
					}
					v96 = int32(0)
					*(*uint8)(unsafe.Add(mBase, uint32(v23+v90))) = uint8(v96)
					switch v36 {
					default:
						v99 = v90 << (uint(int32(3)) % 32)
						*(*uint8)(unsafe.Add(mBase, uint32(v33))) = uint8(v99)
					case 1:
						*(*uint8)(unsafe.Add(mBase, uint32(v23+int32(-3)))) = uint8(v90)
					case 2:
						*(*uint16)(unsafe.Add(mBase, uint32(v23+int32(-5)))) = uint16(v90)
					case 3:
						*(*int32)(unsafe.Add(mBase, uint32(v23+int32(-9)))) = v90
					case 4:
						*(*int64)(unsafe.Add(mBase, uint32(v23+int32(-17)))) = base.I64_extend_i32_u(v90)
					}
				}
			case 1:
				v41 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23+int32(-3)))))
				v51 = v41
				if v51 == int32(0) {
				} else {
					v57 = int32(-1)&v51 + v25
					v61 = v24>>(uint(int32(31))%32)&v51 + v24
					v64 = v57 - v61 + int32(1)
					switch v36 {
					default:
						v79 = int32(base.Ui32(v34) >> (uint(int32(3)) % 32))
					case 1:
						v69 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23+int32(-3)))))
						v79 = v69
					case 2:
						v72 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v23+int32(-5)))))
						v79 = v72
					case 3:
						v75 = *(*int32)(unsafe.Add(mBase, uint32(v23+int32(-9))))
						v79 = v75
					case 4:
						v78 = *(*int32)(unsafe.Add(mBase, uint32(v23+int32(-17))))
						v79 = v78
					}
					v80 = int32(0)
					v82 = base.B2i32(base.Ui32(v61) < base.Ui32(v79))
					if base.Ui32(v61) < base.Ui32(v79) {
						v83 = v61
					} else {
						v83 = v80
					}
					v84 = v79 - v83
					if base.Ui32(v64) < base.Ui32(v84) {
						v86 = v64
					} else {
						v86 = v84
					}
					if v57 < v61 {
						v88 = v80
					} else {
						v88 = v86
					}
					if base.Ui32(v61) < base.Ui32(v79) {
						v90 = v88
					} else {
						v90 = int32(0)
					}
					if v90 == int32(0) {
					} else {
						v94 = F_memmove(m, v23, v23+v83, v90)
						mBase = m.M
					}
					v96 = int32(0)
					*(*uint8)(unsafe.Add(mBase, uint32(v23+v90))) = uint8(v96)
					switch v36 {
					default:
						v99 = v90 << (uint(int32(3)) % 32)
						*(*uint8)(unsafe.Add(mBase, uint32(v33))) = uint8(v99)
					case 1:
						*(*uint8)(unsafe.Add(mBase, uint32(v23+int32(-3)))) = uint8(v90)
					case 2:
						*(*uint16)(unsafe.Add(mBase, uint32(v23+int32(-5)))) = uint16(v90)
					case 3:
						*(*int32)(unsafe.Add(mBase, uint32(v23+int32(-9)))) = v90
					case 4:
						*(*int64)(unsafe.Add(mBase, uint32(v23+int32(-17)))) = base.I64_extend_i32_u(v90)
					}
				}
			case 2:
				v44 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v23+int32(-5)))))
				v51 = v44
				if v51 == int32(0) {
				} else {
					v57 = int32(-1)&v51 + v25
					v61 = v24>>(uint(int32(31))%32)&v51 + v24
					v64 = v57 - v61 + int32(1)
					switch v36 {
					default:
						v79 = int32(base.Ui32(v34) >> (uint(int32(3)) % 32))
					case 1:
						v69 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23+int32(-3)))))
						v79 = v69
					case 2:
						v72 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v23+int32(-5)))))
						v79 = v72
					case 3:
						v75 = *(*int32)(unsafe.Add(mBase, uint32(v23+int32(-9))))
						v79 = v75
					case 4:
						v78 = *(*int32)(unsafe.Add(mBase, uint32(v23+int32(-17))))
						v79 = v78
					}
					v80 = int32(0)
					v82 = base.B2i32(base.Ui32(v61) < base.Ui32(v79))
					if base.Ui32(v61) < base.Ui32(v79) {
						v83 = v61
					} else {
						v83 = v80
					}
					v84 = v79 - v83
					if base.Ui32(v64) < base.Ui32(v84) {
						v86 = v64
					} else {
						v86 = v84
					}
					if v57 < v61 {
						v88 = v80
					} else {
						v88 = v86
					}
					if base.Ui32(v61) < base.Ui32(v79) {
						v90 = v88
					} else {
						v90 = int32(0)
					}
					if v90 == int32(0) {
					} else {
						v94 = F_memmove(m, v23, v23+v83, v90)
						mBase = m.M
					}
					v96 = int32(0)
					*(*uint8)(unsafe.Add(mBase, uint32(v23+v90))) = uint8(v96)
					switch v36 {
					default:
						v99 = v90 << (uint(int32(3)) % 32)
						*(*uint8)(unsafe.Add(mBase, uint32(v33))) = uint8(v99)
					case 1:
						*(*uint8)(unsafe.Add(mBase, uint32(v23+int32(-3)))) = uint8(v90)
					case 2:
						*(*uint16)(unsafe.Add(mBase, uint32(v23+int32(-5)))) = uint16(v90)
					case 3:
						*(*int32)(unsafe.Add(mBase, uint32(v23+int32(-9)))) = v90
					case 4:
						*(*int64)(unsafe.Add(mBase, uint32(v23+int32(-17)))) = base.I64_extend_i32_u(v90)
					}
				}
			case 3:
				v47 = *(*int32)(unsafe.Add(mBase, uint32(v23+int32(-9))))
				v51 = v47
				if v51 == int32(0) {
				} else {
					v57 = int32(-1)&v51 + v25
					v61 = v24>>(uint(int32(31))%32)&v51 + v24
					v64 = v57 - v61 + int32(1)
					switch v36 {
					default:
						v79 = int32(base.Ui32(v34) >> (uint(int32(3)) % 32))
					case 1:
						v69 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23+int32(-3)))))
						v79 = v69
					case 2:
						v72 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v23+int32(-5)))))
						v79 = v72
					case 3:
						v75 = *(*int32)(unsafe.Add(mBase, uint32(v23+int32(-9))))
						v79 = v75
					case 4:
						v78 = *(*int32)(unsafe.Add(mBase, uint32(v23+int32(-17))))
						v79 = v78
					}
					v80 = int32(0)
					v82 = base.B2i32(base.Ui32(v61) < base.Ui32(v79))
					if base.Ui32(v61) < base.Ui32(v79) {
						v83 = v61
					} else {
						v83 = v80
					}
					v84 = v79 - v83
					if base.Ui32(v64) < base.Ui32(v84) {
						v86 = v64
					} else {
						v86 = v84
					}
					if v57 < v61 {
						v88 = v80
					} else {
						v88 = v86
					}
					if base.Ui32(v61) < base.Ui32(v79) {
						v90 = v88
					} else {
						v90 = int32(0)
					}
					if v90 == int32(0) {
					} else {
						v94 = F_memmove(m, v23, v23+v83, v90)
						mBase = m.M
					}
					v96 = int32(0)
					*(*uint8)(unsafe.Add(mBase, uint32(v23+v90))) = uint8(v96)
					switch v36 {
					default:
						v99 = v90 << (uint(int32(3)) % 32)
						*(*uint8)(unsafe.Add(mBase, uint32(v33))) = uint8(v99)
					case 1:
						*(*uint8)(unsafe.Add(mBase, uint32(v23+int32(-3)))) = uint8(v90)
					case 2:
						*(*uint16)(unsafe.Add(mBase, uint32(v23+int32(-5)))) = uint16(v90)
					case 3:
						*(*int32)(unsafe.Add(mBase, uint32(v23+int32(-9)))) = v90
					case 4:
						*(*int64)(unsafe.Add(mBase, uint32(v23+int32(-17)))) = base.I64_extend_i32_u(v90)
					}
				}
			case 4:
				v50 = *(*int32)(unsafe.Add(mBase, uint32(v23+int32(-17))))
				v51 = v50
				if v51 == int32(0) {
				} else {
					v57 = int32(-1)&v51 + v25
					v61 = v24>>(uint(int32(31))%32)&v51 + v24
					v64 = v57 - v61 + int32(1)
					switch v36 {
					default:
						v79 = int32(base.Ui32(v34) >> (uint(int32(3)) % 32))
					case 1:
						v69 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23+int32(-3)))))
						v79 = v69
					case 2:
						v72 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v23+int32(-5)))))
						v79 = v72
					case 3:
						v75 = *(*int32)(unsafe.Add(mBase, uint32(v23+int32(-9))))
						v79 = v75
					case 4:
						v78 = *(*int32)(unsafe.Add(mBase, uint32(v23+int32(-17))))
						v79 = v78
					}
					v80 = int32(0)
					v82 = base.B2i32(base.Ui32(v61) < base.Ui32(v79))
					if base.Ui32(v61) < base.Ui32(v79) {
						v83 = v61
					} else {
						v83 = v80
					}
					v84 = v79 - v83
					if base.Ui32(v64) < base.Ui32(v84) {
						v86 = v64
					} else {
						v86 = v84
					}
					if v57 < v61 {
						v88 = v80
					} else {
						v88 = v86
					}
					if base.Ui32(v61) < base.Ui32(v79) {
						v90 = v88
					} else {
						v90 = int32(0)
					}
					if v90 == int32(0) {
					} else {
						v94 = F_memmove(m, v23, v23+v83, v90)
						mBase = m.M
					}
					v96 = int32(0)
					*(*uint8)(unsafe.Add(mBase, uint32(v23+v90))) = uint8(v96)
					switch v36 {
					default:
						v99 = v90 << (uint(int32(3)) % 32)
						*(*uint8)(unsafe.Add(mBase, uint32(v33))) = uint8(v99)
					case 1:
						*(*uint8)(unsafe.Add(mBase, uint32(v23+int32(-3)))) = uint8(v90)
					case 2:
						*(*uint16)(unsafe.Add(mBase, uint32(v23+int32(-5)))) = uint16(v90)
					case 3:
						*(*int32)(unsafe.Add(mBase, uint32(v23+int32(-9)))) = v90
					case 4:
						*(*int64)(unsafe.Add(mBase, uint32(v23+int32(-17)))) = base.I64_extend_i32_u(v90)
					}
				}
			default:
			}
			v126 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
			v127 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
			v128 = *(*int32)(unsafe.Add(mBase, uint32(v127)+56))
			*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v126 - v128
			*(*int64)(unsafe.Add(mBase, uint32(v127)+56)) = int64(0)
		}
		v137 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+201)))
		if v137&int32(4) == int32(0) {
			v144 = F_updateClientMemUsageAndBucket(m, l0)
			mBase = m.M
			v145 = m.ExcPending
			if v145 != 0 {
				return
			} else {
				v147 = *(*int32)(unsafe.Add(mBase, _consts[254]))
				if v147 == int32(1) {
					return
				} else {
					v150 = F_trySendWriteToIOThreads(m, l0)
					mBase = m.M
					v151 = m.ExcPending
					if v151 != 0 {
						return
					} else {
						return
					}
				}
			}
		} else {
			v142 = F_freeClient(m, l0)
			mBase = m.M
			v143 = m.ExcPending
			if v143 != 0 {
				return
			} else {
				return
			}
		}
	}
}
func F_beforeSleep(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v22 int64
	_ = v22
	var v24 int64
	_ = v24
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v77 int32
	_ = v77
	var v86 int32
	_ = v86
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v99 int32
	_ = v99
	var v108 int32
	_ = v108
	var v112 int32
	_ = v112
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v119 int32
	_ = v119
	var v126 int32
	_ = v126
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v133 int32
	_ = v133
	var v136 int32
	_ = v136
	var v139 int32
	_ = v139
	var v142 int32
	_ = v142
	var v145 int32
	_ = v145
	var v148 int32
	_ = v148
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v156 int32
	_ = v156
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v170 int32
	_ = v170
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v183 int32
	_ = v183
	var v185 int32
	_ = v185
	var v193 int32
	_ = v193
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v209 int32
	_ = v209
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v225 int64
	_ = v225
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v234 int64
	_ = v234
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v241 int64
	_ = v241
	var v247 int32
	_ = v247
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v254 int32
	_ = v254
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v258 int32
	_ = v258
	var v259 int32
	_ = v259
	var v263 int32
	_ = v263
	var v265 int32
	_ = v265
	var v266 int32
	_ = v266
	var v267 int32
	_ = v267
	var v268 int64
	_ = v268
	var v269 int64
	_ = v269
	var v271 int32
	_ = v271
	var v275 int32
	_ = v275
	var v276 int32
	_ = v276
	var v277 int32
	_ = v277
	var v281 int32
	_ = v281
	var v282 int32
	_ = v282
	var v283 int32
	_ = v283
	var v287 int32
	_ = v287
	var v289 int64
	_ = v289
	var v290 int32
	_ = v290
	var v291 int64
	_ = v291
	var v294 int32
	_ = v294
	var v295 int32
	_ = v295
	var v296 int32
	_ = v296
	var v301 int32
	_ = v301
	var v304 int32
	_ = v304
	var v305 int32
	_ = v305
	var v306 int32
	_ = v306
	var v310 int32
	_ = v310
	var v311 int32
	_ = v311
	var v312 int32
	_ = v312
	var v313 int32
	_ = v313
	var v316 int64
	_ = v316
	var v323 int32
	_ = v323
	var v324 int32
	_ = v324
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
	var v338 int32
	_ = v338
	var v339 int64
	_ = v339
	var v341 int64
	_ = v341
	var v343 int32
	_ = v343
	var v350 int32
	_ = v350
	var v353 int32
	_ = v353
	var v354 int64
	_ = v354
	var v355 int64
	_ = v355
	var v365 int32
	_ = v365
	var v366 int64
	_ = v366
	var v371 int32
	_ = v371
	var v372 int64
	_ = v372
	var v377 int64
	_ = v377
	var v384 int32
	_ = v384
	var v388 int64
	_ = v388
	var v391 int32
	_ = v391
	var v393 int64
	_ = v393
	var v397 int32
	_ = v397
	var v399 int32
	_ = v399
	var v400 int32
	_ = v400
	var v401 int32
	_ = v401
	var v402 int32
	_ = v402
	var v404 int32
	_ = v404
	var v405 int32
	_ = v405
	var v409 int32
	_ = v409
	var v412 int32
	_ = v412
	var v415 int32
	_ = v415
	var v416 int32
	_ = v416
	var v420 int32
	_ = v420
	var v422 int32
	_ = v422
	var v423 int32
	_ = v423
	var v425 int32
	_ = v425
	var v426 int64
	_ = v426
	var v427 int32
	_ = v427
	var v428 int32
	_ = v428
	var v429 int32
	_ = v429
	var v430 int32
	_ = v430
	var v435 int32
	_ = v435
	var v437 int32
	_ = v437
	var v438 int32
	_ = v438
	var v439 int32
	_ = v439
	var v440 int64
	_ = v440
	var v442 int64
	_ = v442
	var v446 int64
	_ = v446
	var v456 int32
	_ = v456
	var v457 int64
	_ = v457
	var v462 int32
	_ = v462
	var v463 int64
	_ = v463
	var v468 int64
	_ = v468
	var v474 int32
	_ = v474
	var v476 int32
	_ = v476
	var v477 int64
	_ = v477
	var v479 int64
	_ = v479
	var v483 int32
	_ = v483
	var v488 int64
	_ = v488
	var v489 int64
	_ = v489
	var v501 int32
	_ = v501
	var v502 int64
	_ = v502
	var v507 int32
	_ = v507
	var v508 int64
	_ = v508
	var v513 int64
	_ = v513
	var v519 int32
	_ = v519
	var v523 int64
	_ = v523
	var v525 int64
	_ = v525
	var v527 int64
	_ = v527
	var v529 int64
	_ = v529
	var v534 int32
	_ = v534
	var v535 int32
	_ = v535
	var v536 int32
	_ = v536
	var v546 int32
	_ = v546
	var v549 int32
	_ = v549
	var v550 int32
	_ = v550
	var v551 int32
	_ = v551
	var v556 int32
	_ = v556
	var v574 int32
	_ = v574
	var v580 int32
	_ = v580
	v11 = m.G0
	v13 = v11 - int32(16)
	m.G0 = v13
	v18 = *(*int32)(unsafe.Add(mBase, _consts[347]))
	if v18 < int32(2) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	v99 = int32(0)
	v108 = *(*int32)(unsafe.Add(mBase, _consts[315]))
	if v108 < int32(261) {
		goto L16
	} else {
		goto L17
	}
L2:
	;
	goto L1
L3:
	;
	v21 = int32(_a20)
	v22 = *(*int64)(unsafe.Add(mBase, _consts[361]))
	v24 = *(*int64)(unsafe.Add(mBase, _consts[362]))
	if base.I32_wrap_i64(v22+v24) == int32(0) {
		goto L2
	} else {
		goto L4
	}
L4:
	;
	v30 = *(*int32)(unsafe.Add(mBase, _consts[363]))
	if v30 != 0 {
		goto L2
	} else {
		goto L5
	}
L5:
	;
	v31 = int32(_a20)
	v32 = int32(1)
	*(*int32)(unsafe.Add(mBase, _consts[363])) = v32
	v35 = *(*int32)(unsafe.Add(mBase, _consts[223]))
	F_aeSetPollProtect(m, v35, v32)
	mBase = m.M
	v39 = v35 | int32(4)
	v41 = *(*int32)(unsafe.Add(mBase, _consts[347]))
	if int32(9) < v41 {
		goto L7
	} else {
		goto L8
	}
L6:
	;
	v86 = *(*int32)(unsafe.Add(mBase, _consts[223]))
	F_aeSetCustomPollProc(m, v86, int32(542))
	mBase = m.M
	v89 = int32(0)
	v91 = *(*int32)(unsafe.Add(mBase, _consts[357]))
	*(*int32)(unsafe.Add(mBase, _consts[357])) = v91 + int32(1)
	goto L2
L7:
	;
	v53 = int32(0)
	v55 = *(*int32)(unsafe.Add(mBase, _consts[717]))
	v58 = base.I32_rem_s(v55, v41+int32(-1))
	v60 = v58 + int32(1)
	*(*int32)(unsafe.Add(mBase, _consts[717])) = v60
	v66 = F_spscIsFull(m, v60*int32(192)+int32(_a674))
	mBase = m.M
	if v66 == v53 {
		goto L10
	} else {
		goto L11
	}
L8:
	;
	v45 = F_spmcEnqueue(m, int32(_a2182), v39)
	mBase = m.M
	if v45 != 0 {
		goto L6
	} else {
		goto L9
	}
L9:
	;
	v46 = int32(_a20)
	v47 = int32(0)
	*(*int32)(unsafe.Add(mBase, _consts[363])) = v47
	v50 = *(*int32)(unsafe.Add(mBase, _consts[223]))
	F_aeSetPollProtect(m, v50, v47)
	mBase = m.M
	goto L1
L10:
	;
	v77 = *(*int32)(unsafe.Add(mBase, _consts[717]))
	F_spscEnqueue(m, v77*int32(192)+int32(_a674), v39, int32(1))
	mBase = m.M
	goto L6
L11:
	;
	v69 = int32(_a20)
	v70 = int32(0)
	*(*int32)(unsafe.Add(mBase, _consts[363])) = v70
	v73 = *(*int32)(unsafe.Add(mBase, _consts[223]))
	F_aeSetPollProtect(m, v73, v70)
	mBase = m.M
	goto L1
L12:
	;
	v198 = *(*int32)(unsafe.Add(mBase, _consts[718]))
	v199 = F_processIOThreadsResponses(m)
	mBase = m.M
	v200 = m.ExcPending
	if v200 != 0 {
		goto L30
	} else {
		goto L31
	}
L13:
	;
	v193 = *(*int32)(unsafe.Add(mBase, _consts[554]))
	if base.Ui32(v185) <= base.Ui32(v193) {
		goto L12
	} else {
		goto L29
	}
L14:
	;
	goto L13
L15:
	;
	v119 = v117 & int32(3)
	if base.Ui32(int32(4)) <= base.Ui32(v117) {
		goto L20
	} else {
		goto L21
	}
L16:
	;
	if v108 < int32(1) {
		v185 = v99
		goto L14
	} else {
		goto L18
	}
L17:
	;
	v112 = *(*int32)(unsafe.Add(mBase, _consts[316]))
	v116 = v112
	v117 = int32(260)
	goto L15
L18:
	;
	v116 = v99
	v117 = v108
	goto L15
L19:
	;
	if v119 == int32(0) {
		v185 = v158
		goto L14
	} else {
		goto L25
	}
L20:
	;
	v126 = int32(0)
	v128 = v116
	v129 = v126
	v133 = v126
	goto L22
L21:
	;
	v158 = v116
	v159 = int32(0)
	goto L19
L22:
	;
	v136 = v129 << (uint(int32(2)) % 32)
	v139 = *(*int32)(unsafe.Add(mBase, uint32(v136)+uint32(_consts[317])))
	v142 = *(*int32)(unsafe.Add(mBase, uint32(v136)+uint32(_consts[318])))
	v145 = *(*int32)(unsafe.Add(mBase, uint32(v136)+uint32(_consts[319])))
	v148 = *(*int32)(unsafe.Add(mBase, uint32(v136)+uint32(_consts[320])))
	v152 = v139 + (v142 + (v145 + (v148 + v128)))
	v153 = int32(4)
	v154 = v129 + v153
	v156 = v133 + v153
	if v156 != v117&int32(2147483644) {
		v128 = v152
		v129 = v154
		v133 = v156
		goto L22
	} else {
		goto L24
	}
L23:
	;
	v158 = v152
	v159 = v154
	goto L19
L24:
	;
	goto L23
L25:
	;
	v167 = v158
	v168 = v159
	v170 = int32(0)
	goto L26
L26:
	;
	v178 = *(*int32)(unsafe.Add(mBase, uint32(v168<<(uint(int32(2))%32))+uint32(_consts[320])))
	v179 = v178 + v167
	v180 = int32(1)
	v183 = v170 + v180
	if v183 != v119 {
		v167 = v179
		v168 = v168 + v180
		v170 = v183
		goto L26
	} else {
		goto L28
	}
L27:
	;
	v185 = v179
	goto L14
L28:
	;
	goto L27
L29:
	;
	*(*int32)(unsafe.Add(mBase, _consts[554])) = v185
	goto L12
L30:
	;
	return
L31:
	;
	if v198 == int32(0) {
		goto L35
	} else {
		goto L36
	}
L32:
	;
	F__serverAssert(m, int32(_a2183), int32(_a2157), int32(1904))
	mBase = m.M
	v580 = m.ExcPending
	if v580 != 0 {
		goto L30
	} else {
		goto L132
	}
L33:
	;
	F__serverAssert(m, int32(_a2184), int32(_a2157), int32(1903))
	mBase = m.M
	v574 = m.ExcPending
	if v574 != 0 {
		goto L30
	} else {
		goto L131
	}
L34:
	;
	m.G0 = v13 + int32(16)
	return
L35:
	;
	if v199 < int32(1) {
		goto L47
	} else {
		goto L48
	}
L36:
	;
	v203 = F_connTypeProcessPendingData(m)
	mBase = m.M
	v204 = m.ExcPending
	if v204 != 0 {
		goto L30
	} else {
		goto L37
	}
L37:
	;
	v209 = *(*int32)(unsafe.Add(mBase, _consts[38]))
	if base.Ui32(int32(1)) < base.Ui32(v209+int32(-1)) {
		goto L38
	} else {
		goto L39
	}
L38:
	;
	v217 = F_handleClientsWithPendingWrites(m)
	mBase = m.M
	v218 = m.ExcPending
	if v218 != 0 {
		goto L30
	} else {
		goto L41
	}
L39:
	;
	F_flushAppendOnlyFile(m, int32(0))
	mBase = m.M
	v216 = m.ExcPending
	if v216 != 0 {
		goto L30
	} else {
		goto L40
	}
L40:
	;
	goto L38
L41:
	;
	v225 = base.I64_extend_i32_s(v203) + base.I64_extend_i32_s(v199) + base.I64_extend_i32_s(v217)
	goto L42
L42:
	;
	v231 = F_processIOThreadsResponses(m)
	mBase = m.M
	v232 = m.ExcPending
	if v232 != 0 {
		goto L30
	} else {
		goto L44
	}
L43:
	;
	v236 = F_freeClientsInAsyncFreeQueue(m)
	mBase = m.M
	v237 = m.ExcPending
	if v237 != 0 {
		goto L30
	} else {
		goto L46
	}
L44:
	;
	v234 = v225 + base.I64_extend_i32_s(v231)
	if v231 != 0 {
		v225 = v234
		goto L42
	} else {
		goto L45
	}
L45:
	;
	goto L43
L46:
	;
	v241 = *(*int64)(unsafe.Add(mBase, _consts[719]))
	*(*int64)(unsafe.Add(mBase, _consts[719])) = v234 + base.I64_extend_i32_s(v236) + v241
	goto L34
L47:
	;
	v249 = F_connTypeProcessPendingData(m)
	mBase = m.M
	v250 = m.ExcPending
	if v250 != 0 {
		goto L30
	} else {
		goto L50
	}
L48:
	;
	v247 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _consts[720])) = uint8(v247)
	goto L47
L49:
	;
	v256 = F_connTypeHasPendingData(m)
	mBase = m.M
	v257 = m.ExcPending
	if v257 != 0 {
		goto L30
	} else {
		goto L52
	}
L50:
	;
	if v249 < int32(1) {
		goto L49
	} else {
		goto L51
	}
L51:
	;
	v254 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _consts[720])) = uint8(v254)
	goto L49
L52:
	;
	v258 = int32(0)
	v259 = *(*int32)(unsafe.Add(mBase, _consts[110]))
	if v259 == v258 {
		goto L53
	} else {
		goto L54
	}
L53:
	;
	F_blockedBeforeSleep(m)
	mBase = m.M
	v265 = m.ExcPending
	if v265 != 0 {
		goto L30
	} else {
		goto L56
	}
L54:
	;
	F_clusterBeforeSleep(m)
	mBase = m.M
	v263 = m.ExcPending
	if v263 != 0 {
		goto L30
	} else {
		goto L55
	}
L55:
	;
	goto L53
L56:
	;
	v266 = int32(0)
	v267 = *(*int32)(unsafe.Add(mBase, _consts[34]))
	v268 = m.T0[v267].(func(*base.Module) int64)(m)
	mBase = m.M
	v269 = int64(0)
	v271 = *(*int32)(unsafe.Add(mBase, _consts[721]))
	if v271 == v266 {
		v291 = v269
		goto L57
	} else {
		goto L58
	}
L57:
	;
	v294 = *(*int32)(unsafe.Add(mBase, _consts[81]))
	v295 = *(*int32)(unsafe.Add(mBase, uint32(v294)+16))
	v296 = *(*int32)(unsafe.Add(mBase, uint32(v294)+12))
	goto L69
L58:
	;
	v275 = *(*int32)(unsafe.Add(mBase, _consts[505]))
	if v275 != 0 {
		v291 = v269
		goto L57
	} else {
		goto L59
	}
L59:
	;
	v276 = int32(0)
	v277 = *(*int32)(unsafe.Add(mBase, _consts[110]))
	if v277 == v276 {
		goto L61
	} else {
		goto L62
	}
L60:
	;
	v289 = F_activeExpireCycle(m, int32(1))
	mBase = m.M
	v290 = m.ExcPending
	if v290 != 0 {
		goto L30
	} else {
		goto L67
	}
L61:
	;
	v287 = *(*int32)(unsafe.Add(mBase, _consts[166]))
	if v287 != 0 {
		v291 = v269
		goto L57
	} else {
		goto L66
	}
L62:
	;
	v281 = *(*int32)(unsafe.Add(mBase, _consts[111]))
	v282 = *(*int32)(unsafe.Add(mBase, uint32(v281)))
	goto L63
L63:
	;
	v283 = *(*int32)(unsafe.Add(mBase, uint32(v282)+88))
	goto L64
L64:
	;
	if v283&int32(1) != 0 {
		goto L60
	} else {
		goto L65
	}
L65:
	;
	v291 = v269
	goto L57
L66:
	;
	goto L60
L67:
	;
	v291 = v289
	goto L57
L68:
	;
	v305 = int32(0)
	v306 = *(*int32)(unsafe.Add(mBase, _consts[637]))
	if v306 == v305 {
		goto L72
	} else {
		goto L73
	}
L69:
	;
	if v295+v296 == int32(0) {
		goto L68
	} else {
		goto L70
	}
L70:
	;
	v301 = int32(0)
	F_moduleFireServerEvent(m, int64(15), v301, v301)
	mBase = m.M
	v304 = m.ExcPending
	if v304 != 0 {
		goto L30
	} else {
		goto L71
	}
L71:
	;
	goto L68
L72:
	;
	F_updateFailoverStatus(m)
	mBase = m.M
	v328 = m.ExcPending
	if v328 != 0 {
		goto L30
	} else {
		goto L77
	}
L73:
	;
	v310 = F_isPausedActionsWithUpdate(m, int32(16))
	mBase = m.M
	v311 = m.ExcPending
	if v311 != 0 {
		goto L30
	} else {
		goto L74
	}
L74:
	;
	if v310 != 0 {
		goto L72
	} else {
		goto L75
	}
L75:
	;
	v312 = int32(0)
	v313 = *(*int32)(unsafe.Add(mBase, _consts[722]))
	*(*int32)(unsafe.Add(mBase, uint32(v13)+4)) = v313
	v316 = *(*int64)(unsafe.Add(mBase, _consts[723]))
	*(*int64)(unsafe.Add(mBase, uint32(v13)+8)) = v316
	F_replicationFeedReplicas(m, int32(-1), v13+int32(4), int32(3))
	mBase = m.M
	v323 = m.ExcPending
	if v323 != 0 {
		goto L30
	} else {
		goto L76
	}
L76:
	;
	v324 = int32(0)
	*(*int32)(unsafe.Add(mBase, _consts[637])) = v324
	goto L72
L77:
	;
	v330 = *(*int32)(unsafe.Add(mBase, _consts[724]))
	v331 = *(*int32)(unsafe.Add(mBase, uint32(v330)+20))
	if v331 != 0 {
		goto L33
	} else {
		goto L78
	}
L78:
	;
	v333 = *(*int32)(unsafe.Add(mBase, _consts[485]))
	v334 = *(*int32)(unsafe.Add(mBase, uint32(v333)+20))
	if v334 != 0 {
		goto L32
	} else {
		goto L79
	}
L79:
	;
	F_trackingBroadcastInvalidationMessages(m)
	mBase = m.M
	v336 = m.ExcPending
	if v336 != 0 {
		goto L30
	} else {
		goto L80
	}
L80:
	;
	v337 = int32(0)
	v338 = *(*int32)(unsafe.Add(mBase, _consts[34]))
	v339 = m.T0[v338].(func(*base.Module) int64)(m)
	mBase = m.M
	v341 = *(*int64)(unsafe.Add(mBase, _consts[49]))
	v343 = *(*int32)(unsafe.Add(mBase, _consts[38]))
	if base.Ui32(int32(1)) < base.Ui32(v343+int32(-1)) {
		goto L81
	} else {
		goto L82
	}
L81:
	;
	v353 = *(*int32)(unsafe.Add(mBase, _consts[34]))
	v354 = m.T0[v353].(func(*base.Module) int64)(m)
	mBase = m.M
	v355 = v354 - v339
	goto L86
L82:
	;
	F_flushAppendOnlyFile(m, int32(0))
	mBase = m.M
	v350 = m.ExcPending
	if v350 != 0 {
		goto L30
	} else {
		goto L83
	}
L83:
	;
	goto L81
L84:
	;
	v384 = *(*int32)(unsafe.Add(mBase, _consts[38]))
	if v384 != int32(1) {
		v402 = v256
		goto L88
	} else {
		goto L89
	}
L85:
	;
	goto L84
L86:
	;
	v365 = int32(_a2185)
	v366 = *(*int64)(unsafe.Add(mBase, _consts[725]))
	*(*int64)(unsafe.Add(mBase, _consts[725])) = v366 + int64(1)
	v371 = int32(_a2186)
	v372 = *(*int64)(unsafe.Add(mBase, _consts[726]))
	*(*int64)(unsafe.Add(mBase, _consts[726])) = v372 + v355
	v377 = *(*int64)(unsafe.Add(mBase, _consts[727]))
	if base.Ui64(v355) <= base.Ui64(v377) {
		goto L85
	} else {
		goto L87
	}
L87:
	;
	*(*int64)(unsafe.Add(mBase, _consts[727])) = v355
	goto L85
L88:
	;
	v404 = F_handleClientsWithPendingWrites(m)
	mBase = m.M
	v405 = m.ExcPending
	if v405 != 0 {
		goto L30
	} else {
		goto L98
	}
L89:
	;
	v388 = *(*int64)(unsafe.Add(mBase, _consts[49]))
	if v388 == int64(-1) {
		v402 = v256
		goto L88
	} else {
		goto L90
	}
L90:
	;
	v391 = int32(0)
	v393 = *(*int64)(unsafe.Add(mBase, _consts[48]))
	*(*int64)(unsafe.Add(mBase, _consts[49])) = v393
	if v341 == v393 {
		goto L91
	} else {
		goto L92
	}
L91:
	;
	v397 = v256
	goto L93
L92:
	;
	v397 = int32(1)
	goto L93
L93:
	;
	v399 = *(*int32)(unsafe.Add(mBase, _consts[80]))
	v400 = *(*int32)(unsafe.Add(mBase, uint32(v399)+20))
	if v400 != 0 {
		goto L94
	} else {
		goto L95
	}
L94:
	;
	v401 = v397
	goto L96
L95:
	;
	v401 = v256
	goto L96
L96:
	;
	v402 = v401
	goto L88
L97:
	;
	v412 = *(*int32)(unsafe.Add(mBase, _consts[254]))
	if v412 == int32(1) {
		goto L100
	} else {
		goto L101
	}
L98:
	;
	if v404 < int32(1) {
		goto L97
	} else {
		goto L99
	}
L99:
	;
	v409 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _consts[720])) = uint8(v409)
	goto L97
L100:
	;
	v425 = *(*int32)(unsafe.Add(mBase, _consts[34]))
	v426 = m.T0[v425].(func(*base.Module) int64)(m)
	mBase = m.M
	v427 = F_freeClientsInAsyncFreeQueue(m)
	mBase = m.M
	v428 = m.ExcPending
	if v428 != 0 {
		goto L30
	} else {
		goto L105
	}
L101:
	;
	v415 = F_processIOThreadsResponses(m)
	mBase = m.M
	v416 = m.ExcPending
	if v416 != 0 {
		goto L30
	} else {
		goto L102
	}
L102:
	;
	if v415 < int32(1) {
		goto L100
	} else {
		goto L103
	}
L103:
	;
	v420 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _consts[720])) = uint8(v420)
	v422 = F_handleClientsWithPendingWrites(m)
	mBase = m.M
	v423 = m.ExcPending
	if v423 != 0 {
		goto L30
	} else {
		goto L104
	}
L104:
	;
	goto L100
L105:
	;
	v429 = int32(0)
	v430 = *(*int32)(unsafe.Add(mBase, _consts[293]))
	if v430 == v429 {
		goto L106
	} else {
		goto L107
	}
L106:
	;
	F_evictClients(m)
	mBase = m.M
	v437 = m.ExcPending
	if v437 != 0 {
		goto L30
	} else {
		goto L109
	}
L107:
	;
	F_incrementalTrimReplicationBacklog(m, int32(640))
	mBase = m.M
	v435 = m.ExcPending
	if v435 != 0 {
		goto L30
	} else {
		goto L108
	}
L108:
	;
	goto L106
L109:
	;
	v438 = int32(0)
	v439 = *(*int32)(unsafe.Add(mBase, _consts[34]))
	v440 = m.T0[v439].(func(*base.Module) int64)(m)
	mBase = m.M
	v442 = *(*int64)(unsafe.Add(mBase, _consts[728]))
	if v442 == int64(0) {
		goto L110
	} else {
		goto L111
	}
L110:
	;
	v483 = int32(0)
	v488 = *(*int64)(unsafe.Add(mBase, _consts[729]))
	v489 = v339 - (v268 + v426) + v440 + v488
	*(*int64)(unsafe.Add(mBase, _consts[729])) = v489
	goto L121
L111:
	;
	v446 = v440 - v442
	goto L114
L112:
	;
	v474 = int32(0)
	v476 = int32(*(*uint8)(unsafe.Add(mBase, _consts[720])))
	if v476 != 0 {
		goto L116
	} else {
		goto L117
	}
L113:
	;
	goto L112
L114:
	;
	v456 = int32(_a2187)
	v457 = *(*int64)(unsafe.Add(mBase, _consts[370]))
	*(*int64)(unsafe.Add(mBase, _consts[370])) = v457 + int64(1)
	v462 = int32(_a2188)
	v463 = *(*int64)(unsafe.Add(mBase, _consts[371]))
	*(*int64)(unsafe.Add(mBase, _consts[371])) = v463 + v446
	v468 = *(*int64)(unsafe.Add(mBase, _consts[372]))
	if base.Ui64(v446) <= base.Ui64(v468) {
		goto L113
	} else {
		goto L115
	}
L115:
	;
	*(*int64)(unsafe.Add(mBase, _consts[372])) = v446
	goto L113
L116:
	;
	v477 = v446
	goto L118
L117:
	;
	v477 = v291
	goto L118
L118:
	;
	v479 = *(*int64)(unsafe.Add(mBase, _consts[349]))
	*(*int64)(unsafe.Add(mBase, _consts[349])) = v477 + v479
	goto L110
L119:
	;
	v519 = int32(0)
	*(*int64)(unsafe.Add(mBase, _consts[729])) = int64(0)
	v523 = *(*int64)(unsafe.Add(mBase, _consts[730]))
	v525 = *(*int64)(unsafe.Add(mBase, _consts[731]))
	if v523 <= v525 {
		goto L123
	} else {
		goto L124
	}
L120:
	;
	goto L119
L121:
	;
	v501 = int32(_a2189)
	v502 = *(*int64)(unsafe.Add(mBase, _consts[732]))
	*(*int64)(unsafe.Add(mBase, _consts[732])) = v502 + int64(1)
	v507 = int32(_a2190)
	v508 = *(*int64)(unsafe.Add(mBase, _consts[733]))
	*(*int64)(unsafe.Add(mBase, _consts[733])) = v508 + v489
	v513 = *(*int64)(unsafe.Add(mBase, _consts[734]))
	if base.Ui64(v489) <= base.Ui64(v513) {
		goto L120
	} else {
		goto L122
	}
L122:
	;
	*(*int64)(unsafe.Add(mBase, _consts[734])) = v489
	goto L120
L123:
	;
	v534 = int32(0)
	v535 = *(*int32)(unsafe.Add(mBase, _consts[223]))
	v536 = *(*int32)(unsafe.Add(mBase, uint32(v535)+72))
	*(*int32)(unsafe.Add(mBase, uint32(v535)+72)) = v536&int32(-5) | base.B2i32(v402 != v534)<<(uint(int32(2))%32)
	goto L126
L124:
	;
	v527 = v523 - v525
	v529 = *(*int64)(unsafe.Add(mBase, _consts[735]))
	if v527 <= v529 {
		goto L123
	} else {
		goto L125
	}
L125:
	;
	*(*int64)(unsafe.Add(mBase, _consts[735])) = v527
	goto L123
L126:
	;
	F_IOThreadsBeforeSleep(m, v440)
	mBase = m.M
	v546 = m.ExcPending
	if v546 != 0 {
		goto L30
	} else {
		goto L127
	}
L127:
	;
	v549 = *(*int32)(unsafe.Add(mBase, _consts[81]))
	v550 = *(*int32)(unsafe.Add(mBase, uint32(v549)+16))
	v551 = *(*int32)(unsafe.Add(mBase, uint32(v549)+12))
	goto L128
L128:
	;
	if v550+v551 == int32(0) {
		goto L34
	} else {
		goto L129
	}
L129:
	;
	v556 = F___pthread_mutex_unlock(m, int32(_a2191))
	mBase = m.M
	goto L130
L130:
	;
	goto L34
L131:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L132:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_bg_unlink(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v18 int32
	_ = v18
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	v6 = F_open(m, l0, int32(2048), int32(0))
	mBase = m.M
	v7 = F_unlink(m, l0)
	mBase = m.M
	if v6 == int32(-1) {
		v25 = v7
		return v25
	} else {
		if v7 != int32(-1) {
			v18 = int32(0)
			F_bioCreateCloseJob(m, v6, v18, v18)
			mBase = m.M
			v24 = m.ExcPending
			if v24 != 0 {
				return int32(0)
			} else {
				v25 = v18
				return v25
			}
		} else {
			v12 = int32(9116376)
			v13 = *(*int32)(unsafe.Add(mBase, _consts[18]))
			v14 = F_close(m, v6)
			mBase = m.M
			*(*int32)(unsafe.Add(mBase, _consts[18])) = v13
			return int32(-1)
		}
	}
}
func F_bgrewriteaofCommand(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v14 int32
	_ = v14
	var v18 int32
	_ = v18
	var v25 int32
	_ = v25
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	v3 = *(*int32)(unsafe.Add(mBase, _consts[60]))
	if v3 != int32(2) {
		v10 = *(*int32)(unsafe.Add(mBase, _consts[61]))
		v14 = *(*int32)(unsafe.Add(mBase, _consts[62]))
		if base.B2i32(v10 != int32(-1))|v14 == int32(0) {
			v36 = F_rewriteAppendOnlyFileBackground(m)
			mBase = m.M
			v37 = m.ExcPending
			if v37 != 0 {
				return
			} else {
				if v36 != 0 {
					F_addReplyError(m, l0, int32(_a161))
					mBase = m.M
					v43 = m.ExcPending
					if v43 != 0 {
						return
					} else {
						return
					}
				} else {
					F_addReplyStatus(m, l0, int32(_a162))
					mBase = m.M
					v40 = m.ExcPending
					if v40 != 0 {
						return
					} else {
						return
					}
				}
			}
		} else {
			v18 = int32(_a20)
			*(*int64)(unsafe.Add(mBase, _consts[52])) = int64(0)
			*(*int32)(unsafe.Add(mBase, _consts[63])) = int32(1)
			v25 = *(*int32)(unsafe.Add(mBase, _consts[28]))
			if int32(2) < v25 {
				F_addReplyStatus(m, l0, int32(_a163))
				mBase = m.M
				v35 = m.ExcPending
				if v35 != 0 {
					return
				} else {
					return
				}
			} else {
				F__serverLog(m, int32(2), int32(_a164), int32(0))
				mBase = m.M
				v32 = m.ExcPending
				if v32 != 0 {
					return
				} else {
					F_addReplyStatus(m, l0, int32(_a163))
					mBase = m.M
					v35 = m.ExcPending
					if v35 != 0 {
						return
					} else {
						return
					}
				}
			}
		}
	} else {
		F_addReplyError(m, l0, int32(_a165))
		mBase = m.M
		v8 = m.ExcPending
		if v8 != 0 {
			return
		} else {
			return
		}
	}
}
func F_bitcountCommand(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
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
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
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
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
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
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
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
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v120 int32
	_ = v120
	var v122 int32
	_ = v122
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
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
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v151 int64
	_ = v151
	var v152 int32
	_ = v152
	var v156 int64
	_ = v156
	var v158 int64
	_ = v158
	var v159 int64
	_ = v159
	var v160 int64
	_ = v160
	var v162 int32
	_ = v162
	var v168 int64
	_ = v168
	var v169 int64
	_ = v169
	var v171 int64
	_ = v171
	var v174 int64
	_ = v174
	var v176 int64
	_ = v176
	var v179 int64
	_ = v179
	var v182 int64
	_ = v182
	var v185 int64
	_ = v185
	var v188 int64
	_ = v188
	var v190 int64
	_ = v190
	var v195 int32
	_ = v195
	var v202 int32
	_ = v202
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
	var v210 int32
	_ = v210
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v223 int32
	_ = v223
	var v229 int32
	_ = v229
	var v231 int32
	_ = v231
	var v232 int64
	_ = v232
	var v233 int64
	_ = v233
	var v237 int32
	_ = v237
	var v241 int32
	_ = v241
	var v255 int32
	_ = v255
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v258 int64
	_ = v258
	var v259 int64
	_ = v259
	var v261 int32
	_ = v261
	var v265 int32
	_ = v265
	var v266 int32
	_ = v266
	var v267 int32
	_ = v267
	var v268 int64
	_ = v268
	var v269 int64
	_ = v269
	var v271 int32
	_ = v271
	var v275 int32
	_ = v275
	var v277 int32
	_ = v277
	var v280 int32
	_ = v280
	var v282 int32
	_ = v282
	var v284 int32
	_ = v284
	var v286 int32
	_ = v286
	var v288 int32
	_ = v288
	var v300 int32
	_ = v300
	var v301 int32
	_ = v301
	var v304 int64
	_ = v304
	var v306 int32
	_ = v306
	var v310 int32
	_ = v310
	var v311 int32
	_ = v311
	var v314 int64
	_ = v314
	var v315 int64
	_ = v315
	var v317 int32
	_ = v317
	var v321 int32
	_ = v321
	var v322 int32
	_ = v322
	var v325 int64
	_ = v325
	var v326 int64
	_ = v326
	var v327 int32
	_ = v327
	var v328 int32
	_ = v328
	var v335 int32
	_ = v335
	var v338 int64
	_ = v338
	var v340 int32
	_ = v340
	var v341 int64
	_ = v341
	var v342 int32
	_ = v342
	var v347 int32
	_ = v347
	var v348 int64
	_ = v348
	var v349 int32
	_ = v349
	var v353 int32
	_ = v353
	var v354 int32
	_ = v354
	var v356 int32
	_ = v356
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
	var v369 int32
	_ = v369
	var v371 int32
	_ = v371
	var v376 int32
	_ = v376
	var v383 int32
	_ = v383
	var v390 int32
	_ = v390
	var v395 int32
	_ = v395
	var v402 int32
	_ = v402
	var v409 int32
	_ = v409
	var v414 int32
	_ = v414
	var v421 int32
	_ = v421
	var v428 int32
	_ = v428
	var v433 int32
	_ = v433
	var v440 int32
	_ = v440
	var v447 int32
	_ = v447
	var v452 int32
	_ = v452
	var v459 int32
	_ = v459
	var v466 int32
	_ = v466
	var v471 int32
	_ = v471
	var v478 int32
	_ = v478
	var v490 int64
	_ = v490
	var v492 int32
	_ = v492
	var v496 int32
	_ = v496
	var v499 int64
	_ = v499
	var v500 int32
	_ = v500
	var v502 int32
	_ = v502
	var v507 int32
	_ = v507
	var v509 int32
	_ = v509
	var v510 int32
	_ = v510
	var v511 int64
	_ = v511
	var v513 int32
	_ = v513
	var v517 int32
	_ = v517
	var v518 int32
	_ = v518
	var v521 int64
	_ = v521
	var v522 int64
	_ = v522
	var v523 int32
	_ = v523
	var v524 int32
	_ = v524
	var v526 int32
	_ = v526
	var v529 int32
	_ = v529
	var v530 int64
	_ = v530
	var v531 int32
	_ = v531
	var v538 int32
	_ = v538
	var v539 int64
	_ = v539
	var v540 int32
	_ = v540
	var v544 int32
	_ = v544
	var v547 int64
	_ = v547
	var v549 int32
	_ = v549
	var v552 int64
	_ = v552
	var v554 int32
	_ = v554
	var v557 int64
	_ = v557
	var v559 int32
	_ = v559
	var v562 int64
	_ = v562
	var v563 int64
	_ = v563
	var v567 int32
	_ = v567
	var v570 int64
	_ = v570
	var v576 int32
	_ = v576
	var v579 int32
	_ = v579
	var v583 int32
	_ = v583
	var v584 int32
	_ = v584
	var v590 int32
	_ = v590
	var v591 int32
	_ = v591
	var v594 int32
	_ = v594
	var v608 int32
	_ = v608
	var v611 int64
	_ = v611
	var v613 int32
	_ = v613
	var v617 int32
	_ = v617
	var v618 int32
	_ = v618
	var v621 int64
	_ = v621
	var v622 int64
	_ = v622
	var v624 int32
	_ = v624
	var v647 int32
	_ = v647
	var v648 int64
	_ = v648
	var v649 int32
	_ = v649
	var v654 int32
	_ = v654
	var v655 int64
	_ = v655
	var v656 int32
	_ = v656
	var v660 int32
	_ = v660
	var v661 int32
	_ = v661
	var v663 int32
	_ = v663
	var v665 int32
	_ = v665
	var v666 int32
	_ = v666
	var v668 int32
	_ = v668
	var v672 int32
	_ = v672
	var v673 int32
	_ = v673
	var v676 int32
	_ = v676
	var v678 int32
	_ = v678
	var v683 int32
	_ = v683
	var v690 int32
	_ = v690
	var v697 int32
	_ = v697
	var v702 int32
	_ = v702
	var v709 int32
	_ = v709
	var v716 int32
	_ = v716
	var v721 int32
	_ = v721
	var v728 int32
	_ = v728
	var v735 int32
	_ = v735
	var v740 int32
	_ = v740
	var v747 int32
	_ = v747
	var v754 int32
	_ = v754
	var v759 int32
	_ = v759
	var v766 int32
	_ = v766
	var v773 int32
	_ = v773
	var v778 int32
	_ = v778
	var v785 int32
	_ = v785
	var v797 int64
	_ = v797
	var v799 int32
	_ = v799
	var v803 int32
	_ = v803
	var v806 int64
	_ = v806
	var v807 int32
	_ = v807
	var v809 int32
	_ = v809
	var v814 int32
	_ = v814
	var v816 int32
	_ = v816
	var v817 int32
	_ = v817
	var v818 int64
	_ = v818
	var v820 int32
	_ = v820
	var v824 int32
	_ = v824
	var v825 int32
	_ = v825
	var v828 int64
	_ = v828
	var v829 int64
	_ = v829
	var v830 int32
	_ = v830
	var v831 int32
	_ = v831
	var v833 int32
	_ = v833
	var v836 int32
	_ = v836
	var v837 int64
	_ = v837
	var v838 int32
	_ = v838
	var v845 int32
	_ = v845
	var v846 int64
	_ = v846
	var v847 int32
	_ = v847
	var v851 int32
	_ = v851
	var v854 int64
	_ = v854
	var v856 int32
	_ = v856
	var v859 int64
	_ = v859
	var v861 int32
	_ = v861
	var v864 int64
	_ = v864
	var v866 int32
	_ = v866
	var v869 int64
	_ = v869
	var v870 int64
	_ = v870
	var v874 int32
	_ = v874
	var v877 int64
	_ = v877
	var v883 int64
	_ = v883
	var v885 int32
	_ = v885
	v13 = m.G0
	v15 = v13 - int32(64)
	m.G0 = v15
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	switch v17 + int32(-2) {
	case 0:
		goto L6
	case 1, 2, 3:
		goto L7
	default:
		goto L5
	}
L1:
	;
	m.G0 = v15 + int32(64)
	return
L2:
	;
	if v266 != 0 {
		goto L83
	} else {
		goto L84
	}
L3:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v15)+48)) = v259
	v265 = v255
	v266 = v256
	v267 = v257
	v268 = v258
	v269 = v259
	v271 = v261
	goto L2
L4:
	;
	v232 = int64(3)
	v233 = int64(base.Ui64(v182) >> (uint(v232) % 64))
	*(*int64)(unsafe.Add(mBase, uint32(v15)+56)) = v233
	v237 = int32(-1)
	v241 = int32(7)
	v255 = v237<<(uint((base.I32_wrap_i64(v190)^v237)&v241)%32) ^ v237
	v256 = v140
	v257 = v149
	v258 = v233
	v259 = int64(base.Ui64(v190) >> (uint(v232) % 64))
	v261 = v237 << (uint(int32(8)-base.I32_wrap_i64(v182)&v241) % 32) & int32(65535)
	goto L3
L5:
	;
	v229 = *(*int32)(unsafe.Add(mBase, _consts[33]))
	F_addReplyErrorObject(m, l0, v229)
	mBase = m.M
	v231 = m.ExcPending
	if v231 != 0 {
		goto L8
	} else {
		goto L82
	}
L6:
	;
	v206 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v207 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v208 = *(*int32)(unsafe.Add(mBase, uint32(v207)+4))
	v209 = F_lookupKeyRead(m, v206, v208)
	mBase = m.M
	v210 = m.ExcPending
	if v210 != 0 {
		goto L8
	} else {
		goto L78
	}
L7:
	;
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v21 = *(*int32)(unsafe.Add(mBase, uint32(v20)+8))
	v25 = F_getLongLongFromObjectOrReply(m, l0, v21, v15+int32(56), int32(0))
	mBase = m.M
	v26 = m.ExcPending
	if v26 != 0 {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	return
L9:
	;
	if v25 != 0 {
		goto L1
	} else {
		goto L10
	}
L10:
	;
	v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v28 != int32(5) {
		v126 = int32(1)
		v127 = v28
		goto L11
	} else {
		goto L12
	}
L11:
	;
	if v127 < int32(4) {
		goto L41
	} else {
		goto L42
	}
L12:
	;
	v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v31)+16))
	v33 = F_objectGetVal(m, v32)
	mBase = m.M
	v34 = int32(_a179)
	v37 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v33))))
	if v37 != 0 {
		goto L16
	} else {
		goto L17
	}
L13:
	;
	v125 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v126 = base.B2i32(v72 != int32(0))
	v127 = v125
	goto L11
L14:
	;
	if v72 == int32(0) {
		goto L13
	} else {
		goto L26
	}
L15:
	;
	v69 = F_tolower(m, v65)
	mBase = m.M
	v70 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v66))))
	v71 = F_tolower(m, v70)
	mBase = m.M
	v72 = v69 - v71
	goto L14
L16:
	;
	v39 = v33
	v40 = v34
	v41 = v37
	goto L19
L17:
	;
	v65 = int32(0)
	v66 = v34
	goto L15
L18:
	;
	v65 = v62 & int32(255)
	v66 = v61
	goto L15
L19:
	;
	v43 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v40))))
	if v43 == int32(0) {
		v61 = v40
		v62 = v41
		goto L18
	} else {
		goto L21
	}
L20:
	;
	v61 = v55
	v62 = int32(0)
	goto L18
L21:
	;
	v47 = v41 & int32(255)
	if v47 == v43 {
		goto L22
	} else {
		goto L23
	}
L22:
	;
	v54 = int32(1)
	v55 = v40 + v54
	v56 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v39)+1)))
	if v56 != 0 {
		v39 = v39 + v54
		v40 = v55
		v41 = v56
		goto L19
	} else {
		goto L25
	}
L23:
	;
	v49 = F_tolower(m, v47)
	mBase = m.M
	v50 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v40))))
	v51 = F_tolower(m, v50)
	mBase = m.M
	if v49 == v51 {
		goto L22
	} else {
		goto L24
	}
L24:
	;
	v53 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v39))))
	v61 = v40
	v62 = v53
	goto L18
L25:
	;
	goto L20
L26:
	;
	v75 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v75)+16))
	v77 = F_objectGetVal(m, v76)
	mBase = m.M
	v78 = int32(_a180)
	v81 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v77))))
	if v81 != 0 {
		goto L29
	} else {
		goto L30
	}
L27:
	;
	if v113-v115 == int32(0) {
		goto L13
	} else {
		goto L39
	}
L28:
	;
	v113 = F_tolower(m, v109)
	mBase = m.M
	v114 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v110))))
	v115 = F_tolower(m, v114)
	mBase = m.M
	goto L27
L29:
	;
	v83 = v77
	v84 = v78
	v85 = v81
	goto L32
L30:
	;
	v109 = int32(0)
	v110 = v78
	goto L28
L31:
	;
	v109 = v106 & int32(255)
	v110 = v105
	goto L28
L32:
	;
	v87 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v84))))
	if v87 == int32(0) {
		v105 = v84
		v106 = v85
		goto L31
	} else {
		goto L34
	}
L33:
	;
	v105 = v99
	v106 = int32(0)
	goto L31
L34:
	;
	v91 = v85 & int32(255)
	if v91 == v87 {
		goto L35
	} else {
		goto L36
	}
L35:
	;
	v98 = int32(1)
	v99 = v84 + v98
	v100 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v83)+1)))
	if v100 != 0 {
		v83 = v83 + v98
		v84 = v99
		v85 = v100
		goto L32
	} else {
		goto L38
	}
L36:
	;
	v93 = F_tolower(m, v91)
	mBase = m.M
	v94 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v84))))
	v95 = F_tolower(m, v94)
	mBase = m.M
	if v93 == v95 {
		goto L35
	} else {
		goto L37
	}
L37:
	;
	v97 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v83))))
	v105 = v84
	v106 = v97
	goto L31
L38:
	;
	goto L33
L39:
	;
	v120 = *(*int32)(unsafe.Add(mBase, _consts[33]))
	F_addReplyErrorObject(m, l0, v120)
	mBase = m.M
	v122 = m.ExcPending
	if v122 != 0 {
		goto L8
	} else {
		goto L40
	}
L40:
	;
	goto L1
L41:
	;
	v137 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v138 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v139 = *(*int32)(unsafe.Add(mBase, uint32(v138)+4))
	v140 = F_lookupKeyRead(m, v137, v139)
	mBase = m.M
	v141 = m.ExcPending
	if v141 != 0 {
		goto L8
	} else {
		goto L45
	}
L42:
	;
	v130 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v131 = *(*int32)(unsafe.Add(mBase, uint32(v130)+12))
	v135 = F_getLongLongFromObjectOrReply(m, l0, v131, v15+int32(48), int32(0))
	mBase = m.M
	v136 = m.ExcPending
	if v136 != 0 {
		goto L8
	} else {
		goto L43
	}
L43:
	;
	if v135 != 0 {
		goto L1
	} else {
		goto L44
	}
L44:
	;
	goto L41
L45:
	;
	v143 = F_checkType(m, l0, v140, int32(0))
	mBase = m.M
	v144 = m.ExcPending
	if v144 != 0 {
		goto L8
	} else {
		goto L46
	}
L46:
	;
	if v143 != 0 {
		goto L1
	} else {
		goto L47
	}
L47:
	;
	v149 = F_getObjectReadOnlyString(m, v140, v15+int32(44), v15+int32(16))
	mBase = m.M
	v150 = m.ExcPending
	if v150 != 0 {
		goto L8
	} else {
		goto L48
	}
L48:
	;
	v151 = int64(*(*int32)(unsafe.Add(mBase, uint32(v15)+44)))
	v152 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if int32(3) < v152 {
		goto L50
	} else {
		goto L51
	}
L49:
	;
	v160 = *(*int64)(unsafe.Add(mBase, uint32(v15)+56))
	v162 = base.B2i32(int64(-1) < v160)
	if int64(-1) < v160 {
		goto L53
	} else {
		goto L54
	}
L50:
	;
	v158 = *(*int64)(unsafe.Add(mBase, uint32(v15)+48))
	v159 = v158
	goto L49
L51:
	;
	v156 = v151 + int64(-1)
	*(*int64)(unsafe.Add(mBase, uint32(v15)+48)) = v156
	v159 = v156
	goto L49
L52:
	;
	v202 = *(*int32)(unsafe.Add(mBase, _consts[71]))
	F_addReply(m, l0, v202)
	mBase = m.M
	v204 = m.ExcPending
	if v204 != 0 {
		goto L8
	} else {
		goto L77
	}
L53:
	;
	if v126 != 0 {
		goto L57
	} else {
		goto L58
	}
L54:
	;
	if int64(-1) < v159 {
		goto L53
	} else {
		goto L55
	}
L55:
	;
	if v159 < v160 {
		goto L52
	} else {
		goto L56
	}
L56:
	;
	goto L53
L57:
	;
	v168 = v151
	goto L59
L58:
	;
	v168 = v151 << (uint(int64(3)) % 64)
	goto L59
L59:
	;
	if int64(-1) < v160 {
		v171 = v160
		goto L60
	} else {
		goto L61
	}
L60:
	;
	if int64(-1) < v159 {
		v176 = v159
		goto L62
	} else {
		goto L63
	}
L61:
	;
	v169 = v160 + v168
	*(*int64)(unsafe.Add(mBase, uint32(v15)+56)) = v169
	v171 = v169
	goto L60
L62:
	;
	if int64(-1) < v171 {
		v182 = v171
		goto L64
	} else {
		goto L65
	}
L63:
	;
	v174 = v159 + v168
	*(*int64)(unsafe.Add(mBase, uint32(v15)+48)) = v174
	v176 = v174
	goto L62
L64:
	;
	v185 = int64(0)
	if v185 < v176 {
		goto L66
	} else {
		goto L67
	}
L65:
	;
	v179 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v15)+56)) = v179
	v182 = v179
	goto L64
L66:
	;
	v188 = v176
	goto L68
L67:
	;
	v188 = v185
	goto L68
L68:
	;
	if v168 <= v188 {
		goto L69
	} else {
		goto L70
	}
L69:
	;
	v190 = v168 + int64(-1)
	goto L71
L70:
	;
	v190 = v188
	goto L71
L71:
	;
	if v176 < int64(0) {
		goto L73
	} else {
		goto L74
	}
L72:
	;
	v195 = int32(0)
	if v126|base.B2i32(v190 < v182) == v195 {
		goto L4
	} else {
		goto L76
	}
L73:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v15)+48)) = v190
	goto L72
L74:
	;
	if v188 < v168 {
		goto L72
	} else {
		goto L75
	}
L75:
	;
	goto L73
L76:
	;
	v265 = int32(0)
	v266 = v140
	v267 = v149
	v268 = v182
	v269 = v190
	v271 = v195
	goto L2
L77:
	;
	goto L1
L78:
	;
	v212 = F_checkType(m, l0, v209, int32(0))
	mBase = m.M
	v213 = m.ExcPending
	if v213 != 0 {
		goto L8
	} else {
		goto L79
	}
L79:
	;
	if v212 != 0 {
		goto L1
	} else {
		goto L80
	}
L80:
	;
	v219 = F_getObjectReadOnlyString(m, v209, v15+int32(44), v15+int32(16))
	mBase = m.M
	v220 = m.ExcPending
	if v220 != 0 {
		goto L8
	} else {
		goto L81
	}
L81:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v15)+56)) = int64(0)
	v223 = *(*int32)(unsafe.Add(mBase, uint32(v15)+44))
	v255 = int32(0)
	v256 = v209
	v257 = v219
	v258 = int64(0)
	v259 = base.I64_extend_i32_s(v223 + int32(-1))
	v261 = int32(0)
	goto L3
L82:
	;
	goto L1
L83:
	;
	if v268 <= v269 {
		goto L86
	} else {
		goto L87
	}
L84:
	;
	v275 = *(*int32)(unsafe.Add(mBase, _consts[71]))
	F_addReply(m, l0, v275)
	mBase = m.M
	v277 = m.ExcPending
	if v277 != 0 {
		goto L8
	} else {
		goto L85
	}
L85:
	;
	goto L1
L86:
	;
	v284 = v267 + base.I32_wrap_i64(v268)
	v286 = base.I32_wrap_i64(v269 - v268)
	v288 = v286 + int32(1)
	if v284&int32(3) == int32(0) {
		goto L92
	} else {
		goto L93
	}
L87:
	;
	v280 = *(*int32)(unsafe.Add(mBase, _consts[71]))
	F_addReply(m, l0, v280)
	mBase = m.M
	v282 = m.ExcPending
	if v282 != 0 {
		goto L8
	} else {
		goto L88
	}
L88:
	;
	goto L1
L89:
	;
	v576 = v271 & int32(255)
	if v576 != 0 {
		goto L126
	} else {
		goto L127
	}
L90:
	;
	if int32(27) < v340 {
		goto L108
	} else {
		goto L109
	}
L91:
	;
	v300 = v286 + int32(0)
	v301 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v284))))
	v304 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v301)+uint32(_consts[72]))))
	v306 = v284 + int32(1)
	if v306&int32(3) != 0 {
		goto L95
	} else {
		goto L96
	}
L92:
	;
	v340 = v288
	v341 = int64(0)
	v342 = v284
	goto L90
L93:
	;
	if v288 != 0 {
		goto L91
	} else {
		goto L94
	}
L94:
	;
	goto L92
L95:
	;
	if v300 != 0 {
		goto L97
	} else {
		goto L98
	}
L96:
	;
	v340 = v300
	v341 = v304
	v342 = v306
	goto L90
L97:
	;
	v310 = v286 + int32(-1)
	v311 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v284)+1)))
	v314 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v311)+uint32(_consts[72]))))
	v315 = v304 + v314
	v317 = v284 + int32(2)
	if v317&int32(3) != 0 {
		goto L99
	} else {
		goto L100
	}
L98:
	;
	v340 = v300
	v341 = v304
	v342 = v306
	goto L90
L99:
	;
	if v310 != 0 {
		goto L101
	} else {
		goto L102
	}
L100:
	;
	v340 = v310
	v341 = v315
	v342 = v317
	goto L90
L101:
	;
	v321 = v286 + int32(-2)
	v322 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v284)+2)))
	v325 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v322)+uint32(_consts[72]))))
	v326 = v315 + v325
	v327 = int32(3)
	v328 = v284 + v327
	if v328&v327 != 0 {
		goto L103
	} else {
		goto L104
	}
L102:
	;
	v340 = v310
	v341 = v315
	v342 = v317
	goto L90
L103:
	;
	if v321 != 0 {
		goto L105
	} else {
		goto L106
	}
L104:
	;
	v340 = v321
	v341 = v326
	v342 = v328
	goto L90
L105:
	;
	v335 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v284)+3)))
	v338 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v335)+uint32(_consts[72]))))
	v340 = v286 + int32(-3)
	v341 = v326 + v338
	v342 = v284 + int32(4)
	goto L90
L106:
	;
	v340 = v321
	v341 = v326
	v342 = v328
	goto L90
L107:
	;
	if v502 == int32(0) {
		v570 = v499
		goto L113
	} else {
		goto L114
	}
L108:
	;
	v347 = v340
	v348 = v341
	v349 = v342
	goto L110
L109:
	;
	v499 = v341
	v500 = v342
	v502 = v340
	goto L107
L110:
	;
	v353 = *(*int32)(unsafe.Add(mBase, uint32(v349)+4))
	v354 = int32(1)
	v356 = int32(1431655765)
	v358 = v353 - int32(base.Ui32(v353)>>(uint(v354)%32))&v356
	v359 = int32(2)
	v361 = int32(858993459)
	v365 = int32(base.Ui32(v358)>>(uint(v359)%32))&v361 + v358&v361
	v366 = int32(4)
	v369 = int32(252645135)
	v371 = *(*int32)(unsafe.Add(mBase, uint32(v349)))
	v376 = v371 - int32(base.Ui32(v371)>>(uint(v354)%32))&v356
	v383 = int32(base.Ui32(v376)>>(uint(v359)%32))&v361 + v376&v361
	v390 = *(*int32)(unsafe.Add(mBase, uint32(v349)+8))
	v395 = v390 - int32(base.Ui32(v390)>>(uint(v354)%32))&v356
	v402 = int32(base.Ui32(v395)>>(uint(v359)%32))&v361 + v395&v361
	v409 = *(*int32)(unsafe.Add(mBase, uint32(v349)+12))
	v414 = v409 - int32(base.Ui32(v409)>>(uint(v354)%32))&v356
	v421 = int32(base.Ui32(v414)>>(uint(v359)%32))&v361 + v414&v361
	v428 = *(*int32)(unsafe.Add(mBase, uint32(v349)+16))
	v433 = v428 - int32(base.Ui32(v428)>>(uint(v354)%32))&v356
	v440 = int32(base.Ui32(v433)>>(uint(v359)%32))&v361 + v433&v361
	v447 = *(*int32)(unsafe.Add(mBase, uint32(v349)+20))
	v452 = v447 - int32(base.Ui32(v447)>>(uint(v354)%32))&v356
	v459 = int32(base.Ui32(v452)>>(uint(v359)%32))&v361 + v452&v361
	v466 = *(*int32)(unsafe.Add(mBase, uint32(v349)+24))
	v471 = v466 - int32(base.Ui32(v466)>>(uint(v354)%32))&v356
	v478 = int32(base.Ui32(v471)>>(uint(v359)%32))&v361 + v471&v361
	v490 = v348 + base.I64_extend_i32_u(int32(base.Ui32(((int32(base.Ui32(v365)>>(uint(v366)%32))+v365)&v369+(int32(base.Ui32(v383)>>(uint(v366)%32))+v383)&v369+(int32(base.Ui32(v402)>>(uint(v366)%32))+v402)&v369+(int32(base.Ui32(v421)>>(uint(v366)%32))+v421)&v369+(int32(base.Ui32(v440)>>(uint(v366)%32))+v440)&v369+(int32(base.Ui32(v459)>>(uint(v366)%32))+v459)&v369+(int32(base.Ui32(v478)>>(uint(v366)%32))+v478)&v369)*int32(16843009))>>(uint(int32(24))%32)))
	v492 = v349 + int32(28)
	v496 = v347 + int32(-28)
	if base.Ui32(int32(55)) < base.Ui32(v347) {
		v347 = v496
		v348 = v490
		v349 = v492
		goto L110
	} else {
		goto L112
	}
L111:
	;
	v499 = v490
	v500 = v492
	v502 = v496
	goto L107
L112:
	;
	goto L111
L113:
	;
	goto L89
L114:
	;
	v507 = v502 & int32(3)
	if v507 != 0 {
		goto L116
	} else {
		goto L117
	}
L115:
	;
	if base.Ui32(v502) < base.Ui32(int32(4)) {
		v570 = v530
		goto L113
	} else {
		goto L121
	}
L116:
	;
	v509 = v500
	v510 = v502
	v511 = v499
	v513 = int32(0)
	goto L118
L117:
	;
	v529 = v502
	v530 = v499
	v531 = v500
	goto L115
L118:
	;
	v517 = v510 + int32(-1)
	v518 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v509))))
	v521 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v518)+uint32(_consts[72]))))
	v522 = v511 + v521
	v523 = int32(1)
	v524 = v509 + v523
	v526 = v513 + v523
	if v526 != v507 {
		v509 = v524
		v510 = v517
		v511 = v522
		v513 = v526
		goto L118
	} else {
		goto L120
	}
L119:
	;
	v529 = v517
	v530 = v522
	v531 = v524
	goto L115
L120:
	;
	goto L119
L121:
	;
	v538 = v529
	v539 = v530
	v540 = v531
	goto L122
L122:
	;
	v544 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v540))))
	v547 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v544)+uint32(_consts[72]))))
	v549 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v540)+1)))
	v552 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v549)+uint32(_consts[72]))))
	v554 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v540)+2)))
	v557 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v554)+uint32(_consts[72]))))
	v559 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v540)+3)))
	v562 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v559)+uint32(_consts[72]))))
	v563 = v539 + v547 + v552 + v557 + v562
	v567 = v538 + int32(-4)
	if v567 != 0 {
		v538 = v567
		v539 = v563
		v540 = v540 + int32(4)
		goto L122
	} else {
		goto L124
	}
L123:
	;
	v570 = v563
	goto L113
L124:
	;
	goto L123
L125:
	;
	F_addReplyLongLong(m, l0, v883)
	mBase = m.M
	v885 = m.ExcPending
	if v885 != 0 {
		goto L8
	} else {
		goto L169
	}
L126:
	;
	v579 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v15)+14)) = uint16(v579)
	if v576 == v579 {
		goto L129
	} else {
		goto L130
	}
L127:
	;
	if v265 == int32(0) {
		v883 = v570
		goto L125
	} else {
		goto L128
	}
L128:
	;
	goto L126
L129:
	;
	if v265 == int32(0) {
		goto L131
	} else {
		goto L132
	}
L130:
	;
	v583 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v284))))
	v584 = v583 & v271
	*(*uint8)(unsafe.Add(mBase, uint32(v15)+14)) = uint8(v584)
	goto L129
L131:
	;
	v594 = v15 + int32(14)
	if v594&int32(3) == int32(0) {
		goto L136
	} else {
		goto L137
	}
L132:
	;
	v590 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v267+base.I32_wrap_i64(v269)))))
	v591 = v590 & v265
	*(*uint8)(unsafe.Add(mBase, uint32(v15)+15)) = uint8(v591)
	goto L131
L133:
	;
	v883 = v570 - v877
	goto L125
L134:
	;
	if int32(27) < v647 {
		goto L152
	} else {
		goto L153
	}
L135:
	;
	v608 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v594))))
	v611 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v608)+uint32(_consts[72]))))
	v613 = v15 + int32(15)
	if v613&int32(3) != 0 {
		goto L139
	} else {
		goto L140
	}
L136:
	;
	v647 = int32(2)
	v648 = int64(0)
	v649 = v594
	goto L134
L137:
	;
	goto L135
L139:
	;
	goto L141
L140:
	;
	v647 = int32(1)
	v648 = v611
	v649 = v613
	goto L134
L141:
	;
	v617 = int32(0)
	v618 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v594)+1)))
	v621 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v618)+uint32(_consts[72]))))
	v622 = v611 + v621
	v624 = v15 + int32(16)
	if v624&int32(3) != 0 {
		goto L143
	} else {
		goto L144
	}
L143:
	;
	goto L146
L144:
	;
	v647 = v617
	v648 = v622
	v649 = v624
	goto L134
L146:
	;
	v647 = v617
	v648 = v622
	v649 = v624
	goto L134
L151:
	;
	if v809 == int32(0) {
		v877 = v806
		goto L157
	} else {
		goto L158
	}
L152:
	;
	v654 = v647
	v655 = v648
	v656 = v649
	goto L154
L153:
	;
	v806 = v648
	v807 = v649
	v809 = v647
	goto L151
L154:
	;
	v660 = *(*int32)(unsafe.Add(mBase, uint32(v656)+4))
	v661 = int32(1)
	v663 = int32(1431655765)
	v665 = v660 - int32(base.Ui32(v660)>>(uint(v661)%32))&v663
	v666 = int32(2)
	v668 = int32(858993459)
	v672 = int32(base.Ui32(v665)>>(uint(v666)%32))&v668 + v665&v668
	v673 = int32(4)
	v676 = int32(252645135)
	v678 = *(*int32)(unsafe.Add(mBase, uint32(v656)))
	v683 = v678 - int32(base.Ui32(v678)>>(uint(v661)%32))&v663
	v690 = int32(base.Ui32(v683)>>(uint(v666)%32))&v668 + v683&v668
	v697 = *(*int32)(unsafe.Add(mBase, uint32(v656)+8))
	v702 = v697 - int32(base.Ui32(v697)>>(uint(v661)%32))&v663
	v709 = int32(base.Ui32(v702)>>(uint(v666)%32))&v668 + v702&v668
	v716 = *(*int32)(unsafe.Add(mBase, uint32(v656)+12))
	v721 = v716 - int32(base.Ui32(v716)>>(uint(v661)%32))&v663
	v728 = int32(base.Ui32(v721)>>(uint(v666)%32))&v668 + v721&v668
	v735 = *(*int32)(unsafe.Add(mBase, uint32(v656)+16))
	v740 = v735 - int32(base.Ui32(v735)>>(uint(v661)%32))&v663
	v747 = int32(base.Ui32(v740)>>(uint(v666)%32))&v668 + v740&v668
	v754 = *(*int32)(unsafe.Add(mBase, uint32(v656)+20))
	v759 = v754 - int32(base.Ui32(v754)>>(uint(v661)%32))&v663
	v766 = int32(base.Ui32(v759)>>(uint(v666)%32))&v668 + v759&v668
	v773 = *(*int32)(unsafe.Add(mBase, uint32(v656)+24))
	v778 = v773 - int32(base.Ui32(v773)>>(uint(v661)%32))&v663
	v785 = int32(base.Ui32(v778)>>(uint(v666)%32))&v668 + v778&v668
	v797 = v655 + base.I64_extend_i32_u(int32(base.Ui32(((int32(base.Ui32(v672)>>(uint(v673)%32))+v672)&v676+(int32(base.Ui32(v690)>>(uint(v673)%32))+v690)&v676+(int32(base.Ui32(v709)>>(uint(v673)%32))+v709)&v676+(int32(base.Ui32(v728)>>(uint(v673)%32))+v728)&v676+(int32(base.Ui32(v747)>>(uint(v673)%32))+v747)&v676+(int32(base.Ui32(v766)>>(uint(v673)%32))+v766)&v676+(int32(base.Ui32(v785)>>(uint(v673)%32))+v785)&v676)*int32(16843009))>>(uint(int32(24))%32)))
	v799 = v656 + int32(28)
	v803 = v654 + int32(-28)
	if base.Ui32(int32(55)) < base.Ui32(v654) {
		v654 = v803
		v655 = v797
		v656 = v799
		goto L154
	} else {
		goto L156
	}
L155:
	;
	v806 = v797
	v807 = v799
	v809 = v803
	goto L151
L156:
	;
	goto L155
L157:
	;
	goto L133
L158:
	;
	v814 = v809 & int32(3)
	if v814 != 0 {
		goto L160
	} else {
		goto L161
	}
L159:
	;
	if base.Ui32(v809) < base.Ui32(int32(4)) {
		v877 = v837
		goto L157
	} else {
		goto L165
	}
L160:
	;
	v816 = v807
	v817 = v809
	v818 = v806
	v820 = int32(0)
	goto L162
L161:
	;
	v836 = v809
	v837 = v806
	v838 = v807
	goto L159
L162:
	;
	v824 = v817 + int32(-1)
	v825 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v816))))
	v828 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v825)+uint32(_consts[72]))))
	v829 = v818 + v828
	v830 = int32(1)
	v831 = v816 + v830
	v833 = v820 + v830
	if v833 != v814 {
		v816 = v831
		v817 = v824
		v818 = v829
		v820 = v833
		goto L162
	} else {
		goto L164
	}
L163:
	;
	v836 = v824
	v837 = v829
	v838 = v831
	goto L159
L164:
	;
	goto L163
L165:
	;
	v845 = v836
	v846 = v837
	v847 = v838
	goto L166
L166:
	;
	v851 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v847))))
	v854 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v851)+uint32(_consts[72]))))
	v856 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v847)+1)))
	v859 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v856)+uint32(_consts[72]))))
	v861 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v847)+2)))
	v864 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v861)+uint32(_consts[72]))))
	v866 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v847)+3)))
	v869 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v866)+uint32(_consts[72]))))
	v870 = v846 + v854 + v859 + v864 + v869
	v874 = v845 + int32(-4)
	if v874 != 0 {
		v845 = v874
		v846 = v870
		v847 = v847 + int32(4)
		goto L166
	} else {
		goto L168
	}
L167:
	;
	v877 = v870
	goto L157
L168:
	;
	goto L167
L169:
	;
	goto L1
}
func F_bitposCommand(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
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
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
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
	var v107 int32
	_ = v107
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
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
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v136 int32
	_ = v136
	var v138 int32
	_ = v138
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
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
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v168 int64
	_ = v168
	var v169 int32
	_ = v169
	var v177 int64
	_ = v177
	var v178 int64
	_ = v178
	var v181 int64
	_ = v181
	var v183 int64
	_ = v183
	var v184 int64
	_ = v184
	var v187 int64
	_ = v187
	var v189 int64
	_ = v189
	var v192 int64
	_ = v192
	var v195 int64
	_ = v195
	var v198 int64
	_ = v198
	var v201 int64
	_ = v201
	var v203 int64
	_ = v203
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v234 int32
	_ = v234
	var v237 int64
	_ = v237
	var v239 int32
	_ = v239
	var v242 int32
	_ = v242
	var v244 int32
	_ = v244
	var v245 int64
	_ = v245
	var v246 int64
	_ = v246
	var v249 int64
	_ = v249
	var v251 int32
	_ = v251
	var v254 int32
	_ = v254
	var v267 int32
	_ = v267
	var v268 int32
	_ = v268
	var v269 int32
	_ = v269
	var v270 int32
	_ = v270
	var v271 int64
	_ = v271
	var v272 int64
	_ = v272
	var v276 int32
	_ = v276
	var v279 int32
	_ = v279
	var v280 int64
	_ = v280
	var v282 int32
	_ = v282
	var v286 int32
	_ = v286
	var v288 int32
	_ = v288
	var v290 int32
	_ = v290
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
	var v310 int32
	_ = v310
	var v312 int32
	_ = v312
	var v313 int32
	_ = v313
	var v320 int32
	_ = v320
	var v325 int64
	_ = v325
	var v326 int32
	_ = v326
	var v331 int64
	_ = v331
	var v335 int64
	_ = v335
	var v337 int64
	_ = v337
	var v339 int32
	_ = v339
	var v340 int64
	_ = v340
	var v347 int32
	_ = v347
	var v352 int32
	_ = v352
	var v353 int64
	_ = v353
	var v354 int32
	_ = v354
	var v357 int64
	_ = v357
	var v363 int64
	_ = v363
	var v367 int32
	_ = v367
	var v370 int32
	_ = v370
	var v372 int32
	_ = v372
	var v377 int32
	_ = v377
	var v378 int32
	_ = v378
	var v379 int32
	_ = v379
	var v380 int32
	_ = v380
	var v385 int32
	_ = v385
	var v390 int64
	_ = v390
	var v391 int32
	_ = v391
	var v393 int32
	_ = v393
	var v394 int64
	_ = v394
	var v399 int32
	_ = v399
	var v411 int32
	_ = v411
	var v412 int64
	_ = v412
	var v413 int64
	_ = v413
	var v419 int64
	_ = v419
	var v421 int32
	_ = v421
	v16 = m.G0
	v18 = v16 - int32(64)
	m.G0 = v18
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v21 = *(*int32)(unsafe.Add(mBase, uint32(v20)+8))
	v25 = F_getLongFromObjectOrReply(m, l0, v21, v18+int32(44), int32(0))
	mBase = m.M
	v26 = m.ExcPending
	if v26 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v18 + int32(64)
	return
L2:
	;
	return
L3:
	;
	if v25 != 0 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v18)+44))
	if base.Ui32(v27) < base.Ui32(int32(2)) {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	v33 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	switch v33 + int32(-3) {
	case 0:
		goto L11
	case 1, 2, 3:
		goto L12
	default:
		goto L10
	}
L6:
	;
	F_addReplyError(m, l0, int32(_a181))
	mBase = m.M
	v32 = m.ExcPending
	if v32 != 0 {
		goto L2
	} else {
		goto L7
	}
L7:
	;
	goto L1
L8:
	;
	if v269 != 0 {
		goto L80
	} else {
		goto L81
	}
L9:
	;
	v245 = int64(3)
	v246 = int64(base.Ui64(v203) >> (uint(v245) % 64))
	*(*int64)(unsafe.Add(mBase, uint32(v18)+48)) = v246
	v249 = int64(base.Ui64(v195) >> (uint(v245) % 64))
	*(*int64)(unsafe.Add(mBase, uint32(v18)+56)) = v249
	v251 = int32(-1)
	v254 = int32(7)
	v267 = v251<<(uint((base.I32_wrap_i64(v203)^v251)&v254)%32) ^ v251
	v268 = v209
	v269 = v157
	v270 = v166
	v271 = v249
	v272 = v246
	v276 = v251 << (uint(int32(8)-base.I32_wrap_i64(v195)&v254) % 32)
	goto L8
L10:
	;
	v242 = *(*int32)(unsafe.Add(mBase, _consts[33]))
	F_addReplyErrorObject(m, l0, v242)
	mBase = m.M
	v244 = m.ExcPending
	if v244 != 0 {
		goto L2
	} else {
		goto L79
	}
L11:
	;
	v217 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v218 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v219 = *(*int32)(unsafe.Add(mBase, uint32(v218)+4))
	v220 = F_lookupKeyRead(m, v217, v219)
	mBase = m.M
	v221 = m.ExcPending
	if v221 != 0 {
		goto L2
	} else {
		goto L75
	}
L12:
	;
	v36 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v36)+12))
	v41 = F_getLongLongFromObjectOrReply(m, l0, v37, v18+int32(56), int32(0))
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L2
	} else {
		goto L13
	}
L13:
	;
	if v41 != 0 {
		goto L1
	} else {
		goto L14
	}
L14:
	;
	v44 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v44 != int32(6) {
		v142 = int32(1)
		v143 = v44
		goto L15
	} else {
		goto L16
	}
L15:
	;
	if v143 < int32(5) {
		goto L45
	} else {
		goto L46
	}
L16:
	;
	v47 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v47)+20))
	v49 = F_objectGetVal(m, v48)
	mBase = m.M
	v50 = int32(_a179)
	v53 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v49))))
	if v53 != 0 {
		goto L20
	} else {
		goto L21
	}
L17:
	;
	v141 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v142 = base.B2i32(v88 != int32(0))
	v143 = v141
	goto L15
L18:
	;
	if v88 == int32(0) {
		goto L17
	} else {
		goto L30
	}
L19:
	;
	v85 = F_tolower(m, v81)
	mBase = m.M
	v86 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v82))))
	v87 = F_tolower(m, v86)
	mBase = m.M
	v88 = v85 - v87
	goto L18
L20:
	;
	v55 = v49
	v56 = v50
	v57 = v53
	goto L23
L21:
	;
	v81 = int32(0)
	v82 = v50
	goto L19
L22:
	;
	v81 = v78 & int32(255)
	v82 = v77
	goto L19
L23:
	;
	v59 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v56))))
	if v59 == int32(0) {
		v77 = v56
		v78 = v57
		goto L22
	} else {
		goto L25
	}
L24:
	;
	v77 = v71
	v78 = int32(0)
	goto L22
L25:
	;
	v63 = v57 & int32(255)
	if v63 == v59 {
		goto L26
	} else {
		goto L27
	}
L26:
	;
	v70 = int32(1)
	v71 = v56 + v70
	v72 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v55)+1)))
	if v72 != 0 {
		v55 = v55 + v70
		v56 = v71
		v57 = v72
		goto L23
	} else {
		goto L29
	}
L27:
	;
	v65 = F_tolower(m, v63)
	mBase = m.M
	v66 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v56))))
	v67 = F_tolower(m, v66)
	mBase = m.M
	if v65 == v67 {
		goto L26
	} else {
		goto L28
	}
L28:
	;
	v69 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v55))))
	v77 = v56
	v78 = v69
	goto L22
L29:
	;
	goto L24
L30:
	;
	v91 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v92 = *(*int32)(unsafe.Add(mBase, uint32(v91)+20))
	v93 = F_objectGetVal(m, v92)
	mBase = m.M
	v94 = int32(_a180)
	v97 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v93))))
	if v97 != 0 {
		goto L33
	} else {
		goto L34
	}
L31:
	;
	if v129-v131 == int32(0) {
		goto L17
	} else {
		goto L43
	}
L32:
	;
	v129 = F_tolower(m, v125)
	mBase = m.M
	v130 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v126))))
	v131 = F_tolower(m, v130)
	mBase = m.M
	goto L31
L33:
	;
	v99 = v93
	v100 = v94
	v101 = v97
	goto L36
L34:
	;
	v125 = int32(0)
	v126 = v94
	goto L32
L35:
	;
	v125 = v122 & int32(255)
	v126 = v121
	goto L32
L36:
	;
	v103 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v100))))
	if v103 == int32(0) {
		v121 = v100
		v122 = v101
		goto L35
	} else {
		goto L38
	}
L37:
	;
	v121 = v115
	v122 = int32(0)
	goto L35
L38:
	;
	v107 = v101 & int32(255)
	if v107 == v103 {
		goto L39
	} else {
		goto L40
	}
L39:
	;
	v114 = int32(1)
	v115 = v100 + v114
	v116 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v99)+1)))
	if v116 != 0 {
		v99 = v99 + v114
		v100 = v115
		v101 = v116
		goto L36
	} else {
		goto L42
	}
L40:
	;
	v109 = F_tolower(m, v107)
	mBase = m.M
	v110 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v100))))
	v111 = F_tolower(m, v110)
	mBase = m.M
	if v109 == v111 {
		goto L39
	} else {
		goto L41
	}
L41:
	;
	v113 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v99))))
	v121 = v100
	v122 = v113
	goto L35
L42:
	;
	goto L37
L43:
	;
	v136 = *(*int32)(unsafe.Add(mBase, _consts[33]))
	F_addReplyErrorObject(m, l0, v136)
	mBase = m.M
	v138 = m.ExcPending
	if v138 != 0 {
		goto L2
	} else {
		goto L44
	}
L44:
	;
	goto L1
L45:
	;
	v154 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v155 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v156 = *(*int32)(unsafe.Add(mBase, uint32(v155)+4))
	v157 = F_lookupKeyRead(m, v154, v156)
	mBase = m.M
	v158 = m.ExcPending
	if v158 != 0 {
		goto L2
	} else {
		goto L49
	}
L46:
	;
	v147 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v148 = *(*int32)(unsafe.Add(mBase, uint32(v147)+16))
	v152 = F_getLongLongFromObjectOrReply(m, l0, v148, v18+int32(48), int32(0))
	mBase = m.M
	v153 = m.ExcPending
	if v153 != 0 {
		goto L2
	} else {
		goto L47
	}
L47:
	;
	if v152 != 0 {
		goto L1
	} else {
		goto L48
	}
L48:
	;
	goto L45
L49:
	;
	v160 = F_checkType(m, l0, v157, int32(0))
	mBase = m.M
	v161 = m.ExcPending
	if v161 != 0 {
		goto L2
	} else {
		goto L50
	}
L50:
	;
	if v160 != 0 {
		goto L1
	} else {
		goto L51
	}
L51:
	;
	v166 = F_getObjectReadOnlyString(m, v157, v18+int32(40), v18+int32(16))
	mBase = m.M
	v167 = m.ExcPending
	if v167 != 0 {
		goto L2
	} else {
		goto L52
	}
L52:
	;
	v168 = int64(*(*int32)(unsafe.Add(mBase, uint32(v18)+40)))
	v169 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if int32(4) < v169 {
		goto L53
	} else {
		goto L54
	}
L53:
	;
	if v142 != 0 {
		goto L55
	} else {
		goto L56
	}
L54:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v18)+48)) = v168 + int64(-1)
	goto L53
L55:
	;
	v177 = v168
	goto L57
L56:
	;
	v177 = v168 << (uint(int64(3)) % 64)
	goto L57
L57:
	;
	v178 = *(*int64)(unsafe.Add(mBase, uint32(v18)+56))
	if int64(-1) < v178 {
		v183 = v178
		goto L58
	} else {
		goto L59
	}
L58:
	;
	v184 = *(*int64)(unsafe.Add(mBase, uint32(v18)+48))
	if int64(-1) < v184 {
		v189 = v184
		goto L60
	} else {
		goto L61
	}
L59:
	;
	v181 = v178 + v177
	*(*int64)(unsafe.Add(mBase, uint32(v18)+56)) = v181
	v183 = v181
	goto L58
L60:
	;
	if int64(-1) < v183 {
		v195 = v183
		goto L62
	} else {
		goto L63
	}
L61:
	;
	v187 = v184 + v177
	*(*int64)(unsafe.Add(mBase, uint32(v18)+48)) = v187
	v189 = v187
	goto L60
L62:
	;
	v198 = int64(0)
	if v198 < v189 {
		goto L64
	} else {
		goto L65
	}
L63:
	;
	v192 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v18)+56)) = v192
	v195 = v192
	goto L62
L64:
	;
	v201 = v189
	goto L66
L65:
	;
	v201 = v198
	goto L66
L66:
	;
	if v177 <= v201 {
		goto L67
	} else {
		goto L68
	}
L67:
	;
	v203 = v177 + int64(-1)
	goto L69
L68:
	;
	v203 = v201
	goto L69
L69:
	;
	if v189 < int64(0) {
		goto L71
	} else {
		goto L72
	}
L70:
	;
	v209 = base.B2i32(int32(4) < v143)
	v210 = int32(0)
	if v142|base.B2i32(v203 < v195) == v210 {
		goto L9
	} else {
		goto L74
	}
L71:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v18)+48)) = v203
	goto L70
L72:
	;
	if v201 < v177 {
		goto L70
	} else {
		goto L73
	}
L73:
	;
	goto L71
L74:
	;
	v267 = int32(0)
	v268 = v209
	v269 = v157
	v270 = v166
	v271 = v195
	v272 = v203
	v276 = v210
	goto L8
L75:
	;
	v223 = F_checkType(m, l0, v220, int32(0))
	mBase = m.M
	v224 = m.ExcPending
	if v224 != 0 {
		goto L2
	} else {
		goto L76
	}
L76:
	;
	if v223 != 0 {
		goto L1
	} else {
		goto L77
	}
L77:
	;
	v230 = F_getObjectReadOnlyString(m, v220, v18+int32(40), v18+int32(16))
	mBase = m.M
	v231 = m.ExcPending
	if v231 != 0 {
		goto L2
	} else {
		goto L78
	}
L78:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v18)+56)) = int64(0)
	v234 = *(*int32)(unsafe.Add(mBase, uint32(v18)+40))
	v237 = base.I64_extend_i32_s(v234 + int32(-1))
	*(*int64)(unsafe.Add(mBase, uint32(v18)+48)) = v237
	v239 = int32(0)
	v267 = v239
	v268 = int32(0)
	v269 = v220
	v270 = v230
	v271 = int64(0)
	v272 = v237
	v276 = v239
	goto L8
L79:
	;
	goto L1
L80:
	;
	if v271 <= v272 {
		goto L86
	} else {
		goto L87
	}
L81:
	;
	v279 = *(*int32)(unsafe.Add(mBase, uint32(v18)+44))
	if v279 != 0 {
		goto L82
	} else {
		goto L83
	}
L82:
	;
	v280 = int64(-1)
	goto L84
L83:
	;
	v280 = int64(0)
	goto L84
L84:
	;
	F_addReplyLongLong(m, l0, v280)
	mBase = m.M
	v282 = m.ExcPending
	if v282 != 0 {
		goto L2
	} else {
		goto L85
	}
L85:
	;
	goto L1
L86:
	;
	v288 = base.I32_wrap_i64(v272 - v271)
	v290 = v288 + int32(1)
	if v276&int32(255) == int32(0) {
		v339 = v290
		v340 = v271
		goto L91
	} else {
		goto L92
	}
L87:
	;
	F_addReplyLongLong(m, l0, int64(-1))
	mBase = m.M
	v286 = m.ExcPending
	if v286 != 0 {
		goto L2
	} else {
		goto L88
	}
L88:
	;
	goto L1
L89:
	;
	v399 = *(*int32)(unsafe.Add(mBase, uint32(v18)+44))
	if v268&base.B2i32(v399 == int32(0)) != int32(1) {
		goto L116
	} else {
		goto L117
	}
L90:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v18)+15)) = uint8(v379)
	v390 = F_serverBitpos(m, v18+int32(15), int32(1), v385)
	mBase = m.M
	v391 = m.ExcPending
	if v391 != 0 {
		goto L2
	} else {
		goto L115
	}
L91:
	;
	v347 = v339 - base.B2i32(v267 != int32(0))
	if v347 < int32(1) {
		v367 = v339
		goto L105
	} else {
		goto L106
	}
L92:
	;
	v297 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v270+base.I32_wrap_i64(v271)))))
	v298 = *(*int32)(unsafe.Add(mBase, uint32(v18)+44))
	if v298 != 0 {
		goto L94
	} else {
		goto L95
	}
L93:
	;
	v325 = F_serverBitpos(m, v18+int32(15), int32(1), v298)
	mBase = m.M
	v326 = m.ExcPending
	if v326 != 0 {
		goto L2
	} else {
		goto L100
	}
L94:
	;
	v310 = v297 & (v276 ^ int32(-1))
	*(*uint8)(unsafe.Add(mBase, uint32(v18)+15)) = uint8(v310)
	v312 = int32(0)
	v313 = base.B2i32(v288 == v312)
	if v267 == v312 {
		v320 = v313
		goto L93
	} else {
		goto L98
	}
L95:
	;
	v299 = v297 | v276
	*(*uint8)(unsafe.Add(mBase, uint32(v18)+15)) = uint8(v299)
	v301 = int32(0)
	v302 = base.B2i32(v288 == v301)
	if v267 == v301 {
		v320 = v302
		goto L93
	} else {
		goto L96
	}
L96:
	;
	if v288 != 0 {
		v320 = v302
		goto L93
	} else {
		goto L97
	}
L97:
	;
	v379 = v299 | v267
	v380 = int32(1)
	v385 = int32(0)
	goto L90
L98:
	;
	if v288 != 0 {
		v320 = v313
		goto L93
	} else {
		goto L99
	}
L99:
	;
	v379 = v310 & (v267 ^ int32(-1))
	v380 = int32(1)
	v385 = v298
	goto L90
L100:
	;
	if v320 == int32(0) {
		goto L101
	} else {
		goto L102
	}
L101:
	;
	v331 = v325 + int64(1)
	if base.Ui64(int64(9)) < base.Ui64(v331) {
		v393 = v290
		v394 = v325
		goto L89
	} else {
		goto L103
	}
L102:
	;
	v393 = int32(1)
	v394 = v325
	goto L89
L103:
	;
	switch base.I32_wrap_i64(v331) {
	default:
		goto L104
	case 1, 2, 3, 4, 5, 6, 7, 8:
		v393 = v290
		v394 = v325
		goto L89
	}
L104:
	;
	v335 = *(*int64)(unsafe.Add(mBase, uint32(v18)+56))
	v337 = v335 + int64(1)
	*(*int64)(unsafe.Add(mBase, uint32(v18)+56)) = v337
	v339 = v288
	v340 = v337
	goto L91
L105:
	;
	v370 = *(*int32)(unsafe.Add(mBase, uint32(v18)+48))
	v372 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v270+v370))))
	v377 = *(*int32)(unsafe.Add(mBase, uint32(v18)+44))
	if v377 != 0 {
		goto L112
	} else {
		goto L113
	}
L106:
	;
	v352 = *(*int32)(unsafe.Add(mBase, uint32(v18)+44))
	v353 = F_serverBitpos(m, v270+base.I32_wrap_i64(v340), v347, v352)
	mBase = m.M
	v354 = m.ExcPending
	if v354 != 0 {
		goto L2
	} else {
		goto L107
	}
L107:
	;
	if v267 == int32(0) {
		v393 = v339
		v394 = v353
		goto L89
	} else {
		goto L108
	}
L108:
	;
	v357 = base.I64_extend_i32_u(v347)
	if v353 == int64(-1) {
		goto L109
	} else {
		goto L110
	}
L109:
	;
	v363 = *(*int64)(unsafe.Add(mBase, uint32(v18)+56))
	*(*int64)(unsafe.Add(mBase, uint32(v18)+56)) = v363 + v357
	v367 = int32(1)
	goto L105
L110:
	;
	if v353 != v357<<(uint(int64(3))%64) {
		v393 = v339
		v394 = v353
		goto L89
	} else {
		goto L111
	}
L111:
	;
	goto L109
L112:
	;
	v378 = v372 & (v267 ^ int32(-1))
	goto L114
L113:
	;
	v378 = v372 | v267
	goto L114
L114:
	;
	v379 = v378
	v380 = v367
	v385 = v377
	goto L90
L115:
	;
	v393 = v380
	v394 = v390
	goto L89
L116:
	;
	v412 = int64(-1)
	v413 = *(*int64)(unsafe.Add(mBase, uint32(v18)+56))
	if v394 == v412 {
		goto L120
	} else {
		goto L121
	}
L117:
	;
	if v394 != base.I64_extend_i32_s(v393)<<(uint(int64(3))%64) {
		goto L116
	} else {
		goto L118
	}
L118:
	;
	F_addReplyLongLong(m, l0, int64(-1))
	mBase = m.M
	v411 = m.ExcPending
	if v411 != 0 {
		goto L2
	} else {
		goto L119
	}
L119:
	;
	goto L1
L120:
	;
	v419 = v412
	goto L122
L121:
	;
	v419 = v413<<(uint(int64(3))%64) + v394
	goto L122
L122:
	;
	F_addReplyLongLong(m, l0, v419)
	mBase = m.M
	v421 = m.ExcPending
	if v421 != 0 {
		goto L2
	} else {
		goto L123
	}
L123:
	;
	goto L1
}
func F_blmoveCommand(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
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
	var v43 int32
	_ = v43
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
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v104 int32
	_ = v104
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
	var v114 int32
	_ = v114
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v146 int32
	_ = v146
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v152 int32
	_ = v152
	var v156 int32
	_ = v156
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v191 int64
	_ = v191
	var v193 int32
	_ = v193
	var v195 int32
	_ = v195
	var v197 int32
	_ = v197
	var v199 int32
	_ = v199
	var v201 int32
	_ = v201
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v10 = int32(1)
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v13 = *(*int32)(unsafe.Add(mBase, uint32(v12)+12))
	v14 = F_objectGetVal(m, v13)
	mBase = m.M
	v15 = int32(_a2386)
	v18 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14))))
	if v18 != 0 {
		goto L7
	} else {
		goto L8
	}
L1:
	;
	m.G0 = v8 + int32(16)
	return
L2:
	;
	v199 = *(*int32)(unsafe.Add(mBase, _consts[33]))
	F_addReplyErrorObject(m, l0, v199)
	mBase = m.M
	v201 = m.ExcPending
	if v201 != 0 {
		goto L58
	} else {
		goto L63
	}
L3:
	;
	v195 = *(*int32)(unsafe.Add(mBase, _consts[33]))
	F_addReplyErrorObject(m, l0, v195)
	mBase = m.M
	v197 = m.ExcPending
	if v197 != 0 {
		goto L58
	} else {
		goto L62
	}
L4:
	;
	v98 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v99 = *(*int32)(unsafe.Add(mBase, uint32(v98)+16))
	v100 = F_objectGetVal(m, v99)
	mBase = m.M
	v101 = int32(_a2386)
	v104 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v100))))
	if v104 != 0 {
		goto L34
	} else {
		goto L35
	}
L5:
	;
	if v50-v52 == int32(0) {
		v97 = v10
		goto L4
	} else {
		goto L17
	}
L6:
	;
	v50 = F_tolower(m, v46)
	mBase = m.M
	v51 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v47))))
	v52 = F_tolower(m, v51)
	mBase = m.M
	goto L5
L7:
	;
	v20 = v14
	v21 = v15
	v22 = v18
	goto L10
L8:
	;
	v46 = int32(0)
	v47 = v15
	goto L6
L9:
	;
	v46 = v43 & int32(255)
	v47 = v42
	goto L6
L10:
	;
	v24 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21))))
	if v24 == int32(0) {
		v42 = v21
		v43 = v22
		goto L9
	} else {
		goto L12
	}
L11:
	;
	v42 = v36
	v43 = int32(0)
	goto L9
L12:
	;
	v28 = v22 & int32(255)
	if v28 == v24 {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	v35 = int32(1)
	v36 = v21 + v35
	v37 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20)+1)))
	if v37 != 0 {
		v20 = v20 + v35
		v21 = v36
		v22 = v37
		goto L10
	} else {
		goto L16
	}
L14:
	;
	v30 = F_tolower(m, v28)
	mBase = m.M
	v31 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21))))
	v32 = F_tolower(m, v31)
	mBase = m.M
	if v30 == v32 {
		goto L13
	} else {
		goto L15
	}
L15:
	;
	v34 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20))))
	v42 = v21
	v43 = v34
	goto L9
L16:
	;
	goto L11
L17:
	;
	v56 = F_objectGetVal(m, v13)
	mBase = m.M
	v57 = int32(_a2387)
	v60 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v56))))
	if v60 != 0 {
		goto L20
	} else {
		goto L21
	}
L18:
	;
	if v92-v94 != 0 {
		goto L3
	} else {
		goto L30
	}
L19:
	;
	v92 = F_tolower(m, v88)
	mBase = m.M
	v93 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v89))))
	v94 = F_tolower(m, v93)
	mBase = m.M
	goto L18
L20:
	;
	v62 = v56
	v63 = v57
	v64 = v60
	goto L23
L21:
	;
	v88 = int32(0)
	v89 = v57
	goto L19
L22:
	;
	v88 = v85 & int32(255)
	v89 = v84
	goto L19
L23:
	;
	v66 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v63))))
	if v66 == int32(0) {
		v84 = v63
		v85 = v64
		goto L22
	} else {
		goto L25
	}
L24:
	;
	v84 = v78
	v85 = int32(0)
	goto L22
L25:
	;
	v70 = v64 & int32(255)
	if v70 == v66 {
		goto L26
	} else {
		goto L27
	}
L26:
	;
	v77 = int32(1)
	v78 = v63 + v77
	v79 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v62)+1)))
	if v79 != 0 {
		v62 = v62 + v77
		v63 = v78
		v64 = v79
		goto L23
	} else {
		goto L29
	}
L27:
	;
	v72 = F_tolower(m, v70)
	mBase = m.M
	v73 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v63))))
	v74 = F_tolower(m, v73)
	mBase = m.M
	if v72 == v74 {
		goto L26
	} else {
		goto L28
	}
L28:
	;
	v76 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v62))))
	v84 = v63
	v85 = v76
	goto L22
L29:
	;
	goto L24
L30:
	;
	v97 = int32(0)
	goto L4
L31:
	;
	v184 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v185 = *(*int32)(unsafe.Add(mBase, uint32(v184)+20))
	v189 = F_getTimeoutFromObjectOrReply(m, l0, v185, v8+int32(8), int32(0))
	mBase = m.M
	v190 = m.ExcPending
	if v190 != 0 {
		goto L58
	} else {
		goto L59
	}
L32:
	;
	if v136-v138 == int32(0) {
		v183 = v10
		goto L31
	} else {
		goto L44
	}
L33:
	;
	v136 = F_tolower(m, v132)
	mBase = m.M
	v137 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v133))))
	v138 = F_tolower(m, v137)
	mBase = m.M
	goto L32
L34:
	;
	v106 = v100
	v107 = v101
	v108 = v104
	goto L37
L35:
	;
	v132 = int32(0)
	v133 = v101
	goto L33
L36:
	;
	v132 = v129 & int32(255)
	v133 = v128
	goto L33
L37:
	;
	v110 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v107))))
	if v110 == int32(0) {
		v128 = v107
		v129 = v108
		goto L36
	} else {
		goto L39
	}
L38:
	;
	v128 = v122
	v129 = int32(0)
	goto L36
L39:
	;
	v114 = v108 & int32(255)
	if v114 == v110 {
		goto L40
	} else {
		goto L41
	}
L40:
	;
	v121 = int32(1)
	v122 = v107 + v121
	v123 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v106)+1)))
	if v123 != 0 {
		v106 = v106 + v121
		v107 = v122
		v108 = v123
		goto L37
	} else {
		goto L43
	}
L41:
	;
	v116 = F_tolower(m, v114)
	mBase = m.M
	v117 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v107))))
	v118 = F_tolower(m, v117)
	mBase = m.M
	if v116 == v118 {
		goto L40
	} else {
		goto L42
	}
L42:
	;
	v120 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v106))))
	v128 = v107
	v129 = v120
	goto L36
L43:
	;
	goto L38
L44:
	;
	v142 = F_objectGetVal(m, v99)
	mBase = m.M
	v143 = int32(_a2387)
	v146 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v142))))
	if v146 != 0 {
		goto L47
	} else {
		goto L48
	}
L45:
	;
	if v178-v180 != 0 {
		goto L2
	} else {
		goto L57
	}
L46:
	;
	v178 = F_tolower(m, v174)
	mBase = m.M
	v179 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v175))))
	v180 = F_tolower(m, v179)
	mBase = m.M
	goto L45
L47:
	;
	v148 = v142
	v149 = v143
	v150 = v146
	goto L50
L48:
	;
	v174 = int32(0)
	v175 = v143
	goto L46
L49:
	;
	v174 = v171 & int32(255)
	v175 = v170
	goto L46
L50:
	;
	v152 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v149))))
	if v152 == int32(0) {
		v170 = v149
		v171 = v150
		goto L49
	} else {
		goto L52
	}
L51:
	;
	v170 = v164
	v171 = int32(0)
	goto L49
L52:
	;
	v156 = v150 & int32(255)
	if v156 == v152 {
		goto L53
	} else {
		goto L54
	}
L53:
	;
	v163 = int32(1)
	v164 = v149 + v163
	v165 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v148)+1)))
	if v165 != 0 {
		v148 = v148 + v163
		v149 = v164
		v150 = v165
		goto L50
	} else {
		goto L56
	}
L54:
	;
	v158 = F_tolower(m, v156)
	mBase = m.M
	v159 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v149))))
	v160 = F_tolower(m, v159)
	mBase = m.M
	if v158 == v160 {
		goto L53
	} else {
		goto L55
	}
L55:
	;
	v162 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v148))))
	v170 = v149
	v171 = v162
	goto L49
L56:
	;
	goto L51
L57:
	;
	v183 = int32(0)
	goto L31
L58:
	;
	return
L59:
	;
	if v189 != 0 {
		goto L1
	} else {
		goto L60
	}
L60:
	;
	v191 = *(*int64)(unsafe.Add(mBase, uint32(v8)+8))
	F_blmoveGenericCommand(m, l0, v97, v183, v191)
	mBase = m.M
	v193 = m.ExcPending
	if v193 != 0 {
		goto L58
	} else {
		goto L61
	}
L61:
	;
	goto L1
L62:
	;
	goto L1
L63:
	;
	goto L1
}
func F_block(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
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
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
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
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v90 int32
	_ = v90
	var v102 int32
	_ = v102
	var v104 int32
	_ = v104
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v112 int32
	_ = v112
	var v114 int32
	_ = v114
	var v121 int32
	_ = v121
	var v127 int32
	_ = v127
	var v131 int32
	_ = v131
	var v137 int32
	_ = v137
	var v162 int32
	_ = v162
	var v166 int32
	_ = v166
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v172 int32
	_ = v172
	var v174 int32
	_ = v174
	v2 = int32(0)
	v13 = m.G0
	v15 = v13 - int32(16)
	m.G0 = v15
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	*(*uint8)(unsafe.Add(mBase, uint32(v15)+14)) = uint8(v2)
	*(*int32)(unsafe.Add(mBase, uint32(v15)+8)) = int32(-1)
	v23 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17)+50)))
	*(*uint8)(unsafe.Add(mBase, uint32(v15)+13)) = uint8(v2)
	*(*uint8)(unsafe.Add(mBase, uint32(v15)+12)) = uint8(v23)
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v17)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+4)) = v27
	*(*int32)(unsafe.Add(mBase, uint32(v17)+20)) = v15 + int32(4)
	F_chunk(m, l0)
	mBase = m.M
	v33 = m.ExcPending
	if v33 != 0 {
		return
	} else {
		v34 = *(*int32)(unsafe.Add(mBase, uint32(v17)+20))
		v35 = *(*int32)(unsafe.Add(mBase, uint32(v34)))
		*(*int32)(unsafe.Add(mBase, uint32(v17)+20)) = v35
		v37 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v34)+8)))
		v38 = *(*int32)(unsafe.Add(mBase, uint32(v17)+12))
		v39 = *(*int32)(unsafe.Add(mBase, uint32(v38)+48))
		v40 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v39)+50)))
		if base.Ui32(v40) <= base.Ui32(v37) {
		} else {
			v43 = v39 + int32(172)
			v44 = *(*int32)(unsafe.Add(mBase, uint32(v39)+24))
			v45 = *(*int32)(unsafe.Add(mBase, uint32(v39)))
			v46 = *(*int32)(unsafe.Add(mBase, uint32(v45)+24))
			v49 = (v40 - v37) & int32(3)
			if v49 != 0 {
				v50 = v40
				v53 = v2
				for {
					v63 = v50 + int32(-1)
					v64 = int32(1)
					v67 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v43+v63<<(uint(v64)%32)))))
					*(*int32)(unsafe.Add(mBase, uint32(v46+v67*int32(12))+8)) = v44
					v73 = v53 + v64
					if v73 != v49 {
						v50 = v63
						v53 = v73
						continue
					} else {
						break
					}
					break
				}
				v75 = v63
			} else {
				v75 = v40
			}
			if base.Ui32(int32(-4)) < base.Ui32(v37-v40) {
				v137 = v75
			} else {
				v90 = v75
				for {
					v102 = int32(1)
					v104 = v90<<(uint(v102)%32) + v43
					v107 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v104+int32(-2)))))
					v108 = int32(12)
					*(*int32)(unsafe.Add(mBase, uint32(v46+v107*v108)+8)) = v44
					v112 = int32(-4)
					v114 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v104+v112))))
					*(*int32)(unsafe.Add(mBase, uint32(v46+v114*v108)+8)) = v44
					v121 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v104+int32(-6)))))
					*(*int32)(unsafe.Add(mBase, uint32(v46+v121*v108)+8)) = v44
					v127 = v90 + v112
					v131 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v43+v127<<(uint(v102)%32)))))
					*(*int32)(unsafe.Add(mBase, uint32(v46+v131*v108)+8)) = v44
					if base.Ui32(v37) < base.Ui32(v127) {
						v90 = v127
						continue
					} else {
						break
					}
					break
				}
				v137 = v127
			}
			*(*uint8)(unsafe.Add(mBase, uint32(v39)+50)) = uint8(v137)
		}
		v162 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v34)+9)))
		if v162 == int32(0) {
			v170 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17)+50)))
			*(*int32)(unsafe.Add(mBase, uint32(v17)+36)) = v170
			v172 = *(*int32)(unsafe.Add(mBase, uint32(v34)+4))
			F_luaK_patchtohere(m, v17, v172)
			mBase = m.M
			v174 = m.ExcPending
			if v174 != 0 {
				return
			} else {
				m.G0 = v15 + int32(16)
				return
			}
		} else {
			v166 = int32(0)
			v168 = F_luaK_codeABC(m, v17, int32(35), v37, v166, v166)
			mBase = m.M
			v169 = m.ExcPending
			if v169 != 0 {
				return
			} else {
				v170 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17)+50)))
				*(*int32)(unsafe.Add(mBase, uint32(v17)+36)) = v170
				v172 = *(*int32)(unsafe.Add(mBase, uint32(v34)+4))
				F_luaK_patchtohere(m, v17, v172)
				mBase = m.M
				v174 = m.ExcPending
				if v174 != 0 {
					return
				} else {
					m.G0 = v15 + int32(16)
					return
				}
			}
		}
	}
}
func F_blockClient(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int64
	_ = v43
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v71 int32
	_ = v71
	var v75 int32
	_ = v75
	var v80 int32
	_ = v80
	v6 = int32(1)
	v7 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+200)))
	if v7&v6 != 0 {
		v14 = v6
		v17 = v14
	} else {
		v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+216))
		if v10 != 0 {
			v12 = F_isImportSlotMigrationJob(m, v10)
			mBase = m.M
			v14 = v12
			v17 = v14
		} else {
			v17 = int32(0)
		}
	}
	if v17 == int32(0) {
		v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+116))
		if v28 != 0 {
			v53 = v28
			v55 = *(*int32)(unsafe.Add(mBase, uint32(l0)+200))
			*(*int32)(unsafe.Add(mBase, uint32(l0)+200)) = v55 | int32(16)
			*(*int32)(unsafe.Add(mBase, uint32(v53))) = l1
			v60 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+203)))
			if v60&int32(64) != 0 {
			} else {
				v63 = int32(_a20)
				v65 = *(*int32)(unsafe.Add(mBase, _consts[73]))
				*(*int32)(unsafe.Add(mBase, _consts[73])) = v65 + int32(1)
			}
			v71 = l1 << (uint(int32(2)) % 32)
			v75 = *(*int32)(unsafe.Add(mBase, uint32(v71)+uint32(_consts[74])))
			*(*int32)(unsafe.Add(mBase, uint32(v71)+uint32(_consts[74]))) = v75 + int32(1)
			F_addClientToTimeoutTable(m, l0)
			mBase = m.M
			v80 = m.ExcPending
			if v80 != 0 {
				return
			} else {
				return
			}
		} else {
			v30 = F_valkey_malloc(m, int32(56))
			mBase = m.M
			v31 = m.ExcPending
			if v31 != 0 {
				return
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(l0)+116)) = v30
				v33 = int32(0)
				*(*int32)(unsafe.Add(mBase, uint32(v30)+16)) = v33
				*(*int64)(unsafe.Add(mBase, uint32(v30)+8)) = int64(0)
				*(*int32)(unsafe.Add(mBase, uint32(v30))) = v33
				v40 = F_dictCreate(m, int32(_a182))
				mBase = m.M
				v41 = m.ExcPending
				if v41 != 0 {
					return
				} else {
					v42 = *(*int32)(unsafe.Add(mBase, uint32(l0)+116))
					v43 = int64(0)
					*(*int64)(unsafe.Add(mBase, uint32(v42)+40)) = v43
					*(*int64)(unsafe.Add(mBase, uint32(v42)+28)) = v43
					*(*int32)(unsafe.Add(mBase, uint32(v42)+24)) = v40
					*(*int32)(unsafe.Add(mBase, uint32(v42)+20)) = int32(0)
					v50 = *(*int32)(unsafe.Add(mBase, uint32(l0)+116))
					*(*int64)(unsafe.Add(mBase, uint32(v50)+48)) = v43
					v53 = v50
					v55 = *(*int32)(unsafe.Add(mBase, uint32(l0)+200))
					*(*int32)(unsafe.Add(mBase, uint32(l0)+200)) = v55 | int32(16)
					*(*int32)(unsafe.Add(mBase, uint32(v53))) = l1
					v60 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+203)))
					if v60&int32(64) != 0 {
					} else {
						v63 = int32(_a20)
						v65 = *(*int32)(unsafe.Add(mBase, _consts[73]))
						*(*int32)(unsafe.Add(mBase, _consts[73])) = v65 + int32(1)
					}
					v71 = l1 << (uint(int32(2)) % 32)
					v75 = *(*int32)(unsafe.Add(mBase, uint32(v71)+uint32(_consts[74])))
					*(*int32)(unsafe.Add(mBase, uint32(v71)+uint32(_consts[74]))) = v75 + int32(1)
					F_addClientToTimeoutTable(m, l0)
					mBase = m.M
					v80 = m.ExcPending
					if v80 != 0 {
						return
					} else {
						return
					}
				}
			}
		}
	} else {
		switch l1 + int32(-3) {
		case 0, 3:
			v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+116))
			if v28 != 0 {
				v53 = v28
				v55 = *(*int32)(unsafe.Add(mBase, uint32(l0)+200))
				*(*int32)(unsafe.Add(mBase, uint32(l0)+200)) = v55 | int32(16)
				*(*int32)(unsafe.Add(mBase, uint32(v53))) = l1
				v60 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+203)))
				if v60&int32(64) != 0 {
				} else {
					v63 = int32(_a20)
					v65 = *(*int32)(unsafe.Add(mBase, _consts[73]))
					*(*int32)(unsafe.Add(mBase, _consts[73])) = v65 + int32(1)
				}
				v71 = l1 << (uint(int32(2)) % 32)
				v75 = *(*int32)(unsafe.Add(mBase, uint32(v71)+uint32(_consts[74])))
				*(*int32)(unsafe.Add(mBase, uint32(v71)+uint32(_consts[74]))) = v75 + int32(1)
				F_addClientToTimeoutTable(m, l0)
				mBase = m.M
				v80 = m.ExcPending
				if v80 != 0 {
					return
				} else {
					return
				}
			} else {
				v30 = F_valkey_malloc(m, int32(56))
				mBase = m.M
				v31 = m.ExcPending
				if v31 != 0 {
					return
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(l0)+116)) = v30
					v33 = int32(0)
					*(*int32)(unsafe.Add(mBase, uint32(v30)+16)) = v33
					*(*int64)(unsafe.Add(mBase, uint32(v30)+8)) = int64(0)
					*(*int32)(unsafe.Add(mBase, uint32(v30))) = v33
					v40 = F_dictCreate(m, int32(_a182))
					mBase = m.M
					v41 = m.ExcPending
					if v41 != 0 {
						return
					} else {
						v42 = *(*int32)(unsafe.Add(mBase, uint32(l0)+116))
						v43 = int64(0)
						*(*int64)(unsafe.Add(mBase, uint32(v42)+40)) = v43
						*(*int64)(unsafe.Add(mBase, uint32(v42)+28)) = v43
						*(*int32)(unsafe.Add(mBase, uint32(v42)+24)) = v40
						*(*int32)(unsafe.Add(mBase, uint32(v42)+20)) = int32(0)
						v50 = *(*int32)(unsafe.Add(mBase, uint32(l0)+116))
						*(*int64)(unsafe.Add(mBase, uint32(v50)+48)) = v43
						v53 = v50
						v55 = *(*int32)(unsafe.Add(mBase, uint32(l0)+200))
						*(*int32)(unsafe.Add(mBase, uint32(l0)+200)) = v55 | int32(16)
						*(*int32)(unsafe.Add(mBase, uint32(v53))) = l1
						v60 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+203)))
						if v60&int32(64) != 0 {
						} else {
							v63 = int32(_a20)
							v65 = *(*int32)(unsafe.Add(mBase, _consts[73]))
							*(*int32)(unsafe.Add(mBase, _consts[73])) = v65 + int32(1)
						}
						v71 = l1 << (uint(int32(2)) % 32)
						v75 = *(*int32)(unsafe.Add(mBase, uint32(v71)+uint32(_consts[74])))
						*(*int32)(unsafe.Add(mBase, uint32(v71)+uint32(_consts[74]))) = v75 + int32(1)
						F_addClientToTimeoutTable(m, l0)
						mBase = m.M
						v80 = m.ExcPending
						if v80 != 0 {
							return
						} else {
							return
						}
					}
				}
			}
		default:
			F__serverAssert(m, int32(_a183), int32(_a184), int32(108))
			mBase = m.M
			v26 = m.ExcPending
			if v26 != 0 {
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
func F_blockClientForReplicaAck(m *base.Module, l0 int32, l1 int64, l2 int64, l3 int32, l4 int32) {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int64
	_ = v23
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
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
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v59 int32
	_ = v59
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+116))
	if v8 != 0 {
		v33 = v8
		*(*int64)(unsafe.Add(mBase, uint32(v33)+40)) = l2
		*(*int64)(unsafe.Add(mBase, uint32(v33)+8)) = l1
		*(*int32)(unsafe.Add(mBase, uint32(v33)+32)) = l4
		*(*int32)(unsafe.Add(mBase, uint32(v33)+28)) = l3
		v40 = *(*int32)(unsafe.Add(mBase, _consts[80]))
		v41 = F_listAddNodeHead(m, v40, l0)
		mBase = m.M
		v42 = m.ExcPending
		if v42 != 0 {
			return
		} else {
			v43 = *(*int32)(unsafe.Add(mBase, uint32(l0)+116))
			v44 = *(*int32)(unsafe.Add(mBase, uint32(v43)+20))
			if v44 == int32(0) {
				v54 = *(*int32)(unsafe.Add(mBase, _consts[80]))
				v55 = *(*int32)(unsafe.Add(mBase, uint32(v54)))
				*(*int32)(unsafe.Add(mBase, uint32(v43)+20)) = v55
				F_blockClient(m, l0, int32(2))
				mBase = m.M
				v59 = m.ExcPending
				if v59 != 0 {
					return
				} else {
					return
				}
			} else {
				F__serverAssert(m, int32(_a192), int32(_a184), int32(701))
				mBase = m.M
				v51 = m.ExcPending
				if v51 != 0 {
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
	} else {
		v10 = F_valkey_malloc(m, int32(56))
		mBase = m.M
		v11 = m.ExcPending
		if v11 != 0 {
			return
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(l0)+116)) = v10
			v13 = int32(0)
			*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = v13
			*(*int64)(unsafe.Add(mBase, uint32(v10)+8)) = int64(0)
			*(*int32)(unsafe.Add(mBase, uint32(v10))) = v13
			v20 = F_dictCreate(m, int32(_a182))
			mBase = m.M
			v21 = m.ExcPending
			if v21 != 0 {
				return
			} else {
				v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+116))
				v23 = int64(0)
				*(*int64)(unsafe.Add(mBase, uint32(v22)+40)) = v23
				*(*int64)(unsafe.Add(mBase, uint32(v22)+28)) = v23
				*(*int32)(unsafe.Add(mBase, uint32(v22)+24)) = v20
				*(*int32)(unsafe.Add(mBase, uint32(v22)+20)) = int32(0)
				v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)+116))
				*(*int64)(unsafe.Add(mBase, uint32(v30)+48)) = v23
				v33 = v30
				*(*int64)(unsafe.Add(mBase, uint32(v33)+40)) = l2
				*(*int64)(unsafe.Add(mBase, uint32(v33)+8)) = l1
				*(*int32)(unsafe.Add(mBase, uint32(v33)+32)) = l4
				*(*int32)(unsafe.Add(mBase, uint32(v33)+28)) = l3
				v40 = *(*int32)(unsafe.Add(mBase, _consts[80]))
				v41 = F_listAddNodeHead(m, v40, l0)
				mBase = m.M
				v42 = m.ExcPending
				if v42 != 0 {
					return
				} else {
					v43 = *(*int32)(unsafe.Add(mBase, uint32(l0)+116))
					v44 = *(*int32)(unsafe.Add(mBase, uint32(v43)+20))
					if v44 == int32(0) {
						v54 = *(*int32)(unsafe.Add(mBase, _consts[80]))
						v55 = *(*int32)(unsafe.Add(mBase, uint32(v54)))
						*(*int32)(unsafe.Add(mBase, uint32(v43)+20)) = v55
						F_blockClient(m, l0, int32(2))
						mBase = m.M
						v59 = m.ExcPending
						if v59 != 0 {
							return
						} else {
							return
						}
					} else {
						F__serverAssert(m, int32(_a192), int32(_a184), int32(701))
						mBase = m.M
						v51 = m.ExcPending
						if v51 != 0 {
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
	}
}
func F_blockForKeys(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int64, l5 int32) {
	mBase := m.M
	_ = mBase
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int64
	_ = v32
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v59 int32
	_ = v59
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
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
	var v76 int32
	_ = v76
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
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
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
	var v101 int32
	_ = v101
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v114 int32
	_ = v114
	var v116 int32
	_ = v116
	var v119 int32
	_ = v119
	var v121 int64
	_ = v121
	var v127 int32
	_ = v127
	var v129 int32
	_ = v129
	var v137 int32
	_ = v137
	var v145 int32
	_ = v145
	var v150 int32
	_ = v150
	v13 = m.G0
	v15 = v13 - int32(16)
	m.G0 = v15
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+116))
	if v17 != 0 {
		v42 = v17
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v44 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+206)))
	if v44&int32(32) != 0 {
		goto L6
	} else {
		goto L7
	}
L2:
	;
	v19 = F_valkey_malloc(m, int32(56))
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	return
L4:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+116)) = v19
	v22 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v19)+16)) = v22
	*(*int64)(unsafe.Add(mBase, uint32(v19)+8)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v19))) = v22
	v29 = F_dictCreate(m, int32(_a182))
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		goto L3
	} else {
		goto L5
	}
L5:
	;
	v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)+116))
	v32 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v31)+40)) = v32
	*(*int64)(unsafe.Add(mBase, uint32(v31)+28)) = v32
	*(*int32)(unsafe.Add(mBase, uint32(v31)+24)) = v29
	*(*int32)(unsafe.Add(mBase, uint32(v31)+20)) = int32(0)
	v39 = *(*int32)(unsafe.Add(mBase, uint32(l0)+116))
	*(*int64)(unsafe.Add(mBase, uint32(v39)+48)) = v32
	v42 = v39
	goto L1
L6:
	;
	if l3 < int32(1) {
		v137 = v42
		goto L8
	} else {
		goto L9
	}
L7:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v42)+8)) = l4
	goto L6
L8:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v137)+16)) = l5
	if l1 == int32(3) {
		goto L34
	} else {
		goto L35
	}
L9:
	;
	v59 = int32(0)
	goto L10
L10:
	;
	v63 = *(*int32)(unsafe.Add(mBase, uint32(l0)+116))
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v63)+24))
	v67 = l2 + v59<<(uint(int32(2))%32)
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v67)))
	v70 = F_dictAddRaw(m, v64, v68, int32(0))
	mBase = m.M
	v71 = m.ExcPending
	if v71 != 0 {
		goto L3
	} else {
		goto L13
	}
L11:
	;
	v129 = *(*int32)(unsafe.Add(mBase, uint32(l0)+116))
	v137 = v129
	goto L8
L12:
	;
	v127 = v59 + int32(1)
	if v127 != l3 {
		v59 = v127
		goto L10
	} else {
		goto L33
	}
L13:
	;
	if v70 == int32(0) {
		goto L12
	} else {
		goto L14
	}
L14:
	;
	v74 = *(*int32)(unsafe.Add(mBase, uint32(v67)))
	F_incrRefCount(m, v74)
	mBase = m.M
	v76 = m.ExcPending
	if v76 != 0 {
		goto L3
	} else {
		goto L15
	}
L15:
	;
	v77 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v77)+12))
	v79 = *(*int32)(unsafe.Add(mBase, uint32(v67)))
	v82 = F_dictAddRaw(m, v78, v79, v15+int32(12))
	mBase = m.M
	v83 = m.ExcPending
	if v83 != 0 {
		goto L3
	} else {
		goto L18
	}
L16:
	;
	v97 = F_listAddNodeTail(m, v96, l0)
	mBase = m.M
	v98 = m.ExcPending
	if v98 != 0 {
		goto L3
	} else {
		goto L24
	}
L17:
	;
	v94 = *(*int32)(unsafe.Add(mBase, uint32(v15)+12))
	v95 = *(*int32)(unsafe.Add(mBase, uint32(v94)+8))
	goto L23
L18:
	;
	if v82 == int32(0) {
		goto L17
	} else {
		goto L19
	}
L19:
	;
	v86 = F_listCreate(m)
	mBase = m.M
	v87 = m.ExcPending
	if v87 != 0 {
		goto L3
	} else {
		goto L20
	}
L20:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v82)+8)) = v86
	goto L21
L21:
	;
	v91 = *(*int32)(unsafe.Add(mBase, uint32(v67)))
	F_incrRefCount(m, v91)
	mBase = m.M
	v93 = m.ExcPending
	if v93 != 0 {
		goto L3
	} else {
		goto L22
	}
L22:
	;
	v96 = v86
	goto L16
L23:
	;
	v96 = v95
	goto L16
L24:
	;
	v101 = *(*int32)(unsafe.Add(mBase, uint32(v96)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v70)+8)) = v101
	goto L25
L25:
	;
	if l5 == int32(0) {
		goto L12
	} else {
		goto L26
	}
L26:
	;
	v105 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v106 = *(*int32)(unsafe.Add(mBase, uint32(v105)+16))
	v107 = *(*int32)(unsafe.Add(mBase, uint32(v67)))
	v110 = F_dictAddRaw(m, v106, v107, v15+int32(12))
	mBase = m.M
	v111 = m.ExcPending
	if v111 != 0 {
		goto L3
	} else {
		goto L28
	}
L27:
	;
	v119 = *(*int32)(unsafe.Add(mBase, uint32(v15)+12))
	v121 = *(*int64)(unsafe.Add(mBase, uint32(v119)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v119)+8)) = v121 + int64(1)
	goto L32
L28:
	;
	if v110 == int32(0) {
		goto L27
	} else {
		goto L29
	}
L29:
	;
	v114 = *(*int32)(unsafe.Add(mBase, uint32(v67)))
	F_incrRefCount(m, v114)
	mBase = m.M
	v116 = m.ExcPending
	if v116 != 0 {
		goto L3
	} else {
		goto L30
	}
L30:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v110)+8)) = int64(1)
	goto L31
L31:
	;
	goto L12
L32:
	;
	goto L12
L33:
	;
	goto L11
L34:
	;
	F_blockClient(m, l0, l1)
	mBase = m.M
	v150 = m.ExcPending
	if v150 != 0 {
		goto L3
	} else {
		goto L36
	}
L35:
	;
	v145 = *(*int32)(unsafe.Add(mBase, uint32(l0)+204))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+204)) = v145 | int32(2)
	goto L34
L36:
	;
	m.G0 = v15 + int32(16)
	return
}
func F_blockPostponeClient(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int64
	_ = v19
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+116))
	if v4 != 0 {
		v29 = v4
		*(*int64)(unsafe.Add(mBase, uint32(v29)+8)) = int64(0)
		F_blockClient(m, l0, int32(6))
		mBase = m.M
		v35 = m.ExcPending
		if v35 != 0 {
			return
		} else {
			v37 = *(*int32)(unsafe.Add(mBase, _consts[76]))
			v38 = F_listAddNodeTail(m, v37, l0)
			mBase = m.M
			v39 = m.ExcPending
			if v39 != 0 {
				return
			} else {
				v40 = *(*int32)(unsafe.Add(mBase, uint32(l0)+116))
				v41 = *(*int32)(unsafe.Add(mBase, uint32(v40)+20))
				if v41 == int32(0) {
					v51 = *(*int32)(unsafe.Add(mBase, _consts[76]))
					v52 = *(*int32)(unsafe.Add(mBase, uint32(v51)+4))
					*(*int32)(unsafe.Add(mBase, uint32(v40)+20)) = v52
					v54 = *(*int32)(unsafe.Add(mBase, uint32(l0)+204))
					*(*int32)(unsafe.Add(mBase, uint32(l0)+204)) = v54 | int32(2)
					return
				} else {
					F__serverAssert(m, int32(_a193), int32(_a184), int32(714))
					mBase = m.M
					v48 = m.ExcPending
					if v48 != 0 {
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
	} else {
		v6 = F_valkey_malloc(m, int32(56))
		mBase = m.M
		v7 = m.ExcPending
		if v7 != 0 {
			return
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(l0)+116)) = v6
			v9 = int32(0)
			*(*int32)(unsafe.Add(mBase, uint32(v6)+16)) = v9
			*(*int64)(unsafe.Add(mBase, uint32(v6)+8)) = int64(0)
			*(*int32)(unsafe.Add(mBase, uint32(v6))) = v9
			v16 = F_dictCreate(m, int32(_a182))
			mBase = m.M
			v17 = m.ExcPending
			if v17 != 0 {
				return
			} else {
				v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+116))
				v19 = int64(0)
				*(*int64)(unsafe.Add(mBase, uint32(v18)+40)) = v19
				*(*int64)(unsafe.Add(mBase, uint32(v18)+28)) = v19
				*(*int32)(unsafe.Add(mBase, uint32(v18)+24)) = v16
				*(*int32)(unsafe.Add(mBase, uint32(v18)+20)) = int32(0)
				v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+116))
				*(*int64)(unsafe.Add(mBase, uint32(v26)+48)) = v19
				v29 = v26
				*(*int64)(unsafe.Add(mBase, uint32(v29)+8)) = int64(0)
				F_blockClient(m, l0, int32(6))
				mBase = m.M
				v35 = m.ExcPending
				if v35 != 0 {
					return
				} else {
					v37 = *(*int32)(unsafe.Add(mBase, _consts[76]))
					v38 = F_listAddNodeTail(m, v37, l0)
					mBase = m.M
					v39 = m.ExcPending
					if v39 != 0 {
						return
					} else {
						v40 = *(*int32)(unsafe.Add(mBase, uint32(l0)+116))
						v41 = *(*int32)(unsafe.Add(mBase, uint32(v40)+20))
						if v41 == int32(0) {
							v51 = *(*int32)(unsafe.Add(mBase, _consts[76]))
							v52 = *(*int32)(unsafe.Add(mBase, uint32(v51)+4))
							*(*int32)(unsafe.Add(mBase, uint32(v40)+20)) = v52
							v54 = *(*int32)(unsafe.Add(mBase, uint32(l0)+204))
							*(*int32)(unsafe.Add(mBase, uint32(l0)+204)) = v54 | int32(2)
							return
						} else {
							F__serverAssert(m, int32(_a193), int32(_a184), int32(714))
							mBase = m.M
							v48 = m.ExcPending
							if v48 != 0 {
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
		}
	}
}
func F_blockedBeforeSleep(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	F_handleBlockedClientsTimeout(m)
	mBase = m.M
	v2 = m.ExcPending
	if v2 != 0 {
		return
	} else {
		v4 = *(*int32)(unsafe.Add(mBase, _consts[80]))
		v5 = *(*int32)(unsafe.Add(mBase, uint32(v4)+20))
		if v5 == int32(0) {
			F_handleClientsBlockedOnKeys(m)
			mBase = m.M
			v11 = m.ExcPending
			if v11 != 0 {
				return
			} else {
				v14 = *(*int32)(unsafe.Add(mBase, _consts[81]))
				v15 = *(*int32)(unsafe.Add(mBase, uint32(v14)+16))
				v16 = *(*int32)(unsafe.Add(mBase, uint32(v14)+12))
				if v15+v16 == int32(0) {
					v23 = *(*int32)(unsafe.Add(mBase, _consts[75]))
					v24 = *(*int32)(unsafe.Add(mBase, uint32(v23)+20))
					if v24 == int32(0) {
						return
					} else {
						F_processUnblockedClients(m)
						mBase = m.M
						v28 = m.ExcPending
						if v28 != 0 {
							return
						} else {
							return
						}
					}
				} else {
					F_moduleHandleBlockedClients(m)
					mBase = m.M
					v21 = m.ExcPending
					if v21 != 0 {
						return
					} else {
						v23 = *(*int32)(unsafe.Add(mBase, _consts[75]))
						v24 = *(*int32)(unsafe.Add(mBase, uint32(v23)+20))
						if v24 == int32(0) {
							return
						} else {
							F_processUnblockedClients(m)
							mBase = m.M
							v28 = m.ExcPending
							if v28 != 0 {
								return
							} else {
								return
							}
						}
					}
				}
			}
		} else {
			F_processClientsWaitingReplicas(m)
			mBase = m.M
			v9 = m.ExcPending
			if v9 != 0 {
				return
			} else {
				F_handleClientsBlockedOnKeys(m)
				mBase = m.M
				v11 = m.ExcPending
				if v11 != 0 {
					return
				} else {
					v14 = *(*int32)(unsafe.Add(mBase, _consts[81]))
					v15 = *(*int32)(unsafe.Add(mBase, uint32(v14)+16))
					v16 = *(*int32)(unsafe.Add(mBase, uint32(v14)+12))
					if v15+v16 == int32(0) {
						v23 = *(*int32)(unsafe.Add(mBase, _consts[75]))
						v24 = *(*int32)(unsafe.Add(mBase, uint32(v23)+20))
						if v24 == int32(0) {
							return
						} else {
							F_processUnblockedClients(m)
							mBase = m.M
							v28 = m.ExcPending
							if v28 != 0 {
								return
							} else {
								return
							}
						}
					} else {
						F_moduleHandleBlockedClients(m)
						mBase = m.M
						v21 = m.ExcPending
						if v21 != 0 {
							return
						} else {
							v23 = *(*int32)(unsafe.Add(mBase, _consts[75]))
							v24 = *(*int32)(unsafe.Add(mBase, uint32(v23)+20))
							if v24 == int32(0) {
								return
							} else {
								F_processUnblockedClients(m)
								mBase = m.M
								v28 = m.ExcPending
								if v28 != 0 {
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
func F_blockingGenericZpopCommand(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32) {
	mBase := m.M
	_ = mBase
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v20 int32
	_ = v20
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v33 int32
	_ = v33
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v59 int32
	_ = v59
	var v63 int32
	_ = v63
	var v72 int32
	_ = v72
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
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	var v104 int32
	_ = v104
	var v117 int32
	_ = v117
	var v123 int32
	_ = v123
	var v125 int64
	_ = v125
	var v128 int32
	_ = v128
	v12 = m.G0
	v14 = v12 - int32(48)
	m.G0 = v14
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v20 = *(*int32)(unsafe.Add(mBase, uint32(v16+l4<<(uint(int32(2))%32))))
	v24 = F_getTimeoutFromObjectOrReply(m, l0, v20, v14+int32(32), int32(0))
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v14 + int32(48)
	return
L2:
	;
	return
L3:
	;
	if v24 != 0 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	if l2 < int32(1) {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	v117 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+205)))
	if v117&int32(16) == int32(0) {
		goto L33
	} else {
		goto L34
	}
L6:
	;
	v33 = int32(0)
	goto L7
L7:
	;
	v40 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v44 = *(*int32)(unsafe.Add(mBase, uint32(l1+v33<<(uint(int32(2))%32))))
	v45 = F_lookupKeyWrite(m, v40, v44)
	mBase = m.M
	v46 = m.ExcPending
	if v46 != 0 {
		goto L2
	} else {
		goto L10
	}
L8:
	;
	goto L5
L9:
	;
	v104 = v33 + int32(1)
	if v104 != l2 {
		v33 = v104
		goto L7
	} else {
		goto L32
	}
L10:
	;
	if v45 == int32(0) {
		goto L9
	} else {
		goto L11
	}
L11:
	;
	v50 = F_checkType(m, l0, v45, int32(3))
	mBase = m.M
	v51 = m.ExcPending
	if v51 != 0 {
		goto L2
	} else {
		goto L12
	}
L12:
	;
	if v50 != 0 {
		goto L1
	} else {
		goto L13
	}
L13:
	;
	v52 = F_zsetLength(m, v45)
	mBase = m.M
	v53 = m.ExcPending
	if v53 != 0 {
		goto L2
	} else {
		goto L14
	}
L14:
	;
	if v52 == int32(0) {
		goto L9
	} else {
		goto L15
	}
L15:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+44)) = v44
	v59 = int32(1)
	F_genericZpopCommand(m, l0, v14+int32(44), v59, l3, v59, l5, l6, l7, int32(0))
	mBase = m.M
	v63 = m.ExcPending
	if v63 != 0 {
		goto L2
	} else {
		goto L16
	}
L16:
	;
	if l5 != int32(-1) {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	if l5 < v52 {
		goto L23
	} else {
		goto L24
	}
L18:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+4)) = v44
	if l3 == int32(1) {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	v72 = int32(280)
	goto L21
L20:
	;
	v72 = int32(276)
	goto L21
L21:
	;
	v74 = *(*int32)(unsafe.Add(mBase, uint32(v72)+uint32(_consts[27])))
	*(*int32)(unsafe.Add(mBase, uint32(v14))) = v74
	F_rewriteClientCommandVector(m, l0, int32(2), v14)
	mBase = m.M
	v78 = m.ExcPending
	if v78 != 0 {
		goto L2
	} else {
		goto L22
	}
L22:
	;
	goto L1
L23:
	;
	v80 = l5
	goto L25
L24:
	;
	v80 = v52
	goto L25
L25:
	;
	v82 = F_createStringObjectFromLongLong(m, base.I64_extend_i32_s(v80))
	mBase = m.M
	v83 = m.ExcPending
	if v83 != 0 {
		goto L2
	} else {
		goto L26
	}
L26:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+20)) = v44
	*(*int32)(unsafe.Add(mBase, uint32(v14)+24)) = v82
	if l3 == int32(1) {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	v91 = int32(280)
	goto L29
L28:
	;
	v91 = int32(276)
	goto L29
L29:
	;
	v93 = *(*int32)(unsafe.Add(mBase, uint32(v91)+uint32(_consts[27])))
	*(*int32)(unsafe.Add(mBase, uint32(v14)+16)) = v93
	F_rewriteClientCommandVector(m, l0, int32(3), v14+int32(16))
	mBase = m.M
	v99 = m.ExcPending
	if v99 != 0 {
		goto L2
	} else {
		goto L30
	}
L30:
	;
	F_decrRefCount(m, v82)
	mBase = m.M
	v101 = m.ExcPending
	if v101 != 0 {
		goto L2
	} else {
		goto L31
	}
L31:
	;
	goto L1
L32:
	;
	goto L8
L33:
	;
	v125 = *(*int64)(unsafe.Add(mBase, uint32(v14)+32))
	F_blockForKeys(m, l0, int32(5), l1, l2, v125, int32(0))
	mBase = m.M
	v128 = m.ExcPending
	if v128 != 0 {
		goto L2
	} else {
		goto L36
	}
L34:
	;
	F_addReplyNullArray(m, l0)
	mBase = m.M
	v123 = m.ExcPending
	if v123 != 0 {
		goto L2
	} else {
		goto L35
	}
L35:
	;
	goto L1
L36:
	;
	goto L1
}
func F_blockingOperationEnds(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v1 int32
	_ = v1
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	v1 = int32(0)
	v4 = *(*int32)(unsafe.Add(mBase, _consts[416]))
	v6 = v4 + int32(-1)
	*(*int32)(unsafe.Add(mBase, _consts[416])) = v6
	if v6 != 0 {
	} else {
		*(*int64)(unsafe.Add(mBase, _consts[417])) = int64(0)
	}
	return
}
func F_blockingOperationStarts(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v1 int32
	_ = v1
	var v6 int32
	_ = v6
	var v10 int32
	_ = v10
	var v11 int64
	_ = v11
	var v14 int64
	_ = v14
	var v15 int64
	_ = v15
	var v19 int64
	_ = v19
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
	var v31 int64
	_ = v31
	var v35 int64
	_ = v35
	var v40 int32
	_ = v40
	var v42 int64
	_ = v42
	v1 = int32(0)
	v6 = *(*int32)(unsafe.Add(mBase, _consts[416]))
	*(*int32)(unsafe.Add(mBase, _consts[416])) = v6 + int32(1)
	if v6 != 0 {
	} else {
		v10 = int32(0)
		v11 = F_ustime(m)
		mBase = m.M
		*(*int64)(unsafe.Add(mBase, _consts[178])) = v11
		v14 = int64(1000)
		v15 = base.I64_div_s(v11, v14)
		*(*int64)(unsafe.Add(mBase, _consts[54])) = v15
		v19 = base.I64_div_s(v11, int64(1000000))
		*(*int64)(unsafe.Add(mBase, _consts[109])) = v19
		v22 = *(*int32)(unsafe.Add(mBase, _consts[179]))
		v26 = int32(base.Ui32(v22&int32(2)) >> (uint(int32(1)) % 32))
		*(*uint8)(unsafe.Add(mBase, _consts[378])) = uint8(v26)
		v31 = base.I64_div_s(v15, int64(60000))
		*(*uint16)(unsafe.Add(mBase, _consts[376])) = uint16(v31)
		v35 = base.I64_div_s(v15, v14)
		*(*int32)(unsafe.Add(mBase, _consts[379])) = base.I32_wrap_i64(v35) & int32(16777215)
		v40 = int32(0)
		v42 = *(*int64)(unsafe.Add(mBase, _consts[54]))
		*(*int64)(unsafe.Add(mBase, _consts[417])) = v42
	}
	return
}
func F_boolConfigGet(m *base.Module, l0 int32) int32 {
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
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	v2 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+9)))
	if v2&int32(1) == int32(0) {
		v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
		v13 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
		v14 = v13
		if v14 != 0 {
			v17 = int32(_a510)
		} else {
			v17 = int32(_a512)
		}
		v18 = F_sdsnew(m, v17)
		mBase = m.M
		v19 = m.ExcPending
		if v19 != 0 {
			return int32(0)
		} else {
			return v18
		}
	} else {
		v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+84))
		v8 = F_getModuleBoolConfig(m, v7)
		mBase = m.M
		v11 = m.ExcPending
		if v11 != 0 {
			return int32(0)
		} else {
			v14 = v8
			if v14 != 0 {
				v17 = int32(_a510)
			} else {
				v17 = int32(_a512)
			}
			v18 = F_sdsnew(m, v17)
			mBase = m.M
			v19 = m.ExcPending
			if v19 != 0 {
				return int32(0)
			} else {
				return v18
			}
		}
	}
}
func F_boolConfigInit(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v2))) = v3
	return
}
func F_boolConfigSet(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
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
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
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
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	var v95 int32
	_ = v95
	var v98 int32
	_ = v98
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
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
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v119 int32
	_ = v119
	var v127 int32
	_ = v127
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v8 = int32(_a510)
	v11 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7))))
	if v11 != 0 {
		goto L6
	} else {
		goto L7
	}
L1:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = int32(_a511)
	return int32(0)
L2:
	;
	v89 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	if v89 == int32(0) {
		goto L30
	} else {
		goto L31
	}
L3:
	;
	v48 = int32(_a512)
	v51 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7))))
	if v51 != 0 {
		goto L19
	} else {
		goto L20
	}
L4:
	;
	if v43-v45 != 0 {
		goto L3
	} else {
		goto L16
	}
L5:
	;
	v43 = F_tolower(m, v39)
	mBase = m.M
	v44 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v40))))
	v45 = F_tolower(m, v44)
	mBase = m.M
	goto L4
L6:
	;
	v13 = v7
	v14 = v8
	v15 = v11
	goto L9
L7:
	;
	v39 = int32(0)
	v40 = v8
	goto L5
L8:
	;
	v39 = v36 & int32(255)
	v40 = v35
	goto L5
L9:
	;
	v17 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14))))
	if v17 == int32(0) {
		v35 = v14
		v36 = v15
		goto L8
	} else {
		goto L11
	}
L10:
	;
	v35 = v29
	v36 = int32(0)
	goto L8
L11:
	;
	v21 = v15 & int32(255)
	if v21 == v17 {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	v28 = int32(1)
	v29 = v14 + v28
	v30 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+1)))
	if v30 != 0 {
		v13 = v13 + v28
		v14 = v29
		v15 = v30
		goto L9
	} else {
		goto L15
	}
L13:
	;
	v23 = F_tolower(m, v21)
	mBase = m.M
	v24 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14))))
	v25 = F_tolower(m, v24)
	mBase = m.M
	if v23 == v25 {
		goto L12
	} else {
		goto L14
	}
L14:
	;
	v27 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13))))
	v35 = v14
	v36 = v27
	goto L8
L15:
	;
	goto L10
L16:
	;
	v88 = int32(1)
	goto L2
L17:
	;
	if v83-v85 != 0 {
		goto L1
	} else {
		goto L29
	}
L18:
	;
	v83 = F_tolower(m, v79)
	mBase = m.M
	v84 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v80))))
	v85 = F_tolower(m, v84)
	mBase = m.M
	goto L17
L19:
	;
	v53 = v7
	v54 = v48
	v55 = v51
	goto L22
L20:
	;
	v79 = int32(0)
	v80 = v48
	goto L18
L21:
	;
	v79 = v76 & int32(255)
	v80 = v75
	goto L18
L22:
	;
	v57 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v54))))
	if v57 == int32(0) {
		v75 = v54
		v76 = v55
		goto L21
	} else {
		goto L24
	}
L23:
	;
	v75 = v69
	v76 = int32(0)
	goto L21
L24:
	;
	v61 = v55 & int32(255)
	if v61 == v57 {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	v68 = int32(1)
	v69 = v54 + v68
	v70 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v53)+1)))
	if v70 != 0 {
		v53 = v53 + v68
		v54 = v69
		v55 = v70
		goto L22
	} else {
		goto L28
	}
L26:
	;
	v63 = F_tolower(m, v61)
	mBase = m.M
	v64 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v54))))
	v65 = F_tolower(m, v64)
	mBase = m.M
	if v63 == v65 {
		goto L25
	} else {
		goto L27
	}
L27:
	;
	v67 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v53))))
	v75 = v54
	v76 = v67
	goto L21
L28:
	;
	goto L23
L29:
	;
	v88 = int32(0)
	goto L2
L30:
	;
	v98 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+9)))
	if v98&int32(1) == int32(0) {
		goto L36
	} else {
		goto L37
	}
L31:
	;
	v92 = m.T0[v89].(func(*base.Module, int32, int32) int32)(m, v88, l3)
	mBase = m.M
	v95 = m.ExcPending
	if v95 != 0 {
		goto L32
	} else {
		goto L33
	}
L32:
	;
	return int32(0)
L33:
	;
	if v92 != 0 {
		goto L30
	} else {
		goto L34
	}
L34:
	;
	return int32(0)
L35:
	;
	v109 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v108 == v88 {
		goto L39
	} else {
		goto L40
	}
L36:
	;
	v106 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	v107 = *(*int32)(unsafe.Add(mBase, uint32(v106)))
	v108 = v107
	goto L35
L37:
	;
	v103 = *(*int32)(unsafe.Add(mBase, uint32(l0)+84))
	v104 = F_getModuleBoolConfig(m, v103)
	mBase = m.M
	v105 = m.ExcPending
	if v105 != 0 {
		goto L32
	} else {
		goto L38
	}
L38:
	;
	v108 = v104
	goto L35
L39:
	;
	if v109&int32(512) != 0 {
		goto L44
	} else {
		goto L45
	}
L40:
	;
	if v109&int32(256) == int32(0) {
		goto L41
	} else {
		goto L42
	}
L41:
	;
	v119 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v119))) = v88
	return int32(1)
L42:
	;
	v115 = *(*int32)(unsafe.Add(mBase, uint32(l0)+84))
	v116 = F_setModuleBoolConfig(m, v115, v88, l3)
	mBase = m.M
	v117 = m.ExcPending
	if v117 != 0 {
		goto L32
	} else {
		goto L43
	}
L43:
	;
	return v116
L44:
	;
	v127 = int32(1)
	goto L46
L45:
	;
	v127 = int32(2)
	goto L46
L46:
	;
	return v127
}
func F_bulkStringCallback(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v8 < int32(0) {
		v32 = v7
	} else {
		v13 = l0 + v8<<(uint(int32(2))%32)
		v14 = *(*int32)(unsafe.Add(mBase, uint32(v13)+16))
		v15 = int32(1)
		v16 = v14 + v15
		*(*int32)(unsafe.Add(mBase, uint32(v13)+16)) = v16
		v18 = *(*int32)(unsafe.Add(mBase, uint32(v13)+1040))
		if v18 != v15 {
			v32 = v7
		} else {
			v23 = *(*int32)(unsafe.Add(mBase, uint32(v7)+8))
			*(*int32)(unsafe.Add(mBase, uint32(v23)+8)) = int32(3)
			*(*float64)(unsafe.Add(mBase, uint32(v23))) = base.F64_convert_i32_u(v16)
			v27 = *(*int32)(unsafe.Add(mBase, uint32(v7)+8))
			*(*int32)(unsafe.Add(mBase, uint32(v7)+8)) = v27 + int32(16)
			v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			v32 = v31
		}
	}
	v36 = F_lua_checkstack(m, v32, int32(1))
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		return
	} else {
		if v36 != 0 {
			F_lua_pushlstring(m, v32, l1, l2)
			mBase = m.M
			v42 = m.ExcPending
			if v42 != 0 {
				return
			} else {
				F_processCollectionElementEnd(m, l0)
				mBase = m.M
				v44 = m.ExcPending
				if v44 != 0 {
					return
				} else {
					return
				}
			}
		} else {
			F__serverPanic_2(m, int32(915))
			mBase = m.M
			v40 = m.ExcPending
			if v40 != 0 {
				return
			} else {
				base.Wasm_trap_unreachable()
				for {
				}
			}
		}
	}
}
func F_bzmpopGetKeys(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	v9 = F_genericGetKeys(m, int32(0), int32(2), int32(3), int32(1), l1, l2, l3)
	v12 = m.ExcPending
	if v12 != 0 {
		return int32(0)
	} else {
		return v9
	}
}
