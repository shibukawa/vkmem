package p1

import (
	base "github.com/shibukawa/vkmem/internal/aot/vkaot/base"
	"unsafe"
)

func F_newkey(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v55 int32
	_ = v55
	var v58 float64
	_ = v58
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v63 int64
	_ = v63
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v75 int32
	_ = v75
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v111 int32
	_ = v111
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
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v132 int32
	_ = v132
	var v136 int32
	_ = v136
	var v139 int32
	_ = v139
	var v149 int32
	_ = v149
	var v165 int32
	_ = v165
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v174 int32
	_ = v174
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v209 int32
	_ = v209
	var v212 int32
	_ = v212
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v232 int32
	_ = v232
	var v240 int32
	_ = v240
	var v244 int32
	_ = v244
	var v257 int32
	_ = v257
	var v258 int32
	_ = v258
	var v259 int32
	_ = v259
	var v261 int32
	_ = v261
	var v262 int32
	_ = v262
	var v266 int32
	_ = v266
	var v271 int32
	_ = v271
	var v276 int32
	_ = v276
	var v279 int32
	_ = v279
	var v281 int32
	_ = v281
	var v283 int32
	_ = v283
	var v291 int32
	_ = v291
	var v295 int32
	_ = v295
	var v317 int32
	_ = v317
	var v318 int32
	_ = v318
	var v321 int32
	_ = v321
	var v338 int32
	_ = v338
	var v341 int32
	_ = v341
	var v342 int32
	_ = v342
	var v345 int32
	_ = v345
	var v353 int32
	_ = v353
	var v379 int32
	_ = v379
	var v383 int32
	_ = v383
	var v399 int32
	_ = v399
	var v400 int32
	_ = v400
	var v403 int32
	_ = v403
	var v405 int32
	_ = v405
	var v407 int32
	_ = v407
	var v421 int32
	_ = v421
	var v434 int32
	_ = v434
	var v435 int32
	_ = v435
	var v439 int32
	_ = v439
	var v447 int32
	_ = v447
	var v448 int32
	_ = v448
	var v450 int32
	_ = v450
	var v465 int32
	_ = v465
	var v468 int32
	_ = v468
	var v469 int32
	_ = v469
	var v472 int32
	_ = v472
	var v473 int32
	_ = v473
	var v476 float64
	_ = v476
	var v482 int32
	_ = v482
	var v484 int32
	_ = v484
	var v488 int32
	_ = v488
	var v495 int32
	_ = v495
	var v498 int32
	_ = v498
	var v499 int32
	_ = v499
	var v502 int32
	_ = v502
	var v503 int32
	_ = v503
	var v507 int32
	_ = v507
	var v509 int32
	_ = v509
	var v510 int32
	_ = v510
	var v512 int32
	_ = v512
	var v516 int32
	_ = v516
	var v520 int32
	_ = v520
	var v521 int32
	_ = v521
	var v527 int32
	_ = v527
	var v533 int32
	_ = v533
	var v534 int32
	_ = v534
	var v539 int32
	_ = v539
	var v541 int32
	_ = v541
	var v544 float64
	_ = v544
	var v550 int32
	_ = v550
	var v552 int32
	_ = v552
	var v558 int32
	_ = v558
	var v565 int32
	_ = v565
	var v568 int32
	_ = v568
	var v569 int32
	_ = v569
	var v572 int32
	_ = v572
	var v573 int32
	_ = v573
	var v577 int32
	_ = v577
	var v579 int32
	_ = v579
	var v580 int32
	_ = v580
	var v582 int32
	_ = v582
	var v586 int32
	_ = v586
	var v590 int32
	_ = v590
	var v591 int32
	_ = v591
	var v597 int32
	_ = v597
	var v598 int32
	_ = v598
	var v599 int32
	_ = v599
	var v603 int32
	_ = v603
	var v612 int32
	_ = v612
	var v613 int32
	_ = v613
	var v616 int32
	_ = v616
	var v623 int32
	_ = v623
	var v624 int32
	_ = v624
	var v625 int32
	_ = v625
	var v633 int32
	_ = v633
	var v635 int32
	_ = v635
	var v636 int32
	_ = v636
	var v638 int32
	_ = v638
	var v639 int32
	_ = v639
	var v640 int32
	_ = v640
	var v641 int32
	_ = v641
	var v643 int32
	_ = v643
	var v646 int32
	_ = v646
	var v668 int32
	_ = v668
	var v669 int32
	_ = v669
	var v681 int32
	_ = v681
	var v682 int32
	_ = v682
	var v683 int32
	_ = v683
	var v685 int32
	_ = v685
	var v687 int32
	_ = v687
	var v688 int32
	_ = v688
	var v691 int32
	_ = v691
	var v692 float64
	_ = v692
	var v697 int32
	_ = v697
	var v701 int32
	_ = v701
	var v702 float64
	_ = v702
	var v705 int64
	_ = v705
	var v710 int32
	_ = v710
	var v711 int32
	_ = v711
	var v717 int32
	_ = v717
	var v721 int32
	_ = v721
	var v722 int32
	_ = v722
	var v723 int32
	_ = v723
	var v724 int32
	_ = v724
	var v732 int32
	_ = v732
	var v733 int32
	_ = v733
	var v734 int32
	_ = v734
	var v742 int32
	_ = v742
	var v743 int32
	_ = v743
	var v744 int32
	_ = v744
	var v750 int32
	_ = v750
	var v754 int32
	_ = v754
	var v755 int32
	_ = v755
	var v756 int32
	_ = v756
	var v762 int32
	_ = v762
	var v767 int32
	_ = v767
	var v776 int32
	_ = v776
	var v794 int32
	_ = v794
	var v801 int64
	_ = v801
	var v807 int64
	_ = v807
	var v812 int32
	_ = v812
	var v813 int64
	_ = v813
	var v815 int64
	_ = v815
	var v817 int32
	_ = v817
	var v823 int32
	_ = v823
	var v833 int32
	_ = v833
	var v850 int64
	_ = v850
	var v852 int32
	_ = v852
	var v856 int32
	_ = v856
	var v857 int32
	_ = v857
	var v862 int32
	_ = v862
	var v867 int32
	_ = v867
	var v868 int32
	_ = v868
	var v870 int32
	_ = v870
	var v872 int32
	_ = v872
	var v882 int32
	_ = v882
	v25 = m.G0
	v27 = v25 - int32(112)
	m.G0 = v27
	v30 = v27 | int32(4)
	goto L8
L1:
	;
	m.G0 = v27 + int32(112)
	return v882
L2:
	;
	v850 = *(*int64)(unsafe.Add(mBase, uint32(l2)))
	*(*int64)(unsafe.Add(mBase, uint32(v833)+16)) = v850
	v852 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v833)+24)) = v852
	if v852 < int32(4) {
		v882 = v833
		goto L1
	} else {
		goto L109
	}
