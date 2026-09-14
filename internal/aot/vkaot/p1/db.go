package p1

import (
	base "github.com/shibukawa/vkmem/internal/aot/vkaot/base"
	"unsafe"
)

func F_dbAddInternal(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
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
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v24 int32
	_ = v24
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
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v100 int32
	_ = v100
	var v103 int32
	_ = v103
	var v105 int32
	_ = v105
	var v112 int32
	_ = v112
	v7 = F_objectGetVal(m, l1)
	mBase = m.M
	v9 = *(*int32)(unsafe.Add(mBase, _consts[110]))
	if v9 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	if l3 == int32(0) {
		goto L8
	} else {
		goto L9
	}
L2:
	;
	v11 = F_getKeySlot(m, v7)
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		goto L4
	} else {
		goto L5
	}
L3:
	;
	v13 = int32(0)
	goto L1
L4:
	;
	return
L5:
	;
	v13 = v11
	goto L1
L6:
	;
	F__serverAssertWithInfo(m, int32(0), l1, int32(_a559), int32(_a560), int32(212))
	mBase = m.M
	v112 = m.ExcPending
	if v112 != 0 {
		goto L4
	} else {
		goto L35
	}
L7:
	;
	v34 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	v35 = F_objectGetVal(m, l1)
	mBase = m.M
	v37 = F_objectSetKeyAndExpire(m, v34, v35, int64(-1))
	mBase = m.M
	v38 = m.ExcPending
	if v38 != 0 {
		goto L4
	} else {
		goto L17
	}
L8:
	;
	v26 = *(*int32)(unsafe.Add(mBase, _consts[113]))
	if v26 == int32(0) {
		goto L7
	} else {
		goto L13
	}
L9:
	;
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v17 = F_objectGetVal(m, l1)
	mBase = m.M
	v18 = F_kvstoreHashtableFindRef(m, v16, v13, v17)
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		goto L4
	} else {
		goto L10
	}
L10:
	;
	if v18 == int32(0) {
		goto L7
	} else {
		goto L11
	}
L11:
	;
	F_dbSetValue(m, l0, l1, l2, int32(1), v18)
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		goto L4
	} else {
		goto L12
	}
L12:
	;
	return
L13:
	;
	v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v30 = F_objectGetVal(m, l1)
	mBase = m.M
	v31 = F_kvstoreHashtableFindRef(m, v29, v13, v30)
	mBase = m.M
	v32 = m.ExcPending
	if v32 != 0 {
		goto L4
	} else {
		goto L14
	}
L14:
	;
	if v31 != 0 {
		goto L6
	} else {
		goto L15
	}
L15:
	;
	goto L7
L16:
	;
	v83 = *(*int32)(unsafe.Add(mBase, uint32(v37)+4))
	if base.Ui32(int32(-9)) < base.Ui32(v83) {
		goto L30
	} else {
		goto L31
	}
L17:
	;
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v37)))
	if v39&int32(15) != int32(4) {
		goto L16
	} else {
		goto L18
	}
L18:
	;
	v44 = F_hashTypeHasVolatileFields(m, v37)
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L4
	} else {
		goto L19
	}
L19:
	;
	if v44 == int32(0) {
		goto L16
	} else {
		goto L20
	}
L20:
	;
	v48 = int32(0)
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v37)+4))
	if v51&int32(2) == v48 {
		v71 = v48
		goto L22
	} else {
		goto L23
	}
L21:
	;
	v73 = *(*int32)(unsafe.Add(mBase, _consts[110]))
	if v73 != 0 {
		goto L25
	} else {
		goto L26
	}
L22:
	;
	goto L21
L23:
	;
	v65 = v37 + (v51&int32(4) ^ int32(12)) + v51<<(uint(int32(3))%32)&int32(8)
	v66 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v65))))
	v71 = v65 + v66 + int32(1)
	goto L22
L24:
	;
	v78 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v79 = F_kvstoreHashtableAdd(m, v78, v77, v37)
	mBase = m.M
	v80 = m.ExcPending
	if v80 != 0 {
		goto L4
	} else {
		goto L28
	}
L25:
	;
	v75 = F_getKeySlot(m, v71)
	mBase = m.M
	v76 = m.ExcPending
	if v76 != 0 {
		goto L4
	} else {
		goto L27
	}
L26:
	;
	v77 = int32(0)
	goto L24
L27:
	;
	v77 = v75
	goto L24
L28:
	;
	goto L16
L29:
	;
	v93 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v94 = F_kvstoreHashtableAdd(m, v93, v13, v37)
	mBase = m.M
	v95 = m.ExcPending
	if v95 != 0 {
		goto L4
	} else {
		goto L32
	}
L30:
	;
	goto L29
L31:
	;
	v86 = F_lrulfu_init(m)
	mBase = m.M
	v87 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v37))))
	*(*int32)(unsafe.Add(mBase, uint32(v37))) = v87 | v86<<(uint(int32(8))%32)
	goto L30
L32:
	;
	v96 = *(*int32)(unsafe.Add(mBase, uint32(v37)))
	F_signalKeyAsReady(m, l0, l1, v96&int32(15))
	mBase = m.M
	v100 = m.ExcPending
	if v100 != 0 {
		goto L4
	} else {
		goto L33
	}
L33:
	;
	v103 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	F_notifyKeyspaceEvent(m, int32(16384), int32(_a558), l1, v103)
	mBase = m.M
	v105 = m.ExcPending
	if v105 != 0 {
		goto L4
	} else {
		goto L34
	}
L34:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v37
	return
