package p0

import (
	base "github.com/shibukawa/vkmem/internal/aot/vkaot/base"
	"unsafe"
)

func F___getitimer(m *base.Module, l0 int32, l1 int32, l2 float64) {
	mBase := m.M
	_ = mBase
	var v4 float64
	_ = v4
	var v8 int32
	_ = v8
	var v11 float64
	_ = v11
	var v12 float64
	_ = v12
	var v16 float64
	_ = v16
	var v18 float64
	_ = v18
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v29 float64
	_ = v29
	var v35 int64
	_ = v35
	var v37 int64
	_ = v37
	var v41 float64
	_ = v41
	var v43 float64
	_ = v43
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v54 float64
	_ = v54
	var v60 int64
	_ = v60
	var v62 int64
	_ = v62
	v4 = float64(0)
	v8 = l0 << (uint(int32(3)) % 32)
	v11 = *(*float64)(unsafe.Add(mBase, uint32(v8)+uint32(_c_F___getitimer[0])))
	v12 = base.F64_sub(v11, l2)
	if base.F64_gt(v12, v4) != 0 {
		v16 = v12
	} else {
		v16 = v4
	}
	v18 = base.F64_mul(v16, float64(1000))
	if base.F64_lt(base.F64_abs(v18), float64(2.147483648e+09)) == int32(0) {
		v26 = int32(-2147483648)
	} else {
		v24 = base.I32_trunc_f64_s(v18)
		v26 = v24
	}
	*(*int32)(unsafe.Add(mBase, uint32(l1)+24)) = v26
	v29 = base.F64_div(v16, float64(1000))
	if base.F64_lt(base.F64_abs(v29), float64(9.223372036854776e+18)) == int32(0) {
		v37 = int64(-9223372036854775807 - 1)
	} else {
		v35 = base.I64_trunc_f64_s(v29)
		v37 = v35
	}
	*(*int64)(unsafe.Add(mBase, uint32(l1)+16)) = v37
	v41 = *(*float64)(unsafe.Add(mBase, uint32(v8)+uint32(_c_F___getitimer[1])))
	v43 = base.F64_mul(v41, float64(1000))
	if base.F64_lt(base.F64_abs(v43), float64(2.147483648e+09)) == int32(0) {
		v51 = int32(-2147483648)
	} else {
		v49 = base.I32_trunc_f64_s(v43)
		v51 = v49
	}
	*(*int32)(unsafe.Add(mBase, uint32(l1)+8)) = v51
	v54 = base.F64_div(v41, float64(1000))
	if base.F64_lt(base.F64_abs(v54), float64(9.223372036854776e+18)) == int32(0) {
		v62 = int64(-9223372036854775807 - 1)
	} else {
		v60 = base.I64_trunc_f64_s(v54)
		v62 = v60
	}
	*(*int64)(unsafe.Add(mBase, uint32(l1))) = v62
	return
}
func F_genModulesInfoStringRenderModuleOptions(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
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
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
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
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v87 int32
	_ = v87
	var v93 int32
	_ = v93
	var v97 int32
	_ = v97
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v105 int32
	_ = v105
	var v109 int32
	_ = v109
	var v114 int32
	_ = v114
	var v116 int32
	_ = v116
	var v118 int32
	_ = v118
	var v120 int32
	_ = v120
	var v124 int32
	_ = v124
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	v5 = F_sdsnew(m, int32(_a_F_genModulesInfoStringRenderModuleOptions_0))
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	if v9&int32(1) == int32(0) {
		v18 = v5
		v19 = v9
		goto L3
	} else {
		goto L4
	}
L3:
	;
	if v19&int32(4) == int32(0) {
		v28 = v18
		v29 = v19
		goto L6
	} else {
		goto L7
	}
L4:
	;
	v15 = F_sdscat(m, v5, int32(_a_F_genModulesInfoStringRenderModuleOptions_1))
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		goto L1
	} else {
		goto L5
	}
L5:
	;
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v18 = v15
	v19 = v17
	goto L3
L6:
	;
	if v29&int32(2) == int32(0) {
		v38 = v28
		v39 = v29
		goto L9
	} else {
		goto L10
	}
L7:
	;
	v25 = F_sdscat(m, v18, int32(_a_F_genModulesInfoStringRenderModuleOptions_2))
	mBase = m.M
	v26 = m.ExcPending
	if v26 != 0 {
		goto L1
	} else {
		goto L8
	}
L8:
	;
	v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v28 = v25
	v29 = v27
	goto L6
L9:
	;
	if v39&int32(32) == int32(0) {
		v47 = v38
		goto L12
	} else {
		goto L13
	}
L10:
	;
	v35 = F_sdscat(m, v28, int32(_a_F_genModulesInfoStringRenderModuleOptions_3))
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
		goto L1
	} else {
		goto L11
	}
L11:
	;
	v37 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v38 = v35
	v39 = v37
	goto L9
L12:
	;
	v48 = int32(_a_F_genModulesInfoStringRenderModuleOptions_4)
	v54 = v47 + int32(-1)
	v55 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v54))))
	switch v55 & int32(7) {
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
		v72 = int32(0)
		goto L16
	}
L13:
	;
	v45 = F_sdscat(m, v38, int32(_a_F_genModulesInfoStringRenderModuleOptions_5))
	mBase = m.M
	v46 = m.ExcPending
	if v46 != 0 {
		goto L1
	} else {
		goto L14
	}
L14:
	;
	v47 = v45
	goto L12
L15:
	;
	v140 = F_sdscat(m, v47, int32(_a_F_genModulesInfoStringRenderModuleOptions_6))
	mBase = m.M
	v141 = m.ExcPending
	if v141 != 0 {
		goto L1
	} else {
		goto L42
	}
L16:
	;
	v75 = v47 + v72 + int32(-1)
	if base.Ui32(v75) < base.Ui32(v47) {
		v93 = v47
		goto L22
	} else {
		goto L23
	}
L17:
	;
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v47+int32(-17))))
	v72 = v71
	goto L16
L18:
	;
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v47+int32(-9))))
	v72 = v68
	goto L16
L19:
	;
	v65 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v47+int32(-5)))))
	v72 = v65
	goto L16
L20:
	;
	v62 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v47+int32(-3)))))
	v72 = v62
	goto L16
L21:
	;
	v72 = int32(base.Ui32(v55) >> (uint(int32(3)) % 32))
	goto L16
L22:
	;
	if base.Ui32(v75) <= base.Ui32(v93) {
		v109 = v75
		goto L28
	} else {
		goto L29
	}
L23:
	;
	v81 = v47
	goto L24
L24:
	;
	v82 = int32(*(*int8)(unsafe.Add(mBase, uint32(v81))))
	v83 = F_strchr(m, v48, v82)
	mBase = m.M
	if v83 == int32(0) {
		v93 = v81
		goto L22
	} else {
		goto L26
	}
L25:
	;
	v93 = v87
	goto L22
L26:
	;
	v87 = v81 + int32(1)
	if base.Ui32(v87) <= base.Ui32(v75) {
		v81 = v87
		goto L24
	} else {
		goto L27
	}
L27:
	;
	goto L25
L28:
	;
	v114 = v109 - v93 + int32(1)
	if v47 == v93 {
		goto L34
	} else {
		goto L35
	}
L29:
	;
	v97 = v75
	goto L30
L30:
	;
	v100 = int32(*(*int8)(unsafe.Add(mBase, uint32(v97))))
	v101 = F_strchr(m, v48, v100)
	mBase = m.M
	if v101 == int32(0) {
		v109 = v97
		goto L28
	} else {
		goto L32
	}
L31:
	;
	v109 = v93
	goto L28
L32:
	;
	v105 = v97 + int32(-1)
	if base.Ui32(v93) < base.Ui32(v105) {
		v97 = v105
		goto L30
	} else {
		goto L33
	}
L33:
	;
	goto L31
L34:
	;
	v118 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v47+v114))) = uint8(v118)
	v120 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v54))))
	switch v120 & int32(7) {
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
		goto L36
	}
L35:
	;
	v116 = F_memmove(m, v47, v93, v114)
	mBase = m.M
	goto L34
L36:
	;
	goto L15
L37:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v47+int32(-17)))) = base.I64_extend_i32_u(v114)
	goto L36
L38:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v47+int32(-9)))) = v114
	goto L15
L39:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v47+int32(-5)))) = uint16(v114)
	goto L15
L40:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v47+int32(-3)))) = uint8(v114)
	goto L15
L41:
	;
	v124 = v114 << (uint(int32(3)) % 32)
	*(*uint8)(unsafe.Add(mBase, uint32(v54))) = uint8(v124)
	goto L15
L42:
	;
	return v140
}
func F_generateSkyline(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v21 int32
	_ = v21
	var v26 int32
	_ = v26
	var v28 int64
	_ = v28
	var v32 int64
	_ = v32
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v44 int64
	_ = v44
	var v48 int64
	_ = v48
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v62 int64
	_ = v62
	var v66 int64
	_ = v66
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v82 int32
	_ = v82
	var v96 int32
	_ = v96
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v109 int32
	_ = v109
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v130 int32
	_ = v130
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v161 int32
	_ = v161
	var v166 int32
	_ = v166
	var v168 int32
	_ = v168
	var v187 int32
	_ = v187
	var v189 int32
	_ = v189
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v200 int32
	_ = v200
	var v204 int32
	_ = v204
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v222 int32
	_ = v222
	var v229 int32
	_ = v229
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
	var v241 int32
	_ = v241
	var v254 int32
	_ = v254
	var v273 int32
	_ = v273
	var v274 int32
	_ = v274
	var v284 int32
	_ = v284
	var v289 int32
	_ = v289
	var v291 int64
	_ = v291
	var v295 int64
	_ = v295
	var v301 int32
	_ = v301
	var v302 int32
	_ = v302
	var v305 int32
	_ = v305
	var v307 int64
	_ = v307
	var v311 int64
	_ = v311
	var v317 int32
	_ = v317
	var v321 int32
	_ = v321
	var v323 int32
	_ = v323
	var v325 int64
	_ = v325
	var v329 int64
	_ = v329
	var v334 int32
	_ = v334
	var v337 int32
	_ = v337
	var v339 int32
	_ = v339
	var v341 int32
	_ = v341
	var v345 int32
	_ = v345
	var v359 int32
	_ = v359
	var v362 int32
	_ = v362
	var v364 int32
	_ = v364
	var v365 int32
	_ = v365
	var v372 int32
	_ = v372
	var v376 int32
	_ = v376
	var v377 int32
	_ = v377
	var v393 int32
	_ = v393
	var v400 int32
	_ = v400
	var v401 int32
	_ = v401
	var v408 int32
	_ = v408
	var v409 int32
	_ = v409
	var v419 int32
	_ = v419
	var v420 int32
	_ = v420
	var v424 int32
	_ = v424
	var v429 int32
	_ = v429
	var v431 int32
	_ = v431
	var v450 int32
	_ = v450
	var v452 int32
	_ = v452
	var v454 int32
	_ = v454
	var v455 int32
	_ = v455
	var v463 int32
	_ = v463
	var v467 int32
	_ = v467
	var v480 int32
	_ = v480
	var v481 int32
	_ = v481
	var v485 int32
	_ = v485
	var v492 int32
	_ = v492
	var v498 int32
	_ = v498
	var v499 int32
	_ = v499
	var v504 int32
	_ = v504
	var v517 int32
	_ = v517
	var v536 int32
	_ = v536
	var v537 int32
	_ = v537
	var v541 int32
	_ = v541
	var v547 int32
	_ = v547
	var v552 int32
	_ = v552
	var v554 int64
	_ = v554
	var v558 int64
	_ = v558
	var v564 int32
	_ = v564
	var v565 int32
	_ = v565
	var v568 int32
	_ = v568
	var v570 int64
	_ = v570
	var v574 int64
	_ = v574
	var v580 int32
	_ = v580
	var v582 int32
	_ = v582
	var v589 int32
	_ = v589
	var v591 int32
	_ = v591
	var v593 int32
	_ = v593
	var v595 int32
	_ = v595
	var v597 int64
	_ = v597
	var v601 int64
	_ = v601
	var v606 int32
	_ = v606
	var v607 int32
	_ = v607
	var v609 int32
	_ = v609
	var v611 int32
	_ = v611
	var v613 int32
	_ = v613
	var v617 int32
	_ = v617
	var v631 int32
	_ = v631
	var v634 int32
	_ = v634
	var v636 int32
	_ = v636
	var v637 int32
	_ = v637
	var v644 int32
	_ = v644
	var v648 int32
	_ = v648
	var v649 int32
	_ = v649
	var v665 int32
	_ = v665
	var v672 int32
	_ = v672
	var v673 int32
	_ = v673
	var v680 int32
	_ = v680
	var v681 int32
	_ = v681
	var v691 int32
	_ = v691
	var v692 int32
	_ = v692
	var v696 int32
	_ = v696
	var v701 int32
	_ = v701
	var v703 int32
	_ = v703
	var v722 int32
	_ = v722
	var v724 int32
	_ = v724
	var v726 int32
	_ = v726
	var v727 int32
	_ = v727
	var v735 int32
	_ = v735
	var v739 int32
	_ = v739
	var v752 int32
	_ = v752
	var v753 int32
	_ = v753
	var v757 int32
	_ = v757
	var v764 int32
	_ = v764
	var v770 int32
	_ = v770
	var v771 int32
	_ = v771
	var v776 int32
	_ = v776
	var v789 int32
	_ = v789
	var v808 int32
	_ = v808
	var v809 int32
	_ = v809
	v8 = m.G0
	v10 = v8 - int32(32)
	m.G0 = v10
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(v10)+28)) = int32(2)
	v15 = int32(-10)
	if v12 <= v15 {
		*(*int32)(unsafe.Add(mBase, uint32(v10)+28)) = int32(1)
		*(*int32)(unsafe.Add(mBase, uint32(v10)+28)) = int32(0)
	} else {
		v21 = v15
		for {
			v26 = int32(0)
			v28 = *(*int64)(unsafe.Add(mBase, _c_F_generateSkyline[0]))
			v32 = v28*int64(6364136223846793005) + int64(1)
			*(*int64)(unsafe.Add(mBase, _c_F_generateSkyline[0])) = v32
			v38 = base.I32_rem_s(base.I32_wrap_i64(int64(base.Ui64(v32)>>(uint(int64(33))%64))), int32(8))
			v39 = v38 + v21
			*(*int32)(unsafe.Add(mBase, uint32(v10)+12)) = v39
			v42 = int32(0)
			v44 = *(*int64)(unsafe.Add(mBase, _c_F_generateSkyline[0]))
			v48 = v44*int64(6364136223846793005) + int64(1)
			*(*int64)(unsafe.Add(mBase, _c_F_generateSkyline[0])) = v48
			v54 = base.I32_rem_s(base.I32_wrap_i64(int64(base.Ui64(v48)>>(uint(int64(33))%64))), int32(9))
			v56 = v54 + int32(10)
			*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = v56
			v58 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			v60 = int32(0)
			v62 = *(*int64)(unsafe.Add(mBase, _c_F_generateSkyline[0]))
			v66 = v62*int64(6364136223846793005) + int64(1)
			*(*int64)(unsafe.Add(mBase, _c_F_generateSkyline[0])) = v66
			v71 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			*(*int32)(unsafe.Add(mBase, uint32(v10)+24)) = int32(0)
			v74 = base.I32_rem_s(base.I32_wrap_i64(int64(base.Ui64(v66)>>(uint(int64(33))%64))), v71)
			v75 = int32(2)
			v76 = base.I32_div_s(v74, v75)
			v78 = base.I32_div_s(v58, v75)
			*(*int32)(unsafe.Add(mBase, uint32(v10)+20)) = v76 + v78
			v82 = v10 + int32(12)
			v96 = *(*int32)(unsafe.Add(mBase, uint32(v82)+8))
			if v96 < int32(1) {
			} else {
				v99 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
				v101 = v99 + int32(-1)
				v102 = v101 - v96
				v109 = *(*int32)(unsafe.Add(mBase, uint32(v82)+4))
				v113 = v109
				v114 = v101
				for {
					if v113 < int32(1) {
						v241 = v113
					} else {
						v130 = *(*int32)(unsafe.Add(mBase, uint32(v82)))
						v137 = v130
						v138 = v113
						v145 = v130
						v146 = v113 + v130
						for {
							if v114 != v102+int32(1) {
								v156 = *(*int32)(unsafe.Add(mBase, uint32(v82)+16))
								v157 = *(*int32)(unsafe.Add(mBase, uint32(v82)+12))
								if v157 == int32(0) {
									v204 = v156
								} else {
									v161 = v145 + int32(1)
									if v137 <= v161 {
										v204 = v156
									} else {
										if v146+int32(-2) <= v137 {
											v204 = v156
										} else {
											if v114 <= v102+int32(2) {
												v204 = v156
											} else {
												if v99+int32(-2) <= v114 {
													v204 = v156
												} else {
													v166 = v137 - v161
													v168 = base.I32_div_s(v166, int32(2))
													if v168&((v114-v102)&int32(1)) == int32(0) {
														v204 = v156
													} else {
														for {
															v187 = F_rand(m)
															mBase = m.M
															v189 = base.I32_rem_s(v187, int32(2))
															v191 = v189 + int32(1)
															v192 = *(*int32)(unsafe.Add(mBase, uint32(v82)+16))
															if v191 == v192 {
																continue
															} else {
																break
															}
															break
														}
														if v166&int32(1) == int32(0) {
															v204 = v191
														} else {
															v200 = F_lwGetPixel(m, l0, v137+int32(-1), v114)
															mBase = m.M
															v204 = v200
														}
													}
												}
											}
										}
									}
								}
								F_lwDrawPixel(m, l0, v137, v114, v204)
								mBase = m.M
								v217 = *(*int32)(unsafe.Add(mBase, uint32(v82)+4))
								v218 = *(*int32)(unsafe.Add(mBase, uint32(v82)))
								v222 = v217
								v229 = v218
							} else {
								if v137 <= v145+int32(1) {
									v222 = v138
									v229 = v145
								} else {
									if v146+int32(-2) <= v137 {
										v222 = v138
										v229 = v145
									} else {
										v156 = *(*int32)(unsafe.Add(mBase, uint32(v82)+16))
										v157 = *(*int32)(unsafe.Add(mBase, uint32(v82)+12))
										if v157 == int32(0) {
											v204 = v156
										} else {
											v161 = v145 + int32(1)
											if v137 <= v161 {
												v204 = v156
											} else {
												if v146+int32(-2) <= v137 {
													v204 = v156
												} else {
													if v114 <= v102+int32(2) {
														v204 = v156
													} else {
														if v99+int32(-2) <= v114 {
															v204 = v156
														} else {
															v166 = v137 - v161
															v168 = base.I32_div_s(v166, int32(2))
															if v168&((v114-v102)&int32(1)) == int32(0) {
																v204 = v156
															} else {
																for {
																	v187 = F_rand(m)
																	mBase = m.M
																	v189 = base.I32_rem_s(v187, int32(2))
																	v191 = v189 + int32(1)
																	v192 = *(*int32)(unsafe.Add(mBase, uint32(v82)+16))
																	if v191 == v192 {
																		continue
																	} else {
																		break
																	}
																	break
																}
																if v166&int32(1) == int32(0) {
																	v204 = v191
																} else {
																	v200 = F_lwGetPixel(m, l0, v137+int32(-1), v114)
																	mBase = m.M
																	v204 = v200
																}
															}
														}
													}
												}
											}
										}
										F_lwDrawPixel(m, l0, v137, v114, v204)
										mBase = m.M
										v217 = *(*int32)(unsafe.Add(mBase, uint32(v82)+4))
										v218 = *(*int32)(unsafe.Add(mBase, uint32(v82)))
										v222 = v217
										v229 = v218
									}
								}
							}
							v235 = v137 + int32(1)
							v236 = v222 + v229
							if v235 < v236 {
								v137 = v235
								v138 = v222
								v145 = v229
								v146 = v236
								continue
							} else {
								break
							}
							break
						}
						v241 = v222
					}
					v254 = v114 + int32(-1)
					if v102 < v254 {
						v113 = v241
						v114 = v254
						continue
					} else {
						break
					}
					break
				}
			}
			v273 = int32(base.Ui32(v56)>>(uint(int32(1))%32)) + v39
			v274 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			if v273 < v274 {
				v21 = v273
				continue
			} else {
				break
			}
			break
		}
		*(*int32)(unsafe.Add(mBase, uint32(v10)+28)) = int32(1)
		if v274 < int32(-9) {
			*(*int32)(unsafe.Add(mBase, uint32(v10)+28)) = int32(0)
		} else {
			v284 = int32(-10)
			for {
				v289 = int32(0)
				v291 = *(*int64)(unsafe.Add(mBase, _c_F_generateSkyline[0]))
				v295 = v291*int64(6364136223846793005) + int64(1)
				*(*int64)(unsafe.Add(mBase, _c_F_generateSkyline[0])) = v295
				v301 = base.I32_rem_s(base.I32_wrap_i64(int64(base.Ui64(v295)>>(uint(int64(33))%64))), int32(8))
				v302 = v301 + v284
				*(*int32)(unsafe.Add(mBase, uint32(v10)+12)) = v302
				v305 = int32(0)
				v307 = *(*int64)(unsafe.Add(mBase, _c_F_generateSkyline[0]))
				v311 = v307*int64(6364136223846793005) + int64(1)
				*(*int64)(unsafe.Add(mBase, _c_F_generateSkyline[0])) = v311
				v317 = base.I32_rem_s(base.I32_wrap_i64(int64(base.Ui64(v311)>>(uint(int64(33))%64))), int32(9))
				*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = v317 + int32(10)
				v321 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
				v323 = int32(0)
				v325 = *(*int64)(unsafe.Add(mBase, _c_F_generateSkyline[0]))
				v329 = v325*int64(6364136223846793005) + int64(1)
				*(*int64)(unsafe.Add(mBase, _c_F_generateSkyline[0])) = v329
				v334 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
				*(*int32)(unsafe.Add(mBase, uint32(v10)+24)) = int32(0)
				v337 = base.I32_rem_s(base.I32_wrap_i64(int64(base.Ui64(v329)>>(uint(int64(33))%64))), v334)
				v339 = base.I32_div_s(v337, int32(3))
				v341 = base.I32_div_s(v321, int32(2))
				*(*int32)(unsafe.Add(mBase, uint32(v10)+20)) = v339 + v341
				v345 = v10 + int32(12)
				v359 = *(*int32)(unsafe.Add(mBase, uint32(v345)+8))
				if v359 < int32(1) {
				} else {
					v362 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
					v364 = v362 + int32(-1)
					v365 = v364 - v359
					v372 = *(*int32)(unsafe.Add(mBase, uint32(v345)+4))
					v376 = v372
					v377 = v364
					for {
						if v376 < int32(1) {
							v504 = v376
						} else {
							v393 = *(*int32)(unsafe.Add(mBase, uint32(v345)))
							v400 = v393
							v401 = v376
							v408 = v393
							v409 = v376 + v393
							for {
								if v377 != v365+int32(1) {
									v419 = *(*int32)(unsafe.Add(mBase, uint32(v345)+16))
									v420 = *(*int32)(unsafe.Add(mBase, uint32(v345)+12))
									if v420 == int32(0) {
										v467 = v419
									} else {
										v424 = v408 + int32(1)
										if v400 <= v424 {
											v467 = v419
										} else {
											if v409+int32(-2) <= v400 {
												v467 = v419
											} else {
												if v377 <= v365+int32(2) {
													v467 = v419
												} else {
													if v362+int32(-2) <= v377 {
														v467 = v419
													} else {
														v429 = v400 - v424
														v431 = base.I32_div_s(v429, int32(2))
														if v431&((v377-v365)&int32(1)) == int32(0) {
															v467 = v419
														} else {
															for {
																v450 = F_rand(m)
																mBase = m.M
																v452 = base.I32_rem_s(v450, int32(2))
																v454 = v452 + int32(1)
																v455 = *(*int32)(unsafe.Add(mBase, uint32(v345)+16))
																if v454 == v455 {
																	continue
																} else {
																	break
																}
																break
															}
															if v429&int32(1) == int32(0) {
																v467 = v454
															} else {
																v463 = F_lwGetPixel(m, l0, v400+int32(-1), v377)
																mBase = m.M
																v467 = v463
															}
														}
													}
												}
											}
										}
									}
									F_lwDrawPixel(m, l0, v400, v377, v467)
									mBase = m.M
									v480 = *(*int32)(unsafe.Add(mBase, uint32(v345)+4))
									v481 = *(*int32)(unsafe.Add(mBase, uint32(v345)))
									v485 = v480
									v492 = v481
								} else {
									if v400 <= v408+int32(1) {
										v485 = v401
										v492 = v408
									} else {
										if v409+int32(-2) <= v400 {
											v485 = v401
											v492 = v408
										} else {
											v419 = *(*int32)(unsafe.Add(mBase, uint32(v345)+16))
											v420 = *(*int32)(unsafe.Add(mBase, uint32(v345)+12))
											if v420 == int32(0) {
												v467 = v419
											} else {
												v424 = v408 + int32(1)
												if v400 <= v424 {
													v467 = v419
												} else {
													if v409+int32(-2) <= v400 {
														v467 = v419
													} else {
														if v377 <= v365+int32(2) {
															v467 = v419
														} else {
															if v362+int32(-2) <= v377 {
																v467 = v419
															} else {
																v429 = v400 - v424
																v431 = base.I32_div_s(v429, int32(2))
																if v431&((v377-v365)&int32(1)) == int32(0) {
																	v467 = v419
																} else {
																	for {
																		v450 = F_rand(m)
																		mBase = m.M
																		v452 = base.I32_rem_s(v450, int32(2))
																		v454 = v452 + int32(1)
																		v455 = *(*int32)(unsafe.Add(mBase, uint32(v345)+16))
																		if v454 == v455 {
																			continue
																		} else {
																			break
																		}
																		break
																	}
																	if v429&int32(1) == int32(0) {
																		v467 = v454
																	} else {
																		v463 = F_lwGetPixel(m, l0, v400+int32(-1), v377)
																		mBase = m.M
																		v467 = v463
																	}
																}
															}
														}
													}
												}
											}
											F_lwDrawPixel(m, l0, v400, v377, v467)
											mBase = m.M
											v480 = *(*int32)(unsafe.Add(mBase, uint32(v345)+4))
											v481 = *(*int32)(unsafe.Add(mBase, uint32(v345)))
											v485 = v480
											v492 = v481
										}
									}
								}
								v498 = v400 + int32(1)
								v499 = v485 + v492
								if v498 < v499 {
									v400 = v498
									v401 = v485
									v408 = v492
									v409 = v499
									continue
								} else {
									break
								}
								break
							}
							v504 = v485
						}
						v517 = v377 + int32(-1)
						if v365 < v517 {
							v376 = v504
							v377 = v517
							continue
						} else {
							break
						}
						break
					}
				}
				v536 = v317 + v302 + int32(11)
				v537 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
				if v536 < v537 {
					v284 = v536
					continue
				} else {
					break
				}
				break
			}
			*(*int32)(unsafe.Add(mBase, uint32(v10)+28)) = int32(0)
			v541 = int32(-10)
			if v537 <= v541 {
			} else {
				v547 = v541
				for {
					v552 = int32(0)
					v554 = *(*int64)(unsafe.Add(mBase, _c_F_generateSkyline[0]))
					v558 = v554*int64(6364136223846793005) + int64(1)
					*(*int64)(unsafe.Add(mBase, _c_F_generateSkyline[0])) = v558
					v564 = base.I32_rem_s(base.I32_wrap_i64(int64(base.Ui64(v558)>>(uint(int64(33))%64))), int32(8))
					v565 = v564 + v547
					*(*int32)(unsafe.Add(mBase, uint32(v10)+12)) = v565
					v568 = int32(0)
					v570 = *(*int64)(unsafe.Add(mBase, _c_F_generateSkyline[0]))
					v574 = v570*int64(6364136223846793005) + int64(1)
					*(*int64)(unsafe.Add(mBase, _c_F_generateSkyline[0])) = v574
					v580 = base.I32_rem_s(base.I32_wrap_i64(int64(base.Ui64(v574)>>(uint(int64(33))%64))), int32(14))
					v582 = v580 + int32(5)
					if v582&int32(3) == int32(0) {
						v591 = v582
					} else {
						v589 = base.I32_rem_s(base.I32_extend8_s(v582), int32(3))
						v591 = v582 + v589
					}
					*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = v591
					v593 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
					v595 = int32(0)
					v597 = *(*int64)(unsafe.Add(mBase, _c_F_generateSkyline[0]))
					v601 = v597*int64(6364136223846793005) + int64(1)
					*(*int64)(unsafe.Add(mBase, _c_F_generateSkyline[0])) = v601
					v606 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
					v607 = int32(1)
					*(*int32)(unsafe.Add(mBase, uint32(v10)+24)) = v607
					v609 = base.I32_rem_s(base.I32_wrap_i64(int64(base.Ui64(v601)>>(uint(int64(33))%64))), v606)
					v611 = base.I32_div_s(v609, int32(2))
					v613 = base.I32_div_s(v593, int32(3))
					*(*int32)(unsafe.Add(mBase, uint32(v10)+20)) = v611 + v613
					v617 = v10 + int32(12)
					v631 = *(*int32)(unsafe.Add(mBase, uint32(v617)+8))
					if v631 < v607 {
					} else {
						v634 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
						v636 = v634 + int32(-1)
						v637 = v636 - v631
						v644 = *(*int32)(unsafe.Add(mBase, uint32(v617)+4))
						v648 = v644
						v649 = v636
						for {
							if v648 < int32(1) {
								v776 = v648
							} else {
								v665 = *(*int32)(unsafe.Add(mBase, uint32(v617)))
								v672 = v665
								v673 = v648
								v680 = v665
								v681 = v648 + v665
								for {
									if v649 != v637+int32(1) {
										v691 = *(*int32)(unsafe.Add(mBase, uint32(v617)+16))
										v692 = *(*int32)(unsafe.Add(mBase, uint32(v617)+12))
										if v692 == int32(0) {
											v739 = v691
										} else {
											v696 = v680 + int32(1)
											if v672 <= v696 {
												v739 = v691
											} else {
												if v681+int32(-2) <= v672 {
													v739 = v691
												} else {
													if v649 <= v637+int32(2) {
														v739 = v691
													} else {
														if v634+int32(-2) <= v649 {
															v739 = v691
														} else {
															v701 = v672 - v696
															v703 = base.I32_div_s(v701, int32(2))
															if v703&((v649-v637)&int32(1)) == int32(0) {
																v739 = v691
															} else {
																for {
																	v722 = F_rand(m)
																	mBase = m.M
																	v724 = base.I32_rem_s(v722, int32(2))
																	v726 = v724 + int32(1)
																	v727 = *(*int32)(unsafe.Add(mBase, uint32(v617)+16))
																	if v726 == v727 {
																		continue
																	} else {
																		break
																	}
																	break
																}
																if v701&int32(1) == int32(0) {
																	v739 = v726
																} else {
																	v735 = F_lwGetPixel(m, l0, v672+int32(-1), v649)
																	mBase = m.M
																	v739 = v735
																}
															}
														}
													}
												}
											}
										}
										F_lwDrawPixel(m, l0, v672, v649, v739)
										mBase = m.M
										v752 = *(*int32)(unsafe.Add(mBase, uint32(v617)+4))
										v753 = *(*int32)(unsafe.Add(mBase, uint32(v617)))
										v757 = v752
										v764 = v753
									} else {
										if v672 <= v680+int32(1) {
											v757 = v673
											v764 = v680
										} else {
											if v681+int32(-2) <= v672 {
												v757 = v673
												v764 = v680
											} else {
												v691 = *(*int32)(unsafe.Add(mBase, uint32(v617)+16))
												v692 = *(*int32)(unsafe.Add(mBase, uint32(v617)+12))
												if v692 == int32(0) {
													v739 = v691
												} else {
													v696 = v680 + int32(1)
													if v672 <= v696 {
														v739 = v691
													} else {
														if v681+int32(-2) <= v672 {
															v739 = v691
														} else {
															if v649 <= v637+int32(2) {
																v739 = v691
															} else {
																if v634+int32(-2) <= v649 {
																	v739 = v691
																} else {
																	v701 = v672 - v696
																	v703 = base.I32_div_s(v701, int32(2))
																	if v703&((v649-v637)&int32(1)) == int32(0) {
																		v739 = v691
																	} else {
																		for {
																			v722 = F_rand(m)
																			mBase = m.M
																			v724 = base.I32_rem_s(v722, int32(2))
																			v726 = v724 + int32(1)
																			v727 = *(*int32)(unsafe.Add(mBase, uint32(v617)+16))
																			if v726 == v727 {
																				continue
																			} else {
																				break
																			}
																			break
																		}
																		if v701&int32(1) == int32(0) {
																			v739 = v726
																		} else {
																			v735 = F_lwGetPixel(m, l0, v672+int32(-1), v649)
																			mBase = m.M
																			v739 = v735
																		}
																	}
																}
															}
														}
													}
												}
												F_lwDrawPixel(m, l0, v672, v649, v739)
												mBase = m.M
												v752 = *(*int32)(unsafe.Add(mBase, uint32(v617)+4))
												v753 = *(*int32)(unsafe.Add(mBase, uint32(v617)))
												v757 = v752
												v764 = v753
											}
										}
									}
									v770 = v672 + int32(1)
									v771 = v757 + v764
									if v770 < v771 {
										v672 = v770
										v673 = v757
										v680 = v764
										v681 = v771
										continue
									} else {
										break
									}
									break
								}
								v776 = v757
							}
							v789 = v649 + int32(-1)
							if v637 < v789 {
								v648 = v776
								v649 = v789
								continue
							} else {
								break
							}
							break
						}
					}
					v808 = v565 + v591 + int32(5)
					v809 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
					if v808 < v809 {
						v547 = v808
						continue
					} else {
						break
					}
					break
				}
			}
		}
	}
	m.G0 = v10 + int32(32)
	return
}
func F_geoaddCommand(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
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
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
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
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v100 int32
	_ = v100
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v127 int32
	_ = v127
	var v130 int32
	_ = v130
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v136 int32
	_ = v136
	var v140 int32
	_ = v140
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v189 int32
	_ = v189
	var v193 int32
	_ = v193
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v212 int32
	_ = v212
	var v215 int32
	_ = v215
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v221 int32
	_ = v221
	var v226 int32
	_ = v226
	var v240 int32
	_ = v240
	var v242 int32
	_ = v242
	var v244 int32
	_ = v244
	var v247 int32
	_ = v247
	var v249 int32
	_ = v249
	var v254 int32
	_ = v254
	var v255 int32
	_ = v255
	var v268 int32
	_ = v268
	var v278 int32
	_ = v278
	var v283 int32
	_ = v283
	var v284 int32
	_ = v284
	var v285 int32
	_ = v285
	var v289 int32
	_ = v289
	var v290 int32
	_ = v290
	var v291 int32
	_ = v291
	var v293 int32
	_ = v293
	var v294 int32
	_ = v294
	var v295 float64
	_ = v295
	var v296 float64
	_ = v296
	var v309 int32
	_ = v309
	var v320 int32
	_ = v320
	var v333 int32
	_ = v333
	var v337 int32
	_ = v337
	var v339 int32
	_ = v339
	var v357 int32
	_ = v357
	var v361 int32
	_ = v361
	var v363 int32
	_ = v363
	var v370 int64
	_ = v370
	var v372 int64
	_ = v372
	var v376 int64
	_ = v376
	var v378 int64
	_ = v378
	var v385 int32
	_ = v385
	var v386 int32
	_ = v386
	var v387 int32
	_ = v387
	var v392 int32
	_ = v392
	var v394 int32
	_ = v394
	var v400 int32
	_ = v400
	var v402 int32
	_ = v402
	var v420 int32
	_ = v420
	var v422 int32
	_ = v422
	var v439 int32
	_ = v439
	var v441 int32
	_ = v441
	v16 = m.G0
	v18 = v16 - int32(64)
	m.G0 = v18
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v20 < int32(3) {
		goto L4
	} else {
		goto L5
	}
L1:
	;
	m.G0 = v18 + int32(64)
	return
L2:
	;
	v439 = *(*int32)(unsafe.Add(mBase, _c_F_geoaddCommand[0]))
	F_addReplyErrorObject(m, l0, v439)
	mBase = m.M
	v441 = m.ExcPending
	if v441 != 0 {
		goto L55
	} else {
		goto L91
	}
L3:
	;
	v212 = base.I32_div_s(v202, int32(3))
	v215 = v212<<(uint(int32(1))%32) + v201
	v218 = F_valkey_calloc(m, v215<<(uint(int32(2))%32))
	mBase = m.M
	v219 = m.ExcPending
	if v219 != 0 {
		goto L55
	} else {
		goto L56
	}
L4:
	;
	v189 = int32(2)
	v193 = base.I32_rem_u_s(v189-v20, int32(3))
	if v193 != 0 {
		goto L2
	} else {
		goto L54
	}
L5:
	;
	v23 = int32(0)
	v29 = v23
	v30 = v23
	v31 = int32(2)
	goto L7
L6:
	;
	v179 = v175 - v178
	v181 = base.I32_rem_s(v179, int32(3))
	if v181 != 0 {
		goto L2
	} else {
		goto L52
	}
L7:
	;
	v41 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v41+v31<<(uint(int32(2))%32))))
	v46 = F_objectGetVal(m, v45)
	mBase = m.M
	v47 = int32(_a_F_geoaddCommand_0)
	v50 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v46))))
	if v50 != 0 {
		goto L13
	} else {
		goto L14
	}