L3:
	;
	if v767 == v130 {
		goto L104
	} else {
		goto L105
	}
L4:
	;
	v754 = *(*int32)(unsafe.Add(mBase, uint32(v130)+16))
	v755 = int32(-1)
	v756 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+12)))
	v762 = base.I32_rem_u_s(v754, v755<<(uint(v756)%32)^v755|int32(1))
	v767 = v129 + v762<<(uint(int32(5))%32)
	goto L3
L5:
	;
	v742 = *(*int32)(unsafe.Add(mBase, uint32(v130)+16))
	v743 = int32(-1)
	v744 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+12)))
	v750 = base.I32_rem_u_s(v742, v743<<(uint(v744)%32)^v743|int32(1))
	v767 = v129 + v750<<(uint(int32(5))%32)
	goto L3
L6:
	;
	v732 = *(*int32)(unsafe.Add(mBase, uint32(v130)+16))
	v733 = int32(-1)
	v734 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+12)))
	v767 = v129 + v732&(v733<<(uint(v734)%32)^v733)<<(uint(int32(5))%32)
	goto L3
L7:
	;
	v721 = *(*int32)(unsafe.Add(mBase, uint32(v130)+16))
	v722 = *(*int32)(unsafe.Add(mBase, uint32(v721)+8))
	v723 = int32(-1)
	v724 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+12)))
	v767 = v129 + v722&(v723<<(uint(v724)%32)^v723)<<(uint(int32(5))%32)
	goto L3
L8:
	;
	v55 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	switch v55 + int32(-1) {
	case 0:
		goto L13
	case 1:
		goto L12
	case 2:
		goto L15
	case 3:
		goto L14
	default:
		goto L11
	}
L9:
	;
	v702 = *(*float64)(unsafe.Add(mBase, uint32(v130)+16))
	if base.F64_eq(v702, float64(0)) != 0 {
		v767 = v129
		goto L3
	} else {
		goto L103
	}
L10:
	;
	v132 = m.G3
	if v130 == v132+int32(_a2685) {
		goto L18
	} else {
		goto L19
	}
L11:
	;
	v115 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
	v116 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	v117 = int32(-1)
	v118 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+12)))
	v124 = base.I32_rem_u_s(v116, v117<<(uint(v118)%32)^v117|int32(1))
	v129 = v115
	v130 = v115 + v124<<(uint(int32(5))%32)
	goto L10
L12:
	;
	v102 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
	v103 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	v104 = int32(-1)
	v105 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+12)))
	v111 = base.I32_rem_u_s(v103, v104<<(uint(v105)%32)^v104|int32(1))
	v129 = v102
	v130 = v102 + v111<<(uint(int32(5))%32)
	goto L10
L13:
	;
	v91 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
	v92 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	v93 = int32(-1)
	v94 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+12)))
	v129 = v91
	v130 = v91 + v92&(v93<<(uint(v94)%32)^v93)<<(uint(int32(5))%32)
	goto L10
L14:
	;
	v79 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
	v80 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	v81 = *(*int32)(unsafe.Add(mBase, uint32(v80)+8))
	v82 = int32(-1)
	v83 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+12)))
	v129 = v79
	v130 = v79 + v81&(v82<<(uint(v83)%32)^v82)<<(uint(int32(5))%32)
	goto L10