L35:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_dbDelete(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	v5 = int32(_a20)
	v6 = *(*int32)(unsafe.Add(mBase, _consts[175]))
	v7 = F_objectGetVal(m, l1)
	mBase = m.M
	v9 = *(*int32)(unsafe.Add(mBase, _consts[110]))
	if v9 != 0 {
		v11 = F_getKeySlot(m, v7)
		mBase = m.M
		v14 = m.ExcPending
		if v14 != 0 {
			return int32(0)
		} else {
			v15 = v11
			v17 = F_dbGenericDeleteWithDictIndex(m, l0, l1, v6, int32(1), v15)
			mBase = m.M
			v18 = m.ExcPending
			if v18 != 0 {
				return int32(0)
			} else {
				return v17
			}
		}
	} else {
		v15 = int32(0)
		v17 = F_dbGenericDeleteWithDictIndex(m, l0, l1, v6, int32(1), v15)
		mBase = m.M
		v18 = m.ExcPending
		if v18 != 0 {
			return int32(0)
		} else {
			return v17
		}
	}
}
func F_dbFind(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
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
	v3 = int32(0)
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v11 = *(*int32)(unsafe.Add(mBase, _consts[110]))
	if v11 == v3 {
		v18 = v3
		*(*int32)(unsafe.Add(mBase, uint32(v7)+12)) = int32(0)
		v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		v24 = F_kvstoreHashtableFind(m, v21, v18, l1, v7+int32(12))
		mBase = m.M
		v25 = m.ExcPending
		if v25 != 0 {
			return int32(0)
		} else {
			v26 = *(*int32)(unsafe.Add(mBase, uint32(v7)+12))
			m.G0 = v7 + int32(16)
			return v26
		}
	} else {
		v14 = F_getKeySlot(m, l1)
		mBase = m.M
		v17 = m.ExcPending
		if v17 != 0 {
			return int32(0)
		} else {
			v18 = v14
			*(*int32)(unsafe.Add(mBase, uint32(v7)+12)) = int32(0)
			v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			v24 = F_kvstoreHashtableFind(m, v21, v18, l1, v7+int32(12))
			mBase = m.M
			v25 = m.ExcPending
			if v25 != 0 {
				return int32(0)
			} else {
				v26 = *(*int32)(unsafe.Add(mBase, uint32(v7)+12))
				m.G0 = v7 + int32(16)
				return v26
			}
		}
	}
}
func F_dbHasNoKeys(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v21 int64
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v28 int64
	_ = v28
	var v32 int32
	_ = v32
	v3 = int32(1)
	if l0 < int32(0) {
		v32 = v3
	} else {
		v7 = *(*int32)(unsafe.Add(mBase, _consts[172]))
		if v7 <= l0 {
			v32 = v3
		} else {
			v9 = int32(0)
			v10 = *(*int32)(unsafe.Add(mBase, _consts[173]))
			v14 = *(*int32)(unsafe.Add(mBase, uint32(v10+l0<<(uint(int32(2))%32))))
			if v14 == v9 {
				v32 = v3
			} else {
				v17 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
				v18 = *(*int32)(unsafe.Add(mBase, uint32(v17)+12))
				if v18 == int32(1) {
					v22 = *(*int32)(unsafe.Add(mBase, uint32(v17)+8))
					v23 = *(*int32)(unsafe.Add(mBase, uint32(v22)))
					if v23 != 0 {
						v25 = F_hashtableSize(m, v23)
						mBase = m.M
						v28 = base.I64_extend_i32_u(v25)
					} else {
						v28 = int64(0)
					}
				} else {
					v21 = *(*int64)(unsafe.Add(mBase, uint32(v17)+40))
					v28 = v21
				}
				v32 = base.B2i32(v28 == int64(0))
			}
		}
	}
	return v32
}
func F_dbReplaceValue(m *base.Module, l0 int32, l1 int32, l2 int32) {
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	v4 = int32(0)
	F_dbSetValue(m, l0, l1, l2, v4, v4)
	v7 = m.ExcPending
	if v7 != 0 {
		return
	} else {
		return
	}
}
func F_dbTotalServerKeyCount(m *base.Module) int64 {
	mBase := m.M
	_ = mBase
	var v1 int64
	_ = v1
	var v5 int32
	_ = v5
	var v9 int64
	_ = v9
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v27 int64
	_ = v27
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v42 int64
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v49 int64
	_ = v49
	var v51 int64
	_ = v51
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v57 int64
	_ = v57
	v1 = int64(0)
	v5 = *(*int32)(unsafe.Add(mBase, _consts[172]))
	if v5 < int32(1) {
		v57 = v1
	} else {
		v9 = v1
		v10 = int32(0)
		for {
			v12 = int32(1)
			if v10 < int32(0) {
				v31 = v12
			} else {
				v16 = *(*int32)(unsafe.Add(mBase, _consts[172]))
				if v16 <= v10 {
					v31 = v12
				} else {
					v18 = int32(0)
					v19 = *(*int32)(unsafe.Add(mBase, _consts[173]))
					v23 = *(*int32)(unsafe.Add(mBase, uint32(v19+v10<<(uint(int32(2))%32))))
					if v23 == v18 {
						v31 = v12
					} else {
						v26 = *(*int32)(unsafe.Add(mBase, uint32(v23)))
						v27 = F_kvstoreSize(m, v26)
						mBase = m.M
						v31 = base.B2i32(v27 == int64(0))
					}
				}
			}
			if v31 != 0 {
				v51 = v9
			} else {
				v33 = *(*int32)(unsafe.Add(mBase, _consts[173]))
				v37 = *(*int32)(unsafe.Add(mBase, uint32(v33+v10<<(uint(int32(2))%32))))
				v38 = *(*int32)(unsafe.Add(mBase, uint32(v37)))
				v39 = *(*int32)(unsafe.Add(mBase, uint32(v38)+12))
				if v39 == int32(1) {
					v43 = *(*int32)(unsafe.Add(mBase, uint32(v38)+8))
					v44 = *(*int32)(unsafe.Add(mBase, uint32(v43)))
					if v44 != 0 {
						v46 = F_hashtableSize(m, v44)
						mBase = m.M
						v49 = base.I64_extend_i32_u(v46)
					} else {
						v49 = int64(0)
					}
				} else {
					v42 = *(*int64)(unsafe.Add(mBase, uint32(v38)+40))
					v49 = v42
				}
				v51 = v49 + v9
			}
			v53 = v10 + int32(1)
			v55 = *(*int32)(unsafe.Add(mBase, _consts[172]))
			if v53 < v55 {
				v9 = v51
				v10 = v53
				continue
			} else {
				break
			}
			break
		}
		v57 = v51
	}
	return v57
}
func F_db_errorfb(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v17 int32
	_ = v17
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v130 int32
	_ = v130
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v143 int32
	_ = v143
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v166 int32
	_ = v166
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v184 int32
	_ = v184
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v194 int32
	_ = v194
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v205 int32
	_ = v205
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v237 int32
	_ = v237
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v245 int32
	_ = v245
	var v252 int32
	_ = v252
	var v255 int32
	_ = v255
	var v260 int32
	_ = v260
	var v261 int32
	_ = v261
	var v266 int32
	_ = v266
	var v272 int32
	_ = v272
	var v275 int32
	_ = v275
	var v278 int32
	_ = v278
	var v280 int32
	_ = v280
	var v285 int32
	_ = v285
	var v286 int32
	_ = v286
	var v287 int32
	_ = v287
	var v288 int32
	_ = v288
	var v289 int32
	_ = v289
	var v291 int32
	_ = v291
	var v293 int32
	_ = v293
	var v297 int32
	_ = v297
	var v299 int32
	_ = v299
	var v304 int32
	_ = v304
	var v308 int32
	_ = v308
	var v310 int32
	_ = v310
	var v319 int32
	_ = v319
	var v323 int32
	_ = v323
	var v329 int32
	_ = v329
	var v330 int32
	_ = v330
	var v339 int32
	_ = v339
	var v343 int32
	_ = v343
	var v344 int32
	_ = v344
	var v351 int32
	_ = v351
	var v357 int32
	_ = v357
	var v360 int32
	_ = v360
	var v363 int32
	_ = v363
	var v365 int32
	_ = v365
	var v370 int32
	_ = v370
	var v371 int32
	_ = v371
	var v372 int32
	_ = v372
	var v373 int32
	_ = v373
	var v374 int32
	_ = v374
	var v376 int32
	_ = v376
	var v378 int32
	_ = v378
	var v382 int32
	_ = v382
	var v384 int32
	_ = v384
	var v389 int32
	_ = v389
	var v393 int32
	_ = v393
	var v395 int32
	_ = v395
	var v404 int32
	_ = v404
	var v407 int32
	_ = v407
	var v412 int32
	_ = v412
	var v417 int32
	_ = v417
	var v421 int32
	_ = v421
	var v424 int32
	_ = v424
	var v430 int32
	_ = v430
	var v433 int32
	_ = v433
	var v436 int32
	_ = v436
	var v438 int32
	_ = v438
	var v443 int32
	_ = v443
	var v444 int32
	_ = v444
	var v445 int32
	_ = v445
	var v446 int32
	_ = v446
	var v447 int32
	_ = v447
	var v449 int32
	_ = v449
	var v451 int32
	_ = v451
	var v455 int32
	_ = v455
	var v457 int32
	_ = v457
	var v462 int32
	_ = v462
	var v466 int32
	_ = v466
	var v468 int32
	_ = v468
	var v477 int32
	_ = v477
	var v483 int32
	_ = v483
	var v492 int32
	_ = v492
	var v495 int32
	_ = v495
	var v498 int32
	_ = v498
	var v500 int32
	_ = v500
	var v505 int32
	_ = v505
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
	var v513 int32
	_ = v513
	var v517 int32
	_ = v517
	var v519 int32
	_ = v519
	var v524 int32
	_ = v524
	var v528 int32
	_ = v528
	var v530 int32
	_ = v530
	var v539 int32
	_ = v539
	var v540 int32
	_ = v540
	var v545 int32
	_ = v545
	var v550 int32
	_ = v550
	var v551 int32
	_ = v551
	var v557 int32
	_ = v557
	var v558 int32
	_ = v558
	var v559 int32
	_ = v559
	var v563 int32
	_ = v563
	var v568 int32
	_ = v568
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
	var v581 int32
	_ = v581
	var v582 int32
	_ = v582
	var v583 int32
	_ = v583
	var v584 int32
	_ = v584
	var v587 int32
	_ = v587
	var v591 int32
	_ = v591
	var v592 int32
	_ = v592
	var v596 int32
	_ = v596
	var v598 int32
	_ = v598
	var v601 int32
	_ = v601
	var v602 int32
	_ = v602
	var v603 int32
	_ = v603
	var v608 int32
	_ = v608
	var v610 int32
	_ = v610
	var v611 int32
	_ = v611
	var v617 int32
	_ = v617
	var v623 int32
	_ = v623
	var v626 int32
	_ = v626
	var v629 int32
	_ = v629
	var v631 int32
	_ = v631
	var v636 int32
	_ = v636
	var v637 int32
	_ = v637
	var v638 int32
	_ = v638
	var v639 int32
	_ = v639
	var v640 int32
	_ = v640
	var v642 int32
	_ = v642
	var v644 int32
	_ = v644
	var v648 int32
	_ = v648
	var v650 int32
	_ = v650
	var v655 int32
	_ = v655
	var v659 int32
	_ = v659
	var v661 int32
	_ = v661
	var v670 int32
	_ = v670
	var v679 int32
	_ = v679
	var v680 int32
	_ = v680
	var v686 int32
	_ = v686
	v9 = m.G0
	v11 = v9 - int32(160)
	m.G0 = v11
	goto L5
L1:
	;
	v139 = v136 | int32(2)
	v140 = F_lua_isnumber(m, l0, v139)
	mBase = m.M
	v143 = m.ExcPending
	if v143 != 0 {
		goto L38
	} else {
		goto L39
	}
L2:
	;
	if v71 != int32(8) {
		v136 = int32(0)
		v137 = l0
		goto L1
	} else {
		goto L17
	}
L3:
	;
	v65 = m.G398
	if v22 != v65 {
		goto L15
	} else {
		goto L16
	}
L5:
	;
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v22 = v17 + int32(0)
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if base.Ui32(v22) < base.Ui32(v23) {
		goto L3
	} else {
		goto L6
	}
L6:
	;
	v71 = int32(-1)
	goto L2
L15:
	;
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v22)+8))
	v71 = v68
	goto L2
