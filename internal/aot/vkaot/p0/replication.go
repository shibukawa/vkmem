package p0

import (
	base "github.com/shibukawa/vkmem/internal/aot/vkaot/base"
	"unsafe"
)

func F_clearReplicationId2(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v1 int32
	_ = v1
	var v2 int64
	_ = v2
	var v8 int32
	_ = v8
	v1 = int32(_a69)
	v2 = int64(3472328296227680304)
	*(*int64)(unsafe.Add(mBase, _consts[512])) = v2
	*(*int64)(unsafe.Add(mBase, _consts[508])) = int64(-1)
	v8 = int32(0)
	*(*uint8)(unsafe.Add(mBase, _consts[513])) = uint8(v8)
	*(*int64)(unsafe.Add(mBase, _consts[514])) = v2
	*(*int64)(unsafe.Add(mBase, _consts[515])) = v2
	*(*int64)(unsafe.Add(mBase, _consts[516])) = v2
	*(*int64)(unsafe.Add(mBase, _consts[517])) = v2
	return
}
func F_freeReplicationBacklogRefMemAsync(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v13 int64
	_ = v13
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int64
	_ = v18
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if base.Ui32(int32(64)) < base.Ui32(v10) {
		v17 = v10
		v18 = *(*int64)(unsafe.Add(mBase, uint32(l1)+8))
		v19 = int32(0)
		v21 = *(*int32)(unsafe.Add(mBase, _consts[323]))
		*(*int32)(unsafe.Add(mBase, _consts[323])) = v21 + (v17 + base.I32_wrap_i64(v18))
		*(*int32)(unsafe.Add(mBase, uint32(v8))) = l0
		*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = l1
		F_bioCreateLazyFreeJob(m, int32(557), int32(2), v8)
		mBase = m.M
		v31 = m.ExcPending
		if v31 != 0 {
			return
		} else {
			m.G0 = v8 + int32(16)
			return
		}
	} else {
		v13 = *(*int64)(unsafe.Add(mBase, uint32(l1)+8))
		if base.Ui64(v13) < base.Ui64(int64(65)) {
			F_listRelease(m, l0)
			mBase = m.M
			v33 = m.ExcPending
			if v33 != 0 {
				return
			} else {
				F_raxFree(m, l1)
				mBase = m.M
				v35 = m.ExcPending
				if v35 != 0 {
					return
				} else {
					m.G0 = v8 + int32(16)
					return
				}
			}
		} else {
			v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
			v17 = v16
			v18 = *(*int64)(unsafe.Add(mBase, uint32(l1)+8))
			v19 = int32(0)
			v21 = *(*int32)(unsafe.Add(mBase, _consts[323]))
			*(*int32)(unsafe.Add(mBase, _consts[323])) = v21 + (v17 + base.I32_wrap_i64(v18))
			*(*int32)(unsafe.Add(mBase, uint32(v8))) = l0
			*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = l1
			F_bioCreateLazyFreeJob(m, int32(557), int32(2), v8)
			mBase = m.M
			v31 = m.ExcPending
			if v31 != 0 {
				return
			} else {
				m.G0 = v8 + int32(16)
				return
			}
		}
	}
}
func F_replicationCachePrimaryUsingMyself(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v13 int64
	_ = v13
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v24 int64
	_ = v24
	var v31 int32
	_ = v31
	var v38 int64
	_ = v38
	var v45 int64
	_ = v45
	var v52 int64
	_ = v52
	var v59 int64
	_ = v59
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	v3 = *(*int32)(unsafe.Add(mBase, _consts[15]))
	if int32(2) < v3 {
		v11 = int32(_a69)
		v13 = *(*int64)(unsafe.Add(mBase, _consts[31]))
		*(*int64)(unsafe.Add(mBase, _consts[574])) = v13
		F_replicationCreatePrimaryClientWithHandler(m, int32(0), int32(-1), int32(107))
		mBase = m.M
		v19 = m.ExcPending
		if v19 != 0 {
			return
		} else {
			v20 = int32(_a69)
			v21 = *(*int32)(unsafe.Add(mBase, _consts[133]))
			v22 = *(*int32)(unsafe.Add(mBase, uint32(v21)+104))
			v24 = *(*int64)(unsafe.Add(mBase, _consts[528]))
			*(*int64)(unsafe.Add(mBase, uint32(v22)+104)) = v24
			v31 = int32(*(*uint8)(unsafe.Add(mBase, _consts[511])))
			*(*uint8)(unsafe.Add(mBase, uint32(v22+int32(144)))) = uint8(v31)
			v38 = *(*int64)(unsafe.Add(mBase, _consts[532]))
			*(*int64)(unsafe.Add(mBase, uint32(v22+int32(136)))) = v38
			v45 = *(*int64)(unsafe.Add(mBase, _consts[531]))
			*(*int64)(unsafe.Add(mBase, uint32(v22+int32(128)))) = v45
			v52 = *(*int64)(unsafe.Add(mBase, _consts[530]))
			*(*int64)(unsafe.Add(mBase, uint32(v22+int32(120)))) = v52
			v59 = *(*int64)(unsafe.Add(mBase, _consts[529]))
			*(*int64)(unsafe.Add(mBase, uint32(v22+int32(112)))) = v59
			v62 = *(*int32)(unsafe.Add(mBase, _consts[133]))
			F_unlinkClient(m, v62)
			mBase = m.M
			v64 = m.ExcPending
			if v64 != 0 {
				return
			} else {
				v65 = int32(_a69)
				v67 = *(*int32)(unsafe.Add(mBase, _consts[133]))
				*(*int32)(unsafe.Add(mBase, _consts[134])) = v67
				*(*int32)(unsafe.Add(mBase, _consts[133])) = int32(0)
				return
			}
		}
	} else {
		F__serverLog(m, int32(2), int32(_a1044), int32(0))
		mBase = m.M
		v10 = m.ExcPending
		if v10 != 0 {
			return
		} else {
			v11 = int32(_a69)
			v13 = *(*int64)(unsafe.Add(mBase, _consts[31]))
			*(*int64)(unsafe.Add(mBase, _consts[574])) = v13
			F_replicationCreatePrimaryClientWithHandler(m, int32(0), int32(-1), int32(107))
			mBase = m.M
			v19 = m.ExcPending
			if v19 != 0 {
				return
			} else {
				v20 = int32(_a69)
				v21 = *(*int32)(unsafe.Add(mBase, _consts[133]))
				v22 = *(*int32)(unsafe.Add(mBase, uint32(v21)+104))
				v24 = *(*int64)(unsafe.Add(mBase, _consts[528]))
				*(*int64)(unsafe.Add(mBase, uint32(v22)+104)) = v24
				v31 = int32(*(*uint8)(unsafe.Add(mBase, _consts[511])))
				*(*uint8)(unsafe.Add(mBase, uint32(v22+int32(144)))) = uint8(v31)
				v38 = *(*int64)(unsafe.Add(mBase, _consts[532]))
				*(*int64)(unsafe.Add(mBase, uint32(v22+int32(136)))) = v38
				v45 = *(*int64)(unsafe.Add(mBase, _consts[531]))
				*(*int64)(unsafe.Add(mBase, uint32(v22+int32(128)))) = v45
				v52 = *(*int64)(unsafe.Add(mBase, _consts[530]))
				*(*int64)(unsafe.Add(mBase, uint32(v22+int32(120)))) = v52
				v59 = *(*int64)(unsafe.Add(mBase, _consts[529]))
				*(*int64)(unsafe.Add(mBase, uint32(v22+int32(112)))) = v59
				v62 = *(*int32)(unsafe.Add(mBase, _consts[133]))
				F_unlinkClient(m, v62)
				mBase = m.M
				v64 = m.ExcPending
				if v64 != 0 {
					return
				} else {
					v65 = int32(_a69)
					v67 = *(*int32)(unsafe.Add(mBase, _consts[133]))
					*(*int32)(unsafe.Add(mBase, _consts[134])) = v67
					*(*int32)(unsafe.Add(mBase, _consts[133])) = int32(0)
					return
				}
			}
		}
	}
}
func F_replicationCountAOFAcksByOffset(m *base.Module, l0 int64) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v30 int32
	_ = v30
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
	var v43 int64
	_ = v43
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v59 int32
	_ = v59
	var v63 int32
	_ = v63
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v10 = *(*int32)(unsafe.Add(mBase, _consts[188]))
	v12 = v7 + int32(8)
	v13 = *(*int32)(unsafe.Add(mBase, uint32(v10)))
	*(*int32)(unsafe.Add(mBase, uint32(v12)+4)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v12))) = v13
	v17 = int32(0)
	v19 = v7 + int32(8)
	v21 = *(*int32)(unsafe.Add(mBase, uint32(v19)))
	if v21 == v17 {
	} else {
		v24 = *(*int32)(unsafe.Add(mBase, uint32(v19)+4))
		v30 = *(*int32)(unsafe.Add(mBase, uint32(v21+base.B2i32(v24 == int32(0))<<(uint(int32(2))%32))))
		*(*int32)(unsafe.Add(mBase, uint32(v19))) = v30
	}
	if v21 == int32(0) {
		v63 = v17
	} else {
		v36 = v17
		v37 = v21
		for {
			v38 = *(*int32)(unsafe.Add(mBase, uint32(v37)+8))
			v39 = *(*int32)(unsafe.Add(mBase, uint32(v38)+104))
			v40 = *(*int32)(unsafe.Add(mBase, uint32(v39)))
			if v40 != int32(9) {
				v46 = v36
			} else {
				v43 = *(*int64)(unsafe.Add(mBase, uint32(v39)+72))
				v46 = v36 + base.B2i32(l0 <= v43)
			}
			v48 = v7 + int32(8)
			v50 = *(*int32)(unsafe.Add(mBase, uint32(v48)))
			if v50 == int32(0) {
			} else {
				v53 = *(*int32)(unsafe.Add(mBase, uint32(v48)+4))
				v59 = *(*int32)(unsafe.Add(mBase, uint32(v50+base.B2i32(v53 == int32(0))<<(uint(int32(2))%32))))
				*(*int32)(unsafe.Add(mBase, uint32(v48))) = v59
			}
			if v50 != 0 {
				v36 = v46
				v37 = v50
				continue
			} else {
				break
			}
			break
		}
		v63 = v46
	}
	m.G0 = v7 + int32(16)
	return v63
}
func F_replicationCron(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v16 int32
	_ = v16
	var v22 int64
	_ = v22
	var v23 int32
	_ = v23
	var v24 int64
	_ = v24
	var v27 int64
	_ = v27
	var v30 int32
	_ = v30
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v50 int64
	_ = v50
	var v51 int32
	_ = v51
	var v52 int64
	_ = v52
	var v55 int64
	_ = v55
	var v58 int32
	_ = v58
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
	var v72 int32
	_ = v72
	var v78 int64
	_ = v78
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v81 int64
	_ = v81
	var v84 int64
	_ = v84
	var v87 int32
	_ = v87
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v107 int32
	_ = v107
	var v111 int64
	_ = v111
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v122 int32
	_ = v122
	var v126 int32
	_ = v126
	var v129 int32
	_ = v129
	var v133 int32
	_ = v133
	var v136 int64
	_ = v136
	var v138 int64
	_ = v138
	var v139 int64
	_ = v139
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v148 int32
	_ = v148
	var v152 int32
	_ = v152
	var v153 int64
	_ = v153
	var v155 int64
	_ = v155
	var v160 int64
	_ = v160
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v167 int32
	_ = v167
	var v174 int32
	_ = v174
	var v176 int32
	_ = v176
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v184 int32
	_ = v184
	var v186 int32
	_ = v186
	var v189 int32
	_ = v189
	var v195 int32
	_ = v195
	var v200 int32
	_ = v200
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v211 int32
	_ = v211
	var v214 int32
	_ = v214
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v223 int32
	_ = v223
	var v225 int32
	_ = v225
	var v228 int32
	_ = v228
	var v234 int32
	_ = v234
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v251 int32
	_ = v251
	var v253 int32
	_ = v253
	var v256 int32
	_ = v256
	var v262 int32
	_ = v262
	var v267 int32
	_ = v267
	var v270 int32
	_ = v270
	var v271 int32
	_ = v271
	var v272 int32
	_ = v272
	var v275 int32
	_ = v275
	var v278 int32
	_ = v278
	var v279 int64
	_ = v279
	var v280 int32
	_ = v280
	var v281 int64
	_ = v281
	var v284 int64
	_ = v284
	var v286 int32
	_ = v286
	var v287 int32
	_ = v287
	var v288 int32
	_ = v288
	var v292 int32
	_ = v292
	var v295 int64
	_ = v295
	var v298 int32
	_ = v298
	var v299 int64
	_ = v299
	var v300 int32
	_ = v300
	var v301 int64
	_ = v301
	var v304 int64
	_ = v304
	var v307 int32
	_ = v307
	var v312 int32
	_ = v312
	var v316 int32
	_ = v316
	var v318 int32
	_ = v318
	var v319 int32
	_ = v319
	var v325 int32
	_ = v325
	var v328 int32
	_ = v328
	var v329 int32
	_ = v329
	var v333 int32
	_ = v333
	var v335 int32
	_ = v335
	var v338 int32
	_ = v338
	var v344 int32
	_ = v344
	var v351 int32
	_ = v351
	var v352 int32
	_ = v352
	var v356 int32
	_ = v356
	var v362 int64
	_ = v362
	var v366 int32
	_ = v366
	var v370 int32
	_ = v370
	var v371 int32
	_ = v371
	var v372 int64
	_ = v372
	var v374 int64
	_ = v374
	var v377 int64
	_ = v377
	var v384 int32
	_ = v384
	var v385 int32
	_ = v385
	var v386 int64
	_ = v386
	var v389 int32
	_ = v389
	var v418 int32
	_ = v418
	var v420 int32
	_ = v420
	var v424 int64
	_ = v424
	var v429 int32
	_ = v429
	var v430 int32
	_ = v430
	var v442 int32
	_ = v442
	var v445 int32
	_ = v445
	var v446 int32
	_ = v446
	var v447 int32
	_ = v447
	var v448 int32
	_ = v448
	var v449 int32
	_ = v449
	var v451 int32
	_ = v451
	var v453 int32
	_ = v453
	var v455 int32
	_ = v455
	var v456 int32
	_ = v456
	var v459 int32
	_ = v459
	var v460 int32
	_ = v460
	var v461 int32
	_ = v461
	var v464 int32
	_ = v464
	var v465 int32
	_ = v465
	var v466 int32
	_ = v466
	var v468 int32
	_ = v468
	var v469 int64
	_ = v469
	var v477 int32
	_ = v477
	var v481 int32
	_ = v481
	var v485 int32
	_ = v485
	var v487 int32
	_ = v487
	var v488 int32
	_ = v488
	var v492 int32
	_ = v492
	var v494 int32
	_ = v494
	var v496 int32
	_ = v496
	var v499 int32
	_ = v499
	var v505 int32
	_ = v505
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
	var v518 int32
	_ = v518
	var v519 int64
	_ = v519
	var v520 int64
	_ = v520
	var v523 int64
	_ = v523
	var v526 int32
	_ = v526
	var v528 int32
	_ = v528
	var v530 int32
	_ = v530
	var v533 int32
	_ = v533
	var v539 int32
	_ = v539
	var v543 int32
	_ = v543
	var v551 int32
	_ = v551
	var v553 int64
	_ = v553
	var v564 int32
	_ = v564
	v5 = m.G0
	v7 = v5 - int32(80)
	m.G0 = v7
	F_updateFailoverStatus(m)
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v12 = *(*int32)(unsafe.Add(mBase, _consts[64]))
	if v12 == int32(0) {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v42 = *(*int32)(unsafe.Add(mBase, _consts[64]))
	if v42 == int32(0) {
		goto L11
	} else {
		goto L12
	}
L4:
	;
	v16 = *(*int32)(unsafe.Add(mBase, _consts[518]))
	if base.Ui32(v16+int32(-13)) < base.Ui32(int32(-11)) {
		goto L3
	} else {
		goto L5
	}
L5:
	;
	v22 = F___time(m, int32(0))
	mBase = m.M
	v23 = int32(_a69)
	v24 = *(*int64)(unsafe.Add(mBase, _consts[523]))
	v27 = int64(*(*int32)(unsafe.Add(mBase, _consts[538])))
	if v22-v24 <= v27 {
		goto L3
	} else {
		goto L6
	}
L6:
	;
	v30 = *(*int32)(unsafe.Add(mBase, _consts[15]))
	if int32(3) < v30 {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	v39 = F_cancelReplicationHandshake(m, int32(1))
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
		goto L1
	} else {
		goto L10
	}
L8:
	;
	F__serverLog(m, int32(3), int32(_a1063), int32(0))
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		goto L1
	} else {
		goto L9
	}
L9:
	;
	goto L7
L10:
	;
	goto L3
L11:
	;
	v69 = int32(_a69)
	v70 = *(*int32)(unsafe.Add(mBase, _consts[518]))
	v72 = *(*int32)(unsafe.Add(mBase, _consts[64]))
	if v72 == int32(0) {
		v103 = v70
		goto L19
	} else {
		goto L20
	}
L12:
	;
	v46 = *(*int32)(unsafe.Add(mBase, _consts[518]))
	if v46 != int32(13) {
		goto L11
	} else {
		goto L13
	}
L13:
	;
	v50 = F___time(m, int32(0))
	mBase = m.M
	v51 = int32(_a69)
	v52 = *(*int64)(unsafe.Add(mBase, _consts[523]))
	v55 = int64(*(*int32)(unsafe.Add(mBase, _consts[538])))
	if v50-v52 <= v55 {
		goto L11
	} else {
		goto L14
	}
L14:
	;
	v58 = *(*int32)(unsafe.Add(mBase, _consts[15]))
	if int32(3) < v58 {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	v67 = F_cancelReplicationHandshake(m, int32(1))
	mBase = m.M
	v68 = m.ExcPending
	if v68 != 0 {
		goto L1
	} else {
		goto L18
	}
L16:
	;
	F__serverLog(m, int32(3), int32(_a1064), int32(0))
	mBase = m.M
	v65 = m.ExcPending
	if v65 != 0 {
		goto L1
	} else {
		goto L17
	}
L17:
	;
	goto L15
L18:
	;
	goto L11
L19:
	;
	if v103 != int32(1) {
		goto L28
	} else {
		goto L29
	}
L20:
	;
	if v70 != int32(14) {
		v103 = v70
		goto L19
	} else {
		goto L21
	}
L21:
	;
	v78 = F___time(m, int32(0))
	mBase = m.M
	v79 = int32(_a69)
	v80 = *(*int32)(unsafe.Add(mBase, _consts[133]))
	v81 = *(*int64)(unsafe.Add(mBase, uint32(v80)+88))
	v84 = int64(*(*int32)(unsafe.Add(mBase, _consts[538])))
	if v78-v81 <= v84 {
		goto L22
	} else {
		goto L23
	}
L22:
	;
	v102 = *(*int32)(unsafe.Add(mBase, _consts[518]))
	v103 = v102
	goto L19
L23:
	;
	v87 = *(*int32)(unsafe.Add(mBase, _consts[15]))
	if int32(3) < v87 {
		v97 = v80
		goto L24
	} else {
		goto L25
	}
L24:
	;
	v98 = F_freeClient(m, v97)
	mBase = m.M
	v99 = m.ExcPending
	if v99 != 0 {
		goto L1
	} else {
		goto L27
	}
L25:
	;
	F__serverLog(m, int32(3), int32(_a1065), int32(0))
	mBase = m.M
	v94 = m.ExcPending
	if v94 != 0 {
		goto L1
	} else {
		goto L26
	}
L26:
	;
	v96 = *(*int32)(unsafe.Add(mBase, _consts[133]))
	v97 = v96
	goto L24
L27:
	;
	goto L22
L28:
	;
	v122 = *(*int32)(unsafe.Add(mBase, _consts[64]))
	if v122 == int32(0) {
		goto L34
	} else {
		goto L35
	}
L29:
	;
	v107 = *(*int32)(unsafe.Add(mBase, _consts[15]))
	if int32(2) < v107 {
		goto L30
	} else {
		goto L31
	}
L30:
	;
	v119 = F_connectWithPrimary(m)
	mBase = m.M
	v120 = m.ExcPending
	if v120 != 0 {
		goto L1
	} else {
		goto L33
	}
L31:
	;
	v111 = *(*int64)(unsafe.Add(mBase, _consts[64]))
	*(*int64)(unsafe.Add(mBase, uint32(v7)+32)) = v111
	F__serverLog(m, int32(2), int32(_a1043), v7+int32(32))
	mBase = m.M
	v118 = m.ExcPending
	if v118 != 0 {
		goto L1
	} else {
		goto L32
	}
L32:
	;
	goto L30
L33:
	;
	goto L28
L34:
	;
	v136 = *(*int64)(unsafe.Add(mBase, _consts[597]))
	v138 = int64(*(*int32)(unsafe.Add(mBase, _consts[598])))
	v139 = base.I64_rem_s(v136, v138)
	if v139 != int64(0) {
		goto L39
	} else {
		goto L40
	}
L35:
	;
	v126 = *(*int32)(unsafe.Add(mBase, _consts[133]))
	if v126 == int32(0) {
		goto L34
	} else {
		goto L36
	}
L36:
	;
	v129 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v126)+202)))
	if v129&int32(1) != 0 {
		goto L34
	} else {
		goto L37
	}
