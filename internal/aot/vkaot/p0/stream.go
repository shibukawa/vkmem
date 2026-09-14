package p0

import (
	base "github.com/shibukawa/valkeymem/internal/aot/vkaot/base"
	"unsafe"
)

func F_rewriteStreamObject(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v15 int32
	_ = v15
	var v22 int32
	_ = v22
	var v23 int64
	_ = v23
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v45 int32
	_ = v45
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int64
	_ = v66
	var v68 int64
	_ = v68
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v85 int32
	_ = v85
	var v88 int32
	_ = v88
	var v91 int32
	_ = v91
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v102 int64
	_ = v102
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
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
	var v137 int32
	_ = v137
	var v140 int64
	_ = v140
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v169 int32
	_ = v169
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
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
	var v205 int64
	_ = v205
	var v207 int64
	_ = v207
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v215 int32
	_ = v215
	var v222 int32
	_ = v222
	var v225 int32
	_ = v225
	var v228 int32
	_ = v228
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
	var v236 int32
	_ = v236
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v247 int32
	_ = v247
	var v248 int32
	_ = v248
	var v255 int32
	_ = v255
	var v267 int32
	_ = v267
	var v268 int32
	_ = v268
	var v273 int32
	_ = v273
	var v274 int32
	_ = v274
	var v277 int32
	_ = v277
	var v278 int32
	_ = v278
	var v283 int32
	_ = v283
	var v284 int32
	_ = v284
	var v289 int32
	_ = v289
	var v290 int32
	_ = v290
	var v293 int64
	_ = v293
	var v294 int32
	_ = v294
	var v295 int32
	_ = v295
	var v300 int32
	_ = v300
	var v301 int32
	_ = v301
	var v306 int32
	_ = v306
	var v307 int32
	_ = v307
	var v311 int32
	_ = v311
	var v313 int32
	_ = v313
	var v317 int32
	_ = v317
	var v323 int64
	_ = v323
	var v338 int32
	_ = v338
	var v340 int32
	_ = v340
	var v341 int32
	_ = v341
	var v344 int32
	_ = v344
	var v345 int32
	_ = v345
	var v356 int32
	_ = v356
	var v359 int32
	_ = v359
	var v360 int32
	_ = v360
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
	var v376 int32
	_ = v376
	var v379 int32
	_ = v379
	var v380 int32
	_ = v380
	var v381 int32
	_ = v381
	var v382 int32
	_ = v382
	var v385 int32
	_ = v385
	var v386 int32
	_ = v386
	var v391 int32
	_ = v391
	var v392 int32
	_ = v392
	var v395 int64
	_ = v395
	var v396 int32
	_ = v396
	var v397 int32
	_ = v397
	var v401 int32
	_ = v401
	var v405 int32
	_ = v405
	var v408 int32
	_ = v408
	var v409 int32
	_ = v409
	var v415 int64
	_ = v415
	var v430 int32
	_ = v430
	var v432 int32
	_ = v432
	var v433 int32
	_ = v433
	var v436 int32
	_ = v436
	var v437 int32
	_ = v437
	var v448 int32
	_ = v448
	var v449 int32
	_ = v449
	var v450 int64
	_ = v450
	var v453 int32
	_ = v453
	var v454 int32
	_ = v454
	var v455 int32
	_ = v455
	var v456 int32
	_ = v456
	var v460 int32
	_ = v460
	var v464 int32
	_ = v464
	var v468 int32
	_ = v468
	var v471 int32
	_ = v471
	var v472 int32
	_ = v472
	var v478 int64
	_ = v478
	var v493 int32
	_ = v493
	var v495 int32
	_ = v495
	var v496 int32
	_ = v496
	var v507 int32
	_ = v507
	var v508 int32
	_ = v508
	var v511 int32
	_ = v511
	var v512 int32
	_ = v512
	var v513 int32
	_ = v513
	var v514 int32
	_ = v514
	var v515 int32
	_ = v515
	var v516 int32
	_ = v516
	var v520 int32
	_ = v520
	var v524 int32
	_ = v524
	var v528 int32
	_ = v528
	var v532 int32
	_ = v532
	var v537 int32
	_ = v537
	var v548 int32
	_ = v548
	var v549 int32
	_ = v549
	var v561 int32
	_ = v561
	var v564 int32
	_ = v564
	var v565 int32
	_ = v565
	var v577 int32
	_ = v577
	var v589 int32
	_ = v589
	var v597 int32
	_ = v597
	var v599 int32
	_ = v599
	v4 = int32(0)
	v9 = m.G0
	v11 = v9 - int32(1424)
	m.G0 = v11
	v15 = F_objectGetVal(m, l2)
	mBase = m.M
	F_streamIteratorStart(m, v11+int32(976), v15, v4, v4, v4)
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
	v23 = *(*int64)(unsafe.Add(mBase, uint32(v15)+8))
	if v23 == int64(0) {
		goto L6
	} else {
		goto L7
	}
L3:
	;
	m.G0 = v11 + int32(1424)
	return v599
L4:
	;
	F_streamIteratorStop(m, v11+int32(976))
	mBase = m.M
	v597 = m.ExcPending
	if v597 != 0 {
		goto L1
	} else {
		goto L155
	}
L5:
	;
	v267 = F_rioWriteBulkCount(m, l0, int32(42), int32(7))
	mBase = m.M
	v268 = m.ExcPending
	if v268 != 0 {
		goto L1
	} else {
		goto L72
	}
L6:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v11)+968)) = int64(1)
	*(*int64)(unsafe.Add(mBase, uint32(v11)+960)) = int64(0)
	v177 = F_rioWriteBulkCount(m, l0, int32(42), int32(7))
	mBase = m.M
	v178 = m.ExcPending
	if v178 != 0 {
		goto L1
	} else {
		goto L44
	}
L7:
	;
	v32 = F_streamIteratorGetID(m, v11+int32(976), v11+int32(960), v11+int32(952))
	mBase = m.M
	v33 = m.ExcPending
	if v33 != 0 {
		goto L1
	} else {
		goto L8
	}
L8:
	;
	if v32 == int32(0) {
		goto L5
	} else {
		goto L9
	}
L9:
	;
	goto L10
L10:
	;
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v11)+952))
	v50 = F_rioWriteBulkCount(m, l0, int32(42), v45<<(uint(int32(1))%32)+int32(3))
	mBase = m.M
	v51 = m.ExcPending
	if v51 != 0 {
		goto L1
	} else {
		goto L12
	}
L11:
	;
	F_streamIteratorStop(m, v11+int32(976))
	mBase = m.M
	v169 = m.ExcPending
	if v169 != 0 {
		goto L1
	} else {
		goto L42
	}
L12:
	;
	if v50 == int32(0) {
		goto L4
	} else {
		goto L13
	}
L13:
	;
	v56 = F_rioWriteBulkString(m, l0, int32(_a114), int32(4))
	mBase = m.M
	v57 = m.ExcPending
	if v57 != 0 {
		goto L1
	} else {
		goto L14
	}
L14:
	;
	if v56 == int32(0) {
		goto L4
	} else {
		goto L15
	}
L15:
	;
	v60 = F_rioWriteBulkObject(m, l0, l1)
	mBase = m.M
	v61 = m.ExcPending
	if v61 != 0 {
		goto L1
	} else {
		goto L16
	}
L16:
	;
	if v60 == int32(0) {
		goto L4
	} else {
		goto L17
	}
L17:
	;
	v64 = F_sdsempty(m)
	mBase = m.M
	v65 = m.ExcPending
	if v65 != 0 {
		goto L1
	} else {
		goto L18
	}
L18:
	;
	v66 = *(*int64)(unsafe.Add(mBase, uint32(v11)+960))
	*(*int64)(unsafe.Add(mBase, uint32(v11)+16)) = v66
	v68 = *(*int64)(unsafe.Add(mBase, uint32(v11)+968))
	*(*int64)(unsafe.Add(mBase, uint32(v11)+24)) = v68
	v74 = F_sdscatfmt(m, v64, int32(_a105), v11+int32(16))
	mBase = m.M
	v75 = m.ExcPending
	if v75 != 0 {
		goto L1
	} else {
		goto L25
	}
L19:
	;
	v96 = F_rioWriteBulkString(m, l0, v74, v95)
	mBase = m.M
	v97 = m.ExcPending
	if v97 != 0 {
		goto L1
	} else {
		goto L26
	}
L20:
	;
	v94 = *(*int32)(unsafe.Add(mBase, uint32(v74+int32(-17))))
	v95 = v94
	goto L19
L21:
	;
	v91 = *(*int32)(unsafe.Add(mBase, uint32(v74+int32(-9))))
	v95 = v91
	goto L19
L22:
	;
	v88 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v74+int32(-5)))))
	v95 = v88
	goto L19
L23:
	;
	v85 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v74+int32(-3)))))
	v95 = v85
	goto L19
L24:
	;
	v95 = int32(base.Ui32(v78) >> (uint(int32(3)) % 32))
	goto L19
L25:
	;
	v78 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v74+int32(-1)))))
	switch v78 & int32(7) {
	case 0:
		goto L24
	case 1:
		goto L23
	case 2:
		goto L22
	case 3:
		goto L21
	case 4:
		goto L20
	default:
		v95 = int32(0)
		goto L19
	}
L26:
	;
	F_sdsfree(m, v74)
	mBase = m.M
	v99 = m.ExcPending
	if v99 != 0 {
		goto L1
	} else {
		goto L27
	}
L27:
	;
	if v96 == int32(0) {
		goto L4
	} else {
		goto L28
	}
L28:
	;
	v102 = *(*int64)(unsafe.Add(mBase, uint32(v11)+952))
	*(*int64)(unsafe.Add(mBase, uint32(v11)+952)) = v102 + int64(-1)
	if v102 == int64(0) {
		goto L30
	} else {
		goto L31
	}
L29:
	;
	goto L11
L30:
	;
	v162 = F_streamIteratorGetID(m, v11+int32(976), v11+int32(960), v11+int32(952))
	mBase = m.M
	v163 = m.ExcPending
	if v163 != 0 {
		goto L1
	} else {
		goto L40
	}
L31:
	;
	goto L32
L32:
	;
	F_streamIteratorGetField(m, v11+int32(976), v11+int32(32), v11+int32(948), v11+int32(640), v11+int32(336))
	mBase = m.M
	v127 = m.ExcPending
	if v127 != 0 {
		goto L1
	} else {
		goto L34
	}
L33:
	;
	goto L30
L34:
	;
	v128 = *(*int32)(unsafe.Add(mBase, uint32(v11)+32))
	v129 = *(*int32)(unsafe.Add(mBase, uint32(v11)+640))
	v130 = F_rioWriteBulkString(m, l0, v128, v129)
	mBase = m.M
	v131 = m.ExcPending
	if v131 != 0 {
		goto L1
	} else {
		goto L35
	}
L35:
	;
	if v130 == int32(0) {
		goto L29
	} else {
		goto L36
	}
L36:
	;
	v134 = *(*int32)(unsafe.Add(mBase, uint32(v11)+948))
	v135 = *(*int32)(unsafe.Add(mBase, uint32(v11)+336))
	v136 = F_rioWriteBulkString(m, l0, v134, v135)
	mBase = m.M
	v137 = m.ExcPending
	if v137 != 0 {
		goto L1
	} else {
		goto L37
	}
L37:
	;
	if v136 == int32(0) {
		goto L29
	} else {
		goto L38
	}
L38:
	;
	v140 = *(*int64)(unsafe.Add(mBase, uint32(v11)+952))
	*(*int64)(unsafe.Add(mBase, uint32(v11)+952)) = v140 + int64(-1)
	if base.B2i32(v140 == int64(0)) == int32(0) {
		goto L32
	} else {
		goto L39
	}
L39:
	;
	goto L33
L40:
	;
	if v162 == int32(0) {
		goto L5
	} else {
		goto L41
	}
L41:
	;
	goto L10
L42:
	;
	v599 = int32(0)
	goto L3
L43:
	;
	F_streamIteratorStop(m, v11+int32(976))
	mBase = m.M
	v255 = m.ExcPending
	if v255 != 0 {
		goto L1
	} else {
		goto L69
	}
L44:
	;
	if v177 == int32(0) {
		goto L43
	} else {
		goto L45
	}
L45:
	;
	v183 = F_rioWriteBulkString(m, l0, int32(_a114), int32(4))
	mBase = m.M
	v184 = m.ExcPending
	if v184 != 0 {
		goto L1
	} else {
		goto L46
	}
L46:
	;
	if v183 == int32(0) {
		goto L43
	} else {
		goto L47
	}
L47:
	;
	v187 = F_rioWriteBulkObject(m, l0, l1)
	mBase = m.M
	v188 = m.ExcPending
	if v188 != 0 {
		goto L1
	} else {
		goto L48
	}
L48:
	;
	if v187 == int32(0) {
		goto L43
	} else {
		goto L49
	}
L49:
	;
	v193 = F_rioWriteBulkString(m, l0, int32(_a115), int32(6))
	mBase = m.M
	v194 = m.ExcPending
	if v194 != 0 {
		goto L1
	} else {
		goto L50
	}
L50:
	;
	if v193 == int32(0) {
		goto L43
	} else {
		goto L51
	}
L51:
	;
	v199 = F_rioWriteBulkString(m, l0, int32(_a107), int32(1))
	mBase = m.M
	v200 = m.ExcPending
	if v200 != 0 {
		goto L1
	} else {
		goto L52
	}
L52:
	;
	if v199 == int32(0) {
		goto L43
	} else {
		goto L53
	}
L53:
	;
	v203 = F_sdsempty(m)
	mBase = m.M
	v204 = m.ExcPending
	if v204 != 0 {
		goto L1
	} else {
		goto L54
	}
L54:
	;
	v205 = *(*int64)(unsafe.Add(mBase, uint32(v11)+960))
	*(*int64)(unsafe.Add(mBase, uint32(v11))) = v205
	v207 = *(*int64)(unsafe.Add(mBase, uint32(v11)+968))
	*(*int64)(unsafe.Add(mBase, uint32(v11)+8)) = v207
	v211 = F_sdscatfmt(m, v203, int32(_a105), v11)
	mBase = m.M
	v212 = m.ExcPending
	if v212 != 0 {
		goto L1
	} else {
		goto L61
	}
L55:
	;
	v233 = F_rioWriteBulkString(m, l0, v211, v232)
	mBase = m.M
	v234 = m.ExcPending
	if v234 != 0 {
		goto L1
	} else {
		goto L62
	}
L56:
	;
	v231 = *(*int32)(unsafe.Add(mBase, uint32(v211+int32(-17))))
	v232 = v231
	goto L55
L57:
	;
	v228 = *(*int32)(unsafe.Add(mBase, uint32(v211+int32(-9))))
	v232 = v228
	goto L55
L58:
	;
	v225 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v211+int32(-5)))))
	v232 = v225
	goto L55
L59:
	;
	v222 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v211+int32(-3)))))
	v232 = v222
	goto L55
L60:
	;
	v232 = int32(base.Ui32(v215) >> (uint(int32(3)) % 32))
	goto L55
L61:
	;
	v215 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v211+int32(-1)))))
	switch v215 & int32(7) {
	case 0:
		goto L60
	case 1:
		goto L59
	case 2:
		goto L58
	case 3:
		goto L57
	case 4:
		goto L56
	default:
		v232 = int32(0)
		goto L55
	}
L62:
	;
	F_sdsfree(m, v211)
	mBase = m.M
	v236 = m.ExcPending
	if v236 != 0 {
		goto L1
	} else {
		goto L63
	}
L63:
	;
	if v233 == int32(0) {
		goto L43
	} else {
		goto L64
	}
L64:
	;
	v241 = F_rioWriteBulkString(m, l0, int32(_a116), int32(1))
	mBase = m.M
	v242 = m.ExcPending
	if v242 != 0 {
		goto L1
	} else {
		goto L65
	}
L65:
	;
	if v241 == int32(0) {
		goto L43
	} else {
		goto L66
	}
L66:
	;
	v247 = F_rioWriteBulkString(m, l0, int32(_a117), int32(1))
	mBase = m.M
	v248 = m.ExcPending
	if v248 != 0 {
		goto L1
	} else {
		goto L67
	}
L67:
	;
	if v247 != 0 {
		goto L5
	} else {
		goto L68
	}
L68:
	;
	goto L43
L69:
	;
	v599 = int32(0)
	goto L3
L70:
	;
	v313 = *(*int32)(unsafe.Add(mBase, uint32(v15)))
	if v313 == int32(0) {
		goto L89
	} else {
		goto L90
	}
L71:
	;
	F_streamIteratorStop(m, v11+int32(976))
	mBase = m.M
	v311 = m.ExcPending
	if v311 != 0 {
		goto L1
	} else {
		goto L88
	}
L72:
	;
	if v267 == int32(0) {
		goto L71
	} else {
		goto L73
	}
L73:
	;
	v273 = F_rioWriteBulkString(m, l0, int32(_a118), int32(6))
	mBase = m.M
	v274 = m.ExcPending
	if v274 != 0 {
		goto L1
	} else {
		goto L74
	}
L74:
	;
	if v273 == int32(0) {
		goto L71
	} else {
		goto L75
	}
L75:
	;
	v277 = F_rioWriteBulkObject(m, l0, l1)
	mBase = m.M
	v278 = m.ExcPending
	if v278 != 0 {
		goto L1
	} else {
		goto L76
	}
L76:
	;
	if v277 == int32(0) {
		goto L71
	} else {
		goto L77
	}
L77:
	;
	v283 = F_rioWriteBulkStreamID(m, l0, v15+int32(16))
	mBase = m.M
	v284 = m.ExcPending
	if v284 != 0 {
		goto L1
	} else {
		goto L78
	}
L78:
	;
	if v283 == int32(0) {
		goto L71
	} else {
		goto L79
	}
L79:
	;
	v289 = F_rioWriteBulkString(m, l0, int32(_a119), int32(12))
	mBase = m.M
	v290 = m.ExcPending
	if v290 != 0 {
		goto L1
	} else {
		goto L80
	}
L80:
	;
	if v289 == int32(0) {
		goto L71
	} else {
		goto L81
	}
L81:
	;
	v293 = *(*int64)(unsafe.Add(mBase, uint32(v15)+64))
	v294 = F_rioWriteBulkLongLong(m, l0, v293)
	mBase = m.M
	v295 = m.ExcPending
	if v295 != 0 {
		goto L1
	} else {
		goto L82
	}
L82:
	;
	if v294 == int32(0) {
		goto L71
	} else {
		goto L83
	}
L83:
	;
	v300 = F_rioWriteBulkString(m, l0, int32(_a120), int32(12))
	mBase = m.M
	v301 = m.ExcPending
	if v301 != 0 {
		goto L1
	} else {
		goto L84
	}
L84:
	;
	if v300 == int32(0) {
		goto L71
	} else {
		goto L85
	}
L85:
	;
	v306 = F_rioWriteBulkStreamID(m, l0, v15+int32(48))
	mBase = m.M
	v307 = m.ExcPending
	if v307 != 0 {
		goto L1
	} else {
		goto L86
	}
L86:
	;
	if v306 != 0 {
		goto L70
	} else {
		goto L87
	}
L87:
	;
	goto L71
L88:
	;
	v599 = int32(0)
	goto L3
L89:
	;
	F_streamIteratorStop(m, v11+int32(976))
	mBase = m.M
	v589 = m.ExcPending
	if v589 != 0 {
		goto L1
	} else {
		goto L154
	}
L90:
	;
	v317 = v11 + int32(640)
	*(*int32)(unsafe.Add(mBase, uint32(v317)+4)) = v313
	*(*int32)(unsafe.Add(mBase, uint32(v317))) = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v317)+20)) = int32(128)
	v323 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v317)+12)) = v323
	*(*int64)(unsafe.Add(mBase, uint32(v317)+296)) = v323
	*(*int64)(unsafe.Add(mBase, uint32(v317)+160)) = int64(137438953472)
	*(*int32)(unsafe.Add(mBase, uint32(v317)+8)) = v11 + int32(664)
	*(*int32)(unsafe.Add(mBase, uint32(v317)+156)) = v11 + int32(808)
	goto L91
L91:
	;
	v338 = int32(0)
	v340 = F_raxSeek(m, v11+int32(640), int32(_a4), v338, v338)
	mBase = m.M
	v341 = m.ExcPending
	if v341 != 0 {
		goto L1
	} else {
		goto L92
	}
L92:
	;
	v344 = F_raxNext(m, v11+int32(640))
	mBase = m.M
	v345 = m.ExcPending
	if v345 != 0 {
		goto L1
	} else {
		goto L94
	}
L93:
	;
	F_raxStop(m, v11+int32(640))
	mBase = m.M
	v577 = m.ExcPending
	if v577 != 0 {
		goto L1
	} else {
		goto L153
	}
L94:
	;
	if v344 == int32(0) {
		goto L93
	} else {
		goto L95
	}
L95:
	;
	goto L96
L96:
	;
	v356 = *(*int32)(unsafe.Add(mBase, uint32(v11)+652))
	v359 = F_rioWriteBulkCount(m, l0, int32(42), int32(7))
	mBase = m.M
	v360 = m.ExcPending
	if v360 != 0 {
		goto L1
	} else {
		goto L100
	}
L97:
	;
	goto L93
L98:
	;
	v408 = v11 + int32(336)
	v409 = *(*int32)(unsafe.Add(mBase, uint32(v356)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v408)+4)) = v409
	*(*int32)(unsafe.Add(mBase, uint32(v408))) = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v408)+20)) = int32(128)
	v415 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v408)+12)) = v415
	*(*int64)(unsafe.Add(mBase, uint32(v408)+296)) = v415
	*(*int64)(unsafe.Add(mBase, uint32(v408)+160)) = int64(137438953472)
	*(*int32)(unsafe.Add(mBase, uint32(v408)+8)) = v11 + int32(360)
	*(*int32)(unsafe.Add(mBase, uint32(v408)+156)) = v11 + int32(504)
	goto L118
L99:
	;
	F_raxStop(m, v11+int32(640))
	mBase = m.M
	v401 = m.ExcPending
	if v401 != 0 {
		goto L1
	} else {
		goto L116
	}
L100:
	;
	if v359 == int32(0) {
		goto L99
	} else {
		goto L101
	}
L101:
	;
	v365 = F_rioWriteBulkString(m, l0, int32(_a112), int32(6))
	mBase = m.M
	v366 = m.ExcPending
	if v366 != 0 {
		goto L1
	} else {
		goto L102
	}
L102:
	;
	if v365 == int32(0) {
		goto L99
	} else {
		goto L103
	}
L103:
	;
	v371 = F_rioWriteBulkString(m, l0, int32(_a121), int32(6))
	mBase = m.M
	v372 = m.ExcPending
	if v372 != 0 {
		goto L1
	} else {
		goto L104
	}
L104:
	;
	if v371 == int32(0) {
		goto L99
	} else {
		goto L105
	}
L105:
	;
	v375 = F_rioWriteBulkObject(m, l0, l1)
	mBase = m.M
	v376 = m.ExcPending
	if v376 != 0 {
		goto L1
	} else {
		goto L106
	}
L106:
	;
	if v375 == int32(0) {
		goto L99
	} else {
		goto L107
	}
L107:
	;
	v379 = *(*int32)(unsafe.Add(mBase, uint32(v11)+648))
	v380 = *(*int32)(unsafe.Add(mBase, uint32(v11)+656))
	v381 = F_rioWriteBulkString(m, l0, v379, v380)
	mBase = m.M
	v382 = m.ExcPending
	if v382 != 0 {
		goto L1
	} else {
		goto L108
	}
L108:
	;
	if v381 == int32(0) {
		goto L99
	} else {
		goto L109
	}
