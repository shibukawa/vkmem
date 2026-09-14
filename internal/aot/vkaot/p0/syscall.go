package p0

import (
	base "github.com/shibukawa/vkmem/internal/aot/vkaot/base"
	"unsafe"
)

func F___syscall_getrusage(m *base.Module, l0 int32, l1 int32) int32 {
	var v8 int32
	_ = v8
	if l1 == int32(0) {
	} else {
		v8 = F__emscripten_memset_bulkmem(m, l1, base.I32_extend8_s(int32(0)), int32(152))
	}
	return int32(0)
}
func F___syscall_mmap2(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int64) int32 {
	mBase := m.M
	_ = mBase
	var v13 int64
	_ = v13
	var v22 int32
	_ = v22
	var v28 int32
	_ = v28
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	var v97 int32
	_ = v97
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v108 int32
	_ = v108
	var __phi108 int32
	_ = __phi108
	var v109 int32
	_ = v109
	var __phi109 int32
	_ = __phi109
	var v115 int32
	_ = v115
	var v118 int32
	_ = v118
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v140 int32
	_ = v140
	var v149 int32
	_ = v149
	var v151 int32
	_ = v151
	var v154 int32
	_ = v154
	var v157 int32
	_ = v157
	var v159 int32
	_ = v159
	var v164 int32
	_ = v164
	var v171 int32
	_ = v171
	var v176 int32
	_ = v176
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v191 int32
	_ = v191
	var v199 int32
	_ = v199
	var v201 int32
	_ = v201
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v212 int32
	_ = v212
	var v214 int32
	_ = v214
	var v221 int32
	_ = v221
	var v223 int32
	_ = v223
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v241 int32
	_ = v241
	var v243 int32
	_ = v243
	var v245 int32
	_ = v245
	var v254 int32
	_ = v254
	var v256 int32
	_ = v256
	var v259 int32
	_ = v259
	var v264 int32
	_ = v264
	var v269 int32
	_ = v269
	var v270 int32
	_ = v270
	var v275 int32
	_ = v275
	var __phi275 int32
	_ = __phi275
	var v276 int32
	_ = v276
	var __phi276 int32
	_ = __phi276
	var v282 int32
	_ = v282
	var v285 int32
	_ = v285
	var v299 int32
	_ = v299
	var v308 int32
	_ = v308
	var v310 int32
	_ = v310
	var v313 int32
	_ = v313
	var v316 int32
	_ = v316
	var v318 int32
	_ = v318
	var v323 int32
	_ = v323
	var v330 int32
	_ = v330
	var v335 int32
	_ = v335
	var v357 int32
	_ = v357
	var v369 int32
	_ = v369
	var v371 int32
	_ = v371
	var v373 int32
	_ = v373
	var v377 int32
	_ = v377
	var v382 int32
	_ = v382
	var v383 int32
	_ = v383
	var v394 int32
	_ = v394
	var v397 int32
	_ = v397
	var v404 int32
	_ = v404
	var v409 int32
	_ = v409
	var v413 int32
	_ = v413
	var v415 int32
	_ = v415
	var v430 int32
	_ = v430
	var v432 int32
	_ = v432
	var v435 int32
	_ = v435
	var v438 int32
	_ = v438
	var v442 int32
	_ = v442
	var v452 int32
	_ = v452
	var v453 int32
	_ = v453
	var v459 int32
	_ = v459
	var v461 int32
	_ = v461
	var v464 int32
	_ = v464
	var v468 int32
	_ = v468
	var v474 int32
	_ = v474
	var v476 int32
	_ = v476
	var v477 int32
	_ = v477
	var v478 int32
	_ = v478
	var v479 int32
	_ = v479
	var v488 int32
	_ = v488
	var v490 int32
	_ = v490
	var v491 int32
	_ = v491
	var v492 int32
	_ = v492
	var v494 int32
	_ = v494
	var v517 int32
	_ = v517
	var v518 int32
	_ = v518
	var v522 int32
	_ = v522
	var v530 int32
	_ = v530
	var v531 int32
	_ = v531
	var v536 int32
	_ = v536
	if l0 == int32(0) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v13 = l5 << (uint(int64(12)) % 64)
	if l3&int32(32) == int32(0) {
		goto L5
	} else {
		goto L6
	}
L2:
	;
	return int32(-28)
L3:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v522)+32)) = l2
	*(*int64)(unsafe.Add(mBase, uint32(v522)+24)) = v13
	*(*int32)(unsafe.Add(mBase, uint32(v522)+16)) = l3
	*(*int32)(unsafe.Add(mBase, uint32(v522)+4)) = l1
	goto L114
