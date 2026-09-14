package p2

import (
	base "github.com/shibukawa/valkeymem/internal/aot/vkaot/base"
	"unsafe"
)

func F_IOThreadsAfterSleep(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v29 int32
	_ = v29
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v45 int64
	_ = v45
	var v53 int64
	_ = v53
	var v56 int64
	_ = v56
	var v59 int64
	_ = v59
	var v62 int64
	_ = v62
	var v65 int64
	_ = v65
	var v68 int64
	_ = v68
	var v71 int64
	_ = v71
	var v74 int64
	_ = v74
	var v77 int64
	_ = v77
	var v80 int64
	_ = v80
	var v83 int64
	_ = v83
	var v86 int64
	_ = v86
	var v89 int64
	_ = v89
	var v92 int64
	_ = v92
	var v95 int64
	_ = v95
	var v98 int64
	_ = v98
	var v115 int64
	_ = v115
	var v125 int32
	_ = v125
	var v127 int32
	_ = v127
	var v129 int32
	_ = v129
	var v131 int32
	_ = v131
	var v134 int32
	_ = v134
	var v141 int32
	_ = v141
	var v143 int64
	_ = v143
	var v147 int32
	_ = v147
	var v150 int32
	_ = v150
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
	var v164 int32
	_ = v164
	var v166 int32
	_ = v166
	var v169 int64
	_ = v169
	var v170 int64
	_ = v170
	var v171 int64
	_ = v171
	var v177 int32
	_ = v177
	var v178 int64
	_ = v178
	var v181 int64
	_ = v181
	var v185 int64
	_ = v185
	var v188 int64
	_ = v188
	var v189 int64
	_ = v189
	var v190 int32
	_ = v190
	var v200 int32
	_ = v200
	var v207 int32
	_ = v207
	var v215 int64
	_ = v215
	var v218 int64
	_ = v218
	var v221 int64
	_ = v221
	var v224 int64
	_ = v224
	var v227 int64
	_ = v227
	var v230 int64
	_ = v230
	var v233 int64
	_ = v233
	var v236 int64
	_ = v236
	var v239 int64
	_ = v239
	var v242 int64
	_ = v242
	var v245 int64
	_ = v245
	var v248 int64
	_ = v248
	var v251 int64
	_ = v251
	var v254 int64
	_ = v254
	var v257 int64
	_ = v257
	var v260 int64
	_ = v260
	var v277 int64
	_ = v277
	var v279 int32
	_ = v279
	var v280 int32
	_ = v280
	var v284 int32
	_ = v284
	var v289 int64
	_ = v289
	var v296 int32
	_ = v296
	var v298 int32
	_ = v298
	var v302 int32
	_ = v302
	var v311 int32
	_ = v311
	var v313 int32
	_ = v313
	var v315 int32
	_ = v315
	var v318 int32
	_ = v318
	var v328 int32
	_ = v328
	var v333 int32
	_ = v333
	var v338 int32
	_ = v338
	var v339 int32
	_ = v339
	var v341 int32
	_ = v341
	var v343 int32
	_ = v343
	var v345 int32
	_ = v345
	var v350 int32
	_ = v350
	var v353 int32
	_ = v353
	var v354 int32
	_ = v354
	var v356 int32
	_ = v356
	var v358 int32
	_ = v358
	var v360 int32
	_ = v360
	var v368 int32
	_ = v368
	var v370 int32
	_ = v370
	var v372 int32
	_ = v372
	var v375 int32
	_ = v375
	var v385 int32
	_ = v385
	var v399 int32
	_ = v399
	v7 = m.G0
	v9 = v7 - int32(48)
	m.G0 = v9
	v12 = *(*int32)(unsafe.Add(mBase, _consts[436]))
	if v12 == int32(1) {
		m.G0 = v9 + int32(48)
		return
	} else {
		v16 = *(*int32)(unsafe.Add(mBase, _consts[430]))
		if v16 != 0 {
			F__serverAssert(m, int32(_a845), int32(_a846), int32(154))
			mBase = m.M
			v399 = m.ExcPending
			if v399 != 0 {
				return
			} else {
				F_abort(m)
				mBase = m.M
				base.Wasm_trap_unreachable()
				for {
				}
			}
		} else {
			v17 = int32(_a44)
			v18 = *(*int32)(unsafe.Add(mBase, _consts[431]))
			v20 = *(*int32)(unsafe.Add(mBase, _consts[437]))
			if v20 == int32(0) {
				v45 = *(*int64)(unsafe.Add(mBase, _consts[35]))
				if v18 != int32(1) {
					v143 = *(*int64)(unsafe.Add(mBase, _consts[438]))
					if v45-v143 < int64(10) {
						m.G0 = v9 + int32(48)
						return
					} else {
						v147 = int32(0)
						*(*int64)(unsafe.Add(mBase, _consts[438])) = v45
						v150 = int32(_a847)
						v153 = *(*int32)(unsafe.Add(mBase, _consts[439]))
						v154 = *(*int32)(unsafe.Add(mBase, _consts[440]))
						v155 = v153 - v154
						if base.Ui32(v153) < base.Ui32(v155) {
							v157 = v147
						} else {
							v157 = v155
						}
						v158 = int32(0)
						v159 = *(*int32)(unsafe.Add(mBase, _consts[441]))
						v160 = v157 + v159
						*(*int32)(unsafe.Add(mBase, _consts[441])) = v160
						v164 = *(*int32)(unsafe.Add(mBase, _consts[442]))
						v166 = v164 + int32(1)
						*(*int32)(unsafe.Add(mBase, _consts[442])) = v166
						v169 = base.I64_extend_i32_u(v160)
						v170 = base.I64_extend_i32_u(v166)
						v171 = int64(1)
						v177 = int32(_a535)
						v178 = *(*int64)(unsafe.Add(mBase, _consts[443]))
						if v178 < v171 {
						} else {
							v181 = v170 - v178
							if int64(1) <= v181 {
								v185 = *(*int64)(unsafe.Add(mBase, _consts[444]))
								v188 = base.I64_div_s((v169-v185)*v171, v181)
								v189 = v188
							} else {
								v189 = int64(0)
							}
							v190 = *(*int32)(unsafe.Add(mBase, _consts[445]))
							*(*int64)(unsafe.Add(mBase, uint32(v190<<(uint(int32(3))%32))+uint32(_consts[446]))) = v189
							v200 = base.I32_rem_s(v190+int32(1), int32(16))
							*(*int32)(unsafe.Add(mBase, _consts[445])) = v200
						}
						*(*int64)(unsafe.Add(mBase, _consts[444])) = v169
						*(*int64)(unsafe.Add(mBase, _consts[443])) = v170
						v207 = int32(*(*uint8)(unsafe.Add(mBase, _consts[442])))
						if v207&int32(15) != 0 {
							m.G0 = v9 + int32(48)
							return
						} else {
							v215 = *(*int64)(unsafe.Add(mBase, _consts[447]))
							v218 = *(*int64)(unsafe.Add(mBase, _consts[448]))
							v221 = *(*int64)(unsafe.Add(mBase, _consts[449]))
							v224 = *(*int64)(unsafe.Add(mBase, _consts[450]))
							v227 = *(*int64)(unsafe.Add(mBase, _consts[451]))
							v230 = *(*int64)(unsafe.Add(mBase, _consts[452]))
							v233 = *(*int64)(unsafe.Add(mBase, _consts[453]))
							v236 = *(*int64)(unsafe.Add(mBase, _consts[454]))
							v239 = *(*int64)(unsafe.Add(mBase, _consts[455]))
							v242 = *(*int64)(unsafe.Add(mBase, _consts[456]))
							v245 = *(*int64)(unsafe.Add(mBase, _consts[457]))
							v248 = *(*int64)(unsafe.Add(mBase, _consts[458]))
							v251 = *(*int64)(unsafe.Add(mBase, _consts[459]))
							v254 = *(*int64)(unsafe.Add(mBase, _consts[460]))
							v257 = *(*int64)(unsafe.Add(mBase, _consts[461]))
							v260 = *(*int64)(unsafe.Add(mBase, _consts[446]))
							v277 = base.I64_div_s(v215+(v218+(v221+(v224+(v227+(v230+(v233+(v236+(v239+(v242+(v245+(v248+(v251+(v254+(v257+v260)))))))))))))), int64(16))
							v279 = *(*int32)(unsafe.Add(mBase, _consts[431]))
							v280 = base.I32_wrap_i64(v277)
							if base.Ui32(v280) < base.Ui32(int32(2)) {
								if v280 != 0 {
									m.G0 = v9 + int32(48)
									return
								} else {
									v289 = *(*int64)(unsafe.Add(mBase, _consts[462]))
									v296 = v279 - base.B2i32(int64(1000) < v45-v289)&base.B2i32(base.Ui32(int32(1)) < base.Ui32(v279))
									if base.Ui32(v296) <= base.Ui32(v279) {
										if base.Ui32(v279) <= base.Ui32(v296) {
											m.G0 = v9 + int32(48)
											return
										} else {
											v333 = (v279 + int32(-1)) * int32(192)
											v338 = *(*int32)(unsafe.Add(mBase, uint32(v333)+uint32(_consts[432])))
											v339 = *(*int32)(unsafe.Add(mBase, uint32(v333)+uint32(_consts[463])))
											if v338 == v339 {
												v345 = int32(1)
											} else {
												v341 = *(*int32)(unsafe.Add(mBase, uint32(v333)+uint32(_consts[464])))
												*(*int32)(unsafe.Add(mBase, uint32(v333)+uint32(_consts[463]))) = v341
												v343 = *(*int32)(unsafe.Add(mBase, uint32(v333)+uint32(_consts[432])))
												v345 = base.B2i32(v341 == v343)
											}
											if v345 == int32(0) {
												m.G0 = v9 + int32(48)
												return
											} else {
												if v296 != int32(1) {
													v368 = int32(_a44)
													v370 = *(*int32)(unsafe.Add(mBase, _consts[431]))
													v372 = v370 + int32(-1)
													*(*int32)(unsafe.Add(mBase, _consts[431])) = v372
													v375 = *(*int32)(unsafe.Add(mBase, _consts[6]))
													if int32(0) < v375 {
														m.G0 = v9 + int32(48)
														return
													} else {
														*(*int32)(unsafe.Add(mBase, uint32(v9)+36)) = v372
														*(*int32)(unsafe.Add(mBase, uint32(v9)+32)) = v279
														F__serverLog(m, int32(0), int32(_a848), v9+int32(32))
														mBase = m.M
														v385 = m.ExcPending
														if v385 != 0 {
															return
														} else {
															m.G0 = v9 + int32(48)
															return
														}
													}
												} else {
													v350 = int32(_a847)
													v353 = *(*int32)(unsafe.Add(mBase, _consts[439]))
													v354 = *(*int32)(unsafe.Add(mBase, _consts[465]))
													if v353 == v354 {
														v360 = int32(1)
													} else {
														v356 = *(*int32)(unsafe.Add(mBase, _consts[440]))
														*(*int32)(unsafe.Add(mBase, _consts[465])) = v356
														v358 = *(*int32)(unsafe.Add(mBase, _consts[439]))
														v360 = base.B2i32(v356 == v358)
													}
													if v360 == int32(0) {
														m.G0 = v9 + int32(48)
														return
													} else {
														v368 = int32(_a44)
														v370 = *(*int32)(unsafe.Add(mBase, _consts[431]))
														v372 = v370 + int32(-1)
														*(*int32)(unsafe.Add(mBase, _consts[431])) = v372
														v375 = *(*int32)(unsafe.Add(mBase, _consts[6]))
														if int32(0) < v375 {
															m.G0 = v9 + int32(48)
															return
														} else {
															*(*int32)(unsafe.Add(mBase, uint32(v9)+36)) = v372
															*(*int32)(unsafe.Add(mBase, uint32(v9)+32)) = v279
															F__serverLog(m, int32(0), int32(_a848), v9+int32(32))
															mBase = m.M
															v385 = m.ExcPending
															if v385 != 0 {
																return
															} else {
																m.G0 = v9 + int32(48)
																return
															}
														}
													}
												}
											}
										}
									} else {
										v298 = v296
										v302 = v279
										for {
											v311 = v302 + int32(1)
											if v311 != v298 {
												v302 = v311
												continue
											} else {
												break
											}
											break
										}
										v313 = int32(_a44)
										*(*int32)(unsafe.Add(mBase, _consts[431])) = v298
										v315 = int32(0)
										*(*int64)(unsafe.Add(mBase, _consts[462])) = v45
										v318 = *(*int32)(unsafe.Add(mBase, _consts[6]))
										if v315 < v318 {
											m.G0 = v9 + int32(48)
											return
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(v9)+20)) = v298
											*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = v279
											F__serverLog(m, int32(0), int32(_a849), v9+int32(16))
											mBase = m.M
											v328 = m.ExcPending
											if v328 != 0 {
												return
											} else {
												m.G0 = v9 + int32(48)
												return
											}
										}
									}
								}
							} else {
								v284 = *(*int32)(unsafe.Add(mBase, _consts[436]))
								if base.Ui32(v284) <= base.Ui32(v279) {
									if v280 != 0 {
										m.G0 = v9 + int32(48)
										return
									} else {
										v289 = *(*int64)(unsafe.Add(mBase, _consts[462]))
										v296 = v279 - base.B2i32(int64(1000) < v45-v289)&base.B2i32(base.Ui32(int32(1)) < base.Ui32(v279))
										if base.Ui32(v296) <= base.Ui32(v279) {
											if base.Ui32(v279) <= base.Ui32(v296) {
												m.G0 = v9 + int32(48)
												return
											} else {
												v333 = (v279 + int32(-1)) * int32(192)
												v338 = *(*int32)(unsafe.Add(mBase, uint32(v333)+uint32(_consts[432])))
												v339 = *(*int32)(unsafe.Add(mBase, uint32(v333)+uint32(_consts[463])))
												if v338 == v339 {
													v345 = int32(1)
												} else {
													v341 = *(*int32)(unsafe.Add(mBase, uint32(v333)+uint32(_consts[464])))
													*(*int32)(unsafe.Add(mBase, uint32(v333)+uint32(_consts[463]))) = v341
													v343 = *(*int32)(unsafe.Add(mBase, uint32(v333)+uint32(_consts[432])))
													v345 = base.B2i32(v341 == v343)
												}
												if v345 == int32(0) {
													m.G0 = v9 + int32(48)
													return
												} else {
													if v296 != int32(1) {
														v368 = int32(_a44)
														v370 = *(*int32)(unsafe.Add(mBase, _consts[431]))
														v372 = v370 + int32(-1)
														*(*int32)(unsafe.Add(mBase, _consts[431])) = v372
														v375 = *(*int32)(unsafe.Add(mBase, _consts[6]))
														if int32(0) < v375 {
															m.G0 = v9 + int32(48)
															return
														} else {
															*(*int32)(unsafe.Add(mBase, uint32(v9)+36)) = v372
															*(*int32)(unsafe.Add(mBase, uint32(v9)+32)) = v279
															F__serverLog(m, int32(0), int32(_a848), v9+int32(32))
															mBase = m.M
															v385 = m.ExcPending
															if v385 != 0 {
																return
															} else {
																m.G0 = v9 + int32(48)
																return
															}
														}
													} else {
														v350 = int32(_a847)
														v353 = *(*int32)(unsafe.Add(mBase, _consts[439]))
														v354 = *(*int32)(unsafe.Add(mBase, _consts[465]))
														if v353 == v354 {
															v360 = int32(1)
														} else {
															v356 = *(*int32)(unsafe.Add(mBase, _consts[440]))
															*(*int32)(unsafe.Add(mBase, _consts[465])) = v356
															v358 = *(*int32)(unsafe.Add(mBase, _consts[439]))
															v360 = base.B2i32(v356 == v358)
														}
														if v360 == int32(0) {
															m.G0 = v9 + int32(48)
															return
														} else {
															v368 = int32(_a44)
															v370 = *(*int32)(unsafe.Add(mBase, _consts[431]))
															v372 = v370 + int32(-1)
															*(*int32)(unsafe.Add(mBase, _consts[431])) = v372
															v375 = *(*int32)(unsafe.Add(mBase, _consts[6]))
															if int32(0) < v375 {
																m.G0 = v9 + int32(48)
																return
															} else {
																*(*int32)(unsafe.Add(mBase, uint32(v9)+36)) = v372
																*(*int32)(unsafe.Add(mBase, uint32(v9)+32)) = v279
																F__serverLog(m, int32(0), int32(_a848), v9+int32(32))
																mBase = m.M
																v385 = m.ExcPending
																if v385 != 0 {
																	return
																} else {
																	m.G0 = v9 + int32(48)
																	return
																}
															}
														}
													}
												}
											}
										} else {
											v298 = v296
											v302 = v279
											for {
												v311 = v302 + int32(1)
												if v311 != v298 {
													v302 = v311
													continue
												} else {
													break
												}
												break
											}
											v313 = int32(_a44)
											*(*int32)(unsafe.Add(mBase, _consts[431])) = v298
											v315 = int32(0)
											*(*int64)(unsafe.Add(mBase, _consts[462])) = v45
											v318 = *(*int32)(unsafe.Add(mBase, _consts[6]))
											if v315 < v318 {
												m.G0 = v9 + int32(48)
												return
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(v9)+20)) = v298
												*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = v279
												F__serverLog(m, int32(0), int32(_a849), v9+int32(16))
												mBase = m.M
												v328 = m.ExcPending
												if v328 != 0 {
													return
												} else {
													m.G0 = v9 + int32(48)
													return
												}
											}
										}
									}
								} else {
									v298 = v279 + int32(1)
									v302 = v279
									for {
										v311 = v302 + int32(1)
										if v311 != v298 {
											v302 = v311
											continue
										} else {
											break
										}
										break
									}
									v313 = int32(_a44)
									*(*int32)(unsafe.Add(mBase, _consts[431])) = v298
									v315 = int32(0)
									*(*int64)(unsafe.Add(mBase, _consts[462])) = v45
									v318 = *(*int32)(unsafe.Add(mBase, _consts[6]))
									if v315 < v318 {
										m.G0 = v9 + int32(48)
										return
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v9)+20)) = v298
										*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = v279
										F__serverLog(m, int32(0), int32(_a849), v9+int32(16))
										mBase = m.M
										v328 = m.ExcPending
										if v328 != 0 {
											return
										} else {
											m.G0 = v9 + int32(48)
											return
										}
									}
								}
							}
						}
					}
				} else {
					v53 = *(*int64)(unsafe.Add(mBase, _consts[466]))
					v56 = *(*int64)(unsafe.Add(mBase, _consts[467]))
					v59 = *(*int64)(unsafe.Add(mBase, _consts[468]))
					v62 = *(*int64)(unsafe.Add(mBase, _consts[469]))
					v65 = *(*int64)(unsafe.Add(mBase, _consts[470]))
					v68 = *(*int64)(unsafe.Add(mBase, _consts[471]))
					v71 = *(*int64)(unsafe.Add(mBase, _consts[472]))
					v74 = *(*int64)(unsafe.Add(mBase, _consts[473]))
					v77 = *(*int64)(unsafe.Add(mBase, _consts[474]))
					v80 = *(*int64)(unsafe.Add(mBase, _consts[475]))
					v83 = *(*int64)(unsafe.Add(mBase, _consts[476]))
					v86 = *(*int64)(unsafe.Add(mBase, _consts[477]))
					v89 = *(*int64)(unsafe.Add(mBase, _consts[478]))
					v92 = *(*int64)(unsafe.Add(mBase, _consts[479]))
					v95 = *(*int64)(unsafe.Add(mBase, _consts[480]))
					v98 = *(*int64)(unsafe.Add(mBase, _consts[481]))
					v115 = base.I64_div_s(v53+(v56+(v59+(v62+(v65+(v68+(v71+(v74+(v77+(v80+(v83+(v86+(v89+(v92+(v95+v98)))))))))))))), int64(16))
					if base.F32_gt(base.F32_div(base.F32_convert_i64_s(v115), float32(10000)), float32(30)) == int32(0) {
						m.G0 = v9 + int32(48)
						return
					} else {
						v125 = int32(0)
						*(*int64)(unsafe.Add(mBase, _consts[462])) = v45
						v127 = int32(_a44)
						v129 = *(*int32)(unsafe.Add(mBase, _consts[431]))
						v131 = v129 + int32(1)
						*(*int32)(unsafe.Add(mBase, _consts[431])) = v131
						v134 = *(*int32)(unsafe.Add(mBase, _consts[6]))
						if v125 < v134 {
							m.G0 = v9 + int32(48)
							return
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v9))) = v131
							F__serverLog(m, int32(0), int32(_a850), v9)
							mBase = m.M
							v141 = m.ExcPending
							if v141 != 0 {
								return
							} else {
								m.G0 = v9 + int32(48)
								return
							}
						}
					}
				}
			} else {
				if l0 < int32(1) {
				} else {
					if v12 <= v18 {
					} else {
						v29 = v18
						for {
							v38 = v29 + int32(1)
							v40 = *(*int32)(unsafe.Add(mBase, _consts[436]))
							if v38 < v40 {
								v29 = v38
								continue
							} else {
								break
							}
							break
						}
						*(*int32)(unsafe.Add(mBase, _consts[431])) = v40
					}
				}
				m.G0 = v9 + int32(48)
				return
			}
		}
	}
}
func F_initIOThreads(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v26 int32
	_ = v26
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v60 int32
	_ = v60
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
	var v89 int32
	_ = v89
	var v95 int32
	_ = v95
	var v104 int32
	_ = v104
	var v111 int32
	_ = v111
	var v114 int32
	_ = v114
	var v119 int32
	_ = v119
	var v144 int32
	_ = v144
	var v150 int32
	_ = v150
	var v156 int32
	_ = v156
	v5 = m.G0
	v7 = v5 - int32(64)
	m.G0 = v7
	v10 = *(*int32)(unsafe.Add(mBase, _consts[436]))
	if v10 == int32(1) {
		m.G0 = v7 + int32(64)
		return
	} else {
		if int32(257) <= v10 {
			F__serverAssert(m, int32(_a851), int32(_a846), int32(489))
			mBase = m.M
			v144 = m.ExcPending
			if v144 != 0 {
				return
			} else {
				F_abort(m)
				mBase = m.M
				base.Wasm_trap_unreachable()
				for {
				}
			}
		} else {
			v16 = int32(*(*uint8)(unsafe.Add(mBase, _consts[482])))
			if v16 != 0 {
				v44 = v10
				if v44 <= l0 {
					m.G0 = v7 + int32(64)
					return
				} else {
					if v44 <= int32(0) {
						F__serverAssert(m, int32(_a852), int32(_a846), int32(387))
						mBase = m.M
						v150 = m.ExcPending
						if v150 != 0 {
							return
						} else {
							F_abort(m)
							mBase = m.M
							base.Wasm_trap_unreachable()
							for {
							}
						}
					} else {
						if l0 <= int32(0) {
							F__serverAssert(m, int32(_a853), int32(_a846), int32(388))
							mBase = m.M
							v156 = m.ExcPending
							if v156 != 0 {
								return
							} else {
								F_abort(m)
								mBase = m.M
								base.Wasm_trap_unreachable()
								for {
								}
							}
						} else {
							F_spscInit(m, l0*int32(192)+int32(_a854), int32(4096))
							mBase = m.M
							v60 = m.ExcPending
							if v60 != 0 {
								return
							} else {
								v68 = int32(16)
								v69 = v7 + v68
								v73 = m.G0
								v75 = v73 - v68
								m.G0 = v75
								v77 = F_pthread_attr_init(m, v69)
								mBase = m.M
								v80 = F_pthread_attr_getstacksize(m, v69, v75+int32(12))
								mBase = m.M
								v81 = *(*int32)(unsafe.Add(mBase, uint32(v75)+12))
								v82 = int32(1)
								if base.Ui32(v82) < base.Ui32(v81) {
									v85 = v81
								} else {
									v85 = v82
								}
								v89 = v85
								for {
									if base.Ui32(v89) < base.Ui32(int32(4194304)) {
										v89 = v89 << (uint(int32(1)) % 32)
										continue
									} else {
										break
									}
									break
								}
								*(*int32)(unsafe.Add(mBase, uint32(v75)+12)) = v89
								v95 = F_pthread_attr_setstacksize(m, v69, v89)
								mBase = m.M
								m.G0 = v75 + int32(16)
								v104 = int32(6)
								v111 = *(*int32)(unsafe.Add(mBase, _consts[6]))
								if int32(3) < v111 {
									m.Env.Exit(m, int32(1))
									mBase = m.M
									base.Wasm_trap_unreachable()
									for {
									}
								} else {
									v114 = F___strerror_l(m, v104, v104)
									mBase = m.M
									*(*int32)(unsafe.Add(mBase, uint32(v7))) = v114
									F__serverLog(m, int32(3), int32(_a855), v7)
									mBase = m.M
									v119 = m.ExcPending
									if v119 != 0 {
										return
									} else {
										m.Env.Exit(m, int32(1))
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
				v17 = int32(_a44)
				*(*int64)(unsafe.Add(mBase, _consts[483])) = int64(0)
				*(*int32)(unsafe.Add(mBase, _consts[431])) = int32(1)
				F_spmcInit(m, int32(_a847), int32(4096))
				mBase = m.M
				v26 = m.ExcPending
				if v26 != 0 {
					return
				} else {
					F_mpscInit(m, int32(_a856), int32(16384))
					mBase = m.M
					v30 = m.ExcPending
					if v30 != 0 {
						return
					} else {
						v31 = int32(0)
						*(*int32)(unsafe.Add(mBase, _consts[435])) = v31
						*(*int32)(unsafe.Add(mBase, _consts[434])) = v31
						F_prefetchCommandsBatchInit(m)
						mBase = m.M
						v38 = m.ExcPending
						if v38 != 0 {
							return
						} else {
							v40 = int32(1)
							*(*uint8)(unsafe.Add(mBase, _consts[482])) = uint8(v40)
							v43 = *(*int32)(unsafe.Add(mBase, _consts[436]))
							v44 = v43
							if v44 <= l0 {
								m.G0 = v7 + int32(64)
								return
							} else {
								if v44 <= int32(0) {
									F__serverAssert(m, int32(_a852), int32(_a846), int32(387))
									mBase = m.M
									v150 = m.ExcPending
									if v150 != 0 {
										return
									} else {
										F_abort(m)
										mBase = m.M
										base.Wasm_trap_unreachable()
										for {
										}
									}
								} else {
									if l0 <= int32(0) {
										F__serverAssert(m, int32(_a853), int32(_a846), int32(388))
										mBase = m.M
										v156 = m.ExcPending
										if v156 != 0 {
											return
										} else {
											F_abort(m)
											mBase = m.M
											base.Wasm_trap_unreachable()
											for {
											}
										}
									} else {
										F_spscInit(m, l0*int32(192)+int32(_a854), int32(4096))
										mBase = m.M
										v60 = m.ExcPending
										if v60 != 0 {
											return
										} else {
											v68 = int32(16)
											v69 = v7 + v68
											v73 = m.G0
											v75 = v73 - v68
											m.G0 = v75
											v77 = F_pthread_attr_init(m, v69)
											mBase = m.M
											v80 = F_pthread_attr_getstacksize(m, v69, v75+int32(12))
											mBase = m.M
											v81 = *(*int32)(unsafe.Add(mBase, uint32(v75)+12))
											v82 = int32(1)
											if base.Ui32(v82) < base.Ui32(v81) {
												v85 = v81
											} else {
												v85 = v82
											}
											v89 = v85
											for {
												if base.Ui32(v89) < base.Ui32(int32(4194304)) {
													v89 = v89 << (uint(int32(1)) % 32)
													continue
												} else {
													break
												}
												break
											}
											*(*int32)(unsafe.Add(mBase, uint32(v75)+12)) = v89
											v95 = F_pthread_attr_setstacksize(m, v69, v89)
											mBase = m.M
											m.G0 = v75 + int32(16)
											v104 = int32(6)
											v111 = *(*int32)(unsafe.Add(mBase, _consts[6]))
											if int32(3) < v111 {
												m.Env.Exit(m, int32(1))
												mBase = m.M
												base.Wasm_trap_unreachable()
												for {
												}
											} else {
												v114 = F___strerror_l(m, v104, v104)
												mBase = m.M
												*(*int32)(unsafe.Add(mBase, uint32(v7))) = v114
												F__serverLog(m, int32(3), int32(_a855), v7)
												mBase = m.M
												v119 = m.ExcPending
												if v119 != 0 {
													return
												} else {
													m.Env.Exit(m, int32(1))
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
				}
			}
		}
	}
}
func F_processIOThreadsResponses(m *base.Module) int32 {
	mBase := m.M
	_ = mBase
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v18 int64
	_ = v18
	var v20 int64
	_ = v20
	var v24 int32
	_ = v24
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v109 int32
	_ = v109
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v121 int32
	_ = v121
	var v132 int32
	_ = v132
	var v134 int32
	_ = v134
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v148 int32
	_ = v148
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v162 int32
	_ = v162
	var v170 int32
	_ = v170
	var v176 int32
	_ = v176
	var v183 int32
	_ = v183
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v215 int32
	_ = v215
	var v217 int64
	_ = v217
	var v218 int64
	_ = v218
	var v219 int64
	_ = v219
	var v223 int32
	_ = v223
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
	var v245 int32
	_ = v245
	var v246 int64
	_ = v246
	var v247 int32
	_ = v247
	var v248 int32
	_ = v248
	var v259 int32
	_ = v259
	var v261 int32
	_ = v261
	var v270 int32
	_ = v270
	var v275 int32
	_ = v275
	var v277 int64
	_ = v277
	var v281 int32
	_ = v281
	var v290 int32
	_ = v290
	var v301 int64
	_ = v301
	var v304 int32
	_ = v304
	var v305 int32
	_ = v305
	var v306 int32
	_ = v306
	var v308 int64
	_ = v308
	var v310 int64
	_ = v310
	var v312 int64
	_ = v312
	var v315 int64
	_ = v315
	var v317 int64
	_ = v317
	var v319 int64
	_ = v319
	var v321 int64
	_ = v321
	var v347 int32
	_ = v347
	var v348 int32
	_ = v348
	var v353 int32
	_ = v353
	var v354 int32
	_ = v354
	var v360 int32
	_ = v360
	var v363 int32
	_ = v363
	var v364 int32
	_ = v364
	var v366 int32
	_ = v366
	var v369 int32
	_ = v369
	var v370 int32
	_ = v370
	var v371 int32
	_ = v371
	var v373 int64
	_ = v373
	var v375 int64
	_ = v375
	var v377 int64
	_ = v377
	var v380 int64
	_ = v380
	var v382 int64
	_ = v382
	var v384 int64
	_ = v384
	var v386 int64
	_ = v386
	var v412 int32
	_ = v412
	var v413 int32
	_ = v413
	var v418 int32
	_ = v418
	var v419 int32
	_ = v419
	var v425 int32
	_ = v425
	var v429 int32
	_ = v429
	var v431 int32
	_ = v431
	var v434 int32
	_ = v434
	var v435 int32
	_ = v435
	var v436 int32
	_ = v436
	var v437 int32
	_ = v437
	var v440 int32
	_ = v440
	var v443 int32
	_ = v443
	var v447 int32
	_ = v447
	var v449 int32
	_ = v449
	var v450 int32
	_ = v450
	var v451 int32
	_ = v451
	var v455 int32
	_ = v455
	var v461 int32
	_ = v461
	var v477 int32
	_ = v477
	var v479 int64
	_ = v479
	var v481 int64
	_ = v481
	var v485 int32
	_ = v485
	var v494 int32
	_ = v494
	var v500 int32
	_ = v500
	var v502 int64
	_ = v502
	var v511 int32
	_ = v511
	var v513 int32
	_ = v513
	var v515 int32
	_ = v515
	var v529 int32
	_ = v529
	var v533 int32
	_ = v533
	var v552 int32
	_ = v552
	var v558 int32
	_ = v558
	v13 = m.G0
	v15 = v13 - int32(336)
	m.G0 = v15
	v17 = int32(_a44)
	v18 = *(*int64)(unsafe.Add(mBase, _consts[486]))
	v20 = *(*int64)(unsafe.Add(mBase, _consts[488]))
	if base.I32_wrap_i64(v18+v20) != 0 {
		goto L4
	} else {
		goto L5
	}
L1:
	;
	F__serverAssert(m, int32(_a858), int32(_a846), int32(889))
	mBase = m.M
	v558 = m.ExcPending
	if v558 != 0 {
		goto L39
	} else {
		goto L91
	}
L2:
	;
	F__serverAssert(m, int32(_a859), int32(_a846), int32(856))
	mBase = m.M
	v552 = m.ExcPending
	if v552 != 0 {
		goto L39
	} else {
		goto L90
	}
L3:
	;
	m.G0 = v15 + int32(336)
	return v533
L4:
	;
	v24 = int32(0)
	v29 = v24
	v30 = v24
	v31 = v24
	v32 = v24
	goto L6
L5:
	;
	v533 = int32(0)
	goto L3
L6:
	;
	v40 = int32(_a856)
	v44 = int32(16) - v32
	v49 = *(*int32)(unsafe.Add(mBase, _consts[489]))
	v50 = *(*int32)(unsafe.Add(mBase, _consts[490]))
	if v49 != v50 {
		v55 = v50
		goto L12
	} else {
		goto L13
	}
L7:
	;
	v533 = v202
	goto L3
L8:
	;
	if v204 == int32(0) {
		goto L44
	} else {
		goto L45
	}
L9:
	;
	if v109 == int32(0) {
		v202 = v29
		v203 = v30
		v204 = v31
		goto L8
	} else {
		goto L25
	}
L10:
	;
	*(*int32)(unsafe.Add(mBase, _consts[489])) = v99
	v109 = v100
	goto L9
L11:
	;
	v109 = int32(0)
	goto L9
L12:
	;
	v56 = v55 - v49
	if base.Ui32(v56) < base.Ui32(v44) {
		goto L15
	} else {
		goto L16
	}
L13:
	;
	v52 = *(*int32)(unsafe.Add(mBase, _consts[491]))
	*(*int32)(unsafe.Add(mBase, _consts[490])) = v52
	if v49 == v52 {
		goto L11
	} else {
		goto L14
	}
L14:
	;
	v55 = v52
	goto L12
L15:
	;
	v58 = v56
	goto L17
L16:
	;
	v58 = v44
	goto L17
L17:
	;
	if v58 == int32(0) {
		goto L11
	} else {
		goto L18
	}
L18:
	;
	v65 = v49
	v66 = int32(0)
	goto L20
L19:
	;
	if v66 != 0 {
		v99 = v65
		v100 = v66
		goto L10
	} else {
		goto L24
	}
L20:
	;
	v69 = *(*int32)(unsafe.Add(mBase, _consts[492]))
	v70 = *(*int32)(unsafe.Add(mBase, _consts[493]))
	v76 = v69 + (v70+int32(-1))&v65<<(uint(int32(2))%32)
	v77 = *(*int32)(unsafe.Add(mBase, uint32(v76)))
	if v77 == int32(0) {
		goto L19
	} else {
		goto L22
	}
L21:
	;
	v99 = v87
	v100 = v58
	goto L10
L22:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15+int32(144)+v66<<(uint(int32(2))%32)))) = v77
	*(*int32)(unsafe.Add(mBase, uint32(v76))) = int32(0)
	v86 = int32(1)
	v87 = v65 + v86
	v89 = v66 + v86
	if v89 != v58 {
		v65 = v87
		v66 = v89
		goto L20
	} else {
		goto L23
	}
L23:
	;
	goto L21
L24:
	;
	goto L11
L25:
	;
	if v109 < int32(1) {
		v187 = v30
		v188 = v31
		goto L26
	} else {
		goto L27
	}
L26:
	;
	v197 = v109 + v29
	v198 = v109 + v32
	if v198 <= int32(15) {
		v29 = v197
		v30 = v187
		v31 = v188
		v32 = v198
		goto L6
	} else {
		goto L43
	}
L27:
	;
	v117 = v30
	v118 = v31
	v121 = int32(0)
	goto L31
L28:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15))) = v136
	F__serverPanic_1(m, int32(_a846), int32(938), int32(_a860), v15)
	mBase = m.M
	v183 = m.ExcPending
	if v183 != 0 {
		goto L39
	} else {
		goto L42
	}
L29:
	;
	F__serverAssert(m, int32(_a861), int32(_a846), int32(935))
	mBase = m.M
	v176 = m.ExcPending
	if v176 != 0 {
		goto L39
	} else {
		goto L41
	}
L30:
	;
	F__serverAssert(m, int32(_a862), int32(_a846), int32(932))
	mBase = m.M
	v170 = m.ExcPending
	if v170 != 0 {
		goto L39
	} else {
		goto L40
	}
L31:
	;
	v132 = *(*int32)(unsafe.Add(mBase, uint32(v15+int32(144)+v121<<(uint(int32(2))%32))))
	v134 = v132 & int32(-8)
	v136 = v132 & int32(7)
	switch v136 {
	case 0:
		goto L35
	case 1:
		goto L34
	default:
		goto L28
	}
L33:
	;
	v162 = v121 + int32(1)
	if v162 == v109 {
		v187 = v159
		v188 = v160
		goto L26
	} else {
		goto L38
	}
L34:
	;
	v148 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v134)+223)))
	if v148 != int32(2) {
		goto L29
	} else {
		goto L37
	}
L35:
	;
	v137 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v134)+222)))
	if v137 != int32(2) {
		goto L30
	} else {
		goto L36
	}
L36:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15+int32(80)+v118<<(uint(int32(2))%32)))) = v134
	v159 = v117
	v160 = v118 + int32(1)
	goto L33
