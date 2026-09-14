package p1

import (
	base "github.com/shibukawa/vkmem/internal/aot/vkaot/base"
	"unsafe"
)

func F_zsetConvertAndExpand(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v20 int32
	_ = v20
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v94 int64
	_ = v94
	var v99 int64
	_ = v99
	var v107 int32
	_ = v107
	var v114 int32
	_ = v114
	var v120 float64
	_ = v120
	var v124 int64
	_ = v124
	var v126 float64
	_ = v126
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v133 int64
	_ = v133
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v149 int32
	_ = v149
	var v158 int64
	_ = v158
	var v162 int64
	_ = v162
	var v163 int64
	_ = v163
	var v171 int64
	_ = v171
	var v173 int64
	_ = v173
	var v177 int32
	_ = v177
	var v181 int64
	_ = v181
	var v184 int64
	_ = v184
	var v186 int64
	_ = v186
	var v189 int64
	_ = v189
	var v200 int64
	_ = v200
	var v203 int64
	_ = v203
	var v214 int64
	_ = v214
	var v217 int64
	_ = v217
	var v230 int64
	_ = v230
	var v237 int64
	_ = v237
	var v242 int32
	_ = v242
	var v245 int64
	_ = v245
	var v251 int32
	_ = v251
	var v252 int32
	_ = v252
	var v255 int32
	_ = v255
	var v256 int32
	_ = v256
	var v261 int64
	_ = v261
	var v269 int64
	_ = v269
	var v272 int64
	_ = v272
	var v284 int64
	_ = v284
	var v286 int32
	_ = v286
	var v288 int64
	_ = v288
	var v295 int32
	_ = v295
	var v296 int32
	_ = v296
	var v299 int32
	_ = v299
	var v300 int32
	_ = v300
	var v305 int64
	_ = v305
	var v313 int64
	_ = v313
	var v316 int64
	_ = v316
	var v327 int32
	_ = v327
	var v328 int32
	_ = v328
	var v330 int64
	_ = v330
	var v338 int64
	_ = v338
	var v340 int64
	_ = v340
	var v345 int64
	_ = v345
	var v354 int32
	_ = v354
	var v355 int64
	_ = v355
	var v366 int64
	_ = v366
	var v371 int64
	_ = v371
	var v376 int64
	_ = v376
	var v379 int64
	_ = v379
	var v382 int32
	_ = v382
	var v388 int32
	_ = v388
	var v389 int32
	_ = v389
	var v390 int32
	_ = v390
	var v391 int32
	_ = v391
	var v392 int32
	_ = v392
	var v394 int32
	_ = v394
	var v395 int32
	_ = v395
	var v396 int32
	_ = v396
	var v404 int32
	_ = v404
	var v405 int32
	_ = v405
	var v408 int32
	_ = v408
	var v422 int32
	_ = v422
	var v424 int32
	_ = v424
	var v426 int32
	_ = v426
	var v427 int32
	_ = v427
	var v434 int32
	_ = v434
	var v435 int32
	_ = v435
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
	var v445 int32
	_ = v445
	var v449 int32
	_ = v449
	var v453 int32
	_ = v453
	var v461 int32
	_ = v461
	var v462 int32
	_ = v462
	var v465 int32
	_ = v465
	var v466 int32
	_ = v466
	var v470 float64
	_ = v470
	var v471 int32
	_ = v471
	var v472 int32
	_ = v472
	var v473 int32
	_ = v473
	var v475 int32
	_ = v475
	var v481 int32
	_ = v481
	var v488 int32
	_ = v488
	var v490 int32
	_ = v490
	var v491 int32
	_ = v491
	var v516 int32
	_ = v516
	var v522 int32
	_ = v522
	var v528 int32
	_ = v528
	var v535 int32
	_ = v535
	var v542 int32
	_ = v542
	v12 = m.G0
	v14 = v12 - int32(48)
	m.G0 = v14
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v20 = int32(base.Ui32(v16)>>(uint(int32(4))%32)) & int32(15)
	if v20 == l1 {
		goto L6
	} else {
		goto L7
	}
L1:
	;
	F__serverPanic_1(m, int32(_a_F_zsetConvertAndExpand_0), int32(1399), int32(_a_F_zsetConvertAndExpand_1), int32(0))
	mBase = m.M
	v542 = m.ExcPending
	if v542 != 0 {
		goto L11
	} else {
		goto L90
	}
L2:
	;
	F__serverPanic_1(m, int32(_a_F_zsetConvertAndExpand_0), int32(1378), int32(_a_F_zsetConvertAndExpand_2), int32(0))
	mBase = m.M
	v535 = m.ExcPending
	if v535 != 0 {
		goto L11
	} else {
		goto L89
	}
L3:
	;
	F__serverAssert(m, int32(_a_F_zsetConvertAndExpand_3), int32(_a_F_zsetConvertAndExpand_0), int32(1368))
	mBase = m.M
	v528 = m.ExcPending
	if v528 != 0 {
		goto L11
	} else {
		goto L88
	}
L4:
	;
	F__serverAssert(m, int32(_a_F_zsetConvertAndExpand_4), int32(_a_F_zsetConvertAndExpand_0), int32(857))
	mBase = m.M
	v522 = m.ExcPending
	if v522 != 0 {
		goto L11
	} else {
		goto L87
	}
L5:
	;
	F__serverPanic_1(m, int32(_a_F_zsetConvertAndExpand_0), int32(1343), int32(_a_F_zsetConvertAndExpand_2), int32(0))
	mBase = m.M
	v516 = m.ExcPending
	if v516 != 0 {
		goto L11
	} else {
		goto L86
	}
L6:
	;
	m.G0 = v14 + int32(48)
	return
L7:
	;
	switch v20 + int32(-7) {
	case 0:
		goto L8
	default:
		goto L1
	case 4:
		goto L9
	}
L8:
	;
	v434 = F_lpNew(m, int32(0))
	mBase = m.M
	v435 = m.ExcPending
	if v435 != 0 {
		goto L11
	} else {
		goto L73
	}
L9:
	;
	v24 = F_objectGetVal(m, l0)
	mBase = m.M
	if l1 != int32(7) {
		goto L5
	} else {
		goto L10
	}
L10:
	;
	v28 = F_valkey_malloc(m, int32(8))
	mBase = m.M
	v29 = m.ExcPending
	if v29 != 0 {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	return
L12:
	;
	v31 = F_hashtableCreate(m, int32(_a_F_zsetConvertAndExpand_5))
	mBase = m.M
	v32 = m.ExcPending
	if v32 != 0 {
		goto L11
	} else {
		goto L13
	}
L13:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28))) = v31
	v35 = F_valkey_calloc(m, int32(272))
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
		goto L11
	} else {
		goto L14
	}
