package p0

import (
	base "github.com/shibukawa/vkmem/internal/aot/vkaot/base"
	"unsafe"
)

func F_createZsetListpackObject(m *base.Module) int32 {
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
			*(*int64)(unsafe.Add(mBase, uint32(v18))) = int64(34359738371)
			v23 = *(*int32)(unsafe.Add(mBase, uint32(v18)))
			*(*int32)(unsafe.Add(mBase, uint32(v18))) = v23&int32(-241) | int32(176)
			m.G0 = v6 + int32(16)
			return v18
		}
	}
}
func F_zsetAdd(m *base.Module, l0 int32, l1 float64, l2 int32, l3 int32, l4 int32, l5 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
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
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v58 float64
	_ = v58
	var v61 float64
	_ = v61
	var v67 int32
	_ = v67
	var v72 float64
	_ = v72
	var v73 int32
	_ = v73
	var v88 int32
	_ = v88
	var v96 int32
	_ = v96
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v120 int32
	_ = v120
	var v124 int32
	_ = v124
	var v131 int32
	_ = v131
	var v134 int32
	_ = v134
	var v137 int32
	_ = v137
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v143 int32
	_ = v143
	var v146 int32
	_ = v146
	var v152 int32
	_ = v152
	var v159 int32
	_ = v159
	var v162 int32
	_ = v162
	var v165 int32
	_ = v165
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v171 int32
	_ = v171
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v184 int32
	_ = v184
	var v188 int32
	_ = v188
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v207 int32
	_ = v207
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
	var v220 int32
	_ = v220
	var v221 float64
	_ = v221
	var v225 float64
	_ = v225
	var v231 int32
	_ = v231
	var v236 float64
	_ = v236
	var v237 int32
	_ = v237
	var v252 int32
	_ = v252
	var v260 int32
	_ = v260
	var v261 int32
	_ = v261
	var v262 int32
	_ = v262
	var v266 int32
	_ = v266
	var v270 int32
	_ = v270
	var v271 int32
	_ = v271
	var v281 int32
	_ = v281
	var v290 int64
	_ = v290
	var v294 int64
	_ = v294
	var v295 int64
	_ = v295
	var v303 int64
	_ = v303
	var v305 int64
	_ = v305
	var v309 int32
	_ = v309
	var v313 int64
	_ = v313
	var v316 int64
	_ = v316
	var v318 int64
	_ = v318
	var v321 int64
	_ = v321
	var v332 int64
	_ = v332
	var v335 int64
	_ = v335
	var v346 int64
	_ = v346
	var v349 int64
	_ = v349
	var v362 int64
	_ = v362
	var v369 int64
	_ = v369
	var v374 int32
	_ = v374
	var v377 int64
	_ = v377
	var v383 int32
	_ = v383
	var v384 int32
	_ = v384
	var v387 int32
	_ = v387
	var v388 int32
	_ = v388
	var v393 int64
	_ = v393
	var v401 int64
	_ = v401
	var v404 int64
	_ = v404
	var v416 int64
	_ = v416
	var v418 int32
	_ = v418
	var v420 int64
	_ = v420
	var v427 int32
	_ = v427
	var v428 int32
	_ = v428
	var v431 int32
	_ = v431
	var v432 int32
	_ = v432
	var v437 int64
	_ = v437
	var v445 int64
	_ = v445
	var v448 int64
	_ = v448
	var v459 int32
	_ = v459
	var v460 int32
	_ = v460
	var v462 int64
	_ = v462
	var v470 int64
	_ = v470
	var v472 int64
	_ = v472
	var v477 int64
	_ = v477
	var v486 int32
	_ = v486
	var v487 int64
	_ = v487
	var v498 int64
	_ = v498
	var v503 int64
	_ = v503
	var v508 int64
	_ = v508
	var v511 int64
	_ = v511
	var v514 int32
	_ = v514
	var v520 int32
	_ = v520
	var v521 int32
	_ = v521
	var v522 int32
	_ = v522
	var v523 int32
	_ = v523
	var v524 int32
	_ = v524
	var v525 int32
	_ = v525
	var v526 int32
	_ = v526
	var v527 int32
	_ = v527
	var v530 int32
	_ = v530
	var v537 int32
	_ = v537
	var v538 int32
	_ = v538
	var v547 int32
	_ = v547
	var v549 int32
	_ = v549
	var v550 int32
	_ = v550
	var v559 int32
	_ = v559
	var v574 int32
	_ = v574
	v7 = int32(0)
	v15 = m.G0
	v17 = v15 - int32(16)
	m.G0 = v17
	*(*int32)(unsafe.Add(mBase, uint32(l4))) = v7
	if base.Ui64(base.I64_reinterpret_f64(l1)&int64(9223372036854775807)) < base.Ui64(int64(9218868437227405313)) {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	F__serverAssert(m, int32(_a1610), int32(_a1609), int32(1595))
	mBase = m.M
	v574 = m.ExcPending
	if v574 != 0 {
		goto L9
	} else {
		goto L105
	}
L2:
	;
	m.G0 = v17 + int32(16)
	return v559
L3:
	;
	v30 = l3 & int32(16)
	v32 = l3 & int32(8)
	v34 = l3 & int32(4)
	v36 = l3 & int32(2)
	v38 = l3 & int32(1)
	v39 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v39&int32(240) != int32(176) {
		v202 = v39
		goto L6
	} else {
		goto L7
	}
L4:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l4))) = int32(2)
	v559 = v7
	goto L2
