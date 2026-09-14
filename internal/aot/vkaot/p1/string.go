package p1

import (
	base "github.com/shibukawa/valkeymem/internal/aot/vkaot/base"
	"unsafe"
)

func F_compareStringObjectsWithFlags(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
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
	var v25 int32
	_ = v25
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v66 int32
	_ = v66
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v92 int32
	_ = v92
	var v99 int32
	_ = v99
	var v102 int32
	_ = v102
	var v105 int32
	_ = v105
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v129 int32
	_ = v129
	var v136 int32
	_ = v136
	var v139 int32
	_ = v139
	var v142 int32
	_ = v142
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v150 int32
	_ = v150
	var v165 int32
	_ = v165
	var v168 int32
	_ = v168
	var v170 int32
	_ = v170
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v178 int64
	_ = v178
	var v187 int32
	_ = v187
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v198 int64
	_ = v198
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v224 int32
	_ = v224
	var v229 int32
	_ = v229
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v249 int32
	_ = v249
	var v256 int32
	_ = v256
	var v259 int32
	_ = v259
	var v262 int32
	_ = v262
	var v265 int32
	_ = v265
	var v266 int32
	_ = v266
	var v270 int32
	_ = v270
	var v285 int32
	_ = v285
	var v288 int32
	_ = v288
	var v289 int32
	_ = v289
	var v296 int32
	_ = v296
	var v303 int32
	_ = v303
	var v306 int32
	_ = v306
	var v309 int32
	_ = v309
	var v312 int32
	_ = v312
	var v313 int32
	_ = v313
	var v322 int32
	_ = v322
	var v323 int32
	_ = v323
	var v328 int32
	_ = v328
	var v329 int32
	_ = v329
	var v333 int32
	_ = v333
	var v340 int32
	_ = v340
	var v343 int32
	_ = v343
	var v346 int32
	_ = v346
	var v349 int32
	_ = v349
	var v350 int32
	_ = v350
	var v354 int32
	_ = v354
	var v369 int32
	_ = v369
	var v372 int32
	_ = v372
	var v374 int32
	_ = v374
	var v378 int64
	_ = v378
	var v387 int32
	_ = v387
	var v391 int32
	_ = v391
	var v396 int32
	_ = v396
	var v397 int32
	_ = v397
	var v398 int64
	_ = v398
	var v399 int32
	_ = v399
	var v400 int32
	_ = v400
	var v419 int32
	_ = v419
	var v420 int32
	_ = v420
	var v421 int32
	_ = v421
	var v428 int32
	_ = v428
	var v429 int32
	_ = v429
	var v431 int32
	_ = v431
	var v439 int32
	_ = v439
	var v440 int32
	_ = v440
	var v441 int32
	_ = v441
	var v444 int32
	_ = v444
	var v445 int32
	_ = v445
	var v447 int32
	_ = v447
	var v448 int32
	_ = v448
	var v450 int32
	_ = v450
	var v452 int32
	_ = v452
	var v455 int32
	_ = v455
	var v456 int32
	_ = v456
	var v457 int32
	_ = v457
	var v462 int32
	_ = v462
	var v463 int32
	_ = v463
	var v464 int32
	_ = v464
	var v467 int32
	_ = v467
	var v468 int32
	_ = v468
	var v469 int32
	_ = v469
	var v472 int32
	_ = v472
	var v473 int32
	_ = v473
	var v475 int32
	_ = v475
	var v480 int32
	_ = v480
	var v495 int32
	_ = v495
	var v497 int32
	_ = v497
	var v499 int32
	_ = v499
	var v515 int32
	_ = v515
	v9 = m.G0
	v11 = v9 - int32(256)
	m.G0 = v11
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v13&int32(15) != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	F__serverAssertWithInfo(m, int32(0), l0, int32(_a1761), int32(_a1758), int32(958))
	mBase = m.M
	v515 = m.ExcPending
	if v515 != 0 {
		goto L134
	} else {
		goto L135
	}
L2:
	;
	v16 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	if v16&int32(15) != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	if l0 != l1 {
		goto L5
	} else {
		goto L6
	}
L4:
	;
	m.G0 = v11 + int32(256)
	return v499
L5:
	;
	switch int32(base.Ui32(v13)>>(uint(int32(4))%32)) & int32(15) {
	case 0, 8:
		goto L9
	default:
		goto L8
	}
L6:
	;
	v499 = int32(0)
	goto L4
L7:
	;
	v224 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	switch int32(base.Ui32(v224)>>(uint(int32(4))%32)) & int32(15) {
	case 0, 8:
		goto L60
	default:
		goto L59
	}
L8:
	;
	v109 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v109&int32(4) == int32(0) {
		goto L33
	} else {
		goto L34
	}
L9:
	;
	v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v25&int32(4) == int32(0) {
		goto L11
	} else {
		goto L12
	}
L10:
	;
	v92 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v85+int32(-1)))))
	switch v92 & int32(7) {
	case 0:
		goto L31
	case 1:
		goto L30
	case 2:
		goto L29
	case 3:
		goto L28
	case 4:
		goto L27
	default:
		v220 = v85
		v221 = int32(0)
		goto L7
	}
L11:
	;
	v84 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v85 = v84
	goto L10
L12:
	;
	if v25&int32(1) != 0 {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	v34 = int32(16)
	goto L15
L14:
	;
	v34 = int32(8)
	goto L15
L15:
	;
	v35 = l0 + v34
	if v25&int32(2) == int32(0) {
		v66 = v35
		goto L16
	} else {
		goto L17
	}
L16:
	;
	goto L26
L17:
	;
	v40 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v35))))
	v41 = v35 + v40
	v45 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v41))))
	switch v45 & int32(7) {
	case 0:
		goto L23
	case 1:
		goto L22
	case 2:
		goto L21
	case 3:
		goto L20
	case 4:
		goto L19
	default:
		v62 = int32(0)
		goto L18
	}
L18:
	;
	v66 = v41 + int32(1) + v62 + int32(1)
	goto L16
L19:
	;
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v41+int32(-16))))
	v62 = v61
	goto L18
L20:
	;
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v41+int32(-8))))
	v62 = v58
	goto L18
L21:
	;
	v55 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v41+int32(-4)))))
	v62 = v55
	goto L18
L22:
	;
	v52 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v41+int32(-2)))))
	v62 = v52
	goto L18