L37:
	;
	F_replicationSendAck(m)
	mBase = m.M
	v133 = m.ExcPending
	if v133 != 0 {
		goto L1
	} else {
		goto L38
	}
L38:
	;
	goto L34
L39:
	;
	v176 = *(*int32)(unsafe.Add(mBase, _consts[188]))
	v178 = v7 + int32(56)
	v179 = *(*int32)(unsafe.Add(mBase, uint32(v176)))
	*(*int32)(unsafe.Add(mBase, uint32(v178)+4)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v178))) = v179
	goto L52
L40:
	;
	v143 = *(*int32)(unsafe.Add(mBase, _consts[188]))
	v144 = *(*int32)(unsafe.Add(mBase, uint32(v143)+20))
	if v144 == int32(0) {
		goto L39
	} else {
		goto L41
	}
L41:
	;
	v148 = *(*int32)(unsafe.Add(mBase, _consts[63]))
	if v148 == int32(0) {
		goto L44
	} else {
		goto L45
	}
L42:
	;
	v167 = *(*int32)(unsafe.Add(mBase, _consts[599]))
	*(*int32)(unsafe.Add(mBase, uint32(v7)+52)) = v167
	F_replicationFeedReplicas(m, int32(-1), v7+int32(52), int32(1))
	mBase = m.M
	v174 = m.ExcPending
	if v174 != 0 {
		goto L1
	} else {
		goto L51
	}
