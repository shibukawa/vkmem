package p2

import (
	base "github.com/shibukawa/vkmem/internal/aot/vkaot/base"
	"unsafe"
)

func F_aofInfoFree(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v16 int32
	_ = v16
	if l0 == int32(0) {
		F__serverAssert(m, int32(_a_F_aofInfoFree_0), int32(_a_F_aofInfoFree_1), int32(105))
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
	} else {
		v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		if v5 == int32(0) {
			F_valkey_free(m, l0)
			mBase = m.M
			v11 = m.ExcPending
			if v11 != 0 {
				return
			} else {
				return
			}
		} else {
			F_sdsfree(m, v5)
			mBase = m.M
			v9 = m.ExcPending
			if v9 != 0 {
				return
			} else {
				F_valkey_free(m, l0)
				mBase = m.M
				v11 = m.ExcPending
				if v11 != 0 {
					return
				} else {
					return
				}
			}
		}
	}
}
func F_aofLoadManifestFromFile(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
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
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v52 int32
	_ = v52
	var v57 int32
	_ = v57
	var v63 int64
	_ = v63
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
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
	var v95 int32
	_ = v95
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v112 int32
	_ = v112
	var v115 int32
	_ = v115
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v134 int32
	_ = v134
	var v137 int32
	_ = v137
	var v140 int32
	_ = v140
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v147 int32
	_ = v147
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v159 int32
	_ = v159
	var v165 int32
	_ = v165
	var v169 int32
	_ = v169
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v177 int32
	_ = v177
	var v181 int32
	_ = v181
	var v186 int32
	_ = v186
	var v188 int32
	_ = v188
	var v190 int32
	_ = v190
	var v192 int32
	_ = v192
	var v196 int32
	_ = v196
	var v213 int32
	_ = v213
	var v220 int32
	_ = v220
	var v223 int32
	_ = v223
	var v226 int32
	_ = v226
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v260 int32
	_ = v260
	var v263 int32
	_ = v263
	var v267 int32
	_ = v267
	var v268 int32
	_ = v268
	var v269 int32
	_ = v269
	var v272 int32
	_ = v272
	var v274 int32
	_ = v274
	var v275 int32
	_ = v275
	var v276 int32
	_ = v276
	var v278 int32
	_ = v278
	var v282 int32
	_ = v282
	var v284 int32
	_ = v284
	var v285 int32
	_ = v285
	var v286 int32
	_ = v286
	var v288 int32
	_ = v288
	var v289 int32
	_ = v289
	var v290 int32
	_ = v290
	var v291 int32
	_ = v291
	var v296 int32
	_ = v296
	var v297 int32
	_ = v297
	var v300 int32
	_ = v300
	var v301 int32
	_ = v301
	var v304 int32
	_ = v304
	var v305 int32
	_ = v305
	var v306 int32
	_ = v306
	var v310 int32
	_ = v310
	var v311 int32
	_ = v311
	var v312 int32
	_ = v312
	var v317 int32
	_ = v317
	var v319 int32
	_ = v319
	var v322 int32
	_ = v322
	var v323 int32
	_ = v323
	var v325 int32
	_ = v325
	var v328 int32
	_ = v328
	var v330 int32
	_ = v330
	var v331 int32
	_ = v331
	var v332 int32
	_ = v332
	var v334 int32
	_ = v334
	var v338 int32
	_ = v338
	var v340 int32
	_ = v340
	var v341 int32
	_ = v341
	var v342 int32
	_ = v342
	var v344 int32
	_ = v344
	var v345 int32
	_ = v345
	var v346 int32
	_ = v346
	var v347 int32
	_ = v347
	var v352 int32
	_ = v352
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
	var v362 int32
	_ = v362
	var v366 int32
	_ = v366
	var v371 int32
	_ = v371
	var v377 int32
	_ = v377
	var v378 int32
	_ = v378
	var v379 int32
	_ = v379
	var v380 int32
	_ = v380
	var v386 int32
	_ = v386
	var v387 int32
	_ = v387
	var v388 int32
	_ = v388
	var v389 int32
	_ = v389
	var v390 int32
	_ = v390
	var v393 int32
	_ = v393
	var v397 int32
	_ = v397
	var v398 int32
	_ = v398
	var v401 int64
	_ = v401
	var v405 int64
	_ = v405
	var v406 int32
	_ = v406
	var v410 int32
	_ = v410
	var v417 int64
	_ = v417
	var v420 int64
	_ = v420
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
	var v463 int32
	_ = v463
	var v464 int32
	_ = v464
	var v467 int32
	_ = v467
	var v469 int32
	_ = v469
	var v483 int32
	_ = v483
	var v485 int32
	_ = v485
	var v486 int32
	_ = v486
	var v489 int64
	_ = v489
	var v492 int32
	_ = v492
	var v496 int32
	_ = v496
	var v498 int32
	_ = v498
	var v501 int32
	_ = v501
	var v506 int32
	_ = v506
	var v508 int64
	_ = v508
	var v511 int32
	_ = v511
	var v512 int32
	_ = v512
	var v513 int32
	_ = v513
	var v515 int32
	_ = v515
	var v516 int64
	_ = v516
	var v520 int32
	_ = v520
	var v521 int32
	_ = v521
	var v522 int32
	_ = v522
	var v523 int64
	_ = v523
	var v526 int32
	_ = v526
	var v527 int32
	_ = v527
	var v528 int32
	_ = v528
	var v536 int32
	_ = v536
	var v545 int32
	_ = v545
	var v549 int32
	_ = v549
	var v552 int32
	_ = v552
	var v561 int32
	_ = v561
	var v568 int32
	_ = v568
	var v569 int32
	_ = v569
	var v576 int32
	_ = v576
	var v577 int32
	_ = v577
	var v584 int32
	_ = v584
	var v591 int32
	_ = v591
	var v593 int32
	_ = v593
	var v594 int32
	_ = v594
	var v605 int32
	_ = v605
	var v607 int32
	_ = v607
	var v616 int32
	_ = v616
	var v618 int32
	_ = v618
	var v619 int32
	_ = v619
	var v628 int32
	_ = v628
	v15 = m.G0
	v17 = v15 - int32(1104)
	m.G0 = v17
	v20 = F_valkey_calloc(m, int32(40))
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v24 = F_listCreate(m)
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+4)) = v24
	v27 = F_listCreate(m)
	mBase = m.M
	v28 = m.ExcPending
	if v28 != 0 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+8)) = v27
	v30 = int32(15)
	*(*int32)(unsafe.Add(mBase, uint32(v24)+8)) = v30
	v32 = int32(16)
	*(*int32)(unsafe.Add(mBase, uint32(v24)+12)) = v32
	*(*int32)(unsafe.Add(mBase, uint32(v27)+8)) = v30
	*(*int32)(unsafe.Add(mBase, uint32(v27)+12)) = v32
	v39 = F_fopen(m, l0, int32(_a_F_aofLoadManifestFromFile_0))
	mBase = m.M
	if v39 != 0 {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	v57 = int32(0)
	v63 = int64(0)
	goto L17
L6:
	;
	v41 = *(*int32)(unsafe.Add(mBase, _c_F_aofLoadManifestFromFile[0]))
	if int32(3) < v41 {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	m.Env.Exit(m, int32(1))
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L8:
	;
	goto L9
L9:
	;
	v45 = *(*int32)(unsafe.Add(mBase, _c_F_aofLoadManifestFromFile[1]))
	v46 = F___strerror_l(m, v45, v45)
	mBase = m.M
	goto L10
L10:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+4)) = v46
	*(*int32)(unsafe.Add(mBase, uint32(v17))) = l0
	F__serverLog(m, int32(3), int32(_a_F_aofLoadManifestFromFile_1), v17)
	mBase = m.M
	v52 = m.ExcPending
	if v52 != 0 {
		goto L1
	} else {
		goto L11
	}
L11:
	;
	goto L7
L12:
	;
	v584 = *(*int32)(unsafe.Add(mBase, _c_F_aofLoadManifestFromFile[0]))
	if int32(3) < v584 {
		v594 = v584
		goto L167
	} else {
		goto L168
	}
