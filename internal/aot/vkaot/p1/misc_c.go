package p1

import (
	base "github.com/shibukawa/vkmem/internal/aot/vkaot/base"
	"unsafe"
)

func F___cos(m *base.Module, l0 float64, l1 float64) float64 {
	var v6 float64
	_ = v6
	var v7 float64
	_ = v7
	var v9 float64
	_ = v9
	var v10 float64
	_ = v10
	var v22 float64
	_ = v22
	v6 = float64(1)
	v7 = base.F64_mul(l0, l0)
	v9 = base.F64_mul(v7, float64(0.5))
	v10 = base.F64_sub(v6, v9)
	v22 = base.F64_mul(v7, v7)
	return base.F64_add(v10, base.F64_add(base.F64_sub(base.F64_sub(v6, v10), v9), base.F64_sub(base.F64_mul(v7, base.F64_add(base.F64_mul(v7, base.F64_add(base.F64_mul(v7, base.F64_add(base.F64_mul(v7, float64(2.480158728947673e-05)), float64(-0.001388888888887411))), float64(0.0416666666666666))), base.F64_mul(base.F64_mul(v22, v22), base.F64_add(base.F64_mul(v7, base.F64_add(base.F64_mul(v7, float64(-1.1359647557788195e-11)), float64(2.087572321298175e-09))), float64(-2.7557314351390663e-07))))), base.F64_mul(l0, l1))))
}
func F_call(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v37 int32
	_ = v37
	var v40 int64
	_ = v40
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v50 int32
	_ = v50
	var v53 int64
	_ = v53
	var v56 int64
	_ = v56
	var v58 int64
	_ = v58
	var v59 int64
	_ = v59
	var v62 int32
	_ = v62
	var v68 int64
	_ = v68
	var v69 int64
	_ = v69
	var v70 int32
	_ = v70
	var v73 int64
	_ = v73
	var v74 int64
	_ = v74
	var v78 int64
	_ = v78
	var v81 int32
	_ = v81
	var v85 int32
	_ = v85
	var v90 int64
	_ = v90
	var v94 int64
	_ = v94
	var v99 int32
	_ = v99
	var v101 int64
	_ = v101
	var v105 int32
	_ = v105
	var v113 int32
	_ = v113
	var v119 int32
	_ = v119
	var v120 int64
	_ = v120
	var v121 int64
	_ = v121
	var v122 int32
	_ = v122
	var v125 int32
	_ = v125
	var v128 int32
	_ = v128
	var v132 int32
	_ = v132
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v140 int32
	_ = v140
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v149 int32
	_ = v149
	var v151 int32
	_ = v151
	var v157 int32
	_ = v157
	var v177 int32
	_ = v177
	var v179 int32
	_ = v179
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v186 int32
	_ = v186
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v214 int32
	_ = v214
	var v216 int32
	_ = v216
	var v236 int32
	_ = v236
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v247 int32
	_ = v247
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v272 int32
	_ = v272
	var v273 int32
	_ = v273
	var v274 int32
	_ = v274
	var v282 int32
	_ = v282
	var v283 int32
	_ = v283
	var v285 int32
	_ = v285
	var v286 int32
	_ = v286
	var v291 int32
	_ = v291
	var v292 int32
	_ = v292
	var v296 int32
	_ = v296
	var v297 int32
	_ = v297
	var v299 int32
	_ = v299
	var v300 int32
	_ = v300
	var v302 int32
	_ = v302
	var v305 int32
	_ = v305
	var v308 int32
	_ = v308
	var v309 int32
	_ = v309
	var v318 int32
	_ = v318
	var v323 int32
	_ = v323
	var v324 int32
	_ = v324
	var v333 int32
	_ = v333
	var v348 int32
	_ = v348
	var v351 int32
	_ = v351
	var v352 int32
	_ = v352
	var v353 int32
	_ = v353
	var v354 int32
	_ = v354
	var v355 int32
	_ = v355
	var v359 int32
	_ = v359
	var v360 int32
	_ = v360
	var v368 int32
	_ = v368
	var v369 int32
	_ = v369
	var v370 int32
	_ = v370
	var v371 int32
	_ = v371
	var v372 int32
	_ = v372
	var v375 int32
	_ = v375
	var v376 int32
	_ = v376
	var v378 int32
	_ = v378
	var v379 int32
	_ = v379
	var v380 int32
	_ = v380
	var v384 int32
	_ = v384
	var v385 int32
	_ = v385
	var v393 int32
	_ = v393
	var v394 int32
	_ = v394
	var v395 int32
	_ = v395
	var v398 int32
	_ = v398
	var v400 int32
	_ = v400
	var v401 int32
	_ = v401
	var v402 int32
	_ = v402
	var v403 int32
	_ = v403
	var v409 int32
	_ = v409
	var v412 int32
	_ = v412
	var v439 int32
	_ = v439
	var v441 int32
	_ = v441
	var v466 int32
	_ = v466
	var v468 int32
	_ = v468
	var v472 int32
	_ = v472
	var v475 int32
	_ = v475
	var v480 int32
	_ = v480
	var v486 int32
	_ = v486
	var v487 int64
	_ = v487
	var v489 int64
	_ = v489
	var v491 int64
	_ = v491
	var v492 int32
	_ = v492
	var v496 int32
	_ = v496
	var v498 int64
	_ = v498
	var v499 int64
	_ = v499
	var v500 int64
	_ = v500
	var v503 int64
	_ = v503
	var v504 int32
	_ = v504
	var v505 int64
	_ = v505
	var v509 int64
	_ = v509
	var v511 int64
	_ = v511
	var v517 int32
	_ = v517
	var v519 int32
	_ = v519
	var v522 int64
	_ = v522
	var v527 int32
	_ = v527
	var v529 int32
	_ = v529
	var v530 int32
	_ = v530
	var v538 int32
	_ = v538
	var v543 int64
	_ = v543
	var v551 int64
	_ = v551
	var v556 int32
	_ = v556
	var v558 int32
	_ = v558
	var v560 int32
	_ = v560
	var v571 int32
	_ = v571
	var v572 int64
	_ = v572
	var v577 int32
	_ = v577
	var v578 int64
	_ = v578
	var v583 int64
	_ = v583
	var v591 int32
	_ = v591
	var v595 int32
	_ = v595
	var v596 int32
	_ = v596
	var v597 int32
	_ = v597
	var v601 int32
	_ = v601
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
	var v611 int32
	_ = v611
	var v613 int32
	_ = v613
	var v616 int32
	_ = v616
	var v619 int64
	_ = v619
	var v623 int64
	_ = v623
	var v624 int32
	_ = v624
	var v625 int64
	_ = v625
	var v628 int32
	_ = v628
	var v629 int32
	_ = v629
	var v632 int32
	_ = v632
	var v641 int32
	_ = v641
	var v642 int64
	_ = v642
	var v643 int64
	_ = v643
	var v645 int32
	_ = v645
	var v650 int32
	_ = v650
	var v655 int32
	_ = v655
	var v662 int32
	_ = v662
	var v665 int32
	_ = v665
	var v670 int32
	_ = v670
	var v675 int32
	_ = v675
	var v682 int32
	_ = v682
	var v689 int32
	_ = v689
	var v696 int32
	_ = v696
	var v703 int32
	_ = v703
	var v706 int32
	_ = v706
	var v707 int32
	_ = v707
	var v708 int32
	_ = v708
	var v709 int32
	_ = v709
	var v710 int32
	_ = v710
	var v712 int32
	_ = v712
	var v713 int32
	_ = v713
	var v714 int32
	_ = v714
	var v715 int32
	_ = v715
	var v716 int32
	_ = v716
	var v722 int32
	_ = v722
	var v724 int32
	_ = v724
	var v729 int32
	_ = v729
	var v736 int32
	_ = v736
	var v737 int32
	_ = v737
	var v740 int32
	_ = v740
	var v746 int32
	_ = v746
	var v747 int32
	_ = v747
	var v749 int32
	_ = v749
	var v752 int32
	_ = v752
	var v753 int32
	_ = v753
	var v756 int64
	_ = v756
	var v760 int32
	_ = v760
	var v762 int64
	_ = v762
	var v767 int32
	_ = v767
	var v776 int32
	_ = v776
	var v780 int32
	_ = v780
	var v784 int32
	_ = v784
	var v785 int32
	_ = v785
	var v787 int32
	_ = v787
	var v794 int32
	_ = v794
	var v796 int32
	_ = v796
	var v797 int32
	_ = v797
	var v801 int32
	_ = v801
	var v804 int32
	_ = v804
	var v807 int32
	_ = v807
	var v810 int32
	_ = v810
	var v813 int32
	_ = v813
	var v816 int32
	_ = v816
	var v820 int32
	_ = v820
	var v821 int32
	_ = v821
	var v822 int32
	_ = v822
	var v824 int32
	_ = v824
	var v826 int32
	_ = v826
	var v827 int32
	_ = v827
	var v835 int32
	_ = v835
	var v836 int32
	_ = v836
	var v838 int32
	_ = v838
	var v846 int32
	_ = v846
	var v847 int32
	_ = v847
	var v848 int32
	_ = v848
	var v851 int32
	_ = v851
	var v853 int32
	_ = v853
	var v861 int32
	_ = v861
	var v866 int32
	_ = v866
	var v868 int32
	_ = v868
	var v870 int32
	_ = v870
	var v872 int32
	_ = v872
	var v873 int32
	_ = v873
	var v874 int32
	_ = v874
	var v875 int32
	_ = v875
	var v879 int32
	_ = v879
	var v882 int32
	_ = v882
	var v883 int32
	_ = v883
	var v889 int32
	_ = v889
	var v891 int32
	_ = v891
	var v901 int64
	_ = v901
	var v905 int32
	_ = v905
	var v906 int32
	_ = v906
	var v907 int32
	_ = v907
	var v910 int32
	_ = v910
	var v922 int32
	_ = v922
	var v928 int32
	_ = v928
	v3 = int32(0)
	v25 = m.G0
	v27 = v25 - int32(48)
	m.G0 = v27
	v30 = *(*int32)(unsafe.Add(mBase, _consts[300]))
	v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
	v32 = *(*int32)(unsafe.Add(mBase, uint32(l0)+200))
	*(*int32)(unsafe.Add(mBase, _consts[300])) = l0
	v37 = *(*int32)(unsafe.Add(mBase, _consts[14]))
	if v37 == v3 {
		v43 = v3
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v44 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+196)) = v44
	v46 = *(*int32)(unsafe.Add(mBase, uint32(l0)+200))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+200)) = v46 & int32(-2146305)
	v50 = *(*int32)(unsafe.Add(mBase, uint32(l0)+204))
	v53 = *(*int64)(unsafe.Add(mBase, _consts[421]))
	*(*int64)(unsafe.Add(mBase, _consts[422])) = v53
	v56 = *(*int64)(unsafe.Add(mBase, _consts[47]))
	v58 = *(*int64)(unsafe.Add(mBase, _consts[180]))
	v59 = F_ustime(m)
	mBase = m.M
	v62 = *(*int32)(unsafe.Add(mBase, _consts[177]))
	*(*int32)(unsafe.Add(mBase, _consts[177])) = v62 + int32(1)
	if v62 != 0 {
		goto L3
	} else {
		goto L4
	}
L2:
	;
	v40 = *(*int64)(unsafe.Add(mBase, uint32(v37)))
	v43 = base.B2i32(v40 == int64(-1))
	goto L1
L3:
	;
	v105 = *(*int32)(unsafe.Add(mBase, uint32(l0)+204))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+204)) = v105&int32(1073741822) | int32(1)
	v113 = *(*int32)(unsafe.Add(mBase, _consts[34]))
	goto L9
L4:
	;
	if v59 != int64(0) {
		v69 = v59
		goto L5
	} else {
		goto L6
	}
L5:
	;
	v70 = int32(0)
	*(*int64)(unsafe.Add(mBase, _consts[178])) = v69
	v73 = int64(1000)
	v74 = base.I64_div_s(v69, v73)
	*(*int64)(unsafe.Add(mBase, _consts[54])) = v74
	v78 = base.I64_div_s(v69, int64(1000000))
	*(*int64)(unsafe.Add(mBase, _consts[109])) = v78
	v81 = *(*int32)(unsafe.Add(mBase, _consts[179]))
	v85 = int32(base.Ui32(v81&int32(2)) >> (uint(int32(1)) % 32))
	*(*uint8)(unsafe.Add(mBase, _consts[378])) = uint8(v85)
	v90 = base.I64_div_s(v74, int64(60000))
	*(*uint16)(unsafe.Add(mBase, _consts[376])) = uint16(v90)
	v94 = base.I64_div_s(v74, v73)
	*(*int32)(unsafe.Add(mBase, _consts[379])) = base.I32_wrap_i64(v94) & int32(16777215)
	goto L7
L6:
	;
	v68 = F_ustime(m)
	mBase = m.M
	v69 = v68
	goto L5
L7:
	;
	v99 = int32(0)
	v101 = *(*int64)(unsafe.Add(mBase, _consts[54]))
	*(*int64)(unsafe.Add(mBase, _consts[12])) = v101
	goto L3
L8:
	;
	v122 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+208)))
	if v122&int32(1) != 0 {
		goto L12
	} else {
		goto L13
	}
L9:
	;
	if base.B2i32(v113 != int32(948)) != int32(1) {
		v121 = int64(0)
		goto L8
	} else {
		goto L10
	}
L10:
	;
	v119 = *(*int32)(unsafe.Add(mBase, _consts[34]))
	v120 = m.T0[v119].(func(*base.Module) int64)(m)
	mBase = m.M
	v121 = v120
	goto L8
L11:
	;
	v282 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
	v283 = *(*int32)(unsafe.Add(mBase, uint32(v282)+48))
	m.T0[v283].(func(*base.Module, int32))(m, l0)
	mBase = m.M
	v285 = m.ExcPending
	if v285 != 0 {
		goto L18
	} else {
		goto L34
	}
L12:
	;
	v128 = int32(0)
	v132 = *(*int32)(unsafe.Add(mBase, _consts[113]))
	if v132 == v128 {
		v272 = v128
		v273 = v128
		v274 = v128
		goto L11
	} else {
		goto L14
	}
L13:
	;
	v125 = int32(0)
	v272 = v125
	v273 = v125
	v274 = v125
	goto L11
L14:
	;
	v137 = *(*int32)(unsafe.Add(mBase, uint32(l0)+192))
	if v137 != 0 {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	v138 = int32(188)
	goto L17
L16:
	;
	v138 = int32(24)
	goto L17
L17:
	;
	v140 = *(*int32)(unsafe.Add(mBase, uint32(l0+v138)))
	v142 = v140 << (uint(int32(2)) % 32)
	v143 = F_valkey_malloc(m, v142)
	mBase = m.M
	v144 = m.ExcPending
	if v144 != 0 {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	return
L19:
	;
	v145 = F_valkey_malloc(m, v142)
	mBase = m.M
	v146 = m.ExcPending
	if v146 != 0 {
		goto L18
	} else {
		goto L20
	}
L20:
	;
	if v140 < int32(1) {
		v272 = v145
		v273 = v140
		v274 = v143
		goto L11
	} else {
		goto L21
	}
L21:
	;
	v149 = int32(1)
	v151 = *(*int32)(unsafe.Add(mBase, uint32(l0)+192))
	if v140 == v149 {
		v236 = int32(0)
		goto L22
	} else {
		goto L23
	}
L22:
	;
	if v140&v149 == int32(0) {
		v272 = v145
		v273 = v140
		v274 = v143
		goto L11
	} else {
		goto L31
	}
L23:
	;
	v157 = int32(0)
	v177 = v157
	v179 = v157
	goto L24
L24:
	;
	if v151 != 0 {
		v184 = v151
		goto L26
	} else {
		goto L27
	}
L25:
	;
	v236 = v214
	goto L22
L26:
	;
	v186 = v177 << (uint(int32(2)) % 32)
	v188 = v184 + v186
	v189 = *(*int32)(unsafe.Add(mBase, uint32(v188)))
	*(*int32)(unsafe.Add(mBase, uint32(v143+v186))) = v189
	v192 = *(*int32)(unsafe.Add(mBase, uint32(v188)))
	v193 = *(*int32)(unsafe.Add(mBase, uint32(v192)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v145+v186))) = int32(base.Ui32(v193) >> (uint(int32(3)) % 32))
	if v151 != 0 {
		v198 = v151
		goto L28
	} else {
		goto L29
	}
L27:
	;
	v183 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v184 = v183
	goto L26
L28:
	;
	v201 = int32(2)
	v202 = (v177 | int32(1)) << (uint(v201) % 32)
	v204 = v198 + v202
	v205 = *(*int32)(unsafe.Add(mBase, uint32(v204)))
	*(*int32)(unsafe.Add(mBase, uint32(v143+v202))) = v205
	v208 = *(*int32)(unsafe.Add(mBase, uint32(v204)))
	v209 = *(*int32)(unsafe.Add(mBase, uint32(v208)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v145+v202))) = int32(base.Ui32(v209) >> (uint(int32(3)) % 32))
	v214 = v177 + v201
	v216 = v179 + v201
	if v216 != v140&int32(2147483646) {
		v177 = v214
		v179 = v216
		goto L24
	} else {
		goto L30
	}
L29:
	;
	v197 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v198 = v197
	goto L28
L30:
	;
	goto L25
L31:
	;
	if v151 != 0 {
		v245 = v151
		goto L32
	} else {
		goto L33
	}
L32:
	;
	v247 = v236 << (uint(int32(2)) % 32)
	v249 = v245 + v247
	v250 = *(*int32)(unsafe.Add(mBase, uint32(v249)))
	*(*int32)(unsafe.Add(mBase, uint32(v143+v247))) = v250
	v253 = *(*int32)(unsafe.Add(mBase, uint32(v249)))
	v254 = *(*int32)(unsafe.Add(mBase, uint32(v253)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v145+v247))) = int32(base.Ui32(v254) >> (uint(int32(3)) % 32))
	v272 = v145
	v273 = v140
	v274 = v143
	goto L11
L33:
	;
	v244 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v245 = v244
	goto L32
L34:
	;
	v286 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+208)))
	if v286&int32(1) == int32(0) {
		goto L37
	} else {
		goto L38
	}
L35:
	;
	F__serverAssert(m, int32(_a2221), int32(_a2157), int32(3914))
	mBase = m.M
	v928 = m.ExcPending
	if v928 != 0 {
		goto L18
	} else {
		goto L183
	}
L36:
	;
	F__serverAssert(m, int32(_a2222), int32(_a2157), int32(3919))
	mBase = m.M
	v922 = m.ExcPending
	if v922 != 0 {
		goto L18
	} else {
		goto L182
	}
L37:
	;
	v466 = int32(0)
	v468 = *(*int32)(unsafe.Add(mBase, _consts[177]))
	*(*int32)(unsafe.Add(mBase, _consts[177])) = v468 + int32(-1)
	v472 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+200)))
	if v472&int32(16) != 0 {
		goto L66
	} else {
		goto L67
	}
L38:
	;
	v291 = int32(0)
	v292 = *(*int32)(unsafe.Add(mBase, _consts[113]))
	if v292 == v291 {
		goto L37
	} else {
		goto L39
	}
L39:
	;
	v296 = *(*int32)(unsafe.Add(mBase, uint32(l0)+192))
	if v296 != 0 {
		v299 = int32(188)
		v300 = v296
		goto L40
	} else {
		goto L41
	}
L40:
	;
	v302 = *(*int32)(unsafe.Add(mBase, uint32(l0+v299)))
	if v302 == v273 {
		goto L42
	} else {
		goto L43
	}
L41:
	;
	v297 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v299 = int32(24)
	v300 = v297
	goto L40
L42:
	;
	if v302 != v273 {
		goto L35
	} else {
		goto L46
	}
L43:
	;
	v305 = *(*int32)(unsafe.Add(mBase, _consts[28]))
	if int32(3) < v305 {
		goto L42
	} else {
		goto L44
	}
L44:
	;
	v308 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
	v309 = *(*int32)(unsafe.Add(mBase, uint32(v308)+144))
	*(*int32)(unsafe.Add(mBase, uint32(v27)+40)) = v302
	*(*int32)(unsafe.Add(mBase, uint32(v27)+36)) = v273
	*(*int32)(unsafe.Add(mBase, uint32(v27)+32)) = v309
	F__serverLog(m, int32(3), int32(_a2223), v27+int32(32))
	mBase = m.M
	v318 = m.ExcPending
	if v318 != 0 {
		goto L18
	} else {
		goto L45
	}
L45:
	;
	goto L42
L46:
	;
	if v273 < int32(1) {
		goto L47
	} else {
		goto L48
	}
L47:
	;
	F_valkey_free(m, v274)
	mBase = m.M
	v439 = m.ExcPending
	if v439 != 0 {
		goto L18
	} else {
		goto L64
	}
L48:
	;
	v323 = int32(0)
	v324 = *(*int32)(unsafe.Add(mBase, _consts[28]))
	v333 = v323
	v348 = v324
	goto L49
L49:
	;
	v351 = v333 << (uint(int32(2)) % 32)
	v352 = v300 + v351
	v353 = *(*int32)(unsafe.Add(mBase, uint32(v352)))
	v354 = v274 + v351
	v355 = *(*int32)(unsafe.Add(mBase, uint32(v354)))
	if v353 == v355 {
		v371 = v353
		v372 = v355
		goto L51
	} else {
		goto L52
	}
L50:
	;
	goto L47
L51:
	;
	if v372 != v371 {
		goto L36
	} else {
		goto L55
	}
L52:
	;
	if int32(3) < v348 {
		v371 = v353
		v372 = v355
		goto L51
	} else {
		goto L53
	}
L53:
	;
	v359 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
	v360 = *(*int32)(unsafe.Add(mBase, uint32(v359)+144))
	*(*int32)(unsafe.Add(mBase, uint32(v27)+20)) = v333
	*(*int32)(unsafe.Add(mBase, uint32(v27)+16)) = v360
	F__serverLog(m, int32(3), int32(_a2224), v27+int32(16))
	mBase = m.M
	v368 = m.ExcPending
	if v368 != 0 {
		goto L18
	} else {
		goto L54
	}
L54:
	;
	v369 = *(*int32)(unsafe.Add(mBase, uint32(v352)))
	v370 = *(*int32)(unsafe.Add(mBase, uint32(v354)))
	v371 = v369
	v372 = v370
	goto L51
L55:
	;
	v375 = *(*int32)(unsafe.Add(mBase, _consts[28]))
	v376 = *(*int32)(unsafe.Add(mBase, uint32(v371)+4))
	v378 = int32(base.Ui32(v376) >> (uint(int32(3)) % 32))
	v379 = v272 + v351
	v380 = *(*int32)(unsafe.Add(mBase, uint32(v379)))
	if v380 <= v378 {
		v401 = v378
		v402 = v380
		v403 = v375
		goto L56
	} else {
		goto L57
	}
L56:
	;
	if v402 <= v401 {
		goto L60
	} else {
		goto L61
	}
L57:
	;
	if int32(3) < v375 {
		v401 = v378
		v402 = v380
		v403 = v375
		goto L56
	} else {
		goto L58
	}
L58:
	;
	v384 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
	v385 = *(*int32)(unsafe.Add(mBase, uint32(v384)+144))
	*(*int32)(unsafe.Add(mBase, uint32(v27)+12)) = v378
	*(*int32)(unsafe.Add(mBase, uint32(v27)+8)) = v380
	*(*int32)(unsafe.Add(mBase, uint32(v27)+4)) = v333
	*(*int32)(unsafe.Add(mBase, uint32(v27))) = v385
	F__serverLog(m, int32(3), int32(_a2225), v27)
	mBase = m.M
	v393 = m.ExcPending
	if v393 != 0 {
		goto L18
	} else {
		goto L59
	}
L59:
	;
	v394 = *(*int32)(unsafe.Add(mBase, uint32(v352)))
	v395 = *(*int32)(unsafe.Add(mBase, uint32(v394)+4))
	v398 = *(*int32)(unsafe.Add(mBase, uint32(v379)))
	v400 = *(*int32)(unsafe.Add(mBase, _consts[28]))
	v401 = int32(base.Ui32(v395) >> (uint(int32(3)) % 32))
	v402 = v398
	v403 = v400
	goto L56
L60:
	;
	v412 = v333 + int32(1)
	if v412 != v273 {
		v333 = v412
		v348 = v403
		goto L49
	} else {
		goto L63
	}
L61:
	;
	F__serverAssert(m, int32(_a2226), int32(_a2157), int32(3924))
	mBase = m.M
	v409 = m.ExcPending
	if v409 != 0 {
		goto L18
	} else {
		goto L62
	}
L62:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L63:
	;
	goto L50
L64:
	;
	F_valkey_free(m, v272)
	mBase = m.M
	v441 = m.ExcPending
	if v441 != 0 {
		goto L18
	} else {
		goto L65
	}
L65:
	;
	goto L37
L66:
	;
	v480 = *(*int32)(unsafe.Add(mBase, _consts[34]))
	goto L70
L67:
	;
	v475 = *(*int32)(unsafe.Add(mBase, uint32(l0)+204))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+204)) = v475 & int32(-2)
	goto L66
L68:
	;
	v492 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+120)) = v492 + base.I32_wrap_i64(v491)
	v496 = int32(0)
	v498 = *(*int64)(unsafe.Add(mBase, _consts[180]))
	v499 = v498 - v58
	v500 = int64(0)
	if v500 < v499 {
		goto L72
	} else {
		goto L73
	}
L69:
	;
	v489 = F_ustime(m)
	mBase = m.M
	v491 = v489 - v59
	goto L68
L70:
	;
	if base.B2i32(v480 != int32(948)) != int32(1) {
		goto L69
	} else {
		goto L71
	}
L71:
	;
	v486 = *(*int32)(unsafe.Add(mBase, _consts[34]))
	v487 = m.T0[v486].(func(*base.Module) int64)(m)
	mBase = m.M
	v491 = v487 - v121
	goto L68
L72:
	;
	v503 = v499
	goto L74
L73:
	;
	v503 = v500
	goto L74
L74:
	;
	v504 = int32(0)
	v505 = *(*int64)(unsafe.Add(mBase, _consts[421]))
	if v31 == v504 {
		goto L77
	} else {
		goto L78
	}
L75:
	;
	F_moduleFireCommandResultEvent(m, l0, v31, v527, v491, v503)
	mBase = m.M
	v529 = m.ExcPending
	if v529 != 0 {
		goto L18
	} else {
		goto L81
	}
L76:
	;
	v527 = int32(1)
	goto L75
L77:
	;
	v517 = int32(0)
	*(*int64)(unsafe.Add(mBase, _consts[422])) = v505
	v519 = *(*int32)(unsafe.Add(mBase, uint32(l0)+344))
	if v519 == v517 {
		v527 = v496
		goto L75
	} else {
		goto L80
	}
L78:
	;
	v509 = *(*int64)(unsafe.Add(mBase, _consts[422]))
	if v505 <= v509 {
		goto L77
	} else {
		goto L79
	}
L79:
	;
	v511 = *(*int64)(unsafe.Add(mBase, uint32(v31)+128))
	*(*int64)(unsafe.Add(mBase, uint32(v31)+128)) = v511 + int64(1)
	*(*int64)(unsafe.Add(mBase, _consts[422])) = v505
	goto L76
L80:
	;
	v522 = *(*int64)(unsafe.Add(mBase, uint32(v31)+128))
	*(*int64)(unsafe.Add(mBase, uint32(v31)+128)) = v522 + int64(1)
	goto L76
L81:
	;
	v530 = *(*int32)(unsafe.Add(mBase, uint32(l0)+204))
	if v530&int32(2048) == int32(0) {
		goto L82
	} else {
		goto L83
	}
L82:
	;
	if v43 != 0 {
		goto L84
	} else {
		goto L85
	}
L83:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+204)) = v530 & int32(-2049)
	v538 = *(*int32)(unsafe.Add(mBase, uint32(l0)+200))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+200)) = v538 | int32(64)
	goto L82
L84:
	;
	v650 = *(*int32)(unsafe.Add(mBase, uint32(l0)+200))
	if v650&int32(16) != 0 {
		goto L114
	} else {
		goto L115
	}
L85:
	;
	v543 = *(*int64)(unsafe.Add(mBase, _consts[45]))
	if v543 == int64(0) {
		goto L86
	} else {
		goto L87
	}
L86:
	;
	v560 = *(*int32)(unsafe.Add(mBase, _consts[177]))
	if v560 != 0 {
		goto L93
	} else {
		goto L94
	}
L87:
	;
	if v491 < v543*int64(1000) {
		goto L86
	} else {
		goto L88
	}
L88:
	;
	v551 = *(*int64)(unsafe.Add(mBase, uint32(v31)+56))
	if v551&int64(16384) == int64(0) {
		goto L89
	} else {
		goto L90
	}
L89:
	;
	v556 = int32(_a622)
	goto L91
L90:
	;
	v556 = int32(_a693)
	goto L91
L91:
	;
	F_latencyAddSample(m, v556, v491)
	mBase = m.M
	v558 = m.ExcPending
	if v558 != 0 {
		goto L18
	} else {
		goto L92
	}
L92:
	;
	goto L86
L93:
	;
	v591 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+200)))
	if v591&int32(16) != 0 {
		goto L99
	} else {
		goto L100
	}
L94:
	;
	goto L97
L95:
	;
	goto L93
L96:
	;
	goto L95
L97:
	;
	v571 = int32(_a2227)
	v572 = *(*int64)(unsafe.Add(mBase, _consts[771]))
	*(*int64)(unsafe.Add(mBase, _consts[771])) = v572 + int64(1)
	v577 = int32(_a2228)
	v578 = *(*int64)(unsafe.Add(mBase, _consts[772]))
	*(*int64)(unsafe.Add(mBase, _consts[772])) = v578 + v491
	v583 = *(*int64)(unsafe.Add(mBase, _consts[773]))
	if base.Ui64(v491) <= base.Ui64(v583) {
		goto L96
	} else {
		goto L98
	}
L98:
	;
	*(*int64)(unsafe.Add(mBase, _consts[773])) = v491
	goto L96
L99:
	;
	if v50&int32(2097152) != 0 {
		goto L102
	} else {
		goto L103
	}
L100:
	;
	F_commandlogPushCurrentCommand(m, l0, v31)
	mBase = m.M
	v595 = m.ExcPending
	if v595 != 0 {
		goto L18
	} else {
		goto L101
	}
L101:
	;
	goto L99
L102:
	;
	v616 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+200)))
	if v616&int32(16) != 0 {
		goto L84
	} else {
		goto L108
	}
L103:
	;
	v596 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
	v597 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v596)+56)))
	if v597&int32(2064) != 0 {
		goto L102
	} else {
		goto L104
	}
L104:
	;
	v601 = *(*int32)(unsafe.Add(mBase, uint32(l0)+192))
	if v601 != 0 {
		v604 = int32(188)
		v605 = v601
		goto L105
	} else {
		goto L106
	}
L105:
	;
	v607 = *(*int32)(unsafe.Add(mBase, _consts[336]))
	v608 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v609 = *(*int32)(unsafe.Add(mBase, uint32(v608)+28))
	v611 = *(*int32)(unsafe.Add(mBase, uint32(l0+v604)))
	F_replicationFeedMonitors(m, l0, v607, v609, v605, v611)
	mBase = m.M
	v613 = m.ExcPending
	if v613 != 0 {
		goto L18
	} else {
		goto L107
	}
L106:
	;
	v602 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v604 = int32(24)
	v605 = v602
	goto L105
L107:
	;
	goto L102
L108:
	;
	v619 = *(*int64)(unsafe.Add(mBase, uint32(v31)+112))
	*(*int64)(unsafe.Add(mBase, uint32(v31)+112)) = v619 + int64(1)
	v623 = *(*int64)(unsafe.Add(mBase, uint32(v31)+104))
	v624 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
	v625 = base.I64_extend_i32_s(v624)
	*(*int64)(unsafe.Add(mBase, uint32(v31)+104)) = v623 + v625
	v628 = int32(0)
	v629 = *(*int32)(unsafe.Add(mBase, _consts[774]))
	if v629 == v628 {
		v643 = v625
		goto L109
	} else {
		goto L110
	}
L109:
	;
	F_clusterSlotStatsAddCpuDuration(m, l0, v643)
	mBase = m.M
	v645 = m.ExcPending
	if v645 != 0 {
		goto L18
	} else {
		goto L113
	}
L110:
	;
	v632 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+200)))
	if v632&int32(16) != 0 {
		v643 = v625
		goto L109
	} else {
		goto L111
	}
L111:
	;
	F_updateCommandLatencyHistogram(m, v31+int32(148), base.I64_extend_i32_s(v624*int32(1000)))
	mBase = m.M
	v641 = m.ExcPending
	if v641 != 0 {
		goto L18
	} else {
		goto L112
	}
L112:
	;
	v642 = int64(*(*int32)(unsafe.Add(mBase, uint32(l0)+120)))
	v643 = v642
	goto L109
L113:
	;
	goto L84
L114:
	;
	v655 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
	if l1&int32(3) == int32(0) {
		v715 = v655
		v716 = v650
		goto L116
	} else {
		goto L117
	}
L115:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+120)) = int32(0)
	goto L114
L116:
	;
	v722 = v716&int32(-2146305) | v32&int32(2146304)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+200)) = v722
	v724 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v715)+56)))
	if v724&int32(2) == int32(0) {
		v749 = v722
		goto L136
	} else {
		goto L137
	}
L117:
	;
	if v650&int32(2097152) != 0 {
		v715 = v655
		v716 = v650
		goto L116
	} else {
		goto L118
	}
L118:
	;
	v662 = *(*int32)(unsafe.Add(mBase, uint32(v655)+48))
	if v662 == int32(17) {
		v715 = v655
		v716 = v650
		goto L116
	} else {
		goto L119
	}
L119:
	;
	v665 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v655)+56)))
	if v665&int32(8) != 0 {
		v715 = v655
		v716 = v650
		goto L116
	} else {
		goto L120
	}
L120:
	;
	v670 = int32(3)
	if v499 < int64(1) {
		goto L121
	} else {
		goto L122
	}
L121:
	;
	v675 = int32(base.Ui32(v650)>>(uint(int32(14))%32)) & v670
	goto L123
L122:
	;
	v675 = v670
	goto L123
L123:
	;
	if v650&int32(1048576) != 0 {
		goto L125
	} else {
		goto L126
	}
L124:
	;
	if v650&int32(524288) != 0 {
		goto L130
	} else {
		goto L131
	}
L125:
	;
	v689 = v675 & int32(1)
	goto L124
L126:
	;
	if l1&int32(2) == int32(0) {
		goto L125
	} else {
		goto L127
	}
L127:
	;
	v682 = *(*int32)(unsafe.Add(mBase, uint32(l0)+204))
	if v682&int32(1048576) == int32(0) {
		v689 = v675
		goto L124
	} else {
		goto L128
	}
L128:
	;
	goto L125
L129:
	;
	if v703 == int32(0) {
		v715 = v655
		v716 = v650
		goto L116
	} else {
		goto L134
	}
L130:
	;
	v703 = v689 & int32(2)
	goto L129
L131:
	;
	if l1&int32(1) == int32(0) {
		goto L130
	} else {
		goto L132
	}
L132:
	;
	v696 = *(*int32)(unsafe.Add(mBase, uint32(l0)+204))
	if v696&int32(524288) == int32(0) {
		v703 = v689
		goto L129
	} else {
		goto L133
	}
L133:
	;
	goto L130
L134:
	;
	v706 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v707 = *(*int32)(unsafe.Add(mBase, uint32(v706)+28))
	v708 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v709 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v710 = *(*int32)(unsafe.Add(mBase, uint32(l0)+292))
	F_alsoPropagate(m, v707, v708, v709, v703, v710)
	mBase = m.M
	v712 = m.ExcPending
	if v712 != 0 {
		goto L18
	} else {
		goto L135
	}
L135:
	;
	v713 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
	v714 = *(*int32)(unsafe.Add(mBase, uint32(l0)+200))
	v715 = v713
	v716 = v714
	goto L116
L136:
	;
	if v749&int32(16) != 0 {
		goto L144
	} else {
		goto L145
	}
L137:
	;
	v729 = *(*int32)(unsafe.Add(mBase, uint32(v715)+48))
	if v729 == int32(294) {
		v749 = v722
		goto L136
	} else {
		goto L138
	}
L138:
	;
	if v729 == int32(293) {
		v749 = v722
		goto L136
	} else {
		goto L139
	}
L139:
	;
	if v729 == int32(297) {
		v749 = v722
		goto L136
	} else {
		goto L140
	}
L140:
	;
	v736 = int32(0)
	v737 = *(*int32)(unsafe.Add(mBase, _consts[14]))
	if v737 == v736 {
		v749 = v722
		goto L136
	} else {
		goto L141
	}
L141:
	;
	v740 = *(*int32)(unsafe.Add(mBase, uint32(v737)+204))
	if v740&int32(20) != int32(4) {
		v749 = v722
		goto L136
	} else {
		goto L142
	}
L142:
	;
	F_trackingRememberKeys(m, v737, l0)
	mBase = m.M
	v746 = m.ExcPending
	if v746 != 0 {
		goto L18
	} else {
		goto L143
	}
L143:
	;
	v747 = *(*int32)(unsafe.Add(mBase, uint32(l0)+200))
	v749 = v747
	goto L136
L144:
	;
	v767 = int32(0)
	v776 = *(*int32)(unsafe.Add(mBase, _consts[315]))
	if v776 < int32(261) {
		goto L152
	} else {
		goto L153
	}
L145:
	;
	v752 = int32(0)
	v753 = *(*int32)(unsafe.Add(mBase, _consts[14]))
	if v753 == v752 {
		goto L146
	} else {
		goto L147
	}
L146:
	;
	v760 = int32(0)
	v762 = *(*int64)(unsafe.Add(mBase, _consts[730]))
	*(*int64)(unsafe.Add(mBase, _consts[730])) = v762 + int64(1)
	goto L144
L147:
	;
	v756 = *(*int64)(unsafe.Add(mBase, uint32(v753)+256))
	*(*int64)(unsafe.Add(mBase, uint32(v753)+256)) = v756 + int64(1)
	goto L146
L148:
	;
	F_postExecutionUnitOperations(m)
	mBase = m.M
	v866 = m.ExcPending
	if v866 != 0 {
		goto L18
	} else {
		goto L166
	}
L149:
	;
	v861 = *(*int32)(unsafe.Add(mBase, _consts[554]))
	if base.Ui32(v853) <= base.Ui32(v861) {
		goto L148
	} else {
		goto L165
	}
L150:
	;
	goto L149
L151:
	;
	v787 = v785 & int32(3)
	if base.Ui32(int32(4)) <= base.Ui32(v785) {
		goto L156
	} else {
		goto L157
	}