L43:
	;
	v164 = F_isPausedActionsWithUpdate(m, int32(16))
	mBase = m.M
	v165 = m.ExcPending
	if v165 != 0 {
		goto L1
	} else {
		goto L49
	}
L44:
	;
	v160 = *(*int64)(unsafe.Add(mBase, _consts[600]))
	if v160 == int64(0) {
		goto L42
	} else {
		goto L48
	}
L45:
	;
	v152 = *(*int32)(unsafe.Add(mBase, _consts[90]))
	v153 = *(*int64)(unsafe.Add(mBase, uint32(v152)+uint32(_consts[113])))
	goto L46
L46:
	;
	v155 = *(*int64)(unsafe.Add(mBase, _consts[600]))
	if v153|v155 != int64(0) {
		goto L43
	} else {
		goto L47
	}
L47:
	;
	goto L42
L48:
	;
	goto L43
L49:
	;
	if v164 != 0 {
		goto L39
	} else {
		goto L50
	}
L50:
	;
	goto L42
L51:
	;
	goto L39
L52:
	;
	v184 = v7 + int32(56)
	v186 = *(*int32)(unsafe.Add(mBase, uint32(v184)))
	if v186 == int32(0) {
		goto L55
	} else {
		goto L56
	}
L53:
	;
	v241 = *(*int32)(unsafe.Add(mBase, _consts[188]))
	v242 = *(*int32)(unsafe.Add(mBase, uint32(v241)+20))
	if v242 != 0 {
		goto L71
	} else {
		goto L72
	}
L54:
	;
	if v186 == int32(0) {
		goto L53
	} else {
		goto L57
	}
L55:
	;
	goto L54
L56:
	;
	v189 = *(*int32)(unsafe.Add(mBase, uint32(v184)+4))
	v195 = *(*int32)(unsafe.Add(mBase, uint32(v186+base.B2i32(v189 == int32(0))<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v184))) = v195
	goto L55
L57:
	;
	v200 = v186
	goto L58
L58:
	;
	v203 = *(*int32)(unsafe.Add(mBase, uint32(v200)+8))
	v204 = *(*int32)(unsafe.Add(mBase, uint32(v203)+104))
	v205 = *(*int32)(unsafe.Add(mBase, uint32(v204)))
	if v205 == int32(6) {
		goto L61
	} else {
		goto L62
	}
L59:
	;
	goto L53
L60:
	;
	v223 = v7 + int32(56)
	v225 = *(*int32)(unsafe.Add(mBase, uint32(v223)))
	if v225 == int32(0) {
		goto L67
	} else {
		goto L68
	}
L61:
	;
	v214 = *(*int32)(unsafe.Add(mBase, uint32(v203)+8))
	v217 = *(*int32)(unsafe.Add(mBase, uint32(v214)))
	v218 = *(*int32)(unsafe.Add(mBase, uint32(v217)+68))
	v219 = m.T0[v218].(func(*base.Module, int32, int32, int32) int32)(m, v214, int32(_a397), int32(1))
	mBase = m.M
	v220 = m.ExcPending
	if v220 != 0 {
		goto L1
	} else {
		goto L65
	}
L62:
	;
	if v205 != int32(7) {
		goto L60
	} else {
		goto L63
	}
L63:
	;
	v211 = *(*int32)(unsafe.Add(mBase, _consts[601]))
	if v211 == int32(2) {
		goto L60
	} else {
		goto L64
	}
L64:
	;
	goto L61
L65:
	;
	goto L60
L66:
	;
	if v225 != 0 {
		v200 = v225
		goto L58
	} else {
		goto L69
	}
L67:
	;
	goto L66
L68:
	;
	v228 = *(*int32)(unsafe.Add(mBase, uint32(v223)+4))
	v234 = *(*int32)(unsafe.Add(mBase, uint32(v225+base.B2i32(v228 == int32(0))<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v223))) = v234
	goto L67
L69:
	;
	goto L59
L70:
	;
	if v356 == int32(0) {
		goto L102
	} else {
		goto L103
	}
L71:
	;
	v245 = v7 + int32(72)
	v246 = *(*int32)(unsafe.Add(mBase, uint32(v241)))
	*(*int32)(unsafe.Add(mBase, uint32(v245)+4)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v245))) = v246
	goto L73
L72:
	;
	v356 = int32(1)
	goto L70
L73:
	;
	v251 = v7 + int32(72)
	v253 = *(*int32)(unsafe.Add(mBase, uint32(v251)))
	if v253 == int32(0) {
		goto L76
	} else {
		goto L77
	}
L74:
	;
	v351 = *(*int32)(unsafe.Add(mBase, _consts[188]))
	v352 = *(*int32)(unsafe.Add(mBase, uint32(v351)+20))
	v356 = base.B2i32(v352 == int32(0))
	goto L70
L75:
	;
	if v253 == int32(0) {
		goto L74
	} else {
		goto L78
	}
L76:
	;
	goto L75
L77:
	;
	v256 = *(*int32)(unsafe.Add(mBase, uint32(v251)+4))
	v262 = *(*int32)(unsafe.Add(mBase, uint32(v253+base.B2i32(v256 == int32(0))<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v251))) = v262
	goto L76
L78:
	;
	v267 = v253
	goto L79
L79:
	;
	v270 = *(*int32)(unsafe.Add(mBase, uint32(v267)+8))
	v271 = *(*int32)(unsafe.Add(mBase, uint32(v270)+104))
	v272 = *(*int32)(unsafe.Add(mBase, uint32(v271)))
	if v272 != int32(9) {
		v287 = v272
		v288 = v271
		goto L85
	} else {
		goto L86
	}
L80:
	;
	goto L74
L81:
	;
	v333 = v7 + int32(72)
	v335 = *(*int32)(unsafe.Add(mBase, uint32(v333)))
	if v335 == int32(0) {
		goto L99
	} else {
		goto L100
	}
L82:
	;
	v328 = F_freeClient(m, v270)
	mBase = m.M
	v329 = m.ExcPending
	if v329 != 0 {
		goto L1
	} else {
		goto L97
	}
L83:
	;
	v318 = F_replicationGetReplicaName(m, v270)
	mBase = m.M
	v319 = m.ExcPending
	if v319 != 0 {
		goto L1
	} else {
		goto L95
	}
L84:
	;
	v312 = *(*int32)(unsafe.Add(mBase, _consts[15]))
	if int32(3) < v312 {
		goto L82
	} else {
		goto L94
	}
L85:
	;
	if v287 != int32(7) {
		goto L81
	} else {
		goto L89
	}
L86:
	;
	v275 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v270)+202)))
	if v275&int32(1) != 0 {
		goto L81
	} else {
		goto L87
	}
L87:
	;
	v278 = int32(_a69)
	v279 = *(*int64)(unsafe.Add(mBase, _consts[37]))
	v280 = *(*int32)(unsafe.Add(mBase, uint32(v270)+104))
	v281 = *(*int64)(unsafe.Add(mBase, uint32(v280)+80))
	v284 = int64(*(*int32)(unsafe.Add(mBase, _consts[538])))
	if v284 < v279-v281 {
		goto L84
	} else {
		goto L88
	}
L88:
	;
	v286 = *(*int32)(unsafe.Add(mBase, uint32(v280)))
	v287 = v286
	v288 = v280
	goto L85
L89:
	;
	v292 = *(*int32)(unsafe.Add(mBase, _consts[601]))
	if v292 != int32(2) {
		goto L81
	} else {
		goto L90
	}
L90:
	;
	v295 = *(*int64)(unsafe.Add(mBase, uint32(v288)+88))
	if v295 == int64(0) {
		goto L81
	} else {
		goto L91
	}
L91:
	;
	v298 = int32(_a69)
	v299 = *(*int64)(unsafe.Add(mBase, _consts[37]))
	v300 = *(*int32)(unsafe.Add(mBase, uint32(v270)+104))
	v301 = *(*int64)(unsafe.Add(mBase, uint32(v300)+88))
	v304 = int64(*(*int32)(unsafe.Add(mBase, _consts[538])))
	if v299-v301 <= v304 {
		goto L81
	} else {
		goto L92
	}
L92:
	;
	v307 = *(*int32)(unsafe.Add(mBase, _consts[15]))
	if int32(3) < v307 {
		goto L82
	} else {
		goto L93
	}
L93:
	;
	v316 = int32(_a1066)
	goto L83
L94:
	;
	v316 = int32(_a1067)
	goto L83
L95:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v7)+16)) = v318
	F__serverLog(m, int32(3), v316, v7+int32(16))
	mBase = m.M
	v325 = m.ExcPending
	if v325 != 0 {
		goto L1
	} else {
		goto L96
	}
L96:
	;
	goto L82
L97:
	;
	goto L81
L98:
	;
	if v335 != 0 {
		v267 = v335
		goto L79
	} else {
		goto L101
	}
L99:
	;
	goto L98
L100:
	;
	v338 = *(*int32)(unsafe.Add(mBase, uint32(v333)+4))
	v344 = *(*int32)(unsafe.Add(mBase, uint32(v335+base.B2i32(v338 == int32(0))<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v333))) = v344
	goto L99
L101:
	;
	goto L80
L102:
	;
	v430 = int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(v7)+72)) = v430
	*(*int32)(unsafe.Add(mBase, uint32(v7)+68)) = v430
	*(*int32)(unsafe.Add(mBase, uint32(v7)+64)) = v430
	v442 = F_shouldStartChildReplication(m, v7+int32(72), v7+int32(68), v7+int32(64))
	mBase = m.M
	if v442 == int32(0) {
		goto L112
	} else {
		goto L113
	}
L103:
	;
	v362 = *(*int64)(unsafe.Add(mBase, _consts[602]))
	if v362 == int64(0) {
		goto L102
	} else {
		goto L104
	}
L104:
	;
	v366 = *(*int32)(unsafe.Add(mBase, _consts[448]))
	if v366 == int32(0) {
		goto L102
	} else {
		goto L105
	}
L105:
	;
	v370 = *(*int32)(unsafe.Add(mBase, _consts[64]))
	if v370 != 0 {
		goto L102
	} else {
		goto L106
	}