L109:
	;
	v385 = F_rioWriteBulkStreamID(m, l0, v356)
	mBase = m.M
	v386 = m.ExcPending
	if v386 != 0 {
		goto L1
	} else {
		goto L110
	}
L110:
	;
	if v385 == int32(0) {
		goto L99
	} else {
		goto L111
	}
L111:
	;
	v391 = F_rioWriteBulkString(m, l0, int32(_a122), int32(11))
	mBase = m.M
	v392 = m.ExcPending
	if v392 != 0 {
		goto L1
	} else {
		goto L112
	}
L112:
	;
	if v391 == int32(0) {
		goto L99
	} else {
		goto L113
	}
L113:
	;
	v395 = *(*int64)(unsafe.Add(mBase, uint32(v356)+16))
	v396 = F_rioWriteBulkLongLong(m, l0, v395)
	mBase = m.M
	v397 = m.ExcPending
	if v397 != 0 {
		goto L1
	} else {
		goto L114
	}
L114:
	;
	if v396 != 0 {
		goto L98
	} else {
		goto L115
	}
L115:
	;
	goto L99
L116:
	;
	F_streamIteratorStop(m, v11+int32(976))
	mBase = m.M
	v405 = m.ExcPending
	if v405 != 0 {
		goto L1
	} else {
		goto L117
	}
L117:
	;
	v599 = int32(0)
	goto L3
L118:
	;
	v430 = int32(0)
	v432 = F_raxSeek(m, v11+int32(336), int32(_a4), v430, v430)
	mBase = m.M
	v433 = m.ExcPending
	if v433 != 0 {
		goto L1
	} else {
		goto L119
	}
L119:
	;
	v436 = F_raxNext(m, v11+int32(336))
	mBase = m.M
	v437 = m.ExcPending
	if v437 != 0 {
		goto L1
	} else {
		goto L121
	}
L120:
	;
	F_raxStop(m, v11+int32(336))
	mBase = m.M
	v561 = m.ExcPending
	if v561 != 0 {
		goto L1
	} else {
		goto L150
	}
L121:
	;
	if v436 == int32(0) {
		goto L120
	} else {
		goto L122
	}
L122:
	;
	goto L123
L123:
	;
	v448 = *(*int32)(unsafe.Add(mBase, uint32(v11)+348))
	v449 = *(*int32)(unsafe.Add(mBase, uint32(v448)+20))
	v450 = *(*int64)(unsafe.Add(mBase, uint32(v449)+8))
	goto L127
L124:
	;
	goto L120
L125:
	;
	v548 = F_raxNext(m, v11+int32(336))
	mBase = m.M
	v549 = m.ExcPending
	if v549 != 0 {
		goto L1
	} else {
		goto L148
	}
L126:
	;
	v471 = v11 + int32(32)
	v472 = *(*int32)(unsafe.Add(mBase, uint32(v448)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v471)+4)) = v472
	*(*int32)(unsafe.Add(mBase, uint32(v471))) = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v471)+20)) = int32(128)
	v478 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v471)+12)) = v478
	*(*int64)(unsafe.Add(mBase, uint32(v471)+296)) = v478
	*(*int64)(unsafe.Add(mBase, uint32(v471)+160)) = int64(137438953472)
	*(*int32)(unsafe.Add(mBase, uint32(v471)+8)) = v11 + int32(56)
	*(*int32)(unsafe.Add(mBase, uint32(v471)+156)) = v11 + int32(200)
	goto L134
L127:
	;
	if v450 != int64(0) {
		goto L126
	} else {
		goto L128
	}
L128:
	;
	v453 = *(*int32)(unsafe.Add(mBase, uint32(v11)+648))
	v454 = *(*int32)(unsafe.Add(mBase, uint32(v11)+656))
	v455 = F_rioWriteStreamEmptyConsumer(m, l0, l1, v453, v454, v448)
	mBase = m.M
	v456 = m.ExcPending
	if v456 != 0 {
		goto L1
	} else {
		goto L129
	}
L129:
	;
	if v455 != 0 {
		goto L125
	} else {
		goto L130
	}
L130:
	;
	F_raxStop(m, v11+int32(336))
	mBase = m.M
	v460 = m.ExcPending
	if v460 != 0 {
		goto L1
	} else {
		goto L131
	}
L131:
	;
	F_raxStop(m, v11+int32(640))
	mBase = m.M
	v464 = m.ExcPending
	if v464 != 0 {
		goto L1
	} else {
		goto L132
	}
L132:
	;
	F_streamIteratorStop(m, v11+int32(976))
	mBase = m.M
	v468 = m.ExcPending
	if v468 != 0 {
		goto L1
	} else {
		goto L133
	}
L133:
	;
	v599 = int32(0)
	goto L3
L134:
	;
	v493 = int32(0)
	v495 = F_raxSeek(m, v11+int32(32), int32(_a4), v493, v493)
	mBase = m.M
	v496 = m.ExcPending
	if v496 != 0 {
		goto L1
	} else {
		goto L135
	}
L135:
	;
	goto L137
L136:
	;
	F_raxStop(m, v11+int32(32))
	mBase = m.M
	v537 = m.ExcPending
	if v537 != 0 {
		goto L1
	} else {
		goto L147
	}
L137:
	;
	v507 = F_raxNext(m, v11+int32(32))
	mBase = m.M
	v508 = m.ExcPending
	if v508 != 0 {
		goto L1
	} else {
		goto L139
	}
L138:
	;
	F_raxStop(m, v11+int32(32))
	mBase = m.M
	v520 = m.ExcPending
	if v520 != 0 {
		goto L1
	} else {
		goto L143
	}
L139:
	;
	if v507 == int32(0) {
		goto L136
	} else {
		goto L140
	}
L140:
	;
	v511 = *(*int32)(unsafe.Add(mBase, uint32(v11)+648))
	v512 = *(*int32)(unsafe.Add(mBase, uint32(v11)+656))
	v513 = *(*int32)(unsafe.Add(mBase, uint32(v11)+40))
	v514 = *(*int32)(unsafe.Add(mBase, uint32(v11)+44))
	v515 = F_rioWriteStreamPendingEntry(m, l0, l1, v511, v512, v448, v513, v514)
	mBase = m.M
	v516 = m.ExcPending
	if v516 != 0 {
		goto L1
	} else {
		goto L141
	}
L141:
	;
	if v515 != 0 {
		goto L137
	} else {
		goto L142
	}
L142:
	;
	goto L138
L143:
	;
	F_raxStop(m, v11+int32(336))
	mBase = m.M
	v524 = m.ExcPending
	if v524 != 0 {
		goto L1
	} else {
		goto L144
	}
L144:
	;
	F_raxStop(m, v11+int32(640))
	mBase = m.M
	v528 = m.ExcPending
	if v528 != 0 {
		goto L1
	} else {
		goto L145
	}
L145:
	;
	F_streamIteratorStop(m, v11+int32(976))
	mBase = m.M
	v532 = m.ExcPending
	if v532 != 0 {
		goto L1
	} else {
		goto L146
	}
L146:
	;
	v599 = int32(0)
	goto L3
L147:
	;
	goto L125
L148:
	;
	if v548 != 0 {
		goto L123
	} else {
		goto L149
	}
L149:
	;
	goto L124
L150:
	;
	v564 = F_raxNext(m, v11+int32(640))
	mBase = m.M
	v565 = m.ExcPending
	if v565 != 0 {
		goto L1
	} else {
		goto L151
	}
L151:
	;
	if v564 != 0 {
		goto L96
	} else {
		goto L152
	}
L152:
	;
	goto L97
L153:
	;
	goto L89
L154:
	;
	v599 = int32(1)
	goto L3
L155:
	;
	v599 = int32(0)
	goto L3
}
func F_streamCreateConsumer(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v20 int32
	_ = v20
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v56 int64
	_ = v56
	var v60 int32
	_ = v60
	var v62 int64
	_ = v62
	var v69 int32
	_ = v69
	if l0 != 0 {
		v12 = F_valkey_malloc(m, int32(24))
		mBase = m.M
		v15 = m.ExcPending
		if v15 != 0 {
			return int32(0)
		} else {
			v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
			v20 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1+int32(-1)))))
			switch v20 & int32(7) {
			case 0:
				v37 = int32(base.Ui32(v20) >> (uint(int32(3)) % 32))
			case 1:
				v27 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1+int32(-3)))))
				v37 = v27
			case 2:
				v30 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1+int32(-5)))))
				v37 = v30
			case 3:
				v33 = *(*int32)(unsafe.Add(mBase, uint32(l1+int32(-9))))
				v37 = v33
			case 4:
				v36 = *(*int32)(unsafe.Add(mBase, uint32(l1+int32(-17))))
				v37 = v36
			default:
				v37 = int32(0)
			}
			v39 = F_raxTryInsert(m, v16, l1, v37, v12, int32(0))
			mBase = m.M
			v40 = m.ExcPending
			if v40 != 0 {
				return int32(0)
			} else {
				if v39 != 0 {
					v47 = F_sdsdup(m, l1)
					mBase = m.M
					v48 = m.ExcPending
					if v48 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = v47
						v50 = F_raxNew(m)
						mBase = m.M
						v51 = m.ExcPending
						if v51 != 0 {
							return int32(0)
						} else {
							*(*int64)(unsafe.Add(mBase, uint32(v12)+8)) = int64(-1)
							*(*int32)(unsafe.Add(mBase, uint32(v12)+20)) = v50
							v56 = *(*int64)(unsafe.Add(mBase, _consts[78]))
							*(*int64)(unsafe.Add(mBase, uint32(v12))) = v56
							if l4&int32(2) != 0 {
							} else {
								v60 = int32(_a69)
								v62 = *(*int64)(unsafe.Add(mBase, _consts[60]))
								*(*int64)(unsafe.Add(mBase, _consts[60])) = v62 + int64(1)
							}
							if l4&int32(1) != 0 {
								return v12
							} else {
								F_notifyKeyspaceEvent(m, int32(1024), int32(_a1547), l2, l3)
								mBase = m.M
								v69 = m.ExcPending
								if v69 != 0 {
									return int32(0)
								} else {
									return v12
								}
							}
						}
					}
				} else {
					F_valkey_free(m, v12)
					mBase = m.M
					v42 = m.ExcPending
					if v42 != 0 {
						return int32(0)
					} else {
						return int32(0)
					}
				}
			}
		}
	} else {
		return int32(0)
	}
}
func F_streamDeleteItem(m *base.Module, l0 int32, l1 int32) int32 {
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v16 int32
	_ = v16
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v38 int32
	_ = v38
	v3 = int32(0)
	v5 = m.G0
	v7 = v5 - int32(480)
	m.G0 = v7
	F_streamIteratorStart(m, v7+int32(32), l0, l1, l1, v3)
	v16 = m.ExcPending
	if v16 != 0 {
		return int32(0)
	} else {
		v23 = F_streamIteratorGetID(m, v7+int32(32), v7+int32(16), v7+int32(8))
		v24 = m.ExcPending
		if v24 != 0 {
			return int32(0)
		} else {
			if v23 == int32(0) {
				v34 = v3
				F_raxStop(m, v7+int32(120))
				v38 = m.ExcPending
				if v38 != 0 {
					return int32(0)
				} else {
					m.G0 = v7 + int32(480)
					return v34
				}
			} else {
				F_streamIteratorRemoveEntry(m, v7+int32(32), v7+int32(16))
				v32 = m.ExcPending
				if v32 != 0 {
					return int32(0)
				} else {
					v34 = int32(1)
					F_raxStop(m, v7+int32(120))
					v38 = m.ExcPending
					if v38 != 0 {
						return int32(0)
					} else {
						m.G0 = v7 + int32(480)
						return v34
					}
				}
			}
		}
	}
}
func F_streamDup(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v38 int64
	_ = v38
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v82 int64
	_ = v82
	var v86 int64
	_ = v86
	var v88 int32
	_ = v88
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v109 int64
	_ = v109
	var v111 int32
	_ = v111
	var v115 int64
	_ = v115
	var v117 int64
	_ = v117
	var v119 int32
	_ = v119
	var v123 int64
	_ = v123
	var v125 int64
	_ = v125
	var v127 int64
	_ = v127
	var v129 int32
	_ = v129
	var v133 int64
	_ = v133
	var v135 int64
	_ = v135
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v145 int32
	_ = v145
	var v151 int64
	_ = v151
	var v166 int32
	_ = v166
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v189 int64
	_ = v189
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v202 int64
	_ = v202
	var v217 int32
	_ = v217
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v237 int32
	_ = v237
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v242 int64
	_ = v242
	var v243 int32
	_ = v243
	var v248 int64
	_ = v248
	var v250 int64
	_ = v250
	var v252 int32
	_ = v252
	var v253 int32
	_ = v253
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v260 int32
	_ = v260
	var v261 int32
	_ = v261
	var v275 int32
	_ = v275
	var v277 int32
	_ = v277
	var v278 int32
	_ = v278
	var v284 int64
	_ = v284
	var v299 int32
	_ = v299
	var v301 int32
	_ = v301
	var v302 int32
	_ = v302
	var v305 int32
	_ = v305
	var v306 int32
	_ = v306
	var v319 int32
	_ = v319
	var v321 int32
	_ = v321
	var v322 int32
	_ = v322
	var v323 int32
	_ = v323
	var v324 int32
	_ = v324
	var v325 int32
	_ = v325
	var v327 int32
	_ = v327
	var v328 int32
	_ = v328
	var v331 int32
	_ = v331
	var v334 int32
	_ = v334
	var v341 int32
	_ = v341
	var v344 int32
	_ = v344
	var v347 int32
	_ = v347
	var v350 int32
	_ = v350
	var v351 int32
	_ = v351
	var v352 int32
	_ = v352
	var v354 int32
	_ = v354
	var v355 int32
	_ = v355
	var v356 int64
	_ = v356
	var v358 int64
	_ = v358
	var v361 int32
	_ = v361
	var v362 int32
	_ = v362
	var v368 int64
	_ = v368
	var v383 int32
	_ = v383
	var v385 int32
	_ = v385
	var v386 int32
	_ = v386
	var v389 int32
	_ = v389
	var v390 int32
	_ = v390
	var v403 int32
	_ = v403
	var v404 int32
	_ = v404
	var v405 int32
	_ = v405
	var v407 int32
	_ = v407
	var v416 int32
	_ = v416
	var v417 int32
	_ = v417
	var v428 int32
	_ = v428
	var v429 int32
	_ = v429
	var v431 int32
	_ = v431
	var v437 int32
	_ = v437
	var v438 int32
	_ = v438
	var v439 int32
	_ = v439
	var v444 int32
	_ = v444
	var v454 int32
	_ = v454
	var v457 int32
	_ = v457
	var v460 int32
	_ = v460
	var v462 int32
	_ = v462
	var v464 int32
	_ = v464
	var v465 int32
	_ = v465
	var v467 int32
	_ = v467
	var v477 int32
	_ = v477
	var v480 int32
	_ = v480
	var v485 int32
	_ = v485
	var v488 int32
	_ = v488
	var v501 int32
	_ = v501
	var v503 int32
	_ = v503
	var v505 int32
	_ = v505
	var v510 int32
	_ = v510
	var v517 int32
	_ = v517
	var v522 int32
	_ = v522
	var v532 int32
	_ = v532
	var v533 int32
	_ = v533
	var v541 int32
	_ = v541
	var v542 int32
	_ = v542
	var v544 int32
	_ = v544
	var v547 int32
	_ = v547
	var v551 int32
	_ = v551
	var v557 int32
	_ = v557
	var v558 int32
	_ = v558
	var v560 int32
	_ = v560
	var v564 int32
	_ = v564
	var v566 int32
	_ = v566
	var v572 int32
	_ = v572
	var v576 int32
	_ = v576
	var v582 int32
	_ = v582
	var v583 int32
	_ = v583
	var v593 int32
	_ = v593
	var v597 int32
	_ = v597
	var v598 int32
	_ = v598
	var v601 int32
	_ = v601
	var v605 int32
	_ = v605
	var v607 int32
	_ = v607
	var v608 int32
	_ = v608
	var v611 int32
	_ = v611
	var v612 int32
	_ = v612
	var v615 int32
	_ = v615
	var v616 int32
	_ = v616
	var v630 int32
	_ = v630
	var v633 int32
	_ = v633
	var v634 int32
	_ = v634
	var v648 int32
	_ = v648
	var v651 int32
	_ = v651
	var v652 int32
	_ = v652
	var v666 int32
	_ = v666
	var v685 int32
	_ = v685
	var v692 int32
	_ = v692
	var v698 int32
	_ = v698
	var v704 int32
	_ = v704
	v11 = m.G0
	v13 = v11 - int32(1552)
	m.G0 = v13
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v15&int32(15) != int32(6) {
		goto L4
	} else {
		goto L5
	}
L1:
	;
	F__serverAssert(m, int32(_a1524), int32(_a1525), int32(247))
	mBase = m.M
	v704 = m.ExcPending
	if v704 != 0 {
		goto L7
	} else {
		goto L125
	}
L2:
	;
	F__serverAssert(m, int32(_a1526), int32(_a1525), int32(209))
	mBase = m.M
	v698 = m.ExcPending
	if v698 != 0 {
		goto L7
	} else {
		goto L124
	}
L3:
	;
	F__serverPanic_1(m, int32(_a1525), int32(168), int32(_a1527), int32(0))
	mBase = m.M
	v692 = m.ExcPending
	if v692 != 0 {
		goto L7
	} else {
		goto L123
	}
L4:
	;
	F__serverAssert(m, int32(_a1528), int32(_a1525), int32(164))
	mBase = m.M
	v685 = m.ExcPending
	if v685 != 0 {
		goto L7
	} else {
		goto L122
	}
L5:
	;
	if v15&int32(240) != int32(160) {
		goto L3
	} else {
		goto L6
	}
L6:
	;
	v24 = F_createStreamObject(m)
	mBase = m.M
	v27 = m.ExcPending
	if v27 != 0 {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	return int32(0)
L8:
	;
	v28 = F_objectGetVal(m, l0)
	mBase = m.M
	v29 = F_objectGetVal(m, v24)
	mBase = m.M
	v31 = v13 + int32(1248)
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v28)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v31)+4)) = v32
	*(*int32)(unsafe.Add(mBase, uint32(v31))) = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v31)+20)) = int32(128)
	v38 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v31)+12)) = v38
	*(*int64)(unsafe.Add(mBase, uint32(v31)+296)) = v38
	*(*int64)(unsafe.Add(mBase, uint32(v31)+160)) = int64(137438953472)
	*(*int32)(unsafe.Add(mBase, uint32(v31)+8)) = v13 + int32(1272)
	*(*int32)(unsafe.Add(mBase, uint32(v31)+156)) = v13 + int32(1416)
	goto L9
L9:
	;
	v53 = int32(0)
	v55 = F_raxSeek(m, v13+int32(1248), int32(_a4), v53, v53)
	mBase = m.M
	v56 = m.ExcPending
	if v56 != 0 {
		goto L7
	} else {
		goto L10
	}
L10:
	;
	v59 = F_raxNext(m, v13+int32(1248))
	mBase = m.M
	v60 = m.ExcPending
	if v60 != 0 {
		goto L7
	} else {
		goto L12
	}
L11:
	;
	v109 = *(*int64)(unsafe.Add(mBase, uint32(v28)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v29)+8)) = v109
	v111 = int32(40)
	v115 = *(*int64)(unsafe.Add(mBase, uint32(v28+v111)))
	*(*int64)(unsafe.Add(mBase, uint32(v29+v111))) = v115
	v117 = *(*int64)(unsafe.Add(mBase, uint32(v28)+32))
	*(*int64)(unsafe.Add(mBase, uint32(v29)+32)) = v117
	v119 = int32(24)
	v123 = *(*int64)(unsafe.Add(mBase, uint32(v28+v119)))
	*(*int64)(unsafe.Add(mBase, uint32(v29+v119))) = v123
	v125 = *(*int64)(unsafe.Add(mBase, uint32(v28)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v29)+16)) = v125
	v127 = *(*int64)(unsafe.Add(mBase, uint32(v28)+48))
	*(*int64)(unsafe.Add(mBase, uint32(v29)+48)) = v127
	v129 = int32(56)
	v133 = *(*int64)(unsafe.Add(mBase, uint32(v28+v129)))
	*(*int64)(unsafe.Add(mBase, uint32(v29+v129))) = v133
	v135 = *(*int64)(unsafe.Add(mBase, uint32(v28)+64))
	*(*int64)(unsafe.Add(mBase, uint32(v29)+64)) = v135
	F_raxStop(m, v13+int32(1248))
	mBase = m.M
	v140 = m.ExcPending
	if v140 != 0 {
		goto L7
	} else {
		goto L24
	}
L12:
	;
	if v59 == int32(0) {
		goto L11
	} else {
		goto L13
	}
L13:
	;
	goto L14
L14:
	;
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v13)+1260))
	v74 = *(*int32)(unsafe.Add(mBase, uint32(v73)))
	goto L16
L15:
	;
	goto L11
L16:
	;
	v75 = F_valkey_malloc(m, v74)
	mBase = m.M
	v76 = m.ExcPending
	if v76 != 0 {
		goto L7
	} else {
		goto L17
	}
L17:
	;
	if v74 == int32(0) {
		v80 = v75
		goto L19
	} else {
		goto L20
	}
L18:
	;
	v81 = *(*int32)(unsafe.Add(mBase, uint32(v13)+1256))
	v82 = *(*int64)(unsafe.Add(mBase, uint32(v81)))
	*(*int64)(unsafe.Add(mBase, uint32(v13)+1232)) = v82
	v86 = *(*int64)(unsafe.Add(mBase, uint32(v81+int32(8))))
	*(*int64)(unsafe.Add(mBase, uint32(v13)+1240)) = v86
	v88 = *(*int32)(unsafe.Add(mBase, uint32(v29)+4))
	v93 = F_raxInsert(m, v88, v13+int32(1232), int32(16), v80, int32(0))
	mBase = m.M
	v94 = m.ExcPending
	if v94 != 0 {
		goto L7
	} else {
		goto L21
	}
L19:
	;
	goto L18
L20:
	;
	v79 = F__emscripten_memcpy_bulkmem(m, v75, v73, v74)
	mBase = m.M
	v80 = v79
	goto L19
L21:
	;
	v97 = F_raxNext(m, v13+int32(1248))
	mBase = m.M
	v98 = m.ExcPending
	if v98 != 0 {
		goto L7
	} else {
		goto L22
	}
L22:
	;
	if v97 != 0 {
		goto L14
	} else {
		goto L23
	}
L23:
	;
	goto L15
L24:
	;
	v141 = *(*int32)(unsafe.Add(mBase, uint32(v28)))
	if v141 == int32(0) {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	m.G0 = v13 + int32(1552)
	return v24
L26:
	;
	v145 = v13 + int32(928)
	*(*int32)(unsafe.Add(mBase, uint32(v145)+4)) = v141
	*(*int32)(unsafe.Add(mBase, uint32(v145))) = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v145)+20)) = int32(128)
	v151 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v145)+12)) = v151
	*(*int64)(unsafe.Add(mBase, uint32(v145)+296)) = v151
	*(*int64)(unsafe.Add(mBase, uint32(v145)+160)) = int64(137438953472)
	*(*int32)(unsafe.Add(mBase, uint32(v145)+8)) = v13 + int32(952)
	*(*int32)(unsafe.Add(mBase, uint32(v145)+156)) = v13 + int32(1096)
	goto L27
L27:
	;
	v166 = int32(0)
	v168 = F_raxSeek(m, v13+int32(928), int32(_a4), v166, v166)
	mBase = m.M
	v169 = m.ExcPending
	if v169 != 0 {
		goto L7
	} else {
		goto L28
	}
L28:
	;
	v172 = F_raxNext(m, v13+int32(928))
	mBase = m.M
	v173 = m.ExcPending
	if v173 != 0 {
		goto L7
	} else {
		goto L30
	}