L4:
	;
	v517 = F__emscripten_memset_bulkmem(m, v28, base.I32_extend8_s(int32(0)), v22)
	mBase = m.M
	goto L113
L5:
	;
	v33 = F_emscripten_builtin_malloc(m, int32(40))
	mBase = m.M
	v36 = m.Env.X_mmap_js(m, l1, l2, l3, l4, v13, v33+int32(8), v33)
	mBase = m.M
	if v36 < int32(0) {
		goto L11
	} else {
		goto L12
	}
L6:
	;
	v22 = (l1 + int32(15)) & int32(-16)
	goto L8
L7:
	;
	if v28 != 0 {
		goto L4
	} else {
		goto L10
	}
L8:
	;
	v28 = F_internal_memalign(m, int32(65536), v22+int32(40))
	mBase = m.M
	goto L7
L10:
	;
	return int32(-48)
L11:
	;
	if v33 == int32(0) {
		goto L14
	} else {
		goto L15
	}
L12:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v33)+12)) = l4
	v522 = v33
	goto L3
L13:
	;
	return v36
L14:
	;
	goto L13
L15:
	;
	v50 = int32(-8)
	v51 = v33 + v50
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v33+int32(-4))))
	v56 = v54 & v50
	v57 = v51 + v56
	if v54&int32(1) != 0 {
		v181 = v56
		v182 = v51
		goto L16
	} else {
		goto L17
	}
L16:
	;
	if base.Ui32(v57) <= base.Ui32(v182) {
		goto L14
	} else {
		goto L51
	}
L17:
	;
	if v54&int32(2) == int32(0) {
		goto L14
	} else {
		goto L18
	}
L18:
	;
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v51)))
	v65 = v51 - v64
	v67 = *(*int32)(unsafe.Add(mBase, _consts[1207]))
	if base.Ui32(v65) < base.Ui32(v67) {
		goto L14
	} else {
		goto L19
	}
L19:
	;
	v69 = v64 + v56
	v71 = *(*int32)(unsafe.Add(mBase, _consts[1208]))
	if v65 == v71 {
		goto L23
	} else {
		goto L24
	}
L20:
	;
	if v87 == int32(0) {
		v181 = v69
		v182 = v65
		goto L16
	} else {
		goto L39
	}
L21:
	;
	v140 = int32(0)
	goto L20
L22:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v76)+12)) = v73
	*(*int32)(unsafe.Add(mBase, uint32(v73)+8)) = v76
	v181 = v69
	v182 = v65
	goto L16
L23:
	;
	v121 = *(*int32)(unsafe.Add(mBase, uint32(v57)+4))
	v122 = int32(3)
	if v121&v122 != v122 {
		v181 = v69
		v182 = v65
		goto L16
	} else {
		goto L38
	}
L24:
	;
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v65)+12))
	if base.Ui32(int32(255)) < base.Ui32(v64) {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	v87 = *(*int32)(unsafe.Add(mBase, uint32(v65)+24))
	if v73 == v65 {
		goto L28
	} else {
		goto L29
	}
L26:
	;
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v65)+8))
	if v73 != v76 {
		goto L22
	} else {
		goto L27
	}
L27:
	;
	v78 = int32(0)
	v80 = *(*int32)(unsafe.Add(mBase, _consts[1209]))
	*(*int32)(unsafe.Add(mBase, _consts[1209])) = v80 & base.I32_rotl(int32(-2), int32(base.Ui32(v64)>>(uint(int32(3))%32)))
	v181 = v69
	v182 = v65
	goto L16
L28:
	;
	v92 = *(*int32)(unsafe.Add(mBase, uint32(v65)+20))
	if v92 == int32(0) {
		goto L31
	} else {
		goto L32
	}
L29:
	;
	v89 = *(*int32)(unsafe.Add(mBase, uint32(v65)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v89)+12)) = v73
	*(*int32)(unsafe.Add(mBase, uint32(v73)+8)) = v89
	v140 = v73
	goto L20
L30:
	;
	__phi108 = v102
	__phi109 = v103
	v108 = __phi108
	v109 = __phi109
	goto L34