L5:
	;
	v549 = int32(1)
	v550 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
	*(*int32)(unsafe.Add(mBase, uint32(l4))) = v550 | v549
	v559 = v549
	goto L2
L6:
	;
	if v202&int32(240) != int32(112) {
		goto L58
	} else {
		goto L59
	}
L7:
	;
	v44 = F_objectGetVal(m, l0)
	mBase = m.M
	v45 = F_zzlFind(m, v44, l2, v17)
	mBase = m.M
	v48 = m.ExcPending
	if v48 != 0 {
		goto L9
	} else {
		goto L10
	}
L8:
	;
	if v34 != 0 {
		goto L5
	} else {
		goto L28
	}
L9:
	;
	return int32(0)
L10:
	;
	if v45 == int32(0) {
		goto L8
	} else {
		goto L11
	}
L11:
	;
	if v36 == int32(0) {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	v58 = *(*float64)(unsafe.Add(mBase, uint32(v17)))
	if v38 == int32(0) {
		v72 = l1
		goto L14
	} else {
		goto L15
	}
L13:
	;
	v53 = int32(1)
	v54 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
	*(*int32)(unsafe.Add(mBase, uint32(l4))) = v54 | v53
	v559 = v53
	goto L2
L14:
	;
	v73 = int32(1)
	if base.B2i32(v30 == int32(0))|(base.F64_ge(v72, v58)^v73) != v73 {
		goto L18
	} else {
		goto L19
	}
L15:
	;
	v61 = base.F64_add(l1, v58)
	if base.Ui64(base.I64_reinterpret_f64(v61)&int64(9223372036854775807)) < base.Ui64(int64(9218868437227405313)) {
		v72 = v61
		goto L14
	} else {
		goto L16
	}
L16:
	;
	v67 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
	*(*int32)(unsafe.Add(mBase, uint32(l4))) = v67 | int32(2)
	v559 = int32(0)
	goto L2
L17:
	;
	if l5 == int32(0) {
		goto L21
	} else {
		goto L22
	}
L18:
	;
	v88 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
	*(*int32)(unsafe.Add(mBase, uint32(l4))) = v88 | int32(1)
	v559 = v73
	goto L2
L19:
	;
	if base.B2i32(v32 == int32(0))|(base.F64_le(v72, v58)^int32(1)) != 0 {
		goto L17
	} else {
		goto L20
	}
L20:
	;
	goto L18
L21:
	;
	if base.F64_eq(v72, v58) != 0 {
		v559 = v73
		goto L2
	} else {
		goto L23
	}
L22:
	;
	*(*float64)(unsafe.Add(mBase, uint32(l5))) = v72
	goto L21
L23:
	;
	v96 = F_objectGetVal(m, l0)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v17)+12)) = v45
	v101 = F_lpDeleteRangeWithEntry(m, v96, v17+int32(12), int32(2))
	mBase = m.M
	v102 = m.ExcPending
	if v102 != 0 {
		goto L9
	} else {
		goto L24
	}
L24:
	;
	F_objectSetVal(m, l0, v101)
	mBase = m.M
	v104 = m.ExcPending
	if v104 != 0 {
		goto L9
	} else {
		goto L25
	}
L25:
	;
	v105 = F_objectGetVal(m, l0)
	mBase = m.M
	v106 = F_zzlInsert(m, v105, l2, v72)
	mBase = m.M
	v107 = m.ExcPending
	if v107 != 0 {
		goto L9
	} else {
		goto L26
	}
L26:
	;
	F_objectSetVal(m, l0, v106)
	mBase = m.M
	v109 = m.ExcPending
	if v109 != 0 {
		goto L9
	} else {
		goto L27
	}
L27:
	;
	v110 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
	*(*int32)(unsafe.Add(mBase, uint32(l4))) = v110 | int32(8)
	v559 = v73
	goto L2
L28:
	;
	v114 = F_objectGetVal(m, l0)
	mBase = m.M
	v115 = F_lpLength(m, v114)
	mBase = m.M
	v116 = m.ExcPending
	if v116 != 0 {
		goto L9
	} else {
		goto L30
	}
L29:
	;
	v195 = F_zsetLength(m, l0)
	mBase = m.M
	v196 = m.ExcPending
	if v196 != 0 {
		goto L9
	} else {
		goto L56
	}
L30:
	;
	v120 = *(*int32)(unsafe.Add(mBase, _consts[1103]))
	if base.Ui32(v120) <= base.Ui32(int32(base.Ui32(v115)>>(uint(int32(1))%32))) {
		goto L29
	} else {
		goto L31
	}
L31:
	;
	v124 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2+int32(-1)))))
	switch v124 & int32(7) {
	case 0:
		goto L38
	case 1:
		goto L37
	case 2:
		goto L36
	case 3:
		goto L35
	case 4:
		goto L34
	default:
		goto L32
	}
L32:
	;
	v146 = F_objectGetVal(m, l0)
	mBase = m.M
	v152 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2+int32(-1)))))
	switch v152 & int32(7) {
	case 0:
		goto L46
	case 1:
		goto L45
	case 2:
		goto L44
	case 3:
		goto L43
	case 4:
		goto L42
	default:
		v169 = int32(0)
		goto L41
	}
L33:
	;
	v143 = *(*int32)(unsafe.Add(mBase, _consts[1104]))
	if base.Ui32(v143) < base.Ui32(v141) {
		goto L29
	} else {
		goto L39
	}
