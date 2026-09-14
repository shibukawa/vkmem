package p0

import (
	base "github.com/shibukawa/vkmem/internal/aot/vkaot/base"
	"unsafe"
)

func F___multf3(m *base.Module, l0 int32, l1 int64, l2 int64, l3 int64, l4 int64) {
	mBase := m.M
	_ = mBase
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v30 int64
	_ = v30
	var v31 int64
	_ = v31
	var v34 int64
	_ = v34
	var v36 int64
	_ = v36
	var v38 int64
	_ = v38
	var v39 int64
	_ = v39
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v48 int32
	_ = v48
	var v54 int32
	_ = v54
	var v62 int64
	_ = v62
	var v63 int64
	_ = v63
	var v67 int32
	_ = v67
	var v73 int64
	_ = v73
	var v74 int64
	_ = v74
	var v78 int32
	_ = v78
	var v102 int64
	_ = v102
	var v122 int32
	_ = v122
	var v124 int32
	_ = v124
	var v125 int64
	_ = v125
	var v131 int32
	_ = v131
	var v133 int32
	_ = v133
	var v150 int64
	_ = v150
	var v154 int64
	_ = v154
	var v155 int64
	_ = v155
	var v163 int64
	_ = v163
	var v166 int64
	_ = v166
	var v167 int64
	_ = v167
	var v168 int64
	_ = v168
	var v169 int64
	_ = v169
	var v170 int32
	_ = v170
	var v174 int32
	_ = v174
	var v176 int32
	_ = v176
	var v177 int64
	_ = v177
	var v183 int32
	_ = v183
	var v185 int32
	_ = v185
	var v202 int64
	_ = v202
	var v206 int64
	_ = v206
	var v207 int64
	_ = v207
	var v216 int64
	_ = v216
	var v217 int64
	_ = v217
	var v218 int64
	_ = v218
	var v220 int64
	_ = v220
	var v221 int64
	_ = v221
	var v222 int64
	_ = v222
	var v223 int64
	_ = v223
	var v224 int32
	_ = v224
	var v227 int64
	_ = v227
	var v228 int64
	_ = v228
	var v230 int64
	_ = v230
	var v231 int64
	_ = v231
	var v232 int64
	_ = v232
	var v233 int64
	_ = v233
	var v235 int64
	_ = v235
	var v236 int64
	_ = v236
	var v237 int64
	_ = v237
	var v239 int64
	_ = v239
	var v241 int64
	_ = v241
	var v243 int64
	_ = v243
	var v247 int64
	_ = v247
	var v248 int64
	_ = v248
	var v250 int64
	_ = v250
	var v254 int64
	_ = v254
	var v257 int64
	_ = v257
	var v259 int64
	_ = v259
	var v267 int64
	_ = v267
	var v269 int64
	_ = v269
	var v270 int64
	_ = v270
	var v272 int64
	_ = v272
	var v276 int64
	_ = v276
	var v278 int64
	_ = v278
	var v280 int64
	_ = v280
	var v283 int64
	_ = v283
	var v284 int64
	_ = v284
	var v286 int32
	_ = v286
	var v289 int64
	_ = v289
	var v291 int64
	_ = v291
	var v295 int64
	_ = v295
	var v306 int64
	_ = v306
	var v310 int64
	_ = v310
	var v312 int64
	_ = v312
	var v323 int64
	_ = v323
	var v340 int64
	_ = v340
	var v349 int64
	_ = v349
	var v352 int64
	_ = v352
	var v359 int64
	_ = v359
	var v361 int64
	_ = v361
	var v376 int64
	_ = v376
	var v377 int64
	_ = v377
	var v379 int64
	_ = v379
	var v380 int32
	_ = v380
	var v381 int64
	_ = v381
	var v390 int32
	_ = v390
	var v394 int32
	_ = v394
	var v396 int32
	_ = v396
	var v413 int64
	_ = v413
	var v417 int64
	_ = v417
	var v418 int64
	_ = v418
	var v423 int32
	_ = v423
	var v440 int64
	_ = v440
	var v444 int64
	_ = v444
	var v445 int64
	_ = v445
	var v450 int32
	_ = v450
	var v467 int64
	_ = v467
	var v471 int64
	_ = v471
	var v472 int64
	_ = v472
	var v492 int64
	_ = v492
	var v496 int64
	_ = v496
	var v497 int64
	_ = v497
	var v501 int64
	_ = v501
	var v502 int64
	_ = v502
	var v504 int64
	_ = v504
	var v509 int64
	_ = v509
	var v519 int64
	_ = v519
	var v524 int64
	_ = v524
	var v528 int64
	_ = v528
	var v529 int64
	_ = v529
	var v537 int64
	_ = v537
	var v538 int64
	_ = v538
	var v539 int64
	_ = v539
	var v542 int64
	_ = v542
	var v543 int64
	_ = v543
	var v550 int32
	_ = v550
	var v552 int64
	_ = v552
	var v564 int64
	_ = v564
	var v568 int64
	_ = v568
	var v573 int64
	_ = v573
	v26 = m.G0
	v28 = v26 - int32(96)
	m.G0 = v28
	v30 = int64(281474976710655)
	v31 = l4 & v30
	v34 = (l4 ^ l2) & int64(-9223372036854775807-1)
	v36 = l2 & v30
	v38 = int64(base.Ui64(v36) >> (uint(int64(32)) % 64))
	v39 = int64(48)
	v42 = int32(32767)
	v43 = base.I32_wrap_i64(int64(base.Ui64(l4)>>(uint(v39)%64))) & v42
	v48 = base.I32_wrap_i64(int64(base.Ui64(l2)>>(uint(v39)%64))) & v42
	if base.Ui32(v48+int32(-32767)) < base.Ui32(int32(-32766)) {
		v62 = l2 & int64(9223372036854775807)
		v63 = int64(9223090561878065152)
		if v62 == v63 {
			v67 = base.B2i32(l1 == int64(0))
		} else {
			v67 = base.B2i32(base.Ui64(v62) < base.Ui64(v63))
		}
		if v67 != 0 {
			v73 = l4 & int64(9223372036854775807)
			v74 = int64(9223090561878065152)
			if v73 == v74 {
				v78 = base.B2i32(l3 == int64(0))
			} else {
				v78 = base.B2i32(base.Ui64(v73) < base.Ui64(v74))
			}
			if v78 != 0 {
				if l1|(v62^int64(9223090561878065152)) != int64(0) {
					if l3|(v73^int64(9223090561878065152)) != int64(0) {
						if l1|v62 != int64(0) {
							if l3|v73 != int64(0) {
								if base.Ui64(int64(281474976710655)) < base.Ui64(v62) {
									v167 = l1
									v168 = v36
									v169 = v38
									v170 = int32(0)
								} else {
									v122 = v28 + int32(80)
									v124 = base.B2i32(v36 == int64(0))
									if v36 == int64(0) {
										v125 = l1
									} else {
										v125 = v36
									}
									v131 = base.I32_wrap_i64(base.I64_clz(v125) + base.I64_extend_i32_u(v124<<(uint(int32(6))%32)))
									v133 = v131 + int32(-15)
									if v133&int32(64) == int32(0) {
										if v133 == int32(0) {
											v154 = l1
											v155 = v36
										} else {
											v150 = base.I64_extend_i32_u(v133)
											v154 = l1 << (uint(v150) % 64)
											v155 = int64(base.Ui64(l1)>>(uint(base.I64_extend_i32_u(int32(64)-v133))%64)) | v36<<(uint(v150)%64)
										}
									} else {
										v154 = int64(0)
										v155 = l1 << (uint(base.I64_extend_i32_u(v131+int32(-79))) % 64)
									}
									*(*int64)(unsafe.Add(mBase, uint32(v122))) = v154
									*(*int64)(unsafe.Add(mBase, uint32(v122)+8)) = v155
									v163 = *(*int64)(unsafe.Add(mBase, uint32(v28+int32(88))))
									v166 = *(*int64)(unsafe.Add(mBase, uint32(v28)+80))
									v167 = v166
									v168 = v163
									v169 = int64(base.Ui64(v163) >> (uint(int64(32)) % 64))
									v170 = int32(16) - v131
								}
								if base.Ui64(int64(281474976710655)) < base.Ui64(v73) {
									v218 = v167
									v220 = l3
									v221 = v31
									v222 = v168
									v223 = v169
									v224 = v170
								} else {
									v174 = v28 + int32(64)
									v176 = base.B2i32(v31 == int64(0))
									if v31 == int64(0) {
										v177 = l3
									} else {
										v177 = v31
									}
									v183 = base.I32_wrap_i64(base.I64_clz(v177) + base.I64_extend_i32_u(v176<<(uint(int32(6))%32)))
									v185 = v183 + int32(-15)
									if v185&int32(64) == int32(0) {
										if v185 == int32(0) {
											v206 = l3
											v207 = v31
										} else {
											v202 = base.I64_extend_i32_u(v185)
											v206 = l3 << (uint(v202) % 64)
											v207 = int64(base.Ui64(l3)>>(uint(base.I64_extend_i32_u(int32(64)-v185))%64)) | v31<<(uint(v202)%64)
										}
									} else {
										v206 = int64(0)
										v207 = l3 << (uint(base.I64_extend_i32_u(v183+int32(-79))) % 64)
									}
									*(*int64)(unsafe.Add(mBase, uint32(v174))) = v206
									*(*int64)(unsafe.Add(mBase, uint32(v174)+8)) = v207
									v216 = *(*int64)(unsafe.Add(mBase, uint32(v28+int32(72))))
									v217 = *(*int64)(unsafe.Add(mBase, uint32(v28)+64))
									v218 = v167
									v220 = v217
									v221 = v216
									v222 = v168
									v223 = v169
									v224 = v170 - v183 + int32(16)
								}
								v227 = int64(15)
								v228 = v220 << (uint(v227) % 64)
								v230 = v228 & int64(4294934528)
								v231 = int64(32)
								v232 = int64(base.Ui64(v218) >> (uint(v231) % 64))
								v233 = v230 * v232
								v235 = int64(base.Ui64(v228) >> (uint(v231) % 64))
								v236 = int64(4294967295)
								v237 = v218 & v236
								v239 = v233 + v235*v237
								v241 = v239 << (uint(v231) % 64)
								v243 = v241 + v230*v237
								v247 = v222 & v236
								v248 = v230 * v247
								v250 = v248 + v235*v232
								v254 = v221 << (uint(v227) % 64)
								v257 = (int64(base.Ui64(v220)>>(uint(int64(49))%64)) | v254) & v236
								v259 = v250 + v257*v237
								v267 = v259 + (int64(base.Ui64(v239)>>(uint(v231)%64)) | base.I64_extend_i32_u(base.B2i32(base.Ui64(v239) < base.Ui64(v233)))<<(uint(v231)%64))
								v269 = v223 | int64(65536)
								v270 = v230 * v269
								v272 = v270 + v235*v247
								v276 = int64(base.Ui64(v254)>>(uint(v231)%64)) | int64(2147483648)
								v278 = v272 + v276*v237
								v280 = v278 + v257*v232
								v283 = v267 + v280<<(uint(v231)%64)
								v284 = base.I64_extend_i32_u(base.B2i32(base.Ui64(v243) < base.Ui64(v241))) + v283
								v286 = v48 + v43 + v224
								v289 = v276 * v232
								v291 = v289 + v235*v269
								v295 = v291 + v257*v247
								v306 = v295 + (base.I64_extend_i32_u(base.B2i32(base.Ui64(v250) < base.Ui64(v248))) + base.I64_extend_i32_u(base.B2i32(base.Ui64(v259) < base.Ui64(v250))))
								v310 = v257 * v269
								v312 = v310 + v276*v247
								v323 = v306 + v312<<(uint(v231)%64)
								v340 = v323 + (int64(base.Ui64(v280)>>(uint(v231)%64)) | (base.I64_extend_i32_u(base.B2i32(base.Ui64(v272) < base.Ui64(v270)))+base.I64_extend_i32_u(base.B2i32(base.Ui64(v278) < base.Ui64(v272)))+base.I64_extend_i32_u(base.B2i32(base.Ui64(v280) < base.Ui64(v278))))<<(uint(v231)%64))
								v349 = v340 + (base.I64_extend_i32_u(base.B2i32(base.Ui64(v267) < base.Ui64(v259))) + base.I64_extend_i32_u(base.B2i32(base.Ui64(v283) < base.Ui64(v267))))
								v352 = base.I64_extend_i32_u(base.B2i32(base.Ui64(v291) < base.Ui64(v289))) + base.I64_extend_i32_u(base.B2i32(base.Ui64(v295) < base.Ui64(v291))) + v276*v269 + base.I64_extend_i32_u(base.B2i32(base.Ui64(v306) < base.Ui64(v295))) + (base.I64_extend_i32_u(base.B2i32(base.Ui64(v312) < base.Ui64(v310)))<<(uint(v231)%64) | int64(base.Ui64(v312)>>(uint(v231)%64))) + base.I64_extend_i32_u(base.B2i32(base.Ui64(v323) < base.Ui64(v306))) + base.I64_extend_i32_u(base.B2i32(base.Ui64(v340) < base.Ui64(v323))) + base.I64_extend_i32_u(base.B2i32(base.Ui64(v349) < base.Ui64(v340)))
								if v352&int64(281474976710656) == int64(0) {
									v359 = int64(63)
									v361 = int64(1)
									v376 = int64(base.Ui64(v243)>>(uint(v359)%64)) | v284<<(uint(v361)%64)
									v377 = v349<<(uint(v361)%64) | int64(base.Ui64(v284)>>(uint(v359)%64))
									v379 = v352<<(uint(v361)%64) | int64(base.Ui64(v349)>>(uint(v359)%64))
									v380 = v286 + int32(-16383)
									v381 = v243 << (uint(v361) % 64)
								} else {
									v376 = v284
									v377 = v349
									v379 = v352
									v380 = v286 + int32(-16382)
									v381 = v243
								}
								if v380 < int32(32767) {
									if int32(0) < v380 {
										v537 = v376
										v538 = v377
										v539 = base.I64_extend_i32_u(v380)<<(uint(int64(48))%64) | v379&int64(281474976710655)
										v542 = v381
										v543 = v539 | v34
										if v537 == int64(-9223372036854775807-1) {
											v550 = base.B2i32(v542 == int64(0))
										} else {
											v550 = base.B2i32(int64(-1) < v537)
										}
										if v550 != 0 {
											if v542|(v537^int64(-9223372036854775807-1)) == int64(0) {
												v564 = v538 + v538&int64(1)
												v568 = v564
												v573 = v543 + base.I64_extend_i32_u(base.B2i32(base.Ui64(v564) < base.Ui64(v538)))
											} else {
												v568 = v538
												v573 = v543
											}
										} else {
											v552 = v538 + int64(1)
											v568 = v552
											v573 = v543 + base.I64_extend_i32_u(base.B2i32(v552 == int64(0)))
										}
									} else {
										v390 = int32(1) - v380
										if base.Ui32(int32(127)) < base.Ui32(v390) {
											v568 = int64(0)
											v573 = v34
										} else {
											v394 = v28 + int32(48)
											v396 = v380 + int32(127)
											if v396&int32(64) == int32(0) {
												if v396 == int32(0) {
													v417 = v381
													v418 = v376
												} else {
													v413 = base.I64_extend_i32_u(v396)
													v417 = v381 << (uint(v413) % 64)
													v418 = int64(base.Ui64(v381)>>(uint(base.I64_extend_i32_u(int32(64)-v396))%64)) | v376<<(uint(v413)%64)
												}
											} else {
												v417 = int64(0)
												v418 = v381 << (uint(base.I64_extend_i32_u(v380+int32(63))) % 64)
											}
											*(*int64)(unsafe.Add(mBase, uint32(v394))) = v417
											*(*int64)(unsafe.Add(mBase, uint32(v394)+8)) = v418
											v423 = v28 + int32(32)
											if v396&int32(64) == int32(0) {
												if v396 == int32(0) {
													v444 = v377
													v445 = v379
												} else {
													v440 = base.I64_extend_i32_u(v396)
													v444 = v377 << (uint(v440) % 64)
													v445 = int64(base.Ui64(v377)>>(uint(base.I64_extend_i32_u(int32(64)-v396))%64)) | v379<<(uint(v440)%64)
												}
											} else {
												v444 = int64(0)
												v445 = v377 << (uint(base.I64_extend_i32_u(v380+int32(63))) % 64)
											}
											*(*int64)(unsafe.Add(mBase, uint32(v423))) = v444
											*(*int64)(unsafe.Add(mBase, uint32(v423)+8)) = v445
											v450 = v28 + int32(16)
											if v390&int32(64) == int32(0) {
												if v390 == int32(0) {
													v471 = v381
													v472 = v376
												} else {
													v467 = base.I64_extend_i32_u(v390)
													v471 = v376<<(uint(base.I64_extend_i32_u(int32(64)-v390))%64) | int64(base.Ui64(v381)>>(uint(v467)%64))
													v472 = int64(base.Ui64(v376) >> (uint(v467) % 64))
												}
											} else {
												v471 = int64(base.Ui64(v376) >> (uint(base.I64_extend_i32_u(v390+int32(-64))) % 64))
												v472 = int64(0)
											}
											*(*int64)(unsafe.Add(mBase, uint32(v450))) = v471
											*(*int64)(unsafe.Add(mBase, uint32(v450)+8)) = v472
											if v390&int32(64) == int32(0) {
												if v390 == int32(0) {
													v496 = v377
													v497 = v379
												} else {
													v492 = base.I64_extend_i32_u(v390)
													v496 = v379<<(uint(base.I64_extend_i32_u(int32(64)-v390))%64) | int64(base.Ui64(v377)>>(uint(v492)%64))
													v497 = int64(base.Ui64(v379) >> (uint(v492) % 64))
												}
											} else {
												v496 = int64(base.Ui64(v379) >> (uint(base.I64_extend_i32_u(v390+int32(-64))) % 64))
												v497 = int64(0)
											}
											*(*int64)(unsafe.Add(mBase, uint32(v28))) = v496
											*(*int64)(unsafe.Add(mBase, uint32(v28)+8)) = v497
											v501 = *(*int64)(unsafe.Add(mBase, uint32(v28)+32))
											v502 = *(*int64)(unsafe.Add(mBase, uint32(v28)+16))
											v504 = *(*int64)(unsafe.Add(mBase, uint32(v28)+48))
											v509 = *(*int64)(unsafe.Add(mBase, uint32(v28+int32(56))))
											v519 = *(*int64)(unsafe.Add(mBase, uint32(v28+int32(40))))
											v524 = *(*int64)(unsafe.Add(mBase, uint32(v28+int32(24))))
											v528 = *(*int64)(unsafe.Add(mBase, uint32(v28+int32(8))))
											v529 = *(*int64)(unsafe.Add(mBase, uint32(v28)))
											v537 = v519 | v524
											v538 = v529
											v539 = v528
											v542 = v501 | v502 | base.I64_extend_i32_u(base.B2i32(v504|v509 != int64(0)))
											v543 = v539 | v34
											if v537 == int64(-9223372036854775807-1) {
												v550 = base.B2i32(v542 == int64(0))
											} else {
												v550 = base.B2i32(int64(-1) < v537)
											}
											if v550 != 0 {
												if v542|(v537^int64(-9223372036854775807-1)) == int64(0) {
													v564 = v538 + v538&int64(1)
													v568 = v564
													v573 = v543 + base.I64_extend_i32_u(base.B2i32(base.Ui64(v564) < base.Ui64(v538)))
												} else {
													v568 = v538
													v573 = v543
												}
											} else {
												v552 = v538 + int64(1)
												v568 = v552
												v573 = v543 + base.I64_extend_i32_u(base.B2i32(v552 == int64(0)))
											}
										}
									}
								} else {
									v568 = int64(0)
									v573 = v34 | int64(9223090561878065152)
								}
							} else {
								v568 = int64(0)
								v573 = v34
							}
						} else {
							v568 = int64(0)
							v573 = v34
						}
					} else {
						v102 = int64(0)
						if base.B2i32(l1|v62 == v102) == int32(0) {
							v568 = v102
							v573 = v34 | int64(9223090561878065152)
						} else {
							v568 = v102
							v573 = int64(9223231299366420480)
						}
					}
				} else {
					if base.B2i32(l3|v73 == int64(0)) == int32(0) {
						v568 = int64(0)
						v573 = v34 | int64(9223090561878065152)
					} else {
						v568 = int64(0)
						v573 = int64(9223231299366420480)
					}
				}
			} else {
				v568 = l3
				v573 = l4 | int64(140737488355328)
			}
		} else {
			v568 = l1
			v573 = l2 | int64(140737488355328)
		}
	} else {
		v54 = int32(-32767)
		if base.Ui32(v54) < base.Ui32(v43+v54) {
			v218 = l1
			v220 = l3
			v221 = v31
			v222 = v36
			v223 = v38
			v224 = int32(0)
			v227 = int64(15)
			v228 = v220 << (uint(v227) % 64)
			v230 = v228 & int64(4294934528)
			v231 = int64(32)
			v232 = int64(base.Ui64(v218) >> (uint(v231) % 64))
			v233 = v230 * v232
			v235 = int64(base.Ui64(v228) >> (uint(v231) % 64))
			v236 = int64(4294967295)
			v237 = v218 & v236
			v239 = v233 + v235*v237
			v241 = v239 << (uint(v231) % 64)
			v243 = v241 + v230*v237
			v247 = v222 & v236
			v248 = v230 * v247
			v250 = v248 + v235*v232
			v254 = v221 << (uint(v227) % 64)
			v257 = (int64(base.Ui64(v220)>>(uint(int64(49))%64)) | v254) & v236
			v259 = v250 + v257*v237
			v267 = v259 + (int64(base.Ui64(v239)>>(uint(v231)%64)) | base.I64_extend_i32_u(base.B2i32(base.Ui64(v239) < base.Ui64(v233)))<<(uint(v231)%64))
			v269 = v223 | int64(65536)
			v270 = v230 * v269
			v272 = v270 + v235*v247
			v276 = int64(base.Ui64(v254)>>(uint(v231)%64)) | int64(2147483648)
			v278 = v272 + v276*v237
			v280 = v278 + v257*v232
			v283 = v267 + v280<<(uint(v231)%64)
			v284 = base.I64_extend_i32_u(base.B2i32(base.Ui64(v243) < base.Ui64(v241))) + v283
			v286 = v48 + v43 + v224
			v289 = v276 * v232
			v291 = v289 + v235*v269
			v295 = v291 + v257*v247
			v306 = v295 + (base.I64_extend_i32_u(base.B2i32(base.Ui64(v250) < base.Ui64(v248))) + base.I64_extend_i32_u(base.B2i32(base.Ui64(v259) < base.Ui64(v250))))
			v310 = v257 * v269
			v312 = v310 + v276*v247
			v323 = v306 + v312<<(uint(v231)%64)
			v340 = v323 + (int64(base.Ui64(v280)>>(uint(v231)%64)) | (base.I64_extend_i32_u(base.B2i32(base.Ui64(v272) < base.Ui64(v270)))+base.I64_extend_i32_u(base.B2i32(base.Ui64(v278) < base.Ui64(v272)))+base.I64_extend_i32_u(base.B2i32(base.Ui64(v280) < base.Ui64(v278))))<<(uint(v231)%64))
			v349 = v340 + (base.I64_extend_i32_u(base.B2i32(base.Ui64(v267) < base.Ui64(v259))) + base.I64_extend_i32_u(base.B2i32(base.Ui64(v283) < base.Ui64(v267))))
			v352 = base.I64_extend_i32_u(base.B2i32(base.Ui64(v291) < base.Ui64(v289))) + base.I64_extend_i32_u(base.B2i32(base.Ui64(v295) < base.Ui64(v291))) + v276*v269 + base.I64_extend_i32_u(base.B2i32(base.Ui64(v306) < base.Ui64(v295))) + (base.I64_extend_i32_u(base.B2i32(base.Ui64(v312) < base.Ui64(v310)))<<(uint(v231)%64) | int64(base.Ui64(v312)>>(uint(v231)%64))) + base.I64_extend_i32_u(base.B2i32(base.Ui64(v323) < base.Ui64(v306))) + base.I64_extend_i32_u(base.B2i32(base.Ui64(v340) < base.Ui64(v323))) + base.I64_extend_i32_u(base.B2i32(base.Ui64(v349) < base.Ui64(v340)))
			if v352&int64(281474976710656) == int64(0) {
				v359 = int64(63)
				v361 = int64(1)
				v376 = int64(base.Ui64(v243)>>(uint(v359)%64)) | v284<<(uint(v361)%64)
				v377 = v349<<(uint(v361)%64) | int64(base.Ui64(v284)>>(uint(v359)%64))
				v379 = v352<<(uint(v361)%64) | int64(base.Ui64(v349)>>(uint(v359)%64))
				v380 = v286 + int32(-16383)
				v381 = v243 << (uint(v361) % 64)
			} else {
				v376 = v284
				v377 = v349
				v379 = v352
				v380 = v286 + int32(-16382)
				v381 = v243
			}
			if v380 < int32(32767) {
				if int32(0) < v380 {
					v537 = v376
					v538 = v377
					v539 = base.I64_extend_i32_u(v380)<<(uint(int64(48))%64) | v379&int64(281474976710655)
					v542 = v381
					v543 = v539 | v34
					if v537 == int64(-9223372036854775807-1) {
						v550 = base.B2i32(v542 == int64(0))
					} else {
						v550 = base.B2i32(int64(-1) < v537)
					}
					if v550 != 0 {
						if v542|(v537^int64(-9223372036854775807-1)) == int64(0) {
							v564 = v538 + v538&int64(1)
							v568 = v564
							v573 = v543 + base.I64_extend_i32_u(base.B2i32(base.Ui64(v564) < base.Ui64(v538)))
						} else {
							v568 = v538
							v573 = v543
						}
					} else {
						v552 = v538 + int64(1)
						v568 = v552
						v573 = v543 + base.I64_extend_i32_u(base.B2i32(v552 == int64(0)))
					}
				} else {
					v390 = int32(1) - v380
					if base.Ui32(int32(127)) < base.Ui32(v390) {
						v568 = int64(0)
						v573 = v34
					} else {
						v394 = v28 + int32(48)
						v396 = v380 + int32(127)
						if v396&int32(64) == int32(0) {
							if v396 == int32(0) {
								v417 = v381
								v418 = v376
							} else {
								v413 = base.I64_extend_i32_u(v396)
								v417 = v381 << (uint(v413) % 64)
								v418 = int64(base.Ui64(v381)>>(uint(base.I64_extend_i32_u(int32(64)-v396))%64)) | v376<<(uint(v413)%64)
							}
						} else {
							v417 = int64(0)
							v418 = v381 << (uint(base.I64_extend_i32_u(v380+int32(63))) % 64)
						}
						*(*int64)(unsafe.Add(mBase, uint32(v394))) = v417
						*(*int64)(unsafe.Add(mBase, uint32(v394)+8)) = v418
						v423 = v28 + int32(32)
						if v396&int32(64) == int32(0) {
							if v396 == int32(0) {
								v444 = v377
								v445 = v379
							} else {
								v440 = base.I64_extend_i32_u(v396)
								v444 = v377 << (uint(v440) % 64)
								v445 = int64(base.Ui64(v377)>>(uint(base.I64_extend_i32_u(int32(64)-v396))%64)) | v379<<(uint(v440)%64)
							}
						} else {
							v444 = int64(0)
							v445 = v377 << (uint(base.I64_extend_i32_u(v380+int32(63))) % 64)
						}
						*(*int64)(unsafe.Add(mBase, uint32(v423))) = v444
						*(*int64)(unsafe.Add(mBase, uint32(v423)+8)) = v445
						v450 = v28 + int32(16)
						if v390&int32(64) == int32(0) {
							if v390 == int32(0) {
								v471 = v381
								v472 = v376
							} else {
								v467 = base.I64_extend_i32_u(v390)
								v471 = v376<<(uint(base.I64_extend_i32_u(int32(64)-v390))%64) | int64(base.Ui64(v381)>>(uint(v467)%64))
								v472 = int64(base.Ui64(v376) >> (uint(v467) % 64))
							}
						} else {
							v471 = int64(base.Ui64(v376) >> (uint(base.I64_extend_i32_u(v390+int32(-64))) % 64))
							v472 = int64(0)
						}
						*(*int64)(unsafe.Add(mBase, uint32(v450))) = v471
						*(*int64)(unsafe.Add(mBase, uint32(v450)+8)) = v472
						if v390&int32(64) == int32(0) {
							if v390 == int32(0) {
								v496 = v377
								v497 = v379
							} else {
								v492 = base.I64_extend_i32_u(v390)
								v496 = v379<<(uint(base.I64_extend_i32_u(int32(64)-v390))%64) | int64(base.Ui64(v377)>>(uint(v492)%64))
								v497 = int64(base.Ui64(v379) >> (uint(v492) % 64))
							}
						} else {
							v496 = int64(base.Ui64(v379) >> (uint(base.I64_extend_i32_u(v390+int32(-64))) % 64))
							v497 = int64(0)
						}
						*(*int64)(unsafe.Add(mBase, uint32(v28))) = v496
						*(*int64)(unsafe.Add(mBase, uint32(v28)+8)) = v497
						v501 = *(*int64)(unsafe.Add(mBase, uint32(v28)+32))
						v502 = *(*int64)(unsafe.Add(mBase, uint32(v28)+16))
						v504 = *(*int64)(unsafe.Add(mBase, uint32(v28)+48))
						v509 = *(*int64)(unsafe.Add(mBase, uint32(v28+int32(56))))
						v519 = *(*int64)(unsafe.Add(mBase, uint32(v28+int32(40))))
						v524 = *(*int64)(unsafe.Add(mBase, uint32(v28+int32(24))))
						v528 = *(*int64)(unsafe.Add(mBase, uint32(v28+int32(8))))
						v529 = *(*int64)(unsafe.Add(mBase, uint32(v28)))
						v537 = v519 | v524
						v538 = v529
						v539 = v528
						v542 = v501 | v502 | base.I64_extend_i32_u(base.B2i32(v504|v509 != int64(0)))
						v543 = v539 | v34
						if v537 == int64(-9223372036854775807-1) {
							v550 = base.B2i32(v542 == int64(0))
						} else {
							v550 = base.B2i32(int64(-1) < v537)
						}
						if v550 != 0 {
							if v542|(v537^int64(-9223372036854775807-1)) == int64(0) {
								v564 = v538 + v538&int64(1)
								v568 = v564
								v573 = v543 + base.I64_extend_i32_u(base.B2i32(base.Ui64(v564) < base.Ui64(v538)))
							} else {
								v568 = v538
								v573 = v543
							}
						} else {
							v552 = v538 + int64(1)
							v568 = v552
							v573 = v543 + base.I64_extend_i32_u(base.B2i32(v552 == int64(0)))
						}
					}
				}
			} else {
				v568 = int64(0)
				v573 = v34 | int64(9223090561878065152)
			}
		} else {
			v62 = l2 & int64(9223372036854775807)
			v63 = int64(9223090561878065152)
			if v62 == v63 {
				v67 = base.B2i32(l1 == int64(0))
			} else {
				v67 = base.B2i32(base.Ui64(v62) < base.Ui64(v63))
			}
			if v67 != 0 {
				v73 = l4 & int64(9223372036854775807)
				v74 = int64(9223090561878065152)
				if v73 == v74 {
					v78 = base.B2i32(l3 == int64(0))
				} else {
					v78 = base.B2i32(base.Ui64(v73) < base.Ui64(v74))
				}
				if v78 != 0 {
					if l1|(v62^int64(9223090561878065152)) != int64(0) {
						if l3|(v73^int64(9223090561878065152)) != int64(0) {
							if l1|v62 != int64(0) {
								if l3|v73 != int64(0) {
									if base.Ui64(int64(281474976710655)) < base.Ui64(v62) {
										v167 = l1
										v168 = v36
										v169 = v38
										v170 = int32(0)
									} else {
										v122 = v28 + int32(80)
										v124 = base.B2i32(v36 == int64(0))
										if v36 == int64(0) {
											v125 = l1
										} else {
											v125 = v36
										}
										v131 = base.I32_wrap_i64(base.I64_clz(v125) + base.I64_extend_i32_u(v124<<(uint(int32(6))%32)))
										v133 = v131 + int32(-15)
										if v133&int32(64) == int32(0) {
											if v133 == int32(0) {
												v154 = l1
												v155 = v36
											} else {
												v150 = base.I64_extend_i32_u(v133)
												v154 = l1 << (uint(v150) % 64)
												v155 = int64(base.Ui64(l1)>>(uint(base.I64_extend_i32_u(int32(64)-v133))%64)) | v36<<(uint(v150)%64)
											}
										} else {
											v154 = int64(0)
											v155 = l1 << (uint(base.I64_extend_i32_u(v131+int32(-79))) % 64)
										}
										*(*int64)(unsafe.Add(mBase, uint32(v122))) = v154
										*(*int64)(unsafe.Add(mBase, uint32(v122)+8)) = v155
										v163 = *(*int64)(unsafe.Add(mBase, uint32(v28+int32(88))))
										v166 = *(*int64)(unsafe.Add(mBase, uint32(v28)+80))
										v167 = v166
										v168 = v163
										v169 = int64(base.Ui64(v163) >> (uint(int64(32)) % 64))
										v170 = int32(16) - v131
									}
									if base.Ui64(int64(281474976710655)) < base.Ui64(v73) {
										v218 = v167
										v220 = l3
										v221 = v31
										v222 = v168
										v223 = v169
										v224 = v170
									} else {
										v174 = v28 + int32(64)
										v176 = base.B2i32(v31 == int64(0))
										if v31 == int64(0) {
											v177 = l3
										} else {
											v177 = v31
										}
										v183 = base.I32_wrap_i64(base.I64_clz(v177) + base.I64_extend_i32_u(v176<<(uint(int32(6))%32)))
										v185 = v183 + int32(-15)
										if v185&int32(64) == int32(0) {
											if v185 == int32(0) {
												v206 = l3
												v207 = v31
											} else {
												v202 = base.I64_extend_i32_u(v185)
												v206 = l3 << (uint(v202) % 64)
												v207 = int64(base.Ui64(l3)>>(uint(base.I64_extend_i32_u(int32(64)-v185))%64)) | v31<<(uint(v202)%64)
											}
										} else {
											v206 = int64(0)
											v207 = l3 << (uint(base.I64_extend_i32_u(v183+int32(-79))) % 64)
										}
										*(*int64)(unsafe.Add(mBase, uint32(v174))) = v206
										*(*int64)(unsafe.Add(mBase, uint32(v174)+8)) = v207
										v216 = *(*int64)(unsafe.Add(mBase, uint32(v28+int32(72))))
										v217 = *(*int64)(unsafe.Add(mBase, uint32(v28)+64))
										v218 = v167
										v220 = v217
										v221 = v216
										v222 = v168
										v223 = v169
										v224 = v170 - v183 + int32(16)
									}
									v227 = int64(15)
									v228 = v220 << (uint(v227) % 64)
									v230 = v228 & int64(4294934528)
									v231 = int64(32)
									v232 = int64(base.Ui64(v218) >> (uint(v231) % 64))
									v233 = v230 * v232
									v235 = int64(base.Ui64(v228) >> (uint(v231) % 64))
									v236 = int64(4294967295)
									v237 = v218 & v236
									v239 = v233 + v235*v237
									v241 = v239 << (uint(v231) % 64)
									v243 = v241 + v230*v237
									v247 = v222 & v236
									v248 = v230 * v247
									v250 = v248 + v235*v232
									v254 = v221 << (uint(v227) % 64)
									v257 = (int64(base.Ui64(v220)>>(uint(int64(49))%64)) | v254) & v236
									v259 = v250 + v257*v237
									v267 = v259 + (int64(base.Ui64(v239)>>(uint(v231)%64)) | base.I64_extend_i32_u(base.B2i32(base.Ui64(v239) < base.Ui64(v233)))<<(uint(v231)%64))
									v269 = v223 | int64(65536)
									v270 = v230 * v269
									v272 = v270 + v235*v247
									v276 = int64(base.Ui64(v254)>>(uint(v231)%64)) | int64(2147483648)
									v278 = v272 + v276*v237
									v280 = v278 + v257*v232
									v283 = v267 + v280<<(uint(v231)%64)
									v284 = base.I64_extend_i32_u(base.B2i32(base.Ui64(v243) < base.Ui64(v241))) + v283
									v286 = v48 + v43 + v224
									v289 = v276 * v232
									v291 = v289 + v235*v269
									v295 = v291 + v257*v247
									v306 = v295 + (base.I64_extend_i32_u(base.B2i32(base.Ui64(v250) < base.Ui64(v248))) + base.I64_extend_i32_u(base.B2i32(base.Ui64(v259) < base.Ui64(v250))))
									v310 = v257 * v269
									v312 = v310 + v276*v247
									v323 = v306 + v312<<(uint(v231)%64)
									v340 = v323 + (int64(base.Ui64(v280)>>(uint(v231)%64)) | (base.I64_extend_i32_u(base.B2i32(base.Ui64(v272) < base.Ui64(v270)))+base.I64_extend_i32_u(base.B2i32(base.Ui64(v278) < base.Ui64(v272)))+base.I64_extend_i32_u(base.B2i32(base.Ui64(v280) < base.Ui64(v278))))<<(uint(v231)%64))
									v349 = v340 + (base.I64_extend_i32_u(base.B2i32(base.Ui64(v267) < base.Ui64(v259))) + base.I64_extend_i32_u(base.B2i32(base.Ui64(v283) < base.Ui64(v267))))
									v352 = base.I64_extend_i32_u(base.B2i32(base.Ui64(v291) < base.Ui64(v289))) + base.I64_extend_i32_u(base.B2i32(base.Ui64(v295) < base.Ui64(v291))) + v276*v269 + base.I64_extend_i32_u(base.B2i32(base.Ui64(v306) < base.Ui64(v295))) + (base.I64_extend_i32_u(base.B2i32(base.Ui64(v312) < base.Ui64(v310)))<<(uint(v231)%64) | int64(base.Ui64(v312)>>(uint(v231)%64))) + base.I64_extend_i32_u(base.B2i32(base.Ui64(v323) < base.Ui64(v306))) + base.I64_extend_i32_u(base.B2i32(base.Ui64(v340) < base.Ui64(v323))) + base.I64_extend_i32_u(base.B2i32(base.Ui64(v349) < base.Ui64(v340)))
									if v352&int64(281474976710656) == int64(0) {
										v359 = int64(63)
										v361 = int64(1)
										v376 = int64(base.Ui64(v243)>>(uint(v359)%64)) | v284<<(uint(v361)%64)
										v377 = v349<<(uint(v361)%64) | int64(base.Ui64(v284)>>(uint(v359)%64))
										v379 = v352<<(uint(v361)%64) | int64(base.Ui64(v349)>>(uint(v359)%64))
										v380 = v286 + int32(-16383)
										v381 = v243 << (uint(v361) % 64)
									} else {
										v376 = v284
										v377 = v349
										v379 = v352
										v380 = v286 + int32(-16382)
										v381 = v243
									}
									if v380 < int32(32767) {
										if int32(0) < v380 {
											v537 = v376
											v538 = v377
											v539 = base.I64_extend_i32_u(v380)<<(uint(int64(48))%64) | v379&int64(281474976710655)
											v542 = v381
											v543 = v539 | v34
											if v537 == int64(-9223372036854775807-1) {
												v550 = base.B2i32(v542 == int64(0))
											} else {
												v550 = base.B2i32(int64(-1) < v537)
											}
											if v550 != 0 {
												if v542|(v537^int64(-9223372036854775807-1)) == int64(0) {
													v564 = v538 + v538&int64(1)
													v568 = v564
													v573 = v543 + base.I64_extend_i32_u(base.B2i32(base.Ui64(v564) < base.Ui64(v538)))
												} else {
													v568 = v538
													v573 = v543
												}
											} else {
												v552 = v538 + int64(1)
												v568 = v552
												v573 = v543 + base.I64_extend_i32_u(base.B2i32(v552 == int64(0)))
											}
										} else {
											v390 = int32(1) - v380
											if base.Ui32(int32(127)) < base.Ui32(v390) {
												v568 = int64(0)
												v573 = v34
											} else {
												v394 = v28 + int32(48)
												v396 = v380 + int32(127)
												if v396&int32(64) == int32(0) {
													if v396 == int32(0) {
														v417 = v381
														v418 = v376
													} else {
														v413 = base.I64_extend_i32_u(v396)
														v417 = v381 << (uint(v413) % 64)
														v418 = int64(base.Ui64(v381)>>(uint(base.I64_extend_i32_u(int32(64)-v396))%64)) | v376<<(uint(v413)%64)
													}
												} else {
													v417 = int64(0)
													v418 = v381 << (uint(base.I64_extend_i32_u(v380+int32(63))) % 64)
												}
												*(*int64)(unsafe.Add(mBase, uint32(v394))) = v417
												*(*int64)(unsafe.Add(mBase, uint32(v394)+8)) = v418
												v423 = v28 + int32(32)
												if v396&int32(64) == int32(0) {
													if v396 == int32(0) {
														v444 = v377
														v445 = v379
													} else {
														v440 = base.I64_extend_i32_u(v396)
														v444 = v377 << (uint(v440) % 64)
														v445 = int64(base.Ui64(v377)>>(uint(base.I64_extend_i32_u(int32(64)-v396))%64)) | v379<<(uint(v440)%64)
													}
												} else {
													v444 = int64(0)
													v445 = v377 << (uint(base.I64_extend_i32_u(v380+int32(63))) % 64)
												}
												*(*int64)(unsafe.Add(mBase, uint32(v423))) = v444
												*(*int64)(unsafe.Add(mBase, uint32(v423)+8)) = v445
												v450 = v28 + int32(16)
												if v390&int32(64) == int32(0) {
													if v390 == int32(0) {
														v471 = v381
														v472 = v376
													} else {
														v467 = base.I64_extend_i32_u(v390)
														v471 = v376<<(uint(base.I64_extend_i32_u(int32(64)-v390))%64) | int64(base.Ui64(v381)>>(uint(v467)%64))
														v472 = int64(base.Ui64(v376) >> (uint(v467) % 64))
													}
												} else {
													v471 = int64(base.Ui64(v376) >> (uint(base.I64_extend_i32_u(v390+int32(-64))) % 64))
													v472 = int64(0)
												}
												*(*int64)(unsafe.Add(mBase, uint32(v450))) = v471
												*(*int64)(unsafe.Add(mBase, uint32(v450)+8)) = v472
												if v390&int32(64) == int32(0) {
													if v390 == int32(0) {
														v496 = v377
														v497 = v379
													} else {
														v492 = base.I64_extend_i32_u(v390)
														v496 = v379<<(uint(base.I64_extend_i32_u(int32(64)-v390))%64) | int64(base.Ui64(v377)>>(uint(v492)%64))
														v497 = int64(base.Ui64(v379) >> (uint(v492) % 64))
													}
												} else {
													v496 = int64(base.Ui64(v379) >> (uint(base.I64_extend_i32_u(v390+int32(-64))) % 64))
													v497 = int64(0)
												}
												*(*int64)(unsafe.Add(mBase, uint32(v28))) = v496
												*(*int64)(unsafe.Add(mBase, uint32(v28)+8)) = v497
												v501 = *(*int64)(unsafe.Add(mBase, uint32(v28)+32))
												v502 = *(*int64)(unsafe.Add(mBase, uint32(v28)+16))
												v504 = *(*int64)(unsafe.Add(mBase, uint32(v28)+48))
												v509 = *(*int64)(unsafe.Add(mBase, uint32(v28+int32(56))))
												v519 = *(*int64)(unsafe.Add(mBase, uint32(v28+int32(40))))
												v524 = *(*int64)(unsafe.Add(mBase, uint32(v28+int32(24))))
												v528 = *(*int64)(unsafe.Add(mBase, uint32(v28+int32(8))))
												v529 = *(*int64)(unsafe.Add(mBase, uint32(v28)))
												v537 = v519 | v524
												v538 = v529
												v539 = v528
												v542 = v501 | v502 | base.I64_extend_i32_u(base.B2i32(v504|v509 != int64(0)))
												v543 = v539 | v34
												if v537 == int64(-9223372036854775807-1) {
													v550 = base.B2i32(v542 == int64(0))
												} else {
													v550 = base.B2i32(int64(-1) < v537)
												}
												if v550 != 0 {
													if v542|(v537^int64(-9223372036854775807-1)) == int64(0) {
														v564 = v538 + v538&int64(1)
														v568 = v564
														v573 = v543 + base.I64_extend_i32_u(base.B2i32(base.Ui64(v564) < base.Ui64(v538)))
													} else {
														v568 = v538
														v573 = v543
													}
												} else {
													v552 = v538 + int64(1)
													v568 = v552
													v573 = v543 + base.I64_extend_i32_u(base.B2i32(v552 == int64(0)))
												}
											}
										}
									} else {
										v568 = int64(0)
										v573 = v34 | int64(9223090561878065152)
									}
								} else {
									v568 = int64(0)
									v573 = v34
								}
							} else {
								v568 = int64(0)
								v573 = v34
							}
						} else {
							v102 = int64(0)
							if base.B2i32(l1|v62 == v102) == int32(0) {
								v568 = v102
								v573 = v34 | int64(9223090561878065152)
							} else {
								v568 = v102
								v573 = int64(9223231299366420480)
							}
						}
					} else {
						if base.B2i32(l3|v73 == int64(0)) == int32(0) {
							v568 = int64(0)
							v573 = v34 | int64(9223090561878065152)
						} else {
							v568 = int64(0)
							v573 = int64(9223231299366420480)
						}
					}
				} else {
					v568 = l3
					v573 = l4 | int64(140737488355328)
				}
			} else {
				v568 = l1
				v573 = l2 | int64(140737488355328)
			}
		}
	}
	*(*int64)(unsafe.Add(mBase, uint32(l0))) = v568
	*(*int64)(unsafe.Add(mBase, uint32(l0)+8)) = v573
	m.G0 = v28 + int32(96)
	return
}
func F___multi3(m *base.Module, l0 int32, l1 int64, l2 int64, l3 int64, l4 int64) {
	mBase := m.M
	_ = mBase
	var v10 int64
	_ = v10
	var v11 int64
	_ = v11
	var v13 int64
	_ = v13
	var v16 int64
	_ = v16
	var v17 int64
	_ = v17
	var v19 int64
	_ = v19
	var v20 int64
	_ = v20
	var v24 int64
	_ = v24
	var v31 int64
	_ = v31
	v10 = int64(32)
	v11 = int64(base.Ui64(l3) >> (uint(v10) % 64))
	v13 = int64(base.Ui64(l1) >> (uint(v10) % 64))
	v16 = int64(4294967295)
	v17 = l3 & v16
	v19 = l1 & v16
	v20 = v17 * v19
	v24 = int64(base.Ui64(v20)>>(uint(v10)%64)) + v17*v13
	v31 = v24&v16 + v11*v19
	*(*int64)(unsafe.Add(mBase, uint32(l0)+8)) = l4*l1 + l2*l3 + v11*v13 + int64(base.Ui64(v24)>>(uint(v10)%64)) + int64(base.Ui64(v31)>>(uint(v10)%64))
	*(*int64)(unsafe.Add(mBase, uint32(l0))) = v31<<(uint(v10)%64) | v20&v16
	return
}
func F_markmt(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
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
	var v47 int32
	_ = v47
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v102 int32
	_ = v102
	var v104 int32
	_ = v104
	var v107 int32
	_ = v107
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v134 int32
	_ = v134
	var v136 int32
	_ = v136
	var v138 int32
	_ = v138
	var v140 int32
	_ = v140
	var v142 int32
	_ = v142
	var v144 int32
	_ = v144
	var v153 int32
	_ = v153
	var v156 int32
	_ = v156
	var v163 int32
	_ = v163
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v169 int32
	_ = v169
	var v171 int32
	_ = v171
	var v177 int32
	_ = v177
	var v179 int32
	_ = v179
	var v182 int32
	_ = v182
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v209 int32
	_ = v209
	var v211 int32
	_ = v211
	var v213 int32
	_ = v213
	var v215 int32
	_ = v215
	var v217 int32
	_ = v217
	var v219 int32
	_ = v219
	var v228 int32
	_ = v228
	var v231 int32
	_ = v231
	var v238 int32
	_ = v238
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v244 int32
	_ = v244
	var v246 int32
	_ = v246
	var v252 int32
	_ = v252
	var v254 int32
	_ = v254
	var v257 int32
	_ = v257
	var v263 int32
	_ = v263
	var v264 int32
	_ = v264
	var v267 int32
	_ = v267
	var v268 int32
	_ = v268
	var v271 int32
	_ = v271
	var v272 int32
	_ = v272
	var v278 int32
	_ = v278
	var v279 int32
	_ = v279
	var v284 int32
	_ = v284
	var v286 int32
	_ = v286
	var v288 int32
	_ = v288
	var v290 int32
	_ = v290
	var v292 int32
	_ = v292
	var v294 int32
	_ = v294
	var v303 int32
	_ = v303
	var v306 int32
	_ = v306
	var v313 int32
	_ = v313
	var v315 int32
	_ = v315
	var v316 int32
	_ = v316
	var v319 int32
	_ = v319
	var v321 int32
	_ = v321
	var v327 int32
	_ = v327
	var v329 int32
	_ = v329
	var v332 int32
	_ = v332
	var v338 int32
	_ = v338
	var v339 int32
	_ = v339
	var v342 int32
	_ = v342
	var v343 int32
	_ = v343
	var v346 int32
	_ = v346
	var v347 int32
	_ = v347
	var v353 int32
	_ = v353
	var v354 int32
	_ = v354
	var v359 int32
	_ = v359
	var v361 int32
	_ = v361
	var v363 int32
	_ = v363
	var v365 int32
	_ = v365
	var v367 int32
	_ = v367
	var v369 int32
	_ = v369
	var v378 int32
	_ = v378
	var v381 int32
	_ = v381
	var v388 int32
	_ = v388
	var v390 int32
	_ = v390
	var v391 int32
	_ = v391
	var v394 int32
	_ = v394
	var v396 int32
	_ = v396
	var v402 int32
	_ = v402
	var v404 int32
	_ = v404
	var v407 int32
	_ = v407
	var v413 int32
	_ = v413
	var v414 int32
	_ = v414
	var v417 int32
	_ = v417
	var v418 int32
	_ = v418
	var v421 int32
	_ = v421
	var v422 int32
	_ = v422
	var v428 int32
	_ = v428
	var v429 int32
	_ = v429
	var v434 int32
	_ = v434
	var v436 int32
	_ = v436
	var v438 int32
	_ = v438
	var v440 int32
	_ = v440
	var v442 int32
	_ = v442
	var v444 int32
	_ = v444
	var v453 int32
	_ = v453
	var v456 int32
	_ = v456
	var v463 int32
	_ = v463
	var v465 int32
	_ = v465
	var v466 int32
	_ = v466
	var v469 int32
	_ = v469
	var v471 int32
	_ = v471
	var v477 int32
	_ = v477
	var v479 int32
	_ = v479
	var v482 int32
	_ = v482
	var v488 int32
	_ = v488
	var v489 int32
	_ = v489
	var v492 int32
	_ = v492
	var v493 int32
	_ = v493
	var v496 int32
	_ = v496
	var v497 int32
	_ = v497
	var v503 int32
	_ = v503
	var v504 int32
	_ = v504
	var v509 int32
	_ = v509
	var v511 int32
	_ = v511
	var v513 int32
	_ = v513
	var v515 int32
	_ = v515
	var v517 int32
	_ = v517
	var v519 int32
	_ = v519
	var v528 int32
	_ = v528
	var v531 int32
	_ = v531
	var v538 int32
	_ = v538
	var v540 int32
	_ = v540
	var v541 int32
	_ = v541
	var v544 int32
	_ = v544
	var v546 int32
	_ = v546
	var v552 int32
	_ = v552
	var v554 int32
	_ = v554
	var v557 int32
	_ = v557
	var v563 int32
	_ = v563
	var v564 int32
	_ = v564
	var v567 int32
	_ = v567
	var v568 int32
	_ = v568
	var v571 int32
	_ = v571
	var v572 int32
	_ = v572
	var v578 int32
	_ = v578
	var v579 int32
	_ = v579
	var v584 int32
	_ = v584
	var v586 int32
	_ = v586
	var v588 int32
	_ = v588
	var v590 int32
	_ = v590
	var v592 int32
	_ = v592
	var v594 int32
	_ = v594
	var v603 int32
	_ = v603
	var v606 int32
	_ = v606
	var v613 int32
	_ = v613
	var v615 int32
	_ = v615
	var v616 int32
	_ = v616
	var v619 int32
	_ = v619
	var v621 int32
	_ = v621
	var v627 int32
	_ = v627
	var v629 int32
	_ = v629
	var v632 int32
	_ = v632
	var v638 int32
	_ = v638
	var v639 int32
	_ = v639
	var v642 int32
	_ = v642
	var v643 int32
	_ = v643
	var v646 int32
	_ = v646
	var v647 int32
	_ = v647
	var v653 int32
	_ = v653
	var v654 int32
	_ = v654
	var v659 int32
	_ = v659
	var v661 int32
	_ = v661
	var v663 int32
	_ = v663
	var v665 int32
	_ = v665
	var v667 int32
	_ = v667
	var v669 int32
	_ = v669
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+152))
	if v3 == int32(0) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v78 = *(*int32)(unsafe.Add(mBase, uint32(l0)+156))
	if v78 == int32(0) {
		goto L24
	} else {
		goto L25
	}
