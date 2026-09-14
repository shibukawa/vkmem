package p1

import (
	base "github.com/shibukawa/vkmem/internal/aot/vkaot/base"
	"unsafe"
)

func F___isspace_1(m *base.Module, l0 int32) int32 {
	return base.B2i32(l0 == int32(32)) | base.B2i32(base.Ui32(l0+int32(-9)) < base.Ui32(int32(5)))
}
func F___isspace_6(m *base.Module, l0 int32) int32 {
	return base.B2i32(l0 == int32(32)) | base.B2i32(base.Ui32(l0+int32(-9)) < base.Ui32(int32(5)))
}
func F_ignore_err(m *base.Module, l0 int32, l1 int32) int32 {
	return int32(0)
}
func F_incrementalTrimReplicationBacklog(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v18 int64
	_ = v18
	var v20 int64
	_ = v20
	var v29 int32
	_ = v29
	var v30 int64
	_ = v30
	var v31 int64
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v48 int64
	_ = v48
	var v49 int64
	_ = v49
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v63 int64
	_ = v63
	var v64 int64
	_ = v64
	var v66 int64
	_ = v66
	var v68 int64
	_ = v68
	var v71 int64
	_ = v71
	var v73 int64
	_ = v73
	var v75 int64
	_ = v75
	var v77 int64
	_ = v77
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
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
	var v111 int32
	_ = v111
	var v113 int32
	_ = v113
	var v119 int32
	_ = v119
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v124 int64
	_ = v124
	var v126 int64
	_ = v126
	var v131 int32
	_ = v131
	var v133 int32
	_ = v133
	var v140 int64
	_ = v140
	var v143 int32
	_ = v143
	var v144 int64
	_ = v144
	var v151 int64
	_ = v151
	var v163 int32
	_ = v163
	var v169 int32
	_ = v169
	var v175 int32
	_ = v175
	var v182 int32
	_ = v182
	v10 = m.G0
	v12 = v10 - int32(16)
	m.G0 = v12
	v15 = *(*int32)(unsafe.Add(mBase, _c_F_incrementalTrimReplicationBacklog[0]))
	if v15 == int32(0) {
		goto L4
	} else {
		goto L5
	}
L1:
	;
	F__serverAssert(m, int32(_a_F_incrementalTrimReplicationBacklog_0), int32(_a_F_incrementalTrimReplicationBacklog_1), int32(414))
	mBase = m.M
	v182 = m.ExcPending
	if v182 != 0 {
		goto L17
	} else {
		goto L27
	}
L2:
	;
	F__serverAssert(m, int32(_a_F_incrementalTrimReplicationBacklog_2), int32(_a_F_incrementalTrimReplicationBacklog_1), int32(405))
	mBase = m.M
	v175 = m.ExcPending
	if v175 != 0 {
		goto L17
	} else {
		goto L26
	}
L3:
	;
	F__serverAssert(m, int32(_a_F_incrementalTrimReplicationBacklog_3), int32(_a_F_incrementalTrimReplicationBacklog_1), int32(389))
	mBase = m.M
	v169 = m.ExcPending
	if v169 != 0 {
		goto L17
	} else {
		goto L25
	}
L4:
	;
	F__serverAssert(m, int32(_a_F_incrementalTrimReplicationBacklog_4), int32(_a_F_incrementalTrimReplicationBacklog_1), int32(375))
	mBase = m.M
	v163 = m.ExcPending
	if v163 != 0 {
		goto L17
	} else {
		goto L24
	}
L5:
	;
	v18 = *(*int64)(unsafe.Add(mBase, uint32(v15)+16))
	v20 = *(*int64)(unsafe.Add(mBase, _c_F_incrementalTrimReplicationBacklog[1]))
	if v18 <= v20 {
		v143 = v15
		v144 = v18
		goto L6
	} else {
		goto L7
	}
L6:
	;
	v151 = *(*int64)(unsafe.Add(mBase, _c_F_incrementalTrimReplicationBacklog[2]))
	*(*int64)(unsafe.Add(mBase, uint32(v143)+24)) = v151 - v144 + int64(1)
	m.G0 = v12 + int32(16)
	return
L7:
	;
	if l0 == int32(0) {
		v143 = v15
		v144 = v18
		goto L6
	} else {
		goto L8
	}
L8:
	;
	v29 = v15
	v30 = v18
	v31 = v20
	v32 = v15 + int32(16)
	v33 = int32(0)
	goto L10
L9:
	;
	v140 = *(*int64)(unsafe.Add(mBase, uint32(v133)+16))
	v143 = v133
	v144 = v140
	goto L6
L10:
	;
	v37 = *(*int32)(unsafe.Add(mBase, _c_F_incrementalTrimReplicationBacklog[3]))
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v37)+20))
	if base.Ui32(v38) < base.Ui32(int32(2)) {
		v133 = v29
		goto L9
	} else {
		goto L12
	}
L11:
	;
	v133 = v123
	goto L9
L12:
	;
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v37)))
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v29)))
	if v41 != v42 {
		goto L3
	} else {
		goto L13
	}
L13:
	;
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v41)+8))
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v44)))
	if v45 != int32(1) {
		v133 = v29
		goto L9
	} else {
		goto L14
	}
L14:
	;
	v48 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v44)+24)))
	v49 = v30 - v48
	if v49 <= v31 {
		v133 = v29
		goto L9
	} else {
		goto L15
	}
L15:
	;
	v51 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v44))) = v51
	*(*int64)(unsafe.Add(mBase, uint32(v32))) = v49
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v41)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v29))) = v54
	if v54 == v51 {
		goto L2
	} else {
		goto L16
	}
