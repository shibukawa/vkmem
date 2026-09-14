package p2

import (
	base "github.com/shibukawa/vkmem/internal/aot/vkaot/base"
	"unsafe"
)

func F_addModuleConfigApply(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v30 int32
	_ = v30
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v57 int32
	_ = v57
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	if v9 == int32(0) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	m.G0 = v7 + int32(16)
	return
L2:
	;
	v13 = v7 + int32(8)
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(v13)+4)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v13))) = v14
	goto L3
L3:
	;
	v19 = v7 + int32(8)
	v21 = *(*int32)(unsafe.Add(mBase, uint32(v19)))
	if v21 == int32(0) {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	v63 = F_listAddNodeTail(m, l0, l1)
	mBase = m.M
	v64 = m.ExcPending
	if v64 != 0 {
		goto L18
	} else {
		goto L19
	}
L5:
	;
	if v21 == int32(0) {
		goto L4
	} else {
		goto L8
	}
L6:
	;
	goto L5
L7:
	;
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v19)+4))
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v21+base.B2i32(v24 == int32(0))<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v19))) = v30
	goto L6
L8:
	;
	v37 = v21
	goto L9
L9:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v37)+8))
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v38)+16))
	v40 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	if v39 != v40 {
		goto L11
	} else {
		goto L12
	}
L10:
	;
	goto L4
L11:
	;
	v46 = v7 + int32(8)
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v46)))
	if v48 == int32(0) {
		goto L15
	} else {
		goto L16
	}
L12:
	;
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v38)+4))
	v43 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v42 == v43 {
		goto L1
	} else {
		goto L13
	}
L13:
	;
	goto L11
L14:
	;
	if v48 != 0 {
		v37 = v48
		goto L9
	} else {
		goto L17
	}
L15:
	;
	goto L14
L16:
	;
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v46)+4))
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v48+base.B2i32(v51 == int32(0))<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v46))) = v57
	goto L15
L17:
	;
	goto L10
L18:
	;
	return
L19:
	;
	goto L1
}
func F_checkModuleAuthentication(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v40 int64
	_ = v40
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v73 int64
	_ = v73
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v83 int64
	_ = v83
	var v87 int64
	_ = v87
	var v91 int32
	_ = v91
	var v95 int32
	_ = v95
	var v103 int64
	_ = v103
	var v105 int32
	_ = v105
	var v109 int64
	_ = v109
	var v113 int64
	_ = v113
	var v116 int32
	_ = v116
	var v124 int64
	_ = v124
	var v128 int32
	_ = v128
	var v132 int32
	_ = v132
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v139 int32
	_ = v139
	var v141 int32
	_ = v141
	var v143 int32
	_ = v143
	var v145 int32
	_ = v145
	var v148 int32
	_ = v148
	var v151 int32
	_ = v151
	var v154 int64
	_ = v154
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v189 int64
	_ = v189
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v199 int64
	_ = v199
	var v203 int64
	_ = v203
	var v207 int32
	_ = v207
	var v211 int32
	_ = v211
	var v219 int64
	_ = v219
	var v221 int32
	_ = v221
	var v225 int64
	_ = v225
	var v229 int64
	_ = v229
	var v232 int32
	_ = v232
	var v240 int64
	_ = v240
	var v244 int32
	_ = v244
	var v246 int32
	_ = v246
	var v248 int32
	_ = v248
	var v250 int32
	_ = v250
	var v252 int32
	_ = v252
	var v255 int32
	_ = v255
	var v258 int32
	_ = v258
	var v259 int64
	_ = v259
	var v260 int64
	_ = v260
	var v263 int32
	_ = v263
	var v264 int32
	_ = v264
	var v269 int32
	_ = v269
	var v270 int32
	_ = v270
	var v273 int32
	_ = v273
	var v277 int32
	_ = v277
	var v280 int32
	_ = v280
	var v283 int32
	_ = v283
	var v287 int32
	_ = v287
	var v288 int32
	_ = v288
	var v290 int32
	_ = v290
	var v291 int32
	_ = v291
	var v315 int32
	_ = v315
	var v317 int32
	_ = v317
	var v327 int32
	_ = v327
	var v329 int32
	_ = v329
	var v332 int32
	_ = v332
	var v338 int32
	_ = v338
	var v347 int32
	_ = v347
	var v348 int32
	_ = v348
	var v358 int32
	_ = v358
	var v359 int32
	_ = v359
	var v360 int32
	_ = v360
	var v363 int32
	_ = v363
	var v365 int32
	_ = v365
	var v368 int32
	_ = v368
	var v374 int32
	_ = v374
	var v376 int32
	_ = v376
	var v380 int32
	_ = v380
	var v381 int64
	_ = v381
	var v403 int32
	_ = v403
	var v404 int32
	_ = v404
	var v405 int32
	_ = v405
	var v406 int64
	_ = v406
	var v411 int32
	_ = v411
	var v412 int32
	_ = v412
	var v416 int64
	_ = v416
	var v420 int64
	_ = v420
	var v424 int32
	_ = v424
	var v428 int32
	_ = v428
	var v436 int64
	_ = v436
	var v438 int32
	_ = v438
	var v442 int64
	_ = v442
	var v446 int64
	_ = v446
	var v449 int32
	_ = v449
	var v457 int64
	_ = v457
	var v464 int32
	_ = v464
	var v466 int32
	_ = v466
	var v467 int32
	_ = v467
	var v469 int32
	_ = v469
	var v471 int32
	_ = v471
	var v472 int32
	_ = v472
	var v473 int32
	_ = v473
	var v475 int32
	_ = v475
	var v485 int32
	_ = v485
	var v494 int32
	_ = v494
	var v506 int32
	_ = v506
	var v508 int32
	_ = v508
	var v510 int32
	_ = v510
	var v513 int32
	_ = v513
	var v514 int32
	_ = v514
	var v515 int32
	_ = v515
	var v535 int32
	_ = v535
	var v536 int32
	_ = v536
	var v539 int32
	_ = v539
	var v540 int32
	_ = v540
	var v541 int32
	_ = v541
	var v542 int64
	_ = v542
	var v543 int32
	_ = v543
	var v545 int32
	_ = v545
	var v546 int32
	_ = v546
	var v547 int64
	_ = v547
	var v550 int32
	_ = v550
	var v552 int64
	_ = v552
	var v560 int64
	_ = v560
	var v566 int32
	_ = v566
	var v572 int32
	_ = v572
	v5 = int32(0)
	v17 = m.G0
	v19 = v17 - int32(80)
	m.G0 = v19
	v23 = *(*int32)(unsafe.Add(mBase, _consts[559]))
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v23)+20))
	if v24 == v5 {
		v572 = int32(2)
		goto L1
	} else {
		goto L2
	}
L1:
	;
	m.G0 = v19 + int32(80)
	return v572
L2:
	;
	v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+108))
	if v28 == int32(0) {
		v283 = int32(1)
		goto L4
	} else {
		goto L5
	}
L3:
	;
	v494 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+200)))
	if v494&int32(16) == int32(0) {
		goto L70
	} else {
		goto L71
	}
L4:
	;
	v287 = int32(0)
	v288 = *(*int32)(unsafe.Add(mBase, _consts[559]))
	v290 = v19 + int32(72)
	v291 = *(*int32)(unsafe.Add(mBase, uint32(v288)))
	*(*int32)(unsafe.Add(mBase, uint32(v290)+4)) = v287
	*(*int32)(unsafe.Add(mBase, uint32(v290))) = v291
	goto L40
L5:
	;
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v28)))
	if v31 == int32(0) {
		v277 = v28
		goto L6
	} else {
		goto L7
	}
L6:
	;
	v280 = *(*int32)(unsafe.Add(mBase, uint32(v277)+4))
	v283 = base.B2i32(v280 == int32(0))
	goto L4
L7:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31))) = l0
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v31)+12))
	if v35 != 0 {
		goto L9
	} else {
		goto L10
	}
L8:
	;
	v145 = *(*int32)(unsafe.Add(mBase, uint32(v31)+28))
	if v145 == int32(0) {
		goto L23
	} else {
		goto L24
	}
L9:
	;
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v31)+4))
	v40 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v19+int32(24)))) = v40
	*(*int64)(unsafe.Add(mBase, uint32(v19+int32(64)))) = v40
	*(*int64)(unsafe.Add(mBase, uint32(v19+int32(56)))) = v40
	*(*int64)(unsafe.Add(mBase, uint32(v19+int32(48)))) = v40
	*(*int64)(unsafe.Add(mBase, uint32(v19+int32(40)))) = v40
	*(*int64)(unsafe.Add(mBase, uint32(v19+int32(32)))) = v40
	*(*int64)(unsafe.Add(mBase, uint32(v19)+16)) = v40
	*(*int32)(unsafe.Add(mBase, uint32(v19)+4)) = v37
	*(*int32)(unsafe.Add(mBase, uint32(v19))) = int32(561)
	*(*int32)(unsafe.Add(mBase, uint32(v19)+28)) = int32(4)
	v70 = *(*int32)(unsafe.Add(mBase, _consts[131]))
	v71 = int32(0)
	v72 = *(*int32)(unsafe.Add(mBase, _consts[23]))
	v73 = m.T0[v72].(func(*base.Module) int64)(m)
	mBase = m.M
	if v70 == v71 {
		goto L12
	} else {
		goto L13
	}
L10:
	;
	v143 = int32(1)
	goto L8
L11:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v19)+56)) = v87
	v91 = int32(0)
	v95 = *(*int32)(unsafe.Add(mBase, _consts[95]))
	*(*int32)(unsafe.Add(mBase, _consts[95])) = v95 + int32(1)
	goto L16
L12:
	;
	v83 = *(*int64)(unsafe.Add(mBase, _consts[529]))
	v87 = v83*int64(1000) + v73
	goto L11
L13:
	;
	v78 = *(*int32)(unsafe.Add(mBase, _consts[219]))
	v79 = base.I32_div_s(int32(1000000), v78)
	v87 = v73 + base.I64_extend_i32_s(v79)
	goto L11
L14:
	;
	v128 = *(*int32)(unsafe.Add(mBase, uint32(v31)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v19)+44)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v19)+40)) = v128
	v132 = *(*int32)(unsafe.Add(mBase, uint32(v31)))
	*(*int32)(unsafe.Add(mBase, uint32(v19)+12)) = v31
	*(*int32)(unsafe.Add(mBase, uint32(v19)+8)) = v132
	v135 = *(*int32)(unsafe.Add(mBase, uint32(v31)+12))
	v136 = m.T0[v135].(func(*base.Module, int32, int32, int32, int32) int32)(m, v19, l1, l2, l3)
	mBase = m.M
	v139 = m.ExcPending
	if v139 != 0 {
		goto L20
	} else {
		goto L21
	}
L15:
	;
	goto L14
L16:
	;
	if v95 != 0 {
		goto L15
	} else {
		goto L17
	}
L17:
	;
	goto L19
L18:
	;
	v105 = int32(0)
	*(*int64)(unsafe.Add(mBase, _consts[96])) = v103
	v109 = base.I64_div_s(v103, int64(1000))
	*(*int64)(unsafe.Add(mBase, _consts[35])) = v109
	v113 = base.I64_div_s(v103, int64(1000000))
	*(*int64)(unsafe.Add(mBase, _consts[47])) = v113
	v116 = *(*int32)(unsafe.Add(mBase, _consts[97]))
	F_lrulfu_updateClockAndPolicy(m, v109, int32(base.Ui32(v116&int32(2))>>(uint(int32(1))%32)))
	mBase = m.M
	v124 = *(*int64)(unsafe.Add(mBase, _consts[35]))
	*(*int64)(unsafe.Add(mBase, _consts[98])) = v124
	goto L15
L19:
	;
	v103 = F_ustime(m)
	mBase = m.M
	goto L18
L20:
	;
	return int32(0)
L21:
	;
	F_moduleFreeContext(m, v19)
	mBase = m.M
	v141 = m.ExcPending
	if v141 != 0 {
		goto L20
	} else {
		goto L22
	}
L22:
	;
	v143 = v136
	goto L8
L23:
	;
	v255 = *(*int32)(unsafe.Add(mBase, uint32(l0)+108))
	*(*int32)(unsafe.Add(mBase, uint32(v255))) = int32(0)
	v258 = *(*int32)(unsafe.Add(mBase, uint32(l0)+72))
	v259 = *(*int64)(unsafe.Add(mBase, uint32(v258)+104))
	v260 = *(*int64)(unsafe.Add(mBase, uint32(v31)+64))
	*(*int64)(unsafe.Add(mBase, uint32(v258)+104)) = v259 + v260
	v263 = *(*int32)(unsafe.Add(mBase, uint32(v31)+4))
	v264 = *(*int32)(unsafe.Add(mBase, uint32(v263)+52))
	*(*int32)(unsafe.Add(mBase, uint32(v263)+52)) = v264 + int32(-1)
	F_valkey_free(m, v31)
	mBase = m.M
	v269 = m.ExcPending
	if v269 != 0 {
		goto L20
	} else {
		goto L37
	}
L24:
	;
	v148 = *(*int32)(unsafe.Add(mBase, uint32(v31)+24))
	if v148 == int32(0) {
		goto L23
	} else {
		goto L25
	}
L25:
	;
	v151 = *(*int32)(unsafe.Add(mBase, uint32(v31)+4))
	v154 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v19+int32(64)))) = v154
	*(*int64)(unsafe.Add(mBase, uint32(v19+int32(56)))) = v154
	*(*int64)(unsafe.Add(mBase, uint32(v19+int32(48)))) = v154
	*(*int64)(unsafe.Add(mBase, uint32(v19+int32(40)))) = v154
	*(*int64)(unsafe.Add(mBase, uint32(v19+int32(32)))) = v154
	*(*int64)(unsafe.Add(mBase, uint32(v19+int32(24)))) = v154
	*(*int64)(unsafe.Add(mBase, uint32(v19+int32(16)))) = v154
	*(*int64)(unsafe.Add(mBase, uint32(v19)+8)) = v154
	*(*int32)(unsafe.Add(mBase, uint32(v19)+4)) = v151
	*(*int32)(unsafe.Add(mBase, uint32(v19))) = int32(561)
	v186 = *(*int32)(unsafe.Add(mBase, _consts[131]))
	v187 = int32(0)
	v188 = *(*int32)(unsafe.Add(mBase, _consts[23]))
	v189 = m.T0[v188].(func(*base.Module) int64)(m)
	mBase = m.M
	if v186 == v187 {
		goto L27
	} else {
		goto L28
	}
L26:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v19)+56)) = v203
	v207 = int32(0)
	v211 = *(*int32)(unsafe.Add(mBase, _consts[95]))
	*(*int32)(unsafe.Add(mBase, _consts[95])) = v211 + int32(1)
	goto L31
L27:
	;
	v199 = *(*int64)(unsafe.Add(mBase, _consts[529]))
	v203 = v199*int64(1000) + v189
	goto L26
L28:
	;
	v194 = *(*int32)(unsafe.Add(mBase, _consts[219]))
	v195 = base.I32_div_s(int32(1000000), v194)
	v203 = v189 + base.I64_extend_i32_s(v195)
	goto L26
L29:
	;
	v244 = *(*int32)(unsafe.Add(mBase, uint32(v31)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v19)+40)) = v244
	v246 = *(*int32)(unsafe.Add(mBase, uint32(v31)))
	*(*int32)(unsafe.Add(mBase, uint32(v19)+8)) = v246
	v248 = *(*int32)(unsafe.Add(mBase, uint32(v31)+24))
	m.T0[v248].(func(*base.Module, int32, int32))(m, v19, v244)
	mBase = m.M
	v250 = m.ExcPending
	if v250 != 0 {
		goto L20
	} else {
		goto L35
	}
L30:
	;
	goto L29
L31:
	;
	if v211 != 0 {
		goto L30
	} else {
		goto L32
	}
L32:
	;
	goto L34
L33:
	;
	v221 = int32(0)
	*(*int64)(unsafe.Add(mBase, _consts[96])) = v219
	v225 = base.I64_div_s(v219, int64(1000))
	*(*int64)(unsafe.Add(mBase, _consts[35])) = v225
	v229 = base.I64_div_s(v219, int64(1000000))
	*(*int64)(unsafe.Add(mBase, _consts[47])) = v229
	v232 = *(*int32)(unsafe.Add(mBase, _consts[97]))
	F_lrulfu_updateClockAndPolicy(m, v225, int32(base.Ui32(v232&int32(2))>>(uint(int32(1))%32)))
	mBase = m.M
	v240 = *(*int64)(unsafe.Add(mBase, _consts[35]))
	*(*int64)(unsafe.Add(mBase, _consts[98])) = v240
	goto L30
L34:
	;
	v219 = F_ustime(m)
	mBase = m.M
	goto L33
L35:
	;
	F_moduleFreeContext(m, v19)
	mBase = m.M
	v252 = m.ExcPending
	if v252 != 0 {
		goto L20
	} else {
		goto L36
	}
L36:
	;
	goto L23
L37:
	;
	v270 = int32(1)
	if v143 != v270 {
		v485 = v143
		goto L3
	} else {
		goto L38
	}
L38:
	;
	v273 = *(*int32)(unsafe.Add(mBase, uint32(l0)+108))
	if v273 == int32(0) {
		v283 = v270
		goto L4
	} else {
		goto L39
	}
L39:
	;
	v277 = v273
	goto L6
L40:
	;
	v315 = v283
	v317 = int32(1)
	goto L41
L41:
	;
	v327 = v19 + int32(72)
	v329 = *(*int32)(unsafe.Add(mBase, uint32(v327)))
	if v329 == int32(0) {
		goto L44
	} else {
		goto L45
	}
L42:
	;
	v485 = int32(0)
	goto L3
L43:
	;
	if v329 == int32(0) {
		v485 = v317
		goto L3
	} else {
		goto L46
	}
L44:
	;
	goto L43
L45:
	;
	v332 = *(*int32)(unsafe.Add(mBase, uint32(v327)+4))
	v338 = *(*int32)(unsafe.Add(mBase, uint32(v329+base.B2i32(v332 == int32(0))<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v327))) = v338
	goto L44
L46:
	;
	v347 = v315
	v348 = v329
	goto L48
L47:
	;
	v376 = *(*int32)(unsafe.Add(mBase, uint32(l0)+204))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+204)) = v376 & int32(-262145)
	v380 = *(*int32)(unsafe.Add(mBase, uint32(v358)))
	v381 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v19+int32(64)))) = v381
	*(*int64)(unsafe.Add(mBase, uint32(v19+int32(56)))) = v381
	*(*int64)(unsafe.Add(mBase, uint32(v19+int32(48)))) = v381
	*(*int64)(unsafe.Add(mBase, uint32(v19+int32(40)))) = v381
	*(*int64)(unsafe.Add(mBase, uint32(v19+int32(32)))) = v381
	*(*int64)(unsafe.Add(mBase, uint32(v19+int32(24)))) = v381
	*(*int64)(unsafe.Add(mBase, uint32(v19+int32(16)))) = v381
	*(*int64)(unsafe.Add(mBase, uint32(v19+int32(8)))) = v381
	*(*int32)(unsafe.Add(mBase, uint32(v19)+4)) = v380
	*(*int32)(unsafe.Add(mBase, uint32(v19))) = int32(561)
	v403 = *(*int32)(unsafe.Add(mBase, _consts[131]))
	v404 = int32(0)
	v405 = *(*int32)(unsafe.Add(mBase, _consts[23]))
	v406 = m.T0[v405].(func(*base.Module) int64)(m)
	mBase = m.M
	if v403 == v404 {
		goto L56
	} else {
		goto L57
	}
L48:
	;
	v358 = *(*int32)(unsafe.Add(mBase, uint32(v348)+8))
	if v347 != 0 {
		goto L47
	} else {
		goto L50
	}
L50:
	;
	v359 = *(*int32)(unsafe.Add(mBase, uint32(l0)+108))
	v360 = *(*int32)(unsafe.Add(mBase, uint32(v359)+4))
	v363 = v19 + int32(72)
	v365 = *(*int32)(unsafe.Add(mBase, uint32(v363)))
	if v365 == int32(0) {
		goto L52
	} else {
		goto L53
	}
L51:
	;
	if v365 != 0 {
		v347 = base.B2i32(v358 == v360)
		v348 = v365
		goto L48
	} else {
		goto L54
	}
L52:
	;
	goto L51
L53:
	;
	v368 = *(*int32)(unsafe.Add(mBase, uint32(v363)+4))
	v374 = *(*int32)(unsafe.Add(mBase, uint32(v365+base.B2i32(v368 == int32(0))<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v363))) = v374
	goto L52
L54:
	;
	v485 = v317
	goto L3
L55:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v19)+56)) = v420
	v424 = int32(0)
	v428 = *(*int32)(unsafe.Add(mBase, _consts[95]))
	*(*int32)(unsafe.Add(mBase, _consts[95])) = v428 + int32(1)
	goto L60
L56:
	;
	v416 = *(*int64)(unsafe.Add(mBase, _consts[529]))
	v420 = v416*int64(1000) + v406
	goto L55
L57:
	;
	v411 = *(*int32)(unsafe.Add(mBase, _consts[219]))
	v412 = base.I32_div_s(int32(1000000), v411)
	v420 = v406 + base.I64_extend_i32_s(v412)
	goto L55
L58:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+8)) = l0
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = int32(0)
	v464 = *(*int32)(unsafe.Add(mBase, uint32(l0)+108))
	if v464 != 0 {
		v469 = v464
		goto L64
	} else {
		goto L65
	}
L59:
	;
	goto L58
L60:
	;
	if v428 != 0 {
		goto L59
	} else {
		goto L61
	}
L61:
	;
	goto L63
L62:
	;
	v438 = int32(0)
	*(*int64)(unsafe.Add(mBase, _consts[96])) = v436
	v442 = base.I64_div_s(v436, int64(1000))
	*(*int64)(unsafe.Add(mBase, _consts[35])) = v442
	v446 = base.I64_div_s(v436, int64(1000000))
	*(*int64)(unsafe.Add(mBase, _consts[47])) = v446
	v449 = *(*int32)(unsafe.Add(mBase, _consts[97]))
	F_lrulfu_updateClockAndPolicy(m, v442, int32(base.Ui32(v449&int32(2))>>(uint(int32(1))%32)))
	mBase = m.M
	v457 = *(*int64)(unsafe.Add(mBase, _consts[35]))
	*(*int64)(unsafe.Add(mBase, _consts[98])) = v457
	goto L59
L63:
	;
	v436 = F_ustime(m)
	mBase = m.M
	goto L62
L64:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v469)+4)) = v358
	v471 = *(*int32)(unsafe.Add(mBase, uint32(v358)+4))
	v472 = m.T0[v471].(func(*base.Module, int32, int32, int32, int32) int32)(m, v19, l1, l2, l3)
	mBase = m.M
	v473 = m.ExcPending
	if v473 != 0 {
		goto L20
	} else {
		goto L67
	}
L65:
	;
	v466 = F_valkey_calloc(m, int32(20))
	mBase = m.M
	v467 = m.ExcPending
	if v467 != 0 {
		goto L20
	} else {
		goto L66
	}
L66:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+108)) = v466
	v469 = v466
	goto L64
L67:
	;
	F_moduleFreeContext(m, v19)
	mBase = m.M
	v475 = m.ExcPending
	if v475 != 0 {
		goto L20
	} else {
		goto L68
	}
L68:
	;
	if v472 != 0 {
		v315 = int32(1)
		v317 = v472
		goto L41
	} else {
		goto L69
	}
L69:
	;
	goto L42
L70:
	;
	v508 = *(*int32)(unsafe.Add(mBase, uint32(l0)+108))
	if v508 != 0 {
		goto L75
	} else {
		goto L76
	}
L71:
	;
	if v485 == int32(0) {
		v572 = int32(3)
		goto L1
	} else {
		goto L72
	}
L72:
	;
	F__serverAssert(m, int32(_a942), int32(_a917), int32(8586))
	mBase = m.M
	v506 = m.ExcPending
	if v506 != 0 {
		goto L20
	} else {
		goto L73
	}
L73:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L74:
	;
	v514 = int32(1)
	v515 = *(*int32)(unsafe.Add(mBase, uint32(l0)+204))
	if v485 != v514 {
		goto L77
	} else {
		goto L78
	}
L75:
	;
	v510 = *(*int32)(unsafe.Add(mBase, uint32(v508)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v508)+4)) = int32(0)
	v513 = v510
	goto L74
L76:
	;
	v513 = int32(0)
	goto L74
L77:
	;
	if v515&int32(262144) == int32(0) {
		v535 = v514
		goto L79
	} else {
		goto L80
	}
L78:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+204)) = v515 & int32(-262145)
	v572 = int32(2)
	goto L1
L79:
	;
	v536 = int32(0)
	if v513 == v536 {
		v541 = v536
		goto L81
	} else {
		goto L82
	}
L80:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+204)) = v515 & int32(-262145)
	v535 = int32(base.Ui32(v515^int32(-1))>>(uint(int32(23))%32)) & int32(1)
	goto L79
L81:
	;
	v542 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	v543 = F_objectGetVal(m, l1)
	mBase = m.M
	v545 = v19 + int32(24)
	v546 = int32(0)
	v547 = *(*int64)(unsafe.Add(mBase, _consts[560]))
	*(*int64)(unsafe.Add(mBase, uint32(v545))) = v547
	v550 = v19 + int32(16)
	v552 = *(*int64)(unsafe.Add(mBase, _consts[561]))
	*(*int64)(unsafe.Add(mBase, uint32(v550))) = v552
	*(*int64)(unsafe.Add(mBase, uint32(v19+int32(8)))) = v542
	*(*int32)(unsafe.Add(mBase, uint32(v545))) = v535
	*(*int32)(unsafe.Add(mBase, uint32(v550))) = v543
	v560 = *(*int64)(unsafe.Add(mBase, _consts[562]))
	*(*int64)(unsafe.Add(mBase, uint32(v19))) = v560
	*(*int32)(unsafe.Add(mBase, uint32(v19)+20)) = v541
	F_moduleFireServerEvent(m, int64(18), v546, v19)
	mBase = m.M
	v566 = m.ExcPending
	if v566 != 0 {
		goto L20
	} else {
		goto L83
	}
L82:
	;
	v539 = *(*int32)(unsafe.Add(mBase, uint32(v513)))
	v540 = *(*int32)(unsafe.Add(mBase, uint32(v539)+4))
	v541 = v540
	goto L81
L83:
	;
	v572 = v535
	goto L1
}
func F_freeModuleObject(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
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
	var v46 int32
	_ = v46
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v5&int32(4) == int32(0) {
		v64 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		v65 = v64
	} else {
		if v5&int32(1) != 0 {
			v14 = int32(16)
		} else {
			v14 = int32(8)
		}
		v15 = l0 + v14
		if v5&int32(2) == int32(0) {
			v46 = v15
		} else {
			v20 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15))))
			v21 = v15 + v20
			v25 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21))))
			switch v25 & int32(7) {
			case 0:
				v42 = int32(base.Ui32(v25) >> (uint(int32(3)) % 32))
			case 1:
				v32 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21+int32(-2)))))
				v42 = v32
			case 2:
				v35 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v21+int32(-4)))))
				v42 = v35
			case 3:
				v38 = *(*int32)(unsafe.Add(mBase, uint32(v21+int32(-8))))
				v42 = v38
			case 4:
				v41 = *(*int32)(unsafe.Add(mBase, uint32(v21+int32(-16))))
				v42 = v41
			default:
				v42 = int32(0)
			}
			v46 = v21 + int32(1) + v42 + int32(1)
		}
		v61 = *(*int32)(unsafe.Add(mBase, _consts[601]))
		v65 = v46 + v61
	}
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v65)+4))
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v65)))
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v70)+32))
	m.T0[v71].(func(*base.Module, int32))(m, v69)
	mBase = m.M
	v73 = m.ExcPending
	if v73 != 0 {
		return
	} else {
		F_valkey_free(m, v65)
		mBase = m.M
		v75 = m.ExcPending
		if v75 != 0 {
			return
		} else {
			return
		}
	}
}
func F_getModuleBoolConfig(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v5 = m.T0[v4].(func(*base.Module, int32, int32) int32)(m, v2, v3)
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		return int32(0)
	} else {
		return v5
	}
}
func F_getModuleUnsignedNumericConfig(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v5 int64
	_ = v5
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v5 = m.T0[v4].(func(*base.Module, int32, int32) int64)(m, v2, v3)
	mBase = m.M
	return v5
}
func F_moduleAcquireGIL(m *base.Module) {
	return
}
func F_moduleAllModulesHandleReplAsyncLoad(m *base.Module) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v30 int32
	_ = v30
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v43 int64
	_ = v43
	var v44 int64
	_ = v44
	var v45 int64
	_ = v45
	var v46 int64
	_ = v46
	var v47 int64
	_ = v47
	var v48 int64
	_ = v48
	var v49 int64
	_ = v49
	var v51 int64
	_ = v51
	var v53 int64
	_ = v53
	var v55 int64
	_ = v55
	var v57 int64
	_ = v57
	var v59 int64
	_ = v59
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v81 int32
	_ = v81
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v101 int32
	_ = v101
	var v106 int32
	_ = v106
	var v110 int32
	_ = v110
	var v114 int32
	_ = v114
	var v119 int32
	_ = v119
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v127 int32
	_ = v127
	v4 = *(*int32)(unsafe.Add(mBase, _consts[541]))
	v5 = F_dictGetIterator(m, v4)
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
	goto L3