L2:
	;
	v6 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3)+5)))
	if v6&int32(3) == int32(0) {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v13 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3)+5)))
	v15 = v3
	v16 = v13
	goto L12
L4:
	;
	goto L1
L5:
	;
	goto L4
L6:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = v15
	goto L5
L7:
	;
	v69 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+68)) = v69
	goto L6
L8:
	;
	v67 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+108)) = v67
	goto L6
L9:
	;
	v65 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+32)) = v65
	goto L6
L10:
	;
	v63 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+8)) = v63
	goto L6
L11:
	;
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v15)+8))
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v42)+8))
	if v43 < int32(4) {
		v54 = v42
		goto L20
	} else {
		goto L21
	}
L12:
	;
	v19 = v16 & int32(252)
	*(*uint8)(unsafe.Add(mBase, uint32(v15)+5)) = uint8(v19)
	v21 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15)+4)))
	if v21 == int32(7) {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v27 = v19 | int32(4)
	*(*uint8)(unsafe.Add(mBase, uint32(v15)+5)) = uint8(v27)
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v15)+8))
	if v29 == int32(0) {
		goto L16
	} else {
		goto L17
	}
L15:
	;
	switch v21 + int32(-5) {
	case 0:
		goto L9
	case 1:
		goto L10
	default:
		goto L5
	case 3:
		goto L8
	case 4:
		goto L7
	case 5:
		goto L11
	}
L16:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v15)+12))
	v39 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v38)+5)))
	if v39&int32(3) != 0 {
		v15 = v38
		v16 = v39
		goto L12
	} else {
		goto L19
	}
L17:
	;
	v32 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v29)+5)))
	if v32&int32(3) == int32(0) {
		goto L16
	} else {
		goto L18
	}
L18:
	;
	F_reallymarkobject(m, l0, v29)
	mBase = m.M
	goto L16
L19:
	;
	goto L5
L20:
	;
	if v54 != v15+int32(16) {
		goto L5
	} else {
		goto L23
	}
L21:
	;
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v42)))
	v47 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v46)+5)))
	if v47&int32(3) == int32(0) {
		v54 = v42
		goto L20
	} else {
		goto L22
	}
L22:
	;
	F_reallymarkobject(m, l0, v46)
	mBase = m.M
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v15)+8))
	v54 = v53
	goto L20
L23:
	;
	v59 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15)+5)))
	v61 = v59 | int32(4)
	*(*uint8)(unsafe.Add(mBase, uint32(v15)+5)) = uint8(v61)
	goto L4
L24:
	;
	v153 = *(*int32)(unsafe.Add(mBase, uint32(l0)+160))
	if v153 == int32(0) {
		goto L47
	} else {
		goto L48
	}
L25:
	;
	v81 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v78)+5)))
	if v81&int32(3) == int32(0) {
		goto L24
	} else {
		goto L26
	}
L26:
	;
	v88 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v78)+5)))
	v90 = v78
	v91 = v88
	goto L35
L27:
	;
	goto L24
L28:
	;
	goto L27
L29:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = v90
	goto L28
L30:
	;
	v144 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v90)+68)) = v144
	goto L29
L31:
	;
	v142 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v90)+108)) = v142
	goto L29
L32:
	;
	v140 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v90)+32)) = v140
	goto L29
L33:
	;
	v138 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v90)+8)) = v138
	goto L29
L34:
	;
	v117 = *(*int32)(unsafe.Add(mBase, uint32(v90)+8))
	v118 = *(*int32)(unsafe.Add(mBase, uint32(v117)+8))
	if v118 < int32(4) {
		v129 = v117
		goto L43
	} else {
		goto L44
	}