L8:
	;
	v175 = v173
	v176 = v169
	v177 = v170
	v178 = v172
	goto L6
L9:
	;
	v172 = v31 + int32(1)
	v173 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v172 < v173 {
		v29 = v169
		v30 = v170
		v31 = v172
		goto L7
	} else {
		goto L51
	}
L10:
	;
	v87 = int32(_a_F_geoaddCommand_1)
	v90 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v46))))
	if v90 != 0 {
		goto L27
	} else {
		goto L28
	}
L11:
	;
	if v82-v84 != 0 {
		goto L10
	} else {
		goto L23
	}
L12:
	;
	v82 = F_tolower(m, v78)
	mBase = m.M
	v83 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v79))))
	v84 = F_tolower(m, v83)
	mBase = m.M
	goto L11
L13:
	;
	v52 = v46
	v53 = v47
	v54 = v50
	goto L16
L14:
	;
	v78 = int32(0)
	v79 = v47
	goto L12
L15:
	;
	v78 = v75 & int32(255)
	v79 = v74
	goto L12
L16:
	;
	v56 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v53))))
	if v56 == int32(0) {
		v74 = v53
		v75 = v54
		goto L15
	} else {
		goto L18
	}
L17:
	;
	v74 = v68
	v75 = int32(0)
	goto L15
L18:
	;
	v60 = v54 & int32(255)
	if v60 == v56 {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	v67 = int32(1)
	v68 = v53 + v67
	v69 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v52)+1)))
	if v69 != 0 {
		v52 = v52 + v67
		v53 = v68
		v54 = v69
		goto L16
	} else {
		goto L22
	}
L20:
	;
	v62 = F_tolower(m, v60)
	mBase = m.M
	v63 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v53))))
	v64 = F_tolower(m, v63)
	mBase = m.M
	if v62 == v64 {
		goto L19
	} else {
		goto L21
	}
L21:
	;
	v66 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v52))))
	v74 = v53
	v75 = v66
	goto L15
L22:
	;
	goto L17
L23:
	;
	v169 = v29
	v170 = int32(1)
	goto L9
L24:
	;
	v127 = int32(_a_F_geoaddCommand_2)
	v130 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v46))))
	if v130 != 0 {
		goto L40
	} else {
		goto L41
	}
L25:
	;
	if v122-v124 != 0 {
		goto L24
	} else {
		goto L37
	}
L26:
	;
	v122 = F_tolower(m, v118)
	mBase = m.M
	v123 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v119))))
	v124 = F_tolower(m, v123)
	mBase = m.M
	goto L25
L27:
	;
	v92 = v46
	v93 = v87
	v94 = v90
	goto L30
L28:
	;
	v118 = int32(0)
	v119 = v87
	goto L26
L29:
	;
	v118 = v115 & int32(255)
	v119 = v114
	goto L26
L30:
	;
	v96 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v93))))
	if v96 == int32(0) {
		v114 = v93
		v115 = v94
		goto L29
	} else {
		goto L32
	}
L31:
	;
	v114 = v108
	v115 = int32(0)
	goto L29
L32:
	;
	v100 = v94 & int32(255)
	if v100 == v96 {
		goto L33
	} else {
		goto L34
	}
L33:
	;
	v107 = int32(1)
	v108 = v93 + v107
	v109 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v92)+1)))
	if v109 != 0 {
		v92 = v92 + v107
		v93 = v108
		v94 = v109
		goto L30
	} else {
		goto L36
	}
L34:
	;
	v102 = F_tolower(m, v100)
	mBase = m.M
	v103 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v93))))
	v104 = F_tolower(m, v103)
	mBase = m.M
	if v102 == v104 {
		goto L33
	} else {
		goto L35
	}
L35:
	;
	v106 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v92))))
	v114 = v93
	v115 = v106
	goto L29
L36:
	;
	goto L31
L37:
	;
	v169 = int32(1)
	v170 = v30
	goto L9
L38:
	;
	if v162-v164 == int32(0) {
		v169 = v29
		v170 = v30
		goto L9
	} else {
		goto L50
	}
L39:
	;
	v162 = F_tolower(m, v158)
	mBase = m.M
	v163 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v159))))
	v164 = F_tolower(m, v163)
	mBase = m.M
	goto L38
L40:
	;
	v132 = v46
	v133 = v127
	v134 = v130
	goto L43
L41:
	;
	v158 = int32(0)
	v159 = v127
	goto L39
L42:
	;
	v158 = v155 & int32(255)
	v159 = v154
	goto L39
L43:
	;
	v136 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v133))))
	if v136 == int32(0) {
		v154 = v133
		v155 = v134
		goto L42
	} else {
		goto L45
	}
L44:
	;
	v154 = v148
	v155 = int32(0)
	goto L42
L45:
	;
	v140 = v134 & int32(255)
	if v140 == v136 {
		goto L46
	} else {
		goto L47
	}
L46:
	;
	v147 = int32(1)
	v148 = v133 + v147
	v149 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v132)+1)))
	if v149 != 0 {
		v132 = v132 + v147
		v133 = v148
		v134 = v149
		goto L43
	} else {
		goto L49
	}
L47:
	;
	v142 = F_tolower(m, v140)
	mBase = m.M
	v143 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v133))))
	v144 = F_tolower(m, v143)
	mBase = m.M
	if v142 == v144 {
		goto L46
	} else {
		goto L48
	}
L48:
	;
	v146 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v132))))
	v154 = v133
	v155 = v146
	goto L42
L49:
	;
	goto L44
L50:
	;
	v168 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v175 = v168
	v176 = v29
	v177 = v30
	v178 = v31
	goto L6
L51:
	;
	goto L8
L52:
	;
	v182 = int32(0)
	if base.B2i32(v176 != v182)&base.B2i32(v177 != v182) == v182 {
		v201 = v178
		v202 = v179
		goto L3
	} else {
		goto L53
	}
L53:
	;
	goto L2
L54:
	;
	v201 = v189
	v202 = v20 + int32(-2)
	goto L3
L55:
	;
	return
L56:
	;
	v221 = *(*int32)(unsafe.Add(mBase, _c_F_geoaddCommand[1]))
	*(*int32)(unsafe.Add(mBase, uint32(v218))) = v221
	v226 = int32(1)
	goto L57
L57:
	;
	v240 = v226 << (uint(int32(2)) % 32)
	v242 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v244 = *(*int32)(unsafe.Add(mBase, uint32(v242+v240)))
	*(*int32)(unsafe.Add(mBase, uint32(v218+v240))) = v244
	F_incrRefCount(m, v244)
	mBase = m.M
	v247 = m.ExcPending
	if v247 != 0 {
		goto L55
	} else {
		goto L59
	}
L58:
	;
	if v202 < int32(3) {
		goto L61
	} else {
		goto L62
	}
L59:
	;
	v249 = v226 + int32(1)
	if v249 != v201 {
		v226 = v249
		goto L57
	} else {
		goto L60
	}
L60:
	;
	goto L58
L61:
	;
	F_replaceClientCommandVector(m, l0, v215, v218)
	mBase = m.M
	v420 = m.ExcPending
	if v420 != 0 {
		goto L55
	} else {
		goto L89
	}
L62:
	;
	v254 = v201 << (uint(int32(2)) % 32)
	v255 = v218 + v254
	v268 = int32(0)
	goto L63
L63:
	;
	v278 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v283 = v268 * int32(3) << (uint(int32(2)) % 32)
	v284 = v278 + v254 + v283
	v285 = *(*int32)(unsafe.Add(mBase, uint32(v284)))
	v289 = F_getDoubleFromObjectOrReply(m, l0, v285, v18+int32(48), int32(0))
	mBase = m.M
	v290 = m.ExcPending
	if v290 != 0 {
		goto L55
	} else {
		goto L67
	}
L64:
	;
	goto L61
L65:
	;
	v361 = F_geohashEncodeType(m, v296, v295, int32(26), v18+int32(32))
	mBase = m.M
	goto L84
L66:
	;
	if v215 < int32(1) {
		goto L75
	} else {
		goto L76
	}
L67:
	;
	if v289 != 0 {
		goto L66
	} else {
		goto L68
	}
L68:
	;
	v291 = *(*int32)(unsafe.Add(mBase, uint32(v284)+4))
	v293 = F_getDoubleFromObjectOrReply(m, l0, v291, v18+int32(48)|int32(8), int32(0))
	mBase = m.M
	v294 = m.ExcPending
	if v294 != 0 {
		goto L55
	} else {
		goto L69
	}
L69:
	;
	if v293 != 0 {
		goto L66
	} else {
		goto L70
	}
L70:
	;
	v295 = *(*float64)(unsafe.Add(mBase, uint32(v18)+56))
	v296 = *(*float64)(unsafe.Add(mBase, uint32(v18)+48))
	if base.F64_gt(base.F64_abs(v296), float64(180)) != 0 {
		goto L71
	} else {
		goto L72
	}
L71:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v18)+8)) = v295
	*(*float64)(unsafe.Add(mBase, uint32(v18))) = v296
	F_addReplyErrorFormat(m, l0, int32(_a_F_geoaddCommand_3), v18)
	mBase = m.M
	v309 = m.ExcPending
	if v309 != 0 {
		goto L55
	} else {
		goto L74
	}
L72:
	;
	if base.F64_gt(base.F64_abs(v295), float64(85.05112878)) == int32(0) {
		goto L65
	} else {
		goto L73
	}
L73:
	;
	goto L71
L74:
	;
	goto L66
L75:
	;
	F_valkey_free(m, v218)
	mBase = m.M
	v357 = m.ExcPending
	if v357 != 0 {
		goto L55
	} else {
		goto L83
	}
L76:
	;
	v320 = int32(0)
	goto L77
L77:
	;
	v333 = *(*int32)(unsafe.Add(mBase, uint32(v218+v320<<(uint(int32(2))%32))))
	if v333 == int32(0) {
		goto L79
	} else {
		goto L80
	}
L78:
	;
	goto L75
L79:
	;
	v339 = v320 + int32(1)
	if v339 != v215 {
		v320 = v339
		goto L77
	} else {
		goto L82
	}
L80:
	;
	F_decrRefCount(m, v333)
	mBase = m.M
	v337 = m.ExcPending
	if v337 != 0 {
		goto L55
	} else {
		goto L81
	}
L81:
	;
	goto L79
L82:
	;
	goto L78
L83:
	;
	goto L1
L84:
	;
	v363 = v18 + int32(16)
	v370 = *(*int64)(unsafe.Add(mBase, uint32(v18+int32(40))))
	*(*int64)(unsafe.Add(mBase, uint32(v18+int32(24)))) = v370
	v372 = *(*int64)(unsafe.Add(mBase, uint32(v18)+32))
	*(*int64)(unsafe.Add(mBase, uint32(v18)+16)) = v372
	v376 = *(*int64)(unsafe.Add(mBase, uint32(v363)))
	v378 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v363)+8)))
	goto L85
L85:
	;
	v385 = F_createStringObjectFromLongLongWithSds(m, v376<<(uint((int64(52)-v378<<(uint(int64(1))%64))&int64(4294967294))%64))
	mBase = m.M
	v386 = m.ExcPending
	if v386 != 0 {
		goto L55
	} else {
		goto L86
	}
L86:
	;
	v387 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v392 = *(*int32)(unsafe.Add(mBase, uint32(v387+v254+v283+int32(8))))
	v394 = v268 << (uint(int32(3)) % 32)
	*(*int32)(unsafe.Add(mBase, uint32(v255+v394))) = v385
	*(*int32)(unsafe.Add(mBase, uint32(v255+int32(4)+v394))) = v392
	F_incrRefCount(m, v392)
	mBase = m.M
	v400 = m.ExcPending
	if v400 != 0 {
		goto L55
	} else {
		goto L87
	}
L87:
	;
	v402 = v268 + int32(1)
	if v402 != v212 {
		v268 = v402
		goto L63
	} else {
		goto L88
	}
L88:
	;
	goto L64
L89:
	;
	F_zaddCommand(m, l0)
	mBase = m.M
	v422 = m.ExcPending
	if v422 != 0 {
		goto L55
	} else {
		goto L90
	}
L90:
	;
	goto L1
L91:
	;
	goto L1
}
func F_georadiusGeneric(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v56 int32
	_ = v56
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v76 float64
	_ = v76
	var v77 float64
	_ = v77
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v100 float64
	_ = v100
	var v107 int32
	_ = v107
	var v109 int32
	_ = v109
	var v110 float64
	_ = v110
	var v111 int32
	_ = v111
	var v117 int32
	_ = v117
	var v119 int32
	_ = v119
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v132 int32
	_ = v132
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v140 float64
	_ = v140
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v149 int32
	_ = v149
	var v157 int64
	_ = v157
	var v166 int64
	_ = v166
	var v168 int64
	_ = v168
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v186 int64
	_ = v186
	var v188 int64
	_ = v188
	var v190 int32
	_ = v190
	var v196 int32
	_ = v196
	var v198 int32
	_ = v198
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v211 float64
	_ = v211
	var v218 int32
	_ = v218
	var v220 int32
	_ = v220
	var v221 float64
	_ = v221
	var v222 int32
	_ = v222
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v240 int32
	_ = v240
	var v243 int32
	_ = v243
	var v248 int32
	_ = v248
	var v251 int32
	_ = v251
	var v252 int32
	_ = v252
	var v255 int32
	_ = v255
	var v257 int32
	_ = v257
	var v270 int32
	_ = v270
	var v272 int32
	_ = v272
	var v276 int32
	_ = v276
	var v279 int32
	_ = v279
	var v280 int32
	_ = v280
	var v284 int32
	_ = v284
	var v289 int32
	_ = v289
	var v302 int32
	_ = v302
	var v306 int32
	_ = v306
	var v314 int32
	_ = v314
	var v315 int32
	_ = v315
	var v316 int32
	_ = v316
	var v317 int32
	_ = v317
	var v318 int32
	_ = v318
	var v328 int32
	_ = v328
	var v329 int32
	_ = v329
	var v330 int32
	_ = v330
	var v331 int32
	_ = v331
	var v332 int32
	_ = v332
	var v333 int32
	_ = v333
	var v335 int32
	_ = v335
	var v338 int32
	_ = v338
	var v340 int32
	_ = v340
	var v341 int32
	_ = v341
	var v342 int32
	_ = v342
	var v345 int32
	_ = v345
	var v347 int32
	_ = v347
	var v348 int32
	_ = v348
	var v349 int32
	_ = v349
	var v351 int32
	_ = v351
	var v355 int32
	_ = v355
	var v357 int32
	_ = v357
	var v358 int32
	_ = v358
	var v359 int32
	_ = v359
	var v361 int32
	_ = v361
	var v362 int32
	_ = v362
	var v363 int32
	_ = v363
	var v364 int32
	_ = v364
	var v369 int32
	_ = v369
	var v370 int32
	_ = v370
	var v373 int32
	_ = v373
	var v374 int32
	_ = v374
	var v377 int32
	_ = v377
	var v378 int32
	_ = v378
	var v379 int32
	_ = v379
	var v382 int32
	_ = v382
	var v385 int32
	_ = v385
	var v387 int32
	_ = v387
	var v388 int32
	_ = v388
	var v389 int32
	_ = v389
	var v391 int32
	_ = v391
	var v395 int32
	_ = v395
	var v397 int32
	_ = v397
	var v398 int32
	_ = v398
	var v399 int32
	_ = v399
	var v401 int32
	_ = v401
	var v402 int32
	_ = v402
	var v403 int32
	_ = v403
	var v404 int32
	_ = v404
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
	var v418 int32
	_ = v418
	var v419 int32
	_ = v419
	var v422 int32
	_ = v422
	var v425 int32
	_ = v425
	var v427 int32
	_ = v427
	var v428 int32
	_ = v428
	var v429 int32
	_ = v429
	var v431 int32
	_ = v431
	var v435 int32
	_ = v435
	var v437 int32
	_ = v437
	var v438 int32
	_ = v438
	var v439 int32
	_ = v439
	var v441 int32
	_ = v441
	var v442 int32
	_ = v442
	var v443 int32
	_ = v443
	var v444 int32
	_ = v444
	var v449 int32
	_ = v449
	var v450 int32
	_ = v450
	var v453 int32
	_ = v453
	var v454 int32
	_ = v454
	var v457 int32
	_ = v457
	var v458 int32
	_ = v458
	var v459 int32
	_ = v459
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
	var v502 int32
	_ = v502
	var v505 int32
	_ = v505
	var v507 int32
	_ = v507
	var v508 int32
	_ = v508
	var v509 int32
	_ = v509
	var v511 int32
	_ = v511
	var v515 int32
	_ = v515
	var v517 int32
	_ = v517
	var v518 int32
	_ = v518
	var v519 int32
	_ = v519
	var v521 int32
	_ = v521
	var v522 int32
	_ = v522
	var v523 int32
	_ = v523
	var v524 int32
	_ = v524
	var v529 int32
	_ = v529
	var v530 int32
	_ = v530
	var v533 int32
	_ = v533
	var v534 int32
	_ = v534
	var v537 int32
	_ = v537
	var v538 int32
	_ = v538
	var v539 int32
	_ = v539
	var v542 int32
	_ = v542
	var v545 int32
	_ = v545
	var v547 int32
	_ = v547
	var v548 int32
	_ = v548
	var v549 int32
	_ = v549
	var v551 int32
	_ = v551
	var v555 int32
	_ = v555
	var v557 int32
	_ = v557
	var v558 int32
	_ = v558
	var v559 int32
	_ = v559
	var v561 int32
	_ = v561
	var v562 int32
	_ = v562
	var v563 int32
	_ = v563
	var v564 int32
	_ = v564
	var v569 int32
	_ = v569
	var v570 int32
	_ = v570
	var v573 int32
	_ = v573
	var v574 int32
	_ = v574
	var v577 int32
	_ = v577
	var v578 int32
	_ = v578
	var v579 int32
	_ = v579
	var v582 int32
	_ = v582
	var v585 int32
	_ = v585
	var v587 int32
	_ = v587
	var v588 int32
	_ = v588
	var v589 int32
	_ = v589
	var v591 int32
	_ = v591
	var v595 int32
	_ = v595
	var v597 int32
	_ = v597
	var v598 int32
	_ = v598
	var v599 int32
	_ = v599
	var v601 int32
	_ = v601
	var v602 int32
	_ = v602
	var v603 int32
	_ = v603
	var v604 int32
	_ = v604
	var v609 int32
	_ = v609
	var v610 int32
	_ = v610
	var v613 int32
	_ = v613
	var v614 int32
	_ = v614
	var v617 int32
	_ = v617
	var v618 int32
	_ = v618
	var v619 int32
	_ = v619
	var v622 int32
	_ = v622
	var v624 int32
	_ = v624
	var v628 int32
	_ = v628
	var v632 int32
	_ = v632
	var v633 int32
	_ = v633
	var v636 int32
	_ = v636
	var v639 int32
	_ = v639
	var v643 int32
	_ = v643
	var v644 int64
	_ = v644
	var v649 int32
	_ = v649
	var v650 int32
	_ = v650
	var v653 int32
	_ = v653
	var v657 int32
	_ = v657
	var v659 int32
	_ = v659
	var v662 int32
	_ = v662
	var v664 int32
	_ = v664
	var v665 int32
	_ = v665
	var v666 int32
	_ = v666
	var v668 int32
	_ = v668
	var v672 int32
	_ = v672
	var v674 int32
	_ = v674
	var v675 int32
	_ = v675
	var v676 int32
	_ = v676
	var v678 int32
	_ = v678
	var v679 int32
	_ = v679
	var v680 int32
	_ = v680
	var v681 int32
	_ = v681
	var v686 int32
	_ = v686
	var v687 int32
	_ = v687
	var v690 int32
	_ = v690
	var v691 int32
	_ = v691
	var v694 int32
	_ = v694
	var v695 int32
	_ = v695
	var v696 int32
	_ = v696
	var v699 int32
	_ = v699
	var v701 int32
	_ = v701
	var v705 int32
	_ = v705
	var v708 int32
	_ = v708
	var v711 int32
	_ = v711
	var v713 int32
	_ = v713
	var v714 int32
	_ = v714
	var v715 int32
	_ = v715
	var v717 int32
	_ = v717
	var v721 int32
	_ = v721
	var v723 int32
	_ = v723
	var v724 int32
	_ = v724
	var v725 int32
	_ = v725
	var v727 int32
	_ = v727
	var v728 int32
	_ = v728
	var v729 int32
	_ = v729
	var v730 int32
	_ = v730
	var v735 int32
	_ = v735
	var v736 int32
	_ = v736
	var v739 int32
	_ = v739
	var v740 int32
	_ = v740
	var v743 int32
	_ = v743
	var v744 int32
	_ = v744
	var v745 int32
	_ = v745
	var v748 int32
	_ = v748
	var v750 int32
	_ = v750
	var v754 int32
	_ = v754
	var v761 int32
	_ = v761
	var v764 int32
	_ = v764
	var v766 int32
	_ = v766
	var v767 int32
	_ = v767
	var v768 int32
	_ = v768
	var v770 int32
	_ = v770
	var v774 int32
	_ = v774
	var v776 int32
	_ = v776
	var v777 int32
	_ = v777
	var v778 int32
	_ = v778
	var v780 int32
	_ = v780
	var v781 int32
	_ = v781
	var v782 int32
	_ = v782
	var v783 int32
	_ = v783
	var v788 int32
	_ = v788
	var v789 int32
	_ = v789
	var v792 int32
	_ = v792
	var v793 int32
	_ = v793
	var v796 int32
	_ = v796
	var v797 int32
	_ = v797
	var v798 int32
	_ = v798
	var v801 int32
	_ = v801
	var v805 int32
	_ = v805
	var v806 int32
	_ = v806
	var v809 int32
	_ = v809
	var v813 int32
	_ = v813
	var v814 int32
	_ = v814
	var v815 int32
	_ = v815
	var v818 int32
	_ = v818
	var v821 int32
	_ = v821
	var v825 int32
	_ = v825
	var v827 int32
	_ = v827
	var v830 int32
	_ = v830
	var v832 int32
	_ = v832
	var v833 int32
	_ = v833
	var v834 int32
	_ = v834
	var v836 int32
	_ = v836
	var v840 int32
	_ = v840
	var v842 int32
	_ = v842
	var v843 int32
	_ = v843
	var v844 int32
	_ = v844
	var v846 int32
	_ = v846
	var v847 int32
	_ = v847
	var v848 int32
	_ = v848
	var v849 int32
	_ = v849
	var v854 int32
	_ = v854
	var v855 int32
	_ = v855
	var v858 int32
	_ = v858
	var v859 int32
	_ = v859
	var v862 int32
	_ = v862
	var v863 int32
	_ = v863
	var v864 int32
	_ = v864
	var v867 int32
	_ = v867
	var v871 int32
	_ = v871
	var v872 int32
	_ = v872
	var v880 int32
	_ = v880
	var v881 int32
	_ = v881
	var v885 int32
	_ = v885
	var v887 int32
	_ = v887
	var v890 int32
	_ = v890
	var v894 int32
	_ = v894
	var v896 int32
	_ = v896
	var v899 int32
	_ = v899
	var v901 int32
	_ = v901
	var v902 int32
	_ = v902
	var v903 int32
	_ = v903
	var v905 int32
	_ = v905
	var v909 int32
	_ = v909
	var v911 int32
	_ = v911
	var v912 int32
	_ = v912
	var v913 int32
	_ = v913
	var v915 int32
	_ = v915
	var v916 int32
	_ = v916
	var v917 int32
	_ = v917
	var v918 int32
	_ = v918
	var v923 int32
	_ = v923
	var v924 int32
	_ = v924
	var v927 int32
	_ = v927
	var v928 int32
	_ = v928
	var v931 int32
	_ = v931
	var v932 int32
	_ = v932
	var v933 int32
	_ = v933
	var v936 int32
	_ = v936
	var v940 int32
	_ = v940
	var v941 int32
	_ = v941
	var v949 int32
	_ = v949
	var v950 int32
	_ = v950
	var v953 int32
	_ = v953
	var v956 int32
	_ = v956
	var v960 int32
	_ = v960
	var v962 int32
	_ = v962
	var v965 int32
	_ = v965
	var v967 int32
	_ = v967
	var v968 int32
	_ = v968
	var v969 int32
	_ = v969
	var v971 int32
	_ = v971
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
	var v982 int32
	_ = v982
	var v983 int32
	_ = v983
	var v984 int32
	_ = v984
	var v989 int32
	_ = v989
	var v990 int32
	_ = v990
	var v993 int32
	_ = v993
	var v994 int32
	_ = v994
	var v997 int32
	_ = v997
	var v998 int32
	_ = v998
	var v999 int32
	_ = v999
	var v1002 int32
	_ = v1002
	var v1006 int32
	_ = v1006
	var v1007 int32
	_ = v1007
	var v1015 int32
	_ = v1015
	var v1016 int32
	_ = v1016
	var v1019 int32
	_ = v1019
	var v1022 int32
	_ = v1022
	var v1026 int32
	_ = v1026
	var v1030 int32
	_ = v1030
	var v1033 int32
	_ = v1033
	var v1036 int32
	_ = v1036
	var v1038 int32
	_ = v1038
	var v1039 int32
	_ = v1039
	var v1040 int32
	_ = v1040
	var v1042 int32
	_ = v1042
	var v1046 int32
	_ = v1046
	var v1048 int32
	_ = v1048
	var v1049 int32
	_ = v1049
	var v1050 int32
	_ = v1050
	var v1052 int32
	_ = v1052
	var v1053 int32
	_ = v1053
	var v1054 int32
	_ = v1054
	var v1055 int32
	_ = v1055
	var v1060 int32
	_ = v1060
	var v1061 int32
	_ = v1061
	var v1064 int32
	_ = v1064
	var v1065 int32
	_ = v1065
	var v1068 int32
	_ = v1068
	var v1069 int32
	_ = v1069
	var v1070 int32
	_ = v1070
	var v1079 int32
	_ = v1079
	var v1083 int32
	_ = v1083
	var v1087 int32
	_ = v1087
	var v1088 int32
	_ = v1088
	var v1091 int32
	_ = v1091
	var v1094 int32
	_ = v1094
	var v1099 int32
	_ = v1099
	var v1103 int32
	_ = v1103
	var v1104 int32
	_ = v1104
	var v1112 int32
	_ = v1112
	var v1113 int32
	_ = v1113
	var v1117 int32
	_ = v1117
	var v1118 int32
	_ = v1118
	var v1132 int32
	_ = v1132
	var v1157 int32
	_ = v1157
	var v1158 int32
	_ = v1158
	var v1169 int32
	_ = v1169
	var v1173 int32
	_ = v1173
	var v1174 int32
	_ = v1174
	var v1178 int32
	_ = v1178
	var v1179 int32
	_ = v1179
	var v1181 int32
	_ = v1181
	var v1183 int32
	_ = v1183
	var v1220 int32
	_ = v1220
	var v1224 int32
	_ = v1224
	var v1226 int32
	_ = v1226
	var v1228 int32
	_ = v1228
	var v1229 int32
	_ = v1229
	var v1232 int32
	_ = v1232
	var v1236 int32
	_ = v1236
	var v1237 int32
	_ = v1237
	var v1240 int32
	_ = v1240
	var v1260 int32
	_ = v1260
	var v1276 int32
	_ = v1276
	var v1280 int32
	_ = v1280
	var v1289 int32
	_ = v1289
	var v1297 int32
	_ = v1297
	var v1298 int32
	_ = v1298
	var v1299 int32
	_ = v1299
	var v1300 int32
	_ = v1300
	var v1301 int32
	_ = v1301
	var v1302 int32
	_ = v1302
	var v1311 int32
	_ = v1311
	var v1312 int32
	_ = v1312
	var v1313 int32
	_ = v1313
	var v1314 int32
	_ = v1314
	var v1315 int32
	_ = v1315
	var v1316 int32
	_ = v1316
	var v1319 int32
	_ = v1319
	var v1321 int32
	_ = v1321
	var v1338 int32
	_ = v1338
	var v1342 int32
	_ = v1342
	var v1345 int32
	_ = v1345
	var v1346 int32
	_ = v1346
	var v1348 int32
	_ = v1348
	var v1349 int32
	_ = v1349
	var v1350 int32
	_ = v1350
	var v1351 int32
	_ = v1351
	var v1352 int32
	_ = v1352
	var v1353 int32
	_ = v1353
	var v1354 int32
	_ = v1354
	var v1355 int32
	_ = v1355
	var v1382 int32
	_ = v1382
	var v1388 int32
	_ = v1388
	var v1389 int32
	_ = v1389
	var v1392 int32
	_ = v1392
	var v1396 int32
	_ = v1396
	var v1398 int32
	_ = v1398
	var v1400 int32
	_ = v1400
	var v1406 int32
	_ = v1406
	var v1407 int32
	_ = v1407
	var v1408 int32
	_ = v1408
	var v1414 int32
	_ = v1414
	var v1415 int32
	_ = v1415
	var v1418 int32
	_ = v1418
	var v1422 int32
	_ = v1422
	var v1428 int32
	_ = v1428
	var v1429 int32
	_ = v1429
	var v1430 int32
	_ = v1430
	var v1436 int32
	_ = v1436
	var v1437 int32
	_ = v1437
	var v1440 int32
	_ = v1440
	var v1444 int32
	_ = v1444
	var v1445 int64
	_ = v1445
	var v1452 int32
	_ = v1452
	var v1453 int32
	_ = v1453
	var v1456 int32
	_ = v1456
	var v1460 int32
	_ = v1460
	var v1464 int32
	_ = v1464
	var v1465 int32
	_ = v1465
	var v1466 int32
	_ = v1466
	var v1470 int32
	_ = v1470
	var v1472 int32
	_ = v1472
	var v1475 int32
	_ = v1475
	var v1476 int32
	_ = v1476
	var v1478 int32
	_ = v1478
	var v1479 int32
	_ = v1479
	var v1481 int64
	_ = v1481
	var v1488 int32
	_ = v1488
	var v1489 int32
	_ = v1489
	var v1491 int32
	_ = v1491
	var v1492 int32
	_ = v1492
	var v1495 int32
	_ = v1495
	var v1499 int32
	_ = v1499
	var v1501 int32
	_ = v1501
	var v1503 int32
	_ = v1503
	var v1516 int32
	_ = v1516
	var v1518 int32
	_ = v1518
	var v1519 int32
	_ = v1519
	var v1520 int32
	_ = v1520
	var v1521 int32
	_ = v1521
	var v1523 int32
	_ = v1523
	var v1525 int32
	_ = v1525
	var v1529 int32
	_ = v1529
	var v1530 int32
	_ = v1530
	var v1533 int64
	_ = v1533
	var v1537 int32
	_ = v1537
	var v1540 int32
	_ = v1540
	var v1544 int32
	_ = v1544
	var v1547 int32
	_ = v1547
	var v1549 int32
	_ = v1549
	var v1550 int32
	_ = v1550
	var v1554 int32
	_ = v1554
	var v1560 int32
	_ = v1560
	var v1564 int32
	_ = v1564
	var v1569 int32
	_ = v1569
	var v1570 int32
	_ = v1570
	var v1571 int32
	_ = v1571
	var v1577 int32
	_ = v1577
	var v1609 int32
	_ = v1609
	var v1612 int32
	_ = v1612
	var v1613 float64
	_ = v1613
	var v1614 float64
	_ = v1614
	var v1620 int32
	_ = v1620
	var v1621 int32
	_ = v1621
	var v1623 int32
	_ = v1623
	var v1624 int32
	_ = v1624
	var v1629 int32
	_ = v1629
	var v1633 float64
	_ = v1633
	var v1635 int32
	_ = v1635
	var v1637 int32
	_ = v1637
	var v1640 float64
	_ = v1640
	var v1646 int64
	_ = v1646
	var v1648 int64
	_ = v1648
	var v1650 int32
	_ = v1650
	var v1657 int32
	_ = v1657
	var v1658 int32
	_ = v1658
	var v1659 int32
	_ = v1659
	var v1660 float64
	_ = v1660
	var v1667 int32
	_ = v1667
	var v1669 int32
	_ = v1669
	var v1671 int64
	_ = v1671
	var v1673 int64
	_ = v1673
	var v1677 int64
	_ = v1677
	var v1697 int64
	_ = v1697
	var v1711 int32
	_ = v1711
	var v1720 int64
	_ = v1720
	var v1723 int64
	_ = v1723
	var v1724 int64
	_ = v1724
	var v1725 int64
	_ = v1725
	var v1726 int64
	_ = v1726
	var v1739 int64
	_ = v1739
	var v1744 int64
	_ = v1744
	var v1746 int32
	_ = v1746
	var v1747 float64
	_ = v1747
	var v1754 int32
	_ = v1754
	var v1756 int32
	_ = v1756
	var v1758 int64
	_ = v1758
	var v1760 int64
	_ = v1760
	var v1764 int64
	_ = v1764
	var v1784 int64
	_ = v1784
	var v1798 int32
	_ = v1798
	var v1807 int64
	_ = v1807
	var v1810 int64
	_ = v1810
	var v1811 int64
	_ = v1811
	var v1812 int64
	_ = v1812
	var v1813 int64
	_ = v1813
	var v1826 int64
	_ = v1826
	var v1829 int64
	_ = v1829
	var v1831 int32
	_ = v1831
	var v1833 int32
	_ = v1833
	var v1837 int32
	_ = v1837
	var v1838 int32
	_ = v1838
	var v1840 int32
	_ = v1840
	var v1841 int32
	_ = v1841
	var v1845 int32
	_ = v1845
	var v1850 int32
	_ = v1850
	var v1860 int32
	_ = v1860
	var v1866 int32
	_ = v1866
	var v1882 int32
	_ = v1882
	var v1885 int32
	_ = v1885
	var v1886 float64
	_ = v1886
	var v1887 float64
	_ = v1887
	var v1888 float64
	_ = v1888
	var v1892 float64
	_ = v1892
	var v1893 float64
	_ = v1893
	var v1895 int32
	_ = v1895
	var v1898 int32
	_ = v1898
	var v1905 int32
	_ = v1905
	var v1908 int32
	_ = v1908
	var v1911 int32
	_ = v1911
	var v1914 int32
	_ = v1914
	var v1915 int32
	_ = v1915
	var v1916 int32
	_ = v1916
	var v1917 int32
	_ = v1917
	var v1918 int32
	_ = v1918
	var v1919 int32
	_ = v1919
	var v1920 int32
	_ = v1920
	var v1921 int32
	_ = v1921
	var v1924 int32
	_ = v1924
	var v1926 int32
	_ = v1926
	var v1930 int32
	_ = v1930
	var v1931 int32
	_ = v1931
	var v1933 int32
	_ = v1933
	var v1947 int32
	_ = v1947
	var v1953 int32
	_ = v1953
	var v1970 int32
	_ = v1970
	var v1971 int32
	_ = v1971
	var v1976 int32
	_ = v1976
	var v1980 int32
	_ = v1980
	var v1981 int32
	_ = v1981
	var v1982 int32
	_ = v1982
	var v1984 int32
	_ = v1984
	var v1986 int32
	_ = v1986
	var v1987 int32
	_ = v1987
	var v1988 int32
	_ = v1988
	var v1991 int32
	_ = v1991
	var v1993 int32
	_ = v1993
	var v1996 int32
	_ = v1996
	var v1997 int32
	_ = v1997
	var v1999 int32
	_ = v1999
	var v2011 int64
	_ = v2011
	var v2035 int32
	_ = v2035
	var v2037 int64
	_ = v2037
	var v2076 int32
	_ = v2076
	var v2114 int32
	_ = v2114
	var v2115 int32
	_ = v2115
	var v2152 int32
	_ = v2152
	var v2156 int32
	_ = v2156
	var v2161 int32
	_ = v2161
	var v2165 int32
	_ = v2165
	v35 = m.G0
	v37 = v35 - int32(880)
	m.G0 = v37
	v39 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v40 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v40+l1<<(uint(int32(2))%32))))
	v45 = F_lookupKeyRead(m, v39, v44)
	mBase = m.M
	v46 = m.ExcPending
	if v46 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v37 + int32(880)
	return