L16:
	;
	v71 = int32(-1)
	goto L2
L17:
	;
	goto L21
L18:
	;
	v136 = int32(1)
	v137 = v134
	goto L1
L19:
	;
	v130 = *(*int32)(unsafe.Add(mBase, uint32(v87)+8))
	if v130 != int32(8) {
		v134 = int32(0)
		goto L34
	} else {
		goto L35
	}
L21:
	;
	v78 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v79 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v84 = v79 + int32(0)
	v85 = m.G398
	if base.Ui32(v84) < base.Ui32(v78) {
		goto L22
	} else {
		goto L23
	}
L22:
	;
	v87 = v84
	goto L24
L23:
	;
	v87 = v85
	goto L24
L24:
	;
	goto L19
L34:
	;
	goto L18
L35:
	;
	v133 = *(*int32)(unsafe.Add(mBase, uint32(v87)))
	v134 = v133
	goto L34
L36:
	;
	v178 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v179 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	goto L53
L37:
	;
	v177 = base.B2i32(l0 == v137)
	goto L36
L38:
	;
	return int32(0)
L39:
	;
	if v140 == int32(0) {
		goto L37
	} else {
		goto L40
	}
L40:
	;
	v146 = F_lua_tointeger(m, l0, v139)
	mBase = m.M
	v147 = m.ExcPending
	if v147 != 0 {
		goto L38
	} else {
		goto L41
	}
L41:
	;
	goto L44
L42:
	;
	v177 = v146
	goto L36
L43:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v166 + int32(-16)
	goto L42
L44:
	;
	v166 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	goto L43
L50:
	;
	m.G0 = v11 + int32(160)
	return int32(1)
L51:
	;
	v261 = m.G3
	F_lua_pushlstring(m, l0, v261+int32(_a2704), int32(16))
	mBase = m.M
	v266 = m.ExcPending
	if v266 != 0 {
		goto L38
	} else {
		goto L73
	}
L52:
	;
	v190 = int32(1)
	v191 = v136 + v190
	if v191 < v190 {
		goto L58
	} else {
		goto L59
	}
L53:
	;
	if (v178-v179)>>(uint(int32(4))%32) != v136 {
		goto L52
	} else {
		goto L54
	}
L54:
	;
	v184 = m.G3
	F_lua_pushlstring(m, l0, v184+int32(_a320), int32(0))
	mBase = m.M
	v189 = m.ExcPending
	if v189 != 0 {
		goto L38
	} else {
		goto L55
	}
L55:
	;
	goto L51
L56:
	;
	if v252 == int32(0) {
		goto L50
	} else {
		goto L71
	}
L57:
	;
	v242 = m.G398
	if v241 != v242 {
		goto L69
	} else {
		goto L70
	}
L58:
	;
	if v191 < int32(-9999) {
		goto L61
	} else {
		goto L62
	}
L59:
	;
	v194 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v199 = v194 + v191<<(uint(int32(4))%32) + int32(-16)
	v200 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if base.Ui32(v199) < base.Ui32(v200) {
		v241 = v199
		goto L57
	} else {
		goto L60
	}
L60:
	;
	v252 = int32(0)
	goto L56
L61:
	;
	switch v136 + int32(10003) {
	case 0:
		goto L65
	case 1:
		goto L66
	case 2:
		goto L63
	default:
		goto L64
	}
L62:
	;
	v205 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v241 = v205 + v191<<(uint(int32(4))%32)
	goto L57
L63:
	;
	v237 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v241 = v237 + int32(96)
	goto L57
L64:
	;
	v224 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v225 = *(*int32)(unsafe.Add(mBase, uint32(v224)+4))
	v226 = *(*int32)(unsafe.Add(mBase, uint32(v225)))
	v227 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v226)+7)))
	if base.Ui32(v227) < base.Ui32(int32(-10002)-v191) {
		goto L67
	} else {
		goto L68
	}
L65:
	;
	v241 = l0 + int32(72)
	goto L57
L66:
	;
	v211 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v212 = *(*int32)(unsafe.Add(mBase, uint32(v211)+4))
	v213 = *(*int32)(unsafe.Add(mBase, uint32(v212)))
	v214 = *(*int32)(unsafe.Add(mBase, uint32(v213)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+96)) = int32(5)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+88)) = v214
	v241 = l0 + int32(88)
	goto L57
L67:
	;
	v252 = int32(0)
	goto L56
L68:
	;
	v241 = v226 + (int32(-10003)-v191)<<(uint(int32(4))%32) + int32(24)
	goto L57
L69:
	;
	v245 = *(*int32)(unsafe.Add(mBase, uint32(v241)+8))
	v252 = base.B2i32(base.Ui32(v245+int32(-3)) < base.Ui32(int32(2)))
	goto L56
L70:
	;
	v252 = int32(0)
	goto L56
L71:
	;
	v255 = m.G3
	F_lua_pushlstring(m, l0, v255+int32(_a26), int32(1))
	mBase = m.M
	v260 = m.ExcPending
	if v260 != 0 {
		goto L38
	} else {
		goto L72
	}
L72:
	;
	goto L51
L73:
	;
	v272 = *(*int32)(unsafe.Add(mBase, uint32(v137)+20))
	if v177 < int32(1) {
		v297 = v177
		v299 = v272
		goto L77
	} else {
		goto L78
	}
L74:
	;
	v679 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v680 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	goto L175
L75:
	;
	if v319 == int32(0) {
		goto L74
	} else {
		goto L88
	}
L76:
	;
	goto L75
L77:
	;
	if v297 != 0 {
		v310 = int32(0)
		goto L85
	} else {
		goto L86
	}
L78:
	;
	v275 = *(*int32)(unsafe.Add(mBase, uint32(v137)+40))
	v278 = v177
	v280 = v272
	goto L79
L79:
	;
	if base.Ui32(v280) <= base.Ui32(v275) {
		v319 = int32(0)
		goto L76
	} else {
		goto L81
	}
L80:
	;
	v297 = v291
	v299 = v293
	goto L77
L81:
	;
	v285 = v278 + int32(-1)
	v286 = *(*int32)(unsafe.Add(mBase, uint32(v280)+4))
	v287 = *(*int32)(unsafe.Add(mBase, uint32(v286)))
	v288 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v287)+6)))
	if v288 != 0 {
		v291 = v285
		goto L82
	} else {
		goto L83
	}
L82:
	;
	v293 = v280 + int32(-24)
	if int32(0) < v291 {
		v278 = v291
		v280 = v293
		goto L79
	} else {
		goto L84
	}
L83:
	;
	v289 = *(*int32)(unsafe.Add(mBase, uint32(v280)+20))
	v291 = v285 - v289
	goto L82
L84:
	;
	goto L80
L85:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11+int32(60))+96)) = v310
	v319 = int32(1)
	goto L76
L86:
	;
	v304 = *(*int32)(unsafe.Add(mBase, uint32(v137)+40))
	if base.Ui32(v299) <= base.Ui32(v304) {
		v319 = int32(0)
		goto L76
	} else {
		goto L87
	}
L87:
	;
	v308 = base.I32_div_s(v299-v304, int32(24))
	v310 = v308
	goto L85