L35:
	;
	v94 = v91 & int32(252)
	*(*uint8)(unsafe.Add(mBase, uint32(v90)+5)) = uint8(v94)
	v96 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v90)+4)))
	if v96 == int32(7) {
		goto L37
	} else {
		goto L38
	}
L37:
	;
	v102 = v94 | int32(4)
	*(*uint8)(unsafe.Add(mBase, uint32(v90)+5)) = uint8(v102)
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v90)+8))
	if v104 == int32(0) {
		goto L39
	} else {
		goto L40
	}
L38:
	;
	switch v96 + int32(-5) {
	case 0:
		goto L32
	case 1:
		goto L33
	default:
		goto L28
	case 3:
		goto L31
	case 4:
		goto L30
	case 5:
		goto L34
	}
L39:
	;
	v113 = *(*int32)(unsafe.Add(mBase, uint32(v90)+12))
	v114 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v113)+5)))
	if v114&int32(3) != 0 {
		v90 = v113
		v91 = v114
		goto L35
	} else {
		goto L42
	}
L40:
	;
	v107 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v104)+5)))
	if v107&int32(3) == int32(0) {
		goto L39
	} else {
		goto L41
	}
L41:
	;
	F_reallymarkobject(m, l0, v104)
	mBase = m.M
	goto L39
L42:
	;
	goto L28
L43:
	;
	if v129 != v90+int32(16) {
		goto L28
	} else {
		goto L46
	}
L44:
	;
	v121 = *(*int32)(unsafe.Add(mBase, uint32(v117)))
	v122 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v121)+5)))
	if v122&int32(3) == int32(0) {
		v129 = v117
		goto L43
	} else {
		goto L45
	}
L45:
	;
	F_reallymarkobject(m, l0, v121)
	mBase = m.M
	v128 = *(*int32)(unsafe.Add(mBase, uint32(v90)+8))
	v129 = v128
	goto L43
L46:
	;
	v134 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v90)+5)))
	v136 = v134 | int32(4)
	*(*uint8)(unsafe.Add(mBase, uint32(v90)+5)) = uint8(v136)
	goto L27
L47:
	;
	v228 = *(*int32)(unsafe.Add(mBase, uint32(l0)+164))
	if v228 == int32(0) {
		goto L70
	} else {
		goto L71
	}
L48:
	;
	v156 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v153)+5)))
	if v156&int32(3) == int32(0) {
		goto L47
	} else {
		goto L49
	}
L49:
	;
	v163 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v153)+5)))
	v165 = v153
	v166 = v163
	goto L58
L50:
	;
	goto L47
L51:
	;
	goto L50
L52:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = v165
	goto L51
L53:
	;
	v219 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v165)+68)) = v219
	goto L52
L54:
	;
	v217 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v165)+108)) = v217
	goto L52
L55:
	;
	v215 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v165)+32)) = v215
	goto L52
L56:
	;
	v213 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v165)+8)) = v213
	goto L52
L57:
	;
	v192 = *(*int32)(unsafe.Add(mBase, uint32(v165)+8))
	v193 = *(*int32)(unsafe.Add(mBase, uint32(v192)+8))
	if v193 < int32(4) {
		v204 = v192
		goto L66
	} else {
		goto L67
	}
L58:
	;
	v169 = v166 & int32(252)
	*(*uint8)(unsafe.Add(mBase, uint32(v165)+5)) = uint8(v169)
	v171 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v165)+4)))
	if v171 == int32(7) {
		goto L60
	} else {
		goto L61
	}
L60:
	;
	v177 = v169 | int32(4)
	*(*uint8)(unsafe.Add(mBase, uint32(v165)+5)) = uint8(v177)
	v179 = *(*int32)(unsafe.Add(mBase, uint32(v165)+8))
	if v179 == int32(0) {
		goto L62
	} else {
		goto L63
	}
L61:
	;
	switch v171 + int32(-5) {
	case 0:
		goto L55
	case 1:
		goto L56
	default:
		goto L51
	case 3:
		goto L54
	case 4:
		goto L53
	case 5:
		goto L57
	}
L62:
	;
	v188 = *(*int32)(unsafe.Add(mBase, uint32(v165)+12))
	v189 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v188)+5)))
	if v189&int32(3) != 0 {
		v165 = v188
		v166 = v189
		goto L58
	} else {
		goto L65
	}
L63:
	;
	v182 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v179)+5)))
	if v182&int32(3) == int32(0) {
		goto L62
	} else {
		goto L64
	}
L64:
	;
	F_reallymarkobject(m, l0, v179)
	mBase = m.M
	goto L62
L65:
	;
	goto L51
L66:
	;
	if v204 != v165+int32(16) {
		goto L51
	} else {
		goto L69
	}
L67:
	;
	v196 = *(*int32)(unsafe.Add(mBase, uint32(v192)))
	v197 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v196)+5)))
	if v197&int32(3) == int32(0) {
		v204 = v192
		goto L66
	} else {
		goto L68
	}
L68:
	;
	F_reallymarkobject(m, l0, v196)
	mBase = m.M
	v203 = *(*int32)(unsafe.Add(mBase, uint32(v165)+8))
	v204 = v203
	goto L66
L69:
	;
	v209 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v165)+5)))
	v211 = v209 | int32(4)
	*(*uint8)(unsafe.Add(mBase, uint32(v165)+5)) = uint8(v211)
	goto L50
L70:
	;
	v303 = *(*int32)(unsafe.Add(mBase, uint32(l0)+168))
	if v303 == int32(0) {
		goto L93
	} else {
		goto L94
	}
L71:
	;
	v231 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v228)+5)))
	if v231&int32(3) == int32(0) {
		goto L70
	} else {
		goto L72
	}
L72:
	;
	v238 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v228)+5)))
	v240 = v228
	v241 = v238
	goto L81
L73:
	;
	goto L70
L74:
	;
	goto L73
L75:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = v240
	goto L74
L76:
	;
	v294 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v240)+68)) = v294
	goto L75
L77:
	;
	v292 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v240)+108)) = v292
	goto L75
L78:
	;
	v290 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v240)+32)) = v290
	goto L75
L79:
	;
	v288 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v240)+8)) = v288
	goto L75
L80:
	;
	v267 = *(*int32)(unsafe.Add(mBase, uint32(v240)+8))
	v268 = *(*int32)(unsafe.Add(mBase, uint32(v267)+8))
	if v268 < int32(4) {
		v279 = v267
		goto L89
	} else {
		goto L90
	}
L81:
	;
	v244 = v241 & int32(252)
	*(*uint8)(unsafe.Add(mBase, uint32(v240)+5)) = uint8(v244)
	v246 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v240)+4)))
	if v246 == int32(7) {
		goto L83
	} else {
		goto L84
	}
L83:
	;
	v252 = v244 | int32(4)
	*(*uint8)(unsafe.Add(mBase, uint32(v240)+5)) = uint8(v252)
	v254 = *(*int32)(unsafe.Add(mBase, uint32(v240)+8))
	if v254 == int32(0) {
		goto L85
	} else {
		goto L86
	}
L84:
	;
	switch v246 + int32(-5) {
	case 0:
		goto L78
	case 1:
		goto L79
	default:
		goto L74
	case 3:
		goto L77
	case 4:
		goto L76
	case 5:
		goto L80
	}
L85:
	;
	v263 = *(*int32)(unsafe.Add(mBase, uint32(v240)+12))
	v264 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v263)+5)))
	if v264&int32(3) != 0 {
		v240 = v263
		v241 = v264
		goto L81
	} else {
		goto L88
	}
L86:
	;
	v257 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v254)+5)))
	if v257&int32(3) == int32(0) {
		goto L85
	} else {
		goto L87
	}
L87:
	;
	F_reallymarkobject(m, l0, v254)
	mBase = m.M
	goto L85
L88:
	;
	goto L74
L89:
	;
	if v279 != v240+int32(16) {
		goto L74
	} else {
		goto L92
	}
L90:
	;
	v271 = *(*int32)(unsafe.Add(mBase, uint32(v267)))
	v272 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v271)+5)))
	if v272&int32(3) == int32(0) {
		v279 = v267
		goto L89
	} else {
		goto L91
	}
L91:
	;
	F_reallymarkobject(m, l0, v271)
	mBase = m.M
	v278 = *(*int32)(unsafe.Add(mBase, uint32(v240)+8))
	v279 = v278
	goto L89
L92:
	;
	v284 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v240)+5)))
	v286 = v284 | int32(4)
	*(*uint8)(unsafe.Add(mBase, uint32(v240)+5)) = uint8(v286)
	goto L73
L93:
	;
	v378 = *(*int32)(unsafe.Add(mBase, uint32(l0)+172))
	if v378 == int32(0) {
		goto L116
	} else {
		goto L117
	}
L94:
	;
	v306 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v303)+5)))
	if v306&int32(3) == int32(0) {
		goto L93
	} else {
		goto L95
	}
L95:
	;
	v313 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v303)+5)))
	v315 = v303
	v316 = v313
	goto L104
L96:
	;
	goto L93
L97:
	;
	goto L96
L98:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = v315
	goto L97
L99:
	;
	v369 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v315)+68)) = v369
	goto L98
L100:
	;
	v367 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v315)+108)) = v367
	goto L98
L101:
	;
	v365 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v315)+32)) = v365
	goto L98
L102:
	;
	v363 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v315)+8)) = v363
	goto L98
L103:
	;
	v342 = *(*int32)(unsafe.Add(mBase, uint32(v315)+8))
	v343 = *(*int32)(unsafe.Add(mBase, uint32(v342)+8))
	if v343 < int32(4) {
		v354 = v342
		goto L112
	} else {
		goto L113
	}
L104:
	;
	v319 = v316 & int32(252)
	*(*uint8)(unsafe.Add(mBase, uint32(v315)+5)) = uint8(v319)
	v321 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v315)+4)))
	if v321 == int32(7) {
		goto L106
	} else {
		goto L107
	}
L106:
	;
	v327 = v319 | int32(4)
	*(*uint8)(unsafe.Add(mBase, uint32(v315)+5)) = uint8(v327)
	v329 = *(*int32)(unsafe.Add(mBase, uint32(v315)+8))
	if v329 == int32(0) {
		goto L108
	} else {
		goto L109
	}
L107:
	;
	switch v321 + int32(-5) {
	case 0:
		goto L101
	case 1:
		goto L102
	default:
		goto L97
	case 3:
		goto L100
	case 4:
		goto L99
	case 5:
		goto L103
	}
L108:
	;
	v338 = *(*int32)(unsafe.Add(mBase, uint32(v315)+12))
	v339 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v338)+5)))
	if v339&int32(3) != 0 {
		v315 = v338
		v316 = v339
		goto L104
	} else {
		goto L111
	}
L109:
	;
	v332 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v329)+5)))
	if v332&int32(3) == int32(0) {
		goto L108
	} else {
		goto L110
	}
L110:
	;
	F_reallymarkobject(m, l0, v329)
	mBase = m.M
	goto L108
L111:
	;
	goto L97
L112:
	;
	if v354 != v315+int32(16) {
		goto L97
	} else {
		goto L115
	}
L113:
	;
	v346 = *(*int32)(unsafe.Add(mBase, uint32(v342)))
	v347 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v346)+5)))
	if v347&int32(3) == int32(0) {
		v354 = v342
		goto L112
	} else {
		goto L114
	}
L114:
	;
	F_reallymarkobject(m, l0, v346)
	mBase = m.M
	v353 = *(*int32)(unsafe.Add(mBase, uint32(v315)+8))
	v354 = v353
	goto L112
L115:
	;
	v359 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v315)+5)))
	v361 = v359 | int32(4)
	*(*uint8)(unsafe.Add(mBase, uint32(v315)+5)) = uint8(v361)
	goto L96
L116:
	;
	v453 = *(*int32)(unsafe.Add(mBase, uint32(l0)+176))
	if v453 == int32(0) {
		goto L139
	} else {
		goto L140
	}
L117:
	;
	v381 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v378)+5)))
	if v381&int32(3) == int32(0) {
		goto L116
	} else {
		goto L118
	}
L118:
	;
	v388 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v378)+5)))
	v390 = v378
	v391 = v388
	goto L127
L119:
	;
	goto L116
L120:
	;
	goto L119
L121:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = v390
	goto L120
L122:
	;
	v444 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v390)+68)) = v444
	goto L121
L123:
	;
	v442 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v390)+108)) = v442
	goto L121
L124:
	;
	v440 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v390)+32)) = v440
	goto L121
L125:
	;
	v438 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v390)+8)) = v438
	goto L121
L126:
	;
	v417 = *(*int32)(unsafe.Add(mBase, uint32(v390)+8))
	v418 = *(*int32)(unsafe.Add(mBase, uint32(v417)+8))
	if v418 < int32(4) {
		v429 = v417
		goto L135
	} else {
		goto L136
	}
L127:
	;
	v394 = v391 & int32(252)
	*(*uint8)(unsafe.Add(mBase, uint32(v390)+5)) = uint8(v394)
	v396 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v390)+4)))
	if v396 == int32(7) {
		goto L129
	} else {
		goto L130
	}
L129:
	;
	v402 = v394 | int32(4)
	*(*uint8)(unsafe.Add(mBase, uint32(v390)+5)) = uint8(v402)
	v404 = *(*int32)(unsafe.Add(mBase, uint32(v390)+8))
	if v404 == int32(0) {
		goto L131
	} else {
		goto L132
	}
L130:
	;
	switch v396 + int32(-5) {
	case 0:
		goto L124
	case 1:
		goto L125
	default:
		goto L120
	case 3:
		goto L123
	case 4:
		goto L122
	case 5:
		goto L126
	}
L131:
	;
	v413 = *(*int32)(unsafe.Add(mBase, uint32(v390)+12))
	v414 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v413)+5)))
	if v414&int32(3) != 0 {
		v390 = v413
		v391 = v414
		goto L127
	} else {
		goto L134
	}
L132:
	;
	v407 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v404)+5)))
	if v407&int32(3) == int32(0) {
		goto L131
	} else {
		goto L133
	}
L133:
	;
	F_reallymarkobject(m, l0, v404)
	mBase = m.M
	goto L131
L134:
	;
	goto L120
L135:
	;
	if v429 != v390+int32(16) {
		goto L120
	} else {
		goto L138
	}
L136:
	;
	v421 = *(*int32)(unsafe.Add(mBase, uint32(v417)))
	v422 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v421)+5)))
	if v422&int32(3) == int32(0) {
		v429 = v417
		goto L135
	} else {
		goto L137
	}
L137:
	;
	F_reallymarkobject(m, l0, v421)
	mBase = m.M
	v428 = *(*int32)(unsafe.Add(mBase, uint32(v390)+8))
	v429 = v428
	goto L135
L138:
	;
	v434 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v390)+5)))
	v436 = v434 | int32(4)
	*(*uint8)(unsafe.Add(mBase, uint32(v390)+5)) = uint8(v436)
	goto L119
L139:
	;
	v528 = *(*int32)(unsafe.Add(mBase, uint32(l0)+180))
	if v528 == int32(0) {
		goto L162
	} else {
		goto L163
	}
L140:
	;
	v456 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v453)+5)))
	if v456&int32(3) == int32(0) {
		goto L139
	} else {
		goto L141
	}
L141:
	;
	v463 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v453)+5)))
	v465 = v453
	v466 = v463
	goto L150
L142:
	;
	goto L139
L143:
	;
	goto L142
L144:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = v465
	goto L143
L145:
	;
	v519 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v465)+68)) = v519
	goto L144
L146:
	;
	v517 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v465)+108)) = v517
	goto L144
L147:
	;
	v515 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v465)+32)) = v515
	goto L144
L148:
	;
	v513 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v465)+8)) = v513
	goto L144
L149:
	;
	v492 = *(*int32)(unsafe.Add(mBase, uint32(v465)+8))
	v493 = *(*int32)(unsafe.Add(mBase, uint32(v492)+8))
	if v493 < int32(4) {
		v504 = v492
		goto L158
	} else {
		goto L159
	}
L150:
	;
	v469 = v466 & int32(252)
	*(*uint8)(unsafe.Add(mBase, uint32(v465)+5)) = uint8(v469)
	v471 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v465)+4)))
	if v471 == int32(7) {
		goto L152
	} else {
		goto L153
	}
L152:
	;
	v477 = v469 | int32(4)
	*(*uint8)(unsafe.Add(mBase, uint32(v465)+5)) = uint8(v477)
	v479 = *(*int32)(unsafe.Add(mBase, uint32(v465)+8))
	if v479 == int32(0) {
		goto L154
	} else {
		goto L155
	}
L153:
	;
	switch v471 + int32(-5) {
	case 0:
		goto L147
	case 1:
		goto L148
	default:
		goto L143
	case 3:
		goto L146
	case 4:
		goto L145
	case 5:
		goto L149
	}
L154:
	;
	v488 = *(*int32)(unsafe.Add(mBase, uint32(v465)+12))
	v489 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v488)+5)))
	if v489&int32(3) != 0 {
		v465 = v488
		v466 = v489
		goto L150
	} else {
		goto L157
	}
L155:
	;
	v482 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v479)+5)))
	if v482&int32(3) == int32(0) {
		goto L154
	} else {
		goto L156
	}
L156:
	;
	F_reallymarkobject(m, l0, v479)
	mBase = m.M
	goto L154
L157:
	;
	goto L143
L158:
	;
	if v504 != v465+int32(16) {
		goto L143
	} else {
		goto L161
	}
L159:
	;
	v496 = *(*int32)(unsafe.Add(mBase, uint32(v492)))
	v497 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v496)+5)))
	if v497&int32(3) == int32(0) {
		v504 = v492
		goto L158
	} else {
		goto L160
	}
L160:
	;
	F_reallymarkobject(m, l0, v496)
	mBase = m.M
	v503 = *(*int32)(unsafe.Add(mBase, uint32(v465)+8))
	v504 = v503
	goto L158
L161:
	;
	v509 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v465)+5)))
	v511 = v509 | int32(4)
	*(*uint8)(unsafe.Add(mBase, uint32(v465)+5)) = uint8(v511)
	goto L142
L162:
	;
	v603 = *(*int32)(unsafe.Add(mBase, uint32(l0)+184))
	if v603 == int32(0) {
		goto L185
	} else {
		goto L186
	}
L163:
	;
	v531 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v528)+5)))
	if v531&int32(3) == int32(0) {
		goto L162
	} else {
		goto L164
	}
L164:
	;
	v538 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v528)+5)))
	v540 = v528
	v541 = v538
	goto L173
L165:
	;
	goto L162
L166:
	;
	goto L165
L167:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = v540
	goto L166
L168:
	;
	v594 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v540)+68)) = v594
	goto L167
L169:
	;
	v592 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v540)+108)) = v592
	goto L167
L170:
	;
	v590 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v540)+32)) = v590
	goto L167
L171:
	;
	v588 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v540)+8)) = v588
	goto L167
L172:
	;
	v567 = *(*int32)(unsafe.Add(mBase, uint32(v540)+8))
	v568 = *(*int32)(unsafe.Add(mBase, uint32(v567)+8))
	if v568 < int32(4) {
		v579 = v567
		goto L181
	} else {
		goto L182
	}
L173:
	;
	v544 = v541 & int32(252)
	*(*uint8)(unsafe.Add(mBase, uint32(v540)+5)) = uint8(v544)
	v546 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v540)+4)))
	if v546 == int32(7) {
		goto L175
	} else {
		goto L176
	}
L175:
	;
	v552 = v544 | int32(4)
	*(*uint8)(unsafe.Add(mBase, uint32(v540)+5)) = uint8(v552)
	v554 = *(*int32)(unsafe.Add(mBase, uint32(v540)+8))
	if v554 == int32(0) {
		goto L177
	} else {
		goto L178
	}
L176:
	;
	switch v546 + int32(-5) {
	case 0:
		goto L170
	case 1:
		goto L171
	default:
		goto L166
	case 3:
		goto L169
	case 4:
		goto L168
	case 5:
		goto L172
	}
L177:
	;
	v563 = *(*int32)(unsafe.Add(mBase, uint32(v540)+12))
	v564 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v563)+5)))
	if v564&int32(3) != 0 {
		v540 = v563
		v541 = v564
		goto L173
	} else {
		goto L180
	}
L178:
	;
	v557 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v554)+5)))
	if v557&int32(3) == int32(0) {
		goto L177
	} else {
		goto L179
	}
L179:
	;
	F_reallymarkobject(m, l0, v554)
	mBase = m.M
	goto L177
L180:
	;
	goto L166
L181:
	;
	if v579 != v540+int32(16) {
		goto L166
	} else {
		goto L184
	}
L182:
	;
	v571 = *(*int32)(unsafe.Add(mBase, uint32(v567)))
	v572 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v571)+5)))
	if v572&int32(3) == int32(0) {
		v579 = v567
		goto L181
	} else {
		goto L183
	}
L183:
	;
	F_reallymarkobject(m, l0, v571)
	mBase = m.M
	v578 = *(*int32)(unsafe.Add(mBase, uint32(v540)+8))
	v579 = v578
	goto L181
L184:
	;
	v584 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v540)+5)))
	v586 = v584 | int32(4)
	*(*uint8)(unsafe.Add(mBase, uint32(v540)+5)) = uint8(v586)
	goto L165
L185:
	;
	return
L186:
	;
	v606 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v603)+5)))
	if v606&int32(3) == int32(0) {
		goto L185
	} else {
		goto L187
	}
L187:
	;
	v613 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v603)+5)))
	v615 = v603
	v616 = v613
	goto L196
L188:
	;
	goto L185
L189:
	;
	goto L188
L190:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = v615
	goto L189
L191:
	;
	v669 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v615)+68)) = v669
	goto L190
L192:
	;
	v667 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v615)+108)) = v667
	goto L190
L193:
	;
	v665 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v615)+32)) = v665
	goto L190
L194:
	;
	v663 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v615)+8)) = v663
	goto L190
L195:
	;
	v642 = *(*int32)(unsafe.Add(mBase, uint32(v615)+8))
	v643 = *(*int32)(unsafe.Add(mBase, uint32(v642)+8))
	if v643 < int32(4) {
		v654 = v642
		goto L204
	} else {
		goto L205
	}
L196:
	;
	v619 = v616 & int32(252)
	*(*uint8)(unsafe.Add(mBase, uint32(v615)+5)) = uint8(v619)
	v621 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v615)+4)))
	if v621 == int32(7) {
		goto L198
	} else {
		goto L199
	}
L198:
	;
	v627 = v619 | int32(4)
	*(*uint8)(unsafe.Add(mBase, uint32(v615)+5)) = uint8(v627)
	v629 = *(*int32)(unsafe.Add(mBase, uint32(v615)+8))
	if v629 == int32(0) {
		goto L200
	} else {
		goto L201
	}
L199:
	;
	switch v621 + int32(-5) {
	case 0:
		goto L193
	case 1:
		goto L194
	default:
		goto L189
	case 3:
		goto L192
	case 4:
		goto L191
	case 5:
		goto L195
	}
L200:
	;
	v638 = *(*int32)(unsafe.Add(mBase, uint32(v615)+12))
	v639 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v638)+5)))
	if v639&int32(3) != 0 {
		v615 = v638
		v616 = v639
		goto L196
	} else {
		goto L203
	}
L201:
	;
	v632 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v629)+5)))
	if v632&int32(3) == int32(0) {
		goto L200
	} else {
		goto L202
	}
L202:
	;
	F_reallymarkobject(m, l0, v629)
	mBase = m.M
	goto L200
L203:
	;
	goto L189
L204:
	;
	if v654 != v615+int32(16) {
		goto L189
	} else {
		goto L207
	}
L205:
	;
	v646 = *(*int32)(unsafe.Add(mBase, uint32(v642)))
	v647 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v646)+5)))
	if v647&int32(3) == int32(0) {
		v654 = v642
		goto L204
	} else {
		goto L206
	}
L206:
	;
	F_reallymarkobject(m, l0, v646)
	mBase = m.M
	v653 = *(*int32)(unsafe.Add(mBase, uint32(v615)+8))
	v654 = v653
	goto L204
L207:
	;
	v659 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v615)+5)))
	v661 = v659 | int32(4)
	*(*uint8)(unsafe.Add(mBase, uint32(v615)+5)) = uint8(v661)
	goto L188
}
func F_match_class(m *base.Module, l0 int32, l1 int32) int32 {
	var v9 int32
	_ = v9
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v78 int32
	_ = v78
	var v85 int32
	_ = v85
	if base.Ui32(l1+int32(-65)) < base.Ui32(int32(26)) {
		v9 = l1 | int32(32)
	} else {
		v9 = l1
	}
	switch v9 + int32(-97) {
	case 0:
		v78 = base.B2i32(base.Ui32(l0|int32(32)+int32(-97)) < base.Ui32(int32(26)))
		if base.Ui32(l1+int32(-97)) < base.Ui32(int32(26)) {
			v85 = v78
		} else {
			v85 = base.B2i32(v78 == int32(0))
		}
		return v85
	default:
		return base.B2i32(l1 == l0)
	case 2:
		v78 = base.B2i32(base.Ui32(l0) < base.Ui32(int32(32))) | base.B2i32(l0 == int32(127))
		if base.Ui32(l1+int32(-97)) < base.Ui32(int32(26)) {
			v85 = v78
		} else {
			v85 = base.B2i32(v78 == int32(0))
		}
		return v85
	case 3:
		v78 = base.B2i32(base.Ui32(l0+int32(-48)) < base.Ui32(int32(10)))
		if base.Ui32(l1+int32(-97)) < base.Ui32(int32(26)) {
			v85 = v78
		} else {
			v85 = base.B2i32(v78 == int32(0))
		}
		return v85
	case 11:
		v78 = base.B2i32(base.Ui32(l0+int32(-97)) < base.Ui32(int32(26)))
		if base.Ui32(l1+int32(-97)) < base.Ui32(int32(26)) {
			v85 = v78
		} else {
			v85 = base.B2i32(v78 == int32(0))
		}
		return v85
	case 15:
		if base.Ui32(int32(93)) < base.Ui32(l0+int32(-33)) {
			v34 = int32(0)
		} else {
			v31 = F_isalnum(m, l0)
			v34 = base.B2i32(v31 == int32(0))
		}
		v78 = v34
		if base.Ui32(l1+int32(-97)) < base.Ui32(int32(26)) {
			v85 = v78
		} else {
			v85 = base.B2i32(v78 == int32(0))
		}
		return v85
	case 18:
		v78 = base.B2i32(l0 == int32(32)) | base.B2i32(base.Ui32(l0+int32(-9)) < base.Ui32(int32(5)))
		if base.Ui32(l1+int32(-97)) < base.Ui32(int32(26)) {
			v85 = v78
		} else {
			v85 = base.B2i32(v78 == int32(0))
		}
		return v85
	case 20:
		v78 = base.B2i32(base.Ui32(l0+int32(-65)) < base.Ui32(int32(26)))
		if base.Ui32(l1+int32(-97)) < base.Ui32(int32(26)) {
			v85 = v78
		} else {
			v85 = base.B2i32(v78 == int32(0))
		}
		return v85
	case 22:
		v78 = base.B2i32(base.Ui32(l0+int32(-48)) < base.Ui32(int32(10))) | base.B2i32(base.Ui32(l0|int32(32)+int32(-97)) < base.Ui32(int32(26)))
		if base.Ui32(l1+int32(-97)) < base.Ui32(int32(26)) {
			v85 = v78
		} else {
			v85 = base.B2i32(v78 == int32(0))
		}
		return v85
	case 23:
		v78 = base.B2i32(base.Ui32(l0+int32(-48)) < base.Ui32(int32(10))) | base.B2i32(base.Ui32(l0|int32(32)+int32(-97)) < base.Ui32(int32(6)))
		if base.Ui32(l1+int32(-97)) < base.Ui32(int32(26)) {
			v85 = v78
		} else {
			v85 = base.B2i32(v78 == int32(0))
		}
		return v85
	case 25:
		v78 = base.B2i32(l0 == int32(0))
		if base.Ui32(l1+int32(-97)) < base.Ui32(int32(26)) {
			v85 = v78
		} else {
			v85 = base.B2i32(v78 == int32(0))
		}
		return v85
	}
}
func F_maxn(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v17 float64
	_ = v17
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v24 float64
	_ = v24
	var v44 int32
	_ = v44
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v105 int32
	_ = v105
	var v108 int32
	_ = v108
	var v111 int32
	_ = v111
	var v115 float64
	_ = v115
	var v116 int32
	_ = v116
	var v118 float64
	_ = v118
	var v119 float64
	_ = v119
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v125 float64
	_ = v125
	var v128 int32
	_ = v128
	var v132 int32
	_ = v132
	F_luaL_checktype(m, l0, int32(1), int32(5))
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+8)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v11 + int32(16)
	goto L3