L16:
	;
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v54)+8))
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v58)))
	*(*int32)(unsafe.Add(mBase, uint32(v58))) = v59 + int32(1)
	v63 = *(*int64)(unsafe.Add(mBase, uint32(v44)+16))
	v64 = int64(56)
	v66 = int64(65280)
	v68 = int64(40)
	v71 = int64(16711680)
	v73 = int64(24)
	v75 = int64(4278190080)
	v77 = int64(8)
	*(*int64)(unsafe.Add(mBase, uint32(v12)+8)) = v63<<(uint(v64)%64) | v63&v66<<(uint(v68)%64) | (v63&v71<<(uint(v73)%64) | v63&v75<<(uint(v77)%64)) | (int64(base.Ui64(v63)>>(uint(v77)%64))&v75 | int64(base.Ui64(v63)>>(uint(v73)%64))&v71 | (int64(base.Ui64(v63)>>(uint(v68)%64))&v66 | int64(base.Ui64(v63)>>(uint(v64)%64))))
	v100 = *(*int32)(unsafe.Add(mBase, uint32(v29)+8))
	v101 = int32(8)
	v105 = F_raxRemove(m, v100, v12+v101, v101, int32(0))
	mBase = m.M
	v106 = m.ExcPending
	if v106 != 0 {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	return
L18:
	;
	v107 = *(*int32)(unsafe.Add(mBase, uint32(v44)))
	if v107 != 0 {
		goto L1
	} else {
		goto L19
	}
L19:
	;
	v108 = *(*int32)(unsafe.Add(mBase, uint32(v44)+28))
	v109 = *(*int32)(unsafe.Add(mBase, uint32(v44)+24))
	if v108 != v109 {
		goto L1
	} else {
		goto L20
	}
L20:
	;
	v111 = int32(_a_F_incrementalTrimReplicationBacklog_5)
	v113 = *(*int32)(unsafe.Add(mBase, _c_F_incrementalTrimReplicationBacklog[4]))
	*(*int32)(unsafe.Add(mBase, _c_F_incrementalTrimReplicationBacklog[4])) = v113 - v108 + int32(-44)
	v119 = *(*int32)(unsafe.Add(mBase, _c_F_incrementalTrimReplicationBacklog[3]))
	F_listDelNode(m, v119, v41)
	mBase = m.M
	v121 = m.ExcPending
	if v121 != 0 {
		goto L17
	} else {
		goto L21
	}
L21:
	;
	v122 = int32(_a_F_incrementalTrimReplicationBacklog_5)
	v123 = *(*int32)(unsafe.Add(mBase, _c_F_incrementalTrimReplicationBacklog[0]))
	v124 = *(*int64)(unsafe.Add(mBase, uint32(v123)+16))
	v126 = *(*int64)(unsafe.Add(mBase, _c_F_incrementalTrimReplicationBacklog[1]))
	if v124 <= v126 {
		v133 = v123
		goto L9
	} else {
		goto L22
	}
L22:
	;
	v131 = v33 + int32(1)
	if base.Ui32(v131) < base.Ui32(l0) {
		v29 = v123
		v30 = v124
		v31 = v126
		v32 = v123 + int32(16)
		v33 = v131
		goto L10
	} else {
		goto L23
	}
L23:
	;
	goto L11
L24:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L25:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
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
func F_initSharedQueryBuf(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v1 int32
	_ = v1
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v18 int32
	_ = v18
	var v22 int32
	_ = v22
	var v32 int32
	_ = v32
	v1 = int32(0)
	v5 = F_sdsnewlen(m, v1, int32(16384))
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return
	} else {
		*(*int32)(unsafe.Add(mBase, _c_F_initSharedQueryBuf[0])) = v5
		v10 = v5 + int32(-1)
		v11 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10))))
		switch v11 & int32(7) {
		case 0:
			v14 = int32(0)
			*(*uint8)(unsafe.Add(mBase, uint32(v10))) = uint8(v14)
		case 1:
			v18 = int32(0)
			*(*uint8)(unsafe.Add(mBase, uint32(v5+int32(-3)))) = uint8(v18)
		case 2:
			v22 = int32(0)
			*(*uint16)(unsafe.Add(mBase, uint32(v5+int32(-5)))) = uint16(v22)
		case 3:
			*(*int32)(unsafe.Add(mBase, uint32(v5+int32(-9)))) = int32(0)
		case 4:
			*(*int64)(unsafe.Add(mBase, uint32(v5+int32(-17)))) = int64(0)
		default:
		}
		v32 = int32(0)
		*(*uint8)(unsafe.Add(mBase, uint32(v5))) = uint8(v32)
		return
	}
}
func F_insertToBucket_RAX(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int64) int32 {
	mBase := m.M
	_ = mBase
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v34 int32
	_ = v34
	var v41 int32
	_ = v41
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v58 int32
	_ = v58
	var v72 int32
	_ = v72
	var v80 int64
	_ = v80
	var v82 int64
	_ = v82
	var v85 int64
	_ = v85
	var v87 int64
	_ = v87
	var v89 int64
	_ = v89
	var v91 int64
	_ = v91
	var v93 int64
	_ = v93
	var v117 int64
	_ = v117
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v130 int64
	_ = v130
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
	var v148 int32
	_ = v148
	var v151 int32
	_ = v151
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v158 int32
	_ = v158
	var v163 int32
	_ = v163
	var v172 int32
	_ = v172
	var v175 int32
	_ = v175
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v186 int32
	_ = v186
	var v191 int32
	_ = v191
	var v206 int32
	_ = v206
	var v209 int64
	_ = v209
	var v214 int32
	_ = v214
	var v215 int64
	_ = v215
	var v221 int32
	_ = v221
	var v224 int64
	_ = v224
	var v237 int32
	_ = v237
	var v238 int64
	_ = v238
	var v239 int64
	_ = v239
	var v240 int32
	_ = v240
	var v247 int32
	_ = v247
	var v251 int32
	_ = v251
	var v252 int64
	_ = v252
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v255 int64
	_ = v255
	var v256 int32
	_ = v256
	var v259 int64
	_ = v259
	var v261 int64
	_ = v261
	var v263 int32
	_ = v263
	var v264 int64
	_ = v264
	var v266 int64
	_ = v266
	var v268 int64
	_ = v268
	var v270 int64
	_ = v270
	var v273 int64
	_ = v273
	var v275 int64
	_ = v275
	var v277 int64
	_ = v277
	var v279 int64
	_ = v279
	var v307 int32
	_ = v307
	var v308 int32
	_ = v308
	var v311 int32
	_ = v311
	var v317 int32
	_ = v317
	var v320 int64
	_ = v320
	var v326 int32
	_ = v326
	var v327 int32
	_ = v327
	var v329 int32
	_ = v329
	var v333 int32
	_ = v333
	var v351 int32
	_ = v351
	var v354 int32
	_ = v354
	var v358 int32
	_ = v358
	var v359 int32
	_ = v359
	var v366 int32
	_ = v366
	var v367 int64
	_ = v367
	var v368 int32
	_ = v368
	var v369 int32
	_ = v369
	var v376 int32
	_ = v376
	var v377 int64
	_ = v377
	var v378 int32
	_ = v378
	var v380 int64
	_ = v380
	var v383 int64
	_ = v383
	var v384 int64
	_ = v384
	var v387 int64
	_ = v387
	var v392 int64
	_ = v392
	var v397 int32
	_ = v397
	var v400 int32
	_ = v400
	var v402 int32
	_ = v402
	var v405 int32
	_ = v405
	var v410 int32
	_ = v410
	var v411 int64
	_ = v411
	var v412 int32
	_ = v412
	var v413 int32
	_ = v413
	var v420 int32
	_ = v420
	var v421 int64
	_ = v421
	var v422 int32
	_ = v422
	var v424 int64
	_ = v424
	var v427 int64
	_ = v427
	var v428 int64
	_ = v428
	var v431 int64
	_ = v431
	var v436 int64
	_ = v436
	var v443 int32
	_ = v443
	var v446 int32
	_ = v446
	var v450 int64
	_ = v450
	var v452 int32
	_ = v452
	var v456 int32
	_ = v456
	var v466 int32
	_ = v466
	var v472 int32
	_ = v472
	var v477 int32
	_ = v477
	var v478 int32
	_ = v478
	var v481 int32
	_ = v481
	var v484 int32
	_ = v484
	var v485 int32
	_ = v485
	var v492 int32
	_ = v492
	var v495 int32
	_ = v495
	var v500 int32
	_ = v500
	var v501 int32
	_ = v501
	var v507 int32
	_ = v507
	var v510 int32
	_ = v510
	var v517 int32
	_ = v517
	var v520 int32
	_ = v520
	var v521 int32
	_ = v521
	var v528 int32
	_ = v528
	var v533 int32
	_ = v533
	var v554 int64
	_ = v554
	var v561 int32
	_ = v561
	var v566 int64
	_ = v566
	var v568 int64
	_ = v568
	var v570 int64
	_ = v570
	var v573 int64
	_ = v573
	var v575 int64
	_ = v575
	var v577 int64
	_ = v577
	var v579 int64
	_ = v579
	var v606 int32
	_ = v606
	var v607 int32
	_ = v607
	var v608 int32
	_ = v608
	var v609 int32
	_ = v609
	var v611 int32
	_ = v611
	var v612 int32
	_ = v612
	var v619 int32
	_ = v619
	var v627 int32
	_ = v627
	var v630 int32
	_ = v630
	var v631 int32
	_ = v631
	var v638 int32
	_ = v638
	var v643 int32
	_ = v643
	var v658 int32
	_ = v658
	var v659 int32
	_ = v659
	var v661 int32
	_ = v661
	var v664 int32
	_ = v664
	var v667 int32
	_ = v667
	var v668 int32
	_ = v668
	var v675 int32
	_ = v675
	var v680 int32
	_ = v680
	var v696 int32
	_ = v696
	var v697 int32
	_ = v697
	var v702 int32
	_ = v702
	var v709 int32
	_ = v709
	var v715 int32
	_ = v715
	var v721 int32
	_ = v721
	var v727 int32
	_ = v727
	var v733 int32
	_ = v733
	var v739 int32
	_ = v739
	var v745 int32
	_ = v745
	var v751 int32
	_ = v751
	var v757 int32
	_ = v757
	var v763 int32
	_ = v763
	var v769 int32
	_ = v769
	var v775 int32
	_ = v775
	var v781 int32
	_ = v781
	var v787 int32
	_ = v787
	var v793 int32
	_ = v793
	var v799 int32
	_ = v799
	var v805 int32
	_ = v805
	var v811 int32
	_ = v811
	var v817 int32
	_ = v817
	var v842 int32
	_ = v842
	var v845 int32
	_ = v845
	v20 = m.G0
	v22 = v20 - int32(48)
	m.G0 = v22
	*(*int64)(unsafe.Add(mBase, uint32(v22)+24)) = int64(0)
	switch l1 + int32(1) {
	case 0:
		goto L20
	case 1:
		goto L22
	default:
		goto L21
	}
L1:
	;
	m.G0 = v22 + int32(48)
	return v845
L2:
	;
	F__serverAssert(m, int32(_a_F_insertToBucket_RAX_0), int32(_a_F_insertToBucket_RAX_1), int32(1172))
	mBase = m.M
	v842 = m.ExcPending
	if v842 != 0 {
		goto L23
	} else {
		goto L185
	}
L3:
	;
	F__serverAssert(m, int32(_a_F_insertToBucket_RAX_2), int32(_a_F_insertToBucket_RAX_1), int32(1283))
	mBase = m.M
	v817 = m.ExcPending
	if v817 != 0 {
		goto L23
	} else {
		goto L184
	}
L4:
	;
	F__serverAssert(m, int32(_a_F_insertToBucket_RAX_3), int32(_a_F_insertToBucket_RAX_1), int32(1178))
	mBase = m.M
	v811 = m.ExcPending
	if v811 != 0 {
		goto L23
	} else {
		goto L183
	}
L5:
	;
	F__serverAssert(m, int32(_a_F_insertToBucket_RAX_4), int32(_a_F_insertToBucket_RAX_1), int32(797))
	mBase = m.M
	v805 = m.ExcPending
	if v805 != 0 {
		goto L23
	} else {
		goto L182
	}
L6:
	;
	F__serverAssert(m, int32(_a_F_insertToBucket_RAX_5), int32(_a_F_insertToBucket_RAX_1), int32(1177))
	mBase = m.M
	v799 = m.ExcPending
	if v799 != 0 {
		goto L23
	} else {
		goto L181
	}
L7:
	;
	F__serverAssert(m, int32(_a_F_insertToBucket_RAX_4), int32(_a_F_insertToBucket_RAX_1), int32(797))
	mBase = m.M
	v793 = m.ExcPending
	if v793 != 0 {
		goto L23
	} else {
		goto L180
	}
L8:
	;
	F__serverAssert(m, int32(_a_F_insertToBucket_RAX_6), int32(_a_F_insertToBucket_RAX_1), int32(816))
	mBase = m.M
	v787 = m.ExcPending
	if v787 != 0 {
		goto L23
	} else {
		goto L179
	}
L9:
	;
	F__serverAssert(m, int32(_a_F_insertToBucket_RAX_6), int32(_a_F_insertToBucket_RAX_1), int32(816))
	mBase = m.M
	v781 = m.ExcPending
	if v781 != 0 {
		goto L23
	} else {
		goto L178
	}
L10:
	;
	F__serverAssert(m, int32(_a_F_insertToBucket_RAX_7), int32(_a_F_insertToBucket_RAX_1), int32(1173))
	mBase = m.M
	v775 = m.ExcPending
	if v775 != 0 {
		goto L23
	} else {
		goto L177
	}
L11:
	;
	F__serverAssert(m, int32(_a_F_insertToBucket_RAX_8), int32(_a_F_insertToBucket_RAX_1), int32(600))
	mBase = m.M
	v769 = m.ExcPending
	if v769 != 0 {
		goto L23
	} else {
		goto L176
	}
L12:
	;
	F__serverAssert(m, int32(_a_F_insertToBucket_RAX_8), int32(_a_F_insertToBucket_RAX_1), int32(600))
	mBase = m.M
	v763 = m.ExcPending
	if v763 != 0 {
		goto L23
	} else {
		goto L175
	}
L13:
	;
	F__serverAssert(m, int32(_a_F_insertToBucket_RAX_8), int32(_a_F_insertToBucket_RAX_1), int32(600))
	mBase = m.M
	v757 = m.ExcPending
	if v757 != 0 {
		goto L23
	} else {
		goto L174
	}
L14:
	;
	F__serverAssert(m, int32(_a_F_insertToBucket_RAX_8), int32(_a_F_insertToBucket_RAX_1), int32(600))
	mBase = m.M
	v751 = m.ExcPending
	if v751 != 0 {
		goto L23
	} else {
		goto L173
	}
L15:
	;
	F__serverAssert(m, int32(_a_F_insertToBucket_RAX_9), int32(_a_F_insertToBucket_RAX_1), int32(1165))
	mBase = m.M
	v745 = m.ExcPending
	if v745 != 0 {
		goto L23
	} else {
		goto L172
	}
L16:
	;
	F__serverAssert(m, int32(_a_F_insertToBucket_RAX_10), int32(_a_F_insertToBucket_RAX_1), int32(853))
	mBase = m.M
	v739 = m.ExcPending
	if v739 != 0 {
		goto L23
	} else {
		goto L171
	}
L17:
	;
	F__serverAssert(m, int32(_a_F_insertToBucket_RAX_11), int32(_a_F_insertToBucket_RAX_1), int32(1272))
	mBase = m.M
	v733 = m.ExcPending
	if v733 != 0 {
		goto L23
	} else {
		goto L170
	}
L18:
	;
	F__serverAssert(m, int32(_a_F_insertToBucket_RAX_6), int32(_a_F_insertToBucket_RAX_1), int32(816))
	mBase = m.M
	v727 = m.ExcPending
	if v727 != 0 {
		goto L23
	} else {
		goto L169
	}
L19:
	;
	F__serverAssert(m, int32(_a_F_insertToBucket_RAX_12), int32(_a_F_insertToBucket_RAX_1), int32(1265))
	mBase = m.M
	v721 = m.ExcPending
	if v721 != 0 {
		goto L23
	} else {
		goto L168
	}
L20:
	;
	F__serverAssert(m, int32(_a_F_insertToBucket_RAX_13), int32(_a_F_insertToBucket_RAX_1), int32(807))
	mBase = m.M
	v715 = m.ExcPending
	if v715 != 0 {
		goto L23
	} else {
		goto L167
	}
L21:
	;
	if l1&int32(7) != int32(6) {
		goto L20
	} else {
		goto L25
	}
L22:
	;
	F__serverAssert(m, int32(_a_F_insertToBucket_RAX_14), int32(_a_F_insertToBucket_RAX_1), int32(781))
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
		goto L23
	} else {
		goto L24
	}