L23:
	;
	v62 = int32(base.Ui32(v45) >> (uint(int32(3)) % 32))
	goto L18
L24:
	;
	v85 = v66 + v81
	goto L10
L25:
	;
	goto L24
L26:
	;
	v81 = *(*int32)(unsafe.Add(mBase, _consts[308]))
	goto L25
L27:
	;
	v108 = *(*int32)(unsafe.Add(mBase, uint32(v85+int32(-17))))
	v220 = v85
	v221 = v108
	goto L7
L28:
	;
	v105 = *(*int32)(unsafe.Add(mBase, uint32(v85+int32(-9))))
	v220 = v85
	v221 = v105
	goto L7
L29:
	;
	v102 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v85+int32(-5)))))
	v220 = v85
	v221 = v102
	goto L7
L30:
	;
	v99 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v85+int32(-3)))))
	v220 = v85
	v221 = v99
	goto L7
L31:
	;
	v220 = v85
	v221 = int32(base.Ui32(v92) >> (uint(int32(3)) % 32))
	goto L7
L32:
	;
	v173 = int32(128)
	v174 = v11 + v173
	v178 = base.I64_extend_i32_s(v170)
	if v178 <= int64(-1) {
		goto L53
	} else {
		goto L54
	}
L33:
	;
	v168 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v170 = v168
	goto L32
L34:
	;
	if v109&int32(1) != 0 {
		goto L35
	} else {
		goto L36
	}
L35:
	;
	v118 = int32(16)
	goto L37
L36:
	;
	v118 = int32(8)
	goto L37
L37:
	;
	v119 = l0 + v118
	if v109&int32(2) == int32(0) {
		v150 = v119
		goto L38
	} else {
		goto L39
	}
L38:
	;
	goto L48
L39:
	;
	v124 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v119))))
	v125 = v119 + v124
	v129 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v125))))
	switch v129 & int32(7) {
	case 0:
		goto L45
	case 1:
		goto L44
	case 2:
		goto L43
	case 3:
		goto L42
	case 4:
		goto L41
	default:
		v146 = int32(0)
		goto L40
	}
L40:
	;
	v150 = v125 + int32(1) + v146 + int32(1)
	goto L38
L41:
	;
	v145 = *(*int32)(unsafe.Add(mBase, uint32(v125+int32(-16))))
	v146 = v145
	goto L40
L42:
	;
	v142 = *(*int32)(unsafe.Add(mBase, uint32(v125+int32(-8))))
	v146 = v142
	goto L40
L43:
	;
	v139 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v125+int32(-4)))))
	v146 = v139
	goto L40
L44:
	;
	v136 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v125+int32(-2)))))
	v146 = v136
	goto L40
L45:
	;
	v146 = int32(base.Ui32(v129) >> (uint(int32(3)) % 32))
	goto L40
L46:
	;
	v170 = v150 + v165
	goto L32
L47:
	;
	goto L46
L48:
	;
	v165 = *(*int32)(unsafe.Add(mBase, _consts[308]))
	goto L47
L49:
	;
	v220 = v174
	v221 = v219
	goto L7
L50:
	;
	v219 = int32(0)
	goto L49
L52:
	;
	v200 = F_ull2string(m, v196, v197, v198)
	mBase = m.M
	if v200 == int32(0) {
		goto L50
	} else {
		goto L56
	}
L53:
	;
	goto L55
L54:
	;
	v196 = v174
	v197 = v173
	v198 = v178
	v199 = int32(0)
	goto L52
L55:
	;
	v187 = int32(45)
	*(*uint8)(unsafe.Add(mBase, uint32(v174))) = uint8(v187)
	v196 = v11 + int32(129)
	v197 = int32(127)
	v198 = int64(0) - v178
	v199 = int32(1)
	goto L52
L56:
	;
	v219 = v200 + v199
	goto L49
L58:
	;
	if l2&int32(2) == int32(0) {
		goto L109
	} else {
		goto L110
	}
L59:
	;
	v313 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v313&int32(4) == int32(0) {
		goto L84
	} else {
		goto L85
	}
L60:
	;
	v229 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v229&int32(4) == int32(0) {
		goto L62
	} else {
		goto L63
	}
L61:
	;
	v296 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v289+int32(-1)))))
	switch v296 & int32(7) {
	case 0:
		goto L82
	case 1:
		goto L81
	case 2:
		goto L80
	case 3:
		goto L79
	case 4:
		goto L78
	default:
		v420 = v289
		v421 = int32(0)
		goto L58
	}
L62:
	;
	v288 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v289 = v288
	goto L61
L63:
	;
	if v229&int32(1) != 0 {
		goto L64
	} else {
		goto L65
	}
L64:
	;
	v238 = int32(16)
	goto L66
L65:
	;
	v238 = int32(8)
	goto L66
L66:
	;
	v239 = l1 + v238
	if v229&int32(2) == int32(0) {
		v270 = v239
		goto L67
	} else {
		goto L68
	}
L67:
	;
	goto L77
L68:
	;
	v244 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v239))))
	v245 = v239 + v244
	v249 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v245))))
	switch v249 & int32(7) {
	case 0:
		goto L74
	case 1:
		goto L73
	case 2:
		goto L72
	case 3:
		goto L71
	case 4:
		goto L70
	default:
		v266 = int32(0)
		goto L69
	}
L69:
	;
	v270 = v245 + int32(1) + v266 + int32(1)
	goto L67
L70:
	;
	v265 = *(*int32)(unsafe.Add(mBase, uint32(v245+int32(-16))))
	v266 = v265
	goto L69
L71:
	;
	v262 = *(*int32)(unsafe.Add(mBase, uint32(v245+int32(-8))))
	v266 = v262
	goto L69
L72:
	;
	v259 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v245+int32(-4)))))
	v266 = v259
	goto L69
L73:
	;
	v256 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v245+int32(-2)))))
	v266 = v256
	goto L69
L74:
	;
	v266 = int32(base.Ui32(v249) >> (uint(int32(3)) % 32))
	goto L69
L75:
	;
	v289 = v270 + v285
	goto L61
L76:
	;
	goto L75
L77:
	;
	v285 = *(*int32)(unsafe.Add(mBase, _consts[308]))
	goto L76
L78:
	;
	v312 = *(*int32)(unsafe.Add(mBase, uint32(v289+int32(-17))))
	v420 = v289
	v421 = v312
	goto L58