L29:
	;
	F_raxStop(m, v13+int32(928))
	mBase = m.M
	v666 = m.ExcPending
	if v666 != 0 {
		goto L7
	} else {
		goto L121
	}
L30:
	;
	if v172 == int32(0) {
		goto L29
	} else {
		goto L31
	}
L31:
	;
	goto L32
L32:
	;
	v186 = *(*int32)(unsafe.Add(mBase, uint32(v13)+936))
	v187 = *(*int32)(unsafe.Add(mBase, uint32(v13)+944))
	v188 = *(*int32)(unsafe.Add(mBase, uint32(v13)+940))
	v189 = *(*int64)(unsafe.Add(mBase, uint32(v188)+16))
	v190 = F_streamCreateCG(m, v29, v186, v187, v188, v189)
	mBase = m.M
	v191 = m.ExcPending
	if v191 != 0 {
		goto L7
	} else {
		goto L34
	}
L33:
	;
	goto L29
L34:
	;
	if v190 == int32(0) {
		goto L2
	} else {
		goto L35
	}
L35:
	;
	v195 = v13 + int32(624)
	v196 = *(*int32)(unsafe.Add(mBase, uint32(v188)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v195)+4)) = v196
	*(*int32)(unsafe.Add(mBase, uint32(v195))) = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v195)+20)) = int32(128)
	v202 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v195)+12)) = v202
	*(*int64)(unsafe.Add(mBase, uint32(v195)+296)) = v202
	*(*int64)(unsafe.Add(mBase, uint32(v195)+160)) = int64(137438953472)
	*(*int32)(unsafe.Add(mBase, uint32(v195)+8)) = v13 + int32(648)
	*(*int32)(unsafe.Add(mBase, uint32(v195)+156)) = v13 + int32(792)
	goto L36
L36:
	;
	v217 = int32(0)
	v219 = F_raxSeek(m, v13+int32(624), int32(_a4), v217, v217)
	mBase = m.M
	v220 = m.ExcPending
	if v220 != 0 {
		goto L7
	} else {
		goto L37
	}
L37:
	;
	v223 = F_raxNext(m, v13+int32(624))
	mBase = m.M
	v224 = m.ExcPending
	if v224 != 0 {
		goto L7
	} else {
		goto L39
	}
L38:
	;
	F_raxStop(m, v13+int32(624))
	mBase = m.M
	v275 = m.ExcPending
	if v275 != 0 {
		goto L7
	} else {
		goto L48
	}
L39:
	;
	if v223 == int32(0) {
		goto L38
	} else {
		goto L40
	}
L40:
	;
	goto L41
L41:
	;
	v237 = *(*int32)(unsafe.Add(mBase, uint32(v13)+636))
	v239 = F_valkey_malloc(m, int32(24))
	mBase = m.M
	v240 = m.ExcPending
	if v240 != 0 {
		goto L7
	} else {
		goto L43
	}
L42:
	;
	goto L38
L43:
	;
	v242 = *(*int64)(unsafe.Add(mBase, _consts[78]))
	goto L44
L44:
	;
	v243 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v239)+16)) = v243
	*(*int64)(unsafe.Add(mBase, uint32(v239)+8)) = int64(1)
	*(*int64)(unsafe.Add(mBase, uint32(v239))) = v242
	v248 = *(*int64)(unsafe.Add(mBase, uint32(v237)))
	*(*int64)(unsafe.Add(mBase, uint32(v239))) = v248
	v250 = *(*int64)(unsafe.Add(mBase, uint32(v237)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v239)+8)) = v250
	v252 = *(*int32)(unsafe.Add(mBase, uint32(v190)+24))
	v253 = *(*int32)(unsafe.Add(mBase, uint32(v13)+632))
	v256 = F_raxInsert(m, v252, v253, int32(16), v239, v243)
	mBase = m.M
	v257 = m.ExcPending
	if v257 != 0 {
		goto L7
	} else {
		goto L45
	}
L45:
	;
	v260 = F_raxNext(m, v13+int32(624))
	mBase = m.M
	v261 = m.ExcPending
	if v261 != 0 {
		goto L7
	} else {
		goto L46
	}
L46:
	;
	if v260 != 0 {
		goto L41
	} else {
		goto L47
	}
L47:
	;
	goto L42
L48:
	;
	v277 = v13 + int32(320)
	v278 = *(*int32)(unsafe.Add(mBase, uint32(v188)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v277)+4)) = v278
	*(*int32)(unsafe.Add(mBase, uint32(v277))) = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v277)+20)) = int32(128)
	v284 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v277)+12)) = v284
	*(*int64)(unsafe.Add(mBase, uint32(v277)+296)) = v284
	*(*int64)(unsafe.Add(mBase, uint32(v277)+160)) = int64(137438953472)
	*(*int32)(unsafe.Add(mBase, uint32(v277)+8)) = v13 + int32(344)
	*(*int32)(unsafe.Add(mBase, uint32(v277)+156)) = v13 + int32(488)
	goto L49
L49:
	;
	v299 = int32(0)
	v301 = F_raxSeek(m, v13+int32(320), int32(_a4), v299, v299)
	mBase = m.M
	v302 = m.ExcPending
	if v302 != 0 {
		goto L7
	} else {
		goto L50
	}
L50:
	;
	v305 = F_raxNext(m, v13+int32(320))
	mBase = m.M
	v306 = m.ExcPending
	if v306 != 0 {
		goto L7
	} else {
		goto L52
	}
L51:
	;
	F_raxStop(m, v13+int32(320))
	mBase = m.M
	v648 = m.ExcPending
	if v648 != 0 {
		goto L7
	} else {
		goto L118
	}
L52:
	;
	if v305 == int32(0) {
		goto L51
	} else {
		goto L53
	}
L53:
	;
	goto L54
L54:
	;
	v319 = *(*int32)(unsafe.Add(mBase, uint32(v13)+332))
	v321 = F_valkey_malloc(m, int32(24))
	mBase = m.M
	v322 = m.ExcPending
	if v322 != 0 {
		goto L7
	} else {
		goto L56
	}
L55:
	;
	goto L51
L56:
	;
	v323 = *(*int32)(unsafe.Add(mBase, uint32(v319)+16))
	v324 = F_sdsdup(m, v323)
	mBase = m.M
	v325 = m.ExcPending
	if v325 != 0 {
		goto L7
	} else {
		goto L57
	}
L57:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v321)+16)) = v324
	v327 = F_raxNew(m)
	mBase = m.M
	v328 = m.ExcPending
	if v328 != 0 {
		goto L7
	} else {
		goto L58
	}
L58:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v321)+20)) = v327
	v331 = *(*int32)(unsafe.Add(mBase, uint32(v321)+16))
	v334 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v331+int32(-1)))))
	switch v334 & int32(7) {
	case 0:
		goto L64
	case 1:
		goto L63
	case 2:
		goto L62
	case 3:
		goto L61
	case 4:
		goto L60
	default:
		v351 = int32(0)
		goto L59
	}
L59:
	;
	v352 = *(*int32)(unsafe.Add(mBase, uint32(v190)+28))
	v354 = F_raxInsert(m, v352, v331, v351, v321, int32(0))
	mBase = m.M
	v355 = m.ExcPending
	if v355 != 0 {
		goto L7
	} else {
		goto L65
	}
L60:
	;
	v350 = *(*int32)(unsafe.Add(mBase, uint32(v331+int32(-17))))
	v351 = v350
	goto L59
L61:
	;
	v347 = *(*int32)(unsafe.Add(mBase, uint32(v331+int32(-9))))
	v351 = v347
	goto L59
L62:
	;
	v344 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v331+int32(-5)))))
	v351 = v344
	goto L59
L63:
	;
	v341 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v331+int32(-3)))))
	v351 = v341
	goto L59
L64:
	;
	v351 = int32(base.Ui32(v334) >> (uint(int32(3)) % 32))
	goto L59
L65:
	;
	v356 = *(*int64)(unsafe.Add(mBase, uint32(v319)))
	*(*int64)(unsafe.Add(mBase, uint32(v321))) = v356
	v358 = *(*int64)(unsafe.Add(mBase, uint32(v319)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v321)+8)) = v358
	v361 = v13 + int32(16)
	v362 = *(*int32)(unsafe.Add(mBase, uint32(v319)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v361)+4)) = v362
	*(*int32)(unsafe.Add(mBase, uint32(v361))) = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v361)+20)) = int32(128)
	v368 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v361)+12)) = v368
	*(*int64)(unsafe.Add(mBase, uint32(v361)+296)) = v368
	*(*int64)(unsafe.Add(mBase, uint32(v361)+160)) = int64(137438953472)
	*(*int32)(unsafe.Add(mBase, uint32(v361)+8)) = v13 + int32(40)
	*(*int32)(unsafe.Add(mBase, uint32(v361)+156)) = v13 + int32(184)
	goto L66
L66:
	;
	v383 = int32(0)
	v385 = F_raxSeek(m, v13+int32(16), int32(_a4), v383, v383)
	mBase = m.M
	v386 = m.ExcPending
	if v386 != 0 {
		goto L7
	} else {
		goto L67
	}
L67:
	;
	v389 = F_raxNext(m, v13+int32(16))
	mBase = m.M
	v390 = m.ExcPending
	if v390 != 0 {
		goto L7
	} else {
		goto L69
	}
L68:
	;
	F_raxStop(m, v13+int32(16))
	mBase = m.M
	v630 = m.ExcPending
	if v630 != 0 {
		goto L7
	} else {
		goto L115
	}
L69:
	;
	if v389 == int32(0) {
		goto L68
	} else {
		goto L70
	}
L70:
	;
	goto L71
L71:
	;
	v403 = *(*int32)(unsafe.Add(mBase, uint32(v190)+24))
	v404 = *(*int32)(unsafe.Add(mBase, uint32(v13)+24))
	v405 = int32(16)
	v407 = v13 + int32(12)
	v416 = *(*int32)(unsafe.Add(mBase, uint32(v403)))
	v417 = *(*int32)(unsafe.Add(mBase, uint32(v416)))
	goto L76
L72:
	;
	goto L68
L73:
	;
	if v601 == int32(0) {
		goto L1
	} else {
		goto L111
	}
L74:
	;
	if v560 != v405 {
		v601 = int32(0)
		goto L101
	} else {
		goto L102
	}
L75:
	;
	v551 = int32(0)
	v557 = v416
	v558 = v417
	v560 = v551
	v564 = v551
	goto L74
L76:
	;
	if base.Ui32(v417) < base.Ui32(int32(8)) {
		goto L75
	} else {
		goto L77
	}
L77:
	;
	v428 = v416
	v429 = v417
	v431 = int32(0)
	goto L79
L78:
	;
	v557 = v541
	v558 = v542
	v560 = v544
	v564 = base.B2i32(v547 != int32(0))
	goto L74
L79:
	;
	v437 = int32(base.Ui32(v429) >> (uint(int32(3)) % 32))
	v438 = int32(4)
	v439 = v428 + v438
	if v429&v438 == int32(0) {
		goto L82
	} else {
		goto L83
	}
L80:
	;
	v541 = v532
	v542 = v533
	v544 = v517
	v547 = v522
	goto L78
L81:
	;
	v522 = int32(0)
	v532 = *(*int32)(unsafe.Add(mBase, uint32(v439+v437+(v522-v437)&int32(3)+v510<<(uint(int32(2))%32))))
	v533 = *(*int32)(unsafe.Add(mBase, uint32(v532)))
	if base.Ui32(v533) < base.Ui32(int32(8)) {
		v541 = v532
		v542 = v533
		v544 = v517
		v547 = v522
		goto L78
	} else {
		goto L99
	}
L82:
	;
	v485 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v404+v431))))
	v488 = int32(0)
	goto L93
L83:
	;
	v444 = int32(0)
	if base.Ui32(v405) <= base.Ui32(v431) {
		v477 = v431
		v480 = v444
		goto L84
	} else {
		goto L85
	}
L84:
	;
	if v480 == v437 {
		v510 = v444
		v517 = v477
		goto L81
	} else {
		goto L91
	}
L85:
	;
	v454 = v431
	v457 = v444
	goto L86
L86:
	;
	v460 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v439+v457))))
	v462 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v404+v454))))
	if v460 != v462 {
		v477 = v454
		v480 = v457
		goto L84
	} else {
		goto L88
	}
L87:
	;
	v477 = v465
	v480 = v467
	goto L84
L88:
	;
	v464 = int32(1)
	v465 = v454 + v464
	v467 = v457 + v464
	if base.Ui32(v437) <= base.Ui32(v467) {
		v477 = v465
		v480 = v467
		goto L84
	} else {
		goto L89
	}
L89:
	;
	if base.Ui32(v465) < base.Ui32(v405) {
		v454 = v465
		v457 = v467
		goto L86
	} else {
		goto L90
	}
L90:
	;
	goto L87
L91:
	;
	v541 = v428
	v542 = v429
	v544 = v477
	v547 = v480
	goto L78
L92:
	;
	if v488 != v437 {
		goto L97
	} else {
		goto L98
	}
L93:
	;
	v501 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v439+v488))))
	if v501 == v485&int32(255) {
		goto L92
	} else {
		goto L95
	}
L95:
	;
	v503 = int32(1)
	v505 = v488 + v503
	if v505 != v437 {
		v488 = v505
		goto L93
	} else {
		goto L96
	}
L96:
	;
	v557 = v428
	v558 = v429
	v560 = v431
	v564 = v503
	goto L74
L97:
	;
	v510 = v488
	v517 = v431 + int32(1)
	goto L81
L98:
	;
	v541 = v428
	v542 = v429
	v544 = v431
	v547 = v437
	goto L78
L99:
	;
	if base.Ui32(v517) < base.Ui32(v405) {
		v428 = v532
		v429 = v533
		v431 = v517
		goto L79
	} else {
		goto L100
	}
L100:
	;
	goto L80
L101:
	;
	goto L73
L102:
	;
	v566 = int32(0)
	if v558&int32(1) == v566 {
		v601 = v566
		goto L101
	} else {
		goto L103
	}
L103:
	;
	v572 = v558 & int32(4)
	if v564&base.B2i32(v572 != int32(0)) != 0 {
		v601 = v566
		goto L101
	} else {
		goto L104
	}
L104:
	;
	v576 = int32(1)
	if v407 == int32(0) {
		v601 = v576
		goto L101
	} else {
		goto L105
	}
L105:
	;
	if v558&int32(2) != 0 {
		v598 = int32(0)
		goto L106
	} else {
		goto L107
	}
L106:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v407))) = v598
	v601 = v576
	goto L101
L107:
	;
	v582 = int32(3)
	v583 = int32(base.Ui32(v558) >> (uint(v582) % 32))
	if v572 != 0 {
		goto L108
	} else {
		goto L109
	}
L108:
	;
	v593 = int32(4)
	goto L110
L109:
	;
	v593 = v583 << (uint(int32(2)) % 32)
	goto L110
L110:
	;
	v597 = *(*int32)(unsafe.Add(mBase, uint32(v557+v583+(int32(0)-v583)&v582+v593+int32(4))))
	v598 = v597
	goto L106
L111:
	;
	v605 = *(*int32)(unsafe.Add(mBase, uint32(v13)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v605)+16)) = v321
	v607 = *(*int32)(unsafe.Add(mBase, uint32(v321)+20))
	v608 = *(*int32)(unsafe.Add(mBase, uint32(v13)+24))
	v611 = F_raxInsert(m, v607, v608, int32(16), v605, int32(0))
	mBase = m.M
	v612 = m.ExcPending
	if v612 != 0 {
		goto L7
	} else {
		goto L112
	}
L112:
	;
	v615 = F_raxNext(m, v13+int32(16))
	mBase = m.M
	v616 = m.ExcPending
	if v616 != 0 {
		goto L7
	} else {
		goto L113
	}
L113:
	;
	if v615 != 0 {
		goto L71
	} else {
		goto L114
	}
L114:
	;
	goto L72
L115:
	;
	v633 = F_raxNext(m, v13+int32(320))
	mBase = m.M
	v634 = m.ExcPending
	if v634 != 0 {
		goto L7
	} else {
		goto L116
	}
L116:
	;
	if v633 != 0 {
		goto L54
	} else {
		goto L117
	}
L117:
	;
	goto L55
L118:
	;
	v651 = F_raxNext(m, v13+int32(928))
	mBase = m.M
	v652 = m.ExcPending
	if v652 != 0 {
		goto L7
	} else {
		goto L119
	}
L119:
	;
	if v651 != 0 {
		goto L32
	} else {
		goto L120
	}
L120:
	;
	goto L33
L121:
	;
	goto L25
L122:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L123:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L124:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L125:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_streamFreeCGVoid(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	F_raxFreeWithCallback(m, v2, int32(102))
	mBase = m.M
	v5 = m.ExcPending
	if v5 != 0 {
		return
	} else {
		v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		F_raxFreeWithCallback(m, v6, int32(1094))
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
func F_streamFreeConsumerVoid(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	F_raxFree(m, v2)
	mBase = m.M
	v4 = m.ExcPending
	if v4 != 0 {
		return
	} else {
		v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
		F_sdsfree(m, v5)
		mBase = m.M
		v7 = m.ExcPending
		if v7 != 0 {
			return
		} else {
			F_valkey_free(m, l0)
			mBase = m.M
			v9 = m.ExcPending
			if v9 != 0 {
				return
			} else {
				return
			}
		}
	}
}
func F_streamIncrID(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int64
	_ = v3
	var v6 int64
	_ = v6
	var v9 int64
	_ = v9
	v3 = *(*int64)(unsafe.Add(mBase, uint32(l0)+8))
	if v3 != int64(-1) {
		*(*int64)(unsafe.Add(mBase, uint32(l0)+8)) = v3 + int64(1)
		return int32(0)
	} else {
		v6 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
		if v6 != int64(-1) {
			*(*int64)(unsafe.Add(mBase, uint32(l0)+8)) = int64(0)
			*(*int64)(unsafe.Add(mBase, uint32(l0))) = v6 + int64(1)
			return int32(0)
		} else {
			v9 = int64(0)
			*(*int64)(unsafe.Add(mBase, uint32(l0))) = v9
			*(*int64)(unsafe.Add(mBase, uint32(l0+int32(8)))) = v9
			return int32(-1)
		}
	}
}
func F_streamIteratorGetField(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
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
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	v8 = l0 + int32(404)
	v9 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+40)))
	if v9&int32(2) == int32(0) {
		v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+396))
		v25 = F_lpGet(m, v24, l3, v8)
		mBase = m.M
		v26 = m.ExcPending
		if v26 != 0 {
			return
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(l1))) = v25
			v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+392))
			v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)+396))
			v30 = F_lpNext(m, v28, v29)
			mBase = m.M
			v31 = m.ExcPending
			if v31 != 0 {
				return
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(l0)+396)) = v30
				v33 = v30
				v36 = F_lpGet(m, v33, l4, l0+int32(425))
				mBase = m.M
				v37 = m.ExcPending
				if v37 != 0 {
					return
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(l2))) = v36
					v39 = *(*int32)(unsafe.Add(mBase, uint32(l0)+392))
					v40 = *(*int32)(unsafe.Add(mBase, uint32(l0)+396))
					v41 = F_lpNext(m, v39, v40)
					mBase = m.M
					v42 = m.ExcPending
					if v42 != 0 {
						return
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(l0)+396)) = v41
						return
					}
				}
			}
		}
	} else {
		v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
		v15 = F_lpGet(m, v14, l3, v8)
		mBase = m.M
		v16 = m.ExcPending
		if v16 != 0 {
			return
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(l1))) = v15
			v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+392))
			v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
			v20 = F_lpNext(m, v18, v19)
			mBase = m.M
			v21 = m.ExcPending
			if v21 != 0 {
				return
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = v20
				v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+396))
				v33 = v23
				v36 = F_lpGet(m, v33, l4, l0+int32(425))
				mBase = m.M
				v37 = m.ExcPending
				if v37 != 0 {
					return
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(l2))) = v36
					v39 = *(*int32)(unsafe.Add(mBase, uint32(l0)+392))
					v40 = *(*int32)(unsafe.Add(mBase, uint32(l0)+396))
					v41 = F_lpNext(m, v39, v40)
					mBase = m.M
					v42 = m.ExcPending
					if v42 != 0 {
						return
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(l0)+396)) = v41
						return
					}
				}
			}
		}
	}
}
func F_streamIteratorRemoveEntry(m *base.Module, l0 int32, l1 int32) {
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
	var v22 int32
	_ = v22
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v70 int64
	_ = v70
	var v76 int32
	_ = v76
	var v78 int64
	_ = v78
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v92 int64
	_ = v92
	var v97 int64
	_ = v97
	var v101 int32
	_ = v101
	var v103 int64
	_ = v103
	var v105 int32
	_ = v105
	var v113 int64
	_ = v113
	var v137 int64
	_ = v137
	var v156 int32
	_ = v156
	var v165 int64
	_ = v165
	var v166 int64
	_ = v166
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v181 int64
	_ = v181
	var v182 int32
	_ = v182
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v202 int32
	_ = v202
	var v218 int32
	_ = v218
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v233 int64
	_ = v233
	var v239 int32
	_ = v239
	var v241 int64
	_ = v241
	var v243 int32
	_ = v243
	var v245 int32
	_ = v245
	var v255 int64
	_ = v255
	var v260 int64
	_ = v260
	var v264 int32
	_ = v264
	var v266 int64
	_ = v266
	var v268 int32
	_ = v268
	var v276 int64
	_ = v276
	var v300 int64
	_ = v300
	var v319 int32
	_ = v319
	var v328 int64
	_ = v328
	var v329 int64
	_ = v329
	var v333 int32
	_ = v333
	var v334 int32
	_ = v334
	var v335 int32
	_ = v335
	var v336 int32
	_ = v336
	var v337 int32
	_ = v337
	var v339 int32
	_ = v339
	var v340 int32
	_ = v340
	var v345 int32
	_ = v345
	var v346 int32
	_ = v346
	var v347 int32
	_ = v347
	var v348 int32
	_ = v348
	var v349 int32
	_ = v349
	var v354 int32
	_ = v354
	var v355 int32
	_ = v355
	var v356 int64
	_ = v356
	var v357 int32
	_ = v357
	var v359 int32
	_ = v359
	var v360 int32
	_ = v360
	var v372 int32
	_ = v372
	var v373 int32
	_ = v373
	var v377 int32
	_ = v377
	var v393 int32
	_ = v393
	var v395 int32
	_ = v395
	var v396 int32
	_ = v396
	var v397 int32
	_ = v397
	var v408 int64
	_ = v408
	var v414 int32
	_ = v414
	var v416 int64
	_ = v416
	var v418 int32
	_ = v418
	var v420 int32
	_ = v420
	var v430 int64
	_ = v430
	var v435 int64
	_ = v435
	var v439 int32
	_ = v439
	var v441 int64
	_ = v441
	var v443 int32
	_ = v443
	var v451 int64
	_ = v451
	var v475 int64
	_ = v475
	var v494 int32
	_ = v494
	var v503 int64
	_ = v503
	var v504 int64
	_ = v504
	var v509 int32
	_ = v509
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
	var v516 int32
	_ = v516
	var v518 int32
	_ = v518
	var v519 int32
	_ = v519
	var v523 int32
	_ = v523
	var v524 int64
	_ = v524
	var v528 int32
	_ = v528
	var v535 int64
	_ = v535
	var v537 int64
	_ = v537
	var v545 int64
	_ = v545
	var v547 int64
	_ = v547
	var v551 int32
	_ = v551
	var v558 int64
	_ = v558
	var v560 int64
	_ = v560
	var v565 int32
	_ = v565
	var v566 int32
	_ = v566
	var v571 int32
	_ = v571
	var v573 int32
	_ = v573
	var v581 int32
	_ = v581
	var v587 int32
	_ = v587
	var v593 int32
	_ = v593
	v7 = m.G0
	v9 = v7 - int32(48)
	m.G0 = v9
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+392))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+400))
	v16 = F_lpGet(m, v12, v9+int32(24), int32(0))
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		goto L6
	} else {
		goto L7
	}