L23:
	;
	return int32(0)
L24:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L25:
	;
	v41 = l1 & int32(-8)
	v50 = F_findBucket_2(m, v41, l3, v22+int32(24), v22+int32(20), v22+int32(8), v22+int32(4))
	mBase = m.M
	v51 = m.ExcPending
	if v51 != 0 {
		goto L23
	} else {
		goto L33
	}
L26:
	;
	F__serverPanic_1(m, int32(_a_F_insertToBucket_RAX_1), int32(1301), int32(_a_F_insertToBucket_RAX_15), int32(0))
	mBase = m.M
	v709 = m.ExcPending
	if v709 != 0 {
		goto L23
	} else {
		goto L166
	}
L27:
	;
	v696 = F_hashtableAdd(m, v50&int32(-8), l2)
	mBase = m.M
	v697 = m.ExcPending
	if v697 != 0 {
		goto L23
	} else {
		goto L163
	}
L28:
	;
	v206 = v50 & int32(-8)
	if v206 == int32(0) {
		goto L67
	} else {
		goto L68
	}
L29:
	;
	v126 = F_valkey_malloc(m, int32(16))
	mBase = m.M
	v127 = m.ExcPending
	if v127 != 0 {
		goto L23
	} else {
		goto L44
	}
L30:
	;
	switch l2 + int32(1) {
	case 0:
		goto L19
	case 1:
		goto L37
	default:
		goto L36
	}
L31:
	;
	if v50&int32(1) != 0 {
		goto L29
	} else {
		goto L35
	}
L32:
	;
	F__serverAssert(m, int32(_a_F_insertToBucket_RAX_14), int32(_a_F_insertToBucket_RAX_1), int32(781))
	mBase = m.M
	v58 = m.ExcPending
	if v58 != 0 {
		goto L23
	} else {
		goto L34
	}
L33:
	;
	switch v50 + int32(1) {
	case 0:
		goto L30
	case 1:
		goto L32
	default:
		goto L31
	}
L34:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L35:
	;
	switch v50&int32(6) + int32(-2) {
	case 0:
		goto L28
	default:
		goto L26
	case 2:
		goto L27
	}
L36:
	;
	if l2&int32(1) == int32(0) {
		goto L19
	} else {
		goto L39
	}
L37:
	;
	F__serverAssert(m, int32(_a_F_insertToBucket_RAX_14), int32(_a_F_insertToBucket_RAX_1), int32(781))
	mBase = m.M
	v72 = m.ExcPending
	if v72 != 0 {
		goto L23
	} else {
		goto L38
	}
L38:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L39:
	;
	v80 = l3 & int64(-8192)
	v82 = v80 + int64(8192)
	v85 = int64(40)
	v87 = int64(16711680)
	v89 = int64(24)
	v91 = int64(4278190080)
	v93 = int64(8)
	if v80 == int64(9223372036854767616) {
		goto L40
	} else {
		goto L41
	}
L40:
	;
	v117 = int64(-129)
	goto L42
L41:
	;
	v117 = v82&int64(57344)<<(uint(v85)%64) | (v82&v87<<(uint(v89)%64) | v82&v91<<(uint(v93)%64)) | (int64(base.Ui64(v82)>>(uint(v93)%64))&v91 | int64(base.Ui64(v82)>>(uint(v89)%64))&v87 | (int64(base.Ui64(v82)>>(uint(v85)%64))&int64(65280) | int64(base.Ui64(v82)>>(uint(int64(56))%64))))
	goto L42
L42:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v22)+24)) = v117
	v123 = F_raxInsert(m, v41, v22+int32(24), int32(8), l2, int32(0))
	mBase = m.M
	v124 = m.ExcPending
	if v124 != 0 {
		goto L23
	} else {
		goto L43
	}
L43:
	;
	v845 = l1
	goto L1
L44:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v126))) = int64(17179869184)
	v130 = m.T0[l0].(func(*base.Module, int32) int64)(m, v50)
	mBase = m.M
	v131 = m.ExcPending
	if v131 != 0 {
		goto L23
	} else {
		goto L47
	}
L45:
	;
	if v158 == int32(0) {
		goto L18
	} else {
		goto L57
	}
L46:
	;
	v145 = int32(0)
	v147 = F_pvInsertAt(m, v126, l2, v145)
	mBase = m.M
	v148 = m.ExcPending
	if v148 != 0 {
		goto L23
	} else {
		goto L54
	}