L79:
	;
	v309 = *(*int32)(unsafe.Add(mBase, uint32(v289+int32(-9))))
	v420 = v289
	v421 = v309
	goto L58
L80:
	;
	v306 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v289+int32(-5)))))
	v420 = v289
	v421 = v306
	goto L58
L81:
	;
	v303 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v289+int32(-3)))))
	v420 = v289
	v421 = v303
	goto L58
L82:
	;
	v420 = v289
	v421 = int32(base.Ui32(v296) >> (uint(int32(3)) % 32))
	goto L58
L83:
	;
	v378 = base.I64_extend_i32_s(v374)
	if v378 <= int64(-1) {
		goto L104
	} else {
		goto L105
	}
L84:
	;
	v372 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v374 = v372
	goto L83
L85:
	;
	if v313&int32(1) != 0 {
		goto L86
	} else {
		goto L87
	}
L86:
	;
	v322 = int32(16)
	goto L88
L87:
	;
	v322 = int32(8)
	goto L88
L88:
	;
	v323 = l1 + v322
	if v313&int32(2) == int32(0) {
		v354 = v323
		goto L89
	} else {
		goto L90
	}
L89:
	;
	goto L99
L90:
	;
	v328 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v323))))
	v329 = v323 + v328
	v333 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v329))))
	switch v333 & int32(7) {
	case 0:
		goto L96
	case 1:
		goto L95
	case 2:
		goto L94
	case 3:
		goto L93
	case 4:
		goto L92
	default:
		v350 = int32(0)
		goto L91
	}
L91:
	;
	v354 = v329 + int32(1) + v350 + int32(1)
	goto L89
L92:
	;
	v349 = *(*int32)(unsafe.Add(mBase, uint32(v329+int32(-16))))
	v350 = v349
	goto L91
L93:
	;
	v346 = *(*int32)(unsafe.Add(mBase, uint32(v329+int32(-8))))
	v350 = v346
	goto L91
L94:
	;
	v343 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v329+int32(-4)))))
	v350 = v343
	goto L91
L95:
	;
	v340 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v329+int32(-2)))))
	v350 = v340
	goto L91
L96:
	;
	v350 = int32(base.Ui32(v333) >> (uint(int32(3)) % 32))
	goto L91
L97:
	;
	v374 = v354 + v369
	goto L83
L98:
	;
	goto L97
L99:
	;
	v369 = *(*int32)(unsafe.Add(mBase, _consts[308]))
	goto L98
L100:
	;
	v420 = v11
	v421 = v419
	goto L58
L101:
	;
	v419 = int32(0)
	goto L100
L103:
	;
	v400 = F_ull2string(m, v396, v397, v398)
	mBase = m.M
	if v400 == int32(0) {
		goto L101
	} else {
		goto L107
	}
L104:
	;
	goto L106
L105:
	;
	v396 = v11
	v397 = int32(128)
	v398 = v378
	v399 = int32(0)
	goto L103
L106:
	;
	v387 = int32(45)
	*(*uint8)(unsafe.Add(mBase, uint32(v11))) = uint8(v387)
	v391 = int32(1)
	v396 = v11 + v391
	v397 = int32(127)
	v398 = int64(0) - v378
	v399 = v391
	goto L103
L107:
	;
	v419 = v400 + v399
	goto L100
L109:
	;
	if base.Ui32(v221) < base.Ui32(v421) {
		goto L112
	} else {
		goto L113
	}
L110:
	;
	v428 = F___get_tp(m)
	mBase = m.M
	v429 = F___strcoll_l(m, v220, v420, v420)
	mBase = m.M
	goto L111
L111:
	;
	v499 = v429
	goto L4
L112:
	;
	v431 = v221
	goto L114
L113:
	;
	v431 = v421
	goto L114
L114:
	;
	if base.Ui32(v431) < base.Ui32(int32(4)) {
		v455 = v220
		v456 = v420
		v457 = v431
		goto L118
	} else {
		goto L119
	}
L115:
	;
	if v495 != 0 {
		goto L131
	} else {
		goto L132
	}
L116:
	;
	v495 = int32(0)
	goto L115
L117:
	;
	v467 = v462
	v468 = v463
	v469 = v464
	goto L127
L118:
	;
	if v457 == int32(0) {
		goto L116
	} else {
		goto L125
	}
L119:
	;
	if (v420|v220)&int32(3) != 0 {
		v462 = v220
		v463 = v420
		v464 = v431
		goto L117
	} else {
		goto L120
	}
L120:
	;
	v439 = v220
	v440 = v420
	v441 = v431
	goto L121
L121:
	;
	v444 = *(*int32)(unsafe.Add(mBase, uint32(v439)))
	v445 = *(*int32)(unsafe.Add(mBase, uint32(v440)))
	if v444 != v445 {
		v462 = v439
		v463 = v440
		v464 = v441
		goto L117
	} else {
		goto L123
	}
L122:
	;
	v455 = v450
	v456 = v448
	v457 = v452
	goto L118
L123:
	;
	v447 = int32(4)
	v448 = v440 + v447
	v450 = v439 + v447
	v452 = v441 + int32(-4)
	if base.Ui32(int32(3)) < base.Ui32(v452) {
		v439 = v450
		v440 = v448
		v441 = v452
		goto L121
	} else {
		goto L124
	}
L124:
	;
	goto L122
L125:
	;
	v462 = v455
	v463 = v456
	v464 = v457
	goto L117
L126:
	;
	v495 = v472 - v473
	goto L115
L127:
	;
	v472 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v467))))
	v473 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v468))))
	if v472 != v473 {
		goto L126
	} else {
		goto L129
	}
L129:
	;
	v475 = int32(1)
	v480 = v469 + int32(-1)
	if v480 == int32(0) {
		goto L116
	} else {
		goto L130
	}
L130:
	;
	v467 = v467 + v475
	v468 = v468 + v475
	v469 = v480
	goto L127
L131:
	;
	v497 = v495
	goto L133
L132:
	;
	v497 = v221 - v421
	goto L133
L133:
	;
	v499 = v497
	goto L4
L134:
	;
	return int32(0)