L152:
	;
	if v776 < int32(1) {
		v853 = v767
		goto L150
	} else {
		goto L154
	}
L153:
	;
	v780 = *(*int32)(unsafe.Add(mBase, _consts[316]))
	v784 = v780
	v785 = int32(260)
	goto L151
L154:
	;
	v784 = v767
	v785 = v776
	goto L151
L155:
	;
	if v787 == int32(0) {
		v853 = v826
		goto L150
	} else {
		goto L161
	}
L156:
	;
	v794 = int32(0)
	v796 = v784
	v797 = v794
	v801 = v794
	goto L158
L157:
	;
	v826 = v784
	v827 = int32(0)
	goto L155
L158:
	;
	v804 = v797 << (uint(int32(2)) % 32)
	v807 = *(*int32)(unsafe.Add(mBase, uint32(v804)+uint32(_consts[317])))
	v810 = *(*int32)(unsafe.Add(mBase, uint32(v804)+uint32(_consts[318])))
	v813 = *(*int32)(unsafe.Add(mBase, uint32(v804)+uint32(_consts[319])))
	v816 = *(*int32)(unsafe.Add(mBase, uint32(v804)+uint32(_consts[320])))
	v820 = v807 + (v810 + (v813 + (v816 + v796)))
	v821 = int32(4)
	v822 = v797 + v821
	v824 = v801 + v821
	if v824 != v785&int32(2147483644) {
		v796 = v820
		v797 = v822
		v801 = v824
		goto L158
	} else {
		goto L160
	}
L159:
	;
	v826 = v820
	v827 = v822
	goto L155
L160:
	;
	goto L159
L161:
	;
	v835 = v826
	v836 = v827
	v838 = int32(0)
	goto L162
L162:
	;
	v846 = *(*int32)(unsafe.Add(mBase, uint32(v836<<(uint(int32(2))%32))+uint32(_consts[320])))
	v847 = v846 + v835
	v848 = int32(1)
	v851 = v838 + v848
	if v851 != v787 {
		v835 = v847
		v836 = v836 + v848
		v838 = v851
		goto L162
	} else {
		goto L164
	}
L163:
	;
	v853 = v847
	goto L150
L164:
	;
	goto L163
L165:
	;
	*(*int32)(unsafe.Add(mBase, _consts[554])) = v853
	goto L148
L166:
	;
	F_trackingHandlePendingKeyInvalidations(m)
	mBase = m.M
	v868 = m.ExcPending
	if v868 != 0 {
		goto L18
	} else {
		goto L167
	}
L167:
	;
	F_clusterSlotStatsAddNetworkBytesOutForUserClient(m, l0)
	mBase = m.M
	v870 = m.ExcPending
	if v870 != 0 {
		goto L18
	} else {
		goto L168
	}
L168:
	;
	v872 = *(*int32)(unsafe.Add(mBase, _consts[177]))
	if v872 != 0 {
		goto L169
	} else {
		goto L170
	}
L169:
	;
	v901 = *(*int64)(unsafe.Add(mBase, _consts[47]))
	if v56 == v901 {
		goto L177
	} else {
		goto L178
	}
L170:
	;
	v873 = *(*int32)(unsafe.Add(mBase, uint32(l0)+132))
	v874 = int32(0)
	v875 = *(*int32)(unsafe.Add(mBase, _consts[485]))
	v879 = *(*int32)(unsafe.Add(mBase, uint32(v875)+20))
	if v879 == v874 {
		goto L172
	} else {
		goto L173
	}
L171:
	;
	goto L169
L172:
	;
	goto L171
L173:
	;
	v882 = *(*int32)(unsafe.Add(mBase, uint32(v875)))
	v883 = *(*int32)(unsafe.Add(mBase, uint32(v873)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v882))) = v883
	if v883 == int32(0) {
		goto L175
	} else {
		goto L176
	}
L174:
	;
	v889 = *(*int32)(unsafe.Add(mBase, uint32(v875)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v873)+4)) = v889
	v891 = *(*int32)(unsafe.Add(mBase, uint32(v873)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v873)+20)) = v891 + v879
	*(*int32)(unsafe.Add(mBase, uint32(v875)+20)) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v875))) = int64(0)
	goto L172
L175:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v873))) = v882
	goto L174
L176:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v883)+4)) = v882
	goto L174
L177:
	;
	v905 = *(*int32)(unsafe.Add(mBase, _consts[62]))
	if v905 != 0 {
		goto L179
	} else {
		goto L180
	}
L178:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l0)+48)) = v901
	goto L177
L179:
	;
	*(*int32)(unsafe.Add(mBase, _consts[300])) = v30
	m.G0 = v27 + int32(48)
	return
L180:
	;
	v906 = int32(0)
	v907 = *(*int32)(unsafe.Add(mBase, _consts[775]))
	if v907 == v906 {
		goto L179
	} else {
		goto L181
	}
L181:
	;
	v910 = int32(0)
	*(*int32)(unsafe.Add(mBase, _consts[775])) = v910
	goto L179
L182:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L183:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_callArgvOnAvailableCallback(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
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
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v21 int32
	_ = v21
	v5 = int32(0)
	v6 = m.G6
	v7 = *(*int32)(unsafe.Add(mBase, uint32(v6)))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(v6)+264))
	if base.B2i32(v7 != v5)&base.B2i32(v10 != v5) == int32(0) {
		return int32(1)
	} else {
		v16 = m.G19
		v17 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
		m.T0[v17].(func(*base.Module, int32))(m, l2)
		mBase = m.M
		v21 = m.ExcPending
		if v21 != 0 {
			return int32(0)
		} else {
			return int32(1)
		}
	}
}
func F_call_orderTM(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
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
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v74 float64
	_ = v74
	var v75 float64
	_ = v75
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v93 int64
	_ = v93
	var v95 int32
	_ = v95
	var v97 int32
	_ = v97
	var v98 int64
	_ = v98
	var v100 int32
	_ = v100
	var v102 int32
	_ = v102
	var v103 int64
	_ = v103
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v127 int32
	_ = v127
	var v129 int32
	_ = v129
	var v131 int32
	_ = v131
	var v132 int64
	_ = v132
	var v136 int32
	_ = v136
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v149 int32
	_ = v149
	v7 = int32(-1)
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	switch v9 + int32(-5) {
	case 0:
		v12 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
		v24 = v12 + int32(16)
	default:
		v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
		v24 = v18 + v9<<(uint(int32(2))%32) + int32(152)
	case 2:
		v15 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
		v24 = v15 + int32(8)
	}
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v24)))
	v26 = m.G398
	if v25 == int32(0) {
		v35 = v26
	} else {
		v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
		v33 = *(*int32)(unsafe.Add(mBase, uint32(v29+l3<<(uint(int32(2))%32))+188))
		v34 = F_luaH_getstr(m, v25, v33)
		mBase = m.M
		v35 = v34
	}
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v35)+8))
	if v36 == int32(0) {
		v149 = v7
		return v149
	} else {
		v40 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
		switch v40 + int32(-5) {
		case 0:
			v43 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
			v55 = v43 + int32(16)
		default:
			v49 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
			v55 = v49 + v40<<(uint(int32(2))%32) + int32(152)
		case 2:
			v46 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
			v55 = v46 + int32(8)
		}
		v56 = *(*int32)(unsafe.Add(mBase, uint32(v55)))
		v57 = m.G398
		if v56 == int32(0) {
			v66 = v57
		} else {
			v60 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
			v64 = *(*int32)(unsafe.Add(mBase, uint32(v60+l3<<(uint(int32(2))%32))+188))
			v65 = F_luaH_getstr(m, v56, v64)
			mBase = m.M
			v66 = v65
		}
		v69 = *(*int32)(unsafe.Add(mBase, uint32(v35)+8))
		v70 = *(*int32)(unsafe.Add(mBase, uint32(v66)+8))
		if v69 == v70 {
			switch v69 {
			case 0:
				v86 = int32(1)
				v88 = v86
			case 1:
				v77 = *(*int32)(unsafe.Add(mBase, uint32(v35)))
				v78 = *(*int32)(unsafe.Add(mBase, uint32(v66)))
				v88 = base.B2i32(v77 == v78)
			case 2:
				v80 = *(*int32)(unsafe.Add(mBase, uint32(v35)))
				v81 = *(*int32)(unsafe.Add(mBase, uint32(v66)))
				v88 = base.B2i32(v80 == v81)
			case 3:
				v74 = *(*float64)(unsafe.Add(mBase, uint32(v35)))
				v75 = *(*float64)(unsafe.Add(mBase, uint32(v66)))
				v88 = base.F64_eq(v74, v75)
			default:
				v83 = *(*int32)(unsafe.Add(mBase, uint32(v35)))
				v84 = *(*int32)(unsafe.Add(mBase, uint32(v66)))
				v86 = base.B2i32(v83 == v84)
				v88 = v86
			}
		} else {
			v88 = int32(0)
		}
		if v88 == int32(0) {
			v149 = v7
			return v149
		} else {
			v91 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
			v92 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			v93 = *(*int64)(unsafe.Add(mBase, uint32(v35)))
			*(*int64)(unsafe.Add(mBase, uint32(v92))) = v93
			v95 = *(*int32)(unsafe.Add(mBase, uint32(v35)+8))
			*(*int32)(unsafe.Add(mBase, uint32(v92)+8)) = v95
			v97 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			v98 = *(*int64)(unsafe.Add(mBase, uint32(l1)))
			*(*int64)(unsafe.Add(mBase, uint32(v97)+16)) = v98
			v100 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
			*(*int32)(unsafe.Add(mBase, uint32(v97)+24)) = v100
			v102 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			v103 = *(*int64)(unsafe.Add(mBase, uint32(l2)))
			*(*int64)(unsafe.Add(mBase, uint32(v102)+32)) = v103
			v105 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
			*(*int32)(unsafe.Add(mBase, uint32(v102)+40)) = v105
			v107 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
			v108 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			if int32(48) < v107-v108 {
				v118 = v108
				*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v118 + int32(48)
				F_luaD_call(m, l0, v118, int32(1))
				mBase = m.M
				v124 = m.ExcPending
				if v124 != 0 {
					return int32(0)
				} else {
					v125 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
					v127 = v125 + int32(-16)
					*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v127
					v129 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
					v131 = v129 + (v92 - v91)
					v132 = *(*int64)(unsafe.Add(mBase, uint32(v127)))
					*(*int64)(unsafe.Add(mBase, uint32(v131))) = v132
					v136 = *(*int32)(unsafe.Add(mBase, uint32(v125+int32(-8))))
					*(*int32)(unsafe.Add(mBase, uint32(v131)+8)) = v136
					v138 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
					v139 = *(*int32)(unsafe.Add(mBase, uint32(v138)+8))
					switch v139 {
					case 0:
						v149 = v139
						return v149
					case 1:
						v140 = *(*int32)(unsafe.Add(mBase, uint32(v138)))
						return base.B2i32(v140 != int32(0))
					default:
						v149 = int32(1)
						return v149
					}
				}
			} else {
				F_luaD_growstack(m, l0, int32(3))
				mBase = m.M
				v116 = m.ExcPending
				if v116 != 0 {
					return int32(0)
				} else {
					v117 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
					v118 = v117
					*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v118 + int32(48)
					F_luaD_call(m, l0, v118, int32(1))
					mBase = m.M
					v124 = m.ExcPending
					if v124 != 0 {
						return int32(0)
					} else {
						v125 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
						v127 = v125 + int32(-16)
						*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v127
						v129 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
						v131 = v129 + (v92 - v91)
						v132 = *(*int64)(unsafe.Add(mBase, uint32(v127)))
						*(*int64)(unsafe.Add(mBase, uint32(v131))) = v132
						v136 = *(*int32)(unsafe.Add(mBase, uint32(v125+int32(-8))))
						*(*int32)(unsafe.Add(mBase, uint32(v131)+8)) = v136
						v138 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
						v139 = *(*int32)(unsafe.Add(mBase, uint32(v138)+8))
						switch v139 {
						case 0:
							v149 = v139
							return v149
						case 1:
							v140 = *(*int32)(unsafe.Add(mBase, uint32(v138)))
							return base.B2i32(v140 != int32(0))
						default:
							v149 = int32(1)
							return v149
						}
					}
				}
			}
		}
	}
}
func F_callallgcTM(m *base.Module, l0 int32, l1 int32) {
	var v4 int32
	_ = v4
	F_luaC_callGCTM(m, l0)
	v4 = m.ExcPending
	if v4 != 0 {
		return
	} else {
		return
	}
}
func F_cancelReplicationHandshake(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v14 int32
	_ = v14
	var v18 int32
	_ = v18
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v29 int64
	_ = v29
	var v32 int32
	_ = v32
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v61 int32
	_ = v61
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v80 int32
	_ = v80
	var v87 int32
	_ = v87
	var v91 int64
	_ = v91
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v14 = *(*int32)(unsafe.Add(mBase, _consts[587]))
	goto L2
L1:
	;
	v28 = int32(_a20)
	v29 = int64(0)
	*(*int64)(unsafe.Add(mBase, _consts[588])) = v29
	v32 = int32(0)
	*(*int32)(unsafe.Add(mBase, _consts[589])) = v32
	*(*int64)(unsafe.Add(mBase, _consts[590])) = v29
	v40 = *(*int32)(unsafe.Add(mBase, _consts[591]))
	if v40 == v32 {
		goto L6
	} else {
		goto L7
	}
L2:
	;
	if v14 == int32(0) {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v18 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _consts[593])) = uint8(v18)
	F_bioDrainWorker(m, int32(4))
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	return int32(0)
L5:
	;
	v26 = int32(0)
	*(*uint8)(unsafe.Add(mBase, _consts[593])) = uint8(v26)
	goto L1
L6:
	;
	v46 = *(*int32)(unsafe.Add(mBase, _consts[423]))
	switch v46 + int32(-2) {
	case 0:
		goto L11
	default:
		goto L12
	case 11:
		goto L13
	}
L7:
	;
	F_replicationAbortDualChannelSyncTransfer(m)
	mBase = m.M
	v44 = m.ExcPending
	if v44 != 0 {
		goto L4
	} else {
		goto L8
	}
L8:
	;
	goto L6
L9:
	;
	m.G0 = v7 + int32(16)
	return v99
L10:
	;
	v80 = int32(1)
	*(*int32)(unsafe.Add(mBase, _consts[423])) = v80
	if l0 == int32(0) {
		v99 = v80
		goto L9
	} else {
		goto L21
	}
L11:
	;
	v69 = *(*int32)(unsafe.Add(mBase, _consts[592]))
	if v69 == int32(0) {
		goto L10
	} else {
		goto L19
	}
L12:
	;
	if base.Ui32(v46+int32(-13)) < base.Ui32(int32(-10)) {
		v99 = int32(0)
		goto L9
	} else {
		goto L18
	}
L13:
	;
	v50 = *(*int32)(unsafe.Add(mBase, _consts[592]))
	if v50 == int32(0) {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	F_cleanupTransferResources(m)
	mBase = m.M
	v61 = m.ExcPending
	if v61 != 0 {
		goto L4
	} else {
		goto L17
	}
L15:
	;
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v50)))
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v53)+52))
	m.T0[v54].(func(*base.Module, int32))(m, v50)
	mBase = m.M
	v56 = m.ExcPending
	if v56 != 0 {
		goto L4
	} else {
		goto L16
	}
L16:
	;
	*(*int32)(unsafe.Add(mBase, _consts[592])) = int32(0)
	goto L14
L17:
	;
	goto L10
L18:
	;
	goto L11
L19:
	;
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v69)))
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v72)+52))
	m.T0[v73].(func(*base.Module, int32))(m, v69)
	mBase = m.M
	v75 = m.ExcPending
	if v75 != 0 {
		goto L4
	} else {
		goto L20
	}
L20:
	;
	*(*int32)(unsafe.Add(mBase, _consts[592])) = int32(0)
	goto L10
L21:
	;
	v87 = *(*int32)(unsafe.Add(mBase, _consts[28]))
	if int32(2) < v87 {
		goto L22
	} else {
		goto L23
	}
L22:
	;
	v97 = F_connectWithPrimary(m)
	mBase = m.M
	v98 = m.ExcPending
	if v98 != 0 {
		goto L4
	} else {
		goto L25
	}
L23:
	;
	v91 = *(*int64)(unsafe.Add(mBase, _consts[166]))
	*(*int64)(unsafe.Add(mBase, uint32(v7))) = v91
	F__serverLog(m, int32(2), int32(_a1940), v7)
	mBase = m.M
	v96 = m.ExcPending
	if v96 != 0 {
		goto L4
	} else {
		goto L24
	}
L24:
	;
	goto L22
