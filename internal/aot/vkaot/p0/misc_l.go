package p0

import (
	base "github.com/shibukawa/vkmem/internal/aot/vkaot/base"
	"unsafe"
)

func F___localtime_r(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v8 int64
	_ = v8
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	v3 = int32(9116396)
	F___lock(m, v3)
	mBase = m.M
	F_do_tzset(m)
	mBase = m.M
	F___unlock(m, v3)
	mBase = m.M
	v8 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	m.Env.X_localtime_js(m, v8, l1)
	mBase = m.M
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l1)+32))
	if v14 != 0 {
		v15 = int32(9116392)
	} else {
		v15 = int32(9116388)
	}
	v16 = *(*int32)(unsafe.Add(mBase, uint32(v15)))
	*(*int32)(unsafe.Add(mBase, uint32(l1)+40)) = v16
	return l1
}
func F_l_alloc(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v29 int32
	_ = v29
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
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v62 int32
	_ = v62
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v73 int32
	_ = v73
	var __phi73 int32
	_ = __phi73
	var v74 int32
	_ = v74
	var __phi74 int32
	_ = __phi74
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v105 int32
	_ = v105
	var v114 int32
	_ = v114
	var v116 int32
	_ = v116
	var v119 int32
	_ = v119
	var v122 int32
	_ = v122
	var v124 int32
	_ = v124
	var v129 int32
	_ = v129
	var v136 int32
	_ = v136
	var v141 int32
	_ = v141
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v156 int32
	_ = v156
	var v164 int32
	_ = v164
	var v166 int32
	_ = v166
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v177 int32
	_ = v177
	var v179 int32
	_ = v179
	var v186 int32
	_ = v186
	var v188 int32
	_ = v188
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v206 int32
	_ = v206
	var v208 int32
	_ = v208
	var v210 int32
	_ = v210
	var v219 int32
	_ = v219
	var v221 int32
	_ = v221
	var v224 int32
	_ = v224
	var v229 int32
	_ = v229
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v240 int32
	_ = v240
	var __phi240 int32
	_ = __phi240
	var v241 int32
	_ = v241
	var __phi241 int32
	_ = __phi241
	var v247 int32
	_ = v247
	var v250 int32
	_ = v250
	var v264 int32
	_ = v264
	var v273 int32
	_ = v273
	var v275 int32
	_ = v275
	var v278 int32
	_ = v278
	var v281 int32
	_ = v281
	var v283 int32
	_ = v283
	var v288 int32
	_ = v288
	var v295 int32
	_ = v295
	var v300 int32
	_ = v300
	var v322 int32
	_ = v322
	var v334 int32
	_ = v334
	var v336 int32
	_ = v336
	var v338 int32
	_ = v338
	var v342 int32
	_ = v342
	var v347 int32
	_ = v347
	var v348 int32
	_ = v348
	var v359 int32
	_ = v359
	var v362 int32
	_ = v362
	var v369 int32
	_ = v369
	var v374 int32
	_ = v374
	var v378 int32
	_ = v378
	var v380 int32
	_ = v380
	var v395 int32
	_ = v395
	var v397 int32
	_ = v397
	var v400 int32
	_ = v400
	var v403 int32
	_ = v403
	var v407 int32
	_ = v407
	var v417 int32
	_ = v417
	var v418 int32
	_ = v418
	var v424 int32
	_ = v424
	var v426 int32
	_ = v426
	var v429 int32
	_ = v429
	var v433 int32
	_ = v433
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
	var v453 int32
	_ = v453
	var v455 int32
	_ = v455
	var v456 int32
	_ = v456
	var v457 int32
	_ = v457
	var v459 int32
	_ = v459
	var v483 int32
	_ = v483
	var v486 int32
	_ = v486
	var v490 int32
	_ = v490
	var v493 int32
	_ = v493
	var v499 int32
	_ = v499
	var v500 int32
	_ = v500
	var v505 int32
	_ = v505
	var v507 int32
	_ = v507
	var v511 int32
	_ = v511
	var v514 int32
	_ = v514
	var v517 int32
	_ = v517
	var v519 int32
	_ = v519
	var v520 int32
	_ = v520
	var v524 int32
	_ = v524
	if l3 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	if l1 != 0 {
		goto L104
	} else {
		goto L105
	}
L2:
	;
	if l1 == int32(0) {
		goto L4
	} else {
		goto L5
	}
L3:
	;
	return int32(0)
L4:
	;
	goto L3
L5:
	;
	v15 = int32(-8)
	v16 = l1 + v15
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l1+int32(-4))))
	v21 = v19 & v15
	v22 = v16 + v21
	if v19&int32(1) != 0 {
		v146 = v21
		v147 = v16
		goto L6
	} else {
		goto L7
	}
L6:
	;
	if base.Ui32(v22) <= base.Ui32(v147) {
		goto L4
	} else {
		goto L41
	}
L7:
	;
	if v19&int32(2) == int32(0) {
		goto L4
	} else {
		goto L8
	}
L8:
	;
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
	v30 = v16 - v29
	v32 = *(*int32)(unsafe.Add(mBase, _c_F_l_alloc[0]))
	if base.Ui32(v30) < base.Ui32(v32) {
		goto L4
	} else {
		goto L9
	}
L9:
	;
	v34 = v29 + v21
	v36 = *(*int32)(unsafe.Add(mBase, _c_F_l_alloc[1]))
	if v30 == v36 {
		goto L13
	} else {
		goto L14
	}
L10:
	;
	if v52 == int32(0) {
		v146 = v34
		v147 = v30
		goto L6
	} else {
		goto L29
	}
L11:
	;
	v105 = int32(0)
	goto L10
L12:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v41)+12)) = v38
	*(*int32)(unsafe.Add(mBase, uint32(v38)+8)) = v41
	v146 = v34
	v147 = v30
	goto L6
L13:
	;
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v22)+4))
	v87 = int32(3)
	if v86&v87 != v87 {
		v146 = v34
		v147 = v30
		goto L6
	} else {
		goto L28
	}
L14:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v30)+12))
	if base.Ui32(int32(255)) < base.Ui32(v29) {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v30)+24))
	if v38 == v30 {
		goto L18
	} else {
		goto L19
	}
L16:
	;
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v30)+8))
	if v38 != v41 {
		goto L12
	} else {
		goto L17
	}
L17:
	;
	v43 = int32(0)
	v45 = *(*int32)(unsafe.Add(mBase, _c_F_l_alloc[2]))
	*(*int32)(unsafe.Add(mBase, _c_F_l_alloc[2])) = v45 & base.I32_rotl(int32(-2), int32(base.Ui32(v29)>>(uint(int32(3))%32)))
	v146 = v34
	v147 = v30
	goto L6
L18:
	;
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v30)+20))
	if v57 == int32(0) {
		goto L21
	} else {
		goto L22
	}
L19:
	;
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v30)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v54)+12)) = v38
	*(*int32)(unsafe.Add(mBase, uint32(v38)+8)) = v54
	v105 = v38
	goto L10
L20:
	;
	__phi73 = v67
	__phi74 = v68
	v73 = __phi73
	v74 = __phi74
	goto L24
L21:
	;
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v30)+16))
	if v62 == int32(0) {
		goto L11
	} else {
		goto L23
	}
L22:
	;
	v67 = v57
	v68 = v30 + int32(20)
	goto L20
L23:
	;
	v67 = v62
	v68 = v30 + int32(16)
	goto L20
L24:
	;
	v80 = *(*int32)(unsafe.Add(mBase, uint32(v73)+20))
	if v80 != 0 {
		__phi73 = v80
		__phi74 = v73 + int32(20)
		v73 = __phi73
		v74 = __phi74
		goto L24
	} else {
		goto L26
	}
L25:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v74))) = int32(0)
	v105 = v73
	goto L10
L26:
	;
	v83 = *(*int32)(unsafe.Add(mBase, uint32(v73)+16))
	if v83 != 0 {
		__phi73 = v83
		__phi74 = v73 + int32(16)
		v73 = __phi73
		v74 = __phi74
		goto L24
	} else {
		goto L27
	}
L27:
	;
	goto L25
L28:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_l_alloc[3])) = v34
	*(*int32)(unsafe.Add(mBase, uint32(v22)+4)) = v86 & int32(-2)
	*(*int32)(unsafe.Add(mBase, uint32(v30)+4)) = v34 | int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v22))) = v34
	goto L3
L29:
	;
	v114 = *(*int32)(unsafe.Add(mBase, uint32(v30)+28))
	v116 = v114 << (uint(int32(2)) % 32)
	v119 = *(*int32)(unsafe.Add(mBase, uint32(v116)+uint32(_c_F_l_alloc[4])))
	if v30 != v119 {
		goto L31
	} else {
		goto L32
	}
L30:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v105)+24)) = v52
	v136 = *(*int32)(unsafe.Add(mBase, uint32(v30)+16))
	if v136 == int32(0) {
		goto L38
	} else {
		goto L39
	}
L31:
	;
	v129 = *(*int32)(unsafe.Add(mBase, uint32(v52)+16))
	if v129 != v30 {
		goto L35
	} else {
		goto L36
	}
L32:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v116)+uint32(_c_F_l_alloc[4]))) = v105
	if v105 != 0 {
		goto L30
	} else {
		goto L33
	}
L33:
	;
	v122 = int32(0)
	v124 = *(*int32)(unsafe.Add(mBase, _c_F_l_alloc[5]))
	*(*int32)(unsafe.Add(mBase, _c_F_l_alloc[5])) = v124 & base.I32_rotl(int32(-2), v114)
	v146 = v34
	v147 = v30
	goto L6
L34:
	;
	if v105 == int32(0) {
		v146 = v34
		v147 = v30
		goto L6
	} else {
		goto L37
	}
L35:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v52)+20)) = v105
	goto L34
L36:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v52)+16)) = v105
	goto L34
L37:
	;
	goto L30
L38:
	;
	v141 = *(*int32)(unsafe.Add(mBase, uint32(v30)+20))
	if v141 == int32(0) {
		v146 = v34
		v147 = v30
		goto L6
	} else {
		goto L40
	}
L39:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v105)+16)) = v136
	*(*int32)(unsafe.Add(mBase, uint32(v136)+24)) = v105
	goto L38
L40:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v105)+20)) = v141
	*(*int32)(unsafe.Add(mBase, uint32(v141)+24)) = v105
	v146 = v34
	v147 = v30
	goto L6
L41:
	;
	v156 = *(*int32)(unsafe.Add(mBase, uint32(v22)+4))
	if v156&int32(1) == int32(0) {
		goto L4
	} else {
		goto L42
	}
L42:
	;
	if v156&int32(2) != 0 {
		goto L47
	} else {
		goto L48
	}
L43:
	;
	if base.Ui32(int32(255)) < base.Ui32(v322) {
		goto L81
	} else {
		goto L82
	}
L44:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v147)+4)) = v202 | int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v147+v202))) = v202
	if v147 != v186 {
		v322 = v202
		goto L43
	} else {
		goto L80
	}
L45:
	;
	if v219 == int32(0) {
		goto L44
	} else {
		goto L68
	}
L46:
	;
	v264 = int32(0)
	goto L45
L47:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+4)) = v156 & int32(-2)
	*(*int32)(unsafe.Add(mBase, uint32(v147)+4)) = v146 | int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v147+v146))) = v146
	v322 = v146
	goto L43
L48:
	;
	v164 = *(*int32)(unsafe.Add(mBase, _c_F_l_alloc[6]))
	if v22 != v164 {
		goto L49
	} else {
		goto L50
	}
L49:
	;
	v186 = *(*int32)(unsafe.Add(mBase, _c_F_l_alloc[1]))
	if v22 != v186 {
		goto L52
	} else {
		goto L53
	}
L50:
	;
	v166 = int32(0)
	*(*int32)(unsafe.Add(mBase, _c_F_l_alloc[6])) = v147
	v170 = *(*int32)(unsafe.Add(mBase, _c_F_l_alloc[7]))
	v171 = v170 + v146
	*(*int32)(unsafe.Add(mBase, _c_F_l_alloc[7])) = v171
	*(*int32)(unsafe.Add(mBase, uint32(v147)+4)) = v171 | int32(1)
	v177 = *(*int32)(unsafe.Add(mBase, _c_F_l_alloc[1]))
	if v147 != v177 {
		goto L4
	} else {
		goto L51
	}
L51:
	;
	v179 = int32(0)
	*(*int32)(unsafe.Add(mBase, _c_F_l_alloc[3])) = v179
	*(*int32)(unsafe.Add(mBase, _c_F_l_alloc[1])) = v179
	goto L3
L52:
	;
	v202 = v156&int32(-8) + v146
	v203 = *(*int32)(unsafe.Add(mBase, uint32(v22)+12))
	if base.Ui32(int32(255)) < base.Ui32(v156) {
		goto L54
	} else {
		goto L55
	}
L53:
	;
	v188 = int32(0)
	*(*int32)(unsafe.Add(mBase, _c_F_l_alloc[1])) = v147
	v192 = *(*int32)(unsafe.Add(mBase, _c_F_l_alloc[3]))
	v193 = v192 + v146
	*(*int32)(unsafe.Add(mBase, _c_F_l_alloc[3])) = v193
	*(*int32)(unsafe.Add(mBase, uint32(v147)+4)) = v193 | int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v147+v193))) = v193
	goto L3
L54:
	;
	v219 = *(*int32)(unsafe.Add(mBase, uint32(v22)+24))
	if v203 == v22 {
		goto L58
	} else {
		goto L59
	}
L55:
	;
	v206 = *(*int32)(unsafe.Add(mBase, uint32(v22)+8))
	if v203 != v206 {
		goto L56
	} else {
		goto L57
	}
L56:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v206)+12)) = v203
	*(*int32)(unsafe.Add(mBase, uint32(v203)+8)) = v206
	goto L44
L57:
	;
	v208 = int32(0)
	v210 = *(*int32)(unsafe.Add(mBase, _c_F_l_alloc[2]))
	*(*int32)(unsafe.Add(mBase, _c_F_l_alloc[2])) = v210 & base.I32_rotl(int32(-2), int32(base.Ui32(v156)>>(uint(int32(3))%32)))
	goto L44
L58:
	;
	v224 = *(*int32)(unsafe.Add(mBase, uint32(v22)+20))
	if v224 == int32(0) {
		goto L61
	} else {
		goto L62
	}
L59:
	;
	v221 = *(*int32)(unsafe.Add(mBase, uint32(v22)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v221)+12)) = v203
	*(*int32)(unsafe.Add(mBase, uint32(v203)+8)) = v221
	v264 = v203
	goto L45
L60:
	;
	__phi240 = v234
	__phi241 = v235
	v240 = __phi240
	v241 = __phi241
	goto L64
L61:
	;
	v229 = *(*int32)(unsafe.Add(mBase, uint32(v22)+16))
	if v229 == int32(0) {
		goto L46
	} else {
		goto L63
	}
L62:
	;
	v234 = v224
	v235 = v22 + int32(20)
	goto L60
L63:
	;
	v234 = v229
	v235 = v22 + int32(16)
	goto L60
L64:
	;
	v247 = *(*int32)(unsafe.Add(mBase, uint32(v240)+20))
	if v247 != 0 {
		__phi240 = v247
		__phi241 = v240 + int32(20)
		v240 = __phi240
		v241 = __phi241
		goto L64
	} else {
		goto L66
	}
L65:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v241))) = int32(0)
	v264 = v240
	goto L45
L66:
	;
	v250 = *(*int32)(unsafe.Add(mBase, uint32(v240)+16))
	if v250 != 0 {
		__phi240 = v250
		__phi241 = v240 + int32(16)
		v240 = __phi240
		v241 = __phi241
		goto L64
	} else {
		goto L67
	}
L67:
	;
	goto L65
L68:
	;
	v273 = *(*int32)(unsafe.Add(mBase, uint32(v22)+28))
	v275 = v273 << (uint(int32(2)) % 32)
	v278 = *(*int32)(unsafe.Add(mBase, uint32(v275)+uint32(_c_F_l_alloc[4])))
	if v22 != v278 {
		goto L70
	} else {
		goto L71
	}
L69:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v264)+24)) = v219
	v295 = *(*int32)(unsafe.Add(mBase, uint32(v22)+16))
	if v295 == int32(0) {
		goto L77
	} else {
		goto L78
	}
L70:
	;
	v288 = *(*int32)(unsafe.Add(mBase, uint32(v219)+16))
	if v288 != v22 {
		goto L74
	} else {
		goto L75
	}
L71:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v275)+uint32(_c_F_l_alloc[4]))) = v264
	if v264 != 0 {
		goto L69
	} else {
		goto L72
	}
L72:
	;
	v281 = int32(0)
	v283 = *(*int32)(unsafe.Add(mBase, _c_F_l_alloc[5]))
	*(*int32)(unsafe.Add(mBase, _c_F_l_alloc[5])) = v283 & base.I32_rotl(int32(-2), v273)
	goto L44
L73:
	;
	if v264 == int32(0) {
		goto L44
	} else {
		goto L76
	}
L74:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v219)+20)) = v264
	goto L73
L75:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v219)+16)) = v264
	goto L73
L76:
	;
	goto L69
L77:
	;
	v300 = *(*int32)(unsafe.Add(mBase, uint32(v22)+20))
	if v300 == int32(0) {
		goto L44
	} else {
		goto L79
	}
L78:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v264)+16)) = v295
	*(*int32)(unsafe.Add(mBase, uint32(v295)+24)) = v264
	goto L77
L79:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v264)+20)) = v300
	*(*int32)(unsafe.Add(mBase, uint32(v300)+24)) = v264
	goto L44
L80:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_l_alloc[3])) = v202
	goto L3
L81:
	;
	if base.Ui32(int32(16777215)) < base.Ui32(v322) {
		v369 = int32(31)
		goto L86
	} else {
		goto L87
	}
L82:
	;
	v334 = v322 & int32(-8)
	v336 = v334 + int32(9128464)
	v338 = *(*int32)(unsafe.Add(mBase, _c_F_l_alloc[2]))
	v342 = int32(1) << (uint(int32(base.Ui32(v322)>>(uint(int32(3))%32))) % 32)
	if v338&v342 != 0 {
		goto L84
	} else {
		goto L85
	}
L83:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v334)+uint32(_c_F_l_alloc[8]))) = v147
	*(*int32)(unsafe.Add(mBase, uint32(v348)+12)) = v147
	*(*int32)(unsafe.Add(mBase, uint32(v147)+12)) = v336
	*(*int32)(unsafe.Add(mBase, uint32(v147)+8)) = v348
	goto L3
L84:
	;
	v347 = *(*int32)(unsafe.Add(mBase, uint32(v334)+uint32(_c_F_l_alloc[8])))
	v348 = v347
	goto L83
L85:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_l_alloc[2])) = v338 | v342
	v348 = v336
	goto L83
L86:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v147)+28)) = v369
	*(*int64)(unsafe.Add(mBase, uint32(v147)+16)) = int64(0)
	v374 = v369 << (uint(int32(2)) % 32)
	v378 = *(*int32)(unsafe.Add(mBase, _c_F_l_alloc[5]))
	v380 = int32(1) << (uint(v369) % 32)
	if v378&v380 != 0 {
		goto L91
	} else {
		goto L92
	}
L87:
	;
	v359 = base.I32_clz(int32(base.Ui32(v322) >> (uint(int32(8)) % 32)))
	v362 = int32(1)
	v369 = int32(base.Ui32(v322)>>(uint(int32(38)-v359)%32))&v362 - v359<<(uint(v362)%32) + int32(62)
	goto L86
L88:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v147+v441))) = v444
	*(*int32)(unsafe.Add(mBase, uint32(v147)+12)) = v443
	*(*int32)(unsafe.Add(mBase, uint32(v147+v439))) = v442
	v453 = int32(0)
	v455 = *(*int32)(unsafe.Add(mBase, _c_F_l_alloc[9]))
	v456 = int32(-1)
	v457 = v455 + v456
	if v457 != 0 {
		goto L100
	} else {
		goto L101
	}
L89:
	;
	v433 = *(*int32)(unsafe.Add(mBase, uint32(v403)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v433)+12)) = v147
	*(*int32)(unsafe.Add(mBase, uint32(v403)+8)) = v147
	v439 = int32(24)
	v441 = int32(8)
	v442 = int32(0)
	v443 = v403
	v444 = v433
	goto L88
L90:
	;
	v439 = v424
	v441 = v426
	v442 = v147
	v443 = v147
	v444 = v429
	goto L88
L91:
	;
	if v369 == int32(31) {
		goto L93
	} else {
		goto L94
	}
L92:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_l_alloc[5])) = v378 | v380
	*(*int32)(unsafe.Add(mBase, uint32(v374)+uint32(_c_F_l_alloc[4]))) = v147
	v424 = int32(8)
	v426 = int32(24)
	v429 = v374 + int32(9128728)
	goto L90
L93:
	;
	v395 = int32(0)
	goto L95
L94:
	;
	v395 = int32(25) - int32(base.Ui32(v369)>>(uint(int32(1))%32))
	goto L95
L95:
	;
	v397 = *(*int32)(unsafe.Add(mBase, uint32(v374)+uint32(_c_F_l_alloc[4])))
	v400 = v322 << (uint(v395) % 32)
	v403 = v397
	goto L96
L96:
	;
	v407 = *(*int32)(unsafe.Add(mBase, uint32(v403)+4))
	if v407&int32(-8) == v322 {
		goto L89
	} else {
		goto L98
	}
L97:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v417+int32(16)))) = v147
	v424 = int32(8)
	v426 = int32(24)
	v429 = v403
	goto L90
L98:
	;
	v417 = v403 + int32(base.Ui32(v400)>>(uint(int32(29))%32))&int32(4)
	v418 = *(*int32)(unsafe.Add(mBase, uint32(v417)+16))
	if v418 != 0 {
		v400 = v400 << (uint(int32(1)) % 32)
		v403 = v418
		goto L96
	} else {
		goto L99
	}
L99:
	;
	goto L97
L100:
	;
	v459 = v457
	goto L102
L101:
	;
	v459 = v456
	goto L102
L102:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_l_alloc[9])) = v459
	goto L4
L103:
	;
	return v524
L104:
	;
	if base.Ui32(l3) < base.Ui32(int32(-64)) {
		goto L106
	} else {
		goto L107
	}
L105:
	;
	v483 = F_emscripten_builtin_malloc(m, l3)
	mBase = m.M
	v524 = v483
	goto L103
L106:
	;
	v490 = int32(-8)
	v493 = int32(11)
	if base.Ui32(l3) < base.Ui32(v493) {
		goto L109
	} else {
		goto L110
	}
L107:
	;
	v486 = F___errno_location(m)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v486))) = int32(48)
	v524 = int32(0)
	goto L103
L108:
	;
	v505 = F_emscripten_builtin_malloc(m, l3)
	mBase = m.M
	if v505 != 0 {
		goto L113
	} else {
		goto L114
	}
L109:
	;
	v499 = int32(16)
	goto L111
L110:
	;
	v499 = (l3 + v493) & v490
	goto L111
L111:
	;
	v500 = F_try_realloc_chunk(m, l1+v490, v499)
	mBase = m.M
	if v500 == int32(0) {
		goto L108
	} else {
		goto L112
	}
L112:
	;
	v524 = v500 + int32(8)
	goto L103
L113:
	;
	v507 = int32(-4)
	v511 = *(*int32)(unsafe.Add(mBase, uint32(l1+v507)))
	if v511&int32(3) != 0 {
		goto L115
	} else {
		goto L116
	}
L114:
	;
	v524 = int32(0)
	goto L103
L115:
	;
	v514 = v507
	goto L117
L116:
	;
	v514 = int32(-8)
	goto L117
L117:
	;
	v517 = v514 + v511&int32(-8)
	if base.Ui32(v517) < base.Ui32(l3) {
		goto L118
	} else {
		goto L119
	}
L118:
	;
	v519 = v517
	goto L120
L119:
	;
	v519 = l3
	goto L120