L106:
	;
	v371 = int32(_a69)
	v372 = *(*int64)(unsafe.Add(mBase, _consts[37]))
	v374 = *(*int64)(unsafe.Add(mBase, _consts[603]))
	v377 = *(*int64)(unsafe.Add(mBase, _consts[602]))
	if v372-v374 <= v377 {
		goto L102
	} else {
		goto L107
	}
L107:
	;
	F_getRandomHexChars(m, int32(_a920), int32(40))
	mBase = m.M
	v384 = m.ExcPending
	if v384 != 0 {
		goto L1
	} else {
		goto L108
	}
L108:
	;
	v385 = int32(_a69)
	v386 = int64(3472328296227680304)
	*(*int64)(unsafe.Add(mBase, _consts[512])) = v386
	v389 = int32(0)
	*(*uint8)(unsafe.Add(mBase, _consts[511])) = uint8(v389)
	*(*int64)(unsafe.Add(mBase, _consts[517])) = v386
	*(*int64)(unsafe.Add(mBase, _consts[516])) = v386
	*(*int64)(unsafe.Add(mBase, _consts[515])) = v386
	*(*int64)(unsafe.Add(mBase, _consts[514])) = v386
	*(*int64)(unsafe.Add(mBase, _consts[508])) = int64(-1)
	*(*uint8)(unsafe.Add(mBase, _consts[513])) = uint8(v389)
	F_freeReplicationBacklog(m)
	mBase = m.M
	v418 = m.ExcPending
	if v418 != 0 {
		goto L1
	} else {
		goto L109
	}
L109:
	;
	v420 = *(*int32)(unsafe.Add(mBase, _consts[15]))
	if int32(2) < v420 {
		goto L102
	} else {
		goto L110
	}
L110:
	;
	v424 = *(*int64)(unsafe.Add(mBase, _consts[602]))
	*(*uint32)(unsafe.Add(mBase, uint32(v7))) = uint32(v424)
	F__serverLog(m, int32(2), int32(_a1068), v7)
	mBase = m.M
	v429 = m.ExcPending
	if v429 != 0 {
		goto L1
	} else {
		goto L111
	}
L111:
	;
	goto L102
L112:
	;
	F_handleBioThreadFinishedRDBDownload(m)
	mBase = m.M
	v451 = m.ExcPending
	if v451 != 0 {
		goto L1
	} else {
		goto L115
	}
L113:
	;
	v445 = *(*int32)(unsafe.Add(mBase, uint32(v7)+72))
	v446 = *(*int32)(unsafe.Add(mBase, uint32(v7)+68))
	v447 = *(*int32)(unsafe.Add(mBase, uint32(v7)+64))
	v448 = F_startBgsaveForReplication(m, v445, v446, v447)
	mBase = m.M
	v449 = m.ExcPending
	if v449 != 0 {
		goto L1
	} else {
		goto L114
	}
L114:
	;
	goto L112
L115:
	;
	F_removeRDBUsedToSyncReplicas(m)
	mBase = m.M
	v453 = m.ExcPending
	if v453 != 0 {
		goto L1
	} else {
		goto L116
	}
L116:
	;
	v455 = *(*int32)(unsafe.Add(mBase, _consts[314]))
	v456 = *(*int32)(unsafe.Add(mBase, uint32(v455)+20))
	if v456 == int32(0) {
		goto L118
	} else {
		goto L119
	}
L117:
	;
	F__serverAssert(m, int32(_a1069), int32(_a913), int32(5435))
	mBase = m.M
	v564 = m.ExcPending
	if v564 != 0 {
		goto L1
	} else {
		goto L140
	}
L118:
	;
	v477 = *(*int32)(unsafe.Add(mBase, _consts[186]))
	if v477 == int32(0) {
		goto L123
	} else {
		goto L124
	}
L119:
	;
	v459 = *(*int32)(unsafe.Add(mBase, uint32(v455)))
	v460 = *(*int32)(unsafe.Add(mBase, uint32(v459)+8))
	v461 = *(*int32)(unsafe.Add(mBase, uint32(v460)))
	if v461 < int32(1) {
		goto L117
	} else {
		goto L120
	}
L120:
	;
	v464 = int32(_a69)
	v465 = *(*int32)(unsafe.Add(mBase, _consts[188]))
	v466 = *(*int32)(unsafe.Add(mBase, uint32(v465)+20))
	v468 = *(*int32)(unsafe.Add(mBase, _consts[505]))
	v469 = *(*int64)(unsafe.Add(mBase, uint32(v468)+8))
	goto L121
L121:
	;
	if v466+base.I32_wrap_i64(v469)+int32(1) < v461 {
		goto L117
	} else {
		goto L122
	}
L122:
	;
	goto L118
L123:
	;
	v551 = int32(0)
	v553 = *(*int64)(unsafe.Add(mBase, _consts[597]))
	*(*int64)(unsafe.Add(mBase, _consts[597])) = v553 + int64(1)
	m.G0 = v7 + int32(80)
	return
L124:
	;
	v481 = *(*int32)(unsafe.Add(mBase, _consts[187]))
	if v481 == int32(0) {
		goto L123
	} else {
		goto L125
	}
L125:
	;
	v485 = *(*int32)(unsafe.Add(mBase, _consts[188]))
	v487 = v7 + int32(72)
	v488 = *(*int32)(unsafe.Add(mBase, uint32(v485)))
	*(*int32)(unsafe.Add(mBase, uint32(v487)+4)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v487))) = v488
	goto L126
L126:
	;
	v492 = int32(0)
	v494 = v7 + int32(72)
	v496 = *(*int32)(unsafe.Add(mBase, uint32(v494)))
	if v496 == v492 {
		goto L129
	} else {
		goto L130
	}
L127:
	;
	*(*int32)(unsafe.Add(mBase, _consts[189])) = v543
	goto L123
L128:
	;
	if v496 == int32(0) {
		v543 = v492
		goto L127
	} else {
		goto L131
	}
L129:
	;
	goto L128
L130:
	;
	v499 = *(*int32)(unsafe.Add(mBase, uint32(v494)+4))
	v505 = *(*int32)(unsafe.Add(mBase, uint32(v496+base.B2i32(v499 == int32(0))<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v494))) = v505
	goto L129
L131:
	;
	v510 = v496
	v511 = v492
	goto L132
L132:
	;
	v513 = *(*int32)(unsafe.Add(mBase, uint32(v510)+8))
	v514 = *(*int32)(unsafe.Add(mBase, uint32(v513)+104))
	v515 = *(*int32)(unsafe.Add(mBase, uint32(v514)))
	if v515 != int32(9) {
		v526 = v511
		goto L134
	} else {
		goto L135
	}
L133:
	;
	v543 = v526
	goto L127
L134:
	;
	v528 = v7 + int32(72)
	v530 = *(*int32)(unsafe.Add(mBase, uint32(v528)))
	if v530 == int32(0) {
		goto L137
	} else {
		goto L138
	}
L135:
	;
	v518 = int32(_a69)
	v519 = *(*int64)(unsafe.Add(mBase, _consts[37]))
	v520 = *(*int64)(unsafe.Add(mBase, uint32(v514)+80))
	v523 = int64(*(*int32)(unsafe.Add(mBase, _consts[187])))
	v526 = v511 + base.B2i32(v519-v520 <= v523)
	goto L134
L136:
	;
	if v530 != 0 {
		v510 = v530
		v511 = v526
		goto L132
	} else {
		goto L139
	}
L137:
	;
	goto L136
L138:
	;
	v533 = *(*int32)(unsafe.Add(mBase, uint32(v528)+4))
	v539 = *(*int32)(unsafe.Add(mBase, uint32(v530+base.B2i32(v533 == int32(0))<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v528))) = v539
	goto L137
L139:
	;
	goto L133
L140:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_replicationFeedMonitors(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) {
	mBase := m.M
	_ = mBase
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v19 int32
	_ = v19
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v30 int32
	_ = v30
	var v31 int64
	_ = v31
	var v33 int32
	_ = v33
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v57 int32
	_ = v57
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v81 int32
	_ = v81
	var v85 int32
	_ = v85
	var v94 int32
	_ = v94
	var v103 int32
	_ = v103
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v115 int32
	_ = v115
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v124 int32
	_ = v124
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
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v162 int32
	_ = v162
	var v169 int32
	_ = v169
	var v172 int32
	_ = v172
	var v175 int32
	_ = v175
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v194 int32
	_ = v194
	var v202 int32
	_ = v202
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v222 int32
	_ = v222
	var v224 int32
	_ = v224
	var v227 int32
	_ = v227
	var v233 int32
	_ = v233
	var v239 int32
	_ = v239
	var v249 int32
	_ = v249
	var v251 int32
	_ = v251
	var v252 int32
	_ = v252
	var v253 int32
	_ = v253
	var v255 int32
	_ = v255
	var v257 int32
	_ = v257
	var v260 int32
	_ = v260
	var v266 int32
	_ = v266
	var v281 int32
	_ = v281
	v13 = m.G0
	v15 = v13 - int32(96)
	m.G0 = v15
	if l1 == int32(0) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	m.G0 = v15 + int32(96)
	return
L2:
	;
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	if v19 == int32(0) {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v23 = *(*int32)(unsafe.Add(mBase, _consts[116]))
	if v23 != 0 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v25 = F_sdsnew(m, int32(_a788))
	mBase = m.M
	v26 = m.ExcPending
	if v26 != 0 {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	return
L6:
	;
	v30 = F___gettimeofday(m, v15+int32(72), int32(0))
	mBase = m.M
	v31 = *(*int64)(unsafe.Add(mBase, uint32(v15)+72))
	*(*uint32)(unsafe.Add(mBase, uint32(v15)+64)) = uint32(v31)
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v15)+80))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+68)) = v33
	v38 = F_sdscatprintf(m, v25, int32(_a914), v15+int32(64))
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L5
	} else {
		goto L7
	}
L7:
	;
	v40 = *(*int32)(unsafe.Add(mBase, uint32(l0)+200))
	if v40&int32(256) == int32(0) {
		goto L9
	} else {
		goto L10
	}
L8:
	;
	if l4 < int32(1) {
		v202 = v73
		goto L17
	} else {
		goto L18
	}
L9:
	;
	if v40&int32(2048) == int32(0) {
		goto L12
	} else {
		goto L13
	}
L10:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+48)) = l2
	v49 = F_sdscatprintf(m, v38, int32(_a915), v15+int32(48))
	mBase = m.M
	v50 = m.ExcPending
	if v50 != 0 {
		goto L5
	} else {
		goto L11
	}
L11:
	;
	v73 = v49
	goto L8
L12:
	;
	v64 = F_getClientPeerId(m, l0)
	mBase = m.M
	v65 = m.ExcPending
	if v65 != 0 {
		goto L5
	} else {
		goto L15
	}
L13:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+32)) = l2
	v57 = *(*int32)(unsafe.Add(mBase, _consts[507]))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+36)) = v57
	v62 = F_sdscatprintf(m, v38, int32(_a916), v15+int32(32))
	mBase = m.M
	v63 = m.ExcPending
	if v63 != 0 {
		goto L5
	} else {
		goto L14
	}
L14:
	;
	v73 = v62
	goto L8
