package p2

import (
	base "github.com/shibukawa/valkeymem/internal/aot/vkaot/base"
	"unsafe"
)

func F_pubsubPublishMessageInternal(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
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
	var v50 int32
	_ = v50
	var v59 int32
	_ = v59
	var v63 int32
	_ = v63
	var v67 int32
	_ = v67
	var v71 int32
	_ = v71
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v94 int32
	_ = v94
	var v103 int32
	_ = v103
	var v110 int32
	_ = v110
	var v113 int32
	_ = v113
	var v117 int32
	_ = v117
	var v120 int32
	_ = v120
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
	var v132 int32
	_ = v132
	var v135 int32
	_ = v135
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v161 int32
	_ = v161
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
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v191 int32
	_ = v191
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v203 int32
	_ = v203
	var v214 int32
	_ = v214
	var v220 int32
	_ = v220
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v247 int32
	_ = v247
	var v250 int32
	_ = v250
	var v251 int32
	_ = v251
	var v255 int32
	_ = v255
	var v259 int32
	_ = v259
	var v260 int32
	_ = v260
	var v261 int32
	_ = v261
	var v264 int32
	_ = v264
	var v266 int32
	_ = v266
	var v268 int64
	_ = v268
	var v269 int64
	_ = v269
	var v270 int64
	_ = v270
	var v271 int64
	_ = v271
	var v272 int64
	_ = v272
	var v273 int64
	_ = v273
	var v274 int64
	_ = v274
	var v276 int64
	_ = v276
	var v278 int64
	_ = v278
	var v280 int64
	_ = v280
	var v282 int64
	_ = v282
	var v284 int64
	_ = v284
	var v286 int32
	_ = v286
	var v287 int32
	_ = v287
	var v288 int32
	_ = v288
	var v291 int32
	_ = v291
	var v294 int32
	_ = v294
	var v295 int32
	_ = v295
	var v296 int32
	_ = v296
	var v298 int32
	_ = v298
	var v301 int32
	_ = v301
	var v302 int32
	_ = v302
	var v306 int32
	_ = v306
	var v310 int32
	_ = v310
	var v312 int32
	_ = v312
	var v319 int32
	_ = v319
	var v320 int32
	_ = v320
	var v321 int32
	_ = v321
	var v326 int32
	_ = v326
	var v331 int32
	_ = v331
	var v335 int32
	_ = v335
	var v339 int32
	_ = v339
	var v347 int32
	_ = v347
	var v350 int32
	_ = v350
	var v358 int32
	_ = v358
	var v359 int32
	_ = v359
	var v360 int32
	_ = v360
	var v362 int32
	_ = v362
	var v365 int32
	_ = v365
	var v372 int32
	_ = v372
	var v375 int32
	_ = v375
	var v378 int32
	_ = v378
	var v381 int32
	_ = v381
	var v382 int32
	_ = v382
	var v383 int32
	_ = v383
	var v385 int32
	_ = v385
	var v388 int32
	_ = v388
	var v395 int32
	_ = v395
	var v398 int32
	_ = v398
	var v401 int32
	_ = v401
	var v404 int32
	_ = v404
	var v405 int32
	_ = v405
	var v406 int32
	_ = v406
	var v408 int32
	_ = v408
	var v409 int32
	_ = v409
	var v410 int32
	_ = v410
	var v417 int32
	_ = v417
	var v424 int32
	_ = v424
	var v425 int32
	_ = v425
	var v447 int32
	_ = v447
	var v448 int32
	_ = v448
	var v456 int32
	_ = v456
	var v464 int32
	_ = v464
	var v466 int32
	_ = v466
	var v467 int32
	_ = v467
	var v468 int32
	_ = v468
	var v469 int32
	_ = v469
	var v471 int32
	_ = v471
	var v476 int32
	_ = v476
	var v477 int32
	_ = v477
	var v483 int32
	_ = v483
	var v494 int32
	_ = v494
	var v500 int32
	_ = v500
	var v515 int32
	_ = v515
	var v516 int32
	_ = v516
	var v519 int32
	_ = v519
	var v522 int32
	_ = v522
	var v523 int32
	_ = v523
	var v527 int32
	_ = v527
	var v531 int32
	_ = v531
	var v532 int32
	_ = v532
	var v533 int32
	_ = v533
	var v536 int32
	_ = v536
	var v538 int32
	_ = v538
	var v540 int64
	_ = v540
	var v541 int64
	_ = v541
	var v542 int64
	_ = v542
	var v543 int64
	_ = v543
	var v544 int64
	_ = v544
	var v545 int64
	_ = v545
	var v546 int64
	_ = v546
	var v548 int64
	_ = v548
	var v550 int64
	_ = v550
	var v552 int64
	_ = v552
	var v554 int64
	_ = v554
	var v556 int64
	_ = v556
	var v558 int32
	_ = v558
	var v559 int32
	_ = v559
	var v560 int32
	_ = v560
	var v563 int32
	_ = v563
	var v566 int32
	_ = v566
	var v567 int32
	_ = v567
	var v568 int32
	_ = v568
	var v570 int32
	_ = v570
	var v573 int32
	_ = v573
	var v574 int32
	_ = v574
	var v578 int32
	_ = v578
	var v582 int32
	_ = v582
	var v584 int32
	_ = v584
	var v591 int32
	_ = v591
	var v592 int32
	_ = v592
	var v593 int32
	_ = v593
	var v598 int32
	_ = v598
	var v603 int32
	_ = v603
	var v607 int32
	_ = v607
	var v611 int32
	_ = v611
	var v620 int32
	_ = v620
	var v629 int32
	_ = v629
	var v631 int32
	_ = v631
	var v637 int32
	_ = v637
	v14 = m.G0
	v16 = v14 - int32(80)
	m.G0 = v16
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	v20 = int32(-1)
	v22 = *(*int32)(unsafe.Add(mBase, _consts[139]))
	if v22 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	v125 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	v126 = *(*int32)(unsafe.Add(mBase, uint32(v125)))
	v129 = F_kvstoreHashtableFind(m, v126, v122, l0, v16+int32(76))
	mBase = m.M
	v132 = m.ExcPending
	if v132 != 0 {
		goto L36
	} else {
		goto L37
	}
L2:
	;
	v24 = int32(0)
	if v18 == v24 {
		v121 = v20
		v122 = v24
		goto L1
	} else {
		goto L4
	}
L3:
	;
	v121 = v20
	v122 = int32(0)
	goto L1
L4:
	;
	v27 = F_objectGetVal(m, l0)
	mBase = m.M
	v29 = F_objectGetVal(m, l0)
	mBase = m.M
	v32 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v29+int32(-1)))))
	switch v32 & int32(7) {
	case 0:
		goto L10
	case 1:
		goto L9
	case 2:
		goto L8
	case 3:
		goto L7
	case 4:
		goto L6
	default:
		v49 = int32(0)
		goto L5
	}