L3:
	;
	v18 = v5 + int32(20)
	v19 = *(*int32)(unsafe.Add(mBase, uint32(v5)+16))
	if v19 != 0 {
		goto L8
	} else {
		goto L9
	}
L4:
	;
	F_dictReleaseIterator(m, v5)
	mBase = m.M
	v127 = m.ExcPending
	if v127 != 0 {
		goto L1
	} else {
		goto L36
	}
L5:
	;
	v122 = *(*int32)(unsafe.Add(mBase, uint32(v114)+8))
	goto L34
L6:
	;
	if v114 != 0 {
		goto L5
	} else {
		goto L32
	}
L7:
	;
	v25 = v18
	v26 = v22
	goto L10
L8:
	;
	v22 = int32(1)
	goto L7
L9:
	;
	v22 = int32(0)
	goto L7
L10:
	;
	switch v26 {
	case 0:
		goto L15
	default:
		goto L14
	}
L12:
	;
	v26 = int32(0)
	goto L10
L13:
	;
	goto L6
L14:
	;
	v106 = *(*int32)(unsafe.Add(mBase, uint32(v25)))
	*(*int32)(unsafe.Add(mBase, uint32(v5)+16)) = v106
	if v106 == int32(0) {
		goto L12
	} else {
		goto L31
	}
L15:
	;
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v5)+4))
	if v30 != int32(-1) {
		v69 = v30
		goto L16
	} else {
		goto L17
	}
L16:
	;
	v70 = int32(1)
	v71 = v69 + v70
	*(*int32)(unsafe.Add(mBase, uint32(v5)+4)) = v71
	v73 = int32(0)
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v5)))
	v77 = *(*int32)(unsafe.Add(mBase, uint32(v5)+8))
	v81 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v76+v77+int32(26)))))
	if v81 == int32(255) {
		goto L25
	} else {
		goto L26
	}
L17:
	;
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v5)+8))
	if v34 != 0 {
		v69 = int32(-1)
		goto L16
	} else {
		goto L18
	}
L18:
	;
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v5)))
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v5)+12))
	if v36 == int32(0) {
		goto L20
	} else {
		goto L21
	}
L19:
	;
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v62)+20))
	if v63 != int32(-1) {
		goto L22
	} else {
		goto L23
	}
L20:
	;
	v43 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v35)+16)))
	v44 = int64(*(*int8)(unsafe.Add(mBase, uint32(v35)+27)))
	v45 = int64(*(*int32)(unsafe.Add(mBase, uint32(v35)+8)))
	v46 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v35)+12)))
	v47 = int64(*(*int8)(unsafe.Add(mBase, uint32(v35)+26)))
	v48 = int64(*(*int32)(unsafe.Add(mBase, uint32(v35)+4)))
	v49 = F_wangHash64(m, v48)
	mBase = m.M
	v51 = F_wangHash64(m, v47+v49)
	mBase = m.M
	v53 = F_wangHash64(m, v46+v51)
	mBase = m.M
	v55 = F_wangHash64(m, v45+v53)
	mBase = m.M
	v57 = F_wangHash64(m, v44+v55)
	mBase = m.M
	v59 = F_wangHash64(m, v43+v57)
	mBase = m.M
	*(*int64)(unsafe.Add(mBase, uint32(v5)+24)) = v59
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v5)))
	v62 = v61
	goto L19
L21:
	;
	v39 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v35)+24)))
	v41 = v39 + int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v35)+24)) = uint16(v41)
	v62 = v35
	goto L19
L22:
	;
	v69 = v63 + int32(-1)
	goto L16
L23:
	;
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v5)+4))
	v69 = v66
	goto L16
L24:
	;
	v96 = int32(2)
	v101 = *(*int32)(unsafe.Add(mBase, uint32(v76+v94<<(uint(v96)%32)+int32(4))))
	v25 = v101 + v95<<(uint(v96)%32)
	v26 = int32(1)
	goto L10
L25:
	;
	v85 = v73
	goto L27
L26:
	;
	v85 = v70 << (uint(v81) % 32)
	goto L27
L27:
	;
	if v71 < v85 {
		v94 = v77
		v95 = v71
		goto L24
	} else {
		goto L28
	}
L28:
	;
	if v77 != 0 {
		v114 = v73
		goto L13
	} else {
		goto L29
	}
L29:
	;
	v87 = *(*int32)(unsafe.Add(mBase, uint32(v76)+20))
	if v87 == int32(-1) {
		v114 = v73
		goto L13
	} else {
		goto L30
	}
L30:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v5)+4)) = int64(4294967296)
	v94 = int32(1)
	v95 = int32(0)
	goto L24
L31:
	;
	v110 = *(*int32)(unsafe.Add(mBase, uint32(v106)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v18))) = v110
	v114 = v106
	goto L13
L32:
	;
	F_dictReleaseIterator(m, v5)
	mBase = m.M
	v119 = m.ExcPending
	if v119 != 0 {
		goto L1
	} else {
		goto L33
	}
L33:
	;
	return int32(1)
L34:
	;
	v123 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v122)+48)))
	if v123&int32(4) != 0 {
		goto L3
	} else {
		goto L35
	}
L35:
	;
	goto L4
L36:
	;
	return int32(0)
}
func F_moduleBlockedClientTimedOut(m *base.Module, l0 int32, l1 int32) {
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
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v28 int64
	_ = v28
	var v29 int32
	_ = v29
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v45 int64
	_ = v45
	var v50 int32
	_ = v50
	v7 = m.G0
	v9 = v7 - int32(80)
	m.G0 = v9
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+116))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(v11)+48))
	v13 = *(*int32)(unsafe.Add(mBase, uint32(v12)+48))
	if v13 != 0 {
		m.G0 = v9 + int32(80)
		return
	} else {
		v14 = int32(8)
		v16 = *(*int32)(unsafe.Add(mBase, uint32(v12)+4))
		if l1 != 0 {
			v19 = int32(24)
		} else {
			v19 = v14
		}
		F_moduleCreateContext(m, v9+v14, v16, v19)
		mBase = m.M
		v21 = m.ExcPending
		if v21 != 0 {
			return
		} else {
			v22 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
			*(*int32)(unsafe.Add(mBase, uint32(v9)+20)) = v12
			*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = v22
			v25 = *(*int32)(unsafe.Add(mBase, uint32(v12)+28))
			*(*int32)(unsafe.Add(mBase, uint32(v9)+48)) = v25
			v28 = *(*int64)(unsafe.Add(mBase, _consts[94]))
			v29 = *(*int32)(unsafe.Add(mBase, uint32(v12)+16))
			if v29 == int32(0) {
				F_moduleFreeContext(m, v9+int32(8))
				mBase = m.M
				v41 = m.ExcPending
				if v41 != 0 {
					return
				} else {
					if l1 != 0 {
						*(*int32)(unsafe.Add(mBase, uint32(v12)+20)) = int32(0)
						m.G0 = v9 + int32(80)
						return
					} else {
						v42 = *(*int32)(unsafe.Add(mBase, uint32(v12)+64))
						v45 = *(*int64)(unsafe.Add(mBase, _consts[94]))
						F_updateStatsOnUnblock(m, l0, v42, int32(0), base.B2i32(v45 != v28)<<(uint(int32(1))%32))
						mBase = m.M
						v50 = m.ExcPending
						if v50 != 0 {
							return
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v12)+20)) = int32(0)
							m.G0 = v9 + int32(80)
							return
						}
					}
				}
			} else {
				v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
				v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
				v36 = m.T0[v29].(func(*base.Module, int32, int32, int32) int32)(m, v9+int32(8), v34, v35)
				mBase = m.M
				v37 = m.ExcPending
				if v37 != 0 {
					return
				} else {
					F_moduleFreeContext(m, v9+int32(8))
					mBase = m.M
					v41 = m.ExcPending
					if v41 != 0 {
						return
					} else {
						if l1 != 0 {
							*(*int32)(unsafe.Add(mBase, uint32(v12)+20)) = int32(0)
							m.G0 = v9 + int32(80)
							return
						} else {
							v42 = *(*int32)(unsafe.Add(mBase, uint32(v12)+64))
							v45 = *(*int64)(unsafe.Add(mBase, _consts[94]))
							F_updateStatsOnUnblock(m, l0, v42, int32(0), base.B2i32(v45 != v28)<<(uint(int32(1))%32))
							mBase = m.M
							v50 = m.ExcPending
							if v50 != 0 {
								return
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v12)+20)) = int32(0)
								m.G0 = v9 + int32(80)
								return
							}
						}
					}
				}
			}
		}
	}
}
func F_moduleCallClusterReceivers(m *base.Module, l0 int32, l1 int64, l2 int32, l3 int32, l4 int32) {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v18 int32
	_ = v18
	var v27 int32
	_ = v27
	var v30 int64
	_ = v30
	var v32 int32
	_ = v32
	var v35 int64
	_ = v35
	var v37 int32
	_ = v37
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v94 int32
	_ = v94
	var v98 int32
	_ = v98
	var v102 int32
	_ = v102
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v109 int64
	_ = v109
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v119 int64
	_ = v119
	var v123 int64
	_ = v123
	var v127 int32
	_ = v127
	var v131 int32
	_ = v131
	var v139 int64
	_ = v139
	var v141 int32
	_ = v141
	var v145 int64
	_ = v145
	var v149 int64
	_ = v149
	var v152 int32
	_ = v152
	var v160 int64
	_ = v160
	var v166 int32
	_ = v166
	var v168 int32
	_ = v168
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	v10 = m.G0
	v12 = v10 - int32(80)
	m.G0 = v12
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l2<<(uint(int32(2))%32))+uint32(_consts[565])))
	if v18 == int32(0) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	m.G0 = v12 + int32(80)
	return
L2:
	;
	v27 = v18
	goto L3
L3:
	;
	v30 = *(*int64)(unsafe.Add(mBase, uint32(v27)))
	if v30 != l1 {
		goto L5
	} else {
		goto L6
	}
L4:
	;
	goto L1
L5:
	;
	v173 = *(*int32)(unsafe.Add(mBase, uint32(v27)+16))
	if v173 != 0 {
		v27 = v173
		goto L3
	} else {
		goto L24
	}
L6:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v27)+12))
	v35 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v12+int32(32)))) = v35
	v37 = int32(64)
	*(*int64)(unsafe.Add(mBase, uint32(v12+v37))) = v35
	*(*int64)(unsafe.Add(mBase, uint32(v12+int32(56)))) = v35
	*(*int64)(unsafe.Add(mBase, uint32(v12+int32(48)))) = v35
	*(*int64)(unsafe.Add(mBase, uint32(v12+int32(40)))) = v35
	*(*int64)(unsafe.Add(mBase, uint32(v12+int32(24)))) = v35
	*(*int32)(unsafe.Add(mBase, uint32(v12)+36)) = v37
	*(*int64)(unsafe.Add(mBase, uint32(v12+int32(72)))) = v35
	*(*int64)(unsafe.Add(mBase, uint32(v12)+16)) = v35
	*(*int32)(unsafe.Add(mBase, uint32(v12)+12)) = v32
	*(*int32)(unsafe.Add(mBase, uint32(v12)+8)) = int32(561)
	v70 = int32(0)
	v71 = *(*int32)(unsafe.Add(mBase, _consts[546]))
	if v71 == v70 {
		goto L8
	} else {
		goto L9
	}
L7:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = v102
	v106 = *(*int32)(unsafe.Add(mBase, _consts[131]))
	v107 = int32(0)
	v108 = *(*int32)(unsafe.Add(mBase, _consts[23]))
	v109 = m.T0[v108].(func(*base.Module) int64)(m)
	mBase = m.M
	if v106 == v107 {
		goto L14
	} else {
		goto L15
	}
L8:
	;
	v90 = F_createClient(m, int32(0))
	mBase = m.M
	v91 = m.ExcPending
	if v91 != 0 {
		goto L11
	} else {
		goto L12
	}
L9:
	;
	v74 = int32(0)
	v76 = v71 + int32(-1)
	*(*int32)(unsafe.Add(mBase, _consts[546])) = v76
	v79 = *(*int32)(unsafe.Add(mBase, _consts[547]))
	v83 = *(*int32)(unsafe.Add(mBase, uint32(v79+v76<<(uint(int32(2))%32))))
	v85 = *(*int32)(unsafe.Add(mBase, _consts[548]))
	if base.Ui32(v85) <= base.Ui32(v76) {
		v102 = v83
		goto L7
	} else {
		goto L10
	}
L10:
	;
	*(*int32)(unsafe.Add(mBase, _consts[548])) = v76
	v102 = v83
	goto L7
L11:
	;
	return
L12:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v90)+328)) = int32(0)
	v94 = *(*int32)(unsafe.Add(mBase, uint32(v90)+200))
	*(*int32)(unsafe.Add(mBase, uint32(v90)+200)) = v94 | int32(1073741824)
	v98 = *(*int32)(unsafe.Add(mBase, uint32(v90)+204))
	*(*int32)(unsafe.Add(mBase, uint32(v90)+204)) = v98 | int32(268435456)
	v102 = v90
	goto L7
L13:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v12)+64)) = v123
	v127 = int32(0)
	v131 = *(*int32)(unsafe.Add(mBase, _consts[95]))
	*(*int32)(unsafe.Add(mBase, _consts[95])) = v131 + int32(1)
	goto L18
L14:
	;
	v119 = *(*int64)(unsafe.Add(mBase, _consts[529]))
	v123 = v119*int64(1000) + v109
	goto L13
L15:
	;
	v114 = *(*int32)(unsafe.Add(mBase, _consts[219]))
	v115 = base.I32_div_s(int32(1000000), v114)
	v123 = v109 + base.I64_extend_i32_s(v115)
	goto L13
L16:
	;
	v166 = *(*int32)(unsafe.Add(mBase, uint32(v27)+8))
	m.T0[v166].(func(*base.Module, int32, int32, int32, int32, int32))(m, v12+int32(8), l0, l2, l3, l4)
	mBase = m.M
	v168 = m.ExcPending
	if v168 != 0 {
		goto L11
	} else {
		goto L22
	}
L17:
	;
	goto L16
L18:
	;
	if v131 != 0 {
		goto L17
	} else {
		goto L19
	}
L19:
	;
	goto L21
L20:
	;
	v141 = int32(0)
	*(*int64)(unsafe.Add(mBase, _consts[96])) = v139
	v145 = base.I64_div_s(v139, int64(1000))
	*(*int64)(unsafe.Add(mBase, _consts[35])) = v145
	v149 = base.I64_div_s(v139, int64(1000000))
	*(*int64)(unsafe.Add(mBase, _consts[47])) = v149
	v152 = *(*int32)(unsafe.Add(mBase, _consts[97]))
	F_lrulfu_updateClockAndPolicy(m, v145, int32(base.Ui32(v152&int32(2))>>(uint(int32(1))%32)))
	mBase = m.M
	v160 = *(*int64)(unsafe.Add(mBase, _consts[35]))
	*(*int64)(unsafe.Add(mBase, _consts[98])) = v160
	goto L17
L21:
	;
	v139 = F_ustime(m)
	mBase = m.M
	goto L20
L22:
	;
	F_moduleFreeContext(m, v12+int32(8))
	mBase = m.M
	v172 = m.ExcPending
	if v172 != 0 {
		goto L11
	} else {
		goto L23
	}
L23:
	;
	goto L1
L24:
	;
	goto L4
}
func F_moduleConvertKeySpecsFlags(m *base.Module, l0 int64, l1 int32) int64 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v12 int64
	_ = v12
	var v23 int64
	_ = v23
	var v24 int64
	_ = v24
	var v27 int64
	_ = v27
	var v35 int64
	_ = v35
	var v37 int64
	_ = v37
	var v39 int32
	_ = v39
	var v42 int64
	_ = v42
	var v50 int64
	_ = v50
	var v52 int64
	_ = v52
	var v55 int64
	_ = v55
	var v63 int64
	_ = v63
	var v65 int64
	_ = v65
	var v67 int32
	_ = v67
	var v70 int64
	_ = v70
	var v78 int64
	_ = v78
	var v80 int64
	_ = v80
	var v83 int64
	_ = v83
	var v91 int64
	_ = v91
	var v93 int64
	_ = v93
	var v95 int32
	_ = v95
	var v98 int64
	_ = v98
	var v106 int64
	_ = v106
	var v108 int64
	_ = v108
	var v111 int64
	_ = v111
	var v119 int64
	_ = v119
	var v121 int64
	_ = v121
	var v123 int32
	_ = v123
	var v126 int64
	_ = v126
	var v134 int64
	_ = v134
	var v136 int64
	_ = v136
	var v139 int64
	_ = v139
	var v147 int64
	_ = v147
	var v149 int64
	_ = v149
	var v154 int64
	_ = v154
	var v162 int64
	_ = v162
	var v164 int64
	_ = v164
	v7 = l1 ^ int32(1)
	v9 = v7 << (uint(int32(3)) % 32)
	v12 = *(*int64)(unsafe.Add(mBase, uint32(v9)+uint32(_consts[530])))
	if base.B2i32(v12&l0 == int64(0)) == int32(0) {
		v23 = *(*int64)(unsafe.Add(mBase, uint32(l1<<(uint(int32(3))%32))+uint32(_consts[530])))
		v24 = v23
	} else {
		v24 = int64(0)
	}
	v27 = *(*int64)(unsafe.Add(mBase, uint32(v9)+uint32(_consts[531])))
	if v27&l0 == int64(0) {
		v37 = v24
	} else {
		v35 = *(*int64)(unsafe.Add(mBase, uint32(l1<<(uint(int32(3))%32))+uint32(_consts[531])))
		v37 = v35 | v24
	}
	v39 = v7 << (uint(int32(3)) % 32)
	v42 = *(*int64)(unsafe.Add(mBase, uint32(v39)+uint32(_consts[532])))
	if v42&l0 == int64(0) {
		v52 = v37
	} else {
		v50 = *(*int64)(unsafe.Add(mBase, uint32(l1<<(uint(int32(3))%32))+uint32(_consts[532])))
		v52 = v50 | v37
	}
	v55 = *(*int64)(unsafe.Add(mBase, uint32(v39)+uint32(_consts[533])))
	if v55&l0 == int64(0) {
		v65 = v52
	} else {
		v63 = *(*int64)(unsafe.Add(mBase, uint32(l1<<(uint(int32(3))%32))+uint32(_consts[533])))
		v65 = v63 | v52
	}
	v67 = v7 << (uint(int32(3)) % 32)
	v70 = *(*int64)(unsafe.Add(mBase, uint32(v67)+uint32(_consts[534])))
	if v70&l0 == int64(0) {
		v80 = v65
	} else {
		v78 = *(*int64)(unsafe.Add(mBase, uint32(l1<<(uint(int32(3))%32))+uint32(_consts[534])))
		v80 = v78 | v65
	}
	v83 = *(*int64)(unsafe.Add(mBase, uint32(v67)+uint32(_consts[535])))
	if v83&l0 == int64(0) {
		v93 = v80
	} else {
		v91 = *(*int64)(unsafe.Add(mBase, uint32(l1<<(uint(int32(3))%32))+uint32(_consts[535])))
		v93 = v91 | v80
	}
	v95 = v7 << (uint(int32(3)) % 32)
	v98 = *(*int64)(unsafe.Add(mBase, uint32(v95)+uint32(_consts[536])))
	if v98&l0 == int64(0) {
		v108 = v93
	} else {
		v106 = *(*int64)(unsafe.Add(mBase, uint32(l1<<(uint(int32(3))%32))+uint32(_consts[536])))
		v108 = v106 | v93
	}
	v111 = *(*int64)(unsafe.Add(mBase, uint32(v95)+uint32(_consts[537])))
	if v111&l0 == int64(0) {
		v121 = v108
	} else {
		v119 = *(*int64)(unsafe.Add(mBase, uint32(l1<<(uint(int32(3))%32))+uint32(_consts[537])))
		v121 = v119 | v108
	}
	v123 = v7 << (uint(int32(3)) % 32)
	v126 = *(*int64)(unsafe.Add(mBase, uint32(v123)+uint32(_consts[538])))
	if v126&l0 == int64(0) {
		v136 = v121
	} else {
		v134 = *(*int64)(unsafe.Add(mBase, uint32(l1<<(uint(int32(3))%32))+uint32(_consts[538])))
		v136 = v134 | v121
	}
	v139 = *(*int64)(unsafe.Add(mBase, uint32(v123)+uint32(_consts[539])))
	if v139&l0 == int64(0) {
		v149 = v136
	} else {
		v147 = *(*int64)(unsafe.Add(mBase, uint32(l1<<(uint(int32(3))%32))+uint32(_consts[539])))
		v149 = v147 | v136
	}
	v154 = *(*int64)(unsafe.Add(mBase, uint32(v7<<(uint(int32(3))%32))+uint32(_consts[540])))
	if v154&l0 == int64(0) {
		v164 = v149
	} else {
		v162 = *(*int64)(unsafe.Add(mBase, uint32(l1<<(uint(int32(3))%32))+uint32(_consts[540])))
		v164 = v162 | v149
	}
	return v164
}
func F_moduleCopyCommandArgs(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v14 int32
	_ = v14
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v103 int32
	_ = v103
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v110 int32
	_ = v110
	var v125 int32
	_ = v125
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v14 = int32(0)
	goto L1
L1:
	;
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l0+v9*v14)))
	if v23 != 0 {
		v14 = v14 + int32(1)
		goto L1
	} else {
		goto L3
	}
L2:
	;
	if base.Ui32(int32(97612893)) <= base.Ui32(v14) {
		goto L4
	} else {
		goto L5
	}
L3:
	;
	goto L2
L4:
	;
	F__serverAssert(m, int32(_a918), int32(_a917), int32(2320))
	mBase = m.M
	v125 = m.ExcPending
	if v125 != 0 {
		goto L6
	} else {
		goto L37
	}
L5:
	;
	v26 = int32(44)
	v30 = F_valkey_calloc(m, v14*v26+v26)
	mBase = m.M
	v33 = m.ExcPending
	if v33 != 0 {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	return int32(0)
L7:
	;
	if v14 == int32(0) {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	return v30
L9:
	;
	v43 = int32(0)
	goto L10
L10:
	;
	v47 = v30 + v43*int32(44)
	v48 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v50 = l0 + v48*v43
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v50)))
	v52 = F_zstrdup(m, v51)
	mBase = m.M
	v53 = m.ExcPending
	if v53 != 0 {
		goto L6
	} else {
		goto L12
	}
L11:
	;
	goto L8
L12:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v47))) = v52
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v50)+4))
	if base.Ui32(int32(8)) < base.Ui32(v55) {
		goto L15
	} else {
		goto L16
	}
L13:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v47)+8)) = v67
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v50)+12))
	if v69 == int32(0) {
		goto L18
	} else {
		goto L19
	}
L14:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v47)+4)) = v64
	v67 = int32(-1)
	goto L13
L15:
	;
	v64 = int32(-1)
	goto L14
L16:
	;
	if v55 != int32(3) {
		v64 = v55
		goto L14
	} else {
		goto L17
	}
L17:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v47)+4)) = int32(3)
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v50)+8))
	v67 = v62
	goto L13
L18:
	;
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v50)+16))
	if v75 == int32(0) {
		goto L21
	} else {
		goto L22
	}
L19:
	;
	v72 = F_zstrdup(m, v69)
	mBase = m.M
	v73 = m.ExcPending
	if v73 != 0 {
		goto L6
	} else {
		goto L20
	}
L20:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v47)+12)) = v72
	goto L18
L21:
	;
	v81 = *(*int32)(unsafe.Add(mBase, uint32(v50)+20))
	if v81 == int32(0) {
		goto L24
	} else {
		goto L25
	}
L22:
	;
	v78 = F_zstrdup(m, v75)
	mBase = m.M
	v79 = m.ExcPending
	if v79 != 0 {
		goto L6
	} else {
		goto L23
	}
L23:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v47)+16)) = v78
	goto L21
L24:
	;
	v87 = *(*int32)(unsafe.Add(mBase, uint32(v50)+28))
	if v87 == int32(0) {
		goto L27
	} else {
		goto L28
	}
L25:
	;
	v84 = F_zstrdup(m, v81)
	mBase = m.M
	v85 = m.ExcPending
	if v85 != 0 {
		goto L6
	} else {
		goto L26
	}
L26:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v47)+20)) = v84
	goto L24
L27:
	;
	v93 = *(*int32)(unsafe.Add(mBase, uint32(v50)+36))
	if v93 == int32(0) {
		goto L30
	} else {
		goto L31
	}
L28:
	;
	v90 = F_zstrdup(m, v87)
	mBase = m.M
	v91 = m.ExcPending
	if v91 != 0 {
		goto L6
	} else {
		goto L29
	}
L29:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v47)+28)) = v90
	goto L27
L30:
	;
	v99 = *(*int32)(unsafe.Add(mBase, uint32(v50)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v47)+24)) = v99 & int32(7)
	v103 = *(*int32)(unsafe.Add(mBase, uint32(v50)+32))
	if v103 == int32(0) {
		goto L33
	} else {
		goto L34
	}
L31:
	;
	v96 = F_zstrdup(m, v93)
	mBase = m.M
	v97 = m.ExcPending
	if v97 != 0 {
		goto L6
	} else {
		goto L32
	}
L32:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v47)+40)) = v96
	goto L30
L33:
	;
	v110 = v43 + int32(1)
	if v110 != v14 {
		v43 = v110
		goto L10
	} else {
		goto L36
	}
L34:
	;
	v106 = F_moduleCopyCommandArgs(m, v103, l1)
	mBase = m.M
	v107 = m.ExcPending
	if v107 != 0 {
		goto L6
	} else {
		goto L35
	}
L35:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v47)+36)) = v106
	goto L33
L36:
	;
	goto L11