L47:
	;
	if l3 <= v130 {
		goto L46
	} else {
		goto L48
	}
L48:
	;
	v133 = int32(0)
	v135 = F_pvInsertAt(m, v126, v50, v133)
	mBase = m.M
	v136 = m.ExcPending
	if v136 != 0 {
		goto L23
	} else {
		goto L50
	}
L49:
	;
	v143 = F_pvInsertAt(m, v135, l2, v142)
	mBase = m.M
	v144 = m.ExcPending
	if v144 != 0 {
		goto L23
	} else {
		goto L52
	}
L50:
	;
	if v135 == int32(0) {
		v142 = v133
		goto L49
	} else {
		goto L51
	}
L51:
	;
	v139 = *(*int32)(unsafe.Add(mBase, uint32(v135)))
	v142 = v139 & int32(1073741823)
	goto L49
L52:
	;
	v158 = v143
	goto L45
L53:
	;
	v155 = F_pvInsertAt(m, v147, v50, v154)
	mBase = m.M
	v156 = m.ExcPending
	if v156 != 0 {
		goto L23
	} else {
		goto L56
	}
L54:
	;
	if v147 == int32(0) {
		v154 = v145
		goto L53
	} else {
		goto L55
	}
L55:
	;
	v151 = *(*int32)(unsafe.Add(mBase, uint32(v147)))
	v154 = v151 & int32(1073741823)
	goto L53
L56:
	;
	v158 = v155
	goto L45
L57:
	;
	v163 = v158 | int32(2)
	if v163 == int32(-1) {
		goto L17
	} else {
		goto L58
	}
L58:
	;
	if v158&int32(1) != 0 {
		goto L17
	} else {
		goto L59
	}
L59:
	;
	if v163&int32(6) != int32(2) {
		goto L17
	} else {
		goto L60
	}
L60:
	;
	v172 = *(*int32)(unsafe.Add(mBase, uint32(v22)+4))
	v175 = *(*int32)(unsafe.Add(mBase, uint32(v172)))
	if v163 == int32(0) {
		goto L62
	} else {
		goto L63
	}
L61:
	;
	v845 = l1
	goto L1
L62:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v172))) = v175 | int32(3)
	goto L61
L63:
	;
	v178 = int32(3)
	v179 = int32(base.Ui32(v175) >> (uint(v178) % 32))
	v186 = int32(4)
	if v175&v186 != 0 {
		goto L64
	} else {
		goto L65
	}
L64:
	;
	v191 = v186
	goto L66
L65:
	;
	v191 = v179 << (uint(int32(2)) % 32)
	goto L66
L66:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v172+v179+(int32(0)-v179)&v178+v191+int32(4)))) = v163
	*(*int32)(unsafe.Add(mBase, uint32(v172))) = v175&int32(-4) | int32(1)
	goto L61
L67:
	;
	v658 = F_insertToBucket_VECTOR(m, v50, l2, int32(-1))
	mBase = m.M
	v659 = m.ExcPending
	if v659 != 0 {
		goto L23
	} else {
		goto L155
	}
L68:
	;
	v209 = *(*int64)(unsafe.Add(mBase, uint32(v206)))
	if v209&int64(1073741823) != int64(127) {
		goto L67
	} else {
		goto L69
	}
L69:
	;
	v214 = *(*int32)(unsafe.Add(mBase, uint32(v22)+4))
	v215 = *(*int64)(unsafe.Add(mBase, uint32(v22)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v22)+40)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v22)+36)) = int32(-1)
	v221 = *(*int32)(unsafe.Add(mBase, _c_F_insertToBucket_RAX[0]))
	if v221 != 0 {
		goto L16
	} else {
		goto L70
	}
L70:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_insertToBucket_RAX[0])) = l0
	v224 = *(*int64)(unsafe.Add(mBase, uint32(v206)))
	if v224&int64(1073741822) == int64(0) {
		v239 = v224
		goto L71
	} else {
		goto L72
	}
L71:
	;
	v240 = int32(0)
	*(*int32)(unsafe.Add(mBase, _c_F_insertToBucket_RAX[0])) = v240
	v247 = v206 + int32(8)
	v251 = *(*int32)(unsafe.Add(mBase, uint32(base.I32_wrap_i64(v239)<<(uint(int32(2))%32)+v247+int32(-4))))
	v252 = m.T0[l0].(func(*base.Module, int32) int64)(m, v251)
	mBase = m.M
	v253 = m.ExcPending
	if v253 != 0 {
		goto L23
	} else {
		goto L74
	}
L72:
	;
	F_qsort(m, v206+int32(8), base.I32_wrap_i64(v224)&int32(1073741823), int32(4), int32(1128))
	mBase = m.M
	v237 = m.ExcPending
	if v237 != 0 {
		goto L23
	} else {
		goto L73
	}
L73:
	;
	v238 = *(*int64)(unsafe.Add(mBase, uint32(v206)))
	v239 = v238
	goto L71
L74:
	;
	v254 = *(*int32)(unsafe.Add(mBase, uint32(v206)+8))
	v255 = m.T0[l0].(func(*base.Module, int32) int64)(m, v254)
	mBase = m.M
	v256 = m.ExcPending
	if v256 != 0 {
		goto L23
	} else {
		goto L75
	}
L75:
	;
	v259 = v252 & int64(-16)
	v261 = v259 + int64(16)
	v263 = base.B2i32(v259 == int64(9223372036854775792))
	if v259 == int64(9223372036854775792) {
		goto L79
	} else {
		goto L80
	}
L76:
	;
	v611 = F_insertToBucket_VECTOR(m, v50, l2, int32(-1))
	mBase = m.M
	v612 = m.ExcPending
	if v612 != 0 {
		goto L23
	} else {
		goto L146
	}
L77:
	;
	v566 = int64(56)
	v568 = int64(65280)
	v570 = int64(40)
	v573 = int64(16711680)
	v575 = int64(24)
	v577 = int64(4278190080)
	v579 = int64(8)
	*(*int64)(unsafe.Add(mBase, uint32(v22)+40)) = v554<<(uint(v566)%64) | v554&v568<<(uint(v570)%64) | (v554&v573<<(uint(v575)%64) | v554&v577<<(uint(v579)%64)) | (int64(base.Ui64(v554)>>(uint(v579)%64))&v577 | int64(base.Ui64(v554)>>(uint(v575)%64))&v573 | (int64(base.Ui64(v554)>>(uint(v570)%64))&v568 | int64(base.Ui64(v554)>>(uint(v566)%64))))
	v606 = F_raxInsert(m, v41, v22+int32(40), int32(8), v561, int32(0))
	mBase = m.M
	v607 = m.ExcPending
	if v607 != 0 {
		goto L23
	} else {
		goto L142
	}
L78:
	;
	v320 = v255 & int64(-16)
	if v320 == int64(9223372036854775792) {
		goto L87
	} else {
		goto L88
	}
L79:
	;
	v264 = int64(9223372036854775807)
	goto L81
L80:
	;
	v264 = v261
	goto L81
L81:
	;
	if v215 <= v264 {
		goto L78
	} else {
		goto L82
	}
L82:
	;
	v266 = int64(56)
	v268 = int64(65280)
	v270 = int64(40)
	v273 = int64(16711680)
	v275 = int64(24)
	v277 = int64(4278190080)
	v279 = int64(8)
	*(*int64)(unsafe.Add(mBase, uint32(v22)+40)) = v215<<(uint(v266)%64) | v215&v268<<(uint(v270)%64) | (v215&v273<<(uint(v275)%64) | v215&v277<<(uint(v279)%64)) | (int64(base.Ui64(v215)>>(uint(v279)%64))&v277 | int64(base.Ui64(v215)>>(uint(v275)%64))&v273 | (int64(base.Ui64(v215)>>(uint(v270)%64))&v268 | int64(base.Ui64(v215)>>(uint(v266)%64))))
	v307 = F_raxRemove(m, v41, v22+int32(40), int32(8), v22+int32(36))
	mBase = m.M
	v308 = m.ExcPending
	if v308 != 0 {
		goto L23
	} else {
		goto L83
	}
L83:
	;
	if v307 == int32(0) {
		goto L15
	} else {
		goto L84
	}
L84:
	;
	v311 = *(*int32)(unsafe.Add(mBase, uint32(v22)+36))
	if v311 == v50 {
		v554 = v261
		v561 = v311
		goto L77
	} else {
		goto L85
	}
L85:
	;
	F__serverAssert(m, int32(_a_F_insertToBucket_RAX_16), int32(_a_F_insertToBucket_RAX_1), int32(1166))
	mBase = m.M
	v317 = m.ExcPending
	if v317 != 0 {
		goto L23
	} else {
		goto L86
	}
L86:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L87:
	;
	v326 = v263
	goto L89
L88:
	;
	v326 = base.B2i32(v320+int64(16) == v264)
	goto L89
L89:
	;
	if v326 != 0 {
		goto L76
	} else {
		goto L90
	}