L2:
	;
	return
L3:
	;
	v48 = F_checkType(m, l0, v45, int32(3))
	mBase = m.M
	v49 = m.ExcPending
	if v49 != 0 {
		goto L2
	} else {
		goto L4
	}
L4:
	;
	if v48 != 0 {
		goto L1
	} else {
		goto L5
	}
L5:
	;
	v50 = int32(0)
	v56 = F__emscripten_memset_bulkmem(m, v37+int32(672), base.I32_extend8_s(v50), int32(80))
	mBase = m.M
	goto L6
L6:
	;
	if l2&int32(1) == int32(0) {
		goto L9
	} else {
		goto L10
	}
L7:
	;
	F_addReplyError(m, l0, int32(_a_F_georadiusGeneric_0))
	mBase = m.M
	v2165 = m.ExcPending
	if v2165 != 0 {
		goto L2
	} else {
		goto L522
	}
L8:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v37)+664)) = int64(0)
	v248 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v243 < v248 {
		goto L55
	} else {
		goto L56
	}
L9:
	;
	v117 = int32(5)
	v119 = l2 & int32(2)
	if v119 == int32(0) {
		goto L27
	} else {
		goto L28
	}
L10:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v37)+672)) = int32(1)
	v63 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v63)+8))
	v68 = F_getDoubleFromObjectOrReply(m, l0, v64, v37+int32(680), int32(0))
	mBase = m.M
	v69 = m.ExcPending
	if v69 != 0 {
		goto L2
	} else {
		goto L11
	}
L11:
	;
	if v68 != 0 {
		goto L1
	} else {
		goto L12
	}
L12:
	;
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v63)+12))
	v74 = F_getDoubleFromObjectOrReply(m, l0, v70, v37+int32(688), int32(0))
	mBase = m.M
	v75 = m.ExcPending
	if v75 != 0 {
		goto L2
	} else {
		goto L13
	}
L13:
	;
	if v74 != 0 {
		goto L1
	} else {
		goto L14
	}
L14:
	;
	v76 = *(*float64)(unsafe.Add(mBase, uint32(v37)+688))
	v77 = *(*float64)(unsafe.Add(mBase, uint32(v37)+680))
	if base.F64_gt(base.F64_abs(v77), float64(180)) != 0 {
		goto L16
	} else {
		goto L17
	}
L15:
	;
	v93 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v94 = *(*int32)(unsafe.Add(mBase, uint32(v93)+16))
	v98 = F_getDoubleFromObjectOrReply(m, l0, v94, v37+int32(136), int32(_a_F_georadiusGeneric_1))
	mBase = m.M
	v99 = m.ExcPending
	if v99 != 0 {
		goto L2
	} else {
		goto L20
	}
L16:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v37)+120)) = v76
	*(*float64)(unsafe.Add(mBase, uint32(v37)+112)) = v77
	F_addReplyErrorFormat(m, l0, int32(_a_F_georadiusGeneric_2), v37+int32(112))
	mBase = m.M
	v92 = m.ExcPending
	if v92 != 0 {
		goto L2
	} else {
		goto L19
	}
L17:
	;
	if base.F64_gt(base.F64_abs(v76), float64(85.05112878)) == int32(0) {
		goto L15
	} else {
		goto L18
	}
L18:
	;
	goto L16
L19:
	;
	goto L1
L20:
	;
	if v98 != 0 {
		goto L1
	} else {
		goto L21
	}
L21:
	;
	v100 = *(*float64)(unsafe.Add(mBase, uint32(v37)+136))
	if base.F64_lt(v100, float64(0)) == int32(0) {
		goto L22
	} else {
		goto L23
	}
L22:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v37)+736)) = v100
	v109 = *(*int32)(unsafe.Add(mBase, uint32(v93)+20))
	v110 = F_extractUnitOrReply(m, l0, v109)
	mBase = m.M
	v111 = m.ExcPending
	if v111 != 0 {
		goto L2
	} else {
		goto L25
	}
L23:
	;
	F_addReplyError(m, l0, int32(_a_F_georadiusGeneric_3))
	mBase = m.M
	v107 = m.ExcPending
	if v107 != 0 {
		goto L2
	} else {
		goto L24
	}
L24:
	;
	goto L1
L25:
	;
	if base.F64_lt(v110, float64(0)) != 0 {
		goto L1
	} else {
		goto L26
	}
L26:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v37)+696)) = v110
	v240 = int32(0)
	v243 = int32(6)
	goto L8
L27:
	;
	if v119 == int32(0) {
		goto L30
	} else {
		goto L31
	}
L28:
	;
	if v45 == int32(0) {
		v240 = v50
		v243 = v117
		goto L8
	} else {
		goto L29
	}
L29:
	;
	goto L27
L30:
	;
	if l2&int32(8) == int32(0) {
		goto L7
	} else {
		goto L50
	}
L31:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v37)+672)) = int32(1)
	v128 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v129 = *(*int32)(unsafe.Add(mBase, uint32(v128)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v37)+472)) = int64(0)
	v132 = F_objectGetVal(m, v129)
	mBase = m.M
	v135 = F_zsetScore(m, v45, v132, v37+int32(472))
	mBase = m.M
	v136 = m.ExcPending
	if v136 != 0 {
		goto L2
	} else {
		goto L35
	}
L32:
	;
	v204 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v205 = *(*int32)(unsafe.Add(mBase, uint32(v204)+12))
	v209 = F_getDoubleFromObjectOrReply(m, l0, v205, v37+int32(136), int32(_a_F_georadiusGeneric_1))
	mBase = m.M
	v210 = m.ExcPending
	if v210 != 0 {
		goto L2
	} else {
		goto L43
	}
L33:
	;
	v198 = F_objectGetVal(m, v129)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v37)+80)) = v198
	F_addReplyErrorFormat(m, l0, v196, v37+int32(80))
	mBase = m.M
	v203 = m.ExcPending
	if v203 != 0 {
		goto L2
	} else {
		goto L42
	}
L34:
	;
	v140 = *(*float64)(unsafe.Add(mBase, uint32(v37)+472))
	v144 = v37 + int32(144)
	v145 = int32(26)
	*(*uint8)(unsafe.Add(mBase, uint32(v144))) = uint8(v145)
	v149 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v37+int32(148)))) = v149
	*(*int32)(unsafe.Add(mBase, uint32(v37)+145)) = v149
	v157 = *(*int64)(unsafe.Add(mBase, uint32(v144)))
	*(*int64)(unsafe.Add(mBase, uint32(v37+int32(104)))) = v157
	if base.F64_lt(v140, float64(1.8446744073709552e+19))&base.F64_ge(v140, float64(0)) == v149 {
		goto L38
	} else {
		goto L39
	}
L35:
	;
	if v135 != int32(-1) {
		goto L34
	} else {
		goto L36
	}
L36:
	;
	v196 = int32(_a_F_georadiusGeneric_4)
	goto L33
L37:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v37)+136)) = v168
	*(*int64)(unsafe.Add(mBase, uint32(v37)+96)) = v168
	v178 = m.G0
	v179 = int32(16)
	v180 = v178 - v179
	m.G0 = v180
	v186 = *(*int64)(unsafe.Add(mBase, uint32(v37+int32(104))))
	*(*int64)(unsafe.Add(mBase, uint32(v180+int32(8)))) = v186
	v188 = *(*int64)(unsafe.Add(mBase, uint32(v37+int32(96))))
	*(*int64)(unsafe.Add(mBase, uint32(v180))) = v188
	v190 = F_geohashDecodeToLongLatType(m, v180, v37+int32(680))
	mBase = m.M
	m.G0 = v180 + v179
	goto L40
L38:
	;
	v168 = int64(0)
	goto L37
L39:
	;
	v166 = base.I64_trunc_f64_u(v140)
	v168 = v166
	goto L37
L40:
	;
	if v190 != 0 {
		goto L32
	} else {
		goto L41
	}
L41:
	;
	v196 = int32(_a_F_georadiusGeneric_5)
	goto L33
L42:
	;
	goto L1
L43:
	;
	if v209 != 0 {
		goto L1
	} else {
		goto L44
	}
L44:
	;
	v211 = *(*float64)(unsafe.Add(mBase, uint32(v37)+136))
	if base.F64_lt(v211, float64(0)) == int32(0) {
		goto L45
	} else {
		goto L46
	}
L45:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v37)+736)) = v211
	v220 = *(*int32)(unsafe.Add(mBase, uint32(v204)+16))
	v221 = F_extractUnitOrReply(m, l0, v220)
	mBase = m.M
	v222 = m.ExcPending
	if v222 != 0 {
		goto L2
	} else {
		goto L48
	}
L46:
	;
	F_addReplyError(m, l0, int32(_a_F_georadiusGeneric_3))
	mBase = m.M
	v218 = m.ExcPending
	if v218 != 0 {
		goto L2
	} else {
		goto L47
	}
L47:
	;
	goto L1
L48:
	;
	if base.F64_lt(v221, float64(0)) != 0 {
		goto L1
	} else {
		goto L49
	}
L49:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v37)+696)) = v221
	v240 = int32(0)
	v243 = v117
	goto L8
L50:
	;
	if l2&int32(16) == int32(0) {
		v240 = v50
		v243 = int32(2)
		goto L8
	} else {
		goto L51
	}
L51:
	;
	v236 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v237 = *(*int32)(unsafe.Add(mBase, uint32(v236)+4))
	v240 = v237
	v243 = int32(3)
	goto L8
L52:
	;
	if v1342 == int32(0) {
		goto L351
	} else {
		goto L352
	}
L53:
	;
	v270 = l2 & int32(8)
	v272 = l2 & int32(12)
	v276 = v37 + int32(736)
	v279 = int32(24)
	v280 = v37 + int32(696)
	v284 = v37 + int32(680)
	v289 = int32(0)
	v302 = v289
	v306 = v240
	v314 = v289
	v315 = v289
	v316 = v289
	v317 = v289
	v318 = v289
	v328 = v289
	v329 = v289
	v330 = v289
	v331 = v289
	v332 = v289
	v333 = v289
	goto L58
L54:
	;
	v257 = int32(0)
	v1338 = v257
	v1342 = v240
	v1345 = v257
	v1346 = v257
	v1348 = v255
	v1349 = v257
	v1350 = v257
	v1351 = v257
	v1352 = v257
	v1353 = v257
	v1354 = v257
	v1355 = v257
	goto L52
L55:
	;
	v251 = int32(1)
	v252 = v248 - v243
	if v251 <= v252 {
		goto L53
	} else {
		goto L57
	}
L56:
	;
	v255 = int32(1)
	goto L54
L57:
	;
	v255 = v251
	goto L54
L58:
	;
	v335 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v338 = (v302 + v243) << (uint(int32(2)) % 32)
	v340 = *(*int32)(unsafe.Add(mBase, uint32(v335+v338)))
	v341 = F_objectGetVal(m, v340)
	mBase = m.M
	v342 = int32(_a_F_georadiusGeneric_6)
	v345 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v341))))
	if v345 != 0 {
		goto L64
	} else {
		goto L65
	}
L59:
	;
	v1321 = int32(0)
	v1338 = base.B2i32(v1316 != v1321)
	v1342 = v1289
	v1345 = base.B2i32(v1315 != v1321)
	v1346 = base.B2i32(v1312 != v1321)
	v1348 = base.B2i32(v1311 == v1321)
	v1349 = base.B2i32(v1314 != v1321)
	v1350 = v1297
	v1351 = v1298
	v1352 = base.B2i32(v1299 != v1321)
	v1353 = base.B2i32(v1300 != v1321)
	v1354 = v1301
	v1355 = base.B2i32(v1313 != v1321)
	goto L52
L60:
	;
	v1319 = v1302 + int32(1)
	if v1319 < v252 {
		v302 = v1319
		v306 = v1289
		v314 = v1297
		v315 = v1298
		v316 = v1299
		v317 = v1300
		v318 = v1301
		v328 = v1311
		v329 = v1312
		v330 = v1313
		v331 = v1314
		v332 = v1315
		v333 = v1316
		goto L58
	} else {
		goto L350
	}
L61:
	;
	v382 = int32(_a_F_georadiusGeneric_7)
	v385 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v341))))
	if v385 != 0 {
		goto L78
	} else {
		goto L79
	}
L62:
	;
	if v377-v379 != 0 {
		goto L61
	} else {
		goto L74
	}
L63:
	;
	v377 = F_tolower(m, v373)
	mBase = m.M
	v378 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v374))))
	v379 = F_tolower(m, v378)
	mBase = m.M
	goto L62
L64:
	;
	v347 = v341
	v348 = v342
	v349 = v345
	goto L67
L65:
	;
	v373 = int32(0)
	v374 = v342
	goto L63
L66:
	;
	v373 = v370 & int32(255)
	v374 = v369
	goto L63
L67:
	;
	v351 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v348))))
	if v351 == int32(0) {
		v369 = v348
		v370 = v349
		goto L66
	} else {
		goto L69
	}
L68:
	;
	v369 = v363
	v370 = int32(0)
	goto L66
L69:
	;
	v355 = v349 & int32(255)
	if v355 == v351 {
		goto L70
	} else {
		goto L71
	}
L70:
	;
	v362 = int32(1)
	v363 = v348 + v362
	v364 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v347)+1)))
	if v364 != 0 {
		v347 = v347 + v362
		v348 = v363
		v349 = v364
		goto L67
	} else {
		goto L73
	}
L71:
	;
	v357 = F_tolower(m, v355)
	mBase = m.M
	v358 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v348))))
	v359 = F_tolower(m, v358)
	mBase = m.M
	if v357 == v359 {
		goto L70
	} else {
		goto L72
	}
L72:
	;
	v361 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v347))))
	v369 = v348
	v370 = v361
	goto L66
L73:
	;
	goto L68
L74:
	;
	v1289 = v306
	v1297 = v314
	v1298 = v315
	v1299 = v316
	v1300 = v317
	v1301 = v318
	v1302 = v302
	v1311 = v328
	v1312 = int32(1)
	v1313 = v330
	v1314 = v331
	v1315 = v332
	v1316 = v333
	goto L60
L75:
	;
	v422 = int32(_a_F_georadiusGeneric_8)
	v425 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v341))))
	if v425 != 0 {
		goto L92
	} else {
		goto L93
	}
L76:
	;
	if v417-v419 != 0 {
		goto L75
	} else {
		goto L88
	}
L77:
	;
	v417 = F_tolower(m, v413)
	mBase = m.M
	v418 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v414))))
	v419 = F_tolower(m, v418)
	mBase = m.M
	goto L76
L78:
	;
	v387 = v341
	v388 = v382
	v389 = v385
	goto L81
L79:
	;
	v413 = int32(0)
	v414 = v382
	goto L77
L80:
	;
	v413 = v410 & int32(255)
	v414 = v409
	goto L77
L81:
	;
	v391 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v388))))
	if v391 == int32(0) {
		v409 = v388
		v410 = v389
		goto L80
	} else {
		goto L83
	}
L82:
	;
	v409 = v403
	v410 = int32(0)
	goto L80
L83:
	;
	v395 = v389 & int32(255)
	if v395 == v391 {
		goto L84
	} else {
		goto L85
	}
L84:
	;
	v402 = int32(1)
	v403 = v388 + v402
	v404 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v387)+1)))
	if v404 != 0 {
		v387 = v387 + v402
		v388 = v403
		v389 = v404
		goto L81
	} else {
		goto L87
	}
L85:
	;
	v397 = F_tolower(m, v395)
	mBase = m.M
	v398 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v388))))
	v399 = F_tolower(m, v398)
	mBase = m.M
	if v397 == v399 {
		goto L84
	} else {
		goto L86
	}
L86:
	;
	v401 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v387))))
	v409 = v388
	v410 = v401
	goto L80
L87:
	;
	goto L82
L88:
	;
	v1289 = v306
	v1297 = v314
	v1298 = v315
	v1299 = v316
	v1300 = v317
	v1301 = v318
	v1302 = v302
	v1311 = v328
	v1312 = v329
	v1313 = int32(1)
	v1314 = v331
	v1315 = v332
	v1316 = v333
	goto L60
L89:
	;
	v462 = int32(_a_F_georadiusGeneric_9)
	v465 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v341))))
	if v465 != 0 {
		goto L106
	} else {
		goto L107
	}
L90:
	;
	if v457-v459 != 0 {
		goto L89
	} else {
		goto L102
	}
L91:
	;
	v457 = F_tolower(m, v453)
	mBase = m.M
	v458 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v454))))
	v459 = F_tolower(m, v458)
	mBase = m.M
	goto L90
L92:
	;
	v427 = v341
	v428 = v422
	v429 = v425
	goto L95
L93:
	;
	v453 = int32(0)
	v454 = v422
	goto L91
L94:
	;
	v453 = v450 & int32(255)
	v454 = v449
	goto L91
L95:
	;
	v431 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v428))))
	if v431 == int32(0) {
		v449 = v428
		v450 = v429
		goto L94
	} else {
		goto L97
	}
L96:
	;
	v449 = v443
	v450 = int32(0)
	goto L94
L97:
	;
	v435 = v429 & int32(255)
	if v435 == v431 {
		goto L98
	} else {
		goto L99
	}
L98:
	;
	v442 = int32(1)
	v443 = v428 + v442
	v444 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v427)+1)))
	if v444 != 0 {
		v427 = v427 + v442
		v428 = v443
		v429 = v444
		goto L95
	} else {
		goto L101
	}
L99:
	;
	v437 = F_tolower(m, v435)
	mBase = m.M
	v438 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v428))))
	v439 = F_tolower(m, v438)
	mBase = m.M
	if v437 == v439 {
		goto L98
	} else {
		goto L100
	}
L100:
	;
	v441 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v427))))
	v449 = v428
	v450 = v441
	goto L94
L101:
	;
	goto L96
L102:
	;
	v1289 = v306
	v1297 = v314
	v1298 = v315
	v1299 = v316
	v1300 = v317
	v1301 = int32(1)
	v1302 = v302
	v1311 = v328
	v1312 = v329
	v1313 = v330
	v1314 = v331
	v1315 = v332
	v1316 = v333
	goto L60
L103:
	;
	v502 = int32(_a_F_georadiusGeneric_10)
	v505 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v341))))
	if v505 != 0 {
		goto L120
	} else {
		goto L121
	}
L104:
	;
	if v497-v499 != 0 {
		goto L103
	} else {
		goto L116
	}
L105:
	;
	v497 = F_tolower(m, v493)
	mBase = m.M
	v498 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v494))))
	v499 = F_tolower(m, v498)
	mBase = m.M
	goto L104
L106:
	;
	v467 = v341
	v468 = v462
	v469 = v465
	goto L109
L107:
	;
	v493 = int32(0)
	v494 = v462
	goto L105
L108:
	;
	v493 = v490 & int32(255)
	v494 = v489
	goto L105
L109:
	;
	v471 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v468))))
	if v471 == int32(0) {
		v489 = v468
		v490 = v469
		goto L108
	} else {
		goto L111
	}
L110:
	;
	v489 = v483
	v490 = int32(0)
	goto L108
L111:
	;
	v475 = v469 & int32(255)
	if v475 == v471 {
		goto L112
	} else {
		goto L113
	}
L112:
	;
	v482 = int32(1)
	v483 = v468 + v482
	v484 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v467)+1)))
	if v484 != 0 {
		v467 = v467 + v482
		v468 = v483
		v469 = v484
		goto L109
	} else {
		goto L115
	}
L113:
	;
	v477 = F_tolower(m, v475)
	mBase = m.M
	v478 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v468))))
	v479 = F_tolower(m, v478)
	mBase = m.M
	if v477 == v479 {
		goto L112
	} else {
		goto L114
	}
L114:
	;
	v481 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v467))))
	v489 = v468
	v490 = v481
	goto L108
L115:
	;
	goto L110
L116:
	;
	v1289 = v306
	v1297 = v314
	v1298 = int32(1)
	v1299 = v316
	v1300 = v317
	v1301 = v318
	v1302 = v302
	v1311 = v328
	v1312 = v329
	v1313 = v330
	v1314 = v331
	v1315 = v332
	v1316 = v333
	goto L60
L117:
	;
	v542 = int32(_a_F_georadiusGeneric_11)
	v545 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v341))))
	if v545 != 0 {
		goto L134
	} else {
		goto L135
	}
L118:
	;
	if v537-v539 != 0 {
		goto L117
	} else {
		goto L130
	}
L119:
	;
	v537 = F_tolower(m, v533)
	mBase = m.M
	v538 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v534))))
	v539 = F_tolower(m, v538)
	mBase = m.M
	goto L118
L120:
	;
	v507 = v341
	v508 = v502
	v509 = v505
	goto L123
L121:
	;
	v533 = int32(0)
	v534 = v502
	goto L119
L122:
	;
	v533 = v530 & int32(255)
	v534 = v529
	goto L119
L123:
	;
	v511 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v508))))
	if v511 == int32(0) {
		v529 = v508
		v530 = v509
		goto L122
	} else {
		goto L125
	}
L124:
	;
	v529 = v523
	v530 = int32(0)
	goto L122
L125:
	;
	v515 = v509 & int32(255)
	if v515 == v511 {
		goto L126
	} else {
		goto L127
	}
L126:
	;
	v522 = int32(1)
	v523 = v508 + v522
	v524 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v507)+1)))
	if v524 != 0 {
		v507 = v507 + v522
		v508 = v523
		v509 = v524
		goto L123
	} else {
		goto L129
	}
L127:
	;
	v517 = F_tolower(m, v515)
	mBase = m.M
	v518 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v508))))
	v519 = F_tolower(m, v518)
	mBase = m.M
	if v517 == v519 {
		goto L126
	} else {
		goto L128
	}
L128:
	;
	v521 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v507))))
	v529 = v508
	v530 = v521
	goto L122
L129:
	;
	goto L124
L130:
	;
	v1289 = v306
	v1297 = int32(1)
	v1298 = v315
	v1299 = v316
	v1300 = v317
	v1301 = v318
	v1302 = v302
	v1311 = v328
	v1312 = v329
	v1313 = v330
	v1314 = v331
	v1315 = v332
	v1316 = v333
	goto L60
L131:
	;
	v582 = int32(_a_F_georadiusGeneric_12)
	v585 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v341))))
	if v585 != 0 {
		goto L148
	} else {
		goto L149
	}
L132:
	;
	if v577-v579 != 0 {
		goto L131
	} else {
		goto L144
	}
L133:
	;
	v577 = F_tolower(m, v573)
	mBase = m.M
	v578 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v574))))
	v579 = F_tolower(m, v578)
	mBase = m.M
	goto L132
L134:
	;
	v547 = v341
	v548 = v542
	v549 = v545
	goto L137
L135:
	;
	v573 = int32(0)
	v574 = v542
	goto L133
L136:
	;
	v573 = v570 & int32(255)
	v574 = v569
	goto L133
L137:
	;
	v551 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v548))))
	if v551 == int32(0) {
		v569 = v548
		v570 = v549
		goto L136
	} else {
		goto L139
	}
L138:
	;
	v569 = v563
	v570 = int32(0)
	goto L136
L139:
	;
	v555 = v549 & int32(255)
	if v555 == v551 {
		goto L140
	} else {
		goto L141
	}
L140:
	;
	v562 = int32(1)
	v563 = v548 + v562
	v564 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v547)+1)))
	if v564 != 0 {
		v547 = v547 + v562
		v548 = v563
		v549 = v564
		goto L137
	} else {
		goto L143
	}
L141:
	;
	v557 = F_tolower(m, v555)
	mBase = m.M
	v558 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v548))))
	v559 = F_tolower(m, v558)
	mBase = m.M
	if v557 == v559 {
		goto L140
	} else {
		goto L142
	}
L142:
	;
	v561 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v547))))
	v569 = v548
	v570 = v561
	goto L136
L143:
	;
	goto L138
L144:
	;
	v1289 = v306
	v1297 = int32(2)
	v1298 = v315
	v1299 = v316
	v1300 = v317
	v1301 = v318
	v1302 = v302
	v1311 = v328
	v1312 = v329
	v1313 = v330
	v1314 = v331
	v1315 = v332
	v1316 = v333
	goto L60
L145:
	;
	v659 = int32(_a_F_georadiusGeneric_13)
	v662 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v341))))
	if v662 != 0 {
		goto L174
	} else {
		goto L175
	}
L146:
	;
	if v617-v619 != 0 {
		goto L145
	} else {
		goto L158
	}
L147:
	;
	v617 = F_tolower(m, v613)
	mBase = m.M
	v618 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v614))))
	v619 = F_tolower(m, v618)
	mBase = m.M
	goto L146
L148:
	;
	v587 = v341
	v588 = v582
	v589 = v585
	goto L151
L149:
	;
	v613 = int32(0)
	v614 = v582
	goto L147
L150:
	;
	v613 = v610 & int32(255)
	v614 = v609
	goto L147
L151:
	;
	v591 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v588))))
	if v591 == int32(0) {
		v609 = v588
		v610 = v589
		goto L150
	} else {
		goto L153
	}
L152:
	;
	v609 = v603
	v610 = int32(0)
	goto L150
L153:
	;
	v595 = v589 & int32(255)
	if v595 == v591 {
		goto L154
	} else {
		goto L155
	}
L154:
	;
	v602 = int32(1)
	v603 = v588 + v602
	v604 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v587)+1)))
	if v604 != 0 {
		v587 = v587 + v602
		v588 = v603
		v589 = v604
		goto L151
	} else {
		goto L157
	}
L155:
	;
	v597 = F_tolower(m, v595)
	mBase = m.M
	v598 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v588))))
	v599 = F_tolower(m, v598)
	mBase = m.M
	if v597 == v599 {
		goto L154
	} else {
		goto L156
	}
L156:
	;
	v601 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v587))))
	v609 = v588
	v610 = v601
	goto L150
L157:
	;
	goto L152
L158:
	;
	v622 = v302 + int32(1)
	if v252 <= v622 {
		goto L145
	} else {
		goto L159
	}
L159:
	;
	v624 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v628 = *(*int32)(unsafe.Add(mBase, uint32(v624+v338+int32(4))))
	v632 = F_getLongLongFromObjectOrReply(m, l0, v628, v37+int32(664), int32(0))
	mBase = m.M
	v633 = m.ExcPending
	if v633 != 0 {
		goto L2
	} else {
		goto L161
	}
L160:
	;
	v644 = *(*int64)(unsafe.Add(mBase, uint32(v37)+664))
	if int64(0) < v644 {
		v1289 = v306
		v1297 = v314
		v1298 = v315
		v1299 = v316
		v1300 = v317
		v1301 = v318
		v1302 = v622
		v1311 = v328
		v1312 = v329
		v1313 = v330
		v1314 = v331
		v1315 = v332
		v1316 = v333
		goto L60
	} else {
		goto L166
	}
L161:
	;
	if v632 == int32(0) {
		goto L160
	} else {
		goto L162
	}
L162:
	;
	v636 = *(*int32)(unsafe.Add(mBase, uint32(v37)+672))
	if v636 != int32(3) {
		goto L1
	} else {
		goto L163
	}
L163:
	;
	v639 = *(*int32)(unsafe.Add(mBase, uint32(v37)+740))
	if v639 == int32(0) {
		goto L1
	} else {
		goto L164
	}
L164:
	;
	F_valkey_free(m, v639)
	mBase = m.M
	v643 = m.ExcPending
	if v643 != 0 {
		goto L2
	} else {
		goto L165
	}
L165:
	;
	goto L1
L166:
	;
	F_addReplyError(m, l0, int32(_a_F_georadiusGeneric_14))
	mBase = m.M
	v649 = m.ExcPending
	if v649 != 0 {
		goto L2
	} else {
		goto L167
	}
L167:
	;
	v650 = *(*int32)(unsafe.Add(mBase, uint32(v37)+672))
	if v650 != int32(3) {
		goto L1
	} else {
		goto L168
	}
L168:
	;
	v653 = *(*int32)(unsafe.Add(mBase, uint32(v37)+740))
	if v653 == int32(0) {
		goto L1
	} else {
		goto L169
	}
L169:
	;
	F_valkey_free(m, v653)
	mBase = m.M
	v657 = m.ExcPending
	if v657 != 0 {
		goto L2
	} else {
		goto L170
	}
L170:
	;
	goto L1
L171:
	;
	v708 = int32(_a_F_georadiusGeneric_15)
	v711 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v341))))
	if v711 != 0 {
		goto L190
	} else {
		goto L191
	}
L172:
	;
	if v694-v696 != 0 {
		goto L171
	} else {
		goto L184
	}
L173:
	;
	v694 = F_tolower(m, v690)
	mBase = m.M
	v695 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v691))))
	v696 = F_tolower(m, v695)
	mBase = m.M
	goto L172
L174:
	;
	v664 = v341
	v665 = v659
	v666 = v662
	goto L177
L175:
	;
	v690 = int32(0)
	v691 = v659
	goto L173
L176:
	;
	v690 = v687 & int32(255)
	v691 = v686
	goto L173
L177:
	;
	v668 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v665))))
	if v668 == int32(0) {
		v686 = v665
		v687 = v666
		goto L176
	} else {
		goto L179
	}
L178:
	;
	v686 = v680
	v687 = int32(0)
	goto L176
L179:
	;
	v672 = v666 & int32(255)
	if v672 == v668 {
		goto L180
	} else {
		goto L181
	}
L180:
	;
	v679 = int32(1)
	v680 = v665 + v679
	v681 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v664)+1)))
	if v681 != 0 {
		v664 = v664 + v679
		v665 = v680
		v666 = v681
		goto L177
	} else {
		goto L183
	}
L181:
	;
	v674 = F_tolower(m, v672)
	mBase = m.M
	v675 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v665))))
	v676 = F_tolower(m, v675)
	mBase = m.M
	if v674 == v676 {
		goto L180
	} else {
		goto L182
	}
L182:
	;
	v678 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v664))))
	v686 = v665
	v687 = v678
	goto L176
L183:
	;
	goto L178
L184:
	;
	if v272 != 0 {
		goto L171
	} else {
		goto L185
	}
L185:
	;
	v699 = v302 + int32(1)
	if v252 <= v699 {
		goto L171
	} else {
		goto L186
	}
L186:
	;
	v701 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v705 = *(*int32)(unsafe.Add(mBase, uint32(v701+v338+int32(4))))
	v1289 = v705
	v1297 = v314
	v1298 = v315
	v1299 = v316
	v1300 = v317
	v1301 = v318
	v1302 = v699
	v1311 = int32(0)
	v1312 = v329
	v1313 = v330
	v1314 = v331
	v1315 = v332
	v1316 = v333
	goto L60
L187:
	;
	v761 = int32(_a_F_georadiusGeneric_16)
	v764 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v341))))
	if v764 != 0 {
		goto L208
	} else {
		goto L209
	}
L188:
	;
	if v743-v745 != 0 {
		goto L187
	} else {
		goto L200
	}
L189:
	;
	v743 = F_tolower(m, v739)
	mBase = m.M
	v744 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v740))))
	v745 = F_tolower(m, v744)
	mBase = m.M
	goto L188
L190:
	;
	v713 = v341
	v714 = v708
	v715 = v711
	goto L193
L191:
	;
	v739 = int32(0)
	v740 = v708
	goto L189
L192:
	;
	v739 = v736 & int32(255)
	v740 = v735
	goto L189
L193:
	;
	v717 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v714))))
	if v717 == int32(0) {
		v735 = v714
		v736 = v715
		goto L192
	} else {
		goto L195
	}
L194:
	;
	v735 = v729
	v736 = int32(0)
	goto L192
L195:
	;
	v721 = v715 & int32(255)
	if v721 == v717 {
		goto L196
	} else {
		goto L197
	}
L196:
	;
	v728 = int32(1)
	v729 = v714 + v728
	v730 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v713)+1)))
	if v730 != 0 {
		v713 = v713 + v728
		v714 = v729
		v715 = v730
		goto L193
	} else {
		goto L199
	}
L197:
	;
	v723 = F_tolower(m, v721)
	mBase = m.M
	v724 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v714))))
	v725 = F_tolower(m, v724)
	mBase = m.M
	if v723 == v725 {
		goto L196
	} else {
		goto L198
	}
L198:
	;
	v727 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v713))))
	v735 = v714
	v736 = v727
	goto L192
L199:
	;
	goto L194
L200:
	;
	if v272 != 0 {
		goto L201
	} else {
		goto L202
	}
L201:
	;
	if base.B2i32(l2&v279 == v279) == int32(0) {
		goto L187
	} else {
		goto L204
	}
L202:
	;
	v748 = v302 + int32(1)
	if v252 <= v748 {
		goto L201
	} else {
		goto L203
	}
L203:
	;
	v750 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v754 = *(*int32)(unsafe.Add(mBase, uint32(v750+v338+int32(4))))
	v1289 = v754
	v1297 = v314
	v1298 = v315
	v1299 = v316
	v1300 = v317
	v1301 = v318
	v1302 = v748
	v1311 = int32(1)
	v1312 = v329
	v1313 = v330
	v1314 = v331
	v1315 = v332
	v1316 = v333
	goto L60
L204:
	;
	v1289 = v306
	v1297 = v314
	v1298 = v315
	v1299 = v316
	v1300 = v317
	v1301 = v318
	v1302 = v302
	v1311 = int32(1)
	v1312 = v329
	v1313 = v330
	v1314 = v331
	v1315 = v332
	v1316 = v333
	goto L60
L205:
	;
	v827 = int32(_a_F_georadiusGeneric_17)
	v830 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v341))))
	if v830 != 0 {
		goto L233
	} else {
		goto L234
	}
L206:
	;
	if v796-v798 != 0 {
		goto L205
	} else {
		goto L218
	}
L207:
	;
	v796 = F_tolower(m, v792)
	mBase = m.M
	v797 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v793))))
	v798 = F_tolower(m, v797)
	mBase = m.M
	goto L206
L208:
	;
	v766 = v341
	v767 = v761
	v768 = v764
	goto L211