L88:
	;
	v323 = v11 + int32(96)
	v329 = int32(1)
	v330 = v177
	goto L89
L89:
	;
	v339 = v330
	goto L91
L90:
	;
	goto L74
L91:
	;
	v343 = int32(1)
	v344 = v339 + v343
	if (base.B2i32(v339 < int32(12))|(v329^int32(-1)))&v343 != 0 {
		goto L93
	} else {
		goto L94
	}
L92:
	;
	goto L90
L93:
	;
	v540 = m.G3
	F_lua_pushlstring(m, l0, v540+int32(_a2705), int32(2))
	mBase = m.M
	v545 = m.ExcPending
	if v545 != 0 {
		goto L38
	} else {
		goto L141
	}
L94:
	;
	v351 = v339 + int32(11)
	v357 = *(*int32)(unsafe.Add(mBase, uint32(v137)+20))
	if v351 < int32(1) {
		v382 = v351
		v384 = v357
		goto L98
	} else {
		goto L99
	}
L95:
	;
	v492 = *(*int32)(unsafe.Add(mBase, uint32(v137)+20))
	if v483 < int32(1) {
		v517 = v483
		v519 = v492
		goto L129
	} else {
		goto L130
	}
L96:
	;
	if v404 == int32(0) {
		v483 = v339
		goto L95
	} else {
		goto L109
	}
L97:
	;
	goto L96
L98:
	;
	if v382 != 0 {
		v395 = int32(0)
		goto L106
	} else {
		goto L107
	}
L99:
	;
	v360 = *(*int32)(unsafe.Add(mBase, uint32(v137)+40))
	v363 = v351
	v365 = v357
	goto L100
L100:
	;
	if base.Ui32(v365) <= base.Ui32(v360) {
		v404 = int32(0)
		goto L97
	} else {
		goto L102
	}
L101:
	;
	v382 = v376
	v384 = v378
	goto L98
L102:
	;
	v370 = v363 + int32(-1)
	v371 = *(*int32)(unsafe.Add(mBase, uint32(v365)+4))
	v372 = *(*int32)(unsafe.Add(mBase, uint32(v371)))
	v373 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v372)+6)))
	if v373 != 0 {
		v376 = v370
		goto L103
	} else {
		goto L104
	}
L103:
	;
	v378 = v365 + int32(-24)
	if int32(0) < v376 {
		v363 = v376
		v365 = v378
		goto L100
	} else {
		goto L105
	}
L104:
	;
	v374 = *(*int32)(unsafe.Add(mBase, uint32(v365)+20))
	v376 = v370 - v374
	goto L103
L105:
	;
	goto L101
L106:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11+int32(60))+96)) = v395
	v404 = int32(1)
	goto L97
L107:
	;
	v389 = *(*int32)(unsafe.Add(mBase, uint32(v137)+40))
	if base.Ui32(v384) <= base.Ui32(v389) {
		v404 = int32(0)
		goto L97
	} else {
		goto L108
	}
L108:
	;
	v393 = base.I32_div_s(v384-v389, int32(24))
	v395 = v393
	goto L106
L109:
	;
	v407 = m.G3
	F_lua_pushlstring(m, l0, v407+int32(_a2706), int32(5))
	mBase = m.M
	v412 = m.ExcPending
	if v412 != 0 {
		goto L38
	} else {
		goto L110
	}
L110:
	;
	v417 = v344
	goto L111
L111:
	;
	v421 = int32(1)
	v424 = v417 + int32(10)
	v430 = *(*int32)(unsafe.Add(mBase, uint32(v137)+20))
	if v424 < v421 {
		v455 = v424
		v457 = v430
		goto L115
	} else {
		goto L116
	}
L112:
	;
	v483 = v417
	goto L95
L113:
	;
	if v477 != 0 {
		v417 = v417 + v421
		goto L111
	} else {
		goto L126
	}
L114:
	;
	goto L113
L115:
	;
	if v455 != 0 {
		v468 = int32(0)
		goto L123
	} else {
		goto L124
	}
L116:
	;
	v433 = *(*int32)(unsafe.Add(mBase, uint32(v137)+40))
	v436 = v424
	v438 = v430
	goto L117
L117:
	;
	if base.Ui32(v438) <= base.Ui32(v433) {
		v477 = int32(0)
		goto L114
	} else {
		goto L119
	}
L118:
	;
	v455 = v449
	v457 = v451
	goto L115
L119:
	;
	v443 = v436 + int32(-1)
	v444 = *(*int32)(unsafe.Add(mBase, uint32(v438)+4))
	v445 = *(*int32)(unsafe.Add(mBase, uint32(v444)))
	v446 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v445)+6)))
	if v446 != 0 {
		v449 = v443
		goto L120
	} else {
		goto L121
	}
L120:
	;
	v451 = v438 + int32(-24)
	if int32(0) < v449 {
		v436 = v449
		v438 = v451
		goto L117
	} else {
		goto L122
	}
L121:
	;
	v447 = *(*int32)(unsafe.Add(mBase, uint32(v438)+20))
	v449 = v443 - v447
	goto L120
L122:
	;
	goto L118
L123:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11+int32(60))+96)) = v468
	v477 = int32(1)
	goto L114
L124:
	;
	v462 = *(*int32)(unsafe.Add(mBase, uint32(v137)+40))
	if base.Ui32(v457) <= base.Ui32(v462) {
		v477 = int32(0)
		goto L114
	} else {
		goto L125
	}
L125:
	;
	v466 = base.I32_div_s(v457-v462, int32(24))
	v468 = v466
	goto L123
L126:
	;
	goto L112
L127:
	;
	if v539 != 0 {
		v329 = int32(0)
		v330 = v483
		goto L89
	} else {
		goto L140
	}
L128:
	;
	goto L127
L129:
	;
	if v517 != 0 {
		v530 = int32(0)
		goto L137
	} else {
		goto L138
	}
L130:
	;
	v495 = *(*int32)(unsafe.Add(mBase, uint32(v137)+40))
	v498 = v483
	v500 = v492
	goto L131
L131:
	;
	if base.Ui32(v500) <= base.Ui32(v495) {
		v539 = int32(0)
		goto L128
	} else {
		goto L133
	}
L132:
	;
	v517 = v511
	v519 = v513
	goto L129
L133:
	;
	v505 = v498 + int32(-1)
	v506 = *(*int32)(unsafe.Add(mBase, uint32(v500)+4))
	v507 = *(*int32)(unsafe.Add(mBase, uint32(v506)))
	v508 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v507)+6)))
	if v508 != 0 {
		v511 = v505
		goto L134
	} else {
		goto L135
	}
L134:
	;
	v513 = v500 + int32(-24)
	if int32(0) < v511 {
		v498 = v511
		v500 = v513
		goto L131
	} else {
		goto L136
	}
L135:
	;
	v509 = *(*int32)(unsafe.Add(mBase, uint32(v500)+20))
	v511 = v505 - v509
	goto L134
L136:
	;
	goto L132
L137:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11+int32(60))+96)) = v530
	v539 = int32(1)
	goto L128
L138:
	;
	v524 = *(*int32)(unsafe.Add(mBase, uint32(v137)+40))
	if base.Ui32(v519) <= base.Ui32(v524) {
		v539 = int32(0)
		goto L128
	} else {
		goto L139
	}
L139:
	;
	v528 = base.I32_div_s(v519-v524, int32(24))
	v530 = v528
	goto L137
L140:
	;
	goto L74
L141:
	;
	v550 = F_lua_getinfo(m, v137, v540+int32(_a2597), v11+int32(60))
	mBase = m.M
	v551 = m.ExcPending
	if v551 != 0 {
		goto L38
	} else {
		goto L142
	}
L142:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+48)) = v323
	v557 = F_lua_pushfstring(m, l0, v540+int32(_a2707), v11+int32(48))
	mBase = m.M
	v558 = m.ExcPending
	if v558 != 0 {
		goto L38
	} else {
		goto L143
	}
L143:
	;
	v559 = *(*int32)(unsafe.Add(mBase, uint32(v11)+80))
	if v559 < int32(1) {
		goto L144
	} else {
		goto L145
	}
L144:
	;
	v570 = *(*int32)(unsafe.Add(mBase, uint32(v11)+68))
	v571 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v570))))
	if v571 == int32(0) {
		goto L148
	} else {
		goto L149
	}
L145:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+32)) = v559
	v563 = m.G3
	v568 = F_lua_pushfstring(m, l0, v563+int32(_a2708), v11+int32(32))
	mBase = m.M
	v569 = m.ExcPending
	if v569 != 0 {
		goto L38
	} else {
		goto L146
	}