L90:
	;
	v327 = *(*int32)(unsafe.Add(mBase, uint32(v206)))
	v329 = v327 & int32(1073741823)
	if base.Ui32(v329) < base.Ui32(int32(2)) {
		goto L2
	} else {
		goto L91
	}
L91:
	;
	v333 = int32(base.Ui32(v329) >> (uint(int32(1)) % 32))
	v351 = int32(0)
	goto L93
L92:
	;
	if v215 <= v450 {
		goto L2
	} else {
		goto L120
	}
L93:
	;
	v354 = v333 - v351
	if v354 < int32(1) {
		goto L95
	} else {
		goto L96
	}
L95:
	;
	v397 = v351 + v333
	if v397 == int32(0) {
		goto L106
	} else {
		goto L107
	}
L96:
	;
	v358 = v354 + int32(-1)
	v359 = *(*int32)(unsafe.Add(mBase, uint32(v206)))
	if base.Ui32(v359&int32(1073741823)) <= base.Ui32(v358) {
		goto L14
	} else {
		goto L97
	}
L97:
	;
	v366 = *(*int32)(unsafe.Add(mBase, uint32(v247+v358<<(uint(int32(2))%32))))
	v367 = m.T0[l0].(func(*base.Module, int32) int64)(m, v366)
	mBase = m.M
	v368 = m.ExcPending
	if v368 != 0 {
		goto L23
	} else {
		goto L98
	}
L98:
	;
	v369 = *(*int32)(unsafe.Add(mBase, uint32(v206)))
	if base.Ui32(v369&int32(1073741823)) <= base.Ui32(v354) {
		goto L13
	} else {
		goto L99
	}
L99:
	;
	v376 = *(*int32)(unsafe.Add(mBase, uint32(v247+v354<<(uint(int32(2))%32))))
	v377 = m.T0[l0].(func(*base.Module, int32) int64)(m, v376)
	mBase = m.M
	v378 = m.ExcPending
	if v378 != 0 {
		goto L23
	} else {
		goto L100
	}
L100:
	;
	v380 = v367 & int64(-16)
	if v380 == int64(9223372036854775792) {
		goto L95
	} else {
		goto L101
	}
L101:
	;
	v383 = int64(16)
	v384 = v380 + v383
	v387 = v377 & int64(-16)
	if v387 == int64(9223372036854775792) {
		goto L102
	} else {
		goto L103
	}
L102:
	;
	v392 = int64(9223372036854775807)
	goto L104
L103:
	;
	v392 = v387 + v383
	goto L104
L104:
	;
	if v384 < v392 {
		v450 = v384
		v452 = v354
		goto L92
	} else {
		goto L105
	}
L105:
	;
	goto L95
L106:
	;
	v443 = v351 + int32(1)
	if base.Ui32(v443) < base.Ui32(v333) {
		v351 = v443
		goto L93
	} else {
		goto L118
	}
L107:
	;
	v400 = *(*int32)(unsafe.Add(mBase, uint32(v206)))
	v402 = v400 & int32(1073741823)
	if base.Ui32(v402) <= base.Ui32(v397) {
		goto L106
	} else {
		goto L108
	}
L108:
	;
	v405 = v397 + int32(-1)
	if base.Ui32(v402) <= base.Ui32(v405) {
		goto L12
	} else {
		goto L109
	}
L109:
	;
	v410 = *(*int32)(unsafe.Add(mBase, uint32(v247+v405<<(uint(int32(2))%32))))
	v411 = m.T0[l0].(func(*base.Module, int32) int64)(m, v410)
	mBase = m.M
	v412 = m.ExcPending
	if v412 != 0 {
		goto L23
	} else {
		goto L110
	}
L110:
	;
	v413 = *(*int32)(unsafe.Add(mBase, uint32(v206)))
	if base.Ui32(v413&int32(1073741823)) <= base.Ui32(v397) {
		goto L11
	} else {
		goto L111
	}
L111:
	;
	v420 = *(*int32)(unsafe.Add(mBase, uint32(v247+v397<<(uint(int32(2))%32))))
	v421 = m.T0[l0].(func(*base.Module, int32) int64)(m, v420)
	mBase = m.M
	v422 = m.ExcPending
	if v422 != 0 {
		goto L23
	} else {
		goto L112
	}
L112:
	;
	v424 = v411 & int64(-16)
	if v424 == int64(9223372036854775792) {
		goto L106
	} else {
		goto L113
	}
L113:
	;
	v427 = int64(16)
	v428 = v424 + v427
	v431 = v421 & int64(-16)
	if v431 == int64(9223372036854775792) {
		goto L114
	} else {
		goto L115
	}
L114:
	;
	v436 = int64(9223372036854775807)
	goto L116
L115:
	;
	v436 = v431 + v427
	goto L116
L116:
	;
	if v428 < v436 {
		v450 = v428
		v452 = v397
		goto L92
	} else {
		goto L117
	}
L117:
	;
	goto L106
L118:
	;
	v446 = *(*int32)(unsafe.Add(mBase, uint32(v206)))
	if base.Ui32(v443+v333) < base.Ui32(v446&int32(1073741823)) {
		v351 = v443
		goto L93
	} else {
		goto L119
	}
L119:
	;
	goto L2
L120:
	;
	v456 = *(*int32)(unsafe.Add(mBase, uint32(v206)))
	if v452 == v456&int32(1073741823) {
		goto L10
	} else {
		goto L121
	}
L121:
	;
	switch v50 + int32(1) {
	case 0:
		goto L123
	case 1:
		goto L124
	default:
		goto L122
	}
L122:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+32)) = v206
	v477 = F_pvSplit(m, v22+int32(32), v452)
	mBase = m.M
	v478 = m.ExcPending
	if v478 != 0 {
		goto L23
	} else {
		goto L127
	}
L123:
	;
	F__serverAssert(m, int32(_a_F_insertToBucket_RAX_4), int32(_a_F_insertToBucket_RAX_1), int32(797))
	mBase = m.M
	v472 = m.ExcPending
	if v472 != 0 {
		goto L23
	} else {
		goto L126
	}
L124:
	;
	F__serverAssert(m, int32(_a_F_insertToBucket_RAX_14), int32(_a_F_insertToBucket_RAX_1), int32(781))
	mBase = m.M
	v466 = m.ExcPending
	if v466 != 0 {
		goto L23
	} else {
		goto L125
	}
L125:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L126:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L127:
	;
	if v477 == int32(0) {
		goto L9
	} else {
		goto L128
	}
L128:
	;
	v481 = *(*int32)(unsafe.Add(mBase, uint32(v22)+32))
	if v481 == int32(0) {
		goto L8
	} else {
		goto L129
	}
L129:
	;
	v484 = int32(2)
	v485 = v481 | v484
	*(*int32)(unsafe.Add(mBase, uint32(v22)+36)) = v485
	if v485&int32(7) != v484 {
		goto L7
	} else {
		goto L130
	}
L130:
	;
	v492 = v481 & int32(-8)
	if v492 == int32(0) {
		goto L6
	} else {
		goto L131
	}
L131:
	;
	v495 = *(*int32)(unsafe.Add(mBase, uint32(v492)))
	if v495&int32(1073741823) == int32(0) {
		goto L6
	} else {
		goto L132
	}
L132:
	;
	v500 = int32(2)
	v501 = v477 | v500
	if v501&int32(7) != v500 {
		goto L5
	} else {
		goto L133
	}
L133:
	;
	v507 = v477 & int32(-8)
	if v507 == int32(0) {
		goto L4
	} else {
		goto L134
	}
L134:
	;
	v510 = *(*int32)(unsafe.Add(mBase, uint32(v507)))
	if v510&int32(1073741823) == int32(0) {
		goto L4
	} else {
		goto L135
	}
L135:
	;
	v517 = *(*int32)(unsafe.Add(mBase, uint32(v214)))
	if v501 == int32(0) {
		goto L137
	} else {
		goto L138
	}
L136:
	;
	v554 = v450
	v561 = v485
	goto L77
L137:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v214))) = v517 | int32(3)
	goto L136
L138:
	;
	v520 = int32(3)
	v521 = int32(base.Ui32(v517) >> (uint(v520) % 32))
	v528 = int32(4)
	if v517&v528 != 0 {
		goto L139
	} else {
		goto L140
	}
L139:
	;
	v533 = v528
	goto L141
L140:
	;
	v533 = v521 << (uint(int32(2)) % 32)
	goto L141
L141:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v214+v521+(int32(0)-v521)&v520+v533+int32(4)))) = v501
	*(*int32)(unsafe.Add(mBase, uint32(v214))) = v517&int32(-4) | int32(1)
	goto L136
L142:
	;
	v608 = F_insertToBucket_RAX(m, l0, l1, l2, l3)
	mBase = m.M
	v609 = m.ExcPending
	if v609 != 0 {
		goto L23
	} else {
		goto L143
	}
L143:
	;
	v845 = v608
	goto L1