L135:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_createStringObjectFromLongLong(m *base.Module, l0 int64) int32 {
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	v3 = F_createStringObjectFromLongLongWithOptions(m, l0, int32(0))
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		return v3
	}
}
func F_createStringObjectFromLongLongWithOptions(m *base.Module, l0 int64, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v1 int64
	_ = v1
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v16 int32
	_ = v16
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v33 int32
	_ = v33
	var v48 int32
	_ = v48
	var v52 int32
	_ = v52
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v59 int64
	_ = v59
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v80 int32
	_ = v80
	var v94 int32
	_ = v94
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v117 int32
	_ = v117
	v1 = l0
	v5 = m.G0
	v7 = v5 - int32(32)
	m.G0 = v7
	if base.Ui64(int64(9999)) < base.Ui64(v1) {
		if base.Ui64(int64(4294967295)) < base.Ui64(v1+int64(2147483648)) {
			if v1 <= int64(-1) {
				v48 = int32(45)
				*(*uint8)(unsafe.Add(mBase, uint32(v7))) = uint8(v48)
				v52 = int32(1)
				v57 = v7 + v52
				v58 = int32(20)
				v59 = int64(0) - v1
				v60 = v52
			} else {
				v57 = v7
				v58 = int32(21)
				v59 = v1
				v60 = int32(0)
			}
			v61 = F_ull2string(m, v57, v58, v59)
			mBase = m.M
			if v61 == int32(0) {
				v80 = int32(0)
			} else {
				v80 = v61 + v60
			}
			if base.Ui32(int32(255)) < base.Ui32(v80) {
				v105 = F_sdsnewlen(m, v7, v80)
				mBase = m.M
				v106 = m.ExcPending
				if v106 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v7)+28)) = int32(0)
					v112 = F_zmalloc_usable(m, int32(12), v7+int32(28))
					mBase = m.M
					v113 = m.ExcPending
					if v113 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v112)+8)) = v105
						*(*int64)(unsafe.Add(mBase, uint32(v112))) = int64(34359738368)
						v117 = v112
						m.G0 = v7 + int32(32)
						return v117
					}
				}
			} else {
				v94 = *(*int32)(unsafe.Add(mBase, _consts[308]))
				if base.Ui32(int32(128)) < base.Ui32(v80+v94+int32(9)) {
					v105 = F_sdsnewlen(m, v7, v80)
					mBase = m.M
					v106 = m.ExcPending
					if v106 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v7)+28)) = int32(0)
						v112 = F_zmalloc_usable(m, int32(12), v7+int32(28))
						mBase = m.M
						v113 = m.ExcPending
						if v113 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v112)+8)) = v105
							*(*int64)(unsafe.Add(mBase, uint32(v112))) = int64(34359738368)
							v117 = v112
							m.G0 = v7 + int32(32)
							return v117
						}
					}
				} else {
					v103 = F_createEmbeddedStringObjectWithKeyAndExpire(m, v7, v80, int32(0), int64(-1))
					mBase = m.M
					v104 = m.ExcPending
					if v104 != 0 {
						return int32(0)
					} else {
						v117 = v103
						m.G0 = v7 + int32(32)
						return v117
					}
				}
			}
		} else {
			if l1 == int32(2) {
				if v1 <= int64(-1) {
					v48 = int32(45)
					*(*uint8)(unsafe.Add(mBase, uint32(v7))) = uint8(v48)
					v52 = int32(1)
					v57 = v7 + v52
					v58 = int32(20)
					v59 = int64(0) - v1
					v60 = v52
				} else {
					v57 = v7
					v58 = int32(21)
					v59 = v1
					v60 = int32(0)
				}
				v61 = F_ull2string(m, v57, v58, v59)
				mBase = m.M
				if v61 == int32(0) {
					v80 = int32(0)
				} else {
					v80 = v61 + v60
				}
				if base.Ui32(int32(255)) < base.Ui32(v80) {
					v105 = F_sdsnewlen(m, v7, v80)
					mBase = m.M
					v106 = m.ExcPending
					if v106 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v7)+28)) = int32(0)
						v112 = F_zmalloc_usable(m, int32(12), v7+int32(28))
						mBase = m.M
						v113 = m.ExcPending
						if v113 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v112)+8)) = v105
							*(*int64)(unsafe.Add(mBase, uint32(v112))) = int64(34359738368)
							v117 = v112
							m.G0 = v7 + int32(32)
							return v117
						}
					}
				} else {
					v94 = *(*int32)(unsafe.Add(mBase, _consts[308]))
					if base.Ui32(int32(128)) < base.Ui32(v80+v94+int32(9)) {
						v105 = F_sdsnewlen(m, v7, v80)
						mBase = m.M
						v106 = m.ExcPending
						if v106 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v7)+28)) = int32(0)
							v112 = F_zmalloc_usable(m, int32(12), v7+int32(28))
							mBase = m.M
							v113 = m.ExcPending
							if v113 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v112)+8)) = v105
								*(*int64)(unsafe.Add(mBase, uint32(v112))) = int64(34359738368)
								v117 = v112
								m.G0 = v7 + int32(32)
								return v117
							}
						}
					} else {
						v103 = F_createEmbeddedStringObjectWithKeyAndExpire(m, v7, v80, int32(0), int64(-1))
						mBase = m.M
						v104 = m.ExcPending
						if v104 != 0 {
							return int32(0)
						} else {
							v117 = v103
							m.G0 = v7 + int32(32)
							return v117
						}
					}
				}
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v7))) = int32(0)
				v26 = F_zmalloc_usable(m, int32(12), v7)
				mBase = m.M
				v29 = m.ExcPending
				if v29 != 0 {
					return int32(0)
				} else {
					*(*int64)(unsafe.Add(mBase, uint32(v26))) = int64(34359738368)
					*(*uint32)(unsafe.Add(mBase, uint32(v26)+8)) = uint32(v1)
					v33 = *(*int32)(unsafe.Add(mBase, uint32(v26)))
					*(*int32)(unsafe.Add(mBase, uint32(v26))) = v33&int32(-241) | int32(16)
					v117 = v26
					m.G0 = v7 + int32(32)
					return v117
				}
			}
		}
	} else {
		if l1 != 0 {
			if base.Ui64(int64(4294967295)) < base.Ui64(v1+int64(2147483648)) {
				if v1 <= int64(-1) {
					v48 = int32(45)
					*(*uint8)(unsafe.Add(mBase, uint32(v7))) = uint8(v48)
					v52 = int32(1)
					v57 = v7 + v52
					v58 = int32(20)
					v59 = int64(0) - v1
					v60 = v52
				} else {
					v57 = v7
					v58 = int32(21)
					v59 = v1
					v60 = int32(0)
				}
				v61 = F_ull2string(m, v57, v58, v59)
				mBase = m.M
				if v61 == int32(0) {
					v80 = int32(0)
				} else {
					v80 = v61 + v60
				}
				if base.Ui32(int32(255)) < base.Ui32(v80) {
					v105 = F_sdsnewlen(m, v7, v80)
					mBase = m.M
					v106 = m.ExcPending
					if v106 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v7)+28)) = int32(0)
						v112 = F_zmalloc_usable(m, int32(12), v7+int32(28))
						mBase = m.M
						v113 = m.ExcPending
						if v113 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v112)+8)) = v105
							*(*int64)(unsafe.Add(mBase, uint32(v112))) = int64(34359738368)
							v117 = v112
							m.G0 = v7 + int32(32)
							return v117
						}
					}
				} else {
					v94 = *(*int32)(unsafe.Add(mBase, _consts[308]))
					if base.Ui32(int32(128)) < base.Ui32(v80+v94+int32(9)) {
						v105 = F_sdsnewlen(m, v7, v80)
						mBase = m.M
						v106 = m.ExcPending
						if v106 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v7)+28)) = int32(0)
							v112 = F_zmalloc_usable(m, int32(12), v7+int32(28))
							mBase = m.M
							v113 = m.ExcPending
							if v113 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v112)+8)) = v105
								*(*int64)(unsafe.Add(mBase, uint32(v112))) = int64(34359738368)
								v117 = v112
								m.G0 = v7 + int32(32)
								return v117
							}
						}
					} else {
						v103 = F_createEmbeddedStringObjectWithKeyAndExpire(m, v7, v80, int32(0), int64(-1))
						mBase = m.M
						v104 = m.ExcPending
						if v104 != 0 {
							return int32(0)
						} else {
							v117 = v103
							m.G0 = v7 + int32(32)
							return v117
						}
					}
				}
			} else {
				if l1 == int32(2) {
					if v1 <= int64(-1) {
						v48 = int32(45)
						*(*uint8)(unsafe.Add(mBase, uint32(v7))) = uint8(v48)
						v52 = int32(1)
						v57 = v7 + v52
						v58 = int32(20)
						v59 = int64(0) - v1
						v60 = v52
					} else {
						v57 = v7
						v58 = int32(21)
						v59 = v1
						v60 = int32(0)
					}
					v61 = F_ull2string(m, v57, v58, v59)
					mBase = m.M
					if v61 == int32(0) {
						v80 = int32(0)
					} else {
						v80 = v61 + v60
					}
					if base.Ui32(int32(255)) < base.Ui32(v80) {
						v105 = F_sdsnewlen(m, v7, v80)
						mBase = m.M
						v106 = m.ExcPending
						if v106 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v7)+28)) = int32(0)
							v112 = F_zmalloc_usable(m, int32(12), v7+int32(28))
							mBase = m.M
							v113 = m.ExcPending
							if v113 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v112)+8)) = v105
								*(*int64)(unsafe.Add(mBase, uint32(v112))) = int64(34359738368)
								v117 = v112
								m.G0 = v7 + int32(32)
								return v117
							}
						}
					} else {
						v94 = *(*int32)(unsafe.Add(mBase, _consts[308]))
						if base.Ui32(int32(128)) < base.Ui32(v80+v94+int32(9)) {
							v105 = F_sdsnewlen(m, v7, v80)
							mBase = m.M
							v106 = m.ExcPending
							if v106 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v7)+28)) = int32(0)
								v112 = F_zmalloc_usable(m, int32(12), v7+int32(28))
								mBase = m.M
								v113 = m.ExcPending
								if v113 != 0 {
									return int32(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v112)+8)) = v105
									*(*int64)(unsafe.Add(mBase, uint32(v112))) = int64(34359738368)
									v117 = v112
									m.G0 = v7 + int32(32)
									return v117
								}
							}
						} else {
							v103 = F_createEmbeddedStringObjectWithKeyAndExpire(m, v7, v80, int32(0), int64(-1))
							mBase = m.M
							v104 = m.ExcPending
							if v104 != 0 {
								return int32(0)
							} else {
								v117 = v103
								m.G0 = v7 + int32(32)
								return v117
							}
						}
					}
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v7))) = int32(0)
					v26 = F_zmalloc_usable(m, int32(12), v7)
					mBase = m.M
					v29 = m.ExcPending
					if v29 != 0 {
						return int32(0)
					} else {
						*(*int64)(unsafe.Add(mBase, uint32(v26))) = int64(34359738368)
						*(*uint32)(unsafe.Add(mBase, uint32(v26)+8)) = uint32(v1)
						v33 = *(*int32)(unsafe.Add(mBase, uint32(v26)))
						*(*int32)(unsafe.Add(mBase, uint32(v26))) = v33&int32(-241) | int32(16)
						v117 = v26
						m.G0 = v7 + int32(32)
						return v117
					}
				}
			}
		} else {
			v16 = *(*int32)(unsafe.Add(mBase, uint32(base.I32_wrap_i64(v1)<<(uint(int32(2))%32))+uint32(_consts[506])))
			v117 = v16
			m.G0 = v7 + int32(32)
			return v117
		}
	}
}
func F_createStringObjectFromLongLongWithSds(m *base.Module, l0 int64) int32 {
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	v3 = F_createStringObjectFromLongLongWithOptions(m, l0, int32(2))
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		return v3
	}
}
func F_getStringObjectLen(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	var v31 int32
	_ = v31
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v48 int32
	_ = v48
	var v52 int32
	_ = v52
	var v56 int32
	_ = v56
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v72 int32
	_ = v72
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v4&int32(15) != 0 {
		F__serverAssertWithInfo(m, int32(0), l0, int32(_a1629), int32(_a1630), int32(183))
		mBase = m.M
		v72 = m.ExcPending
		if v72 != 0 {
			return int32(0)
		} else {
			F_abort(m)
			mBase = m.M
			base.Wasm_trap_unreachable()
			for {
			}
		}
	} else {
		v7 = int32(0)
		switch int32(base.Ui32(v4)>>(uint(int32(4))%32)) & int32(15) {
		case 0:
			v12 = F_objectGetVal(m, l0)
			mBase = m.M
			v15 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12+int32(-1)))))
			switch v15 & int32(7) {
			case 0:
				return int32(base.Ui32(v15) >> (uint(int32(3)) % 32))
			case 1:
				v23 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12+int32(-3)))))
				return v23
			case 2:
				v27 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v12+int32(-5)))))
				return v27
			case 3:
				v31 = *(*int32)(unsafe.Add(mBase, uint32(v12+int32(-9))))
				return v31
			case 4:
				v35 = *(*int32)(unsafe.Add(mBase, uint32(v12+int32(-17))))
				return v35
			default:
				v63 = v7
				return v63
			}
		default:
			v63 = v7
			return v63
		case 8:
			v37 = F_objectGetVal(m, l0)
			mBase = m.M
			v40 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v37+int32(-1)))))
			switch v40 & int32(7) {
			case 0:
				return int32(base.Ui32(v40) >> (uint(int32(3)) % 32))
			case 1:
				v48 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v37+int32(-3)))))
				return v48
			case 2:
				v52 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v37+int32(-5)))))
				return v52
			case 3:
				v56 = *(*int32)(unsafe.Add(mBase, uint32(v37+int32(-9))))
				return v56
			case 4:
				v60 = *(*int32)(unsafe.Add(mBase, uint32(v37+int32(-17))))
				v63 = v60
				return v63
			default:
				v63 = v7
				return v63
			}
		}
	}
}
func F_lookupStringForBitCommand(m *base.Module, l0 int32, l1 int64, l2 int32) int32 {
	mBase := m.M
	_ = mBase
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
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
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
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v85 int32
	_ = v85
	var v88 int32
	_ = v88
	var v95 int32
	_ = v95
	var v98 int32
	_ = v98
	var v101 int32
	_ = v101
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v119 int32
	_ = v119
	var v122 int32
	_ = v122
	v9 = m.G0
	v11 = v9 - int32(16)
	m.G0 = v11
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v16 = *(*int32)(unsafe.Add(mBase, uint32(v15)+4))
	v17 = F_lookupKeyWrite(m, v14, v16)
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		return int32(0)
	} else {
		v22 = F_checkType(m, l0, v17, int32(0))
		mBase = m.M
		v23 = m.ExcPending
		if v23 != 0 {
			return int32(0)
		} else {
			if v22 != 0 {
				v122 = int32(0)
				m.G0 = v11 + int32(16)
				return v122
			} else {
				if l2 == int32(0) {
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(l2))) = int32(0)
				}
				v30 = base.I32_wrap_i64(int64(base.Ui64(l1) >> (uint(int64(3)) % 64)))
				if v17 != 0 {
					v47 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
					v48 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
					v49 = *(*int32)(unsafe.Add(mBase, uint32(v48)+4))
					v50 = F_dbUnshareStringValue(m, v47, v49, v17)
					mBase = m.M
					v51 = m.ExcPending
					if v51 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v11)+12)) = v50
						v54 = F_objectGetVal(m, v50)
						mBase = m.M
						v57 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v54+int32(-1)))))
						switch v57 & int32(7) {
						case 0:
							v74 = int32(base.Ui32(v57) >> (uint(int32(3)) % 32))
						case 1:
							v64 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v54+int32(-3)))))
							v74 = v64
						case 2:
							v67 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v54+int32(-5)))))
							v74 = v67
						case 3:
							v70 = *(*int32)(unsafe.Add(mBase, uint32(v54+int32(-9))))
							v74 = v70
						case 4:
							v73 = *(*int32)(unsafe.Add(mBase, uint32(v54+int32(-17))))
							v74 = v73
						default:
							v74 = int32(0)
						}
						v75 = F_objectGetVal(m, v50)
						mBase = m.M
						v78 = F_sdsgrowzero(m, v75, v30+int32(1))
						mBase = m.M
						v79 = m.ExcPending
						if v79 != 0 {
							return int32(0)
						} else {
							F_objectSetVal(m, v50, v78)
							mBase = m.M
							v81 = m.ExcPending
							if v81 != 0 {
								return int32(0)
							} else {
								if l2 == int32(0) {
								} else {
									v85 = F_objectGetVal(m, v50)
									mBase = m.M
									v88 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v85+int32(-1)))))
									switch v88 & int32(7) {
									case 0:
										v105 = int32(base.Ui32(v88) >> (uint(int32(3)) % 32))
									case 1:
										v95 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v85+int32(-3)))))
										v105 = v95
									case 2:
										v98 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v85+int32(-5)))))
										v105 = v98
									case 3:
										v101 = *(*int32)(unsafe.Add(mBase, uint32(v85+int32(-9))))
										v105 = v101
									case 4:
										v104 = *(*int32)(unsafe.Add(mBase, uint32(v85+int32(-17))))
										v105 = v104
									default:
										v105 = int32(0)
									}
									if v74 == v105 {
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(l2))) = int32(1)
									}
								}
								v119 = *(*int32)(unsafe.Add(mBase, uint32(v11)+12))
								v122 = v119
								m.G0 = v11 + int32(16)
								return v122
							}
						}
					}
				} else {
					v31 = int32(0)
					v35 = F_sdsnewlen(m, v31, v30+int32(1))
					mBase = m.M
					v36 = m.ExcPending
					if v36 != 0 {
						return int32(0)
					} else {
						v37 = F_createObject(m, v31, v35)
						mBase = m.M
						v38 = m.ExcPending
						if v38 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v11)+12)) = v37
							v40 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
							v41 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
							v42 = *(*int32)(unsafe.Add(mBase, uint32(v41)+4))
							F_dbAdd(m, v40, v42, v11+int32(12))
							mBase = m.M
							v46 = m.ExcPending
							if v46 != 0 {
								return int32(0)
							} else {
								if l2 != 0 {
									*(*int32)(unsafe.Add(mBase, uint32(l2))) = int32(1)
								} else {
								}
								v119 = *(*int32)(unsafe.Add(mBase, uint32(v11)+12))
								v122 = v119
								m.G0 = v11 + int32(16)
								return v122
							}
						}
					}
				}
			}
		}
	}
}
func F_stringObjectLen(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v12 int32
	_ = v12
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v32 int32
	_ = v32
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v53 int32
	_ = v53
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v79 int32
	_ = v79
	var v87 int32
	_ = v87
	var v91 int32
	_ = v91
	var v95 int32
	_ = v95
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v121 int32
	_ = v121
	var v128 int32
	_ = v128
	var v131 int32
	_ = v131
	var v134 int32
	_ = v134
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v142 int32
	_ = v142
	var v157 int32
	_ = v157
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v165 int64
	_ = v165
	var v170 int32
	_ = v170
	var v172 int64
	_ = v172
	var v175 int64
	_ = v175
	var v176 int32
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
	var v222 int32
	_ = v222
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v231 int32
	_ = v231
	var v234 int64
	_ = v234
	var v235 int32
	_ = v235
	var v259 int32
	_ = v259
	var v265 int32
	_ = v265
	var v273 int32
	_ = v273
	var v279 int32
	_ = v279
	var v282 int32
	_ = v282
	var v286 int64
	_ = v286
	var v288 int32
	_ = v288
	var v295 int32
	_ = v295
	var v297 int32
	_ = v297
	var v308 int32
	_ = v308
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v5&int32(15) != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	F__serverAssertWithInfo(m, int32(0), l0, int32(_a1629), int32(_a1758), int32(1018))
	mBase = m.M
	v308 = m.ExcPending
	if v308 != 0 {
		goto L111
	} else {
		goto L112
	}
L2:
	;
	switch int32(base.Ui32(v5)>>(uint(int32(4))%32)) & int32(15) {
	case 0, 8:
		goto L5
	default:
		goto L4
	}
L3:
	;
	return v297
L4:
	;
	v101 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v101&int32(4) == int32(0) {
		goto L29
	} else {
		goto L30
	}
L5:
	;
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v12&int32(4) == int32(0) {
		goto L7
	} else {
		goto L8
	}
L6:
	;
	v79 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v72+int32(-1)))))
	switch v79 & int32(7) {
	case 0:
		goto L27
	case 1:
		goto L26
	case 2:
		goto L25
	case 3:
		goto L24
	case 4:
		goto L23
	default:
		v297 = int32(0)
		goto L3
	}