L37:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_moduleCount(m *base.Module) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	v3 = *(*int32)(unsafe.Add(mBase, _consts[541]))
	v4 = *(*int32)(unsafe.Add(mBase, uint32(v3)+16))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(v3)+12))
	return v4 + v5
}
func F_moduleCreateArgvFromUserFormat(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v38 int32
	_ = v38
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v53 int32
	_ = v53
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v83 int32
	_ = v83
	var v86 int32
	_ = v86
	var v90 int32
	_ = v90
	var v95 int32
	_ = v95
	var v97 int32
	_ = v97
	var v101 int32
	_ = v101
	var v107 int32
	_ = v107
	var v110 int32
	_ = v110
	var v116 int32
	_ = v116
	var v120 int32
	_ = v120
	var v122 int32
	_ = v122
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
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
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v157 int32
	_ = v157
	var v164 int32
	_ = v164
	var v166 int32
	_ = v166
	var v168 int32
	_ = v168
	var v182 int32
	_ = v182
	var v185 int32
	_ = v185
	var v196 int32
	_ = v196
	var v199 int32
	_ = v199
	var v203 int32
	_ = v203
	var v208 int32
	_ = v208
	var v210 int32
	_ = v210
	var v214 int32
	_ = v214
	var v220 int32
	_ = v220
	var v223 int32
	_ = v223
	var v229 int32
	_ = v229
	var v233 int32
	_ = v233
	var v235 int32
	_ = v235
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v247 int32
	_ = v247
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v259 int32
	_ = v259
	var v261 int32
	_ = v261
	var v264 int32
	_ = v264
	var v269 int32
	_ = v269
	var v270 int32
	_ = v270
	var v273 int32
	_ = v273
	var v274 int32
	_ = v274
	var v275 int32
	_ = v275
	var v278 int32
	_ = v278
	var v279 int32
	_ = v279
	var v280 int32
	_ = v280
	var v283 int32
	_ = v283
	var v284 int32
	_ = v284
	var v285 int32
	_ = v285
	var v288 int32
	_ = v288
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
	var v304 int32
	_ = v304
	var v311 int32
	_ = v311
	var v312 int32
	_ = v312
	var v313 int32
	_ = v313
	var v314 int32
	_ = v314
	var v316 int32
	_ = v316
	var v328 int32
	_ = v328
	var v329 int64
	_ = v329
	var v330 int32
	_ = v330
	var v331 int32
	_ = v331
	var v333 int32
	_ = v333
	var v340 int32
	_ = v340
	var v341 int32
	_ = v341
	var v343 int32
	_ = v343
	var v346 int32
	_ = v346
	var v349 int32
	_ = v349
	var v350 int32
	_ = v350
	var v353 int32
	_ = v353
	var v360 int32
	_ = v360
	var v366 int32
	_ = v366
	var v367 int32
	_ = v367
	var v369 int32
	_ = v369
	var v373 int32
	_ = v373
	var v375 int32
	_ = v375
	var v376 int32
	_ = v376
	var v378 int32
	_ = v378
	var v384 int32
	_ = v384
	var v385 int32
	_ = v385
	var v392 int32
	_ = v392
	var v400 int32
	_ = v400
	var v408 int32
	_ = v408
	var v416 int32
	_ = v416
	var v424 int32
	_ = v424
	var v432 int32
	_ = v432
	var v440 int32
	_ = v440
	var v448 int32
	_ = v448
	var v456 int32
	_ = v456
	var v464 int32
	_ = v464
	var v472 int32
	_ = v472
	var v480 int32
	_ = v480
	v6 = int32(0)
	if l1&int32(3) == v6 {
		v34 = l1
		goto L3
	} else {
		goto L4
	}
L1:
	;
	v69 = v67 + int32(1)
	v72 = F_valkey_realloc(m, v6, v69<<(uint(int32(2))%32))
	mBase = m.M
	v75 = m.ExcPending
	if v75 != 0 {
		goto L17
	} else {
		goto L18
	}
L2:
	;
	v67 = v59 - l1
	goto L1
L3:
	;
	v38 = v34
	goto L11
L4:
	;
	v20 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	if v20 != 0 {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	v23 = l1
	goto L7
L6:
	;
	v67 = l1 - l1
	goto L1
L7:
	;
	v27 = v23 + int32(1)
	if v27&int32(3) == int32(0) {
		v34 = v27
		goto L3
	} else {
		goto L9
	}
L9:
	;
	v32 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v27))))
	if v32 != 0 {
		v23 = v27
		goto L7
	} else {
		goto L10
	}
L10:
	;
	v59 = v27
	goto L2
L11:
	;
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v38)))
	v47 = int32(-2139062144)
	if (int32(16843008)-v44|v44)&v47 == v47 {
		v38 = v38 + int32(4)
		goto L11
	} else {
		goto L13
	}
L12:
	;
	v53 = v38
	goto L14
L13:
	;
	goto L12
L14:
	;
	v57 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v53))))
	if v57 != 0 {
		v53 = v53 + int32(1)
		goto L14
	} else {
		goto L16
	}
L15:
	;
	v59 = v53
	goto L2
L16:
	;
	goto L15
L17:
	;
	return int32(0)
L18:
	;
	if l0&int32(3) == int32(0) {
		v97 = l0
		goto L21
	} else {
		goto L22
	}
L19:
	;
	v131 = F_createStringObject_1(m, l0, v130)
	mBase = m.M
	v132 = m.ExcPending
	if v132 != 0 {
		goto L17
	} else {
		goto L35
	}
L20:
	;
	v130 = v122 - l0
	goto L19
L21:
	;
	v101 = v97
	goto L29
L22:
	;
	v83 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v83 != 0 {
		goto L23
	} else {
		goto L24
	}
L23:
	;
	v86 = l0
	goto L25
L24:
	;
	v130 = l0 - l0
	goto L19
L25:
	;
	v90 = v86 + int32(1)
	if v90&int32(3) == int32(0) {
		v97 = v90
		goto L21
	} else {
		goto L27
	}
L27:
	;
	v95 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v90))))
	if v95 != 0 {
		v86 = v90
		goto L25
	} else {
		goto L28
	}
L28:
	;
	v122 = v90
	goto L20
L29:
	;
	v107 = *(*int32)(unsafe.Add(mBase, uint32(v101)))
	v110 = int32(-2139062144)
	if (int32(16843008)-v107|v107)&v110 == v110 {
		v101 = v101 + int32(4)
		goto L29
	} else {
		goto L31
	}
L30:
	;
	v116 = v101
	goto L32
L31:
	;
	goto L30
L32:
	;
	v120 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v116))))
	if v120 != 0 {
		v116 = v116 + int32(1)
		goto L32
	} else {
		goto L34
	}
L33:
	;
	v122 = v116
	goto L20
L34:
	;
	goto L33
L35:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v72))) = v131
	v135 = int32(1)
	v136 = l1
	v139 = l4
	v140 = v69
	v141 = v72
	goto L37
L36:
	;
	if l2 == int32(0) {
		goto L120
	} else {
		goto L121
	}
L37:
	;
	v146 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v136))))
	switch v146 {
	case 0:
		goto L36
	default:
		goto L58
	case 33:
		goto L52
	case 48:
		goto L48
	case 51:
		goto L49
	case 65:
		goto L51
	case 67:
		goto L47
	case 68:
		goto L42
	case 69:
		goto L43
	case 75:
		goto L41
	case 77:
		goto L44
	case 82:
		goto L50
	case 83:
		goto L46
	case 87:
		goto L45
	case 88:
		goto L40
	case 98:
		goto L55
	case 99:
		goto L57
	case 108:
		goto L54
	case 115:
		goto L56
	case 118:
		goto L53
	}
L39:
	;
	v136 = v136 + int32(1)
	goto L37
L40:
	;
	if l3 == int32(0) {
		goto L39
	} else {
		goto L119
	}
L41:
	;
	if l3 == int32(0) {
		goto L39
	} else {
		goto L118
	}
L42:
	;
	if l3 == int32(0) {
		goto L39
	} else {
		goto L117
	}
L43:
	;
	if l3 == int32(0) {
		goto L39
	} else {
		goto L116
	}
L44:
	;
	if l3 == int32(0) {
		goto L39
	} else {
		goto L115
	}
L45:
	;
	if l3 == int32(0) {
		goto L39
	} else {
		goto L114
	}
L46:
	;
	if l3 == int32(0) {
		goto L39
	} else {
		goto L113
	}
L47:
	;
	if l3 == int32(0) {
		goto L39
	} else {
		goto L112
	}
L48:
	;
	if l3 == int32(0) {
		goto L39
	} else {
		goto L111
	}
L49:
	;
	if l3 == int32(0) {
		goto L39
	} else {
		goto L110
	}
L50:
	;
	if l3 == int32(0) {
		goto L39
	} else {
		goto L109
	}
L51:
	;
	if l3 == int32(0) {
		goto L39
	} else {
		goto L108
	}
L52:
	;
	if l3 == int32(0) {
		goto L39
	} else {
		goto L107
	}
L53:
	;
	v340 = v139 + int32(8)
	v341 = *(*int32)(unsafe.Add(mBase, uint32(v139)))
	v343 = *(*int32)(unsafe.Add(mBase, uint32(v139)+4))
	v346 = v140 + v343 + int32(-1)
	v349 = F_valkey_realloc(m, v141, v346<<(uint(int32(2))%32))
	mBase = m.M
	v350 = m.ExcPending
	if v350 != 0 {
		goto L17
	} else {
		goto L100
	}
L54:
	;
	v328 = (v139 + int32(7)) & int32(-8)
	v329 = *(*int64)(unsafe.Add(mBase, uint32(v328)))
	v330 = F_createStringObjectFromLongLongWithSds(m, v329)
	mBase = m.M
	v331 = m.ExcPending
	if v331 != 0 {
		goto L17
	} else {
		goto L99
	}
L55:
	;
	v311 = *(*int32)(unsafe.Add(mBase, uint32(v139)))
	v312 = *(*int32)(unsafe.Add(mBase, uint32(v139)+4))
	v313 = F_createStringObject_1(m, v311, v312)
	mBase = m.M
	v314 = m.ExcPending
	if v314 != 0 {
		goto L17
	} else {
		goto L98
	}
L56:
	;
	v253 = *(*int32)(unsafe.Add(mBase, uint32(v139)))
	v254 = *(*int32)(unsafe.Add(mBase, uint32(v253)+4))
	if v254&int32(-8) != int32(-16) {
		goto L84
	} else {
		goto L85
	}
L57:
	;
	v185 = *(*int32)(unsafe.Add(mBase, uint32(v139)))
	if v185&int32(3) == int32(0) {
		v210 = v185
		goto L68
	} else {
		goto L69
	}
L58:
	;
	v147 = int32(0)
	if v135 <= v147 {
		goto L59
	} else {
		goto L60
	}
L59:
	;
	F_valkey_free(m, v141)
	mBase = m.M
	v182 = m.ExcPending
	if v182 != 0 {
		goto L17
	} else {
		goto L65
	}
L60:
	;
	v157 = v147
	goto L61
L61:
	;
	v164 = *(*int32)(unsafe.Add(mBase, uint32(v141+v157<<(uint(int32(2))%32))))
	F_decrRefCount(m, v164)
	mBase = m.M
	v166 = m.ExcPending
	if v166 != 0 {
		goto L17
	} else {
		goto L63
	}
L62:
	;
	goto L59
L63:
	;
	v168 = v157 + int32(1)
	if v168 != v135 {
		v157 = v168
		goto L61
	} else {
		goto L64
	}
L64:
	;
	goto L62
L65:
	;
	return int32(0)
L66:
	;
	v244 = F_createStringObject_1(m, v185, v243)
	mBase = m.M
	v245 = m.ExcPending
	if v245 != 0 {
		goto L17
	} else {
		goto L82
	}
L67:
	;
	v243 = v235 - v185
	goto L66
L68:
	;
	v214 = v210
	goto L76
L69:
	;
	v196 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v185))))
	if v196 != 0 {
		goto L70
	} else {
		goto L71
	}
L70:
	;
	v199 = v185
	goto L72
L71:
	;
	v243 = v185 - v185
	goto L66
L72:
	;
	v203 = v199 + int32(1)
	if v203&int32(3) == int32(0) {
		v210 = v203
		goto L68
	} else {
		goto L74
	}
L74:
	;
	v208 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v203))))
	if v208 != 0 {
		v199 = v203
		goto L72
	} else {
		goto L75
	}
L75:
	;
	v235 = v203
	goto L67
L76:
	;
	v220 = *(*int32)(unsafe.Add(mBase, uint32(v214)))
	v223 = int32(-2139062144)
	if (int32(16843008)-v220|v220)&v223 == v223 {
		v214 = v214 + int32(4)
		goto L76
	} else {
		goto L78
	}
L77:
	;
	v229 = v214
	goto L79
L78:
	;
	goto L77
L79:
	;
	v233 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v229))))
	if v233 != 0 {
		v229 = v229 + int32(1)
		goto L79
	} else {
		goto L81
	}
L80:
	;
	v235 = v229
	goto L67
L81:
	;
	goto L80
L82:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v141+v135<<(uint(int32(2))%32)))) = v244
	v247 = int32(1)
	v135 = v135 + v247
	v136 = v136 + v247
	v139 = v139 + int32(4)
	goto L37
L83:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v141+v135<<(uint(int32(2))%32)))) = v294
	v304 = int32(1)
	v135 = v135 + v304
	v136 = v136 + v304
	v139 = v139 + int32(4)
	goto L37
L84:
	;
	F_incrRefCount(m, v253)
	mBase = m.M
	v293 = m.ExcPending
	if v293 != 0 {
		goto L17
	} else {
		goto L97
	}
L85:
	;
	v259 = F_objectGetVal(m, v253)
	mBase = m.M
	v261 = F_objectGetVal(m, v253)
	mBase = m.M
	v264 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v261+int32(-1)))))
	switch v264 & int32(7) {
	case 0:
		goto L91
	case 1:
		goto L90
	case 2:
		goto L89
	case 3:
		goto L88
	case 4:
		goto L87
	default:
		v289 = int32(0)
		goto L86
	}
L86:
	;
	v290 = F_createStringObject_1(m, v259, v289)
	mBase = m.M
	v291 = m.ExcPending
	if v291 != 0 {
		goto L17
	} else {
		goto L96
	}
L87:
	;
	v288 = *(*int32)(unsafe.Add(mBase, uint32(v261+int32(-17))))
	v289 = v288
	goto L86
L88:
	;
	v283 = *(*int32)(unsafe.Add(mBase, uint32(v261+int32(-9))))
	v284 = F_createStringObject_1(m, v259, v283)
	mBase = m.M
	v285 = m.ExcPending
	if v285 != 0 {
		goto L17
	} else {
		goto L95
	}
L89:
	;
	v278 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v261+int32(-5)))))
	v279 = F_createStringObject_1(m, v259, v278)
	mBase = m.M
	v280 = m.ExcPending
	if v280 != 0 {
		goto L17
	} else {
		goto L94
	}
L90:
	;
	v273 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v261+int32(-3)))))
	v274 = F_createStringObject_1(m, v259, v273)
	mBase = m.M
	v275 = m.ExcPending
	if v275 != 0 {
		goto L17
	} else {
		goto L93
	}
L91:
	;
	v269 = F_createStringObject_1(m, v259, int32(base.Ui32(v264)>>(uint(int32(3))%32)))
	mBase = m.M
	v270 = m.ExcPending
	if v270 != 0 {
		goto L17
	} else {
		goto L92
	}
L92:
	;
	v294 = v269
	goto L83
L93:
	;
	v294 = v274
	goto L83
L94:
	;
	v294 = v279
	goto L83
L95:
	;
	v294 = v284
	goto L83
L96:
	;
	v294 = v290
	goto L83
L97:
	;
	v294 = v253
	goto L83
L98:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v141+v135<<(uint(int32(2))%32)))) = v313
	v316 = int32(1)
	v135 = v135 + v316
	v136 = v136 + v316
	v139 = v139 + int32(8)
	goto L37
L99:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v141+v135<<(uint(int32(2))%32)))) = v330
	v333 = int32(1)
	v135 = v135 + v333
	v136 = v136 + v333
	v139 = v328 + int32(8)
	goto L37
L100:
	;
	if v343 != 0 {
		goto L101
	} else {
		goto L102
	}
L101:
	;
	v353 = v135
	v360 = int32(0)
	goto L103
L102:
	;
	v136 = v136 + int32(1)
	v139 = v340
	v140 = v346
	v141 = v349
	goto L37
L103:
	;
	v366 = v341 + v360<<(uint(int32(2))%32)
	v367 = *(*int32)(unsafe.Add(mBase, uint32(v366)))
	F_incrRefCount(m, v367)
	mBase = m.M
	v369 = m.ExcPending
	if v369 != 0 {
		goto L17
	} else {
		goto L105
	}
L104:
	;
	v135 = v376
	v136 = v136 + int32(1)
	v139 = v340
	v140 = v346
	v141 = v349
	goto L37
L105:
	;
	v373 = *(*int32)(unsafe.Add(mBase, uint32(v366)))
	*(*int32)(unsafe.Add(mBase, uint32(v349+v353<<(uint(int32(2))%32)))) = v373
	v375 = int32(1)
	v376 = v353 + v375
	v378 = v360 + v375
	if v378 != v343 {
		v353 = v376
		v360 = v378
		goto L103
	} else {
		goto L106
	}
L106:
	;
	goto L104
L107:
	;
	v384 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	v385 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v384 | v385
	v136 = v136 + v385
	goto L37
L108:
	;
	v392 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v392 | int32(2)
	v136 = v136 + int32(1)
	goto L37
L109:
	;
	v400 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v400 | int32(4)
	v136 = v136 + int32(1)
	goto L37
L110:
	;
	v408 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v408 | int32(8)
	v136 = v136 + int32(1)
	goto L37
L111:
	;
	v416 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v416 | int32(16)
	v136 = v136 + int32(1)
	goto L37
L112:
	;
	v424 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v424 | int32(32)
	v136 = v136 + int32(1)
	goto L37
L113:
	;
	v432 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v432 | int32(64)
	v136 = v136 + int32(1)
	goto L37
L114:
	;
	v440 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v440 | int32(128)
	v136 = v136 + int32(1)
	goto L37
L115:
	;
	v448 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v448 | int32(512)
	v136 = v136 + int32(1)
	goto L37
L116:
	;
	v456 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v456 | int32(256)
	v136 = v136 + int32(1)
	goto L37
L117:
	;
	v464 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v464 | int32(1280)
	v136 = v136 + int32(1)
	goto L37
L118:
	;
	v472 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v472 | int32(2048)
	v136 = v136 + int32(1)
	goto L37
L119:
	;
	v480 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v480 | int32(4096)
	goto L39
L120:
	;
	return v141
L121:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v135
	goto L120
}
func F_moduleCreateCommandProxy(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int64, l5 int32, l6 int32, l7 int32) int32 {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v61 int64
	_ = v61
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v70 int64
	_ = v70
	v11 = F_valkey_calloc(m, int32(12))
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		return int32(0)
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(v11)+4)) = l3
		*(*int32)(unsafe.Add(mBase, uint32(v11))) = l0
		v18 = F_valkey_calloc(m, int32(224))
		mBase = m.M
		v19 = m.ExcPending
		if v19 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v11)+8)) = v18
			*(*int32)(unsafe.Add(mBase, uint32(v18)+144)) = l2
			*(*int32)(unsafe.Add(mBase, uint32(v18)+140)) = l2
			*(*int32)(unsafe.Add(mBase, uint32(v18))) = l1
			*(*int32)(unsafe.Add(mBase, uint32(v18)+208)) = v11
			*(*int64)(unsafe.Add(mBase, uint32(v18)+56)) = l4 | int64(8)
			*(*int32)(unsafe.Add(mBase, uint32(v18)+48)) = int32(562)
			*(*int32)(unsafe.Add(mBase, uint32(v18)+28)) = int32(17)
			if l5 == int32(0) {
				*(*int64)(unsafe.Add(mBase, uint32(v18)+76)) = int64(0)
				v66 = v18
				F_populateCommandLegacyRangeSpec(m, v66)
				mBase = m.M
				v68 = m.ExcPending
				if v68 != 0 {
					return int32(0)
				} else {
					v69 = *(*int32)(unsafe.Add(mBase, uint32(v11)+8))
					v70 = int64(0)
					*(*int64)(unsafe.Add(mBase, uint32(v69)+104)) = v70
					*(*int64)(unsafe.Add(mBase, uint32(v69+int32(128)))) = v70
					*(*int64)(unsafe.Add(mBase, uint32(v69+int32(120)))) = v70
					*(*int64)(unsafe.Add(mBase, uint32(v69+int32(112)))) = v70
					return v11
				}
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v18)+80)) = int32(1)
				v37 = F_valkey_calloc(m, int32(48))
				mBase = m.M
				v38 = m.ExcPending
				if v38 != 0 {
					return int32(0)
				} else {
					v39 = *(*int32)(unsafe.Add(mBase, uint32(v11)+8))
					*(*int32)(unsafe.Add(mBase, uint32(v39)+76)) = v37
					v41 = int32(0)
					*(*int32)(unsafe.Add(mBase, uint32(v37)+40)) = v41
					*(*int32)(unsafe.Add(mBase, uint32(v37)+36)) = l7
					if l6 < v41 {
						v47 = v41
					} else {
						v47 = l5
					}
					*(*int32)(unsafe.Add(mBase, uint32(v37)+32)) = l6 - v47
					v50 = int32(2)
					*(*int32)(unsafe.Add(mBase, uint32(v37)+28)) = v50
					*(*int32)(unsafe.Add(mBase, uint32(v37)+20)) = l5
					*(*int32)(unsafe.Add(mBase, uint32(v37)+16)) = v50
					if l4&int64(2097152) == int64(0) {
						v61 = int64(50)
					} else {
						v61 = int64(1074)
					}
					*(*int64)(unsafe.Add(mBase, uint32(v37)+8)) = v61
					v66 = v39
					F_populateCommandLegacyRangeSpec(m, v66)
					mBase = m.M
					v68 = m.ExcPending
					if v68 != 0 {
						return int32(0)
					} else {
						v69 = *(*int32)(unsafe.Add(mBase, uint32(v11)+8))
						v70 = int64(0)
						*(*int64)(unsafe.Add(mBase, uint32(v69)+104)) = v70
						*(*int64)(unsafe.Add(mBase, uint32(v69+int32(128)))) = v70
						*(*int64)(unsafe.Add(mBase, uint32(v69+int32(120)))) = v70
						*(*int64)(unsafe.Add(mBase, uint32(v69+int32(112)))) = v70
						return v11
					}
				}
			}
		}
	}
}
func F_moduleEnqueueLoadModule(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
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
	var v32 int32
	_ = v32
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	v12 = F_valkey_malloc(m, int32(12))
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v14 = int32(0)
	if l2 == v14 {
		v21 = v14
		goto L3
	} else {
		goto L4
	}
L3:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+8)) = v21
	v23 = F_sdsnew(m, l0)
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		goto L1
	} else {
		goto L6
	}
L4:
	;
	v19 = F_valkey_malloc(m, l2<<(uint(int32(2))%32))
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		goto L1
	} else {
		goto L5
	}
L5:
	;
	v21 = v19
	goto L3
L6:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+4)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(v12))) = v23
	if l2 < int32(1) {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	v80 = *(*int32)(unsafe.Add(mBase, _consts[525]))
	v81 = F_listAddNodeTail(m, v80, v12)
	mBase = m.M
	v82 = m.ExcPending
	if v82 != 0 {
		goto L1
	} else {
		goto L19
	}
L8:
	;
	v32 = int32(0)
	goto L9
L9:
	;
	v40 = v32 << (uint(int32(2)) % 32)
	v42 = *(*int32)(unsafe.Add(mBase, uint32(l1+v40)))
	v45 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v42+int32(-1)))))
	switch v45 & int32(7) {
	case 0:
		goto L16
	case 1:
		goto L15
	case 2:
		goto L14
	case 3:
		goto L13
	case 4:
		goto L12
	default:
		v62 = int32(0)
		goto L11
	}
L10:
	;
	goto L7
L11:
	;
	v64 = F_createRawStringObject(m, v42, v62)
	mBase = m.M
	v65 = m.ExcPending
	if v65 != 0 {
		goto L1
	} else {
		goto L17
	}
L12:
	;
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v42+int32(-17))))
	v62 = v61
	goto L11
L13:
	;
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v42+int32(-9))))
	v62 = v58
	goto L11
L14:
	;
	v55 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v42+int32(-5)))))
	v62 = v55
	goto L11
L15:
	;
	v52 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v42+int32(-3)))))
	v62 = v52
	goto L11
L16:
	;
	v62 = int32(base.Ui32(v45) >> (uint(int32(3)) % 32))
	goto L11
L17:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21+v40))) = v64
	v68 = v32 + int32(1)
	if v68 != l2 {
		v32 = v68
		goto L9
	} else {
		goto L18
	}
L18:
	;
	goto L10
L19:
	;
	return
}
func F_moduleFireAuthenticationEvent(m *base.Module, l0 int64, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v15 int64
	_ = v15
	var v18 int32
	_ = v18
	var v20 int64
	_ = v20
	var v30 int64
	_ = v30
	var v36 int32
	_ = v36
	v5 = int32(0)
	v8 = m.G0
	v10 = v8 - int32(32)
	m.G0 = v10
	v13 = v10 + int32(24)
	v15 = *(*int64)(unsafe.Add(mBase, _consts[560]))
	*(*int64)(unsafe.Add(mBase, uint32(v13))) = v15
	v18 = v10 + int32(16)
	v20 = *(*int64)(unsafe.Add(mBase, _consts[561]))
	*(*int64)(unsafe.Add(mBase, uint32(v18))) = v20
	*(*int64)(unsafe.Add(mBase, uint32(v10+int32(8)))) = l0
	*(*int32)(unsafe.Add(mBase, uint32(v13))) = base.B2i32(l3 == v5)
	*(*int32)(unsafe.Add(mBase, uint32(v18))) = l1
	v30 = *(*int64)(unsafe.Add(mBase, _consts[562]))
	*(*int64)(unsafe.Add(mBase, uint32(v10))) = v30
	*(*int32)(unsafe.Add(mBase, uint32(v10)+20)) = l2
	F_moduleFireServerEvent(m, int64(18), v5, v10)
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
		return
	} else {
		m.G0 = v10 + int32(32)
		return
	}
}
func F_moduleFireCommandRejectedEvent(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v12 int32
	_ = v12
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v24 int64
	_ = v24
	var v28 int32
	_ = v28
	var v33 int64
	_ = v33
	var v35 int32
	_ = v35
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v51 int32
	_ = v51
	v3 = int32(0)
	v6 = m.G0
	v8 = v6 - int32(64)
	m.G0 = v8
	v12 = *(*int32)(unsafe.Add(mBase, _consts[579]))
	if v12 == v3 {
		m.G0 = v8 + int32(64)
		return
	} else {
		*(*int64)(unsafe.Add(mBase, uint32(v8)+8)) = int64(1)
		v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
		if v17 == int32(0) {
			v21 = v3
		} else {
			v20 = *(*int32)(unsafe.Add(mBase, uint32(v17)+140))
			v21 = v20
		}
		v24 = int64(0)
		*(*int64)(unsafe.Add(mBase, uint32(v8+int32(28)))) = v24
		v28 = int32(0)
		*(*int32)(unsafe.Add(mBase, uint32(v8+int32(36)))) = v28
		*(*int64)(unsafe.Add(mBase, uint32(v8)+20)) = v24
		*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = v21
		v33 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
		*(*int64)(unsafe.Add(mBase, uint32(v8)+40)) = v33
		v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)+200))
		*(*int32)(unsafe.Add(mBase, uint32(v8)+48)) = int32(base.Ui32(v35)>>(uint(int32(30))%32)) & int32(1)
		v41 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
		*(*int32)(unsafe.Add(mBase, uint32(v8)+52)) = v41
		v43 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
		*(*int32)(unsafe.Add(mBase, uint32(v8)+60)) = l1
		*(*int32)(unsafe.Add(mBase, uint32(v8)+56)) = v43
		F_moduleFireServerEvent(m, int64(22), v28, v8+int32(8))
		mBase = m.M
		v51 = m.ExcPending
		if v51 != 0 {
			return
		} else {
			m.G0 = v8 + int32(64)
			return
		}
	}
}
func F_moduleFireServerEvent(m *base.Module, l0 int64, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v47 int32
	_ = v47
	var v78 int32
	_ = v78
	var v93 int32
	_ = v93
	var v94 int64
	_ = v94
	var v96 int32
	_ = v96
	var v97 int64
	_ = v97
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v124 int64
	_ = v124
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v139 int32
	_ = v139
	var v141 int32
	_ = v141
	var v144 int32
	_ = v144
	var v148 int32
	_ = v148
	var v150 int32
	_ = v150
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v159 int32
	_ = v159
	var v163 int32
	_ = v163
	var v167 int32
	_ = v167
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v174 int64
	_ = v174
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v184 int64
	_ = v184
	var v188 int64
	_ = v188
	var v192 int32
	_ = v192
	var v196 int32
	_ = v196
	var v204 int64
	_ = v204
	var v206 int32
	_ = v206
	var v210 int64
	_ = v210
	var v214 int64
	_ = v214
	var v217 int32
	_ = v217
	var v225 int64
	_ = v225
	var v238 int32
	_ = v238
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v250 int32
	_ = v250
	var v253 int64
	_ = v253
	var v257 int64
	_ = v257
	var v261 int32
	_ = v261
	var v265 int32
	_ = v265
	var v273 int64
	_ = v273
	var v275 int32
	_ = v275
	var v279 int64
	_ = v279
	var v283 int64
	_ = v283
	var v286 int32
	_ = v286
	var v294 int64
	_ = v294
	var v308 int32
	_ = v308
	var v317 int32
	_ = v317
	var v318 int32
	_ = v318
	var v321 int64
	_ = v321
	var v324 int64
	_ = v324
	var v327 int32
	_ = v327
	var v329 int32
	_ = v329
	var v340 int32
	_ = v340
	var v342 int32
	_ = v342
	var v345 int32
	_ = v345
	var v346 int32
	_ = v346
	var v347 int32
	_ = v347
	var v348 int32
	_ = v348
	var v352 int32
	_ = v352
	var v354 int32
	_ = v354
	var v358 int32
	_ = v358
	var v359 int32
	_ = v359
	var v360 int32
	_ = v360
	var v361 int32
	_ = v361
	var v362 int32
	_ = v362
	var v363 int32
	_ = v363
	var v364 int32
	_ = v364
	var v368 int32
	_ = v368
	var v369 int32
	_ = v369
	var v373 int32
	_ = v373
	var v375 int32
	_ = v375
	var v379 int32
	_ = v379
	var v382 int32
	_ = v382
	var v395 int32
	_ = v395
	var v400 int32
	_ = v400
	var v401 int32
	_ = v401
	var v405 int32
	_ = v405
	var v406 int32
	_ = v406
	var v414 int64
	_ = v414
	var v416 int64
	_ = v416
	var v423 int32
	_ = v423
	var v424 int32
	_ = v424
	var v433 int32
	_ = v433
	var v443 int32
	_ = v443
	var v451 int32
	_ = v451
	var v453 int32
	_ = v453
	var v456 int32
	_ = v456
	var v462 int32
	_ = v462
	var v493 int32
	_ = v493
	v4 = int32(0)
	v20 = m.G0
	v22 = v20 - int32(352)
	m.G0 = v22
	v25 = *(*int32)(unsafe.Add(mBase, _consts[563]))
	v26 = *(*int32)(unsafe.Add(mBase, uint32(v25)+20))
	if v26 == v4 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	m.G0 = v22 + int32(352)
	return
L2:
	;
	v30 = v22 + int32(344)
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v25)))
	*(*int32)(unsafe.Add(mBase, uint32(v30)+4)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v30))) = v31
	goto L3