L144:
	;
	if v611&int32(7) != int32(4) {
		goto L3
	} else {
		goto L148
	}
L145:
	;
	F__serverAssert(m, int32(_a_F_insertToBucket_RAX_14), int32(_a_F_insertToBucket_RAX_1), int32(781))
	mBase = m.M
	v619 = m.ExcPending
	if v619 != 0 {
		goto L23
	} else {
		goto L147
	}
L146:
	;
	switch v611 + int32(1) {
	case 0:
		goto L3
	case 1:
		goto L145
	default:
		goto L144
	}
L147:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L148:
	;
	v627 = *(*int32)(unsafe.Add(mBase, uint32(v214)))
	if v611 == int32(0) {
		goto L150
	} else {
		goto L151
	}
L149:
	;
	v845 = l1
	goto L1
L150:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v214))) = v627 | int32(3)
	goto L149
L151:
	;
	v630 = int32(3)
	v631 = int32(base.Ui32(v627) >> (uint(v630) % 32))
	v638 = int32(4)
	if v627&v638 != 0 {
		goto L152
	} else {
		goto L153
	}
L152:
	;
	v643 = v638
	goto L154
L153:
	;
	v643 = v631 << (uint(int32(2)) % 32)
	goto L154
L154:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v214+v631+(int32(0)-v631)&v630+v643+int32(4)))) = v611
	*(*int32)(unsafe.Add(mBase, uint32(v214))) = v627&int32(-4) | int32(1)
	goto L149
L155:
	;
	if v658 == v50 {
		v845 = l1
		goto L1
	} else {
		goto L156
	}
L156:
	;
	v661 = *(*int32)(unsafe.Add(mBase, uint32(v22)+4))
	v664 = *(*int32)(unsafe.Add(mBase, uint32(v661)))
	if v658 == int32(0) {
		goto L158
	} else {
		goto L159
	}
L157:
	;
	v845 = l1
	goto L1
L158:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v661))) = v664 | int32(3)
	goto L157
L159:
	;
	v667 = int32(3)
	v668 = int32(base.Ui32(v664) >> (uint(v667) % 32))
	v675 = int32(4)
	if v664&v675 != 0 {
		goto L160
	} else {
		goto L161
	}
L160:
	;
	v680 = v675
	goto L162
L161:
	;
	v680 = v668 << (uint(int32(2)) % 32)
	goto L162
L162:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v661+v668+(int32(0)-v668)&v667+v680+int32(4)))) = v658
	*(*int32)(unsafe.Add(mBase, uint32(v661))) = v664&int32(-4) | int32(1)
	goto L157
L163:
	;
	if v696 != 0 {
		v845 = l1
		goto L1
	} else {
		goto L164
	}
L164:
	;
	F__serverAssert(m, int32(_a_F_insertToBucket_RAX_17), int32(_a_F_insertToBucket_RAX_1), int32(1250))
	mBase = m.M
	v702 = m.ExcPending
	if v702 != 0 {
		goto L23
	} else {
		goto L165
	}
L165:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L166:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L167:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L168:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L169:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L170:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L171:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L172:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L173:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L174:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L175:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L176:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L177:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L178:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L179:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L180:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L181:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
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
L184:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L185:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_instanceLinkCloseConnection(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v9 int32
	_ = v9
	var v18 int32
	_ = v18
	if l1 == int32(0) {
		return
	} else {
		v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
		if v5 != l1 {
		} else {
			*(*int64)(unsafe.Add(mBase, uint32(l0)+8)) = int64(0)
		}
		v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
		if v9 != l1 {
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = int32(0)
		}
		*(*int32)(unsafe.Add(mBase, uint32(l1)+212)) = int32(0)
		*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = int32(1)
		F_valkeyAsyncFree(m, l1)
		mBase = m.M
		v18 = m.ExcPending
		if v18 != 0 {
			return
		} else {
			return
		}
	}
}
func F_interleave64(m *base.Module, l0 int32, l1 int32) int64 {
	var v4 int64
	_ = v4
	var v5 int64
	_ = v5
	var v8 int64
	_ = v8
	var v9 int64
	_ = v9
	var v10 int64
	_ = v10
	var v13 int64
	_ = v13
	var v14 int64
	_ = v14
	var v15 int64
	_ = v15
	var v18 int64
	_ = v18
	var v19 int64
	_ = v19
	var v20 int64
	_ = v20
	var v23 int64
	_ = v23
	var v24 int64
	_ = v24
	var v27 int64
	_ = v27
	var v32 int64
	_ = v32
	var v37 int64
	_ = v37
	var v42 int64
	_ = v42
	var v47 int64
	_ = v47
	var v52 int64
	_ = v52
	v4 = base.I64_extend_i32_u(l1)
	v5 = int64(16)
	v8 = int64(281470681808895)
	v9 = (v4<<(uint(v5)%64) | v4) & v8
	v10 = int64(8)
	v13 = int64(71777214294589695)
	v14 = (v9<<(uint(v10)%64) | v9) & v13
	v15 = int64(4)
	v18 = int64(1085102592571150095)
	v19 = (v14<<(uint(v15)%64) | v14) & v18
	v20 = int64(2)
	v23 = int64(3689348814741910323)
	v24 = (v19<<(uint(v20)%64) | v19) & v23
	v27 = int64(1)
	v32 = base.I64_extend_i32_u(l0)
	v37 = (v32<<(uint(v5)%64) | v32) & v8
	v42 = (v37<<(uint(v10)%64) | v37) & v13
	v47 = (v42<<(uint(v15)%64) | v42) & v18
	v52 = (v47<<(uint(v20)%64) | v47) & v23
	return (v24<<(uint(v20)%64)|v24<<(uint(v27)%64))&int64(-6148914691236517206) | (v52<<(uint(v27)%64)|v52)&int64(6148914691236517205)
}
func F_invalidateCommandCache(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v1 int32
	_ = v1
	var v3 int32
	_ = v3
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	v1 = int32(0)
	v3 = *(*int32)(unsafe.Add(mBase, _c_F_invalidateCommandCache[0]))
	if v3 == v1 {
		v11 = int32(0)
		v12 = *(*int32)(unsafe.Add(mBase, _c_F_invalidateCommandCache[1]))
		if v12 == v11 {
			return
		} else {
			F_sdsfree(m, v12)
			mBase = m.M
			v16 = m.ExcPending
			if v16 != 0 {
				return
			} else {
				v17 = int32(0)
				*(*int32)(unsafe.Add(mBase, _c_F_invalidateCommandCache[1])) = v17
				return
			}
		}
	} else {
		F_sdsfree(m, v3)
		mBase = m.M
		v7 = m.ExcPending
		if v7 != 0 {
			return
		} else {
			v8 = int32(0)
			*(*int32)(unsafe.Add(mBase, _c_F_invalidateCommandCache[0])) = v8
			v11 = int32(0)
			v12 = *(*int32)(unsafe.Add(mBase, _c_F_invalidateCommandCache[1]))
			if v12 == v11 {
				return
			} else {
				F_sdsfree(m, v12)
				mBase = m.M
				v16 = m.ExcPending
				if v16 != 0 {
					return
				} else {
					v17 = int32(0)
					*(*int32)(unsafe.Add(mBase, _c_F_invalidateCommandCache[1])) = v17
					return
				}
			}
		}
	}
}
func F_ioctl(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v14 int32
	_ = v14
	var v19 int32
	_ = v19
	var v27 int32
	_ = v27
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v44 int32
	_ = v44
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v87 int32
	_ = v87
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v100 int32
	_ = v100
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v112 int32
	_ = v112
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v120 int32
	_ = v120
	var v121 int64
	_ = v121
	var v124 int32
	_ = v124
	var v127 int32
	_ = v127
	var v144 int32
	_ = v144
	var v146 int32
	_ = v146
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v162 int32
	_ = v162
	var v165 int32
	_ = v165
	var v172 int32
	_ = v172
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v193 int32
	_ = v193
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v224 int32
	_ = v224
	var v235 int32
	_ = v235
	var v237 int32
	_ = v237
	var v239 int32
	_ = v239
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v257 int32
	_ = v257
	var v260 int32
	_ = v260
	var v261 int32
	_ = v261
	var v267 int32
	_ = v267
	var v276 int32
	_ = v276
	var v277 int32
	_ = v277
	var v280 int32
	_ = v280
	var v286 int32
	_ = v286
	var v287 int32
	_ = v287
	var v288 int32
	_ = v288
	var v292 int32
	_ = v292
	var v310 int32
	_ = v310
	var v311 int32
	_ = v311
	var v315 int32
	_ = v315
	var v316 int64
	_ = v316
	var v319 int32
	_ = v319
	var v322 int32
	_ = v322
	var v324 int32
	_ = v324
	var v326 int32
	_ = v326
	var v328 int32
	_ = v328
	var v329 int32
	_ = v329
	var v342 int32
	_ = v342
	var v345 int32
	_ = v345
	var v352 int32
	_ = v352
	var v360 int32
	_ = v360
	var v361 int32
	_ = v361
	var v376 int32
	_ = v376
	var v396 int32
	_ = v396
	var v402 int32
	_ = v402
	var v407 int32
	_ = v407
	v7 = m.G0
	v9 = v7 - int32(288)
	m.G0 = v9
	*(*int32)(unsafe.Add(mBase, uint32(v9)+284)) = l2 + int32(4)
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	*(*int64)(unsafe.Add(mBase, uint32(v9)+16)) = base.I64_extend_i32_u(v14)
	v19 = m.Env.X__syscall_ioctl(m, l0, l1, v9+int32(16))
	mBase = m.M
	if l1 == int32(0) {
		v396 = v19
		goto L1
	} else {
		goto L2
	}
L1:
	;
	if base.Ui32(v396) < base.Ui32(int32(-4095)) {
		v407 = v396
		goto L49
	} else {
		goto L50
	}
L2:
	;
	if v19 != int32(-59) {
		v396 = v19
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v27 = int32(0)
	goto L5
L4:
	;
	v44 = int32(1)
	v55 = m.G0
	v57 = v55 - int32(16)
	m.G0 = v57
	v59 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v32)+uint32(_c_F_ioctl[0]))))
	if v44&v59 == int32(0) {
		goto L10
	} else {
		goto L11
	}