L5:
	;
	v50 = int32(0)
	if v49 < int32(1) {
		v71 = v50
		goto L15
	} else {
		goto L16
	}
L6:
	;
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v29+int32(-17))))
	v49 = v48
	goto L5
L7:
	;
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v29+int32(-9))))
	v49 = v45
	goto L5
L8:
	;
	v42 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v29+int32(-5)))))
	v49 = v42
	goto L5
L9:
	;
	v39 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v29+int32(-3)))))
	v49 = v39
	goto L5
L10:
	;
	v49 = int32(base.Ui32(v32) >> (uint(int32(3)) % 32))
	goto L5
L11:
	;
	if v117 == int32(-1) {
		goto L32
	} else {
		goto L33
	}
L12:
	;
	v117 = v113 & int32(16383)
	goto L11
L13:
	;
	v82 = v71 + int32(1)
	if v49 <= v82 {
		goto L23
	} else {
		goto L24
	}
L14:
	;
	v80 = F_crc16(m, v27, v49)
	mBase = m.M
	v113 = v80
	goto L12
L15:
	;
	if v71 != v49 {
		goto L13
	} else {
		goto L21
	}
L16:
	;
	v59 = v50
	goto L17
L17:
	;
	v63 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v27+v59))))
	if v63 == int32(123) {
		v71 = v59
		goto L15
	} else {
		goto L19
	}
L19:
	;
	v67 = v59 + int32(1)
	if v67 != v49 {
		v59 = v67
		goto L17
	} else {
		goto L20
	}
L20:
	;
	goto L14
L21:
	;
	goto L14
L22:
	;
	v110 = F_crc16(m, v27+v71+int32(1), v88+(v71^int32(-1)))
	mBase = m.M
	v113 = v110
	goto L12
L23:
	;
	v103 = F_crc16(m, v27, v49)
	mBase = m.M
	v113 = v103
	goto L12
L24:
	;
	v88 = v82
	goto L26
L25:
	;
	if v88 == v49 {
		goto L23
	} else {
		goto L30
	}
L26:
	;
	v90 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v27+v88))))
	if v90 == int32(125) {
		goto L25
	} else {
		goto L28
	}
L28:
	;
	v94 = v88 + int32(1)
	if v94 != v49 {
		v88 = v94
		goto L26
	} else {
		goto L29
	}
L29:
	;
	goto L23
L30:
	;
	if v88 != v82 {
		goto L22
	} else {
		goto L31
	}
L31:
	;
	goto L23
L32:
	;
	v120 = v50
	goto L34
L33:
	;
	v120 = v117
	goto L34
L34:
	;
	v121 = v117
	v122 = v120
	goto L1
L35:
	;
	if v18 != 0 {
		v637 = v220
		goto L54
	} else {
		goto L55
	}
L36:
	;
	return int32(0)
L37:
	;
	if v129 == int32(0) {
		v220 = int32(0)
		goto L35
	} else {
		goto L38
	}
L38:
	;
	v135 = int32(0)
	v137 = v16 + int32(16)
	v138 = *(*int32)(unsafe.Add(mBase, uint32(v16)+76))
	*(*uint8)(unsafe.Add(mBase, uint32(v137)+14)) = uint8(v135)
	*(*int32)(unsafe.Add(mBase, uint32(v137))) = v138
	*(*int32)(unsafe.Add(mBase, uint32(v137)+24)) = v135
	*(*uint8)(unsafe.Add(mBase, uint32(v137)+15)) = uint8(v135)
	*(*int32)(unsafe.Add(mBase, uint32(v137)+8)) = int32(-1)
	if v138 == v135 {
		goto L40
	} else {
		goto L41
	}
L39:
	;
	v161 = F_hashtableNext(m, v16+int32(16), v16+int32(12))
	mBase = m.M
	v162 = m.ExcPending
	if v162 != 0 {
		goto L36
	} else {
		goto L44
	}
L40:
	;
	goto L39
L41:
	;
	goto L40
L43:
	;
	F_hashtableCleanupIterator(m, v16+int32(16))
	mBase = m.M
	v214 = m.ExcPending
	if v214 != 0 {
		goto L36
	} else {
		goto L53
	}
L44:
	;
	if v161 == int32(0) {
		v203 = v135
		goto L43
	} else {
		goto L45
	}
L45:
	;
	v165 = *(*int32)(unsafe.Add(mBase, uint32(l2)+24))
	v172 = int32(0)
	goto L46
L46:
	;
	v180 = *(*int32)(unsafe.Add(mBase, uint32(v16)+12))
	v181 = *(*int32)(unsafe.Add(mBase, uint32(v165)))
	F_addReplyPubsubMessage(m, v180, l0, l1, v181)
	mBase = m.M
	v183 = m.ExcPending
	if v183 != 0 {
		goto L36
	} else {
		goto L48
	}
