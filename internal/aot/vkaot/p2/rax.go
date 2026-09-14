package p2

import (
	base "github.com/shibukawa/vkmem/internal/aot/vkaot/base"
	"unsafe"
)

func F_raxAddChild(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v14 int32
	_ = v14
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
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
	var v60 int32
	_ = v60
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v82 int32
	_ = v82
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	var v101 int32
	_ = v101
	var v109 int32
	_ = v109
	var v115 int32
	_ = v115
	var v121 int32
	_ = v121
	var v123 int32
	_ = v123
	var v126 int32
	_ = v126
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v138 int32
	_ = v138
	var v141 int32
	_ = v141
	var v144 int32
	_ = v144
	var v148 int32
	_ = v148
	var v155 int32
	_ = v155
	var v158 int32
	_ = v158
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
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
	var v189 int32
	_ = v189
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v198 int32
	_ = v198
	var v204 int32
	_ = v204
	var v211 int32
	_ = v211
	var v215 int32
	_ = v215
	var v218 int32
	_ = v218
	var v224 int32
	_ = v224
	var v231 int32
	_ = v231
	var v235 int32
	_ = v235
	var v238 int32
	_ = v238
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v248 int32
	_ = v248
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v252 int32
	_ = v252
	var v254 int32
	_ = v254
	var v255 int32
	_ = v255
	var v257 int32
	_ = v257
	var v259 int32
	_ = v259
	var v263 int32
	_ = v263
	var v264 int32
	_ = v264
	var v265 int32
	_ = v265
	var v270 int32
	_ = v270
	var v271 int32
	_ = v271
	var v272 int32
	_ = v272
	var v274 int32
	_ = v274
	var v276 int32
	_ = v276
	var v281 int32
	_ = v281
	var v295 int32
	_ = v295
	var v296 int32
	_ = v296
	var v297 int32
	_ = v297
	var v303 int32
	_ = v303
	var v304 int32
	_ = v304
	var v308 int32
	_ = v308
	var v315 int32
	_ = v315
	var v318 int32
	_ = v318
	var v325 int32
	_ = v325
	var v326 int32
	_ = v326
	var v327 int32
	_ = v327
	var v331 int32
	_ = v331
	var v333 int32
	_ = v333
	var v334 int32
	_ = v334
	var v336 int32
	_ = v336
	var v338 int32
	_ = v338
	var v349 int32
	_ = v349
	var v355 int32
	_ = v355
	var v356 int32
	_ = v356
	var v358 int32
	_ = v358
	var v364 int32
	_ = v364
	var v371 int32
	_ = v371
	var v375 int32
	_ = v375
	var v378 int32
	_ = v378
	var v384 int32
	_ = v384
	var v391 int32
	_ = v391
	var v395 int32
	_ = v395
	var v398 int32
	_ = v398
	var v401 int32
	_ = v401
	var v402 int32
	_ = v402
	var v403 int32
	_ = v403
	var v408 int32
	_ = v408
	var v409 int32
	_ = v409
	var v410 int32
	_ = v410
	var v412 int32
	_ = v412
	var v414 int32
	_ = v414
	var v415 int32
	_ = v415
	var v417 int32
	_ = v417
	var v419 int32
	_ = v419
	var v423 int32
	_ = v423
	var v424 int32
	_ = v424
	var v425 int32
	_ = v425
	var v430 int32
	_ = v430
	var v431 int32
	_ = v431
	var v432 int32
	_ = v432
	var v434 int32
	_ = v434
	var v436 int32
	_ = v436
	var v441 int32
	_ = v441
	var v454 int32
	_ = v454
	var v456 int32
	_ = v456
	var v457 int32
	_ = v457
	var v460 int32
	_ = v460
	var v464 int32
	_ = v464
	var v471 int32
	_ = v471
	var v474 int32
	_ = v474
	var v481 int32
	_ = v481
	var v482 int32
	_ = v482
	var v483 int32
	_ = v483
	var v487 int32
	_ = v487
	var v489 int32
	_ = v489
	var v490 int32
	_ = v490
	var v492 int32
	_ = v492
	var v494 int32
	_ = v494
	var v505 int32
	_ = v505
	var v511 int32
	_ = v511
	var v512 int32
	_ = v512
	var v514 int32
	_ = v514
	var v520 int32
	_ = v520
	var v527 int32
	_ = v527
	var v531 int32
	_ = v531
	var v534 int32
	_ = v534
	var v540 int32
	_ = v540
	var v547 int32
	_ = v547
	var v551 int32
	_ = v551
	var v554 int32
	_ = v554
	var v557 int32
	_ = v557
	var v558 int32
	_ = v558
	var v559 int32
	_ = v559
	var v564 int32
	_ = v564
	var v565 int32
	_ = v565
	var v566 int32
	_ = v566
	var v568 int32
	_ = v568
	var v570 int32
	_ = v570
	var v571 int32
	_ = v571
	var v573 int32
	_ = v573
	var v575 int32
	_ = v575
	var v579 int32
	_ = v579
	var v580 int32
	_ = v580
	var v581 int32
	_ = v581
	var v586 int32
	_ = v586
	var v587 int32
	_ = v587
	var v588 int32
	_ = v588
	var v590 int32
	_ = v590
	var v592 int32
	_ = v592
	var v597 int32
	_ = v597
	var v613 int32
	_ = v613
	var v618 int32
	_ = v618
	var v619 int32
	_ = v619
	var v626 int32
	_ = v626
	var v635 int32
	_ = v635
	var v638 int32
	_ = v638
	v2 = l1
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v14&int32(4) != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	F_valkey_free(m, v18)
	mBase = m.M
	v638 = m.ExcPending
	if v638 != 0 {
		goto L5
	} else {
		goto L144
	}
L2:
	;
	F__serverAssert(m, int32(_a_F_raxAddChild_0), int32(_a_F_raxAddChild_1), int32(243))
	mBase = m.M
	v635 = m.ExcPending
	if v635 != 0 {
		goto L5
	} else {
		goto L143
	}
L3:
	;
	v18 = F_valkey_malloc(m, int32(4))
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		goto L5
	} else {
		goto L6
	}
L4:
	;
	v24 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v18))) = v24
	v30 = v14&int32(-8) + int32(8)
	v31 = int32(3)
	v32 = int32(base.Ui32(v30) >> (uint(v31) % 32))
	v33 = int32(1)
	v41 = int32(4)
	v42 = v32 + int32(base.Ui32(v30)>>(uint(v33)%32)) + (v24-v32)&v31 + v41
	v53 = (v14 ^ int32(-1)) << (uint(v33) % 32) & (v24 - v14&v33) & v41
	v54 = v42 + v53
	v55 = F_valkey_realloc(m, l0, v54)
	mBase = m.M
	v56 = m.ExcPending
	if v56 != 0 {
		goto L5
	} else {
		goto L8
	}
L5:
	;
	return int32(0)
L6:
	;
	if v18 != 0 {
		goto L4
	} else {
		goto L7
	}
L7:
	;
	return int32(0)
L8:
	;
	if v55 == int32(0) {
		goto L1
	} else {
		goto L9
	}
L9:
	;
	v59 = int32(3)
	v60 = int32(base.Ui32(v14) >> (uint(v59) % 32))
	v70 = v55 + int32(4)
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v55)))
	v73 = int32(base.Ui32(v71) >> (uint(v59) % 32))
	if base.Ui32(v71) < base.Ui32(int32(8)) {
		v101 = v24
		goto L10
	} else {
		goto L11
	}
L10:
	;
	v109 = v60 + (int32(0)-v60)&v59 + v60<<(uint(int32(2))%32) + int32(4)
	if v71&int32(3) != int32(1) {
		v126 = v73
		goto L16
	} else {
		goto L17
	}
L11:
	;
	v82 = v24
	goto L12
L12:
	;
	v90 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v70+v82))))
	if base.Ui32(v2) < base.Ui32(v90) {
		v101 = v82
		goto L10
	} else {
		goto L14
	}
L13:
	;
	v101 = v73
	goto L10
L14:
	;
	v93 = v82 + int32(1)
	if v93 != v73 {
		v82 = v93
		goto L12
	} else {
		goto L15
	}
L15:
	;
	goto L13
L16:
	;
	v133 = int32(2)
	v134 = v101 << (uint(v133) % 32)
	v135 = v70 + v126 + (int32(0)-v126)&int32(3) + v134
	v138 = v42 - v109 + int32(-4)
	v141 = v135 + v138 + int32(4)
	v144 = (v126 - v101) << (uint(v133) % 32)
	if v141 == v135 {
		goto L19
	} else {
		goto L20
	}
L17:
	;
	v115 = int32(-4)
	v121 = *(*int32)(unsafe.Add(mBase, uint32(v55+v109+v53+v115)))
	*(*int32)(unsafe.Add(mBase, uint32(v55+v54+v115))) = v121
	v123 = *(*int32)(unsafe.Add(mBase, uint32(v55)))
	v126 = int32(base.Ui32(v123) >> (uint(int32(3)) % 32))
	goto L16
L18:
	;
	if v138 == int32(0) {
		goto L59
	} else {
		goto L60
	}
L19:
	;
	goto L18
L20:
	;
	v148 = v144 + v141
	if base.Ui32(int32(0)-v144<<(uint(int32(1))%32)) < base.Ui32(v135-v148) {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	v158 = (v135 ^ v141) & int32(3)
	if base.Ui32(v135) <= base.Ui32(v141) {
		goto L25
	} else {
		goto L26
	}
L22:
	;
	v155 = F___memcpy(m, v141, v135, v144)
	mBase = m.M
	goto L18
L23:
	;
	if v264 == int32(0) {
		goto L19
	} else {
		goto L55
	}
L24:
	;
	if base.Ui32(v242) <= base.Ui32(int32(3)) {
		v263 = v241
		v264 = v242
		v265 = v243
		goto L23
	} else {
		goto L51
	}
L25:
	;
	if v158 != 0 {
		v224 = v144
		goto L35
	} else {
		goto L36
	}
L26:
	;
	if v158 == int32(0) {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	if v141&int32(3) != 0 {
		goto L29
	} else {
		goto L30
	}
L28:
	;
	v263 = v135
	v264 = v144
	v265 = v141
	goto L23
L29:
	;
	v165 = v135
	v166 = v144
	v167 = v141
	goto L31
L30:
	;
	v241 = v135
	v242 = v144
	v243 = v141
	goto L24
L31:
	;
	if v166 == int32(0) {
		goto L19
	} else {
		goto L33
	}
L33:
	;
	v171 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v165))))
	*(*uint8)(unsafe.Add(mBase, uint32(v167))) = uint8(v171)
	v173 = int32(1)
	v174 = v165 + v173
	v176 = v166 + int32(-1)
	v178 = v167 + v173
	if v178&int32(3) == int32(0) {
		v241 = v174
		v242 = v176
		v243 = v178
		goto L24
	} else {
		goto L34
	}
L34:
	;
	v165 = v174
	v166 = v176
	v167 = v178
	goto L31
L35:
	;
	if v224 == int32(0) {
		goto L19
	} else {
		goto L47
	}
L36:
	;
	if v148&int32(3) == int32(0) {
		v204 = v144
		goto L37
	} else {
		goto L38
	}
L37:
	;
	if base.Ui32(v204) <= base.Ui32(int32(3)) {
		v224 = v204
		goto L35
	} else {
		goto L43
	}
L38:
	;
	v189 = v144
	goto L39
L39:
	;
	if v189 == int32(0) {
		goto L19
	} else {
		goto L41
	}
L40:
	;
	v204 = v195
	goto L37
L41:
	;
	v195 = v189 + int32(-1)
	v196 = v141 + v195
	v198 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v135+v195))))
	*(*uint8)(unsafe.Add(mBase, uint32(v196))) = uint8(v198)
	if v196&int32(3) != 0 {
		v189 = v195
		goto L39
	} else {
		goto L42
	}
L42:
	;
	goto L40
L43:
	;
	v211 = v204
	goto L44
L44:
	;
	v215 = v211 + int32(-4)
	v218 = *(*int32)(unsafe.Add(mBase, uint32(v135+v215)))
	*(*int32)(unsafe.Add(mBase, uint32(v141+v215))) = v218
	if base.Ui32(int32(3)) < base.Ui32(v215) {
		v211 = v215
		goto L44
	} else {
		goto L46
	}
L45:
	;
	v224 = v215
	goto L35
L46:
	;
	goto L45
L47:
	;
	v231 = v224
	goto L48
L48:
	;
	v235 = v231 + int32(-1)
	v238 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v135+v235))))
	*(*uint8)(unsafe.Add(mBase, uint32(v141+v235))) = uint8(v238)
	if v235 != 0 {
		v231 = v235
		goto L48
	} else {
		goto L50
	}
L50:
	;
	goto L19
L51:
	;
	v248 = v241
	v249 = v242
	v250 = v243
	goto L52
L52:
	;
	v252 = *(*int32)(unsafe.Add(mBase, uint32(v248)))
	*(*int32)(unsafe.Add(mBase, uint32(v250))) = v252
	v254 = int32(4)
	v255 = v248 + v254
	v257 = v250 + v254
	v259 = v249 + int32(-4)
	if base.Ui32(int32(3)) < base.Ui32(v259) {
		v248 = v255
		v249 = v259
		v250 = v257
		goto L52
	} else {
		goto L54
	}
L53:
	;
	v263 = v255
	v264 = v259
	v265 = v257
	goto L23
L54:
	;
	goto L53
L55:
	;
	v270 = v263
	v271 = v264
	v272 = v265
	goto L56
L56:
	;
	v274 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v270))))
	*(*uint8)(unsafe.Add(mBase, uint32(v272))) = uint8(v274)
	v276 = int32(1)
	v281 = v271 + int32(-1)
	if v281 != 0 {
		v270 = v270 + v276
		v271 = v281
		v272 = v272 + v276
		goto L56
	} else {
		goto L58
	}
L57:
	;
	goto L19
L58:
	;
	goto L57
L59:
	;
	v454 = v70 + v101
	v456 = v454 + int32(1)
	v457 = *(*int32)(unsafe.Add(mBase, uint32(v55)))
	v460 = int32(base.Ui32(v457)>>(uint(int32(3))%32)) - v101
	if v456 == v454 {
		goto L103
	} else {
		goto L104
	}
L60:
	;
	v295 = *(*int32)(unsafe.Add(mBase, uint32(v55)))
	v296 = int32(3)
	v297 = int32(base.Ui32(v295) >> (uint(v296) % 32))
	v303 = v70 + v297 + (int32(0)-v297)&v296
	v304 = v303 + v138
	if v304 == v303 {
		goto L62
	} else {
		goto L63
	}
L61:
	;
	goto L59
L62:
	;
	goto L61
L63:
	;
	v308 = v134 + v304
	if base.Ui32(int32(0)-v134<<(uint(int32(1))%32)) < base.Ui32(v303-v308) {
		goto L64
	} else {
		goto L65
	}
L64:
	;
	v318 = (v303 ^ v304) & int32(3)
	if base.Ui32(v303) <= base.Ui32(v304) {
		goto L68
	} else {
		goto L69
	}
L65:
	;
	v315 = F___memcpy(m, v304, v303, v134)
	mBase = m.M
	goto L61
L66:
	;
	if v424 == int32(0) {
		goto L62
	} else {
		goto L98
	}
