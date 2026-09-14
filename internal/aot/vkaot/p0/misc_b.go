package p0

import (
	base "github.com/shibukawa/vkmem/internal/aot/vkaot/base"
	"unsafe"
)

func F___bswap_16_1(m *base.Module, l0 int32) int32 {
	var v2 int32
	_ = v2
	v2 = int32(8)
	return (l0<<(uint(v2)%32) | int32(base.Ui32(l0)>>(uint(v2)%32))) & int32(65535)
}
func F___bswap_16_2(m *base.Module, l0 int32) int32 {
	var v2 int32
	_ = v2
	v2 = int32(8)
	return (l0<<(uint(v2)%32) | int32(base.Ui32(l0)>>(uint(v2)%32))) & int32(65535)
}
func F_b_size(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v62 int32
	_ = v62
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v97 int32
	_ = v97
	var v104 int32
	_ = v104
	var v109 int32
	_ = v109
	v2 = int32(0)
	v9 = m.G0
	v11 = v9 - int32(16)
	m.G0 = v11
	v16 = F_luaL_checklstring(m, l0, int32(1), v2)
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+12)) = int32(1)
	v22 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16))))
	if v22 == int32(0) {
		v97 = v2
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v104 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v104)+8)) = int32(3)
	*(*float64)(unsafe.Add(mBase, uint32(v104))) = base.F64_convert_i32_s(v97)
	v109 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v109 + int32(16)
	goto L26
L4:
	;
	v27 = v2
	v28 = v16
	v29 = v22
	goto L5
L5:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+4)) = v28 + int32(1)
	v36 = int32(0)
	v37 = base.I32_extend8_s(v29)
	v40 = F_optsize(m, l0, v37, v11+int32(4))
	mBase = m.M
	v41 = m.ExcPending
	if v41 != 0 {
		goto L1
	} else {
		goto L7
	}
L6:
	;
	v97 = v92
	goto L3
L7:
	;
	v43 = v29 & int32(255)
	if v43 == int32(99) {
		v56 = v36
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v57 = m.G3
	if v43 != int32(115) {
		goto L16
	} else {
		goto L17
	}
L9:
	;
	if v40 == int32(0) {
		v56 = v36
		goto L8
	} else {
		goto L10
	}
L10:
	;
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v11)+12))
	if base.Ui32(v40) < base.Ui32(v48) {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v50 = v40
	goto L13
L12:
	;
	v50 = v48
	goto L13
L13:
	;
	v52 = v50 + int32(-1)
	v56 = (v50 - v52&v27) & v52
	goto L8
L14:
	;
	goto L22
L15:
	;
	v70 = F_luaL_argerror(m, l0, int32(1), v67)
	mBase = m.M
	v71 = m.ExcPending
	if v71 != 0 {
		goto L1
	} else {
		goto L20
	}
L16:
	;
	v62 = m.G3
	if v43 != int32(99) {
		goto L14
	} else {
		goto L18
	}
L17:
	;
	v67 = v57 + int32(_a_F_b_size_0)
	goto L15
L18:
	;
	if v40 != 0 {
		goto L14
	} else {
		goto L19
	}
L19:
	;
	v67 = v62 + int32(_a_F_b_size_1)
	goto L15
L20:
	;
	goto L14
L21:
	;
	v92 = v40 + v27 + v56
	v93 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
	v94 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v93))))
	if v94 != 0 {
		v27 = v92
		v28 = v93
		v29 = v94
		goto L5
	} else {
		goto L25
	}
L22:
	;
	if base.B2i32(base.Ui32(v37+int32(-48)) < base.Ui32(int32(10)))|base.B2i32(base.Ui32(v37|int32(32)+int32(-97)) < base.Ui32(int32(26))) != 0 {
		goto L21
	} else {
		goto L23
	}
L23:
	;
	F_controloptions(m, l0, v37, v11+int32(4), v11+int32(8))
	mBase = m.M
	v90 = m.ExcPending
	if v90 != 0 {
		goto L1
	} else {
		goto L24
	}
L24:
	;
	goto L21
L25:
	;
	goto L6
L26:
	;
	m.G0 = v11 + int32(16)
	return int32(1)
}
func F_b_unpack(m *base.Module, l0 int32) int32 {
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
	var v26 int32
	_ = v26
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
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
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v96 int32
	_ = v96
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v102 int32
	_ = v102
	var v106 int32
	_ = v106
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v128 int32
	_ = v128
	var v131 int32
	_ = v131
	var v141 int32
	_ = v141
	var v143 int32
	_ = v143
	var v145 int32
	_ = v145
	var v148 int32
	_ = v148
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v159 int32
	_ = v159
	var v164 int32
	_ = v164
	var v169 int32
	_ = v169
	var v172 int32
	_ = v172
	var v186 int32
	_ = v186
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v201 int32
	_ = v201
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v206 int32
	_ = v206
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v220 int32
	_ = v220
	var v225 int32
	_ = v225
	var v228 int32
	_ = v228
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v240 int32
	_ = v240
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v248 int32
	_ = v248
	var v249 int32
	_ = v249
	var v254 int32
	_ = v254
	var v260 int32
	_ = v260
	var v265 int32
	_ = v265
	var v268 int32
	_ = v268
	var v282 int32
	_ = v282
	var v284 int32
	_ = v284
	var v285 int32
	_ = v285
	var v297 int32
	_ = v297
	var v298 int32
	_ = v298
	var v299 int32
	_ = v299
	var v302 int32
	_ = v302
	var v310 int32
	_ = v310
	var v322 int32
	_ = v322
	var v327 int32
	_ = v327
	var v330 int32
	_ = v330
	var v334 float64
	_ = v334
	var v336 int32
	_ = v336
	var v340 int32
	_ = v340
	var v351 int32
	_ = v351
	var v353 int32
	_ = v353
	var v356 int32
	_ = v356
	var v357 int32
	_ = v357
	var v360 int32
	_ = v360
	var v361 int32
	_ = v361
	var v365 float32
	_ = v365
	var v368 int32
	_ = v368
	var v372 int32
	_ = v372
	var v383 int32
	_ = v383
	var v385 int32
	_ = v385
	var v388 int32
	_ = v388
	var v389 int32
	_ = v389
	var v392 int32
	_ = v392
	var v393 int32
	_ = v393
	var v396 int32
	_ = v396
	var v397 int32
	_ = v397
	var v400 int32
	_ = v400
	var v401 int32
	_ = v401
	var v405 float64
	_ = v405
	var v407 int32
	_ = v407
	var v411 int32
	_ = v411
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
	var v429 float64
	_ = v429
	var v430 int32
	_ = v430
	var v449 int32
	_ = v449
	var v466 int32
	_ = v466
	var v468 int32
	_ = v468
	var v470 int32
	_ = v470
	var v471 int32
	_ = v471
	var v476 int32
	_ = v476
	var v479 int32
	_ = v479
	var v480 int32
	_ = v480
	var v482 int32
	_ = v482
	var v483 int32
	_ = v483
	var v487 int32
	_ = v487
	var v490 int32
	_ = v490
	var v491 int32
	_ = v491
	var v492 int32
	_ = v492
	var v493 int32
	_ = v493
	var v497 int32
	_ = v497
	var v506 int32
	_ = v506
	var v508 int32
	_ = v508
	var v511 int32
	_ = v511
	var v514 int32
	_ = v514
	var v515 int32
	_ = v515
	var v516 int32
	_ = v516
	var v518 int32
	_ = v518
	var v523 int32
	_ = v523
	var v525 int32
	_ = v525
	var v526 int32
	_ = v526
	var v530 int32
	_ = v530
	var v540 int32
	_ = v540
	var v542 int32
	_ = v542
	var v546 int32
	_ = v546
	var v547 int32
	_ = v547
	var v550 int32
	_ = v550
	var v555 int32
	_ = v555
	var v557 int32
	_ = v557
	var v560 int32
	_ = v560
	var v562 int32
	_ = v562
	var v567 int32
	_ = v567
	var v569 int32
	_ = v569
	var v574 int32
	_ = v574
	var v576 int32
	_ = v576
	var v579 int32
	_ = v579
	var v584 int32
	_ = v584
	var v596 int32
	_ = v596
	var v597 int32
	_ = v597
	var v601 int32
	_ = v601
	var v602 int32
	_ = v602
	var v603 int32
	_ = v603
	var v605 int32
	_ = v605
	var v606 int32
	_ = v606
	var v615 int32
	_ = v615
	var v623 int32
	_ = v623
	var v624 int32
	_ = v624
	var v631 int32
	_ = v631
	var v632 int32
	_ = v632
	var v633 int32
	_ = v633
	var v634 int32
	_ = v634
	var v640 int32
	_ = v640
	var v643 int32
	_ = v643
	var v654 int32
	_ = v654
	var v659 int32
	_ = v659
	v16 = m.G0
	v18 = v16 - int32(32)
	m.G0 = v18
	v20 = int32(1)
	v23 = F_luaL_checklstring(m, l0, v20, int32(0))
	mBase = m.M
	v26 = m.ExcPending
	if v26 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v30 = F_luaL_checklstring(m, l0, int32(2), v18+int32(16))
	mBase = m.M
	v31 = m.ExcPending
	if v31 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v34 = F_luaL_optinteger(m, l0, int32(3), int32(1))
	mBase = m.M
	v35 = m.ExcPending
	if v35 != 0 {
		goto L1
	} else {
		goto L5
	}
L4:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v18)+24)) = int64(4294967297)
	v44 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23))))
	if v44 == int32(0) {
		v640 = v20
		v643 = v34
		goto L8
	} else {
		goto L9
	}
L5:
	;
	if v34 != 0 {
		goto L4
	} else {
		goto L6
	}
L6:
	;
	v37 = m.G3
	v40 = F_luaL_argerror(m, l0, int32(3), v37+int32(_a_F_b_unpack_0))
	mBase = m.M
	v41 = m.ExcPending
	if v41 != 0 {
		goto L1
	} else {
		goto L7
	}
L7:
	;
	goto L4
L8:
	;
	v654 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v654)+8)) = int32(3)
	*(*float64)(unsafe.Add(mBase, uint32(v654))) = base.F64_convert_i32_s(v643)
	v659 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v659 + int32(16)
	goto L132
L9:
	;
	v53 = v23
	v55 = v34 + int32(-1)
	v56 = v44
	v57 = int32(0)
	goto L10
L10:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+20)) = v53 + int32(1)
	v68 = base.I32_extend8_s(v56)
	v71 = F_optsize(m, l0, v68, v18+int32(20))
	mBase = m.M
	v72 = m.ExcPending
	if v72 != 0 {
		goto L1
	} else {
		goto L12
	}
L11:
	;
	v634 = int32(1)
	v640 = v623 + v634
	v643 = v631 + v634
	goto L8
L12:
	;
	v75 = v56 & int32(255)
	if v75 == int32(99) {
		v89 = int32(0)
		goto L13
	} else {
		goto L14
	}
L13:
	;
	v90 = v89 + v55
	v91 = *(*int32)(unsafe.Add(mBase, uint32(v18)+16))
	if base.Ui32(v91) < base.Ui32(v71) {
		goto L20
	} else {
		goto L21
	}
L14:
	;
	v78 = int32(0)
	if v71 == v78 {
		v89 = v78
		goto L13
	} else {
		goto L15
	}
L15:
	;
	v81 = *(*int32)(unsafe.Add(mBase, uint32(v18)+28))
	if base.Ui32(v71) < base.Ui32(v81) {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	v83 = v71
	goto L18
L17:
	;
	v83 = v81
	goto L18
L18:
	;
	v85 = v83 + int32(-1)
	v89 = (v83 - v85&v55) & v85
	goto L13
L19:
	;
	v102 = m.G3
	F_luaL_checkstack(m, l0, int32(2), v102+int32(_a_F_b_unpack_1))
	mBase = m.M
	v106 = m.ExcPending
	if v106 != 0 {
		goto L1
	} else {
		goto L24
	}
L20:
	;
	v96 = m.G3
	v99 = F_luaL_argerror(m, l0, int32(2), v96+int32(_a_F_b_unpack_2))
	mBase = m.M
	v100 = m.ExcPending
	if v100 != 0 {
		goto L1
	} else {
		goto L23
	}
L21:
	;
	if base.Ui32(v90) <= base.Ui32(v91-v71) {
		goto L19
	} else {
		goto L22
	}
L22:
	;
	goto L20
L23:
	;
	goto L19
L24:
	;
	switch v75 + int32(-66) {
	case 0, 6, 7, 10, 18, 32, 38, 39, 42:
		goto L31
	default:
		goto L26
	case 33:
		goto L28
	case 34:
		goto L29
	case 36:
		goto L30
	case 49:
		goto L27
	case 54:
		v623 = v57
		v624 = v71
		goto L25
	}
L25:
	;
	v631 = v624 + v90
	v632 = *(*int32)(unsafe.Add(mBase, uint32(v18)+20))
	v633 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v632))))
	if v633 != 0 {
		v53 = v632
		v55 = v631
		v56 = v633
		v57 = v623
		goto L10
	} else {
		goto L131
	}
L26:
	;
	F_controloptions(m, l0, v68, v18+int32(20), v18+int32(24))
	mBase = m.M
	v615 = m.ExcPending
	if v615 != 0 {
		goto L1
	} else {
		goto L130
	}
L27:
	;
	v490 = v30 + v90
	v491 = int32(0)
	v492 = *(*int32)(unsafe.Add(mBase, uint32(v18)+16))
	v493 = v492 - v90
	v497 = base.B2i32(v493 != v491)
	if v490&int32(3) == v491 {
		v523 = v490
		v525 = v493
		v526 = v497
		goto L105
	} else {
		goto L106
	}
L28:
	;
	if v71 != 0 {
		v482 = v57
		v483 = v71
		goto L76
	} else {
		goto L77
	}
L29:
	;
	if v71 == int32(0) {
		goto L71
	} else {
		goto L72
	}
L30:
	;
	if v71 == int32(0) {
		goto L65
	} else {
		goto L66
	}
L31:
	;
	v109 = v30 + v90
	v110 = *(*int32)(unsafe.Add(mBase, uint32(v18)+24))
	if v110 == int32(0) {
		goto L33
	} else {
		goto L34
	}
L32:
	;
	if base.Ui32(int32(96)) < base.Ui32(v75) {
		goto L58
	} else {
		goto L59
	}
L33:
	;
	if int32(1) <= v71 {
		goto L46
	} else {
		goto L47
	}
L34:
	;
	if int32(1) <= v71 {
		goto L35
	} else {
		goto L36
	}
L35:
	;
	v117 = v71 & int32(3)
	v118 = int32(0)
	if base.Ui32(v71) < base.Ui32(int32(4)) {
		v169 = v71
		v172 = v118
		goto L37
	} else {
		goto L38
	}
L36:
	;
	v310 = int32(0)
	goto L32
L37:
	;
	if v117 == int32(0) {
		v310 = v172
		goto L32
	} else {
		goto L42
	}
L38:
	;
	v128 = v71
	v131 = int32(0)
	goto L39
L39:
	;
	v141 = v128 + int32(-4)
	v143 = v131 + int32(4)
	if v143 != v71&int32(2147483644) {
		v128 = v141
		v131 = v143
		goto L39
	} else {
		goto L41
	}
L40:
	;
	v145 = v109 + v128
	v148 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v145+int32(-1)))))
	v153 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v145+int32(-2)))))
	v154 = int32(8)
	v159 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v145+int32(-3)))))
	v164 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v109+v141))))
	v169 = v141
	v172 = (v148<<(uint(int32(16))%32)|v153<<(uint(v154)%32)|v159)<<(uint(v154)%32) | v164
	goto L37
L41:
	;
	goto L40
L42:
	;
	v186 = v169
	v188 = v118
	v189 = v172
	goto L43
L43:
	;
	v201 = v186 + int32(-1)
	v203 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v109+v201))))
	v204 = v189<<(uint(int32(8))%32) | v203
	v206 = v188 + int32(1)
	if v206 != v117 {
		v186 = v201
		v188 = v206
		v189 = v204
		goto L43
	} else {
		goto L45
	}
L45:
	;
	v310 = v204
	goto L32
L46:
	;
	v212 = v71 & int32(3)
	v213 = int32(0)
	if base.Ui32(v71) < base.Ui32(int32(4)) {
		v265 = v213
		v268 = v213
		goto L48
	} else {
		goto L49
	}
L47:
	;
	v310 = int32(0)
	goto L32
L48:
	;
	if v212 == int32(0) {
		v310 = v268
		goto L32
	} else {
		goto L53
	}
L49:
	;
	v220 = int32(0)
	v225 = v220
	v228 = v220
	goto L50
L50:
	;
	v237 = int32(4)
	v238 = v225 + v237
	v240 = v228 + v237
	if v240 != v71&int32(2147483644) {
		v225 = v238
		v228 = v240
		goto L50
	} else {
		goto L52
	}
L51:
	;
	v242 = v109 + v225
	v243 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v242))))
	v248 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v242+int32(1)))))
	v249 = int32(8)
	v254 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v242+int32(2)))))
	v260 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v242+int32(3)))))
	v265 = v238
	v268 = (v243<<(uint(int32(16))%32)|v248<<(uint(v249)%32)|v254)<<(uint(v249)%32) | v260
	goto L48
L52:
	;
	goto L51
L53:
	;
	v282 = v265
	v284 = v213
	v285 = v268
	goto L54
L54:
	;
	v297 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v109+v282))))
	v298 = v285<<(uint(int32(8))%32) | v297
	v299 = int32(1)
	v302 = v284 + v299
	if v302 != v212 {
		v282 = v282 + v299
		v284 = v302
		v285 = v298
		goto L54
	} else {
		goto L56
	}
L55:
	;
	v310 = v298
	goto L32
L56:
	;
	goto L55
L57:
	;
	v336 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v336)+8)) = int32(3)
	*(*float64)(unsafe.Add(mBase, uint32(v336))) = v334
	v340 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v340 + int32(16)
	goto L63
L58:
	;
	v322 = int32(-1)
	v327 = v322 << (uint(v71<<(uint(int32(3))%32)+v322) % 32)
	if v310&v327 != 0 {
		goto L60
	} else {
		goto L61
	}
L59:
	;
	v334 = base.F64_convert_i32_u(v310)
	goto L57
L60:
	;
	v330 = v327
	goto L62
L61:
	;
	v330 = int32(0)
	goto L62