L47:
	;
	v203 = v191
	goto L43
L48:
	;
	v184 = *(*int32)(unsafe.Add(mBase, uint32(v16)+12))
	F_clusterSlotStatsAddNetworkBytesOutForShardedPubSubInternalPropagation(m, v184, v121)
	mBase = m.M
	v186 = m.ExcPending
	if v186 != 0 {
		goto L36
	} else {
		goto L49
	}
L49:
	;
	v187 = *(*int32)(unsafe.Add(mBase, uint32(v16)+12))
	v188 = F_updateClientMemUsageAndBucket(m, v187)
	mBase = m.M
	v189 = m.ExcPending
	if v189 != 0 {
		goto L36
	} else {
		goto L50
	}
L50:
	;
	v191 = v172 + int32(1)
	v196 = F_hashtableNext(m, v16+int32(16), v16+int32(12))
	mBase = m.M
	v197 = m.ExcPending
	if v197 != 0 {
		goto L36
	} else {
		goto L51
	}
L51:
	;
	if v196 != 0 {
		v172 = v191
		goto L46
	} else {
		goto L52
	}
L52:
	;
	goto L47
L53:
	;
	v220 = v203
	goto L35
L54:
	;
	m.G0 = v16 + int32(80)
	return v637
L55:
	;
	v229 = *(*int32)(unsafe.Add(mBase, _consts[18]))
	v230 = F_dictGetIterator(m, v229)
	mBase = m.M
	v231 = m.ExcPending
	if v231 != 0 {
		goto L36
	} else {
		goto L56
	}
L56:
	;
	if v230 == int32(0) {
		v637 = v220
		goto L54
	} else {
		goto L57
	}
L57:
	;
	v234 = F_getDecodedObject(m, l0)
	mBase = m.M
	v235 = m.ExcPending
	if v235 != 0 {
		goto L36
	} else {
		goto L58
	}
L58:
	;
	v243 = v230 + int32(20)
	v244 = *(*int32)(unsafe.Add(mBase, uint32(v230)+16))
	if v244 != 0 {
		goto L62
	} else {
		goto L63
	}
L59:
	;
	F_decrRefCount(m, v234)
	mBase = m.M
	v629 = m.ExcPending
	if v629 != 0 {
		goto L36
	} else {
		goto L147
	}
L60:
	;
	if v339 == int32(0) {
		v620 = v220
		goto L59
	} else {
		goto L86
	}
L61:
	;
	v250 = v243
	v251 = v247
	goto L64
L62:
	;
	v247 = int32(1)
	goto L61
L63:
	;
	v247 = int32(0)
	goto L61
L64:
	;
	switch v251 {
	case 0:
		goto L69
	default:
		goto L68
	}
L66:
	;
	v251 = int32(0)
	goto L64
L67:
	;
	goto L60
L68:
	;
	v331 = *(*int32)(unsafe.Add(mBase, uint32(v250)))
	*(*int32)(unsafe.Add(mBase, uint32(v230)+16)) = v331
	if v331 == int32(0) {
		goto L66
	} else {
		goto L85
	}
L69:
	;
	v255 = *(*int32)(unsafe.Add(mBase, uint32(v230)+4))
	if v255 != int32(-1) {
		v294 = v255
		goto L70
	} else {
		goto L71
	}
L70:
	;
	v295 = int32(1)
	v296 = v294 + v295
	*(*int32)(unsafe.Add(mBase, uint32(v230)+4)) = v296
	v298 = int32(0)
	v301 = *(*int32)(unsafe.Add(mBase, uint32(v230)))
	v302 = *(*int32)(unsafe.Add(mBase, uint32(v230)+8))
	v306 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v301+v302+int32(26)))))
	if v306 == int32(255) {
		goto L79
	} else {
		goto L80
	}
L71:
	;
	v259 = *(*int32)(unsafe.Add(mBase, uint32(v230)+8))
	if v259 != 0 {
		v294 = int32(-1)
		goto L70
	} else {
		goto L72
	}
L72:
	;
	v260 = *(*int32)(unsafe.Add(mBase, uint32(v230)))
	v261 = *(*int32)(unsafe.Add(mBase, uint32(v230)+12))
	if v261 == int32(0) {
		goto L74
	} else {
		goto L75
	}
L73:
	;
	v288 = *(*int32)(unsafe.Add(mBase, uint32(v287)+20))
	if v288 != int32(-1) {
		goto L76
	} else {
		goto L77
	}
L74:
	;
	v268 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v260)+16)))
	v269 = int64(*(*int8)(unsafe.Add(mBase, uint32(v260)+27)))
	v270 = int64(*(*int32)(unsafe.Add(mBase, uint32(v260)+8)))
	v271 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v260)+12)))
	v272 = int64(*(*int8)(unsafe.Add(mBase, uint32(v260)+26)))
	v273 = int64(*(*int32)(unsafe.Add(mBase, uint32(v260)+4)))
	v274 = F_wangHash64(m, v273)
	mBase = m.M
	v276 = F_wangHash64(m, v272+v274)
	mBase = m.M
	v278 = F_wangHash64(m, v271+v276)
	mBase = m.M
	v280 = F_wangHash64(m, v270+v278)
	mBase = m.M
	v282 = F_wangHash64(m, v269+v280)
	mBase = m.M
	v284 = F_wangHash64(m, v268+v282)
	mBase = m.M
	*(*int64)(unsafe.Add(mBase, uint32(v230)+24)) = v284
	v286 = *(*int32)(unsafe.Add(mBase, uint32(v230)))
	v287 = v286
	goto L73
L75:
	;
	v264 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v260)+24)))
	v266 = v264 + int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v260)+24)) = uint16(v266)
	v287 = v260
	goto L73