L25:
	;
	v99 = v80
	goto L9
}
func F_catClientInfoShortString(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
	var v30 int32
	_ = v30
	var v34 int32
	_ = v34
	var v38 int64
	_ = v38
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
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
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	v12 = m.G0
	v14 = v12 - int32(96)
	m.G0 = v14
	v17 = *(*int32)(unsafe.Add(mBase, _consts[299]))
	if v17 != 0 {
	} else {
		v18 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+222)))
		if v18 != 0 {
			v22 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+222)))
			if v22 != int32(1) {
			} else {
				for {
					v26 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+222)))
					if v26 == int32(1) {
						continue
					} else {
						break
					}
					break
				}
			}
			v30 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+223)))
			if v30 != int32(1) {
			} else {
				for {
					v34 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+223)))
					if v34 == int32(1) {
						continue
					} else {
						break
					}
					break
				}
			}
		} else {
			v19 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+223)))
			if v19 == int32(0) {
			} else {
				v22 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+222)))
				if v22 != int32(1) {
				} else {
					for {
						v26 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+222)))
						if v26 == int32(1) {
							continue
						} else {
							break
						}
						break
					}
				}
				v30 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+223)))
				if v30 != int32(1) {
				} else {
					for {
						v34 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+223)))
						if v34 == int32(1) {
							continue
						} else {
							break
						}
						break
					}
				}
			}
		}
	}
	v38 = *(*int64)(unsafe.Add(mBase, uint32(l1)))
	v39 = F_getClientPeerId(m, l1)
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		return int32(0)
	} else {
		v43 = F_getClientSockname(m, l1)
		mBase = m.M
		v44 = m.ExcPending
		if v44 != 0 {
			return int32(0)
		} else {
			v45 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
			if v45 != 0 {
				v47 = *(*int32)(unsafe.Add(mBase, uint32(v45)+12))
				v48 = v47
			} else {
				v48 = int32(-1)
			}
			*(*int32)(unsafe.Add(mBase, uint32(v14)+48)) = v48
			v56 = F_snprintf(m, v14+int32(64), int32(31), int32(_a1644), v14+int32(48))
			mBase = m.M
			v57 = m.ExcPending
			if v57 != 0 {
				return int32(0)
			} else {
				if l2 == int32(0) {
					v62 = *(*int32)(unsafe.Add(mBase, uint32(l1)+348))
					if v62 != 0 {
						v64 = F_objectGetVal(m, v62)
						mBase = m.M
						v65 = v64
					} else {
						v65 = int32(_a320)
					}
					v66 = *(*int32)(unsafe.Add(mBase, uint32(l1)+328))
					if v66 != 0 {
						v68 = *(*int32)(unsafe.Add(mBase, uint32(v66)))
						v70 = v65
						v71 = v68
					} else {
						v70 = v65
						v71 = int32(_a1647)
					}
				} else {
					v60 = int32(_a1645)
					v70 = v60
					v71 = v60
				}
				v72 = int32(_a320)
				v74 = *(*int32)(unsafe.Add(mBase, uint32(l1)+352))
				if v74 == int32(0) {
					v78 = v72
				} else {
					v77 = F_objectGetVal(m, v74)
					mBase = m.M
					v78 = v77
				}
				v79 = *(*int32)(unsafe.Add(mBase, uint32(l1)+356))
				if v79 == int32(0) {
					v83 = v72
				} else {
					v82 = F_objectGetVal(m, v79)
					mBase = m.M
					v83 = v82
				}
				*(*int32)(unsafe.Add(mBase, uint32(v14+int32(32)))) = v83
				*(*int32)(unsafe.Add(mBase, uint32(v14+int32(28)))) = v78
				*(*int32)(unsafe.Add(mBase, uint32(v14+int32(24)))) = v71
				*(*int32)(unsafe.Add(mBase, uint32(v14+int32(20)))) = v70
				*(*int32)(unsafe.Add(mBase, uint32(v14)+12)) = v43
				*(*int32)(unsafe.Add(mBase, uint32(v14)+8)) = v39
				*(*int64)(unsafe.Add(mBase, uint32(v14))) = v38
				*(*int32)(unsafe.Add(mBase, uint32(v14+int32(16)))) = v14 + int32(64)
				v105 = F_sdscatfmt(m, l0, int32(_a1695), v14)
				mBase = m.M
				v106 = m.ExcPending
				if v106 != 0 {
					return int32(0)
				} else {
					m.G0 = v14 + int32(96)
					return v105
				}
			}
		}
	}
}
func F_catClientInfoString(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v38 int32
	_ = v38
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
	var v52 int32
	_ = v52
	var v56 int32
	_ = v56
	var v60 int32
	_ = v60
	var v66 int32
	_ = v66
	var v75 int32
	_ = v75
	var v81 int32
	_ = v81
	var v86 int32
	_ = v86
	var v90 int32
	_ = v90
	var v95 int32
	_ = v95
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v106 int32
	_ = v106
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v117 int32
	_ = v117
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v127 int32
	_ = v127
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v138 int32
	_ = v138
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v149 int32
	_ = v149
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v159 int32
	_ = v159
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v170 int32
	_ = v170
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v181 int32
	_ = v181
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v192 int32
	_ = v192
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v203 int32
	_ = v203
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v214 int32
	_ = v214
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v224 int32
	_ = v224
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v235 int32
	_ = v235
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v246 int32
	_ = v246
	var v250 int32
	_ = v250
	var v251 int32
	_ = v251
	var v254 int32
	_ = v254
	var v259 int32
	_ = v259
	var v263 int32
	_ = v263
	var v264 int32
	_ = v264
	var v267 int32
	_ = v267
	var v270 int32
	_ = v270
	var v274 int32
	_ = v274
	var v279 int32
	_ = v279
	var v283 int32
	_ = v283
	var v284 int32
	_ = v284
	var v286 int32
	_ = v286
	var v289 int32
	_ = v289
	var v292 int32
	_ = v292
	var v296 int32
	_ = v296
	var v297 int32
	_ = v297
	var v300 int32
	_ = v300
	var v304 int32
	_ = v304
	var v305 int32
	_ = v305
	var v309 int32
	_ = v309
	var v314 int32
	_ = v314
	var v320 int32
	_ = v320
	var v321 int32
	_ = v321
	var v325 int32
	_ = v325
	var v328 int32
	_ = v328
	var v329 int64
	_ = v329
	var v330 int32
	_ = v330
	var v333 int32
	_ = v333
	var v337 int32
	_ = v337
	var v338 int32
	_ = v338
	var v339 int32
	_ = v339
	var v340 int64
	_ = v340
	var v341 int32
	_ = v341
	var v342 int64
	_ = v342
	var v349 int64
	_ = v349
	var v350 int64
	_ = v350
	var v351 int32
	_ = v351
	var v352 int32
	_ = v352
	var v353 int32
	_ = v353
	var v354 int32
	_ = v354
	var v355 int32
	_ = v355
	var v357 int32
	_ = v357
	var v358 int32
	_ = v358
	var v366 int32
	_ = v366
	var v367 int32
	_ = v367
	var v371 int32
	_ = v371
	var v373 int32
	_ = v373
	var v375 int32
	_ = v375
	var v377 int64
	_ = v377
	var v378 int64
	_ = v378
	var v380 int64
	_ = v380
	var v381 int64
	_ = v381
	var v382 int32
	_ = v382
	var v383 int32
	_ = v383
	var v384 int32
	_ = v384
	var v388 int32
	_ = v388
	var v391 int32
	_ = v391
	var v392 int32
	_ = v392
	var v393 int32
	_ = v393
	var v394 int32
	_ = v394
	var v395 int32
	_ = v395
	var v396 int32
	_ = v396
	var v398 int32
	_ = v398
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
	var v406 int32
	_ = v406
	var v409 int32
	_ = v409
	var v410 int32
	_ = v410
	var v411 int32
	_ = v411
	var v412 int32
	_ = v412
	var v414 int32
	_ = v414
	var v415 int32
	_ = v415
	var v416 int32
	_ = v416
	var v417 int32
	_ = v417
	var v418 int64
	_ = v418
	var v421 int32
	_ = v421
	var v424 int32
	_ = v424
	var v428 int32
	_ = v428
	var v436 int32
	_ = v436
	var v439 int32
	_ = v439
	var v444 int32
	_ = v444
	var v447 int32
	_ = v447
	var v452 int32
	_ = v452
	var v455 int32
	_ = v455
	var v460 int64
	_ = v460
	var v463 int64
	_ = v463
	var v468 int64
	_ = v468
	var v470 int32
	_ = v470
	var v472 int64
	_ = v472
	var v473 int64
	_ = v473
	var v479 int64
	_ = v479
	var v480 int64
	_ = v480
	var v481 int32
	_ = v481
	var v483 int32
	_ = v483
	var v484 int32
	_ = v484
	var v488 int32
	_ = v488
	var v490 int32
	_ = v490
	var v492 int32
	_ = v492
	var v493 int32
	_ = v493
	var v495 int32
	_ = v495
	var v500 int32
	_ = v500
	var v501 int64
	_ = v501
	var v502 int64
	_ = v502
	var v503 int64
	_ = v503
	var v504 int64
	_ = v504
	var v505 int64
	_ = v505
	var v506 int64
	_ = v506
	var v507 int64
	_ = v507
	var v508 int64
	_ = v508
	var v509 int32
	_ = v509
	var v510 int32
	_ = v510
	var v512 int32
	_ = v512
	var v515 int32
	_ = v515
	var v516 int32
	_ = v516
	var v517 int32
	_ = v517
	var v520 int32
	_ = v520
	var v521 int32
	_ = v521
	var v522 int64
	_ = v522
	var v523 int64
	_ = v523
	var v526 int64
	_ = v526
	var v612 int64
	_ = v612
	var v639 int32
	_ = v639
	var v640 int32
	_ = v640
	v38 = m.G0
	v40 = v38 - int32(320)
	m.G0 = v40
	v43 = *(*int32)(unsafe.Add(mBase, _consts[299]))
	if v43 != 0 {
	} else {
		v44 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+222)))
		if v44 != 0 {
			v48 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+222)))
			if v48 != int32(1) {
			} else {
				for {
					v52 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+222)))
					if v52 == int32(1) {
						continue
					} else {
						break
					}
					break
				}
			}
			v56 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+223)))
			if v56 != int32(1) {
			} else {
				for {
					v60 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+223)))
					if v60 == int32(1) {
						continue
					} else {
						break
					}
					break
				}
			}
		} else {
			v45 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+223)))
			if v45 == int32(0) {
			} else {
				v48 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+222)))
				if v48 != int32(1) {
				} else {
					for {
						v52 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+222)))
						if v52 == int32(1) {
							continue
						} else {
							break
						}
						break
					}
				}
				v56 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+223)))
				if v56 != int32(1) {
				} else {
					for {
						v60 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+223)))
						if v60 == int32(1) {
							continue
						} else {
							break
						}
						break
					}
				}
			}
		}
	}
	v66 = *(*int32)(unsafe.Add(mBase, uint32(l1)+200))
	if v66&int32(2) == int32(0) {
		v81 = v40 + int32(288)
	} else {
		if v66&int32(4) != 0 {
			v75 = int32(79)
		} else {
			v75 = int32(83)
		}
		*(*uint8)(unsafe.Add(mBase, uint32(v40)+288)) = uint8(v75)
		v81 = v40 + int32(288) | int32(1)
	}
	if v66&int32(1) == int32(0) {
		v90 = v81
	} else {
		v86 = int32(77)
		*(*uint8)(unsafe.Add(mBase, uint32(v81))) = uint8(v86)
		v90 = v81 + int32(1)
	}
	if v66&int32(262144) == int32(0) {
		v100 = v90
		v101 = v66
	} else {
		v95 = int32(80)
		*(*uint8)(unsafe.Add(mBase, uint32(v90))) = uint8(v95)
		v99 = *(*int32)(unsafe.Add(mBase, uint32(l1)+200))
		v100 = v90 + int32(1)
		v101 = v99
	}
	if v101&int32(8) == int32(0) {
		v111 = v100
		v112 = v101
	} else {
		v106 = int32(120)
		*(*uint8)(unsafe.Add(mBase, uint32(v100))) = uint8(v106)
		v110 = *(*int32)(unsafe.Add(mBase, uint32(l1)+200))
		v111 = v100 + int32(1)
		v112 = v110
	}
	if v112&int32(16) == int32(0) {
		v121 = v111
	} else {
		v117 = int32(98)
		*(*uint8)(unsafe.Add(mBase, uint32(v111))) = uint8(v117)
		v121 = v111 + int32(1)
	}
	v122 = *(*int32)(unsafe.Add(mBase, uint32(l1)+204))
	if v122&int32(4) == int32(0) {
		v132 = v121
		v133 = v122
	} else {
		v127 = int32(116)
		*(*uint8)(unsafe.Add(mBase, uint32(v121))) = uint8(v127)
		v131 = *(*int32)(unsafe.Add(mBase, uint32(l1)+204))
		v132 = v121 + int32(1)
		v133 = v131
	}
	if v133&int32(8) == int32(0) {
		v143 = v132
		v144 = v133
	} else {
		v138 = int32(82)
		*(*uint8)(unsafe.Add(mBase, uint32(v132))) = uint8(v138)
		v142 = *(*int32)(unsafe.Add(mBase, uint32(l1)+204))
		v143 = v132 + int32(1)
		v144 = v142
	}
	if v144&int32(16) == int32(0) {
		v153 = v143
	} else {
		v149 = int32(66)
		*(*uint8)(unsafe.Add(mBase, uint32(v143))) = uint8(v149)
		v153 = v143 + int32(1)
	}
	v154 = *(*int32)(unsafe.Add(mBase, uint32(l1)+200))
	if v154&int32(32) == int32(0) {
		v164 = v153
		v165 = v154
	} else {
		v159 = int32(100)
		*(*uint8)(unsafe.Add(mBase, uint32(v153))) = uint8(v159)
		v163 = *(*int32)(unsafe.Add(mBase, uint32(l1)+200))
		v164 = v153 + int32(1)
		v165 = v163
	}
	if v165&int32(64) == int32(0) {
		v175 = v164
		v176 = v165
	} else {
		v170 = int32(99)
		*(*uint8)(unsafe.Add(mBase, uint32(v164))) = uint8(v170)
		v174 = *(*int32)(unsafe.Add(mBase, uint32(l1)+200))
		v175 = v164 + int32(1)
		v176 = v174
	}
	if v176&int32(128) == int32(0) {
		v186 = v175
		v187 = v176
	} else {
		v181 = int32(117)
		*(*uint8)(unsafe.Add(mBase, uint32(v175))) = uint8(v181)
		v185 = *(*int32)(unsafe.Add(mBase, uint32(l1)+200))
		v186 = v175 + int32(1)
		v187 = v185
	}
	if v187&int32(1024) == int32(0) {
		v197 = v186
		v198 = v187
	} else {
		v192 = int32(65)
		*(*uint8)(unsafe.Add(mBase, uint32(v186))) = uint8(v192)
		v196 = *(*int32)(unsafe.Add(mBase, uint32(l1)+200))
		v197 = v186 + int32(1)
		v198 = v196
	}
	if v198&int32(2048) == int32(0) {
		v208 = v197
		v209 = v198
	} else {
		v203 = int32(85)
		*(*uint8)(unsafe.Add(mBase, uint32(v197))) = uint8(v203)
		v207 = *(*int32)(unsafe.Add(mBase, uint32(l1)+200))
		v208 = v197 + int32(1)
		v209 = v207
	}
	if v209&int32(131072) == int32(0) {
		v218 = v208
	} else {
		v214 = int32(114)
		*(*uint8)(unsafe.Add(mBase, uint32(v208))) = uint8(v214)
		v218 = v208 + int32(1)
	}
	v219 = *(*int32)(unsafe.Add(mBase, uint32(l1)+204))
	if v219&int32(16384) == int32(0) {
		v229 = v218
		v230 = v219
	} else {
		v224 = int32(101)
		*(*uint8)(unsafe.Add(mBase, uint32(v218))) = uint8(v224)
		v228 = *(*int32)(unsafe.Add(mBase, uint32(l1)+204))
		v229 = v218 + int32(1)
		v230 = v228
	}
	if v230&int32(65536) == int32(0) {
		v240 = v229
		v241 = v230
	} else {
		v235 = int32(84)
		*(*uint8)(unsafe.Add(mBase, uint32(v229))) = uint8(v235)
		v239 = *(*int32)(unsafe.Add(mBase, uint32(l1)+204))
		v240 = v229 + int32(1)
		v241 = v239
	}
	if v241&int32(536870912) == int32(0) {
		v250 = v240
	} else {
		v246 = int32(73)
		*(*uint8)(unsafe.Add(mBase, uint32(v240))) = uint8(v246)
		v250 = v240 + int32(1)
	}
	v251 = *(*int32)(unsafe.Add(mBase, uint32(l1)+216))
	if v251 == int32(0) {
		v274 = v250
	} else {
		v254 = *(*int32)(unsafe.Add(mBase, uint32(v251)))
		if base.B2i32(v254 == int32(1)) == int32(0) {
			v263 = v250
		} else {
			v259 = int32(105)
			*(*uint8)(unsafe.Add(mBase, uint32(v250))) = uint8(v259)
			v263 = v250 + int32(1)
		}
		v264 = *(*int32)(unsafe.Add(mBase, uint32(l1)+216))
		if v264 == int32(0) {
			v274 = v263
		} else {
			v267 = *(*int32)(unsafe.Add(mBase, uint32(v264)))
			if v267 == int32(1) {
				v274 = v263
			} else {
				v270 = int32(69)
				*(*uint8)(unsafe.Add(mBase, uint32(v263))) = uint8(v270)
				v274 = v263 + int32(1)
			}
		}
	}
	if v274 != v40+int32(288) {
		v283 = v274
	} else {
		v279 = int32(78)
		*(*uint8)(unsafe.Add(mBase, uint32(v274))) = uint8(v279)
		v283 = v274 + int32(1)
	}
	v284 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v283))) = uint8(v284)
	v286 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	if v286 != 0 {
		v289 = *(*int32)(unsafe.Add(mBase, uint32(v286)+36))
		if v289 != 0 {
			v292 = int32(114)
			*(*uint8)(unsafe.Add(mBase, uint32(v40)+285)) = uint8(v292)
			v296 = v40 + int32(286)
		} else {
			v296 = v40 + int32(285)
		}
		v297 = *(*int32)(unsafe.Add(mBase, uint32(v286)+32))
		if v297 == int32(0) {
			v304 = v296
		} else {
			v300 = int32(119)
			*(*uint8)(unsafe.Add(mBase, uint32(v296))) = uint8(v300)
			v304 = v296 + int32(1)
		}
	} else {
		v304 = v40 + int32(285)
	}
	v305 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v304))) = uint8(v305)
	v309 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+227)))
	if v309&int32(1) == v305 {
		v320 = v40 + int32(276)
	} else {
		v314 = int32(114)
		*(*uint8)(unsafe.Add(mBase, uint32(v40)+276)) = uint8(v314)
		v320 = v40 + int32(277)
	}
	v321 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v320))) = uint8(v321)
	v325 = F_getClientMemoryUsage(m, l1, v40+int32(236))
	mBase = m.M
	v328 = m.ExcPending
	if v328 != 0 {
		return int32(0)
	} else {
		v329 = int64(0)
		v330 = *(*int32)(unsafe.Add(mBase, uint32(l1)+104))
		if v330 == int32(0) {
			v349 = v329
		} else {
			v333 = *(*int32)(unsafe.Add(mBase, uint32(v330)+184))
			if v333 == int32(0) {
				v349 = v329
			} else {
				v337 = *(*int32)(unsafe.Add(mBase, _consts[475]))
				v338 = *(*int32)(unsafe.Add(mBase, uint32(v337)+4))
				v339 = *(*int32)(unsafe.Add(mBase, uint32(v338)+8))
				v340 = *(*int64)(unsafe.Add(mBase, uint32(v339)+8))
				v341 = *(*int32)(unsafe.Add(mBase, uint32(v333)+8))
				v342 = *(*int64)(unsafe.Add(mBase, uint32(v341)+8))
				v349 = (v340 - v342 + int64(1)) & int64(4294967295)
			}
		}
		v350 = *(*int64)(unsafe.Add(mBase, uint32(l1)))
		v351 = F_getClientPeerId(m, l1)
		mBase = m.M
		v352 = m.ExcPending
		if v352 != 0 {
			return int32(0)
		} else {
			v353 = F_getClientSockname(m, l1)
			mBase = m.M
			v354 = m.ExcPending
			if v354 != 0 {
				return int32(0)
			} else {
				v355 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
				if v355 != 0 {
					v357 = *(*int32)(unsafe.Add(mBase, uint32(v355)+12))
					v358 = v357
				} else {
					v358 = int32(-1)
				}
				*(*int32)(unsafe.Add(mBase, uint32(v40)+224)) = v358
				v366 = F_snprintf(m, v40+int32(240), int32(31), int32(_a1644), v40+int32(224))
				mBase = m.M
				v367 = m.ExcPending
				if v367 != 0 {
					return int32(0)
				} else {
					if l2 == int32(0) {
						v371 = *(*int32)(unsafe.Add(mBase, uint32(l1)+348))
						if v371 != 0 {
							v373 = F_objectGetVal(m, v371)
							mBase = m.M
							v375 = v373
						} else {
							v375 = int32(_a320)
						}
					} else {
						v375 = int32(_a1645)
					}
					v377 = *(*int64)(unsafe.Add(mBase, _consts[12]))
					v378 = *(*int64)(unsafe.Add(mBase, uint32(l1)+88))
					v380 = *(*int64)(unsafe.Add(mBase, _consts[109]))
					v381 = *(*int64)(unsafe.Add(mBase, uint32(l1)+368))
					v382 = *(*int32)(unsafe.Add(mBase, uint32(l1)+96))
					v383 = *(*int32)(unsafe.Add(mBase, uint32(v382)+28))
					v384 = int32(0)
					v388 = *(*int32)(unsafe.Add(mBase, uint32(l1)+100))
					if v388 == v384 {
						v409 = v384
						v410 = v384
						v411 = v384
					} else {
						v391 = *(*int32)(unsafe.Add(mBase, uint32(v388)))
						v392 = *(*int32)(unsafe.Add(mBase, uint32(v391)+20))
						v393 = *(*int32)(unsafe.Add(mBase, uint32(v391)+16))
						v394 = v392 + v393
						v395 = *(*int32)(unsafe.Add(mBase, uint32(l1)+100))
						if v395 != 0 {
							v398 = *(*int32)(unsafe.Add(mBase, uint32(v395)+4))
							v399 = *(*int32)(unsafe.Add(mBase, uint32(v398)+20))
							v400 = *(*int32)(unsafe.Add(mBase, uint32(v398)+16))
							v401 = v399 + v400
							v402 = *(*int32)(unsafe.Add(mBase, uint32(l1)+100))
							if v402 != 0 {
								v404 = *(*int32)(unsafe.Add(mBase, uint32(v402)+8))
								v405 = *(*int32)(unsafe.Add(mBase, uint32(v404)+20))
								v406 = *(*int32)(unsafe.Add(mBase, uint32(v404)+16))
								v409 = v401
								v410 = v394
								v411 = v405 + v406
							} else {
								v409 = v401
								v410 = v394
								v411 = int32(0)
							}
						} else {
							v396 = int32(0)
							v409 = v396
							v410 = v394
							v411 = v396
						}
					}
					v412 = *(*int32)(unsafe.Add(mBase, uint32(l1)+112))
					if v412 != 0 {
						v414 = *(*int32)(unsafe.Add(mBase, uint32(v412)+44))
						v415 = *(*int32)(unsafe.Add(mBase, uint32(v412)+4))
						v416 = v414
						v417 = v415
					} else {
						v416 = v384
						v417 = int32(-1)
					}
					v418 = int64(0)
					v421 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
					if v421 == int32(0) {
						v472 = v418
						v473 = v418
					} else {
						v424 = int32(0)
						v428 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v421+int32(-1)))))
						switch v428 & int32(7) {
						case 0:
							v468 = base.I64_extend_i32_u(int32(base.Ui32(v428) >> (uint(int32(3)) % 32)))
							v470 = v424
						case 1:
							v436 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v421+int32(-2)))))
							v439 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v421+int32(-3)))))
							v468 = base.I64_extend_i32_u(v439)
							v470 = v436 - v439
						case 2:
							v444 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v421+int32(-3)))))
							v447 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v421+int32(-5)))))
							v468 = base.I64_extend_i32_u(v447)
							v470 = v444 - v447
						case 3:
							v452 = *(*int32)(unsafe.Add(mBase, uint32(v421+int32(-5))))
							v455 = *(*int32)(unsafe.Add(mBase, uint32(v421+int32(-9))))
							v468 = base.I64_extend_i32_u(v455)
							v470 = v452 - v455
						case 4:
							v460 = *(*int64)(unsafe.Add(mBase, uint32(v421+int32(-9))))
							v463 = *(*int64)(unsafe.Add(mBase, uint32(v421+int32(-17))))
							v468 = v463 & int64(4294967295)
							v470 = base.I32_wrap_i64(v460 - v463)
						default:
							v468 = int64(0)
							v470 = v424
						}
						v472 = v468
						v473 = base.I64_extend_i32_u(v470)
					}
					if v412 == int32(0) {
						v480 = v418
					} else {
						v479 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v412)+16)))
						v480 = v479
					}
					v481 = *(*int32)(unsafe.Add(mBase, uint32(l1)+72))
					if v481 != 0 {
						v483 = *(*int32)(unsafe.Add(mBase, uint32(v481)+140))
						v484 = v483
					} else {
						v484 = int32(_a1646)
					}
					if l2 == int32(0) {
						v488 = *(*int32)(unsafe.Add(mBase, uint32(l1)+328))
						if v488 != 0 {
							v490 = *(*int32)(unsafe.Add(mBase, uint32(v488)))
							v492 = v490
						} else {
							v492 = int32(_a1647)
						}
					} else {
						v492 = int32(_a1645)
					}
					v493 = *(*int32)(unsafe.Add(mBase, uint32(l1)+132))
					v495 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+204)))
					if v495&int32(4) == int32(0) {
						v502 = int64(-1)
					} else {
						v500 = *(*int32)(unsafe.Add(mBase, uint32(l1)+100))
						v501 = *(*int64)(unsafe.Add(mBase, uint32(v500)+16))
						v502 = v501
					}
					v503 = int64(*(*uint32)(unsafe.Add(mBase, uint32(l1)+32)))
					v504 = int64(*(*uint32)(unsafe.Add(mBase, uint32(l1)+180)))
					v505 = int64(*(*uint32)(unsafe.Add(mBase, uint32(l1)+276)))
					v506 = int64(*(*uint32)(unsafe.Add(mBase, uint32(l1)+128)))
					v507 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v493)+20)))
					v508 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v40)+236)))
					v509 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+224)))
					v510 = int32(_a320)
					v512 = *(*int32)(unsafe.Add(mBase, uint32(l1)+352))
					if v512 == int32(0) {
						v516 = v510
					} else {
						v515 = F_objectGetVal(m, v512)
						mBase = m.M
						v516 = v515
					}
					v517 = *(*int32)(unsafe.Add(mBase, uint32(l1)+356))
					if v517 == int32(0) {
						v521 = v510
					} else {
						v520 = F_objectGetVal(m, v517)
						mBase = m.M
						v521 = v520
					}
					v522 = *(*int64)(unsafe.Add(mBase, uint32(l1)+232))
					v523 = *(*int64)(unsafe.Add(mBase, uint32(l1)+248))
					v526 = *(*int64)(unsafe.Add(mBase, uint32(l1)+256))
					*(*int64)(unsafe.Add(mBase, uint32(v40+int32(208)))) = v526
					*(*int64)(unsafe.Add(mBase, uint32(v40+int32(200)))) = v523
					*(*int64)(unsafe.Add(mBase, uint32(v40+int32(192)))) = v522
					*(*int32)(unsafe.Add(mBase, uint32(v40+int32(184)))) = v521
					*(*int32)(unsafe.Add(mBase, uint32(v40+int32(180)))) = v516
					*(*int32)(unsafe.Add(mBase, uint32(v40+int32(176)))) = v509
					*(*int64)(unsafe.Add(mBase, uint32(v40+int32(168)))) = v502
					*(*int32)(unsafe.Add(mBase, uint32(v40+int32(160)))) = v492
					*(*int32)(unsafe.Add(mBase, uint32(v40+int32(156)))) = v484
					*(*int64)(unsafe.Add(mBase, uint32(v40+int32(144)))) = base.I64_extend_i32_u(v325)
					*(*int64)(unsafe.Add(mBase, uint32(v40+int32(136)))) = v508
					*(*int64)(unsafe.Add(mBase, uint32(v40+int32(128)))) = v349 + v507
					*(*int64)(unsafe.Add(mBase, uint32(v40+int32(120)))) = v504
					*(*int64)(unsafe.Add(mBase, uint32(v40+int32(112)))) = v505
					*(*int64)(unsafe.Add(mBase, uint32(v40+int32(104)))) = v506
					*(*int64)(unsafe.Add(mBase, uint32(v40+int32(96)))) = v480
					*(*int64)(unsafe.Add(mBase, uint32(v40+int32(88)))) = v503
					*(*int64)(unsafe.Add(mBase, uint32(v40+int32(80)))) = v473
					*(*int64)(unsafe.Add(mBase, uint32(v40+int32(72)))) = v472
					*(*int32)(unsafe.Add(mBase, uint32(v40+int32(68)))) = v416
					*(*int32)(unsafe.Add(mBase, uint32(v40+int32(64)))) = v417
					*(*int32)(unsafe.Add(mBase, uint32(v40+int32(60)))) = v411
					*(*int32)(unsafe.Add(mBase, uint32(v40+int32(56)))) = v409
					*(*int32)(unsafe.Add(mBase, uint32(v40+int32(52)))) = v410
					*(*int32)(unsafe.Add(mBase, uint32(v40+int32(48)))) = v383
					*(*int64)(unsafe.Add(mBase, uint32(v40+int32(32)))) = v380 - v378
					*(*int32)(unsafe.Add(mBase, uint32(v40+int32(20)))) = v375
					v612 = base.I64_div_s(v377, int64(1000))
					*(*int64)(unsafe.Add(mBase, uint32(v40+int32(24)))) = v612 - v381
					*(*int32)(unsafe.Add(mBase, uint32(v40+int32(152)))) = v40 + int32(285)
					*(*int32)(unsafe.Add(mBase, uint32(v40+int32(44)))) = v40 + int32(276)
					*(*int32)(unsafe.Add(mBase, uint32(v40+int32(40)))) = v40 + int32(288)
					*(*int32)(unsafe.Add(mBase, uint32(v40+int32(16)))) = v40 + int32(240)
					*(*int32)(unsafe.Add(mBase, uint32(v40)+12)) = v353
					*(*int32)(unsafe.Add(mBase, uint32(v40)+8)) = v351
					*(*int64)(unsafe.Add(mBase, uint32(v40))) = v350
					v639 = F_sdscatfmt(m, l0, int32(_a1648), v40)
					mBase = m.M
					v640 = m.ExcPending
					if v640 != 0 {
						return int32(0)
					} else {
						m.G0 = v40 + int32(320)
						return v639
					}
				}
			}
		}
	}
}
func F_checkChildrenDone(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v1 int32
	_ = v1
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v65 int32
	_ = v65
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v80 int32
	_ = v80
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v91 int32
	_ = v91
	var v99 int32
	_ = v99
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v104 int64
	_ = v104
	var v120 int32
	_ = v120
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v134 int32
	_ = v134
	var v138 int32
	_ = v138
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v159 int32
	_ = v159
	var v168 int32
	_ = v168
	var v173 int32
	_ = v173
	v1 = int32(0)
	v5 = m.G0
	v7 = v5 - int32(48)
	m.G0 = v7
	*(*int32)(unsafe.Add(mBase, uint32(v7)+44)) = v1
	v16 = F___syscall_wait4(m, int32(-1), v7+int32(44), int32(1), v1)
	mBase = m.M
	v17 = F___syscall_ret(m, v16)
	mBase = m.M
	if v17 == int32(0) {
		m.G0 = v7 + int32(48)
		return
	} else {
		if v17 != int32(-1) {
			v52 = *(*int32)(unsafe.Add(mBase, _consts[61]))
			if v17 != v52 {
				v156 = F_scriptingEngineDebuggerRemoveChild(m, v17)
				mBase = m.M
				v157 = m.ExcPending
				if v157 != 0 {
					return
				} else {
					if v156 != 0 {
						F_replicationStartPendingFork(m)
						mBase = m.M
						v173 = m.ExcPending
						if v173 != 0 {
							return
						} else {
							m.G0 = v7 + int32(48)
							return
						}
					} else {
						v159 = *(*int32)(unsafe.Add(mBase, _consts[28]))
						if int32(3) < v159 {
							F_replicationStartPendingFork(m)
							mBase = m.M
							v173 = m.ExcPending
							if v173 != 0 {
								return
							} else {
								m.G0 = v7 + int32(48)
								return
							}
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v7)+32)) = v17
							F__serverLog(m, int32(3), int32(_a2159), v7+int32(32))
							mBase = m.M
							v168 = m.ExcPending
							if v168 != 0 {
								return
							} else {
								F_replicationStartPendingFork(m)
								mBase = m.M
								v173 = m.ExcPending
								if v173 != 0 {
									return
								} else {
									m.G0 = v7 + int32(48)
									return
								}
							}
						}
					}
				}
			} else {
				v55 = *(*int32)(unsafe.Add(mBase, uint32(v7)+44))
				v57 = v55 & int32(127)
				if base.Ui32(v55&int32(65535)+int32(-1)) < base.Ui32(int32(255)) {
					v65 = v57
				} else {
					v65 = int32(0)
				}
				if v57 != 0 {
					v71 = int32(-1)
				} else {
					v71 = int32(base.Ui32(v55)>>(uint(int32(8))%32)) & int32(255)
				}
				v73 = base.B2i32(v71 == int32(255))
				if v71 == int32(255) {
					v74 = int32(10)
				} else {
					v74 = v65
				}
				if v71 == int32(255) {
					v76 = int32(1)
				} else {
					v76 = v71
				}
				if v74|v76 != 0 {
				} else {
					F_receiveChildInfo(m)
					mBase = m.M
				}
				v80 = *(*int32)(unsafe.Add(mBase, _consts[60]))
				switch v80 + int32(-1) {
				case 0:
					F_backgroundSaveDoneHandler(m, v76, v74)
					mBase = m.M
					v102 = m.ExcPending
					if v102 != 0 {
						return
					} else {
						v103 = int32(0)
						v104 = int64(0)
						*(*int64)(unsafe.Add(mBase, _consts[704])) = v104
						*(*int64)(unsafe.Add(mBase, _consts[61])) = int64(4294967295)
						*(*int64)(unsafe.Add(mBase, _consts[705])) = v104
						*(*int64)(unsafe.Add(mBase, _consts[706])) = v104
						*(*int64)(unsafe.Add(mBase, _consts[707])) = v104
						v120 = *(*int32)(unsafe.Add(mBase, _consts[708]))
						v126 = *(*int32)(unsafe.Add(mBase, _consts[709]))
						if v126 != 0 {
							v127 = int32(2)
						} else {
							v127 = base.B2i32(v120 == v103) << (uint(int32(1)) % 32)
						}
						*(*int32)(unsafe.Add(mBase, _consts[302])) = v127
						*(*int32)(unsafe.Add(mBase, _consts[340])) = v127
						v134 = *(*int32)(unsafe.Add(mBase, _consts[98]))
						if v134 != int32(-1) {
							v141 = F_close(m, v134)
							mBase = m.M
							v142 = int32(_a20)
							v143 = *(*int32)(unsafe.Add(mBase, _consts[99]))
							v144 = F_close(m, v143)
							mBase = m.M
							*(*int32)(unsafe.Add(mBase, _consts[100])) = int32(0)
							*(*int64)(unsafe.Add(mBase, _consts[98])) = int64(-1)
						} else {
							v138 = *(*int32)(unsafe.Add(mBase, _consts[99]))
							if v138 == int32(-1) {
							} else {
								v141 = F_close(m, v134)
								mBase = m.M
								v142 = int32(_a20)
								v143 = *(*int32)(unsafe.Add(mBase, _consts[99]))
								v144 = F_close(m, v143)
								mBase = m.M
								*(*int32)(unsafe.Add(mBase, _consts[100])) = int32(0)
								*(*int64)(unsafe.Add(mBase, _consts[98])) = int64(-1)
							}
						}
						F_moduleFireServerEvent(m, int64(13), int32(1), int32(0))
						mBase = m.M
						v155 = m.ExcPending
						if v155 != 0 {
							return
						} else {
							F_replicationStartPendingFork(m)
							mBase = m.M
							v173 = m.ExcPending
							if v173 != 0 {
								return
							} else {
								m.G0 = v7 + int32(48)
								return
							}
						}
					}
				case 1:
					F_backgroundRewriteDoneHandler(m, v76, v74)
					mBase = m.M
					v84 = m.ExcPending
					if v84 != 0 {
						return
					} else {
						v103 = int32(0)
						v104 = int64(0)
						*(*int64)(unsafe.Add(mBase, _consts[704])) = v104
						*(*int64)(unsafe.Add(mBase, _consts[61])) = int64(4294967295)
						*(*int64)(unsafe.Add(mBase, _consts[705])) = v104
						*(*int64)(unsafe.Add(mBase, _consts[706])) = v104
						*(*int64)(unsafe.Add(mBase, _consts[707])) = v104
						v120 = *(*int32)(unsafe.Add(mBase, _consts[708]))
						v126 = *(*int32)(unsafe.Add(mBase, _consts[709]))
						if v126 != 0 {
							v127 = int32(2)
						} else {
							v127 = base.B2i32(v120 == v103) << (uint(int32(1)) % 32)
						}
						*(*int32)(unsafe.Add(mBase, _consts[302])) = v127
						*(*int32)(unsafe.Add(mBase, _consts[340])) = v127
						v134 = *(*int32)(unsafe.Add(mBase, _consts[98]))
						if v134 != int32(-1) {
							v141 = F_close(m, v134)
							mBase = m.M
							v142 = int32(_a20)
							v143 = *(*int32)(unsafe.Add(mBase, _consts[99]))
							v144 = F_close(m, v143)
							mBase = m.M
							*(*int32)(unsafe.Add(mBase, _consts[100])) = int32(0)
							*(*int64)(unsafe.Add(mBase, _consts[98])) = int64(-1)
						} else {
							v138 = *(*int32)(unsafe.Add(mBase, _consts[99]))
							if v138 == int32(-1) {
							} else {
								v141 = F_close(m, v134)
								mBase = m.M
								v142 = int32(_a20)
								v143 = *(*int32)(unsafe.Add(mBase, _consts[99]))
								v144 = F_close(m, v143)
								mBase = m.M
								*(*int32)(unsafe.Add(mBase, _consts[100])) = int32(0)
								*(*int64)(unsafe.Add(mBase, _consts[98])) = int64(-1)
							}
						}
						F_moduleFireServerEvent(m, int64(13), int32(1), int32(0))
						mBase = m.M
						v155 = m.ExcPending
						if v155 != 0 {
							return
						} else {
							F_replicationStartPendingFork(m)
							mBase = m.M
							v173 = m.ExcPending
							if v173 != 0 {
								return
							} else {
								m.G0 = v7 + int32(48)
								return
							}
						}
					}
				default:
					*(*int32)(unsafe.Add(mBase, uint32(v7)+16)) = v80
					v91 = *(*int32)(unsafe.Add(mBase, _consts[61]))
					*(*int32)(unsafe.Add(mBase, uint32(v7)+20)) = v91
					F__serverPanic_1(m, int32(_a2157), int32(1427), int32(_a2160), v7+int32(16))
					mBase = m.M
					v99 = m.ExcPending
					if v99 != 0 {
						return
					} else {
						F_abort(m)
						mBase = m.M
						base.Wasm_trap_unreachable()
						for {
						}
					}
				case 3:
					F_ModuleForkDoneHandler(m, v76, v74)
					mBase = m.M
					v86 = m.ExcPending
					if v86 != 0 {
						return
					} else {
						v103 = int32(0)
						v104 = int64(0)
						*(*int64)(unsafe.Add(mBase, _consts[704])) = v104
						*(*int64)(unsafe.Add(mBase, _consts[61])) = int64(4294967295)
						*(*int64)(unsafe.Add(mBase, _consts[705])) = v104
						*(*int64)(unsafe.Add(mBase, _consts[706])) = v104
						*(*int64)(unsafe.Add(mBase, _consts[707])) = v104
						v120 = *(*int32)(unsafe.Add(mBase, _consts[708]))
						v126 = *(*int32)(unsafe.Add(mBase, _consts[709]))
						if v126 != 0 {
							v127 = int32(2)
						} else {
							v127 = base.B2i32(v120 == v103) << (uint(int32(1)) % 32)
						}
						*(*int32)(unsafe.Add(mBase, _consts[302])) = v127
						*(*int32)(unsafe.Add(mBase, _consts[340])) = v127
						v134 = *(*int32)(unsafe.Add(mBase, _consts[98]))
						if v134 != int32(-1) {
							v141 = F_close(m, v134)
							mBase = m.M
							v142 = int32(_a20)
							v143 = *(*int32)(unsafe.Add(mBase, _consts[99]))
							v144 = F_close(m, v143)
							mBase = m.M
							*(*int32)(unsafe.Add(mBase, _consts[100])) = int32(0)
							*(*int64)(unsafe.Add(mBase, _consts[98])) = int64(-1)
						} else {
							v138 = *(*int32)(unsafe.Add(mBase, _consts[99]))
							if v138 == int32(-1) {
							} else {
								v141 = F_close(m, v134)
								mBase = m.M
								v142 = int32(_a20)
								v143 = *(*int32)(unsafe.Add(mBase, _consts[99]))
								v144 = F_close(m, v143)
								mBase = m.M
								*(*int32)(unsafe.Add(mBase, _consts[100])) = int32(0)
								*(*int64)(unsafe.Add(mBase, _consts[98])) = int64(-1)
							}
						}
						F_moduleFireServerEvent(m, int64(13), int32(1), int32(0))
						mBase = m.M
						v155 = m.ExcPending
						if v155 != 0 {
							return
						} else {
							F_replicationStartPendingFork(m)
							mBase = m.M
							v173 = m.ExcPending
							if v173 != 0 {
								return
							} else {
								m.G0 = v7 + int32(48)
								return
							}
						}
					}
				case 4:
					F_backgroundSlotMigrationDoneHandler(m, v76, v74)
					mBase = m.M
					v88 = m.ExcPending
					if v88 != 0 {
						return
					} else {
						v103 = int32(0)
						v104 = int64(0)
						*(*int64)(unsafe.Add(mBase, _consts[704])) = v104
						*(*int64)(unsafe.Add(mBase, _consts[61])) = int64(4294967295)
						*(*int64)(unsafe.Add(mBase, _consts[705])) = v104
						*(*int64)(unsafe.Add(mBase, _consts[706])) = v104
						*(*int64)(unsafe.Add(mBase, _consts[707])) = v104
						v120 = *(*int32)(unsafe.Add(mBase, _consts[708]))
						v126 = *(*int32)(unsafe.Add(mBase, _consts[709]))
						if v126 != 0 {
							v127 = int32(2)
						} else {
							v127 = base.B2i32(v120 == v103) << (uint(int32(1)) % 32)
						}
						*(*int32)(unsafe.Add(mBase, _consts[302])) = v127
						*(*int32)(unsafe.Add(mBase, _consts[340])) = v127
						v134 = *(*int32)(unsafe.Add(mBase, _consts[98]))
						if v134 != int32(-1) {
							v141 = F_close(m, v134)
							mBase = m.M
							v142 = int32(_a20)
							v143 = *(*int32)(unsafe.Add(mBase, _consts[99]))
							v144 = F_close(m, v143)
							mBase = m.M
							*(*int32)(unsafe.Add(mBase, _consts[100])) = int32(0)
							*(*int64)(unsafe.Add(mBase, _consts[98])) = int64(-1)
						} else {
							v138 = *(*int32)(unsafe.Add(mBase, _consts[99]))
							if v138 == int32(-1) {
							} else {
								v141 = F_close(m, v134)
								mBase = m.M
								v142 = int32(_a20)
								v143 = *(*int32)(unsafe.Add(mBase, _consts[99]))
								v144 = F_close(m, v143)
								mBase = m.M
								*(*int32)(unsafe.Add(mBase, _consts[100])) = int32(0)
								*(*int64)(unsafe.Add(mBase, _consts[98])) = int64(-1)
							}
						}
						F_moduleFireServerEvent(m, int64(13), int32(1), int32(0))
						mBase = m.M
						v155 = m.ExcPending
						if v155 != 0 {
							return
						} else {
							F_replicationStartPendingFork(m)
							mBase = m.M
							v173 = m.ExcPending
							if v173 != 0 {
								return
							} else {
								m.G0 = v7 + int32(48)
								return
							}
						}
					}
				}
			}
		} else {
			v23 = *(*int32)(unsafe.Add(mBase, _consts[28]))
			if int32(3) < v23 {
				F_replicationStartPendingFork(m)
				mBase = m.M
				v173 = m.ExcPending
				if v173 != 0 {
					return
				} else {
					m.G0 = v7 + int32(48)
					return
				}
			} else {
				v27 = *(*int32)(unsafe.Add(mBase, _consts[18]))
				v28 = F___strerror_l(m, v27, v27)
				mBase = m.M
				v31 = *(*int32)(unsafe.Add(mBase, _consts[60]))
				v33 = v31 + int32(-1)
				if base.Ui32(int32(4)) < base.Ui32(v33) {
					v41 = int32(_a2161)
				} else {
					v40 = *(*int32)(unsafe.Add(mBase, uint32(v33<<(uint(int32(2))%32))+uint32(_consts[710])))
					v41 = v40
				}
				*(*int32)(unsafe.Add(mBase, uint32(v7))) = v28
				*(*int32)(unsafe.Add(mBase, uint32(v7)+4)) = v41
				v45 = *(*int32)(unsafe.Add(mBase, _consts[61]))
				*(*int32)(unsafe.Add(mBase, uint32(v7)+8)) = v45
				F__serverLog(m, int32(3), int32(_a2162), v7)
				mBase = m.M
				v50 = m.ExcPending
				if v50 != 0 {
					return
				} else {
					F_replicationStartPendingFork(m)
					mBase = m.M
					v173 = m.ExcPending
					if v173 != 0 {
						return
					} else {
						m.G0 = v7 + int32(48)
						return
					}
				}
			}
		}
	}
}
func F_checkSingleAof(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v19 int32
	_ = v19
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v33 int32
	_ = v33
	var v37 int32
	_ = v37
	var v41 int32
	_ = v41
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v54 int64
	_ = v54
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v83 int32
	_ = v83
	var v87 int64
	_ = v87
	var v90 int64
	_ = v90
	var v91 int32
	_ = v91
	var v92 int64
	_ = v92
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v100 int32
	_ = v100
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v109 int32
	_ = v109
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v129 int32
	_ = v129
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
	var v142 int32
	_ = v142
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v154 int32
	_ = v154
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v163 int32
	_ = v163
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v185 int64
	_ = v185
	var v190 int64
	_ = v190
	var v195 int64
	_ = v195
	var v198 int64
	_ = v198
	var v201 int64
	_ = v201
	var v204 int64
	_ = v204
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v228 int64
	_ = v228
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v247 int64
	_ = v247
	var v251 int64
	_ = v251
	var v258 int32
	_ = v258
	var v259 int32
	_ = v259
	var v277 int32
	_ = v277
	var v278 int32
	_ = v278
	var v281 int32
	_ = v281
	var v282 int32
	_ = v282
	var v287 int32
	_ = v287
	var v288 int32
	_ = v288
	var v289 int32
	_ = v289
	var v293 int32
	_ = v293
	var v294 int32
	_ = v294
	var v299 int32
	_ = v299
	var v301 int32
	_ = v301
	var v302 int32
	_ = v302
	var v303 int32
	_ = v303
	var v304 int32
	_ = v304
	var v306 int32
	_ = v306
	var v310 int32
	_ = v310
	var v314 int32
	_ = v314
	var v316 int32
	_ = v316
	var v317 int32
	_ = v317
	var v318 int32
	_ = v318
	var v320 int32
	_ = v320
	var v321 int32
	_ = v321
	var v322 int32
	_ = v322
	var v323 int32
	_ = v323
	var v328 int32
	_ = v328
	var v330 int32
	_ = v330
	var v333 int32
	_ = v333
	var v334 int32
	_ = v334
	var v338 int32
	_ = v338
	var v339 int32
	_ = v339
	var v340 int32
	_ = v340
	var v350 int32
	_ = v350
	var v353 int32
	_ = v353
	var v354 int32
	_ = v354
	var v355 int32
	_ = v355
	var v359 int32
	_ = v359
	var v363 int32
	_ = v363
	var v367 int32
	_ = v367
	var v368 int32
	_ = v368
	var v371 int32
	_ = v371
	var v372 int32
	_ = v372
	var v378 int32
	_ = v378
	var v379 int32
	_ = v379
	var v382 int32
	_ = v382
	var v383 int32
	_ = v383
	var v392 int32
	_ = v392
	var v402 int32
	_ = v402
	var v403 int32
	_ = v403
	var v407 int32
	_ = v407
	var v408 int32
	_ = v408
	var v415 int32
	_ = v415
	var v416 int32
	_ = v416
	var v417 int32
	_ = v417
	var v418 int32
	_ = v418
	var v422 int32
	_ = v422
	var v423 int32
	_ = v423
	var v427 int32
	_ = v427
	var v428 int32
	_ = v428
	var v434 int32
	_ = v434
	var v435 int32
	_ = v435
	var v442 int32
	_ = v442
	var v443 int32
	_ = v443
	var v447 int32
	_ = v447
	var v448 int32
	_ = v448
	var v455 int32
	_ = v455
	var v456 int32
	_ = v456
	v6 = int32(0)
	v12 = m.G0
	v14 = v12 - int32(1376)
	m.G0 = v14
	*(*int32)(unsafe.Add(mBase, uint32(v14)+1372)) = v6
	v19 = F_fopen(m, l1, int32(_a2476))
	mBase = m.M
	if v19 == v6 {
		goto L7
	} else {
		goto L8
	}
L1:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+96)) = l0
	v455 = F_iprintf(m, int32(_a2477), v14+int32(96))
	mBase = m.M
	v456 = m.ExcPending
	if v456 != 0 {
		goto L23
	} else {
		goto L125
	}
L2:
	;
	v447 = F_puts(m, int32(_a2478))
	mBase = m.M
	v448 = m.ExcPending
	if v448 != 0 {
		goto L23
	} else {
		goto L124
	}
L3:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+80)) = l0
	v442 = F_iprintf(m, int32(_a2479), v14+int32(80))
	mBase = m.M
	v443 = m.ExcPending
	if v443 != 0 {
		goto L23
	} else {
		goto L123
	}
L4:
	;
	goto L120
L5:
	;
	v422 = F_puts(m, int32(_a2480))
	mBase = m.M
	v423 = m.ExcPending
	if v423 != 0 {
		goto L23
	} else {
		goto L119
	}
L6:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+16)) = l0
	v415 = F_iprintf(m, int32(_a2481), v14+int32(16))
	mBase = m.M
	v416 = m.ExcPending
	if v416 != 0 {
		goto L23
	} else {
		goto L117
	}
L7:
	;
	goto L114
L8:
	;
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v19)+76))
	if int32(-1) < v24 {
		goto L11
	} else {
		goto L12
	}
L9:
	;
	if int32(-1) < v41 {
		goto L17
	} else {
		goto L18
	}
L10:
	;
	if int32(-1) < v33 {
		v41 = v33
		goto L14
	} else {
		goto L15
	}
L11:
	;
	v28 = F___lockfile(m, v19)
	mBase = m.M
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v19)+60))
	if v28 == int32(0) {
		v33 = v29
		goto L10
	} else {
		goto L13
	}
L12:
	;
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v19)+60))
	v33 = v27
	goto L10
L13:
	;
	F___unlockfile(m, v19)
	mBase = m.M
	v33 = v29
	goto L10
L14:
	;
	goto L9
L15:
	;
	v37 = F___errno_location(m)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v37))) = int32(8)
	v41 = int32(-1)
	goto L14
L16:
	;
	if v51 == int32(-1) {
		goto L6
	} else {
		goto L19
	}
L17:
	;
	v50 = F___fstatat(m, v41, int32(_a320), v14+int32(1272), int32(4096))
	mBase = m.M
	v51 = v50
	goto L16
L18:
	;
	v47 = F___syscall_ret(m, int32(-8))
	mBase = m.M
	v51 = v47
	goto L16
L19:
	;
	v54 = *(*int64)(unsafe.Add(mBase, uint32(v14)+1296))
	if v54 != int64(0) {
		goto L21
	} else {
		goto L22
	}
L20:
	;
	m.G0 = v14 + int32(1376)
	return v392
L21:
	;
	if l4 == int32(0) {
		goto L25
	} else {
		goto L26
	}
L22:
	;
	v57 = F_fclose(m, v19)
	mBase = m.M
	v60 = m.ExcPending
	if v60 != 0 {
		goto L23
	} else {
		goto L24
	}
L23:
	;
	return int32(0)
L24:
	;
	v392 = int32(1)
	goto L20
L25:
	;
	v83 = int32(0)
	v87 = int64(0)
	goto L32
L26:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+244)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v14)+240)) = int32(0)
	v70 = F_redis_check_rdb_main(m, int32(2), v14+int32(240), v19)
	mBase = m.M
	v71 = m.ExcPending
	if v71 != 0 {
		goto L23
	} else {
		goto L27
	}
L27:
	;
	if v70 == int32(-1) {
		goto L5
	} else {
		goto L28
	}
L28:
	;
	v75 = F_puts(m, int32(_a2482))
	mBase = m.M
	v76 = m.ExcPending
	if v76 != 0 {
		goto L23
	} else {
		goto L29
	}
L29:
	;
	goto L25
L30:
	;
	v154 = *(*int32)(unsafe.Add(mBase, uint32(v19)+76))
	if int32(-1) < v154 {
		goto L60
	} else {
		goto L61
	}
L31:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+224)) = l0
	v148 = F_iprintf(m, int32(_a2483), v14+int32(224))
	mBase = m.M
	v149 = m.ExcPending
	if v149 != 0 {
		goto L23
	} else {
		goto L57
	}
L32:
	;
	if v83 != 0 {
		v92 = v87
		goto L34
	} else {
		goto L35
	}
L33:
	;
	v141 = F_fclose(m, v19)
	mBase = m.M
	v142 = m.ExcPending
	if v142 != 0 {
		goto L23
	} else {
		goto L56
	}
L34:
	;
	v96 = F_fgets(m, v14+int32(1370), int32(2), v19)
	mBase = m.M
	v97 = m.ExcPending
	if v97 != 0 {
		goto L23
	} else {
		goto L38
	}
L35:
	;
	v90 = F___ftello(m, v19)
	mBase = m.M
	v91 = m.ExcPending
	if v91 != 0 {
		goto L23
	} else {
		goto L36
	}
L36:
	;
	v92 = v90
	goto L34
L37:
	;
	v125 = F_fseek(m, v19, int32(-1), int32(1))
	mBase = m.M
	v126 = m.ExcPending
	if v126 != 0 {
		goto L23
	} else {
		goto L47
	}
L38:
	;
	if v96 != 0 {
		goto L37
	} else {
		goto L39
	}
L39:
	;
	v100 = *(*int32)(unsafe.Add(mBase, uint32(v19)+76))
	if int32(-1) < v100 {
		goto L42
	} else {
		goto L43
	}
L40:
	;
	if int32(base.Ui32(v109)>>(uint(int32(4))%32))&int32(1) != 0 {
		v150 = v83
		goto L30
	} else {
		goto L45
	}
L41:
	;
	goto L40
L42:
	;
	v104 = F___lockfile(m, v19)
	mBase = m.M
	v105 = *(*int32)(unsafe.Add(mBase, uint32(v19)))
	if v104 == int32(0) {
		v109 = v105
		goto L41
	} else {
		goto L44
	}
L43:
	;
	v103 = *(*int32)(unsafe.Add(mBase, uint32(v19)))
	v109 = v103
	goto L41
L44:
	;
	F___unlockfile(m, v19)
	mBase = m.M
	v109 = v105
	goto L41
L45:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+32)) = l0
	v119 = F_iprintf(m, int32(_a2484), v14+int32(32))
	mBase = m.M
	v120 = m.ExcPending
	if v120 != 0 {
		goto L23
	} else {
		goto L46
	}
L46:
	;
	m.Env.Exit(m, int32(1))
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L47:
	;
	if v125 == int32(-1) {
		goto L4
	} else {
		goto L48
	}
L48:
	;
	v129 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+1370)))
	if v129 == int32(35) {
		goto L49
	} else {
		goto L50
	}
L49:
	;
	v139 = F_processAnnotations(m, v19, l1, l2)
	mBase = m.M
	v140 = m.ExcPending
	if v140 != 0 {
		goto L23
	} else {
		goto L54
	}
L50:
	;
	if v129 != int32(42) {
		goto L31
	} else {
		goto L51
	}
L51:
	;
	v136 = F_processRESP(m, v19, l1, v14+int32(1372))
	mBase = m.M
	v137 = m.ExcPending
	if v137 != 0 {
		goto L23
	} else {
		goto L52
	}
L52:
	;
	v138 = *(*int32)(unsafe.Add(mBase, uint32(v14)+1372))
	if v136 != 0 {
		v83 = v138
		v87 = v92
		goto L32
	} else {
		goto L53
	}
L53:
	;
	v150 = v138
	goto L30
L54:
	;
	if v139 != 0 {
		v87 = v92
		goto L32
	} else {
		goto L55
	}
L55:
	;
	goto L33
L56:
	;
	v392 = int32(3)
	goto L20
L57:
	;
	v150 = v83
	goto L30
L58:
	;
	v169 = int32(0)
	v170 = int32(*(*uint8)(unsafe.Add(mBase, _consts[905])))
	if int32(base.Ui32(v163)>>(uint(int32(4))%32))&int32(1) == v169 {
		v218 = v170
		goto L63
	} else {
		goto L64
	}
L59:
	;
	goto L58
L60:
	;
	v158 = F___lockfile(m, v19)
	mBase = m.M
	v159 = *(*int32)(unsafe.Add(mBase, uint32(v19)))
	if v158 == int32(0) {
		v163 = v159
		goto L59
	} else {
		goto L62
	}
L61:
	;
	v157 = *(*int32)(unsafe.Add(mBase, uint32(v19)))
	v163 = v157
	goto L59
L62:
	;
	F___unlockfile(m, v19)
	mBase = m.M
	v163 = v159
	goto L59
L63:
	;
	if v218&int32(255) == int32(0) {
		goto L68
	} else {
		goto L69
	}
L64:
	;
	if v150 == int32(0) {
		v218 = v170
		goto L63
	} else {
		goto L65
	}
L65:
	;
	if v170&int32(255) != 0 {
		v218 = v170
		goto L63
	} else {
		goto L66
	}
L66:
	;
	v179 = int32(0)
	v180 = int32(*(*uint16)(unsafe.Add(mBase, _consts[906])))
	*(*uint16)(unsafe.Add(mBase, uint32(v14+int32(280)))) = uint16(v180)
	v185 = *(*int64)(unsafe.Add(mBase, _consts[907]))
	*(*int64)(unsafe.Add(mBase, uint32(v14+int32(272)))) = v185
	v190 = *(*int64)(unsafe.Add(mBase, _consts[908]))
	*(*int64)(unsafe.Add(mBase, uint32(v14+int32(264)))) = v190
	v195 = *(*int64)(unsafe.Add(mBase, _consts[909]))
	*(*int64)(unsafe.Add(mBase, uint32(v14+int32(256)))) = v195
	v198 = *(*int64)(unsafe.Add(mBase, _consts[910]))
	*(*int64)(unsafe.Add(mBase, uint32(v14)+248)) = v198
	v201 = *(*int64)(unsafe.Add(mBase, _consts[911]))
	*(*int64)(unsafe.Add(mBase, uint32(v14)+240)) = v201
	v204 = *(*int64)(unsafe.Add(mBase, _consts[903]))
	*(*int64)(unsafe.Add(mBase, uint32(v14)+192)) = v204
	*(*int32)(unsafe.Add(mBase, uint32(v14)+200)) = v14 + int32(240)
	v214 = F_snprintf(m, int32(_a2485), int32(1044), int32(_a2486), v14+int32(192))
	mBase = m.M
	v215 = m.ExcPending
	if v215 != 0 {
		goto L23
	} else {
		goto L67
	}
L67:
	;
	v217 = int32(*(*uint8)(unsafe.Add(mBase, _consts[905])))
	v218 = v217
	goto L63
L68:
	;
	if v54 != v92 {
		goto L72
	} else {
		goto L73
	}