L37:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15+int32(16)+v117<<(uint(int32(2))%32)))) = v134
	v159 = v117 + int32(1)
	v160 = v118
	goto L33
L38:
	;
	v117 = v159
	v118 = v160
	v121 = v162
	goto L31
L39:
	;
	return int32(0)
L40:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L41:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L42:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L43:
	;
	v202 = v197
	v203 = v187
	v204 = v188
	goto L8
L44:
	;
	if v203 == int32(0) {
		goto L81
	} else {
		goto L82
	}
L45:
	;
	v215 = int32(_a44)
	v217 = *(*int64)(unsafe.Add(mBase, _consts[486]))
	v218 = base.I64_extend_i32_s(v204)
	v219 = v217 - v218
	*(*int64)(unsafe.Add(mBase, _consts[486])) = v219
	if v219 <= int64(-1) {
		goto L2
	} else {
		goto L46
	}
L46:
	;
	v223 = int32(0)
	if v204 < int32(1) {
		v270 = v223
		goto L47
	} else {
		goto L48
	}
L47:
	;
	v275 = int32(_a44)
	v277 = *(*int64)(unsafe.Add(mBase, _consts[308]))
	*(*int64)(unsafe.Add(mBase, _consts[308])) = v277 + v218
	F_processClientsCommandsBatch(m)
	mBase = m.M
	v281 = m.ExcPending
	if v281 != 0 {
		goto L39
	} else {
		goto L55
	}