L120:
	;
	v520 = F___memcpy(m, v505, l1, v519)
	mBase = m.M
	F_emscripten_builtin_free(m, l1)
	mBase = m.M
	v524 = v505
	goto L103
}
func F_lazyFreeErrors(m *base.Module, l0 int32) {
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
	var v10 int32
	_ = v10
	var v15 int32
	_ = v15
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v4 = *(*int32)(unsafe.Add(mBase, uint32(v3)+8))
	F_raxFreeWithCallback(m, v3, int32(102))
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return
	} else {
		v8 = int32(0)
		v10 = *(*int32)(unsafe.Add(mBase, _c_F_lazyFreeErrors[0]))
		*(*int32)(unsafe.Add(mBase, _c_F_lazyFreeErrors[0])) = v10 - v4
		v15 = *(*int32)(unsafe.Add(mBase, _c_F_lazyFreeErrors[1]))
		*(*int32)(unsafe.Add(mBase, _c_F_lazyFreeErrors[1])) = v4 + v15
		return
	}
}
func F_lazyFreeReplicaKeysWithExpire(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
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
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v17 int32
	_ = v17
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(v4)+16))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(v4)+12))
	F_dictRelease(m, v4)
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		return
	} else {
		v9 = int32(0)
		v11 = *(*int32)(unsafe.Add(mBase, _c_F_lazyFreeReplicaKeysWithExpire[0]))
		v12 = v5 + v6
		*(*int32)(unsafe.Add(mBase, _c_F_lazyFreeReplicaKeysWithExpire[0])) = v11 - v12
		v17 = *(*int32)(unsafe.Add(mBase, _c_F_lazyFreeReplicaKeysWithExpire[1]))
		*(*int32)(unsafe.Add(mBase, _c_F_lazyFreeReplicaKeysWithExpire[1])) = v17 + v12
		return
	}
}
func F_lazyFreeReplicationBacklogRefMem(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v8 int64
	_ = v8
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v22 int32
	_ = v22
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(v5)+20))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v8 = *(*int64)(unsafe.Add(mBase, uint32(v7)+8))
	F_listRelease(m, v5)
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		return
	} else {
		F_raxFree(m, v7)
		mBase = m.M
		v12 = m.ExcPending
		if v12 != 0 {
			return
		} else {
			v13 = int32(0)
			v15 = *(*int32)(unsafe.Add(mBase, _c_F_lazyFreeReplicationBacklogRefMem[0]))
			v17 = v6 + base.I32_wrap_i64(v8)
			*(*int32)(unsafe.Add(mBase, _c_F_lazyFreeReplicationBacklogRefMem[0])) = v15 - v17
			v22 = *(*int32)(unsafe.Add(mBase, _c_F_lazyFreeReplicationBacklogRefMem[1]))
			*(*int32)(unsafe.Add(mBase, _c_F_lazyFreeReplicationBacklogRefMem[1])) = v22 + v17
			return
		}
	}
}
func F_lazyfreeFreeDatabase(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v11 int64
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v18 int64
	_ = v18
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v33 int32
	_ = v33
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(v7)+12))
	if v8 == int32(1) {
		v12 = *(*int32)(unsafe.Add(mBase, uint32(v7)+8))
		v13 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
		if v13 != 0 {
			v15 = F_hashtableSize(m, v13)
			mBase = m.M
			v18 = base.I64_extend_i32_u(v15)
		} else {
			v18 = int64(0)
		}
	} else {
		v11 = *(*int64)(unsafe.Add(mBase, uint32(v7)+40))
		v18 = v11
	}
	F_kvstoreRelease(m, v7)
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		return
	} else {
		F_kvstoreRelease(m, v6)
		mBase = m.M
		v22 = m.ExcPending
		if v22 != 0 {
			return
		} else {
			F_kvstoreRelease(m, v5)
			mBase = m.M
			v24 = m.ExcPending
			if v24 != 0 {
				return
			} else {
				v25 = int32(0)
				v27 = *(*int32)(unsafe.Add(mBase, _c_F_lazyfreeFreeDatabase[0]))
				v28 = base.I32_wrap_i64(v18)
				*(*int32)(unsafe.Add(mBase, _c_F_lazyfreeFreeDatabase[0])) = v27 - v28
				v33 = *(*int32)(unsafe.Add(mBase, _c_F_lazyfreeFreeDatabase[1]))
				*(*int32)(unsafe.Add(mBase, _c_F_lazyfreeFreeDatabase[1])) = v33 + v28
				return
			}
		}
	}
}
func F_lazyfreeGetFreedObjectsCount(m *base.Module) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	v2 = *(*int32)(unsafe.Add(mBase, _c_F_lazyfreeGetFreedObjectsCount[0]))
	return v2
}
func F_lcg31(m *base.Module, l0 int32) int32 {
	return (l0*int32(1103515245) + int32(12345)) & int32(2147483647)
}
func F_lcg64(m *base.Module, l0 int64) int64 {
	return l0*int64(6364136223846793005) + int64(1)
}
func F_lcsCommand(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
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
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
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
	var v54 int32
	_ = v54
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v86 int32
	_ = v86
	var v90 int32
	_ = v90
	var v94 int32
	_ = v94
	var v100 int32
	_ = v100
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
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
	var v132 int32
	_ = v132
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
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
	var v146 int32
	_ = v146
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v169 int32
	_ = v169
	var v172 int32
	_ = v172
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v178 int32
	_ = v178
	var v182 int32
	_ = v182
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v209 int32
	_ = v209
	var v212 int32
	_ = v212
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v218 int32
	_ = v218
	var v222 int32
	_ = v222
	var v224 int32
	_ = v224
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
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v249 int32
	_ = v249
	var v252 int32
	_ = v252
	var v254 int32
	_ = v254
	var v255 int32
	_ = v255
	var v256 int32
	_ = v256
	var v258 int32
	_ = v258
	var v262 int32
	_ = v262
	var v264 int32
	_ = v264
	var v265 int32
	_ = v265
	var v266 int32
	_ = v266
	var v268 int32
	_ = v268
	var v269 int32
	_ = v269
	var v270 int32
	_ = v270
	var v271 int32
	_ = v271
	var v276 int32
	_ = v276
	var v277 int32
	_ = v277
	var v280 int32
	_ = v280
	var v281 int32
	_ = v281
	var v284 int32
	_ = v284
	var v285 int32
	_ = v285
	var v286 int32
	_ = v286
	var v289 int32
	_ = v289
	var v291 int32
	_ = v291
	var v295 int32
	_ = v295
	var v299 int32
	_ = v299
	var v300 int32
	_ = v300
	var v301 int64
	_ = v301
	var v306 int32
	_ = v306
	var v307 int32
	_ = v307
	var v308 int32
	_ = v308
	var v309 int32
	_ = v309
	var v311 int32
	_ = v311
	var v312 int32
	_ = v312
	var v314 int32
	_ = v314
	var v315 int32
	_ = v315
	var v317 int32
	_ = v317
	var v324 int32
	_ = v324
	var v327 int32
	_ = v327
	var v329 int32
	_ = v329
	var v339 int32
	_ = v339
	var v340 int32
	_ = v340
	var v341 int32
	_ = v341
	var v362 int32
	_ = v362
	var v364 int32
	_ = v364
	var v369 int32
	_ = v369
	var v372 int32
	_ = v372
	var v373 int32
	_ = v373
	var v379 int32
	_ = v379
	var v381 int32
	_ = v381
	var v386 int32
	_ = v386
	var v389 int32
	_ = v389
	var v390 int32
	_ = v390
	var v398 int32
	_ = v398
	var v400 int32
	_ = v400
	var v406 int32
	_ = v406
	var v409 int32
	_ = v409
	var v412 int32
	_ = v412
	var v415 int32
	_ = v415
	var v416 int32
	_ = v416
	var v421 int32
	_ = v421
	var v424 int32
	_ = v424
	var v427 int32
	_ = v427
	var v430 int32
	_ = v430
	var v431 int32
	_ = v431
	var v432 int32
	_ = v432
	var v433 int32
	_ = v433
	var v438 int64
	_ = v438
	var v440 int64
	_ = v440
	var v443 int64
	_ = v443
	var v447 int64
	_ = v447
	var v451 int32
	_ = v451
	var v452 int32
	_ = v452
	var v453 int32
	_ = v453
	var v460 int32
	_ = v460
	var v462 int32
	_ = v462
	var v463 int32
	_ = v463
	var v468 int32
	_ = v468
	var v471 int32
	_ = v471
	var v473 int32
	_ = v473
	var v479 int32
	_ = v479
	var v482 int32
	_ = v482
	var v484 int32
	_ = v484
	var v488 int32
	_ = v488
	var v491 int32
	_ = v491
	var v498 int32
	_ = v498
	var v517 int32
	_ = v517
	var v535 int32
	_ = v535
	var v536 int32
	_ = v536
	var v537 int32
	_ = v537
	var v543 int32
	_ = v543
	var v547 int32
	_ = v547
	var v553 int32
	_ = v553
	var v585 int32
	_ = v585
	var v587 int32
	_ = v587
	var v589 int32
	_ = v589
	var v591 int32
	_ = v591
	var v597 int32
	_ = v597
	var v601 int32
	_ = v601
	var v603 int32
	_ = v603
	var v606 int32
	_ = v606
	var v610 int32
	_ = v610
	var v612 int32
	_ = v612
	var v654 int32
	_ = v654
	var v661 int32
	_ = v661
	var v662 int32
	_ = v662
	var v668 int32
	_ = v668
	var v670 int32
	_ = v670
	var v671 int32
	_ = v671
	var v672 int32
	_ = v672
	var v673 int32
	_ = v673
	var v678 int32
	_ = v678
	var v681 int32
	_ = v681
	var v682 int32
	_ = v682
	var v683 int32
	_ = v683
	var v684 int32
	_ = v684
	var v693 int32
	_ = v693
	var v707 int32
	_ = v707
	var v709 int32
	_ = v709
	var v710 int32
	_ = v710
	var v716 int32
	_ = v716
	var v717 int32
	_ = v717
	var v718 int32
	_ = v718
	var v719 int32
	_ = v719
	var v723 int32
	_ = v723
	var v727 int32
	_ = v727
	var v728 int32
	_ = v728
	var v730 int32
	_ = v730
	var v732 int32
	_ = v732
	var v734 int32
	_ = v734
	var v741 int32
	_ = v741
	var v742 int32
	_ = v742
	var v744 int32
	_ = v744
	var v749 int32
	_ = v749
	var v755 int32
	_ = v755
	var v762 int32
	_ = v762
	var v770 int32
	_ = v770
	var v771 int32
	_ = v771
	var v772 int32
	_ = v772
	var v773 int32
	_ = v773
	var v775 int32
	_ = v775
	var v776 int32
	_ = v776
	var v777 int32
	_ = v777
	var v778 int32
	_ = v778
	var v779 int32
	_ = v779
	var v780 int32
	_ = v780
	var v781 int32
	_ = v781
	var v783 int32
	_ = v783
	var v784 int32
	_ = v784
	var v785 int32
	_ = v785
	var v786 int32
	_ = v786
	var v787 int32
	_ = v787
	var v788 int32
	_ = v788
	var v789 int32
	_ = v789
	var v790 int32
	_ = v790
	var v792 int32
	_ = v792
	var v793 int64
	_ = v793
	var v801 int32
	_ = v801
	var v804 int32
	_ = v804
	var v807 int32
	_ = v807
	var v810 int32
	_ = v810
	var v813 int32
	_ = v813
	var v816 int32
	_ = v816
	var v819 int32
	_ = v819
	var v824 int32
	_ = v824
	var v827 int32
	_ = v827
	var v828 int32
	_ = v828
	var v829 int32
	_ = v829
	var v830 int32
	_ = v830
	var v834 int32
	_ = v834
	var v835 int32
	_ = v835
	var v836 int32
	_ = v836
	var v837 int32
	_ = v837
	var v838 int32
	_ = v838
	var v839 int32
	_ = v839
	var v840 int32
	_ = v840
	var v842 int32
	_ = v842
	var v844 int32
	_ = v844
	var v845 int32
	_ = v845
	var v846 int32
	_ = v846
	var v875 int32
	_ = v875
	var v883 int32
	_ = v883
	var v886 int32
	_ = v886
	var v888 int32
	_ = v888
	var v914 int32
	_ = v914
	var v923 int32
	_ = v923
	var v925 int32
	_ = v925
	var v952 int32
	_ = v952
	var v958 int32
	_ = v958
	var v960 int32
	_ = v960
	var v964 int32
	_ = v964
	var v998 int32
	_ = v998
	var v1002 int32
	_ = v1002
	v31 = m.G0
	v33 = v31 - int32(16)
	m.G0 = v33
	*(*int64)(unsafe.Add(mBase, uint32(v33)+8)) = int64(0)
	v37 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v38)+4))
	v40 = F_lookupKeyRead(m, v37, v39)
	mBase = m.M
	v41 = m.ExcPending
	if v41 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v42 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v43 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v43)+8))
	v45 = F_lookupKeyRead(m, v42, v44)
	mBase = m.M
	v46 = m.ExcPending
	if v46 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	if v40 == int32(0) {
		goto L7
	} else {
		goto L8
	}
L4:
	;
	m.G0 = v33 + int32(16)
	return
L5:
	;
	if v40 == int32(0) {
		goto L14
	} else {
		goto L15
	}
L6:
	;
	F_addReplyError(m, l0, int32(_a_F_lcsCommand_0))
	mBase = m.M
	v61 = m.ExcPending
	if v61 != 0 {
		goto L1
	} else {
		goto L12
	}
L7:
	;
	if v45 == int32(0) {
		goto L5
	} else {
		goto L10
	}
L8:
	;
	v49 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v40))))
	if v49&int32(15) != 0 {
		goto L6
	} else {
		goto L9
	}
L9:
	;
	goto L7
L10:
	;
	v54 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v45))))
	if v54&int32(15) == int32(0) {
		goto L5
	} else {
		goto L11
	}
L11:
	;
	goto L6
L12:
	;
	goto L4
L13:
	;
	if v45 == int32(0) {
		goto L19
	} else {
		goto L20
	}
L14:
	;
	v68 = F_createStringObject_1(m, int32(_a_F_lcsCommand_1), int32(0))
	mBase = m.M
	v69 = m.ExcPending
	if v69 != 0 {
		goto L1
	} else {
		goto L17
	}
L15:
	;
	v64 = F_getDecodedObject(m, v40)
	mBase = m.M
	v65 = m.ExcPending
	if v65 != 0 {
		goto L1
	} else {
		goto L16
	}
L16:
	;
	v70 = v64
	goto L13
L17:
	;
	v70 = v68
	goto L13
L18:
	;
	v80 = F_objectGetVal(m, v70)
	mBase = m.M
	v81 = F_objectGetVal(m, v79)
	mBase = m.M
	v83 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if base.Ui32(int32(4)) <= base.Ui32(v83) {
		goto L25
	} else {
		goto L26
	}
L19:
	;
	v77 = F_createStringObject_1(m, int32(_a_F_lcsCommand_1), int32(0))
	mBase = m.M
	v78 = m.ExcPending
	if v78 != 0 {
		goto L1
	} else {
		goto L22
	}
L20:
	;
	v73 = F_getDecodedObject(m, v45)
	mBase = m.M
	v74 = m.ExcPending
	if v74 != 0 {
		goto L1
	} else {
		goto L21
	}
L21:
	;
	v79 = v73
	goto L18
L22:
	;
	v79 = v77
	goto L18
L23:
	;
	if v70 == int32(0) {
		goto L219
	} else {
		goto L220
	}
L24:
	;
	v362 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v80+int32(-1)))))
	v364 = v362 & int32(7)
	switch v364 + int32(-3) {
	case 0:
		goto L100
	case 1:
		goto L99
	default:
		goto L97
	}
L25:
	;
	v90 = int32(0)
	v94 = int32(3)
	v100 = int32(0)
	v103 = v90
	v104 = v90
	goto L28
L26:
	;
	v86 = int32(0)
	v339 = v86
	v340 = v86
	v341 = v86
	goto L24
L27:
	;
	v327 = *(*int32)(unsafe.Add(mBase, _c_F_lcsCommand[0]))
	F_addReplyErrorObject(m, l0, v327)
	mBase = m.M
	v329 = m.ExcPending
	if v329 != 0 {
		goto L1
	} else {
		goto L94
	}
L28:
	;
	v122 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v126 = *(*int32)(unsafe.Add(mBase, uint32(v122+v94<<(uint(int32(2))%32))))
	v127 = F_objectGetVal(m, v126)
	mBase = m.M
	v128 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v129 = int32(_a_F_lcsCommand_2)
	v132 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v127))))
	if v132 != 0 {
		goto L34
	} else {
		goto L35
	}
L29:
	;
	v314 = int32(0)
	v315 = base.B2i32(v309 != v314)
	v317 = base.B2i32(v307 != v314)
	if v307 == v314 {
		v339 = v317
		v340 = v315
		v341 = v308
		goto L24
	} else {
		goto L91
	}
L30:
	;
	v311 = v306 + int32(1)
	v312 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if base.Ui32(v311) < base.Ui32(v312) {
		v94 = v311
		v100 = v307
		v103 = v308
		v104 = v309
		goto L28
	} else {
		goto L90
	}
L31:
	;
	v169 = int32(_a_F_lcsCommand_3)
	v172 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v127))))
	if v172 != 0 {
		goto L48
	} else {
		goto L49
	}
L32:
	;
	if v164-v166 != 0 {
		goto L31
	} else {
		goto L44
	}
L33:
	;
	v164 = F_tolower(m, v160)
	mBase = m.M
	v165 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v161))))
	v166 = F_tolower(m, v165)
	mBase = m.M
	goto L32
L34:
	;
	v134 = v127
	v135 = v129
	v136 = v132
	goto L37
L35:
	;
	v160 = int32(0)
	v161 = v129
	goto L33
L36:
	;
	v160 = v157 & int32(255)
	v161 = v156
	goto L33
L37:
	;
	v138 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v135))))
	if v138 == int32(0) {
		v156 = v135
		v157 = v136
		goto L36
	} else {
		goto L39
	}
L38:
	;
	v156 = v150
	v157 = int32(0)
	goto L36
L39:
	;
	v142 = v136 & int32(255)
	if v142 == v138 {
		goto L40
	} else {
		goto L41
	}
L40:
	;
	v149 = int32(1)
	v150 = v135 + v149
	v151 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v134)+1)))
	if v151 != 0 {
		v134 = v134 + v149
		v135 = v150
		v136 = v151
		goto L37
	} else {
		goto L43
	}
L41:
	;
	v144 = F_tolower(m, v142)
	mBase = m.M
	v145 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v135))))
	v146 = F_tolower(m, v145)
	mBase = m.M
	if v144 == v146 {
		goto L40
	} else {
		goto L42
	}
L42:
	;
	v148 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v134))))
	v156 = v135
	v157 = v148
	goto L36
L43:
	;
	goto L38
L44:
	;
	v306 = v94
	v307 = v100
	v308 = v103
	v309 = int32(1)
	goto L30
L45:
	;
	v209 = int32(_a_F_lcsCommand_4)
	v212 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v127))))
	if v212 != 0 {
		goto L62
	} else {
		goto L63
	}
L46:
	;
	if v204-v206 != 0 {
		goto L45
	} else {
		goto L58
	}
L47:
	;
	v204 = F_tolower(m, v200)
	mBase = m.M
	v205 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v201))))
	v206 = F_tolower(m, v205)
	mBase = m.M
	goto L46
L48:
	;
	v174 = v127
	v175 = v169
	v176 = v172
	goto L51
L49:
	;
	v200 = int32(0)
	v201 = v169
	goto L47
L50:
	;
	v200 = v197 & int32(255)
	v201 = v196
	goto L47
L51:
	;
	v178 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v175))))
	if v178 == int32(0) {
		v196 = v175
		v197 = v176
		goto L50
	} else {
		goto L53
	}
L52:
	;
	v196 = v190
	v197 = int32(0)
	goto L50
L53:
	;
	v182 = v176 & int32(255)
	if v182 == v178 {
		goto L54
	} else {
		goto L55
	}
L54:
	;
	v189 = int32(1)
	v190 = v175 + v189
	v191 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v174)+1)))
	if v191 != 0 {
		v174 = v174 + v189
		v175 = v190
		v176 = v191
		goto L51
	} else {
		goto L57
	}
L55:
	;
	v184 = F_tolower(m, v182)
	mBase = m.M
	v185 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v175))))
	v186 = F_tolower(m, v185)
	mBase = m.M
	if v184 == v186 {
		goto L54
	} else {
		goto L56
	}
L56:
	;
	v188 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v174))))
	v196 = v175
	v197 = v188
	goto L50
L57:
	;
	goto L52
L58:
	;
	v306 = v94
	v307 = int32(1)
	v308 = v103
	v309 = v104
	goto L30
L59:
	;
	v249 = int32(_a_F_lcsCommand_5)
	v252 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v127))))
	if v252 != 0 {
		goto L75
	} else {
		goto L76
	}
L60:
	;
	if v244-v246 != 0 {
		goto L59
	} else {
		goto L72
	}
L61:
	;
	v244 = F_tolower(m, v240)
	mBase = m.M
	v245 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v241))))
	v246 = F_tolower(m, v245)
	mBase = m.M
	goto L60
L62:
	;
	v214 = v127
	v215 = v209
	v216 = v212
	goto L65
L63:
	;
	v240 = int32(0)
	v241 = v209
	goto L61
L64:
	;
	v240 = v237 & int32(255)
	v241 = v236
	goto L61
L65:
	;
	v218 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v215))))
	if v218 == int32(0) {
		v236 = v215
		v237 = v216
		goto L64
	} else {
		goto L67
	}
L66:
	;
	v236 = v230
	v237 = int32(0)
	goto L64
L67:
	;
	v222 = v216 & int32(255)
	if v222 == v218 {
		goto L68
	} else {
		goto L69
	}
L68:
	;
	v229 = int32(1)
	v230 = v215 + v229
	v231 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v214)+1)))
	if v231 != 0 {
		v214 = v214 + v229
		v215 = v230
		v216 = v231
		goto L65
	} else {
		goto L71
	}
L69:
	;
	v224 = F_tolower(m, v222)
	mBase = m.M
	v225 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v215))))
	v226 = F_tolower(m, v225)
	mBase = m.M
	if v224 == v226 {
		goto L68
	} else {
		goto L70
	}
L70:
	;
	v228 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v214))))
	v236 = v215
	v237 = v228
	goto L64
L71:
	;
	goto L66
L72:
	;
	v306 = v94
	v307 = v100
	v308 = int32(1)
	v309 = v104
	goto L30
L73:
	;
	if v284-v286 != 0 {
		goto L27
	} else {
		goto L85
	}
L74:
	;
	v284 = F_tolower(m, v280)
	mBase = m.M
	v285 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v281))))
	v286 = F_tolower(m, v285)
	mBase = m.M
	goto L73
L75:
	;
	v254 = v127
	v255 = v249
	v256 = v252
	goto L78
L76:
	;
	v280 = int32(0)
	v281 = v249
	goto L74
L77:
	;
	v280 = v277 & int32(255)
	v281 = v276
	goto L74
L78:
	;
	v258 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v255))))
	if v258 == int32(0) {
		v276 = v255
		v277 = v256
		goto L77
	} else {
		goto L80
	}
L79:
	;
	v276 = v270
	v277 = int32(0)
	goto L77
L80:
	;
	v262 = v256 & int32(255)
	if v262 == v258 {
		goto L81
	} else {
		goto L82
	}
L81:
	;
	v269 = int32(1)
	v270 = v255 + v269
	v271 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v254)+1)))
	if v271 != 0 {
		v254 = v254 + v269
		v255 = v270
		v256 = v271
		goto L78
	} else {
		goto L84
	}
L82:
	;
	v264 = F_tolower(m, v262)
	mBase = m.M
	v265 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v255))))
	v266 = F_tolower(m, v265)
	mBase = m.M
	if v264 == v266 {
		goto L81
	} else {
		goto L83
	}
L83:
	;
	v268 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v254))))
	v276 = v255
	v277 = v268
	goto L77
L84:
	;
	goto L79
L85:
	;
	v289 = v94 + int32(1)
	if v128 == v289 {
		goto L27
	} else {
		goto L86
	}
L86:
	;
	v291 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v295 = *(*int32)(unsafe.Add(mBase, uint32(v291+v289<<(uint(int32(2))%32))))
	v299 = F_getLongLongFromObjectOrReply(m, l0, v295, v33+int32(8), int32(0))
	mBase = m.M
	v300 = m.ExcPending
	if v300 != 0 {
		goto L1
	} else {
		goto L87
	}
L87:
	;
	if v299 != 0 {
		goto L23
	} else {
		goto L88
	}
L88:
	;
	v301 = *(*int64)(unsafe.Add(mBase, uint32(v33)+8))
	if int64(-1) < v301 {
		v306 = v289
		v307 = v100
		v308 = v103
		v309 = v104
		goto L30
	} else {
		goto L89
	}
L89:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v33)+8)) = int64(0)
	v306 = v289
	v307 = v100
	v308 = v103
	v309 = v104
	goto L30
L90:
	;
	goto L29
L91:
	;
	if v309 == int32(0) {
		v339 = v317
		v340 = v315
		v341 = v308
		goto L24
	} else {
		goto L92
	}