L15:
	;
	v58 = *(*float64)(unsafe.Add(mBase, uint32(l2)))
	if base.F64_ne(v58, float64(0)) != 0 {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	v62 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
	v63 = base.I64_reinterpret_f64(v58)
	v68 = int32(-1)
	v69 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+12)))
	v75 = base.I32_rem_u_s(base.I32_wrap_i64(int64(base.Ui64(v63)>>(uint(int64(32))%64))+v63), v68<<(uint(v69)%32)^v68|int32(1))
	v129 = v62
	v130 = v62 + v75<<(uint(int32(5))%32)
	goto L10
L17:
	;
	v61 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
	v129 = v61
	v130 = v61
	goto L10
L18:
	;
	v139 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
	v149 = v139
	goto L23
L19:
	;
	v136 = *(*int32)(unsafe.Add(mBase, uint32(v130)+8))
	if v136 == int32(0) {
		v833 = v130
		goto L2
	} else {
		goto L20
	}
L20:
	;
	goto L18
L21:
	;
	goto L9
L22:
	;
	v174 = int32(0)
	v178 = F__emscripten_memset_bulkmem(m, v27, base.I32_extend8_s(v174), int32(108))
	mBase = m.M
	goto L27
L23:
	;
	v165 = v149 + int32(-32)
	*(*int32)(unsafe.Add(mBase, uint32(l1)+28)) = v165
	if base.Ui32(v149) <= base.Ui32(v129) {
		goto L22
	} else {
		goto L25
	}
L24:
	;
	v171 = *(*int32)(unsafe.Add(mBase, uint32(v130)+24))
	switch v171 + int32(-1) {
	case 0:
		goto L6
	case 1:
		goto L5
	case 2:
		goto L21
	case 3:
		goto L7
	default:
		goto L4
	}
L25:
	;
	v170 = *(*int32)(unsafe.Add(mBase, uint32(v149+int32(-8))))
	if v170 != 0 {
		v149 = v165
		goto L23
	} else {
		goto L26
	}
L26:
	;
	goto L24
L27:
	;
	v179 = *(*int32)(unsafe.Add(mBase, uint32(l1)+36))
	v180 = int32(1)
	v193 = v180
	v194 = v174
	v197 = v180
	v198 = int32(0)
	goto L29
L28:
	;
	v434 = int32(-1)
	v435 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+12)))
	v439 = int32(0)
	v447 = v434<<(uint(v435)%32) ^ v434
	v448 = v439
	v450 = v439
	goto L48
L29:
	;
	if v197 <= v179 {
		v209 = v197
		goto L31
	} else {
		goto L32
	}
L30:
	;
	v421 = v405
	goto L28
L31:
	;
	if v209 < v193 {
		v379 = int32(0)
		v383 = v193
		goto L34
	} else {
		goto L35
	}
L32:
	;
	if v179 < v193 {
		v421 = v194
		goto L28
	} else {
		goto L33
	}
L33:
	;
	v209 = v179
	goto L31
L34:
	;
	v399 = v178 + v198<<(uint(int32(2))%32)
	v400 = *(*int32)(unsafe.Add(mBase, uint32(v399)))
	*(*int32)(unsafe.Add(mBase, uint32(v399))) = v400 + v379
	v403 = int32(1)
	v405 = v379 + v194
	v407 = v198 + v403
	if v407 != int32(27) {
		v193 = v383
		v194 = v405
		v197 = v197 << (uint(v403) % 32)
		v198 = v407
		goto L29
	} else {
		goto L47
	}
L35:
	;
	v212 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	v214 = v212 + int32(-8)
	v215 = v209 - v193
	v217 = v215 + int32(1)
	v218 = int32(3)
	v219 = v217 & v218
	if base.Ui32(v218) <= base.Ui32(v215) {
		goto L37
	} else {
		goto L38
	}
L36:
	;
	if v219 == int32(0) {
		v353 = v291
		goto L42
	} else {
		goto L43
	}
L37:
	;
	v232 = int32(0)
	v240 = v232
	v244 = v193
	v257 = v232
	goto L39
L38:
	;
	v291 = int32(0)
	v295 = v193
	goto L36
L39:
	;
	v258 = int32(4)
	v259 = v244 << (uint(v258) % 32)
	v261 = *(*int32)(unsafe.Add(mBase, uint32(v214+v259)))
	v262 = int32(0)
	v266 = *(*int32)(unsafe.Add(mBase, uint32(v212+int32(8)+v259)))
	v271 = *(*int32)(unsafe.Add(mBase, uint32(v212+int32(24)+v259)))
	v276 = *(*int32)(unsafe.Add(mBase, uint32(v212+int32(40)+v259)))
	v279 = v240 + base.B2i32(v261 != v262) + base.B2i32(v266 != v262) + base.B2i32(v271 != v262) + base.B2i32(v276 != v262)
	v281 = v244 + v258
	v283 = v257 + v258
	if v283 != v217&int32(-4) {
		v240 = v279
		v244 = v281
		v257 = v283
		goto L39
	} else {
		goto L41
	}
L40:
	;
	v291 = v279
	v295 = v281
	goto L36
L41:
	;
	goto L40
L42:
	;
	v379 = v353
	v383 = v209 + int32(1)
	goto L34
L43:
	;
	v317 = v291
	v318 = int32(0)
	v321 = v295
	goto L44
