package p0

import (
	base "github.com/shibukawa/valkeymem/internal/aot/vkaot/base"
	"unsafe"
)

func F_IOThreadMain(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v41 int32
	_ = v41
	var v48 int32
	_ = v48
	var v52 int32
	_ = v52
	var v58 int32
	_ = v58
	var v67 int64
	_ = v67
	var v68 int32
	_ = v68
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v77 int64
	_ = v77
	var v80 int64
	_ = v80
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v104 int32
	_ = v104
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v123 int32
	_ = v123
	var v129 int32
	_ = v129
	var v135 int32
	_ = v135
	var v137 int32
	_ = v137
	var v145 int32
	_ = v145
	var v148 int32
	_ = v148
	var v157 int32
	_ = v157
	var v160 int32
	_ = v160
	var v169 int32
	_ = v169
	var v172 int32
	_ = v172
	var v181 int32
	_ = v181
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v186 int32
	_ = v186
	var v192 int32
	_ = v192
	var v206 int32
	_ = v206
	var v210 int32
	_ = v210
	var v214 int32
	_ = v214
	var v222 int32
	_ = v222
	var v224 int32
	_ = v224
	var v227 int32
	_ = v227
	var v248 int32
	_ = v248
	var v264 int32
	_ = v264
	var v266 int32
	_ = v266
	var v282 int32
	_ = v282
	var v291 int32
	_ = v291
	var v293 int32
	_ = v293
	var v296 int32
	_ = v296
	var v299 int32
	_ = v299
	var v316 int32
	_ = v316
	var v319 int32
	_ = v319
	var v323 int32
	_ = v323
	var v327 int64
	_ = v327
	var v333 int32
	_ = v333
	var v334 int32
	_ = v334
	var v335 int32
	_ = v335
	var v347 int32
	_ = v347
	var v353 int32
	_ = v353
	var v355 int32
	_ = v355
	var v371 int32
	_ = v371
	var v373 int32
	_ = v373
	var v375 int32
	_ = v375
	var v376 int32
	_ = v376
	var v386 int32
	_ = v386
	var v387 int32
	_ = v387
	var v389 int32
	_ = v389
	var v393 int32
	_ = v393
	var v395 int32
	_ = v395
	var v397 int32
	_ = v397
	var v401 int32
	_ = v401
	var v402 int32
	_ = v402
	var v404 int32
	_ = v404
	var v405 int32
	_ = v405
	var v406 int32
	_ = v406
	var v412 int32
	_ = v412
	var v418 int32
	_ = v418
	var v424 int32
	_ = v424
	var v426 int32
	_ = v426
	var v434 int32
	_ = v434
	var v437 int32
	_ = v437
	var v446 int32
	_ = v446
	var v449 int32
	_ = v449
	var v458 int32
	_ = v458
	var v461 int32
	_ = v461
	var v470 int32
	_ = v470
	var v472 int32
	_ = v472
	var v473 int32
	_ = v473
	var v475 int32
	_ = v475
	var v481 int32
	_ = v481
	var v495 int32
	_ = v495
	var v499 int32
	_ = v499
	var v503 int32
	_ = v503
	var v511 int32
	_ = v511
	var v513 int32
	_ = v513
	var v516 int32
	_ = v516
	var v537 int32
	_ = v537
	var v551 int32
	_ = v551
	var v558 int32
	_ = v558
	var v563 int32
	_ = v563
	var v565 int32
	_ = v565
	var v569 int32
	_ = v569
	var v570 int32
	_ = v570
	var v576 int32
	_ = v576
	var v577 int32
	_ = v577
	var v579 int32
	_ = v579
	var v581 int32
	_ = v581
	var v583 int32
	_ = v583
	var v586 int32
	_ = v586
	var v589 int32
	_ = v589
	var v592 int32
	_ = v592
	var v594 int32
	_ = v594
	var v598 int32
	_ = v598
	var v600 int32
	_ = v600
	var v602 int32
	_ = v602
	var v604 int32
	_ = v604
	var v605 int32
	_ = v605
	var v607 int32
	_ = v607
	var v608 int32
	_ = v608
	var v609 int32
	_ = v609
	var v610 int32
	_ = v610
	var v611 int32
	_ = v611
	var v613 int32
	_ = v613
	var v614 int32
	_ = v614
	var v619 int32
	_ = v619
	var v621 int32
	_ = v621
	var v625 int32
	_ = v625
	var v626 int32
	_ = v626
	var v628 int32
	_ = v628
	var v630 int32
	_ = v630
	var v631 int32
	_ = v631
	var v633 int32
	_ = v633
	var v634 int32
	_ = v634
	var v635 int32
	_ = v635
	var v636 int64
	_ = v636
	var v642 int32
	_ = v642
	var v643 int32
	_ = v643
	var v644 int32
	_ = v644
	var v654 int32
	_ = v654
	var v657 int32
	_ = v657
	var v662 int32
	_ = v662
	var v666 int32
	_ = v666
	var v668 int32
	_ = v668
	var v672 int32
	_ = v672
	var v673 int32
	_ = v673
	var v678 int32
	_ = v678
	v15 = m.G0
	v17 = v15 - int32(240)
	m.G0 = v17
	*(*int32)(unsafe.Add(mBase, uint32(v17)+32)) = l0
	v22 = int32(32)
	v26 = F_snprintf(m, v17+int32(192), v22, int32(_a662), v17+v22)
	mBase = m.M
	v29 = m.ExcPending
	if v29 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	goto L3
L3:
	;
	F_initSharedQueryBuf(m)
	mBase = m.M
	v33 = m.ExcPending
	if v33 != 0 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v35 = v17 + int32(180)
	*(*int32)(unsafe.Add(mBase, uint32(v35)+4)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v35))) = int32(541)
	F_dummy_4(m, v35)
	mBase = m.M
	goto L5