L31:
	;
	v97 = *(*int32)(unsafe.Add(mBase, uint32(v65)+16))
	if v97 == int32(0) {
		goto L21
	} else {
		goto L33
	}
L32:
	;
	v102 = v92
	v103 = v65 + int32(20)
	goto L30
L33:
	;
	v102 = v97
	v103 = v65 + int32(16)
	goto L30
L34:
	;
	v115 = *(*int32)(unsafe.Add(mBase, uint32(v108)+20))
	if v115 != 0 {
		__phi108 = v115
		__phi109 = v108 + int32(20)
		v108 = __phi108
		v109 = __phi109
		goto L34
	} else {
		goto L36
	}
L35:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v109))) = int32(0)
	v140 = v108
	goto L20
L36:
	;
	v118 = *(*int32)(unsafe.Add(mBase, uint32(v108)+16))
	if v118 != 0 {
		__phi108 = v118
		__phi109 = v108 + int32(16)
		v108 = __phi108
		v109 = __phi109
		goto L34
	} else {
		goto L37
	}
L37:
	;
	goto L35
L38:
	;
	*(*int32)(unsafe.Add(mBase, _consts[1210])) = v69
	*(*int32)(unsafe.Add(mBase, uint32(v57)+4)) = v121 & int32(-2)
	*(*int32)(unsafe.Add(mBase, uint32(v65)+4)) = v69 | int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v57))) = v69
	goto L13
L39:
	;
	v149 = *(*int32)(unsafe.Add(mBase, uint32(v65)+28))
	v151 = v149 << (uint(int32(2)) % 32)
	v154 = *(*int32)(unsafe.Add(mBase, uint32(v151)+uint32(_consts[1211])))
	if v65 != v154 {
		goto L41
	} else {
		goto L42
	}
L40:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v140)+24)) = v87
	v171 = *(*int32)(unsafe.Add(mBase, uint32(v65)+16))
	if v171 == int32(0) {
		goto L48
	} else {
		goto L49
	}
L41:
	;
	v164 = *(*int32)(unsafe.Add(mBase, uint32(v87)+16))
	if v164 != v65 {
		goto L45
	} else {
		goto L46
	}
L42:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v151)+uint32(_consts[1211]))) = v140
	if v140 != 0 {
		goto L40
	} else {
		goto L43
	}
L43:
	;
	v157 = int32(0)
	v159 = *(*int32)(unsafe.Add(mBase, _consts[1212]))
	*(*int32)(unsafe.Add(mBase, _consts[1212])) = v159 & base.I32_rotl(int32(-2), v149)
	v181 = v69
	v182 = v65
	goto L16
L44:
	;
	if v140 == int32(0) {
		v181 = v69
		v182 = v65
		goto L16
	} else {
		goto L47
	}
L45:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v87)+20)) = v140
	goto L44
L46:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v87)+16)) = v140
	goto L44
L47:
	;
	goto L40
L48:
	;
	v176 = *(*int32)(unsafe.Add(mBase, uint32(v65)+20))
	if v176 == int32(0) {
		v181 = v69
		v182 = v65
		goto L16
	} else {
		goto L50
	}
L49:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v140)+16)) = v171
	*(*int32)(unsafe.Add(mBase, uint32(v171)+24)) = v140
	goto L48
L50:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v140)+20)) = v176
	*(*int32)(unsafe.Add(mBase, uint32(v176)+24)) = v140
	v181 = v69
	v182 = v65
	goto L16
L51:
	;
	v191 = *(*int32)(unsafe.Add(mBase, uint32(v57)+4))
	if v191&int32(1) == int32(0) {
		goto L14
	} else {
		goto L52
	}
L52:
	;
	if v191&int32(2) != 0 {
		goto L57
	} else {
		goto L58
	}
L53:
	;
	if base.Ui32(int32(255)) < base.Ui32(v357) {
		goto L91
	} else {
		goto L92
	}
L54:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v182)+4)) = v237 | int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v182+v237))) = v237
	if v182 != v221 {
		v357 = v237
		goto L53
	} else {
		goto L90
	}
L55:
	;
	if v254 == int32(0) {
		goto L54
	} else {
		goto L78
	}
L56:
	;
	v299 = int32(0)
	goto L55