L15:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+20)) = v64
	*(*int32)(unsafe.Add(mBase, uint32(v15)+16)) = l2
	v71 = F_sdscatprintf(m, v38, int32(_a917), v15+int32(16))
	mBase = m.M
	v72 = m.ExcPending
	if v72 != 0 {
		goto L5
	} else {
		goto L16
	}
L16:
	;
	v73 = v71
	goto L8
L17:
	;
	v211 = F_sdscatlen(m, v202, int32(_a727), int32(2))
	mBase = m.M
	v212 = m.ExcPending
	if v212 != 0 {
		goto L5
	} else {
		goto L54
	}
L18:
	;
	v81 = int32(0)
	v85 = v73
	goto L19
L19:
	;
	if int32(1) <= v81 {
		goto L24
	} else {
		goto L25
	}
L20:
	;
	v202 = v192
	goto L17
L21:
	;
	if v81 == l4+int32(-1) {
		v192 = v182
		goto L50
	} else {
		goto L51
	}
L22:
	;
	v145 = l3 + v81<<(uint(int32(2))%32)
	v146 = *(*int32)(unsafe.Add(mBase, uint32(v145)))
	v147 = *(*int32)(unsafe.Add(mBase, uint32(v146)))
	v148 = F_objectGetVal(m, v146)
	mBase = m.M
	if v147&int32(240) != int32(16) {
		goto L40
	} else {
		goto L41
	}
L23:
	;
	if v103 == int32(0) {
		goto L22
	} else {
		goto L28
	}
L24:
	;
	v94 = *(*int32)(unsafe.Add(mBase, uint32(l0)+196))
	if base.Ui32(v81) < base.Ui32(int32(32)) {
		goto L26
	} else {
		goto L27
	}
L25:
	;
	v103 = int32(0)
	goto L23
L26:
	;
	v103 = int32(base.Ui32(v94)>>(uint(v81)%32)) & int32(1)
	goto L23
L27:
	;
	v103 = v94 & int32(1)
	goto L23
L28:
	;
	v106 = int32(_a918)
	v107 = *(*int32)(unsafe.Add(mBase, _consts[164]))
	v108 = F_objectGetVal(m, v107)
	mBase = m.M
	v111 = *(*int32)(unsafe.Add(mBase, _consts[164]))
	v112 = F_objectGetVal(m, v111)
	mBase = m.M
	v115 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v112+int32(-1)))))
	switch v115 & int32(7) {
	case 0:
		goto L34
	case 1:
		goto L33
	case 2:
		goto L32
	case 3:
		goto L31
	case 4:
		goto L30
	default:
		v140 = int32(0)
		goto L29
	}
L29:
	;
	v141 = F_sdscatrepr(m, v85, v108, v140)
	mBase = m.M
	v142 = m.ExcPending
	if v142 != 0 {
		goto L5
	} else {
		goto L39
	}
L30:
	;
	v139 = *(*int32)(unsafe.Add(mBase, uint32(v112+int32(-17))))
	v140 = v139
	goto L29
L31:
	;
	v134 = *(*int32)(unsafe.Add(mBase, uint32(v112+int32(-9))))
	v135 = F_sdscatrepr(m, v85, v108, v134)
	mBase = m.M
	v136 = m.ExcPending
	if v136 != 0 {
		goto L5
	} else {
		goto L38
	}
L32:
	;
	v129 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v112+int32(-5)))))
	v130 = F_sdscatrepr(m, v85, v108, v129)
	mBase = m.M
	v131 = m.ExcPending
	if v131 != 0 {
		goto L5
	} else {
		goto L37
	}
L33:
	;
	v124 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v112+int32(-3)))))
	v125 = F_sdscatrepr(m, v85, v108, v124)
	mBase = m.M
	v126 = m.ExcPending
	if v126 != 0 {
		goto L5
	} else {
		goto L36
	}
L34:
	;
	v120 = F_sdscatrepr(m, v85, v108, int32(base.Ui32(v115)>>(uint(int32(3))%32)))
	mBase = m.M
	v121 = m.ExcPending
	if v121 != 0 {
		goto L5
	} else {
		goto L35
	}
L35:
	;
	v182 = v120
	goto L21
L36:
	;
	v182 = v125
	goto L21
L37:
	;
	v182 = v130
	goto L21
L38:
	;
	v182 = v135
	goto L21
L39:
	;
	v182 = v141
	goto L21
L40:
	;
	v158 = *(*int32)(unsafe.Add(mBase, uint32(v145)))
	v159 = F_objectGetVal(m, v158)
	mBase = m.M
	v162 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v159+int32(-1)))))
	switch v162 & int32(7) {
	case 0:
		goto L48
	case 1:
		goto L47
	case 2:
		goto L46
	case 3:
		goto L45
	case 4:
		goto L44
	default:
		v179 = int32(0)
		goto L43
	}
L41:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15))) = v148
	v155 = F_sdscatprintf(m, v85, int32(_a919), v15)
	mBase = m.M
	v156 = m.ExcPending
	if v156 != 0 {
		goto L5
	} else {
		goto L42
	}
L42:
	;
	v182 = v155
	goto L21
L43:
	;
	v180 = F_sdscatrepr(m, v85, v148, v179)
	mBase = m.M
	v181 = m.ExcPending
	if v181 != 0 {
		goto L5
	} else {
		goto L49
	}
L44:
	;
	v178 = *(*int32)(unsafe.Add(mBase, uint32(v159+int32(-17))))
	v179 = v178
	goto L43
L45:
	;
	v175 = *(*int32)(unsafe.Add(mBase, uint32(v159+int32(-9))))
	v179 = v175
	goto L43
L46:
	;
	v172 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v159+int32(-5)))))
	v179 = v172
	goto L43
L47:
	;
	v169 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v159+int32(-3)))))
	v179 = v169
	goto L43
L48:
	;
	v179 = int32(base.Ui32(v162) >> (uint(int32(3)) % 32))
	goto L43
L49:
	;
	v182 = v180
	goto L21
L50:
	;
	v194 = v81 + int32(1)
	if v194 != l4 {
		v81 = v194
		v85 = v192
		goto L19
	} else {
		goto L53
	}
L51:
	;
	v190 = F_sdscatlen(m, v182, int32(_a10), int32(1))
	mBase = m.M
	v191 = m.ExcPending
	if v191 != 0 {
		goto L5
	} else {
		goto L52
	}
L52:
	;
	v192 = v190
	goto L50
L53:
	;
	goto L20
L54:
	;
	v213 = F_createObject(m, int32(0), v211)
	mBase = m.M
	v214 = m.ExcPending
	if v214 != 0 {
		goto L5
	} else {
		goto L55
	}
L55:
	;
	v216 = v15 + int32(88)
	v217 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	*(*int32)(unsafe.Add(mBase, uint32(v216)+4)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v216))) = v217
	goto L56
L56:
	;
	v222 = v15 + int32(88)
	v224 = *(*int32)(unsafe.Add(mBase, uint32(v222)))
	if v224 == int32(0) {
		goto L59
	} else {
		goto L60
	}
L57:
	;
	F_decrRefCount(m, v213)
	mBase = m.M
	v281 = m.ExcPending
	if v281 != 0 {
		goto L5
	} else {
		goto L70
	}
L58:
	;
	if v224 == int32(0) {
		goto L57
	} else {
		goto L61
	}
L59:
	;
	goto L58
L60:
	;
	v227 = *(*int32)(unsafe.Add(mBase, uint32(v222)+4))
	v233 = *(*int32)(unsafe.Add(mBase, uint32(v224+base.B2i32(v227 == int32(0))<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v222))) = v233
	goto L59
L61:
	;
	v239 = v224
	goto L62
L62:
	;
	v249 = *(*int32)(unsafe.Add(mBase, uint32(v239)+8))
	F_addReply(m, v249, v213)
	mBase = m.M
	v251 = m.ExcPending
	if v251 != 0 {
		goto L5
	} else {
		goto L64
	}
L63:
	;
	goto L57
L64:
	;
	v252 = F_updateClientMemUsageAndBucket(m, v249)
	mBase = m.M
	v253 = m.ExcPending
	if v253 != 0 {
		goto L5
	} else {
		goto L65
	}
L65:
	;
	v255 = v15 + int32(88)
	v257 = *(*int32)(unsafe.Add(mBase, uint32(v255)))
	if v257 == int32(0) {
		goto L67
	} else {
		goto L68
	}
L66:
	;
	if v257 != 0 {
		v239 = v257
		goto L62
	} else {
		goto L69
	}
L67:
	;
	goto L66
L68:
	;
	v260 = *(*int32)(unsafe.Add(mBase, uint32(v255)+4))
	v266 = *(*int32)(unsafe.Add(mBase, uint32(v257+base.B2i32(v260 == int32(0))<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v255))) = v266
	goto L67
L69:
	;
	goto L63
L70:
	;
	goto L1
}
func F_replicationFeedStreamFromPrimaryStream(m *base.Module, l0 int32, l1 int32) {
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
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v39 int32
	_ = v39
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v70 int32
	_ = v70
	var v78 int32
	_ = v78
	var v91 int32
	_ = v91
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v10 = int32(_a69)
	v11 = *(*int32)(unsafe.Add(mBase, _consts[448]))
	v13 = *(*int32)(unsafe.Add(mBase, _consts[188]))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(v13)+20))
	if v14 == int32(0) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	F__serverAssert(m, int32(_a912), int32(_a913), int32(685))
	mBase = m.M
	v91 = m.ExcPending
	if v91 != 0 {
		goto L18
	} else {
		goto L25
	}
L2:
	;
	if v11 == int32(0) {
		goto L5
	} else {
		goto L6
	}
L3:
	;
	if v11 == int32(0) {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	goto L2
L5:
	;
	m.G0 = v8 + int32(16)
	return
L6:
	;
	v22 = v8 + int32(8)
	v23 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
	*(*int32)(unsafe.Add(mBase, uint32(v22)+4)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v22))) = v23
	goto L7
L7:
	;
	v28 = v8 + int32(8)
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v28)))
	if v30 == int32(0) {
		goto L10
	} else {
		goto L11
	}
L8:
	;
	F_feedReplicationBuffer(m, l0, l1)
	mBase = m.M
	v78 = m.ExcPending
	if v78 != 0 {
		goto L18
	} else {
		goto L24
	}
L9:
	;
	if v30 == int32(0) {
		goto L8
	} else {
		goto L12
	}
L10:
	;
	goto L9
L11:
	;
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v28)+4))
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v30+base.B2i32(v33 == int32(0))<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v28))) = v39
	goto L10
L12:
	;
	v46 = v30
	goto L13
L13:
	;
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v46)+8))
	v49 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v48)+205)))
	if v49&int32(32) != 0 {
		goto L15
	} else {
		goto L16
	}
L14:
	;
	goto L8
L15:
	;
	v59 = v8 + int32(8)
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v59)))
	if v61 == int32(0) {
		goto L21
	} else {
		goto L22
	}
L16:
	;
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v48)+104))
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v52)))
	if v53 == int32(6) {
		goto L15
	} else {
		goto L17
	}
