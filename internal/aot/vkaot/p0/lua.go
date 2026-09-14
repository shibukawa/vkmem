package p0

import (
	base "github.com/shibukawa/valkeymem/internal/aot/vkaot/base"
	"unsafe"
)

func F_luaArgsToServerArgv(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v53 int32
	_ = v53
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v82 int32
	_ = v82
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v114 int32
	_ = v114
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v122 int32
	_ = v122
	var v125 int32
	_ = v125
	var v128 float64
	_ = v128
	var v129 int32
	_ = v129
	var v130 float64
	_ = v130
	var v137 int64
	_ = v137
	var v139 int64
	_ = v139
	var v142 int32
	_ = v142
	var v147 int32
	_ = v147
	var v153 int64
	_ = v153
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v166 int32
	_ = v166
	var v176 int64
	_ = v176
	var v197 int32
	_ = v197
	var v202 int32
	_ = v202
	var v209 int32
	_ = v209
	var v214 int32
	_ = v214
	var v216 int32
	_ = v216
	var v220 int64
	_ = v220
	var v227 int32
	_ = v227
	var v236 int32
	_ = v236
	var v239 int32
	_ = v239
	var v242 int32
	_ = v242
	var v245 int32
	_ = v245
	var v254 int32
	_ = v254
	var v259 int64
	_ = v259
	var v269 int32
	_ = v269
	var v272 int64
	_ = v272
	var v273 int64
	_ = v273
	var v281 int32
	_ = v281
	var v284 int32
	_ = v284
	var v293 int32
	_ = v293
	var v298 int64
	_ = v298
	var v310 int32
	_ = v310
	var v315 int32
	_ = v315
	var v322 int32
	_ = v322
	var v328 int32
	_ = v328
	var v350 int32
	_ = v350
	var v352 int32
	_ = v352
	var v356 int32
	_ = v356
	var v357 int32
	_ = v357
	var v360 int32
	_ = v360
	var v361 int32
	_ = v361
	var v368 int32
	_ = v368
	var v387 int32
	_ = v387
	var v398 int32
	_ = v398
	var v402 int32
	_ = v402
	var v403 int32
	_ = v403
	var v404 int32
	_ = v404
	var v405 int32
	_ = v405
	var v407 int32
	_ = v407
	var v413 int32
	_ = v413
	var v415 int32
	_ = v415
	var v428 int32
	_ = v428
	var v432 int32
	_ = v432
	var v433 int32
	_ = v433
	var v436 int32
	_ = v436
	var v440 int32
	_ = v440
	var v444 int32
	_ = v444
	var v446 int32
	_ = v446
	var v453 int32
	_ = v453
	var v456 int32
	_ = v456
	var v461 int32
	_ = v461
	var v462 int32
	_ = v462
	var v468 int32
	_ = v468
	var v476 int32
	_ = v476
	var v482 int32
	_ = v482
	var v490 int32
	_ = v490
	var v491 int32
	_ = v491
	var v492 int32
	_ = v492
	var v493 int32
	_ = v493
	var v495 int32
	_ = v495
	var v498 int32
	_ = v498
	var v499 int32
	_ = v499
	var v501 int32
	_ = v501
	var v504 int32
	_ = v504
	var v505 int32
	_ = v505
	var v507 int32
	_ = v507
	var v510 int32
	_ = v510
	var v511 int32
	_ = v511
	var v513 int32
	_ = v513
	var v514 int32
	_ = v514
	var v515 int32
	_ = v515
	var v517 int32
	_ = v517
	var v525 int32
	_ = v525
	var v545 int32
	_ = v545
	var v547 int32
	_ = v547
	var v560 int32
	_ = v560
	var v561 int32
	_ = v561
	var v562 int32
	_ = v562
	var v564 int32
	_ = v564
	var v565 int32
	_ = v565
	var v568 int32
	_ = v568
	var v588 int32
	_ = v588
	var v589 int32
	_ = v589
	var v591 int32
	_ = v591
	var v592 int32
	_ = v592
	var v596 int32
	_ = v596
	var v621 int32
	_ = v621
	v19 = m.G0
	v21 = v19 - int32(80)
	m.G0 = v21
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v27 = (v23 - v24) >> (uint(int32(4)) % 32)
	goto L1
L1:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v27
	if v27 != 0 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	m.G0 = v21 + int32(80)
	return v621
L3:
	;
	v621 = int32(0)
	goto L2
L4:
	;
	v38 = m.G22
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v38)))
	v40 = m.T0[v39].(func(*base.Module, int32) int32)(m, v27<<(uint(int32(2))%32))
	mBase = m.M
	v41 = m.ExcPending
	if v41 != 0 {
		goto L6
	} else {
		goto L8
	}
L5:
	;
	v29 = m.G3
	F_luaPushErrorBuff(m, l1, v29+int32(_a2224))
	mBase = m.M
	v35 = m.ExcPending
	if v35 != 0 {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	return int32(0)
L7:
	;
	goto L3
L8:
	;
	v42 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	if int32(1) <= v42 {
		goto L10
	} else {
		goto L11
	}
L9:
	;
	v428 = v415 ^ int32(-1)
	if v428 < int32(0) {
		goto L93
	} else {
		goto L94
	}
L10:
	;
	v53 = int32(0)
	goto L12
L11:
	;
	v413 = int32(0)
	v415 = v42
	goto L9
L12:
	;
	v67 = int32(1)
	v68 = v53 + v67
	if v68 < v67 {
		goto L19
	} else {
		goto L20
	}
L13:
	;
	v413 = v68
	v415 = v407
	goto L9
L14:
	;
	v402 = m.G13
	v403 = *(*int32)(unsafe.Add(mBase, uint32(v402)))
	v404 = m.T0[v403].(func(*base.Module, int32, int32, int32) int32)(m, l0, v398, v387)
	mBase = m.M
	v405 = m.ExcPending
	if v405 != 0 {
		goto L6
	} else {
		goto L89
	}
L15:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+76)) = v368
	v387 = v368
	v398 = v21
	goto L14
L16:
	;
	v356 = F_lua_tolstring(m, l1, v68, v21+int32(76))
	mBase = m.M
	v357 = m.ExcPending
	if v357 != 0 {
		goto L6
	} else {
		goto L87
	}
L17:
	;
	if v125 != int32(3) {
		goto L16
	} else {
		goto L32
	}
L18:
	;
	v119 = m.G398
	if v118 != v119 {
		goto L30
	} else {
		goto L31
	}
L19:
	;
	if v68 < int32(-9999) {
		goto L22
	} else {
		goto L23
	}
L20:
	;
	v71 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v76 = v71 + v68<<(uint(int32(4))%32) + int32(-16)
	v77 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	if base.Ui32(v76) < base.Ui32(v77) {
		v118 = v76
		goto L18
	} else {
		goto L21
	}
L21:
	;
	v125 = int32(-1)
	goto L17
L22:
	;
	switch v53 + int32(10003) {
	case 0:
		goto L26
	case 1:
		goto L27
	case 2:
		goto L24
	default:
		goto L25
	}
L23:
	;
	v82 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v118 = v82 + v68<<(uint(int32(4))%32)
	goto L18
L24:
	;
	v114 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	v118 = v114 + int32(96)
	goto L18
L25:
	;
	v101 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	v102 = *(*int32)(unsafe.Add(mBase, uint32(v101)+4))
	v103 = *(*int32)(unsafe.Add(mBase, uint32(v102)))
	v104 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v103)+7)))
	if base.Ui32(v104) < base.Ui32(int32(-10002)-v68) {
		goto L28
	} else {
		goto L29
	}
L26:
	;
	v118 = l1 + int32(72)
	goto L18
L27:
	;
	v88 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	v89 = *(*int32)(unsafe.Add(mBase, uint32(v88)+4))
	v90 = *(*int32)(unsafe.Add(mBase, uint32(v89)))
	v91 = *(*int32)(unsafe.Add(mBase, uint32(v90)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l1)+96)) = int32(5)
	*(*int32)(unsafe.Add(mBase, uint32(l1)+88)) = v91
	v118 = l1 + int32(88)
	goto L18
L28:
	;
	v125 = int32(-1)
	goto L17
L29:
	;
	v118 = v103 + (int32(-10003)-v68)<<(uint(int32(4))%32) + int32(24)
	goto L18
L30:
	;
	v122 = *(*int32)(unsafe.Add(mBase, uint32(v118)+8))
	v125 = v122
	goto L17
L31:
	;
	v125 = int32(-1)
	goto L17
L32:
	;
	v128 = F_lua_tonumber(m, l1, v68)
	mBase = m.M
	v129 = m.ExcPending
	if v129 != 0 {
		goto L6
	} else {
		goto L34
	}
L33:
	;
	v350 = F_fpconv_dtoa(m, v128, v21)
	mBase = m.M
	v352 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v21+v350))) = uint8(v352)
	v368 = v350
	goto L15
L34:
	;
	v130 = base.F64_abs(v128)
	if base.F64_gt(v130, float64(4.611686018427388e+18)) != 0 {
		goto L33
	} else {
		goto L35
	}
L35:
	;
	if base.F64_lt(v130, float64(9.223372036854776e+18)) == int32(0) {
		goto L37
	} else {
		goto L38
	}
L36:
	;
	if base.F64_ne(v128, base.F64_convert_i64_s(v139)) != 0 {
		goto L33
	} else {
		goto L39
	}
L37:
	;
	v139 = int64(-9223372036854775807 - 1)
	goto L36
L38:
	;
	v137 = base.I64_trunc_f64_s(v128)
	v139 = v137
	goto L36
L39:
	;
	v142 = int32(0)
	if v139 <= int64(-1) {
		goto L41
	} else {
		goto L42
	}
L40:
	;
	v157 = int32(1)
	if base.Ui64(v153) < base.Ui64(int64(10)) {
		v227 = v142
		v236 = v157
		goto L43
	} else {
		goto L44
	}
L41:
	;
	v147 = int32(45)
	*(*uint8)(unsafe.Add(mBase, uint32(v21))) = uint8(v147)
	v153 = int64(0) - v139
	v154 = v21 | int32(1)
	v155 = int32(63)
	v156 = int32(1)
	goto L40
L42:
	;
	v153 = v139
	v154 = v21
	v155 = int32(64)
	v156 = int32(0)
	goto L40
L43:
	;
	v239 = v236 + v227
	if base.Ui32(v155) <= base.Ui32(v239) {
		goto L75
	} else {
		goto L76
	}
L44:
	;
	v166 = v142
	v176 = v153
	goto L45
L45:
	;
	if base.Ui64(int64(100)) <= base.Ui64(v176) {
		goto L47
	} else {
		goto L48
	}
L46:
	;
	v227 = v216
	v236 = v157
	goto L43
L47:
	;
	if base.Ui64(int64(1000)) <= base.Ui64(v176) {
		goto L49
	} else {
		goto L50
	}
L48:
	;
	v227 = v166
	v236 = int32(2)
	goto L43
L49:
	;
	if base.Ui64(int64(999999999999)) < base.Ui64(v176) {
		goto L51
	} else {
		goto L52
	}
L50:
	;
	v227 = v166
	v236 = int32(3)
	goto L43
L51:
	;
	v216 = v166 + int32(12)
	v220 = base.I64_div_u_s(v176, int64(1000000000000))
	if base.Ui64(int64(9999999999999)) < base.Ui64(v176) {
		v166 = v216
		v176 = v220
		goto L45
	} else {
		goto L73
	}
L52:
	;
	if base.Ui64(int64(99999999)) < base.Ui64(v176) {
		goto L53
	} else {
		goto L54
	}
L53:
	;
	if base.Ui64(int64(9999999999)) < base.Ui64(v176) {
		goto L65
	} else {
		goto L66
	}
L54:
	;
	if base.Ui64(int64(999999)) < base.Ui64(v176) {
		goto L55
	} else {
		goto L56
	}
L55:
	;
	if base.Ui64(int64(9999999)) < base.Ui64(v176) {
		goto L62
	} else {
		goto L63
	}
L56:
	;
	if base.Ui64(int64(10000)) <= base.Ui64(v176) {
		goto L57
	} else {
		goto L58
	}
L57:
	;
	if base.Ui64(int64(99999)) < base.Ui64(v176) {
		goto L59
	} else {
		goto L60
	}
L58:
	;
	v227 = v166
	v236 = int32(4)
	goto L43
L59:
	;
	v197 = int32(6)
	goto L61
L60:
	;
	v197 = int32(5)
	goto L61
L61:
	;
	v227 = v166
	v236 = v197
	goto L43
L62:
	;
	v202 = int32(8)
	goto L64
L63:
	;
	v202 = int32(7)
	goto L64
L64:
	;
	v227 = v166
	v236 = v202
	goto L43
L65:
	;
	if base.Ui64(int64(99999999999)) < base.Ui64(v176) {
		goto L70
	} else {
		goto L71
	}
L66:
	;
	if base.Ui64(int64(999999999)) < base.Ui64(v176) {
		goto L67
	} else {
		goto L68
	}
L67:
	;
	v209 = int32(10)
	goto L69
L68:
	;
	v209 = int32(9)
	goto L69
L69:
	;
	v227 = v166
	v236 = v209
	goto L43
L70:
	;
	v214 = int32(12)
	goto L72
L71:
	;
	v214 = int32(11)
	goto L72
L72:
	;
	v227 = v166
	v236 = v214
	goto L43
L73:
	;
	goto L46
L74:
	;
	v368 = int32(0)
	goto L15
L75:
	;
	v328 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v21|v156))) = uint8(v328)
	goto L74
L76:
	;
	v242 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v154+v239))) = uint8(v242)
	v245 = v239 + int32(-1)
	if base.Ui64(v153) < base.Ui64(int64(100)) {
		v293 = v245
		v298 = v153
		goto L77
	} else {
		goto L78
	}
L77:
	;
	if base.Ui64(int64(9)) < base.Ui64(v298) {
		goto L83
	} else {
		goto L84
	}
L78:
	;
	v254 = v245
	v259 = v153
	goto L79
L79:
	;
	v269 = m.G3
	v272 = int64(100)
	v273 = base.I64_div_u_s(v259, v272)
	v281 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v269+int32(_a2225)+base.I32_wrap_i64(v259-v273*v272)<<(uint(int32(1))%32)))))
	*(*uint16)(unsafe.Add(mBase, uint32(v154+v254+int32(-1)))) = uint16(v281)
	v284 = v254 + int32(-2)
	if base.Ui64(int64(9999)) < base.Ui64(v259) {
		v254 = v284
		v259 = v273
		goto L79
	} else {
		goto L81
	}
L80:
	;
	v293 = v284
	v298 = v273
	goto L77
L81:
	;
	goto L80
L82:
	;
	if v239 == int32(0) {
		goto L74
	} else {
		goto L85
	}
L83:
	;
	v315 = m.G3
	v322 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v315+int32(_a2225)+base.I32_wrap_i64(v298)<<(uint(int32(1))%32)))))
	*(*uint16)(unsafe.Add(mBase, uint32(v154+v293+int32(-1)))) = uint16(v322)
	goto L82
L84:
	;
	v310 = base.I32_wrap_i64(v298) | int32(48)
	*(*uint8)(unsafe.Add(mBase, uint32(v154+v293))) = uint8(v310)
	goto L82
L85:
	;
	v368 = v239 + v156
	goto L15
L86:
	;
	v361 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	v413 = v53
	v415 = v361
	goto L9
L87:
	;
	if v356 == int32(0) {
		goto L86
	} else {
		goto L88
	}
L88:
	;
	v360 = *(*int32)(unsafe.Add(mBase, uint32(v21)+76))
	v387 = v360
	v398 = v356
	goto L14
L89:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v40+v53<<(uint(int32(2))%32)))) = v404
	v407 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	if v68 < v407 {
		v53 = v68
		goto L12
	} else {
		goto L90
	}
L90:
	;
	goto L13
L91:
	;
	v456 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	if v413 == v456 {
		v621 = v40
		goto L2
	} else {
		goto L99
	}
L92:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+8)) = v453
	goto L91
L93:
	;
	v446 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v453 = v446 + v428<<(uint(int32(4))%32) + int32(16)
	goto L92
L94:
	;
	v432 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v433 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v436 = v433 + v428<<(uint(int32(4))%32)
	if base.Ui32(v436) <= base.Ui32(v432) {
		v453 = v436
		goto L92
	} else {
		goto L95
	}
L95:
	;
	v440 = v432
	goto L96
L96:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v440)+8)) = int32(0)
	v444 = v440 + int32(16)
	if base.Ui32(v444) < base.Ui32(v436) {
		v440 = v444
		goto L96
	} else {
		goto L98
	}
L98:
	;
	v453 = v436
	goto L92
L99:
	;
	if v413 == int32(0) {
		goto L100
	} else {
		goto L101
	}
L100:
	;
	v588 = m.G11
	v589 = *(*int32)(unsafe.Add(mBase, uint32(v588)))
	m.T0[v589].(func(*base.Module, int32))(m, v40)
	mBase = m.M
	v591 = m.ExcPending
	if v591 != 0 {
		goto L6
	} else {
		goto L116
	}
L101:
	;
	v461 = v413 & int32(3)
	v462 = int32(0)
	if base.Ui32(v413) < base.Ui32(int32(4)) {
		v525 = v462
		goto L102
	} else {
		goto L103
	}
L102:
	;
	if v461 == int32(0) {
		goto L100
	} else {
		goto L111
	}
L103:
	;
	v468 = int32(0)
	v476 = v468
	v482 = v468
	goto L104
L104:
	;
	v490 = v40 + v476<<(uint(int32(2))%32)
	v491 = *(*int32)(unsafe.Add(mBase, uint32(v490)))
	v492 = m.G17
	v493 = *(*int32)(unsafe.Add(mBase, uint32(v492)))
	m.T0[v493].(func(*base.Module, int32, int32))(m, l0, v491)
	mBase = m.M
	v495 = m.ExcPending
	if v495 != 0 {
		goto L6
	} else {
		goto L106
	}
L105:
	;
	v525 = v515
	goto L102
L106:
	;
	v498 = *(*int32)(unsafe.Add(mBase, uint32(v490+int32(4))))
	v499 = *(*int32)(unsafe.Add(mBase, uint32(v492)))
	m.T0[v499].(func(*base.Module, int32, int32))(m, l0, v498)
	mBase = m.M
	v501 = m.ExcPending
	if v501 != 0 {
		goto L6
	} else {
		goto L107
	}
L107:
	;
	v504 = *(*int32)(unsafe.Add(mBase, uint32(v490+int32(8))))
	v505 = *(*int32)(unsafe.Add(mBase, uint32(v492)))
	m.T0[v505].(func(*base.Module, int32, int32))(m, l0, v504)
	mBase = m.M
	v507 = m.ExcPending
	if v507 != 0 {
		goto L6
	} else {
		goto L108
	}
L108:
	;
	v510 = *(*int32)(unsafe.Add(mBase, uint32(v490+int32(12))))
	v511 = *(*int32)(unsafe.Add(mBase, uint32(v492)))
	m.T0[v511].(func(*base.Module, int32, int32))(m, l0, v510)
	mBase = m.M
	v513 = m.ExcPending
	if v513 != 0 {
		goto L6
	} else {
		goto L109
	}
L109:
	;
	v514 = int32(4)
	v515 = v476 + v514
	v517 = v482 + v514
	if v517 != v413&int32(-4) {
		v476 = v515
		v482 = v517
		goto L104
	} else {
		goto L110
	}
L110:
	;
	goto L105
L111:
	;
	v545 = v525
	v547 = v462
	goto L112
L112:
	;
	v560 = *(*int32)(unsafe.Add(mBase, uint32(v40+v545<<(uint(int32(2))%32))))
	v561 = m.G17
	v562 = *(*int32)(unsafe.Add(mBase, uint32(v561)))
	m.T0[v562].(func(*base.Module, int32, int32))(m, l0, v560)
	mBase = m.M
	v564 = m.ExcPending
	if v564 != 0 {
		goto L6
	} else {
		goto L114
	}
L113:
	;
	goto L100
L114:
	;
	v565 = int32(1)
	v568 = v547 + v565
	if v568 != v461 {
		v545 = v545 + v565
		v547 = v568
		goto L112
	} else {
		goto L115
	}
L115:
	;
	goto L113
L116:
	;
	v592 = m.G3
	F_luaPushErrorBuff(m, l1, v592+int32(_a2226))
	mBase = m.M
	v596 = m.ExcPending
	if v596 != 0 {
		goto L6
	} else {
		goto L117
	}
L117:
	;
	goto L3
}
func F_luaEngineCompileCode(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32) int32 {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v56 int32
	_ = v56
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	var v117 int32
	_ = v117
	var v120 int32
	_ = v120
	var v123 int32
	_ = v123
	var v126 int32
	_ = v126
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
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v142 int64
	_ = v142
	var v155 int32
	_ = v155
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v168 int32
	_ = v168
	var v174 int32
	_ = v174
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v183 int32
	_ = v183
	v10 = m.G0
	v12 = v10 - int32(16)
	m.G0 = v12
	if l2 != 0 {
		v163 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
		v164 = F_luaFunctionLibraryCreate(m, v163, l3, l5, l6, l7)
		mBase = m.M
		v165 = m.ExcPending
		if v165 != 0 {
			return int32(0)
		} else {
			v168 = v164
			m.G0 = v12 + int32(16)
			return v168
		}
	} else {
		v14 = m.G3
		v15 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
		v18 = F_luaL_loadbuffer(m, v15, l3, l4, v14+int32(_a2127))
		mBase = m.M
		v21 = m.ExcPending
		if v21 != 0 {
			return int32(0)
		} else {
			if v18 == int32(0) {
				v80 = *(*int32)(unsafe.Add(mBase, uint32(v15)+8))
				v83 = v80 + int32(-16)
				v117 = m.G398
				if v83 != v117 {
					v120 = *(*int32)(unsafe.Add(mBase, uint32(v83)+8))
					v123 = v120
				} else {
					v123 = int32(-1)
				}
				if v123 != int32(6) {
					v174 = m.G3
					v180 = m.G8
					v181 = *(*int32)(unsafe.Add(mBase, uint32(v180)))
					m.T0[v181].(func(*base.Module, int32, int32, int32))(m, v174+int32(_a2128), v174+int32(_a2119), int32(263))
					mBase = m.M
					v183 = m.ExcPending
					if v183 != 0 {
						return int32(0)
					} else {
						m.Env.Exit(m, int32(1))
						mBase = m.M
						base.Wasm_trap_unreachable()
						for {
						}
					}
				} else {
					v126 = m.G9
					v128 = F_luaL_ref(m, v15, int32(-10000))
					mBase = m.M
					v129 = m.ExcPending
					if v129 != 0 {
						return int32(0)
					} else {
						v132 = *(*int32)(unsafe.Add(mBase, uint32(v126)))
						v133 = m.T0[v132].(func(*base.Module, int32, int32) int32)(m, int32(1), int32(8))
						mBase = m.M
						v134 = m.ExcPending
						if v134 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v133)+4)) = v128
							*(*int32)(unsafe.Add(mBase, uint32(v133))) = v15
							v138 = m.G22
							v139 = *(*int32)(unsafe.Add(mBase, uint32(v138)))
							v140 = m.T0[v139].(func(*base.Module, int32) int32)(m, int32(32))
							mBase = m.M
							v141 = m.ExcPending
							if v141 != 0 {
								return int32(0)
							} else {
								v142 = int64(0)
								*(*int64)(unsafe.Add(mBase, uint32(v140))) = v142
								*(*int64)(unsafe.Add(mBase, uint32(v140)+16)) = v142
								*(*int32)(unsafe.Add(mBase, uint32(v140)+12)) = v133
								*(*int32)(unsafe.Add(mBase, uint32(v140+int32(8)))) = int32(0)
								*(*int64)(unsafe.Add(mBase, uint32(v140+int32(24)))) = v142
								v155 = int32(1)
								*(*int32)(unsafe.Add(mBase, uint32(l6))) = v155
								v159 = *(*int32)(unsafe.Add(mBase, uint32(v126)))
								v160 = m.T0[v159].(func(*base.Module, int32, int32) int32)(m, v155, int32(4))
								mBase = m.M
								v161 = m.ExcPending
								if v161 != 0 {
									return int32(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v160))) = v140
									v168 = v160
									m.G0 = v12 + int32(16)
									return v168
								}
							}
						}
					}
				}
			} else {
				v24 = m.G15
				v25 = *(*int32)(unsafe.Add(mBase, uint32(v24)))
				v26 = int32(0)
				v29 = F_lua_tolstring(m, v15, int32(-1), v26)
				mBase = m.M
				v30 = m.ExcPending
				if v30 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v12))) = v29
					v32 = m.G3
					v35 = m.T0[v25].(func(*base.Module, int32, int32, int32) int32)(m, l0, v32+int32(_a2129), v12)
					mBase = m.M
					v36 = m.ExcPending
					if v36 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(l7))) = v35
						v56 = *(*int32)(unsafe.Add(mBase, uint32(v15)+8))
						*(*int32)(unsafe.Add(mBase, uint32(v15)+8)) = v56 + int32(-16)
						v168 = v26
						m.G0 = v12 + int32(16)
						return v168
					}
				}
			}
		}
	}
}
func F_luaEngineDebuggerDisable(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	v4 = int32(0)
	v5 = m.G6
	*(*int32)(unsafe.Add(mBase, uint32(v5))) = v4
	*(*int32)(unsafe.Add(mBase, uint32(v5)+264)) = v4
	return
}
func F_luaEngineFunctionCall(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32) {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
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
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v89 int64
	_ = v89
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v111 int32
	_ = v111
	var v114 int32
	_ = v114
	var v148 int32
	_ = v148
	var v151 int32
	_ = v151
	var v154 int32
	_ = v154
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v162 int32
	_ = v162
	var v181 int32
	_ = v181
	var v191 int32
	_ = v191
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v200 int32
	_ = v200
	v10 = int32(0)
	v13 = m.G3
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l3)+12))
	v16 = *(*int32)(unsafe.Add(mBase, uint32(v15)+4))
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l1+base.B2i32(l4 != v10)<<(uint(int32(2))%32))))
	F_lua_pushstring(m, v22, v13+int32(_a2130))
	mBase = m.M
	v26 = m.ExcPending
	if v26 != 0 {
		return
	} else {
		F_lua_gettable(m, v22, int32(-10000))
		mBase = m.M
		v29 = m.ExcPending
		if v29 != 0 {
			return
		} else {
			switch int32(2) {
			case 0:
				v83 = v22 + int32(72)
			case 1:
				v56 = *(*int32)(unsafe.Add(mBase, uint32(v22)+20))
				v57 = *(*int32)(unsafe.Add(mBase, uint32(v56)+4))
				v58 = *(*int32)(unsafe.Add(mBase, uint32(v57)))
				v59 = *(*int32)(unsafe.Add(mBase, uint32(v58)+12))
				*(*int32)(unsafe.Add(mBase, uint32(v22)+96)) = int32(5)
				*(*int32)(unsafe.Add(mBase, uint32(v22)+88)) = v59
				v83 = v22 + int32(88)
			case 2:
				v53 = *(*int32)(unsafe.Add(mBase, uint32(v22)+16))
				v83 = v53 + int32(96)
			default:
				v67 = *(*int32)(unsafe.Add(mBase, uint32(v22)+20))
				v68 = *(*int32)(unsafe.Add(mBase, uint32(v67)+4))
				v69 = *(*int32)(unsafe.Add(mBase, uint32(v68)))
				v70 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v69)+7)))
				v71 = m.G398
				if base.Ui32(v70) < base.Ui32(int32(-2)) {
					v82 = v71
				} else {
					v82 = v69 + int32(-24)
				}
				v83 = v82
			}
			v86 = *(*int32)(unsafe.Add(mBase, uint32(v83)))
			v87 = F_luaH_getnum(m, v86, v16)
			mBase = m.M
			v88 = *(*int32)(unsafe.Add(mBase, uint32(v22)+8))
			v89 = *(*int64)(unsafe.Add(mBase, uint32(v87)))
			*(*int64)(unsafe.Add(mBase, uint32(v88))) = v89
			v91 = *(*int32)(unsafe.Add(mBase, uint32(v87)+8))
			*(*int32)(unsafe.Add(mBase, uint32(v88)+8)) = v91
			v93 = *(*int32)(unsafe.Add(mBase, uint32(v22)+8))
			*(*int32)(unsafe.Add(mBase, uint32(v22)+8)) = v93 + int32(16)
			v111 = *(*int32)(unsafe.Add(mBase, uint32(v22)+8))
			v114 = v111 + int32(-16)
			v148 = m.G398
			if v114 != v148 {
				v151 = *(*int32)(unsafe.Add(mBase, uint32(v114)+8))
				v154 = v151
			} else {
				v154 = int32(-1)
			}
			if v154 == int32(0) {
				v191 = m.G3
				v197 = m.G8
				v198 = *(*int32)(unsafe.Add(mBase, uint32(v197)))
				m.T0[v198].(func(*base.Module, int32, int32, int32))(m, v191+int32(_a2131), v191+int32(_a2119), int32(312))
				mBase = m.M
				v200 = m.ExcPending
				if v200 != 0 {
					return
				} else {
					m.Env.Exit(m, int32(1))
					mBase = m.M
					base.Wasm_trap_unreachable()
					for {
					}
				}
			} else {
				if l4 != 0 {
					v159 = v10
				} else {
					v157 = m.G6
					v158 = *(*int32)(unsafe.Add(mBase, uint32(v157)))
					v159 = v158
				}
				v160 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
				F_luaCallFunction(m, l0, l2, l4, v22, l5, l6, l7, l8, v159, v160)
				mBase = m.M
				v162 = m.ExcPending
				if v162 != 0 {
					return
				} else {
					v181 = *(*int32)(unsafe.Add(mBase, uint32(v22)+8))
					*(*int32)(unsafe.Add(mBase, uint32(v22)+8)) = v181 + int32(-16)
					return
				}
			}
		}
	}
}
func F_luaEngineGetMemoryInfo(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v6 int64
	_ = v6
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v24 int32
	_ = v24
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
	v6 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(l0))) = v6
	*(*int64)(unsafe.Add(mBase, uint32(l0+int32(8)))) = v6
	if l3&int32(-3) != 0 {
		v19 = int32(0)
		if base.Ui32(int32(1)) < base.Ui32(l3+int32(-1)) {
			v29 = m.G332
			v30 = *(*int32)(unsafe.Add(mBase, uint32(v29)))
			v31 = m.T0[v30].(func(*base.Module, int32) int32)(m, l2)
			mBase = m.M
			v32 = m.ExcPending
			if v32 != 0 {
				return
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v31
				return
			}
		} else {
			v24 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
			v25 = F_luaMemory(m, v24)
			mBase = m.M
			v26 = m.ExcPending
			if v26 != 0 {
				return
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v19 + v25
				v29 = m.G332
				v30 = *(*int32)(unsafe.Add(mBase, uint32(v29)))
				v31 = m.T0[v30].(func(*base.Module, int32) int32)(m, l2)
				mBase = m.M
				v32 = m.ExcPending
				if v32 != 0 {
					return
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v31
					return
				}
			}
		}
	} else {
		v15 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
		v16 = F_luaMemory(m, v15)
		mBase = m.M
		v17 = m.ExcPending
		if v17 != 0 {
			return
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v16
			v19 = v16
			if base.Ui32(int32(1)) < base.Ui32(l3+int32(-1)) {
				v29 = m.G332
				v30 = *(*int32)(unsafe.Add(mBase, uint32(v29)))
				v31 = m.T0[v30].(func(*base.Module, int32) int32)(m, l2)
				mBase = m.M
				v32 = m.ExcPending
				if v32 != 0 {
					return
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v31
					return
				}
			} else {
				v24 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
				v25 = F_luaMemory(m, v24)
				mBase = m.M
				v26 = m.ExcPending
				if v26 != 0 {
					return
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v19 + v25
					v29 = m.G332
					v30 = *(*int32)(unsafe.Add(mBase, uint32(v29)))
					v31 = m.T0[v30].(func(*base.Module, int32) int32)(m, l2)
					mBase = m.M
					v32 = m.ExcPending
					if v32 != 0 {
						return
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v31
						return
					}
				}
			}
		}
	}
}
func F_luaErrorInformationDiscard(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v3 == int32(0) {
		v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		if v10 == int32(0) {
			v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			if v17 == int32(0) {
				return
			} else {
				v20 = m.G11
				v21 = *(*int32)(unsafe.Add(mBase, uint32(v20)))
				m.T0[v21].(func(*base.Module, int32))(m, v17)
				mBase = m.M
				v23 = m.ExcPending
				if v23 != 0 {
					return
				} else {
					return
				}
			}
		} else {
			v13 = m.G11
			v14 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
			m.T0[v14].(func(*base.Module, int32))(m, v10)
			mBase = m.M
			v16 = m.ExcPending
			if v16 != 0 {
				return
			} else {
				v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				if v17 == int32(0) {
					return
				} else {
					v20 = m.G11
					v21 = *(*int32)(unsafe.Add(mBase, uint32(v20)))
					m.T0[v21].(func(*base.Module, int32))(m, v17)
					mBase = m.M
					v23 = m.ExcPending
					if v23 != 0 {
						return
					} else {
						return
					}
				}
			}
		}
	} else {
		v6 = m.G11
		v7 = *(*int32)(unsafe.Add(mBase, uint32(v6)))
		m.T0[v7].(func(*base.Module, int32))(m, v3)
		mBase = m.M
		v9 = m.ExcPending
		if v9 != 0 {
			return
		} else {
			v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			if v10 == int32(0) {
				v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				if v17 == int32(0) {
					return
				} else {
					v20 = m.G11
					v21 = *(*int32)(unsafe.Add(mBase, uint32(v20)))
					m.T0[v21].(func(*base.Module, int32))(m, v17)
					mBase = m.M
					v23 = m.ExcPending
					if v23 != 0 {
						return
					} else {
						return
					}
				}
			} else {
				v13 = m.G11
				v14 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
				m.T0[v14].(func(*base.Module, int32))(m, v10)
				mBase = m.M
				v16 = m.ExcPending
				if v16 != 0 {
					return
				} else {
					v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
					if v17 == int32(0) {
						return
					} else {
						v20 = m.G11
						v21 = *(*int32)(unsafe.Add(mBase, uint32(v20)))
						m.T0[v21].(func(*base.Module, int32))(m, v17)
						mBase = m.M
						v23 = m.ExcPending
						if v23 != 0 {
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
func F_luaExtractErrorInformation(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v71 int32
	_ = v71
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v90 int32
	_ = v90
	var v94 int32
	_ = v94
	var v109 int32
	_ = v109
	var v112 int32
	_ = v112
	var v146 int32
	_ = v146
	var v149 int32
	_ = v149
	var v156 int32
	_ = v156
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v170 int32
	_ = v170
	var v173 int32
	_ = v173
	var v177 int32
	_ = v177
	var v182 int32
	_ = v182
	var v184 int32
	_ = v184
	var v188 int32
	_ = v188
	var v194 int32
	_ = v194
	var v197 int32
	_ = v197
	var v203 int32
	_ = v203
	var v207 int32
	_ = v207
	var v209 int32
	_ = v209
	var v217 int32
	_ = v217
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v249 int32
	_ = v249
	var v260 int32
	_ = v260
	var v264 int32
	_ = v264
	var v279 int32
	_ = v279
	var v282 int32
	_ = v282
	var v316 int32
	_ = v316
	var v319 int32
	_ = v319
	var v326 int32
	_ = v326
	var v331 int32
	_ = v331
	var v332 int32
	_ = v332
	var v340 int32
	_ = v340
	var v343 int32
	_ = v343
	var v347 int32
	_ = v347
	var v352 int32
	_ = v352
	var v354 int32
	_ = v354
	var v358 int32
	_ = v358
	var v364 int32
	_ = v364
	var v367 int32
	_ = v367
	var v373 int32
	_ = v373
	var v377 int32
	_ = v377
	var v379 int32
	_ = v379
	var v387 int32
	_ = v387
	var v389 int32
	_ = v389
	var v390 int32
	_ = v390
	var v391 int32
	_ = v391
	var v392 int32
	_ = v392
	var v393 int32
	_ = v393
	var v396 int32
	_ = v396
	var v397 int32
	_ = v397
	var v419 int32
	_ = v419
	var v430 int32
	_ = v430
	var v434 int32
	_ = v434
	var v449 int32
	_ = v449
	var v452 int32
	_ = v452
	var v486 int32
	_ = v486
	var v489 int32
	_ = v489
	var v496 int32
	_ = v496
	var v501 int32
	_ = v501
	var v502 int32
	_ = v502
	var v510 int32
	_ = v510
	var v513 int32
	_ = v513
	var v517 int32
	_ = v517
	var v522 int32
	_ = v522
	var v524 int32
	_ = v524
	var v528 int32
	_ = v528
	var v534 int32
	_ = v534
	var v537 int32
	_ = v537
	var v543 int32
	_ = v543
	var v547 int32
	_ = v547
	var v549 int32
	_ = v549
	var v557 int32
	_ = v557
	var v559 int32
	_ = v559
	var v560 int32
	_ = v560
	var v561 int32
	_ = v561
	var v562 int32
	_ = v562
	var v563 int32
	_ = v563
	var v566 int32
	_ = v566
	var v567 int32
	_ = v567
	var v589 int32
	_ = v589
	var v600 int32
	_ = v600
	var v604 int32
	_ = v604
	var v619 int32
	_ = v619
	var v622 int32
	_ = v622
	var v656 int32
	_ = v656
	var v659 int32
	_ = v659
	var v662 int32
	_ = v662
	var v681 int32
	_ = v681
	var v684 int32
	_ = v684
	var v720 int32
	_ = v720
	var v721 int32
	_ = v721
	var v725 int32
	_ = v725
	var v727 int32
	_ = v727
	var v747 int32
	_ = v747
	var v757 int32
	_ = v757
	var v759 int32
	_ = v759
	var v760 int32
	_ = v760
	var v761 int32
	_ = v761
	var v762 int32
	_ = v762
	var v765 int32
	_ = v765
	var v770 int32
	_ = v770
	var v776 int64
	_ = v776
	var v778 int64
	_ = v778
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	goto L5
L1:
	;
	m.G0 = v8 + int32(16)
	return
L2:
	;
	v90 = m.G3
	F_lua_getfield(m, l0, int32(-1), v90+int32(_a1411))
	mBase = m.M
	v94 = m.ExcPending
	if v94 != 0 {
		goto L19
	} else {
		goto L22
	}
L3:
	;
	if v71 == int32(0) {
		goto L2
	} else {
		goto L18
	}
L4:
	;
	v61 = m.G398
	if v27 != v61 {
		goto L16
	} else {
		goto L17
	}
L5:
	;
	goto L9
L9:
	;
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v27 = v24 + int32(-16)
	goto L4
L16:
	;
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v27)+8))
	v71 = base.B2i32(base.Ui32(v64+int32(-3)) < base.Ui32(int32(2)))
	goto L3
L17:
	;
	v71 = int32(0)
	goto L3
L18:
	;
	v76 = F_lua_tolstring(m, l0, int32(-1), int32(0))
	mBase = m.M
	v77 = m.ExcPending
	if v77 != 0 {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	return
L20:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8))) = v76
	v79 = m.G3
	v82 = F_lm_asprintf(m, v79+int32(_a2227), v8)
	mBase = m.M
	v83 = m.ExcPending
	if v83 != 0 {
		goto L19
	} else {
		goto L21
	}
L21:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v82
	*(*int32)(unsafe.Add(mBase, uint32(l1)+12)) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(l1)+4)) = int64(0)
	goto L1
L22:
	;
	goto L26
L23:
	;
	goto L63
L24:
	;
	if v156 == int32(0) {
		goto L23
	} else {
		goto L39
	}
L25:
	;
	v146 = m.G398
	if v112 != v146 {
		goto L37
	} else {
		goto L38
	}
L26:
	;
	goto L30
L30:
	;
	v109 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v112 = v109 + int32(-16)
	goto L25
L37:
	;
	v149 = *(*int32)(unsafe.Add(mBase, uint32(v112)+8))
	v156 = base.B2i32(base.Ui32(v149+int32(-3)) < base.Ui32(int32(2)))
	goto L24
L38:
	;
	v156 = int32(0)
	goto L24
L39:
	;
	v161 = F_lua_tolstring(m, l0, int32(-1), int32(0))
	mBase = m.M
	v162 = m.ExcPending
	if v162 != 0 {
		goto L19
	} else {
		goto L40
	}
L40:
	;
	if v161&int32(3) == int32(0) {
		v184 = v161
		goto L43
	} else {
		goto L44
	}
L41:
	;
	v219 = v217 + int32(1)
	v220 = m.G22
	v221 = *(*int32)(unsafe.Add(mBase, uint32(v220)))
	v222 = m.T0[v221].(func(*base.Module, int32) int32)(m, v219)
	mBase = m.M
	v223 = m.ExcPending
	if v223 != 0 {
		goto L19
	} else {
		goto L57
	}
L42:
	;
	v217 = v209 - v161
	goto L41
L43:
	;
	v188 = v184
	goto L51
L44:
	;
	v170 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v161))))
	if v170 != 0 {
		goto L45
	} else {
		goto L46
	}
L45:
	;
	v173 = v161
	goto L47
L46:
	;
	v217 = v161 - v161
	goto L41
L47:
	;
	v177 = v173 + int32(1)
	if v177&int32(3) == int32(0) {
		v184 = v177
		goto L43
	} else {
		goto L49
	}
L49:
	;
	v182 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v177))))
	if v182 != 0 {
		v173 = v177
		goto L47
	} else {
		goto L50
	}
L50:
	;
	v209 = v177
	goto L42
L51:
	;
	v194 = *(*int32)(unsafe.Add(mBase, uint32(v188)))
	v197 = int32(-2139062144)
	if (int32(16843008)-v194|v194)&v197 == v197 {
		v188 = v188 + int32(4)
		goto L51
	} else {
		goto L53
	}
L52:
	;
	v203 = v188
	goto L54
L53:
	;
	goto L52
L54:
	;
	v207 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v203))))
	if v207 != 0 {
		v203 = v203 + int32(1)
		goto L54
	} else {
		goto L56
	}
L55:
	;
	v209 = v203
	goto L42
L56:
	;
	goto L55
L57:
	;
	if v219 == int32(0) {
		v227 = v222
		goto L59
	} else {
		goto L60
	}
L58:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v227
	goto L23
L59:
	;
	goto L58
L60:
	;
	v226 = F__emscripten_memcpy_bulkmem(m, v222, v161, v219)
	mBase = m.M
	v227 = v226
	goto L59
L61:
	;
	v260 = m.G3
	F_lua_getfield(m, l0, int32(-1), v260+int32(_a2228))
	mBase = m.M
	v264 = m.ExcPending
	if v264 != 0 {
		goto L19
	} else {
		goto L69
	}
L62:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v249 + int32(-16)
	goto L61
L63:
	;
	v249 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	goto L62
L69:
	;
	goto L73
L70:
	;
	goto L110
L71:
	;
	if v326 == int32(0) {
		goto L70
	} else {
		goto L86
	}
L72:
	;
	v316 = m.G398
	if v282 != v316 {
		goto L84
	} else {
		goto L85
	}
L73:
	;
	goto L77
L77:
	;
	v279 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v282 = v279 + int32(-16)
	goto L72
L84:
	;
	v319 = *(*int32)(unsafe.Add(mBase, uint32(v282)+8))
	v326 = base.B2i32(base.Ui32(v319+int32(-3)) < base.Ui32(int32(2)))
	goto L71
L85:
	;
	v326 = int32(0)
	goto L71
L86:
	;
	v331 = F_lua_tolstring(m, l0, int32(-1), int32(0))
	mBase = m.M
	v332 = m.ExcPending
	if v332 != 0 {
		goto L19
	} else {
		goto L87
	}
L87:
	;
	if v331&int32(3) == int32(0) {
		v354 = v331
		goto L90
	} else {
		goto L91
	}
L88:
	;
	v389 = v387 + int32(1)
	v390 = m.G22
	v391 = *(*int32)(unsafe.Add(mBase, uint32(v390)))
	v392 = m.T0[v391].(func(*base.Module, int32) int32)(m, v389)
	mBase = m.M
	v393 = m.ExcPending
	if v393 != 0 {
		goto L19
	} else {
		goto L104
	}
L89:
	;
	v387 = v379 - v331
	goto L88
L90:
	;
	v358 = v354
	goto L98
L91:
	;
	v340 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v331))))
	if v340 != 0 {
		goto L92
	} else {
		goto L93
	}
L92:
	;
	v343 = v331
	goto L94
L93:
	;
	v387 = v331 - v331
	goto L88
L94:
	;
	v347 = v343 + int32(1)
	if v347&int32(3) == int32(0) {
		v354 = v347
		goto L90
	} else {
		goto L96
	}
L96:
	;
	v352 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v347))))
	if v352 != 0 {
		v343 = v347
		goto L94
	} else {
		goto L97
	}
L97:
	;
	v379 = v347
	goto L89
L98:
	;
	v364 = *(*int32)(unsafe.Add(mBase, uint32(v358)))
	v367 = int32(-2139062144)
	if (int32(16843008)-v364|v364)&v367 == v367 {
		v358 = v358 + int32(4)
		goto L98
	} else {
		goto L100
	}
L99:
	;
	v373 = v358
	goto L101
L100:
	;
	goto L99
L101:
	;
	v377 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v373))))
	if v377 != 0 {
		v373 = v373 + int32(1)
		goto L101
	} else {
		goto L103
	}
L102:
	;
	v379 = v373
	goto L89
L103:
	;
	goto L102
L104:
	;
	if v389 == int32(0) {
		v397 = v392
		goto L106
	} else {
		goto L107
	}
L105:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v397
	goto L70
L106:
	;
	goto L105
L107:
	;
	v396 = F__emscripten_memcpy_bulkmem(m, v392, v331, v389)
	mBase = m.M
	v397 = v396
	goto L106
L108:
	;
	v430 = m.G3
	F_lua_getfield(m, l0, int32(-1), v430+int32(_a2229))
	mBase = m.M
	v434 = m.ExcPending
	if v434 != 0 {
		goto L19
	} else {
		goto L116
	}
L109:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v419 + int32(-16)
	goto L108
L110:
	;
	v419 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	goto L109
L116:
	;
	goto L120
L117:
	;
	goto L157
L118:
	;
	if v496 == int32(0) {
		goto L117
	} else {
		goto L133
	}
L119:
	;
	v486 = m.G398
	if v452 != v486 {
		goto L131
	} else {
		goto L132
	}
L120:
	;
	goto L124
L124:
	;
	v449 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v452 = v449 + int32(-16)
	goto L119
L131:
	;
	v489 = *(*int32)(unsafe.Add(mBase, uint32(v452)+8))
	v496 = base.B2i32(base.Ui32(v489+int32(-3)) < base.Ui32(int32(2)))
	goto L118
L132:
	;
	v496 = int32(0)
	goto L118
L133:
	;
	v501 = F_lua_tolstring(m, l0, int32(-1), int32(0))
	mBase = m.M
	v502 = m.ExcPending
	if v502 != 0 {
		goto L19
	} else {
		goto L134
	}
L134:
	;
	if v501&int32(3) == int32(0) {
		v524 = v501
		goto L137
	} else {
		goto L138
	}
L135:
	;
	v559 = v557 + int32(1)
	v560 = m.G22
	v561 = *(*int32)(unsafe.Add(mBase, uint32(v560)))
	v562 = m.T0[v561].(func(*base.Module, int32) int32)(m, v559)
	mBase = m.M
	v563 = m.ExcPending
	if v563 != 0 {
		goto L19
	} else {
		goto L151
	}
L136:
	;
	v557 = v549 - v501
	goto L135
L137:
	;
	v528 = v524
	goto L145
L138:
	;
	v510 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v501))))
	if v510 != 0 {
		goto L139
	} else {
		goto L140
	}
L139:
	;
	v513 = v501
	goto L141
L140:
	;
	v557 = v501 - v501
	goto L135
L141:
	;
	v517 = v513 + int32(1)
	if v517&int32(3) == int32(0) {
		v524 = v517
		goto L137
	} else {
		goto L143
	}
L143:
	;
	v522 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v517))))
	if v522 != 0 {
		v513 = v517
		goto L141
	} else {
		goto L144
	}
L144:
	;
	v549 = v517
	goto L136
L145:
	;
	v534 = *(*int32)(unsafe.Add(mBase, uint32(v528)))
	v537 = int32(-2139062144)
	if (int32(16843008)-v534|v534)&v537 == v537 {
		v528 = v528 + int32(4)
		goto L145
	} else {
		goto L147
	}
L146:
	;
	v543 = v528
	goto L148
L147:
	;
	goto L146
L148:
	;
	v547 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v543))))
	if v547 != 0 {
		v543 = v543 + int32(1)
		goto L148
	} else {
		goto L150
	}
L149:
	;
	v549 = v543
	goto L136
L150:
	;
	goto L149
L151:
	;
	if v559 == int32(0) {
		v567 = v562
		goto L153
	} else {
		goto L154
	}
L152:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+8)) = v567
	goto L117
L153:
	;
	goto L152
L154:
	;
	v566 = F__emscripten_memcpy_bulkmem(m, v562, v501, v559)
	mBase = m.M
	v567 = v566
	goto L153
L155:
	;
	v600 = m.G3
	F_lua_getfield(m, l0, int32(-1), v600+int32(_a2230))
	mBase = m.M
	v604 = m.ExcPending
	if v604 != 0 {
		goto L19
	} else {
		goto L163
	}
L156:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v589 + int32(-16)
	goto L155
L157:
	;
	v589 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	goto L156
L163:
	;
	goto L167
L164:
	;
	goto L202
L165:
	;
	if v662 != int32(1) {
		goto L164
	} else {
		goto L180
	}
L166:
	;
	v656 = m.G398
	if v622 != v656 {
		goto L178
	} else {
		goto L179
	}
L167:
	;
	goto L171
L171:
	;
	v619 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v622 = v619 + int32(-16)
	goto L166
L178:
	;
	v659 = *(*int32)(unsafe.Add(mBase, uint32(v622)+8))
	v662 = v659
	goto L165
L179:
	;
	v662 = int32(-1)
	goto L165
L180:
	;
	goto L183
L181:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+12)) = v727
	goto L164
L182:
	;
	v720 = *(*int32)(unsafe.Add(mBase, uint32(v684)+8))
	switch v720 {
	case 0:
		v725 = v720
		goto L197
	case 1:
		goto L199
	default:
		goto L198
	}
L183:
	;
	goto L189
L189:
	;
	v681 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v684 = v681 + int32(-16)
	goto L182
L197:
	;
	v727 = v725
	goto L181
L198:
	;
	v725 = int32(1)
	goto L197
L199:
	;
	v721 = *(*int32)(unsafe.Add(mBase, uint32(v684)))
	v727 = base.B2i32(v721 != int32(0))
	goto L181
L200:
	;
	v757 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	if v757 != 0 {
		goto L1
	} else {
		goto L208
	}
L201:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v747 + int32(-16)
	goto L200
L202:
	;
	v747 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	goto L201
L208:
	;
	v759 = m.G22
	v760 = *(*int32)(unsafe.Add(mBase, uint32(v759)))
	v761 = m.T0[v760].(func(*base.Module, int32) int32)(m, int32(18))
	mBase = m.M
	v762 = m.ExcPending
	if v762 != 0 {
		goto L19
	} else {
		goto L209
	}
L209:
	;
	v765 = m.G3
	v770 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v765)+uint32(_consts[1197]))))
	*(*uint16)(unsafe.Add(mBase, uint32(v761+int32(16)))) = uint16(v770)
	v776 = *(*int64)(unsafe.Add(mBase, uint32(v765)+uint32(_consts[1198])))
	*(*int64)(unsafe.Add(mBase, uint32(v761+int32(8)))) = v776
	v778 = *(*int64)(unsafe.Add(mBase, uint32(v765)+uint32(_consts[1199])))
	*(*int64)(unsafe.Add(mBase, uint32(v761))) = v778
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v761
	goto L1
}
func F_luaFunctionRegisterFunction(m *base.Module, l0 int32) int32 {
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
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
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
	var v46 int32
	_ = v46
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v57 int32
	_ = v57
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v105 int32
	_ = v105
	var v108 int32
	_ = v108
	var v111 int32
	_ = v111
	var v117 int32
	_ = v117
	var v119 int32
	_ = v119
	var v125 int32
	_ = v125
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v132 int32
	_ = v132
	var v134 int32
	_ = v134
	var v142 int64
	_ = v142
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v149 int32
	_ = v149
	var v164 int32
	_ = v164
	var v167 int32
	_ = v167
	var v201 int32
	_ = v201
	var v204 int32
	_ = v204
	var v211 int32
	_ = v211
	var v214 int32
	_ = v214
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v220 int32
	_ = v220
	var v223 int32
	_ = v223
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v229 int32
	_ = v229
	var v233 int32
	_ = v233
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
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
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v259 int32
	_ = v259
	var v274 int32
	_ = v274
	var v277 int32
	_ = v277
	var v311 int32
	_ = v311
	var v314 int32
	_ = v314
	var v321 int32
	_ = v321
	var v324 int32
	_ = v324
	var v325 int32
	_ = v325
	var v330 int32
	_ = v330
	var v331 int32
	_ = v331
	var v332 int32
	_ = v332
	var v333 int32
	_ = v333
	var v334 int32
	_ = v334
	var v335 int32
	_ = v335
	var v338 int32
	_ = v338
	var v340 int32
	_ = v340
	var v343 int32
	_ = v343
	var v345 int32
	_ = v345
	var v346 int32
	_ = v346
	var v347 int32
	_ = v347
	var v349 int32
	_ = v349
	var v353 int32
	_ = v353
	var v355 int32
	_ = v355
	var v356 int32
	_ = v356
	var v357 int32
	_ = v357
	var v359 int32
	_ = v359
	var v360 int32
	_ = v360
	var v361 int32
	_ = v361
	var v362 int32
	_ = v362
	var v367 int32
	_ = v367
	var v368 int32
	_ = v368
	var v371 int32
	_ = v371
	var v372 int32
	_ = v372
	var v375 int32
	_ = v375
	var v376 int32
	_ = v376
	var v377 int32
	_ = v377
	var v379 int32
	_ = v379
	var v394 int32
	_ = v394
	var v397 int32
	_ = v397
	var v431 int32
	_ = v431
	var v434 int32
	_ = v434
	var v441 int32
	_ = v441
	var v445 int32
	_ = v445
	var v446 int32
	_ = v446
	var v451 int32
	_ = v451
	var v452 int32
	_ = v452
	var v453 int32
	_ = v453
	var v454 int32
	_ = v454
	var v455 int32
	_ = v455
	var v456 int32
	_ = v456
	var v460 int32
	_ = v460
	var v462 int32
	_ = v462
	var v465 int32
	_ = v465
	var v467 int32
	_ = v467
	var v468 int32
	_ = v468
	var v469 int32
	_ = v469
	var v471 int32
	_ = v471
	var v475 int32
	_ = v475
	var v477 int32
	_ = v477
	var v478 int32
	_ = v478
	var v479 int32
	_ = v479
	var v481 int32
	_ = v481
	var v482 int32
	_ = v482
	var v483 int32
	_ = v483
	var v484 int32
	_ = v484
	var v489 int32
	_ = v489
	var v490 int32
	_ = v490
	var v493 int32
	_ = v493
	var v494 int32
	_ = v494
	var v497 int32
	_ = v497
	var v498 int32
	_ = v498
	var v499 int32
	_ = v499
	var v501 int32
	_ = v501
	var v516 int32
	_ = v516
	var v519 int32
	_ = v519
	var v553 int32
	_ = v553
	var v556 int32
	_ = v556
	var v559 int32
	_ = v559
	var v564 int32
	_ = v564
	var v566 int32
	_ = v566
	var v567 int32
	_ = v567
	var v569 int32
	_ = v569
	var v570 int32
	_ = v570
	var v571 int32
	_ = v571
	var v574 int32
	_ = v574
	var v576 int32
	_ = v576
	var v579 int32
	_ = v579
	var v581 int32
	_ = v581
	var v582 int32
	_ = v582
	var v583 int32
	_ = v583
	var v585 int32
	_ = v585
	var v589 int32
	_ = v589
	var v591 int32
	_ = v591
	var v592 int32
	_ = v592
	var v593 int32
	_ = v593
	var v595 int32
	_ = v595
	var v596 int32
	_ = v596
	var v597 int32
	_ = v597
	var v598 int32
	_ = v598
	var v603 int32
	_ = v603
	var v604 int32
	_ = v604
	var v607 int32
	_ = v607
	var v608 int32
	_ = v608
	var v611 int32
	_ = v611
	var v612 int32
	_ = v612
	var v613 int32
	_ = v613
	var v619 int32
	_ = v619
	var v634 int32
	_ = v634
	var v637 int32
	_ = v637
	var v671 int32
	_ = v671
	var v674 int32
	_ = v674
	var v677 int32
	_ = v677
	var v684 int32
	_ = v684
	var v688 int32
	_ = v688
	var v694 int32
	_ = v694
	var v709 int32
	_ = v709
	var v712 int32
	_ = v712
	var v746 int32
	_ = v746
	var v749 int32
	_ = v749
	var v752 int32
	_ = v752
	var v761 int32
	_ = v761
	var v767 int32
	_ = v767
	var v783 int32
	_ = v783
	var v786 int32
	_ = v786
	var v820 int32
	_ = v820
	var v823 int32
	_ = v823
	var v830 int32
	_ = v830
	var v833 int32
	_ = v833
	var v838 int32
	_ = v838
	var v839 int32
	_ = v839
	var v842 int32
	_ = v842
	var v844 int32
	_ = v844
	var v845 int32
	_ = v845
	var v846 int32
	_ = v846
	var v848 int32
	_ = v848
	var v852 int32
	_ = v852
	var v854 int32
	_ = v854
	var v855 int32
	_ = v855
	var v856 int32
	_ = v856
	var v858 int32
	_ = v858
	var v859 int32
	_ = v859
	var v860 int32
	_ = v860
	var v861 int32
	_ = v861
	var v866 int32
	_ = v866
	var v867 int32
	_ = v867
	var v870 int32
	_ = v870
	var v871 int32
	_ = v871
	var v874 int32
	_ = v874
	var v875 int32
	_ = v875
	var v876 int32
	_ = v876
	var v880 int32
	_ = v880
	var v885 int32
	_ = v885
	var v887 int32
	_ = v887
	var v888 int32
	_ = v888
	var v889 int32
	_ = v889
	var v891 int32
	_ = v891
	var v895 int32
	_ = v895
	var v897 int32
	_ = v897
	var v898 int32
	_ = v898
	var v899 int32
	_ = v899
	var v901 int32
	_ = v901
	var v902 int32
	_ = v902
	var v903 int32
	_ = v903
	var v904 int32
	_ = v904
	var v909 int32
	_ = v909
	var v910 int32
	_ = v910
	var v913 int32
	_ = v913
	var v914 int32
	_ = v914
	var v917 int32
	_ = v917
	var v918 int32
	_ = v918
	var v919 int32
	_ = v919
	var v925 int32
	_ = v925
	var v930 int32
	_ = v930
	var v932 int32
	_ = v932
	var v933 int32
	_ = v933
	var v934 int32
	_ = v934
	var v936 int32
	_ = v936
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
	var v947 int32
	_ = v947
	var v948 int32
	_ = v948
	var v949 int32
	_ = v949
	var v954 int32
	_ = v954
	var v955 int32
	_ = v955
	var v958 int32
	_ = v958
	var v959 int32
	_ = v959
	var v962 int32
	_ = v962
	var v963 int32
	_ = v963
	var v964 int32
	_ = v964
	var v970 int32
	_ = v970
	var v975 int32
	_ = v975
	var v977 int32
	_ = v977
	var v978 int32
	_ = v978
	var v979 int32
	_ = v979
	var v981 int32
	_ = v981
	var v985 int32
	_ = v985
	var v987 int32
	_ = v987
	var v988 int32
	_ = v988
	var v989 int32
	_ = v989
	var v991 int32
	_ = v991
	var v992 int32
	_ = v992
	var v993 int32
	_ = v993
	var v994 int32
	_ = v994
	var v999 int32
	_ = v999
	var v1000 int32
	_ = v1000
	var v1003 int32
	_ = v1003
	var v1004 int32
	_ = v1004
	var v1007 int32
	_ = v1007
	var v1008 int32
	_ = v1008
	var v1009 int32
	_ = v1009
	var v1015 int32
	_ = v1015
	var v1020 int32
	_ = v1020
	var v1022 int32
	_ = v1022
	var v1023 int32
	_ = v1023
	var v1024 int32
	_ = v1024
	var v1026 int32
	_ = v1026
	var v1030 int32
	_ = v1030
	var v1032 int32
	_ = v1032
	var v1033 int32
	_ = v1033
	var v1034 int32
	_ = v1034
	var v1036 int32
	_ = v1036
	var v1037 int32
	_ = v1037
	var v1038 int32
	_ = v1038
	var v1039 int32
	_ = v1039
	var v1044 int32
	_ = v1044
	var v1045 int32
	_ = v1045
	var v1048 int32
	_ = v1048
	var v1049 int32
	_ = v1049
	var v1052 int32
	_ = v1052
	var v1053 int32
	_ = v1053
	var v1054 int32
	_ = v1054
	var v1060 int32
	_ = v1060
	var v1080 int32
	_ = v1080
	var v1092 int32
	_ = v1092
	var v1096 int32
	_ = v1096
	var v1102 int32
	_ = v1102
	var v1105 int32
	_ = v1105
	var v1106 int32
	_ = v1106
	var v1121 int32
	_ = v1121
	var v1124 int32
	_ = v1124
	var v1158 int32
	_ = v1158
	var v1161 int32
	_ = v1161
	var v1164 int32
	_ = v1164
	var v1186 int32
	_ = v1186
	var v1196 int32
	_ = v1196
	var v1203 int32
	_ = v1203
	var v1206 int32
	_ = v1206
	var v1219 int64
	_ = v1219
	var v1244 int32
	_ = v1244
	var v1259 int64
	_ = v1259
	var v1261 int32
	_ = v1261
	var v1262 int32
	_ = v1262
	var v1284 int32
	_ = v1284
	var v1299 int64
	_ = v1299
	var v1300 int32
	_ = v1300
	var v1301 int32
	_ = v1301
	var v1302 int32
	_ = v1302
	var v1307 int32
	_ = v1307
	var v1308 int32
	_ = v1308
	var v1309 int32
	_ = v1309
	var v1312 int32
	_ = v1312
	var v1320 int32
	_ = v1320
	var v1322 int32
	_ = v1322
	var v1323 int32
	_ = v1323
	var v1324 int32
	_ = v1324
	var v1329 int32
	_ = v1329
	var v1330 int32
	_ = v1330
	var v1332 int32
	_ = v1332
	var v1337 int32
	_ = v1337
	var v1339 int32
	_ = v1339
	var v1340 int32
	_ = v1340
	var v1348 int32
	_ = v1348
	var v1349 int32
	_ = v1349
	var v1351 int32
	_ = v1351
	var v1355 int32
	_ = v1355
	var v1357 int32
	_ = v1357
	var v1358 int32
	_ = v1358
	var v1359 int32
	_ = v1359
	var v1361 int32
	_ = v1361
	var v1363 int32
	_ = v1363
	var v1372 int32
	_ = v1372
	var v1376 int32
	_ = v1376
	var v1381 int32
	_ = v1381
	var v1382 int32
	_ = v1382
	var v1424 int32
	_ = v1424
	var v1427 int32
	_ = v1427
	var v1434 int32
	_ = v1434
	var v1437 int32
	_ = v1437
	var v1438 int32
	_ = v1438
	var v1443 int32
	_ = v1443
	var v1444 int32
	_ = v1444
	var v1445 int32
	_ = v1445
	var v1446 int32
	_ = v1446
	var v1447 int32
	_ = v1447
	var v1448 int32
	_ = v1448
	var v1454 int32
	_ = v1454
	var v1459 int32
	_ = v1459
	var v1460 int32
	_ = v1460
	var v1502 int32
	_ = v1502
	var v1505 int32
	_ = v1505
	var v1508 int32
	_ = v1508
	var v1511 int32
	_ = v1511
	var v1513 int32
	_ = v1513
	var v1514 int32
	_ = v1514
	var v1516 int32
	_ = v1516
	var v1517 int32
	_ = v1517
	var v1518 int32
	_ = v1518
	var v1521 int64
	_ = v1521
	var v1532 int32
	_ = v1532
	var v1533 int32
	_ = v1533
	var v1535 int32
	_ = v1535
	var v1536 int32
	_ = v1536
	var v1539 int32
	_ = v1539
	var v1542 int32
	_ = v1542
	var v1555 int32
	_ = v1555
	var v1556 int32
	_ = v1556
	var v1558 int32
	_ = v1558
	var v1559 int32
	_ = v1559
	var v1560 int32
	_ = v1560
	var v1573 int32
	_ = v1573
	var v1575 int32
	_ = v1575
	var v1577 int32
	_ = v1577
	v13 = m.G0
	v15 = v13 - int32(16)
	m.G0 = v15
	v17 = m.G3
	v20 = F_luaGetFromRegistry(m, l0, v17+int32(_a2133))
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	m.G0 = v15 + int32(16)
	return v1577
L2:
	;
	v33 = m.G9
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v33)))
	v35 = m.T0[v34].(func(*base.Module, int32, int32) int32)(m, int32(1), int32(32))
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
		goto L3
	} else {
		goto L8
	}
L3:
	;
	return int32(0)
L4:
	;
	if v20 != 0 {
		goto L2
	} else {
		goto L5
	}
L5:
	;
	v24 = m.G3
	F_luaPushError(m, l0, v24+int32(_a2134))
	mBase = m.M
	v28 = m.ExcPending
	if v28 != 0 {
		goto L3
	} else {
		goto L6
	}
L6:
	;
	v29 = F_luaError(m, l0)
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		goto L3
	} else {
		goto L7
	}
L7:
	;
	v1577 = v29
	goto L1
L8:
	;
	v37 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v41 = (v37 - v38) >> (uint(int32(4)) % 32)
	goto L12
L9:
	;
	v1573 = *(*int32)(unsafe.Add(mBase, uint32(v20)))
	F_list_add(m, v1573, v35)
	mBase = m.M
	v1575 = m.ExcPending
	if v1575 != 0 {
		goto L3
	} else {
		goto L418
	}
L10:
	;
	v1555 = m.G11
	v1556 = *(*int32)(unsafe.Add(mBase, uint32(v1555)))
	m.T0[v1556].(func(*base.Module, int32))(m, v35)
	mBase = m.M
	v1558 = m.ExcPending
	if v1558 != 0 {
		goto L3
	} else {
		goto L416
	}
L11:
	;
	if v41 != int32(1) {
		goto L15
	} else {
		goto L16
	}
L12:
	;
	v42 = int32(-3)
	if base.Ui32(v42) < base.Ui32(v41+v42) {
		goto L11
	} else {
		goto L13
	}
L13:
	;
	v46 = m.G3
	F_luaPushError(m, l0, v46+int32(_a2135))
	mBase = m.M
	v50 = m.ExcPending
	if v50 != 0 {
		goto L3
	} else {
		goto L14
	}
L14:
	;
	goto L10
L15:
	;
	v1372 = m.G3
	goto L378
L16:
	;
	v53 = m.G3
	goto L21
L17:
	;
	v119 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v119)+8)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v119 + int32(16)
	goto L35
L18:
	;
	if v111 == int32(5) {
		goto L17
	} else {
		goto L33
	}
L19:
	;
	v105 = m.G398
	if v62 != v105 {
		goto L31
	} else {
		goto L32
	}
L21:
	;
	v57 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v62 = v57 + int32(0)
	v63 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if base.Ui32(v62) < base.Ui32(v63) {
		goto L19
	} else {
		goto L22
	}
L22:
	;
	v111 = int32(-1)
	goto L18
L31:
	;
	v108 = *(*int32)(unsafe.Add(mBase, uint32(v62)+8))
	v111 = v108
	goto L18
L32:
	;
	v111 = int32(-1)
	goto L18
L33:
	;
	F_luaPushError(m, l0, v53+int32(_a2136))
	mBase = m.M
	v117 = m.ExcPending
	if v117 != 0 {
		goto L3
	} else {
		goto L34
	}
L34:
	;
	goto L10
L35:
	;
	v125 = m.G3
	v127 = F_lua_next(m, l0, int32(-2))
	mBase = m.M
	v128 = m.ExcPending
	if v128 != 0 {
		goto L3
	} else {
		goto L37
	}
L36:
	;
	v134 = int32(0)
	v142 = int64(0)
	v143 = v134
	v144 = v134
	v145 = v134
	goto L43
L37:
	;
	if v127 != 0 {
		goto L36
	} else {
		goto L38
	}
L38:
	;
	F_luaPushError(m, l0, v125+int32(_a2137))
	mBase = m.M
	v132 = m.ExcPending
	if v132 != 0 {
		goto L3
	} else {
		goto L39
	}
L39:
	;
	goto L10
L40:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v35)+24)) = v1299
	*(*int32)(unsafe.Add(mBase, uint32(v35)+20)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v35)+16)) = v1301
	*(*int32)(unsafe.Add(mBase, uint32(v35)+12)) = v1300
	*(*int32)(unsafe.Add(mBase, uint32(v35)+8)) = v1302
	*(*int64)(unsafe.Add(mBase, uint32(v35))) = int64(0)
	goto L9
L41:
	;
	if v1340 == int32(0) {
		goto L365
	} else {
		goto L366
	}
L42:
	;
	v1329 = m.G17
	v1330 = *(*int32)(unsafe.Add(mBase, uint32(v1329)))
	m.T0[v1330].(func(*base.Module, int32, int32))(m, int32(0), v1324)
	mBase = m.M
	v1332 = m.ExcPending
	if v1332 != 0 {
		goto L3
	} else {
		goto L364
	}
L43:
	;
	v149 = m.G3
	goto L53
L44:
	;
	v1309 = m.G3
	if v1302 != 0 {
		goto L361
	} else {
		goto L362
	}
L45:
	;
	v1307 = F_lua_next(m, l0, int32(-2))
	mBase = m.M
	v1308 = m.ExcPending
	if v1308 != 0 {
		goto L3
	} else {
		goto L359
	}
L46:
	;
	goto L353
L47:
	;
	goto L345
L48:
	;
	v1219 = int64(0)
	goto L47
L49:
	;
	if v145 == int32(0) {
		v1337 = v1203
		v1339 = v143
		v1340 = v1206
		goto L41
	} else {
		goto L342
	}
L50:
	;
	v214 = m.G3
	v217 = F_lua_tolstring(m, l0, int32(-2), int32(0))
	mBase = m.M
	v218 = m.ExcPending
	if v218 != 0 {
		goto L3
	} else {
		goto L68
	}
L51:
	;
	if v211 != 0 {
		goto L50
	} else {
		goto L66
	}
L52:
	;
	v201 = m.G398
	if v167 != v201 {
		goto L64
	} else {
		goto L65
	}
L53:
	;
	goto L57
L57:
	;
	v164 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v167 = v164 + int32(-32)
	goto L52
L64:
	;
	v204 = *(*int32)(unsafe.Add(mBase, uint32(v167)+8))
	v211 = base.B2i32(base.Ui32(v204+int32(-3)) < base.Ui32(int32(2)))
	goto L51
L65:
	;
	v211 = int32(0)
	goto L51
L66:
	;
	v1203 = v149 + int32(_a2138)
	v1206 = v144
	goto L49
L67:
	;
	v338 = m.G3
	v340 = v338 + int32(_a2139)
	v343 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v217))))
	if v343 != 0 {
		goto L105
	} else {
		goto L106
	}
L68:
	;
	v220 = v214 + int32(_a2140)
	v223 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v217))))
	if v223 != 0 {
		goto L71
	} else {
		goto L72
	}
L69:
	;
	if v255-v257 != 0 {
		goto L67
	} else {
		goto L81
	}
L70:
	;
	v255 = F_tolower(m, v251)
	mBase = m.M
	v256 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v252))))
	v257 = F_tolower(m, v256)
	mBase = m.M
	goto L69
L71:
	;
	v225 = v217
	v226 = v220
	v227 = v223
	goto L74
L72:
	;
	v251 = int32(0)
	v252 = v220
	goto L70
L73:
	;
	v251 = v248 & int32(255)
	v252 = v247
	goto L70
L74:
	;
	v229 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v226))))
	if v229 == int32(0) {
		v247 = v226
		v248 = v227
		goto L73
	} else {
		goto L76
	}
L75:
	;
	v247 = v241
	v248 = int32(0)
	goto L73
L76:
	;
	v233 = v227 & int32(255)
	if v233 == v229 {
		goto L77
	} else {
		goto L78
	}
L77:
	;
	v240 = int32(1)
	v241 = v226 + v240
	v242 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v225)+1)))
	if v242 != 0 {
		v225 = v225 + v240
		v226 = v241
		v227 = v242
		goto L74
	} else {
		goto L80
	}
L78:
	;
	v235 = F_tolower(m, v233)
	mBase = m.M
	v236 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v226))))
	v237 = F_tolower(m, v236)
	mBase = m.M
	if v235 == v237 {
		goto L77
	} else {
		goto L79
	}
L79:
	;
	v239 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v225))))
	v247 = v226
	v248 = v239
	goto L73
L80:
	;
	goto L75
L81:
	;
	v259 = m.G3
	goto L85
L82:
	;
	v324 = m.G3
	v325 = m.G13
	v330 = F_lua_tolstring(m, l0, int32(-1), v15+int32(4))
	mBase = m.M
	v331 = m.ExcPending
	if v331 != 0 {
		goto L3
	} else {
		goto L99
	}
L83:
	;
	if v321 != 0 {
		goto L82
	} else {
		goto L98
	}
L84:
	;
	v311 = m.G398
	if v277 != v311 {
		goto L96
	} else {
		goto L97
	}
L85:
	;
	goto L89
L89:
	;
	v274 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v277 = v274 + int32(-16)
	goto L84
L96:
	;
	v314 = *(*int32)(unsafe.Add(mBase, uint32(v277)+8))
	v321 = base.B2i32(base.Ui32(v314+int32(-3)) < base.Ui32(int32(2)))
	goto L83
L97:
	;
	v321 = int32(0)
	goto L83
L98:
	;
	v1337 = v259 + int32(_a2141)
	v1339 = v143
	v1340 = v144
	goto L41
L99:
	;
	v332 = *(*int32)(unsafe.Add(mBase, uint32(v15)+4))
	v333 = *(*int32)(unsafe.Add(mBase, uint32(v325)))
	v334 = m.T0[v333].(func(*base.Module, int32, int32, int32) int32)(m, int32(0), v330, v332)
	mBase = m.M
	v335 = m.ExcPending
	if v335 != 0 {
		goto L3
	} else {
		goto L100
	}
L100:
	;
	if v334 != 0 {
		v1259 = v142
		v1261 = v144
		v1262 = v334
		goto L46
	} else {
		goto L101
	}
L101:
	;
	v1337 = v324 + int32(_a2141)
	v1339 = v143
	v1340 = v144
	goto L41
L102:
	;
	v460 = m.G3
	v462 = v460 + int32(_a2142)
	v465 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v217))))
	if v465 != 0 {
		goto L139
	} else {
		goto L140
	}
L103:
	;
	if v375-v377 != 0 {
		goto L102
	} else {
		goto L115
	}
L104:
	;
	v375 = F_tolower(m, v371)
	mBase = m.M
	v376 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v372))))
	v377 = F_tolower(m, v376)
	mBase = m.M
	goto L103
L105:
	;
	v345 = v217
	v346 = v340
	v347 = v343
	goto L108
L106:
	;
	v371 = int32(0)
	v372 = v340
	goto L104
L107:
	;
	v371 = v368 & int32(255)
	v372 = v367
	goto L104
L108:
	;
	v349 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v346))))
	if v349 == int32(0) {
		v367 = v346
		v368 = v347
		goto L107
	} else {
		goto L110
	}
L109:
	;
	v367 = v361
	v368 = int32(0)
	goto L107
L110:
	;
	v353 = v347 & int32(255)
	if v353 == v349 {
		goto L111
	} else {
		goto L112
	}
L111:
	;
	v360 = int32(1)
	v361 = v346 + v360
	v362 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v345)+1)))
	if v362 != 0 {
		v345 = v345 + v360
		v346 = v361
		v347 = v362
		goto L108
	} else {
		goto L114
	}
L112:
	;
	v355 = F_tolower(m, v353)
	mBase = m.M
	v356 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v346))))
	v357 = F_tolower(m, v356)
	mBase = m.M
	if v355 == v357 {
		goto L111
	} else {
		goto L113
	}
L113:
	;
	v359 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v345))))
	v367 = v346
	v368 = v359
	goto L107
L114:
	;
	goto L109
L115:
	;
	v379 = m.G3
	goto L119
L116:
	;
	v445 = m.G3
	v446 = m.G13
	v451 = F_lua_tolstring(m, l0, int32(-1), v15+int32(8))
	mBase = m.M
	v452 = m.ExcPending
	if v452 != 0 {
		goto L3
	} else {
		goto L133
	}
L117:
	;
	if v441 != 0 {
		goto L116
	} else {
		goto L132
	}
L118:
	;
	v431 = m.G398
	if v397 != v431 {
		goto L130
	} else {
		goto L131
	}
L119:
	;
	goto L123
L123:
	;
	v394 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v397 = v394 + int32(-16)
	goto L118
L130:
	;
	v434 = *(*int32)(unsafe.Add(mBase, uint32(v397)+8))
	v441 = base.B2i32(base.Ui32(v434+int32(-3)) < base.Ui32(int32(2)))
	goto L117
L131:
	;
	v441 = int32(0)
	goto L117
L132:
	;
	v1203 = v379 + int32(_a2143)
	v1206 = int32(0)
	goto L49
L133:
	;
	v453 = *(*int32)(unsafe.Add(mBase, uint32(v15)+8))
	v454 = *(*int32)(unsafe.Add(mBase, uint32(v446)))
	v455 = m.T0[v454].(func(*base.Module, int32, int32, int32) int32)(m, int32(0), v451, v453)
	mBase = m.M
	v456 = m.ExcPending
	if v456 != 0 {
		goto L3
	} else {
		goto L134
	}
L134:
	;
	if v455 != 0 {
		v1259 = v142
		v1261 = v455
		v1262 = v145
		goto L46
	} else {
		goto L135
	}
L135:
	;
	v1203 = v445 + int32(_a2143)
	v1206 = int32(0)
	goto L49
L136:
	;
	v574 = m.G3
	v576 = v574 + int32(_a2144)
	v579 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v217))))
	if v579 != 0 {
		goto L172
	} else {
		goto L173
	}
L137:
	;
	if v497-v499 != 0 {
		goto L136
	} else {
		goto L149
	}
L138:
	;
	v497 = F_tolower(m, v493)
	mBase = m.M
	v498 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v494))))
	v499 = F_tolower(m, v498)
	mBase = m.M
	goto L137
L139:
	;
	v467 = v217
	v468 = v462
	v469 = v465
	goto L142
L140:
	;
	v493 = int32(0)
	v494 = v462
	goto L138
L141:
	;
	v493 = v490 & int32(255)
	v494 = v489
	goto L138
L142:
	;
	v471 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v468))))
	if v471 == int32(0) {
		v489 = v468
		v490 = v469
		goto L141
	} else {
		goto L144
	}
L143:
	;
	v489 = v483
	v490 = int32(0)
	goto L141
L144:
	;
	v475 = v469 & int32(255)
	if v475 == v471 {
		goto L145
	} else {
		goto L146
	}
L145:
	;
	v482 = int32(1)
	v483 = v468 + v482
	v484 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v467)+1)))
	if v484 != 0 {
		v467 = v467 + v482
		v468 = v483
		v469 = v484
		goto L142
	} else {
		goto L148
	}
L146:
	;
	v477 = F_tolower(m, v475)
	mBase = m.M
	v478 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v468))))
	v479 = F_tolower(m, v478)
	mBase = m.M
	if v477 == v479 {
		goto L145
	} else {
		goto L147
	}
L147:
	;
	v481 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v467))))
	v489 = v468
	v490 = v481
	goto L141
L148:
	;
	goto L143
L149:
	;
	v501 = m.G3
	goto L153
L150:
	;
	v564 = m.G22
	v566 = F_luaL_ref(m, l0, int32(-10000))
	mBase = m.M
	v567 = m.ExcPending
	if v567 != 0 {
		goto L3
	} else {
		goto L167
	}
L151:
	;
	if v559 == int32(6) {
		goto L150
	} else {
		goto L166
	}
L152:
	;
	v553 = m.G398
	if v519 != v553 {
		goto L164
	} else {
		goto L165
	}
L153:
	;
	goto L157
L157:
	;
	v516 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v519 = v516 + int32(-16)
	goto L152
L164:
	;
	v556 = *(*int32)(unsafe.Add(mBase, uint32(v519)+8))
	v559 = v556
	goto L151
L165:
	;
	v559 = int32(-1)
	goto L151
L166:
	;
	v1203 = v501 + int32(_a2145)
	v1206 = v144
	goto L49
L167:
	;
	v569 = *(*int32)(unsafe.Add(mBase, uint32(v564)))
	v570 = m.T0[v569].(func(*base.Module, int32) int32)(m, int32(8))
	mBase = m.M
	v571 = m.ExcPending
	if v571 != 0 {
		goto L3
	} else {
		goto L168
	}
L168:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v570)+4)) = v566
	*(*int32)(unsafe.Add(mBase, uint32(v570))) = l0
	v1299 = v142
	v1300 = v570
	v1301 = v144
	v1302 = v145
	goto L45
L169:
	;
	v619 = m.G3
	goto L186
L170:
	;
	if v611-v613 == int32(0) {
		goto L169
	} else {
		goto L182
	}
L171:
	;
	v611 = F_tolower(m, v607)
	mBase = m.M
	v612 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v608))))
	v613 = F_tolower(m, v612)
	mBase = m.M
	goto L170
L172:
	;
	v581 = v217
	v582 = v576
	v583 = v579
	goto L175
L173:
	;
	v607 = int32(0)
	v608 = v576
	goto L171
L174:
	;
	v607 = v604 & int32(255)
	v608 = v603
	goto L171
L175:
	;
	v585 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v582))))
	if v585 == int32(0) {
		v603 = v582
		v604 = v583
		goto L174
	} else {
		goto L177
	}
L176:
	;
	v603 = v597
	v604 = int32(0)
	goto L174
L177:
	;
	v589 = v583 & int32(255)
	if v589 == v585 {
		goto L178
	} else {
		goto L179
	}
L178:
	;
	v596 = int32(1)
	v597 = v582 + v596
	v598 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v581)+1)))
	if v598 != 0 {
		v581 = v581 + v596
		v582 = v597
		v583 = v598
		goto L175
	} else {
		goto L181
	}
L179:
	;
	v591 = F_tolower(m, v589)
	mBase = m.M
	v592 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v582))))
	v593 = F_tolower(m, v592)
	mBase = m.M
	if v591 == v593 {
		goto L178
	} else {
		goto L180
	}
L180:
	;
	v595 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v581))))
	v603 = v582
	v604 = v595
	goto L174
L181:
	;
	goto L176
L182:
	;
	v1203 = v574 + int32(_a2146)
	v1206 = v144
	goto L49
L183:
	;
	v684 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v684)+8)) = int32(3)
	*(*float64)(unsafe.Add(mBase, uint32(v684))) = float64(1)
	v688 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v688 + int32(16)
	goto L200
L184:
	;
	if v677 == int32(5) {
		goto L183
	} else {
		goto L199
	}
L185:
	;
	v671 = m.G398
	if v637 != v671 {
		goto L197
	} else {
		goto L198
	}
L186:
	;
	goto L190
L190:
	;
	v634 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v637 = v634 + int32(-16)
	goto L185
L197:
	;
	v674 = *(*int32)(unsafe.Add(mBase, uint32(v637)+8))
	v677 = v674
	goto L184
L198:
	;
	v677 = int32(-1)
	goto L184
L199:
	;
	v1203 = v619 + int32(_a2147)
	v1206 = v144
	goto L49
L200:
	;
	F_lua_gettable(m, l0, int32(-2))
	mBase = m.M
	v694 = m.ExcPending
	if v694 != 0 {
		goto L3
	} else {
		goto L201
	}
L201:
	;
	goto L204
L202:
	;
	if v752 == int32(0) {
		goto L48
	} else {
		goto L217
	}
L203:
	;
	v746 = m.G398
	if v712 != v746 {
		goto L215
	} else {
		goto L216
	}
L204:
	;
	goto L208
L208:
	;
	v709 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v712 = v709 + int32(-16)
	goto L203
L215:
	;
	v749 = *(*int32)(unsafe.Add(mBase, uint32(v712)+8))
	v752 = v749
	goto L202
L216:
	;
	v752 = int32(-1)
	goto L202
L217:
	;
	v761 = int32(2)
	v767 = int32(0)
	goto L219
L218:
	;
	goto L336
L219:
	;
	goto L223
L220:
	;
	v1219 = base.I64_extend_i32_s(v1106)
	goto L47
L221:
	;
	if v830 == int32(0) {
		goto L218
	} else {
		goto L236
	}
L222:
	;
	v820 = m.G398
	if v786 != v820 {
		goto L234
	} else {
		goto L235
	}
L223:
	;
	goto L227
L227:
	;
	v783 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v786 = v783 + int32(-16)
	goto L222
L234:
	;
	v823 = *(*int32)(unsafe.Add(mBase, uint32(v786)+8))
	v830 = base.B2i32(base.Ui32(v823+int32(-3)) < base.Ui32(int32(2)))
	goto L221
L235:
	;
	v830 = int32(0)
	goto L221
L236:
	;
	v833 = m.G3
	v838 = F_lua_tolstring(m, l0, int32(-1), int32(0))
	mBase = m.M
	v839 = m.ExcPending
	if v839 != 0 {
		goto L3
	} else {
		goto L239
	}
L237:
	;
	goto L310
L238:
	;
	v880 = m.G3
	v885 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v880)+uint32(_consts[1144]))))
	if v885 != 0 {
		goto L256
	} else {
		goto L257
	}
L239:
	;
	v842 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v833)+uint32(_consts[1145]))))
	if v842 != 0 {
		goto L242
	} else {
		goto L243
	}
L240:
	;
	if v874-v876 != 0 {
		goto L238
	} else {
		goto L252
	}
L241:
	;
	v874 = F_tolower(m, v870)
	mBase = m.M
	v875 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v871))))
	v876 = F_tolower(m, v875)
	mBase = m.M
	goto L240
L242:
	;
	v844 = v833 + int32(_a2148)
	v845 = v838
	v846 = v842
	goto L245
L243:
	;
	v870 = int32(0)
	v871 = v838
	goto L241
L244:
	;
	v870 = v867 & int32(255)
	v871 = v866
	goto L241
L245:
	;
	v848 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v845))))
	if v848 == int32(0) {
		v866 = v845
		v867 = v846
		goto L244
	} else {
		goto L247
	}
L246:
	;
	v866 = v860
	v867 = int32(0)
	goto L244
L247:
	;
	v852 = v846 & int32(255)
	if v852 == v848 {
		goto L248
	} else {
		goto L249
	}
L248:
	;
	v859 = int32(1)
	v860 = v845 + v859
	v861 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v844)+1)))
	if v861 != 0 {
		v844 = v844 + v859
		v845 = v860
		v846 = v861
		goto L245
	} else {
		goto L251
	}
L249:
	;
	v854 = F_tolower(m, v852)
	mBase = m.M
	v855 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v845))))
	v856 = F_tolower(m, v855)
	mBase = m.M
	if v854 == v856 {
		goto L248
	} else {
		goto L250
	}
L250:
	;
	v858 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v844))))
	v866 = v845
	v867 = v858
	goto L244
L251:
	;
	goto L246
L252:
	;
	v1060 = v833 + int32(_a2149)
	goto L237
L253:
	;
	v925 = m.G3
	v930 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v925)+uint32(_consts[1146]))))
	if v930 != 0 {
		goto L270
	} else {
		goto L271
	}
L254:
	;
	if v917-v919 != 0 {
		goto L253
	} else {
		goto L266
	}
L255:
	;
	v917 = F_tolower(m, v913)
	mBase = m.M
	v918 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v914))))
	v919 = F_tolower(m, v918)
	mBase = m.M
	goto L254
L256:
	;
	v887 = v880 + int32(_a2150)
	v888 = v838
	v889 = v885
	goto L259
L257:
	;
	v913 = int32(0)
	v914 = v838
	goto L255
L258:
	;
	v913 = v910 & int32(255)
	v914 = v909
	goto L255
L259:
	;
	v891 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v888))))
	if v891 == int32(0) {
		v909 = v888
		v910 = v889
		goto L258
	} else {
		goto L261
	}
L260:
	;
	v909 = v903
	v910 = int32(0)
	goto L258
L261:
	;
	v895 = v889 & int32(255)
	if v895 == v891 {
		goto L262
	} else {
		goto L263
	}
L262:
	;
	v902 = int32(1)
	v903 = v888 + v902
	v904 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v887)+1)))
	if v904 != 0 {
		v887 = v887 + v902
		v888 = v903
		v889 = v904
		goto L259
	} else {
		goto L265
	}
L263:
	;
	v897 = F_tolower(m, v895)
	mBase = m.M
	v898 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v888))))
	v899 = F_tolower(m, v898)
	mBase = m.M
	if v897 == v899 {
		goto L262
	} else {
		goto L264
	}
L264:
	;
	v901 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v887))))
	v909 = v888
	v910 = v901
	goto L258
L265:
	;
	goto L260
L266:
	;
	v1060 = v880 + int32(_a2151)
	goto L237
L267:
	;
	v970 = m.G3
	v975 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v970)+uint32(_consts[1147]))))
	if v975 != 0 {
		goto L284
	} else {
		goto L285
	}
L268:
	;
	if v962-v964 != 0 {
		goto L267
	} else {
		goto L280
	}
L269:
	;
	v962 = F_tolower(m, v958)
	mBase = m.M
	v963 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v959))))
	v964 = F_tolower(m, v963)
	mBase = m.M
	goto L268
L270:
	;
	v932 = v925 + int32(_a2152)
	v933 = v838
	v934 = v930
	goto L273
L271:
	;
	v958 = int32(0)
	v959 = v838
	goto L269
L272:
	;
	v958 = v955 & int32(255)
	v959 = v954
	goto L269
L273:
	;
	v936 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v933))))
	if v936 == int32(0) {
		v954 = v933
		v955 = v934
		goto L272
	} else {
		goto L275
	}
L274:
	;
	v954 = v948
	v955 = int32(0)
	goto L272
L275:
	;
	v940 = v934 & int32(255)
	if v940 == v936 {
		goto L276
	} else {
		goto L277
	}
L276:
	;
	v947 = int32(1)
	v948 = v933 + v947
	v949 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v932)+1)))
	if v949 != 0 {
		v932 = v932 + v947
		v933 = v948
		v934 = v949
		goto L273
	} else {
		goto L279
	}
L277:
	;
	v942 = F_tolower(m, v940)
	mBase = m.M
	v943 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v933))))
	v944 = F_tolower(m, v943)
	mBase = m.M
	if v942 == v944 {
		goto L276
	} else {
		goto L278
	}
L278:
	;
	v946 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v932))))
	v954 = v933
	v955 = v946
	goto L272
L279:
	;
	goto L274
L280:
	;
	v1060 = v925 + int32(_a2153)
	goto L237
L281:
	;
	v1015 = m.G3
	v1020 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1015)+uint32(_consts[1148]))))
	if v1020 != 0 {
		goto L297
	} else {
		goto L298
	}
L282:
	;
	if v1007-v1009 != 0 {
		goto L281
	} else {
		goto L294
	}
L283:
	;
	v1007 = F_tolower(m, v1003)
	mBase = m.M
	v1008 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1004))))
	v1009 = F_tolower(m, v1008)
	mBase = m.M
	goto L282
L284:
	;
	v977 = v970 + int32(_a2154)
	v978 = v838
	v979 = v975
	goto L287
L285:
	;
	v1003 = int32(0)
	v1004 = v838
	goto L283
L286:
	;
	v1003 = v1000 & int32(255)
	v1004 = v999
	goto L283
L287:
	;
	v981 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v978))))
	if v981 == int32(0) {
		v999 = v978
		v1000 = v979
		goto L286
	} else {
		goto L289
	}
L288:
	;
	v999 = v993
	v1000 = int32(0)
	goto L286
L289:
	;
	v985 = v979 & int32(255)
	if v985 == v981 {
		goto L290
	} else {
		goto L291
	}
L290:
	;
	v992 = int32(1)
	v993 = v978 + v992
	v994 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v977)+1)))
	if v994 != 0 {
		v977 = v977 + v992
		v978 = v993
		v979 = v994
		goto L287
	} else {
		goto L293
	}
L291:
	;
	v987 = F_tolower(m, v985)
	mBase = m.M
	v988 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v978))))
	v989 = F_tolower(m, v988)
	mBase = m.M
	if v987 == v989 {
		goto L290
	} else {
		goto L292
	}
L292:
	;
	v991 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v977))))
	v999 = v978
	v1000 = v991
	goto L286
L293:
	;
	goto L288
L294:
	;
	v1060 = v970 + int32(_a2155)
	goto L237
L295:
	;
	if v1052-v1054 != 0 {
		goto L218
	} else {
		goto L307
	}
L296:
	;
	v1052 = F_tolower(m, v1048)
	mBase = m.M
	v1053 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1049))))
	v1054 = F_tolower(m, v1053)
	mBase = m.M
	goto L295
L297:
	;
	v1022 = v1015 + int32(_a2156)
	v1023 = v838
	v1024 = v1020
	goto L300
L298:
	;
	v1048 = int32(0)
	v1049 = v838
	goto L296
L299:
	;
	v1048 = v1045 & int32(255)
	v1049 = v1044
	goto L296
L300:
	;
	v1026 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1023))))
	if v1026 == int32(0) {
		v1044 = v1023
		v1045 = v1024
		goto L299
	} else {
		goto L302
	}
L301:
	;
	v1044 = v1038
	v1045 = int32(0)
	goto L299
L302:
	;
	v1030 = v1024 & int32(255)
	if v1030 == v1026 {
		goto L303
	} else {
		goto L304
	}
L303:
	;
	v1037 = int32(1)
	v1038 = v1023 + v1037
	v1039 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1022)+1)))
	if v1039 != 0 {
		v1022 = v1022 + v1037
		v1023 = v1038
		v1024 = v1039
		goto L300
	} else {
		goto L306
	}
L304:
	;
	v1032 = F_tolower(m, v1030)
	mBase = m.M
	v1033 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1023))))
	v1034 = F_tolower(m, v1033)
	mBase = m.M
	if v1032 == v1034 {
		goto L303
	} else {
		goto L305
	}
L305:
	;
	v1036 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1022))))
	v1044 = v1023
	v1045 = v1036
	goto L299
L306:
	;
	goto L301
L307:
	;
	v1060 = v1015 + int32(_a2157)
	goto L237
L308:
	;
	v1092 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v1092)+8)) = int32(3)
	*(*float64)(unsafe.Add(mBase, uint32(v1092))) = base.F64_convert_i32_u(v761)
	v1096 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v1096 + int32(16)
	goto L316
L309:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v1080 + int32(-16)
	goto L308
L310:
	;
	v1080 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	goto L309
L316:
	;
	F_lua_gettable(m, l0, int32(-2))
	mBase = m.M
	v1102 = m.ExcPending
	if v1102 != 0 {
		goto L3
	} else {
		goto L317
	}
L317:
	;
	v1105 = *(*int32)(unsafe.Add(mBase, uint32(v1060)))
	v1106 = v1105 | v767
	goto L320
L318:
	;
	if v1164 != 0 {
		v761 = v761 + int32(1)
		v767 = v1106
		goto L219
	} else {
		goto L333
	}
L319:
	;
	v1158 = m.G398
	if v1124 != v1158 {
		goto L331
	} else {
		goto L332
	}
L320:
	;
	goto L324
L324:
	;
	v1121 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1124 = v1121 + int32(-16)
	goto L319
L331:
	;
	v1161 = *(*int32)(unsafe.Add(mBase, uint32(v1124)+8))
	v1164 = v1161
	goto L318
L332:
	;
	v1164 = int32(-1)
	goto L318
L333:
	;
	goto L220
L334:
	;
	v1196 = m.G3
	v1203 = v1196 + int32(_a2158)
	v1206 = v144
	goto L49
L335:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v1186 + int32(-16)
	goto L334
L336:
	;
	v1186 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	goto L335
L342:
	;
	v1320 = v1203
	v1322 = v143
	v1323 = v1206
	v1324 = v145
	goto L42
L343:
	;
	v1259 = v1219
	v1261 = v144
	v1262 = v145
	goto L46
L344:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v1244 + int32(-16)
	goto L343
L345:
	;
	v1244 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	goto L344
L351:
	;
	v1299 = v1259
	v1300 = v143
	v1301 = v1261
	v1302 = v1262
	goto L45
L352:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v1284 + int32(-16)
	goto L351
L353:
	;
	v1284 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	goto L352
L359:
	;
	if v1307 != 0 {
		v142 = v1299
		v143 = v1300
		v144 = v1301
		v145 = v1302
		goto L43
	} else {
		goto L360
	}
L360:
	;
	goto L44
L361:
	;
	v1312 = m.G3
	if v1300 != 0 {
		goto L40
	} else {
		goto L363
	}
L362:
	;
	v1337 = v1309 + int32(_a2137)
	v1339 = v1300
	v1340 = v1301
	goto L41
L363:
	;
	v1320 = v1312 + int32(_a2159)
	v1322 = int32(0)
	v1323 = v1301
	v1324 = v1302
	goto L42
L364:
	;
	v1337 = v1320
	v1339 = v1322
	v1340 = v1323
	goto L41
L365:
	;
	if v1339 == int32(0) {
		goto L368
	} else {
		goto L369
	}
L366:
	;
	v1348 = m.G17
	v1349 = *(*int32)(unsafe.Add(mBase, uint32(v1348)))
	m.T0[v1349].(func(*base.Module, int32, int32))(m, int32(0), v1340)
	mBase = m.M
	v1351 = m.ExcPending
	if v1351 != 0 {
		goto L3
	} else {
		goto L367
	}
L367:
	;
	goto L365
L368:
	;
	F_luaPushError(m, l0, v1337)
	mBase = m.M
	v1363 = m.ExcPending
	if v1363 != 0 {
		goto L3
	} else {
		goto L372
	}
L369:
	;
	v1355 = *(*int32)(unsafe.Add(mBase, uint32(v1339)+4))
	F_luaL_unref(m, l0, int32(-10000), v1355)
	mBase = m.M
	v1357 = m.ExcPending
	if v1357 != 0 {
		goto L3
	} else {
		goto L370
	}
L370:
	;
	v1358 = m.G11
	v1359 = *(*int32)(unsafe.Add(mBase, uint32(v1358)))
	m.T0[v1359].(func(*base.Module, int32))(m, v1339)
	mBase = m.M
	v1361 = m.ExcPending
	if v1361 != 0 {
		goto L3
	} else {
		goto L371
	}
L371:
	;
	goto L368
L372:
	;
	goto L10
L373:
	;
	F_luaPushError(m, l0, v1539)
	mBase = m.M
	v1542 = m.ExcPending
	if v1542 != 0 {
		goto L3
	} else {
		goto L415
	}
L374:
	;
	v1437 = m.G3
	v1438 = m.G13
	v1443 = F_lua_tolstring(m, l0, int32(1), v15+int32(12))
	mBase = m.M
	v1444 = m.ExcPending
	if v1444 != 0 {
		goto L3
	} else {
		goto L392
	}
L375:
	;
	if v1434 != 0 {
		goto L374
	} else {
		goto L390
	}
L376:
	;
	v1424 = m.G398
	if v1381 != v1424 {
		goto L388
	} else {
		goto L389
	}
L378:
	;
	v1376 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v1381 = v1376 + int32(0)
	v1382 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if base.Ui32(v1381) < base.Ui32(v1382) {
		goto L376
	} else {
		goto L379
	}
L379:
	;
	v1434 = int32(0)
	goto L375
L388:
	;
	v1427 = *(*int32)(unsafe.Add(mBase, uint32(v1381)+8))
	v1434 = base.B2i32(base.Ui32(v1427+int32(-3)) < base.Ui32(int32(2)))
	goto L375
L389:
	;
	v1434 = int32(0)
	goto L375
L390:
	;
	v1539 = v1372 + int32(_a2160)
	goto L373
L391:
	;
	goto L399
L392:
	;
	v1445 = *(*int32)(unsafe.Add(mBase, uint32(v15)+12))
	v1446 = *(*int32)(unsafe.Add(mBase, uint32(v1438)))
	v1447 = m.T0[v1446].(func(*base.Module, int32, int32, int32) int32)(m, int32(0), v1443, v1445)
	mBase = m.M
	v1448 = m.ExcPending
	if v1448 != 0 {
		goto L3
	} else {
		goto L393
	}
L393:
	;
	if v1447 != 0 {
		goto L391
	} else {
		goto L394
	}
L394:
	;
	v1539 = v1437 + int32(_a2160)
	goto L373
L395:
	;
	v1532 = m.G17
	v1533 = *(*int32)(unsafe.Add(mBase, uint32(v1532)))
	m.T0[v1533].(func(*base.Module, int32, int32))(m, int32(0), v1447)
	mBase = m.M
	v1535 = m.ExcPending
	if v1535 != 0 {
		goto L3
	} else {
		goto L414
	}
L396:
	;
	if v1508 != int32(6) {
		goto L395
	} else {
		goto L411
	}
L397:
	;
	v1502 = m.G398
	if v1459 != v1502 {
		goto L409
	} else {
		goto L410
	}
L399:
	;
	v1454 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v1459 = v1454 + int32(16)
	v1460 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if base.Ui32(v1459) < base.Ui32(v1460) {
		goto L397
	} else {
		goto L400
	}
L400:
	;
	v1508 = int32(-1)
	goto L396
L409:
	;
	v1505 = *(*int32)(unsafe.Add(mBase, uint32(v1459)+8))
	v1508 = v1505
	goto L396
L410:
	;
	v1508 = int32(-1)
	goto L396
L411:
	;
	v1511 = m.G22
	v1513 = F_luaL_ref(m, l0, int32(-10000))
	mBase = m.M
	v1514 = m.ExcPending
	if v1514 != 0 {
		goto L3
	} else {
		goto L412
	}
L412:
	;
	v1516 = *(*int32)(unsafe.Add(mBase, uint32(v1511)))
	v1517 = m.T0[v1516].(func(*base.Module, int32) int32)(m, int32(8))
	mBase = m.M
	v1518 = m.ExcPending
	if v1518 != 0 {
		goto L3
	} else {
		goto L413
	}
L413:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1517)+4)) = v1513
	*(*int32)(unsafe.Add(mBase, uint32(v1517))) = l0
	v1521 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v35)+16)) = v1521
	*(*int32)(unsafe.Add(mBase, uint32(v35)+12)) = v1517
	*(*int32)(unsafe.Add(mBase, uint32(v35)+8)) = v1447
	*(*int64)(unsafe.Add(mBase, uint32(v35))) = v1521
	*(*int64)(unsafe.Add(mBase, uint32(v35+int32(24)))) = v1521
	goto L9
L414:
	;
	v1536 = m.G3
	v1539 = v1536 + int32(_a2161)
	goto L373
L415:
	;
	goto L10
L416:
	;
	v1559 = F_luaError(m, l0)
	mBase = m.M
	v1560 = m.ExcPending
	if v1560 != 0 {
		goto L3
	} else {
		goto L417
	}
L417:
	;
	v1577 = v1559
	goto L1
L418:
	;
	v1577 = int32(0)
	goto L1
}
func F_luaNewIndexAllowList(m *base.Module, l0 int32) int32 {
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
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v91 int32
	_ = v91
	var v94 int32
	_ = v94
	var v97 int32
	_ = v97
	var v100 int32
	_ = v100
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v120 int32
	_ = v120
	var v123 int32
	_ = v123
	var v157 int32
	_ = v157
	var v160 int32
	_ = v160
	var v167 int32
	_ = v167
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
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
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
	var v237 int32
	_ = v237
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v253 int32
	_ = v253
	var v258 int32
	_ = v258
	var v259 int32
	_ = v259
	var v265 int32
	_ = v265
	var v266 int32
	_ = v266
	var v269 int32
	_ = v269
	var v270 int32
	_ = v270
	var v273 int32
	_ = v273
	var v282 int32
	_ = v282
	var v283 int32
	_ = v283
	var v289 int32
	_ = v289
	var v294 int32
	_ = v294
	var v295 int32
	_ = v295
	var v301 int32
	_ = v301
	var v302 int32
	_ = v302
	var v305 int32
	_ = v305
	var v306 int32
	_ = v306
	var v309 int32
	_ = v309
	var v318 int32
	_ = v318
	var v319 int32
	_ = v319
	var v325 int32
	_ = v325
	var v330 int32
	_ = v330
	var v331 int32
	_ = v331
	var v337 int32
	_ = v337
	var v338 int32
	_ = v338
	var v341 int32
	_ = v341
	var v342 int32
	_ = v342
	var v345 int32
	_ = v345
	var v354 int32
	_ = v354
	var v355 int32
	_ = v355
	var v361 int32
	_ = v361
	var v366 int32
	_ = v366
	var v367 int32
	_ = v367
	var v373 int32
	_ = v373
	var v374 int32
	_ = v374
	var v377 int32
	_ = v377
	var v378 int32
	_ = v378
	var v381 int32
	_ = v381
	var v390 int32
	_ = v390
	var v391 int32
	_ = v391
	var v397 int32
	_ = v397
	var v402 int32
	_ = v402
	var v403 int32
	_ = v403
	var v409 int32
	_ = v409
	var v410 int32
	_ = v410
	var v413 int32
	_ = v413
	var v414 int32
	_ = v414
	var v417 int32
	_ = v417
	var v426 int32
	_ = v426
	var v427 int32
	_ = v427
	var v433 int32
	_ = v433
	var v438 int32
	_ = v438
	var v439 int32
	_ = v439
	var v445 int32
	_ = v445
	var v446 int32
	_ = v446
	var v449 int32
	_ = v449
	var v450 int32
	_ = v450
	var v453 int32
	_ = v453
	var v462 int32
	_ = v462
	var v463 int32
	_ = v463
	var v469 int32
	_ = v469
	var v474 int32
	_ = v474
	var v475 int32
	_ = v475
	var v481 int32
	_ = v481
	var v482 int32
	_ = v482
	var v485 int32
	_ = v485
	var v486 int32
	_ = v486
	var v489 int32
	_ = v489
	var v498 int32
	_ = v498
	var v499 int32
	_ = v499
	var v505 int32
	_ = v505
	var v510 int32
	_ = v510
	var v511 int32
	_ = v511
	var v517 int32
	_ = v517
	var v518 int32
	_ = v518
	var v521 int32
	_ = v521
	var v522 int32
	_ = v522
	var v525 int32
	_ = v525
	var v534 int32
	_ = v534
	var v535 int32
	_ = v535
	var v541 int32
	_ = v541
	var v546 int32
	_ = v546
	var v547 int32
	_ = v547
	var v553 int32
	_ = v553
	var v554 int32
	_ = v554
	var v557 int32
	_ = v557
	var v558 int32
	_ = v558
	var v561 int32
	_ = v561
	var v570 int32
	_ = v570
	var v571 int32
	_ = v571
	var v577 int32
	_ = v577
	var v582 int32
	_ = v582
	var v583 int32
	_ = v583
	var v589 int32
	_ = v589
	var v590 int32
	_ = v590
	var v593 int32
	_ = v593
	var v594 int32
	_ = v594
	var v597 int32
	_ = v597
	var v606 int32
	_ = v606
	var v607 int32
	_ = v607
	var v613 int32
	_ = v613
	var v618 int32
	_ = v618
	var v619 int32
	_ = v619
	var v625 int32
	_ = v625
	var v626 int32
	_ = v626
	var v629 int32
	_ = v629
	var v630 int32
	_ = v630
	var v633 int32
	_ = v633
	var v642 int32
	_ = v642
	var v643 int32
	_ = v643
	var v649 int32
	_ = v649
	var v654 int32
	_ = v654
	var v655 int32
	_ = v655
	var v661 int32
	_ = v661
	var v662 int32
	_ = v662
	var v665 int32
	_ = v665
	var v666 int32
	_ = v666
	var v669 int32
	_ = v669
	var v678 int32
	_ = v678
	var v679 int32
	_ = v679
	var v685 int32
	_ = v685
	var v690 int32
	_ = v690
	var v691 int32
	_ = v691
	var v697 int32
	_ = v697
	var v698 int32
	_ = v698
	var v701 int32
	_ = v701
	var v702 int32
	_ = v702
	var v705 int32
	_ = v705
	var v714 int32
	_ = v714
	var v715 int32
	_ = v715
	var v721 int32
	_ = v721
	var v726 int32
	_ = v726
	var v727 int32
	_ = v727
	var v733 int32
	_ = v733
	var v734 int32
	_ = v734
	var v737 int32
	_ = v737
	var v738 int32
	_ = v738
	var v741 int32
	_ = v741
	var v750 int32
	_ = v750
	var v751 int32
	_ = v751
	var v757 int32
	_ = v757
	var v762 int32
	_ = v762
	var v763 int32
	_ = v763
	var v769 int32
	_ = v769
	var v770 int32
	_ = v770
	var v773 int32
	_ = v773
	var v774 int32
	_ = v774
	var v777 int32
	_ = v777
	var v786 int32
	_ = v786
	var v787 int32
	_ = v787
	var v793 int32
	_ = v793
	var v798 int32
	_ = v798
	var v799 int32
	_ = v799
	var v805 int32
	_ = v805
	var v806 int32
	_ = v806
	var v809 int32
	_ = v809
	var v810 int32
	_ = v810
	var v813 int32
	_ = v813
	var v822 int32
	_ = v822
	var v823 int32
	_ = v823
	var v829 int32
	_ = v829
	var v834 int32
	_ = v834
	var v835 int32
	_ = v835
	var v841 int32
	_ = v841
	var v842 int32
	_ = v842
	var v845 int32
	_ = v845
	var v846 int32
	_ = v846
	var v849 int32
	_ = v849
	var v858 int32
	_ = v858
	var v859 int32
	_ = v859
	var v865 int32
	_ = v865
	var v870 int32
	_ = v870
	var v871 int32
	_ = v871
	var v877 int32
	_ = v877
	var v878 int32
	_ = v878
	var v881 int32
	_ = v881
	var v882 int32
	_ = v882
	var v885 int32
	_ = v885
	var v894 int32
	_ = v894
	var v895 int32
	_ = v895
	var v901 int32
	_ = v901
	var v906 int32
	_ = v906
	var v907 int32
	_ = v907
	var v913 int32
	_ = v913
	var v914 int32
	_ = v914
	var v917 int32
	_ = v917
	var v918 int32
	_ = v918
	var v921 int32
	_ = v921
	var v930 int32
	_ = v930
	var v931 int32
	_ = v931
	var v937 int32
	_ = v937
	var v942 int32
	_ = v942
	var v943 int32
	_ = v943
	var v949 int32
	_ = v949
	var v950 int32
	_ = v950
	var v953 int32
	_ = v953
	var v954 int32
	_ = v954
	var v957 int32
	_ = v957
	var v966 int32
	_ = v966
	var v967 int32
	_ = v967
	var v973 int32
	_ = v973
	var v978 int32
	_ = v978
	var v979 int32
	_ = v979
	var v985 int32
	_ = v985
	var v986 int32
	_ = v986
	var v989 int32
	_ = v989
	var v990 int32
	_ = v990
	var v993 int32
	_ = v993
	var v1002 int32
	_ = v1002
	var v1003 int32
	_ = v1003
	var v1009 int32
	_ = v1009
	var v1014 int32
	_ = v1014
	var v1015 int32
	_ = v1015
	var v1021 int32
	_ = v1021
	var v1022 int32
	_ = v1022
	var v1025 int32
	_ = v1025
	var v1026 int32
	_ = v1026
	var v1029 int32
	_ = v1029
	var v1038 int32
	_ = v1038
	var v1039 int32
	_ = v1039
	var v1045 int32
	_ = v1045
	var v1050 int32
	_ = v1050
	var v1051 int32
	_ = v1051
	var v1057 int32
	_ = v1057
	var v1058 int32
	_ = v1058
	var v1061 int32
	_ = v1061
	var v1062 int32
	_ = v1062
	var v1065 int32
	_ = v1065
	var v1074 int32
	_ = v1074
	var v1075 int32
	_ = v1075
	var v1081 int32
	_ = v1081
	var v1086 int32
	_ = v1086
	var v1087 int32
	_ = v1087
	var v1093 int32
	_ = v1093
	var v1094 int32
	_ = v1094
	var v1097 int32
	_ = v1097
	var v1098 int32
	_ = v1098
	var v1101 int32
	_ = v1101
	var v1110 int32
	_ = v1110
	var v1111 int32
	_ = v1111
	var v1117 int32
	_ = v1117
	var v1122 int32
	_ = v1122
	var v1123 int32
	_ = v1123
	var v1129 int32
	_ = v1129
	var v1130 int32
	_ = v1130
	var v1133 int32
	_ = v1133
	var v1134 int32
	_ = v1134
	var v1137 int32
	_ = v1137
	var v1146 int32
	_ = v1146
	var v1147 int32
	_ = v1147
	var v1153 int32
	_ = v1153
	var v1158 int32
	_ = v1158
	var v1159 int32
	_ = v1159
	var v1165 int32
	_ = v1165
	var v1166 int32
	_ = v1166
	var v1169 int32
	_ = v1169
	var v1170 int32
	_ = v1170
	var v1173 int32
	_ = v1173
	var v1182 int32
	_ = v1182
	var v1183 int32
	_ = v1183
	var v1189 int32
	_ = v1189
	var v1194 int32
	_ = v1194
	var v1195 int32
	_ = v1195
	var v1201 int32
	_ = v1201
	var v1202 int32
	_ = v1202
	var v1205 int32
	_ = v1205
	var v1206 int32
	_ = v1206
	var v1209 int32
	_ = v1209
	var v1218 int32
	_ = v1218
	var v1219 int32
	_ = v1219
	var v1225 int32
	_ = v1225
	var v1230 int32
	_ = v1230
	var v1231 int32
	_ = v1231
	var v1237 int32
	_ = v1237
	var v1238 int32
	_ = v1238
	var v1241 int32
	_ = v1241
	var v1242 int32
	_ = v1242
	var v1245 int32
	_ = v1245
	var v1254 int32
	_ = v1254
	var v1255 int32
	_ = v1255
	var v1261 int32
	_ = v1261
	var v1266 int32
	_ = v1266
	var v1267 int32
	_ = v1267
	var v1273 int32
	_ = v1273
	var v1274 int32
	_ = v1274
	var v1277 int32
	_ = v1277
	var v1278 int32
	_ = v1278
	var v1281 int32
	_ = v1281
	var v1290 int32
	_ = v1290
	var v1291 int32
	_ = v1291
	var v1297 int32
	_ = v1297
	var v1302 int32
	_ = v1302
	var v1303 int32
	_ = v1303
	var v1309 int32
	_ = v1309
	var v1310 int32
	_ = v1310
	var v1313 int32
	_ = v1313
	var v1314 int32
	_ = v1314
	var v1317 int32
	_ = v1317
	var v1326 int32
	_ = v1326
	var v1327 int32
	_ = v1327
	var v1333 int32
	_ = v1333
	var v1338 int32
	_ = v1338
	var v1339 int32
	_ = v1339
	var v1345 int32
	_ = v1345
	var v1346 int32
	_ = v1346
	var v1349 int32
	_ = v1349
	var v1350 int32
	_ = v1350
	var v1353 int32
	_ = v1353
	var v1362 int32
	_ = v1362
	var v1363 int32
	_ = v1363
	var v1369 int32
	_ = v1369
	var v1374 int32
	_ = v1374
	var v1375 int32
	_ = v1375
	var v1381 int32
	_ = v1381
	var v1382 int32
	_ = v1382
	var v1385 int32
	_ = v1385
	var v1386 int32
	_ = v1386
	var v1389 int32
	_ = v1389
	var v1398 int32
	_ = v1398
	var v1399 int32
	_ = v1399
	var v1405 int32
	_ = v1405
	var v1410 int32
	_ = v1410
	var v1411 int32
	_ = v1411
	var v1417 int32
	_ = v1417
	var v1418 int32
	_ = v1418
	var v1421 int32
	_ = v1421
	var v1422 int32
	_ = v1422
	var v1425 int32
	_ = v1425
	var v1434 int32
	_ = v1434
	var v1435 int32
	_ = v1435
	var v1441 int32
	_ = v1441
	var v1446 int32
	_ = v1446
	var v1447 int32
	_ = v1447
	var v1453 int32
	_ = v1453
	var v1454 int32
	_ = v1454
	var v1457 int32
	_ = v1457
	var v1458 int32
	_ = v1458
	var v1461 int32
	_ = v1461
	var v1470 int32
	_ = v1470
	var v1471 int32
	_ = v1471
	var v1477 int32
	_ = v1477
	var v1482 int32
	_ = v1482
	var v1483 int32
	_ = v1483
	var v1489 int32
	_ = v1489
	var v1490 int32
	_ = v1490
	var v1493 int32
	_ = v1493
	var v1494 int32
	_ = v1494
	var v1497 int32
	_ = v1497
	var v1506 int32
	_ = v1506
	var v1507 int32
	_ = v1507
	var v1513 int32
	_ = v1513
	var v1518 int32
	_ = v1518
	var v1519 int32
	_ = v1519
	var v1525 int32
	_ = v1525
	var v1526 int32
	_ = v1526
	var v1529 int32
	_ = v1529
	var v1530 int32
	_ = v1530
	var v1533 int32
	_ = v1533
	var v1542 int32
	_ = v1542
	var v1543 int32
	_ = v1543
	var v1549 int32
	_ = v1549
	var v1554 int32
	_ = v1554
	var v1555 int32
	_ = v1555
	var v1561 int32
	_ = v1561
	var v1562 int32
	_ = v1562
	var v1565 int32
	_ = v1565
	var v1566 int32
	_ = v1566
	var v1569 int32
	_ = v1569
	var v1578 int32
	_ = v1578
	var v1579 int32
	_ = v1579
	var v1585 int32
	_ = v1585
	var v1590 int32
	_ = v1590
	var v1591 int32
	_ = v1591
	var v1597 int32
	_ = v1597
	var v1598 int32
	_ = v1598
	var v1601 int32
	_ = v1601
	var v1602 int32
	_ = v1602
	var v1605 int32
	_ = v1605
	var v1614 int32
	_ = v1614
	var v1615 int32
	_ = v1615
	var v1622 int32
	_ = v1622
	var v1624 int32
	_ = v1624
	var v1625 int32
	_ = v1625
	var v1630 int32
	_ = v1630
	var v1631 int32
	_ = v1631
	var v1637 int32
	_ = v1637
	var v1638 int32
	_ = v1638
	var v1641 int32
	_ = v1641
	var v1642 int32
	_ = v1642
	var v1645 int32
	_ = v1645
	var v1654 int32
	_ = v1654
	var v1655 int32
	_ = v1655
	var v1661 int32
	_ = v1661
	var v1666 int32
	_ = v1666
	var v1667 int32
	_ = v1667
	var v1673 int32
	_ = v1673
	var v1674 int32
	_ = v1674
	var v1677 int32
	_ = v1677
	var v1678 int32
	_ = v1678
	var v1681 int32
	_ = v1681
	var v1690 int32
	_ = v1690
	var v1691 int32
	_ = v1691
	var v1697 int32
	_ = v1697
	var v1702 int32
	_ = v1702
	var v1703 int32
	_ = v1703
	var v1709 int32
	_ = v1709
	var v1710 int32
	_ = v1710
	var v1713 int32
	_ = v1713
	var v1714 int32
	_ = v1714
	var v1717 int32
	_ = v1717
	var v1726 int32
	_ = v1726
	var v1727 int32
	_ = v1727
	var v1734 int32
	_ = v1734
	var v1736 int32
	_ = v1736
	var v1741 int32
	_ = v1741
	var v1742 int32
	_ = v1742
	var v1744 int32
	_ = v1744
	var v1747 int32
	_ = v1747
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v9 = m.G3
	v12 = F_luaGetFromRegistry(m, l0, v9+int32(_a2168))
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	goto L4
L3:
	;
	goto L11
L4:
	;
	if (v16-v17)>>(uint(int32(4))%32) == int32(3) {
		goto L3
	} else {
		goto L5
	}
L5:
	;
	v23 = m.G3
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
	v30 = m.G10
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v30)))
	m.T0[v31].(func(*base.Module, int32, int32, int32, int32))(m, v24, v23+int32(_a716), v23+int32(_a2169), int32(0))
	mBase = m.M
	v33 = m.ExcPending
	if v33 != 0 {
		goto L1
	} else {
		goto L6
	}
L6:
	;
	v37 = F_luaL_error(m, l0, v23+int32(_a2170), int32(0))
	mBase = m.M
	v38 = m.ExcPending
	if v38 != 0 {
		goto L1
	} else {
		goto L7
	}
L7:
	;
	goto L3
L8:
	;
	goto L29
L9:
	;
	if v97 == int32(5) {
		goto L8
	} else {
		goto L24
	}
L10:
	;
	v91 = m.G398
	if v57 != v91 {
		goto L22
	} else {
		goto L23
	}
L11:
	;
	goto L15
L15:
	;
	v54 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v57 = v54 + int32(-48)
	goto L10
L22:
	;
	v94 = *(*int32)(unsafe.Add(mBase, uint32(v57)+8))
	v97 = v94
	goto L9
L23:
	;
	v97 = int32(-1)
	goto L9
L24:
	;
	v100 = m.G3
	v104 = F_luaL_error(m, l0, v100+int32(_a2171), int32(0))
	mBase = m.M
	v105 = m.ExcPending
	if v105 != 0 {
		goto L1
	} else {
		goto L25
	}
L25:
	;
	goto L8
L26:
	;
	v177 = m.G3
	v182 = F_lua_tolstring(m, l0, int32(-2), int32(0))
	mBase = m.M
	v183 = m.ExcPending
	if v183 != 0 {
		goto L1
	} else {
		goto L48
	}
L27:
	;
	if v167 != 0 {
		goto L26
	} else {
		goto L42
	}
L28:
	;
	v157 = m.G398
	if v123 != v157 {
		goto L40
	} else {
		goto L41
	}
L29:
	;
	goto L33
L33:
	;
	v120 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v123 = v120 + int32(-32)
	goto L28
L40:
	;
	v160 = *(*int32)(unsafe.Add(mBase, uint32(v123)+8))
	v167 = base.B2i32(base.Ui32(v160+int32(-3)) < base.Ui32(int32(2)))
	goto L27
L41:
	;
	v167 = int32(0)
	goto L27
L42:
	;
	v169 = F_lua_isnumber(m, l0, int32(-2))
	mBase = m.M
	v170 = m.ExcPending
	if v170 != 0 {
		goto L1
	} else {
		goto L43
	}
L43:
	;
	if v169 != 0 {
		goto L26
	} else {
		goto L44
	}
L44:
	;
	v171 = m.G3
	v175 = F_luaL_error(m, l0, v171+int32(_a2172), int32(0))
	mBase = m.M
	v176 = m.ExcPending
	if v176 != 0 {
		goto L1
	} else {
		goto L45
	}
L45:
	;
	goto L26
L46:
	;
	m.G0 = v7 + int32(16)
	return int32(0)
L47:
	;
	F_lua_rawset(m, l0, int32(-3))
	mBase = m.M
	v1747 = m.ExcPending
	if v1747 != 0 {
		goto L1
	} else {
		goto L440
	}
L48:
	;
	v186 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v182))))
	v187 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v177)+uint32(_consts[1149]))))
	if v187 == int32(0) {
		v210 = v186
		v211 = v187
		goto L50
	} else {
		goto L51
	}
L49:
	;
	if v211-v210&int32(255) == int32(0) {
		goto L47
	} else {
		goto L57
	}
L50:
	;
	goto L49
L51:
	;
	if v187 != v186&int32(255) {
		v210 = v186
		v211 = v187
		goto L50
	} else {
		goto L52
	}
L52:
	;
	v193 = v177 + int32(_a2173)
	v194 = v182
	goto L53
L53:
	;
	v197 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v194)+1)))
	v198 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v193)+1)))
	if v198 == int32(0) {
		v210 = v197
		v211 = v198
		goto L50
	} else {
		goto L55
	}
L54:
	;
	v210 = v197
	v211 = v198
	goto L50
L55:
	;
	v201 = int32(1)
	if v198 == v197&int32(255) {
		v193 = v193 + v201
		v194 = v194 + v201
		goto L53
	} else {
		goto L56
	}
L56:
	;
	goto L54
L57:
	;
	v217 = m.G3
	v222 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v182))))
	v223 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v217)+uint32(_consts[1150]))))
	if v223 == int32(0) {
		v246 = v222
		v247 = v223
		goto L59
	} else {
		goto L60
	}
L58:
	;
	if v247-v246&int32(255) == int32(0) {
		goto L47
	} else {
		goto L66
	}
L59:
	;
	goto L58
L60:
	;
	if v223 != v222&int32(255) {
		v246 = v222
		v247 = v223
		goto L59
	} else {
		goto L61
	}
L61:
	;
	v229 = v217 + int32(_a2174)
	v230 = v182
	goto L62
L62:
	;
	v233 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v230)+1)))
	v234 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v229)+1)))
	if v234 == int32(0) {
		v246 = v233
		v247 = v234
		goto L59
	} else {
		goto L64
	}
L63:
	;
	v246 = v233
	v247 = v234
	goto L59
L64:
	;
	v237 = int32(1)
	if v234 == v233&int32(255) {
		v229 = v229 + v237
		v230 = v230 + v237
		goto L62
	} else {
		goto L65
	}
L65:
	;
	goto L63
L66:
	;
	v253 = m.G3
	v258 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v182))))
	v259 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v253)+uint32(_consts[1151]))))
	if v259 == int32(0) {
		v282 = v258
		v283 = v259
		goto L68
	} else {
		goto L69
	}
L67:
	;
	if v283-v282&int32(255) == int32(0) {
		goto L47
	} else {
		goto L75
	}
L68:
	;
	goto L67
L69:
	;
	if v259 != v258&int32(255) {
		v282 = v258
		v283 = v259
		goto L68
	} else {
		goto L70
	}
L70:
	;
	v265 = v253 + int32(_a2175)
	v266 = v182
	goto L71
L71:
	;
	v269 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v266)+1)))
	v270 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v265)+1)))
	if v270 == int32(0) {
		v282 = v269
		v283 = v270
		goto L68
	} else {
		goto L73
	}
L72:
	;
	v282 = v269
	v283 = v270
	goto L68
L73:
	;
	v273 = int32(1)
	if v270 == v269&int32(255) {
		v265 = v265 + v273
		v266 = v266 + v273
		goto L71
	} else {
		goto L74
	}
L74:
	;
	goto L72
L75:
	;
	v289 = m.G3
	v294 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v182))))
	v295 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v289)+uint32(_consts[1152]))))
	if v295 == int32(0) {
		v318 = v294
		v319 = v295
		goto L77
	} else {
		goto L78
	}
L76:
	;
	if v319-v318&int32(255) == int32(0) {
		goto L47
	} else {
		goto L84
	}
L77:
	;
	goto L76
L78:
	;
	if v295 != v294&int32(255) {
		v318 = v294
		v319 = v295
		goto L77
	} else {
		goto L79
	}
L79:
	;
	v301 = v289 + int32(_a2176)
	v302 = v182
	goto L80
L80:
	;
	v305 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v302)+1)))
	v306 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v301)+1)))
	if v306 == int32(0) {
		v318 = v305
		v319 = v306
		goto L77
	} else {
		goto L82
	}
L81:
	;
	v318 = v305
	v319 = v306
	goto L77
L82:
	;
	v309 = int32(1)
	if v306 == v305&int32(255) {
		v301 = v301 + v309
		v302 = v302 + v309
		goto L80
	} else {
		goto L83
	}
L83:
	;
	goto L81
L84:
	;
	v325 = m.G3
	v330 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v182))))
	v331 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v325)+uint32(_consts[1153]))))
	if v331 == int32(0) {
		v354 = v330
		v355 = v331
		goto L86
	} else {
		goto L87
	}
L85:
	;
	if v355-v354&int32(255) == int32(0) {
		goto L47
	} else {
		goto L93
	}
L86:
	;
	goto L85
L87:
	;
	if v331 != v330&int32(255) {
		v354 = v330
		v355 = v331
		goto L86
	} else {
		goto L88
	}
L88:
	;
	v337 = v325 + int32(_a2177)
	v338 = v182
	goto L89
L89:
	;
	v341 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v338)+1)))
	v342 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v337)+1)))
	if v342 == int32(0) {
		v354 = v341
		v355 = v342
		goto L86
	} else {
		goto L91
	}
L90:
	;
	v354 = v341
	v355 = v342
	goto L86
L91:
	;
	v345 = int32(1)
	if v342 == v341&int32(255) {
		v337 = v337 + v345
		v338 = v338 + v345
		goto L89
	} else {
		goto L92
	}
L92:
	;
	goto L90
L93:
	;
	v361 = m.G3
	v366 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v182))))
	v367 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v361)+uint32(_consts[1125]))))
	if v367 == int32(0) {
		v390 = v366
		v391 = v367
		goto L95
	} else {
		goto L96
	}
L94:
	;
	if v391-v390&int32(255) == int32(0) {
		goto L47
	} else {
		goto L102
	}
L95:
	;
	goto L94
L96:
	;
	if v367 != v366&int32(255) {
		v390 = v366
		v391 = v367
		goto L95
	} else {
		goto L97
	}
L97:
	;
	v373 = v361 + int32(_a1662)
	v374 = v182
	goto L98
L98:
	;
	v377 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v374)+1)))
	v378 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v373)+1)))
	if v378 == int32(0) {
		v390 = v377
		v391 = v378
		goto L95
	} else {
		goto L100
	}
L99:
	;
	v390 = v377
	v391 = v378
	goto L95
L100:
	;
	v381 = int32(1)
	if v378 == v377&int32(255) {
		v373 = v373 + v381
		v374 = v374 + v381
		goto L98
	} else {
		goto L101
	}
L101:
	;
	goto L99
L102:
	;
	v397 = m.G3
	v402 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v182))))
	v403 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v397)+uint32(_consts[1154]))))
	if v403 == int32(0) {
		v426 = v402
		v427 = v403
		goto L104
	} else {
		goto L105
	}
L103:
	;
	if v427-v426&int32(255) == int32(0) {
		goto L47
	} else {
		goto L111
	}
L104:
	;
	goto L103
L105:
	;
	if v403 != v402&int32(255) {
		v426 = v402
		v427 = v403
		goto L104
	} else {
		goto L106
	}
L106:
	;
	v409 = v397 + int32(_a2178)
	v410 = v182
	goto L107
L107:
	;
	v413 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v410)+1)))
	v414 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v409)+1)))
	if v414 == int32(0) {
		v426 = v413
		v427 = v414
		goto L104
	} else {
		goto L109
	}
L108:
	;
	v426 = v413
	v427 = v414
	goto L104
L109:
	;
	v417 = int32(1)
	if v414 == v413&int32(255) {
		v409 = v409 + v417
		v410 = v410 + v417
		goto L107
	} else {
		goto L110
	}
L110:
	;
	goto L108
L111:
	;
	v433 = m.G3
	v438 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v182))))
	v439 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v433)+uint32(_consts[1155]))))
	if v439 == int32(0) {
		v462 = v438
		v463 = v439
		goto L113
	} else {
		goto L114
	}
L112:
	;
	if v463-v462&int32(255) == int32(0) {
		goto L47
	} else {
		goto L120
	}
L113:
	;
	goto L112
L114:
	;
	if v439 != v438&int32(255) {
		v462 = v438
		v463 = v439
		goto L113
	} else {
		goto L115
	}
L115:
	;
	v445 = v433 + int32(_a2179)
	v446 = v182
	goto L116
L116:
	;
	v449 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v446)+1)))
	v450 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v445)+1)))
	if v450 == int32(0) {
		v462 = v449
		v463 = v450
		goto L113
	} else {
		goto L118
	}
L117:
	;
	v462 = v449
	v463 = v450
	goto L113
L118:
	;
	v453 = int32(1)
	if v450 == v449&int32(255) {
		v445 = v445 + v453
		v446 = v446 + v453
		goto L116
	} else {
		goto L119
	}
L119:
	;
	goto L117
L120:
	;
	v469 = m.G3
	v474 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v182))))
	v475 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v469)+uint32(_consts[1156]))))
	if v475 == int32(0) {
		v498 = v474
		v499 = v475
		goto L122
	} else {
		goto L123
	}
L121:
	;
	if v499-v498&int32(255) == int32(0) {
		goto L47
	} else {
		goto L129
	}
L122:
	;
	goto L121
L123:
	;
	if v475 != v474&int32(255) {
		v498 = v474
		v499 = v475
		goto L122
	} else {
		goto L124
	}
L124:
	;
	v481 = v469 + int32(_a1390)
	v482 = v182
	goto L125
L125:
	;
	v485 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v482)+1)))
	v486 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v481)+1)))
	if v486 == int32(0) {
		v498 = v485
		v499 = v486
		goto L122
	} else {
		goto L127
	}
L126:
	;
	v498 = v485
	v499 = v486
	goto L122
L127:
	;
	v489 = int32(1)
	if v486 == v485&int32(255) {
		v481 = v481 + v489
		v482 = v482 + v489
		goto L125
	} else {
		goto L128
	}
L128:
	;
	goto L126
L129:
	;
	v505 = m.G3
	v510 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v182))))
	v511 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v505)+uint32(_consts[1157]))))
	if v511 == int32(0) {
		v534 = v510
		v535 = v511
		goto L131
	} else {
		goto L132
	}
L130:
	;
	if v535-v534&int32(255) == int32(0) {
		goto L47
	} else {
		goto L138
	}
L131:
	;
	goto L130
L132:
	;
	if v511 != v510&int32(255) {
		v534 = v510
		v535 = v511
		goto L131
	} else {
		goto L133
	}
L133:
	;
	v517 = v505 + int32(_a1396)
	v518 = v182
	goto L134
L134:
	;
	v521 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v518)+1)))
	v522 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v517)+1)))
	if v522 == int32(0) {
		v534 = v521
		v535 = v522
		goto L131
	} else {
		goto L136
	}
L135:
	;
	v534 = v521
	v535 = v522
	goto L131
L136:
	;
	v525 = int32(1)
	if v522 == v521&int32(255) {
		v517 = v517 + v525
		v518 = v518 + v525
		goto L134
	} else {
		goto L137
	}
L137:
	;
	goto L135
L138:
	;
	v541 = m.G3
	v546 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v182))))
	v547 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v541)+uint32(_consts[1158]))))
	if v547 == int32(0) {
		v570 = v546
		v571 = v547
		goto L140
	} else {
		goto L141
	}
L139:
	;
	if v571-v570&int32(255) == int32(0) {
		goto L47
	} else {
		goto L147
	}
L140:
	;
	goto L139
L141:
	;
	if v547 != v546&int32(255) {
		v570 = v546
		v571 = v547
		goto L140
	} else {
		goto L142
	}
L142:
	;
	v553 = v541 + int32(_a2180)
	v554 = v182
	goto L143
L143:
	;
	v557 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v554)+1)))
	v558 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v553)+1)))
	if v558 == int32(0) {
		v570 = v557
		v571 = v558
		goto L140
	} else {
		goto L145
	}
L144:
	;
	v570 = v557
	v571 = v558
	goto L140
L145:
	;
	v561 = int32(1)
	if v558 == v557&int32(255) {
		v553 = v553 + v561
		v554 = v554 + v561
		goto L143
	} else {
		goto L146
	}
L146:
	;
	goto L144
L147:
	;
	v577 = m.G3
	v582 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v182))))
	v583 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v577)+uint32(_consts[1159]))))
	if v583 == int32(0) {
		v606 = v582
		v607 = v583
		goto L149
	} else {
		goto L150
	}
L148:
	;
	if v607-v606&int32(255) == int32(0) {
		goto L47
	} else {
		goto L156
	}
L149:
	;
	goto L148
L150:
	;
	if v583 != v582&int32(255) {
		v606 = v582
		v607 = v583
		goto L149
	} else {
		goto L151
	}
L151:
	;
	v589 = v577 + int32(_a2181)
	v590 = v182
	goto L152
L152:
	;
	v593 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v590)+1)))
	v594 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v589)+1)))
	if v594 == int32(0) {
		v606 = v593
		v607 = v594
		goto L149
	} else {
		goto L154
	}
L153:
	;
	v606 = v593
	v607 = v594
	goto L149
L154:
	;
	v597 = int32(1)
	if v594 == v593&int32(255) {
		v589 = v589 + v597
		v590 = v590 + v597
		goto L152
	} else {
		goto L155
	}
L155:
	;
	goto L153
L156:
	;
	v613 = m.G3
	v618 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v182))))
	v619 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v613)+uint32(_consts[1160]))))
	if v619 == int32(0) {
		v642 = v618
		v643 = v619
		goto L158
	} else {
		goto L159
	}
L157:
	;
	if v643-v642&int32(255) == int32(0) {
		goto L47
	} else {
		goto L165
	}
L158:
	;
	goto L157
L159:
	;
	if v619 != v618&int32(255) {
		v642 = v618
		v643 = v619
		goto L158
	} else {
		goto L160
	}
L160:
	;
	v625 = v613 + int32(_a2182)
	v626 = v182
	goto L161
L161:
	;
	v629 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v626)+1)))
	v630 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v625)+1)))
	if v630 == int32(0) {
		v642 = v629
		v643 = v630
		goto L158
	} else {
		goto L163
	}
L162:
	;
	v642 = v629
	v643 = v630
	goto L158
L163:
	;
	v633 = int32(1)
	if v630 == v629&int32(255) {
		v625 = v625 + v633
		v626 = v626 + v633
		goto L161
	} else {
		goto L164
	}
L164:
	;
	goto L162
L165:
	;
	v649 = m.G3
	v654 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v182))))
	v655 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v649)+uint32(_consts[1161]))))
	if v655 == int32(0) {
		v678 = v654
		v679 = v655
		goto L167
	} else {
		goto L168
	}
L166:
	;
	if v679-v678&int32(255) == int32(0) {
		goto L47
	} else {
		goto L174
	}
L167:
	;
	goto L166
L168:
	;
	if v655 != v654&int32(255) {
		v678 = v654
		v679 = v655
		goto L167
	} else {
		goto L169
	}
L169:
	;
	v661 = v649 + int32(_a2183)
	v662 = v182
	goto L170
L170:
	;
	v665 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v662)+1)))
	v666 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v661)+1)))
	if v666 == int32(0) {
		v678 = v665
		v679 = v666
		goto L167
	} else {
		goto L172
	}
L171:
	;
	v678 = v665
	v679 = v666
	goto L167
L172:
	;
	v669 = int32(1)
	if v666 == v665&int32(255) {
		v661 = v661 + v669
		v662 = v662 + v669
		goto L170
	} else {
		goto L173
	}
L173:
	;
	goto L171
L174:
	;
	v685 = m.G3
	v690 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v182))))
	v691 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v685)+uint32(_consts[1162]))))
	if v691 == int32(0) {
		v714 = v690
		v715 = v691
		goto L176
	} else {
		goto L177
	}
L175:
	;
	if v715-v714&int32(255) == int32(0) {
		goto L47
	} else {
		goto L183
	}
L176:
	;
	goto L175
L177:
	;
	if v691 != v690&int32(255) {
		v714 = v690
		v715 = v691
		goto L176
	} else {
		goto L178
	}
L178:
	;
	v697 = v685 + int32(_a2184)
	v698 = v182
	goto L179
L179:
	;
	v701 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v698)+1)))
	v702 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v697)+1)))
	if v702 == int32(0) {
		v714 = v701
		v715 = v702
		goto L176
	} else {
		goto L181
	}
L180:
	;
	v714 = v701
	v715 = v702
	goto L176
L181:
	;
	v705 = int32(1)
	if v702 == v701&int32(255) {
		v697 = v697 + v705
		v698 = v698 + v705
		goto L179
	} else {
		goto L182
	}
L182:
	;
	goto L180
L183:
	;
	v721 = m.G3
	v726 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v182))))
	v727 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v721)+uint32(_consts[1163]))))
	if v727 == int32(0) {
		v750 = v726
		v751 = v727
		goto L185
	} else {
		goto L186
	}
L184:
	;
	if v751-v750&int32(255) == int32(0) {
		goto L47
	} else {
		goto L192
	}
L185:
	;
	goto L184
L186:
	;
	if v727 != v726&int32(255) {
		v750 = v726
		v751 = v727
		goto L185
	} else {
		goto L187
	}
L187:
	;
	v733 = v721 + int32(_a2185)
	v734 = v182
	goto L188
L188:
	;
	v737 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v734)+1)))
	v738 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v733)+1)))
	if v738 == int32(0) {
		v750 = v737
		v751 = v738
		goto L185
	} else {
		goto L190
	}
L189:
	;
	v750 = v737
	v751 = v738
	goto L185
L190:
	;
	v741 = int32(1)
	if v738 == v737&int32(255) {
		v733 = v733 + v741
		v734 = v734 + v741
		goto L188
	} else {
		goto L191
	}
L191:
	;
	goto L189
L192:
	;
	v757 = m.G3
	v762 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v182))))
	v763 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v757)+uint32(_consts[1164]))))
	if v763 == int32(0) {
		v786 = v762
		v787 = v763
		goto L194
	} else {
		goto L195
	}
L193:
	;
	if v787-v786&int32(255) == int32(0) {
		goto L47
	} else {
		goto L201
	}
L194:
	;
	goto L193
L195:
	;
	if v763 != v762&int32(255) {
		v786 = v762
		v787 = v763
		goto L194
	} else {
		goto L196
	}
L196:
	;
	v769 = v757 + int32(_a2186)
	v770 = v182
	goto L197
L197:
	;
	v773 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v770)+1)))
	v774 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v769)+1)))
	if v774 == int32(0) {
		v786 = v773
		v787 = v774
		goto L194
	} else {
		goto L199
	}
L198:
	;
	v786 = v773
	v787 = v774
	goto L194
L199:
	;
	v777 = int32(1)
	if v774 == v773&int32(255) {
		v769 = v769 + v777
		v770 = v770 + v777
		goto L197
	} else {
		goto L200
	}
L200:
	;
	goto L198
L201:
	;
	v793 = m.G3
	v798 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v182))))
	v799 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v793)+uint32(_consts[1165]))))
	if v799 == int32(0) {
		v822 = v798
		v823 = v799
		goto L203
	} else {
		goto L204
	}
L202:
	;
	if v823-v822&int32(255) == int32(0) {
		goto L47
	} else {
		goto L210
	}
L203:
	;
	goto L202
L204:
	;
	if v799 != v798&int32(255) {
		v822 = v798
		v823 = v799
		goto L203
	} else {
		goto L205
	}
L205:
	;
	v805 = v793 + int32(_a2187)
	v806 = v182
	goto L206
L206:
	;
	v809 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v806)+1)))
	v810 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v805)+1)))
	if v810 == int32(0) {
		v822 = v809
		v823 = v810
		goto L203
	} else {
		goto L208
	}
L207:
	;
	v822 = v809
	v823 = v810
	goto L203
L208:
	;
	v813 = int32(1)
	if v810 == v809&int32(255) {
		v805 = v805 + v813
		v806 = v806 + v813
		goto L206
	} else {
		goto L209
	}
L209:
	;
	goto L207
L210:
	;
	v829 = m.G3
	v834 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v182))))
	v835 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v829)+uint32(_consts[1166]))))
	if v835 == int32(0) {
		v858 = v834
		v859 = v835
		goto L212
	} else {
		goto L213
	}
L211:
	;
	if v859-v858&int32(255) == int32(0) {
		goto L47
	} else {
		goto L219
	}
L212:
	;
	goto L211
L213:
	;
	if v835 != v834&int32(255) {
		v858 = v834
		v859 = v835
		goto L212
	} else {
		goto L214
	}
L214:
	;
	v841 = v829 + int32(_a2188)
	v842 = v182
	goto L215
L215:
	;
	v845 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v842)+1)))
	v846 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v841)+1)))
	if v846 == int32(0) {
		v858 = v845
		v859 = v846
		goto L212
	} else {
		goto L217
	}
L216:
	;
	v858 = v845
	v859 = v846
	goto L212
L217:
	;
	v849 = int32(1)
	if v846 == v845&int32(255) {
		v841 = v841 + v849
		v842 = v842 + v849
		goto L215
	} else {
		goto L218
	}
L218:
	;
	goto L216
L219:
	;
	v865 = m.G3
	v870 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v182))))
	v871 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v865)+uint32(_consts[1167]))))
	if v871 == int32(0) {
		v894 = v870
		v895 = v871
		goto L221
	} else {
		goto L222
	}
L220:
	;
	if v895-v894&int32(255) == int32(0) {
		goto L47
	} else {
		goto L228
	}
L221:
	;
	goto L220
L222:
	;
	if v871 != v870&int32(255) {
		v894 = v870
		v895 = v871
		goto L221
	} else {
		goto L223
	}
L223:
	;
	v877 = v865 + int32(_a2189)
	v878 = v182
	goto L224
L224:
	;
	v881 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v878)+1)))
	v882 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v877)+1)))
	if v882 == int32(0) {
		v894 = v881
		v895 = v882
		goto L221
	} else {
		goto L226
	}
L225:
	;
	v894 = v881
	v895 = v882
	goto L221
L226:
	;
	v885 = int32(1)
	if v882 == v881&int32(255) {
		v877 = v877 + v885
		v878 = v878 + v885
		goto L224
	} else {
		goto L227
	}
L227:
	;
	goto L225
L228:
	;
	v901 = m.G3
	v906 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v182))))
	v907 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v901)+uint32(_consts[1168]))))
	if v907 == int32(0) {
		v930 = v906
		v931 = v907
		goto L230
	} else {
		goto L231
	}
L229:
	;
	if v931-v930&int32(255) == int32(0) {
		goto L47
	} else {
		goto L237
	}
L230:
	;
	goto L229
L231:
	;
	if v907 != v906&int32(255) {
		v930 = v906
		v931 = v907
		goto L230
	} else {
		goto L232
	}
L232:
	;
	v913 = v901 + int32(_a2190)
	v914 = v182
	goto L233
L233:
	;
	v917 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v914)+1)))
	v918 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v913)+1)))
	if v918 == int32(0) {
		v930 = v917
		v931 = v918
		goto L230
	} else {
		goto L235
	}
L234:
	;
	v930 = v917
	v931 = v918
	goto L230
L235:
	;
	v921 = int32(1)
	if v918 == v917&int32(255) {
		v913 = v913 + v921
		v914 = v914 + v921
		goto L233
	} else {
		goto L236
	}
L236:
	;
	goto L234
L237:
	;
	v937 = m.G3
	v942 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v182))))
	v943 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v937)+uint32(_consts[1169]))))
	if v943 == int32(0) {
		v966 = v942
		v967 = v943
		goto L239
	} else {
		goto L240
	}
L238:
	;
	if v967-v966&int32(255) == int32(0) {
		goto L47
	} else {
		goto L246
	}
L239:
	;
	goto L238
L240:
	;
	if v943 != v942&int32(255) {
		v966 = v942
		v967 = v943
		goto L239
	} else {
		goto L241
	}
L241:
	;
	v949 = v937 + int32(_a2191)
	v950 = v182
	goto L242
L242:
	;
	v953 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v950)+1)))
	v954 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v949)+1)))
	if v954 == int32(0) {
		v966 = v953
		v967 = v954
		goto L239
	} else {
		goto L244
	}
L243:
	;
	v966 = v953
	v967 = v954
	goto L239
L244:
	;
	v957 = int32(1)
	if v954 == v953&int32(255) {
		v949 = v949 + v957
		v950 = v950 + v957
		goto L242
	} else {
		goto L245
	}
L245:
	;
	goto L243
L246:
	;
	v973 = m.G3
	v978 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v182))))
	v979 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v973)+uint32(_consts[1170]))))
	if v979 == int32(0) {
		v1002 = v978
		v1003 = v979
		goto L248
	} else {
		goto L249
	}
L247:
	;
	if v1003-v1002&int32(255) == int32(0) {
		goto L47
	} else {
		goto L255
	}
L248:
	;
	goto L247
L249:
	;
	if v979 != v978&int32(255) {
		v1002 = v978
		v1003 = v979
		goto L248
	} else {
		goto L250
	}
L250:
	;
	v985 = v973 + int32(_a2192)
	v986 = v182
	goto L251
L251:
	;
	v989 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v986)+1)))
	v990 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v985)+1)))
	if v990 == int32(0) {
		v1002 = v989
		v1003 = v990
		goto L248
	} else {
		goto L253
	}
L252:
	;
	v1002 = v989
	v1003 = v990
	goto L248
L253:
	;
	v993 = int32(1)
	if v990 == v989&int32(255) {
		v985 = v985 + v993
		v986 = v986 + v993
		goto L251
	} else {
		goto L254
	}
L254:
	;
	goto L252
L255:
	;
	v1009 = m.G3
	v1014 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v182))))
	v1015 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1009)+uint32(_consts[1171]))))
	if v1015 == int32(0) {
		v1038 = v1014
		v1039 = v1015
		goto L257
	} else {
		goto L258
	}
L256:
	;
	if v1039-v1038&int32(255) == int32(0) {
		goto L47
	} else {
		goto L264
	}
L257:
	;
	goto L256
L258:
	;
	if v1015 != v1014&int32(255) {
		v1038 = v1014
		v1039 = v1015
		goto L257
	} else {
		goto L259
	}
L259:
	;
	v1021 = v1009 + int32(_a2193)
	v1022 = v182
	goto L260
L260:
	;
	v1025 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1022)+1)))
	v1026 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1021)+1)))
	if v1026 == int32(0) {
		v1038 = v1025
		v1039 = v1026
		goto L257
	} else {
		goto L262
	}
L261:
	;
	v1038 = v1025
	v1039 = v1026
	goto L257
L262:
	;
	v1029 = int32(1)
	if v1026 == v1025&int32(255) {
		v1021 = v1021 + v1029
		v1022 = v1022 + v1029
		goto L260
	} else {
		goto L263
	}
L263:
	;
	goto L261
L264:
	;
	v1045 = m.G3
	v1050 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v182))))
	v1051 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1045)+uint32(_consts[1172]))))
	if v1051 == int32(0) {
		v1074 = v1050
		v1075 = v1051
		goto L266
	} else {
		goto L267
	}
L265:
	;
	if v1075-v1074&int32(255) == int32(0) {
		goto L47
	} else {
		goto L273
	}
L266:
	;
	goto L265
L267:
	;
	if v1051 != v1050&int32(255) {
		v1074 = v1050
		v1075 = v1051
		goto L266
	} else {
		goto L268
	}
L268:
	;
	v1057 = v1045 + int32(_a63)
	v1058 = v182
	goto L269
L269:
	;
	v1061 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1058)+1)))
	v1062 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1057)+1)))
	if v1062 == int32(0) {
		v1074 = v1061
		v1075 = v1062
		goto L266
	} else {
		goto L271
	}
L270:
	;
	v1074 = v1061
	v1075 = v1062
	goto L266
L271:
	;
	v1065 = int32(1)
	if v1062 == v1061&int32(255) {
		v1057 = v1057 + v1065
		v1058 = v1058 + v1065
		goto L269
	} else {
		goto L272
	}
L272:
	;
	goto L270
L273:
	;
	v1081 = m.G3
	v1086 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v182))))
	v1087 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1081)+uint32(_consts[1173]))))
	if v1087 == int32(0) {
		v1110 = v1086
		v1111 = v1087
		goto L275
	} else {
		goto L276
	}
L274:
	;
	if v1111-v1110&int32(255) == int32(0) {
		goto L47
	} else {
		goto L282
	}
L275:
	;
	goto L274
L276:
	;
	if v1087 != v1086&int32(255) {
		v1110 = v1086
		v1111 = v1087
		goto L275
	} else {
		goto L277
	}
L277:
	;
	v1093 = v1081 + int32(_a2194)
	v1094 = v182
	goto L278
L278:
	;
	v1097 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1094)+1)))
	v1098 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1093)+1)))
	if v1098 == int32(0) {
		v1110 = v1097
		v1111 = v1098
		goto L275
	} else {
		goto L280
	}
L279:
	;
	v1110 = v1097
	v1111 = v1098
	goto L275
L280:
	;
	v1101 = int32(1)
	if v1098 == v1097&int32(255) {
		v1093 = v1093 + v1101
		v1094 = v1094 + v1101
		goto L278
	} else {
		goto L281
	}
L281:
	;
	goto L279
L282:
	;
	v1117 = m.G3
	v1122 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v182))))
	v1123 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1117)+uint32(_consts[1174]))))
	if v1123 == int32(0) {
		v1146 = v1122
		v1147 = v1123
		goto L284
	} else {
		goto L285
	}
L283:
	;
	if v1147-v1146&int32(255) == int32(0) {
		goto L47
	} else {
		goto L291
	}
L284:
	;
	goto L283
L285:
	;
	if v1123 != v1122&int32(255) {
		v1146 = v1122
		v1147 = v1123
		goto L284
	} else {
		goto L286
	}
L286:
	;
	v1129 = v1117 + int32(_a1400)
	v1130 = v182
	goto L287
L287:
	;
	v1133 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1130)+1)))
	v1134 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1129)+1)))
	if v1134 == int32(0) {
		v1146 = v1133
		v1147 = v1134
		goto L284
	} else {
		goto L289
	}
L288:
	;
	v1146 = v1133
	v1147 = v1134
	goto L284
L289:
	;
	v1137 = int32(1)
	if v1134 == v1133&int32(255) {
		v1129 = v1129 + v1137
		v1130 = v1130 + v1137
		goto L287
	} else {
		goto L290
	}
L290:
	;
	goto L288
L291:
	;
	v1153 = m.G3
	v1158 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v182))))
	v1159 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1153)+uint32(_consts[1175]))))
	if v1159 == int32(0) {
		v1182 = v1158
		v1183 = v1159
		goto L293
	} else {
		goto L294
	}
L292:
	;
	if v1183-v1182&int32(255) == int32(0) {
		goto L47
	} else {
		goto L300
	}
L293:
	;
	goto L292
L294:
	;
	if v1159 != v1158&int32(255) {
		v1182 = v1158
		v1183 = v1159
		goto L293
	} else {
		goto L295
	}
L295:
	;
	v1165 = v1153 + int32(_a2195)
	v1166 = v182
	goto L296
L296:
	;
	v1169 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1166)+1)))
	v1170 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1165)+1)))
	if v1170 == int32(0) {
		v1182 = v1169
		v1183 = v1170
		goto L293
	} else {
		goto L298
	}
L297:
	;
	v1182 = v1169
	v1183 = v1170
	goto L293
L298:
	;
	v1173 = int32(1)
	if v1170 == v1169&int32(255) {
		v1165 = v1165 + v1173
		v1166 = v1166 + v1173
		goto L296
	} else {
		goto L299
	}
L299:
	;
	goto L297
L300:
	;
	v1189 = m.G3
	v1194 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v182))))
	v1195 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1189)+uint32(_consts[1176]))))
	if v1195 == int32(0) {
		v1218 = v1194
		v1219 = v1195
		goto L302
	} else {
		goto L303
	}
L301:
	;
	if v1219-v1218&int32(255) == int32(0) {
		goto L47
	} else {
		goto L309
	}
L302:
	;
	goto L301
L303:
	;
	if v1195 != v1194&int32(255) {
		v1218 = v1194
		v1219 = v1195
		goto L302
	} else {
		goto L304
	}
L304:
	;
	v1201 = v1189 + int32(_a2196)
	v1202 = v182
	goto L305
L305:
	;
	v1205 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1202)+1)))
	v1206 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1201)+1)))
	if v1206 == int32(0) {
		v1218 = v1205
		v1219 = v1206
		goto L302
	} else {
		goto L307
	}
L306:
	;
	v1218 = v1205
	v1219 = v1206
	goto L302
L307:
	;
	v1209 = int32(1)
	if v1206 == v1205&int32(255) {
		v1201 = v1201 + v1209
		v1202 = v1202 + v1209
		goto L305
	} else {
		goto L308
	}
L308:
	;
	goto L306
L309:
	;
	v1225 = m.G3
	v1230 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v182))))
	v1231 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1225)+uint32(_consts[1177]))))
	if v1231 == int32(0) {
		v1254 = v1230
		v1255 = v1231
		goto L311
	} else {
		goto L312
	}
L310:
	;
	if v1255-v1254&int32(255) == int32(0) {
		goto L47
	} else {
		goto L318
	}
L311:
	;
	goto L310
L312:
	;
	if v1231 != v1230&int32(255) {
		v1254 = v1230
		v1255 = v1231
		goto L311
	} else {
		goto L313
	}
L313:
	;
	v1237 = v1225 + int32(_a2197)
	v1238 = v182
	goto L314
L314:
	;
	v1241 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1238)+1)))
	v1242 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1237)+1)))
	if v1242 == int32(0) {
		v1254 = v1241
		v1255 = v1242
		goto L311
	} else {
		goto L316
	}
L315:
	;
	v1254 = v1241
	v1255 = v1242
	goto L311
L316:
	;
	v1245 = int32(1)
	if v1242 == v1241&int32(255) {
		v1237 = v1237 + v1245
		v1238 = v1238 + v1245
		goto L314
	} else {
		goto L317
	}
L317:
	;
	goto L315
L318:
	;
	v1261 = m.G3
	v1266 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v182))))
	v1267 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1261)+uint32(_consts[1178]))))
	if v1267 == int32(0) {
		v1290 = v1266
		v1291 = v1267
		goto L320
	} else {
		goto L321
	}
L319:
	;
	if v1291-v1290&int32(255) == int32(0) {
		goto L47
	} else {
		goto L327
	}
L320:
	;
	goto L319
L321:
	;
	if v1267 != v1266&int32(255) {
		v1290 = v1266
		v1291 = v1267
		goto L320
	} else {
		goto L322
	}
L322:
	;
	v1273 = v1261 + int32(_a2198)
	v1274 = v182
	goto L323
L323:
	;
	v1277 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1274)+1)))
	v1278 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1273)+1)))
	if v1278 == int32(0) {
		v1290 = v1277
		v1291 = v1278
		goto L320
	} else {
		goto L325
	}
L324:
	;
	v1290 = v1277
	v1291 = v1278
	goto L320
L325:
	;
	v1281 = int32(1)
	if v1278 == v1277&int32(255) {
		v1273 = v1273 + v1281
		v1274 = v1274 + v1281
		goto L323
	} else {
		goto L326
	}
L326:
	;
	goto L324
L327:
	;
	v1297 = m.G3
	v1302 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v182))))
	v1303 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1297)+uint32(_consts[1179]))))
	if v1303 == int32(0) {
		v1326 = v1302
		v1327 = v1303
		goto L329
	} else {
		goto L330
	}
L328:
	;
	if v1327-v1326&int32(255) == int32(0) {
		goto L47
	} else {
		goto L336
	}
L329:
	;
	goto L328
L330:
	;
	if v1303 != v1302&int32(255) {
		v1326 = v1302
		v1327 = v1303
		goto L329
	} else {
		goto L331
	}
L331:
	;
	v1309 = v1297 + int32(_a2199)
	v1310 = v182
	goto L332
L332:
	;
	v1313 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1310)+1)))
	v1314 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1309)+1)))
	if v1314 == int32(0) {
		v1326 = v1313
		v1327 = v1314
		goto L329
	} else {
		goto L334
	}
L333:
	;
	v1326 = v1313
	v1327 = v1314
	goto L329
L334:
	;
	v1317 = int32(1)
	if v1314 == v1313&int32(255) {
		v1309 = v1309 + v1317
		v1310 = v1310 + v1317
		goto L332
	} else {
		goto L335
	}
L335:
	;
	goto L333
L336:
	;
	v1333 = m.G3
	v1338 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v182))))
	v1339 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1333)+uint32(_consts[1180]))))
	if v1339 == int32(0) {
		v1362 = v1338
		v1363 = v1339
		goto L338
	} else {
		goto L339
	}
L337:
	;
	if v1363-v1362&int32(255) == int32(0) {
		goto L47
	} else {
		goto L345
	}
L338:
	;
	goto L337
L339:
	;
	if v1339 != v1338&int32(255) {
		v1362 = v1338
		v1363 = v1339
		goto L338
	} else {
		goto L340
	}
L340:
	;
	v1345 = v1333 + int32(_a2200)
	v1346 = v182
	goto L341
L341:
	;
	v1349 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1346)+1)))
	v1350 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1345)+1)))
	if v1350 == int32(0) {
		v1362 = v1349
		v1363 = v1350
		goto L338
	} else {
		goto L343
	}
L342:
	;
	v1362 = v1349
	v1363 = v1350
	goto L338
L343:
	;
	v1353 = int32(1)
	if v1350 == v1349&int32(255) {
		v1345 = v1345 + v1353
		v1346 = v1346 + v1353
		goto L341
	} else {
		goto L344
	}
L344:
	;
	goto L342
L345:
	;
	v1369 = m.G3
	v1374 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v182))))
	v1375 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1369)+uint32(_consts[1181]))))
	if v1375 == int32(0) {
		v1398 = v1374
		v1399 = v1375
		goto L347
	} else {
		goto L348
	}
L346:
	;
	if v1399-v1398&int32(255) == int32(0) {
		goto L47
	} else {
		goto L354
	}
L347:
	;
	goto L346
L348:
	;
	if v1375 != v1374&int32(255) {
		v1398 = v1374
		v1399 = v1375
		goto L347
	} else {
		goto L349
	}
L349:
	;
	v1381 = v1369 + int32(_a2201)
	v1382 = v182
	goto L350
L350:
	;
	v1385 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1382)+1)))
	v1386 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1381)+1)))
	if v1386 == int32(0) {
		v1398 = v1385
		v1399 = v1386
		goto L347
	} else {
		goto L352
	}
L351:
	;
	v1398 = v1385
	v1399 = v1386
	goto L347
L352:
	;
	v1389 = int32(1)
	if v1386 == v1385&int32(255) {
		v1381 = v1381 + v1389
		v1382 = v1382 + v1389
		goto L350
	} else {
		goto L353
	}
L353:
	;
	goto L351
L354:
	;
	v1405 = m.G3
	v1410 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v182))))
	v1411 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1405)+uint32(_consts[1182]))))
	if v1411 == int32(0) {
		v1434 = v1410
		v1435 = v1411
		goto L356
	} else {
		goto L357
	}
L355:
	;
	if v1435-v1434&int32(255) == int32(0) {
		goto L47
	} else {
		goto L363
	}
L356:
	;
	goto L355
L357:
	;
	if v1411 != v1410&int32(255) {
		v1434 = v1410
		v1435 = v1411
		goto L356
	} else {
		goto L358
	}
L358:
	;
	v1417 = v1405 + int32(_a756)
	v1418 = v182
	goto L359
L359:
	;
	v1421 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1418)+1)))
	v1422 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1417)+1)))
	if v1422 == int32(0) {
		v1434 = v1421
		v1435 = v1422
		goto L356
	} else {
		goto L361
	}
L360:
	;
	v1434 = v1421
	v1435 = v1422
	goto L356
L361:
	;
	v1425 = int32(1)
	if v1422 == v1421&int32(255) {
		v1417 = v1417 + v1425
		v1418 = v1418 + v1425
		goto L359
	} else {
		goto L362
	}
L362:
	;
	goto L360
L363:
	;
	v1441 = m.G3
	v1446 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v182))))
	v1447 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1441)+uint32(_consts[1183]))))
	if v1447 == int32(0) {
		v1470 = v1446
		v1471 = v1447
		goto L365
	} else {
		goto L366
	}
L364:
	;
	if v1471-v1470&int32(255) == int32(0) {
		goto L47
	} else {
		goto L372
	}
L365:
	;
	goto L364
L366:
	;
	if v1447 != v1446&int32(255) {
		v1470 = v1446
		v1471 = v1447
		goto L365
	} else {
		goto L367
	}
L367:
	;
	v1453 = v1441 + int32(_a2202)
	v1454 = v182
	goto L368
L368:
	;
	v1457 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1454)+1)))
	v1458 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1453)+1)))
	if v1458 == int32(0) {
		v1470 = v1457
		v1471 = v1458
		goto L365
	} else {
		goto L370
	}
L369:
	;
	v1470 = v1457
	v1471 = v1458
	goto L365
L370:
	;
	v1461 = int32(1)
	if v1458 == v1457&int32(255) {
		v1453 = v1453 + v1461
		v1454 = v1454 + v1461
		goto L368
	} else {
		goto L371
	}
L371:
	;
	goto L369
L372:
	;
	v1477 = m.G3
	v1482 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v182))))
	v1483 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1477)+uint32(_consts[1184]))))
	if v1483 == int32(0) {
		v1506 = v1482
		v1507 = v1483
		goto L374
	} else {
		goto L375
	}
L373:
	;
	if v1507-v1506&int32(255) == int32(0) {
		goto L47
	} else {
		goto L381
	}
L374:
	;
	goto L373
L375:
	;
	if v1483 != v1482&int32(255) {
		v1506 = v1482
		v1507 = v1483
		goto L374
	} else {
		goto L376
	}
L376:
	;
	v1489 = v1477 + int32(_a713)
	v1490 = v182
	goto L377
L377:
	;
	v1493 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1490)+1)))
	v1494 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1489)+1)))
	if v1494 == int32(0) {
		v1506 = v1493
		v1507 = v1494
		goto L374
	} else {
		goto L379
	}
L378:
	;
	v1506 = v1493
	v1507 = v1494
	goto L374
L379:
	;
	v1497 = int32(1)
	if v1494 == v1493&int32(255) {
		v1489 = v1489 + v1497
		v1490 = v1490 + v1497
		goto L377
	} else {
		goto L380
	}
L380:
	;
	goto L378
L381:
	;
	v1513 = m.G3
	v1518 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v182))))
	v1519 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1513)+uint32(_consts[1185]))))
	if v1519 == int32(0) {
		v1542 = v1518
		v1543 = v1519
		goto L385
	} else {
		goto L386
	}
L382:
	;
	v1625 = m.G3
	v1630 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v182))))
	v1631 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1625)+uint32(_consts[1186]))))
	if v1631 == int32(0) {
		v1654 = v1630
		v1655 = v1631
		goto L413
	} else {
		goto L414
	}
L383:
	;
	v1622 = *(*int32)(unsafe.Add(mBase, uint32(v12)+20))
	if v1622 != 0 {
		goto L47
	} else {
		goto L411
	}
L384:
	;
	if v1543-v1542&int32(255) == int32(0) {
		goto L383
	} else {
		goto L392
	}
L385:
	;
	goto L384
L386:
	;
	if v1519 != v1518&int32(255) {
		v1542 = v1518
		v1543 = v1519
		goto L385
	} else {
		goto L387
	}
L387:
	;
	v1525 = v1513 + int32(_a2203)
	v1526 = v182
	goto L388
L388:
	;
	v1529 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1526)+1)))
	v1530 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1525)+1)))
	if v1530 == int32(0) {
		v1542 = v1529
		v1543 = v1530
		goto L385
	} else {
		goto L390
	}
L389:
	;
	v1542 = v1529
	v1543 = v1530
	goto L385
L390:
	;
	v1533 = int32(1)
	if v1530 == v1529&int32(255) {
		v1525 = v1525 + v1533
		v1526 = v1526 + v1533
		goto L388
	} else {
		goto L391
	}
L391:
	;
	goto L389
L392:
	;
	v1549 = m.G3
	v1554 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v182))))
	v1555 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1549)+uint32(_consts[1187]))))
	if v1555 == int32(0) {
		v1578 = v1554
		v1579 = v1555
		goto L394
	} else {
		goto L395
	}
L393:
	;
	if v1579-v1578&int32(255) == int32(0) {
		goto L383
	} else {
		goto L401
	}
L394:
	;
	goto L393
L395:
	;
	if v1555 != v1554&int32(255) {
		v1578 = v1554
		v1579 = v1555
		goto L394
	} else {
		goto L396
	}
L396:
	;
	v1561 = v1549 + int32(_a2204)
	v1562 = v182
	goto L397
L397:
	;
	v1565 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1562)+1)))
	v1566 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1561)+1)))
	if v1566 == int32(0) {
		v1578 = v1565
		v1579 = v1566
		goto L394
	} else {
		goto L399
	}
L398:
	;
	v1578 = v1565
	v1579 = v1566
	goto L394
L399:
	;
	v1569 = int32(1)
	if v1566 == v1565&int32(255) {
		v1561 = v1561 + v1569
		v1562 = v1562 + v1569
		goto L397
	} else {
		goto L400
	}
L400:
	;
	goto L398
L401:
	;
	v1585 = m.G3
	v1590 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v182))))
	v1591 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1585)+uint32(_consts[1188]))))
	if v1591 == int32(0) {
		v1614 = v1590
		v1615 = v1591
		goto L403
	} else {
		goto L404
	}
L402:
	;
	if v1615-v1614&int32(255) == int32(0) {
		goto L383
	} else {
		goto L410
	}
L403:
	;
	goto L402
L404:
	;
	if v1591 != v1590&int32(255) {
		v1614 = v1590
		v1615 = v1591
		goto L403
	} else {
		goto L405
	}
L405:
	;
	v1597 = v1585 + int32(_a2205)
	v1598 = v182
	goto L406
L406:
	;
	v1601 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1598)+1)))
	v1602 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1597)+1)))
	if v1602 == int32(0) {
		v1614 = v1601
		v1615 = v1602
		goto L403
	} else {
		goto L408
	}
L407:
	;
	v1614 = v1601
	v1615 = v1602
	goto L403
L408:
	;
	v1605 = int32(1)
	if v1602 == v1601&int32(255) {
		v1597 = v1597 + v1605
		v1598 = v1598 + v1605
		goto L406
	} else {
		goto L409
	}
L409:
	;
	goto L407
L410:
	;
	v1624 = int32(0)
	goto L382
L411:
	;
	v1624 = int32(1)
	goto L382
L412:
	;
	if v1655-v1654&int32(255) == int32(0) {
		goto L46
	} else {
		goto L420
	}
L413:
	;
	goto L412
L414:
	;
	if v1631 != v1630&int32(255) {
		v1654 = v1630
		v1655 = v1631
		goto L413
	} else {
		goto L415
	}
L415:
	;
	v1637 = v1625 + int32(_a2206)
	v1638 = v182
	goto L416
L416:
	;
	v1641 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1638)+1)))
	v1642 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1637)+1)))
	if v1642 == int32(0) {
		v1654 = v1641
		v1655 = v1642
		goto L413
	} else {
		goto L418
	}
L417:
	;
	v1654 = v1641
	v1655 = v1642
	goto L413
L418:
	;
	v1645 = int32(1)
	if v1642 == v1641&int32(255) {
		v1637 = v1637 + v1645
		v1638 = v1638 + v1645
		goto L416
	} else {
		goto L419
	}
L419:
	;
	goto L417
L420:
	;
	v1661 = m.G3
	v1666 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v182))))
	v1667 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1661)+uint32(_consts[1189]))))
	if v1667 == int32(0) {
		v1690 = v1666
		v1691 = v1667
		goto L422
	} else {
		goto L423
	}
L421:
	;
	if v1691-v1690&int32(255) == int32(0) {
		goto L46
	} else {
		goto L429
	}
L422:
	;
	goto L421
L423:
	;
	if v1667 != v1666&int32(255) {
		v1690 = v1666
		v1691 = v1667
		goto L422
	} else {
		goto L424
	}
L424:
	;
	v1673 = v1661 + int32(_a2207)
	v1674 = v182
	goto L425
L425:
	;
	v1677 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1674)+1)))
	v1678 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1673)+1)))
	if v1678 == int32(0) {
		v1690 = v1677
		v1691 = v1678
		goto L422
	} else {
		goto L427
	}
L426:
	;
	v1690 = v1677
	v1691 = v1678
	goto L422
L427:
	;
	v1681 = int32(1)
	if v1678 == v1677&int32(255) {
		v1673 = v1673 + v1681
		v1674 = v1674 + v1681
		goto L425
	} else {
		goto L428
	}
L428:
	;
	goto L426
L429:
	;
	v1697 = m.G3
	v1702 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v182))))
	v1703 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1697)+uint32(_consts[1190]))))
	if v1703 == int32(0) {
		v1726 = v1702
		v1727 = v1703
		goto L431
	} else {
		goto L432
	}
L430:
	;
	if base.B2i32(v1727-v1726&int32(255) == int32(0))|v1624 != 0 {
		goto L46
	} else {
		goto L438
	}
L431:
	;
	goto L430
L432:
	;
	if v1703 != v1702&int32(255) {
		v1726 = v1702
		v1727 = v1703
		goto L431
	} else {
		goto L433
	}
L433:
	;
	v1709 = v1697 + int32(_a2208)
	v1710 = v182
	goto L434
L434:
	;
	v1713 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1710)+1)))
	v1714 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1709)+1)))
	if v1714 == int32(0) {
		v1726 = v1713
		v1727 = v1714
		goto L431
	} else {
		goto L436
	}
L435:
	;
	v1726 = v1713
	v1727 = v1714
	goto L431
L436:
	;
	v1717 = int32(1)
	if v1714 == v1713&int32(255) {
		v1709 = v1709 + v1717
		v1710 = v1710 + v1717
		goto L434
	} else {
		goto L437
	}
L437:
	;
	goto L435
L438:
	;
	v1734 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
	*(*int32)(unsafe.Add(mBase, uint32(v7))) = v182
	v1736 = m.G3
	v1741 = m.G10
	v1742 = *(*int32)(unsafe.Add(mBase, uint32(v1741)))
	m.T0[v1742].(func(*base.Module, int32, int32, int32, int32))(m, v1734, v1736+int32(_a716), v1736+int32(_a2209), v7)
	mBase = m.M
	v1744 = m.ExcPending
	if v1744 != 0 {
		goto L1
	} else {
		goto L439
	}
L439:
	;
	goto L46
L440:
	;
	goto L46
}
func F_luaRedisAclCheckCmdPermissionsCommand(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
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
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
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
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v64 int32
	_ = v64
	var v69 int32
	_ = v69
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v89 int32
	_ = v89
	var v93 int32
	_ = v93
	var v99 int32
	_ = v99
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v109 int32
	_ = v109
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v115 int32
	_ = v115
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v121 int32
	_ = v121
	var v124 int32
	_ = v124
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
	var v135 int32
	_ = v135
	var v148 int32
	_ = v148
	var v153 int32
	_ = v153
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v168 int32
	_ = v168
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v196 int32
	_ = v196
	var v199 int32
	_ = v199
	var v202 int32
	_ = v202
	var v208 int32
	_ = v208
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v215 int32
	_ = v215
	var v218 int32
	_ = v218
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v227 int32
	_ = v227
	var v229 int32
	_ = v229
	var v231 int32
	_ = v231
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v247 int32
	_ = v247
	var v250 int32
	_ = v250
	var v251 int32
	_ = v251
	var v253 int32
	_ = v253
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v259 int32
	_ = v259
	var v262 int32
	_ = v262
	var v263 int32
	_ = v263
	var v265 int32
	_ = v265
	var v266 int32
	_ = v266
	var v267 int32
	_ = v267
	var v269 int32
	_ = v269
	var v273 int32
	_ = v273
	var v286 int32
	_ = v286
	var v291 int32
	_ = v291
	var v298 int32
	_ = v298
	var v299 int32
	_ = v299
	var v300 int32
	_ = v300
	var v302 int32
	_ = v302
	var v303 int32
	_ = v303
	var v306 int32
	_ = v306
	var v319 int32
	_ = v319
	var v320 int32
	_ = v320
	var v322 int32
	_ = v322
	var v326 int32
	_ = v326
	var v339 int32
	_ = v339
	var v345 int32
	_ = v345
	var v346 int32
	_ = v346
	var v348 int32
	_ = v348
	v12 = m.G0
	v14 = v12 - int32(16)
	m.G0 = v14
	v16 = m.G3
	v19 = F_luaGetFromRegistry(m, l0, v16+int32(_a2168))
	mBase = m.M
	v22 = m.ExcPending
	if v22 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	v339 = m.G3
	v345 = m.G8
	v346 = *(*int32)(unsafe.Add(mBase, uint32(v345)))
	m.T0[v346].(func(*base.Module, int32, int32, int32))(m, v339+int32(_a2210), v339+int32(_a2211), int32(1386))
	mBase = m.M
	v348 = m.ExcPending
	if v348 != 0 {
		goto L2
	} else {
		goto L61
	}
L2:
	;
	return int32(0)
L3:
	;
	if v19 == int32(0) {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+12)) = int32(0)
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v19)))
	v30 = F_luaArgsToServerArgv(m, v27, l0, v14+int32(12))
	mBase = m.M
	v31 = m.ExcPending
	if v31 != 0 {
		goto L2
	} else {
		goto L7
	}
L5:
	;
	m.G0 = v14 + int32(16)
	return v326
L6:
	;
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v19)))
	v35 = m.G342
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v35)))
	v37 = m.G343
	v38 = m.T0[v36].(func(*base.Module, int32) int32)(m, v34)
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L2
	} else {
		goto L10
	}
L7:
	;
	if v30 != 0 {
		goto L6
	} else {
		goto L8
	}
L8:
	;
	v32 = F_lua_error(m, l0)
	mBase = m.M
	v33 = m.ExcPending
	if v33 != 0 {
		goto L2
	} else {
		goto L9
	}
L9:
	;
	v326 = v32
	goto L5
L10:
	;
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v37)))
	v41 = m.T0[v40].(func(*base.Module, int32) int32)(m, v38)
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L2
	} else {
		goto L11
	}
L11:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v19)))
	v44 = m.G66
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v44)))
	v46 = m.T0[v45].(func(*base.Module, int32) int32)(m, v43)
	mBase = m.M
	v47 = m.ExcPending
	if v47 != 0 {
		goto L2
	} else {
		goto L12
	}
L12:
	;
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v19)))
	v49 = m.G17
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v49)))
	m.T0[v50].(func(*base.Module, int32, int32))(m, v48, v38)
	mBase = m.M
	v52 = m.ExcPending
	if v52 != 0 {
		goto L2
	} else {
		goto L13
	}
L13:
	;
	v53 = m.G347
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v14)+12))
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v53)))
	v58 = m.T0[v57].(func(*base.Module, int32, int32, int32, int32, int32) int32)(m, v41, v30, v55, v46, int32(0))
	mBase = m.M
	v59 = m.ExcPending
	if v59 != 0 {
		goto L2
	} else {
		goto L15
	}
L14:
	;
	v202 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v202)+8)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v202))) = base.B2i32(v199 != int32(0))
	v208 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v208 + int32(16)
	goto L42
L15:
	;
	if v58 == int32(0) {
		v199 = int32(1)
		goto L14
	} else {
		goto L16
	}
L16:
	;
	goto L17
L17:
	;
	v64 = *(*int32)(unsafe.Add(mBase, _consts[9]))
	if v64 == int32(2) {
		v199 = int32(0)
		goto L14
	} else {
		goto L18
	}
L18:
	;
	if v64 != int32(28) {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	v187 = m.G3
	v193 = m.G8
	v194 = *(*int32)(unsafe.Add(mBase, uint32(v193)))
	m.T0[v194].(func(*base.Module, int32, int32, int32))(m, v187+int32(_a2218), v187+int32(_a2211), int32(1406))
	mBase = m.M
	v196 = m.ExcPending
	if v196 != 0 {
		goto L2
	} else {
		goto L41
	}
L20:
	;
	v69 = m.G3
	F_luaPushErrorBuff(m, l0, v69+int32(_a2219))
	mBase = m.M
	v73 = m.ExcPending
	if v73 != 0 {
		goto L2
	} else {
		goto L21
	}
L21:
	;
	v74 = m.G337
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v74)))
	m.T0[v75].(func(*base.Module, int32))(m, v41)
	mBase = m.M
	v77 = m.ExcPending
	if v77 != 0 {
		goto L2
	} else {
		goto L22
	}
L22:
	;
	if v55 < int32(1) {
		goto L23
	} else {
		goto L24
	}
L23:
	;
	v181 = m.G11
	v182 = *(*int32)(unsafe.Add(mBase, uint32(v181)))
	m.T0[v182].(func(*base.Module, int32))(m, v30)
	mBase = m.M
	v184 = m.ExcPending
	if v184 != 0 {
		goto L2
	} else {
		goto L39
	}
L24:
	;
	v80 = *(*int32)(unsafe.Add(mBase, uint32(v19)))
	v82 = v55 & int32(3)
	v83 = int32(0)
	if base.Ui32(v55) < base.Ui32(int32(4)) {
		v135 = v83
		goto L25
	} else {
		goto L26
	}
L25:
	;
	if v82 == int32(0) {
		goto L23
	} else {
		goto L34
	}
L26:
	;
	v89 = int32(0)
	v93 = v89
	v99 = v89
	goto L27
L27:
	;
	v104 = v30 + v93<<(uint(int32(2))%32)
	v105 = *(*int32)(unsafe.Add(mBase, uint32(v104)))
	v106 = m.G17
	v107 = *(*int32)(unsafe.Add(mBase, uint32(v106)))
	m.T0[v107].(func(*base.Module, int32, int32))(m, v80, v105)
	mBase = m.M
	v109 = m.ExcPending
	if v109 != 0 {
		goto L2
	} else {
		goto L29
	}
L28:
	;
	v135 = v129
	goto L25
L29:
	;
	v112 = *(*int32)(unsafe.Add(mBase, uint32(v104+int32(4))))
	v113 = *(*int32)(unsafe.Add(mBase, uint32(v106)))
	m.T0[v113].(func(*base.Module, int32, int32))(m, v80, v112)
	mBase = m.M
	v115 = m.ExcPending
	if v115 != 0 {
		goto L2
	} else {
		goto L30
	}
L30:
	;
	v118 = *(*int32)(unsafe.Add(mBase, uint32(v104+int32(8))))
	v119 = *(*int32)(unsafe.Add(mBase, uint32(v106)))
	m.T0[v119].(func(*base.Module, int32, int32))(m, v80, v118)
	mBase = m.M
	v121 = m.ExcPending
	if v121 != 0 {
		goto L2
	} else {
		goto L31
	}
L31:
	;
	v124 = *(*int32)(unsafe.Add(mBase, uint32(v104+int32(12))))
	v125 = *(*int32)(unsafe.Add(mBase, uint32(v106)))
	m.T0[v125].(func(*base.Module, int32, int32))(m, v80, v124)
	mBase = m.M
	v127 = m.ExcPending
	if v127 != 0 {
		goto L2
	} else {
		goto L32
	}
L32:
	;
	v128 = int32(4)
	v129 = v93 + v128
	v131 = v99 + v128
	if v131 != v55&int32(2147483644) {
		v93 = v129
		v99 = v131
		goto L27
	} else {
		goto L33
	}
L33:
	;
	goto L28
L34:
	;
	v148 = v135
	v153 = v83
	goto L35
L35:
	;
	v160 = *(*int32)(unsafe.Add(mBase, uint32(v30+v148<<(uint(int32(2))%32))))
	v161 = m.G17
	v162 = *(*int32)(unsafe.Add(mBase, uint32(v161)))
	m.T0[v162].(func(*base.Module, int32, int32))(m, v80, v160)
	mBase = m.M
	v164 = m.ExcPending
	if v164 != 0 {
		goto L2
	} else {
		goto L37
	}
L36:
	;
	goto L23
L37:
	;
	v165 = int32(1)
	v168 = v153 + v165
	if v168 != v82 {
		v148 = v148 + v165
		v153 = v168
		goto L35
	} else {
		goto L38
	}
L38:
	;
	goto L36
L39:
	;
	v185 = F_lua_error(m, l0)
	mBase = m.M
	v186 = m.ExcPending
	if v186 != 0 {
		goto L2
	} else {
		goto L40
	}
L40:
	;
	v326 = v185
	goto L5
L41:
	;
	m.Env.Exit(m, int32(1))
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L42:
	;
	v212 = m.G337
	v213 = *(*int32)(unsafe.Add(mBase, uint32(v212)))
	m.T0[v213].(func(*base.Module, int32))(m, v41)
	mBase = m.M
	v215 = m.ExcPending
	if v215 != 0 {
		goto L2
	} else {
		goto L43
	}
L43:
	;
	if v55 < int32(1) {
		goto L44
	} else {
		goto L45
	}
L44:
	;
	v319 = m.G11
	v320 = *(*int32)(unsafe.Add(mBase, uint32(v319)))
	m.T0[v320].(func(*base.Module, int32))(m, v30)
	mBase = m.M
	v322 = m.ExcPending
	if v322 != 0 {
		goto L2
	} else {
		goto L60
	}
L45:
	;
	v218 = *(*int32)(unsafe.Add(mBase, uint32(v19)))
	v220 = v55 & int32(3)
	v221 = int32(0)
	if base.Ui32(v55) < base.Ui32(int32(4)) {
		v273 = v221
		goto L46
	} else {
		goto L47
	}
L46:
	;
	if v220 == int32(0) {
		goto L44
	} else {
		goto L55
	}
L47:
	;
	v227 = int32(0)
	v229 = v227
	v231 = v227
	goto L48
L48:
	;
	v242 = v30 + v231<<(uint(int32(2))%32)
	v243 = *(*int32)(unsafe.Add(mBase, uint32(v242)))
	v244 = m.G17
	v245 = *(*int32)(unsafe.Add(mBase, uint32(v244)))
	m.T0[v245].(func(*base.Module, int32, int32))(m, v218, v243)
	mBase = m.M
	v247 = m.ExcPending
	if v247 != 0 {
		goto L2
	} else {
		goto L50
	}
L49:
	;
	v273 = v267
	goto L46
L50:
	;
	v250 = *(*int32)(unsafe.Add(mBase, uint32(v242+int32(4))))
	v251 = *(*int32)(unsafe.Add(mBase, uint32(v244)))
	m.T0[v251].(func(*base.Module, int32, int32))(m, v218, v250)
	mBase = m.M
	v253 = m.ExcPending
	if v253 != 0 {
		goto L2
	} else {
		goto L51
	}
L51:
	;
	v256 = *(*int32)(unsafe.Add(mBase, uint32(v242+int32(8))))
	v257 = *(*int32)(unsafe.Add(mBase, uint32(v244)))
	m.T0[v257].(func(*base.Module, int32, int32))(m, v218, v256)
	mBase = m.M
	v259 = m.ExcPending
	if v259 != 0 {
		goto L2
	} else {
		goto L52
	}
L52:
	;
	v262 = *(*int32)(unsafe.Add(mBase, uint32(v242+int32(12))))
	v263 = *(*int32)(unsafe.Add(mBase, uint32(v244)))
	m.T0[v263].(func(*base.Module, int32, int32))(m, v218, v262)
	mBase = m.M
	v265 = m.ExcPending
	if v265 != 0 {
		goto L2
	} else {
		goto L53
	}
L53:
	;
	v266 = int32(4)
	v267 = v231 + v266
	v269 = v229 + v266
	if v269 != v55&int32(2147483644) {
		v229 = v269
		v231 = v267
		goto L48
	} else {
		goto L54
	}
L54:
	;
	goto L49
L55:
	;
	v286 = v273
	v291 = v221
	goto L56
L56:
	;
	v298 = *(*int32)(unsafe.Add(mBase, uint32(v30+v286<<(uint(int32(2))%32))))
	v299 = m.G17
	v300 = *(*int32)(unsafe.Add(mBase, uint32(v299)))
	m.T0[v300].(func(*base.Module, int32, int32))(m, v218, v298)
	mBase = m.M
	v302 = m.ExcPending
	if v302 != 0 {
		goto L2
	} else {
		goto L58
	}
L57:
	;
	goto L44
L58:
	;
	v303 = int32(1)
	v306 = v291 + v303
	if v306 != v220 {
		v286 = v286 + v303
		v291 = v306
		goto L56
	} else {
		goto L59
	}
L59:
	;
	goto L57
L60:
	;
	v326 = int32(1)
	goto L5
L61:
	;
	m.Env.Exit(m, int32(1))
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_luaRedisCallCommand(m *base.Module, l0 int32) int32 {
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	v3 = F_luaServerGenericCommand(m, l0, int32(1))
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		return v3
	}
}
func F_luaRedisErrorReplyCommand(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v83 int32
	_ = v83
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v92 int32
	_ = v92
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v104 int32
	_ = v104
	var v107 int32
	_ = v107
	var v111 int32
	_ = v111
	var v116 int32
	_ = v116
	var v118 int32
	_ = v118
	var v122 int32
	_ = v122
	var v128 int32
	_ = v128
	var v131 int32
	_ = v131
	var v137 int32
	_ = v137
	var v141 int32
	_ = v141
	var v143 int32
	_ = v143
	var v151 int32
	_ = v151
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
	var v160 int32
	_ = v160
	var v162 int32
	_ = v162
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v169 int32
	_ = v169
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	goto L4
L1:
	;
	m.G0 = v8 + int32(16)
	return int32(1)
L2:
	;
	v86 = F_lua_tolstring(m, l0, int32(-1), int32(0))
	mBase = m.M
	v87 = m.ExcPending
	if v87 != 0 {
		goto L22
	} else {
		goto L26
	}
L3:
	;
	v77 = m.G3
	F_luaPushErrorBuff(m, l0, v77+int32(_a2215))
	mBase = m.M
	v83 = m.ExcPending
	if v83 != 0 {
		goto L22
	} else {
		goto L23
	}
L4:
	;
	if (v10-v11)>>(uint(int32(4))%32) != int32(1) {
		goto L3
	} else {
		goto L5
	}
L5:
	;
	goto L8
L6:
	;
	if v74 == int32(4) {
		goto L2
	} else {
		goto L21
	}
L7:
	;
	v68 = m.G398
	if v34 != v68 {
		goto L19
	} else {
		goto L20
	}
L8:
	;
	goto L12
L12:
	;
	v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v34 = v31 + int32(-16)
	goto L7
L19:
	;
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v34)+8))
	v74 = v71
	goto L6
L20:
	;
	v74 = int32(-1)
	goto L6
L21:
	;
	goto L3
L22:
	;
	return int32(0)
L23:
	;
	goto L1
L24:
	;
	F_luaPushErrorBuff(m, l0, v162)
	mBase = m.M
	v165 = m.ExcPending
	if v165 != 0 {
		goto L22
	} else {
		goto L49
	}
L25:
	;
	if v86&int32(3) == int32(0) {
		v118 = v86
		goto L31
	} else {
		goto L32
	}
L26:
	;
	v88 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v86))))
	if v88 == int32(45) {
		goto L25
	} else {
		goto L27
	}
L27:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8))) = v86
	v92 = m.G3
	v95 = F_lm_asprintf(m, v92+int32(_a704), v8)
	mBase = m.M
	v96 = m.ExcPending
	if v96 != 0 {
		goto L22
	} else {
		goto L28
	}
L28:
	;
	v162 = v95
	goto L24
L29:
	;
	v153 = v151 + int32(1)
	v154 = m.G22
	v155 = *(*int32)(unsafe.Add(mBase, uint32(v154)))
	v156 = m.T0[v155].(func(*base.Module, int32) int32)(m, v153)
	mBase = m.M
	v157 = m.ExcPending
	if v157 != 0 {
		goto L22
	} else {
		goto L45
	}
L30:
	;
	v151 = v143 - v86
	goto L29
L31:
	;
	v122 = v118
	goto L39
L32:
	;
	v104 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v86))))
	if v104 != 0 {
		goto L33
	} else {
		goto L34
	}
L33:
	;
	v107 = v86
	goto L35
L34:
	;
	v151 = v86 - v86
	goto L29
L35:
	;
	v111 = v107 + int32(1)
	if v111&int32(3) == int32(0) {
		v118 = v111
		goto L31
	} else {
		goto L37
	}
L37:
	;
	v116 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v111))))
	if v116 != 0 {
		v107 = v111
		goto L35
	} else {
		goto L38
	}
L38:
	;
	v143 = v111
	goto L30
L39:
	;
	v128 = *(*int32)(unsafe.Add(mBase, uint32(v122)))
	v131 = int32(-2139062144)
	if (int32(16843008)-v128|v128)&v131 == v131 {
		v122 = v122 + int32(4)
		goto L39
	} else {
		goto L41
	}
L40:
	;
	v137 = v122
	goto L42
L41:
	;
	goto L40
L42:
	;
	v141 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v137))))
	if v141 != 0 {
		v137 = v137 + int32(1)
		goto L42
	} else {
		goto L44
	}
L43:
	;
	v143 = v137
	goto L30
L44:
	;
	goto L43
L45:
	;
	if v153 == int32(0) {
		goto L47
	} else {
		goto L48
	}
L46:
	;
	v162 = v156
	goto L24
L47:
	;
	goto L46
L48:
	;
	v160 = F__emscripten_memcpy_bulkmem(m, v156, v86, v153)
	mBase = m.M
	goto L47
L49:
	;
	v166 = m.G11
	v167 = *(*int32)(unsafe.Add(mBase, uint32(v166)))
	m.T0[v167].(func(*base.Module, int32))(m, v162)
	mBase = m.M
	v169 = m.ExcPending
	if v169 != 0 {
		goto L22
	} else {
		goto L50
	}
L50:
	;
	goto L1
}
func F_luaRedisPCallCommand(m *base.Module, l0 int32) int32 {
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	v3 = F_luaServerGenericCommand(m, l0, int32(0))
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		return v3
	}
}
func F_luaRedisSetReplCommand(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v21 int32
	_ = v21
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v30 float64
	_ = v30
	var v31 int32
	_ = v31
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v53 int32
	_ = v53
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	v5 = m.G3
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v13 = F_luaGetFromRegistry(m, l0, v5+int32(_a2168))
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		return int32(0)
	} else {
		if v13 == int32(0) {
			v53 = m.G3
			v59 = m.G8
			v60 = *(*int32)(unsafe.Add(mBase, uint32(v59)))
			m.T0[v60].(func(*base.Module, int32, int32, int32))(m, v53+int32(_a2210), v53+int32(_a2211), int32(1363))
			mBase = m.M
			v62 = m.ExcPending
			if v62 != 0 {
				return int32(0)
			} else {
				m.Env.Exit(m, int32(1))
				mBase = m.M
				base.Wasm_trap_unreachable()
				for {
				}
			}
		} else {
			if (v6-v7)>>(uint(int32(4))%32) == int32(1) {
				v30 = F_lua_tonumber(m, l0, int32(-1))
				mBase = m.M
				v31 = m.ExcPending
				if v31 != 0 {
					return int32(0)
				} else {
					if base.F64_lt(base.F64_abs(v30), float64(2.147483648e+09)) == int32(0) {
						v39 = int32(-2147483648)
					} else {
						v37 = base.I32_trunc_f64_s(v30)
						v39 = v37
					}
					if base.Ui32(v39) < base.Ui32(int32(4)) {
						*(*int32)(unsafe.Add(mBase, uint32(v13)+12)) = v39
						return int32(0)
					} else {
						v42 = m.G3
						F_luaPushErrorBuff(m, l0, v42+int32(_a2216))
						mBase = m.M
						v46 = m.ExcPending
						if v46 != 0 {
							return int32(0)
						} else {
							v47 = F_lua_error(m, l0)
							mBase = m.M
							v48 = m.ExcPending
							if v48 != 0 {
								return int32(0)
							} else {
								return v47
							}
						}
					}
				}
			} else {
				v21 = m.G3
				F_luaPushErrorBuff(m, l0, v21+int32(_a2217))
				mBase = m.M
				v25 = m.ExcPending
				if v25 != 0 {
					return int32(0)
				} else {
					v26 = F_lua_error(m, l0)
					mBase = m.M
					v27 = m.ExcPending
					if v27 != 0 {
						return int32(0)
					} else {
						return v26
					}
				}
			}
		}
	}
}
func F_luaRedisSha1hexCommand(m *base.Module, l0 int32) int32 {
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
	var v16 int32
	_ = v16
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
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v60 int32
	_ = v60
	var v65 int32
	_ = v65
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	v4 = m.G0
	v6 = v4 - int32(64)
	m.G0 = v6
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if (v9-v10)>>(uint(int32(4))%32) == int32(1) {
		v26 = v6 + int32(16)
		v30 = F_lua_tolstring(m, l0, int32(1), v6+int32(12))
		mBase = m.M
		v31 = m.ExcPending
		if v31 != 0 {
			return int32(0)
		} else {
			v32 = *(*int32)(unsafe.Add(mBase, uint32(v6)+12))
			v35 = m.G0
			v37 = v35 - int32(112)
			m.G0 = v37
			v40 = v37 + int32(20)
			F_SHA1Init(m, v40)
			mBase = m.M
			F_SHA1Update(m, v40, v30, v32)
			mBase = m.M
			F_SHA1Final(m, v37, v40)
			mBase = m.M
			v51 = int32(0)
			for {
				v54 = int32(1)
				v56 = v26 + v51<<(uint(v54)%32)
				v60 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v37+v51))))
				v65 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v60&int32(15))+uint32(_consts[259]))))
				*(*uint8)(unsafe.Add(mBase, uint32(v56+v54))) = uint8(v65)
				v71 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v60)>>(uint(int32(4))%32)))+uint32(_consts[259]))))
				*(*uint8)(unsafe.Add(mBase, uint32(v56))) = uint8(v71)
				v74 = v51 + v54
				if v74 != int32(20) {
					v51 = v74
					continue
				} else {
					break
				}
				break
			}
			v77 = int32(0)
			*(*uint8)(unsafe.Add(mBase, uint32(v26)+40)) = uint8(v77)
			m.G0 = v37 + int32(112)
			F_lua_pushstring(m, l0, v6+int32(16))
			mBase = m.M
			v85 = m.ExcPending
			if v85 != 0 {
				return int32(0)
			} else {
				v86 = int32(1)
				m.G0 = v6 + int32(64)
				return v86
			}
		}
	} else {
		v16 = m.G3
		F_luaPushErrorBuff(m, l0, v16+int32(_a2214))
		mBase = m.M
		v22 = m.ExcPending
		if v22 != 0 {
			return int32(0)
		} else {
			v23 = F_lua_error(m, l0)
			mBase = m.M
			v24 = m.ExcPending
			if v24 != 0 {
				return int32(0)
			} else {
				v86 = v23
				m.G0 = v6 + int32(64)
				return v86
			}
		}
	}
}
func F_luaRedisStatusReplyCommand(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v86 int32
	_ = v86
	var v104 int32
	_ = v104
	var v107 int32
	_ = v107
	var v143 int32
	_ = v143
	var v144 int64
	_ = v144
	var v146 int32
	_ = v146
	var v148 int32
	_ = v148
	var v154 int32
	_ = v154
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if (v2-v3)>>(uint(int32(4))%32) != int32(1) {
		v69 = m.G3
		F_luaPushErrorBuff(m, l0, v69+int32(_a2215))
		mBase = m.M
		v75 = m.ExcPending
		if v75 != 0 {
			return int32(0)
		} else {
			return int32(1)
		}
	} else {
		v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		v26 = v23 + int32(-16)
		v60 = m.G398
		if v26 != v60 {
			v63 = *(*int32)(unsafe.Add(mBase, uint32(v26)+8))
			v66 = v63
		} else {
			v66 = int32(-1)
		}
		if v66 == int32(4) {
			v78 = int32(0)
			F_lua_createtable(m, l0, v78, v78)
			mBase = m.M
			v81 = m.ExcPending
			if v81 != 0 {
				return int32(0)
			} else {
				v82 = m.G3
				F_lua_pushstring(m, l0, v82+int32(_a294))
				mBase = m.M
				v86 = m.ExcPending
				if v86 != 0 {
					return int32(0)
				} else {
					v104 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
					v107 = v104 + int32(-48)
					v143 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
					v144 = *(*int64)(unsafe.Add(mBase, uint32(v107)))
					*(*int64)(unsafe.Add(mBase, uint32(v143))) = v144
					v146 = *(*int32)(unsafe.Add(mBase, uint32(v107)+8))
					*(*int32)(unsafe.Add(mBase, uint32(v143)+8)) = v146
					v148 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
					*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v148 + int32(16)
					F_lua_settable(m, l0, int32(-3))
					mBase = m.M
					v154 = m.ExcPending
					if v154 != 0 {
						return int32(0)
					} else {
						return int32(1)
					}
				}
			}
		} else {
			v69 = m.G3
			F_luaPushErrorBuff(m, l0, v69+int32(_a2215))
			mBase = m.M
			v75 = m.ExcPending
			if v75 != 0 {
				return int32(0)
			} else {
				return int32(1)
			}
		}
	}
}
func F_luaRegisterVersion(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v16 int32
	_ = v16
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v50 int32
	_ = v50
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	v4 = m.G3
	F_lua_pushstring(m, l1, v4+int32(_a2163))
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		return
	} else {
		v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
		v12 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
		*(*int32)(unsafe.Add(mBase, uint32(v12)+8)) = int32(3)
		*(*float64)(unsafe.Add(mBase, uint32(v12))) = base.F64_convert_i32_u(v9)
		v16 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
		*(*int32)(unsafe.Add(mBase, uint32(l1)+8)) = v16 + int32(16)
		F_lua_settable(m, l1, int32(-3))
		mBase = m.M
		v22 = m.ExcPending
		if v22 != 0 {
			return
		} else {
			F_lua_pushstring(m, l1, v4+int32(_a2164))
			mBase = m.M
			v26 = m.ExcPending
			if v26 != 0 {
				return
			} else {
				v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				F_lua_pushstring(m, l1, v27)
				mBase = m.M
				v29 = m.ExcPending
				if v29 != 0 {
					return
				} else {
					F_lua_settable(m, l1, int32(-3))
					mBase = m.M
					v32 = m.ExcPending
					if v32 != 0 {
						return
					} else {
						F_lua_pushstring(m, l1, v4+int32(_a2165))
						mBase = m.M
						v36 = m.ExcPending
						if v36 != 0 {
							return
						} else {
							v37 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
							v40 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
							*(*int32)(unsafe.Add(mBase, uint32(v40)+8)) = int32(3)
							*(*float64)(unsafe.Add(mBase, uint32(v40))) = base.F64_convert_i32_u(v37)
							v44 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
							*(*int32)(unsafe.Add(mBase, uint32(l1)+8)) = v44 + int32(16)
							F_lua_settable(m, l1, int32(-3))
							mBase = m.M
							v50 = m.ExcPending
							if v50 != 0 {
								return
							} else {
								F_lua_pushstring(m, l1, v4+int32(_a2166))
								mBase = m.M
								v54 = m.ExcPending
								if v54 != 0 {
									return
								} else {
									v55 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
									F_lua_pushstring(m, l1, v55)
									mBase = m.M
									v57 = m.ExcPending
									if v57 != 0 {
										return
									} else {
										F_lua_settable(m, l1, int32(-3))
										mBase = m.M
										v60 = m.ExcPending
										if v60 != 0 {
											return
										} else {
											F_lua_pushstring(m, l1, v4+int32(_a2167))
											mBase = m.M
											v64 = m.ExcPending
											if v64 != 0 {
												return
											} else {
												v65 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
												F_lua_pushstring(m, l1, v65)
												mBase = m.M
												v67 = m.ExcPending
												if v67 != 0 {
													return
												} else {
													F_lua_settable(m, l1, int32(-3))
													mBase = m.M
													v70 = m.ExcPending
													if v70 != 0 {
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
func F_luaSaveOnRegistry(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v9 int32
	_ = v9
	var v13 int32
	_ = v13
	var v18 int32
	_ = v18
	var v26 int32
	_ = v26
	F_lua_pushstring(m, l0, l1)
	mBase = m.M
	v5 = m.ExcPending
	if v5 != 0 {
		return
	} else {
		if l2 == int32(0) {
			v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			*(*int32)(unsafe.Add(mBase, uint32(v18)+8)) = int32(0)
			*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v18 + int32(16)
		} else {
			v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			*(*int32)(unsafe.Add(mBase, uint32(v9)+8)) = int32(2)
			*(*int32)(unsafe.Add(mBase, uint32(v9))) = l2
			v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v13 + int32(16)
		}
		F_lua_settable(m, l0, int32(-10000))
		mBase = m.M
		v26 = m.ExcPending
		if v26 != 0 {
			return
		} else {
			return
		}
	}
}
func F_luaServerDebugCommand(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v68 int32
	_ = v68
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v10 = m.G6
	v11 = *(*int32)(unsafe.Add(mBase, uint32(v10)))
	goto L2
L1:
	;
	m.G0 = v8 + int32(16)
	return int32(0)
L2:
	;
	if v11 == int32(0) {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v14 = m.G15
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v19 = (v15 - v16) >> (uint(int32(4)) % 32)
	goto L4
L4:
	;
	v20 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
	v21 = m.G6
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v21)+280))
	goto L5
L5:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8))) = v22
	v25 = m.G3
	v28 = m.T0[v20].(func(*base.Module, int32, int32, int32) int32)(m, int32(0), v25+int32(_a2132), v8)
	mBase = m.M
	v31 = m.ExcPending
	if v31 != 0 {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	return int32(0)
L7:
	;
	if v19 == int32(0) {
		v64 = v28
		goto L8
	} else {
		goto L9
	}
L8:
	;
	F_ldbLog(m, v64)
	mBase = m.M
	v68 = m.ExcPending
	if v68 != 0 {
		goto L6
	} else {
		goto L17
	}
L9:
	;
	v36 = F_ldbCatStackValue(m, v28, l0, int32(0)-v19)
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		goto L6
	} else {
		goto L10
	}
L10:
	;
	v39 = v19 + int32(-1)
	if v39 == int32(0) {
		v64 = v36
		goto L8
	} else {
		goto L11
	}
L11:
	;
	v44 = v36
	v45 = v39
	goto L12
L12:
	;
	v48 = m.G3
	v52 = m.G16
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v52)))
	v54 = m.T0[v53].(func(*base.Module, int32, int32, int32, int32) int32)(m, int32(0), v44, v48+int32(_a438), int32(2))
	mBase = m.M
	v55 = m.ExcPending
	if v55 != 0 {
		goto L6
	} else {
		goto L14
	}
L13:
	;
	v64 = v58
	goto L8
L14:
	;
	v58 = F_ldbCatStackValue(m, v44, l0, int32(0)-v45)
	mBase = m.M
	v59 = m.ExcPending
	if v59 != 0 {
		goto L6
	} else {
		goto L15
	}
L15:
	;
	v61 = v45 + int32(-1)
	if v61 != 0 {
		v44 = v58
		v45 = v61
		goto L12
	} else {
		goto L16
	}
L16:
	;
	goto L13
L17:
	;
	goto L1
}
func F_luaServerGenericCommand(m *base.Module, l0 int32, l1 int32) int32 {
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
	var v23 int32
	_ = v23
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
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
	var v81 int32
	_ = v81
	var v86 int32
	_ = v86
	var v88 int64
	_ = v88
	var v90 int32
	_ = v90
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v120 int32
	_ = v120
	var v124 int32
	_ = v124
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v142 int32
	_ = v142
	var v144 int32
	_ = v144
	var v153 int32
	_ = v153
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v163 int32
	_ = v163
	var v172 int32
	_ = v172
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v185 int32
	_ = v185
	var v187 int32
	_ = v187
	var v198 int32
	_ = v198
	var v201 int32
	_ = v201
	var v207 int32
	_ = v207
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
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
	var v231 int32
	_ = v231
	var v238 int32
	_ = v238
	var v241 int32
	_ = v241
	var v247 int32
	_ = v247
	var v248 int32
	_ = v248
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v252 int32
	_ = v252
	var v255 int32
	_ = v255
	var v256 int32
	_ = v256
	var v258 int32
	_ = v258
	var v261 int32
	_ = v261
	var v262 int32
	_ = v262
	var v264 int32
	_ = v264
	var v267 int32
	_ = v267
	var v268 int32
	_ = v268
	var v270 int32
	_ = v270
	var v271 int32
	_ = v271
	var v272 int32
	_ = v272
	var v274 int32
	_ = v274
	var v281 int32
	_ = v281
	var v295 int32
	_ = v295
	var v299 int32
	_ = v299
	var v305 int32
	_ = v305
	var v306 int32
	_ = v306
	var v307 int32
	_ = v307
	var v309 int32
	_ = v309
	var v310 int32
	_ = v310
	var v313 int32
	_ = v313
	var v327 int32
	_ = v327
	var v328 int32
	_ = v328
	var v330 int32
	_ = v330
	var v331 int32
	_ = v331
	var v334 int32
	_ = v334
	var v338 int32
	_ = v338
	var v341 int32
	_ = v341
	var v356 int32
	_ = v356
	var v357 int32
	_ = v357
	var v363 int32
	_ = v363
	var v374 int32
	_ = v374
	var v380 int32
	_ = v380
	var v381 int32
	_ = v381
	var v383 int32
	_ = v383
	var v386 int32
	_ = v386
	var v392 int32
	_ = v392
	var v393 int32
	_ = v393
	var v395 int32
	_ = v395
	v13 = m.G0
	v15 = v13 - int32(2112)
	m.G0 = v15
	v17 = m.G3
	v20 = F_luaGetFromRegistry(m, l0, v17+int32(_a2168))
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	v386 = m.G3
	v392 = m.G8
	v393 = *(*int32)(unsafe.Add(mBase, uint32(v392)))
	m.T0[v393].(func(*base.Module, int32, int32, int32))(m, v386+int32(_a2220), v386+int32(_a2211), int32(1216))
	mBase = m.M
	v395 = m.ExcPending
	if v395 != 0 {
		goto L3
	} else {
		goto L70
	}
L2:
	;
	v374 = m.G3
	v380 = m.G8
	v381 = *(*int32)(unsafe.Add(mBase, uint32(v380)))
	m.T0[v381].(func(*base.Module, int32, int32, int32))(m, v374+int32(_a2210), v374+int32(_a2211), int32(1140))
	mBase = m.M
	v383 = m.ExcPending
	if v383 != 0 {
		goto L3
	} else {
		goto L69
	}
L3:
	;
	return int32(0)
L4:
	;
	if v20 == int32(0) {
		goto L2
	} else {
		goto L5
	}
L5:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+2108)) = int32(0)
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v20)))
	v31 = F_luaArgsToServerArgv(m, v28, l0, v15+int32(2108))
	mBase = m.M
	v32 = m.ExcPending
	if v32 != 0 {
		goto L3
	} else {
		goto L9
	}
L6:
	;
	m.G0 = v15 + int32(2112)
	return v363
L7:
	;
	v356 = F_lua_error(m, l0)
	mBase = m.M
	v357 = m.ExcPending
	if v357 != 0 {
		goto L3
	} else {
		goto L68
	}
L8:
	;
	v34 = m.G3
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v34)+uint32(_consts[1194])))
	if v37 == int32(0) {
		goto L12
	} else {
		goto L13
	}
L9:
	;
	if v31 != 0 {
		goto L8
	} else {
		goto L10
	}
L10:
	;
	if l1 != 0 {
		goto L7
	} else {
		goto L11
	}
L11:
	;
	v363 = int32(1)
	goto L6
L12:
	;
	v58 = m.G3
	*(*int32)(unsafe.Add(mBase, uint32(v58)+uint32(_consts[1194]))) = int32(1)
	v63 = int32(0)
	v64 = m.G6
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v64)))
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v64)+264))
	goto L18
L13:
	;
	v40 = m.G3
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v20)))
	v43 = v40 + int32(_a2221)
	*(*int32)(unsafe.Add(mBase, uint32(v15)+32)) = v43
	v51 = m.G10
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v51)))
	m.T0[v52].(func(*base.Module, int32, int32, int32, int32))(m, v41, v40+int32(_a716), v40+int32(_a57), v15+int32(32))
	mBase = m.M
	v54 = m.ExcPending
	if v54 != 0 {
		goto L3
	} else {
		goto L14
	}
L14:
	;
	F_luaPushErrorBuff(m, l0, v43)
	mBase = m.M
	v56 = m.ExcPending
	if v56 != 0 {
		goto L3
	} else {
		goto L15
	}
L15:
	;
	v363 = int32(1)
	goto L6
L16:
	;
	v176 = *(*int32)(unsafe.Add(mBase, uint32(v20)+16))
	v177 = *(*int32)(unsafe.Add(mBase, uint32(v20)+12))
	v185 = F__emscripten_memset_bulkmem(m, v15+int32(52), base.I32_extend8_s(int32(0)), int32(2056))
	mBase = m.M
	goto L36
L17:
	;
	v75 = m.G9
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v75)))
	v77 = m.T0[v76].(func(*base.Module, int32, int32) int32)(m, int32(10), int32(1))
	mBase = m.M
	v78 = m.ExcPending
	if v78 != 0 {
		goto L3
	} else {
		goto L20
	}
L18:
	;
	if base.B2i32(v65 != v63)&base.B2i32(v68 != v63) != 0 {
		goto L17
	} else {
		goto L19
	}
L19:
	;
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v15)+2108))
	v172 = v72
	goto L16
L20:
	;
	v81 = m.G3
	v86 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v81)+uint32(_consts[1195]))))
	*(*uint16)(unsafe.Add(mBase, uint32(v77+int32(8)))) = uint16(v86)
	v88 = *(*int64)(unsafe.Add(mBase, uint32(v81)+uint32(_consts[1196])))
	*(*int64)(unsafe.Add(mBase, uint32(v77))) = v88
	v90 = *(*int32)(unsafe.Add(mBase, uint32(v15)+2108))
	if int32(1) <= v90 {
		goto L22
	} else {
		goto L23
	}
L21:
	;
	F_ldbLogCString(m, v153)
	mBase = m.M
	v159 = m.ExcPending
	if v159 != 0 {
		goto L3
	} else {
		goto L34
	}
L22:
	;
	v99 = int32(0)
	v100 = v77
	goto L24
L23:
	;
	v153 = v77
	goto L21
L24:
	;
	if v99 != int32(10) {
		goto L26
	} else {
		goto L27
	}
L25:
	;
	v153 = v138
	goto L21
L26:
	;
	v124 = *(*int32)(unsafe.Add(mBase, uint32(v31+v99<<(uint(int32(2))%32))))
	v126 = m.G7
	v127 = *(*int32)(unsafe.Add(mBase, uint32(v126)))
	v128 = m.T0[v127].(func(*base.Module, int32, int32) int32)(m, v124, int32(0))
	mBase = m.M
	v129 = m.ExcPending
	if v129 != 0 {
		goto L3
	} else {
		goto L30
	}
L27:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15))) = v100
	*(*int32)(unsafe.Add(mBase, uint32(v15)+4)) = v90 + int32(-11)
	v112 = m.G3
	v113 = m.G11
	v116 = F_lm_asprintf(m, v112+int32(_a2222), v15)
	mBase = m.M
	v117 = m.ExcPending
	if v117 != 0 {
		goto L3
	} else {
		goto L28
	}
L28:
	;
	v118 = *(*int32)(unsafe.Add(mBase, uint32(v113)))
	m.T0[v118].(func(*base.Module, int32))(m, v100)
	mBase = m.M
	v120 = m.ExcPending
	if v120 != 0 {
		goto L3
	} else {
		goto L29
	}
L29:
	;
	v153 = v116
	goto L21
L30:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+20)) = v128
	*(*int32)(unsafe.Add(mBase, uint32(v15)+16)) = v100
	v132 = m.G3
	v133 = m.G11
	v138 = F_lm_asprintf(m, v132+int32(_a423), v15+int32(16))
	mBase = m.M
	v139 = m.ExcPending
	if v139 != 0 {
		goto L3
	} else {
		goto L31
	}
L31:
	;
	v140 = *(*int32)(unsafe.Add(mBase, uint32(v133)))
	m.T0[v140].(func(*base.Module, int32))(m, v100)
	mBase = m.M
	v142 = m.ExcPending
	if v142 != 0 {
		goto L3
	} else {
		goto L32
	}
L32:
	;
	v144 = v99 + int32(1)
	if v144 != v90 {
		v99 = v144
		v100 = v138
		goto L24
	} else {
		goto L33
	}
L33:
	;
	goto L25
L34:
	;
	v160 = m.G11
	v161 = *(*int32)(unsafe.Add(mBase, uint32(v160)))
	m.T0[v161].(func(*base.Module, int32))(m, v153)
	mBase = m.M
	v163 = m.ExcPending
	if v163 != 0 {
		goto L3
	} else {
		goto L35
	}
L35:
	;
	v172 = v90
	goto L16
L36:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+44)) = l0
	v187 = *(*int32)(unsafe.Add(mBase, uint32(v20)))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+56)) = int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(v15)+48)) = v187
	goto L37
L37:
	;
	*(*int32)(unsafe.Add(mBase, _consts[9])) = int32(0)
	if v177&int32(1) != 0 {
		goto L38
	} else {
		goto L39
	}
L38:
	;
	v198 = int32(4929)
	goto L40
L39:
	;
	v198 = int32(4931)
	goto L40
L40:
	;
	v201 = int32(4)
	if v177 != 0 {
		goto L41
	} else {
		goto L42
	}
L41:
	;
	v207 = v198 | v177<<(uint(int32(1))%32)&v201 ^ v201
	goto L43
L42:
	;
	v207 = int32(4935)
	goto L43
L43:
	;
	if v176 == int32(3) {
		goto L44
	} else {
		goto L45
	}
L44:
	;
	v212 = v207 | int32(8)
	goto L46
L45:
	;
	v212 = v207
	goto L46
L46:
	;
	v213 = m.G3
	v218 = F_VM_CallArgv(m, v187, v31, v172, v212, v213+int32(_a2223), v15+int32(44))
	mBase = m.M
	v219 = m.ExcPending
	if v219 != 0 {
		goto L3
	} else {
		goto L47
	}
L47:
	;
	if v218 != 0 {
		goto L1
	} else {
		goto L48
	}
L48:
	;
	if v172 < int32(1) {
		goto L49
	} else {
		goto L50
	}
L49:
	;
	v327 = m.G11
	v328 = *(*int32)(unsafe.Add(mBase, uint32(v327)))
	m.T0[v328].(func(*base.Module, int32))(m, v31)
	mBase = m.M
	v330 = m.ExcPending
	if v330 != 0 {
		goto L3
	} else {
		goto L65
	}
L50:
	;
	v222 = *(*int32)(unsafe.Add(mBase, uint32(v20)))
	v224 = v172 & int32(3)
	v225 = int32(0)
	if base.Ui32(v172) < base.Ui32(int32(4)) {
		v281 = v225
		goto L51
	} else {
		goto L52
	}
L51:
	;
	if v224 == int32(0) {
		goto L49
	} else {
		goto L60
	}
L52:
	;
	v231 = int32(0)
	v238 = v231
	v241 = v231
	goto L53
L53:
	;
	v247 = v31 + v238<<(uint(int32(2))%32)
	v248 = *(*int32)(unsafe.Add(mBase, uint32(v247)))
	v249 = m.G17
	v250 = *(*int32)(unsafe.Add(mBase, uint32(v249)))
	m.T0[v250].(func(*base.Module, int32, int32))(m, v222, v248)
	mBase = m.M
	v252 = m.ExcPending
	if v252 != 0 {
		goto L3
	} else {
		goto L55
	}
L54:
	;
	v281 = v272
	goto L51
L55:
	;
	v255 = *(*int32)(unsafe.Add(mBase, uint32(v247+int32(4))))
	v256 = *(*int32)(unsafe.Add(mBase, uint32(v249)))
	m.T0[v256].(func(*base.Module, int32, int32))(m, v222, v255)
	mBase = m.M
	v258 = m.ExcPending
	if v258 != 0 {
		goto L3
	} else {
		goto L56
	}
L56:
	;
	v261 = *(*int32)(unsafe.Add(mBase, uint32(v247+int32(8))))
	v262 = *(*int32)(unsafe.Add(mBase, uint32(v249)))
	m.T0[v262].(func(*base.Module, int32, int32))(m, v222, v261)
	mBase = m.M
	v264 = m.ExcPending
	if v264 != 0 {
		goto L3
	} else {
		goto L57
	}
L57:
	;
	v267 = *(*int32)(unsafe.Add(mBase, uint32(v247+int32(12))))
	v268 = *(*int32)(unsafe.Add(mBase, uint32(v249)))
	m.T0[v268].(func(*base.Module, int32, int32))(m, v222, v267)
	mBase = m.M
	v270 = m.ExcPending
	if v270 != 0 {
		goto L3
	} else {
		goto L58
	}
L58:
	;
	v271 = int32(4)
	v272 = v238 + v271
	v274 = v241 + v271
	if v274 != v172&int32(2147483644) {
		v238 = v272
		v241 = v274
		goto L53
	} else {
		goto L59
	}
L59:
	;
	goto L54
L60:
	;
	v295 = v281
	v299 = v225
	goto L61
L61:
	;
	v305 = *(*int32)(unsafe.Add(mBase, uint32(v31+v295<<(uint(int32(2))%32))))
	v306 = m.G17
	v307 = *(*int32)(unsafe.Add(mBase, uint32(v306)))
	m.T0[v307].(func(*base.Module, int32, int32))(m, v222, v305)
	mBase = m.M
	v309 = m.ExcPending
	if v309 != 0 {
		goto L3
	} else {
		goto L63
	}
L62:
	;
	goto L49
L63:
	;
	v310 = int32(1)
	v313 = v299 + v310
	if v313 != v224 {
		v295 = v295 + v310
		v299 = v313
		goto L61
	} else {
		goto L64
	}
L64:
	;
	goto L62
L65:
	;
	v331 = m.G3
	v334 = *(*int32)(unsafe.Add(mBase, uint32(v331)+uint32(_consts[1194])))
	*(*int32)(unsafe.Add(mBase, uint32(v331)+uint32(_consts[1194]))) = v334 + int32(-1)
	v338 = int32(1)
	if l1 == int32(0) {
		v363 = v338
		goto L6
	} else {
		goto L66
	}
L66:
	;
	v341 = *(*int32)(unsafe.Add(mBase, uint32(v15)+52))
	if v341 == int32(0) {
		v363 = v338
		goto L6
	} else {
		goto L67
	}
L67:
	;
	goto L7
L68:
	;
	v363 = v356
	goto L6
L69:
	;
	m.Env.Exit(m, int32(1))
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L70:
	;
	m.Env.Exit(m, int32(1))
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_luaSetResp(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v21 int32
	_ = v21
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v30 float64
	_ = v30
	var v31 int32
	_ = v31
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v55 int32
	_ = v55
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	v5 = m.G3
	v8 = F_luaGetFromRegistry(m, l0, v5+int32(_a2168))
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return int32(0)
	} else {
		if v8 == int32(0) {
			v55 = m.G3
			v61 = m.G8
			v62 = *(*int32)(unsafe.Add(mBase, uint32(v61)))
			m.T0[v62].(func(*base.Module, int32, int32, int32))(m, v55+int32(_a2210), v55+int32(_a2211), int32(1478))
			mBase = m.M
			v64 = m.ExcPending
			if v64 != 0 {
				return int32(0)
			} else {
				m.Env.Exit(m, int32(1))
				mBase = m.M
				base.Wasm_trap_unreachable()
				for {
				}
			}
		} else {
			v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
			if (v14-v15)>>(uint(int32(4))%32) == int32(1) {
				v30 = F_lua_tonumber(m, l0, int32(-1))
				mBase = m.M
				v31 = m.ExcPending
				if v31 != 0 {
					return int32(0)
				} else {
					if base.F64_lt(base.F64_abs(v30), float64(2.147483648e+09)) == int32(0) {
						v39 = int32(-2147483648)
					} else {
						v37 = base.I32_trunc_f64_s(v30)
						v39 = v37
					}
					if base.Ui32(int32(-3)) < base.Ui32(v39+int32(-4)) {
						*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = v39
						return int32(0)
					} else {
						v44 = m.G3
						F_luaPushErrorBuff(m, l0, v44+int32(_a2212))
						mBase = m.M
						v48 = m.ExcPending
						if v48 != 0 {
							return int32(0)
						} else {
							v49 = F_lua_error(m, l0)
							mBase = m.M
							v50 = m.ExcPending
							if v50 != 0 {
								return int32(0)
							} else {
								return v49
							}
						}
					}
				}
			} else {
				v21 = m.G3
				F_luaPushErrorBuff(m, l0, v21+int32(_a2213))
				mBase = m.M
				v25 = m.ExcPending
				if v25 != 0 {
					return int32(0)
				} else {
					v26 = F_lua_error(m, l0)
					mBase = m.M
					v27 = m.ExcPending
					if v27 != 0 {
						return int32(0)
					} else {
						return v26
					}
				}
			}
		}
	}
}
func F_luaSetTableProtectionForBasicTypes(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v24 int32
	_ = v24
	var v30 int32
	_ = v30
	var v36 int32
	_ = v36
	var v41 int32
	_ = v41
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v58 int32
	_ = v58
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v118 int32
	_ = v118
	var v121 int32
	_ = v121
	var v124 int32
	_ = v124
	var v127 int32
	_ = v127
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v136 int32
	_ = v136
	var v140 int32
	_ = v140
	var v146 int32
	_ = v146
	var v150 int32
	_ = v150
	var v169 int32
	_ = v169
	var v197 int32
	_ = v197
	var v208 int32
	_ = v208
	v5 = int32(0)
	goto L1
L1:
	;
	v6 = m.G3
	v12 = *(*int32)(unsafe.Add(mBase, uint32(v6+int32(_a2162)+v5<<(uint(int32(2))%32))))
	switch v12 {
	case 0:
		goto L7
	case 1:
		goto L8
	case 2:
		goto L4
	case 3:
		goto L9
	case 4:
		goto L10
	default:
		goto L3
	case 6:
		goto L6
	case 8:
		goto L5
	}
L2:
	;
	return
L3:
	;
	goto L22
L4:
	;
	v54 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v54)+8)) = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v54))) = l0
	v58 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v58 + int32(16)
	goto L18
L5:
	;
	v51 = F_lua_newthread(m, l0)
	mBase = m.M
	v52 = m.ExcPending
	if v52 != 0 {
		goto L11
	} else {
		goto L17
	}
L6:
	;
	v47 = int32(0)
	F_lua_pushcclosure(m, l0, v47, v47)
	mBase = m.M
	v50 = m.ExcPending
	if v50 != 0 {
		goto L11
	} else {
		goto L16
	}
L7:
	;
	v41 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v41)+8)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v41 + int32(16)
	goto L15
L8:
	;
	v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v30)+8)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v30))) = int32(0)
	v36 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v36 + int32(16)
	goto L14
L9:
	;
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v20)+8)) = int32(3)
	*(*float64)(unsafe.Add(mBase, uint32(v20))) = float64(0)
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v24 + int32(16)
	goto L13
L10:
	;
	v13 = m.G3
	F_lua_pushstring(m, l0, v13+int32(_a188))
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	return
L12:
	;
	goto L3
L13:
	;
	goto L3
L14:
	;
	goto L3
L15:
	;
	goto L3
L16:
	;
	goto L3
L17:
	;
	goto L3
L18:
	;
	goto L3
L19:
	;
	goto L54
L20:
	;
	if v146 == int32(0) {
		goto L19
	} else {
		goto L42
	}
L21:
	;
	v118 = *(*int32)(unsafe.Add(mBase, uint32(v82)+8))
	switch v118 + int32(-5) {
	case 0:
		goto L39
	default:
		goto L37
	case 2:
		goto L38
	}
L22:
	;
	goto L28
L28:
	;
	v79 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v82 = v79 + int32(-16)
	goto L21
L36:
	;
	v134 = *(*int32)(unsafe.Add(mBase, uint32(v133)))
	if v134 != 0 {
		goto L40
	} else {
		goto L41
	}
L37:
	;
	v127 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v133 = v127 + v118<<(uint(int32(2))%32) + int32(152)
	goto L36
L38:
	;
	v124 = *(*int32)(unsafe.Add(mBase, uint32(v82)))
	v133 = v124 + int32(8)
	goto L36
L39:
	;
	v121 = *(*int32)(unsafe.Add(mBase, uint32(v82)))
	v133 = v121 + int32(16)
	goto L36
L40:
	;
	v136 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v136)+8)) = int32(5)
	*(*int32)(unsafe.Add(mBase, uint32(v136))) = v134
	v140 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v140 + int32(16)
	v146 = int32(1)
	goto L20
L41:
	;
	v146 = int32(0)
	goto L20
L42:
	;
	F_luaSetTableProtectionRecursively(m, l0)
	mBase = m.M
	v150 = m.ExcPending
	if v150 != 0 {
		goto L11
	} else {
		goto L43
	}
L43:
	;
	goto L46
L44:
	;
	goto L19
L45:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v169 + int32(-16)
	goto L44
L46:
	;
	v169 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	goto L45
L52:
	;
	v208 = v5 + int32(1)
	if v208 != int32(7) {
		v5 = v208
		goto L1
	} else {
		goto L60
	}
L53:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v197 + int32(-16)
	goto L52
L54:
	;
	v197 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	goto L53
L60:
	;
	goto L2
}
func F_luaSetTableProtectionRecursively(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v18 int32
	_ = v18
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v76 int32
	_ = v76
	var v115 int32
	_ = v115
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v121 int32
	_ = v121
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v147 int32
	_ = v147
	var v150 int32
	_ = v150
	var v184 int32
	_ = v184
	var v187 int32
	_ = v187
	var v190 int32
	_ = v190
	var v194 int32
	_ = v194
	var v213 int32
	_ = v213
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v244 int32
	_ = v244
	var v247 int32
	_ = v247
	var v283 int32
	_ = v283
	var v286 int32
	_ = v286
	var v289 int32
	_ = v289
	var v292 int32
	_ = v292
	var v298 int32
	_ = v298
	var v299 int32
	_ = v299
	var v301 int32
	_ = v301
	var v305 int32
	_ = v305
	var v311 int32
	_ = v311
	var v315 int32
	_ = v315
	var v334 int32
	_ = v334
	goto L4
L1:
	;
	return
L2:
	;
	if v58 != 0 {
		goto L1
	} else {
		goto L18
	}
L3:
	;
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v18+int32(-16))))
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v57)+8))
	goto L2
L4:
	;
	goto L10
L10:
	;
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	goto L3
L18:
	;
	goto L21
L19:
	;
	v118 = F_lua_checkstack(m, l0, int32(2))
	mBase = m.M
	v119 = m.ExcPending
	if v119 != 0 {
		goto L35
	} else {
		goto L36
	}
L20:
	;
	v115 = *(*int32)(unsafe.Add(mBase, uint32(v76+int32(-16))))
	*(*int32)(unsafe.Add(mBase, uint32(v115)+8)) = int32(1)
	goto L19
L21:
	;
	goto L27
L27:
	;
	v76 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	goto L20
L35:
	;
	return
L36:
	;
	v121 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v121)+8)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v121 + int32(16)
	goto L37
L37:
	;
	v128 = F_lua_next(m, l0, int32(-2))
	mBase = m.M
	v129 = m.ExcPending
	if v129 != 0 {
		goto L35
	} else {
		goto L39
	}
L38:
	;
	goto L73
L39:
	;
	if v128 == int32(0) {
		goto L38
	} else {
		goto L40
	}
L40:
	;
	goto L41
L41:
	;
	goto L46
L42:
	;
	goto L38
L43:
	;
	goto L63
L44:
	;
	if v190 != int32(5) {
		goto L43
	} else {
		goto L59
	}
L45:
	;
	v184 = m.G398
	if v150 != v184 {
		goto L57
	} else {
		goto L58
	}
L46:
	;
	goto L50
L50:
	;
	v147 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v150 = v147 + int32(-16)
	goto L45
L57:
	;
	v187 = *(*int32)(unsafe.Add(mBase, uint32(v150)+8))
	v190 = v187
	goto L44
L58:
	;
	v190 = int32(-1)
	goto L44
L59:
	;
	F_luaSetTableProtectionRecursively(m, l0)
	mBase = m.M
	v194 = m.ExcPending
	if v194 != 0 {
		goto L35
	} else {
		goto L60
	}
L60:
	;
	goto L43
L61:
	;
	v224 = F_lua_next(m, l0, int32(-2))
	mBase = m.M
	v225 = m.ExcPending
	if v225 != 0 {
		goto L35
	} else {
		goto L69
	}
L62:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v213 + int32(-16)
	goto L61
L63:
	;
	v213 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	goto L62
L69:
	;
	if v224 != 0 {
		goto L41
	} else {
		goto L70
	}
L70:
	;
	goto L42
L71:
	;
	if v311 == int32(0) {
		goto L1
	} else {
		goto L93
	}
L72:
	;
	v283 = *(*int32)(unsafe.Add(mBase, uint32(v247)+8))
	switch v283 + int32(-5) {
	case 0:
		goto L90
	default:
		goto L88
	case 2:
		goto L89
	}
L73:
	;
	goto L79
L79:
	;
	v244 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v247 = v244 + int32(-16)
	goto L72
L87:
	;
	v299 = *(*int32)(unsafe.Add(mBase, uint32(v298)))
	if v299 != 0 {
		goto L91
	} else {
		goto L92
	}
L88:
	;
	v292 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v298 = v292 + v283<<(uint(int32(2))%32) + int32(152)
	goto L87
L89:
	;
	v289 = *(*int32)(unsafe.Add(mBase, uint32(v247)))
	v298 = v289 + int32(8)
	goto L87
L90:
	;
	v286 = *(*int32)(unsafe.Add(mBase, uint32(v247)))
	v298 = v286 + int32(16)
	goto L87
L91:
	;
	v301 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v301)+8)) = int32(5)
	*(*int32)(unsafe.Add(mBase, uint32(v301))) = v299
	v305 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v305 + int32(16)
	v311 = int32(1)
	goto L71
L92:
	;
	v311 = int32(0)
	goto L71
L93:
	;
	F_luaSetTableProtectionRecursively(m, l0)
	mBase = m.M
	v315 = m.ExcPending
	if v315 != 0 {
		goto L35
	} else {
		goto L94
	}
L94:
	;
	goto L97
L95:
	;
	goto L1
L96:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v334 + int32(-16)
	goto L95
L97:
	;
	v334 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	goto L96
}
func F_lua_checkstack(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
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
	var v33 int32
	_ = v33
	var v38 int32
	_ = v38
	v3 = int32(0)
	if int32(8000) < l1 {
		v38 = v3
		return v38
	} else {
		v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
		if int32(8000) < (v9-v10)>>(uint(int32(4))%32)+l1 {
			v38 = v3
			return v38
		} else {
			v17 = int32(1)
			if l1 < v17 {
				v38 = v17
				return v38
			} else {
				v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
				v23 = l1 << (uint(int32(4)) % 32)
				if v23 < v20-v9 {
					v30 = v9
					v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
					v32 = *(*int32)(unsafe.Add(mBase, uint32(v31)+8))
					v33 = v30 + v23
					if base.Ui32(v33) <= base.Ui32(v32) {
						v38 = v17
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v31)+8)) = v33
						v38 = v17
					}
					return v38
				} else {
					F_luaD_growstack(m, l0, l1)
					mBase = m.M
					v28 = m.ExcPending
					if v28 != 0 {
						return int32(0)
					} else {
						v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
						v30 = v29
						v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
						v32 = *(*int32)(unsafe.Add(mBase, uint32(v31)+8))
						v33 = v30 + v23
						if base.Ui32(v33) <= base.Ui32(v32) {
							v38 = v17
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v31)+8)) = v33
							v38 = v17
						}
						return v38
					}
				}
			}
		}
	}
}
func F_lua_concat(m *base.Module, l0 int32, l1 int32) {
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
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v38 int32
	_ = v38
	if l1 < int32(2) {
		if l1 != 0 {
			return
		} else {
			v28 = m.G3
			v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			v33 = F_luaS_newlstr(m, l0, v28+int32(_a188), int32(0))
			mBase = m.M
			v34 = m.ExcPending
			if v34 != 0 {
				return
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v29)+8)) = int32(4)
				*(*int32)(unsafe.Add(mBase, uint32(v29))) = v33
				v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v38 + int32(16)
				return
			}
		}
	} else {
		v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
		v7 = *(*int32)(unsafe.Add(mBase, uint32(v6)+68))
		v8 = *(*int32)(unsafe.Add(mBase, uint32(v6)+64))
		if base.Ui32(v7) < base.Ui32(v8) {
			v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
			F_luaV_concat(m, l0, l1, (v12-v13)>>(uint(int32(4))%32)+int32(-1))
			mBase = m.M
			v20 = m.ExcPending
			if v20 != 0 {
				return
			} else {
				v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v21 + (int32(1)-l1)<<(uint(int32(4))%32)
				return
			}
		} else {
			F_luaC_step(m, l0)
			mBase = m.M
			v11 = m.ExcPending
			if v11 != 0 {
				return
			} else {
				v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
				F_luaV_concat(m, l0, l1, (v12-v13)>>(uint(int32(4))%32)+int32(-1))
				mBase = m.M
				v20 = m.ExcPending
				if v20 != 0 {
					return
				} else {
					v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
					*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v21 + (int32(1)-l1)<<(uint(int32(4))%32)
					return
				}
			}
		}
	}
}
func F_lua_createtable(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v17 int32
	_ = v17
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(v5)+68))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(v5)+64))
	if base.Ui32(v6) < base.Ui32(v7) {
		v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		v12 = F_luaH_new(m, l0, l1, l2)
		mBase = m.M
		v13 = m.ExcPending
		if v13 != 0 {
			return
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v11)+8)) = int32(5)
			*(*int32)(unsafe.Add(mBase, uint32(v11))) = v12
			v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v17 + int32(16)
			return
		}
	} else {
		F_luaC_step(m, l0)
		mBase = m.M
		v10 = m.ExcPending
		if v10 != 0 {
			return
		} else {
			v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			v12 = F_luaH_new(m, l0, l1, l2)
			mBase = m.M
			v13 = m.ExcPending
			if v13 != 0 {
				return
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v11)+8)) = int32(5)
				*(*int32)(unsafe.Add(mBase, uint32(v11))) = v12
				v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v17 + int32(16)
				return
			}
		}
	}
}
func F_lua_dump(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	v6 = int32(1)
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(v7+int32(-8))))
	if v10 != int32(6) {
		v23 = v6
		return v23
	} else {
		v15 = *(*int32)(unsafe.Add(mBase, uint32(v7+int32(-16))))
		v16 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15)+6)))
		if v16 != 0 {
			v23 = v6
			return v23
		} else {
			v17 = *(*int32)(unsafe.Add(mBase, uint32(v15)+16))
			v19 = F_luaU_dump(m, l0, v17, l1, l2, int32(0))
			mBase = m.M
			v22 = m.ExcPending
			if v22 != 0 {
				return int32(0)
			} else {
				v23 = v19
				return v23
			}
		}
	}
}
func F_lua_getfenv(m *base.Module, l0 int32, l1 int32) {
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
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
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
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v73 int64
	_ = v73
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	if l1 < int32(1) {
		if l1 < int32(-9999) {
			switch l1 + int32(10002) {
			case 0:
				v55 = l0 + int32(72)
			case 1:
				v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
				v29 = *(*int32)(unsafe.Add(mBase, uint32(v28)+4))
				v30 = *(*int32)(unsafe.Add(mBase, uint32(v29)))
				v31 = *(*int32)(unsafe.Add(mBase, uint32(v30)+12))
				*(*int32)(unsafe.Add(mBase, uint32(l0)+96)) = int32(5)
				*(*int32)(unsafe.Add(mBase, uint32(l0)+88)) = v31
				v55 = l0 + int32(88)
			case 2:
				v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
				v55 = v25 + int32(96)
			default:
				v39 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
				v40 = *(*int32)(unsafe.Add(mBase, uint32(v39)+4))
				v41 = *(*int32)(unsafe.Add(mBase, uint32(v40)))
				v42 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v41)+7)))
				v43 = m.G398
				if base.Ui32(v42) < base.Ui32(int32(-10002)-l1) {
					v54 = v43
				} else {
					v54 = v41 + (int32(-10003)-l1)<<(uint(int32(4))%32) + int32(24)
				}
				v55 = v54
			}
		} else {
			v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			v55 = v19 + l1<<(uint(int32(4))%32)
		}
	} else {
		v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
		v13 = v8 + l1<<(uint(int32(4))%32) + int32(-16)
		v14 = m.G398
		if base.Ui32(v13) < base.Ui32(v7) {
			v16 = v13
		} else {
			v16 = v14
		}
		v55 = v16
	}
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v55)+8))
	switch v58 + int32(-6) {
	case 0:
		v61 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		v62 = *(*int32)(unsafe.Add(mBase, uint32(v55)))
		v63 = *(*int32)(unsafe.Add(mBase, uint32(v62)+12))
		*(*int32)(unsafe.Add(mBase, uint32(v61))) = v63
		v78 = int32(5)
		v79 = v61
	case 1:
		v66 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		v67 = *(*int32)(unsafe.Add(mBase, uint32(v55)))
		v68 = *(*int32)(unsafe.Add(mBase, uint32(v67)+12))
		*(*int32)(unsafe.Add(mBase, uint32(v66))) = v68
		v78 = int32(5)
		v79 = v66
	case 2:
		v71 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		v72 = *(*int32)(unsafe.Add(mBase, uint32(v55)))
		v73 = *(*int64)(unsafe.Add(mBase, uint32(v72)+72))
		*(*int64)(unsafe.Add(mBase, uint32(v71))) = v73
		v75 = *(*int32)(unsafe.Add(mBase, uint32(v72)+80))
		v78 = v75
		v79 = v71
	default:
		v76 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		v78 = int32(0)
		v79 = v76
	}
	*(*int32)(unsafe.Add(mBase, uint32(v79)+8)) = v78
	v81 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v81 + int32(16)
	return
}
func F_lua_gethook(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
	return v2
}
func F_lua_getinfo(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v50 int32
	_ = v50
	var v57 int32
	_ = v57
	var v65 int32
	_ = v65
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v108 int32
	_ = v108
	var v113 int32
	_ = v113
	var v117 int32
	_ = v117
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v125 int32
	_ = v125
	var v127 int32
	_ = v127
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v138 int32
	_ = v138
	var v145 int32
	_ = v145
	var v147 int64
	_ = v147
	var v150 int32
	_ = v150
	var v152 int32
	_ = v152
	var v154 int32
	_ = v154
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v162 int32
	_ = v162
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v169 int32
	_ = v169
	var v175 int32
	_ = v175
	var v184 int32
	_ = v184
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
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
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v206 int32
	_ = v206
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v222 int32
	_ = v222
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
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
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v247 int32
	_ = v247
	var v248 int32
	_ = v248
	var v250 int32
	_ = v250
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
	var v263 int32
	_ = v263
	var v276 int32
	_ = v276
	var v283 int32
	_ = v283
	var v292 int32
	_ = v292
	var v298 int32
	_ = v298
	var v306 int32
	_ = v306
	var v311 int32
	_ = v311
	var v315 int32
	_ = v315
	var v318 int32
	_ = v318
	var v322 int32
	_ = v322
	var v327 int32
	_ = v327
	var v331 int32
	_ = v331
	var v334 int32
	_ = v334
	var v335 int32
	_ = v335
	var v336 int32
	_ = v336
	var v339 int32
	_ = v339
	var v341 int32
	_ = v341
	var v347 int32
	_ = v347
	var v348 int32
	_ = v348
	var v349 int32
	_ = v349
	var v352 int32
	_ = v352
	var v359 int32
	_ = v359
	var v361 int64
	_ = v361
	var v364 int32
	_ = v364
	var v366 int32
	_ = v366
	var v368 int32
	_ = v368
	var v371 int32
	_ = v371
	var v372 int32
	_ = v372
	var v376 int32
	_ = v376
	var v378 int32
	_ = v378
	var v379 int32
	_ = v379
	var v380 int32
	_ = v380
	var v383 int32
	_ = v383
	var v389 int32
	_ = v389
	var v396 int32
	_ = v396
	var v399 int32
	_ = v399
	var v400 int32
	_ = v400
	var v402 int32
	_ = v402
	var v406 int32
	_ = v406
	var v407 int32
	_ = v407
	var v409 int32
	_ = v409
	var v410 int32
	_ = v410
	var v413 int32
	_ = v413
	var v415 int32
	_ = v415
	var v416 int32
	_ = v416
	var v418 int32
	_ = v418
	var v422 int32
	_ = v422
	var v425 int32
	_ = v425
	var v429 int32
	_ = v429
	var v432 int32
	_ = v432
	var v433 int32
	_ = v433
	var v434 int32
	_ = v434
	var v435 int32
	_ = v435
	var v442 int32
	_ = v442
	var v447 int32
	_ = v447
	var v455 int32
	_ = v455
	var v456 int32
	_ = v456
	var v461 int32
	_ = v461
	var v463 int32
	_ = v463
	var v464 int32
	_ = v464
	var v471 int32
	_ = v471
	var v476 int32
	_ = v476
	var v477 int32
	_ = v477
	var v479 int32
	_ = v479
	var v483 int32
	_ = v483
	var v486 int32
	_ = v486
	var v487 int32
	_ = v487
	var v489 int32
	_ = v489
	var v493 int32
	_ = v493
	var v500 int32
	_ = v500
	var v507 int32
	_ = v507
	var v512 int32
	_ = v512
	var v526 int32
	_ = v526
	var v531 int32
	_ = v531
	var v534 int32
	_ = v534
	var v537 int32
	_ = v537
	var v538 int32
	_ = v538
	var v539 int32
	_ = v539
	var v540 int32
	_ = v540
	var v543 int32
	_ = v543
	var v547 int32
	_ = v547
	var v563 int32
	_ = v563
	var v564 int32
	_ = v564
	var v565 int32
	_ = v565
	var v566 int32
	_ = v566
	var v571 int32
	_ = v571
	var v572 int32
	_ = v572
	var v573 int32
	_ = v573
	var v591 int32
	_ = v591
	var v595 int32
	_ = v595
	var v599 int32
	_ = v599
	var v607 int32
	_ = v607
	var v612 int32
	_ = v612
	var v618 int32
	_ = v618
	var v619 int32
	_ = v619
	var v620 int32
	_ = v620
	var v635 int32
	_ = v635
	v17 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	if v17 != int32(62) {
		goto L12
	} else {
		goto L13
	}
L1:
	;
	return v635
L2:
	;
	v612 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if int32(16) < v612-v599 {
		v620 = v599
		goto L114
	} else {
		goto L115
	}
L3:
	;
	v534 = int32(0)
	v537 = F_luaH_new(m, l0, v534, v534)
	mBase = m.M
	v538 = m.ExcPending
	if v538 != 0 {
		goto L93
	} else {
		goto L107
	}
L4:
	;
	v531 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v531)+8)) = int32(0)
	v599 = v531
	v607 = v526
	goto L2
L5:
	;
	v512 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v500)+6)))
	if v512 == int32(0) {
		goto L3
	} else {
		goto L106
	}
L6:
	;
	v486 = int32(76)
	v487 = F___strchrnul(m, v37, v486)
	mBase = m.M
	v489 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v487))))
	if v489 == v486 {
		goto L102
	} else {
		goto L103
	}
L7:
	;
	v476 = int32(76)
	v477 = F___strchrnul(m, v461, v476)
	mBase = m.M
	v479 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v477))))
	if v479 == v476 {
		goto L96
	} else {
		goto L97
	}
L8:
	;
	v447 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if int32(16) < v447-v433 {
		goto L91
	} else {
		goto L92
	}
L9:
	;
	v415 = int32(102)
	v416 = F___strchrnul(m, v37, v415)
	mBase = m.M
	v418 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v416))))
	if v418 == v415 {
		goto L87
	} else {
		goto L88
	}
L10:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l2)+28)) = int64(-1)
	*(*int32)(unsafe.Add(mBase, uint32(l2)+20)) = int32(-1)
	v306 = m.G3
	*(*int32)(unsafe.Add(mBase, uint32(l2)+12)) = v306 + int32(_a2243)
	v311 = v306 + int32(_a188)
	*(*int32)(unsafe.Add(mBase, uint32(l2)+4)) = v311
	*(*int32)(unsafe.Add(mBase, uint32(l2)+8)) = v311
	v315 = v306 + int32(_a2244)
	*(*int32)(unsafe.Add(mBase, uint32(l2)+16)) = v315
	v318 = l2 + int32(36)
	v322 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v306)+uint32(_consts[1200]))))
	switch v322 + int32(-61) {
	case 0:
		goto L71
	default:
		goto L69
	case 3:
		goto L70
	}
L11:
	;
	if v39 == int32(0) {
		v298 = v37
		goto L10
	} else {
		goto L15
	}
L12:
	;
	v28 = *(*int32)(unsafe.Add(mBase, uint32(l2)+96))
	if v28 == int32(0) {
		v298 = l1
		goto L10
	} else {
		goto L14
	}
L13:
	;
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v22 = v20 + int32(-16)
	v23 = *(*int32)(unsafe.Add(mBase, uint32(v22)))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v22
	v37 = l1 + int32(1)
	v39 = v23
	v40 = int32(0)
	goto L11
L14:
	;
	v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v34 = v31 + v28*int32(24)
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v34)+4))
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v35)))
	v37 = l1
	v39 = v36
	v40 = v34
	goto L11
L15:
	;
	v44 = l2 + int32(36)
	v46 = v40 + int32(-12)
	v50 = v40 + int32(-24)
	v57 = v37
	v65 = int32(1)
	goto L16
L16:
	;
	v70 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v57))))
	switch v70 + int32(-76) {
	case 0, 26:
		v292 = v65
		goto L18
	case 1, 2, 3, 4, 5, 6, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 27, 28, 29, 30, 31, 33, 35, 36, 37, 38, 39, 40:
		goto L19
	case 7:
		goto L24
	case 32:
		goto L23
	case 34:
		goto L21
	case 41:
		goto L22
	default:
		goto L20
	}
L18:
	;
	v57 = v57 + int32(1)
	v65 = v292
	goto L16
L19:
	;
	v292 = int32(0)
	goto L18
L20:
	;
	if v70 == int32(0) {
		goto L9
	} else {
		goto L67
	}
L21:
	;
	if v40 == int32(0) {
		goto L53
	} else {
		goto L54
	}
L22:
	;
	v222 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v39)+7)))
	*(*int32)(unsafe.Add(mBase, uint32(l2)+24)) = v222
	v57 = v57 + int32(1)
	goto L16
L23:
	;
	v184 = int32(-1)
	if v40 == int32(0) {
		v215 = v184
		goto L43
	} else {
		goto L44
	}
L24:
	;
	v73 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v39)+6)))
	if v73 == int32(0) {
		goto L26
	} else {
		goto L27
	}
L25:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2)+12)) = v101
	*(*int32)(unsafe.Add(mBase, uint32(l2)+32)) = v100
	v108 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v102))))
	switch v108 + int32(-61) {
	case 0:
		goto L34
	default:
		goto L32
	case 3:
		goto L33
	}
L26:
	;
	v85 = *(*int32)(unsafe.Add(mBase, uint32(v39)+16))
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v85)+32))
	v88 = v86 + int32(16)
	*(*int32)(unsafe.Add(mBase, uint32(l2)+16)) = v88
	v90 = *(*int32)(unsafe.Add(mBase, uint32(v39)+16))
	v91 = *(*int32)(unsafe.Add(mBase, uint32(v90)+60))
	*(*int32)(unsafe.Add(mBase, uint32(l2)+28)) = v91
	v93 = m.G3
	if v91 != 0 {
		goto L28
	} else {
		goto L29
	}
L27:
	;
	v76 = int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(l2)+28)) = v76
	v79 = m.G3
	v81 = v79 + int32(_a2245)
	*(*int32)(unsafe.Add(mBase, uint32(l2)+16)) = v81
	v100 = v76
	v101 = v79 + int32(_a2246)
	v102 = v81
	goto L25
L28:
	;
	v96 = int32(_a2247)
	goto L30
L29:
	;
	v96 = int32(_a2248)
	goto L30
L30:
	;
	v98 = *(*int32)(unsafe.Add(mBase, uint32(v39)+16))
	v99 = *(*int32)(unsafe.Add(mBase, uint32(v98)+64))
	v100 = v99
	v101 = v93 + v96
	v102 = v88
	goto L25
L31:
	;
	v57 = v57 + int32(1)
	goto L16
L32:
	;
	v135 = m.G3
	v138 = F_strcspn(m, v102, v135+int32(_a2249))
	mBase = m.M
	v145 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v135)+uint32(_consts[1201]))))
	*(*uint16)(unsafe.Add(mBase, uint32(l2+int32(44)))) = uint16(v145)
	v147 = *(*int64)(unsafe.Add(mBase, uint32(v135)+uint32(_consts[1202])))
	*(*int64)(unsafe.Add(mBase, uint32(v44))) = v147
	v150 = int32(43)
	if base.Ui32(v138) < base.Ui32(v150) {
		goto L39
	} else {
		goto L40
	}
L33:
	;
	v120 = v102 + int32(1)
	v121 = F_strlen(m, v120)
	mBase = m.M
	v122 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v44))) = uint8(v122)
	v125 = int32(52)
	if base.Ui32(v121) <= base.Ui32(v125) {
		v133 = v120
		goto L35
	} else {
		goto L36
	}
L34:
	;
	v113 = F_strncpy(m, v44, v102+int32(1), int32(60))
	mBase = m.M
	v117 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v113+int32(59)))) = uint8(v117)
	goto L31
L35:
	;
	v134 = F_strcat(m, v44, v133)
	mBase = m.M
	goto L31
L36:
	;
	v127 = F_strlen(m, v44)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v44+v127))) = int32(3026478)
	v133 = v120 + (v121 - v125)
	goto L35
L37:
	;
	v164 = F_strlen(m, v44)
	mBase = m.M
	v165 = v44 + v164
	v166 = m.G3
	v169 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v166)+uint32(_consts[1203]))))
	*(*uint16)(unsafe.Add(mBase, uint32(v165))) = uint16(v169)
	v175 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v166)+uint32(_consts[1204]))))
	*(*uint8)(unsafe.Add(mBase, uint32(v165+int32(2)))) = uint8(v175)
	goto L31
L38:
	;
	v162 = F_strcat(m, v44, v102)
	mBase = m.M
	goto L37
L39:
	;
	v152 = v138
	goto L41
L40:
	;
	v152 = v150
	goto L41
L41:
	;
	v154 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v102+v152))))
	if v154 == int32(0) {
		goto L38
	} else {
		goto L42
	}
L42:
	;
	v157 = F_strncat(m, v44, v102, v152)
	mBase = m.M
	v158 = F_strlen(m, v157)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v157+v158))) = int32(3026478)
	goto L37
L43:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2)+20)) = v215
	v57 = v57 + int32(1)
	goto L16
L44:
	;
	v187 = *(*int32)(unsafe.Add(mBase, uint32(v40)+4))
	v188 = *(*int32)(unsafe.Add(mBase, uint32(v187)+8))
	if v188 != int32(6) {
		v215 = v184
		goto L43
	} else {
		goto L45
	}
L45:
	;
	v191 = *(*int32)(unsafe.Add(mBase, uint32(v187)))
	v192 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v191)+6)))
	if v192 != 0 {
		v215 = v184
		goto L43
	} else {
		goto L46
	}
L46:
	;
	v193 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v40 == v193 {
		goto L48
	} else {
		goto L49
	}
L47:
	;
	v201 = *(*int32)(unsafe.Add(mBase, uint32(v199)+16))
	v202 = *(*int32)(unsafe.Add(mBase, uint32(v201)+12))
	v203 = v200 - v202
	if v203 < int32(4) {
		v215 = v184
		goto L43
	} else {
		goto L50
	}
L48:
	;
	v196 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v40)+12)) = v196
	v198 = *(*int32)(unsafe.Add(mBase, uint32(v187)))
	v199 = v198
	v200 = v196
	goto L47
L49:
	;
	v195 = *(*int32)(unsafe.Add(mBase, uint32(v40)+12))
	v199 = v191
	v200 = v195
	goto L47
L50:
	;
	v206 = *(*int32)(unsafe.Add(mBase, uint32(v201)+20))
	if v206 != 0 {
		goto L51
	} else {
		goto L52
	}
L51:
	;
	v214 = *(*int32)(unsafe.Add(mBase, uint32(v206+v203+int32(-4))))
	v215 = v214
	goto L43
L52:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2)+20)) = int32(0)
	v57 = v57 + int32(1)
	goto L16
L53:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2)+4)) = int32(0)
	v283 = m.G3
	*(*int32)(unsafe.Add(mBase, uint32(l2)+8)) = v283 + int32(_a188)
	v57 = v57 + int32(1)
	goto L16
L54:
	;
	v228 = *(*int32)(unsafe.Add(mBase, uint32(v40)+4))
	v229 = *(*int32)(unsafe.Add(mBase, uint32(v228)+8))
	if v229 != int32(6) {
		goto L55
	} else {
		goto L56
	}
L55:
	;
	v237 = *(*int32)(unsafe.Add(mBase, uint32(v40+int32(-20))))
	v238 = *(*int32)(unsafe.Add(mBase, uint32(v237)+8))
	if v238 != int32(6) {
		goto L53
	} else {
		goto L59
	}
L56:
	;
	v232 = *(*int32)(unsafe.Add(mBase, uint32(v228)))
	v233 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v232)+6)))
	if v233 != 0 {
		goto L55
	} else {
		goto L57
	}
L57:
	;
	v234 = *(*int32)(unsafe.Add(mBase, uint32(v40)+20))
	if int32(0) < v234 {
		goto L53
	} else {
		goto L58
	}
L58:
	;
	goto L55
L59:
	;
	v241 = *(*int32)(unsafe.Add(mBase, uint32(v237)))
	v242 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v241)+6)))
	if v242 != 0 {
		goto L53
	} else {
		goto L60
	}
L60:
	;
	v243 = *(*int32)(unsafe.Add(mBase, uint32(v241)+16))
	v244 = *(*int32)(unsafe.Add(mBase, uint32(v243)+12))
	v245 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v50 == v245 {
		goto L62
	} else {
		goto L63
	}
L61:
	;
	v259 = *(*int32)(unsafe.Add(mBase, uint32(v244+(v254-v253)+int32(-4))))
	v263 = v259&int32(63) + int32(-28)
	if base.Ui32(int32(5)) < base.Ui32(v263) {
		goto L53
	} else {
		goto L64
	}
L62:
	;
	v248 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v46))) = v248
	v250 = *(*int32)(unsafe.Add(mBase, uint32(v237)))
	v251 = *(*int32)(unsafe.Add(mBase, uint32(v250)+16))
	v252 = *(*int32)(unsafe.Add(mBase, uint32(v251)+12))
	v253 = v252
	v254 = v248
	goto L61
L63:
	;
	v247 = *(*int32)(unsafe.Add(mBase, uint32(v46)))
	v253 = v244
	v254 = v247
	goto L61
L64:
	;
	if int32(1)<<(uint(v263)%32)&int32(35) == int32(0) {
		goto L53
	} else {
		goto L65
	}
L65:
	;
	v276 = F_getobjname(m, l0, v50, int32(base.Ui32(v259)>>(uint(int32(6))%32))&int32(255), l2+int32(4))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(l2)+8)) = v276
	if v276 != 0 {
		v292 = v65
		goto L18
	} else {
		goto L66
	}
L66:
	;
	goto L53
L67:
	;
	goto L19
L68:
	;
	v396 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l2)+24)) = v396
	v399 = int32(102)
	v400 = F___strchrnul(m, v298, v399)
	mBase = m.M
	v402 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v400))))
	if v402 == v399 {
		goto L82
	} else {
		goto L83
	}
L69:
	;
	v349 = m.G3
	v352 = F_strcspn(m, v315, v349+int32(_a2249))
	mBase = m.M
	v359 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v349)+uint32(_consts[1201]))))
	*(*uint16)(unsafe.Add(mBase, uint32(l2+int32(44)))) = uint16(v359)
	v361 = *(*int64)(unsafe.Add(mBase, uint32(v349)+uint32(_consts[1202])))
	*(*int64)(unsafe.Add(mBase, uint32(v318))) = v361
	v364 = int32(43)
	if base.Ui32(v352) < base.Ui32(v364) {
		goto L76
	} else {
		goto L77
	}
L70:
	;
	v334 = v306 + int32(_a2250)
	v335 = F_strlen(m, v334)
	mBase = m.M
	v336 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v318))) = uint8(v336)
	v339 = int32(52)
	if base.Ui32(v335) <= base.Ui32(v339) {
		v347 = v334
		goto L72
	} else {
		goto L73
	}
L71:
	;
	v327 = F_strncpy(m, v318, v306+int32(_a2250), int32(60))
	mBase = m.M
	v331 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v327+int32(59)))) = uint8(v331)
	goto L68
L72:
	;
	v348 = F_strcat(m, v318, v347)
	mBase = m.M
	goto L68
L73:
	;
	v341 = F_strlen(m, v318)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v318+v341))) = int32(3026478)
	v347 = v334 + (v335 - v339)
	goto L72
L74:
	;
	v378 = F_strlen(m, v318)
	mBase = m.M
	v379 = v318 + v378
	v380 = m.G3
	v383 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v380)+uint32(_consts[1203]))))
	*(*uint16)(unsafe.Add(mBase, uint32(v379))) = uint16(v383)
	v389 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v380)+uint32(_consts[1204]))))
	*(*uint8)(unsafe.Add(mBase, uint32(v379+int32(2)))) = uint8(v389)
	goto L68
L75:
	;
	v376 = F_strcat(m, v318, v315)
	mBase = m.M
	goto L74
L76:
	;
	v366 = v352
	goto L78
L77:
	;
	v366 = v364
	goto L78
L78:
	;
	v368 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v315+v366))))
	if v368 == int32(0) {
		goto L75
	} else {
		goto L79
	}
L79:
	;
	v371 = F_strncat(m, v318, v315, v366)
	mBase = m.M
	v372 = F_strlen(m, v371)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v371+v372))) = int32(3026478)
	goto L74
L80:
	;
	v409 = int32(0)
	v410 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v410)+8)) = v409
	v413 = int32(1)
	v432 = v298
	v433 = v410
	v434 = v413
	v435 = v409
	v442 = v413
	goto L8
L81:
	;
	if v406 != 0 {
		goto L80
	} else {
		goto L85
	}
L82:
	;
	v406 = v400
	goto L84
L83:
	;
	v406 = v396
	goto L84
L84:
	;
	goto L81
L85:
	;
	v407 = int32(1)
	v461 = v298
	v463 = v407
	v464 = v396
	v471 = v407
	goto L7
L86:
	;
	if v422 == int32(0) {
		goto L6
	} else {
		goto L90
	}
L87:
	;
	v422 = v416
	goto L89
L88:
	;
	v422 = int32(0)
	goto L89
L89:
	;
	goto L86
L90:
	;
	v425 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v425)+8)) = int32(6)
	*(*int32)(unsafe.Add(mBase, uint32(v425))) = v39
	v429 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v432 = v37
	v433 = v429
	v434 = int32(0)
	v435 = v39
	v442 = v65
	goto L8
L91:
	;
	v456 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v456 + int32(16)
	v461 = v432
	v463 = v434
	v464 = v435
	v471 = v442
	goto L7
L92:
	;
	F_luaD_growstack(m, l0, int32(1))
	mBase = m.M
	v455 = m.ExcPending
	if v455 != 0 {
		goto L93
	} else {
		goto L94
	}
L93:
	;
	return int32(0)
L94:
	;
	goto L91
L95:
	;
	if v483 == int32(0) {
		v635 = v471
		goto L1
	} else {
		goto L99
	}
L96:
	;
	v483 = v477
	goto L98
L97:
	;
	v483 = int32(0)
	goto L98
L98:
	;
	goto L95
L99:
	;
	if v463 != 0 {
		v526 = v471
		goto L4
	} else {
		goto L100
	}
L100:
	;
	v500 = v464
	v507 = v471
	goto L5
L101:
	;
	if v493 == int32(0) {
		v635 = v65
		goto L1
	} else {
		goto L105
	}
L102:
	;
	v493 = v487
	goto L104
L103:
	;
	v493 = int32(0)
	goto L104
L104:
	;
	goto L101
L105:
	;
	v500 = v39
	v507 = v65
	goto L5
L106:
	;
	v526 = v507
	goto L4
L107:
	;
	v539 = *(*int32)(unsafe.Add(mBase, uint32(v500)+16))
	v540 = *(*int32)(unsafe.Add(mBase, uint32(v539)+48))
	if v540 < int32(1) {
		goto L108
	} else {
		goto L109
	}
L108:
	;
	v591 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v591)+8)) = int32(5)
	*(*int32)(unsafe.Add(mBase, uint32(v591))) = v537
	v595 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v599 = v595
	v607 = v507
	goto L2
L109:
	;
	v543 = *(*int32)(unsafe.Add(mBase, uint32(v539)+20))
	v547 = v534
	goto L110
L110:
	;
	v563 = *(*int32)(unsafe.Add(mBase, uint32(v543+v547<<(uint(int32(2))%32))))
	v564 = F_luaH_setnum(m, l0, v537, v563)
	mBase = m.M
	v565 = m.ExcPending
	if v565 != 0 {
		goto L93
	} else {
		goto L112
	}
L111:
	;
	goto L108
L112:
	;
	v566 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v564)+8)) = v566
	*(*int32)(unsafe.Add(mBase, uint32(v564))) = v566
	v571 = v547 + v566
	v572 = *(*int32)(unsafe.Add(mBase, uint32(v500)+16))
	v573 = *(*int32)(unsafe.Add(mBase, uint32(v572)+48))
	if v571 < v573 {
		v547 = v571
		goto L110
	} else {
		goto L113
	}
L113:
	;
	goto L111
L114:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v620 + int32(16)
	v635 = v607
	goto L1
L115:
	;
	F_luaD_growstack(m, l0, int32(1))
	mBase = m.M
	v618 = m.ExcPending
	if v618 != 0 {
		goto L93
	} else {
		goto L116
	}
L116:
	;
	v619 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v620 = v619
	goto L114
}
func F_lua_getlocal(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
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
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
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
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
	var v113 int64
	_ = v113
	var v115 int32
	_ = v115
	var v117 int32
	_ = v117
	var v122 int32
	_ = v122
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l1)+96))
	v11 = v7 + v8*int32(24)
	v12 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
	v13 = *(*int32)(unsafe.Add(mBase, uint32(v12)+8))
	if v13 != int32(6) {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	return v122
L2:
	;
	v110 = v102 + l2<<(uint(int32(4))%32) + int32(-16)
	v112 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v113 = *(*int64)(unsafe.Add(mBase, uint32(v110)))
	*(*int64)(unsafe.Add(mBase, uint32(v112))) = v113
	v115 = *(*int32)(unsafe.Add(mBase, uint32(v110)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v112)+8)) = v115
	v117 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v117 + int32(16)
	goto L27
L3:
	;
	v87 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v11 == v87 {
		goto L22
	} else {
		goto L23
	}
L4:
	;
	v16 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
	v17 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16)+6)))
	if v17 != 0 {
		goto L3
	} else {
		goto L5
	}
L5:
	;
	v18 = *(*int32)(unsafe.Add(mBase, uint32(v16)+16))
	if v18 == int32(0) {
		goto L3
	} else {
		goto L6
	}
L6:
	;
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v11 == v21 {
		goto L8
	} else {
		goto L9
	}
L7:
	;
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v28)+12))
	v35 = (v29-v30)>>(uint(int32(2))%32) + int32(-1)
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v18)+56))
	if int32(1) <= v39 {
		goto L11
	} else {
		goto L12
	}
L8:
	;
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+12)) = v24
	v26 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v26)+16))
	v28 = v27
	v29 = v24
	goto L7
L9:
	;
	v23 = *(*int32)(unsafe.Add(mBase, uint32(v11)+12))
	v28 = v18
	v29 = v23
	goto L7
L10:
	;
	if v76 == int32(0) {
		goto L3
	} else {
		goto L21
	}
L11:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v18)+24))
	v45 = int32(0)
	v46 = l2
	goto L14
L12:
	;
	v76 = int32(0)
	goto L10
L13:
	;
	v76 = int32(0)
	goto L10
L14:
	;
	v53 = v43 + v45*int32(12)
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v53)+4))
	if v35 < v54 {
		goto L13
	} else {
		goto L16
	}
L15:
	;
	goto L13
L16:
	;
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v53)+8))
	if v56 <= v35 {
		v63 = v46
		goto L17
	} else {
		goto L18
	}
L17:
	;
	v65 = v45 + int32(1)
	if v65 != v39 {
		v45 = v65
		v46 = v63
		goto L14
	} else {
		goto L20
	}
L18:
	;
	v59 = v46 + int32(-1)
	if v59 != 0 {
		v63 = v59
		goto L17
	} else {
		goto L19
	}
L19:
	;
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v53)))
	v76 = v60 + int32(16)
	goto L10
L20:
	;
	goto L15
L21:
	;
	v79 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
	v102 = v79
	v103 = v76
	goto L2
L22:
	;
	v89 = l0 + int32(8)
	goto L24
L23:
	;
	v89 = v11 + int32(28)
	goto L24
L24:
	;
	v90 = *(*int32)(unsafe.Add(mBase, uint32(v89)))
	v91 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
	v92 = m.G3
	v93 = int32(0)
	if l2 < int32(1) {
		v122 = v93
		goto L1
	} else {
		goto L25
	}
L25:
	;
	if (v90-v91)>>(uint(int32(4))%32) < l2 {
		v122 = v93
		goto L1
	} else {
		goto L26
	}
L26:
	;
	v102 = v91
	v103 = v92 + int32(_a2242)
	goto L2
L27:
	;
	v122 = v103
	goto L1
}
func F_lua_getmetatable(m *base.Module, l0 int32, l1 int32) int32 {
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
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v81 int32
	_ = v81
	if l1 < int32(1) {
		if l1 < int32(-9999) {
			switch l1 + int32(10002) {
			case 0:
				v55 = l0 + int32(72)
			case 1:
				v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
				v29 = *(*int32)(unsafe.Add(mBase, uint32(v28)+4))
				v30 = *(*int32)(unsafe.Add(mBase, uint32(v29)))
				v31 = *(*int32)(unsafe.Add(mBase, uint32(v30)+12))
				*(*int32)(unsafe.Add(mBase, uint32(l0)+96)) = int32(5)
				*(*int32)(unsafe.Add(mBase, uint32(l0)+88)) = v31
				v55 = l0 + int32(88)
			case 2:
				v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
				v55 = v25 + int32(96)
			default:
				v39 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
				v40 = *(*int32)(unsafe.Add(mBase, uint32(v39)+4))
				v41 = *(*int32)(unsafe.Add(mBase, uint32(v40)))
				v42 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v41)+7)))
				v43 = m.G398
				if base.Ui32(v42) < base.Ui32(int32(-10002)-l1) {
					v54 = v43
				} else {
					v54 = v41 + (int32(-10003)-l1)<<(uint(int32(4))%32) + int32(24)
				}
				v55 = v54
			}
		} else {
			v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			v55 = v19 + l1<<(uint(int32(4))%32)
		}
	} else {
		v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
		v13 = v8 + l1<<(uint(int32(4))%32) + int32(-16)
		v14 = m.G398
		if base.Ui32(v13) < base.Ui32(v7) {
			v16 = v13
		} else {
			v16 = v14
		}
		v55 = v16
	}
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v55)+8))
	switch v58 + int32(-5) {
	case 0:
		v61 = *(*int32)(unsafe.Add(mBase, uint32(v55)))
		v73 = v61 + int32(16)
	default:
		v67 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
		v73 = v67 + v58<<(uint(int32(2))%32) + int32(152)
	case 2:
		v64 = *(*int32)(unsafe.Add(mBase, uint32(v55)))
		v73 = v64 + int32(8)
	}
	v74 = *(*int32)(unsafe.Add(mBase, uint32(v73)))
	if v74 != 0 {
		v77 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		*(*int32)(unsafe.Add(mBase, uint32(v77)+8)) = int32(5)
		*(*int32)(unsafe.Add(mBase, uint32(v77))) = v74
		v81 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v81 + int32(16)
		return int32(1)
	} else {
		return int32(0)
	}
}
func F_lua_isnumber(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	if l1 < int32(1) {
		if l1 < int32(-9999) {
			switch l1 + int32(10002) {
			case 0:
				v60 = l0 + int32(72)
			case 1:
				v32 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
				v33 = *(*int32)(unsafe.Add(mBase, uint32(v32)+4))
				v34 = *(*int32)(unsafe.Add(mBase, uint32(v33)))
				v35 = *(*int32)(unsafe.Add(mBase, uint32(v34)+12))
				*(*int32)(unsafe.Add(mBase, uint32(l0)+96)) = int32(5)
				*(*int32)(unsafe.Add(mBase, uint32(l0)+88)) = v35
				v60 = l0 + int32(88)
			case 2:
				v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
				v60 = v29 + int32(96)
			default:
				v43 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
				v44 = *(*int32)(unsafe.Add(mBase, uint32(v43)+4))
				v45 = *(*int32)(unsafe.Add(mBase, uint32(v44)))
				v46 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v45)+7)))
				v47 = m.G398
				if base.Ui32(v46) < base.Ui32(int32(-10002)-l1) {
					v58 = v47
				} else {
					v58 = v45 + (int32(-10003)-l1)<<(uint(int32(4))%32) + int32(24)
				}
				v60 = v58
			}
		} else {
			v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			v60 = v23 + l1<<(uint(int32(4))%32)
		}
	} else {
		v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
		v17 = v12 + l1<<(uint(int32(4))%32) + int32(-16)
		v18 = m.G398
		if base.Ui32(v17) < base.Ui32(v11) {
			v20 = v17
		} else {
			v20 = v18
		}
		v60 = v20
	}
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v60)+8))
	if v63 == int32(3) {
		v72 = int32(1)
		m.G0 = v7 + int32(16)
		return v72
	} else {
		v66 = F_luaV_tonumber(m, v60, v7)
		mBase = m.M
		v69 = m.ExcPending
		if v69 != 0 {
			return int32(0)
		} else {
			v72 = base.B2i32(v66 != int32(0))
			m.G0 = v7 + int32(16)
			return v72
		}
	}
}
func F_lua_isreadonlytable(m *base.Module, l0 int32, l1 int32) int32 {
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
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	if l1 < int32(1) {
		if l1 < int32(-9999) {
			switch l1 + int32(10002) {
			case 0:
				v55 = l0 + int32(72)
			case 1:
				v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
				v28 = *(*int32)(unsafe.Add(mBase, uint32(v27)+4))
				v29 = *(*int32)(unsafe.Add(mBase, uint32(v28)))
				v30 = *(*int32)(unsafe.Add(mBase, uint32(v29)+12))
				*(*int32)(unsafe.Add(mBase, uint32(l0)+96)) = int32(5)
				*(*int32)(unsafe.Add(mBase, uint32(l0)+88)) = v30
				v55 = l0 + int32(88)
			case 2:
				v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
				v55 = v24 + int32(96)
			default:
				v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
				v39 = *(*int32)(unsafe.Add(mBase, uint32(v38)+4))
				v40 = *(*int32)(unsafe.Add(mBase, uint32(v39)))
				v41 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v40)+7)))
				v42 = m.G398
				if base.Ui32(v41) < base.Ui32(int32(-10002)-l1) {
					v53 = v42
				} else {
					v53 = v40 + (int32(-10003)-l1)<<(uint(int32(4))%32) + int32(24)
				}
				v55 = v53
			}
		} else {
			v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			v55 = v18 + l1<<(uint(int32(4))%32)
		}
	} else {
		v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
		v12 = v7 + l1<<(uint(int32(4))%32) + int32(-16)
		v13 = m.G398
		if base.Ui32(v12) < base.Ui32(v6) {
			v15 = v12
		} else {
			v15 = v13
		}
		v55 = v15
	}
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v55)))
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v57)+8))
	return v58
}
func F_lua_isstring(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v17 int32
	_ = v17
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v50 int32
	_ = v50
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v59 int32
	_ = v59
	if l1 < int32(1) {
		if l1 < int32(-9999) {
			switch l1 + int32(10002) {
			case 0:
				v54 = l0 + int32(72)
				v55 = m.G398
				if v54 != v55 {
					v59 = *(*int32)(unsafe.Add(mBase, uint32(v54)+8))
					return base.B2i32(base.Ui32(v59+int32(-3)) < base.Ui32(int32(2)))
				} else {
					return int32(0)
				}
			case 1:
				v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
				v24 = *(*int32)(unsafe.Add(mBase, uint32(v23)+4))
				v25 = *(*int32)(unsafe.Add(mBase, uint32(v24)))
				v26 = *(*int32)(unsafe.Add(mBase, uint32(v25)+12))
				*(*int32)(unsafe.Add(mBase, uint32(l0)+96)) = int32(5)
				*(*int32)(unsafe.Add(mBase, uint32(l0)+88)) = v26
				v54 = l0 + int32(88)
				v55 = m.G398
				if v54 != v55 {
					v59 = *(*int32)(unsafe.Add(mBase, uint32(v54)+8))
					return base.B2i32(base.Ui32(v59+int32(-3)) < base.Ui32(int32(2)))
				} else {
					return int32(0)
				}
			case 2:
				v50 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
				v54 = v50 + int32(96)
				v55 = m.G398
				if v54 != v55 {
					v59 = *(*int32)(unsafe.Add(mBase, uint32(v54)+8))
					return base.B2i32(base.Ui32(v59+int32(-3)) < base.Ui32(int32(2)))
				} else {
					return int32(0)
				}
			default:
				v36 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
				v37 = *(*int32)(unsafe.Add(mBase, uint32(v36)+4))
				v38 = *(*int32)(unsafe.Add(mBase, uint32(v37)))
				v39 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v38)+7)))
				if base.Ui32(v39) < base.Ui32(int32(-10002)-l1) {
					return int32(0)
				} else {
					v54 = v38 + (int32(-10003)-l1)<<(uint(int32(4))%32) + int32(24)
					v55 = m.G398
					if v54 != v55 {
						v59 = *(*int32)(unsafe.Add(mBase, uint32(v54)+8))
						return base.B2i32(base.Ui32(v59+int32(-3)) < base.Ui32(int32(2)))
					} else {
						return int32(0)
					}
				}
			}
		} else {
			v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			v54 = v17 + l1<<(uint(int32(4))%32)
			v55 = m.G398
			if v54 != v55 {
				v59 = *(*int32)(unsafe.Add(mBase, uint32(v54)+8))
				return base.B2i32(base.Ui32(v59+int32(-3)) < base.Ui32(int32(2)))
			} else {
				return int32(0)
			}
		}
	} else {
		v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
		v10 = v5 + l1<<(uint(int32(4))%32) + int32(-16)
		v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		if base.Ui32(v10) < base.Ui32(v11) {
			v54 = v10
			v55 = m.G398
			if v54 != v55 {
				v59 = *(*int32)(unsafe.Add(mBase, uint32(v54)+8))
				return base.B2i32(base.Ui32(v59+int32(-3)) < base.Ui32(int32(2)))
			} else {
				return int32(0)
			}
		} else {
			return int32(0)
		}
	}
}
func F_lua_load(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	v6 = m.G0
	v8 = v6 - int32(32)
	m.G0 = v8
	v11 = v8 + int32(12)
	*(*int32)(unsafe.Add(mBase, uint32(v11)+12)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(v11)+8)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = l0
	*(*int64)(unsafe.Add(mBase, uint32(v11))) = int64(0)
	v19 = m.G3
	if l3 != 0 {
		v22 = l3
	} else {
		v22 = v19 + int32(_a176)
	}
	v23 = F_luaD_protectedparser(m, l0, v8+int32(12), v22)
	mBase = m.M
	v26 = m.ExcPending
	if v26 != 0 {
		return int32(0)
	} else {
		m.G0 = v8 + int32(32)
		return v23
	}
}
func F_lua_objlen(m *base.Module, l0 int32, l1 int32) int32 {
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
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v105 int32
	_ = v105
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v126 int32
	_ = v126
	var __phi126 int32
	_ = __phi126
	var v128 int32
	_ = v128
	var __phi128 int32
	_ = __phi128
	var v137 int32
	_ = v137
	var v143 float64
	_ = v143
	var v147 float64
	_ = v147
	var v148 int64
	_ = v148
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v160 int32
	_ = v160
	var v164 int32
	_ = v164
	var v165 float64
	_ = v165
	var v169 int32
	_ = v169
	var v177 int32
	_ = v177
	var v180 float64
	_ = v180
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v189 int32
	_ = v189
	var v194 int32
	_ = v194
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v208 int32
	_ = v208
	var v210 int32
	_ = v210
	var v214 int32
	_ = v214
	var v221 float64
	_ = v221
	var v222 int64
	_ = v222
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v234 int32
	_ = v234
	var v238 int32
	_ = v238
	var v239 float64
	_ = v239
	var v243 int32
	_ = v243
	var v251 int32
	_ = v251
	var v254 float64
	_ = v254
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v263 int32
	_ = v263
	var v268 int32
	_ = v268
	var v269 int32
	_ = v269
	var v270 int32
	_ = v270
	var v274 int32
	_ = v274
	var v276 int32
	_ = v276
	var v281 int32
	_ = v281
	var v293 int32
	_ = v293
	var v299 float64
	_ = v299
	var v303 float64
	_ = v303
	var v304 int64
	_ = v304
	var v309 int32
	_ = v309
	var v310 int32
	_ = v310
	var v316 int32
	_ = v316
	var v320 int32
	_ = v320
	var v321 float64
	_ = v321
	var v325 int32
	_ = v325
	var v333 int32
	_ = v333
	var v336 float64
	_ = v336
	var v338 int32
	_ = v338
	var v339 int32
	_ = v339
	var v345 int32
	_ = v345
	var v352 int32
	_ = v352
	var v358 int32
	_ = v358
	var v375 int32
	_ = v375
	var v377 int32
	_ = v377
	var v380 int32
	_ = v380
	var v383 int32
	_ = v383
	var v384 int32
	_ = v384
	var v385 int32
	_ = v385
	if l1 < int32(1) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	v58 = int32(0)
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v55)+8))
	switch v59 + int32(-3) {
	case 0:
		goto L17
	case 1:
		goto L20
	case 2:
		goto L18
	default:
		v385 = v58
		goto L16
	case 4:
		goto L19
	}
L2:
	;
	if l1 < int32(-9999) {
		goto L7
	} else {
		goto L8
	}
L3:
	;
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v13 = v8 + l1<<(uint(int32(4))%32) + int32(-16)
	v14 = m.G398
	if base.Ui32(v13) < base.Ui32(v7) {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v16 = v13
	goto L6
L5:
	;
	v16 = v14
	goto L6
L6:
	;
	v55 = v16
	goto L1
L7:
	;
	switch l1 + int32(10002) {
	case 0:
		goto L10
	case 1:
		goto L11
	case 2:
		goto L12
	default:
		goto L9
	}
L8:
	;
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v55 = v19 + l1<<(uint(int32(4))%32)
	goto L1
L9:
	;
	v39 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v39)+4))
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v40)))
	v42 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v41)+7)))
	v43 = m.G398
	if base.Ui32(v42) < base.Ui32(int32(-10002)-l1) {
		goto L13
	} else {
		goto L14
	}
L10:
	;
	v55 = l0 + int32(72)
	goto L1
L11:
	;
	v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v28)+4))
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v29)))
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v30)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+96)) = int32(5)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+88)) = v31
	v55 = l0 + int32(88)
	goto L1
L12:
	;
	v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v55 = v25 + int32(96)
	goto L1
L13:
	;
	v54 = v43
	goto L15
L14:
	;
	v54 = v41 + (int32(-10003)-l1)<<(uint(int32(4))%32) + int32(24)
	goto L15
L15:
	;
	v55 = v54
	goto L1
L16:
	;
	return v385
L17:
	;
	v377 = F_luaV_tostring(m, l0, v55)
	mBase = m.M
	v380 = m.ExcPending
	if v380 != 0 {
		goto L96
	} else {
		goto L97
	}
L18:
	;
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v55)))
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v68)+36))
	if v78 == int32(0) {
		goto L23
	} else {
		goto L24
	}
L19:
	;
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v55)))
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v65)+16))
	return v66
L20:
	;
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v55)))
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v62)+12))
	return v63
L21:
	;
	return v375
L22:
	;
	v375 = v358
	goto L21
L23:
	;
	v116 = m.G3
	v117 = *(*int32)(unsafe.Add(mBase, uint32(v68)+24))
	if v117 != v116+int32(_a2240) {
		goto L36
	} else {
		goto L37
	}
L24:
	;
	v81 = *(*int32)(unsafe.Add(mBase, uint32(v68)+20))
	v87 = *(*int32)(unsafe.Add(mBase, uint32(v81+v78<<(uint(int32(4))%32)+int32(-8))))
	if v87 != 0 {
		goto L23
	} else {
		goto L25
	}
L25:
	;
	v88 = int32(0)
	if v78 == int32(1) {
		v358 = v88
		goto L22
	} else {
		goto L26
	}
L26:
	;
	v94 = v78
	v96 = v88
	goto L27
L27:
	;
	v105 = int32(base.Ui32(v96+v94) >> (uint(int32(1)) % 32))
	v109 = *(*int32)(unsafe.Add(mBase, uint32(v81+int32(-8)+v105<<(uint(int32(4))%32))))
	if v109 != 0 {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	v110 = v94
	goto L31
L30:
	;
	v110 = v105
	goto L31
L31:
	;
	if v109 != 0 {
		goto L32
	} else {
		goto L33
	}
L32:
	;
	v111 = v105
	goto L34
L33:
	;
	v111 = v96
	goto L34
L34:
	;
	if base.Ui32(int32(1)) < base.Ui32(v110-v111) {
		v94 = v110
		v96 = v111
		goto L27
	} else {
		goto L35
	}
L35:
	;
	v358 = v111
	goto L22
L36:
	;
	__phi126 = v78
	__phi128 = v78 + int32(1)
	v126 = __phi126
	v128 = __phi128
	goto L38
L37:
	;
	v375 = v78
	goto L21
L38:
	;
	if v128 < int32(1) {
		goto L43
	} else {
		goto L44
	}
L39:
	;
	v281 = v274
	goto L79
L40:
	;
	v194 = *(*int32)(unsafe.Add(mBase, uint32(v189)+8))
	if v194 != 0 {
		goto L54
	} else {
		goto L55
	}
L41:
	;
	v169 = v164
	goto L48
L42:
	;
	v148 = base.I64_reinterpret_f64(v147)
	v153 = int32(-1)
	v154 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v68)+12)))
	v160 = base.I32_rem_u_s(base.I32_wrap_i64(int64(base.Ui64(v148)>>(uint(int64(32))%64))+v148), v153<<(uint(v154)%32)^v153|int32(1))
	v164 = v117 + v160<<(uint(int32(5))%32)
	v165 = v147
	goto L41
L43:
	;
	v143 = base.F64_convert_i32_s(v128)
	if v128 == int32(0) {
		v164 = v117
		v165 = v143
		goto L41
	} else {
		goto L47
	}
L44:
	;
	if v128 <= v78 {
		goto L45
	} else {
		goto L46
	}
L45:
	;
	v137 = *(*int32)(unsafe.Add(mBase, uint32(v68)+20))
	v189 = v137 + v128<<(uint(int32(4))%32) + int32(-16)
	goto L40
L46:
	;
	v147 = base.F64_convert_i32_u(v128)
	goto L42
L47:
	;
	v147 = v143
	goto L42
L48:
	;
	v177 = *(*int32)(unsafe.Add(mBase, uint32(v169)+24))
	if v177 != int32(3) {
		goto L50
	} else {
		goto L51
	}
L49:
	;
	v189 = v183
	goto L40
L50:
	;
	v182 = *(*int32)(unsafe.Add(mBase, uint32(v169)+28))
	v183 = m.G398
	if v182 != 0 {
		v169 = v182
		goto L48
	} else {
		goto L53
	}
L51:
	;
	v180 = *(*float64)(unsafe.Add(mBase, uint32(v169)+16))
	if base.F64_ne(v180, v165) != 0 {
		goto L50
	} else {
		goto L52
	}
L52:
	;
	v189 = v169
	goto L40
L53:
	;
	goto L49
L54:
	;
	v274 = int32(1)
	v276 = v128 << (uint(v274) % 32)
	if base.Ui32(v276) < base.Ui32(int32(2147483646)) {
		__phi126 = v128
		__phi128 = v276
		v126 = __phi126
		v128 = __phi128
		goto L38
	} else {
		goto L78
	}
L55:
	;
	if base.Ui32(v128-v126) <= base.Ui32(int32(1)) {
		v358 = v126
		goto L22
	} else {
		goto L56
	}
L56:
	;
	v201 = v126
	v202 = v128
	goto L57
L57:
	;
	v208 = v202 + v201
	v210 = int32(base.Ui32(v208) >> (uint(int32(1)) % 32))
	if base.Ui32(v208) < base.Ui32(int32(2)) {
		goto L62
	} else {
		goto L63
	}
L59:
	;
	v268 = *(*int32)(unsafe.Add(mBase, uint32(v263)+8))
	if v268 != 0 {
		goto L71
	} else {
		goto L72
	}
L60:
	;
	v243 = v238
	goto L65
L61:
	;
	v221 = base.F64_convert_i32_u(v210)
	v222 = base.I64_reinterpret_f64(v221)
	v227 = int32(-1)
	v228 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v68)+12)))
	v234 = base.I32_rem_u_s(base.I32_wrap_i64(int64(base.Ui64(v222)>>(uint(int64(32))%64))+v222), v227<<(uint(v228)%32)^v227|int32(1))
	v238 = v117 + v234<<(uint(int32(5))%32)
	v239 = v221
	goto L60
L62:
	;
	v238 = v117
	v239 = base.F64_convert_i32_u(v210)
	goto L60
L63:
	;
	if v78 < v210 {
		goto L61
	} else {
		goto L64
	}
L64:
	;
	v214 = *(*int32)(unsafe.Add(mBase, uint32(v68)+20))
	v263 = v214 + v210<<(uint(int32(4))%32) + int32(-16)
	goto L59
L65:
	;
	v251 = *(*int32)(unsafe.Add(mBase, uint32(v243)+24))
	if v251 != int32(3) {
		goto L67
	} else {
		goto L68
	}
L66:
	;
	v263 = v257
	goto L59
L67:
	;
	v256 = *(*int32)(unsafe.Add(mBase, uint32(v243)+28))
	v257 = m.G398
	if v256 != 0 {
		v243 = v256
		goto L65
	} else {
		goto L70
	}
L68:
	;
	v254 = *(*float64)(unsafe.Add(mBase, uint32(v243)+16))
	if base.F64_ne(v254, v239) != 0 {
		goto L67
	} else {
		goto L69
	}
L69:
	;
	v263 = v243
	goto L59
L70:
	;
	goto L66
L71:
	;
	v269 = v202
	goto L73
L72:
	;
	v269 = v210
	goto L73
L73:
	;
	if v268 != 0 {
		goto L74
	} else {
		goto L75
	}
L74:
	;
	v270 = v210
	goto L76
L75:
	;
	v270 = v201
	goto L76
L76:
	;
	if base.Ui32(int32(1)) < base.Ui32(v269-v270) {
		v201 = v270
		v202 = v269
		goto L57
	} else {
		goto L77
	}
L77:
	;
	v358 = v270
	goto L22
L78:
	;
	goto L39
L79:
	;
	if v281 < int32(1) {
		goto L84
	} else {
		goto L85
	}
L80:
	;
	v375 = v281 + int32(-1)
	goto L21
L81:
	;
	v352 = *(*int32)(unsafe.Add(mBase, uint32(v345)+8))
	if v352 != 0 {
		v281 = v281 + int32(1)
		goto L79
	} else {
		goto L95
	}
L82:
	;
	v325 = v320
	goto L89
L83:
	;
	v304 = base.I64_reinterpret_f64(v303)
	v309 = int32(-1)
	v310 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v68)+12)))
	v316 = base.I32_rem_u_s(base.I32_wrap_i64(int64(base.Ui64(v304)>>(uint(int64(32))%64))+v304), v309<<(uint(v310)%32)^v309|int32(1))
	v320 = v117 + v316<<(uint(int32(5))%32)
	v321 = v303
	goto L82
L84:
	;
	v299 = base.F64_convert_i32_s(v281)
	if v281 == int32(0) {
		v320 = v117
		v321 = v299
		goto L82
	} else {
		goto L88
	}
L85:
	;
	if v281 <= v78 {
		goto L86
	} else {
		goto L87
	}
L86:
	;
	v293 = *(*int32)(unsafe.Add(mBase, uint32(v68)+20))
	v345 = v293 + v281<<(uint(int32(4))%32) + int32(-16)
	goto L81
L87:
	;
	v303 = base.F64_convert_i32_u(v281)
	goto L83
L88:
	;
	v303 = v299
	goto L83
L89:
	;
	v333 = *(*int32)(unsafe.Add(mBase, uint32(v325)+24))
	if v333 != int32(3) {
		goto L91
	} else {
		goto L92
	}
L90:
	;
	v345 = v339
	goto L81
L91:
	;
	v338 = *(*int32)(unsafe.Add(mBase, uint32(v325)+28))
	v339 = m.G398
	if v338 != 0 {
		v325 = v338
		goto L89
	} else {
		goto L94
	}
L92:
	;
	v336 = *(*float64)(unsafe.Add(mBase, uint32(v325)+16))
	if base.F64_ne(v336, v321) != 0 {
		goto L91
	} else {
		goto L93
	}
L93:
	;
	v345 = v325
	goto L81
L94:
	;
	goto L90
L95:
	;
	goto L80
L96:
	;
	return int32(0)
L97:
	;
	if v377 == int32(0) {
		v385 = v58
		goto L16
	} else {
		goto L98
	}
L98:
	;
	v383 = *(*int32)(unsafe.Add(mBase, uint32(v55)))
	v384 = *(*int32)(unsafe.Add(mBase, uint32(v383)+12))
	v385 = v384
	goto L16
}
func F_lua_pushcclosure(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
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
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
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
	var v47 int32
	_ = v47
	var v48 int64
	_ = v48
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v58 int32
	_ = v58
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v71 int64
	_ = v71
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v82 int64
	_ = v82
	var v84 int32
	_ = v84
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	var v104 int32
	_ = v104
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v9 = *(*int32)(unsafe.Add(mBase, uint32(v8)+68))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(v8)+64))
	if base.Ui32(v9) < base.Ui32(v10) {
		v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
		v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
		if v14 != v15 {
			v19 = *(*int32)(unsafe.Add(mBase, uint32(v14)+4))
			v20 = *(*int32)(unsafe.Add(mBase, uint32(v19)))
			v23 = v20 + int32(12)
		} else {
			v23 = l0 + int32(72)
		}
		v24 = *(*int32)(unsafe.Add(mBase, uint32(v23)))
		v25 = F_luaF_newCclosure(m, l0, l2, v24)
		mBase = m.M
		v26 = m.ExcPending
		if v26 != 0 {
			return
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v25)+16)) = l1
			v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			v31 = v28 - l2<<(uint(int32(4))%32)
			*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v31
			if l2 == int32(0) {
				v95 = v31
			} else {
				v36 = v25 + int32(24)
				if l2&int32(1) == int32(0) {
					v52 = l2
				} else {
					v42 = l2 + int32(-1)
					v44 = v42 << (uint(int32(4)) % 32)
					v45 = v36 + v44
					v46 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
					v47 = v46 + v44
					v48 = *(*int64)(unsafe.Add(mBase, uint32(v47)))
					*(*int64)(unsafe.Add(mBase, uint32(v45))) = v48
					v50 = *(*int32)(unsafe.Add(mBase, uint32(v47)+8))
					*(*int32)(unsafe.Add(mBase, uint32(v45)+8)) = v50
					v52 = v42
				}
				if l2 == int32(1) {
				} else {
					v58 = v52
					for {
						v64 = int32(4)
						v67 = v58<<(uint(v64)%32) + int32(-16)
						v68 = v36 + v67
						v69 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
						v70 = v69 + v67
						v71 = *(*int64)(unsafe.Add(mBase, uint32(v70)))
						*(*int64)(unsafe.Add(mBase, uint32(v68))) = v71
						v73 = *(*int32)(unsafe.Add(mBase, uint32(v70)+8))
						*(*int32)(unsafe.Add(mBase, uint32(v68)+8)) = v73
						v76 = v58 + int32(-2)
						v78 = v76 << (uint(v64) % 32)
						v79 = v36 + v78
						v80 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
						v81 = v80 + v78
						v82 = *(*int64)(unsafe.Add(mBase, uint32(v81)))
						*(*int64)(unsafe.Add(mBase, uint32(v79))) = v82
						v84 = *(*int32)(unsafe.Add(mBase, uint32(v81)+8))
						*(*int32)(unsafe.Add(mBase, uint32(v79)+8)) = v84
						if v76 != 0 {
							v58 = v76
							continue
						} else {
							break
						}
						break
					}
				}
				v93 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				v95 = v93
			}
			*(*int32)(unsafe.Add(mBase, uint32(v95)+8)) = int32(6)
			*(*int32)(unsafe.Add(mBase, uint32(v95))) = v25
			v104 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v104 + int32(16)
			return
		}
	} else {
		F_luaC_step(m, l0)
		mBase = m.M
		v13 = m.ExcPending
		if v13 != 0 {
			return
		} else {
			v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
			v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
			if v14 != v15 {
				v19 = *(*int32)(unsafe.Add(mBase, uint32(v14)+4))
				v20 = *(*int32)(unsafe.Add(mBase, uint32(v19)))
				v23 = v20 + int32(12)
			} else {
				v23 = l0 + int32(72)
			}
			v24 = *(*int32)(unsafe.Add(mBase, uint32(v23)))
			v25 = F_luaF_newCclosure(m, l0, l2, v24)
			mBase = m.M
			v26 = m.ExcPending
			if v26 != 0 {
				return
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v25)+16)) = l1
				v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				v31 = v28 - l2<<(uint(int32(4))%32)
				*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v31
				if l2 == int32(0) {
					v95 = v31
				} else {
					v36 = v25 + int32(24)
					if l2&int32(1) == int32(0) {
						v52 = l2
					} else {
						v42 = l2 + int32(-1)
						v44 = v42 << (uint(int32(4)) % 32)
						v45 = v36 + v44
						v46 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
						v47 = v46 + v44
						v48 = *(*int64)(unsafe.Add(mBase, uint32(v47)))
						*(*int64)(unsafe.Add(mBase, uint32(v45))) = v48
						v50 = *(*int32)(unsafe.Add(mBase, uint32(v47)+8))
						*(*int32)(unsafe.Add(mBase, uint32(v45)+8)) = v50
						v52 = v42
					}
					if l2 == int32(1) {
					} else {
						v58 = v52
						for {
							v64 = int32(4)
							v67 = v58<<(uint(v64)%32) + int32(-16)
							v68 = v36 + v67
							v69 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
							v70 = v69 + v67
							v71 = *(*int64)(unsafe.Add(mBase, uint32(v70)))
							*(*int64)(unsafe.Add(mBase, uint32(v68))) = v71
							v73 = *(*int32)(unsafe.Add(mBase, uint32(v70)+8))
							*(*int32)(unsafe.Add(mBase, uint32(v68)+8)) = v73
							v76 = v58 + int32(-2)
							v78 = v76 << (uint(v64) % 32)
							v79 = v36 + v78
							v80 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
							v81 = v80 + v78
							v82 = *(*int64)(unsafe.Add(mBase, uint32(v81)))
							*(*int64)(unsafe.Add(mBase, uint32(v79))) = v82
							v84 = *(*int32)(unsafe.Add(mBase, uint32(v81)+8))
							*(*int32)(unsafe.Add(mBase, uint32(v79)+8)) = v84
							if v76 != 0 {
								v58 = v76
								continue
							} else {
								break
							}
							break
						}
					}
					v93 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
					v95 = v93
				}
				*(*int32)(unsafe.Add(mBase, uint32(v95)+8)) = int32(6)
				*(*int32)(unsafe.Add(mBase, uint32(v95))) = v25
				v104 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v104 + int32(16)
				return
			}
		}
	}
}
func F_lua_pushfstring(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(v10)+68))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(v10)+64))
	if base.Ui32(v11) < base.Ui32(v12) {
		*(*int32)(unsafe.Add(mBase, uint32(v8)+12)) = l2
		v19 = F_luaO_pushvfstring(m, l0, l1, l2)
		mBase = m.M
		v20 = m.ExcPending
		if v20 != 0 {
			return int32(0)
		} else {
			m.G0 = v8 + int32(16)
			return v19
		}
	} else {
		F_luaC_step(m, l0)
		mBase = m.M
		v17 = m.ExcPending
		if v17 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v8)+12)) = l2
			v19 = F_luaO_pushvfstring(m, l0, l1, l2)
			mBase = m.M
			v20 = m.ExcPending
			if v20 != 0 {
				return int32(0)
			} else {
				m.G0 = v8 + int32(16)
				return v19
			}
		}
	}
}
func F_lua_pushthread(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v7 int32
	_ = v7
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v3)+8)) = int32(8)
	*(*int32)(unsafe.Add(mBase, uint32(v3))) = l0
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v7 + int32(16)
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(v11)+112))
	return base.B2i32(v12 == l0)
}
func F_lua_rawget(m *base.Module, l0 int32, l1 int32) {
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
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v66 int64
	_ = v66
	var v70 int32
	_ = v70
	if l1 < int32(1) {
		if l1 < int32(-9999) {
			switch l1 + int32(10002) {
			case 0:
				v55 = l0 + int32(72)
			case 1:
				v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
				v29 = *(*int32)(unsafe.Add(mBase, uint32(v28)+4))
				v30 = *(*int32)(unsafe.Add(mBase, uint32(v29)))
				v31 = *(*int32)(unsafe.Add(mBase, uint32(v30)+12))
				*(*int32)(unsafe.Add(mBase, uint32(l0)+96)) = int32(5)
				*(*int32)(unsafe.Add(mBase, uint32(l0)+88)) = v31
				v55 = l0 + int32(88)
			case 2:
				v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
				v55 = v25 + int32(96)
			default:
				v39 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
				v40 = *(*int32)(unsafe.Add(mBase, uint32(v39)+4))
				v41 = *(*int32)(unsafe.Add(mBase, uint32(v40)))
				v42 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v41)+7)))
				v43 = m.G398
				if base.Ui32(v42) < base.Ui32(int32(-10002)-l1) {
					v54 = v43
				} else {
					v54 = v41 + (int32(-10003)-l1)<<(uint(int32(4))%32) + int32(24)
				}
				v55 = v54
			}
		} else {
			v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			v55 = v19 + l1<<(uint(int32(4))%32)
		}
	} else {
		v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
		v13 = v8 + l1<<(uint(int32(4))%32) + int32(-16)
		v14 = m.G398
		if base.Ui32(v13) < base.Ui32(v7) {
			v16 = v13
		} else {
			v16 = v14
		}
		v55 = v16
	}
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v55)))
	v59 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v60 = int32(-16)
	v62 = F_luaH_get(m, v58, v59+v60)
	mBase = m.M
	v63 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v66 = *(*int64)(unsafe.Add(mBase, uint32(v62)))
	*(*int64)(unsafe.Add(mBase, uint32(v63+v60))) = v66
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v62)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v63+int32(-8)))) = v70
	return
}
func F_lua_rawseti(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
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
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v76 int64
	_ = v76
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	if l1 < int32(1) {
		if l1 < int32(-9999) {
			switch l1 + int32(10002) {
			case 0:
				v57 = l0 + int32(72)
			case 1:
				v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
				v30 = *(*int32)(unsafe.Add(mBase, uint32(v29)+4))
				v31 = *(*int32)(unsafe.Add(mBase, uint32(v30)))
				v32 = *(*int32)(unsafe.Add(mBase, uint32(v31)+12))
				*(*int32)(unsafe.Add(mBase, uint32(l0)+96)) = int32(5)
				*(*int32)(unsafe.Add(mBase, uint32(l0)+88)) = v32
				v57 = l0 + int32(88)
			case 2:
				v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
				v57 = v26 + int32(96)
			default:
				v40 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
				v41 = *(*int32)(unsafe.Add(mBase, uint32(v40)+4))
				v42 = *(*int32)(unsafe.Add(mBase, uint32(v41)))
				v43 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v42)+7)))
				v44 = m.G398
				if base.Ui32(v43) < base.Ui32(int32(-10002)-l1) {
					v55 = v44
				} else {
					v55 = v42 + (int32(-10003)-l1)<<(uint(int32(4))%32) + int32(24)
				}
				v57 = v55
			}
		} else {
			v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			v57 = v20 + l1<<(uint(int32(4))%32)
		}
	} else {
		v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
		v14 = v9 + l1<<(uint(int32(4))%32) + int32(-16)
		v15 = m.G398
		if base.Ui32(v14) < base.Ui32(v8) {
			v17 = v14
		} else {
			v17 = v15
		}
		v57 = v17
	}
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v57)))
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v59)+8))
	if v60 == int32(0) {
		v70 = v59
		v71 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		v72 = F_luaH_setnum(m, l0, v70, l2)
		mBase = m.M
		v73 = m.ExcPending
		if v73 != 0 {
			return
		} else {
			v76 = *(*int64)(unsafe.Add(mBase, uint32(v71+int32(-16))))
			*(*int64)(unsafe.Add(mBase, uint32(v72))) = v76
			v78 = int32(-8)
			v80 = *(*int32)(unsafe.Add(mBase, uint32(v71+v78)))
			*(*int32)(unsafe.Add(mBase, uint32(v72)+8)) = v80
			v82 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			v85 = *(*int32)(unsafe.Add(mBase, uint32(v82+v78)))
			if v85 < int32(4) {
				v111 = v82
			} else {
				v90 = *(*int32)(unsafe.Add(mBase, uint32(v82+int32(-16))))
				v91 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v90)+5)))
				if v91&int32(3) == int32(0) {
					v111 = v82
				} else {
					v96 = *(*int32)(unsafe.Add(mBase, uint32(v57)))
					v97 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v96)+5)))
					if v97&int32(4) == int32(0) {
						v111 = v82
					} else {
						v102 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
						v103 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v96)+5)))
						v105 = v103 & int32(251)
						*(*uint8)(unsafe.Add(mBase, uint32(v96)+5)) = uint8(v105)
						v107 = *(*int32)(unsafe.Add(mBase, uint32(v102)+40))
						*(*int32)(unsafe.Add(mBase, uint32(v96)+32)) = v107
						*(*int32)(unsafe.Add(mBase, uint32(v102)+40)) = v96
						v110 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
						v111 = v110
					}
				}
			}
			*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v111 + int32(-16)
			return
		}
	} else {
		v63 = m.G3
		F_luaG_runerror(m, l0, v63+int32(_a2241), int32(0))
		mBase = m.M
		v68 = m.ExcPending
		if v68 != 0 {
			return
		} else {
			v69 = *(*int32)(unsafe.Add(mBase, uint32(v57)))
			v70 = v69
			v71 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			v72 = F_luaH_setnum(m, l0, v70, l2)
			mBase = m.M
			v73 = m.ExcPending
			if v73 != 0 {
				return
			} else {
				v76 = *(*int64)(unsafe.Add(mBase, uint32(v71+int32(-16))))
				*(*int64)(unsafe.Add(mBase, uint32(v72))) = v76
				v78 = int32(-8)
				v80 = *(*int32)(unsafe.Add(mBase, uint32(v71+v78)))
				*(*int32)(unsafe.Add(mBase, uint32(v72)+8)) = v80
				v82 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				v85 = *(*int32)(unsafe.Add(mBase, uint32(v82+v78)))
				if v85 < int32(4) {
					v111 = v82
				} else {
					v90 = *(*int32)(unsafe.Add(mBase, uint32(v82+int32(-16))))
					v91 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v90)+5)))
					if v91&int32(3) == int32(0) {
						v111 = v82
					} else {
						v96 = *(*int32)(unsafe.Add(mBase, uint32(v57)))
						v97 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v96)+5)))
						if v97&int32(4) == int32(0) {
							v111 = v82
						} else {
							v102 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
							v103 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v96)+5)))
							v105 = v103 & int32(251)
							*(*uint8)(unsafe.Add(mBase, uint32(v96)+5)) = uint8(v105)
							v107 = *(*int32)(unsafe.Add(mBase, uint32(v102)+40))
							*(*int32)(unsafe.Add(mBase, uint32(v96)+32)) = v107
							*(*int32)(unsafe.Add(mBase, uint32(v102)+40)) = v96
							v110 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
							v111 = v110
						}
					}
				}
				*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v111 + int32(-16)
				return
			}
		}
	}
}
func F_lua_replace(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v62 int64
	_ = v62
	var v66 int32
	_ = v66
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
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
	var v97 int32
	_ = v97
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
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v117 int32
	_ = v117
	var v119 int32
	_ = v119
	var v122 int32
	_ = v122
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v133 int32
	_ = v133
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v143 int32
	_ = v143
	var v146 int32
	_ = v146
	var v149 int32
	_ = v149
	var v154 int32
	_ = v154
	if l1 != int32(-10001) {
		if l1 < int32(1) {
			if l1 < int32(-9999) {
				switch l1 + int32(10002) {
				case 0:
					v57 = l0 + int32(72)
					v59 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
					v62 = *(*int64)(unsafe.Add(mBase, uint32(v59+int32(-16))))
					*(*int64)(unsafe.Add(mBase, uint32(v57))) = v62
					v66 = *(*int32)(unsafe.Add(mBase, uint32(v59+int32(-8))))
					*(*int32)(unsafe.Add(mBase, uint32(v57)+8)) = v66
					if int32(-10003) < l1 {
					} else {
						v70 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
						v73 = *(*int32)(unsafe.Add(mBase, uint32(v70+int32(-8))))
						if v73 < int32(4) {
						} else {
							v78 = *(*int32)(unsafe.Add(mBase, uint32(v70+int32(-16))))
							v79 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v78)+5)))
							if v79&int32(3) == int32(0) {
							} else {
								v84 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
								v85 = *(*int32)(unsafe.Add(mBase, uint32(v84)+4))
								v86 = *(*int32)(unsafe.Add(mBase, uint32(v85)))
								v87 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v86)+5)))
								if v87&int32(4) == int32(0) {
								} else {
									v92 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
									v93 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v92)+21)))
									if v93 != int32(1) {
										v97 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v92)+20)))
										v100 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v86)+5)))
										v103 = v97&int32(3) | v100&int32(248)
										*(*uint8)(unsafe.Add(mBase, uint32(v86)+5)) = uint8(v103)
									} else {
										F_reallymarkobject(m, v92, v78)
										mBase = m.M
									}
								}
							}
						}
					}
				case 1:
					v105 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
					v106 = *(*int32)(unsafe.Add(mBase, uint32(v105)+4))
					v107 = *(*int32)(unsafe.Add(mBase, uint32(v106)))
					v108 = *(*int32)(unsafe.Add(mBase, uint32(v107)+12))
					*(*int32)(unsafe.Add(mBase, uint32(l0)+96)) = int32(5)
					*(*int32)(unsafe.Add(mBase, uint32(l0)+88)) = v108
					v112 = *(*int32)(unsafe.Add(mBase, uint32(v105)+4))
					v113 = *(*int32)(unsafe.Add(mBase, uint32(v112)))
					v114 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
					v117 = *(*int32)(unsafe.Add(mBase, uint32(v114+int32(-16))))
					*(*int32)(unsafe.Add(mBase, uint32(v113)+12)) = v117
					v119 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
					v122 = *(*int32)(unsafe.Add(mBase, uint32(v119+int32(-8))))
					if v122 < int32(4) {
					} else {
						v127 = *(*int32)(unsafe.Add(mBase, uint32(v119+int32(-16))))
						v128 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v127)+5)))
						if v128&int32(3) == int32(0) {
						} else {
							v133 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v113)+5)))
							if v133&int32(4) == int32(0) {
							} else {
								v138 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
								v139 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v138)+21)))
								if v139 != int32(1) {
									v143 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v138)+20)))
									v146 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v113)+5)))
									v149 = v143&int32(3) | v146&int32(248)
									*(*uint8)(unsafe.Add(mBase, uint32(v113)+5)) = uint8(v149)
								} else {
									F_reallymarkobject(m, v138, v127)
									mBase = m.M
								}
							}
						}
					}
				case 2:
					v54 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
					v57 = v54 + int32(96)
					v59 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
					v62 = *(*int64)(unsafe.Add(mBase, uint32(v59+int32(-16))))
					*(*int64)(unsafe.Add(mBase, uint32(v57))) = v62
					v66 = *(*int32)(unsafe.Add(mBase, uint32(v59+int32(-8))))
					*(*int32)(unsafe.Add(mBase, uint32(v57)+8)) = v66
					if int32(-10003) < l1 {
					} else {
						v70 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
						v73 = *(*int32)(unsafe.Add(mBase, uint32(v70+int32(-8))))
						if v73 < int32(4) {
						} else {
							v78 = *(*int32)(unsafe.Add(mBase, uint32(v70+int32(-16))))
							v79 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v78)+5)))
							if v79&int32(3) == int32(0) {
							} else {
								v84 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
								v85 = *(*int32)(unsafe.Add(mBase, uint32(v84)+4))
								v86 = *(*int32)(unsafe.Add(mBase, uint32(v85)))
								v87 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v86)+5)))
								if v87&int32(4) == int32(0) {
								} else {
									v92 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
									v93 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v92)+21)))
									if v93 != int32(1) {
										v97 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v92)+20)))
										v100 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v86)+5)))
										v103 = v97&int32(3) | v100&int32(248)
										*(*uint8)(unsafe.Add(mBase, uint32(v86)+5)) = uint8(v103)
									} else {
										F_reallymarkobject(m, v92, v78)
										mBase = m.M
									}
								}
							}
						}
					}
				default:
					v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
					v39 = *(*int32)(unsafe.Add(mBase, uint32(v38)+4))
					v40 = *(*int32)(unsafe.Add(mBase, uint32(v39)))
					v41 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v40)+7)))
					v42 = m.G398
					if base.Ui32(v41) < base.Ui32(int32(-10002)-l1) {
						v53 = v42
					} else {
						v53 = v40 + (int32(-10003)-l1)<<(uint(int32(4))%32) + int32(24)
					}
					v57 = v53
					v59 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
					v62 = *(*int64)(unsafe.Add(mBase, uint32(v59+int32(-16))))
					*(*int64)(unsafe.Add(mBase, uint32(v57))) = v62
					v66 = *(*int32)(unsafe.Add(mBase, uint32(v59+int32(-8))))
					*(*int32)(unsafe.Add(mBase, uint32(v57)+8)) = v66
					if int32(-10003) < l1 {
					} else {
						v70 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
						v73 = *(*int32)(unsafe.Add(mBase, uint32(v70+int32(-8))))
						if v73 < int32(4) {
						} else {
							v78 = *(*int32)(unsafe.Add(mBase, uint32(v70+int32(-16))))
							v79 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v78)+5)))
							if v79&int32(3) == int32(0) {
							} else {
								v84 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
								v85 = *(*int32)(unsafe.Add(mBase, uint32(v84)+4))
								v86 = *(*int32)(unsafe.Add(mBase, uint32(v85)))
								v87 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v86)+5)))
								if v87&int32(4) == int32(0) {
								} else {
									v92 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
									v93 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v92)+21)))
									if v93 != int32(1) {
										v97 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v92)+20)))
										v100 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v86)+5)))
										v103 = v97&int32(3) | v100&int32(248)
										*(*uint8)(unsafe.Add(mBase, uint32(v86)+5)) = uint8(v103)
									} else {
										F_reallymarkobject(m, v92, v78)
										mBase = m.M
									}
								}
							}
						}
					}
				}
			} else {
				v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				v57 = v30 + l1<<(uint(int32(4))%32)
				v59 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				v62 = *(*int64)(unsafe.Add(mBase, uint32(v59+int32(-16))))
				*(*int64)(unsafe.Add(mBase, uint32(v57))) = v62
				v66 = *(*int32)(unsafe.Add(mBase, uint32(v59+int32(-8))))
				*(*int32)(unsafe.Add(mBase, uint32(v57)+8)) = v66
				if int32(-10003) < l1 {
				} else {
					v70 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
					v73 = *(*int32)(unsafe.Add(mBase, uint32(v70+int32(-8))))
					if v73 < int32(4) {
					} else {
						v78 = *(*int32)(unsafe.Add(mBase, uint32(v70+int32(-16))))
						v79 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v78)+5)))
						if v79&int32(3) == int32(0) {
						} else {
							v84 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
							v85 = *(*int32)(unsafe.Add(mBase, uint32(v84)+4))
							v86 = *(*int32)(unsafe.Add(mBase, uint32(v85)))
							v87 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v86)+5)))
							if v87&int32(4) == int32(0) {
							} else {
								v92 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
								v93 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v92)+21)))
								if v93 != int32(1) {
									v97 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v92)+20)))
									v100 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v86)+5)))
									v103 = v97&int32(3) | v100&int32(248)
									*(*uint8)(unsafe.Add(mBase, uint32(v86)+5)) = uint8(v103)
								} else {
									F_reallymarkobject(m, v92, v78)
									mBase = m.M
								}
							}
						}
					}
				}
			}
		} else {
			v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
			v24 = v19 + l1<<(uint(int32(4))%32) + int32(-16)
			v25 = m.G398
			if base.Ui32(v24) < base.Ui32(v18) {
				v27 = v24
			} else {
				v27 = v25
			}
			v57 = v27
			v59 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			v62 = *(*int64)(unsafe.Add(mBase, uint32(v59+int32(-16))))
			*(*int64)(unsafe.Add(mBase, uint32(v57))) = v62
			v66 = *(*int32)(unsafe.Add(mBase, uint32(v59+int32(-8))))
			*(*int32)(unsafe.Add(mBase, uint32(v57)+8)) = v66
			if int32(-10003) < l1 {
			} else {
				v70 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				v73 = *(*int32)(unsafe.Add(mBase, uint32(v70+int32(-8))))
				if v73 < int32(4) {
				} else {
					v78 = *(*int32)(unsafe.Add(mBase, uint32(v70+int32(-16))))
					v79 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v78)+5)))
					if v79&int32(3) == int32(0) {
					} else {
						v84 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
						v85 = *(*int32)(unsafe.Add(mBase, uint32(v84)+4))
						v86 = *(*int32)(unsafe.Add(mBase, uint32(v85)))
						v87 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v86)+5)))
						if v87&int32(4) == int32(0) {
						} else {
							v92 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
							v93 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v92)+21)))
							if v93 != int32(1) {
								v97 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v92)+20)))
								v100 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v86)+5)))
								v103 = v97&int32(3) | v100&int32(248)
								*(*uint8)(unsafe.Add(mBase, uint32(v86)+5)) = uint8(v103)
							} else {
								F_reallymarkobject(m, v92, v78)
								mBase = m.M
							}
						}
					}
				}
			}
		}
		v154 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v154 + int32(-16)
		return
	} else {
		v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
		v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
		if v7 != v8 {
			v105 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
			v106 = *(*int32)(unsafe.Add(mBase, uint32(v105)+4))
			v107 = *(*int32)(unsafe.Add(mBase, uint32(v106)))
			v108 = *(*int32)(unsafe.Add(mBase, uint32(v107)+12))
			*(*int32)(unsafe.Add(mBase, uint32(l0)+96)) = int32(5)
			*(*int32)(unsafe.Add(mBase, uint32(l0)+88)) = v108
			v112 = *(*int32)(unsafe.Add(mBase, uint32(v105)+4))
			v113 = *(*int32)(unsafe.Add(mBase, uint32(v112)))
			v114 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			v117 = *(*int32)(unsafe.Add(mBase, uint32(v114+int32(-16))))
			*(*int32)(unsafe.Add(mBase, uint32(v113)+12)) = v117
			v119 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			v122 = *(*int32)(unsafe.Add(mBase, uint32(v119+int32(-8))))
			if v122 < int32(4) {
			} else {
				v127 = *(*int32)(unsafe.Add(mBase, uint32(v119+int32(-16))))
				v128 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v127)+5)))
				if v128&int32(3) == int32(0) {
				} else {
					v133 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v113)+5)))
					if v133&int32(4) == int32(0) {
					} else {
						v138 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
						v139 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v138)+21)))
						if v139 != int32(1) {
							v143 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v138)+20)))
							v146 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v113)+5)))
							v149 = v143&int32(3) | v146&int32(248)
							*(*uint8)(unsafe.Add(mBase, uint32(v113)+5)) = uint8(v149)
						} else {
							F_reallymarkobject(m, v138, v127)
							mBase = m.M
						}
					}
				}
			}
			v154 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v154 + int32(-16)
			return
		} else {
			v10 = m.G3
			F_luaG_runerror(m, l0, v10+int32(_a2239), int32(0))
			mBase = m.M
			v15 = m.ExcPending
			if v15 != 0 {
				return
			} else {
				v105 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
				v106 = *(*int32)(unsafe.Add(mBase, uint32(v105)+4))
				v107 = *(*int32)(unsafe.Add(mBase, uint32(v106)))
				v108 = *(*int32)(unsafe.Add(mBase, uint32(v107)+12))
				*(*int32)(unsafe.Add(mBase, uint32(l0)+96)) = int32(5)
				*(*int32)(unsafe.Add(mBase, uint32(l0)+88)) = v108
				v112 = *(*int32)(unsafe.Add(mBase, uint32(v105)+4))
				v113 = *(*int32)(unsafe.Add(mBase, uint32(v112)))
				v114 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				v117 = *(*int32)(unsafe.Add(mBase, uint32(v114+int32(-16))))
				*(*int32)(unsafe.Add(mBase, uint32(v113)+12)) = v117
				v119 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				v122 = *(*int32)(unsafe.Add(mBase, uint32(v119+int32(-8))))
				if v122 < int32(4) {
				} else {
					v127 = *(*int32)(unsafe.Add(mBase, uint32(v119+int32(-16))))
					v128 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v127)+5)))
					if v128&int32(3) == int32(0) {
					} else {
						v133 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v113)+5)))
						if v133&int32(4) == int32(0) {
						} else {
							v138 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
							v139 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v138)+21)))
							if v139 != int32(1) {
								v143 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v138)+20)))
								v146 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v113)+5)))
								v149 = v143&int32(3) | v146&int32(248)
								*(*uint8)(unsafe.Add(mBase, uint32(v113)+5)) = uint8(v149)
							} else {
								F_reallymarkobject(m, v138, v127)
								mBase = m.M
							}
						}
					}
				}
				v154 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v154 + int32(-16)
				return
			}
		}
	}
}
func F_lua_setfield(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v78 int32
	_ = v78
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v89 int32
	_ = v89
	var v95 int32
	_ = v95
	var v98 int32
	_ = v98
	var v104 int32
	_ = v104
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v124 int32
	_ = v124
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	if l1 < int32(1) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	if l2&int32(3) == int32(0) {
		v85 = l2
		goto L18
	} else {
		goto L19
	}
L2:
	;
	if l1 < int32(-9999) {
		goto L7
	} else {
		goto L8
	}
L3:
	;
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v19 = v14 + l1<<(uint(int32(4))%32) + int32(-16)
	v20 = m.G398
	if base.Ui32(v19) < base.Ui32(v13) {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v22 = v19
	goto L6
L5:
	;
	v22 = v20
	goto L6
L6:
	;
	v61 = v22
	goto L1
L7:
	;
	switch l1 + int32(10002) {
	case 0:
		goto L10
	case 1:
		goto L11
	case 2:
		goto L12
	default:
		goto L9
	}
L8:
	;
	v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v61 = v25 + l1<<(uint(int32(4))%32)
	goto L1
L9:
	;
	v45 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v45)+4))
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v46)))
	v48 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v47)+7)))
	v49 = m.G398
	if base.Ui32(v48) < base.Ui32(int32(-10002)-l1) {
		goto L13
	} else {
		goto L14
	}
L10:
	;
	v61 = l0 + int32(72)
	goto L1
L11:
	;
	v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v34)+4))
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v35)))
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v36)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+96)) = int32(5)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+88)) = v37
	v61 = l0 + int32(88)
	goto L1
L12:
	;
	v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v61 = v31 + int32(96)
	goto L1
L13:
	;
	v60 = v49
	goto L15
L14:
	;
	v60 = v47 + (int32(-10003)-l1)<<(uint(int32(4))%32) + int32(24)
	goto L15
L15:
	;
	v61 = v60
	goto L1
L16:
	;
	v119 = F_luaS_newlstr(m, l0, l2, v118)
	mBase = m.M
	v120 = m.ExcPending
	if v120 != 0 {
		goto L32
	} else {
		goto L33
	}
L17:
	;
	v118 = v110 - l2
	goto L16
L18:
	;
	v89 = v85
	goto L26
L19:
	;
	v71 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2))))
	if v71 != 0 {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	v74 = l2
	goto L22
L21:
	;
	v118 = l2 - l2
	goto L16
L22:
	;
	v78 = v74 + int32(1)
	if v78&int32(3) == int32(0) {
		v85 = v78
		goto L18
	} else {
		goto L24
	}
L24:
	;
	v83 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v78))))
	if v83 != 0 {
		v74 = v78
		goto L22
	} else {
		goto L25
	}
L25:
	;
	v110 = v78
	goto L17
L26:
	;
	v95 = *(*int32)(unsafe.Add(mBase, uint32(v89)))
	v98 = int32(-2139062144)
	if (int32(16843008)-v95|v95)&v98 == v98 {
		v89 = v89 + int32(4)
		goto L26
	} else {
		goto L28
	}
L27:
	;
	v104 = v89
	goto L29
L28:
	;
	goto L27
L29:
	;
	v108 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v104))))
	if v108 != 0 {
		v104 = v104 + int32(1)
		goto L29
	} else {
		goto L31
	}
L30:
	;
	v110 = v104
	goto L17
L31:
	;
	goto L30
L32:
	;
	return
L33:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9)+8)) = int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v9))) = v119
	v124 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	F_luaV_settable(m, l0, v61, v9, v124+int32(-16))
	mBase = m.M
	v128 = m.ExcPending
	if v128 != 0 {
		goto L32
	} else {
		goto L34
	}
L34:
	;
	v129 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v129 + int32(-16)
	m.G0 = v9 + int32(16)
	return
}
func F_lua_setlocal(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
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
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
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
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
	var v114 int64
	_ = v114
	var v116 int32
	_ = v116
	var v120 int32
	_ = v120
	var v124 int32
	_ = v124
	var v127 int32
	_ = v127
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l1)+96))
	v11 = v7 + v8*int32(24)
	v12 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
	v13 = *(*int32)(unsafe.Add(mBase, uint32(v12)+8))
	if v13 != int32(6) {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	v127 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v127 + int32(-16)
	return v124
L2:
	;
	v108 = v102 + l2<<(uint(int32(4))%32)
	v109 = int32(-16)
	v111 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v114 = *(*int64)(unsafe.Add(mBase, uint32(v111+v109)))
	*(*int64)(unsafe.Add(mBase, uint32(v108+v109))) = v114
	v116 = int32(-8)
	v120 = *(*int32)(unsafe.Add(mBase, uint32(v111+v116)))
	*(*int32)(unsafe.Add(mBase, uint32(v108+v116))) = v120
	v124 = v103
	goto L1
L3:
	;
	v87 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v11 == v87 {
		goto L22
	} else {
		goto L23
	}
L4:
	;
	v16 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
	v17 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16)+6)))
	if v17 != 0 {
		goto L3
	} else {
		goto L5
	}
L5:
	;
	v18 = *(*int32)(unsafe.Add(mBase, uint32(v16)+16))
	if v18 == int32(0) {
		goto L3
	} else {
		goto L6
	}
L6:
	;
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v11 == v21 {
		goto L8
	} else {
		goto L9
	}
L7:
	;
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v28)+12))
	v35 = (v29-v30)>>(uint(int32(2))%32) + int32(-1)
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v18)+56))
	if int32(1) <= v39 {
		goto L11
	} else {
		goto L12
	}
L8:
	;
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+12)) = v24
	v26 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v26)+16))
	v28 = v27
	v29 = v24
	goto L7
L9:
	;
	v23 = *(*int32)(unsafe.Add(mBase, uint32(v11)+12))
	v28 = v18
	v29 = v23
	goto L7
L10:
	;
	if v76 == int32(0) {
		goto L3
	} else {
		goto L21
	}
L11:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v18)+24))
	v45 = int32(0)
	v46 = l2
	goto L14
L12:
	;
	v76 = int32(0)
	goto L10
L13:
	;
	v76 = int32(0)
	goto L10
L14:
	;
	v53 = v43 + v45*int32(12)
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v53)+4))
	if v35 < v54 {
		goto L13
	} else {
		goto L16
	}
L15:
	;
	goto L13
L16:
	;
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v53)+8))
	if v56 <= v35 {
		v63 = v46
		goto L17
	} else {
		goto L18
	}
L17:
	;
	v65 = v45 + int32(1)
	if v65 != v39 {
		v45 = v65
		v46 = v63
		goto L14
	} else {
		goto L20
	}
L18:
	;
	v59 = v46 + int32(-1)
	if v59 != 0 {
		v63 = v59
		goto L17
	} else {
		goto L19
	}
L19:
	;
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v53)))
	v76 = v60 + int32(16)
	goto L10
L20:
	;
	goto L15
L21:
	;
	v79 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
	v102 = v79
	v103 = v76
	goto L2
L22:
	;
	v89 = l0 + int32(8)
	goto L24
L23:
	;
	v89 = v11 + int32(28)
	goto L24
L24:
	;
	v90 = *(*int32)(unsafe.Add(mBase, uint32(v89)))
	v91 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
	v92 = m.G3
	v93 = int32(0)
	if l2 < int32(1) {
		v124 = v93
		goto L1
	} else {
		goto L25
	}
L25:
	;
	if (v90-v91)>>(uint(int32(4))%32) < l2 {
		v124 = v93
		goto L1
	} else {
		goto L26
	}
L26:
	;
	v102 = v91
	v103 = v92 + int32(_a2242)
	goto L2
}
func F_lua_settop(m *base.Module, l0 int32, l1 int32) {
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
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v27 int32
	_ = v27
	if l1 < int32(0) {
		v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		v27 = v20 + l1<<(uint(int32(4))%32) + int32(16)
	} else {
		v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
		v10 = v7 + l1<<(uint(int32(4))%32)
		if base.Ui32(v10) <= base.Ui32(v6) {
			v27 = v10
		} else {
			v14 = v6
			for {
				*(*int32)(unsafe.Add(mBase, uint32(v14)+8)) = int32(0)
				v18 = v14 + int32(16)
				if base.Ui32(v18) < base.Ui32(v10) {
					v14 = v18
					continue
				} else {
					break
				}
				break
			}
			v27 = v10
		}
	}
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v27
	return
}
func F_lua_setupvalue(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
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
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v102 int32
	_ = v102
	var v104 int64
	_ = v104
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v131 int32
	_ = v131
	var v134 int32
	_ = v134
	var v137 int32
	_ = v137
	var v141 int32
	_ = v141
	if l1 < int32(1) {
		if l1 < int32(-9999) {
			switch l1 + int32(10002) {
			case 0:
				v57 = l0 + int32(72)
			case 1:
				v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
				v31 = *(*int32)(unsafe.Add(mBase, uint32(v30)+4))
				v32 = *(*int32)(unsafe.Add(mBase, uint32(v31)))
				v33 = *(*int32)(unsafe.Add(mBase, uint32(v32)+12))
				*(*int32)(unsafe.Add(mBase, uint32(l0)+96)) = int32(5)
				*(*int32)(unsafe.Add(mBase, uint32(l0)+88)) = v33
				v57 = l0 + int32(88)
			case 2:
				v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
				v57 = v27 + int32(96)
			default:
				v41 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
				v42 = *(*int32)(unsafe.Add(mBase, uint32(v41)+4))
				v43 = *(*int32)(unsafe.Add(mBase, uint32(v42)))
				v44 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v43)+7)))
				v45 = m.G398
				if base.Ui32(v44) < base.Ui32(int32(-10002)-l1) {
					v56 = v45
				} else {
					v56 = v43 + (int32(-10003)-l1)<<(uint(int32(4))%32) + int32(24)
				}
				v57 = v56
			}
		} else {
			v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			v57 = v21 + l1<<(uint(int32(4))%32)
		}
	} else {
		v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
		v15 = v10 + l1<<(uint(int32(4))%32) + int32(-16)
		v16 = m.G398
		if base.Ui32(v15) < base.Ui32(v9) {
			v18 = v15
		} else {
			v18 = v16
		}
		v57 = v18
	}
	v60 = int32(0)
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v57)+8))
	if v61 != int32(6) {
		v141 = v60
	} else {
		v64 = *(*int32)(unsafe.Add(mBase, uint32(v57)))
		v65 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v64)+6)))
		if v65 == int32(0) {
			if l2 < int32(1) {
				v141 = v60
			} else {
				v82 = *(*int32)(unsafe.Add(mBase, uint32(v64)+16))
				v83 = *(*int32)(unsafe.Add(mBase, uint32(v82)+36))
				if v83 < l2 {
					v141 = v60
				} else {
					v85 = *(*int32)(unsafe.Add(mBase, uint32(v82)+28))
					v89 = l2<<(uint(int32(2))%32) + int32(-4)
					v91 = *(*int32)(unsafe.Add(mBase, uint32(v85+v89)))
					v95 = *(*int32)(unsafe.Add(mBase, uint32(v64+v89)+20))
					v96 = *(*int32)(unsafe.Add(mBase, uint32(v95)+8))
					v97 = v96
					v98 = v91 + int32(16)
					v100 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
					v102 = v100 + int32(-16)
					*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v102
					v104 = *(*int64)(unsafe.Add(mBase, uint32(v102)))
					*(*int64)(unsafe.Add(mBase, uint32(v97))) = v104
					v108 = *(*int32)(unsafe.Add(mBase, uint32(v100+int32(-8))))
					*(*int32)(unsafe.Add(mBase, uint32(v97)+8)) = v108
					v110 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
					v111 = *(*int32)(unsafe.Add(mBase, uint32(v110)+8))
					if v111 < int32(4) {
						v141 = v98
					} else {
						v114 = *(*int32)(unsafe.Add(mBase, uint32(v110)))
						v115 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v114)+5)))
						if v115&int32(3) == int32(0) {
							v141 = v98
						} else {
							v120 = *(*int32)(unsafe.Add(mBase, uint32(v57)))
							v121 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v120)+5)))
							if v121&int32(4) == int32(0) {
								v141 = v98
							} else {
								v126 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
								v127 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v126)+21)))
								if v127 != int32(1) {
									v131 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v126)+20)))
									v134 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v120)+5)))
									v137 = v131&int32(3) | v134&int32(248)
									*(*uint8)(unsafe.Add(mBase, uint32(v120)+5)) = uint8(v137)
								} else {
									F_reallymarkobject(m, v126, v114)
									mBase = m.M
								}
								v141 = v98
							}
						}
					}
				}
			}
		} else {
			if l2 < int32(1) {
				v141 = v60
			} else {
				v70 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v64)+7)))
				if base.Ui32(v70) < base.Ui32(l2) {
					v141 = v60
				} else {
					v72 = m.G3
					v97 = l2<<(uint(int32(4))%32) + v64 + int32(8)
					v98 = v72 + int32(_a188)
					v100 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
					v102 = v100 + int32(-16)
					*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v102
					v104 = *(*int64)(unsafe.Add(mBase, uint32(v102)))
					*(*int64)(unsafe.Add(mBase, uint32(v97))) = v104
					v108 = *(*int32)(unsafe.Add(mBase, uint32(v100+int32(-8))))
					*(*int32)(unsafe.Add(mBase, uint32(v97)+8)) = v108
					v110 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
					v111 = *(*int32)(unsafe.Add(mBase, uint32(v110)+8))
					if v111 < int32(4) {
						v141 = v98
					} else {
						v114 = *(*int32)(unsafe.Add(mBase, uint32(v110)))
						v115 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v114)+5)))
						if v115&int32(3) == int32(0) {
							v141 = v98
						} else {
							v120 = *(*int32)(unsafe.Add(mBase, uint32(v57)))
							v121 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v120)+5)))
							if v121&int32(4) == int32(0) {
								v141 = v98
							} else {
								v126 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
								v127 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v126)+21)))
								if v127 != int32(1) {
									v131 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v126)+20)))
									v134 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v120)+5)))
									v137 = v131&int32(3) | v134&int32(248)
									*(*uint8)(unsafe.Add(mBase, uint32(v120)+5)) = uint8(v137)
								} else {
									F_reallymarkobject(m, v126, v114)
									mBase = m.M
								}
								v141 = v98
							}
						}
					}
				}
			}
		}
	}
	return v141
}
func F_lua_toboolean(m *base.Module, l0 int32, l1 int32) int32 {
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
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v63 int32
	_ = v63
	if l1 < int32(1) {
		if l1 < int32(-9999) {
			switch l1 + int32(10002) {
			case 0:
				v55 = l0 + int32(72)
			case 1:
				v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
				v28 = *(*int32)(unsafe.Add(mBase, uint32(v27)+4))
				v29 = *(*int32)(unsafe.Add(mBase, uint32(v28)))
				v30 = *(*int32)(unsafe.Add(mBase, uint32(v29)+12))
				*(*int32)(unsafe.Add(mBase, uint32(l0)+96)) = int32(5)
				*(*int32)(unsafe.Add(mBase, uint32(l0)+88)) = v30
				v55 = l0 + int32(88)
			case 2:
				v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
				v55 = v24 + int32(96)
			default:
				v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
				v39 = *(*int32)(unsafe.Add(mBase, uint32(v38)+4))
				v40 = *(*int32)(unsafe.Add(mBase, uint32(v39)))
				v41 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v40)+7)))
				v42 = m.G398
				if base.Ui32(v41) < base.Ui32(int32(-10002)-l1) {
					v53 = v42
				} else {
					v53 = v40 + (int32(-10003)-l1)<<(uint(int32(4))%32) + int32(24)
				}
				v55 = v53
			}
		} else {
			v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			v55 = v18 + l1<<(uint(int32(4))%32)
		}
	} else {
		v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
		v12 = v7 + l1<<(uint(int32(4))%32) + int32(-16)
		v13 = m.G398
		if base.Ui32(v12) < base.Ui32(v6) {
			v15 = v12
		} else {
			v15 = v13
		}
		v55 = v15
	}
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v55)+8))
	switch v57 {
	case 0:
		v63 = v57
		return v63
	case 1:
		v58 = *(*int32)(unsafe.Add(mBase, uint32(v55)))
		return base.B2i32(v58 != int32(0))
	default:
		v63 = int32(1)
		return v63
	}
}
func F_lua_tointeger(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v72 float64
	_ = v72
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	if l1 < int32(1) {
		if l1 < int32(-9999) {
			switch l1 + int32(10002) {
			case 0:
				v61 = l0 + int32(72)
			case 1:
				v33 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
				v34 = *(*int32)(unsafe.Add(mBase, uint32(v33)+4))
				v35 = *(*int32)(unsafe.Add(mBase, uint32(v34)))
				v36 = *(*int32)(unsafe.Add(mBase, uint32(v35)+12))
				*(*int32)(unsafe.Add(mBase, uint32(l0)+96)) = int32(5)
				*(*int32)(unsafe.Add(mBase, uint32(l0)+88)) = v36
				v61 = l0 + int32(88)
			case 2:
				v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
				v61 = v30 + int32(96)
			default:
				v44 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
				v45 = *(*int32)(unsafe.Add(mBase, uint32(v44)+4))
				v46 = *(*int32)(unsafe.Add(mBase, uint32(v45)))
				v47 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v46)+7)))
				v48 = m.G398
				if base.Ui32(v47) < base.Ui32(int32(-10002)-l1) {
					v59 = v48
				} else {
					v59 = v46 + (int32(-10003)-l1)<<(uint(int32(4))%32) + int32(24)
				}
				v61 = v59
			}
		} else {
			v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			v61 = v24 + l1<<(uint(int32(4))%32)
		}
	} else {
		v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
		v18 = v13 + l1<<(uint(int32(4))%32) + int32(-16)
		v19 = m.G398
		if base.Ui32(v18) < base.Ui32(v12) {
			v21 = v18
		} else {
			v21 = v19
		}
		v61 = v21
	}
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v61)+8))
	if v63 == int32(3) {
		v71 = v61
		v72 = *(*float64)(unsafe.Add(mBase, uint32(v71)))
		if base.F64_lt(base.F64_abs(v72), float64(2.147483648e+09)) == int32(0) {
			v80 = int32(-2147483648)
		} else {
			v78 = base.I32_trunc_f64_s(v72)
			v80 = v78
		}
		m.G0 = v8 + int32(16)
		return v80
	} else {
		v66 = F_luaV_tonumber(m, v61, v8)
		mBase = m.M
		v69 = m.ExcPending
		if v69 != 0 {
			return int32(0)
		} else {
			if v66 != 0 {
				v71 = v66
				v72 = *(*float64)(unsafe.Add(mBase, uint32(v71)))
				if base.F64_lt(base.F64_abs(v72), float64(2.147483648e+09)) == int32(0) {
					v80 = int32(-2147483648)
				} else {
					v78 = base.I32_trunc_f64_s(v72)
					v80 = v78
				}
			} else {
				v80 = int32(0)
			}
			m.G0 = v8 + int32(16)
			return v80
		}
	}
}
func F_lua_topointer(m *base.Module, l0 int32, l1 int32) int32 {
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
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
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
	var v112 int32
	_ = v112
	var v114 int32
	_ = v114
	var v116 int32
	_ = v116
	var v119 int32
	_ = v119
	var v123 int32
	_ = v123
	var v127 int32
	_ = v127
	if l1 < int32(1) {
		if l1 < int32(-9999) {
			switch l1 + int32(10002) {
			case 0:
				v55 = l0 + int32(72)
			case 1:
				v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
				v29 = *(*int32)(unsafe.Add(mBase, uint32(v28)+4))
				v30 = *(*int32)(unsafe.Add(mBase, uint32(v29)))
				v31 = *(*int32)(unsafe.Add(mBase, uint32(v30)+12))
				*(*int32)(unsafe.Add(mBase, uint32(l0)+96)) = int32(5)
				*(*int32)(unsafe.Add(mBase, uint32(l0)+88)) = v31
				v55 = l0 + int32(88)
			case 2:
				v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
				v55 = v25 + int32(96)
			default:
				v39 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
				v40 = *(*int32)(unsafe.Add(mBase, uint32(v39)+4))
				v41 = *(*int32)(unsafe.Add(mBase, uint32(v40)))
				v42 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v41)+7)))
				v43 = m.G398
				if base.Ui32(v42) < base.Ui32(int32(-10002)-l1) {
					v54 = v43
				} else {
					v54 = v41 + (int32(-10003)-l1)<<(uint(int32(4))%32) + int32(24)
				}
				v55 = v54
			}
		} else {
			v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			v55 = v19 + l1<<(uint(int32(4))%32)
		}
	} else {
		v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
		v13 = v8 + l1<<(uint(int32(4))%32) + int32(-16)
		v14 = m.G398
		if base.Ui32(v13) < base.Ui32(v7) {
			v16 = v13
		} else {
			v16 = v14
		}
		v55 = v16
	}
	v57 = int32(0)
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v55)+8))
	switch v58 + int32(-2) {
	case 0, 5:
		if l1 < int32(1) {
			if l1 < int32(-9999) {
				switch l1 + int32(10002) {
				case 0:
					v114 = l0 + int32(72)
				case 1:
					v86 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
					v87 = *(*int32)(unsafe.Add(mBase, uint32(v86)+4))
					v88 = *(*int32)(unsafe.Add(mBase, uint32(v87)))
					v89 = *(*int32)(unsafe.Add(mBase, uint32(v88)+12))
					*(*int32)(unsafe.Add(mBase, uint32(l0)+96)) = int32(5)
					*(*int32)(unsafe.Add(mBase, uint32(l0)+88)) = v89
					v114 = l0 + int32(88)
				case 2:
					v83 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
					v114 = v83 + int32(96)
				default:
					v97 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
					v98 = *(*int32)(unsafe.Add(mBase, uint32(v97)+4))
					v99 = *(*int32)(unsafe.Add(mBase, uint32(v98)))
					v100 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v99)+7)))
					v101 = m.G398
					if base.Ui32(v100) < base.Ui32(int32(-10002)-l1) {
						v112 = v101
					} else {
						v112 = v99 + (int32(-10003)-l1)<<(uint(int32(4))%32) + int32(24)
					}
					v114 = v112
				}
			} else {
				v77 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				v114 = v77 + l1<<(uint(int32(4))%32)
			}
		} else {
			v65 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			v66 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
			v71 = v66 + l1<<(uint(int32(4))%32) + int32(-16)
			v72 = m.G398
			if base.Ui32(v71) < base.Ui32(v65) {
				v74 = v71
			} else {
				v74 = v72
			}
			v114 = v74
		}
		v116 = *(*int32)(unsafe.Add(mBase, uint32(v114)+8))
		switch v116 + int32(-2) {
		case 0:
			v123 = *(*int32)(unsafe.Add(mBase, uint32(v114)))
			v127 = v123
			return v127
		default:
			v127 = v57
			return v127
		case 5:
			v119 = *(*int32)(unsafe.Add(mBase, uint32(v114)))
			return v119 + int32(24)
		}
	default:
		v127 = v57
		return v127
	case 3, 4, 6:
		v61 = *(*int32)(unsafe.Add(mBase, uint32(v55)))
		return v61
	}
}
func F_lua_touserdata(m *base.Module, l0 int32, l1 int32) int32 {
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
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	if l1 < int32(1) {
		if l1 < int32(-9999) {
			switch l1 + int32(10002) {
			case 0:
				v55 = l0 + int32(72)
			case 1:
				v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
				v28 = *(*int32)(unsafe.Add(mBase, uint32(v27)+4))
				v29 = *(*int32)(unsafe.Add(mBase, uint32(v28)))
				v30 = *(*int32)(unsafe.Add(mBase, uint32(v29)+12))
				*(*int32)(unsafe.Add(mBase, uint32(l0)+96)) = int32(5)
				*(*int32)(unsafe.Add(mBase, uint32(l0)+88)) = v30
				v55 = l0 + int32(88)
			case 2:
				v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
				v55 = v24 + int32(96)
			default:
				v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
				v39 = *(*int32)(unsafe.Add(mBase, uint32(v38)+4))
				v40 = *(*int32)(unsafe.Add(mBase, uint32(v39)))
				v41 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v40)+7)))
				v42 = m.G398
				if base.Ui32(v41) < base.Ui32(int32(-10002)-l1) {
					v53 = v42
				} else {
					v53 = v40 + (int32(-10003)-l1)<<(uint(int32(4))%32) + int32(24)
				}
				v55 = v53
			}
		} else {
			v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			v55 = v18 + l1<<(uint(int32(4))%32)
		}
	} else {
		v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
		v12 = v7 + l1<<(uint(int32(4))%32) + int32(-16)
		v13 = m.G398
		if base.Ui32(v12) < base.Ui32(v6) {
			v15 = v12
		} else {
			v15 = v13
		}
		v55 = v15
	}
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v55)+8))
	switch v58 + int32(-2) {
	case 0:
		v65 = *(*int32)(unsafe.Add(mBase, uint32(v55)))
		v66 = v65
		return v66
	default:
		v66 = int32(0)
		return v66
	case 5:
		v61 = *(*int32)(unsafe.Add(mBase, uint32(v55)))
		return v61 + int32(24)
	}
}
func F_resetLuaContext(m *base.Module, l0 int32) {
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	v4 = F_lua_gc(m, l0, int32(2), int32(0))
	v5 = m.ExcPending
	if v5 != 0 {
		return
	} else {
		F_lua_close(m, l0)
		v7 = m.ExcPending
		if v7 != 0 {
			return
		} else {
			return
		}
	}
}