L44:
	;
	v338 = *(*int32)(unsafe.Add(mBase, uint32(v214+v321<<(uint(int32(4))%32))))
	v341 = v317 + base.B2i32(v338 != int32(0))
	v342 = int32(1)
	v345 = v318 + v342
	if v345 != v219 {
		v317 = v341
		v318 = v345
		v321 = v321 + v342
		goto L44
	} else {
		goto L46
	}
L45:
	;
	v353 = v341
	goto L42
L46:
	;
	goto L45
L47:
	;
	goto L30
L48:
	;
	v465 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
	v468 = v465 + v447<<(uint(int32(5))%32)
	v469 = *(*int32)(unsafe.Add(mBase, uint32(v468)+8))
	if v469 == int32(0) {
		v533 = v448
		v534 = v450
		goto L50
	} else {
		goto L51
	}
L49:
	;
	v539 = int32(0)
	v541 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	if v541 != int32(3) {
		v597 = v539
		goto L67
	} else {
		goto L68
	}
L50:
	;
	if v447 != 0 {
		v447 = v447 + int32(-1)
		v448 = v533
		v450 = v534
		goto L48
	} else {
		goto L66
	}
L51:
	;
	v472 = int32(0)
	v473 = *(*int32)(unsafe.Add(mBase, uint32(v468)+24))
	if v473 != int32(3) {
		v527 = v472
		goto L52
	} else {
		goto L53
	}
L52:
	;
	v533 = v527 + v448
	v534 = v450 + int32(1)
	goto L50
L53:
	;
	v476 = *(*float64)(unsafe.Add(mBase, uint32(v468)+16))
	if base.F64_lt(base.F64_abs(v476), float64(2.147483648e+09)) == int32(0) {
		goto L55
	} else {
		goto L56
	}
L54:
	;
	if base.F64_ne(v476, base.F64_convert_i32_s(v484)) != 0 {
		v527 = v472
		goto L52
	} else {
		goto L57
	}
L55:
	;
	v484 = int32(-2147483648)
	goto L54
L56:
	;
	v482 = base.I32_trunc_f64_s(v476)
	v484 = v482
	goto L54
L57:
	;
	v488 = v484 + int32(-1)
	if base.Ui32(int32(67108863)) < base.Ui32(v488) {
		v527 = v472
		goto L52
	} else {
		goto L58
	}
L58:
	;
	v495 = int32(-1)
	if base.Ui32(int32(256)) <= base.Ui32(v488) {
		goto L61
	} else {
		goto L62
	}
L59:
	;
	v520 = v30 + (v509+v516)<<(uint(int32(2))%32)
	v521 = *(*int32)(unsafe.Add(mBase, uint32(v520)))
	*(*int32)(unsafe.Add(mBase, uint32(v520))) = v521 + int32(1)
	v527 = int32(1)
	goto L52
L60:
	;
	v512 = m.G3
	v516 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v512+int32(_a2665)+v510))))
	goto L59
L61:
	;
	v498 = v488
	v499 = v495
	goto L63
L62:
	;
	v509 = v495
	v510 = v488
	goto L60
L63:
	;
	v502 = int32(8)
	v503 = v499 + v502
	v507 = int32(base.Ui32(v498) >> (uint(v502) % 32))
	if base.Ui32(int32(65535)) < base.Ui32(v498) {
		v498 = v507
		v499 = v503
		goto L63
	} else {
		goto L65
	}
L64:
	;
	v509 = v503
	v510 = v507
	goto L60
L65:
	;
	goto L64
L66:
	;
	goto L49
L67:
	;
	v598 = int32(1)
	v599 = int32(0)
	v603 = v533 + v421 + v597
	if v603 < v598 {
		v668 = v599
		v669 = v599
		goto L81
	} else {
		goto L82
	}
L68:
	;
	v544 = *(*float64)(unsafe.Add(mBase, uint32(l2)))
	if base.F64_lt(base.F64_abs(v544), float64(2.147483648e+09)) == int32(0) {
		goto L70
	} else {
		goto L71
	}
L69:
	;
	if base.F64_ne(v544, base.F64_convert_i32_s(v552)) != 0 {
		v597 = int32(0)
		goto L67
	} else {
		goto L72
	}
L70:
	;
	v552 = int32(-2147483648)
	goto L69
L71:
	;
	v550 = base.I32_trunc_f64_s(v544)
	v552 = v550
	goto L69
L72:
	;
	v558 = v552 + int32(-1)
	if base.Ui32(int32(67108863)) < base.Ui32(v558) {
		v597 = int32(0)
		goto L67
	} else {
		goto L73
	}
L73:
	;
	v565 = int32(-1)
	if base.Ui32(int32(256)) <= base.Ui32(v558) {
		goto L76
	} else {
		goto L77
	}
L74:
	;
	v590 = v30 + (v579+v586)<<(uint(int32(2))%32)
	v591 = *(*int32)(unsafe.Add(mBase, uint32(v590)))
	*(*int32)(unsafe.Add(mBase, uint32(v590))) = v591 + int32(1)
	v597 = int32(1)
	goto L67
L75:
	;
	v582 = m.G3
	v586 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v582+int32(_a2665)+v580))))
	goto L74