L69:
	;
	v224 = F_puts(m, int32(_a2485))
	mBase = m.M
	v225 = m.ExcPending
	if v225 != 0 {
		goto L23
	} else {
		goto L70
	}
L70:
	;
	goto L68
L71:
	;
	v392 = int32(0)
	goto L20
L72:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v14+int32(160)))) = v92
	v247 = *(*int64)(unsafe.Add(mBase, _consts[912]))
	*(*int64)(unsafe.Add(mBase, uint32(v14+int32(168)))) = v247
	v251 = v54 - v92
	*(*int64)(unsafe.Add(mBase, uint32(v14+int32(176)))) = v251
	*(*int32)(unsafe.Add(mBase, uint32(v14)+144)) = l0
	*(*int64)(unsafe.Add(mBase, uint32(v14)+152)) = v54
	v258 = F_iprintf(m, int32(_a2487), v14+int32(144))
	mBase = m.M
	v259 = m.ExcPending
	if v259 != 0 {
		goto L23
	} else {
		goto L77
	}
L73:
	;
	v228 = *(*int64)(unsafe.Add(mBase, _consts[904]))
	if v228 == int64(0) {
		goto L72
	} else {
		goto L74
	}
L74:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v14)+56)) = v228
	*(*int32)(unsafe.Add(mBase, uint32(v14)+48)) = l0
	v236 = F_iprintf(m, int32(_a2488), v14+int32(48))
	mBase = m.M
	v237 = m.ExcPending
	if v237 != 0 {
		goto L23
	} else {
		goto L75
	}
L75:
	;
	v238 = F_fclose(m, v19)
	mBase = m.M
	v239 = m.ExcPending
	if v239 != 0 {
		goto L23
	} else {
		goto L76
	}
L76:
	;
	goto L71
L77:
	;
	if v251 < int64(1) {
		goto L78
	} else {
		goto L79
	}
L78:
	;
	v382 = F_fclose(m, v19)
	mBase = m.M
	v383 = m.ExcPending
	if v383 != 0 {
		goto L23
	} else {
		goto L113
	}
L79:
	;
	if l3 == int32(0) {
		goto L80
	} else {
		goto L81
	}
L80:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+64)) = l0
	v378 = F_iprintf(m, int32(_a2489), v14+int32(64))
	mBase = m.M
	v379 = m.ExcPending
	if v379 != 0 {
		goto L23
	} else {
		goto L112
	}
L81:
	;
	if l2 == int32(0) {
		goto L3
	} else {
		goto L82
	}
L82:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v14+int32(128)))) = v251
	*(*int64)(unsafe.Add(mBase, uint32(v14+int32(136)))) = v92
	*(*int32)(unsafe.Add(mBase, uint32(v14)+112)) = l0
	*(*int64)(unsafe.Add(mBase, uint32(v14)+120)) = v54
	v277 = F_iprintf(m, int32(_a2490), v14+int32(112))
	mBase = m.M
	v278 = m.ExcPending
	if v278 != 0 {
		goto L23
	} else {
		goto L83
	}
L83:
	;
	v281 = F_iprintf(m, int32(_a2491), int32(0))
	mBase = m.M
	v282 = m.ExcPending
	if v282 != 0 {
		goto L23
	} else {
		goto L84
	}
L84:
	;
	v287 = *(*int32)(unsafe.Add(mBase, _consts[913]))
	v288 = F_fgets(m, v14+int32(240), int32(2), v287)
	mBase = m.M
	v289 = m.ExcPending
	if v289 != 0 {
		goto L23
	} else {
		goto L85
	}
L85:
	;
	if v288 == int32(0) {
		goto L2
	} else {
		goto L86
	}
L86:
	;
	v293 = v14 + int32(240)
	v294 = int32(_a2492)
	goto L88
L87:
	;
	if v338-v340 != 0 {
		goto L2
	} else {
		goto L102
	}
L88:
	;
	v299 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v293))))
	if v299 != 0 {
		goto L91
	} else {
		goto L92
	}
L90:
	;
	v338 = F_tolower(m, v333)
	mBase = m.M
	v339 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v334))))
	v340 = F_tolower(m, v339)
	mBase = m.M
	goto L87
L91:
	;
	v301 = v293
	v302 = v294
	v303 = int32(1)
	v304 = v299
	goto L94
L92:
	;
	v333 = int32(0)
	v334 = v294
	goto L90
L93:
	;
	v333 = v330 & int32(255)
	v334 = v328
	goto L90
L94:
	;
	v306 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v302))))
	if v306 == int32(0) {
		v328 = v302
		v330 = v304
		goto L93
	} else {
		goto L96
	}
L95:
	;
	v328 = v322
	v330 = int32(0)
	goto L93
L96:
	;
	v310 = v303 + int32(-1)
	if v310 == int32(0) {
		v328 = v302
		v330 = v304
		goto L93
	} else {
		goto L97
	}
L97:
	;
	v314 = v304 & int32(255)
	if v314 == v306 {
		goto L98
	} else {
		goto L99
	}
L98:
	;
	v321 = int32(1)
	v322 = v302 + v321
	v323 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v301)+1)))
	if v323 != 0 {
		v301 = v301 + v321
		v302 = v322
		v303 = v310
		v304 = v323
		goto L94
	} else {
		goto L101
	}
L99:
	;
	v316 = F_tolower(m, v314)
	mBase = m.M
	v317 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v302))))
	v318 = F_tolower(m, v317)
	mBase = m.M
	if v316 == v318 {
		goto L98
	} else {
		goto L100
	}
L100:
	;
	v320 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v301))))
	v328 = v302
	v330 = v320
	goto L93
L101:
	;
	goto L95
L102:
	;
	v350 = *(*int32)(unsafe.Add(mBase, uint32(v19)+76))
	if int32(-1) < v350 {
		goto L105
	} else {
		goto L106
	}
L103:
	;
	v368 = F_ftruncate(m, v367, v92)
	mBase = m.M
	if v368 == int32(-1) {
		goto L1
	} else {
		goto L110
	}
L104:
	;
	if int32(-1) < v359 {
		v367 = v359
		goto L108
	} else {
		goto L109
	}
L105:
	;
	v354 = F___lockfile(m, v19)
	mBase = m.M
	v355 = *(*int32)(unsafe.Add(mBase, uint32(v19)+60))
	if v354 == int32(0) {
		v359 = v355
		goto L104
	} else {
		goto L107
	}
L106:
	;
	v353 = *(*int32)(unsafe.Add(mBase, uint32(v19)+60))
	v359 = v353
	goto L104
L107:
	;
	F___unlockfile(m, v19)
	mBase = m.M
	v359 = v355
	goto L104
L108:
	;
	goto L103
L109:
	;
	v363 = F___errno_location(m)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v363))) = int32(8)
	v367 = int32(-1)
	goto L108
L110:
	;
	v371 = F_fclose(m, v19)
	mBase = m.M
	v372 = m.ExcPending
	if v372 != 0 {
		goto L23
	} else {
		goto L111
	}
L111:
	;
	v392 = int32(2)
	goto L20
L112:
	;
	m.Env.Exit(m, int32(1))
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L113:
	;
	goto L71
L114:
	;
	v402 = *(*int32)(unsafe.Add(mBase, _consts[18]))
	v403 = F___strerror_l(m, v402, v402)
	mBase = m.M
	goto L115
L115:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+4)) = v403
	*(*int32)(unsafe.Add(mBase, uint32(v14))) = l1
	v407 = F_iprintf(m, int32(_a2493), v14)
	mBase = m.M
	v408 = m.ExcPending
	if v408 != 0 {
		goto L23
	} else {
		goto L116
	}
L116:
	;
	m.Env.Exit(m, int32(1))
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L117:
	;
	v417 = F_fclose(m, v19)
	mBase = m.M
	v418 = m.ExcPending
	if v418 != 0 {
		goto L23
	} else {
		goto L118
	}
L118:
	;
	m.Env.Exit(m, int32(1))
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L119:
	;
	m.Env.Exit(m, int32(1))
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L120:
	;
	v427 = *(*int32)(unsafe.Add(mBase, _consts[18]))
	v428 = F___strerror_l(m, v427, v427)
	mBase = m.M
	goto L121
L121:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+212)) = v428
	*(*int32)(unsafe.Add(mBase, uint32(v14)+208)) = l0
	v434 = F_iprintf(m, int32(_a2494), v14+int32(208))
	mBase = m.M
	v435 = m.ExcPending
	if v435 != 0 {
		goto L23
	} else {
		goto L122
	}
L122:
	;
	m.Env.Exit(m, int32(1))
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L123:
	;
	m.Env.Exit(m, int32(1))
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L124:
	;
	m.Env.Exit(m, int32(1))
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L125:
	;
	m.Env.Exit(m, int32(1))
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_cleanupThreadResources(m *base.Module, l0 int32) {
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	F_flushPendingIOResponses(m, int32(1))
	v4 = m.ExcPending
	if v4 != 0 {
		return
	} else {
		F_freeSharedQueryBuf(m)
		v6 = m.ExcPending
		if v6 != 0 {
			return
		} else {
			return
		}
	}
}
func F_cleanupTransferResources(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v16 int32
	_ = v16
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v55 int32
	_ = v55
	v4 = int32(_a20)
	v5 = *(*int32)(unsafe.Add(mBase, _consts[604]))
	v7 = *(*int32)(unsafe.Add(mBase, _consts[605]))
	if v7 != int32(-1) {
		if v5 == int32(0) {
			F__serverAssert(m, int32(_a1957), int32(_a1913), int32(4373))
			mBase = m.M
			v55 = m.ExcPending
			if v55 != 0 {
				return
			} else {
				F_abort(m)
				mBase = m.M
				base.Wasm_trap_unreachable()
				for {
				}
			}
		} else {
			v20 = F_close(m, v7)
			mBase = m.M
			v22 = *(*int32)(unsafe.Add(mBase, _consts[604]))
			v25 = F_open(m, v22, int32(2048), int32(0))
			mBase = m.M
			v26 = F_unlink(m, v22)
			mBase = m.M
			if v25 == int32(-1) {
				v42 = *(*int32)(unsafe.Add(mBase, _consts[604]))
				F_valkey_free(m, v42)
				mBase = m.M
				v44 = m.ExcPending
				if v44 != 0 {
					return
				} else {
					*(*int64)(unsafe.Add(mBase, _consts[605])) = int64(4294967295)
					return
				}
			} else {
				if v26 != int32(-1) {
					v35 = int32(0)
					F_bioCreateCloseJob(m, v25, v35, v35)
					mBase = m.M
					v38 = m.ExcPending
					if v38 != 0 {
						return
					} else {
						v42 = *(*int32)(unsafe.Add(mBase, _consts[604]))
						F_valkey_free(m, v42)
						mBase = m.M
						v44 = m.ExcPending
						if v44 != 0 {
							return
						} else {
							*(*int64)(unsafe.Add(mBase, _consts[605])) = int64(4294967295)
							return
						}
					}
				} else {
					v31 = int32(9116376)
					v32 = *(*int32)(unsafe.Add(mBase, _consts[18]))
					v33 = F_close(m, v25)
					mBase = m.M
					*(*int32)(unsafe.Add(mBase, _consts[18])) = v32
					v42 = *(*int32)(unsafe.Add(mBase, _consts[604]))
					F_valkey_free(m, v42)
					mBase = m.M
					v44 = m.ExcPending
					if v44 != 0 {
						return
					} else {
						*(*int64)(unsafe.Add(mBase, _consts[605])) = int64(4294967295)
						return
					}
				}
			}
		}
	} else {
		if v5 == int32(0) {
			return
		} else {
			F__serverAssert(m, int32(_a1958), int32(_a1913), int32(4369))
			mBase = m.M
			v16 = m.ExcPending
			if v16 != 0 {
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
func F_clock(m *base.Module) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v12 int64
	_ = v12
	var v17 int64
	_ = v17
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v26 int32
	_ = v26
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v9 = int32(-1)
	v11 = F___clock_gettime(m, int32(2), v7)
	mBase = m.M
	if v11 != 0 {
		v26 = v9
	} else {
		v12 = *(*int64)(unsafe.Add(mBase, uint32(v7)))
		if int64(2147) < v12 {
			v26 = v9
		} else {
			v17 = v12 * int64(1000000)
			v19 = *(*int32)(unsafe.Add(mBase, uint32(v7)+8))
			v21 = base.I32_div_s(v19, int32(1000))
			if int64(2147483647)-v17 < base.I64_extend_i32_s(v21) {
				v26 = v9
			} else {
				v26 = v21 + base.I32_wrap_i64(v17)
			}
		}
	}
	m.G0 = v7 + int32(16)
	return v26
}
func F_commandlogCommand(m *base.Module, l0 int32) {
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
	var v57 int64
	_ = v57
	var v62 int64
	_ = v62
	var v67 int64
	_ = v67
	var v72 int64
	_ = v72
	var v77 int64
	_ = v77
	var v80 int64
	_ = v80
	var v83 int64
	_ = v83
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
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v141 int32
	_ = v141
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v151 int32
	_ = v151
	var v153 int32
	_ = v153
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v163 int32
	_ = v163
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v175 int32
	_ = v175
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v181 int32
	_ = v181
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
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v221 int32
	_ = v221
	var v222 int64
	_ = v222
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v229 int32
	_ = v229
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
	var v271 int32
	_ = v271
	var v272 int32
	_ = v272
	var v276 int32
	_ = v276
	var v277 int32
	_ = v277
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
	var v293 int32
	_ = v293
	var v294 int32
	_ = v294
	var v296 int32
	_ = v296
	var v298 int32
	_ = v298
	var v300 int32
	_ = v300
	v5 = m.G0
	v7 = v5 - int32(64)
	m.G0 = v7
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v9 != int32(2) {
		v88 = v9
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v7 + int32(64)
	return
L2:
	;
	if v88 != int32(3) {
		v226 = v88
		goto L20
	} else {
		goto L21
	}
L3:
	;
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v13 = *(*int32)(unsafe.Add(mBase, uint32(v12)+4))
	v14 = F_objectGetVal(m, v13)
	mBase = m.M
	v15 = int32(_a98)
	v18 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14))))
	if v18 != 0 {
		goto L7
	} else {
		goto L8
	}
L4:
	;
	v87 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v88 = v87
	goto L2
L5:
	;
	if v50-v52 != 0 {
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
	v56 = int32(0)
	v57 = *(*int64)(unsafe.Add(mBase, _consts[236]))
	*(*int64)(unsafe.Add(mBase, uint32(v7+int32(48)))) = v57
	v62 = *(*int64)(unsafe.Add(mBase, _consts[237]))
	*(*int64)(unsafe.Add(mBase, uint32(v7+int32(40)))) = v62
	v67 = *(*int64)(unsafe.Add(mBase, _consts[238]))
	*(*int64)(unsafe.Add(mBase, uint32(v7+int32(32)))) = v67
	v72 = *(*int64)(unsafe.Add(mBase, _consts[239]))
	*(*int64)(unsafe.Add(mBase, uint32(v7+int32(24)))) = v72
	v77 = *(*int64)(unsafe.Add(mBase, _consts[240]))
	*(*int64)(unsafe.Add(mBase, uint32(v7+int32(16)))) = v77
	v80 = *(*int64)(unsafe.Add(mBase, _consts[241]))
	*(*int64)(unsafe.Add(mBase, uint32(v7)+8)) = v80
	v83 = *(*int64)(unsafe.Add(mBase, _consts[242]))
	*(*int64)(unsafe.Add(mBase, uint32(v7))) = v83
	F_addReplyHelp(m, l0, v7)
	mBase = m.M
	v86 = m.ExcPending
	if v86 != 0 {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	return
L19:
	;
	goto L1
L20:
	;
	if v226 != int32(4) {
		goto L63
	} else {
		goto L64
	}
L21:
	;
	v91 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v92 = *(*int32)(unsafe.Add(mBase, uint32(v91)+4))
	v93 = F_objectGetVal(m, v92)
	mBase = m.M
	v94 = int32(_a83)
	v97 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v93))))
	if v97 != 0 {
		goto L25
	} else {
		goto L26
	}
L22:
	;
	v166 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v166 != int32(3) {
		v226 = v166
		goto L20
	} else {
		goto L45
	}
L23:
	;
	if v129-v131 != 0 {
		goto L22
	} else {
		goto L35
	}
L24:
	;
	v129 = F_tolower(m, v125)
	mBase = m.M
	v130 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v126))))
	v131 = F_tolower(m, v130)
	mBase = m.M
	goto L23
L25:
	;
	v99 = v93
	v100 = v94
	v101 = v97
	goto L28
L26:
	;
	v125 = int32(0)
	v126 = v94
	goto L24
L27:
	;
	v125 = v122 & int32(255)
	v126 = v121
	goto L24
L28:
	;
	v103 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v100))))
	if v103 == int32(0) {
		v121 = v100
		v122 = v101
		goto L27
	} else {
		goto L30
	}
L29:
	;
	v121 = v115
	v122 = int32(0)
	goto L27
L30:
	;
	v107 = v101 & int32(255)
	if v107 == v103 {
		goto L31
	} else {
		goto L32
	}
L31:
	;
	v114 = int32(1)
	v115 = v100 + v114
	v116 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v99)+1)))
	if v116 != 0 {
		v99 = v99 + v114
		v100 = v115
		v101 = v116
		goto L28
	} else {
		goto L34
	}
L32:
	;
	v109 = F_tolower(m, v107)
	mBase = m.M
	v110 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v100))))
	v111 = F_tolower(m, v110)
	mBase = m.M
	if v109 == v111 {
		goto L31
	} else {
		goto L33
	}
L33:
	;
	v113 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v99))))
	v121 = v100
	v122 = v113
	goto L27
L34:
	;
	goto L29
L35:
	;
	v133 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v134 = *(*int32)(unsafe.Add(mBase, uint32(v133)+8))
	v135 = F_commandlogGetTypeOrReply(m, l0, v134)
	mBase = m.M
	v136 = m.ExcPending
	if v136 != 0 {
		goto L18
	} else {
		goto L36
	}
L36:
	;
	if v135 == int32(-1) {
		goto L1
	} else {
		goto L37
	}
L37:
	;
	v141 = v135 << (uint(int32(5)) % 32)
	v143 = *(*int32)(unsafe.Add(mBase, uint32(v141)+uint32(_consts[230])))
	v144 = *(*int32)(unsafe.Add(mBase, uint32(v143)+20))
	if v144 == int32(0) {
		goto L38
	} else {
		goto L39
	}
L38:
	;
	v163 = *(*int32)(unsafe.Add(mBase, _consts[27]))
	F_addReply(m, l0, v163)
	mBase = m.M
	v165 = m.ExcPending
	if v165 != 0 {
		goto L18
	} else {
		goto L44
	}
L39:
	;
	v151 = v143
	goto L40
L40:
	;
	v153 = *(*int32)(unsafe.Add(mBase, uint32(v151)+4))
	F_listDelNode(m, v151, v153)
	mBase = m.M
	v155 = m.ExcPending
	if v155 != 0 {
		goto L18
	} else {
		goto L42
	}
L41:
	;
	goto L38
L42:
	;
	v156 = *(*int32)(unsafe.Add(mBase, uint32(v141)+uint32(_consts[230])))
	v157 = *(*int32)(unsafe.Add(mBase, uint32(v156)+20))
	if v157 != 0 {
		v151 = v156
		goto L40
	} else {
		goto L43
	}
L43:
	;
	goto L41
L44:
	;
	goto L1
L45:
	;
	v169 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v170 = *(*int32)(unsafe.Add(mBase, uint32(v169)+4))
	v171 = F_objectGetVal(m, v170)
	mBase = m.M
	v172 = int32(_a462)
	v175 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v171))))
	if v175 != 0 {
		goto L49
	} else {
		goto L50
	}
L46:
	;
	v225 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v226 = v225
	goto L20
L47:
	;
	if v207-v209 != 0 {
		goto L46
	} else {
		goto L59
	}
L48:
	;
	v207 = F_tolower(m, v203)
	mBase = m.M
	v208 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v204))))
	v209 = F_tolower(m, v208)
	mBase = m.M
	goto L47
L49:
	;
	v177 = v171
	v178 = v172
	v179 = v175
	goto L52
L50:
	;
	v203 = int32(0)
	v204 = v172
	goto L48
L51:
	;
	v203 = v200 & int32(255)
	v204 = v199
	goto L48
L52:
	;
	v181 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v178))))
	if v181 == int32(0) {
		v199 = v178
		v200 = v179
		goto L51
	} else {
		goto L54
	}
L53:
	;
	v199 = v193
	v200 = int32(0)
	goto L51
L54:
	;
	v185 = v179 & int32(255)
	if v185 == v181 {
		goto L55
	} else {
		goto L56
	}
L55:
	;
	v192 = int32(1)
	v193 = v178 + v192
	v194 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v177)+1)))
	if v194 != 0 {
		v177 = v177 + v192
		v178 = v193
		v179 = v194
		goto L52
	} else {
		goto L58
	}
L56:
	;
	v187 = F_tolower(m, v185)
	mBase = m.M
	v188 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v178))))
	v189 = F_tolower(m, v188)
	mBase = m.M
	if v187 == v189 {
		goto L55
	} else {
		goto L57
	}
L57:
	;
	v191 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v177))))
	v199 = v178
	v200 = v191
	goto L51
L58:
	;
	goto L53
L59:
	;
	v211 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v212 = *(*int32)(unsafe.Add(mBase, uint32(v211)+8))
	v213 = F_commandlogGetTypeOrReply(m, l0, v212)
	mBase = m.M
	v214 = m.ExcPending
	if v214 != 0 {
		goto L18
	} else {
		goto L60
	}
L60:
	;
	if v213 == int32(-1) {
		goto L1
	} else {
		goto L61
	}
L61:
	;
	v221 = *(*int32)(unsafe.Add(mBase, uint32(v213<<(uint(int32(5))%32))+uint32(_consts[230])))
	v222 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v221)+20)))
	F_addReplyLongLong(m, l0, v222)
	mBase = m.M
	v224 = m.ExcPending
	if v224 != 0 {
		goto L18
	} else {
		goto L62
	}
L62:
	;
	goto L1
L63:
	;
	F_addReplySubcommandSyntaxError(m, l0)
	mBase = m.M
	v300 = m.ExcPending
	if v300 != 0 {
		goto L18
	} else {
		goto L85
	}
L64:
	;
	v229 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v230 = *(*int32)(unsafe.Add(mBase, uint32(v229)+4))
	v231 = F_objectGetVal(m, v230)
	mBase = m.M
	v232 = int32(_a463)
	v235 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v231))))
	if v235 != 0 {
		goto L67
	} else {
		goto L68
	}
L65:
	;
	if v267-v269 != 0 {
		goto L63
	} else {
		goto L77
	}
L66:
	;
	v267 = F_tolower(m, v263)
	mBase = m.M
	v268 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v264))))
	v269 = F_tolower(m, v268)
	mBase = m.M
	goto L65
L67:
	;
	v237 = v231
	v238 = v232
	v239 = v235
	goto L70
L68:
	;
	v263 = int32(0)
	v264 = v232
	goto L66
L69:
	;
	v263 = v260 & int32(255)
	v264 = v259
	goto L66
L70:
	;
	v241 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v238))))
	if v241 == int32(0) {
		v259 = v238
		v260 = v239
		goto L69
	} else {
		goto L72
	}
L71:
	;
	v259 = v253
	v260 = int32(0)
	goto L69
L72:
	;
	v245 = v239 & int32(255)
	if v245 == v241 {
		goto L73
	} else {
		goto L74
	}
L73:
	;
	v252 = int32(1)
	v253 = v238 + v252
	v254 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v237)+1)))
	if v254 != 0 {
		v237 = v237 + v252
		v238 = v253
		v239 = v254
		goto L70
	} else {
		goto L76
	}
L74:
	;
	v247 = F_tolower(m, v245)
	mBase = m.M
	v248 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v238))))
	v249 = F_tolower(m, v248)
	mBase = m.M
	if v247 == v249 {
		goto L73
	} else {
		goto L75
	}
L75:
	;
	v251 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v237))))
	v259 = v238
	v260 = v251
	goto L69
L76:
	;
	goto L71
L77:
	;
	v271 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v272 = *(*int32)(unsafe.Add(mBase, uint32(v271)+8))
	v276 = F_getRangeLongFromObjectOrReply(m, l0, v272, int32(-1), int32(2147483647), v7, int32(_a464))
	mBase = m.M
	v277 = m.ExcPending
	if v277 != 0 {
		goto L18
	} else {
		goto L78
	}
L78:
	;
	if v276 != 0 {
		goto L1
	} else {
		goto L79
	}
L79:
	;
	v278 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v279 = *(*int32)(unsafe.Add(mBase, uint32(v278)+12))
	v280 = F_commandlogGetTypeOrReply(m, l0, v279)
	mBase = m.M
	v281 = m.ExcPending
	if v281 != 0 {
		goto L18
	} else {
		goto L80
	}
L80:
	;
	if v280 == int32(-1) {
		goto L1
	} else {
		goto L81
	}
L81:
	;
	v284 = *(*int32)(unsafe.Add(mBase, uint32(v7)))
	if v284 != int32(-1) {
		v296 = v284
		goto L82
	} else {
		goto L83
	}
L82:
	;
	F_commandlogGetReply(m, l0, v280, v296)
	mBase = m.M
	v298 = m.ExcPending
	if v298 != 0 {
		goto L18
	} else {
		goto L84
	}
L83:
	;
	v293 = *(*int32)(unsafe.Add(mBase, uint32(v280<<(uint(int32(5))%32))+uint32(_consts[230])))
	v294 = *(*int32)(unsafe.Add(mBase, uint32(v293)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v7))) = v294
	v296 = v294
	goto L82
L84:
	;
	goto L1
L85:
	;
	goto L1
}
func F_commandlogFreeEntry(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v3 < int32(1) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	F_valkey_free(m, v22)
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		goto L5
	} else {
		goto L8
	}
L2:
	;
	v8 = int32(0)
	goto L3
L3:
	;
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v13 = *(*int32)(unsafe.Add(mBase, uint32(v9+v8<<(uint(int32(2))%32))))
	F_decrRefCount(m, v13)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		goto L5
	} else {
		goto L6
	}
L4:
	;
	goto L1
L5:
	;
	return
L6:
	;
	v17 = v8 + int32(1)
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v17 < v18 {
		v8 = v17
		goto L3
	} else {
		goto L7
	}
L7:
	;
	goto L4
L8:
	;
	v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	F_sdsfree(m, v25)
	mBase = m.M
	v27 = m.ExcPending
	if v27 != 0 {
		goto L5
	} else {
		goto L9
	}
L9:
	;
	v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	F_sdsfree(m, v28)
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		goto L5
	} else {
		goto L10
	}
L10:
	;
	F_valkey_free(m, l0)
	mBase = m.M
	v32 = m.ExcPending
	if v32 != 0 {
		goto L5
	} else {
		goto L11
	}
L11:
	;
	return
}
func F_commandlogGetReply(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v15 int32
	_ = v15
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v60 int64
	_ = v60
	var v62 int32
	_ = v62
	var v63 int64
	_ = v63
	var v65 int32
	_ = v65
	var v66 int64
	_ = v66
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v78 int32
	_ = v78
	var v84 int32
	_ = v84
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v103 int32
	_ = v103
	var v105 int32
	_ = v105
	var v108 int32
	_ = v108
	var v115 int32
	_ = v115
	var v118 int32
	_ = v118
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
	var v131 int32
	_ = v131
	var v138 int32
	_ = v138
	var v141 int32
	_ = v141
	var v144 int32
	_ = v144
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v150 int32
	_ = v150
	var v152 int32
	_ = v152
	v9 = m.G0
	v11 = v9 - int32(16)
	m.G0 = v11
	v15 = l1 << (uint(int32(5)) % 32)
	v19 = *(*int32)(unsafe.Add(mBase, uint32(v15)+uint32(_consts[230])))
	v20 = *(*int32)(unsafe.Add(mBase, uint32(v19)+20))
	if l2 < v20 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v22 = l2
	goto L3
L2:
	;
	v22 = v20
	goto L3
L3:
	;
	F_addReplyArrayLen(m, l0, v22)
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	return
L5:
	;
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v15)+uint32(_consts[230])))
	v27 = v11 + int32(8)
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v25)))
	*(*int32)(unsafe.Add(mBase, uint32(v27)+4)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v27))) = v28
	goto L6
L6:
	;
	if v22 == int32(0) {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	m.G0 = v11 + int32(16)
	return
L8:
	;
	v39 = v22
	goto L9
L9:
	;
	v43 = v11 + int32(8)
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v43)))
	if v45 == int32(0) {
		goto L12
	} else {
		goto L13
	}
L10:
	;
	goto L7
L11:
	;
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v45)+8))
	F_addReplyArrayLen(m, l0, int32(6))
	mBase = m.M
	v59 = m.ExcPending
	if v59 != 0 {
		goto L4
	} else {
		goto L14
	}
L12:
	;
	goto L11
L13:
	;
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v43)+4))
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v45+base.B2i32(v48 == int32(0))<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v43))) = v54
	goto L12
L14:
	;
	v60 = *(*int64)(unsafe.Add(mBase, uint32(v56)+8))
	F_addReplyLongLong(m, l0, v60)
	mBase = m.M
	v62 = m.ExcPending
	if v62 != 0 {
		goto L4
	} else {
		goto L15
	}
L15:
	;
	v63 = *(*int64)(unsafe.Add(mBase, uint32(v56)+24))
	F_addReplyLongLong(m, l0, v63)
	mBase = m.M
	v65 = m.ExcPending
	if v65 != 0 {
		goto L4
	} else {
		goto L16
	}
L16:
	;
	v66 = *(*int64)(unsafe.Add(mBase, uint32(v56)+16))
	F_addReplyLongLong(m, l0, v66)
	mBase = m.M
	v68 = m.ExcPending
	if v68 != 0 {
		goto L4
	} else {
		goto L17
	}
L17:
	;
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v56)+4))
	F_addReplyArrayLen(m, l0, v69)
	mBase = m.M
	v71 = m.ExcPending
	if v71 != 0 {
		goto L4
	} else {
		goto L18
	}
L18:
	;
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v56)+4))
	if v73 < int32(1) {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	v103 = int32(0)
	v105 = *(*int32)(unsafe.Add(mBase, uint32(v56)+36))
	v108 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v105+int32(-1)))))
	switch v108 & int32(7) {
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
		v125 = v103
		goto L25
	}
L20:
	;
	v78 = int32(0)
	goto L21
L21:
	;
	v84 = *(*int32)(unsafe.Add(mBase, uint32(v56)))
	v88 = *(*int32)(unsafe.Add(mBase, uint32(v84+v78<<(uint(int32(2))%32))))
	F_addReplyBulk(m, l0, v88)
	mBase = m.M
	v90 = m.ExcPending
	if v90 != 0 {
		goto L4
	} else {
		goto L23
	}
L22:
	;
	goto L19
L23:
	;
	v92 = v78 + int32(1)
	v93 = *(*int32)(unsafe.Add(mBase, uint32(v56)+4))
	if v92 < v93 {
		v78 = v92
		goto L21
	} else {
		goto L24
	}
L24:
	;
	goto L22
L25:
	;
	F_addReplyBulkCBuffer(m, l0, v105, v125)
	mBase = m.M
	v127 = m.ExcPending
	if v127 != 0 {
		goto L4
	} else {
		goto L31
	}
L26:
	;
	v124 = *(*int32)(unsafe.Add(mBase, uint32(v105+int32(-17))))
	v125 = v124
	goto L25
L27:
	;
	v121 = *(*int32)(unsafe.Add(mBase, uint32(v105+int32(-9))))
	v125 = v121
	goto L25
L28:
	;
	v118 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v105+int32(-5)))))
	v125 = v118
	goto L25
L29:
	;
	v115 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v105+int32(-3)))))
	v125 = v115
	goto L25
L30:
	;
	v125 = int32(base.Ui32(v108) >> (uint(int32(3)) % 32))
	goto L25
L31:
	;
	v128 = *(*int32)(unsafe.Add(mBase, uint32(v56)+32))
	v131 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v128+int32(-1)))))
	switch v131 & int32(7) {
	case 0:
		goto L37
	case 1:
		goto L36
	case 2:
		goto L35
	case 3:
		goto L34
	case 4:
		goto L33
	default:
		v148 = v103
		goto L32
	}
L32:
	;
	F_addReplyBulkCBuffer(m, l0, v128, v148)
	mBase = m.M
	v150 = m.ExcPending
	if v150 != 0 {
		goto L4
	} else {
		goto L38
	}
L33:
	;
	v147 = *(*int32)(unsafe.Add(mBase, uint32(v128+int32(-17))))
	v148 = v147
	goto L32
L34:
	;
	v144 = *(*int32)(unsafe.Add(mBase, uint32(v128+int32(-9))))
	v148 = v144
	goto L32
L35:
	;
	v141 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v128+int32(-5)))))
	v148 = v141
	goto L32
L36:
	;
	v138 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v128+int32(-3)))))
	v148 = v138
	goto L32
L37:
	;
	v148 = int32(base.Ui32(v131) >> (uint(int32(3)) % 32))
	goto L32
L38:
	;
	v152 = v39 + int32(-1)
	if v152 != 0 {
		v39 = v152
		goto L9
	} else {
		goto L39
	}
L39:
	;
	goto L10
}
func F_commandlogGetTypeOrReply(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
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
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v133 int32
	_ = v133
	v3 = F_objectGetVal(m, l1)
	mBase = m.M
	v4 = int32(_a465)
	v7 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3))))
	if v7 != 0 {
		goto L4
	} else {
		goto L5
	}
L1:
	;
	v45 = F_objectGetVal(m, l1)
	mBase = m.M
	v46 = int32(_a466)
	v49 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v45))))
	if v49 != 0 {
		goto L18
	} else {
		goto L19
	}
L2:
	;
	if v39-v41 != 0 {
		goto L1
	} else {
		goto L14
	}
L3:
	;
	v39 = F_tolower(m, v35)
	mBase = m.M
	v40 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v36))))
	v41 = F_tolower(m, v40)
	mBase = m.M
	goto L2
L4:
	;
	v9 = v3
	v10 = v4
	v11 = v7
	goto L7
L5:
	;
	v35 = int32(0)
	v36 = v4
	goto L3
L6:
	;
	v35 = v32 & int32(255)
	v36 = v31
	goto L3
L7:
	;
	v13 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10))))
	if v13 == int32(0) {
		v31 = v10
		v32 = v11
		goto L6
	} else {
		goto L9
	}
L8:
	;
	v31 = v25
	v32 = int32(0)
	goto L6
L9:
	;
	v17 = v11 & int32(255)
	if v17 == v13 {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	v24 = int32(1)
	v25 = v10 + v24
	v26 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+1)))
	if v26 != 0 {
		v9 = v9 + v24
		v10 = v25
		v11 = v26
		goto L7
	} else {
		goto L13
	}
L11:
	;
	v19 = F_tolower(m, v17)
	mBase = m.M
	v20 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10))))
	v21 = F_tolower(m, v20)
	mBase = m.M
	if v19 == v21 {
		goto L10
	} else {
		goto L12
	}
L12:
	;
	v23 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9))))
	v31 = v10
	v32 = v23
	goto L6
L13:
	;
	goto L8
L14:
	;
	return int32(0)
L15:
	;
	v87 = F_objectGetVal(m, l1)
	mBase = m.M
	v88 = int32(_a467)
	v91 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v87))))
	if v91 != 0 {
		goto L32
	} else {
		goto L33
	}
L16:
	;
	if v81-v83 != 0 {
		goto L15
	} else {
		goto L28
	}
L17:
	;
	v81 = F_tolower(m, v77)
	mBase = m.M
	v82 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v78))))
	v83 = F_tolower(m, v82)
	mBase = m.M
	goto L16
L18:
	;
	v51 = v45
	v52 = v46
	v53 = v49
	goto L21
L19:
	;
	v77 = int32(0)
	v78 = v46
	goto L17
L20:
	;
	v77 = v74 & int32(255)
	v78 = v73
	goto L17
L21:
	;
	v55 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v52))))
	if v55 == int32(0) {
		v73 = v52
		v74 = v53
		goto L20
	} else {
		goto L23
	}
L22:
	;
	v73 = v67
	v74 = int32(0)
	goto L20
L23:
	;
	v59 = v53 & int32(255)
	if v59 == v55 {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	v66 = int32(1)
	v67 = v52 + v66
	v68 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v51)+1)))
	if v68 != 0 {
		v51 = v51 + v66
		v52 = v67
		v53 = v68
		goto L21
	} else {
		goto L27
	}
L25:
	;
	v61 = F_tolower(m, v59)
	mBase = m.M
	v62 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v52))))
	v63 = F_tolower(m, v62)
	mBase = m.M
	if v61 == v63 {
		goto L24
	} else {
		goto L26
	}
L26:
	;
	v65 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v51))))
	v73 = v52
	v74 = v65
	goto L20
L27:
	;
	goto L22
L28:
	;
	return int32(1)
L29:
	;
	F_addReplyError(m, l0, int32(_a468))
	mBase = m.M
	v133 = m.ExcPending
	if v133 != 0 {
		goto L43
	} else {
		goto L44
	}
L30:
	;
	if v123-v125 != 0 {
		goto L29
	} else {
		goto L42
	}
L31:
	;
	v123 = F_tolower(m, v119)
	mBase = m.M
	v124 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v120))))
	v125 = F_tolower(m, v124)
	mBase = m.M
	goto L30
L32:
	;
	v93 = v87
	v94 = v88
	v95 = v91
	goto L35
L33:
	;
	v119 = int32(0)
	v120 = v88
	goto L31
L34:
	;
	v119 = v116 & int32(255)
	v120 = v115
	goto L31
L35:
	;
	v97 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v94))))
	if v97 == int32(0) {
		v115 = v94
		v116 = v95
		goto L34
	} else {
		goto L37
	}
L36:
	;
	v115 = v109
	v116 = int32(0)
	goto L34
L37:
	;
	v101 = v95 & int32(255)
	if v101 == v97 {
		goto L38
	} else {
		goto L39
	}
L38:
	;
	v108 = int32(1)
	v109 = v94 + v108
	v110 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v93)+1)))
	if v110 != 0 {
		v93 = v93 + v108
		v94 = v109
		v95 = v110
		goto L35
	} else {
		goto L41
	}
L39:
	;
	v103 = F_tolower(m, v101)
	mBase = m.M
	v104 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v94))))
	v105 = F_tolower(m, v104)
	mBase = m.M
	if v103 == v105 {
		goto L38
	} else {
		goto L40
	}
L40:
	;
	v107 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v93))))
	v115 = v94
	v116 = v107
	goto L34
L41:
	;
	goto L36
L42:
	;
	return int32(2)
L43:
	;
	return int32(0)