L14:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v35)+16)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v28)+4)) = v35
	v40 = F_hashtableExpand(m, v31, l2)
	mBase = m.M
	v41 = m.ExcPending
	if v41 != 0 {
		goto L11
	} else {
		goto L15
	}
L15:
	;
	v43 = F_lpSeek(m, v24, int32(0))
	mBase = m.M
	v44 = m.ExcPending
	if v44 != 0 {
		goto L11
	} else {
		goto L16
	}
L16:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+28)) = v43
	if v43 == int32(0) {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	v422 = F_objectGetVal(m, l0)
	mBase = m.M
	F_valkey_free(m, v422)
	mBase = m.M
	v424 = m.ExcPending
	if v424 != 0 {
		goto L11
	} else {
		goto L71
	}
L18:
	;
	v48 = F_lpNext(m, v24, v43)
	mBase = m.M
	v49 = m.ExcPending
	if v49 != 0 {
		goto L11
	} else {
		goto L19
	}
L19:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+24)) = v48
	if v48 == int32(0) {
		goto L21
	} else {
		goto L22
	}
L20:
	;
	v61 = v43
	v64 = v48
	goto L24
L21:
	;
	F__serverAssertWithInfo(m, int32(0), l0, int32(_a_F_zsetConvertAndExpand_4), int32(_a_F_zsetConvertAndExpand_0), int32(1355))
	mBase = m.M
	v58 = m.ExcPending
	if v58 != 0 {
		goto L11
	} else {
		goto L23
	}