L1:
	;
	F__serverAssert(m, int32(_a1529), int32(_a1525), int32(282))
	mBase = m.M
	v593 = m.ExcPending
	if v593 != 0 {
		goto L6
	} else {
		goto L120
	}
L2:
	;
	F__serverAssert(m, int32(_a1529), int32(_a1525), int32(282))
	mBase = m.M
	v587 = m.ExcPending
	if v587 != 0 {
		goto L6
	} else {
		goto L119
	}
L3:
	;
	F__serverAssert(m, int32(_a1529), int32(_a1525), int32(282))
	mBase = m.M
	v581 = m.ExcPending
	if v581 != 0 {
		goto L6
	} else {
		goto L118
	}
L4:
	;
	v171 = F_lpReplaceInteger(m, v11, l0+int32(400), v166|int64(1))
	mBase = m.M
	v172 = m.ExcPending
	if v172 != 0 {
		goto L6
	} else {
		goto L37
	}
L5:
	;
	v19 = *(*int32)(unsafe.Add(mBase, uint32(v9)+24))
	v21 = v9 + int32(8)
	v22 = int32(0)
	if base.Ui32(v19+int32(-21)) < base.Ui32(int32(-20)) {
		v156 = v22
		goto L10
	} else {
		goto L11
	}
L6:
	;
	return
L7:
	;
	if v16 != 0 {
		goto L5
	} else {
		goto L8
	}
L8:
	;
	v18 = *(*int64)(unsafe.Add(mBase, uint32(v9)+24))
	v166 = v18
	goto L4
L9:
	;
	if v156 == int32(0) {
		goto L3
	} else {
		goto L36
	}
L10:
	;
	goto L9
L11:
	;
	v34 = int32(1)
	v35 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16))))
	if v19 != v34 {
		goto L14
	} else {
		goto L15
	}
L12:
	;
	v156 = int32(1)
	goto L10
L13:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v21))) = v137
	goto L12
L14:
	;
	if v35&int32(255) == int32(45) {
		goto L19
	} else {
		goto L20
	}
L15:
	;
	v39 = v35 + int32(-48)
	if base.Ui32(int32(9)) < base.Ui32(v39&int32(255)) {
		v156 = v22
		goto L10
	} else {
		goto L16
	}
L16:
	;
	if v21 == int32(0) {
		goto L12
	} else {
		goto L17
	}
L17:
	;
	v137 = base.I64_extend_i32_u(v39) & int64(255)
	goto L13
L18:
	;
	if base.Ui32(int32(8)) < base.Ui32((v58+int32(-49))&int32(255)) {
		v156 = v22
		goto L10
	} else {
		goto L21
	}
L19:
	;
	v55 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16)+1)))
	v57 = int32(2)
	v58 = v55
	v59 = v16 + int32(1)
	goto L18
L20:
	;
	v57 = v34
	v58 = v35
	v59 = v16
	goto L18
L21:
	;
	v70 = base.I64_extend_i32_u(v58+int32(-48)) & int64(255)
	if base.Ui32(v19) <= base.Ui32(v57) {
		v113 = v70
		goto L22
	} else {
		goto L23
	}
L22:
	;
	if v35&int32(255) != int32(45) {
		goto L30
	} else {
		goto L31
	}
L23:
	;
	v76 = v57
	v78 = v70
	v80 = v59
	goto L24
L24:
	;
	v82 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v80)+1)))
	if base.Ui32((v82+int32(-58))&int32(255)) < base.Ui32(int32(246)) {
		v156 = v22
		goto L10
	} else {
		goto L26
	}
L25:
	;
	v113 = v103
	goto L22
L26:
	;
	if base.Ui64(int64(1844674407370955161)) < base.Ui64(v78) {
		v156 = v22
		goto L10
	} else {
		goto L27
	}
L27:
	;
	v92 = v78 * int64(10)
	v97 = base.I64_extend_i32_u(v82+int32(-48)) & int64(255)
	if base.Ui64(v97^int64(-1)) < base.Ui64(v92) {
		v156 = v22
		goto L10
	} else {
		goto L28
	}
L28:
	;
	v101 = int32(1)
	v103 = v92 + v97
	v105 = v76 + v101
	if v105 != v19 {
		v76 = v105
		v78 = v103
		v80 = v80 + v101
		goto L24
	} else {
		goto L29
	}
L29:
	;
	goto L25
L30:
	;
	if v113 < int64(0) {
		v156 = v22
		goto L10
	} else {
		goto L34
	}
L31:
	;
	if base.Ui64(int64(-9223372036854775807-1)) < base.Ui64(v113) {
		v156 = v22
		goto L10
	} else {
		goto L32
	}
L32:
	;
	if v21 == int32(0) {
		goto L12
	} else {
		goto L33
	}
L33:
	;
	v137 = int64(0) - v113
	goto L13
L34:
	;
	if v21 == int32(0) {
		goto L12
	} else {
		goto L35
	}
L35:
	;
	v137 = v113
	goto L13
L36:
	;
	v165 = *(*int64)(unsafe.Add(mBase, uint32(v9)+8))
	v166 = v165
	goto L4
L37:
	;
	v173 = F_lpFirst(m, v171)
	mBase = m.M
	v174 = m.ExcPending
	if v174 != 0 {
		goto L6
	} else {
		goto L38
	}
L38:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9)+44)) = v173
	v179 = F_lpGet(m, v173, v9+int32(24), int32(0))
	mBase = m.M
	v180 = m.ExcPending
	if v180 != 0 {
		goto L6
	} else {
		goto L41
	}
L39:
	;
	if v329 != int64(1) {
		goto L72
	} else {
		goto L73
	}
L40:
	;
	v182 = *(*int32)(unsafe.Add(mBase, uint32(v9)+24))
	v184 = v9 + int32(8)
	v185 = int32(0)
	if base.Ui32(v182+int32(-21)) < base.Ui32(int32(-20)) {
		v319 = v185
		goto L44
	} else {
		goto L45
	}
L41:
	;
	if v179 != 0 {
		goto L40
	} else {
		goto L42
	}
L42:
	;
	v181 = *(*int64)(unsafe.Add(mBase, uint32(v9)+24))
	v329 = v181
	goto L39
L43:
	;
	if v319 == int32(0) {
		goto L2
	} else {
		goto L70
	}
L44:
	;
	goto L43
L45:
	;
	v197 = int32(1)
	v198 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v179))))
	if v182 != v197 {
		goto L48
	} else {
		goto L49
	}
L46:
	;
	v319 = int32(1)
	goto L44
L47:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v184))) = v300
	goto L46
L48:
	;
	if v198&int32(255) == int32(45) {
		goto L53
	} else {
		goto L54
	}
L49:
	;
	v202 = v198 + int32(-48)
	if base.Ui32(int32(9)) < base.Ui32(v202&int32(255)) {
		v319 = v185
		goto L44
	} else {
		goto L50
	}
L50:
	;
	if v184 == int32(0) {
		goto L46
	} else {
		goto L51
	}
L51:
	;
	v300 = base.I64_extend_i32_u(v202) & int64(255)
	goto L47
L52:
	;
	if base.Ui32(int32(8)) < base.Ui32((v221+int32(-49))&int32(255)) {
		v319 = v185
		goto L44
	} else {
		goto L55
	}
L53:
	;
	v218 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v179)+1)))
	v220 = int32(2)
	v221 = v218
	v222 = v179 + int32(1)
	goto L52
L54:
	;
	v220 = v197
	v221 = v198
	v222 = v179
	goto L52
L55:
	;
	v233 = base.I64_extend_i32_u(v221+int32(-48)) & int64(255)
	if base.Ui32(v182) <= base.Ui32(v220) {
		v276 = v233
		goto L56
	} else {
		goto L57
	}
L56:
	;
	if v198&int32(255) != int32(45) {
		goto L64
	} else {
		goto L65
	}
L57:
	;
	v239 = v220
	v241 = v233
	v243 = v222
	goto L58
L58:
	;
	v245 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v243)+1)))
	if base.Ui32((v245+int32(-58))&int32(255)) < base.Ui32(int32(246)) {
		v319 = v185
		goto L44
	} else {
		goto L60
	}
L59:
	;
	v276 = v266
	goto L56
L60:
	;
	if base.Ui64(int64(1844674407370955161)) < base.Ui64(v241) {
		v319 = v185
		goto L44
	} else {
		goto L61
	}
L61:
	;
	v255 = v241 * int64(10)
	v260 = base.I64_extend_i32_u(v245+int32(-48)) & int64(255)
	if base.Ui64(v260^int64(-1)) < base.Ui64(v255) {
		v319 = v185
		goto L44
	} else {
		goto L62
	}
L62:
	;
	v264 = int32(1)
	v266 = v255 + v260
	v268 = v239 + v264
	if v268 != v182 {
		v239 = v268
		v241 = v266
		v243 = v243 + v264
		goto L58
	} else {
		goto L63
	}
L63:
	;
	goto L59
L64:
	;
	if v276 < int64(0) {
		v319 = v185
		goto L44
	} else {
		goto L68
	}
L65:
	;
	if base.Ui64(int64(-9223372036854775807-1)) < base.Ui64(v276) {
		v319 = v185
		goto L44
	} else {
		goto L66
	}
L66:
	;
	if v184 == int32(0) {
		goto L46
	} else {
		goto L67
	}
L67:
	;
	v300 = int64(0) - v276
	goto L47
L68:
	;
	if v184 == int32(0) {
		goto L46
	} else {
		goto L69
	}
L69:
	;
	v300 = v276
	goto L47
L70:
	;
	v328 = *(*int64)(unsafe.Add(mBase, uint32(v9)+8))
	v329 = v328
	goto L39
L71:
	;
	v523 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v524 = *(*int64)(unsafe.Add(mBase, uint32(v523)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v523)+8)) = v524 + int64(-1)
	v528 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	if v528 == int32(0) {
		goto L114
	} else {
		goto L115
	}
L72:
	;
	v345 = F_lpReplaceInteger(m, v171, v9+int32(44), v329+int64(-1))
	mBase = m.M
	v346 = m.ExcPending
	if v346 != 0 {
		goto L6
	} else {
		goto L76
	}
L73:
	;
	F_lpFree(m, v171)
	mBase = m.M
	v333 = m.ExcPending
	if v333 != 0 {
		goto L6
	} else {
		goto L74
	}
L74:
	;
	v334 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v335 = *(*int32)(unsafe.Add(mBase, uint32(v334)+4))
	v336 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v337 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
	v339 = F_raxRemove(m, v335, v336, v337, int32(0))
	mBase = m.M
	v340 = m.ExcPending
	if v340 != 0 {
		goto L6
	} else {
		goto L75
	}
L75:
	;
	goto L71
L76:
	;
	v347 = *(*int32)(unsafe.Add(mBase, uint32(v9)+44))
	v348 = F_lpNext(m, v345, v347)
	mBase = m.M
	v349 = m.ExcPending
	if v349 != 0 {
		goto L6
	} else {
		goto L77
	}
L77:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9)+44)) = v348
	v354 = F_lpGet(m, v348, v9+int32(24), int32(0))
	mBase = m.M
	v355 = m.ExcPending
	if v355 != 0 {
		goto L6
	} else {
		goto L80
	}
L78:
	;
	v509 = F_lpReplaceInteger(m, v345, v9+int32(44), v504+int64(1))
	mBase = m.M
	v510 = m.ExcPending
	if v510 != 0 {
		goto L6
	} else {
		goto L110
	}
L79:
	;
	v357 = *(*int32)(unsafe.Add(mBase, uint32(v9)+24))
	v359 = v9 + int32(8)
	v360 = int32(0)
	if base.Ui32(v357+int32(-21)) < base.Ui32(int32(-20)) {
		v494 = v360
		goto L83
	} else {
		goto L84
	}
L80:
	;
	if v354 != 0 {
		goto L79
	} else {
		goto L81
	}
L81:
	;
	v356 = *(*int64)(unsafe.Add(mBase, uint32(v9)+24))
	v504 = v356
	goto L78
L82:
	;
	if v494 == int32(0) {
		goto L1
	} else {
		goto L109
	}
L83:
	;
	goto L82
L84:
	;
	v372 = int32(1)
	v373 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v354))))
	if v357 != v372 {
		goto L87
	} else {
		goto L88
	}
L85:
	;
	v494 = int32(1)
	goto L83
L86:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v359))) = v475
	goto L85
L87:
	;
	if v373&int32(255) == int32(45) {
		goto L92
	} else {
		goto L93
	}
L88:
	;
	v377 = v373 + int32(-48)
	if base.Ui32(int32(9)) < base.Ui32(v377&int32(255)) {
		v494 = v360
		goto L83
	} else {
		goto L89
	}
L89:
	;
	if v359 == int32(0) {
		goto L85
	} else {
		goto L90
	}
L90:
	;
	v475 = base.I64_extend_i32_u(v377) & int64(255)
	goto L86
L91:
	;
	if base.Ui32(int32(8)) < base.Ui32((v396+int32(-49))&int32(255)) {
		v494 = v360
		goto L83
	} else {
		goto L94
	}
L92:
	;
	v393 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v354)+1)))
	v395 = int32(2)
	v396 = v393
	v397 = v354 + int32(1)
	goto L91
L93:
	;
	v395 = v372
	v396 = v373
	v397 = v354
	goto L91
L94:
	;
	v408 = base.I64_extend_i32_u(v396+int32(-48)) & int64(255)
	if base.Ui32(v357) <= base.Ui32(v395) {
		v451 = v408
		goto L95
	} else {
		goto L96
	}
L95:
	;
	if v373&int32(255) != int32(45) {
		goto L103
	} else {
		goto L104
	}
L96:
	;
	v414 = v395
	v416 = v408
	v418 = v397
	goto L97
L97:
	;
	v420 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v418)+1)))
	if base.Ui32((v420+int32(-58))&int32(255)) < base.Ui32(int32(246)) {
		v494 = v360
		goto L83
	} else {
		goto L99
	}
L98:
	;
	v451 = v441
	goto L95
L99:
	;
	if base.Ui64(int64(1844674407370955161)) < base.Ui64(v416) {
		v494 = v360
		goto L83
	} else {
		goto L100
	}
L100:
	;
	v430 = v416 * int64(10)
	v435 = base.I64_extend_i32_u(v420+int32(-48)) & int64(255)
	if base.Ui64(v435^int64(-1)) < base.Ui64(v430) {
		v494 = v360
		goto L83
	} else {
		goto L101
	}
L101:
	;
	v439 = int32(1)
	v441 = v430 + v435
	v443 = v414 + v439
	if v443 != v357 {
		v414 = v443
		v416 = v441
		v418 = v418 + v439
		goto L97
	} else {
		goto L102
	}
L102:
	;
	goto L98
L103:
	;
	if v451 < int64(0) {
		v494 = v360
		goto L83
	} else {
		goto L107
	}
L104:
	;
	if base.Ui64(int64(-9223372036854775807-1)) < base.Ui64(v451) {
		v494 = v360
		goto L83
	} else {
		goto L105
	}
L105:
	;
	if v359 == int32(0) {
		goto L85
	} else {
		goto L106
	}
L106:
	;
	v475 = int64(0) - v451
	goto L86
L107:
	;
	if v359 == int32(0) {
		goto L85
	} else {
		goto L108
	}
L108:
	;
	v475 = v451
	goto L86
L109:
	;
	v503 = *(*int64)(unsafe.Add(mBase, uint32(v9)+8))
	v504 = v503
	goto L78
L110:
	;
	v511 = *(*int32)(unsafe.Add(mBase, uint32(l0)+392))
	if v509 == v511 {
		goto L71
	} else {
		goto L111
	}
L111:
	;
	v513 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v514 = *(*int32)(unsafe.Add(mBase, uint32(v513)+4))
	v515 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v516 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
	v518 = F_raxInsert(m, v514, v515, v516, v509, int32(0))
	mBase = m.M
	v519 = m.ExcPending
	if v519 != 0 {
		goto L6
	} else {
		goto L112
	}
L112:
	;
	goto L71
L113:
	;
	v558 = *(*int64)(unsafe.Add(mBase, uint32(v551+int32(8))))
	*(*int64)(unsafe.Add(mBase, uint32(v9+int32(16)))) = v558
	v560 = *(*int64)(unsafe.Add(mBase, uint32(v551)))
	*(*int64)(unsafe.Add(mBase, uint32(v9)+8)) = v560
	F_raxStop(m, l0+int32(88))
	mBase = m.M
	v565 = m.ExcPending
	if v565 != 0 {
		goto L6
	} else {
		goto L116
	}
L114:
	;
	v545 = *(*int64)(unsafe.Add(mBase, uint32(l1+int32(8))))
	*(*int64)(unsafe.Add(mBase, uint32(v9+int32(32)))) = v545
	v547 = *(*int64)(unsafe.Add(mBase, uint32(l1)))
	*(*int64)(unsafe.Add(mBase, uint32(v9)+24)) = v547
	v551 = l0 + int32(72)
	goto L113
L115:
	;
	v535 = *(*int64)(unsafe.Add(mBase, uint32(l0+int32(64))))
	*(*int64)(unsafe.Add(mBase, uint32(v9+int32(32)))) = v535
	v537 = *(*int64)(unsafe.Add(mBase, uint32(l0)+56))
	*(*int64)(unsafe.Add(mBase, uint32(v9)+24)) = v537
	v551 = l1
	goto L113
L116:
	;
	v566 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v571 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	F_streamIteratorStart(m, l0, v566, v9+int32(24), v9+int32(8), v571)
	mBase = m.M
	v573 = m.ExcPending
	if v573 != 0 {
		goto L6
	} else {
		goto L117
	}
L117:
	;
	m.G0 = v9 + int32(48)
	return
L118:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L119:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L120:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_streamIteratorStop(m *base.Module, l0 int32) {
	var v5 int32
	_ = v5
	F_raxStop(m, l0+int32(88))
	v5 = m.ExcPending
	if v5 != 0 {
		return
	} else {
		return
	}
}
func F_streamParseAddOrTrimArgsOrReply(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v99 int32
	_ = v99
	var v102 int32
	_ = v102
	var v109 int32
	_ = v109
	var v114 int32
	_ = v114
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v124 int32
	_ = v124
	var v129 int32
	_ = v129
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v136 int32
	_ = v136
	var v138 int32
	_ = v138
	var v142 int32
	_ = v142
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v150 int64
	_ = v150
	var v155 int32
	_ = v155
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
	var v199 int32
	_ = v199
	var v204 int32
	_ = v204
	var v209 int32
	_ = v209
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v219 int32
	_ = v219
	var v224 int32
	_ = v224
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
	var v237 int32
	_ = v237
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v250 int32
	_ = v250
	var v253 int32
	_ = v253
	var v255 int32
	_ = v255
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v259 int32
	_ = v259
	var v263 int32
	_ = v263
	var v265 int32
	_ = v265
	var v266 int32
	_ = v266
	var v267 int32
	_ = v267
	var v269 int32
	_ = v269
	var v270 int32
	_ = v270
	var v271 int32
	_ = v271
	var v272 int32
	_ = v272
	var v277 int32
	_ = v277
	var v278 int32
	_ = v278
	var v281 int32
	_ = v281
	var v282 int32
	_ = v282
	var v285 int32
	_ = v285
	var v286 int32
	_ = v286
	var v287 int32
	_ = v287
	var v291 int32
	_ = v291
	var v293 int32
	_ = v293
	var v297 int32
	_ = v297
	var v299 int32
	_ = v299
	var v300 int32
	_ = v300
	var v306 int64
	_ = v306
	var v311 int32
	_ = v311
	var v316 int32
	_ = v316
	var v319 int32
	_ = v319
	var v321 int32
	_ = v321
	var v322 int32
	_ = v322
	var v323 int32
	_ = v323
	var v325 int32
	_ = v325
	var v329 int32
	_ = v329
	var v331 int32
	_ = v331
	var v332 int32
	_ = v332
	var v333 int32
	_ = v333
	var v335 int32
	_ = v335
	var v336 int32
	_ = v336
	var v337 int32
	_ = v337
	var v338 int32
	_ = v338
	var v343 int32
	_ = v343
	var v344 int32
	_ = v344
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
	var v357 int32
	_ = v357
	var v361 int32
	_ = v361
	var v366 int32
	_ = v366
	var v367 int32
	_ = v367
	var v375 int32
	_ = v375
	var v377 int32
	_ = v377
	var v383 int32
	_ = v383
	var v384 int32
	_ = v384
	var v389 int32
	_ = v389
	var v390 int32
	_ = v390
	var v393 int32
	_ = v393
	var v395 int32
	_ = v395
	var v405 int64
	_ = v405
	var v408 int32
	_ = v408
	var v411 int32
	_ = v411
	var v414 int32
	_ = v414
	var v417 int32
	_ = v417
	var v421 int64
	_ = v421
	var v425 int32
	_ = v425
	var v426 int32
	_ = v426
	var v429 int32
	_ = v429
	var v431 int32
	_ = v431
	var v433 int32
	_ = v433
	var v436 int32
	_ = v436
	var v442 int32
	_ = v442
	var v447 int32
	_ = v447
	var v454 int64
	_ = v454
	var v456 int64
	_ = v456
	var v457 int64
	_ = v457
	var v460 int64
	_ = v460
	var v463 int64
	_ = v463
	v4 = int32(0)
	v18 = F__emscripten_memset_bulkmem(m, l1, base.I32_extend8_s(v4), int32(72))
	mBase = m.M
	goto L1
L1:
	;
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v20 < int32(3) {
		v393 = int32(2)
		v395 = v4
		goto L2
	} else {
		goto L3
	}
L2:
	;
	v405 = *(*int64)(unsafe.Add(mBase, uint32(v18)+40))
	if v405 == int64(0) {
		goto L116
	} else {
		goto L117
	}
L3:
	;
	v24 = v18 + int32(48)
	v28 = v18 + int32(40)
	v32 = int32(2)
	v34 = int32(0)
	v36 = v20
	goto L4
L4:
	;
	v44 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v44+v32<<(uint(int32(2))%32))))
	v49 = F_objectGetVal(m, v48)
	mBase = m.M
	if l2 == int32(0) {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	v393 = v389
	v395 = v384
	goto L2
L6:
	;
	v58 = int32(_a1537)
	v61 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v49))))
	if v61 != 0 {
		goto L12
	} else {
		goto L13
	}
L7:
	;
	v52 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v49))))
	if v52 != int32(42) {
		goto L6
	} else {
		goto L8
	}
L8:
	;
	v55 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v49)+1)))
	if v55 == int32(0) {
		v393 = v32
		v395 = v34
		goto L2
	} else {
		goto L9
	}
L9:
	;
	goto L6
L10:
	;
	v99 = v36 + (v32 ^ int32(-1))
	if v99 == int32(0) {
		goto L24
	} else {
		goto L25
	}
L11:
	;
	v93 = F_tolower(m, v89)
	mBase = m.M
	v94 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v90))))
	v95 = F_tolower(m, v94)
	mBase = m.M
	goto L10
L12:
	;
	v63 = v49
	v64 = v58
	v65 = v61
	goto L15
L13:
	;
	v89 = int32(0)
	v90 = v58
	goto L11
L14:
	;
	v89 = v86 & int32(255)
	v90 = v85
	goto L11
L15:
	;
	v67 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v64))))
	if v67 == int32(0) {
		v85 = v64
		v86 = v65
		goto L14
	} else {
		goto L17
	}