L3:
	;
	v36 = v22 + int32(344)
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v36)))
	if v38 == int32(0) {
		goto L5
	} else {
		goto L6
	}
L4:
	;
	if v38 == int32(0) {
		goto L1
	} else {
		goto L7
	}
L5:
	;
	goto L4
L6:
	;
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v36)+4))
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v38+base.B2i32(v41 == int32(0))<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v36))) = v47
	goto L5
L7:
	;
	v78 = v38
	goto L8
L8:
	;
	v93 = *(*int32)(unsafe.Add(mBase, uint32(v78)+8))
	v94 = *(*int64)(unsafe.Add(mBase, uint32(v93)+8))
	if v94 != l0 {
		goto L10
	} else {
		goto L11
	}
L9:
	;
	goto L1
L10:
	;
	v451 = v22 + int32(344)
	v453 = *(*int32)(unsafe.Add(mBase, uint32(v451)))
	if v453 == int32(0) {
		goto L67
	} else {
		goto L68
	}
L11:
	;
	v96 = *(*int32)(unsafe.Add(mBase, uint32(v93)))
	v97 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v22+int32(336)))) = v97
	*(*int64)(unsafe.Add(mBase, uint32(v22+int32(328)))) = v97
	*(*int64)(unsafe.Add(mBase, uint32(v22+int32(320)))) = v97
	*(*int64)(unsafe.Add(mBase, uint32(v22+int32(312)))) = v97
	*(*int64)(unsafe.Add(mBase, uint32(v22+int32(304)))) = v97
	*(*int64)(unsafe.Add(mBase, uint32(v22+int32(296)))) = v97
	*(*int64)(unsafe.Add(mBase, uint32(v22+int32(288)))) = v97
	*(*int64)(unsafe.Add(mBase, uint32(v22+int32(280)))) = v97
	*(*int32)(unsafe.Add(mBase, uint32(v22)+276)) = v96
	*(*int32)(unsafe.Add(mBase, uint32(v22)+272)) = int32(561)
	if l0 != int64(4) {
		goto L18
	} else {
		goto L19
	}
L12:
	;
	F_moduleFreeContext(m, v22+int32(272))
	mBase = m.M
	v443 = m.ExcPending
	if v443 != 0 {
		goto L25
	} else {
		goto L65
	}
L13:
	;
	v400 = *(*int32)(unsafe.Add(mBase, uint32(v93)))
	v401 = *(*int32)(unsafe.Add(mBase, uint32(v400)+44))
	*(*int32)(unsafe.Add(mBase, uint32(v400)+44)) = v401 + int32(1)
	v405 = *(*int32)(unsafe.Add(mBase, uint32(v93)+24))
	v406 = int32(8)
	v493 = int32(16)
	v414 = *(*int64)(unsafe.Add(mBase, uint32(v93+v493)))
	*(*int64)(unsafe.Add(mBase, uint32(v22+v493))) = v414
	v416 = *(*int64)(unsafe.Add(mBase, uint32(v93+v406)))
	*(*int64)(unsafe.Add(mBase, uint32(v22)+8)) = v416
	m.T0[v405].(func(*base.Module, int32, int32, int64, int32))(m, v22+int32(272), v22+v406, base.I64_extend_i32_s(l1), v395)
	mBase = m.M
	v423 = *(*int32)(unsafe.Add(mBase, uint32(v93)))
	v424 = *(*int32)(unsafe.Add(mBase, uint32(v423)+44))
	*(*int32)(unsafe.Add(mBase, uint32(v423)+44)) = v424 + int32(-1)
	if l0 != int64(17) {
		goto L12
	} else {
		goto L63
	}
L14:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v22)+328)) = v257
	v261 = int32(0)
	v265 = *(*int32)(unsafe.Add(mBase, _consts[95]))
	*(*int32)(unsafe.Add(mBase, _consts[95])) = v265 + int32(1)
	goto L39
L15:
	;
	v253 = *(*int64)(unsafe.Add(mBase, _consts[529]))
	v257 = v253*int64(1000) + v174
	goto L14
L16:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v22)+328)) = v188
	v192 = int32(0)
	v196 = *(*int32)(unsafe.Add(mBase, _consts[95]))
	*(*int32)(unsafe.Add(mBase, _consts[95])) = v196 + int32(1)
	goto L30
L17:
	;
	v184 = *(*int64)(unsafe.Add(mBase, _consts[529]))
	v188 = v184*int64(1000) + v124
	goto L16
L18:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+300)) = int32(64)
	v135 = int32(0)
	v136 = *(*int32)(unsafe.Add(mBase, _consts[546]))
	if v136 == v135 {
		goto L22
	} else {
		goto L23
	}
L19:
	;
	v121 = *(*int32)(unsafe.Add(mBase, _consts[131]))
	v122 = int32(0)
	v123 = *(*int32)(unsafe.Add(mBase, _consts[23]))
	v124 = m.T0[v123].(func(*base.Module) int64)(m)
	mBase = m.M
	if v121 == v122 {
		goto L17
	} else {
		goto L20
	}
L20:
	;
	v129 = *(*int32)(unsafe.Add(mBase, _consts[219]))
	v130 = base.I32_div_s(int32(1000000), v129)
	v188 = v124 + base.I64_extend_i32_s(v130)
	goto L16
L21:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+280)) = v167
	v171 = *(*int32)(unsafe.Add(mBase, _consts[131]))
	v172 = int32(0)
	v173 = *(*int32)(unsafe.Add(mBase, _consts[23]))
	v174 = m.T0[v173].(func(*base.Module) int64)(m)
	mBase = m.M
	if v171 == v172 {
		goto L15
	} else {
		goto L27
	}
L22:
	;
	v155 = F_createClient(m, int32(0))
	mBase = m.M
	v156 = m.ExcPending
	if v156 != 0 {
		goto L25
	} else {
		goto L26
	}
L23:
	;
	v139 = int32(0)
	v141 = v136 + int32(-1)
	*(*int32)(unsafe.Add(mBase, _consts[546])) = v141
	v144 = *(*int32)(unsafe.Add(mBase, _consts[547]))
	v148 = *(*int32)(unsafe.Add(mBase, uint32(v144+v141<<(uint(int32(2))%32))))
	v150 = *(*int32)(unsafe.Add(mBase, _consts[548]))
	if base.Ui32(v150) <= base.Ui32(v141) {
		v167 = v148
		goto L21
	} else {
		goto L24
	}
L24:
	;
	*(*int32)(unsafe.Add(mBase, _consts[548])) = v141
	v167 = v148
	goto L21
L25:
	;
	return
L26:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v155)+328)) = int32(0)
	v159 = *(*int32)(unsafe.Add(mBase, uint32(v155)+200))
	*(*int32)(unsafe.Add(mBase, uint32(v155)+200)) = v159 | int32(1073741824)
	v163 = *(*int32)(unsafe.Add(mBase, uint32(v155)+204))
	*(*int32)(unsafe.Add(mBase, uint32(v155)+204)) = v163 | int32(268435456)
	v167 = v155
	goto L21
L27:
	;
	v179 = *(*int32)(unsafe.Add(mBase, _consts[219]))
	v180 = base.I32_div_s(int32(1000000), v179)
	v257 = v174 + base.I64_extend_i32_s(v180)
	goto L14
L28:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+280)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(v22)+36)) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v22)+24)) = int64(1)
	*(*int32)(unsafe.Add(mBase, uint32(v22)+32)) = v22 + int32(40)
	v238 = v22 + int32(192)
	v241 = *(*int32)(unsafe.Add(mBase, uint32(v93)+16))
	v242 = F_modulePopulateClientInfoStructure(m, v238, l2, v241)
	mBase = m.M
	v243 = m.ExcPending
	if v243 != 0 {
		goto L25
	} else {
		goto L34
	}
L29:
	;
	goto L28
L30:
	;
	if v196 != 0 {
		goto L29
	} else {
		goto L31
	}
L31:
	;
	goto L33
L32:
	;
	v206 = int32(0)
	*(*int64)(unsafe.Add(mBase, _consts[96])) = v204
	v210 = base.I64_div_s(v204, int64(1000))
	*(*int64)(unsafe.Add(mBase, _consts[35])) = v210
	v214 = base.I64_div_s(v204, int64(1000000))
	*(*int64)(unsafe.Add(mBase, _consts[47])) = v214
	v217 = *(*int32)(unsafe.Add(mBase, _consts[97]))
	F_lrulfu_updateClockAndPolicy(m, v210, int32(base.Ui32(v217&int32(2))>>(uint(int32(1))%32)))
	mBase = m.M
	v225 = *(*int64)(unsafe.Add(mBase, _consts[35]))
	*(*int64)(unsafe.Add(mBase, _consts[98])) = v225
	goto L29
L33:
	;
	v204 = F_ustime(m)
	mBase = m.M
	goto L32
L34:
	;
	if v242 == int32(0) {
		v395 = v238
		goto L13
	} else {
		goto L35
	}
L35:
	;
	F__serverAssert(m, int32(_a943), int32(_a917), int32(12848))
	mBase = m.M
	v250 = m.ExcPending
	if v250 != 0 {
		goto L25
	} else {
		goto L36
	}
L36:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L37:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+36)) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v22)+24)) = int64(1)
	*(*int32)(unsafe.Add(mBase, uint32(v22)+32)) = v22 + int32(40)
	if base.Ui64(int64(23)) < base.Ui64(l0) {
		goto L43
	} else {
		goto L44
	}
L38:
	;
	goto L37
L39:
	;
	if v265 != 0 {
		goto L38
	} else {
		goto L40
	}
L40:
	;
	goto L42
L41:
	;
	v275 = int32(0)
	*(*int64)(unsafe.Add(mBase, _consts[96])) = v273
	v279 = base.I64_div_s(v273, int64(1000))
	*(*int64)(unsafe.Add(mBase, _consts[35])) = v279
	v283 = base.I64_div_s(v273, int64(1000000))
	*(*int64)(unsafe.Add(mBase, _consts[47])) = v283
	v286 = *(*int32)(unsafe.Add(mBase, _consts[97]))
	F_lrulfu_updateClockAndPolicy(m, v279, int32(base.Ui32(v286&int32(2))>>(uint(int32(1))%32)))
	mBase = m.M
	v294 = *(*int64)(unsafe.Add(mBase, _consts[35]))
	*(*int64)(unsafe.Add(mBase, _consts[98])) = v294
	goto L38
L42:
	;
	v273 = F_ustime(m)
	mBase = m.M
	goto L41
L43:
	;
	v395 = int32(0)
	goto L13
L44:
	;
	switch base.I32_wrap_i64(l0) {
	default:
		goto L48
	case 1, 3, 4, 5, 6, 7, 12, 13, 14, 15:
		goto L43
	case 2:
		goto L47
	case 8, 10, 11, 16, 18, 19, 20, 21, 22, 23:
		v395 = l2
		goto L13
	case 9:
		goto L46
	case 17:
		goto L45
	}
L45:
	;
	v358 = *(*int32)(unsafe.Add(mBase, uint32(v22)+280))
	v359 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	v360 = F_selectDb(m, v358, v359)
	mBase = m.M
	v361 = m.ExcPending
	if v361 != 0 {
		goto L25
	} else {
		goto L58
	}
L46:
	;
	v348 = *(*int32)(unsafe.Add(mBase, uint32(v93)))
	if l2 == v348 {
		goto L12
	} else {
		goto L57
	}
L47:
	;
	v342 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	if v342 == int32(-1) {
		v395 = l2
		goto L13
	} else {
		goto L55
	}
L48:
	;
	v308 = *(*int32)(unsafe.Add(mBase, uint32(v93)+16))
	if v308 != int32(1) {
		goto L49
	} else {
		goto L50
	}
L49:
	;
	F__serverAssert(m, int32(_a944), int32(_a917), int32(12851))
	mBase = m.M
	v340 = m.ExcPending
	if v340 != 0 {
		goto L25
	} else {
		goto L54
	}
L50:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v22)+168)) = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v22)+144)) = int64(1)
	*(*int32)(unsafe.Add(mBase, uint32(v22)+168)) = int32(_a945)
	*(*int32)(unsafe.Add(mBase, uint32(v22)+164)) = int32(_a946)
	v317 = int32(_a44)
	v318 = *(*int32)(unsafe.Add(mBase, _consts[290]))
	*(*int32)(unsafe.Add(mBase, uint32(v22)+160)) = v318
	v321 = *(*int64)(unsafe.Add(mBase, _consts[40]))
	*(*int64)(unsafe.Add(mBase, uint32(v22)+176)) = v321
	v324 = *(*int64)(unsafe.Add(mBase, _consts[385]))
	*(*int64)(unsafe.Add(mBase, uint32(v22)+184)) = v324
	v327 = *(*int32)(unsafe.Add(mBase, _consts[130]))
	if v327 != 0 {
		goto L51
	} else {
		goto L52
	}
L51:
	;
	v329 = v327
	goto L53
L52:
	;
	v329 = int32(_a139)
	goto L53
L53:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+156)) = v329
	*(*int32)(unsafe.Add(mBase, uint32(v22)+152)) = base.B2i32(v327 == int32(0))
	v395 = v22 + int32(144)
	goto L13
L54:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L55:
	;
	v345 = *(*int32)(unsafe.Add(mBase, uint32(v22)+280))
	v346 = F_selectDb(m, v345, v342)
	mBase = m.M
	v347 = m.ExcPending
	if v347 != 0 {
		goto L25
	} else {
		goto L56
	}
L56:
	;
	v395 = l2
	goto L13
L57:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v22)+128)) = int64(1)
	v352 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v22)+136)) = v352
	v354 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v22)+140)) = v354
	v395 = v22 + int32(128)
	goto L13
L58:
	;
	v362 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	v363 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	v364 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v22)+40)) = v22 + int32(272)
	v368 = *(*int32)(unsafe.Add(mBase, uint32(v22)+280))
	v369 = *(*int32)(unsafe.Add(mBase, uint32(v368)+96))
	*(*int32)(unsafe.Add(mBase, uint32(v22)+48)) = v364
	*(*int32)(unsafe.Add(mBase, uint32(v22)+44)) = v369
	F_incrRefCount(m, v364)
	mBase = m.M
	v373 = m.ExcPending
	if v373 != 0 {
		goto L25
	} else {
		goto L59
	}
L59:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+60)) = v363
	v375 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v22)+56)) = v375
	*(*int32)(unsafe.Add(mBase, uint32(v22)+52)) = v362
	v379 = v22 + int32(24)
	if v362 == v375 {
		v395 = v379
		goto L13
	} else {
		goto L60
	}
L60:
	;
	v382 = *(*int32)(unsafe.Add(mBase, uint32(v362)))
	switch v382&int32(15) + int32(-3) {
	case 0:
		goto L62
	default:
		v395 = v379
		goto L13
	case 3:
		goto L61
	}
L61:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+88)) = int32(0)
	v395 = v379
	goto L13
L62:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v22)+120)) = int64(4294967296)
	*(*int32)(unsafe.Add(mBase, uint32(v22)+64)) = int32(0)
	v395 = v379
	goto L13
L63:
	;
	F_moduleCloseKey(m, v22+int32(40))
	mBase = m.M
	v433 = m.ExcPending
	if v433 != 0 {
		goto L25
	} else {
		goto L64
	}
L64:
	;
	goto L12
L65:
	;
	goto L10
L66:
	;
	if v453 != 0 {
		v78 = v453
		goto L8
	} else {
		goto L69
	}
L67:
	;
	goto L66
L68:
	;
	v456 = *(*int32)(unsafe.Add(mBase, uint32(v451)+4))
	v462 = *(*int32)(unsafe.Add(mBase, uint32(v453+base.B2i32(v456 == int32(0))<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v451))) = v462
	goto L67
L69:
	;
	goto L9
}
func F_moduleFreeArgs(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v11 int32
	_ = v11
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v50 int32
	_ = v50
	v3 = int32(0)
	if l1 <= v3 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	F_valkey_free(m, l0)
	mBase = m.M
	v50 = m.ExcPending
	if v50 != 0 {
		goto L5
	} else {
		goto L16
	}
L2:
	;
	v11 = v3
	goto L3
L3:
	;
	v16 = l0 + v11*int32(44)
	v17 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
	F_valkey_free(m, v17)
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
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
	v20 = *(*int32)(unsafe.Add(mBase, uint32(v16)+12))
	F_valkey_free(m, v20)
	mBase = m.M
	v22 = m.ExcPending
	if v22 != 0 {
		goto L5
	} else {
		goto L7
	}
L7:
	;
	v23 = *(*int32)(unsafe.Add(mBase, uint32(v16)+16))
	F_valkey_free(m, v23)
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		goto L5
	} else {
		goto L8
	}
L8:
	;
	v26 = *(*int32)(unsafe.Add(mBase, uint32(v16)+20))
	F_valkey_free(m, v26)
	mBase = m.M
	v28 = m.ExcPending
	if v28 != 0 {
		goto L5
	} else {
		goto L9
	}
L9:
	;
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v16)+28))
	F_valkey_free(m, v29)
	mBase = m.M
	v31 = m.ExcPending
	if v31 != 0 {
		goto L5
	} else {
		goto L10
	}
L10:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v16)+40))
	F_valkey_free(m, v32)
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
		goto L5
	} else {
		goto L11
	}
L11:
	;
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v16)+36))
	if v35 == int32(0) {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	v42 = v11 + int32(1)
	if v42 != l1 {
		v11 = v42
		goto L3
	} else {
		goto L15
	}
L13:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v16)+32))
	F_moduleFreeArgs(m, v35, v38)
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
		goto L5
	} else {
		goto L14
	}
L14:
	;
	goto L12
L15:
	;
	goto L4
L16:
	;
	return
}
func F_moduleFreeCommand(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v81 int32
	_ = v81
	var v85 int32
	_ = v85
	var v88 int32
	_ = v88
	var v93 int32
	_ = v93
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v112 int32
	_ = v112
	var v114 int32
	_ = v114
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v121 int32
	_ = v121
	var v123 int32
	_ = v123
	var v126 int32
	_ = v126
	var v131 int32
	_ = v131
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v154 int32
	_ = v154
	var v157 int32
	_ = v157
	var v161 int32
	_ = v161
	var v164 int32
	_ = v164
	var v168 int32
	_ = v168
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v174 int32
	_ = v174
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v198 int32
	_ = v198
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v226 int32
	_ = v226
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v231 int32
	_ = v231
	var v233 int32
	_ = v233
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v251 int32
	_ = v251
	var v252 int32
	_ = v252
	var v254 int32
	_ = v254
	var v258 int32
	_ = v258
	var v271 int32
	_ = v271
	v9 = m.G0
	v11 = v9 - int32(64)
	m.G0 = v11
	v13 = int32(-1)
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
	if v14 != int32(562) {
		v258 = v13
		goto L2
	} else {
		goto L3
	}
L1:
	;
	F__serverAssert(m, int32(_a956), int32(_a917), int32(13278))
	mBase = m.M
	v271 = m.ExcPending
	if v271 != 0 {
		goto L11
	} else {
		goto L74
	}
L2:
	;
	m.G0 = v11 + int32(64)
	return v258
L3:
	;
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l1)+208))
	v18 = *(*int32)(unsafe.Add(mBase, uint32(v17)))
	if v18 != l0 {
		v258 = v13
		goto L2
	} else {
		goto L4
	}
L4:
	;
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l1)+80))
	if v20 < int32(1) {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	v64 = *(*int32)(unsafe.Add(mBase, uint32(l1)+76))
	F_valkey_free(m, v64)
	mBase = m.M
	v66 = m.ExcPending
	if v66 != 0 {
		goto L11
	} else {
		goto L17
	}
L6:
	;
	v29 = int32(0)
	goto L7
L7:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(l1)+76))
	v34 = v29 * int32(48)
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v32+v34)))
	if v36 == int32(0) {
		v44 = v32
		goto L9
	} else {
		goto L10
	}
L8:
	;
	goto L5
L9:
	;
	v45 = v44 + v34
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v45)+16))
	if v46 != int32(3) {
		goto L13
	} else {
		goto L14
	}
L10:
	;
	F_valkey_free(m, v36)
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	return int32(0)
L12:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(l1)+76))
	v44 = v43
	goto L9
L13:
	;
	v53 = v29 + int32(1)
	v54 = *(*int32)(unsafe.Add(mBase, uint32(l1)+80))
	if v53 < v54 {
		v29 = v53
		goto L7
	} else {
		goto L16
	}
L14:
	;
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v45)+20))
	F_valkey_free(m, v49)
	mBase = m.M
	v51 = m.ExcPending
	if v51 != 0 {
		goto L11
	} else {
		goto L15
	}
L15:
	;
	goto L13
L16:
	;
	goto L8
L17:
	;
	v68 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
	if v68 != 0 {
		goto L19
	} else {
		goto L20
	}
L18:
	;
	F_valkey_free(m, v93)
	mBase = m.M
	v99 = m.ExcPending
	if v99 != 0 {
		goto L11
	} else {
		goto L26
	}
L19:
	;
	v73 = v68
	v75 = int32(0)
	goto L21
L20:
	;
	v93 = int32(0)
	goto L18
L21:
	;
	v81 = *(*int32)(unsafe.Add(mBase, uint32(v73+v75<<(uint(int32(2))%32))))
	if v81 == int32(0) {
		v93 = v73
		goto L18
	} else {
		goto L23
	}
L22:
	;
	v93 = int32(0)
	goto L18
L23:
	;
	F_valkey_free(m, v81)
	mBase = m.M
	v85 = m.ExcPending
	if v85 != 0 {
		goto L11
	} else {
		goto L24
	}
L24:
	;
	v88 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
	if v88 != 0 {
		v73 = v88
		v75 = v75 + int32(1)
		goto L21
	} else {
		goto L25
	}
L25:
	;
	goto L22
L26:
	;
	v101 = *(*int32)(unsafe.Add(mBase, uint32(l1)+32))
	if v101 != 0 {
		goto L28
	} else {
		goto L29
	}
L27:
	;
	F_valkey_free(m, v131)
	mBase = m.M
	v137 = m.ExcPending
	if v137 != 0 {
		goto L11
	} else {
		goto L36
	}
L28:
	;
	v106 = v101
	v108 = int32(0)
	goto L30
L29:
	;
	v131 = int32(0)
	goto L27
L30:
	;
	v112 = v108 << (uint(int32(3)) % 32)
	v114 = *(*int32)(unsafe.Add(mBase, uint32(v106+v112)))
	if v114 == int32(0) {
		v131 = v106
		goto L27
	} else {
		goto L32
	}
L31:
	;
	v131 = int32(0)
	goto L27
L32:
	;
	F_valkey_free(m, v114)
	mBase = m.M
	v118 = m.ExcPending
	if v118 != 0 {
		goto L11
	} else {
		goto L33
	}
L33:
	;
	v119 = *(*int32)(unsafe.Add(mBase, uint32(l1)+32))
	v121 = *(*int32)(unsafe.Add(mBase, uint32(v119+v112)+4))
	F_valkey_free(m, v121)
	mBase = m.M
	v123 = m.ExcPending
	if v123 != 0 {
		goto L11
	} else {
		goto L34
	}
L34:
	;
	v126 = *(*int32)(unsafe.Add(mBase, uint32(l1)+32))
	if v126 != 0 {
		v106 = v126
		v108 = v108 + int32(1)
		goto L30
	} else {
		goto L35
	}
L35:
	;
	goto L31
L36:
	;
	v138 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	F_valkey_free(m, v138)
	mBase = m.M
	v140 = m.ExcPending
	if v140 != 0 {
		goto L11
	} else {
		goto L37
	}
L37:
	;
	v141 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	F_valkey_free(m, v141)
	mBase = m.M
	v143 = m.ExcPending
	if v143 != 0 {
		goto L11
	} else {
		goto L38
	}
L38:
	;
	v144 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
	F_valkey_free(m, v144)
	mBase = m.M
	v146 = m.ExcPending
	if v146 != 0 {
		goto L11
	} else {
		goto L39
	}
L39:
	;
	v147 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	F_valkey_free(m, v147)
	mBase = m.M
	v149 = m.ExcPending
	if v149 != 0 {
		goto L11
	} else {
		goto L40
	}
L40:
	;
	v150 = *(*int32)(unsafe.Add(mBase, uint32(l1)+148))
	if v150 == int32(0) {
		goto L41
	} else {
		goto L42
	}
L41:
	;
	v157 = *(*int32)(unsafe.Add(mBase, uint32(l1)+212))
	if v157 == int32(0) {
		goto L44
	} else {
		goto L45
	}
L42:
	;
	F_hdr_close(m, v150)
	mBase = m.M
	v154 = m.ExcPending
	if v154 != 0 {
		goto L11
	} else {
		goto L43
	}
L43:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+148)) = int32(0)
	goto L41
L44:
	;
	v164 = *(*int32)(unsafe.Add(mBase, uint32(l1)+216))
	if v164 == int32(0) {
		goto L47
	} else {
		goto L48
	}
L45:
	;
	F_sdsfree(m, v157)
	mBase = m.M
	v161 = m.ExcPending
	if v161 != 0 {
		goto L11
	} else {
		goto L46
	}
L46:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+212)) = int32(0)
	goto L44
L47:
	;
	v171 = *(*int32)(unsafe.Add(mBase, uint32(l1)+96))
	v172 = *(*int32)(unsafe.Add(mBase, uint32(l1)+88))
	F_moduleFreeArgs(m, v171, v172)
	mBase = m.M
	v174 = m.ExcPending
	if v174 != 0 {
		goto L11
	} else {
		goto L50
	}
L48:
	;
	F_sdsfree(m, v164)
	mBase = m.M
	v168 = m.ExcPending
	if v168 != 0 {
		goto L11
	} else {
		goto L49
	}
L49:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+216)) = int32(0)
	goto L47
L50:
	;
	F_valkey_free(m, v17)
	mBase = m.M
	v176 = m.ExcPending
	if v176 != 0 {
		goto L11
	} else {
		goto L51
	}