L22:
	;
	goto L20
L23:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L24:
	;
	v75 = F_lpGetValue(m, v64, v14+int32(44), v14+int32(32))
	mBase = m.M
	v76 = m.ExcPending
	if v76 != 0 {
		goto L11
	} else {
		goto L28
	}
L26:
	;
	v131 = F_lpGetValue(m, v61, v14+int32(20), v14+int32(8))
	mBase = m.M
	v132 = m.ExcPending
	if v132 != 0 {
		goto L11
	} else {
		goto L40
	}
L27:
	;
	v124 = *(*int64)(unsafe.Add(mBase, uint32(v14)+32))
	v126 = base.F64_convert_i64_s(v124)
	goto L26
L28:
	;
	if v75 == int32(0) {
		goto L27
	} else {
		goto L29
	}
L29:
	;
	v79 = *(*int32)(unsafe.Add(mBase, uint32(v14)+44))
	v80 = int32(0)
	v84 = m.G0
	v86 = v84 - int32(32)
	m.G0 = v86
	v88 = F___errno_location(m)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v88))) = v80
	v94 = *(*int64)(unsafe.Add(mBase, _c_F_zsetConvertAndExpand[0]))
	*(*int64)(unsafe.Add(mBase, uint32(v86+int32(8)))) = v94
	*(*int64)(unsafe.Add(mBase, uint32(v86)+24)) = int64(0)
	v99 = *(*int64)(unsafe.Add(mBase, _c_F_zsetConvertAndExpand[1]))
	*(*int64)(unsafe.Add(mBase, uint32(v86))) = v99
	F_ffc_from_chars_double_options(m, v86+int32(16), v75, v75+v79, v86+int32(24), v86)
	mBase = m.M
	v107 = *(*int32)(unsafe.Add(mBase, uint32(v86)+20))
	if v107 == v80 {
		goto L31
	} else {
		goto L32
	}
L30:
	;
	v126 = v120
	goto L26
L31:
	;
	goto L36
L32:
	;
	if v107 == int32(2) {
		goto L33
	} else {
		goto L34
	}
L33:
	;
	v114 = int32(68)
	goto L35
L34:
	;
	v114 = int32(28)
	goto L35
L35:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v88))) = v114
	goto L31
L36:
	;
	v120 = *(*float64)(unsafe.Add(mBase, uint32(v86)+24))
	m.G0 = v86 + int32(32)
	goto L30
L38:
	;
	v149 = *(*int32)(unsafe.Add(mBase, _c_F_zsetConvertAndExpand[2]))
	if int32(311) < v149 {
		goto L46
	} else {
		goto L47
	}
L39:
	;
	v136 = *(*int32)(unsafe.Add(mBase, uint32(v14)+20))
	v137 = F_sdsnewlen(m, v131, v136)
	mBase = m.M
	v138 = m.ExcPending
	if v138 != 0 {
		goto L11
	} else {
		goto L43
	}
L40:
	;
	if v131 != 0 {
		goto L39
	} else {
		goto L41
	}
L41:
	;
	v133 = *(*int64)(unsafe.Add(mBase, uint32(v14)+8))
	v134 = F_sdsfromlonglong(m, v133)
	mBase = m.M
	v135 = m.ExcPending
	if v135 != 0 {
		goto L11
	} else {
		goto L42
	}
L42:
	;
	v139 = v134
	goto L38
L43:
	;
	v139 = v137
	goto L38
L44:
	;
	v382 = int32(1)
	if v379 == int64(0) {
		goto L60
	} else {
		goto L61
	}
L45:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_zsetConvertAndExpand[2])) = v354
	v366 = int64(base.Ui64(v355)>>(uint(int64(29))%64))&int64(22906492245) ^ v355
	v371 = v366<<(uint(int64(17))%64)&int64(8202884508482404352) ^ v366
	v376 = v371<<(uint(int64(37))%64)&int64(-2270628950310912) ^ v371
	v379 = int64(base.Ui64(v376)>>(uint(int64(43))%64)) ^ v376
	goto L44