L34:
	;
	v140 = *(*int32)(unsafe.Add(mBase, uint32(l2+int32(-17))))
	v141 = v140
	goto L33
L35:
	;
	v137 = *(*int32)(unsafe.Add(mBase, uint32(l2+int32(-9))))
	v141 = v137
	goto L33
L36:
	;
	v134 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l2+int32(-5)))))
	v141 = v134
	goto L33
L37:
	;
	v131 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2+int32(-3)))))
	v141 = v131
	goto L33
L38:
	;
	v141 = int32(base.Ui32(v124) >> (uint(int32(3)) % 32))
	goto L33
L39:
	;
	goto L32
L40:
	;
	if v146 != 0 {
		goto L49
	} else {
		goto L50
	}
L41:
	;
	v171 = v169
	goto L40
L42:
	;
	v168 = *(*int32)(unsafe.Add(mBase, uint32(l2+int32(-17))))
	v169 = v168
	goto L41
L43:
	;
	v165 = *(*int32)(unsafe.Add(mBase, uint32(l2+int32(-9))))
	v171 = v165
	goto L40
L44:
	;
	v162 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l2+int32(-5)))))
	v171 = v162
	goto L40
L45:
	;
	v159 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2+int32(-3)))))
	v171 = v159
	goto L40
L46:
	;
	v171 = int32(base.Ui32(v152) >> (uint(int32(3)) % 32))
	goto L40
L47:
	;
	if base.B2i32(base.Ui32(v174+v171) < base.Ui32(int32(1073741825))) == int32(0) {
		goto L29
	} else {
		goto L51
	}
L48:
	;
	goto L47
L49:
	;
	v173 = *(*int32)(unsafe.Add(mBase, uint32(v146)))
	v174 = v173
	goto L48
L50:
	;
	v174 = int32(0)
	goto L48
L51:
	;
	v180 = F_objectGetVal(m, l0)
	mBase = m.M
	v181 = F_zzlInsert(m, v180, l2, l1)
	mBase = m.M
	v182 = m.ExcPending
	if v182 != 0 {
		goto L9
	} else {
		goto L52
	}
L52:
	;
	F_objectSetVal(m, l0, v181)
	mBase = m.M
	v184 = m.ExcPending
	if v184 != 0 {
		goto L9
	} else {
		goto L53
	}
L53:
	;
	if l5 == int32(0) {
		goto L54
	} else {
		goto L55
	}
L54:
	;
	v188 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
	*(*int32)(unsafe.Add(mBase, uint32(l4))) = v188 | int32(4)
	v559 = int32(1)
	goto L2
L55:
	;
	*(*float64)(unsafe.Add(mBase, uint32(l5))) = l1
	goto L54
L56:
	;
	F_zsetConvertAndExpand(m, l0, int32(7), v195+int32(1))
	mBase = m.M
	v200 = m.ExcPending
	if v200 != 0 {
		goto L9
	} else {
		goto L57
	}
L57:
	;
	v201 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v202 = v201
	goto L6
L58:
	;
	F__serverPanic_1(m, int32(_a1609), int32(1604), int32(_a854), int32(0))
	mBase = m.M
	v547 = m.ExcPending
	if v547 != 0 {
		goto L9
	} else {
		goto L104
	}
L59:
	;
	v207 = F_objectGetVal(m, l0)
	mBase = m.M
	v208 = *(*int32)(unsafe.Add(mBase, uint32(v207)))
	v209 = F_hashtableFindRef(m, v208, l2)
	mBase = m.M
	v210 = m.ExcPending
	if v210 != 0 {
		goto L9
	} else {
		goto L61
	}
L60:
	;
	if v34 != 0 {
		goto L78
	} else {
		goto L79
	}
L61:
	;
	if v209 == int32(0) {
		goto L60
	} else {
		goto L62
	}
L62:
	;
	if v36 == int32(0) {
		goto L63
	} else {
		goto L64
	}
L63:
	;
	v220 = *(*int32)(unsafe.Add(mBase, uint32(v209)))
	v221 = *(*float64)(unsafe.Add(mBase, uint32(v220)))
	*(*float64)(unsafe.Add(mBase, uint32(v17))) = v221
	if v38 == int32(0) {
		v236 = l1
		goto L65
	} else {
		goto L66
	}
L64:
	;
	v215 = int32(1)
	v216 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
	*(*int32)(unsafe.Add(mBase, uint32(l4))) = v216 | v215
	v559 = v215
	goto L2
L65:
	;
	v237 = int32(1)
	if base.B2i32(v30 == int32(0))|(base.F64_ge(v236, v221)^v237) != v237 {
		goto L69
	} else {
		goto L70
	}
L66:
	;
	v225 = base.F64_add(l1, v221)
	if base.Ui64(base.I64_reinterpret_f64(v225)&int64(9223372036854775807)) < base.Ui64(int64(9218868437227405313)) {
		v236 = v225
		goto L65
	} else {
		goto L67
	}
L67:
	;
	v231 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
	*(*int32)(unsafe.Add(mBase, uint32(l4))) = v231 | int32(2)
	v559 = int32(0)
	goto L2
L68:
	;
	if l5 == int32(0) {
		goto L72
	} else {
		goto L73
	}
L69:
	;
	v252 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
	*(*int32)(unsafe.Add(mBase, uint32(l4))) = v252 | int32(1)
	v559 = v237
	goto L2