L48:
	;
	v235 = v223
	v236 = v223
	goto L49
L49:
	;
	v245 = *(*int32)(unsafe.Add(mBase, uint32(v15+int32(80)+v236<<(uint(int32(2))%32))))
	v246 = *(*int64)(unsafe.Add(mBase, uint32(v245)))
	v247 = F_processClientIOReadsDone(m, v245)
	mBase = m.M
	v248 = m.ExcPending
	if v248 != 0 {
		goto L39
	} else {
		goto L52
	}
L50:
	;
	v270 = v259
	goto L47
L51:
	;
	v261 = v236 + int32(1)
	if v261 != v204 {
		v235 = v259
		v236 = v261
		goto L49
	} else {
		goto L54
	}
L52:
	;
	if v247 == int32(0) {
		v259 = v235
		goto L51
	} else {
		goto L53
	}
L53:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v15+int32(208)+v235<<(uint(int32(3))%32)))) = v246
	v259 = v235 + int32(1)
	goto L51
L54:
	;
	goto L50
L55:
	;
	if v270 < int32(1) {
		goto L44
	} else {
		goto L56
	}
L56:
	;
	v290 = v223
	goto L57
L57:
	;
	v301 = *(*int64)(unsafe.Add(mBase, uint32(v15+int32(208)+v290<<(uint(int32(3))%32))))
	v304 = m.G0
	v305 = int32(16)
	v306 = v304 - v305
	m.G0 = v306
	v308 = int64(56)
	v310 = int64(65280)
	v312 = int64(40)
	v315 = int64(16711680)
	v317 = int64(24)
	v319 = int64(4278190080)
	v321 = int64(8)
	*(*int64)(unsafe.Add(mBase, uint32(v306)+8)) = v301<<(uint(v308)%64) | v301&v310<<(uint(v312)%64) | (v301&v315<<(uint(v317)%64) | v301&v319<<(uint(v321)%64)) | (int64(base.Ui64(v301)>>(uint(v321)%64))&v319 | int64(base.Ui64(v301)>>(uint(v317)%64))&v315 | (int64(base.Ui64(v301)>>(uint(v312)%64))&v310 | int64(base.Ui64(v301)>>(uint(v308)%64))))
	*(*int32)(unsafe.Add(mBase, uint32(v306)+4)) = int32(0)
	v347 = *(*int32)(unsafe.Add(mBase, _consts[404]))
	v348 = int32(8)
	v353 = F_raxFind(m, v347, v306+v348, v348, v306+int32(4))
	mBase = m.M
	v354 = *(*int32)(unsafe.Add(mBase, uint32(v306)+4))
	m.G0 = v306 + v305
	goto L60