L44:
	;
	return int32(-1)
}
func F_compareReplicasForPromotion(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v14 int64
	_ = v14
	var v15 int64
	_ = v15
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(v7)+176))
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(v9)+176))
	if v8 == v10 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v14 = *(*int64)(unsafe.Add(mBase, uint32(v7)+208))
	v15 = *(*int64)(unsafe.Add(mBase, uint32(v9)+208))
	if base.Ui64(v14) <= base.Ui64(v15) {
		goto L3
	} else {
		goto L4
	}
L2:
	;
	return v8 - v10
L3:
	;
	if base.Ui64(v15) <= base.Ui64(v14) {
		goto L5
	} else {
		goto L6
	}
L4:
	;
	return int32(-1)
L5:
	;
	v23 = *(*int32)(unsafe.Add(mBase, uint32(v7)+8))
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v9)+8))
	if v23 != 0 {
		goto L7
	} else {
		goto L8
	}
L6:
	;
	return int32(1)
L7:
	;
	v28 = int32(-1)
	goto L9
L8:
	;
	v28 = base.B2i32(v23|v24 != int32(0))
	goto L9
L9:
	;
	if v23 == int32(0) {
		v71 = v28
		goto L10
	} else {
		goto L11
	}
L10:
	;
	return v71
L11:
	;
	if v24 == int32(0) {
		v71 = v28
		goto L10
	} else {
		goto L12
	}
L12:
	;
	v35 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23))))
	if v35 != 0 {
		goto L15
	} else {
		goto L16
	}
L13:
	;
	v71 = v67 - v69
	goto L10
L14:
	;
	v67 = F_tolower(m, v63)
	mBase = m.M
	v68 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v64))))
	v69 = F_tolower(m, v68)
	mBase = m.M
	goto L13
L15:
	;
	v37 = v23
	v38 = v24
	v39 = v35
	goto L18
L16:
	;
	v63 = int32(0)
	v64 = v24
	goto L14
L17:
	;
	v63 = v60 & int32(255)
	v64 = v59
	goto L14
L18:
	;
	v41 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v38))))
	if v41 == int32(0) {
		v59 = v38
		v60 = v39
		goto L17
	} else {
		goto L20
	}
L19:
	;
	v59 = v53
	v60 = int32(0)
	goto L17
L20:
	;
	v45 = v39 & int32(255)
	if v45 == v41 {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	v52 = int32(1)
	v53 = v38 + v52
	v54 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v37)+1)))
	if v54 != 0 {
		v37 = v37 + v52
		v38 = v53
		v39 = v54
		goto L18
	} else {
		goto L24
	}
L22:
	;
	v47 = F_tolower(m, v45)
	mBase = m.M
	v48 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v38))))
	v49 = F_tolower(m, v48)
	mBase = m.M
	if v47 == v49 {
		goto L21
	} else {
		goto L23
	}
L23:
	;
	v51 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v37))))
	v59 = v38
	v60 = v51
	goto L17
L24:
	;
	goto L19
}
func F_computeDatasetProfile(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int64) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v56 int32
	_ = v56
	var v58 int64
	_ = v58
	var v60 int32
	_ = v60
	var v64 int32
	_ = v64
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v81 int64
	_ = v81
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v95 int64
	_ = v95
	var v99 int32
	_ = v99
	var v102 int32
	_ = v102
	var v110 int32
	_ = v110
	var v113 int32
	_ = v113
	var v115 int32
	_ = v115
	var v118 int32
	_ = v118
	var v121 int32
	_ = v121
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
	var v133 int64
	_ = v133
	var v134 int64
	_ = v134
	var v137 int64
	_ = v137
	var v141 int64
	_ = v141
	var v143 int64
	_ = v143
	var v145 int64
	_ = v145
	var v147 int64
	_ = v147
	var v163 int32
	_ = v163
	var v166 int32
	_ = v166
	var v169 int32
	_ = v169
	var v172 int32
	_ = v172
	var v173 int64
	_ = v173
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v187 int64
	_ = v187
	var v191 int32
	_ = v191
	var v194 int32
	_ = v194
	var v202 int32
	_ = v202
	var v205 int32
	_ = v205
	var v207 int32
	_ = v207
	var v210 int32
	_ = v210
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
	var v224 int32
	_ = v224
	var v225 int64
	_ = v225
	var v226 int64
	_ = v226
	var v229 int64
	_ = v229
	var v233 int64
	_ = v233
	var v235 int64
	_ = v235
	var v237 int64
	_ = v237
	var v239 int64
	_ = v239
	var v242 int64
	_ = v242
	var v257 int32
	_ = v257
	var v258 int32
	_ = v258
	var v261 int32
	_ = v261
	var v262 int32
	_ = v262
	var v276 int32
	_ = v276
	var v277 int32
	_ = v277
	var v278 int32
	_ = v278
	var v279 int32
	_ = v279
	var v280 int32
	_ = v280
	var v283 int32
	_ = v283
	var v286 int32
	_ = v286
	var v289 int32
	_ = v289
	var v290 int64
	_ = v290
	var v301 int32
	_ = v301
	var v302 int32
	_ = v302
	var v304 int64
	_ = v304
	var v308 int32
	_ = v308
	var v311 int32
	_ = v311
	var v319 int32
	_ = v319
	var v322 int32
	_ = v322
	var v324 int32
	_ = v324
	var v327 int32
	_ = v327
	var v330 int32
	_ = v330
	var v332 int32
	_ = v332
	var v335 int32
	_ = v335
	var v337 int32
	_ = v337
	var v338 int32
	_ = v338
	var v341 int32
	_ = v341
	var v342 int64
	_ = v342
	var v343 int64
	_ = v343
	var v346 int64
	_ = v346
	var v350 int64
	_ = v350
	var v352 int64
	_ = v352
	var v354 int64
	_ = v354
	var v356 int64
	_ = v356
	var v359 int64
	_ = v359
	var v373 int32
	_ = v373
	var v376 int32
	_ = v376
	var v377 int32
	_ = v377
	var v388 int32
	_ = v388
	var v389 int32
	_ = v389
	var v390 int32
	_ = v390
	var v393 int32
	_ = v393
	var v396 int32
	_ = v396
	var v399 int32
	_ = v399
	var v400 int64
	_ = v400
	var v411 int32
	_ = v411
	var v412 int32
	_ = v412
	var v414 int64
	_ = v414
	var v418 int32
	_ = v418
	var v421 int32
	_ = v421
	var v429 int32
	_ = v429
	var v432 int32
	_ = v432
	var v434 int32
	_ = v434
	var v437 int32
	_ = v437
	var v440 int32
	_ = v440
	var v442 int32
	_ = v442
	var v445 int32
	_ = v445
	var v447 int32
	_ = v447
	var v448 int32
	_ = v448
	var v451 int32
	_ = v451
	var v452 int64
	_ = v452
	var v453 int64
	_ = v453
	var v456 int64
	_ = v456
	var v460 int64
	_ = v460
	var v462 int64
	_ = v462
	var v464 int64
	_ = v464
	var v466 int64
	_ = v466
	var v469 int64
	_ = v469
	var v482 int32
	_ = v482
	var v483 int32
	_ = v483
	var v484 int32
	_ = v484
	var v485 int32
	_ = v485
	var v489 int32
	_ = v489
	var v500 int32
	_ = v500
	var v507 int32
	_ = v507
	var v510 int32
	_ = v510
	var v513 int32
	_ = v513
	var v516 int32
	_ = v516
	var v517 int32
	_ = v517
	var v518 int32
	_ = v518
	var v521 int32
	_ = v521
	var v524 int32
	_ = v524
	var v527 int32
	_ = v527
	var v528 int64
	_ = v528
	var v539 int32
	_ = v539
	var v540 int32
	_ = v540
	var v542 int64
	_ = v542
	var v546 int32
	_ = v546
	var v549 int32
	_ = v549
	var v557 int32
	_ = v557
	var v560 int32
	_ = v560
	var v562 int32
	_ = v562
	var v565 int32
	_ = v565
	var v568 int32
	_ = v568
	var v570 int32
	_ = v570
	var v573 int32
	_ = v573
	var v575 int32
	_ = v575
	var v576 int32
	_ = v576
	var v579 int32
	_ = v579
	var v580 int64
	_ = v580
	var v581 int64
	_ = v581
	var v584 int64
	_ = v584
	var v588 int64
	_ = v588
	var v590 int64
	_ = v590
	var v592 int64
	_ = v592
	var v594 int64
	_ = v594
	var v597 int64
	_ = v597
	var v611 int32
	_ = v611
	var v612 int32
	_ = v612
	var v613 int32
	_ = v613
	var v624 int32
	_ = v624
	var v625 int32
	_ = v625
	var v626 int32
	_ = v626
	var v629 int32
	_ = v629
	var v632 int32
	_ = v632
	var v635 int32
	_ = v635
	var v636 int64
	_ = v636
	var v647 int32
	_ = v647
	var v648 int32
	_ = v648
	var v650 int64
	_ = v650
	var v654 int32
	_ = v654
	var v657 int32
	_ = v657
	var v665 int32
	_ = v665
	var v668 int32
	_ = v668
	var v670 int32
	_ = v670
	var v673 int32
	_ = v673
	var v676 int32
	_ = v676
	var v678 int32
	_ = v678
	var v681 int32
	_ = v681
	var v683 int32
	_ = v683
	var v684 int32
	_ = v684
	var v687 int32
	_ = v687
	var v688 int64
	_ = v688
	var v689 int64
	_ = v689
	var v692 int64
	_ = v692
	var v696 int64
	_ = v696
	var v698 int64
	_ = v698
	var v700 int64
	_ = v700
	var v702 int64
	_ = v702
	var v705 int64
	_ = v705
	var v724 int32
	_ = v724
	var v726 int32
	_ = v726
	var v727 int32
	_ = v727
	var v731 int32
	_ = v731
	var v732 int32
	_ = v732
	var v741 int32
	_ = v741
	var v747 int32
	_ = v747
	var v748 int32
	_ = v748
	var v749 int32
	_ = v749
	var v750 float64
	_ = v750
	var v751 int32
	_ = v751
	var v754 int32
	_ = v754
	var v756 int32
	_ = v756
	var v758 int64
	_ = v758
	var v767 int32
	_ = v767
	var v776 int32
	_ = v776
	var v777 int32
	_ = v777
	var v778 int64
	_ = v778
	var v780 int32
	_ = v780
	var v801 int32
	_ = v801
	var v809 int32
	_ = v809
	var v812 int32
	_ = v812
	var v816 int32
	_ = v816
	var v821 int32
	_ = v821
	var v823 int32
	_ = v823
	var v827 int32
	_ = v827
	var v833 int32
	_ = v833
	var v836 int32
	_ = v836
	var v842 int32
	_ = v842
	var v846 int32
	_ = v846
	var v848 int32
	_ = v848
	var v856 int32
	_ = v856
	var v857 int32
	_ = v857
	var v859 int32
	_ = v859
	var v862 int32
	_ = v862
	var v864 int32
	_ = v864
	var v866 int32
	_ = v866
	var v876 int32
	_ = v876
	var v879 int32
	_ = v879
	var v883 int32
	_ = v883
	var v888 int32
	_ = v888
	var v890 int32
	_ = v890
	var v894 int32
	_ = v894
	var v900 int32
	_ = v900
	var v903 int32
	_ = v903
	var v909 int32
	_ = v909
	var v913 int32
	_ = v913
	var v915 int32
	_ = v915
	var v923 int32
	_ = v923
	var v924 int32
	_ = v924
	var v927 int32
	_ = v927
	var v930 int32
	_ = v930
	var v933 int32
	_ = v933
	var v934 int64
	_ = v934
	var v945 int32
	_ = v945
	var v946 int32
	_ = v946
	var v948 int64
	_ = v948
	var v952 int32
	_ = v952
	var v955 int32
	_ = v955
	var v963 int32
	_ = v963
	var v966 int32
	_ = v966
	var v968 int32
	_ = v968
	var v971 int32
	_ = v971
	var v974 int32
	_ = v974
	var v976 int32
	_ = v976
	var v979 int32
	_ = v979
	var v981 int32
	_ = v981
	var v982 int32
	_ = v982
	var v985 int32
	_ = v985
	var v986 int64
	_ = v986
	var v987 int64
	_ = v987
	var v990 int64
	_ = v990
	var v994 int64
	_ = v994
	var v996 int64
	_ = v996
	var v998 int64
	_ = v998
	var v1000 int64
	_ = v1000
	var v1003 int64
	_ = v1003
	var v1021 int32
	_ = v1021
	var v1022 int32
	_ = v1022
	var v1023 int32
	_ = v1023
	var v1024 int32
	_ = v1024
	var v1025 int32
	_ = v1025
	var v1028 int32
	_ = v1028
	var v1031 int32
	_ = v1031
	var v1034 int32
	_ = v1034
	var v1035 int64
	_ = v1035
	var v1046 int32
	_ = v1046
	var v1047 int32
	_ = v1047
	var v1049 int64
	_ = v1049
	var v1053 int32
	_ = v1053
	var v1056 int32
	_ = v1056
	var v1064 int32
	_ = v1064
	var v1067 int32
	_ = v1067
	var v1069 int32
	_ = v1069
	var v1072 int32
	_ = v1072
	var v1075 int32
	_ = v1075
	var v1077 int32
	_ = v1077
	var v1080 int32
	_ = v1080
	var v1082 int32
	_ = v1082
	var v1083 int32
	_ = v1083
	var v1086 int32
	_ = v1086
	var v1087 int64
	_ = v1087
	var v1088 int64
	_ = v1088
	var v1091 int64
	_ = v1091
	var v1095 int64
	_ = v1095
	var v1097 int64
	_ = v1097
	var v1099 int64
	_ = v1099
	var v1101 int64
	_ = v1101
	var v1104 int64
	_ = v1104
	var v1118 int32
	_ = v1118
	var v1119 int32
	_ = v1119
	var v1120 int32
	_ = v1120
	var v1121 int32
	_ = v1121
	var v1143 int32
	_ = v1143
	var v1144 int32
	_ = v1144
	var v1156 int32
	_ = v1156
	var v1158 int32
	_ = v1158
	var v1159 int32
	_ = v1159
	var v1160 float64
	_ = v1160
	var v1163 int32
	_ = v1163
	var v1168 int32
	_ = v1168
	var v1169 int32
	_ = v1169
	var v1172 int32
	_ = v1172
	var v1173 int32
	_ = v1173
	var v1174 int32
	_ = v1174
	var v1179 int32
	_ = v1179
	var v1186 int32
	_ = v1186
	var v1189 int32
	_ = v1189
	var v1192 int32
	_ = v1192
	var v1195 int32
	_ = v1195
	var v1196 int32
	_ = v1196
	var v1197 int32
	_ = v1197
	var v1199 int32
	_ = v1199
	var v1207 int32
	_ = v1207
	var v1210 int32
	_ = v1210
	var v1214 int32
	_ = v1214
	var v1219 int32
	_ = v1219
	var v1221 int32
	_ = v1221
	var v1225 int32
	_ = v1225
	var v1231 int32
	_ = v1231
	var v1234 int32
	_ = v1234
	var v1240 int32
	_ = v1240
	var v1244 int32
	_ = v1244
	var v1246 int32
	_ = v1246
	var v1254 int32
	_ = v1254
	var v1255 int32
	_ = v1255
	var v1258 int32
	_ = v1258
	var v1261 int32
	_ = v1261
	var v1264 int32
	_ = v1264
	var v1265 int64
	_ = v1265
	var v1276 int32
	_ = v1276
	var v1277 int32
	_ = v1277
	var v1279 int64
	_ = v1279
	var v1283 int32
	_ = v1283
	var v1286 int32
	_ = v1286
	var v1294 int32
	_ = v1294
	var v1297 int32
	_ = v1297
	var v1299 int32
	_ = v1299
	var v1302 int32
	_ = v1302
	var v1305 int32
	_ = v1305
	var v1307 int32
	_ = v1307
	var v1310 int32
	_ = v1310
	var v1312 int32
	_ = v1312
	var v1313 int32
	_ = v1313
	var v1316 int32
	_ = v1316
	var v1317 int64
	_ = v1317
	var v1318 int64
	_ = v1318
	var v1321 int64
	_ = v1321
	var v1325 int64
	_ = v1325
	var v1327 int64
	_ = v1327
	var v1329 int64
	_ = v1329
	var v1331 int64
	_ = v1331
	var v1334 int64
	_ = v1334
	var v1351 int32
	_ = v1351
	var v1352 int32
	_ = v1352
	var v1365 int32
	_ = v1365
	var v1366 int32
	_ = v1366
	var v1367 int32
	_ = v1367
	var v1368 int32
	_ = v1368
	var v1369 int32
	_ = v1369
	var v1372 int32
	_ = v1372
	var v1375 int32
	_ = v1375
	var v1378 int32
	_ = v1378
	var v1379 int64
	_ = v1379
	var v1390 int32
	_ = v1390
	var v1391 int32
	_ = v1391
	var v1393 int64
	_ = v1393
	var v1397 int32
	_ = v1397
	var v1400 int32
	_ = v1400
	var v1408 int32
	_ = v1408
	var v1411 int32
	_ = v1411
	var v1413 int32
	_ = v1413
	var v1416 int32
	_ = v1416
	var v1419 int32
	_ = v1419
	var v1421 int32
	_ = v1421
	var v1424 int32
	_ = v1424
	var v1426 int32
	_ = v1426
	var v1427 int32
	_ = v1427
	var v1430 int32
	_ = v1430
	var v1431 int64
	_ = v1431
	var v1432 int64
	_ = v1432
	var v1435 int64
	_ = v1435
	var v1439 int64
	_ = v1439
	var v1441 int64
	_ = v1441
	var v1443 int64
	_ = v1443
	var v1445 int64
	_ = v1445
	var v1448 int64
	_ = v1448
	var v1466 int32
	_ = v1466
	var v1471 int32
	_ = v1471
	var v1474 int32
	_ = v1474
	var v1475 int32
	_ = v1475
	var v1487 int32
	_ = v1487
	var v1492 int32
	_ = v1492
	var v1493 int32
	_ = v1493
	var v1496 int32
	_ = v1496
	var v1503 int32
	_ = v1503
	var v1506 int32
	_ = v1506
	var v1509 int32
	_ = v1509
	var v1512 int32
	_ = v1512
	var v1513 int32
	_ = v1513
	var v1515 int32
	_ = v1515
	var v1519 int32
	_ = v1519
	var v1520 int32
	_ = v1520
	var v1523 int32
	_ = v1523
	var v1530 int32
	_ = v1530
	var v1533 int32
	_ = v1533
	var v1536 int32
	_ = v1536
	var v1539 int32
	_ = v1539
	var v1540 int32
	_ = v1540
	var v1542 int32
	_ = v1542
	var v1543 int32
	_ = v1543
	var v1544 int32
	_ = v1544
	var v1547 int32
	_ = v1547
	var v1550 int32
	_ = v1550
	var v1553 int32
	_ = v1553
	var v1554 int64
	_ = v1554
	var v1565 int32
	_ = v1565
	var v1566 int32
	_ = v1566
	var v1568 int64
	_ = v1568
	var v1572 int32
	_ = v1572
	var v1575 int32
	_ = v1575
	var v1583 int32
	_ = v1583
	var v1586 int32
	_ = v1586
	var v1588 int32
	_ = v1588
	var v1591 int32
	_ = v1591
	var v1594 int32
	_ = v1594
	var v1596 int32
	_ = v1596
	var v1599 int32
	_ = v1599
	var v1601 int32
	_ = v1601
	var v1602 int32
	_ = v1602
	var v1605 int32
	_ = v1605
	var v1606 int64
	_ = v1606
	var v1607 int64
	_ = v1607
	var v1610 int64
	_ = v1610
	var v1614 int64
	_ = v1614
	var v1616 int64
	_ = v1616
	var v1618 int64
	_ = v1618
	var v1620 int64
	_ = v1620
	var v1623 int64
	_ = v1623
	var v1638 int32
	_ = v1638
	var v1639 int32
	_ = v1639
	var v1654 int32
	_ = v1654
	var v1655 int32
	_ = v1655
	var v1656 int32
	_ = v1656
	var v1659 int32
	_ = v1659
	var v1662 int32
	_ = v1662
	var v1665 int32
	_ = v1665
	var v1666 int64
	_ = v1666
	var v1677 int32
	_ = v1677
	var v1678 int32
	_ = v1678
	var v1680 int64
	_ = v1680
	var v1684 int32
	_ = v1684
	var v1687 int32
	_ = v1687
	var v1695 int32
	_ = v1695
	var v1698 int32
	_ = v1698
	var v1700 int32
	_ = v1700
	var v1703 int32
	_ = v1703
	var v1706 int32
	_ = v1706
	var v1708 int32
	_ = v1708
	var v1711 int32
	_ = v1711
	var v1713 int32
	_ = v1713
	var v1714 int32
	_ = v1714
	var v1717 int32
	_ = v1717
	var v1718 int64
	_ = v1718
	var v1719 int64
	_ = v1719
	var v1722 int64
	_ = v1722
	var v1726 int64
	_ = v1726
	var v1728 int64
	_ = v1728
	var v1730 int64
	_ = v1730
	var v1732 int64
	_ = v1732
	var v1735 int64
	_ = v1735
	var v1750 int32
	_ = v1750
	var v1751 int32
	_ = v1751
	var v1755 int32
	_ = v1755
	var v1762 int32
	_ = v1762
	var v1763 int32
	_ = v1763
	var v1775 int64
	_ = v1775
	var v1799 int32
	_ = v1799
	var v1800 int32
	_ = v1800
	var v1801 int64
	_ = v1801
	var v1802 int64
	_ = v1802
	var v1803 int64
	_ = v1803
	var v1804 int32
	_ = v1804
	var v1807 int32
	_ = v1807
	var v1810 int32
	_ = v1810
	var v1813 int32
	_ = v1813
	var v1815 int64
	_ = v1815
	var v1826 int32
	_ = v1826
	var v1827 int32
	_ = v1827
	var v1829 int64
	_ = v1829
	var v1833 int32
	_ = v1833
	var v1836 int32
	_ = v1836
	var v1844 int32
	_ = v1844
	var v1847 int32
	_ = v1847
	var v1849 int32
	_ = v1849
	var v1852 int32
	_ = v1852
	var v1855 int32
	_ = v1855
	var v1857 int32
	_ = v1857
	var v1860 int32
	_ = v1860
	var v1862 int32
	_ = v1862
	var v1863 int32
	_ = v1863
	var v1866 int32
	_ = v1866
	var v1867 int64
	_ = v1867
	var v1868 int64
	_ = v1868
	var v1871 int64
	_ = v1871
	var v1875 int64
	_ = v1875
	var v1877 int64
	_ = v1877
	var v1879 int64
	_ = v1879
	var v1881 int64
	_ = v1881
	var v1884 int64
	_ = v1884
	var v1897 int64
	_ = v1897
	var v1920 int32
	_ = v1920
	var v1921 int32
	_ = v1921
	var v1934 int32
	_ = v1934
	var v1935 int32
	_ = v1935
	var v1936 int32
	_ = v1936
	var v1939 int32
	_ = v1939
	var v1942 int32
	_ = v1942
	var v1945 int32
	_ = v1945
	var v1946 int64
	_ = v1946
	var v1957 int32
	_ = v1957
	var v1958 int32
	_ = v1958
	var v1960 int64
	_ = v1960
	var v1964 int32
	_ = v1964
	var v1967 int32
	_ = v1967
	var v1975 int32
	_ = v1975
	var v1978 int32
	_ = v1978
	var v1980 int32
	_ = v1980
	var v1983 int32
	_ = v1983
	var v1986 int32
	_ = v1986
	var v1988 int32
	_ = v1988
	var v1991 int32
	_ = v1991
	var v1993 int32
	_ = v1993
	var v1994 int32
	_ = v1994
	var v1997 int32
	_ = v1997
	var v1998 int64
	_ = v1998
	var v1999 int64
	_ = v1999
	var v2002 int64
	_ = v2002
	var v2006 int64
	_ = v2006
	var v2008 int64
	_ = v2008
	var v2010 int64
	_ = v2010
	var v2012 int64
	_ = v2012
	var v2015 int64
	_ = v2015
	var v2033 int32
	_ = v2033
	var v2039 int32
	_ = v2039
	var v2045 int32
	_ = v2045
	var v2047 int32
	_ = v2047
	var v2051 int32
	_ = v2051
	var v2054 int32
	_ = v2054
	var v2055 int64
	_ = v2055
	var v2066 int32
	_ = v2066
	var v2067 int32
	_ = v2067
	var v2069 int64
	_ = v2069
	var v2073 int32
	_ = v2073
	var v2076 int32
	_ = v2076
	var v2084 int32
	_ = v2084
	var v2087 int32
	_ = v2087
	var v2089 int32
	_ = v2089
	var v2092 int32
	_ = v2092
	var v2095 int32
	_ = v2095
	var v2097 int32
	_ = v2097
	var v2100 int32
	_ = v2100
	var v2102 int32
	_ = v2102
	var v2103 int32
	_ = v2103
	var v2106 int32
	_ = v2106
	var v2107 int64
	_ = v2107
	var v2108 int64
	_ = v2108
	var v2111 int64
	_ = v2111
	var v2115 int64
	_ = v2115
	var v2117 int64
	_ = v2117
	var v2119 int64
	_ = v2119
	var v2121 int64
	_ = v2121
	v5 = int32(0)
	v10 = m.G0
	v12 = v10 - int32(768)
	m.G0 = v12
	v15 = *(*int32)(unsafe.Add(mBase, _consts[914]))
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v15+v16&int32(15)<<(uint(int32(2))%32)+l0*int32(28))))
	v27 = F_objectGetVal(m, l1)
	mBase = m.M
	v30 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v27+int32(-1)))))
	switch v30 & int32(7) {
	case 0:
		goto L6
	case 1:
		goto L5
	case 2:
		goto L4
	case 3:
		goto L3
	case 4:
		goto L2
	default:
		v47 = v5
		goto L1
	}
L1:
	;
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v25)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v25)+16)) = v48 + v47
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v25)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v25)+4)) = v51 + int32(1)
	v56 = base.B2i32(l3 == int64(-1))
	if l3 == int64(-1) {
		goto L9
	} else {
		goto L10
	}
L2:
	;
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v27+int32(-17))))
	v47 = v46
	goto L1
L3:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v27+int32(-9))))
	v47 = v43
	goto L1
L4:
	;
	v40 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v27+int32(-5)))))
	v47 = v40
	goto L1
L5:
	;
	v37 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v27+int32(-3)))))
	v47 = v37
	goto L1
L6:
	;
	v47 = int32(base.Ui32(v30) >> (uint(int32(3)) % 32))
	goto L1
L7:
	;
	v68 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	switch v68 & int32(15) {
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
	case 5:
		goto L14
	case 6:
		goto L18
	default:
		goto L17
	}
L8:
	;
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v25)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v25)+8)) = v64 + int32(1)
	goto L7
L9:
	;
	if l3 == int64(-1) {
		goto L7
	} else {
		goto L12
	}
L10:
	;
	v58 = *(*int64)(unsafe.Add(mBase, _consts[915]))
	if v58 <= l3 {
		goto L9
	} else {
		goto L11
	}
L11:
	;
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v25)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v25)+12)) = v60 + int32(1)
	goto L8
L12:
	;
	goto L8
L13:
	;
	m.G0 = v12 + int32(768)
	return
L14:
	;
	v2047 = *(*int32)(unsafe.Add(mBase, uint32(v25)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v25)+24)) = v2047 + int32(1)
	v2051 = *(*int32)(unsafe.Add(mBase, uint32(v25)+32))
	if v2051 != 0 {
		goto L552
	} else {
		goto L553
	}
L15:
	;
	F__serverAssert(m, int32(_a574), int32(_a2495), int32(297))
	mBase = m.M
	v2045 = m.ExcPending
	if v2045 != 0 {
		goto L24
	} else {
		goto L551
	}
L16:
	;
	F__serverAssert(m, int32(_a575), int32(_a2495), int32(299))
	mBase = m.M
	v2039 = m.ExcPending
	if v2039 != 0 {
		goto L24
	} else {
		goto L550
	}
L17:
	;
	F__serverPanic_1(m, int32(_a2495), int32(378), int32(_a573), int32(0))
	mBase = m.M
	v2033 = m.ExcPending
	if v2033 != 0 {
		goto L24
	} else {
		goto L549
	}
L18:
	;
	v1750 = F_objectGetVal(m, l2)
	mBase = m.M
	v1751 = int32(0)
	F_streamIteratorStart(m, v12+int32(32), v1750, v1751, v1751, v1751)
	mBase = m.M
	v1755 = m.ExcPending
	if v1755 != 0 {
		goto L24
	} else {
		goto L482
	}
L19:
	;
	F_hashTypeInitIterator(m, l2, v12+int32(32))
	mBase = m.M
	v1471 = m.ExcPending
	if v1471 != 0 {
		goto L24
	} else {
		goto L405
	}
L20:
	;
	switch int32(base.Ui32(v68)>>(uint(int32(4))%32))&int32(15) + int32(-7) {
	case 0:
		goto L209
	default:
		goto L208
	case 4:
		goto L210
	}
L21:
	;
	v482 = F_setTypeInitIterator(m, l2)
	mBase = m.M
	v483 = m.ExcPending
	if v483 != 0 {
		goto L24
	} else {
		goto L141
	}
L22:
	;
	v257 = F_listTypeInitIterator(m, l2, int32(0), int32(1))
	mBase = m.M
	v258 = m.ExcPending
	if v258 != 0 {
		goto L24
	} else {
		goto L77
	}
L23:
	;
	v71 = F_stringObjectLen(m, l2)
	mBase = m.M
	v72 = m.ExcPending
	if v72 != 0 {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	return
L25:
	;
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v25)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v25)+24)) = v73 + int32(1)
	v77 = *(*int32)(unsafe.Add(mBase, uint32(v25)+32))
	if v77 != 0 {
		goto L26
	} else {
		goto L27
	}
L26:
	;
	v80 = *(*int32)(unsafe.Add(mBase, uint32(v25)+40))
	v81 = int64(1)
	goto L29
L27:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25)+32)) = int32(1)
	goto L26
L28:
	;
	v163 = *(*int32)(unsafe.Add(mBase, uint32(v25)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v25)+20)) = v163 + v71
	v166 = *(*int32)(unsafe.Add(mBase, uint32(v25)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v25)+28)) = v166 + v71
	v169 = *(*int32)(unsafe.Add(mBase, uint32(v25)+36))
	if base.Ui32(v71) <= base.Ui32(v169) {
		goto L51
	} else {
		goto L52
	}
L29:
	;
	v92 = *(*int32)(unsafe.Add(mBase, uint32(v80)+24))
	v93 = *(*int32)(unsafe.Add(mBase, uint32(v80)+16))
	v95 = *(*int64)(unsafe.Add(mBase, uint32(v80)+32))
	v99 = v92 + v93 + base.I32_wrap_i64(base.I64_clz(v95|v81))
	v102 = *(*int32)(unsafe.Add(mBase, uint32(v80)+28))
	v110 = (int32(64)-v99)<<(uint(v92)%32) - v102 + base.I32_wrap_i64(int64(base.Ui64(v81)>>(uint(base.I64_extend_i32_u(v93-v99+int32(63)))%64)))
	if v110 < int32(0) {
		goto L31
	} else {
		goto L32
	}
L31:
	;
	goto L28
L32:
	;
	v113 = *(*int32)(unsafe.Add(mBase, uint32(v80)+80))
	if v113 <= v110 {
		goto L31
	} else {
		goto L33
	}
L33:
	;
	v115 = *(*int32)(unsafe.Add(mBase, uint32(v80)+64))
	if v115 == int32(0) {
		v128 = v110
		goto L34
	} else {
		goto L35
	}
L34:
	;
	v129 = *(*int32)(unsafe.Add(mBase, uint32(v80)+96))
	v132 = v129 + v128<<(uint(int32(3))%32)
	v133 = *(*int64)(unsafe.Add(mBase, uint32(v132)))
	v134 = int64(1)
	*(*int64)(unsafe.Add(mBase, uint32(v132))) = v133 + v134
	v137 = *(*int64)(unsafe.Add(mBase, uint32(v80)+88))
	*(*int64)(unsafe.Add(mBase, uint32(v80)+88)) = v137 + v134
	v141 = *(*int64)(unsafe.Add(mBase, uint32(v80)+56))
	if v141 < v81 {
		goto L42
	} else {
		goto L43
	}
L35:
	;
	v118 = int32(0)
	v121 = v110 - v115
	if v121 < v113 {
		goto L36
	} else {
		goto L37
	}
L36:
	;
	v123 = v118
	goto L38
L37:
	;
	v123 = v118 - v113
	goto L38
L38:
	;
	if v121 < int32(0) {
		goto L39
	} else {
		goto L40
	}
L39:
	;
	v126 = v113
	goto L41
L40:
	;
	v126 = v123
	goto L41
L41:
	;
	v128 = v126 + v121
	goto L34
L42:
	;
	v143 = v81
	goto L44
L43:
	;
	v143 = v141
	goto L44
L44:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v80)+56)) = v143
	v145 = *(*int64)(unsafe.Add(mBase, uint32(v80)+48))
	if v81 < v145 {
		goto L45
	} else {
		goto L46
	}
L45:
	;
	v147 = v81
	goto L47
L46:
	;
	v147 = v145
	goto L47
L47:
	;
	goto L49
L49:
	;
	goto L50
L50:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v80)+48)) = v147
	goto L31
L51:
	;
	v172 = *(*int32)(unsafe.Add(mBase, uint32(v25)+44))
	v173 = base.I64_extend_i32_u(v71)
	if int64(0) <= v173 {
		goto L54
	} else {
		goto L55
	}
L52:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25)+36)) = v71
	goto L51
L53:
	;
	goto L13
L54:
	;
	v184 = *(*int32)(unsafe.Add(mBase, uint32(v172)+24))
	v185 = *(*int32)(unsafe.Add(mBase, uint32(v172)+16))
	v187 = *(*int64)(unsafe.Add(mBase, uint32(v172)+32))
	v191 = v184 + v185 + base.I32_wrap_i64(base.I64_clz(v187|v173))
	v194 = *(*int32)(unsafe.Add(mBase, uint32(v172)+28))
	v202 = (int32(64)-v191)<<(uint(v184)%32) - v194 + base.I32_wrap_i64(int64(base.Ui64(v173)>>(uint(base.I64_extend_i32_u(v185-v191+int32(63)))%64)))
	if v202 < int32(0) {
		goto L56
	} else {
		goto L57
	}
L55:
	;
	goto L53
L56:
	;
	goto L53
L57:
	;
	v205 = *(*int32)(unsafe.Add(mBase, uint32(v172)+80))
	if v205 <= v202 {
		goto L56
	} else {
		goto L58
	}
L58:
	;
	v207 = *(*int32)(unsafe.Add(mBase, uint32(v172)+64))
	if v207 == int32(0) {
		v220 = v202
		goto L59
	} else {
		goto L60
	}
L59:
	;
	v221 = *(*int32)(unsafe.Add(mBase, uint32(v172)+96))
	v224 = v221 + v220<<(uint(int32(3))%32)
	v225 = *(*int64)(unsafe.Add(mBase, uint32(v224)))
	v226 = int64(1)
	*(*int64)(unsafe.Add(mBase, uint32(v224))) = v225 + v226
	v229 = *(*int64)(unsafe.Add(mBase, uint32(v172)+88))
	*(*int64)(unsafe.Add(mBase, uint32(v172)+88)) = v229 + v226
	v233 = *(*int64)(unsafe.Add(mBase, uint32(v172)+56))
	if v233 < v173 {
		goto L67
	} else {
		goto L68
	}
L60:
	;
	v210 = int32(0)
	v213 = v202 - v207
	if v213 < v205 {
		goto L61
	} else {
		goto L62
	}
L61:
	;
	v215 = v210
	goto L63
L62:
	;
	v215 = v210 - v205
	goto L63
L63:
	;
	if v213 < int32(0) {
		goto L64
	} else {
		goto L65
	}
L64:
	;
	v218 = v205
	goto L66
L65:
	;
	v218 = v215
	goto L66
L66:
	;
	v220 = v218 + v213
	goto L59
L67:
	;
	v235 = v173
	goto L69
L68:
	;
	v235 = v233
	goto L69
L69:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v172)+56)) = v235
	v237 = *(*int64)(unsafe.Add(mBase, uint32(v172)+48))
	if v173 < v237 {
		goto L70
	} else {
		goto L71
	}
L70:
	;
	v239 = v173
	goto L72
L71:
	;
	v239 = v237
	goto L72
L72:
	;
	if v173 == int64(0) {
		goto L73
	} else {
		goto L74
	}
L73:
	;
	v242 = v237
	goto L75
L74:
	;
	v242 = v239
	goto L75
L75:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v172)+48)) = v242
	goto L56
L76:
	;
	F_listTypeReleaseIterator(m, v257)
	mBase = m.M
	v388 = m.ExcPending
	if v388 != 0 {
		goto L24
	} else {
		goto L112
	}
L77:
	;
	v261 = F_listTypeNext(m, v257, v12+int32(32))
	mBase = m.M
	v262 = m.ExcPending
	if v262 != 0 {
		goto L24
	} else {
		goto L78
	}
L78:
	;
	if v261 == int32(0) {
		goto L76
	} else {
		goto L79
	}
L79:
	;
	goto L80
L80:
	;
	v276 = F_listTypeGet(m, v12+int32(32))
	mBase = m.M
	v277 = m.ExcPending
	if v277 != 0 {
		goto L24
	} else {
		goto L82
	}
L81:
	;
	goto L76
L82:
	;
	v278 = F_stringObjectLen(m, v276)
	mBase = m.M
	v279 = m.ExcPending
	if v279 != 0 {
		goto L24
	} else {
		goto L83
	}
L83:
	;
	v280 = *(*int32)(unsafe.Add(mBase, uint32(v25)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v25)+20)) = v278 + v280
	v283 = *(*int32)(unsafe.Add(mBase, uint32(v25)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v25)+28)) = v278 + v283
	v286 = *(*int32)(unsafe.Add(mBase, uint32(v25)+36))
	if base.Ui32(v278) <= base.Ui32(v286) {
		goto L84
	} else {
		goto L85
	}
L84:
	;
	v289 = *(*int32)(unsafe.Add(mBase, uint32(v25)+44))
	v290 = base.I64_extend_i32_u(v278)
	if int64(0) <= v290 {
		goto L87
	} else {
		goto L88
	}
L85:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25)+36)) = v278
	goto L84
L86:
	;
	F_decrRefCount(m, v276)
	mBase = m.M
	v373 = m.ExcPending
	if v373 != 0 {
		goto L24
	} else {
		goto L109
	}