L92:
	;
	F_addReplyError(m, l0, int32(_a_F_lcsCommand_6))
	mBase = m.M
	v324 = m.ExcPending
	if v324 != 0 {
		goto L1
	} else {
		goto L93
	}
L93:
	;
	goto L23
L94:
	;
	goto L23
L95:
	;
	v400 = int32(0)
	switch v364 {
	case 0:
		goto L112
	case 1:
		goto L111
	case 2:
		goto L110
	case 3:
		goto L109
	case 4:
		goto L108
	default:
		v416 = v400
		goto L107
	}
L96:
	;
	F_addReplyError(m, l0, int32(_a_F_lcsCommand_7))
	mBase = m.M
	v398 = m.ExcPending
	if v398 != 0 {
		goto L1
	} else {
		goto L106
	}
L97:
	;
	v379 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v81+int32(-1)))))
	v381 = v379 & int32(7)
	switch v381 + int32(-3) {
	case 0:
		goto L104
	case 1:
		goto L103
	default:
		goto L95
	}
L98:
	;
	if base.Ui32(int32(-3)) < base.Ui32(v373) {
		goto L96
	} else {
		goto L101
	}
L99:
	;
	v372 = *(*int32)(unsafe.Add(mBase, uint32(v80+int32(-17))))
	v373 = v372
	goto L98
L100:
	;
	v369 = *(*int32)(unsafe.Add(mBase, uint32(v80+int32(-9))))
	v373 = v369
	goto L98
L101:
	;
	goto L97
L102:
	;
	if base.Ui32(v390) < base.Ui32(int32(-2)) {
		goto L95
	} else {
		goto L105
	}
L103:
	;
	v389 = *(*int32)(unsafe.Add(mBase, uint32(v81+int32(-17))))
	v390 = v389
	goto L102
L104:
	;
	v386 = *(*int32)(unsafe.Add(mBase, uint32(v81+int32(-9))))
	v390 = v386
	goto L102
L105:
	;
	goto L96
L106:
	;
	goto L23
L107:
	;
	switch v381 {
	case 0:
		goto L118
	case 1:
		goto L117
	case 2:
		goto L116
	case 3:
		goto L115
	case 4:
		goto L114
	default:
		v431 = v400
		goto L113
	}
L108:
	;
	v415 = *(*int32)(unsafe.Add(mBase, uint32(v80+int32(-17))))
	v416 = v415
	goto L107
L109:
	;
	v412 = *(*int32)(unsafe.Add(mBase, uint32(v80+int32(-9))))
	v416 = v412
	goto L107
L110:
	;
	v409 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v80+int32(-5)))))
	v416 = v409
	goto L107
L111:
	;
	v406 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v80+int32(-3)))))
	v416 = v406
	goto L107
L112:
	;
	v416 = int32(base.Ui32(v362) >> (uint(int32(3)) % 32))
	goto L107
L113:
	;
	v432 = int32(1)
	v433 = v431 + v432
	v438 = base.I64_extend_i32_u(v433) * base.I64_extend_i32_u(v416+v432)
	v440 = v438 << (uint(int64(2)) % 64)
	if base.Ui64(int64(4294967294)) < base.Ui64(v440) {
		goto L119
	} else {
		goto L120
	}
L114:
	;
	v430 = *(*int32)(unsafe.Add(mBase, uint32(v81+int32(-17))))
	v431 = v430
	goto L113
L115:
	;
	v427 = *(*int32)(unsafe.Add(mBase, uint32(v81+int32(-9))))
	v431 = v427
	goto L113
L116:
	;
	v424 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v81+int32(-5)))))
	v431 = v424
	goto L113
L117:
	;
	v421 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v81+int32(-3)))))
	v431 = v421
	goto L113
L118:
	;
	v431 = int32(base.Ui32(v379) >> (uint(int32(3)) % 32))
	goto L113
L119:
	;
	F_addReplyError(m, l0, int32(_a_F_lcsCommand_8))
	mBase = m.M
	v964 = m.ExcPending
	if v964 != 0 {
		goto L1
	} else {
		goto L218
	}
L120:
	;
	v443 = base.I64_div_u_s(v440, v438)
	if v443 != int64(4) {
		goto L119
	} else {
		goto L121
	}
L121:
	;
	v447 = int64(*(*uint32)(unsafe.Add(mBase, _c_F_lcsCommand[1])))
	if base.Ui64(v440) <= base.Ui64(v447) {
		goto L122
	} else {
		goto L123
	}
L122:
	;
	v452 = base.I32_wrap_i64(v440)
	v453 = int32(0)
	if base.Ui32(int32(2147483646)) < base.Ui32(v452) {
		v498 = v453
		goto L126
	} else {
		goto L127
	}
L123:
	;
	F_addReplyError(m, l0, int32(_a_F_lcsCommand_9))
	mBase = m.M
	v451 = m.ExcPending
	if v451 != 0 {
		goto L1
	} else {
		goto L124
	}
L124:
	;
	goto L23
L125:
	;
	if v498 == int32(0) {
		goto L119
	} else {
		goto L137
	}
L126:
	;
	goto L125
L127:
	;
	if v452 != 0 {
		goto L128
	} else {
		goto L129
	}
L128:
	;
	v460 = v452
	goto L130
L129:
	;
	v460 = int32(4)
	goto L130
L130:
	;
	v462 = v460 + int32(8)
	v463 = F_emscripten_builtin_malloc(m, v462)
	mBase = m.M
	if v463 == int32(0) {
		v498 = v453
		goto L126
	} else {
		goto L131
	}
L131:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v463))) = v460
	v468 = *(*int32)(unsafe.Add(mBase, _c_F_lcsCommand[2]))
	if v468 != int32(-1) {
		v479 = v468
		goto L132
	} else {
		goto L133
	}
L132:
	;
	if v479 < int32(260) {
		goto L135
	} else {
		goto L136
	}
L133:
	;
	v471 = int32(0)
	v473 = *(*int32)(unsafe.Add(mBase, _c_F_lcsCommand[3]))
	*(*int32)(unsafe.Add(mBase, _c_F_lcsCommand[2])) = v473
	*(*int32)(unsafe.Add(mBase, _c_F_lcsCommand[3])) = v473 + int32(1)
	v479 = v473
	goto L132
L134:
	;
	v498 = v463 + int32(8)
	goto L126
L135:
	;
	v488 = v479 << (uint(int32(2)) % 32)
	v491 = *(*int32)(unsafe.Add(mBase, uint32(v488)+uint32(_c_F_lcsCommand[4])))
	*(*int32)(unsafe.Add(mBase, uint32(v488)+uint32(_c_F_lcsCommand[4]))) = v491 + v462
	goto L134
L136:
	;
	v482 = int32(0)
	v484 = *(*int32)(unsafe.Add(mBase, _c_F_lcsCommand[5]))
	*(*int32)(unsafe.Add(mBase, _c_F_lcsCommand[5])) = v484 + v462
	goto L134
L137:
	;
	v517 = int32(0)
	goto L138
L138:
	;
	v535 = int32(2)
	v536 = v517 * v433 << (uint(v535) % 32)
	v537 = v498 + v536
	*(*int32)(unsafe.Add(mBase, uint32(v537))) = int32(0)
	if base.Ui32(v433) < base.Ui32(v535) {
		goto L140
	} else {
		goto L141
	}
L139:
	;
	v654 = int32(2)
	v661 = *(*int32)(unsafe.Add(mBase, uint32(v498+v431<<(uint(v654)%32)+v433*v416<<(uint(v654)%32))))
	v662 = int32(0)
	if v340|(v339^int32(1)) == v662 {
		v914 = v662
		goto L155
	} else {
		goto L156
	}
L140:
	;
	if base.B2i32(v517 == v416) == int32(0) {
		v517 = v517 + int32(1)
		goto L138
	} else {
		goto L153
	}
L141:
	;
	v543 = v517 + int32(-1)
	v547 = v543 * v433 << (uint(int32(2)) % 32)
	v553 = int32(1)
	goto L142
L142:
	;
	if v517 != 0 {
		goto L145
	} else {
		goto L146
	}
L143:
	;
	goto L140
L144:
	;
	if v431 != v553 {
		v553 = v553 + int32(1)
		goto L142
	} else {
		goto L152
	}
L145:
	;
	v585 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v80+v543))))
	v587 = v553 + int32(-1)
	v589 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v81+v587))))
	if v585 != v589 {
		goto L147
	} else {
		goto L148
	}
L146:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v537+v553<<(uint(int32(2))%32)))) = int32(0)
	goto L144
L147:
	;
	v601 = int32(2)
	v603 = v498 + v553<<(uint(v601)%32)
	v606 = *(*int32)(unsafe.Add(mBase, uint32(v603+v547)))
	v610 = *(*int32)(unsafe.Add(mBase, uint32(v537+v587<<(uint(v601)%32))))
	if base.Ui32(v610) < base.Ui32(v606) {
		goto L149
	} else {
		goto L150
	}
L148:
	;
	v591 = int32(2)
	v597 = *(*int32)(unsafe.Add(mBase, uint32(v498+v547+v587<<(uint(v591)%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v537+v553<<(uint(v591)%32)))) = v597 + int32(1)
	goto L144
L149:
	;
	v612 = v606
	goto L151
L150:
	;
	v612 = v610
	goto L151
L151:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v603+v536))) = v612
	goto L144
L152:
	;
	goto L143
L153:
	;
	goto L139
L154:
	;
	F_sdsfree(m, v952)
	mBase = m.M
	v958 = m.ExcPending
	if v958 != 0 {
		goto L1
	} else {
		goto L216
	}
L155:
	;
	if v339 == int32(0) {
		goto L212
	} else {
		goto L213
	}
L156:
	;
	v668 = int32(0)
	v670 = *(*int32)(unsafe.Add(mBase, _c_F_lcsCommand[6]))
	v671 = F_sdsnewlen(m, v670, v661)
	mBase = m.M
	v672 = m.ExcPending
	if v672 != 0 {
		goto L1
	} else {
		goto L157
	}
L157:
	;
	v673 = int32(0)
	if v340 == v673 {
		v684 = v673
		goto L158
	} else {
		goto L159
	}
L158:
	;
	if v416 == int32(0) {
		v875 = v668
		goto L163
	} else {
		goto L164
	}
L159:
	;
	F_addReplyMapLen(m, l0, int32(2))
	mBase = m.M
	v678 = m.ExcPending
	if v678 != 0 {
		goto L1
	} else {
		goto L160
	}
L160:
	;
	F_addReplyBulkCString(m, l0, int32(_a_F_lcsCommand_10))
	mBase = m.M
	v681 = m.ExcPending
	if v681 != 0 {
		goto L1
	} else {
		goto L161
	}
L161:
	;
	v682 = F_addReplyDeferredLen(m, l0)
	mBase = m.M
	v683 = m.ExcPending
	if v683 != 0 {
		goto L1
	} else {
		goto L162
	}
L162:
	;
	v684 = v682
	goto L158
L163:
	;
	if v684 == int32(0) {
		v914 = v671
		goto L155
	} else {
		goto L208
	}
L164:
	;
	if v431 == int32(0) {
		v875 = v668
		goto L163
	} else {
		goto L165
	}
L165:
	;
	v693 = int32(0)
	v707 = v693
	v709 = v431
	v710 = v416
	v716 = v693
	v717 = v416
	v718 = v693
	v719 = v661
	v723 = v693
	goto L166
L166:
	;
	v727 = int32(-1)
	v728 = v717 + v727
	v730 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v80+v728))))
	v732 = v709 + v727
	v734 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v81+v732))))
	if v730 != v734 {
		goto L173
	} else {
		goto L174
	}
L167:
	;
	v875 = v846
	goto L163
L168:
	;
	if v837 == int32(0) {
		v875 = v846
		goto L163
	} else {
		goto L206
	}
L169:
	;
	v837 = v827
	v838 = v828
	v839 = v829
	v840 = v830
	v842 = v416
	v844 = v834
	v845 = v835
	v846 = v836
	goto L168
L170:
	;
	v792 = v787 + int32(1)
	v793 = *(*int64)(unsafe.Add(mBase, uint32(v33)+8))
	if v793 == int64(0) {
		goto L192
	} else {
		goto L193
	}
L171:
	;
	v783 = v775
	v784 = v776
	v785 = v777
	v786 = v778
	v787 = v777 - v779
	v788 = v779
	v789 = v780
	v790 = v781
	goto L170
L172:
	;
	v775 = v728
	v776 = v732
	v777 = v716
	v778 = v707
	v779 = v710
	v780 = v749
	v781 = v718
	goto L171
L173:
	;
	v755 = int32(2)
	v762 = *(*int32)(unsafe.Add(mBase, uint32(v498+v709<<(uint(v755)%32)+v728*v433<<(uint(v755)%32))))
	v770 = *(*int32)(unsafe.Add(mBase, uint32(v498+v732<<(uint(v755)%32)+v717*v433<<(uint(v755)%32))))
	v771 = base.B2i32(base.Ui32(v770) < base.Ui32(v762))
	if base.Ui32(v770) < base.Ui32(v762) {
		goto L185
	} else {
		goto L186
	}
L174:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v671+int32(-1)+v719))) = uint8(v730)
	if v710 == v416 {
		v741 = v728
		v742 = v732
		goto L176
	} else {
		goto L177
	}
L175:
	;
	v749 = v719 + int32(-1)
	if v710 == int32(0) {
		goto L172
	} else {
		goto L183
	}
L176:
	;
	v744 = v719 + int32(-1)
	if v728 != 0 {
		goto L180
	} else {
		goto L181
	}
L177:
	;
	if v710 != v717 {
		goto L175
	} else {
		goto L178
	}
L178:
	;
	if v707 != v709 {
		goto L175
	} else {
		goto L179
	}
L179:
	;
	v741 = v716
	v742 = v718
	goto L176
L180:
	;
	if v732 == int32(0) {
		v775 = v728
		v776 = v732
		v777 = v741
		v778 = v732
		v779 = v728
		v780 = v744
		v781 = v742
		goto L171
	} else {
		goto L182
	}
L181:
	;
	v783 = v728
	v784 = v732
	v785 = v741
	v786 = v732
	v787 = v741 - v728
	v788 = v728
	v789 = v744
	v790 = v742
	goto L170
L182:
	;
	v837 = v728
	v838 = v732
	v839 = v741
	v840 = v732
	v842 = v728
	v844 = v744
	v845 = v742
	v846 = v723
	goto L168
L183:
	;
	if v707 == int32(0) {
		goto L172
	} else {
		goto L184
	}
L184:
	;
	v783 = v728
	v784 = v732
	v785 = v716
	v786 = v707
	v787 = v716 - v710
	v788 = v710
	v789 = v749
	v790 = v718
	goto L170
L185:
	;
	v772 = v728
	goto L187
L186:
	;
	v772 = v717
	goto L187
L187:
	;
	if base.Ui32(v770) < base.Ui32(v762) {
		goto L188
	} else {
		goto L189
	}
L188:
	;
	v773 = v709
	goto L190
L189:
	;
	v773 = v732
	goto L190
L190:
	;
	if v710 != v416 {
		v775 = v772
		v776 = v773
		v777 = v716
		v778 = v707
		v779 = v710
		v780 = v719
		v781 = v718
		goto L171
	} else {
		goto L191
	}
L191:
	;
	v827 = v772
	v828 = v773
	v829 = v716
	v830 = v707
	v834 = v719
	v835 = v718
	v836 = v723
	goto L169
L192:
	;
	if v684 == int32(0) {
		v827 = v783
		v828 = v784
		v829 = v785
		v830 = v786
		v834 = v789
		v835 = v790
		v836 = v723
		goto L169
	} else {
		goto L195
	}
L193:
	;
	if base.I64_extend_i32_u(v792) < v793 {
		v827 = v783
		v828 = v784
		v829 = v785
		v830 = v786
		v834 = v789
		v835 = v790
		v836 = v723
		goto L169
	} else {
		goto L194
	}
L194:
	;
	goto L192
L195:
	;
	F_addReplyArrayLen(m, l0, v341+int32(2))
	mBase = m.M
	v801 = m.ExcPending
	if v801 != 0 {
		goto L1
	} else {
		goto L196
	}
L196:
	;
	F_addReplyArrayLen(m, l0, int32(2))
	mBase = m.M
	v804 = m.ExcPending
	if v804 != 0 {
		goto L1
	} else {
		goto L197
	}
L197:
	;
	F_addReplyLongLong(m, l0, base.I64_extend_i32_u(v788))
	mBase = m.M
	v807 = m.ExcPending
	if v807 != 0 {
		goto L1
	} else {
		goto L198
	}
L198:
	;
	F_addReplyLongLong(m, l0, base.I64_extend_i32_u(v785))
	mBase = m.M
	v810 = m.ExcPending
	if v810 != 0 {
		goto L1
	} else {
		goto L199
	}
L199:
	;
	F_addReplyArrayLen(m, l0, int32(2))
	mBase = m.M
	v813 = m.ExcPending
	if v813 != 0 {
		goto L1
	} else {
		goto L200
	}
L200:
	;
	F_addReplyLongLong(m, l0, base.I64_extend_i32_u(v786))
	mBase = m.M
	v816 = m.ExcPending
	if v816 != 0 {
		goto L1
	} else {
		goto L201
	}
L201:
	;
	F_addReplyLongLong(m, l0, base.I64_extend_i32_u(v790))
	mBase = m.M
	v819 = m.ExcPending
	if v819 != 0 {
		goto L1
	} else {
		goto L202
	}
L202:
	;
	if v341 == int32(0) {
		goto L203
	} else {
		goto L204
	}
L203:
	;
	v827 = v783
	v828 = v784
	v829 = v785
	v830 = v786
	v834 = v789
	v835 = v790
	v836 = v723 + int32(1)
	goto L169
L204:
	;
	F_addReplyLongLong(m, l0, base.I64_extend_i32_u(v792))
	mBase = m.M
	v824 = m.ExcPending
	if v824 != 0 {
		goto L1
	} else {
		goto L205
	}
L205:
	;
	goto L203
L206:
	;
	if v838 != 0 {
		v707 = v840
		v709 = v838
		v710 = v842
		v716 = v839
		v717 = v837
		v718 = v845
		v719 = v844
		v723 = v846
		goto L166
	} else {
		goto L207
	}
L207:
	;
	goto L167
L208:
	;
	F_addReplyBulkCString(m, l0, int32(_a_F_lcsCommand_11))
	mBase = m.M
	v883 = m.ExcPending
	if v883 != 0 {
		goto L1
	} else {
		goto L209
	}
L209:
	;
	F_addReplyLongLong(m, l0, base.I64_extend_i32_u(v661))
	mBase = m.M
	v886 = m.ExcPending
	if v886 != 0 {
		goto L1
	} else {
		goto L210
	}
L210:
	;
	F_setDeferredArrayLen(m, l0, v684, v875)
	mBase = m.M
	v888 = m.ExcPending
	if v888 != 0 {
		goto L1
	} else {
		goto L211
	}
L211:
	;
	v952 = v671
	goto L154
L212:
	;
	F_addReplyBulkSds(m, l0, v914)
	mBase = m.M
	v925 = m.ExcPending
	if v925 != 0 {
		goto L1
	} else {
		goto L215
	}
L213:
	;
	F_addReplyLongLong(m, l0, base.I64_extend_i32_u(v661))
	mBase = m.M
	v923 = m.ExcPending
	if v923 != 0 {
		goto L1
	} else {
		goto L214
	}
L214:
	;
	v952 = v914
	goto L154
L215:
	;
	v952 = int32(0)
	goto L154
L216:
	;
	F_valkey_free(m, v498)
	mBase = m.M
	v960 = m.ExcPending
	if v960 != 0 {
		goto L1
	} else {
		goto L217
	}
L217:
	;
	goto L23
L218:
	;
	goto L23
L219:
	;
	if v79 == int32(0) {
		goto L4
	} else {
		goto L222
	}
L220:
	;
	F_decrRefCount(m, v70)
	mBase = m.M
	v998 = m.ExcPending
	if v998 != 0 {
		goto L1
	} else {
		goto L221
	}
L221:
	;
	goto L219
L222:
	;
	F_decrRefCount(m, v79)
	mBase = m.M
	v1002 = m.ExcPending
	if v1002 != 0 {
		goto L1
	} else {
		goto L223
	}