L67:
	;
	if base.Ui32(v402) <= base.Ui32(int32(3)) {
		v423 = v401
		v424 = v402
		v425 = v403
		goto L66
	} else {
		goto L94
	}
L68:
	;
	if v318 != 0 {
		v384 = v134
		goto L78
	} else {
		goto L79
	}
L69:
	;
	if v318 == int32(0) {
		goto L70
	} else {
		goto L71
	}
L70:
	;
	if v304&int32(3) != 0 {
		goto L72
	} else {
		goto L73
	}
L71:
	;
	v423 = v303
	v424 = v134
	v425 = v304
	goto L66
L72:
	;
	v325 = v303
	v326 = v134
	v327 = v304
	goto L74
L73:
	;
	v401 = v303
	v402 = v134
	v403 = v304
	goto L67
L74:
	;
	if v326 == int32(0) {
		goto L62
	} else {
		goto L76
	}
L76:
	;
	v331 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v325))))
	*(*uint8)(unsafe.Add(mBase, uint32(v327))) = uint8(v331)
	v333 = int32(1)
	v334 = v325 + v333
	v336 = v326 + int32(-1)
	v338 = v327 + v333
	if v338&int32(3) == int32(0) {
		v401 = v334
		v402 = v336
		v403 = v338
		goto L67
	} else {
		goto L77
	}
L77:
	;
	v325 = v334
	v326 = v336
	v327 = v338
	goto L74
L78:
	;
	if v384 == int32(0) {
		goto L62
	} else {
		goto L90
	}
L79:
	;
	if v308&int32(3) == int32(0) {
		v364 = v134
		goto L80
	} else {
		goto L81
	}
L80:
	;
	if base.Ui32(v364) <= base.Ui32(int32(3)) {
		v384 = v364
		goto L78
	} else {
		goto L86
	}
L81:
	;
	v349 = v134
	goto L82
L82:
	;
	if v349 == int32(0) {
		goto L62
	} else {
		goto L84
	}
L83:
	;
	v364 = v355
	goto L80
L84:
	;
	v355 = v349 + int32(-1)
	v356 = v304 + v355
	v358 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v303+v355))))
	*(*uint8)(unsafe.Add(mBase, uint32(v356))) = uint8(v358)
	if v356&int32(3) != 0 {
		v349 = v355
		goto L82
	} else {
		goto L85
	}
L85:
	;
	goto L83
L86:
	;
	v371 = v364
	goto L87
L87:
	;
	v375 = v371 + int32(-4)
	v378 = *(*int32)(unsafe.Add(mBase, uint32(v303+v375)))
	*(*int32)(unsafe.Add(mBase, uint32(v304+v375))) = v378
	if base.Ui32(int32(3)) < base.Ui32(v375) {
		v371 = v375
		goto L87
	} else {
		goto L89
	}
L88:
	;
	v384 = v375
	goto L78
L89:
	;
	goto L88
L90:
	;
	v391 = v384
	goto L91
L91:
	;
	v395 = v391 + int32(-1)
	v398 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v303+v395))))
	*(*uint8)(unsafe.Add(mBase, uint32(v304+v395))) = uint8(v398)
	if v395 != 0 {
		v391 = v395
		goto L91
	} else {
		goto L93
	}
L93:
	;
	goto L62
L94:
	;
	v408 = v401
	v409 = v402
	v410 = v403
	goto L95
L95:
	;
	v412 = *(*int32)(unsafe.Add(mBase, uint32(v408)))
	*(*int32)(unsafe.Add(mBase, uint32(v410))) = v412
	v414 = int32(4)
	v415 = v408 + v414
	v417 = v410 + v414
	v419 = v409 + int32(-4)
	if base.Ui32(int32(3)) < base.Ui32(v419) {
		v408 = v415
		v409 = v419
		v410 = v417
		goto L95
	} else {
		goto L97
	}
L96:
	;
	v423 = v415
	v424 = v419
	v425 = v417
	goto L66
L97:
	;
	goto L96
L98:
	;
	v430 = v423
	v431 = v424
	v432 = v425
	goto L99
L99:
	;
	v434 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v430))))
	*(*uint8)(unsafe.Add(mBase, uint32(v432))) = uint8(v434)
	v436 = int32(1)
	v441 = v431 + int32(-1)
	if v441 != 0 {
		v430 = v430 + v436
		v431 = v441
		v432 = v432 + v436
		goto L99
	} else {
		goto L101
	}
L100:
	;
	goto L62
L101:
	;
	goto L100
L102:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v454))) = uint8(v2)
	v613 = v457&int32(-8) + int32(8)
	*(*int32)(unsafe.Add(mBase, uint32(v55))) = v613 | v457&int32(7)
	v618 = int32(3)
	v619 = int32(base.Ui32(v613) >> (uint(v618) % 32))
	v626 = v70 + v619 + (int32(0)-v619)&v618 + v134
	*(*int32)(unsafe.Add(mBase, uint32(v626))) = v18
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v18
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v626
	return v55
L103:
	;
	goto L102
L104:
	;
	v464 = v460 + v456
	if base.Ui32(int32(0)-v460<<(uint(int32(1))%32)) < base.Ui32(v454-v464) {
		goto L105
	} else {
		goto L106
	}
L105:
	;
	v474 = (v454 ^ v456) & int32(3)
	if base.Ui32(v454) <= base.Ui32(v456) {
		goto L109
	} else {
		goto L110
	}
L106:
	;
	v471 = F___memcpy(m, v456, v454, v460)
	mBase = m.M
	goto L102
L107:
	;
	if v580 == int32(0) {
		goto L103
	} else {
		goto L139
	}
L108:
	;
	if base.Ui32(v558) <= base.Ui32(int32(3)) {
		v579 = v557
		v580 = v558
		v581 = v559
		goto L107
	} else {
		goto L135
	}
L109:
	;
	if v474 != 0 {
		v540 = v460
		goto L119
	} else {
		goto L120
	}
L110:
	;
	if v474 == int32(0) {
		goto L111
	} else {
		goto L112
	}
L111:
	;
	if v456&int32(3) != 0 {
		goto L113
	} else {
		goto L114
	}
L112:
	;
	v579 = v454
	v580 = v460
	v581 = v456
	goto L107
L113:
	;
	v481 = v454
	v482 = v460
	v483 = v456
	goto L115
L114:
	;
	v557 = v454
	v558 = v460
	v559 = v456
	goto L108
L115:
	;
	if v482 == int32(0) {
		goto L103
	} else {
		goto L117
	}
L117:
	;
	v487 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v481))))
	*(*uint8)(unsafe.Add(mBase, uint32(v483))) = uint8(v487)
	v489 = int32(1)
	v490 = v481 + v489
	v492 = v482 + int32(-1)
	v494 = v483 + v489
	if v494&int32(3) == int32(0) {
		v557 = v490
		v558 = v492
		v559 = v494
		goto L108
	} else {
		goto L118
	}
L118:
	;
	v481 = v490
	v482 = v492
	v483 = v494
	goto L115
L119:
	;
	if v540 == int32(0) {
		goto L103
	} else {
		goto L131
	}
L120:
	;
	if v464&int32(3) == int32(0) {
		v520 = v460
		goto L121
	} else {
		goto L122
	}
L121:
	;
	if base.Ui32(v520) <= base.Ui32(int32(3)) {
		v540 = v520
		goto L119
	} else {
		goto L127
	}
L122:
	;
	v505 = v460
	goto L123
L123:
	;
	if v505 == int32(0) {
		goto L103
	} else {
		goto L125
	}
L124:
	;
	v520 = v511
	goto L121
L125:
	;
	v511 = v505 + int32(-1)
	v512 = v456 + v511
	v514 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v454+v511))))
	*(*uint8)(unsafe.Add(mBase, uint32(v512))) = uint8(v514)
	if v512&int32(3) != 0 {
		v505 = v511
		goto L123
	} else {
		goto L126
	}
L126:
	;
	goto L124
L127:
	;
	v527 = v520
	goto L128
L128:
	;
	v531 = v527 + int32(-4)
	v534 = *(*int32)(unsafe.Add(mBase, uint32(v454+v531)))
	*(*int32)(unsafe.Add(mBase, uint32(v456+v531))) = v534
	if base.Ui32(int32(3)) < base.Ui32(v531) {
		v527 = v531
		goto L128
	} else {
		goto L130
	}
L129:
	;
	v540 = v531
	goto L119
L130:
	;
	goto L129
L131:
	;
	v547 = v540
	goto L132
L132:
	;
	v551 = v547 + int32(-1)
	v554 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v454+v551))))
	*(*uint8)(unsafe.Add(mBase, uint32(v456+v551))) = uint8(v554)
	if v551 != 0 {
		v547 = v551
		goto L132
	} else {
		goto L134
	}
L134:
	;
	goto L103
L135:
	;
	v564 = v557
	v565 = v558
	v566 = v559
	goto L136
L136:
	;
	v568 = *(*int32)(unsafe.Add(mBase, uint32(v564)))
	*(*int32)(unsafe.Add(mBase, uint32(v566))) = v568
	v570 = int32(4)
	v571 = v564 + v570
	v573 = v566 + v570
	v575 = v565 + int32(-4)
	if base.Ui32(int32(3)) < base.Ui32(v575) {
		v564 = v571
		v565 = v575
		v566 = v573
		goto L136
	} else {
		goto L138
	}
L137:
	;
	v579 = v571
	v580 = v575
	v581 = v573
	goto L107
L138:
	;
	goto L137
L139:
	;
	v586 = v579
	v587 = v580
	v588 = v581
	goto L140
L140:
	;
	v590 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v586))))
	*(*uint8)(unsafe.Add(mBase, uint32(v588))) = uint8(v590)
	v592 = int32(1)
	v597 = v587 + int32(-1)
	if v597 != 0 {
		v586 = v586 + v592
		v587 = v597
		v588 = v588 + v592
		goto L140
	} else {
		goto L142
	}
L141:
	;
	goto L103
L142:
	;
	goto L141
L143:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L144:
	;
	return int32(0)
}
func F_raxCompare(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v14 int32
	_ = v14
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v27 int32
	_ = v27
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
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v85 int32
	_ = v85
	var v100 int32
	_ = v100
	v10 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	if v10 == int32(61) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	v22 = base.B2i32(v10 == int32(62))
	if v10 == int32(62) {
		v32 = int32(1)
		goto L5
	} else {
		goto L6
	}
L2:
	;
	v19 = int32(1)
	goto L1
L3:
	;
	v14 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+1)))
	if v14 != int32(61) {
		v19 = int32(0)
		goto L1
	} else {
		goto L4
	}
L4:
	;
	goto L2
L5:
	;
	v33 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v35 = base.B2i32(base.Ui32(l3) < base.Ui32(v34))
	if base.Ui32(l3) < base.Ui32(v34) {
		goto L9
	} else {
		goto L10
	}
L6:
	;
	if v10 == int32(60) {
		v32 = int32(0)
		goto L5
	} else {
		goto L7
	}
L7:
	;
	v27 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+1)))
	if v27 == int32(61) {
		v32 = int32(1)
		goto L5
	} else {
		goto L8
	}
L8:
	;
	return int32(0)
L9:
	;
	v36 = l3
	goto L11
L10:
	;
	v36 = v34
	goto L11
L11:
	;
	if base.Ui32(v36) < base.Ui32(int32(4)) {
		v60 = v33
		v61 = l2
		v62 = v36
		goto L15
	} else {
		goto L16
	}
L12:
	;
	if v10 == int32(62) {
		goto L28
	} else {
		goto L29
	}
L13:
	;
	v100 = int32(0)
	goto L12
L14:
	;
	v72 = v67
	v73 = v68
	v74 = v69
	goto L24
L15:
	;
	if v62 == int32(0) {
		goto L13
	} else {
		goto L22
	}
L16:
	;
	if (l2|v33)&int32(3) != 0 {
		v67 = v33
		v68 = l2
		v69 = v36
		goto L14
	} else {
		goto L17
	}
L17:
	;
	v44 = v33
	v45 = l2
	v46 = v36
	goto L18
L18:
	;
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v44)))
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v45)))
	if v49 != v50 {
		v67 = v44
		v68 = v45
		v69 = v46
		goto L14
	} else {
		goto L20
	}
L19:
	;
	v60 = v55
	v61 = v53
	v62 = v57
	goto L15
L20:
	;
	v52 = int32(4)
	v53 = v45 + v52
	v55 = v44 + v52
	v57 = v46 + int32(-4)
	if base.Ui32(int32(3)) < base.Ui32(v57) {
		v44 = v55
		v45 = v53
		v46 = v57
		goto L18
	} else {
		goto L21
	}
L21:
	;
	goto L19
L22:
	;
	v67 = v60
	v68 = v61
	v69 = v62
	goto L14
L23:
	;
	v100 = v77 - v78
	goto L12
L24:
	;
	v77 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v72))))
	v78 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v73))))
	if v77 != v78 {
		goto L23
	} else {
		goto L26
	}
L26:
	;
	v80 = int32(1)
	v85 = v74 + int32(-1)
	if v85 == int32(0) {
		goto L13
	} else {
		goto L27
	}
L27:
	;
	v72 = v72 + v80
	v73 = v73 + v80
	v74 = v85
	goto L24
L28:
	;
	if v100 != 0 {
		goto L31
	} else {
		goto L32
	}
L29:
	;
	if v32 == int32(0) {
		goto L28
	} else {
		goto L30
	}
L30:
	;
	return base.B2i32(v100 == int32(0)) & base.B2i32(l3 == v34)
L31:
	;
	if v100 < int32(1) {
		goto L37
	} else {
		goto L38
	}
L32:
	;
	if v19&base.B2i32(l3 == v34) == int32(0) {
		goto L33
	} else {
		goto L34
	}
L33:
	;
	if v32 != 0 {
		goto L35
	} else {
		goto L36
	}
L34:
	;
	return int32(1)
L35:
	;
	return v35 & base.B2i32(v10 == int32(62))
L36:
	;
	return base.B2i32(base.Ui32(v34) < base.Ui32(l3))
L37:
	;
	return v32 ^ int32(1)
L38:
	;
	return base.B2i32(v10 == int32(62))
}
func F_raxEOF(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	return v2 & int32(2)
}
func F_raxFind(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v41 int32
	_ = v41
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v102 int32
	_ = v102
	var v107 int32
	_ = v107
	var v114 int32
	_ = v114
	var v119 int32
	_ = v119
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v141 int32
	_ = v141
	var v144 int32
	_ = v144
	var v148 int32
	_ = v148
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v157 int32
	_ = v157
	var v161 int32
	_ = v161
	var v163 int32
	_ = v163
	var v169 int32
	_ = v169
	var v173 int32
	_ = v173
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v190 int32
	_ = v190
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v198 int32
	_ = v198
	v5 = int32(0)
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
	if l2 == v5 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	if v157 != l2 {
		v198 = v5
		goto L28
	} else {
		goto L29
	}
L2:
	;
	v148 = int32(0)
	v154 = v13
	v155 = v14
	v157 = v148
	v161 = v148
	goto L1
L3:
	;
	if base.Ui32(v14) < base.Ui32(int32(8)) {
		goto L2
	} else {
		goto L4
	}
L4:
	;
	v25 = v13
	v26 = v14
	v28 = int32(0)
	goto L6
L5:
	;
	v154 = v138
	v155 = v139
	v157 = v141
	v161 = base.B2i32(v144 != int32(0))
	goto L1
L6:
	;
	v34 = int32(base.Ui32(v26) >> (uint(int32(3)) % 32))
	v35 = int32(4)
	v36 = v25 + v35
	if v26&v35 == int32(0) {
		goto L9
	} else {
		goto L10
	}
L7:
	;
	v138 = v129
	v139 = v130
	v141 = v114
	v144 = v119
	goto L5
L8:
	;
	v119 = int32(0)
	v129 = *(*int32)(unsafe.Add(mBase, uint32(v36+v34+(v119-v34)&int32(3)+v107<<(uint(int32(2))%32))))
	v130 = *(*int32)(unsafe.Add(mBase, uint32(v129)))
	if base.Ui32(v130) < base.Ui32(int32(8)) {
		v138 = v129
		v139 = v130
		v141 = v114
		v144 = v119
		goto L5
	} else {
		goto L26
	}
L9:
	;
	v82 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1+v28))))
	v85 = int32(0)
	goto L20