L76:
	;
	v568 = v558
	v569 = v565
	goto L78
L77:
	;
	v579 = v565
	v580 = v558
	goto L75
L78:
	;
	v572 = int32(8)
	v573 = v569 + v572
	v577 = int32(base.Ui32(v568) >> (uint(v572) % 32))
	if base.Ui32(int32(65535)) < base.Ui32(v568) {
		v568 = v577
		v569 = v573
		goto L78
	} else {
		goto L80
	}
L79:
	;
	v579 = v573
	v580 = v577
	goto L75
L80:
	;
	goto L79
L81:
	;
	F_resize_2(m, l0, l1, v668, v421+v534-v669+int32(1))
	mBase = m.M
	v681 = m.ExcPending
	if v681 != 0 {
		goto L96
	} else {
		goto L97
	}
L82:
	;
	v612 = v599
	v613 = v539
	v616 = v598
	v623 = v599
	v624 = v599
	v625 = v599
	goto L83
L83:
	;
	v633 = *(*int32)(unsafe.Add(mBase, uint32(v178+v623<<(uint(int32(2))%32))))
	v635 = base.B2i32(int32(0) < v633)
	v636 = v633 + v612
	v638 = v635 & base.B2i32(v613 < v636)
	if v638 != 0 {
		goto L85
	} else {
		goto L86
	}
L84:
	;
	v668 = v639
	v669 = v640
	goto L81
L85:
	;
	v639 = v616
	goto L87
L86:
	;
	v639 = v624
	goto L87
L87:
	;
	if v638 != 0 {
		goto L88
	} else {
		goto L89
	}
L88:
	;
	v640 = v636
	goto L90
L89:
	;
	v640 = v625
	goto L90
L90:
	;
	if int32(0) < v633 {
		goto L91
	} else {
		goto L92
	}
L91:
	;
	v641 = v636
	goto L93
L92:
	;
	v641 = v612
	goto L93
L93:
	;
	if v641 == v603 {
		v668 = v639
		v669 = v640
		goto L81
	} else {
		goto L94
	}
L94:
	;
	v643 = int32(1)
	v646 = v616 & int32(2147483647)
	if base.Ui32(v646) < base.Ui32(v603) {
		v612 = v641
		v613 = v646
		v616 = v616 << (uint(v643) % 32)
		v623 = v623 + v643
		v624 = v639
		v625 = v640
		goto L83
	} else {
		goto L95
	}
L95:
	;
	goto L84
L96:
	;
	return int32(0)
L97:
	;
	v682 = F_luaH_get(m, l1, l2)
	mBase = m.M
	v683 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l1)+6)) = uint8(v683)
	v685 = m.G398
	if v682 != v685 {
		v882 = v682
		goto L1
	} else {
		goto L98
	}
L98:
	;
	v687 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	v688 = m.G3
	switch v687 {
	case 0:
		v697 = v688 + int32(_a2686)
		goto L99
	default:
		goto L8
	case 3:
		goto L100
	}
L99:
	;
	F_luaG_runerror(m, l0, v697, int32(0))
	mBase = m.M
	v701 = m.ExcPending
	if v701 != 0 {
		goto L96
	} else {
		goto L102
	}
L100:
	;
	v691 = m.G3
	v692 = *(*float64)(unsafe.Add(mBase, uint32(l2)))
	if base.F64_eq(v692, v692) != 0 {
		goto L8
	} else {
		goto L101
	}
L101:
	;
	v697 = v691 + int32(_a2687)
	goto L99
L102:
	;
	goto L8
L103:
	;
	v705 = base.I64_reinterpret_f64(v702)
	v710 = int32(-1)
	v711 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+12)))
	v717 = base.I32_rem_u_s(base.I32_wrap_i64(int64(base.Ui64(v705)>>(uint(int64(32))%64))+v705), v710<<(uint(v711)%32)^v710|int32(1))
	v767 = v129 + v717<<(uint(int32(5))%32)
	goto L3
L104:
	;
	v823 = *(*int32)(unsafe.Add(mBase, uint32(v130)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v149+int32(-4)))) = v823
	*(*int32)(unsafe.Add(mBase, uint32(v130)+28)) = v165
	v833 = v165
	goto L2
L105:
	;
	v776 = v767
	goto L106
L106:
	;
	v794 = *(*int32)(unsafe.Add(mBase, uint32(v776)+28))
	if v794 != v130 {
		v776 = v794
		goto L106
	} else {
		goto L108
	}
L107:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v776)+28)) = v165
	v801 = *(*int64)(unsafe.Add(mBase, uint32(v130+int32(24))))
	*(*int64)(unsafe.Add(mBase, uint32(v149+int32(-8)))) = v801
	v807 = *(*int64)(unsafe.Add(mBase, uint32(v130+int32(16))))
	*(*int64)(unsafe.Add(mBase, uint32(v149+int32(-16)))) = v807
	v812 = v130 + int32(8)
	v813 = *(*int64)(unsafe.Add(mBase, uint32(v812)))
	*(*int64)(unsafe.Add(mBase, uint32(v149+int32(-24)))) = v813
	v815 = *(*int64)(unsafe.Add(mBase, uint32(v130)))
	*(*int64)(unsafe.Add(mBase, uint32(v165))) = v815
	v817 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v130)+28)) = v817
	*(*int32)(unsafe.Add(mBase, uint32(v812))) = v817
	v833 = v130
	goto L2