L223:
	;
	goto L4
}
func F_ld2string(m *base.Module, l0 int32, l1 int32, l2 int64, l3 int64, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v16 int64
	_ = v16
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v30 int32
	_ = v30
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v41 int64
	_ = v41
	var v46 int32
	_ = v46
	var v50 int64
	_ = v50
	var v51 int64
	_ = v51
	var v55 int32
	_ = v55
	var v79 int32
	_ = v79
	var v88 int32
	_ = v88
	var v94 int32
	_ = v94
	var v98 int32
	_ = v98
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v107 int32
	_ = v107
	var v117 int64
	_ = v117
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v131 int32
	_ = v131
	var v135 int32
	_ = v135
	var v137 int32
	_ = v137
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v146 int32
	_ = v146
	var v154 int32
	_ = v154
	var v157 int32
	_ = v157
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v174 int32
	_ = v174
	var v178 int32
	_ = v178
	var v183 int32
	_ = v183
	var v186 int32
	_ = v186
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v209 int32
	_ = v209
	var v212 int32
	_ = v212
	var v218 int32
	_ = v218
	var v223 int32
	_ = v223
	var v228 int32
	_ = v228
	var v231 int32
	_ = v231
	var v238 int32
	_ = v238
	var v242 int32
	_ = v242
	var v248 int32
	_ = v248
	v8 = m.G0
	v10 = v8 - int32(48)
	m.G0 = v10
	v16 = l3 & int64(281474976710655)
	v20 = int32(32767)
	v21 = base.I32_wrap_i64(int64(base.Ui64(l3)>>(uint(int64(48))%64))) & v20
	if v21 == v20 {
		v34 = base.B2i32(v16|l2 == int64(0))
		v36 = v34
	} else {
		if v21 != 0 {
			v34 = int32(4)
			v36 = v34
		} else {
			if v16|l2 == int64(0) {
				v30 = int32(2)
			} else {
				v30 = int32(3)
			}
			v36 = v30
		}
	}
	if v36 != int32(1) {
		v117 = l3 & int64(281474976710655)
		v121 = int32(32767)
		v122 = base.I32_wrap_i64(int64(base.Ui64(l3)>>(uint(int64(48))%64))) & v121
		if v122 == v121 {
			v135 = base.B2i32(v117|l2 == int64(0))
			v137 = v135
		} else {
			if v122 != 0 {
				v135 = int32(4)
				v137 = v135
			} else {
				if v117|l2 == int64(0) {
					v131 = int32(2)
				} else {
					v131 = int32(3)
				}
				v137 = v131
			}
		}
		if v137 != 0 {
			switch l4 {
			case 0:
				*(*int64)(unsafe.Add(mBase, uint32(v10))) = l2
				*(*int64)(unsafe.Add(mBase, uint32(v10)+8)) = l3
				v203 = F_snprintf(m, l0, l1, int32(_a_F_ld2string_0), v10)
				mBase = m.M
				v204 = m.ExcPending
				if v204 != 0 {
					return int32(0)
				} else {
					if base.Ui32(v203+int32(1)) <= base.Ui32(l1) {
						v238 = v203
						v242 = int32(0)
						*(*uint8)(unsafe.Add(mBase, uint32(l0+v238))) = uint8(v242)
						v248 = v238
					} else {
						v209 = int32(0)
						if l1 == v209 {
							v248 = v209
						} else {
							v212 = int32(0)
							*(*uint8)(unsafe.Add(mBase, uint32(l0))) = uint8(v212)
							v248 = v209
						}
					}
					m.G0 = v10 + int32(48)
					return v248
				}
			case 1:
				*(*int64)(unsafe.Add(mBase, uint32(v10)+32)) = l2
				*(*int64)(unsafe.Add(mBase, uint32(v10)+40)) = l3
				v166 = F_snprintf(m, l0, l1, int32(_a_F_ld2string_1), v10+int32(32))
				mBase = m.M
				v167 = m.ExcPending
				if v167 != 0 {
					return int32(0)
				} else {
					if base.Ui32(l1) < base.Ui32(v166+int32(1)) {
						v209 = int32(0)
						if l1 == v209 {
							v248 = v209
						} else {
							v212 = int32(0)
							*(*uint8)(unsafe.Add(mBase, uint32(l0))) = uint8(v212)
							v248 = v209
						}
					} else {
						v171 = int32(46)
						v172 = F___strchrnul(m, l0, v171)
						mBase = m.M
						v174 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v172))))
						if v174 == v171 {
							v178 = v172
						} else {
							v178 = int32(0)
						}
						if v178 == int32(0) {
							v218 = v166
						} else {
							v183 = l0 + v166
							v186 = v166
							for {
								v190 = v183 + int32(-1)
								v191 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v190))))
								if v191 == int32(48) {
									v183 = v190
									v186 = v186 + int32(-1)
									continue
								} else {
									break
								}
								break
							}
							if v191 != int32(46) {
								v218 = v186
							} else {
								v218 = v186 + int32(-1)
							}
						}
						if v218 != int32(2) {
							v238 = v218
						} else {
							v223 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
							if v223 == int32(45) {
								v228 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+1)))
								if v228 != int32(48) {
									v238 = int32(2)
								} else {
									v231 = int32(48)
									*(*uint8)(unsafe.Add(mBase, uint32(l0))) = uint8(v231)
									v238 = int32(1)
								}
							} else {
								v238 = int32(2)
							}
						}
						v242 = int32(0)
						*(*uint8)(unsafe.Add(mBase, uint32(l0+v238))) = uint8(v242)
						v248 = v238
					}
					m.G0 = v10 + int32(48)
					return v248
				}
			case 2:
				*(*int64)(unsafe.Add(mBase, uint32(v10)+16)) = l2
				*(*int64)(unsafe.Add(mBase, uint32(v10)+24)) = l3
				v154 = F_snprintf(m, l0, l1, int32(_a_F_ld2string_2), v10+int32(16))
				mBase = m.M
				v157 = m.ExcPending
				if v157 != 0 {
					return int32(0)
				} else {
					if base.Ui32(v154+int32(1)) <= base.Ui32(l1) {
						v238 = v154
						v242 = int32(0)
						*(*uint8)(unsafe.Add(mBase, uint32(l0+v238))) = uint8(v242)
						v248 = v238
					} else {
						v209 = int32(0)
						if l1 == v209 {
							v248 = v209
						} else {
							v212 = int32(0)
							*(*uint8)(unsafe.Add(mBase, uint32(l0))) = uint8(v212)
							v248 = v209
						}
					}
					m.G0 = v10 + int32(48)
					return v248
				}
			default:
				v209 = int32(0)
				if l1 == v209 {
					v248 = v209
				} else {
					v212 = int32(0)
					*(*uint8)(unsafe.Add(mBase, uint32(l0))) = uint8(v212)
					v248 = v209
				}
				m.G0 = v10 + int32(48)
				return v248
			}
		} else {
			if base.Ui32(l1) < base.Ui32(int32(4)) {
				v209 = int32(0)
				if l1 == v209 {
					v248 = v209
				} else {
					v212 = int32(0)
					*(*uint8)(unsafe.Add(mBase, uint32(l0))) = uint8(v212)
					v248 = v209
				}
			} else {
				v142 = int32(0)
				v143 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_ld2string[0])))
				*(*uint8)(unsafe.Add(mBase, uint32(l0+int32(2)))) = uint8(v143)
				v146 = int32(*(*uint16)(unsafe.Add(mBase, _c_F_ld2string[1])))
				*(*uint16)(unsafe.Add(mBase, uint32(l0))) = uint16(v146)
				v238 = int32(3)
				v242 = int32(0)
				*(*uint8)(unsafe.Add(mBase, uint32(l0+v238))) = uint8(v242)
				v248 = v238
			}
			m.G0 = v10 + int32(48)
			return v248
		}
	} else {
		if base.Ui32(l1) < base.Ui32(int32(5)) {
			v209 = int32(0)
			if l1 == v209 {
				v248 = v209
			} else {
				v212 = int32(0)
				*(*uint8)(unsafe.Add(mBase, uint32(l0))) = uint8(v212)
				v248 = v209
			}
		} else {
			v41 = int64(0)
			v46 = int32(-1)
			v50 = l3 & int64(9223372036854775807)
			v51 = int64(9223090561878065152)
			if v50 == v51 {
				v55 = base.B2i32(l2 != v41)
			} else {
				v55 = base.B2i32(base.Ui64(v51) < base.Ui64(v50))
			}
			if v55 != 0 {
				v94 = v46
				v98 = v94
			} else {
				if base.B2i32(v41|l2|(int64(0)|v50) == int64(0)) == int32(0) {
					if v41&l3 < int64(0) {
						if l3 == v41 {
							v88 = base.B2i32(base.Ui64(v41) < base.Ui64(l2))
						} else {
							v88 = base.B2i32(v41 < l3)
						}
						if v88 != 0 {
							v94 = v46
						} else {
							v94 = base.B2i32(l2^v41|(l3^v41) != int64(0))
						}
						v98 = v94
					} else {
						if l3 == v41 {
							v79 = base.B2i32(base.Ui64(l2) < base.Ui64(v41))
						} else {
							v79 = base.B2i32(l3 < v41)
						}
						if v79 != 0 {
							v94 = v46
							v98 = v94
						} else {
							v98 = base.B2i32(l2^v41|(l3^v41) != int64(0))
						}
					}
				} else {
					v98 = int32(0)
				}
			}
			if v98 < int32(1) {
				*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(1718511917)
				v238 = int32(4)
			} else {
				v103 = int32(0)
				v104 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_ld2string[2])))
				*(*uint8)(unsafe.Add(mBase, uint32(l0+int32(2)))) = uint8(v104)
				v107 = int32(*(*uint16)(unsafe.Add(mBase, _c_F_ld2string[3])))
				*(*uint16)(unsafe.Add(mBase, uint32(l0))) = uint16(v107)
				v238 = int32(3)
			}
			v242 = int32(0)
			*(*uint8)(unsafe.Add(mBase, uint32(l0+v238))) = uint8(v242)
			v248 = v238
		}
		m.G0 = v10 + int32(48)
		return v248
	}
}
func F_ldexp(m *base.Module, l0 float64, l1 int32) float64 {
	var v6 float64
	_ = v6
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v22 float64
	_ = v22
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v35 float64
	_ = v35
	var v36 int32
	_ = v36
	if l1 < int32(1024) {
		if int32(-1023) < l1 {
			v35 = l0
			v36 = l1
		} else {
			v22 = base.F64_mul(l0, float64(2.004168360008973e-292))
			if base.Ui32(l1) <= base.Ui32(int32(-1992)) {
				v29 = int32(-2960)
				if base.Ui32(v29) < base.Ui32(l1) {
					v32 = l1
				} else {
					v32 = v29
				}
				v35 = base.F64_mul(v22, float64(2.004168360008973e-292))
				v36 = v32 + int32(1938)
			} else {
				v35 = v22
				v36 = l1 + int32(969)
			}
		}
	} else {
		v6 = base.F64_mul(l0, float64(8.98846567431158e+307))
		if base.Ui32(int32(2047)) <= base.Ui32(l1) {
			v13 = int32(3069)
			if base.Ui32(l1) < base.Ui32(v13) {
				v16 = l1
			} else {
				v16 = v13
			}
			v35 = base.F64_mul(v6, float64(8.98846567431158e+307))
			v36 = v16 + int32(-2046)
		} else {
			v35 = v6
			v36 = l1 + int32(-1023)
		}
	}
	return base.F64_mul(v35, base.F64_reinterpret_i64(base.I64_extend_i32_u(v36+int32(1023))<<(uint(int64(52))%64)))
}
func F_libraryLink(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v31 int32
	_ = v31
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
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
	var v50 int64
	_ = v50
	var v52 int64
	_ = v52
	var v54 int64
	_ = v54
	var v56 int64
	_ = v56
	var v58 int64
	_ = v58
	var v60 int64
	_ = v60
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v82 int32
	_ = v82
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v102 int32
	_ = v102
	var v107 int32
	_ = v107
	var v111 int32
	_ = v111
	var v115 int32
	_ = v115
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v132 int32
	_ = v132
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v143 int32
	_ = v143
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v154 int32
	_ = v154
	var v159 int32
	_ = v159
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v175 int32
	_ = v175
	var v182 int32
	_ = v182
	var v186 int32
	_ = v186
	var v190 int32
	_ = v190
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v199 int32
	_ = v199
	var v201 int32
	_ = v201
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v217 int32
	_ = v217
	var v224 int32
	_ = v224
	var v228 int32
	_ = v228
	var v232 int32
	_ = v232
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v241 int32
	_ = v241
	var v243 int32
	_ = v243
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v248 int32
	_ = v248
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v255 int32
	_ = v255
	var v257 int32
	_ = v257
	var v261 int32
	_ = v261
	var v262 int32
	_ = v262
	var v263 int32
	_ = v263
	var v265 int32
	_ = v265
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v6 = F_dictGetIterator(m, v5)
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	goto L4
L3:
	;
	F_dictReleaseIterator(m, v6)
	mBase = m.M
	v147 = m.ExcPending
	if v147 != 0 {
		goto L1
	} else {
		goto L38
	}
L4:
	;
	v19 = v6 + int32(20)
	v20 = *(*int32)(unsafe.Add(mBase, uint32(v6)+16))
	if v20 != 0 {
		goto L8
	} else {
		goto L9
	}
L6:
	;
	if v115 == int32(0) {
		goto L3
	} else {
		goto L32
	}
L7:
	;
	v26 = v19
	v27 = v23
	goto L10
L8:
	;
	v23 = int32(1)
	goto L7
L9:
	;
	v23 = int32(0)
	goto L7
L10:
	;
	switch v27 {
	case 0:
		goto L15
	default:
		goto L14
	}
L12:
	;
	v27 = int32(0)
	goto L10
L13:
	;
	goto L6
L14:
	;
	v107 = *(*int32)(unsafe.Add(mBase, uint32(v26)))
	*(*int32)(unsafe.Add(mBase, uint32(v6)+16)) = v107
	if v107 == int32(0) {
		goto L12
	} else {
		goto L31
	}
L15:
	;
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v6)+4))
	if v31 != int32(-1) {
		v70 = v31
		goto L16
	} else {
		goto L17
	}
L16:
	;
	v71 = int32(1)
	v72 = v70 + v71
	*(*int32)(unsafe.Add(mBase, uint32(v6)+4)) = v72
	v74 = int32(0)
	v77 = *(*int32)(unsafe.Add(mBase, uint32(v6)))
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v6)+8))
	v82 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v77+v78+int32(26)))))
	if v82 == int32(255) {
		goto L25
	} else {
		goto L26
	}
L17:
	;
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v6)+8))
	if v35 != 0 {
		v70 = int32(-1)
		goto L16
	} else {
		goto L18
	}
L18:
	;
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v6)))
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v6)+12))
	if v37 == int32(0) {
		goto L20
	} else {
		goto L21
	}
L19:
	;
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v63)+20))
	if v64 != int32(-1) {
		goto L22
	} else {
		goto L23
	}
L20:
	;
	v44 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v36)+16)))
	v45 = int64(*(*int8)(unsafe.Add(mBase, uint32(v36)+27)))
	v46 = int64(*(*int32)(unsafe.Add(mBase, uint32(v36)+8)))
	v47 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v36)+12)))
	v48 = int64(*(*int8)(unsafe.Add(mBase, uint32(v36)+26)))
	v49 = int64(*(*int32)(unsafe.Add(mBase, uint32(v36)+4)))
	v50 = F_wangHash64(m, v49)
	mBase = m.M
	v52 = F_wangHash64(m, v48+v50)
	mBase = m.M
	v54 = F_wangHash64(m, v47+v52)
	mBase = m.M
	v56 = F_wangHash64(m, v46+v54)
	mBase = m.M
	v58 = F_wangHash64(m, v45+v56)
	mBase = m.M
	v60 = F_wangHash64(m, v44+v58)
	mBase = m.M
	*(*int64)(unsafe.Add(mBase, uint32(v6)+24)) = v60
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v6)))
	v63 = v62
	goto L19
L21:
	;
	v40 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v36)+24)))
	v42 = v40 + int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v36)+24)) = uint16(v42)
	v63 = v36
	goto L19
L22:
	;
	v70 = v64 + int32(-1)
	goto L16
L23:
	;
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v6)+4))
	v70 = v67
	goto L16
L24:
	;
	v97 = int32(2)
	v102 = *(*int32)(unsafe.Add(mBase, uint32(v77+v95<<(uint(v97)%32)+int32(4))))
	v26 = v102 + v96<<(uint(v97)%32)
	v27 = int32(1)
	goto L10
L25:
	;
	v86 = v74
	goto L27
L26:
	;
	v86 = v71 << (uint(v82) % 32)
	goto L27
L27:
	;
	if v72 < v86 {
		v95 = v78
		v96 = v72
		goto L24
	} else {
		goto L28
	}
L28:
	;
	if v78 != 0 {
		v115 = v74
		goto L13
	} else {
		goto L29
	}
L29:
	;
	v88 = *(*int32)(unsafe.Add(mBase, uint32(v77)+20))
	if v88 == int32(-1) {
		v115 = v74
		goto L13
	} else {
		goto L30
	}
L30:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v6)+4)) = int64(4294967296)
	v95 = int32(1)
	v96 = int32(0)
	goto L24
L31:
	;
	v111 = *(*int32)(unsafe.Add(mBase, uint32(v107)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v19))) = v111
	v115 = v107
	goto L13
L32:
	;
	v121 = *(*int32)(unsafe.Add(mBase, uint32(v115)+8))
	goto L33
L33:
	;
	v122 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v123 = *(*int32)(unsafe.Add(mBase, uint32(v121)))
	v124 = *(*int32)(unsafe.Add(mBase, uint32(v123)+8))
	v125 = F_objectGetVal(m, v124)
	mBase = m.M
	v126 = F_sdsnew(m, v125)
	mBase = m.M
	v127 = m.ExcPending
	if v127 != 0 {
		goto L1
	} else {
		goto L34
	}
L34:
	;
	v128 = F_dictAdd(m, v122, v126, v121)
	mBase = m.M
	v129 = m.ExcPending
	if v129 != 0 {
		goto L1
	} else {
		goto L35
	}
L35:
	;
	v132 = *(*int32)(unsafe.Add(mBase, uint32(v121+int32(-8))))
	goto L36
L36:
	;
	v137 = *(*int32)(unsafe.Add(mBase, uint32(v121)+4))
	v138 = *(*int32)(unsafe.Add(mBase, uint32(v137)+8))
	v139 = *(*int32)(unsafe.Add(mBase, uint32(v121)))
	v140 = F_scriptingEngineCallGetFunctionMemoryOverhead(m, v138, v139)
	mBase = m.M
	v141 = m.ExcPending
	if v141 != 0 {
		goto L1
	} else {
		goto L37
	}
L37:
	;
	v143 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v132&int32(2147483647) + int32(8) + v140 + v143
	goto L4
L38:
	;
	v148 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v149 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v150 = F_dictAdd(m, v148, v149, l1)
	mBase = m.M
	v151 = m.ExcPending
	if v151 != 0 {
		goto L1
	} else {
		goto L39
	}
L39:
	;
	v154 = *(*int32)(unsafe.Add(mBase, uint32(l1+int32(-8))))
	goto L40
L40:
	;
	v159 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v166 = v159 + int32(-1)
	v167 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v166))))
	v169 = v167 & int32(7)
	switch v169 {
	case 0:
		goto L47
	case 1:
		v175 = int32(4)
		goto L42
	case 2:
		goto L46
	case 3:
		goto L45
	case 4:
		goto L44
	default:
		goto L43
	}
L41:
	;
	v201 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v208 = v201 + int32(-1)
	v209 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v208))))
	v211 = v209 & int32(7)
	switch v211 {
	case 0:
		goto L60
	case 1:
		v217 = int32(4)
		goto L55
	case 2:
		goto L59
	case 3:
		goto L58
	case 4:
		goto L57
	default:
		goto L56
	}
L42:
	;
	switch v169 {
	case 0:
		goto L53
	case 1:
		goto L52
	case 2:
		goto L51
	case 3:
		goto L50
	case 4:
		goto L49
	default:
		v195 = int32(0)
		goto L48
	}
L43:
	;
	v175 = int32(1)
	goto L42
L44:
	;
	v175 = int32(18)
	goto L42
L45:
	;
	v175 = int32(10)
	goto L42
L46:
	;
	v175 = int32(6)
	goto L42
L47:
	;
	v170 = F_zmalloc_usable_size(m, v166)
	mBase = m.M
	v199 = v170
	goto L41
L48:
	;
	v199 = v175 + v195
	goto L41
L49:
	;
	v194 = *(*int32)(unsafe.Add(mBase, uint32(v159+int32(-9))))
	v195 = v194
	goto L48
L50:
	;
	v190 = *(*int32)(unsafe.Add(mBase, uint32(v159+int32(-5))))
	v199 = v175 + v190
	goto L41
L51:
	;
	v186 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v159+int32(-3)))))
	v199 = v175 + v186
	goto L41
L52:
	;
	v182 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v159+int32(-2)))))
	v199 = v175 + v182
	goto L41
L53:
	;
	v199 = v175 + int32(base.Ui32(v167)>>(uint(int32(3))%32))
	goto L41
L54:
	;
	v243 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v154&int32(2147483647) + int32(8) + v199 + v241 + v243
	v246 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v247 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v248 = *(*int32)(unsafe.Add(mBase, uint32(v247)))
	goto L68
L55:
	;
	switch v211 {
	case 0:
		goto L66
	case 1:
		goto L65
	case 2:
		goto L64
	case 3:
		goto L63
	case 4:
		goto L62
	default:
		v237 = int32(0)
		goto L61
	}
L56:
	;
	v217 = int32(1)
	goto L55
L57:
	;
	v217 = int32(18)
	goto L55
L58:
	;
	v217 = int32(10)
	goto L55
L59:
	;
	v217 = int32(6)
	goto L55
L60:
	;
	v212 = F_zmalloc_usable_size(m, v208)
	mBase = m.M
	v241 = v212
	goto L54
L61:
	;
	v241 = v217 + v237
	goto L54
L62:
	;
	v236 = *(*int32)(unsafe.Add(mBase, uint32(v201+int32(-9))))
	v237 = v236
	goto L61
L63:
	;
	v232 = *(*int32)(unsafe.Add(mBase, uint32(v201+int32(-5))))
	v241 = v217 + v232
	goto L54
L64:
	;
	v228 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v201+int32(-3)))))
	v241 = v217 + v228
	goto L54
L65:
	;
	v224 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v201+int32(-2)))))
	v241 = v217 + v224
	goto L54
L66:
	;
	v241 = v217 + int32(base.Ui32(v209)>>(uint(int32(3))%32))
	goto L54
L67:
	;
	v257 = *(*int32)(unsafe.Add(mBase, uint32(v249)))
	*(*int32)(unsafe.Add(mBase, uint32(v249))) = v257 + int32(1)
	v261 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v262 = *(*int32)(unsafe.Add(mBase, uint32(v261)+16))
	v263 = *(*int32)(unsafe.Add(mBase, uint32(v261)+12))
	v265 = *(*int32)(unsafe.Add(mBase, uint32(v249)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v249)+4)) = v262 + v263 + v265
	return
L68:
	;
	v249 = F_dictFetchValue(m, v246, v248)
	mBase = m.M
	v250 = m.ExcPending
	if v250 != 0 {
		goto L1
	} else {
		goto L69
	}
L69:
	;
	if v249 != 0 {
		goto L67
	} else {
		goto L70
	}
L70:
	;
	F__serverAssert(m, int32(_a_F_libraryLink_0), int32(_a_F_libraryLink_1), int32(357))
	mBase = m.M
	v255 = m.ExcPending
	if v255 != 0 {
		goto L1
	} else {
		goto L71
	}
L71:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_lindexCommand(m *base.Module, l0 int32) {
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
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
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
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v71 int64
	_ = v71
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v77 int64
	_ = v77
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	v6 = m.G0
	v8 = v6 - int32(64)
	m.G0 = v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(v10)+4))
	v13 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+224)))
	v17 = *(*int32)(unsafe.Add(mBase, uint32(v13<<(uint(int32(2))%32))+uint32(_c_F_lindexCommand[0])))
	v18 = F_lookupKeyReadOrReply(m, l0, v11, v17)
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		return
	} else {
		if v18 == int32(0) {
			m.G0 = v8 + int32(64)
			return
		} else {
			v23 = F_checkType(m, l0, v18, int32(1))
			mBase = m.M
			v24 = m.ExcPending
			if v24 != 0 {
				return
			} else {
				if v23 != 0 {
					m.G0 = v8 + int32(64)
					return
				} else {
					v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
					v26 = *(*int32)(unsafe.Add(mBase, uint32(v25)+8))
					v30 = F_getLongFromObjectOrReply(m, l0, v26, v8+int32(56), int32(0))
					mBase = m.M
					v31 = m.ExcPending
					if v31 != 0 {
						return
					} else {
						if v30 != 0 {
							m.G0 = v8 + int32(64)
							return
						} else {
							v32 = *(*int32)(unsafe.Add(mBase, uint32(v8)+56))
							v34 = F_listTypeInitIterator(m, v18, v32, int32(1))
							mBase = m.M
							v35 = m.ExcPending
							if v35 != 0 {
								return
							} else {
								v38 = F_listTypeNext(m, v34, v8+int32(16))
								mBase = m.M
								v39 = m.ExcPending
								if v39 != 0 {
									return
								} else {
									if v38 == int32(0) {
										F_addReplyNull(m, l0)
										mBase = m.M
										v76 = m.ExcPending
										if v76 != 0 {
											return
										} else {
											v82 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v34)+4)))
											if v82 != int32(9) {
												F_valkey_free(m, v34)
												mBase = m.M
												v89 = m.ExcPending
												if v89 != 0 {
													return
												} else {
													m.G0 = v8 + int32(64)
													return
												}
											} else {
												v85 = *(*int32)(unsafe.Add(mBase, uint32(v34)+12))
												F_quicklistReleaseIterator(m, v85)
												mBase = m.M
												v87 = m.ExcPending
												if v87 != 0 {
													return
												} else {
													F_valkey_free(m, v34)
													mBase = m.M
													v89 = m.ExcPending
													if v89 != 0 {
														return
													} else {
														m.G0 = v8 + int32(64)
														return
													}
												}
											}
										}
									} else {
										v42 = *(*int32)(unsafe.Add(mBase, uint32(v8)+16))
										v43 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v42)+4)))
										switch v43 + int32(-9) {
										case 0:
											v46 = *(*int32)(unsafe.Add(mBase, uint32(v8)+36))
											if v46 == int32(0) {
												v71 = *(*int64)(unsafe.Add(mBase, uint32(v8)+40))
												*(*int64)(unsafe.Add(mBase, uint32(v8)+8)) = v71
												F_addReplyBulkLongLong(m, l0, v71)
												mBase = m.M
												v74 = m.ExcPending
												if v74 != 0 {
													return
												} else {
													v82 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v34)+4)))
													if v82 != int32(9) {
														F_valkey_free(m, v34)
														mBase = m.M
														v89 = m.ExcPending
														if v89 != 0 {
															return
														} else {
															m.G0 = v8 + int32(64)
															return
														}
													} else {
														v85 = *(*int32)(unsafe.Add(mBase, uint32(v34)+12))
														F_quicklistReleaseIterator(m, v85)
														mBase = m.M
														v87 = m.ExcPending
														if v87 != 0 {
															return
														} else {
															F_valkey_free(m, v34)
															mBase = m.M
															v89 = m.ExcPending
															if v89 != 0 {
																return
															} else {
																m.G0 = v8 + int32(64)
																return
															}
														}
													}
												}
											} else {
												v49 = *(*int32)(unsafe.Add(mBase, uint32(v8)+48))
												F_addReplyBulkCBuffer(m, l0, v46, v49)
												mBase = m.M
												v51 = m.ExcPending
												if v51 != 0 {
													return
												} else {
													v82 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v34)+4)))
													if v82 != int32(9) {
														F_valkey_free(m, v34)
														mBase = m.M
														v89 = m.ExcPending
														if v89 != 0 {
															return
														} else {
															m.G0 = v8 + int32(64)
															return
														}
													} else {
														v85 = *(*int32)(unsafe.Add(mBase, uint32(v34)+12))
														F_quicklistReleaseIterator(m, v85)
														mBase = m.M
														v87 = m.ExcPending
														if v87 != 0 {
															return
														} else {
															F_valkey_free(m, v34)
															mBase = m.M
															v89 = m.ExcPending
															if v89 != 0 {
																return
															} else {
																m.G0 = v8 + int32(64)
																return
															}
														}
													}
												}
											}
										default:
											F__serverPanic_1(m, int32(_a_F_lindexCommand_0), int32(307), int32(_a_F_lindexCommand_1), int32(0))
											mBase = m.M
											v57 = m.ExcPending
											if v57 != 0 {
												return
											} else {
												F_abort(m)
												mBase = m.M
												base.Wasm_trap_unreachable()
												for {
												}
											}
										case 2:
											v59 = *(*int32)(unsafe.Add(mBase, uint32(v8)+20))
											v64 = F_lpGetValue(m, v59, v8+int32(60), v8+int32(8))
											mBase = m.M
											v65 = m.ExcPending
											if v65 != 0 {
												return
											} else {
												if v64 == int32(0) {
													v77 = *(*int64)(unsafe.Add(mBase, uint32(v8)+8))
													F_addReplyBulkLongLong(m, l0, v77)
													mBase = m.M
													v79 = m.ExcPending
													if v79 != 0 {
														return
													} else {
														v82 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v34)+4)))
														if v82 != int32(9) {
															F_valkey_free(m, v34)
															mBase = m.M
															v89 = m.ExcPending
															if v89 != 0 {
																return
															} else {
																m.G0 = v8 + int32(64)
																return
															}
														} else {
															v85 = *(*int32)(unsafe.Add(mBase, uint32(v34)+12))
															F_quicklistReleaseIterator(m, v85)
															mBase = m.M
															v87 = m.ExcPending
															if v87 != 0 {
																return
															} else {
																F_valkey_free(m, v34)
																mBase = m.M
																v89 = m.ExcPending
																if v89 != 0 {
																	return
																} else {
																	m.G0 = v8 + int32(64)
																	return
																}
															}
														}
													}
												} else {
													v68 = *(*int32)(unsafe.Add(mBase, uint32(v8)+60))
													F_addReplyBulkCBuffer(m, l0, v64, v68)
													mBase = m.M
													v70 = m.ExcPending
													if v70 != 0 {
														return
													} else {
														v82 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v34)+4)))
														if v82 != int32(9) {
															F_valkey_free(m, v34)
															mBase = m.M
															v89 = m.ExcPending
															if v89 != 0 {
																return
															} else {
																m.G0 = v8 + int32(64)
																return
															}
														} else {
															v85 = *(*int32)(unsafe.Add(mBase, uint32(v34)+12))
															F_quicklistReleaseIterator(m, v85)
															mBase = m.M
															v87 = m.ExcPending
															if v87 != 0 {
																return
															} else {
																F_valkey_free(m, v34)
																mBase = m.M
																v89 = m.ExcPending
																if v89 != 0 {
																	return
																} else {
																	m.G0 = v8 + int32(64)
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
			}
		}
	}
}
func F_ll2string(m *base.Module, l0 int32, l1 int32, l2 int64) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v12 int32
	_ = v12
	var v16 int32
	_ = v16
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int64
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v29 int32
	_ = v29
	var v37 int32
	_ = v37
	var v38 int64
	_ = v38
	var v59 int32
	_ = v59
	var v64 int32
	_ = v64
	var v71 int32
	_ = v71
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v82 int64
	_ = v82
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	var v96 int32
	_ = v96
	var v103 int64
	_ = v103
	var v105 int32
	_ = v105
	var v109 int64
	_ = v109
	var v110 int64
	_ = v110
	var v119 int32
	_ = v119
	var v122 int32
	_ = v122
	var v129 int32
	_ = v129
	var v130 int64
	_ = v130
	var v132 int32
	_ = v132
	var v137 int32
	_ = v137
	var v146 int32
	_ = v146
	var v150 int32
	_ = v150
	var v160 int32
	_ = v160
	var v167 int32
	_ = v167
	var v173 int32
	_ = v173
	v4 = int32(0)
	if l2 <= int64(-1) {
		goto L4
	} else {
		goto L5
	}