L10:
	;
	v41 = int32(0)
	if base.Ui32(l2) <= base.Ui32(v28) {
		v74 = v28
		v77 = v41
		goto L11
	} else {
		goto L12
	}
L11:
	;
	if v77 == v34 {
		v107 = v41
		v114 = v74
		goto L8
	} else {
		goto L18
	}
L12:
	;
	v51 = v28
	v54 = v41
	goto L13
L13:
	;
	v57 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v36+v54))))
	v59 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1+v51))))
	if v57 != v59 {
		v74 = v51
		v77 = v54
		goto L11
	} else {
		goto L15
	}
L14:
	;
	v74 = v62
	v77 = v64
	goto L11
L15:
	;
	v61 = int32(1)
	v62 = v51 + v61
	v64 = v54 + v61
	if base.Ui32(v34) <= base.Ui32(v64) {
		v74 = v62
		v77 = v64
		goto L11
	} else {
		goto L16
	}
L16:
	;
	if base.Ui32(v62) < base.Ui32(l2) {
		v51 = v62
		v54 = v64
		goto L13
	} else {
		goto L17
	}
L17:
	;
	goto L14
L18:
	;
	v138 = v25
	v139 = v26
	v141 = v74
	v144 = v77
	goto L5
L19:
	;
	if v85 != v34 {
		goto L24
	} else {
		goto L25
	}
L20:
	;
	v98 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v36+v85))))
	if v98 == v82&int32(255) {
		goto L19
	} else {
		goto L22
	}
L22:
	;
	v100 = int32(1)
	v102 = v85 + v100
	if v102 != v34 {
		v85 = v102
		goto L20
	} else {
		goto L23
	}
L23:
	;
	v154 = v25
	v155 = v26
	v157 = v28
	v161 = v100
	goto L1
L24:
	;
	v107 = v85
	v114 = v28 + int32(1)
	goto L8
L25:
	;
	v138 = v25
	v139 = v26
	v141 = v28
	v144 = v34
	goto L5
L26:
	;
	if base.Ui32(v114) < base.Ui32(l2) {
		v25 = v129
		v26 = v130
		v28 = v114
		goto L6
	} else {
		goto L27
	}
L27:
	;
	goto L7
L28:
	;
	return v198
L29:
	;
	v163 = int32(0)
	if v155&int32(1) == v163 {
		v198 = v163
		goto L28
	} else {
		goto L30
	}
L30:
	;
	v169 = v155 & int32(4)
	if v161&base.B2i32(v169 != int32(0)) != 0 {
		v198 = v163
		goto L28
	} else {
		goto L31
	}
L31:
	;
	v173 = int32(1)
	if l3 == int32(0) {
		v198 = v173
		goto L28
	} else {
		goto L32
	}
L32:
	;
	if v155&int32(2) != 0 {
		v195 = int32(0)
		goto L33
	} else {
		goto L34
	}
L33:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v195
	v198 = v173
	goto L28
L34:
	;
	v179 = int32(3)
	v180 = int32(base.Ui32(v155) >> (uint(v179) % 32))
	if v169 != 0 {
		goto L35
	} else {
		goto L36
	}
L35:
	;
	v190 = int32(4)
	goto L37
L36:
	;
	v190 = v180 << (uint(int32(2)) % 32)
	goto L37
L37:
	;
	v194 = *(*int32)(unsafe.Add(mBase, uint32(v154+v180+(int32(0)-v180)&v179+v190+int32(4))))
	v195 = v194
	goto L33
}
func F_raxFree(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v5 int32
	_ = v5
	var v6 int64
	_ = v6
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	F_raxRecursiveFree(m, l0, v2, int32(0))
	mBase = m.M
	v5 = m.ExcPending
	if v5 != 0 {
		return
	} else {
		v6 = *(*int64)(unsafe.Add(mBase, uint32(l0)+16))
		if v6 == int64(0) {
			F_valkey_free(m, l0)
			mBase = m.M
			v16 = m.ExcPending
			if v16 != 0 {
				return
			} else {
				return
			}
		} else {
			F__serverAssert(m, int32(_a_F_raxFree_0), int32(_a_F_raxFree_1), int32(1236))
			mBase = m.M
			v13 = m.ExcPending
			if v13 != 0 {
				return
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
func F_raxIteratorNextStep(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v20 int32
	_ = v20
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v51 int32
	_ = v51
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v93 int32
	_ = v93
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
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
	var v122 int32
	_ = v122
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v129 int32
	_ = v129
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
	var v143 int32
	_ = v143
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v151 int32
	_ = v151
	var v154 int32
	_ = v154
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v163 int32
	_ = v163
	var v167 int32
	_ = v167
	var v174 int32
	_ = v174
	var v177 int32
	_ = v177
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v190 int32
	_ = v190
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v195 int32
	_ = v195
	var v197 int32
	_ = v197
	var v208 int32
	_ = v208
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v217 int32
	_ = v217
	var v223 int32
	_ = v223
	var v230 int32
	_ = v230
	var v234 int32
	_ = v234
	var v237 int32
	_ = v237
	var v243 int32
	_ = v243
	var v250 int32
	_ = v250
	var v254 int32
	_ = v254
	var v257 int32
	_ = v257
	var v260 int32
	_ = v260
	var v261 int32
	_ = v261
	var v262 int32
	_ = v262
	var v267 int32
	_ = v267
	var v268 int32
	_ = v268
	var v269 int32
	_ = v269
	var v271 int32
	_ = v271
	var v273 int32
	_ = v273
	var v274 int32
	_ = v274
	var v276 int32
	_ = v276
	var v278 int32
	_ = v278
	var v282 int32
	_ = v282
	var v283 int32
	_ = v283
	var v284 int32
	_ = v284
	var v289 int32
	_ = v289
	var v290 int32
	_ = v290
	var v291 int32
	_ = v291
	var v293 int32
	_ = v293
	var v295 int32
	_ = v295
	var v300 int32
	_ = v300
	var v312 int32
	_ = v312
	var v319 int32
	_ = v319
	var v323 int32
	_ = v323
	var v324 int32
	_ = v324
	var v326 int32
	_ = v326
	var v329 int32
	_ = v329
	var v330 int32
	_ = v330
	var v331 int32
	_ = v331
	var v332 int32
	_ = v332
	var v334 int32
	_ = v334
	var v335 int32
	_ = v335
	var v336 int32
	_ = v336
	var v344 int32
	_ = v344
	var v345 int32
	_ = v345
	var v352 int32
	_ = v352
	var v357 int32
	_ = v357
	var v361 int32
	_ = v361
	var v362 int32
	_ = v362
	var v371 int32
	_ = v371
	var v375 int32
	_ = v375
	var v376 int32
	_ = v376
	var v386 int32
	_ = v386
	var v387 int32
	_ = v387
	var v392 int32
	_ = v392
	var v396 int32
	_ = v396
	var v397 int32
	_ = v397
	var v401 int32
	_ = v401
	var v402 int32
	_ = v402
	var v403 int32
	_ = v403
	var v405 int32
	_ = v405
	var v407 int32
	_ = v407
	var v408 int32
	_ = v408
	var v410 int32
	_ = v410
	var v412 int32
	_ = v412
	var v416 int32
	_ = v416
	var v418 int32
	_ = v418
	var v420 int32
	_ = v420
	var v421 int32
	_ = v421
	var v422 int32
	_ = v422
	var v423 int32
	_ = v423
	var v424 int32
	_ = v424
	var v425 int32
	_ = v425
	var v431 int32
	_ = v431
	var v432 int32
	_ = v432
	var v434 int32
	_ = v434
	var v437 int32
	_ = v437
	var v439 int32
	_ = v439
	var v443 int32
	_ = v443
	var v445 int32
	_ = v445
	var v456 int32
	_ = v456
	var v457 int32
	_ = v457
	var v472 int32
	_ = v472
	var v473 int32
	_ = v473
	var v478 int32
	_ = v478
	var v482 int32
	_ = v482
	var v484 int32
	_ = v484
	var v488 int32
	_ = v488
	var v490 int32
	_ = v490
	var v491 int32
	_ = v491
	var v492 int32
	_ = v492
	var v496 int32
	_ = v496
	var v499 int32
	_ = v499
	var v501 int32
	_ = v501
	var v502 int32
	_ = v502
	var v504 int32
	_ = v504
	var v507 int32
	_ = v507
	var v508 int32
	_ = v508
	var v510 int32
	_ = v510
	var v512 int32
	_ = v512
	var v516 int32
	_ = v516
	var v523 int32
	_ = v523
	var v532 int32
	_ = v532
	var v533 int32
	_ = v533
	var v534 int32
	_ = v534
	var v535 int32
	_ = v535
	var v538 int32
	_ = v538
	var v540 int32
	_ = v540
	var v541 int32
	_ = v541
	var v551 int32
	_ = v551
	var v552 int32
	_ = v552
	var v556 int32
	_ = v556
	var v564 int32
	_ = v564
	var v566 int32
	_ = v566
	var v569 int32
	_ = v569
	var v571 int32
	_ = v571
	var v572 int32
	_ = v572
	var v576 int32
	_ = v576
	var v577 int32
	_ = v577
	var v578 int32
	_ = v578
	var v586 int32
	_ = v586
	var v588 int32
	_ = v588
	var v591 int32
	_ = v591
	var v592 int32
	_ = v592
	var v593 int32
	_ = v593
	var v594 int32
	_ = v594
	var v596 int32
	_ = v596
	var v597 int32
	_ = v597
	var v598 int32
	_ = v598
	var v606 int32
	_ = v606
	var v607 int32
	_ = v607
	var v614 int32
	_ = v614
	var v619 int32
	_ = v619
	var v623 int32
	_ = v623
	var v624 int32
	_ = v624
	var v647 int32
	_ = v647
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v20&int32(2) != 0 {
		v647 = int32(1)
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return v647
L2:
	;
	if v20&int32(1) == int32(0) {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v33 = l0 + int32(168)
	v35 = l0 + int32(24)
	v37 = l0 + int32(152)
	v40 = *(*int32)(unsafe.Add(mBase, uint32(l0)+160))
	v41 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v42 = *(*int32)(unsafe.Add(mBase, uint32(l0)+152))
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v42)))
	v51 = base.B2i32(l1 == int32(0))
	v55 = v43
	v56 = v42
	goto L6
L4:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v20 & int32(-4)
	return int32(1)
L5:
	;
	v647 = int32(1)
	goto L1
L6:
	;
	if v51&int32(1) == int32(0) {
		goto L11
	} else {
		goto L12
	}
L7:
	;
	if v598&int32(2) != 0 {
		v624 = int32(0)
		goto L161
	} else {
		goto L162
	}
L8:
	;
	v532 = *(*int32)(unsafe.Add(mBase, uint32(l0)+156))
	v533 = *(*int32)(unsafe.Add(mBase, uint32(l0)+152))
	v534 = *(*int32)(unsafe.Add(mBase, uint32(l0)+160))
	v535 = *(*int32)(unsafe.Add(mBase, uint32(l0)+164))
	if v534 != v535 {
		v577 = v534
		v578 = v532
		goto L139
	} else {
		goto L140
	}
L9:
	;
	if v488 != 0 {
		goto L135
	} else {
		goto L136
	}
L10:
	;
	if v143 != 0 {
		goto L131
	} else {
		goto L132
	}
L11:
	;
	v371 = v51
	v375 = v55
	v376 = v56
	goto L101
L12:
	;
	if base.Ui32(v55) < base.Ui32(int32(4)) {
		goto L11
	} else {
		goto L13
	}
L13:
	;
	v68 = *(*int32)(unsafe.Add(mBase, uint32(l0)+156))
	v69 = *(*int32)(unsafe.Add(mBase, uint32(l0)+160))
	v70 = *(*int32)(unsafe.Add(mBase, uint32(l0)+164))
	if v69 == v70 {
		goto L15
	} else {
		goto L16
	}
L14:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v116+v117<<(uint(int32(2))%32)))) = v56
	v122 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+160)) = v117 + v122
	v125 = int32(4)
	v126 = v115 + v125
	v127 = *(*int32)(unsafe.Add(mBase, uint32(v115)))
	v129 = int32(base.Ui32(v127) >> (uint(int32(3)) % 32))
	if v127&v125 != 0 {
		goto L33
	} else {
		goto L34
	}
L15:
	;
	v73 = v69 << (uint(int32(3)) % 32)
	if v68 != v33 {
		goto L19
	} else {
		goto L20
	}
L16:
	;
	v115 = v56
	v116 = v68
	v117 = v69
	goto L14
L17:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+164)) = v109 << (uint(int32(1)) % 32)
	v113 = *(*int32)(unsafe.Add(mBase, uint32(l0)+152))
	v114 = *(*int32)(unsafe.Add(mBase, uint32(l0)+160))
	v115 = v113
	v116 = v108
	v117 = v114
	goto L14
L18:
	;
	v101 = *(*int32)(unsafe.Add(mBase, uint32(l0)+164))
	v103 = v101 << (uint(int32(2)) % 32)
	if v103 == int32(0) {
		goto L30
	} else {
		goto L31
	}
L19:
	;
	v88 = F_valkey_realloc(m, v68, v73)
	mBase = m.M
	v89 = m.ExcPending
	if v89 != 0 {
		goto L21
	} else {
		goto L26
	}
L20:
	;
	v75 = F_valkey_malloc(m, v73)
	mBase = m.M
	v78 = m.ExcPending
	if v78 != 0 {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	return int32(0)
L22:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+156)) = v75
	if v75 != 0 {
		goto L18
	} else {
		goto L23
	}
L23:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+296)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+156)) = v33
	goto L24
L24:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_raxIteratorNextStep[0])) = int32(48)
	return int32(0)
L25:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+296)) = int32(1)
	goto L28
L26:
	;
	if v88 == int32(0) {
		goto L25
	} else {
		goto L27
	}
L27:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+156)) = v88
	v93 = *(*int32)(unsafe.Add(mBase, uint32(l0)+164))
	v108 = v88
	v109 = v93
	goto L17
L28:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_raxIteratorNextStep[0])) = int32(48)
	return int32(0)
L29:
	;
	v108 = v75
	v109 = v101
	goto L17
L30:
	;
	goto L29
L31:
	;
	v106 = F__emscripten_memcpy_bulkmem(m, v75, v33, v103)
	mBase = m.M
	goto L30