L62:
	;
	v334 = base.F64_convert_i32_s(v330 | v310)
	goto L57
L63:
	;
	v623 = v57 + int32(1)
	v624 = v71
	goto L25
L64:
	;
	v353 = *(*int32)(unsafe.Add(mBase, uint32(v18)+24))
	if v353 == int32(1) {
		goto L67
	} else {
		goto L68
	}
L65:
	;
	goto L64
L66:
	;
	v351 = F__emscripten_memcpy_bulkmem(m, v18+int32(8), v30+v90, v71)
	mBase = m.M
	goto L65
L67:
	;
	v365 = *(*float32)(unsafe.Add(mBase, uint32(v18)+8))
	v368 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v368)+8)) = int32(3)
	*(*float64)(unsafe.Add(mBase, uint32(v368))) = base.F64_promote_f32(v365)
	v372 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v372 + int32(16)
	goto L69
L68:
	;
	v356 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18)+11)))
	v357 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18)+8)))
	*(*uint8)(unsafe.Add(mBase, uint32(v18)+11)) = uint8(v357)
	*(*uint8)(unsafe.Add(mBase, uint32(v18)+8)) = uint8(v356)
	v360 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18)+9)))
	v361 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18)+10)))
	*(*uint8)(unsafe.Add(mBase, uint32(v18)+9)) = uint8(v361)
	*(*uint8)(unsafe.Add(mBase, uint32(v18)+10)) = uint8(v360)
	goto L67
L69:
	;
	v623 = v57 + int32(1)
	v624 = v71
	goto L25
L70:
	;
	v385 = *(*int32)(unsafe.Add(mBase, uint32(v18)+24))
	if v385 == int32(1) {
		goto L73
	} else {
		goto L74
	}
L71:
	;
	goto L70
L72:
	;
	v383 = F__emscripten_memcpy_bulkmem(m, v18+int32(8), v30+v90, v71)
	mBase = m.M
	goto L71
L73:
	;
	v405 = *(*float64)(unsafe.Add(mBase, uint32(v18)+8))
	v407 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v407)+8)) = int32(3)
	*(*float64)(unsafe.Add(mBase, uint32(v407))) = v405
	v411 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v411 + int32(16)
	goto L75
L74:
	;
	v388 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18)+15)))
	v389 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18)+8)))
	*(*uint8)(unsafe.Add(mBase, uint32(v18)+15)) = uint8(v389)
	*(*uint8)(unsafe.Add(mBase, uint32(v18)+8)) = uint8(v388)
	v392 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18)+9)))
	v393 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18)+14)))
	*(*uint8)(unsafe.Add(mBase, uint32(v18)+9)) = uint8(v393)
	*(*uint8)(unsafe.Add(mBase, uint32(v18)+14)) = uint8(v392)
	v396 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18)+10)))
	v397 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18)+13)))
	*(*uint8)(unsafe.Add(mBase, uint32(v18)+10)) = uint8(v397)
	*(*uint8)(unsafe.Add(mBase, uint32(v18)+13)) = uint8(v396)
	v400 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18)+11)))
	v401 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18)+12)))
	*(*uint8)(unsafe.Add(mBase, uint32(v18)+11)) = uint8(v401)
	*(*uint8)(unsafe.Add(mBase, uint32(v18)+12)) = uint8(v400)
	goto L73
L75:
	;
	v623 = v57 + int32(1)
	v624 = v71
	goto L25
L76:
	;
	F_lua_pushlstring(m, l0, v30+v90, v483)
	mBase = m.M
	v487 = m.ExcPending
	if v487 != 0 {
		goto L1
	} else {
		goto L100
	}
L77:
	;
	if v57 == int32(0) {
		goto L79
	} else {
		goto L80
	}
L78:
	;
	v429 = F_lua_tonumber(m, l0, int32(-1))
	mBase = m.M
	v430 = m.ExcPending
	if v430 != 0 {
		goto L1
	} else {
		goto L84
	}
L79:
	;
	v422 = m.G3
	v426 = F_luaL_error(m, l0, v422+int32(_a_F_b_unpack_3), int32(0))
	mBase = m.M
	v427 = m.ExcPending
	if v427 != 0 {
		goto L1
	} else {
		goto L83
	}
L80:
	;
	v420 = F_lua_isnumber(m, l0, int32(-1))
	mBase = m.M
	v421 = m.ExcPending
	if v421 != 0 {
		goto L1
	} else {
		goto L81
	}
L81:
	;
	if v420 != 0 {
		goto L78
	} else {
		goto L82
	}
L82:
	;
	goto L79
L83:
	;
	goto L78
L84:
	;
	goto L87
L85:
	;
	if base.F64_lt(v429, float64(4.294967296e+09))&base.F64_ge(v429, float64(0)) == int32(0) {
		goto L94
	} else {
		goto L95
	}
L86:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v449 + int32(-16)
	goto L85
L87:
	;
	v449 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	goto L86
L93:
	;
	v470 = v57 + int32(-1)
	v471 = *(*int32)(unsafe.Add(mBase, uint32(v18)+16))
	if base.Ui32(v471) < base.Ui32(v468) {
		goto L96
	} else {
		goto L97
	}
L94:
	;
	v468 = int32(0)
	goto L93
L95:
	;
	v466 = base.I32_trunc_f64_u(v429)
	v468 = v466
	goto L93
L96:
	;
	v476 = m.G3
	v479 = F_luaL_argerror(m, l0, int32(2), v476+int32(_a_F_b_unpack_2))
	mBase = m.M
	v480 = m.ExcPending
	if v480 != 0 {
		goto L1
	} else {
		goto L99
	}
L97:
	;
	if base.Ui32(v90) <= base.Ui32(v471-v468) {
		v482 = v470
		v483 = v468
		goto L76
	} else {
		goto L98
	}
L98:
	;
	goto L96
L99:
	;
	v482 = v470
	v483 = v468
	goto L76
L100:
	;
	v623 = v482 + int32(1)
	v624 = v483
	goto L25
L101:
	;
	v603 = v596 - v490
	F_lua_pushlstring(m, l0, v490, v603)
	mBase = m.M
	v605 = m.ExcPending
	if v605 != 0 {
		goto L1
	} else {
		goto L129
	}
L102:
	;
	if v596 != 0 {
		goto L101
	} else {
		goto L127
	}
L103:
	;
	v596 = int32(0)
	goto L102
L104:
	;
	v574 = v567
	v576 = v569
	goto L122
L105:
	;
	if v526 == int32(0) {
		goto L103
	} else {
		goto L113
	}
L106:
	;
	if v493 == int32(0) {
		v523 = v490
		v525 = v493
		v526 = v497
		goto L105
	} else {
		goto L107
	}
L107:
	;
	v506 = v490
	v508 = v493
	goto L108
L108:
	;
	v511 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v506))))
	if v511 == int32(0) {
		v567 = v506
		v569 = v508
		goto L104
	} else {
		goto L110
	}
L109:
	;
	v523 = v518
	v525 = v514
	v526 = v516
	goto L105
L110:
	;
	v514 = v508 + int32(-1)
	v515 = int32(0)
	v516 = base.B2i32(v514 != v515)
	v518 = v506 + int32(1)
	if v518&int32(3) == v515 {
		v523 = v518
		v525 = v514
		v526 = v516
		goto L105
	} else {
		goto L111
	}
L111:
	;
	if v514 != 0 {
		v506 = v518
		v508 = v514
		goto L108
	} else {
		goto L112
	}
L112:
	;
	goto L109
L113:
	;
	v530 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v523))))
	if v530 == int32(0) {
		v560 = v523
		v562 = v525
		goto L114
	} else {
		goto L115
	}
L114:
	;
	if v562 == int32(0) {
		goto L103
	} else {
		goto L121
	}
L115:
	;
	if base.Ui32(v525) < base.Ui32(int32(4)) {
		v560 = v523
		v562 = v525
		goto L114
	} else {
		goto L116
	}
L116:
	;
	v540 = v523
	v542 = v525
	goto L117
L117:
	;
	v546 = *(*int32)(unsafe.Add(mBase, uint32(v540)))
	v547 = v546 ^ int32(0)
	v550 = int32(-2139062144)
	if (int32(16843008)-v547|v547)&v550 != v550 {
		v567 = v540
		v569 = v542
		goto L104
	} else {
		goto L119
	}
L118:
	;
	v560 = v555
	v562 = v557
	goto L114
L119:
	;
	v555 = v540 + int32(4)
	v557 = v542 + int32(-4)
	if base.Ui32(int32(3)) < base.Ui32(v557) {
		v540 = v555
		v542 = v557
		goto L117
	} else {
		goto L120
	}
L120:
	;
	goto L118
L121:
	;
	v567 = v560
	v569 = v562
	goto L104
L122:
	;
	v579 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v574))))
	if v579 != int32(0) {
		goto L124
	} else {
		goto L125
	}
L123:
	;
	goto L103
L124:
	;
	v584 = v576 + int32(-1)
	if v584 != 0 {
		v574 = v574 + int32(1)
		v576 = v584
		goto L122
	} else {
		goto L126
	}
L125:
	;
	v596 = v574
	goto L102
L126:
	;
	goto L123
L127:
	;
	v597 = m.G3
	v601 = F_luaL_error(m, l0, v597+int32(_a_F_b_unpack_4), int32(0))
	mBase = m.M
	v602 = m.ExcPending
	if v602 != 0 {
		goto L1
	} else {
		goto L128
	}
L128:
	;
	goto L101
L129:
	;
	v606 = int32(1)
	v623 = v57 + v606
	v624 = v603 + v606
	goto L25
L130:
	;
	v623 = v57
	v624 = v71
	goto L25
L131:
	;
	goto L11
L132:
	;
	m.G0 = v18 + int32(32)
	return v640
}
func F_backgroundSlotMigrationDoneHandler(m *base.Module, l0 int32, l1 int32) {
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
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v30 int32
	_ = v30
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v86 int32
	_ = v86
	var v92 int32
	_ = v92
	var v96 int32
	_ = v96
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v107 int32
	_ = v107
	var v109 int32
	_ = v109
	var v112 int32
	_ = v112
	var v115 int32
	_ = v115
	var v120 int32
	_ = v120
	var v124 int32
	_ = v124
	var v126 int32
	_ = v126
	var v128 int32
	_ = v128
	var v131 int32
	_ = v131
	var v137 int32
	_ = v137
	v6 = m.G0
	v8 = v6 - int32(32)
	m.G0 = v8
	v10 = l1 | l0
	if v10 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	v42 = *(*int32)(unsafe.Add(mBase, _c_F_backgroundSlotMigrationDoneHandler[0]))
	if v42 == int32(-1) {
		goto L14
	} else {
		goto L15
	}
L2:
	;
	v21 = *(*int32)(unsafe.Add(mBase, _c_F_backgroundSlotMigrationDoneHandler[1]))
	if l0 == int32(0) {
		goto L7
	} else {
		goto L8
	}
L3:
	;
	v12 = *(*int32)(unsafe.Add(mBase, _c_F_backgroundSlotMigrationDoneHandler[1]))
	if int32(2) < v12 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	F__serverLog(m, int32(2), int32(_a_F_backgroundSlotMigrationDoneHandler_0), int32(0))
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	return
L6:
	;
	goto L1
L7:
	;
	if int32(3) < v21 {
		goto L1
	} else {
		goto L12
	}
L8:
	;
	if l1 != 0 {
		goto L7
	} else {
		goto L9
	}
L9:
	;
	if int32(3) < v21 {
		goto L1
	} else {
		goto L10
	}
L10:
	;
	F__serverLog(m, int32(3), int32(_a_F_backgroundSlotMigrationDoneHandler_1), int32(0))
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		goto L5
	} else {
		goto L11
	}
L11:
	;
	goto L1
L12:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = l1
	F__serverLog(m, int32(3), int32(_a_F_backgroundSlotMigrationDoneHandler_2), v8+int32(16))
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L5
	} else {
		goto L13
	}
L13:
	;
	goto L1
L14:
	;
	v47 = *(*int32)(unsafe.Add(mBase, _c_F_backgroundSlotMigrationDoneHandler[2]))
	if v47 < int32(1) {
		goto L16
	} else {
		goto L17
	}
L15:
	;
	v45 = F_close(m, v42)
	mBase = m.M
	goto L14
L16:
	;
	v58 = int32(_a_F_backgroundSlotMigrationDoneHandler_3)
	*(*int32)(unsafe.Add(mBase, _c_F_backgroundSlotMigrationDoneHandler[3])) = int32(0)
	*(*int64)(unsafe.Add(mBase, _c_F_backgroundSlotMigrationDoneHandler[2])) = int64(-1)
	v65 = *(*int32)(unsafe.Add(mBase, _c_F_backgroundSlotMigrationDoneHandler[4]))
	F_valkey_free(m, v65)
	mBase = m.M
	v67 = m.ExcPending
	if v67 != 0 {
		goto L5
	} else {
		goto L19
	}
L17:
	;
	v51 = *(*int32)(unsafe.Add(mBase, _c_F_backgroundSlotMigrationDoneHandler[5]))
	F_aeDeleteFileEvent(m, v51, v47, int32(1))
	mBase = m.M
	v54 = m.ExcPending
	if v54 != 0 {
		goto L5
	} else {
		goto L18
	}
L18:
	;
	v56 = *(*int32)(unsafe.Add(mBase, _c_F_backgroundSlotMigrationDoneHandler[2]))
	v57 = F_close(m, v56)
	mBase = m.M
	goto L16
L19:
	;
	v68 = int32(_a_F_backgroundSlotMigrationDoneHandler_3)
	*(*int64)(unsafe.Add(mBase, _c_F_backgroundSlotMigrationDoneHandler[4])) = int64(0)
	v72 = *(*int32)(unsafe.Add(mBase, _c_F_backgroundSlotMigrationDoneHandler[6]))
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v72)+uint32(_c_F_backgroundSlotMigrationDoneHandler[7])))
	v75 = v8 + int32(24)
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v73)))
	*(*int32)(unsafe.Add(mBase, uint32(v75)+4)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v75))) = v76
	goto L20
L20:
	;
	v81 = v8 + int32(24)
	v83 = *(*int32)(unsafe.Add(mBase, uint32(v81)))
	if v83 == int32(0) {
		goto L23
	} else {
		goto L24
	}
L21:
	;
	m.G0 = v8 + int32(32)
	return
L22:
	;
	if v83 == int32(0) {
		goto L21
	} else {
		goto L25
	}
L23:
	;
	goto L22
L24:
	;
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v81)+4))
	v92 = *(*int32)(unsafe.Add(mBase, uint32(v83+base.B2i32(v86 == int32(0))<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v81))) = v92
	goto L23
L25:
	;
	v96 = v83
	goto L26
L26:
	;
	v101 = *(*int32)(unsafe.Add(mBase, uint32(v96)+8))
	v102 = *(*int32)(unsafe.Add(mBase, uint32(v101)))
	if v102 != 0 {
		goto L28
	} else {
		goto L29
	}
L27:
	;
	goto L21
L28:
	;
	v126 = v8 + int32(24)
	v128 = *(*int32)(unsafe.Add(mBase, uint32(v126)))
	if v128 == int32(0) {
		goto L39
	} else {
		goto L40
	}
L29:
	;
	v103 = *(*int32)(unsafe.Add(mBase, uint32(v101)+156))
	if v103 != int32(13) {
		goto L28
	} else {
		goto L30
	}
L30:
	;
	if v10 != 0 {
		goto L31
	} else {
		goto L32
	}
L31:
	;
	v112 = *(*int32)(unsafe.Add(mBase, _c_F_backgroundSlotMigrationDoneHandler[1]))
	if int32(3) < v112 {
		goto L34
	} else {
		goto L35
	}
L32:
	;
	F_slotExportBeginStreaming(m, v101)
	mBase = m.M
	v107 = m.ExcPending
	if v107 != 0 {
		goto L5
	} else {
		goto L33
	}
L33:
	;
	v109 = *(*int32)(unsafe.Add(mBase, _c_F_backgroundSlotMigrationDoneHandler[8]))
	*(*int32)(unsafe.Add(mBase, uint32(v101)+192)) = v109
	goto L28
L34:
	;
	F_finishSlotMigrationJob(m, v101, int32(18), int32(_a_F_backgroundSlotMigrationDoneHandler_4))
	mBase = m.M
	v124 = m.ExcPending
	if v124 != 0 {
		goto L5
	} else {
		goto L37
	}
L35:
	;
	v115 = *(*int32)(unsafe.Add(mBase, uint32(v101)+188))
	*(*int32)(unsafe.Add(mBase, uint32(v8))) = v115
	F__serverLog(m, int32(3), int32(_a_F_backgroundSlotMigrationDoneHandler_5), v8)
	mBase = m.M
	v120 = m.ExcPending
	if v120 != 0 {
		goto L5
	} else {
		goto L36
	}
L36:
	;
	goto L34
L37:
	;
	goto L28
L38:
	;
	if v128 != 0 {
		v96 = v128
		goto L26
	} else {
		goto L41
	}
L39:
	;
	goto L38
L40:
	;
	v131 = *(*int32)(unsafe.Add(mBase, uint32(v126)+4))
	v137 = *(*int32)(unsafe.Add(mBase, uint32(v128+base.B2i32(v131 == int32(0))<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v126))) = v137
	goto L39
L41:
	;
	goto L27
}
func F_backupAndUpdateClientArgv(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v29 int32
	_ = v29
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v63 int32
	_ = v63
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
	var v96 int32
	_ = v96
	var v100 int32
	_ = v100
	var v102 int32
	_ = v102
	var v113 int32
	_ = v113
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+192))
	if v11 != 0 {
		v14 = v11
		goto L1
	} else {
		goto L2
	}
L1:
	;
	if l2 == int32(0) {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+192)) = v10
	*(*int32)(unsafe.Add(mBase, uint32(l0)+188)) = v9
	v14 = v10
	goto L1
L3:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = l1
	if v71 == v10 {
		goto L22
	} else {
		goto L23
	}
L4:
	;
	if v14 == v10 {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = l2
	v71 = l2
	goto L3
L6:
	;
	v23 = l1 << (uint(int32(2)) % 32)
	v24 = F_valkey_malloc(m, v23)
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		goto L9
	} else {
		goto L10
	}
L7:
	;
	if v9 < l1 {
		goto L6
	} else {
		goto L8
	}
L8:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = l1
	return