L3:
	;
	v17 = float64(0)
	v19 = F_lua_next(m, l0, int32(1))
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		goto L1
	} else {
		goto L5
	}
L4:
	;
	v128 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v128)+8)) = int32(3)
	*(*float64)(unsafe.Add(mBase, uint32(v128))) = v125
	v132 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v132 + int32(16)
	goto L40
L5:
	;
	if v19 == int32(0) {
		v125 = v17
		goto L4
	} else {
		goto L6
	}
L6:
	;
	v24 = v17
	goto L7
L7:
	;
	goto L11
L8:
	;
	v125 = v119
	goto L4
L9:
	;
	goto L20
L10:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v44 + int32(-16)
	goto L9
L11:
	;
	v44 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	goto L10
L17:
	;
	v122 = F_lua_next(m, l0, int32(1))
	mBase = m.M
	v123 = m.ExcPending
	if v123 != 0 {
		goto L1
	} else {
		goto L38
	}
L18:
	;
	if v111 != int32(3) {
		v119 = v24
		goto L17
	} else {
		goto L33
	}
L19:
	;
	v105 = m.G398
	if v71 != v105 {
		goto L31
	} else {
		goto L32
	}
L20:
	;
	goto L24
L24:
	;
	v68 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v71 = v68 + int32(-16)
	goto L19
L31:
	;
	v108 = *(*int32)(unsafe.Add(mBase, uint32(v71)+8))
	v111 = v108
	goto L18
L32:
	;
	v111 = int32(-1)
	goto L18
L33:
	;
	v115 = F_lua_tonumber(m, l0, int32(-1))
	mBase = m.M
	v116 = m.ExcPending
	if v116 != 0 {
		goto L1
	} else {
		goto L34
	}
L34:
	;
	if base.F64_gt(v115, v24) != 0 {
		goto L35
	} else {
		goto L36
	}
L35:
	;
	v118 = v115
	goto L37
L36:
	;
	v118 = v24
	goto L37
L37:
	;
	v119 = v118
	goto L17
L38:
	;
	if v122 != 0 {
		v24 = v119
		goto L7
	} else {
		goto L39
	}
L39:
	;
	goto L8
L40:
	;
	return int32(1)
}
func F_maybeConvertIntset(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v27 int32
	_ = v27
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v3&int32(240) != int32(96) {
		F__serverAssert(m, int32(_a1515), int32(_a1511), int32(82))
		mBase = m.M
		v27 = m.ExcPending
		if v27 != 0 {
			return
		} else {
			F_abort(m)
			mBase = m.M
			base.Wasm_trap_unreachable()
			for {
			}
		}
	} else {
		v8 = F_objectGetVal(m, l0)
		mBase = m.M
		v9 = *(*int32)(unsafe.Add(mBase, uint32(v8)+4))
		v11 = *(*int32)(unsafe.Add(mBase, _consts[1078]))
		v12 = int32(1073741824)
		if base.Ui32(v11) < base.Ui32(v12) {
			v15 = v11
		} else {
			v15 = v12
		}
		if base.Ui32(v9) <= base.Ui32(v15) {
			return
		} else {
			v18 = F_setTypeSize(m, l0)
			mBase = m.M
			v19 = m.ExcPending
			if v19 != 0 {
				return
			} else {
				v21 = F_setTypeConvertAndExpand(m, l0, int32(2), v18, int32(1))
				mBase = m.M
				v22 = m.ExcPending
				if v22 != 0 {
					return
				} else {
					return
				}
			}
		}
	}
}
func F_mbrtowc(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v41 int32
	_ = v41
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
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
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v121 int32
	_ = v121
	var v128 int32
	_ = v128
	if l3 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v10 = l3
	goto L3
L2:
	;
	v10 = int32(9128420)
	goto L3
L3:
	;
	v11 = *(*int32)(unsafe.Add(mBase, uint32(v10)))
	if l1 != 0 {
		goto L7
	} else {
		goto L8
	}
L4:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10))) = v128
	return int32(-2)
L5:
	;
	return v121
L6:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10))) = int32(0)
	goto L33
L7:
	;
	if l2 == int32(0) {
		v121 = int32(-2)
		goto L5
	} else {
		goto L10
	}
L8:
	;
	if v11 != 0 {
		goto L6
	} else {
		goto L9
	}
L9:
	;
	return int32(0)
L10:
	;
	if v11 == int32(0) {
		goto L12
	} else {
		goto L13
	}
L11:
	;
	v58 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v55))))
	v60 = int32(base.Ui32(v58) >> (uint(int32(3)) % 32))
	if base.Ui32(int32(7)) < base.Ui32(v60+int32(-16)|(v56>>(uint(int32(26))%32)+v60)) {
		goto L6
	} else {
		goto L24
	}
L12:
	;
	v19 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	v20 = base.I32_extend8_s(v19)
	if v20 < int32(0) {
		goto L14
	} else {
		goto L15
	}
L13:
	;
	v55 = l1
	v56 = v11
	v57 = l2
	goto L11
L14:
	;
	goto L19
L15:
	;
	if l0 == int32(0) {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	return base.B2i32(v20 != int32(0))
L17:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v19
	goto L16
L18:
	;
	v41 = v19 + int32(-194)
	if base.Ui32(int32(50)) < base.Ui32(v41) {
		goto L6
	} else {
		goto L22
	}
L19:
	;
	v30 = *(*int32)(unsafe.Add(mBase, _consts[1237]))
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v30)))
	if v31 != 0 {
		goto L18
	} else {
		goto L20
	}
L20:
	;
	if l0 == int32(0) {
		v121 = int32(1)
		goto L5
	} else {
		goto L21
	}
L21:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v20 & int32(57343)
	return int32(1)
L22:
	;
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v41<<(uint(int32(2))%32))+uint32(_consts[1236])))
	v50 = l2 + int32(-1)
	if v50 == int32(0) {
		v128 = v48
		goto L4
	} else {
		goto L23
	}
L23:
	;
	v55 = l1 + int32(1)
	v56 = v48
	v57 = v50
	goto L11
L24:
	;
	v70 = v55
	v72 = v56
	v74 = v57
	v75 = v58
	goto L25
L25:
	;
	v78 = v74 + int32(-1)
	v85 = v75&int32(255) + int32(-128) | v72<<(uint(int32(6))%32)
	if v85 < int32(0) {
		goto L27
	} else {
		goto L28
	}
L26:
	;
	goto L6
L27:
	;
	if v78 == int32(0) {
		v128 = v85
		goto L4
	} else {
		goto L31
	}
L28:
	;
	v88 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v10))) = v88
	if l0 == v88 {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	return l2 - v78
L30:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v85
	goto L29
L31:
	;
	v98 = v70 + int32(1)
	v99 = int32(*(*int8)(unsafe.Add(mBase, uint32(v98))))
	if v99 < int32(-64) {
		v70 = v98
		v72 = v85
		v74 = v78
		v75 = v99
		goto L25
	} else {
		goto L32
	}
L32:
	;
	goto L26
L33:
	;
	*(*int32)(unsafe.Add(mBase, _consts[9])) = int32(25)
	v121 = int32(-1)
	goto L5
}
func F_mbtowc(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v32 int32
	_ = v32
	var v39 int32
	_ = v39
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v64 int32
	_ = v64
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v110 int32
	_ = v110
	if l1 != 0 {
		if l2 == int32(0) {
			*(*int32)(unsafe.Add(mBase, _consts[9])) = int32(25)
			v110 = int32(-1)
			return v110
		} else {
			v10 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
			v11 = base.I32_extend8_s(v10)
			if v11 < int32(0) {
				v21 = *(*int32)(unsafe.Add(mBase, _consts[1237]))
				v22 = *(*int32)(unsafe.Add(mBase, uint32(v21)))
				if v22 != 0 {
					v32 = v10 + int32(-194)
					if base.Ui32(int32(50)) < base.Ui32(v32) {
						*(*int32)(unsafe.Add(mBase, _consts[9])) = int32(25)
						v110 = int32(-1)
						return v110
					} else {
						v39 = *(*int32)(unsafe.Add(mBase, uint32(v32<<(uint(int32(2))%32))+uint32(_consts[1236])))
						if base.Ui32(int32(3)) < base.Ui32(l2) {
							v49 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+1)))
							v51 = int32(base.Ui32(v49) >> (uint(int32(3)) % 32))
							if base.Ui32(int32(7)) < base.Ui32(v51+int32(-16)|(v51+v39>>(uint(int32(26))%32))) {
								*(*int32)(unsafe.Add(mBase, _consts[9])) = int32(25)
								v110 = int32(-1)
								return v110
							} else {
								v64 = v49 + int32(-128) | v39<<(uint(int32(6))%32)
								if v64 < int32(0) {
									v73 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+2)))
									v75 = v73 + int32(-128)
									if base.Ui32(int32(63)) < base.Ui32(v75) {
										*(*int32)(unsafe.Add(mBase, _consts[9])) = int32(25)
										v110 = int32(-1)
										return v110
									} else {
										v79 = v64 << (uint(int32(6)) % 32)
										v80 = v75 | v79
										if v79 < int32(0) {
											v89 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+3)))
											v91 = v89 + int32(-128)
											if base.Ui32(int32(63)) < base.Ui32(v91) {
												*(*int32)(unsafe.Add(mBase, _consts[9])) = int32(25)
												v110 = int32(-1)
												return v110
											} else {
												if l0 == int32(0) {
													v110 = int32(4)
													return v110
												} else {
													*(*int32)(unsafe.Add(mBase, uint32(l0))) = v91 | v80<<(uint(int32(6))%32)
													return int32(4)
												}
											}
										} else {
											if l0 == int32(0) {
												v110 = int32(3)
												return v110
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(l0))) = v80
												return int32(3)
											}
										}
									}
								} else {
									if l0 == int32(0) {
										v110 = int32(2)
										return v110
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(l0))) = v64
										return int32(2)
									}
								}
							}
						} else {
							if v39<<(uint(l2*int32(6)+int32(-6))%32) < int32(0) {
								*(*int32)(unsafe.Add(mBase, _consts[9])) = int32(25)
								v110 = int32(-1)
								return v110
							} else {
								v49 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+1)))
								v51 = int32(base.Ui32(v49) >> (uint(int32(3)) % 32))
								if base.Ui32(int32(7)) < base.Ui32(v51+int32(-16)|(v51+v39>>(uint(int32(26))%32))) {
									*(*int32)(unsafe.Add(mBase, _consts[9])) = int32(25)
									v110 = int32(-1)
									return v110
								} else {
									v64 = v49 + int32(-128) | v39<<(uint(int32(6))%32)
									if v64 < int32(0) {
										v73 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+2)))
										v75 = v73 + int32(-128)
										if base.Ui32(int32(63)) < base.Ui32(v75) {
											*(*int32)(unsafe.Add(mBase, _consts[9])) = int32(25)
											v110 = int32(-1)
											return v110
										} else {
											v79 = v64 << (uint(int32(6)) % 32)
											v80 = v75 | v79
											if v79 < int32(0) {
												v89 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+3)))
												v91 = v89 + int32(-128)
												if base.Ui32(int32(63)) < base.Ui32(v91) {
													*(*int32)(unsafe.Add(mBase, _consts[9])) = int32(25)
													v110 = int32(-1)
													return v110
												} else {
													if l0 == int32(0) {
														v110 = int32(4)
														return v110
													} else {
														*(*int32)(unsafe.Add(mBase, uint32(l0))) = v91 | v80<<(uint(int32(6))%32)
														return int32(4)
													}
												}
											} else {
												if l0 == int32(0) {
													v110 = int32(3)
													return v110
												} else {
													*(*int32)(unsafe.Add(mBase, uint32(l0))) = v80
													return int32(3)
												}
											}
										}
									} else {
										if l0 == int32(0) {
											v110 = int32(2)
											return v110
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(l0))) = v64
											return int32(2)
										}
									}
								}
							}
						}
					}
				} else {
					if l0 == int32(0) {
						v110 = int32(1)
						return v110
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(l0))) = v11 & int32(57343)
						return int32(1)
					}
				}
			} else {
				if l0 == int32(0) {
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(l0))) = v10
				}
				return base.B2i32(v11 != int32(0))
			}
		}
	} else {
		return int32(0)
	}
}
func F_memchr(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v40 int32
	_ = v40
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v89 int32
	_ = v89
	var v95 int32
	_ = v95
	v4 = int32(0)
	v7 = base.B2i32(l2 != v4)
	if l0&int32(3) == v4 {
		v33 = l0
		v35 = l2
		v36 = v7
		goto L3
	} else {
		goto L4
	}
L1:
	;
	return int32(0)
L2:
	;
	v84 = v77
	v86 = v79
	goto L20
L3:
	;
	if v36 == int32(0) {
		goto L1
	} else {
		goto L11
	}
L4:
	;
	if l2 == int32(0) {
		v33 = l0
		v35 = l2
		v36 = v7
		goto L3
	} else {
		goto L5
	}
L5:
	;
	v16 = l0
	v18 = l2
	goto L6
L6:
	;
	v21 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16))))
	if v21 == l1&int32(255) {
		v77 = v16
		v79 = v18
		goto L2
	} else {
		goto L8
	}
L7:
	;
	v33 = v28
	v35 = v24
	v36 = v26
	goto L3
L8:
	;
	v24 = v18 + int32(-1)
	v25 = int32(0)
	v26 = base.B2i32(v24 != v25)
	v28 = v16 + int32(1)
	if v28&int32(3) == v25 {
		v33 = v28
		v35 = v24
		v36 = v26
		goto L3
	} else {
		goto L9
	}
L9:
	;
	if v24 != 0 {
		v16 = v28
		v18 = v24
		goto L6
	} else {
		goto L10
	}
L10:
	;
	goto L7
L11:
	;
	v40 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v33))))
	if v40 == l1&int32(255) {
		v70 = v33
		v72 = v35
		goto L12
	} else {
		goto L13
	}
L12:
	;
	if v72 == int32(0) {
		goto L1
	} else {
		goto L19
	}
L13:
	;
	if base.Ui32(v35) < base.Ui32(int32(4)) {
		v70 = v33
		v72 = v35
		goto L12
	} else {
		goto L14
	}
L14:
	;
	v50 = v33
	v52 = v35
	goto L15
L15:
	;
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v50)))
	v57 = v56 ^ l1&int32(255)*int32(16843009)
	v60 = int32(-2139062144)
	if (int32(16843008)-v57|v57)&v60 != v60 {
		v77 = v50
		v79 = v52
		goto L2
	} else {
		goto L17
	}
L16:
	;
	v70 = v65
	v72 = v67
	goto L12
L17:
	;
	v65 = v50 + int32(4)
	v67 = v52 + int32(-4)
	if base.Ui32(int32(3)) < base.Ui32(v67) {
		v50 = v65
		v52 = v67
		goto L15
	} else {
		goto L18
	}
L18:
	;
	goto L16
L19:
	;
	v77 = v70
	v79 = v72
	goto L2
L20:
	;
	v89 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v84))))
	if v89 != l1&int32(255) {
		goto L22
	} else {
		goto L23
	}
L21:
	;
	goto L1
L22:
	;
	v95 = v86 + int32(-1)
	if v95 != 0 {
		v84 = v84 + int32(1)
		v86 = v95
		goto L20
	} else {
		goto L24
	}
L23:
	;
	return v84
L24:
	;
	goto L21
}
func F_memmove(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v21 int32
	_ = v21
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v52 int32
	_ = v52
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v67 int32
	_ = v67
	var v74 int32
	_ = v74
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v87 int32
	_ = v87
	var v94 int32
	_ = v94
	var v98 int32
	_ = v98
	var v101 int32
	_ = v101
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v115 int32
	_ = v115
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v120 int32
	_ = v120
	var v122 int32
	_ = v122
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v137 int32
	_ = v137
	var v139 int32
	_ = v139
	var v144 int32
	_ = v144
	if l0 == l1 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return l0
L2:
	;
	v7 = l2 + l0
	if base.Ui32(int32(0)-l2<<(uint(int32(1))%32)) < base.Ui32(l1-v7) {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v21 = (l1 ^ l0) & int32(3)
	if base.Ui32(l1) <= base.Ui32(l0) {
		goto L10
	} else {
		goto L11
	}
L4:
	;
	if l2 == int32(0) {
		v17 = l0
		goto L6
	} else {
		goto L7
	}
L5:
	;
	return v17
L6:
	;
	goto L5
L7:
	;
	v16 = F__emscripten_memcpy_bulkmem(m, l0, l1, l2)
	mBase = m.M
	v17 = v16
	goto L6
L8:
	;
	if v127 == int32(0) {
		goto L1
	} else {
		goto L40
	}
L9:
	;
	if base.Ui32(v105) <= base.Ui32(int32(3)) {
		v126 = v104
		v127 = v105
		v128 = v106
		goto L8
	} else {
		goto L36
	}
L10:
	;
	if v21 != 0 {
		v87 = l2
		goto L20
	} else {
		goto L21
	}
L11:
	;
	if v21 == int32(0) {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	if l0&int32(3) != 0 {
		goto L14
	} else {
		goto L15
	}
L13:
	;
	v126 = l1
	v127 = l2
	v128 = l0
	goto L8
L14:
	;
	v28 = l1
	v29 = l2
	v30 = l0
	goto L16
L15:
	;
	v104 = l1
	v105 = l2
	v106 = l0
	goto L9
L16:
	;
	if v29 == int32(0) {
		goto L1
	} else {
		goto L18
	}
L18:
	;
	v34 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28))))
	*(*uint8)(unsafe.Add(mBase, uint32(v30))) = uint8(v34)
	v36 = int32(1)
	v37 = v28 + v36
	v39 = v29 + int32(-1)
	v41 = v30 + v36
	if v41&int32(3) == int32(0) {
		v104 = v37
		v105 = v39
		v106 = v41
		goto L9
	} else {
		goto L19
	}
L19:
	;
	v28 = v37
	v29 = v39
	v30 = v41
	goto L16
L20:
	;
	if v87 == int32(0) {
		goto L1
	} else {
		goto L32
	}
L21:
	;
	if v7&int32(3) == int32(0) {
		v67 = l2
		goto L22
	} else {
		goto L23
	}
L22:
	;
	if base.Ui32(v67) <= base.Ui32(int32(3)) {
		v87 = v67
		goto L20
	} else {
		goto L28
	}
L23:
	;
	v52 = l2
	goto L24
L24:
	;
	if v52 == int32(0) {
		goto L1
	} else {
		goto L26
	}
L25:
	;
	v67 = v58
	goto L22
L26:
	;
	v58 = v52 + int32(-1)
	v59 = l0 + v58
	v61 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1+v58))))
	*(*uint8)(unsafe.Add(mBase, uint32(v59))) = uint8(v61)
	if v59&int32(3) != 0 {
		v52 = v58
		goto L24
	} else {
		goto L27
	}
L27:
	;
	goto L25
L28:
	;
	v74 = v67
	goto L29
L29:
	;
	v78 = v74 + int32(-4)
	v81 = *(*int32)(unsafe.Add(mBase, uint32(l1+v78)))
	*(*int32)(unsafe.Add(mBase, uint32(l0+v78))) = v81
	if base.Ui32(int32(3)) < base.Ui32(v78) {
		v74 = v78
		goto L29
	} else {
		goto L31
	}
L30:
	;
	v87 = v78
	goto L20
L31:
	;
	goto L30
L32:
	;
	v94 = v87
	goto L33
L33:
	;
	v98 = v94 + int32(-1)
	v101 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1+v98))))
	*(*uint8)(unsafe.Add(mBase, uint32(l0+v98))) = uint8(v101)
	if v98 != 0 {
		v94 = v98
		goto L33
	} else {
		goto L35
	}
L35:
	;
	goto L1
L36:
	;
	v111 = v104
	v112 = v105
	v113 = v106
	goto L37
L37:
	;
	v115 = *(*int32)(unsafe.Add(mBase, uint32(v111)))
	*(*int32)(unsafe.Add(mBase, uint32(v113))) = v115
	v117 = int32(4)
	v118 = v111 + v117
	v120 = v113 + v117
	v122 = v112 + int32(-4)
	if base.Ui32(int32(3)) < base.Ui32(v122) {
		v111 = v118
		v112 = v122
		v113 = v120
		goto L37
	} else {
		goto L39
	}
L38:
	;
	v126 = v118
	v127 = v122
	v128 = v120
	goto L8
L39:
	;
	goto L38
L40:
	;
	v133 = v126
	v134 = v127
	v135 = v128
	goto L41
L41:
	;
	v137 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v133))))
	*(*uint8)(unsafe.Add(mBase, uint32(v135))) = uint8(v137)
	v139 = int32(1)
	v144 = v134 + int32(-1)
	if v144 != 0 {
		v133 = v133 + v139
		v134 = v144
		v135 = v135 + v139
		goto L41
	} else {
		goto L43
	}
L42:
	;
	goto L1
L43:
	;
	goto L42
}
func F_memtoull(m *base.Module, l0 int32, l1 int32) int64 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v19 int32
	_ = v19
	var v23 int32
	_ = v23
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v52 int64
	_ = v52
	var v57 int32
	_ = v57
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
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v98 int32
	_ = v98
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	var v111 int32
	_ = v111
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
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
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v147 int32
	_ = v147
	var v151 int32
	_ = v151
	var v153 int32
	_ = v153
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
	var v160 int32
	_ = v160
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v178 int32
	_ = v178
	var v181 int32
	_ = v181
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v187 int32
	_ = v187
	var v191 int32
	_ = v191
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v218 int32
	_ = v218
	var v221 int32
	_ = v221
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v227 int32
	_ = v227
	var v231 int32
	_ = v231
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v255 int32
	_ = v255
	var v258 int32
	_ = v258
	var v261 int32
	_ = v261
	var v263 int32
	_ = v263
	var v264 int32
	_ = v264
	var v265 int32
	_ = v265
	var v267 int32
	_ = v267
	var v271 int32
	_ = v271
	var v273 int32
	_ = v273
	var v274 int32
	_ = v274
	var v275 int32
	_ = v275
	var v277 int32
	_ = v277
	var v278 int32
	_ = v278
	var v279 int32
	_ = v279
	var v280 int32
	_ = v280
	var v285 int32
	_ = v285
	var v286 int32
	_ = v286
	var v289 int32
	_ = v289
	var v290 int32
	_ = v290
	var v293 int32
	_ = v293
	var v294 int32
	_ = v294
	var v295 int32
	_ = v295
	var v298 int32
	_ = v298
	var v301 int32
	_ = v301
	var v303 int32
	_ = v303
	var v304 int32
	_ = v304
	var v305 int32
	_ = v305
	var v307 int32
	_ = v307
	var v311 int32
	_ = v311
	var v313 int32
	_ = v313
	var v314 int32
	_ = v314
	var v315 int32
	_ = v315
	var v317 int32
	_ = v317
	var v318 int32
	_ = v318
	var v319 int32
	_ = v319
	var v320 int32
	_ = v320
	var v325 int32
	_ = v325
	var v326 int32
	_ = v326
	var v329 int32
	_ = v329
	var v330 int32
	_ = v330
	var v333 int32
	_ = v333
	var v334 int32
	_ = v334
	var v335 int32
	_ = v335
	var v338 int64
	_ = v338
	var v346 int64
	_ = v346
	var v347 int32
	_ = v347
	var v350 int64
	_ = v350
	var v359 int32
	_ = v359
	var v364 int32
	_ = v364
	var v366 int32
	_ = v366
	var v375 int64
	_ = v375
	var v376 int32
	_ = v376
	var v383 int32
	_ = v383
	var v384 int32
	_ = v384
	var v387 int64
	_ = v387
	var v396 int64
	_ = v396
	v9 = m.G0
	v11 = v9 - int32(144)
	m.G0 = v11
	if l1 != 0 {
		goto L5
	} else {
		goto L6
	}
L1:
	;
	m.G0 = v11 + int32(144)
	return v396
L2:
	;
	v347 = v49 - l0
	if base.Ui32(v347) < base.Ui32(int32(128)) {
		goto L113
	} else {
		goto L114
	}
L3:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = int32(1)
	v396 = int64(0)
	goto L1
L4:
	;
	if base.Ui32(int32(10)) <= base.Ui32(base.I32_extend8_s(v23)+int32(-48)) {
		v48 = v23
		v49 = l0
		goto L9
	} else {
		goto L10
	}
L5:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = int32(0)
	v19 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v19 == int32(45) {
		goto L3
	} else {
		goto L8
	}
L6:
	;
	v14 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v14 != int32(45) {
		v23 = v14
		goto L4
	} else {
		goto L7
	}
L7:
	;
	v396 = int64(0)
	goto L1
L8:
	;
	v23 = v19
	goto L4
L9:
	;
	v52 = int64(1)
	if v48&int32(255) == int32(0) {
		v346 = v52
		goto L2
	} else {
		goto L14
	}
L10:
	;
	v35 = l0
	goto L11
L11:
	;
	v37 = int32(*(*int8)(unsafe.Add(mBase, uint32(v35)+1)))
	v39 = v35 + int32(1)
	if base.Ui32(v37+int32(-48)) < base.Ui32(int32(10)) {
		v35 = v39
		goto L11
	} else {
		goto L13
	}
L12:
	;
	v48 = v37
	v49 = v39
	goto L9
L13:
	;
	goto L12
L14:
	;
	v57 = int32(_a1638)
	v60 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v49))))
	if v60 != 0 {
		goto L17
	} else {
		goto L18
	}
L15:
	;
	if v92-v94 == int32(0) {
		v346 = v52
		goto L2
	} else {
		goto L27
	}
L16:
	;
	v92 = F_tolower(m, v88)
	mBase = m.M
	v93 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v89))))
	v94 = F_tolower(m, v93)
	mBase = m.M
	goto L15
L17:
	;
	v62 = v49
	v63 = v57
	v64 = v60
	goto L20
L18:
	;
	v88 = int32(0)
	v89 = v57
	goto L16
L19:
	;
	v88 = v85 & int32(255)
	v89 = v84
	goto L16
L20:
	;
	v66 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v63))))
	if v66 == int32(0) {
		v84 = v63
		v85 = v64
		goto L19
	} else {
		goto L22
	}
L21:
	;
	v84 = v78
	v85 = int32(0)
	goto L19
L22:
	;
	v70 = v64 & int32(255)
	if v70 == v66 {
		goto L23
	} else {
		goto L24
	}
L23:
	;
	v77 = int32(1)
	v78 = v63 + v77
	v79 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v62)+1)))
	if v79 != 0 {
		v62 = v62 + v77
		v63 = v78
		v64 = v79
		goto L20
	} else {
		goto L26
	}
L24:
	;
	v72 = F_tolower(m, v70)
	mBase = m.M
	v73 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v63))))
	v74 = F_tolower(m, v73)
	mBase = m.M
	if v72 == v74 {
		goto L23
	} else {
		goto L25
	}
L25:
	;
	v76 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v62))))
	v84 = v63
	v85 = v76
	goto L19
L26:
	;
	goto L21
L27:
	;
	v98 = int32(_a1639)
	v101 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v49))))
	if v101 != 0 {
		goto L31
	} else {
		goto L32
	}
L28:
	;
	v138 = int32(_a1640)
	v141 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v49))))
	if v141 != 0 {
		goto L45
	} else {
		goto L46
	}
L29:
	;
	if v133-v135 != 0 {
		goto L28
	} else {
		goto L41
	}
L30:
	;
	v133 = F_tolower(m, v129)
	mBase = m.M
	v134 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v130))))
	v135 = F_tolower(m, v134)
	mBase = m.M
	goto L29
L31:
	;
	v103 = v49
	v104 = v98
	v105 = v101
	goto L34
L32:
	;
	v129 = int32(0)
	v130 = v98
	goto L30
L33:
	;
	v129 = v126 & int32(255)
	v130 = v125
	goto L30
L34:
	;
	v107 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v104))))
	if v107 == int32(0) {
		v125 = v104
		v126 = v105
		goto L33
	} else {
		goto L36
	}
L35:
	;
	v125 = v119
	v126 = int32(0)
	goto L33
L36:
	;
	v111 = v105 & int32(255)
	if v111 == v107 {
		goto L37
	} else {
		goto L38
	}
L37:
	;
	v118 = int32(1)
	v119 = v104 + v118
	v120 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v103)+1)))
	if v120 != 0 {
		v103 = v103 + v118
		v104 = v119
		v105 = v120
		goto L34
	} else {
		goto L40
	}
L38:
	;
	v113 = F_tolower(m, v111)
	mBase = m.M
	v114 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v104))))
	v115 = F_tolower(m, v114)
	mBase = m.M
	if v113 == v115 {
		goto L37
	} else {
		goto L39
	}
L39:
	;
	v117 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v103))))
	v125 = v104
	v126 = v117
	goto L33