L70:
	;
	if base.B2i32(v32 == int32(0))|(base.F64_le(v236, v221)^int32(1)) != 0 {
		goto L68
	} else {
		goto L71
	}
L71:
	;
	goto L69
L72:
	;
	if base.F64_eq(v236, v221) != 0 {
		v559 = v237
		goto L2
	} else {
		goto L74
	}
L73:
	;
	*(*float64)(unsafe.Add(mBase, uint32(l5))) = v236
	goto L72
L74:
	;
	v260 = *(*int32)(unsafe.Add(mBase, uint32(v207)+4))
	v261 = F_zslUpdateScore(m, v260, v220, v236)
	mBase = m.M
	v262 = m.ExcPending
	if v262 != 0 {
		goto L9
	} else {
		goto L76
	}
L75:
	;
	v266 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
	*(*int32)(unsafe.Add(mBase, uint32(l4))) = v266 | int32(8)
	v559 = v237
	goto L2
L76:
	;
	if v261 == int32(0) {
		goto L75
	} else {
		goto L77
	}
L77:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v209))) = v261
	goto L75
L78:
	;
	v537 = int32(1)
	v538 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
	*(*int32)(unsafe.Add(mBase, uint32(l4))) = v538 | v537
	v559 = v537
	goto L2
L79:
	;
	v270 = int32(1)
	v271 = *(*int32)(unsafe.Add(mBase, uint32(v207)+4))
	v281 = *(*int32)(unsafe.Add(mBase, _consts[1094]))
	if int32(311) < v281 {
		goto L82
	} else {
		goto L83
	}
L80:
	;
	v514 = int32(1)
	if v511 == int64(0) {
		goto L96
	} else {
		goto L97
	}
L81:
	;
	*(*int32)(unsafe.Add(mBase, _consts[1094])) = v486
	v498 = int64(base.Ui64(v487)>>(uint(int64(29))%64))&int64(22906492245) ^ v487
	v503 = v498<<(uint(int64(17))%64)&int64(8202884508482404352) ^ v498
	v508 = v503<<(uint(int64(37))%64)&int64(-2270628950310912) ^ v503
	v511 = int64(base.Ui64(v508)>>(uint(int64(43))%64)) ^ v508
	goto L80
L82:
	;
	if v281 == int32(313) {
		goto L85
	} else {
		goto L86
	}
L83:
	;
	v290 = *(*int64)(unsafe.Add(mBase, uint32(v281<<(uint(int32(3))%32))+uint32(_consts[1095])))
	v486 = v281 + int32(1)
	v487 = v290
	goto L81
L84:
	;
	v374 = int32(0)
	v377 = v369
	goto L90
L85:
	;
	v295 = int64(5489)
	*(*int64)(unsafe.Add(mBase, _consts[1095])) = v295
	v303 = int64(1)
	v305 = v295
	goto L87
L86:
	;
	v294 = *(*int64)(unsafe.Add(mBase, _consts[1095]))
	v369 = v294
	goto L84
L87:
	;
	v309 = int32(3)
	v313 = int64(62)
	v316 = int64(6364136223846793005)
	v318 = (int64(base.Ui64(v305)>>(uint(v313)%64))^v305)*v316 + v303
	*(*int64)(unsafe.Add(mBase, uint32(base.I32_wrap_i64(v303)<<(uint(v309)%32))+uint32(_consts[1095]))) = v318
	v321 = v303 + int64(1)
	v332 = (int64(base.Ui64(v318)>>(uint(v313)%64))^v318)*v316 + v321
	*(*int64)(unsafe.Add(mBase, uint32(base.I32_wrap_i64(v321)<<(uint(v309)%32))+uint32(_consts[1095]))) = v332
	v335 = v303 + int64(2)
	v346 = (int64(base.Ui64(v332)>>(uint(v313)%64))^v332)*v316 + v335
	*(*int64)(unsafe.Add(mBase, uint32(base.I32_wrap_i64(v335)<<(uint(v309)%32))+uint32(_consts[1095]))) = v346
	v349 = v303 + int64(3)
	if v349 == int64(312) {
		v369 = v295
		goto L84
	} else {
		goto L89
	}
L89:
	;
	v362 = (int64(base.Ui64(v346)>>(uint(int64(62))%64))^v346)*int64(6364136223846793005) + v349
	*(*int64)(unsafe.Add(mBase, uint32(base.I32_wrap_i64(v349)<<(uint(int32(3))%32))+uint32(_consts[1095]))) = v362
	v303 = v303 + int64(4)
	v305 = v362
	goto L87
L90:
	;
	v383 = int32(3)
	v384 = v374 << (uint(v383) % 32)
	v387 = int32(1)
	v388 = v374 + v387
	v393 = *(*int64)(unsafe.Add(mBase, uint32(v388<<(uint(v383)%32))+uint32(_consts[1095])))
	v401 = *(*int64)(unsafe.Add(mBase, uint32(base.I32_wrap_i64(v393)&v387<<(uint(v383)%32))+uint32(_consts[1096])))
	v404 = *(*int64)(unsafe.Add(mBase, uint32(v384)+uint32(_consts[1097])))
	*(*int64)(unsafe.Add(mBase, uint32(v384)+uint32(_consts[1095]))) = v401 ^ v404 ^ int64(base.Ui64(v377&int64(-2147483648)|v393&int64(2147483646))>>(uint(int64(1))%64))
	if v388 != int32(156) {
		v374 = v388
		v377 = v393
		goto L90
	} else {
		goto L92
	}