L58:
	;
	goto L44
L59:
	;
	v461 = v290 + int32(1)
	if v461 != v270 {
		v290 = v461
		goto L57
	} else {
		goto L80
	}
L60:
	;
	if v354 == int32(0) {
		goto L59
	} else {
		goto L61
	}
L61:
	;
	v360 = *(*int32)(unsafe.Add(mBase, uint32(v354)+8))
	if v360 == int32(0) {
		goto L59
	} else {
		goto L62
	}
L62:
	;
	v363 = F_processPendingCommandAndInputBuffer(m, v354)
	mBase = m.M
	v364 = m.ExcPending
	if v364 != 0 {
		goto L39
	} else {
		goto L64
	}
L63:
	;
	v369 = m.G0
	v370 = int32(16)
	v371 = v369 - v370
	m.G0 = v371
	v373 = int64(56)
	v375 = int64(65280)
	v377 = int64(40)
	v380 = int64(16711680)
	v382 = int64(24)
	v384 = int64(4278190080)
	v386 = int64(8)
	*(*int64)(unsafe.Add(mBase, uint32(v371)+8)) = v301<<(uint(v373)%64) | v301&v375<<(uint(v377)%64) | (v301&v380<<(uint(v382)%64) | v301&v384<<(uint(v386)%64)) | (int64(base.Ui64(v301)>>(uint(v386)%64))&v384 | int64(base.Ui64(v301)>>(uint(v382)%64))&v380 | (int64(base.Ui64(v301)>>(uint(v377)%64))&v375 | int64(base.Ui64(v301)>>(uint(v373)%64))))
	*(*int32)(unsafe.Add(mBase, uint32(v371)+4)) = int32(0)
	v412 = *(*int32)(unsafe.Add(mBase, _consts[404]))
	v413 = int32(8)
	v418 = F_raxFind(m, v412, v371+v413, v413, v371+int32(4))
	mBase = m.M
	v419 = *(*int32)(unsafe.Add(mBase, uint32(v371)+4))
	m.G0 = v371 + v370
	goto L67