L146:
	;
	goto L144
L147:
	;
	v610 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v611 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	goto L159
L148:
	;
	v583 = *(*int32)(unsafe.Add(mBase, uint32(v11)+72))
	v584 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v583))))
	switch v584 + int32(-109) {
	case 0:
		goto L154
	case 1, 2, 3, 4, 5, 6:
		goto L152
	case 7:
		goto L151
	default:
		goto L153
	}
L149:
	;
	v574 = *(*int32)(unsafe.Add(mBase, uint32(v11)+64))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = v574
	v576 = m.G3
	v581 = F_lua_pushfstring(m, l0, v576+int32(_a2709), v11+int32(16))
	mBase = m.M
	v582 = m.ExcPending
	if v582 != 0 {
		goto L38
	} else {
		goto L150
	}
L150:
	;
	goto L147
L151:
	;
	v603 = m.G3
	F_lua_pushlstring(m, l0, v603+int32(_a2710), int32(2))
	mBase = m.M
	v608 = m.ExcPending
	if v608 != 0 {
		goto L38
	} else {
		goto L158
	}
L152:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11))) = v323
	v596 = *(*int32)(unsafe.Add(mBase, uint32(v11)+88))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+4)) = v596
	v598 = m.G3
	v601 = F_lua_pushfstring(m, l0, v598+int32(_a2711), v11)
	mBase = m.M
	v602 = m.ExcPending
	if v602 != 0 {
		goto L38
	} else {
		goto L157
	}
L153:
	;
	if v584 == int32(67) {
		goto L151
	} else {
		goto L156
	}
L154:
	;
	v587 = m.G3
	v591 = F_lua_pushfstring(m, l0, v587+int32(_a2712), int32(0))
	mBase = m.M
	v592 = m.ExcPending
	if v592 != 0 {
		goto L38
	} else {
		goto L155
	}
L155:
	;
	goto L147
L156:
	;
	goto L152
L157:
	;
	goto L147
L158:
	;
	goto L147
L159:
	;
	F_lua_concat(m, l0, (v610-v611)>>(uint(int32(4))%32)-v136)
	mBase = m.M
	v617 = m.ExcPending
	if v617 != 0 {
		goto L38
	} else {
		goto L160
	}
L160:
	;
	v623 = *(*int32)(unsafe.Add(mBase, uint32(v137)+20))
	if v344 < int32(1) {
		v648 = v344
		v650 = v623
		goto L163
	} else {
		goto L164
	}
L161:
	;
	if v670 != 0 {
		v339 = v344
		goto L91
	} else {
		goto L174
	}
L162:
	;
	goto L161
L163:
	;
	if v648 != 0 {
		v661 = int32(0)
		goto L171
	} else {
		goto L172
	}
L164:
	;
	v626 = *(*int32)(unsafe.Add(mBase, uint32(v137)+40))
	v629 = v344
	v631 = v623
	goto L165
L165:
	;
	if base.Ui32(v631) <= base.Ui32(v626) {
		v670 = int32(0)
		goto L162
	} else {
		goto L167
	}
L166:
	;
	v648 = v642
	v650 = v644
	goto L163
L167:
	;
	v636 = v629 + int32(-1)
	v637 = *(*int32)(unsafe.Add(mBase, uint32(v631)+4))
	v638 = *(*int32)(unsafe.Add(mBase, uint32(v637)))
	v639 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v638)+6)))
	if v639 != 0 {
		v642 = v636
		goto L168
	} else {
		goto L169
	}
L168:
	;
	v644 = v631 + int32(-24)
	if int32(0) < v642 {
		v629 = v642
		v631 = v644
		goto L165
	} else {
		goto L170
	}
L169:
	;
	v640 = *(*int32)(unsafe.Add(mBase, uint32(v631)+20))
	v642 = v636 - v640
	goto L168
L170:
	;
	goto L166
L171:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11+int32(60))+96)) = v661
	v670 = int32(1)
	goto L162
L172:
	;
	v655 = *(*int32)(unsafe.Add(mBase, uint32(v137)+40))
	if base.Ui32(v650) <= base.Ui32(v655) {
		v670 = int32(0)
		goto L162
	} else {
		goto L173
	}
L173:
	;
	v659 = base.I32_div_s(v650-v655, int32(24))
	v661 = v659
	goto L171
L174:
	;
	goto L92
L175:
	;
	F_lua_concat(m, l0, (v679-v680)>>(uint(int32(4))%32)-v136)
	mBase = m.M
	v686 = m.ExcPending
	if v686 != 0 {
		goto L38
	} else {
		goto L176
	}