L51:
	;
	v177 = int32(0)
	v178 = *(*int32)(unsafe.Add(mBase, uint32(l1)+200))
	if v178 == v177 {
		v258 = v177
		goto L2
	} else {
		goto L52
	}
L52:
	;
	v182 = v11 + int32(16)
	v183 = int32(1)
	v184 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v182)+14)) = uint8(v184)
	*(*int32)(unsafe.Add(mBase, uint32(v182))) = v178
	*(*int32)(unsafe.Add(mBase, uint32(v182)+24)) = v184
	*(*uint8)(unsafe.Add(mBase, uint32(v182)+15)) = uint8(v183)
	*(*int32)(unsafe.Add(mBase, uint32(v182)+8)) = int32(-1)
	if v178 == v184 {
		goto L54
	} else {
		goto L55
	}
L53:
	;
	v205 = F_hashtableNext(m, v11+int32(16), v11+int32(12))
	mBase = m.M
	v206 = m.ExcPending
	if v206 != 0 {
		goto L11
	} else {
		goto L58
	}
L54:
	;
	goto L53
L55:
	;
	goto L56
L56:
	;
	v198 = *(*int32)(unsafe.Add(mBase, uint32(v178)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v182)+24)) = v198
	*(*int32)(unsafe.Add(mBase, uint32(v178)+40)) = v182
	goto L54
L57:
	;
	F_hashtableCleanupIterator(m, v11+int32(16))
	mBase = m.M
	v251 = m.ExcPending
	if v251 != 0 {
		goto L11
	} else {
		goto L72
	}
L58:
	;
	if v205 == int32(0) {
		goto L57
	} else {
		goto L59
	}
L59:
	;
	goto L60
L60:
	;
	v217 = *(*int32)(unsafe.Add(mBase, uint32(v11)+12))
	v218 = F_moduleFreeCommand(m, l0, v217)
	mBase = m.M
	v219 = m.ExcPending
	if v219 != 0 {
		goto L11
	} else {
		goto L63
	}
L61:
	;
	goto L57
L62:
	;
	v238 = F_hashtableNext(m, v11+int32(16), v11+int32(12))
	mBase = m.M
	v239 = m.ExcPending
	if v239 != 0 {
		goto L11
	} else {
		goto L70
	}
L63:
	;
	if v218 != 0 {
		goto L62
	} else {
		goto L64
	}
L64:
	;
	v220 = *(*int32)(unsafe.Add(mBase, uint32(l1)+200))
	v221 = *(*int32)(unsafe.Add(mBase, uint32(v217)))
	v222 = F_hashtableDelete(m, v220, v221)
	mBase = m.M
	v223 = m.ExcPending
	if v223 != 0 {
		goto L11
	} else {
		goto L65
	}
L65:
	;
	if v222 == int32(0) {
		goto L1
	} else {
		goto L66
	}
L66:
	;
	v226 = *(*int32)(unsafe.Add(mBase, uint32(v217)))
	F_sdsfree(m, v226)
	mBase = m.M
	v228 = m.ExcPending
	if v228 != 0 {
		goto L11
	} else {
		goto L67
	}
L67:
	;
	v229 = *(*int32)(unsafe.Add(mBase, uint32(v217)+140))
	F_sdsfree(m, v229)
	mBase = m.M
	v231 = m.ExcPending
	if v231 != 0 {
		goto L11
	} else {
		goto L68
	}
L68:
	;
	F_valkey_free(m, v217)
	mBase = m.M
	v233 = m.ExcPending
	if v233 != 0 {
		goto L11
	} else {
		goto L69
	}
L69:
	;
	goto L62
L70:
	;
	if v238 != 0 {
		goto L60
	} else {
		goto L71
	}
L71:
	;
	goto L61
L72:
	;
	v252 = *(*int32)(unsafe.Add(mBase, uint32(l1)+200))
	F_hashtableRelease(m, v252)
	mBase = m.M
	v254 = m.ExcPending
	if v254 != 0 {
		goto L11
	} else {
		goto L73
	}
L73:
	;
	v258 = v177
	goto L2
L74:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_moduleFreeListIterator(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v8 int32
	_ = v8
	var v12 int32
	_ = v12
	var v17 int32
	_ = v17
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v3 = *(*int32)(unsafe.Add(mBase, uint32(v2)))
	if v3&int32(15) != int32(1) {
		F__serverAssert(m, int32(_a926), int32(_a917), int32(811))
		mBase = m.M
		v17 = m.ExcPending
		if v17 != 0 {
			return
		} else {
			F_abort(m)
			mBase = m.M
			base.Wasm_trap_unreachable()
			for {
			}
		}
	} else {
		v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
		if v8 == int32(0) {
			return
		} else {
			F_moduleFreeKeyIterator(m, l0)
			mBase = m.M
			v12 = m.ExcPending
			if v12 != 0 {
				return
			} else {
				return
			}
		}
	}
}
func F_moduleGetCommandKeysViaAPI(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int64
	_ = v7
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v53 int64
	_ = v53
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v63 int64
	_ = v63
	var v67 int64
	_ = v67
	var v71 int32
	_ = v71
	var v75 int32
	_ = v75
	var v83 int64
	_ = v83
	var v85 int32
	_ = v85
	var v89 int64
	_ = v89
	var v93 int64
	_ = v93
	var v96 int32
	_ = v96
	var v104 int64
	_ = v104
	var v109 int32
	_ = v109
	var v112 int32
	_ = v112
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	v5 = int32(0)
	v7 = int64(0)
	v8 = m.G0
	v10 = v8 - int32(80)
	m.G0 = v10
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+208))
	v13 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
	*(*int64)(unsafe.Add(mBase, uint32(v10+int32(32)))) = v7
	*(*int64)(unsafe.Add(mBase, uint32(v10+int32(72)))) = v7
	*(*int64)(unsafe.Add(mBase, uint32(v10+int32(64)))) = v7
	*(*int64)(unsafe.Add(mBase, uint32(v10+int32(56)))) = v7
	*(*int64)(unsafe.Add(mBase, uint32(v10+int32(48)))) = v7
	*(*int64)(unsafe.Add(mBase, uint32(v10+int32(40)))) = v7
	*(*int64)(unsafe.Add(mBase, uint32(v10+int32(24)))) = v7
	*(*int64)(unsafe.Add(mBase, uint32(v10)+16)) = v7
	*(*int32)(unsafe.Add(mBase, uint32(v10)+12)) = v13
	*(*int32)(unsafe.Add(mBase, uint32(v10)+8)) = int32(561)
	*(*int32)(unsafe.Add(mBase, uint32(v10)+36)) = int32(2)
	v50 = *(*int32)(unsafe.Add(mBase, _consts[131]))
	v52 = *(*int32)(unsafe.Add(mBase, _consts[23]))
	v53 = m.T0[v52].(func(*base.Module) int64)(m)
	mBase = m.M
	if v50 == v5 {
		v63 = *(*int64)(unsafe.Add(mBase, _consts[529]))
		v67 = v63*int64(1000) + v53
	} else {
		v58 = *(*int32)(unsafe.Add(mBase, _consts[219]))
		v59 = base.I32_div_s(int32(1000000), v58)
		v67 = v53 + base.I64_extend_i32_s(v59)
	}
	*(*int64)(unsafe.Add(mBase, uint32(v10)+64)) = v67
	v71 = int32(0)
	v75 = *(*int32)(unsafe.Add(mBase, _consts[95]))
	*(*int32)(unsafe.Add(mBase, _consts[95])) = v75 + int32(1)
	if v75 != 0 {
	} else {
		v83 = F_ustime(m)
		mBase = m.M
		v85 = int32(0)
		*(*int64)(unsafe.Add(mBase, _consts[96])) = v83
		v89 = base.I64_div_s(v83, int64(1000))
		*(*int64)(unsafe.Add(mBase, _consts[35])) = v89
		v93 = base.I64_div_s(v83, int64(1000000))
		*(*int64)(unsafe.Add(mBase, _consts[47])) = v93
		v96 = *(*int32)(unsafe.Add(mBase, _consts[97]))
		F_lrulfu_updateClockAndPolicy(m, v89, int32(base.Ui32(v96&int32(2))>>(uint(int32(1))%32)))
		mBase = m.M
		v104 = *(*int64)(unsafe.Add(mBase, _consts[35]))
		*(*int64)(unsafe.Add(mBase, _consts[98])) = v104
	}
	v109 = F_getKeysPrepareResult(m, l3, int32(256))
	mBase = m.M
	v112 = m.ExcPending
	if v112 != 0 {
		return int32(0)
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(v10)+56)) = l3
		v116 = *(*int32)(unsafe.Add(mBase, uint32(v12)+4))
		v117 = m.T0[v116].(func(*base.Module, int32, int32, int32) int32)(m, v10+int32(8), l1, l2)
		mBase = m.M
		v118 = m.ExcPending
		if v118 != 0 {
			return int32(0)
		} else {
			F_moduleFreeContext(m, v10+int32(8))
			mBase = m.M
			v122 = m.ExcPending
			if v122 != 0 {
				return int32(0)
			} else {
				v123 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
				m.G0 = v10 + int32(80)
				return v123
			}
		}
	}
}
func F_moduleGetHandleByName(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	v3 = *(*int32)(unsafe.Add(mBase, _consts[541]))
	v4 = F_dictFetchValue(m, v3, l0)
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		return v4
	}
}
func F_moduleHandleBlockedClients(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v43 int32
	_ = v43
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v64 int64
	_ = v64
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v72 int64
	_ = v72
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v99 int64
	_ = v99
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v109 int64
	_ = v109
	var v113 int64
	_ = v113
	var v117 int32
	_ = v117
	var v121 int32
	_ = v121
	var v129 int64
	_ = v129
	var v131 int32
	_ = v131
	var v135 int64
	_ = v135
	var v139 int64
	_ = v139
	var v142 int32
	_ = v142
	var v150 int64
	_ = v150
	var v154 int32
	_ = v154
	var v156 int32
	_ = v156
	var v158 int32
	_ = v158
	var v162 int32
	_ = v162
	var v163 int64
	_ = v163
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v172 int32
	_ = v172
	var v173 int64
	_ = v173
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v182 int32
	_ = v182
	var v185 int32
	_ = v185
	var v189 int32
	_ = v189
	var v192 int32
	_ = v192
	var v195 int32
	_ = v195
	var v200 int32
	_ = v200
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v209 int32
	_ = v209
	var v213 int32
	_ = v213
	var v215 int32
	_ = v215
	var v219 int32
	_ = v219
	var v221 int32
	_ = v221
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v249 int32
	_ = v249
	var v253 int64
	_ = v253
	var v255 int32
	_ = v255
	var v256 int32
	_ = v256
	var v259 int32
	_ = v259
	var v261 int32
	_ = v261
	var v267 int32
	_ = v267
	var v269 int64
	_ = v269
	var v271 int32
	_ = v271
	var v274 int32
	_ = v274
	var v275 int32
	_ = v275
	var v276 int32
	_ = v276
	var v279 int32
	_ = v279
	var v282 int32
	_ = v282
	var v289 int32
	_ = v289
	var v291 int32
	_ = v291
	var v294 int32
	_ = v294
	var v296 int32
	_ = v296
	var v301 int32
	_ = v301
	var v303 int32
	_ = v303
	var v310 int32
	_ = v310
	var v313 int32
	_ = v313
	var v318 int32
	_ = v318
	var v319 int32
	_ = v319
	var v324 int32
	_ = v324
	var v332 int32
	_ = v332
	var v333 int32
	_ = v333
	v15 = m.G0
	v17 = v15 - int32(80)
	m.G0 = v17
	goto L1
L1:
	;
	v21 = int32(0)
	v22 = *(*int32)(unsafe.Add(mBase, _consts[557]))
	v23 = *(*int32)(unsafe.Add(mBase, uint32(v22)+20))
	if v23 == v21 {
		goto L2
	} else {
		goto L3
	}
L2:
	;
	goto L73
L3:
	;
	v43 = v22
	goto L4
L4:
	;
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v43)))
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v56)+8))
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v57)))
	F_listDelNode(m, v43, v56)
	mBase = m.M
	v60 = m.ExcPending
	if v60 != 0 {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	goto L2
L6:
	;
	return
L7:
	;
	goto L8
L8:
	;
	v64 = *(*int64)(unsafe.Add(mBase, _consts[94]))
	if v58 != 0 {
		goto L13
	} else {
		goto L14
	}
L9:
	;
	goto L71
L10:
	;
	v318 = *(*int32)(unsafe.Add(mBase, uint32(v57)+4))
	v319 = *(*int32)(unsafe.Add(mBase, uint32(v318)+52))
	*(*int32)(unsafe.Add(mBase, uint32(v318)+52)) = v319 + int32(-1)
	F_valkey_free(m, v57)
	mBase = m.M
	v324 = m.ExcPending
	if v324 != 0 {
		goto L6
	} else {
		goto L70
	}
L11:
	;
	F_commitDeferredReplyBuffer(m, v58, int32(0))
	mBase = m.M
	v232 = m.ExcPending
	if v232 != 0 {
		goto L6
	} else {
		goto L40
	}
L12:
	;
	v192 = *(*int32)(unsafe.Add(mBase, uint32(v57)+28))
	if v192 == int32(0) {
		goto L31
	} else {
		goto L32
	}
L13:
	;
	v66 = int32(0)
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v57)+44))
	if v67 != 0 {
		v180 = v66
		goto L15
	} else {
		goto L16
	}
L14:
	;
	v189 = int32(0)
	goto L12
L15:
	;
	v182 = *(*int32)(unsafe.Add(mBase, uint32(v58)+108))
	if v182 == int32(0) {
		v189 = v180
		goto L12
	} else {
		goto L29
	}
L16:
	;
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v57)+8))
	if v68 == int32(0) {
		v180 = v66
		goto L15
	} else {
		goto L17
	}
L17:
	;
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v57)+4))
	v72 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v17+int32(72)))) = v72
	*(*int64)(unsafe.Add(mBase, uint32(v17+int32(64)))) = v72
	*(*int64)(unsafe.Add(mBase, uint32(v17+int32(56)))) = v72
	*(*int64)(unsafe.Add(mBase, uint32(v17+int32(48)))) = v72
	*(*int64)(unsafe.Add(mBase, uint32(v17+int32(40)))) = v72
	*(*int64)(unsafe.Add(mBase, uint32(v17+int32(32)))) = v72
	*(*int64)(unsafe.Add(mBase, uint32(v17+int32(24)))) = v72
	*(*int64)(unsafe.Add(mBase, uint32(v17+int32(16)))) = v72
	*(*int32)(unsafe.Add(mBase, uint32(v17)+36)) = int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v17)+12)) = v71
	*(*int32)(unsafe.Add(mBase, uint32(v17)+8)) = int32(561)
	v96 = *(*int32)(unsafe.Add(mBase, _consts[131]))
	v97 = int32(0)
	v98 = *(*int32)(unsafe.Add(mBase, _consts[23]))
	v99 = m.T0[v98].(func(*base.Module) int64)(m)
	mBase = m.M
	if v96 == v97 {
		goto L19
	} else {
		goto L20
	}
L18:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v17)+64)) = v113
	v117 = int32(0)
	v121 = *(*int32)(unsafe.Add(mBase, _consts[95]))
	*(*int32)(unsafe.Add(mBase, _consts[95])) = v121 + int32(1)
	goto L23
L19:
	;
	v109 = *(*int64)(unsafe.Add(mBase, _consts[529]))
	v113 = v109*int64(1000) + v99
	goto L18
L20:
	;
	v104 = *(*int32)(unsafe.Add(mBase, _consts[219]))
	v105 = base.I32_div_s(int32(1000000), v104)
	v113 = v99 + base.I64_extend_i32_s(v105)
	goto L18
L21:
	;
	v154 = *(*int32)(unsafe.Add(mBase, uint32(v57)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v17)+48)) = v154
	v156 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v17)+52)) = v156
	v158 = *(*int32)(unsafe.Add(mBase, uint32(v57)))
	*(*int32)(unsafe.Add(mBase, uint32(v17)+20)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v17)+16)) = v158
	v162 = *(*int32)(unsafe.Add(mBase, _consts[23]))
	v163 = m.T0[v162].(func(*base.Module) int64)(m)
	mBase = m.M
	v166 = *(*int32)(unsafe.Add(mBase, uint32(v58)+20))
	v167 = *(*int32)(unsafe.Add(mBase, uint32(v58)+24))
	v168 = *(*int32)(unsafe.Add(mBase, uint32(v57)+8))
	v169 = m.T0[v168].(func(*base.Module, int32, int32, int32) int32)(m, v17+int32(8), v166, v167)
	mBase = m.M
	v170 = m.ExcPending
	if v170 != 0 {
		goto L6
	} else {
		goto L27
	}
L22:
	;
	goto L21
L23:
	;
	if v121 != 0 {
		goto L22
	} else {
		goto L24
	}
L24:
	;
	goto L26
L25:
	;
	v131 = int32(0)
	*(*int64)(unsafe.Add(mBase, _consts[96])) = v129
	v135 = base.I64_div_s(v129, int64(1000))
	*(*int64)(unsafe.Add(mBase, _consts[35])) = v135
	v139 = base.I64_div_s(v129, int64(1000000))
	*(*int64)(unsafe.Add(mBase, _consts[47])) = v139
	v142 = *(*int32)(unsafe.Add(mBase, _consts[97]))
	F_lrulfu_updateClockAndPolicy(m, v135, int32(base.Ui32(v142&int32(2))>>(uint(int32(1))%32)))
	mBase = m.M
	v150 = *(*int64)(unsafe.Add(mBase, _consts[35]))
	*(*int64)(unsafe.Add(mBase, _consts[98])) = v150
	goto L22
L26:
	;
	v129 = F_ustime(m)
	mBase = m.M
	goto L25
L27:
	;
	v172 = *(*int32)(unsafe.Add(mBase, _consts[23]))
	v173 = m.T0[v172].(func(*base.Module) int64)(m)
	mBase = m.M
	F_moduleFreeContext(m, v17+int32(8))
	mBase = m.M
	v179 = m.ExcPending
	if v179 != 0 {
		goto L6
	} else {
		goto L28
	}
L28:
	;
	v180 = base.I32_wrap_i64(v173 - v163)
	goto L15
L29:
	;
	v185 = *(*int32)(unsafe.Add(mBase, uint32(v182)+4))
	if v185 == int32(0) {
		v189 = v180
		goto L12
	} else {
		goto L30
	}
L30:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v182))) = v57
	v227 = v180
	goto L11
L31:
	;
	if v58 != 0 {
		v227 = v189
		goto L11
	} else {
		goto L37
	}
L32:
	;
	v195 = *(*int32)(unsafe.Add(mBase, uint32(v57)+24))
	if v195 == int32(0) {
		goto L31
	} else {
		goto L33
	}
L33:
	;
	v200 = *(*int32)(unsafe.Add(mBase, uint32(v57)+4))
	F_moduleCreateContext(m, v17+int32(8), v200, base.B2i32(v58 == int32(0))<<(uint(int32(5))%32))
	mBase = m.M
	v206 = m.ExcPending
	if v206 != 0 {
		goto L6
	} else {
		goto L34
	}
L34:
	;
	v207 = *(*int32)(unsafe.Add(mBase, uint32(v57)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v17)+48)) = v207
	v209 = *(*int32)(unsafe.Add(mBase, uint32(v57)))
	*(*int32)(unsafe.Add(mBase, uint32(v17)+16)) = v209
	v213 = *(*int32)(unsafe.Add(mBase, uint32(v57)+24))
	m.T0[v213].(func(*base.Module, int32, int32))(m, v17+int32(8), v207)
	mBase = m.M
	v215 = m.ExcPending
	if v215 != 0 {
		goto L6
	} else {
		goto L35
	}
L35:
	;
	F_moduleFreeContext(m, v17+int32(8))
	mBase = m.M
	v219 = m.ExcPending
	if v219 != 0 {
		goto L6
	} else {
		goto L36
	}
L36:
	;
	goto L31
L37:
	;
	v221 = *(*int32)(unsafe.Add(mBase, uint32(v57)+36))
	F_moduleReleaseTempClient(m, v221)
	mBase = m.M
	v223 = m.ExcPending
	if v223 != 0 {
		goto L6
	} else {
		goto L38
	}
L38:
	;
	v224 = *(*int32)(unsafe.Add(mBase, uint32(v57)+32))
	F_moduleReleaseTempClient(m, v224)
	mBase = m.M
	v226 = m.ExcPending
	if v226 != 0 {
		goto L6
	} else {
		goto L39
	}
L39:
	;
	goto L10
L40:
	;
	v233 = *(*int32)(unsafe.Add(mBase, uint32(v57)+36))
	F_AddReplyFromClient(m, v58, v233)
	mBase = m.M
	v235 = m.ExcPending
	if v235 != 0 {
		goto L6
	} else {
		goto L41
	}
L41:
	;
	v236 = *(*int32)(unsafe.Add(mBase, uint32(v57)+36))
	F_moduleReleaseTempClient(m, v236)
	mBase = m.M
	v238 = m.ExcPending
	if v238 != 0 {
		goto L6
	} else {
		goto L42
	}
L42:
	;
	v239 = *(*int32)(unsafe.Add(mBase, uint32(v57)+32))
	F_moduleReleaseTempClient(m, v239)
	mBase = m.M
	v241 = m.ExcPending
	if v241 != 0 {
		goto L6
	} else {
		goto L43
	}
L43:
	;
	v242 = *(*int32)(unsafe.Add(mBase, uint32(v58)+108))
	if v242 == int32(0) {
		goto L45
	} else {
		goto L46
	}
L44:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v57)+20)) = int32(0)
	F_unblockClient(m, v58, int32(1))
	mBase = m.M
	v267 = m.ExcPending
	if v267 != 0 {
		goto L6
	} else {
		goto L55
	}
L45:
	;
	v246 = *(*int32)(unsafe.Add(mBase, uint32(v58)+344))
	if v246 == int32(0) {
		goto L49
	} else {
		goto L50
	}
L46:
	;
	v245 = *(*int32)(unsafe.Add(mBase, uint32(v242)+4))
	if v245 != 0 {
		goto L44
	} else {
		goto L47
	}
L47:
	;
	goto L45
L48:
	;
	v256 = *(*int32)(unsafe.Add(mBase, uint32(v57)+64))
	if v255 != 0 {
		goto L51
	} else {
		goto L52
	}
L49:
	;
	v253 = *(*int64)(unsafe.Add(mBase, _consts[94]))
	v255 = base.B2i32(v253 != v64)
	goto L48
L50:
	;
	v249 = *(*int32)(unsafe.Add(mBase, uint32(v246)+20))
	v255 = base.B2i32(v249 != int32(0))
	goto L48
L51:
	;
	v259 = int32(2)
	goto L53
L52:
	;
	v259 = int32(0)
	goto L53
L53:
	;
	F_updateStatsOnUnblock(m, v58, v256, v227, v259)
	mBase = m.M
	v261 = m.ExcPending
	if v261 != 0 {
		goto L6
	} else {
		goto L54
	}
L54:
	;
	goto L44
L55:
	;
	v269 = *(*int64)(unsafe.Add(mBase, _consts[40]))
	*(*int64)(unsafe.Add(mBase, uint32(v58)+48)) = v269
	v271 = *(*int32)(unsafe.Add(mBase, uint32(v58)+108))
	if v271 == int32(0) {
		goto L57
	} else {
		goto L58
	}
L56:
	;
	v310 = *(*int32)(unsafe.Add(mBase, uint32(v58)+108))
	if v310 == int32(0) {
		goto L10
	} else {
		goto L68
	}
L57:
	;
	v275 = F_clientHasPendingReplies(m, v58)
	mBase = m.M
	v276 = m.ExcPending
	if v276 != 0 {
		goto L6
	} else {
		goto L60
	}
L58:
	;
	v274 = *(*int32)(unsafe.Add(mBase, uint32(v271)+4))
	if v274 != 0 {
		goto L56
	} else {
		goto L59
	}
L59:
	;
	goto L57
L60:
	;
	if v275 == int32(0) {
		goto L56
	} else {
		goto L61
	}
L61:
	;
	v279 = *(*int32)(unsafe.Add(mBase, uint32(v58)+200))
	if v279&int32(4194304) != 0 {
		goto L56
	} else {
		goto L62
	}
L62:
	;
	v282 = *(*int32)(unsafe.Add(mBase, uint32(v58)+8))
	if v282 == int32(0) {
		goto L56
	} else {
		goto L63
	}
L63:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v58)+200)) = v279 | int32(4194304)
	v289 = *(*int32)(unsafe.Add(mBase, _consts[216]))
	v291 = v58 + int32(168)
	v294 = *(*int32)(unsafe.Add(mBase, uint32(v289)+20))
	if v294 != 0 {
		goto L66
	} else {
		goto L67
	}
L64:
	;
	goto L56
L65:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v289))) = v291
	*(*int32)(unsafe.Add(mBase, uint32(v291)+4)) = v303
	*(*int32)(unsafe.Add(mBase, uint32(v289)+20)) = v294 + int32(1)
	goto L64
L66:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v291))) = int32(0)
	v301 = *(*int32)(unsafe.Add(mBase, uint32(v289)))
	*(*int32)(unsafe.Add(mBase, uint32(v301))) = v291
	v303 = v301
	goto L65
L67:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v289)+4)) = v291
	v296 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v291))) = v296
	v303 = v296
	goto L65
L68:
	;
	v313 = *(*int32)(unsafe.Add(mBase, uint32(v310)+4))
	if v313 != 0 {
		goto L9
	} else {
		goto L69
	}
L69:
	;
	goto L10
L70:
	;
	goto L9
L71:
	;
	v332 = *(*int32)(unsafe.Add(mBase, _consts[557]))
	v333 = *(*int32)(unsafe.Add(mBase, uint32(v332)+20))
	if v333 != 0 {
		v43 = v332
		goto L4
	} else {
		goto L72
	}
L72:
	;
	goto L5
L73:
	;
	m.G0 = v17 + int32(80)
	return
}
func F_moduleIsModuleCommand(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	v3 = int32(0)
	if l0 == v3 {
		v13 = v3
	} else {
		v7 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
		if v7 != int32(562) {
			v13 = v3
		} else {
			v10 = *(*int32)(unsafe.Add(mBase, uint32(l1)+208))
			v11 = *(*int32)(unsafe.Add(mBase, uint32(v10)))
			v13 = base.B2i32(v11 == l0)
		}
	}
	return v13
}
func F_moduleListConfigMatch(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v9 int32
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
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v6 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3))))
	if v6 != 0 {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	return base.B2i32(v38-v40 == int32(0))
L2:
	;
	v38 = F_tolower(m, v34)
	mBase = m.M
	v39 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v35))))
	v40 = F_tolower(m, v39)
	mBase = m.M
	goto L1
L3:
	;
	v8 = v3
	v9 = l1
	v10 = v6
	goto L6
L4:
	;
	v34 = int32(0)
	v35 = l1
	goto L2
L5:
	;
	v34 = v31 & int32(255)
	v35 = v30
	goto L2
L6:
	;
	v12 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9))))
	if v12 == int32(0) {
		v30 = v9
		v31 = v10
		goto L5
	} else {
		goto L8
	}
L7:
	;
	v30 = v24
	v31 = int32(0)
	goto L5
L8:
	;
	v16 = v10 & int32(255)
	if v16 == v12 {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	v23 = int32(1)
	v24 = v9 + v23
	v25 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+1)))
	if v25 != 0 {
		v8 = v8 + v23
		v9 = v24
		v10 = v25
		goto L6
	} else {
		goto L12
	}
L10:
	;
	v18 = F_tolower(m, v16)
	mBase = m.M
	v19 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9))))
	v20 = F_tolower(m, v19)
	mBase = m.M
	if v18 == v20 {
		goto L9
	} else {
		goto L11
	}
L11:
	;
	v22 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8))))
	v30 = v9
	v31 = v22
	goto L5