L46:
	;
	if v149 == int32(313) {
		goto L49
	} else {
		goto L50
	}
L47:
	;
	v158 = *(*int64)(unsafe.Add(mBase, uint32(v149<<(uint(int32(3))%32))+uint32(_c_F_zsetConvertAndExpand[3])))
	v354 = v149 + int32(1)
	v355 = v158
	goto L45
L48:
	;
	v242 = int32(0)
	v245 = v237
	goto L54
L49:
	;
	v163 = int64(5489)
	*(*int64)(unsafe.Add(mBase, _c_F_zsetConvertAndExpand[3])) = v163
	v171 = int64(1)
	v173 = v163
	goto L51
L50:
	;
	v162 = *(*int64)(unsafe.Add(mBase, _c_F_zsetConvertAndExpand[3]))
	v237 = v162
	goto L48
L51:
	;
	v177 = int32(3)
	v181 = int64(62)
	v184 = int64(6364136223846793005)
	v186 = (int64(base.Ui64(v173)>>(uint(v181)%64))^v173)*v184 + v171
	*(*int64)(unsafe.Add(mBase, uint32(base.I32_wrap_i64(v171)<<(uint(v177)%32))+uint32(_c_F_zsetConvertAndExpand[3]))) = v186
	v189 = v171 + int64(1)
	v200 = (int64(base.Ui64(v186)>>(uint(v181)%64))^v186)*v184 + v189
	*(*int64)(unsafe.Add(mBase, uint32(base.I32_wrap_i64(v189)<<(uint(v177)%32))+uint32(_c_F_zsetConvertAndExpand[3]))) = v200
	v203 = v171 + int64(2)
	v214 = (int64(base.Ui64(v200)>>(uint(v181)%64))^v200)*v184 + v203
	*(*int64)(unsafe.Add(mBase, uint32(base.I32_wrap_i64(v203)<<(uint(v177)%32))+uint32(_c_F_zsetConvertAndExpand[3]))) = v214
	v217 = v171 + int64(3)
	if v217 == int64(312) {
		v237 = v163
		goto L48
	} else {
		goto L53
	}
L53:
	;
	v230 = (int64(base.Ui64(v214)>>(uint(int64(62))%64))^v214)*int64(6364136223846793005) + v217
	*(*int64)(unsafe.Add(mBase, uint32(base.I32_wrap_i64(v217)<<(uint(int32(3))%32))+uint32(_c_F_zsetConvertAndExpand[3]))) = v230
	v171 = v171 + int64(4)
	v173 = v230
	goto L51
L54:
	;
	v251 = int32(3)
	v252 = v242 << (uint(v251) % 32)
	v255 = int32(1)
	v256 = v242 + v255
	v261 = *(*int64)(unsafe.Add(mBase, uint32(v256<<(uint(v251)%32))+uint32(_c_F_zsetConvertAndExpand[3])))
	v269 = *(*int64)(unsafe.Add(mBase, uint32(base.I32_wrap_i64(v261)&v255<<(uint(v251)%32))+uint32(_c_F_zsetConvertAndExpand[4])))
	v272 = *(*int64)(unsafe.Add(mBase, uint32(v252)+uint32(_c_F_zsetConvertAndExpand[5])))
	*(*int64)(unsafe.Add(mBase, uint32(v252)+uint32(_c_F_zsetConvertAndExpand[3]))) = v269 ^ v272 ^ int64(base.Ui64(v245&int64(-2147483648)|v261&int64(2147483646))>>(uint(int64(1))%64))
	if v256 != int32(156) {
		v242 = v256
		v245 = v261
		goto L54
	} else {
		goto L56
	}
L55:
	;
	v284 = *(*int64)(unsafe.Add(mBase, _c_F_zsetConvertAndExpand[5]))
	v286 = int32(156)
	v288 = v284
	goto L57
L56:
	;
	goto L55