L176:
	;
	goto L50
}
func F_db_getfenv(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
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
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v78 int64
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
	F_luaL_checkany(m, l0, int32(1))
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
		v18 = v13 + int32(0)
		v19 = m.G398
		if base.Ui32(v18) < base.Ui32(v12) {
			v21 = v18
		} else {
			v21 = v19
		}
		v63 = *(*int32)(unsafe.Add(mBase, uint32(v21)+8))
		switch v63 + int32(-6) {
		case 0:
			v66 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			v67 = *(*int32)(unsafe.Add(mBase, uint32(v21)))
			v68 = *(*int32)(unsafe.Add(mBase, uint32(v67)+12))
			*(*int32)(unsafe.Add(mBase, uint32(v66))) = v68
			v83 = int32(5)
			v84 = v66
		case 1:
			v71 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			v72 = *(*int32)(unsafe.Add(mBase, uint32(v21)))
			v73 = *(*int32)(unsafe.Add(mBase, uint32(v72)+12))
			*(*int32)(unsafe.Add(mBase, uint32(v71))) = v73
			v83 = int32(5)
			v84 = v71
		case 2:
			v76 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			v77 = *(*int32)(unsafe.Add(mBase, uint32(v21)))
			v78 = *(*int64)(unsafe.Add(mBase, uint32(v77)+72))
			*(*int64)(unsafe.Add(mBase, uint32(v76))) = v78
			v80 = *(*int32)(unsafe.Add(mBase, uint32(v77)+80))
			v83 = v80
			v84 = v76
		default:
			v81 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			v83 = int32(0)
			v84 = v81
		}
		*(*int32)(unsafe.Add(mBase, uint32(v84)+8)) = v83
		v86 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v86 + int32(16)
		return int32(1)
	}
}
func F_db_gethook(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v13 int32
	_ = v13
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v126 int32
	_ = v126
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v136 int32
	_ = v136
	var v140 int32
	_ = v140
	var v147 int32
	_ = v147
	var v149 int32
	_ = v149
	var v151 int32
	_ = v151
	var v155 int32
	_ = v155
	var v176 int32
	_ = v176
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v223 int64
	_ = v223
	var v227 int32
	_ = v227
	var v246 int32
	_ = v246
	var v286 int32
	_ = v286
	var v287 int32
	_ = v287
	var v290 int32
	_ = v290
	var __phi290 int32
	_ = __phi290
	var v291 int32
	_ = v291
	var __phi291 int32
	_ = __phi291
	var v293 int64
	_ = v293
	var v295 int32
	_ = v295
	var v298 int32
	_ = v298
	var v299 int32
	_ = v299
	var v304 int32
	_ = v304
	var v308 int32
	_ = v308
	var v313 int32
	_ = v313
	var v316 int32
	_ = v316
	var v324 int32
	_ = v324
	var v328 int32
	_ = v328
	var v336 int32
	_ = v336
	var v340 int32
	_ = v340
	var v342 int32
	_ = v342
	var v344 int32
	_ = v344
	var v349 int32
	_ = v349
	var v350 int32
	_ = v350
	var v352 int32
	_ = v352
	var v357 int32
	_ = v357
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v18 = v13 + int32(0)
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if base.Ui32(v18) < base.Ui32(v19) {
		v61 = m.G398
		if v18 != v61 {
			v64 = *(*int32)(unsafe.Add(mBase, uint32(v18)+8))
			v67 = v64
		} else {
			v67 = int32(-1)
		}
	} else {
		v67 = int32(-1)
	}
	if v67 != int32(8) {
		v131 = l0
	} else {
		v74 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		v75 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
		v80 = v75 + int32(0)
		v81 = m.G398
		if base.Ui32(v80) < base.Ui32(v74) {
			v83 = v80
		} else {
			v83 = v81
		}
		v126 = *(*int32)(unsafe.Add(mBase, uint32(v83)+8))
		if v126 != int32(8) {
			v130 = int32(0)
		} else {
			v129 = *(*int32)(unsafe.Add(mBase, uint32(v83)))
			v130 = v129
		}
		v131 = v130
	}
	v132 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v131)+56)))
	v133 = *(*int32)(unsafe.Add(mBase, uint32(v131)+68))
	if v133 == int32(0) {
		F_gethooktable(m, l0)
		mBase = m.M
		v149 = m.ExcPending
		if v149 != 0 {
			return int32(0)
		} else {
			v151 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			*(*int32)(unsafe.Add(mBase, uint32(v151)+8)) = int32(2)
			*(*int32)(unsafe.Add(mBase, uint32(v151))) = v131
			v155 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v155 + int32(16)
			v176 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			v215 = *(*int32)(unsafe.Add(mBase, uint32(v176+int32(-32))))
			v216 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			v217 = int32(-16)
			v219 = F_luaH_get(m, v215, v216+v217)
			mBase = m.M
			v220 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			v223 = *(*int64)(unsafe.Add(mBase, uint32(v219)))
			*(*int64)(unsafe.Add(mBase, uint32(v220+v217))) = v223
			v227 = *(*int32)(unsafe.Add(mBase, uint32(v219)+8))
			*(*int32)(unsafe.Add(mBase, uint32(v220+int32(-8)))) = v227
			v246 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			v286 = v246 + int32(-16)
			v287 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			if base.Ui32(v287) <= base.Ui32(v286) {
				v304 = v287
			} else {
				__phi290 = v246 + int32(-32)
				__phi291 = v286
				v290 = __phi290
				v291 = __phi291
				for {
					v293 = *(*int64)(unsafe.Add(mBase, uint32(v290)+16))
					*(*int64)(unsafe.Add(mBase, uint32(v290))) = v293
					v295 = *(*int32)(unsafe.Add(mBase, uint32(v290)+24))
					*(*int32)(unsafe.Add(mBase, uint32(v290)+8)) = v295
					v298 = v291 + int32(16)
					v299 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
					if base.Ui32(v298) < base.Ui32(v299) {
						__phi290 = v291
						__phi291 = v298
						v290 = __phi290
						v291 = __phi291
						continue
					} else {
						break
					}
					break
				}
				v304 = v299
			}
			*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v304 + int32(-16)
			v308 = int32(0)
			if v132&int32(1) == v308 {
				v316 = v308
			} else {
				v313 = int32(99)
				*(*uint8)(unsafe.Add(mBase, uint32(v8)+11)) = uint8(v313)
				v316 = int32(1)
			}
			if v132&int32(2) == int32(0) {
				v328 = v316
			} else {
				v324 = int32(114)
				*(*uint8)(unsafe.Add(mBase, uint32(v8+int32(11)+v316))) = uint8(v324)
				v328 = v316 + int32(1)
			}
			if v132&int32(4) == int32(0) {
				v340 = v328
			} else {
				v336 = int32(108)
				*(*uint8)(unsafe.Add(mBase, uint32(v8+int32(11)+v328))) = uint8(v336)
				v340 = v328 + int32(1)
			}
			v342 = v8 + int32(11)
			v344 = int32(0)
			*(*uint8)(unsafe.Add(mBase, uint32(v342+v340))) = uint8(v344)
			F_lua_pushstring(m, l0, v342)
			mBase = m.M
			v349 = m.ExcPending
			if v349 != 0 {
				return int32(0)
			} else {
				v350 = *(*int32)(unsafe.Add(mBase, uint32(v131)+60))
				v352 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				*(*int32)(unsafe.Add(mBase, uint32(v352)+8)) = int32(3)
				*(*float64)(unsafe.Add(mBase, uint32(v352))) = base.F64_convert_i32_s(v350)
				v357 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v357 + int32(16)
				m.G0 = v8 + int32(16)
				return int32(3)
			}
		}
	} else {
		v136 = m.G5
		if v133 == v136+int32(1274) {
			F_gethooktable(m, l0)
			mBase = m.M
			v149 = m.ExcPending
			if v149 != 0 {
				return int32(0)
			} else {
				v151 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				*(*int32)(unsafe.Add(mBase, uint32(v151)+8)) = int32(2)
				*(*int32)(unsafe.Add(mBase, uint32(v151))) = v131
				v155 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v155 + int32(16)
				v176 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				v215 = *(*int32)(unsafe.Add(mBase, uint32(v176+int32(-32))))
				v216 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				v217 = int32(-16)
				v219 = F_luaH_get(m, v215, v216+v217)
				mBase = m.M
				v220 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				v223 = *(*int64)(unsafe.Add(mBase, uint32(v219)))
				*(*int64)(unsafe.Add(mBase, uint32(v220+v217))) = v223
				v227 = *(*int32)(unsafe.Add(mBase, uint32(v219)+8))
				*(*int32)(unsafe.Add(mBase, uint32(v220+int32(-8)))) = v227
				v246 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				v286 = v246 + int32(-16)
				v287 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				if base.Ui32(v287) <= base.Ui32(v286) {
					v304 = v287
				} else {
					__phi290 = v246 + int32(-32)
					__phi291 = v286
					v290 = __phi290
					v291 = __phi291
					for {
						v293 = *(*int64)(unsafe.Add(mBase, uint32(v290)+16))
						*(*int64)(unsafe.Add(mBase, uint32(v290))) = v293
						v295 = *(*int32)(unsafe.Add(mBase, uint32(v290)+24))
						*(*int32)(unsafe.Add(mBase, uint32(v290)+8)) = v295
						v298 = v291 + int32(16)
						v299 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
						if base.Ui32(v298) < base.Ui32(v299) {
							__phi290 = v291
							__phi291 = v298
							v290 = __phi290
							v291 = __phi291
							continue
						} else {
							break
						}
						break
					}
					v304 = v299
				}
				*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v304 + int32(-16)
				v308 = int32(0)
				if v132&int32(1) == v308 {
					v316 = v308
				} else {
					v313 = int32(99)
					*(*uint8)(unsafe.Add(mBase, uint32(v8)+11)) = uint8(v313)
					v316 = int32(1)
				}
				if v132&int32(2) == int32(0) {
					v328 = v316
				} else {
					v324 = int32(114)
					*(*uint8)(unsafe.Add(mBase, uint32(v8+int32(11)+v316))) = uint8(v324)
					v328 = v316 + int32(1)
				}
				if v132&int32(4) == int32(0) {
					v340 = v328
				} else {
					v336 = int32(108)
					*(*uint8)(unsafe.Add(mBase, uint32(v8+int32(11)+v328))) = uint8(v336)
					v340 = v328 + int32(1)
				}
				v342 = v8 + int32(11)
				v344 = int32(0)
				*(*uint8)(unsafe.Add(mBase, uint32(v342+v340))) = uint8(v344)
				F_lua_pushstring(m, l0, v342)
				mBase = m.M
				v349 = m.ExcPending
				if v349 != 0 {
					return int32(0)
				} else {
					v350 = *(*int32)(unsafe.Add(mBase, uint32(v131)+60))
					v352 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
					*(*int32)(unsafe.Add(mBase, uint32(v352)+8)) = int32(3)
					*(*float64)(unsafe.Add(mBase, uint32(v352))) = base.F64_convert_i32_s(v350)
					v357 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
					*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v357 + int32(16)
					m.G0 = v8 + int32(16)
					return int32(3)
				}
			}
		} else {
			v140 = m.G3
			F_lua_pushlstring(m, l0, v140+int32(_a2699), int32(13))
			mBase = m.M
			v147 = m.ExcPending
			if v147 != 0 {
				return int32(0)
			} else {
				v308 = int32(0)
				if v132&int32(1) == v308 {
					v316 = v308
				} else {
					v313 = int32(99)
					*(*uint8)(unsafe.Add(mBase, uint32(v8)+11)) = uint8(v313)
					v316 = int32(1)
				}
				if v132&int32(2) == int32(0) {
					v328 = v316
				} else {
					v324 = int32(114)
					*(*uint8)(unsafe.Add(mBase, uint32(v8+int32(11)+v316))) = uint8(v324)
					v328 = v316 + int32(1)
				}
				if v132&int32(4) == int32(0) {
					v340 = v328
				} else {
					v336 = int32(108)
					*(*uint8)(unsafe.Add(mBase, uint32(v8+int32(11)+v328))) = uint8(v336)
					v340 = v328 + int32(1)
				}
				v342 = v8 + int32(11)
				v344 = int32(0)
				*(*uint8)(unsafe.Add(mBase, uint32(v342+v340))) = uint8(v344)
				F_lua_pushstring(m, l0, v342)
				mBase = m.M
				v349 = m.ExcPending
				if v349 != 0 {
					return int32(0)
				} else {
					v350 = *(*int32)(unsafe.Add(mBase, uint32(v131)+60))
					v352 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
					*(*int32)(unsafe.Add(mBase, uint32(v352)+8)) = int32(3)
					*(*float64)(unsafe.Add(mBase, uint32(v352))) = base.F64_convert_i32_s(v350)
					v357 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
					*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v357 + int32(16)
					m.G0 = v8 + int32(16)
					return int32(3)
				}
			}
		}
	}
}
func F_db_getlocal(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v14 int32
	_ = v14
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v127 int32
	_ = v127
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v140 int32
	_ = v140
	var v146 int32
	_ = v146
	var v149 int32
	_ = v149
	var v152 int32
	_ = v152
	var v154 int32
	_ = v154
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v165 int32
	_ = v165
	var v167 int32
	_ = v167
	var v171 int32
	_ = v171
	var v173 int32
	_ = v173
	var v178 int32
	_ = v178
	var v182 int32
	_ = v182
	var v184 int32
	_ = v184
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
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
	var v226 int32
	_ = v226
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v238 int32
	_ = v238
	var v241 int32
	_ = v241
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
	var v255 int32
	_ = v255
	var v264 int32
	_ = v264
	var v265 int32
	_ = v265
	var v275 int32
	_ = v275
	var v288 int32
	_ = v288
	var v355 int32
	_ = v355
	var v356 int32
	_ = v356
	var v362 int32
	_ = v362
	var v363 int64
	_ = v363
	var v365 int32
	_ = v365
	var v377 int32
	_ = v377
	var v395 int32
	_ = v395
	var v398 int32
	_ = v398
	var v434 int32
	_ = v434
	var v435 int64
	_ = v435
	var v437 int32
	_ = v437
	var v439 int32
	_ = v439
	var v444 int32
	_ = v444
	var v452 int32
	_ = v452
	v6 = m.G0
	v8 = v6 - int32(112)
	m.G0 = v8
	goto L5
L1:
	;
	v136 = v133 + int32(1)
	v137 = F_luaL_checkinteger(m, l0, v136)
	mBase = m.M
	v140 = m.ExcPending
	if v140 != 0 {
		goto L38
	} else {
		goto L39
	}
L2:
	;
	if v68 != int32(8) {
		v133 = int32(0)
		v134 = l0
		goto L1
	} else {
		goto L17
	}
L3:
	;
	v62 = m.G398
	if v19 != v62 {
		goto L15
	} else {
		goto L16
	}
L5:
	;
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v19 = v14 + int32(0)
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if base.Ui32(v19) < base.Ui32(v20) {
		goto L3
	} else {
		goto L6
	}
L6:
	;
	v68 = int32(-1)
	goto L2
L15:
	;
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v19)+8))
	v68 = v65
	goto L2