L7:
	;
	v71 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v72 = v71
	goto L6
L8:
	;
	if v12&int32(1) != 0 {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	v21 = int32(16)
	goto L11
L10:
	;
	v21 = int32(8)
	goto L11
L11:
	;
	v22 = l0 + v21
	if v12&int32(2) == int32(0) {
		v53 = v22
		goto L12
	} else {
		goto L13
	}
L12:
	;
	goto L22
L13:
	;
	v27 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22))))
	v28 = v22 + v27
	v32 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28))))
	switch v32 & int32(7) {
	case 0:
		goto L19
	case 1:
		goto L18
	case 2:
		goto L17
	case 3:
		goto L16
	case 4:
		goto L15
	default:
		v49 = int32(0)
		goto L14
	}
L14:
	;
	v53 = v28 + int32(1) + v49 + int32(1)
	goto L12
L15:
	;
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v28+int32(-16))))
	v49 = v48
	goto L14
L16:
	;
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v28+int32(-8))))
	v49 = v45
	goto L14
L17:
	;
	v42 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v28+int32(-4)))))
	v49 = v42
	goto L14
L18:
	;
	v39 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28+int32(-2)))))
	v49 = v39
	goto L14
L19:
	;
	v49 = int32(base.Ui32(v32) >> (uint(int32(3)) % 32))
	goto L14