L5:
	;
	v32 = v27 * int32(20)
	v34 = v32 + int32(_a_F_ioctl_0)
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v32)+uint32(_c_F_ioctl[1])))
	if v35 == l1 {
		goto L4
	} else {
		goto L7
	}
L6:
	;
	v396 = int32(-59)
	goto L1
L7:
	;
	v38 = v27 + int32(1)
	if v38 != int32(20) {
		v27 = v38
		goto L5
	} else {
		goto L8
	}
L8:
	;
	goto L6
L9:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v9))) = base.I64_extend_i32_u(v9 + int32(24))
	v218 = *(*int32)(unsafe.Add(mBase, uint32(v32)+uint32(_c_F_ioctl[2])))
	v219 = m.Env.X__syscall_ioctl(m, l0, v218, v9)
	mBase = m.M
	if v219 < int32(0) {
		v396 = v219
		goto L1
	} else {
		goto L28
	}
L10:
	;
	m.G0 = v57 + int32(16)
	goto L9
L11:
	;
	v63 = v34
	v64 = v9 + int32(24)
	v65 = v14
	goto L13
L12:
	;
	goto L27
L13:
	;
	v77 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v63)+8)))
	if v77 == int32(0) {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	F_convert_ioctl_struct(m, v63+int32(20), v64, v65, v44)
	mBase = m.M
	F_convert_ioctl_struct(m, v63+int32(40), v64+int32(4), v65+int32(8), v44)
	mBase = m.M
	v162 = v63 + int32(60)
	v165 = int32(72)
	F_convert_ioctl_struct(m, v162, v64+int32(68), v65+v165, v44)
	mBase = m.M
	v172 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v63)+69)))
	if v44&v172 != 0 {
		v63 = v162
		v64 = v64 + v165
		v65 = v65 + int32(76)
		goto L13
	} else {
		goto L25
	}
L16:
	;
	v80 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v63)+11)))
	if v80 != 0 {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	v87 = int32(0)
	v96 = v87
	v97 = v87
	v100 = v87
	goto L19
L18:
	;
	v81 = int32(0)
	v180 = v81
	v181 = v81
	goto L12
L19:
	;
	v106 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v63+int32(12)+v100))))
	v107 = v106 - v96
	v108 = v107 + v97
	v112 = (int32(0)-v108)&int32(7) + v108
	goto L23
L21:
	;
	v144 = v106 + int32(4)
	v146 = v112 + int32(8)
	v148 = v100 + int32(1)
	v149 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v63)+11)))
	if base.Ui32(v149) <= base.Ui32(v148) {
		v180 = v144
		v181 = v146
		goto L12
	} else {
		goto L24
	}
L23:
	;
	v115 = F___memcpy(m, v64+v96, v65+v97, v107)
	mBase = m.M
	v116 = int32(8)
	v120 = F___memcpy(m, v57+v116, v65+v112, v116)
	mBase = m.M
	v121 = *(*int64)(unsafe.Add(mBase, uint32(v57)+8))
	*(*uint32)(unsafe.Add(mBase, uint32(v57)+4)) = uint32(v121)
	v124 = int32(4)
	v127 = F___memcpy(m, v64+v106, v57+v124, v124)
	mBase = m.M
	goto L21
L24:
	;
	v96 = v144
	v97 = v146
	v100 = v148
	goto L19
L25:
	;
	goto L10
L27:
	;
	v193 = F___memcpy(m, v64+v180, v65+v181, v77-v180)
	mBase = m.M
	goto L10
L28:
	;
	v224 = int32(2)
	v235 = m.G0
	v237 = v235 - int32(16)
	m.G0 = v237
	v239 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v32)+uint32(_c_F_ioctl[0]))))
	if v224&v239 == int32(0) {
		goto L30
	} else {
		goto L31
	}
L29:
	;
	v396 = v219
	goto L1
L30:
	;
	m.G0 = v237 + int32(16)
	goto L29
L31:
	;
	v243 = v34
	v244 = v9 + int32(24)
	v245 = v14
	goto L33
L32:
	;
	goto L46
L33:
	;
	v257 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v243)+8)))
	if v257 == int32(0) {
		goto L35
	} else {
		goto L36
	}
L35:
	;
	F_convert_ioctl_struct(m, v243+int32(20), v244, v245, v224)
	mBase = m.M
	F_convert_ioctl_struct(m, v243+int32(40), v244+int32(4), v245+int32(8), v224)
	mBase = m.M
	v342 = v243 + int32(60)
	v345 = int32(72)
	F_convert_ioctl_struct(m, v342, v244+int32(68), v245+v345, v224)
	mBase = m.M
	v352 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v243)+69)))
	if v224&v352 != 0 {
		v243 = v342
		v244 = v244 + v345
		v245 = v245 + int32(76)
		goto L33
	} else {
		goto L45
	}
L36:
	;
	v260 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v243)+11)))
	if v260 != 0 {
		goto L37
	} else {
		goto L38
	}
L37:
	;
	v267 = int32(0)
	v276 = v267
	v277 = v267
	v280 = v267
	goto L39
L38:
	;
	v261 = int32(0)
	v360 = v261
	v361 = v261
	goto L32
L39:
	;
	v286 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v243+int32(12)+v280))))
	v287 = v286 - v276
	v288 = v287 + v277
	v292 = (int32(0)-v288)&int32(7) + v288
	goto L42
L41:
	;
	v324 = v286 + int32(4)
	v326 = v292 + int32(8)
	v328 = v280 + int32(1)
	v329 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v243)+11)))
	if base.Ui32(v329) <= base.Ui32(v328) {
		v360 = v324
		v361 = v326
		goto L32
	} else {
		goto L44
	}
L42:
	;
	v310 = F___memcpy(m, v245+v277, v244+v276, v287)
	mBase = m.M
	v311 = int32(4)
	v315 = F___memcpy(m, v237+v311, v244+v286, v311)
	mBase = m.M
	v316 = int64(*(*int32)(unsafe.Add(mBase, uint32(v237)+4)))
	*(*int64)(unsafe.Add(mBase, uint32(v237)+8)) = v316
	v319 = int32(8)
	v322 = F___memcpy(m, v245+v292, v237+v319, v319)
	mBase = m.M
	goto L41
L44:
	;
	v276 = v324
	v277 = v326
	v280 = v328
	goto L39
L45:
	;
	goto L30
L46:
	;
	v376 = F___memcpy(m, v245+v361, v244+v360, v257-v360)
	mBase = m.M
	goto L30
L48:
	;
	m.G0 = v9 + int32(288)
	return v407
L49:
	;
	goto L48