L64:
	;
	if v363 != 0 {
		goto L63
	} else {
		goto L65
	}
L65:
	;
	F_beforeNextClient(m, v354)
	mBase = m.M
	v366 = m.ExcPending
	if v366 != 0 {
		goto L39
	} else {
		goto L66
	}
L66:
	;
	goto L63
L67:
	;
	if v419 == int32(0) {
		goto L59
	} else {
		goto L68
	}
L68:
	;
	v425 = *(*int32)(unsafe.Add(mBase, uint32(v419)+8))
	if v425 == int32(0) {
		goto L59
	} else {
		goto L69
	}
L69:
	;
	v429 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v419)+222)))
	v431 = base.B2i32(v429 != int32(0))
	v434 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v419)+223)))
	if v434 != 0 {
		goto L71
	} else {
		goto L72
	}
L70:
	;
	v440 = *(*int32)(unsafe.Add(mBase, uint32(v425)))
	if v440 == int32(0) {
		goto L74
	} else {
		goto L75
	}
L71:
	;
	v435 = v431 | int32(2)
	goto L73
L72:
	;
	v435 = v431
	goto L73
L73:
	;
	v436 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v419)+62)))
	v437 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v419)+60)))
	goto L70
L74:
	;
	v449 = *(*int32)(unsafe.Add(mBase, uint32(v419)+8))
	v450 = *(*int32)(unsafe.Add(mBase, uint32(v449)))
	v451 = *(*int32)(unsafe.Add(mBase, uint32(v450)+116))
	if v451 == int32(0) {
		goto L59
	} else {
		goto L78
	}