L20:
	;
	v72 = v53 + v68
	goto L6
L21:
	;
	goto L20
L22:
	;
	v68 = *(*int32)(unsafe.Add(mBase, _consts[308]))
	goto L21
L23:
	;
	v99 = *(*int32)(unsafe.Add(mBase, uint32(v72+int32(-17))))
	return v99
L24:
	;
	v95 = *(*int32)(unsafe.Add(mBase, uint32(v72+int32(-9))))
	return v95
L25:
	;
	v91 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v72+int32(-5)))))
	return v91
L26:
	;
	v87 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v72+int32(-3)))))
	return v87
L27:
	;
	return int32(base.Ui32(v79) >> (uint(int32(3)) % 32))
L28:
	;
	v165 = base.I64_extend_i32_s(v161)
	if int64(-1) < v165 {
		goto L46
	} else {
		goto L47
	}
L29:
	;
	v160 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v161 = v160
	goto L28
L30:
	;
	if v101&int32(1) != 0 {
		goto L31
	} else {
		goto L32
	}
L31:
	;
	v110 = int32(16)
	goto L33
L32:
	;
	v110 = int32(8)
	goto L33
L33:
	;
	v111 = l0 + v110
	if v101&int32(2) == int32(0) {
		v142 = v111
		goto L34
	} else {
		goto L35
	}