L57:
	;
	v295 = int32(3)
	v296 = v286 << (uint(v295) % 32)
	v299 = int32(1)
	v300 = v286 + v299
	v305 = *(*int64)(unsafe.Add(mBase, uint32(v300<<(uint(v295)%32))+uint32(_c_F_zsetConvertAndExpand[3])))
	v313 = *(*int64)(unsafe.Add(mBase, uint32(base.I32_wrap_i64(v305)&v299<<(uint(v295)%32))+uint32(_c_F_zsetConvertAndExpand[4])))
	v316 = *(*int64)(unsafe.Add(mBase, uint32(v296)+uint32(_c_F_zsetConvertAndExpand[6])))
	*(*int64)(unsafe.Add(mBase, uint32(v296)+uint32(_c_F_zsetConvertAndExpand[3]))) = v313 ^ v316 ^ int64(base.Ui64(v288&int64(-2147483648)|v305&int64(2147483646))>>(uint(int64(1))%64))
	if v300 != int32(311) {
		v286 = v300
		v288 = v305
		goto L57
	} else {
		goto L59
	}
L58:
	;
	v327 = int32(1)
	v328 = int32(0)
	v330 = *(*int64)(unsafe.Add(mBase, _c_F_zsetConvertAndExpand[3]))
	v338 = *(*int64)(unsafe.Add(mBase, uint32(base.I32_wrap_i64(v330)&v327<<(uint(int32(3))%32))+uint32(_c_F_zsetConvertAndExpand[4])))
	v340 = *(*int64)(unsafe.Add(mBase, _c_F_zsetConvertAndExpand[7]))
	v345 = *(*int64)(unsafe.Add(mBase, _c_F_zsetConvertAndExpand[8]))
	*(*int64)(unsafe.Add(mBase, _c_F_zsetConvertAndExpand[8])) = v338 ^ v340 ^ int64(base.Ui64(v330&int64(2147483646)|v345&int64(-2147483648))>>(uint(int64(1))%64))
	v354 = v327
	v355 = v330
	goto L45
L59:
	;
	goto L58
L60:
	;
	v388 = int32(32)
	goto L62
L61:
	;
	v388 = int32(base.Ui32(base.I32_wrap_i64(base.I64_clz(v379)))>>(uint(v382)%32)) + v382
	goto L62
L62:
	;
	v389 = F_zslCreateNode(m, v388, v126, v139)
	mBase = m.M
	v390 = m.ExcPending
	if v390 != 0 {
		goto L11
	} else {
		goto L63
	}
L63:
	;
	v391 = F_zslInsertNode(m, v35, v389)
	mBase = m.M
	v392 = m.ExcPending
	if v392 != 0 {
		goto L11
	} else {
		goto L64
	}
L64:
	;
	F_sdsfree(m, v139)
	mBase = m.M
	v394 = m.ExcPending
	if v394 != 0 {
		goto L11
	} else {
		goto L65
	}
L65:
	;
	v395 = F_hashtableAdd(m, v31, v391)
	mBase = m.M
	v396 = m.ExcPending
	if v396 != 0 {
		goto L11
	} else {
		goto L66
	}
L66:
	;
	if v395 == int32(0) {
		goto L3
	} else {
		goto L67
	}
L67:
	;
	F_zzlNext(m, v24, v14+int32(28), v14+int32(24))
	mBase = m.M
	v404 = m.ExcPending
	if v404 != 0 {
		goto L11
	} else {
		goto L68
	}
L68:
	;
	v405 = *(*int32)(unsafe.Add(mBase, uint32(v14)+28))
	if v405 == int32(0) {
		goto L17
	} else {
		goto L69
	}
L69:
	;
	v408 = *(*int32)(unsafe.Add(mBase, uint32(v14)+24))
	if v408 == int32(0) {
		goto L4
	} else {
		goto L70
	}
L70:
	;
	v61 = v405
	v64 = v408
	goto L24
L71:
	;
	F_objectSetVal(m, l0, v28)
	mBase = m.M
	v426 = m.ExcPending
	if v426 != 0 {
		goto L11
	} else {
		goto L72
	}
L72:
	;
	v427 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v427&int32(-241) | int32(112)
	goto L6
L73:
	;
	if l1 != int32(11) {
		goto L2
	} else {
		goto L74
	}