L209:
	;
	v792 = int32(0)
	v793 = v761
	goto L207
L210:
	;
	v792 = v789 & int32(255)
	v793 = v788
	goto L207
L211:
	;
	v770 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v767))))
	if v770 == int32(0) {
		v788 = v767
		v789 = v768
		goto L210
	} else {
		goto L213
	}
L212:
	;
	v788 = v782
	v789 = int32(0)
	goto L210
L213:
	;
	v774 = v768 & int32(255)
	if v774 == v770 {
		goto L214
	} else {
		goto L215
	}
L214:
	;
	v781 = int32(1)
	v782 = v767 + v781
	v783 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v766)+1)))
	if v783 != 0 {
		v766 = v766 + v781
		v767 = v782
		v768 = v783
		goto L211
	} else {
		goto L217
	}
L215:
	;
	v776 = F_tolower(m, v774)
	mBase = m.M
	v777 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v767))))
	v778 = F_tolower(m, v777)
	mBase = m.M
	if v776 == v778 {
		goto L214
	} else {
		goto L216
	}
L216:
	;
	v780 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v766))))
	v788 = v767
	v789 = v780
	goto L210
L217:
	;
	goto L212
L218:
	;
	v801 = v302 + int32(1)
	if v252 <= v801 {
		goto L205
	} else {
		goto L219
	}
L219:
	;
	if v270 == int32(0) {
		goto L205
	} else {
		goto L220
	}
L220:
	;
	if v332 != 0 {
		goto L205
	} else {
		goto L221
	}
L221:
	;
	if v333 != 0 {
		goto L205
	} else {
		goto L222
	}
L222:
	;
	v805 = int32(1)
	v806 = int32(0)
	if v45 != 0 {
		goto L223
	} else {
		goto L224
	}
L223:
	;
	v809 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v813 = *(*int32)(unsafe.Add(mBase, uint32(v809+v338+int32(4))))
	v814 = F_longLatFromMemberOrReply(m, l0, v45, v813, v284)
	mBase = m.M
	v815 = m.ExcPending
	if v815 != 0 {
		goto L2
	} else {
		goto L225
	}
L224:
	;
	v1289 = v306
	v1297 = v314
	v1298 = v315
	v1299 = v316
	v1300 = v317
	v1301 = v318
	v1302 = v801
	v1311 = v328
	v1312 = v329
	v1313 = v330
	v1314 = v805
	v1315 = int32(0)
	v1316 = v806
	goto L60
L225:
	;
	if v814 != int32(-1) {
		v1289 = v306
		v1297 = v314
		v1298 = v315
		v1299 = v316
		v1300 = v317
		v1301 = v318
		v1302 = v801
		v1311 = v328
		v1312 = v329
		v1313 = v330
		v1314 = v805
		v1315 = int32(0)
		v1316 = v806
		goto L60
	} else {
		goto L226
	}
L226:
	;
	v818 = *(*int32)(unsafe.Add(mBase, uint32(v37)+672))
	if v818 != int32(3) {
		goto L1
	} else {
		goto L227
	}
L227:
	;
	v821 = *(*int32)(unsafe.Add(mBase, uint32(v37)+740))
	if v821 == int32(0) {
		goto L1
	} else {
		goto L228
	}
L228:
	;
	F_valkey_free(m, v821)
	mBase = m.M
	v825 = m.ExcPending
	if v825 != 0 {
		goto L2
	} else {
		goto L229
	}
L229:
	;
	goto L1
L230:
	;
	v896 = int32(_a_F_georadiusGeneric_18)
	v899 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v341))))
	if v899 != 0 {
		goto L259
	} else {
		goto L260
	}
L231:
	;
	if v862-v864 != 0 {
		goto L230
	} else {
		goto L243
	}
L232:
	;
	v862 = F_tolower(m, v858)
	mBase = m.M
	v863 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v859))))
	v864 = F_tolower(m, v863)
	mBase = m.M
	goto L231
L233:
	;
	v832 = v341
	v833 = v827
	v834 = v830
	goto L236
L234:
	;
	v858 = int32(0)
	v859 = v827
	goto L232
L235:
	;
	v858 = v855 & int32(255)
	v859 = v854
	goto L232
L236:
	;
	v836 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v833))))
	if v836 == int32(0) {
		v854 = v833
		v855 = v834
		goto L235
	} else {
		goto L238
	}
L237:
	;
	v854 = v848
	v855 = int32(0)
	goto L235
L238:
	;
	v840 = v834 & int32(255)
	if v840 == v836 {
		goto L239
	} else {
		goto L240
	}
L239:
	;
	v847 = int32(1)
	v848 = v833 + v847
	v849 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v832)+1)))
	if v849 != 0 {
		v832 = v832 + v847
		v833 = v848
		v834 = v849
		goto L236
	} else {
		goto L242
	}
L240:
	;
	v842 = F_tolower(m, v840)
	mBase = m.M
	v843 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v833))))
	v844 = F_tolower(m, v843)
	mBase = m.M
	if v842 == v844 {
		goto L239
	} else {
		goto L241
	}
L241:
	;
	v846 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v832))))
	v854 = v833
	v855 = v846
	goto L235
L242:
	;
	goto L237
L243:
	;
	v867 = v302 + int32(2)
	if v252 <= v867 {
		goto L230
	} else {
		goto L244
	}
L244:
	;
	if v270 == int32(0) {
		goto L230
	} else {
		goto L245
	}
L245:
	;
	if v331 != 0 {
		goto L230
	} else {
		goto L246
	}
L246:
	;
	if v333 != 0 {
		goto L230
	} else {
		goto L247
	}
L247:
	;
	v871 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v872 = int32(2)
	v880 = F_extractLongLatOrReply(m, l0, v871+v243<<(uint(v872)%32)+v302<<(uint(v872)%32)+int32(4), v284)
	mBase = m.M
	v881 = m.ExcPending
	if v881 != 0 {
		goto L2
	} else {
		goto L249
	}
L248:
	;
	v887 = *(*int32)(unsafe.Add(mBase, uint32(v37)+672))
	if v887 != int32(3) {
		goto L1
	} else {
		goto L251
	}
L249:
	;
	if v880 == int32(-1) {
		goto L248
	} else {
		goto L250
	}
L250:
	;
	v885 = int32(0)
	v1289 = v306
	v1297 = v314
	v1298 = v315
	v1299 = v316
	v1300 = v317
	v1301 = v318
	v1302 = v867
	v1311 = v328
	v1312 = v329
	v1313 = v330
	v1314 = v885
	v1315 = int32(1)
	v1316 = v885
	goto L60
L251:
	;
	v890 = *(*int32)(unsafe.Add(mBase, uint32(v37)+740))
	if v890 == int32(0) {
		goto L1
	} else {
		goto L252
	}
L252:
	;
	F_valkey_free(m, v890)
	mBase = m.M
	v894 = m.ExcPending
	if v894 != 0 {
		goto L2
	} else {
		goto L253
	}
L253:
	;
	goto L1
L254:
	;
	v1276 = int32(1)
	v1280 = int32(0)
	v1289 = v306
	v1297 = v314
	v1298 = v315
	v1299 = v1280
	v1300 = v1280
	v1301 = v318
	v1302 = v302 + v1260 + v1276
	v1311 = v328
	v1312 = v329
	v1313 = v330
	v1314 = v1280
	v1315 = v1280
	v1316 = v1276
	goto L60
L255:
	;
	v1237 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v37)+672)) = v1237
	v1240 = int32(0)
	v1289 = v306
	v1297 = v314
	v1298 = v315
	v1299 = v1240
	v1300 = v1237
	v1301 = v318
	v1302 = v936
	v1311 = v328
	v1312 = v329
	v1313 = v330
	v1314 = v331
	v1315 = v332
	v1316 = v1240
	goto L60
L256:
	;
	v962 = int32(_a_F_georadiusGeneric_19)
	v965 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v341))))
	if v965 != 0 {
		goto L282
	} else {
		goto L283
	}
L257:
	;
	if v931-v933 != 0 {
		goto L256
	} else {
		goto L269
	}
L258:
	;
	v931 = F_tolower(m, v927)
	mBase = m.M
	v932 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v928))))
	v933 = F_tolower(m, v932)
	mBase = m.M
	goto L257
L259:
	;
	v901 = v341
	v902 = v896
	v903 = v899
	goto L262
L260:
	;
	v927 = int32(0)
	v928 = v896
	goto L258
L261:
	;
	v927 = v924 & int32(255)
	v928 = v923
	goto L258
L262:
	;
	v905 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v902))))
	if v905 == int32(0) {
		v923 = v902
		v924 = v903
		goto L261
	} else {
		goto L264
	}
L263:
	;
	v923 = v917
	v924 = int32(0)
	goto L261
L264:
	;
	v909 = v903 & int32(255)
	if v909 == v905 {
		goto L265
	} else {
		goto L266
	}
L265:
	;
	v916 = int32(1)
	v917 = v902 + v916
	v918 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v901)+1)))
	if v918 != 0 {
		v901 = v901 + v916
		v902 = v917
		v903 = v918
		goto L262
	} else {
		goto L268
	}
L266:
	;
	v911 = F_tolower(m, v909)
	mBase = m.M
	v912 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v902))))
	v913 = F_tolower(m, v912)
	mBase = m.M
	if v911 == v913 {
		goto L265
	} else {
		goto L267
	}
L267:
	;
	v915 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v901))))
	v923 = v902
	v924 = v915
	goto L261
L268:
	;
	goto L263
L269:
	;
	v936 = v302 + int32(2)
	if v252 <= v936 {
		goto L256
	} else {
		goto L270
	}
L270:
	;
	if v270 == int32(0) {
		goto L256
	} else {
		goto L271
	}
L271:
	;
	if v316 != 0 {
		goto L256
	} else {
		goto L272
	}
L272:
	;
	if v333 != 0 {
		goto L256
	} else {
		goto L273
	}
L273:
	;
	v940 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v941 = int32(2)
	v949 = F_extractDistanceOrReply(m, l0, v940+v243<<(uint(v941)%32)+v302<<(uint(v941)%32)+int32(4), v280, v276)
	mBase = m.M
	v950 = m.ExcPending
	if v950 != 0 {
		goto L2
	} else {
		goto L274
	}
L274:
	;
	if v949 == int32(0) {
		goto L255
	} else {
		goto L275
	}
L275:
	;
	v953 = *(*int32)(unsafe.Add(mBase, uint32(v37)+672))
	if v953 != int32(3) {
		goto L1
	} else {
		goto L276
	}
L276:
	;
	v956 = *(*int32)(unsafe.Add(mBase, uint32(v37)+740))
	if v956 == int32(0) {
		goto L1
	} else {
		goto L277
	}
L277:
	;
	F_valkey_free(m, v956)
	mBase = m.M
	v960 = m.ExcPending
	if v960 != 0 {
		goto L2
	} else {
		goto L278
	}
L278:
	;
	goto L1
L279:
	;
	v1033 = int32(_a_F_georadiusGeneric_20)
	v1036 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v341))))
	if v1036 != 0 {
		goto L306
	} else {
		goto L307
	}
L280:
	;
	if v997-v999 != 0 {
		goto L279
	} else {
		goto L292
	}
L281:
	;
	v997 = F_tolower(m, v993)
	mBase = m.M
	v998 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v994))))
	v999 = F_tolower(m, v998)
	mBase = m.M
	goto L280
L282:
	;
	v967 = v341
	v968 = v962
	v969 = v965
	goto L285
L283:
	;
	v993 = int32(0)
	v994 = v962
	goto L281
L284:
	;
	v993 = v990 & int32(255)
	v994 = v989
	goto L281
L285:
	;
	v971 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v968))))
	if v971 == int32(0) {
		v989 = v968
		v990 = v969
		goto L284
	} else {
		goto L287
	}
L286:
	;
	v989 = v983
	v990 = int32(0)
	goto L284
L287:
	;
	v975 = v969 & int32(255)
	if v975 == v971 {
		goto L288
	} else {
		goto L289
	}
L288:
	;
	v982 = int32(1)
	v983 = v968 + v982
	v984 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v967)+1)))
	if v984 != 0 {
		v967 = v967 + v982
		v968 = v983
		v969 = v984
		goto L285
	} else {
		goto L291
	}
L289:
	;
	v977 = F_tolower(m, v975)
	mBase = m.M
	v978 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v968))))
	v979 = F_tolower(m, v978)
	mBase = m.M
	if v977 == v979 {
		goto L288
	} else {
		goto L290
	}
L290:
	;
	v981 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v967))))
	v989 = v968
	v990 = v981
	goto L284
L291:
	;
	goto L286
L292:
	;
	v1002 = v302 + int32(3)
	if v252 <= v1002 {
		goto L279
	} else {
		goto L293
	}
L293:
	;
	if v270 == int32(0) {
		goto L279
	} else {
		goto L294
	}
L294:
	;
	if v317 != 0 {
		goto L279
	} else {
		goto L295
	}
L295:
	;
	if v333 != 0 {
		goto L279
	} else {
		goto L296
	}
L296:
	;
	v1006 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v1007 = int32(2)
	v1015 = F_extractBoxOrReply(m, l0, v1006+v243<<(uint(v1007)%32)+v302<<(uint(v1007)%32)+int32(4), v280, v37+int32(744), v276)
	mBase = m.M
	v1016 = m.ExcPending
	if v1016 != 0 {
		goto L2
	} else {
		goto L298
	}
L297:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v37)+672)) = int32(2)
	v1030 = int32(0)
	v1289 = v306
	v1297 = v314
	v1298 = v315
	v1299 = int32(1)
	v1300 = v1030
	v1301 = v318
	v1302 = v1002
	v1311 = v328
	v1312 = v329
	v1313 = v330
	v1314 = v331
	v1315 = v332
	v1316 = v1030
	goto L60
L298:
	;
	if v1015 == int32(0) {
		goto L297
	} else {
		goto L299
	}
L299:
	;
	v1019 = *(*int32)(unsafe.Add(mBase, uint32(v37)+672))
	if v1019 != int32(3) {
		goto L1
	} else {
		goto L300
	}
L300:
	;
	v1022 = *(*int32)(unsafe.Add(mBase, uint32(v37)+740))
	if v1022 == int32(0) {
		goto L1
	} else {
		goto L301
	}
L301:
	;
	F_valkey_free(m, v1022)
	mBase = m.M
	v1026 = m.ExcPending
	if v1026 != 0 {
		goto L2
	} else {
		goto L302
	}
L302:
	;
	goto L1
L303:
	;
	v1226 = *(*int32)(unsafe.Add(mBase, _c_F_georadiusGeneric[0]))
	F_addReplyErrorObject(m, l0, v1226)
	mBase = m.M
	v1228 = m.ExcPending
	if v1228 != 0 {
		goto L2
	} else {
		goto L346
	}
L304:
	;
	if v1068-v1070 != 0 {
		goto L303
	} else {
		goto L316
	}
L305:
	;
	v1068 = F_tolower(m, v1064)
	mBase = m.M
	v1069 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1065))))
	v1070 = F_tolower(m, v1069)
	mBase = m.M
	goto L304
L306:
	;
	v1038 = v341
	v1039 = v1033
	v1040 = v1036
	goto L309
L307:
	;
	v1064 = int32(0)
	v1065 = v1033
	goto L305
L308:
	;
	v1064 = v1061 & int32(255)
	v1065 = v1060
	goto L305
L309:
	;
	v1042 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1039))))
	if v1042 == int32(0) {
		v1060 = v1039
		v1061 = v1040
		goto L308
	} else {
		goto L311
	}
L310:
	;
	v1060 = v1054
	v1061 = int32(0)
	goto L308
L311:
	;
	v1046 = v1040 & int32(255)
	if v1046 == v1042 {
		goto L312
	} else {
		goto L313
	}
L312:
	;
	v1053 = int32(1)
	v1054 = v1039 + v1053
	v1055 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1038)+1)))
	if v1055 != 0 {
		v1038 = v1038 + v1053
		v1039 = v1054
		v1040 = v1055
		goto L309
	} else {
		goto L315
	}
L313:
	;
	v1048 = F_tolower(m, v1046)
	mBase = m.M
	v1049 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1039))))
	v1050 = F_tolower(m, v1049)
	mBase = m.M
	if v1048 == v1050 {
		goto L312
	} else {
		goto L314
	}
L314:
	;
	v1052 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1038))))
	v1060 = v1039
	v1061 = v1052
	goto L308
L315:
	;
	goto L310
L316:
	;
	if v252 <= v302+int32(2) {
		goto L303
	} else {
		goto L317
	}
L317:
	;
	if v270 == int32(0) {
		goto L303
	} else {
		goto L318
	}
L318:
	;
	if v317 != 0 {
		goto L303
	} else {
		goto L319
	}
L319:
	;
	if v316 != 0 {
		goto L303
	} else {
		goto L320
	}
L320:
	;
	if v331 != 0 {
		goto L303
	} else {
		goto L321
	}
L321:
	;
	if v332 != 0 {
		goto L303
	} else {
		goto L322
	}
L322:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v37)+136)) = int32(0)
	v1079 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v1083 = *(*int32)(unsafe.Add(mBase, uint32(v1079+v338+int32(4))))
	v1087 = F_getIntFromObjectOrReply(m, l0, v1083, v37+int32(136), int32(_a_F_georadiusGeneric_21))
	mBase = m.M
	v1088 = m.ExcPending
	if v1088 != 0 {
		goto L2
	} else {
		goto L325
	}
L323:
	;
	v1220 = *(*int32)(unsafe.Add(mBase, uint32(v37)+740))
	if v1220 == int32(0) {
		goto L1
	} else {
		goto L344
	}
L324:
	;
	v1094 = *(*int32)(unsafe.Add(mBase, uint32(v37)+136))
	if v1094 < int32(3) {
		goto L329
	} else {
		goto L330
	}
L325:
	;
	if v1087 == int32(0) {
		goto L324
	} else {
		goto L326
	}
L326:
	;
	v1091 = *(*int32)(unsafe.Add(mBase, uint32(v37)+672))
	if v1091 == int32(3) {
		goto L323
	} else {
		goto L327
	}
L327:
	;
	goto L1
L328:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v37)+696)) = int64(4607182418800017408)
	*(*int32)(unsafe.Add(mBase, uint32(v37)+736)) = v1094
	v1112 = F_valkey_malloc(m, v1094<<(uint(int32(4))%32))
	mBase = m.M
	v1113 = m.ExcPending
	if v1113 != 0 {
		goto L2
	} else {
		goto L334
	}
L329:
	;
	F_addReplyError(m, l0, int32(_a_F_georadiusGeneric_22))
	mBase = m.M
	v1103 = m.ExcPending
	if v1103 != 0 {
		goto L2
	} else {
		goto L332
	}
L330:
	;
	v1099 = base.I32_div_s(v252+int32(-2)-v302, int32(2))
	if v1094 <= v1099 {
		goto L328
	} else {
		goto L331
	}
L331:
	;
	goto L329
L332:
	;
	v1104 = *(*int32)(unsafe.Add(mBase, uint32(v37)+672))
	if v1104 == int32(3) {
		goto L323
	} else {
		goto L333
	}
L333:
	;
	goto L1
L334:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v37)+672)) = int32(3)
	*(*int32)(unsafe.Add(mBase, uint32(v37)+740)) = v1112
	v1117 = int32(0)
	v1118 = *(*int32)(unsafe.Add(mBase, uint32(v37)+136))
	if v1117 < v1118 {
		goto L335
	} else {
		goto L336
	}
L335:
	;
	v1132 = v1117
	goto L338
L336:
	;
	v1260 = v1118 << (uint(int32(1)) % 32)
	goto L254
L337:
	;
	v1183 = *(*int32)(unsafe.Add(mBase, uint32(v37)+672))
	if v1183 != int32(3) {
		goto L1
	} else {
		goto L343
	}
L338:
	;
	v1157 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v1158 = int32(2)
	v1169 = *(*int32)(unsafe.Add(mBase, uint32(v37)+740))
	v1173 = F_extractLongLatOrReply(m, l0, v1157+v243<<(uint(v1158)%32)+v302<<(uint(v1158)%32)+v1132<<(uint(v1158)%32)+int32(8), v1169+v1132<<(uint(int32(3))%32))
	mBase = m.M
	v1174 = m.ExcPending
	if v1174 != 0 {
		goto L2
	} else {
		goto L340
	}
L340:
	;
	if v1173 == int32(-1) {
		goto L337
	} else {
		goto L341
	}
L341:
	;
	v1178 = v1132 + int32(2)
	v1179 = *(*int32)(unsafe.Add(mBase, uint32(v37)+136))
	v1181 = v1179 << (uint(int32(1)) % 32)
	if v1181 <= v1178 {
		v1260 = v1181
		goto L254
	} else {
		goto L342
	}
L342:
	;
	v1132 = v1178
	goto L338
L343:
	;
	goto L323
L344:
	;
	F_valkey_free(m, v1220)
	mBase = m.M
	v1224 = m.ExcPending
	if v1224 != 0 {
		goto L2
	} else {
		goto L345
	}
L345:
	;
	goto L1
L346:
	;
	v1229 = *(*int32)(unsafe.Add(mBase, uint32(v37)+672))
	if v1229 != int32(3) {
		goto L1
	} else {
		goto L347
	}
L347:
	;
	v1232 = *(*int32)(unsafe.Add(mBase, uint32(v37)+740))
	if v1232 == int32(0) {
		goto L1
	} else {
		goto L348
	}
L348:
	;
	F_valkey_free(m, v1232)
	mBase = m.M
	v1236 = m.ExcPending
	if v1236 != 0 {
		goto L2
	} else {
		goto L349
	}
L349:
	;
	goto L1
L350:
	;
	goto L59
L351:
	;
	v1398 = l2 & int32(8)
	v1400 = base.B2i32(v1398 == int32(0))
	if (v1400|v1349|v1345|v1338)&int32(1) != 0 {
		goto L363
	} else {
		goto L364
	}
L352:
	;
	if (v1346|v1355)&int32(1) != 0 {
		goto L353
	} else {
		goto L354
	}
L353:
	;
	if l2&int32(16) != 0 {
		goto L356
	} else {
		goto L357
	}
L354:
	;
	if v1354 == int32(0) {
		goto L351
	} else {
		goto L355
	}
L355:
	;
	goto L353
L356:
	;
	v1382 = int32(_a_F_georadiusGeneric_23)
	goto L358
L357:
	;
	v1382 = int32(_a_F_georadiusGeneric_24)
	goto L358
L358:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v37)+64)) = v1382
	F_addReplyErrorFormat(m, l0, int32(_a_F_georadiusGeneric_25), v37+int32(64))
	mBase = m.M
	v1388 = m.ExcPending
	if v1388 != 0 {
		goto L2
	} else {
		goto L359
	}
L359:
	;
	v1389 = *(*int32)(unsafe.Add(mBase, uint32(v37)+672))
	if v1389 != int32(3) {
		goto L1
	} else {
		goto L360
	}
L360:
	;
	v1392 = *(*int32)(unsafe.Add(mBase, uint32(v37)+740))
	if v1392 == int32(0) {
		goto L1
	} else {
		goto L361
	}
L361:
	;
	F_valkey_free(m, v1392)
	mBase = m.M
	v1396 = m.ExcPending
	if v1396 != 0 {
		goto L2
	} else {
		goto L362
	}
L362:
	;
	goto L1
L363:
	;
	if (v1400|v1353|v1352|v1338)&int32(1) != 0 {
		goto L369
	} else {
		goto L370
	}
L364:
	;
	v1406 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v1407 = *(*int32)(unsafe.Add(mBase, uint32(v1406)))
	v1408 = F_objectGetVal(m, v1407)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v37)+48)) = v1408
	F_addReplyErrorFormat(m, l0, int32(_a_F_georadiusGeneric_26), v37+int32(48))
	mBase = m.M
	v1414 = m.ExcPending
	if v1414 != 0 {
		goto L2
	} else {
		goto L365
	}
L365:
	;
	v1415 = *(*int32)(unsafe.Add(mBase, uint32(v37)+672))
	if v1415 != int32(3) {
		goto L1
	} else {
		goto L366
	}
L366:
	;
	v1418 = *(*int32)(unsafe.Add(mBase, uint32(v37)+740))
	if v1418 == int32(0) {
		goto L1
	} else {
		goto L367
	}
L367:
	;
	F_valkey_free(m, v1418)
	mBase = m.M
	v1422 = m.ExcPending
	if v1422 != 0 {
		goto L2
	} else {
		goto L368
	}
L368:
	;
	goto L1
L369:
	;
	v1445 = *(*int64)(unsafe.Add(mBase, uint32(v37)+664))
	if v1351 == int32(0) {
		goto L375
	} else {
		goto L376
	}
L370:
	;
	v1428 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v1429 = *(*int32)(unsafe.Add(mBase, uint32(v1428)))
	v1430 = F_objectGetVal(m, v1429)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v37)+32)) = v1430
	F_addReplyErrorFormat(m, l0, int32(_a_F_georadiusGeneric_27), v37+int32(32))
	mBase = m.M
	v1436 = m.ExcPending
	if v1436 != 0 {
		goto L2
	} else {
		goto L371
	}
L371:
	;
	v1437 = *(*int32)(unsafe.Add(mBase, uint32(v37)+672))
	if v1437 != int32(3) {
		goto L1
	} else {
		goto L372
	}
L372:
	;
	v1440 = *(*int32)(unsafe.Add(mBase, uint32(v37)+740))
	if v1440 == int32(0) {
		goto L1
	} else {
		goto L373
	}
L373:
	;
	F_valkey_free(m, v1440)
	mBase = m.M
	v1444 = m.ExcPending
	if v1444 != 0 {
		goto L2
	} else {
		goto L374
	}
L374:
	;
	goto L1
L375:
	;
	if v45 != 0 {
		goto L382
	} else {
		goto L383
	}
L376:
	;
	if v1445 != int64(0) {
		goto L375
	} else {
		goto L377
	}
L377:
	;
	F_addReplyError(m, l0, int32(_a_F_georadiusGeneric_28))
	mBase = m.M
	v1452 = m.ExcPending
	if v1452 != 0 {
		goto L2
	} else {
		goto L378
	}
L378:
	;
	v1453 = *(*int32)(unsafe.Add(mBase, uint32(v37)+672))
	if v1453 != int32(3) {
		goto L1
	} else {
		goto L379
	}
L379:
	;
	v1456 = *(*int32)(unsafe.Add(mBase, uint32(v37)+740))
	if v1456 == int32(0) {
		goto L1
	} else {
		goto L380
	}
L380:
	;
	F_valkey_free(m, v1456)
	mBase = m.M
	v1460 = m.ExcPending
	if v1460 != 0 {
		goto L2
	} else {
		goto L381
	}
L381:
	;
	goto L1
L382:
	;
	v1501 = v37 + int32(472)
	v1503 = v37 + int32(672)
	F_geohashCalculateAreasByShapeWGS84(m, v1501, v1503)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v37)+136)) = v37 + int32(152)
	*(*int64)(unsafe.Add(mBase, uint32(v37)+140)) = int64(8)
	v1516 = *(*int32)(unsafe.Add(mBase, uint32(v37)+664))
	if v1351 != 0 {
		goto L396
	} else {
		goto L397
	}
L383:
	;
	if v1342 != 0 {
		goto L385
	} else {
		goto L386
	}
L384:
	;
	v1489 = *(*int32)(unsafe.Add(mBase, uint32(v1488)))
	F_addReply(m, l0, v1489)
	mBase = m.M
	v1491 = m.ExcPending
	if v1491 != 0 {
		goto L2
	} else {
		goto L392
	}
L385:
	;
	v1464 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v1465 = F_dbDelete(m, v1464, v1342)
	mBase = m.M
	v1466 = m.ExcPending
	if v1466 != 0 {
		goto L2
	} else {
		goto L388
	}
L386:
	;
	v1488 = int32(_a_F_georadiusGeneric_29)
	goto L384
L387:
	;
	v1470 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	F_signalModifiedKey(m, l0, v1470, v1342)
	mBase = m.M
	v1472 = m.ExcPending
	if v1472 != 0 {
		goto L2
	} else {
		goto L390
	}
L388:
	;
	if v1465 != 0 {
		goto L387
	} else {
		goto L389
	}
L389:
	;
	v1488 = int32(_a_F_georadiusGeneric_30)
	goto L384
L390:
	;
	v1475 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v1476 = *(*int32)(unsafe.Add(mBase, uint32(v1475)+28))
	F_notifyKeyspaceEvent(m, int32(4), int32(_a_F_georadiusGeneric_31), v1342, v1476)
	mBase = m.M
	v1478 = m.ExcPending
	if v1478 != 0 {
		goto L2
	} else {
		goto L391
	}
L391:
	;
	v1479 = int32(_a_F_georadiusGeneric_32)
	v1481 = *(*int64)(unsafe.Add(mBase, _c_F_georadiusGeneric[1]))
	*(*int64)(unsafe.Add(mBase, _c_F_georadiusGeneric[1])) = v1481 + int64(1)
	v1488 = int32(_a_F_georadiusGeneric_30)
	goto L384
L392:
	;
	v1492 = *(*int32)(unsafe.Add(mBase, uint32(v37)+672))
	if v1492 != int32(3) {
		goto L1
	} else {
		goto L393
	}
L393:
	;
	v1495 = *(*int32)(unsafe.Add(mBase, uint32(v37)+740))
	if v1495 == int32(0) {
		goto L1
	} else {
		goto L394
	}
L394:
	;
	F_valkey_free(m, v1495)
	mBase = m.M
	v1499 = m.ExcPending
	if v1499 != 0 {
		goto L2
	} else {
		goto L395
	}
L395:
	;
	goto L1
L396:
	;
	v1518 = v1516
	goto L398
L397:
	;
	v1518 = int32(0)
	goto L398
L398:
	;
	v1519 = F_membersOfAllNeighbors(m, v45, v1501, v1503, v37+int32(136), v1518)
	mBase = m.M
	v1520 = m.ExcPending
	if v1520 != 0 {
		goto L2
	} else {
		goto L399
	}
L399:
	;
	v1521 = *(*int32)(unsafe.Add(mBase, uint32(v37)+144))
	if v1342 != 0 {
		goto L402
	} else {
		goto L403
	}
L400:
	;
	F__serverAssert(m, int32(_a_F_georadiusGeneric_33), int32(_a_F_georadiusGeneric_34), int32(844))
	mBase = m.M
	v2161 = m.ExcPending
	if v2161 != 0 {
		goto L2
	} else {
		goto L521
	}
L401:
	;
	v2152 = *(*int32)(unsafe.Add(mBase, uint32(v37)+740))
	if v2152 == int32(0) {
		goto L1
	} else {
		goto L519
	}
L402:
	;
	v1533 = *(*int64)(unsafe.Add(mBase, uint32(v37)+664))
	if base.I64_extend_i32_s(v1521) < v1533 {
		goto L408
	} else {
		goto L409
	}
L403:
	;
	if v1521 != 0 {
		goto L402
	} else {
		goto L404
	}
L404:
	;
	v1523 = *(*int32)(unsafe.Add(mBase, _c_F_georadiusGeneric[2]))
	F_addReply(m, l0, v1523)
	mBase = m.M
	v1525 = m.ExcPending
	if v1525 != 0 {
		goto L2
	} else {
		goto L405
	}
L405:
	;
	F_geoArrayCleanup(m, v37+int32(136))
	mBase = m.M
	v1529 = m.ExcPending
	if v1529 != 0 {
		goto L2
	} else {
		goto L406
	}
L406:
	;
	v1530 = *(*int32)(unsafe.Add(mBase, uint32(v37)+672))
	if v1530 == int32(3) {
		goto L401
	} else {
		goto L407
	}
L407:
	;
	goto L1
L408:
	;
	v1537 = v1521
	goto L410
L409:
	;
	v1537 = base.I32_wrap_i64(v1533)
	goto L410
L410:
	;
	if v1533 == int64(0) {
		goto L411
	} else {
		goto L412
	}
L411:
	;
	v1540 = v1521
	goto L413
L412:
	;
	v1540 = v1537
	goto L413
L413:
	;
	if v1350|v1351 != 0 {
		goto L417
	} else {
		goto L418
	}
L414:
	;
	if v1342 != 0 {
		goto L428
	} else {
		goto L429
	}
L415:
	;
	v1550 = *(*int32)(unsafe.Add(mBase, uint32(v37)+136))
	if v1540 != v1521 {
		goto L423
	} else {
		goto L424
	}
L416:
	;
	v1549 = int32(538)
	goto L415
L417:
	;
	v1544 = v1350
	goto L419
L418:
	;
	v1544 = int32(1)
	goto L419
L419:
	;
	if v1445 != int64(0) {
		goto L420
	} else {
		goto L421
	}
L420:
	;
	v1547 = v1544
	goto L422
L421:
	;
	v1547 = v1350
	goto L422
L422:
	;
	switch v1547 {
	case 0:
		goto L414
	case 1:
		v1549 = int32(537)
		goto L415
	default:
		goto L416
	}
L423:
	;
	F_pqsort(m, v1550, v1521, int32(40), v1549, int32(0), v1540+int32(-1))
	mBase = m.M
	v1560 = m.ExcPending
	if v1560 != 0 {
		goto L2
	} else {
		goto L426
	}
L424:
	;
	F_qsort(m, v1550, v1521, int32(40), v1549)
	mBase = m.M
	v1554 = m.ExcPending
	if v1554 != 0 {
		goto L2
	} else {
		goto L425
	}
L425:
	;
	goto L414
L426:
	;
	goto L414
L427:
	;
	F_geoArrayCleanup(m, v37+int32(136))
	mBase = m.M
	v2114 = m.ExcPending
	if v2114 != 0 {
		goto L2
	} else {
		goto L517
	}
L428:
	;
	if v1540 == int32(0) {
		goto L483
	} else {
		goto L484
	}
L429:
	;
	F_addReplyArrayLen(m, l0, v1540)
	mBase = m.M
	v1564 = m.ExcPending
	if v1564 != 0 {
		goto L2
	} else {
		goto L430
	}
L430:
	;
	if v1540 < int32(1) {
		goto L427
	} else {
		goto L431
	}
L431:
	;
	if v1346 != 0 {
		goto L432
	} else {
		goto L433
	}
L432:
	;
	v1569 = int32(2)
	goto L434
L433:
	;
	v1569 = int32(1)
	goto L434
L434:
	;
	if v1354 != 0 {
		goto L435
	} else {
		goto L436
	}
L435:
	;
	v1570 = v1569
	goto L437
L436:
	;
	v1570 = v1346
	goto L437
L437:
	;
	v1571 = v1570 + v1355
	v1577 = int32(0)
	goto L438
L438:
	;
	v1609 = *(*int32)(unsafe.Add(mBase, uint32(v37)+136))
	v1612 = v1609 + v1577*int32(40)
	v1613 = *(*float64)(unsafe.Add(mBase, uint32(v1612)+16))
	v1614 = *(*float64)(unsafe.Add(mBase, uint32(v37)+696))
	*(*float64)(unsafe.Add(mBase, uint32(v1612)+16)) = base.F64_div(v1613, v1614)
	if v1571 == int32(0) {
		goto L440
	} else {
		goto L441
	}
L440:
	;
	v1621 = *(*int32)(unsafe.Add(mBase, uint32(v1612)+32))
	F_addReplyBulkSds(m, l0, v1621)
	mBase = m.M
	v1623 = m.ExcPending
	if v1623 != 0 {
		goto L2
	} else {
		goto L443
	}
L441:
	;
	F_addReplyArrayLen(m, l0, v1571+int32(1))
	mBase = m.M
	v1620 = m.ExcPending
	if v1620 != 0 {
		goto L2
	} else {
		goto L442
	}
L442:
	;
	goto L440
L443:
	;
	v1624 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v1612)+32)) = v1624
	if v1346 == v1624 {
		goto L444
	} else {
		goto L445
	}