L87:
	;
	v301 = *(*int32)(unsafe.Add(mBase, uint32(v289)+24))
	v302 = *(*int32)(unsafe.Add(mBase, uint32(v289)+16))
	v304 = *(*int64)(unsafe.Add(mBase, uint32(v289)+32))
	v308 = v301 + v302 + base.I32_wrap_i64(base.I64_clz(v304|v290))
	v311 = *(*int32)(unsafe.Add(mBase, uint32(v289)+28))
	v319 = (int32(64)-v308)<<(uint(v301)%32) - v311 + base.I32_wrap_i64(int64(base.Ui64(v290)>>(uint(base.I64_extend_i32_u(v302-v308+int32(63)))%64)))
	if v319 < int32(0) {
		goto L89
	} else {
		goto L90
	}
L88:
	;
	goto L86
L89:
	;
	goto L86
L90:
	;
	v322 = *(*int32)(unsafe.Add(mBase, uint32(v289)+80))
	if v322 <= v319 {
		goto L89
	} else {
		goto L91
	}
L91:
	;
	v324 = *(*int32)(unsafe.Add(mBase, uint32(v289)+64))
	if v324 == int32(0) {
		v337 = v319
		goto L92
	} else {
		goto L93
	}
L92:
	;
	v338 = *(*int32)(unsafe.Add(mBase, uint32(v289)+96))
	v341 = v338 + v337<<(uint(int32(3))%32)
	v342 = *(*int64)(unsafe.Add(mBase, uint32(v341)))
	v343 = int64(1)
	*(*int64)(unsafe.Add(mBase, uint32(v341))) = v342 + v343
	v346 = *(*int64)(unsafe.Add(mBase, uint32(v289)+88))
	*(*int64)(unsafe.Add(mBase, uint32(v289)+88)) = v346 + v343
	v350 = *(*int64)(unsafe.Add(mBase, uint32(v289)+56))
	if v350 < v290 {
		goto L100
	} else {
		goto L101
	}
L93:
	;
	v327 = int32(0)
	v330 = v319 - v324
	if v330 < v322 {
		goto L94
	} else {
		goto L95
	}
L94:
	;
	v332 = v327
	goto L96
L95:
	;
	v332 = v327 - v322
	goto L96
L96:
	;
	if v330 < int32(0) {
		goto L97
	} else {
		goto L98
	}
L97:
	;
	v335 = v322
	goto L99
L98:
	;
	v335 = v332
	goto L99
L99:
	;
	v337 = v335 + v330
	goto L92
L100:
	;
	v352 = v290
	goto L102
L101:
	;
	v352 = v350
	goto L102
L102:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v289)+56)) = v352
	v354 = *(*int64)(unsafe.Add(mBase, uint32(v289)+48))
	if v290 < v354 {
		goto L103
	} else {
		goto L104
	}
L103:
	;
	v356 = v290
	goto L105
L104:
	;
	v356 = v354
	goto L105
L105:
	;
	if v290 == int64(0) {
		goto L106
	} else {
		goto L107
	}
L106:
	;
	v359 = v354
	goto L108
L107:
	;
	v359 = v356
	goto L108
L108:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v289)+48)) = v359
	goto L89
L109:
	;
	v376 = F_listTypeNext(m, v257, v12+int32(32))
	mBase = m.M
	v377 = m.ExcPending
	if v377 != 0 {
		goto L24
	} else {
		goto L110
	}
L110:
	;
	if v376 != 0 {
		goto L80
	} else {
		goto L111
	}
L111:
	;
	goto L81
L112:
	;
	v389 = F_listTypeLength(m, l2)
	mBase = m.M
	v390 = m.ExcPending
	if v390 != 0 {
		goto L24
	} else {
		goto L113
	}
L113:
	;
	if v25 == int32(0) {
		goto L13
	} else {
		goto L114
	}
L114:
	;
	v393 = *(*int32)(unsafe.Add(mBase, uint32(v25)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v25)+24)) = v393 + v389
	v396 = *(*int32)(unsafe.Add(mBase, uint32(v25)+32))
	if base.Ui32(v389) <= base.Ui32(v396) {
		goto L115
	} else {
		goto L116
	}
L115:
	;
	v399 = *(*int32)(unsafe.Add(mBase, uint32(v25)+40))
	v400 = base.I64_extend_i32_u(v389)
	if int64(0) <= v400 {
		goto L118
	} else {
		goto L119
	}
L116:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25)+32)) = v389
	goto L115
L117:
	;
	goto L13
L118:
	;
	v411 = *(*int32)(unsafe.Add(mBase, uint32(v399)+24))
	v412 = *(*int32)(unsafe.Add(mBase, uint32(v399)+16))
	v414 = *(*int64)(unsafe.Add(mBase, uint32(v399)+32))
	v418 = v411 + v412 + base.I32_wrap_i64(base.I64_clz(v414|v400))
	v421 = *(*int32)(unsafe.Add(mBase, uint32(v399)+28))
	v429 = (int32(64)-v418)<<(uint(v411)%32) - v421 + base.I32_wrap_i64(int64(base.Ui64(v400)>>(uint(base.I64_extend_i32_u(v412-v418+int32(63)))%64)))
	if v429 < int32(0) {
		goto L120
	} else {
		goto L121
	}
L119:
	;
	goto L117
L120:
	;
	goto L117
L121:
	;
	v432 = *(*int32)(unsafe.Add(mBase, uint32(v399)+80))
	if v432 <= v429 {
		goto L120
	} else {
		goto L122
	}
L122:
	;
	v434 = *(*int32)(unsafe.Add(mBase, uint32(v399)+64))
	if v434 == int32(0) {
		v447 = v429
		goto L123
	} else {
		goto L124
	}
L123:
	;
	v448 = *(*int32)(unsafe.Add(mBase, uint32(v399)+96))
	v451 = v448 + v447<<(uint(int32(3))%32)
	v452 = *(*int64)(unsafe.Add(mBase, uint32(v451)))
	v453 = int64(1)
	*(*int64)(unsafe.Add(mBase, uint32(v451))) = v452 + v453
	v456 = *(*int64)(unsafe.Add(mBase, uint32(v399)+88))
	*(*int64)(unsafe.Add(mBase, uint32(v399)+88)) = v456 + v453
	v460 = *(*int64)(unsafe.Add(mBase, uint32(v399)+56))
	if v460 < v400 {
		goto L131
	} else {
		goto L132
	}
L124:
	;
	v437 = int32(0)
	v440 = v429 - v434
	if v440 < v432 {
		goto L125
	} else {
		goto L126
	}
L125:
	;
	v442 = v437
	goto L127
L126:
	;
	v442 = v437 - v432
	goto L127
L127:
	;
	if v440 < int32(0) {
		goto L128
	} else {
		goto L129
	}
L128:
	;
	v445 = v432
	goto L130
L129:
	;
	v445 = v442
	goto L130
L130:
	;
	v447 = v445 + v440
	goto L123
L131:
	;
	v462 = v400
	goto L133
L132:
	;
	v462 = v460
	goto L133
L133:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v399)+56)) = v462
	v464 = *(*int64)(unsafe.Add(mBase, uint32(v399)+48))
	if v400 < v464 {
		goto L134
	} else {
		goto L135
	}
L134:
	;
	v466 = v400
	goto L136
L135:
	;
	v466 = v464
	goto L136
L136:
	;
	if v400 == int64(0) {
		goto L137
	} else {
		goto L138
	}
L137:
	;
	v469 = v464
	goto L139
L138:
	;
	v469 = v466
	goto L139
L139:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v399)+48)) = v469
	goto L120
L140:
	;
	F_setTypeReleaseIterator(m, v482)
	mBase = m.M
	v624 = m.ExcPending
	if v624 != 0 {
		goto L24
	} else {
		goto L180
	}
L141:
	;
	v484 = F_setTypeNextObject(m, v482)
	mBase = m.M
	v485 = m.ExcPending
	if v485 != 0 {
		goto L24
	} else {
		goto L142
	}
L142:
	;
	if v484 == int32(0) {
		goto L140
	} else {
		goto L143
	}
L143:
	;
	v489 = v484
	goto L144
L144:
	;
	v500 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v489+int32(-1)))))
	switch v500 & int32(7) {
	case 0:
		goto L151
	case 1:
		goto L150
	case 2:
		goto L149
	case 3:
		goto L148
	case 4:
		goto L147
	default:
		v517 = int32(0)
		goto L146
	}
L145:
	;
	goto L140
L146:
	;
	v518 = *(*int32)(unsafe.Add(mBase, uint32(v25)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v25)+20)) = v518 + v517
	v521 = *(*int32)(unsafe.Add(mBase, uint32(v25)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v25)+28)) = v521 + v517
	v524 = *(*int32)(unsafe.Add(mBase, uint32(v25)+36))
	if base.Ui32(v517) <= base.Ui32(v524) {
		goto L152
	} else {
		goto L153
	}
L147:
	;
	v516 = *(*int32)(unsafe.Add(mBase, uint32(v489+int32(-17))))
	v517 = v516
	goto L146
L148:
	;
	v513 = *(*int32)(unsafe.Add(mBase, uint32(v489+int32(-9))))
	v517 = v513
	goto L146
L149:
	;
	v510 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v489+int32(-5)))))
	v517 = v510
	goto L146
L150:
	;
	v507 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v489+int32(-3)))))
	v517 = v507
	goto L146
L151:
	;
	v517 = int32(base.Ui32(v500) >> (uint(int32(3)) % 32))
	goto L146
L152:
	;
	v527 = *(*int32)(unsafe.Add(mBase, uint32(v25)+44))
	v528 = base.I64_extend_i32_u(v517)
	if int64(0) <= v528 {
		goto L155
	} else {
		goto L156
	}
L153:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25)+36)) = v517
	goto L152
L154:
	;
	F_sdsfree(m, v489)
	mBase = m.M
	v611 = m.ExcPending
	if v611 != 0 {
		goto L24
	} else {
		goto L177
	}
L155:
	;
	v539 = *(*int32)(unsafe.Add(mBase, uint32(v527)+24))
	v540 = *(*int32)(unsafe.Add(mBase, uint32(v527)+16))
	v542 = *(*int64)(unsafe.Add(mBase, uint32(v527)+32))
	v546 = v539 + v540 + base.I32_wrap_i64(base.I64_clz(v542|v528))
	v549 = *(*int32)(unsafe.Add(mBase, uint32(v527)+28))
	v557 = (int32(64)-v546)<<(uint(v539)%32) - v549 + base.I32_wrap_i64(int64(base.Ui64(v528)>>(uint(base.I64_extend_i32_u(v540-v546+int32(63)))%64)))
	if v557 < int32(0) {
		goto L157
	} else {
		goto L158
	}
L156:
	;
	goto L154
L157:
	;
	goto L154
L158:
	;
	v560 = *(*int32)(unsafe.Add(mBase, uint32(v527)+80))
	if v560 <= v557 {
		goto L157
	} else {
		goto L159
	}
L159:
	;
	v562 = *(*int32)(unsafe.Add(mBase, uint32(v527)+64))
	if v562 == int32(0) {
		v575 = v557
		goto L160
	} else {
		goto L161
	}
L160:
	;
	v576 = *(*int32)(unsafe.Add(mBase, uint32(v527)+96))
	v579 = v576 + v575<<(uint(int32(3))%32)
	v580 = *(*int64)(unsafe.Add(mBase, uint32(v579)))
	v581 = int64(1)
	*(*int64)(unsafe.Add(mBase, uint32(v579))) = v580 + v581
	v584 = *(*int64)(unsafe.Add(mBase, uint32(v527)+88))
	*(*int64)(unsafe.Add(mBase, uint32(v527)+88)) = v584 + v581
	v588 = *(*int64)(unsafe.Add(mBase, uint32(v527)+56))
	if v588 < v528 {
		goto L168
	} else {
		goto L169
	}
L161:
	;
	v565 = int32(0)
	v568 = v557 - v562
	if v568 < v560 {
		goto L162
	} else {
		goto L163
	}
L162:
	;
	v570 = v565
	goto L164
L163:
	;
	v570 = v565 - v560
	goto L164
L164:
	;
	if v568 < int32(0) {
		goto L165
	} else {
		goto L166
	}
L165:
	;
	v573 = v560
	goto L167
L166:
	;
	v573 = v570
	goto L167
L167:
	;
	v575 = v573 + v568
	goto L160
L168:
	;
	v590 = v528
	goto L170
L169:
	;
	v590 = v588
	goto L170
L170:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v527)+56)) = v590
	v592 = *(*int64)(unsafe.Add(mBase, uint32(v527)+48))
	if v528 < v592 {
		goto L171
	} else {
		goto L172
	}
L171:
	;
	v594 = v528
	goto L173
L172:
	;
	v594 = v592
	goto L173
L173:
	;
	if v528 == int64(0) {
		goto L174
	} else {
		goto L175
	}
L174:
	;
	v597 = v592
	goto L176
L175:
	;
	v597 = v594
	goto L176
L176:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v527)+48)) = v597
	goto L157
L177:
	;
	v612 = F_setTypeNextObject(m, v482)
	mBase = m.M
	v613 = m.ExcPending
	if v613 != 0 {
		goto L24
	} else {
		goto L178
	}
L178:
	;
	if v612 != 0 {
		v489 = v612
		goto L144
	} else {
		goto L179
	}
L179:
	;
	goto L145
L180:
	;
	v625 = F_setTypeSize(m, l2)
	mBase = m.M
	v626 = m.ExcPending
	if v626 != 0 {
		goto L24
	} else {
		goto L181
	}
L181:
	;
	if v25 == int32(0) {
		goto L13
	} else {
		goto L182
	}
L182:
	;
	v629 = *(*int32)(unsafe.Add(mBase, uint32(v25)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v25)+24)) = v629 + v625
	v632 = *(*int32)(unsafe.Add(mBase, uint32(v25)+32))
	if base.Ui32(v625) <= base.Ui32(v632) {
		goto L183
	} else {
		goto L184
	}
L183:
	;
	v635 = *(*int32)(unsafe.Add(mBase, uint32(v25)+40))
	v636 = base.I64_extend_i32_u(v625)
	if int64(0) <= v636 {
		goto L186
	} else {
		goto L187
	}
L184:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25)+32)) = v625
	goto L183
L185:
	;
	goto L13
L186:
	;
	v647 = *(*int32)(unsafe.Add(mBase, uint32(v635)+24))
	v648 = *(*int32)(unsafe.Add(mBase, uint32(v635)+16))
	v650 = *(*int64)(unsafe.Add(mBase, uint32(v635)+32))
	v654 = v647 + v648 + base.I32_wrap_i64(base.I64_clz(v650|v636))
	v657 = *(*int32)(unsafe.Add(mBase, uint32(v635)+28))
	v665 = (int32(64)-v654)<<(uint(v647)%32) - v657 + base.I32_wrap_i64(int64(base.Ui64(v636)>>(uint(base.I64_extend_i32_u(v648-v654+int32(63)))%64)))
	if v665 < int32(0) {
		goto L188
	} else {
		goto L189
	}
L187:
	;
	goto L185
L188:
	;
	goto L185
L189:
	;
	v668 = *(*int32)(unsafe.Add(mBase, uint32(v635)+80))
	if v668 <= v665 {
		goto L188
	} else {
		goto L190
	}
L190:
	;
	v670 = *(*int32)(unsafe.Add(mBase, uint32(v635)+64))
	if v670 == int32(0) {
		v683 = v665
		goto L191
	} else {
		goto L192
	}
L191:
	;
	v684 = *(*int32)(unsafe.Add(mBase, uint32(v635)+96))
	v687 = v684 + v683<<(uint(int32(3))%32)
	v688 = *(*int64)(unsafe.Add(mBase, uint32(v687)))
	v689 = int64(1)
	*(*int64)(unsafe.Add(mBase, uint32(v687))) = v688 + v689
	v692 = *(*int64)(unsafe.Add(mBase, uint32(v635)+88))
	*(*int64)(unsafe.Add(mBase, uint32(v635)+88)) = v692 + v689
	v696 = *(*int64)(unsafe.Add(mBase, uint32(v635)+56))
	if v696 < v636 {
		goto L199
	} else {
		goto L200
	}
L192:
	;
	v673 = int32(0)
	v676 = v665 - v670
	if v676 < v668 {
		goto L193
	} else {
		goto L194
	}
L193:
	;
	v678 = v673
	goto L195
L194:
	;
	v678 = v673 - v668
	goto L195
L195:
	;
	if v676 < int32(0) {
		goto L196
	} else {
		goto L197
	}
L196:
	;
	v681 = v668
	goto L198
L197:
	;
	v681 = v678
	goto L198
L198:
	;
	v683 = v681 + v676
	goto L191
L199:
	;
	v698 = v636
	goto L201
L200:
	;
	v698 = v696
	goto L201
L201:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v635)+56)) = v698
	v700 = *(*int64)(unsafe.Add(mBase, uint32(v635)+48))
	if v636 < v700 {
		goto L202
	} else {
		goto L203
	}
L202:
	;
	v702 = v636
	goto L204
L203:
	;
	v702 = v700
	goto L204
L204:
	;
	if v636 == int64(0) {
		goto L205
	} else {
		goto L206
	}
L205:
	;
	v705 = v700
	goto L207
L206:
	;
	v705 = v702
	goto L207
L207:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v635)+48)) = v705
	goto L188
L208:
	;
	F__serverPanic_1(m, int32(_a2495), int32(339), int32(_a576), int32(0))
	mBase = m.M
	v1466 = m.ExcPending
	if v1466 != 0 {
		goto L24
	} else {
		goto L404
	}
L209:
	;
	v1118 = v12 + int32(720)
	v1119 = F_objectGetVal(m, l2)
	mBase = m.M
	v1120 = *(*int32)(unsafe.Add(mBase, uint32(v1119)))
	v1121 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v1118)+14)) = uint8(v1121)
	*(*int32)(unsafe.Add(mBase, uint32(v1118))) = v1120
	*(*int32)(unsafe.Add(mBase, uint32(v1118)+24)) = v1121
	*(*uint8)(unsafe.Add(mBase, uint32(v1118)+15)) = uint8(v1121)
	*(*int32)(unsafe.Add(mBase, uint32(v1118)+8)) = int32(-1)
	if v1120 == v1121 {
		goto L318
	} else {
		goto L319
	}
L210:
	;
	v724 = F_objectGetVal(m, l2)
	mBase = m.M
	v726 = F_lpSeek(m, v724, int32(0))
	mBase = m.M
	v727 = m.ExcPending
	if v727 != 0 {
		goto L24
	} else {
		goto L211
	}
L211:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+24)) = v726
	if v726 == int32(0) {
		goto L15
	} else {
		goto L212
	}
L212:
	;
	v731 = F_lpNext(m, v724, v726)
	mBase = m.M
	v732 = m.ExcPending
	if v732 != 0 {
		goto L24
	} else {
		goto L213
	}
L213:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+8)) = v731
	if v731 == int32(0) {
		goto L16
	} else {
		goto L214
	}
L214:
	;
	v741 = v726
	goto L215
L215:
	;
	v747 = F_lpGetValue(m, v741, v12, v12+int32(720))
	mBase = m.M
	v748 = m.ExcPending
	if v748 != 0 {
		goto L24
	} else {
		goto L217
	}
L216:
	;
	v1023 = F_objectGetVal(m, l2)
	mBase = m.M
	v1024 = F_lpLength(m, v1023)
	mBase = m.M
	v1025 = m.ExcPending
	if v1025 != 0 {
		goto L24
	} else {
		goto L290
	}
L217:
	;
	v749 = *(*int32)(unsafe.Add(mBase, uint32(v12)+8))
	v750 = F_zzlGetScore(m, v749)
	mBase = m.M
	v751 = m.ExcPending
	if v751 != 0 {
		goto L24
	} else {
		goto L218
	}
L218:
	;
	if v747 == int32(0) {
		goto L220
	} else {
		goto L221
	}
L219:
	;
	v859 = v12 + int32(32)
	v862 = F_fpconv_dtoa(m, v750, v859)
	mBase = m.M
	v864 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v859+v862))) = uint8(v864)
	v866 = *(*int32)(unsafe.Add(mBase, uint32(v25)+20))
	if v859&int32(3) == v864 {
		v890 = v859
		goto L249
	} else {
		goto L250
	}
L220:
	;
	v756 = v12 + int32(32)
	v758 = *(*int64)(unsafe.Add(mBase, uint32(v12)+720))
	if v758 <= int64(-1) {
		goto L226
	} else {
		goto L227
	}
L221:
	;
	v754 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
	v857 = v754
	goto L219
L222:
	;
	v801 = v12 + int32(32)
	if v801&int32(3) == int32(0) {
		v823 = v801
		goto L233
	} else {
		goto L234
	}
L223:
	;
	goto L222
L225:
	;
	v780 = F_ull2string(m, v776, v777, v778)
	mBase = m.M
	if v780 == int32(0) {
		goto L223
	} else {
		goto L229
	}
L226:
	;
	goto L228
L227:
	;
	v776 = v756
	v777 = int32(128)
	v778 = v758
	goto L225
L228:
	;
	v767 = int32(45)
	*(*uint8)(unsafe.Add(mBase, uint32(v756))) = uint8(v767)
	v776 = v12 + int32(33)
	v777 = int32(127)
	v778 = int64(0) - v758
	goto L225
L229:
	;
	goto L222
L231:
	;
	v857 = v856
	goto L219
L232:
	;
	v856 = v848 - v801
	goto L231
L233:
	;
	v827 = v823
	goto L241
L234:
	;
	v809 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v801))))
	if v809 != 0 {
		goto L235
	} else {
		goto L236
	}
L235:
	;
	v812 = v801
	goto L237
L236:
	;
	v856 = v801 - v801
	goto L231
L237:
	;
	v816 = v812 + int32(1)
	if v816&int32(3) == int32(0) {
		v823 = v816
		goto L233
	} else {
		goto L239
	}
L239:
	;
	v821 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v816))))
	if v821 != 0 {
		v812 = v816
		goto L237
	} else {
		goto L240
	}
L240:
	;
	v848 = v816
	goto L232
L241:
	;
	v833 = *(*int32)(unsafe.Add(mBase, uint32(v827)))
	v836 = int32(-2139062144)
	if (int32(16843008)-v833|v833)&v836 == v836 {
		v827 = v827 + int32(4)
		goto L241
	} else {
		goto L243
	}
L242:
	;
	v842 = v827
	goto L244
L243:
	;
	goto L242
L244:
	;
	v846 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v842))))
	if v846 != 0 {
		v842 = v842 + int32(1)
		goto L244
	} else {
		goto L246
	}
L245:
	;
	v848 = v842
	goto L232
L246:
	;
	goto L245
L247:
	;
	v924 = v923 + v857
	*(*int32)(unsafe.Add(mBase, uint32(v25)+20)) = v866 + v924
	v927 = *(*int32)(unsafe.Add(mBase, uint32(v25)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v25)+28)) = v927 + v924
	v930 = *(*int32)(unsafe.Add(mBase, uint32(v25)+36))
	if base.Ui32(v924) <= base.Ui32(v930) {
		goto L263
	} else {
		goto L264
	}
L248:
	;
	v923 = v915 - v859
	goto L247
L249:
	;
	v894 = v890
	goto L257
L250:
	;
	v876 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v859))))
	if v876 != 0 {
		goto L251
	} else {
		goto L252
	}
L251:
	;
	v879 = v859
	goto L253
L252:
	;
	v923 = v859 - v859
	goto L247
L253:
	;
	v883 = v879 + int32(1)
	if v883&int32(3) == int32(0) {
		v890 = v883
		goto L249
	} else {
		goto L255
	}
L255:
	;
	v888 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v883))))
	if v888 != 0 {
		v879 = v883
		goto L253
	} else {
		goto L256
	}
L256:
	;
	v915 = v883
	goto L248
L257:
	;
	v900 = *(*int32)(unsafe.Add(mBase, uint32(v894)))
	v903 = int32(-2139062144)
	if (int32(16843008)-v900|v900)&v903 == v903 {
		v894 = v894 + int32(4)
		goto L257
	} else {
		goto L259
	}
L258:
	;
	v909 = v894
	goto L260
L259:
	;
	goto L258
L260:
	;
	v913 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v909))))
	if v913 != 0 {
		v909 = v909 + int32(1)
		goto L260
	} else {
		goto L262
	}
L261:
	;
	v915 = v909
	goto L248
L262:
	;
	goto L261
L263:
	;
	v933 = *(*int32)(unsafe.Add(mBase, uint32(v25)+44))
	v934 = base.I64_extend_i32_u(v924)
	if int64(0) <= v934 {
		goto L266
	} else {
		goto L267
	}
L264:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25)+36)) = v924
	goto L263
L265:
	;
	F_zzlNext(m, v724, v12+int32(24), v12+int32(8))
	mBase = m.M
	v1021 = m.ExcPending
	if v1021 != 0 {
		goto L24
	} else {
		goto L288
	}
L266:
	;
	v945 = *(*int32)(unsafe.Add(mBase, uint32(v933)+24))
	v946 = *(*int32)(unsafe.Add(mBase, uint32(v933)+16))
	v948 = *(*int64)(unsafe.Add(mBase, uint32(v933)+32))
	v952 = v945 + v946 + base.I32_wrap_i64(base.I64_clz(v948|v934))
	v955 = *(*int32)(unsafe.Add(mBase, uint32(v933)+28))
	v963 = (int32(64)-v952)<<(uint(v945)%32) - v955 + base.I32_wrap_i64(int64(base.Ui64(v934)>>(uint(base.I64_extend_i32_u(v946-v952+int32(63)))%64)))
	if v963 < int32(0) {
		goto L268
	} else {
		goto L269
	}
L267:
	;
	goto L265
L268:
	;
	goto L265
L269:
	;
	v966 = *(*int32)(unsafe.Add(mBase, uint32(v933)+80))
	if v966 <= v963 {
		goto L268
	} else {
		goto L270
	}
L270:
	;
	v968 = *(*int32)(unsafe.Add(mBase, uint32(v933)+64))
	if v968 == int32(0) {
		v981 = v963
		goto L271
	} else {
		goto L272
	}
L271:
	;
	v982 = *(*int32)(unsafe.Add(mBase, uint32(v933)+96))
	v985 = v982 + v981<<(uint(int32(3))%32)
	v986 = *(*int64)(unsafe.Add(mBase, uint32(v985)))
	v987 = int64(1)
	*(*int64)(unsafe.Add(mBase, uint32(v985))) = v986 + v987
	v990 = *(*int64)(unsafe.Add(mBase, uint32(v933)+88))
	*(*int64)(unsafe.Add(mBase, uint32(v933)+88)) = v990 + v987
	v994 = *(*int64)(unsafe.Add(mBase, uint32(v933)+56))
	if v994 < v934 {
		goto L279
	} else {
		goto L280
	}
L272:
	;
	v971 = int32(0)
	v974 = v963 - v968
	if v974 < v966 {
		goto L273
	} else {
		goto L274
	}
L273:
	;
	v976 = v971
	goto L275
L274:
	;
	v976 = v971 - v966
	goto L275
L275:
	;
	if v974 < int32(0) {
		goto L276
	} else {
		goto L277
	}
L276:
	;
	v979 = v966
	goto L278
L277:
	;
	v979 = v976
	goto L278
L278:
	;
	v981 = v979 + v974
	goto L271
L279:
	;
	v996 = v934
	goto L281
L280:
	;
	v996 = v994
	goto L281
L281:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v933)+56)) = v996
	v998 = *(*int64)(unsafe.Add(mBase, uint32(v933)+48))
	if v934 < v998 {
		goto L282
	} else {
		goto L283
	}
L282:
	;
	v1000 = v934
	goto L284
L283:
	;
	v1000 = v998
	goto L284
L284:
	;
	if v934 == int64(0) {
		goto L285
	} else {
		goto L286
	}
L285:
	;
	v1003 = v998
	goto L287
L286:
	;
	v1003 = v1000
	goto L287
L287:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v933)+48)) = v1003
	goto L268
L288:
	;
	v1022 = *(*int32)(unsafe.Add(mBase, uint32(v12)+24))
	if v1022 != 0 {
		v741 = v1022
		goto L215
	} else {
		goto L289
	}
L289:
	;
	goto L216
L290:
	;
	if v25 == int32(0) {
		goto L13
	} else {
		goto L291
	}
L291:
	;
	v1028 = *(*int32)(unsafe.Add(mBase, uint32(v25)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v25)+24)) = v1028 + v1024
	v1031 = *(*int32)(unsafe.Add(mBase, uint32(v25)+32))
	if base.Ui32(v1024) <= base.Ui32(v1031) {
		goto L292
	} else {
		goto L293
	}
L292:
	;
	v1034 = *(*int32)(unsafe.Add(mBase, uint32(v25)+40))
	v1035 = base.I64_extend_i32_u(v1024)
	if int64(0) <= v1035 {
		goto L295
	} else {
		goto L296
	}
L293:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25)+32)) = v1024
	goto L292
L294:
	;
	goto L13
L295:
	;
	v1046 = *(*int32)(unsafe.Add(mBase, uint32(v1034)+24))
	v1047 = *(*int32)(unsafe.Add(mBase, uint32(v1034)+16))
	v1049 = *(*int64)(unsafe.Add(mBase, uint32(v1034)+32))
	v1053 = v1046 + v1047 + base.I32_wrap_i64(base.I64_clz(v1049|v1035))
	v1056 = *(*int32)(unsafe.Add(mBase, uint32(v1034)+28))
	v1064 = (int32(64)-v1053)<<(uint(v1046)%32) - v1056 + base.I32_wrap_i64(int64(base.Ui64(v1035)>>(uint(base.I64_extend_i32_u(v1047-v1053+int32(63)))%64)))
	if v1064 < int32(0) {
		goto L297
	} else {
		goto L298
	}
L296:
	;
	goto L294
L297:
	;
	goto L294
L298:
	;
	v1067 = *(*int32)(unsafe.Add(mBase, uint32(v1034)+80))
	if v1067 <= v1064 {
		goto L297
	} else {
		goto L299
	}
L299:
	;
	v1069 = *(*int32)(unsafe.Add(mBase, uint32(v1034)+64))
	if v1069 == int32(0) {
		v1082 = v1064
		goto L300
	} else {
		goto L301
	}
L300:
	;
	v1083 = *(*int32)(unsafe.Add(mBase, uint32(v1034)+96))
	v1086 = v1083 + v1082<<(uint(int32(3))%32)
	v1087 = *(*int64)(unsafe.Add(mBase, uint32(v1086)))
	v1088 = int64(1)
	*(*int64)(unsafe.Add(mBase, uint32(v1086))) = v1087 + v1088
	v1091 = *(*int64)(unsafe.Add(mBase, uint32(v1034)+88))
	*(*int64)(unsafe.Add(mBase, uint32(v1034)+88)) = v1091 + v1088
	v1095 = *(*int64)(unsafe.Add(mBase, uint32(v1034)+56))
	if v1095 < v1035 {
		goto L308
	} else {
		goto L309
	}
L301:
	;
	v1072 = int32(0)
	v1075 = v1064 - v1069
	if v1075 < v1067 {
		goto L302
	} else {
		goto L303
	}
L302:
	;
	v1077 = v1072
	goto L304
L303:
	;
	v1077 = v1072 - v1067
	goto L304
L304:
	;
	if v1075 < int32(0) {
		goto L305
	} else {
		goto L306
	}
L305:
	;
	v1080 = v1067
	goto L307
L306:
	;
	v1080 = v1077
	goto L307
L307:
	;
	v1082 = v1080 + v1075
	goto L300
L308:
	;
	v1097 = v1035
	goto L310
L309:
	;
	v1097 = v1095
	goto L310
L310:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v1034)+56)) = v1097
	v1099 = *(*int64)(unsafe.Add(mBase, uint32(v1034)+48))
	if v1035 < v1099 {
		goto L311
	} else {
		goto L312
	}
L311:
	;
	v1101 = v1035
	goto L313
L312:
	;
	v1101 = v1099
	goto L313
L313:
	;
	if v1035 == int64(0) {
		goto L314
	} else {
		goto L315
	}
L314:
	;
	v1104 = v1099
	goto L316
L315:
	;
	v1104 = v1101
	goto L316
L316:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v1034)+48)) = v1104
	goto L297
L317:
	;
	v1143 = F_hashtableNext(m, v12+int32(720), v12+int32(24))
	mBase = m.M
	v1144 = m.ExcPending
	if v1144 != 0 {
		goto L24
	} else {
		goto L322
	}
L318:
	;
	goto L317
L319:
	;
	goto L318
L321:
	;
	F_hashtableCleanupIterator(m, v12+int32(720))
	mBase = m.M
	v1365 = m.ExcPending
	if v1365 != 0 {
		goto L24
	} else {
		goto L376
	}
L322:
	;
	if v1143 == int32(0) {
		goto L321
	} else {
		goto L323
	}
L323:
	;
	goto L324
L324:
	;
	v1156 = int32(0)
	v1158 = v12 + int32(32)
	v1159 = *(*int32)(unsafe.Add(mBase, uint32(v12)+24))
	v1160 = *(*float64)(unsafe.Add(mBase, uint32(v1159)))
	v1163 = F_fpconv_dtoa(m, v1160, v1158)
	mBase = m.M
	*(*uint8)(unsafe.Add(mBase, uint32(v1158+v1163))) = uint8(v1156)
	v1168 = v1159 + int32(16)
	v1169 = *(*int32)(unsafe.Add(mBase, uint32(v1168)))
	v1172 = v1168 + v1169<<(uint(int32(3))%32)
	v1173 = int32(*(*int8)(unsafe.Add(mBase, uint32(v1172))))
	v1174 = v1172 + v1173
	goto L332
L325:
	;
	goto L321
L326:
	;
	v1197 = *(*int32)(unsafe.Add(mBase, uint32(v25)+20))
	v1199 = v12 + int32(32)
	if v1199&int32(3) == int32(0) {
		v1221 = v1199
		goto L335
	} else {
		goto L336
	}
L327:
	;
	v1195 = *(*int32)(unsafe.Add(mBase, uint32(v1174+int32(-16))))
	v1196 = v1195
	goto L326
L328:
	;
	v1192 = *(*int32)(unsafe.Add(mBase, uint32(v1174+int32(-8))))
	v1196 = v1192
	goto L326
L329:
	;
	v1189 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1174+int32(-4)))))
	v1196 = v1189
	goto L326
L330:
	;
	v1186 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1174+int32(-2)))))
	v1196 = v1186
	goto L326
L331:
	;
	v1196 = int32(base.Ui32(v1179) >> (uint(int32(3)) % 32))
	goto L326
L332:
	;
	v1179 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1174+int32(0)))))
	switch v1179 & int32(7) {
	case 0:
		goto L331
	case 1:
		goto L330
	case 2:
		goto L329
	case 3:
		goto L328
	case 4:
		goto L327
	default:
		v1196 = v1156
		goto L326
	}
L333:
	;
	v1255 = v1254 + v1196
	*(*int32)(unsafe.Add(mBase, uint32(v25)+20)) = v1197 + v1255
	v1258 = *(*int32)(unsafe.Add(mBase, uint32(v25)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v25)+28)) = v1258 + v1255
	v1261 = *(*int32)(unsafe.Add(mBase, uint32(v25)+36))
	if base.Ui32(v1255) <= base.Ui32(v1261) {
		goto L349
	} else {
		goto L350
	}
L334:
	;
	v1254 = v1246 - v1199
	goto L333
L335:
	;
	v1225 = v1221
	goto L343
L336:
	;
	v1207 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1199))))
	if v1207 != 0 {
		goto L337
	} else {
		goto L338
	}
L337:
	;
	v1210 = v1199
	goto L339
L338:
	;
	v1254 = v1199 - v1199
	goto L333
L339:
	;
	v1214 = v1210 + int32(1)
	if v1214&int32(3) == int32(0) {
		v1221 = v1214
		goto L335
	} else {
		goto L341
	}
L341:
	;
	v1219 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1214))))
	if v1219 != 0 {
		v1210 = v1214
		goto L339
	} else {
		goto L342
	}
L342:
	;
	v1246 = v1214
	goto L334
L343:
	;
	v1231 = *(*int32)(unsafe.Add(mBase, uint32(v1225)))
	v1234 = int32(-2139062144)
	if (int32(16843008)-v1231|v1231)&v1234 == v1234 {
		v1225 = v1225 + int32(4)
		goto L343
	} else {
		goto L345
	}
L344:
	;
	v1240 = v1225
	goto L346
L345:
	;
	goto L344
L346:
	;
	v1244 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1240))))
	if v1244 != 0 {
		v1240 = v1240 + int32(1)
		goto L346
	} else {
		goto L348
	}
L347:
	;
	v1246 = v1240
	goto L334
L348:
	;
	goto L347
L349:
	;
	v1264 = *(*int32)(unsafe.Add(mBase, uint32(v25)+44))
	v1265 = base.I64_extend_i32_u(v1255)
	if int64(0) <= v1265 {
		goto L352
	} else {
		goto L353
	}
L350:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25)+36)) = v1255
	goto L349
L351:
	;
	v1351 = F_hashtableNext(m, v12+int32(720), v12+int32(24))
	mBase = m.M
	v1352 = m.ExcPending
	if v1352 != 0 {
		goto L24
	} else {
		goto L374
	}
L352:
	;
	v1276 = *(*int32)(unsafe.Add(mBase, uint32(v1264)+24))
	v1277 = *(*int32)(unsafe.Add(mBase, uint32(v1264)+16))
	v1279 = *(*int64)(unsafe.Add(mBase, uint32(v1264)+32))
	v1283 = v1276 + v1277 + base.I32_wrap_i64(base.I64_clz(v1279|v1265))
	v1286 = *(*int32)(unsafe.Add(mBase, uint32(v1264)+28))
	v1294 = (int32(64)-v1283)<<(uint(v1276)%32) - v1286 + base.I32_wrap_i64(int64(base.Ui64(v1265)>>(uint(base.I64_extend_i32_u(v1277-v1283+int32(63)))%64)))
	if v1294 < int32(0) {
		goto L354
	} else {
		goto L355
	}
L353:
	;
	goto L351
L354:
	;
	goto L351
L355:
	;
	v1297 = *(*int32)(unsafe.Add(mBase, uint32(v1264)+80))
	if v1297 <= v1294 {
		goto L354
	} else {
		goto L356
	}
L356:
	;
	v1299 = *(*int32)(unsafe.Add(mBase, uint32(v1264)+64))
	if v1299 == int32(0) {
		v1312 = v1294
		goto L357
	} else {
		goto L358
	}
L357:
	;
	v1313 = *(*int32)(unsafe.Add(mBase, uint32(v1264)+96))
	v1316 = v1313 + v1312<<(uint(int32(3))%32)
	v1317 = *(*int64)(unsafe.Add(mBase, uint32(v1316)))
	v1318 = int64(1)
	*(*int64)(unsafe.Add(mBase, uint32(v1316))) = v1317 + v1318
	v1321 = *(*int64)(unsafe.Add(mBase, uint32(v1264)+88))
	*(*int64)(unsafe.Add(mBase, uint32(v1264)+88)) = v1321 + v1318
	v1325 = *(*int64)(unsafe.Add(mBase, uint32(v1264)+56))
	if v1325 < v1265 {
		goto L365
	} else {
		goto L366
	}