L17:
	;
	v56 = F_prepareClientToWrite(m, v48)
	mBase = m.M
	v57 = m.ExcPending
	if v57 != 0 {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	return
L19:
	;
	goto L15
L20:
	;
	if v61 != 0 {
		v46 = v61
		goto L13
	} else {
		goto L23
	}
L21:
	;
	goto L20
L22:
	;
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v59)+4))
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v61+base.B2i32(v64 == int32(0))<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v59))) = v70
	goto L21
L23:
	;
	goto L14
L24:
	;
	goto L5
L25:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_replicationGetReplicaName(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
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
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v51 int32
	_ = v51
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v72 int64
	_ = v72
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	v2 = int32(0)
	v5 = m.G0
	v7 = v5 - int32(96)
	m.G0 = v7
	*(*uint8)(unsafe.Add(mBase, _consts[502])) = uint8(v2)
	*(*uint8)(unsafe.Add(mBase, uint32(v7)+48)) = uint8(v2)
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
	v15 = *(*int32)(unsafe.Add(mBase, uint32(v14)+152))
	if v15 != 0 {
		v36 = v14
		v37 = v15
		if v37 != 0 {
			v40 = v37
		} else {
			v40 = v7 + int32(48)
		}
		v41 = *(*int32)(unsafe.Add(mBase, uint32(v36)+148))
		if v41 == int32(0) {
			*(*int32)(unsafe.Add(mBase, uint32(v7)+16)) = v40
			v69 = F_snprintf(m, int32(_a905), int32(288), int32(_a906), v7+int32(16))
			mBase = m.M
			v70 = m.ExcPending
			if v70 != 0 {
				return int32(0)
			} else {
				m.G0 = v7 + int32(96)
				return int32(_a905)
			}
		} else {
			v44 = int32(58)
			v45 = F___strchrnul(m, v40, v44)
			mBase = m.M
			v47 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v45))))
			if v47 == v44 {
				v51 = v45
			} else {
				v51 = int32(0)
			}
			*(*int32)(unsafe.Add(mBase, uint32(v7)+36)) = v41
			*(*int32)(unsafe.Add(mBase, uint32(v7)+32)) = v40
			if v51 != 0 {
				v58 = int32(_a798)
			} else {
				v58 = int32(_a799)
			}
			v61 = F_snprintf(m, int32(_a905), int32(288), v58, v7+int32(32))
			mBase = m.M
			v62 = m.ExcPending
			if v62 != 0 {
				return int32(0)
			} else {
				m.G0 = v7 + int32(96)
				return int32(_a905)
			}
		}
	} else {
		v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		if v16 == int32(0) {
			v72 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
			*(*int64)(unsafe.Add(mBase, uint32(v7))) = v72
			v77 = F_snprintf(m, int32(_a905), int32(288), int32(_a907), v7)
			mBase = m.M
			v78 = m.ExcPending
			if v78 != 0 {
				return int32(0)
			} else {
				m.G0 = v7 + int32(96)
				return int32(_a905)
			}
		} else {
			v19 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
			v20 = *(*int32)(unsafe.Add(mBase, uint32(v19)+24))
			if v20 == int32(0) {
				v72 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
				*(*int64)(unsafe.Add(mBase, uint32(v7))) = v72
				v77 = F_snprintf(m, int32(_a905), int32(288), int32(_a907), v7)
				mBase = m.M
				v78 = m.ExcPending
				if v78 != 0 {
					return int32(0)
				} else {
					m.G0 = v7 + int32(96)
					return int32(_a905)
				}
			} else {
				v28 = m.T0[v20].(func(*base.Module, int32, int32, int32, int32, int32) int32)(m, v16, v7+int32(48), int32(46), int32(0), int32(1))
				mBase = m.M
				v31 = m.ExcPending
				if v31 != 0 {
					return int32(0)
				} else {
					if v28 == int32(-1) {
						v72 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
						*(*int64)(unsafe.Add(mBase, uint32(v7))) = v72
						v77 = F_snprintf(m, int32(_a905), int32(288), int32(_a907), v7)
						mBase = m.M
						v78 = m.ExcPending
						if v78 != 0 {
							return int32(0)
						} else {
							m.G0 = v7 + int32(96)
							return int32(_a905)
						}
					} else {
						v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
						v35 = *(*int32)(unsafe.Add(mBase, uint32(v34)+152))
						v36 = v34
						v37 = v35
						if v37 != 0 {
							v40 = v37
						} else {
							v40 = v7 + int32(48)
						}
						v41 = *(*int32)(unsafe.Add(mBase, uint32(v36)+148))
						if v41 == int32(0) {
							*(*int32)(unsafe.Add(mBase, uint32(v7)+16)) = v40
							v69 = F_snprintf(m, int32(_a905), int32(288), int32(_a906), v7+int32(16))
							mBase = m.M
							v70 = m.ExcPending
							if v70 != 0 {
								return int32(0)
							} else {
								m.G0 = v7 + int32(96)
								return int32(_a905)
							}
						} else {
							v44 = int32(58)
							v45 = F___strchrnul(m, v40, v44)
							mBase = m.M
							v47 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v45))))
							if v47 == v44 {
								v51 = v45
							} else {
								v51 = int32(0)
							}
							*(*int32)(unsafe.Add(mBase, uint32(v7)+36)) = v41
							*(*int32)(unsafe.Add(mBase, uint32(v7)+32)) = v40
							if v51 != 0 {
								v58 = int32(_a798)
							} else {
								v58 = int32(_a799)
							}
							v61 = F_snprintf(m, int32(_a905), int32(288), v58, v7+int32(32))
							mBase = m.M
							v62 = m.ExcPending
							if v62 != 0 {
								return int32(0)
							} else {
								m.G0 = v7 + int32(96)
								return int32(_a905)
							}
						}
					}
				}
			}
		}
	}
}
func F_replicationHandlePrimaryDisconnection(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v35 int64
	_ = v35
	var v39 int32
	_ = v39
	var v44 int64
	_ = v44
	var v48 int32
	_ = v48
	var v52 int32
	_ = v52
	var v56 int32
	_ = v56
	var v61 int32
	_ = v61
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	v3 = m.G0
	v5 = v3 - int32(16)
	m.G0 = v5
	v8 = *(*int32)(unsafe.Add(mBase, _consts[518]))
	if v8 == int32(14) {
		F_moduleFireServerEvent(m, int64(7), int32(1), int32(0))
		mBase = m.M
		v18 = m.ExcPending
		if v18 != 0 {
			return
		} else {
			v19 = int32(_a69)
			*(*int32)(unsafe.Add(mBase, _consts[133])) = int32(0)
			v23 = *(*int32)(unsafe.Add(mBase, _consts[518]))
			if v23 != int32(14) {
				if v23 != int32(14) {
				} else {
					v39 = int32(_a69)
					*(*int32)(unsafe.Add(mBase, _consts[518])) = int32(0)
					v44 = *(*int64)(unsafe.Add(mBase, _consts[37]))
					*(*int64)(unsafe.Add(mBase, _consts[519])) = v44
				}
			} else {
				v27 = *(*int32)(unsafe.Add(mBase, _consts[64]))
				if v27 == int32(0) {
					if v23 != int32(14) {
					} else {
						v39 = int32(_a69)
						*(*int32)(unsafe.Add(mBase, _consts[518])) = int32(0)
						v44 = *(*int64)(unsafe.Add(mBase, _consts[37]))
						*(*int64)(unsafe.Add(mBase, _consts[519])) = v44
					}
				} else {
					v30 = int32(_a69)
					*(*int32)(unsafe.Add(mBase, _consts[518])) = int32(1)
					v35 = *(*int64)(unsafe.Add(mBase, _consts[37]))
					*(*int64)(unsafe.Add(mBase, _consts[519])) = v35
				}
			}
			v48 = *(*int32)(unsafe.Add(mBase, _consts[518]))
			if v48 != int32(1) {
				m.G0 = v5 + int32(16)
				return
			} else {
				v52 = *(*int32)(unsafe.Add(mBase, _consts[64]))
				if v52 == int32(0) {
					m.G0 = v5 + int32(16)
					return
				} else {
					v56 = *(*int32)(unsafe.Add(mBase, _consts[15]))
					if int32(2) < v56 {
						v67 = F_connectWithPrimary(m)
						mBase = m.M
						v68 = m.ExcPending
						if v68 != 0 {
							return
						} else {
							m.G0 = v5 + int32(16)
							return
						}
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v5))) = v52
						v61 = *(*int32)(unsafe.Add(mBase, _consts[520]))
						*(*int32)(unsafe.Add(mBase, uint32(v5)+4)) = v61
						F__serverLog(m, int32(2), int32(_a930), v5)
						mBase = m.M
						v66 = m.ExcPending
						if v66 != 0 {
							return
						} else {
							v67 = F_connectWithPrimary(m)
							mBase = m.M
							v68 = m.ExcPending
							if v68 != 0 {
								return
							} else {
								m.G0 = v5 + int32(16)
								return
							}
						}
					}
				}
			}
		}
	} else {
		*(*int32)(unsafe.Add(mBase, _consts[133])) = int32(0)
		v48 = *(*int32)(unsafe.Add(mBase, _consts[518]))
		if v48 != int32(1) {
			m.G0 = v5 + int32(16)
			return
		} else {
			v52 = *(*int32)(unsafe.Add(mBase, _consts[64]))
			if v52 == int32(0) {
				m.G0 = v5 + int32(16)
				return
			} else {
				v56 = *(*int32)(unsafe.Add(mBase, _consts[15]))
				if int32(2) < v56 {
					v67 = F_connectWithPrimary(m)
					mBase = m.M
					v68 = m.ExcPending
					if v68 != 0 {
						return
					} else {
						m.G0 = v5 + int32(16)
						return
					}
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v5))) = v52
					v61 = *(*int32)(unsafe.Add(mBase, _consts[520]))
					*(*int32)(unsafe.Add(mBase, uint32(v5)+4)) = v61
					F__serverLog(m, int32(2), int32(_a930), v5)
					mBase = m.M
					v66 = m.ExcPending
					if v66 != 0 {
						return
					} else {
						v67 = F_connectWithPrimary(m)
						mBase = m.M
						v68 = m.ExcPending
						if v68 != 0 {
							return
						} else {
							m.G0 = v5 + int32(16)
							return
						}
					}
				}
			}
		}
	}
}
func F_replicationSendAck(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v9 int64
	_ = v9
	var v10 int32
	_ = v10
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int64
	_ = v28
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v35 int64
	_ = v35
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	v5 = *(*int32)(unsafe.Add(mBase, _consts[133]))
	if v5 == int32(0) {
		return
	} else {
		v9 = *(*int64)(unsafe.Add(mBase, _consts[524]))
		v10 = *(*int32)(unsafe.Add(mBase, uint32(v5)+200))
		*(*int32)(unsafe.Add(mBase, uint32(v5)+200)) = v10 | int32(8192)
		v17 = base.B2i32(v9 == int64(-1))
		if v9 == int64(-1) {
			v18 = int32(3)
		} else {
			v18 = int32(5)
		}
		F_addReplyArrayLen(m, v5, v18)
		mBase = m.M
		v20 = m.ExcPending
		if v20 != 0 {
			return
		} else {
			F_addReplyBulkCString(m, v5, int32(_a935))
			mBase = m.M
			v23 = m.ExcPending
			if v23 != 0 {
				return
			} else {
				F_addReplyBulkCString(m, v5, int32(_a936))
				mBase = m.M
				v26 = m.ExcPending
				if v26 != 0 {
					return
				} else {
					v27 = *(*int32)(unsafe.Add(mBase, uint32(v5)+104))
					v28 = *(*int64)(unsafe.Add(mBase, uint32(v27)+48))
					F_addReplyBulkLongLong(m, v5, v28)
					mBase = m.M
					v30 = m.ExcPending
					if v30 != 0 {
						return
					} else {
						if v9 == int64(-1) {
							*(*int64)(unsafe.Add(mBase, uint32(v5)+264)) = int64(0)
							v40 = *(*int32)(unsafe.Add(mBase, uint32(v5)+200))
							*(*int32)(unsafe.Add(mBase, uint32(v5)+200)) = v40 & int32(-8193)
							return
						} else {
							F_addReplyBulkCString(m, v5, int32(_a937))
							mBase = m.M
							v33 = m.ExcPending
							if v33 != 0 {
								return
							} else {
								v35 = *(*int64)(unsafe.Add(mBase, _consts[524]))
								F_addReplyBulkLongLong(m, v5, v35)
								mBase = m.M
								v37 = m.ExcPending
								if v37 != 0 {
									return
								} else {
									*(*int64)(unsafe.Add(mBase, uint32(v5)+264)) = int64(0)
									v40 = *(*int32)(unsafe.Add(mBase, uint32(v5)+200))
									*(*int32)(unsafe.Add(mBase, uint32(v5)+200)) = v40 & int32(-8193)
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
func F_replicationSendAuth(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v21 int64
	_ = v21
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v29 int64
	_ = v29
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v55 int32
	_ = v55
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v66 int32
	_ = v66
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v81 int32
	_ = v81
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v95 int32
	_ = v95
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v102 int32
	_ = v102
	var v106 int32
	_ = v106
	var v113 int32
	_ = v113
	var v116 int32
	_ = v116
	var v119 int32
	_ = v119
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v127 int32
	_ = v127
	var v130 int32
	_ = v130
	v2 = int32(0)
	v8 = m.G0
	v10 = v8 - int32(32)
	m.G0 = v10
	v16 = v10 + int32(24)
	v18 = *(*int32)(unsafe.Add(mBase, _consts[563]))
	*(*int32)(unsafe.Add(mBase, uint32(v16))) = v18
	v21 = *(*int64)(unsafe.Add(mBase, _consts[564]))
	*(*int64)(unsafe.Add(mBase, uint32(v10)+16)) = v21
	v24 = v10 + int32(8)
	v26 = *(*int32)(unsafe.Add(mBase, _consts[565]))
	*(*int32)(unsafe.Add(mBase, uint32(v24))) = v26
	v29 = *(*int64)(unsafe.Add(mBase, _consts[566]))
	*(*int64)(unsafe.Add(mBase, uint32(v10))) = v29
	v32 = *(*int32)(unsafe.Add(mBase, _consts[567]))
	if v32 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	v102 = *(*int32)(unsafe.Add(mBase, _consts[147]))
	*(*int32)(unsafe.Add(mBase, uint32(v98))) = v102
	v106 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v102+int32(-1)))))
	switch v106 & int32(7) {
	case 0:
		goto L25
	case 1:
		goto L24
	case 2:
		goto L23
	case 3:
		goto L22
	case 4:
		goto L21
	default:
		v123 = v2
		goto L20
	}
L2:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+20)) = v32
	if v32&int32(3) == int32(0) {
		v62 = v32
		goto L6
	} else {
		goto L7
	}
L3:
	;
	v35 = int32(4)
	v98 = v10 + int32(16) | v35
	v99 = v10 | v35
	v100 = int32(2)
	goto L1
L4:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+4)) = v95
	v98 = v16
	v99 = v24
	v100 = int32(3)
	goto L1
L5:
	;
	v95 = v87 - v32
	goto L4
L6:
	;
	v66 = v62
	goto L14
L7:
	;
	v48 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v32))))
	if v48 != 0 {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v51 = v32
	goto L10
L9:
	;
	v95 = v32 - v32
	goto L4
L10:
	;
	v55 = v51 + int32(1)
	if v55&int32(3) == int32(0) {
		v62 = v55
		goto L6
	} else {
		goto L12
	}
L12:
	;
	v60 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v55))))
	if v60 != 0 {
		v51 = v55
		goto L10
	} else {
		goto L13
	}