L16:
	;
	v85 = v79
	v86 = int32(0)
	goto L14
L17:
	;
	v71 = v65 & int32(255)
	if v71 == v67 {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	v78 = int32(1)
	v79 = v64 + v78
	v80 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v63)+1)))
	if v80 != 0 {
		v63 = v63 + v78
		v64 = v79
		v65 = v80
		goto L15
	} else {
		goto L21
	}
L19:
	;
	v73 = F_tolower(m, v71)
	mBase = m.M
	v74 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v64))))
	v75 = F_tolower(m, v74)
	mBase = m.M
	if v73 == v75 {
		goto L18
	} else {
		goto L20
	}
L20:
	;
	v77 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v63))))
	v85 = v64
	v86 = v77
	goto L14
L21:
	;
	goto L16
L22:
	;
	v389 = v383 + int32(1)
	v390 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v389 < v390 {
		v32 = v389
		v34 = v384
		v36 = v390
		goto L4
	} else {
		goto L115
	}
L23:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+32)) = v138
	*(*int32)(unsafe.Add(mBase, uint32(v18)+28)) = int32(1)
	v383 = v138
	v384 = v34
	goto L22
L24:
	;
	v158 = int32(_a1538)
	v161 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v49))))
	if v161 != 0 {
		goto L47
	} else {
		goto L48
	}
L25:
	;
	if v93-v95 != 0 {
		goto L24
	} else {
		goto L26
	}
L26:
	;
	v102 = *(*int32)(unsafe.Add(mBase, uint32(v18)+28))
	if v102 == int32(0) {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+36)) = int32(0)
	v114 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v116 = v32 + int32(1)
	v117 = int32(2)
	v120 = *(*int32)(unsafe.Add(mBase, uint32(v114+v116<<(uint(v117)%32))))
	v121 = F_objectGetVal(m, v120)
	mBase = m.M
	if v99 < v117 {
		v134 = v32
		goto L31
	} else {
		goto L32
	}
L28:
	;
	F_addReplyError(m, l0, int32(_a1539))
	mBase = m.M
	v109 = m.ExcPending
	if v109 != 0 {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	return int32(0)
L30:
	;
	return int32(-1)
L31:
	;
	v136 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v138 = v134 + int32(1)
	v142 = *(*int32)(unsafe.Add(mBase, uint32(v136+v138<<(uint(int32(2))%32))))
	v144 = F_getLongLongFromObjectOrReply(m, l0, v142, v24, int32(0))
	mBase = m.M
	v145 = m.ExcPending
	if v145 != 0 {
		goto L29
	} else {
		goto L41
	}
L32:
	;
	v124 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v121))))
	if v124 == int32(61) {
		goto L33
	} else {
		goto L34
	}
L33:
	;
	v132 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v121)+1)))
	if v132 != 0 {
		goto L37
	} else {
		goto L38
	}
L34:
	;
	if v124 != int32(126) {
		v134 = v32
		goto L31
	} else {
		goto L35
	}
L35:
	;
	v129 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v121)+1)))
	if v129 != 0 {
		v134 = v32
		goto L31
	} else {
		goto L36
	}
L36:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+36)) = int32(1)
	v134 = v116
	goto L31
L37:
	;
	v133 = v32
	goto L39
L38:
	;
	v133 = v116
	goto L39
L39:
	;
	v134 = v133
	goto L31
L40:
	;
	v150 = *(*int64)(unsafe.Add(mBase, uint32(v24)))
	if int64(-1) < v150 {
		goto L23
	} else {
		goto L43
	}
L41:
	;
	if v144 == int32(0) {
		goto L40
	} else {
		goto L42
	}
L42:
	;
	return int32(-1)
L43:
	;
	F_addReplyError(m, l0, int32(_a1540))
	mBase = m.M
	v155 = m.ExcPending
	if v155 != 0 {
		goto L29
	} else {
		goto L44
	}
L44:
	;
	return int32(-1)
L45:
	;
	if v99 == int32(0) {
		goto L57
	} else {
		goto L58
	}
L46:
	;
	v193 = F_tolower(m, v189)
	mBase = m.M
	v194 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v190))))
	v195 = F_tolower(m, v194)
	mBase = m.M
	goto L45
L47:
	;
	v163 = v49
	v164 = v158
	v165 = v161
	goto L50
L48:
	;
	v189 = int32(0)
	v190 = v158
	goto L46
L49:
	;
	v189 = v186 & int32(255)
	v190 = v185
	goto L46
L50:
	;
	v167 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v164))))
	if v167 == int32(0) {
		v185 = v164
		v186 = v165
		goto L49
	} else {
		goto L52
	}
L51:
	;
	v185 = v179
	v186 = int32(0)
	goto L49
L52:
	;
	v171 = v165 & int32(255)
	if v171 == v167 {
		goto L53
	} else {
		goto L54
	}
L53:
	;
	v178 = int32(1)
	v179 = v164 + v178
	v180 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v163)+1)))
	if v180 != 0 {
		v163 = v163 + v178
		v164 = v179
		v165 = v180
		goto L50
	} else {
		goto L56
	}
L54:
	;
	v173 = F_tolower(m, v171)
	mBase = m.M
	v174 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v164))))
	v175 = F_tolower(m, v174)
	mBase = m.M
	if v173 == v175 {
		goto L53
	} else {
		goto L55
	}
L55:
	;
	v177 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v163))))
	v185 = v164
	v186 = v177
	goto L49
L56:
	;
	goto L51
L57:
	;
	v250 = int32(_a1541)
	v253 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v49))))
	if v253 != 0 {
		goto L77
	} else {
		goto L78
	}
L58:
	;
	if v193-v195 != 0 {
		goto L57
	} else {
		goto L59
	}
L59:
	;
	v199 = *(*int32)(unsafe.Add(mBase, uint32(v18)+28))
	if v199 == int32(0) {
		goto L60
	} else {
		goto L61
	}
L60:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+36)) = int32(0)
	v209 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v211 = v32 + int32(1)
	v212 = int32(2)
	v215 = *(*int32)(unsafe.Add(mBase, uint32(v209+v211<<(uint(v212)%32))))
	v216 = F_objectGetVal(m, v215)
	mBase = m.M
	if v99 < v212 {
		v229 = v32
		goto L63
	} else {
		goto L64
	}
L61:
	;
	F_addReplyError(m, l0, int32(_a1539))
	mBase = m.M
	v204 = m.ExcPending
	if v204 != 0 {
		goto L29
	} else {
		goto L62
	}
L62:
	;
	return int32(-1)
L63:
	;
	v231 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v232 = int32(1)
	v233 = v229 + v232
	v237 = *(*int32)(unsafe.Add(mBase, uint32(v231+v233<<(uint(int32(2))%32))))
	v241 = F_streamGenericParseIDOrReply(m, l0, v237, v18+int32(56), int64(0), v232, int32(0))
	mBase = m.M
	v242 = m.ExcPending
	if v242 != 0 {
		goto L29
	} else {
		goto L73
	}
L64:
	;
	v219 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v216))))
	if v219 == int32(61) {
		goto L65
	} else {
		goto L66
	}
L65:
	;
	v227 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v216)+1)))
	if v227 != 0 {
		goto L69
	} else {
		goto L70
	}
L66:
	;
	if v219 != int32(126) {
		v229 = v32
		goto L63
	} else {
		goto L67
	}
L67:
	;
	v224 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v216)+1)))
	if v224 != 0 {
		v229 = v32
		goto L63
	} else {
		goto L68
	}
L68:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+36)) = int32(1)
	v229 = v211
	goto L63
L69:
	;
	v228 = v32
	goto L71
L70:
	;
	v228 = v211
	goto L71
L71:
	;
	v229 = v228
	goto L63
L72:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+32)) = v233
	*(*int32)(unsafe.Add(mBase, uint32(v18)+28)) = int32(2)
	v383 = v233
	v384 = v34
	goto L22
L73:
	;
	if v241 == int32(0) {
		goto L72
	} else {
		goto L74
	}
L74:
	;
	return int32(-1)
L75:
	;
	if v99 == int32(0) {
		goto L87
	} else {
		goto L88
	}
L76:
	;
	v285 = F_tolower(m, v281)
	mBase = m.M
	v286 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v282))))
	v287 = F_tolower(m, v286)
	mBase = m.M
	goto L75
L77:
	;
	v255 = v49
	v256 = v250
	v257 = v253
	goto L80
L78:
	;
	v281 = int32(0)
	v282 = v250
	goto L76
L79:
	;
	v281 = v278 & int32(255)
	v282 = v277
	goto L76
L80:
	;
	v259 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v256))))
	if v259 == int32(0) {
		v277 = v256
		v278 = v257
		goto L79
	} else {
		goto L82
	}
L81:
	;
	v277 = v271
	v278 = int32(0)
	goto L79
L82:
	;
	v263 = v257 & int32(255)
	if v263 == v259 {
		goto L83
	} else {
		goto L84
	}
L83:
	;
	v270 = int32(1)
	v271 = v256 + v270
	v272 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v255)+1)))
	if v272 != 0 {
		v255 = v255 + v270
		v256 = v271
		v257 = v272
		goto L80
	} else {
		goto L86
	}
L84:
	;
	v265 = F_tolower(m, v263)
	mBase = m.M
	v266 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v256))))
	v267 = F_tolower(m, v266)
	mBase = m.M
	if v265 == v267 {
		goto L83
	} else {
		goto L85
	}
L85:
	;
	v269 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v255))))
	v277 = v256
	v278 = v269
	goto L79
L86:
	;
	goto L81
L87:
	;
	if l2 == int32(0) {
		goto L95
	} else {
		goto L96
	}
L88:
	;
	if v285-v287 != 0 {
		goto L87
	} else {
		goto L89
	}
L89:
	;
	v291 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v293 = v32 + int32(1)
	v297 = *(*int32)(unsafe.Add(mBase, uint32(v291+v293<<(uint(int32(2))%32))))
	v299 = F_getLongLongFromObjectOrReply(m, l0, v297, v28, int32(0))
	mBase = m.M
	v300 = m.ExcPending
	if v300 != 0 {
		goto L29
	} else {
		goto L91
	}
L90:
	;
	v306 = *(*int64)(unsafe.Add(mBase, uint32(v28)))
	if int64(-1) < v306 {
		v383 = v293
		v384 = int32(1)
		goto L22
	} else {
		goto L93
	}
L91:
	;
	if v299 == int32(0) {
		goto L90
	} else {
		goto L92
	}
L92:
	;
	return int32(-1)
L93:
	;
	F_addReplyError(m, l0, int32(_a1542))
	mBase = m.M
	v311 = m.ExcPending
	if v311 != 0 {
		goto L29
	} else {
		goto L94
	}
L94:
	;
	return int32(-1)
L95:
	;
	v375 = *(*int32)(unsafe.Add(mBase, _consts[59]))
	F_addReplyErrorObject(m, l0, v375)
	mBase = m.M
	v377 = m.ExcPending
	if v377 != 0 {
		goto L29
	} else {
		goto L114
	}
L96:
	;
	v316 = int32(_a1543)
	v319 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v49))))
	if v319 != 0 {
		goto L100
	} else {
		goto L101
	}
L97:
	;
	v357 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v361 = *(*int32)(unsafe.Add(mBase, uint32(v357+v32<<(uint(int32(2))%32))))
	v366 = F_streamGenericParseIDOrReply(m, l0, v361, v18, int64(0), int32(1), v18+int32(20))
	mBase = m.M
	v367 = m.ExcPending
	if v367 != 0 {
		goto L29
	} else {
		goto L112
	}
L98:
	;
	if v351-v353 != 0 {
		goto L97
	} else {
		goto L110
	}
L99:
	;
	v351 = F_tolower(m, v347)
	mBase = m.M
	v352 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v348))))
	v353 = F_tolower(m, v352)
	mBase = m.M
	goto L98
L100:
	;
	v321 = v49
	v322 = v316
	v323 = v319
	goto L103
L101:
	;
	v347 = int32(0)
	v348 = v316
	goto L99
L102:
	;
	v347 = v344 & int32(255)
	v348 = v343
	goto L99
L103:
	;
	v325 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v322))))
	if v325 == int32(0) {
		v343 = v322
		v344 = v323
		goto L102
	} else {
		goto L105
	}
L104:
	;
	v343 = v337
	v344 = int32(0)
	goto L102
L105:
	;
	v329 = v323 & int32(255)
	if v329 == v325 {
		goto L106
	} else {
		goto L107
	}
L106:
	;
	v336 = int32(1)
	v337 = v322 + v336
	v338 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v321)+1)))
	if v338 != 0 {
		v321 = v321 + v336
		v322 = v337
		v323 = v338
		goto L103
	} else {
		goto L109
	}
L107:
	;
	v331 = F_tolower(m, v329)
	mBase = m.M
	v332 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v322))))
	v333 = F_tolower(m, v332)
	mBase = m.M
	if v331 == v333 {
		goto L106
	} else {
		goto L108
	}
L108:
	;
	v335 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v321))))
	v343 = v322
	v344 = v335
	goto L102
L109:
	;
	goto L104
L110:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+24)) = int32(1)
	v383 = v32
	v384 = v34
	goto L22
L111:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+16)) = int32(1)
	v393 = v32
	v395 = v34
	goto L2
L112:
	;
	if v366 == int32(0) {
		goto L111
	} else {
		goto L113
	}
L113:
	;
	return int32(-1)
L114:
	;
	return int32(-1)
L115:
	;
	goto L5
L116:
	;
	if l2 != 0 {
		goto L120
	} else {
		goto L121
	}
L117:
	;
	v408 = *(*int32)(unsafe.Add(mBase, uint32(v18)+28))
	if v408 != 0 {
		goto L116
	} else {
		goto L118
	}
L118:
	;
	F_addReplyError(m, l0, int32(_a1544))
	mBase = m.M
	v411 = m.ExcPending
	if v411 != 0 {
		goto L29
	} else {
		goto L119
	}
L119:
	;
	return int32(-1)
L120:
	;
	v421 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	if v421 != int64(-1) {
		goto L126
	} else {
		goto L127
	}
L121:
	;
	v414 = *(*int32)(unsafe.Add(mBase, uint32(v18)+28))
	if v414 != 0 {
		goto L120
	} else {
		goto L122
	}
L122:
	;
	F_addReplyError(m, l0, int32(_a1545))
	mBase = m.M
	v417 = m.ExcPending
	if v417 != 0 {
		goto L29
	} else {
		goto L123
	}
L123:
	;
	return int32(-1)
L124:
	;
	v442 = *(*int32)(unsafe.Add(mBase, uint32(v18)+36))
	if v395 == int32(0) {
		goto L134
	} else {
		goto L135
	}
L125:
	;
	if v436 == int32(0) {
		goto L124
	} else {
		goto L132
	}
L126:
	;
	v425 = int32(1)
	v426 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+200)))
	if v426&v425 != 0 {
		v433 = v425
		goto L128
	} else {
		goto L129
	}
L127:
	;
	v436 = int32(1)
	goto L125
L128:
	;
	v436 = v433
	goto L125
L129:
	;
	v429 = *(*int32)(unsafe.Add(mBase, uint32(l0)+216))
	if v429 != 0 {
		goto L130
	} else {
		goto L131
	}
L130:
	;
	v431 = F_isImportSlotMigrationJob(m, v429)
	mBase = m.M
	v433 = v431
	goto L128
L131:
	;
	v436 = int32(0)
	goto L125
L132:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v18)+40)) = int64(0)
	return v393
L133:
	;
	return v393
L134:
	;
	if v442 == int32(0) {
		goto L138
	} else {
		goto L139
	}
L135:
	;
	if v442 != 0 {
		goto L133
	} else {
		goto L136
	}
L136:
	;
	F_addReplyError(m, l0, int32(_a1546))
	mBase = m.M
	v447 = m.ExcPending
	if v447 != 0 {
		goto L29
	} else {
		goto L137
	}
L137:
	;
	return int32(-1)
L138:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v18)+40)) = int64(0)
	goto L133
L139:
	;
	v454 = *(*int64)(unsafe.Add(mBase, _consts[1080]))
	v456 = v454 * int64(100)
	v457 = int64(1000000)
	if v456 < v457 {
		goto L140
	} else {
		goto L141
	}
L140:
	;
	v460 = v456
	goto L142
L141:
	;
	v460 = v457
	goto L142
L142:
	;
	if v454 < int64(1) {
		goto L143
	} else {
		goto L144
	}
L143:
	;
	v463 = int64(10000)
	goto L145
L144:
	;
	v463 = v460
	goto L145