L91:
	;
	v416 = *(*int64)(unsafe.Add(mBase, _consts[1097]))
	v418 = int32(156)
	v420 = v416
	goto L93
L92:
	;
	goto L91
L93:
	;
	v427 = int32(3)
	v428 = v418 << (uint(v427) % 32)
	v431 = int32(1)
	v432 = v418 + v431
	v437 = *(*int64)(unsafe.Add(mBase, uint32(v432<<(uint(v427)%32))+uint32(_consts[1095])))
	v445 = *(*int64)(unsafe.Add(mBase, uint32(base.I32_wrap_i64(v437)&v431<<(uint(v427)%32))+uint32(_consts[1096])))
	v448 = *(*int64)(unsafe.Add(mBase, uint32(v428)+uint32(_consts[1098])))
	*(*int64)(unsafe.Add(mBase, uint32(v428)+uint32(_consts[1095]))) = v445 ^ v448 ^ int64(base.Ui64(v420&int64(-2147483648)|v437&int64(2147483646))>>(uint(int64(1))%64))
	if v432 != int32(311) {
		v418 = v432
		v420 = v437
		goto L93
	} else {
		goto L95
	}
L94:
	;
	v459 = int32(1)
	v460 = int32(0)
	v462 = *(*int64)(unsafe.Add(mBase, _consts[1095]))
	v470 = *(*int64)(unsafe.Add(mBase, uint32(base.I32_wrap_i64(v462)&v459<<(uint(int32(3))%32))+uint32(_consts[1096])))
	v472 = *(*int64)(unsafe.Add(mBase, _consts[1099]))
	v477 = *(*int64)(unsafe.Add(mBase, _consts[1100]))
	*(*int64)(unsafe.Add(mBase, _consts[1100])) = v470 ^ v472 ^ int64(base.Ui64(v462&int64(2147483646)|v477&int64(-2147483648))>>(uint(int64(1))%64))
	v486 = v459
	v487 = v462
	goto L81
L95:
	;
	goto L94
L96:
	;
	v520 = int32(32)
	goto L98
L97:
	;
	v520 = int32(base.Ui32(base.I32_wrap_i64(base.I64_clz(v511)))>>(uint(v514)%32)) + v514
	goto L98
L98:
	;
	v521 = F_zslCreateNode(m, v520, l1, l2)
	mBase = m.M
	v522 = m.ExcPending
	if v522 != 0 {
		goto L9
	} else {
		goto L99
	}
L99:
	;
	v523 = F_zslInsertNode(m, v271, v521)
	mBase = m.M
	v524 = m.ExcPending
	if v524 != 0 {
		goto L9
	} else {
		goto L100
	}
L100:
	;
	v525 = *(*int32)(unsafe.Add(mBase, uint32(v207)))
	v526 = F_hashtableAdd(m, v525, v523)
	mBase = m.M
	v527 = m.ExcPending
	if v527 != 0 {
		goto L9
	} else {
		goto L101
	}
L101:
	;
	if v526 == int32(0) {
		goto L1
	} else {
		goto L102
	}
L102:
	;
	v530 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
	*(*int32)(unsafe.Add(mBase, uint32(l4))) = v530 | int32(4)
	if l5 == int32(0) {
		v559 = v270
		goto L2
	} else {
		goto L103
	}
L103:
	;
	*(*float64)(unsafe.Add(mBase, uint32(l5))) = l1
	v559 = v270
	goto L2