L32:
	;
	v319 = int32(0)
	v323 = v126 + v129 + (v319-v129)&int32(3)
	v324 = *(*int32)(unsafe.Add(mBase, uint32(v323)))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+152)) = v324
	v326 = *(*int32)(unsafe.Add(mBase, uint32(l0)+300))
	if v326 == v319 {
		v334 = v324
		goto L90
	} else {
		goto L91
	}
L33:
	;
	v133 = v129
	goto L35
L34:
	;
	v133 = v122
	goto L35
L35:
	;
	if v133 == int32(0) {
		goto L32
	} else {
		goto L36
	}
L36:
	;
	v136 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v137 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v138 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v139 = v138 + v133
	if base.Ui32(v139) <= base.Ui32(v137) {
		v160 = v136
		v161 = v138
		goto L37
	} else {
		goto L38
	}
L37:
	;
	v163 = v160 + v161
	if v163 == v126 {
		goto L50
	} else {
		goto L51
	}
L38:
	;
	if v136 == v35 {
		goto L39
	} else {
		goto L40
	}
L39:
	;
	v143 = int32(0)
	goto L41
L40:
	;
	v143 = v136
	goto L41
L41:
	;
	v145 = v139 << (uint(int32(1)) % 32)
	v146 = F_valkey_realloc(m, v143, v145)
	mBase = m.M
	v147 = m.ExcPending
	if v147 != 0 {
		goto L21
	} else {
		goto L42
	}
L42:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v146
	if v146 == int32(0) {
		goto L10
	} else {
		goto L43
	}
L43:
	;
	if v143 != 0 {
		v157 = v146
		goto L44
	} else {
		goto L45
	}
L44:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v145
	v159 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v160 = v157
	v161 = v159
	goto L37
L45:
	;
	v151 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v151 == int32(0) {
		goto L47
	} else {
		goto L48
	}
L46:
	;
	v156 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v157 = v156
	goto L44
L47:
	;
	goto L46
L48:
	;
	v154 = F__emscripten_memcpy_bulkmem(m, v146, v35, v151)
	mBase = m.M
	goto L47
L49:
	;
	v312 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v312 + v133
	goto L32
L50:
	;
	goto L49
L51:
	;
	v167 = v133 + v163
	if base.Ui32(int32(0)-v133<<(uint(int32(1))%32)) < base.Ui32(v126-v167) {
		goto L52
	} else {
		goto L53
	}
L52:
	;
	v177 = (v126 ^ v163) & int32(3)
	if base.Ui32(v126) <= base.Ui32(v163) {
		goto L56
	} else {
		goto L57
	}
L53:
	;
	v174 = F___memcpy(m, v163, v126, v133)
	mBase = m.M
	goto L49
L54:
	;
	if v283 == int32(0) {
		goto L50
	} else {
		goto L86
	}
L55:
	;
	if base.Ui32(v261) <= base.Ui32(int32(3)) {
		v282 = v260
		v283 = v261
		v284 = v262
		goto L54
	} else {
		goto L82
	}
L56:
	;
	if v177 != 0 {
		v243 = v133
		goto L66
	} else {
		goto L67
	}
L57:
	;
	if v177 == int32(0) {
		goto L58
	} else {
		goto L59
	}
L58:
	;
	if v163&int32(3) != 0 {
		goto L60
	} else {
		goto L61
	}
L59:
	;
	v282 = v126
	v283 = v133
	v284 = v163
	goto L54
L60:
	;
	v184 = v126
	v185 = v133
	v186 = v163
	goto L62
L61:
	;
	v260 = v126
	v261 = v133
	v262 = v163
	goto L55
L62:
	;
	if v185 == int32(0) {
		goto L50
	} else {
		goto L64
	}
L64:
	;
	v190 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v184))))
	*(*uint8)(unsafe.Add(mBase, uint32(v186))) = uint8(v190)
	v192 = int32(1)
	v193 = v184 + v192
	v195 = v185 + int32(-1)
	v197 = v186 + v192
	if v197&int32(3) == int32(0) {
		v260 = v193
		v261 = v195
		v262 = v197
		goto L55
	} else {
		goto L65
	}
L65:
	;
	v184 = v193
	v185 = v195
	v186 = v197
	goto L62
L66:
	;
	if v243 == int32(0) {
		goto L50
	} else {
		goto L78
	}
L67:
	;
	if v167&int32(3) == int32(0) {
		v223 = v133
		goto L68
	} else {
		goto L69
	}
L68:
	;
	if base.Ui32(v223) <= base.Ui32(int32(3)) {
		v243 = v223
		goto L66
	} else {
		goto L74
	}
L69:
	;
	v208 = v133
	goto L70
L70:
	;
	if v208 == int32(0) {
		goto L50
	} else {
		goto L72
	}
L71:
	;
	v223 = v214
	goto L68
L72:
	;
	v214 = v208 + int32(-1)
	v215 = v163 + v214
	v217 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v126+v214))))
	*(*uint8)(unsafe.Add(mBase, uint32(v215))) = uint8(v217)
	if v215&int32(3) != 0 {
		v208 = v214
		goto L70
	} else {
		goto L73
	}
L73:
	;
	goto L71
L74:
	;
	v230 = v223
	goto L75
L75:
	;
	v234 = v230 + int32(-4)
	v237 = *(*int32)(unsafe.Add(mBase, uint32(v126+v234)))
	*(*int32)(unsafe.Add(mBase, uint32(v163+v234))) = v237
	if base.Ui32(int32(3)) < base.Ui32(v234) {
		v230 = v234
		goto L75
	} else {
		goto L77
	}
L76:
	;
	v243 = v234
	goto L66
L77:
	;
	goto L76
L78:
	;
	v250 = v243
	goto L79
L79:
	;
	v254 = v250 + int32(-1)
	v257 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v126+v254))))
	*(*uint8)(unsafe.Add(mBase, uint32(v163+v254))) = uint8(v257)
	if v254 != 0 {
		v250 = v254
		goto L79
	} else {
		goto L81
	}
L81:
	;
	goto L50
L82:
	;
	v267 = v260
	v268 = v261
	v269 = v262
	goto L83
L83:
	;
	v271 = *(*int32)(unsafe.Add(mBase, uint32(v267)))
	*(*int32)(unsafe.Add(mBase, uint32(v269))) = v271
	v273 = int32(4)
	v274 = v267 + v273
	v276 = v269 + v273
	v278 = v268 + int32(-4)
	if base.Ui32(int32(3)) < base.Ui32(v278) {
		v267 = v274
		v268 = v278
		v269 = v276
		goto L83
	} else {
		goto L85
	}
L84:
	;
	v282 = v274
	v283 = v278
	v284 = v276
	goto L54
L85:
	;
	goto L84
L86:
	;
	v289 = v282
	v290 = v283
	v291 = v284
	goto L87
L87:
	;
	v293 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v289))))
	*(*uint8)(unsafe.Add(mBase, uint32(v291))) = uint8(v293)
	v295 = int32(1)
	v300 = v290 + int32(-1)
	if v300 != 0 {
		v289 = v289 + v295
		v290 = v300
		v291 = v291 + v295
		goto L87
	} else {
		goto L89
	}
L88:
	;
	goto L50
L89:
	;
	goto L88
L90:
	;
	v335 = int32(1)
	v336 = *(*int32)(unsafe.Add(mBase, uint32(v334)))
	if v336&v335 == int32(0) {
		v51 = v335
		v55 = v336
		v56 = v334
		goto L6
	} else {
		goto L95
	}
L91:
	;
	v329 = m.T0[v326].(func(*base.Module, int32) int32)(m, v37)
	mBase = m.M
	v330 = m.ExcPending
	if v330 != 0 {
		goto L21
	} else {
		goto L93
	}
L92:
	;
	v332 = *(*int32)(unsafe.Add(mBase, uint32(v37)))
	*(*int32)(unsafe.Add(mBase, uint32(v323))) = v332
	v334 = v332
	goto L90
L93:
	;
	if v329 != 0 {
		goto L92
	} else {
		goto L94
	}
L94:
	;
	v331 = *(*int32)(unsafe.Add(mBase, uint32(v37)))
	v334 = v331
	goto L90
L95:
	;
	if v336&int32(2) != 0 {
		v362 = int32(0)
		goto L96
	} else {
		goto L97
	}
L96:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v362
	goto L5
L97:
	;
	v344 = int32(3)
	v345 = int32(base.Ui32(v336) >> (uint(v344) % 32))
	v352 = int32(4)
	if v336&v352 != 0 {
		goto L98
	} else {
		goto L99
	}
L98:
	;
	v357 = v352
	goto L100
L99:
	;
	v357 = v345 << (uint(int32(2)) % 32)
	goto L100
L100:
	;
	v361 = *(*int32)(unsafe.Add(mBase, uint32(v334+v345+(int32(0)-v345)&v344+v357+int32(4))))
	v362 = v361
	goto L96
L101:
	;
	if v371&int32(1) == int32(0) {
		goto L105
	} else {
		goto L106
	}
L102:
	;
	v482 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v484 = v432 + int32(1)
	if base.Ui32(v484) <= base.Ui32(v482) {
		v507 = v424
		v508 = v432
		goto L119
	} else {
		goto L120
	}
L103:
	;
	v425 = int32(1)
	if v421&int32(4) != 0 {
		goto L108
	} else {
		goto L109
	}
L104:
	;
	v402 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v403 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v405 = int32(-1)
	v407 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v402+v403+v405))))
	v408 = *(*int32)(unsafe.Add(mBase, uint32(l0)+160))
	v410 = v408 + v405
	*(*int32)(unsafe.Add(mBase, uint32(l0)+160)) = v410
	v412 = *(*int32)(unsafe.Add(mBase, uint32(l0)+156))
	v416 = *(*int32)(unsafe.Add(mBase, uint32(v412+v410<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+152)) = v416
	v418 = *(*int32)(unsafe.Add(mBase, uint32(v416)))
	v420 = v403
	v421 = v418
	v422 = v416
	v423 = v407
	v424 = v402
	goto L103
L105:
	;
	v396 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v397 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v401 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v396+v397+int32(-1)))))
	v420 = v397
	v421 = v375
	v422 = v376
	v423 = v401
	v424 = v396
	goto L103
L106:
	;
	v386 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v387 = *(*int32)(unsafe.Add(mBase, uint32(v386)))
	if v376 != v387 {
		goto L104
	} else {
		goto L107
	}
L107:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+160)) = v40
	*(*int32)(unsafe.Add(mBase, uint32(l0)+152)) = v42
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v41
	v392 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v392 | int32(2)
	goto L5
L108:
	;
	v431 = int32(base.Ui32(v421) >> (uint(int32(3)) % 32))
	goto L110
L109:
	;
	v431 = v425
	goto L110
L110:
	;
	v432 = v420 - v431
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v432
	v434 = *(*int32)(unsafe.Add(mBase, uint32(v422)))
	if v434&int32(4) != 0 {
		v371 = v425
		v375 = v434
		v376 = v422
		goto L101
	} else {
		goto L111
	}
L111:
	;
	v437 = int32(1)
	v439 = int32(base.Ui32(v434) >> (uint(int32(3)) % 32))
	if base.Ui32(v439) <= base.Ui32(v371&v437) {
		v371 = v437
		v375 = v434
		v376 = v422
		goto L101
	} else {
		goto L112
	}
L112:
	;
	v443 = int32(0)
	v445 = v422 + int32(4)
	v456 = v443
	v457 = v445 + v439 + (v443-v439)&int32(3)
	goto L114
L113:
	;
	if v456 == v439 {
		v371 = v437
		v375 = v434
		v376 = v422
		goto L101
	} else {
		goto L118
	}
L114:
	;
	v472 = v445 + v456
	v473 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v472))))
	if base.Ui32(v423&int32(255)) < base.Ui32(v473) {
		goto L113
	} else {
		goto L116
	}
L115:
	;
	v371 = int32(1)
	v375 = v434
	v376 = v422
	goto L101
L116:
	;
	v478 = v456 + int32(1)
	if v478 != v439 {
		v456 = v478
		v457 = v457 + int32(4)
		goto L114
	} else {
		goto L117
	}
L117:
	;
	goto L115
L118:
	;
	goto L102
L119:
	;
	v510 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v472))))
	*(*uint8)(unsafe.Add(mBase, uint32(v507+v508))) = uint8(v510)
	v512 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v512 + int32(1)
	goto L8
L120:
	;
	if v424 == v35 {
		goto L121
	} else {
		goto L122
	}
L121:
	;
	v488 = int32(0)
	goto L123
L122:
	;
	v488 = v424
	goto L123
L123:
	;
	v490 = v484 << (uint(int32(1)) % 32)
	v491 = F_valkey_realloc(m, v488, v490)
	mBase = m.M
	v492 = m.ExcPending
	if v492 != 0 {
		goto L21
	} else {
		goto L124
	}
L124:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v491
	if v491 == int32(0) {
		goto L9
	} else {
		goto L125
	}
L125:
	;
	if v488 != 0 {
		v502 = v491
		goto L126
	} else {
		goto L127
	}
L126:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v490
	v504 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v507 = v502
	v508 = v504
	goto L119
L127:
	;
	v496 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v496 == int32(0) {
		goto L129
	} else {
		goto L130
	}
L128:
	;
	v501 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v502 = v501
	goto L126
L129:
	;
	goto L128
L130:
	;
	v499 = F__emscripten_memcpy_bulkmem(m, v491, v35, v496)
	mBase = m.M
	goto L129
L131:
	;
	v516 = v143
	goto L133
L132:
	;
	v516 = v35
	goto L133
L133:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v516
	goto L134
L134:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_raxIteratorNextStep[0])) = int32(48)
	return int32(0)
L135:
	;
	v523 = v488
	goto L137
L136:
	;
	v523 = v35
	goto L137
L137:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v523
	goto L138
L138:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_raxIteratorNextStep[0])) = int32(48)
	goto L8
L139:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v578+v577<<(uint(int32(2))%32)))) = v533
	*(*int32)(unsafe.Add(mBase, uint32(l0)+160)) = v577 + int32(1)
	v586 = *(*int32)(unsafe.Add(mBase, uint32(v457)))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+152)) = v586
	v588 = *(*int32)(unsafe.Add(mBase, uint32(l0)+300))
	if v588 == int32(0) {
		v596 = v586
		goto L155
	} else {
		goto L156
	}
L140:
	;
	v538 = v534 << (uint(int32(3)) % 32)
	if v532 != v33 {
		goto L143
	} else {
		goto L144
	}
L141:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+164)) = v571 << (uint(int32(1)) % 32)
	v576 = *(*int32)(unsafe.Add(mBase, uint32(l0)+160))
	v577 = v576
	v578 = v572
	goto L139
L142:
	;
	v564 = *(*int32)(unsafe.Add(mBase, uint32(l0)+164))
	v566 = v564 << (uint(int32(2)) % 32)
	if v566 == int32(0) {
		goto L153
	} else {
		goto L154
	}
L143:
	;
	v551 = F_valkey_realloc(m, v532, v538)
	mBase = m.M
	v552 = m.ExcPending
	if v552 != 0 {
		goto L21
	} else {
		goto L149
	}
L144:
	;
	v540 = F_valkey_malloc(m, v538)
	mBase = m.M
	v541 = m.ExcPending
	if v541 != 0 {
		goto L21
	} else {
		goto L145
	}
L145:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+156)) = v540
	if v540 != 0 {
		goto L142
	} else {
		goto L146
	}
L146:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+296)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+156)) = v33
	goto L147