L34:
	;
	goto L44
L35:
	;
	v116 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v111))))
	v117 = v111 + v116
	v121 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v117))))
	switch v121 & int32(7) {
	case 0:
		goto L41
	case 1:
		goto L40
	case 2:
		goto L39
	case 3:
		goto L38
	case 4:
		goto L37
	default:
		v138 = int32(0)
		goto L36
	}
L36:
	;
	v142 = v117 + int32(1) + v138 + int32(1)
	goto L34
L37:
	;
	v137 = *(*int32)(unsafe.Add(mBase, uint32(v117+int32(-16))))
	v138 = v137
	goto L36
L38:
	;
	v134 = *(*int32)(unsafe.Add(mBase, uint32(v117+int32(-8))))
	v138 = v134
	goto L36
L39:
	;
	v131 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v117+int32(-4)))))
	v138 = v131
	goto L36
L40:
	;
	v128 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v117+int32(-2)))))
	v138 = v128
	goto L36
L41:
	;
	v138 = int32(base.Ui32(v121) >> (uint(int32(3)) % 32))
	goto L36
L42:
	;
	v161 = v142 + v157
	goto L28
L43:
	;
	goto L42
L44:
	;
	v157 = *(*int32)(unsafe.Add(mBase, _consts[308]))
	goto L43
L45:
	;
	v297 = v295
	goto L3
L46:
	;
	v231 = int32(0)
	if base.Ui64(v165) < base.Ui64(int64(10)) {
		v288 = v231
		goto L80
	} else {
		goto L81
	}
L47:
	;
	v170 = int32(0)
	v172 = int64(0) - v165
	if base.Ui64(v172) < base.Ui64(int64(10)) {
		v222 = v170
		goto L49
	} else {
		goto L50
	}