L50:
	;
	v402 = F___errno_location(m)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v402))) = int32(0) - v396
	v407 = int32(-1)
	goto L49
}
func F_isAnySlotInManualImportingState(m *base.Module) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	v3 = *(*int32)(unsafe.Add(mBase, _c_F_isAnySlotInManualImportingState[0]))
	v4 = *(*int32)(unsafe.Add(mBase, uint32(v3)+48))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(v4)+12))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(v4)+16))
	return base.B2i32(v5 != int32(0)-v7)
}
func F_isPausedActionsWithUpdate(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	v3 = *(*int32)(unsafe.Add(mBase, _c_F_isPausedActionsWithUpdate[0]))
	if v3&l0 != 0 {
		F_updatePausedActions(m)
		mBase = m.M
		v10 = m.ExcPending
		if v10 != 0 {
			return int32(0)
		} else {
			v12 = *(*int32)(unsafe.Add(mBase, _c_F_isPausedActionsWithUpdate[0]))
			return v12 & l0
		}
	} else {
		return int32(0)
	}
}
func F_is_leap(m *base.Module, l0 int32) int32 {
	var v6 int32
	_ = v6
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v22 int32
	_ = v22
	if int32(2147481747) < l0 {
		v6 = l0 + int32(-2000)
	} else {
		v6 = l0
	}
	if v6&int32(3) == int32(0) {
		v14 = v6 + int32(1900)
		v16 = base.I32_rem_s(v14, int32(100))
		if v16 == int32(0) {
			v22 = base.I32_rem_s(v14, int32(400))
			return base.B2i32(v22 == int32(0))
		} else {
			return int32(1)
		}
	} else {
		return int32(0)
	}
}
func F_isatty(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v22 int32
	_ = v22
	v4 = m.G0
	v6 = v4 - int32(32)
	m.G0 = v6
	v10 = m.Wasi_snapshot_preview1.Fd_fdstat_get(m, l0, v6+int32(8))
	mBase = m.M
	if v10 != 0 {
		v16 = v10
		*(*int32)(unsafe.Add(mBase, _c_F_isatty[0])) = v16
		v22 = int32(0)
	} else {
		v13 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6)+8)))
		if v13 == int32(2) {
			v22 = int32(1)
		} else {
			v16 = int32(59)
			*(*int32)(unsafe.Add(mBase, _c_F_isatty[0])) = v16
			v22 = int32(0)
		}
	}
	m.G0 = v6 + int32(32)
	return v22
}
func F_iswalnum(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v14 int32
	_ = v14
	var v24 int32
	_ = v24
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	if base.Ui32(l0+int32(-48)) < base.Ui32(int32(10)) {
		v35 = int32(1)
	} else {
		if base.Ui32(int32(131071)) < base.Ui32(l0) {
			v32 = base.B2i32(base.Ui32(l0) < base.Ui32(int32(196606)))
		} else {
			v14 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(l0)>>(uint(int32(8))%32)))+uint32(_c_F_iswalnum[0]))))
			v24 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14<<(uint(int32(5))%32)|int32(base.Ui32(l0)>>(uint(int32(3))%32))&int32(31))+uint32(_c_F_iswalnum[0]))))
			v32 = int32(base.Ui32(v24)>>(uint(l0&int32(7))%32)) & int32(1)
		}
		v35 = base.B2i32(v32 != int32(0))
	}
	return v35
}
func F_iswalpha(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v18 int32
	_ = v18
	if base.Ui32(int32(131071)) < base.Ui32(l0) {
		return base.B2i32(base.Ui32(l0) < base.Ui32(int32(196606)))
	} else {
		v8 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(l0)>>(uint(int32(8))%32)))+uint32(_c_F_iswalpha[0]))))
		v18 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8<<(uint(int32(5))%32)|int32(base.Ui32(l0)>>(uint(int32(3))%32))&int32(31))+uint32(_c_F_iswalpha[0]))))
		return int32(base.Ui32(v18)>>(uint(l0&int32(7))%32)) & int32(1)
	}
}
func F_iswcntrl(m *base.Module, l0 int32) int32 {
	return base.B2i32(l0&int32(-2) == int32(8232)) | (base.B2i32(base.Ui32(l0) < base.Ui32(int32(32))) | base.B2i32(base.Ui32(l0+int32(-127)) < base.Ui32(int32(33)))) | base.B2i32(base.Ui32(l0+int32(-65529)) < base.Ui32(int32(3)))
}
func F_iswctype(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v24 int32
	_ = v24
	var v34 int32
	_ = v34
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v89 int32
	_ = v89
	var v104 int32
	_ = v104
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
	var v121 int32
	_ = v121
	var v131 int32
	_ = v131
	var v137 int32
	_ = v137
	var v141 int32
	_ = v141
	var v144 int32
	_ = v144
	var v146 int32
	_ = v146
	var v160 int32
	_ = v160
	switch l1 + int32(-1) {
	case 0:
		if base.Ui32(l0+int32(-48)) < base.Ui32(int32(10)) {
			v16 = int32(1)
		} else {
			v13 = F_iswalpha(m, l0)
			mBase = m.M
			v16 = base.B2i32(v13 != int32(0))
		}
		return v16
	case 1:
		if base.Ui32(int32(131071)) < base.Ui32(l0) {
			v42 = base.B2i32(base.Ui32(l0) < base.Ui32(int32(196606)))
		} else {
			v24 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(l0)>>(uint(int32(8))%32)))+uint32(_c_F_iswctype[0]))))
			v34 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24<<(uint(int32(5))%32)|int32(base.Ui32(l0)>>(uint(int32(3))%32))&int32(31))+uint32(_c_F_iswctype[0]))))
			v42 = int32(base.Ui32(v34)>>(uint(l0&int32(7))%32)) & int32(1)
		}
		return v42
	case 2:
		v44 = F_isblank(m, l0)
		mBase = m.M
		return v44
	case 3:
		return base.B2i32(l0&int32(-2) == int32(8232)) | (base.B2i32(base.Ui32(l0) < base.Ui32(int32(32))) | base.B2i32(base.Ui32(l0+int32(-127)) < base.Ui32(int32(33)))) | base.B2i32(base.Ui32(l0+int32(-65529)) < base.Ui32(int32(3)))
	case 4:
		return base.B2i32(base.Ui32(l0+int32(-48)) < base.Ui32(int32(10)))
	case 5:
		v71 = F_iswspace(m, l0)
		mBase = m.M
		if v71 != 0 {
			v75 = int32(0)
		} else {
			v72 = F_iswprint(m, l0)
			mBase = m.M
			v75 = base.B2i32(v72 != int32(0))
		}
		return v75
	case 6:
		v77 = F_towupper(m, l0)
		mBase = m.M
		return base.B2i32(v77 != l0)
	case 7:
		if base.Ui32(int32(254)) < base.Ui32(l0) {
			v89 = int32(1)
			if base.Ui32(l0+int32(-57344)) < base.Ui32(int32(8185)) {
				v109 = v89
			} else {
				if base.Ui32(l0) < base.Ui32(int32(8232)) {
					v109 = v89
				} else {
					if base.Ui32(l0+int32(-8234)) < base.Ui32(int32(47062)) {
						v109 = v89
					} else {
						v104 = int32(65534)
						v109 = base.B2i32(base.Ui32(l0+int32(-65532)) < base.Ui32(int32(1048580))) & base.B2i32(l0&v104 != v104)
					}
				}
			}
			v111 = v109
		} else {
			v111 = base.B2i32(base.Ui32(int32(32)) < base.Ui32((l0+int32(1))&int32(127)))
		}
		return v111
	case 8:
		if base.Ui32(int32(131071)) < base.Ui32(l0) {
			v137 = int32(0)
		} else {
			v121 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(l0)>>(uint(int32(8))%32)))+uint32(_c_F_iswctype[1]))))
			v131 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v121<<(uint(int32(5))%32)|int32(base.Ui32(l0)>>(uint(int32(3))%32))&int32(31))+uint32(_c_F_iswctype[1]))))
			v137 = int32(base.Ui32(v131)>>(uint(l0&int32(7))%32)) & int32(1)
		}
		return v137
	case 9:
		if l0 != 0 {
			v141 = F_wcschr(m, int32(_a_F_iswctype_0), l0)
			mBase = m.M
			v144 = base.B2i32(v141 != int32(0))
		} else {
			v144 = int32(0)
		}
		return v144
	case 10:
		v146 = F_towlower(m, l0)
		mBase = m.M
		return base.B2i32(v146 != l0)
	case 11:
		v160 = base.B2i32(base.Ui32(l0+int32(-48)) < base.Ui32(int32(10))) | base.B2i32(base.Ui32(l0|int32(32)+int32(-97)) < base.Ui32(int32(6)))
		return v160
	default:
		v160 = int32(0)
		return v160
	}
}
func F_iswgraph(m *base.Module, l0 int32) int32 {
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v19 int32
	_ = v19
	var v34 int32
	_ = v34
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	if l0 != 0 {
		v6 = F_wcschr(m, int32(_a_F_iswgraph_0), l0)
		v9 = base.B2i32(v6 != int32(0))
	} else {
		v9 = int32(0)
	}
	if v9 != 0 {
		v44 = int32(0)
	} else {
		if base.Ui32(int32(254)) < base.Ui32(l0) {
			v19 = int32(1)
			if base.Ui32(l0+int32(-57344)) < base.Ui32(int32(8185)) {
				v39 = v19
			} else {
				if base.Ui32(l0) < base.Ui32(int32(8232)) {
					v39 = v19
				} else {
					if base.Ui32(l0+int32(-8234)) < base.Ui32(int32(47062)) {
						v39 = v19
					} else {
						v34 = int32(65534)
						v39 = base.B2i32(base.Ui32(l0+int32(-65532)) < base.Ui32(int32(1048580))) & base.B2i32(l0&v34 != v34)
					}
				}
			}
			v41 = v39
		} else {
			v41 = base.B2i32(base.Ui32(int32(32)) < base.Ui32((l0+int32(1))&int32(127)))
		}
		v44 = base.B2i32(v41 != int32(0))
	}
	return v44
}