L12:
	;
	goto L7
}
func F_moduleListFree(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	F_sdsfree(m, v2)
	mBase = m.M
	v4 = m.ExcPending
	if v4 != 0 {
		return
	} else {
		F_valkey_free(m, l0)
		mBase = m.M
		v6 = m.ExcPending
		if v6 != 0 {
			return
		} else {
			return
		}
	}
}
func F_moduleListIteratorSeek(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v26 int32
	_ = v26
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v88 int32
	_ = v88
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v109 int32
	_ = v109
	var v118 int32
	_ = v118
	var v124 int32
	_ = v124
	var v130 int32
	_ = v130
	if l0 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v13 == int32(0) {
		goto L5
	} else {
		goto L6
	}
L2:
	;
	goto L3
L3:
	;
	*(*int32)(unsafe.Add(mBase, _consts[5])) = int32(28)
	return int32(0)
L4:
	;
	v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v26&l2 != 0 {
		goto L9
	} else {
		goto L10
	}
L5:
	;
	goto L8
L6:
	;
	v16 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
	if v16&int32(15) == int32(1) {
		goto L4
	} else {
		goto L7
	}
L7:
	;
	goto L5
L8:
	;
	*(*int32)(unsafe.Add(mBase, _consts[5])) = int32(138)
	return int32(0)
L9:
	;
	v34 = F_listTypeLength(m, v13)
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		goto L14
	} else {
		goto L15
	}
L10:
	;
	goto L11
L11:
	;
	*(*int32)(unsafe.Add(mBase, _consts[5])) = int32(8)
	return int32(0)
L12:
	;
	v46 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v46 != 0 {
		goto L22
	} else {
		goto L23
	}
L13:
	;
	goto L18
L14:
	;
	return int32(0)
L15:
	;
	if l1 < int32(0)-v34 {
		goto L13
	} else {
		goto L16
	}
L16:
	;
	if l1 < v34 {
		goto L12
	} else {
		goto L17
	}
L17:
	;
	goto L13
L18:
	;
	*(*int32)(unsafe.Add(mBase, _consts[5])) = int32(18)
	return int32(0)
L19:
	;
	F__serverAssert(m, int32(_a924), int32(_a917), int32(4713))
	mBase = m.M
	v130 = m.ExcPending
	if v130 != 0 {
		goto L14
	} else {
		goto L48
	}
L20:
	;
	F__serverAssert(m, int32(_a924), int32(_a917), int32(4695))
	mBase = m.M
	v124 = m.ExcPending
	if v124 != 0 {
		goto L14
	} else {
		goto L47
	}
L21:
	;
	F__serverAssert(m, int32(_a925), int32(_a917), int32(4694))
	mBase = m.M
	v118 = m.ExcPending
	if v118 != 0 {
		goto L14
	} else {
		goto L46
	}
L22:
	;
	v63 = int32(1)
	v64 = int32(0)
	v65 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
	if v65 < v64 {
		goto L29
	} else {
		goto L30
	}
L23:
	;
	v47 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v49 = F_listTypeInitIterator(m, v47, l1, int32(1))
	mBase = m.M
	v50 = m.ExcPending
	if v50 != 0 {
		goto L14
	} else {
		goto L24
	}
L24:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v49
	if v49 == int32(0) {
		goto L21
	} else {
		goto L25
	}
L25:
	;
	v56 = F_listTypeNext(m, v49, l0+int32(24))
	mBase = m.M
	v57 = m.ExcPending
	if v57 != 0 {
		goto L14
	} else {
		goto L26
	}
L26:
	;
	if v56 == int32(0) {
		goto L20
	} else {
		goto L27
	}
L27:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+64)) = l1
	return int32(1)
L28:
	;
	return v109
L29:
	;
	v68 = v64
	goto L31
L30:
	;
	v68 = v34
	goto L31
L31:
	;
	v69 = int32(0)
	if l1 < v69 {
		goto L32
	} else {
		goto L33
	}
L32:
	;
	v76 = v68
	goto L34
L33:
	;
	v76 = v69 - v65>>(uint(int32(31))%32)&v34
	goto L34
L34:
	;
	v77 = l1 + v76
	if v77 == v65 {
		v109 = v63
		goto L28
	} else {
		goto L35
	}
L35:
	;
	v80 = l0 + int32(24)
	v81 = base.B2i32(v65 < v77)
	F_listTypeSetIteratorDirection(m, v46, v80, v81)
	mBase = m.M
	v83 = m.ExcPending
	if v83 != 0 {
		goto L14
	} else {
		goto L36
	}
L36:
	;
	v84 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
	if v84 == v77 {
		v109 = v63
		goto L28
	} else {
		goto L37
	}
L37:
	;
	if v65 < v77 {
		goto L38
	} else {
		goto L39
	}
L38:
	;
	v88 = int32(1)
	goto L40
L39:
	;
	v88 = int32(-1)
	goto L40
L40:
	;
	goto L41
L41:
	;
	v96 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v97 = F_listTypeNext(m, v96, v80)
	mBase = m.M
	v98 = m.ExcPending
	if v98 != 0 {
		goto L14
	} else {
		goto L43
	}
L42:
	;
	v109 = int32(1)
	goto L28
L43:
	;
	if v97 == int32(0) {
		goto L19
	} else {
		goto L44
	}
L44:
	;
	v101 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
	v102 = v101 + v88
	*(*int32)(unsafe.Add(mBase, uint32(l0)+64)) = v102
	if v102 != v77 {
		goto L41
	} else {
		goto L45
	}
L45:
	;
	goto L42
L46:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L47:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L48:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_moduleLoadStatic(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	*(*int32)(unsafe.Add(mBase, uint32(v9)+8)) = int32(0)
	v19 = F_moduleLoadStaticSymbol(m, v9+int32(12), v9+int32(8), int32(_a959), l0)
	mBase = m.M
	v22 = m.ExcPending
	if v22 != 0 {
		return int32(0)
	} else {
		if v19 != 0 {
			v28 = int32(-1)
			m.G0 = v9 + int32(16)
			return v28
		} else {
			v23 = *(*int32)(unsafe.Add(mBase, uint32(v9)+12))
			v24 = *(*int32)(unsafe.Add(mBase, uint32(v9)+8))
			v26 = F_moduleInitPostOnLoadResolved(m, v23, v24, l0, l1, l2, l3, int32(1))
			mBase = m.M
			v27 = m.ExcPending
			if v27 != 0 {
				return int32(0)
			} else {
				v28 = v26
				m.G0 = v9 + int32(16)
				return v28
			}
		}
	}
}
func F_moduleLoadStaticSymbol(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v34 int32
	_ = v34
	var v65 int32
	_ = v65
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v89 int32
	_ = v89
	var v94 int32
	_ = v94
	var v104 int32
	_ = v104
	var v109 int32
	_ = v109
	v7 = m.G0
	v9 = v7 - int32(176)
	m.G0 = v9
	*(*int32)(unsafe.Add(mBase, uint32(v9)+32)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(v9)+36)) = l3
	v19 = F_snprintf(m, v9+int32(48), int32(128), int32(_a960), v9+int32(32))
	mBase = m.M
	v22 = m.ExcPending
	if v22 != 0 {
		return int32(0)
	} else {
		if v19 < int32(128) {
			*(*int32)(unsafe.Add(mBase, uint32(l1))) = int32(_a961)
			v65 = v9 + int32(48)
			v69 = F_strcmp(m, int32(_a962), v65)
			mBase = m.M
			if v69 != 0 {
				v73 = F_strcmp(m, int32(_a963), v65)
				mBase = m.M
				if v73 != 0 {
					v76 = int32(0)
					*(*int32)(unsafe.Add(mBase, _consts[2])) = int32(_a964)
					v80 = v76
				} else {
					v75 = *(*int32)(unsafe.Add(mBase, _consts[584]))
					v80 = v75
				}
			} else {
				v71 = *(*int32)(unsafe.Add(mBase, _consts[585]))
				v80 = v71
			}
			*(*int32)(unsafe.Add(mBase, uint32(l0))) = v80
			if v80 != 0 {
				v109 = int32(0)
				m.G0 = v9 + int32(176)
				return v109
			} else {
				v82 = int32(0)
				v84 = *(*int32)(unsafe.Add(mBase, _consts[2]))
				*(*int32)(unsafe.Add(mBase, _consts[2])) = v82
				v89 = *(*int32)(unsafe.Add(mBase, _consts[6]))
				if int32(3) < v89 {
					v109 = int32(-1)
					m.G0 = v9 + int32(176)
					return v109
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = l3
					if v84 != 0 {
						v94 = v84
					} else {
						v94 = int32(_a965)
					}
					*(*int32)(unsafe.Add(mBase, uint32(v9)+24)) = v94
					*(*int32)(unsafe.Add(mBase, uint32(v9)+20)) = v9 + int32(48)
					F__serverLog(m, int32(3), int32(_a966), v9+int32(16))
					mBase = m.M
					v104 = m.ExcPending
					if v104 != 0 {
						return int32(0)
					} else {
						v109 = int32(-1)
						m.G0 = v9 + int32(176)
						return v109
					}
				}
			}
		} else {
			v25 = int32(-1)
			v27 = *(*int32)(unsafe.Add(mBase, _consts[6]))
			if int32(3) < v27 {
				v109 = v25
				m.G0 = v9 + int32(176)
				return v109
			} else {
				F__serverLog(m, int32(3), int32(_a967), int32(0))
				mBase = m.M
				v34 = m.ExcPending
				if v34 != 0 {
					return int32(0)
				} else {
					v109 = v25
					m.G0 = v9 + int32(176)
					return v109
				}
			}
		}
	}
}
func F_moduleLoadString(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v15 int64
	_ = v15
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v34 int64
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v57 int32
	_ = v57
	var v63 int32
	_ = v63
	v7 = m.G0
	v9 = v7 - int32(32)
	m.G0 = v9
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v12 != 0 {
		v63 = int32(0)
		m.G0 = v9 + int32(32)
		return v63
	} else {
		v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		v15 = F_rdbLoadLen(m, v13, int32(0))
		mBase = m.M
		v18 = m.ExcPending
		if v18 != 0 {
			return int32(0)
		} else {
			if v15 != int64(5) {
				v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				v30 = *(*int32)(unsafe.Add(mBase, uint32(v29)+8))
				v31 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v30)+48)))
				if v31&int32(1) != 0 {
					*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = int32(1)
					v63 = int32(0)
					m.G0 = v9 + int32(32)
					return v63
				} else {
					v34 = int64(*(*uint32)(unsafe.Add(mBase, uint32(l0))))
					v35 = *(*int32)(unsafe.Add(mBase, uint32(v30)+4))
					v36 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
					if v36 != 0 {
						v39 = *(*int32)(unsafe.Add(mBase, _consts[209]))
						if v39 == int32(0) {
							v43 = F_objectGetVal(m, v36)
							mBase = m.M
							v44 = v43
						} else {
							v44 = int32(_a938)
						}
					} else {
						v44 = int32(_a939)
					}
					*(*int32)(unsafe.Add(mBase, uint32(v9+int32(16)))) = v44
					*(*int64)(unsafe.Add(mBase, uint32(v9)+8)) = v34
					*(*int32)(unsafe.Add(mBase, uint32(v9))) = v35
					*(*int32)(unsafe.Add(mBase, uint32(v9)+4)) = v29 + int32(84)
					F__serverPanic_1(m, int32(_a917), int32(7585), int32(_a940), v9)
					mBase = m.M
					v57 = m.ExcPending
					if v57 != 0 {
						return int32(0)
					} else {
						F_abort(m)
						mBase = m.M
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			} else {
				v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
				v26 = F_rdbGenericLoadStringObject(m, v21, base.B2i32(l1 != int32(0))<<(uint(int32(1))%32), l2)
				mBase = m.M
				v27 = m.ExcPending
				if v27 != 0 {
					return int32(0)
				} else {
					if v26 != 0 {
						v63 = v26
						m.G0 = v9 + int32(32)
						return v63
					} else {
						v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
						v30 = *(*int32)(unsafe.Add(mBase, uint32(v29)+8))
						v31 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v30)+48)))
						if v31&int32(1) != 0 {
							*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = int32(1)
							v63 = int32(0)
							m.G0 = v9 + int32(32)
							return v63
						} else {
							v34 = int64(*(*uint32)(unsafe.Add(mBase, uint32(l0))))
							v35 = *(*int32)(unsafe.Add(mBase, uint32(v30)+4))
							v36 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
							if v36 != 0 {
								v39 = *(*int32)(unsafe.Add(mBase, _consts[209]))
								if v39 == int32(0) {
									v43 = F_objectGetVal(m, v36)
									mBase = m.M
									v44 = v43
								} else {
									v44 = int32(_a938)
								}
							} else {
								v44 = int32(_a939)
							}
							*(*int32)(unsafe.Add(mBase, uint32(v9+int32(16)))) = v44
							*(*int64)(unsafe.Add(mBase, uint32(v9)+8)) = v34
							*(*int32)(unsafe.Add(mBase, uint32(v9))) = v35
							*(*int32)(unsafe.Add(mBase, uint32(v9)+4)) = v29 + int32(84)
							F__serverPanic_1(m, int32(_a917), int32(7585), int32(_a940), v9)
							mBase = m.M
							v57 = m.ExcPending
							if v57 != 0 {
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
func F_moduleNotifyKeyUnlink(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v29 int32
	_ = v29
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	v7 = m.G0
	v9 = v7 - int32(32)
	m.G0 = v9
	v11 = int32(_a44)
	v13 = *(*int32)(unsafe.Add(mBase, _consts[582]))
	v14 = int32(1)
	*(*int32)(unsafe.Add(mBase, _consts[582])) = v13 + v14
	if l3&int32(2) != 0 {
		v29 = v14
	} else {
		if l3&int32(4) != 0 {
			v29 = int32(2)
		} else {
			v29 = l3 << (uint(int32(28)) % 32) >> (uint(int32(31)) % 32) & int32(3)
		}
	}
	*(*int32)(unsafe.Add(mBase, uint32(v9)+28)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v9)+20)) = l0
	*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(v9)+24)) = l1
	F_moduleFireServerEvent(m, int64(17), v29, v9+int32(16))
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		return
	} else {
		v40 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
		if v40&int32(15) != int32(5) {
			v68 = int32(_a44)
			v70 = *(*int32)(unsafe.Add(mBase, _consts[582]))
			*(*int32)(unsafe.Add(mBase, _consts[582])) = v70 + int32(-1)
			m.G0 = v9 + int32(32)
			return
		} else {
			v45 = F_objectGetVal(m, l1)
			mBase = m.M
			v46 = *(*int32)(unsafe.Add(mBase, uint32(v45)))
			v47 = *(*int32)(unsafe.Add(mBase, uint32(v46)+68))
			if v47 == int32(0) {
				v59 = *(*int32)(unsafe.Add(mBase, uint32(v46)+40))
				if v59 == int32(0) {
					v68 = int32(_a44)
					v70 = *(*int32)(unsafe.Add(mBase, _consts[582]))
					*(*int32)(unsafe.Add(mBase, _consts[582])) = v70 + int32(-1)
					m.G0 = v9 + int32(32)
					return
				} else {
					v62 = *(*int32)(unsafe.Add(mBase, uint32(v45)+4))
					m.T0[v59].(func(*base.Module, int32, int32))(m, l0, v62)
					mBase = m.M
					v64 = m.ExcPending
					if v64 != 0 {
						return
					} else {
						v68 = int32(_a44)
						v70 = *(*int32)(unsafe.Add(mBase, _consts[582]))
						*(*int32)(unsafe.Add(mBase, _consts[582])) = v70 + int32(-1)
						m.G0 = v9 + int32(32)
						return
					}
				}
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v9)+12)) = int32(-1)
				*(*int32)(unsafe.Add(mBase, uint32(v9)+8)) = l2
				*(*int32)(unsafe.Add(mBase, uint32(v9)+4)) = int32(0)
				*(*int32)(unsafe.Add(mBase, uint32(v9))) = l0
				v56 = *(*int32)(unsafe.Add(mBase, uint32(v45)+4))
				m.T0[v47].(func(*base.Module, int32, int32))(m, v9, v56)
				mBase = m.M
				v58 = m.ExcPending
				if v58 != 0 {
					return
				} else {
					v68 = int32(_a44)
					v70 = *(*int32)(unsafe.Add(mBase, _consts[582]))
					*(*int32)(unsafe.Add(mBase, _consts[582])) = v70 + int32(-1)
					m.G0 = v9 + int32(32)
					return
				}
			}
		}
	}
}
func F_moduleNotifyUserChanged(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v10 int64
	_ = v10
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+108))
	if v4 == int32(0) {
	} else {
		v7 = *(*int32)(unsafe.Add(mBase, uint32(v4)+8))
		if v7 == int32(0) {
		} else {
			v10 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
			v11 = *(*int32)(unsafe.Add(mBase, uint32(v4)+12))
			m.T0[v7].(func(*base.Module, int64, int32))(m, v10, v11)
			mBase = m.M
			v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+108))
			*(*int32)(unsafe.Add(mBase, uint32(v13)+16)) = int32(0)
			*(*int64)(unsafe.Add(mBase, uint32(v13)+8)) = int64(0)
		}
	}
	return
}
func F_modulePostExecutionUnitOperations(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	v3 = *(*int32)(unsafe.Add(mBase, _consts[95]))
	if v3 != 0 {
		return
	} else {
		v5 = *(*int32)(unsafe.Add(mBase, _consts[526]))
		if v5 == int32(0) {
			return
		} else {
			v8 = int32(0)
			v11 = *(*int32)(unsafe.Add(mBase, _consts[527]))
			v13 = v11 + int32(-1)
			*(*int32)(unsafe.Add(mBase, _consts[527])) = v13
			if v13 != 0 {
			} else {
				*(*int64)(unsafe.Add(mBase, _consts[528])) = int64(0)
			}
			v18 = int32(_a44)
			v19 = int32(0)
			*(*int32)(unsafe.Add(mBase, _consts[526])) = v19
			v22 = *(*int32)(unsafe.Add(mBase, _consts[67]))
			if v22 == v19 {
				F_unblockPostponedClients(m)
				mBase = m.M
				v28 = m.ExcPending
				if v28 != 0 {
					return
				} else {
					return
				}
			} else {
				F_unprotectClient(m, v22)
				mBase = m.M
				v26 = m.ExcPending
				if v26 != 0 {
					return
				} else {
					F_unblockPostponedClients(m)
					mBase = m.M
					v28 = m.ExcPending
					if v28 != 0 {
						return
					} else {
						return
					}
				}
			}
		}
	}
}
func F_moduleRemoveConfigs(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v27 int32
	_ = v27
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
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v60 int32
	_ = v60
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	v10 = v6 + int32(8)
	v11 = *(*int32)(unsafe.Add(mBase, uint32(v8)))
	*(*int32)(unsafe.Add(mBase, uint32(v10)+4)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v10))) = v11
	goto L1
L1:
	;
	v16 = v6 + int32(8)
	v18 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
	if v18 == int32(0) {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	m.G0 = v6 + int32(16)
	return
L3:
	;
	if v18 == int32(0) {
		goto L2
	} else {
		goto L6
	}
L4:
	;
	goto L3
L5:
	;
	v21 = *(*int32)(unsafe.Add(mBase, uint32(v16)+4))
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v18+base.B2i32(v21 == int32(0))<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v16))) = v27
	goto L4
L6:
	;
	v33 = v18
	goto L7
L7:
	;
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v33)+8))
	v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v36 = F_sdsnew(m, v35)
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		goto L9
	} else {
		goto L10
	}
L8:
	;
	goto L2
L9:
	;
	return
L10:
	;
	v39 = F_sdscat(m, v36, int32(_a955))
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
		goto L9
	} else {
		goto L11
	}
L11:
	;
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v34)))
	v42 = F_sdscat(m, v39, v41)
	mBase = m.M
	v43 = m.ExcPending
	if v43 != 0 {
		goto L9
	} else {
		goto L12
	}
L12:
	;
	F_removeConfig(m, v42)
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L9
	} else {
		goto L13
	}
L13:
	;
	F_sdsfree(m, v42)
	mBase = m.M
	v47 = m.ExcPending
	if v47 != 0 {
		goto L9
	} else {
		goto L14
	}
L14:
	;
	v49 = v6 + int32(8)
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v49)))
	if v51 == int32(0) {
		goto L16
	} else {
		goto L17
	}
L15:
	;
	if v51 != 0 {
		v33 = v51
		goto L7
	} else {
		goto L18
	}
L16:
	;
	goto L15
L17:
	;
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v49)+4))
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v51+base.B2i32(v54 == int32(0))<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v49))) = v60
	goto L16