L9:
	;
	return
L10:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v24
	if v9 < l1 {
		goto L12
	} else {
		goto L13
	}
L11:
	;
	if l1 <= v9 {
		v71 = v55
		goto L3
	} else {
		goto L20
	}
L12:
	;
	v29 = v9
	goto L14
L13:
	;
	v29 = l1
	goto L14
L14:
	;
	if v29 <= int32(0) {
		v55 = v24
		goto L11
	} else {
		goto L15
	}
L15:
	;
	v37 = int32(0)
	goto L16
L16:
	;
	v40 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v42 = v37 << (uint(int32(2)) % 32)
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v10+v42)))
	*(*int32)(unsafe.Add(mBase, uint32(v40+v42))) = v45
	F_incrRefCount(m, v45)
	mBase = m.M
	v48 = m.ExcPending
	if v48 != 0 {
		goto L9
	} else {
		goto L18
	}
L17:
	;
	v52 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v55 = v52
	goto L11
L18:
	;
	v50 = v37 + int32(1)
	if v50 != v29 {
		v37 = v50
		goto L16
	} else {
		goto L19
	}
L19:
	;
	goto L17
L20:
	;
	v63 = v9 << (uint(int32(2)) % 32)
	v68 = F__emscripten_memset_bulkmem(m, v55+v63, base.I32_extend8_s(int32(0)), v23-v63)
	mBase = m.M
	goto L21
L21:
	;
	v71 = v55
	goto L3
L22:
	;
	return
L23:
	;
	v80 = *(*int32)(unsafe.Add(mBase, uint32(l0)+192))
	if v80 == v10 {
		goto L22
	} else {
		goto L24
	}
L24:
	;
	v82 = int32(0)
	if v9 <= v82 {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	F_valkey_free(m, v10)
	mBase = m.M
	v113 = m.ExcPending
	if v113 != 0 {
		goto L9
	} else {
		goto L33
	}
L26:
	;
	v85 = v82
	goto L27
L27:
	;
	v96 = *(*int32)(unsafe.Add(mBase, uint32(v10+v85<<(uint(int32(2))%32))))
	if v96 == int32(0) {
		goto L29
	} else {
		goto L30
	}
L28:
	;
	goto L25
L29:
	;
	v102 = v85 + int32(1)
	if v102 != v9 {
		v85 = v102
		goto L27
	} else {
		goto L32
	}
L30:
	;
	F_decrRefCount(m, v96)
	mBase = m.M
	v100 = m.ExcPending
	if v100 != 0 {
		goto L9
	} else {
		goto L31
	}
L31:
	;
	goto L29
L32:
	;
	goto L28
L33:
	;
	goto L22
}
func F_bgsaveCommand(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v15 int32
	_ = v15
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
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v31 int32
	_ = v31
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
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
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
	var v109 int32
	_ = v109
	var v116 int32
	_ = v116
	var v118 int32
	_ = v118
	var v122 int32
	_ = v122
	var v126 int32
	_ = v126
	var v131 int32
	_ = v131
	var v133 int32
	_ = v133
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v139 int32
	_ = v139
	var v141 int32
	_ = v141
	var v145 int32
	_ = v145
	var v152 int32
	_ = v152
	var v158 int32
	_ = v158
	var v161 int32
	_ = v161
	var v163 int32
	_ = v163
	var v165 int32
	_ = v165
	var v168 int32
	_ = v168
	var v169 int64
	_ = v169
	var v174 int64
	_ = v174
	var v179 int64
	_ = v179
	var v184 int64
	_ = v184
	var v189 int64
	_ = v189
	var v194 int64
	_ = v194
	var v199 int64
	_ = v199
	var v202 int64
	_ = v202
	var v205 int32
	_ = v205
	var v207 int32
	_ = v207
	var v212 int32
	_ = v212
	var v215 int32
	_ = v215
	var v217 int32
	_ = v217
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v223 int32
	_ = v223
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v231 int32
	_ = v231
	var v233 int32
	_ = v233
	var v238 int32
	_ = v238
	var v240 int32
	_ = v240
	var v244 int32
	_ = v244
	var v252 int32
	_ = v252
	var v256 int32
	_ = v256
	var v264 int32
	_ = v264
	var v267 int32
	_ = v267
	var v270 int32
	_ = v270
	var v273 int32
	_ = v273
	var v274 int32
	_ = v274
	var v276 int32
	_ = v276
	var v278 int32
	_ = v278
	var v279 int32
	_ = v279
	var v282 int32
	_ = v282
	var v284 int32
	_ = v284
	var v286 int32
	_ = v286
	v6 = m.G0
	v8 = v6 - int32(80)
	m.G0 = v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v10 < int32(2) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v8 + int32(80)
	return
L2:
	;
	v168 = int32(0)
	v169 = *(*int64)(unsafe.Add(mBase, _c_F_bgsaveCommand[0]))
	*(*int64)(unsafe.Add(mBase, uint32(v8+int32(72)))) = v169
	v174 = *(*int64)(unsafe.Add(mBase, _c_F_bgsaveCommand[1]))
	*(*int64)(unsafe.Add(mBase, uint32(v8+int32(64)))) = v174
	v179 = *(*int64)(unsafe.Add(mBase, _c_F_bgsaveCommand[2]))
	*(*int64)(unsafe.Add(mBase, uint32(v8+int32(56)))) = v179
	v184 = *(*int64)(unsafe.Add(mBase, _c_F_bgsaveCommand[3]))
	*(*int64)(unsafe.Add(mBase, uint32(v8+int32(48)))) = v184
	v189 = *(*int64)(unsafe.Add(mBase, _c_F_bgsaveCommand[4]))
	*(*int64)(unsafe.Add(mBase, uint32(v8+int32(40)))) = v189
	v194 = *(*int64)(unsafe.Add(mBase, _c_F_bgsaveCommand[5]))
	*(*int64)(unsafe.Add(mBase, uint32(v8+int32(32)))) = v194
	v199 = *(*int64)(unsafe.Add(mBase, _c_F_bgsaveCommand[6]))
	*(*int64)(unsafe.Add(mBase, uint32(v8+int32(24)))) = v199
	v202 = *(*int64)(unsafe.Add(mBase, _c_F_bgsaveCommand[7]))
	*(*int64)(unsafe.Add(mBase, uint32(v8)+16)) = v202
	v205 = *(*int32)(unsafe.Add(mBase, _c_F_bgsaveCommand[8]))
	if v205 != 0 {
		goto L55
	} else {
		goto L56
	}
L3:
	;
	if v10 != int32(2) {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v163 = *(*int32)(unsafe.Add(mBase, _c_F_bgsaveCommand[9]))
	F_addReplyErrorObject(m, l0, v163)
	mBase = m.M
	v165 = m.ExcPending
	if v165 != 0 {
		goto L38
	} else {
		goto L52
	}
L5:
	;
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v16 = *(*int32)(unsafe.Add(mBase, uint32(v15)+4))
	v17 = F_objectGetVal(m, v16)
	mBase = m.M
	v18 = int32(_a_F_bgsaveCommand_0)
	v21 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17))))
	if v21 != 0 {
		goto L8
	} else {
		goto L9
	}
L6:
	;
	if v53-v55 == int32(0) {
		goto L2
	} else {
		goto L18
	}
L7:
	;
	v53 = F_tolower(m, v49)
	mBase = m.M
	v54 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v50))))
	v55 = F_tolower(m, v54)
	mBase = m.M
	goto L6
L8:
	;
	v23 = v17
	v24 = v18
	v25 = v21
	goto L11
L9:
	;
	v49 = int32(0)
	v50 = v18
	goto L7
L10:
	;
	v49 = v46 & int32(255)
	v50 = v45
	goto L7
L11:
	;
	v27 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24))))
	if v27 == int32(0) {
		v45 = v24
		v46 = v25
		goto L10
	} else {
		goto L13
	}
L12:
	;
	v45 = v39
	v46 = int32(0)
	goto L10
L13:
	;
	v31 = v25 & int32(255)
	if v31 == v27 {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v38 = int32(1)
	v39 = v24 + v38
	v40 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23)+1)))
	if v40 != 0 {
		v23 = v23 + v38
		v24 = v39
		v25 = v40
		goto L11
	} else {
		goto L17
	}
L15:
	;
	v33 = F_tolower(m, v31)
	mBase = m.M
	v34 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24))))
	v35 = F_tolower(m, v34)
	mBase = m.M
	if v33 == v35 {
		goto L14
	} else {
		goto L16
	}
L16:
	;
	v37 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23))))
	v45 = v24
	v46 = v37
	goto L10
L17:
	;
	goto L12
L18:
	;
	v59 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v59 != int32(2) {
		goto L4
	} else {
		goto L19
	}
L19:
	;
	v62 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v62)+4))
	v64 = F_objectGetVal(m, v63)
	mBase = m.M
	v65 = int32(_a_F_bgsaveCommand_1)
	v68 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v64))))
	if v68 != 0 {
		goto L22
	} else {
		goto L23
	}
L20:
	;
	if v100-v102 != 0 {
		goto L4
	} else {
		goto L32
	}
L21:
	;
	v100 = F_tolower(m, v96)
	mBase = m.M
	v101 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v97))))
	v102 = F_tolower(m, v101)
	mBase = m.M
	goto L20
L22:
	;
	v70 = v64
	v71 = v65
	v72 = v68
	goto L25
L23:
	;
	v96 = int32(0)
	v97 = v65
	goto L21
L24:
	;
	v96 = v93 & int32(255)
	v97 = v92
	goto L21
L25:
	;
	v74 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v71))))
	if v74 == int32(0) {
		v92 = v71
		v93 = v72
		goto L24
	} else {
		goto L27
	}
L26:
	;
	v92 = v86
	v93 = int32(0)
	goto L24
L27:
	;
	v78 = v72 & int32(255)
	if v78 == v74 {
		goto L28
	} else {
		goto L29
	}
L28:
	;
	v85 = int32(1)
	v86 = v71 + v85
	v87 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v70)+1)))
	if v87 != 0 {
		v70 = v70 + v85
		v71 = v86
		v72 = v87
		goto L25
	} else {
		goto L31
	}
L29:
	;
	v80 = F_tolower(m, v78)
	mBase = m.M
	v81 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v71))))
	v82 = F_tolower(m, v81)
	mBase = m.M
	if v80 == v82 {
		goto L28
	} else {
		goto L30
	}
L30:
	;
	v84 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v70))))
	v92 = v71
	v93 = v84
	goto L24
L31:
	;
	goto L26
L32:
	;
	v105 = *(*int32)(unsafe.Add(mBase, _c_F_bgsaveCommand[10]))
	if v105 != int32(1) {
		goto L33
	} else {
		goto L34
	}
L33:
	;
	v141 = *(*int32)(unsafe.Add(mBase, _c_F_bgsaveCommand[11]))
	if v141 != int32(1) {
		goto L45
	} else {
		goto L46
	}
L34:
	;
	v109 = *(*int32)(unsafe.Add(mBase, _c_F_bgsaveCommand[12]))
	if int32(2) < v109 {
		goto L36
	} else {
		goto L37
	}
L35:
	;
	F_addReplyStatus(m, l0, int32(_a_F_bgsaveCommand_2))
	mBase = m.M
	v139 = m.ExcPending
	if v139 != 0 {
		goto L38
	} else {
		goto L44
	}
L36:
	;
	v133 = *(*int32)(unsafe.Add(mBase, _c_F_bgsaveCommand[13]))
	v135 = F_kill(m, v133, int32(10))
	mBase = m.M
	v136 = m.ExcPending
	if v136 != 0 {
		goto L38
	} else {
		goto L43
	}
L37:
	;
	F__serverLog(m, int32(2), int32(_a_F_bgsaveCommand_3), int32(0))
	mBase = m.M
	v116 = m.ExcPending
	if v116 != 0 {
		goto L38
	} else {
		goto L39
	}
L38:
	;
	return
L39:
	;
	v118 = *(*int32)(unsafe.Add(mBase, _c_F_bgsaveCommand[10]))
	if v118 != int32(1) {
		goto L35
	} else {
		goto L40
	}
L40:
	;
	v122 = *(*int32)(unsafe.Add(mBase, _c_F_bgsaveCommand[12]))
	if int32(2) < v122 {
		goto L36
	} else {
		goto L41
	}
L41:
	;
	v126 = *(*int32)(unsafe.Add(mBase, _c_F_bgsaveCommand[13]))
	*(*int32)(unsafe.Add(mBase, uint32(v8))) = v126
	F__serverLog(m, int32(2), int32(_a_F_bgsaveCommand_4), v8)
	mBase = m.M
	v131 = m.ExcPending
	if v131 != 0 {
		goto L38
	} else {
		goto L42
	}
L42:
	;
	goto L36
L43:
	;
	goto L35
L44:
	;
	goto L1
L45:
	;
	F_addReplyError(m, l0, int32(_a_F_bgsaveCommand_5))
	mBase = m.M
	v161 = m.ExcPending
	if v161 != 0 {
		goto L38
	} else {
		goto L51
	}
L46:
	;
	v145 = *(*int32)(unsafe.Add(mBase, _c_F_bgsaveCommand[12]))
	if int32(2) < v145 {
		goto L47
	} else {
		goto L48
	}
L47:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_bgsaveCommand[11])) = int32(0)
	F_addReplyStatus(m, l0, int32(_a_F_bgsaveCommand_6))
	mBase = m.M
	v158 = m.ExcPending
	if v158 != 0 {
		goto L38
	} else {
		goto L50
	}
L48:
	;
	F__serverLog(m, int32(2), int32(_a_F_bgsaveCommand_7), int32(0))
	mBase = m.M
	v152 = m.ExcPending
	if v152 != 0 {
		goto L38
	} else {
		goto L49
	}
L49:
	;
	goto L47
L50:
	;
	goto L1
L51:
	;
	goto L1
L52:
	;
	goto L1
L53:
	;
	v233 = *(*int32)(unsafe.Add(mBase, _c_F_bgsaveCommand[10]))
	if v233 != int32(1) {
		goto L65
	} else {
		goto L66
	}
L54:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = v227
	v231 = v8 + int32(16)
	goto L53
L55:
	;
	v217 = *(*int32)(unsafe.Add(mBase, _c_F_bgsaveCommand[14]))
	if v217 == int32(0) {
		goto L61
	} else {
		goto L62
	}
L56:
	;
	v207 = *(*int32)(unsafe.Add(mBase, _c_F_bgsaveCommand[15]))
	if v207 == int32(0) {
		goto L55
	} else {
		goto L57
	}
L57:
	;
	v212 = *(*int32)(unsafe.Add(mBase, _c_F_bgsaveCommand[16]))
	if v212 == int32(-1) {
		goto L58
	} else {
		goto L59
	}
L58:
	;
	v215 = int32(0)
	goto L60
L59:
	;
	v215 = v212
	goto L60
L60:
	;
	v227 = v215
	goto L54
L61:
	;
	v223 = *(*int32)(unsafe.Add(mBase, _c_F_bgsaveCommand[17]))
	if v223 != 0 {
		goto L63
	} else {
		goto L64
	}
L62:
	;
	v220 = *(*int32)(unsafe.Add(mBase, uint32(v217)+96))
	v221 = *(*int32)(unsafe.Add(mBase, uint32(v220)+28))
	v227 = v221
	goto L54
L63:
	;
	v225 = *(*int32)(unsafe.Add(mBase, uint32(v223)+96))
	v226 = *(*int32)(unsafe.Add(mBase, uint32(v225)+28))
	v227 = v226
	goto L54
L64:
	;
	v231 = int32(0)
	goto L53
L65:
	;
	v240 = *(*int32)(unsafe.Add(mBase, _c_F_bgsaveCommand[13]))
	goto L69
L66:
	;
	F_addReplyError(m, l0, int32(_a_F_bgsaveCommand_8))
	mBase = m.M
	v238 = m.ExcPending
	if v238 != 0 {
		goto L38
	} else {
		goto L67
	}
L67:
	;
	goto L1
L68:
	;
	v274 = int32(0)
	v276 = *(*int32)(unsafe.Add(mBase, _c_F_bgsaveCommand[18]))
	v278 = F_rdbSaveBackground(m, v274, v276, v231, v274)
	mBase = m.M
	v279 = m.ExcPending
	if v279 != 0 {
		goto L38
	} else {
		goto L84
	}
L69:
	;
	v244 = *(*int32)(unsafe.Add(mBase, _c_F_bgsaveCommand[19]))
	if base.B2i32(v240 != int32(-1))|v244 == int32(0) {
		goto L68
	} else {
		goto L70
	}
L70:
	;
	if int32(1) < v10 {
		goto L72
	} else {
		goto L73
	}
L71:
	;
	F_addReplyError(m, l0, int32(_a_F_bgsaveCommand_9))
	mBase = m.M
	v273 = m.ExcPending
	if v273 != 0 {
		goto L38
	} else {
		goto L82
	}
L72:
	;
	v252 = int32(_a_F_bgsaveCommand_10)
	*(*int32)(unsafe.Add(mBase, _c_F_bgsaveCommand[11])) = int32(1)
	v256 = *(*int32)(unsafe.Add(mBase, _c_F_bgsaveCommand[12]))
	if int32(2) < v256 {
		goto L75
	} else {
		goto L76
	}
L73:
	;
	if v244 == int32(0) {
		goto L71
	} else {
		goto L74
	}
L74:
	;
	goto L72
L75:
	;
	F_addReplyStatus(m, l0, int32(_a_F_bgsaveCommand_11))
	mBase = m.M
	v270 = m.ExcPending
	if v270 != 0 {
		goto L38
	} else {
		goto L81
	}
L76:
	;
	if int32(1) < v10 {
		goto L77
	} else {
		goto L78
	}
L77:
	;
	v264 = int32(_a_F_bgsaveCommand_12)
	goto L79
L78:
	;
	v264 = int32(_a_F_bgsaveCommand_13)
	goto L79
L79:
	;
	F__serverLog(m, int32(2), v264, int32(0))
	mBase = m.M
	v267 = m.ExcPending
	if v267 != 0 {
		goto L38
	} else {
		goto L80
	}
L80:
	;
	goto L75
L81:
	;
	goto L1
L82:
	;
	goto L1
L83:
	;
	v284 = *(*int32)(unsafe.Add(mBase, _c_F_bgsaveCommand[20]))
	F_addReplyErrorObject(m, l0, v284)
	mBase = m.M
	v286 = m.ExcPending
	if v286 != 0 {
		goto L38
	} else {
		goto L87
	}