L104:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L105:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_zsetInitLexRange(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v25 int32
	_ = v25
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	v7 = int32(1)
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v8 == int32(0) {
		v74 = v7
		return v74
	} else {
		v11 = *(*int32)(unsafe.Add(mBase, uint32(v8)))
		if v11&int32(15) != int32(3) {
			v74 = v7
			return v74
		} else {
			v16 = int32(1)
			v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
			if v17 != v16 {
				*(*int64)(unsafe.Add(mBase, uint32(l0)+80)) = int64(0)
				*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = int32(0)
				v31 = l0 + int32(56)
				v32 = F_zsetParseLexRange(m, l1, l2, v31)
				mBase = m.M
				v33 = m.ExcPending
				if v33 != 0 {
					return int32(0)
				} else {
					if v32 == int32(-1) {
						v74 = v16
						return v74
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = int32(1)
						v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
						v39 = *(*int32)(unsafe.Add(mBase, uint32(v38)))
						switch int32(base.Ui32(v39)>>(uint(int32(4))%32))&int32(15) + int32(-7) {
						case 0:
							v51 = F_objectGetVal(m, v38)
							mBase = m.M
							v52 = *(*int32)(unsafe.Add(mBase, uint32(v51)+4))
							if l3 == int32(0) {
								v58 = F_zslNthInLexRange(m, v52, v31, int32(-1))
								mBase = m.M
								v69 = v58
							} else {
								v56 = F_zslNthInLexRange(m, v52, v31, int32(0))
								mBase = m.M
								v69 = v56
							}
							*(*int32)(unsafe.Add(mBase, uint32(l0)+80)) = v69
							v71 = int32(0)
							if v69 != 0 {
								v74 = v71
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(l0)+84)) = int32(1)
								v74 = v71
							}
							return v74
						default:
							F__serverPanic_1(m, int32(_a694), int32(5197), int32(_a708), int32(0))
							mBase = m.M
							v64 = m.ExcPending
							if v64 != 0 {
								return int32(0)
							} else {
								F_abort(m)
								mBase = m.M
								base.Wasm_trap_unreachable()
								for {
								}
							}
						case 4:
							v46 = F_objectGetVal(m, v38)
							mBase = m.M
							if l3 == int32(0) {
								v66 = F_zzlLastInLexRange(m, v46, v31)
								mBase = m.M
								v67 = m.ExcPending
								if v67 != 0 {
									return int32(0)
								} else {
									v69 = v66
									*(*int32)(unsafe.Add(mBase, uint32(l0)+80)) = v69
									v71 = int32(0)
									if v69 != 0 {
										v74 = v71
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(l0)+84)) = int32(1)
										v74 = v71
									}
									return v74
								}
							} else {
								v49 = F_zzlFirstInLexRange(m, v46, v31)
								mBase = m.M
								v50 = m.ExcPending
								if v50 != 0 {
									return int32(0)
								} else {
									v69 = v49
									*(*int32)(unsafe.Add(mBase, uint32(l0)+80)) = v69
									v71 = int32(0)
									if v69 != 0 {
										v74 = v71
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(l0)+84)) = int32(1)
										v74 = v71
									}
									return v74
								}
							}
						}
					}
				}
			} else {
				F_zsetFreeLexRange(m, l0+int32(56))
				mBase = m.M
				v25 = m.ExcPending
				if v25 != 0 {
					return int32(0)
				} else {
					*(*int64)(unsafe.Add(mBase, uint32(l0)+80)) = int64(0)
					*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = int32(0)
					v31 = l0 + int32(56)
					v32 = F_zsetParseLexRange(m, l1, l2, v31)
					mBase = m.M
					v33 = m.ExcPending
					if v33 != 0 {
						return int32(0)
					} else {
						if v32 == int32(-1) {
							v74 = v16
							return v74
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = int32(1)
							v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
							v39 = *(*int32)(unsafe.Add(mBase, uint32(v38)))
							switch int32(base.Ui32(v39)>>(uint(int32(4))%32))&int32(15) + int32(-7) {
							case 0:
								v51 = F_objectGetVal(m, v38)
								mBase = m.M
								v52 = *(*int32)(unsafe.Add(mBase, uint32(v51)+4))
								if l3 == int32(0) {
									v58 = F_zslNthInLexRange(m, v52, v31, int32(-1))
									mBase = m.M
									v69 = v58
								} else {
									v56 = F_zslNthInLexRange(m, v52, v31, int32(0))
									mBase = m.M
									v69 = v56
								}
								*(*int32)(unsafe.Add(mBase, uint32(l0)+80)) = v69
								v71 = int32(0)
								if v69 != 0 {
									v74 = v71
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(l0)+84)) = int32(1)
									v74 = v71
								}
								return v74
							default:
								F__serverPanic_1(m, int32(_a694), int32(5197), int32(_a708), int32(0))
								mBase = m.M
								v64 = m.ExcPending
								if v64 != 0 {
									return int32(0)
								} else {
									F_abort(m)
									mBase = m.M
									base.Wasm_trap_unreachable()
									for {
									}
								}
							case 4:
								v46 = F_objectGetVal(m, v38)
								mBase = m.M
								if l3 == int32(0) {
									v66 = F_zzlLastInLexRange(m, v46, v31)
									mBase = m.M
									v67 = m.ExcPending
									if v67 != 0 {
										return int32(0)
									} else {
										v69 = v66
										*(*int32)(unsafe.Add(mBase, uint32(l0)+80)) = v69
										v71 = int32(0)
										if v69 != 0 {
											v74 = v71
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(l0)+84)) = int32(1)
											v74 = v71
										}
										return v74
									}
								} else {
									v49 = F_zzlFirstInLexRange(m, v46, v31)
									mBase = m.M
									v50 = m.ExcPending
									if v50 != 0 {
										return int32(0)
									} else {
										v69 = v49
										*(*int32)(unsafe.Add(mBase, uint32(l0)+80)) = v69
										v71 = int32(0)
										if v69 != 0 {
											v74 = v71
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(l0)+84)) = int32(1)
											v74 = v71
										}
										return v74
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
func F_zsetRemoveFromSkiplist(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
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
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v66 int32
	_ = v66
	var v69 float64
	_ = v69
	var v70 float64
	_ = v70
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	var v100 int32
	_ = v100
	var v107 int32
	_ = v107
	var v110 int32
	_ = v110
	var v113 int32
	_ = v113
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v120 int32
	_ = v120
	var v127 int32
	_ = v127
	var v130 int32
	_ = v130
	var v133 int32
	_ = v133
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
	var v143 int32
	_ = v143
	var v146 int32
	_ = v146
	var v148 int32
	_ = v148
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v157 int32
	_ = v157
	var v177 int32
	_ = v177
	var v188 int32
	_ = v188
	var v191 int32
	_ = v191
	var v198 int32
	_ = v198
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v206 int32
	_ = v206
	var v208 int32
	_ = v208
	var v219 int32
	_ = v219
	var v226 int32
	_ = v226
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v233 int32
	_ = v233
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
	var v244 int32
	_ = v244
	var v247 int32
	_ = v247
	var v248 int32
	_ = v248
	var v255 int32
	_ = v255
	var v256 int32
	_ = v256
	var v261 int32
	_ = v261
	var v267 int32
	_ = v267
	var v268 int32
	_ = v268
	var v272 int32
	_ = v272
	var v274 int32
	_ = v274
	var v282 int32
	_ = v282
	var v289 int32
	_ = v289
	var v293 int32
	_ = v293
	var v306 int32
	_ = v306
	var v311 int32
	_ = v311
	var v316 int32
	_ = v316
	var v333 int32
	_ = v333
	v13 = m.G0
	v15 = v13 - int32(144)
	m.G0 = v15
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v21 = F_hashtablePop(m, v18, l1, v15+int32(12))
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	F__serverAssert(m, int32(_a1611), int32(_a1609), int32(350))
	mBase = m.M
	v333 = m.ExcPending
	if v333 != 0 {
		goto L3
	} else {
		goto L68
	}
L2:
	;
	m.G0 = v15 + int32(144)
	return v316
L3:
	;
	return int32(0)
L4:
	;
	if v21 == int32(0) {
		v316 = int32(0)
		goto L2
	} else {
		goto L5
	}
L5:
	;
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v15)+12))
	v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v28)+16))
	if v29 < int32(1) {
		v177 = v28
		goto L6
	} else {
		goto L7
	}