L16:
	;
	v68 = int32(-1)
	goto L2
L17:
	;
	goto L21
L18:
	;
	v133 = int32(1)
	v134 = v131
	goto L1
L19:
	;
	v127 = *(*int32)(unsafe.Add(mBase, uint32(v84)+8))
	if v127 != int32(8) {
		v131 = int32(0)
		goto L34
	} else {
		goto L35
	}
L21:
	;
	v75 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v76 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v81 = v76 + int32(0)
	v82 = m.G398
	if base.Ui32(v81) < base.Ui32(v75) {
		goto L22
	} else {
		goto L23
	}
L22:
	;
	v84 = v81
	goto L24
L23:
	;
	v84 = v82
	goto L24
L24:
	;
	goto L19
L34:
	;
	goto L18
L35:
	;
	v130 = *(*int32)(unsafe.Add(mBase, uint32(v84)))
	v131 = v130
	goto L34
L36:
	;
	m.G0 = v8 + int32(112)
	return v452
L37:
	;
	v199 = int32(2)
	v204 = F_luaL_checkinteger(m, l0, v133|v199)
	mBase = m.M
	v205 = m.ExcPending
	if v205 != 0 {
		goto L38
	} else {
		goto L56
	}
L38:
	;
	return int32(0)
L39:
	;
	v146 = *(*int32)(unsafe.Add(mBase, uint32(v134)+20))
	if v137 < int32(1) {
		v171 = v137
		v173 = v146
		goto L42
	} else {
		goto L43
	}
L40:
	;
	if v193 != 0 {
		goto L37
	} else {
		goto L53
	}
L41:
	;
	goto L40
L42:
	;
	if v171 != 0 {
		v184 = int32(0)
		goto L50
	} else {
		goto L51
	}
L43:
	;
	v149 = *(*int32)(unsafe.Add(mBase, uint32(v134)+40))
	v152 = v137
	v154 = v146
	goto L44
L44:
	;
	if base.Ui32(v154) <= base.Ui32(v149) {
		v193 = int32(0)
		goto L41
	} else {
		goto L46
	}
L45:
	;
	v171 = v165
	v173 = v167
	goto L42
L46:
	;
	v159 = v152 + int32(-1)
	v160 = *(*int32)(unsafe.Add(mBase, uint32(v154)+4))
	v161 = *(*int32)(unsafe.Add(mBase, uint32(v160)))
	v162 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v161)+6)))
	if v162 != 0 {
		v165 = v159
		goto L47
	} else {
		goto L48
	}
L47:
	;
	v167 = v154 + int32(-24)
	if int32(0) < v165 {
		v152 = v165
		v154 = v167
		goto L44
	} else {
		goto L49
	}
L48:
	;
	v163 = *(*int32)(unsafe.Add(mBase, uint32(v154)+20))
	v165 = v159 - v163
	goto L47
L49:
	;
	goto L45
L50:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8+int32(12))+96)) = v184
	v193 = int32(1)
	goto L41
L51:
	;
	v178 = *(*int32)(unsafe.Add(mBase, uint32(v134)+40))
	if base.Ui32(v173) <= base.Ui32(v178) {
		v193 = int32(0)
		goto L41
	} else {
		goto L52
	}
L52:
	;
	v182 = base.I32_div_s(v173-v178, int32(24))
	v184 = v182
	goto L50
L53:
	;
	v194 = m.G3
	v197 = F_luaL_argerror(m, l0, v136, v194+int32(_a2702))
	mBase = m.M
	v198 = m.ExcPending
	if v198 != 0 {
		goto L38
	} else {
		goto L54
	}
L54:
	;
	v452 = v197
	goto L36
L55:
	;
	v444 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v444)+8)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v444 + int32(16)
	goto L101
L56:
	;
	v209 = *(*int32)(unsafe.Add(mBase, uint32(v134)+40))
	v210 = *(*int32)(unsafe.Add(mBase, uint32(v8+int32(12))+96))
	v213 = v209 + v210*int32(24)
	v214 = *(*int32)(unsafe.Add(mBase, uint32(v213)+4))
	v215 = *(*int32)(unsafe.Add(mBase, uint32(v214)+8))
	if v215 != int32(6) {
		goto L60
	} else {
		goto L61
	}
L57:
	;
	if v275 == int32(0) {
		goto L55
	} else {
		goto L73
	}
L58:
	;
	goto L57
L59:
	;
	F_luaA_pushobject(m, v134, v264+v204<<(uint(int32(4))%32)+int32(-16))
	mBase = m.M
	v275 = v265
	goto L58
L60:
	;
	v249 = *(*int32)(unsafe.Add(mBase, uint32(v134)+20))
	if v213 == v249 {
		goto L68
	} else {
		goto L69
	}
L61:
	;
	v218 = *(*int32)(unsafe.Add(mBase, uint32(v214)))
	v219 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v218)+6)))
	if v219 != 0 {
		goto L60
	} else {
		goto L62
	}
L62:
	;
	v220 = *(*int32)(unsafe.Add(mBase, uint32(v218)+16))
	if v220 == int32(0) {
		goto L60
	} else {
		goto L63
	}
L63:
	;
	v223 = *(*int32)(unsafe.Add(mBase, uint32(v134)+20))
	if v213 == v223 {
		goto L65
	} else {
		goto L66
	}