L48:
	;
	v295 = v226 + v227 + int32(1)
	goto L45
L49:
	;
	v226 = v222
	v227 = int32(1)
	goto L48
L50:
	;
	v175 = v172
	v176 = v170
	goto L51
L51:
	;
	if base.Ui64(int64(100)) <= base.Ui64(v175) {
		goto L53
	} else {
		goto L54
	}
L52:
	;
	v222 = v216
	goto L49
L53:
	;
	if base.Ui64(int64(1000)) <= base.Ui64(v175) {
		goto L55
	} else {
		goto L56
	}
L54:
	;
	v226 = v176
	v227 = int32(2)
	goto L48
L55:
	;
	if base.Ui64(int64(999999999999)) < base.Ui64(v175) {
		goto L57
	} else {
		goto L58
	}
L56:
	;
	v226 = v176
	v227 = int32(3)
	goto L48
L57:
	;
	v216 = v176 + int32(12)
	v220 = base.I64_div_u_s(v175, int64(1000000000000))
	if base.Ui64(int64(9999999999999)) < base.Ui64(v175) {
		v175 = v220
		v176 = v216
		goto L51
	} else {
		goto L79
	}
L58:
	;
	if base.Ui64(int64(99999999)) < base.Ui64(v175) {
		goto L59
	} else {
		goto L60
	}
L59:
	;
	if base.Ui64(int64(9999999999)) < base.Ui64(v175) {
		goto L71
	} else {
		goto L72
	}
L60:
	;
	if base.Ui64(int64(999999)) < base.Ui64(v175) {
		goto L61
	} else {
		goto L62
	}
L61:
	;
	if base.Ui64(int64(9999999)) < base.Ui64(v175) {
		goto L68
	} else {
		goto L69
	}
L62:
	;
	if base.Ui64(int64(10000)) <= base.Ui64(v175) {
		goto L63
	} else {
		goto L64
	}
L63:
	;
	if base.Ui64(int64(99999)) < base.Ui64(v175) {
		goto L65
	} else {
		goto L66
	}
L64:
	;
	v226 = v176
	v227 = int32(4)
	goto L48
L65:
	;
	v197 = int32(6)
	goto L67
L66:
	;
	v197 = int32(5)
	goto L67
L67:
	;
	v226 = v176
	v227 = v197
	goto L48
L68:
	;
	v202 = int32(8)
	goto L70
L69:
	;
	v202 = int32(7)
	goto L70
L70:
	;
	v226 = v176
	v227 = v202
	goto L48
L71:
	;
	if base.Ui64(int64(99999999999)) < base.Ui64(v175) {
		goto L76
	} else {
		goto L77
	}
L72:
	;
	if base.Ui64(int64(999999999)) < base.Ui64(v175) {
		goto L73
	} else {
		goto L74
	}
L73:
	;
	v209 = int32(10)
	goto L75
L74:
	;
	v209 = int32(9)
	goto L75
L75:
	;
	v226 = v176
	v227 = v209
	goto L48
L76:
	;
	v214 = int32(12)
	goto L78
L77:
	;
	v214 = int32(11)
	goto L78
L78:
	;
	v226 = v176
	v227 = v214
	goto L48
L79:
	;
	goto L52
L80:
	;
	v295 = int32(1) + v288
	goto L45
L81:
	;
	v234 = v165
	v235 = v231
	goto L82
L82:
	;
	if base.Ui64(int64(100)) <= base.Ui64(v234) {
		goto L84
	} else {
		goto L85
	}
L83:
	;
	v288 = v282
	goto L80
L84:
	;
	if base.Ui64(int64(1000)) <= base.Ui64(v234) {
		goto L86
	} else {
		goto L87
	}
L85:
	;
	v295 = int32(2) + v235
	goto L45
L86:
	;
	if base.Ui64(int64(999999999999)) < base.Ui64(v234) {
		goto L88
	} else {
		goto L89
	}
L87:
	;
	v295 = int32(3) + v235
	goto L45
L88:
	;
	v282 = v235 + int32(12)
	v286 = base.I64_div_u_s(v234, int64(1000000000000))
	if base.Ui64(int64(9999999999999)) < base.Ui64(v234) {
		v234 = v286
		v235 = v282
		goto L82
	} else {
		goto L110
	}
L89:
	;
	if base.Ui64(int64(99999999)) < base.Ui64(v234) {
		goto L90
	} else {
		goto L91
	}
L90:
	;
	if base.Ui64(int64(9999999999)) < base.Ui64(v234) {
		goto L102
	} else {
		goto L103
	}
L91:
	;
	if base.Ui64(int64(999999)) < base.Ui64(v234) {
		goto L92
	} else {
		goto L93
	}
L92:
	;
	if base.Ui64(int64(9999999)) < base.Ui64(v234) {
		goto L99
	} else {
		goto L100
	}
L93:
	;
	if base.Ui64(int64(10000)) <= base.Ui64(v234) {
		goto L94
	} else {
		goto L95
	}
L94:
	;
	if base.Ui64(int64(99999)) < base.Ui64(v234) {
		goto L96
	} else {
		goto L97
	}
L95:
	;
	v295 = int32(4) + v235
	goto L45
L96:
	;
	v259 = int32(6)
	goto L98
L97:
	;
	v259 = int32(5)
	goto L98
L98:
	;
	v295 = v259 + v235
	goto L45
L99:
	;
	v265 = int32(8)
	goto L101
L100:
	;
	v265 = int32(7)
	goto L101
L101:
	;
	v295 = v265 + v235
	goto L45
L102:
	;
	if base.Ui64(int64(99999999999)) < base.Ui64(v234) {
		goto L107
	} else {
		goto L108
	}
L103:
	;
	if base.Ui64(int64(999999999)) < base.Ui64(v234) {
		goto L104
	} else {
		goto L105
	}
L104:
	;
	v273 = int32(10)
	goto L106
L105:
	;
	v273 = int32(9)
	goto L106
L106:
	;
	v295 = v273 + v235
	goto L45
L107:
	;
	v279 = int32(12)
	goto L109
L108:
	;
	v279 = int32(11)
	goto L109
L109:
	;
	v295 = v279 + v235
	goto L45
L110:
	;
	goto L83
L111:
	;
	return int32(0)
L112:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
}