L40:
	;
	goto L35
L41:
	;
	v346 = int64(1000)
	goto L2
L42:
	;
	v178 = int32(_a820)
	v181 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v49))))
	if v181 != 0 {
		goto L59
	} else {
		goto L60
	}
L43:
	;
	if v173-v175 != 0 {
		goto L42
	} else {
		goto L55
	}
L44:
	;
	v173 = F_tolower(m, v169)
	mBase = m.M
	v174 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v170))))
	v175 = F_tolower(m, v174)
	mBase = m.M
	goto L43
L45:
	;
	v143 = v49
	v144 = v138
	v145 = v141
	goto L48
L46:
	;
	v169 = int32(0)
	v170 = v138
	goto L44
L47:
	;
	v169 = v166 & int32(255)
	v170 = v165
	goto L44
L48:
	;
	v147 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v144))))
	if v147 == int32(0) {
		v165 = v144
		v166 = v145
		goto L47
	} else {
		goto L50
	}
L49:
	;
	v165 = v159
	v166 = int32(0)
	goto L47
L50:
	;
	v151 = v145 & int32(255)
	if v151 == v147 {
		goto L51
	} else {
		goto L52
	}
L51:
	;
	v158 = int32(1)
	v159 = v144 + v158
	v160 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v143)+1)))
	if v160 != 0 {
		v143 = v143 + v158
		v144 = v159
		v145 = v160
		goto L48
	} else {
		goto L54
	}
L52:
	;
	v153 = F_tolower(m, v151)
	mBase = m.M
	v154 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v144))))
	v155 = F_tolower(m, v154)
	mBase = m.M
	if v153 == v155 {
		goto L51
	} else {
		goto L53
	}
L53:
	;
	v157 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v143))))
	v165 = v144
	v166 = v157
	goto L47
L54:
	;
	goto L49
L55:
	;
	v346 = int64(1024)
	goto L2
L56:
	;
	v218 = int32(_a1641)
	v221 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v49))))
	if v221 != 0 {
		goto L73
	} else {
		goto L74
	}
L57:
	;
	if v213-v215 != 0 {
		goto L56
	} else {
		goto L69
	}
L58:
	;
	v213 = F_tolower(m, v209)
	mBase = m.M
	v214 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v210))))
	v215 = F_tolower(m, v214)
	mBase = m.M
	goto L57
L59:
	;
	v183 = v49
	v184 = v178
	v185 = v181
	goto L62
L60:
	;
	v209 = int32(0)
	v210 = v178
	goto L58
L61:
	;
	v209 = v206 & int32(255)
	v210 = v205
	goto L58
L62:
	;
	v187 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v184))))
	if v187 == int32(0) {
		v205 = v184
		v206 = v185
		goto L61
	} else {
		goto L64
	}
L63:
	;
	v205 = v199
	v206 = int32(0)
	goto L61
L64:
	;
	v191 = v185 & int32(255)
	if v191 == v187 {
		goto L65
	} else {
		goto L66
	}
L65:
	;
	v198 = int32(1)
	v199 = v184 + v198
	v200 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v183)+1)))
	if v200 != 0 {
		v183 = v183 + v198
		v184 = v199
		v185 = v200
		goto L62
	} else {
		goto L68
	}
L66:
	;
	v193 = F_tolower(m, v191)
	mBase = m.M
	v194 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v184))))
	v195 = F_tolower(m, v194)
	mBase = m.M
	if v193 == v195 {
		goto L65
	} else {
		goto L67
	}
L67:
	;
	v197 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v183))))
	v205 = v184
	v206 = v197
	goto L61
L68:
	;
	goto L63
L69:
	;
	v346 = int64(1000000)
	goto L2
L70:
	;
	v258 = int32(_a827)
	v261 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v49))))
	if v261 != 0 {
		goto L87
	} else {
		goto L88
	}
L71:
	;
	if v253-v255 != 0 {
		goto L70
	} else {
		goto L83
	}
L72:
	;
	v253 = F_tolower(m, v249)
	mBase = m.M
	v254 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v250))))
	v255 = F_tolower(m, v254)
	mBase = m.M
	goto L71
L73:
	;
	v223 = v49
	v224 = v218
	v225 = v221
	goto L76
L74:
	;
	v249 = int32(0)
	v250 = v218
	goto L72
L75:
	;
	v249 = v246 & int32(255)
	v250 = v245
	goto L72
L76:
	;
	v227 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v224))))
	if v227 == int32(0) {
		v245 = v224
		v246 = v225
		goto L75
	} else {
		goto L78
	}
L77:
	;
	v245 = v239
	v246 = int32(0)
	goto L75
L78:
	;
	v231 = v225 & int32(255)
	if v231 == v227 {
		goto L79
	} else {
		goto L80
	}
L79:
	;
	v238 = int32(1)
	v239 = v224 + v238
	v240 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v223)+1)))
	if v240 != 0 {
		v223 = v223 + v238
		v224 = v239
		v225 = v240
		goto L76
	} else {
		goto L82
	}
L80:
	;
	v233 = F_tolower(m, v231)
	mBase = m.M
	v234 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v224))))
	v235 = F_tolower(m, v234)
	mBase = m.M
	if v233 == v235 {
		goto L79
	} else {
		goto L81
	}
L81:
	;
	v237 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v223))))
	v245 = v224
	v246 = v237
	goto L75
L82:
	;
	goto L77
L83:
	;
	v346 = int64(1048576)
	goto L2
L84:
	;
	v298 = int32(_a1642)
	v301 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v49))))
	if v301 != 0 {
		goto L101
	} else {
		goto L102
	}
L85:
	;
	if v293-v295 != 0 {
		goto L84
	} else {
		goto L97
	}
L86:
	;
	v293 = F_tolower(m, v289)
	mBase = m.M
	v294 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v290))))
	v295 = F_tolower(m, v294)
	mBase = m.M
	goto L85
L87:
	;
	v263 = v49
	v264 = v258
	v265 = v261
	goto L90
L88:
	;
	v289 = int32(0)
	v290 = v258
	goto L86
L89:
	;
	v289 = v286 & int32(255)
	v290 = v285
	goto L86
L90:
	;
	v267 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v264))))
	if v267 == int32(0) {
		v285 = v264
		v286 = v265
		goto L89
	} else {
		goto L92
	}
L91:
	;
	v285 = v279
	v286 = int32(0)
	goto L89
L92:
	;
	v271 = v265 & int32(255)
	if v271 == v267 {
		goto L93
	} else {
		goto L94
	}
L93:
	;
	v278 = int32(1)
	v279 = v264 + v278
	v280 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v263)+1)))
	if v280 != 0 {
		v263 = v263 + v278
		v264 = v279
		v265 = v280
		goto L90
	} else {
		goto L96
	}
L94:
	;
	v273 = F_tolower(m, v271)
	mBase = m.M
	v274 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v264))))
	v275 = F_tolower(m, v274)
	mBase = m.M
	if v273 == v275 {
		goto L93
	} else {
		goto L95
	}
L95:
	;
	v277 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v263))))
	v285 = v264
	v286 = v277
	goto L89
L96:
	;
	goto L91
L97:
	;
	v346 = int64(1000000000)
	goto L2
L98:
	;
	v338 = int64(0)
	if l1 == int32(0) {
		v396 = v338
		goto L1
	} else {
		goto L112
	}
L99:
	;
	if v333-v335 != 0 {
		goto L98
	} else {
		goto L111
	}
L100:
	;
	v333 = F_tolower(m, v329)
	mBase = m.M
	v334 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v330))))
	v335 = F_tolower(m, v334)
	mBase = m.M
	goto L99
L101:
	;
	v303 = v49
	v304 = v298
	v305 = v301
	goto L104
L102:
	;
	v329 = int32(0)
	v330 = v298
	goto L100
L103:
	;
	v329 = v326 & int32(255)
	v330 = v325
	goto L100
L104:
	;
	v307 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v304))))
	if v307 == int32(0) {
		v325 = v304
		v326 = v305
		goto L103
	} else {
		goto L106
	}
L105:
	;
	v325 = v319
	v326 = int32(0)
	goto L103
L106:
	;
	v311 = v305 & int32(255)
	if v311 == v307 {
		goto L107
	} else {
		goto L108
	}
L107:
	;
	v318 = int32(1)
	v319 = v304 + v318
	v320 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v303)+1)))
	if v320 != 0 {
		v303 = v303 + v318
		v304 = v319
		v305 = v320
		goto L104
	} else {
		goto L110
	}
L108:
	;
	v313 = F_tolower(m, v311)
	mBase = m.M
	v314 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v304))))
	v315 = F_tolower(m, v314)
	mBase = m.M
	if v313 == v315 {
		goto L107
	} else {
		goto L109
	}
L109:
	;
	v317 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v303))))
	v325 = v304
	v326 = v317
	goto L103
L110:
	;
	goto L105
L111:
	;
	v346 = int64(1073741824)
	goto L2
L112:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = int32(1)
	v396 = v338
	goto L1
L113:
	;
	if v347 == int32(0) {
		goto L117
	} else {
		goto L118
	}
L114:
	;
	v350 = int64(0)
	if l1 == int32(0) {
		v396 = v350
		goto L1
	} else {
		goto L115
	}
L115:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = int32(1)
	v396 = v350
	goto L1
L116:
	;
	v364 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v11+int32(16)+v347))) = uint8(v364)
	v366 = int32(9116376)
	goto L119
L117:
	;
	goto L116
L118:
	;
	v359 = F__emscripten_memcpy_bulkmem(m, v11+int32(16), l0, v347)
	mBase = m.M
	goto L117
L119:
	;
	*(*int32)(unsafe.Add(mBase, _consts[9])) = int32(0)
	v375 = F_strtox_2(m, v11+int32(16), v11+int32(12), int32(10), int64(-1))
	mBase = m.M
	goto L120
L120:
	;
	v376 = *(*int32)(unsafe.Add(mBase, _consts[9]))
	if v376 == int32(68) {
		goto L122
	} else {
		goto L123
	}
L121:
	;
	v396 = v375 * v346
	goto L1
L122:
	;
	v387 = int64(0)
	if l1 == int32(0) {
		v396 = v387
		goto L1
	} else {
		goto L128
	}
L123:
	;
	if v375 != int64(0) {
		goto L124
	} else {
		goto L125
	}
L124:
	;
	v383 = *(*int32)(unsafe.Add(mBase, uint32(v11)+12))
	v384 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v383))))
	if v384 == int32(0) {
		goto L121
	} else {
		goto L127
	}
L125:
	;
	if v376 == int32(28) {
		goto L122
	} else {
		goto L126
	}
L126:
	;
	goto L124
L127:
	;
	goto L122
L128:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = int32(1)
	v396 = v387
	goto L1
}
func F_migrateCommand(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	var v53 int32
	_ = v53
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
	var v116 int32
	_ = v116
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v143 int32
	_ = v143
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v150 int32
	_ = v150
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v156 int32
	_ = v156
	var v160 int32
	_ = v160
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
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v197 int32
	_ = v197
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v203 int32
	_ = v203
	var v207 int32
	_ = v207
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v234 int32
	_ = v234
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v244 int32
	_ = v244
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v248 int32
	_ = v248
	var v250 int32
	_ = v250
	var v254 int32
	_ = v254
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v258 int32
	_ = v258
	var v260 int32
	_ = v260
	var v261 int32
	_ = v261
	var v262 int32
	_ = v262
	var v263 int32
	_ = v263
	var v268 int32
	_ = v268
	var v269 int32
	_ = v269
	var v272 int32
	_ = v272
	var v273 int32
	_ = v273
	var v276 int32
	_ = v276
	var v277 int32
	_ = v277
	var v278 int32
	_ = v278
	var v283 int32
	_ = v283
	var v285 int32
	_ = v285
	var v286 int32
	_ = v286
	var v288 int32
	_ = v288
	var v292 int32
	_ = v292
	var v293 int32
	_ = v293
	var v295 int32
	_ = v295
	var v296 int32
	_ = v296
	var v297 int32
	_ = v297
	var v298 int32
	_ = v298
	var v302 int32
	_ = v302
	var v306 int32
	_ = v306
	var v307 int32
	_ = v307
	var v308 int32
	_ = v308
	var v311 int32
	_ = v311
	var v313 int32
	_ = v313
	var v314 int32
	_ = v314
	var v315 int32
	_ = v315
	var v317 int32
	_ = v317
	var v321 int32
	_ = v321
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
	var v329 int32
	_ = v329
	var v330 int32
	_ = v330
	var v335 int32
	_ = v335
	var v336 int32
	_ = v336
	var v339 int32
	_ = v339
	var v340 int32
	_ = v340
	var v343 int32
	_ = v343
	var v344 int32
	_ = v344
	var v345 int32
	_ = v345
	var v347 int32
	_ = v347
	var v348 int32
	_ = v348
	var v349 int32
	_ = v349
	var v352 int32
	_ = v352
	var v359 int32
	_ = v359
	var v362 int32
	_ = v362
	var v365 int32
	_ = v365
	var v368 int32
	_ = v368
	var v369 int32
	_ = v369
	var v374 int32
	_ = v374
	var v376 int32
	_ = v376
	var v378 int32
	_ = v378
	var v382 int32
	_ = v382
	var v384 int32
	_ = v384
	var v386 int32
	_ = v386
	var v390 int32
	_ = v390
	var v391 int32
	_ = v391
	var v392 int32
	_ = v392
	var v393 int32
	_ = v393
	var v394 int32
	_ = v394
	var v396 int32
	_ = v396
	var v399 int32
	_ = v399
	var v400 int32
	_ = v400
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
	var v412 int32
	_ = v412
	var v414 int32
	_ = v414
	var v415 int32
	_ = v415
	var v416 int32
	_ = v416
	var v417 int32
	_ = v417
	var v418 int32
	_ = v418
	var v449 int32
	_ = v449
	var v450 int32
	_ = v450
	var v454 int32
	_ = v454
	var v455 int32
	_ = v455
	var v456 int32
	_ = v456
	var v457 int32
	_ = v457
	var v461 int32
	_ = v461
	var v462 int32
	_ = v462
	var v463 int32
	_ = v463
	var v468 int32
	_ = v468
	var v471 int32
	_ = v471
	var v472 int32
	_ = v472
	var v473 int32
	_ = v473
	var v475 int32
	_ = v475
	var v476 int32
	_ = v476
	var v489 int32
	_ = v489
	var v493 int32
	_ = v493
	var v519 int32
	_ = v519
	var v520 int32
	_ = v520
	var v522 int32
	_ = v522
	var v523 int32
	_ = v523
	var v526 int32
	_ = v526
	var v528 int32
	_ = v528
	var v529 int32
	_ = v529
	var v530 int32
	_ = v530
	var v535 int32
	_ = v535
	var v537 int32
	_ = v537
	var v541 int32
	_ = v541
	var v543 int32
	_ = v543
	var v549 int32
	_ = v549
	var v552 int32
	_ = v552
	var v553 int32
	_ = v553
	var v555 int32
	_ = v555
	var v557 int32
	_ = v557
	var v559 int32
	_ = v559
	var v561 int32
	_ = v561
	var v588 int32
	_ = v588
	var v602 int32
	_ = v602
	var v603 int32
	_ = v603
	var v614 int32
	_ = v614
	var v615 int32
	_ = v615
	var v616 int32
	_ = v616
	var v617 int32
	_ = v617
	var v618 int32
	_ = v618
	var v619 int32
	_ = v619
	var v621 int32
	_ = v621
	var v623 int32
	_ = v623
	var v626 int32
	_ = v626
	var v627 int32
	_ = v627
	var v630 int32
	_ = v630
	var v639 int32
	_ = v639
	var v640 int32
	_ = v640
	var v647 int32
	_ = v647
	var v648 int32
	_ = v648
	var v654 int32
	_ = v654
	var v659 int32
	_ = v659
	var v660 int32
	_ = v660
	var v661 int32
	_ = v661
	var v662 int32
	_ = v662
	var v663 int32
	_ = v663
	var v666 int32
	_ = v666
	var v667 int32
	_ = v667
	var v673 int32
	_ = v673
	var v678 int32
	_ = v678
	var v679 int32
	_ = v679
	var v680 int32
	_ = v680
	var v681 int32
	_ = v681
	var v682 int32
	_ = v682
	var v685 int32
	_ = v685
	var v686 int32
	_ = v686
	var v691 int32
	_ = v691
	var v692 int32
	_ = v692
	var v693 int32
	_ = v693
	var v698 int32
	_ = v698
	var v699 int32
	_ = v699
	var v706 int32
	_ = v706
	var v707 int32
	_ = v707
	var v712 int64
	_ = v712
	var v713 int32
	_ = v713
	var v714 int32
	_ = v714
	var v717 int32
	_ = v717
	var v725 int32
	_ = v725
	var v732 int32
	_ = v732
	var v762 int32
	_ = v762
	var v763 int32
	_ = v763
	var v764 int32
	_ = v764
	var v768 int32
	_ = v768
	var v778 int64
	_ = v778
	var v779 int64
	_ = v779
	var v784 int64
	_ = v784
	var v785 int64
	_ = v785
	var v788 int64
	_ = v788
	var v791 int64
	_ = v791
	var v792 int64
	_ = v792
	var v794 int32
	_ = v794
	var v796 int32
	_ = v796
	var v799 int32
	_ = v799
	var v800 int32
	_ = v800
	var v805 int32
	_ = v805
	var v806 int32
	_ = v806
	var v810 int32
	_ = v810
	var v817 int32
	_ = v817
	var v818 int32
	_ = v818
	var v824 int32
	_ = v824
	var v830 int32
	_ = v830
	var v831 int32
	_ = v831
	var v834 int32
	_ = v834
	var v835 int32
	_ = v835
	var v845 int32
	_ = v845
	var v847 int32
	_ = v847
	var v849 int32
	_ = v849
	var v850 int32
	_ = v850
	var v853 int32
	_ = v853
	var v860 int32
	_ = v860
	var v863 int32
	_ = v863
	var v866 int32
	_ = v866
	var v869 int32
	_ = v869
	var v870 int32
	_ = v870
	var v873 int32
	_ = v873
	var v874 int32
	_ = v874
	var v879 int32
	_ = v879
	var v880 int32
	_ = v880
	var v885 int32
	_ = v885
	var v886 int32
	_ = v886
	var v887 int32
	_ = v887
	var v889 int32
	_ = v889
	var v891 int32
	_ = v891
	var v894 int32
	_ = v894
	var v901 int32
	_ = v901
	var v904 int32
	_ = v904
	var v907 int32
	_ = v907
	var v910 int32
	_ = v910
	var v911 int32
	_ = v911
	var v914 int32
	_ = v914
	var v915 int32
	_ = v915
	var v919 int32
	_ = v919
	var v920 int32
	_ = v920
	var v922 int32
	_ = v922
	var v929 int32
	_ = v929
	var v930 int32
	_ = v930
	var v935 int32
	_ = v935
	var v942 int32
	_ = v942
	var v954 int32
	_ = v954
	var v983 int32
	_ = v983
	var v986 int32
	_ = v986
	var v1006 int32
	_ = v1006
	var v1037 int32
	_ = v1037
	var v1042 int32
	_ = v1042
	var v1043 int32
	_ = v1043
	var v1044 int32
	_ = v1044
	var v1045 int32
	_ = v1045
	var v1046 int32
	_ = v1046
	var v1047 int32
	_ = v1047
	var v1048 int32
	_ = v1048
	var v1050 int32
	_ = v1050
	var v1051 int32
	_ = v1051
	var v1054 int32
	_ = v1054
	var v1055 int64
	_ = v1055
	var v1056 int32
	_ = v1056
	var v1057 int32
	_ = v1057
	var v1058 int32
	_ = v1058
	var v1059 int32
	_ = v1059
	var v1065 int32
	_ = v1065
	var v1069 int64
	_ = v1069
	var v1070 int32
	_ = v1070
	var v1071 int32
	_ = v1071
	var v1072 int32
	_ = v1072
	var v1073 int32
	_ = v1073
	var v1078 int32
	_ = v1078
	var v1082 int64
	_ = v1082
	var v1083 int32
	_ = v1083
	var v1084 int32
	_ = v1084
	var v1085 int32
	_ = v1085
	var v1086 int32
	_ = v1086
	var v1095 int32
	_ = v1095
	var v1096 int32
	_ = v1096
	var v1100 int32
	_ = v1100
	var v1102 int32
	_ = v1102
	var v1105 int32
	_ = v1105
	var v1106 int32
	_ = v1106
	var v1111 int32
	_ = v1111
	var v1112 int32
	_ = v1112
	var v1118 int32
	_ = v1118
	var v1148 int32
	_ = v1148
	var v1152 int64
	_ = v1152
	var v1153 int32
	_ = v1153
	var v1154 int32
	_ = v1154
	var v1155 int32
	_ = v1155
	var v1156 int32
	_ = v1156
	var v1161 int32
	_ = v1161
	var v1166 int32
	_ = v1166
	var v1168 int32
	_ = v1168
	var v1174 int32
	_ = v1174
	var v1191 int32
	_ = v1191
	var v1196 int32
	_ = v1196
	var v1200 int32
	_ = v1200
	var v1201 int32
	_ = v1201
	var v1209 int32
	_ = v1209
	var v1210 int32
	_ = v1210
	var v1213 int32
	_ = v1213
	var v1214 int32
	_ = v1214
	var v1215 int32
	_ = v1215
	var v1216 int32
	_ = v1216
	var v1217 int32
	_ = v1217
	var v1218 int32
	_ = v1218
	var v1220 int32
	_ = v1220
	var v1223 int32
	_ = v1223
	var v1224 int32
	_ = v1224
	var v1225 int32
	_ = v1225
	var v1227 int32
	_ = v1227
	var v1228 int32
	_ = v1228
	var v1230 int64
	_ = v1230
	var v1237 int32
	_ = v1237
	var v1240 int32
	_ = v1240
	var v1243 int32
	_ = v1243
	var v1244 int32
	_ = v1244
	var v1247 int32
	_ = v1247
	var v1249 int32
	_ = v1249
	var v1250 int32
	_ = v1250
	var v1260 int32
	_ = v1260
	var v1265 int32
	_ = v1265
	var v1266 int32
	_ = v1266
	var v1267 int32
	_ = v1267
	var v1268 int32
	_ = v1268
	var v1270 int32
	_ = v1270
	var v1275 int32
	_ = v1275
	var v1276 int32
	_ = v1276
	var v1280 int32
	_ = v1280
	var v1282 int32
	_ = v1282
	var v1283 int32
	_ = v1283
	var v1288 int32
	_ = v1288
	var v1289 int32
	_ = v1289
	var v1292 int32
	_ = v1292
	var v1293 int32
	_ = v1293
	var v1299 int32
	_ = v1299
	var v1300 int32
	_ = v1300
	var v1318 int32
	_ = v1318
	var v1331 int32
	_ = v1331
	var v1332 int32
	_ = v1332
	var v1343 int32
	_ = v1343
	var v1344 int32
	_ = v1344
	var v1346 int32
	_ = v1346
	var v1347 int32
	_ = v1347
	var v1348 int32
	_ = v1348
	var v1349 int32
	_ = v1349
	var v1351 int32
	_ = v1351
	var v1354 int32
	_ = v1354
	var v1363 int32
	_ = v1363
	var v1365 int32
	_ = v1365
	var v1366 int32
	_ = v1366
	var v1367 int32
	_ = v1367
	var v1370 int32
	_ = v1370
	var v1373 int32
	_ = v1373
	var v1374 int32
	_ = v1374
	var v1376 int32
	_ = v1376
	var v1417 int32
	_ = v1417
	var v1419 int32
	_ = v1419
	var v1421 int32
	_ = v1421
	var v1422 int32
	_ = v1422
	var v1424 int32
	_ = v1424
	var v1430 int32
	_ = v1430
	var v1437 int32
	_ = v1437
	var v1444 int32
	_ = v1444
	var v1451 int32
	_ = v1451
	var v1458 int32
	_ = v1458
	var v1465 int32
	_ = v1465
	var v1472 int32
	_ = v1472
	var v1479 int32
	_ = v1479
	var v1486 int32
	_ = v1486
	var v1493 int32
	_ = v1493
	var v1500 int32
	_ = v1500
	var v1507 int32
	_ = v1507
	var v1514 int32
	_ = v1514
	var v1518 int32
	_ = v1518
	var v1520 int32
	_ = v1520
	var v1538 int32
	_ = v1538
	var v1563 int32
	_ = v1563
	var v1566 int32
	_ = v1566
	var v1568 int32
	_ = v1568
	var v1583 int32
	_ = v1583
	var v1608 int32
	_ = v1608
	var v1610 int32
	_ = v1610
	var v1612 int32
	_ = v1612
	var v1614 int32
	_ = v1614
	var v1616 int32
	_ = v1616
	v40 = m.G0
	v42 = v40 - int32(3280)
	m.G0 = v42
	v44 = int32(3)
	v45 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if int32(7) <= v45 {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	m.G0 = v42 + int32(3280)
	return
L2:
	;
	v449 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v450 = *(*int32)(unsafe.Add(mBase, uint32(v449)+20))
	v454 = F_getLongFromObjectOrReply(m, l0, v450, v42+int32(3276), int32(0))
	mBase = m.M
	v455 = m.ExcPending
	if v455 != 0 {
		goto L53
	} else {
		goto L99
	}
L3:
	;
	v53 = int32(0)
	v61 = v45
	v63 = v53
	v64 = v53
	v65 = v53
	v66 = v53
	v67 = int32(6)
	goto L5
L4:
	;
	v49 = int32(0)
	v412 = v44
	v414 = int32(1)
	v415 = v49
	v416 = v49
	v417 = v49
	v418 = v49
	goto L2
L5:
	;
	v97 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v99 = v67 << (uint(int32(2)) % 32)
	v101 = *(*int32)(unsafe.Add(mBase, uint32(v97+v99)))
	v102 = F_objectGetVal(m, v101)
	mBase = m.M
	v103 = int32(_a137)
	v106 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v102))))
	if v106 != 0 {
		goto L11
	} else {
		goto L12
	}
L6:
	;
	v412 = v44
	v414 = v405
	v415 = v399
	v416 = v400
	v417 = v401
	v418 = v402
	goto L2
L7:
	;
	v405 = int32(1)
	v407 = v403 + v405
	v408 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v407 < v408 {
		v61 = v408
		v63 = v399
		v64 = v400
		v65 = v401
		v66 = v402
		v67 = v407
		goto L5
	} else {
		goto L98
	}
L8:
	;
	v143 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v145 = *(*int32)(unsafe.Add(mBase, uint32(v143+v99)))
	v146 = F_objectGetVal(m, v145)
	mBase = m.M
	v147 = int32(_a138)
	v150 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v146))))
	if v150 != 0 {
		goto L25
	} else {
		goto L26
	}
L9:
	;
	if v138-v140 != 0 {
		goto L8
	} else {
		goto L21
	}
L10:
	;
	v138 = F_tolower(m, v134)
	mBase = m.M
	v139 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v135))))
	v140 = F_tolower(m, v139)
	mBase = m.M
	goto L9
L11:
	;
	v108 = v102
	v109 = v103
	v110 = v106
	goto L14
L12:
	;
	v134 = int32(0)
	v135 = v103
	goto L10
L13:
	;
	v134 = v131 & int32(255)
	v135 = v130
	goto L10
L14:
	;
	v112 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v109))))
	if v112 == int32(0) {
		v130 = v109
		v131 = v110
		goto L13
	} else {
		goto L16
	}
L15:
	;
	v130 = v124
	v131 = int32(0)
	goto L13
L16:
	;
	v116 = v110 & int32(255)
	if v116 == v112 {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	v123 = int32(1)
	v124 = v109 + v123
	v125 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v108)+1)))
	if v125 != 0 {
		v108 = v108 + v123
		v109 = v124
		v110 = v125
		goto L14
	} else {
		goto L20
	}
L18:
	;
	v118 = F_tolower(m, v116)
	mBase = m.M
	v119 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v109))))
	v120 = F_tolower(m, v119)
	mBase = m.M
	if v118 == v120 {
		goto L17
	} else {
		goto L19
	}
L19:
	;
	v122 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v108))))
	v130 = v109
	v131 = v122
	goto L13
L20:
	;
	goto L15
L21:
	;
	v399 = v63
	v400 = v64
	v401 = v65
	v402 = int32(1)
	v403 = v67
	goto L7
L22:
	;
	v188 = v67 ^ int32(-1)
	v189 = v61 + v188
	v190 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v192 = *(*int32)(unsafe.Add(mBase, uint32(v190+v99)))
	v193 = F_objectGetVal(m, v192)
	mBase = m.M
	v194 = int32(_a139)
	v197 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v193))))
	if v197 != 0 {
		goto L41
	} else {
		goto L42
	}