L84:
	;
	if v278 != 0 {
		goto L83
	} else {
		goto L85
	}
L85:
	;
	F_addReplyStatus(m, l0, int32(_a_F_bgsaveCommand_14))
	mBase = m.M
	v282 = m.ExcPending
	if v282 != 0 {
		goto L38
	} else {
		goto L86
	}
L86:
	;
	goto L1
L87:
	;
	goto L1
}
func F_bitfieldGetKeys(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v58 int32
	_ = v58
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
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
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v110 int32
	_ = v110
	var v117 int32
	_ = v117
	var v120 int32
	_ = v120
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v126 int32
	_ = v126
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
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v158 int32
	_ = v158
	var v161 int32
	_ = v161
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v167 int32
	_ = v167
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
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v215 int32
	_ = v215
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v221 int32
	_ = v221
	var v225 int32
	_ = v225
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v247 int32
	_ = v247
	var v248 int32
	_ = v248
	var v249 int32
	_ = v249
	var v254 int32
	_ = v254
	var v258 int32
	_ = v258
	var v277 int32
	_ = v277
	var v286 int32
	_ = v286
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l3)+8))
	if v10 != 0 {
		v15 = v10
		goto L2
	} else {
		goto L3
	}
L1:
	;
	F__serverAssert(m, int32(_a_F_bitfieldGetKeys_0), int32(_a_F_bitfieldGetKeys_1), int32(2292))
	mBase = m.M
	v286 = m.ExcPending
	if v286 != 0 {
		goto L10
	} else {
		goto L83
	}
L2:
	;
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l3)+4))
	if int32(0) < v16 {
		v45 = v15
		goto L5
	} else {
		goto L6
	}
L3:
	;
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	if v11 != 0 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v13 = l3 + int32(12)
	*(*int32)(unsafe.Add(mBase, uint32(l3)+8)) = v13
	v15 = v13
	goto L2
L5:
	;
	v48 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v45))) = v48
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v48
	if l2 < int32(3) {
		goto L18
	} else {
		goto L19
	}
L6:
	;
	v20 = l3 + int32(12)
	if v15 == v20 {
		goto L8
	} else {
		goto L9
	}
L7:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3)+4)) = int32(1)
	v45 = v41
	goto L5
L8:
	;
	v29 = F_valkey_malloc(m, int32(8))
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		goto L10
	} else {
		goto L12
	}
L9:
	;
	v23 = F_valkey_realloc(m, v15, int32(8))
	mBase = m.M
	v26 = m.ExcPending
	if v26 != 0 {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	return int32(0)
L11:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3)+8)) = v23
	v41 = v23
	goto L7
L12:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3)+8)) = v29
	v32 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	if v32 == int32(0) {
		v41 = v29
		goto L7
	} else {
		goto L13
	}
L13:
	;
	v36 = v32 << (uint(int32(3)) % 32)
	if v36 == int32(0) {
		goto L15
	} else {
		goto L16
	}
L14:
	;
	v41 = v29
	goto L7
L15:
	;
	goto L14
L16:
	;
	v39 = F__emscripten_memcpy_bulkmem(m, v29, v20, v36)
	mBase = m.M
	goto L15
L17:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v45)+4)) = v277
	return int32(1)
L18:
	;
	v277 = int32(17)
	goto L17
L19:
	;
	v58 = int32(2)
	goto L20
L20:
	;
	v67 = *(*int32)(unsafe.Add(mBase, uint32(l1+v58<<(uint(int32(2))%32))))
	v68 = F_objectGetVal(m, v67)
	mBase = m.M
	v69 = int32(_a_F_bitfieldGetKeys_2)
	v72 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v68))))
	if v72 != 0 {
		goto L24
	} else {
		goto L25
	}
L21:
	;
	goto L18
L22:
	;
	v110 = l2 + (v58 ^ int32(-1))
	if v110 < int32(2) {
		goto L35
	} else {
		goto L36
	}
L23:
	;
	v104 = F_tolower(m, v100)
	mBase = m.M
	v105 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v101))))
	v106 = F_tolower(m, v105)
	mBase = m.M
	goto L22
L24:
	;
	v74 = v68
	v75 = v69
	v76 = v72
	goto L27
L25:
	;
	v100 = int32(0)
	v101 = v69
	goto L23
L26:
	;
	v100 = v97 & int32(255)
	v101 = v96
	goto L23
L27:
	;
	v78 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v75))))
	if v78 == int32(0) {
		v96 = v75
		v97 = v76
		goto L26
	} else {
		goto L29
	}
L28:
	;
	v96 = v90
	v97 = int32(0)
	goto L26
L29:
	;
	v82 = v76 & int32(255)
	if v82 == v78 {
		goto L30
	} else {
		goto L31
	}
L30:
	;
	v89 = int32(1)
	v90 = v75 + v89
	v91 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v74)+1)))
	if v91 != 0 {
		v74 = v74 + v89
		v75 = v90
		v76 = v91
		goto L27
	} else {
		goto L33
	}
L31:
	;
	v84 = F_tolower(m, v82)
	mBase = m.M
	v85 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v75))))
	v86 = F_tolower(m, v85)
	mBase = m.M
	if v84 == v86 {
		goto L30
	} else {
		goto L32
	}
L32:
	;
	v88 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v74))))
	v96 = v75
	v97 = v88
	goto L26
L33:
	;
	goto L28
L34:
	;
	v258 = v58 + v254 + int32(1)
	if v258 < l2 {
		v58 = v258
		goto L20
	} else {
		goto L82
	}
L35:
	;
	v117 = int32(_a_F_bitfieldGetKeys_3)
	v120 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v68))))
	if v120 != 0 {
		goto L42
	} else {
		goto L43
	}
L36:
	;
	if v104-v106 == int32(0) {
		v254 = int32(2)
		goto L34
	} else {
		goto L37
	}
L37:
	;
	goto L35
L38:
	;
	v211 = int32(50)
	v212 = int32(_a_F_bitfieldGetKeys_4)
	v215 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v68))))
	if v215 != 0 {
		goto L70
	} else {
		goto L71
	}
L39:
	;
	if v110 <= int32(2) {
		goto L38
	} else {
		goto L67
	}
L40:
	;
	if v152-v154 == int32(0) {
		goto L39
	} else {
		goto L52
	}
L41:
	;
	v152 = F_tolower(m, v148)
	mBase = m.M
	v153 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v149))))
	v154 = F_tolower(m, v153)
	mBase = m.M
	goto L40
L42:
	;
	v122 = v68
	v123 = v117
	v124 = v120
	goto L45
L43:
	;
	v148 = int32(0)
	v149 = v117
	goto L41
L44:
	;
	v148 = v145 & int32(255)
	v149 = v144
	goto L41
L45:
	;
	v126 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v123))))
	if v126 == int32(0) {
		v144 = v123
		v145 = v124
		goto L44
	} else {
		goto L47
	}
L46:
	;
	v144 = v138
	v145 = int32(0)
	goto L44
L47:
	;
	v130 = v124 & int32(255)
	if v130 == v126 {
		goto L48
	} else {
		goto L49
	}
L48:
	;
	v137 = int32(1)
	v138 = v123 + v137
	v139 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v122)+1)))
	if v139 != 0 {
		v122 = v122 + v137
		v123 = v138
		v124 = v139
		goto L45
	} else {
		goto L51
	}
L49:
	;
	v132 = F_tolower(m, v130)
	mBase = m.M
	v133 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v123))))
	v134 = F_tolower(m, v133)
	mBase = m.M
	if v132 == v134 {
		goto L48
	} else {
		goto L50
	}
L50:
	;
	v136 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v122))))
	v144 = v123
	v145 = v136
	goto L44
L51:
	;
	goto L46
L52:
	;
	v158 = int32(_a_F_bitfieldGetKeys_5)
	v161 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v68))))
	if v161 != 0 {
		goto L55
	} else {
		goto L56
	}
L53:
	;
	if v110 < int32(3) {
		goto L38
	} else {
		goto L65
	}
L54:
	;
	v193 = F_tolower(m, v189)
	mBase = m.M
	v194 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v190))))
	v195 = F_tolower(m, v194)
	mBase = m.M
	goto L53
L55:
	;
	v163 = v68
	v164 = v158
	v165 = v161
	goto L58
L56:
	;
	v189 = int32(0)
	v190 = v158
	goto L54
L57:
	;
	v189 = v186 & int32(255)
	v190 = v185
	goto L54
L58:
	;
	v167 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v164))))
	if v167 == int32(0) {
		v185 = v164
		v186 = v165
		goto L57
	} else {
		goto L60
	}
L59:
	;
	v185 = v179
	v186 = int32(0)
	goto L57
L60:
	;
	v171 = v165 & int32(255)
	if v171 == v167 {
		goto L61
	} else {
		goto L62
	}
L61:
	;
	v178 = int32(1)
	v179 = v164 + v178
	v180 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v163)+1)))
	if v180 != 0 {
		v163 = v163 + v178
		v164 = v179
		v165 = v180
		goto L58
	} else {
		goto L64
	}
L62:
	;
	v173 = F_tolower(m, v171)
	mBase = m.M
	v174 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v164))))
	v175 = F_tolower(m, v174)
	mBase = m.M
	if v173 == v175 {
		goto L61
	} else {
		goto L63
	}
L63:
	;
	v177 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v163))))
	v185 = v164
	v186 = v177
	goto L57
L64:
	;
	goto L59
L65:
	;
	if v193-v195 != 0 {
		goto L38
	} else {
		goto L66
	}
L66:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v45)+4)) = int32(50)
	return int32(1)
L67:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v45)+4)) = int32(50)
	return int32(1)
L68:
	;
	if v110 < int32(1) {
		v277 = v211
		goto L17
	} else {
		goto L80
	}
L69:
	;
	v247 = F_tolower(m, v243)
	mBase = m.M
	v248 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v244))))
	v249 = F_tolower(m, v248)
	mBase = m.M
	goto L68
L70:
	;
	v217 = v68
	v218 = v212
	v219 = v215
	goto L73
L71:
	;
	v243 = int32(0)
	v244 = v212
	goto L69
L72:
	;
	v243 = v240 & int32(255)
	v244 = v239
	goto L69
L73:
	;
	v221 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v218))))
	if v221 == int32(0) {
		v239 = v218
		v240 = v219
		goto L72
	} else {
		goto L75
	}
L74:
	;
	v239 = v233
	v240 = int32(0)
	goto L72
L75:
	;
	v225 = v219 & int32(255)
	if v225 == v221 {
		goto L76
	} else {
		goto L77
	}
L76:
	;
	v232 = int32(1)
	v233 = v218 + v232
	v234 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v217)+1)))
	if v234 != 0 {
		v217 = v217 + v232
		v218 = v233
		v219 = v234
		goto L73
	} else {
		goto L79
	}
L77:
	;
	v227 = F_tolower(m, v225)
	mBase = m.M
	v228 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v218))))
	v229 = F_tolower(m, v228)
	mBase = m.M
	if v227 == v229 {
		goto L76
	} else {
		goto L78
	}
L78:
	;
	v231 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v217))))
	v239 = v218
	v240 = v231
	goto L72
L79:
	;
	goto L74
L80:
	;
	if v247-v249 != 0 {
		v277 = v211
		goto L17
	} else {
		goto L81
	}
L81:
	;
	v254 = int32(1)
	goto L34
L82:
	;
	goto L21
L83:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_bitopCommand(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
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
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
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
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
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
	var v84 int32
	_ = v84
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
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
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v123 int32
	_ = v123
	var v130 int32
	_ = v130
	var v133 int32
	_ = v133
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v139 int32
	_ = v139
	var v143 int32
	_ = v143
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v174 int32
	_ = v174
	var v177 int32
	_ = v177
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v183 int32
	_ = v183
	var v187 int32
	_ = v187
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v216 int32
	_ = v216
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v223 int32
	_ = v223
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v242 int32
	_ = v242
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v248 int32
	_ = v248
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v254 int32
	_ = v254
	var v255 int32
	_ = v255
	var v258 int32
	_ = v258
	var v259 int32
	_ = v259
	var v260 int32
	_ = v260
	var v264 int32
	_ = v264
	var v266 int32
	_ = v266
	var v267 int32
	_ = v267
	var v268 int32
	_ = v268
	var v269 int32
	_ = v269
	var v271 int32
	_ = v271
	var v273 int32
	_ = v273
	var v274 int32
	_ = v274
	var v277 int32
	_ = v277
	var v282 int32
	_ = v282
	var v285 int32
	_ = v285
	var v288 int32
	_ = v288
	var v291 int32
	_ = v291
	var v294 int32
	_ = v294
	var v301 int32
	_ = v301
	var v303 int32
	_ = v303
	var v323 int32
	_ = v323
	var v324 int32
	_ = v324
	var v325 int32
	_ = v325
	var v331 int32
	_ = v331
	var v332 int32
	_ = v332
	var v334 int32
	_ = v334
	var v338 int32
	_ = v338
	var v339 int32
	_ = v339
	var v340 int32
	_ = v340
	var v342 int32
	_ = v342
	var v352 int32
	_ = v352
	var v354 int32
	_ = v354
	var v355 int32
	_ = v355
	var v378 int32
	_ = v378
	var v386 int32
	_ = v386
	var v390 int32
	_ = v390
	var v392 int32
	_ = v392
	var v394 int32
	_ = v394
	var v395 int32
	_ = v395
	var v396 int32
	_ = v396
	var v399 int32
	_ = v399
	var v402 int32
	_ = v402
	var v403 int32
	_ = v403
	var v406 int32
	_ = v406
	var v413 int32
	_ = v413
	var v416 int32
	_ = v416
	var v419 int32
	_ = v419
	var v422 int32
	_ = v422
	var v423 int32
	_ = v423
	var v427 int32
	_ = v427
	var v429 int32
	_ = v429
	var v431 int32
	_ = v431
	var v432 int32
	_ = v432
	var v436 int32
	_ = v436
	var v455 int32
	_ = v455
	var v457 int32
	_ = v457
	var v466 int32
	_ = v466
	var v468 int32
	_ = v468
	var v469 int32
	_ = v469
	var v476 int32
	_ = v476
	var v477 int32
	_ = v477
	var v478 int32
	_ = v478
	var v481 int32
	_ = v481
	var v482 int32
	_ = v482
	var v493 int32
	_ = v493
	var v506 int32
	_ = v506
	var v514 int32
	_ = v514
	var v515 int32
	_ = v515
	var v516 int32
	_ = v516
	var v517 int32
	_ = v517
	var v524 int32
	_ = v524
	var v527 int32
	_ = v527
	var v530 int32
	_ = v530
	var v537 int32
	_ = v537
	var v540 int32
	_ = v540
	var v546 int32
	_ = v546
	var v547 int32
	_ = v547
	var v548 int32
	_ = v548
	var v549 int32
	_ = v549
	var v551 int32
	_ = v551
	var v552 int32
	_ = v552
	var v554 int32
	_ = v554
	var v555 int32
	_ = v555
	var v557 int32
	_ = v557
	var v558 int32
	_ = v558
	var v564 int32
	_ = v564
	var v594 int32
	_ = v594
	var v607 int32
	_ = v607
	var v620 int32
	_ = v620
	var v628 int32
	_ = v628
	var v629 int32
	_ = v629
	var v630 int32
	_ = v630
	var v631 int32
	_ = v631
	var v638 int32
	_ = v638
	var v641 int32
	_ = v641
	var v644 int32
	_ = v644
	var v651 int32
	_ = v651
	var v654 int32
	_ = v654
	var v660 int32
	_ = v660
	var v661 int32
	_ = v661
	var v662 int32
	_ = v662
	var v663 int32
	_ = v663
	var v665 int32
	_ = v665
	var v666 int32
	_ = v666
	var v668 int32
	_ = v668
	var v669 int32
	_ = v669
	var v671 int32
	_ = v671
	var v672 int32
	_ = v672
	var v678 int32
	_ = v678
	var v708 int32
	_ = v708
	var v721 int32
	_ = v721
	var v734 int32
	_ = v734
	var v742 int32
	_ = v742
	var v743 int32
	_ = v743
	var v744 int32
	_ = v744
	var v745 int32
	_ = v745
	var v752 int32
	_ = v752
	var v755 int32
	_ = v755
	var v758 int32
	_ = v758
	var v765 int32
	_ = v765
	var v768 int32
	_ = v768
	var v774 int32
	_ = v774
	var v775 int32
	_ = v775
	var v776 int32
	_ = v776
	var v777 int32
	_ = v777
	var v779 int32
	_ = v779
	var v780 int32
	_ = v780
	var v782 int32
	_ = v782
	var v783 int32
	_ = v783
	var v785 int32
	_ = v785
	var v786 int32
	_ = v786
	var v792 int32
	_ = v792
	var v822 int32
	_ = v822
	var v825 int32
	_ = v825
	var v829 int32
	_ = v829
	var v832 int32
	_ = v832
	var v833 int32
	_ = v833
	var v836 int32
	_ = v836
	var v840 int32
	_ = v840
	var v844 int32
	_ = v844
	var v850 int32
	_ = v850
	var v851 int32
	_ = v851
	var v853 int32
	_ = v853
	var v858 int32
	_ = v858
	var v874 int32
	_ = v874
	var v881 int32
	_ = v881
	var v882 int32
	_ = v882
	var v885 int32
	_ = v885
	var v889 int32
	_ = v889
	var v893 int32
	_ = v893
	var v897 int32
	_ = v897
	var v901 int32
	_ = v901
	var v905 int32
	_ = v905
	var v909 int32
	_ = v909
	var v916 int32
	_ = v916
	var v939 int32
	_ = v939
	var v967 int32
	_ = v967
	var v973 int32
	_ = v973
	var v975 int32
	_ = v975
	var v977 int32
	_ = v977
	var v978 int32
	_ = v978
	var v979 int32
	_ = v979
	var v980 int32
	_ = v980
	var v985 int32
	_ = v985
	var v987 int32
	_ = v987
	var v1010 int32
	_ = v1010
	var v1012 int32
	_ = v1012
	var v1015 int32
	_ = v1015
	var v1017 int32
	_ = v1017
	var v1018 int32
	_ = v1018
	var v1020 int32
	_ = v1020
	var v1024 int32
	_ = v1024
	var v1025 int32
	_ = v1025
	var v1030 int32
	_ = v1030
	var v1033 int32
	_ = v1033
	var v1046 int32
	_ = v1046
	var v1063 int32
	_ = v1063
	var v1106 int64
	_ = v1106
	var v1107 int32
	_ = v1107
	var v1140 int32
	_ = v1140
	var v1146 int32
	_ = v1146
	var v1150 int32
	_ = v1150
	var v1152 int32
	_ = v1152
	var v1158 int32
	_ = v1158
	var v1169 int64
	_ = v1169
	var v1170 int32
	_ = v1170
	var v1180 int32
	_ = v1180
	var v1182 int32
	_ = v1182
	var v1184 int32
	_ = v1184
	var v1186 int32
	_ = v1186
	var v1187 int32
	_ = v1187
	var v1189 int32
	_ = v1189
	var v1194 int32
	_ = v1194
	var v1197 int32
	_ = v1197
	var v1198 int32
	_ = v1198
	var v1200 int32
	_ = v1200
	var v1201 int32
	_ = v1201
	var v1202 int32
	_ = v1202
	var v1203 int32
	_ = v1203
	var v1206 int32
	_ = v1206
	var v1208 int32
	_ = v1208
	var v1211 int32
	_ = v1211
	var v1212 int32
	_ = v1212
	var v1214 int32
	_ = v1214
	var v1215 int32
	_ = v1215
	var v1217 int64
	_ = v1217
	var v1222 int32
	_ = v1222
	var v1249 int32
	_ = v1249
	var v1251 int32
	_ = v1251
	var v1253 int32
	_ = v1253
	v26 = m.G0
	v28 = v26 - int32(80)
	m.G0 = v28
	v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v30)+4))
	v32 = F_objectGetVal(m, v31)
	mBase = m.M
	v33 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v33)+8))
	v35 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v32))))
	switch v35 + int32(-65) {
	case 0, 32:
		goto L9
	default:
		goto L5
	case 13, 45:
		goto L6
	case 14, 46:
		goto L8
	case 23, 55:
		goto L7
	}