L145:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v18)+40)) = v463
	return v393
}
func F_streamParseIntervalIDOrReply(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int64) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v23 int32
	_ = v23
	var v28 int32
	_ = v28
	var v33 int32
	_ = v33
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v77 int32
	_ = v77
	var v81 int32
	_ = v81
	v9 = F_objectGetVal(m, l1)
	mBase = m.M
	v12 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9+int32(-1)))))
	switch v12 & int32(7) {
	case 0:
		if l3 == int32(0) {
			v69 = int32(0)
			v71 = F_streamGenericParseIDOrReply(m, l0, l1, l2, l4, v69, v69)
			mBase = m.M
			v72 = m.ExcPending
			if v72 != 0 {
				return int32(0)
			} else {
				v73 = v71
				v77 = int32(-1)
				if v73 == v77 {
					v81 = v77
				} else {
					v81 = int32(0)
				}
				return v81
			}
		} else {
			v39 = int32(base.Ui32(v12) >> (uint(int32(3)) % 32))
			if base.Ui32(v39) <= base.Ui32(int32(1)) {
				*(*int32)(unsafe.Add(mBase, uint32(l3))) = int32(0)
				v69 = int32(0)
				v71 = F_streamGenericParseIDOrReply(m, l0, l1, l2, l4, v69, v69)
				mBase = m.M
				v72 = m.ExcPending
				if v72 != 0 {
					return int32(0)
				} else {
					v73 = v71
					v77 = int32(-1)
					if v73 == v77 {
						v81 = v77
					} else {
						v81 = int32(0)
					}
					return v81
				}
			} else {
				v42 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9))))
				v43 = int32(40)
				*(*int32)(unsafe.Add(mBase, uint32(l3))) = base.B2i32(v42 == v43)
				if v42 != v43 {
					v69 = int32(0)
					v71 = F_streamGenericParseIDOrReply(m, l0, l1, l2, l4, v69, v69)
					mBase = m.M
					v72 = m.ExcPending
					if v72 != 0 {
						return int32(0)
					} else {
						v73 = v71
						v77 = int32(-1)
						if v73 == v77 {
							v81 = v77
						} else {
							v81 = int32(0)
						}
						return v81
					}
				} else {
					v52 = F_createStringObject_1(m, v9+int32(1), v39+int32(-1))
					mBase = m.M
					v55 = m.ExcPending
					if v55 != 0 {
						return int32(0)
					} else {
						v58 = F_streamGenericParseIDOrReply(m, l0, v52, l2, l4, int32(1), int32(0))
						mBase = m.M
						v59 = m.ExcPending
						if v59 != 0 {
							return int32(0)
						} else {
							F_decrRefCount(m, v52)
							mBase = m.M
							v61 = m.ExcPending
							if v61 != 0 {
								return int32(0)
							} else {
								v73 = v58
								v77 = int32(-1)
								if v73 == v77 {
									v81 = v77
								} else {
									v81 = int32(0)
								}
								return v81
							}
						}
					}
				}
			}
		}
	case 1:
		if l3 == int32(0) {
			v69 = int32(0)
			v71 = F_streamGenericParseIDOrReply(m, l0, l1, l2, l4, v69, v69)
			mBase = m.M
			v72 = m.ExcPending
			if v72 != 0 {
				return int32(0)
			} else {
				v73 = v71
				v77 = int32(-1)
				if v73 == v77 {
					v81 = v77
				} else {
					v81 = int32(0)
				}
				return v81
			}
		} else {
			v23 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9+int32(-3)))))
			v39 = v23
			if base.Ui32(v39) <= base.Ui32(int32(1)) {
				*(*int32)(unsafe.Add(mBase, uint32(l3))) = int32(0)
				v69 = int32(0)
				v71 = F_streamGenericParseIDOrReply(m, l0, l1, l2, l4, v69, v69)
				mBase = m.M
				v72 = m.ExcPending
				if v72 != 0 {
					return int32(0)
				} else {
					v73 = v71
					v77 = int32(-1)
					if v73 == v77 {
						v81 = v77
					} else {
						v81 = int32(0)
					}
					return v81
				}
			} else {
				v42 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9))))
				v43 = int32(40)
				*(*int32)(unsafe.Add(mBase, uint32(l3))) = base.B2i32(v42 == v43)
				if v42 != v43 {
					v69 = int32(0)
					v71 = F_streamGenericParseIDOrReply(m, l0, l1, l2, l4, v69, v69)
					mBase = m.M
					v72 = m.ExcPending
					if v72 != 0 {
						return int32(0)
					} else {
						v73 = v71
						v77 = int32(-1)
						if v73 == v77 {
							v81 = v77
						} else {
							v81 = int32(0)
						}
						return v81
					}
				} else {
					v52 = F_createStringObject_1(m, v9+int32(1), v39+int32(-1))
					mBase = m.M
					v55 = m.ExcPending
					if v55 != 0 {
						return int32(0)
					} else {
						v58 = F_streamGenericParseIDOrReply(m, l0, v52, l2, l4, int32(1), int32(0))
						mBase = m.M
						v59 = m.ExcPending
						if v59 != 0 {
							return int32(0)
						} else {
							F_decrRefCount(m, v52)
							mBase = m.M
							v61 = m.ExcPending
							if v61 != 0 {
								return int32(0)
							} else {
								v73 = v58
								v77 = int32(-1)
								if v73 == v77 {
									v81 = v77
								} else {
									v81 = int32(0)
								}
								return v81
							}
						}
					}
				}
			}
		}
	case 2:
		if l3 == int32(0) {
			v69 = int32(0)
			v71 = F_streamGenericParseIDOrReply(m, l0, l1, l2, l4, v69, v69)
			mBase = m.M
			v72 = m.ExcPending
			if v72 != 0 {
				return int32(0)
			} else {
				v73 = v71
				v77 = int32(-1)
				if v73 == v77 {
					v81 = v77
				} else {
					v81 = int32(0)
				}
				return v81
			}
		} else {
			v28 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v9+int32(-5)))))
			v39 = v28
			if base.Ui32(v39) <= base.Ui32(int32(1)) {
				*(*int32)(unsafe.Add(mBase, uint32(l3))) = int32(0)
				v69 = int32(0)
				v71 = F_streamGenericParseIDOrReply(m, l0, l1, l2, l4, v69, v69)
				mBase = m.M
				v72 = m.ExcPending
				if v72 != 0 {
					return int32(0)
				} else {
					v73 = v71
					v77 = int32(-1)
					if v73 == v77 {
						v81 = v77
					} else {
						v81 = int32(0)
					}
					return v81
				}
			} else {
				v42 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9))))
				v43 = int32(40)
				*(*int32)(unsafe.Add(mBase, uint32(l3))) = base.B2i32(v42 == v43)
				if v42 != v43 {
					v69 = int32(0)
					v71 = F_streamGenericParseIDOrReply(m, l0, l1, l2, l4, v69, v69)
					mBase = m.M
					v72 = m.ExcPending
					if v72 != 0 {
						return int32(0)
					} else {
						v73 = v71
						v77 = int32(-1)
						if v73 == v77 {
							v81 = v77
						} else {
							v81 = int32(0)
						}
						return v81
					}
				} else {
					v52 = F_createStringObject_1(m, v9+int32(1), v39+int32(-1))
					mBase = m.M
					v55 = m.ExcPending
					if v55 != 0 {
						return int32(0)
					} else {
						v58 = F_streamGenericParseIDOrReply(m, l0, v52, l2, l4, int32(1), int32(0))
						mBase = m.M
						v59 = m.ExcPending
						if v59 != 0 {
							return int32(0)
						} else {
							F_decrRefCount(m, v52)
							mBase = m.M
							v61 = m.ExcPending
							if v61 != 0 {
								return int32(0)
							} else {
								v73 = v58
								v77 = int32(-1)
								if v73 == v77 {
									v81 = v77
								} else {
									v81 = int32(0)
								}
								return v81
							}
						}
					}
				}
			}
		}
	case 3:
		if l3 == int32(0) {
			v69 = int32(0)
			v71 = F_streamGenericParseIDOrReply(m, l0, l1, l2, l4, v69, v69)
			mBase = m.M
			v72 = m.ExcPending
			if v72 != 0 {
				return int32(0)
			} else {
				v73 = v71
				v77 = int32(-1)
				if v73 == v77 {
					v81 = v77
				} else {
					v81 = int32(0)
				}
				return v81
			}
		} else {
			v33 = *(*int32)(unsafe.Add(mBase, uint32(v9+int32(-9))))
			v39 = v33
			if base.Ui32(v39) <= base.Ui32(int32(1)) {
				*(*int32)(unsafe.Add(mBase, uint32(l3))) = int32(0)
				v69 = int32(0)
				v71 = F_streamGenericParseIDOrReply(m, l0, l1, l2, l4, v69, v69)
				mBase = m.M
				v72 = m.ExcPending
				if v72 != 0 {
					return int32(0)
				} else {
					v73 = v71
					v77 = int32(-1)
					if v73 == v77 {
						v81 = v77
					} else {
						v81 = int32(0)
					}
					return v81
				}
			} else {
				v42 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9))))
				v43 = int32(40)
				*(*int32)(unsafe.Add(mBase, uint32(l3))) = base.B2i32(v42 == v43)
				if v42 != v43 {
					v69 = int32(0)
					v71 = F_streamGenericParseIDOrReply(m, l0, l1, l2, l4, v69, v69)
					mBase = m.M
					v72 = m.ExcPending
					if v72 != 0 {
						return int32(0)
					} else {
						v73 = v71
						v77 = int32(-1)
						if v73 == v77 {
							v81 = v77
						} else {
							v81 = int32(0)
						}
						return v81
					}
				} else {
					v52 = F_createStringObject_1(m, v9+int32(1), v39+int32(-1))
					mBase = m.M
					v55 = m.ExcPending
					if v55 != 0 {
						return int32(0)
					} else {
						v58 = F_streamGenericParseIDOrReply(m, l0, v52, l2, l4, int32(1), int32(0))
						mBase = m.M
						v59 = m.ExcPending
						if v59 != 0 {
							return int32(0)
						} else {
							F_decrRefCount(m, v52)
							mBase = m.M
							v61 = m.ExcPending
							if v61 != 0 {
								return int32(0)
							} else {
								v73 = v58
								v77 = int32(-1)
								if v73 == v77 {
									v81 = v77
								} else {
									v81 = int32(0)
								}
								return v81
							}
						}
					}
				}
			}
		}
	case 4:
		if l3 == int32(0) {
			v69 = int32(0)
			v71 = F_streamGenericParseIDOrReply(m, l0, l1, l2, l4, v69, v69)
			mBase = m.M
			v72 = m.ExcPending
			if v72 != 0 {
				return int32(0)
			} else {
				v73 = v71
				v77 = int32(-1)
				if v73 == v77 {
					v81 = v77
				} else {
					v81 = int32(0)
				}
				return v81
			}
		} else {
			v38 = *(*int32)(unsafe.Add(mBase, uint32(v9+int32(-17))))
			v39 = v38
			if base.Ui32(v39) <= base.Ui32(int32(1)) {
				*(*int32)(unsafe.Add(mBase, uint32(l3))) = int32(0)
				v69 = int32(0)
				v71 = F_streamGenericParseIDOrReply(m, l0, l1, l2, l4, v69, v69)
				mBase = m.M
				v72 = m.ExcPending
				if v72 != 0 {
					return int32(0)
				} else {
					v73 = v71
					v77 = int32(-1)
					if v73 == v77 {
						v81 = v77
					} else {
						v81 = int32(0)
					}
					return v81
				}
			} else {
				v42 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9))))
				v43 = int32(40)
				*(*int32)(unsafe.Add(mBase, uint32(l3))) = base.B2i32(v42 == v43)
				if v42 != v43 {
					v69 = int32(0)
					v71 = F_streamGenericParseIDOrReply(m, l0, l1, l2, l4, v69, v69)
					mBase = m.M
					v72 = m.ExcPending
					if v72 != 0 {
						return int32(0)
					} else {
						v73 = v71
						v77 = int32(-1)
						if v73 == v77 {
							v81 = v77
						} else {
							v81 = int32(0)
						}
						return v81
					}
				} else {
					v52 = F_createStringObject_1(m, v9+int32(1), v39+int32(-1))
					mBase = m.M
					v55 = m.ExcPending
					if v55 != 0 {
						return int32(0)
					} else {
						v58 = F_streamGenericParseIDOrReply(m, l0, v52, l2, l4, int32(1), int32(0))
						mBase = m.M
						v59 = m.ExcPending
						if v59 != 0 {
							return int32(0)
						} else {
							F_decrRefCount(m, v52)
							mBase = m.M
							v61 = m.ExcPending
							if v61 != 0 {
								return int32(0)
							} else {
								v73 = v58
								v77 = int32(-1)
								if v73 == v77 {
									v81 = v77
								} else {
									v81 = int32(0)
								}
								return v81
							}
						}
					}
				}
			}
		}
	default:
		if l3 == int32(0) {
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(l3))) = int32(0)
		}
		v69 = int32(0)
		v71 = F_streamGenericParseIDOrReply(m, l0, l1, l2, l4, v69, v69)
		mBase = m.M
		v72 = m.ExcPending
		if v72 != 0 {
			return int32(0)
		} else {
			v73 = v71
			v77 = int32(-1)
			if v73 == v77 {
				v81 = v77
			} else {
				v81 = int32(0)
			}
			return v81
		}
	}
}
func F_streamTrim(m *base.Module, l0 int32, l1 int32) int64 {
	mBase := m.M
	_ = mBase
	var v4 int64
	_ = v4
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v28 int64
	_ = v28
	var v29 int32
	_ = v29
	var v30 int64
	_ = v30
	var v33 int64
	_ = v33
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v43 int64
	_ = v43
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v72 int32
	_ = v72
	var v88 int64
	_ = v88
	var v97 int64
	_ = v97
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v107 int64
	_ = v107
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v128 int32
	_ = v128
	var v144 int32
	_ = v144
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v159 int64
	_ = v159
	var v165 int32
	_ = v165
	var v167 int64
	_ = v167
	var v169 int32
	_ = v169
	var v171 int32
	_ = v171
	var v181 int64
	_ = v181
	var v186 int64
	_ = v186
	var v190 int32
	_ = v190
	var v192 int64
	_ = v192
	var v194 int32
	_ = v194
	var v202 int64
	_ = v202
	var v226 int64
	_ = v226
	var v245 int32
	_ = v245
	var v254 int64
	_ = v254
	var v255 int64
	_ = v255
	var v256 int64
	_ = v256
	var v260 int64
	_ = v260
	var v264 int64
	_ = v264
	var v267 int32
	_ = v267
	var v268 int64
	_ = v268
	var v269 int64
	_ = v269
	var v270 int64
	_ = v270
	var v272 int64
	_ = v272
	var v274 int64
	_ = v274
	var v277 int64
	_ = v277
	var v279 int64
	_ = v279
	var v281 int64
	_ = v281
	var v283 int64
	_ = v283
	var v346 int64
	_ = v346
	var v350 int32
	_ = v350
	var v356 int32
	_ = v356
	var v357 int32
	_ = v357
	var v358 int64
	_ = v358
	var v359 int64
	_ = v359
	var v364 int64
	_ = v364
	var v365 int64
	_ = v365
	var v368 int32
	_ = v368
	var v373 int32
	_ = v373
	var v374 int32
	_ = v374
	var v375 int32
	_ = v375
	var v376 int32
	_ = v376
	var v380 int32
	_ = v380
	var v381 int32
	_ = v381
	var v382 int64
	_ = v382
	var v383 int32
	_ = v383
	var v385 int32
	_ = v385
	var v386 int32
	_ = v386
	var v398 int32
	_ = v398
	var v399 int32
	_ = v399
	var v403 int32
	_ = v403
	var v419 int32
	_ = v419
	var v421 int32
	_ = v421
	var v422 int32
	_ = v422
	var v423 int32
	_ = v423
	var v434 int64
	_ = v434
	var v440 int32
	_ = v440
	var v442 int64
	_ = v442
	var v444 int32
	_ = v444
	var v446 int32
	_ = v446
	var v456 int64
	_ = v456
	var v461 int64
	_ = v461
	var v465 int32
	_ = v465
	var v467 int64
	_ = v467
	var v469 int32
	_ = v469
	var v477 int64
	_ = v477
	var v501 int64
	_ = v501
	var v520 int32
	_ = v520
	var v529 int64
	_ = v529
	var v530 int64
	_ = v530
	var v532 int32
	_ = v532
	var v533 int32
	_ = v533
	var v549 int32
	_ = v549
	var v551 int64
	_ = v551
	var v555 int32
	_ = v555
	var v556 int32
	_ = v556
	var v558 int64
	_ = v558
	var v573 int32
	_ = v573
	var v579 int32
	_ = v579
	var v580 int32
	_ = v580
	var v597 int32
	_ = v597
	var v599 int32
	_ = v599
	var v603 int64
	_ = v603
	var v609 int32
	_ = v609
	var v610 int32
	_ = v610
	var v611 int64
	_ = v611
	var v612 int32
	_ = v612
	var v614 int32
	_ = v614
	var v615 int32
	_ = v615
	var v627 int32
	_ = v627
	var v628 int32
	_ = v628
	var v632 int32
	_ = v632
	var v648 int32
	_ = v648
	var v650 int32
	_ = v650
	var v651 int32
	_ = v651
	var v652 int32
	_ = v652
	var v663 int64
	_ = v663
	var v669 int32
	_ = v669
	var v671 int64
	_ = v671
	var v673 int32
	_ = v673
	var v675 int32
	_ = v675
	var v685 int64
	_ = v685
	var v690 int64
	_ = v690
	var v694 int32
	_ = v694
	var v696 int64
	_ = v696
	var v698 int32
	_ = v698
	var v706 int64
	_ = v706
	var v730 int64
	_ = v730
	var v749 int32
	_ = v749
	var v758 int64
	_ = v758
	var v759 int64
	_ = v759
	var v760 int32
	_ = v760
	var v761 int32
	_ = v761
	var v766 int32
	_ = v766
	var v767 int32
	_ = v767
	var v768 int64
	_ = v768
	var v769 int32
	_ = v769
	var v771 int32
	_ = v771
	var v772 int32
	_ = v772
	var v784 int32
	_ = v784
	var v785 int32
	_ = v785
	var v789 int32
	_ = v789
	var v805 int32
	_ = v805
	var v807 int32
	_ = v807
	var v808 int32
	_ = v808
	var v809 int32
	_ = v809
	var v820 int64
	_ = v820
	var v826 int32
	_ = v826
	var v828 int64
	_ = v828
	var v830 int32
	_ = v830
	var v832 int32
	_ = v832
	var v842 int64
	_ = v842
	var v847 int64
	_ = v847
	var v851 int32
	_ = v851
	var v853 int64
	_ = v853
	var v855 int32
	_ = v855
	var v863 int64
	_ = v863
	var v887 int64
	_ = v887
	var v906 int32
	_ = v906
	var v915 int64
	_ = v915
	var v916 int64
	_ = v916
	var v917 int32
	_ = v917
	var v918 int32
	_ = v918
	var v923 int32
	_ = v923
	var v924 int32
	_ = v924
	var v925 int64
	_ = v925
	var v926 int32
	_ = v926
	var v928 int32
	_ = v928
	var v929 int32
	_ = v929
	var v941 int32
	_ = v941
	var v942 int32
	_ = v942
	var v946 int32
	_ = v946
	var v962 int32
	_ = v962
	var v964 int32
	_ = v964
	var v965 int32
	_ = v965
	var v966 int32
	_ = v966
	var v977 int64
	_ = v977
	var v983 int32
	_ = v983
	var v985 int64
	_ = v985
	var v987 int32
	_ = v987
	var v989 int32
	_ = v989
	var v999 int64
	_ = v999
	var v1004 int64
	_ = v1004
	var v1008 int32
	_ = v1008
	var v1010 int64
	_ = v1010
	var v1012 int32
	_ = v1012
	var v1020 int64
	_ = v1020
	var v1044 int64
	_ = v1044
	var v1063 int32
	_ = v1063
	var v1072 int64
	_ = v1072
	var v1073 int64
	_ = v1073
	var v1074 int32
	_ = v1074
	var v1075 int32
	_ = v1075
	var v1079 int64
	_ = v1079
	var v1080 int64
	_ = v1080
	var v1082 int64
	_ = v1082
	var v1083 int64
	_ = v1083
	var v1084 int64
	_ = v1084
	var v1086 int64
	_ = v1086
	var v1088 int64
	_ = v1088
	var v1089 int64
	_ = v1089
	var v1090 int64
	_ = v1090
	var v1092 int64
	_ = v1092
	var v1094 int64
	_ = v1094
	var v1100 int32
	_ = v1100
	var v1106 int32
	_ = v1106
	var v1112 int32
	_ = v1112
	var v1124 int32
	_ = v1124
	var v1125 int32
	_ = v1125
	var v1126 int64
	_ = v1126
	var v1127 int32
	_ = v1127
	var v1129 int32
	_ = v1129
	var v1130 int32
	_ = v1130
	var v1142 int32
	_ = v1142
	var v1143 int32
	_ = v1143
	var v1147 int32
	_ = v1147
	var v1163 int32
	_ = v1163
	var v1165 int32
	_ = v1165
	var v1166 int32
	_ = v1166
	var v1167 int32
	_ = v1167
	var v1178 int64
	_ = v1178
	var v1184 int32
	_ = v1184
	var v1186 int64
	_ = v1186
	var v1188 int32
	_ = v1188
	var v1190 int32
	_ = v1190
	var v1200 int64
	_ = v1200
	var v1205 int64
	_ = v1205
	var v1209 int32
	_ = v1209
	var v1211 int64
	_ = v1211
	var v1213 int32
	_ = v1213
	var v1221 int64
	_ = v1221
	var v1245 int64
	_ = v1245
	var v1264 int32
	_ = v1264
	var v1273 int64
	_ = v1273
	var v1274 int64
	_ = v1274
	var v1275 int32
	_ = v1275
	var v1276 int32
	_ = v1276
	var v1281 int32
	_ = v1281
	var v1282 int64
	_ = v1282
	var v1298 int32
	_ = v1298
	var v1300 int64
	_ = v1300
	var v1304 int32
	_ = v1304
	var v1305 int32
	_ = v1305
	var v1308 int64
	_ = v1308
	var v1324 int32
	_ = v1324
	var v1330 int32
	_ = v1330
	var v1331 int32
	_ = v1331
	var v1340 int32
	_ = v1340
	var v1341 int32
	_ = v1341
	var v1342 int32
	_ = v1342
	var v1346 int32
	_ = v1346
	var v1347 int32
	_ = v1347
	var v1349 int32
	_ = v1349
	var v1350 int64
	_ = v1350
	var v1357 int32
	_ = v1357
	var v1358 int32
	_ = v1358
	var v1359 int64
	_ = v1359
	var v1371 int32
	_ = v1371
	var v1377 int64
	_ = v1377
	var v1379 int32
	_ = v1379
	var v1380 int32
	_ = v1380
	var v1385 int32
	_ = v1385
	var v1386 int32
	_ = v1386
	var v1387 int32
	_ = v1387
	var v1388 int32
	_ = v1388
	var v1389 int32
	_ = v1389
	var v1394 int32
	_ = v1394
	var v1395 int32
	_ = v1395
	var v1396 int64
	_ = v1396
	var v1397 int32
	_ = v1397
	var v1399 int32
	_ = v1399
	var v1400 int32
	_ = v1400
	var v1412 int32
	_ = v1412
	var v1413 int32
	_ = v1413
	var v1417 int32
	_ = v1417
	var v1433 int32
	_ = v1433
	var v1435 int32
	_ = v1435
	var v1436 int32
	_ = v1436
	var v1437 int32
	_ = v1437
	var v1448 int64
	_ = v1448
	var v1454 int32
	_ = v1454
	var v1456 int64
	_ = v1456
	var v1458 int32
	_ = v1458
	var v1460 int32
	_ = v1460
	var v1470 int64
	_ = v1470
	var v1475 int64
	_ = v1475
	var v1479 int32
	_ = v1479
	var v1481 int64
	_ = v1481
	var v1483 int32
	_ = v1483
	var v1491 int64
	_ = v1491
	var v1515 int64
	_ = v1515
	var v1534 int32
	_ = v1534
	var v1543 int64
	_ = v1543
	var v1544 int64
	_ = v1544
	var v1549 int32
	_ = v1549
	var v1550 int32
	_ = v1550
	var v1551 int32
	_ = v1551
	var v1552 int32
	_ = v1552
	var v1553 int32
	_ = v1553
	var v1555 int32
	_ = v1555
	var v1556 int32
	_ = v1556
	var v1557 int32
	_ = v1557
	var v1559 int32
	_ = v1559
	var v1560 int32
	_ = v1560
	var v1565 int32
	_ = v1565
	var v1571 int32
	_ = v1571
	var v1574 int32
	_ = v1574
	var v1575 int32
	_ = v1575
	var v1576 int32
	_ = v1576
	var v1577 int32
	_ = v1577
	var v1579 int32
	_ = v1579
	var v1580 int32
	_ = v1580
	var v1584 int32
	_ = v1584
	var v1585 int32
	_ = v1585
	var v1586 int32
	_ = v1586
	var v1587 int32
	_ = v1587
	var v1588 int64
	_ = v1588
	var v1593 int32
	_ = v1593
	var v1594 int32
	_ = v1594
	var v1599 int32
	_ = v1599
	var v1605 int32
	_ = v1605
	var v1610 int64
	_ = v1610
	var v1629 int32
	_ = v1629
	var v1630 int64
	_ = v1630
	var v1633 int64
	_ = v1633
	var v1643 int64
	_ = v1643
	var v1647 int64
	_ = v1647
	var v1654 int32
	_ = v1654
	var v1655 int32
	_ = v1655
	var v1674 int32
	_ = v1674
	var v1676 int32
	_ = v1676
	var v1677 int32
	_ = v1677
	var v1686 int32
	_ = v1686
	var v1689 int32
	_ = v1689
	var v1690 int32
	_ = v1690
	var v1691 int64
	_ = v1691
	var v1698 int32
	_ = v1698
	var v1702 int64
	_ = v1702
	v4 = int64(0)
	v20 = m.G0
	v22 = v20 - int32(800)
	m.G0 = v22
	v25 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
	if v25 == int32(0) {
		v1702 = v4
		goto L1
	} else {
		goto L2
	}
L1:
	;
	m.G0 = v22 + int32(800)
	return v1702
L2:
	;
	v28 = *(*int64)(unsafe.Add(mBase, uint32(l1)+40))
	v29 = *(*int32)(unsafe.Add(mBase, uint32(l1)+36))
	v30 = int64(*(*uint32)(unsafe.Add(mBase, uint32(l1)+48)))
	if v25 != int32(1) {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v36 = v22 + int32(40)
	v37 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v36)+4)) = v37
	*(*int32)(unsafe.Add(mBase, uint32(v36))) = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v36)+20)) = int32(128)
	v43 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v36)+12)) = v43
	*(*int64)(unsafe.Add(mBase, uint32(v36)+296)) = v43
	*(*int64)(unsafe.Add(mBase, uint32(v36)+160)) = int64(137438953472)
	*(*int32)(unsafe.Add(mBase, uint32(v36)+8)) = v22 + int32(64)
	*(*int32)(unsafe.Add(mBase, uint32(v36)+156)) = v22 + int32(208)
	goto L6
L4:
	;
	v33 = *(*int64)(unsafe.Add(mBase, uint32(l0)+8))
	if base.Ui64(v33) <= base.Ui64(v30) {
		v1702 = v4
		goto L1
	} else {
		goto L5
	}
L5:
	;
	goto L3
L6:
	;
	v58 = int32(0)
	v60 = F_raxSeek(m, v22+int32(40), int32(_a4), v58, v58)
	mBase = m.M
	v63 = m.ExcPending
	if v63 != 0 {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	return int64(0)
L8:
	;
	v67 = F_raxNext(m, v22+int32(40))
	mBase = m.M
	v68 = m.ExcPending
	if v68 != 0 {
		goto L7
	} else {
		goto L10
	}
L9:
	;
	F_raxStop(m, v22+int32(40))
	mBase = m.M
	v1629 = m.ExcPending
	if v1629 != 0 {
		goto L7
	} else {
		goto L330
	}
L10:
	;
	if v67 == int32(0) {
		v1610 = int64(0)
		goto L9
	} else {
		goto L11
	}
L11:
	;
	v72 = base.B2i32(v25 != int32(1))
	v88 = int64(0)
	goto L13
L12:
	;
	F__serverAssert(m, int32(_a1529), int32(_a1525), int32(282))
	mBase = m.M
	v1605 = m.ExcPending
	if v1605 != 0 {
		goto L7
	} else {
		goto L329
	}
L13:
	;
	if v25 != int32(1) {
		goto L15
	} else {
		goto L16
	}
L14:
	;
	F__serverAssert(m, int32(_a1529), int32(_a1525), int32(282))
	mBase = m.M
	v1599 = m.ExcPending
	if v1599 != 0 {
		goto L7
	} else {
		goto L328
	}
L15:
	;
	v99 = *(*int32)(unsafe.Add(mBase, uint32(v22)+52))
	v100 = F_lpFirst(m, v99)
	mBase = m.M
	v101 = m.ExcPending
	if v101 != 0 {
		goto L7
	} else {
		goto L20
	}
L16:
	;
	v97 = *(*int64)(unsafe.Add(mBase, uint32(l0)+8))
	if base.Ui64(v30) < base.Ui64(v97) {
		goto L15
	} else {
		goto L17
	}
L17:
	;
	v1610 = v88
	goto L9
L18:
	;
	v256 = v255 + v88
	if v28 == int64(0) {
		goto L51
	} else {
		goto L52
	}
L19:
	;
	v108 = *(*int32)(unsafe.Add(mBase, uint32(v22)+352))
	v110 = v22 + int32(16)
	v111 = int32(0)
	if base.Ui32(v108+int32(-21)) < base.Ui32(int32(-20)) {
		v245 = v111
		goto L24
	} else {
		goto L25
	}
L20:
	;
	v105 = F_lpGet(m, v100, v22+int32(352), int32(0))
	mBase = m.M
	v106 = m.ExcPending
	if v106 != 0 {
		goto L7
	} else {
		goto L21
	}
L21:
	;
	if v105 != 0 {
		goto L19
	} else {
		goto L22
	}
L22:
	;
	v107 = *(*int64)(unsafe.Add(mBase, uint32(v22)+352))
	v255 = v107
	goto L18
L23:
	;
	if v245 == int32(0) {
		goto L12
	} else {
		goto L50
	}
L24:
	;
	goto L23
L25:
	;
	v123 = int32(1)
	v124 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v105))))
	if v108 != v123 {
		goto L28
	} else {
		goto L29
	}
L26:
	;
	v245 = int32(1)
	goto L24
L27:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v110))) = v226
	goto L26
L28:
	;
	if v124&int32(255) == int32(45) {
		goto L33
	} else {
		goto L34
	}