L1:
	;
	return v173
L2:
	;
	if l1 != int32(1) {
		v173 = v4
		goto L1
	} else {
		goto L52
	}
L3:
	;
	v25 = int32(0)
	v29 = int32(1)
	if base.Ui64(v23) < base.Ui64(int64(10)) {
		v86 = v29
		v87 = v25
		goto L8
	} else {
		goto L9
	}
L4:
	;
	if base.Ui32(l1) < base.Ui32(int32(2)) {
		goto L2
	} else {
		goto L6
	}
L5:
	;
	v21 = l0
	v22 = l1
	v23 = l2
	v24 = int32(0)
	goto L3
L6:
	;
	v12 = int32(45)
	*(*uint8)(unsafe.Add(mBase, uint32(l0))) = uint8(v12)
	v16 = int32(1)
	v21 = l0 + v16
	v22 = l1 + int32(-1)
	v23 = int64(0) - l2
	v24 = v16
	goto L3
L7:
	;
	if v160 == int32(0) {
		v173 = v4
		goto L1
	} else {
		goto L51
	}
L8:
	;
	v90 = v86 + v87
	if base.Ui32(v22) <= base.Ui32(v90) {
		goto L39
	} else {
		goto L40
	}
L9:
	;
	v37 = v25
	v38 = v23
	goto L10
L10:
	;
	if base.Ui64(int64(100)) <= base.Ui64(v38) {
		goto L12
	} else {
		goto L13
	}
L11:
	;
	v86 = v29
	v87 = v78
	goto L8
L12:
	;
	if base.Ui64(int64(1000)) <= base.Ui64(v38) {
		goto L14
	} else {
		goto L15
	}
L13:
	;
	v86 = int32(2)
	v87 = v37
	goto L8
L14:
	;
	if base.Ui64(int64(999999999999)) < base.Ui64(v38) {
		goto L16
	} else {
		goto L17
	}
L15:
	;
	v86 = int32(3)
	v87 = v37
	goto L8
L16:
	;
	v78 = v37 + int32(12)
	v82 = base.I64_div_u_s(v38, int64(1000000000000))
	if base.Ui64(int64(9999999999999)) < base.Ui64(v38) {
		v37 = v78
		v38 = v82
		goto L10
	} else {
		goto L38
	}
L17:
	;
	if base.Ui64(int64(99999999)) < base.Ui64(v38) {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	if base.Ui64(int64(9999999999)) < base.Ui64(v38) {
		goto L30
	} else {
		goto L31
	}
L19:
	;
	if base.Ui64(int64(999999)) < base.Ui64(v38) {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	if base.Ui64(int64(9999999)) < base.Ui64(v38) {
		goto L27
	} else {
		goto L28
	}
L21:
	;
	if base.Ui64(int64(10000)) <= base.Ui64(v38) {
		goto L22
	} else {
		goto L23
	}
L22:
	;
	if base.Ui64(int64(99999)) < base.Ui64(v38) {
		goto L24
	} else {
		goto L25
	}
L23:
	;
	v86 = int32(4)
	v87 = v37
	goto L8
L24:
	;
	v59 = int32(6)
	goto L26
L25:
	;
	v59 = int32(5)
	goto L26
L26:
	;
	v86 = v59
	v87 = v37
	goto L8
L27:
	;
	v64 = int32(8)
	goto L29
L28:
	;
	v64 = int32(7)
	goto L29
L29:
	;
	v86 = v64
	v87 = v37
	goto L8
L30:
	;
	if base.Ui64(int64(99999999999)) < base.Ui64(v38) {
		goto L35
	} else {
		goto L36
	}
L31:
	;
	if base.Ui64(int64(999999999)) < base.Ui64(v38) {
		goto L32
	} else {
		goto L33
	}
L32:
	;
	v71 = int32(10)
	goto L34
L33:
	;
	v71 = int32(9)
	goto L34
L34:
	;
	v86 = v71
	v87 = v37
	goto L8
L35:
	;
	v76 = int32(12)
	goto L37
L36:
	;
	v76 = int32(11)
	goto L37
L37:
	;
	v86 = v76
	v87 = v37
	goto L8
L38:
	;
	goto L11
L39:
	;
	if v22 == int32(0) {
		goto L49
	} else {
		goto L50
	}
L40:
	;
	v93 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v21+v90))) = uint8(v93)
	v96 = v90 + int32(-1)
	if base.Ui64(int64(100)) <= base.Ui64(v23) {
		goto L42
	} else {
		goto L43
	}
L41:
	;
	v132 = v21 + v129
	if base.Ui64(int64(9)) < base.Ui64(v130) {
		goto L47
	} else {
		goto L48
	}
L42:
	;
	v103 = v23
	v105 = v96
	goto L44
L43:
	;
	v129 = v96
	v130 = v23
	goto L41
L44:
	;
	v109 = int64(100)
	v110 = base.I64_div_u_s(v103, v109)
	v119 = int32(*(*uint16)(unsafe.Add(mBase, uint32(base.I32_wrap_i64(v103-v110*v109)<<(uint(int32(1))%32))+uint32(_c_F_ll2string[0]))))
	*(*uint16)(unsafe.Add(mBase, uint32(v21+int32(-1)+v105))) = uint16(v119)
	v122 = v105 + int32(-2)
	if base.Ui64(int64(9999)) < base.Ui64(v103) {
		v103 = v110
		v105 = v122
		goto L44
	} else {
		goto L46
	}
L45:
	;
	v129 = v122
	v130 = v110
	goto L41
L46:
	;
	goto L45
L47:
	;
	v146 = int32(*(*uint16)(unsafe.Add(mBase, uint32(base.I32_wrap_i64(v130)<<(uint(int32(1))%32))+uint32(_c_F_ll2string[0]))))
	*(*uint16)(unsafe.Add(mBase, uint32(v132+int32(-1)))) = uint16(v146)
	v160 = v90
	goto L7
L48:
	;
	v137 = base.I32_wrap_i64(v130) | int32(48)
	*(*uint8)(unsafe.Add(mBase, uint32(v132))) = uint8(v137)
	v160 = v90
	goto L7
L49:
	;
	v160 = int32(0)
	goto L7
L50:
	;
	v150 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v21))) = uint8(v150)
	goto L49
L51:
	;
	return v160 + v24
L52:
	;
	v167 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l0))) = uint8(v167)
	v173 = v167
	goto L1
}
func F_llenCommand(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v4 = *(*int32)(unsafe.Add(mBase, uint32(v3)+4))
	v6 = *(*int32)(unsafe.Add(mBase, _c_F_llenCommand[0]))
	v7 = F_lookupKeyReadOrReply(m, l0, v4, v6)
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		return
	} else {
		if v7 == int32(0) {
			return
		} else {
			v12 = F_checkType(m, l0, v7, int32(1))
			mBase = m.M
			v13 = m.ExcPending
			if v13 != 0 {
				return
			} else {
				if v12 != 0 {
					return
				} else {
					v14 = F_listTypeLength(m, v7)
					mBase = m.M
					v15 = m.ExcPending
					if v15 != 0 {
						return
					} else {
						F_addReplyLongLong(m, l0, base.I64_extend_i32_u(v14))
						mBase = m.M
						v18 = m.ExcPending
						if v18 != 0 {
							return
						} else {
							return
						}
					}
				}
			}
		}
	}
}
func F_lm_asprintf(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v13 int32
	_ = v13
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
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	v3 = int32(0)
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	*(*int32)(unsafe.Add(mBase, uint32(v8)+12)) = l1
	v13 = F_vsnprintf(m, v3, v3, l0, l1)
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		return int32(0)
	} else {
		v18 = v13 + int32(1)
		v19 = m.G22
		v20 = *(*int32)(unsafe.Add(mBase, uint32(v19)))
		v21 = m.T0[v20].(func(*base.Module, int32) int32)(m, v18)
		mBase = m.M
		v22 = m.ExcPending
		if v22 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v8)+12)) = l1
			v24 = F_vsnprintf(m, v21, v18, l0, l1)
			mBase = m.M
			v25 = m.ExcPending
			if v25 != 0 {
				return int32(0)
			} else {
				m.G0 = v8 + int32(16)
				return v21
			}
		}
	}
}
func F_lm_strcpy(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v17 int32
	_ = v17
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	if l0&int32(3) == int32(0) {
		v24 = l0
		goto L3
	} else {
		goto L4
	}
L1:
	;
	v59 = v57 + int32(1)
	v60 = m.G22
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v60)))
	v62 = m.T0[v61].(func(*base.Module, int32) int32)(m, v59)
	mBase = m.M
	v65 = m.ExcPending
	if v65 != 0 {
		goto L17
	} else {
		goto L18
	}
L2:
	;
	v57 = v49 - l0
	goto L1
L3:
	;
	v28 = v24
	goto L11
L4:
	;
	v10 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v10 != 0 {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	v13 = l0
	goto L7
L6:
	;
	v57 = l0 - l0
	goto L1
L7:
	;
	v17 = v13 + int32(1)
	if v17&int32(3) == int32(0) {
		v24 = v17
		goto L3
	} else {
		goto L9
	}
L9:
	;
	v22 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17))))
	if v22 != 0 {
		v13 = v17
		goto L7
	} else {
		goto L10
	}
L10:
	;
	v49 = v17
	goto L2
L11:
	;
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v28)))
	v37 = int32(-2139062144)
	if (int32(16843008)-v34|v34)&v37 == v37 {
		v28 = v28 + int32(4)
		goto L11
	} else {
		goto L13
	}
L12:
	;
	v43 = v28
	goto L14
L13:
	;
	goto L12
L14:
	;
	v47 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v43))))
	if v47 != 0 {
		v43 = v43 + int32(1)
		goto L14
	} else {
		goto L16
	}
L15:
	;
	v49 = v43
	goto L2
L16:
	;
	goto L15
L17:
	;
	return int32(0)
L18:
	;
	if v59 == int32(0) {
		v69 = v62
		goto L20
	} else {
		goto L21
	}
L19:
	;
	return v69
L20:
	;
	goto L19