L444:
	;
	if v1355 == int32(0) {
		goto L447
	} else {
		goto L448
	}
L445:
	;
	v1629 = v37 + int32(752)
	v1633 = *(*float64)(unsafe.Add(mBase, uint32(v1612)+16))
	v1635 = F_fixedpoint_d2string(m, v1629, int32(128), v1633, int32(4))
	mBase = m.M
	F_addReplyBulkCBuffer(m, l0, v1629, v1635)
	mBase = m.M
	v1637 = m.ExcPending
	if v1637 != 0 {
		goto L2
	} else {
		goto L446
	}
L446:
	;
	goto L444
L447:
	;
	if v1354 == int32(0) {
		goto L453
	} else {
		goto L454
	}
L448:
	;
	v1640 = *(*float64)(unsafe.Add(mBase, uint32(v1612)+24))
	if base.F64_lt(base.F64_abs(v1640), float64(9.223372036854776e+18)) == int32(0) {
		goto L450
	} else {
		goto L451
	}
L449:
	;
	F_addReplyLongLong(m, l0, v1648)
	mBase = m.M
	v1650 = m.ExcPending
	if v1650 != 0 {
		goto L2
	} else {
		goto L452
	}
L450:
	;
	v1648 = int64(-9223372036854775807 - 1)
	goto L449
L451:
	;
	v1646 = base.I64_trunc_f64_s(v1640)
	v1648 = v1646
	goto L449
L452:
	;
	goto L447
L453:
	;
	v1833 = v1577 + int32(1)
	if v1833 != v1540 {
		v1577 = v1833
		goto L438
	} else {
		goto L480
	}
L454:
	;
	F_addReplyArrayLen(m, l0, int32(2))
	mBase = m.M
	v1657 = m.ExcPending
	if v1657 != 0 {
		goto L2
	} else {
		goto L455
	}
L455:
	;
	v1658 = int32(16)
	v1659 = v37 + v1658
	v1660 = *(*float64)(unsafe.Add(mBase, uint32(v1612)))
	v1667 = m.G0
	v1669 = v1667 - v1658
	m.G0 = v1669
	v1671 = base.I64_reinterpret_f64(v1660)
	v1673 = v1671 & int64(4503599627370495)
	v1677 = int64(base.Ui64(v1671)>>(uint(int64(52))%64)) & int64(2047)
	if v1677 == int64(0) {
		goto L458
	} else {
		goto L459
	}
L456:
	;
	v1739 = *(*int64)(unsafe.Add(mBase, uint32(v37)+16))
	v1744 = *(*int64)(unsafe.Add(mBase, uint32(v37+int32(24))))
	F_addReplyHumanLongDouble(m, l0, v1739, v1744)
	mBase = m.M
	v1746 = m.ExcPending
	if v1746 != 0 {
		goto L2
	} else {
		goto L467
	}
L457:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v1659))) = v1724
	*(*int64)(unsafe.Add(mBase, uint32(v1659)+8)) = v1725<<(uint(int64(48))%64) | v1671&int64(-9223372036854775807-1) | v1726
	m.G0 = v1669 + int32(16)
	goto L456
L458:
	;
	if base.B2i32(v1673 == int64(0)) == int32(0) {
		goto L462
	} else {
		goto L463
	}
L459:
	;
	if v1677 == int64(2047) {
		goto L460
	} else {
		goto L461
	}
L460:
	;
	v1724 = v1673 << (uint(int64(60)) % 64)
	v1725 = int64(32767)
	v1726 = int64(base.Ui64(v1673) >> (uint(int64(4)) % 64))
	goto L457
L461:
	;
	v1724 = v1673 << (uint(int64(60)) % 64)
	v1725 = v1677 + int64(15360)
	v1726 = int64(base.Ui64(v1673) >> (uint(int64(4)) % 64))
	goto L457
L462:
	;
	if base.Ui64(v1673) < base.Ui64(int64(4294967296)) {
		goto L464
	} else {
		goto L465
	}
L463:
	;
	v1697 = int64(0)
	v1724 = v1697
	v1725 = v1697
	v1726 = v1697
	goto L457
L464:
	;
	v1711 = base.I32_clz(base.I32_wrap_i64(v1671)) | int32(32)
	goto L466
L465:
	;
	v1711 = base.I32_clz(base.I32_wrap_i64(int64(base.Ui64(v1673) >> (uint(int64(32)) % 64))))
	goto L466
L466:
	;
	F___ashlti3(m, v1669, v1673, int64(0), v1711+int32(49))
	mBase = m.M
	v1720 = *(*int64)(unsafe.Add(mBase, uint32(v1669+int32(8))))
	v1723 = *(*int64)(unsafe.Add(mBase, uint32(v1669)))
	v1724 = v1723
	v1725 = base.I64_extend_i32_u(int32(15372) - v1711)
	v1726 = v1720 ^ int64(281474976710656)
	goto L457
L467:
	;
	v1747 = *(*float64)(unsafe.Add(mBase, uint32(v1612)+8))
	v1754 = m.G0
	v1756 = v1754 - int32(16)
	m.G0 = v1756
	v1758 = base.I64_reinterpret_f64(v1747)
	v1760 = v1758 & int64(4503599627370495)
	v1764 = int64(base.Ui64(v1758)>>(uint(int64(52))%64)) & int64(2047)
	if v1764 == int64(0) {
		goto L470
	} else {
		goto L471
	}
L468:
	;
	v1826 = *(*int64)(unsafe.Add(mBase, uint32(v37)))
	v1829 = *(*int64)(unsafe.Add(mBase, uint32(v37+int32(8))))
	F_addReplyHumanLongDouble(m, l0, v1826, v1829)
	mBase = m.M
	v1831 = m.ExcPending
	if v1831 != 0 {
		goto L2
	} else {
		goto L479
	}
L469:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v37))) = v1811
	*(*int64)(unsafe.Add(mBase, uint32(v37)+8)) = v1812<<(uint(int64(48))%64) | v1758&int64(-9223372036854775807-1) | v1813
	m.G0 = v1756 + int32(16)
	goto L468
L470:
	;
	if base.B2i32(v1760 == int64(0)) == int32(0) {
		goto L474
	} else {
		goto L475
	}
L471:
	;
	if v1764 == int64(2047) {
		goto L472
	} else {
		goto L473
	}
L472:
	;
	v1811 = v1760 << (uint(int64(60)) % 64)
	v1812 = int64(32767)
	v1813 = int64(base.Ui64(v1760) >> (uint(int64(4)) % 64))
	goto L469
L473:
	;
	v1811 = v1760 << (uint(int64(60)) % 64)
	v1812 = v1764 + int64(15360)
	v1813 = int64(base.Ui64(v1760) >> (uint(int64(4)) % 64))
	goto L469
L474:
	;
	if base.Ui64(v1760) < base.Ui64(int64(4294967296)) {
		goto L476
	} else {
		goto L477
	}
L475:
	;
	v1784 = int64(0)
	v1811 = v1784
	v1812 = v1784
	v1813 = v1784
	goto L469
L476:
	;
	v1798 = base.I32_clz(base.I32_wrap_i64(v1758)) | int32(32)
	goto L478
L477:
	;
	v1798 = base.I32_clz(base.I32_wrap_i64(int64(base.Ui64(v1760) >> (uint(int64(32)) % 64))))
	goto L478
L478:
	;
	F___ashlti3(m, v1756, v1760, int64(0), v1798+int32(49))
	mBase = m.M
	v1807 = *(*int64)(unsafe.Add(mBase, uint32(v1756+int32(8))))
	v1810 = *(*int64)(unsafe.Add(mBase, uint32(v1756)))
	v1811 = v1810
	v1812 = base.I64_extend_i32_u(int32(15372) - v1798)
	v1813 = v1807 ^ int64(281474976710656)
	goto L469
L479:
	;
	goto L453
L480:
	;
	goto L427
L481:
	;
	F_addReplyLongLong(m, l0, base.I64_extend_i32_s(v1540))
	mBase = m.M
	v2076 = m.ExcPending
	if v2076 != 0 {
		goto L2
	} else {
		goto L516
	}
L482:
	;
	v2035 = int32(_a_F_georadiusGeneric_32)
	v2037 = *(*int64)(unsafe.Add(mBase, _c_F_georadiusGeneric[1]))
	*(*int64)(unsafe.Add(mBase, _c_F_georadiusGeneric[1])) = v2037 + v2011
	goto L481
L483:
	;
	v1986 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v1987 = F_dbDelete(m, v1986, v1342)
	mBase = m.M
	v1988 = m.ExcPending
	if v1988 != 0 {
		goto L2
	} else {
		goto L512
	}
L484:
	;
	v1837 = F_createZsetObject(m)
	mBase = m.M
	v1838 = m.ExcPending
	if v1838 != 0 {
		goto L2
	} else {
		goto L485
	}
L485:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v37)+752)) = v1837
	v1840 = F_objectGetVal(m, v1837)
	mBase = m.M
	v1841 = int32(0)
	if v1540 < int32(1) {
		v1947 = v1841
		v1953 = v1841
		goto L486
	} else {
		goto L487
	}
L486:
	;
	F_zsetConvertToListpackIfNeeded(m, v1837, v1953, v1947)
	mBase = m.M
	v1970 = m.ExcPending
	if v1970 != 0 {
		goto L2
	} else {
		goto L506
	}
L487:
	;
	v1845 = int32(0)
	v1850 = v1845
	v1860 = v1845
	v1866 = v1845
	goto L488
L488:
	;
	v1882 = *(*int32)(unsafe.Add(mBase, uint32(v37)+136))
	v1885 = v1882 + v1850*int32(40)
	v1886 = *(*float64)(unsafe.Add(mBase, uint32(v1885)+16))
	v1887 = *(*float64)(unsafe.Add(mBase, uint32(v37)+696))
	v1888 = base.F64_div(v1886, v1887)
	*(*float64)(unsafe.Add(mBase, uint32(v1885)+16)) = v1888
	if v1348 == int32(0) {
		v1893 = v1888
		goto L490
	} else {
		goto L491
	}
L489:
	;
	v1947 = v1931
	v1953 = v1930
	goto L486
L490:
	;
	v1895 = *(*int32)(unsafe.Add(mBase, uint32(v1885)+32))
	v1898 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1895+int32(-1)))))
	switch v1898 & int32(7) {
	case 0:
		goto L497
	case 1:
		goto L496
	case 2:
		goto L495
	case 3:
		goto L494
	case 4:
		goto L493
	default:
		v1915 = int32(0)
		goto L492
	}
L491:
	;
	v1892 = *(*float64)(unsafe.Add(mBase, uint32(v1885)+24))
	v1893 = v1892
	goto L490
L492:
	;
	v1916 = *(*int32)(unsafe.Add(mBase, uint32(v1840)+4))
	v1917 = F_zslInsert(m, v1916, v1893, v1895)
	mBase = m.M
	v1918 = m.ExcPending
	if v1918 != 0 {
		goto L2
	} else {
		goto L498
	}
L493:
	;
	v1914 = *(*int32)(unsafe.Add(mBase, uint32(v1895+int32(-17))))
	v1915 = v1914
	goto L492
L494:
	;
	v1911 = *(*int32)(unsafe.Add(mBase, uint32(v1895+int32(-9))))
	v1915 = v1911
	goto L492
L495:
	;
	v1908 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1895+int32(-5)))))
	v1915 = v1908
	goto L492
L496:
	;
	v1905 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1895+int32(-3)))))
	v1915 = v1905
	goto L492
L497:
	;
	v1915 = int32(base.Ui32(v1898) >> (uint(int32(3)) % 32))
	goto L492
L498:
	;
	v1919 = *(*int32)(unsafe.Add(mBase, uint32(v1840)))
	v1920 = F_hashtableAdd(m, v1919, v1917)
	mBase = m.M
	v1921 = m.ExcPending
	if v1921 != 0 {
		goto L2
	} else {
		goto L499
	}
L499:
	;
	if v1920 == int32(0) {
		goto L400
	} else {
		goto L500
	}
L500:
	;
	v1924 = *(*int32)(unsafe.Add(mBase, uint32(v1885)+32))
	F_sdsfree(m, v1924)
	mBase = m.M
	v1926 = m.ExcPending
	if v1926 != 0 {
		goto L2
	} else {
		goto L501
	}
L501:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1885)+32)) = int32(0)
	if base.Ui32(v1915) < base.Ui32(v1866) {
		goto L502
	} else {
		goto L503
	}
L502:
	;
	v1930 = v1866
	goto L504
L503:
	;
	v1930 = v1915
	goto L504
L504:
	;
	v1931 = v1915 + v1860
	v1933 = v1850 + int32(1)
	if v1933 != v1540 {
		v1850 = v1933
		v1860 = v1931
		v1866 = v1930
		goto L488
	} else {
		goto L505
	}
L505:
	;
	goto L489
L506:
	;
	v1971 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	F_setKey(m, l0, v1971, v1342, v37+int32(752), int32(0))
	mBase = m.M
	v1976 = m.ExcPending
	if v1976 != 0 {
		goto L2
	} else {
		goto L507
	}
L507:
	;
	if v1398 != 0 {
		goto L508
	} else {
		goto L509
	}
L508:
	;
	v1980 = int32(_a_F_georadiusGeneric_35)
	goto L510
L509:
	;
	v1980 = int32(_a_F_georadiusGeneric_36)
	goto L510
L510:
	;
	v1981 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v1982 = *(*int32)(unsafe.Add(mBase, uint32(v1981)+28))
	F_notifyKeyspaceEvent(m, int32(128), v1980, v1342, v1982)
	mBase = m.M
	v1984 = m.ExcPending
	if v1984 != 0 {
		goto L2
	} else {
		goto L511
	}
L511:
	;
	v2011 = base.I64_extend_i32_s(v1540)
	goto L482
L512:
	;
	if v1987 == int32(0) {
		goto L481
	} else {
		goto L513
	}
L513:
	;
	v1991 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	F_signalModifiedKey(m, l0, v1991, v1342)
	mBase = m.M
	v1993 = m.ExcPending
	if v1993 != 0 {
		goto L2
	} else {
		goto L514
	}
L514:
	;
	v1996 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v1997 = *(*int32)(unsafe.Add(mBase, uint32(v1996)+28))
	F_notifyKeyspaceEvent(m, int32(4), int32(_a_F_georadiusGeneric_31), v1342, v1997)
	mBase = m.M
	v1999 = m.ExcPending
	if v1999 != 0 {
		goto L2
	} else {
		goto L515
	}
L515:
	;
	v2011 = int64(1)
	goto L482
L516:
	;
	goto L427
L517:
	;
	v2115 = *(*int32)(unsafe.Add(mBase, uint32(v37)+672))
	if v2115 != int32(3) {
		goto L1
	} else {
		goto L518
	}
L518:
	;
	goto L401
L519:
	;
	F_valkey_free(m, v2152)
	mBase = m.M
	v2156 = m.ExcPending
	if v2156 != 0 {
		goto L2
	} else {
		goto L520
	}
L520:
	;
	goto L1
L521:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L522:
	;
	goto L1
}
func F_georadiusbymemberCommand(m *base.Module, l0 int32) {
	var v5 int32
	_ = v5
	F_georadiusGeneric(m, l0, int32(1), int32(2))
	v5 = m.ExcPending
	if v5 != 0 {
		return
	} else {
		return
	}
}
func F_georadiusroCommand(m *base.Module, l0 int32) {
	var v5 int32
	_ = v5
	F_georadiusGeneric(m, l0, int32(1), int32(5))
	v5 = m.ExcPending
	if v5 != 0 {
		return
	} else {
		return
	}
}
func F_geosearchCommand(m *base.Module, l0 int32) {
	var v5 int32
	_ = v5
	F_georadiusGeneric(m, l0, int32(1), int32(8))
	v5 = m.ExcPending
	if v5 != 0 {
		return
	} else {
		return
	}
}
func F_getAllClientsInfoString(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v33 int32
	_ = v33
	var v37 int32
	_ = v37
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v62 int32
	_ = v62
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v78 int32
	_ = v78
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v101 int32
	_ = v101
	var v105 int32
	_ = v105
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	v8 = m.G0
	v10 = v8 - int32(16)
	m.G0 = v10
	v13 = *(*int32)(unsafe.Add(mBase, _c_F_getAllClientsInfoString[0]))
	v15 = *(*int32)(unsafe.Add(mBase, _c_F_getAllClientsInfoString[1]))
	v16 = *(*int32)(unsafe.Add(mBase, uint32(v15)+20))
	v19 = F_sdsnewlen(m, v13, v16*int32(200))
	mBase = m.M
	v22 = m.ExcPending
	if v22 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v25 = v19 + int32(-1)
	v26 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v25))))
	switch v26 & int32(7) {
	case 0:
		goto L9
	case 1:
		goto L8
	case 2:
		goto L7
	case 3:
		goto L6
	case 4:
		goto L5
	default:
		goto L4
	}
L3:
	;
	v50 = *(*int32)(unsafe.Add(mBase, _c_F_getAllClientsInfoString[1]))
	v52 = v10 + int32(8)
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v50)))
	*(*int32)(unsafe.Add(mBase, uint32(v52)+4)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v52))) = v53
	goto L10
L4:
	;
	v47 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v19))) = uint8(v47)
	goto L3
L5:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v19+int32(-17)))) = int64(0)
	goto L4
L6:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19+int32(-9)))) = int32(0)
	goto L4
L7:
	;
	v37 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v19+int32(-5)))) = uint16(v37)
	goto L4
L8:
	;
	v33 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v19+int32(-3)))) = uint8(v33)
	goto L4
L9:
	;
	v29 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v25))) = uint8(v29)
	goto L4
L10:
	;
	v62 = v19
	goto L12
L11:
	;
	m.G0 = v10 + int32(16)
	return v62
L12:
	;
	v67 = v10 + int32(8)
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v67)))
	if v69 == int32(0) {
		goto L15
	} else {
		goto L16
	}
L14:
	;
	if v69 == int32(0) {
		goto L11
	} else {
		goto L17
	}
L15:
	;
	goto L14
L16:
	;
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v67)+4))
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v69+base.B2i32(v72 == int32(0))<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v67))) = v78
	goto L15
L17:
	;
	v82 = *(*int32)(unsafe.Add(mBase, uint32(v69)+8))
	if l0 == int32(-1) {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	v112 = F_catClientInfoString(m, v62, v82, l1)
	mBase = m.M
	v113 = m.ExcPending
	if v113 != 0 {
		goto L1
	} else {
		goto L35
	}
L19:
	;
	v83 = *(*int32)(unsafe.Add(mBase, uint32(v82)+200))
	if v83&int32(1) == int32(0) {
		goto L21
	} else {
		goto L22
	}
L20:
	;
	if v109 != l0 {
		goto L12
	} else {
		goto L34
	}
L21:
	;
	if v83&int32(2) == int32(0) {
		goto L23
	} else {
		goto L24
	}
L22:
	;
	v109 = int32(3)
	goto L20
L23:
	;
	if v83&int32(262144) == int32(0) {
		goto L26
	} else {
		goto L27
	}
L24:
	;
	if v83&int32(4) != 0 {
		goto L23
	} else {
		goto L25
	}
L25:
	;
	v109 = int32(1)
	goto L20
L26:
	;
	v101 = *(*int32)(unsafe.Add(mBase, uint32(v82)+216))
	if v101 != 0 {
		goto L28
	} else {
		goto L29
	}
L27:
	;
	v109 = int32(2)
	goto L20
L28:
	;
	v105 = *(*int32)(unsafe.Add(mBase, uint32(v101)))
	goto L30
L29:
	;
	v109 = int32(0)
	goto L20
L30:
	;
	if v105 == int32(1) {
		goto L31
	} else {
		goto L32
	}
L31:
	;
	v108 = int32(4)
	goto L33
L32:
	;
	v108 = int32(5)
	goto L33
L33:
	;
	v109 = v108
	goto L20
L34:
	;
	goto L18
L35:
	;
	v116 = F_sdscatlen(m, v112, int32(_a_F_getAllClientsInfoString_0), int32(1))
	mBase = m.M
	v117 = m.ExcPending
	if v117 != 0 {
		goto L1
	} else {
		goto L36
	}
L36:
	;
	v62 = v116
	goto L12
}
func F_getBaseAndIncrAppendOnlyFilesSize(m *base.Module, l0 int32, l1 int32) int64 {
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
	var v15 int32
	_ = v15
	var v16 int64
	_ = v16
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int64
	_ = v21
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v33 int64
	_ = v33
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v46 int32
	_ = v46
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v55 int64
	_ = v55
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v71 int64
	_ = v71
	var v80 int32
	_ = v80
	var v86 int32
	_ = v86
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v10 != 0 {
		goto L6
	} else {
		goto L7
	}
L1:
	;
	F__serverAssert(m, int32(_a_F_getBaseAndIncrAppendOnlyFilesSize_0), int32(_a_F_getBaseAndIncrAppendOnlyFilesSize_1), int32(2759))
	mBase = m.M
	v86 = m.ExcPending
	if v86 != 0 {
		goto L9
	} else {
		goto L23
	}
L2:
	;
	F__serverAssert(m, int32(_a_F_getBaseAndIncrAppendOnlyFilesSize_2), int32(_a_F_getBaseAndIncrAppendOnlyFilesSize_1), int32(2750))
	mBase = m.M
	v80 = m.ExcPending
	if v80 != 0 {
		goto L9
	} else {
		goto L22
	}
L3:
	;
	m.G0 = v8 + int32(16)
	return v71
L4:
	;
	v71 = int64(0)
	goto L3
L5:
	;
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v24 = v8 + int32(8)
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v22)))
	*(*int32)(unsafe.Add(mBase, uint32(v24)+4)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v24))) = v25
	goto L12
L6:
	;
	v12 = *(*int32)(unsafe.Add(mBase, uint32(v10)+16))
	if v12 != int32(98) {
		goto L2
	} else {
		goto L8
	}
L7:
	;
	v21 = int64(0)
	goto L5
L8:
	;
	v15 = *(*int32)(unsafe.Add(mBase, uint32(v10)))
	v16 = F_getAppendOnlyFileSize(m, v15, l1)
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	return int64(0)
L10:
	;
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	if v20 != 0 {
		goto L4
	} else {
		goto L11
	}
L11:
	;
	v21 = v16
	goto L5
L12:
	;
	v33 = v21
	goto L13
L13:
	;
	v35 = v8 + int32(8)
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v35)))
	if v37 == int32(0) {
		goto L16
	} else {
		goto L17
	}
L14:
	;
	goto L4
L15:
	;
	if v37 == int32(0) {
		v71 = v33
		goto L3
	} else {
		goto L18
	}
L16:
	;
	goto L15
L17:
	;
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v35)+4))
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v37+base.B2i32(v40 == int32(0))<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v35))) = v46
	goto L16
L18:
	;
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v37)+8))
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v50)+16))
	if v51 != int32(105) {
		goto L1
	} else {
		goto L19
	}
L19:
	;
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v50)))
	v55 = F_getAppendOnlyFileSize(m, v54, l1)
	mBase = m.M
	v56 = m.ExcPending
	if v56 != 0 {
		goto L9
	} else {
		goto L20
	}
L20:
	;
	v58 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	if v58 == int32(0) {
		v33 = v55 + v33
		goto L13
	} else {
		goto L21
	}
L21:
	;
	goto L14
L22:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L23:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_getBitfieldTypeFromArgument(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v72 int32
	_ = v72
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v103 int64
	_ = v103
	var v109 int32
	_ = v109
	var v111 int64
	_ = v111
	var v113 int32
	_ = v113
	var v115 int32
	_ = v115
	var v125 int64
	_ = v125
	var v130 int64
	_ = v130
	var v134 int32
	_ = v134
	var v136 int64
	_ = v136
	var v138 int32
	_ = v138
	var v146 int64
	_ = v146
	var v170 int64
	_ = v170
	var v189 int32
	_ = v189
	var v198 int64
	_ = v198
	var v201 int32
	_ = v201
	var v212 int32
	_ = v212
	var v216 int32
	_ = v216
	var v223 int32
	_ = v223
	v9 = m.G0
	v11 = v9 - int32(16)
	m.G0 = v11
	v14 = F_objectGetVal(m, l1)
	mBase = m.M
	v17 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14+int32(-1)))))
	switch v17 & int32(7) {
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
		v34 = int32(0)
		goto L1
	}
L1:
	;
	v36 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14))))
	switch v36 + int32(-105) {
	case 0:
		v46 = int32(1)
		goto L8
	default:
		goto L10
	case 12:
		goto L9
	}
L2:
	;
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v14+int32(-17))))
	v34 = v33
	goto L1
L3:
	;
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v14+int32(-9))))
	v34 = v30
	goto L1
L4:
	;
	v27 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v14+int32(-5)))))
	v34 = v27
	goto L1
L5:
	;
	v24 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14+int32(-3)))))
	v34 = v24
	goto L1
L6:
	;
	v34 = int32(base.Ui32(v17) >> (uint(int32(3)) % 32))
	goto L1
L7:
	;
	m.G0 = v11 + int32(16)
	return v216
L8:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v46
	v48 = int32(-1)
	v50 = v14 + int32(1)
	v52 = v34 + v48
	v54 = v11 + int32(8)
	v55 = int32(0)
	if base.Ui32(v34+int32(-22)) < base.Ui32(int32(-20)) {
		v189 = v55
		goto L16
	} else {
		goto L17
	}
L9:
	;
	v46 = int32(0)
	goto L8
L10:
	;
	F_addReplyError(m, l0, int32(_a_F_getBitfieldTypeFromArgument_0))
	mBase = m.M
	v43 = m.ExcPending
	if v43 != 0 {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	return int32(0)
L12:
	;
	v216 = int32(-1)
	goto L7
L13:
	;
	*(*uint32)(unsafe.Add(mBase, uint32(l3))) = uint32(v198)
	v216 = int32(0)
	goto L7
L14:
	;
	F_addReplyError(m, l0, int32(_a_F_getBitfieldTypeFromArgument_0))
	mBase = m.M
	v212 = m.ExcPending
	if v212 != 0 {
		goto L11
	} else {
		goto L49
	}
L15:
	;
	if v189 == int32(0) {
		goto L14
	} else {
		goto L42
	}
L16:
	;
	goto L15
L17:
	;
	v67 = int32(1)
	v68 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v50))))
	if v52 != v67 {
		goto L20
	} else {
		goto L21
	}
L18:
	;
	v189 = int32(1)
	goto L16
L19:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v54))) = v170
	goto L18
L20:
	;
	if v68&int32(255) == int32(45) {
		goto L25
	} else {
		goto L26
	}
L21:
	;
	v72 = v68 + int32(-48)
	if base.Ui32(int32(9)) < base.Ui32(v72&int32(255)) {
		v189 = v55
		goto L16
	} else {
		goto L22
	}
L22:
	;
	if v54 == int32(0) {
		goto L18
	} else {
		goto L23
	}
L23:
	;
	v170 = base.I64_extend_i32_u(v72) & int64(255)
	goto L19
L24:
	;
	if base.Ui32(int32(8)) < base.Ui32((v91+int32(-49))&int32(255)) {
		v189 = v55
		goto L16
	} else {
		goto L27
	}
L25:
	;
	v223 = int32(2)
	v88 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v50)+1)))
	v90 = v223
	v91 = v88
	v92 = v14 + v223
	goto L24
L26:
	;
	v90 = v67
	v91 = v68
	v92 = v50
	goto L24
L27:
	;
	v103 = base.I64_extend_i32_u(v91+int32(-48)) & int64(255)
	if base.Ui32(v52) <= base.Ui32(v90) {
		v146 = v103
		goto L28
	} else {
		goto L29
	}
L28:
	;
	if v68&int32(255) != int32(45) {
		goto L36
	} else {
		goto L37
	}
L29:
	;
	v109 = v90
	v111 = v103
	v113 = v92
	goto L30
L30:
	;
	v115 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v113)+1)))
	if base.Ui32((v115+int32(-58))&int32(255)) < base.Ui32(int32(246)) {
		v189 = v55
		goto L16
	} else {
		goto L32
	}
L31:
	;
	v146 = v136
	goto L28
L32:
	;
	if base.Ui64(int64(1844674407370955161)) < base.Ui64(v111) {
		v189 = v55
		goto L16
	} else {
		goto L33
	}
L33:
	;
	v125 = v111 * int64(10)
	v130 = base.I64_extend_i32_u(v115+int32(-48)) & int64(255)
	if base.Ui64(v130^int64(-1)) < base.Ui64(v125) {
		v189 = v55
		goto L16
	} else {
		goto L34
	}
L34:
	;
	v134 = int32(1)
	v136 = v125 + v130
	v138 = v109 + v134
	if v138 != v52 {
		v109 = v138
		v111 = v136
		v113 = v113 + v134
		goto L30
	} else {
		goto L35
	}
L35:
	;
	goto L31
L36:
	;
	if v146 < int64(0) {
		v189 = v55
		goto L16
	} else {
		goto L40
	}
L37:
	;
	if base.Ui64(int64(-9223372036854775807-1)) < base.Ui64(v146) {
		v189 = v55
		goto L16
	} else {
		goto L38
	}
L38:
	;
	if v54 == int32(0) {
		goto L18
	} else {
		goto L39
	}
L39:
	;
	v170 = int64(0) - v146
	goto L19
L40:
	;
	if v54 == int32(0) {
		goto L18
	} else {
		goto L41
	}
L41:
	;
	v170 = v146
	goto L19
L42:
	;
	v198 = *(*int64)(unsafe.Add(mBase, uint32(v11)+8))
	if v198 < int64(1) {
		goto L14
	} else {
		goto L43
	}
L43:
	;
	v201 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	if v201 != int32(1) {
		goto L44
	} else {
		goto L45
	}
L44:
	;
	if v201 != 0 {
		goto L13
	} else {
		goto L47
	}
L45:
	;
	if base.Ui64(int64(64)) < base.Ui64(v198) {
		goto L14
	} else {
		goto L46
	}
L46:
	;
	goto L44
L47:
	;
	if base.Ui64(v198) < base.Ui64(int64(64)) {
		goto L13
	} else {
		goto L48
	}
L48:
	;
	goto L14