L13:
	;
	F_aofInfoFree(m, v245)
	mBase = m.M
	v568 = m.ExcPending
	if v568 != 0 {
		goto L1
	} else {
		goto L166
	}
L14:
	;
	F_sdsfreesplitres(m, v235, v549)
	mBase = m.M
	v552 = m.ExcPending
	if v552 != 0 {
		goto L1
	} else {
		goto L165
	}
L15:
	;
	if v235 == int32(0) {
		v569 = v99
		v576 = v118
		v577 = v115
		goto L12
	} else {
		goto L163
	}
L16:
	;
	v527 = F_fclose(m, v39)
	mBase = m.M
	v528 = m.ExcPending
	if v528 != 0 {
		goto L1
	} else {
		goto L162
	}
L17:
	;
	v74 = F_fgets(m, v17+int32(64), int32(1025), v39)
	mBase = m.M
	v75 = m.ExcPending
	if v75 != 0 {
		goto L1
	} else {
		goto L20
	}
L19:
	;
	v99 = v57 + int32(1)
	v100 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17)+64)))
	if v100 == int32(35) {
		v57 = v99
		goto L17
	} else {
		goto L30
	}
L20:
	;
	if v74 != 0 {
		goto L19
	} else {
		goto L21
	}
L21:
	;
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v39)+76))
	if int32(-1) < v78 {
		goto L25
	} else {
		goto L26
	}
L22:
	;
	if v57 != 0 {
		goto L16
	} else {
		goto L29
	}
L23:
	;
	if int32(base.Ui32(v87)>>(uint(int32(4))%32))&int32(1) != 0 {
		goto L22
	} else {
		goto L28
	}
L24:
	;
	goto L23
L25:
	;
	v82 = F___lockfile(m, v39)
	mBase = m.M
	v83 = *(*int32)(unsafe.Add(mBase, uint32(v39)))
	if v82 == int32(0) {
		v87 = v83
		goto L24
	} else {
		goto L27
	}
L26:
	;
	v81 = *(*int32)(unsafe.Add(mBase, uint32(v39)))
	v87 = v81
	goto L24
L27:
	;
	F___unlockfile(m, v39)
	mBase = m.M
	v87 = v83
	goto L24
L28:
	;
	v569 = v57
	v576 = int32(0)
	v577 = int32(_a_F_aofLoadManifestFromFile_2)
	goto L12
L29:
	;
	v95 = int32(0)
	v569 = v95
	v576 = v95
	v577 = int32(_a_F_aofLoadManifestFromFile_3)
	goto L12
L30:
	;
	v105 = int32(10)
	v106 = F___strchrnul(m, v17+int32(64), v105)
	mBase = m.M
	v108 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v106))))
	if v108 == v105 {
		goto L33
	} else {
		goto L34
	}
L31:
	;
	v115 = int32(_a_F_aofLoadManifestFromFile_4)
	v118 = F_sdsnew(m, v17+int32(64))
	mBase = m.M
	v119 = m.ExcPending
	if v119 != 0 {
		goto L1
	} else {
		goto L43
	}
L32:
	;
	if v112 != 0 {
		goto L31
	} else {
		goto L36
	}
L33:
	;
	v112 = v106
	goto L35
L34:
	;
	v112 = int32(0)
	goto L35
L35:
	;
	goto L32
L36:
	;
	v569 = v99
	v576 = int32(0)
	v577 = int32(_a_F_aofLoadManifestFromFile_5)
	goto L12
L37:
	;
	if v230 == int32(0) {
		v569 = v99
		v576 = v118
		v577 = v115
		goto L12
	} else {
		goto L71
	}
L38:
	;
	v229 = *(*int32)(unsafe.Add(mBase, uint32(v118+int32(-17))))
	v230 = v229
	goto L37
L39:
	;
	v226 = *(*int32)(unsafe.Add(mBase, uint32(v118+int32(-9))))
	v230 = v226
	goto L37
L40:
	;
	v223 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v118+int32(-5)))))
	v230 = v223
	goto L37
L41:
	;
	v220 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v118+int32(-3)))))
	v230 = v220
	goto L37
L42:
	;
	v230 = int32(base.Ui32(v213) >> (uint(int32(3)) % 32))
	goto L37
L43:
	;
	v120 = int32(_a_F_aofLoadManifestFromFile_6)
	v126 = v118 + int32(-1)
	v127 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v126))))
	switch v127 & int32(7) {
	case 0:
		goto L50
	case 1:
		goto L49
	case 2:
		goto L48
	case 3:
		goto L47
	case 4:
		goto L46
	default:
		v144 = int32(0)
		goto L45
	}
L44:
	;
	v213 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v118+int32(-1)))))
	switch v213 & int32(7) {
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
		v569 = v99
		v576 = v118
		v577 = v115
		goto L12
	}
L45:
	;
	v147 = v118 + v144 + int32(-1)
	if base.Ui32(v147) < base.Ui32(v118) {
		v165 = v118
		goto L51
	} else {
		goto L52
	}
L46:
	;
	v143 = *(*int32)(unsafe.Add(mBase, uint32(v118+int32(-17))))
	v144 = v143
	goto L45
L47:
	;
	v140 = *(*int32)(unsafe.Add(mBase, uint32(v118+int32(-9))))
	v144 = v140
	goto L45
L48:
	;
	v137 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v118+int32(-5)))))
	v144 = v137
	goto L45
L49:
	;
	v134 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v118+int32(-3)))))
	v144 = v134
	goto L45
L50:
	;
	v144 = int32(base.Ui32(v127) >> (uint(int32(3)) % 32))
	goto L45
L51:
	;
	if base.Ui32(v147) <= base.Ui32(v165) {
		v181 = v147
		goto L57
	} else {
		goto L58
	}
L52:
	;
	v153 = v118
	goto L53
L53:
	;
	v154 = int32(*(*int8)(unsafe.Add(mBase, uint32(v153))))
	v155 = F_strchr(m, v120, v154)
	mBase = m.M
	if v155 == int32(0) {
		v165 = v153
		goto L51
	} else {
		goto L55
	}
L54:
	;
	v165 = v159
	goto L51
L55:
	;
	v159 = v153 + int32(1)
	if base.Ui32(v159) <= base.Ui32(v147) {
		v153 = v159
		goto L53
	} else {
		goto L56
	}
L56:
	;
	goto L54
L57:
	;
	v186 = v181 - v165 + int32(1)
	if v118 == v165 {
		goto L63
	} else {
		goto L64
	}
L58:
	;
	v169 = v147
	goto L59
L59:
	;
	v172 = int32(*(*int8)(unsafe.Add(mBase, uint32(v169))))
	v173 = F_strchr(m, v120, v172)
	mBase = m.M
	if v173 == int32(0) {
		v181 = v169
		goto L57
	} else {
		goto L61
	}
L60:
	;
	v181 = v165
	goto L57
L61:
	;
	v177 = v169 + int32(-1)
	if base.Ui32(v165) < base.Ui32(v177) {
		v169 = v177
		goto L59
	} else {
		goto L62
	}
L62:
	;
	goto L60
L63:
	;
	v190 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v118+v186))) = uint8(v190)
	v192 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v126))))
	switch v192 & int32(7) {
	case 0:
		goto L70
	case 1:
		goto L69
	case 2:
		goto L68
	case 3:
		goto L67
	case 4:
		goto L66
	default:
		goto L65
	}
L64:
	;
	v188 = F_memmove(m, v118, v165, v186)
	mBase = m.M
	goto L63
L65:
	;
	goto L44
L66:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v118+int32(-17)))) = base.I64_extend_i32_u(v186)
	goto L65
L67:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v118+int32(-9)))) = v186
	goto L44
L68:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v118+int32(-5)))) = uint16(v186)
	goto L44
L69:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v118+int32(-3)))) = uint8(v186)
	goto L44
L70:
	;
	v196 = v186 << (uint(int32(3)) % 32)
	*(*uint8)(unsafe.Add(mBase, uint32(v126))) = uint8(v196)
	goto L44
L71:
	;
	v235 = F_sdssplitargs(m, v118, v17+int32(60))
	mBase = m.M
	v236 = m.ExcPending
	if v236 != 0 {
		goto L1
	} else {
		goto L72
	}