L76:
	;
	v294 = v288 + int32(-1)
	goto L70
L77:
	;
	v291 = *(*int32)(unsafe.Add(mBase, uint32(v230)+4))
	v294 = v291
	goto L70
L78:
	;
	v321 = int32(2)
	v326 = *(*int32)(unsafe.Add(mBase, uint32(v301+v319<<(uint(v321)%32)+int32(4))))
	v250 = v326 + v320<<(uint(v321)%32)
	v251 = int32(1)
	goto L64
L79:
	;
	v310 = v298
	goto L81
L80:
	;
	v310 = v295 << (uint(v306) % 32)
	goto L81
L81:
	;
	if v296 < v310 {
		v319 = v302
		v320 = v296
		goto L78
	} else {
		goto L82
	}
L82:
	;
	if v302 != 0 {
		v339 = v298
		goto L67
	} else {
		goto L83
	}
L83:
	;
	v312 = *(*int32)(unsafe.Add(mBase, uint32(v301)+20))
	if v312 == int32(-1) {
		v339 = v298
		goto L67
	} else {
		goto L84
	}
L84:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v230)+4)) = int64(4294967296)
	v319 = int32(1)
	v320 = int32(0)
	goto L78
L85:
	;
	v335 = *(*int32)(unsafe.Add(mBase, uint32(v331)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v243))) = v335
	v339 = v331
	goto L67
L86:
	;
	v347 = v339
	v350 = v220
	goto L87
L87:
	;
	v358 = *(*int32)(unsafe.Add(mBase, uint32(v347)))
	goto L89
L88:
	;
	v620 = v500
	goto L59
L89:
	;
	v359 = *(*int32)(unsafe.Add(mBase, uint32(v347)+8))
	goto L90
L90:
	;
	v360 = F_objectGetVal(m, v358)
	mBase = m.M
	v362 = F_objectGetVal(m, v358)
	mBase = m.M
	v365 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v362+int32(-1)))))
	switch v365 & int32(7) {
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
		v382 = int32(0)
		goto L91
	}
L91:
	;
	v383 = F_objectGetVal(m, v234)
	mBase = m.M
	v385 = F_objectGetVal(m, v234)
	mBase = m.M
	v388 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v385+int32(-1)))))
	switch v388 & int32(7) {
	case 0:
		goto L102
	case 1:
		goto L101
	case 2:
		goto L100
	case 3:
		goto L99
	case 4:
		goto L98
	default:
		v405 = int32(0)
		goto L97
	}
L92:
	;
	v381 = *(*int32)(unsafe.Add(mBase, uint32(v362+int32(-17))))
	v382 = v381
	goto L91
L93:
	;
	v378 = *(*int32)(unsafe.Add(mBase, uint32(v362+int32(-9))))
	v382 = v378
	goto L91
L94:
	;
	v375 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v362+int32(-5)))))
	v382 = v375
	goto L91
L95:
	;
	v372 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v362+int32(-3)))))
	v382 = v372
	goto L91
L96:
	;
	v382 = int32(base.Ui32(v365) >> (uint(int32(3)) % 32))
	goto L91
L97:
	;
	v406 = int32(0)
	v408 = m.G0
	v409 = int32(16)
	v410 = v408 - v409
	m.G0 = v410
	*(*int32)(unsafe.Add(mBase, uint32(v410)+12)) = v406
	v417 = F_stringmatchlen_impl(m, v360, v382, v383, v405, v406, v410+int32(12), v406)
	mBase = m.M
	m.G0 = v410 + v409
	goto L104
L98:
	;
	v404 = *(*int32)(unsafe.Add(mBase, uint32(v385+int32(-17))))
	v405 = v404
	goto L97
L99:
	;
	v401 = *(*int32)(unsafe.Add(mBase, uint32(v385+int32(-9))))
	v405 = v401
	goto L97
L100:
	;
	v398 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v385+int32(-5)))))
	v405 = v398
	goto L97
L101:
	;
	v395 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v385+int32(-3)))))
	v405 = v395
	goto L97
L102:
	;
	v405 = int32(base.Ui32(v388) >> (uint(int32(3)) % 32))
	goto L97
L103:
	;
	v515 = v230 + int32(20)
	v516 = *(*int32)(unsafe.Add(mBase, uint32(v230)+16))
	if v516 != 0 {
		goto L122
	} else {
		goto L123
	}
L104:
	;
	if v417 == int32(0) {
		v500 = v350
		goto L103
	} else {
		goto L105
	}
L105:
	;
	v424 = v16 + int32(16)
	v425 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v424)+14)) = uint8(v425)
	*(*int32)(unsafe.Add(mBase, uint32(v424))) = v359
	*(*int32)(unsafe.Add(mBase, uint32(v424)+24)) = v425
	*(*uint8)(unsafe.Add(mBase, uint32(v424)+15)) = uint8(v425)
	*(*int32)(unsafe.Add(mBase, uint32(v424)+8)) = int32(-1)
	if v359 == v425 {
		goto L107
	} else {
		goto L108
	}
L106:
	;
	v447 = F_hashtableNext(m, v16+int32(16), v16+int32(12))
	mBase = m.M
	v448 = m.ExcPending
	if v448 != 0 {
		goto L36
	} else {
		goto L111
	}
L107:
	;
	goto L106
L108:
	;
	goto L107
L110:
	;
	F_hashtableCleanupIterator(m, v16+int32(16))
	mBase = m.M
	v494 = m.ExcPending
	if v494 != 0 {
		goto L36
	} else {
		goto L119
	}
L111:
	;
	if v447 == int32(0) {
		v483 = v350
		goto L110
	} else {
		goto L112
	}
L112:
	;
	v456 = v350
	goto L113