L108:
	;
	goto L107
L109:
	;
	v856 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	v857 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v856)+5)))
	if v857&int32(3) == int32(0) {
		v882 = v833
		goto L1
	} else {
		goto L110
	}
L110:
	;
	v862 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+5)))
	if v862&int32(4) == int32(0) {
		v882 = v833
		goto L1
	} else {
		goto L111
	}
L111:
	;
	v867 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v868 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+5)))
	v870 = v868 & int32(251)
	*(*uint8)(unsafe.Add(mBase, uint32(l1)+5)) = uint8(v870)
	v872 = *(*int32)(unsafe.Add(mBase, uint32(v867)+40))
	*(*int32)(unsafe.Add(mBase, uint32(l1)+32)) = v872
	*(*int32)(unsafe.Add(mBase, uint32(v867)+40)) = l1
	goto L112
L112:
	;
	v882 = v833
	goto L1
}
func F_ntohs(m *base.Module, l0 int32) int32 {
	var v2 int32
	_ = v2
	v2 = int32(8)
	return (l0<<(uint(v2)%32) | int32(base.Ui32(l0)>>(uint(v2)%32))) & int32(65535)
}
func F_nullBulkString(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v21 int32
	_ = v21
	var v25 int32
	_ = v25
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v47 int32
	_ = v47
	var v52 int32
	_ = v52
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v6 < int32(0) {
		v30 = v5
	} else {
		v11 = l0 + v6<<(uint(int32(2))%32)
		v12 = *(*int32)(unsafe.Add(mBase, uint32(v11)+16))
		v13 = int32(1)
		v14 = v12 + v13
		*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = v14
		v16 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1040))
		if v16 != v13 {
			v30 = v5
		} else {
			v21 = *(*int32)(unsafe.Add(mBase, uint32(v5)+8))
			*(*int32)(unsafe.Add(mBase, uint32(v21)+8)) = int32(3)
			*(*float64)(unsafe.Add(mBase, uint32(v21))) = base.F64_convert_i32_u(v14)
			v25 = *(*int32)(unsafe.Add(mBase, uint32(v5)+8))
			*(*int32)(unsafe.Add(mBase, uint32(v5)+8)) = v25 + int32(16)
			v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			v30 = v29
		}
	}
	v34 = F_lua_checkstack(m, v30, int32(1))
	mBase = m.M
	v35 = m.ExcPending
	if v35 != 0 {
		return
	} else {
		if v34 != 0 {
			v41 = *(*int32)(unsafe.Add(mBase, uint32(v30)+8))
			*(*int32)(unsafe.Add(mBase, uint32(v41)+8)) = int32(1)
			*(*int32)(unsafe.Add(mBase, uint32(v41))) = int32(0)
			v47 = *(*int32)(unsafe.Add(mBase, uint32(v30)+8))
			*(*int32)(unsafe.Add(mBase, uint32(v30)+8)) = v47 + int32(16)
			F_processCollectionElementEnd(m, l0)
			mBase = m.M
			v52 = m.ExcPending
			if v52 != 0 {
				return
			} else {
				return
			}
		} else {
			F__serverPanic_2(m, int32(903))
			mBase = m.M
			v38 = m.ExcPending
			if v38 != 0 {
				return
			} else {
				base.Wasm_trap_unreachable()
				for {
				}
			}
		}
	}
}
func F_nullCallback(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v21 int32
	_ = v21
	var v25 int32
	_ = v25
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v47 int32
	_ = v47
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v6 < int32(0) {
		v30 = v5
	} else {
		v11 = l0 + v6<<(uint(int32(2))%32)
		v12 = *(*int32)(unsafe.Add(mBase, uint32(v11)+16))
		v13 = int32(1)
		v14 = v12 + v13
		*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = v14
		v16 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1040))
		if v16 != v13 {
			v30 = v5
		} else {
			v21 = *(*int32)(unsafe.Add(mBase, uint32(v5)+8))
			*(*int32)(unsafe.Add(mBase, uint32(v21)+8)) = int32(3)
			*(*float64)(unsafe.Add(mBase, uint32(v21))) = base.F64_convert_i32_u(v14)
			v25 = *(*int32)(unsafe.Add(mBase, uint32(v5)+8))
			*(*int32)(unsafe.Add(mBase, uint32(v5)+8)) = v25 + int32(16)
			v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			v30 = v29
		}
	}
	v34 = F_lua_checkstack(m, v30, int32(1))
	mBase = m.M
	v35 = m.ExcPending
	if v35 != 0 {
		return
	} else {
		if v34 != 0 {
			v40 = *(*int32)(unsafe.Add(mBase, uint32(v30)+8))
			*(*int32)(unsafe.Add(mBase, uint32(v40)+8)) = int32(0)
			*(*int32)(unsafe.Add(mBase, uint32(v30)+8)) = v40 + int32(16)
			F_processCollectionElementEnd(m, l0)
			mBase = m.M
			v47 = m.ExcPending
			if v47 != 0 {
				return
			} else {
				return
			}
		} else {
			F__serverPanic_2(m, int32(879))
			mBase = m.M
			v38 = m.ExcPending
			if v38 != 0 {
				return
			} else {
				base.Wasm_trap_unreachable()
				for {
				}
			}
		}
	}
}
func F_numericConfigInit(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v2 int64
	_ = v2
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	v2 = *(*int64)(unsafe.Add(mBase, uint32(l0)+64))
	v4 = F_setNumericType(m, l0, v2, int32(0))
	mBase = m.M
	v5 = m.ExcPending
	if v5 != 0 {
		return
	} else {
		return
	}
}
func F_numericConfigRewrite(m *base.Module, l0 int32, l1 int32, l2 int32) {
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
	var v16 int64
	_ = v16
	var v17 int32
	_ = v17
	var v22 int32
	_ = v22
	var v23 int64
	_ = v23
	var v24 int32
	_ = v24
	var v25 int64
	_ = v25
	var v26 int32
	_ = v26
	var v31 int32
	_ = v31
	var v32 int64
	_ = v32
	var v33 int32
	_ = v33
	var v34 int64
	_ = v34
	var v35 int32
	_ = v35
	var v36 int64
	_ = v36
	var v37 int32
	_ = v37
	var v38 int64
	_ = v38
	var v39 int32
	_ = v39
	var v40 int64
	_ = v40
	var v41 int32
	_ = v41
	var v42 int64
	_ = v42
	var v43 int64
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int64
	_ = v46
	var v47 int64
	_ = v47
	var v48 int32
	_ = v48
	var v55 int64
	_ = v55
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v60 int64
	_ = v60
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v76 int64
	_ = v76
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v89 int64
	_ = v89
	var v90 int32
	_ = v90
	var v95 int64
	_ = v95
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v119 int64
	_ = v119
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v132 int64
	_ = v132
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
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
	v9 = m.G0
	v11 = v9 - int32(160)
	m.G0 = v11
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	switch v14 {
	case 0:
		goto L4
	case 1, 3, 6:
		goto L6
	case 2:
		goto L12
	case 4:
		goto L11
	case 5:
		goto L10
	case 7:
		goto L9
	case 8:
		goto L8
	case 9:
		goto L7
	default:
		v43 = int64(0)
		goto L5
	}
L1:
	;
	m.G0 = v11 + int32(160)
	return
L2:
	;
	if v90&int32(1) == int32(0) {
		goto L29
	} else {
		goto L30
	}
L3:
	;
	v48 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	if v48&int32(2) == int32(0) {
		goto L17
	} else {
		goto L18
	}
L4:
	;
	v45 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	v46 = int64(*(*int32)(unsafe.Add(mBase, uint32(v45))))
	v47 = v46
	goto L3
L5:
	;
	v44 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v89 = v43
	v90 = v44
	goto L2
L6:
	;
	v41 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	v42 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v41))))
	v43 = v42
	goto L5