L147:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_raxIteratorNextStep[0])) = int32(48)
	return int32(0)
L148:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+296)) = int32(1)
	goto L151
L149:
	;
	if v551 == int32(0) {
		goto L148
	} else {
		goto L150
	}
L150:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+156)) = v551
	v556 = *(*int32)(unsafe.Add(mBase, uint32(l0)+164))
	v571 = v556
	v572 = v551
	goto L141
L151:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_raxIteratorNextStep[0])) = int32(48)
	return int32(0)
L152:
	;
	v571 = v564
	v572 = v540
	goto L141
L153:
	;
	goto L152
L154:
	;
	v569 = F__emscripten_memcpy_bulkmem(m, v540, v33, v566)
	mBase = m.M
	goto L153
L155:
	;
	v597 = int32(1)
	v598 = *(*int32)(unsafe.Add(mBase, uint32(v596)))
	if v598&v597 == int32(0) {
		v51 = v597
		v55 = v598
		v56 = v596
		goto L6
	} else {
		goto L160
	}
L156:
	;
	v591 = m.T0[v588].(func(*base.Module, int32) int32)(m, v37)
	mBase = m.M
	v592 = m.ExcPending
	if v592 != 0 {
		goto L21
	} else {
		goto L158
	}
L157:
	;
	v594 = *(*int32)(unsafe.Add(mBase, uint32(v37)))
	*(*int32)(unsafe.Add(mBase, uint32(v457))) = v594
	v596 = v594
	goto L155
L158:
	;
	if v591 != 0 {
		goto L157
	} else {
		goto L159
	}
L159:
	;
	v593 = *(*int32)(unsafe.Add(mBase, uint32(v37)))
	v596 = v593
	goto L155
L160:
	;
	goto L7
L161:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v624
	goto L5
L162:
	;
	v606 = int32(3)
	v607 = int32(base.Ui32(v598) >> (uint(v606) % 32))
	v614 = int32(4)
	if v598&v614 != 0 {
		goto L163
	} else {
		goto L164
	}
L163:
	;
	v619 = v614
	goto L165
L164:
	;
	v619 = v607 << (uint(int32(2)) % 32)
	goto L165
L165:
	;
	v623 = *(*int32)(unsafe.Add(mBase, uint32(v596+v607+(int32(0)-v607)&v606+v619+int32(4))))
	v624 = v623
	goto L161
}
func F_raxIteratorPrevStep(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v18 int32
	_ = v18
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v47 int32
	_ = v47
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v66 int32
	_ = v66
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
	var v115 int32
	_ = v115
	var v118 int32
	_ = v118
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v133 int32
	_ = v133
	var v135 int32
	_ = v135
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v174 int32
	_ = v174
	var v176 int32
	_ = v176
	var v180 int32
	_ = v180
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v186 int32
	_ = v186
	var v190 int32
	_ = v190
	var v193 int32
	_ = v193
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v207 int32
	_ = v207
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v217 int32
	_ = v217
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v235 int32
	_ = v235
	var v243 int32
	_ = v243
	var v245 int32
	_ = v245
	var v248 int32
	_ = v248
	var v250 int32
	_ = v250
	var v251 int32
	_ = v251
	var v255 int32
	_ = v255
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v265 int32
	_ = v265
	var v267 int32
	_ = v267
	var v268 int32
	_ = v268
	var v271 int32
	_ = v271
	var v272 int32
	_ = v272
	var v281 int32
	_ = v281
	var v286 int32
	_ = v286
	var v290 int32
	_ = v290
	var v298 int32
	_ = v298
	var v299 int32
	_ = v299
	var v306 int32
	_ = v306
	var v311 int32
	_ = v311
	var v315 int32
	_ = v315
	var v316 int32
	_ = v316
	var v337 int32
	_ = v337
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v18&int32(2) != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	if v180 != 0 {
		goto L65
	} else {
		goto L66
	}
L2:
	;
	return int32(1)
L3:
	;
	if v18&int32(1) == int32(0) {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v31 = l0 + int32(168)
	v33 = l0 + int32(24)
	v36 = *(*int32)(unsafe.Add(mBase, uint32(l0)+160))
	v37 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)+152))
	v41 = base.B2i32(l1 == int32(0))
	v47 = v38
	goto L6
L5:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v18 & int32(-4)
	return int32(1)
L6:
	;
	v57 = v41 & int32(1)
	if v57 == int32(0) {
		goto L10
	} else {
		goto L11
	}
L7:
	;
	if v286&int32(2) != 0 {
		v316 = int32(0)
		goto L60
	} else {
		goto L61
	}
L8:
	;
	v102 = *(*int32)(unsafe.Add(mBase, uint32(v99)))
	if v102&int32(4) != 0 {
		goto L16
	} else {
		goto L17
	}
L9:
	;
	v78 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v79 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v83 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v78+v79+int32(-1)))))
	v84 = *(*int32)(unsafe.Add(mBase, uint32(l0)+160))
	if v84 != 0 {
		goto L14
	} else {
		goto L15
	}
L10:
	;
	v72 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v73 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v77 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v72+v73+int32(-1)))))
	v98 = v73
	v99 = v47
	v100 = v72
	v101 = v77
	goto L8
L11:
	;
	v60 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v60)))
	if v47 != v61 {
		goto L9
	} else {
		goto L12
	}
L12:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+160)) = v36
	*(*int32)(unsafe.Add(mBase, uint32(l0)+152)) = v38
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v37
	v66 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v66 | int32(2)
	return int32(1)
L13:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+152)) = v95
	v98 = v79
	v99 = v95
	v100 = v78
	v101 = v83
	goto L8
L14:
	;
	v87 = v84 + int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+160)) = v87
	v89 = *(*int32)(unsafe.Add(mBase, uint32(l0)+156))
	v93 = *(*int32)(unsafe.Add(mBase, uint32(v89+v87<<(uint(int32(2))%32))))
	v95 = v93
	goto L13
L15:
	;
	v95 = int32(0)
	goto L13
L16:
	;
	v108 = int32(base.Ui32(v102) >> (uint(int32(3)) % 32))
	goto L18
L17:
	;
	v108 = int32(1)
	goto L18
L18:
	;
	v109 = v98 - v108
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v109
	v111 = *(*int32)(unsafe.Add(mBase, uint32(v99)))
	if v111&int32(4) != 0 {
		v281 = v99
		v286 = v111
		goto L19
	} else {
		goto L20
	}
L19:
	;
	v290 = int32(1)
	if v286&v290 == int32(0) {
		v41 = v290
		v47 = v281
		goto L6
	} else {
		goto L59
	}
L20:
	;
	v115 = int32(base.Ui32(v111) >> (uint(int32(3)) % 32))
	if base.Ui32(v115) <= base.Ui32(v57) {
		v281 = v99
		v286 = v111
		goto L19
	} else {
		goto L21
	}
L21:
	;
	v118 = int32(0)
	v126 = int32(1)
	v127 = v111 << (uint(v126) % 32)
	v133 = v118 - v111&v126
	v135 = int32(4)
	v149 = v99 + v115 + (v118-v115)&int32(3) + v115<<(uint(int32(2))%32) + (v127^int32(-1))&v133&v135 + v133&(v127|int32(-5)+v126)
	v150 = v115
	goto L23
L22:
	;
	v174 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v176 = v109 + int32(1)
	if base.Ui32(v176) <= base.Ui32(v174) {
		v202 = v100
		v203 = v168
		v204 = v109
		goto L27
	} else {
		goto L28
	}
L23:
	;
	v166 = v150 + int32(-1)
	v167 = v99 + v135 + v166
	v168 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v167))))
	if base.Ui32(v168) < base.Ui32(v101&int32(255)) {
		goto L22
	} else {
		goto L25
	}
L25:
	;
	if base.Ui32(int32(1)) < base.Ui32(v150) {
		v149 = v149 + int32(-4)
		v150 = v166
		goto L23
	} else {
		goto L26
	}
L26:
	;
	v281 = v99
	v286 = v111
	goto L19
L27:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v202+v204))) = uint8(v203)
	v207 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v207 + int32(1)
	v211 = *(*int32)(unsafe.Add(mBase, uint32(l0)+156))
	v212 = *(*int32)(unsafe.Add(mBase, uint32(l0)+152))
	v213 = *(*int32)(unsafe.Add(mBase, uint32(l0)+160))
	v214 = *(*int32)(unsafe.Add(mBase, uint32(l0)+164))
	if v213 != v214 {
		v256 = v211
		v257 = v213
		goto L40
	} else {
		goto L41
	}
L28:
	;
	if v100 == v33 {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	v180 = int32(0)
	goto L31
L30:
	;
	v180 = v100
	goto L31
L31:
	;
	v182 = v176 << (uint(int32(1)) % 32)
	v183 = F_valkey_realloc(m, v180, v182)
	mBase = m.M
	v186 = m.ExcPending
	if v186 != 0 {
		goto L32
	} else {
		goto L33
	}
L32:
	;
	return int32(0)
L33:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v183
	if v183 == int32(0) {
		goto L1
	} else {
		goto L34
	}
L34:
	;
	if v180 != 0 {
		v196 = v183
		goto L35
	} else {
		goto L36
	}
L35:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v182
	v198 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v199 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v167))))
	v202 = v196
	v203 = v199
	v204 = v198
	goto L27
L36:
	;
	v190 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v190 == int32(0) {
		goto L38
	} else {
		goto L39
	}
L37:
	;
	v195 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v196 = v195
	goto L35
L38:
	;
	goto L37
L39:
	;
	v193 = F__emscripten_memcpy_bulkmem(m, v183, v33, v190)
	mBase = m.M
	goto L38
L40:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v256+v257<<(uint(int32(2))%32)))) = v212
	*(*int32)(unsafe.Add(mBase, uint32(l0)+160)) = v257 + int32(1)
	v265 = *(*int32)(unsafe.Add(mBase, uint32(v149)))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+152)) = v265
	v267 = F_raxSeekGreatest(m, l0)
	mBase = m.M
	v268 = m.ExcPending
	if v268 != 0 {
		goto L32
	} else {
		goto L57
	}
L41:
	;
	v217 = v213 << (uint(int32(3)) % 32)
	if v211 != v31 {
		goto L44
	} else {
		goto L45
	}
L42:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+164)) = v251 << (uint(int32(1)) % 32)
	v255 = *(*int32)(unsafe.Add(mBase, uint32(l0)+160))
	v256 = v250
	v257 = v255
	goto L40
L43:
	;
	v243 = *(*int32)(unsafe.Add(mBase, uint32(l0)+164))
	v245 = v243 << (uint(int32(2)) % 32)
	if v245 == int32(0) {
		goto L54
	} else {
		goto L55
	}
L44:
	;
	v230 = F_valkey_realloc(m, v211, v217)
	mBase = m.M
	v231 = m.ExcPending
	if v231 != 0 {
		goto L32
	} else {
		goto L50
	}
L45:
	;
	v219 = F_valkey_malloc(m, v217)
	mBase = m.M
	v220 = m.ExcPending
	if v220 != 0 {
		goto L32
	} else {
		goto L46
	}
L46:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+156)) = v219
	if v219 != 0 {
		goto L43
	} else {
		goto L47
	}
L47:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+296)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+156)) = v31
	goto L48
L48:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_raxIteratorPrevStep[0])) = int32(48)
	return int32(0)
L49:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+296)) = int32(1)
	goto L52
L50:
	;
	if v230 == int32(0) {
		goto L49
	} else {
		goto L51
	}
L51:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+156)) = v230
	v235 = *(*int32)(unsafe.Add(mBase, uint32(l0)+164))
	v250 = v230
	v251 = v235
	goto L42
L52:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_raxIteratorPrevStep[0])) = int32(48)
	return int32(0)
L53:
	;
	v250 = v219
	v251 = v243
	goto L42
L54:
	;
	goto L53
L55:
	;
	v248 = F__emscripten_memcpy_bulkmem(m, v219, v31, v245)
	mBase = m.M
	goto L54
L56:
	;
	v271 = *(*int32)(unsafe.Add(mBase, uint32(l0)+152))
	v272 = *(*int32)(unsafe.Add(mBase, uint32(v271)))
	v281 = v271
	v286 = v272
	goto L19
L57:
	;
	if v267 != 0 {
		goto L56
	} else {
		goto L58
	}
L58:
	;
	return int32(0)
L59:
	;
	goto L7
L60:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v316
	goto L2
L61:
	;
	v298 = int32(3)
	v299 = int32(base.Ui32(v286) >> (uint(v298) % 32))
	v306 = int32(4)
	if v286&v306 != 0 {
		goto L62
	} else {
		goto L63
	}
L62:
	;
	v311 = v306
	goto L64
L63:
	;
	v311 = v299 << (uint(int32(2)) % 32)
	goto L64
L64:
	;
	v315 = *(*int32)(unsafe.Add(mBase, uint32(v281+v299+(int32(0)-v299)&v298+v311+int32(4))))
	v316 = v315
	goto L60
L65:
	;
	v337 = v180
	goto L67
L66:
	;
	v337 = v33
	goto L67
L67:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v337
	goto L68
L68:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_raxIteratorPrevStep[0])) = int32(48)
	return int32(0)
}
func F_raxRandomWalk(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v14 int32
	_ = v14
	var v15 int64
	_ = v15
	var v18 int32
	_ = v18
	var v25 float64
	_ = v25
	var v28 float64
	_ = v28
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v42 int64
	_ = v42
	var v46 int64
	_ = v46
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v79 int32
	_ = v79
	var v81 int64
	_ = v81
	var v85 int64
	_ = v85
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v100 int32
	_ = v100
	var v102 int32
	_ = v102
	var v104 int32
	_ = v104
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v118 int32
	_ = v118
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v139 int32
	_ = v139
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v145 int32
	_ = v145
	var v149 int32
	_ = v149
	var v152 int32
	_ = v152
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v161 int32
	_ = v161
	var v165 int32
	_ = v165
	var v172 int32
	_ = v172
	var v175 int32
	_ = v175
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v188 int32
	_ = v188
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v193 int32
	_ = v193
	var v195 int32
	_ = v195
	var v206 int32
	_ = v206
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v215 int32
	_ = v215
	var v221 int32
	_ = v221
	var v228 int32
	_ = v228
	var v232 int32
	_ = v232
	var v235 int32
	_ = v235
	var v241 int32
	_ = v241
	var v248 int32
	_ = v248
	var v252 int32
	_ = v252
	var v255 int32
	_ = v255
	var v258 int32
	_ = v258
	var v259 int32
	_ = v259
	var v260 int32
	_ = v260
	var v265 int32
	_ = v265
	var v266 int32
	_ = v266
	var v267 int32
	_ = v267
	var v269 int32
	_ = v269
	var v271 int32
	_ = v271
	var v272 int32
	_ = v272
	var v274 int32
	_ = v274
	var v276 int32
	_ = v276
	var v280 int32
	_ = v280
	var v281 int32
	_ = v281
	var v282 int32
	_ = v282
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
	var v298 int32
	_ = v298
	var v310 int32
	_ = v310
	var v317 int32
	_ = v317
	var v318 int32
	_ = v318
	var v319 int32
	_ = v319
	var v321 int32
	_ = v321
	var v325 int32
	_ = v325
	var v327 int32
	_ = v327
	var v328 int32
	_ = v328
	var v329 int32
	_ = v329
	var v333 int32
	_ = v333
	var v336 int32
	_ = v336
	var v338 int32
	_ = v338
	var v339 int32
	_ = v339
	var v341 int32
	_ = v341
	var v342 int32
	_ = v342
	var v343 int32
	_ = v343
	var v347 int32
	_ = v347
	var v350 int32
	_ = v350
	var v354 int32
	_ = v354
	var v361 int32
	_ = v361
	var v362 int32
	_ = v362
	var v363 int32
	_ = v363
	var v364 int32
	_ = v364
	var v367 int32
	_ = v367
	var v369 int32
	_ = v369
	var v370 int32
	_ = v370
	var v380 int32
	_ = v380
	var v381 int32
	_ = v381
	var v385 int32
	_ = v385
	var v393 int32
	_ = v393
	var v395 int32
	_ = v395
	var v398 int32
	_ = v398
	var v400 int32
	_ = v400
	var v401 int32
	_ = v401
	var v405 int32
	_ = v405
	var v406 int32
	_ = v406
	var v407 int32
	_ = v407
	var v408 int32
	_ = v408
	var v415 int32
	_ = v415
	var v416 int32
	_ = v416
	var v428 int32
	_ = v428
	var v429 int32
	_ = v429
	var v435 int32
	_ = v435
	var v437 int32
	_ = v437
	var v438 int32
	_ = v438
	var v443 int32
	_ = v443
	var v446 int32
	_ = v446
	var v447 int32
	_ = v447
	var v454 int32
	_ = v454
	var v459 int32
	_ = v459
	var v466 int32
	_ = v466
	var v467 int32
	_ = v467
	var v471 int32
	_ = v471
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v15 = *(*int64)(unsafe.Add(mBase, uint32(v14)+8))
	if v15 != int64(0) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	if l1 != 0 {
		v56 = l1
		goto L3
	} else {
		goto L4
	}
L2:
	;
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v18 | int32(2)
	return int32(0)
L3:
	;
	v60 = l0 + int32(168)
	v62 = l0 + int32(24)
	v63 = *(*int32)(unsafe.Add(mBase, uint32(l0)+152))
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v63)))
	v66 = v56
	v69 = v63
	v72 = v64
	goto L10