L13:
	;
	v87 = v55
	goto L5
L14:
	;
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v66)))
	v75 = int32(-2139062144)
	if (int32(16843008)-v72|v72)&v75 == v75 {
		v66 = v66 + int32(4)
		goto L14
	} else {
		goto L16
	}
L15:
	;
	v81 = v66
	goto L17
L16:
	;
	goto L15
L17:
	;
	v85 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v81))))
	if v85 != 0 {
		v81 = v81 + int32(1)
		goto L17
	} else {
		goto L19
	}
L18:
	;
	v87 = v81
	goto L5
L19:
	;
	goto L18
L20:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v99))) = v123
	v127 = F_sendCommandArgv(m, l0, v100, v10+int32(16), v10)
	mBase = m.M
	v130 = m.ExcPending
	if v130 != 0 {
		goto L26
	} else {
		goto L27
	}
L21:
	;
	v122 = *(*int32)(unsafe.Add(mBase, uint32(v102+int32(-17))))
	v123 = v122
	goto L20
L22:
	;
	v119 = *(*int32)(unsafe.Add(mBase, uint32(v102+int32(-9))))
	v123 = v119
	goto L20
L23:
	;
	v116 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v102+int32(-5)))))
	v123 = v116
	goto L20
L24:
	;
	v113 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v102+int32(-3)))))
	v123 = v113
	goto L20
L25:
	;
	v123 = int32(base.Ui32(v106) >> (uint(int32(3)) % 32))
	goto L20
L26:
	;
	return int32(0)
L27:
	;
	m.G0 = v10 + int32(32)
	return v127
}
func F_replicationSetPrimary(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
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
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
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
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v57 int32
	_ = v57
	var v64 int32
	_ = v64
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
	var v75 int32
	_ = v75
	var v81 int32
	_ = v81
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v91 int64
	_ = v91
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v102 int64
	_ = v102
	var v109 int32
	_ = v109
	var v116 int64
	_ = v116
	var v123 int64
	_ = v123
	var v130 int64
	_ = v130
	var v137 int64
	_ = v137
	var v140 int32
	_ = v140
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v145 int32
	_ = v145
	var v155 int32
	_ = v155
	var v157 int32
	_ = v157
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v169 int32
	_ = v169
	var v173 int64
	_ = v173
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	v8 = m.G0
	v10 = v8 - int32(16)
	m.G0 = v10
	v13 = *(*int32)(unsafe.Add(mBase, _consts[64]))
	F_sdsfree(m, v13)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v16 = int32(_a69)
	v17 = int32(0)
	*(*int32)(unsafe.Add(mBase, _consts[64])) = v17
	v20 = *(*int32)(unsafe.Add(mBase, _consts[133]))
	if v20 == v17 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v36 = F_sdsnew(m, l0)
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		goto L1
	} else {
		goto L6
	}
L4:
	;
	v23 = *(*int32)(unsafe.Add(mBase, uint32(v20)+204))
	*(*int32)(unsafe.Add(mBase, uint32(v20)+204)) = v23&int32(-134217729) | l2<<(uint(int32(27))%32)&int32(134217728)
	v33 = *(*int32)(unsafe.Add(mBase, _consts[133]))
	v34 = F_freeClient(m, v33)
	mBase = m.M
	v35 = m.ExcPending
	if v35 != 0 {
		goto L1
	} else {
		goto L5
	}
L5:
	;
	goto L3
L6:
	;
	v38 = int32(_a69)
	*(*int32)(unsafe.Add(mBase, _consts[520])) = l1
	*(*int32)(unsafe.Add(mBase, _consts[64])) = v36
	if l3 == int32(0) {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	v47 = F_setOOMScoreAdj(m, int32(-1))
	mBase = m.M
	v48 = m.ExcPending
	if v48 != 0 {
		goto L1
	} else {
		goto L10
	}
L8:
	;
	F_disconnectOrRedirectAllBlockedClients(m)
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L1
	} else {
		goto L9
	}
L9:
	;
	goto L7
L10:
	;
	v50 = F_cancelReplicationHandshake(m, int32(0))
	mBase = m.M
	v51 = m.ExcPending
	if v51 != 0 {
		goto L1
	} else {
		goto L11
	}
L11:
	;
	if l2 != 0 {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	F_moduleFireServerEvent(m, int64(0), int32(1), int32(0))
	mBase = m.M
	v155 = m.ExcPending
	if v155 != 0 {
		goto L1
	} else {
		goto L26
	}
L13:
	;
	if v13 != 0 {
		goto L12
	} else {
		goto L14
	}
L14:
	;
	v53 = *(*int32)(unsafe.Add(mBase, _consts[134]))
	if v53 == int32(0) {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	v81 = *(*int32)(unsafe.Add(mBase, _consts[15]))
	if int32(2) < v81 {
		goto L21
	} else {
		goto L22
	}
L16:
	;
	v57 = *(*int32)(unsafe.Add(mBase, _consts[15]))
	if int32(2) < v57 {
		v67 = v53
		goto L17
	} else {
		goto L18
	}
L17:
	;
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v67)+200))
	*(*int32)(unsafe.Add(mBase, uint32(v67)+200)) = v68 & int32(-2)
	v73 = *(*int32)(unsafe.Add(mBase, _consts[134]))
	v74 = F_freeClient(m, v73)
	mBase = m.M
	v75 = m.ExcPending
	if v75 != 0 {
		goto L1
	} else {
		goto L20
	}
L18:
	;
	F__serverLog(m, int32(2), int32(_a1045), int32(0))
	mBase = m.M
	v64 = m.ExcPending
	if v64 != 0 {
		goto L1
	} else {
		goto L19
	}
L19:
	;
	v66 = *(*int32)(unsafe.Add(mBase, _consts[134]))
	v67 = v66
	goto L17
L20:
	;
	*(*int32)(unsafe.Add(mBase, _consts[134])) = int32(0)
	goto L15
L21:
	;
	v89 = int32(_a69)
	v91 = *(*int64)(unsafe.Add(mBase, _consts[31]))
	*(*int64)(unsafe.Add(mBase, _consts[574])) = v91
	F_replicationCreatePrimaryClientWithHandler(m, int32(0), int32(-1), int32(107))
	mBase = m.M
	v97 = m.ExcPending
	if v97 != 0 {
		goto L1
	} else {
		goto L24
	}
L22:
	;
	F__serverLog(m, int32(2), int32(_a1044), int32(0))
	mBase = m.M
	v88 = m.ExcPending
	if v88 != 0 {
		goto L1
	} else {
		goto L23
	}