L75:
	;
	v443 = *(*int32)(unsafe.Add(mBase, uint32(v440)+112))
	if v443 == int32(0) {
		goto L74
	} else {
		goto L76
	}
L76:
	;
	m.T0[v443].(func(*base.Module, int32, int32))(m, v425, v435|base.B2i32(base.Ui32(v436) < base.Ui32(v437)))
	mBase = m.M
	v447 = m.ExcPending
	if v447 != 0 {
		goto L39
	} else {
		goto L77
	}
L77:
	;
	goto L74
L78:
	;
	m.T0[v451].(func(*base.Module, int32))(m, v449)
	mBase = m.M
	v455 = m.ExcPending
	if v455 != 0 {
		goto L39
	} else {
		goto L79
	}
L79:
	;
	goto L59
L80:
	;
	goto L58
L81:
	;
	v529 = int32(0)
	if v109 != 0 {
		v29 = v202
		v30 = v529
		v31 = v529
		v32 = v529
		goto L6
	} else {
		goto L89
	}
L82:
	;
	v477 = int32(_a44)
	v479 = *(*int64)(unsafe.Add(mBase, _consts[488]))
	v481 = v479 - base.I64_extend_i32_s(v203)
	*(*int64)(unsafe.Add(mBase, _consts[488])) = v481
	if v481 < int64(0) {
		goto L1
	} else {
		goto L83
	}
L83:
	;
	v485 = int32(0)
	if v203 <= v485 {
		goto L81
	} else {
		goto L84
	}
L84:
	;
	v494 = v485
	goto L85
L85:
	;
	v500 = int32(_a44)
	v502 = *(*int64)(unsafe.Add(mBase, _consts[297]))
	*(*int64)(unsafe.Add(mBase, _consts[297])) = v502 + int64(1)
	v511 = *(*int32)(unsafe.Add(mBase, uint32(v15+int32(16)+v494<<(uint(int32(2))%32))))
	F_processClientIOWriteDone(m, v511)
	mBase = m.M
	v513 = m.ExcPending
	if v513 != 0 {
		goto L39
	} else {
		goto L87
	}
L86:
	;
	goto L81
L87:
	;
	v515 = v494 + int32(1)
	if v515 != v203 {
		v494 = v515
		goto L85
	} else {
		goto L88
	}
L88:
	;
	goto L86
L89:
	;
	goto L7
L90:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L91:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
}