L113:
	;
	v464 = *(*int32)(unsafe.Add(mBase, uint32(v16)+12))
	F_addReplyPubsubPatMessage(m, v464, v358, v234, l1)
	mBase = m.M
	v466 = m.ExcPending
	if v466 != 0 {
		goto L36
	} else {
		goto L115
	}
L114:
	;
	v483 = v471
	goto L110
L115:
	;
	v467 = *(*int32)(unsafe.Add(mBase, uint32(v16)+12))
	v468 = F_updateClientMemUsageAndBucket(m, v467)
	mBase = m.M
	v469 = m.ExcPending
	if v469 != 0 {
		goto L36
	} else {
		goto L116
	}
L116:
	;
	v471 = v456 + int32(1)
	v476 = F_hashtableNext(m, v16+int32(16), v16+int32(12))
	mBase = m.M
	v477 = m.ExcPending
	if v477 != 0 {
		goto L36
	} else {
		goto L117
	}
L117:
	;
	if v476 != 0 {
		v456 = v471
		goto L113
	} else {
		goto L118
	}
L118:
	;
	goto L114
L119:
	;
	v500 = v483
	goto L103
L120:
	;
	if v611 != 0 {
		v347 = v611
		v350 = v500
		goto L87
	} else {
		goto L146
	}
L121:
	;
	v522 = v515
	v523 = v519
	goto L124
L122:
	;
	v519 = int32(1)
	goto L121
L123:
	;
	v519 = int32(0)
	goto L121
L124:
	;
	switch v523 {
	case 0:
		goto L129
	default:
		goto L128
	}
L126:
	;
	v523 = int32(0)
	goto L124
L127:
	;
	goto L120
L128:
	;
	v603 = *(*int32)(unsafe.Add(mBase, uint32(v522)))
	*(*int32)(unsafe.Add(mBase, uint32(v230)+16)) = v603
	if v603 == int32(0) {
		goto L126
	} else {
		goto L145
	}
L129:
	;
	v527 = *(*int32)(unsafe.Add(mBase, uint32(v230)+4))
	if v527 != int32(-1) {
		v566 = v527
		goto L130
	} else {
		goto L131
	}
L130:
	;
	v567 = int32(1)
	v568 = v566 + v567
	*(*int32)(unsafe.Add(mBase, uint32(v230)+4)) = v568
	v570 = int32(0)
	v573 = *(*int32)(unsafe.Add(mBase, uint32(v230)))
	v574 = *(*int32)(unsafe.Add(mBase, uint32(v230)+8))
	v578 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v573+v574+int32(26)))))
	if v578 == int32(255) {
		goto L139
	} else {
		goto L140
	}
L131:
	;
	v531 = *(*int32)(unsafe.Add(mBase, uint32(v230)+8))
	if v531 != 0 {
		v566 = int32(-1)
		goto L130
	} else {
		goto L132
	}
L132:
	;
	v532 = *(*int32)(unsafe.Add(mBase, uint32(v230)))
	v533 = *(*int32)(unsafe.Add(mBase, uint32(v230)+12))
	if v533 == int32(0) {
		goto L134
	} else {
		goto L135
	}
L133:
	;
	v560 = *(*int32)(unsafe.Add(mBase, uint32(v559)+20))
	if v560 != int32(-1) {
		goto L136
	} else {
		goto L137
	}
L134:
	;
	v540 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v532)+16)))
	v541 = int64(*(*int8)(unsafe.Add(mBase, uint32(v532)+27)))
	v542 = int64(*(*int32)(unsafe.Add(mBase, uint32(v532)+8)))
	v543 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v532)+12)))
	v544 = int64(*(*int8)(unsafe.Add(mBase, uint32(v532)+26)))
	v545 = int64(*(*int32)(unsafe.Add(mBase, uint32(v532)+4)))
	v546 = F_wangHash64(m, v545)
	mBase = m.M
	v548 = F_wangHash64(m, v544+v546)
	mBase = m.M
	v550 = F_wangHash64(m, v543+v548)
	mBase = m.M
	v552 = F_wangHash64(m, v542+v550)
	mBase = m.M
	v554 = F_wangHash64(m, v541+v552)
	mBase = m.M
	v556 = F_wangHash64(m, v540+v554)
	mBase = m.M
	*(*int64)(unsafe.Add(mBase, uint32(v230)+24)) = v556
	v558 = *(*int32)(unsafe.Add(mBase, uint32(v230)))
	v559 = v558
	goto L133
L135:
	;
	v536 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v532)+24)))
	v538 = v536 + int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v532)+24)) = uint16(v538)
	v559 = v532
	goto L133
L136:
	;
	v566 = v560 + int32(-1)
	goto L130
L137:
	;
	v563 = *(*int32)(unsafe.Add(mBase, uint32(v230)+4))
	v566 = v563
	goto L130
L138:
	;
	v593 = int32(2)
	v598 = *(*int32)(unsafe.Add(mBase, uint32(v573+v591<<(uint(v593)%32)+int32(4))))
	v522 = v598 + v592<<(uint(v593)%32)
	v523 = int32(1)
	goto L124
L139:
	;
	v582 = v570
	goto L141
L140:
	;
	v582 = v567 << (uint(v578) % 32)
	goto L141
L141:
	;
	if v568 < v582 {
		v591 = v574
		v592 = v568
		goto L138
	} else {
		goto L142
	}
L142:
	;
	if v574 != 0 {
		v611 = v570
		goto L127
	} else {
		goto L143
	}
L143:
	;
	v584 = *(*int32)(unsafe.Add(mBase, uint32(v573)+20))
	if v584 == int32(-1) {
		v611 = v570
		goto L127
	} else {
		goto L144
	}
L144:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v230)+4)) = int64(4294967296)
	v591 = int32(1)
	v592 = int32(0)
	goto L138
L145:
	;
	v607 = *(*int32)(unsafe.Add(mBase, uint32(v603)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v515))) = v607
	v611 = v603
	goto L127