L358:
	;
	v1302 = int32(0)
	v1305 = v1294 - v1299
	if v1305 < v1297 {
		goto L359
	} else {
		goto L360
	}
L359:
	;
	v1307 = v1302
	goto L361
L360:
	;
	v1307 = v1302 - v1297
	goto L361
L361:
	;
	if v1305 < int32(0) {
		goto L362
	} else {
		goto L363
	}
L362:
	;
	v1310 = v1297
	goto L364
L363:
	;
	v1310 = v1307
	goto L364
L364:
	;
	v1312 = v1310 + v1305
	goto L357
L365:
	;
	v1327 = v1265
	goto L367
L366:
	;
	v1327 = v1325
	goto L367
L367:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v1264)+56)) = v1327
	v1329 = *(*int64)(unsafe.Add(mBase, uint32(v1264)+48))
	if v1265 < v1329 {
		goto L368
	} else {
		goto L369
	}
L368:
	;
	v1331 = v1265
	goto L370
L369:
	;
	v1331 = v1329
	goto L370
L370:
	;
	if v1265 == int64(0) {
		goto L371
	} else {
		goto L372
	}
L371:
	;
	v1334 = v1329
	goto L373
L372:
	;
	v1334 = v1331
	goto L373
L373:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v1264)+48)) = v1334
	goto L354
L374:
	;
	if v1351 != 0 {
		goto L324
	} else {
		goto L375
	}
L375:
	;
	goto L325
L376:
	;
	v1366 = *(*int32)(unsafe.Add(mBase, uint32(v1119)))
	v1367 = *(*int32)(unsafe.Add(mBase, uint32(v1366)+20))
	v1368 = *(*int32)(unsafe.Add(mBase, uint32(v1366)+16))
	v1369 = v1367 + v1368
	goto L377
L377:
	;
	if v25 == int32(0) {
		goto L13
	} else {
		goto L378
	}
L378:
	;
	v1372 = *(*int32)(unsafe.Add(mBase, uint32(v25)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v25)+24)) = v1372 + v1369
	v1375 = *(*int32)(unsafe.Add(mBase, uint32(v25)+32))
	if base.Ui32(v1369) <= base.Ui32(v1375) {
		goto L379
	} else {
		goto L380
	}
L379:
	;
	v1378 = *(*int32)(unsafe.Add(mBase, uint32(v25)+40))
	v1379 = base.I64_extend_i32_u(v1369)
	if int64(0) <= v1379 {
		goto L382
	} else {
		goto L383
	}
L380:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25)+32)) = v1369
	goto L379
L381:
	;
	goto L13
L382:
	;
	v1390 = *(*int32)(unsafe.Add(mBase, uint32(v1378)+24))
	v1391 = *(*int32)(unsafe.Add(mBase, uint32(v1378)+16))
	v1393 = *(*int64)(unsafe.Add(mBase, uint32(v1378)+32))
	v1397 = v1390 + v1391 + base.I32_wrap_i64(base.I64_clz(v1393|v1379))
	v1400 = *(*int32)(unsafe.Add(mBase, uint32(v1378)+28))
	v1408 = (int32(64)-v1397)<<(uint(v1390)%32) - v1400 + base.I32_wrap_i64(int64(base.Ui64(v1379)>>(uint(base.I64_extend_i32_u(v1391-v1397+int32(63)))%64)))
	if v1408 < int32(0) {
		goto L384
	} else {
		goto L385
	}
L383:
	;
	goto L381
L384:
	;
	goto L381
L385:
	;
	v1411 = *(*int32)(unsafe.Add(mBase, uint32(v1378)+80))
	if v1411 <= v1408 {
		goto L384
	} else {
		goto L386
	}
L386:
	;
	v1413 = *(*int32)(unsafe.Add(mBase, uint32(v1378)+64))
	if v1413 == int32(0) {
		v1426 = v1408
		goto L387
	} else {
		goto L388
	}
L387:
	;
	v1427 = *(*int32)(unsafe.Add(mBase, uint32(v1378)+96))
	v1430 = v1427 + v1426<<(uint(int32(3))%32)
	v1431 = *(*int64)(unsafe.Add(mBase, uint32(v1430)))
	v1432 = int64(1)
	*(*int64)(unsafe.Add(mBase, uint32(v1430))) = v1431 + v1432
	v1435 = *(*int64)(unsafe.Add(mBase, uint32(v1378)+88))
	*(*int64)(unsafe.Add(mBase, uint32(v1378)+88)) = v1435 + v1432
	v1439 = *(*int64)(unsafe.Add(mBase, uint32(v1378)+56))
	if v1439 < v1379 {
		goto L395
	} else {
		goto L396
	}
L388:
	;
	v1416 = int32(0)
	v1419 = v1408 - v1413
	if v1419 < v1411 {
		goto L389
	} else {
		goto L390
	}
L389:
	;
	v1421 = v1416
	goto L391
L390:
	;
	v1421 = v1416 - v1411
	goto L391
L391:
	;
	if v1419 < int32(0) {
		goto L392
	} else {
		goto L393
	}
L392:
	;
	v1424 = v1411
	goto L394
L393:
	;
	v1424 = v1421
	goto L394
L394:
	;
	v1426 = v1424 + v1419
	goto L387
L395:
	;
	v1441 = v1379
	goto L397
L396:
	;
	v1441 = v1439
	goto L397
L397:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v1378)+56)) = v1441
	v1443 = *(*int64)(unsafe.Add(mBase, uint32(v1378)+48))
	if v1379 < v1443 {
		goto L398
	} else {
		goto L399
	}
L398:
	;
	v1445 = v1379
	goto L400
L399:
	;
	v1445 = v1443
	goto L400
L400:
	;
	if v1379 == int64(0) {
		goto L401
	} else {
		goto L402
	}
L401:
	;
	v1448 = v1443
	goto L403
L402:
	;
	v1448 = v1445
	goto L403
L403:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v1378)+48)) = v1448
	goto L384
L404:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L405:
	;
	v1474 = F_hashTypeNext(m, v12+int32(32))
	mBase = m.M
	v1475 = m.ExcPending
	if v1475 != 0 {
		goto L24
	} else {
		goto L407
	}
L406:
	;
	F_hashTypeResetIterator(m, v12+int32(32))
	mBase = m.M
	v1654 = m.ExcPending
	if v1654 != 0 {
		goto L24
	} else {
		goto L454
	}
L407:
	;
	if v1474 == int32(-1) {
		goto L406
	} else {
		goto L408
	}
L408:
	;
	goto L409
L409:
	;
	v1487 = int32(0)
	v1492 = F_hashTypeCurrentObjectNewSds(m, v12+int32(32), int32(1))
	mBase = m.M
	v1493 = m.ExcPending
	if v1493 != 0 {
		goto L24
	} else {
		goto L417
	}
L410:
	;
	goto L406
L411:
	;
	F_sdsfree(m, v1492)
	mBase = m.M
	v1515 = m.ExcPending
	if v1515 != 0 {
		goto L24
	} else {
		goto L418
	}
L412:
	;
	v1512 = *(*int32)(unsafe.Add(mBase, uint32(v1492+int32(-17))))
	v1513 = v1512
	goto L411
L413:
	;
	v1509 = *(*int32)(unsafe.Add(mBase, uint32(v1492+int32(-9))))
	v1513 = v1509
	goto L411
L414:
	;
	v1506 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1492+int32(-5)))))
	v1513 = v1506
	goto L411
L415:
	;
	v1503 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1492+int32(-3)))))
	v1513 = v1503
	goto L411
L416:
	;
	v1513 = int32(base.Ui32(v1496) >> (uint(int32(3)) % 32))
	goto L411
L417:
	;
	v1496 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1492+int32(-1)))))
	switch v1496 & int32(7) {
	case 0:
		goto L416
	case 1:
		goto L415
	case 2:
		goto L414
	case 3:
		goto L413
	case 4:
		goto L412
	default:
		v1513 = v1487
		goto L411
	}
L418:
	;
	v1519 = F_hashTypeCurrentObjectNewSds(m, v12+int32(32), int32(2))
	mBase = m.M
	v1520 = m.ExcPending
	if v1520 != 0 {
		goto L24
	} else {
		goto L425
	}
L419:
	;
	F_sdsfree(m, v1519)
	mBase = m.M
	v1542 = m.ExcPending
	if v1542 != 0 {
		goto L24
	} else {
		goto L426
	}
L420:
	;
	v1539 = *(*int32)(unsafe.Add(mBase, uint32(v1519+int32(-17))))
	v1540 = v1539
	goto L419
L421:
	;
	v1536 = *(*int32)(unsafe.Add(mBase, uint32(v1519+int32(-9))))
	v1540 = v1536
	goto L419
L422:
	;
	v1533 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1519+int32(-5)))))
	v1540 = v1533
	goto L419
L423:
	;
	v1530 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1519+int32(-3)))))
	v1540 = v1530
	goto L419
L424:
	;
	v1540 = int32(base.Ui32(v1523) >> (uint(int32(3)) % 32))
	goto L419
L425:
	;
	v1523 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1519+int32(-1)))))
	switch v1523 & int32(7) {
	case 0:
		goto L424
	case 1:
		goto L423
	case 2:
		goto L422
	case 3:
		goto L421
	case 4:
		goto L420
	default:
		v1540 = v1487
		goto L419
	}
L426:
	;
	v1543 = *(*int32)(unsafe.Add(mBase, uint32(v25)+20))
	v1544 = v1540 + v1513
	*(*int32)(unsafe.Add(mBase, uint32(v25)+20)) = v1543 + v1544
	v1547 = *(*int32)(unsafe.Add(mBase, uint32(v25)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v25)+28)) = v1547 + v1544
	v1550 = *(*int32)(unsafe.Add(mBase, uint32(v25)+36))
	if base.Ui32(v1544) <= base.Ui32(v1550) {
		goto L427
	} else {
		goto L428
	}
L427:
	;
	v1553 = *(*int32)(unsafe.Add(mBase, uint32(v25)+44))
	v1554 = base.I64_extend_i32_u(v1544)
	if int64(0) <= v1554 {
		goto L430
	} else {
		goto L431
	}
L428:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25)+36)) = v1544
	goto L427
L429:
	;
	v1638 = F_hashTypeNext(m, v12+int32(32))
	mBase = m.M
	v1639 = m.ExcPending
	if v1639 != 0 {
		goto L24
	} else {
		goto L452
	}
L430:
	;
	v1565 = *(*int32)(unsafe.Add(mBase, uint32(v1553)+24))
	v1566 = *(*int32)(unsafe.Add(mBase, uint32(v1553)+16))
	v1568 = *(*int64)(unsafe.Add(mBase, uint32(v1553)+32))
	v1572 = v1565 + v1566 + base.I32_wrap_i64(base.I64_clz(v1568|v1554))
	v1575 = *(*int32)(unsafe.Add(mBase, uint32(v1553)+28))
	v1583 = (int32(64)-v1572)<<(uint(v1565)%32) - v1575 + base.I32_wrap_i64(int64(base.Ui64(v1554)>>(uint(base.I64_extend_i32_u(v1566-v1572+int32(63)))%64)))
	if v1583 < int32(0) {
		goto L432
	} else {
		goto L433
	}
L431:
	;
	goto L429
L432:
	;
	goto L429
L433:
	;
	v1586 = *(*int32)(unsafe.Add(mBase, uint32(v1553)+80))
	if v1586 <= v1583 {
		goto L432
	} else {
		goto L434
	}
L434:
	;
	v1588 = *(*int32)(unsafe.Add(mBase, uint32(v1553)+64))
	if v1588 == int32(0) {
		v1601 = v1583
		goto L435
	} else {
		goto L436
	}
L435:
	;
	v1602 = *(*int32)(unsafe.Add(mBase, uint32(v1553)+96))
	v1605 = v1602 + v1601<<(uint(int32(3))%32)
	v1606 = *(*int64)(unsafe.Add(mBase, uint32(v1605)))
	v1607 = int64(1)
	*(*int64)(unsafe.Add(mBase, uint32(v1605))) = v1606 + v1607
	v1610 = *(*int64)(unsafe.Add(mBase, uint32(v1553)+88))
	*(*int64)(unsafe.Add(mBase, uint32(v1553)+88)) = v1610 + v1607
	v1614 = *(*int64)(unsafe.Add(mBase, uint32(v1553)+56))
	if v1614 < v1554 {
		goto L443
	} else {
		goto L444
	}
L436:
	;
	v1591 = int32(0)
	v1594 = v1583 - v1588
	if v1594 < v1586 {
		goto L437
	} else {
		goto L438
	}
L437:
	;
	v1596 = v1591
	goto L439
L438:
	;
	v1596 = v1591 - v1586
	goto L439
L439:
	;
	if v1594 < int32(0) {
		goto L440
	} else {
		goto L441
	}
L440:
	;
	v1599 = v1586
	goto L442
L441:
	;
	v1599 = v1596
	goto L442
L442:
	;
	v1601 = v1599 + v1594
	goto L435
L443:
	;
	v1616 = v1554
	goto L445
L444:
	;
	v1616 = v1614
	goto L445
L445:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v1553)+56)) = v1616
	v1618 = *(*int64)(unsafe.Add(mBase, uint32(v1553)+48))
	if v1554 < v1618 {
		goto L446
	} else {
		goto L447
	}
L446:
	;
	v1620 = v1554
	goto L448
L447:
	;
	v1620 = v1618
	goto L448
L448:
	;
	if v1554 == int64(0) {
		goto L449
	} else {
		goto L450
	}
L449:
	;
	v1623 = v1618
	goto L451
L450:
	;
	v1623 = v1620
	goto L451
L451:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v1553)+48)) = v1623
	goto L432
L452:
	;
	if v1638 != int32(-1) {
		goto L409
	} else {
		goto L453
	}
L453:
	;
	goto L410
L454:
	;
	v1655 = F_hashTypeLength(m, l2)
	mBase = m.M
	v1656 = m.ExcPending
	if v1656 != 0 {
		goto L24
	} else {
		goto L455
	}
L455:
	;
	if v25 == int32(0) {
		goto L13
	} else {
		goto L456
	}
L456:
	;
	v1659 = *(*int32)(unsafe.Add(mBase, uint32(v25)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v25)+24)) = v1659 + v1655
	v1662 = *(*int32)(unsafe.Add(mBase, uint32(v25)+32))
	if base.Ui32(v1655) <= base.Ui32(v1662) {
		goto L457
	} else {
		goto L458
	}
L457:
	;
	v1665 = *(*int32)(unsafe.Add(mBase, uint32(v25)+40))
	v1666 = base.I64_extend_i32_u(v1655)
	if int64(0) <= v1666 {
		goto L460
	} else {
		goto L461
	}
L458:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25)+32)) = v1655
	goto L457
L459:
	;
	goto L13
L460:
	;
	v1677 = *(*int32)(unsafe.Add(mBase, uint32(v1665)+24))
	v1678 = *(*int32)(unsafe.Add(mBase, uint32(v1665)+16))
	v1680 = *(*int64)(unsafe.Add(mBase, uint32(v1665)+32))
	v1684 = v1677 + v1678 + base.I32_wrap_i64(base.I64_clz(v1680|v1666))
	v1687 = *(*int32)(unsafe.Add(mBase, uint32(v1665)+28))
	v1695 = (int32(64)-v1684)<<(uint(v1677)%32) - v1687 + base.I32_wrap_i64(int64(base.Ui64(v1666)>>(uint(base.I64_extend_i32_u(v1678-v1684+int32(63)))%64)))
	if v1695 < int32(0) {
		goto L462
	} else {
		goto L463
	}
L461:
	;
	goto L459
L462:
	;
	goto L459
L463:
	;
	v1698 = *(*int32)(unsafe.Add(mBase, uint32(v1665)+80))
	if v1698 <= v1695 {
		goto L462
	} else {
		goto L464
	}
L464:
	;
	v1700 = *(*int32)(unsafe.Add(mBase, uint32(v1665)+64))
	if v1700 == int32(0) {
		v1713 = v1695
		goto L465
	} else {
		goto L466
	}
L465:
	;
	v1714 = *(*int32)(unsafe.Add(mBase, uint32(v1665)+96))
	v1717 = v1714 + v1713<<(uint(int32(3))%32)
	v1718 = *(*int64)(unsafe.Add(mBase, uint32(v1717)))
	v1719 = int64(1)
	*(*int64)(unsafe.Add(mBase, uint32(v1717))) = v1718 + v1719
	v1722 = *(*int64)(unsafe.Add(mBase, uint32(v1665)+88))
	*(*int64)(unsafe.Add(mBase, uint32(v1665)+88)) = v1722 + v1719
	v1726 = *(*int64)(unsafe.Add(mBase, uint32(v1665)+56))
	if v1726 < v1666 {
		goto L473
	} else {
		goto L474
	}
L466:
	;
	v1703 = int32(0)
	v1706 = v1695 - v1700
	if v1706 < v1698 {
		goto L467
	} else {
		goto L468
	}
L467:
	;
	v1708 = v1703
	goto L469
L468:
	;
	v1708 = v1703 - v1698
	goto L469
L469:
	;
	if v1706 < int32(0) {
		goto L470
	} else {
		goto L471
	}
L470:
	;
	v1711 = v1698
	goto L472
L471:
	;
	v1711 = v1708
	goto L472
L472:
	;
	v1713 = v1711 + v1706
	goto L465
L473:
	;
	v1728 = v1666
	goto L475
L474:
	;
	v1728 = v1726
	goto L475
L475:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v1665)+56)) = v1728
	v1730 = *(*int64)(unsafe.Add(mBase, uint32(v1665)+48))
	if v1666 < v1730 {
		goto L476
	} else {
		goto L477
	}
L476:
	;
	v1732 = v1666
	goto L478
L477:
	;
	v1732 = v1730
	goto L478
L478:
	;
	if v1666 == int64(0) {
		goto L479
	} else {
		goto L480
	}
L479:
	;
	v1735 = v1730
	goto L481
L480:
	;
	v1735 = v1732
	goto L481
L481:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v1665)+48)) = v1735
	goto L462
L482:
	;
	v1762 = F_streamIteratorGetID(m, v12+int32(32), v12+int32(720), v12+int32(24))
	mBase = m.M
	v1763 = m.ExcPending
	if v1763 != 0 {
		goto L24
	} else {
		goto L484
	}
L483:
	;
	F_streamIteratorStop(m, v12+int32(32))
	mBase = m.M
	v1934 = m.ExcPending
	if v1934 != 0 {
		goto L24
	} else {
		goto L521
	}
L484:
	;
	if v1762 == int32(0) {
		goto L483
	} else {
		goto L485
	}
L485:
	;
	goto L486
L486:
	;
	v1775 = *(*int64)(unsafe.Add(mBase, uint32(v12)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v12)+24)) = v1775 + int64(-1)
	if v1775 == int64(0) {
		goto L488
	} else {
		goto L489
	}
L487:
	;
	goto L483
L488:
	;
	v1920 = F_streamIteratorGetID(m, v12+int32(32), v12+int32(720), v12+int32(24))
	mBase = m.M
	v1921 = m.ExcPending
	if v1921 != 0 {
		goto L24
	} else {
		goto L519
	}
L489:
	;
	goto L490
L490:
	;
	F_streamIteratorGetField(m, v12+int32(32), v12+int32(20), v12+int32(16), v12+int32(8), v12)
	mBase = m.M
	v1799 = m.ExcPending
	if v1799 != 0 {
		goto L24
	} else {
		goto L492
	}
L491:
	;
	goto L488
L492:
	;
	v1800 = *(*int32)(unsafe.Add(mBase, uint32(v25)+20))
	v1801 = *(*int64)(unsafe.Add(mBase, uint32(v12)))
	v1802 = *(*int64)(unsafe.Add(mBase, uint32(v12)+8))
	v1803 = v1801 + v1802
	v1804 = base.I32_wrap_i64(v1803)
	*(*int32)(unsafe.Add(mBase, uint32(v25)+20)) = v1800 + v1804
	v1807 = *(*int32)(unsafe.Add(mBase, uint32(v25)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v25)+28)) = v1807 + v1804
	v1810 = *(*int32)(unsafe.Add(mBase, uint32(v25)+36))
	if base.Ui32(v1804) <= base.Ui32(v1810) {
		goto L493
	} else {
		goto L494
	}
L493:
	;
	v1813 = *(*int32)(unsafe.Add(mBase, uint32(v25)+44))
	v1815 = v1803 & int64(4294967295)
	if int64(0) <= v1815 {
		goto L496
	} else {
		goto L497
	}
L494:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25)+36)) = v1804
	goto L493
L495:
	;
	v1897 = *(*int64)(unsafe.Add(mBase, uint32(v12)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v12)+24)) = v1897 + int64(-1)
	if base.B2i32(v1897 == int64(0)) == int32(0) {
		goto L490
	} else {
		goto L518
	}
L496:
	;
	v1826 = *(*int32)(unsafe.Add(mBase, uint32(v1813)+24))
	v1827 = *(*int32)(unsafe.Add(mBase, uint32(v1813)+16))
	v1829 = *(*int64)(unsafe.Add(mBase, uint32(v1813)+32))
	v1833 = v1826 + v1827 + base.I32_wrap_i64(base.I64_clz(v1829|v1815))
	v1836 = *(*int32)(unsafe.Add(mBase, uint32(v1813)+28))
	v1844 = (int32(64)-v1833)<<(uint(v1826)%32) - v1836 + base.I32_wrap_i64(int64(base.Ui64(v1815)>>(uint(base.I64_extend_i32_u(v1827-v1833+int32(63)))%64)))
	if v1844 < int32(0) {
		goto L498
	} else {
		goto L499
	}
L497:
	;
	goto L495
L498:
	;
	goto L495
L499:
	;
	v1847 = *(*int32)(unsafe.Add(mBase, uint32(v1813)+80))
	if v1847 <= v1844 {
		goto L498
	} else {
		goto L500
	}
L500:
	;
	v1849 = *(*int32)(unsafe.Add(mBase, uint32(v1813)+64))
	if v1849 == int32(0) {
		v1862 = v1844
		goto L501
	} else {
		goto L502
	}
L501:
	;
	v1863 = *(*int32)(unsafe.Add(mBase, uint32(v1813)+96))
	v1866 = v1863 + v1862<<(uint(int32(3))%32)
	v1867 = *(*int64)(unsafe.Add(mBase, uint32(v1866)))
	v1868 = int64(1)
	*(*int64)(unsafe.Add(mBase, uint32(v1866))) = v1867 + v1868
	v1871 = *(*int64)(unsafe.Add(mBase, uint32(v1813)+88))
	*(*int64)(unsafe.Add(mBase, uint32(v1813)+88)) = v1871 + v1868
	v1875 = *(*int64)(unsafe.Add(mBase, uint32(v1813)+56))
	if v1875 < v1815 {
		goto L509
	} else {
		goto L510
	}
L502:
	;
	v1852 = int32(0)
	v1855 = v1844 - v1849
	if v1855 < v1847 {
		goto L503
	} else {
		goto L504
	}
L503:
	;
	v1857 = v1852
	goto L505
L504:
	;
	v1857 = v1852 - v1847
	goto L505
L505:
	;
	if v1855 < int32(0) {
		goto L506
	} else {
		goto L507
	}
L506:
	;
	v1860 = v1847
	goto L508
L507:
	;
	v1860 = v1857
	goto L508
L508:
	;
	v1862 = v1860 + v1855
	goto L501
L509:
	;
	v1877 = v1815
	goto L511
L510:
	;
	v1877 = v1875
	goto L511
L511:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v1813)+56)) = v1877
	v1879 = *(*int64)(unsafe.Add(mBase, uint32(v1813)+48))
	if v1815 < v1879 {
		goto L512
	} else {
		goto L513
	}
L512:
	;
	v1881 = v1815
	goto L514
L513:
	;
	v1881 = v1879
	goto L514
L514:
	;
	if v1815 == int64(0) {
		goto L515
	} else {
		goto L516
	}
L515:
	;
	v1884 = v1879
	goto L517
L516:
	;
	v1884 = v1881
	goto L517
L517:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v1813)+48)) = v1884
	goto L498
L518:
	;
	goto L491
L519:
	;
	if v1920 != 0 {
		goto L486
	} else {
		goto L520
	}
L520:
	;
	goto L487
L521:
	;
	v1935 = F_objectGetVal(m, l2)
	mBase = m.M
	v1936 = *(*int32)(unsafe.Add(mBase, uint32(v1935)+8))
	goto L522
L522:
	;
	if v25 == int32(0) {
		goto L13
	} else {
		goto L523
	}
L523:
	;
	v1939 = *(*int32)(unsafe.Add(mBase, uint32(v25)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v25)+24)) = v1939 + v1936
	v1942 = *(*int32)(unsafe.Add(mBase, uint32(v25)+32))
	if base.Ui32(v1936) <= base.Ui32(v1942) {
		goto L524
	} else {
		goto L525
	}
L524:
	;
	v1945 = *(*int32)(unsafe.Add(mBase, uint32(v25)+40))
	v1946 = base.I64_extend_i32_u(v1936)
	if int64(0) <= v1946 {
		goto L527
	} else {
		goto L528
	}
L525:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25)+32)) = v1936
	goto L524
L526:
	;
	goto L13
L527:
	;
	v1957 = *(*int32)(unsafe.Add(mBase, uint32(v1945)+24))
	v1958 = *(*int32)(unsafe.Add(mBase, uint32(v1945)+16))
	v1960 = *(*int64)(unsafe.Add(mBase, uint32(v1945)+32))
	v1964 = v1957 + v1958 + base.I32_wrap_i64(base.I64_clz(v1960|v1946))
	v1967 = *(*int32)(unsafe.Add(mBase, uint32(v1945)+28))
	v1975 = (int32(64)-v1964)<<(uint(v1957)%32) - v1967 + base.I32_wrap_i64(int64(base.Ui64(v1946)>>(uint(base.I64_extend_i32_u(v1958-v1964+int32(63)))%64)))
	if v1975 < int32(0) {
		goto L529
	} else {
		goto L530
	}
L528:
	;
	goto L526
L529:
	;
	goto L526
L530:
	;
	v1978 = *(*int32)(unsafe.Add(mBase, uint32(v1945)+80))
	if v1978 <= v1975 {
		goto L529
	} else {
		goto L531
	}
L531:
	;
	v1980 = *(*int32)(unsafe.Add(mBase, uint32(v1945)+64))
	if v1980 == int32(0) {
		v1993 = v1975
		goto L532
	} else {
		goto L533
	}
L532:
	;
	v1994 = *(*int32)(unsafe.Add(mBase, uint32(v1945)+96))
	v1997 = v1994 + v1993<<(uint(int32(3))%32)
	v1998 = *(*int64)(unsafe.Add(mBase, uint32(v1997)))
	v1999 = int64(1)
	*(*int64)(unsafe.Add(mBase, uint32(v1997))) = v1998 + v1999
	v2002 = *(*int64)(unsafe.Add(mBase, uint32(v1945)+88))
	*(*int64)(unsafe.Add(mBase, uint32(v1945)+88)) = v2002 + v1999
	v2006 = *(*int64)(unsafe.Add(mBase, uint32(v1945)+56))
	if v2006 < v1946 {
		goto L540
	} else {
		goto L541
	}
L533:
	;
	v1983 = int32(0)
	v1986 = v1975 - v1980
	if v1986 < v1978 {
		goto L534
	} else {
		goto L535
	}
L534:
	;
	v1988 = v1983
	goto L536
L535:
	;
	v1988 = v1983 - v1978
	goto L536
L536:
	;
	if v1986 < int32(0) {
		goto L537
	} else {
		goto L538
	}
L537:
	;
	v1991 = v1978
	goto L539
L538:
	;
	v1991 = v1988
	goto L539
L539:
	;
	v1993 = v1991 + v1986
	goto L532
L540:
	;
	v2008 = v1946
	goto L542
L541:
	;
	v2008 = v2006
	goto L542
L542:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v1945)+56)) = v2008
	v2010 = *(*int64)(unsafe.Add(mBase, uint32(v1945)+48))
	if v1946 < v2010 {
		goto L543
	} else {
		goto L544
	}
L543:
	;
	v2012 = v1946
	goto L545
L544:
	;
	v2012 = v2010
	goto L545
L545:
	;
	if v1946 == int64(0) {
		goto L546
	} else {
		goto L547
	}
L546:
	;
	v2015 = v2010
	goto L548
L547:
	;
	v2015 = v2012
	goto L548
L548:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v1945)+48)) = v2015
	goto L529
L549:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L550:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L551:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L552:
	;
	v2054 = *(*int32)(unsafe.Add(mBase, uint32(v25)+40))
	v2055 = int64(1)
	goto L555
L553:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25)+32)) = int32(1)
	goto L552
L554:
	;
	goto L13
L555:
	;
	v2066 = *(*int32)(unsafe.Add(mBase, uint32(v2054)+24))
	v2067 = *(*int32)(unsafe.Add(mBase, uint32(v2054)+16))
	v2069 = *(*int64)(unsafe.Add(mBase, uint32(v2054)+32))
	v2073 = v2066 + v2067 + base.I32_wrap_i64(base.I64_clz(v2069|v2055))
	v2076 = *(*int32)(unsafe.Add(mBase, uint32(v2054)+28))
	v2084 = (int32(64)-v2073)<<(uint(v2066)%32) - v2076 + base.I32_wrap_i64(int64(base.Ui64(v2055)>>(uint(base.I64_extend_i32_u(v2067-v2073+int32(63)))%64)))
	if v2084 < int32(0) {
		goto L557
	} else {
		goto L558
	}
L557:
	;
	goto L554
L558:
	;
	v2087 = *(*int32)(unsafe.Add(mBase, uint32(v2054)+80))
	if v2087 <= v2084 {
		goto L557
	} else {
		goto L559
	}
L559:
	;
	v2089 = *(*int32)(unsafe.Add(mBase, uint32(v2054)+64))
	if v2089 == int32(0) {
		v2102 = v2084
		goto L560
	} else {
		goto L561
	}
L560:
	;
	v2103 = *(*int32)(unsafe.Add(mBase, uint32(v2054)+96))
	v2106 = v2103 + v2102<<(uint(int32(3))%32)
	v2107 = *(*int64)(unsafe.Add(mBase, uint32(v2106)))
	v2108 = int64(1)
	*(*int64)(unsafe.Add(mBase, uint32(v2106))) = v2107 + v2108
	v2111 = *(*int64)(unsafe.Add(mBase, uint32(v2054)+88))
	*(*int64)(unsafe.Add(mBase, uint32(v2054)+88)) = v2111 + v2108
	v2115 = *(*int64)(unsafe.Add(mBase, uint32(v2054)+56))
	if v2115 < v2055 {
		goto L568
	} else {
		goto L569
	}
L561:
	;
	v2092 = int32(0)
	v2095 = v2084 - v2089
	if v2095 < v2087 {
		goto L562
	} else {
		goto L563
	}
L562:
	;
	v2097 = v2092
	goto L564
L563:
	;
	v2097 = v2092 - v2087
	goto L564
L564:
	;
	if v2095 < int32(0) {
		goto L565
	} else {
		goto L566
	}
L565:
	;
	v2100 = v2087
	goto L567
L566:
	;
	v2100 = v2097
	goto L567
L567:
	;
	v2102 = v2100 + v2095
	goto L560
L568:
	;
	v2117 = v2055
	goto L570
L569:
	;
	v2117 = v2115
	goto L570
L570:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v2054)+56)) = v2117
	v2119 = *(*int64)(unsafe.Add(mBase, uint32(v2054)+48))
	if v2055 < v2119 {
		goto L571
	} else {
		goto L572
	}
L571:
	;
	v2121 = v2055
	goto L573
L572:
	;
	v2121 = v2119
	goto L573
L573:
	;
	goto L575
L575:
	;
	goto L576
L576:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v2054)+48)) = v2121
	goto L557
}
func F_connectionTypeTls(m *base.Module) int32 {
	mBase := m.M
	_ = mBase
	var v1 int32
	_ = v1
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	v1 = int32(0)
	v3 = m.G0
	v5 = v3 - int32(16)
	m.G0 = v5
	v8 = int32(*(*uint8)(unsafe.Add(mBase, _consts[278])))
	if v8 == v1 {
		v13 = int32(0)
		v14 = int32(1)
		*(*uint8)(unsafe.Add(mBase, _consts[278])) = uint8(v14)
		v17 = *(*int32)(unsafe.Add(mBase, _consts[279]))
		if v17 != 0 {
			*(*int32)(unsafe.Add(mBase, _consts[280])) = v17
			v32 = v17
			m.G0 = v5 + int32(16)
			return v32
		} else {
			v19 = *(*int32)(unsafe.Add(mBase, _consts[28]))
			if int32(3) < v19 {
				*(*int32)(unsafe.Add(mBase, _consts[280])) = v17
				v32 = v17
				m.G0 = v5 + int32(16)
				return v32
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v5))) = int32(_a553)
				F__serverLog(m, int32(3), int32(_a554), v5)
				mBase = m.M
				v29 = m.ExcPending
				if v29 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, _consts[280])) = v17
					v32 = v17
					m.G0 = v5 + int32(16)
					return v32
				}
			}
		}
	} else {
		v12 = *(*int32)(unsafe.Add(mBase, _consts[280]))
		v32 = v12
		m.G0 = v5 + int32(16)
		return v32
	}
}
func F_continueCommandHandler(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	return int32(0)
}
func F_cos(m *base.Module, l0 float64) float64 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v14 int32
	_ = v14
	var v24 float64
	_ = v24
	var v25 float64
	_ = v25
	var v27 float64
	_ = v27
	var v28 float64
	_ = v28
	var v40 float64
	_ = v40
	var v59 int32
	_ = v59
	var v60 float64
	_ = v60
	var v61 float64
	_ = v61
	var v67 float64
	_ = v67
	var v68 float64
	_ = v68
	var v70 float64
	_ = v70
	var v71 float64
	_ = v71
	var v83 float64
	_ = v83
	var v103 float64
	_ = v103
	var v119 float64
	_ = v119
	var v140 float64
	_ = v140
	var v141 float64
	_ = v141
	var v143 float64
	_ = v143
	var v144 float64
	_ = v144
	var v156 float64
	_ = v156
	var v177 float64
	_ = v177
	var v193 float64
	_ = v193
	var v212 float64
	_ = v212
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v14 = base.I32_wrap_i64(int64(base.Ui64(base.I64_reinterpret_f64(l0))>>(uint(int64(32))%64))) & int32(2147483647)
	if base.Ui32(int32(1072243195)) < base.Ui32(v14) {
		if base.Ui32(v14) < base.Ui32(int32(2146435072)) {
			v59 = F___rem_pio2(m, l0, v7)
			mBase = m.M
			v60 = *(*float64)(unsafe.Add(mBase, uint32(v7)+8))
			v61 = *(*float64)(unsafe.Add(mBase, uint32(v7)))
			switch v59 & int32(3) {
			default:
				v67 = float64(1)
				v68 = base.F64_mul(v61, v61)
				v70 = base.F64_mul(v68, float64(0.5))
				v71 = base.F64_sub(v67, v70)
				v83 = base.F64_mul(v68, v68)
				v212 = base.F64_add(v71, base.F64_add(base.F64_sub(base.F64_sub(v67, v71), v70), base.F64_sub(base.F64_mul(v68, base.F64_add(base.F64_mul(v68, base.F64_add(base.F64_mul(v68, base.F64_add(base.F64_mul(v68, float64(2.480158728947673e-05)), float64(-0.001388888888887411))), float64(0.0416666666666666))), base.F64_mul(base.F64_mul(v83, v83), base.F64_add(base.F64_mul(v68, base.F64_add(base.F64_mul(v68, float64(-1.1359647557788195e-11)), float64(2.087572321298175e-09))), float64(-2.7557314351390663e-07))))), base.F64_mul(v61, v60))))
			case 1:
				v103 = base.F64_mul(v61, v61)
				v119 = base.F64_mul(v61, v103)
				v212 = base.F64_neg(base.F64_sub(v61, base.F64_add(base.F64_sub(base.F64_mul(v103, base.F64_sub(base.F64_mul(v60, float64(0.5)), base.F64_mul(v119, base.F64_add(base.F64_mul(base.F64_mul(v103, base.F64_mul(v103, v103)), base.F64_add(base.F64_mul(v103, float64(1.58969099521155e-10)), float64(-2.5050760253406863e-08))), base.F64_add(base.F64_mul(v103, base.F64_add(base.F64_mul(v103, float64(2.7557313707070068e-06)), float64(-0.0001984126982985795))), float64(0.00833333333332249)))))), v60), base.F64_mul(v119, float64(0.16666666666666632)))))
			case 2:
				v140 = float64(1)
				v141 = base.F64_mul(v61, v61)
				v143 = base.F64_mul(v141, float64(0.5))
				v144 = base.F64_sub(v140, v143)
				v156 = base.F64_mul(v141, v141)
				v212 = base.F64_neg(base.F64_add(v144, base.F64_add(base.F64_sub(base.F64_sub(v140, v144), v143), base.F64_sub(base.F64_mul(v141, base.F64_add(base.F64_mul(v141, base.F64_add(base.F64_mul(v141, base.F64_add(base.F64_mul(v141, float64(2.480158728947673e-05)), float64(-0.001388888888887411))), float64(0.0416666666666666))), base.F64_mul(base.F64_mul(v156, v156), base.F64_add(base.F64_mul(v141, base.F64_add(base.F64_mul(v141, float64(-1.1359647557788195e-11)), float64(2.087572321298175e-09))), float64(-2.7557314351390663e-07))))), base.F64_mul(v61, v60)))))
			case 3:
				v177 = base.F64_mul(v61, v61)
				v193 = base.F64_mul(v61, v177)
				v212 = base.F64_sub(v61, base.F64_add(base.F64_sub(base.F64_mul(v177, base.F64_sub(base.F64_mul(v60, float64(0.5)), base.F64_mul(v193, base.F64_add(base.F64_mul(base.F64_mul(v177, base.F64_mul(v177, v177)), base.F64_add(base.F64_mul(v177, float64(1.58969099521155e-10)), float64(-2.5050760253406863e-08))), base.F64_add(base.F64_mul(v177, base.F64_add(base.F64_mul(v177, float64(2.7557313707070068e-06)), float64(-0.0001984126982985795))), float64(0.00833333333332249)))))), v60), base.F64_mul(v193, float64(0.16666666666666632))))
			}
		} else {
			v212 = base.F64_sub(l0, l0)
		}
	} else {
		if base.Ui32(v14) < base.Ui32(int32(1044816030)) {
			v212 = float64(1)
		} else {
			v24 = float64(1)
			v25 = base.F64_mul(l0, l0)
			v27 = base.F64_mul(v25, float64(0.5))
			v28 = base.F64_sub(v24, v27)
			v40 = base.F64_mul(v25, v25)
			v212 = base.F64_add(v28, base.F64_add(base.F64_sub(base.F64_sub(v24, v28), v27), base.F64_sub(base.F64_mul(v25, base.F64_add(base.F64_mul(v25, base.F64_add(base.F64_mul(v25, base.F64_add(base.F64_mul(v25, float64(2.480158728947673e-05)), float64(-0.001388888888887411))), float64(0.0416666666666666))), base.F64_mul(base.F64_mul(v40, v40), base.F64_add(base.F64_mul(v25, base.F64_add(base.F64_mul(v25, float64(-1.1359647557788195e-11)), float64(2.087572321298175e-09))), float64(-2.7557314351390663e-07))))), base.F64_mul(l0, float64(0)))))
		}
	}
	m.G0 = v7 + int32(16)
	return v212
}
func F_crc16(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v35 int32
	_ = v35
	var v43 int32
	_ = v43
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v76 int32
	_ = v76
	var v81 int32
	_ = v81
	if int32(1) <= l1 {
		v10 = int32(1)
		if l1 != v10 {
			v17 = int32(0)
			v19 = l0
			v20 = v17
			v23 = v17
			for {
				v25 = int32(65280)
				v27 = int32(8)
				v29 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19))))
				v31 = int32(1)
				v35 = int32(*(*uint16)(unsafe.Add(mBase, uint32((int32(base.Ui32(v20&v25)>>(uint(v27)%32))^v29)<<(uint(v31)%32))+uint32(_consts[281]))))
				v43 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19)+1)))
				v49 = int32(*(*uint16)(unsafe.Add(mBase, uint32((int32(base.Ui32((v35^v20<<(uint(v27)%32))&v25)>>(uint(v27)%32))^v43)<<(uint(v31)%32))+uint32(_consts[281]))))
				v52 = v49 ^ v35<<(uint(v27)%32)
				v53 = int32(2)
				v54 = v19 + v53
				v56 = v23 + v53
				if v56 != l1&int32(2147483646) {
					v19 = v54
					v20 = v52
					v23 = v56
					continue
				} else {
					break
				}
				break
			}
			v58 = v54
			v59 = v52
		} else {
			v58 = l0
			v59 = int32(0)
		}
		if l1&v10 == int32(0) {
			v81 = v59
		} else {
			v68 = int32(8)
			v70 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v58))))
			v76 = int32(*(*uint16)(unsafe.Add(mBase, uint32((int32(base.Ui32(v59&int32(65280))>>(uint(v68)%32))^v70)<<(uint(int32(1))%32))+uint32(_consts[281]))))
			v81 = v76 ^ v59<<(uint(v68)%32)
		}
	} else {
		v81 = int32(0)
	}
	return v81 & int32(65535)
}
func F_crc64_init(m *base.Module) {
	F_crcspeed64little_init(m, int32(513), int32(_a555))
	return
}
func F_crcspeed64native_init(m *base.Module, l0 int32, l1 int32) {
	F_crcspeed64little_init(m, l0, l1)
	return
}
func F_createCachedResponseClient(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v1 int32
	_ = v1
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	v1 = l0
	v4 = F_createClient(m, int32(0))
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		*(*uint8)(unsafe.Add(mBase, uint32(v4)+224)) = uint8(v1)
		*(*int64)(unsafe.Add(mBase, uint32(v4))) = int64(-2)
		v12 = F_valkey_calloc(m, int32(40))
		mBase = m.M
		v13 = m.ExcPending
		if v13 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v4)+8)) = v12
			v15 = *(*int32)(unsafe.Add(mBase, uint32(v4)+204))
			*(*int32)(unsafe.Add(mBase, uint32(v4)+204)) = v15 | int32(268435456)
			return v4
		}
	}
}
func F_createDatabase(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
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
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v99 int64
	_ = v99
	v6 = *(*int32)(unsafe.Add(mBase, _consts[110]))
	v8 = F_valkey_malloc(m, int32(64))
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	if v6 != 0 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v15 = int32(14)
	goto L5
L4:
	;
	v15 = int32(0)
	goto L5
L5:
	;
	if v6 != 0 {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	v18 = int32(3)
	goto L8
L7:
	;
	v18 = int32(1)
	goto L8
L8:
	;
	v19 = F_kvstoreCreate(m, int32(_a2214), v15, v18)
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		goto L1
	} else {
		goto L9
	}
L9:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8))) = v19
	v23 = F_kvstoreCreate(m, int32(_a2215), v15, v18)
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		goto L1
	} else {
		goto L10
	}