L1:
	;
	m.G0 = v28 + int32(80)
	return
L2:
	;
	v240 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v242 = v240 + int32(-3)
	v244 = v242 << (uint(int32(2)) % 32)
	v245 = F_valkey_malloc(m, v244)
	mBase = m.M
	v246 = m.ExcPending
	if v246 != 0 {
		goto L68
	} else {
		goto L73
	}
L3:
	;
	v233 = int32(0)
	v235 = v233
	v236 = v230
	v237 = v231
	v238 = v232
	v239 = v233
	goto L2
L4:
	;
	v219 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v219 != int32(4) {
		goto L70
	} else {
		goto L71
	}
L5:
	;
	v216 = *(*int32)(unsafe.Add(mBase, _c_F_bitopCommand[0]))
	F_addReplyErrorObject(m, l0, v216)
	mBase = m.M
	v218 = m.ExcPending
	if v218 != 0 {
		goto L68
	} else {
		goto L69
	}
L6:
	;
	v174 = int32(_a_F_bitopCommand_0)
	v177 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v32))))
	if v177 != 0 {
		goto L57
	} else {
		goto L58
	}
L7:
	;
	v130 = int32(_a_F_bitopCommand_1)
	v133 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v32))))
	if v133 != 0 {
		goto L43
	} else {
		goto L44
	}
L8:
	;
	v84 = int32(_a_F_bitopCommand_2)
	v87 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v32))))
	if v87 != 0 {
		goto L27
	} else {
		goto L28
	}
L9:
	;
	v38 = int32(_a_F_bitopCommand_3)
	v41 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v32))))
	if v41 != 0 {
		goto L13
	} else {
		goto L14
	}
L10:
	;
	switch v35 + int32(-78) {
	case 0:
		goto L6
	case 1:
		goto L8
	default:
		goto L5
	case 10:
		goto L7
	}
L11:
	;
	if v73-v75 != 0 {
		goto L10
	} else {
		goto L23
	}
L12:
	;
	v73 = F_tolower(m, v69)
	mBase = m.M
	v74 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v70))))
	v75 = F_tolower(m, v74)
	mBase = m.M
	goto L11
L13:
	;
	v43 = v32
	v44 = v38
	v45 = v41
	goto L16
L14:
	;
	v69 = int32(0)
	v70 = v38
	goto L12
L15:
	;
	v69 = v66 & int32(255)
	v70 = v65
	goto L12
L16:
	;
	v47 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v44))))
	if v47 == int32(0) {
		v65 = v44
		v66 = v45
		goto L15
	} else {
		goto L18
	}
L17:
	;
	v65 = v59
	v66 = int32(0)
	goto L15
L18:
	;
	v51 = v45 & int32(255)
	if v51 == v47 {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	v58 = int32(1)
	v59 = v44 + v58
	v60 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v43)+1)))
	if v60 != 0 {
		v43 = v43 + v58
		v44 = v59
		v45 = v60
		goto L16
	} else {
		goto L22
	}
L20:
	;
	v53 = F_tolower(m, v51)
	mBase = m.M
	v54 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v44))))
	v55 = F_tolower(m, v54)
	mBase = m.M
	if v53 == v55 {
		goto L19
	} else {
		goto L21
	}
L21:
	;
	v57 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v43))))
	v65 = v44
	v66 = v57
	goto L15
L22:
	;
	goto L17
L23:
	;
	v78 = int32(0)
	v235 = int32(1)
	v236 = v78
	v237 = v78
	v238 = v78
	v239 = v78
	goto L2
L24:
	;
	if v35 == int32(110) {
		goto L6
	} else {
		goto L38
	}
L25:
	;
	if v119-v121 != 0 {
		goto L24
	} else {
		goto L37
	}
L26:
	;
	v119 = F_tolower(m, v115)
	mBase = m.M
	v120 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v116))))
	v121 = F_tolower(m, v120)
	mBase = m.M
	goto L25
L27:
	;
	v89 = v32
	v90 = v84
	v91 = v87
	goto L30
L28:
	;
	v115 = int32(0)
	v116 = v84
	goto L26
L29:
	;
	v115 = v112 & int32(255)
	v116 = v111
	goto L26
L30:
	;
	v93 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v90))))
	if v93 == int32(0) {
		v111 = v90
		v112 = v91
		goto L29
	} else {
		goto L32
	}
L31:
	;
	v111 = v105
	v112 = int32(0)
	goto L29
L32:
	;
	v97 = v91 & int32(255)
	if v97 == v93 {
		goto L33
	} else {
		goto L34
	}
L33:
	;
	v104 = int32(1)
	v105 = v90 + v104
	v106 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v89)+1)))
	if v106 != 0 {
		v89 = v89 + v104
		v90 = v105
		v91 = v106
		goto L30
	} else {
		goto L36
	}
L34:
	;
	v99 = F_tolower(m, v97)
	mBase = m.M
	v100 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v90))))
	v101 = F_tolower(m, v100)
	mBase = m.M
	if v99 == v101 {
		goto L33
	} else {
		goto L35
	}
L35:
	;
	v103 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v89))))
	v111 = v90
	v112 = v103
	goto L29
L36:
	;
	goto L31
L37:
	;
	v123 = int32(1)
	v230 = v123
	v231 = int32(0)
	v232 = v123
	goto L3
L38:
	;
	if v35 != int32(88) {
		goto L5
	} else {
		goto L39
	}
L39:
	;
	goto L7
L40:
	;
	if v35 != int32(110) {
		goto L5
	} else {
		goto L54
	}
L41:
	;
	if v165-v167 != 0 {
		goto L40
	} else {
		goto L53
	}
L42:
	;
	v165 = F_tolower(m, v161)
	mBase = m.M
	v166 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v162))))
	v167 = F_tolower(m, v166)
	mBase = m.M
	goto L41
L43:
	;
	v135 = v32
	v136 = v130
	v137 = v133
	goto L46
L44:
	;
	v161 = int32(0)
	v162 = v130
	goto L42
L45:
	;
	v161 = v158 & int32(255)
	v162 = v157
	goto L42
L46:
	;
	v139 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v136))))
	if v139 == int32(0) {
		v157 = v136
		v158 = v137
		goto L45
	} else {
		goto L48
	}
L47:
	;
	v157 = v151
	v158 = int32(0)
	goto L45
L48:
	;
	v143 = v137 & int32(255)
	if v143 == v139 {
		goto L49
	} else {
		goto L50
	}
L49:
	;
	v150 = int32(1)
	v151 = v136 + v150
	v152 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v135)+1)))
	if v152 != 0 {
		v135 = v135 + v150
		v136 = v151
		v137 = v152
		goto L46
	} else {
		goto L52
	}
L50:
	;
	v145 = F_tolower(m, v143)
	mBase = m.M
	v146 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v136))))
	v147 = F_tolower(m, v146)
	mBase = m.M
	if v145 == v147 {
		goto L49
	} else {
		goto L51
	}
L51:
	;
	v149 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v135))))
	v157 = v136
	v158 = v149
	goto L45
L52:
	;
	goto L47
L53:
	;
	v230 = int32(2)
	v231 = int32(1)
	v232 = int32(0)
	goto L3
L54:
	;
	goto L6
L55:
	;
	if v209-v211 == int32(0) {
		goto L4
	} else {
		goto L67
	}
L56:
	;
	v209 = F_tolower(m, v205)
	mBase = m.M
	v210 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v206))))
	v211 = F_tolower(m, v210)
	mBase = m.M
	goto L55
L57:
	;
	v179 = v32
	v180 = v174
	v181 = v177
	goto L60
L58:
	;
	v205 = int32(0)
	v206 = v174
	goto L56
L59:
	;
	v205 = v202 & int32(255)
	v206 = v201
	goto L56
L60:
	;
	v183 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v180))))
	if v183 == int32(0) {
		v201 = v180
		v202 = v181
		goto L59
	} else {
		goto L62
	}
L61:
	;
	v201 = v195
	v202 = int32(0)
	goto L59
L62:
	;
	v187 = v181 & int32(255)
	if v187 == v183 {
		goto L63
	} else {
		goto L64
	}
L63:
	;
	v194 = int32(1)
	v195 = v180 + v194
	v196 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v179)+1)))
	if v196 != 0 {
		v179 = v179 + v194
		v180 = v195
		v181 = v196
		goto L60
	} else {
		goto L66
	}
L64:
	;
	v189 = F_tolower(m, v187)
	mBase = m.M
	v190 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v180))))
	v191 = F_tolower(m, v190)
	mBase = m.M
	if v189 == v191 {
		goto L63
	} else {
		goto L65
	}
L65:
	;
	v193 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v179))))
	v201 = v180
	v202 = v193
	goto L59
L66:
	;
	goto L61
L67:
	;
	goto L5
L68:
	;
	return
L69:
	;
	goto L1
L70:
	;
	F_addReplyError(m, l0, int32(_a_F_bitopCommand_4))
	mBase = m.M
	v229 = m.ExcPending
	if v229 != 0 {
		goto L68
	} else {
		goto L72
	}
L71:
	;
	v223 = int32(0)
	v235 = v223
	v236 = int32(3)
	v237 = v223
	v238 = v223
	v239 = int32(1)
	goto L2
L72:
	;
	goto L1
L73:
	;
	v247 = F_valkey_malloc(m, v244)
	mBase = m.M
	v248 = m.ExcPending
	if v248 != 0 {
		goto L68
	} else {
		goto L74
	}
L74:
	;
	v249 = F_valkey_malloc(m, v244)
	mBase = m.M
	v250 = m.ExcPending
	if v250 != 0 {
		goto L68
	} else {
		goto L75
	}
L75:
	;
	if v242 != 0 {
		goto L78
	} else {
		goto L79
	}
L76:
	;
	F_valkey_free(m, v245)
	mBase = m.M
	v1249 = m.ExcPending
	if v1249 != 0 {
		goto L68
	} else {
		goto L215
	}
L77:
	;
	F_valkey_free(m, v245)
	mBase = m.M
	v1180 = m.ExcPending
	if v1180 != 0 {
		goto L68
	} else {
		goto L200
	}
L78:
	;
	v254 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v255 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v258 = *(*int32)(unsafe.Add(mBase, uint32(v255+int32(12))))
	v259 = F_lookupKeyRead(m, v254, v258)
	mBase = m.M
	v260 = m.ExcPending
	if v260 != 0 {
		goto L68
	} else {
		goto L80
	}
L79:
	;
	v1158 = int32(1)
	v1169 = int64(0)
	v1170 = int32(0)
	goto L77
L80:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+76)) = v259
	if v259 == int32(0) {
		goto L82
	} else {
		goto L83
	}
L81:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v247))) = v301
	v303 = int32(1)
	if v242 != v303 {
		goto L93
	} else {
		goto L94
	}
L82:
	;
	v294 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v249))) = v294
	*(*int32)(unsafe.Add(mBase, uint32(v245))) = v294
	v301 = v294
	goto L81
L83:
	;
	v264 = int32(0)
	v266 = F_checkType(m, l0, v259, v264)
	mBase = m.M
	v267 = m.ExcPending
	if v267 != 0 {
		goto L68
	} else {
		goto L84
	}
L84:
	;
	if v266 != 0 {
		goto L76
	} else {
		goto L85
	}
L85:
	;
	v268 = F_getDecodedObject(m, v259)
	mBase = m.M
	v269 = m.ExcPending
	if v269 != 0 {
		goto L68
	} else {
		goto L86
	}
L86:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v249))) = v268
	v271 = F_objectGetVal(m, v268)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v245))) = v271
	v273 = *(*int32)(unsafe.Add(mBase, uint32(v249)))
	v274 = F_objectGetVal(m, v273)
	mBase = m.M
	v277 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v274+int32(-1)))))
	switch v277 & int32(7) {
	case 0:
		goto L87
	case 1:
		goto L88
	case 2:
		goto L89
	case 3:
		goto L90
	case 4:
		goto L91
	default:
		v301 = v264
		goto L81
	}
L87:
	;
	v301 = int32(base.Ui32(v277) >> (uint(int32(3)) % 32))
	goto L81
L88:
	;
	v291 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v274+int32(-3)))))
	v301 = v291
	goto L81
L89:
	;
	v288 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v274+int32(-5)))))
	v301 = v288
	goto L81
L90:
	;
	v285 = *(*int32)(unsafe.Add(mBase, uint32(v274+int32(-9))))
	v301 = v285
	goto L81
L91:
	;
	v282 = *(*int32)(unsafe.Add(mBase, uint32(v274+int32(-17))))
	v301 = v282
	goto L81
L92:
	;
	if v457 != 0 {
		goto L125
	} else {
		goto L126
	}
L93:
	;
	v323 = v301
	v324 = v303
	v325 = v301
	goto L95
L94:
	;
	v455 = v301
	v457 = v301
	goto L92
L95:
	;
	v331 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v332 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v334 = v324 << (uint(int32(2)) % 32)
	v338 = *(*int32)(unsafe.Add(mBase, uint32(v332+v334+int32(12))))
	v339 = F_lookupKeyRead(m, v331, v338)
	mBase = m.M
	v340 = m.ExcPending
	if v340 != 0 {
		goto L68
	} else {
		goto L97
	}
L96:
	;
	v455 = v431
	v457 = v432
	goto L92
L97:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+76)) = v339
	if v339 != 0 {
		goto L99
	} else {
		goto L100
	}
L98:
	;
	v436 = v324 + int32(1)
	if v436 != v242 {
		v323 = v431
		v324 = v436
		v325 = v432
		goto L95
	} else {
		goto L123
	}
L99:
	;
	v352 = int32(0)
	v354 = F_checkType(m, l0, v339, v352)
	mBase = m.M
	v355 = m.ExcPending
	if v355 != 0 {
		goto L68
	} else {
		goto L102
	}
L100:
	;
	v342 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v249+v334))) = v342
	*(*int32)(unsafe.Add(mBase, uint32(v245+v334))) = v342
	*(*int32)(unsafe.Add(mBase, uint32(v247+v334))) = v342
	v431 = v342
	v432 = v325
	goto L98
L101:
	;
	v394 = v249 + v334
	v395 = F_getDecodedObject(m, v339)
	mBase = m.M
	v396 = m.ExcPending
	if v396 != 0 {
		goto L68
	} else {
		goto L110
	}
L102:
	;
	if v354 == int32(0) {
		goto L101
	} else {
		goto L103
	}
L103:
	;
	v378 = v352
	goto L104
L104:
	;
	v386 = *(*int32)(unsafe.Add(mBase, uint32(v249+v378<<(uint(int32(2))%32))))
	if v386 == int32(0) {
		goto L106
	} else {
		goto L107
	}
L106:
	;
	v392 = v378 + int32(1)
	if v392 != v324 {
		v378 = v392
		goto L104
	} else {
		goto L109
	}
L107:
	;
	F_decrRefCount(m, v386)
	mBase = m.M
	v390 = m.ExcPending
	if v390 != 0 {
		goto L68
	} else {
		goto L108
	}
L108:
	;
	goto L106
L109:
	;
	goto L76
L110:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v394))) = v395
	v399 = F_objectGetVal(m, v395)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v245+v334))) = v399
	v402 = *(*int32)(unsafe.Add(mBase, uint32(v394)))
	v403 = F_objectGetVal(m, v402)
	mBase = m.M
	v406 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v403+int32(-1)))))
	switch v406 & int32(7) {
	case 0:
		goto L116
	case 1:
		goto L115
	case 2:
		goto L114
	case 3:
		goto L113
	case 4:
		goto L112
	default:
		v423 = int32(0)
		goto L111
	}
L111:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v247+v334))) = v423
	if base.Ui32(v423) < base.Ui32(v323) {
		goto L117
	} else {
		goto L118
	}
L112:
	;
	v422 = *(*int32)(unsafe.Add(mBase, uint32(v403+int32(-17))))
	v423 = v422
	goto L111
L113:
	;
	v419 = *(*int32)(unsafe.Add(mBase, uint32(v403+int32(-9))))
	v423 = v419
	goto L111
L114:
	;
	v416 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v403+int32(-5)))))
	v423 = v416
	goto L111
L115:
	;
	v413 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v403+int32(-3)))))
	v423 = v413
	goto L111
L116:
	;
	v423 = int32(base.Ui32(v406) >> (uint(int32(3)) % 32))
	goto L111
L117:
	;
	v427 = v423
	goto L119
L118:
	;
	v427 = v323
	goto L119