L18:
	;
	goto L8
}
func F_moduleScanCallback(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v15 int32
	_ = v15
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
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
	var v42 int64
	_ = v42
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v77 int32
	_ = v77
	var v83 int32
	_ = v83
	var v94 int32
	_ = v94
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v104 int32
	_ = v104
	var v106 int32
	_ = v106
	v4 = int32(0)
	v7 = m.G0
	v9 = v7 - int32(96)
	m.G0 = v9
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v15&int32(2) == v4 {
		v35 = v4
	} else {
		v29 = l1 + (v15&int32(4) ^ int32(12)) + v15<<(uint(int32(3))%32)&int32(8)
		v30 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v29))))
		v35 = v29 + v30 + int32(1)
	}
	v36 = F_sdsdup(m, v35)
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		return
	} else {
		v38 = F_createObject(m, v4, v36)
		mBase = m.M
		v39 = m.ExcPending
		if v39 != 0 {
			return
		} else {
			v42 = int64(0)
			*(*int64)(unsafe.Add(mBase, uint32(v9+int32(88)))) = v42
			*(*int64)(unsafe.Add(mBase, uint32(v9+int32(80)))) = v42
			*(*int64)(unsafe.Add(mBase, uint32(v9+int32(72)))) = v42
			*(*int64)(unsafe.Add(mBase, uint32(v9+int32(64)))) = v42
			*(*int64)(unsafe.Add(mBase, uint32(v9+int32(56)))) = v42
			*(*int64)(unsafe.Add(mBase, uint32(v9+int32(48)))) = v42
			*(*int64)(unsafe.Add(mBase, uint32(v9+int32(40)))) = v42
			*(*int64)(unsafe.Add(mBase, uint32(v9)+32)) = v42
			v70 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			*(*int32)(unsafe.Add(mBase, uint32(v9)+8)) = v70
			v72 = *(*int32)(unsafe.Add(mBase, uint32(v70)+8))
			v73 = *(*int32)(unsafe.Add(mBase, uint32(v72)+96))
			*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = v38
			*(*int32)(unsafe.Add(mBase, uint32(v9)+12)) = v73
			F_incrRefCount(m, v38)
			mBase = m.M
			v77 = m.ExcPending
			if v77 != 0 {
				return
			} else {
				*(*int64)(unsafe.Add(mBase, uint32(v9)+24)) = int64(4294967296)
				*(*int32)(unsafe.Add(mBase, uint32(v9)+20)) = l1
				if l1 == int32(0) {
				} else {
					v83 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
					switch v83&int32(15) + int32(-3) {
					case 0:
						*(*int64)(unsafe.Add(mBase, uint32(v9)+88)) = int64(4294967296)
						*(*int32)(unsafe.Add(mBase, uint32(v9)+32)) = int32(0)
					default:
					case 3:
						*(*int32)(unsafe.Add(mBase, uint32(v9)+56)) = int32(0)
					}
				}
				v94 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
				v97 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
				v98 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				m.T0[v98].(func(*base.Module, int32, int32, int32, int32))(m, v94, v38, v9+int32(8), v97)
				mBase = m.M
				v100 = m.ExcPending
				if v100 != 0 {
					return
				} else {
					F_moduleCloseKey(m, v9+int32(8))
					mBase = m.M
					v104 = m.ExcPending
					if v104 != 0 {
						return
					} else {
						F_decrRefCount(m, v38)
						mBase = m.M
						v106 = m.ExcPending
						if v106 != 0 {
							return
						} else {
							m.G0 = v9 + int32(96)
							return
						}
					}
				}
			}
		}
	}
}
func F_moduleTryServeClientBlockedOnKey(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
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
	var v18 int64
	_ = v18
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v51 int64
	_ = v51
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v61 int64
	_ = v61
	var v65 int64
	_ = v65
	var v69 int32
	_ = v69
	var v73 int32
	_ = v73
	var v81 int64
	_ = v81
	var v83 int32
	_ = v83
	var v87 int64
	_ = v87
	var v91 int64
	_ = v91
	var v94 int32
	_ = v94
	var v102 int64
	_ = v102
	var v107 int32
	_ = v107
	var v109 int32
	_ = v109
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v120 int32
	_ = v120
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	v7 = m.G0
	v9 = v7 - int32(80)
	m.G0 = v9
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+116))
	v13 = *(*int32)(unsafe.Add(mBase, uint32(v12)+48))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(v13)+48))
	if v14 != 0 {
		v127 = int32(0)
		m.G0 = v9 + int32(80)
		return v127
	} else {
		v15 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
		v18 = int64(0)
		*(*int64)(unsafe.Add(mBase, uint32(v9+int32(32)))) = v18
		*(*int64)(unsafe.Add(mBase, uint32(v9+int32(72)))) = v18
		*(*int64)(unsafe.Add(mBase, uint32(v9+int32(64)))) = v18
		*(*int64)(unsafe.Add(mBase, uint32(v9+int32(56)))) = v18
		*(*int64)(unsafe.Add(mBase, uint32(v9+int32(48)))) = v18
		*(*int64)(unsafe.Add(mBase, uint32(v9+int32(40)))) = v18
		*(*int64)(unsafe.Add(mBase, uint32(v9)+24)) = v18
		*(*int32)(unsafe.Add(mBase, uint32(v9)+12)) = v15
		*(*int32)(unsafe.Add(mBase, uint32(v9)+8)) = int32(561)
		*(*int32)(unsafe.Add(mBase, uint32(v9)+36)) = int32(4)
		v48 = *(*int32)(unsafe.Add(mBase, _consts[131]))
		v49 = int32(0)
		v50 = *(*int32)(unsafe.Add(mBase, _consts[23]))
		v51 = m.T0[v50].(func(*base.Module) int64)(m)
		mBase = m.M
		if v48 == v49 {
			v61 = *(*int64)(unsafe.Add(mBase, _consts[529]))
			v65 = v61*int64(1000) + v51
		} else {
			v56 = *(*int32)(unsafe.Add(mBase, _consts[219]))
			v57 = base.I32_div_s(int32(1000000), v56)
			v65 = v51 + base.I64_extend_i32_s(v57)
		}
		*(*int64)(unsafe.Add(mBase, uint32(v9)+64)) = v65
		v69 = int32(0)
		v73 = *(*int32)(unsafe.Add(mBase, _consts[95]))
		*(*int32)(unsafe.Add(mBase, _consts[95])) = v73 + int32(1)
		if v73 != 0 {
		} else {
			v81 = F_ustime(m)
			mBase = m.M
			v83 = int32(0)
			*(*int64)(unsafe.Add(mBase, _consts[96])) = v81
			v87 = base.I64_div_s(v81, int64(1000))
			*(*int64)(unsafe.Add(mBase, _consts[35])) = v87
			v91 = base.I64_div_s(v81, int64(1000000))
			*(*int64)(unsafe.Add(mBase, _consts[47])) = v91
			v94 = *(*int32)(unsafe.Add(mBase, _consts[97]))
			F_lrulfu_updateClockAndPolicy(m, v87, int32(base.Ui32(v94&int32(2))>>(uint(int32(1))%32)))
			mBase = m.M
			v102 = *(*int64)(unsafe.Add(mBase, _consts[35]))
			*(*int64)(unsafe.Add(mBase, _consts[98])) = v102
		}
		*(*int32)(unsafe.Add(mBase, uint32(v9)+52)) = l1
		v107 = *(*int32)(unsafe.Add(mBase, uint32(v13)+28))
		*(*int32)(unsafe.Add(mBase, uint32(v9)+48)) = v107
		v109 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
		*(*int32)(unsafe.Add(mBase, uint32(v9)+20)) = v13
		*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = v109
		v114 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
		v115 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
		v116 = *(*int32)(unsafe.Add(mBase, uint32(v13)+8))
		v117 = m.T0[v116].(func(*base.Module, int32, int32, int32) int32)(m, v9+int32(8), v114, v115)
		mBase = m.M
		v120 = m.ExcPending
		if v120 != 0 {
			return int32(0)
		} else {
			F_moduleFreeContext(m, v9+int32(8))
			mBase = m.M
			v126 = m.ExcPending
			if v126 != 0 {
				return int32(0)
			} else {
				v127 = base.B2i32(v117 == int32(0))
				m.G0 = v9 + int32(80)
				return v127
			}
		}
	}
}
func F_moduleTypeDupOrReply(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v22 int32
	_ = v22
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
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
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	v10 = m.G0
	v12 = v10 - int32(16)
	m.G0 = v12
	v14 = F_objectGetVal(m, l4)
	mBase = m.M
	v15 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
	v16 = *(*int32)(unsafe.Add(mBase, uint32(v15)+72))
	v17 = *(*int32)(unsafe.Add(mBase, uint32(v15)+44))
	if v17 != 0 {
		if v16 == int32(0) {
			v35 = *(*int32)(unsafe.Add(mBase, uint32(v14)+4))
			v36 = m.T0[v17].(func(*base.Module, int32, int32, int32) int32)(m, l1, l2, v35)
			mBase = m.M
			v37 = m.ExcPending
			if v37 != 0 {
				return int32(0)
			} else {
				v39 = v36
				if v39 != 0 {
					v44 = F_createModuleObject(m, v15, v39)
					mBase = m.M
					v45 = m.ExcPending
					if v45 != 0 {
						return int32(0)
					} else {
						v47 = v44
						m.G0 = v12 + int32(16)
						return v47
					}
				} else {
					F_addReplyError(m, l0, int32(_a936))
					mBase = m.M
					v42 = m.ExcPending
					if v42 != 0 {
						return int32(0)
					} else {
						v47 = int32(0)
						m.G0 = v12 + int32(16)
						return v47
					}
				}
			}
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v12)+4)) = l2
			*(*int32)(unsafe.Add(mBase, uint32(v12))) = l1
			v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
			v29 = *(*int32)(unsafe.Add(mBase, uint32(v28)+28))
			*(*int32)(unsafe.Add(mBase, uint32(v12)+12)) = l3
			*(*int32)(unsafe.Add(mBase, uint32(v12)+8)) = v29
			v32 = *(*int32)(unsafe.Add(mBase, uint32(v14)+4))
			v33 = m.T0[v16].(func(*base.Module, int32, int32) int32)(m, v12, v32)
			mBase = m.M
			v34 = m.ExcPending
			if v34 != 0 {
				return int32(0)
			} else {
				v39 = v33
				if v39 != 0 {
					v44 = F_createModuleObject(m, v15, v39)
					mBase = m.M
					v45 = m.ExcPending
					if v45 != 0 {
						return int32(0)
					} else {
						v47 = v44
						m.G0 = v12 + int32(16)
						return v47
					}
				} else {
					F_addReplyError(m, l0, int32(_a936))
					mBase = m.M
					v42 = m.ExcPending
					if v42 != 0 {
						return int32(0)
					} else {
						v47 = int32(0)
						m.G0 = v12 + int32(16)
						return v47
					}
				}
			}
		}
	} else {
		if v16 != 0 {
			*(*int32)(unsafe.Add(mBase, uint32(v12)+4)) = l2
			*(*int32)(unsafe.Add(mBase, uint32(v12))) = l1
			v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
			v29 = *(*int32)(unsafe.Add(mBase, uint32(v28)+28))
			*(*int32)(unsafe.Add(mBase, uint32(v12)+12)) = l3
			*(*int32)(unsafe.Add(mBase, uint32(v12)+8)) = v29
			v32 = *(*int32)(unsafe.Add(mBase, uint32(v14)+4))
			v33 = m.T0[v16].(func(*base.Module, int32, int32) int32)(m, v12, v32)
			mBase = m.M
			v34 = m.ExcPending
			if v34 != 0 {
				return int32(0)
			} else {
				v39 = v33
				if v39 != 0 {
					v44 = F_createModuleObject(m, v15, v39)
					mBase = m.M
					v45 = m.ExcPending
					if v45 != 0 {
						return int32(0)
					} else {
						v47 = v44
						m.G0 = v12 + int32(16)
						return v47
					}
				} else {
					F_addReplyError(m, l0, int32(_a936))
					mBase = m.M
					v42 = m.ExcPending
					if v42 != 0 {
						return int32(0)
					} else {
						v47 = int32(0)
						m.G0 = v12 + int32(16)
						return v47
					}
				}
			}
		} else {
			F_addReplyError(m, l0, int32(_a937))
			mBase = m.M
			v22 = m.ExcPending
			if v22 != 0 {
				return int32(0)
			} else {
				v47 = int32(0)
				m.G0 = v12 + int32(16)
				return v47
			}
		}
	}
}
func F_moduleTypeLookupModuleByID(m *base.Module, l0 int64) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v17 int64
	_ = v17
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v24 int64
	_ = v24
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v32 int64
	_ = v32
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v63 int32
	_ = v63
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
	var v76 int64
	_ = v76
	var v77 int64
	_ = v77
	var v78 int64
	_ = v78
	var v79 int64
	_ = v79
	var v80 int64
	_ = v80
	var v81 int64
	_ = v81
	var v82 int64
	_ = v82
	var v84 int64
	_ = v84
	var v86 int64
	_ = v86
	var v88 int64
	_ = v88
	var v90 int64
	_ = v90
	var v92 int64
	_ = v92
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v99 int32
	_ = v99
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v106 int32
	_ = v106
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v114 int32
	_ = v114
	var v118 int32
	_ = v118
	var v120 int32
	_ = v120
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v134 int32
	_ = v134
	var v139 int32
	_ = v139
	var v143 int32
	_ = v143
	var v147 int32
	_ = v147
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v176 int32
	_ = v176
	var v178 int32
	_ = v178
	var v181 int32
	_ = v181
	var v187 int32
	_ = v187
	var v190 int32
	_ = v190
	var v191 int64
	_ = v191
	var v195 int32
	_ = v195
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v207 int32
	_ = v207
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v215 int32
	_ = v215
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v224 int32
	_ = v224
	var v226 int32
	_ = v226
	var v228 int64
	_ = v228
	var v229 int64
	_ = v229
	var v230 int64
	_ = v230
	var v231 int64
	_ = v231
	var v232 int64
	_ = v232
	var v233 int64
	_ = v233
	var v234 int64
	_ = v234
	var v236 int64
	_ = v236
	var v238 int64
	_ = v238
	var v240 int64
	_ = v240
	var v242 int64
	_ = v242
	var v244 int64
	_ = v244
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v248 int32
	_ = v248
	var v251 int32
	_ = v251
	var v254 int32
	_ = v254
	var v255 int32
	_ = v255
	var v256 int32
	_ = v256
	var v258 int32
	_ = v258
	var v261 int32
	_ = v261
	var v262 int32
	_ = v262
	var v266 int32
	_ = v266
	var v270 int32
	_ = v270
	var v272 int32
	_ = v272
	var v279 int32
	_ = v279
	var v280 int32
	_ = v280
	var v281 int32
	_ = v281
	var v286 int32
	_ = v286
	var v291 int32
	_ = v291
	var v295 int32
	_ = v295
	var v299 int32
	_ = v299
	var v304 int32
	_ = v304
	var v311 int32
	_ = v311
	var v319 int32
	_ = v319
	var v323 int32
	_ = v323
	v8 = m.G0
	v10 = v8 - int32(16)
	m.G0 = v10
	v13 = *(*int32)(unsafe.Add(mBase, _consts[550]))
	if v13 != 0 {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	m.G0 = v10 + int32(16)
	return v323
L2:
	;
	v39 = *(*int32)(unsafe.Add(mBase, _consts[541]))
	v40 = F_dictGetIterator(m, v39)
	mBase = m.M
	v43 = m.ExcPending
	if v43 != 0 {
		goto L13
	} else {
		goto L14
	}
L3:
	;
	v17 = *(*int64)(unsafe.Add(mBase, _consts[551]))
	if v17 == l0 {
		v323 = v13
		goto L1
	} else {
		goto L5
	}
L4:
	;
	v36 = int32(1)
	v37 = int32(0)
	goto L2
L5:
	;
	v20 = *(*int32)(unsafe.Add(mBase, _consts[552]))
	if v20 != 0 {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	v24 = *(*int64)(unsafe.Add(mBase, _consts[553]))
	if v24 == l0 {
		v323 = v20
		goto L1
	} else {
		goto L8
	}
L7:
	;
	v21 = int32(1)
	v36 = v21
	v37 = v21
	goto L2
L8:
	;
	v27 = *(*int32)(unsafe.Add(mBase, _consts[554]))
	if v27 != 0 {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	v30 = int32(0)
	v32 = *(*int64)(unsafe.Add(mBase, _consts[555]))
	if v32 == l0 {
		v323 = v27
		goto L1
	} else {
		goto L11
	}
L10:
	;
	v36 = int32(1)
	v37 = int32(2)
	goto L2
L11:
	;
	v36 = v30
	v37 = int32(3)
	goto L2
L12:
	;
	F_dictReleaseIterator(m, v40)
	mBase = m.M
	v319 = m.ExcPending
	if v319 != 0 {
		goto L13
	} else {
		goto L86
	}
L13:
	;
	return int32(0)
L14:
	;
	v51 = v40 + int32(20)
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v40)+16))
	if v52 != 0 {
		goto L17
	} else {
		goto L18
	}
L15:
	;
	if v147 == int32(0) {
		goto L12
	} else {
		goto L41
	}
L16:
	;
	v58 = v51
	v59 = v55
	goto L19
L17:
	;
	v55 = int32(1)
	goto L16
L18:
	;
	v55 = int32(0)
	goto L16
L19:
	;
	switch v59 {
	case 0:
		goto L24
	default:
		goto L23
	}
L21:
	;
	v59 = int32(0)
	goto L19
L22:
	;
	goto L15
L23:
	;
	v139 = *(*int32)(unsafe.Add(mBase, uint32(v58)))
	*(*int32)(unsafe.Add(mBase, uint32(v40)+16)) = v139
	if v139 == int32(0) {
		goto L21
	} else {
		goto L40
	}
L24:
	;
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v40)+4))
	if v63 != int32(-1) {
		v102 = v63
		goto L25
	} else {
		goto L26
	}
L25:
	;
	v103 = int32(1)
	v104 = v102 + v103
	*(*int32)(unsafe.Add(mBase, uint32(v40)+4)) = v104
	v106 = int32(0)
	v109 = *(*int32)(unsafe.Add(mBase, uint32(v40)))
	v110 = *(*int32)(unsafe.Add(mBase, uint32(v40)+8))
	v114 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v109+v110+int32(26)))))
	if v114 == int32(255) {
		goto L34
	} else {
		goto L35
	}
L26:
	;
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v40)+8))
	if v67 != 0 {
		v102 = int32(-1)
		goto L25
	} else {
		goto L27
	}
L27:
	;
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v40)))
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v40)+12))
	if v69 == int32(0) {
		goto L29
	} else {
		goto L30
	}
L28:
	;
	v96 = *(*int32)(unsafe.Add(mBase, uint32(v95)+20))
	if v96 != int32(-1) {
		goto L31
	} else {
		goto L32
	}
L29:
	;
	v76 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v68)+16)))
	v77 = int64(*(*int8)(unsafe.Add(mBase, uint32(v68)+27)))
	v78 = int64(*(*int32)(unsafe.Add(mBase, uint32(v68)+8)))
	v79 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v68)+12)))
	v80 = int64(*(*int8)(unsafe.Add(mBase, uint32(v68)+26)))
	v81 = int64(*(*int32)(unsafe.Add(mBase, uint32(v68)+4)))
	v82 = F_wangHash64(m, v81)
	mBase = m.M
	v84 = F_wangHash64(m, v80+v82)
	mBase = m.M
	v86 = F_wangHash64(m, v79+v84)
	mBase = m.M
	v88 = F_wangHash64(m, v78+v86)
	mBase = m.M
	v90 = F_wangHash64(m, v77+v88)
	mBase = m.M
	v92 = F_wangHash64(m, v76+v90)
	mBase = m.M
	*(*int64)(unsafe.Add(mBase, uint32(v40)+24)) = v92
	v94 = *(*int32)(unsafe.Add(mBase, uint32(v40)))
	v95 = v94
	goto L28
L30:
	;
	v72 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v68)+24)))
	v74 = v72 + int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v68)+24)) = uint16(v74)
	v95 = v68
	goto L28
L31:
	;
	v102 = v96 + int32(-1)
	goto L25
L32:
	;
	v99 = *(*int32)(unsafe.Add(mBase, uint32(v40)+4))
	v102 = v99
	goto L25
L33:
	;
	v129 = int32(2)
	v134 = *(*int32)(unsafe.Add(mBase, uint32(v109+v127<<(uint(v129)%32)+int32(4))))
	v58 = v134 + v128<<(uint(v129)%32)
	v59 = int32(1)
	goto L19
L34:
	;
	v118 = v106
	goto L36
L35:
	;
	v118 = v103 << (uint(v114) % 32)
	goto L36
L36:
	;
	if v104 < v118 {
		v127 = v110
		v128 = v104
		goto L33
	} else {
		goto L37
	}
L37:
	;
	if v110 != 0 {
		v147 = v106
		goto L22
	} else {
		goto L38
	}
L38:
	;
	v120 = *(*int32)(unsafe.Add(mBase, uint32(v109)+20))
	if v120 == int32(-1) {
		v147 = v106
		goto L22
	} else {
		goto L39
	}
L39:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v40)+4)) = int64(4294967296)
	v127 = int32(1)
	v128 = int32(0)
	goto L33
L40:
	;
	v143 = *(*int32)(unsafe.Add(mBase, uint32(v139)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v51))) = v143
	v147 = v139
	goto L22
L41:
	;
	v159 = v147
	goto L42
L42:
	;
	v160 = *(*int32)(unsafe.Add(mBase, uint32(v159)+8))
	goto L44
L43:
	;
	F_dictReleaseIterator(m, v40)
	mBase = m.M
	v304 = m.ExcPending
	if v304 != 0 {
		goto L13
	} else {
		goto L84
	}
L44:
	;
	v161 = *(*int32)(unsafe.Add(mBase, uint32(v160)+16))
	v163 = v10 + int32(8)
	v164 = *(*int32)(unsafe.Add(mBase, uint32(v161)))
	*(*int32)(unsafe.Add(mBase, uint32(v163)+4)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v163))) = v164
	goto L45
L45:
	;
	goto L47
L46:
	;
	v203 = v40 + int32(20)
	v204 = *(*int32)(unsafe.Add(mBase, uint32(v40)+16))
	if v204 != 0 {
		goto L57
	} else {
		goto L58
	}
L47:
	;
	v176 = v10 + int32(8)
	v178 = *(*int32)(unsafe.Add(mBase, uint32(v176)))
	if v178 == int32(0) {
		goto L51
	} else {
		goto L52
	}
L48:
	;
	v195 = v190
	goto L46
L49:
	;
	v190 = *(*int32)(unsafe.Add(mBase, uint32(v178)+8))
	v191 = *(*int64)(unsafe.Add(mBase, uint32(v190)))
	if base.Ui64(int64(1023)) < base.Ui64(v191^l0) {
		goto L47
	} else {
		goto L54
	}
L50:
	;
	if v178 != 0 {
		goto L49
	} else {
		goto L53
	}
L51:
	;
	goto L50
L52:
	;
	v181 = *(*int32)(unsafe.Add(mBase, uint32(v176)+4))
	v187 = *(*int32)(unsafe.Add(mBase, uint32(v178+base.B2i32(v181 == int32(0))<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v176))) = v187
	goto L51
L53:
	;
	v195 = int32(0)
	goto L46
L54:
	;
	goto L48
L55:
	;
	if v195 != 0 {
		goto L81
	} else {
		goto L82
	}
L56:
	;
	v210 = v203
	v211 = v207
	goto L59
L57:
	;
	v207 = int32(1)
	goto L56
L58:
	;
	v207 = int32(0)
	goto L56
L59:
	;
	switch v211 {
	case 0:
		goto L64
	default:
		goto L63
	}
L61:
	;
	v211 = int32(0)
	goto L59
L62:
	;
	goto L55
L63:
	;
	v291 = *(*int32)(unsafe.Add(mBase, uint32(v210)))
	*(*int32)(unsafe.Add(mBase, uint32(v40)+16)) = v291
	if v291 == int32(0) {
		goto L61
	} else {
		goto L80
	}
L64:
	;
	v215 = *(*int32)(unsafe.Add(mBase, uint32(v40)+4))
	if v215 != int32(-1) {
		v254 = v215
		goto L65
	} else {
		goto L66
	}
L65:
	;
	v255 = int32(1)
	v256 = v254 + v255
	*(*int32)(unsafe.Add(mBase, uint32(v40)+4)) = v256
	v258 = int32(0)
	v261 = *(*int32)(unsafe.Add(mBase, uint32(v40)))
	v262 = *(*int32)(unsafe.Add(mBase, uint32(v40)+8))
	v266 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v261+v262+int32(26)))))
	if v266 == int32(255) {
		goto L74
	} else {
		goto L75
	}
L66:
	;
	v219 = *(*int32)(unsafe.Add(mBase, uint32(v40)+8))
	if v219 != 0 {
		v254 = int32(-1)
		goto L65
	} else {
		goto L67
	}
L67:
	;
	v220 = *(*int32)(unsafe.Add(mBase, uint32(v40)))
	v221 = *(*int32)(unsafe.Add(mBase, uint32(v40)+12))
	if v221 == int32(0) {
		goto L69
	} else {
		goto L70
	}
L68:
	;
	v248 = *(*int32)(unsafe.Add(mBase, uint32(v247)+20))
	if v248 != int32(-1) {
		goto L71
	} else {
		goto L72
	}
L69:
	;
	v228 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v220)+16)))
	v229 = int64(*(*int8)(unsafe.Add(mBase, uint32(v220)+27)))
	v230 = int64(*(*int32)(unsafe.Add(mBase, uint32(v220)+8)))
	v231 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v220)+12)))
	v232 = int64(*(*int8)(unsafe.Add(mBase, uint32(v220)+26)))
	v233 = int64(*(*int32)(unsafe.Add(mBase, uint32(v220)+4)))
	v234 = F_wangHash64(m, v233)
	mBase = m.M
	v236 = F_wangHash64(m, v232+v234)
	mBase = m.M
	v238 = F_wangHash64(m, v231+v236)
	mBase = m.M
	v240 = F_wangHash64(m, v230+v238)
	mBase = m.M
	v242 = F_wangHash64(m, v229+v240)
	mBase = m.M
	v244 = F_wangHash64(m, v228+v242)
	mBase = m.M
	*(*int64)(unsafe.Add(mBase, uint32(v40)+24)) = v244
	v246 = *(*int32)(unsafe.Add(mBase, uint32(v40)))
	v247 = v246
	goto L68
L70:
	;
	v224 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v220)+24)))
	v226 = v224 + int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v220)+24)) = uint16(v226)
	v247 = v220
	goto L68
L71:
	;
	v254 = v248 + int32(-1)
	goto L65
L72:
	;
	v251 = *(*int32)(unsafe.Add(mBase, uint32(v40)+4))
	v254 = v251
	goto L65
L73:
	;
	v281 = int32(2)
	v286 = *(*int32)(unsafe.Add(mBase, uint32(v261+v279<<(uint(v281)%32)+int32(4))))
	v210 = v286 + v280<<(uint(v281)%32)
	v211 = int32(1)
	goto L59
L74:
	;
	v270 = v258
	goto L76
L75:
	;
	v270 = v255 << (uint(v266) % 32)
	goto L76
L76:
	;
	if v256 < v270 {
		v279 = v262
		v280 = v256
		goto L73
	} else {
		goto L77
	}
L77:
	;
	if v262 != 0 {
		v299 = v258
		goto L62
	} else {
		goto L78
	}
L78:
	;
	v272 = *(*int32)(unsafe.Add(mBase, uint32(v261)+20))
	if v272 == int32(-1) {
		v299 = v258
		goto L62
	} else {
		goto L79
	}
L79:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v40)+4)) = int64(4294967296)
	v279 = int32(1)
	v280 = int32(0)
	goto L73
L80:
	;
	v295 = *(*int32)(unsafe.Add(mBase, uint32(v291)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v203))) = v295
	v299 = v291
	goto L62
L81:
	;
	goto L43
L82:
	;
	if v299 != 0 {
		v159 = v299
		goto L42
	} else {
		goto L83
	}
L83:
	;
	goto L81
L84:
	;
	if v36^int32(1)|base.B2i32(v195 == int32(0)) != 0 {
		v323 = v195
		goto L1
	} else {
		goto L85
	}
L85:
	;
	v311 = v37 << (uint(int32(4)) % 32)
	*(*int32)(unsafe.Add(mBase, uint32(v311)+uint32(_consts[550]))) = v195
	*(*int64)(unsafe.Add(mBase, uint32(v311)+uint32(_consts[551]))) = l0
	v323 = v195
	goto L1
L86:
	;
	v323 = int32(0)
	goto L1
}
func F_moduleTypeNameByID(m *base.Module, l0 int32, l1 int64) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v22 int32
	_ = v22
	var v29 int32
	_ = v29
	var v37 int32
	_ = v37
	var v45 int32
	_ = v45
	var v53 int32
	_ = v53
	var v61 int32
	_ = v61
	var v69 int32
	_ = v69
	var v75 int32
	_ = v75
	v3 = int32(0)
	v6 = *(*int32)(unsafe.Add(mBase, _consts[556]))
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+9)) = uint8(v3)
	v9 = base.I32_wrap_i64(l1)
	v12 = int32(63)
	v15 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6+int32(base.Ui32(v9)>>(uint(int32(10))%32))&v12))))
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+8)) = uint8(v15)
	v22 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6+int32(base.Ui32(v9)>>(uint(int32(16))%32))&v12))))
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+7)) = uint8(v22)
	v29 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6+int32(base.Ui32(v9)>>(uint(int32(22))%32))&v12))))
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+6)) = uint8(v29)
	v37 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6+base.I32_wrap_i64(int64(base.Ui64(l1)>>(uint(int64(28))%64)))&v12))))
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+5)) = uint8(v37)
	v45 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6+base.I32_wrap_i64(int64(base.Ui64(l1)>>(uint(int64(34))%64)))&v12))))
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+4)) = uint8(v45)
	v53 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6+base.I32_wrap_i64(int64(base.Ui64(l1)>>(uint(int64(40))%64)))&v12))))
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+3)) = uint8(v53)
	v61 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6+base.I32_wrap_i64(int64(base.Ui64(l1)>>(uint(int64(46))%64)))&v12))))
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+2)) = uint8(v61)
	v69 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6+base.I32_wrap_i64(int64(base.Ui64(l1)>>(uint(int64(52))%64)))&v12))))
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+1)) = uint8(v69)
	v75 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6+base.I32_wrap_i64(int64(base.Ui64(l1)>>(uint(int64(58))%64)))))))
	*(*uint8)(unsafe.Add(mBase, uint32(l0))) = uint8(v75)
	return
}
func F_moduleUnregisterCleanup(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
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
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v50 int64
	_ = v50
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v60 int32
	_ = v60
	var v64 int32
	_ = v64
	var v81 int32
	_ = v81
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v96 int32
	_ = v96
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v114 int32
	_ = v114
	var v116 int32
	_ = v116
	var v119 int32
	_ = v119
	var v125 int32
	_ = v125
	var v131 int32
	_ = v131
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v138 int32
	_ = v138
	var v140 int32
	_ = v140
	var v142 int32
	_ = v142
	var v144 int32
	_ = v144
	var v146 int32
	_ = v146
	var v149 int32
	_ = v149
	var v155 int32
	_ = v155
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v172 int32
	_ = v172
	var v174 int32
	_ = v174
	var v177 int32
	_ = v177
	var v183 int32
	_ = v183
	var v189 int32
	_ = v189
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v198 int32
	_ = v198
	var v200 int32
	_ = v200
	var v202 int32
	_ = v202
	var v204 int32
	_ = v204
	var v207 int32
	_ = v207
	var v213 int32
	_ = v213
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
	var v228 int32
	_ = v228
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v236 int32
	_ = v236
	var v238 int32
	_ = v238
	var v241 int32
	_ = v241
	var v247 int32
	_ = v247
	var v253 int32
	_ = v253
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v260 int32
	_ = v260
	var v262 int32
	_ = v262
	var v264 int32
	_ = v264
	var v266 int32
	_ = v266
	var v268 int32
	_ = v268
	var v271 int32
	_ = v271
	var v277 int32
	_ = v277
	var v285 int32
	_ = v285
	var v292 int32
	_ = v292
	var v294 int32
	_ = v294
	var v297 int32
	_ = v297
	var v300 int32
	_ = v300
	var v302 int32
	_ = v302
	var v305 int32
	_ = v305
	var v308 int32
	_ = v308
	var v311 int32
	_ = v311
	var v314 int32
	_ = v314
	var v316 int32
	_ = v316
	var v319 int32
	_ = v319
	var v326 int32
	_ = v326
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v11 = *(*int32)(unsafe.Add(mBase, _consts[3]))
	v13 = v8 + int32(8)
	v14 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
	*(*int32)(unsafe.Add(mBase, uint32(v13)+4)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v13))) = v14
	goto L1
L1:
	;
	v19 = v8 + int32(8)
	v21 = *(*int32)(unsafe.Add(mBase, uint32(v19)))
	if v21 == int32(0) {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	F_moduleUnregisterCommands(m, l0)
	mBase = m.M
	v104 = m.ExcPending
	if v104 != 0 {
		goto L19
	} else {
		goto L25
	}
L3:
	;
	if v21 == int32(0) {
		goto L2
	} else {
		goto L6
	}
L4:
	;
	goto L3
L5:
	;
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v19)+4))
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v21+base.B2i32(v24 == int32(0))<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v19))) = v30
	goto L4
L6:
	;
	v36 = v21
	goto L7
L7:
	;
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v36)+8))
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v39)+108))
	if v40 == int32(0) {
		goto L9
	} else {
		goto L10
	}
L8:
	;
	goto L2
L9:
	;
	v85 = v8 + int32(8)
	v87 = *(*int32)(unsafe.Add(mBase, uint32(v85)))
	if v87 == int32(0) {
		goto L22
	} else {
		goto L23
	}
L10:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v40)+16))
	if v43 == int32(0) {
		goto L9
	} else {
		goto L11
	}
L11:
	;
	if v43 != l0 {
		goto L9
	} else {
		goto L12
	}
L12:
	;
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v40)+8))
	if v47 == int32(0) {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	v60 = *(*int32)(unsafe.Add(mBase, _consts[4]))
	*(*int32)(unsafe.Add(mBase, uint32(v39)+328)) = v60
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v39)+204))
	goto L17
L14:
	;
	v50 = *(*int64)(unsafe.Add(mBase, uint32(v39)))
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v40)+12))
	m.T0[v47].(func(*base.Module, int64, int32))(m, v50, v51)
	mBase = m.M
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v39)+108))
	*(*int32)(unsafe.Add(mBase, uint32(v53)+16)) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v53)+8)) = int64(0)
	goto L13
L15:
	;
	F_freeClientOrCloseLater(m, v39, int32(1))
	mBase = m.M
	v81 = m.ExcPending
	if v81 != 0 {
		goto L19
	} else {
		goto L20
	}
L17:
	;
	goto L18
L18:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v39)+204)) = v64 & int32(-8388609)
	goto L15
L19:
	;
	return
L20:
	;
	goto L9
L21:
	;
	if v87 != 0 {
		v36 = v87
		goto L7
	} else {
		goto L24
	}
L22:
	;
	goto L21
L23:
	;
	v90 = *(*int32)(unsafe.Add(mBase, uint32(v85)+4))
	v96 = *(*int32)(unsafe.Add(mBase, uint32(v87+base.B2i32(v90 == int32(0))<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v85))) = v96
	goto L22
L24:
	;
	goto L8
L25:
	;
	v105 = int32(0)
	v106 = *(*int32)(unsafe.Add(mBase, _consts[583]))
	v108 = v8 + int32(8)
	v109 = *(*int32)(unsafe.Add(mBase, uint32(v106)))
	*(*int32)(unsafe.Add(mBase, uint32(v108)+4)) = v105
	*(*int32)(unsafe.Add(mBase, uint32(v108))) = v109
	goto L26