L146:
	;
	goto L88
L147:
	;
	F_dictReleaseIterator(m, v230)
	mBase = m.M
	v631 = m.ExcPending
	if v631 != 0 {
		goto L36
	} else {
		goto L148
	}
L148:
	;
	v637 = v620
	goto L54
}
func F_pubsubShardUnsubscribeAllChannelsInSlot(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v104 int64
	_ = v104
	var v107 int64
	_ = v107
	var v110 int64
	_ = v110
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v133 int32
	_ = v133
	var v141 int32
	_ = v141
	var v143 int32
	_ = v143
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v166 int32
	_ = v166
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v185 int32
	_ = v185
	var v202 int32
	_ = v202
	v10 = m.G0
	v12 = v10 - int32(96)
	m.G0 = v12
	v15 = *(*int32)(unsafe.Add(mBase, _consts[20]))
	v16 = *(*int32)(unsafe.Add(mBase, uint32(v15)+8))
	v20 = *(*int32)(unsafe.Add(mBase, uint32(v16+l0<<(uint(int32(2))%32))))
	if v20 != 0 {
		goto L4
	} else {
		goto L5
	}
L1:
	;
	F__serverAssertWithInfo(m, v93, v55, int32(_a1107), int32(_a1108), int32(384))
	mBase = m.M
	v202 = m.ExcPending
	if v202 != 0 {
		goto L8
	} else {
		goto L40
	}
L2:
	;
	m.G0 = v12 + int32(96)
	return
L3:
	;
	if v23 == int32(0) {
		goto L2
	} else {
		goto L6
	}
L4:
	;
	v22 = F_hashtableSize(m, v20)
	mBase = m.M
	v23 = v22
	goto L3
L5:
	;
	v23 = int32(0)
	goto L3
L6:
	;
	v27 = *(*int32)(unsafe.Add(mBase, _consts[20]))
	v29 = F_kvstoreGetHashtableIterator(m, v27, l0, int32(1))
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		goto L8
	} else {
		goto L9
	}
L7:
	;
	F_kvstoreReleaseHashtableIterator(m, v29)
	mBase = m.M
	v185 = m.ExcPending
	if v185 != 0 {
		goto L8
	} else {
		goto L39
	}
L8:
	;
	return
L9:
	;
	v33 = F_kvstoreHashtableIteratorNext(m, v29, v12+int32(92))
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
		goto L8
	} else {
		goto L10
	}
L10:
	;
	if v33 == int32(0) {
		goto L7
	} else {
		goto L11
	}
L11:
	;
	goto L12
L12:
	;
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v12)+92))
	goto L14
L13:
	;
	goto L7
L14:
	;
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v52+int32(44))))
	v57 = v12 + int32(32)
	v58 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v57)+14)) = uint8(v58)
	*(*int32)(unsafe.Add(mBase, uint32(v57))) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v57)+24)) = v58
	*(*uint8)(unsafe.Add(mBase, uint32(v57)+15)) = uint8(v58)
	*(*int32)(unsafe.Add(mBase, uint32(v57)+8)) = int32(-1)
	if v52 == v58 {
		goto L16
	} else {
		goto L17
	}
L15:
	;
	v80 = F_hashtableNext(m, v12+int32(32), v12+int32(28))
	mBase = m.M
	v81 = m.ExcPending
	if v81 != 0 {
		goto L8
	} else {
		goto L20
	}
L16:
	;
	goto L15
L17:
	;
	goto L16
L19:
	;
	F_hashtableCleanupIterator(m, v12+int32(32))
	mBase = m.M
	v166 = m.ExcPending
	if v166 != 0 {
		goto L8
	} else {
		goto L35
	}
L20:
	;
	if v80 == int32(0) {
		goto L19
	} else {
		goto L21
	}
L21:
	;
	goto L22
L22:
	;
	v93 = *(*int32)(unsafe.Add(mBase, uint32(v12)+28))
	v94 = *(*int32)(unsafe.Add(mBase, uint32(v93)+100))
	v95 = *(*int32)(unsafe.Add(mBase, uint32(v94)+8))
	v96 = F_hashtableDelete(m, v95, v55)
	mBase = m.M
	v97 = m.ExcPending
	if v97 != 0 {
		goto L8
	} else {
		goto L24
	}
L23:
	;
	goto L19
L24:
	;
	if v96 == int32(0) {
		goto L1
	} else {
		goto L25
	}
L25:
	;
	v100 = int32(0)
	v101 = *(*int32)(unsafe.Add(mBase, _consts[619]))
	*(*int32)(unsafe.Add(mBase, uint32(v12+int32(24)))) = v101
	v104 = *(*int64)(unsafe.Add(mBase, _consts[620]))
	*(*int64)(unsafe.Add(mBase, uint32(v12+int32(16)))) = v104
	v107 = *(*int64)(unsafe.Add(mBase, _consts[621]))
	*(*int64)(unsafe.Add(mBase, uint32(v12+int32(8)))) = v107
	v110 = *(*int64)(unsafe.Add(mBase, _consts[622]))
	*(*int64)(unsafe.Add(mBase, uint32(v12))) = v110
	F_addReplyPubsubUnsubscribed(m, v93, v55, v12)
	mBase = m.M
	v113 = m.ExcPending
	if v113 != 0 {
		goto L8
	} else {
		goto L26
	}
L26:
	;
	v114 = *(*int32)(unsafe.Add(mBase, uint32(v93)+100))
	v115 = *(*int32)(unsafe.Add(mBase, uint32(v114)))
	v116 = *(*int32)(unsafe.Add(mBase, uint32(v115)+20))
	v117 = *(*int32)(unsafe.Add(mBase, uint32(v115)+16))
	goto L28