L74:
	;
	v438 = F_objectGetVal(m, l0)
	mBase = m.M
	v439 = *(*int32)(unsafe.Add(mBase, uint32(v438)))
	F_hashtableRelease(m, v439)
	mBase = m.M
	v441 = m.ExcPending
	if v441 != 0 {
		goto L11
	} else {
		goto L75
	}
L75:
	;
	v442 = *(*int32)(unsafe.Add(mBase, uint32(v438)+4))
	v443 = *(*int32)(unsafe.Add(mBase, uint32(v442)+12))
	F_valkey_free(m, v442)
	mBase = m.M
	v445 = m.ExcPending
	if v445 != 0 {
		goto L11
	} else {
		goto L76
	}
L76:
	;
	if v443 == int32(0) {
		v481 = v434
		goto L77
	} else {
		goto L78
	}
L77:
	;
	F_valkey_free(m, v438)
	mBase = m.M
	v488 = m.ExcPending
	if v488 != 0 {
		goto L11
	} else {
		goto L84
	}
L78:
	;
	v449 = v443
	v453 = v434
	goto L79
L79:
	;
	v461 = v449 + int32(16)
	v462 = *(*int32)(unsafe.Add(mBase, uint32(v461)))
	v465 = v461 + v462<<(uint(int32(3))%32)
	v466 = int32(*(*int8)(unsafe.Add(mBase, uint32(v465))))
	v470 = *(*float64)(unsafe.Add(mBase, uint32(v449)))
	v471 = F_zzlInsertAt(m, v453, int32(0), v465+v466+int32(1), v470)
	mBase = m.M
	v472 = m.ExcPending
	if v472 != 0 {
		goto L11
	} else {
		goto L81
	}
L80:
	;
	v481 = v471
	goto L77
L81:
	;
	v473 = *(*int32)(unsafe.Add(mBase, uint32(v449)+12))
	F_valkey_free(m, v449)
	mBase = m.M
	v475 = m.ExcPending
	if v475 != 0 {
		goto L11
	} else {
		goto L82
	}
L82:
	;
	if v473 != 0 {
		v449 = v473
		v453 = v471
		goto L79
	} else {
		goto L83
	}
L83:
	;
	goto L80
L84:
	;
	F_objectSetVal(m, l0, v481)
	mBase = m.M
	v490 = m.ExcPending
	if v490 != 0 {
		goto L11
	} else {
		goto L85
	}
L85:
	;
	v491 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v491&int32(-241) | int32(176)
	goto L6
L86:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L87:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L88:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L89:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L90:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_zsetHashtableGetKey(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	v3 = l0 + int32(16)
	v4 = *(*int32)(unsafe.Add(mBase, uint32(v3)))
	v7 = v3 + v4<<(uint(int32(3))%32)
	v8 = int32(*(*int8)(unsafe.Add(mBase, uint32(v7))))
	return v7 + v8 + int32(1)
}
func F_zsetLength(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	switch int32(base.Ui32(v2)>>(uint(int32(4))%32))&int32(15) + int32(-7) {
	case 0:
		v9 = F_objectGetVal(m, l0)
		mBase = m.M
		v10 = *(*int32)(unsafe.Add(mBase, uint32(v9)+4))
		v11 = *(*int32)(unsafe.Add(mBase, uint32(v10)))
		return v11
	default:
		F__serverPanic_1(m, int32(_a_F_zsetLength_0), int32(1288), int32(_a_F_zsetLength_1), int32(0))
		mBase = m.M
		v20 = m.ExcPending
		if v20 != 0 {
			return int32(0)
		} else {
			F_abort(m)
			mBase = m.M
			base.Wasm_trap_unreachable()
			for {
			}
		}
	case 4:
		v22 = F_objectGetVal(m, l0)
		mBase = m.M
		v23 = F_lpLength(m, v22)
		mBase = m.M
		v24 = m.ExcPending
		if v24 != 0 {
			return int32(0)
		} else {
			return int32(base.Ui32(v23) >> (uint(int32(1)) % 32))
		}
	}
}