L23:
	;
	goto L21
L24:
	;
	v98 = int32(_a69)
	v99 = *(*int32)(unsafe.Add(mBase, _consts[133]))
	v100 = *(*int32)(unsafe.Add(mBase, uint32(v99)+104))
	v102 = *(*int64)(unsafe.Add(mBase, _consts[528]))
	*(*int64)(unsafe.Add(mBase, uint32(v100)+104)) = v102
	v109 = int32(*(*uint8)(unsafe.Add(mBase, _consts[511])))
	*(*uint8)(unsafe.Add(mBase, uint32(v100+int32(144)))) = uint8(v109)
	v116 = *(*int64)(unsafe.Add(mBase, _consts[532]))
	*(*int64)(unsafe.Add(mBase, uint32(v100+int32(136)))) = v116
	v123 = *(*int64)(unsafe.Add(mBase, _consts[531]))
	*(*int64)(unsafe.Add(mBase, uint32(v100+int32(128)))) = v123
	v130 = *(*int64)(unsafe.Add(mBase, _consts[530]))
	*(*int64)(unsafe.Add(mBase, uint32(v100+int32(120)))) = v130
	v137 = *(*int64)(unsafe.Add(mBase, _consts[529]))
	*(*int64)(unsafe.Add(mBase, uint32(v100+int32(112)))) = v137
	v140 = *(*int32)(unsafe.Add(mBase, _consts[133]))
	F_unlinkClient(m, v140)
	mBase = m.M
	v142 = m.ExcPending
	if v142 != 0 {
		goto L1
	} else {
		goto L25
	}
L25:
	;
	v143 = int32(_a69)
	v145 = *(*int32)(unsafe.Add(mBase, _consts[133]))
	*(*int32)(unsafe.Add(mBase, _consts[134])) = v145
	*(*int32)(unsafe.Add(mBase, _consts[133])) = int32(0)
	goto L12
L26:
	;
	v157 = *(*int32)(unsafe.Add(mBase, _consts[518]))
	if v157 != int32(14) {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	v165 = int32(_a69)
	*(*int32)(unsafe.Add(mBase, _consts[518])) = int32(1)
	v169 = *(*int32)(unsafe.Add(mBase, _consts[15]))
	if int32(2) < v169 {
		goto L30
	} else {
		goto L31
	}
L28:
	;
	F_moduleFireServerEvent(m, int64(7), int32(1), int32(0))
	mBase = m.M
	v164 = m.ExcPending
	if v164 != 0 {
		goto L1
	} else {
		goto L29
	}
L29:
	;
	goto L27
L30:
	;
	v179 = F_connectWithPrimary(m)
	mBase = m.M
	v180 = m.ExcPending
	if v180 != 0 {
		goto L1
	} else {
		goto L33
	}
L31:
	;
	v173 = *(*int64)(unsafe.Add(mBase, _consts[64]))
	*(*int64)(unsafe.Add(mBase, uint32(v10))) = v173
	F__serverLog(m, int32(2), int32(_a1043), v10)
	mBase = m.M
	v178 = m.ExcPending
	if v178 != 0 {
		goto L1
	} else {
		goto L32
	}
L32:
	;
	goto L30
L33:
	;
	m.G0 = v10 + int32(16)
	return
}
func F_replicationSteadyStateInit(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v18 int32
	_ = v18
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v53 int32
	_ = v53
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	v3 = m.G0
	v5 = v3 - int32(32)
	m.G0 = v5
	v8 = *(*int32)(unsafe.Add(mBase, _consts[133]))
	v9 = *(*int32)(unsafe.Add(mBase, uint32(v8)+8))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(v9)))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(v11)+84))
	v13 = m.T0[v12].(func(*base.Module, int32, int32) int32)(m, v9, int32(107))
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		return
	} else {
		if v13 == int32(0) {
			v36 = *(*int32)(unsafe.Add(mBase, _consts[133]))
			v37 = F_clientHasPendingReplies(m, v36)
			mBase = m.M
			v38 = m.ExcPending
			if v38 != 0 {
				return
			} else {
				if v37 == int32(0) {
					m.G0 = v5 + int32(32)
					return
				} else {
					v42 = *(*int32)(unsafe.Add(mBase, _consts[133]))
					v43 = *(*int32)(unsafe.Add(mBase, uint32(v42)+8))
					v46 = *(*int32)(unsafe.Add(mBase, uint32(v43)))
					v47 = *(*int32)(unsafe.Add(mBase, uint32(v46)+80))
					v48 = m.T0[v47].(func(*base.Module, int32, int32, int32) int32)(m, v43, int32(954), int32(0))
					mBase = m.M
					v49 = m.ExcPending
					if v49 != 0 {
						return
					} else {
						if v48 == int32(0) {
							m.G0 = v5 + int32(32)
							return
						} else {
							v53 = *(*int32)(unsafe.Add(mBase, _consts[15]))
							if int32(3) < v53 {
								v65 = *(*int32)(unsafe.Add(mBase, _consts[133]))
								F_freeClientAsync(m, v65)
								mBase = m.M
								v67 = m.ExcPending
								if v67 != 0 {
									return
								} else {
									m.G0 = v5 + int32(32)
									return
								}
							} else {
								v57 = *(*int32)(unsafe.Add(mBase, _consts[9]))
								v58 = F___strerror_l(m, v57, v57)
								mBase = m.M
								*(*int32)(unsafe.Add(mBase, uint32(v5))) = v58
								F__serverLog(m, int32(3), int32(_a995), v5)
								mBase = m.M
								v63 = m.ExcPending
								if v63 != 0 {
									return
								} else {
									v65 = *(*int32)(unsafe.Add(mBase, _consts[133]))
									F_freeClientAsync(m, v65)
									mBase = m.M
									v67 = m.ExcPending
									if v67 != 0 {
										return
									} else {
										m.G0 = v5 + int32(32)
										return
									}
								}
							}
						}
					}
				}
			}
		} else {
			v18 = *(*int32)(unsafe.Add(mBase, _consts[15]))
			if int32(3) < v18 {
				v32 = *(*int32)(unsafe.Add(mBase, _consts[133]))
				F_freeClientAsync(m, v32)
				mBase = m.M
				v34 = m.ExcPending
				if v34 != 0 {
					return
				} else {
					v36 = *(*int32)(unsafe.Add(mBase, _consts[133]))
					v37 = F_clientHasPendingReplies(m, v36)
					mBase = m.M
					v38 = m.ExcPending
					if v38 != 0 {
						return
					} else {
						if v37 == int32(0) {
							m.G0 = v5 + int32(32)
							return
						} else {
							v42 = *(*int32)(unsafe.Add(mBase, _consts[133]))
							v43 = *(*int32)(unsafe.Add(mBase, uint32(v42)+8))
							v46 = *(*int32)(unsafe.Add(mBase, uint32(v43)))
							v47 = *(*int32)(unsafe.Add(mBase, uint32(v46)+80))
							v48 = m.T0[v47].(func(*base.Module, int32, int32, int32) int32)(m, v43, int32(954), int32(0))
							mBase = m.M
							v49 = m.ExcPending
							if v49 != 0 {
								return
							} else {
								if v48 == int32(0) {
									m.G0 = v5 + int32(32)
									return
								} else {
									v53 = *(*int32)(unsafe.Add(mBase, _consts[15]))
									if int32(3) < v53 {
										v65 = *(*int32)(unsafe.Add(mBase, _consts[133]))
										F_freeClientAsync(m, v65)
										mBase = m.M
										v67 = m.ExcPending
										if v67 != 0 {
											return
										} else {
											m.G0 = v5 + int32(32)
											return
										}
									} else {
										v57 = *(*int32)(unsafe.Add(mBase, _consts[9]))
										v58 = F___strerror_l(m, v57, v57)
										mBase = m.M
										*(*int32)(unsafe.Add(mBase, uint32(v5))) = v58
										F__serverLog(m, int32(3), int32(_a995), v5)
										mBase = m.M
										v63 = m.ExcPending
										if v63 != 0 {
											return
										} else {
											v65 = *(*int32)(unsafe.Add(mBase, _consts[133]))
											F_freeClientAsync(m, v65)
											mBase = m.M
											v67 = m.ExcPending
											if v67 != 0 {
												return
											} else {
												m.G0 = v5 + int32(32)
												return
											}
										}
									}
								}
							}
						}
					}
				}
			} else {
				v22 = *(*int32)(unsafe.Add(mBase, _consts[9]))
				v23 = F___strerror_l(m, v22, v22)
				mBase = m.M
				*(*int32)(unsafe.Add(mBase, uint32(v5)+16)) = v23
				F__serverLog(m, int32(3), int32(_a996), v5+int32(16))
				mBase = m.M
				v30 = m.ExcPending
				if v30 != 0 {
					return
				} else {
					v32 = *(*int32)(unsafe.Add(mBase, _consts[133]))
					F_freeClientAsync(m, v32)
					mBase = m.M
					v34 = m.ExcPending
					if v34 != 0 {
						return
					} else {
						v36 = *(*int32)(unsafe.Add(mBase, _consts[133]))
						v37 = F_clientHasPendingReplies(m, v36)
						mBase = m.M
						v38 = m.ExcPending
						if v38 != 0 {
							return
						} else {
							if v37 == int32(0) {
								m.G0 = v5 + int32(32)
								return
							} else {
								v42 = *(*int32)(unsafe.Add(mBase, _consts[133]))
								v43 = *(*int32)(unsafe.Add(mBase, uint32(v42)+8))
								v46 = *(*int32)(unsafe.Add(mBase, uint32(v43)))
								v47 = *(*int32)(unsafe.Add(mBase, uint32(v46)+80))
								v48 = m.T0[v47].(func(*base.Module, int32, int32, int32) int32)(m, v43, int32(954), int32(0))
								mBase = m.M
								v49 = m.ExcPending
								if v49 != 0 {
									return
								} else {
									if v48 == int32(0) {
										m.G0 = v5 + int32(32)
										return
									} else {
										v53 = *(*int32)(unsafe.Add(mBase, _consts[15]))
										if int32(3) < v53 {
											v65 = *(*int32)(unsafe.Add(mBase, _consts[133]))
											F_freeClientAsync(m, v65)
											mBase = m.M
											v67 = m.ExcPending
											if v67 != 0 {
												return
											} else {
												m.G0 = v5 + int32(32)
												return
											}
										} else {
											v57 = *(*int32)(unsafe.Add(mBase, _consts[9]))
											v58 = F___strerror_l(m, v57, v57)
											mBase = m.M
											*(*int32)(unsafe.Add(mBase, uint32(v5))) = v58
											F__serverLog(m, int32(3), int32(_a995), v5)
											mBase = m.M
											v63 = m.ExcPending
											if v63 != 0 {
												return
											} else {
												v65 = *(*int32)(unsafe.Add(mBase, _consts[133]))
												F_freeClientAsync(m, v65)
												mBase = m.M
												v67 = m.ExcPending
												if v67 != 0 {
													return
												} else {
													m.G0 = v5 + int32(32)
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