L5:
	;
	v41 = int32(0)
	*(*int32)(unsafe.Add(mBase, _consts[299])) = l0
	v48 = l0 * int32(192)
	v52 = l0 << (uint(int32(3)) % 32)
	v58 = v17 + int32(232)
	v67 = int64(0)
	v68 = v41
	goto L6
L6:
	;
	goto L8
L8:
	;
	v75 = int32(0)
	v76 = *(*int32)(unsafe.Add(mBase, _consts[271]))
	v77 = m.T0[v76].(func(*base.Module) int64)(m)
	mBase = m.M
	if v68 == v75 {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	v84 = int32(0)
	v86 = v17 + int32(48)
	v87 = int32(32)
	v97 = *(*int32)(unsafe.Add(mBase, uint32(v48)+uint32(_consts[303])))
	v98 = *(*int32)(unsafe.Add(mBase, uint32(v48)+uint32(_consts[304])))
	if v97 != v98 {
		v104 = v98
		goto L14
	} else {
		goto L15
	}
L10:
	;
	v80 = *(*int64)(unsafe.Add(mBase, uint32(v52)+uint32(_consts[302])))
	*(*int64)(unsafe.Add(mBase, uint32(v52)+uint32(_consts[302]))) = v80 + (v77 - v67)
	goto L9
L11:
	;
	v558 = int32(_a476)
	v563 = *(*int32)(unsafe.Add(mBase, _consts[305]))
	v565 = v563
	goto L74
L12:
	;
	if v248 == int32(0) {
		v551 = v84
		goto L11
	} else {
		goto L31
	}
L13:
	;
	goto L12
L14:
	;
	v106 = v104 - v97
	if base.Ui32(v87) < base.Ui32(v106) {
		goto L18
	} else {
		goto L19
	}
L15:
	;
	v100 = *(*int32)(unsafe.Add(mBase, uint32(v48)+uint32(_consts[306])))
	*(*int32)(unsafe.Add(mBase, uint32(v48)+uint32(_consts[304]))) = v100
	if v97 == v100 {
		v248 = int32(0)
		goto L13
	} else {
		goto L16
	}
L16:
	;
	v104 = v100
	goto L14
L17:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v48)+uint32(_consts[303]))) = v108 + v97
	v248 = v108
	goto L13
L18:
	;
	v108 = v87
	goto L20
L19:
	;
	v108 = v106
	goto L20
L20:
	;
	if v108 == int32(0) {
		goto L17
	} else {
		goto L21
	}
L21:
	;
	v112 = v108 & int32(3)
	v113 = *(*int32)(unsafe.Add(mBase, uint32(v48)+uint32(_consts[307])))
	v115 = v113 + int32(-1)
	v116 = *(*int32)(unsafe.Add(mBase, uint32(v48)+uint32(_consts[308])))
	v117 = int32(0)
	if base.Ui32(v108) < base.Ui32(int32(4)) {
		v192 = v117
		goto L22
	} else {
		goto L23
	}