L6:
	;
	v188 = *(*int32)(unsafe.Add(mBase, uint32(v177)+12))
	if v188 != v27 {
		goto L1
	} else {
		goto L44
	}
L7:
	;
	v33 = v27 + int32(16)
	v35 = v28
	v39 = v29
	goto L8
L8:
	;
	v47 = v39 + int32(-1)
	v49 = v47 << (uint(int32(3)) % 32)
	v52 = v35 + v49 + int32(12)
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v52)))
	if v53 == v27 {
		v157 = v35
		goto L10
	} else {
		goto L11
	}
L9:
	;
	v177 = v157
	goto L6
L10:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15+int32(16)+v47<<(uint(int32(2))%32)))) = v157
	if int32(1) < v39 {
		v35 = v157
		v39 = v47
		goto L8
	} else {
		goto L43
	}
L11:
	;
	if v53 == int32(0) {
		v157 = v35
		goto L10
	} else {
		goto L12
	}
L12:
	;
	v57 = v53
	v58 = v35
	v66 = v52
	goto L13
L13:
	;
	if v27 != 0 {
		goto L16
	} else {
		goto L17
	}
L14:
	;
	v157 = v148
	goto L10
L15:
	;
	v153 = v148 + v49 + int32(12)
	v154 = *(*int32)(unsafe.Add(mBase, uint32(v153)))
	if v154 == v27 {
		v157 = v148
		goto L10
	} else {
		goto L41
	}
L16:
	;
	v69 = *(*float64)(unsafe.Add(mBase, uint32(v57)))
	v70 = *(*float64)(unsafe.Add(mBase, uint32(v27)))
	if base.F64_gt(v69, v70) != 0 {
		v157 = v58
		goto L10
	} else {
		goto L18
	}
L17:
	;
	v148 = v57
	goto L15
L18:
	;
	if base.F64_lt(v69, v70) == int32(0) {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	v76 = v57 + int32(16)
	v77 = *(*int32)(unsafe.Add(mBase, uint32(v76)))
	v78 = int32(3)
	v80 = v76 + v77<<(uint(v78)%32)
	v81 = int32(*(*int8)(unsafe.Add(mBase, uint32(v80))))
	v82 = v80 + v81
	v83 = int32(1)
	v85 = *(*int32)(unsafe.Add(mBase, uint32(v33)))
	v88 = v33 + v85<<(uint(v78)%32)
	v89 = int32(*(*int8)(unsafe.Add(mBase, uint32(v88))))
	v90 = v88 + v89
	v93 = int32(0)
	v100 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v82+v93))))
	switch v100 & int32(7) {
	case 0:
		goto L27
	case 1:
		goto L26
	case 2:
		goto L25
	case 3:
		goto L24
	case 4:
		goto L23
	default:
		v117 = v93
		goto L22
	}
L20:
	;
	v148 = v57
	goto L15
L21:
	;
	if int32(-1) < v143 {
		v157 = v58
		goto L10
	} else {
		goto L40
	}
L22:
	;
	v120 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v90+int32(0)))))
	switch v120 & int32(7) {
	case 0:
		goto L33
	case 1:
		goto L32
	case 2:
		goto L31
	case 3:
		goto L30
	case 4:
		goto L29
	default:
		v137 = v93
		goto L28
	}
L23:
	;
	v116 = *(*int32)(unsafe.Add(mBase, uint32(v82+int32(-16))))
	v117 = v116
	goto L22
L24:
	;
	v113 = *(*int32)(unsafe.Add(mBase, uint32(v82+int32(-8))))
	v117 = v113
	goto L22
L25:
	;
	v110 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v82+int32(-4)))))
	v117 = v110
	goto L22
L26:
	;
	v107 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v82+int32(-2)))))
	v117 = v107
	goto L22
L27:
	;
	v117 = int32(base.Ui32(v100) >> (uint(int32(3)) % 32))
	goto L22
L28:
	;
	v138 = base.B2i32(base.Ui32(v117) < base.Ui32(v137))
	if base.Ui32(v117) < base.Ui32(v137) {
		goto L34
	} else {
		goto L35
	}
L29:
	;
	v136 = *(*int32)(unsafe.Add(mBase, uint32(v90+int32(-16))))
	v137 = v136
	goto L28