L57:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v57)+4)) = v191 & int32(-2)
	*(*int32)(unsafe.Add(mBase, uint32(v182)+4)) = v181 | int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v182+v181))) = v181
	v357 = v181
	goto L53
L58:
	;
	v199 = *(*int32)(unsafe.Add(mBase, _consts[1213]))
	if v57 != v199 {
		goto L59
	} else {
		goto L60
	}
L59:
	;
	v221 = *(*int32)(unsafe.Add(mBase, _consts[1208]))
	if v57 != v221 {
		goto L62
	} else {
		goto L63
	}
L60:
	;
	v201 = int32(0)
	*(*int32)(unsafe.Add(mBase, _consts[1213])) = v182
	v205 = *(*int32)(unsafe.Add(mBase, _consts[1214]))
	v206 = v205 + v181
	*(*int32)(unsafe.Add(mBase, _consts[1214])) = v206
	*(*int32)(unsafe.Add(mBase, uint32(v182)+4)) = v206 | int32(1)
	v212 = *(*int32)(unsafe.Add(mBase, _consts[1208]))
	if v182 != v212 {
		goto L14
	} else {
		goto L61
	}
L61:
	;
	v214 = int32(0)
	*(*int32)(unsafe.Add(mBase, _consts[1210])) = v214
	*(*int32)(unsafe.Add(mBase, _consts[1208])) = v214
	goto L13
L62:
	;
	v237 = v191&int32(-8) + v181
	v238 = *(*int32)(unsafe.Add(mBase, uint32(v57)+12))
	if base.Ui32(int32(255)) < base.Ui32(v191) {
		goto L64
	} else {
		goto L65
	}
L63:
	;
	v223 = int32(0)
	*(*int32)(unsafe.Add(mBase, _consts[1208])) = v182
	v227 = *(*int32)(unsafe.Add(mBase, _consts[1210]))
	v228 = v227 + v181
	*(*int32)(unsafe.Add(mBase, _consts[1210])) = v228
	*(*int32)(unsafe.Add(mBase, uint32(v182)+4)) = v228 | int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v182+v228))) = v228
	goto L13
L64:
	;
	v254 = *(*int32)(unsafe.Add(mBase, uint32(v57)+24))
	if v238 == v57 {
		goto L68
	} else {
		goto L69
	}
L65:
	;
	v241 = *(*int32)(unsafe.Add(mBase, uint32(v57)+8))
	if v238 != v241 {
		goto L66
	} else {
		goto L67
	}
L66:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v241)+12)) = v238
	*(*int32)(unsafe.Add(mBase, uint32(v238)+8)) = v241
	goto L54
L67:
	;
	v243 = int32(0)
	v245 = *(*int32)(unsafe.Add(mBase, _consts[1209]))
	*(*int32)(unsafe.Add(mBase, _consts[1209])) = v245 & base.I32_rotl(int32(-2), int32(base.Ui32(v191)>>(uint(int32(3))%32)))
	goto L54
L68:
	;
	v259 = *(*int32)(unsafe.Add(mBase, uint32(v57)+20))
	if v259 == int32(0) {
		goto L71
	} else {
		goto L72
	}
L69:
	;
	v256 = *(*int32)(unsafe.Add(mBase, uint32(v57)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v256)+12)) = v238
	*(*int32)(unsafe.Add(mBase, uint32(v238)+8)) = v256
	v299 = v238
	goto L55
L70:
	;
	__phi275 = v269
	__phi276 = v270
	v275 = __phi275
	v276 = __phi276
	goto L74
L71:
	;
	v264 = *(*int32)(unsafe.Add(mBase, uint32(v57)+16))
	if v264 == int32(0) {
		goto L56
	} else {
		goto L73
	}
L72:
	;
	v269 = v259
	v270 = v57 + int32(20)
	goto L70
L73:
	;
	v269 = v264
	v270 = v57 + int32(16)
	goto L70
L74:
	;
	v282 = *(*int32)(unsafe.Add(mBase, uint32(v275)+20))
	if v282 != 0 {
		__phi275 = v282
		__phi276 = v275 + int32(20)
		v275 = __phi275
		v276 = __phi276
		goto L74
	} else {
		goto L76
	}
L75:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v276))) = int32(0)
	v299 = v275
	goto L55