L7:
	;
	v39 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	v40 = *(*int64)(unsafe.Add(mBase, uint32(v39)))
	v47 = v40
	goto L3
L8:
	;
	v37 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	v38 = *(*int64)(unsafe.Add(mBase, uint32(v37)))
	v47 = v38
	goto L3
L9:
	;
	v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	v36 = int64(*(*int32)(unsafe.Add(mBase, uint32(v35))))
	v47 = v36
	goto L3
L10:
	;
	v26 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+9)))
	if v26&int32(1) == int32(0) {
		goto L15
	} else {
		goto L16
	}
L11:
	;
	v17 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+9)))
	if v17&int32(1) == int32(0) {
		goto L13
	} else {
		goto L14
	}
L12:
	;
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	v16 = int64(*(*int32)(unsafe.Add(mBase, uint32(v15))))
	v47 = v16
	goto L3
L13:
	;
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	v25 = *(*int64)(unsafe.Add(mBase, uint32(v24)))
	v47 = v25
	goto L3
L14:
	;
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+84))
	v23 = F_getModuleNumericConfig(m, v22)
	mBase = m.M
	v47 = v23
	goto L3
L15:
	;
	v33 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	v34 = *(*int64)(unsafe.Add(mBase, uint32(v33)))
	v47 = v34
	goto L3
L16:
	;
	v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)+84))
	v32 = F_getModuleUnsignedNumericConfig(m, v31)
	mBase = m.M
	v47 = v32
	goto L3
L17:
	;
	if v48&int32(16) == int32(0) {
		v89 = v47
		v90 = v48
		goto L2
	} else {
		goto L24
	}
L18:
	;
	if int64(-1) < v47 {
		goto L17
	} else {
		goto L19
	}
L19:
	;
	v55 = *(*int64)(unsafe.Add(mBase, uint32(l0)+64))
	v56 = F_sdsempty(m)
	mBase = m.M
	v57 = m.ExcPending
	if v57 != 0 {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	return
L21:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+64)) = l1
	v60 = int64(0) - v47
	*(*int64)(unsafe.Add(mBase, uint32(v11)+72)) = v60
	v65 = F_sdscatprintf(m, v56, int32(_a523), v11+int32(64))
	mBase = m.M
	v66 = m.ExcPending
	if v66 != 0 {
		goto L20
	} else {
		goto L22
	}