L27:
	;
	v152 = F_hashtableNext(m, v12+int32(32), v12+int32(28))
	mBase = m.M
	v153 = m.ExcPending
	if v153 != 0 {
		goto L8
	} else {
		goto L33
	}
L28:
	;
	v119 = *(*int32)(unsafe.Add(mBase, uint32(v93)+100))
	v120 = *(*int32)(unsafe.Add(mBase, uint32(v119)+4))
	v121 = *(*int32)(unsafe.Add(mBase, uint32(v120)+20))
	v122 = *(*int32)(unsafe.Add(mBase, uint32(v120)+16))
	goto L29
L29:
	;
	v126 = *(*int32)(unsafe.Add(mBase, uint32(v93)+100))
	v127 = *(*int32)(unsafe.Add(mBase, uint32(v126)+8))
	v128 = *(*int32)(unsafe.Add(mBase, uint32(v127)+20))
	v129 = *(*int32)(unsafe.Add(mBase, uint32(v127)+16))
	goto L30
L30:
	;
	if v116+v117+(v121+v122) != int32(0)-(v128+v129) {
		goto L27
	} else {
		goto L31
	}
L31:
	;
	v133 = *(*int32)(unsafe.Add(mBase, uint32(v93)+200))
	if v133&int32(262144) == int32(0) {
		goto L27
	} else {
		goto L32
	}
L32:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v93)+200)) = v133 & int32(-262145)
	v141 = int32(_a44)
	v143 = *(*int32)(unsafe.Add(mBase, _consts[623]))
	*(*int32)(unsafe.Add(mBase, _consts[623])) = v143 + int32(-1)
	goto L27
L33:
	;
	if v152 != 0 {
		goto L22
	} else {
		goto L34
	}
L34:
	;
	goto L23
L35:
	;
	v168 = *(*int32)(unsafe.Add(mBase, _consts[20]))
	v169 = F_kvstoreHashtableDelete(m, v168, l0, v55)
	mBase = m.M
	v170 = m.ExcPending
	if v170 != 0 {
		goto L8
	} else {
		goto L36
	}
L36:
	;
	v173 = F_kvstoreHashtableIteratorNext(m, v29, v12+int32(92))
	mBase = m.M
	v174 = m.ExcPending
	if v174 != 0 {
		goto L8
	} else {
		goto L37
	}
L37:
	;
	if v173 != 0 {
		goto L12
	} else {
		goto L38
	}
L38:
	;
	goto L13
L39:
	;
	goto L2