L4:
	;
	v25 = F_log(m, base.F64_convert_i64_u(v15))
	mBase = m.M
	v28 = base.F64_add(base.F64_floor(v25), float64(1))
	if base.F64_lt(v28, float64(4.294967296e+09))&base.F64_ge(v28, float64(0)) == int32(0) {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	v40 = int32(0)
	v42 = *(*int64)(unsafe.Add(mBase, _c_F_raxRandomWalk[0]))
	v46 = v42*int64(6364136223846793005) + int64(1)
	*(*int64)(unsafe.Add(mBase, _c_F_raxRandomWalk[0])) = v46
	goto L8
L6:
	;
	v38 = int32(0)
	goto L5
L7:
	;
	v36 = base.I32_trunc_f64_u(v28)
	v38 = v36
	goto L5
L8:
	;
	v51 = int32(1)
	v53 = base.I32_rem_u_s(base.I32_wrap_i64(int64(base.Ui64(v46)>>(uint(int64(33))%64))), v38<<(uint(v51)%32))
	v56 = v53 + v51
	goto L3
L9:
	;
	if v325 != 0 {
		goto L121
	} else {
		goto L122
	}
L10:
	;
	v79 = int32(0)
	v81 = *(*int64)(unsafe.Add(mBase, _c_F_raxRandomWalk[0]))
	v85 = v81*int64(6364136223846793005) + int64(1)
	*(*int64)(unsafe.Add(mBase, _c_F_raxRandomWalk[0])) = v85
	goto L14
L11:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+152)) = v429
	v443 = *(*int32)(unsafe.Add(mBase, uint32(v429)))
	if v443&int32(2) != 0 {
		v467 = int32(0)
		goto L116
	} else {
		goto L117
	}
L12:
	;
	v435 = *(*int32)(unsafe.Add(mBase, uint32(v429)))
	v437 = v435 & int32(1)
	v438 = v66 - v437
	if v438 != 0 {
		v66 = v438
		v69 = v429
		v72 = v435
		goto L10
	} else {
		goto L114
	}
L13:
	;
	v121 = int32(4)
	v122 = v69 + v121
	v123 = *(*int32)(unsafe.Add(mBase, uint32(v69)))
	if v123&v121 == int32(0) {
		goto L24
	} else {
		goto L25
	}
L14:
	;
	if v72&int32(4) != 0 {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	v95 = int32(1)
	goto L17
L16:
	;
	v95 = int32(base.Ui32(v72) >> (uint(int32(3)) % 32))
	goto L17
L17:
	;
	v96 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v97 = *(*int32)(unsafe.Add(mBase, uint32(v96)))
	v100 = base.I32_rem_s(base.I32_wrap_i64(int64(base.Ui64(v85)>>(uint(int64(33))%64))), v95+base.B2i32(v69 != v97))
	if v100 != v95 {
		goto L13
	} else {
		goto L18
	}
L18:
	;
	v102 = *(*int32)(unsafe.Add(mBase, uint32(l0)+160))
	v104 = v102 + int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+160)) = v104
	v106 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v107 = *(*int32)(unsafe.Add(mBase, uint32(l0)+156))
	v111 = *(*int32)(unsafe.Add(mBase, uint32(v107+v104<<(uint(int32(2))%32))))
	v112 = *(*int32)(unsafe.Add(mBase, uint32(v111)))
	if v112&int32(4) != 0 {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	v118 = int32(base.Ui32(v112) >> (uint(int32(3)) % 32))
	goto L21
L20:
	;
	v118 = int32(1)
	goto L21
L21:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v106 - v118
	v429 = v111
	goto L12
L22:
	;
	v361 = *(*int32)(unsafe.Add(mBase, uint32(l0)+156))
	v362 = *(*int32)(unsafe.Add(mBase, uint32(v69)))
	v363 = *(*int32)(unsafe.Add(mBase, uint32(l0)+160))
	v364 = *(*int32)(unsafe.Add(mBase, uint32(l0)+164))
	if v363 != v364 {
		v406 = v361
		v407 = v363
		goto L98
	} else {
		goto L99
	}
L23:
	;
	v354 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v354 + v350
	goto L22
L24:
	;
	v317 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v318 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v319 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v321 = v319 + int32(1)
	if base.Ui32(v321) <= base.Ui32(v318) {
		v342 = v317
		v343 = v319
		goto L86
	} else {
		goto L87
	}
L25:
	;
	if base.Ui32(v123) < base.Ui32(int32(8)) {
		goto L22
	} else {
		goto L26
	}
L26:
	;
	v130 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v131 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v132 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v134 = int32(base.Ui32(v123) >> (uint(int32(3)) % 32))
	v135 = v132 + v134
	if base.Ui32(v135) <= base.Ui32(v131) {
		v158 = v130
		v159 = v132
		goto L28
	} else {
		goto L29
	}
L27:
	;
	if v139 != 0 {
		goto L82
	} else {
		goto L83
	}
L28:
	;
	v161 = v158 + v159
	if v161 == v122 {
		goto L42
	} else {
		goto L43
	}
L29:
	;
	if v130 == v62 {
		goto L30
	} else {
		goto L31
	}
L30:
	;
	v139 = int32(0)
	goto L32
L31:
	;
	v139 = v130
	goto L32
L32:
	;
	v141 = v135 << (uint(int32(1)) % 32)
	v142 = F_valkey_realloc(m, v139, v141)
	mBase = m.M
	v145 = m.ExcPending
	if v145 != 0 {
		goto L33
	} else {
		goto L34
	}
L33:
	;
	return int32(0)
L34:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v142
	if v142 == int32(0) {
		goto L27
	} else {
		goto L35
	}
L35:
	;
	if v139 != 0 {
		v155 = v142
		goto L36
	} else {
		goto L37
	}
L36:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v141
	v157 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v158 = v155
	v159 = v157
	goto L28
L37:
	;
	v149 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v149 == int32(0) {
		goto L39
	} else {
		goto L40
	}
L38:
	;
	v154 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v155 = v154
	goto L36
L39:
	;
	goto L38
L40:
	;
	v152 = F__emscripten_memcpy_bulkmem(m, v142, v62, v149)
	mBase = m.M
	goto L39
L41:
	;
	v350 = v134
	goto L23
L42:
	;
	goto L41
L43:
	;
	v165 = v134 + v161
	if base.Ui32(int32(0)-v134<<(uint(int32(1))%32)) < base.Ui32(v122-v165) {
		goto L44
	} else {
		goto L45
	}
L44:
	;
	v175 = (v122 ^ v161) & int32(3)
	if base.Ui32(v122) <= base.Ui32(v161) {
		goto L48
	} else {
		goto L49
	}
L45:
	;
	v172 = F___memcpy(m, v161, v122, v134)
	mBase = m.M
	goto L41
L46:
	;
	if v281 == int32(0) {
		goto L42
	} else {
		goto L78
	}
L47:
	;
	if base.Ui32(v259) <= base.Ui32(int32(3)) {
		v280 = v258
		v281 = v259
		v282 = v260
		goto L46
	} else {
		goto L74
	}
L48:
	;
	if v175 != 0 {
		v241 = v134
		goto L58
	} else {
		goto L59
	}
L49:
	;
	if v175 == int32(0) {
		goto L50
	} else {
		goto L51
	}
L50:
	;
	if v161&int32(3) != 0 {
		goto L52
	} else {
		goto L53
	}
L51:
	;
	v280 = v122
	v281 = v134
	v282 = v161
	goto L46
L52:
	;
	v182 = v122
	v183 = v134
	v184 = v161
	goto L54
L53:
	;
	v258 = v122
	v259 = v134
	v260 = v161
	goto L47
L54:
	;
	if v183 == int32(0) {
		goto L42
	} else {
		goto L56
	}
L56:
	;
	v188 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v182))))
	*(*uint8)(unsafe.Add(mBase, uint32(v184))) = uint8(v188)
	v190 = int32(1)
	v191 = v182 + v190
	v193 = v183 + int32(-1)
	v195 = v184 + v190
	if v195&int32(3) == int32(0) {
		v258 = v191
		v259 = v193
		v260 = v195
		goto L47
	} else {
		goto L57
	}
L57:
	;
	v182 = v191
	v183 = v193
	v184 = v195
	goto L54
L58:
	;
	if v241 == int32(0) {
		goto L42
	} else {
		goto L70
	}
L59:
	;
	if v165&int32(3) == int32(0) {
		v221 = v134
		goto L60
	} else {
		goto L61
	}
L60:
	;
	if base.Ui32(v221) <= base.Ui32(int32(3)) {
		v241 = v221
		goto L58
	} else {
		goto L66
	}
L61:
	;
	v206 = v134
	goto L62
L62:
	;
	if v206 == int32(0) {
		goto L42
	} else {
		goto L64
	}
L63:
	;
	v221 = v212
	goto L60
L64:
	;
	v212 = v206 + int32(-1)
	v213 = v161 + v212
	v215 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v122+v212))))
	*(*uint8)(unsafe.Add(mBase, uint32(v213))) = uint8(v215)
	if v213&int32(3) != 0 {
		v206 = v212
		goto L62
	} else {
		goto L65
	}
L65:
	;
	goto L63
L66:
	;
	v228 = v221
	goto L67
L67:
	;
	v232 = v228 + int32(-4)
	v235 = *(*int32)(unsafe.Add(mBase, uint32(v122+v232)))
	*(*int32)(unsafe.Add(mBase, uint32(v161+v232))) = v235
	if base.Ui32(int32(3)) < base.Ui32(v232) {
		v228 = v232
		goto L67
	} else {
		goto L69
	}
L68:
	;
	v241 = v232
	goto L58
L69:
	;
	goto L68
L70:
	;
	v248 = v241
	goto L71
L71:
	;
	v252 = v248 + int32(-1)
	v255 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v122+v252))))
	*(*uint8)(unsafe.Add(mBase, uint32(v161+v252))) = uint8(v255)
	if v252 != 0 {
		v248 = v252
		goto L71
	} else {
		goto L73
	}
L73:
	;
	goto L42
L74:
	;
	v265 = v258
	v266 = v259
	v267 = v260
	goto L75
L75:
	;
	v269 = *(*int32)(unsafe.Add(mBase, uint32(v265)))
	*(*int32)(unsafe.Add(mBase, uint32(v267))) = v269
	v271 = int32(4)
	v272 = v265 + v271
	v274 = v267 + v271
	v276 = v266 + int32(-4)
	if base.Ui32(int32(3)) < base.Ui32(v276) {
		v265 = v272
		v266 = v276
		v267 = v274
		goto L75
	} else {
		goto L77
	}
L76:
	;
	v280 = v272
	v281 = v276
	v282 = v274
	goto L46
L77:
	;
	goto L76
L78:
	;
	v287 = v280
	v288 = v281
	v289 = v282
	goto L79
L79:
	;
	v291 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v287))))
	*(*uint8)(unsafe.Add(mBase, uint32(v289))) = uint8(v291)
	v293 = int32(1)
	v298 = v288 + int32(-1)
	if v298 != 0 {
		v287 = v287 + v293
		v288 = v298
		v289 = v289 + v293
		goto L79
	} else {
		goto L81
	}
L80:
	;
	goto L42
L81:
	;
	goto L80
L82:
	;
	v310 = v139
	goto L84
L83:
	;
	v310 = v62
	goto L84
L84:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v310
	goto L85
L85:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_raxRandomWalk[1])) = int32(48)
	return int32(0)
L86:
	;
	v347 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v122+v100))))
	*(*uint8)(unsafe.Add(mBase, uint32(v342+v343))) = uint8(v347)
	v350 = int32(1)
	goto L23
L87:
	;
	if v317 == v62 {
		goto L88
	} else {
		goto L89
	}
L88:
	;
	v325 = int32(0)
	goto L90
L89:
	;
	v325 = v317
	goto L90
L90:
	;
	v327 = v321 << (uint(int32(1)) % 32)
	v328 = F_valkey_realloc(m, v325, v327)
	mBase = m.M
	v329 = m.ExcPending
	if v329 != 0 {
		goto L33
	} else {
		goto L91
	}
L91:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v328
	if v328 == int32(0) {
		goto L9
	} else {
		goto L92
	}
L92:
	;
	if v325 != 0 {
		v339 = v328
		goto L93
	} else {
		goto L94
	}
L93:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v327
	v341 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v342 = v339
	v343 = v341
	goto L86
L94:
	;
	v333 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v333 == int32(0) {
		goto L96
	} else {
		goto L97
	}
L95:
	;
	v338 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v339 = v338
	goto L93
L96:
	;
	goto L95
L97:
	;
	v336 = F__emscripten_memcpy_bulkmem(m, v328, v62, v333)
	mBase = m.M
	goto L96
L98:
	;
	v408 = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v406+v407<<(uint(v408)%32)))) = v69
	*(*int32)(unsafe.Add(mBase, uint32(l0)+160)) = v407 + int32(1)
	v415 = int32(3)
	v416 = int32(base.Ui32(v362) >> (uint(v415) % 32))
	v428 = *(*int32)(unsafe.Add(mBase, uint32(v69+v416+(int32(0)-v416)&v415+v100<<(uint(v408)%32)+int32(4))))
	v429 = v428
	goto L12
L99:
	;
	v367 = v363 << (uint(int32(3)) % 32)
	if v361 != v60 {
		goto L102
	} else {
		goto L103
	}