L76:
	;
	v285 = *(*int32)(unsafe.Add(mBase, uint32(v275)+16))
	if v285 != 0 {
		__phi275 = v285
		__phi276 = v275 + int32(16)
		v275 = __phi275
		v276 = __phi276
		goto L74
	} else {
		goto L77
	}
L77:
	;
	goto L75
L78:
	;
	v308 = *(*int32)(unsafe.Add(mBase, uint32(v57)+28))
	v310 = v308 << (uint(int32(2)) % 32)
	v313 = *(*int32)(unsafe.Add(mBase, uint32(v310)+uint32(_consts[1211])))
	if v57 != v313 {
		goto L80
	} else {
		goto L81
	}
L79:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v299)+24)) = v254
	v330 = *(*int32)(unsafe.Add(mBase, uint32(v57)+16))
	if v330 == int32(0) {
		goto L87
	} else {
		goto L88
	}
L80:
	;
	v323 = *(*int32)(unsafe.Add(mBase, uint32(v254)+16))
	if v323 != v57 {
		goto L84
	} else {
		goto L85
	}
L81:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v310)+uint32(_consts[1211]))) = v299
	if v299 != 0 {
		goto L79
	} else {
		goto L82
	}
L82:
	;
	v316 = int32(0)
	v318 = *(*int32)(unsafe.Add(mBase, _consts[1212]))
	*(*int32)(unsafe.Add(mBase, _consts[1212])) = v318 & base.I32_rotl(int32(-2), v308)
	goto L54
L83:
	;
	if v299 == int32(0) {
		goto L54
	} else {
		goto L86
	}
L84:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v254)+20)) = v299
	goto L83
L85:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v254)+16)) = v299
	goto L83
L86:
	;
	goto L79
L87:
	;
	v335 = *(*int32)(unsafe.Add(mBase, uint32(v57)+20))
	if v335 == int32(0) {
		goto L54
	} else {
		goto L89
	}
L88:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v299)+16)) = v330
	*(*int32)(unsafe.Add(mBase, uint32(v330)+24)) = v299
	goto L87
L89:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v299)+20)) = v335
	*(*int32)(unsafe.Add(mBase, uint32(v335)+24)) = v299
	goto L54
L90:
	;
	*(*int32)(unsafe.Add(mBase, _consts[1210])) = v237
	goto L13
L91:
	;
	if base.Ui32(int32(16777215)) < base.Ui32(v357) {
		v404 = int32(31)
		goto L96
	} else {
		goto L97
	}
L92:
	;
	v369 = v357 & int32(-8)
	v371 = v369 + int32(9128464)
	v373 = *(*int32)(unsafe.Add(mBase, _consts[1209]))
	v377 = int32(1) << (uint(int32(base.Ui32(v357)>>(uint(int32(3))%32))) % 32)
	if v373&v377 != 0 {
		goto L94
	} else {
		goto L95
	}
L93:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v369)+uint32(_consts[1215]))) = v182
	*(*int32)(unsafe.Add(mBase, uint32(v383)+12)) = v182
	*(*int32)(unsafe.Add(mBase, uint32(v182)+12)) = v371
	*(*int32)(unsafe.Add(mBase, uint32(v182)+8)) = v383
	goto L13
L94:
	;
	v382 = *(*int32)(unsafe.Add(mBase, uint32(v369)+uint32(_consts[1215])))
	v383 = v382
	goto L93
L95:
	;
	*(*int32)(unsafe.Add(mBase, _consts[1209])) = v373 | v377
	v383 = v371
	goto L93
L96:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v182)+28)) = v404
	*(*int64)(unsafe.Add(mBase, uint32(v182)+16)) = int64(0)
	v409 = v404 << (uint(int32(2)) % 32)
	v413 = *(*int32)(unsafe.Add(mBase, _consts[1212]))
	v415 = int32(1) << (uint(v404) % 32)
	if v413&v415 != 0 {
		goto L101
	} else {
		goto L102
	}
L97:
	;
	v394 = base.I32_clz(int32(base.Ui32(v357) >> (uint(int32(8)) % 32)))
	v397 = int32(1)
	v404 = int32(base.Ui32(v357)>>(uint(int32(38)-v394)%32))&v397 - v394<<(uint(v397)%32) + int32(62)
	goto L96