L22:
	;
	if v112 == int32(0) {
		goto L17
	} else {
		goto L27
	}
L23:
	;
	v123 = int32(0)
	v129 = v123
	v135 = v123
	goto L24
L24:
	;
	v137 = int32(2)
	v145 = *(*int32)(unsafe.Add(mBase, uint32(v116+v115&(v129+v97)<<(uint(v137)%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v86+v129<<(uint(v137)%32)))) = v145
	v148 = v129 | int32(1)
	v157 = *(*int32)(unsafe.Add(mBase, uint32(v116+v115&(v148+v97)<<(uint(v137)%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v86+v148<<(uint(v137)%32)))) = v157
	v160 = v129 | v137
	v169 = *(*int32)(unsafe.Add(mBase, uint32(v116+v115&(v160+v97)<<(uint(v137)%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v86+v160<<(uint(v137)%32)))) = v169
	v172 = v129 | int32(3)
	v181 = *(*int32)(unsafe.Add(mBase, uint32(v116+v115&(v172+v97)<<(uint(v137)%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v86+v172<<(uint(v137)%32)))) = v181
	v183 = int32(4)
	v184 = v129 + v183
	v186 = v135 + v183
	if v186 != v108&int32(-4) {
		v129 = v184
		v135 = v186
		goto L24
	} else {
		goto L26
	}
L25:
	;
	v192 = v184
	goto L22
L26:
	;
	goto L25
L27:
	;
	v206 = v192
	v210 = v117
	goto L28
L28:
	;
	v214 = int32(2)
	v222 = *(*int32)(unsafe.Add(mBase, uint32(v116+v115&(v206+v97)<<(uint(v214)%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v86+v206<<(uint(v214)%32)))) = v222
	v224 = int32(1)
	v227 = v210 + v224
	if v227 != v112 {
		v206 = v206 + v224
		v210 = v227
		goto L28
	} else {
		goto L30
	}
L29:
	;
	goto L17
L30:
	;
	goto L29
L31:
	;
	v264 = v84
	v266 = v248
	goto L32
L32:
	;
	v282 = int32(0)
	goto L34
L33:
	;
	v551 = v373
	goto L11
L34:
	;
	v291 = *(*int32)(unsafe.Add(mBase, uint32(v17+int32(48)+v282<<(uint(int32(2))%32))))
	v293 = v291 & int32(-8)
	v296 = v291 & int32(7)
	switch v296 + int32(-2) {
	case 0:
		goto L40
	default:
		goto L38
	case 2:
		goto L39
	}
L35:
	;
	v373 = v266 + v264
	v375 = v17 + int32(48)
	v376 = int32(32)
	v386 = *(*int32)(unsafe.Add(mBase, uint32(v48)+uint32(_consts[303])))
	v387 = *(*int32)(unsafe.Add(mBase, uint32(v48)+uint32(_consts[304])))
	if v386 != v387 {
		v393 = v387
		goto L54
	} else {
		goto L55
	}
L36:
	;
	v371 = v282 + int32(1)
	if v371 != v266 {
		v282 = v371
		goto L34
	} else {
		goto L51
	}
L37:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v316)+4)) = v319 | int32(8)
	F_decrRefCount(m, v316)
	mBase = m.M
	v353 = m.ExcPending
	if v353 != 0 {
		goto L1
	} else {
		goto L49
	}
L38:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+16)) = v296
	F__serverPanic_1(m, int32(_a663), int32(327), int32(_a664), v17+int32(16))
	mBase = m.M
	v347 = m.ExcPending
	if v347 != 0 {
		goto L1
	} else {
		goto L48
	}
L39:
	;
	v327 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v58))) = v327
	*(*int64)(unsafe.Add(mBase, uint32(v17)+224)) = v327
	v333 = F_aePoll(m, v293, v17+int32(224))
	mBase = m.M
	v334 = m.ExcPending
	if v334 != 0 {
		goto L1
	} else {
		goto L47
	}
L40:
	;
	v299 = int32(0)
	goto L41
L41:
	;
	v316 = *(*int32)(unsafe.Add(mBase, uint32(v293+v299<<(uint(int32(2))%32))))
	if v316 == int32(0) {
		goto L43
	} else {
		goto L44
	}
L43:
	;
	v299 = v299 + int32(1)
	goto L41
