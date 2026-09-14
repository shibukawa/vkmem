package p0

import (
	base "github.com/shibukawa/vkmem/internal/aot/vkaot/base"
	"unsafe"
)

func F_geohashAlign52Bits(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int64
	_ = v2
	var v4 int64
	_ = v4
	v2 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	v4 = int64(*(*uint8)(unsafe.Add(mBase, uint32(l0)+8)))
	return v2 << (uint((int64(52)-v4<<(uint(int64(1))%64))&int64(4294967294)) % 64)
}
func F_geohashBoundingBox(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v19 float64
	_ = v19
	var v21 int32
	_ = v21
	var v24 float64
	_ = v24
	var v25 float64
	_ = v25
	var v27 float64
	_ = v27
	var v29 float64
	_ = v29
	var v33 int32
	_ = v33
	var v36 float64
	_ = v36
	var v43 int32
	_ = v43
	var v49 float64
	_ = v49
	var v57 float64
	_ = v57
	var v58 float64
	_ = v58
	var v59 float64
	_ = v59
	var v60 float64
	_ = v60
	var v61 float64
	_ = v61
	var v62 float64
	_ = v62
	var v63 float64
	_ = v63
	var v65 int32
	_ = v65
	var v70 int32
	_ = v70
	var v71 float64
	_ = v71
	var v73 float64
	_ = v73
	var v75 float64
	_ = v75
	var v76 float64
	_ = v76
	var v78 float64
	_ = v78
	var v80 float64
	_ = v80
	var v82 float64
	_ = v82
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v95 int32
	_ = v95
	var v102 float64
	_ = v102
	var v106 int32
	_ = v106
	var v107 float64
	_ = v107
	var v108 float64
	_ = v108
	var v111 float64
	_ = v111
	var v113 float64
	_ = v113
	var v115 float64
	_ = v115
	var v118 float64
	_ = v118
	var v121 float64
	_ = v121
	var v126 float64
	_ = v126
	var v130 int32
	_ = v130
	var v132 int32
	_ = v132
	var v139 int32
	_ = v139
	var v146 float64
	_ = v146
	var v150 int32
	_ = v150
	var v151 float64
	_ = v151
	var v152 float64
	_ = v152
	var v156 float64
	_ = v156
	var v157 float64
	_ = v157
	var v159 float64
	_ = v159
	var v161 float64
	_ = v161
	var v163 float64
	_ = v163
	var v170 float64
	_ = v170
	var v174 int32
	_ = v174
	var v176 int32
	_ = v176
	var v183 int32
	_ = v183
	var v190 float64
	_ = v190
	var v194 int32
	_ = v194
	var v195 float64
	_ = v195
	var v196 float64
	_ = v196
	var v199 float64
	_ = v199
	var v201 float64
	_ = v201
	var v203 float64
	_ = v203
	var v206 float64
	_ = v206
	var v209 float64
	_ = v209
	var v214 float64
	_ = v214
	var v218 int32
	_ = v218
	var v220 int32
	_ = v220
	var v227 int32
	_ = v227
	var v234 float64
	_ = v234
	var v238 int32
	_ = v238
	var v239 float64
	_ = v239
	var v240 float64
	_ = v240
	var v244 float64
	_ = v244
	var v245 float64
	_ = v245
	var v247 float64
	_ = v247
	var v249 float64
	_ = v249
	var v251 float64
	_ = v251
	var v257 float64
	_ = v257
	var v259 int32
	_ = v259
	var v261 float64
	_ = v261
	var v262 float64
	_ = v262
	var v263 float64
	_ = v263
	var v264 float64
	_ = v264
	var v265 float64
	_ = v265
	var v266 float64
	_ = v266
	var v267 float64
	_ = v267
	var v270 float64
	_ = v270
	var v271 float64
	_ = v271
	var v272 float64
	_ = v272
	var v274 float64
	_ = v274
	var v277 float64
	_ = v277
	var v281 int32
	_ = v281
	var v283 int32
	_ = v283
	var v290 int32
	_ = v290
	var v297 float64
	_ = v297
	var v301 int32
	_ = v301
	var v302 float64
	_ = v302
	var v303 float64
	_ = v303
	var v306 float64
	_ = v306
	var v308 float64
	_ = v308
	var v310 float64
	_ = v310
	var v313 float64
	_ = v313
	var v316 float64
	_ = v316
	var v323 float64
	_ = v323
	var v327 int32
	_ = v327
	var v329 int32
	_ = v329
	var v336 int32
	_ = v336
	var v343 float64
	_ = v343
	var v347 int32
	_ = v347
	var v348 float64
	_ = v348
	var v349 float64
	_ = v349
	var v352 float64
	_ = v352
	var v354 float64
	_ = v354
	var v356 float64
	_ = v356
	var v359 float64
	_ = v359
	var v362 float64
	_ = v362
	var v368 float64
	_ = v368
	var v371 float64
	_ = v371
	var v383 float64
	_ = v383
	var v384 float64
	_ = v384
	var v385 float64
	_ = v385
	var v386 float64
	_ = v386
	var v387 float64
	_ = v387
	var v388 float64
	_ = v388
	var v389 float64
	_ = v389
	var v405 int64
	_ = v405
	var v410 int64
	_ = v410
	var v416 int64
	_ = v416
	var v419 int32
	_ = v419
	var v422 int32
	_ = v422
	var v424 float64
	_ = v424
	var v428 int32
	_ = v428
	var v429 int64
	_ = v429
	var v433 int32
	_ = v433
	var v438 int32
	_ = v438
	var v444 int32
	_ = v444
	var v456 float64
	_ = v456
	var v472 float64
	_ = v472
	var v473 float64
	_ = v473
	var v474 float64
	_ = v474
	var v488 float64
	_ = v488
	var v490 float64
	_ = v490
	var v498 float64
	_ = v498
	var v505 float64
	_ = v505
	var v513 int64
	_ = v513
	var v518 int64
	_ = v518
	var v524 int64
	_ = v524
	var v527 int32
	_ = v527
	var v530 int32
	_ = v530
	var v532 float64
	_ = v532
	var v536 int32
	_ = v536
	var v537 int64
	_ = v537
	var v541 int32
	_ = v541
	var v546 int32
	_ = v546
	var v552 int32
	_ = v552
	var v564 float64
	_ = v564
	var v580 float64
	_ = v580
	var v581 float64
	_ = v581
	var v582 float64
	_ = v582
	var v596 float64
	_ = v596
	var v598 float64
	_ = v598
	var v606 float64
	_ = v606
	if l1 != 0 {
		v19 = float64(0)
		v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		switch v21 + int32(-1) {
		case 0:
			v261 = *(*float64)(unsafe.Add(mBase, uint32(l0)+24))
			v262 = *(*float64)(unsafe.Add(mBase, uint32(l0)+64))
			v263 = base.F64_mul(v261, v262)
			v264 = v263
			v265 = v263
			v266 = *(*float64)(unsafe.Add(mBase, uint32(l0)+8))
			v267 = *(*float64)(unsafe.Add(mBase, uint32(l0)+16))
			v270 = float64(0.017453292519943295)
			v271 = base.F64_div(base.F64_div(v265, float64(6.372797560856e+06)), v270)
			v272 = base.F64_add(v267, v271)
			*(*float64)(unsafe.Add(mBase, uint32(l1)+24)) = v272
			v274 = base.F64_sub(v267, v271)
			*(*float64)(unsafe.Add(mBase, uint32(l1)+8)) = v274
			v277 = base.F64_mul(v272, v270)
			v281 = m.G0
			v283 = v281 - int32(16)
			m.G0 = v283
			v290 = base.I32_wrap_i64(int64(base.Ui64(base.I64_reinterpret_f64(v277))>>(uint(int64(32))%64))) & int32(2147483647)
			if base.Ui32(int32(1072243195)) < base.Ui32(v290) {
				if base.Ui32(v290) < base.Ui32(int32(2146435072)) {
					v301 = F___rem_pio2(m, v277, v283)
					mBase = m.M
					v302 = *(*float64)(unsafe.Add(mBase, uint32(v283)+8))
					v303 = *(*float64)(unsafe.Add(mBase, uint32(v283)))
					switch v301 & int32(3) {
					default:
						v306 = F___cos(m, v303, v302)
						mBase = m.M
						v316 = v306
					case 1:
						v308 = F___sin(m, v303, v302, int32(1))
						mBase = m.M
						v316 = base.F64_neg(v308)
					case 2:
						v310 = F___cos(m, v303, v302)
						mBase = m.M
						v316 = base.F64_neg(v310)
					case 3:
						v313 = F___sin(m, v303, v302, int32(1))
						mBase = m.M
						v316 = v313
					}
				} else {
					v316 = base.F64_sub(v277, v277)
				}
			} else {
				if base.Ui32(v290) < base.Ui32(int32(1044816030)) {
					v316 = float64(1)
				} else {
					v297 = F___cos(m, v277, float64(0))
					mBase = m.M
					v316 = v297
				}
			}
			m.G0 = v283 + int32(16)
			v323 = base.F64_mul(v274, float64(0.017453292519943295))
			v327 = m.G0
			v329 = v327 - int32(16)
			m.G0 = v329
			v336 = base.I32_wrap_i64(int64(base.Ui64(base.I64_reinterpret_f64(v323))>>(uint(int64(32))%64))) & int32(2147483647)
			if base.Ui32(int32(1072243195)) < base.Ui32(v336) {
				if base.Ui32(v336) < base.Ui32(int32(2146435072)) {
					v347 = F___rem_pio2(m, v323, v329)
					mBase = m.M
					v348 = *(*float64)(unsafe.Add(mBase, uint32(v329)+8))
					v349 = *(*float64)(unsafe.Add(mBase, uint32(v329)))
					switch v347 & int32(3) {
					default:
						v352 = F___cos(m, v349, v348)
						mBase = m.M
						v362 = v352
					case 1:
						v354 = F___sin(m, v349, v348, int32(1))
						mBase = m.M
						v362 = base.F64_neg(v354)
					case 2:
						v356 = F___cos(m, v349, v348)
						mBase = m.M
						v362 = base.F64_neg(v356)
					case 3:
						v359 = F___sin(m, v349, v348, int32(1))
						mBase = m.M
						v362 = v359
					}
				} else {
					v362 = base.F64_sub(v323, v323)
				}
			} else {
				if base.Ui32(v336) < base.Ui32(int32(1044816030)) {
					v362 = float64(1)
				} else {
					v343 = F___cos(m, v323, float64(0))
					mBase = m.M
					v362 = v343
				}
			}
			m.G0 = v329 + int32(16)
			if base.F64_lt(v267, float64(0)) != 0 {
				v368 = v362
			} else {
				v368 = v316
			}
			v371 = base.F64_div(base.F64_div(base.F64_div(v264, float64(6.372797560856e+06)), v368), float64(0.017453292519943295))
			*(*float64)(unsafe.Add(mBase, uint32(l1)+16)) = base.F64_add(v266, v371)
			*(*float64)(unsafe.Add(mBase, uint32(l1))) = base.F64_sub(v266, v371)
			return int32(1)
		case 1:
			v24 = *(*float64)(unsafe.Add(mBase, uint32(l0)+24))
			v25 = *(*float64)(unsafe.Add(mBase, uint32(l0)+72))
			v27 = float64(0.5)
			v29 = *(*float64)(unsafe.Add(mBase, uint32(l0)+64))
			v264 = base.F64_mul(base.F64_mul(v24, v25), v27)
			v265 = base.F64_mul(base.F64_mul(v24, v29), v27)
			v266 = *(*float64)(unsafe.Add(mBase, uint32(l0)+8))
			v267 = *(*float64)(unsafe.Add(mBase, uint32(l0)+16))
			v270 = float64(0.017453292519943295)
			v271 = base.F64_div(base.F64_div(v265, float64(6.372797560856e+06)), v270)
			v272 = base.F64_add(v267, v271)
			*(*float64)(unsafe.Add(mBase, uint32(l1)+24)) = v272
			v274 = base.F64_sub(v267, v271)
			*(*float64)(unsafe.Add(mBase, uint32(l1)+8)) = v274
			v277 = base.F64_mul(v272, v270)
			v281 = m.G0
			v283 = v281 - int32(16)
			m.G0 = v283
			v290 = base.I32_wrap_i64(int64(base.Ui64(base.I64_reinterpret_f64(v277))>>(uint(int64(32))%64))) & int32(2147483647)
			if base.Ui32(int32(1072243195)) < base.Ui32(v290) {
				if base.Ui32(v290) < base.Ui32(int32(2146435072)) {
					v301 = F___rem_pio2(m, v277, v283)
					mBase = m.M
					v302 = *(*float64)(unsafe.Add(mBase, uint32(v283)+8))
					v303 = *(*float64)(unsafe.Add(mBase, uint32(v283)))
					switch v301 & int32(3) {
					default:
						v306 = F___cos(m, v303, v302)
						mBase = m.M
						v316 = v306
					case 1:
						v308 = F___sin(m, v303, v302, int32(1))
						mBase = m.M
						v316 = base.F64_neg(v308)
					case 2:
						v310 = F___cos(m, v303, v302)
						mBase = m.M
						v316 = base.F64_neg(v310)
					case 3:
						v313 = F___sin(m, v303, v302, int32(1))
						mBase = m.M
						v316 = v313
					}
				} else {
					v316 = base.F64_sub(v277, v277)
				}
			} else {
				if base.Ui32(v290) < base.Ui32(int32(1044816030)) {
					v316 = float64(1)
				} else {
					v297 = F___cos(m, v277, float64(0))
					mBase = m.M
					v316 = v297
				}
			}
			m.G0 = v283 + int32(16)
			v323 = base.F64_mul(v274, float64(0.017453292519943295))
			v327 = m.G0
			v329 = v327 - int32(16)
			m.G0 = v329
			v336 = base.I32_wrap_i64(int64(base.Ui64(base.I64_reinterpret_f64(v323))>>(uint(int64(32))%64))) & int32(2147483647)
			if base.Ui32(int32(1072243195)) < base.Ui32(v336) {
				if base.Ui32(v336) < base.Ui32(int32(2146435072)) {
					v347 = F___rem_pio2(m, v323, v329)
					mBase = m.M
					v348 = *(*float64)(unsafe.Add(mBase, uint32(v329)+8))
					v349 = *(*float64)(unsafe.Add(mBase, uint32(v329)))
					switch v347 & int32(3) {
					default:
						v352 = F___cos(m, v349, v348)
						mBase = m.M
						v362 = v352
					case 1:
						v354 = F___sin(m, v349, v348, int32(1))
						mBase = m.M
						v362 = base.F64_neg(v354)
					case 2:
						v356 = F___cos(m, v349, v348)
						mBase = m.M
						v362 = base.F64_neg(v356)
					case 3:
						v359 = F___sin(m, v349, v348, int32(1))
						mBase = m.M
						v362 = v359
					}
				} else {
					v362 = base.F64_sub(v323, v323)
				}
			} else {
				if base.Ui32(v336) < base.Ui32(int32(1044816030)) {
					v362 = float64(1)
				} else {
					v343 = F___cos(m, v323, float64(0))
					mBase = m.M
					v362 = v343
				}
			}
			m.G0 = v329 + int32(16)
			if base.F64_lt(v267, float64(0)) != 0 {
				v368 = v362
			} else {
				v368 = v316
			}
			v371 = base.F64_div(base.F64_div(base.F64_div(v264, float64(6.372797560856e+06)), v368), float64(0.017453292519943295))
			*(*float64)(unsafe.Add(mBase, uint32(l1)+16)) = base.F64_add(v266, v371)
			*(*float64)(unsafe.Add(mBase, uint32(l1))) = base.F64_sub(v266, v371)
			return int32(1)
		case 2:
			v33 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
			if int32(1) <= v33 {
				v43 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
				v49 = float64(0)
				v57 = v49
				v58 = float64(180)
				v59 = float64(-180)
				v60 = float64(85.05112878)
				v61 = float64(-85.05112878)
				v62 = v49
				v63 = v49
				v65 = int32(0)
				for {
					v70 = v43 + v65<<(uint(int32(4))%32)
					v71 = *(*float64)(unsafe.Add(mBase, uint32(v70)+8))
					if base.F64_gt(v71, v61) != 0 {
						v73 = v71
					} else {
						v73 = v61
					}
					if base.F64_lt(v71, v60) != 0 {
						v75 = v71
					} else {
						v75 = v60
					}
					v76 = *(*float64)(unsafe.Add(mBase, uint32(v70)))
					if base.F64_gt(v76, v59) != 0 {
						v78 = v76
					} else {
						v78 = v59
					}
					if base.F64_lt(v76, v58) != 0 {
						v80 = v76
					} else {
						v80 = v58
					}
					v82 = base.F64_mul(v71, float64(0.017453292519943295))
					v86 = m.G0
					v88 = v86 - int32(16)
					m.G0 = v88
					v95 = base.I32_wrap_i64(int64(base.Ui64(base.I64_reinterpret_f64(v82))>>(uint(int64(32))%64))) & int32(2147483647)
					if base.Ui32(int32(1072243195)) < base.Ui32(v95) {
						if base.Ui32(v95) < base.Ui32(int32(2146435072)) {
							v106 = F___rem_pio2(m, v82, v88)
							mBase = m.M
							v107 = *(*float64)(unsafe.Add(mBase, uint32(v88)+8))
							v108 = *(*float64)(unsafe.Add(mBase, uint32(v88)))
							switch v106 & int32(3) {
							default:
								v111 = F___cos(m, v108, v107)
								mBase = m.M
								v121 = v111
							case 1:
								v113 = F___sin(m, v108, v107, int32(1))
								mBase = m.M
								v121 = base.F64_neg(v113)
							case 2:
								v115 = F___cos(m, v108, v107)
								mBase = m.M
								v121 = base.F64_neg(v115)
							case 3:
								v118 = F___sin(m, v108, v107, int32(1))
								mBase = m.M
								v121 = v118
							}
						} else {
							v121 = base.F64_sub(v82, v82)
						}
					} else {
						if base.Ui32(v95) < base.Ui32(int32(1044816030)) {
							v121 = float64(1)
						} else {
							v102 = F___cos(m, v82, float64(0))
							mBase = m.M
							v121 = v102
						}
					}
					m.G0 = v88 + int32(16)
					v126 = base.F64_mul(v76, float64(0.017453292519943295))
					v130 = m.G0
					v132 = v130 - int32(16)
					m.G0 = v132
					v139 = base.I32_wrap_i64(int64(base.Ui64(base.I64_reinterpret_f64(v126))>>(uint(int64(32))%64))) & int32(2147483647)
					if base.Ui32(int32(1072243195)) < base.Ui32(v139) {
						if base.Ui32(v139) < base.Ui32(int32(2146435072)) {
							v150 = F___rem_pio2(m, v126, v132)
							mBase = m.M
							v151 = *(*float64)(unsafe.Add(mBase, uint32(v132)+8))
							v152 = *(*float64)(unsafe.Add(mBase, uint32(v132)))
							switch v150 & int32(3) {
							default:
								v156 = F___sin(m, v152, v151, int32(1))
								mBase = m.M
								v163 = v156
							case 1:
								v157 = F___cos(m, v152, v151)
								mBase = m.M
								v163 = v157
							case 2:
								v159 = F___sin(m, v152, v151, int32(1))
								mBase = m.M
								v163 = base.F64_neg(v159)
							case 3:
								v161 = F___cos(m, v152, v151)
								mBase = m.M
								v163 = base.F64_neg(v161)
							}
						} else {
							v163 = base.F64_sub(v126, v126)
						}
					} else {
						if base.Ui32(v139) < base.Ui32(int32(1045430272)) {
							v163 = v126
						} else {
							v146 = F___sin(m, v126, float64(0), int32(0))
							mBase = m.M
							v163 = v146
						}
					}
					m.G0 = v132 + int32(16)
					v170 = base.F64_add(v62, base.F64_mul(v121, v163))
					v174 = m.G0
					v176 = v174 - int32(16)
					m.G0 = v176
					v183 = base.I32_wrap_i64(int64(base.Ui64(base.I64_reinterpret_f64(v126))>>(uint(int64(32))%64))) & int32(2147483647)
					if base.Ui32(int32(1072243195)) < base.Ui32(v183) {
						if base.Ui32(v183) < base.Ui32(int32(2146435072)) {
							v194 = F___rem_pio2(m, v126, v176)
							mBase = m.M
							v195 = *(*float64)(unsafe.Add(mBase, uint32(v176)+8))
							v196 = *(*float64)(unsafe.Add(mBase, uint32(v176)))
							switch v194 & int32(3) {
							default:
								v199 = F___cos(m, v196, v195)
								mBase = m.M
								v209 = v199
							case 1:
								v201 = F___sin(m, v196, v195, int32(1))
								mBase = m.M
								v209 = base.F64_neg(v201)
							case 2:
								v203 = F___cos(m, v196, v195)
								mBase = m.M
								v209 = base.F64_neg(v203)
							case 3:
								v206 = F___sin(m, v196, v195, int32(1))
								mBase = m.M
								v209 = v206
							}
						} else {
							v209 = base.F64_sub(v126, v126)
						}
					} else {
						if base.Ui32(v183) < base.Ui32(int32(1044816030)) {
							v209 = float64(1)
						} else {
							v190 = F___cos(m, v126, float64(0))
							mBase = m.M
							v209 = v190
						}
					}
					m.G0 = v176 + int32(16)
					v214 = base.F64_add(v63, base.F64_mul(v121, v209))
					v218 = m.G0
					v220 = v218 - int32(16)
					m.G0 = v220
					v227 = base.I32_wrap_i64(int64(base.Ui64(base.I64_reinterpret_f64(v82))>>(uint(int64(32))%64))) & int32(2147483647)
					if base.Ui32(int32(1072243195)) < base.Ui32(v227) {
						if base.Ui32(v227) < base.Ui32(int32(2146435072)) {
							v238 = F___rem_pio2(m, v82, v220)
							mBase = m.M
							v239 = *(*float64)(unsafe.Add(mBase, uint32(v220)+8))
							v240 = *(*float64)(unsafe.Add(mBase, uint32(v220)))
							switch v238 & int32(3) {
							default:
								v244 = F___sin(m, v240, v239, int32(1))
								mBase = m.M
								v251 = v244
							case 1:
								v245 = F___cos(m, v240, v239)
								mBase = m.M
								v251 = v245
							case 2:
								v247 = F___sin(m, v240, v239, int32(1))
								mBase = m.M
								v251 = base.F64_neg(v247)
							case 3:
								v249 = F___cos(m, v240, v239)
								mBase = m.M
								v251 = base.F64_neg(v249)
							}
						} else {
							v251 = base.F64_sub(v82, v82)
						}
					} else {
						if base.Ui32(v227) < base.Ui32(int32(1045430272)) {
							v251 = v82
						} else {
							v234 = F___sin(m, v82, float64(0), int32(0))
							mBase = m.M
							v251 = v234
						}
					}
					m.G0 = v220 + int32(16)
					v257 = base.F64_add(v57, v251)
					v259 = v65 + int32(1)
					if v259 != v33 {
						v57 = v257
						v58 = v80
						v59 = v78
						v60 = v75
						v61 = v73
						v62 = v170
						v63 = v214
						v65 = v259
						continue
					} else {
						break
					}
					break
				}
				v383 = v257
				v384 = v80
				v385 = v78
				v386 = v75
				v387 = v73
				v388 = v170
				v389 = v214
			} else {
				v36 = float64(0)
				v383 = v36
				v384 = float64(180)
				v385 = float64(-180)
				v386 = float64(85.05112878)
				v387 = float64(-85.05112878)
				v388 = v36
				v389 = v36
			}
			*(*float64)(unsafe.Add(mBase, uint32(l1)+24)) = v387
			*(*float64)(unsafe.Add(mBase, uint32(l1)+16)) = v385
			*(*float64)(unsafe.Add(mBase, uint32(l1)+8)) = v386
			*(*float64)(unsafe.Add(mBase, uint32(l1))) = v384
			v405 = F___DOUBLE_BITS_2(m, v389)
			mBase = m.M
			if base.Ui64(int64(9218868437227405312)) < base.Ui64(v405&int64(9223372036854775807)) {
				v498 = base.F64_add(v388, v389)
			} else {
				v410 = F___DOUBLE_BITS_2(m, v388)
				mBase = m.M
				if base.Ui64(v410&int64(9223372036854775807)) < base.Ui64(int64(9218868437227405313)) {
					v416 = base.I64_reinterpret_f64(v389)
					v419 = base.I32_wrap_i64(int64(base.Ui64(v416) >> (uint(int64(32)) % 64)))
					v422 = base.I32_wrap_i64(v416)
					if v419+int32(-1072693248)|v422 != 0 {
						v428 = int32(base.Ui32(v419)>>(uint(int32(30))%32)) & int32(2)
						v429 = base.I64_reinterpret_f64(v388)
						v433 = v428 | base.I32_wrap_i64(int64(base.Ui64(v429)>>(uint(int64(63))%64)))
						v438 = base.I32_wrap_i64(int64(base.Ui64(v429)>>(uint(int64(32))%64))) & int32(2147483647)
						if v438|base.I32_wrap_i64(v429) != 0 {
							v444 = v419 & int32(2147483647)
							if v444|v422 != 0 {
								if v444 != int32(2146435072) {
									if v438 == int32(2146435072) {
										v498 = base.F64_copysign(float64(1.5707963267948966), v388)
									} else {
										if base.Ui32(v438) <= base.Ui32(v444+int32(67108864)) {
											if v428 == int32(0) {
												v472 = F_fabs(m, base.F64_div(v388, v389))
												mBase = m.M
												v473 = F_atan(m, v472)
												mBase = m.M
												v474 = v473
											} else {
												if base.Ui32(v438+int32(67108864)) < base.Ui32(v444) {
													v474 = float64(0)
												} else {
													v472 = F_fabs(m, base.F64_div(v388, v389))
													mBase = m.M
													v473 = F_atan(m, v472)
													mBase = m.M
													v474 = v473
												}
											}
											switch v433 {
											default:
												v490 = v474
												v498 = v490
											case 1:
												v498 = base.F64_neg(v474)
											case 2:
												v498 = base.F64_sub(float64(3.141592653589793), base.F64_add(v474, float64(-1.2246467991473532e-16)))
											case 3:
												v498 = base.F64_add(base.F64_add(v474, float64(-1.2246467991473532e-16)), float64(-3.141592653589793))
											}
										} else {
											v498 = base.F64_copysign(float64(1.5707963267948966), v388)
										}
									}
								} else {
									if v438 != int32(2146435072) {
										v488 = *(*float64)(unsafe.Add(mBase, uint32(v433<<(uint(int32(3))%32))+uint32(_consts[290])))
										v490 = v488
										v498 = v490
									} else {
										v456 = *(*float64)(unsafe.Add(mBase, uint32(v433<<(uint(int32(3))%32))+uint32(_consts[291])))
										v498 = v456
									}
								}
							} else {
								v498 = base.F64_copysign(float64(1.5707963267948966), v388)
							}
						} else {
							switch v433 {
							default:
								v490 = v388
								v498 = v490
							case 2:
								v498 = float64(3.141592653589793)
							case 3:
								v498 = float64(-3.141592653589793)
							}
						}
					} else {
						v424 = F_atan(m, v388)
						mBase = m.M
						v498 = v424
					}
				} else {
					v498 = base.F64_add(v388, v389)
				}
			}
			*(*float64)(unsafe.Add(mBase, uint32(l0)+8)) = base.F64_div(v498, float64(0.017453292519943295))
			v505 = base.F64_sqrt(base.F64_add(base.F64_mul(v389, v389), base.F64_mul(v388, v388)))
			v513 = F___DOUBLE_BITS_2(m, v505)
			mBase = m.M
			if base.Ui64(int64(9218868437227405312)) < base.Ui64(v513&int64(9223372036854775807)) {
				v606 = base.F64_add(v383, v505)
			} else {
				v518 = F___DOUBLE_BITS_2(m, v383)
				mBase = m.M
				if base.Ui64(v518&int64(9223372036854775807)) < base.Ui64(int64(9218868437227405313)) {
					v524 = base.I64_reinterpret_f64(v505)
					v527 = base.I32_wrap_i64(int64(base.Ui64(v524) >> (uint(int64(32)) % 64)))
					v530 = base.I32_wrap_i64(v524)
					if v527+int32(-1072693248)|v530 != 0 {
						v536 = int32(base.Ui32(v527)>>(uint(int32(30))%32)) & int32(2)
						v537 = base.I64_reinterpret_f64(v383)
						v541 = v536 | base.I32_wrap_i64(int64(base.Ui64(v537)>>(uint(int64(63))%64)))
						v546 = base.I32_wrap_i64(int64(base.Ui64(v537)>>(uint(int64(32))%64))) & int32(2147483647)
						if v546|base.I32_wrap_i64(v537) != 0 {
							v552 = v527 & int32(2147483647)
							if v552|v530 != 0 {
								if v552 != int32(2146435072) {
									if v546 == int32(2146435072) {
										v606 = base.F64_copysign(float64(1.5707963267948966), v383)
									} else {
										if base.Ui32(v546) <= base.Ui32(v552+int32(67108864)) {
											if v536 == int32(0) {
												v580 = F_fabs(m, base.F64_div(v383, v505))
												mBase = m.M
												v581 = F_atan(m, v580)
												mBase = m.M
												v582 = v581
											} else {
												if base.Ui32(v546+int32(67108864)) < base.Ui32(v552) {
													v582 = float64(0)
												} else {
													v580 = F_fabs(m, base.F64_div(v383, v505))
													mBase = m.M
													v581 = F_atan(m, v580)
													mBase = m.M
													v582 = v581
												}
											}
											switch v541 {
											default:
												v598 = v582
												v606 = v598
											case 1:
												v606 = base.F64_neg(v582)
											case 2:
												v606 = base.F64_sub(float64(3.141592653589793), base.F64_add(v582, float64(-1.2246467991473532e-16)))
											case 3:
												v606 = base.F64_add(base.F64_add(v582, float64(-1.2246467991473532e-16)), float64(-3.141592653589793))
											}
										} else {
											v606 = base.F64_copysign(float64(1.5707963267948966), v383)
										}
									}
								} else {
									if v546 != int32(2146435072) {
										v596 = *(*float64)(unsafe.Add(mBase, uint32(v541<<(uint(int32(3))%32))+uint32(_consts[290])))
										v598 = v596
										v606 = v598
									} else {
										v564 = *(*float64)(unsafe.Add(mBase, uint32(v541<<(uint(int32(3))%32))+uint32(_consts[291])))
										v606 = v564
									}
								}
							} else {
								v606 = base.F64_copysign(float64(1.5707963267948966), v383)
							}
						} else {
							switch v541 {
							default:
								v598 = v383
								v606 = v598
							case 2:
								v606 = float64(3.141592653589793)
							case 3:
								v606 = float64(-3.141592653589793)
							}
						}
					} else {
						v532 = F_atan(m, v383)
						mBase = m.M
						v606 = v532
					}
				} else {
					v606 = base.F64_add(v383, v505)
				}
			}
			*(*float64)(unsafe.Add(mBase, uint32(l0)+16)) = base.F64_div(v606, float64(0.017453292519943295))
			return int32(1)
		default:
			v264 = v19
			v265 = v19
			v266 = *(*float64)(unsafe.Add(mBase, uint32(l0)+8))
			v267 = *(*float64)(unsafe.Add(mBase, uint32(l0)+16))
			v270 = float64(0.017453292519943295)
			v271 = base.F64_div(base.F64_div(v265, float64(6.372797560856e+06)), v270)
			v272 = base.F64_add(v267, v271)
			*(*float64)(unsafe.Add(mBase, uint32(l1)+24)) = v272
			v274 = base.F64_sub(v267, v271)
			*(*float64)(unsafe.Add(mBase, uint32(l1)+8)) = v274
			v277 = base.F64_mul(v272, v270)
			v281 = m.G0
			v283 = v281 - int32(16)
			m.G0 = v283
			v290 = base.I32_wrap_i64(int64(base.Ui64(base.I64_reinterpret_f64(v277))>>(uint(int64(32))%64))) & int32(2147483647)
			if base.Ui32(int32(1072243195)) < base.Ui32(v290) {
				if base.Ui32(v290) < base.Ui32(int32(2146435072)) {
					v301 = F___rem_pio2(m, v277, v283)
					mBase = m.M
					v302 = *(*float64)(unsafe.Add(mBase, uint32(v283)+8))
					v303 = *(*float64)(unsafe.Add(mBase, uint32(v283)))
					switch v301 & int32(3) {
					default:
						v306 = F___cos(m, v303, v302)
						mBase = m.M
						v316 = v306
					case 1:
						v308 = F___sin(m, v303, v302, int32(1))
						mBase = m.M
						v316 = base.F64_neg(v308)
					case 2:
						v310 = F___cos(m, v303, v302)
						mBase = m.M
						v316 = base.F64_neg(v310)
					case 3:
						v313 = F___sin(m, v303, v302, int32(1))
						mBase = m.M
						v316 = v313
					}
				} else {
					v316 = base.F64_sub(v277, v277)
				}
			} else {
				if base.Ui32(v290) < base.Ui32(int32(1044816030)) {
					v316 = float64(1)
				} else {
					v297 = F___cos(m, v277, float64(0))
					mBase = m.M
					v316 = v297
				}
			}
			m.G0 = v283 + int32(16)
			v323 = base.F64_mul(v274, float64(0.017453292519943295))
			v327 = m.G0
			v329 = v327 - int32(16)
			m.G0 = v329
			v336 = base.I32_wrap_i64(int64(base.Ui64(base.I64_reinterpret_f64(v323))>>(uint(int64(32))%64))) & int32(2147483647)
			if base.Ui32(int32(1072243195)) < base.Ui32(v336) {
				if base.Ui32(v336) < base.Ui32(int32(2146435072)) {
					v347 = F___rem_pio2(m, v323, v329)
					mBase = m.M
					v348 = *(*float64)(unsafe.Add(mBase, uint32(v329)+8))
					v349 = *(*float64)(unsafe.Add(mBase, uint32(v329)))
					switch v347 & int32(3) {
					default:
						v352 = F___cos(m, v349, v348)
						mBase = m.M
						v362 = v352
					case 1:
						v354 = F___sin(m, v349, v348, int32(1))
						mBase = m.M
						v362 = base.F64_neg(v354)
					case 2:
						v356 = F___cos(m, v349, v348)
						mBase = m.M
						v362 = base.F64_neg(v356)
					case 3:
						v359 = F___sin(m, v349, v348, int32(1))
						mBase = m.M
						v362 = v359
					}
				} else {
					v362 = base.F64_sub(v323, v323)
				}
			} else {
				if base.Ui32(v336) < base.Ui32(int32(1044816030)) {
					v362 = float64(1)
				} else {
					v343 = F___cos(m, v323, float64(0))
					mBase = m.M
					v362 = v343
				}
			}
			m.G0 = v329 + int32(16)
			if base.F64_lt(v267, float64(0)) != 0 {
				v368 = v362
			} else {
				v368 = v316
			}
			v371 = base.F64_div(base.F64_div(base.F64_div(v264, float64(6.372797560856e+06)), v368), float64(0.017453292519943295))
			*(*float64)(unsafe.Add(mBase, uint32(l1)+16)) = base.F64_add(v266, v371)
			*(*float64)(unsafe.Add(mBase, uint32(l1))) = base.F64_sub(v266, v371)
			return int32(1)
		}
	} else {
		return int32(0)
	}
}
func F_geohashDecode(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v17 int64
	_ = v17
	var v20 int32
	_ = v20
	var v26 float64
	_ = v26
	var v27 float64
	_ = v27
	var v32 float64
	_ = v32
	var v33 float64
	_ = v33
	var v38 int64
	_ = v38
	var v40 int32
	_ = v40
	var v44 int64
	_ = v44
	var v46 int64
	_ = v46
	var v52 int64
	_ = v52
	var v58 int64
	_ = v58
	var v59 int64
	_ = v59
	var v60 int64
	_ = v60
	var v63 int64
	_ = v63
	var v64 int64
	_ = v64
	var v65 int64
	_ = v65
	var v67 int64
	_ = v67
	var v68 int64
	_ = v68
	var v70 int64
	_ = v70
	var v72 int64
	_ = v72
	var v75 int32
	_ = v75
	var v82 float64
	_ = v82
	var v84 float64
	_ = v84
	var v89 int64
	_ = v89
	var v94 int64
	_ = v94
	var v99 int64
	_ = v99
	var v104 int64
	_ = v104
	var v107 int64
	_ = v107
	var v115 int32
	_ = v115
	var v118 float64
	_ = v118
	var v122 int32
	_ = v122
	var v139 int32
	_ = v139
	v5 = int32(0)
	if l3 == v5 {
		v139 = v5
	} else {
		v17 = *(*int64)(unsafe.Add(mBase, uint32(l2)))
		v20 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+8)))
		if base.B2i32(v17 == int64(0))&base.B2i32(v20&int32(255) == int32(0)) != 0 {
			v139 = v5
		} else {
			v26 = *(*float64)(unsafe.Add(mBase, uint32(l1)))
			v27 = *(*float64)(unsafe.Add(mBase, uint32(l1)+8))
			if base.F64_ne(v27, float64(0)) != 0 {
				v32 = *(*float64)(unsafe.Add(mBase, uint32(l0)))
				v33 = *(*float64)(unsafe.Add(mBase, uint32(l0)+8))
				if base.F64_ne(v33, float64(0)) != 0 {
					v38 = *(*int64)(unsafe.Add(mBase, uint32(l2)))
					*(*int64)(unsafe.Add(mBase, uint32(l3))) = v38
					v40 = int32(8)
					v44 = *(*int64)(unsafe.Add(mBase, uint32(l2+v40)))
					*(*int64)(unsafe.Add(mBase, uint32(l3+v40))) = v44
					v46 = int64(1)
					v52 = int64(base.Ui64(v17)>>(uint(v46)%64))&int64(4919131752989213764) | v17&int64(2459565876494606882)
					v58 = int64(1085102592571150095)
					v59 = (int64(base.Ui64(v52)>>(uint(v46)%64)) | int64(base.Ui64(v52)>>(uint(int64(3))%64))) & v58
					v60 = int64(4)
					v63 = int64(71777214294589695)
					v64 = (int64(base.Ui64(v59)>>(uint(v60)%64)) | v59) & v63
					v65 = int64(8)
					v67 = int64(base.Ui64(v64)>>(uint(v65)%64)) | v64
					v68 = int64(16)
					v70 = int64(4294901760)
					v72 = int64(65535)
					v75 = base.I32_wrap_i64(int64(base.Ui64(v67)>>(uint(v68)%64))&v70 | v67&v72)
					v82 = base.F64_convert_i64_u(v46 << (uint(base.I64_extend_i32_u(v20)&int64(255)) % 64))
					v84 = base.F64_sub(v33, v32)
					*(*float64)(unsafe.Add(mBase, uint32(l3)+16)) = base.F64_add(base.F64_mul(base.F64_div(base.F64_convert_i32_u(v75), v82), v84), v32)
					v89 = v17 & int64(6148914691236517205)
					v94 = (int64(base.Ui64(v89)>>(uint(v46)%64)) | v89) & int64(3689348814741910323)
					v99 = (int64(base.Ui64(v94)>>(uint(int64(2))%64)) | v94) & v58
					v104 = (int64(base.Ui64(v99)>>(uint(v60)%64)) | v99) & v63
					v107 = int64(base.Ui64(v104)>>(uint(v65)%64)) | v104
					v115 = base.I32_wrap_i64(int64(base.Ui64(v107)>>(uint(v68)%64))&v70 | v107&v72)
					v118 = base.F64_sub(v27, v26)
					*(*float64)(unsafe.Add(mBase, uint32(l3)+32)) = base.F64_add(base.F64_mul(base.F64_div(base.F64_convert_i32_u(v115), v82), v118), v26)
					v122 = int32(1)
					*(*float64)(unsafe.Add(mBase, uint32(l3)+24)) = base.F64_add(base.F64_mul(base.F64_div(base.F64_convert_i32_u(v75+v122), v82), v84), v32)
					*(*float64)(unsafe.Add(mBase, uint32(l3)+40)) = base.F64_add(base.F64_mul(base.F64_div(base.F64_convert_i32_u(v115+v122), v82), v118), v26)
					v139 = v122
				} else {
					if base.F64_eq(v32, float64(0)) != 0 {
						v139 = v5
					} else {
						v38 = *(*int64)(unsafe.Add(mBase, uint32(l2)))
						*(*int64)(unsafe.Add(mBase, uint32(l3))) = v38
						v40 = int32(8)
						v44 = *(*int64)(unsafe.Add(mBase, uint32(l2+v40)))
						*(*int64)(unsafe.Add(mBase, uint32(l3+v40))) = v44
						v46 = int64(1)
						v52 = int64(base.Ui64(v17)>>(uint(v46)%64))&int64(4919131752989213764) | v17&int64(2459565876494606882)
						v58 = int64(1085102592571150095)
						v59 = (int64(base.Ui64(v52)>>(uint(v46)%64)) | int64(base.Ui64(v52)>>(uint(int64(3))%64))) & v58
						v60 = int64(4)
						v63 = int64(71777214294589695)
						v64 = (int64(base.Ui64(v59)>>(uint(v60)%64)) | v59) & v63
						v65 = int64(8)
						v67 = int64(base.Ui64(v64)>>(uint(v65)%64)) | v64
						v68 = int64(16)
						v70 = int64(4294901760)
						v72 = int64(65535)
						v75 = base.I32_wrap_i64(int64(base.Ui64(v67)>>(uint(v68)%64))&v70 | v67&v72)
						v82 = base.F64_convert_i64_u(v46 << (uint(base.I64_extend_i32_u(v20)&int64(255)) % 64))
						v84 = base.F64_sub(v33, v32)
						*(*float64)(unsafe.Add(mBase, uint32(l3)+16)) = base.F64_add(base.F64_mul(base.F64_div(base.F64_convert_i32_u(v75), v82), v84), v32)
						v89 = v17 & int64(6148914691236517205)
						v94 = (int64(base.Ui64(v89)>>(uint(v46)%64)) | v89) & int64(3689348814741910323)
						v99 = (int64(base.Ui64(v94)>>(uint(int64(2))%64)) | v94) & v58
						v104 = (int64(base.Ui64(v99)>>(uint(v60)%64)) | v99) & v63
						v107 = int64(base.Ui64(v104)>>(uint(v65)%64)) | v104
						v115 = base.I32_wrap_i64(int64(base.Ui64(v107)>>(uint(v68)%64))&v70 | v107&v72)
						v118 = base.F64_sub(v27, v26)
						*(*float64)(unsafe.Add(mBase, uint32(l3)+32)) = base.F64_add(base.F64_mul(base.F64_div(base.F64_convert_i32_u(v115), v82), v118), v26)
						v122 = int32(1)
						*(*float64)(unsafe.Add(mBase, uint32(l3)+24)) = base.F64_add(base.F64_mul(base.F64_div(base.F64_convert_i32_u(v75+v122), v82), v84), v32)
						*(*float64)(unsafe.Add(mBase, uint32(l3)+40)) = base.F64_add(base.F64_mul(base.F64_div(base.F64_convert_i32_u(v115+v122), v82), v118), v26)
						v139 = v122
					}
				}
			} else {
				if base.F64_eq(v26, float64(0)) != 0 {
					v139 = v5
				} else {
					v32 = *(*float64)(unsafe.Add(mBase, uint32(l0)))
					v33 = *(*float64)(unsafe.Add(mBase, uint32(l0)+8))
					if base.F64_ne(v33, float64(0)) != 0 {
						v38 = *(*int64)(unsafe.Add(mBase, uint32(l2)))
						*(*int64)(unsafe.Add(mBase, uint32(l3))) = v38
						v40 = int32(8)
						v44 = *(*int64)(unsafe.Add(mBase, uint32(l2+v40)))
						*(*int64)(unsafe.Add(mBase, uint32(l3+v40))) = v44
						v46 = int64(1)
						v52 = int64(base.Ui64(v17)>>(uint(v46)%64))&int64(4919131752989213764) | v17&int64(2459565876494606882)
						v58 = int64(1085102592571150095)
						v59 = (int64(base.Ui64(v52)>>(uint(v46)%64)) | int64(base.Ui64(v52)>>(uint(int64(3))%64))) & v58
						v60 = int64(4)
						v63 = int64(71777214294589695)
						v64 = (int64(base.Ui64(v59)>>(uint(v60)%64)) | v59) & v63
						v65 = int64(8)
						v67 = int64(base.Ui64(v64)>>(uint(v65)%64)) | v64
						v68 = int64(16)
						v70 = int64(4294901760)
						v72 = int64(65535)
						v75 = base.I32_wrap_i64(int64(base.Ui64(v67)>>(uint(v68)%64))&v70 | v67&v72)
						v82 = base.F64_convert_i64_u(v46 << (uint(base.I64_extend_i32_u(v20)&int64(255)) % 64))
						v84 = base.F64_sub(v33, v32)
						*(*float64)(unsafe.Add(mBase, uint32(l3)+16)) = base.F64_add(base.F64_mul(base.F64_div(base.F64_convert_i32_u(v75), v82), v84), v32)
						v89 = v17 & int64(6148914691236517205)
						v94 = (int64(base.Ui64(v89)>>(uint(v46)%64)) | v89) & int64(3689348814741910323)
						v99 = (int64(base.Ui64(v94)>>(uint(int64(2))%64)) | v94) & v58
						v104 = (int64(base.Ui64(v99)>>(uint(v60)%64)) | v99) & v63
						v107 = int64(base.Ui64(v104)>>(uint(v65)%64)) | v104
						v115 = base.I32_wrap_i64(int64(base.Ui64(v107)>>(uint(v68)%64))&v70 | v107&v72)
						v118 = base.F64_sub(v27, v26)
						*(*float64)(unsafe.Add(mBase, uint32(l3)+32)) = base.F64_add(base.F64_mul(base.F64_div(base.F64_convert_i32_u(v115), v82), v118), v26)
						v122 = int32(1)
						*(*float64)(unsafe.Add(mBase, uint32(l3)+24)) = base.F64_add(base.F64_mul(base.F64_div(base.F64_convert_i32_u(v75+v122), v82), v84), v32)
						*(*float64)(unsafe.Add(mBase, uint32(l3)+40)) = base.F64_add(base.F64_mul(base.F64_div(base.F64_convert_i32_u(v115+v122), v82), v118), v26)
						v139 = v122
					} else {
						if base.F64_eq(v32, float64(0)) != 0 {
							v139 = v5
						} else {
							v38 = *(*int64)(unsafe.Add(mBase, uint32(l2)))
							*(*int64)(unsafe.Add(mBase, uint32(l3))) = v38
							v40 = int32(8)
							v44 = *(*int64)(unsafe.Add(mBase, uint32(l2+v40)))
							*(*int64)(unsafe.Add(mBase, uint32(l3+v40))) = v44
							v46 = int64(1)
							v52 = int64(base.Ui64(v17)>>(uint(v46)%64))&int64(4919131752989213764) | v17&int64(2459565876494606882)
							v58 = int64(1085102592571150095)
							v59 = (int64(base.Ui64(v52)>>(uint(v46)%64)) | int64(base.Ui64(v52)>>(uint(int64(3))%64))) & v58
							v60 = int64(4)
							v63 = int64(71777214294589695)
							v64 = (int64(base.Ui64(v59)>>(uint(v60)%64)) | v59) & v63
							v65 = int64(8)
							v67 = int64(base.Ui64(v64)>>(uint(v65)%64)) | v64
							v68 = int64(16)
							v70 = int64(4294901760)
							v72 = int64(65535)
							v75 = base.I32_wrap_i64(int64(base.Ui64(v67)>>(uint(v68)%64))&v70 | v67&v72)
							v82 = base.F64_convert_i64_u(v46 << (uint(base.I64_extend_i32_u(v20)&int64(255)) % 64))
							v84 = base.F64_sub(v33, v32)
							*(*float64)(unsafe.Add(mBase, uint32(l3)+16)) = base.F64_add(base.F64_mul(base.F64_div(base.F64_convert_i32_u(v75), v82), v84), v32)
							v89 = v17 & int64(6148914691236517205)
							v94 = (int64(base.Ui64(v89)>>(uint(v46)%64)) | v89) & int64(3689348814741910323)
							v99 = (int64(base.Ui64(v94)>>(uint(int64(2))%64)) | v94) & v58
							v104 = (int64(base.Ui64(v99)>>(uint(v60)%64)) | v99) & v63
							v107 = int64(base.Ui64(v104)>>(uint(v65)%64)) | v104
							v115 = base.I32_wrap_i64(int64(base.Ui64(v107)>>(uint(v68)%64))&v70 | v107&v72)
							v118 = base.F64_sub(v27, v26)
							*(*float64)(unsafe.Add(mBase, uint32(l3)+32)) = base.F64_add(base.F64_mul(base.F64_div(base.F64_convert_i32_u(v115), v82), v118), v26)
							v122 = int32(1)
							*(*float64)(unsafe.Add(mBase, uint32(l3)+24)) = base.F64_add(base.F64_mul(base.F64_div(base.F64_convert_i32_u(v75+v122), v82), v84), v32)
							*(*float64)(unsafe.Add(mBase, uint32(l3)+40)) = base.F64_add(base.F64_mul(base.F64_div(base.F64_convert_i32_u(v115+v122), v82), v118), v26)
							v139 = v122
						}
					}
				}
			}
		}
	}
	return v139
}
func F_geohashGetDistanceIfInPolygon(m *base.Module, l0 float64, l1 float64, l2 int32, l3 int32, l4 int32, l5 int32) int32 {
	mBase := m.M
	_ = mBase
	var v20 int32
	_ = v20
	var v24 float64
	_ = v24
	var v25 float64
	_ = v25
	var v26 int32
	_ = v26
	var v34 int32
	_ = v34
	var __phi34 int32
	_ = __phi34
	var v35 float64
	_ = v35
	var __phi35 float64
	_ = __phi35
	var v37 int32
	_ = v37
	var __phi37 int32
	_ = __phi37
	var v38 int32
	_ = v38
	var __phi38 int32
	_ = __phi38
	var v45 int32
	_ = v45
	var v46 float64
	_ = v46
	var v49 float64
	_ = v49
	var v50 float64
	_ = v50
	var v55 float64
	_ = v55
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v73 float64
	_ = v73
	var v74 float64
	_ = v74
	var v80 float64
	_ = v80
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v93 int32
	_ = v93
	var v100 float64
	_ = v100
	var v104 int32
	_ = v104
	var v105 float64
	_ = v105
	var v106 float64
	_ = v106
	var v110 float64
	_ = v110
	var v111 float64
	_ = v111
	var v113 float64
	_ = v113
	var v115 float64
	_ = v115
	var v117 float64
	_ = v117
	var v123 float64
	_ = v123
	var v129 float64
	_ = v129
	var v140 float64
	_ = v140
	var v144 int32
	_ = v144
	var v146 int32
	_ = v146
	var v153 int32
	_ = v153
	var v160 float64
	_ = v160
	var v164 int32
	_ = v164
	var v165 float64
	_ = v165
	var v166 float64
	_ = v166
	var v169 float64
	_ = v169
	var v171 float64
	_ = v171
	var v173 float64
	_ = v173
	var v176 float64
	_ = v176
	var v179 float64
	_ = v179
	var v184 float64
	_ = v184
	var v188 int32
	_ = v188
	var v190 int32
	_ = v190
	var v197 int32
	_ = v197
	var v204 float64
	_ = v204
	var v208 int32
	_ = v208
	var v209 float64
	_ = v209
	var v210 float64
	_ = v210
	var v213 float64
	_ = v213
	var v215 float64
	_ = v215
	var v217 float64
	_ = v217
	var v220 float64
	_ = v220
	var v223 float64
	_ = v223
	var v229 float64
	_ = v229
	var v233 int32
	_ = v233
	var v235 int32
	_ = v235
	var v242 int32
	_ = v242
	var v249 float64
	_ = v249
	var v253 int32
	_ = v253
	var v254 float64
	_ = v254
	var v255 float64
	_ = v255
	var v259 float64
	_ = v259
	var v260 float64
	_ = v260
	var v262 float64
	_ = v262
	var v264 float64
	_ = v264
	var v266 float64
	_ = v266
	var v277 float64
	_ = v277
	var v283 int64
	_ = v283
	var v288 int32
	_ = v288
	var v309 float64
	_ = v309
	var v313 float64
	_ = v313
	var v316 float64
	_ = v316
	var v317 float64
	_ = v317
	var v318 float64
	_ = v318
	var v323 float64
	_ = v323
	var v328 float64
	_ = v328
	var v332 float64
	_ = v332
	var v341 float64
	_ = v341
	var v348 float64
	_ = v348
	var v353 float64
	_ = v353
	var v354 float64
	_ = v354
	var v362 float64
	_ = v362
	if int32(1) <= l4 {
		v20 = l4 + int32(-1)
		v24 = *(*float64)(unsafe.Add(mBase, uint32(l3+v20<<(uint(int32(4))%32))+8))
		v25 = *(*float64)(unsafe.Add(mBase, uint32(l2)+8))
		v26 = int32(0)
		__phi34 = v20
		__phi35 = v24
		__phi37 = v26
		__phi38 = v26
		v34 = __phi34
		v35 = __phi35
		v37 = __phi37
		v38 = __phi38
		for {
			v45 = l3 + v37<<(uint(int32(4))%32)
			v46 = *(*float64)(unsafe.Add(mBase, uint32(v45)+8))
			if base.F64_gt(v35, v25) == base.F64_gt(v46, v25) {
				v66 = v38
			} else {
				v49 = *(*float64)(unsafe.Add(mBase, uint32(l2)))
				v50 = *(*float64)(unsafe.Add(mBase, uint32(v45)))
				v55 = *(*float64)(unsafe.Add(mBase, uint32(l3+v34<<(uint(int32(4))%32))))
				if base.F64_lt(v49, base.F64_add(v50, base.F64_div(base.F64_mul(base.F64_sub(v25, v46), base.F64_sub(v55, v50)), base.F64_sub(v35, v46)))) == int32(0) {
					v66 = v38
				} else {
					v66 = base.B2i32(v38 == int32(0))
				}
			}
			v69 = v37 + int32(1)
			if v69 != l4 {
				__phi34 = v37
				__phi35 = v46
				__phi37 = v69
				__phi38 = v66
				v34 = __phi34
				v35 = __phi35
				v37 = __phi37
				v38 = __phi38
				continue
			} else {
				break
			}
			break
		}
		if v66 != 0 {
			v73 = *(*float64)(unsafe.Add(mBase, uint32(l2)))
			v74 = float64(0.017453292519943295)
			v80 = base.F64_mul(base.F64_sub(base.F64_mul(v73, v74), base.F64_mul(l0, v74)), float64(0.5))
			v84 = m.G0
			v86 = v84 - int32(16)
			m.G0 = v86
			v93 = base.I32_wrap_i64(int64(base.Ui64(base.I64_reinterpret_f64(v80))>>(uint(int64(32))%64))) & int32(2147483647)
			if base.Ui32(int32(1072243195)) < base.Ui32(v93) {
				if base.Ui32(v93) < base.Ui32(int32(2146435072)) {
					v104 = F___rem_pio2(m, v80, v86)
					mBase = m.M
					v105 = *(*float64)(unsafe.Add(mBase, uint32(v86)+8))
					v106 = *(*float64)(unsafe.Add(mBase, uint32(v86)))
					switch v104 & int32(3) {
					default:
						v110 = F___sin(m, v106, v105, int32(1))
						mBase = m.M
						v117 = v110
					case 1:
						v111 = F___cos(m, v106, v105)
						mBase = m.M
						v117 = v111
					case 2:
						v113 = F___sin(m, v106, v105, int32(1))
						mBase = m.M
						v117 = base.F64_neg(v113)
					case 3:
						v115 = F___cos(m, v106, v105)
						mBase = m.M
						v117 = base.F64_neg(v115)
					}
				} else {
					v117 = base.F64_sub(v80, v80)
				}
			} else {
				if base.Ui32(v93) < base.Ui32(int32(1045430272)) {
					v117 = v80
				} else {
					v100 = F___sin(m, v80, float64(0), int32(0))
					mBase = m.M
					v117 = v100
				}
			}
			m.G0 = v86 + int32(16)
			v123 = *(*float64)(unsafe.Add(mBase, uint32(l2)+8))
			if base.F64_le(base.F64_abs(v117), float64(1e-15)) == int32(0) {
				v140 = base.F64_mul(l1, float64(0.017453292519943295))
				v144 = m.G0
				v146 = v144 - int32(16)
				m.G0 = v146
				v153 = base.I32_wrap_i64(int64(base.Ui64(base.I64_reinterpret_f64(v140))>>(uint(int64(32))%64))) & int32(2147483647)
				if base.Ui32(int32(1072243195)) < base.Ui32(v153) {
					if base.Ui32(v153) < base.Ui32(int32(2146435072)) {
						v164 = F___rem_pio2(m, v140, v146)
						mBase = m.M
						v165 = *(*float64)(unsafe.Add(mBase, uint32(v146)+8))
						v166 = *(*float64)(unsafe.Add(mBase, uint32(v146)))
						switch v164 & int32(3) {
						default:
							v169 = F___cos(m, v166, v165)
							mBase = m.M
							v179 = v169
						case 1:
							v171 = F___sin(m, v166, v165, int32(1))
							mBase = m.M
							v179 = base.F64_neg(v171)
						case 2:
							v173 = F___cos(m, v166, v165)
							mBase = m.M
							v179 = base.F64_neg(v173)
						case 3:
							v176 = F___sin(m, v166, v165, int32(1))
							mBase = m.M
							v179 = v176
						}
					} else {
						v179 = base.F64_sub(v140, v140)
					}
				} else {
					if base.Ui32(v153) < base.Ui32(int32(1044816030)) {
						v179 = float64(1)
					} else {
						v160 = F___cos(m, v140, float64(0))
						mBase = m.M
						v179 = v160
					}
				}
				m.G0 = v146 + int32(16)
				v184 = base.F64_mul(v123, float64(0.017453292519943295))
				v188 = m.G0
				v190 = v188 - int32(16)
				m.G0 = v190
				v197 = base.I32_wrap_i64(int64(base.Ui64(base.I64_reinterpret_f64(v184))>>(uint(int64(32))%64))) & int32(2147483647)
				if base.Ui32(int32(1072243195)) < base.Ui32(v197) {
					if base.Ui32(v197) < base.Ui32(int32(2146435072)) {
						v208 = F___rem_pio2(m, v184, v190)
						mBase = m.M
						v209 = *(*float64)(unsafe.Add(mBase, uint32(v190)+8))
						v210 = *(*float64)(unsafe.Add(mBase, uint32(v190)))
						switch v208 & int32(3) {
						default:
							v213 = F___cos(m, v210, v209)
							mBase = m.M
							v223 = v213
						case 1:
							v215 = F___sin(m, v210, v209, int32(1))
							mBase = m.M
							v223 = base.F64_neg(v215)
						case 2:
							v217 = F___cos(m, v210, v209)
							mBase = m.M
							v223 = base.F64_neg(v217)
						case 3:
							v220 = F___sin(m, v210, v209, int32(1))
							mBase = m.M
							v223 = v220
						}
					} else {
						v223 = base.F64_sub(v184, v184)
					}
				} else {
					if base.Ui32(v197) < base.Ui32(int32(1044816030)) {
						v223 = float64(1)
					} else {
						v204 = F___cos(m, v184, float64(0))
						mBase = m.M
						v223 = v204
					}
				}
				m.G0 = v190 + int32(16)
				v229 = base.F64_mul(base.F64_sub(v184, v140), float64(0.5))
				v233 = m.G0
				v235 = v233 - int32(16)
				m.G0 = v235
				v242 = base.I32_wrap_i64(int64(base.Ui64(base.I64_reinterpret_f64(v229))>>(uint(int64(32))%64))) & int32(2147483647)
				if base.Ui32(int32(1072243195)) < base.Ui32(v242) {
					if base.Ui32(v242) < base.Ui32(int32(2146435072)) {
						v253 = F___rem_pio2(m, v229, v235)
						mBase = m.M
						v254 = *(*float64)(unsafe.Add(mBase, uint32(v235)+8))
						v255 = *(*float64)(unsafe.Add(mBase, uint32(v235)))
						switch v253 & int32(3) {
						default:
							v259 = F___sin(m, v255, v254, int32(1))
							mBase = m.M
							v266 = v259
						case 1:
							v260 = F___cos(m, v255, v254)
							mBase = m.M
							v266 = v260
						case 2:
							v262 = F___sin(m, v255, v254, int32(1))
							mBase = m.M
							v266 = base.F64_neg(v262)
						case 3:
							v264 = F___cos(m, v255, v254)
							mBase = m.M
							v266 = base.F64_neg(v264)
						}
					} else {
						v266 = base.F64_sub(v229, v229)
					}
				} else {
					if base.Ui32(v242) < base.Ui32(int32(1045430272)) {
						v266 = v229
					} else {
						v249 = F___sin(m, v229, float64(0), int32(0))
						mBase = m.M
						v266 = v249
					}
				}
				m.G0 = v235 + int32(16)
				v277 = base.F64_sqrt(base.F64_add(base.F64_mul(v266, v266), base.F64_mul(v117, base.F64_mul(base.F64_mul(v179, v223), v117))))
				v283 = base.I64_reinterpret_f64(v277)
				v288 = base.I32_wrap_i64(int64(base.Ui64(v283)>>(uint(int64(32))%64))) & int32(2147483647)
				if base.Ui32(v288) < base.Ui32(int32(1072693248)) {
					if base.Ui32(int32(1071644671)) < base.Ui32(v288) {
						v313 = F_fabs(m, v277)
						mBase = m.M
						v316 = base.F64_mul(base.F64_sub(float64(1), v313), float64(0.5))
						v317 = F_sqrt(m, v316)
						mBase = m.M
						v318 = F_R_2(m, v316)
						mBase = m.M
						if base.Ui32(v288) < base.Ui32(int32(1072640819)) {
							v328 = float64(0.7853981633974483)
							v332 = base.F64_reinterpret_i64(base.I64_reinterpret_f64(v317) & int64(-4294967296))
							v341 = base.F64_div(base.F64_sub(v316, base.F64_mul(v332, v332)), base.F64_add(v317, v332))
							v348 = base.F64_add(base.F64_sub(base.F64_sub(v328, base.F64_add(v332, v332)), base.F64_sub(base.F64_mul(base.F64_add(v317, v317), v318), base.F64_sub(float64(6.123233995736766e-17), base.F64_add(v341, v341)))), v328)
						} else {
							v323 = base.F64_add(base.F64_mul(v317, v318), v317)
							v348 = base.F64_sub(float64(1.5707963267948966), base.F64_add(base.F64_add(v323, v323), float64(-6.123233995736766e-17)))
						}
						if v283 < int64(0) {
							v353 = base.F64_neg(v348)
						} else {
							v353 = v348
						}
						v354 = v353
						v362 = v354
					} else {
						if base.Ui32(v288+int32(-1048576)) < base.Ui32(int32(1044381696)) {
							v354 = v277
							v362 = v354
						} else {
							v309 = F_R_2(m, base.F64_mul(v277, v277))
							mBase = m.M
							v362 = base.F64_add(base.F64_mul(v277, v309), v277)
						}
					}
				} else {
					if v288+int32(-1072693248)|base.I32_wrap_i64(v283) != 0 {
						v362 = base.F64_div(float64(0), base.F64_sub(v277, v277))
					} else {
						v362 = base.F64_add(base.F64_mul(v277, float64(1.5707963267948966)), float64(7.52316384526264e-37))
					}
				}
				*(*float64)(unsafe.Add(mBase, uint32(l5))) = base.F64_mul(v362, float64(1.2745595121712e+07))
				return v66
			} else {
				v129 = float64(0.017453292519943295)
				*(*float64)(unsafe.Add(mBase, uint32(l5))) = base.F64_mul(base.F64_abs(base.F64_sub(base.F64_mul(v123, v129), base.F64_mul(l1, v129))), float64(6.372797560856e+06))
				return v66
			}
		} else {
			return int32(0)
		}
	} else {
		return int32(0)
	}
}