L49:
	;
	v216 = v48
	goto L7
}
func F_getChannelsFromCommand(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v76 int32
	_ = v76
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v95 int32
	_ = v95
	var v98 int32
	_ = v98
	var v102 int32
	_ = v102
	var v104 int32
	_ = v104
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v134 int32
	_ = v134
	var v136 int32
	_ = v136
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v145 int32
	_ = v145
	var v147 int32
	_ = v147
	var v151 int32
	_ = v151
	var v160 int32
	_ = v160
	var v164 int32
	_ = v164
	var v166 int32
	_ = v166
	var v173 int32
	_ = v173
	var v176 int32
	_ = v176
	var v181 int32
	_ = v181
	var v198 int32
	_ = v198
	v12 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+59)))
	if v12&int32(8) != 0 {
		v138 = F_moduleGetCommandChannelsViaAPI(m, l0, l1, l2, l3)
		mBase = m.M
		v139 = m.ExcPending
		if v139 != 0 {
			return int32(0)
		} else {
			return v138
		}
	} else {
		v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
		v17 = int32(_a_F_getChannelsFromCommand_0)
		for {
			v30 = *(*int32)(unsafe.Add(mBase, uint32(v17)))
			if v15 != v30 {
				v17 = v17 + int32(24)
				continue
			} else {
				break
			}
			break
		}
		v32 = *(*int32)(unsafe.Add(mBase, uint32(v17)+16))
		v33 = *(*int32)(unsafe.Add(mBase, uint32(v17)+20))
		v34 = *(*int32)(unsafe.Add(mBase, uint32(l3)+8))
		if v34 != 0 {
			v39 = v34
			v40 = v33 + v32
			if v40 < l2 {
				v42 = v40
			} else {
				v42 = l2
			}
			if v33 == int32(-1) {
				v45 = l2
			} else {
				v45 = v42
			}
			v46 = v45 - v32
			v47 = *(*int32)(unsafe.Add(mBase, uint32(l3)+4))
			if v46 <= v47 {
				v76 = v39
				if v45 <= v32 {
					v198 = int32(0)
				} else {
					v80 = v46 & int32(3)
					v81 = *(*int32)(unsafe.Add(mBase, uint32(v17)+8))
					if base.Ui32(v32-v45) <= base.Ui32(int32(-4)) {
						v89 = int32(0)
						v91 = v32
						v95 = v89
						v98 = v89
						for {
							v102 = int32(3)
							v104 = v76 + v95<<(uint(v102)%32)
							*(*int32)(unsafe.Add(mBase, uint32(v104)+4)) = v81
							*(*int32)(unsafe.Add(mBase, uint32(v104))) = v91
							*(*int32)(unsafe.Add(mBase, uint32(v104+int32(8)))) = v91 + int32(1)
							*(*int32)(unsafe.Add(mBase, uint32(v104+int32(12)))) = v81
							*(*int32)(unsafe.Add(mBase, uint32(v104+int32(20)))) = v81
							*(*int32)(unsafe.Add(mBase, uint32(v104+int32(16)))) = v91 + int32(2)
							*(*int32)(unsafe.Add(mBase, uint32(v104+int32(28)))) = v81
							*(*int32)(unsafe.Add(mBase, uint32(v104+int32(24)))) = v91 + v102
							v131 = int32(4)
							v132 = v91 + v131
							v134 = v95 + v131
							v136 = v98 + v131
							if v136 != v46&int32(-4) {
								v91 = v132
								v95 = v134
								v98 = v136
								continue
							} else {
								break
							}
							break
						}
						v147 = v132
						v151 = v134
					} else {
						v147 = v32
						v151 = int32(0)
					}
					if v80 == int32(0) {
					} else {
						v160 = v147
						v164 = v151
						v166 = int32(0)
						for {
							v173 = v76 + v164<<(uint(int32(3))%32)
							*(*int32)(unsafe.Add(mBase, uint32(v173)+4)) = v81
							*(*int32)(unsafe.Add(mBase, uint32(v173))) = v160
							v176 = int32(1)
							v181 = v166 + v176
							if v181 != v80 {
								v160 = v160 + v176
								v164 = v164 + v176
								v166 = v181
								continue
							} else {
								break
							}
							break
						}
					}
					v198 = v46
				}
				*(*int32)(unsafe.Add(mBase, uint32(l3))) = v198
				return v198
			} else {
				v50 = v46 << (uint(int32(3)) % 32)
				v52 = l3 + int32(12)
				if v39 == v52 {
					v59 = F_valkey_malloc(m, v50)
					mBase = m.M
					v60 = m.ExcPending
					if v60 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(l3)+8)) = v59
						v62 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
						if v62 == int32(0) {
							v72 = v59
						} else {
							v66 = v62 << (uint(int32(3)) % 32)
							if v66 == int32(0) {
							} else {
								v69 = F__emscripten_memcpy_bulkmem(m, v59, v52, v66)
								mBase = m.M
							}
							v72 = v59
						}
						*(*int32)(unsafe.Add(mBase, uint32(l3)+4)) = v46
						v76 = v72
						if v45 <= v32 {
							v198 = int32(0)
						} else {
							v80 = v46 & int32(3)
							v81 = *(*int32)(unsafe.Add(mBase, uint32(v17)+8))
							if base.Ui32(v32-v45) <= base.Ui32(int32(-4)) {
								v89 = int32(0)
								v91 = v32
								v95 = v89
								v98 = v89
								for {
									v102 = int32(3)
									v104 = v76 + v95<<(uint(v102)%32)
									*(*int32)(unsafe.Add(mBase, uint32(v104)+4)) = v81
									*(*int32)(unsafe.Add(mBase, uint32(v104))) = v91
									*(*int32)(unsafe.Add(mBase, uint32(v104+int32(8)))) = v91 + int32(1)
									*(*int32)(unsafe.Add(mBase, uint32(v104+int32(12)))) = v81
									*(*int32)(unsafe.Add(mBase, uint32(v104+int32(20)))) = v81
									*(*int32)(unsafe.Add(mBase, uint32(v104+int32(16)))) = v91 + int32(2)
									*(*int32)(unsafe.Add(mBase, uint32(v104+int32(28)))) = v81
									*(*int32)(unsafe.Add(mBase, uint32(v104+int32(24)))) = v91 + v102
									v131 = int32(4)
									v132 = v91 + v131
									v134 = v95 + v131
									v136 = v98 + v131
									if v136 != v46&int32(-4) {
										v91 = v132
										v95 = v134
										v98 = v136
										continue
									} else {
										break
									}
									break
								}
								v147 = v132
								v151 = v134
							} else {
								v147 = v32
								v151 = int32(0)
							}
							if v80 == int32(0) {
							} else {
								v160 = v147
								v164 = v151
								v166 = int32(0)
								for {
									v173 = v76 + v164<<(uint(int32(3))%32)
									*(*int32)(unsafe.Add(mBase, uint32(v173)+4)) = v81
									*(*int32)(unsafe.Add(mBase, uint32(v173))) = v160
									v176 = int32(1)
									v181 = v166 + v176
									if v181 != v80 {
										v160 = v160 + v176
										v164 = v164 + v176
										v166 = v181
										continue
									} else {
										break
									}
									break
								}
							}
							v198 = v46
						}
						*(*int32)(unsafe.Add(mBase, uint32(l3))) = v198
						return v198
					}
				} else {
					v54 = F_valkey_realloc(m, v39, v50)
					mBase = m.M
					v57 = m.ExcPending
					if v57 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(l3)+8)) = v54
						v72 = v54
						*(*int32)(unsafe.Add(mBase, uint32(l3)+4)) = v46
						v76 = v72
						if v45 <= v32 {
							v198 = int32(0)
						} else {
							v80 = v46 & int32(3)
							v81 = *(*int32)(unsafe.Add(mBase, uint32(v17)+8))
							if base.Ui32(v32-v45) <= base.Ui32(int32(-4)) {
								v89 = int32(0)
								v91 = v32
								v95 = v89
								v98 = v89
								for {
									v102 = int32(3)
									v104 = v76 + v95<<(uint(v102)%32)
									*(*int32)(unsafe.Add(mBase, uint32(v104)+4)) = v81
									*(*int32)(unsafe.Add(mBase, uint32(v104))) = v91
									*(*int32)(unsafe.Add(mBase, uint32(v104+int32(8)))) = v91 + int32(1)
									*(*int32)(unsafe.Add(mBase, uint32(v104+int32(12)))) = v81
									*(*int32)(unsafe.Add(mBase, uint32(v104+int32(20)))) = v81
									*(*int32)(unsafe.Add(mBase, uint32(v104+int32(16)))) = v91 + int32(2)
									*(*int32)(unsafe.Add(mBase, uint32(v104+int32(28)))) = v81
									*(*int32)(unsafe.Add(mBase, uint32(v104+int32(24)))) = v91 + v102
									v131 = int32(4)
									v132 = v91 + v131
									v134 = v95 + v131
									v136 = v98 + v131
									if v136 != v46&int32(-4) {
										v91 = v132
										v95 = v134
										v98 = v136
										continue
									} else {
										break
									}
									break
								}
								v147 = v132
								v151 = v134
							} else {
								v147 = v32
								v151 = int32(0)
							}
							if v80 == int32(0) {
							} else {
								v160 = v147
								v164 = v151
								v166 = int32(0)
								for {
									v173 = v76 + v164<<(uint(int32(3))%32)
									*(*int32)(unsafe.Add(mBase, uint32(v173)+4)) = v81
									*(*int32)(unsafe.Add(mBase, uint32(v173))) = v160
									v176 = int32(1)
									v181 = v166 + v176
									if v181 != v80 {
										v160 = v160 + v176
										v164 = v164 + v176
										v166 = v181
										continue
									} else {
										break
									}
									break
								}
							}
							v198 = v46
						}
						*(*int32)(unsafe.Add(mBase, uint32(l3))) = v198
						return v198
					}
				}
			}
		} else {
			v35 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
			if v35 != 0 {
				F__serverAssert(m, int32(_a_F_getChannelsFromCommand_1), int32(_a_F_getChannelsFromCommand_2), int32(2292))
				mBase = m.M
				v145 = m.ExcPending
				if v145 != 0 {
					return int32(0)
				} else {
					F_abort(m)
					mBase = m.M
					base.Wasm_trap_unreachable()
					for {
					}
				}
			} else {
				v37 = l3 + int32(12)
				*(*int32)(unsafe.Add(mBase, uint32(l3)+8)) = v37
				v39 = v37
				v40 = v33 + v32
				if v40 < l2 {
					v42 = v40
				} else {
					v42 = l2
				}
				if v33 == int32(-1) {
					v45 = l2
				} else {
					v45 = v42
				}
				v46 = v45 - v32
				v47 = *(*int32)(unsafe.Add(mBase, uint32(l3)+4))
				if v46 <= v47 {
					v76 = v39
					if v45 <= v32 {
						v198 = int32(0)
					} else {
						v80 = v46 & int32(3)
						v81 = *(*int32)(unsafe.Add(mBase, uint32(v17)+8))
						if base.Ui32(v32-v45) <= base.Ui32(int32(-4)) {
							v89 = int32(0)
							v91 = v32
							v95 = v89
							v98 = v89
							for {
								v102 = int32(3)
								v104 = v76 + v95<<(uint(v102)%32)
								*(*int32)(unsafe.Add(mBase, uint32(v104)+4)) = v81
								*(*int32)(unsafe.Add(mBase, uint32(v104))) = v91
								*(*int32)(unsafe.Add(mBase, uint32(v104+int32(8)))) = v91 + int32(1)
								*(*int32)(unsafe.Add(mBase, uint32(v104+int32(12)))) = v81
								*(*int32)(unsafe.Add(mBase, uint32(v104+int32(20)))) = v81
								*(*int32)(unsafe.Add(mBase, uint32(v104+int32(16)))) = v91 + int32(2)
								*(*int32)(unsafe.Add(mBase, uint32(v104+int32(28)))) = v81
								*(*int32)(unsafe.Add(mBase, uint32(v104+int32(24)))) = v91 + v102
								v131 = int32(4)
								v132 = v91 + v131
								v134 = v95 + v131
								v136 = v98 + v131
								if v136 != v46&int32(-4) {
									v91 = v132
									v95 = v134
									v98 = v136
									continue
								} else {
									break
								}
								break
							}
							v147 = v132
							v151 = v134
						} else {
							v147 = v32
							v151 = int32(0)
						}
						if v80 == int32(0) {
						} else {
							v160 = v147
							v164 = v151
							v166 = int32(0)
							for {
								v173 = v76 + v164<<(uint(int32(3))%32)
								*(*int32)(unsafe.Add(mBase, uint32(v173)+4)) = v81
								*(*int32)(unsafe.Add(mBase, uint32(v173))) = v160
								v176 = int32(1)
								v181 = v166 + v176
								if v181 != v80 {
									v160 = v160 + v176
									v164 = v164 + v176
									v166 = v181
									continue
								} else {
									break
								}
								break
							}
						}
						v198 = v46
					}
					*(*int32)(unsafe.Add(mBase, uint32(l3))) = v198
					return v198
				} else {
					v50 = v46 << (uint(int32(3)) % 32)
					v52 = l3 + int32(12)
					if v39 == v52 {
						v59 = F_valkey_malloc(m, v50)
						mBase = m.M
						v60 = m.ExcPending
						if v60 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(l3)+8)) = v59
							v62 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
							if v62 == int32(0) {
								v72 = v59
							} else {
								v66 = v62 << (uint(int32(3)) % 32)
								if v66 == int32(0) {
								} else {
									v69 = F__emscripten_memcpy_bulkmem(m, v59, v52, v66)
									mBase = m.M
								}
								v72 = v59
							}
							*(*int32)(unsafe.Add(mBase, uint32(l3)+4)) = v46
							v76 = v72
							if v45 <= v32 {
								v198 = int32(0)
							} else {
								v80 = v46 & int32(3)
								v81 = *(*int32)(unsafe.Add(mBase, uint32(v17)+8))
								if base.Ui32(v32-v45) <= base.Ui32(int32(-4)) {
									v89 = int32(0)
									v91 = v32
									v95 = v89
									v98 = v89
									for {
										v102 = int32(3)
										v104 = v76 + v95<<(uint(v102)%32)
										*(*int32)(unsafe.Add(mBase, uint32(v104)+4)) = v81
										*(*int32)(unsafe.Add(mBase, uint32(v104))) = v91
										*(*int32)(unsafe.Add(mBase, uint32(v104+int32(8)))) = v91 + int32(1)
										*(*int32)(unsafe.Add(mBase, uint32(v104+int32(12)))) = v81
										*(*int32)(unsafe.Add(mBase, uint32(v104+int32(20)))) = v81
										*(*int32)(unsafe.Add(mBase, uint32(v104+int32(16)))) = v91 + int32(2)
										*(*int32)(unsafe.Add(mBase, uint32(v104+int32(28)))) = v81
										*(*int32)(unsafe.Add(mBase, uint32(v104+int32(24)))) = v91 + v102
										v131 = int32(4)
										v132 = v91 + v131
										v134 = v95 + v131
										v136 = v98 + v131
										if v136 != v46&int32(-4) {
											v91 = v132
											v95 = v134
											v98 = v136
											continue
										} else {
											break
										}
										break
									}
									v147 = v132
									v151 = v134
								} else {
									v147 = v32
									v151 = int32(0)
								}
								if v80 == int32(0) {
								} else {
									v160 = v147
									v164 = v151
									v166 = int32(0)
									for {
										v173 = v76 + v164<<(uint(int32(3))%32)
										*(*int32)(unsafe.Add(mBase, uint32(v173)+4)) = v81
										*(*int32)(unsafe.Add(mBase, uint32(v173))) = v160
										v176 = int32(1)
										v181 = v166 + v176
										if v181 != v80 {
											v160 = v160 + v176
											v164 = v164 + v176
											v166 = v181
											continue
										} else {
											break
										}
										break
									}
								}
								v198 = v46
							}
							*(*int32)(unsafe.Add(mBase, uint32(l3))) = v198
							return v198
						}
					} else {
						v54 = F_valkey_realloc(m, v39, v50)
						mBase = m.M
						v57 = m.ExcPending
						if v57 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(l3)+8)) = v54
							v72 = v54
							*(*int32)(unsafe.Add(mBase, uint32(l3)+4)) = v46
							v76 = v72
							if v45 <= v32 {
								v198 = int32(0)
							} else {
								v80 = v46 & int32(3)
								v81 = *(*int32)(unsafe.Add(mBase, uint32(v17)+8))
								if base.Ui32(v32-v45) <= base.Ui32(int32(-4)) {
									v89 = int32(0)
									v91 = v32
									v95 = v89
									v98 = v89
									for {
										v102 = int32(3)
										v104 = v76 + v95<<(uint(v102)%32)
										*(*int32)(unsafe.Add(mBase, uint32(v104)+4)) = v81
										*(*int32)(unsafe.Add(mBase, uint32(v104))) = v91
										*(*int32)(unsafe.Add(mBase, uint32(v104+int32(8)))) = v91 + int32(1)
										*(*int32)(unsafe.Add(mBase, uint32(v104+int32(12)))) = v81
										*(*int32)(unsafe.Add(mBase, uint32(v104+int32(20)))) = v81
										*(*int32)(unsafe.Add(mBase, uint32(v104+int32(16)))) = v91 + int32(2)
										*(*int32)(unsafe.Add(mBase, uint32(v104+int32(28)))) = v81
										*(*int32)(unsafe.Add(mBase, uint32(v104+int32(24)))) = v91 + v102
										v131 = int32(4)
										v132 = v91 + v131
										v134 = v95 + v131
										v136 = v98 + v131
										if v136 != v46&int32(-4) {
											v91 = v132
											v95 = v134
											v98 = v136
											continue
										} else {
											break
										}
										break
									}
									v147 = v132
									v151 = v134
								} else {
									v147 = v32
									v151 = int32(0)
								}
								if v80 == int32(0) {
								} else {
									v160 = v147
									v164 = v151
									v166 = int32(0)
									for {
										v173 = v76 + v164<<(uint(int32(3))%32)
										*(*int32)(unsafe.Add(mBase, uint32(v173)+4)) = v81
										*(*int32)(unsafe.Add(mBase, uint32(v173))) = v160
										v176 = int32(1)
										v181 = v166 + v176
										if v181 != v80 {
											v160 = v160 + v176
											v164 = v164 + v176
											v166 = v181
											continue
										} else {
											break
										}
										break
									}
								}
								v198 = v46
							}
							*(*int32)(unsafe.Add(mBase, uint32(l3))) = v198
							return v198
						}
					}
				}
			}
		}
	}
}
func F_getCurTid(m *base.Module) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	v2 = *(*int32)(unsafe.Add(mBase, _c_F_getCurTid[0]))
	return v2
}
func F_getExpirationPolicyWithFlags(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v6 int32
	_ = v6
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
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v38 int32
	_ = v38
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	v2 = int32(0)
	v6 = *(*int32)(unsafe.Add(mBase, _c_F_getExpirationPolicyWithFlags[0]))
	if v6 != 0 {
		v51 = v2
		return v51
	} else {
		v7 = int32(_a_F_getExpirationPolicyWithFlags_0)
		v8 = *(*int32)(unsafe.Add(mBase, _c_F_getExpirationPolicyWithFlags[1]))
		v10 = *(*int32)(unsafe.Add(mBase, _c_F_getExpirationPolicyWithFlags[2]))
		if v10 == int32(0) {
			if v8 == int32(0) {
				v34 = int32(1)
				if l0&v34 != 0 {
					if l0&int32(2) != 0 {
						v51 = int32(1)
						return v51
					} else {
						v46 = F_isPausedActionsWithUpdate(m, int32(4))
						mBase = m.M
						v49 = m.ExcPending
						if v49 != 0 {
							return int32(0)
						} else {
							if v46 != 0 {
								v50 = int32(1)
							} else {
								v50 = int32(2)
							}
							v51 = v50
							return v51
						}
					}
				} else {
					v38 = *(*int32)(unsafe.Add(mBase, _c_F_getExpirationPolicyWithFlags[3]))
					if v38 != 0 {
						v51 = v34
						return v51
					} else {
						if l0&int32(2) != 0 {
							v51 = int32(1)
							return v51
						} else {
							v46 = F_isPausedActionsWithUpdate(m, int32(4))
							mBase = m.M
							v49 = m.ExcPending
							if v49 != 0 {
								return int32(0)
							} else {
								if v46 != 0 {
									v50 = int32(1)
								} else {
									v50 = int32(2)
								}
								v51 = v50
								return v51
							}
						}
					}
				}
			} else {
				v23 = *(*int32)(unsafe.Add(mBase, uint32(v8)+216))
				if v23 != 0 {
					v51 = v2
					return v51
				} else {
					v25 = *(*int32)(unsafe.Add(mBase, _c_F_getExpirationPolicyWithFlags[3]))
					if v25 == int32(0) {
						if l0&int32(2) != 0 {
							v51 = int32(1)
							return v51
						} else {
							v46 = F_isPausedActionsWithUpdate(m, int32(4))
							mBase = m.M
							v49 = m.ExcPending
							if v49 != 0 {
								return int32(0)
							} else {
								if v46 != 0 {
									v50 = int32(1)
								} else {
									v50 = int32(2)
								}
								v51 = v50
								return v51
							}
						}
					} else {
						v28 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+207)))
						if v28&int32(32) != 0 {
							v51 = v2
							return v51
						} else {
							v31 = int32(1)
							if l0&v31 != 0 {
								if l0&int32(2) != 0 {
									v51 = int32(1)
									return v51
								} else {
									v46 = F_isPausedActionsWithUpdate(m, int32(4))
									mBase = m.M
									v49 = m.ExcPending
									if v49 != 0 {
										return int32(0)
									} else {
										if v46 != 0 {
											v50 = int32(1)
										} else {
											v50 = int32(2)
										}
										v51 = v50
										return v51
									}
								}
							} else {
								v51 = v31
								return v51
							}
						}
					}
				}
			}
		} else {
			if v8 == int32(0) {
				v18 = int32(1)
				if l0&v18 != 0 {
					if l0&int32(2) != 0 {
						v51 = int32(1)
						return v51
					} else {
						v46 = F_isPausedActionsWithUpdate(m, int32(4))
						mBase = m.M
						v49 = m.ExcPending
						if v49 != 0 {
							return int32(0)
						} else {
							if v46 != 0 {
								v50 = int32(1)
							} else {
								v50 = int32(2)
							}
							v51 = v50
							return v51
						}
					}
				} else {
					v51 = v18
					return v51
				}
			} else {
				v15 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+200)))
				if v15&int32(1) != 0 {
					v51 = v2
					return v51
				} else {
					v18 = int32(1)
					if l0&v18 != 0 {
						if l0&int32(2) != 0 {
							v51 = int32(1)
							return v51
						} else {
							v46 = F_isPausedActionsWithUpdate(m, int32(4))
							mBase = m.M
							v49 = m.ExcPending
							if v49 != 0 {
								return int32(0)
							} else {
								if v46 != 0 {
									v50 = int32(1)
								} else {
									v50 = int32(2)
								}
								v51 = v50
								return v51
							}
						}
					} else {
						v51 = v18
						return v51
					}
				}
			}
		}
	}
}
func F_getImportingSlotSource(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	v3 = *(*int32)(unsafe.Add(mBase, _c_F_getImportingSlotSource[0]))
	v4 = *(*int32)(unsafe.Add(mBase, uint32(v3)+48))
	v5 = F_dictFind(m, v4, l0)
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		return int32(0)
	} else {
		if v5 != 0 {
			v11 = *(*int32)(unsafe.Add(mBase, uint32(v5)+8))
			return v11
		} else {
			return int32(0)
		}
	}
}
func F_getInstantaneousMetric(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v6 int64
	_ = v6
	var v9 int64
	_ = v9
	var v12 int64
	_ = v12
	var v15 int64
	_ = v15
	var v18 int64
	_ = v18
	var v21 int64
	_ = v21
	var v24 int64
	_ = v24
	var v27 int64
	_ = v27
	var v30 int64
	_ = v30
	var v33 int64
	_ = v33
	var v36 int64
	_ = v36
	var v39 int64
	_ = v39
	var v42 int64
	_ = v42
	var v45 int64
	_ = v45
	var v48 int64
	_ = v48
	var v51 int64
	_ = v51
	var v68 int64
	_ = v68
	v3 = l0 * int32(152)
	v6 = *(*int64)(unsafe.Add(mBase, uint32(v3)+uint32(_c_F_getInstantaneousMetric[0])))
	v9 = *(*int64)(unsafe.Add(mBase, uint32(v3)+uint32(_c_F_getInstantaneousMetric[1])))
	v12 = *(*int64)(unsafe.Add(mBase, uint32(v3)+uint32(_c_F_getInstantaneousMetric[2])))
	v15 = *(*int64)(unsafe.Add(mBase, uint32(v3)+uint32(_c_F_getInstantaneousMetric[3])))
	v18 = *(*int64)(unsafe.Add(mBase, uint32(v3)+uint32(_c_F_getInstantaneousMetric[4])))
	v21 = *(*int64)(unsafe.Add(mBase, uint32(v3)+uint32(_c_F_getInstantaneousMetric[5])))
	v24 = *(*int64)(unsafe.Add(mBase, uint32(v3)+uint32(_c_F_getInstantaneousMetric[6])))
	v27 = *(*int64)(unsafe.Add(mBase, uint32(v3)+uint32(_c_F_getInstantaneousMetric[7])))
	v30 = *(*int64)(unsafe.Add(mBase, uint32(v3)+uint32(_c_F_getInstantaneousMetric[8])))
	v33 = *(*int64)(unsafe.Add(mBase, uint32(v3)+uint32(_c_F_getInstantaneousMetric[9])))
	v36 = *(*int64)(unsafe.Add(mBase, uint32(v3)+uint32(_c_F_getInstantaneousMetric[10])))
	v39 = *(*int64)(unsafe.Add(mBase, uint32(v3)+uint32(_c_F_getInstantaneousMetric[11])))
	v42 = *(*int64)(unsafe.Add(mBase, uint32(v3)+uint32(_c_F_getInstantaneousMetric[12])))
	v45 = *(*int64)(unsafe.Add(mBase, uint32(v3)+uint32(_c_F_getInstantaneousMetric[13])))
	v48 = *(*int64)(unsafe.Add(mBase, uint32(v3)+uint32(_c_F_getInstantaneousMetric[14])))
	v51 = *(*int64)(unsafe.Add(mBase, uint32(v3)+uint32(_c_F_getInstantaneousMetric[15])))
	v68 = base.I64_div_s(v6+(v9+(v12+(v15+(v18+(v21+(v24+(v27+(v30+(v33+(v36+(v39+(v42+(v45+(v48+v51)))))))))))))), int64(16))
	return v68
}
func F_getKVStoreIndexForKey(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	v3 = *(*int32)(unsafe.Add(mBase, _c_F_getKVStoreIndexForKey[0]))
	if v3 != 0 {
		v6 = F_getKeySlot(m, l0)
		mBase = m.M
		v9 = m.ExcPending
		if v9 != 0 {
			return int32(0)
		} else {
			return v6
		}
	} else {
		return int32(0)
	}
}
func F_getLastIncrAofName(m *base.Module, l0 int32) int32 {
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
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v26 int64
	_ = v26
	var v28 int64
	_ = v28
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v44 int64
	_ = v44
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v57 int32
	_ = v57
	var v66 int32
	_ = v66
	v7 = m.G0
	v9 = v7 - int32(32)
	m.G0 = v9
	if l0 == int32(0) {
		F__serverAssert(m, int32(_a_F_getLastIncrAofName_0), int32(_a_F_getLastIncrAofName_1), int32(472))
		mBase = m.M
		v66 = m.ExcPending
		if v66 != 0 {
			return int32(0)
		} else {
			F_abort(m)
			mBase = m.M
			base.Wasm_trap_unreachable()
			for {
			}
		}
	} else {
		v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		v14 = *(*int32)(unsafe.Add(mBase, uint32(v13)+20))
		if v14 != 0 {
			v51 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
			v52 = *(*int32)(unsafe.Add(mBase, uint32(v51)+8))
			v53 = v52
			v57 = *(*int32)(unsafe.Add(mBase, uint32(v53)))
			m.G0 = v9 + int32(32)
			return v57
		} else {
			v16 = F_valkey_calloc(m, int32(24))
			mBase = m.M
			v19 = m.ExcPending
			if v19 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v16)+16)) = int32(105)
				v22 = F_sdsempty(m)
				mBase = m.M
				v23 = m.ExcPending
				if v23 != 0 {
					return int32(0)
				} else {
					v25 = *(*int32)(unsafe.Add(mBase, _c_F_getLastIncrAofName[0]))
					v26 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
					v28 = v26 + int64(1)
					*(*int64)(unsafe.Add(mBase, uint32(l0)+24)) = v28
					*(*int32)(unsafe.Add(mBase, uint32(v9+int32(20)))) = int32(_a_F_getLastIncrAofName_2)
					*(*int32)(unsafe.Add(mBase, uint32(v9+int32(16)))) = int32(_a_F_getLastIncrAofName_3)
					*(*int64)(unsafe.Add(mBase, uint32(v9)+8)) = v28
					*(*int32)(unsafe.Add(mBase, uint32(v9))) = v25
					v41 = F_sdscatprintf(m, v22, int32(_a_F_getLastIncrAofName_4), v9)
					mBase = m.M
					v42 = m.ExcPending
					if v42 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v16))) = v41
						v44 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
						*(*int64)(unsafe.Add(mBase, uint32(v16)+8)) = v44
						v46 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
						v47 = F_listAddNodeTail(m, v46, v16)
						mBase = m.M
						v48 = m.ExcPending
						if v48 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = int32(1)
							v53 = v16
							v57 = *(*int32)(unsafe.Add(mBase, uint32(v53)))
							m.G0 = v9 + int32(32)
							return v57
						}
					}
				}
			}
		}
	}
}
func F_getMemoryOverheadData(m *base.Module) int32 {
	mBase := m.M
	_ = mBase
	var v1 int32
	_ = v1
	var v27 int32
	_ = v27
	var v31 int32
	_ = v31
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v52 int32
	_ = v52
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
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v102 int32
	_ = v102
	var v104 int32
	_ = v104
	var v112 int32
	_ = v112
	var v115 int32
	_ = v115
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v121 int32
	_ = v121
	var v124 int32
	_ = v124
	var v126 int32
	_ = v126
	var v129 float32
	_ = v129
	var v134 int32
	_ = v134
	var v136 int32
	_ = v136
	var v145 int32
	_ = v145
	var v147 int32
	_ = v147
	var v152 float32
	_ = v152
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v165 int64
	_ = v165
	var v168 int32
	_ = v168
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v176 int32
	_ = v176
	var v179 int32
	_ = v179
	var v180 int64
	_ = v180
	var v181 int64
	_ = v181
	var v186 int32
	_ = v186
	var v188 int32
	_ = v188
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v194 int32
	_ = v194
	var v196 int32
	_ = v196
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v203 int32
	_ = v203
	var v206 int32
	_ = v206
	var v209 int32
	_ = v209
	var v212 int32
	_ = v212
	var v215 int32
	_ = v215
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v231 int32
	_ = v231
	var v238 int32
	_ = v238
	var v242 int32
	_ = v242
	var v246 int32
	_ = v246
	var v250 int32
	_ = v250
	var v251 int32
	_ = v251
	var v255 int32
	_ = v255
	var v256 int32
	_ = v256
	var v258 int32
	_ = v258
	var v260 int32
	_ = v260
	var v262 int32
	_ = v262
	var v263 int32
	_ = v263
	var v266 int32
	_ = v266
	var v267 int32
	_ = v267
	var v268 int32
	_ = v268
	var v274 int32
	_ = v274
	var v275 int32
	_ = v275
	var v278 int32
	_ = v278
	var v280 int32
	_ = v280
	var v281 int32
	_ = v281
	var v282 int32
	_ = v282
	var v283 int32
	_ = v283
	var v284 int32
	_ = v284
	var v287 int32
	_ = v287
	var v288 int32
	_ = v288
	var v290 int32
	_ = v290
	var v293 int32
	_ = v293
	var v304 int32
	_ = v304
	var v306 int32
	_ = v306
	var v310 int32
	_ = v310
	var v317 int32
	_ = v317
	var v319 int32
	_ = v319
	var v322 int32
	_ = v322
	var v323 int32
	_ = v323
	var v332 int32
	_ = v332
	var v335 int32
	_ = v335
	var v336 int32
	_ = v336
	var v339 int64
	_ = v339
	var v340 int32
	_ = v340
	var v341 int32
	_ = v341
	var v343 int32
	_ = v343
	var v346 int64
	_ = v346
	var v347 int32
	_ = v347
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
	var v357 int32
	_ = v357
	var v358 int32
	_ = v358
	var v360 int32
	_ = v360
	var v362 int32
	_ = v362
	var v364 int32
	_ = v364
	var v365 int32
	_ = v365
	var v366 int32
	_ = v366
	var v371 int32
	_ = v371
	var v372 int32
	_ = v372
	var v375 int32
	_ = v375
	var v381 int32
	_ = v381
	var v382 int32
	_ = v382
	var v384 int32
	_ = v384
	var v386 int32
	_ = v386
	var v388 int32
	_ = v388
	var v389 int32
	_ = v389
	var v390 int32
	_ = v390
	var v395 int32
	_ = v395
	var v396 int32
	_ = v396
	var v399 int32
	_ = v399
	var v405 int32
	_ = v405
	var v410 int32
	_ = v410
	var v411 int32
	_ = v411
	var v412 int32
	_ = v412
	var v413 int32
	_ = v413
	var v414 int32
	_ = v414
	var v418 int32
	_ = v418
	var v419 int32
	_ = v419
	var v420 int32
	_ = v420
	var v421 int32
	_ = v421
	var v422 int32
	_ = v422
	var v426 int32
	_ = v426
	var v427 int32
	_ = v427
	var v428 int32
	_ = v428
	var v429 int32
	_ = v429
	var v430 int32
	_ = v430
	var v431 int32
	_ = v431
	var v432 int32
	_ = v432
	var v438 int32
	_ = v438
	var v439 int32
	_ = v439
	var v441 int32
	_ = v441
	var v442 int32
	_ = v442
	var v444 int32
	_ = v444
	var v445 int32
	_ = v445
	var v449 int32
	_ = v449
	var v462 int32
	_ = v462
	var v469 int32
	_ = v469
	var v472 float32
	_ = v472
	var v484 float32
	_ = v484
	var v487 int32
	_ = v487
	var v491 int32
	_ = v491
	v1 = int32(0)
	v27 = *(*int32)(unsafe.Add(mBase, _c_F_getMemoryOverheadData[0]))
	if v27 < int32(261) {
		goto L4
	} else {
		goto L5
	}
L1:
	;
	v112 = F_valkey_calloc(m, int32(128))
	mBase = m.M
	v115 = m.ExcPending
	if v115 != 0 {
		goto L17
	} else {
		goto L18
	}
L2:
	;
	goto L1
L3:
	;
	v38 = v36 & int32(3)
	if base.Ui32(int32(4)) <= base.Ui32(v36) {
		goto L8
	} else {
		goto L9
	}
L4:
	;
	if v27 < int32(1) {
		v104 = v1
		goto L2
	} else {
		goto L6
	}
L5:
	;
	v31 = *(*int32)(unsafe.Add(mBase, _c_F_getMemoryOverheadData[1]))
	v35 = v31
	v36 = int32(260)
	goto L3
L6:
	;
	v35 = v1
	v36 = v27
	goto L3
L7:
	;
	if v38 == int32(0) {
		v104 = v77
		goto L2
	} else {
		goto L13
	}
L8:
	;
	v45 = int32(0)
	v47 = v35
	v48 = v45
	v52 = v45
	goto L10
L9:
	;
	v77 = v35
	v78 = int32(0)
	goto L7
L10:
	;
	v55 = v48 << (uint(int32(2)) % 32)
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v55)+uint32(_c_F_getMemoryOverheadData[2])))
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v55)+uint32(_c_F_getMemoryOverheadData[3])))
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v55)+uint32(_c_F_getMemoryOverheadData[4])))
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v55)+uint32(_c_F_getMemoryOverheadData[5])))
	v71 = v58 + (v61 + (v64 + (v67 + v47)))
	v72 = int32(4)
	v73 = v48 + v72
	v75 = v52 + v72
	if v75 != v36&int32(2147483644) {
		v47 = v71
		v48 = v73
		v52 = v75
		goto L10
	} else {
		goto L12
	}
L11:
	;
	v77 = v71
	v78 = v73
	goto L7
L12:
	;
	goto L11
L13:
	;
	v86 = v77
	v87 = v78
	v89 = int32(0)
	goto L14
L14:
	;
	v97 = *(*int32)(unsafe.Add(mBase, uint32(v87<<(uint(int32(2))%32))+uint32(_c_F_getMemoryOverheadData[5])))
	v98 = v97 + v86
	v99 = int32(1)
	v102 = v89 + v99
	if v102 != v38 {
		v86 = v98
		v87 = v87 + v99
		v89 = v102
		goto L14
	} else {
		goto L16
	}
L15:
	;
	v104 = v98
	goto L2
L16:
	;
	goto L15
L17:
	;
	return int32(0)
L18:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v112)+4)) = v104
	v117 = int32(_a_F_getMemoryOverheadData_0)
	v118 = *(*int32)(unsafe.Add(mBase, _c_F_getMemoryOverheadData[6]))
	*(*int32)(unsafe.Add(mBase, uint32(v112)+8)) = v118
	v121 = *(*int32)(unsafe.Add(mBase, _c_F_getMemoryOverheadData[7]))
	*(*int32)(unsafe.Add(mBase, uint32(v112))) = v121
	v124 = *(*int32)(unsafe.Add(mBase, _c_F_getMemoryOverheadData[8]))
	v126 = *(*int32)(unsafe.Add(mBase, _c_F_getMemoryOverheadData[9]))
	*(*int32)(unsafe.Add(mBase, uint32(v112)+80)) = v124 - v126
	v129 = base.F32_convert_i32_u(v124)
	*(*float32)(unsafe.Add(mBase, uint32(v112)+76)) = base.F32_div(v129, base.F32_convert_i32_u(v126))
	v134 = *(*int32)(unsafe.Add(mBase, _c_F_getMemoryOverheadData[10]))
	v136 = *(*int32)(unsafe.Add(mBase, _c_F_getMemoryOverheadData[11]))
	*(*int32)(unsafe.Add(mBase, uint32(v112)+88)) = v136
	*(*float32)(unsafe.Add(mBase, uint32(v112)+84)) = base.F32_add(base.F32_div(base.F32_convert_i32_u(v136), base.F32_convert_i32_u(v134)), float32(1))
	v145 = *(*int32)(unsafe.Add(mBase, _c_F_getMemoryOverheadData[12]))
	v147 = *(*int32)(unsafe.Add(mBase, _c_F_getMemoryOverheadData[13]))
	*(*int32)(unsafe.Add(mBase, uint32(v112)+104)) = v124 - v147
	*(*int32)(unsafe.Add(mBase, uint32(v112)+96)) = v147 - v145
	v152 = base.F32_convert_i32_u(v147)
	*(*float32)(unsafe.Add(mBase, uint32(v112)+100)) = base.F32_div(v129, v152)
	*(*float32)(unsafe.Add(mBase, uint32(v112)+92)) = base.F32_div(v152, base.F32_convert_i32_u(v145))
	v159 = *(*int32)(unsafe.Add(mBase, _c_F_getMemoryOverheadData[14]))
	v160 = int32(0)
	v162 = *(*int32)(unsafe.Add(mBase, _c_F_getMemoryOverheadData[15]))
	v163 = *(*int32)(unsafe.Add(mBase, uint32(v162)+20))
	if v163 != 0 {
		goto L20
	} else {
		goto L21
	}
L19:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v112)+12)) = v170
	*(*int32)(unsafe.Add(mBase, uint32(v112)+20)) = v171
	v176 = *(*int32)(unsafe.Add(mBase, _c_F_getMemoryOverheadData[16]))
	if v176 == int32(0) {
		v188 = v170
		goto L24
	} else {
		goto L25
	}
L20:
	;
	v165 = *(*int64)(unsafe.Add(mBase, _c_F_getMemoryOverheadData[17]))
	if v165 < base.I64_extend_i32_u(v159) {
		goto L22
	} else {
		goto L23
	}
L21:
	;
	v170 = v159
	v171 = v160
	goto L19
L22:
	;
	v168 = base.I32_wrap_i64(v165)
	v170 = v168
	v171 = v159 - v168
	goto L19
L23:
	;
	v170 = v159
	v171 = v160
	goto L19
L24:
	;
	v190 = int32(_a_F_getMemoryOverheadData_0)
	v191 = *(*int32)(unsafe.Add(mBase, _c_F_getMemoryOverheadData[18]))
	*(*int32)(unsafe.Add(mBase, uint32(v112)+16)) = v191
	v194 = *(*int32)(unsafe.Add(mBase, _c_F_getMemoryOverheadData[19]))
	v196 = *(*int32)(unsafe.Add(mBase, _c_F_getMemoryOverheadData[20]))
	v199 = *(*int32)(unsafe.Add(mBase, _c_F_getMemoryOverheadData[21]))
	v200 = v194 + v196 + v199
	*(*int32)(unsafe.Add(mBase, uint32(v112)+24)) = v200
	v203 = *(*int32)(unsafe.Add(mBase, _c_F_getMemoryOverheadData[22]))
	*(*int32)(unsafe.Add(mBase, uint32(v112)+32)) = v203
	v206 = *(*int32)(unsafe.Add(mBase, _c_F_getMemoryOverheadData[23]))
	*(*int32)(unsafe.Add(mBase, uint32(v112)+36)) = v206
	v209 = *(*int32)(unsafe.Add(mBase, _c_F_getMemoryOverheadData[24]))
	*(*int32)(unsafe.Add(mBase, uint32(v112)+28)) = v209
	v212 = *(*int32)(unsafe.Add(mBase, _c_F_getMemoryOverheadData[25]))
	if v212 != 0 {
		goto L28
	} else {
		goto L29
	}