L64:
	;
	v232 = *(*int32)(unsafe.Add(mBase, uint32(v230)+12))
	v238 = F_luaF_getlocalname(m, v220, v204, (v231-v232)>>(uint(int32(2))%32)+int32(-1))
	mBase = m.M
	if v238 == int32(0) {
		goto L60
	} else {
		goto L67
	}
L65:
	;
	v226 = *(*int32)(unsafe.Add(mBase, uint32(v134)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v213)+12)) = v226
	v228 = *(*int32)(unsafe.Add(mBase, uint32(v214)))
	v229 = *(*int32)(unsafe.Add(mBase, uint32(v228)+16))
	v230 = v229
	v231 = v226
	goto L64
L66:
	;
	v225 = *(*int32)(unsafe.Add(mBase, uint32(v213)+12))
	v230 = v220
	v231 = v225
	goto L64
L67:
	;
	v241 = *(*int32)(unsafe.Add(mBase, uint32(v213)))
	v264 = v241
	v265 = v238
	goto L59
L68:
	;
	v251 = v134 + int32(8)
	goto L70
L69:
	;
	v251 = v213 + int32(28)
	goto L70
L70:
	;
	v252 = *(*int32)(unsafe.Add(mBase, uint32(v251)))
	v253 = *(*int32)(unsafe.Add(mBase, uint32(v213)))
	v254 = m.G3
	v255 = int32(0)
	if v204 < int32(1) {
		v275 = v255
		goto L58
	} else {
		goto L71
	}
L71:
	;
	if (v252-v253)>>(uint(int32(4))%32) < v204 {
		v275 = v255
		goto L58
	} else {
		goto L72
	}
L72:
	;
	v264 = v253
	v265 = v254 + int32(_a2590)
	goto L59
L73:
	;
	if v134 == l0 {
		goto L75
	} else {
		goto L76
	}
L74:
	;
	F_lua_pushstring(m, l0, v275)
	mBase = m.M
	v377 = m.ExcPending
	if v377 != 0 {
		goto L38
	} else {
		goto L84
	}
L75:
	;
	goto L74
L76:
	;
	v288 = *(*int32)(unsafe.Add(mBase, uint32(v134)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v134)+8)) = v288 - int32(16)
	goto L77
L77:
	;
	goto L78
L78:
	;
	goto L83
L83:
	;
	v355 = *(*int32)(unsafe.Add(mBase, uint32(v134)+8))
	v356 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v356 + int32(16)
	v362 = v355 + int32(0)
	v363 = *(*int64)(unsafe.Add(mBase, uint32(v362)))
	*(*int64)(unsafe.Add(mBase, uint32(v356))) = v363
	v365 = *(*int32)(unsafe.Add(mBase, uint32(v362)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v356)+8)) = v365
	goto L75
L84:
	;
	goto L87
L85:
	;
	v452 = v199
	goto L36
L86:
	;
	v434 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v435 = *(*int64)(unsafe.Add(mBase, uint32(v398)))
	*(*int64)(unsafe.Add(mBase, uint32(v434))) = v435
	v437 = *(*int32)(unsafe.Add(mBase, uint32(v398)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v434)+8)) = v437
	v439 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v439 + int32(16)
	goto L85
L87:
	;
	goto L93
L93:
	;
	v395 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v398 = v395 + int32(-32)
	goto L86
L101:
	;
	v452 = int32(1)
	goto L36
}
func F_db_getregistry(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
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
	var v59 int64
	_ = v59
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	switch int32(2) {
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
		if base.Ui32(v42) < base.Ui32(int32(-2)) {
			v54 = v43
		} else {
			v54 = v41 + int32(-24)
		}
		v55 = v54
	}
	v58 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v59 = *(*int64)(unsafe.Add(mBase, uint32(v55)))
	*(*int64)(unsafe.Add(mBase, uint32(v58))) = v59
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v55)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v58)+8)) = v61
	v63 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v63 + int32(16)
	return int32(1)
}
func F_db_setmetatable(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v80 int32
	_ = v80
	var v84 int32
	_ = v84
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v106 int32
	_ = v106
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v10 = v5 + int32(16)
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if base.Ui32(v10) < base.Ui32(v11) {
		v53 = m.G398
		if v10 != v53 {
			v56 = *(*int32)(unsafe.Add(mBase, uint32(v10)+8))
			v59 = v56
		} else {
			v59 = int32(-1)
		}
	} else {
		v59 = int32(-1)
	}
	switch v59 {
	case 0, 5:
		v72 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		v73 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
		v76 = v73 + int32(32)
		if base.Ui32(v76) <= base.Ui32(v72) {
		} else {
			v80 = v72
			for {
				*(*int32)(unsafe.Add(mBase, uint32(v80)+8)) = int32(0)
				v84 = v80 + int32(16)
				if base.Ui32(v84) < base.Ui32(v76) {
					v80 = v84
					continue
				} else {
					break
				}
				break
			}
		}
		*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v76
		v97 = F_lua_setmetatable(m, l0, int32(1))
		mBase = m.M
		v98 = m.ExcPending
		if v98 != 0 {
			return int32(0)
		} else {
			v100 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			*(*int32)(unsafe.Add(mBase, uint32(v100)+8)) = int32(1)
			*(*int32)(unsafe.Add(mBase, uint32(v100))) = base.B2i32(v97 != int32(0))
			v106 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v106 + int32(16)
			return int32(1)
		}
	default:
		v61 = m.G3
		v64 = F_luaL_argerror(m, l0, int32(2), v61+int32(_a2703))
		mBase = m.M
		v67 = m.ExcPending
		if v67 != 0 {
			return int32(0)
		} else {
			v72 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			v73 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
			v76 = v73 + int32(32)
			if base.Ui32(v76) <= base.Ui32(v72) {
			} else {
				v80 = v72
				for {
					*(*int32)(unsafe.Add(mBase, uint32(v80)+8)) = int32(0)
					v84 = v80 + int32(16)
					if base.Ui32(v84) < base.Ui32(v76) {
						v80 = v84
						continue
					} else {
						break
					}
					break
				}
			}
			*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v76
			v97 = F_lua_setmetatable(m, l0, int32(1))
			mBase = m.M
			v98 = m.ExcPending
			if v98 != 0 {
				return int32(0)
			} else {
				v100 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				*(*int32)(unsafe.Add(mBase, uint32(v100)+8)) = int32(1)
				*(*int32)(unsafe.Add(mBase, uint32(v100))) = base.B2i32(v97 != int32(0))
				v106 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v106 + int32(16)
				return int32(1)
			}
		}
	}
}
func F_resetDbExpiryState(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v2 int64
	_ = v2
	v2 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(l0)+32)) = v2
	*(*int64)(unsafe.Add(mBase, uint32(l0+int32(56)))) = v2
	*(*int64)(unsafe.Add(mBase, uint32(l0+int32(48)))) = v2
	*(*int64)(unsafe.Add(mBase, uint32(l0+int32(40)))) = v2
	return
}
func F_selectDbIdArgs(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
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
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v22 int64
	_ = v22
	var v26 int64
	_ = v26
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v35 int32
	_ = v35
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	if l1 < int32(2) {
		v35 = int32(0)
		m.G0 = v9 + int32(16)
		return v35
	} else {
		v14 = int32(0)
		v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		v18 = F_getLongLongFromObject(m, v15, v9+int32(8))
		mBase = m.M
		v21 = m.ExcPending
		if v21 != 0 {
			return int32(0)
		} else {
			if v18 != 0 {
				v35 = v14
				m.G0 = v9 + int32(16)
				return v35
			} else {
				v22 = *(*int64)(unsafe.Add(mBase, uint32(v9)+8))
				if v22 < int64(0) {
					v35 = v14
					m.G0 = v9 + int32(16)
					return v35
				} else {
					v26 = int64(*(*int32)(unsafe.Add(mBase, _consts[172])))
					if v26 <= v22 {
						v35 = v14
						m.G0 = v9 + int32(16)
						return v35
					} else {
						v29 = F_valkey_malloc(m, int32(4))
						mBase = m.M
						v30 = m.ExcPending
						if v30 != 0 {
							return int32(0)
						} else {
							v31 = int32(1)
							*(*int32)(unsafe.Add(mBase, uint32(v29))) = v31
							*(*int32)(unsafe.Add(mBase, uint32(l2))) = v31
							v35 = v29
							m.G0 = v9 + int32(16)
							return v35
						}
					}
				}
			}
		}
	}
}