L21:
	;
	v68 = F__emscripten_memcpy_bulkmem(m, v62, l0, v59)
	mBase = m.M
	v69 = v68
	goto L20
}
func F_lmoveGenericCommand(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
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
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
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
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v59 int64
	_ = v59
	var v60 int64
	_ = v60
	var v62 int32
	_ = v62
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v72 int64
	_ = v72
	var v74 int32
	_ = v74
	var v81 int32
	_ = v81
	var v95 int32
	_ = v95
	v11 = m.G0
	v13 = v11 - int32(48)
	m.G0 = v13
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v16 = *(*int32)(unsafe.Add(mBase, uint32(v15)+4))
	v18 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+224)))
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v18<<(uint(int32(2))%32))+uint32(_c_F_lmoveGenericCommand[0])))
	v23 = F_lookupKeyWriteOrReply(m, l0, v16, v22)
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		return
	} else {
		if v23 == int32(0) {
			m.G0 = v13 + int32(48)
			return
		} else {
			v28 = F_checkType(m, l0, v23, int32(1))
			mBase = m.M
			v29 = m.ExcPending
			if v29 != 0 {
				return
			} else {
				if v28 != 0 {
					m.G0 = v13 + int32(48)
					return
				} else {
					v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
					v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
					v32 = *(*int32)(unsafe.Add(mBase, uint32(v31)+8))
					v33 = F_lookupKeyWrite(m, v30, v32)
					mBase = m.M
					v34 = m.ExcPending
					if v34 != 0 {
						return
					} else {
						v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
						v36 = *(*int32)(unsafe.Add(mBase, uint32(v35)+4))
						v38 = F_checkType(m, l0, v33, int32(1))
						mBase = m.M
						v39 = m.ExcPending
						if v39 != 0 {
							return
						} else {
							if v38 != 0 {
								m.G0 = v13 + int32(48)
								return
							} else {
								v40 = F_listTypePop(m, v23, l1)
								mBase = m.M
								v41 = m.ExcPending
								if v41 != 0 {
									return
								} else {
									if v40 == int32(0) {
										F__serverAssert(m, int32(_a_F_lmoveGenericCommand_0), int32(_a_F_lmoveGenericCommand_1), int32(1117))
										mBase = m.M
										v95 = m.ExcPending
										if v95 != 0 {
											return
										} else {
											F_abort(m)
											mBase = m.M
											base.Wasm_trap_unreachable()
											for {
											}
										}
									} else {
										v44 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
										v45 = *(*int32)(unsafe.Add(mBase, uint32(v44)+8))
										F_lmoveHandlePush(m, l0, v45, v33, v40, l2)
										mBase = m.M
										v47 = m.ExcPending
										if v47 != 0 {
											return
										} else {
											F_listElementsRemoved(m, l0, v36, l1, v23, int32(1), int32(0))
											mBase = m.M
											v51 = m.ExcPending
											if v51 != 0 {
												return
											} else {
												F_decrRefCount(m, v40)
												mBase = m.M
												v53 = m.ExcPending
												if v53 != 0 {
													return
												} else {
													v54 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
													v55 = *(*int32)(unsafe.Add(mBase, uint32(v54)+48))
													if v55 != int32(258) {
														if v55 != int32(263) {
															m.G0 = v13 + int32(48)
															return
														} else {
															v71 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
															v72 = *(*int64)(unsafe.Add(mBase, uint32(v71)+4))
															v74 = *(*int32)(unsafe.Add(mBase, _c_F_lmoveGenericCommand[1]))
															*(*int32)(unsafe.Add(mBase, uint32(v13)+32)) = v74
															*(*int64)(unsafe.Add(mBase, uint32(v13)+36)) = v72
															F_rewriteClientCommandVector(m, l0, int32(3), v13+int32(32))
															mBase = m.M
															v81 = m.ExcPending
															if v81 != 0 {
																return
															} else {
																m.G0 = v13 + int32(48)
																return
															}
														}
													} else {
														v58 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
														v59 = *(*int64)(unsafe.Add(mBase, uint32(v58)+12))
														v60 = *(*int64)(unsafe.Add(mBase, uint32(v58)+4))
														v62 = *(*int32)(unsafe.Add(mBase, _c_F_lmoveGenericCommand[2]))
														*(*int32)(unsafe.Add(mBase, uint32(v13))) = v62
														*(*int64)(unsafe.Add(mBase, uint32(v13)+4)) = v60
														*(*int64)(unsafe.Add(mBase, uint32(v13)+12)) = v59
														F_rewriteClientCommandVector(m, l0, int32(5), v13)
														mBase = m.M
														v68 = m.ExcPending
														if v68 != 0 {
															return
														} else {
															m.G0 = v13 + int32(48)
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
	}
}
func F_lmpopCommand(m *base.Module, l0 int32) {
	var v5 int32
	_ = v5
	F_lmpopGenericCommand(m, l0, int32(1), int32(0))
	v5 = m.ExcPending
	if v5 != 0 {
		return
	} else {
		return
	}
}
func F_loadingAbsProgress(m *base.Module, l0 int64) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v15 int32
	_ = v15
	var v19 int32
	_ = v19
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v101 int32
	_ = v101
	var v110 int32
	_ = v110
	var v114 int32
	_ = v114
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v121 int32
	_ = v121
	var v128 int32
	_ = v128
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v135 int32
	_ = v135
	var v138 int32
	_ = v138
	var v141 int32
	_ = v141
	var v144 int32
	_ = v144
	var v147 int32
	_ = v147
	var v150 int32
	_ = v150
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v158 int32
	_ = v158
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v172 int32
	_ = v172
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v185 int32
	_ = v185
	var v187 int32
	_ = v187
	v2 = int32(_a_F_loadingAbsProgress_0)
	*(*int64)(unsafe.Add(mBase, _c_F_loadingAbsProgress[0])) = l0
	v5 = *(*int32)(unsafe.Add(mBase, _c_F_loadingAbsProgress[1]))
	v6 = int32(0)
	v15 = *(*int32)(unsafe.Add(mBase, _c_F_loadingAbsProgress[2]))
	if v15 < int32(261) {
		if v15 < int32(1) {
			v92 = v6
		} else {
			v23 = v6
			v24 = v15
			v26 = v24 & int32(3)
			if base.Ui32(int32(4)) <= base.Ui32(v24) {
				v33 = int32(0)
				v35 = v23
				v36 = v33
				v40 = v33
				for {
					v43 = v36 << (uint(int32(2)) % 32)
					v46 = *(*int32)(unsafe.Add(mBase, uint32(v43)+uint32(_c_F_loadingAbsProgress[3])))
					v49 = *(*int32)(unsafe.Add(mBase, uint32(v43)+uint32(_c_F_loadingAbsProgress[4])))
					v52 = *(*int32)(unsafe.Add(mBase, uint32(v43)+uint32(_c_F_loadingAbsProgress[5])))
					v55 = *(*int32)(unsafe.Add(mBase, uint32(v43)+uint32(_c_F_loadingAbsProgress[6])))
					v59 = v46 + (v49 + (v52 + (v55 + v35)))
					v60 = int32(4)
					v61 = v36 + v60
					v63 = v40 + v60
					if v63 != v24&int32(2147483644) {
						v35 = v59
						v36 = v61
						v40 = v63
						continue
					} else {
						break
					}
					break
				}
				v65 = v59
				v66 = v61
			} else {
				v65 = v23
				v66 = int32(0)
			}
			if v26 == int32(0) {
				v92 = v65
			} else {
				v74 = v65
				v75 = v66
				v77 = int32(0)
				for {
					v85 = *(*int32)(unsafe.Add(mBase, uint32(v75<<(uint(int32(2))%32))+uint32(_c_F_loadingAbsProgress[6])))
					v86 = v85 + v74
					v87 = int32(1)
					v90 = v77 + v87
					if v90 != v26 {
						v74 = v86
						v75 = v75 + v87
						v77 = v90
						continue
					} else {
						break
					}
					break
				}
				v92 = v86
			}
		}
	} else {
		v19 = *(*int32)(unsafe.Add(mBase, _c_F_loadingAbsProgress[7]))
		v23 = v19
		v24 = int32(260)
		v26 = v24 & int32(3)
		if base.Ui32(int32(4)) <= base.Ui32(v24) {
			v33 = int32(0)
			v35 = v23
			v36 = v33
			v40 = v33
			for {
				v43 = v36 << (uint(int32(2)) % 32)
				v46 = *(*int32)(unsafe.Add(mBase, uint32(v43)+uint32(_c_F_loadingAbsProgress[3])))
				v49 = *(*int32)(unsafe.Add(mBase, uint32(v43)+uint32(_c_F_loadingAbsProgress[4])))
				v52 = *(*int32)(unsafe.Add(mBase, uint32(v43)+uint32(_c_F_loadingAbsProgress[5])))
				v55 = *(*int32)(unsafe.Add(mBase, uint32(v43)+uint32(_c_F_loadingAbsProgress[6])))
				v59 = v46 + (v49 + (v52 + (v55 + v35)))
				v60 = int32(4)
				v61 = v36 + v60
				v63 = v40 + v60
				if v63 != v24&int32(2147483644) {
					v35 = v59
					v36 = v61
					v40 = v63
					continue
				} else {
					break
				}
				break
			}
			v65 = v59
			v66 = v61
		} else {
			v65 = v23
			v66 = int32(0)
		}
		if v26 == int32(0) {
			v92 = v65
		} else {
			v74 = v65
			v75 = v66
			v77 = int32(0)
			for {
				v85 = *(*int32)(unsafe.Add(mBase, uint32(v75<<(uint(int32(2))%32))+uint32(_c_F_loadingAbsProgress[6])))
				v86 = v85 + v74
				v87 = int32(1)
				v90 = v77 + v87
				if v90 != v26 {
					v74 = v86
					v75 = v75 + v87
					v77 = v90
					continue
				} else {
					break
				}
				break
			}
			v92 = v86
		}
	}
	if base.Ui32(v92) <= base.Ui32(v5) {
	} else {
		v101 = int32(0)
		v110 = *(*int32)(unsafe.Add(mBase, _c_F_loadingAbsProgress[2]))
		if v110 < int32(261) {
			if v110 < int32(1) {
				v187 = v101
			} else {
				v118 = v101
				v119 = v110
				v121 = v119 & int32(3)
				if base.Ui32(int32(4)) <= base.Ui32(v119) {
					v128 = int32(0)
					v130 = v118
					v131 = v128
					v135 = v128
					for {
						v138 = v131 << (uint(int32(2)) % 32)
						v141 = *(*int32)(unsafe.Add(mBase, uint32(v138)+uint32(_c_F_loadingAbsProgress[3])))
						v144 = *(*int32)(unsafe.Add(mBase, uint32(v138)+uint32(_c_F_loadingAbsProgress[4])))
						v147 = *(*int32)(unsafe.Add(mBase, uint32(v138)+uint32(_c_F_loadingAbsProgress[5])))
						v150 = *(*int32)(unsafe.Add(mBase, uint32(v138)+uint32(_c_F_loadingAbsProgress[6])))
						v154 = v141 + (v144 + (v147 + (v150 + v130)))
						v155 = int32(4)
						v156 = v131 + v155
						v158 = v135 + v155
						if v158 != v119&int32(2147483644) {
							v130 = v154
							v131 = v156
							v135 = v158
							continue
						} else {
							break
						}
						break
					}
					v160 = v154
					v161 = v156
				} else {
					v160 = v118
					v161 = int32(0)
				}
				if v121 == int32(0) {
					v187 = v160
				} else {
					v169 = v160
					v170 = v161
					v172 = int32(0)
					for {
						v180 = *(*int32)(unsafe.Add(mBase, uint32(v170<<(uint(int32(2))%32))+uint32(_c_F_loadingAbsProgress[6])))
						v181 = v180 + v169
						v182 = int32(1)
						v185 = v172 + v182
						if v185 != v121 {
							v169 = v181
							v170 = v170 + v182
							v172 = v185
							continue
						} else {
							break
						}
						break
					}
					v187 = v181
				}
			}
		} else {
			v114 = *(*int32)(unsafe.Add(mBase, _c_F_loadingAbsProgress[7]))
			v118 = v114
			v119 = int32(260)
			v121 = v119 & int32(3)
			if base.Ui32(int32(4)) <= base.Ui32(v119) {
				v128 = int32(0)
				v130 = v118
				v131 = v128
				v135 = v128
				for {
					v138 = v131 << (uint(int32(2)) % 32)
					v141 = *(*int32)(unsafe.Add(mBase, uint32(v138)+uint32(_c_F_loadingAbsProgress[3])))
					v144 = *(*int32)(unsafe.Add(mBase, uint32(v138)+uint32(_c_F_loadingAbsProgress[4])))
					v147 = *(*int32)(unsafe.Add(mBase, uint32(v138)+uint32(_c_F_loadingAbsProgress[5])))
					v150 = *(*int32)(unsafe.Add(mBase, uint32(v138)+uint32(_c_F_loadingAbsProgress[6])))
					v154 = v141 + (v144 + (v147 + (v150 + v130)))
					v155 = int32(4)
					v156 = v131 + v155
					v158 = v135 + v155
					if v158 != v119&int32(2147483644) {
						v130 = v154
						v131 = v156
						v135 = v158
						continue
					} else {
						break
					}
					break
				}
				v160 = v154
				v161 = v156
			} else {
				v160 = v118
				v161 = int32(0)
			}
			if v121 == int32(0) {
				v187 = v160
			} else {
				v169 = v160
				v170 = v161
				v172 = int32(0)
				for {
					v180 = *(*int32)(unsafe.Add(mBase, uint32(v170<<(uint(int32(2))%32))+uint32(_c_F_loadingAbsProgress[6])))
					v181 = v180 + v169
					v182 = int32(1)
					v185 = v172 + v182
					if v185 != v121 {
						v169 = v181
						v170 = v170 + v182
						v172 = v185
						continue
					} else {
						break
					}
					break
				}
				v187 = v181
			}
		}
		*(*int32)(unsafe.Add(mBase, _c_F_loadingAbsProgress[1])) = v187
	}
	return
}
func F_localeconv(m *base.Module) int32 {
	return int32(_a_F_localeconv_0)
}
func F_log10(m *base.Module, l0 float64) float64 {
	var v12 int64
	_ = v12
	var v29 int32
	_ = v29
	var v31 int64
	_ = v31
	var v41 int64
	_ = v41
	var v46 int64
	_ = v46
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v55 float64
	_ = v55
	var v57 float64
	_ = v57
	var v70 float64
	_ = v70
	var v73 float64
	_ = v73
	var v78 float64
	_ = v78
	var v79 float64
	_ = v79
	var v80 float64
	_ = v80
	var v81 float64
	_ = v81
	var v86 float64
	_ = v86
	var v87 float64
	_ = v87
	var v88 float64
	_ = v88
	var v113 float64
	_ = v113
	var v125 float64
	_ = v125
	v12 = base.I64_reinterpret_f64(l0)
	if int64(4503599627370495) < v12 {
		if base.Ui64(int64(9218868437227405311)) < base.Ui64(v12) {
			v125 = l0
			return v125
		} else {
			v29 = int32(-1023)
			v31 = int64(base.Ui64(v12) >> (uint(int64(32)) % 64))
			if v31 == int64(1072693248) {
				if base.I32_wrap_i64(v12) != 0 {
					v46 = v12
					v47 = v29
					v49 = int32(1072693248)
					v51 = v49 + int32(614242)
					v55 = base.F64_convert_i32_s(v47 + int32(base.Ui32(v51)>>(uint(int32(20))%32)))
					v57 = base.F64_mul(v55, float64(0.30102999566361177))
					v70 = base.F64_add(base.F64_reinterpret_i64(base.I64_extend_i32_u(v51&int32(1048575)+int32(1072079006))<<(uint(int64(32))%64)|v46&int64(4294967295)), float64(-1))
					v73 = base.F64_mul(v70, base.F64_mul(v70, float64(0.5)))
					v78 = base.F64_reinterpret_i64(base.I64_reinterpret_f64(base.F64_sub(v70, v73)) & int64(-4294967296))
					v79 = float64(0.4342944818781689)
					v80 = base.F64_mul(v78, v79)
					v81 = base.F64_add(v57, v80)
					v86 = base.F64_div(v70, base.F64_add(v70, float64(2)))
					v87 = base.F64_mul(v86, v86)
					v88 = base.F64_mul(v87, v87)
					v113 = base.F64_add(base.F64_mul(v86, base.F64_add(v73, base.F64_add(base.F64_mul(v88, base.F64_add(base.F64_mul(v88, base.F64_add(base.F64_mul(v88, float64(0.15313837699209373)), float64(0.22222198432149784))), float64(0.3999999999940942))), base.F64_mul(v87, base.F64_add(base.F64_mul(v88, base.F64_add(base.F64_mul(v88, base.F64_add(base.F64_mul(v88, float64(0.14798198605116586)), float64(0.1818357216161805))), float64(0.2857142874366239))), float64(0.6666666666666735)))))), base.F64_sub(base.F64_sub(v70, v78), v73))
					v125 = base.F64_add(v81, base.F64_add(base.F64_add(v80, base.F64_sub(v57, v81)), base.F64_add(base.F64_mul(v113, v79), base.F64_add(base.F64_mul(v55, float64(3.694239077158931e-13)), base.F64_mul(base.F64_add(v113, v78), float64(2.5082946711645275e-11))))))
					return v125
				} else {
					return float64(0)
				}
			} else {
				v46 = v12
				v47 = v29
				v49 = base.I32_wrap_i64(v31)
				v51 = v49 + int32(614242)
				v55 = base.F64_convert_i32_s(v47 + int32(base.Ui32(v51)>>(uint(int32(20))%32)))
				v57 = base.F64_mul(v55, float64(0.30102999566361177))
				v70 = base.F64_add(base.F64_reinterpret_i64(base.I64_extend_i32_u(v51&int32(1048575)+int32(1072079006))<<(uint(int64(32))%64)|v46&int64(4294967295)), float64(-1))
				v73 = base.F64_mul(v70, base.F64_mul(v70, float64(0.5)))
				v78 = base.F64_reinterpret_i64(base.I64_reinterpret_f64(base.F64_sub(v70, v73)) & int64(-4294967296))
				v79 = float64(0.4342944818781689)
				v80 = base.F64_mul(v78, v79)
				v81 = base.F64_add(v57, v80)
				v86 = base.F64_div(v70, base.F64_add(v70, float64(2)))
				v87 = base.F64_mul(v86, v86)
				v88 = base.F64_mul(v87, v87)
				v113 = base.F64_add(base.F64_mul(v86, base.F64_add(v73, base.F64_add(base.F64_mul(v88, base.F64_add(base.F64_mul(v88, base.F64_add(base.F64_mul(v88, float64(0.15313837699209373)), float64(0.22222198432149784))), float64(0.3999999999940942))), base.F64_mul(v87, base.F64_add(base.F64_mul(v88, base.F64_add(base.F64_mul(v88, base.F64_add(base.F64_mul(v88, float64(0.14798198605116586)), float64(0.1818357216161805))), float64(0.2857142874366239))), float64(0.6666666666666735)))))), base.F64_sub(base.F64_sub(v70, v78), v73))
				v125 = base.F64_add(v81, base.F64_add(base.F64_add(v80, base.F64_sub(v57, v81)), base.F64_add(base.F64_mul(v113, v79), base.F64_add(base.F64_mul(v55, float64(3.694239077158931e-13)), base.F64_mul(base.F64_add(v113, v78), float64(2.5082946711645275e-11))))))
				return v125
			}
		}
	} else {
		if base.F64_ne(l0, float64(0)) != 0 {
			if int64(-1) < v12 {
				v41 = base.I64_reinterpret_f64(base.F64_mul(l0, float64(1.8014398509481984e+16)))
				v46 = v41
				v47 = int32(-1077)
				v49 = base.I32_wrap_i64(int64(base.Ui64(v41) >> (uint(int64(32)) % 64)))
				v51 = v49 + int32(614242)
				v55 = base.F64_convert_i32_s(v47 + int32(base.Ui32(v51)>>(uint(int32(20))%32)))
				v57 = base.F64_mul(v55, float64(0.30102999566361177))
				v70 = base.F64_add(base.F64_reinterpret_i64(base.I64_extend_i32_u(v51&int32(1048575)+int32(1072079006))<<(uint(int64(32))%64)|v46&int64(4294967295)), float64(-1))
				v73 = base.F64_mul(v70, base.F64_mul(v70, float64(0.5)))
				v78 = base.F64_reinterpret_i64(base.I64_reinterpret_f64(base.F64_sub(v70, v73)) & int64(-4294967296))
				v79 = float64(0.4342944818781689)
				v80 = base.F64_mul(v78, v79)
				v81 = base.F64_add(v57, v80)
				v86 = base.F64_div(v70, base.F64_add(v70, float64(2)))
				v87 = base.F64_mul(v86, v86)
				v88 = base.F64_mul(v87, v87)
				v113 = base.F64_add(base.F64_mul(v86, base.F64_add(v73, base.F64_add(base.F64_mul(v88, base.F64_add(base.F64_mul(v88, base.F64_add(base.F64_mul(v88, float64(0.15313837699209373)), float64(0.22222198432149784))), float64(0.3999999999940942))), base.F64_mul(v87, base.F64_add(base.F64_mul(v88, base.F64_add(base.F64_mul(v88, base.F64_add(base.F64_mul(v88, float64(0.14798198605116586)), float64(0.1818357216161805))), float64(0.2857142874366239))), float64(0.6666666666666735)))))), base.F64_sub(base.F64_sub(v70, v78), v73))
				v125 = base.F64_add(v81, base.F64_add(base.F64_add(v80, base.F64_sub(v57, v81)), base.F64_add(base.F64_mul(v113, v79), base.F64_add(base.F64_mul(v55, float64(3.694239077158931e-13)), base.F64_mul(base.F64_add(v113, v78), float64(2.5082946711645275e-11))))))
				return v125
			} else {
				return base.F64_div(base.F64_sub(l0, l0), float64(0))
			}
		} else {
			return base.F64_div(float64(-1), base.F64_mul(l0, l0))
		}
	}
}
func F_lolwut6Command(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
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
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
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
	var v103 int32
	_ = v103
	var v105 int32
	_ = v105
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v124 int32
	_ = v124
	var v128 int32
	_ = v128
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v133 int32
	_ = v133
	var v135 int32
	_ = v135
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v144 int32
	_ = v144
	var v146 int32
	_ = v146
	var v148 int32
	_ = v148
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v161 int32
	_ = v161
	var v167 int32
	_ = v167
	var v171 int32
	_ = v171
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v176 int32
	_ = v176
	var v178 int32
	_ = v178
	var v180 int32
	_ = v180
	var v187 int32
	_ = v187
	var v189 int32
	_ = v189
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v196 int32
	_ = v196
	var v198 int32
	_ = v198
	var v200 int32
	_ = v200
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
	var v240 int32
	_ = v240
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v253 int32
	_ = v253
	var v255 int32
	_ = v255
	var v259 int32
	_ = v259
	var v264 int32
	_ = v264
	var v267 int32
	_ = v267
	var v269 int32
	_ = v269
	var v271 int32
	_ = v271
	var v275 int32
	_ = v275
	var v276 int32
	_ = v276
	var v279 int32
	_ = v279
	var v286 int32
	_ = v286
	var v287 int32
	_ = v287
	var v292 int32
	_ = v292
	var v293 int32
	_ = v293
	var v295 int32
	_ = v295
	var v296 int32
	_ = v296
	var v298 int32
	_ = v298
	var v301 int32
	_ = v301
	var v304 int32
	_ = v304
	var v312 int32
	_ = v312
	var v313 int32
	_ = v313
	var v314 int32
	_ = v314
	var v315 int32
	_ = v315
	var v316 int32
	_ = v316
	var v318 int32
	_ = v318
	var v325 int32
	_ = v325
	var v331 int32
	_ = v331
	var v332 int32
	_ = v332
	var v336 int32
	_ = v336
	var v337 int32
	_ = v337
	var v341 int32
	_ = v341
	var v342 int32
	_ = v342
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
	var v358 int32
	_ = v358
	var v361 int32
	_ = v361
	var v364 int32
	_ = v364
	var v367 int32
	_ = v367
	var v368 int32
	_ = v368
	var v371 int32
	_ = v371
	var v373 int32
	_ = v373
	var v375 int32
	_ = v375
	v9 = m.G0
	v11 = v9 - int32(32)
	m.G0 = v11
	*(*int32)(unsafe.Add(mBase, uint32(v11)+28)) = int32(80)
	*(*int32)(unsafe.Add(mBase, uint32(v11)+24)) = int32(20)
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v17 < int32(2) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v11 + int32(32)
	return
L2:
	;
	v37 = int32(1)
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v11)+28))
	if v38 < v37 {
		v44 = v37
		goto L11
	} else {
		goto L12
	}
L3:
	;
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v21 = *(*int32)(unsafe.Add(mBase, uint32(v20)+4))
	v25 = F_getLongFromObjectOrReply(m, l0, v21, v11+int32(28), int32(0))
	mBase = m.M
	v26 = m.ExcPending
	if v26 != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	return
L5:
	;
	if v25 != 0 {
		goto L1
	} else {
		goto L6
	}
L6:
	;
	v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v27 < int32(3) {
		goto L2
	} else {
		goto L7
	}
L7:
	;
	v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v30)+8))
	v35 = F_getLongFromObjectOrReply(m, l0, v31, v11+int32(24), int32(0))
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
		goto L4
	} else {
		goto L8
	}
L8:
	;
	if v35 != 0 {
		goto L1
	} else {
		goto L9
	}
L9:
	;
	goto L2
L10:
	;
	v48 = int32(1)
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v11)+24))
	if v49 < v48 {
		v55 = v48
		goto L15
	} else {
		goto L16
	}
L11:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+28)) = v44
	v47 = v44
	goto L10
L12:
	;
	if base.Ui32(v38) < base.Ui32(int32(1001)) {
		v47 = v38
		goto L10
	} else {
		goto L13
	}
L13:
	;
	v44 = int32(1000)
	goto L11
L14:
	;
	v60 = F_lwCreateCanvas(m, v47, v57, int32(3))
	mBase = m.M
	v61 = m.ExcPending
	if v61 != 0 {
		goto L4
	} else {
		goto L18
	}
L15:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+24)) = v55
	v57 = v55
	goto L14
L16:
	;
	if base.Ui32(v49) < base.Ui32(int32(1001)) {
		v57 = v49
		goto L14
	} else {
		goto L17
	}
L17:
	;
	v55 = int32(1000)
	goto L15
L18:
	;
	v68 = m.G0
	v70 = v68 - int32(32)
	m.G0 = v70
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v60)))
	*(*int32)(unsafe.Add(mBase, uint32(v70)+28)) = int32(2)
	v75 = int32(-10)
	if v72 <= v75 {
		goto L22
	} else {
		goto L23
	}
L19:
	;
	v232 = F_sdsempty(m)
	mBase = m.M
	v233 = m.ExcPending
	if v233 != 0 {
		goto L4
	} else {
		goto L37
	}
L20:
	;
	m.G0 = v70 + int32(32)
	goto L19
L21:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v70)+28)) = int32(0)
	goto L20
L22:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v70)+28)) = int32(1)
	goto L21
L23:
	;
	v81 = v75
	goto L24
L24:
	;
	v85 = F_rand(m)
	mBase = m.M
	v87 = base.I32_rem_s(v85, int32(8))
	v88 = v87 + v81
	*(*int32)(unsafe.Add(mBase, uint32(v70)+12)) = v88
	v90 = F_rand(m)
	mBase = m.M
	v92 = base.I32_rem_s(v90, int32(9))
	v94 = v92 + int32(10)
	*(*int32)(unsafe.Add(mBase, uint32(v70)+16)) = v94
	v96 = *(*int32)(unsafe.Add(mBase, uint32(v60)+4))
	v97 = F_rand(m)
	mBase = m.M
	v98 = *(*int32)(unsafe.Add(mBase, uint32(v60)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v70)+24)) = int32(0)
	v101 = base.I32_rem_s(v97, v98)
	v102 = int32(2)
	v103 = base.I32_div_s(v101, v102)
	v105 = base.I32_div_s(v96, v102)
	*(*int32)(unsafe.Add(mBase, uint32(v70)+20)) = v103 + v105
	F_generateSkyscraper(m, v60, v70+int32(12))
	mBase = m.M
	v113 = int32(base.Ui32(v94)>>(uint(int32(1))%32)) + v88
	v114 = *(*int32)(unsafe.Add(mBase, uint32(v60)))
	if v113 < v114 {
		v81 = v113
		goto L24
	} else {
		goto L26
	}
L25:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v70)+28)) = int32(1)
	if v114 < int32(-9) {
		goto L21
	} else {
		goto L27
	}
L26:
	;
	goto L25
L27:
	;
	v124 = int32(-10)
	goto L28
L28:
	;
	v128 = F_rand(m)
	mBase = m.M
	v130 = base.I32_rem_s(v128, int32(8))
	v131 = v130 + v124
	*(*int32)(unsafe.Add(mBase, uint32(v70)+12)) = v131
	v133 = F_rand(m)
	mBase = m.M
	v135 = base.I32_rem_s(v133, int32(9))
	*(*int32)(unsafe.Add(mBase, uint32(v70)+16)) = v135 + int32(10)
	v139 = *(*int32)(unsafe.Add(mBase, uint32(v60)+4))
	v140 = F_rand(m)
	mBase = m.M
	v141 = *(*int32)(unsafe.Add(mBase, uint32(v60)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v70)+24)) = int32(0)
	v144 = base.I32_rem_s(v140, v141)
	v146 = base.I32_div_s(v144, int32(3))
	v148 = base.I32_div_s(v139, int32(2))
	*(*int32)(unsafe.Add(mBase, uint32(v70)+20)) = v146 + v148
	F_generateSkyscraper(m, v60, v70+int32(12))
	mBase = m.M
	v156 = v135 + v131 + int32(11)
	v157 = *(*int32)(unsafe.Add(mBase, uint32(v60)))
	if v156 < v157 {
		v124 = v156
		goto L28
	} else {
		goto L30
	}
L29:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v70)+28)) = int32(0)
	v161 = int32(-10)
	if v157 <= v161 {
		goto L20
	} else {
		goto L31
	}
L30:
	;
	goto L29
L31:
	;
	v167 = v161
	goto L32
L32:
	;
	v171 = F_rand(m)
	mBase = m.M
	v173 = base.I32_rem_s(v171, int32(8))
	v174 = v173 + v167
	*(*int32)(unsafe.Add(mBase, uint32(v70)+12)) = v174
	v176 = F_rand(m)
	mBase = m.M
	v178 = base.I32_rem_s(v176, int32(14))
	v180 = v178 + int32(5)
	if v180&int32(3) == int32(0) {
		v189 = v180
		goto L34
	} else {
		goto L35
	}
L34:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v70)+16)) = v189
	v191 = *(*int32)(unsafe.Add(mBase, uint32(v60)+4))
	v192 = F_rand(m)
	mBase = m.M
	v193 = *(*int32)(unsafe.Add(mBase, uint32(v60)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v70)+24)) = int32(1)
	v196 = base.I32_rem_s(v192, v193)
	v198 = base.I32_div_s(v196, int32(2))
	v200 = base.I32_div_s(v191, int32(3))
	*(*int32)(unsafe.Add(mBase, uint32(v70)+20)) = v198 + v200
	F_generateSkyscraper(m, v60, v70+int32(12))
	mBase = m.M
	v208 = v174 + v189 + int32(5)
	v209 = *(*int32)(unsafe.Add(mBase, uint32(v60)))
	if v208 < v209 {
		v167 = v208
		goto L32
	} else {
		goto L36
	}
L35:
	;
	v187 = base.I32_rem_s(base.I32_extend8_s(v180), int32(3))
	v189 = v180 + v187
	goto L34
L36:
	;
	goto L20
L37:
	;
	v234 = *(*int32)(unsafe.Add(mBase, uint32(v60)+4))
	if v234 < int32(1) {
		v325 = v232
		goto L38
	} else {
		goto L39
	}
L38:
	;
	v331 = *(*int32)(unsafe.Add(mBase, _c_F_lolwut6Command[0]))
	if v331 != 0 {
		goto L60
	} else {
		goto L61
	}
L39:
	;
	v240 = v234
	v243 = v232
	v244 = int32(0)
	goto L40
L40:
	;
	v246 = int32(0)
	v247 = *(*int32)(unsafe.Add(mBase, uint32(v60)))
	if v247 <= v246 {
		v301 = v240
		v304 = v243
		goto L42
	} else {
		goto L43
	}
L41:
	;
	v325 = v316
	goto L38
L42:
	;
	if v244 == v301+int32(-1) {
		v315 = v301
		v316 = v304
		goto L56
	} else {
		goto L57
	}
L43:
	;
	v253 = v246
	v255 = v243
	goto L44
L44:
	;
	v259 = int32(0)
	if v253 < v259 {
		v276 = v259
		goto L48
	} else {
		goto L49
	}
L45:
	;
	v298 = *(*int32)(unsafe.Add(mBase, uint32(v60)+4))
	v301 = v298
	v304 = v292
	goto L42
L46:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = v287
	v292 = F_sdscatprintf(m, v255, int32(_a_F_lolwut6Command_0), v11+int32(16))
	mBase = m.M
	v293 = m.ExcPending
	if v293 != 0 {
		goto L4
	} else {
		goto L54
	}
L47:
	;
	v279 = v276 + int32(-1)
	if base.Ui32(int32(2)) < base.Ui32(v279) {
		v287 = int32(_a_F_lolwut6Command_1)
		goto L46
	} else {
		goto L53
	}
L48:
	;
	goto L47
L49:
	;
	v264 = int32(0)
	if v244 < v264 {
		v276 = v264
		goto L48
	} else {
		goto L50
	}
L50:
	;
	v267 = *(*int32)(unsafe.Add(mBase, uint32(v60)))
	if v267 <= v253 {
		v276 = v264
		goto L48
	} else {
		goto L51
	}
L51:
	;
	v269 = *(*int32)(unsafe.Add(mBase, uint32(v60)+4))
	if v269 <= v244 {
		v276 = v264
		goto L48
	} else {
		goto L52
	}
L52:
	;
	v271 = *(*int32)(unsafe.Add(mBase, uint32(v60)+8))
	v275 = int32(*(*int8)(unsafe.Add(mBase, uint32(v271+v253+v267*v244))))
	v276 = v275
	goto L48
L53:
	;
	v286 = *(*int32)(unsafe.Add(mBase, uint32(v279<<(uint(int32(2))%32))+uint32(_c_F_lolwut6Command[1])))
	v287 = v286
	goto L46
L54:
	;
	v295 = v253 + int32(1)
	v296 = *(*int32)(unsafe.Add(mBase, uint32(v60)))
	if v295 < v296 {
		v253 = v295
		v255 = v292
		goto L44
	} else {
		goto L55
	}
L55:
	;
	goto L45
L56:
	;
	v318 = v244 + int32(1)
	if v318 < v315 {
		v240 = v315
		v243 = v316
		v244 = v318
		goto L40
	} else {
		goto L59
	}
L57:
	;
	v312 = F_sdscatlen(m, v304, int32(_a_F_lolwut6Command_2), int32(1))
	mBase = m.M
	v313 = m.ExcPending
	if v313 != 0 {
		goto L4
	} else {
		goto L58
	}
L58:
	;
	v314 = *(*int32)(unsafe.Add(mBase, uint32(v60)+4))
	v315 = v314
	v316 = v312
	goto L56
L59:
	;
	goto L41
L60:
	;
	v332 = int32(_a_F_lolwut6Command_3)
	goto L62
L61:
	;
	v332 = int32(_a_F_lolwut6Command_4)
	goto L62
L62:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11))) = v332
	v336 = F_sdscatprintf(m, v325, int32(_a_F_lolwut6Command_5), v11)
	mBase = m.M
	v337 = m.ExcPending
	if v337 != 0 {
		goto L4
	} else {
		goto L69
	}