L72:
	;
	v237 = *(*int32)(unsafe.Add(mBase, uint32(v17)+60))
	if v235 == int32(0) {
		goto L15
	} else {
		goto L73
	}
L73:
	;
	if v237 < int32(6) {
		goto L15
	} else {
		goto L74
	}
L74:
	;
	if v237&int32(1) != 0 {
		goto L15
	} else {
		goto L75
	}
L75:
	;
	v245 = F_valkey_calloc(m, int32(24))
	mBase = m.M
	v246 = m.ExcPending
	if v246 != 0 {
		goto L1
	} else {
		goto L76
	}
L76:
	;
	v247 = *(*int32)(unsafe.Add(mBase, uint32(v17)+60))
	if v247 < int32(1) {
		v483 = v247
		goto L77
	} else {
		goto L78
	}
L77:
	;
	v485 = int32(_a_F_aofLoadManifestFromFile_4)
	v486 = *(*int32)(unsafe.Add(mBase, uint32(v245)))
	if v486 == int32(0) {
		v545 = v485
		v549 = v483
		goto L14
	} else {
		goto L144
	}
L78:
	;
	v260 = int32(0)
	v263 = v247
	goto L79
L79:
	;
	v267 = v235 + v260<<(uint(int32(2))%32)
	v268 = *(*int32)(unsafe.Add(mBase, uint32(v267)))
	v269 = int32(_a_F_aofLoadManifestFromFile_7)
	v272 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v268))))
	if v272 != 0 {
		goto L85
	} else {
		goto L86
	}
L80:
	;
	v483 = v467
	goto L77
L81:
	;
	v469 = v260 + int32(2)
	if v469 < v467 {
		v260 = v469
		v263 = v467
		goto L79
	} else {
		goto L143
	}
L82:
	;
	v325 = int32(_a_F_aofLoadManifestFromFile_8)
	v328 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v268))))
	if v328 != 0 {
		goto L104
	} else {
		goto L105
	}
L83:
	;
	if v304-v306 != 0 {
		goto L82
	} else {
		goto L95
	}
L84:
	;
	v304 = F_tolower(m, v300)
	mBase = m.M
	v305 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v301))))
	v306 = F_tolower(m, v305)
	mBase = m.M
	goto L83
L85:
	;
	v274 = v268
	v275 = v269
	v276 = v272
	goto L88
L86:
	;
	v300 = int32(0)
	v301 = v269
	goto L84
L87:
	;
	v300 = v297 & int32(255)
	v301 = v296
	goto L84
L88:
	;
	v278 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v275))))
	if v278 == int32(0) {
		v296 = v275
		v297 = v276
		goto L87
	} else {
		goto L90
	}
L89:
	;
	v296 = v290
	v297 = int32(0)
	goto L87
L90:
	;
	v282 = v276 & int32(255)
	if v282 == v278 {
		goto L91
	} else {
		goto L92
	}
L91:
	;
	v289 = int32(1)
	v290 = v275 + v289
	v291 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v274)+1)))
	if v291 != 0 {
		v274 = v274 + v289
		v275 = v290
		v276 = v291
		goto L88
	} else {
		goto L94
	}
L92:
	;
	v284 = F_tolower(m, v282)
	mBase = m.M
	v285 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v275))))
	v286 = F_tolower(m, v285)
	mBase = m.M
	if v284 == v286 {
		goto L91
	} else {
		goto L93
	}
L93:
	;
	v288 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v274))))
	v296 = v275
	v297 = v288
	goto L87
L94:
	;
	goto L89
L95:
	;
	v310 = *(*int32)(unsafe.Add(mBase, uint32(v267+int32(4))))
	v311 = F_sdsnew(m, v310)
	mBase = m.M
	v312 = m.ExcPending
	if v312 != 0 {
		goto L1
	} else {
		goto L96
	}
L96:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v245))) = v311
	v317 = F_strchr(m, v311, int32(47))
	mBase = m.M
	if v317 != 0 {
		v322 = int32(0)
		goto L98
	} else {
		goto L99
	}
L97:
	;
	v323 = *(*int32)(unsafe.Add(mBase, uint32(v17)+60))
	if v322 != 0 {
		v467 = v323
		goto L81
	} else {
		goto L100
	}
L98:
	;
	goto L97
L99:
	;
	v319 = F_strchr(m, v311, int32(92))
	mBase = m.M
	v322 = base.B2i32(v319 == int32(0))
	goto L98
L100:
	;
	v545 = int32(_a_F_aofLoadManifestFromFile_9)
	v549 = v323
	goto L14
L101:
	;
	v422 = int32(_a_F_aofLoadManifestFromFile_10)
	v425 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v268))))
	if v425 != 0 {
		goto L132
	} else {
		goto L133
	}
L102:
	;
	if v360-v362 != 0 {
		goto L101
	} else {
		goto L114
	}
L103:
	;
	v360 = F_tolower(m, v356)
	mBase = m.M
	v361 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v357))))
	v362 = F_tolower(m, v361)
	mBase = m.M
	goto L102
L104:
	;
	v330 = v268
	v331 = v325
	v332 = v328
	goto L107
L105:
	;
	v356 = int32(0)
	v357 = v325
	goto L103
L106:
	;
	v356 = v353 & int32(255)
	v357 = v352
	goto L103
L107:
	;
	v334 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v331))))
	if v334 == int32(0) {
		v352 = v331
		v353 = v332
		goto L106
	} else {
		goto L109
	}
L108:
	;
	v352 = v346
	v353 = int32(0)
	goto L106
L109:
	;
	v338 = v332 & int32(255)
	if v338 == v334 {
		goto L110
	} else {
		goto L111
	}
L110:
	;
	v345 = int32(1)
	v346 = v331 + v345
	v347 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v330)+1)))
	if v347 != 0 {
		v330 = v330 + v345
		v331 = v346
		v332 = v347
		goto L107
	} else {
		goto L113
	}
L111:
	;
	v340 = F_tolower(m, v338)
	mBase = m.M
	v341 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v331))))
	v342 = F_tolower(m, v341)
	mBase = m.M
	if v340 == v342 {
		goto L110
	} else {
		goto L112
	}
L112:
	;
	v344 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v330))))
	v352 = v331
	v353 = v344
	goto L106
L113:
	;
	goto L108
L114:
	;
	v366 = *(*int32)(unsafe.Add(mBase, uint32(v267+int32(4))))
	v371 = v366
	goto L116
L115:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v245)+8)) = v420
	v467 = v263
	goto L81
L116:
	;
	v377 = v371 + int32(1)
	v378 = int32(*(*int8)(unsafe.Add(mBase, uint32(v371))))
	v379 = F___isspace_2(m, v378)
	mBase = m.M
	if v379 != 0 {
		v371 = v377
		goto L116
	} else {
		goto L118
	}
L117:
	;
	v380 = int32(1)
	switch v378&int32(255) + int32(-43) {
	case 0:
		v386 = v380
		goto L120
	default:
		v388 = v371
		v389 = v378
		v390 = v380
		goto L119
	case 2:
		goto L121
	}
L118:
	;
	goto L117
L119:
	;
	v393 = v389 + int32(-48)
	if base.Ui32(int32(9)) < base.Ui32(v393) {
		v417 = int64(0)
		goto L122
	} else {
		goto L123
	}
L120:
	;
	v387 = int32(*(*int8)(unsafe.Add(mBase, uint32(v377))))
	v388 = v377
	v389 = v387
	v390 = v386
	goto L119
L121:
	;
	v386 = int32(0)
	goto L120
L122:
	;
	if v390 != 0 {
		goto L127
	} else {
		goto L128
	}
L123:
	;
	v397 = v393
	v398 = v388
	v401 = int64(0)
	goto L124
L124:
	;
	v405 = v401*int64(10) - base.I64_extend_i32_u(v397)
	v406 = int32(*(*int8)(unsafe.Add(mBase, uint32(v398)+1)))
	v410 = v406 + int32(-48)
	if base.Ui32(v410) < base.Ui32(int32(10)) {
		v397 = v410
		v398 = v398 + int32(1)
		v401 = v405
		goto L124
	} else {
		goto L126
	}