L29:
	;
	v128 = v124 + int32(-48)
	if base.Ui32(int32(9)) < base.Ui32(v128&int32(255)) {
		v245 = v111
		goto L24
	} else {
		goto L30
	}
L30:
	;
	if v110 == int32(0) {
		goto L26
	} else {
		goto L31
	}
L31:
	;
	v226 = base.I64_extend_i32_u(v128) & int64(255)
	goto L27
L32:
	;
	if base.Ui32(int32(8)) < base.Ui32((v147+int32(-49))&int32(255)) {
		v245 = v111
		goto L24
	} else {
		goto L35
	}
L33:
	;
	v144 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v105)+1)))
	v146 = int32(2)
	v147 = v144
	v148 = v105 + int32(1)
	goto L32
L34:
	;
	v146 = v123
	v147 = v124
	v148 = v105
	goto L32
L35:
	;
	v159 = base.I64_extend_i32_u(v147+int32(-48)) & int64(255)
	if base.Ui32(v108) <= base.Ui32(v146) {
		v202 = v159
		goto L36
	} else {
		goto L37
	}
L36:
	;
	if v124&int32(255) != int32(45) {
		goto L44
	} else {
		goto L45
	}
L37:
	;
	v165 = v146
	v167 = v159
	v169 = v148
	goto L38
L38:
	;
	v171 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v169)+1)))
	if base.Ui32((v171+int32(-58))&int32(255)) < base.Ui32(int32(246)) {
		v245 = v111
		goto L24
	} else {
		goto L40
	}
L39:
	;
	v202 = v192
	goto L36
L40:
	;
	if base.Ui64(int64(1844674407370955161)) < base.Ui64(v167) {
		v245 = v111
		goto L24
	} else {
		goto L41
	}
L41:
	;
	v181 = v167 * int64(10)
	v186 = base.I64_extend_i32_u(v171+int32(-48)) & int64(255)
	if base.Ui64(v186^int64(-1)) < base.Ui64(v181) {
		v245 = v111
		goto L24
	} else {
		goto L42
	}
L42:
	;
	v190 = int32(1)
	v192 = v181 + v186
	v194 = v165 + v190
	if v194 != v108 {
		v165 = v194
		v167 = v192
		v169 = v169 + v190
		goto L38
	} else {
		goto L43
	}
L43:
	;
	goto L39
L44:
	;
	if v202 < int64(0) {
		v245 = v111
		goto L24
	} else {
		goto L48
	}
L45:
	;
	if base.Ui64(int64(-9223372036854775807-1)) < base.Ui64(v202) {
		v245 = v111
		goto L24
	} else {
		goto L46
	}
L46:
	;
	if v110 == int32(0) {
		goto L26
	} else {
		goto L47
	}
L47:
	;
	v226 = int64(0) - v202
	goto L27
L48:
	;
	if v110 == int32(0) {
		goto L26
	} else {
		goto L49
	}
L49:
	;
	v226 = v202
	goto L27
L50:
	;
	v254 = *(*int64)(unsafe.Add(mBase, uint32(v22)+16))
	v255 = v254
	goto L18
L51:
	;
	v260 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v22+int32(360)))) = v260
	*(*int64)(unsafe.Add(mBase, uint32(v22)+352)) = v260
	if v25 != int32(1) {
		goto L55
	} else {
		goto L56
	}
L52:
	;
	if v256 <= v28 {
		goto L51
	} else {
		goto L53
	}
L53:
	;
	v1610 = v88
	goto L9
L54:
	;
	if v368 != 0 {
		goto L62
	} else {
		goto L63
	}
L55:
	;
	v267 = *(*int32)(unsafe.Add(mBase, uint32(v22)+48))
	v268 = *(*int64)(unsafe.Add(mBase, uint32(v267)))
	v269 = *(*int64)(unsafe.Add(mBase, uint32(v267)+8))
	v270 = int64(56)
	v272 = int64(65280)
	v274 = int64(40)
	v277 = int64(16711680)
	v279 = int64(24)
	v281 = int64(4278190080)
	v283 = int64(8)
	*(*int64)(unsafe.Add(mBase, uint32(v22)+360)) = v269<<(uint(v270)%64) | v269&v272<<(uint(v274)%64) | (v269&v277<<(uint(v279)%64) | v269&v281<<(uint(v283)%64)) | (int64(base.Ui64(v269)>>(uint(v283)%64))&v281 | int64(base.Ui64(v269)>>(uint(v279)%64))&v277 | (int64(base.Ui64(v269)>>(uint(v274)%64))&v272 | int64(base.Ui64(v269)>>(uint(v270)%64))))
	*(*int64)(unsafe.Add(mBase, uint32(v22)+352)) = v268<<(uint(v270)%64) | v268&v272<<(uint(v274)%64) | (v268&v277<<(uint(v279)%64) | v268&v281<<(uint(v283)%64)) | (int64(base.Ui64(v268)>>(uint(v283)%64))&v281 | int64(base.Ui64(v268)>>(uint(v279)%64))&v277 | (int64(base.Ui64(v268)>>(uint(v274)%64))&v272 | int64(base.Ui64(v268)>>(uint(v270)%64))))
	v346 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v22+int32(24)))) = v346
	*(*int64)(unsafe.Add(mBase, uint32(v22)+16)) = v346
	v350 = int32(0)
	v356 = F_lpGetEdgeStreamID(m, v99, v350, v22+int32(352), v22+int32(16))
	mBase = m.M
	v357 = m.ExcPending
	if v357 != 0 {
		goto L7
	} else {
		goto L57
	}
L56:
	;
	v264 = *(*int64)(unsafe.Add(mBase, uint32(l0)+8))
	v368 = base.B2i32(base.Ui64(v30) <= base.Ui64(v264-v255))
	goto L54
L57:
	;
	v358 = *(*int64)(unsafe.Add(mBase, uint32(v22)+16))
	v359 = *(*int64)(unsafe.Add(mBase, uint32(l1)+56))
	if base.Ui64(v359) < base.Ui64(v358) {
		v368 = v350
		goto L54
	} else {
		goto L58
	}
L58:
	;
	if base.Ui64(v358) < base.Ui64(v359) {
		v368 = int32(1)
		goto L54
	} else {
		goto L59
	}
L59:
	;
	v364 = *(*int64)(unsafe.Add(mBase, uint32(v22)+24))
	v365 = *(*int64)(unsafe.Add(mBase, uint32(l1)+64))
	if base.Ui64(v365) < base.Ui64(v364) {
		v368 = int32(0)
		goto L54
	} else {
		goto L60
	}
L60:
	;
	v368 = base.B2i32(base.Ui64(v364) < base.Ui64(v365))
	goto L54
L61:
	;
	goto L14
L62:
	;
	F_lpFree(m, v99)
	mBase = m.M
	v1574 = m.ExcPending
	if v1574 != 0 {
		goto L7
	} else {
		goto L323
	}
L63:
	;
	if v29 == int32(0) {
		goto L64
	} else {
		goto L65
	}
L64:
	;
	v373 = F_lpNext(m, v99, v100)
	mBase = m.M
	v374 = m.ExcPending
	if v374 != 0 {
		goto L7
	} else {
		goto L68
	}
L65:
	;
	v1610 = v88
	goto L9
L66:
	;
	v532 = F_lpNext(m, v99, v375)
	mBase = m.M
	v533 = m.ExcPending
	if v533 != 0 {
		goto L7
	} else {
		goto L100
	}
L67:
	;
	v383 = *(*int32)(unsafe.Add(mBase, uint32(v22)+16))
	v385 = v22 + int32(344)
	v386 = int32(0)
	if base.Ui32(v383+int32(-21)) < base.Ui32(int32(-20)) {
		v520 = v386
		goto L73
	} else {
		goto L74
	}
L68:
	;
	v375 = F_lpNext(m, v99, v373)
	mBase = m.M
	v376 = m.ExcPending
	if v376 != 0 {
		goto L7
	} else {
		goto L69
	}
L69:
	;
	v380 = F_lpGet(m, v375, v22+int32(16), int32(0))
	mBase = m.M
	v381 = m.ExcPending
	if v381 != 0 {
		goto L7
	} else {
		goto L70
	}
L70:
	;
	if v380 != 0 {
		goto L67
	} else {
		goto L71
	}
L71:
	;
	v382 = *(*int64)(unsafe.Add(mBase, uint32(v22)+16))
	v530 = v382
	goto L66
L72:
	;
	if v520 == int32(0) {
		goto L61
	} else {
		goto L99
	}
L73:
	;
	goto L72
L74:
	;
	v398 = int32(1)
	v399 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v380))))
	if v383 != v398 {
		goto L77
	} else {
		goto L78
	}
L75:
	;
	v520 = int32(1)
	goto L73
L76:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v385))) = v501
	goto L75
L77:
	;
	if v399&int32(255) == int32(45) {
		goto L82
	} else {
		goto L83
	}
L78:
	;
	v403 = v399 + int32(-48)
	if base.Ui32(int32(9)) < base.Ui32(v403&int32(255)) {
		v520 = v386
		goto L73
	} else {
		goto L79
	}
L79:
	;
	if v385 == int32(0) {
		goto L75
	} else {
		goto L80
	}
L80:
	;
	v501 = base.I64_extend_i32_u(v403) & int64(255)
	goto L76
L81:
	;
	if base.Ui32(int32(8)) < base.Ui32((v422+int32(-49))&int32(255)) {
		v520 = v386
		goto L73
	} else {
		goto L84
	}
L82:
	;
	v419 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v380)+1)))
	v421 = int32(2)
	v422 = v419
	v423 = v380 + int32(1)
	goto L81
L83:
	;
	v421 = v398
	v422 = v399
	v423 = v380
	goto L81
L84:
	;
	v434 = base.I64_extend_i32_u(v422+int32(-48)) & int64(255)
	if base.Ui32(v383) <= base.Ui32(v421) {
		v477 = v434
		goto L85
	} else {
		goto L86
	}
L85:
	;
	if v399&int32(255) != int32(45) {
		goto L93
	} else {
		goto L94
	}
L86:
	;
	v440 = v421
	v442 = v434
	v444 = v423
	goto L87
L87:
	;
	v446 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v444)+1)))
	if base.Ui32((v446+int32(-58))&int32(255)) < base.Ui32(int32(246)) {
		v520 = v386
		goto L73
	} else {
		goto L89
	}
L88:
	;
	v477 = v467
	goto L85
L89:
	;
	if base.Ui64(int64(1844674407370955161)) < base.Ui64(v442) {
		v520 = v386
		goto L73
	} else {
		goto L90
	}
L90:
	;
	v456 = v442 * int64(10)
	v461 = base.I64_extend_i32_u(v446+int32(-48)) & int64(255)
	if base.Ui64(v461^int64(-1)) < base.Ui64(v456) {
		v520 = v386
		goto L73
	} else {
		goto L91
	}
L91:
	;
	v465 = int32(1)
	v467 = v456 + v461
	v469 = v440 + v465
	if v469 != v383 {
		v440 = v469
		v442 = v467
		v444 = v444 + v465
		goto L87
	} else {
		goto L92
	}
L92:
	;
	goto L88
L93:
	;
	if v477 < int64(0) {
		v520 = v386
		goto L73
	} else {
		goto L97
	}
L94:
	;
	if base.Ui64(int64(-9223372036854775807-1)) < base.Ui64(v477) {
		v520 = v386
		goto L73
	} else {
		goto L95
	}
L95:
	;
	if v385 == int32(0) {
		goto L75
	} else {
		goto L96
	}
L96:
	;
	v501 = int64(0) - v477
	goto L76
L97:
	;
	if v385 == int32(0) {
		goto L75
	} else {
		goto L98
	}
L98:
	;
	v501 = v477
	goto L76
L99:
	;
	v529 = *(*int64)(unsafe.Add(mBase, uint32(v22)+344))
	v530 = v529
	goto L66
L100:
	;
	if v530 <= int64(0) {
		v573 = v532
		goto L101
	} else {
		goto L102
	}
L101:
	;
	v579 = F_lpNext(m, v99, v573)
	mBase = m.M
	v580 = m.ExcPending
	if v580 != 0 {
		goto L7
	} else {
		goto L107
	}
L102:
	;
	v549 = v532
	v551 = int64(0)
	goto L103
L103:
	;
	v555 = F_lpNext(m, v99, v549)
	mBase = m.M
	v556 = m.ExcPending
	if v556 != 0 {
		goto L7
	} else {
		goto L105
	}
L104:
	;
	v573 = v555
	goto L101
L105:
	;
	v558 = v551 + int64(1)
	if v558 != v530 {
		v549 = v555
		v551 = v558
		goto L103
	} else {
		goto L106
	}
L106:
	;
	goto L104
L107:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+36)) = v579
	if v579 != 0 {
		goto L111
	} else {
		goto L112
	}
L108:
	;
	F__serverAssert(m, int32(_a1529), int32(_a1525), int32(282))
	mBase = m.M
	v1571 = m.ExcPending
	if v1571 != 0 {
		goto L7
	} else {
		goto L322
	}
L109:
	;
	F__serverAssert(m, int32(_a1529), int32(_a1525), int32(282))
	mBase = m.M
	v1565 = m.ExcPending
	if v1565 != 0 {
		goto L7
	} else {
		goto L321
	}
L110:
	;
	v1379 = F_lpFirst(m, v1371)
	mBase = m.M
	v1380 = m.ExcPending
	if v1380 != 0 {
		goto L7
	} else {
		goto L283
	}
L111:
	;
	v597 = v99
	v599 = v579
	v603 = int64(0)
	goto L113
L112:
	;
	v1371 = v99
	v1377 = int64(0)
	goto L110
L113:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+12)) = v599
	v609 = F_lpGet(m, v599, v22+int32(16), int32(0))
	mBase = m.M
	v610 = m.ExcPending
	if v610 != 0 {
		goto L7
	} else {
		goto L121
	}
L114:
	;
	v1371 = v1357
	v1377 = v1359
	goto L110
L115:
	;
	if v759&int64(2) == int64(0) {
		goto L231
	} else {
		goto L232
	}
L116:
	;
	F__serverAssert(m, int32(_a1529), int32(_a1525), int32(282))
	mBase = m.M
	v1112 = m.ExcPending
	if v1112 != 0 {
		goto L7
	} else {
		goto L229
	}
L117:
	;
	F__serverAssert(m, int32(_a1529), int32(_a1525), int32(282))
	mBase = m.M
	v1106 = m.ExcPending
	if v1106 != 0 {
		goto L7
	} else {
		goto L228
	}
L118:
	;
	F__serverAssert(m, int32(_a1529), int32(_a1525), int32(282))
	mBase = m.M
	v1100 = m.ExcPending
	if v1100 != 0 {
		goto L7
	} else {
		goto L227
	}
L119:
	;
	v760 = F_lpNext(m, v597, v599)
	mBase = m.M
	v761 = m.ExcPending
	if v761 != 0 {
		goto L7
	} else {
		goto L151
	}
L120:
	;
	v612 = *(*int32)(unsafe.Add(mBase, uint32(v22)+16))
	v614 = v22 + int32(344)
	v615 = int32(0)
	if base.Ui32(v612+int32(-21)) < base.Ui32(int32(-20)) {
		v749 = v615
		goto L124
	} else {
		goto L125
	}
L121:
	;
	if v609 != 0 {
		goto L120
	} else {
		goto L122
	}
L122:
	;
	v611 = *(*int64)(unsafe.Add(mBase, uint32(v22)+16))
	v759 = v611
	goto L119
L123:
	;
	if v749 == int32(0) {
		goto L118
	} else {
		goto L150
	}
L124:
	;
	goto L123
L125:
	;
	v627 = int32(1)
	v628 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v609))))
	if v612 != v627 {
		goto L128
	} else {
		goto L129
	}
L126:
	;
	v749 = int32(1)
	goto L124
L127:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v614))) = v730
	goto L126
L128:
	;
	if v628&int32(255) == int32(45) {
		goto L133
	} else {
		goto L134
	}
L129:
	;
	v632 = v628 + int32(-48)
	if base.Ui32(int32(9)) < base.Ui32(v632&int32(255)) {
		v749 = v615
		goto L124
	} else {
		goto L130
	}
L130:
	;
	if v614 == int32(0) {
		goto L126
	} else {
		goto L131
	}
L131:
	;
	v730 = base.I64_extend_i32_u(v632) & int64(255)
	goto L127
L132:
	;
	if base.Ui32(int32(8)) < base.Ui32((v651+int32(-49))&int32(255)) {
		v749 = v615
		goto L124
	} else {
		goto L135
	}
L133:
	;
	v648 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v609)+1)))
	v650 = int32(2)
	v651 = v648
	v652 = v609 + int32(1)
	goto L132
L134:
	;
	v650 = v627
	v651 = v628
	v652 = v609
	goto L132
L135:
	;
	v663 = base.I64_extend_i32_u(v651+int32(-48)) & int64(255)
	if base.Ui32(v612) <= base.Ui32(v650) {
		v706 = v663
		goto L136
	} else {
		goto L137
	}
L136:
	;
	if v628&int32(255) != int32(45) {
		goto L144
	} else {
		goto L145
	}
L137:
	;
	v669 = v650
	v671 = v663
	v673 = v652
	goto L138
L138:
	;
	v675 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v673)+1)))
	if base.Ui32((v675+int32(-58))&int32(255)) < base.Ui32(int32(246)) {
		v749 = v615
		goto L124
	} else {
		goto L140
	}
L139:
	;
	v706 = v696
	goto L136
L140:
	;
	if base.Ui64(int64(1844674407370955161)) < base.Ui64(v671) {
		v749 = v615
		goto L124
	} else {
		goto L141
	}
L141:
	;
	v685 = v671 * int64(10)
	v690 = base.I64_extend_i32_u(v675+int32(-48)) & int64(255)
	if base.Ui64(v690^int64(-1)) < base.Ui64(v685) {
		v749 = v615
		goto L124
	} else {
		goto L142
	}
L142:
	;
	v694 = int32(1)
	v696 = v685 + v690
	v698 = v669 + v694
	if v698 != v612 {
		v669 = v698
		v671 = v696
		v673 = v673 + v694
		goto L138
	} else {
		goto L143
	}
L143:
	;
	goto L139
L144:
	;
	if v706 < int64(0) {
		v749 = v615
		goto L124
	} else {
		goto L148
	}
L145:
	;
	if base.Ui64(int64(-9223372036854775807-1)) < base.Ui64(v706) {
		v749 = v615
		goto L124
	} else {
		goto L146
	}
L146:
	;
	if v614 == int32(0) {
		goto L126
	} else {
		goto L147
	}
L147:
	;
	v730 = int64(0) - v706
	goto L127
L148:
	;
	if v614 == int32(0) {
		goto L126
	} else {
		goto L149
	}
L149:
	;
	v730 = v706
	goto L127
L150:
	;
	v758 = *(*int64)(unsafe.Add(mBase, uint32(v22)+344))
	v759 = v758
	goto L119
L151:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+36)) = v760
	v766 = F_lpGet(m, v760, v22+int32(16), int32(0))
	mBase = m.M
	v767 = m.ExcPending
	if v767 != 0 {
		goto L7
	} else {
		goto L154
	}
L152:
	;
	v917 = F_lpNext(m, v597, v760)
	mBase = m.M
	v918 = m.ExcPending
	if v918 != 0 {
		goto L7
	} else {
		goto L184
	}
L153:
	;
	v769 = *(*int32)(unsafe.Add(mBase, uint32(v22)+16))
	v771 = v22 + int32(344)
	v772 = int32(0)
	if base.Ui32(v769+int32(-21)) < base.Ui32(int32(-20)) {
		v906 = v772
		goto L157
	} else {
		goto L158
	}
L154:
	;
	if v766 != 0 {
		goto L153
	} else {
		goto L155
	}
L155:
	;
	v768 = *(*int64)(unsafe.Add(mBase, uint32(v22)+16))
	v916 = v768
	goto L152
L156:
	;
	if v906 == int32(0) {
		goto L117
	} else {
		goto L183
	}
L157:
	;
	goto L156
L158:
	;
	v784 = int32(1)
	v785 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v766))))
	if v769 != v784 {
		goto L161
	} else {
		goto L162
	}
L159:
	;
	v906 = int32(1)
	goto L157
L160:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v771))) = v887
	goto L159
L161:
	;
	if v785&int32(255) == int32(45) {
		goto L166
	} else {
		goto L167
	}
L162:
	;
	v789 = v785 + int32(-48)
	if base.Ui32(int32(9)) < base.Ui32(v789&int32(255)) {
		v906 = v772
		goto L157
	} else {
		goto L163
	}
L163:
	;
	if v771 == int32(0) {
		goto L159
	} else {
		goto L164
	}
L164:
	;
	v887 = base.I64_extend_i32_u(v789) & int64(255)
	goto L160
L165:
	;
	if base.Ui32(int32(8)) < base.Ui32((v808+int32(-49))&int32(255)) {
		v906 = v772
		goto L157
	} else {
		goto L168
	}
L166:
	;
	v805 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v766)+1)))
	v807 = int32(2)
	v808 = v805
	v809 = v766 + int32(1)
	goto L165
L167:
	;
	v807 = v784
	v808 = v785
	v809 = v766
	goto L165
L168:
	;
	v820 = base.I64_extend_i32_u(v808+int32(-48)) & int64(255)
	if base.Ui32(v769) <= base.Ui32(v807) {
		v863 = v820
		goto L169
	} else {
		goto L170
	}
L169:
	;
	if v785&int32(255) != int32(45) {
		goto L177
	} else {
		goto L178
	}
L170:
	;
	v826 = v807
	v828 = v820
	v830 = v809
	goto L171
L171:
	;
	v832 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v830)+1)))
	if base.Ui32((v832+int32(-58))&int32(255)) < base.Ui32(int32(246)) {
		v906 = v772
		goto L157
	} else {
		goto L173
	}
L172:
	;
	v863 = v853
	goto L169
L173:
	;
	if base.Ui64(int64(1844674407370955161)) < base.Ui64(v828) {
		v906 = v772
		goto L157
	} else {
		goto L174
	}
L174:
	;
	v842 = v828 * int64(10)
	v847 = base.I64_extend_i32_u(v832+int32(-48)) & int64(255)
	if base.Ui64(v847^int64(-1)) < base.Ui64(v842) {
		v906 = v772
		goto L157
	} else {
		goto L175
	}
L175:
	;
	v851 = int32(1)
	v853 = v842 + v847
	v855 = v826 + v851
	if v855 != v769 {
		v826 = v855
		v828 = v853
		v830 = v830 + v851
		goto L171
	} else {
		goto L176
	}
L176:
	;
	goto L172
L177:
	;
	if v863 < int64(0) {
		v906 = v772
		goto L157
	} else {
		goto L181
	}
L178:
	;
	if base.Ui64(int64(-9223372036854775807-1)) < base.Ui64(v863) {
		v906 = v772
		goto L157
	} else {
		goto L179
	}
L179:
	;
	if v771 == int32(0) {
		goto L159
	} else {
		goto L180
	}
L180:
	;
	v887 = int64(0) - v863
	goto L160
L181:
	;
	if v771 == int32(0) {
		goto L159
	} else {
		goto L182
	}
L182:
	;
	v887 = v863
	goto L160
L183:
	;
	v915 = *(*int64)(unsafe.Add(mBase, uint32(v22)+344))
	v916 = v915
	goto L152
L184:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+36)) = v917
	v923 = F_lpGet(m, v917, v22+int32(16), int32(0))
	mBase = m.M
	v924 = m.ExcPending
	if v924 != 0 {
		goto L7
	} else {
		goto L187
	}
L185:
	;
	v1074 = F_lpNext(m, v597, v917)
	mBase = m.M
	v1075 = m.ExcPending
	if v1075 != 0 {
		goto L7
	} else {
		goto L217
	}