L98:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v182+v476))) = v479
	*(*int32)(unsafe.Add(mBase, uint32(v182)+12)) = v478
	*(*int32)(unsafe.Add(mBase, uint32(v182+v474))) = v477
	v488 = int32(0)
	v490 = *(*int32)(unsafe.Add(mBase, _consts[1216]))
	v491 = int32(-1)
	v492 = v490 + v491
	if v492 != 0 {
		goto L110
	} else {
		goto L111
	}
L99:
	;
	v468 = *(*int32)(unsafe.Add(mBase, uint32(v438)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v468)+12)) = v182
	*(*int32)(unsafe.Add(mBase, uint32(v438)+8)) = v182
	v474 = int32(24)
	v476 = int32(8)
	v477 = int32(0)
	v478 = v438
	v479 = v468
	goto L98
L100:
	;
	v474 = v459
	v476 = v461
	v477 = v182
	v478 = v182
	v479 = v464
	goto L98
L101:
	;
	if v404 == int32(31) {
		goto L103
	} else {
		goto L104
	}
L102:
	;
	*(*int32)(unsafe.Add(mBase, _consts[1212])) = v413 | v415
	*(*int32)(unsafe.Add(mBase, uint32(v409)+uint32(_consts[1211]))) = v182
	v459 = int32(8)
	v461 = int32(24)
	v464 = v409 + int32(9128728)
	goto L100
L103:
	;
	v430 = int32(0)
	goto L105
L104:
	;
	v430 = int32(25) - int32(base.Ui32(v404)>>(uint(int32(1))%32))
	goto L105
L105:
	;
	v432 = *(*int32)(unsafe.Add(mBase, uint32(v409)+uint32(_consts[1211])))
	v435 = v357 << (uint(v430) % 32)
	v438 = v432
	goto L106
L106:
	;
	v442 = *(*int32)(unsafe.Add(mBase, uint32(v438)+4))
	if v442&int32(-8) == v357 {
		goto L99
	} else {
		goto L108
	}
L107:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v452+int32(16)))) = v182
	v459 = int32(8)
	v461 = int32(24)
	v464 = v438
	goto L100
L108:
	;
	v452 = v438 + int32(base.Ui32(v435)>>(uint(int32(29))%32))&int32(4)
	v453 = *(*int32)(unsafe.Add(mBase, uint32(v452)+16))
	if v453 != 0 {
		v435 = v435 << (uint(int32(1)) % 32)
		v438 = v453
		goto L106
	} else {
		goto L109
	}
L109:
	;
	goto L107
L110:
	;
	v494 = v492
	goto L112
L111:
	;
	v494 = v491
	goto L112
L112:
	;
	*(*int32)(unsafe.Add(mBase, _consts[1216])) = v494
	goto L14
L113:
	;
	v518 = v28 + v22
	*(*int32)(unsafe.Add(mBase, uint32(v518))) = v28
	*(*int64)(unsafe.Add(mBase, uint32(v518)+8)) = int64(-4294967295)
	v522 = v518
	goto L3
L114:
	;
	v530 = int32(0)
	v531 = *(*int32)(unsafe.Add(mBase, _consts[1239]))
	*(*int32)(unsafe.Add(mBase, uint32(v522)+36)) = v531
	*(*int32)(unsafe.Add(mBase, _consts[1239])) = v522
	goto L115
L115:
	;
	v536 = *(*int32)(unsafe.Add(mBase, uint32(v522)))
	return v536
}
func F___syscall_prlimit64(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int64
	_ = v7
	if l3 == int32(0) {
	} else {
		v7 = int64(-1)
		*(*int64)(unsafe.Add(mBase, uint32(l3))) = v7
		*(*int64)(unsafe.Add(mBase, uint32(l3+int32(8)))) = v7
	}
	return int32(0)
}
func F___syscall_ret(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	if base.Ui32(l0) < base.Ui32(int32(-4095)) {
		v9 = l0
	} else {
		*(*int32)(unsafe.Add(mBase, _consts[9])) = int32(0) - l0
		v9 = int32(-1)
	}
	return v9
}
func F___syscall_setsid(m *base.Module) int32 {
	return int32(0)
}
func F___syscall_setsockopt(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32 {
	return int32(0)
}
func F___syscall_umask(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v4 int32
	_ = v4
	v2 = int32(0)
	v4 = *(*int32)(unsafe.Add(mBase, _consts[1235]))
	*(*int32)(unsafe.Add(mBase, _consts[1235])) = l0
	return v4
}
