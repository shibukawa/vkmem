package p1

import (
	base "github.com/shibukawa/vkmem/internal/aot/vkaot/base"
	"unsafe"
)

func F_memtest(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v12 int32
	_ = v12
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
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
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	*(*int32)(unsafe.Add(mBase, uint32(v6))) = int32(_a749)
	v12 = F_ioctl(m, int32(1), int32(21523), v6)
	mBase = m.M
	if v12 != int32(-1) {
	} else {
		*(*int32)(unsafe.Add(mBase, _consts[381])) = int32(5242900)
	}
	F_memtest_alloc_and_test(m, l0, l1)
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		return
	} else {
		v21 = F_puts(m, int32(_a750))
		mBase = m.M
		v22 = m.ExcPending
		if v22 != 0 {
			return
		} else {
			v24 = F_puts(m, int32(_a751))
			mBase = m.M
			v25 = m.ExcPending
			if v25 != 0 {
				return
			} else {
				v27 = F_puts(m, int32(_a752))
				mBase = m.M
				v28 = m.ExcPending
				if v28 != 0 {
					return
				} else {
					v30 = F_puts(m, int32(_a753))
					mBase = m.M
					v31 = m.ExcPending
					if v31 != 0 {
						return
					} else {
						m.Env.Exit(m, int32(0))
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
func F_memtest_test(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v26 int64
	_ = v26
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v36 int32
	_ = v36
	var v40 int32
	_ = v40
	var v48 int32
	_ = v48
	var v55 int32
	_ = v55
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
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v102 int32
	_ = v102
	var v119 int32
	_ = v119
	var v124 int32
	_ = v124
	var v128 int64
	_ = v128
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v134 int32
	_ = v134
	var v139 int64
	_ = v139
	var v142 int32
	_ = v142
	var v145 int32
	_ = v145
	var v147 int32
	_ = v147
	var v151 int32
	_ = v151
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v165 int32
	_ = v165
	var v170 int32
	_ = v170
	var v175 int32
	_ = v175
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v184 int32
	_ = v184
	var v198 int32
	_ = v198
	var v202 int32
	_ = v202
	var v207 int32
	_ = v207
	var v220 int32
	_ = v220
	var v225 int32
	_ = v225
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v248 int32
	_ = v248
	var v265 int32
	_ = v265
	var v270 int32
	_ = v270
	var v274 int64
	_ = v274
	var v276 int32
	_ = v276
	var v277 int32
	_ = v277
	var v280 int32
	_ = v280
	var v285 int64
	_ = v285
	var v288 int32
	_ = v288
	var v291 int32
	_ = v291
	var v293 int32
	_ = v293
	var v301 int32
	_ = v301
	var v308 int32
	_ = v308
	var v313 int32
	_ = v313
	var v317 int64
	_ = v317
	var v319 int32
	_ = v319
	var v320 int32
	_ = v320
	var v323 int32
	_ = v323
	var v328 int64
	_ = v328
	var v333 int32
	_ = v333
	var v341 int32
	_ = v341
	var v346 int32
	_ = v346
	var v350 int64
	_ = v350
	var v352 int32
	_ = v352
	var v353 int32
	_ = v353
	var v356 int32
	_ = v356
	var v361 int64
	_ = v361
	var v364 int32
	_ = v364
	var v374 int32
	_ = v374
	var v379 int32
	_ = v379
	var v383 int64
	_ = v383
	var v385 int32
	_ = v385
	var v386 int32
	_ = v386
	var v389 int32
	_ = v389
	var v394 int64
	_ = v394
	var v397 int32
	_ = v397
	var v400 int32
	_ = v400
	var v415 int32
	_ = v415
	var v416 int64
	_ = v416
	var v417 int32
	_ = v417
	var v418 int32
	_ = v418
	var v419 int32
	_ = v419
	var v422 int32
	_ = v422
	var v428 int64
	_ = v428
	var v434 int32
	_ = v434
	var v438 int32
	_ = v438
	var v439 int32
	_ = v439
	var v440 int32
	_ = v440
	var v443 int32
	_ = v443
	var v444 int32
	_ = v444
	var v445 int32
	_ = v445
	var v447 int32
	_ = v447
	var v452 int32
	_ = v452
	var v457 int32
	_ = v457
	var v460 int32
	_ = v460
	var v461 int32
	_ = v461
	var v464 int32
	_ = v464
	var v465 int32
	_ = v465
	var v466 int32
	_ = v466
	var v469 int32
	_ = v469
	var v470 int32
	_ = v470
	var v473 int32
	_ = v473
	var v474 int32
	_ = v474
	var v475 int32
	_ = v475
	var v478 int32
	_ = v478
	var v479 int32
	_ = v479
	var v482 int32
	_ = v482
	var v483 int32
	_ = v483
	var v484 int32
	_ = v484
	var v487 int32
	_ = v487
	var v488 int32
	_ = v488
	var v491 int32
	_ = v491
	var v492 int32
	_ = v492
	var v493 int32
	_ = v493
	var v496 int32
	_ = v496
	var v497 int32
	_ = v497
	var v501 int32
	_ = v501
	var v504 int32
	_ = v504
	var v510 int32
	_ = v510
	var v517 int32
	_ = v517
	var v522 int32
	_ = v522
	var v524 int32
	_ = v524
	var v530 int32
	_ = v530
	if l2 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v20 = l1 & int32(4095)
	v22 = int32(base.Ui32(l1) >> (uint(int32(3)) % 32))
	v25 = l0 + v22<<(uint(int32(2))%32)
	v26 = base.I64_extend_i32_u(v22)
	v28 = base.B2i32(base.Ui32(l1) < base.Ui32(int32(8)))
	v29 = int32(0)
	v36 = v29
	v40 = v29
	goto L3
L2:
	;
	return int32(0)
L3:
	;
	v48 = v40 + int32(1)
	if l3 == int32(0) {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	return v530
L5:
	;
	v79 = int32(0)
	v93 = v79
	v94 = v79
	goto L20
L6:
	;
	v72 = F_memtest_addressing(m, l0, l1, int32(0))
	mBase = m.M
	v73 = m.ExcPending
	if v73 != 0 {
		goto L8
	} else {
		goto L15
	}
L7:
	;
	F_memtest_progress_start(m, int32(_a741), v48)
	mBase = m.M
	v55 = m.ExcPending
	if v55 != 0 {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	return int32(0)
L9:
	;
	v56 = F_memtest_addressing(m, l0, l1, l3)
	mBase = m.M
	v57 = m.ExcPending
	if v57 != 0 {
		goto L8
	} else {
		goto L10
	}
L10:
	;
	v60 = F_iprintf(m, int32(_a742), int32(0))
	mBase = m.M
	v61 = m.ExcPending
	if v61 != 0 {
		goto L8
	} else {
		goto L11
	}
L11:
	;
	F_memtest_progress_start(m, int32(_a743), v48)
	mBase = m.M
	v64 = m.ExcPending
	if v64 != 0 {
		goto L8
	} else {
		goto L12
	}
L12:
	;
	F_memtest_fill_random(m, l0, l1, l3)
	mBase = m.M
	v66 = m.ExcPending
	if v66 != 0 {
		goto L8
	} else {
		goto L13
	}
L13:
	;
	v69 = F_iprintf(m, int32(_a742), int32(0))
	mBase = m.M
	v70 = m.ExcPending
	if v70 != 0 {
		goto L8
	} else {
		goto L14
	}
L14:
	;
	v77 = v56
	goto L5
L15:
	;
	F_memtest_fill_random(m, l0, l1, int32(0))
	mBase = m.M
	v76 = m.ExcPending
	if v76 != 0 {
		goto L8
	} else {
		goto L16
	}
L16:
	;
	v77 = v72
	goto L5
L17:
	;
	v225 = int32(0)
	v239 = v225
	v240 = v225
	goto L47
L18:
	;
	v202 = int32(0)
	F_memtest_fill_value(m, l0, l1, v202, int32(-1), int32(83), v202)
	mBase = m.M
	v207 = m.ExcPending
	if v207 != 0 {
		goto L8
	} else {
		goto L45
	}
L19:
	;
	F__serverAssert(m, int32(_a744), int32(_a745), int32(199))
	mBase = m.M
	v184 = m.ExcPending
	if v184 != 0 {
		goto L8
	} else {
		goto L44
	}
L20:
	;
	v102 = v94
	goto L23
L21:
	;
	F_memtest_progress_start(m, int32(_a746), v48)
	mBase = m.M
	v170 = m.ExcPending
	if v170 != 0 {
		goto L8
	} else {
		goto L41
	}
L22:
	;
	F_memtest_progress_start(m, int32(_a747), v48)
	mBase = m.M
	v156 = m.ExcPending
	if v156 != 0 {
		goto L8
	} else {
		goto L37
	}
L23:
	;
	if l3 != 0 {
		goto L22
	} else {
		goto L25
	}
L25:
	;
	if v20 != 0 {
		goto L19
	} else {
		goto L26
	}
L26:
	;
	if base.Ui32(l1) < base.Ui32(int32(8)) {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	v151 = v102 + int32(1)
	if v151 != int32(4) {
		v102 = v151
		goto L23
	} else {
		goto L36
	}
L28:
	;
	v119 = v25
	v124 = l0
	v128 = int64(0)
	goto L30
L29:
	;
	v145 = v142 + v93
	v147 = v94 + int32(1)
	if v147 != int32(4) {
		v93 = v145
		v94 = v147
		goto L20
	} else {
		goto L35
	}
L30:
	;
	v130 = *(*int32)(unsafe.Add(mBase, uint32(v124)))
	v131 = *(*int32)(unsafe.Add(mBase, uint32(v119)))
	if v130 == v131 {
		goto L32
	} else {
		goto L33
	}
L31:
	;
	v142 = int32(0)
	goto L29
L32:
	;
	v134 = int32(4)
	v139 = v128 + int64(1)
	if v139 != v26 {
		v119 = v119 + v134
		v124 = v124 + v134
		v128 = v139
		goto L30
	} else {
		goto L34
	}
L33:
	;
	v142 = int32(1)
	goto L29
L34:
	;
	goto L31
L35:
	;
	v198 = v145
	goto L18
L36:
	;
	v198 = v93
	goto L18
L37:
	;
	v157 = F_memtest_compare(m, l0, l1, l3)
	mBase = m.M
	v158 = m.ExcPending
	if v158 != 0 {
		goto L8
	} else {
		goto L38
	}
L38:
	;
	v161 = F_iprintf(m, int32(_a742), int32(0))
	mBase = m.M
	v162 = m.ExcPending
	if v162 != 0 {
		goto L8
	} else {
		goto L39
	}
L39:
	;
	v163 = v157 + v93
	v165 = v94 + int32(1)
	if v165 != int32(4) {
		v93 = v163
		v94 = v165
		goto L20
	} else {
		goto L40
	}
L40:
	;
	goto L21
L41:
	;
	F_memtest_fill_value(m, l0, l1, int32(0), int32(-1), int32(83), l3)
	mBase = m.M
	v175 = m.ExcPending
	if v175 != 0 {
		goto L8
	} else {
		goto L42
	}
L42:
	;
	v178 = F_iprintf(m, int32(_a742), int32(0))
	mBase = m.M
	v179 = m.ExcPending
	if v179 != 0 {
		goto L8
	} else {
		goto L43
	}
L43:
	;
	v220 = v163
	goto L17
L44:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L45:
	;
	v220 = v198
	goto L17
L46:
	;
	v530 = v517 + v522 + (v524 + (v220 + (v77 + v36)))
	if v48 != l2 {
		v36 = v530
		v40 = v48
		goto L3
	} else {
		goto L110
	}
L47:
	;
	v248 = v240
	goto L49
L48:
	;
	v504 = int32(0)
	F_memtest_fill_value(m, l0, l1, int32(-1431655766), int32(1431655765), int32(67), v504)
	mBase = m.M
	v510 = m.ExcPending
	if v510 != 0 {
		goto L8
	} else {
		goto L109
	}
L49:
	;
	if l3 != 0 {
		goto L52
	} else {
		goto L53
	}
L50:
	;
	goto L48
L51:
	;
	v501 = v248 + int32(1)
	if v501 != int32(4) {
		v248 = v501
		goto L49
	} else {
		goto L108
	}
L52:
	;
	F_memtest_progress_start(m, int32(_a747), v48)
	mBase = m.M
	v438 = m.ExcPending
	if v438 != 0 {
		goto L8
	} else {
		goto L89
	}
L53:
	;
	if v20 != 0 {
		goto L54
	} else {
		goto L55
	}
L54:
	;
	F__serverAssert(m, int32(_a744), int32(_a745), int32(199))
	mBase = m.M
	v434 = m.ExcPending
	if v434 != 0 {
		goto L8
	} else {
		goto L88
	}
L55:
	;
	if base.Ui32(l1) < base.Ui32(int32(8)) {
		goto L51
	} else {
		goto L56
	}
L56:
	;
	v265 = v25
	v270 = l0
	v274 = int64(0)
	goto L58
L57:
	;
	v291 = v288 + v239
	v293 = v240 + int32(1)
	if v293 != int32(4) {
		v239 = v291
		v240 = v293
		goto L47
	} else {
		goto L63
	}
L58:
	;
	v276 = *(*int32)(unsafe.Add(mBase, uint32(v270)))
	v277 = *(*int32)(unsafe.Add(mBase, uint32(v265)))
	if v276 == v277 {
		goto L60
	} else {
		goto L61
	}
L59:
	;
	v288 = int32(0)
	goto L57
L60:
	;
	v280 = int32(4)
	v285 = v274 + int64(1)
	if v285 != v26 {
		v265 = v265 + v280
		v270 = v270 + v280
		v274 = v285
		goto L58
	} else {
		goto L62
	}
L61:
	;
	v288 = int32(1)
	goto L57
L62:
	;
	goto L59
L63:
	;
	F_memtest_fill_value(m, l0, l1, int32(-1431655766), int32(1431655765), int32(67), int32(0))
	mBase = m.M
	v301 = m.ExcPending
	if v301 != 0 {
		goto L8
	} else {
		goto L64
	}
L64:
	;
	v308 = v25
	v313 = l0
	v317 = int64(0)
	goto L66
L65:
	;
	v341 = v25
	v346 = l0
	v350 = int64(0)
	goto L72
L66:
	;
	v319 = *(*int32)(unsafe.Add(mBase, uint32(v313)))
	v320 = *(*int32)(unsafe.Add(mBase, uint32(v308)))
	if v319 == v320 {
		goto L68
	} else {
		goto L69
	}
L67:
	;
	v333 = int32(0)
	goto L65
L68:
	;
	v323 = int32(4)
	v328 = v317 + int64(1)
	if v328 != v26 {
		v308 = v308 + v323
		v313 = v313 + v323
		v317 = v328
		goto L66
	} else {
		goto L70
	}
L69:
	;
	v333 = int32(1)
	goto L65
L70:
	;
	goto L67
L71:
	;
	v374 = v25
	v379 = l0
	v383 = int64(0)
	goto L78
L72:
	;
	v352 = *(*int32)(unsafe.Add(mBase, uint32(v346)))
	v353 = *(*int32)(unsafe.Add(mBase, uint32(v341)))
	if v352 == v353 {
		goto L74
	} else {
		goto L75
	}
L73:
	;
	v364 = int32(0)
	goto L71
L74:
	;
	v356 = int32(4)
	v361 = v350 + int64(1)
	if v361 != v26 {
		v341 = v341 + v356
		v346 = v346 + v356
		v350 = v361
		goto L72
	} else {
		goto L76
	}
L75:
	;
	v364 = int32(1)
	goto L71
L76:
	;
	goto L73
L77:
	;
	v400 = v397 + (v364 + v333)
	v415 = v25
	v416 = int64(0)
	v417 = l0
	goto L83
L78:
	;
	v385 = *(*int32)(unsafe.Add(mBase, uint32(v379)))
	v386 = *(*int32)(unsafe.Add(mBase, uint32(v374)))
	if v385 == v386 {
		goto L80
	} else {
		goto L81
	}
L79:
	;
	v397 = int32(0)
	goto L77
L80:
	;
	v389 = int32(4)
	v394 = v383 + int64(1)
	if v394 != v26 {
		v374 = v374 + v389
		v379 = v379 + v389
		v383 = v394
		goto L78
	} else {
		goto L82
	}
L81:
	;
	v397 = int32(1)
	goto L77
L82:
	;
	goto L79
L83:
	;
	v418 = *(*int32)(unsafe.Add(mBase, uint32(v417)))
	v419 = *(*int32)(unsafe.Add(mBase, uint32(v415)))
	if v418 == v419 {
		goto L85
	} else {
		goto L86
	}
L85:
	;
	v422 = int32(4)
	v428 = v416 + int64(1)
	if v428 != v26 {
		v415 = v415 + v422
		v416 = v428
		v417 = v417 + v422
		goto L83
	} else {
		goto L87
	}
L86:
	;
	v517 = int32(1)
	v522 = v400
	v524 = v291
	goto L46
L87:
	;
	v517 = int32(0)
	v522 = v400
	v524 = v291
	goto L46
L88:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L89:
	;
	v439 = F_memtest_compare(m, l0, l1, l3)
	mBase = m.M
	v440 = m.ExcPending
	if v440 != 0 {
		goto L8
	} else {
		goto L90
	}
L90:
	;
	v443 = F_iprintf(m, int32(_a742), int32(0))
	mBase = m.M
	v444 = m.ExcPending
	if v444 != 0 {
		goto L8
	} else {
		goto L91
	}
L91:
	;
	v445 = v439 + v239
	v447 = v240 + int32(1)
	if v447 != int32(4) {
		v239 = v445
		v240 = v447
		goto L47
	} else {
		goto L92
	}
L92:
	;
	F_memtest_progress_start(m, int32(_a748), v48)
	mBase = m.M
	v452 = m.ExcPending
	if v452 != 0 {
		goto L8
	} else {
		goto L93
	}
L93:
	;
	F_memtest_fill_value(m, l0, l1, int32(-1431655766), int32(1431655765), int32(67), l3)
	mBase = m.M
	v457 = m.ExcPending
	if v457 != 0 {
		goto L8
	} else {
		goto L94
	}
L94:
	;
	v460 = F_iprintf(m, int32(_a742), int32(0))
	mBase = m.M
	v461 = m.ExcPending
	if v461 != 0 {
		goto L8
	} else {
		goto L95
	}
L95:
	;
	F_memtest_progress_start(m, int32(_a747), v48)
	mBase = m.M
	v464 = m.ExcPending
	if v464 != 0 {
		goto L8
	} else {
		goto L96
	}
L96:
	;
	v465 = F_memtest_compare(m, l0, l1, l3)
	mBase = m.M
	v466 = m.ExcPending
	if v466 != 0 {
		goto L8
	} else {
		goto L97
	}
L97:
	;
	v469 = F_iprintf(m, int32(_a742), int32(0))
	mBase = m.M
	v470 = m.ExcPending
	if v470 != 0 {
		goto L8
	} else {
		goto L98
	}
L98:
	;
	F_memtest_progress_start(m, int32(_a747), v48)
	mBase = m.M
	v473 = m.ExcPending
	if v473 != 0 {
		goto L8
	} else {
		goto L99
	}
L99:
	;
	v474 = F_memtest_compare(m, l0, l1, l3)
	mBase = m.M
	v475 = m.ExcPending
	if v475 != 0 {
		goto L8
	} else {
		goto L100
	}
L100:
	;
	v478 = F_iprintf(m, int32(_a742), int32(0))
	mBase = m.M
	v479 = m.ExcPending
	if v479 != 0 {
		goto L8
	} else {
		goto L101
	}
L101:
	;
	F_memtest_progress_start(m, int32(_a747), v48)
	mBase = m.M
	v482 = m.ExcPending
	if v482 != 0 {
		goto L8
	} else {
		goto L102
	}
L102:
	;
	v483 = F_memtest_compare(m, l0, l1, l3)
	mBase = m.M
	v484 = m.ExcPending
	if v484 != 0 {
		goto L8
	} else {
		goto L103
	}
L103:
	;
	v487 = F_iprintf(m, int32(_a742), int32(0))
	mBase = m.M
	v488 = m.ExcPending
	if v488 != 0 {
		goto L8
	} else {
		goto L104
	}
L104:
	;
	F_memtest_progress_start(m, int32(_a747), v48)
	mBase = m.M
	v491 = m.ExcPending
	if v491 != 0 {
		goto L8
	} else {
		goto L105
	}
L105:
	;
	v492 = F_memtest_compare(m, l0, l1, l3)
	mBase = m.M
	v493 = m.ExcPending
	if v493 != 0 {
		goto L8
	} else {
		goto L106
	}
L106:
	;
	v496 = F_iprintf(m, int32(_a742), int32(0))
	mBase = m.M
	v497 = m.ExcPending
	if v497 != 0 {
		goto L8
	} else {
		goto L107
	}
L107:
	;
	v517 = v492
	v522 = v483 + (v474 + v465)
	v524 = v445
	goto L46
L108:
	;
	goto L50
L109:
	;
	v517 = int32(0)
	v522 = v504
	v524 = v239
	goto L46
L110:
	;
	goto L4
}