L23:
	;
	if v182-v184 != 0 {
		goto L22
	} else {
		goto L35
	}
L24:
	;
	v182 = F_tolower(m, v178)
	mBase = m.M
	v183 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v179))))
	v184 = F_tolower(m, v183)
	mBase = m.M
	goto L23
L25:
	;
	v152 = v146
	v153 = v147
	v154 = v150
	goto L28
L26:
	;
	v178 = int32(0)
	v179 = v147
	goto L24
L27:
	;
	v178 = v175 & int32(255)
	v179 = v174
	goto L24
L28:
	;
	v156 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v153))))
	if v156 == int32(0) {
		v174 = v153
		v175 = v154
		goto L27
	} else {
		goto L30
	}
L29:
	;
	v174 = v168
	v175 = int32(0)
	goto L27
L30:
	;
	v160 = v154 & int32(255)
	if v160 == v156 {
		goto L31
	} else {
		goto L32
	}
L31:
	;
	v167 = int32(1)
	v168 = v153 + v167
	v169 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v152)+1)))
	if v169 != 0 {
		v152 = v152 + v167
		v153 = v168
		v154 = v169
		goto L28
	} else {
		goto L34
	}
L32:
	;
	v162 = F_tolower(m, v160)
	mBase = m.M
	v163 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v153))))
	v164 = F_tolower(m, v163)
	mBase = m.M
	if v162 == v164 {
		goto L31
	} else {
		goto L33
	}
L33:
	;
	v166 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v152))))
	v174 = v153
	v175 = v166
	goto L27
L34:
	;
	goto L29
L35:
	;
	v399 = v63
	v400 = v64
	v401 = int32(1)
	v402 = v66
	v403 = v67
	goto L7
L36:
	;
	v393 = *(*int32)(unsafe.Add(mBase, uint32(v390)))
	v394 = F_objectGetVal(m, v393)
	mBase = m.M
	F_redactClientCommandArgument(m, l0, v392)
	mBase = m.M
	v396 = m.ExcPending
	if v396 != 0 {
		goto L53
	} else {
		goto L97
	}
L37:
	;
	v384 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v386 = v67 + int32(1)
	v390 = v384 + v386<<(uint(int32(2))%32)
	v391 = v64
	v392 = v386
	goto L36
L38:
	;
	v237 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v239 = *(*int32)(unsafe.Add(mBase, uint32(v237+v99)))
	v240 = F_objectGetVal(m, v239)
	mBase = m.M
	v241 = int32(_a140)
	v244 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v240))))
	if v244 != 0 {
		goto L58
	} else {
		goto L59
	}
L39:
	;
	if v229-v231 != 0 {
		goto L38
	} else {
		goto L51
	}
L40:
	;
	v229 = F_tolower(m, v225)
	mBase = m.M
	v230 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v226))))
	v231 = F_tolower(m, v230)
	mBase = m.M
	goto L39
L41:
	;
	v199 = v193
	v200 = v194
	v201 = v197
	goto L44
L42:
	;
	v225 = int32(0)
	v226 = v194
	goto L40
L43:
	;
	v225 = v222 & int32(255)
	v226 = v221
	goto L40
L44:
	;
	v203 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v200))))
	if v203 == int32(0) {
		v221 = v200
		v222 = v201
		goto L43
	} else {
		goto L46
	}
L45:
	;
	v221 = v215
	v222 = int32(0)
	goto L43
L46:
	;
	v207 = v201 & int32(255)
	if v207 == v203 {
		goto L47
	} else {
		goto L48
	}
L47:
	;
	v214 = int32(1)
	v215 = v200 + v214
	v216 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v199)+1)))
	if v216 != 0 {
		v199 = v199 + v214
		v200 = v215
		v201 = v216
		goto L44
	} else {
		goto L50
	}
L48:
	;
	v209 = F_tolower(m, v207)
	mBase = m.M
	v210 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v200))))
	v211 = F_tolower(m, v210)
	mBase = m.M
	if v209 == v211 {
		goto L47
	} else {
		goto L49
	}
L49:
	;
	v213 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v199))))
	v221 = v200
	v222 = v213
	goto L43
L50:
	;
	goto L45
L51:
	;
	if v189 != 0 {
		goto L37
	} else {
		goto L52
	}
L52:
	;
	v234 = *(*int32)(unsafe.Add(mBase, _consts[59]))
	F_addReplyErrorObject(m, l0, v234)
	mBase = m.M
	v236 = m.ExcPending
	if v236 != 0 {
		goto L53
	} else {
		goto L54
	}
L53:
	;
	return
L54:
	;
	goto L1
L55:
	;
	v302 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v306 = *(*int32)(unsafe.Add(mBase, uint32(v302+v67<<(uint(int32(2))%32))))
	v307 = F_objectGetVal(m, v306)
	mBase = m.M
	v308 = int32(_a141)
	v311 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v307))))
	if v311 != 0 {
		goto L77
	} else {
		goto L78
	}
L56:
	;
	if v276-v278 != 0 {
		goto L55
	} else {
		goto L68
	}
L57:
	;
	v276 = F_tolower(m, v272)
	mBase = m.M
	v277 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v273))))
	v278 = F_tolower(m, v277)
	mBase = m.M
	goto L56
L58:
	;
	v246 = v240
	v247 = v241
	v248 = v244
	goto L61
L59:
	;
	v272 = int32(0)
	v273 = v241
	goto L57
L60:
	;
	v272 = v269 & int32(255)
	v273 = v268
	goto L57
L61:
	;
	v250 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v247))))
	if v250 == int32(0) {
		v268 = v247
		v269 = v248
		goto L60
	} else {
		goto L63
	}
L62:
	;
	v268 = v262
	v269 = int32(0)
	goto L60
L63:
	;
	v254 = v248 & int32(255)
	if v254 == v250 {
		goto L64
	} else {
		goto L65
	}
L64:
	;
	v261 = int32(1)
	v262 = v247 + v261
	v263 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v246)+1)))
	if v263 != 0 {
		v246 = v246 + v261
		v247 = v262
		v248 = v263
		goto L61
	} else {
		goto L67
	}
L65:
	;
	v256 = F_tolower(m, v254)
	mBase = m.M
	v257 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v247))))
	v258 = F_tolower(m, v257)
	mBase = m.M
	if v256 == v258 {
		goto L64
	} else {
		goto L66
	}
L66:
	;
	v260 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v246))))
	v268 = v247
	v269 = v260
	goto L60
L67:
	;
	goto L62
L68:
	;
	if int32(1) < v189 {
		goto L69
	} else {
		goto L70
	}
L69:
	;
	v286 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v288 = v67 + int32(1)
	v292 = *(*int32)(unsafe.Add(mBase, uint32(v286+v288<<(uint(int32(2))%32))))
	v293 = F_objectGetVal(m, v292)
	mBase = m.M
	F_redactClientCommandArgument(m, l0, v288)
	mBase = m.M
	v295 = m.ExcPending
	if v295 != 0 {
		goto L53
	} else {
		goto L72
	}
L70:
	;
	v283 = *(*int32)(unsafe.Add(mBase, _consts[59]))
	F_addReplyErrorObject(m, l0, v283)
	mBase = m.M
	v285 = m.ExcPending
	if v285 != 0 {
		goto L53
	} else {
		goto L71
	}
L71:
	;
	goto L1
L72:
	;
	v296 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v297 = int32(2)
	v298 = v67 + v297
	v390 = v296 + v298<<(uint(v297)%32)
	v391 = v293
	v392 = v298
	goto L36
L73:
	;
	v382 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v412 = v67 + int32(1)
	v414 = v382 + v188
	v415 = v63
	v416 = v64
	v417 = v65
	v418 = v66
	goto L2
L74:
	;
	v376 = *(*int32)(unsafe.Add(mBase, _consts[59]))
	F_addReplyErrorObject(m, l0, v376)
	mBase = m.M
	v378 = m.ExcPending
	if v378 != 0 {
		goto L53
	} else {
		goto L96
	}
L75:
	;
	if v343-v345 != 0 {
		goto L74
	} else {
		goto L87
	}
L76:
	;
	v343 = F_tolower(m, v339)
	mBase = m.M
	v344 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v340))))
	v345 = F_tolower(m, v344)
	mBase = m.M
	goto L75
L77:
	;
	v313 = v307
	v314 = v308
	v315 = v311
	goto L80
L78:
	;
	v339 = int32(0)
	v340 = v308
	goto L76
L79:
	;
	v339 = v336 & int32(255)
	v340 = v335
	goto L76
L80:
	;
	v317 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v314))))
	if v317 == int32(0) {
		v335 = v314
		v336 = v315
		goto L79
	} else {
		goto L82
	}
L81:
	;
	v335 = v329
	v336 = int32(0)
	goto L79
L82:
	;
	v321 = v315 & int32(255)
	if v321 == v317 {
		goto L83
	} else {
		goto L84
	}
L83:
	;
	v328 = int32(1)
	v329 = v314 + v328
	v330 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v313)+1)))
	if v330 != 0 {
		v313 = v313 + v328
		v314 = v329
		v315 = v330
		goto L80
	} else {
		goto L86
	}
L84:
	;
	v323 = F_tolower(m, v321)
	mBase = m.M
	v324 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v314))))
	v325 = F_tolower(m, v324)
	mBase = m.M
	if v323 == v325 {
		goto L83
	} else {
		goto L85
	}
L85:
	;
	v327 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v313))))
	v335 = v314
	v336 = v327
	goto L79
L86:
	;
	goto L81
L87:
	;
	v347 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v348 = *(*int32)(unsafe.Add(mBase, uint32(v347)+12))
	v349 = F_objectGetVal(m, v348)
	mBase = m.M
	v352 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v349+int32(-1)))))
	switch v352 & int32(7) {
	case 0:
		goto L93
	case 1:
		goto L92
	case 2:
		goto L91
	case 3:
		goto L90
	case 4:
		goto L89
	default:
		goto L73
	}
L88:
	;
	if v369 == int32(0) {
		goto L73
	} else {
		goto L94
	}
L89:
	;
	v368 = *(*int32)(unsafe.Add(mBase, uint32(v349+int32(-17))))
	v369 = v368
	goto L88
L90:
	;
	v365 = *(*int32)(unsafe.Add(mBase, uint32(v349+int32(-9))))
	v369 = v365
	goto L88
L91:
	;
	v362 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v349+int32(-5)))))
	v369 = v362
	goto L88
L92:
	;
	v359 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v349+int32(-3)))))
	v369 = v359
	goto L88
L93:
	;
	v369 = int32(base.Ui32(v352) >> (uint(int32(3)) % 32))
	goto L88
L94:
	;
	F_addReplyError(m, l0, int32(_a142))
	mBase = m.M
	v374 = m.ExcPending
	if v374 != 0 {
		goto L53
	} else {
		goto L95
	}
L95:
	;
	goto L1
L96:
	;
	goto L1
L97:
	;
	v399 = v394
	v400 = v391
	v401 = v65
	v402 = v66
	v403 = v392
	goto L7
L98:
	;
	goto L6
L99:
	;
	if v454 != 0 {
		goto L1
	} else {
		goto L100
	}
L100:
	;
	v456 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v457 = *(*int32)(unsafe.Add(mBase, uint32(v456)+16))
	v461 = F_getLongFromObjectOrReply(m, l0, v457, v42+int32(3272), int32(0))
	mBase = m.M
	v462 = m.ExcPending
	if v462 != 0 {
		goto L53
	} else {
		goto L101
	}
L101:
	;
	if v461 != 0 {
		goto L1
	} else {
		goto L102
	}
L102:
	;
	v463 = *(*int32)(unsafe.Add(mBase, uint32(v42)+3276))
	if int32(0) < v463 {
		goto L103
	} else {
		goto L104
	}
L103:
	;
	v468 = int32(0)
	v471 = v414 << (uint(int32(2)) % 32)
	v472 = F_valkey_realloc(m, v468, v471)
	mBase = m.M
	v473 = m.ExcPending
	if v473 != 0 {
		goto L53
	} else {
		goto L105
	}
L104:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v42)+3276)) = int32(1000)
	goto L103
L105:
	;
	v475 = F_valkey_realloc(m, int32(0), v471)
	mBase = m.M
	v476 = m.ExcPending
	if v476 != 0 {
		goto L53
	} else {
		goto L106
	}
L106:
	;
	if v414 < int32(1) {
		goto L123
	} else {
		goto L124
	}
L107:
	;
	v1608 = *(*int32)(unsafe.Add(mBase, uint32(v42)+3240))
	F_sdsfree(m, v1608)
	mBase = m.M
	v1610 = m.ExcPending
	if v1610 != 0 {
		goto L53
	} else {
		goto L345
	}
L108:
	;
	v1563 = *(*int32)(unsafe.Add(mBase, uint32(v42)+3272))
	*(*int32)(unsafe.Add(mBase, uint32(v618)+4)) = v1563
	v1566 = *(*int32)(unsafe.Add(mBase, _consts[77]))
	F_addReply(m, l0, v1566)
	mBase = m.M
	v1568 = m.ExcPending
	if v1568 != 0 {
		goto L53
	} else {
		goto L344
	}
L109:
	;
	if v1518 == int32(0) {
		v1583 = v1520
		goto L107
	} else {
		goto L343
	}
L110:
	;
	F__serverAssertWithInfo(m, l0, int32(0), int32(_a143), int32(_a144), int32(593))
	mBase = m.M
	v1514 = m.ExcPending
	if v1514 != 0 {
		goto L53
	} else {
		goto L342
	}
L111:
	;
	F__serverAssertWithInfo(m, l0, int32(0), int32(_a145), int32(_a144), int32(588))
	mBase = m.M
	v1507 = m.ExcPending
	if v1507 != 0 {
		goto L53
	} else {
		goto L341
	}
L112:
	;
	F__serverAssertWithInfo(m, l0, int32(0), int32(_a146), int32(_a144), int32(583))
	mBase = m.M
	v1500 = m.ExcPending
	if v1500 != 0 {
		goto L53
	} else {
		goto L340
	}
L113:
	;
	F__serverAssertWithInfo(m, l0, int32(0), int32(_a147), int32(_a144), int32(582))
	mBase = m.M
	v1493 = m.ExcPending
	if v1493 != 0 {
		goto L53
	} else {
		goto L339
	}
L114:
	;
	F__serverAssertWithInfo(m, l0, int32(0), int32(_a148), int32(_a144), int32(580))
	mBase = m.M
	v1486 = m.ExcPending
	if v1486 != 0 {
		goto L53
	} else {
		goto L338
	}
L115:
	;
	F__serverAssertWithInfo(m, l0, int32(0), int32(_a149), int32(_a144), int32(575))
	mBase = m.M
	v1479 = m.ExcPending
	if v1479 != 0 {
		goto L53
	} else {
		goto L337
	}
L116:
	;
	F__serverAssertWithInfo(m, l0, int32(0), int32(_a150), int32(_a144), int32(548))
	mBase = m.M
	v1472 = m.ExcPending
	if v1472 != 0 {
		goto L53
	} else {
		goto L336
	}
L117:
	;
	F__serverAssertWithInfo(m, l0, int32(0), int32(_a151), int32(_a144), int32(547))
	mBase = m.M
	v1465 = m.ExcPending
	if v1465 != 0 {
		goto L53
	} else {
		goto L335
	}
L118:
	;
	F__serverAssertWithInfo(m, l0, int32(0), int32(_a152), int32(_a144), int32(546))
	mBase = m.M
	v1458 = m.ExcPending
	if v1458 != 0 {
		goto L53
	} else {
		goto L334
	}
L119:
	;
	F__serverAssertWithInfo(m, l0, int32(0), int32(_a153), int32(_a144), int32(540))
	mBase = m.M
	v1451 = m.ExcPending
	if v1451 != 0 {
		goto L53
	} else {
		goto L333
	}
L120:
	;
	F__serverAssertWithInfo(m, l0, int32(0), int32(_a154), int32(_a144), int32(538))
	mBase = m.M
	v1444 = m.ExcPending
	if v1444 != 0 {
		goto L53
	} else {
		goto L332
	}
L121:
	;
	F__serverAssertWithInfo(m, l0, int32(0), int32(_a155), int32(_a144), int32(536))
	mBase = m.M
	v1437 = m.ExcPending
	if v1437 != 0 {
		goto L53
	} else {
		goto L331
	}
L122:
	;
	F__serverAssertWithInfo(m, l0, int32(0), int32(_a156), int32(_a144), int32(535))
	mBase = m.M
	v1430 = m.ExcPending
	if v1430 != 0 {
		goto L53
	} else {
		goto L330
	}
L123:
	;
	F_valkey_free(m, v472)
	mBase = m.M
	v1417 = m.ExcPending
	if v1417 != 0 {
		goto L53
	} else {
		goto L326
	}
L124:
	;
	v489 = v468
	v493 = int32(0)
	goto L125
L125:
	;
	v519 = int32(2)
	v520 = v493 << (uint(v519) % 32)
	v522 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v523 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v526 = (v489 + v412) << (uint(v519) % 32)
	v528 = *(*int32)(unsafe.Add(mBase, uint32(v523+v526)))
	v529 = F_lookupKeyRead(m, v522, v528)
	mBase = m.M
	v530 = m.ExcPending
	if v530 != 0 {
		goto L53
	} else {
		goto L127
	}
L126:
	;
	if v541 == int32(0) {
		goto L123
	} else {
		goto L131
	}
L127:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v472+v520))) = v529
	if v529 == int32(0) {
		v541 = v493
		goto L128
	} else {
		goto L129
	}
L128:
	;
	v543 = v489 + int32(1)
	if v543 != v414 {
		v489 = v543
		v493 = v541
		goto L125
	} else {
		goto L130
	}
L129:
	;
	v535 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v537 = *(*int32)(unsafe.Add(mBase, uint32(v535+v526)))
	*(*int32)(unsafe.Add(mBase, uint32(v475+v520))) = v537
	v541 = v493 + int32(1)
	goto L128
L130:
	;
	goto L126
L131:
	;
	if v417 != 0 {
		goto L132
	} else {
		goto L133
	}
L132:
	;
	v549 = int32(5)
	goto L134
L133:
	;
	v549 = int32(4)
	goto L134
L134:
	;
	if v416 != 0 {
		goto L135
	} else {
		goto L136
	}
L135:
	;
	v552 = int32(3)
	goto L137
L136:
	;
	v552 = int32(2)
	goto L137
L137:
	;
	v553 = int32(-3)
	v555 = int32(-5)
	v557 = int32(-9)
	v559 = int32(-17)
	v561 = int32(-1)
	v588 = v541
	v602 = int32(1)
	v603 = int32(0)
	goto L138
L138:
	;
	v614 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v615 = *(*int32)(unsafe.Add(mBase, uint32(v614)+4))
	v616 = *(*int32)(unsafe.Add(mBase, uint32(v614)+8))
	v617 = *(*int32)(unsafe.Add(mBase, uint32(v42)+3276))
	v618 = F_migrateGetSocket(m, l0, v615, v616, v617)
	mBase = m.M
	v619 = m.ExcPending
	if v619 != 0 {
		goto L53
	} else {
		goto L141
	}
L139:
	;
	F_valkey_free(m, v472)
	mBase = m.M
	v1363 = m.ExcPending
	if v1363 != 0 {
		goto L53
	} else {
		goto L318
	}
L140:
	;
	v626 = F_sdsempty(m)
	mBase = m.M
	v627 = m.ExcPending
	if v627 != 0 {
		goto L53
	} else {
		goto L145
	}
L141:
	;
	if v618 != 0 {
		goto L140
	} else {
		goto L142
	}
L142:
	;
	F_valkey_free(m, v472)
	mBase = m.M
	v621 = m.ExcPending
	if v621 != 0 {
		goto L53
	} else {
		goto L143
	}
L143:
	;
	F_valkey_free(m, v475)
	mBase = m.M
	v623 = m.ExcPending
	if v623 != 0 {
		goto L53
	} else {
		goto L144
	}
L144:
	;
	goto L1
L145:
	;
	v630 = F___memcpy(m, v42+int32(3192), int32(_a157), int32(80))
	mBase = m.M
	*(*int64)(unsafe.Add(mBase, uint32(v630)+56)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v630)+48)) = v626
	goto L146
L146:
	;
	if v415 == int32(0) {
		goto L147
	} else {
		goto L148
	}
L147:
	;
	v691 = *(*int32)(unsafe.Add(mBase, uint32(v618)+4))
	v692 = *(*int32)(unsafe.Add(mBase, uint32(v42)+3272))
	v693 = base.B2i32(v691 == v692)
	if v691 == v692 {
		goto L171
	} else {
		goto L172
	}
L148:
	;
	v639 = F_rioWriteBulkCount(m, v42+int32(3192), int32(42), v552)
	mBase = m.M
	v640 = m.ExcPending
	if v640 != 0 {
		goto L53
	} else {
		goto L149
	}
L149:
	;
	if v639 == int32(0) {
		goto L122
	} else {
		goto L150
	}
L150:
	;
	v647 = F_rioWriteBulkString(m, v42+int32(3192), int32(_a158), int32(4))
	mBase = m.M
	v648 = m.ExcPending
	if v648 != 0 {
		goto L53
	} else {
		goto L151
	}
L151:
	;
	if v647 == int32(0) {
		goto L121
	} else {
		goto L152
	}
L152:
	;
	if v416 == int32(0) {
		goto L153
	} else {
		goto L154
	}
L153:
	;
	v673 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v415+v561))))
	switch v673 & int32(7) {
	case 0:
		goto L168
	case 1:
		goto L167
	case 2:
		goto L166
	case 3:
		goto L165
	case 4:
		goto L164
	default:
		v682 = int32(0)
		goto L163
	}
L154:
	;
	v654 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v416+v561))))
	switch v654 & int32(7) {
	case 0:
		goto L160
	case 1:
		goto L159
	case 2:
		goto L158
	case 3:
		goto L157
	case 4:
		goto L156
	default:
		v663 = int32(0)
		goto L155
	}
L155:
	;
	v666 = F_rioWriteBulkString(m, v42+int32(3192), v416, v663)
	mBase = m.M
	v667 = m.ExcPending
	if v667 != 0 {
		goto L53
	} else {
		goto L161
	}
L156:
	;
	v662 = *(*int32)(unsafe.Add(mBase, uint32(v416+v559)))
	v663 = v662
	goto L155
L157:
	;
	v661 = *(*int32)(unsafe.Add(mBase, uint32(v416+v557)))
	v663 = v661
	goto L155
L158:
	;
	v660 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v416+v555))))
	v663 = v660
	goto L155
L159:
	;
	v659 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v416+v553))))
	v663 = v659
	goto L155
L160:
	;
	v663 = int32(base.Ui32(v654) >> (uint(int32(3)) % 32))
	goto L155
L161:
	;
	if v666 == int32(0) {
		goto L120
	} else {
		goto L162
	}
L162:
	;
	goto L153
L163:
	;
	v685 = F_rioWriteBulkString(m, v42+int32(3192), v415, v682)
	mBase = m.M
	v686 = m.ExcPending
	if v686 != 0 {
		goto L53
	} else {
		goto L169
	}
L164:
	;
	v681 = *(*int32)(unsafe.Add(mBase, uint32(v415+v559)))
	v682 = v681
	goto L163
L165:
	;
	v680 = *(*int32)(unsafe.Add(mBase, uint32(v415+v557)))
	v682 = v680
	goto L163
L166:
	;
	v679 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v415+v555))))
	v682 = v679
	goto L163
L167:
	;
	v678 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v415+v553))))
	v682 = v678
	goto L163
L168:
	;
	v682 = int32(base.Ui32(v673) >> (uint(int32(3)) % 32))
	goto L163
L169:
	;
	if v685 == int32(0) {
		goto L119
	} else {
		goto L170
	}
L170:
	;
	goto L147
L171:
	;
	v717 = int32(0)
	if v588 < int32(1) {
		v954 = v717
		goto L179
	} else {
		goto L180
	}
L172:
	;
	v698 = F_rioWriteBulkCount(m, v42+int32(3192), int32(42), int32(2))
	mBase = m.M
	v699 = m.ExcPending
	if v699 != 0 {
		goto L53
	} else {
		goto L173
	}
L173:
	;
	if v698 == int32(0) {
		goto L118
	} else {
		goto L174
	}
L174:
	;
	v706 = F_rioWriteBulkString(m, v42+int32(3192), int32(_a159), int32(6))
	mBase = m.M
	v707 = m.ExcPending
	if v707 != 0 {
		goto L53
	} else {
		goto L175
	}
L175:
	;
	if v706 == int32(0) {
		goto L117
	} else {
		goto L176
	}
L176:
	;
	v712 = int64(*(*int32)(unsafe.Add(mBase, uint32(v42)+3272)))
	v713 = F_rioWriteBulkLongLong(m, v42+int32(3192), v712)
	mBase = m.M
	v714 = m.ExcPending
	if v714 != 0 {
		goto L53
	} else {
		goto L177
	}
L177:
	;
	if v713 == int32(0) {
		goto L116
	} else {
		goto L178
	}
L178:
	;
	goto L171
L179:
	;
	v983 = int32(9116376)
	goto L232
L180:
	;
	v725 = v717
	v732 = v717
	goto L181
L181:
	;
	v762 = v725 << (uint(int32(2)) % 32)
	v763 = v472 + v762
	v764 = *(*int32)(unsafe.Add(mBase, uint32(v763)))
	v768 = *(*int32)(unsafe.Add(mBase, uint32(v764)+4))
	if v768&int32(1) == int32(0) {
		v779 = int64(-1)
		goto L187
	} else {
		goto L188
	}
L182:
	;
	v954 = v935
	goto L179
L183:
	;
	v942 = v725 + int32(1)
	if v942 != v588 {
		v725 = v942
		v732 = v935
		goto L181
	} else {
		goto L231
	}
L184:
	;
	v794 = v732 << (uint(int32(2)) % 32)
	v796 = *(*int32)(unsafe.Add(mBase, uint32(v763)))
	*(*int32)(unsafe.Add(mBase, uint32(v472+v794))) = v796
	v799 = v475 + v762
	v800 = *(*int32)(unsafe.Add(mBase, uint32(v799)))
	*(*int32)(unsafe.Add(mBase, uint32(v475+v794))) = v800
	v805 = F_rioWriteBulkCount(m, v42+int32(3192), int32(42), v549)
	mBase = m.M
	v806 = m.ExcPending
	if v806 != 0 {
		goto L53
	} else {
		goto L195
	}
L185:
	;
	v784 = *(*int64)(unsafe.Add(mBase, _consts[78]))
	goto L190
L186:
	;
	if v779 != int64(-1) {
		goto L185
	} else {
		goto L189
	}
L187:
	;
	goto L186
L188:
	;
	v778 = *(*int64)(unsafe.Add(mBase, uint32(v764+(v768&int32(4)^int32(12)))))
	v779 = v778
	goto L187
L189:
	;
	v792 = int64(0)
	goto L184
L190:
	;
	v785 = v779 - v784
	if v785 < int64(0) {
		v935 = v732
		goto L183
	} else {
		goto L191
	}
L191:
	;
	v788 = int64(1)
	if base.Ui64(v788) < base.Ui64(v785) {
		goto L192
	} else {
		goto L193
	}
L192:
	;
	v791 = v785
	goto L194
L193:
	;
	v791 = v788
	goto L194
L194:
	;
	v792 = v791
	goto L184
L195:
	;
	if v805 == int32(0) {
		goto L115
	} else {
		goto L196
	}
L196:
	;
	v810 = *(*int32)(unsafe.Add(mBase, _consts[63]))
	if v810 == int32(0) {
		goto L198
	} else {
		goto L199
	}
L197:
	;
	v834 = *(*int32)(unsafe.Add(mBase, uint32(v799)))
	v835 = *(*int32)(unsafe.Add(mBase, uint32(v834)))
	switch int32(base.Ui32(v835)>>(uint(int32(4))%32)) & int32(15) {
	case 0, 8:
		goto L205
	default:
		goto L206
	}
L198:
	;
	v830 = F_rioWriteBulkString(m, v42+int32(3192), int32(_a160), int32(7))
	mBase = m.M
	v831 = m.ExcPending
	if v831 != 0 {
		goto L53
	} else {
		goto L203
	}
L199:
	;
	v817 = F_rioWriteBulkString(m, v42+int32(3192), int32(_a161), int32(14))
	mBase = m.M
	v818 = m.ExcPending
	if v818 != 0 {
		goto L53
	} else {
		goto L200
	}
L200:
	;
	if v817 != 0 {
		goto L197
	} else {
		goto L201
	}