L22:
	;
	v68 = F_rewriteConfigRewriteLine(m, l2, l1, v65, base.B2i32(v55 != v60))
	mBase = m.M
	v69 = m.ExcPending
	if v69 != 0 {
		goto L20
	} else {
		goto L23
	}
L23:
	;
	goto L1
L24:
	;
	if int64(-1) < v47 {
		v89 = v47
		v90 = v48
		goto L2
	} else {
		goto L25
	}
L25:
	;
	v76 = *(*int64)(unsafe.Add(mBase, uint32(l0)+64))
	v77 = F_sdsempty(m)
	mBase = m.M
	v78 = m.ExcPending
	if v78 != 0 {
		goto L20
	} else {
		goto L26
	}
L26:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+80)) = l1
	*(*int64)(unsafe.Add(mBase, uint32(v11)+88)) = v47
	v84 = F_sdscatprintf(m, v77, int32(_a520), v11+int32(80))
	mBase = m.M
	v85 = m.ExcPending
	if v85 != 0 {
		goto L20
	} else {
		goto L27
	}
L27:
	;
	v87 = F_rewriteConfigRewriteLine(m, l2, l1, v84, base.B2i32(v47 != v76))
	mBase = m.M
	v88 = m.ExcPending
	if v88 != 0 {
		goto L20
	} else {
		goto L28
	}
L28:
	;
	goto L1
L29:
	;
	if v90&int32(4) == int32(0) {
		goto L35
	} else {
		goto L36
	}
L30:
	;
	v95 = *(*int64)(unsafe.Add(mBase, uint32(l0)+64))
	v99 = F_rewriteConfigFormatMemory(m, v11+int32(96), int32(64), v89)
	mBase = m.M
	v100 = m.ExcPending
	if v100 != 0 {
		goto L20
	} else {
		goto L31
	}
L31:
	;
	v101 = F_sdsempty(m)
	mBase = m.M
	v102 = m.ExcPending
	if v102 != 0 {
		goto L20
	} else {
		goto L32
	}
L32:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+48)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v11)+52)) = v11 + int32(96)
	v110 = F_sdscatprintf(m, v101, int32(_a518), v11+int32(48))
	mBase = m.M
	v111 = m.ExcPending
	if v111 != 0 {
		goto L20
	} else {
		goto L33
	}
L33:
	;
	v113 = F_rewriteConfigRewriteLine(m, l2, l1, v110, base.B2i32(v89 != v95))
	mBase = m.M
	v114 = m.ExcPending
	if v114 != 0 {
		goto L20
	} else {
		goto L34
	}
L34:
	;
	goto L1
L35:
	;
	v132 = *(*int64)(unsafe.Add(mBase, uint32(l0)+64))
	v133 = base.B2i32(v89 != v132)
	v134 = F_sdsempty(m)
	mBase = m.M
	v135 = m.ExcPending
	if v135 != 0 {
		goto L20
	} else {
		goto L40
	}
L36:
	;
	v119 = *(*int64)(unsafe.Add(mBase, uint32(l0)+64))
	v120 = F_sdsempty(m)
	mBase = m.M
	v121 = m.ExcPending
	if v121 != 0 {
		goto L20
	} else {
		goto L37
	}
L37:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+32)) = l1
	*(*int64)(unsafe.Add(mBase, uint32(v11)+40)) = v89
	v127 = F_sdscatprintf(m, v120, int32(_a522), v11+int32(32))
	mBase = m.M
	v128 = m.ExcPending
	if v128 != 0 {
		goto L20
	} else {
		goto L38
	}
L38:
	;
	v130 = F_rewriteConfigRewriteLine(m, l2, l1, v127, base.B2i32(v89 != v119))
	mBase = m.M
	v131 = m.ExcPending
	if v131 != 0 {
		goto L20
	} else {
		goto L39
	}
L39:
	;
	goto L1
L40:
	;
	if v90&int32(8) == int32(0) {
		goto L41
	} else {
		goto L42
	}
L41:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v11)+8)) = v89
	*(*int32)(unsafe.Add(mBase, uint32(v11))) = l1
	v152 = F_sdscatprintf(m, v134, int32(_a520), v11)
	mBase = m.M
	v153 = m.ExcPending
	if v153 != 0 {
		goto L20
	} else {
		goto L45
	}
L42:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v11)+24)) = v89
	*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = l1
	v145 = F_sdscatprintf(m, v134, int32(_a521), v11+int32(16))
	mBase = m.M
	v146 = m.ExcPending
	if v146 != 0 {
		goto L20
	} else {
		goto L43
	}
L43:
	;
	v147 = F_rewriteConfigRewriteLine(m, l2, l1, v145, v133)
	mBase = m.M
	v148 = m.ExcPending
	if v148 != 0 {
		goto L20
	} else {
		goto L44
	}
L44:
	;
	goto L1
L45:
	;
	v154 = F_rewriteConfigRewriteLine(m, l2, l1, v152, v133)
	mBase = m.M
	v155 = m.ExcPending
	if v155 != 0 {
		goto L20
	} else {
		goto L46
	}
L46:
	;
	goto L1
}