L119:
	;
	if base.Ui32(v325) < base.Ui32(v423) {
		goto L120
	} else {
		goto L121
	}
L120:
	;
	v429 = v423
	goto L122
L121:
	;
	v429 = v325
	goto L122
L122:
	;
	v431 = v427
	v432 = v429
	goto L98
L123:
	;
	goto L96
L124:
	;
	v1140 = int32(0)
	goto L194
L125:
	;
	v466 = int32(0)
	v468 = F_sdsnewlen(m, v466, v457)
	mBase = m.M
	v469 = m.ExcPending
	if v469 != 0 {
		goto L68
	} else {
		goto L127
	}
L126:
	;
	v1106 = int64(0)
	v1107 = int32(0)
	goto L124
L127:
	;
	if base.Ui32(int32(16)) < base.Ui32(v242) {
		v939 = v466
		goto L128
	} else {
		goto L129
	}
L128:
	;
	if base.Ui32(v457) <= base.Ui32(v939) {
		goto L174
	} else {
		goto L175
	}
L129:
	;
	if base.Ui32(v455) < base.Ui32(int32(16)) {
		v939 = v466
		goto L128
	} else {
		goto L130
	}
L130:
	;
	if v244 == int32(0) {
		v477 = v28
		goto L132
	} else {
		goto L133
	}
L131:
	;
	v478 = *(*int32)(unsafe.Add(mBase, uint32(v245)))
	if v455 == int32(0) {
		v482 = v468
		goto L135
	} else {
		goto L136
	}
L132:
	;
	goto L131
L133:
	;
	v476 = F__emscripten_memcpy_bulkmem(m, v28, v245, v244)
	mBase = m.M
	v477 = v476
	goto L132
L134:
	;
	if v235 == int32(0) {
		goto L137
	} else {
		goto L138
	}
L135:
	;
	goto L134
L136:
	;
	v481 = F__emscripten_memcpy_bulkmem(m, v468, v478, v455)
	mBase = m.M
	v482 = v481
	goto L135
L137:
	;
	if v238 == int32(0) {
		goto L147
	} else {
		goto L148
	}
L138:
	;
	v493 = v482
	v506 = v455
	goto L139
L139:
	;
	if v242 == int32(1) {
		goto L141
	} else {
		goto L142
	}
L141:
	;
	v594 = v506 + int32(-16)
	if base.Ui32(int32(15)) < base.Ui32(v594) {
		v493 = v493 + int32(16)
		v506 = v594
		goto L139
	} else {
		goto L146
	}
L142:
	;
	v514 = *(*int32)(unsafe.Add(mBase, uint32(v493)+12))
	v515 = *(*int32)(unsafe.Add(mBase, uint32(v493)+8))
	v516 = *(*int32)(unsafe.Add(mBase, uint32(v493)+4))
	v517 = *(*int32)(unsafe.Add(mBase, uint32(v493)))
	v524 = v516
	v527 = v517
	v530 = v514
	v537 = int32(1)
	v540 = v515
	goto L143
L143:
	;
	v546 = v477 + v537<<(uint(int32(2))%32)
	v547 = *(*int32)(unsafe.Add(mBase, uint32(v546)))
	v548 = *(*int32)(unsafe.Add(mBase, uint32(v547)))
	v549 = v527 & v548
	*(*int32)(unsafe.Add(mBase, uint32(v493))) = v549
	v551 = *(*int32)(unsafe.Add(mBase, uint32(v547)+4))
	v552 = v524 & v551
	*(*int32)(unsafe.Add(mBase, uint32(v493)+4)) = v552
	v554 = *(*int32)(unsafe.Add(mBase, uint32(v547)+8))
	v555 = v540 & v554
	*(*int32)(unsafe.Add(mBase, uint32(v493)+8)) = v555
	v557 = *(*int32)(unsafe.Add(mBase, uint32(v547)+12))
	v558 = v530 & v557
	*(*int32)(unsafe.Add(mBase, uint32(v493)+12)) = v558
	*(*int32)(unsafe.Add(mBase, uint32(v546))) = v547 + int32(16)
	v564 = v537 + int32(1)
	if v564 != v242 {
		v524 = v552
		v527 = v549
		v530 = v558
		v537 = v564
		v540 = v555
		goto L143
	} else {
		goto L145
	}
L144:
	;
	goto L141
L145:
	;
	goto L144
L146:
	;
	v939 = v455 & int32(-16)
	goto L128
L147:
	;
	if v237 == int32(0) {
		goto L157
	} else {
		goto L158
	}
L148:
	;
	v607 = v482
	v620 = v455
	goto L149
L149:
	;
	if v242 == int32(1) {
		goto L151
	} else {
		goto L152
	}
L151:
	;
	v708 = v620 + int32(-16)
	if base.Ui32(int32(15)) < base.Ui32(v708) {
		v607 = v607 + int32(16)
		v620 = v708
		goto L149
	} else {
		goto L156
	}
L152:
	;
	v628 = *(*int32)(unsafe.Add(mBase, uint32(v607)+12))
	v629 = *(*int32)(unsafe.Add(mBase, uint32(v607)+8))
	v630 = *(*int32)(unsafe.Add(mBase, uint32(v607)+4))
	v631 = *(*int32)(unsafe.Add(mBase, uint32(v607)))
	v638 = v630
	v641 = v631
	v644 = v628
	v651 = int32(1)
	v654 = v629
	goto L153
L153:
	;
	v660 = v477 + v651<<(uint(int32(2))%32)
	v661 = *(*int32)(unsafe.Add(mBase, uint32(v660)))
	v662 = *(*int32)(unsafe.Add(mBase, uint32(v661)))
	v663 = v641 | v662
	*(*int32)(unsafe.Add(mBase, uint32(v607))) = v663
	v665 = *(*int32)(unsafe.Add(mBase, uint32(v661)+4))
	v666 = v638 | v665
	*(*int32)(unsafe.Add(mBase, uint32(v607)+4)) = v666
	v668 = *(*int32)(unsafe.Add(mBase, uint32(v661)+8))
	v669 = v654 | v668
	*(*int32)(unsafe.Add(mBase, uint32(v607)+8)) = v669
	v671 = *(*int32)(unsafe.Add(mBase, uint32(v661)+12))
	v672 = v644 | v671
	*(*int32)(unsafe.Add(mBase, uint32(v607)+12)) = v672
	*(*int32)(unsafe.Add(mBase, uint32(v660))) = v661 + int32(16)
	v678 = v651 + int32(1)
	if v678 != v242 {
		v638 = v666
		v641 = v663
		v644 = v672
		v651 = v678
		v654 = v669
		goto L153
	} else {
		goto L155
	}
L154:
	;
	goto L151
L155:
	;
	goto L154
L156:
	;
	v939 = v455 & int32(-16)
	goto L128
L157:
	;
	v825 = int32(0)
	if v239 == v825 {
		v939 = v825
		goto L128
	} else {
		goto L167
	}
L158:
	;
	v721 = v482
	v734 = v455
	goto L159
L159:
	;
	if v242 == int32(1) {
		goto L161
	} else {
		goto L162
	}
L161:
	;
	v822 = v734 + int32(-16)
	if base.Ui32(int32(15)) < base.Ui32(v822) {
		v721 = v721 + int32(16)
		v734 = v822
		goto L159
	} else {
		goto L166
	}
L162:
	;
	v742 = *(*int32)(unsafe.Add(mBase, uint32(v721)+12))
	v743 = *(*int32)(unsafe.Add(mBase, uint32(v721)+8))
	v744 = *(*int32)(unsafe.Add(mBase, uint32(v721)+4))
	v745 = *(*int32)(unsafe.Add(mBase, uint32(v721)))
	v752 = v744
	v755 = v745
	v758 = v742
	v765 = int32(1)
	v768 = v743
	goto L163
L163:
	;
	v774 = v477 + v765<<(uint(int32(2))%32)
	v775 = *(*int32)(unsafe.Add(mBase, uint32(v774)))
	v776 = *(*int32)(unsafe.Add(mBase, uint32(v775)))
	v777 = v755 ^ v776
	*(*int32)(unsafe.Add(mBase, uint32(v721))) = v777
	v779 = *(*int32)(unsafe.Add(mBase, uint32(v775)+4))
	v780 = v752 ^ v779
	*(*int32)(unsafe.Add(mBase, uint32(v721)+4)) = v780
	v782 = *(*int32)(unsafe.Add(mBase, uint32(v775)+8))
	v783 = v768 ^ v782
	*(*int32)(unsafe.Add(mBase, uint32(v721)+8)) = v783
	v785 = *(*int32)(unsafe.Add(mBase, uint32(v775)+12))
	v786 = v758 ^ v785
	*(*int32)(unsafe.Add(mBase, uint32(v721)+12)) = v786
	*(*int32)(unsafe.Add(mBase, uint32(v774))) = v775 + int32(16)
	v792 = v765 + int32(1)
	if v792 != v242 {
		v752 = v780
		v755 = v777
		v758 = v786
		v765 = v792
		v768 = v783
		goto L163
	} else {
		goto L165
	}
L164:
	;
	goto L161
L165:
	;
	goto L164
L166:
	;
	v939 = v455 & int32(-16)
	goto L128
L167:
	;
	v829 = v455 + int32(-16)
	if v829&int32(16) != 0 {
		v850 = v482
		v851 = v455
		goto L168
	} else {
		goto L169
	}
L168:
	;
	v853 = v455 & int32(-16)
	if base.Ui32(v829) < base.Ui32(int32(16)) {
		v939 = v853
		goto L128
	} else {
		goto L170
	}
L169:
	;
	v832 = *(*int32)(unsafe.Add(mBase, uint32(v482)))
	v833 = int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(v482))) = v832 ^ v833
	v836 = *(*int32)(unsafe.Add(mBase, uint32(v482)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v482)+4)) = v836 ^ v833
	v840 = *(*int32)(unsafe.Add(mBase, uint32(v482)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v482)+8)) = v840 ^ v833
	v844 = *(*int32)(unsafe.Add(mBase, uint32(v482)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v482)+12)) = v844 ^ v833
	v850 = v482 + int32(16)
	v851 = v829
	goto L168
L170:
	;
	v858 = v850
	v874 = v851
	goto L171
L171:
	;
	v881 = *(*int32)(unsafe.Add(mBase, uint32(v858)))
	v882 = int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(v858))) = v881 ^ v882
	v885 = *(*int32)(unsafe.Add(mBase, uint32(v858)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v858)+4)) = v885 ^ v882
	v889 = *(*int32)(unsafe.Add(mBase, uint32(v858)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v858)+8)) = v889 ^ v882
	v893 = *(*int32)(unsafe.Add(mBase, uint32(v858)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v858)+12)) = v893 ^ v882
	v897 = *(*int32)(unsafe.Add(mBase, uint32(v858)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v858)+16)) = v897 ^ v882
	v901 = *(*int32)(unsafe.Add(mBase, uint32(v858)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v858)+20)) = v901 ^ v882
	v905 = *(*int32)(unsafe.Add(mBase, uint32(v858)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v858)+24)) = v905 ^ v882
	v909 = *(*int32)(unsafe.Add(mBase, uint32(v858)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v858)+28)) = v909 ^ v882
	v916 = v874 + int32(-32)
	if base.Ui32(int32(15)) < base.Ui32(v916) {
		v858 = v858 + int32(32)
		v874 = v916
		goto L171
	} else {
		goto L173
	}
L172:
	;
	v939 = v853
	goto L128
L173:
	;
	goto L172
L174:
	;
	v1106 = base.I64_extend_i32_u(v457)
	v1107 = v468
	goto L124
L175:
	;
	v967 = v939
	goto L176
L176:
	;
	v973 = *(*int32)(unsafe.Add(mBase, uint32(v247)))
	if base.Ui32(v973) <= base.Ui32(v967) {
		v978 = int32(0)
		goto L178
	} else {
		goto L179
	}
L177:
	;
	goto L174
L178:
	;
	v979 = int32(1)
	v980 = v978 ^ (int32(0) - v239)
	if v242 == v979 {
		v1046 = v980
		goto L180
	} else {
		goto L181
	}
L179:
	;
	v975 = *(*int32)(unsafe.Add(mBase, uint32(v245)))
	v977 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v975+v967))))
	v978 = v977
	goto L178
L180:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v468+v967))) = uint8(v1046)
	v1063 = v967 + int32(1)
	if v1063 != v457 {
		v967 = v1063
		goto L176
	} else {
		goto L193
	}
L181:
	;
	v985 = v979
	v987 = v980
	goto L182
L182:
	;
	v1010 = v985 << (uint(int32(2)) % 32)
	v1012 = *(*int32)(unsafe.Add(mBase, uint32(v247+v1010)))
	if base.Ui32(v1012) <= base.Ui32(v967) {
		v1018 = int32(0)
		goto L184
	} else {
		goto L185
	}
L183:
	;
	v1046 = v1030
	goto L180
L184:
	;
	switch v236 {
	default:
		goto L188
	case 1:
		goto L187
	case 2:
		goto L189
	case 3:
		v1030 = v987
		goto L186
	}
L185:
	;
	v1015 = *(*int32)(unsafe.Add(mBase, uint32(v245+v1010)))
	v1017 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1015+v967))))
	v1018 = v1017
	goto L184
L186:
	;
	v1033 = v985 + int32(1)
	if v1033 != v242 {
		v985 = v1033
		v987 = v1030
		goto L182
	} else {
		goto L192
	}
L187:
	;
	v1024 = int32(255)
	v1025 = v1018 | v987
	if v1025&v1024 == v1024 {
		v1046 = v1024
		goto L180
	} else {
		goto L191
	}
L188:
	;
	v1020 = v1018 & v987
	if v1020&int32(255) != 0 {
		v1030 = v1020
		goto L186
	} else {
		goto L190
	}
L189:
	;
	v1030 = v1018 ^ v987
	goto L186
L190:
	;
	v1046 = int32(0)
	goto L180
L191:
	;
	v1030 = v1025
	goto L186
L192:
	;
	goto L183
L193:
	;
	goto L177
L194:
	;
	v1146 = *(*int32)(unsafe.Add(mBase, uint32(v249+v1140<<(uint(int32(2))%32))))
	if v1146 == int32(0) {
		goto L196
	} else {
		goto L197
	}
L195:
	;
	v1158 = base.B2i32(v457 == int32(0))
	v1169 = v1106
	v1170 = v1107
	goto L77
L196:
	;
	v1152 = v1140 + int32(1)
	if v1152 != v242 {
		v1140 = v1152
		goto L194
	} else {
		goto L199
	}
L197:
	;
	F_decrRefCount(m, v1146)
	mBase = m.M
	v1150 = m.ExcPending
	if v1150 != 0 {
		goto L68
	} else {
		goto L198
	}
L198:
	;
	goto L196
L199:
	;
	goto L195
L200:
	;
	F_valkey_free(m, v247)
	mBase = m.M
	v1182 = m.ExcPending
	if v1182 != 0 {
		goto L68
	} else {
		goto L201
	}
L201:
	;
	F_valkey_free(m, v249)
	mBase = m.M
	v1184 = m.ExcPending
	if v1184 != 0 {
		goto L68
	} else {
		goto L202
	}
L202:
	;
	if v1158 != 0 {
		goto L205
	} else {
		goto L206
	}
L203:
	;
	F_addReplyLongLong(m, l0, v1169)
	mBase = m.M
	v1222 = m.ExcPending
	if v1222 != 0 {
		goto L68
	} else {
		goto L214
	}
L204:
	;
	v1215 = int32(_a_F_bitopCommand_5)
	v1217 = *(*int64)(unsafe.Add(mBase, _c_F_bitopCommand[1]))
	*(*int64)(unsafe.Add(mBase, _c_F_bitopCommand[1])) = v1217 + int64(1)
	goto L203
L205:
	;
	v1201 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v1202 = F_dbDelete(m, v1201, v34)
	mBase = m.M
	v1203 = m.ExcPending
	if v1203 != 0 {
		goto L68
	} else {
		goto L210
	}
L206:
	;
	v1186 = F_createObject(m, int32(0), v1170)
	mBase = m.M
	v1187 = m.ExcPending
	if v1187 != 0 {
		goto L68
	} else {
		goto L207
	}
L207:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+76)) = v1186
	v1189 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	F_setKey(m, l0, v1189, v34, v28+int32(76), int32(0))
	mBase = m.M
	v1194 = m.ExcPending
	if v1194 != 0 {
		goto L68
	} else {
		goto L208
	}
L208:
	;
	v1197 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v1198 = *(*int32)(unsafe.Add(mBase, uint32(v1197)+28))
	F_notifyKeyspaceEvent(m, int32(8), int32(_a_F_bitopCommand_6), v34, v1198)
	mBase = m.M
	v1200 = m.ExcPending
	if v1200 != 0 {
		goto L68
	} else {
		goto L209
	}
L209:
	;
	goto L204
L210:
	;
	if v1202 == int32(0) {
		goto L203
	} else {
		goto L211
	}
L211:
	;
	v1206 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	F_signalModifiedKey(m, l0, v1206, v34)
	mBase = m.M
	v1208 = m.ExcPending
	if v1208 != 0 {
		goto L68
	} else {
		goto L212
	}
L212:
	;
	v1211 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v1212 = *(*int32)(unsafe.Add(mBase, uint32(v1211)+28))
	F_notifyKeyspaceEvent(m, int32(4), int32(_a_F_bitopCommand_7), v34, v1212)
	mBase = m.M
	v1214 = m.ExcPending
	if v1214 != 0 {
		goto L68
	} else {
		goto L213
	}
L213:
	;
	goto L204
L214:
	;
	goto L1
L215:
	;
	F_valkey_free(m, v247)
	mBase = m.M
	v1251 = m.ExcPending
	if v1251 != 0 {
		goto L68
	} else {
		goto L216
	}
L216:
	;
	F_valkey_free(m, v249)
	mBase = m.M
	v1253 = m.ExcPending
	if v1253 != 0 {
		goto L68
	} else {
		goto L217
	}