L25:
	;
	v179 = *(*int32)(unsafe.Add(mBase, uint32(v176)+8))
	v180 = *(*int64)(unsafe.Add(mBase, uint32(v179)+16))
	v181 = *(*int64)(unsafe.Add(mBase, uint32(v179)+8))
	goto L26
L26:
	;
	v186 = base.I32_wrap_i64(v180+v181)<<(uint(int32(2))%32) + v170
	*(*int32)(unsafe.Add(mBase, uint32(v112)+12)) = v186
	v188 = v186
	goto L24
L27:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v112)+40)) = v256
	v258 = int32(0)
	v260 = *(*int32)(unsafe.Add(mBase, _c_F_getMemoryOverheadData[26]))
	v262 = *(*int32)(unsafe.Add(mBase, _c_F_getMemoryOverheadData[27]))
	v263 = F_dictMemUsage(m, v262)
	mBase = m.M
	v266 = *(*int32)(unsafe.Add(mBase, _c_F_getMemoryOverheadData[27]))
	v267 = *(*int32)(unsafe.Add(mBase, uint32(v266)+16))
	v268 = *(*int32)(unsafe.Add(mBase, uint32(v266)+12))
	v274 = *(*int32)(unsafe.Add(mBase, _c_F_getMemoryOverheadData[28]))
	v275 = *(*int32)(unsafe.Add(mBase, uint32(v274)+20))
	v278 = v260 + v263 + (v267+v268)<<(uint(int32(5))%32) + v275*int32(12)
	goto L43
L28:
	;
	v215 = *(*int32)(unsafe.Add(mBase, _c_F_getMemoryOverheadData[29]))
	v222 = v215 + int32(-1)
	v223 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v222))))
	v225 = v223 & int32(7)
	switch v225 {
	case 0:
		goto L36
	case 1:
		v231 = int32(4)
		goto L31
	case 2:
		goto L35
	case 3:
		goto L34
	case 4:
		goto L33
	default:
		goto L32
	}
L29:
	;
	v256 = int32(0)
	goto L27
L30:
	;
	v256 = v255
	goto L27
L31:
	;
	switch v225 {
	case 0:
		goto L42
	case 1:
		goto L41
	case 2:
		goto L40
	case 3:
		goto L39
	case 4:
		goto L38
	default:
		v251 = int32(0)
		goto L37
	}
L32:
	;
	v231 = int32(1)
	goto L31
L33:
	;
	v231 = int32(18)
	goto L31
L34:
	;
	v231 = int32(10)
	goto L31
L35:
	;
	v231 = int32(6)
	goto L31
L36:
	;
	v226 = F_zmalloc_usable_size(m, v222)
	mBase = m.M
	v255 = v226
	goto L30
L37:
	;
	v255 = v231 + v251
	goto L30
L38:
	;
	v250 = *(*int32)(unsafe.Add(mBase, uint32(v215+int32(-9))))
	v251 = v250
	goto L37
L39:
	;
	v246 = *(*int32)(unsafe.Add(mBase, uint32(v215+int32(-5))))
	v255 = v231 + v246
	goto L30
L40:
	;
	v242 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v215+int32(-3)))))
	v255 = v231 + v242
	goto L30
L41:
	;
	v238 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v215+int32(-2)))))
	v255 = v231 + v238
	goto L30
L42:
	;
	v255 = v231 + int32(base.Ui32(v223)>>(uint(int32(3))%32))
	goto L30
L43:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v112)+44)) = v278
	v280 = F_scriptingEngineManagerGetMemoryUsage(m)
	mBase = m.M
	v281 = int32(0)
	v282 = *(*int32)(unsafe.Add(mBase, _c_F_getMemoryOverheadData[30]))
	v283 = *(*int32)(unsafe.Add(mBase, uint32(v282)+4))
	v284 = F_dictMemUsage(m, v283)
	mBase = m.M
	v287 = *(*int32)(unsafe.Add(mBase, _c_F_getMemoryOverheadData[30]))
	v288 = *(*int32)(unsafe.Add(mBase, uint32(v287)+8))
	v290 = F_scriptingEngineManagerGetTotalMemoryOverhead(m)
	mBase = m.M
	v293 = v280 + v284 + v288 + v290 + int32(16)
	goto L44
L44:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v112)+48)) = v293
	v304 = v293 + (v278 + (v191 + v118 + v188 + v171 + v200 + v203 + v206 + v209 + v256))
	v306 = *(*int32)(unsafe.Add(mBase, _c_F_getMemoryOverheadData[31]))
	if v306 < int32(1) {
		v462 = v304
		goto L45
	} else {
		goto L46
	}
L45:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v112)+52)) = v462
	v469 = v104 - v462
	*(*int32)(unsafe.Add(mBase, uint32(v112)+56)) = v469
	v472 = float32(100)
	*(*float32)(unsafe.Add(mBase, uint32(v112)+72)) = base.F32_div(base.F32_mul(base.F32_convert_i32_u(v104), v472), base.F32_convert_i32_u(v121))
	if base.Ui32(v118) < base.Ui32(v104) {
		goto L70
	} else {
		goto L71
	}
L46:
	;
	v310 = *(*int32)(unsafe.Add(mBase, _c_F_getMemoryOverheadData[32]))
	v317 = int32(0)
	v319 = v306
	v322 = v310
	v323 = v304
	goto L47
L47:
	;
	v332 = *(*int32)(unsafe.Add(mBase, uint32(v322+v317<<(uint(int32(2))%32))))
	if v332 == int32(0) {
		v442 = v319
		v444 = v322
		v445 = v323
		goto L49
	} else {
		goto L50
	}
L48:
	;
	v462 = v445
	goto L45
L49:
	;
	v449 = v317 + int32(1)
	if v449 < v442 {
		v317 = v449
		v319 = v442
		v322 = v444
		v323 = v445
		goto L47
	} else {
		goto L69
	}
L50:
	;
	v335 = *(*int32)(unsafe.Add(mBase, uint32(v332)))
	v336 = *(*int32)(unsafe.Add(mBase, uint32(v335)+12))
	if v336 == int32(1) {
		goto L52
	} else {
		goto L53
	}
L51:
	;
	v347 = *(*int32)(unsafe.Add(mBase, uint32(v112)+60))
	v348 = base.I32_wrap_i64(v346)
	*(*int32)(unsafe.Add(mBase, uint32(v112)+60)) = v347 + v348
	v351 = *(*int32)(unsafe.Add(mBase, uint32(v112)+124))
	v352 = *(*int32)(unsafe.Add(mBase, uint32(v112)+108))
	v353 = int32(12)
	v354 = v352 * v353
	v357 = F_valkey_realloc(m, v351, v354+v353)
	mBase = m.M
	v358 = m.ExcPending
	if v358 != 0 {
		goto L17
	} else {
		goto L56
	}
L52:
	;
	v340 = *(*int32)(unsafe.Add(mBase, uint32(v335)+8))
	v341 = *(*int32)(unsafe.Add(mBase, uint32(v340)))
	if v341 != 0 {
		goto L54
	} else {
		goto L55
	}
L53:
	;
	v339 = *(*int64)(unsafe.Add(mBase, uint32(v335)+40))
	v346 = v339
	goto L51
L54:
	;
	v343 = F_hashtableSize(m, v341)
	mBase = m.M
	v346 = base.I64_extend_i32_u(v343)
	goto L51
L55:
	;
	v346 = int64(0)
	goto L51
L56:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v112)+124)) = v357
	v360 = v357 + v354
	*(*int32)(unsafe.Add(mBase, uint32(v360))) = v317
	v362 = *(*int32)(unsafe.Add(mBase, uint32(v332)))
	v364 = *(*int32)(unsafe.Add(mBase, uint32(v362)+60))
	v365 = *(*int32)(unsafe.Add(mBase, uint32(v362)+20))
	v366 = *(*int32)(unsafe.Add(mBase, uint32(v365)+20))
	v371 = v364 + v366*int32(12) + int32(80)
	v372 = *(*int32)(unsafe.Add(mBase, uint32(v362)+56))
	if v372 == int32(0) {
		v381 = v371
		goto L58
	} else {
		goto L59
	}
L57:
	;
	v382 = int32(12)
	v384 = v381 + v348*v382
	*(*int32)(unsafe.Add(mBase, uint32(v360)+4)) = v384
	v386 = *(*int32)(unsafe.Add(mBase, uint32(v332)+4))
	v388 = *(*int32)(unsafe.Add(mBase, uint32(v386)+60))
	v389 = *(*int32)(unsafe.Add(mBase, uint32(v386)+20))
	v390 = *(*int32)(unsafe.Add(mBase, uint32(v389)+20))
	v395 = v388 + v390*v382 + int32(80)
	v396 = *(*int32)(unsafe.Add(mBase, uint32(v386)+56))
	if v396 == int32(0) {
		v405 = v395
		goto L61
	} else {
		goto L62
	}
L58:
	;
	goto L57
L59:
	;
	v375 = *(*int32)(unsafe.Add(mBase, uint32(v362)+12))
	v381 = v371 + v375<<(uint(int32(3))%32) + int32(8)
	goto L58
L60:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v360)+8)) = v405
	*(*int32)(unsafe.Add(mBase, uint32(v112)+108)) = v352 + int32(1)
	v410 = *(*int32)(unsafe.Add(mBase, uint32(v332)))
	v411 = *(*int32)(unsafe.Add(mBase, uint32(v410)+60))
	goto L63
L61:
	;
	goto L60
L62:
	;
	v399 = *(*int32)(unsafe.Add(mBase, uint32(v386)+12))
	v405 = v395 + v399<<(uint(int32(3))%32) + int32(8)
	goto L61
L63:
	;
	v412 = *(*int32)(unsafe.Add(mBase, uint32(v112)+112))
	v413 = *(*int32)(unsafe.Add(mBase, uint32(v332)+4))
	v414 = *(*int32)(unsafe.Add(mBase, uint32(v413)+60))
	goto L64
L64:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v112)+112)) = v414 + (v412 + v411)
	v418 = *(*int32)(unsafe.Add(mBase, uint32(v332)))
	v419 = *(*int32)(unsafe.Add(mBase, uint32(v418)+64))
	goto L65
L65:
	;
	v420 = *(*int32)(unsafe.Add(mBase, uint32(v112)+116))
	v421 = *(*int32)(unsafe.Add(mBase, uint32(v332)+4))
	v422 = *(*int32)(unsafe.Add(mBase, uint32(v421)+64))
	goto L66
L66:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v112)+116)) = v422 + (v420 + v419)
	v426 = *(*int32)(unsafe.Add(mBase, uint32(v332)))
	v427 = *(*int32)(unsafe.Add(mBase, uint32(v426)+20))
	v428 = *(*int32)(unsafe.Add(mBase, uint32(v427)+20))
	goto L67
L67:
	;
	v429 = *(*int32)(unsafe.Add(mBase, uint32(v112)+120))
	v430 = *(*int32)(unsafe.Add(mBase, uint32(v332)+4))
	v431 = *(*int32)(unsafe.Add(mBase, uint32(v430)+20))
	v432 = *(*int32)(unsafe.Add(mBase, uint32(v431)+20))
	goto L68
L68:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v112)+120)) = v432 + (v429 + v428)
	v438 = int32(_a_F_getMemoryOverheadData_0)
	v439 = *(*int32)(unsafe.Add(mBase, _c_F_getMemoryOverheadData[31]))
	v441 = *(*int32)(unsafe.Add(mBase, _c_F_getMemoryOverheadData[32]))
	v442 = v439
	v444 = v441
	v445 = v405 + (v384 + v323)
	goto L49
L69:
	;
	goto L48
L70:
	;
	v484 = base.F32_convert_i32_u(v104 - v118)
	goto L72
L71:
	;
	v484 = float32(1)
	goto L72
L72:
	;
	*(*float32)(unsafe.Add(mBase, uint32(v112)+68)) = base.F32_div(base.F32_mul(base.F32_convert_i32_u(v469), v472), v484)
	v487 = *(*int32)(unsafe.Add(mBase, uint32(v112)+60))
	if v487 != 0 {
		goto L73
	} else {
		goto L74
	}
L73:
	;
	v491 = base.I32_div_u_s(v469, v487)
	*(*int32)(unsafe.Add(mBase, uint32(v112)+64)) = v491
	return v112
L74:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v112)+64)) = int32(0)
	return v112
}
func F_getMonotonicUs_posix(m *base.Module) int64 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v10 int64
	_ = v10
	var v11 int32
	_ = v11
	var v18 int32
	_ = v18
	v4 = m.G0
	v5 = int32(16)
	v6 = v4 - v5
	m.G0 = v6
	v9 = F___clock_gettime(m, int32(1), v6)
	mBase = m.M
	v10 = *(*int64)(unsafe.Add(mBase, uint32(v6)))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(v6)+8))
	m.G0 = v6 + v5
	v18 = base.I32_div_s(v11, int32(1000))
	return v10*int64(1000000) + base.I64_extend_i32_s(v18)
}
func F_getPausedReason(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	if base.Ui32(int32(4)) < base.Ui32(l0) {
		v11 = int32(_a_F_getPausedReason_0)
	} else {
		v10 = *(*int32)(unsafe.Add(mBase, uint32(l0<<(uint(int32(2))%32))+uint32(_c_F_getPausedReason[0])))
		v11 = v10
	}
	return v11
}
func F_getRangeLongFromObjectOrReply(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32 {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v26 int64
	_ = v26
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
	var v53 int32
	_ = v53
	v10 = m.G0
	v12 = v10 - int32(16)
	m.G0 = v12
	v16 = F_getLongLongFromObject(m, l1, v12+int32(8))
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		return int32(0)
	} else {
		if v16 == int32(0) {
			v26 = *(*int64)(unsafe.Add(mBase, uint32(v12)+8))
			if base.Ui64(int64(-4294967297)) < base.Ui64(v26+int64(-2147483648)) {
				v34 = base.I32_wrap_i64(v26)
				*(*int32)(unsafe.Add(mBase, uint32(l4))) = v34
				if v34 < l2 {
					if l5 != 0 {
						F_addReplyError(m, l0, l5)
						mBase = m.M
						v48 = m.ExcPending
						if v48 != 0 {
							return int32(0)
						} else {
							v53 = int32(-1)
							m.G0 = v12 + int32(16)
							return v53
						}
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v12)+4)) = l3
						*(*int32)(unsafe.Add(mBase, uint32(v12))) = l2
						F_addReplyErrorFormat(m, l0, int32(_a_F_getRangeLongFromObjectOrReply_0), v12)
						mBase = m.M
						v44 = m.ExcPending
						if v44 != 0 {
							return int32(0)
						} else {
							v53 = int32(-1)
							m.G0 = v12 + int32(16)
							return v53
						}
					}
				} else {
					if v34 <= l3 {
						v53 = int32(0)
						m.G0 = v12 + int32(16)
						return v53
					} else {
						if l5 != 0 {
							F_addReplyError(m, l0, l5)
							mBase = m.M
							v48 = m.ExcPending
							if v48 != 0 {
								return int32(0)
							} else {
								v53 = int32(-1)
								m.G0 = v12 + int32(16)
								return v53
							}
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v12)+4)) = l3
							*(*int32)(unsafe.Add(mBase, uint32(v12))) = l2
							F_addReplyErrorFormat(m, l0, int32(_a_F_getRangeLongFromObjectOrReply_0), v12)
							mBase = m.M
							v44 = m.ExcPending
							if v44 != 0 {
								return int32(0)
							} else {
								v53 = int32(-1)
								m.G0 = v12 + int32(16)
								return v53
							}
						}
					}
				}
			} else {
				if l5 != 0 {
					F_addReplyError(m, l0, l5)
					mBase = m.M
					v48 = m.ExcPending
					if v48 != 0 {
						return int32(0)
					} else {
						v53 = int32(-1)
						m.G0 = v12 + int32(16)
						return v53
					}
				} else {
					F_addReplyError(m, l0, int32(_a_F_getRangeLongFromObjectOrReply_1))
					mBase = m.M
					v33 = m.ExcPending
					if v33 != 0 {
						return int32(0)
					} else {
						v53 = int32(-1)
						m.G0 = v12 + int32(16)
						return v53
					}
				}
			}
		} else {
			if l5 != 0 {
				v23 = l5
			} else {
				v23 = int32(_a_F_getRangeLongFromObjectOrReply_2)
			}
			F_addReplyError(m, l0, v23)
			mBase = m.M
			v25 = m.ExcPending
			if v25 != 0 {
				return int32(0)
			} else {
				v53 = int32(-1)
				m.G0 = v12 + int32(16)
				return v53
			}
		}
	}
}
func F_getfunc(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v11 int32
	_ = v11
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v124 int32
	_ = v124
	var v125 int64
	_ = v125
	var v127 int32
	_ = v127
	var v129 int32
	_ = v129
	var v135 int32
	_ = v135
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v146 int32
	_ = v146
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v156 int32
	_ = v156
	var v159 int32
	_ = v159
	var v162 int32
	_ = v162
	var v164 int32
	_ = v164
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v175 int32
	_ = v175
	var v177 int32
	_ = v177
	var v181 int32
	_ = v181
	var v183 int32
	_ = v183
	var v188 int32
	_ = v188
	var v192 int32
	_ = v192
	var v194 int32
	_ = v194
	var v203 int32
	_ = v203
	var v205 int32
	_ = v205
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v231 int32
	_ = v231
	var v234 int32
	_ = v234
	var v268 int32
	_ = v268
	var v271 int32
	_ = v271
	var v274 int32
	_ = v274
	var v276 int32
	_ = v276
	var v279 int32
	_ = v279
	var v280 int32
	_ = v280
	v4 = m.G0
	v6 = v4 - int32(112)
	m.G0 = v6
	goto L6
L1:
	;
	m.G0 = v6 + int32(112)
	return
L2:
	;
	if l1 == int32(0) {
		goto L36
	} else {
		goto L37
	}
L3:
	;
	if v65 != int32(6) {
		goto L2
	} else {
		goto L18
	}
L4:
	;
	v59 = m.G398
	if v16 != v59 {
		goto L16
	} else {
		goto L17
	}
L6:
	;
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v16 = v11 + int32(0)
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if base.Ui32(v16) < base.Ui32(v17) {
		goto L4
	} else {
		goto L7
	}
L7:
	;
	v65 = int32(-1)
	goto L3
L16:
	;
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v16)+8))
	v65 = v62
	goto L3
L17:
	;
	v65 = int32(-1)
	goto L3
L18:
	;
	goto L22
L19:
	;
	goto L1
L20:
	;
	v124 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v125 = *(*int64)(unsafe.Add(mBase, uint32(v82)))
	*(*int64)(unsafe.Add(mBase, uint32(v124))) = v125
	v127 = *(*int32)(unsafe.Add(mBase, uint32(v82)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v124)+8)) = v127
	v129 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v129 + int32(16)
	goto L19
L22:
	;
	v73 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v74 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v79 = v74 + int32(0)
	v80 = m.G398
	if base.Ui32(v79) < base.Ui32(v73) {
		goto L23
	} else {
		goto L24
	}
L23:
	;
	v82 = v79
	goto L25
L24:
	;
	v82 = v80
	goto L25
L25:
	;
	goto L20
L35:
	;
	if int32(-1) < v142 {
		goto L41
	} else {
		goto L42
	}
L36:
	;
	v140 = F_luaL_checkinteger(m, l0, int32(1))
	mBase = m.M
	v141 = m.ExcPending
	if v141 != 0 {
		goto L38
	} else {
		goto L40
	}
L37:
	;
	v135 = int32(1)
	v137 = F_luaL_optinteger(m, l0, v135, v135)
	mBase = m.M
	v138 = m.ExcPending
	if v138 != 0 {
		goto L38
	} else {
		goto L39
	}
L38:
	;
	return
L39:
	;
	v142 = v137
	goto L35
L40:
	;
	v142 = v140
	goto L35
L41:
	;
	v156 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v142 < int32(1) {
		v181 = v142
		v183 = v156
		goto L47
	} else {
		goto L48
	}
L42:
	;
	v146 = m.G3
	v149 = F_luaL_argerror(m, l0, int32(1), v146+int32(_a_F_getfunc_0))
	mBase = m.M
	v150 = m.ExcPending
	if v150 != 0 {
		goto L38
	} else {
		goto L43
	}
L43:
	;
	goto L41
L44:
	;
	v210 = m.G3
	v215 = F_lua_getinfo(m, l0, v210+int32(_a_F_getfunc_1), v6+int32(12))
	mBase = m.M
	v216 = m.ExcPending
	if v216 != 0 {
		goto L38
	} else {
		goto L60
	}
L45:
	;
	if v203 != 0 {
		goto L44
	} else {
		goto L58
	}
L46:
	;
	goto L45
L47:
	;
	if v181 != 0 {
		v194 = int32(0)
		goto L55
	} else {
		goto L56
	}
L48:
	;
	v159 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v162 = v142
	v164 = v156
	goto L49
L49:
	;
	if base.Ui32(v164) <= base.Ui32(v159) {
		v203 = int32(0)
		goto L46
	} else {
		goto L51
	}
L50:
	;
	v181 = v175
	v183 = v177
	goto L47
L51:
	;
	v169 = v162 + int32(-1)
	v170 = *(*int32)(unsafe.Add(mBase, uint32(v164)+4))
	v171 = *(*int32)(unsafe.Add(mBase, uint32(v170)))
	v172 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v171)+6)))
	if v172 != 0 {
		v175 = v169
		goto L52
	} else {
		goto L53
	}
L52:
	;
	v177 = v164 + int32(-24)
	if int32(0) < v175 {
		v162 = v175
		v164 = v177
		goto L49
	} else {
		goto L54
	}
L53:
	;
	v173 = *(*int32)(unsafe.Add(mBase, uint32(v164)+20))
	v175 = v169 - v173
	goto L52
L54:
	;
	goto L50
L55:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v6+int32(12))+96)) = v194
	v203 = int32(1)
	goto L46
L56:
	;
	v188 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	if base.Ui32(v183) <= base.Ui32(v188) {
		v203 = int32(0)
		goto L46
	} else {
		goto L57
	}
L57:
	;
	v192 = base.I32_div_s(v183-v188, int32(24))
	v194 = v192
	goto L55
L58:
	;
	v205 = m.G3
	v208 = F_luaL_argerror(m, l0, int32(1), v205+int32(_a_F_getfunc_2))
	mBase = m.M
	v209 = m.ExcPending
	if v209 != 0 {
		goto L38
	} else {
		goto L59
	}
L59:
	;
	goto L44
L60:
	;
	goto L63
L61:
	;
	if v274 != 0 {
		goto L1
	} else {
		goto L76
	}
L62:
	;
	v268 = m.G398
	if v234 != v268 {
		goto L74
	} else {
		goto L75
	}
L63:
	;
	goto L67
L67:
	;
	v231 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v234 = v231 + int32(-16)
	goto L62
L74:
	;
	v271 = *(*int32)(unsafe.Add(mBase, uint32(v234)+8))
	v274 = v271
	goto L61
L75:
	;
	v274 = int32(-1)
	goto L61
L76:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v6))) = v142
	v276 = m.G3
	v279 = F_luaL_error(m, l0, v276+int32(_a_F_getfunc_3), v6)
	mBase = m.M
	v280 = m.ExcPending
	if v280 != 0 {
		goto L38
	} else {
		goto L77
	}
L77:
	;
	goto L1
}
func F_getgrnam(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	*(*int32)(unsafe.Add(mBase, _c_F_getgrnam[0])) = int32(44)
	return int32(0)
}
func F_getn(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v17 int32
	_ = v17
	F_luaL_checktype(m, l0, int32(1), int32(5))
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		v9 = F_lua_objlen(m, l0, int32(1))
		mBase = m.M
		v10 = m.ExcPending
		if v10 != 0 {
			return int32(0)
		} else {
			v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			*(*int32)(unsafe.Add(mBase, uint32(v12)+8)) = int32(3)
			*(*float64)(unsafe.Add(mBase, uint32(v12))) = base.F64_convert_i32_s(v9)
			v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v17 + int32(16)
			return int32(1)
		}
	}
}
func F_getrusage(m *base.Module, l0 int32, l1 int32) int32 {
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
	var v16 int32
	_ = v16
	var v21 int32
	_ = v21
	var v23 int64
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v29 int64
	_ = v29
	var v30 int32
	_ = v30
	v6 = m.G0
	v7 = int32(16)
	v8 = v6 - v7
	m.G0 = v8
	v11 = l1 + v7
	if v11 == int32(0) {
	} else {
		v16 = F___memset(m, v11, int32(0), int32(152))
		mBase = m.M
	}
	v21 = F__emscripten_memcpy_bulkmem(m, v8, v11, int32(16))
	mBase = m.M
	v23 = int64(*(*int32)(unsafe.Add(mBase, uint32(v8))))
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v8)+4))
	v25 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l1)+12)) = v25
	*(*int32)(unsafe.Add(mBase, uint32(l1)+8)) = v24
	*(*int64)(unsafe.Add(mBase, uint32(l1))) = v23
	v29 = int64(*(*int32)(unsafe.Add(mBase, uint32(v8)+8)))
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v8)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l1)+28)) = v25
	*(*int32)(unsafe.Add(mBase, uint32(l1)+24)) = v30
	*(*int64)(unsafe.Add(mBase, uint32(l1)+16)) = v29
	m.G0 = v8 + int32(16)
	return int32(0)
}
func F_globfree(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v17 int32
	_ = v17
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v77 int32
	_ = v77
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v88 int32
	_ = v88
	var __phi88 int32
	_ = __phi88
	var v89 int32
	_ = v89
	var __phi89 int32
	_ = __phi89
	var v95 int32
	_ = v95
	var v98 int32
	_ = v98
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v120 int32
	_ = v120
	var v129 int32
	_ = v129
	var v131 int32
	_ = v131
	var v134 int32
	_ = v134
	var v137 int32
	_ = v137
	var v139 int32
	_ = v139
	var v144 int32
	_ = v144
	var v151 int32
	_ = v151
	var v156 int32
	_ = v156
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v171 int32
	_ = v171
	var v179 int32
	_ = v179
	var v181 int32
	_ = v181
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v192 int32
	_ = v192
	var v194 int32
	_ = v194
	var v201 int32
	_ = v201
	var v203 int32
	_ = v203
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v221 int32
	_ = v221
	var v223 int32
	_ = v223
	var v225 int32
	_ = v225
	var v234 int32
	_ = v234
	var v236 int32
	_ = v236
	var v239 int32
	_ = v239
	var v244 int32
	_ = v244
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v255 int32
	_ = v255
	var __phi255 int32
	_ = __phi255
	var v256 int32
	_ = v256
	var __phi256 int32
	_ = __phi256
	var v262 int32
	_ = v262
	var v265 int32
	_ = v265
	var v279 int32
	_ = v279
	var v288 int32
	_ = v288
	var v290 int32
	_ = v290
	var v293 int32
	_ = v293
	var v296 int32
	_ = v296
	var v298 int32
	_ = v298
	var v303 int32
	_ = v303
	var v310 int32
	_ = v310
	var v315 int32
	_ = v315
	var v337 int32
	_ = v337
	var v349 int32
	_ = v349
	var v351 int32
	_ = v351
	var v353 int32
	_ = v353
	var v357 int32
	_ = v357
	var v362 int32
	_ = v362
	var v363 int32
	_ = v363
	var v374 int32
	_ = v374
	var v377 int32
	_ = v377
	var v384 int32
	_ = v384
	var v389 int32
	_ = v389
	var v393 int32
	_ = v393
	var v395 int32
	_ = v395
	var v410 int32
	_ = v410
	var v412 int32
	_ = v412
	var v415 int32
	_ = v415
	var v418 int32
	_ = v418
	var v422 int32
	_ = v422
	var v432 int32
	_ = v432
	var v433 int32
	_ = v433
	var v439 int32
	_ = v439
	var v441 int32
	_ = v441
	var v444 int32
	_ = v444
	var v448 int32
	_ = v448
	var v454 int32
	_ = v454
	var v456 int32
	_ = v456
	var v457 int32
	_ = v457
	var v458 int32
	_ = v458
	var v459 int32
	_ = v459
	var v468 int32
	_ = v468
	var v470 int32
	_ = v470
	var v471 int32
	_ = v471
	var v472 int32
	_ = v472
	var v474 int32
	_ = v474
	var v495 int32
	_ = v495
	var v496 int32
	_ = v496
	var v500 int32
	_ = v500
	var v511 int32
	_ = v511
	var v512 int32
	_ = v512
	var v515 int32
	_ = v515
	var v517 int32
	_ = v517
	var v518 int32
	_ = v518
	var v525 int32
	_ = v525
	var v526 int32
	_ = v526
	var v528 int32
	_ = v528
	var v530 int32
	_ = v530
	var v532 int32
	_ = v532
	var v534 int32
	_ = v534
	var v537 int32
	_ = v537
	var v539 int32
	_ = v539
	var v541 int32
	_ = v541
	var v548 int32
	_ = v548
	var v550 int32
	_ = v550
	var v553 int32
	_ = v553
	var v558 int32
	_ = v558
	var v563 int32
	_ = v563
	var v564 int32
	_ = v564
	var v569 int32
	_ = v569
	var __phi569 int32
	_ = __phi569
	var v570 int32
	_ = v570
	var __phi570 int32
	_ = __phi570
	var v576 int32
	_ = v576
	var v579 int32
	_ = v579
	var v582 int32
	_ = v582
	var v583 int32
	_ = v583
	var v601 int32
	_ = v601
	var v610 int32
	_ = v610
	var v612 int32
	_ = v612
	var v615 int32
	_ = v615
	var v618 int32
	_ = v618
	var v620 int32
	_ = v620
	var v625 int32
	_ = v625
	var v632 int32
	_ = v632
	var v637 int32
	_ = v637
	var v642 int32
	_ = v642
	var v643 int32
	_ = v643
	var v652 int32
	_ = v652
	var v660 int32
	_ = v660
	var v662 int32
	_ = v662
	var v666 int32
	_ = v666
	var v667 int32
	_ = v667
	var v673 int32
	_ = v673
	var v675 int32
	_ = v675
	var v682 int32
	_ = v682
	var v684 int32
	_ = v684
	var v688 int32
	_ = v688
	var v689 int32
	_ = v689
	var v698 int32
	_ = v698
	var v699 int32
	_ = v699
	var v702 int32
	_ = v702
	var v704 int32
	_ = v704
	var v706 int32
	_ = v706
	var v715 int32
	_ = v715
	var v717 int32
	_ = v717
	var v720 int32
	_ = v720
	var v725 int32
	_ = v725
	var v730 int32
	_ = v730
	var v731 int32
	_ = v731
	var v736 int32
	_ = v736
	var __phi736 int32
	_ = __phi736
	var v737 int32
	_ = v737
	var __phi737 int32
	_ = __phi737
	var v743 int32
	_ = v743
	var v746 int32
	_ = v746
	var v760 int32
	_ = v760
	var v769 int32
	_ = v769
	var v771 int32
	_ = v771
	var v774 int32
	_ = v774
	var v777 int32
	_ = v777
	var v779 int32
	_ = v779
	var v784 int32
	_ = v784
	var v791 int32
	_ = v791
	var v796 int32
	_ = v796
	var v818 int32
	_ = v818
	var v830 int32
	_ = v830
	var v832 int32
	_ = v832
	var v834 int32
	_ = v834
	var v838 int32
	_ = v838
	var v843 int32
	_ = v843
	var v844 int32
	_ = v844
	var v855 int32
	_ = v855
	var v858 int32
	_ = v858
	var v865 int32
	_ = v865
	var v870 int32
	_ = v870
	var v874 int32
	_ = v874
	var v876 int32
	_ = v876
	var v891 int32
	_ = v891
	var v893 int32
	_ = v893
	var v896 int32
	_ = v896
	var v899 int32
	_ = v899
	var v903 int32
	_ = v903
	var v913 int32
	_ = v913
	var v914 int32
	_ = v914
	var v920 int32
	_ = v920
	var v922 int32
	_ = v922
	var v925 int32
	_ = v925
	var v929 int32
	_ = v929
	var v935 int32
	_ = v935
	var v937 int32
	_ = v937
	var v938 int32
	_ = v938
	var v939 int32
	_ = v939
	var v940 int32
	_ = v940
	var v949 int32
	_ = v949
	var v951 int32
	_ = v951
	var v952 int32
	_ = v952
	var v953 int32
	_ = v953
	var v955 int32
	_ = v955
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v3 == int32(0) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v500 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v500 == int32(0) {
		goto L107
	} else {
		goto L108
	}
L2:
	;
	v8 = int32(0)
	goto L3
L3:
	;
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v11 = int32(2)
	v17 = *(*int32)(unsafe.Add(mBase, uint32(v9+v10<<(uint(v11)%32)+v8<<(uint(v11)%32))))
	if v17+int32(-4) == int32(0) {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	goto L1
L5:
	;
	v495 = v8 + int32(1)
	v496 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if base.Ui32(v495) < base.Ui32(v496) {
		v8 = v495
		goto L3
	} else {
		goto L105
	}
L6:
	;
	goto L5
L7:
	;
	v30 = int32(-8)
	v31 = v17 + int32(-12)
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v17+v30)))
	v36 = v34 & v30
	v37 = v31 + v36
	if v34&int32(1) != 0 {
		v161 = v36
		v162 = v31
		goto L8
	} else {
		goto L9
	}
L8:
	;
	if base.Ui32(v37) <= base.Ui32(v162) {
		goto L6
	} else {
		goto L43
	}
L9:
	;
	if v34&int32(2) == int32(0) {
		goto L6
	} else {
		goto L10
	}
L10:
	;
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v31)))
	v45 = v31 - v44
	v47 = *(*int32)(unsafe.Add(mBase, _c_F_globfree[0]))
	if base.Ui32(v45) < base.Ui32(v47) {
		goto L6
	} else {
		goto L11
	}
L11:
	;
	v49 = v44 + v36
	v51 = *(*int32)(unsafe.Add(mBase, _c_F_globfree[1]))
	if v45 == v51 {
		goto L15
	} else {
		goto L16
	}
L12:
	;
	if v67 == int32(0) {
		v161 = v49
		v162 = v45
		goto L8
	} else {
		goto L31
	}
L13:
	;
	v120 = int32(0)
	goto L12
L14:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v56)+12)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v53)+8)) = v56
	v161 = v49
	v162 = v45
	goto L8
L15:
	;
	v101 = *(*int32)(unsafe.Add(mBase, uint32(v37)+4))
	v102 = int32(3)
	if v101&v102 != v102 {
		v161 = v49
		v162 = v45
		goto L8
	} else {
		goto L30
	}
L16:
	;
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v45)+12))
	if base.Ui32(int32(255)) < base.Ui32(v44) {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v45)+24))
	if v53 == v45 {
		goto L20
	} else {
		goto L21
	}
L18:
	;
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v45)+8))
	if v53 != v56 {
		goto L14
	} else {
		goto L19
	}
L19:
	;
	v58 = int32(0)
	v60 = *(*int32)(unsafe.Add(mBase, _c_F_globfree[2]))
	*(*int32)(unsafe.Add(mBase, _c_F_globfree[2])) = v60 & base.I32_rotl(int32(-2), int32(base.Ui32(v44)>>(uint(int32(3))%32)))
	v161 = v49
	v162 = v45
	goto L8
L20:
	;
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v45)+20))
	if v72 == int32(0) {
		goto L23
	} else {
		goto L24
	}