L125:
	;
	v417 = v405
	goto L122
L126:
	;
	goto L125
L127:
	;
	v420 = int64(0) - v417
	goto L129
L128:
	;
	v420 = v417
	goto L129
L129:
	;
	goto L115
L130:
	;
	if v457-v459 != 0 {
		v467 = v263
		goto L81
	} else {
		goto L142
	}
L131:
	;
	v457 = F_tolower(m, v453)
	mBase = m.M
	v458 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v454))))
	v459 = F_tolower(m, v458)
	mBase = m.M
	goto L130
L132:
	;
	v427 = v268
	v428 = v422
	v429 = v425
	goto L135
L133:
	;
	v453 = int32(0)
	v454 = v422
	goto L131
L134:
	;
	v453 = v450 & int32(255)
	v454 = v449
	goto L131
L135:
	;
	v431 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v428))))
	if v431 == int32(0) {
		v449 = v428
		v450 = v429
		goto L134
	} else {
		goto L137
	}
L136:
	;
	v449 = v443
	v450 = int32(0)
	goto L134
L137:
	;
	v435 = v429 & int32(255)
	if v435 == v431 {
		goto L138
	} else {
		goto L139
	}
L138:
	;
	v442 = int32(1)
	v443 = v428 + v442
	v444 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v427)+1)))
	if v444 != 0 {
		v427 = v427 + v442
		v428 = v443
		v429 = v444
		goto L135
	} else {
		goto L141
	}
L139:
	;
	v437 = F_tolower(m, v435)
	mBase = m.M
	v438 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v428))))
	v439 = F_tolower(m, v438)
	mBase = m.M
	if v437 == v439 {
		goto L138
	} else {
		goto L140
	}
L140:
	;
	v441 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v427))))
	v449 = v428
	v450 = v441
	goto L134
L141:
	;
	goto L136
L142:
	;
	v463 = *(*int32)(unsafe.Add(mBase, uint32(v267+int32(4))))
	v464 = int32(*(*int8)(unsafe.Add(mBase, uint32(v463))))
	*(*int32)(unsafe.Add(mBase, uint32(v245)+16)) = v464
	v467 = v263
	goto L81
L143:
	;
	goto L80
L144:
	;
	v489 = *(*int64)(unsafe.Add(mBase, uint32(v245)+8))
	if v489 == int64(0) {
		v545 = v485
		v549 = v483
		goto L14
	} else {
		goto L145
	}
L145:
	;
	v492 = *(*int32)(unsafe.Add(mBase, uint32(v245)+16))
	if v492 == int32(0) {
		v545 = v485
		v549 = v483
		goto L14
	} else {
		goto L146
	}
L146:
	;
	F_sdsfreesplitres(m, v235, v483)
	mBase = m.M
	v496 = m.ExcPending
	if v496 != 0 {
		goto L1
	} else {
		goto L147
	}
L147:
	;
	v498 = *(*int32)(unsafe.Add(mBase, uint32(v245)+16))
	switch v498 + int32(-98) {
	case 0:
		goto L150
	default:
		v561 = int32(_a_F_aofLoadManifestFromFile_11)
		goto L13
	case 6:
		goto L149
	case 7:
		goto L148
	}
L148:
	;
	v516 = *(*int64)(unsafe.Add(mBase, uint32(v245)+8))
	if v63 < v516 {
		goto L157
	} else {
		goto L158
	}
L149:
	;
	v512 = F_listAddNodeTail(m, v27, v245)
	mBase = m.M
	v513 = m.ExcPending
	if v513 != 0 {
		goto L1
	} else {
		goto L155
	}
L150:
	;
	v501 = *(*int32)(unsafe.Add(mBase, uint32(v20)))
	if v501 == int32(0) {
		goto L151
	} else {
		goto L152
	}
L151:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20))) = v245
	v508 = *(*int64)(unsafe.Add(mBase, uint32(v245)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v20)+16)) = v508
	F_sdsfree(m, v118)
	mBase = m.M
	v511 = m.ExcPending
	if v511 != 0 {
		goto L1
	} else {
		goto L154
	}
L152:
	;
	F_aofInfoFree(m, v245)
	mBase = m.M
	v506 = m.ExcPending
	if v506 != 0 {
		goto L1
	} else {
		goto L153
	}
L153:
	;
	v569 = v99
	v576 = v118
	v577 = int32(_a_F_aofLoadManifestFromFile_12)
	goto L12
L154:
	;
	v57 = v99
	goto L17
L155:
	;
	F_sdsfree(m, v118)
	mBase = m.M
	v515 = m.ExcPending
	if v515 != 0 {
		goto L1
	} else {
		goto L156
	}
L156:
	;
	v57 = v99
	goto L17
L157:
	;
	v521 = F_listAddNodeTail(m, v24, v245)
	mBase = m.M
	v522 = m.ExcPending
	if v522 != 0 {
		goto L1
	} else {
		goto L160
	}
L158:
	;
	F_aofInfoFree(m, v245)
	mBase = m.M
	v520 = m.ExcPending
	if v520 != 0 {
		goto L1
	} else {
		goto L159
	}
L159:
	;
	v569 = v99
	v576 = v118
	v577 = int32(_a_F_aofLoadManifestFromFile_13)
	goto L12
L160:
	;
	v523 = *(*int64)(unsafe.Add(mBase, uint32(v245)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v20)+24)) = v523
	F_sdsfree(m, v118)
	mBase = m.M
	v526 = m.ExcPending
	if v526 != 0 {
		goto L1
	} else {
		goto L161
	}
L161:
	;
	v57 = v99
	v63 = v523
	goto L17
L162:
	;
	m.G0 = v17 + int32(1104)
	return v20
L163:
	;
	F_sdsfreesplitres(m, v235, v237)
	mBase = m.M
	v536 = m.ExcPending
	if v536 != 0 {
		goto L1
	} else {
		goto L164
	}
L164:
	;
	v569 = v99
	v576 = v118
	v577 = v115
	goto L12
L165:
	;
	v561 = v545
	goto L13
L166:
	;
	v569 = v99
	v576 = v118
	v577 = v561
	goto L12
L167:
	;
	if v576 == int32(0) {
		v619 = v594
		goto L171
	} else {
		goto L172
	}
L168:
	;
	F__serverLog(m, int32(3), int32(_a_F_aofLoadManifestFromFile_14), int32(0))
	mBase = m.M
	v591 = m.ExcPending
	if v591 != 0 {
		goto L1
	} else {
		goto L169
	}
L169:
	;
	v593 = *(*int32)(unsafe.Add(mBase, _c_F_aofLoadManifestFromFile[0]))
	v594 = v593
	goto L167
L170:
	;
	m.Env.Exit(m, int32(1))
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L171:
	;
	if int32(3) < v619 {
		goto L170
	} else {
		goto L177
	}
L172:
	;
	if int32(3) < v594 {
		goto L170
	} else {
		goto L173
	}
L173:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+48)) = v569
	F__serverLog(m, int32(3), int32(_a_F_aofLoadManifestFromFile_15), v17+int32(48))
	mBase = m.M
	v605 = m.ExcPending
	if v605 != 0 {
		goto L1
	} else {
		goto L174
	}
L174:
	;
	v607 = *(*int32)(unsafe.Add(mBase, _c_F_aofLoadManifestFromFile[0]))
	if int32(3) < v607 {
		goto L170
	} else {
		goto L175
	}
L175:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+32)) = v576
	F__serverLog(m, int32(3), int32(_a_F_aofLoadManifestFromFile_16), v17+int32(32))
	mBase = m.M
	v616 = m.ExcPending
	if v616 != 0 {
		goto L1
	} else {
		goto L176
	}
L176:
	;
	v618 = *(*int32)(unsafe.Add(mBase, _c_F_aofLoadManifestFromFile[0]))
	v619 = v618
	goto L171
L177:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+16)) = v577
	F__serverLog(m, int32(3), int32(_a_F_aofLoadManifestFromFile_17), v17+int32(16))
	mBase = m.M
	v628 = m.ExcPending
	if v628 != 0 {
		goto L1
	} else {
		goto L178
	}