L217:
	;
	goto L1
}
func F_blmoveGenericCommand(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int64) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v34 int32
	_ = v34
	var v39 int32
	_ = v39
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(v7)+4))
	v9 = F_lookupKeyWrite(m, v6, v8)
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		return
	} else {
		v12 = F_checkType(m, l0, v9, int32(1))
		mBase = m.M
		v13 = m.ExcPending
		if v13 != 0 {
			return
		} else {
			if v12 != 0 {
				return
			} else {
				if v9 != 0 {
					v29 = F_listTypeLength(m, v9)
					mBase = m.M
					v30 = m.ExcPending
					if v30 != 0 {
						return
					} else {
						if v29 == int32(0) {
							F__serverAssertWithInfo(m, l0, v9, int32(_a_F_blmoveGenericCommand_0), int32(_a_F_blmoveGenericCommand_1), int32(1252))
							mBase = m.M
							v39 = m.ExcPending
							if v39 != 0 {
								return
							} else {
								F_abort(m)
								mBase = m.M
								base.Wasm_trap_unreachable()
								for {
								}
							}
						} else {
							F_lmoveGenericCommand(m, l0, l1, l2)
							mBase = m.M
							v34 = m.ExcPending
							if v34 != 0 {
								return
							} else {
								return
							}
						}
					}
				} else {
					v14 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+205)))
					if v14&int32(16) == int32(0) {
						v21 = int32(1)
						v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
						F_blockForKeys(m, l0, v21, v22+int32(4), v21, l3, int32(0))
						mBase = m.M
						v28 = m.ExcPending
						if v28 != 0 {
							return
						} else {
							return
						}
					} else {
						F_addReplyNull(m, l0)
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
}
func F_blmpopGetKeys(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
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
func F_body(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
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
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
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
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v62 int32
	_ = v62
	var v66 int32
	_ = v66
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v104 int32
	_ = v104
	v7 = m.G0
	v9 = v7 - int32(592)
	m.G0 = v9
	F_open_func(m, l0, v9+int32(20))
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		return
	} else {
		v15 = *(*int32)(unsafe.Add(mBase, uint32(v9)+20))
		*(*int32)(unsafe.Add(mBase, uint32(v15)+60)) = l3
		v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
		if v17 == int32(40) {
			F_luaX_next(m, l0)
			mBase = m.M
			v36 = m.ExcPending
			if v36 != 0 {
				return
			} else {
				if l2 == int32(0) {
					F_parlist(m, l0)
					mBase = m.M
					v71 = m.ExcPending
					if v71 != 0 {
						return
					} else {
						v72 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
						if v72 == int32(41) {
							F_luaX_next(m, l0)
							mBase = m.M
							v89 = m.ExcPending
							if v89 != 0 {
								return
							} else {
								F_chunk(m, l0)
								mBase = m.M
								v91 = m.ExcPending
								if v91 != 0 {
									return
								} else {
									v92 = *(*int32)(unsafe.Add(mBase, uint32(v9)+20))
									v93 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
									*(*int32)(unsafe.Add(mBase, uint32(v92)+64)) = v93
									F_check_match(m, l0, int32(262), int32(265), l3)
									mBase = m.M
									v98 = m.ExcPending
									if v98 != 0 {
										return
									} else {
										F_close_func(m, l0)
										mBase = m.M
										v100 = m.ExcPending
										if v100 != 0 {
											return
										} else {
											F_pushclosure(m, l0, v9+int32(20), l1)
											mBase = m.M
											v104 = m.ExcPending
											if v104 != 0 {
												return
											} else {
												m.G0 = v9 + int32(592)
												return
											}
										}
									}
								}
							}
						} else {
							v75 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
							v77 = F_luaX_token2str(m, l0, int32(41))
							mBase = m.M
							v78 = m.ExcPending
							if v78 != 0 {
								return
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v9))) = v77
								v80 = m.G3
								v83 = F_luaO_pushfstring(m, v75, v80+int32(_a_F_body_0), v9)
								mBase = m.M
								v84 = m.ExcPending
								if v84 != 0 {
									return
								} else {
									F_luaX_syntaxerror(m, l0, v83)
									mBase = m.M
									v86 = m.ExcPending
									if v86 != 0 {
										return
									} else {
										F_luaX_next(m, l0)
										mBase = m.M
										v89 = m.ExcPending
										if v89 != 0 {
											return
										} else {
											F_chunk(m, l0)
											mBase = m.M
											v91 = m.ExcPending
											if v91 != 0 {
												return
											} else {
												v92 = *(*int32)(unsafe.Add(mBase, uint32(v9)+20))
												v93 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
												*(*int32)(unsafe.Add(mBase, uint32(v92)+64)) = v93
												F_check_match(m, l0, int32(262), int32(265), l3)
												mBase = m.M
												v98 = m.ExcPending
												if v98 != 0 {
													return
												} else {
													F_close_func(m, l0)
													mBase = m.M
													v100 = m.ExcPending
													if v100 != 0 {
														return
													} else {
														F_pushclosure(m, l0, v9+int32(20), l1)
														mBase = m.M
														v104 = m.ExcPending
														if v104 != 0 {
															return
														} else {
															m.G0 = v9 + int32(592)
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
				} else {
					v39 = m.G3
					v43 = F_luaX_newstring(m, l0, v39+int32(_a_F_body_1), int32(4))
					mBase = m.M
					v44 = m.ExcPending
					if v44 != 0 {
						return
					} else {
						F_new_localvar(m, l0, v43, int32(0))
						mBase = m.M
						v47 = m.ExcPending
						if v47 != 0 {
							return
						} else {
							v48 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
							v49 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v48)+50)))
							v50 = int32(1)
							v51 = v49 + v50
							*(*uint8)(unsafe.Add(mBase, uint32(v48)+50)) = uint8(v51)
							v53 = *(*int32)(unsafe.Add(mBase, uint32(v48)))
							v54 = *(*int32)(unsafe.Add(mBase, uint32(v53)+24))
							v62 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v48+v51&int32(255)<<(uint(v50)%32)+int32(170)))))
							v66 = *(*int32)(unsafe.Add(mBase, uint32(v48)+24))
							*(*int32)(unsafe.Add(mBase, uint32(v54+v62*int32(12))+4)) = v66
							F_parlist(m, l0)
							mBase = m.M
							v71 = m.ExcPending
							if v71 != 0 {
								return
							} else {
								v72 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
								if v72 == int32(41) {
									F_luaX_next(m, l0)
									mBase = m.M
									v89 = m.ExcPending
									if v89 != 0 {
										return
									} else {
										F_chunk(m, l0)
										mBase = m.M
										v91 = m.ExcPending
										if v91 != 0 {
											return
										} else {
											v92 = *(*int32)(unsafe.Add(mBase, uint32(v9)+20))
											v93 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
											*(*int32)(unsafe.Add(mBase, uint32(v92)+64)) = v93
											F_check_match(m, l0, int32(262), int32(265), l3)
											mBase = m.M
											v98 = m.ExcPending
											if v98 != 0 {
												return
											} else {
												F_close_func(m, l0)
												mBase = m.M
												v100 = m.ExcPending
												if v100 != 0 {
													return
												} else {
													F_pushclosure(m, l0, v9+int32(20), l1)
													mBase = m.M
													v104 = m.ExcPending
													if v104 != 0 {
														return
													} else {
														m.G0 = v9 + int32(592)
														return
													}
												}
											}
										}
									}
								} else {
									v75 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
									v77 = F_luaX_token2str(m, l0, int32(41))
									mBase = m.M
									v78 = m.ExcPending
									if v78 != 0 {
										return
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v9))) = v77
										v80 = m.G3
										v83 = F_luaO_pushfstring(m, v75, v80+int32(_a_F_body_0), v9)
										mBase = m.M
										v84 = m.ExcPending
										if v84 != 0 {
											return
										} else {
											F_luaX_syntaxerror(m, l0, v83)
											mBase = m.M
											v86 = m.ExcPending
											if v86 != 0 {
												return
											} else {
												F_luaX_next(m, l0)
												mBase = m.M
												v89 = m.ExcPending
												if v89 != 0 {
													return
												} else {
													F_chunk(m, l0)
													mBase = m.M
													v91 = m.ExcPending
													if v91 != 0 {
														return
													} else {
														v92 = *(*int32)(unsafe.Add(mBase, uint32(v9)+20))
														v93 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
														*(*int32)(unsafe.Add(mBase, uint32(v92)+64)) = v93
														F_check_match(m, l0, int32(262), int32(265), l3)
														mBase = m.M
														v98 = m.ExcPending
														if v98 != 0 {
															return
														} else {
															F_close_func(m, l0)
															mBase = m.M
															v100 = m.ExcPending
															if v100 != 0 {
																return
															} else {
																F_pushclosure(m, l0, v9+int32(20), l1)
																mBase = m.M
																v104 = m.ExcPending
																if v104 != 0 {
																	return
																} else {
																	m.G0 = v9 + int32(592)
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
		} else {
			v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
			v22 = F_luaX_token2str(m, l0, int32(40))
			mBase = m.M
			v23 = m.ExcPending
			if v23 != 0 {
				return
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = v22
				v25 = m.G3
				v30 = F_luaO_pushfstring(m, v20, v25+int32(_a_F_body_0), v9+int32(16))
				mBase = m.M
				v31 = m.ExcPending
				if v31 != 0 {
					return
				} else {
					F_luaX_syntaxerror(m, l0, v30)
					mBase = m.M
					v33 = m.ExcPending
					if v33 != 0 {
						return
					} else {
						F_luaX_next(m, l0)
						mBase = m.M
						v36 = m.ExcPending
						if v36 != 0 {
							return
						} else {
							if l2 == int32(0) {
								F_parlist(m, l0)
								mBase = m.M
								v71 = m.ExcPending
								if v71 != 0 {
									return
								} else {
									v72 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
									if v72 == int32(41) {
										F_luaX_next(m, l0)
										mBase = m.M
										v89 = m.ExcPending
										if v89 != 0 {
											return
										} else {
											F_chunk(m, l0)
											mBase = m.M
											v91 = m.ExcPending
											if v91 != 0 {
												return
											} else {
												v92 = *(*int32)(unsafe.Add(mBase, uint32(v9)+20))
												v93 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
												*(*int32)(unsafe.Add(mBase, uint32(v92)+64)) = v93
												F_check_match(m, l0, int32(262), int32(265), l3)
												mBase = m.M
												v98 = m.ExcPending
												if v98 != 0 {
													return
												} else {
													F_close_func(m, l0)
													mBase = m.M
													v100 = m.ExcPending
													if v100 != 0 {
														return
													} else {
														F_pushclosure(m, l0, v9+int32(20), l1)
														mBase = m.M
														v104 = m.ExcPending
														if v104 != 0 {
															return
														} else {
															m.G0 = v9 + int32(592)
															return
														}
													}
												}
											}
										}
									} else {
										v75 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
										v77 = F_luaX_token2str(m, l0, int32(41))
										mBase = m.M
										v78 = m.ExcPending
										if v78 != 0 {
											return
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(v9))) = v77
											v80 = m.G3
											v83 = F_luaO_pushfstring(m, v75, v80+int32(_a_F_body_0), v9)
											mBase = m.M
											v84 = m.ExcPending
											if v84 != 0 {
												return
											} else {
												F_luaX_syntaxerror(m, l0, v83)
												mBase = m.M
												v86 = m.ExcPending
												if v86 != 0 {
													return
												} else {
													F_luaX_next(m, l0)
													mBase = m.M
													v89 = m.ExcPending
													if v89 != 0 {
														return
													} else {
														F_chunk(m, l0)
														mBase = m.M
														v91 = m.ExcPending
														if v91 != 0 {
															return
														} else {
															v92 = *(*int32)(unsafe.Add(mBase, uint32(v9)+20))
															v93 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
															*(*int32)(unsafe.Add(mBase, uint32(v92)+64)) = v93
															F_check_match(m, l0, int32(262), int32(265), l3)
															mBase = m.M
															v98 = m.ExcPending
															if v98 != 0 {
																return
															} else {
																F_close_func(m, l0)
																mBase = m.M
																v100 = m.ExcPending
																if v100 != 0 {
																	return
																} else {
																	F_pushclosure(m, l0, v9+int32(20), l1)
																	mBase = m.M
																	v104 = m.ExcPending
																	if v104 != 0 {
																		return
																	} else {
																		m.G0 = v9 + int32(592)
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
							} else {
								v39 = m.G3
								v43 = F_luaX_newstring(m, l0, v39+int32(_a_F_body_1), int32(4))
								mBase = m.M
								v44 = m.ExcPending
								if v44 != 0 {
									return
								} else {
									F_new_localvar(m, l0, v43, int32(0))
									mBase = m.M
									v47 = m.ExcPending
									if v47 != 0 {
										return
									} else {
										v48 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
										v49 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v48)+50)))
										v50 = int32(1)
										v51 = v49 + v50
										*(*uint8)(unsafe.Add(mBase, uint32(v48)+50)) = uint8(v51)
										v53 = *(*int32)(unsafe.Add(mBase, uint32(v48)))
										v54 = *(*int32)(unsafe.Add(mBase, uint32(v53)+24))
										v62 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v48+v51&int32(255)<<(uint(v50)%32)+int32(170)))))
										v66 = *(*int32)(unsafe.Add(mBase, uint32(v48)+24))
										*(*int32)(unsafe.Add(mBase, uint32(v54+v62*int32(12))+4)) = v66
										F_parlist(m, l0)
										mBase = m.M
										v71 = m.ExcPending
										if v71 != 0 {
											return
										} else {
											v72 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
											if v72 == int32(41) {
												F_luaX_next(m, l0)
												mBase = m.M
												v89 = m.ExcPending
												if v89 != 0 {
													return
												} else {
													F_chunk(m, l0)
													mBase = m.M
													v91 = m.ExcPending
													if v91 != 0 {
														return
													} else {
														v92 = *(*int32)(unsafe.Add(mBase, uint32(v9)+20))
														v93 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
														*(*int32)(unsafe.Add(mBase, uint32(v92)+64)) = v93
														F_check_match(m, l0, int32(262), int32(265), l3)
														mBase = m.M
														v98 = m.ExcPending
														if v98 != 0 {
															return
														} else {
															F_close_func(m, l0)
															mBase = m.M
															v100 = m.ExcPending
															if v100 != 0 {
																return
															} else {
																F_pushclosure(m, l0, v9+int32(20), l1)
																mBase = m.M
																v104 = m.ExcPending
																if v104 != 0 {
																	return
																} else {
																	m.G0 = v9 + int32(592)
																	return
																}
															}
														}
													}
												}
											} else {
												v75 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
												v77 = F_luaX_token2str(m, l0, int32(41))
												mBase = m.M
												v78 = m.ExcPending
												if v78 != 0 {
													return
												} else {
													*(*int32)(unsafe.Add(mBase, uint32(v9))) = v77
													v80 = m.G3
													v83 = F_luaO_pushfstring(m, v75, v80+int32(_a_F_body_0), v9)
													mBase = m.M
													v84 = m.ExcPending
													if v84 != 0 {
														return
													} else {
														F_luaX_syntaxerror(m, l0, v83)
														mBase = m.M
														v86 = m.ExcPending
														if v86 != 0 {
															return
														} else {
															F_luaX_next(m, l0)
															mBase = m.M
															v89 = m.ExcPending
															if v89 != 0 {
																return
															} else {
																F_chunk(m, l0)
																mBase = m.M
																v91 = m.ExcPending
																if v91 != 0 {
																	return
																} else {
																	v92 = *(*int32)(unsafe.Add(mBase, uint32(v9)+20))
																	v93 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
																	*(*int32)(unsafe.Add(mBase, uint32(v92)+64)) = v93
																	F_check_match(m, l0, int32(262), int32(265), l3)
																	mBase = m.M
																	v98 = m.ExcPending
																	if v98 != 0 {
																		return
																	} else {
																		F_close_func(m, l0)
																		mBase = m.M
																		v100 = m.ExcPending
																		if v100 != 0 {
																			return
																		} else {
																			F_pushclosure(m, l0, v9+int32(20), l1)
																			mBase = m.M
																			v104 = m.ExcPending
																			if v104 != 0 {
																				return
																			} else {
																				m.G0 = v9 + int32(592)
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
func F_boolCallback(m *base.Module, l0 int32, l1 int32) {
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
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v47 int32
	_ = v47
	var v52 int32
	_ = v52
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v7 < int32(0) {
		v31 = v6
	} else {
		v12 = l0 + v7<<(uint(int32(2))%32)
		v13 = *(*int32)(unsafe.Add(mBase, uint32(v12)+16))
		v14 = int32(1)
		v15 = v13 + v14
		*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = v15
		v17 = *(*int32)(unsafe.Add(mBase, uint32(v12)+1040))
		if v17 != v14 {
			v31 = v6
		} else {
			v22 = *(*int32)(unsafe.Add(mBase, uint32(v6)+8))
			*(*int32)(unsafe.Add(mBase, uint32(v22)+8)) = int32(3)
			*(*float64)(unsafe.Add(mBase, uint32(v22))) = base.F64_convert_i32_u(v15)
			v26 = *(*int32)(unsafe.Add(mBase, uint32(v6)+8))
			*(*int32)(unsafe.Add(mBase, uint32(v6)+8)) = v26 + int32(16)
			v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			v31 = v30
		}
	}
	v35 = F_lua_checkstack(m, v31, int32(1))
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
		return
	} else {
		if v35 != 0 {
			v41 = *(*int32)(unsafe.Add(mBase, uint32(v31)+8))
			*(*int32)(unsafe.Add(mBase, uint32(v41)+8)) = int32(1)
			*(*int32)(unsafe.Add(mBase, uint32(v41))) = base.B2i32(l1 != int32(0))
			v47 = *(*int32)(unsafe.Add(mBase, uint32(v31)+8))
			*(*int32)(unsafe.Add(mBase, uint32(v31)+8)) = v47 + int32(16)
			F_processCollectionElementEnd(m, l0)
			mBase = m.M
			v52 = m.ExcPending
			if v52 != 0 {
				return
			} else {
				return
			}
		} else {
			F__serverPanic_2(m, int32(1082))
			mBase = m.M
			v39 = m.ExcPending
			if v39 != 0 {
				return
			} else {
				base.Wasm_trap_unreachable()
				for {
				}
			}
		}
	}
}
func F_boolConfigRewrite(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	v11 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+9)))
	if v11&int32(1) == int32(0) {
		v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
		v20 = *(*int32)(unsafe.Add(mBase, uint32(v19)))
		v21 = v20
		v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
		v23 = F_sdsempty(m)
		mBase = m.M
		v24 = m.ExcPending
		if v24 != 0 {
			return
		} else {
			if v21 != 0 {
				v27 = int32(_a_F_boolConfigRewrite_0)
			} else {
				v27 = int32(_a_F_boolConfigRewrite_1)
			}
			*(*int32)(unsafe.Add(mBase, uint32(v9)+4)) = v27
			*(*int32)(unsafe.Add(mBase, uint32(v9))) = l1
			v31 = F_sdscatprintf(m, v23, int32(_a_F_boolConfigRewrite_2), v9)
			mBase = m.M
			v32 = m.ExcPending
			if v32 != 0 {
				return
			} else {
				v34 = F_rewriteConfigRewriteLine(m, l2, l1, v31, base.B2i32(v21 != v22))
				mBase = m.M
				v35 = m.ExcPending
				if v35 != 0 {
					return
				} else {
					m.G0 = v9 + int32(16)
					return
				}
			}
		}
	} else {
		v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+84))
		v17 = F_getModuleBoolConfig(m, v16)
		mBase = m.M
		v18 = m.ExcPending
		if v18 != 0 {
			return
		} else {
			v21 = v17
			v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
			v23 = F_sdsempty(m)
			mBase = m.M
			v24 = m.ExcPending
			if v24 != 0 {
				return
			} else {
				if v21 != 0 {
					v27 = int32(_a_F_boolConfigRewrite_0)
				} else {
					v27 = int32(_a_F_boolConfigRewrite_1)
				}
				*(*int32)(unsafe.Add(mBase, uint32(v9)+4)) = v27
				*(*int32)(unsafe.Add(mBase, uint32(v9))) = l1
				v31 = F_sdscatprintf(m, v23, int32(_a_F_boolConfigRewrite_2), v9)
				mBase = m.M
				v32 = m.ExcPending
				if v32 != 0 {
					return
				} else {
					v34 = F_rewriteConfigRewriteLine(m, l2, l1, v31, base.B2i32(v21 != v22))
					mBase = m.M
					v35 = m.ExcPending
					if v35 != 0 {
						return
					} else {
						m.G0 = v9 + int32(16)
						return
					}
				}
			}
		}
	}
}
func F_brpoplpushCommand(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v16 int64
	_ = v16
	var v18 int32
	_ = v18
	v3 = m.G0
	v5 = v3 - int32(16)
	m.G0 = v5
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(v7)+12))
	v12 = F_getTimeoutFromObjectOrReply(m, l0, v8, v5+int32(8), int32(0))
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return
	} else {
		if v12 != 0 {
			m.G0 = v5 + int32(16)
			return
		} else {
			v16 = *(*int64)(unsafe.Add(mBase, uint32(v5)+8))
			F_blmoveGenericCommand(m, l0, int32(1), int32(0), v16)
			mBase = m.M
			v18 = m.ExcPending
			if v18 != 0 {
				return
			} else {
				m.G0 = v5 + int32(16)
				return
			}
		}
	}
}
func F_bufferReplData(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v39 int32
	_ = v39
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
	var v50 int32
	_ = v50
	var v54 int32
	_ = v54
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v66 int64
	_ = v66
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v78 int64
	_ = v78
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v92 int64
	_ = v92
	var v96 int64
	_ = v96
	var v99 int32
	_ = v99
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
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
	var v121 int32
	_ = v121
	var v123 int32
	_ = v123
	var v125 int32
	_ = v125
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v145 int32
	_ = v145
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v157 int32
	_ = v157
	var v161 int32
	_ = v161
	var v165 int32
	_ = v165
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v177 int32
	_ = v177
	var v182 int32
	_ = v182
	var v189 int32
	_ = v189
	var v191 int32
	_ = v191
	var v194 int64
	_ = v194
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v202 int32
	_ = v202
	var v205 int32
	_ = v205
	var v207 int64
	_ = v207
	var v212 int32
	_ = v212
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	v9 = m.G0
	v11 = v9 - int32(16)
	m.G0 = v11
	v17 = int32(0)
	v18 = int32(16384)
	goto L2
L1:
	;
	m.G0 = v11 + int32(16)
	return
L2:
	;
	if v18 == int32(0) {
		goto L1
	} else {
		goto L4
	}
L3:
	;
	goto L1
L4:
	;
	v26 = *(*int32)(unsafe.Add(mBase, _c_F_bufferReplData[0]))
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v26)+4))
	if v27 == int32(0) {
		v84 = v17
		v85 = v18
		goto L5
	} else {
		goto L6
	}
L5:
	;
	if v84 != 0 {
		v214 = v84
		v215 = v85
		goto L25
	} else {
		goto L26
	}
L6:
	;
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v27)+8))
	if v30 == int32(0) {
		v84 = v17
		v85 = v18
		goto L5
	} else {
		goto L7
	}
L7:
	;
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v30)))
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v30)+4))
	if base.Ui32(v33) <= base.Ui32(v34) {
		v84 = v17
		v85 = v18
		goto L5
	} else {
		goto L8
	}
L8:
	;
	v39 = v33 - v34
	if base.Ui32(v18) < base.Ui32(v39) {
		goto L10
	} else {
		goto L11
	}
L9:
	;
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v30)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v30)+4)) = v73 + v44
	v76 = int32(_a_F_bufferReplData_0)
	v78 = *(*int64)(unsafe.Add(mBase, _c_F_bufferReplData[1]))
	*(*int64)(unsafe.Add(mBase, _c_F_bufferReplData[1])) = v78 + int64(1)
	v84 = v41 - v44
	v85 = v18 - v41
	goto L5
L10:
	;
	v41 = v18
	goto L12
L11:
	;
	v41 = v39
	goto L12
L12:
	;
	v42 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v42)+76))
	v44 = m.T0[v43].(func(*base.Module, int32, int32, int32) int32)(m, l0, v30+v34+int32(8), v41)
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	return
L14:
	;
	if int32(0) < v44 {
		goto L9
	} else {
		goto L15
	}
L15:
	;
	if v44 == int32(0) {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	v54 = *(*int32)(unsafe.Add(mBase, _c_F_bufferReplData[2]))
	if int32(3) < v54 {
		goto L19
	} else {
		goto L20
	}
L17:
	;
	v50 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v50 == int32(3) {
		goto L1
	} else {
		goto L18
	}
L18:
	;
	goto L16
L19:
	;
	v63 = *(*int32)(unsafe.Add(mBase, _c_F_bufferReplData[3]))
	if v63 == int32(0) {
		goto L22
	} else {
		goto L23
	}
L20:
	;
	F__serverLog(m, int32(3), int32(_a_F_bufferReplData_1), int32(0))
	mBase = m.M
	v61 = m.ExcPending
	if v61 != 0 {
		goto L13
	} else {
		goto L21
	}
L21:
	;
	goto L19
L22:
	;
	v71 = F_cancelReplicationHandshake(m, int32(1))
	mBase = m.M
	v72 = m.ExcPending
	if v72 != 0 {
		goto L13
	} else {
		goto L24
	}
L23:
	;
	v66 = *(*int64)(unsafe.Add(mBase, uint32(v63)+32))
	*(*int64)(unsafe.Add(mBase, uint32(v63)+32)) = v66 | int64(4)
	goto L22
L24:
	;
	goto L1
L25:
	;
	if int32(0) < v214 {
		goto L1
	} else {
		goto L64
	}
L26:
	;
	if v85 == int32(0) {
		v214 = v84
		v215 = v85
		goto L25
	} else {
		goto L27
	}
L27:
	;
	v92 = *(*int64)(unsafe.Add(mBase, _c_F_bufferReplData[4]))
	if v92 == int64(0) {
		goto L28
	} else {
		goto L29
	}
L28:
	;
	v112 = int32(16384)
	if base.Ui32(v112) < base.Ui32(v85) {
		goto L35
	} else {
		goto L36
	}
L29:
	;
	v96 = int64(*(*uint32)(unsafe.Add(mBase, _c_F_bufferReplData[5])))
	if base.Ui64(v96) <= base.Ui64(v92) {
		goto L28
	} else {
		goto L30
	}
L30:
	;
	v99 = *(*int32)(unsafe.Add(mBase, _c_F_bufferReplData[2]))
	if int32(2) < v99 {
		goto L31
	} else {
		goto L32
	}
L31:
	;
	v108 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v109 = *(*int32)(unsafe.Add(mBase, uint32(v108)+84))
	v110 = m.T0[v109].(func(*base.Module, int32, int32) int32)(m, l0, int32(0))
	mBase = m.M
	v111 = m.ExcPending
	if v111 != 0 {
		goto L13
	} else {
		goto L34
	}
L32:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v11))) = v92
	F__serverLog(m, int32(2), int32(_a_F_bufferReplData_2), v11)
	mBase = m.M
	v106 = m.ExcPending
	if v106 != 0 {
		goto L13
	} else {
		goto L33
	}
L33:
	;
	goto L31
L34:
	;
	goto L1
L35:
	;
	v115 = v85
	goto L37
L36:
	;
	v115 = v112
	goto L37
L37:
	;
	v117 = *(*int32)(unsafe.Add(mBase, _c_F_bufferReplData[6]))
	v118 = int32(262159)
	if base.Ui32(v118) < base.Ui32(v117) {
		goto L38
	} else {
		goto L39
	}
L38:
	;
	v121 = v117
	goto L40
L39:
	;
	v121 = v118
	goto L40
L40:
	;
	v123 = int32(base.Ui32(v121) >> (uint(int32(4)) % 32))
	if base.Ui32(v115) < base.Ui32(v123) {
		goto L41
	} else {
		goto L42
	}
L41:
	;
	v125 = v115
	goto L43
L42:
	;
	v125 = v123
	goto L43
L43:
	;
	v130 = F_zmalloc_usable(m, v125+int32(8), v11+int32(12))
	mBase = m.M
	v131 = m.ExcPending
	if v131 != 0 {
		goto L13
	} else {
		goto L44
	}
L44:
	;
	v132 = *(*int32)(unsafe.Add(mBase, uint32(v11)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v130)+4)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v130))) = v132 + int32(-8)
	v139 = *(*int32)(unsafe.Add(mBase, _c_F_bufferReplData[0]))
	v140 = F_listAddNodeTail(m, v139, v130)
	mBase = m.M
	v141 = m.ExcPending
	if v141 != 0 {
		goto L13
	} else {
		goto L45
	}