L100:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+164)) = v401 << (uint(int32(1)) % 32)
	v405 = *(*int32)(unsafe.Add(mBase, uint32(l0)+160))
	v406 = v400
	v407 = v405
	goto L98
L101:
	;
	v393 = *(*int32)(unsafe.Add(mBase, uint32(l0)+164))
	v395 = v393 << (uint(int32(2)) % 32)
	if v395 == int32(0) {
		goto L112
	} else {
		goto L113
	}
L102:
	;
	v380 = F_valkey_realloc(m, v361, v367)
	mBase = m.M
	v381 = m.ExcPending
	if v381 != 0 {
		goto L33
	} else {
		goto L108
	}
L103:
	;
	v369 = F_valkey_malloc(m, v367)
	mBase = m.M
	v370 = m.ExcPending
	if v370 != 0 {
		goto L33
	} else {
		goto L104
	}
L104:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+156)) = v369
	if v369 != 0 {
		goto L101
	} else {
		goto L105
	}
L105:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+296)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+156)) = v60
	goto L106
L106:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_raxRandomWalk[1])) = int32(48)
	return int32(0)
L107:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+296)) = int32(1)
	goto L110
L108:
	;
	if v380 == int32(0) {
		goto L107
	} else {
		goto L109
	}
L109:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+156)) = v380
	v385 = *(*int32)(unsafe.Add(mBase, uint32(l0)+164))
	v400 = v380
	v401 = v385
	goto L100
L110:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_raxRandomWalk[1])) = int32(48)
	return int32(0)
L111:
	;
	v400 = v369
	v401 = v393
	goto L100
L112:
	;
	goto L111
L113:
	;
	v398 = F__emscripten_memcpy_bulkmem(m, v369, v60, v395)
	mBase = m.M
	goto L112
L114:
	;
	if v437 == int32(0) {
		v66 = v438
		v69 = v429
		v72 = v435
		goto L10
	} else {
		goto L115
	}
L115:
	;
	goto L11
L116:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v467
	return int32(1)
L117:
	;
	v446 = int32(3)
	v447 = int32(base.Ui32(v443) >> (uint(v446) % 32))
	v454 = int32(4)
	if v443&v454 != 0 {
		goto L118
	} else {
		goto L119
	}
L118:
	;
	v459 = v454
	goto L120
L119:
	;
	v459 = v447 << (uint(int32(2)) % 32)
	goto L120
L120:
	;
	v466 = *(*int32)(unsafe.Add(mBase, uint32(v429+v447+(int32(0)-v447)&v446+v459+v443<<(uint(int32(2))%32)&int32(4))))
	v467 = v466
	goto L116
L121:
	;
	v471 = v325
	goto L123
L122:
	;
	v471 = v62
	goto L123
L123:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v471
	goto L124
L124:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_raxRandomWalk[1])) = int32(48)
	return int32(0)
}
func F_raxSeek(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v40 int64
	_ = v40
	var v43 int32
	_ = v43
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v78 int64
	_ = v78
	var v81 int32
	_ = v81
	var v88 int32
	_ = v88
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v96 int32
	_ = v96
	var v97 int64
	_ = v97
	var v105 int32
	_ = v105
	var v112 int32
	_ = v112
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v134 int32
	_ = v134
	var v139 int32
	_ = v139
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v157 int32
	_ = v157
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v164 int32
	_ = v164
	var v168 int32
	_ = v168
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v179 int32
	_ = v179
	var v184 int32
	_ = v184
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v196 int32
	_ = v196
	var v200 int32
	_ = v200
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v214 int32
	_ = v214
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
	var v239 int32
	_ = v239
	var v241 int32
	_ = v241
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v250 int32
	_ = v250
	var v255 int32
	_ = v255
	var v257 int32
	_ = v257
	var v258 int32
	_ = v258
	var v261 int32
	_ = v261
	var v264 int32
	_ = v264
	var v265 int32
	_ = v265
	var v268 int32
	_ = v268
	var v269 int32
	_ = v269
	var v271 int32
	_ = v271
	var v275 int32
	_ = v275
	var v276 int32
	_ = v276
	var v290 int32
	_ = v290
	var v293 int32
	_ = v293
	var v294 int32
	_ = v294
	var v301 int32
	_ = v301
	var v306 int32
	_ = v306
	var v313 int32
	_ = v313
	var v314 int32
	_ = v314
	var v318 int32
	_ = v318
	var v320 int32
	_ = v320
	var v321 int32
	_ = v321
	var v327 int32
	_ = v327
	var v329 int32
	_ = v329
	var v330 int32
	_ = v330
	var v334 int32
	_ = v334
	var v335 int32
	_ = v335
	var v339 int32
	_ = v339
	var v347 int32
	_ = v347
	var v352 int32
	_ = v352
	var v358 int32
	_ = v358
	var v359 int32
	_ = v359
	var v362 int32
	_ = v362
	var v366 int32
	_ = v366
	var v369 int32
	_ = v369
	var v370 int32
	_ = v370
	var v377 int32
	_ = v377
	var v382 int32
	_ = v382
	var v389 int32
	_ = v389
	var v390 int32
	_ = v390
	var v392 int32
	_ = v392
	var v395 int32
	_ = v395
	var v398 int32
	_ = v398
	var v399 int32
	_ = v399
	var v403 int32
	_ = v403
	var v404 int32
	_ = v404
	var v408 int32
	_ = v408
	var v409 int32
	_ = v409
	var v413 int32
	_ = v413
	v5 = int32(0)
	v14 = m.G0
	v16 = v14 - int32(16)
	m.G0 = v16
	*(*int32)(unsafe.Add(mBase, uint32(l0)+160)) = v5
	*(*int32)(unsafe.Add(mBase, uint32(l0)+152)) = v5
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v5
	v25 = int32(1)
	v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v30 = v26&int32(-4) | v25
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v30
	v33 = l0 + int32(152)
	v36 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	switch v36 + int32(-60) {
	case 0:
		goto L10
	case 1:
		v70 = l2
		v71 = l3
		v72 = v5
		v74 = v26
		v75 = v5
		v76 = v5
		goto L8
	case 2:
		v60 = l1
		v61 = l2
		v62 = l3
		v63 = v5
		v64 = v25
		v65 = v26
		goto L9
	case 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33:
		goto L11
	case 34:
		goto L13
	default:
		goto L12
	}
L1:
	;
	m.G0 = v16 + int32(16)
	return v413
L2:
	;
	if v75 == int32(0) {
		goto L90
	} else {
		goto L91
	}
L3:
	;
	F__serverAssert(m, int32(_a_F_raxSeek_0), int32(_a_F_raxSeek_1), int32(1542))
	mBase = m.M
	v347 = m.ExcPending
	if v347 != 0 {
		goto L18
	} else {
		goto L89
	}
L4:
	;
	if v76|base.B2i32(v88 != v71) != 0 {
		goto L31
	} else {
		goto L32
	}
L5:
	;
	v112 = *(*int32)(unsafe.Add(mBase, uint32(v96)))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+152)) = v112
	v114 = F_raxSeekGreatest(m, l0)
	mBase = m.M
	v115 = m.ExcPending
	if v115 != 0 {
		goto L18
	} else {
		goto L23
	}
L6:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v105 | int32(3)
	v413 = int32(1)
	goto L1
L7:
	;
	v96 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v97 = *(*int64)(unsafe.Add(mBase, uint32(v96)+8))
	if v97 != int64(0) {
		goto L5
	} else {
		goto L21
	}
L8:
	;
	v77 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v78 = *(*int64)(unsafe.Add(mBase, uint32(v77)+8))
	if v78 == int64(0) {
		v105 = v74
		goto L6
	} else {
		goto L17
	}
L9:
	;
	v66 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v60)+1)))
	v70 = v61
	v71 = v62
	v72 = v64
	v74 = v65
	v75 = v63
	v76 = base.B2i32(v66 != int32(61))
	goto L8
L10:
	;
	v60 = l1
	v61 = l2
	v62 = l3
	v63 = int32(1)
	v64 = int32(0)
	v65 = v26
	goto L9
L11:
	;
	goto L16
L12:
	;
	if v36 == int32(36) {
		goto L7
	} else {
		goto L15
	}
L13:
	;
	v39 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v40 = *(*int64)(unsafe.Add(mBase, uint32(v39)+8))
	if v40 == int64(0) {
		v105 = v26
		goto L6
	} else {
		goto L14
	}
L14:
	;
	v43 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+152)) = v43
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v43
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v30
	v60 = int32(_a_F_raxSeek_2)
	v61 = v43
	v62 = v43
	v63 = v43
	v64 = v25
	v65 = v30
	goto L9
L15:
	;
	goto L11
L16:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_raxSeek[0])) = int32(0)
	v413 = int32(0)
	goto L1
L17:
	;
	v81 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v16)+12)) = v81
	v88 = F_raxLowWalk(m, v77, v70, v71, v33, v81, v16+int32(12), l0+int32(156))
	mBase = m.M
	v91 = m.ExcPending
	if v91 != 0 {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	return int32(0)
L19:
	;
	v92 = *(*int32)(unsafe.Add(mBase, uint32(l0)+296))
	if v92 == int32(0) {
		goto L4
	} else {
		goto L20
	}
L20:
	;
	v413 = int32(0)
	goto L1
L21:
	;
	v105 = v26
	goto L6
L22:
	;
	v117 = *(*int32)(unsafe.Add(mBase, uint32(v33)))
	v118 = *(*int32)(unsafe.Add(mBase, uint32(v117)))
	if v118&int32(1) == int32(0) {
		goto L3
	} else {
		goto L25
	}
L23:
	;
	if v114 != 0 {
		goto L22
	} else {
		goto L24
	}
L24:
	;
	v413 = int32(0)
	goto L1
L25:
	;
	if v118&int32(2) != 0 {
		v144 = int32(0)
		goto L26
	} else {
		goto L27
	}
L26:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v144
	v413 = int32(1)
	goto L1
L27:
	;
	v126 = int32(3)
	v127 = int32(base.Ui32(v118) >> (uint(v126) % 32))
	v134 = int32(4)
	if v118&v134 != 0 {
		goto L28
	} else {
		goto L29
	}
L28:
	;
	v139 = v134
	goto L30
L29:
	;
	v139 = v127 << (uint(int32(2)) % 32)
	goto L30
L30:
	;
	v143 = *(*int32)(unsafe.Add(mBase, uint32(v117+v127+(int32(0)-v127)&v126+v139+int32(4))))
	v144 = v143
	goto L26
L31:
	;
	v196 = int32(1)
	if v72|v75 != v196 {
		goto L44
	} else {
		goto L45
	}
L32:
	;
	v149 = *(*int32)(unsafe.Add(mBase, uint32(v33)))
	v150 = *(*int32)(unsafe.Add(mBase, uint32(v149)))
	if v150&int32(1) == int32(0) {
		goto L31
	} else {
		goto L33
	}
L33:
	;
	v157 = *(*int32)(unsafe.Add(mBase, uint32(v16)+12))
	if int32(base.Ui32(v150)>>(uint(int32(2))%32))&base.B2i32(v157 != int32(0)) != 0 {
		goto L31
	} else {
		goto L34
	}
L34:
	;
	v161 = F_raxIteratorAddChars(m, l0, v70, v71)
	mBase = m.M
	v162 = m.ExcPending
	if v162 != 0 {
		goto L18
	} else {
		goto L36
	}
L35:
	;
	v164 = *(*int32)(unsafe.Add(mBase, uint32(l0)+152))
	v168 = *(*int32)(unsafe.Add(mBase, uint32(v164)))
	if v168&int32(2) != 0 {
		v192 = int32(0)
		goto L39
	} else {
		goto L40
	}
L36:
	;
	if v161 != 0 {
		goto L35
	} else {
		goto L37
	}
L37:
	;
	v413 = int32(0)
	goto L1
L38:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v192
	v413 = int32(1)
	goto L1
L39:
	;
	goto L38
L40:
	;
	v171 = int32(3)
	v172 = int32(base.Ui32(v168) >> (uint(v171) % 32))
	v179 = int32(4)
	if v168&v179 != 0 {
		goto L41
	} else {
		goto L42
	}
L41:
	;
	v184 = v179
	goto L43
L42:
	;
	v184 = v172 << (uint(int32(2)) % 32)
	goto L43
L43:
	;
	v191 = *(*int32)(unsafe.Add(mBase, uint32(v164+v172+(int32(0)-v172)&v171+v184+v168<<(uint(int32(2))%32)&int32(4))))
	v192 = v191
	goto L39
L44:
	;
	v339 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v339 | int32(2)
	v413 = v196
	goto L1
L45:
	;
	v200 = *(*int32)(unsafe.Add(mBase, uint32(v16)+12))
	v202 = F_raxIteratorAddChars(m, l0, v70, v88-v200)
	mBase = m.M
	v203 = m.ExcPending
	if v203 != 0 {
		goto L18
	} else {
		goto L46
	}
L46:
	;
	if v88 == v71 {
		goto L47
	} else {
		goto L48
	}
L47:
	;
	v271 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v271 & int32(-2)
	v275 = *(*int32)(unsafe.Add(mBase, uint32(l0)+152))
	v276 = *(*int32)(unsafe.Add(mBase, uint32(v275)))
	if v276&int32(4) == int32(0) {
		goto L73
	} else {
		goto L74
	}
L48:
	;
	v205 = *(*int32)(unsafe.Add(mBase, uint32(v33)))
	v206 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v205))))
	if v206&int32(4) != 0 {
		goto L49
	} else {
		goto L50
	}
L49:
	;
	v239 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v70+v88))))
	v241 = v205 + int32(4)
	v243 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v241+v200))))
	v244 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v244 & int32(-2)
	if v72 == int32(0) {
		goto L2
	} else {
		goto L62
	}
L50:
	;
	v211 = F_raxIteratorAddChars(m, l0, v70+v88, int32(1))
	mBase = m.M
	v212 = m.ExcPending
	if v212 != 0 {
		goto L18
	} else {
		goto L52
	}
L51:
	;
	v214 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v214 & int32(-2)
	if v75 == int32(0) {
		goto L54
	} else {
		goto L55
	}
L52:
	;
	if v211 != 0 {
		goto L51
	} else {
		goto L53
	}
L53:
	;
	v413 = int32(0)
	goto L1
L54:
	;
	if v72 == int32(0) {
		goto L58
	} else {
		goto L59
	}
L55:
	;
	v221 = F_raxIteratorPrevStep(m, l0, int32(1))
	mBase = m.M
	v222 = m.ExcPending
	if v222 != 0 {
		goto L18
	} else {
		goto L56
	}
L56:
	;
	if v221 != 0 {
		goto L54
	} else {
		goto L57
	}
L57:
	;
	v413 = int32(0)
	goto L1
L58:
	;
	v233 = int32(1)
	v234 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v234 | v233
	v413 = v233
	goto L1
L59:
	;
	v228 = F_raxIteratorNextStep(m, l0, int32(1))
	mBase = m.M
	v229 = m.ExcPending
	if v229 != 0 {
		goto L18
	} else {
		goto L60
	}
L60:
	;
	if v228 == int32(0) {
		v413 = int32(0)
		goto L1
	} else {
		goto L61
	}
L61:
	;
	goto L58