L63:
	;
	F_addReplyVerbatim(m, l0, v347, v368, int32(_a_F_lolwut6Command_6))
	mBase = m.M
	v371 = m.ExcPending
	if v371 != 0 {
		goto L4
	} else {
		goto L75
	}
L64:
	;
	v367 = *(*int32)(unsafe.Add(mBase, uint32(v347+int32(-17))))
	v368 = v367
	goto L63
L65:
	;
	v364 = *(*int32)(unsafe.Add(mBase, uint32(v347+int32(-9))))
	v368 = v364
	goto L63
L66:
	;
	v361 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v347+int32(-5)))))
	v368 = v361
	goto L63
L67:
	;
	v358 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v347+int32(-3)))))
	v368 = v358
	goto L63
L68:
	;
	v368 = int32(base.Ui32(v351) >> (uint(int32(3)) % 32))
	goto L63
L69:
	;
	v341 = *(*int32)(unsafe.Add(mBase, _c_F_lolwut6Command[0]))
	if v341 != 0 {
		goto L70
	} else {
		goto L71
	}
L70:
	;
	v342 = int32(_a_F_lolwut6Command_7)
	goto L72
L71:
	;
	v342 = int32(_a_F_lolwut6Command_8)
	goto L72
L72:
	;
	v343 = F_sdscat(m, v336, v342)
	mBase = m.M
	v344 = m.ExcPending
	if v344 != 0 {
		goto L4
	} else {
		goto L73
	}
L73:
	;
	v347 = F_sdscatlen(m, v343, int32(_a_F_lolwut6Command_2), int32(1))
	mBase = m.M
	v348 = m.ExcPending
	if v348 != 0 {
		goto L4
	} else {
		goto L74
	}
L74:
	;
	v351 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v347+int32(-1)))))
	switch v351 & int32(7) {
	case 0:
		goto L68
	case 1:
		goto L67
	case 2:
		goto L66
	case 3:
		goto L65
	case 4:
		goto L64
	default:
		v368 = int32(0)
		goto L63
	}
L75:
	;
	F_sdsfree(m, v347)
	mBase = m.M
	v373 = m.ExcPending
	if v373 != 0 {
		goto L4
	} else {
		goto L76
	}
L76:
	;
	F_lwFreeCanvas(m, v60)
	mBase = m.M
	v375 = m.ExcPending
	if v375 != 0 {
		goto L4
	} else {
		goto L77
	}
L77:
	;
	goto L1
}
func F_lolwutCommand(m *base.Module, l0 int32) {
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
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v46 int32
	_ = v46
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
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v73 int32
	_ = v73
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v85 int32
	_ = v85
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v97 int32
	_ = v97
	var v100 int32
	_ = v100
	var v103 int32
	_ = v103
	var v106 int32
	_ = v106
	var v109 int32
	_ = v109
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v117 int32
	_ = v117
	var v121 int32
	_ = v121
	var v123 int32
	_ = v123
	var v125 int32
	_ = v125
	var v129 int32
	_ = v129
	var v133 int32
	_ = v133
	v6 = m.G0
	v8 = v6 - int32(80)
	m.G0 = v8
	v10 = int32(_a_F_lolwutCommand_0)
	v11 = int32(_a_F_lolwutCommand_1)
	v12 = int32(_a_F_lolwutCommand_2)
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v13 < int32(3) {
		v91 = v10
		v92 = v11
		v93 = v12
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v8 + int32(80)
	return
L2:
	;
	v94 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v91))))
	switch v94 + int32(-52) {
	case 0:
		goto L27
	case 1:
		goto L28
	case 2:
		goto L25
	default:
		goto L22
	case 5:
		goto L23
	}
L3:
	;
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v17 = *(*int32)(unsafe.Add(mBase, uint32(v16)+4))
	v18 = F_objectGetVal(m, v17)
	mBase = m.M
	v19 = int32(_a_F_lolwutCommand_3)
	v22 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18))))
	if v22 != 0 {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	if v54-v56 != 0 {
		v91 = v10
		v92 = v11
		v93 = v12
		goto L2
	} else {
		goto L16
	}
L5:
	;
	v54 = F_tolower(m, v50)
	mBase = m.M
	v55 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v51))))
	v56 = F_tolower(m, v55)
	mBase = m.M
	goto L4
L6:
	;
	v24 = v18
	v25 = v19
	v26 = v22
	goto L9
L7:
	;
	v50 = int32(0)
	v51 = v19
	goto L5
L8:
	;
	v50 = v47 & int32(255)
	v51 = v46
	goto L5
L9:
	;
	v28 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v25))))
	if v28 == int32(0) {
		v46 = v25
		v47 = v26
		goto L8
	} else {
		goto L11
	}
L10:
	;
	v46 = v40
	v47 = int32(0)
	goto L8
L11:
	;
	v32 = v26 & int32(255)
	if v32 == v28 {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	v39 = int32(1)
	v40 = v25 + v39
	v41 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24)+1)))
	if v41 != 0 {
		v24 = v24 + v39
		v25 = v40
		v26 = v41
		goto L9
	} else {
		goto L15
	}
L13:
	;
	v34 = F_tolower(m, v32)
	mBase = m.M
	v35 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v25))))
	v36 = F_tolower(m, v35)
	mBase = m.M
	if v34 == v36 {
		goto L12
	} else {
		goto L14
	}
L14:
	;
	v38 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24))))
	v46 = v25
	v47 = v38
	goto L8
L15:
	;
	goto L10
L16:
	;
	v58 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v58)+8))
	v63 = F_getLongFromObjectOrReply(m, l0, v59, v8+int32(12), int32(0))
	mBase = m.M
	v64 = m.ExcPending
	if v64 != 0 {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	return
L18:
	;
	if v63 != 0 {
		goto L1
	} else {
		goto L19
	}
L19:
	;
	v66 = v8 + int32(16)
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v8)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v8))) = v73
	v79 = F_snprintf(m, v66, int32(64), int32(_a_F_lolwutCommand_4), v8)
	mBase = m.M
	v80 = m.ExcPending
	if v80 != 0 {
		goto L17
	} else {
		goto L20
	}
L20:
	;
	v81 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v81 + int32(8)
	v85 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = v85 + int32(-2)
	v91 = v8 + int32(16)
	v92 = v66 | int32(2)
	v93 = v66 | int32(1)
	goto L2
L21:
	;
	if v91 != v8+int32(16) {
		goto L1
	} else {
		goto L40
	}
L22:
	;
	F_lolwutUnstableCommand(m, l0)
	mBase = m.M
	v125 = m.ExcPending
	if v125 != 0 {
		goto L17
	} else {
		goto L39
	}
L23:
	;
	F_lolwut9Command(m, l0)
	mBase = m.M
	v123 = m.ExcPending
	if v123 != 0 {
		goto L17
	} else {
		goto L38
	}
L24:
	;
	F_lolwut6Command(m, l0)
	mBase = m.M
	v121 = m.ExcPending
	if v121 != 0 {
		goto L17
	} else {
		goto L37
	}
L25:
	;
	v114 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v93))))
	if v114 != int32(46) {
		goto L22
	} else {
		goto L35
	}
L26:
	;
	F_lolwut5Command(m, l0)
	mBase = m.M
	v113 = m.ExcPending
	if v113 != 0 {
		goto L17
	} else {
		goto L34
	}
L27:
	;
	v106 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v93))))
	if v106 != int32(46) {
		goto L22
	} else {
		goto L32
	}
L28:
	;
	v97 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v93))))
	if v97 != int32(46) {
		goto L22
	} else {
		goto L29
	}
L29:
	;
	v100 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v92))))
	if v100 != int32(57) {
		goto L26
	} else {
		goto L30
	}
L30:
	;
	v103 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v92))))
	if v103 == int32(57) {
		goto L24
	} else {
		goto L31
	}
L31:
	;
	goto L22
L32:
	;
	v109 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v92))))
	if v109 != int32(57) {
		goto L22
	} else {
		goto L33
	}
L33:
	;
	goto L26
L34:
	;
	goto L21
L35:
	;
	v117 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v92))))
	if v117 == int32(57) {
		goto L22
	} else {
		goto L36
	}
L36:
	;
	goto L24
L37:
	;
	goto L21
L38:
	;
	goto L21
L39:
	;
	goto L21
L40:
	;
	v129 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v129 + int32(-8)
	v133 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = v133 + int32(2)
	goto L1
}
func F_lposCommand(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v28 int32
	_ = v28
	var v38 int32
	_ = v38
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
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
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v96 int32
	_ = v96
	var v100 int32
	_ = v100
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
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v121 int32
	_ = v121
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
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v152 int32
	_ = v152
	var v156 int32
	_ = v156
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v164 int32
	_ = v164
	var v167 int32
	_ = v167
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v173 int32
	_ = v173
	var v177 int32
	_ = v177
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v204 int32
	_ = v204
	var v208 int32
	_ = v208
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v219 int32
	_ = v219
	var v222 int32
	_ = v222
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v249 int32
	_ = v249
	var v251 int32
	_ = v251
	var v253 int32
	_ = v253
	var v257 int32
	_ = v257
	var v259 int32
	_ = v259
	var v261 int32
	_ = v261
	var v262 int32
	_ = v262
	var v263 int32
	_ = v263
	var v264 int32
	_ = v264
	var v266 int32
	_ = v266
	var v269 int32
	_ = v269
	var v270 int32
	_ = v270
	var v271 int32
	_ = v271
	var v274 int32
	_ = v274
	var v275 int32
	_ = v275
	var v276 int32
	_ = v276
	var v277 int32
	_ = v277
	var v280 int32
	_ = v280
	var v281 int32
	_ = v281
	var v284 int32
	_ = v284
	var v292 int32
	_ = v292
	var v294 int32
	_ = v294
	var v296 int32
	_ = v296
	var v301 int32
	_ = v301
	var v307 int32
	_ = v307
	var v308 int32
	_ = v308
	var v312 int32
	_ = v312
	var v313 int32
	_ = v313
	var v318 int32
	_ = v318
	var v323 int32
	_ = v323
	var v325 int32
	_ = v325
	var v326 int32
	_ = v326
	var v327 int32
	_ = v327
	var v333 int32
	_ = v333
	var v335 int32
	_ = v335
	var v342 int32
	_ = v342
	var v343 int32
	_ = v343
	var v348 int32
	_ = v348
	var v360 int32
	_ = v360
	var v361 int32
	_ = v361
	var v366 int32
	_ = v366
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
	var v384 int32
	_ = v384
	var v388 int32
	_ = v388
	var v390 int32
	_ = v390
	var v392 int32
	_ = v392
	var v394 int32
	_ = v394
	v15 = m.G0
	v17 = v15 - int32(64)
	m.G0 = v17
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v20 = *(*int32)(unsafe.Add(mBase, uint32(v19)+8))
	v21 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v17)+60)) = v21
	*(*int32)(unsafe.Add(mBase, uint32(v17)+56)) = int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(v17)+52)) = int32(0)
	v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if int32(4) <= v28 {
		goto L4
	} else {
		goto L5
	}
L1:
	;
	m.G0 = v17 + int32(64)
	return
L2:
	;
	v392 = *(*int32)(unsafe.Add(mBase, _c_F_lposCommand[0]))
	F_addReplyErrorObject(m, l0, v392)
	mBase = m.M
	v394 = m.ExcPending
	if v394 != 0 {
		goto L24
	} else {
		goto L113
	}
L3:
	;
	v240 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v241 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v242 = *(*int32)(unsafe.Add(mBase, uint32(v241)+4))
	v243 = F_lookupKeyRead(m, v240, v242)
	mBase = m.M
	v244 = m.ExcPending
	if v244 != 0 {
		goto L24
	} else {
		goto L65
	}
L4:
	;
	v38 = int32(3)
	goto L6
L5:
	;
	v229 = v21
	v230 = int32(1)
	goto L3
L6:
	;
	v48 = v38 + int32(1)
	v49 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v49+v38<<(uint(int32(2))%32))))
	v54 = F_objectGetVal(m, v53)
	mBase = m.M
	v55 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v56 = int32(_a_F_lposCommand_0)
	v59 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v54))))
	if v59 != 0 {
		goto L12
	} else {
		goto L13
	}
L7:
	;
	v219 = *(*int32)(unsafe.Add(mBase, uint32(v17)+60))
	if int32(-1) < v219 {
		v229 = int32(1)
		v230 = v219
		goto L3
	} else {
		goto L63
	}
L8:
	;
	v215 = v38 + int32(2)
	v216 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v215 < v216 {
		v38 = v215
		goto L6
	} else {
		goto L62
	}
L9:
	;
	v112 = int32(_a_F_lposCommand_1)
	v115 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v54))))
	if v115 != 0 {
		goto L32
	} else {
		goto L33
	}
L10:
	;
	if v91-v93 != 0 {
		goto L9
	} else {
		goto L22
	}
L11:
	;
	v91 = F_tolower(m, v87)
	mBase = m.M
	v92 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v88))))
	v93 = F_tolower(m, v92)
	mBase = m.M
	goto L10
L12:
	;
	v61 = v54
	v62 = v56
	v63 = v59
	goto L15
L13:
	;
	v87 = int32(0)
	v88 = v56
	goto L11
L14:
	;
	v87 = v84 & int32(255)
	v88 = v83
	goto L11
L15:
	;
	v65 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v62))))
	if v65 == int32(0) {
		v83 = v62
		v84 = v63
		goto L14
	} else {
		goto L17
	}
L16:
	;
	v83 = v77
	v84 = int32(0)
	goto L14
L17:
	;
	v69 = v63 & int32(255)
	if v69 == v65 {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	v76 = int32(1)
	v77 = v62 + v76
	v78 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v61)+1)))
	if v78 != 0 {
		v61 = v61 + v76
		v62 = v77
		v63 = v78
		goto L15
	} else {
		goto L21
	}
L19:
	;
	v71 = F_tolower(m, v69)
	mBase = m.M
	v72 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v62))))
	v73 = F_tolower(m, v72)
	mBase = m.M
	if v71 == v73 {
		goto L18
	} else {
		goto L20
	}
L20:
	;
	v75 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v61))))
	v83 = v62
	v84 = v75
	goto L14
L21:
	;
	goto L16
L22:
	;
	if v55 == v48 {
		goto L9
	} else {
		goto L23
	}
L23:
	;
	v96 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v100 = *(*int32)(unsafe.Add(mBase, uint32(v96+v48<<(uint(int32(2))%32))))
	v106 = F_getRangeLongFromObjectOrReply(m, l0, v100, int32(-2147483647), int32(2147483647), v17+int32(60), int32(0))
	mBase = m.M
	v107 = m.ExcPending
	if v107 != 0 {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	return
L25:
	;
	if v106 != 0 {
		goto L1
	} else {
		goto L26
	}
L26:
	;
	v108 = *(*int32)(unsafe.Add(mBase, uint32(v17)+60))
	if v108 != 0 {
		goto L8
	} else {
		goto L27
	}
L27:
	;
	F_addReplyError(m, l0, int32(_a_F_lposCommand_2))
	mBase = m.M
	v111 = m.ExcPending
	if v111 != 0 {
		goto L24
	} else {
		goto L28
	}
L28:
	;
	goto L1
L29:
	;
	v164 = int32(_a_F_lposCommand_3)
	v167 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v54))))
	if v167 != 0 {
		goto L48
	} else {
		goto L49
	}
L30:
	;
	if v147-v149 != 0 {
		goto L29
	} else {
		goto L42
	}
L31:
	;
	v147 = F_tolower(m, v143)
	mBase = m.M
	v148 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v144))))
	v149 = F_tolower(m, v148)
	mBase = m.M
	goto L30
L32:
	;
	v117 = v54
	v118 = v112
	v119 = v115
	goto L35
L33:
	;
	v143 = int32(0)
	v144 = v112
	goto L31
L34:
	;
	v143 = v140 & int32(255)
	v144 = v139
	goto L31
L35:
	;
	v121 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v118))))
	if v121 == int32(0) {
		v139 = v118
		v140 = v119
		goto L34
	} else {
		goto L37
	}
L36:
	;
	v139 = v133
	v140 = int32(0)
	goto L34
L37:
	;
	v125 = v119 & int32(255)
	if v125 == v121 {
		goto L38
	} else {
		goto L39
	}
L38:
	;
	v132 = int32(1)
	v133 = v118 + v132
	v134 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v117)+1)))
	if v134 != 0 {
		v117 = v117 + v132
		v118 = v133
		v119 = v134
		goto L35
	} else {
		goto L41
	}
L39:
	;
	v127 = F_tolower(m, v125)
	mBase = m.M
	v128 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v118))))
	v129 = F_tolower(m, v128)
	mBase = m.M
	if v127 == v129 {
		goto L38
	} else {
		goto L40
	}
L40:
	;
	v131 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v117))))
	v139 = v118
	v140 = v131
	goto L34
L41:
	;
	goto L36
L42:
	;
	if v55 == v48 {
		goto L29
	} else {
		goto L43
	}
L43:
	;
	v152 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v156 = *(*int32)(unsafe.Add(mBase, uint32(v152+v48<<(uint(int32(2))%32))))
	v160 = F_getPositiveLongFromObjectOrReply(m, l0, v156, v17+int32(56), int32(_a_F_lposCommand_4))
	mBase = m.M
	v161 = m.ExcPending
	if v161 != 0 {
		goto L24
	} else {
		goto L44
	}
L44:
	;
	if v160 == int32(0) {
		goto L8
	} else {
		goto L45
	}
L45:
	;
	goto L1
L46:
	;
	if v199-v201 != 0 {
		goto L2
	} else {
		goto L58
	}
L47:
	;
	v199 = F_tolower(m, v195)
	mBase = m.M
	v200 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v196))))
	v201 = F_tolower(m, v200)
	mBase = m.M
	goto L46
L48:
	;
	v169 = v54
	v170 = v164
	v171 = v167
	goto L51
L49:
	;
	v195 = int32(0)
	v196 = v164
	goto L47
L50:
	;
	v195 = v192 & int32(255)
	v196 = v191
	goto L47
L51:
	;
	v173 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v170))))
	if v173 == int32(0) {
		v191 = v170
		v192 = v171
		goto L50
	} else {
		goto L53
	}
L52:
	;
	v191 = v185
	v192 = int32(0)
	goto L50
L53:
	;
	v177 = v171 & int32(255)
	if v177 == v173 {
		goto L54
	} else {
		goto L55
	}
L54:
	;
	v184 = int32(1)
	v185 = v170 + v184
	v186 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v169)+1)))
	if v186 != 0 {
		v169 = v169 + v184
		v170 = v185
		v171 = v186
		goto L51
	} else {
		goto L57
	}
L55:
	;
	v179 = F_tolower(m, v177)
	mBase = m.M
	v180 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v170))))
	v181 = F_tolower(m, v180)
	mBase = m.M
	if v179 == v181 {
		goto L54
	} else {
		goto L56
	}
L56:
	;
	v183 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v169))))
	v191 = v170
	v192 = v183
	goto L50
L57:
	;
	goto L52
L58:
	;
	if v55 == v48 {
		goto L2
	} else {
		goto L59
	}
L59:
	;
	v204 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v208 = *(*int32)(unsafe.Add(mBase, uint32(v204+v48<<(uint(int32(2))%32))))
	v212 = F_getPositiveLongFromObjectOrReply(m, l0, v208, v17+int32(52), int32(_a_F_lposCommand_5))
	mBase = m.M
	v213 = m.ExcPending
	if v213 != 0 {
		goto L24
	} else {
		goto L60
	}
L60:
	;
	if v212 != 0 {
		goto L1
	} else {
		goto L61
	}
L61:
	;
	goto L8
L62:
	;
	goto L7
L63:
	;
	v222 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v17)+60)) = v222 - v219
	v229 = v222
	v230 = v219
	goto L3
L64:
	;
	v261 = F_checkType(m, l0, v243, int32(1))
	mBase = m.M
	v262 = m.ExcPending
	if v262 != 0 {
		goto L24
	} else {
		goto L71
	}
L65:
	;
	if v243 != 0 {
		goto L64
	} else {
		goto L66
	}
L66:
	;
	v245 = *(*int32)(unsafe.Add(mBase, uint32(v17)+56))
	if v245 == int32(-1) {
		goto L67
	} else {
		goto L68
	}
L67:
	;
	v253 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+224)))
	v257 = *(*int32)(unsafe.Add(mBase, uint32(v253<<(uint(int32(2))%32))+uint32(_c_F_lposCommand[1])))
	F_addReply(m, l0, v257)
	mBase = m.M
	v259 = m.ExcPending
	if v259 != 0 {
		goto L24
	} else {
		goto L70
	}
L68:
	;
	v249 = *(*int32)(unsafe.Add(mBase, _c_F_lposCommand[2]))
	F_addReply(m, l0, v249)
	mBase = m.M
	v251 = m.ExcPending
	if v251 != 0 {
		goto L24
	} else {
		goto L69
	}
L69:
	;
	goto L1
L70:
	;
	goto L1
L71:
	;
	if v261 != 0 {
		goto L1
	} else {
		goto L72
	}
L72:
	;
	v263 = int32(-1)
	v264 = int32(0)
	v266 = *(*int32)(unsafe.Add(mBase, uint32(v17)+56))
	if v266 == v263 {
		v271 = v264
		goto L73
	} else {
		goto L74
	}
L73:
	;
	v274 = F_listTypeInitIterator(m, v243, v230>>(uint(int32(31))%32), v229)
	mBase = m.M
	v275 = m.ExcPending
	if v275 != 0 {
		goto L24
	} else {
		goto L76
	}
L74:
	;
	v269 = F_addReplyDeferredLen(m, l0)
	mBase = m.M
	v270 = m.ExcPending
	if v270 != 0 {
		goto L24
	} else {
		goto L75
	}
L75:
	;
	v271 = v269
	goto L73
L76:
	;
	v276 = F_listTypeLength(m, v243)
	mBase = m.M
	v277 = m.ExcPending
	if v277 != 0 {
		goto L24
	} else {
		goto L77
	}
L77:
	;
	v280 = F_listTypeNext(m, v274, v17+int32(8))
	mBase = m.M
	v281 = m.ExcPending
	if v281 != 0 {
		goto L24
	} else {
		goto L79
	}
L78:
	;
	v366 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v274)+4)))
	if v366 != int32(9) {
		goto L102
	} else {
		goto L103
	}
L79:
	;
	if v280 == int32(0) {
		v360 = v263
		v361 = v264
		goto L78
	} else {
		goto L80
	}
L80:
	;
	v284 = int32(0)
	v292 = v284
	v294 = v284
	v296 = v284
	goto L82
L81:
	;
	v360 = int32(-1)
	v361 = v348
	goto L78
L82:
	;
	v301 = *(*int32)(unsafe.Add(mBase, uint32(v17)+52))
	if v301 == int32(0) {
		goto L84
	} else {
		goto L85
	}
L83:
	;
	v348 = v335
	goto L81
L84:
	;
	v307 = F_listTypeEqual(m, v17+int32(8), v20)
	mBase = m.M
	v308 = m.ExcPending
	if v308 != 0 {
		goto L24
	} else {
		goto L88
	}
L85:
	;
	if v301 <= v292 {
		v348 = v296
		goto L81
	} else {
		goto L86
	}
L86:
	;
	goto L84
L87:
	;
	v342 = F_listTypeNext(m, v274, v17+int32(8))
	mBase = m.M
	v343 = m.ExcPending
	if v343 != 0 {
		goto L24
	} else {
		goto L100
	}
L88:
	;
	if v307 == int32(0) {
		v333 = v294
		v335 = v296
		goto L87
	} else {
		goto L89
	}
L89:
	;
	v312 = v294 + int32(1)
	v313 = *(*int32)(unsafe.Add(mBase, uint32(v17)+60))
	if v313 <= v312 {
		goto L90
	} else {
		goto L91
	}
L90:
	;
	if v229 != 0 {
		goto L92
	} else {
		goto L93
	}
L91:
	;
	v333 = v312
	v335 = v296
	goto L87
L92:
	;
	v318 = v292
	goto L94
L93:
	;
	v318 = v276 + (v292 ^ int32(-1))
	goto L94
L94:
	;
	if v271 == int32(0) {
		v360 = v318
		v361 = v296
		goto L78
	} else {
		goto L95
	}
L95:
	;
	F_addReplyLongLong(m, l0, base.I64_extend_i32_s(v318))
	mBase = m.M
	v323 = m.ExcPending
	if v323 != 0 {
		goto L24
	} else {
		goto L96
	}
L96:
	;
	v325 = v296 + int32(1)
	v326 = *(*int32)(unsafe.Add(mBase, uint32(v17)+56))
	if v326 != 0 {
		goto L97
	} else {
		goto L98
	}