L44:
	;
	v319 = *(*int32)(unsafe.Add(mBase, uint32(v316)+4))
	if base.Ui32(v319) < base.Ui32(int32(8)) {
		goto L37
	} else {
		goto L45
	}
L45:
	;
	F_decrRefCount(m, v316)
	mBase = m.M
	v323 = m.ExcPending
	if v323 != 0 {
		goto L1
	} else {
		goto L46
	}
L46:
	;
	goto L43
L47:
	;
	v335 = int32(_a69)
	*(*int32)(unsafe.Add(mBase, _consts[309])) = int32(2)
	*(*int32)(unsafe.Add(mBase, _consts[310])) = v333
	goto L36
L48:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L49:
	;
	F_valkey_free(m, v293)
	mBase = m.M
	v355 = m.ExcPending
	if v355 != 0 {
		goto L1
	} else {
		goto L50
	}
L50:
	;
	goto L36
L51:
	;
	goto L35
L52:
	;
	if v537 != 0 {
		v264 = v373
		v266 = v537
		goto L32
	} else {
		goto L71
	}
L53:
	;
	goto L52
L54:
	;
	v395 = v393 - v386
	if base.Ui32(v376) < base.Ui32(v395) {
		goto L58
	} else {
		goto L59
	}
L55:
	;
	v389 = *(*int32)(unsafe.Add(mBase, uint32(v48)+uint32(_consts[306])))
	*(*int32)(unsafe.Add(mBase, uint32(v48)+uint32(_consts[304]))) = v389
	if v386 == v389 {
		v537 = int32(0)
		goto L53
	} else {
		goto L56
	}
L56:
	;
	v393 = v389
	goto L54
L57:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v48)+uint32(_consts[303]))) = v397 + v386
	v537 = v397
	goto L53
L58:
	;
	v397 = v376
	goto L60
L59:
	;
	v397 = v395
	goto L60
L60:
	;
	if v397 == int32(0) {
		goto L57
	} else {
		goto L61
	}
L61:
	;
	v401 = v397 & int32(3)
	v402 = *(*int32)(unsafe.Add(mBase, uint32(v48)+uint32(_consts[307])))
	v404 = v402 + int32(-1)
	v405 = *(*int32)(unsafe.Add(mBase, uint32(v48)+uint32(_consts[308])))
	v406 = int32(0)
	if base.Ui32(v397) < base.Ui32(int32(4)) {
		v481 = v406
		goto L62
	} else {
		goto L63
	}
L62:
	;
	if v401 == int32(0) {
		goto L57
	} else {
		goto L67
	}
L63:
	;
	v412 = int32(0)
	v418 = v412
	v424 = v412
	goto L64