L186:
	;
	v926 = *(*int32)(unsafe.Add(mBase, uint32(v22)+16))
	v928 = v22 + int32(344)
	v929 = int32(0)
	if base.Ui32(v926+int32(-21)) < base.Ui32(int32(-20)) {
		v1063 = v929
		goto L190
	} else {
		goto L191
	}
L187:
	;
	if v923 != 0 {
		goto L186
	} else {
		goto L188
	}
L188:
	;
	v925 = *(*int64)(unsafe.Add(mBase, uint32(v22)+16))
	v1073 = v925
	goto L185
L189:
	;
	if v1063 == int32(0) {
		goto L116
	} else {
		goto L216
	}
L190:
	;
	goto L189
L191:
	;
	v941 = int32(1)
	v942 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v923))))
	if v926 != v941 {
		goto L194
	} else {
		goto L195
	}
L192:
	;
	v1063 = int32(1)
	goto L190
L193:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v928))) = v1044
	goto L192
L194:
	;
	if v942&int32(255) == int32(45) {
		goto L199
	} else {
		goto L200
	}
L195:
	;
	v946 = v942 + int32(-48)
	if base.Ui32(int32(9)) < base.Ui32(v946&int32(255)) {
		v1063 = v929
		goto L190
	} else {
		goto L196
	}
L196:
	;
	if v928 == int32(0) {
		goto L192
	} else {
		goto L197
	}
L197:
	;
	v1044 = base.I64_extend_i32_u(v946) & int64(255)
	goto L193
L198:
	;
	if base.Ui32(int32(8)) < base.Ui32((v965+int32(-49))&int32(255)) {
		v1063 = v929
		goto L190
	} else {
		goto L201
	}
L199:
	;
	v962 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v923)+1)))
	v964 = int32(2)
	v965 = v962
	v966 = v923 + int32(1)
	goto L198
L200:
	;
	v964 = v941
	v965 = v942
	v966 = v923
	goto L198
L201:
	;
	v977 = base.I64_extend_i32_u(v965+int32(-48)) & int64(255)
	if base.Ui32(v926) <= base.Ui32(v964) {
		v1020 = v977
		goto L202
	} else {
		goto L203
	}
L202:
	;
	if v942&int32(255) != int32(45) {
		goto L210
	} else {
		goto L211
	}
L203:
	;
	v983 = v964
	v985 = v977
	v987 = v966
	goto L204
L204:
	;
	v989 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v987)+1)))
	if base.Ui32((v989+int32(-58))&int32(255)) < base.Ui32(int32(246)) {
		v1063 = v929
		goto L190
	} else {
		goto L206
	}
L205:
	;
	v1020 = v1010
	goto L202
L206:
	;
	if base.Ui64(int64(1844674407370955161)) < base.Ui64(v985) {
		v1063 = v929
		goto L190
	} else {
		goto L207
	}
L207:
	;
	v999 = v985 * int64(10)
	v1004 = base.I64_extend_i32_u(v989+int32(-48)) & int64(255)
	if base.Ui64(v1004^int64(-1)) < base.Ui64(v999) {
		v1063 = v929
		goto L190
	} else {
		goto L208
	}
L208:
	;
	v1008 = int32(1)
	v1010 = v999 + v1004
	v1012 = v983 + v1008
	if v1012 != v926 {
		v983 = v1012
		v985 = v1010
		v987 = v987 + v1008
		goto L204
	} else {
		goto L209
	}
L209:
	;
	goto L205
L210:
	;
	if v1020 < int64(0) {
		v1063 = v929
		goto L190
	} else {
		goto L214
	}
L211:
	;
	if base.Ui64(int64(-9223372036854775807-1)) < base.Ui64(v1020) {
		v1063 = v929
		goto L190
	} else {
		goto L212
	}
L212:
	;
	if v928 == int32(0) {
		goto L192
	} else {
		goto L213
	}
L213:
	;
	v1044 = int64(0) - v1020
	goto L193
L214:
	;
	if v928 == int32(0) {
		goto L192
	} else {
		goto L215
	}
L215:
	;
	v1044 = v1020
	goto L193
L216:
	;
	v1072 = *(*int64)(unsafe.Add(mBase, uint32(v22)+344))
	v1073 = v1072
	goto L185
L217:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+36)) = v1074
	if v25 == int32(2) {
		goto L220
	} else {
		goto L221
	}
L218:
	;
	v1094 = *(*int64)(unsafe.Add(mBase, uint32(l0)+8))
	if base.Ui64(v30) < base.Ui64(v1094) {
		goto L115
	} else {
		goto L226
	}
L219:
	;
	if base.Ui64(v1090) < base.Ui64(v1089) {
		goto L115
	} else {
		goto L224
	}
L220:
	;
	v1082 = *(*int64)(unsafe.Add(mBase, uint32(v22)+352))
	v1083 = v1082 + v916
	v1084 = *(*int64)(unsafe.Add(mBase, uint32(l1)+56))
	if base.Ui64(v1084) < base.Ui64(v1083) {
		v1371 = v597
		v1377 = v603
		goto L110
	} else {
		goto L223
	}
L221:
	;
	if v25 == int32(1) {
		goto L218
	} else {
		goto L222
	}
L222:
	;
	v1079 = *(*int64)(unsafe.Add(mBase, uint32(l1)+56))
	v1080 = int64(0)
	v1088 = v1080
	v1089 = v1079
	v1090 = v1080
	goto L219
L223:
	;
	v1086 = *(*int64)(unsafe.Add(mBase, uint32(v22)+360))
	v1088 = v1086 + v1073
	v1089 = v1084
	v1090 = v1083
	goto L219
L224:
	;
	v1092 = *(*int64)(unsafe.Add(mBase, uint32(l1)+64))
	if base.Ui64(v1092) <= base.Ui64(v1088) {
		v1371 = v597
		v1377 = v603
		goto L110
	} else {
		goto L225
	}
L225:
	;
	goto L115
L226:
	;
	v1371 = v597
	v1377 = v603
	goto L110
L227:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L228:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L229:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L230:
	;
	if v1282 == int64(0) {
		v1324 = v1281
		goto L266
	} else {
		goto L267
	}
L231:
	;
	v1124 = F_lpGet(m, v1074, v22+int32(16), int32(0))
	mBase = m.M
	v1125 = m.ExcPending
	if v1125 != 0 {
		goto L7
	} else {
		goto L235
	}
L232:
	;
	v1281 = v1074
	v1282 = v530
	goto L230
L233:
	;
	v1275 = F_lpNext(m, v597, v1074)
	mBase = m.M
	v1276 = m.ExcPending
	if v1276 != 0 {
		goto L7
	} else {
		goto L265
	}
L234:
	;
	v1127 = *(*int32)(unsafe.Add(mBase, uint32(v22)+16))
	v1129 = v22 + int32(344)
	v1130 = int32(0)
	if base.Ui32(v1127+int32(-21)) < base.Ui32(int32(-20)) {
		v1264 = v1130
		goto L238
	} else {
		goto L239
	}
L235:
	;
	if v1124 != 0 {
		goto L234
	} else {
		goto L236
	}
L236:
	;
	v1126 = *(*int64)(unsafe.Add(mBase, uint32(v22)+16))
	v1274 = v1126
	goto L233
L237:
	;
	if v1264 == int32(0) {
		goto L109
	} else {
		goto L264
	}
L238:
	;
	goto L237
L239:
	;
	v1142 = int32(1)
	v1143 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1124))))
	if v1127 != v1142 {
		goto L242
	} else {
		goto L243
	}
L240:
	;
	v1264 = int32(1)
	goto L238
L241:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v1129))) = v1245
	goto L240
L242:
	;
	if v1143&int32(255) == int32(45) {
		goto L247
	} else {
		goto L248
	}
L243:
	;
	v1147 = v1143 + int32(-48)
	if base.Ui32(int32(9)) < base.Ui32(v1147&int32(255)) {
		v1264 = v1130
		goto L238
	} else {
		goto L244
	}
L244:
	;
	if v1129 == int32(0) {
		goto L240
	} else {
		goto L245
	}
L245:
	;
	v1245 = base.I64_extend_i32_u(v1147) & int64(255)
	goto L241
L246:
	;
	if base.Ui32(int32(8)) < base.Ui32((v1166+int32(-49))&int32(255)) {
		v1264 = v1130
		goto L238
	} else {
		goto L249
	}
L247:
	;
	v1163 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1124)+1)))
	v1165 = int32(2)
	v1166 = v1163
	v1167 = v1124 + int32(1)
	goto L246
L248:
	;
	v1165 = v1142
	v1166 = v1143
	v1167 = v1124
	goto L246
L249:
	;
	v1178 = base.I64_extend_i32_u(v1166+int32(-48)) & int64(255)
	if base.Ui32(v1127) <= base.Ui32(v1165) {
		v1221 = v1178
		goto L250
	} else {
		goto L251
	}
L250:
	;
	if v1143&int32(255) != int32(45) {
		goto L258
	} else {
		goto L259
	}
L251:
	;
	v1184 = v1165
	v1186 = v1178
	v1188 = v1167
	goto L252
L252:
	;
	v1190 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1188)+1)))
	if base.Ui32((v1190+int32(-58))&int32(255)) < base.Ui32(int32(246)) {
		v1264 = v1130
		goto L238
	} else {
		goto L254
	}
L253:
	;
	v1221 = v1211
	goto L250
L254:
	;
	if base.Ui64(int64(1844674407370955161)) < base.Ui64(v1186) {
		v1264 = v1130
		goto L238
	} else {
		goto L255
	}
L255:
	;
	v1200 = v1186 * int64(10)
	v1205 = base.I64_extend_i32_u(v1190+int32(-48)) & int64(255)
	if base.Ui64(v1205^int64(-1)) < base.Ui64(v1200) {
		v1264 = v1130
		goto L238
	} else {
		goto L256
	}
L256:
	;
	v1209 = int32(1)
	v1211 = v1200 + v1205
	v1213 = v1184 + v1209
	if v1213 != v1127 {
		v1184 = v1213
		v1186 = v1211
		v1188 = v1188 + v1209
		goto L252
	} else {
		goto L257
	}
L257:
	;
	goto L253
L258:
	;
	if v1221 < int64(0) {
		v1264 = v1130
		goto L238
	} else {
		goto L262
	}
L259:
	;
	if base.Ui64(int64(-9223372036854775807-1)) < base.Ui64(v1221) {
		v1264 = v1130
		goto L238
	} else {
		goto L260
	}
L260:
	;
	if v1129 == int32(0) {
		goto L240
	} else {
		goto L261
	}
L261:
	;
	v1245 = int64(0) - v1221
	goto L241
L262:
	;
	if v1129 == int32(0) {
		goto L240
	} else {
		goto L263
	}
L263:
	;
	v1245 = v1221
	goto L241
L264:
	;
	v1273 = *(*int64)(unsafe.Add(mBase, uint32(v22)+344))
	v1274 = v1273
	goto L233
L265:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+36)) = v1275
	v1281 = v1275
	v1282 = v1274 << (uint(int64(1)) % 64)
	goto L230
L266:
	;
	v1330 = F_lpNext(m, v597, v1324)
	mBase = m.M
	v1331 = m.ExcPending
	if v1331 != 0 {
		goto L7
	} else {
		goto L272
	}
L267:
	;
	v1298 = v1281
	v1300 = v1282
	goto L268
L268:
	;
	v1304 = F_lpNext(m, v597, v1298)
	mBase = m.M
	v1305 = m.ExcPending
	if v1305 != 0 {
		goto L7
	} else {
		goto L270
	}
L269:
	;
	v1324 = v1304
	goto L266
L270:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+36)) = v1304
	v1308 = v1300 + int64(-1)
	if v1308 != int64(0) {
		v1298 = v1304
		v1300 = v1308
		goto L268
	} else {
		goto L271
	}
L271:
	;
	goto L269
L272:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+36)) = v1330
	if base.I32_wrap_i64(v759)&int32(1) != 0 {
		v1357 = v597
		v1358 = v1330
		v1359 = v603
		goto L273
	} else {
		goto L274
	}
L273:
	;
	if v1358 != 0 {
		v597 = v1357
		v599 = v1358
		v603 = v1359
		goto L113
	} else {
		goto L282
	}
L274:
	;
	v1340 = F_lpReplaceInteger(m, v597, v22+int32(12), v759|int64(1))
	mBase = m.M
	v1341 = m.ExcPending
	if v1341 != 0 {
		goto L7
	} else {
		goto L275
	}
L275:
	;
	v1342 = *(*int32)(unsafe.Add(mBase, uint32(v22)+36))
	if v1342 != 0 {
		goto L277
	} else {
		goto L278
	}
L276:
	;
	v1350 = *(*int64)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int64)(unsafe.Add(mBase, uint32(l0)+8)) = v1350 + int64(-1)
	v1357 = v1340
	v1358 = v1349
	v1359 = v603 + int64(1)
	goto L273
L277:
	;
	if v1330 != 0 {
		goto L279
	} else {
		goto L280
	}
L278:
	;
	v1349 = int32(0)
	goto L276
L279:
	;
	v1346 = v1330 - v597
	goto L281
L280:
	;
	v1346 = int32(0)
	goto L281
L281:
	;
	v1347 = v1340 + v1346
	*(*int32)(unsafe.Add(mBase, uint32(v22)+36)) = v1347
	v1349 = v1347
	goto L276
L282:
	;
	goto L114
L283:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+36)) = v1379
	v1385 = F_lpReplaceInteger(m, v1371, v22+int32(36), v255-v1377)
	mBase = m.M
	v1386 = m.ExcPending
	if v1386 != 0 {
		goto L7
	} else {
		goto L284
	}
L284:
	;
	v1387 = *(*int32)(unsafe.Add(mBase, uint32(v22)+36))
	v1388 = F_lpNext(m, v1385, v1387)
	mBase = m.M
	v1389 = m.ExcPending
	if v1389 != 0 {
		goto L7
	} else {
		goto L285
	}
L285:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+36)) = v1388
	v1394 = F_lpGet(m, v1388, v22+int32(16), int32(0))
	mBase = m.M
	v1395 = m.ExcPending
	if v1395 != 0 {
		goto L7
	} else {
		goto L288
	}
L286:
	;
	v1549 = F_lpReplaceInteger(m, v1385, v22+int32(36), v1544+v1377)
	mBase = m.M
	v1550 = m.ExcPending
	if v1550 != 0 {
		goto L7
	} else {
		goto L318
	}
L287:
	;
	v1397 = *(*int32)(unsafe.Add(mBase, uint32(v22)+16))
	v1399 = v22 + int32(344)
	v1400 = int32(0)
	if base.Ui32(v1397+int32(-21)) < base.Ui32(int32(-20)) {
		v1534 = v1400
		goto L291
	} else {
		goto L292
	}
L288:
	;
	if v1394 != 0 {
		goto L287
	} else {
		goto L289
	}
L289:
	;
	v1396 = *(*int64)(unsafe.Add(mBase, uint32(v22)+16))
	v1544 = v1396
	goto L286
L290:
	;
	if v1534 == int32(0) {
		goto L108
	} else {
		goto L317
	}
L291:
	;
	goto L290
L292:
	;
	v1412 = int32(1)
	v1413 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1394))))
	if v1397 != v1412 {
		goto L295
	} else {
		goto L296
	}
L293:
	;
	v1534 = int32(1)
	goto L291
L294:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v1399))) = v1515
	goto L293
L295:
	;
	if v1413&int32(255) == int32(45) {
		goto L300
	} else {
		goto L301
	}
L296:
	;
	v1417 = v1413 + int32(-48)
	if base.Ui32(int32(9)) < base.Ui32(v1417&int32(255)) {
		v1534 = v1400
		goto L291
	} else {
		goto L297
	}
L297:
	;
	if v1399 == int32(0) {
		goto L293
	} else {
		goto L298
	}
L298:
	;
	v1515 = base.I64_extend_i32_u(v1417) & int64(255)
	goto L294
L299:
	;
	if base.Ui32(int32(8)) < base.Ui32((v1436+int32(-49))&int32(255)) {
		v1534 = v1400
		goto L291
	} else {
		goto L302
	}
L300:
	;
	v1433 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1394)+1)))
	v1435 = int32(2)
	v1436 = v1433
	v1437 = v1394 + int32(1)
	goto L299
L301:
	;
	v1435 = v1412
	v1436 = v1413
	v1437 = v1394
	goto L299
L302:
	;
	v1448 = base.I64_extend_i32_u(v1436+int32(-48)) & int64(255)
	if base.Ui32(v1397) <= base.Ui32(v1435) {
		v1491 = v1448
		goto L303
	} else {
		goto L304
	}
L303:
	;
	if v1413&int32(255) != int32(45) {
		goto L311
	} else {
		goto L312
	}
L304:
	;
	v1454 = v1435
	v1456 = v1448
	v1458 = v1437
	goto L305
L305:
	;
	v1460 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1458)+1)))
	if base.Ui32((v1460+int32(-58))&int32(255)) < base.Ui32(int32(246)) {
		v1534 = v1400
		goto L291
	} else {
		goto L307
	}
L306:
	;
	v1491 = v1481
	goto L303
L307:
	;
	if base.Ui64(int64(1844674407370955161)) < base.Ui64(v1456) {
		v1534 = v1400
		goto L291
	} else {
		goto L308
	}
L308:
	;
	v1470 = v1456 * int64(10)
	v1475 = base.I64_extend_i32_u(v1460+int32(-48)) & int64(255)
	if base.Ui64(v1475^int64(-1)) < base.Ui64(v1470) {
		v1534 = v1400
		goto L291
	} else {
		goto L309
	}
L309:
	;
	v1479 = int32(1)
	v1481 = v1470 + v1475
	v1483 = v1454 + v1479
	if v1483 != v1397 {
		v1454 = v1483
		v1456 = v1481
		v1458 = v1458 + v1479
		goto L305
	} else {
		goto L310
	}
L310:
	;
	goto L306
L311:
	;
	if v1491 < int64(0) {
		v1534 = v1400
		goto L291
	} else {
		goto L315
	}
L312:
	;
	if base.Ui64(int64(-9223372036854775807-1)) < base.Ui64(v1491) {
		v1534 = v1400
		goto L291
	} else {
		goto L313
	}
L313:
	;
	if v1399 == int32(0) {
		goto L293
	} else {
		goto L314
	}
L314:
	;
	v1515 = int64(0) - v1491
	goto L294
L315:
	;
	if v1399 == int32(0) {
		goto L293
	} else {
		goto L316
	}
L316:
	;
	v1515 = v1491
	goto L294
L317:
	;
	v1543 = *(*int64)(unsafe.Add(mBase, uint32(v22)+344))
	v1544 = v1543
	goto L286
L318:
	;
	v1551 = *(*int32)(unsafe.Add(mBase, uint32(v22)+36))
	v1552 = F_lpNext(m, v1549, v1551)
	mBase = m.M
	v1553 = m.ExcPending
	if v1553 != 0 {
		goto L7
	} else {
		goto L319
	}
L319:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+36)) = v1552
	v1555 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v1556 = *(*int32)(unsafe.Add(mBase, uint32(v22)+48))
	v1557 = *(*int32)(unsafe.Add(mBase, uint32(v22)+56))
	v1559 = F_raxInsert(m, v1555, v1556, v1557, v1549, int32(0))
	mBase = m.M
	v1560 = m.ExcPending
	if v1560 != 0 {
		goto L7
	} else {
		goto L320
	}
L320:
	;
	v1610 = v1377 + v88
	goto L9
L321:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L322:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L323:
	;
	v1575 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v1576 = *(*int32)(unsafe.Add(mBase, uint32(v22)+48))
	v1577 = *(*int32)(unsafe.Add(mBase, uint32(v22)+56))
	v1579 = F_raxRemove(m, v1575, v1576, v1577, int32(0))
	mBase = m.M
	v1580 = m.ExcPending
	if v1580 != 0 {
		goto L7
	} else {
		goto L324
	}
L324:
	;
	v1584 = *(*int32)(unsafe.Add(mBase, uint32(v22)+48))
	v1585 = *(*int32)(unsafe.Add(mBase, uint32(v22)+56))
	v1586 = F_raxSeek(m, v22+int32(40), int32(_a235), v1584, v1585)
	mBase = m.M
	v1587 = m.ExcPending
	if v1587 != 0 {
		goto L7
	} else {
		goto L325
	}
L325:
	;
	v1588 = *(*int64)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int64)(unsafe.Add(mBase, uint32(l0)+8)) = v1588 - v255
	v1593 = F_raxNext(m, v22+int32(40))
	mBase = m.M
	v1594 = m.ExcPending
	if v1594 != 0 {
		goto L7
	} else {
		goto L326
	}
L326:
	;
	if v1593 != 0 {
		v88 = v256
		goto L13
	} else {
		goto L327
	}
L327:
	;
	v1610 = v256
	goto L9
L328:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L329:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L330:
	;
	v1630 = *(*int64)(unsafe.Add(mBase, uint32(l0)+8))
	if v1630 != int64(0) {
		goto L331
	} else {
		goto L332
	}
L331:
	;
	if v1610 == int64(0) {
		v1702 = v1610
		goto L1
	} else {
		goto L333
	}
L332:
	;
	v1633 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(l0)+32)) = v1633
	*(*int64)(unsafe.Add(mBase, uint32(l0+int32(40)))) = v1633
	v1702 = v1610
	goto L1
L333:
	;
	v1643 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v22+int32(416)))) = v1643
	v1647 = int64(-1)
	*(*int64)(unsafe.Add(mBase, uint32(v22+int32(432)))) = v1647
	*(*int64)(unsafe.Add(mBase, uint32(v22)+408)) = v1643
	*(*int64)(unsafe.Add(mBase, uint32(v22)+424)) = v1647
	v1654 = v22 + int32(440)
	v1655 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v1654)+4)) = v1655
	*(*int32)(unsafe.Add(mBase, uint32(v1654))) = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v1654)+20)) = int32(128)
	*(*int64)(unsafe.Add(mBase, uint32(v1654)+12)) = v1643
	*(*int64)(unsafe.Add(mBase, uint32(v1654)+296)) = v1643
	*(*int64)(unsafe.Add(mBase, uint32(v1654)+160)) = int64(137438953472)
	*(*int32)(unsafe.Add(mBase, uint32(v1654)+8)) = v22 + int32(464)
	*(*int32)(unsafe.Add(mBase, uint32(v1654)+156)) = v22 + int32(608)
	goto L334
L334:
	;
	v1674 = int32(0)
	v1676 = F_raxSeek(m, v1654, int32(_a4), v1674, v1674)
	mBase = m.M
	v1677 = m.ExcPending
	if v1677 != 0 {
		goto L7
	} else {
		goto L335
	}
L335:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v22)+744)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v22)+352)) = l0
	*(*int64)(unsafe.Add(mBase, uint32(v22)+396)) = int64(4294967296)
	v1686 = l0 + int32(32)
	v1689 = F_streamIteratorGetID(m, v22+int32(352), v1686, v22+int32(16))
	mBase = m.M
	v1690 = m.ExcPending
	if v1690 != 0 {
		goto L7
	} else {
		goto L337
	}
L336:
	;
	F_raxStop(m, v1654)
	mBase = m.M
	v1698 = m.ExcPending
	if v1698 != 0 {
		goto L7
	} else {
		goto L339
	}
L337:
	;
	if v1689 != 0 {
		goto L336
	} else {
		goto L338
	}
L338:
	;
	v1691 = int64(-1)
	*(*int64)(unsafe.Add(mBase, uint32(v1686))) = v1691
	*(*int64)(unsafe.Add(mBase, uint32(l0+int32(40)))) = v1691
	goto L336
L339:
	;
	v1702 = v1610
	goto L1
}