L26:
	;
	v114 = v8 + int32(8)
	v116 = *(*int32)(unsafe.Add(mBase, uint32(v114)))
	if v116 == int32(0) {
		goto L29
	} else {
		goto L30
	}
L27:
	;
	v162 = F_moduleUnregisterSharedAPI(m, l0)
	mBase = m.M
	v163 = m.ExcPending
	if v163 != 0 {
		goto L19
	} else {
		goto L42
	}
L28:
	;
	if v116 == int32(0) {
		goto L27
	} else {
		goto L31
	}
L29:
	;
	goto L28
L30:
	;
	v119 = *(*int32)(unsafe.Add(mBase, uint32(v114)+4))
	v125 = *(*int32)(unsafe.Add(mBase, uint32(v116+base.B2i32(v119 == int32(0))<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v114))) = v125
	goto L29
L31:
	;
	v131 = v116
	goto L32
L32:
	;
	v134 = *(*int32)(unsafe.Add(mBase, uint32(v131)+8))
	v135 = *(*int32)(unsafe.Add(mBase, uint32(v134)))
	if v135 != l0 {
		goto L34
	} else {
		goto L35
	}
L33:
	;
	goto L27
L34:
	;
	v144 = v8 + int32(8)
	v146 = *(*int32)(unsafe.Add(mBase, uint32(v144)))
	if v146 == int32(0) {
		goto L39
	} else {
		goto L40
	}
L35:
	;
	v138 = *(*int32)(unsafe.Add(mBase, _consts[583]))
	F_listDelNode(m, v138, v131)
	mBase = m.M
	v140 = m.ExcPending
	if v140 != 0 {
		goto L19
	} else {
		goto L36
	}
L36:
	;
	F_valkey_free(m, v134)
	mBase = m.M
	v142 = m.ExcPending
	if v142 != 0 {
		goto L19
	} else {
		goto L37
	}
L37:
	;
	goto L34
L38:
	;
	if v146 != 0 {
		v131 = v146
		goto L32
	} else {
		goto L41
	}
L39:
	;
	goto L38
L40:
	;
	v149 = *(*int32)(unsafe.Add(mBase, uint32(v144)+4))
	v155 = *(*int32)(unsafe.Add(mBase, uint32(v146+base.B2i32(v149 == int32(0))<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v144))) = v155
	goto L39
L41:
	;
	goto L33
L42:
	;
	v164 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v166 = v8 + int32(8)
	v167 = *(*int32)(unsafe.Add(mBase, uint32(v164)))
	*(*int32)(unsafe.Add(mBase, uint32(v166)+4)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v166))) = v167
	goto L43
L43:
	;
	v172 = v8 + int32(8)
	v174 = *(*int32)(unsafe.Add(mBase, uint32(v172)))
	if v174 == int32(0) {
		goto L46
	} else {
		goto L47
	}
L44:
	;
	v220 = F_moduleUnregisterFilters(m, l0)
	mBase = m.M
	v221 = m.ExcPending
	if v221 != 0 {
		goto L19
	} else {
		goto L59
	}
L45:
	;
	if v174 == int32(0) {
		goto L44
	} else {
		goto L48
	}
L46:
	;
	goto L45
L47:
	;
	v177 = *(*int32)(unsafe.Add(mBase, uint32(v172)+4))
	v183 = *(*int32)(unsafe.Add(mBase, uint32(v174+base.B2i32(v177 == int32(0))<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v172))) = v183
	goto L46
L48:
	;
	v189 = v174
	goto L49
L49:
	;
	v192 = *(*int32)(unsafe.Add(mBase, uint32(v189)+8))
	v193 = *(*int32)(unsafe.Add(mBase, uint32(v192)+20))
	v194 = F_listSearchKey(m, v193, l0)
	mBase = m.M
	v195 = m.ExcPending
	if v195 != 0 {
		goto L19
	} else {
		goto L52
	}
L50:
	;
	goto L44
L51:
	;
	v202 = v8 + int32(8)
	v204 = *(*int32)(unsafe.Add(mBase, uint32(v202)))
	if v204 == int32(0) {
		goto L56
	} else {
		goto L57
	}
L52:
	;
	if v194 == int32(0) {
		goto L51
	} else {
		goto L53
	}
L53:
	;
	v198 = *(*int32)(unsafe.Add(mBase, uint32(v192)+20))
	F_listDelNode(m, v198, v194)
	mBase = m.M
	v200 = m.ExcPending
	if v200 != 0 {
		goto L19
	} else {
		goto L54
	}
L54:
	;
	goto L51
L55:
	;
	if v204 != 0 {
		v189 = v204
		goto L49
	} else {
		goto L58
	}
L56:
	;
	goto L55
L57:
	;
	v207 = *(*int32)(unsafe.Add(mBase, uint32(v202)+4))
	v213 = *(*int32)(unsafe.Add(mBase, uint32(v204+base.B2i32(v207 == int32(0))<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v202))) = v213
	goto L56
L58:
	;
	goto L50
L59:
	;
	F_moduleUnsubscribeAllServerEvents(m, l0)
	mBase = m.M
	v223 = m.ExcPending
	if v223 != 0 {
		goto L19
	} else {
		goto L60
	}
L60:
	;
	F_moduleRemoveConfigs(m, l0)
	mBase = m.M
	v225 = m.ExcPending
	if v225 != 0 {
		goto L19
	} else {
		goto L61
	}
L61:
	;
	v226 = int32(0)
	v228 = *(*int32)(unsafe.Add(mBase, _consts[559]))
	v230 = v8 + int32(8)
	v231 = *(*int32)(unsafe.Add(mBase, uint32(v228)))
	*(*int32)(unsafe.Add(mBase, uint32(v230)+4)) = v226
	*(*int32)(unsafe.Add(mBase, uint32(v230))) = v231
	goto L62
L62:
	;
	v236 = v8 + int32(8)
	v238 = *(*int32)(unsafe.Add(mBase, uint32(v236)))
	if v238 == int32(0) {
		goto L65
	} else {
		goto L66
	}
L63:
	;
	v285 = *(*int32)(unsafe.Add(mBase, _consts[139]))
	if v285 == int32(0) {
		goto L78
	} else {
		goto L79
	}
L64:
	;
	if v238 == int32(0) {
		goto L63
	} else {
		goto L67
	}
L65:
	;
	goto L64
L66:
	;
	v241 = *(*int32)(unsafe.Add(mBase, uint32(v236)+4))
	v247 = *(*int32)(unsafe.Add(mBase, uint32(v238+base.B2i32(v241 == int32(0))<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v236))) = v247
	goto L65
L67:
	;
	v253 = v238
	goto L68
L68:
	;
	v256 = *(*int32)(unsafe.Add(mBase, uint32(v253)+8))
	v257 = *(*int32)(unsafe.Add(mBase, uint32(v256)))
	if v257 != l0 {
		goto L70
	} else {
		goto L71
	}
L69:
	;
	goto L63
L70:
	;
	v266 = v8 + int32(8)
	v268 = *(*int32)(unsafe.Add(mBase, uint32(v266)))
	if v268 == int32(0) {
		goto L75
	} else {
		goto L76
	}
L71:
	;
	v260 = *(*int32)(unsafe.Add(mBase, _consts[559]))
	F_listDelNode(m, v260, v253)
	mBase = m.M
	v262 = m.ExcPending
	if v262 != 0 {
		goto L19
	} else {
		goto L72
	}
L72:
	;
	F_valkey_free(m, v256)
	mBase = m.M
	v264 = m.ExcPending
	if v264 != 0 {
		goto L19
	} else {
		goto L73
	}
L73:
	;
	goto L70
L74:
	;
	if v268 != 0 {
		v253 = v268
		goto L68
	} else {
		goto L77
	}
L75:
	;
	goto L74
L76:
	;
	v271 = *(*int32)(unsafe.Add(mBase, uint32(v266)+4))
	v277 = *(*int32)(unsafe.Add(mBase, uint32(v268+base.B2i32(v271 == int32(0))<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v266))) = v277
	goto L75
L77:
	;
	goto L69
L78:
	;
	m.G0 = v8 + int32(16)
	return
L79:
	;
	v292 = v226
	goto L80
L80:
	;
	v294 = v292 << (uint(int32(2)) % 32)
	v297 = *(*int32)(unsafe.Add(mBase, uint32(v294)+uint32(_consts[565])))
	if v297 == int32(0) {
		goto L82
	} else {
		goto L83
	}
L81:
	;
	goto L78
L82:
	;
	v326 = v292 + int32(1)
	if v326 != int32(256) {
		v292 = v326
		goto L80
	} else {
		goto L92
	}
L83:
	;
	v300 = *(*int32)(unsafe.Add(mBase, uint32(v297)+12))
	if v300 != l0 {
		goto L84
	} else {
		goto L85
	}
L84:
	;
	v308 = v297
	goto L87
L85:
	;
	v302 = *(*int32)(unsafe.Add(mBase, uint32(v297)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v294)+uint32(_consts[565]))) = v302
	F_valkey_free(m, v297)
	mBase = m.M
	v305 = m.ExcPending
	if v305 != 0 {
		goto L19
	} else {
		goto L86
	}
L86:
	;
	goto L82
L87:
	;
	v311 = *(*int32)(unsafe.Add(mBase, uint32(v308)+16))
	if v311 == int32(0) {
		goto L82
	} else {
		goto L89
	}
L88:
	;
	v316 = *(*int32)(unsafe.Add(mBase, uint32(v311)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v308)+16)) = v316
	F_valkey_free(m, v311)
	mBase = m.M
	v319 = m.ExcPending
	if v319 != 0 {
		goto L19
	} else {
		goto L91
	}
L89:
	;
	v314 = *(*int32)(unsafe.Add(mBase, uint32(v311)+12))
	if v314 != l0 {
		v308 = v311
		goto L87
	} else {
		goto L90
	}
L90:
	;
	goto L88
L91:
	;
	goto L82
L92:
	;
	goto L81
}
func F_moduleUnregisterCommands(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v29 int32
	_ = v29
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
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
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v90 int32
	_ = v90
	var v96 int32
	_ = v96
	v4 = m.G0
	v6 = v4 - int32(64)
	m.G0 = v6
	F_drainIOThreadsQueue(m)
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v11 = v6 + int32(16)
	v13 = *(*int32)(unsafe.Add(mBase, _consts[318]))
	v14 = int32(1)
	v15 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v11)+14)) = uint8(v15)
	*(*int32)(unsafe.Add(mBase, uint32(v11))) = v13
	*(*int32)(unsafe.Add(mBase, uint32(v11)+24)) = v15
	*(*uint8)(unsafe.Add(mBase, uint32(v11)+15)) = uint8(v14)
	*(*int32)(unsafe.Add(mBase, uint32(v11)+8)) = int32(-1)
	if v13 == v15 {
		goto L4
	} else {
		goto L5
	}
L3:
	;
	v36 = F_hashtableNext(m, v6+int32(16), v6+int32(12))
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		goto L1
	} else {
		goto L10
	}
L4:
	;
	goto L3
L5:
	;
	goto L6
L6:
	;
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v13)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+24)) = v29
	*(*int32)(unsafe.Add(mBase, uint32(v13)+40)) = v11
	goto L4
L7:
	;
	F__serverAssert(m, int32(_a957), int32(_a917), int32(13302))
	mBase = m.M
	v96 = m.ExcPending
	if v96 != 0 {
		goto L1
	} else {
		goto L29
	}
L8:
	;
	F__serverAssert(m, int32(_a958), int32(_a917), int32(13301))
	mBase = m.M
	v90 = m.ExcPending
	if v90 != 0 {
		goto L1
	} else {
		goto L28
	}
L9:
	;
	F_hashtableCleanupIterator(m, v6+int32(16))
	mBase = m.M
	v80 = m.ExcPending
	if v80 != 0 {
		goto L1
	} else {
		goto L26
	}
L10:
	;
	if v36 == int32(0) {
		goto L9
	} else {
		goto L11
	}
L11:
	;
	goto L12
L12:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v6)+12))
	v44 = F_moduleFreeCommand(m, l0, v43)
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L1
	} else {
		goto L15
	}
L13:
	;
	goto L9
L14:
	;
	v72 = F_hashtableNext(m, v6+int32(16), v6+int32(12))
	mBase = m.M
	v73 = m.ExcPending
	if v73 != 0 {
		goto L1
	} else {
		goto L24
	}
L15:
	;
	if v44 != 0 {
		goto L14
	} else {
		goto L16
	}
L16:
	;
	v47 = *(*int32)(unsafe.Add(mBase, _consts[318]))
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v43)+140))
	v49 = F_hashtableDelete(m, v47, v48)
	mBase = m.M
	v50 = m.ExcPending
	if v50 != 0 {
		goto L1
	} else {
		goto L17
	}
L17:
	;
	if v49 == int32(0) {
		goto L8
	} else {
		goto L18
	}
L18:
	;
	v54 = *(*int32)(unsafe.Add(mBase, _consts[7]))
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v43)+140))
	v56 = F_hashtableDelete(m, v54, v55)
	mBase = m.M
	v57 = m.ExcPending
	if v57 != 0 {
		goto L1
	} else {
		goto L19
	}
L19:
	;
	if v56 == int32(0) {
		goto L7
	} else {
		goto L20
	}
L20:
	;
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v43)))
	F_sdsfree(m, v60)
	mBase = m.M
	v62 = m.ExcPending
	if v62 != 0 {
		goto L1
	} else {
		goto L21
	}
L21:
	;
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v43)+140))
	F_sdsfree(m, v63)
	mBase = m.M
	v65 = m.ExcPending
	if v65 != 0 {
		goto L1
	} else {
		goto L22
	}
L22:
	;
	F_valkey_free(m, v43)
	mBase = m.M
	v67 = m.ExcPending
	if v67 != 0 {
		goto L1
	} else {
		goto L23
	}
L23:
	;
	goto L14
L24:
	;
	if v72 != 0 {
		goto L12
	} else {
		goto L25
	}
L25:
	;
	goto L13
L26:
	;
	F_invalidateCommandCache(m)
	mBase = m.M
	v82 = m.ExcPending
	if v82 != 0 {
		goto L1
	} else {
		goto L27
	}
L27:
	;
	m.G0 = v6 + int32(64)
	return
L28:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L29:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_moduleUnregisterSharedAPI(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v45 int64
	_ = v45
	var v46 int64
	_ = v46
	var v47 int64
	_ = v47
	var v48 int64
	_ = v48
	var v49 int64
	_ = v49
	var v50 int64
	_ = v50
	var v51 int64
	_ = v51
	var v53 int64
	_ = v53
	var v55 int64
	_ = v55
	var v57 int64
	_ = v57
	var v59 int64
	_ = v59
	var v61 int64
	_ = v61
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v83 int32
	_ = v83
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v103 int32
	_ = v103
	var v108 int32
	_ = v108
	var v112 int32
	_ = v112
	var v116 int32
	_ = v116
	var v123 int32
	_ = v123
	var v125 int32
	_ = v125
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
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
	var v139 int32
	_ = v139
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
	var v159 int32
	_ = v159
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v168 int32
	_ = v168
	var v170 int32
	_ = v170
	var v172 int64
	_ = v172
	var v173 int64
	_ = v173
	var v174 int64
	_ = v174
	var v175 int64
	_ = v175
	var v176 int64
	_ = v176
	var v177 int64
	_ = v177
	var v178 int64
	_ = v178
	var v180 int64
	_ = v180
	var v182 int64
	_ = v182
	var v184 int64
	_ = v184
	var v186 int64
	_ = v186
	var v188 int64
	_ = v188
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v195 int32
	_ = v195
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v202 int32
	_ = v202
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v210 int32
	_ = v210
	var v214 int32
	_ = v214
	var v216 int32
	_ = v216
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v230 int32
	_ = v230
	var v235 int32
	_ = v235
	var v239 int32
	_ = v239
	var v243 int32
	_ = v243
	var v248 int32
	_ = v248
	var v253 int32
	_ = v253
	v2 = int32(0)
	v8 = *(*int32)(unsafe.Add(mBase, _consts[578]))
	v9 = F_dictGetSafeIterator(m, v8)
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	F_dictReleaseIterator(m, v9)
	mBase = m.M
	v253 = m.ExcPending
	if v253 != 0 {
		goto L2
	} else {
		goto L66
	}
L2:
	;
	return int32(0)
L3:
	;
	v20 = v9 + int32(20)
	v21 = *(*int32)(unsafe.Add(mBase, uint32(v9)+16))
	if v21 != 0 {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	if v116 == int32(0) {
		v248 = v2
		goto L1
	} else {
		goto L30
	}
L5:
	;
	v27 = v20
	v28 = v24
	goto L8
L6:
	;
	v24 = int32(1)
	goto L5
L7:
	;
	v24 = int32(0)
	goto L5
L8:
	;
	switch v28 {
	case 0:
		goto L13
	default:
		goto L12
	}
L10:
	;
	v28 = int32(0)
	goto L8
L11:
	;
	goto L4
L12:
	;
	v108 = *(*int32)(unsafe.Add(mBase, uint32(v27)))
	*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = v108
	if v108 == int32(0) {
		goto L10
	} else {
		goto L29
	}
L13:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v9)+4))
	if v32 != int32(-1) {
		v71 = v32
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v72 = int32(1)
	v73 = v71 + v72
	*(*int32)(unsafe.Add(mBase, uint32(v9)+4)) = v73
	v75 = int32(0)
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v9)))
	v79 = *(*int32)(unsafe.Add(mBase, uint32(v9)+8))
	v83 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v78+v79+int32(26)))))
	if v83 == int32(255) {
		goto L23
	} else {
		goto L24
	}
L15:
	;
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v9)+8))
	if v36 != 0 {
		v71 = int32(-1)
		goto L14
	} else {
		goto L16
	}
L16:
	;
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v9)))
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v9)+12))
	if v38 == int32(0) {
		goto L18
	} else {
		goto L19
	}
L17:
	;
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v64)+20))
	if v65 != int32(-1) {
		goto L20
	} else {
		goto L21
	}
L18:
	;
	v45 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v37)+16)))
	v46 = int64(*(*int8)(unsafe.Add(mBase, uint32(v37)+27)))
	v47 = int64(*(*int32)(unsafe.Add(mBase, uint32(v37)+8)))
	v48 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v37)+12)))
	v49 = int64(*(*int8)(unsafe.Add(mBase, uint32(v37)+26)))
	v50 = int64(*(*int32)(unsafe.Add(mBase, uint32(v37)+4)))
	v51 = F_wangHash64(m, v50)
	mBase = m.M
	v53 = F_wangHash64(m, v49+v51)
	mBase = m.M
	v55 = F_wangHash64(m, v48+v53)
	mBase = m.M
	v57 = F_wangHash64(m, v47+v55)
	mBase = m.M
	v59 = F_wangHash64(m, v46+v57)
	mBase = m.M
	v61 = F_wangHash64(m, v45+v59)
	mBase = m.M
	*(*int64)(unsafe.Add(mBase, uint32(v9)+24)) = v61
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v9)))
	v64 = v63
	goto L17
L19:
	;
	v41 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v37)+24)))
	v43 = v41 + int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v37)+24)) = uint16(v43)
	v64 = v37
	goto L17
L20:
	;
	v71 = v65 + int32(-1)
	goto L14
L21:
	;
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v9)+4))
	v71 = v68
	goto L14
L22:
	;
	v98 = int32(2)
	v103 = *(*int32)(unsafe.Add(mBase, uint32(v78+v96<<(uint(v98)%32)+int32(4))))
	v27 = v103 + v97<<(uint(v98)%32)
	v28 = int32(1)
	goto L8
L23:
	;
	v87 = v75
	goto L25
L24:
	;
	v87 = v72 << (uint(v83) % 32)
	goto L25
L25:
	;
	if v73 < v87 {
		v96 = v79
		v97 = v73
		goto L22
	} else {
		goto L26
	}
L26:
	;
	if v79 != 0 {
		v116 = v75
		goto L11
	} else {
		goto L27
	}
L27:
	;
	v89 = *(*int32)(unsafe.Add(mBase, uint32(v78)+20))
	if v89 == int32(-1) {
		v116 = v75
		goto L11
	} else {
		goto L28
	}
L28:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v9)+4)) = int64(4294967296)
	v96 = int32(1)
	v97 = int32(0)
	goto L22
L29:
	;
	v112 = *(*int32)(unsafe.Add(mBase, uint32(v108)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v20))) = v112
	v116 = v108
	goto L11
L30:
	;
	v123 = v2
	v125 = v116
	goto L31
L31:
	;
	v127 = *(*int32)(unsafe.Add(mBase, uint32(v125)))
	goto L33
L32:
	;
	v248 = v139
	goto L1
L33:
	;
	v128 = *(*int32)(unsafe.Add(mBase, uint32(v125)+8))
	goto L35
L34:
	;
	v147 = v9 + int32(20)
	v148 = *(*int32)(unsafe.Add(mBase, uint32(v9)+16))
	if v148 != 0 {
		goto L41
	} else {
		goto L42
	}
L35:
	;
	v129 = *(*int32)(unsafe.Add(mBase, uint32(v128)+4))
	if v129 != l0 {
		v139 = v123
		goto L34
	} else {
		goto L36
	}
L36:
	;
	v132 = *(*int32)(unsafe.Add(mBase, _consts[578]))
	v133 = F_dictDelete(m, v132, v127)
	mBase = m.M
	v134 = m.ExcPending
	if v134 != 0 {
		goto L2
	} else {
		goto L37
	}
L37:
	;
	F_valkey_free(m, v128)
	mBase = m.M
	v136 = m.ExcPending
	if v136 != 0 {
		goto L2
	} else {
		goto L38
	}
L38:
	;
	v139 = v123 + int32(1)
	goto L34
L39:
	;
	if v243 != 0 {
		v123 = v139
		v125 = v243
		goto L31
	} else {
		goto L65
	}
L40:
	;
	v154 = v147
	v155 = v151
	goto L43
L41:
	;
	v151 = int32(1)
	goto L40
L42:
	;
	v151 = int32(0)
	goto L40
L43:
	;
	switch v155 {
	case 0:
		goto L48
	default:
		goto L47
	}
L45:
	;
	v155 = int32(0)
	goto L43
L46:
	;
	goto L39
L47:
	;
	v235 = *(*int32)(unsafe.Add(mBase, uint32(v154)))
	*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = v235
	if v235 == int32(0) {
		goto L45
	} else {
		goto L64
	}
L48:
	;
	v159 = *(*int32)(unsafe.Add(mBase, uint32(v9)+4))
	if v159 != int32(-1) {
		v198 = v159
		goto L49
	} else {
		goto L50
	}
L49:
	;
	v199 = int32(1)
	v200 = v198 + v199
	*(*int32)(unsafe.Add(mBase, uint32(v9)+4)) = v200
	v202 = int32(0)
	v205 = *(*int32)(unsafe.Add(mBase, uint32(v9)))
	v206 = *(*int32)(unsafe.Add(mBase, uint32(v9)+8))
	v210 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v205+v206+int32(26)))))
	if v210 == int32(255) {
		goto L58
	} else {
		goto L59
	}
L50:
	;
	v163 = *(*int32)(unsafe.Add(mBase, uint32(v9)+8))
	if v163 != 0 {
		v198 = int32(-1)
		goto L49
	} else {
		goto L51
	}
L51:
	;
	v164 = *(*int32)(unsafe.Add(mBase, uint32(v9)))
	v165 = *(*int32)(unsafe.Add(mBase, uint32(v9)+12))
	if v165 == int32(0) {
		goto L53
	} else {
		goto L54
	}
L52:
	;
	v192 = *(*int32)(unsafe.Add(mBase, uint32(v191)+20))
	if v192 != int32(-1) {
		goto L55
	} else {
		goto L56
	}
L53:
	;
	v172 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v164)+16)))
	v173 = int64(*(*int8)(unsafe.Add(mBase, uint32(v164)+27)))
	v174 = int64(*(*int32)(unsafe.Add(mBase, uint32(v164)+8)))
	v175 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v164)+12)))
	v176 = int64(*(*int8)(unsafe.Add(mBase, uint32(v164)+26)))
	v177 = int64(*(*int32)(unsafe.Add(mBase, uint32(v164)+4)))
	v178 = F_wangHash64(m, v177)
	mBase = m.M
	v180 = F_wangHash64(m, v176+v178)
	mBase = m.M
	v182 = F_wangHash64(m, v175+v180)
	mBase = m.M
	v184 = F_wangHash64(m, v174+v182)
	mBase = m.M
	v186 = F_wangHash64(m, v173+v184)
	mBase = m.M
	v188 = F_wangHash64(m, v172+v186)
	mBase = m.M
	*(*int64)(unsafe.Add(mBase, uint32(v9)+24)) = v188
	v190 = *(*int32)(unsafe.Add(mBase, uint32(v9)))
	v191 = v190
	goto L52
L54:
	;
	v168 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v164)+24)))
	v170 = v168 + int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v164)+24)) = uint16(v170)
	v191 = v164
	goto L52
L55:
	;
	v198 = v192 + int32(-1)
	goto L49
L56:
	;
	v195 = *(*int32)(unsafe.Add(mBase, uint32(v9)+4))
	v198 = v195
	goto L49
L57:
	;
	v225 = int32(2)
	v230 = *(*int32)(unsafe.Add(mBase, uint32(v205+v223<<(uint(v225)%32)+int32(4))))
	v154 = v230 + v224<<(uint(v225)%32)
	v155 = int32(1)
	goto L43
L58:
	;
	v214 = v202
	goto L60
L59:
	;
	v214 = v199 << (uint(v210) % 32)
	goto L60
L60:
	;
	if v200 < v214 {
		v223 = v206
		v224 = v200
		goto L57
	} else {
		goto L61
	}
L61:
	;
	if v206 != 0 {
		v243 = v202
		goto L46
	} else {
		goto L62
	}
L62:
	;
	v216 = *(*int32)(unsafe.Add(mBase, uint32(v205)+20))
	if v216 == int32(-1) {
		v243 = v202
		goto L46
	} else {
		goto L63
	}
L63:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v9)+4)) = int64(4294967296)
	v223 = int32(1)
	v224 = int32(0)
	goto L57
L64:
	;
	v239 = *(*int32)(unsafe.Add(mBase, uint32(v235)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v147))) = v239
	v243 = v235
	goto L46
L65:
	;
	goto L32
L66:
	;
	return v248
}
func F_setModuleUnsignedNumericConfig(m *base.Module, l0 int32, l1 int64, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
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
	var v22 int32
	_ = v22
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v69 int32
	_ = v69
	v4 = int32(0)
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	*(*int32)(unsafe.Add(mBase, uint32(v8)+12)) = v4
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v17 = m.T0[v16].(func(*base.Module, int32, int64, int32, int32) int32)(m, v12, l1, v13, v8+int32(12))
	mBase = m.M
	v18 = *(*int32)(unsafe.Add(mBase, uint32(v8)+12))
	if v18 == v4 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	m.G0 = v8 + int32(16)
	return base.B2i32(v17 == int32(0))
L2:
	;
	v22 = F_objectGetVal(m, v18)
	mBase = m.M
	goto L6
L3:
	;
	F_decrRefCount(m, v18)
	mBase = m.M
	v69 = m.ExcPending
	if v69 != 0 {
		goto L15
	} else {
		goto L16
	}
L4:
	;
	goto L3
L5:
	;
	v53 = v32
	goto L12
L6:
	;
	v28 = int32(_a972)
	v30 = int32(256)
	v32 = v22
	goto L8
L7:
	;
	v43 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v28))) = uint8(v43)
	goto L5
L8:
	;
	v34 = v30 + int32(-1)
	if v34 == int32(0) {
		goto L7
	} else {
		goto L10
	}
L10:
	;
	v37 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v32))))
	*(*uint8)(unsafe.Add(mBase, uint32(v28))) = uint8(v37)
	v39 = int32(1)
	if v37 != 0 {
		v28 = v28 + v39
		v30 = v34
		v32 = v32 + v39
		goto L8
	} else {
		goto L11
	}
L11:
	;
	goto L4
L12:
	;
	v55 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v53))))
	if v55 != 0 {
		v53 = v53 + int32(1)
		goto L12
	} else {
		goto L14
	}
L13:
	;
	goto L4
L14:
	;
	goto L13
L15:
	;
	return int32(0)
L16:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = int32(_a972)
	goto L1
}