L62:
	;
	v250 = int32(255)
	if base.Ui32(v243&v250) <= base.Ui32(v239&v250) {
		goto L63
	} else {
		goto L64
	}
L63:
	;
	v261 = *(*int32)(unsafe.Add(mBase, uint32(v205)))
	v264 = F_raxIteratorAddChars(m, l0, v241, int32(base.Ui32(v261)>>(uint(int32(3))%32)))
	mBase = m.M
	v265 = m.ExcPending
	if v265 != 0 {
		goto L18
	} else {
		goto L68
	}
L64:
	;
	v255 = int32(0)
	v257 = F_raxIteratorNextStep(m, l0, v255)
	mBase = m.M
	v258 = m.ExcPending
	if v258 != 0 {
		goto L18
	} else {
		goto L65
	}
L65:
	;
	if v257 == int32(0) {
		v413 = v255
		goto L1
	} else {
		goto L66
	}
L66:
	;
	goto L2
L67:
	;
	v268 = F_raxIteratorNextStep(m, l0, int32(1))
	mBase = m.M
	v269 = m.ExcPending
	if v269 != 0 {
		goto L18
	} else {
		goto L70
	}
L68:
	;
	if v264 != 0 {
		goto L67
	} else {
		goto L69
	}
L69:
	;
	v413 = int32(0)
	goto L1
L70:
	;
	if v268 != 0 {
		goto L2
	} else {
		goto L71
	}
L71:
	;
	v413 = int32(0)
	goto L1
L72:
	;
	v334 = int32(1)
	v335 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v335 | v334
	v413 = v334
	goto L1
L73:
	;
	if v72 == int32(0) {
		goto L82
	} else {
		goto L83
	}
L74:
	;
	if v75&(base.B2i32(v200 != int32(0))&v276) != int32(1) {
		goto L73
	} else {
		goto L75
	}
L75:
	;
	v290 = *(*int32)(unsafe.Add(mBase, uint32(v275)))
	if v290&int32(2) != 0 {
		v314 = int32(0)
		goto L77
	} else {
		goto L78
	}
L76:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v314
	goto L72
L77:
	;
	goto L76
L78:
	;
	v293 = int32(3)
	v294 = int32(base.Ui32(v290) >> (uint(v293) % 32))
	v301 = int32(4)
	if v290&v301 != 0 {
		goto L79
	} else {
		goto L80
	}
L79:
	;
	v306 = v301
	goto L81
L80:
	;
	v306 = v294 << (uint(int32(2)) % 32)
	goto L81
L81:
	;
	v313 = *(*int32)(unsafe.Add(mBase, uint32(v275+v294+(int32(0)-v294)&v293+v306+v290<<(uint(int32(2))%32)&int32(4))))
	v314 = v313
	goto L77
L82:
	;
	if v75 == int32(0) {
		goto L72
	} else {
		goto L86
	}
L83:
	;
	v318 = int32(0)
	v320 = F_raxIteratorNextStep(m, l0, v318)
	mBase = m.M
	v321 = m.ExcPending
	if v321 != 0 {
		goto L18
	} else {
		goto L84
	}
L84:
	;
	if v320 == int32(0) {
		v413 = v318
		goto L1
	} else {
		goto L85
	}
L85:
	;
	goto L82
L86:
	;
	v327 = int32(0)
	v329 = F_raxIteratorPrevStep(m, l0, v327)
	mBase = m.M
	v330 = m.ExcPending
	if v330 != 0 {
		goto L18
	} else {
		goto L87
	}
L87:
	;
	if v329 == int32(0) {
		v413 = v327
		goto L1
	} else {
		goto L88
	}
L88:
	;
	goto L72
L89:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L90:
	;
	v408 = int32(1)
	v409 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v409 | v408
	v413 = v408
	goto L1
L91:
	;
	v352 = int32(255)
	if base.Ui32(v239&v352) <= base.Ui32(v243&v352) {
		goto L92
	} else {
		goto L93
	}
L92:
	;
	v392 = *(*int32)(unsafe.Add(mBase, uint32(l0)+152))
	v395 = *(*int32)(unsafe.Add(mBase, uint32(v392)))
	v398 = F_raxIteratorAddChars(m, l0, v392+int32(4), int32(base.Ui32(v395)>>(uint(int32(3))%32)))
	mBase = m.M
	v399 = m.ExcPending
	if v399 != 0 {
		goto L18
	} else {
		goto L103
	}
L93:
	;
	v358 = F_raxSeekGreatest(m, l0)
	mBase = m.M
	v359 = m.ExcPending
	if v359 != 0 {
		goto L18
	} else {
		goto L94
	}
L94:
	;
	if v358 == int32(0) {
		v413 = int32(0)
		goto L1
	} else {
		goto L95
	}
L95:
	;
	v362 = *(*int32)(unsafe.Add(mBase, uint32(l0)+152))
	v366 = *(*int32)(unsafe.Add(mBase, uint32(v362)))
	if v366&int32(2) != 0 {
		v390 = int32(0)
		goto L97
	} else {
		goto L98
	}
L96:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v390
	goto L90
L97:
	;
	goto L96
L98:
	;
	v369 = int32(3)
	v370 = int32(base.Ui32(v366) >> (uint(v369) % 32))
	v377 = int32(4)
	if v366&v377 != 0 {
		goto L99
	} else {
		goto L100
	}
L99:
	;
	v382 = v377
	goto L101
L100:
	;
	v382 = v370 << (uint(int32(2)) % 32)
	goto L101
L101:
	;
	v389 = *(*int32)(unsafe.Add(mBase, uint32(v362+v370+(int32(0)-v370)&v369+v382+v366<<(uint(int32(2))%32)&int32(4))))
	v390 = v389
	goto L97
L102:
	;
	v403 = F_raxIteratorPrevStep(m, l0, int32(1))
	mBase = m.M
	v404 = m.ExcPending
	if v404 != 0 {
		goto L18
	} else {
		goto L105
	}
L103:
	;
	if v398 != 0 {
		goto L102
	} else {
		goto L104
	}
L104:
	;
	v413 = int32(0)
	goto L1
L105:
	;
	if v403 == int32(0) {
		v413 = int32(0)
		goto L1
	} else {
		goto L106
	}
L106:
	;
	goto L90
}
func F_raxStart(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v8 int64
	_ = v8
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = int32(128)
	v8 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(l0)+12)) = v8
	*(*int64)(unsafe.Add(mBase, uint32(l0)+296)) = v8
	*(*int64)(unsafe.Add(mBase, uint32(l0)+160)) = int64(137438953472)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = l0 + int32(24)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+156)) = l0 + int32(168)
	return
}
func F_raxTryInsert(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	v7 = F_raxGenericInsert(m, l0, l1, l2, l3, l4, int32(0))
	v10 = m.ExcPending
	if v10 != 0 {
		return int32(0)
	} else {
		return v7
	}
}
func F_shrinkRaxBucketIfPossible(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v19 int32
	_ = v19
	var v26 int32
	_ = v26
	var v27 int64
	_ = v27
	var v35 int64
	_ = v35
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v62 int32
	_ = v62
	var v67 int32
	_ = v67
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v82 int32
	_ = v82
	var v85 int64
	_ = v85
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v95 int32
	_ = v95
	var v101 int32
	_ = v101
	var v102 int64
	_ = v102
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v125 int32
	_ = v125
	var v137 int32
	_ = v137
	var v143 int32
	_ = v143
	var v149 int32
	_ = v149
	var v155 int32
	_ = v155
	v8 = m.G0
	v10 = v8 - int32(304)
	m.G0 = v10
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	switch v12 + int32(1) {
	case 0:
		F__serverAssert(m, int32(_a_F_shrinkRaxBucketIfPossible_0), int32(_a_F_shrinkRaxBucketIfPossible_1), int32(807))
		mBase = m.M
		v137 = m.ExcPending
		if v137 != 0 {
			return
		} else {
			F_abort(m)
			mBase = m.M
			base.Wasm_trap_unreachable()
			for {
			}
		}
	case 1:
		F__serverAssert(m, int32(_a_F_shrinkRaxBucketIfPossible_2), int32(_a_F_shrinkRaxBucketIfPossible_1), int32(781))
		mBase = m.M
		v19 = m.ExcPending
		if v19 != 0 {
			return
		} else {
			F_abort(m)
			mBase = m.M
			base.Wasm_trap_unreachable()
			for {
			}
		}
	default:
		if v12&int32(7) != int32(6) {
			F__serverAssert(m, int32(_a_F_shrinkRaxBucketIfPossible_0), int32(_a_F_shrinkRaxBucketIfPossible_1), int32(807))
			mBase = m.M
			v137 = m.ExcPending
			if v137 != 0 {
				return
			} else {
				F_abort(m)
				mBase = m.M
				base.Wasm_trap_unreachable()
				for {
				}
			}
		} else {
			v26 = v12 & int32(-8)
			v27 = *(*int64)(unsafe.Add(mBase, uint32(v26)+8))
			if v27 != int64(1) {
				m.G0 = v10 + int32(304)
				return
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v10)+4)) = v26
				*(*int32)(unsafe.Add(mBase, uint32(v10))) = int32(2)
				*(*int32)(unsafe.Add(mBase, uint32(v10)+20)) = int32(128)
				v35 = int64(0)
				*(*int64)(unsafe.Add(mBase, uint32(v10)+12)) = v35
				*(*int64)(unsafe.Add(mBase, uint32(v10)+296)) = v35
				*(*int64)(unsafe.Add(mBase, uint32(v10)+160)) = int64(137438953472)
				*(*int32)(unsafe.Add(mBase, uint32(v10)+8)) = v10 + int32(24)
				*(*int32)(unsafe.Add(mBase, uint32(v10)+156)) = v10 + int32(168)
				v48 = int32(0)
				v50 = F_raxSeek(m, v10, int32(_a_F_shrinkRaxBucketIfPossible_3), v48, v48)
				mBase = m.M
				v51 = m.ExcPending
				if v51 != 0 {
					return
				} else {
					if v50 == int32(0) {
						F__serverAssert(m, int32(_a_F_shrinkRaxBucketIfPossible_4), int32(_a_F_shrinkRaxBucketIfPossible_1), int32(1433))
						mBase = m.M
						v143 = m.ExcPending
						if v143 != 0 {
							return
						} else {
							F_abort(m)
							mBase = m.M
							base.Wasm_trap_unreachable()
							for {
							}
						}
					} else {
						v54 = F_raxNext(m, v10)
						mBase = m.M
						v55 = m.ExcPending
						if v55 != 0 {
							return
						} else {
							if v54 == int32(0) {
								F__serverAssert(m, int32(_a_F_shrinkRaxBucketIfPossible_5), int32(_a_F_shrinkRaxBucketIfPossible_1), int32(1434))
								mBase = m.M
								v149 = m.ExcPending
								if v149 != 0 {
									return
								} else {
									F_abort(m)
									mBase = m.M
									base.Wasm_trap_unreachable()
									for {
									}
								}
							} else {
								v58 = *(*int32)(unsafe.Add(mBase, uint32(v10)+12))
								switch v58 + int32(1) {
								case 0:
									F_raxStop(m, v10)
									mBase = m.M
									v62 = m.ExcPending
									if v62 != 0 {
										return
									} else {
										m.G0 = v10 + int32(304)
										return
									}
								case 1:
									F__serverAssert(m, int32(_a_F_shrinkRaxBucketIfPossible_2), int32(_a_F_shrinkRaxBucketIfPossible_1), int32(781))
									mBase = m.M
									v67 = m.ExcPending
									if v67 != 0 {
										return
									} else {
										F_abort(m)
										mBase = m.M
										base.Wasm_trap_unreachable()
										for {
										}
									}
								default:
									if v58&int32(1) == int32(0) {
										F_raxStop(m, v10)
										mBase = m.M
										v76 = m.ExcPending
										if v76 != 0 {
											return
										} else {
											if v58&int32(6) != int32(2) {
												m.G0 = v10 + int32(304)
												return
											} else {
												v82 = v58 & int32(-8)
												if v82 == int32(0) {
													v95 = *(*int32)(unsafe.Add(mBase, _c_F_shrinkRaxBucketIfPossible[0]))
													if v95 != 0 {
														F__serverAssert(m, int32(_a_F_shrinkRaxBucketIfPossible_6), int32(_a_F_shrinkRaxBucketIfPossible_1), int32(853))
														mBase = m.M
														v155 = m.ExcPending
														if v155 != 0 {
															return
														} else {
															F_abort(m)
															mBase = m.M
															base.Wasm_trap_unreachable()
															for {
															}
														}
													} else {
														*(*int32)(unsafe.Add(mBase, _c_F_shrinkRaxBucketIfPossible[0])) = l1
														v116 = int32(_a_F_shrinkRaxBucketIfPossible_7)
														*(*int32)(unsafe.Add(mBase, uint32(v116))) = int32(0)
														*(*int32)(unsafe.Add(mBase, uint32(l0))) = v58
														F_raxFree(m, v26)
														mBase = m.M
														v125 = m.ExcPending
														if v125 != 0 {
															return
														} else {
															m.G0 = v10 + int32(304)
															return
														}
													}
												} else {
													v85 = *(*int64)(unsafe.Add(mBase, uint32(v82)))
													if base.Ui64(int64(127)) <= base.Ui64(v85&int64(1073741823)) {
														m.G0 = v10 + int32(304)
														return
													} else {
														v90 = int32(0)
														v91 = *(*int32)(unsafe.Add(mBase, _c_F_shrinkRaxBucketIfPossible[0]))
														if v91 == v90 {
															*(*int32)(unsafe.Add(mBase, _c_F_shrinkRaxBucketIfPossible[0])) = l1
															v101 = int32(_a_F_shrinkRaxBucketIfPossible_7)
															v102 = *(*int64)(unsafe.Add(mBase, uint32(v82)))
															if v102&int64(1073741822) == int64(0) {
																v116 = v101
																*(*int32)(unsafe.Add(mBase, uint32(v116))) = int32(0)
																*(*int32)(unsafe.Add(mBase, uint32(l0))) = v58
																F_raxFree(m, v26)
																mBase = m.M
																v125 = m.ExcPending
																if v125 != 0 {
																	return
																} else {
																	m.G0 = v10 + int32(304)
																	return
																}
															} else {
																F_qsort(m, v82+int32(8), base.I32_wrap_i64(v102)&int32(1073741823), int32(4), int32(1128))
																mBase = m.M
																v115 = m.ExcPending
																if v115 != 0 {
																	return
																} else {
																	v116 = v101
																	*(*int32)(unsafe.Add(mBase, uint32(v116))) = int32(0)
																	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v58
																	F_raxFree(m, v26)
																	mBase = m.M
																	v125 = m.ExcPending
																	if v125 != 0 {
																		return
																	} else {
																		m.G0 = v10 + int32(304)
																		return
																	}
																}
															}
														} else {
															F__serverAssert(m, int32(_a_F_shrinkRaxBucketIfPossible_6), int32(_a_F_shrinkRaxBucketIfPossible_1), int32(853))
															mBase = m.M
															v155 = m.ExcPending
															if v155 != 0 {
																return
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
									} else {
										F_raxStop(m, v10)
										mBase = m.M
										v74 = m.ExcPending
										if v74 != 0 {
											return
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(l0))) = v58
											F_raxFree(m, v26)
											mBase = m.M
											v125 = m.ExcPending
											if v125 != 0 {
												return
											} else {
												m.G0 = v10 + int32(304)
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