L40:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_pubsubSubscribePattern(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
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
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
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
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v61 int32
	_ = v61
	var v65 int32
	_ = v65
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+100))
	if v5 != 0 {
		v31 = v5
		v33 = *(*int32)(unsafe.Add(mBase, uint32(v31)+4))
		v34 = F_hashtableAdd(m, v33, l1)
		mBase = m.M
		v35 = m.ExcPending
		if v35 != 0 {
			return int32(0)
		} else {
			if v34 == int32(0) {
				F_addReplyPubsubPatSubscribed(m, l0, l1)
				mBase = m.M
				v65 = m.ExcPending
				if v65 != 0 {
					return int32(0)
				} else {
					return v34
				}
			} else {
				F_incrRefCount(m, l1)
				mBase = m.M
				v39 = m.ExcPending
				if v39 != 0 {
					return int32(0)
				} else {
					v41 = *(*int32)(unsafe.Add(mBase, _consts[18]))
					v42 = F_dictFind(m, v41, l1)
					mBase = m.M
					v43 = m.ExcPending
					if v43 != 0 {
						return int32(0)
					} else {
						if v42 != 0 {
							v53 = *(*int32)(unsafe.Add(mBase, uint32(v42)+8))
							v54 = v53
							v55 = F_hashtableAdd(m, v54, l0)
							mBase = m.M
							v56 = m.ExcPending
							if v56 != 0 {
								return int32(0)
							} else {
								if v55 != 0 {
									F_addReplyPubsubPatSubscribed(m, l0, l1)
									mBase = m.M
									v65 = m.ExcPending
									if v65 != 0 {
										return int32(0)
									} else {
										return v34
									}
								} else {
									F__serverAssert(m, int32(_a1109), int32(_a1108), int32(415))
									mBase = m.M
									v61 = m.ExcPending
									if v61 != 0 {
										return int32(0)
									} else {
										F_abort(m)
										mBase = m.M
										base.Wasm_trap_unreachable()
										for {
										}
									}
								}
							}
						} else {
							v45 = F_hashtableCreate(m, int32(_a1110))
							mBase = m.M
							v46 = m.ExcPending
							if v46 != 0 {
								return int32(0)
							} else {
								v48 = *(*int32)(unsafe.Add(mBase, _consts[18]))
								v49 = F_dictAdd(m, v48, l1, v45)
								mBase = m.M
								v50 = m.ExcPending
								if v50 != 0 {
									return int32(0)
								} else {
									F_incrRefCount(m, l1)
									mBase = m.M
									v52 = m.ExcPending
									if v52 != 0 {
										return int32(0)
									} else {
										v54 = v45
										v55 = F_hashtableAdd(m, v54, l0)
										mBase = m.M
										v56 = m.ExcPending
										if v56 != 0 {
											return int32(0)
										} else {
											if v55 != 0 {
												F_addReplyPubsubPatSubscribed(m, l0, l1)
												mBase = m.M
												v65 = m.ExcPending
												if v65 != 0 {
													return int32(0)
												} else {
													return v34
												}
											} else {
												F__serverAssert(m, int32(_a1109), int32(_a1108), int32(415))
												mBase = m.M
												v61 = m.ExcPending
												if v61 != 0 {
													return int32(0)
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
								}
							}
						}
					}
				}
			}
		}
	} else {
		v7 = F_valkey_malloc(m, int32(32))
		mBase = m.M
		v10 = m.ExcPending
		if v10 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(l0)+100)) = v7
			v13 = F_hashtableCreate(m, int32(_a1106))
			mBase = m.M
			v14 = m.ExcPending
			if v14 != 0 {
				return int32(0)
			} else {
				v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+100))
				*(*int32)(unsafe.Add(mBase, uint32(v15))) = v13
				v18 = F_hashtableCreate(m, int32(_a1106))
				mBase = m.M
				v19 = m.ExcPending
				if v19 != 0 {
					return int32(0)
				} else {
					v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+100))
					*(*int32)(unsafe.Add(mBase, uint32(v20)+4)) = v18
					v23 = F_hashtableCreate(m, int32(_a1106))
					mBase = m.M
					v24 = m.ExcPending
					if v24 != 0 {
						return int32(0)
					} else {
						v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+100))
						*(*int32)(unsafe.Add(mBase, uint32(v25)+24)) = int32(0)
						*(*int64)(unsafe.Add(mBase, uint32(v25)+16)) = int64(0)
						*(*int32)(unsafe.Add(mBase, uint32(v25)+8)) = v23
						v31 = v25
						v33 = *(*int32)(unsafe.Add(mBase, uint32(v31)+4))
						v34 = F_hashtableAdd(m, v33, l1)
						mBase = m.M
						v35 = m.ExcPending
						if v35 != 0 {
							return int32(0)
						} else {
							if v34 == int32(0) {
								F_addReplyPubsubPatSubscribed(m, l0, l1)
								mBase = m.M
								v65 = m.ExcPending
								if v65 != 0 {
									return int32(0)
								} else {
									return v34
								}
							} else {
								F_incrRefCount(m, l1)
								mBase = m.M
								v39 = m.ExcPending
								if v39 != 0 {
									return int32(0)
								} else {
									v41 = *(*int32)(unsafe.Add(mBase, _consts[18]))
									v42 = F_dictFind(m, v41, l1)
									mBase = m.M
									v43 = m.ExcPending
									if v43 != 0 {
										return int32(0)
									} else {
										if v42 != 0 {
											v53 = *(*int32)(unsafe.Add(mBase, uint32(v42)+8))
											v54 = v53
											v55 = F_hashtableAdd(m, v54, l0)
											mBase = m.M
											v56 = m.ExcPending
											if v56 != 0 {
												return int32(0)
											} else {
												if v55 != 0 {
													F_addReplyPubsubPatSubscribed(m, l0, l1)
													mBase = m.M
													v65 = m.ExcPending
													if v65 != 0 {
														return int32(0)
													} else {
														return v34
													}
												} else {
													F__serverAssert(m, int32(_a1109), int32(_a1108), int32(415))
													mBase = m.M
													v61 = m.ExcPending
													if v61 != 0 {
														return int32(0)
													} else {
														F_abort(m)
														mBase = m.M
														base.Wasm_trap_unreachable()
														for {
														}
													}
												}
											}
										} else {
											v45 = F_hashtableCreate(m, int32(_a1110))
											mBase = m.M
											v46 = m.ExcPending
											if v46 != 0 {
												return int32(0)
											} else {
												v48 = *(*int32)(unsafe.Add(mBase, _consts[18]))
												v49 = F_dictAdd(m, v48, l1, v45)
												mBase = m.M
												v50 = m.ExcPending
												if v50 != 0 {
													return int32(0)
												} else {
													F_incrRefCount(m, l1)
													mBase = m.M
													v52 = m.ExcPending
													if v52 != 0 {
														return int32(0)
													} else {
														v54 = v45
														v55 = F_hashtableAdd(m, v54, l0)
														mBase = m.M
														v56 = m.ExcPending
														if v56 != 0 {
															return int32(0)
														} else {
															if v55 != 0 {
																F_addReplyPubsubPatSubscribed(m, l0, l1)
																mBase = m.M
																v65 = m.ExcPending
																if v65 != 0 {
																	return int32(0)
																} else {
																	return v34
																}
															} else {
																F__serverAssert(m, int32(_a1109), int32(_a1108), int32(415))
																mBase = m.M
																v61 = m.ExcPending
																if v61 != 0 {
																	return int32(0)
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
func F_pubsubTotalSubscriptions(m *base.Module) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v12 int64
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v19 int64
	_ = v19
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v28 int64
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v35 int64
	_ = v35
	v3 = int32(_a44)
	v4 = *(*int32)(unsafe.Add(mBase, _consts[18]))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(v4)+12))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(v4)+16))
	v8 = *(*int32)(unsafe.Add(mBase, _consts[19]))
	v9 = *(*int32)(unsafe.Add(mBase, uint32(v8)+12))
	if v9 == int32(1) {
		v13 = *(*int32)(unsafe.Add(mBase, uint32(v8)+8))
		v14 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
		if v14 != 0 {
			v16 = F_hashtableSize(m, v14)
			mBase = m.M
			v19 = base.I64_extend_i32_u(v16)
		} else {
			v19 = int64(0)
		}
	} else {
		v12 = *(*int64)(unsafe.Add(mBase, uint32(v8)+40))
		v19 = v12
	}
	v24 = *(*int32)(unsafe.Add(mBase, _consts[20]))
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v24)+12))
	if v25 == int32(1) {
		v29 = *(*int32)(unsafe.Add(mBase, uint32(v24)+8))
		v30 = *(*int32)(unsafe.Add(mBase, uint32(v29)))
		if v30 != 0 {
			v32 = F_hashtableSize(m, v30)
			mBase = m.M
			v35 = base.I64_extend_i32_u(v32)
		} else {
			v35 = int64(0)
		}
	} else {
		v28 = *(*int64)(unsafe.Add(mBase, uint32(v24)+40))
		v35 = v28
	}
	return base.I32_wrap_i64(v19 + base.I64_extend_i32_u(v6+v5) + v35)
}