L64:
	;
	v426 = int32(2)
	v434 = *(*int32)(unsafe.Add(mBase, uint32(v405+v404&(v418+v386)<<(uint(v426)%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v375+v418<<(uint(v426)%32)))) = v434
	v437 = v418 | int32(1)
	v446 = *(*int32)(unsafe.Add(mBase, uint32(v405+v404&(v437+v386)<<(uint(v426)%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v375+v437<<(uint(v426)%32)))) = v446
	v449 = v418 | v426
	v458 = *(*int32)(unsafe.Add(mBase, uint32(v405+v404&(v449+v386)<<(uint(v426)%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v375+v449<<(uint(v426)%32)))) = v458
	v461 = v418 | int32(3)
	v470 = *(*int32)(unsafe.Add(mBase, uint32(v405+v404&(v461+v386)<<(uint(v426)%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v375+v461<<(uint(v426)%32)))) = v470
	v472 = int32(4)
	v473 = v418 + v472
	v475 = v424 + v472
	if v475 != v397&int32(-4) {
		v418 = v473
		v424 = v475
		goto L64
	} else {
		goto L66
	}
L65:
	;
	v481 = v473
	goto L62
L66:
	;
	goto L65
L67:
	;
	v495 = v481
	v499 = v406
	goto L68
L68:
	;
	v503 = int32(2)
	v511 = *(*int32)(unsafe.Add(mBase, uint32(v405+v404&(v495+v386)<<(uint(v503)%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v375+v495<<(uint(v503)%32)))) = v511
	v513 = int32(1)
	v516 = v499 + v513
	if v516 != v401 {
		v495 = v495 + v513
		v499 = v516
		goto L68
	} else {
		goto L70
	}
L69:
	;
	goto L57
L70:
	;
	goto L69
L71:
	;
	goto L33
L72:
	;
	if v662 == int32(0) {
		goto L108
	} else {
		goto L109
	}
L73:
	;
	if v594 == int32(0) {
		v662 = v551
		goto L72
	} else {
		goto L84
	}
L74:
	;
	v569 = *(*int32)(unsafe.Add(mBase, _consts[311]))
	v570 = *(*int32)(unsafe.Add(mBase, _consts[312]))
	v576 = v569 + (v570+int32(-1))&v565<<(uint(int32(6))%32)
	v577 = *(*int32)(unsafe.Add(mBase, uint32(v576)))
	v579 = v565 + int32(1)
	if v577 != v579 {
		goto L76
	} else {
		goto L77
	}
L76:
	;
	if v579 <= v577 {
		goto L82
	} else {
		goto L83
	}
L77:
	;
	v581 = *(*int32)(unsafe.Add(mBase, _consts[305]))
	if v581 == v565 {
		goto L78
	} else {
		goto L79
	}
L78:
	;
	v583 = v577
	goto L80
L79:
	;
	v583 = v581
	goto L80
L80:
	;
	*(*int32)(unsafe.Add(mBase, _consts[305])) = v583
	if v581 != v565 {
		v565 = v581
		goto L74
	} else {
		goto L81
	}
L81:
	;
	v586 = *(*int32)(unsafe.Add(mBase, _consts[312]))
	*(*int32)(unsafe.Add(mBase, uint32(v576))) = v586 + v565
	v589 = *(*int32)(unsafe.Add(mBase, uint32(v576)+4))
	v594 = v589
	goto L73
L82:
	;
	v592 = *(*int32)(unsafe.Add(mBase, _consts[305]))
	v565 = v592
	goto L74
L83:
	;
	v594 = int32(0)
	goto L73
L84:
	;
	v598 = v594 & int32(-8)
	v600 = v594 & int32(7)
	switch v600 {
	case 0:
		goto L86
	case 1:
		goto L91
	default:
		goto L87
	case 3:
		goto L90
	case 4:
		goto L88
	case 5:
		goto L89
	}
L85:
	;
	v662 = v551 + int32(1)
	goto L72
L86:
	;
	F_ioThreadReadQueryFromClient(m, v598)
	mBase = m.M
	v657 = m.ExcPending
	if v657 != 0 {
		goto L1
	} else {
		goto L107
	}
L87:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17))) = v600
	F__serverPanic_1(m, int32(_a663), int32(358), int32(_a665), v17)
	mBase = m.M
	v654 = m.ExcPending
	if v654 != 0 {
		goto L1
	} else {
		goto L106
	}
L88:
	;
	v636 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v58))) = v636
	*(*int64)(unsafe.Add(mBase, uint32(v17)+224)) = v636
	v642 = F_aePoll(m, v598, v17+int32(224))
	mBase = m.M
	v643 = m.ExcPending
	if v643 != 0 {
		goto L1
	} else {
		goto L105
	}
L89:
	;
	v605 = *(*int32)(unsafe.Add(mBase, uint32(v598)+8))
	v607 = *(*int32)(unsafe.Add(mBase, uint32(v605)))
	v608 = *(*int32)(unsafe.Add(mBase, uint32(v607)+64))
	v609 = m.T0[v608].(func(*base.Module, int32, int32) int32)(m, v605, int32(0))
	mBase = m.M
	v610 = m.ExcPending
	if v610 != 0 {
		goto L1
	} else {
		goto L94
	}
L90:
	;
	F_decrRefCount(m, v598)
	mBase = m.M
	v604 = m.ExcPending
	if v604 != 0 {
		goto L1
	} else {
		goto L93
	}
L91:
	;
	F_ioThreadWriteToClient(m, v598)
	mBase = m.M
	v602 = m.ExcPending
	if v602 != 0 {
		goto L1
	} else {
		goto L92
	}
L92:
	;
	goto L85
L93:
	;
	goto L85
L94:
	;
	v611 = int32(2)
	*(*uint8)(unsafe.Add(mBase, uint32(v598)+222)) = uint8(v611)
	v613 = int32(0)
	v614 = *(*int32)(unsafe.Add(mBase, _consts[300]))
	if v614 == v613 {
		goto L96
	} else {
		goto L97
	}