L201:
	;
	F__serverAssertWithInfo(m, l0, int32(0), int32(_a162), int32(_a144), int32(578))
	mBase = m.M
	v824 = m.ExcPending
	if v824 != 0 {
		goto L53
	} else {
		goto L202
	}
L202:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L203:
	;
	if v830 == int32(0) {
		goto L114
	} else {
		goto L204
	}
L204:
	;
	goto L197
L205:
	;
	v847 = F_objectGetVal(m, v834)
	mBase = m.M
	v849 = *(*int32)(unsafe.Add(mBase, uint32(v799)))
	v850 = F_objectGetVal(m, v849)
	mBase = m.M
	v853 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v850+int32(-1)))))
	switch v853 & int32(7) {
	case 0:
		goto L213
	case 1:
		goto L212
	case 2:
		goto L211
	case 3:
		goto L210
	case 4:
		goto L209
	default:
		v870 = int32(0)
		goto L208
	}
L206:
	;
	F__serverAssertWithInfo(m, l0, int32(0), int32(_a163), int32(_a144), int32(581))
	mBase = m.M
	v845 = m.ExcPending
	if v845 != 0 {
		goto L53
	} else {
		goto L207
	}
L207:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L208:
	;
	v873 = F_rioWriteBulkString(m, v42+int32(3192), v847, v870)
	mBase = m.M
	v874 = m.ExcPending
	if v874 != 0 {
		goto L53
	} else {
		goto L214
	}
L209:
	;
	v869 = *(*int32)(unsafe.Add(mBase, uint32(v850+int32(-17))))
	v870 = v869
	goto L208
L210:
	;
	v866 = *(*int32)(unsafe.Add(mBase, uint32(v850+int32(-9))))
	v870 = v866
	goto L208
L211:
	;
	v863 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v850+int32(-5)))))
	v870 = v863
	goto L208
L212:
	;
	v860 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v850+int32(-3)))))
	v870 = v860
	goto L208
L213:
	;
	v870 = int32(base.Ui32(v853) >> (uint(int32(3)) % 32))
	goto L208
L214:
	;
	if v873 == int32(0) {
		goto L113
	} else {
		goto L215
	}
L215:
	;
	v879 = F_rioWriteBulkLongLong(m, v42+int32(3192), v792)
	mBase = m.M
	v880 = m.ExcPending
	if v880 != 0 {
		goto L53
	} else {
		goto L216
	}
L216:
	;
	if v879 == int32(0) {
		goto L112
	} else {
		goto L217
	}
L217:
	;
	v885 = *(*int32)(unsafe.Add(mBase, uint32(v763)))
	v886 = *(*int32)(unsafe.Add(mBase, uint32(v799)))
	v887 = *(*int32)(unsafe.Add(mBase, uint32(v42)+3272))
	F_createDumpPayload(m, v42+int32(3112), v885, v886, v887)
	mBase = m.M
	v889 = m.ExcPending
	if v889 != 0 {
		goto L53
	} else {
		goto L218
	}
L218:
	;
	v891 = *(*int32)(unsafe.Add(mBase, uint32(v42)+3160))
	v894 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v891+int32(-1)))))
	switch v894 & int32(7) {
	case 0:
		goto L224
	case 1:
		goto L223
	case 2:
		goto L222
	case 3:
		goto L221
	case 4:
		goto L220
	default:
		v911 = int32(0)
		goto L219
	}
L219:
	;
	v914 = F_rioWriteBulkString(m, v42+int32(3192), v891, v911)
	mBase = m.M
	v915 = m.ExcPending
	if v915 != 0 {
		goto L53
	} else {
		goto L225
	}
L220:
	;
	v910 = *(*int32)(unsafe.Add(mBase, uint32(v891+int32(-17))))
	v911 = v910
	goto L219
L221:
	;
	v907 = *(*int32)(unsafe.Add(mBase, uint32(v891+int32(-9))))
	v911 = v907
	goto L219
L222:
	;
	v904 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v891+int32(-5)))))
	v911 = v904
	goto L219
L223:
	;
	v901 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v891+int32(-3)))))
	v911 = v901
	goto L219
L224:
	;
	v911 = int32(base.Ui32(v894) >> (uint(int32(3)) % 32))
	goto L219
L225:
	;
	if v914 == int32(0) {
		goto L111
	} else {
		goto L226
	}
L226:
	;
	v919 = v732 + int32(1)
	v920 = *(*int32)(unsafe.Add(mBase, uint32(v42)+3160))
	F_sdsfree(m, v920)
	mBase = m.M
	v922 = m.ExcPending
	if v922 != 0 {
		goto L53
	} else {
		goto L227
	}
L227:
	;
	if v417 == int32(0) {
		v935 = v919
		goto L183
	} else {
		goto L228
	}
L228:
	;
	v929 = F_rioWriteBulkString(m, v42+int32(3192), int32(_a164), int32(7))
	mBase = m.M
	v930 = m.ExcPending
	if v930 != 0 {
		goto L53
	} else {
		goto L229
	}
L229:
	;
	if v929 == int32(0) {
		goto L110
	} else {
		goto L230
	}
L230:
	;
	v935 = v919
	goto L183
L231:
	;
	goto L182
L232:
	;
	*(*int32)(unsafe.Add(mBase, _consts[9])) = int32(0)
	v986 = *(*int32)(unsafe.Add(mBase, uint32(v42)+3240))
	v1006 = v717
	goto L235
L233:
	;
	v1343 = *(*int32)(unsafe.Add(mBase, _consts[9]))
	v1344 = *(*int32)(unsafe.Add(mBase, uint32(v42)+3240))
	F_sdsfree(m, v1344)
	mBase = m.M
	v1346 = m.ExcPending
	if v1346 != 0 {
		goto L53
	} else {
		goto L310
	}
L234:
	;
	if v415 == int32(0) {
		goto L249
	} else {
		goto L250
	}
L235:
	;
	v1037 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v986+int32(-1)))))
	switch v1037 & int32(7) {
	case 0:
		goto L242
	case 1:
		goto L241
	case 2:
		goto L240
	case 3:
		goto L239
	case 4:
		goto L238
	default:
		v1046 = int32(0)
		goto L237
	}
L236:
	;
	v1318 = int32(0)
	v1331 = v602
	v1332 = v603
	goto L233
L237:
	;
	v1047 = base.B2i32(v1046 == v1006)
	if v1046 == v1006 {
		goto L234
	} else {
		goto L243
	}
L238:
	;
	v1045 = *(*int32)(unsafe.Add(mBase, uint32(v986+int32(-17))))
	v1046 = v1045
	goto L237
L239:
	;
	v1044 = *(*int32)(unsafe.Add(mBase, uint32(v986+int32(-9))))
	v1046 = v1044
	goto L237
L240:
	;
	v1043 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v986+int32(-5)))))
	v1046 = v1043
	goto L237
L241:
	;
	v1042 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v986+int32(-3)))))
	v1046 = v1042
	goto L237
L242:
	;
	v1046 = int32(base.Ui32(v1037) >> (uint(int32(3)) % 32))
	goto L237
L243:
	;
	v1048 = *(*int32)(unsafe.Add(mBase, uint32(v618)))
	v1050 = v1046 - v1006
	v1051 = int32(65536)
	if base.Ui32(v1050) < base.Ui32(v1051) {
		goto L244
	} else {
		goto L245
	}
L244:
	;
	v1054 = v1050
	goto L246
L245:
	;
	v1054 = v1051
	goto L246
L246:
	;
	v1055 = int64(*(*int32)(unsafe.Add(mBase, uint32(v42)+3276)))
	v1056 = *(*int32)(unsafe.Add(mBase, uint32(v1048)))
	v1057 = *(*int32)(unsafe.Add(mBase, uint32(v1056)+92))
	v1058 = m.T0[v1057].(func(*base.Module, int32, int32, int32, int64) int32)(m, v1048, v986+v1006, v1054, v1055)
	mBase = m.M
	v1059 = m.ExcPending
	if v1059 != 0 {
		goto L53
	} else {
		goto L247
	}
L247:
	;
	if v1058 == v1054 {
		v1006 = v1058 + v1006
		goto L235
	} else {
		goto L248
	}
L248:
	;
	goto L236
L249:
	;
	if v691 == v692 {
		goto L253
	} else {
		goto L254
	}
L250:
	;
	v1065 = *(*int32)(unsafe.Add(mBase, uint32(v618)))
	v1069 = int64(*(*int32)(unsafe.Add(mBase, uint32(v42)+3276)))
	v1070 = *(*int32)(unsafe.Add(mBase, uint32(v1065)))
	v1071 = *(*int32)(unsafe.Add(mBase, uint32(v1070)+100))
	v1072 = m.T0[v1071].(func(*base.Module, int32, int32, int32, int64) int32)(m, v1065, v42+int32(2080), int32(1024), v1069)
	mBase = m.M
	v1073 = m.ExcPending
	if v1073 != 0 {
		goto L53
	} else {
		goto L251
	}
L251:
	;
	if int32(1) <= v1072 {
		goto L249
	} else {
		goto L252
	}
L252:
	;
	v1318 = int32(0)
	v1331 = v602
	v1332 = v603
	goto L233
L253:
	;
	if v418 != 0 {
		goto L258
	} else {
		goto L259
	}
L254:
	;
	v1078 = *(*int32)(unsafe.Add(mBase, uint32(v618)))
	v1082 = int64(*(*int32)(unsafe.Add(mBase, uint32(v42)+3276)))
	v1083 = *(*int32)(unsafe.Add(mBase, uint32(v1078)))
	v1084 = *(*int32)(unsafe.Add(mBase, uint32(v1083)+100))
	v1085 = m.T0[v1084].(func(*base.Module, int32, int32, int32, int64) int32)(m, v1078, v42+int32(1056), int32(1024), v1082)
	mBase = m.M
	v1086 = m.ExcPending
	if v1086 != 0 {
		goto L53
	} else {
		goto L255
	}
L255:
	;
	if int32(1) <= v1085 {
		goto L253
	} else {
		goto L256
	}
L256:
	;
	v1318 = int32(0)
	v1331 = v602
	v1332 = v603
	goto L233
L257:
	;
	v1106 = int32(0)
	v1111 = int32(1)
	v1112 = v1106
	v1118 = v1106
	goto L268
L258:
	;
	v1102 = int32(0)
	if v954 < int32(1) {
		v1538 = v1102
		goto L108
	} else {
		goto L263
	}
L259:
	;
	v1095 = F_valkey_malloc(m, v954<<(uint(int32(2))%32)+int32(4))
	mBase = m.M
	v1096 = m.ExcPending
	if v1096 != 0 {
		goto L53
	} else {
		goto L260
	}
L260:
	;
	if int32(0) < v954 {
		v1105 = v1095
		goto L257
	} else {
		goto L261
	}
L261:
	;
	F_valkey_free(m, v1095)
	mBase = m.M
	v1100 = m.ExcPending
	if v1100 != 0 {
		goto L53
	} else {
		goto L262
	}
L262:
	;
	v1538 = int32(0)
	goto L108
L263:
	;
	v1105 = v1102
	goto L257
L264:
	;
	if v1280 < int32(2) {
		goto L303
	} else {
		goto L304
	}
L265:
	;
	v1280 = v1111
	v1282 = v1265
	v1283 = v1265
	goto L264
L266:
	;
	if v1276 == int32(0) {
		v1518 = v1275
		v1520 = v1105
		goto L109
	} else {
		goto L302
	}
L267:
	;
	if (v602^int32(-1)|base.B2i32(v1112|v1118 != int32(0)))&int32(1) != 0 {
		goto L297
	} else {
		goto L298
	}
L268:
	;
	v1148 = *(*int32)(unsafe.Add(mBase, uint32(v618)))
	v1152 = int64(*(*int32)(unsafe.Add(mBase, uint32(v42)+3276)))
	v1153 = *(*int32)(unsafe.Add(mBase, uint32(v1148)))
	v1154 = *(*int32)(unsafe.Add(mBase, uint32(v1153)+100))
	v1155 = m.T0[v1154].(func(*base.Module, int32, int32, int32, int64) int32)(m, v1148, v42+int32(32), int32(1024), v1152)
	mBase = m.M
	v1156 = m.ExcPending
	if v1156 != 0 {
		goto L53
	} else {
		goto L270
	}
L269:
	;
	v1249 = int32(0)
	v1250 = base.B2i32(v1244 == v1249)
	if v418 != 0 {
		v1275 = v1250
		v1276 = v1249
		goto L266
	} else {
		goto L296
	}
L270:
	;
	if v1155 < int32(1) {
		goto L267
	} else {
		goto L271
	}
L271:
	;
	v1161 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v42)+2080)))
	v1166 = base.B2i32(v415 != int32(0)) & base.B2i32(v1161&int32(255) == int32(45))
	if v1166 != 0 {
		goto L274
	} else {
		goto L275
	}
L272:
	;
	v1247 = v1118 + int32(1)
	if v1247 != v954 {
		v1111 = v1243
		v1112 = v1244
		v1118 = v1247
		goto L268
	} else {
		goto L295
	}
L273:
	;
	if v418 != 0 {
		v1243 = v1111
		v1244 = v1112
		goto L272
	} else {
		goto L290
	}
L274:
	;
	if v1112 == int32(0) {
		goto L278
	} else {
		goto L279
	}
L275:
	;
	v1168 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v42)+1056)))
	if base.B2i32(v691 != v692)&base.B2i32(v1168&int32(255) == int32(45)) != 0 {
		goto L274
	} else {
		goto L276
	}
L276:
	;
	v1174 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v42)+32)))
	if v1174&int32(255) != int32(45) {
		goto L273
	} else {
		goto L277
	}
L277:
	;
	goto L274
L278:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v618)+4)) = int32(-1)
	v1191 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v42)+1056)))
	if v1191&int32(255) == int32(45) {
		goto L280
	} else {
		goto L281
	}
L279:
	;
	v1243 = v1111
	v1244 = int32(1)
	goto L272
L280:
	;
	v1196 = v42 + int32(1056)
	goto L282
L281:
	;
	v1196 = v42 + int32(32)
	goto L282
L282:
	;
	if v691 != v692 {
		goto L283
	} else {
		goto L284
	}
L283:
	;
	v1200 = v1196
	goto L285
L284:
	;
	v1200 = v42 + int32(32)
	goto L285
L285:
	;
	if v1166 != 0 {
		goto L286
	} else {
		goto L287
	}
L286:
	;
	v1201 = v42 + int32(2080)
	goto L288
L287:
	;
	v1201 = v1200
	goto L288
L288:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v42)+16)) = v1201 | int32(1)
	F_addReplyErrorFormat(m, l0, int32(_a165), v42+int32(16))
	mBase = m.M
	v1209 = m.ExcPending
	if v1209 != 0 {
		goto L53
	} else {
		goto L289
	}
L289:
	;
	v1243 = v1111
	v1244 = int32(1)
	goto L272
L290:
	;
	v1210 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v1213 = v475 + v1118<<(uint(int32(2))%32)
	v1214 = *(*int32)(unsafe.Add(mBase, uint32(v1213)))
	v1215 = F_dbDelete(m, v1210, v1214)
	mBase = m.M
	v1216 = m.ExcPending
	if v1216 != 0 {
		goto L53
	} else {
		goto L291
	}
L291:
	;
	v1217 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v1218 = *(*int32)(unsafe.Add(mBase, uint32(v1213)))
	F_signalModifiedKey(m, l0, v1217, v1218)
	mBase = m.M
	v1220 = m.ExcPending
	if v1220 != 0 {
		goto L53
	} else {
		goto L292
	}
L292:
	;
	v1223 = *(*int32)(unsafe.Add(mBase, uint32(v1213)))
	v1224 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v1225 = *(*int32)(unsafe.Add(mBase, uint32(v1224)+28))
	F_notifyKeyspaceEvent(m, int32(4), int32(_a132), v1223, v1225)
	mBase = m.M
	v1227 = m.ExcPending
	if v1227 != 0 {
		goto L53
	} else {
		goto L293
	}
L293:
	;
	v1228 = int32(_a69)
	v1230 = *(*int64)(unsafe.Add(mBase, _consts[60]))
	*(*int64)(unsafe.Add(mBase, _consts[60])) = v1230 + int64(1)
	v1237 = *(*int32)(unsafe.Add(mBase, uint32(v1213)))
	*(*int32)(unsafe.Add(mBase, uint32(v1105+v1111<<(uint(int32(2))%32)))) = v1237
	F_incrRefCount(m, v1237)
	mBase = m.M
	v1240 = m.ExcPending
	if v1240 != 0 {
		goto L53
	} else {
		goto L294
	}
L294:
	;
	v1243 = v1111 + int32(1)
	v1244 = v1112
	goto L272
L295:
	;
	goto L269
L296:
	;
	v1280 = v1243
	v1282 = v1250
	v1283 = v1249
	goto L264
L297:
	;
	v1265 = base.B2i32(v1112 == int32(0))
	v1266 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v1267 = *(*int32)(unsafe.Add(mBase, uint32(v1266)+4))
	v1268 = *(*int32)(unsafe.Add(mBase, uint32(v1266)+8))
	F_migrateCloseSocket(m, v1267, v1268)
	mBase = m.M
	v1270 = m.ExcPending
	if v1270 != 0 {
		goto L53
	} else {
		goto L300
	}
L298:
	;
	v1260 = *(*int32)(unsafe.Add(mBase, _consts[9]))
	if v1260 == int32(73) {
		goto L297
	} else {
		goto L299
	}
L299:
	;
	v1318 = v1105
	v1331 = int32(1)
	v1332 = v603
	goto L233
L300:
	;
	if v418 == int32(0) {
		goto L265
	} else {
		goto L301
	}
L301:
	;
	v1275 = v1265
	v1276 = v1265
	goto L266
L302:
	;
	v1318 = v1105
	v1331 = int32(0)
	v1332 = v603
	goto L233
L303:
	;
	F_valkey_free(m, v1105)
	mBase = m.M
	v1299 = m.ExcPending
	if v1299 != 0 {
		goto L53
	} else {
		goto L308
	}
L304:
	;
	v1288 = F_createStringObject_1(m, int32(_a166), int32(3))
	mBase = m.M
	v1289 = m.ExcPending
	if v1289 != 0 {
		goto L53
	} else {
		goto L305
	}
L305:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1105))) = v1288
	F_replaceClientCommandVector(m, l0, v1280, v1105)
	mBase = m.M
	v1292 = m.ExcPending
	if v1292 != 0 {
		goto L53
	} else {
		goto L306
	}
L306:
	;
	v1293 = int32(0)
	if v1283 == v1293 {
		v1518 = v1282
		v1520 = v1293
		goto L109
	} else {
		goto L307
	}
L307:
	;
	v1318 = v1293
	v1331 = int32(0)
	v1332 = int32(1)
	goto L233
L308:
	;
	v1300 = int32(0)
	if v1283 == v1300 {
		v1518 = v1282
		v1520 = v1300
		goto L109
	} else {
		goto L309
	}
L309:
	;
	v1318 = v1300
	v1331 = v1300
	v1332 = v603
	goto L233
L310:
	;
	if v1332 != 0 {
		goto L311
	} else {
		goto L312
	}
L311:
	;
	F_valkey_free(m, v1318)
	mBase = m.M
	v1354 = m.ExcPending
	if v1354 != 0 {
		goto L53
	} else {
		goto L314
	}
L312:
	;
	v1347 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v1348 = *(*int32)(unsafe.Add(mBase, uint32(v1347)+4))
	v1349 = *(*int32)(unsafe.Add(mBase, uint32(v1347)+8))
	F_migrateCloseSocket(m, v1348, v1349)
	mBase = m.M
	v1351 = m.ExcPending
	if v1351 != 0 {
		goto L53
	} else {
		goto L313
	}
L313:
	;
	goto L311
L314:
	;
	if v1331 == int32(0) {
		goto L315
	} else {
		goto L316
	}
L315:
	;
	goto L139
L316:
	;
	if v1343 != int32(73) {
		v588 = v954
		v602 = int32(0)
		v603 = v1332
		goto L138
	} else {
		goto L317
	}
L317:
	;
	goto L315
L318:
	;
	F_valkey_free(m, v475)
	mBase = m.M
	v1365 = m.ExcPending
	if v1365 != 0 {
		goto L53
	} else {
		goto L319
	}
L319:
	;
	v1366 = F_sdsempty(m)
	mBase = m.M
	v1367 = m.ExcPending
	if v1367 != 0 {
		goto L53
	} else {
		goto L320
	}
L320:
	;
	if v1046 == v1006 {
		goto L321
	} else {
		goto L322
	}
L321:
	;
	v1370 = int32(_a167)
	goto L323
L322:
	;
	v1370 = int32(_a168)
	goto L323
L323:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v42))) = v1370
	v1373 = F_sdscatprintf(m, v1366, int32(_a169), v42)
	mBase = m.M
	v1374 = m.ExcPending
	if v1374 != 0 {
		goto L53
	} else {
		goto L324
	}
L324:
	;
	F_addReplyErrorSds(m, l0, v1373)
	mBase = m.M
	v1376 = m.ExcPending
	if v1376 != 0 {
		goto L53
	} else {
		goto L325
	}
L325:
	;
	goto L1
L326:
	;
	F_valkey_free(m, v475)
	mBase = m.M
	v1419 = m.ExcPending
	if v1419 != 0 {
		goto L53
	} else {
		goto L327
	}
L327:
	;
	v1421 = F_sdsnew(m, int32(_a170))
	mBase = m.M
	v1422 = m.ExcPending
	if v1422 != 0 {
		goto L53
	} else {
		goto L328
	}
L328:
	;
	F_addReplySds(m, l0, v1421)
	mBase = m.M
	v1424 = m.ExcPending
	if v1424 != 0 {
		goto L53
	} else {
		goto L329
	}
L329:
	;
	goto L1
L330:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L331:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L332:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L333:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L334:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L335:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L336:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L337:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L338:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L339:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L340:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L341:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L342:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L343:
	;
	v1538 = v1520
	goto L108
L344:
	;
	v1583 = v1538
	goto L107
L345:
	;
	F_valkey_free(m, v472)
	mBase = m.M
	v1612 = m.ExcPending
	if v1612 != 0 {
		goto L53
	} else {
		goto L346
	}
L346:
	;
	F_valkey_free(m, v475)
	mBase = m.M
	v1614 = m.ExcPending
	if v1614 != 0 {
		goto L53
	} else {
		goto L347
	}
L347:
	;
	F_valkey_free(m, v1583)
	mBase = m.M
	v1616 = m.ExcPending
	if v1616 != 0 {
		goto L53
	} else {
		goto L348
	}
L348:
	;
	goto L1
}
func F_mixStringObjectDigest(m *base.Module, l0 int32, l1 int32) {
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
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
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
	var v43 int32
	_ = v43
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
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
	var v68 int32
	_ = v68
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
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
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
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v135 int32
	_ = v135
	var v145 int32
	_ = v145
	v8 = m.G0
	v10 = v8 - int32(96)
	m.G0 = v10
	v12 = F_getDecodedObject(m, l1)
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return
	} else {
		v14 = F_objectGetVal(m, v12)
		mBase = m.M
		v16 = F_objectGetVal(m, v12)
		mBase = m.M
		v19 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16+int32(-1)))))
		switch v19 & int32(7) {
		case 0:
			v36 = int32(base.Ui32(v19) >> (uint(int32(3)) % 32))
		case 1:
			v26 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16+int32(-3)))))
			v36 = v26
		case 2:
			v29 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v16+int32(-5)))))
			v36 = v29
		case 3:
			v32 = *(*int32)(unsafe.Add(mBase, uint32(v16+int32(-9))))
			v36 = v32
		case 4:
			v35 = *(*int32)(unsafe.Add(mBase, uint32(v16+int32(-17))))
			v36 = v35
		default:
			v36 = int32(0)
		}
		v38 = m.G0
		v39 = int32(112)
		v40 = v38 - v39
		m.G0 = v40
		v43 = v40 + int32(20)
		F_SHA1Init(m, v43)
		mBase = m.M
		F_SHA1Update(m, v43, v14, v36)
		mBase = m.M
		F_SHA1Final(m, v40, v43)
		mBase = m.M
		v51 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
		v52 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v40))))
		v53 = v51 ^ v52
		*(*uint8)(unsafe.Add(mBase, uint32(l0))) = uint8(v53)
		v55 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+1)))
		v56 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v40)+1)))
		v57 = v55 ^ v56
		*(*uint8)(unsafe.Add(mBase, uint32(l0)+1)) = uint8(v57)
		v59 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+2)))
		v60 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v40)+2)))
		v61 = v59 ^ v60
		*(*uint8)(unsafe.Add(mBase, uint32(l0)+2)) = uint8(v61)
		v63 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+3)))
		v64 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v40)+3)))
		v65 = v63 ^ v64
		*(*uint8)(unsafe.Add(mBase, uint32(l0)+3)) = uint8(v65)
		v67 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+4)))
		v68 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v40)+4)))
		v69 = v67 ^ v68
		*(*uint8)(unsafe.Add(mBase, uint32(l0)+4)) = uint8(v69)
		v71 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+5)))
		v72 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v40)+5)))
		v73 = v71 ^ v72
		*(*uint8)(unsafe.Add(mBase, uint32(l0)+5)) = uint8(v73)
		v75 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+6)))
		v76 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v40)+6)))
		v77 = v75 ^ v76
		*(*uint8)(unsafe.Add(mBase, uint32(l0)+6)) = uint8(v77)
		v79 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+7)))
		v80 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v40)+7)))
		v81 = v79 ^ v80
		*(*uint8)(unsafe.Add(mBase, uint32(l0)+7)) = uint8(v81)
		v83 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+8)))
		v84 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v40)+8)))
		v85 = v83 ^ v84
		*(*uint8)(unsafe.Add(mBase, uint32(l0)+8)) = uint8(v85)
		v87 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+9)))
		v88 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v40)+9)))
		v89 = v87 ^ v88
		*(*uint8)(unsafe.Add(mBase, uint32(l0)+9)) = uint8(v89)
		v91 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+10)))
		v92 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v40)+10)))
		v93 = v91 ^ v92
		*(*uint8)(unsafe.Add(mBase, uint32(l0)+10)) = uint8(v93)
		v95 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+11)))
		v96 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v40)+11)))
		v97 = v95 ^ v96
		*(*uint8)(unsafe.Add(mBase, uint32(l0)+11)) = uint8(v97)
		v99 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+12)))
		v100 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v40)+12)))
		v101 = v99 ^ v100
		*(*uint8)(unsafe.Add(mBase, uint32(l0)+12)) = uint8(v101)
		v103 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+13)))
		v104 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v40)+13)))
		v105 = v103 ^ v104
		*(*uint8)(unsafe.Add(mBase, uint32(l0)+13)) = uint8(v105)
		v107 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+14)))
		v108 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v40)+14)))
		v109 = v107 ^ v108
		*(*uint8)(unsafe.Add(mBase, uint32(l0)+14)) = uint8(v109)
		v111 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+15)))
		v112 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v40)+15)))
		v113 = v111 ^ v112
		*(*uint8)(unsafe.Add(mBase, uint32(l0)+15)) = uint8(v113)
		v115 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)))
		v116 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v40)+16)))
		v117 = v115 ^ v116
		*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v117)
		v119 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+17)))
		v120 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v40)+17)))
		v121 = v119 ^ v120
		*(*uint8)(unsafe.Add(mBase, uint32(l0)+17)) = uint8(v121)
		v123 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+18)))
		v124 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v40)+18)))
		v125 = v123 ^ v124
		*(*uint8)(unsafe.Add(mBase, uint32(l0)+18)) = uint8(v125)
		v127 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+19)))
		v128 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v40)+19)))
		v129 = v127 ^ v128
		*(*uint8)(unsafe.Add(mBase, uint32(l0)+19)) = uint8(v129)
		m.G0 = v40 + v39
		v135 = v10 + int32(4)
		F_SHA1Init(m, v135)
		mBase = m.M
		F_SHA1Update(m, v135, l0, int32(20))
		mBase = m.M
		F_SHA1Final(m, l0, v135)
		mBase = m.M
		F_decrRefCount(m, v12)
		mBase = m.M
		v145 = m.ExcPending
		if v145 != 0 {
			return
		} else {
			m.G0 = v10 + int32(96)
			return
		}
	}
}
func F_mkdir(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v12 int32
	_ = v12
	v4 = m.Env.X__syscall_mkdirat(m, int32(-100), l0, l1)
	mBase = m.M
	if base.Ui32(v4) < base.Ui32(int32(-4095)) {
		v12 = v4
	} else {
		v7 = F___errno_location(m)
		mBase = m.M
		*(*int32)(unsafe.Add(mBase, uint32(v7))) = int32(0) - v4
		v12 = int32(-1)
	}
	return v12
}
func F_mkstemp(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	var v36 int32
	_ = v36
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v57 int32
	_ = v57
	var v69 int32
	_ = v69
	v2 = int32(0)
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	v11 = F_strlen(m, l0)
	mBase = m.M
	if base.Ui32(v11) < base.Ui32(int32(6)) {
		goto L5
	} else {
		goto L6
	}
L1:
	;
	return v69
L2:
	;
	m.G0 = v9 + int32(16)
	goto L1
L3:
	;
	v69 = int32(-1)
	goto L2
L4:
	;
	v36 = int32(100)
	goto L9
L5:
	;
	v27 = F___errno_location(m)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v27))) = int32(28)
	goto L3