L30:
	;
	v133 = *(*int32)(unsafe.Add(mBase, uint32(v90+int32(-8))))
	v137 = v133
	goto L28
L31:
	;
	v130 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v90+int32(-4)))))
	v137 = v130
	goto L28
L32:
	;
	v127 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v90+int32(-2)))))
	v137 = v127
	goto L28
L33:
	;
	v137 = int32(base.Ui32(v120) >> (uint(int32(3)) % 32))
	goto L28
L34:
	;
	v139 = v117
	goto L36
L35:
	;
	v139 = v137
	goto L36
L36:
	;
	v140 = F_memcmp(m, v82+v83, v90+v83, v139)
	mBase = m.M
	if v140 != 0 {
		goto L37
	} else {
		goto L38
	}
L37:
	;
	v143 = v140
	goto L39
L38:
	;
	v143 = base.B2i32(base.Ui32(v137) < base.Ui32(v117)) - v138
	goto L39
L39:
	;
	goto L21
L40:
	;
	v146 = *(*int32)(unsafe.Add(mBase, uint32(v66)))
	v148 = v146
	goto L15
L41:
	;
	if v154 != 0 {
		v57 = v154
		v58 = v148
		v66 = v153
		goto L13
	} else {
		goto L42
	}
L42:
	;
	goto L14
L43:
	;
	goto L9
L44:
	;
	v191 = v15 + int32(16)
	v198 = *(*int32)(unsafe.Add(mBase, uint32(v28)+16))
	if v198 < int32(1) {
		v261 = v198
		goto L46
	} else {
		goto L47
	}
L45:
	;
	F_valkey_free(m, v27)
	mBase = m.M
	v311 = m.ExcPending
	if v311 != 0 {
		goto L3
	} else {
		goto L67
	}
L46:
	;
	v267 = *(*int32)(unsafe.Add(mBase, uint32(v27)+8))
	v268 = *(*int32)(unsafe.Add(mBase, uint32(v27)+12))
	if v268 == int32(0) {
		goto L59
	} else {
		goto L60
	}
L47:
	;
	v202 = v27 + int32(12)
	v203 = *(*int32)(unsafe.Add(mBase, uint32(v191)))
	v204 = *(*int32)(unsafe.Add(mBase, uint32(v203)+12))
	if v204 != v27 {
		goto L48
	} else {
		goto L49
	}
L48:
	;
	v208 = int32(1)
	if v198 != v208 {
		goto L50
	} else {
		goto L51
	}
L49:
	;
	v206 = *(*int32)(unsafe.Add(mBase, uint32(v202)))
	*(*int32)(unsafe.Add(mBase, uint32(v203)+12)) = v206
	goto L48
L50:
	;
	v219 = v208
	goto L52
L51:
	;
	v261 = int32(1)
	goto L46
L52:
	;
	v226 = *(*int32)(unsafe.Add(mBase, uint32(v191+v219<<(uint(int32(2))%32))))
	v228 = v219 << (uint(int32(3)) % 32)
	v229 = v226 + v228
	v230 = *(*int32)(unsafe.Add(mBase, uint32(v229)+12))
	if v230 != v27 {
		goto L55
	} else {
		goto L56
	}
L53:
	;
	v261 = v256
	goto L46
L54:
	;
	v255 = v219 + int32(1)
	v256 = *(*int32)(unsafe.Add(mBase, uint32(v28)+16))
	if v255 < v256 {
		v219 = v255
		goto L52
	} else {
		goto L57
	}
L55:
	;
	v247 = v229 + int32(16)
	v248 = *(*int32)(unsafe.Add(mBase, uint32(v247)))
	*(*int32)(unsafe.Add(mBase, uint32(v247))) = v248 + int32(-1)
	goto L54
L56:
	;
	v233 = v229 + int32(16)
	v235 = *(*int32)(unsafe.Add(mBase, uint32(v27+int32(16)+v228)))
	v236 = *(*int32)(unsafe.Add(mBase, uint32(v233)))
	*(*int32)(unsafe.Add(mBase, uint32(v233))) = v235 + v236 + int32(-1)
	v244 = *(*int32)(unsafe.Add(mBase, uint32(v202+v228)))
	*(*int32)(unsafe.Add(mBase, uint32(v229+int32(12)))) = v244
	goto L54
L57:
	;
	goto L53
L58:
	;
	if v274 < int32(2) {
		goto L61
	} else {
		goto L62
	}
L59:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+8)) = v267
	v274 = v261
	goto L58
L60:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v268)+8)) = v267
	v272 = *(*int32)(unsafe.Add(mBase, uint32(v28)+16))
	v274 = v272
	goto L58
L61:
	;
	v306 = *(*int32)(unsafe.Add(mBase, uint32(v28)))
	*(*int32)(unsafe.Add(mBase, uint32(v28))) = v306 + int32(-1)
	goto L45
L62:
	;
	v282 = v274
	goto L63
L63:
	;
	v289 = v282 + int32(-1)
	v293 = *(*int32)(unsafe.Add(mBase, uint32(v28+int32(12)+v289<<(uint(int32(3))%32))))
	if v293 != 0 {
		goto L61
	} else {
		goto L65
	}
L64:
	;
	goto L61
L65:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+16)) = v289
	if base.Ui32(int32(2)) < base.Ui32(v282) {
		v282 = v289
		goto L63
	} else {
		goto L66
	}
L66:
	;
	goto L64
L67:
	;
	v316 = int32(1)
	goto L2
L68:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
}