L95:
	;
	v634 = F_listAddNodeTail(m, v633, v598)
	mBase = m.M
	v635 = m.ExcPending
	if v635 != 0 {
		goto L1
	} else {
		goto L104
	}
L96:
	;
	v625 = F_mpscEnqueue(m, int32(_a660), v598, int32(_a661))
	mBase = m.M
	v626 = m.ExcPending
	if v626 != 0 {
		goto L1
	} else {
		goto L100
	}
L97:
	;
	F_flushPendingIOResponses(m, int32(0))
	mBase = m.M
	v619 = m.ExcPending
	if v619 != 0 {
		goto L1
	} else {
		goto L98
	}
L98:
	;
	v621 = *(*int32)(unsafe.Add(mBase, _consts[300]))
	if v621 != 0 {
		v633 = v621
		goto L95
	} else {
		goto L99
	}
L99:
	;
	goto L96
L100:
	;
	if v625 != 0 {
		goto L85
	} else {
		goto L101
	}
L101:
	;
	v628 = *(*int32)(unsafe.Add(mBase, _consts[300]))
	if v628 != 0 {
		v633 = v628
		goto L95
	} else {
		goto L102
	}
L102:
	;
	v630 = F_listCreate(m)
	mBase = m.M
	v631 = m.ExcPending
	if v631 != 0 {
		goto L1
	} else {
		goto L103
	}
L103:
	;
	*(*int32)(unsafe.Add(mBase, _consts[300])) = v630
	v633 = v630
	goto L95
L104:
	;
	goto L85
L105:
	;
	v644 = int32(_a69)
	*(*int32)(unsafe.Add(mBase, _consts[309])) = int32(2)
	*(*int32)(unsafe.Add(mBase, _consts[310])) = v642
	goto L85
L106:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L107:
	;
	goto L85
L108:
	;
	v672 = int32(0)
	v673 = *(*int32)(unsafe.Add(mBase, _consts[300]))
	if v673 == v672 {
		goto L110
	} else {
		goto L111
	}
L109:
	;
	v666 = int32(0)
	v668 = *(*int32)(unsafe.Add(mBase, _consts[313]))
	*(*int32)(unsafe.Add(mBase, _consts[313])) = v668 + v662
	v67 = v77
	v68 = int32(1)
	goto L6
L110:
	;
	goto L113
L111:
	;
	F_flushPendingIOResponses(m, int32(0))
	mBase = m.M
	v678 = m.ExcPending
	if v678 != 0 {
		goto L1
	} else {
		goto L112
	}
L112:
	;
	v67 = v77
	v68 = v662
	goto L6
L113:
	;
	goto L114
L114:
	;
	v67 = v77
	v68 = v662
	goto L6
}
func F_getIOThreadActiveTimeMicroseconds(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v6 int64
	_ = v6
	v6 = *(*int64)(unsafe.Add(mBase, uint32(l0<<(uint(int32(3))%32))+uint32(_consts[302])))
	return v6
}
func F_ioThreadWriteToClient(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v20 int32
	_ = v20
	var v25 int32
	_ = v25
	v2 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+223)))
	if v2 != int32(1) {
		F__serverAssert(m, int32(_a819), int32(_a774), int32(6721))
		mBase = m.M
		v25 = m.ExcPending
		if v25 != 0 {
			return
		} else {
			F_abort(m)
			mBase = m.M
			base.Wasm_trap_unreachable()
			for {
			}
		}
	} else {
		v5 = int32(0)
		*(*int32)(unsafe.Add(mBase, uint32(l0)+280)) = v5
		v7 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+220)))
		if v7&int32(2) == v5 {
			v14 = F__writeToClient(m, l0)
			mBase = m.M
			v15 = m.ExcPending
			if v15 != 0 {
				return
			} else {
				v16 = int32(2)
				*(*uint8)(unsafe.Add(mBase, uint32(l0)+223)) = uint8(v16)
				F_sendToMainThread(m, l0, int32(1))
				mBase = m.M
				v20 = m.ExcPending
				if v20 != 0 {
					return
				} else {
					return
				}
			}
		} else {
			F_writeToReplica(m, l0)
			mBase = m.M
			v13 = m.ExcPending
			if v13 != 0 {
				return
			} else {
				v16 = int32(2)
				*(*uint8)(unsafe.Add(mBase, uint32(l0)+223)) = uint8(v16)
				F_sendToMainThread(m, l0, int32(1))
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