L178:
	;
	goto L170
}
func F_aofManifestFreeAndUpdate(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v42 int32
	_ = v42
	if l0 == int32(0) {
		F__serverAssert(m, int32(_a_F_aofManifestFreeAndUpdate_0), int32(_a_F_aofManifestFreeAndUpdate_1), int32(413))
		mBase = m.M
		v42 = m.ExcPending
		if v42 != 0 {
			return
		} else {
			F_abort(m)
			mBase = m.M
			base.Wasm_trap_unreachable()
			for {
			}
		}
	} else {
		v8 = *(*int32)(unsafe.Add(mBase, _c_F_aofManifestFreeAndUpdate[0]))
		if v8 == int32(0) {
			*(*int32)(unsafe.Add(mBase, _c_F_aofManifestFreeAndUpdate[0])) = l0
			return
		} else {
			v11 = *(*int32)(unsafe.Add(mBase, uint32(v8)))
			if v11 == int32(0) {
				v22 = *(*int32)(unsafe.Add(mBase, uint32(v8)+4))
				if v22 == int32(0) {
					v27 = *(*int32)(unsafe.Add(mBase, uint32(v8)+8))
					if v27 == int32(0) {
						F_valkey_free(m, v8)
						mBase = m.M
						v33 = m.ExcPending
						if v33 != 0 {
							return
						} else {
							*(*int32)(unsafe.Add(mBase, _c_F_aofManifestFreeAndUpdate[0])) = l0
							return
						}
					} else {
						F_listRelease(m, v27)
						mBase = m.M
						v31 = m.ExcPending
						if v31 != 0 {
							return
						} else {
							F_valkey_free(m, v8)
							mBase = m.M
							v33 = m.ExcPending
							if v33 != 0 {
								return
							} else {
								*(*int32)(unsafe.Add(mBase, _c_F_aofManifestFreeAndUpdate[0])) = l0
								return
							}
						}
					}
				} else {
					F_listRelease(m, v22)
					mBase = m.M
					v26 = m.ExcPending
					if v26 != 0 {
						return
					} else {
						v27 = *(*int32)(unsafe.Add(mBase, uint32(v8)+8))
						if v27 == int32(0) {
							F_valkey_free(m, v8)
							mBase = m.M
							v33 = m.ExcPending
							if v33 != 0 {
								return
							} else {
								*(*int32)(unsafe.Add(mBase, _c_F_aofManifestFreeAndUpdate[0])) = l0
								return
							}
						} else {
							F_listRelease(m, v27)
							mBase = m.M
							v31 = m.ExcPending
							if v31 != 0 {
								return
							} else {
								F_valkey_free(m, v8)
								mBase = m.M
								v33 = m.ExcPending
								if v33 != 0 {
									return
								} else {
									*(*int32)(unsafe.Add(mBase, _c_F_aofManifestFreeAndUpdate[0])) = l0
									return
								}
							}
						}
					}
				}
			} else {
				v14 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
				if v14 == int32(0) {
					F_valkey_free(m, v11)
					mBase = m.M
					v20 = m.ExcPending
					if v20 != 0 {
						return
					} else {
						v22 = *(*int32)(unsafe.Add(mBase, uint32(v8)+4))
						if v22 == int32(0) {
							v27 = *(*int32)(unsafe.Add(mBase, uint32(v8)+8))
							if v27 == int32(0) {
								F_valkey_free(m, v8)
								mBase = m.M
								v33 = m.ExcPending
								if v33 != 0 {
									return
								} else {
									*(*int32)(unsafe.Add(mBase, _c_F_aofManifestFreeAndUpdate[0])) = l0
									return
								}
							} else {
								F_listRelease(m, v27)
								mBase = m.M
								v31 = m.ExcPending
								if v31 != 0 {
									return
								} else {
									F_valkey_free(m, v8)
									mBase = m.M
									v33 = m.ExcPending
									if v33 != 0 {
										return
									} else {
										*(*int32)(unsafe.Add(mBase, _c_F_aofManifestFreeAndUpdate[0])) = l0
										return
									}
								}
							}
						} else {
							F_listRelease(m, v22)
							mBase = m.M
							v26 = m.ExcPending
							if v26 != 0 {
								return
							} else {
								v27 = *(*int32)(unsafe.Add(mBase, uint32(v8)+8))
								if v27 == int32(0) {
									F_valkey_free(m, v8)
									mBase = m.M
									v33 = m.ExcPending
									if v33 != 0 {
										return
									} else {
										*(*int32)(unsafe.Add(mBase, _c_F_aofManifestFreeAndUpdate[0])) = l0
										return
									}
								} else {
									F_listRelease(m, v27)
									mBase = m.M
									v31 = m.ExcPending
									if v31 != 0 {
										return
									} else {
										F_valkey_free(m, v8)
										mBase = m.M
										v33 = m.ExcPending
										if v33 != 0 {
											return
										} else {
											*(*int32)(unsafe.Add(mBase, _c_F_aofManifestFreeAndUpdate[0])) = l0
											return
										}
									}
								}
							}
						}
					}
				} else {
					F_sdsfree(m, v14)
					mBase = m.M
					v18 = m.ExcPending
					if v18 != 0 {
						return
					} else {
						F_valkey_free(m, v11)
						mBase = m.M
						v20 = m.ExcPending
						if v20 != 0 {
							return
						} else {
							v22 = *(*int32)(unsafe.Add(mBase, uint32(v8)+4))
							if v22 == int32(0) {
								v27 = *(*int32)(unsafe.Add(mBase, uint32(v8)+8))
								if v27 == int32(0) {
									F_valkey_free(m, v8)
									mBase = m.M
									v33 = m.ExcPending
									if v33 != 0 {
										return
									} else {
										*(*int32)(unsafe.Add(mBase, _c_F_aofManifestFreeAndUpdate[0])) = l0
										return
									}
								} else {
									F_listRelease(m, v27)
									mBase = m.M
									v31 = m.ExcPending
									if v31 != 0 {
										return
									} else {
										F_valkey_free(m, v8)
										mBase = m.M
										v33 = m.ExcPending
										if v33 != 0 {
											return
										} else {
											*(*int32)(unsafe.Add(mBase, _c_F_aofManifestFreeAndUpdate[0])) = l0
											return
										}
									}
								}
							} else {
								F_listRelease(m, v22)
								mBase = m.M
								v26 = m.ExcPending
								if v26 != 0 {
									return
								} else {
									v27 = *(*int32)(unsafe.Add(mBase, uint32(v8)+8))
									if v27 == int32(0) {
										F_valkey_free(m, v8)
										mBase = m.M
										v33 = m.ExcPending
										if v33 != 0 {
											return
										} else {
											*(*int32)(unsafe.Add(mBase, _c_F_aofManifestFreeAndUpdate[0])) = l0
											return
										}
									} else {
										F_listRelease(m, v27)
										mBase = m.M
										v31 = m.ExcPending
										if v31 != 0 {
											return
										} else {
											F_valkey_free(m, v8)
											mBase = m.M
											v33 = m.ExcPending
											if v33 != 0 {
												return
											} else {
												*(*int32)(unsafe.Add(mBase, _c_F_aofManifestFreeAndUpdate[0])) = l0
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
func F_aofUpgradePrepare(m *base.Module, l0 int32) {
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
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v63 int32
	_ = v63
	var v71 int32
	_ = v71
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v84 int32
	_ = v84
	var v88 int32
	_ = v88
	var v91 int32
	_ = v91
	var v95 int32
	_ = v95
	var v97 int32
	_ = v97
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
	var v108 int64
	_ = v108
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
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
	var v132 int32
	_ = v132
	var v136 int32
	_ = v136
	var v140 int64
	_ = v140
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v151 int32
	_ = v151
	var v154 int32
	_ = v154
	var v158 int32
	_ = v158
	var v160 int32
	_ = v160
	var v164 int64
	_ = v164
	var v171 int32
	_ = v171
	var v179 int32
	_ = v179
	v6 = m.G0
	v8 = v6 - int32(48)
	m.G0 = v8
	v10 = int32(_a_F_aofUpgradePrepare_0)
	v11 = *(*int32)(unsafe.Add(mBase, _c_F_aofUpgradePrepare[0]))
	v13 = *(*int32)(unsafe.Add(mBase, _c_F_aofUpgradePrepare[1]))
	v14 = F_makePath(m, v11, v13)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		return
	} else {
		v18 = m.G0
		v19 = int32(96)
		v20 = v18 - v19
		m.G0 = v20
		v22 = F_stat(m, v14, v20)
		mBase = m.M
		v23 = *(*int32)(unsafe.Add(mBase, uint32(v20)+4))
		m.G0 = v20 + v19
		F_sdsfree(m, v14)
		mBase = m.M
		v35 = m.ExcPending
		if v35 != 0 {
			return
		} else {
			if base.B2i32(v22 == int32(0))&base.B2i32(v23&int32(61440) == int32(32768)) != 0 {
				F__serverAssert(m, int32(_a_F_aofUpgradePrepare_1), int32(_a_F_aofUpgradePrepare_2), int32(615))
				mBase = m.M
				v179 = m.ExcPending
				if v179 != 0 {
					return
				} else {
					F_abort(m)
					mBase = m.M
					base.Wasm_trap_unreachable()
					for {
					}
				}
			} else {
				v37 = *(*int32)(unsafe.Add(mBase, _c_F_aofUpgradePrepare[0]))
				v38 = int32(0)
				v41 = m.G0
				v43 = v41 - int32(96)
				m.G0 = v43
				v47 = F_mkdir(m, v37, int32(493))
				mBase = m.M
				if v47 == v38 {
					v63 = v38
				} else {
					v50 = F___errno_location(m)
					mBase = m.M
					v51 = *(*int32)(unsafe.Add(mBase, uint32(v50)))
					if v51 != int32(20) {
						v63 = int32(-1)
					} else {
						v54 = F_stat(m, v37, v43)
						mBase = m.M
						if v54 != 0 {
							*(*int32)(unsafe.Add(mBase, uint32(v50))) = int32(54)
							v63 = int32(-1)
						} else {
							v55 = *(*int32)(unsafe.Add(mBase, uint32(v43)+4))
							if v55&int32(61440) == int32(16384) {
								v63 = v38
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v50))) = int32(54)
								v63 = int32(-1)
							}
						}
					}
				}
				m.G0 = v43 + int32(96)
				if v63 != int32(-1) {
					v88 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
					if v88 == int32(0) {
						v100 = F_valkey_calloc(m, int32(24))
						mBase = m.M
						v101 = m.ExcPending
						if v101 != 0 {
							return
						} else {
							v103 = *(*int32)(unsafe.Add(mBase, _c_F_aofUpgradePrepare[1]))
							v104 = F_sdsnew(m, v103)
							mBase = m.M
							v105 = m.ExcPending
							if v105 != 0 {
								return
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v100)+16)) = int32(98)
								v108 = int64(1)
								*(*int64)(unsafe.Add(mBase, uint32(v100)+8)) = v108
								*(*int32)(unsafe.Add(mBase, uint32(v100))) = v104
								*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = int32(1)
								*(*int64)(unsafe.Add(mBase, uint32(l0)+16)) = v108
								*(*int32)(unsafe.Add(mBase, uint32(l0))) = v100
								v116 = F_getAofManifestAsString(m, l0)
								mBase = m.M
								v117 = m.ExcPending
								if v117 != 0 {
									return
								} else {
									v118 = F_writeAofManifestFile(m, v116)
									mBase = m.M
									v119 = m.ExcPending
									if v119 != 0 {
										return
									} else {
										F_sdsfree(m, v116)
										mBase = m.M
										v121 = m.ExcPending
										if v121 != 0 {
											return
										} else {
											if v118 != 0 {
												m.Env.Exit(m, int32(1))
												mBase = m.M
												base.Wasm_trap_unreachable()
												for {
												}
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = int32(0)
												v124 = int32(_a_F_aofUpgradePrepare_0)
												v125 = *(*int32)(unsafe.Add(mBase, _c_F_aofUpgradePrepare[0]))
												v127 = *(*int32)(unsafe.Add(mBase, _c_F_aofUpgradePrepare[1]))
												v128 = F_makePath(m, v125, v127)
												mBase = m.M
												v129 = m.ExcPending
												if v129 != 0 {
													return
												} else {
													v131 = *(*int32)(unsafe.Add(mBase, _c_F_aofUpgradePrepare[1]))
													v132 = F_rename(m, v131, v128)
													mBase = m.M
													if v132 != int32(-1) {
														F_sdsfree(m, v128)
														mBase = m.M
														v158 = m.ExcPending
														if v158 != 0 {
															return
														} else {
															v160 = *(*int32)(unsafe.Add(mBase, _c_F_aofUpgradePrepare[2]))
															if int32(2) < v160 {
																m.G0 = v8 + int32(48)
																return
															} else {
																v164 = *(*int64)(unsafe.Add(mBase, _c_F_aofUpgradePrepare[1]))
																*(*int64)(unsafe.Add(mBase, uint32(v8)+32)) = v164
																F__serverLog(m, int32(2), int32(_a_F_aofUpgradePrepare_3), v8+int32(32))
																mBase = m.M
																v171 = m.ExcPending
																if v171 != 0 {
																	return
																} else {
																	m.G0 = v8 + int32(48)
																	return
																}
															}
														}
													} else {
														v136 = *(*int32)(unsafe.Add(mBase, _c_F_aofUpgradePrepare[2]))
														if int32(3) < v136 {
															F_sdsfree(m, v128)
															mBase = m.M
															v154 = m.ExcPending
															if v154 != 0 {
																return
															} else {
																m.Env.Exit(m, int32(1))
																mBase = m.M
																base.Wasm_trap_unreachable()
																for {
																}
															}
														} else {
															v140 = *(*int64)(unsafe.Add(mBase, _c_F_aofUpgradePrepare[1]))
															v142 = *(*int32)(unsafe.Add(mBase, _c_F_aofUpgradePrepare[3]))
															v143 = F___strerror_l(m, v142, v142)
															mBase = m.M
															*(*int32)(unsafe.Add(mBase, uint32(v8)+24)) = v143
															*(*int64)(unsafe.Add(mBase, uint32(v8)+16)) = v140
															F__serverLog(m, int32(3), int32(_a_F_aofUpgradePrepare_4), v8+int32(16))
															mBase = m.M
															v151 = m.ExcPending
															if v151 != 0 {
																return
															} else {
																F_sdsfree(m, v128)
																mBase = m.M
																v154 = m.ExcPending
																if v154 != 0 {
																	return
																} else {
																	m.Env.Exit(m, int32(1))
																	mBase = m.M
																	base.Wasm_trap_unreachable()
																	for {
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
						v91 = *(*int32)(unsafe.Add(mBase, uint32(v88)))
						if v91 == int32(0) {
							F_valkey_free(m, v88)
							mBase = m.M
							v97 = m.ExcPending
							if v97 != 0 {
								return
							} else {
								v100 = F_valkey_calloc(m, int32(24))
								mBase = m.M
								v101 = m.ExcPending
								if v101 != 0 {
									return
								} else {
									v103 = *(*int32)(unsafe.Add(mBase, _c_F_aofUpgradePrepare[1]))
									v104 = F_sdsnew(m, v103)
									mBase = m.M
									v105 = m.ExcPending
									if v105 != 0 {
										return
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v100)+16)) = int32(98)
										v108 = int64(1)
										*(*int64)(unsafe.Add(mBase, uint32(v100)+8)) = v108
										*(*int32)(unsafe.Add(mBase, uint32(v100))) = v104
										*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = int32(1)
										*(*int64)(unsafe.Add(mBase, uint32(l0)+16)) = v108
										*(*int32)(unsafe.Add(mBase, uint32(l0))) = v100
										v116 = F_getAofManifestAsString(m, l0)
										mBase = m.M
										v117 = m.ExcPending
										if v117 != 0 {
											return
										} else {
											v118 = F_writeAofManifestFile(m, v116)
											mBase = m.M
											v119 = m.ExcPending
											if v119 != 0 {
												return
											} else {
												F_sdsfree(m, v116)
												mBase = m.M
												v121 = m.ExcPending
												if v121 != 0 {
													return
												} else {
													if v118 != 0 {
														m.Env.Exit(m, int32(1))
														mBase = m.M
														base.Wasm_trap_unreachable()
														for {
														}
													} else {
														*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = int32(0)
														v124 = int32(_a_F_aofUpgradePrepare_0)
														v125 = *(*int32)(unsafe.Add(mBase, _c_F_aofUpgradePrepare[0]))
														v127 = *(*int32)(unsafe.Add(mBase, _c_F_aofUpgradePrepare[1]))
														v128 = F_makePath(m, v125, v127)
														mBase = m.M
														v129 = m.ExcPending
														if v129 != 0 {
															return
														} else {
															v131 = *(*int32)(unsafe.Add(mBase, _c_F_aofUpgradePrepare[1]))
															v132 = F_rename(m, v131, v128)
															mBase = m.M
															if v132 != int32(-1) {
																F_sdsfree(m, v128)
																mBase = m.M
																v158 = m.ExcPending
																if v158 != 0 {
																	return
																} else {
																	v160 = *(*int32)(unsafe.Add(mBase, _c_F_aofUpgradePrepare[2]))
																	if int32(2) < v160 {
																		m.G0 = v8 + int32(48)
																		return
																	} else {
																		v164 = *(*int64)(unsafe.Add(mBase, _c_F_aofUpgradePrepare[1]))
																		*(*int64)(unsafe.Add(mBase, uint32(v8)+32)) = v164
																		F__serverLog(m, int32(2), int32(_a_F_aofUpgradePrepare_3), v8+int32(32))
																		mBase = m.M
																		v171 = m.ExcPending
																		if v171 != 0 {
																			return
																		} else {
																			m.G0 = v8 + int32(48)
																			return
																		}
																	}
																}
															} else {
																v136 = *(*int32)(unsafe.Add(mBase, _c_F_aofUpgradePrepare[2]))
																if int32(3) < v136 {
																	F_sdsfree(m, v128)
																	mBase = m.M
																	v154 = m.ExcPending
																	if v154 != 0 {
																		return
																	} else {
																		m.Env.Exit(m, int32(1))
																		mBase = m.M
																		base.Wasm_trap_unreachable()
																		for {
																		}
																	}
																} else {
																	v140 = *(*int64)(unsafe.Add(mBase, _c_F_aofUpgradePrepare[1]))
																	v142 = *(*int32)(unsafe.Add(mBase, _c_F_aofUpgradePrepare[3]))
																	v143 = F___strerror_l(m, v142, v142)
																	mBase = m.M
																	*(*int32)(unsafe.Add(mBase, uint32(v8)+24)) = v143
																	*(*int64)(unsafe.Add(mBase, uint32(v8)+16)) = v140
																	F__serverLog(m, int32(3), int32(_a_F_aofUpgradePrepare_4), v8+int32(16))
																	mBase = m.M
																	v151 = m.ExcPending
																	if v151 != 0 {
																		return
																	} else {
																		F_sdsfree(m, v128)
																		mBase = m.M
																		v154 = m.ExcPending
																		if v154 != 0 {
																			return
																		} else {
																			m.Env.Exit(m, int32(1))
																			mBase = m.M
																			base.Wasm_trap_unreachable()
																			for {
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
							F_sdsfree(m, v91)
							mBase = m.M
							v95 = m.ExcPending
							if v95 != 0 {
								return
							} else {
								F_valkey_free(m, v88)
								mBase = m.M
								v97 = m.ExcPending
								if v97 != 0 {
									return
								} else {
									v100 = F_valkey_calloc(m, int32(24))
									mBase = m.M
									v101 = m.ExcPending
									if v101 != 0 {
										return
									} else {
										v103 = *(*int32)(unsafe.Add(mBase, _c_F_aofUpgradePrepare[1]))
										v104 = F_sdsnew(m, v103)
										mBase = m.M
										v105 = m.ExcPending
										if v105 != 0 {
											return
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(v100)+16)) = int32(98)
											v108 = int64(1)
											*(*int64)(unsafe.Add(mBase, uint32(v100)+8)) = v108
											*(*int32)(unsafe.Add(mBase, uint32(v100))) = v104
											*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = int32(1)
											*(*int64)(unsafe.Add(mBase, uint32(l0)+16)) = v108
											*(*int32)(unsafe.Add(mBase, uint32(l0))) = v100
											v116 = F_getAofManifestAsString(m, l0)
											mBase = m.M
											v117 = m.ExcPending
											if v117 != 0 {
												return
											} else {
												v118 = F_writeAofManifestFile(m, v116)
												mBase = m.M
												v119 = m.ExcPending
												if v119 != 0 {
													return
												} else {
													F_sdsfree(m, v116)
													mBase = m.M
													v121 = m.ExcPending
													if v121 != 0 {
														return
													} else {
														if v118 != 0 {
															m.Env.Exit(m, int32(1))
															mBase = m.M
															base.Wasm_trap_unreachable()
															for {
															}
														} else {
															*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = int32(0)
															v124 = int32(_a_F_aofUpgradePrepare_0)
															v125 = *(*int32)(unsafe.Add(mBase, _c_F_aofUpgradePrepare[0]))
															v127 = *(*int32)(unsafe.Add(mBase, _c_F_aofUpgradePrepare[1]))
															v128 = F_makePath(m, v125, v127)
															mBase = m.M
															v129 = m.ExcPending
															if v129 != 0 {
																return
															} else {
																v131 = *(*int32)(unsafe.Add(mBase, _c_F_aofUpgradePrepare[1]))
																v132 = F_rename(m, v131, v128)
																mBase = m.M
																if v132 != int32(-1) {
																	F_sdsfree(m, v128)
																	mBase = m.M
																	v158 = m.ExcPending
																	if v158 != 0 {
																		return
																	} else {
																		v160 = *(*int32)(unsafe.Add(mBase, _c_F_aofUpgradePrepare[2]))
																		if int32(2) < v160 {
																			m.G0 = v8 + int32(48)
																			return
																		} else {
																			v164 = *(*int64)(unsafe.Add(mBase, _c_F_aofUpgradePrepare[1]))
																			*(*int64)(unsafe.Add(mBase, uint32(v8)+32)) = v164
																			F__serverLog(m, int32(2), int32(_a_F_aofUpgradePrepare_3), v8+int32(32))
																			mBase = m.M
																			v171 = m.ExcPending
																			if v171 != 0 {
																				return
																			} else {
																				m.G0 = v8 + int32(48)
																				return
																			}
																		}
																	}
																} else {
																	v136 = *(*int32)(unsafe.Add(mBase, _c_F_aofUpgradePrepare[2]))
																	if int32(3) < v136 {
																		F_sdsfree(m, v128)
																		mBase = m.M
																		v154 = m.ExcPending
																		if v154 != 0 {
																			return
																		} else {
																			m.Env.Exit(m, int32(1))
																			mBase = m.M
																			base.Wasm_trap_unreachable()
																			for {
																			}
																		}
																	} else {
																		v140 = *(*int64)(unsafe.Add(mBase, _c_F_aofUpgradePrepare[1]))
																		v142 = *(*int32)(unsafe.Add(mBase, _c_F_aofUpgradePrepare[3]))
																		v143 = F___strerror_l(m, v142, v142)
																		mBase = m.M
																		*(*int32)(unsafe.Add(mBase, uint32(v8)+24)) = v143
																		*(*int64)(unsafe.Add(mBase, uint32(v8)+16)) = v140
																		F__serverLog(m, int32(3), int32(_a_F_aofUpgradePrepare_4), v8+int32(16))
																		mBase = m.M
																		v151 = m.ExcPending
																		if v151 != 0 {
																			return
																		} else {
																			F_sdsfree(m, v128)
																			mBase = m.M
																			v154 = m.ExcPending
																			if v154 != 0 {
																				return
																			} else {
																				m.Env.Exit(m, int32(1))
																				mBase = m.M
																				base.Wasm_trap_unreachable()
																				for {
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
				} else {
					v71 = *(*int32)(unsafe.Add(mBase, _c_F_aofUpgradePrepare[2]))
					if int32(3) < v71 {
						m.Env.Exit(m, int32(1))
						mBase = m.M
						base.Wasm_trap_unreachable()
						for {
						}
					} else {
						v75 = *(*int32)(unsafe.Add(mBase, _c_F_aofUpgradePrepare[0]))
						v77 = *(*int32)(unsafe.Add(mBase, _c_F_aofUpgradePrepare[3]))
						v78 = F___strerror_l(m, v77, v77)
						mBase = m.M
						*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = v78
						*(*int32)(unsafe.Add(mBase, uint32(v8))) = v75
						F__serverLog(m, int32(3), int32(_a_F_aofUpgradePrepare_5), v8)
						mBase = m.M
						v84 = m.ExcPending
						if v84 != 0 {
							return
						} else {
							m.Env.Exit(m, int32(1))
							mBase = m.M
							base.Wasm_trap_unreachable()
							for {
							}
						}
					}
				}
			}
		}
	}
}
func F_genAofTimestampAnnotationIfNeeded(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v11 int64
	_ = v11
	var v13 int64
	_ = v13
	var v16 int64
	_ = v16
	var v18 int64
	_ = v18
	var v20 int64
	_ = v20
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v28 int64
	_ = v28
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v63 int32
	_ = v63
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	if l0 != 0 {
		v18 = F___time(m, int32(0))
		mBase = m.M
		v20 = v18
		*(*int64)(unsafe.Add(mBase, _c_F_genAofTimestampAnnotationIfNeeded[0])) = v20
		v23 = F_sdsempty(m)
		mBase = m.M
		v26 = m.ExcPending
		if v26 != 0 {
			return int32(0)
		} else {
			v28 = *(*int64)(unsafe.Add(mBase, _c_F_genAofTimestampAnnotationIfNeeded[0]))
			*(*int64)(unsafe.Add(mBase, uint32(v7))) = v28
			v31 = F_sdscatfmt(m, v23, int32(_a_F_genAofTimestampAnnotationIfNeeded_0), v7)
			mBase = m.M
			v32 = m.ExcPending
			if v32 != 0 {
				return int32(0)
			} else {
				v35 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v31+int32(-1)))))
				switch v35&int32(7) + int32(-2) {
				case 0:
					v42 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v31+int32(-5)))))
					v49 = v42
					if base.Ui32(int32(1025)) <= base.Ui32(v49) {
						F__serverAssert(m, int32(_a_F_genAofTimestampAnnotationIfNeeded_1), int32(_a_F_genAofTimestampAnnotationIfNeeded_2), int32(1433))
						mBase = m.M
						v63 = m.ExcPending
						if v63 != 0 {
							return int32(0)
						} else {
							F_abort(m)
							mBase = m.M
							base.Wasm_trap_unreachable()
							for {
							}
						}
					} else {
						v52 = v31
						m.G0 = v7 + int32(16)
						return v52
					}
				case 1:
					v45 = *(*int32)(unsafe.Add(mBase, uint32(v31+int32(-9))))
					v49 = v45
					if base.Ui32(int32(1025)) <= base.Ui32(v49) {
						F__serverAssert(m, int32(_a_F_genAofTimestampAnnotationIfNeeded_1), int32(_a_F_genAofTimestampAnnotationIfNeeded_2), int32(1433))
						mBase = m.M
						v63 = m.ExcPending
						if v63 != 0 {
							return int32(0)
						} else {
							F_abort(m)
							mBase = m.M
							base.Wasm_trap_unreachable()
							for {
							}
						}
					} else {
						v52 = v31
						m.G0 = v7 + int32(16)
						return v52
					}
				case 2:
					v48 = *(*int32)(unsafe.Add(mBase, uint32(v31+int32(-17))))
					v49 = v48
					if base.Ui32(int32(1025)) <= base.Ui32(v49) {
						F__serverAssert(m, int32(_a_F_genAofTimestampAnnotationIfNeeded_1), int32(_a_F_genAofTimestampAnnotationIfNeeded_2), int32(1433))
						mBase = m.M
						v63 = m.ExcPending
						if v63 != 0 {
							return int32(0)
						} else {
							F_abort(m)
							mBase = m.M
							base.Wasm_trap_unreachable()
							for {
							}
						}
					} else {
						v52 = v31
						m.G0 = v7 + int32(16)
						return v52
					}
				default:
					v52 = v31
					m.G0 = v7 + int32(16)
					return v52
				}
			}
		}
	} else {
		v10 = int32(_a_F_genAofTimestampAnnotationIfNeeded_3)
		v11 = *(*int64)(unsafe.Add(mBase, _c_F_genAofTimestampAnnotationIfNeeded[0]))
		v13 = *(*int64)(unsafe.Add(mBase, _c_F_genAofTimestampAnnotationIfNeeded[1]))
		if v13 <= v11 {
			v52 = int32(0)
			m.G0 = v7 + int32(16)
			return v52
		} else {
			v16 = *(*int64)(unsafe.Add(mBase, _c_F_genAofTimestampAnnotationIfNeeded[1]))
			v20 = v16
			*(*int64)(unsafe.Add(mBase, _c_F_genAofTimestampAnnotationIfNeeded[0])) = v20
			v23 = F_sdsempty(m)
			mBase = m.M
			v26 = m.ExcPending
			if v26 != 0 {
				return int32(0)
			} else {
				v28 = *(*int64)(unsafe.Add(mBase, _c_F_genAofTimestampAnnotationIfNeeded[0]))
				*(*int64)(unsafe.Add(mBase, uint32(v7))) = v28
				v31 = F_sdscatfmt(m, v23, int32(_a_F_genAofTimestampAnnotationIfNeeded_0), v7)
				mBase = m.M
				v32 = m.ExcPending
				if v32 != 0 {
					return int32(0)
				} else {
					v35 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v31+int32(-1)))))
					switch v35&int32(7) + int32(-2) {
					case 0:
						v42 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v31+int32(-5)))))
						v49 = v42
						if base.Ui32(int32(1025)) <= base.Ui32(v49) {
							F__serverAssert(m, int32(_a_F_genAofTimestampAnnotationIfNeeded_1), int32(_a_F_genAofTimestampAnnotationIfNeeded_2), int32(1433))
							mBase = m.M
							v63 = m.ExcPending
							if v63 != 0 {
								return int32(0)
							} else {
								F_abort(m)
								mBase = m.M
								base.Wasm_trap_unreachable()
								for {
								}
							}
						} else {
							v52 = v31
							m.G0 = v7 + int32(16)
							return v52
						}
					case 1:
						v45 = *(*int32)(unsafe.Add(mBase, uint32(v31+int32(-9))))
						v49 = v45
						if base.Ui32(int32(1025)) <= base.Ui32(v49) {
							F__serverAssert(m, int32(_a_F_genAofTimestampAnnotationIfNeeded_1), int32(_a_F_genAofTimestampAnnotationIfNeeded_2), int32(1433))
							mBase = m.M
							v63 = m.ExcPending
							if v63 != 0 {
								return int32(0)
							} else {
								F_abort(m)
								mBase = m.M
								base.Wasm_trap_unreachable()
								for {
								}
							}
						} else {
							v52 = v31
							m.G0 = v7 + int32(16)
							return v52
						}
					case 2:
						v48 = *(*int32)(unsafe.Add(mBase, uint32(v31+int32(-17))))
						v49 = v48
						if base.Ui32(int32(1025)) <= base.Ui32(v49) {
							F__serverAssert(m, int32(_a_F_genAofTimestampAnnotationIfNeeded_1), int32(_a_F_genAofTimestampAnnotationIfNeeded_2), int32(1433))
							mBase = m.M
							v63 = m.ExcPending
							if v63 != 0 {
								return int32(0)
							} else {
								F_abort(m)
								mBase = m.M
								base.Wasm_trap_unreachable()
								for {
								}
							}
						} else {
							v52 = v31
							m.G0 = v7 + int32(16)
							return v52
						}
					default:
						v52 = v31
						m.G0 = v7 + int32(16)
						return v52
					}
				}
			}
		}
	}
}