L6:
	;
	if base.Ui32(v11+int32(-6)) < base.Ui32(v2) {
		goto L5
	} else {
		goto L7
	}
L7:
	;
	v20 = l0 + v11 - v2 + int32(-6)
	v23 = F_memcmp(m, v20, int32(_a2353), int32(6))
	mBase = m.M
	if v23 == int32(0) {
		goto L4
	} else {
		goto L8
	}
L8:
	;
	goto L5
L9:
	;
	v41 = F___randname(m, v20)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v9))) = int32(384)
	v44 = F_open(m, l0, int32(194), v9)
	mBase = m.M
	if int32(-1) < v44 {
		v69 = v44
		goto L2
	} else {
		goto L11
	}
L10:
	;
	v57 = F___memcpy(m, v20, int32(_a2353), int32(6))
	mBase = m.M
	goto L3
L11:
	;
	v48 = v36 + int32(-1)
	if v48 == int32(0) {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	goto L10
L13:
	;
	v51 = F___errno_location(m)
	mBase = m.M
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v51)))
	if v52 == int32(20) {
		v36 = v48
		goto L9
	} else {
		goto L14
	}
L14:
	;
	goto L12
}
func F_modf(m *base.Module, l0 float64, l1 int32) float64 {
	mBase := m.M
	_ = mBase
	var v7 int64
	_ = v7
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v34 int64
	_ = v34
	var v47 int64
	_ = v47
	var v51 float64
	_ = v51
	v7 = base.I64_reinterpret_f64(l0)
	v12 = base.I32_wrap_i64(int64(base.Ui64(v7)>>(uint(int64(52))%64))) & int32(2047)
	v14 = v12 + int32(-1023)
	if base.Ui32(v12) < base.Ui32(int32(1075)) {
		if base.Ui32(int32(1022)) < base.Ui32(v12) {
			v34 = base.I64_extend_i32_u(v14)
			if v7<<(uint(v34)%64)&int64(4503599627370495) != int64(0) {
				v47 = int64(-4503599627370496) >> (uint(v34) % 64) & v7
				*(*int64)(unsafe.Add(mBase, uint32(l1))) = v47
				v51 = base.F64_sub(l0, base.F64_reinterpret_i64(v47))
				return v51
			} else {
				*(*float64)(unsafe.Add(mBase, uint32(l1))) = l0
				return base.F64_reinterpret_i64(v7 & int64(-9223372036854775807-1))
			}
		} else {
			*(*int64)(unsafe.Add(mBase, uint32(l1))) = v7 & int64(-9223372036854775807-1)
			return l0
		}
	} else {
		*(*float64)(unsafe.Add(mBase, uint32(l1))) = l0
		if v7&int64(4503599627370495) == int64(0) {
			return base.F64_reinterpret_i64(v7 & int64(-9223372036854775807-1))
		} else {
			if v14 == int32(1024) {
				v51 = l0
				return v51
			} else {
				return base.F64_reinterpret_i64(v7 & int64(-9223372036854775807-1))
			}
		}
	}
}
func F_modulesCron(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v1 int32
	_ = v1
	var v5 int32
	_ = v5
	var v9 int32
	_ = v9
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v56 int32
	_ = v56
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	v1 = int32(0)
	v5 = *(*int32)(unsafe.Add(mBase, _consts[345]))
	if v5 == v1 {
		v44 = v5
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v47 = int32(0)
	*(*int32)(unsafe.Add(mBase, _consts[349])) = v44
	v50 = *(*int32)(unsafe.Add(mBase, _consts[346]))
	if base.Ui32(v50) < base.Ui32(int32(33)) {
		goto L11
	} else {
		goto L12
	}
L2:
	;
	v9 = *(*int32)(unsafe.Add(mBase, _consts[349]))
	if base.Ui32(v9) < base.Ui32(int32(9)) {
		v44 = v5
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v13 = v5
	v14 = int32(50)
	goto L4
L4:
	;
	v16 = int32(0)
	v18 = v13 + int32(-1)
	*(*int32)(unsafe.Add(mBase, _consts[345])) = v18
	v21 = *(*int32)(unsafe.Add(mBase, _consts[347]))
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v21+v18<<(uint(int32(2))%32))))
	v26 = F_freeClient(m, v25)
	mBase = m.M
	v27 = m.ExcPending
	if v27 != 0 {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	v44 = v35
	goto L1
L6:
	;
	return
L7:
	;
	v28 = int32(0)
	v30 = *(*int32)(unsafe.Add(mBase, _consts[349]))
	v32 = v30 + int32(-1)
	*(*int32)(unsafe.Add(mBase, _consts[349])) = v32
	v35 = *(*int32)(unsafe.Add(mBase, _consts[345]))
	if base.Ui32(v14) < base.Ui32(int32(2)) {
		v44 = v35
		goto L1
	} else {
		goto L8
	}
L8:
	;
	if v35 == int32(0) {
		v44 = v35
		goto L1
	} else {
		goto L9
	}
L9:
	;
	if base.Ui32(int32(8)) < base.Ui32(v32) {
		v13 = v35
		v14 = v14 + int32(-1)
		goto L4
	} else {
		goto L10
	}
L10:
	;
	goto L5
L11:
	;
	return
L12:
	;
	if base.Ui32(v50) <= base.Ui32(v44<<(uint(int32(2))%32)) {
		goto L11
	} else {
		goto L13
	}
L13:
	;
	v56 = int32(0)
	*(*int32)(unsafe.Add(mBase, _consts[346])) = int32(base.Ui32(v50) >> (uint(int32(2)) % 32))
	v62 = *(*int32)(unsafe.Add(mBase, _consts[347]))
	v65 = F_valkey_realloc(m, v62, v50&int32(-4))
	mBase = m.M
	v66 = m.ExcPending
	if v66 != 0 {
		goto L6
	} else {
		goto L14
	}
L14:
	;
	*(*int32)(unsafe.Add(mBase, _consts[347])) = v65
	goto L11
}
func F_mpscDequeueBatch(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v8 != v9 {
		v14 = v9
		goto L3
	} else {
		goto L4
	}
L1:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v59
	return v60
L2:
	;
	return int32(0)
L3:
	;
	v15 = v14 - v8
	if base.Ui32(v15) < base.Ui32(l2) {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v11
	if v8 == v11 {
		goto L2
	} else {
		goto L5
	}
L5:
	;
	v14 = v11
	goto L3
L6:
	;
	v17 = v15
	goto L8
L7:
	;
	v17 = l2
	goto L8
L8:
	;
	if v17 == int32(0) {
		goto L2
	} else {
		goto L9
	}
L9:
	;
	v24 = v8
	v25 = int32(0)
	goto L11
L10:
	;
	if v25 != 0 {
		v59 = v24
		v60 = v25
		goto L1
	} else {
		goto L15
	}
L11:
	;
	v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
	v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)+132))
	v35 = v28 + (v29+int32(-1))&v24<<(uint(int32(2))%32)
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v35)))
	if v36 == int32(0) {
		goto L10
	} else {
		goto L13
	}
L12:
	;
	v59 = v46
	v60 = v17
	goto L1
L13:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1+v25<<(uint(int32(2))%32)))) = v36
	*(*int32)(unsafe.Add(mBase, uint32(v35))) = int32(0)
	v45 = int32(1)
	v46 = v24 + v45
	v48 = v25 + v45
	if v48 != v17 {
		v24 = v46
		v25 = v48
		goto L11
	} else {
		goto L14
	}
L14:
	;
	goto L12
L15:
	;
	goto L2
}
func F_mpscEnqueue(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v38 int32
	_ = v38
	var v48 int32
	_ = v48
	if l1 == int32(0) {
		F__serverAssert(m, int32(_a1684), int32(_a1685), int32(40))
		mBase = m.M
		v48 = m.ExcPending
		if v48 != 0 {
			return int32(0)
		} else {
			F_abort(m)
			mBase = m.M
			base.Wasm_trap_unreachable()
			for {
			}
		}
	} else {
		v8 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+4)))
		if v8 != 0 {
			v13 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
			v14 = v13
		} else {
			v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
			*(*int32)(unsafe.Add(mBase, uint32(l0)+64)) = v9 + int32(1)
			v14 = v9
		}
		v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
		v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+132))
		if base.Ui32(v14-v15) < base.Ui32(v17) {
			v29 = v17
			v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
			*(*int32)(unsafe.Add(mBase, uint32(v30+(v29+int32(-1))&v14<<(uint(int32(2))%32)))) = l1
			v38 = int32(0)
			*(*uint8)(unsafe.Add(mBase, uint32(l2)+4)) = uint8(v38)
			return int32(1)
		} else {
			v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			*(*int32)(unsafe.Add(mBase, uint32(l0)+68)) = v19
			v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+132))
			if base.Ui32(v14-v19) < base.Ui32(v22) {
				v29 = v22
				v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
				*(*int32)(unsafe.Add(mBase, uint32(v30+(v29+int32(-1))&v14<<(uint(int32(2))%32)))) = l1
				v38 = int32(0)
				*(*uint8)(unsafe.Add(mBase, uint32(l2)+4)) = uint8(v38)
				return int32(1)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(l2))) = v14
				v25 = int32(1)
				*(*uint8)(unsafe.Add(mBase, uint32(l2)+4)) = uint8(v25)
				return int32(0)
			}
		}
	}
}
func F_msetCommand(m *base.Module, l0 int32) {
	var v4 int32
	_ = v4
	F_msetGenericCommand(m, l0, int32(0))
	v4 = m.ExcPending
	if v4 != 0 {
		return
	} else {
		return
	}
}
func F_msetexCommand(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v40 int64
	_ = v40
	var v42 int64
	_ = v42
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v73 int64
	_ = v73
	var v86 int64
	_ = v86
	var v87 int64
	_ = v87
	var v88 int64
	_ = v88
	var v94 int32
	_ = v94
	var v98 int32
	_ = v98
	var v114 int64
	_ = v114
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v141 int32
	_ = v141
	var v143 int32
	_ = v143
	var v149 int32
	_ = v149
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v163 int32
	_ = v163
	var v170 int32
	_ = v170
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v184 int32
	_ = v184
	var v186 int32
	_ = v186
	var v188 int32
	_ = v188
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v203 int64
	_ = v203
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v212 int32
	_ = v212
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v218 int64
	_ = v218
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v222 int32
	_ = v222
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v227 int32
	_ = v227
	var v230 int32
	_ = v230
	var v232 int64
	_ = v232
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v247 int32
	_ = v247
	var v248 int32
	_ = v248
	var v250 int32
	_ = v250
	var v251 int32
	_ = v251
	var v257 int32
	_ = v257
	var v260 int32
	_ = v260
	var v271 int32
	_ = v271
	var v274 int32
	_ = v274
	var v277 int64
	_ = v277
	var v278 int32
	_ = v278
	var v279 int32
	_ = v279
	var v280 int32
	_ = v280
	var v282 int32
	_ = v282
	var v284 int32
	_ = v284
	var v285 int32
	_ = v285
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
	v2 = int32(0)
	v12 = m.G0
	v14 = v12 - int32(32)
	m.G0 = v14
	*(*int32)(unsafe.Add(mBase, uint32(v14)+28)) = v2
	*(*int32)(unsafe.Add(mBase, uint32(v14)+24)) = v2
	*(*int32)(unsafe.Add(mBase, uint32(v14)+20)) = int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(v14)+16)) = v2
	*(*int64)(unsafe.Add(mBase, uint32(v14)+8)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v14)+4)) = v2
	v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v28)+4))
	v35 = F_getRangeLongFromObjectOrReply(m, l0, v29, int32(1), int32(2147483647), v14+int32(28), int32(_a1593))
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v14 + int32(32)
	return
L2:
	;
	return
L3:
	;
	if v35 != 0 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v14)+28))
	v40 = base.I64_extend_i32_s(v37) << (uint(int64(1)) % 64)
	v42 = v40 + int64(2)
	v43 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v42 <= base.I64_extend_i32_s(v43) {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	v50 = int32(4)
	v61 = F_parseExtendedCommandArgumentsOrReply(m, l0, v50, base.I32_wrap_i64(v42), v43, v14+v50, v14+int32(24), v14+int32(20), v14+int32(16), int32(0))
	mBase = m.M
	v62 = m.ExcPending
	if v62 != 0 {
		goto L2
	} else {
		goto L8
	}
L6:
	;
	v47 = *(*int32)(unsafe.Add(mBase, _consts[59]))
	F_addReplyErrorObject(m, l0, v47)
	mBase = m.M
	v49 = m.ExcPending
	if v49 != 0 {
		goto L2
	} else {
		goto L7
	}
L7:
	;
	goto L1
L8:
	;
	if v61 != 0 {
		goto L1
	} else {
		goto L9
	}
L9:
	;
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v14)+16))
	if v63 == int32(0) {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	v98 = *(*int32)(unsafe.Add(mBase, uint32(v14)+4))
	if v98&int32(3) == int32(0) {
		v149 = v98
		goto L25
	} else {
		goto L26
	}
L11:
	;
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v14)+24))
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v14)+4))
	v71 = F_getLongLongFromObjectOrReply(m, l0, v63, v14+int32(8), int32(0))
	mBase = m.M
	v72 = m.ExcPending
	if v72 != 0 {
		goto L2
	} else {
		goto L12
	}
L12:
	;
	if v71 != 0 {
		goto L1
	} else {
		goto L13
	}
L13:
	;
	v73 = *(*int64)(unsafe.Add(mBase, uint32(v14)+8))
	if v73 < int64(1) {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	F_addReplyErrorExpireTime(m, l0)
	mBase = m.M
	v94 = m.ExcPending
	if v94 != 0 {
		goto L2
	} else {
		goto L24
	}
L15:
	;
	if v66 != 0 {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	if v66 != 0 {
		goto L19
	} else {
		goto L20
	}
L17:
	;
	if base.Ui64(int64(9223372036854775)) < base.Ui64(v73) {
		goto L14
	} else {
		goto L18
	}
L18:
	;
	goto L16
L19:
	;
	if v67&int32(12) == int32(0) {
		goto L10
	} else {
		goto L21
	}
L20:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v14)+8)) = v73 * int64(1000)
	goto L19
L21:
	;
	v86 = *(*int64)(unsafe.Add(mBase, _consts[78]))
	goto L22
L22:
	;
	v87 = *(*int64)(unsafe.Add(mBase, uint32(v14)+8))
	v88 = v86 + v87
	*(*int64)(unsafe.Add(mBase, uint32(v14)+8)) = v88
	if int64(0) < v88 {
		goto L10
	} else {
		goto L23
	}
L23:
	;
	goto L14
L24:
	;
	goto L1
L25:
	;
	v155 = *(*int32)(unsafe.Add(mBase, uint32(v14)+16))
	v156 = *(*int32)(unsafe.Add(mBase, uint32(v14)+28))
	if v156 < int32(1) {
		v260 = v155
		goto L40
	} else {
		goto L41
	}
L26:
	;
	if v37 < int32(1) {
		v149 = v98
		goto L25
	} else {
		goto L27
	}
L27:
	;
	v114 = int64(2)
	goto L29
L28:
	;
	v141 = *(*int32)(unsafe.Add(mBase, _consts[233]))
	F_addReply(m, l0, v141)
	mBase = m.M
	v143 = m.ExcPending
	if v143 != 0 {
		goto L2
	} else {
		goto L39
	}
L29:
	;
	v117 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v118 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v123 = *(*int32)(unsafe.Add(mBase, uint32(v118+base.I32_wrap_i64(v114)<<(uint(int32(2))%32))))
	v124 = F_lookupKeyWrite(m, v117, v123)
	mBase = m.M
	v125 = m.ExcPending
	if v125 != 0 {
		goto L2
	} else {
		goto L31
	}
L31:
	;
	v126 = *(*int32)(unsafe.Add(mBase, uint32(v14)+4))
	if v126&int32(1) == int32(0) {
		goto L32
	} else {
		goto L33
	}
L32:
	;
	if v126&int32(2) == int32(0) {
		goto L35
	} else {
		goto L36
	}
L33:
	;
	if v124 != 0 {
		goto L28
	} else {
		goto L34
	}
L34:
	;
	goto L32
L35:
	;
	if v40 <= v114 {
		v149 = v126
		goto L25
	} else {
		goto L38
	}
L36:
	;
	if v124 == int32(0) {
		goto L28
	} else {
		goto L37
	}
L37:
	;
	goto L35
L38:
	;
	v114 = v114 + int64(2)
	goto L29
L39:
	;
	goto L1
L40:
	;
	if v260 == int32(0) {
		goto L67
	} else {
		goto L68
	}
L41:
	;
	v163 = base.B2i32(v149&int32(16)|v155 != int32(0))
	v170 = int32(2)
	goto L42
L42:
	;
	v176 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v177 = int32(2)
	v180 = *(*int32)(unsafe.Add(mBase, uint32(v176+v170<<(uint(v177)%32))))
	v181 = int32(1)
	v182 = v170 | v181
	v184 = v182 << (uint(v177) % 32)
	v186 = *(*int32)(unsafe.Add(mBase, uint32(v176+v184)))
	*(*int32)(unsafe.Add(mBase, uint32(v14))) = v186
	v188 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+208)))
	if v188&v181 == int32(0) {
		goto L45
	} else {
		goto L46
	}
L43:
	;
	v257 = *(*int32)(unsafe.Add(mBase, uint32(v14)+16))
	v260 = v257
	goto L40
L44:
	;
	v230 = int32(_a69)
	v232 = *(*int64)(unsafe.Add(mBase, _consts[60]))
	*(*int64)(unsafe.Add(mBase, _consts[60])) = v232 + int64(1)
	v238 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v239 = *(*int32)(unsafe.Add(mBase, uint32(v238)+28))
	F_notifyKeyspaceEvent(m, int32(8), int32(_a131), v180, v239)
	mBase = m.M
	v241 = m.ExcPending
	if v241 != 0 {
		goto L2
	} else {
		goto L61
	}
L45:
	;
	v209 = F_tryObjectEncoding(m, v186)
	mBase = m.M
	v210 = m.ExcPending
	if v210 != 0 {
		goto L2
	} else {
		goto L54
	}
L46:
	;
	F_incrRefCount(m, v186)
	mBase = m.M
	v194 = m.ExcPending
	if v194 != 0 {
		goto L2
	} else {
		goto L47
	}
L47:
	;
	v195 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	F_setKey(m, l0, v195, v180, v14, v163)
	mBase = m.M
	v197 = m.ExcPending
	if v197 != 0 {
		goto L2
	} else {
		goto L48
	}
L48:
	;
	v198 = *(*int32)(unsafe.Add(mBase, uint32(v14)+16))
	if v198 != 0 {
		goto L49
	} else {
		goto L50
	}
L49:
	;
	v202 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v203 = *(*int64)(unsafe.Add(mBase, uint32(v14)+8))
	v204 = F_setExpire(m, l0, v202, v180, v203)
	mBase = m.M
	v205 = m.ExcPending
	if v205 != 0 {
		goto L2
	} else {
		goto L52
	}
L50:
	;
	v199 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
	F_rewriteClientCommandArgument(m, l0, v182, v199)
	mBase = m.M
	v201 = m.ExcPending
	if v201 != 0 {
		goto L2
	} else {
		goto L51
	}
L51:
	;
	goto L44
L52:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14))) = v204
	F_rewriteClientCommandArgument(m, l0, v182, v204)
	mBase = m.M
	v208 = m.ExcPending
	if v208 != 0 {
		goto L2
	} else {
		goto L53
	}
L53:
	;
	goto L44
L54:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14))) = v209
	v212 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	F_setKey(m, l0, v212, v180, v14, v163)
	mBase = m.M
	v214 = m.ExcPending
	if v214 != 0 {
		goto L2
	} else {
		goto L55
	}
L55:
	;
	v215 = *(*int32)(unsafe.Add(mBase, uint32(v14)+16))
	if v215 != 0 {
		goto L57
	} else {
		goto L58
	}
L56:
	;
	F_incrRefCount(m, v222)
	mBase = m.M
	v224 = m.ExcPending
	if v224 != 0 {
		goto L2
	} else {
		goto L60
	}
L57:
	;
	v217 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v218 = *(*int64)(unsafe.Add(mBase, uint32(v14)+8))
	v219 = F_setExpire(m, l0, v217, v180, v218)
	mBase = m.M
	v220 = m.ExcPending
	if v220 != 0 {
		goto L2
	} else {
		goto L59
	}
L58:
	;
	v216 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
	v222 = v216
	goto L56
L59:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14))) = v219
	v222 = v219
	goto L56
L60:
	;
	v225 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v227 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
	*(*int32)(unsafe.Add(mBase, uint32(v225+v184))) = v227
	goto L44
L61:
	;
	v242 = *(*int32)(unsafe.Add(mBase, uint32(v14)+16))
	if v242 == int32(0) {
		goto L62
	} else {
		goto L63
	}
L62:
	;
	v251 = *(*int32)(unsafe.Add(mBase, uint32(v14)+28))
	if v170 < v251<<(uint(int32(1))%32) {
		v170 = v170 + int32(2)
		goto L42
	} else {
		goto L65
	}
L63:
	;
	v247 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v248 = *(*int32)(unsafe.Add(mBase, uint32(v247)+28))
	F_notifyKeyspaceEvent(m, int32(4), int32(_a576), v180, v248)
	mBase = m.M
	v250 = m.ExcPending
	if v250 != 0 {
		goto L2
	} else {
		goto L64
	}
L64:
	;
	goto L62
L65:
	;
	goto L43
L66:
	;
	F__serverAssert(m, int32(_a1594), int32(_a1589), int32(687))
	mBase = m.M
	v301 = m.ExcPending
	if v301 != 0 {
		goto L2
	} else {
		goto L76
	}
L67:
	;
	v294 = *(*int32)(unsafe.Add(mBase, _consts[234]))
	F_addReply(m, l0, v294)
	mBase = m.M
	v296 = m.ExcPending
	if v296 != 0 {
		goto L2
	} else {
		goto L75
	}
L68:
	;
	v271 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+4)))
	if v271&int32(128) != 0 {
		goto L67
	} else {
		goto L69
	}
L69:
	;
	v274 = *(*int32)(unsafe.Add(mBase, uint32(v14)+20))
	if v274 == int32(-1) {
		goto L66
	} else {
		goto L70
	}
L70:
	;
	v277 = *(*int64)(unsafe.Add(mBase, uint32(v14)+8))
	v278 = F_createStringObjectFromLongLong(m, v277)
	mBase = m.M
	v279 = m.ExcPending
	if v279 != 0 {
		goto L2
	} else {
		goto L71
	}
L71:
	;
	v280 = *(*int32)(unsafe.Add(mBase, uint32(v14)+20))
	v282 = *(*int32)(unsafe.Add(mBase, _consts[757]))
	F_rewriteClientCommandArgument(m, l0, v280, v282)
	mBase = m.M
	v284 = m.ExcPending
	if v284 != 0 {
		goto L2
	} else {
		goto L72
	}
L72:
	;
	v285 = *(*int32)(unsafe.Add(mBase, uint32(v14)+20))
	F_rewriteClientCommandArgument(m, l0, v285+int32(1), v278)
	mBase = m.M
	v289 = m.ExcPending
	if v289 != 0 {
		goto L2
	} else {
		goto L73
	}
L73:
	;
	F_decrRefCount(m, v278)
	mBase = m.M
	v291 = m.ExcPending
	if v291 != 0 {
		goto L2
	} else {
		goto L74
	}
L74:
	;
	goto L67
L75:
	;
	goto L1
L76:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_msetnxCommand(m *base.Module, l0 int32) {
	var v4 int32
	_ = v4
	F_msetGenericCommand(m, l0, int32(1))
	v4 = m.ExcPending
	if v4 != 0 {
		return
	} else {
		return
	}
}
func F_mstime(m *base.Module) int64 {
	mBase := m.M
	_ = mBase
	var v1 int32
	_ = v1
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v13 int32
	_ = v13
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int64
	_ = v20
	var v22 int64
	_ = v22
	var v23 int64
	_ = v23
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v33 int64
	_ = v33
	var v36 int64
	_ = v36
	var v37 int64
	_ = v37
	var v40 int64
	_ = v40
	var v42 int64
	_ = v42
	var v48 int64
	_ = v48
	v1 = int32(0)
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	v9 = *(*int32)(unsafe.Add(mBase, _consts[271]))
	if v9 == v1 {
		v30 = int32(0)
		v31 = F___gettimeofday(m, v6, v30)
		mBase = m.M
		v33 = *(*int64)(unsafe.Add(mBase, uint32(v6)))
		v36 = int64(*(*int32)(unsafe.Add(mBase, uint32(v6)+8)))
		v37 = v33*int64(1000000) + v36
		*(*int64)(unsafe.Add(mBase, _consts[1115])) = v37
		v42 = v37
	} else {
		v13 = *(*int32)(unsafe.Add(mBase, _consts[271]))
		if base.B2i32(v13 != int32(948)) != int32(1) {
			v30 = int32(0)
			v31 = F___gettimeofday(m, v6, v30)
			mBase = m.M
			v33 = *(*int64)(unsafe.Add(mBase, uint32(v6)))
			v36 = int64(*(*int32)(unsafe.Add(mBase, uint32(v6)+8)))
			v37 = v33*int64(1000000) + v36
			*(*int64)(unsafe.Add(mBase, _consts[1115])) = v37
			v42 = v37
		} else {
			v18 = int32(0)
			v19 = *(*int32)(unsafe.Add(mBase, _consts[271]))
			v20 = m.T0[v19].(func(*base.Module) int64)(m)
			mBase = m.M
			v22 = *(*int64)(unsafe.Add(mBase, _consts[1116]))
			v23 = v20 - v22
			if base.Ui64(v23) < base.Ui64(int64(1000)) {
				v40 = *(*int64)(unsafe.Add(mBase, _consts[1115]))
				v42 = v40 + v23
			} else {
				*(*int64)(unsafe.Add(mBase, _consts[1116])) = v20
				v30 = int32(0)
				v31 = F___gettimeofday(m, v6, v30)
				mBase = m.M
				v33 = *(*int64)(unsafe.Add(mBase, uint32(v6)))
				v36 = int64(*(*int32)(unsafe.Add(mBase, uint32(v6)+8)))
				v37 = v33*int64(1000000) + v36
				*(*int64)(unsafe.Add(mBase, _consts[1115])) = v37
				v42 = v37
			}
		}
	}
	m.G0 = v6 + int32(16)
	v48 = base.I64_div_s(v42, int64(1000))
	return v48
}
func F_multiCommand(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+112))
	if v4 == int32(0) {
		v9 = F_valkey_calloc(m, int32(56))
		mBase = m.M
		v10 = m.ExcPending
		if v10 != 0 {
			return
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(l0)+112)) = v9
			v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
			v13 = *(*int32)(unsafe.Add(mBase, uint32(v12)+28))
			*(*int32)(unsafe.Add(mBase, uint32(v9)+52)) = v13
			v15 = v9
			v16 = v12
			v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+200))
			*(*int32)(unsafe.Add(mBase, uint32(l0)+200)) = v17 | int32(8)
			v21 = *(*int32)(unsafe.Add(mBase, uint32(v16)+28))
			*(*int32)(unsafe.Add(mBase, uint32(v15)+52)) = v21
			v24 = *(*int32)(unsafe.Add(mBase, _consts[77]))
			F_addReply(m, l0, v24)
			mBase = m.M
			v26 = m.ExcPending
			if v26 != 0 {
				return
			} else {
				return
			}
		}
	} else {
		v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
		v15 = v4
		v16 = v7
		v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+200))
		*(*int32)(unsafe.Add(mBase, uint32(l0)+200)) = v17 | int32(8)
		v21 = *(*int32)(unsafe.Add(mBase, uint32(v16)+28))
		*(*int32)(unsafe.Add(mBase, uint32(v15)+52)) = v21
		v24 = *(*int32)(unsafe.Add(mBase, _consts[77]))
		F_addReply(m, l0, v24)
		mBase = m.M
		v26 = m.ExcPending
		if v26 != 0 {
			return
		} else {
			return
		}
	}
}