L10:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = v23
	v27 = F_kvstoreCreate(m, int32(_a2215), v15, v18)
	mBase = m.M
	v28 = m.ExcPending
	if v28 != 0 {
		goto L1
	} else {
		goto L11
	}
L11:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8)+8)) = v27
	v30 = int32(0)
	v33 = m.G0
	v35 = v33 - int32(16)
	m.G0 = v35
	v39 = *(*int32)(unsafe.Add(mBase, _consts[110]))
	if v39 == v30 {
		v73 = v30
		goto L14
	} else {
		goto L15
	}
L12:
	;
	v83 = F_dictCreate(m, int32(_a2216))
	mBase = m.M
	v84 = m.ExcPending
	if v84 != 0 {
		goto L1
	} else {
		goto L26
	}
L13:
	;
	if v73 == int32(0) {
		goto L12
	} else {
		goto L24
	}
L14:
	;
	m.G0 = v35 + int32(16)
	goto L13
L15:
	;
	v43 = *(*int32)(unsafe.Add(mBase, _consts[111]))
	if v43 == int32(0) {
		v73 = v30
		goto L14
	} else {
		goto L16
	}
L16:
	;
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v43)+uint32(_consts[171])))
	v48 = v35 + int32(8)
	F_listRewind(m, v46, v48)
	mBase = m.M
	v50 = int32(0)
	v53 = F_listNext(m, v48)
	mBase = m.M
	if v53 == v50 {
		v73 = v50
		goto L14
	} else {
		goto L17
	}
L17:
	;
	v58 = v53
	goto L18
L18:
	;
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v58)+8))
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v59)))
	if v60 != int32(1) {
		goto L20
	} else {
		goto L21
	}
L19:
	;
	v73 = v50
	goto L14
L20:
	;
	v71 = F_listNext(m, v35+int32(8))
	mBase = m.M
	if v71 != 0 {
		v58 = v71
		goto L18
	} else {
		goto L23
	}
L21:
	;
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v59)+156))
	if base.Ui32(v63+int32(-18)) <= base.Ui32(int32(2)) {
		goto L20
	} else {
		goto L22
	}
L22:
	;
	v73 = int32(1)
	goto L14
L23:
	;
	goto L19
L24:
	;
	F_clusterMarkImportingSlotsInDb(m, v8)
	mBase = m.M
	v81 = m.ExcPending
	if v81 != 0 {
		goto L1
	} else {
		goto L25
	}
L25:
	;
	goto L12
L26:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8)+12)) = v83
	v87 = F_dictCreate(m, int32(_a2217))
	mBase = m.M
	v88 = m.ExcPending
	if v88 != 0 {
		goto L1
	} else {
		goto L27
	}
L27:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = v87
	v91 = F_dictCreate(m, int32(_a2217))
	mBase = m.M
	v92 = m.ExcPending
	if v92 != 0 {
		goto L1
	} else {
		goto L28
	}
L28:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8)+20)) = v91
	v95 = F_dictCreate(m, int32(_a2216))
	mBase = m.M
	v96 = m.ExcPending
	if v96 != 0 {
		goto L1
	} else {
		goto L29
	}
L29:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8)+28)) = l0
	*(*int32)(unsafe.Add(mBase, uint32(v8)+24)) = v95
	v99 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v8)+32)) = v99
	*(*int64)(unsafe.Add(mBase, uint32(v8+int32(56)))) = v99
	*(*int64)(unsafe.Add(mBase, uint32(v8+int32(48)))) = v99
	*(*int64)(unsafe.Add(mBase, uint32(v8+int32(40)))) = v99
	goto L30
L30:
	;
	return v8
}
func F_createDatabaseIfNeeded(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	v5 = *(*int32)(unsafe.Add(mBase, _consts[173]))
	v7 = l0 << (uint(int32(2)) % 32)
	v9 = *(*int32)(unsafe.Add(mBase, uint32(v5+v7)))
	if v9 != 0 {
		v18 = v9
		return v18
	} else {
		v10 = F_createDatabase(m, l0)
		mBase = m.M
		v13 = m.ExcPending
		if v13 != 0 {
			return int32(0)
		} else {
			v15 = *(*int32)(unsafe.Add(mBase, _consts[173]))
			*(*int32)(unsafe.Add(mBase, uint32(v15+v7))) = v10
			v18 = v10
			return v18
		}
	}
}
func F_createDumpPayload(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
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
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v46 int32
	_ = v46
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v66 int64
	_ = v66
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v80 int32
	_ = v80
	var v86 int32
	_ = v86
	var v92 int32
	_ = v92
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	v11 = F_sdsempty(m)
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		return
	} else {
		v15 = F___memcpy(m, l0, int32(_a201), int32(80))
		mBase = m.M
		*(*int64)(unsafe.Add(mBase, uint32(v15)+56)) = int64(0)
		*(*int32)(unsafe.Add(mBase, uint32(v15)+48)) = v11
		v20 = F_rdbGetObjectType(m, l1, int32(80))
		mBase = m.M
		v21 = m.ExcPending
		if v21 != 0 {
			return
		} else {
			if v20 <= int32(-1) {
				F__serverAssert(m, int32(_a202), int32(_a203), int32(131))
				mBase = m.M
				v80 = m.ExcPending
				if v80 != 0 {
					return
				} else {
					F_abort(m)
					mBase = m.M
					base.Wasm_trap_unreachable()
					for {
					}
				}
			} else {
				v25 = v20 & int32(255)
				v26 = F_rdbSaveType(m, l0, v25)
				mBase = m.M
				v27 = m.ExcPending
				if v27 != 0 {
					return
				} else {
					if v26 == int32(0) {
						F__serverAssert(m, int32(_a204), int32(_a203), int32(132))
						mBase = m.M
						v86 = m.ExcPending
						if v86 != 0 {
							return
						} else {
							F_abort(m)
							mBase = m.M
							base.Wasm_trap_unreachable()
							for {
							}
						}
					} else {
						v30 = F_rdbSaveObject(m, l0, l1, l2, l3, v25)
						mBase = m.M
						v31 = m.ExcPending
						if v31 != 0 {
							return
						} else {
							if v30 == int32(0) {
								F__serverAssert(m, int32(_a205), int32(_a203), int32(133))
								mBase = m.M
								v92 = m.ExcPending
								if v92 != 0 {
									return
								} else {
									F_abort(m)
									mBase = m.M
									base.Wasm_trap_unreachable()
									for {
									}
								}
							} else {
								v34 = int32(80)
								*(*uint16)(unsafe.Add(mBase, uint32(v9)+14)) = uint16(v34)
								v36 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
								v40 = F_sdscatlen(m, v36, v9+int32(14), int32(2))
								mBase = m.M
								v41 = m.ExcPending
								if v41 != 0 {
									return
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(l0)+48)) = v40
									v46 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v40+int32(-1)))))
									switch v46 & int32(7) {
									case 0:
										v63 = int32(base.Ui32(v46) >> (uint(int32(3)) % 32))
									case 1:
										v53 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v40+int32(-3)))))
										v63 = v53
									case 2:
										v56 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v40+int32(-5)))))
										v63 = v56
									case 3:
										v59 = *(*int32)(unsafe.Add(mBase, uint32(v40+int32(-9))))
										v63 = v59
									case 4:
										v62 = *(*int32)(unsafe.Add(mBase, uint32(v40+int32(-17))))
										v63 = v62
									default:
										v63 = int32(0)
									}
									v66 = F_crc64(m, int64(0), v40, base.I64_extend_i32_u(v63))
									mBase = m.M
									*(*int64)(unsafe.Add(mBase, uint32(v9))) = v66
									v68 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
									v70 = F_sdscatlen(m, v68, v9, int32(8))
									mBase = m.M
									v71 = m.ExcPending
									if v71 != 0 {
										return
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(l0)+48)) = v70
										m.G0 = v9 + int32(16)
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
func F_createEmbeddedStringObjectWithKeyAndExpire(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int64) int32 {
	mBase := m.M
	_ = mBase
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v25 int32
	_ = v25
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v58 int32
	_ = v58
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v94 int32
	_ = v94
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v111 int32
	_ = v111
	var v114 int32
	_ = v114
	var v117 int32
	_ = v117
	var v121 int32
	_ = v121
	var v126 int32
	_ = v126
	var v128 int32
	_ = v128
	var v136 int32
	_ = v136
	var v144 int32
	_ = v144
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v159 int32
	_ = v159
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v178 int32
	_ = v178
	var v184 int32
	_ = v184
	v14 = m.G0
	v16 = v14 - int32(16)
	m.G0 = v16
	if l2 != 0 {
		v25 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2+int32(-1)))))
		switch v25 & int32(7) {
		case 0:
			v42 = int32(base.Ui32(v25) >> (uint(int32(3)) % 32))
		case 1:
			v32 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2+int32(-3)))))
			v42 = v32
		case 2:
			v35 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l2+int32(-5)))))
			v42 = v35
		case 3:
			v38 = *(*int32)(unsafe.Add(mBase, uint32(l2+int32(-9))))
			v42 = v38
		case 4:
			v41 = *(*int32)(unsafe.Add(mBase, uint32(l2+int32(-17))))
			v42 = v41
		default:
			v42 = int32(0)
		}
		if base.Ui32(int32(32)) <= base.Ui32(v42) {
			if base.Ui32(int32(253)) <= base.Ui32(v42) {
				if base.Ui32(v42) < base.Ui32(int32(65531)) {
					v53 = int32(2)
				} else {
					v53 = int32(3)
				}
				v54 = v53
			} else {
				v54 = int32(1)
			}
		} else {
			v54 = int32(0)
		}
		v58 = v54 & int32(7)
		if base.Ui32(int32(4)) < base.Ui32(v58) {
			v66 = int32(0)
		} else {
			v65 = *(*int32)(unsafe.Add(mBase, uint32(v58<<(uint(int32(2))%32))+uint32(_consts[309])))
			v66 = v65
		}
		v71 = int32(2)
		v72 = v54
		v73 = v42
		v74 = v42 + v66 + int32(1)
	} else {
		v18 = int32(0)
		v71 = v18
		v72 = v18
		v73 = v18
		v74 = v18
	}
	v87 = *(*int32)(unsafe.Add(mBase, _consts[308]))
	v89 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v16)+12)) = v89
	if l2 != 0 {
		v94 = v74 + int32(1)
	} else {
		v94 = v89
	}
	v98 = base.B2i32(l3 != int64(-1))
	if l3 != int64(-1) {
		v99 = int32(16)
	} else {
		v99 = int32(8)
	}
	v103 = l1 + v87 + int32(1)
	v104 = int32(4)
	if base.Ui32(v104) < base.Ui32(v103) {
		v107 = v103
	} else {
		v107 = v104
	}
	v108 = v94 + v99 + v107
	v111 = F_zmalloc_usable(m, v108, v16+int32(12))
	mBase = m.M
	v114 = m.ExcPending
	if v114 != 0 {
		return int32(0)
	} else {
		v117 = v71 | v98 | int32(12)
		*(*int32)(unsafe.Add(mBase, uint32(v111)+4)) = v117
		*(*int32)(unsafe.Add(mBase, uint32(v111))) = int32(128)
		if l3 != int64(-1) {
			v128 = v117
		} else {
			v121 = *(*int32)(unsafe.Add(mBase, uint32(v16)+12))
			if base.Ui32(v121) < base.Ui32(v108+int32(8)) {
				v128 = v117
			} else {
				v126 = v71 | int32(13)
				*(*int32)(unsafe.Add(mBase, uint32(v111)+4)) = v126
				v128 = v126
			}
		}
		if v128&int32(1) != 0 {
			*(*int64)(unsafe.Add(mBase, uint32(v111)+8)) = l3
			v136 = v111 + int32(16)
		} else {
			v136 = v111 + int32(8)
		}
		if v128&int32(2) == int32(0) {
			v159 = v136
			if base.Ui32(int32(256)) <= base.Ui32(l1) {
				F__serverAssert(m, int32(_a1759), int32(_a1758), int32(215))
				mBase = m.M
				v178 = m.ExcPending
				if v178 != 0 {
					return int32(0)
				} else {
					F_abort(m)
					mBase = m.M
					base.Wasm_trap_unreachable()
					for {
					}
				}
			} else {
				v163 = *(*int32)(unsafe.Add(mBase, uint32(v16)+12))
				v164 = v111 - v159 + v163
				if base.Ui32(int32(256)) <= base.Ui32(v164) {
					F__serverAssert(m, int32(_a1760), int32(_a1758), int32(216))
					mBase = m.M
					v184 = m.ExcPending
					if v184 != 0 {
						return int32(0)
					} else {
						F_abort(m)
						mBase = m.M
						base.Wasm_trap_unreachable()
						for {
						}
					}
				} else {
					v168 = F_sdswrite(m, v159, v164, int32(1), l0, l1)
					mBase = m.M
					v169 = m.ExcPending
					if v169 != 0 {
						return int32(0)
					} else {
						m.G0 = v16 + int32(16)
						return v111
					}
				}
			}
		} else {
			v144 = v72 & int32(7)
			if base.Ui32(int32(4)) < base.Ui32(v144) {
				v152 = int32(0)
			} else {
				v151 = *(*int32)(unsafe.Add(mBase, uint32(v144<<(uint(int32(2))%32))+uint32(_consts[309])))
				v152 = v151
			}
			*(*uint8)(unsafe.Add(mBase, uint32(v136))) = uint8(v152)
			v155 = v136 + int32(1)
			v156 = F_sdswrite(m, v155, v74, v72, l2, v73)
			mBase = m.M
			v157 = m.ExcPending
			if v157 != 0 {
				return int32(0)
			} else {
				v159 = v155 + v74
				if base.Ui32(int32(256)) <= base.Ui32(l1) {
					F__serverAssert(m, int32(_a1759), int32(_a1758), int32(215))
					mBase = m.M
					v178 = m.ExcPending
					if v178 != 0 {
						return int32(0)
					} else {
						F_abort(m)
						mBase = m.M
						base.Wasm_trap_unreachable()
						for {
						}
					}
				} else {
					v163 = *(*int32)(unsafe.Add(mBase, uint32(v16)+12))
					v164 = v111 - v159 + v163
					if base.Ui32(int32(256)) <= base.Ui32(v164) {
						F__serverAssert(m, int32(_a1760), int32(_a1758), int32(216))
						mBase = m.M
						v184 = m.ExcPending
						if v184 != 0 {
							return int32(0)
						} else {
							F_abort(m)
							mBase = m.M
							base.Wasm_trap_unreachable()
							for {
							}
						}
					} else {
						v168 = F_sdswrite(m, v159, v164, int32(1), l0, l1)
						mBase = m.M
						v169 = m.ExcPending
						if v169 != 0 {
							return int32(0)
						} else {
							m.G0 = v16 + int32(16)
							return v111
						}
					}
				}
			}
		}
	}
}
func F_createNilObject(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v37 int32
	_ = v37
	v7 = m.G4
	v8 = *(*int32)(unsafe.Add(mBase, uint32(v7)+4))
	v9 = m.T0[v8].(func(*base.Module, int32, int32) int32)(m, int32(1), int32(48))
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		return int32(0)
	} else {
		if v9 == int32(0) {
			return v9
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v9))) = int32(4)
			v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
			if v17 == int32(0) {
				return v9
			} else {
				v20 = *(*int32)(unsafe.Add(mBase, uint32(v17)+20))
				v21 = *(*int32)(unsafe.Add(mBase, uint32(v20)))
				if base.Ui32(v21+int32(-9)) < base.Ui32(int32(4)) {
					v28 = *(*int32)(unsafe.Add(mBase, uint32(v20)+40))
					v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
					*(*int32)(unsafe.Add(mBase, uint32(v28+v29<<(uint(int32(2))%32)))) = v9
					return v9
				} else {
					if v21 != int32(2) {
						v37 = m.G3
						m.Env.X__assert_fail(m, v37+int32(_a2584), v37+int32(_a2585), int32(269), v37+int32(_a2586))
						mBase = m.M
						base.Wasm_trap_unreachable()
						for {
						}
					} else {
						v28 = *(*int32)(unsafe.Add(mBase, uint32(v20)+40))
						v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
						*(*int32)(unsafe.Add(mBase, uint32(v28+v29<<(uint(int32(2))%32)))) = v9
						return v9
					}
				}
			}
		}
	}
}
func F_createRawStringObject(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	v8 = F_sdsnewlen(m, l0, l1)
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return int32(0)
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(v6)+12)) = int32(0)
		v14 = int32(12)
		v17 = F_zmalloc_usable(m, v14, v6+v14)
		mBase = m.M
		v18 = m.ExcPending
		if v18 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v17)+8)) = v8
			*(*int64)(unsafe.Add(mBase, uint32(v17))) = int64(34359738368)
			m.G0 = v6 + int32(16)
			return v17
		}
	}
}
func F_createSetListpackObject(m *base.Module) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v23 int32
	_ = v23
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	v9 = F_lpNew(m, int32(0))
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		return int32(0)
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(v6)+12)) = int32(0)
		v15 = int32(12)
		v18 = F_zmalloc_usable(m, v15, v6+v15)
		mBase = m.M
		v19 = m.ExcPending
		if v19 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v18)+8)) = v9
			*(*int64)(unsafe.Add(mBase, uint32(v18))) = int64(34359738370)
			v23 = *(*int32)(unsafe.Add(mBase, uint32(v18)))
			*(*int32)(unsafe.Add(mBase, uint32(v18))) = v23&int32(-241) | int32(176)
			m.G0 = v6 + int32(16)
			return v18
		}
	}
}
func F_createSharedObjectsForCompat(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
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
	var v44 int32
	_ = v44
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
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	v5 = m.G0
	v7 = v5 - int32(80)
	m.G0 = v7
	v9 = F_sdsempty(m)
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		return
	} else {
		if l0 != 0 {
			v13 = int32(_a566)
		} else {
			v13 = int32(_a256)
		}
		*(*int32)(unsafe.Add(mBase, uint32(v7)+64)) = v13
		v16 = l0 << (uint(int32(2)) % 32)
		v23 = F_sdscatfmt(m, v9, int32(_a2192), v7+int32(64))
		mBase = m.M
		v24 = m.ExcPending
		if v24 != 0 {
			return
		} else {
			v25 = F_createObject(m, int32(0), v23)
			mBase = m.M
			v26 = m.ExcPending
			if v26 != 0 {
				return
			} else {
				v27 = F_makeObjectShared(m, v25)
				mBase = m.M
				v28 = m.ExcPending
				if v28 != 0 {
					return
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v16)+uint32(_consts[736]))) = v27
					v30 = F_sdsempty(m)
					mBase = m.M
					v31 = m.ExcPending
					if v31 != 0 {
						return
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v7)+48)) = v13
						v39 = F_sdscatfmt(m, v30, int32(_a2193), v7+int32(48))
						mBase = m.M
						v40 = m.ExcPending
						if v40 != 0 {
							return
						} else {
							v41 = F_createObject(m, int32(0), v39)
							mBase = m.M
							v42 = m.ExcPending
							if v42 != 0 {
								return
							} else {
								v43 = F_makeObjectShared(m, v41)
								mBase = m.M
								v44 = m.ExcPending
								if v44 != 0 {
									return
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v16)+uint32(_consts[737]))) = v43
									v46 = F_sdsempty(m)
									mBase = m.M
									v47 = m.ExcPending
									if v47 != 0 {
										return
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v7)+32)) = v13
										v55 = F_sdscatfmt(m, v46, int32(_a2194), v7+int32(32))
										mBase = m.M
										v56 = m.ExcPending
										if v56 != 0 {
											return
										} else {
											v57 = F_createObject(m, int32(0), v55)
											mBase = m.M
											v58 = m.ExcPending
											if v58 != 0 {
												return
											} else {
												v59 = F_makeObjectShared(m, v57)
												mBase = m.M
												v60 = m.ExcPending
												if v60 != 0 {
													return
												} else {
													*(*int32)(unsafe.Add(mBase, uint32(v16)+uint32(_consts[738]))) = v59
													v62 = F_sdsempty(m)
													mBase = m.M
													v63 = m.ExcPending
													if v63 != 0 {
														return
													} else {
														*(*int32)(unsafe.Add(mBase, uint32(v7)+16)) = v13
														v71 = F_sdscatfmt(m, v62, int32(_a2195), v7+int32(16))
														mBase = m.M
														v72 = m.ExcPending
														if v72 != 0 {
															return
														} else {
															v73 = F_createObject(m, int32(0), v71)
															mBase = m.M
															v74 = m.ExcPending
															if v74 != 0 {
																return
															} else {
																v75 = F_makeObjectShared(m, v73)
																mBase = m.M
																v76 = m.ExcPending
																if v76 != 0 {
																	return
																} else {
																	*(*int32)(unsafe.Add(mBase, uint32(v16)+uint32(_consts[739]))) = v75
																	v78 = F_sdsempty(m)
																	mBase = m.M
																	v79 = m.ExcPending
																	if v79 != 0 {
																		return
																	} else {
																		*(*int32)(unsafe.Add(mBase, uint32(v7)+4)) = v13
																		*(*int32)(unsafe.Add(mBase, uint32(v7))) = v13
																		v86 = F_sdscatfmt(m, v78, int32(_a2196), v7)
																		mBase = m.M
																		v87 = m.ExcPending
																		if v87 != 0 {
																			return
																		} else {
																			v88 = F_createObject(m, int32(0), v86)
																			mBase = m.M
																			v89 = m.ExcPending
																			if v89 != 0 {
																				return
																			} else {
																				v90 = F_makeObjectShared(m, v88)
																				mBase = m.M
																				v91 = m.ExcPending
																				if v91 != 0 {
																					return
																				} else {
																					*(*int32)(unsafe.Add(mBase, uint32(v16)+uint32(_consts[740]))) = v90
																					m.G0 = v7 + int32(80)
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
func F_createUnembeddedObjectWithKeyAndExpire(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int64) int32 {
	mBase := m.M
	_ = mBase
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v23 int32
	_ = v23
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v89 int32
	_ = v89
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v114 int32
	_ = v114
	var v117 int32
	_ = v117
	var v120 int32
	_ = v120
	var v127 int32
	_ = v127
	var v130 int32
	_ = v130
	var v132 int32
	_ = v132
	var v142 int32
	_ = v142
	var v146 int32
	_ = v146
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v160 int32
	_ = v160
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	v12 = m.G0
	v14 = v12 - int32(16)
	m.G0 = v14
	if l2 == int32(0) {
		v44 = base.B2i32(l3 != int64(-1))
		if l2 != 0 {
			v50 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2+int32(-1)))))
			v52 = v50
			v53 = v44
			switch v52 & int32(7) {
			case 0:
				v73 = int32(base.Ui32(v52&int32(248)) >> (uint(int32(3)) % 32))
			case 1:
				v63 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2+int32(-3)))))
				v73 = v63
			case 2:
				v66 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l2+int32(-5)))))
				v73 = v66
			case 3:
				v69 = *(*int32)(unsafe.Add(mBase, uint32(l2+int32(-9))))
				v73 = v69
			case 4:
				v72 = *(*int32)(unsafe.Add(mBase, uint32(l2+int32(-17))))
				v73 = v72
			default:
				v73 = int32(0)
			}
			if base.Ui32(int32(32)) <= base.Ui32(v73) {
				if base.Ui32(int32(253)) <= base.Ui32(v73) {
					if base.Ui32(v73) < base.Ui32(int32(65531)) {
						v84 = int32(2)
					} else {
						v84 = int32(3)
					}
					v85 = v84
				} else {
					v85 = int32(1)
				}
			} else {
				v85 = int32(0)
			}
			v89 = v85 & int32(7)
			if base.Ui32(int32(4)) < base.Ui32(v89) {
				v97 = int32(0)
			} else {
				v96 = *(*int32)(unsafe.Add(mBase, uint32(v89<<(uint(int32(2))%32))+uint32(_consts[309])))
				v97 = v96
			}
			v101 = v73
			v103 = v53
			v104 = v85
			v105 = v73 + v97 + int32(1)
		} else {
			v45 = int32(0)
			v101 = v45
			v103 = v44
			v104 = v45
			v105 = v45
		}
	} else {
		if l3 != int64(-1) {
			v44 = base.B2i32(l3 != int64(-1))
			if l2 != 0 {
				v50 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2+int32(-1)))))
				v52 = v50
				v53 = v44
				switch v52 & int32(7) {
				case 0:
					v73 = int32(base.Ui32(v52&int32(248)) >> (uint(int32(3)) % 32))
				case 1:
					v63 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2+int32(-3)))))
					v73 = v63
				case 2:
					v66 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l2+int32(-5)))))
					v73 = v66
				case 3:
					v69 = *(*int32)(unsafe.Add(mBase, uint32(l2+int32(-9))))
					v73 = v69
				case 4:
					v72 = *(*int32)(unsafe.Add(mBase, uint32(l2+int32(-17))))
					v73 = v72
				default:
					v73 = int32(0)
				}
				if base.Ui32(int32(32)) <= base.Ui32(v73) {
					if base.Ui32(int32(253)) <= base.Ui32(v73) {
						if base.Ui32(v73) < base.Ui32(int32(65531)) {
							v84 = int32(2)
						} else {
							v84 = int32(3)
						}
						v85 = v84
					} else {
						v85 = int32(1)
					}
				} else {
					v85 = int32(0)
				}
				v89 = v85 & int32(7)
				if base.Ui32(int32(4)) < base.Ui32(v89) {
					v97 = int32(0)
				} else {
					v96 = *(*int32)(unsafe.Add(mBase, uint32(v89<<(uint(int32(2))%32))+uint32(_consts[309])))
					v97 = v96
				}
				v101 = v73
				v103 = v53
				v104 = v85
				v105 = v73 + v97 + int32(1)
			} else {
				v45 = int32(0)
				v101 = v45
				v103 = v44
				v104 = v45
				v105 = v45
			}
		} else {
			v23 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2+int32(-1)))))
			switch v23 & int32(7) {
			case 0:
				v40 = int32(base.Ui32(v23) >> (uint(int32(3)) % 32))
			case 1:
				v30 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2+int32(-3)))))
				v40 = v30
			case 2:
				v33 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l2+int32(-5)))))
				v40 = v33
			case 3:
				v36 = *(*int32)(unsafe.Add(mBase, uint32(l2+int32(-9))))
				v40 = v36
			case 4:
				v39 = *(*int32)(unsafe.Add(mBase, uint32(l2+int32(-17))))
				v40 = v39
			default:
				v40 = int32(0)
			}
			v52 = v23
			v53 = base.B2i32(base.Ui32(int32(127)) < base.Ui32(v40))
			switch v52 & int32(7) {
			case 0:
				v73 = int32(base.Ui32(v52&int32(248)) >> (uint(int32(3)) % 32))
			case 1:
				v63 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2+int32(-3)))))
				v73 = v63
			case 2:
				v66 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l2+int32(-5)))))
				v73 = v66
			case 3:
				v69 = *(*int32)(unsafe.Add(mBase, uint32(l2+int32(-9))))
				v73 = v69
			case 4:
				v72 = *(*int32)(unsafe.Add(mBase, uint32(l2+int32(-17))))
				v73 = v72
			default:
				v73 = int32(0)
			}
			if base.Ui32(int32(32)) <= base.Ui32(v73) {
				if base.Ui32(int32(253)) <= base.Ui32(v73) {
					if base.Ui32(v73) < base.Ui32(int32(65531)) {
						v84 = int32(2)
					} else {
						v84 = int32(3)
					}
					v85 = v84
				} else {
					v85 = int32(1)
				}
			} else {
				v85 = int32(0)
			}
			v89 = v85 & int32(7)
			if base.Ui32(int32(4)) < base.Ui32(v89) {
				v97 = int32(0)
			} else {
				v96 = *(*int32)(unsafe.Add(mBase, uint32(v89<<(uint(int32(2))%32))+uint32(_consts[309])))
				v97 = v96
			}
			v101 = v73
			v103 = v53
			v104 = v85
			v105 = v73 + v97 + int32(1)
		}
	}
	*(*int32)(unsafe.Add(mBase, uint32(v14)+12)) = int32(0)
	if v103 != 0 {
		v110 = int32(20)
	} else {
		v110 = int32(12)
	}
	v111 = v105 + v110
	if l2 != 0 {
		v114 = v111 + int32(1)
	} else {
		v114 = v110
	}
	v117 = F_zmalloc_usable(m, v114, v14+int32(12))
	mBase = m.M
	v120 = m.ExcPending
	if v120 != 0 {
		return int32(0)
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(v117)+8)) = l1
		*(*int32)(unsafe.Add(mBase, uint32(v117))) = l0 & int32(15)
		if l2 == int32(0) {
			v142 = base.B2i32(l2 != int32(0))<<(uint(int32(1))%32) | v103 | int32(8)
			*(*int32)(unsafe.Add(mBase, uint32(v117)+4)) = v142
			if v103 != 0 {
				v146 = v142
				*(*int64)(unsafe.Add(mBase, uint32(v117)+12)) = l3
				v151 = v146
				v152 = v117 + int32(20)
			} else {
				v151 = v142
				v152 = v117 + int32(12)
			}
		} else {
			if v103 != 0 {
				v142 = base.B2i32(l2 != int32(0))<<(uint(int32(1))%32) | v103 | int32(8)
				*(*int32)(unsafe.Add(mBase, uint32(v117)+4)) = v142
				if v103 != 0 {
					v146 = v142
					*(*int64)(unsafe.Add(mBase, uint32(v117)+12)) = l3
					v151 = v146
					v152 = v117 + int32(20)
				} else {
					v151 = v142
					v152 = v117 + int32(12)
				}
			} else {
				v127 = *(*int32)(unsafe.Add(mBase, uint32(v14)+12))
				v130 = base.B2i32(base.Ui32(v111+int32(9)) <= base.Ui32(v127))
				v132 = v130 | int32(10)
				*(*int32)(unsafe.Add(mBase, uint32(v117)+4)) = v132
				if base.Ui32(v111+int32(9)) <= base.Ui32(v127) {
					v146 = v132
					*(*int64)(unsafe.Add(mBase, uint32(v117)+12)) = l3
					v151 = v146
					v152 = v117 + int32(20)
				} else {
					v151 = v132
					v152 = v117 + int32(12)
				}
			}
		}
		if v151&int32(2) == int32(0) {
			m.G0 = v14 + int32(16)
			return v117
		} else {
			v160 = v104 & int32(7)
			if base.Ui32(int32(4)) < base.Ui32(v160) {
				v168 = int32(0)
			} else {
				v167 = *(*int32)(unsafe.Add(mBase, uint32(v160<<(uint(int32(2))%32))+uint32(_consts[309])))
				v168 = v167
			}
			*(*uint8)(unsafe.Add(mBase, uint32(v152))) = uint8(v168)
			v172 = F_sdswrite(m, v152+int32(1), v105, v104, l2, v101)
			mBase = m.M
			v173 = m.ExcPending
			if v173 != 0 {
				return int32(0)
			} else {
				m.G0 = v14 + int32(16)
				return v117
			}
		}
	}
}
func F_cumulativeKeyCountAdd(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v13 int64
	_ = v13
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v20 int64
	_ = v20
	var v23 int64
	_ = v23
	var v27 int32
	_ = v27
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
	var v57 int32
	_ = v57
	var v61 int32
	_ = v61
	var v70 int32
	_ = v70
	var v71 int64
	_ = v71
	var v78 int32
	_ = v78
	var v92 int32
	_ = v92
	var v98 int32
	_ = v98
	if l1 < int32(0) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	F__serverAssert(m, int32(_a682), int32(_a680), int32(182))
	mBase = m.M
	v98 = m.ExcPending
	if v98 != 0 {
		goto L6
	} else {
		goto L27
	}
L2:
	;
	F__serverAssert(m, int32(_a683), int32(_a680), int32(150))
	mBase = m.M
	v92 = m.ExcPending
	if v92 != 0 {
		goto L6
	} else {
		goto L26
	}
L3:
	;
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v11 <= l1 {
		goto L2
	} else {
		goto L4
	}
L4:
	;
	v13 = base.I64_extend_i32_s(l2)
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
	v16 = F_hashtableFind(m, v14, l1, int32(0))
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	v23 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
	*(*int64)(unsafe.Add(mBase, uint32(l0)+40)) = v23 + v13
	v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v27+l1<<(uint(int32(2))%32))))
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v31)+20))
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v31)+16))
	v34 = v32 + v33
	goto L9
L6:
	;
	return
L7:
	;
	if v16 == int32(0) {
		goto L5
	} else {
		goto L8
	}
L8:
	;
	v20 = *(*int64)(unsafe.Add(mBase, uint32(l0)+72))
	*(*int64)(unsafe.Add(mBase, uint32(l0)+72)) = v20 + v13
	return
L9:
	;
	if int32(-1) < l2 {
		goto L12
	} else {
		goto L13
	}
L10:
	;
	v48 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v48 == int32(1) {
		goto L17
	} else {
		goto L18
	}
L11:
	;
	v44 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v44 + v43
	goto L10
L12:
	;
	v39 = int32(1)
	if l2 < v39 {
		goto L10
	} else {
		goto L15
	}
L13:
	;
	if v34 == int32(0) {
		v43 = int32(-1)
		goto L11
	} else {
		goto L14
	}
L14:
	;
	goto L12
L15:
	;
	if v34 != l2 {
		goto L10
	} else {
		goto L16
	}
L16:
	;
	v43 = v39
	goto L11
L17:
	;
	return
L18:
	;
	if v48 <= l1 {
		goto L17
	} else {
		goto L19
	}
L19:
	;
	v57 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	v61 = l1 + int32(1)
	goto L20
L20:
	;
	v70 = v57 + v61<<(uint(int32(3))%32)
	v71 = *(*int64)(unsafe.Add(mBase, uint32(v70)))
	if int32(-1) < l2 {
		goto L22
	} else {
		goto L23
	}
L21:
	;
	goto L17
L22:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v70))) = v71 + v13
	v78 = v61&(int32(0)-v61) + v61
	if v78 <= v48 {
		v61 = v78
		goto L20
	} else {
		goto L25
	}
L23:
	;
	if base.Ui64(v71) < base.Ui64(base.I64_extend_i32_u(int32(0)-l2)) {
		goto L1
	} else {
		goto L24
	}
L24:
	;
	goto L22
L25:
	;
	goto L21
L26:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L27:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
}