L21:
	;
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v45)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v69)+12)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v53)+8)) = v69
	v120 = v53
	goto L12
L22:
	;
	__phi88 = v82
	__phi89 = v83
	v88 = __phi88
	v89 = __phi89
	goto L26
L23:
	;
	v77 = *(*int32)(unsafe.Add(mBase, uint32(v45)+16))
	if v77 == int32(0) {
		goto L13
	} else {
		goto L25
	}
L24:
	;
	v82 = v72
	v83 = v45 + int32(20)
	goto L22
L25:
	;
	v82 = v77
	v83 = v45 + int32(16)
	goto L22
L26:
	;
	v95 = *(*int32)(unsafe.Add(mBase, uint32(v88)+20))
	if v95 != 0 {
		__phi88 = v95
		__phi89 = v88 + int32(20)
		v88 = __phi88
		v89 = __phi89
		goto L26
	} else {
		goto L28
	}
L27:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v89))) = int32(0)
	v120 = v88
	goto L12
L28:
	;
	v98 = *(*int32)(unsafe.Add(mBase, uint32(v88)+16))
	if v98 != 0 {
		__phi88 = v98
		__phi89 = v88 + int32(16)
		v88 = __phi88
		v89 = __phi89
		goto L26
	} else {
		goto L29
	}
L29:
	;
	goto L27
L30:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_globfree[3])) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v37)+4)) = v101 & int32(-2)
	*(*int32)(unsafe.Add(mBase, uint32(v45)+4)) = v49 | int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v37))) = v49
	goto L5
L31:
	;
	v129 = *(*int32)(unsafe.Add(mBase, uint32(v45)+28))
	v131 = v129 << (uint(int32(2)) % 32)
	v134 = *(*int32)(unsafe.Add(mBase, uint32(v131)+uint32(_c_F_globfree[4])))
	if v45 != v134 {
		goto L33
	} else {
		goto L34
	}
L32:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v120)+24)) = v67
	v151 = *(*int32)(unsafe.Add(mBase, uint32(v45)+16))
	if v151 == int32(0) {
		goto L40
	} else {
		goto L41
	}
L33:
	;
	v144 = *(*int32)(unsafe.Add(mBase, uint32(v67)+16))
	if v144 != v45 {
		goto L37
	} else {
		goto L38
	}
L34:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v131)+uint32(_c_F_globfree[4]))) = v120
	if v120 != 0 {
		goto L32
	} else {
		goto L35
	}
L35:
	;
	v137 = int32(0)
	v139 = *(*int32)(unsafe.Add(mBase, _c_F_globfree[5]))
	*(*int32)(unsafe.Add(mBase, _c_F_globfree[5])) = v139 & base.I32_rotl(int32(-2), v129)
	v161 = v49
	v162 = v45
	goto L8
L36:
	;
	if v120 == int32(0) {
		v161 = v49
		v162 = v45
		goto L8
	} else {
		goto L39
	}
L37:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v67)+20)) = v120
	goto L36
L38:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v67)+16)) = v120
	goto L36
L39:
	;
	goto L32
L40:
	;
	v156 = *(*int32)(unsafe.Add(mBase, uint32(v45)+20))
	if v156 == int32(0) {
		v161 = v49
		v162 = v45
		goto L8
	} else {
		goto L42
	}
L41:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v120)+16)) = v151
	*(*int32)(unsafe.Add(mBase, uint32(v151)+24)) = v120
	goto L40
L42:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v120)+20)) = v156
	*(*int32)(unsafe.Add(mBase, uint32(v156)+24)) = v120
	v161 = v49
	v162 = v45
	goto L8
L43:
	;
	v171 = *(*int32)(unsafe.Add(mBase, uint32(v37)+4))
	if v171&int32(1) == int32(0) {
		goto L6
	} else {
		goto L44
	}
L44:
	;
	if v171&int32(2) != 0 {
		goto L49
	} else {
		goto L50
	}
L45:
	;
	if base.Ui32(int32(255)) < base.Ui32(v337) {
		goto L83
	} else {
		goto L84
	}
L46:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v162)+4)) = v217 | int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v162+v217))) = v217
	if v162 != v201 {
		v337 = v217
		goto L45
	} else {
		goto L82
	}
L47:
	;
	if v234 == int32(0) {
		goto L46
	} else {
		goto L70
	}
L48:
	;
	v279 = int32(0)
	goto L47
L49:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v37)+4)) = v171 & int32(-2)
	*(*int32)(unsafe.Add(mBase, uint32(v162)+4)) = v161 | int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v162+v161))) = v161
	v337 = v161
	goto L45
L50:
	;
	v179 = *(*int32)(unsafe.Add(mBase, _c_F_globfree[6]))
	if v37 != v179 {
		goto L51
	} else {
		goto L52
	}
L51:
	;
	v201 = *(*int32)(unsafe.Add(mBase, _c_F_globfree[1]))
	if v37 != v201 {
		goto L54
	} else {
		goto L55
	}
L52:
	;
	v181 = int32(0)
	*(*int32)(unsafe.Add(mBase, _c_F_globfree[6])) = v162
	v185 = *(*int32)(unsafe.Add(mBase, _c_F_globfree[7]))
	v186 = v185 + v161
	*(*int32)(unsafe.Add(mBase, _c_F_globfree[7])) = v186
	*(*int32)(unsafe.Add(mBase, uint32(v162)+4)) = v186 | int32(1)
	v192 = *(*int32)(unsafe.Add(mBase, _c_F_globfree[1]))
	if v162 != v192 {
		goto L6
	} else {
		goto L53
	}
L53:
	;
	v194 = int32(0)
	*(*int32)(unsafe.Add(mBase, _c_F_globfree[3])) = v194
	*(*int32)(unsafe.Add(mBase, _c_F_globfree[1])) = v194
	goto L5
L54:
	;
	v217 = v171&int32(-8) + v161
	v218 = *(*int32)(unsafe.Add(mBase, uint32(v37)+12))
	if base.Ui32(int32(255)) < base.Ui32(v171) {
		goto L56
	} else {
		goto L57
	}
L55:
	;
	v203 = int32(0)
	*(*int32)(unsafe.Add(mBase, _c_F_globfree[1])) = v162
	v207 = *(*int32)(unsafe.Add(mBase, _c_F_globfree[3]))
	v208 = v207 + v161
	*(*int32)(unsafe.Add(mBase, _c_F_globfree[3])) = v208
	*(*int32)(unsafe.Add(mBase, uint32(v162)+4)) = v208 | int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v162+v208))) = v208
	goto L5
L56:
	;
	v234 = *(*int32)(unsafe.Add(mBase, uint32(v37)+24))
	if v218 == v37 {
		goto L60
	} else {
		goto L61
	}
L57:
	;
	v221 = *(*int32)(unsafe.Add(mBase, uint32(v37)+8))
	if v218 != v221 {
		goto L58
	} else {
		goto L59
	}
L58:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v221)+12)) = v218
	*(*int32)(unsafe.Add(mBase, uint32(v218)+8)) = v221
	goto L46
L59:
	;
	v223 = int32(0)
	v225 = *(*int32)(unsafe.Add(mBase, _c_F_globfree[2]))
	*(*int32)(unsafe.Add(mBase, _c_F_globfree[2])) = v225 & base.I32_rotl(int32(-2), int32(base.Ui32(v171)>>(uint(int32(3))%32)))
	goto L46
L60:
	;
	v239 = *(*int32)(unsafe.Add(mBase, uint32(v37)+20))
	if v239 == int32(0) {
		goto L63
	} else {
		goto L64
	}
L61:
	;
	v236 = *(*int32)(unsafe.Add(mBase, uint32(v37)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v236)+12)) = v218
	*(*int32)(unsafe.Add(mBase, uint32(v218)+8)) = v236
	v279 = v218
	goto L47
L62:
	;
	__phi255 = v249
	__phi256 = v250
	v255 = __phi255
	v256 = __phi256
	goto L66
L63:
	;
	v244 = *(*int32)(unsafe.Add(mBase, uint32(v37)+16))
	if v244 == int32(0) {
		goto L48
	} else {
		goto L65
	}
L64:
	;
	v249 = v239
	v250 = v37 + int32(20)
	goto L62
L65:
	;
	v249 = v244
	v250 = v37 + int32(16)
	goto L62
L66:
	;
	v262 = *(*int32)(unsafe.Add(mBase, uint32(v255)+20))
	if v262 != 0 {
		__phi255 = v262
		__phi256 = v255 + int32(20)
		v255 = __phi255
		v256 = __phi256
		goto L66
	} else {
		goto L68
	}
L67:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v256))) = int32(0)
	v279 = v255
	goto L47
L68:
	;
	v265 = *(*int32)(unsafe.Add(mBase, uint32(v255)+16))
	if v265 != 0 {
		__phi255 = v265
		__phi256 = v255 + int32(16)
		v255 = __phi255
		v256 = __phi256
		goto L66
	} else {
		goto L69
	}
L69:
	;
	goto L67
L70:
	;
	v288 = *(*int32)(unsafe.Add(mBase, uint32(v37)+28))
	v290 = v288 << (uint(int32(2)) % 32)
	v293 = *(*int32)(unsafe.Add(mBase, uint32(v290)+uint32(_c_F_globfree[4])))
	if v37 != v293 {
		goto L72
	} else {
		goto L73
	}
L71:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v279)+24)) = v234
	v310 = *(*int32)(unsafe.Add(mBase, uint32(v37)+16))
	if v310 == int32(0) {
		goto L79
	} else {
		goto L80
	}
L72:
	;
	v303 = *(*int32)(unsafe.Add(mBase, uint32(v234)+16))
	if v303 != v37 {
		goto L76
	} else {
		goto L77
	}
L73:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v290)+uint32(_c_F_globfree[4]))) = v279
	if v279 != 0 {
		goto L71
	} else {
		goto L74
	}
L74:
	;
	v296 = int32(0)
	v298 = *(*int32)(unsafe.Add(mBase, _c_F_globfree[5]))
	*(*int32)(unsafe.Add(mBase, _c_F_globfree[5])) = v298 & base.I32_rotl(int32(-2), v288)
	goto L46
L75:
	;
	if v279 == int32(0) {
		goto L46
	} else {
		goto L78
	}
L76:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v234)+20)) = v279
	goto L75
L77:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v234)+16)) = v279
	goto L75
L78:
	;
	goto L71
L79:
	;
	v315 = *(*int32)(unsafe.Add(mBase, uint32(v37)+20))
	if v315 == int32(0) {
		goto L46
	} else {
		goto L81
	}
L80:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v279)+16)) = v310
	*(*int32)(unsafe.Add(mBase, uint32(v310)+24)) = v279
	goto L79
L81:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v279)+20)) = v315
	*(*int32)(unsafe.Add(mBase, uint32(v315)+24)) = v279
	goto L46
L82:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_globfree[3])) = v217
	goto L5
L83:
	;
	if base.Ui32(int32(16777215)) < base.Ui32(v337) {
		v384 = int32(31)
		goto L88
	} else {
		goto L89
	}
L84:
	;
	v349 = v337 & int32(-8)
	v351 = v349 + int32(9128464)
	v353 = *(*int32)(unsafe.Add(mBase, _c_F_globfree[2]))
	v357 = int32(1) << (uint(int32(base.Ui32(v337)>>(uint(int32(3))%32))) % 32)
	if v353&v357 != 0 {
		goto L86
	} else {
		goto L87
	}
L85:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v349)+uint32(_c_F_globfree[8]))) = v162
	*(*int32)(unsafe.Add(mBase, uint32(v363)+12)) = v162
	*(*int32)(unsafe.Add(mBase, uint32(v162)+12)) = v351
	*(*int32)(unsafe.Add(mBase, uint32(v162)+8)) = v363
	goto L5
L86:
	;
	v362 = *(*int32)(unsafe.Add(mBase, uint32(v349)+uint32(_c_F_globfree[8])))
	v363 = v362
	goto L85
L87:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_globfree[2])) = v353 | v357
	v363 = v351
	goto L85
L88:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v162)+28)) = v384
	*(*int64)(unsafe.Add(mBase, uint32(v162)+16)) = int64(0)
	v389 = v384 << (uint(int32(2)) % 32)
	v393 = *(*int32)(unsafe.Add(mBase, _c_F_globfree[5]))
	v395 = int32(1) << (uint(v384) % 32)
	if v393&v395 != 0 {
		goto L93
	} else {
		goto L94
	}
L89:
	;
	v374 = base.I32_clz(int32(base.Ui32(v337) >> (uint(int32(8)) % 32)))
	v377 = int32(1)
	v384 = int32(base.Ui32(v337)>>(uint(int32(38)-v374)%32))&v377 - v374<<(uint(v377)%32) + int32(62)
	goto L88
L90:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v162+v456))) = v459
	*(*int32)(unsafe.Add(mBase, uint32(v162)+12)) = v458
	*(*int32)(unsafe.Add(mBase, uint32(v162+v454))) = v457
	v468 = int32(0)
	v470 = *(*int32)(unsafe.Add(mBase, _c_F_globfree[9]))
	v471 = int32(-1)
	v472 = v470 + v471
	if v472 != 0 {
		goto L102
	} else {
		goto L103
	}
L91:
	;
	v448 = *(*int32)(unsafe.Add(mBase, uint32(v418)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v448)+12)) = v162
	*(*int32)(unsafe.Add(mBase, uint32(v418)+8)) = v162
	v454 = int32(24)
	v456 = int32(8)
	v457 = int32(0)
	v458 = v418
	v459 = v448
	goto L90
L92:
	;
	v454 = v439
	v456 = v441
	v457 = v162
	v458 = v162
	v459 = v444
	goto L90
L93:
	;
	if v384 == int32(31) {
		goto L95
	} else {
		goto L96
	}
L94:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_globfree[5])) = v393 | v395
	*(*int32)(unsafe.Add(mBase, uint32(v389)+uint32(_c_F_globfree[4]))) = v162
	v439 = int32(8)
	v441 = int32(24)
	v444 = v389 + int32(9128728)
	goto L92
L95:
	;
	v410 = int32(0)
	goto L97
L96:
	;
	v410 = int32(25) - int32(base.Ui32(v384)>>(uint(int32(1))%32))
	goto L97
L97:
	;
	v412 = *(*int32)(unsafe.Add(mBase, uint32(v389)+uint32(_c_F_globfree[4])))
	v415 = v337 << (uint(v410) % 32)
	v418 = v412
	goto L98
L98:
	;
	v422 = *(*int32)(unsafe.Add(mBase, uint32(v418)+4))
	if v422&int32(-8) == v337 {
		goto L91
	} else {
		goto L100
	}
L99:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v432+int32(16)))) = v162
	v439 = int32(8)
	v441 = int32(24)
	v444 = v418
	goto L92
L100:
	;
	v432 = v418 + int32(base.Ui32(v415)>>(uint(int32(29))%32))&int32(4)
	v433 = *(*int32)(unsafe.Add(mBase, uint32(v432)+16))
	if v433 != 0 {
		v415 = v415 << (uint(int32(1)) % 32)
		v418 = v433
		goto L98
	} else {
		goto L101
	}
L101:
	;
	goto L99
L102:
	;
	v474 = v472
	goto L104
L103:
	;
	v474 = v471
	goto L104
L104:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_globfree[9])) = v474
	goto L6
L105:
	;
	goto L4
L106:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l0))) = int64(0)
	return
L107:
	;
	goto L106
L108:
	;
	v511 = int32(-8)
	v512 = v500 + v511
	v515 = *(*int32)(unsafe.Add(mBase, uint32(v500+int32(-4))))
	v517 = v515 & v511
	v518 = v512 + v517
	if v515&int32(1) != 0 {
		v642 = v517
		v643 = v512
		goto L109
	} else {
		goto L110
	}
L109:
	;
	if base.Ui32(v518) <= base.Ui32(v643) {
		goto L107
	} else {
		goto L144
	}
L110:
	;
	if v515&int32(2) == int32(0) {
		goto L107
	} else {
		goto L111
	}
L111:
	;
	v525 = *(*int32)(unsafe.Add(mBase, uint32(v512)))
	v526 = v512 - v525
	v528 = *(*int32)(unsafe.Add(mBase, _c_F_globfree[0]))
	if base.Ui32(v526) < base.Ui32(v528) {
		goto L107
	} else {
		goto L112
	}
L112:
	;
	v530 = v525 + v517
	v532 = *(*int32)(unsafe.Add(mBase, _c_F_globfree[1]))
	if v526 == v532 {
		goto L116
	} else {
		goto L117
	}
L113:
	;
	if v548 == int32(0) {
		v642 = v530
		v643 = v526
		goto L109
	} else {
		goto L132
	}
L114:
	;
	v601 = int32(0)
	goto L113
L115:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v537)+12)) = v534
	*(*int32)(unsafe.Add(mBase, uint32(v534)+8)) = v537
	v642 = v530
	v643 = v526
	goto L109
L116:
	;
	v582 = *(*int32)(unsafe.Add(mBase, uint32(v518)+4))
	v583 = int32(3)
	if v582&v583 != v583 {
		v642 = v530
		v643 = v526
		goto L109
	} else {
		goto L131
	}
L117:
	;
	v534 = *(*int32)(unsafe.Add(mBase, uint32(v526)+12))
	if base.Ui32(int32(255)) < base.Ui32(v525) {
		goto L118
	} else {
		goto L119
	}
L118:
	;
	v548 = *(*int32)(unsafe.Add(mBase, uint32(v526)+24))
	if v534 == v526 {
		goto L121
	} else {
		goto L122
	}
L119:
	;
	v537 = *(*int32)(unsafe.Add(mBase, uint32(v526)+8))
	if v534 != v537 {
		goto L115
	} else {
		goto L120
	}
L120:
	;
	v539 = int32(0)
	v541 = *(*int32)(unsafe.Add(mBase, _c_F_globfree[2]))
	*(*int32)(unsafe.Add(mBase, _c_F_globfree[2])) = v541 & base.I32_rotl(int32(-2), int32(base.Ui32(v525)>>(uint(int32(3))%32)))
	v642 = v530
	v643 = v526
	goto L109
L121:
	;
	v553 = *(*int32)(unsafe.Add(mBase, uint32(v526)+20))
	if v553 == int32(0) {
		goto L124
	} else {
		goto L125
	}
L122:
	;
	v550 = *(*int32)(unsafe.Add(mBase, uint32(v526)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v550)+12)) = v534
	*(*int32)(unsafe.Add(mBase, uint32(v534)+8)) = v550
	v601 = v534
	goto L113
L123:
	;
	__phi569 = v563
	__phi570 = v564
	v569 = __phi569
	v570 = __phi570
	goto L127
L124:
	;
	v558 = *(*int32)(unsafe.Add(mBase, uint32(v526)+16))
	if v558 == int32(0) {
		goto L114
	} else {
		goto L126
	}
L125:
	;
	v563 = v553
	v564 = v526 + int32(20)
	goto L123
L126:
	;
	v563 = v558
	v564 = v526 + int32(16)
	goto L123
L127:
	;
	v576 = *(*int32)(unsafe.Add(mBase, uint32(v569)+20))
	if v576 != 0 {
		__phi569 = v576
		__phi570 = v569 + int32(20)
		v569 = __phi569
		v570 = __phi570
		goto L127
	} else {
		goto L129
	}
L128:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v570))) = int32(0)
	v601 = v569
	goto L113
L129:
	;
	v579 = *(*int32)(unsafe.Add(mBase, uint32(v569)+16))
	if v579 != 0 {
		__phi569 = v579
		__phi570 = v569 + int32(16)
		v569 = __phi569
		v570 = __phi570
		goto L127
	} else {
		goto L130
	}
L130:
	;
	goto L128
L131:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_globfree[3])) = v530
	*(*int32)(unsafe.Add(mBase, uint32(v518)+4)) = v582 & int32(-2)
	*(*int32)(unsafe.Add(mBase, uint32(v526)+4)) = v530 | int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v518))) = v530
	goto L106
L132:
	;
	v610 = *(*int32)(unsafe.Add(mBase, uint32(v526)+28))
	v612 = v610 << (uint(int32(2)) % 32)
	v615 = *(*int32)(unsafe.Add(mBase, uint32(v612)+uint32(_c_F_globfree[4])))
	if v526 != v615 {
		goto L134
	} else {
		goto L135
	}
L133:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v601)+24)) = v548
	v632 = *(*int32)(unsafe.Add(mBase, uint32(v526)+16))
	if v632 == int32(0) {
		goto L141
	} else {
		goto L142
	}
L134:
	;
	v625 = *(*int32)(unsafe.Add(mBase, uint32(v548)+16))
	if v625 != v526 {
		goto L138
	} else {
		goto L139
	}
L135:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v612)+uint32(_c_F_globfree[4]))) = v601
	if v601 != 0 {
		goto L133
	} else {
		goto L136
	}
L136:
	;
	v618 = int32(0)
	v620 = *(*int32)(unsafe.Add(mBase, _c_F_globfree[5]))
	*(*int32)(unsafe.Add(mBase, _c_F_globfree[5])) = v620 & base.I32_rotl(int32(-2), v610)
	v642 = v530
	v643 = v526
	goto L109
L137:
	;
	if v601 == int32(0) {
		v642 = v530
		v643 = v526
		goto L109
	} else {
		goto L140
	}
L138:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v548)+20)) = v601
	goto L137
L139:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v548)+16)) = v601
	goto L137
L140:
	;
	goto L133
L141:
	;
	v637 = *(*int32)(unsafe.Add(mBase, uint32(v526)+20))
	if v637 == int32(0) {
		v642 = v530
		v643 = v526
		goto L109
	} else {
		goto L143
	}
L142:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v601)+16)) = v632
	*(*int32)(unsafe.Add(mBase, uint32(v632)+24)) = v601
	goto L141
L143:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v601)+20)) = v637
	*(*int32)(unsafe.Add(mBase, uint32(v637)+24)) = v601
	v642 = v530
	v643 = v526
	goto L109
L144:
	;
	v652 = *(*int32)(unsafe.Add(mBase, uint32(v518)+4))
	if v652&int32(1) == int32(0) {
		goto L107
	} else {
		goto L145
	}
L145:
	;
	if v652&int32(2) != 0 {
		goto L150
	} else {
		goto L151
	}
L146:
	;
	if base.Ui32(int32(255)) < base.Ui32(v818) {
		goto L184
	} else {
		goto L185
	}
L147:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v643)+4)) = v698 | int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v643+v698))) = v698
	if v643 != v682 {
		v818 = v698
		goto L146
	} else {
		goto L183
	}
L148:
	;
	if v715 == int32(0) {
		goto L147
	} else {
		goto L171
	}
L149:
	;
	v760 = int32(0)
	goto L148
L150:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v518)+4)) = v652 & int32(-2)
	*(*int32)(unsafe.Add(mBase, uint32(v643)+4)) = v642 | int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v643+v642))) = v642
	v818 = v642
	goto L146
L151:
	;
	v660 = *(*int32)(unsafe.Add(mBase, _c_F_globfree[6]))
	if v518 != v660 {
		goto L152
	} else {
		goto L153
	}
L152:
	;
	v682 = *(*int32)(unsafe.Add(mBase, _c_F_globfree[1]))
	if v518 != v682 {
		goto L155
	} else {
		goto L156
	}
L153:
	;
	v662 = int32(0)
	*(*int32)(unsafe.Add(mBase, _c_F_globfree[6])) = v643
	v666 = *(*int32)(unsafe.Add(mBase, _c_F_globfree[7]))
	v667 = v666 + v642
	*(*int32)(unsafe.Add(mBase, _c_F_globfree[7])) = v667
	*(*int32)(unsafe.Add(mBase, uint32(v643)+4)) = v667 | int32(1)
	v673 = *(*int32)(unsafe.Add(mBase, _c_F_globfree[1]))
	if v643 != v673 {
		goto L107
	} else {
		goto L154
	}
L154:
	;
	v675 = int32(0)
	*(*int32)(unsafe.Add(mBase, _c_F_globfree[3])) = v675
	*(*int32)(unsafe.Add(mBase, _c_F_globfree[1])) = v675
	goto L106
L155:
	;
	v698 = v652&int32(-8) + v642
	v699 = *(*int32)(unsafe.Add(mBase, uint32(v518)+12))
	if base.Ui32(int32(255)) < base.Ui32(v652) {
		goto L157
	} else {
		goto L158
	}
L156:
	;
	v684 = int32(0)
	*(*int32)(unsafe.Add(mBase, _c_F_globfree[1])) = v643
	v688 = *(*int32)(unsafe.Add(mBase, _c_F_globfree[3]))
	v689 = v688 + v642
	*(*int32)(unsafe.Add(mBase, _c_F_globfree[3])) = v689
	*(*int32)(unsafe.Add(mBase, uint32(v643)+4)) = v689 | int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v643+v689))) = v689
	goto L106
L157:
	;
	v715 = *(*int32)(unsafe.Add(mBase, uint32(v518)+24))
	if v699 == v518 {
		goto L161
	} else {
		goto L162
	}
L158:
	;
	v702 = *(*int32)(unsafe.Add(mBase, uint32(v518)+8))
	if v699 != v702 {
		goto L159
	} else {
		goto L160
	}
L159:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v702)+12)) = v699
	*(*int32)(unsafe.Add(mBase, uint32(v699)+8)) = v702
	goto L147
L160:
	;
	v704 = int32(0)
	v706 = *(*int32)(unsafe.Add(mBase, _c_F_globfree[2]))
	*(*int32)(unsafe.Add(mBase, _c_F_globfree[2])) = v706 & base.I32_rotl(int32(-2), int32(base.Ui32(v652)>>(uint(int32(3))%32)))
	goto L147
L161:
	;
	v720 = *(*int32)(unsafe.Add(mBase, uint32(v518)+20))
	if v720 == int32(0) {
		goto L164
	} else {
		goto L165
	}
L162:
	;
	v717 = *(*int32)(unsafe.Add(mBase, uint32(v518)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v717)+12)) = v699
	*(*int32)(unsafe.Add(mBase, uint32(v699)+8)) = v717
	v760 = v699
	goto L148
L163:
	;
	__phi736 = v730
	__phi737 = v731
	v736 = __phi736
	v737 = __phi737
	goto L167
L164:
	;
	v725 = *(*int32)(unsafe.Add(mBase, uint32(v518)+16))
	if v725 == int32(0) {
		goto L149
	} else {
		goto L166
	}
L165:
	;
	v730 = v720
	v731 = v518 + int32(20)
	goto L163
L166:
	;
	v730 = v725
	v731 = v518 + int32(16)
	goto L163
L167:
	;
	v743 = *(*int32)(unsafe.Add(mBase, uint32(v736)+20))
	if v743 != 0 {
		__phi736 = v743
		__phi737 = v736 + int32(20)
		v736 = __phi736
		v737 = __phi737
		goto L167
	} else {
		goto L169
	}
L168:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v737))) = int32(0)
	v760 = v736
	goto L148
L169:
	;
	v746 = *(*int32)(unsafe.Add(mBase, uint32(v736)+16))
	if v746 != 0 {
		__phi736 = v746
		__phi737 = v736 + int32(16)
		v736 = __phi736
		v737 = __phi737
		goto L167
	} else {
		goto L170
	}
L170:
	;
	goto L168
L171:
	;
	v769 = *(*int32)(unsafe.Add(mBase, uint32(v518)+28))
	v771 = v769 << (uint(int32(2)) % 32)
	v774 = *(*int32)(unsafe.Add(mBase, uint32(v771)+uint32(_c_F_globfree[4])))
	if v518 != v774 {
		goto L173
	} else {
		goto L174
	}
L172:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v760)+24)) = v715
	v791 = *(*int32)(unsafe.Add(mBase, uint32(v518)+16))
	if v791 == int32(0) {
		goto L180
	} else {
		goto L181
	}
L173:
	;
	v784 = *(*int32)(unsafe.Add(mBase, uint32(v715)+16))
	if v784 != v518 {
		goto L177
	} else {
		goto L178
	}
L174:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v771)+uint32(_c_F_globfree[4]))) = v760
	if v760 != 0 {
		goto L172
	} else {
		goto L175
	}
L175:
	;
	v777 = int32(0)
	v779 = *(*int32)(unsafe.Add(mBase, _c_F_globfree[5]))
	*(*int32)(unsafe.Add(mBase, _c_F_globfree[5])) = v779 & base.I32_rotl(int32(-2), v769)
	goto L147
L176:
	;
	if v760 == int32(0) {
		goto L147
	} else {
		goto L179
	}
L177:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v715)+20)) = v760
	goto L176
L178:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v715)+16)) = v760
	goto L176
L179:
	;
	goto L172
L180:
	;
	v796 = *(*int32)(unsafe.Add(mBase, uint32(v518)+20))
	if v796 == int32(0) {
		goto L147
	} else {
		goto L182
	}
L181:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v760)+16)) = v791
	*(*int32)(unsafe.Add(mBase, uint32(v791)+24)) = v760
	goto L180
L182:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v760)+20)) = v796
	*(*int32)(unsafe.Add(mBase, uint32(v796)+24)) = v760
	goto L147
L183:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_globfree[3])) = v698
	goto L106
L184:
	;
	if base.Ui32(int32(16777215)) < base.Ui32(v818) {
		v865 = int32(31)
		goto L189
	} else {
		goto L190
	}
L185:
	;
	v830 = v818 & int32(-8)
	v832 = v830 + int32(9128464)
	v834 = *(*int32)(unsafe.Add(mBase, _c_F_globfree[2]))
	v838 = int32(1) << (uint(int32(base.Ui32(v818)>>(uint(int32(3))%32))) % 32)
	if v834&v838 != 0 {
		goto L187
	} else {
		goto L188
	}
L186:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v830)+uint32(_c_F_globfree[8]))) = v643
	*(*int32)(unsafe.Add(mBase, uint32(v844)+12)) = v643
	*(*int32)(unsafe.Add(mBase, uint32(v643)+12)) = v832
	*(*int32)(unsafe.Add(mBase, uint32(v643)+8)) = v844
	goto L106
L187:
	;
	v843 = *(*int32)(unsafe.Add(mBase, uint32(v830)+uint32(_c_F_globfree[8])))
	v844 = v843
	goto L186
L188:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_globfree[2])) = v834 | v838
	v844 = v832
	goto L186
L189:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v643)+28)) = v865
	*(*int64)(unsafe.Add(mBase, uint32(v643)+16)) = int64(0)
	v870 = v865 << (uint(int32(2)) % 32)
	v874 = *(*int32)(unsafe.Add(mBase, _c_F_globfree[5]))
	v876 = int32(1) << (uint(v865) % 32)
	if v874&v876 != 0 {
		goto L194
	} else {
		goto L195
	}
L190:
	;
	v855 = base.I32_clz(int32(base.Ui32(v818) >> (uint(int32(8)) % 32)))
	v858 = int32(1)
	v865 = int32(base.Ui32(v818)>>(uint(int32(38)-v855)%32))&v858 - v855<<(uint(v858)%32) + int32(62)
	goto L189
L191:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v643+v937))) = v940
	*(*int32)(unsafe.Add(mBase, uint32(v643)+12)) = v939
	*(*int32)(unsafe.Add(mBase, uint32(v643+v935))) = v938
	v949 = int32(0)
	v951 = *(*int32)(unsafe.Add(mBase, _c_F_globfree[9]))
	v952 = int32(-1)
	v953 = v951 + v952
	if v953 != 0 {
		goto L203
	} else {
		goto L204
	}
L192:
	;
	v929 = *(*int32)(unsafe.Add(mBase, uint32(v899)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v929)+12)) = v643
	*(*int32)(unsafe.Add(mBase, uint32(v899)+8)) = v643
	v935 = int32(24)
	v937 = int32(8)
	v938 = int32(0)
	v939 = v899
	v940 = v929
	goto L191
L193:
	;
	v935 = v920
	v937 = v922
	v938 = v643
	v939 = v643
	v940 = v925
	goto L191
L194:
	;
	if v865 == int32(31) {
		goto L196
	} else {
		goto L197
	}
L195:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_globfree[5])) = v874 | v876
	*(*int32)(unsafe.Add(mBase, uint32(v870)+uint32(_c_F_globfree[4]))) = v643
	v920 = int32(8)
	v922 = int32(24)
	v925 = v870 + int32(9128728)
	goto L193
L196:
	;
	v891 = int32(0)
	goto L198
L197:
	;
	v891 = int32(25) - int32(base.Ui32(v865)>>(uint(int32(1))%32))
	goto L198
L198:
	;
	v893 = *(*int32)(unsafe.Add(mBase, uint32(v870)+uint32(_c_F_globfree[4])))
	v896 = v818 << (uint(v891) % 32)
	v899 = v893
	goto L199
L199:
	;
	v903 = *(*int32)(unsafe.Add(mBase, uint32(v899)+4))
	if v903&int32(-8) == v818 {
		goto L192
	} else {
		goto L201
	}
L200:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v913+int32(16)))) = v643
	v920 = int32(8)
	v922 = int32(24)
	v925 = v899
	goto L193
L201:
	;
	v913 = v899 + int32(base.Ui32(v896)>>(uint(int32(29))%32))&int32(4)
	v914 = *(*int32)(unsafe.Add(mBase, uint32(v913)+16))
	if v914 != 0 {
		v896 = v896 << (uint(int32(1)) % 32)
		v899 = v914
		goto L199
	} else {
		goto L202
	}
L202:
	;
	goto L200
L203:
	;
	v955 = v953
	goto L205
L204:
	;
	v955 = v952
	goto L205
L205:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_globfree[9])) = v955
	goto L107
}
func F_growCI(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
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
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v41 int32
	_ = v41
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v55 int32
	_ = v55
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	if int32(20001) <= v6 {
		F_luaD_throw(m, l0, int32(5))
		mBase = m.M
		v55 = m.ExcPending
		if v55 != 0 {
			return int32(0)
		} else {
			base.Wasm_trap_unreachable()
			for {
			}
		}
	} else {
		v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
		v10 = int32(1)
		v11 = v6 << (uint(v10) % 32)
		if base.Ui32(int32(178956970)) < base.Ui32(v11|v10) {
			v24 = F_luaM_toobig(m, l0)
			mBase = m.M
			v25 = m.ExcPending
			if v25 != 0 {
				return int32(0)
			} else {
				v26 = v24
				*(*int32)(unsafe.Add(mBase, uint32(l0)+48)) = v11
				*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = v26
				v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
				v31 = v26 + (v29 - v9)
				*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v31
				*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = v26 + v11*int32(24) + int32(-24)
				if v6 < int32(10001) {
					v48 = v31
					v50 = v48 + int32(24)
					*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v50
					return v50
				} else {
					v41 = m.G3
					F_luaG_runerror(m, l0, v41+int32(_a_F_growCI_0), int32(0))
					mBase = m.M
					v46 = m.ExcPending
					if v46 != 0 {
						return int32(0)
					} else {
						v47 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
						v48 = v47
						v50 = v48 + int32(24)
						*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v50
						return v50
					}
				}
			}
		} else {
			v20 = F_luaM_realloc_(m, l0, v9, v6*int32(24), v6*int32(48))
			mBase = m.M
			v23 = m.ExcPending
			if v23 != 0 {
				return int32(0)
			} else {
				v26 = v20
				*(*int32)(unsafe.Add(mBase, uint32(l0)+48)) = v11
				*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = v26
				v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
				v31 = v26 + (v29 - v9)
				*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v31
				*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = v26 + v11*int32(24) + int32(-24)
				if v6 < int32(10001) {
					v48 = v31
					v50 = v48 + int32(24)
					*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v50
					return v50
				} else {
					v41 = m.G3
					F_luaG_runerror(m, l0, v41+int32(_a_F_growCI_0), int32(0))
					mBase = m.M
					v46 = m.ExcPending
					if v46 != 0 {
						return int32(0)
					} else {
						v47 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
						v48 = v47
						v50 = v48 + int32(24)
						*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v50
						return v50
					}
				}
			}
		}
	}
}