L45:
	;
	v142 = int32(_a_F_bufferReplData_0)
	v143 = *(*int32)(unsafe.Add(mBase, uint32(v11)+12))
	v145 = *(*int32)(unsafe.Add(mBase, _c_F_bufferReplData[7]))
	*(*int32)(unsafe.Add(mBase, _c_F_bufferReplData[7])) = v143 + v145 + int32(12)
	v152 = *(*int32)(unsafe.Add(mBase, _c_F_bufferReplData[5]))
	v153 = *(*int32)(unsafe.Add(mBase, uint32(v130)))
	v154 = v152 + v153
	*(*int32)(unsafe.Add(mBase, _c_F_bufferReplData[5])) = v154
	v157 = *(*int32)(unsafe.Add(mBase, _c_F_bufferReplData[8]))
	if base.Ui32(v154) <= base.Ui32(v157) {
		goto L46
	} else {
		goto L47
	}
L46:
	;
	v161 = *(*int32)(unsafe.Add(mBase, uint32(v130)+4))
	v165 = *(*int32)(unsafe.Add(mBase, uint32(v130)))
	if base.Ui32(v85) < base.Ui32(v165) {
		goto L50
	} else {
		goto L51
	}
L47:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_bufferReplData[8])) = v154
	goto L46
L48:
	;
	v214 = v212
	v215 = v85 - v167
	goto L25
L49:
	;
	v202 = *(*int32)(unsafe.Add(mBase, uint32(v130)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v130)+4)) = v202 + v170
	v205 = int32(_a_F_bufferReplData_0)
	v207 = *(*int64)(unsafe.Add(mBase, _c_F_bufferReplData[1]))
	*(*int64)(unsafe.Add(mBase, _c_F_bufferReplData[1])) = v207 + int64(1)
	v212 = v167 - v170
	goto L48
L50:
	;
	v167 = v85
	goto L52
L51:
	;
	v167 = v165
	goto L52
L52:
	;
	v168 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v169 = *(*int32)(unsafe.Add(mBase, uint32(v168)+76))
	v170 = m.T0[v169].(func(*base.Module, int32, int32, int32) int32)(m, l0, v130+v161+int32(8), v167)
	mBase = m.M
	v171 = m.ExcPending
	if v171 != 0 {
		goto L13
	} else {
		goto L53
	}
L53:
	;
	if int32(0) < v170 {
		goto L49
	} else {
		goto L54
	}
L54:
	;
	if v170 == int32(0) {
		goto L55
	} else {
		goto L56
	}
L55:
	;
	v182 = *(*int32)(unsafe.Add(mBase, _c_F_bufferReplData[2]))
	if int32(3) < v182 {
		goto L58
	} else {
		goto L59
	}
L56:
	;
	v177 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v177 == int32(3) {
		v212 = int32(-1)
		goto L48
	} else {
		goto L57
	}
L57:
	;
	goto L55
L58:
	;
	v191 = *(*int32)(unsafe.Add(mBase, _c_F_bufferReplData[3]))
	if v191 == int32(0) {
		goto L61
	} else {
		goto L62
	}
L59:
	;
	F__serverLog(m, int32(3), int32(_a_F_bufferReplData_1), int32(0))
	mBase = m.M
	v189 = m.ExcPending
	if v189 != 0 {
		goto L13
	} else {
		goto L60
	}
L60:
	;
	goto L58
L61:
	;
	v199 = F_cancelReplicationHandshake(m, int32(1))
	mBase = m.M
	v200 = m.ExcPending
	if v200 != 0 {
		goto L13
	} else {
		goto L63
	}
L62:
	;
	v194 = *(*int64)(unsafe.Add(mBase, uint32(v191)+32))
	*(*int64)(unsafe.Add(mBase, uint32(v191)+32)) = v194 | int64(4)
	goto L61
L63:
	;
	v212 = int32(-1)
	goto L48
L64:
	;
	if v214 != int32(-1) {
		v17 = v214
		v18 = v215
		goto L2
	} else {
		goto L65
	}
L65:
	;
	goto L3
}
func F_bytesToHuman(m *base.Module, l0 int32, l1 int32, l2 int64) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	v5 = m.G0
	v7 = v5 - int32(112)
	m.G0 = v7
	if base.Ui64(int64(1023)) < base.Ui64(l2) {
		if base.Ui64(int64(1048575)) < base.Ui64(l2) {
			if base.Ui64(int64(1073741823)) < base.Ui64(l2) {
				if base.Ui64(int64(1099511627775)) < base.Ui64(l2) {
					if base.Ui64(int64(1125899906842623)) < base.Ui64(l2) {
						if base.Ui64(int64(1152921504606846975)) < base.Ui64(l2) {
							*(*int64)(unsafe.Add(mBase, uint32(v7)+96)) = l2
							v74 = F_snprintf(m, l0, l1, int32(_a_F_bytesToHuman_0), v7+int32(96))
							mBase = m.M
							v75 = m.ExcPending
							if v75 != 0 {
								return
							} else {
								m.G0 = v7 + int32(112)
								return
							}
						} else {
							*(*float64)(unsafe.Add(mBase, uint32(v7)+80)) = base.F64_mul(base.F64_convert_i64_u(l2), float64(8.881784197001252e-16))
							v68 = F_snprintf(m, l0, l1, int32(_a_F_bytesToHuman_1), v7+int32(80))
							mBase = m.M
							v69 = m.ExcPending
							if v69 != 0 {
								return
							} else {
								m.G0 = v7 + int32(112)
								return
							}
						}
					} else {
						*(*float64)(unsafe.Add(mBase, uint32(v7)+64)) = base.F64_mul(base.F64_convert_i64_u(l2), float64(9.094947017729282e-13))
						v57 = F_snprintf(m, l0, l1, int32(_a_F_bytesToHuman_2), v7+int32(64))
						mBase = m.M
						v58 = m.ExcPending
						if v58 != 0 {
							return
						} else {
							m.G0 = v7 + int32(112)
							return
						}
					}
				} else {
					*(*float64)(unsafe.Add(mBase, uint32(v7)+48)) = base.F64_mul(base.F64_convert_i64_u(l2), float64(9.313225746154785e-10))
					v46 = F_snprintf(m, l0, l1, int32(_a_F_bytesToHuman_3), v7+int32(48))
					mBase = m.M
					v47 = m.ExcPending
					if v47 != 0 {
						return
					} else {
						m.G0 = v7 + int32(112)
						return
					}
				}
			} else {
				*(*float64)(unsafe.Add(mBase, uint32(v7)+32)) = base.F64_mul(base.F64_convert_i64_u(l2), float64(9.5367431640625e-07))
				v35 = F_snprintf(m, l0, l1, int32(_a_F_bytesToHuman_4), v7+int32(32))
				mBase = m.M
				v36 = m.ExcPending
				if v36 != 0 {
					return
				} else {
					m.G0 = v7 + int32(112)
					return
				}
			}
		} else {
			*(*float64)(unsafe.Add(mBase, uint32(v7)+16)) = base.F64_mul(base.F64_convert_i64_u(l2), float64(0.0009765625))
			v24 = F_snprintf(m, l0, l1, int32(_a_F_bytesToHuman_5), v7+int32(16))
			mBase = m.M
			v25 = m.ExcPending
			if v25 != 0 {
				return
			} else {
				m.G0 = v7 + int32(112)
				return
			}
		}
	} else {
		*(*int64)(unsafe.Add(mBase, uint32(v7))) = l2
		v13 = F_snprintf(m, l0, l1, int32(_a_F_bytesToHuman_0), v7)
		mBase = m.M
		v14 = m.ExcPending
		if v14 != 0 {
			return
		} else {
			m.G0 = v7 + int32(112)
			return
		}
	}
}
func F_bzmpopCommand(m *base.Module, l0 int32) {
	var v5 int32
	_ = v5
	F_zmpopGenericCommand(m, l0, int32(2), int32(1))
	v5 = m.ExcPending
	if v5 != 0 {
		return
	} else {
		return
	}
}
func F_bzpopminCommand(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v10 int32
	_ = v10
	var v16 int32
	_ = v16
	v2 = int32(0)
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v10 = int32(-1)
	F_blockingGenericZpopCommand(m, l0, v3+int32(4), v6+int32(-2), v2, v6+v10, v10, v2, v2)
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		return
	} else {
		return
	}
}