L97:
	;
	v327 = *(*int32)(unsafe.Add(mBase, uint32(v17)+60))
	if v326 <= v294-v327+int32(2) {
		v360 = v318
		v361 = v325
		goto L78
	} else {
		goto L99
	}
L98:
	;
	v333 = v312
	v335 = v325
	goto L87
L99:
	;
	v333 = v312
	v335 = v325
	goto L87
L100:
	;
	if v342 != 0 {
		v292 = v292 + int32(1)
		v294 = v333
		v296 = v335
		goto L82
	} else {
		goto L101
	}
L101:
	;
	goto L83
L102:
	;
	F_valkey_free(m, v274)
	mBase = m.M
	v373 = m.ExcPending
	if v373 != 0 {
		goto L24
	} else {
		goto L105
	}
L103:
	;
	v369 = *(*int32)(unsafe.Add(mBase, uint32(v274)+12))
	F_quicklistReleaseIterator(m, v369)
	mBase = m.M
	v371 = m.ExcPending
	if v371 != 0 {
		goto L24
	} else {
		goto L104
	}
L104:
	;
	goto L102
L105:
	;
	if v271 == int32(0) {
		goto L106
	} else {
		goto L107
	}
L106:
	;
	if v360 == int32(-1) {
		goto L109
	} else {
		goto L110
	}
L107:
	;
	F_setDeferredArrayLen(m, l0, v271, v361)
	mBase = m.M
	v377 = m.ExcPending
	if v377 != 0 {
		goto L24
	} else {
		goto L108
	}
L108:
	;
	goto L1
L109:
	;
	v384 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+224)))
	v388 = *(*int32)(unsafe.Add(mBase, uint32(v384<<(uint(int32(2))%32))+uint32(_c_F_lposCommand[1])))
	F_addReply(m, l0, v388)
	mBase = m.M
	v390 = m.ExcPending
	if v390 != 0 {
		goto L24
	} else {
		goto L112
	}
L110:
	;
	F_addReplyLongLong(m, l0, base.I64_extend_i32_s(v360))
	mBase = m.M
	v382 = m.ExcPending
	if v382 != 0 {
		goto L24
	} else {
		goto L111
	}
L111:
	;
	goto L1
L112:
	;
	goto L1
L113:
	;
	goto L1
}
func F_lpushCommand(m *base.Module, l0 int32) {
	var v2 int32
	_ = v2
	var v5 int32
	_ = v5
	v2 = int32(0)
	F_pushGenericCommand(m, l0, v2, v2)
	v5 = m.ExcPending
	if v5 != 0 {
		return
	} else {
		return
	}
}
func F_lpushxCommand(m *base.Module, l0 int32) {
	var v5 int32
	_ = v5
	F_pushGenericCommand(m, l0, int32(0), int32(1))
	v5 = m.ExcPending
	if v5 != 0 {
		return
	} else {
		return
	}
}
func F_lremCommand(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v68 int64
	_ = v68
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
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
	var v123 int32
	_ = v123
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v134 int32
	_ = v134
	var v137 int32
	_ = v137
	v8 = m.G0
	v10 = v8 - int32(48)
	m.G0 = v10
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v13 = *(*int32)(unsafe.Add(mBase, uint32(v12)+12))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(v12)+8))
	v20 = F_getRangeLongFromObjectOrReply(m, l0, v14, int32(-2147483647), int32(2147483647), v10+int32(44), int32(0))
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v10 + int32(48)
	return
L2:
	;
	return
L3:
	;
	if v20 != 0 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v23 = *(*int32)(unsafe.Add(mBase, uint32(v22)+4))
	v25 = *(*int32)(unsafe.Add(mBase, _c_F_lremCommand[0]))
	v26 = F_lookupKeyWriteOrReply(m, l0, v23, v25)
	mBase = m.M
	v27 = m.ExcPending
	if v27 != 0 {
		goto L2
	} else {
		goto L5
	}
L5:
	;
	if v26 == int32(0) {
		goto L1
	} else {
		goto L6
	}
L6:
	;
	v31 = F_checkType(m, l0, v26, int32(1))
	mBase = m.M
	v32 = m.ExcPending
	if v32 != 0 {
		goto L2
	} else {
		goto L7
	}
L7:
	;
	if v31 != 0 {
		goto L1
	} else {
		goto L8
	}
L8:
	;
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v10)+44))
	if int32(-1) < v33 {
		goto L10
	} else {
		goto L11
	}
L9:
	;
	v48 = int32(0)
	v49 = F_listTypeNext(m, v47, v10)
	mBase = m.M
	v50 = m.ExcPending
	if v50 != 0 {
		goto L2
	} else {
		goto L15
	}
L10:
	;
	v45 = F_listTypeInitIterator(m, v26, int32(0), int32(1))
	mBase = m.M
	v46 = m.ExcPending
	if v46 != 0 {
		goto L2
	} else {
		goto L13
	}
L11:
	;
	v36 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v10)+44)) = v36 - v33
	v41 = F_listTypeInitIterator(m, v26, int32(-1), v36)
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L2
	} else {
		goto L12
	}
L12:
	;
	v47 = v41
	goto L9
L13:
	;
	v47 = v45
	goto L9
L14:
	;
	v89 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v47)+4)))
	if v89 != int32(9) {
		goto L27
	} else {
		goto L28
	}
L15:
	;
	if v49 == int32(0) {
		v87 = v48
		goto L14
	} else {
		goto L16
	}
L16:
	;
	v58 = v48
	goto L17
L17:
	;
	v60 = F_listTypeEqual(m, v10, v13)
	mBase = m.M
	v61 = m.ExcPending
	if v61 != 0 {
		goto L2
	} else {
		goto L20
	}
L18:
	;
	v87 = v78
	goto L14
L19:
	;
	v80 = F_listTypeNext(m, v47, v10)
	mBase = m.M
	v81 = m.ExcPending
	if v81 != 0 {
		goto L2
	} else {
		goto L25
	}
L20:
	;
	if v60 == int32(0) {
		v78 = v58
		goto L19
	} else {
		goto L21
	}
L21:
	;
	F_listTypeDelete(m, v47, v10)
	mBase = m.M
	v65 = m.ExcPending
	if v65 != 0 {
		goto L2
	} else {
		goto L22
	}
L22:
	;
	v66 = int32(_a_F_lremCommand_0)
	v68 = *(*int64)(unsafe.Add(mBase, _c_F_lremCommand[1]))
	*(*int64)(unsafe.Add(mBase, _c_F_lremCommand[1])) = v68 + int64(1)
	v73 = v58 + int32(1)
	v74 = *(*int32)(unsafe.Add(mBase, uint32(v10)+44))
	if v74 == int32(0) {
		v78 = v73
		goto L19
	} else {
		goto L23
	}
L23:
	;
	if v73 == v74 {
		v87 = v73
		goto L14
	} else {
		goto L24
	}
L24:
	;
	v78 = v73
	goto L19
L25:
	;
	if v80 != 0 {
		v58 = v78
		goto L17
	} else {
		goto L26
	}
L26:
	;
	goto L18
L27:
	;
	F_valkey_free(m, v47)
	mBase = m.M
	v96 = m.ExcPending
	if v96 != 0 {
		goto L2
	} else {
		goto L30
	}
L28:
	;
	v92 = *(*int32)(unsafe.Add(mBase, uint32(v47)+12))
	F_quicklistReleaseIterator(m, v92)
	mBase = m.M
	v94 = m.ExcPending
	if v94 != 0 {
		goto L2
	} else {
		goto L29
	}
L29:
	;
	goto L27
L30:
	;
	if v87 == int32(0) {
		goto L31
	} else {
		goto L32
	}
L31:
	;
	F_addReplyLongLong(m, l0, base.I64_extend_i32_s(v87))
	mBase = m.M
	v137 = m.ExcPending
	if v137 != 0 {
		goto L2
	} else {
		goto L42
	}
L32:
	;
	v101 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v102 = *(*int32)(unsafe.Add(mBase, uint32(v101)+4))
	v103 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v103)+28))
	F_notifyKeyspaceEvent(m, int32(16), int32(_a_F_lremCommand_1), v102, v104)
	mBase = m.M
	v106 = m.ExcPending
	if v106 != 0 {
		goto L2
	} else {
		goto L33
	}
L33:
	;
	v107 = F_listTypeLength(m, v26)
	mBase = m.M
	v108 = m.ExcPending
	if v108 != 0 {
		goto L2
	} else {
		goto L36
	}
L34:
	;
	v130 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v131 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v132 = *(*int32)(unsafe.Add(mBase, uint32(v131)+4))
	F_signalModifiedKey(m, l0, v130, v132)
	mBase = m.M
	v134 = m.ExcPending
	if v134 != 0 {
		goto L2
	} else {
		goto L41
	}
L35:
	;
	v123 = int32(0)
	F_listTypeTryConversionRaw(m, v26, int32(2), v123, v123, v123, v123, v123)
	mBase = m.M
	v129 = m.ExcPending
	if v129 != 0 {
		goto L2
	} else {
		goto L40
	}
L36:
	;
	if v107 != 0 {
		goto L35
	} else {
		goto L37
	}
L37:
	;
	v109 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v110 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v111 = *(*int32)(unsafe.Add(mBase, uint32(v110)+4))
	v112 = F_dbDelete(m, v109, v111)
	mBase = m.M
	v113 = m.ExcPending
	if v113 != 0 {
		goto L2
	} else {
		goto L38
	}
L38:
	;
	v116 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v117 = *(*int32)(unsafe.Add(mBase, uint32(v116)+4))
	v118 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v119 = *(*int32)(unsafe.Add(mBase, uint32(v118)+28))
	F_notifyKeyspaceEvent(m, int32(4), int32(_a_F_lremCommand_2), v117, v119)
	mBase = m.M
	v121 = m.ExcPending
	if v121 != 0 {
		goto L2
	} else {
		goto L39
	}
L39:
	;
	goto L34
L40:
	;
	goto L34
L41:
	;
	goto L31
L42:
	;
	goto L1
}
func F_lru_getIdleSecs(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	v3 = *(*int32)(unsafe.Add(mBase, _c_F_lru_getIdleSecs[0]))
	return (v3 - l0) & int32(16777215)
}
func F_lrulfu_init(m *base.Module) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v6 int32
	_ = v6
	var v13 int32
	_ = v13
	v2 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_lrulfu_init[0])))
	if v2 != int32(1) {
		v13 = *(*int32)(unsafe.Add(mBase, _c_F_lrulfu_init[1]))
		return v13 & int32(16777215)
	} else {
		v6 = int32(*(*uint16)(unsafe.Add(mBase, _c_F_lrulfu_init[2])))
		return v6<<(uint(int32(8))%32) | int32(5)
	}
}
func F_lrulfu_touch(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v35 int32
	_ = v35
	var v37 int64
	_ = v37
	var v41 int64
	_ = v41
	var v56 float64
	_ = v56
	var v58 int32
	_ = v58
	var v66 int32
	_ = v66
	var v72 int32
	_ = v72
	v2 = int32(0)
	v7 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_lrulfu_touch[0])))
	if v7 != int32(1) {
		v72 = *(*int32)(unsafe.Add(mBase, _c_F_lrulfu_touch[1]))
		return v72 & int32(16777215)
	} else {
		v10 = int32(0)
		v11 = int32(*(*uint16)(unsafe.Add(mBase, _c_F_lrulfu_touch[2])))
		v13 = *(*int32)(unsafe.Add(mBase, _c_F_lrulfu_touch[3]))
		if v13 == v10 {
			v24 = v2
		} else {
			v19 = int32(65535)
			v21 = base.I32_div_s((v11-int32(base.Ui32(l0)>>(uint(int32(8))%32)))&v19, v13)
			v24 = v21 & v19
		}
		v25 = int32(255)
		v28 = l0 & v25
		v29 = v28 - v24
		if base.Ui32(v28) < base.Ui32(v29) {
			v31 = int32(0)
		} else {
			v31 = v29
		}
		if v31 == int32(255) {
			v66 = v25
		} else {
			v35 = int32(0)
			v37 = *(*int64)(unsafe.Add(mBase, _c_F_lrulfu_touch[4]))
			v41 = v37*int64(6364136223846793005) + int64(1)
			*(*int64)(unsafe.Add(mBase, _c_F_lrulfu_touch[4])) = v41
			if base.Ui32(v31) < base.Ui32(int32(5)) {
				v56 = float64(0)
			} else {
				v56 = base.F64_convert_i32_s(v31 + int32(-5))
			}
			v58 = *(*int32)(unsafe.Add(mBase, _c_F_lrulfu_touch[5]))
			v66 = v31 + base.F64_lt(base.F64_div(base.F64_convert_i32_s(base.I32_wrap_i64(int64(base.Ui64(v41)>>(uint(int64(33))%64)))), float64(2.147483647e+09)), base.F64_div(float64(1), base.F64_add(base.F64_mul(v56, base.F64_convert_i32_s(v58)), float64(1))))
		}
		return v66 + v11<<(uint(int32(8))%32)
	}
}
func F_lstat(m *base.Module, l0 int32, l1 int32) int32 {
	var v5 int32
	_ = v5
	v5 = F___fstatat(m, int32(-100), l0, l1, int32(256))
	return v5
}
func F_luaE_newthread(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
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
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v31 int64
	_ = v31
	var v33 int32
	_ = v33
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v82 int64
	_ = v82
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	v2 = int32(0)
	v9 = F_luaM_realloc_(m, l0, v2, v2, int32(120))
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		return int32(0)
	} else {
		v13 = int32(8)
		v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
		v15 = *(*int32)(unsafe.Add(mBase, uint32(v14)+28))
		*(*int32)(unsafe.Add(mBase, uint32(v9))) = v15
		*(*int32)(unsafe.Add(mBase, uint32(v14)+28)) = v9
		v18 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+20)))
		*(*uint8)(unsafe.Add(mBase, uint32(v9)+4)) = uint8(v13)
		v21 = v18 & int32(3)
		*(*uint8)(unsafe.Add(mBase, uint32(v9)+5)) = uint8(v21)
		v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
		v24 = int32(0)
		*(*int32)(unsafe.Add(mBase, uint32(v9)+32)) = v24
		*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = v23
		*(*int32)(unsafe.Add(mBase, uint32(v9)+68)) = v24
		*(*int32)(unsafe.Add(mBase, uint32(v9)+104)) = v24
		v31 = int64(0)
		*(*int64)(unsafe.Add(mBase, uint32(v9)+60)) = v31
		v33 = int32(256)
		*(*uint16)(unsafe.Add(mBase, uint32(v9)+56)) = uint16(v33)
		*(*int64)(unsafe.Add(mBase, uint32(v9)+40)) = v31
		*(*uint8)(unsafe.Add(mBase, uint32(v9)+6)) = uint8(v24)
		*(*int64)(unsafe.Add(mBase, uint32(v9)+112)) = v31
		*(*int64)(unsafe.Add(mBase, uint32(v9)+20)) = v31
		*(*int32)(unsafe.Add(mBase, uint32(v9)+80)) = v24
		*(*int64)(unsafe.Add(mBase, uint32(v9)+48)) = v31
		v50 = F_luaM_realloc_(m, l0, v24, v24, int32(192))
		mBase = m.M
		v51 = m.ExcPending
		if v51 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v9)+48)) = int32(8)
			*(*int32)(unsafe.Add(mBase, uint32(v9)+20)) = v50
			*(*int32)(unsafe.Add(mBase, uint32(v9)+40)) = v50
			*(*int32)(unsafe.Add(mBase, uint32(v9)+36)) = v50 + int32(168)
			v59 = int32(0)
			v62 = F_luaM_realloc_(m, l0, v59, v59, int32(720))
			mBase = m.M
			v63 = m.ExcPending
			if v63 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v9)+44)) = int32(45)
				*(*int32)(unsafe.Add(mBase, uint32(v9)+32)) = v62
				*(*int32)(unsafe.Add(mBase, uint32(v9)+28)) = v62 + int32(624)
				v70 = *(*int32)(unsafe.Add(mBase, uint32(v9)+20))
				*(*int32)(unsafe.Add(mBase, uint32(v70)+4)) = v62
				v73 = v62 + int32(16)
				*(*int32)(unsafe.Add(mBase, uint32(v9)+8)) = v73
				*(*int32)(unsafe.Add(mBase, uint32(v62)+8)) = int32(0)
				*(*int32)(unsafe.Add(mBase, uint32(v70))) = v73
				*(*int32)(unsafe.Add(mBase, uint32(v9)+12)) = v73
				*(*int32)(unsafe.Add(mBase, uint32(v70)+8)) = v62 + int32(336)
				v82 = *(*int64)(unsafe.Add(mBase, uint32(l0)+72))
				*(*int64)(unsafe.Add(mBase, uint32(v9)+72)) = v82
				v84 = *(*int32)(unsafe.Add(mBase, uint32(l0)+80))
				*(*int32)(unsafe.Add(mBase, uint32(v9)+80)) = v84
				v86 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+56)))
				*(*uint8)(unsafe.Add(mBase, uint32(v9)+56)) = uint8(v86)
				v88 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
				*(*int32)(unsafe.Add(mBase, uint32(v9)+60)) = v88
				v90 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
				*(*int32)(unsafe.Add(mBase, uint32(v9)+64)) = v88
				*(*int32)(unsafe.Add(mBase, uint32(v9)+68)) = v90
				return v9
			}
		}
	}
}
func F_luaG_concaterror(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v29 int32
	_ = v29
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v49 int32
	_ = v49
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v65 int32
	_ = v65
	var v75 int32
	_ = v75
	var v82 int32
	_ = v82
	v9 = m.G0
	v11 = v9 - int32(48)
	m.G0 = v11
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+44)) = int32(0)
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v17 = *(*int32)(unsafe.Add(mBase, uint32(v16)+8))
	v18 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
	if base.Ui32(v13+int32(-3)) < base.Ui32(int32(2)) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v23 = l2
	goto L3
L2:
	;
	v23 = l1
	goto L3
L3:
	;
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v23)+8))
	v25 = m.G399
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v25+v24<<(uint(int32(2))%32))))
	if base.Ui32(v17) <= base.Ui32(v18) {
		goto L5
	} else {
		goto L6
	}
L4:
	;
	m.G0 = v11 + int32(48)
	return
L5:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+4)) = v29
	v75 = m.G3
	*(*int32)(unsafe.Add(mBase, uint32(v11))) = v75 + int32(_a_F_luaG_concaterror_0)
	F_luaG_runerror(m, l0, v75+int32(_a_F_luaG_concaterror_1), v11)
	mBase = m.M
	v82 = m.ExcPending
	if v82 != 0 {
		goto L13
	} else {
		goto L15
	}
L6:
	;
	v38 = v18
	goto L8
L7:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v49 = F_getobjname(m, l0, v16, (v23-v43)>>(uint(int32(4))%32), v11+int32(44))
	mBase = m.M
	if v49 == int32(0) {
		goto L5
	} else {
		goto L12
	}
L8:
	;
	if v23 == v38 {
		goto L7
	} else {
		goto L10
	}
L10:
	;
	v41 = v38 + int32(16)
	if base.Ui32(v17) <= base.Ui32(v41) {
		goto L5
	} else {
		goto L11
	}
L11:
	;
	v38 = v41
	goto L8
L12:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+20)) = v49
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v11)+44))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+24)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v11)+28)) = v29
	v56 = m.G3
	*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = v56 + int32(_a_F_luaG_concaterror_0)
	F_luaG_runerror(m, l0, v56+int32(_a_F_luaG_concaterror_2), v11+int32(16))
	mBase = m.M
	v65 = m.ExcPending
	if v65 != 0 {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	return
L14:
	;
	goto L4
L15:
	;
	goto L4
}
func F_luaG_ordererror(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
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
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v40 int32
	_ = v40
	v6 = m.G0
	v8 = v6 - int32(32)
	m.G0 = v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	v11 = m.G399
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v13 = int32(2)
	v16 = *(*int32)(unsafe.Add(mBase, uint32(v11+v12<<(uint(v13)%32))))
	v17 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16)+2)))
	v21 = *(*int32)(unsafe.Add(mBase, uint32(v11+v10<<(uint(v13)%32))))
	v22 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21)+2)))
	if v17 != v22 {
		*(*int32)(unsafe.Add(mBase, uint32(v8)+20)) = v21
		*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = v16
		v34 = m.G3
		F_luaG_runerror(m, l0, v34+int32(_a_F_luaG_ordererror_0), v8+int32(16))
		mBase = m.M
		v40 = m.ExcPending
		if v40 != 0 {
			return int32(0)
		} else {
			m.G0 = v8 + int32(32)
			return int32(0)
		}
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(v8))) = v16
		v25 = m.G3
		F_luaG_runerror(m, l0, v25+int32(_a_F_luaG_ordererror_1), v8)
		mBase = m.M
		v31 = m.ExcPending
		if v31 != 0 {
			return int32(0)
		} else {
			m.G0 = v8 + int32(32)
			return int32(0)
		}
	}
}
func F_luaT_init(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v31 int32
	_ = v31
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v57 int32
	_ = v57
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	v7 = int32(0)
	goto L1
L1:
	;
	v10 = m.G3
	v14 = v7 << (uint(int32(2)) % 32)
	v16 = *(*int32)(unsafe.Add(mBase, uint32(v10+int32(_a_F_luaT_init_0)+v14)))
	if v16&int32(3) == int32(0) {
		v38 = v16
		goto L5
	} else {
		goto L6
	}
L2:
	;
	return
L3:
	;
	v72 = F_luaS_newlstr(m, l0, v16, v71)
	mBase = m.M
	v73 = m.ExcPending
	if v73 != 0 {
		goto L19
	} else {
		goto L20
	}
L4:
	;
	v71 = v63 - v16
	goto L3
L5:
	;
	v42 = v38
	goto L13
L6:
	;
	v24 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16))))
	if v24 != 0 {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	v27 = v16
	goto L9
L8:
	;
	v71 = v16 - v16
	goto L3
L9:
	;
	v31 = v27 + int32(1)
	if v31&int32(3) == int32(0) {
		v38 = v31
		goto L5
	} else {
		goto L11
	}
L11:
	;
	v36 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v31))))
	if v36 != 0 {
		v27 = v31
		goto L9
	} else {
		goto L12
	}
L12:
	;
	v63 = v31
	goto L4
L13:
	;
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v42)))
	v51 = int32(-2139062144)
	if (int32(16843008)-v48|v48)&v51 == v51 {
		v42 = v42 + int32(4)
		goto L13
	} else {
		goto L15
	}
L14:
	;
	v57 = v42
	goto L16
L15:
	;
	goto L14
L16:
	;
	v61 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v57))))
	if v61 != 0 {
		v57 = v57 + int32(1)
		goto L16
	} else {
		goto L18
	}
L17:
	;
	v63 = v57
	goto L4
L18:
	;
	goto L17
L19:
	;
	return
L20:
	;
	v74 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v74+v14)+188)) = v72
	v77 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v72)+5)))
	v79 = v77 | int32(32)
	*(*uint8)(unsafe.Add(mBase, uint32(v72)+5)) = uint8(v79)
	v82 = v7 + int32(1)
	if v82 != int32(17) {
		v7 = v82
		goto L1
	} else {
		goto L21
	}
L21:
	;
	goto L2
}
func F_luaU_dump(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v16 int32
	_ = v16
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	v7 = m.G0
	v9 = v7 - int32(32)
	m.G0 = v9
	*(*int32)(unsafe.Add(mBase, uint32(v9)+12)) = l4
	*(*int32)(unsafe.Add(mBase, uint32(v9)+8)) = l3
	*(*int32)(unsafe.Add(mBase, uint32(v9)+4)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(v9))) = l0
	v16 = v9 + int32(20)
	*(*int64)(unsafe.Add(mBase, uint32(v16)+4)) = int64(2256215107240017)
	*(*int32)(unsafe.Add(mBase, uint32(v16))) = int32(1635077147)
	v24 = m.T0[l2].(func(*base.Module, int32, int32, int32, int32) int32)(m, l0, v9+int32(20), int32(12), l3)
	mBase = m.M
	v27 = m.ExcPending
	if v27 != 0 {
		return int32(0)
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = v24
		F_DumpFunction(m, l1, int32(0), v9)
		mBase = m.M
		v31 = m.ExcPending
		if v31 != 0 {
			return int32(0)
		} else {
			v32 = *(*int32)(unsafe.Add(mBase, uint32(v9)+16))
			m.G0 = v9 + int32(32)
			return v32
		}
	}
}
func F_luaU_header(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	*(*int64)(unsafe.Add(mBase, uint32(l0)+4)) = int64(2256215107240017)
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(1635077147)
	return
}
func F_lwFreeCanvas(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	F_valkey_free(m, v2)
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
