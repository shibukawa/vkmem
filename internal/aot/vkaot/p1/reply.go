package p1

import (
	base "github.com/shibukawa/valkeymem/internal/aot/vkaot/base"
	"unsafe"
)

func F__addReplyPayloadToList(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v26 int32
	_ = v26
	var v29 int64
	_ = v29
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v60 int32
	_ = v60
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v84 int32
	_ = v84
	var v132 int32
	_ = v132
	var v135 int32
	_ = v135
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v149 int32
	_ = v149
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v154 int64
	_ = v154
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v163 int32
	_ = v163
	var v166 int32
	_ = v166
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v176 int32
	_ = v176
	var v180 int32
	_ = v180
	var v182 int32
	_ = v182
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v199 int32
	_ = v199
	var v204 int32
	_ = v204
	var v206 int32
	_ = v206
	var v209 int32
	_ = v209
	var v211 int32
	_ = v211
	var v213 int32
	_ = v213
	var v215 int32
	_ = v215
	var v220 int32
	_ = v220
	var v222 int32
	_ = v222
	var v225 int32
	_ = v225
	var v227 int32
	_ = v227
	var v229 int32
	_ = v229
	var v236 int32
	_ = v236
	var v247 int32
	_ = v247
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v253 int32
	_ = v253
	var v259 int32
	_ = v259
	var v265 int32
	_ = v265
	var v266 int32
	_ = v266
	var v273 int32
	_ = v273
	var v276 int64
	_ = v276
	var v279 int32
	_ = v279
	var v281 int32
	_ = v281
	var v286 int32
	_ = v286
	var v287 int32
	_ = v287
	var v288 int32
	_ = v288
	var v292 int32
	_ = v292
	var v294 int32
	_ = v294
	var v297 int32
	_ = v297
	var v299 int32
	_ = v299
	var v305 int32
	_ = v305
	var v307 int32
	_ = v307
	var v308 int32
	_ = v308
	var v310 int64
	_ = v310
	var v312 int32
	_ = v312
	var v319 int32
	_ = v319
	var v322 int32
	_ = v322
	var v325 int32
	_ = v325
	var v326 int32
	_ = v326
	var v328 int32
	_ = v328
	var v329 int32
	_ = v329
	var v332 int32
	_ = v332
	var v336 int32
	_ = v336
	var v338 int32
	_ = v338
	var v343 int32
	_ = v343
	var v344 int32
	_ = v344
	var v345 int32
	_ = v345
	var v352 int32
	_ = v352
	var v353 int32
	_ = v353
	var v355 int32
	_ = v355
	var v360 int32
	_ = v360
	var v362 int32
	_ = v362
	var v365 int32
	_ = v365
	var v367 int32
	_ = v367
	var v369 int32
	_ = v369
	var v371 int32
	_ = v371
	var v376 int32
	_ = v376
	var v378 int32
	_ = v378
	var v381 int32
	_ = v381
	var v383 int32
	_ = v383
	var v385 int32
	_ = v385
	var v404 int32
	_ = v404
	var v405 int32
	_ = v405
	var v411 int32
	_ = v411
	var v415 int32
	_ = v415
	var v416 int32
	_ = v416
	var v419 int64
	_ = v419
	var v422 int32
	_ = v422
	var v423 int32
	_ = v423
	var v424 int64
	_ = v424
	var v425 int64
	_ = v425
	var v429 int32
	_ = v429
	var v430 int32
	_ = v430
	v11 = m.G0
	v13 = v11 - int32(16)
	m.G0 = v13
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v15 != 0 {
		v17 = *(*int32)(unsafe.Add(mBase, uint32(v15)+8))
		v18 = v17
	} else {
		v18 = int32(0)
	}
	v19 = int32(1)
	if l4 == v19 {
		v135 = v19
	} else {
		v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+204))
		if v26&int32(268435456) != 0 {
			v132 = int32(0)
		} else {
			v29 = *(*int64)(unsafe.Add(mBase, uint32(l0)+384))
			if v29 != int64(-1) {
				v132 = int32(0)
			} else {
				if v26&int32(131072) == int32(0) {
					v60 = *(*int32)(unsafe.Add(mBase, uint32(l0)+200))
					if v60&int32(1) != 0 {
						v132 = int32(0)
					} else {
						if v60&int32(2) == int32(0) {
							if v60&int32(262144) != 0 {
								v79 = int32(_a20)
								v80 = *(*int32)(unsafe.Add(mBase, _consts[479]))
								v84 = *(*int32)(unsafe.Add(mBase, _consts[345]))
								v132 = base.B2i32(v80 != int32(0)) & base.B2i32(v80 <= v84)
							} else {
								v73 = *(*int32)(unsafe.Add(mBase, uint32(l0)+216))
								if v73 == int32(0) {
									v79 = int32(_a20)
									v80 = *(*int32)(unsafe.Add(mBase, _consts[479]))
									v84 = *(*int32)(unsafe.Add(mBase, _consts[345]))
									v132 = base.B2i32(v80 != int32(0)) & base.B2i32(v80 <= v84)
								} else {
									v76 = F_isImportSlotMigrationJob(m, v73)
									mBase = m.M
									v132 = int32(0)
								}
							}
						} else {
							if v60&int32(4) == int32(0) {
								v132 = int32(0)
							} else {
								if v60&int32(262144) != 0 {
									v79 = int32(_a20)
									v80 = *(*int32)(unsafe.Add(mBase, _consts[479]))
									v84 = *(*int32)(unsafe.Add(mBase, _consts[345]))
									v132 = base.B2i32(v80 != int32(0)) & base.B2i32(v80 <= v84)
								} else {
									v73 = *(*int32)(unsafe.Add(mBase, uint32(l0)+216))
									if v73 == int32(0) {
										v79 = int32(_a20)
										v80 = *(*int32)(unsafe.Add(mBase, _consts[479]))
										v84 = *(*int32)(unsafe.Add(mBase, _consts[345]))
										v132 = base.B2i32(v80 != int32(0)) & base.B2i32(v80 <= v84)
									} else {
										v76 = F_isImportSlotMigrationJob(m, v73)
										mBase = m.M
										v132 = int32(0)
									}
								}
							}
						}
					}
				} else {
					v37 = *(*int32)(unsafe.Add(mBase, _consts[14]))
					if l0 != v37 {
						v60 = *(*int32)(unsafe.Add(mBase, uint32(l0)+200))
						if v60&int32(1) != 0 {
							v132 = int32(0)
						} else {
							if v60&int32(2) == int32(0) {
								if v60&int32(262144) != 0 {
									v79 = int32(_a20)
									v80 = *(*int32)(unsafe.Add(mBase, _consts[479]))
									v84 = *(*int32)(unsafe.Add(mBase, _consts[345]))
									v132 = base.B2i32(v80 != int32(0)) & base.B2i32(v80 <= v84)
								} else {
									v73 = *(*int32)(unsafe.Add(mBase, uint32(l0)+216))
									if v73 == int32(0) {
										v79 = int32(_a20)
										v80 = *(*int32)(unsafe.Add(mBase, _consts[479]))
										v84 = *(*int32)(unsafe.Add(mBase, _consts[345]))
										v132 = base.B2i32(v80 != int32(0)) & base.B2i32(v80 <= v84)
									} else {
										v76 = F_isImportSlotMigrationJob(m, v73)
										mBase = m.M
										v132 = int32(0)
									}
								}
							} else {
								if v60&int32(4) == int32(0) {
									v132 = int32(0)
								} else {
									if v60&int32(262144) != 0 {
										v79 = int32(_a20)
										v80 = *(*int32)(unsafe.Add(mBase, _consts[479]))
										v84 = *(*int32)(unsafe.Add(mBase, _consts[345]))
										v132 = base.B2i32(v80 != int32(0)) & base.B2i32(v80 <= v84)
									} else {
										v73 = *(*int32)(unsafe.Add(mBase, uint32(l0)+216))
										if v73 == int32(0) {
											v79 = int32(_a20)
											v80 = *(*int32)(unsafe.Add(mBase, _consts[479]))
											v84 = *(*int32)(unsafe.Add(mBase, _consts[345]))
											v132 = base.B2i32(v80 != int32(0)) & base.B2i32(v80 <= v84)
										} else {
											v76 = F_isImportSlotMigrationJob(m, v73)
											mBase = m.M
											v132 = int32(0)
										}
									}
								}
							}
						}
					} else {
						v40 = *(*int32)(unsafe.Add(mBase, _consts[300]))
						if v40 == int32(0) {
							v60 = *(*int32)(unsafe.Add(mBase, uint32(l0)+200))
							if v60&int32(1) != 0 {
								v132 = int32(0)
							} else {
								if v60&int32(2) == int32(0) {
									if v60&int32(262144) != 0 {
										v79 = int32(_a20)
										v80 = *(*int32)(unsafe.Add(mBase, _consts[479]))
										v84 = *(*int32)(unsafe.Add(mBase, _consts[345]))
										v132 = base.B2i32(v80 != int32(0)) & base.B2i32(v80 <= v84)
									} else {
										v73 = *(*int32)(unsafe.Add(mBase, uint32(l0)+216))
										if v73 == int32(0) {
											v79 = int32(_a20)
											v80 = *(*int32)(unsafe.Add(mBase, _consts[479]))
											v84 = *(*int32)(unsafe.Add(mBase, _consts[345]))
											v132 = base.B2i32(v80 != int32(0)) & base.B2i32(v80 <= v84)
										} else {
											v76 = F_isImportSlotMigrationJob(m, v73)
											mBase = m.M
											v132 = int32(0)
										}
									}
								} else {
									if v60&int32(4) == int32(0) {
										v132 = int32(0)
									} else {
										if v60&int32(262144) != 0 {
											v79 = int32(_a20)
											v80 = *(*int32)(unsafe.Add(mBase, _consts[479]))
											v84 = *(*int32)(unsafe.Add(mBase, _consts[345]))
											v132 = base.B2i32(v80 != int32(0)) & base.B2i32(v80 <= v84)
										} else {
											v73 = *(*int32)(unsafe.Add(mBase, uint32(l0)+216))
											if v73 == int32(0) {
												v79 = int32(_a20)
												v80 = *(*int32)(unsafe.Add(mBase, _consts[479]))
												v84 = *(*int32)(unsafe.Add(mBase, _consts[345]))
												v132 = base.B2i32(v80 != int32(0)) & base.B2i32(v80 <= v84)
											} else {
												v76 = F_isImportSlotMigrationJob(m, v73)
												mBase = m.M
												v132 = int32(0)
											}
										}
									}
								}
							}
						} else {
							v43 = *(*int32)(unsafe.Add(mBase, uint32(v40)+68))
							if v43 == int32(0) {
								v132 = int32(0)
							} else {
								v46 = *(*int32)(unsafe.Add(mBase, uint32(v43)+48))
								if v46 == int32(288) {
									v60 = *(*int32)(unsafe.Add(mBase, uint32(l0)+200))
									if v60&int32(1) != 0 {
										v132 = int32(0)
									} else {
										if v60&int32(2) == int32(0) {
											if v60&int32(262144) != 0 {
												v79 = int32(_a20)
												v80 = *(*int32)(unsafe.Add(mBase, _consts[479]))
												v84 = *(*int32)(unsafe.Add(mBase, _consts[345]))
												v132 = base.B2i32(v80 != int32(0)) & base.B2i32(v80 <= v84)
											} else {
												v73 = *(*int32)(unsafe.Add(mBase, uint32(l0)+216))
												if v73 == int32(0) {
													v79 = int32(_a20)
													v80 = *(*int32)(unsafe.Add(mBase, _consts[479]))
													v84 = *(*int32)(unsafe.Add(mBase, _consts[345]))
													v132 = base.B2i32(v80 != int32(0)) & base.B2i32(v80 <= v84)
												} else {
													v76 = F_isImportSlotMigrationJob(m, v73)
													mBase = m.M
													v132 = int32(0)
												}
											}
										} else {
											if v60&int32(4) == int32(0) {
												v132 = int32(0)
											} else {
												if v60&int32(262144) != 0 {
													v79 = int32(_a20)
													v80 = *(*int32)(unsafe.Add(mBase, _consts[479]))
													v84 = *(*int32)(unsafe.Add(mBase, _consts[345]))
													v132 = base.B2i32(v80 != int32(0)) & base.B2i32(v80 <= v84)
												} else {
													v73 = *(*int32)(unsafe.Add(mBase, uint32(l0)+216))
													if v73 == int32(0) {
														v79 = int32(_a20)
														v80 = *(*int32)(unsafe.Add(mBase, _consts[479]))
														v84 = *(*int32)(unsafe.Add(mBase, _consts[345]))
														v132 = base.B2i32(v80 != int32(0)) & base.B2i32(v80 <= v84)
													} else {
														v76 = F_isImportSlotMigrationJob(m, v73)
														mBase = m.M
														v132 = int32(0)
													}
												}
											}
										}
									}
								} else {
									if v46 == int32(286) {
										v60 = *(*int32)(unsafe.Add(mBase, uint32(l0)+200))
										if v60&int32(1) != 0 {
											v132 = int32(0)
										} else {
											if v60&int32(2) == int32(0) {
												if v60&int32(262144) != 0 {
													v79 = int32(_a20)
													v80 = *(*int32)(unsafe.Add(mBase, _consts[479]))
													v84 = *(*int32)(unsafe.Add(mBase, _consts[345]))
													v132 = base.B2i32(v80 != int32(0)) & base.B2i32(v80 <= v84)
												} else {
													v73 = *(*int32)(unsafe.Add(mBase, uint32(l0)+216))
													if v73 == int32(0) {
														v79 = int32(_a20)
														v80 = *(*int32)(unsafe.Add(mBase, _consts[479]))
														v84 = *(*int32)(unsafe.Add(mBase, _consts[345]))
														v132 = base.B2i32(v80 != int32(0)) & base.B2i32(v80 <= v84)
													} else {
														v76 = F_isImportSlotMigrationJob(m, v73)
														mBase = m.M
														v132 = int32(0)
													}
												}
											} else {
												if v60&int32(4) == int32(0) {
													v132 = int32(0)
												} else {
													if v60&int32(262144) != 0 {
														v79 = int32(_a20)
														v80 = *(*int32)(unsafe.Add(mBase, _consts[479]))
														v84 = *(*int32)(unsafe.Add(mBase, _consts[345]))
														v132 = base.B2i32(v80 != int32(0)) & base.B2i32(v80 <= v84)
													} else {
														v73 = *(*int32)(unsafe.Add(mBase, uint32(l0)+216))
														if v73 == int32(0) {
															v79 = int32(_a20)
															v80 = *(*int32)(unsafe.Add(mBase, _consts[479]))
															v84 = *(*int32)(unsafe.Add(mBase, _consts[345]))
															v132 = base.B2i32(v80 != int32(0)) & base.B2i32(v80 <= v84)
														} else {
															v76 = F_isImportSlotMigrationJob(m, v73)
															mBase = m.M
															v132 = int32(0)
														}
													}
												}
											}
										}
									} else {
										if v46 == int32(284) {
											v60 = *(*int32)(unsafe.Add(mBase, uint32(l0)+200))
											if v60&int32(1) != 0 {
												v132 = int32(0)
											} else {
												if v60&int32(2) == int32(0) {
													if v60&int32(262144) != 0 {
														v79 = int32(_a20)
														v80 = *(*int32)(unsafe.Add(mBase, _consts[479]))
														v84 = *(*int32)(unsafe.Add(mBase, _consts[345]))
														v132 = base.B2i32(v80 != int32(0)) & base.B2i32(v80 <= v84)
													} else {
														v73 = *(*int32)(unsafe.Add(mBase, uint32(l0)+216))
														if v73 == int32(0) {
															v79 = int32(_a20)
															v80 = *(*int32)(unsafe.Add(mBase, _consts[479]))
															v84 = *(*int32)(unsafe.Add(mBase, _consts[345]))
															v132 = base.B2i32(v80 != int32(0)) & base.B2i32(v80 <= v84)
														} else {
															v76 = F_isImportSlotMigrationJob(m, v73)
															mBase = m.M
															v132 = int32(0)
														}
													}
												} else {
													if v60&int32(4) == int32(0) {
														v132 = int32(0)
													} else {
														if v60&int32(262144) != 0 {
															v79 = int32(_a20)
															v80 = *(*int32)(unsafe.Add(mBase, _consts[479]))
															v84 = *(*int32)(unsafe.Add(mBase, _consts[345]))
															v132 = base.B2i32(v80 != int32(0)) & base.B2i32(v80 <= v84)
														} else {
															v73 = *(*int32)(unsafe.Add(mBase, uint32(l0)+216))
															if v73 == int32(0) {
																v79 = int32(_a20)
																v80 = *(*int32)(unsafe.Add(mBase, _consts[479]))
																v84 = *(*int32)(unsafe.Add(mBase, _consts[345]))
																v132 = base.B2i32(v80 != int32(0)) & base.B2i32(v80 <= v84)
															} else {
																v76 = F_isImportSlotMigrationJob(m, v73)
																mBase = m.M
																v132 = int32(0)
															}
														}
													}
												}
											}
										} else {
											if v46 == int32(282) {
												v60 = *(*int32)(unsafe.Add(mBase, uint32(l0)+200))
												if v60&int32(1) != 0 {
													v132 = int32(0)
												} else {
													if v60&int32(2) == int32(0) {
														if v60&int32(262144) != 0 {
															v79 = int32(_a20)
															v80 = *(*int32)(unsafe.Add(mBase, _consts[479]))
															v84 = *(*int32)(unsafe.Add(mBase, _consts[345]))
															v132 = base.B2i32(v80 != int32(0)) & base.B2i32(v80 <= v84)
														} else {
															v73 = *(*int32)(unsafe.Add(mBase, uint32(l0)+216))
															if v73 == int32(0) {
																v79 = int32(_a20)
																v80 = *(*int32)(unsafe.Add(mBase, _consts[479]))
																v84 = *(*int32)(unsafe.Add(mBase, _consts[345]))
																v132 = base.B2i32(v80 != int32(0)) & base.B2i32(v80 <= v84)
															} else {
																v76 = F_isImportSlotMigrationJob(m, v73)
																mBase = m.M
																v132 = int32(0)
															}
														}
													} else {
														if v60&int32(4) == int32(0) {
															v132 = int32(0)
														} else {
															if v60&int32(262144) != 0 {
																v79 = int32(_a20)
																v80 = *(*int32)(unsafe.Add(mBase, _consts[479]))
																v84 = *(*int32)(unsafe.Add(mBase, _consts[345]))
																v132 = base.B2i32(v80 != int32(0)) & base.B2i32(v80 <= v84)
															} else {
																v73 = *(*int32)(unsafe.Add(mBase, uint32(l0)+216))
																if v73 == int32(0) {
																	v79 = int32(_a20)
																	v80 = *(*int32)(unsafe.Add(mBase, _consts[479]))
																	v84 = *(*int32)(unsafe.Add(mBase, _consts[345]))
																	v132 = base.B2i32(v80 != int32(0)) & base.B2i32(v80 <= v84)
																} else {
																	v76 = F_isImportSlotMigrationJob(m, v73)
																	mBase = m.M
																	v132 = int32(0)
																}
															}
														}
													}
												}
											} else {
												if v46 == int32(287) {
													v60 = *(*int32)(unsafe.Add(mBase, uint32(l0)+200))
													if v60&int32(1) != 0 {
														v132 = int32(0)
													} else {
														if v60&int32(2) == int32(0) {
															if v60&int32(262144) != 0 {
																v79 = int32(_a20)
																v80 = *(*int32)(unsafe.Add(mBase, _consts[479]))
																v84 = *(*int32)(unsafe.Add(mBase, _consts[345]))
																v132 = base.B2i32(v80 != int32(0)) & base.B2i32(v80 <= v84)
															} else {
																v73 = *(*int32)(unsafe.Add(mBase, uint32(l0)+216))
																if v73 == int32(0) {
																	v79 = int32(_a20)
																	v80 = *(*int32)(unsafe.Add(mBase, _consts[479]))
																	v84 = *(*int32)(unsafe.Add(mBase, _consts[345]))
																	v132 = base.B2i32(v80 != int32(0)) & base.B2i32(v80 <= v84)
																} else {
																	v76 = F_isImportSlotMigrationJob(m, v73)
																	mBase = m.M
																	v132 = int32(0)
																}
															}
														} else {
															if v60&int32(4) == int32(0) {
																v132 = int32(0)
															} else {
																if v60&int32(262144) != 0 {
																	v79 = int32(_a20)
																	v80 = *(*int32)(unsafe.Add(mBase, _consts[479]))
																	v84 = *(*int32)(unsafe.Add(mBase, _consts[345]))
																	v132 = base.B2i32(v80 != int32(0)) & base.B2i32(v80 <= v84)
																} else {
																	v73 = *(*int32)(unsafe.Add(mBase, uint32(l0)+216))
																	if v73 == int32(0) {
																		v79 = int32(_a20)
																		v80 = *(*int32)(unsafe.Add(mBase, _consts[479]))
																		v84 = *(*int32)(unsafe.Add(mBase, _consts[345]))
																		v132 = base.B2i32(v80 != int32(0)) & base.B2i32(v80 <= v84)
																	} else {
																		v76 = F_isImportSlotMigrationJob(m, v73)
																		mBase = m.M
																		v132 = int32(0)
																	}
																}
															}
														}
													}
												} else {
													if v46 != int32(289) {
														v132 = int32(0)
													} else {
														v60 = *(*int32)(unsafe.Add(mBase, uint32(l0)+200))
														if v60&int32(1) != 0 {
															v132 = int32(0)
														} else {
															if v60&int32(2) == int32(0) {
																if v60&int32(262144) != 0 {
																	v79 = int32(_a20)
																	v80 = *(*int32)(unsafe.Add(mBase, _consts[479]))
																	v84 = *(*int32)(unsafe.Add(mBase, _consts[345]))
																	v132 = base.B2i32(v80 != int32(0)) & base.B2i32(v80 <= v84)
																} else {
																	v73 = *(*int32)(unsafe.Add(mBase, uint32(l0)+216))
																	if v73 == int32(0) {
																		v79 = int32(_a20)
																		v80 = *(*int32)(unsafe.Add(mBase, _consts[479]))
																		v84 = *(*int32)(unsafe.Add(mBase, _consts[345]))
																		v132 = base.B2i32(v80 != int32(0)) & base.B2i32(v80 <= v84)
																	} else {
																		v76 = F_isImportSlotMigrationJob(m, v73)
																		mBase = m.M
																		v132 = int32(0)
																	}
																}
															} else {
																if v60&int32(4) == int32(0) {
																	v132 = int32(0)
																} else {
																	if v60&int32(262144) != 0 {
																		v79 = int32(_a20)
																		v80 = *(*int32)(unsafe.Add(mBase, _consts[479]))
																		v84 = *(*int32)(unsafe.Add(mBase, _consts[345]))
																		v132 = base.B2i32(v80 != int32(0)) & base.B2i32(v80 <= v84)
																	} else {
																		v73 = *(*int32)(unsafe.Add(mBase, uint32(l0)+216))
																		if v73 == int32(0) {
																			v79 = int32(_a20)
																			v80 = *(*int32)(unsafe.Add(mBase, _consts[479]))
																			v84 = *(*int32)(unsafe.Add(mBase, _consts[345]))
																			v132 = base.B2i32(v80 != int32(0)) & base.B2i32(v80 <= v84)
																		} else {
																			v76 = F_isImportSlotMigrationJob(m, v73)
																			mBase = m.M
																			v132 = int32(0)
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
		v135 = base.B2i32(v132 != int32(0))
	}
	if v18 == int32(0) {
		v265 = l2
		v266 = l3
	} else {
		v138 = *(*int32)(unsafe.Add(mBase, uint32(v18)))
		v139 = *(*int32)(unsafe.Add(mBase, uint32(v18)+4))
		v140 = v138 - v139
		v141 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18)+12)))
		if v141&int32(1) == int32(0) {
			if v135 != 0 {
				v265 = l2
				v266 = l3
			} else {
				if base.Ui32(v140) < base.Ui32(l3) {
					v249 = v140
				} else {
					v249 = l3
				}
				v250 = v249
				if v250 == int32(0) {
					v265 = l2
					v266 = l3
				} else {
					v253 = *(*int32)(unsafe.Add(mBase, uint32(v18)+4))
					if v250 == int32(0) {
					} else {
						v259 = F__emscripten_memcpy_bulkmem(m, v18+v253+int32(13), l2, v250)
						mBase = m.M
					}
					*(*int32)(unsafe.Add(mBase, uint32(v18)+4)) = v253 + v250
					v265 = l2 + v250
					v266 = l3 - v250
				}
			}
		} else {
			v149 = v18 + int32(4)
			v151 = v18 + int32(8)
			v152 = *(*int32)(unsafe.Add(mBase, uint32(l0)+292))
			v154 = *(*int64)(unsafe.Add(mBase, _consts[480]))
			v156 = base.B2i32(v154 != int64(-1))
			v157 = int32(0)
			v163 = int32(1)
			if l4 == v163 {
				v166 = l3
			} else {
				v166 = v163
			}
			if base.Ui32(v140) < base.Ui32(v166) {
				v236 = v157
				v247 = v236
			} else {
				v169 = F_clusterSlotStatsEnabled(m, v152)
				mBase = m.M
				if v169 != 0 {
					v170 = v152
				} else {
					v170 = int32(-1)
				}
				if base.Ui32(v140) < base.Ui32(l3) {
					v172 = v140
				} else {
					v172 = l3
				}
				v173 = *(*int32)(unsafe.Add(mBase, uint32(v151)))
				if v173 == int32(0) {
					if base.Ui32(v140) < base.Ui32(v166+int32(12)) {
						v236 = v157
					} else {
						v196 = *(*int32)(unsafe.Add(mBase, uint32(v149)))
						v197 = v18 + int32(13) + v196
						*(*int32)(unsafe.Add(mBase, uint32(v151))) = v197
						v199 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v197)+10)))
						v204 = v199&int32(254) | l4&int32(1)
						*(*uint8)(unsafe.Add(mBase, uint32(v197)+10)) = uint8(v204)
						v206 = *(*int32)(unsafe.Add(mBase, uint32(v151)))
						*(*uint16)(unsafe.Add(mBase, uint32(v206)+8)) = uint16(v170)
						v209 = v140 + int32(-12)
						if base.Ui32(v209) < base.Ui32(l3) {
							v211 = v209
						} else {
							v211 = v172
						}
						*(*int32)(unsafe.Add(mBase, uint32(v206))) = v211
						v213 = int32(0)
						*(*int32)(unsafe.Add(mBase, uint32(v206)+4)) = v213
						v215 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v206)+10)))
						v220 = v215&int32(253) | v156<<(uint(int32(1))%32)
						*(*uint8)(unsafe.Add(mBase, uint32(v206)+10)) = uint8(v220)
						v222 = *(*int32)(unsafe.Add(mBase, uint32(v151)))
						*(*uint8)(unsafe.Add(mBase, uint32(v222)+11)) = uint8(v213)
						v225 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v222)+10)))
						v227 = v225 & int32(3)
						*(*uint8)(unsafe.Add(mBase, uint32(v222)+10)) = uint8(v227)
						v229 = *(*int32)(unsafe.Add(mBase, uint32(v149)))
						*(*int32)(unsafe.Add(mBase, uint32(v149))) = v229 + int32(12)
						v236 = v211
					}
					v247 = v236
				} else {
					v176 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v173)+10)))
					if v176&int32(1) != l4 {
						if base.Ui32(v140) < base.Ui32(v166+int32(12)) {
							v236 = v157
						} else {
							v196 = *(*int32)(unsafe.Add(mBase, uint32(v149)))
							v197 = v18 + int32(13) + v196
							*(*int32)(unsafe.Add(mBase, uint32(v151))) = v197
							v199 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v197)+10)))
							v204 = v199&int32(254) | l4&int32(1)
							*(*uint8)(unsafe.Add(mBase, uint32(v197)+10)) = uint8(v204)
							v206 = *(*int32)(unsafe.Add(mBase, uint32(v151)))
							*(*uint16)(unsafe.Add(mBase, uint32(v206)+8)) = uint16(v170)
							v209 = v140 + int32(-12)
							if base.Ui32(v209) < base.Ui32(l3) {
								v211 = v209
							} else {
								v211 = v172
							}
							*(*int32)(unsafe.Add(mBase, uint32(v206))) = v211
							v213 = int32(0)
							*(*int32)(unsafe.Add(mBase, uint32(v206)+4)) = v213
							v215 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v206)+10)))
							v220 = v215&int32(253) | v156<<(uint(int32(1))%32)
							*(*uint8)(unsafe.Add(mBase, uint32(v206)+10)) = uint8(v220)
							v222 = *(*int32)(unsafe.Add(mBase, uint32(v151)))
							*(*uint8)(unsafe.Add(mBase, uint32(v222)+11)) = uint8(v213)
							v225 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v222)+10)))
							v227 = v225 & int32(3)
							*(*uint8)(unsafe.Add(mBase, uint32(v222)+10)) = uint8(v227)
							v229 = *(*int32)(unsafe.Add(mBase, uint32(v149)))
							*(*int32)(unsafe.Add(mBase, uint32(v149))) = v229 + int32(12)
							v236 = v211
						}
						v247 = v236
					} else {
						v180 = int32(*(*int16)(unsafe.Add(mBase, uint32(v173)+8)))
						if v170 != v180 {
							if base.Ui32(v140) < base.Ui32(v166+int32(12)) {
								v236 = v157
							} else {
								v196 = *(*int32)(unsafe.Add(mBase, uint32(v149)))
								v197 = v18 + int32(13) + v196
								*(*int32)(unsafe.Add(mBase, uint32(v151))) = v197
								v199 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v197)+10)))
								v204 = v199&int32(254) | l4&int32(1)
								*(*uint8)(unsafe.Add(mBase, uint32(v197)+10)) = uint8(v204)
								v206 = *(*int32)(unsafe.Add(mBase, uint32(v151)))
								*(*uint16)(unsafe.Add(mBase, uint32(v206)+8)) = uint16(v170)
								v209 = v140 + int32(-12)
								if base.Ui32(v209) < base.Ui32(l3) {
									v211 = v209
								} else {
									v211 = v172
								}
								*(*int32)(unsafe.Add(mBase, uint32(v206))) = v211
								v213 = int32(0)
								*(*int32)(unsafe.Add(mBase, uint32(v206)+4)) = v213
								v215 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v206)+10)))
								v220 = v215&int32(253) | v156<<(uint(int32(1))%32)
								*(*uint8)(unsafe.Add(mBase, uint32(v206)+10)) = uint8(v220)
								v222 = *(*int32)(unsafe.Add(mBase, uint32(v151)))
								*(*uint8)(unsafe.Add(mBase, uint32(v222)+11)) = uint8(v213)
								v225 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v222)+10)))
								v227 = v225 & int32(3)
								*(*uint8)(unsafe.Add(mBase, uint32(v222)+10)) = uint8(v227)
								v229 = *(*int32)(unsafe.Add(mBase, uint32(v149)))
								*(*int32)(unsafe.Add(mBase, uint32(v149))) = v229 + int32(12)
								v236 = v211
							}
							v247 = v236
						} else {
							v182 = int32(1)
							if v156 != int32(base.Ui32(v176)>>(uint(v182)%32))&v182 {
								if base.Ui32(v140) < base.Ui32(v166+int32(12)) {
									v236 = v157
								} else {
									v196 = *(*int32)(unsafe.Add(mBase, uint32(v149)))
									v197 = v18 + int32(13) + v196
									*(*int32)(unsafe.Add(mBase, uint32(v151))) = v197
									v199 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v197)+10)))
									v204 = v199&int32(254) | l4&int32(1)
									*(*uint8)(unsafe.Add(mBase, uint32(v197)+10)) = uint8(v204)
									v206 = *(*int32)(unsafe.Add(mBase, uint32(v151)))
									*(*uint16)(unsafe.Add(mBase, uint32(v206)+8)) = uint16(v170)
									v209 = v140 + int32(-12)
									if base.Ui32(v209) < base.Ui32(l3) {
										v211 = v209
									} else {
										v211 = v172
									}
									*(*int32)(unsafe.Add(mBase, uint32(v206))) = v211
									v213 = int32(0)
									*(*int32)(unsafe.Add(mBase, uint32(v206)+4)) = v213
									v215 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v206)+10)))
									v220 = v215&int32(253) | v156<<(uint(int32(1))%32)
									*(*uint8)(unsafe.Add(mBase, uint32(v206)+10)) = uint8(v220)
									v222 = *(*int32)(unsafe.Add(mBase, uint32(v151)))
									*(*uint8)(unsafe.Add(mBase, uint32(v222)+11)) = uint8(v213)
									v225 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v222)+10)))
									v227 = v225 & int32(3)
									*(*uint8)(unsafe.Add(mBase, uint32(v222)+10)) = uint8(v227)
									v229 = *(*int32)(unsafe.Add(mBase, uint32(v149)))
									*(*int32)(unsafe.Add(mBase, uint32(v149))) = v229 + int32(12)
									v236 = v211
								}
								v247 = v236
							} else {
								v187 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v173)+11)))
								if v187 != 0 {
									if base.Ui32(v140) < base.Ui32(v166+int32(12)) {
										v236 = v157
									} else {
										v196 = *(*int32)(unsafe.Add(mBase, uint32(v149)))
										v197 = v18 + int32(13) + v196
										*(*int32)(unsafe.Add(mBase, uint32(v151))) = v197
										v199 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v197)+10)))
										v204 = v199&int32(254) | l4&int32(1)
										*(*uint8)(unsafe.Add(mBase, uint32(v197)+10)) = uint8(v204)
										v206 = *(*int32)(unsafe.Add(mBase, uint32(v151)))
										*(*uint16)(unsafe.Add(mBase, uint32(v206)+8)) = uint16(v170)
										v209 = v140 + int32(-12)
										if base.Ui32(v209) < base.Ui32(l3) {
											v211 = v209
										} else {
											v211 = v172
										}
										*(*int32)(unsafe.Add(mBase, uint32(v206))) = v211
										v213 = int32(0)
										*(*int32)(unsafe.Add(mBase, uint32(v206)+4)) = v213
										v215 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v206)+10)))
										v220 = v215&int32(253) | v156<<(uint(int32(1))%32)
										*(*uint8)(unsafe.Add(mBase, uint32(v206)+10)) = uint8(v220)
										v222 = *(*int32)(unsafe.Add(mBase, uint32(v151)))
										*(*uint8)(unsafe.Add(mBase, uint32(v222)+11)) = uint8(v213)
										v225 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v222)+10)))
										v227 = v225 & int32(3)
										*(*uint8)(unsafe.Add(mBase, uint32(v222)+10)) = uint8(v227)
										v229 = *(*int32)(unsafe.Add(mBase, uint32(v149)))
										*(*int32)(unsafe.Add(mBase, uint32(v149))) = v229 + int32(12)
										v236 = v211
									}
									v247 = v236
								} else {
									v188 = *(*int32)(unsafe.Add(mBase, uint32(v151)))
									v189 = *(*int32)(unsafe.Add(mBase, uint32(v188)))
									*(*int32)(unsafe.Add(mBase, uint32(v188))) = v189 + v172
									v247 = v172
								}
							}
						}
					}
				}
			}
			v250 = v247
			if v250 == int32(0) {
				v265 = l2
				v266 = l3
			} else {
				v253 = *(*int32)(unsafe.Add(mBase, uint32(v18)+4))
				if v250 == int32(0) {
				} else {
					v259 = F__emscripten_memcpy_bulkmem(m, v18+v253+int32(13), l2, v250)
					mBase = m.M
				}
				*(*int32)(unsafe.Add(mBase, uint32(v18)+4)) = v253 + v250
				v265 = l2 + v250
				v266 = l3 - v250
			}
		}
	}
	if v266 == int32(0) {
		m.G0 = v13 + int32(16)
		return
	} else {
		if v135 != 0 {
			v273 = v266 + int32(12)
		} else {
			v273 = v266
		}
		v276 = *(*int64)(unsafe.Add(mBase, uint32(l0)+384))
		if v276 == int64(-1) {
			v279 = int32(16384)
		} else {
			v279 = int32(1024)
		}
		if base.Ui32(v279) < base.Ui32(v273) {
			v281 = v273
		} else {
			v281 = v279
		}
		v286 = F_zmalloc_usable(m, v281+int32(16), v13+int32(12))
		mBase = m.M
		v287 = m.ExcPending
		if v287 != 0 {
			return
		} else {
			v288 = *(*int32)(unsafe.Add(mBase, uint32(v13)+12))
			*(*int64)(unsafe.Add(mBase, uint32(v286)+4)) = int64(0)
			v292 = v288 + int32(-16)
			*(*int32)(unsafe.Add(mBase, uint32(v286))) = v292
			v294 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v286)+12)))
			v297 = v294&int32(254) | v135
			*(*uint8)(unsafe.Add(mBase, uint32(v286)+12)) = uint8(v297)
			v299 = int32(0)
			if v135 == v299 {
				v405 = v299
			} else {
				v305 = v286 + int32(4)
				v307 = v286 + int32(8)
				v308 = *(*int32)(unsafe.Add(mBase, uint32(l0)+292))
				v310 = *(*int64)(unsafe.Add(mBase, _consts[480]))
				v312 = base.B2i32(v310 != int64(-1))
				v319 = int32(1)
				if l4 == v319 {
					v322 = v266
				} else {
					v322 = v319
				}
				if base.Ui32(v292) < base.Ui32(v322) {
				} else {
					v325 = F_clusterSlotStatsEnabled(m, v308)
					mBase = m.M
					if v325 != 0 {
						v326 = v308
					} else {
						v326 = int32(-1)
					}
					if base.Ui32(v292) < base.Ui32(v266) {
						v328 = v292
					} else {
						v328 = v266
					}
					v329 = *(*int32)(unsafe.Add(mBase, uint32(v307)))
					if v329 == int32(0) {
						if base.Ui32(v292) < base.Ui32(v322+int32(12)) {
						} else {
							v352 = *(*int32)(unsafe.Add(mBase, uint32(v305)))
							v353 = v286 + int32(13) + v352
							*(*int32)(unsafe.Add(mBase, uint32(v307))) = v353
							v355 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v353)+10)))
							v360 = v355&int32(254) | l4&int32(1)
							*(*uint8)(unsafe.Add(mBase, uint32(v353)+10)) = uint8(v360)
							v362 = *(*int32)(unsafe.Add(mBase, uint32(v307)))
							*(*uint16)(unsafe.Add(mBase, uint32(v362)+8)) = uint16(v326)
							v365 = v288 + int32(-28)
							if base.Ui32(v365) < base.Ui32(v266) {
								v367 = v365
							} else {
								v367 = v328
							}
							*(*int32)(unsafe.Add(mBase, uint32(v362))) = v367
							v369 = int32(0)
							*(*int32)(unsafe.Add(mBase, uint32(v362)+4)) = v369
							v371 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v362)+10)))
							v376 = v371&int32(253) | v312<<(uint(int32(1))%32)
							*(*uint8)(unsafe.Add(mBase, uint32(v362)+10)) = uint8(v376)
							v378 = *(*int32)(unsafe.Add(mBase, uint32(v307)))
							*(*uint8)(unsafe.Add(mBase, uint32(v378)+11)) = uint8(v369)
							v381 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v378)+10)))
							v383 = v381 & int32(3)
							*(*uint8)(unsafe.Add(mBase, uint32(v378)+10)) = uint8(v383)
							v385 = *(*int32)(unsafe.Add(mBase, uint32(v305)))
							*(*int32)(unsafe.Add(mBase, uint32(v305))) = v385 + int32(12)
						}
					} else {
						v332 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v329)+10)))
						if v332&int32(1) != l4 {
							if base.Ui32(v292) < base.Ui32(v322+int32(12)) {
							} else {
								v352 = *(*int32)(unsafe.Add(mBase, uint32(v305)))
								v353 = v286 + int32(13) + v352
								*(*int32)(unsafe.Add(mBase, uint32(v307))) = v353
								v355 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v353)+10)))
								v360 = v355&int32(254) | l4&int32(1)
								*(*uint8)(unsafe.Add(mBase, uint32(v353)+10)) = uint8(v360)
								v362 = *(*int32)(unsafe.Add(mBase, uint32(v307)))
								*(*uint16)(unsafe.Add(mBase, uint32(v362)+8)) = uint16(v326)
								v365 = v288 + int32(-28)
								if base.Ui32(v365) < base.Ui32(v266) {
									v367 = v365
								} else {
									v367 = v328
								}
								*(*int32)(unsafe.Add(mBase, uint32(v362))) = v367
								v369 = int32(0)
								*(*int32)(unsafe.Add(mBase, uint32(v362)+4)) = v369
								v371 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v362)+10)))
								v376 = v371&int32(253) | v312<<(uint(int32(1))%32)
								*(*uint8)(unsafe.Add(mBase, uint32(v362)+10)) = uint8(v376)
								v378 = *(*int32)(unsafe.Add(mBase, uint32(v307)))
								*(*uint8)(unsafe.Add(mBase, uint32(v378)+11)) = uint8(v369)
								v381 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v378)+10)))
								v383 = v381 & int32(3)
								*(*uint8)(unsafe.Add(mBase, uint32(v378)+10)) = uint8(v383)
								v385 = *(*int32)(unsafe.Add(mBase, uint32(v305)))
								*(*int32)(unsafe.Add(mBase, uint32(v305))) = v385 + int32(12)
							}
						} else {
							v336 = int32(*(*int16)(unsafe.Add(mBase, uint32(v329)+8)))
							if v326 != v336 {
								if base.Ui32(v292) < base.Ui32(v322+int32(12)) {
								} else {
									v352 = *(*int32)(unsafe.Add(mBase, uint32(v305)))
									v353 = v286 + int32(13) + v352
									*(*int32)(unsafe.Add(mBase, uint32(v307))) = v353
									v355 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v353)+10)))
									v360 = v355&int32(254) | l4&int32(1)
									*(*uint8)(unsafe.Add(mBase, uint32(v353)+10)) = uint8(v360)
									v362 = *(*int32)(unsafe.Add(mBase, uint32(v307)))
									*(*uint16)(unsafe.Add(mBase, uint32(v362)+8)) = uint16(v326)
									v365 = v288 + int32(-28)
									if base.Ui32(v365) < base.Ui32(v266) {
										v367 = v365
									} else {
										v367 = v328
									}
									*(*int32)(unsafe.Add(mBase, uint32(v362))) = v367
									v369 = int32(0)
									*(*int32)(unsafe.Add(mBase, uint32(v362)+4)) = v369
									v371 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v362)+10)))
									v376 = v371&int32(253) | v312<<(uint(int32(1))%32)
									*(*uint8)(unsafe.Add(mBase, uint32(v362)+10)) = uint8(v376)
									v378 = *(*int32)(unsafe.Add(mBase, uint32(v307)))
									*(*uint8)(unsafe.Add(mBase, uint32(v378)+11)) = uint8(v369)
									v381 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v378)+10)))
									v383 = v381 & int32(3)
									*(*uint8)(unsafe.Add(mBase, uint32(v378)+10)) = uint8(v383)
									v385 = *(*int32)(unsafe.Add(mBase, uint32(v305)))
									*(*int32)(unsafe.Add(mBase, uint32(v305))) = v385 + int32(12)
								}
							} else {
								v338 = int32(1)
								if v312 != int32(base.Ui32(v332)>>(uint(v338)%32))&v338 {
									if base.Ui32(v292) < base.Ui32(v322+int32(12)) {
									} else {
										v352 = *(*int32)(unsafe.Add(mBase, uint32(v305)))
										v353 = v286 + int32(13) + v352
										*(*int32)(unsafe.Add(mBase, uint32(v307))) = v353
										v355 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v353)+10)))
										v360 = v355&int32(254) | l4&int32(1)
										*(*uint8)(unsafe.Add(mBase, uint32(v353)+10)) = uint8(v360)
										v362 = *(*int32)(unsafe.Add(mBase, uint32(v307)))
										*(*uint16)(unsafe.Add(mBase, uint32(v362)+8)) = uint16(v326)
										v365 = v288 + int32(-28)
										if base.Ui32(v365) < base.Ui32(v266) {
											v367 = v365
										} else {
											v367 = v328
										}
										*(*int32)(unsafe.Add(mBase, uint32(v362))) = v367
										v369 = int32(0)
										*(*int32)(unsafe.Add(mBase, uint32(v362)+4)) = v369
										v371 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v362)+10)))
										v376 = v371&int32(253) | v312<<(uint(int32(1))%32)
										*(*uint8)(unsafe.Add(mBase, uint32(v362)+10)) = uint8(v376)
										v378 = *(*int32)(unsafe.Add(mBase, uint32(v307)))
										*(*uint8)(unsafe.Add(mBase, uint32(v378)+11)) = uint8(v369)
										v381 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v378)+10)))
										v383 = v381 & int32(3)
										*(*uint8)(unsafe.Add(mBase, uint32(v378)+10)) = uint8(v383)
										v385 = *(*int32)(unsafe.Add(mBase, uint32(v305)))
										*(*int32)(unsafe.Add(mBase, uint32(v305))) = v385 + int32(12)
									}
								} else {
									v343 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v329)+11)))
									if v343 != 0 {
										if base.Ui32(v292) < base.Ui32(v322+int32(12)) {
										} else {
											v352 = *(*int32)(unsafe.Add(mBase, uint32(v305)))
											v353 = v286 + int32(13) + v352
											*(*int32)(unsafe.Add(mBase, uint32(v307))) = v353
											v355 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v353)+10)))
											v360 = v355&int32(254) | l4&int32(1)
											*(*uint8)(unsafe.Add(mBase, uint32(v353)+10)) = uint8(v360)
											v362 = *(*int32)(unsafe.Add(mBase, uint32(v307)))
											*(*uint16)(unsafe.Add(mBase, uint32(v362)+8)) = uint16(v326)
											v365 = v288 + int32(-28)
											if base.Ui32(v365) < base.Ui32(v266) {
												v367 = v365
											} else {
												v367 = v328
											}
											*(*int32)(unsafe.Add(mBase, uint32(v362))) = v367
											v369 = int32(0)
											*(*int32)(unsafe.Add(mBase, uint32(v362)+4)) = v369
											v371 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v362)+10)))
											v376 = v371&int32(253) | v312<<(uint(int32(1))%32)
											*(*uint8)(unsafe.Add(mBase, uint32(v362)+10)) = uint8(v376)
											v378 = *(*int32)(unsafe.Add(mBase, uint32(v307)))
											*(*uint8)(unsafe.Add(mBase, uint32(v378)+11)) = uint8(v369)
											v381 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v378)+10)))
											v383 = v381 & int32(3)
											*(*uint8)(unsafe.Add(mBase, uint32(v378)+10)) = uint8(v383)
											v385 = *(*int32)(unsafe.Add(mBase, uint32(v305)))
											*(*int32)(unsafe.Add(mBase, uint32(v305))) = v385 + int32(12)
										}
									} else {
										v344 = *(*int32)(unsafe.Add(mBase, uint32(v307)))
										v345 = *(*int32)(unsafe.Add(mBase, uint32(v344)))
										*(*int32)(unsafe.Add(mBase, uint32(v344))) = v345 + v328
									}
								}
							}
						}
					}
				}
				v404 = *(*int32)(unsafe.Add(mBase, uint32(v286)+4))
				v405 = v404
			}
			if v266 == int32(0) {
			} else {
				v411 = F__emscripten_memcpy_bulkmem(m, v286+v405+int32(13), v265, v266)
				mBase = m.M
			}
			*(*int32)(unsafe.Add(mBase, uint32(v286)+4)) = v405 + v266
			v415 = F_listAddNodeTail(m, l1, v286)
			mBase = m.M
			v416 = m.ExcPending
			if v416 != 0 {
				return
			} else {
				v419 = *(*int64)(unsafe.Add(mBase, uint32(l0)+384))
				if v419 == int64(-1) {
					v422 = int32(160)
				} else {
					v422 = int32(384)
				}
				v423 = l0 + v422
				v424 = *(*int64)(unsafe.Add(mBase, uint32(v423)))
				v425 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v286))))
				*(*int64)(unsafe.Add(mBase, uint32(v423))) = v424 + v425
				v429 = F_closeClientOnOutputBufferLimitReached(m, l0, int32(1))
				mBase = m.M
				v430 = m.ExcPending
				if v430 != 0 {
					return
				} else {
					m.G0 = v13 + int32(16)
					return
				}
			}
		}
	}
}
func F__addReplyToBufferOrList(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v25 int32
	_ = v25
	var v32 int64
	_ = v32
	var v36 int32
	_ = v36
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v81 int32
	_ = v81
	var v83 int64
	_ = v83
	var v88 int32
	_ = v88
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
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
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v126 int32
	_ = v126
	var v128 int64
	_ = v128
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v142 int32
	_ = v142
	var v146 int32
	_ = v146
	var v149 int32
	_ = v149
	var v151 int32
	_ = v151
	var v156 int32
	_ = v156
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+200))
	if v10&int32(64) != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	m.G0 = v8 + int32(16)
	return
L2:
	;
	if v10&int32(1) != 0 {
		goto L5
	} else {
		goto L6
	}
L3:
	;
	v83 = *(*int64)(unsafe.Add(mBase, uint32(l0)+384))
	if v83 == int64(-1) {
		goto L35
	} else {
		goto L36
	}
L4:
	;
	v72 = *(*int32)(unsafe.Add(mBase, uint32(l0)+72))
	if v72 != 0 {
		goto L29
	} else {
		goto L30
	}
L5:
	;
	v32 = *(*int64)(unsafe.Add(mBase, uint32(l0)+264))
	*(*int64)(unsafe.Add(mBase, uint32(l0)+264)) = v32 + base.I64_extend_i32_u(l2)
	goto L13
L6:
	;
	if v10&int32(2) == int32(0) {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	if v10&int32(262144) != 0 {
		goto L5
	} else {
		goto L10
	}
L8:
	;
	if v10&int32(4) == int32(0) {
		goto L4
	} else {
		goto L9
	}
L9:
	;
	goto L7
L10:
	;
	v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+216))
	if v25 == int32(0) {
		goto L5
	} else {
		goto L11
	}
L11:
	;
	goto L12
L12:
	;
	goto L5
L13:
	;
	v36 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+206)))
	if v36&int32(2) == int32(0) {
		goto L3
	} else {
		goto L14
	}
L14:
	;
	v42 = *(*int32)(unsafe.Add(mBase, _consts[14]))
	if l0 != v42 {
		goto L3
	} else {
		goto L15
	}
L15:
	;
	v45 = *(*int32)(unsafe.Add(mBase, _consts[300]))
	if v45 == int32(0) {
		goto L3
	} else {
		goto L16
	}
L16:
	;
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v45)+68))
	if v48 == int32(0) {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	if l2 == int32(0) {
		goto L1
	} else {
		goto L25
	}
L18:
	;
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v48)+48))
	if v51 == int32(288) {
		goto L3
	} else {
		goto L19
	}
L19:
	;
	if v51 == int32(286) {
		goto L3
	} else {
		goto L20
	}
L20:
	;
	if v51 == int32(284) {
		goto L3
	} else {
		goto L21
	}
L21:
	;
	if v51 == int32(282) {
		goto L3
	} else {
		goto L22
	}
L22:
	;
	if v51 == int32(287) {
		goto L3
	} else {
		goto L23
	}
L23:
	;
	if v51 == int32(289) {
		goto L3
	} else {
		goto L24
	}
L24:
	;
	goto L17
L25:
	;
	v68 = *(*int32)(unsafe.Add(mBase, _consts[485]))
	F__addReplyPayloadToList(m, l0, v68, l1, l2, int32(0))
	mBase = m.M
	v71 = m.ExcPending
	if v71 != 0 {
		goto L26
	} else {
		goto L27
	}
L26:
	;
	return
L27:
	;
	goto L1
L28:
	;
	if v75 != 0 {
		goto L31
	} else {
		goto L32
	}
L29:
	;
	v74 = *(*int32)(unsafe.Add(mBase, uint32(v72)+140))
	v75 = v74
	goto L28
L30:
	;
	v75 = int32(0)
	goto L28
L31:
	;
	v77 = v75
	goto L33
L32:
	;
	v77 = int32(_a277)
	goto L33
L33:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8))) = v77
	F_logInvalidUseAndFreeClientAsync(m, l0, int32(_a1642), v8)
	mBase = m.M
	v81 = m.ExcPending
	if v81 != 0 {
		goto L26
	} else {
		goto L34
	}
L34:
	;
	goto L1
L35:
	;
	v92 = int32(0)
	if l2 == v92 {
		v146 = v92
		goto L40
	} else {
		goto L41
	}
L36:
	;
	if l2 == int32(0) {
		goto L1
	} else {
		goto L37
	}
L37:
	;
	v88 = *(*int32)(unsafe.Add(mBase, uint32(l0)+376))
	F__addReplyPayloadToList(m, l0, v88, l1, l2, int32(0))
	mBase = m.M
	v91 = m.ExcPending
	if v91 != 0 {
		goto L26
	} else {
		goto L38
	}
L38:
	;
	goto L1
L39:
	;
	if base.Ui32(l2) <= base.Ui32(v146) {
		goto L1
	} else {
		goto L55
	}
L40:
	;
	goto L39
L41:
	;
	v97 = *(*int32)(unsafe.Add(mBase, uint32(l0)+180))
	if v97 != 0 {
		goto L42
	} else {
		goto L43
	}
L42:
	;
	v109 = *(*int32)(unsafe.Add(mBase, _consts[484]))
	if v109 != 0 {
		v146 = v92
		goto L40
	} else {
		goto L44
	}
L43:
	;
	v99 = F_isCopyAvoidPreferred(m, l0, int32(0))
	mBase = m.M
	v100 = *(*int32)(unsafe.Add(mBase, uint32(l0)+200))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+200)) = v100&int32(-16777217) | v99<<(uint(int32(24))%32)
	goto L42
L44:
	;
	v110 = *(*int32)(unsafe.Add(mBase, uint32(l0)+132))
	v111 = *(*int32)(unsafe.Add(mBase, uint32(v110)+20))
	if v111 != 0 {
		v146 = v92
		goto L40
	} else {
		goto L45
	}
L45:
	;
	v112 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
	v113 = *(*int32)(unsafe.Add(mBase, uint32(l0)+180))
	v114 = v112 - v113
	v115 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+203)))
	if v115&int32(1) != 0 {
		goto L47
	} else {
		goto L48
	}
L46:
	;
	if v132 == int32(0) {
		v146 = v92
		goto L40
	} else {
		goto L52
	}
L47:
	;
	v120 = *(*int32)(unsafe.Add(mBase, uint32(l0)+124))
	v126 = *(*int32)(unsafe.Add(mBase, uint32(l0)+292))
	v128 = *(*int64)(unsafe.Add(mBase, _consts[480]))
	v131 = F_upsertPayloadHeader(m, v120, l0+int32(180), l0+int32(184), int32(0), l2, v126, base.B2i32(v128 != int64(-1)), v114)
	mBase = m.M
	v132 = v131
	goto L46
L48:
	;
	if base.Ui32(v114) < base.Ui32(l2) {
		goto L49
	} else {
		goto L50
	}
L49:
	;
	v119 = v114
	goto L51
L50:
	;
	v119 = l2
	goto L51
L51:
	;
	v132 = v119
	goto L46
L52:
	;
	v135 = *(*int32)(unsafe.Add(mBase, uint32(l0)+124))
	v136 = *(*int32)(unsafe.Add(mBase, uint32(l0)+180))
	v138 = F___memcpy(m, v135+v136, l1, v132)
	mBase = m.M
	v139 = *(*int32)(unsafe.Add(mBase, uint32(l0)+180))
	v140 = v139 + v132
	*(*int32)(unsafe.Add(mBase, uint32(l0)+180)) = v140
	v142 = *(*int32)(unsafe.Add(mBase, uint32(l0)+276))
	if base.Ui32(v140) <= base.Ui32(v142) {
		goto L53
	} else {
		goto L54
	}
L53:
	;
	v146 = v132
	goto L40
L54:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+276)) = v140
	goto L53
L55:
	;
	v149 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+184)) = v149
	v151 = *(*int32)(unsafe.Add(mBase, uint32(l0)+132))
	F__addReplyPayloadToList(m, l0, v151, l1+v146, l2-v146, v149)
	mBase = m.M
	v156 = m.ExcPending
	if v156 != 0 {
		goto L26
	} else {
		goto L56
	}
L56:
	;
	goto L1
}
func F_addReply(m *base.Module, l0 int32, l1 int32) {
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
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v60 int64
	_ = v60
	var v69 int32
	_ = v69
	var v73 int32
	_ = v73
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v80 int64
	_ = v80
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	v7 = m.G0
	v9 = v7 - int32(32)
	m.G0 = v9
	v11 = F_prepareClientToWrite(m, l0)
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		return
	} else {
		if v11 != 0 {
			m.G0 = v9 + int32(32)
			return
		} else {
			v13 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
			switch int32(base.Ui32(v13)>>(uint(int32(4))%32)) & int32(15) {
			case 0, 8:
				v18 = F_objectGetVal(m, l1)
				mBase = m.M
				v20 = F_objectGetVal(m, l1)
				mBase = m.M
				v23 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20+int32(-1)))))
				switch v23 & int32(7) {
				case 0:
					F__addReplyToBufferOrList(m, l0, v18, int32(base.Ui32(v23)>>(uint(int32(3))%32)))
					mBase = m.M
					v29 = m.ExcPending
					if v29 != 0 {
						return
					} else {
						m.G0 = v9 + int32(32)
						return
					}
				case 1:
					v32 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20+int32(-3)))))
					F__addReplyToBufferOrList(m, l0, v18, v32)
					mBase = m.M
					v34 = m.ExcPending
					if v34 != 0 {
						return
					} else {
						m.G0 = v9 + int32(32)
						return
					}
				case 2:
					v37 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v20+int32(-5)))))
					F__addReplyToBufferOrList(m, l0, v18, v37)
					mBase = m.M
					v39 = m.ExcPending
					if v39 != 0 {
						return
					} else {
						m.G0 = v9 + int32(32)
						return
					}
				case 3:
					v42 = *(*int32)(unsafe.Add(mBase, uint32(v20+int32(-9))))
					F__addReplyToBufferOrList(m, l0, v18, v42)
					mBase = m.M
					v44 = m.ExcPending
					if v44 != 0 {
						return
					} else {
						m.G0 = v9 + int32(32)
						return
					}
				case 4:
					v47 = *(*int32)(unsafe.Add(mBase, uint32(v20+int32(-17))))
					v48 = v47
					F__addReplyToBufferOrList(m, l0, v18, v48)
					mBase = m.M
					v50 = m.ExcPending
					if v50 != 0 {
						return
					} else {
						m.G0 = v9 + int32(32)
						return
					}
				default:
					v48 = int32(0)
					F__addReplyToBufferOrList(m, l0, v18, v48)
					mBase = m.M
					v50 = m.ExcPending
					if v50 != 0 {
						return
					} else {
						m.G0 = v9 + int32(32)
						return
					}
				}
			case 1:
				v59 = F_objectGetVal(m, l1)
				mBase = m.M
				v60 = base.I64_extend_i32_s(v59)
				if v60 <= int64(-1) {
					v69 = int32(45)
					*(*uint8)(unsafe.Add(mBase, uint32(v9))) = uint8(v69)
					v73 = int32(1)
					v78 = v9 + v73
					v79 = int32(31)
					v80 = int64(0) - v60
					v81 = v73
				} else {
					v78 = v9
					v79 = int32(32)
					v80 = v60
					v81 = int32(0)
				}
				v82 = F_ull2string(m, v78, v79, v80)
				mBase = m.M
				if v82 == int32(0) {
					v101 = int32(0)
				} else {
					v101 = v82 + v81
				}
				F__addReplyToBufferOrList(m, l0, v9, v101)
				mBase = m.M
				v103 = m.ExcPending
				if v103 != 0 {
					return
				} else {
					m.G0 = v9 + int32(32)
					return
				}
			default:
				F__serverPanic_1(m, int32(_a1630), int32(817), int32(_a1649), int32(0))
				mBase = m.M
				v56 = m.ExcPending
				if v56 != 0 {
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
func F_addReplyBulkCBuffer(m *base.Module, l0 int32, l1 int32, l2 int32) {
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
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v32 int64
	_ = v32
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int64
	_ = v52
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v73 int32
	_ = v73
	var v77 int32
	_ = v77
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
	var v89 int32
	_ = v89
	v6 = m.G0
	v8 = v6 - int32(128)
	m.G0 = v8
	v10 = F_prepareClientToWrite(m, l0)
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return
	} else {
		if v10 != 0 {
			m.G0 = v8 + int32(128)
			return
		} else {
			if base.Ui32(int32(31)) < base.Ui32(l2) {
				v27 = int32(36)
				*(*uint8)(unsafe.Add(mBase, uint32(v8))) = uint8(v27)
				v30 = v8 | int32(1)
				v32 = base.I64_extend_i32_u(l2)
				if v32 <= int64(-1) {
					v41 = int32(45)
					*(*uint8)(unsafe.Add(mBase, uint32(v30))) = uint8(v41)
					v45 = int32(1)
					v50 = v30 + v45
					v51 = int32(126)
					v52 = int64(0) - v32
					v53 = v45
				} else {
					v50 = v30
					v51 = int32(127)
					v52 = v32
					v53 = int32(0)
				}
				v54 = F_ull2string(m, v50, v51, v52)
				mBase = m.M
				if v54 == int32(0) {
					v73 = int32(0)
				} else {
					v73 = v54 + v53
				}
				v77 = int32(2573)
				*(*uint16)(unsafe.Add(mBase, uint32(v73+v8+int32(1)))) = uint16(v77)
				F__addReplyToBufferOrList(m, l0, v8, v73+int32(3))
				mBase = m.M
				v82 = m.ExcPending
				if v82 != 0 {
					return
				} else {
					F__addReplyToBufferOrList(m, l0, l1, l2)
					mBase = m.M
					v85 = m.ExcPending
					if v85 != 0 {
						return
					} else {
						F__addReplyToBufferOrList(m, l0, int32(_a823), int32(2))
						mBase = m.M
						v89 = m.ExcPending
						if v89 != 0 {
							return
						} else {
							m.G0 = v8 + int32(128)
							return
						}
					}
				}
			} else {
				v18 = *(*int32)(unsafe.Add(mBase, uint32(l2<<(uint(int32(2))%32))+uint32(_consts[488])))
				v19 = F_objectGetVal(m, v18)
				mBase = m.M
				if base.Ui32(l2) < base.Ui32(int32(10)) {
					v24 = int32(4)
				} else {
					v24 = int32(5)
				}
				F__addReplyToBufferOrList(m, l0, v19, v24)
				mBase = m.M
				v26 = m.ExcPending
				if v26 != 0 {
					return
				} else {
					F__addReplyToBufferOrList(m, l0, l1, l2)
					mBase = m.M
					v85 = m.ExcPending
					if v85 != 0 {
						return
					} else {
						F__addReplyToBufferOrList(m, l0, int32(_a823), int32(2))
						mBase = m.M
						v89 = m.ExcPending
						if v89 != 0 {
							return
						} else {
							m.G0 = v8 + int32(128)
							return
						}
					}
				}
			}
		}
	}
}
func F_addReplyBulkLen(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v32 int64
	_ = v32
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int64
	_ = v52
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v73 int32
	_ = v73
	var v77 int32
	_ = v77
	var v82 int32
	_ = v82
	v4 = m.G0
	v6 = v4 - int32(128)
	m.G0 = v6
	v8 = F_stringObjectLen(m, l1)
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return
	} else {
		v10 = F_prepareClientToWrite(m, l0)
		mBase = m.M
		v11 = m.ExcPending
		if v11 != 0 {
			return
		} else {
			if v10 != 0 {
				m.G0 = v6 + int32(128)
				return
			} else {
				if base.Ui32(int32(31)) < base.Ui32(v8) {
					v27 = int32(36)
					*(*uint8)(unsafe.Add(mBase, uint32(v6))) = uint8(v27)
					v30 = v6 | int32(1)
					v32 = base.I64_extend_i32_u(v8)
					if v32 <= int64(-1) {
						v41 = int32(45)
						*(*uint8)(unsafe.Add(mBase, uint32(v30))) = uint8(v41)
						v45 = int32(1)
						v50 = v30 + v45
						v51 = int32(126)
						v52 = int64(0) - v32
						v53 = v45
					} else {
						v50 = v30
						v51 = int32(127)
						v52 = v32
						v53 = int32(0)
					}
					v54 = F_ull2string(m, v50, v51, v52)
					mBase = m.M
					if v54 == int32(0) {
						v73 = int32(0)
					} else {
						v73 = v54 + v53
					}
					v77 = int32(2573)
					*(*uint16)(unsafe.Add(mBase, uint32(v73+v6+int32(1)))) = uint16(v77)
					F__addReplyToBufferOrList(m, l0, v6, v73+int32(3))
					mBase = m.M
					v82 = m.ExcPending
					if v82 != 0 {
						return
					} else {
						m.G0 = v6 + int32(128)
						return
					}
				} else {
					v18 = *(*int32)(unsafe.Add(mBase, uint32(v8<<(uint(int32(2))%32))+uint32(_consts[488])))
					v19 = F_objectGetVal(m, v18)
					mBase = m.M
					if base.Ui32(v8) < base.Ui32(int32(10)) {
						v24 = int32(4)
					} else {
						v24 = int32(5)
					}
					F__addReplyToBufferOrList(m, l0, v19, v24)
					mBase = m.M
					v26 = m.ExcPending
					if v26 != 0 {
						return
					} else {
						m.G0 = v6 + int32(128)
						return
					}
				}
			}
		}
	}
}
func F_addReplyBulkLongLong(m *base.Module, l0 int32, l1 int64) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v17 int32
	_ = v17
	var v21 int32
	_ = v21
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int64
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	v4 = m.G0
	v5 = int32(64)
	v6 = v4 - v5
	m.G0 = v6
	if l1 <= int64(-1) {
		v17 = int32(45)
		*(*uint8)(unsafe.Add(mBase, uint32(v6))) = uint8(v17)
		v21 = int32(1)
		v26 = v6 + v21
		v27 = int32(63)
		v28 = int64(0) - l1
		v29 = v21
	} else {
		v26 = v6
		v27 = v5
		v28 = l1
		v29 = int32(0)
	}
	v30 = F_ull2string(m, v26, v27, v28)
	mBase = m.M
	if v30 == int32(0) {
		v49 = int32(0)
	} else {
		v49 = v30 + v29
	}
	F_addReplyBulkCBuffer(m, l0, v6, v49)
	mBase = m.M
	v51 = m.ExcPending
	if v51 != 0 {
		return
	} else {
		m.G0 = v6 + int32(64)
		return
	}
}
func F_addReplyClusterLinksDescription(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
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
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v34 int32
	_ = v34
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
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
	var v52 int64
	_ = v52
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
	var v63 int64
	_ = v63
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v85 int32
	_ = v85
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v105 int32
	_ = v105
	var v110 int32
	_ = v110
	var v114 int32
	_ = v114
	var v118 int32
	_ = v118
	var v126 int32
	_ = v126
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
	var v139 int32
	_ = v139
	var v143 int32
	_ = v143
	var v146 int32
	_ = v146
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v158 int32
	_ = v158
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v166 int32
	_ = v166
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v175 int32
	_ = v175
	var v177 int32
	_ = v177
	var v179 int64
	_ = v179
	var v180 int64
	_ = v180
	var v181 int64
	_ = v181
	var v182 int64
	_ = v182
	var v183 int64
	_ = v183
	var v184 int64
	_ = v184
	var v185 int64
	_ = v185
	var v187 int64
	_ = v187
	var v189 int64
	_ = v189
	var v191 int64
	_ = v191
	var v193 int64
	_ = v193
	var v195 int64
	_ = v195
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v202 int32
	_ = v202
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v209 int32
	_ = v209
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v217 int32
	_ = v217
	var v221 int32
	_ = v221
	var v223 int32
	_ = v223
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v237 int32
	_ = v237
	var v242 int32
	_ = v242
	var v246 int32
	_ = v246
	var v250 int32
	_ = v250
	var v256 int32
	_ = v256
	var v261 int32
	_ = v261
	var v263 int32
	_ = v263
	v7 = F_addReplyDeferredLen(m, l0)
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v9 = int32(0)
	v11 = *(*int32)(unsafe.Add(mBase, _consts[111]))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(v11)+32))
	v13 = F_dictGetSafeIterator(m, v12)
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		goto L1
	} else {
		goto L4
	}
L3:
	;
	F_dictReleaseIterator(m, v13)
	mBase = m.M
	v261 = m.ExcPending
	if v261 != 0 {
		goto L1
	} else {
		goto L68
	}
L4:
	;
	v22 = v13 + int32(20)
	v23 = *(*int32)(unsafe.Add(mBase, uint32(v13)+16))
	if v23 != 0 {
		goto L7
	} else {
		goto L8
	}
L5:
	;
	if v118 == int32(0) {
		v256 = v9
		goto L3
	} else {
		goto L31
	}
L6:
	;
	v29 = v22
	v30 = v26
	goto L9
L7:
	;
	v26 = int32(1)
	goto L6
L8:
	;
	v26 = int32(0)
	goto L6
L9:
	;
	switch v30 {
	case 0:
		goto L14
	default:
		goto L13
	}
L11:
	;
	v30 = int32(0)
	goto L9
L12:
	;
	goto L5
L13:
	;
	v110 = *(*int32)(unsafe.Add(mBase, uint32(v29)))
	*(*int32)(unsafe.Add(mBase, uint32(v13)+16)) = v110
	if v110 == int32(0) {
		goto L11
	} else {
		goto L30
	}
L14:
	;
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
	if v34 != int32(-1) {
		v73 = v34
		goto L15
	} else {
		goto L16
	}
L15:
	;
	v74 = int32(1)
	v75 = v73 + v74
	*(*int32)(unsafe.Add(mBase, uint32(v13)+4)) = v75
	v77 = int32(0)
	v80 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
	v81 = *(*int32)(unsafe.Add(mBase, uint32(v13)+8))
	v85 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v80+v81+int32(26)))))
	if v85 == int32(255) {
		goto L24
	} else {
		goto L25
	}
L16:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v13)+8))
	if v38 != 0 {
		v73 = int32(-1)
		goto L15
	} else {
		goto L17
	}
L17:
	;
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v13)+12))
	if v40 == int32(0) {
		goto L19
	} else {
		goto L20
	}
L18:
	;
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v66)+20))
	if v67 != int32(-1) {
		goto L21
	} else {
		goto L22
	}
L19:
	;
	v47 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v39)+16)))
	v48 = int64(*(*int8)(unsafe.Add(mBase, uint32(v39)+27)))
	v49 = int64(*(*int32)(unsafe.Add(mBase, uint32(v39)+8)))
	v50 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v39)+12)))
	v51 = int64(*(*int8)(unsafe.Add(mBase, uint32(v39)+26)))
	v52 = int64(*(*int32)(unsafe.Add(mBase, uint32(v39)+4)))
	v53 = F_wangHash64(m, v52)
	mBase = m.M
	v55 = F_wangHash64(m, v51+v53)
	mBase = m.M
	v57 = F_wangHash64(m, v50+v55)
	mBase = m.M
	v59 = F_wangHash64(m, v49+v57)
	mBase = m.M
	v61 = F_wangHash64(m, v48+v59)
	mBase = m.M
	v63 = F_wangHash64(m, v47+v61)
	mBase = m.M
	*(*int64)(unsafe.Add(mBase, uint32(v13)+24)) = v63
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
	v66 = v65
	goto L18
L20:
	;
	v43 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v39)+24)))
	v45 = v43 + int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v39)+24)) = uint16(v45)
	v66 = v39
	goto L18
L21:
	;
	v73 = v67 + int32(-1)
	goto L15
L22:
	;
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
	v73 = v70
	goto L15
L23:
	;
	v100 = int32(2)
	v105 = *(*int32)(unsafe.Add(mBase, uint32(v80+v98<<(uint(v100)%32)+int32(4))))
	v29 = v105 + v99<<(uint(v100)%32)
	v30 = int32(1)
	goto L9
L24:
	;
	v89 = v77
	goto L26
L25:
	;
	v89 = v74 << (uint(v85) % 32)
	goto L26
L26:
	;
	if v75 < v89 {
		v98 = v81
		v99 = v75
		goto L23
	} else {
		goto L27
	}
L27:
	;
	if v81 != 0 {
		v118 = v77
		goto L12
	} else {
		goto L28
	}
L28:
	;
	v91 = *(*int32)(unsafe.Add(mBase, uint32(v80)+20))
	if v91 == int32(-1) {
		v118 = v77
		goto L12
	} else {
		goto L29
	}
L29:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v13)+4)) = int64(4294967296)
	v98 = int32(1)
	v99 = int32(0)
	goto L23
L30:
	;
	v114 = *(*int32)(unsafe.Add(mBase, uint32(v110)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v22))) = v114
	v118 = v110
	goto L12
L31:
	;
	v126 = v9
	v128 = v118
	goto L32
L32:
	;
	v130 = *(*int32)(unsafe.Add(mBase, uint32(v128)+8))
	goto L35
L33:
	;
	v256 = v146
	goto L3
L34:
	;
	v139 = *(*int32)(unsafe.Add(mBase, uint32(v130)+2348))
	if v139 == int32(0) {
		v146 = v138
		goto L38
	} else {
		goto L39
	}
L35:
	;
	v131 = *(*int32)(unsafe.Add(mBase, uint32(v130)+2344))
	if v131 == int32(0) {
		v138 = v126
		goto L34
	} else {
		goto L36
	}
L36:
	;
	F_addReplyClusterLinkDescription(m, l0, v131)
	mBase = m.M
	v135 = m.ExcPending
	if v135 != 0 {
		goto L1
	} else {
		goto L37
	}
L37:
	;
	v138 = v126 + int32(1)
	goto L34
L38:
	;
	v154 = v13 + int32(20)
	v155 = *(*int32)(unsafe.Add(mBase, uint32(v13)+16))
	if v155 != 0 {
		goto L43
	} else {
		goto L44
	}
L39:
	;
	F_addReplyClusterLinkDescription(m, l0, v139)
	mBase = m.M
	v143 = m.ExcPending
	if v143 != 0 {
		goto L1
	} else {
		goto L40
	}
L40:
	;
	v146 = v138 + int32(1)
	goto L38
L41:
	;
	if v250 != 0 {
		v126 = v146
		v128 = v250
		goto L32
	} else {
		goto L67
	}
L42:
	;
	v161 = v154
	v162 = v158
	goto L45
L43:
	;
	v158 = int32(1)
	goto L42
L44:
	;
	v158 = int32(0)
	goto L42
L45:
	;
	switch v162 {
	case 0:
		goto L50
	default:
		goto L49
	}
L47:
	;
	v162 = int32(0)
	goto L45
L48:
	;
	goto L41
L49:
	;
	v242 = *(*int32)(unsafe.Add(mBase, uint32(v161)))
	*(*int32)(unsafe.Add(mBase, uint32(v13)+16)) = v242
	if v242 == int32(0) {
		goto L47
	} else {
		goto L66
	}
L50:
	;
	v166 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
	if v166 != int32(-1) {
		v205 = v166
		goto L51
	} else {
		goto L52
	}
L51:
	;
	v206 = int32(1)
	v207 = v205 + v206
	*(*int32)(unsafe.Add(mBase, uint32(v13)+4)) = v207
	v209 = int32(0)
	v212 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
	v213 = *(*int32)(unsafe.Add(mBase, uint32(v13)+8))
	v217 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v212+v213+int32(26)))))
	if v217 == int32(255) {
		goto L60
	} else {
		goto L61
	}
L52:
	;
	v170 = *(*int32)(unsafe.Add(mBase, uint32(v13)+8))
	if v170 != 0 {
		v205 = int32(-1)
		goto L51
	} else {
		goto L53
	}
L53:
	;
	v171 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
	v172 = *(*int32)(unsafe.Add(mBase, uint32(v13)+12))
	if v172 == int32(0) {
		goto L55
	} else {
		goto L56
	}
L54:
	;
	v199 = *(*int32)(unsafe.Add(mBase, uint32(v198)+20))
	if v199 != int32(-1) {
		goto L57
	} else {
		goto L58
	}
L55:
	;
	v179 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v171)+16)))
	v180 = int64(*(*int8)(unsafe.Add(mBase, uint32(v171)+27)))
	v181 = int64(*(*int32)(unsafe.Add(mBase, uint32(v171)+8)))
	v182 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v171)+12)))
	v183 = int64(*(*int8)(unsafe.Add(mBase, uint32(v171)+26)))
	v184 = int64(*(*int32)(unsafe.Add(mBase, uint32(v171)+4)))
	v185 = F_wangHash64(m, v184)
	mBase = m.M
	v187 = F_wangHash64(m, v183+v185)
	mBase = m.M
	v189 = F_wangHash64(m, v182+v187)
	mBase = m.M
	v191 = F_wangHash64(m, v181+v189)
	mBase = m.M
	v193 = F_wangHash64(m, v180+v191)
	mBase = m.M
	v195 = F_wangHash64(m, v179+v193)
	mBase = m.M
	*(*int64)(unsafe.Add(mBase, uint32(v13)+24)) = v195
	v197 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
	v198 = v197
	goto L54
L56:
	;
	v175 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v171)+24)))
	v177 = v175 + int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v171)+24)) = uint16(v177)
	v198 = v171
	goto L54
L57:
	;
	v205 = v199 + int32(-1)
	goto L51
L58:
	;
	v202 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
	v205 = v202
	goto L51
L59:
	;
	v232 = int32(2)
	v237 = *(*int32)(unsafe.Add(mBase, uint32(v212+v230<<(uint(v232)%32)+int32(4))))
	v161 = v237 + v231<<(uint(v232)%32)
	v162 = int32(1)
	goto L45
L60:
	;
	v221 = v209
	goto L62
L61:
	;
	v221 = v206 << (uint(v217) % 32)
	goto L62
L62:
	;
	if v207 < v221 {
		v230 = v213
		v231 = v207
		goto L59
	} else {
		goto L63
	}
L63:
	;
	if v213 != 0 {
		v250 = v209
		goto L48
	} else {
		goto L64
	}
L64:
	;
	v223 = *(*int32)(unsafe.Add(mBase, uint32(v212)+20))
	if v223 == int32(-1) {
		v250 = v209
		goto L48
	} else {
		goto L65
	}
L65:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v13)+4)) = int64(4294967296)
	v230 = int32(1)
	v231 = int32(0)
	goto L59
L66:
	;
	v246 = *(*int32)(unsafe.Add(mBase, uint32(v242)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v154))) = v246
	v250 = v242
	goto L48
L67:
	;
	goto L33
L68:
	;
	F_setDeferredArrayLen(m, l0, v7, v256)
	mBase = m.M
	v263 = m.ExcPending
	if v263 != 0 {
		goto L1
	} else {
		goto L69
	}
L69:
	;
	return
}
func F_addReplyCommandArgList(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v17 int32
	_ = v17
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v93 int32
	_ = v93
	var v94 int64
	_ = v94
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v145 int32
	_ = v145
	var v148 int32
	_ = v148
	var v153 int32
	_ = v153
	var v158 int32
	_ = v158
	var v165 int32
	_ = v165
	var v170 int32
	_ = v170
	var v174 int32
	_ = v174
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v185 int32
	_ = v185
	var v187 int32
	_ = v187
	F_addReplyArrayLen(m, l0, l2)
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v11 = int32(0)
	if l2 <= v11 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	return
L4:
	;
	v17 = v11
	goto L5
L5:
	;
	v26 = l1 + v17*int32(44)
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v26)+8))
	if v27 == int32(-1) {
		goto L7
	} else {
		goto L8
	}
L6:
	;
	goto L3
L7:
	;
	v30 = int32(2)
	goto L9
L8:
	;
	v30 = int32(3)
	goto L9
L9:
	;
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v26)+12))
	v32 = int32(0)
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v26)+16))
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v26)+20))
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v26)+28))
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v26)+24))
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v26)+4))
	v53 = v51 + int32(-7)
	v55 = base.B2i32(base.Ui32(v53) < base.Ui32(int32(2)))
	F_addReplyMapLen(m, l0, v30+base.B2i32(v31 != v32)+base.B2i32(v35 != v32)+base.B2i32(v39 != v32)+base.B2i32(v43 != v32)+base.B2i32(v47 != v32)+v55+base.B2i32(base.Ui32(int32(1)) < base.Ui32(v53)))
	mBase = m.M
	v61 = m.ExcPending
	if v61 != 0 {
		goto L1
	} else {
		goto L10
	}
L10:
	;
	F_addReplyBulkCString(m, l0, int32(_a373))
	mBase = m.M
	v64 = m.ExcPending
	if v64 != 0 {
		goto L1
	} else {
		goto L11
	}
L11:
	;
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v26)))
	F_addReplyBulkCString(m, l0, v65)
	mBase = m.M
	v67 = m.ExcPending
	if v67 != 0 {
		goto L1
	} else {
		goto L12
	}
L12:
	;
	F_addReplyBulkCString(m, l0, int32(_a1702))
	mBase = m.M
	v70 = m.ExcPending
	if v70 != 0 {
		goto L1
	} else {
		goto L13
	}
L13:
	;
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v26)+4))
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v71<<(uint(int32(2))%32))+uint32(_consts[785])))
	F_addReplyBulkCString(m, l0, v76)
	mBase = m.M
	v78 = m.ExcPending
	if v78 != 0 {
		goto L1
	} else {
		goto L14
	}
L14:
	;
	if base.Ui32(v53) < base.Ui32(int32(2)) {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	v88 = *(*int32)(unsafe.Add(mBase, uint32(v26)+8))
	if v88 == int32(-1) {
		goto L21
	} else {
		goto L22
	}
L16:
	;
	F_addReplyBulkCString(m, l0, int32(_a2254))
	mBase = m.M
	v81 = m.ExcPending
	if v81 != 0 {
		goto L1
	} else {
		goto L17
	}
L17:
	;
	v82 = *(*int32)(unsafe.Add(mBase, uint32(v26)+40))
	if v82 != 0 {
		v84 = v82
		goto L18
	} else {
		goto L19
	}
L18:
	;
	F_addReplyBulkCString(m, l0, v84)
	mBase = m.M
	v86 = m.ExcPending
	if v86 != 0 {
		goto L1
	} else {
		goto L20
	}
L19:
	;
	v83 = *(*int32)(unsafe.Add(mBase, uint32(v26)))
	v84 = v83
	goto L18
L20:
	;
	goto L15
L21:
	;
	v97 = *(*int32)(unsafe.Add(mBase, uint32(v26)+12))
	if v97 == int32(0) {
		goto L25
	} else {
		goto L26
	}
L22:
	;
	F_addReplyBulkCString(m, l0, int32(_a2255))
	mBase = m.M
	v93 = m.ExcPending
	if v93 != 0 {
		goto L1
	} else {
		goto L23
	}
L23:
	;
	v94 = int64(*(*int32)(unsafe.Add(mBase, uint32(v26)+8)))
	F_addReplyLongLong(m, l0, v94)
	mBase = m.M
	v96 = m.ExcPending
	if v96 != 0 {
		goto L1
	} else {
		goto L24
	}
L24:
	;
	goto L21
L25:
	;
	v106 = *(*int32)(unsafe.Add(mBase, uint32(v26)+16))
	if v106 == int32(0) {
		goto L29
	} else {
		goto L30
	}
L26:
	;
	F_addReplyBulkCString(m, l0, int32(_a2256))
	mBase = m.M
	v102 = m.ExcPending
	if v102 != 0 {
		goto L1
	} else {
		goto L27
	}
L27:
	;
	v103 = *(*int32)(unsafe.Add(mBase, uint32(v26)+12))
	F_addReplyBulkCString(m, l0, v103)
	mBase = m.M
	v105 = m.ExcPending
	if v105 != 0 {
		goto L1
	} else {
		goto L28
	}
L28:
	;
	goto L25
L29:
	;
	v115 = *(*int32)(unsafe.Add(mBase, uint32(v26)+20))
	if v115 == int32(0) {
		goto L33
	} else {
		goto L34
	}
L30:
	;
	F_addReplyBulkCString(m, l0, int32(_a2257))
	mBase = m.M
	v111 = m.ExcPending
	if v111 != 0 {
		goto L1
	} else {
		goto L31
	}
L31:
	;
	v112 = *(*int32)(unsafe.Add(mBase, uint32(v26)+16))
	F_addReplyBulkCString(m, l0, v112)
	mBase = m.M
	v114 = m.ExcPending
	if v114 != 0 {
		goto L1
	} else {
		goto L32
	}
L32:
	;
	goto L29
L33:
	;
	v124 = *(*int32)(unsafe.Add(mBase, uint32(v26)+28))
	if v124 == int32(0) {
		goto L37
	} else {
		goto L38
	}
L34:
	;
	F_addReplyBulkCString(m, l0, int32(_a2258))
	mBase = m.M
	v120 = m.ExcPending
	if v120 != 0 {
		goto L1
	} else {
		goto L35
	}
L35:
	;
	v121 = *(*int32)(unsafe.Add(mBase, uint32(v26)+20))
	F_addReplyBulkCString(m, l0, v121)
	mBase = m.M
	v123 = m.ExcPending
	if v123 != 0 {
		goto L1
	} else {
		goto L36
	}
L36:
	;
	goto L33
L37:
	;
	v133 = *(*int32)(unsafe.Add(mBase, uint32(v26)+24))
	if v133 == int32(0) {
		goto L41
	} else {
		goto L42
	}
L38:
	;
	F_addReplyBulkCString(m, l0, int32(_a2259))
	mBase = m.M
	v129 = m.ExcPending
	if v129 != 0 {
		goto L1
	} else {
		goto L39
	}
L39:
	;
	v130 = *(*int32)(unsafe.Add(mBase, uint32(v26)+28))
	F_addReplyBulkCString(m, l0, v130)
	mBase = m.M
	v132 = m.ExcPending
	if v132 != 0 {
		goto L1
	} else {
		goto L40
	}
L40:
	;
	goto L37
L41:
	;
	v174 = *(*int32)(unsafe.Add(mBase, uint32(v26)+4))
	if base.Ui32(int32(1)) < base.Ui32(v174+int32(-7)) {
		goto L53
	} else {
		goto L54
	}
L42:
	;
	F_addReplyBulkCString(m, l0, int32(_a62))
	mBase = m.M
	v138 = m.ExcPending
	if v138 != 0 {
		goto L1
	} else {
		goto L43
	}
L43:
	;
	v139 = *(*int32)(unsafe.Add(mBase, uint32(v26)+24))
	v140 = int32(1)
	v145 = v139 & v140
	v148 = v139 & int32(4)
	F_addReplySetLen(m, l0, int32(base.Ui32(v139)>>(uint(v140)%32))&v140+v145+int32(base.Ui32(v148)>>(uint(int32(2))%32)))
	mBase = m.M
	v153 = m.ExcPending
	if v153 != 0 {
		goto L1
	} else {
		goto L44
	}
L44:
	;
	if v145 == int32(0) {
		goto L45
	} else {
		goto L46
	}
L45:
	;
	if v139&int32(2) == int32(0) {
		goto L48
	} else {
		goto L49
	}
L46:
	;
	F_addReplyStatus(m, l0, int32(_a2260))
	mBase = m.M
	v158 = m.ExcPending
	if v158 != 0 {
		goto L1
	} else {
		goto L47
	}
L47:
	;
	goto L45
L48:
	;
	if v148 == int32(0) {
		goto L41
	} else {
		goto L51
	}
L49:
	;
	F_addReplyStatus(m, l0, int32(_a2261))
	mBase = m.M
	v165 = m.ExcPending
	if v165 != 0 {
		goto L1
	} else {
		goto L50
	}
L50:
	;
	goto L48
L51:
	;
	F_addReplyStatus(m, l0, int32(_a2262))
	mBase = m.M
	v170 = m.ExcPending
	if v170 != 0 {
		goto L1
	} else {
		goto L52
	}
L52:
	;
	goto L41
L53:
	;
	v187 = v17 + int32(1)
	if v187 != l2 {
		v17 = v187
		goto L5
	} else {
		goto L57
	}
L54:
	;
	F_addReplyBulkCString(m, l0, int32(_a2263))
	mBase = m.M
	v181 = m.ExcPending
	if v181 != 0 {
		goto L1
	} else {
		goto L55
	}
L55:
	;
	v182 = *(*int32)(unsafe.Add(mBase, uint32(v26)+36))
	v183 = *(*int32)(unsafe.Add(mBase, uint32(v26)+32))
	F_addReplyCommandArgList(m, l0, v182, v183)
	mBase = m.M
	v185 = m.ExcPending
	if v185 != 0 {
		goto L1
	} else {
		goto L56
	}
L56:
	;
	goto L53
L57:
	;
	goto L6
}
func F_addReplyCommandDocs(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v21 int32
	_ = v21
	var v25 int32
	_ = v25
	var v29 int32
	_ = v29
	var v33 int32
	_ = v33
	var v37 int32
	_ = v37
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v106 int32
	_ = v106
	var v107 int64
	_ = v107
	var v109 int64
	_ = v109
	var v110 int64
	_ = v110
	var v114 int64
	_ = v114
	var v118 int32
	_ = v118
	var v123 int32
	_ = v123
	var v128 int32
	_ = v128
	var v131 int32
	_ = v131
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v166 int32
	_ = v166
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v173 int32
	_ = v173
	var v175 int32
	_ = v175
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v180 int32
	_ = v180
	var v182 int32
	_ = v182
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v193 int32
	_ = v193
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v208 int32
	_ = v208
	var v212 int32
	_ = v212
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v9 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v10 = int32(2)
	goto L3
L2:
	;
	v10 = int32(1)
	goto L3
L3:
	;
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v12 = int32(0)
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l1)+56))
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v25 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	v29 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
	v33 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	v37 = *(*int32)(unsafe.Add(mBase, uint32(l1)+32))
	v41 = *(*int32)(unsafe.Add(mBase, uint32(l1)+96))
	v45 = *(*int32)(unsafe.Add(mBase, uint32(l1)+200))
	F_addReplyMapLen(m, l0, v10+base.B2i32(v11 != v12)+int32(base.Ui32(v15)>>(uint(int32(3))%32))&int32(1)+base.B2i32(v21 != v12)+base.B2i32(v25 != v12)+base.B2i32(v29 != v12)+base.B2i32(v33 != v12)+base.B2i32(v37 != v12)+base.B2i32(v41 != v12)+base.B2i32(v45 != v12))
	mBase = m.M
	v50 = m.ExcPending
	if v50 != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	return
L5:
	;
	v51 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v51 == int32(0) {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	v60 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	if v60 == int32(0) {
		goto L10
	} else {
		goto L11
	}
L7:
	;
	F_addReplyBulkCString(m, l0, int32(_a2257))
	mBase = m.M
	v56 = m.ExcPending
	if v56 != 0 {
		goto L4
	} else {
		goto L8
	}
L8:
	;
	v57 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	F_addReplyBulkCString(m, l0, v57)
	mBase = m.M
	v59 = m.ExcPending
	if v59 != 0 {
		goto L4
	} else {
		goto L9
	}
L9:
	;
	goto L6
L10:
	;
	F_addReplyBulkCString(m, l0, int32(_a569))
	mBase = m.M
	v71 = m.ExcPending
	if v71 != 0 {
		goto L4
	} else {
		goto L14
	}
L11:
	;
	F_addReplyBulkCString(m, l0, int32(_a2258))
	mBase = m.M
	v65 = m.ExcPending
	if v65 != 0 {
		goto L4
	} else {
		goto L12
	}
L12:
	;
	v66 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	F_addReplyBulkCString(m, l0, v66)
	mBase = m.M
	v68 = m.ExcPending
	if v68 != 0 {
		goto L4
	} else {
		goto L13
	}
L13:
	;
	goto L10
L14:
	;
	v72 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
	v77 = *(*int32)(unsafe.Add(mBase, uint32(v72<<(uint(int32(2))%32))+uint32(_consts[786])))
	goto L15
L15:
	;
	F_addReplyBulkCString(m, l0, v77)
	mBase = m.M
	v79 = m.ExcPending
	if v79 != 0 {
		goto L4
	} else {
		goto L16
	}
L16:
	;
	v80 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	if v80 == int32(0) {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	v89 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+56)))
	if v89&int32(8) == int32(0) {
		goto L21
	} else {
		goto L22
	}
L18:
	;
	F_addReplyBulkCString(m, l0, int32(_a2265))
	mBase = m.M
	v85 = m.ExcPending
	if v85 != 0 {
		goto L4
	} else {
		goto L19
	}
L19:
	;
	v86 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	F_addReplyBulkCString(m, l0, v86)
	mBase = m.M
	v88 = m.ExcPending
	if v88 != 0 {
		goto L4
	} else {
		goto L20
	}
L20:
	;
	goto L17
L21:
	;
	v101 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	if v101 == int32(0) {
		goto L26
	} else {
		goto L27
	}
L22:
	;
	F_addReplyBulkCString(m, l0, int32(_a332))
	mBase = m.M
	v96 = m.ExcPending
	if v96 != 0 {
		goto L4
	} else {
		goto L23
	}
L23:
	;
	v97 = F_moduleNameFromCommand(m, l1)
	mBase = m.M
	v98 = m.ExcPending
	if v98 != 0 {
		goto L4
	} else {
		goto L24
	}
L24:
	;
	F_addReplyBulkCString(m, l0, v97)
	mBase = m.M
	v100 = m.ExcPending
	if v100 != 0 {
		goto L4
	} else {
		goto L25
	}
L25:
	;
	goto L21
L26:
	;
	v131 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
	if v131 == int32(0) {
		goto L35
	} else {
		goto L36
	}
L27:
	;
	F_addReplyBulkCString(m, l0, int32(_a2266))
	mBase = m.M
	v106 = m.ExcPending
	if v106 != 0 {
		goto L4
	} else {
		goto L28
	}
L28:
	;
	v107 = int64(*(*uint32)(unsafe.Add(mBase, uint32(l1)+16)))
	v109 = v107 & int64(2)
	v110 = int64(1)
	v114 = v107 & v110
	F_addReplySetLen(m, l0, base.I32_wrap_i64(int64(base.Ui64(v109)>>(uint(v110)%64)))+base.I32_wrap_i64(v114))
	mBase = m.M
	v118 = m.ExcPending
	if v118 != 0 {
		goto L4
	} else {
		goto L29
	}
L29:
	;
	if v114 == int64(0) {
		goto L30
	} else {
		goto L31
	}
L30:
	;
	if v109 == int64(0) {
		goto L26
	} else {
		goto L33
	}
L31:
	;
	F_addReplyStatus(m, l0, int32(_a2267))
	mBase = m.M
	v123 = m.ExcPending
	if v123 != 0 {
		goto L4
	} else {
		goto L32
	}
L32:
	;
	goto L30
L33:
	;
	F_addReplyStatus(m, l0, int32(_a2268))
	mBase = m.M
	v128 = m.ExcPending
	if v128 != 0 {
		goto L4
	} else {
		goto L34
	}
L34:
	;
	goto L26
L35:
	;
	v140 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	if v140 == int32(0) {
		goto L39
	} else {
		goto L40
	}
L36:
	;
	F_addReplyBulkCString(m, l0, int32(_a2259))
	mBase = m.M
	v136 = m.ExcPending
	if v136 != 0 {
		goto L4
	} else {
		goto L37
	}
L37:
	;
	v137 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
	F_addReplyBulkCString(m, l0, v137)
	mBase = m.M
	v139 = m.ExcPending
	if v139 != 0 {
		goto L4
	} else {
		goto L38
	}
L38:
	;
	goto L35
L39:
	;
	v149 = *(*int32)(unsafe.Add(mBase, uint32(l1)+32))
	if v149 == int32(0) {
		goto L43
	} else {
		goto L44
	}
L40:
	;
	F_addReplyBulkCString(m, l0, int32(_a2269))
	mBase = m.M
	v145 = m.ExcPending
	if v145 != 0 {
		goto L4
	} else {
		goto L41
	}
L41:
	;
	v146 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	F_addReplyBulkCString(m, l0, v146)
	mBase = m.M
	v148 = m.ExcPending
	if v148 != 0 {
		goto L4
	} else {
		goto L42
	}
L42:
	;
	goto L39
L43:
	;
	v193 = *(*int32)(unsafe.Add(mBase, uint32(l1)+96))
	if v193 == int32(0) {
		goto L54
	} else {
		goto L55
	}
L44:
	;
	F_addReplyBulkCString(m, l0, int32(_a2270))
	mBase = m.M
	v154 = m.ExcPending
	if v154 != 0 {
		goto L4
	} else {
		goto L45
	}
L45:
	;
	v155 = *(*int32)(unsafe.Add(mBase, uint32(l1)+36))
	F_addReplySetLen(m, l0, v155)
	mBase = m.M
	v157 = m.ExcPending
	if v157 != 0 {
		goto L4
	} else {
		goto L46
	}
L46:
	;
	v158 = *(*int32)(unsafe.Add(mBase, uint32(l1)+36))
	if v158 < int32(1) {
		goto L43
	} else {
		goto L47
	}
L47:
	;
	v166 = int32(0)
	goto L48
L48:
	;
	F_addReplyArrayLen(m, l0, int32(2))
	mBase = m.M
	v170 = m.ExcPending
	if v170 != 0 {
		goto L4
	} else {
		goto L50
	}
L49:
	;
	goto L43
L50:
	;
	v171 = *(*int32)(unsafe.Add(mBase, uint32(l1)+32))
	v173 = v166 << (uint(int32(3)) % 32)
	v175 = *(*int32)(unsafe.Add(mBase, uint32(v171+v173)))
	F_addReplyBulkCString(m, l0, v175)
	mBase = m.M
	v177 = m.ExcPending
	if v177 != 0 {
		goto L4
	} else {
		goto L51
	}
L51:
	;
	v178 = *(*int32)(unsafe.Add(mBase, uint32(l1)+32))
	v180 = *(*int32)(unsafe.Add(mBase, uint32(v178+v173)+4))
	F_addReplyBulkCString(m, l0, v180)
	mBase = m.M
	v182 = m.ExcPending
	if v182 != 0 {
		goto L4
	} else {
		goto L52
	}
L52:
	;
	v184 = v166 + int32(1)
	v185 = *(*int32)(unsafe.Add(mBase, uint32(l1)+36))
	if v184 < v185 {
		v166 = v184
		goto L48
	} else {
		goto L53
	}
L53:
	;
	goto L49
L54:
	;
	v203 = *(*int32)(unsafe.Add(mBase, uint32(l1)+200))
	if v203 == int32(0) {
		goto L58
	} else {
		goto L59
	}
L55:
	;
	F_addReplyBulkCString(m, l0, int32(_a2263))
	mBase = m.M
	v198 = m.ExcPending
	if v198 != 0 {
		goto L4
	} else {
		goto L56
	}
L56:
	;
	v199 = *(*int32)(unsafe.Add(mBase, uint32(l1)+96))
	v200 = *(*int32)(unsafe.Add(mBase, uint32(l1)+88))
	F_addReplyCommandArgList(m, l0, v199, v200)
	mBase = m.M
	v202 = m.ExcPending
	if v202 != 0 {
		goto L4
	} else {
		goto L57
	}
L57:
	;
	goto L54
L58:
	;
	return
L59:
	;
	F_addReplyBulkCString(m, l0, int32(_a2271))
	mBase = m.M
	v208 = m.ExcPending
	if v208 != 0 {
		goto L4
	} else {
		goto L60
	}
L60:
	;
	F_addReplyCommandSubCommands(m, l0, l1, int32(1030), int32(1))
	mBase = m.M
	v212 = m.ExcPending
	if v212 != 0 {
		goto L4
	} else {
		goto L61
	}
L61:
	;
	goto L58
}
func F_addReplyCommandInfo(m *base.Module, l0 int32, l1 int32) {
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
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v38 int32
	_ = v38
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v85 int32
	_ = v85
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	var v97 int32
	_ = v97
	var v100 int32
	_ = v100
	var v104 int32
	_ = v104
	var v111 int32
	_ = v111
	var v114 int32
	_ = v114
	var v117 int32
	_ = v117
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v123 int32
	_ = v123
	var v136 int32
	_ = v136
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	if l1 != 0 {
		v13 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+224)))
		v18 = l1 + base.B2i32(v13 == int32(3))<<(uint(int32(2))%32)
		v19 = *(*int32)(unsafe.Add(mBase, uint32(v18)+212))
		if v19 != 0 {
			v25 = int32(0)
			v26 = *(*int32)(unsafe.Add(mBase, _consts[113]))
			if v26 == v25 {
				v100 = v19
				v104 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v100+int32(-1)))))
				switch v104 & int32(7) {
				case 0:
					v121 = int32(base.Ui32(v104) >> (uint(int32(3)) % 32))
				case 1:
					v111 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v100+int32(-3)))))
					v121 = v111
				case 2:
					v114 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v100+int32(-5)))))
					v121 = v114
				case 3:
					v117 = *(*int32)(unsafe.Add(mBase, uint32(v100+int32(-9))))
					v121 = v117
				case 4:
					v120 = *(*int32)(unsafe.Add(mBase, uint32(v100+int32(-17))))
					v121 = v120
				default:
					v121 = int32(0)
				}
				F_addReplyProto(m, l0, v100, v121)
				mBase = m.M
				v123 = m.ExcPending
				if v123 != 0 {
					return
				} else {
					m.G0 = v9 + int32(16)
					return
				}
			} else {
				v29 = F_generateCommandInfoResponse(m, l1, v13)
				mBase = m.M
				v30 = m.ExcPending
				if v30 != 0 {
					return
				} else {
					v31 = int32(0)
					v38 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v29+int32(-1)))))
					switch v38 & int32(7) {
					case 0:
						v55 = int32(base.Ui32(v38) >> (uint(int32(3)) % 32))
					case 1:
						v45 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v29+int32(-3)))))
						v55 = v45
					case 2:
						v48 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v29+int32(-5)))))
						v55 = v48
					case 3:
						v51 = *(*int32)(unsafe.Add(mBase, uint32(v29+int32(-9))))
						v55 = v51
					case 4:
						v54 = *(*int32)(unsafe.Add(mBase, uint32(v29+int32(-17))))
						v55 = v54
					default:
						v55 = v31
					}
					v58 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19+int32(-1)))))
					switch v58 & int32(7) {
					case 0:
						v75 = int32(base.Ui32(v58) >> (uint(int32(3)) % 32))
					case 1:
						v65 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19+int32(-3)))))
						v75 = v65
					case 2:
						v68 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v19+int32(-5)))))
						v75 = v68
					case 3:
						v71 = *(*int32)(unsafe.Add(mBase, uint32(v19+int32(-9))))
						v75 = v71
					case 4:
						v74 = *(*int32)(unsafe.Add(mBase, uint32(v19+int32(-17))))
						v75 = v74
					default:
						v75 = v31
					}
					v76 = base.B2i32(base.Ui32(v55) < base.Ui32(v75))
					if base.Ui32(v55) < base.Ui32(v75) {
						v77 = v55
					} else {
						v77 = v75
					}
					v78 = F_memcmp(m, v29, v19, v77)
					mBase = m.M
					if v78 != 0 {
						v81 = v78
					} else {
						v81 = base.B2i32(base.Ui32(v75) < base.Ui32(v55)) - v76
					}
					if v81 == int32(0) {
						F_sdsfree(m, v29)
						mBase = m.M
						v97 = m.ExcPending
						if v97 != 0 {
							return
						} else {
							if v81 != 0 {
								F__serverAssertWithInfo(m, l0, int32(0), int32(_a2264), int32(_a2157), int32(5408))
								mBase = m.M
								v136 = m.ExcPending
								if v136 != 0 {
									return
								} else {
									F_abort(m)
									mBase = m.M
									base.Wasm_trap_unreachable()
									for {
									}
								}
							} else {
								v100 = v19
								v104 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v100+int32(-1)))))
								switch v104 & int32(7) {
								case 0:
									v121 = int32(base.Ui32(v104) >> (uint(int32(3)) % 32))
								case 1:
									v111 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v100+int32(-3)))))
									v121 = v111
								case 2:
									v114 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v100+int32(-5)))))
									v121 = v114
								case 3:
									v117 = *(*int32)(unsafe.Add(mBase, uint32(v100+int32(-9))))
									v121 = v117
								case 4:
									v120 = *(*int32)(unsafe.Add(mBase, uint32(v100+int32(-17))))
									v121 = v120
								default:
									v121 = int32(0)
								}
								F_addReplyProto(m, l0, v100, v121)
								mBase = m.M
								v123 = m.ExcPending
								if v123 != 0 {
									return
								} else {
									m.G0 = v9 + int32(16)
									return
								}
							}
						}
					} else {
						v85 = *(*int32)(unsafe.Add(mBase, _consts[28]))
						if int32(3) < v85 {
							F_sdsfree(m, v29)
							mBase = m.M
							v97 = m.ExcPending
							if v97 != 0 {
								return
							} else {
								if v81 != 0 {
									F__serverAssertWithInfo(m, l0, int32(0), int32(_a2264), int32(_a2157), int32(5408))
									mBase = m.M
									v136 = m.ExcPending
									if v136 != 0 {
										return
									} else {
										F_abort(m)
										mBase = m.M
										base.Wasm_trap_unreachable()
										for {
										}
									}
								} else {
									v100 = v19
									v104 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v100+int32(-1)))))
									switch v104 & int32(7) {
									case 0:
										v121 = int32(base.Ui32(v104) >> (uint(int32(3)) % 32))
									case 1:
										v111 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v100+int32(-3)))))
										v121 = v111
									case 2:
										v114 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v100+int32(-5)))))
										v121 = v114
									case 3:
										v117 = *(*int32)(unsafe.Add(mBase, uint32(v100+int32(-9))))
										v121 = v117
									case 4:
										v120 = *(*int32)(unsafe.Add(mBase, uint32(v100+int32(-17))))
										v121 = v120
									default:
										v121 = int32(0)
									}
									F_addReplyProto(m, l0, v100, v121)
									mBase = m.M
									v123 = m.ExcPending
									if v123 != 0 {
										return
									} else {
										m.G0 = v9 + int32(16)
										return
									}
								}
							}
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v9)+4)) = v19
							*(*int32)(unsafe.Add(mBase, uint32(v9))) = v29
							F__serverLog(m, int32(3), int32(_a230), v9)
							mBase = m.M
							v93 = m.ExcPending
							if v93 != 0 {
								return
							} else {
								F_sdsfree(m, v29)
								mBase = m.M
								v95 = m.ExcPending
								if v95 != 0 {
									return
								} else {
									F__serverAssertWithInfo(m, l0, int32(0), int32(_a2264), int32(_a2157), int32(5408))
									mBase = m.M
									v136 = m.ExcPending
									if v136 != 0 {
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
			}
		} else {
			v22 = F_generateCommandInfoResponse(m, l1, v13)
			mBase = m.M
			v23 = m.ExcPending
			if v23 != 0 {
				return
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v18+int32(212)))) = v22
				v100 = v22
				v104 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v100+int32(-1)))))
				switch v104 & int32(7) {
				case 0:
					v121 = int32(base.Ui32(v104) >> (uint(int32(3)) % 32))
				case 1:
					v111 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v100+int32(-3)))))
					v121 = v111
				case 2:
					v114 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v100+int32(-5)))))
					v121 = v114
				case 3:
					v117 = *(*int32)(unsafe.Add(mBase, uint32(v100+int32(-9))))
					v121 = v117
				case 4:
					v120 = *(*int32)(unsafe.Add(mBase, uint32(v100+int32(-17))))
					v121 = v120
				default:
					v121 = int32(0)
				}
				F_addReplyProto(m, l0, v100, v121)
				mBase = m.M
				v123 = m.ExcPending
				if v123 != 0 {
					return
				} else {
					m.G0 = v9 + int32(16)
					return
				}
			}
		}
	} else {
		F_addReplyNull(m, l0)
		mBase = m.M
		v12 = m.ExcPending
		if v12 != 0 {
			return
		} else {
			m.G0 = v9 + int32(16)
			return
		}
	}
}
func F_addReplyDeferredLen(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v31 int32
	_ = v31
	var v40 int64
	_ = v40
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v60 int32
	_ = v60
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v87 int64
	_ = v87
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
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
	v9 = m.G0
	v11 = v9 - int32(16)
	m.G0 = v11
	v14 = F_prepareClientToWrite(m, l0)
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		return int32(0)
	} else {
		if v14 != 0 {
			v113 = int32(0)
			m.G0 = v11 + int32(16)
			return v113
		} else {
			v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+200))
			if v18&int32(1) != 0 {
				v40 = *(*int64)(unsafe.Add(mBase, uint32(l0)+384))
				v42 = base.B2i32(v40 == int64(-1))
				if v40 == int64(-1) {
					v43 = int32(132)
				} else {
					v43 = int32(376)
				}
				v45 = *(*int32)(unsafe.Add(mBase, uint32(l0+v43)))
				v46 = *(*int32)(unsafe.Add(mBase, uint32(v45)+4))
				if v46 == int32(0) {
					v96 = F_listAddNodeTail(m, v45, int32(0))
					mBase = m.M
					v97 = m.ExcPending
					if v97 != 0 {
						return int32(0)
					} else {
						v98 = *(*int32)(unsafe.Add(mBase, uint32(v45)+4))
						v113 = v98
						m.G0 = v11 + int32(16)
						return v113
					}
				} else {
					v49 = *(*int32)(unsafe.Add(mBase, uint32(v46)+8))
					if v49 == int32(0) {
						v96 = F_listAddNodeTail(m, v45, int32(0))
						mBase = m.M
						v97 = m.ExcPending
						if v97 != 0 {
							return int32(0)
						} else {
							v98 = *(*int32)(unsafe.Add(mBase, uint32(v45)+4))
							v113 = v98
							m.G0 = v11 + int32(16)
							return v113
						}
					} else {
						v52 = *(*int32)(unsafe.Add(mBase, uint32(v49)+4))
						if base.Ui32(int32(16383)) < base.Ui32(v52) {
							v96 = F_listAddNodeTail(m, v45, int32(0))
							mBase = m.M
							v97 = m.ExcPending
							if v97 != 0 {
								return int32(0)
							} else {
								v98 = *(*int32)(unsafe.Add(mBase, uint32(v45)+4))
								v113 = v98
								m.G0 = v11 + int32(16)
								return v113
							}
						} else {
							v55 = *(*int32)(unsafe.Add(mBase, uint32(v49)))
							if base.Ui32(v55-v52) <= base.Ui32(int32(base.Ui32(v55)>>(uint(int32(2))%32))) {
								v96 = F_listAddNodeTail(m, v45, int32(0))
								mBase = m.M
								v97 = m.ExcPending
								if v97 != 0 {
									return int32(0)
								} else {
									v98 = *(*int32)(unsafe.Add(mBase, uint32(v45)+4))
									v113 = v98
									m.G0 = v11 + int32(16)
									return v113
								}
							} else {
								v60 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+223)))
								if v60 == int32(1) {
									v96 = F_listAddNodeTail(m, v45, int32(0))
									mBase = m.M
									v97 = m.ExcPending
									if v97 != 0 {
										return int32(0)
									} else {
										v98 = *(*int32)(unsafe.Add(mBase, uint32(v45)+4))
										v113 = v98
										m.G0 = v11 + int32(16)
										return v113
									}
								} else {
									v65 = *(*int32)(unsafe.Add(mBase, uint32(l0)+144))
									if v49+int32(13) == v65 {
										v96 = F_listAddNodeTail(m, v45, int32(0))
										mBase = m.M
										v97 = m.ExcPending
										if v97 != 0 {
											return int32(0)
										} else {
											v98 = *(*int32)(unsafe.Add(mBase, uint32(v45)+4))
											v113 = v98
											m.G0 = v11 + int32(16)
											return v113
										}
									} else {
										v67 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v49)+12)))
										if v67&int32(1) != 0 {
											v96 = F_listAddNodeTail(m, v45, int32(0))
											mBase = m.M
											v97 = m.ExcPending
											if v97 != 0 {
												return int32(0)
											} else {
												v98 = *(*int32)(unsafe.Add(mBase, uint32(v45)+4))
												v113 = v98
												m.G0 = v11 + int32(16)
												return v113
											}
										} else {
											v74 = F_zrealloc_usable(m, v49, v52+int32(16), v11+int32(12))
											mBase = m.M
											v75 = m.ExcPending
											if v75 != 0 {
												return int32(0)
											} else {
												v76 = *(*int32)(unsafe.Add(mBase, uint32(v11)+12))
												v78 = v76 + int32(-16)
												*(*int32)(unsafe.Add(mBase, uint32(v74))) = v78
												if v40 == int64(-1) {
													v82 = int32(160)
												} else {
													v82 = int32(384)
												}
												v83 = l0 + v82
												v87 = *(*int64)(unsafe.Add(mBase, uint32(v83)))
												*(*int64)(unsafe.Add(mBase, uint32(v83))) = base.I64_extend_i32_u(v78) - base.I64_extend_i32_u(v55) + v87
												*(*int32)(unsafe.Add(mBase, uint32(v46)+8)) = v74
												v96 = F_listAddNodeTail(m, v45, int32(0))
												mBase = m.M
												v97 = m.ExcPending
												if v97 != 0 {
													return int32(0)
												} else {
													v98 = *(*int32)(unsafe.Add(mBase, uint32(v45)+4))
													v113 = v98
													m.G0 = v11 + int32(16)
													return v113
												}
											}
										}
									}
								}
							}
						}
					}
				}
			} else {
				if v18&int32(2) == int32(0) {
					if v18&int32(262144) != 0 {
					} else {
						v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)+216))
						if v31 == int32(0) {
						} else {
						}
					}
					v40 = *(*int64)(unsafe.Add(mBase, uint32(l0)+384))
					v42 = base.B2i32(v40 == int64(-1))
					if v40 == int64(-1) {
						v43 = int32(132)
					} else {
						v43 = int32(376)
					}
					v45 = *(*int32)(unsafe.Add(mBase, uint32(l0+v43)))
					v46 = *(*int32)(unsafe.Add(mBase, uint32(v45)+4))
					if v46 == int32(0) {
						v96 = F_listAddNodeTail(m, v45, int32(0))
						mBase = m.M
						v97 = m.ExcPending
						if v97 != 0 {
							return int32(0)
						} else {
							v98 = *(*int32)(unsafe.Add(mBase, uint32(v45)+4))
							v113 = v98
							m.G0 = v11 + int32(16)
							return v113
						}
					} else {
						v49 = *(*int32)(unsafe.Add(mBase, uint32(v46)+8))
						if v49 == int32(0) {
							v96 = F_listAddNodeTail(m, v45, int32(0))
							mBase = m.M
							v97 = m.ExcPending
							if v97 != 0 {
								return int32(0)
							} else {
								v98 = *(*int32)(unsafe.Add(mBase, uint32(v45)+4))
								v113 = v98
								m.G0 = v11 + int32(16)
								return v113
							}
						} else {
							v52 = *(*int32)(unsafe.Add(mBase, uint32(v49)+4))
							if base.Ui32(int32(16383)) < base.Ui32(v52) {
								v96 = F_listAddNodeTail(m, v45, int32(0))
								mBase = m.M
								v97 = m.ExcPending
								if v97 != 0 {
									return int32(0)
								} else {
									v98 = *(*int32)(unsafe.Add(mBase, uint32(v45)+4))
									v113 = v98
									m.G0 = v11 + int32(16)
									return v113
								}
							} else {
								v55 = *(*int32)(unsafe.Add(mBase, uint32(v49)))
								if base.Ui32(v55-v52) <= base.Ui32(int32(base.Ui32(v55)>>(uint(int32(2))%32))) {
									v96 = F_listAddNodeTail(m, v45, int32(0))
									mBase = m.M
									v97 = m.ExcPending
									if v97 != 0 {
										return int32(0)
									} else {
										v98 = *(*int32)(unsafe.Add(mBase, uint32(v45)+4))
										v113 = v98
										m.G0 = v11 + int32(16)
										return v113
									}
								} else {
									v60 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+223)))
									if v60 == int32(1) {
										v96 = F_listAddNodeTail(m, v45, int32(0))
										mBase = m.M
										v97 = m.ExcPending
										if v97 != 0 {
											return int32(0)
										} else {
											v98 = *(*int32)(unsafe.Add(mBase, uint32(v45)+4))
											v113 = v98
											m.G0 = v11 + int32(16)
											return v113
										}
									} else {
										v65 = *(*int32)(unsafe.Add(mBase, uint32(l0)+144))
										if v49+int32(13) == v65 {
											v96 = F_listAddNodeTail(m, v45, int32(0))
											mBase = m.M
											v97 = m.ExcPending
											if v97 != 0 {
												return int32(0)
											} else {
												v98 = *(*int32)(unsafe.Add(mBase, uint32(v45)+4))
												v113 = v98
												m.G0 = v11 + int32(16)
												return v113
											}
										} else {
											v67 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v49)+12)))
											if v67&int32(1) != 0 {
												v96 = F_listAddNodeTail(m, v45, int32(0))
												mBase = m.M
												v97 = m.ExcPending
												if v97 != 0 {
													return int32(0)
												} else {
													v98 = *(*int32)(unsafe.Add(mBase, uint32(v45)+4))
													v113 = v98
													m.G0 = v11 + int32(16)
													return v113
												}
											} else {
												v74 = F_zrealloc_usable(m, v49, v52+int32(16), v11+int32(12))
												mBase = m.M
												v75 = m.ExcPending
												if v75 != 0 {
													return int32(0)
												} else {
													v76 = *(*int32)(unsafe.Add(mBase, uint32(v11)+12))
													v78 = v76 + int32(-16)
													*(*int32)(unsafe.Add(mBase, uint32(v74))) = v78
													if v40 == int64(-1) {
														v82 = int32(160)
													} else {
														v82 = int32(384)
													}
													v83 = l0 + v82
													v87 = *(*int64)(unsafe.Add(mBase, uint32(v83)))
													*(*int64)(unsafe.Add(mBase, uint32(v83))) = base.I64_extend_i32_u(v78) - base.I64_extend_i32_u(v55) + v87
													*(*int32)(unsafe.Add(mBase, uint32(v46)+8)) = v74
													v96 = F_listAddNodeTail(m, v45, int32(0))
													mBase = m.M
													v97 = m.ExcPending
													if v97 != 0 {
														return int32(0)
													} else {
														v98 = *(*int32)(unsafe.Add(mBase, uint32(v45)+4))
														v113 = v98
														m.G0 = v11 + int32(16)
														return v113
													}
												}
											}
										}
									}
								}
							}
						}
					}
				} else {
					if v18&int32(4) == int32(0) {
						v99 = int32(0)
						v101 = *(*int32)(unsafe.Add(mBase, uint32(l0)+72))
						if v101 == v99 {
							v105 = v99
						} else {
							v104 = *(*int32)(unsafe.Add(mBase, uint32(v101)+140))
							v105 = v104
						}
						if v105 != 0 {
							v107 = v105
						} else {
							v107 = int32(_a277)
						}
						*(*int32)(unsafe.Add(mBase, uint32(v11))) = v107
						F_logInvalidUseAndFreeClientAsync(m, l0, int32(_a1642), v11)
						mBase = m.M
						v111 = m.ExcPending
						if v111 != 0 {
							return int32(0)
						} else {
							v113 = v99
							m.G0 = v11 + int32(16)
							return v113
						}
					} else {
						if v18&int32(262144) != 0 {
						} else {
							v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)+216))
							if v31 == int32(0) {
							} else {
							}
						}
						v40 = *(*int64)(unsafe.Add(mBase, uint32(l0)+384))
						v42 = base.B2i32(v40 == int64(-1))
						if v40 == int64(-1) {
							v43 = int32(132)
						} else {
							v43 = int32(376)
						}
						v45 = *(*int32)(unsafe.Add(mBase, uint32(l0+v43)))
						v46 = *(*int32)(unsafe.Add(mBase, uint32(v45)+4))
						if v46 == int32(0) {
							v96 = F_listAddNodeTail(m, v45, int32(0))
							mBase = m.M
							v97 = m.ExcPending
							if v97 != 0 {
								return int32(0)
							} else {
								v98 = *(*int32)(unsafe.Add(mBase, uint32(v45)+4))
								v113 = v98
								m.G0 = v11 + int32(16)
								return v113
							}
						} else {
							v49 = *(*int32)(unsafe.Add(mBase, uint32(v46)+8))
							if v49 == int32(0) {
								v96 = F_listAddNodeTail(m, v45, int32(0))
								mBase = m.M
								v97 = m.ExcPending
								if v97 != 0 {
									return int32(0)
								} else {
									v98 = *(*int32)(unsafe.Add(mBase, uint32(v45)+4))
									v113 = v98
									m.G0 = v11 + int32(16)
									return v113
								}
							} else {
								v52 = *(*int32)(unsafe.Add(mBase, uint32(v49)+4))
								if base.Ui32(int32(16383)) < base.Ui32(v52) {
									v96 = F_listAddNodeTail(m, v45, int32(0))
									mBase = m.M
									v97 = m.ExcPending
									if v97 != 0 {
										return int32(0)
									} else {
										v98 = *(*int32)(unsafe.Add(mBase, uint32(v45)+4))
										v113 = v98
										m.G0 = v11 + int32(16)
										return v113
									}
								} else {
									v55 = *(*int32)(unsafe.Add(mBase, uint32(v49)))
									if base.Ui32(v55-v52) <= base.Ui32(int32(base.Ui32(v55)>>(uint(int32(2))%32))) {
										v96 = F_listAddNodeTail(m, v45, int32(0))
										mBase = m.M
										v97 = m.ExcPending
										if v97 != 0 {
											return int32(0)
										} else {
											v98 = *(*int32)(unsafe.Add(mBase, uint32(v45)+4))
											v113 = v98
											m.G0 = v11 + int32(16)
											return v113
										}
									} else {
										v60 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+223)))
										if v60 == int32(1) {
											v96 = F_listAddNodeTail(m, v45, int32(0))
											mBase = m.M
											v97 = m.ExcPending
											if v97 != 0 {
												return int32(0)
											} else {
												v98 = *(*int32)(unsafe.Add(mBase, uint32(v45)+4))
												v113 = v98
												m.G0 = v11 + int32(16)
												return v113
											}
										} else {
											v65 = *(*int32)(unsafe.Add(mBase, uint32(l0)+144))
											if v49+int32(13) == v65 {
												v96 = F_listAddNodeTail(m, v45, int32(0))
												mBase = m.M
												v97 = m.ExcPending
												if v97 != 0 {
													return int32(0)
												} else {
													v98 = *(*int32)(unsafe.Add(mBase, uint32(v45)+4))
													v113 = v98
													m.G0 = v11 + int32(16)
													return v113
												}
											} else {
												v67 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v49)+12)))
												if v67&int32(1) != 0 {
													v96 = F_listAddNodeTail(m, v45, int32(0))
													mBase = m.M
													v97 = m.ExcPending
													if v97 != 0 {
														return int32(0)
													} else {
														v98 = *(*int32)(unsafe.Add(mBase, uint32(v45)+4))
														v113 = v98
														m.G0 = v11 + int32(16)
														return v113
													}
												} else {
													v74 = F_zrealloc_usable(m, v49, v52+int32(16), v11+int32(12))
													mBase = m.M
													v75 = m.ExcPending
													if v75 != 0 {
														return int32(0)
													} else {
														v76 = *(*int32)(unsafe.Add(mBase, uint32(v11)+12))
														v78 = v76 + int32(-16)
														*(*int32)(unsafe.Add(mBase, uint32(v74))) = v78
														if v40 == int64(-1) {
															v82 = int32(160)
														} else {
															v82 = int32(384)
														}
														v83 = l0 + v82
														v87 = *(*int64)(unsafe.Add(mBase, uint32(v83)))
														*(*int64)(unsafe.Add(mBase, uint32(v83))) = base.I64_extend_i32_u(v78) - base.I64_extend_i32_u(v55) + v87
														*(*int32)(unsafe.Add(mBase, uint32(v46)+8)) = v74
														v96 = F_listAddNodeTail(m, v45, int32(0))
														mBase = m.M
														v97 = m.ExcPending
														if v97 != 0 {
															return int32(0)
														} else {
															v98 = *(*int32)(unsafe.Add(mBase, uint32(v45)+4))
															v113 = v98
															m.G0 = v11 + int32(16)
															return v113
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
func F_addReplyErrorFormat(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v12 int32
	_ = v12
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	*(*int32)(unsafe.Add(mBase, uint32(v7)+12)) = l2
	F_addReplyErrorFormatInternal(m, l0, int32(0), l1, l2)
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		return
	} else {
		m.G0 = v7 + int32(16)
		return
	}
}
func F_addReplyErrorFormatInternal(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v17 int32
	_ = v17
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
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v61 int32
	_ = v61
	var v67 int32
	_ = v67
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v79 int32
	_ = v79
	var v83 int32
	_ = v83
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v98 int32
	_ = v98
	var v123 int32
	_ = v123
	var v130 int32
	_ = v130
	var v133 int32
	_ = v133
	var v136 int32
	_ = v136
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v150 int32
	_ = v150
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v164 int32
	_ = v164
	var v170 int32
	_ = v170
	var v173 int32
	_ = v173
	var v176 int32
	_ = v176
	var v188 int32
	_ = v188
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v208 int32
	_ = v208
	var v211 int32
	_ = v211
	var v214 int32
	_ = v214
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v228 int32
	_ = v228
	var v231 int32
	_ = v231
	var v234 int32
	_ = v234
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v240 int32
	_ = v240
	var v242 int32
	_ = v242
	v5 = int32(0)
	v9 = m.G0
	v11 = v9 - int32(16)
	m.G0 = v11
	*(*int32)(unsafe.Add(mBase, uint32(v11)+12)) = l3
	*(*int32)(unsafe.Add(mBase, uint32(v11)+8)) = l3
	v17 = F_sdsempty(m)
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		goto L7
	} else {
		goto L8
	}
L1:
	;
	F_addReplyErrorLength(m, l0, v20, v218)
	mBase = m.M
	v220 = m.ExcPending
	if v220 != 0 {
		goto L7
	} else {
		goto L56
	}
L2:
	;
	v217 = *(*int32)(unsafe.Add(mBase, uint32(v20+int32(-17))))
	v218 = v217
	goto L1
L3:
	;
	v214 = *(*int32)(unsafe.Add(mBase, uint32(v20+int32(-9))))
	v218 = v214
	goto L1
L4:
	;
	v211 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v20+int32(-5)))))
	v218 = v211
	goto L1
L5:
	;
	v208 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20+int32(-3)))))
	v218 = v208
	goto L1
L6:
	;
	v218 = int32(base.Ui32(v201) >> (uint(int32(3)) % 32))
	goto L1
L7:
	;
	return
L8:
	;
	v19 = *(*int32)(unsafe.Add(mBase, uint32(v11)+8))
	v20 = F_sdscatvprintf(m, v17, l2, v19)
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		goto L7
	} else {
		goto L9
	}
L9:
	;
	v22 = int32(_a823)
	v28 = v20 + int32(-1)
	v29 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28))))
	switch v29 & int32(7) {
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
		v46 = int32(0)
		goto L11
	}
L10:
	;
	v123 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20+int32(-1)))))
	switch v123 & int32(7) {
	case 0:
		goto L44
	case 1:
		goto L43
	case 2:
		goto L42
	case 3:
		goto L41
	case 4:
		goto L40
	default:
		goto L38
	}
L11:
	;
	v49 = v20 + v46 + int32(-1)
	if base.Ui32(v49) < base.Ui32(v20) {
		v67 = v20
		goto L17
	} else {
		goto L18
	}
L12:
	;
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v20+int32(-17))))
	v46 = v45
	goto L11
L13:
	;
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v20+int32(-9))))
	v46 = v42
	goto L11
L14:
	;
	v39 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v20+int32(-5)))))
	v46 = v39
	goto L11
L15:
	;
	v36 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20+int32(-3)))))
	v46 = v36
	goto L11
L16:
	;
	v46 = int32(base.Ui32(v29) >> (uint(int32(3)) % 32))
	goto L11
L17:
	;
	if base.Ui32(v49) <= base.Ui32(v67) {
		v83 = v49
		goto L23
	} else {
		goto L24
	}
L18:
	;
	v55 = v20
	goto L19
L19:
	;
	v56 = int32(*(*int8)(unsafe.Add(mBase, uint32(v55))))
	v57 = F_strchr(m, v22, v56)
	mBase = m.M
	if v57 == int32(0) {
		v67 = v55
		goto L17
	} else {
		goto L21
	}
L20:
	;
	v67 = v61
	goto L17
L21:
	;
	v61 = v55 + int32(1)
	if base.Ui32(v61) <= base.Ui32(v49) {
		v55 = v61
		goto L19
	} else {
		goto L22
	}
L22:
	;
	goto L20
L23:
	;
	v88 = v83 - v67 + int32(1)
	if v20 == v67 {
		goto L29
	} else {
		goto L30
	}
L24:
	;
	v71 = v49
	goto L25
L25:
	;
	v74 = int32(*(*int8)(unsafe.Add(mBase, uint32(v71))))
	v75 = F_strchr(m, v22, v74)
	mBase = m.M
	if v75 == int32(0) {
		v83 = v71
		goto L23
	} else {
		goto L27
	}
L26:
	;
	v83 = v67
	goto L23
L27:
	;
	v79 = v71 + int32(-1)
	if base.Ui32(v67) < base.Ui32(v79) {
		v71 = v79
		goto L25
	} else {
		goto L28
	}
L28:
	;
	goto L26
L29:
	;
	v92 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v20+v88))) = uint8(v92)
	v94 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28))))
	switch v94 & int32(7) {
	case 0:
		goto L36
	case 1:
		goto L35
	case 2:
		goto L34
	case 3:
		goto L33
	case 4:
		goto L32
	default:
		goto L31
	}
L30:
	;
	v90 = F_memmove(m, v20, v67, v88)
	mBase = m.M
	goto L29
L31:
	;
	goto L10
L32:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v20+int32(-17)))) = base.I64_extend_i32_u(v88)
	goto L31
L33:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20+int32(-9)))) = v88
	goto L10
L34:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v20+int32(-5)))) = uint16(v88)
	goto L10
L35:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v20+int32(-3)))) = uint8(v88)
	goto L10
L36:
	;
	v98 = v88 << (uint(int32(3)) % 32)
	*(*uint8)(unsafe.Add(mBase, uint32(v28))) = uint8(v98)
	goto L10
L37:
	;
	v200 = v20 + int32(-1)
	v201 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v200))))
	switch v201 & int32(7) {
	case 0:
		goto L6
	case 1:
		goto L5
	case 2:
		goto L4
	case 3:
		goto L3
	case 4:
		goto L2
	default:
		v218 = v5
		goto L1
	}
L38:
	;
	goto L37
L39:
	;
	if v140 == int32(0) {
		goto L38
	} else {
		goto L45
	}
L40:
	;
	v139 = *(*int32)(unsafe.Add(mBase, uint32(v20+int32(-17))))
	v140 = v139
	goto L39
L41:
	;
	v136 = *(*int32)(unsafe.Add(mBase, uint32(v20+int32(-9))))
	v140 = v136
	goto L39
L42:
	;
	v133 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v20+int32(-5)))))
	v140 = v133
	goto L39
L43:
	;
	v130 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20+int32(-3)))))
	v140 = v130
	goto L39
L44:
	;
	v140 = int32(base.Ui32(v123) >> (uint(int32(3)) % 32))
	goto L39
L45:
	;
	v150 = int32(0)
	goto L46
L46:
	;
	goto L49
L47:
	;
	goto L38
L48:
	;
	v188 = v150 + int32(1)
	if v188 != v140 {
		v150 = v188
		goto L46
	} else {
		goto L55
	}
L49:
	;
	v156 = v20 + v150
	v157 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v156))))
	v164 = int32(0)
	goto L50
L50:
	;
	v170 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v164)+uint32(_consts[486]))))
	if v157&int32(255) != v170 {
		goto L52
	} else {
		goto L53
	}
L51:
	;
	goto L48
L52:
	;
	v176 = v164 + int32(1)
	if v176 != int32(2) {
		v164 = v176
		goto L50
	} else {
		goto L54
	}
L53:
	;
	v173 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v164)+uint32(_consts[487]))))
	*(*uint8)(unsafe.Add(mBase, uint32(v156))) = uint8(v173)
	goto L48
L54:
	;
	goto L51
L55:
	;
	goto L47
L56:
	;
	v221 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v200))))
	switch v221 & int32(7) {
	case 0:
		goto L62
	case 1:
		goto L61
	case 2:
		goto L60
	case 3:
		goto L59
	case 4:
		goto L58
	default:
		v238 = v5
		goto L57
	}
L57:
	;
	F_afterErrorReply(m, l0, v20, v238, l1)
	mBase = m.M
	v240 = m.ExcPending
	if v240 != 0 {
		goto L7
	} else {
		goto L63
	}
L58:
	;
	v237 = *(*int32)(unsafe.Add(mBase, uint32(v20+int32(-17))))
	v238 = v237
	goto L57
L59:
	;
	v234 = *(*int32)(unsafe.Add(mBase, uint32(v20+int32(-9))))
	v238 = v234
	goto L57
L60:
	;
	v231 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v20+int32(-5)))))
	v238 = v231
	goto L57
L61:
	;
	v228 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20+int32(-3)))))
	v238 = v228
	goto L57
L62:
	;
	v238 = int32(base.Ui32(v221) >> (uint(int32(3)) % 32))
	goto L57
L63:
	;
	F_sdsfree(m, v20)
	mBase = m.M
	v242 = m.ExcPending
	if v242 != 0 {
		goto L7
	} else {
		goto L64
	}
L64:
	;
	m.G0 = v11 + int32(16)
	return
}
func F_addReplyHelp(m *base.Module, l0 int32, l1 int32) {
	var v5 int32
	_ = v5
	F_addExtendedReplyHelp(m, l0, l1, int32(0))
	v5 = m.ExcPending
	if v5 != 0 {
		return
	} else {
		return
	}
}
func F_addReplyHumanLongDouble(m *base.Module, l0 int32, l1 int64, l2 int64) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
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
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
	v6 = m.G0
	v8 = v6 - int32(5120)
	m.G0 = v8
	v10 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+224)))
	if v10 != int32(2) {
		v22 = F_ld2string(m, v8, int32(5120), l1, l2, int32(1))
		mBase = m.M
		v23 = m.ExcPending
		if v23 != 0 {
			return
		} else {
			v24 = F_prepareClientToWrite(m, l0)
			mBase = m.M
			v25 = m.ExcPending
			if v25 != 0 {
				return
			} else {
				if v24 != 0 {
					v30 = F_prepareClientToWrite(m, l0)
					mBase = m.M
					v31 = m.ExcPending
					if v31 != 0 {
						return
					} else {
						if v30 != 0 {
							v34 = F_prepareClientToWrite(m, l0)
							mBase = m.M
							v35 = m.ExcPending
							if v35 != 0 {
								return
							} else {
								if v34 != 0 {
									m.G0 = v8 + int32(5120)
									return
								} else {
									F__addReplyToBufferOrList(m, l0, int32(_a823), int32(2))
									mBase = m.M
									v39 = m.ExcPending
									if v39 != 0 {
										return
									} else {
										m.G0 = v8 + int32(5120)
										return
									}
								}
							}
						} else {
							F__addReplyToBufferOrList(m, l0, v8, v22)
							mBase = m.M
							v33 = m.ExcPending
							if v33 != 0 {
								return
							} else {
								v34 = F_prepareClientToWrite(m, l0)
								mBase = m.M
								v35 = m.ExcPending
								if v35 != 0 {
									return
								} else {
									if v34 != 0 {
										m.G0 = v8 + int32(5120)
										return
									} else {
										F__addReplyToBufferOrList(m, l0, int32(_a823), int32(2))
										mBase = m.M
										v39 = m.ExcPending
										if v39 != 0 {
											return
										} else {
											m.G0 = v8 + int32(5120)
											return
										}
									}
								}
							}
						}
					}
				} else {
					F__addReplyToBufferOrList(m, l0, int32(_a15), int32(1))
					mBase = m.M
					v29 = m.ExcPending
					if v29 != 0 {
						return
					} else {
						v30 = F_prepareClientToWrite(m, l0)
						mBase = m.M
						v31 = m.ExcPending
						if v31 != 0 {
							return
						} else {
							if v30 != 0 {
								v34 = F_prepareClientToWrite(m, l0)
								mBase = m.M
								v35 = m.ExcPending
								if v35 != 0 {
									return
								} else {
									if v34 != 0 {
										m.G0 = v8 + int32(5120)
										return
									} else {
										F__addReplyToBufferOrList(m, l0, int32(_a823), int32(2))
										mBase = m.M
										v39 = m.ExcPending
										if v39 != 0 {
											return
										} else {
											m.G0 = v8 + int32(5120)
											return
										}
									}
								}
							} else {
								F__addReplyToBufferOrList(m, l0, v8, v22)
								mBase = m.M
								v33 = m.ExcPending
								if v33 != 0 {
									return
								} else {
									v34 = F_prepareClientToWrite(m, l0)
									mBase = m.M
									v35 = m.ExcPending
									if v35 != 0 {
										return
									} else {
										if v34 != 0 {
											m.G0 = v8 + int32(5120)
											return
										} else {
											F__addReplyToBufferOrList(m, l0, int32(_a823), int32(2))
											mBase = m.M
											v39 = m.ExcPending
											if v39 != 0 {
												return
											} else {
												m.G0 = v8 + int32(5120)
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
	} else {
		v14 = F_createStringObjectFromLongDouble(m, l1, l2, int32(1))
		mBase = m.M
		v15 = m.ExcPending
		if v15 != 0 {
			return
		} else {
			F_addReplyBulk(m, l0, v14)
			mBase = m.M
			v17 = m.ExcPending
			if v17 != 0 {
				return
			} else {
				F_decrRefCount(m, v14)
				mBase = m.M
				v19 = m.ExcPending
				if v19 != 0 {
					return
				} else {
					m.G0 = v8 + int32(5120)
					return
				}
			}
		}
	}
}
func F_addReplyLoadedModules(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v52 int64
	_ = v52
	var v53 int64
	_ = v53
	var v54 int64
	_ = v54
	var v55 int64
	_ = v55
	var v56 int64
	_ = v56
	var v57 int64
	_ = v57
	var v58 int64
	_ = v58
	var v60 int64
	_ = v60
	var v62 int64
	_ = v62
	var v64 int64
	_ = v64
	var v66 int64
	_ = v66
	var v68 int64
	_ = v68
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v90 int32
	_ = v90
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v110 int32
	_ = v110
	var v115 int32
	_ = v115
	var v119 int32
	_ = v119
	var v123 int32
	_ = v123
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
	var v143 int32
	_ = v143
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v151 int32
	_ = v151
	var v158 int32
	_ = v158
	var v161 int32
	_ = v161
	var v164 int32
	_ = v164
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v170 int32
	_ = v170
	var v173 int32
	_ = v173
	var v174 int64
	_ = v174
	var v176 int32
	_ = v176
	var v179 int32
	_ = v179
	var v182 int32
	_ = v182
	var v189 int32
	_ = v189
	var v192 int32
	_ = v192
	var v195 int32
	_ = v195
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v201 int32
	_ = v201
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v208 int32
	_ = v208
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v222 int32
	_ = v222
	var v226 int32
	_ = v226
	var v228 int32
	_ = v228
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v253 int32
	_ = v253
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v261 int32
	_ = v261
	var v265 int32
	_ = v265
	var v266 int32
	_ = v266
	var v267 int32
	_ = v267
	var v270 int32
	_ = v270
	var v272 int32
	_ = v272
	var v274 int64
	_ = v274
	var v275 int64
	_ = v275
	var v276 int64
	_ = v276
	var v277 int64
	_ = v277
	var v278 int64
	_ = v278
	var v279 int64
	_ = v279
	var v280 int64
	_ = v280
	var v282 int64
	_ = v282
	var v284 int64
	_ = v284
	var v286 int64
	_ = v286
	var v288 int64
	_ = v288
	var v290 int64
	_ = v290
	var v292 int32
	_ = v292
	var v293 int32
	_ = v293
	var v294 int32
	_ = v294
	var v297 int32
	_ = v297
	var v300 int32
	_ = v300
	var v301 int32
	_ = v301
	var v302 int32
	_ = v302
	var v304 int32
	_ = v304
	var v307 int32
	_ = v307
	var v308 int32
	_ = v308
	var v312 int32
	_ = v312
	var v316 int32
	_ = v316
	var v318 int32
	_ = v318
	var v325 int32
	_ = v325
	var v326 int32
	_ = v326
	var v327 int32
	_ = v327
	var v332 int32
	_ = v332
	var v337 int32
	_ = v337
	var v341 int32
	_ = v341
	var v345 int32
	_ = v345
	var v358 int32
	_ = v358
	v10 = *(*int32)(unsafe.Add(mBase, _consts[81]))
	v11 = F_dictGetIterator(m, v10)
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v14 = *(*int32)(unsafe.Add(mBase, _consts[81]))
	v15 = *(*int32)(unsafe.Add(mBase, uint32(v14)+16))
	v16 = *(*int32)(unsafe.Add(mBase, uint32(v14)+12))
	F_addReplyArrayLen(m, l0, v15+v16)
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v27 = v11 + int32(20)
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v11)+16))
	if v28 != 0 {
		goto L7
	} else {
		goto L8
	}
L4:
	;
	F_dictReleaseIterator(m, v11)
	mBase = m.M
	v358 = m.ExcPending
	if v358 != 0 {
		goto L1
	} else {
		goto L90
	}
L5:
	;
	if v123 == int32(0) {
		goto L4
	} else {
		goto L31
	}
L6:
	;
	v34 = v27
	v35 = v31
	goto L9
L7:
	;
	v31 = int32(1)
	goto L6
L8:
	;
	v31 = int32(0)
	goto L6
L9:
	;
	switch v35 {
	case 0:
		goto L14
	default:
		goto L13
	}
L11:
	;
	v35 = int32(0)
	goto L9
L12:
	;
	goto L5
L13:
	;
	v115 = *(*int32)(unsafe.Add(mBase, uint32(v34)))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = v115
	if v115 == int32(0) {
		goto L11
	} else {
		goto L30
	}
L14:
	;
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
	if v39 != int32(-1) {
		v78 = v39
		goto L15
	} else {
		goto L16
	}
L15:
	;
	v79 = int32(1)
	v80 = v78 + v79
	*(*int32)(unsafe.Add(mBase, uint32(v11)+4)) = v80
	v82 = int32(0)
	v85 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v11)+8))
	v90 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v85+v86+int32(26)))))
	if v90 == int32(255) {
		goto L24
	} else {
		goto L25
	}
L16:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v11)+8))
	if v43 != 0 {
		v78 = int32(-1)
		goto L15
	} else {
		goto L17
	}
L17:
	;
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v11)+12))
	if v45 == int32(0) {
		goto L19
	} else {
		goto L20
	}
L18:
	;
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v71)+20))
	if v72 != int32(-1) {
		goto L21
	} else {
		goto L22
	}
L19:
	;
	v52 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v44)+16)))
	v53 = int64(*(*int8)(unsafe.Add(mBase, uint32(v44)+27)))
	v54 = int64(*(*int32)(unsafe.Add(mBase, uint32(v44)+8)))
	v55 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v44)+12)))
	v56 = int64(*(*int8)(unsafe.Add(mBase, uint32(v44)+26)))
	v57 = int64(*(*int32)(unsafe.Add(mBase, uint32(v44)+4)))
	v58 = F_wangHash64(m, v57)
	mBase = m.M
	v60 = F_wangHash64(m, v56+v58)
	mBase = m.M
	v62 = F_wangHash64(m, v55+v60)
	mBase = m.M
	v64 = F_wangHash64(m, v54+v62)
	mBase = m.M
	v66 = F_wangHash64(m, v53+v64)
	mBase = m.M
	v68 = F_wangHash64(m, v52+v66)
	mBase = m.M
	*(*int64)(unsafe.Add(mBase, uint32(v11)+24)) = v68
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
	v71 = v70
	goto L18
L20:
	;
	v48 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v44)+24)))
	v50 = v48 + int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+24)) = uint16(v50)
	v71 = v44
	goto L18
L21:
	;
	v78 = v72 + int32(-1)
	goto L15
L22:
	;
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
	v78 = v75
	goto L15
L23:
	;
	v105 = int32(2)
	v110 = *(*int32)(unsafe.Add(mBase, uint32(v85+v103<<(uint(v105)%32)+int32(4))))
	v34 = v110 + v104<<(uint(v105)%32)
	v35 = int32(1)
	goto L9
L24:
	;
	v94 = v82
	goto L26
L25:
	;
	v94 = v79 << (uint(v90) % 32)
	goto L26
L26:
	;
	if v80 < v94 {
		v103 = v86
		v104 = v80
		goto L23
	} else {
		goto L27
	}
L27:
	;
	if v86 != 0 {
		v123 = v82
		goto L12
	} else {
		goto L28
	}
L28:
	;
	v96 = *(*int32)(unsafe.Add(mBase, uint32(v85)+20))
	if v96 == int32(-1) {
		v123 = v82
		goto L12
	} else {
		goto L29
	}
L29:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v11)+4)) = int64(4294967296)
	v103 = int32(1)
	v104 = int32(0)
	goto L23
L30:
	;
	v119 = *(*int32)(unsafe.Add(mBase, uint32(v115)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v27))) = v119
	v123 = v115
	goto L12
L31:
	;
	v132 = v123
	goto L32
L32:
	;
	v137 = *(*int32)(unsafe.Add(mBase, uint32(v132)))
	goto L34
L33:
	;
	goto L4
L34:
	;
	v138 = *(*int32)(unsafe.Add(mBase, uint32(v132)+8))
	goto L35
L35:
	;
	v139 = *(*int32)(unsafe.Add(mBase, uint32(v138)+64))
	v140 = *(*int32)(unsafe.Add(mBase, uint32(v139)))
	F_addReplyMapLen(m, l0, int32(4))
	mBase = m.M
	v143 = m.ExcPending
	if v143 != 0 {
		goto L1
	} else {
		goto L36
	}
L36:
	;
	F_addReplyBulkCString(m, l0, int32(_a373))
	mBase = m.M
	v146 = m.ExcPending
	if v146 != 0 {
		goto L1
	} else {
		goto L37
	}
L37:
	;
	v147 = int32(0)
	v151 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v137+int32(-1)))))
	switch v151 & int32(7) {
	case 0:
		goto L43
	case 1:
		goto L42
	case 2:
		goto L41
	case 3:
		goto L40
	case 4:
		goto L39
	default:
		v168 = v147
		goto L38
	}
L38:
	;
	F_addReplyBulkCBuffer(m, l0, v137, v168)
	mBase = m.M
	v170 = m.ExcPending
	if v170 != 0 {
		goto L1
	} else {
		goto L44
	}
L39:
	;
	v167 = *(*int32)(unsafe.Add(mBase, uint32(v137+int32(-17))))
	v168 = v167
	goto L38
L40:
	;
	v164 = *(*int32)(unsafe.Add(mBase, uint32(v137+int32(-9))))
	v168 = v164
	goto L38
L41:
	;
	v161 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v137+int32(-5)))))
	v168 = v161
	goto L38
L42:
	;
	v158 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v137+int32(-3)))))
	v168 = v158
	goto L38
L43:
	;
	v168 = int32(base.Ui32(v151) >> (uint(int32(3)) % 32))
	goto L38
L44:
	;
	F_addReplyBulkCString(m, l0, int32(_a1615))
	mBase = m.M
	v173 = m.ExcPending
	if v173 != 0 {
		goto L1
	} else {
		goto L45
	}
L45:
	;
	v174 = int64(*(*int32)(unsafe.Add(mBase, uint32(v138)+8)))
	F_addReplyLongLong(m, l0, v174)
	mBase = m.M
	v176 = m.ExcPending
	if v176 != 0 {
		goto L1
	} else {
		goto L46
	}
L46:
	;
	F_addReplyBulkCString(m, l0, int32(_a1616))
	mBase = m.M
	v179 = m.ExcPending
	if v179 != 0 {
		goto L1
	} else {
		goto L47
	}
L47:
	;
	v182 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v140+int32(-1)))))
	switch v182 & int32(7) {
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
		v199 = v147
		goto L48
	}
L48:
	;
	F_addReplyBulkCBuffer(m, l0, v140, v199)
	mBase = m.M
	v201 = m.ExcPending
	if v201 != 0 {
		goto L1
	} else {
		goto L54
	}
L49:
	;
	v198 = *(*int32)(unsafe.Add(mBase, uint32(v140+int32(-17))))
	v199 = v198
	goto L48
L50:
	;
	v195 = *(*int32)(unsafe.Add(mBase, uint32(v140+int32(-9))))
	v199 = v195
	goto L48
L51:
	;
	v192 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v140+int32(-5)))))
	v199 = v192
	goto L48
L52:
	;
	v189 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v140+int32(-3)))))
	v199 = v189
	goto L48
L53:
	;
	v199 = int32(base.Ui32(v182) >> (uint(int32(3)) % 32))
	goto L48
L54:
	;
	F_addReplyBulkCString(m, l0, int32(_a1617))
	mBase = m.M
	v204 = m.ExcPending
	if v204 != 0 {
		goto L1
	} else {
		goto L55
	}
L55:
	;
	v205 = *(*int32)(unsafe.Add(mBase, uint32(v138)+64))
	v206 = *(*int32)(unsafe.Add(mBase, uint32(v205)+4))
	F_addReplyArrayLen(m, l0, v206)
	mBase = m.M
	v208 = m.ExcPending
	if v208 != 0 {
		goto L1
	} else {
		goto L56
	}
L56:
	;
	v210 = *(*int32)(unsafe.Add(mBase, uint32(v138)+64))
	v211 = *(*int32)(unsafe.Add(mBase, uint32(v210)+4))
	if v211 < int32(1) {
		goto L57
	} else {
		goto L58
	}
L57:
	;
	v249 = v11 + int32(20)
	v250 = *(*int32)(unsafe.Add(mBase, uint32(v11)+16))
	if v250 != 0 {
		goto L65
	} else {
		goto L66
	}
L58:
	;
	v216 = int32(0)
	v217 = v210
	goto L59
L59:
	;
	v222 = *(*int32)(unsafe.Add(mBase, uint32(v217)+8))
	v226 = *(*int32)(unsafe.Add(mBase, uint32(v222+v216<<(uint(int32(2))%32))))
	F_addReplyBulk(m, l0, v226)
	mBase = m.M
	v228 = m.ExcPending
	if v228 != 0 {
		goto L1
	} else {
		goto L61
	}
L60:
	;
	goto L57
L61:
	;
	v230 = v216 + int32(1)
	v231 = *(*int32)(unsafe.Add(mBase, uint32(v138)+64))
	v232 = *(*int32)(unsafe.Add(mBase, uint32(v231)+4))
	if v230 < v232 {
		v216 = v230
		v217 = v231
		goto L59
	} else {
		goto L62
	}
L62:
	;
	goto L60
L63:
	;
	if v345 != 0 {
		v132 = v345
		goto L32
	} else {
		goto L89
	}
L64:
	;
	v256 = v249
	v257 = v253
	goto L67
L65:
	;
	v253 = int32(1)
	goto L64
L66:
	;
	v253 = int32(0)
	goto L64
L67:
	;
	switch v257 {
	case 0:
		goto L72
	default:
		goto L71
	}
L69:
	;
	v257 = int32(0)
	goto L67
L70:
	;
	goto L63
L71:
	;
	v337 = *(*int32)(unsafe.Add(mBase, uint32(v256)))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = v337
	if v337 == int32(0) {
		goto L69
	} else {
		goto L88
	}
L72:
	;
	v261 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
	if v261 != int32(-1) {
		v300 = v261
		goto L73
	} else {
		goto L74
	}
L73:
	;
	v301 = int32(1)
	v302 = v300 + v301
	*(*int32)(unsafe.Add(mBase, uint32(v11)+4)) = v302
	v304 = int32(0)
	v307 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
	v308 = *(*int32)(unsafe.Add(mBase, uint32(v11)+8))
	v312 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v307+v308+int32(26)))))
	if v312 == int32(255) {
		goto L82
	} else {
		goto L83
	}
L74:
	;
	v265 = *(*int32)(unsafe.Add(mBase, uint32(v11)+8))
	if v265 != 0 {
		v300 = int32(-1)
		goto L73
	} else {
		goto L75
	}
L75:
	;
	v266 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
	v267 = *(*int32)(unsafe.Add(mBase, uint32(v11)+12))
	if v267 == int32(0) {
		goto L77
	} else {
		goto L78
	}
L76:
	;
	v294 = *(*int32)(unsafe.Add(mBase, uint32(v293)+20))
	if v294 != int32(-1) {
		goto L79
	} else {
		goto L80
	}
L77:
	;
	v274 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v266)+16)))
	v275 = int64(*(*int8)(unsafe.Add(mBase, uint32(v266)+27)))
	v276 = int64(*(*int32)(unsafe.Add(mBase, uint32(v266)+8)))
	v277 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v266)+12)))
	v278 = int64(*(*int8)(unsafe.Add(mBase, uint32(v266)+26)))
	v279 = int64(*(*int32)(unsafe.Add(mBase, uint32(v266)+4)))
	v280 = F_wangHash64(m, v279)
	mBase = m.M
	v282 = F_wangHash64(m, v278+v280)
	mBase = m.M
	v284 = F_wangHash64(m, v277+v282)
	mBase = m.M
	v286 = F_wangHash64(m, v276+v284)
	mBase = m.M
	v288 = F_wangHash64(m, v275+v286)
	mBase = m.M
	v290 = F_wangHash64(m, v274+v288)
	mBase = m.M
	*(*int64)(unsafe.Add(mBase, uint32(v11)+24)) = v290
	v292 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
	v293 = v292
	goto L76
L78:
	;
	v270 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v266)+24)))
	v272 = v270 + int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v266)+24)) = uint16(v272)
	v293 = v266
	goto L76
L79:
	;
	v300 = v294 + int32(-1)
	goto L73
L80:
	;
	v297 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
	v300 = v297
	goto L73
L81:
	;
	v327 = int32(2)
	v332 = *(*int32)(unsafe.Add(mBase, uint32(v307+v325<<(uint(v327)%32)+int32(4))))
	v256 = v332 + v326<<(uint(v327)%32)
	v257 = int32(1)
	goto L67
L82:
	;
	v316 = v304
	goto L84
L83:
	;
	v316 = v301 << (uint(v312) % 32)
	goto L84
L84:
	;
	if v302 < v316 {
		v325 = v308
		v326 = v302
		goto L81
	} else {
		goto L85
	}
L85:
	;
	if v308 != 0 {
		v345 = v304
		goto L70
	} else {
		goto L86
	}
L86:
	;
	v318 = *(*int32)(unsafe.Add(mBase, uint32(v307)+20))
	if v318 == int32(-1) {
		v345 = v304
		goto L70
	} else {
		goto L87
	}
L87:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v11)+4)) = int64(4294967296)
	v325 = int32(1)
	v326 = int32(0)
	goto L81
L88:
	;
	v341 = *(*int32)(unsafe.Add(mBase, uint32(v337)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v249))) = v341
	v345 = v337
	goto L70
L89:
	;
	goto L33
L90:
	;
	return
}
func F_addReplyMapLen(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v24 int32
	_ = v24
	if l1 <= int32(-1) {
		F__serverAssert(m, int32(_a1652), int32(_a1630), int32(1420))
		mBase = m.M
		v24 = m.ExcPending
		if v24 != 0 {
			return
		} else {
			F_abort(m)
			mBase = m.M
			base.Wasm_trap_unreachable()
			for {
			}
		}
	} else {
		v6 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+224)))
		v7 = F_prepareClientToWrite(m, l0)
		mBase = m.M
		v8 = m.ExcPending
		if v8 != 0 {
			return
		} else {
			if v7 != 0 {
				return
			} else {
				v10 = base.B2i32(v6 == int32(2))
				if v6 == int32(2) {
					v17 = int32(42)
				} else {
					v17 = int32(37)
				}
				F__addReplyLongLongWithPrefix(m, l0, base.I64_extend_i32_u(l1<<(uint(v10)%32)), v17)
				mBase = m.M
				v19 = m.ExcPending
				if v19 != 0 {
					return
				} else {
					return
				}
			}
		}
	}
}
func F_addReplyNullArray(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v12 int32
	_ = v12
	var v16 int32
	_ = v16
	v4 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+224)))
	v5 = F_prepareClientToWrite(m, l0)
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return
	} else {
		if v4 != int32(2) {
			if v5 != 0 {
				return
			} else {
				F__addReplyToBufferOrList(m, l0, int32(_a1655), int32(3))
				mBase = m.M
				v16 = m.ExcPending
				if v16 != 0 {
					return
				} else {
					return
				}
			}
		} else {
			if v5 != 0 {
				return
			} else {
				F__addReplyToBufferOrList(m, l0, int32(_a1656), int32(5))
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
}
func F_addReplyOrErrorObject(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	switch int32(base.Ui32(v6)>>(uint(int32(4))%32)) & int32(15) {
	case 0, 8:
		v17 = F_objectGetVal(m, l1)
		mBase = m.M
		v20 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17+int32(-1)))))
		switch v20 & int32(7) {
		case 0:
			v37 = int32(base.Ui32(v20) >> (uint(int32(3)) % 32))
			if base.Ui32(v37) < base.Ui32(int32(2)) {
				F_addReply(m, l0, l1)
				mBase = m.M
				v75 = m.ExcPending
				if v75 != 0 {
					return
				} else {
					return
				}
			} else {
				v40 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17))))
				if v40 != int32(45) {
					F_addReply(m, l0, l1)
					mBase = m.M
					v75 = m.ExcPending
					if v75 != 0 {
						return
					} else {
						return
					}
				} else {
					F_addReply(m, l0, l1)
					mBase = m.M
					v44 = m.ExcPending
					if v44 != 0 {
						return
					} else {
						v45 = F_objectGetVal(m, l1)
						mBase = m.M
						v47 = F_objectGetVal(m, l1)
						mBase = m.M
						v50 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v47+int32(-1)))))
						switch v50 & int32(7) {
						case 0:
							v67 = int32(base.Ui32(v50) >> (uint(int32(3)) % 32))
						case 1:
							v57 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v47+int32(-3)))))
							v67 = v57
						case 2:
							v60 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v47+int32(-5)))))
							v67 = v60
						case 3:
							v63 = *(*int32)(unsafe.Add(mBase, uint32(v47+int32(-9))))
							v67 = v63
						case 4:
							v66 = *(*int32)(unsafe.Add(mBase, uint32(v47+int32(-17))))
							v67 = v66
						default:
							v67 = int32(0)
						}
						F_afterErrorReply(m, l0, v45, v67+int32(-2), int32(0))
						mBase = m.M
						v72 = m.ExcPending
						if v72 != 0 {
							return
						} else {
							return
						}
					}
				}
			}
		case 1:
			v27 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17+int32(-3)))))
			v37 = v27
			if base.Ui32(v37) < base.Ui32(int32(2)) {
				F_addReply(m, l0, l1)
				mBase = m.M
				v75 = m.ExcPending
				if v75 != 0 {
					return
				} else {
					return
				}
			} else {
				v40 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17))))
				if v40 != int32(45) {
					F_addReply(m, l0, l1)
					mBase = m.M
					v75 = m.ExcPending
					if v75 != 0 {
						return
					} else {
						return
					}
				} else {
					F_addReply(m, l0, l1)
					mBase = m.M
					v44 = m.ExcPending
					if v44 != 0 {
						return
					} else {
						v45 = F_objectGetVal(m, l1)
						mBase = m.M
						v47 = F_objectGetVal(m, l1)
						mBase = m.M
						v50 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v47+int32(-1)))))
						switch v50 & int32(7) {
						case 0:
							v67 = int32(base.Ui32(v50) >> (uint(int32(3)) % 32))
						case 1:
							v57 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v47+int32(-3)))))
							v67 = v57
						case 2:
							v60 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v47+int32(-5)))))
							v67 = v60
						case 3:
							v63 = *(*int32)(unsafe.Add(mBase, uint32(v47+int32(-9))))
							v67 = v63
						case 4:
							v66 = *(*int32)(unsafe.Add(mBase, uint32(v47+int32(-17))))
							v67 = v66
						default:
							v67 = int32(0)
						}
						F_afterErrorReply(m, l0, v45, v67+int32(-2), int32(0))
						mBase = m.M
						v72 = m.ExcPending
						if v72 != 0 {
							return
						} else {
							return
						}
					}
				}
			}
		case 2:
			v30 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v17+int32(-5)))))
			v37 = v30
			if base.Ui32(v37) < base.Ui32(int32(2)) {
				F_addReply(m, l0, l1)
				mBase = m.M
				v75 = m.ExcPending
				if v75 != 0 {
					return
				} else {
					return
				}
			} else {
				v40 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17))))
				if v40 != int32(45) {
					F_addReply(m, l0, l1)
					mBase = m.M
					v75 = m.ExcPending
					if v75 != 0 {
						return
					} else {
						return
					}
				} else {
					F_addReply(m, l0, l1)
					mBase = m.M
					v44 = m.ExcPending
					if v44 != 0 {
						return
					} else {
						v45 = F_objectGetVal(m, l1)
						mBase = m.M
						v47 = F_objectGetVal(m, l1)
						mBase = m.M
						v50 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v47+int32(-1)))))
						switch v50 & int32(7) {
						case 0:
							v67 = int32(base.Ui32(v50) >> (uint(int32(3)) % 32))
						case 1:
							v57 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v47+int32(-3)))))
							v67 = v57
						case 2:
							v60 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v47+int32(-5)))))
							v67 = v60
						case 3:
							v63 = *(*int32)(unsafe.Add(mBase, uint32(v47+int32(-9))))
							v67 = v63
						case 4:
							v66 = *(*int32)(unsafe.Add(mBase, uint32(v47+int32(-17))))
							v67 = v66
						default:
							v67 = int32(0)
						}
						F_afterErrorReply(m, l0, v45, v67+int32(-2), int32(0))
						mBase = m.M
						v72 = m.ExcPending
						if v72 != 0 {
							return
						} else {
							return
						}
					}
				}
			}
		case 3:
			v33 = *(*int32)(unsafe.Add(mBase, uint32(v17+int32(-9))))
			v37 = v33
			if base.Ui32(v37) < base.Ui32(int32(2)) {
				F_addReply(m, l0, l1)
				mBase = m.M
				v75 = m.ExcPending
				if v75 != 0 {
					return
				} else {
					return
				}
			} else {
				v40 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17))))
				if v40 != int32(45) {
					F_addReply(m, l0, l1)
					mBase = m.M
					v75 = m.ExcPending
					if v75 != 0 {
						return
					} else {
						return
					}
				} else {
					F_addReply(m, l0, l1)
					mBase = m.M
					v44 = m.ExcPending
					if v44 != 0 {
						return
					} else {
						v45 = F_objectGetVal(m, l1)
						mBase = m.M
						v47 = F_objectGetVal(m, l1)
						mBase = m.M
						v50 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v47+int32(-1)))))
						switch v50 & int32(7) {
						case 0:
							v67 = int32(base.Ui32(v50) >> (uint(int32(3)) % 32))
						case 1:
							v57 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v47+int32(-3)))))
							v67 = v57
						case 2:
							v60 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v47+int32(-5)))))
							v67 = v60
						case 3:
							v63 = *(*int32)(unsafe.Add(mBase, uint32(v47+int32(-9))))
							v67 = v63
						case 4:
							v66 = *(*int32)(unsafe.Add(mBase, uint32(v47+int32(-17))))
							v67 = v66
						default:
							v67 = int32(0)
						}
						F_afterErrorReply(m, l0, v45, v67+int32(-2), int32(0))
						mBase = m.M
						v72 = m.ExcPending
						if v72 != 0 {
							return
						} else {
							return
						}
					}
				}
			}
		case 4:
			v36 = *(*int32)(unsafe.Add(mBase, uint32(v17+int32(-17))))
			v37 = v36
			if base.Ui32(v37) < base.Ui32(int32(2)) {
				F_addReply(m, l0, l1)
				mBase = m.M
				v75 = m.ExcPending
				if v75 != 0 {
					return
				} else {
					return
				}
			} else {
				v40 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17))))
				if v40 != int32(45) {
					F_addReply(m, l0, l1)
					mBase = m.M
					v75 = m.ExcPending
					if v75 != 0 {
						return
					} else {
						return
					}
				} else {
					F_addReply(m, l0, l1)
					mBase = m.M
					v44 = m.ExcPending
					if v44 != 0 {
						return
					} else {
						v45 = F_objectGetVal(m, l1)
						mBase = m.M
						v47 = F_objectGetVal(m, l1)
						mBase = m.M
						v50 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v47+int32(-1)))))
						switch v50 & int32(7) {
						case 0:
							v67 = int32(base.Ui32(v50) >> (uint(int32(3)) % 32))
						case 1:
							v57 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v47+int32(-3)))))
							v67 = v57
						case 2:
							v60 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v47+int32(-5)))))
							v67 = v60
						case 3:
							v63 = *(*int32)(unsafe.Add(mBase, uint32(v47+int32(-9))))
							v67 = v63
						case 4:
							v66 = *(*int32)(unsafe.Add(mBase, uint32(v47+int32(-17))))
							v67 = v66
						default:
							v67 = int32(0)
						}
						F_afterErrorReply(m, l0, v45, v67+int32(-2), int32(0))
						mBase = m.M
						v72 = m.ExcPending
						if v72 != 0 {
							return
						} else {
							return
						}
					}
				}
			}
		default:
			F_addReply(m, l0, l1)
			mBase = m.M
			v75 = m.ExcPending
			if v75 != 0 {
				return
			} else {
				return
			}
		}
	default:
		F__serverAssert(m, int32(_a1650), int32(_a1630), int32(989))
		mBase = m.M
		v15 = m.ExcPending
		if v15 != 0 {
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
func F_addReplyProto(m *base.Module, l0 int32, l1 int32, l2 int32) {
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	v4 = F_prepareClientToWrite(m, l0)
	v5 = m.ExcPending
	if v5 != 0 {
		return
	} else {
		if v4 != 0 {
			return
		} else {
			F__addReplyToBufferOrList(m, l0, l1, l2)
			v7 = m.ExcPending
			if v7 != 0 {
				return
			} else {
				return
			}
		}
	}
}
func F_addReplyPushLen(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v25 int64
	_ = v25
	var v34 int32
	_ = v34
	var v38 int32
	_ = v38
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int64
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v66 int32
	_ = v66
	var v70 int32
	_ = v70
	var v75 int32
	_ = v75
	var v84 int32
	_ = v84
	var v91 int32
	_ = v91
	var v97 int32
	_ = v97
	v4 = m.G0
	v6 = v4 - int32(128)
	m.G0 = v6
	v8 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+224)))
	if base.Ui32(v8) <= base.Ui32(int32(2)) {
		F__serverAssert(m, int32(_a1653), int32(_a1630), int32(1459))
		mBase = m.M
		v84 = m.ExcPending
		if v84 != 0 {
			return
		} else {
			F_abort(m)
			mBase = m.M
			base.Wasm_trap_unreachable()
			for {
			}
		}
	} else {
		v11 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+206)))
		if v11&int32(2) == int32(0) {
			F__serverAssertWithInfo(m, l0, int32(0), int32(_a1654), int32(_a1630), int32(1460))
			mBase = m.M
			v91 = m.ExcPending
			if v91 != 0 {
				return
			} else {
				F_abort(m)
				mBase = m.M
				base.Wasm_trap_unreachable()
				for {
				}
			}
		} else {
			if l1 <= int32(-1) {
				F__serverAssert(m, int32(_a1652), int32(_a1630), int32(1420))
				mBase = m.M
				v97 = m.ExcPending
				if v97 != 0 {
					return
				} else {
					F_abort(m)
					mBase = m.M
					base.Wasm_trap_unreachable()
					for {
					}
				}
			} else {
				v18 = F_prepareClientToWrite(m, l0)
				mBase = m.M
				v19 = m.ExcPending
				if v19 != 0 {
					return
				} else {
					if v18 != 0 {
						m.G0 = v6 + int32(128)
						return
					} else {
						v20 = int32(62)
						*(*uint8)(unsafe.Add(mBase, uint32(v6))) = uint8(v20)
						v23 = v6 | int32(1)
						v25 = base.I64_extend_i32_u(l1)
						if v25 <= int64(-1) {
							v34 = int32(45)
							*(*uint8)(unsafe.Add(mBase, uint32(v23))) = uint8(v34)
							v38 = int32(1)
							v43 = v23 + v38
							v44 = int32(126)
							v45 = int64(0) - v25
							v46 = v38
						} else {
							v43 = v23
							v44 = int32(127)
							v45 = v25
							v46 = int32(0)
						}
						v47 = F_ull2string(m, v43, v44, v45)
						mBase = m.M
						if v47 == int32(0) {
							v66 = int32(0)
						} else {
							v66 = v47 + v46
						}
						v70 = int32(2573)
						*(*uint16)(unsafe.Add(mBase, uint32(v66+v6+int32(1)))) = uint16(v70)
						F__addReplyToBufferOrList(m, l0, v6, v66+int32(3))
						mBase = m.M
						v75 = m.ExcPending
						if v75 != 0 {
							return
						} else {
							m.G0 = v6 + int32(128)
							return
						}
					}
				}
			}
		}
	}
}
func F_addReplyReplicationBacklog(m *base.Module, l0 int32, l1 int64) int64 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v23 int32
	_ = v23
	var v24 int64
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v29 int64
	_ = v29
	var v34 int32
	_ = v34
	var v38 int32
	_ = v38
	var v42 int64
	_ = v42
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v56 int64
	_ = v56
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v70 int64
	_ = v70
	var v72 int32
	_ = v72
	var v73 int64
	_ = v73
	var v75 int64
	_ = v75
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v85 int64
	_ = v85
	var v86 int64
	_ = v86
	var v88 int32
	_ = v88
	var v95 int32
	_ = v95
	var v97 int32
	_ = v97
	var v98 int64
	_ = v98
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v101 int64
	_ = v101
	var v104 int64
	_ = v104
	var v106 int64
	_ = v106
	var v108 int64
	_ = v108
	var v111 int64
	_ = v111
	var v113 int64
	_ = v113
	var v115 int64
	_ = v115
	var v117 int64
	_ = v117
	var v141 int32
	_ = v141
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v150 int64
	_ = v150
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v172 int32
	_ = v172
	var v180 int32
	_ = v180
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
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
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
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
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v221 int64
	_ = v221
	var v222 int64
	_ = v222
	var v225 int32
	_ = v225
	var v236 int32
	_ = v236
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v245 int32
	_ = v245
	var v247 int64
	_ = v247
	var v248 int64
	_ = v248
	var v251 int32
	_ = v251
	var v252 int64
	_ = v252
	var v257 int64
	_ = v257
	v7 = m.G0
	v9 = v7 - int32(384)
	m.G0 = v9
	v12 = *(*int32)(unsafe.Add(mBase, _consts[28]))
	if int32(0) < v12 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v24 = int64(0)
	v25 = int32(_a20)
	v26 = *(*int32)(unsafe.Add(mBase, _consts[28]))
	v28 = *(*int32)(unsafe.Add(mBase, _consts[293]))
	v29 = *(*int64)(unsafe.Add(mBase, uint32(v28)+16))
	if v29 != v24 {
		goto L6
	} else {
		goto L7
	}
L2:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v9)+64)) = l1
	F__serverLog(m, int32(0), int32(_a1927), v9+int32(64))
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	return int64(0)
L4:
	;
	goto L1
L5:
	;
	m.G0 = v9 + int32(384)
	return v257
L6:
	;
	if int32(0) < v26 {
		v72 = v28
		goto L12
	} else {
		goto L13
	}
L7:
	;
	if int32(0) < v26 {
		v257 = v24
		goto L5
	} else {
		goto L8
	}
L8:
	;
	v34 = int32(0)
	F__serverLog(m, v34, int32(_a1928), v34)
	mBase = m.M
	v38 = m.ExcPending
	if v38 != 0 {
		goto L3
	} else {
		goto L9
	}
L9:
	;
	v257 = v24
	goto L5
L10:
	;
	v100 = *(*int32)(unsafe.Add(mBase, uint32(v99)+8))
	v101 = *(*int64)(unsafe.Add(mBase, uint32(v100)+8))
	goto L23
L11:
	;
	v75 = *(*int64)(unsafe.Add(mBase, uint32(v65)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v9)+16)) = v75
	F__serverLog(m, int32(0), int32(_a1929), v9+int32(16))
	mBase = m.M
	v82 = m.ExcPending
	if v82 != 0 {
		goto L3
	} else {
		goto L18
	}
L12:
	;
	v73 = *(*int64)(unsafe.Add(mBase, uint32(v72)+24))
	v98 = l1 - v73
	v99 = v72
	goto L10
L13:
	;
	v42 = *(*int64)(unsafe.Add(mBase, _consts[322]))
	*(*int64)(unsafe.Add(mBase, uint32(v9)+48)) = v42
	F__serverLog(m, int32(0), int32(_a1930), v9+int32(48))
	mBase = m.M
	v49 = m.ExcPending
	if v49 != 0 {
		goto L3
	} else {
		goto L14
	}
L14:
	;
	v50 = int32(_a20)
	v51 = *(*int32)(unsafe.Add(mBase, _consts[293]))
	v53 = *(*int32)(unsafe.Add(mBase, _consts[28]))
	if int32(0) < v53 {
		v72 = v51
		goto L12
	} else {
		goto L15
	}
L15:
	;
	v56 = *(*int64)(unsafe.Add(mBase, uint32(v51)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v9)+32)) = v56
	F__serverLog(m, int32(0), int32(_a1931), v9+int32(32))
	mBase = m.M
	v63 = m.ExcPending
	if v63 != 0 {
		goto L3
	} else {
		goto L16
	}
L16:
	;
	v64 = int32(_a20)
	v65 = *(*int32)(unsafe.Add(mBase, _consts[293]))
	v67 = *(*int32)(unsafe.Add(mBase, _consts[28]))
	if v67 < int32(1) {
		goto L11
	} else {
		goto L17
	}
L17:
	;
	v70 = *(*int64)(unsafe.Add(mBase, uint32(v65)+24))
	v98 = l1 - v70
	v99 = v65
	goto L10
L18:
	;
	v83 = int32(_a20)
	v84 = *(*int32)(unsafe.Add(mBase, _consts[293]))
	v85 = *(*int64)(unsafe.Add(mBase, uint32(v84)+24))
	v86 = l1 - v85
	v88 = *(*int32)(unsafe.Add(mBase, _consts[28]))
	if int32(0) < v88 {
		v98 = v86
		v99 = v84
		goto L10
	} else {
		goto L19
	}
L19:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v9))) = v86
	F__serverLog(m, int32(0), int32(_a1932), v9)
	mBase = m.M
	v95 = m.ExcPending
	if v95 != 0 {
		goto L3
	} else {
		goto L20
	}
L20:
	;
	v97 = *(*int32)(unsafe.Add(mBase, _consts[293]))
	v98 = v86
	v99 = v97
	goto L10
L21:
	;
	if v211 == int32(0) {
		goto L39
	} else {
		goto L40
	}
L22:
	;
	v209 = *(*int32)(unsafe.Add(mBase, _consts[293]))
	v210 = *(*int32)(unsafe.Add(mBase, uint32(v209)))
	v211 = v210
	goto L21
L23:
	;
	if v101 == int64(0) {
		goto L22
	} else {
		goto L24
	}
L24:
	;
	v104 = int64(56)
	v106 = int64(65280)
	v108 = int64(40)
	v111 = int64(16711680)
	v113 = int64(24)
	v115 = int64(4278190080)
	v117 = int64(8)
	*(*int64)(unsafe.Add(mBase, uint32(v9)+376)) = l1<<(uint(v104)%64) | l1&v106<<(uint(v108)%64) | (l1&v111<<(uint(v113)%64) | l1&v115<<(uint(v117)%64)) | (int64(base.Ui64(l1)>>(uint(v117)%64))&v115 | int64(base.Ui64(l1)>>(uint(v113)%64))&v111 | (int64(base.Ui64(l1)>>(uint(v108)%64))&v106 | int64(base.Ui64(l1)>>(uint(v104)%64))))
	v141 = v9 + int32(72)
	v143 = *(*int32)(unsafe.Add(mBase, _consts[293]))
	v144 = *(*int32)(unsafe.Add(mBase, uint32(v143)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v141)+4)) = v144
	*(*int32)(unsafe.Add(mBase, uint32(v141))) = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v141)+20)) = int32(128)
	v150 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v141)+12)) = v150
	*(*int64)(unsafe.Add(mBase, uint32(v141)+296)) = v150
	*(*int64)(unsafe.Add(mBase, uint32(v141)+160)) = int64(137438953472)
	*(*int32)(unsafe.Add(mBase, uint32(v141)+8)) = v9 + int32(96)
	*(*int32)(unsafe.Add(mBase, uint32(v141)+156)) = v9 + int32(240)
	goto L25
L25:
	;
	v168 = F_raxSeek(m, v9+int32(72), int32(_a104), v9+int32(376), int32(8))
	mBase = m.M
	v169 = m.ExcPending
	if v169 != 0 {
		goto L3
	} else {
		goto L26
	}
L26:
	;
	v172 = *(*int32)(unsafe.Add(mBase, uint32(v9+int32(72))))
	goto L29
L27:
	;
	F_raxStop(m, v9+int32(72))
	mBase = m.M
	v207 = m.ExcPending
	if v207 != 0 {
		goto L3
	} else {
		goto L37
	}
L28:
	;
	v191 = F_raxPrev(m, v9+int32(72))
	mBase = m.M
	v192 = m.ExcPending
	if v192 != 0 {
		goto L3
	} else {
		goto L33
	}
L29:
	;
	if v172&int32(2) == int32(0) {
		goto L28
	} else {
		goto L30
	}
L30:
	;
	v180 = int32(0)
	v182 = F_raxSeek(m, v9+int32(72), int32(_a1933), v180, v180)
	mBase = m.M
	v183 = m.ExcPending
	if v183 != 0 {
		goto L3
	} else {
		goto L31
	}
L31:
	;
	v186 = F_raxPrev(m, v9+int32(72))
	mBase = m.M
	v187 = m.ExcPending
	if v187 != 0 {
		goto L3
	} else {
		goto L32
	}
L32:
	;
	v188 = *(*int32)(unsafe.Add(mBase, uint32(v9)+84))
	v203 = v188
	goto L27
L33:
	;
	v195 = F_raxPrev(m, v9+int32(72))
	mBase = m.M
	v196 = m.ExcPending
	if v196 != 0 {
		goto L3
	} else {
		goto L35
	}
L34:
	;
	v201 = *(*int32)(unsafe.Add(mBase, _consts[293]))
	v202 = *(*int32)(unsafe.Add(mBase, uint32(v201)))
	v203 = v202
	goto L27
L35:
	;
	if v195 == int32(0) {
		goto L34
	} else {
		goto L36
	}
L36:
	;
	v199 = *(*int32)(unsafe.Add(mBase, uint32(v9)+84))
	v203 = v199
	goto L27
L37:
	;
	v211 = v203
	goto L21
L38:
	;
	v238 = F_prepareClientToWrite(m, l0)
	mBase = m.M
	v239 = m.ExcPending
	if v239 != 0 {
		goto L3
	} else {
		goto L46
	}
L39:
	;
	F__serverAssert(m, int32(_a1934), int32(_a1913), int32(789))
	mBase = m.M
	v236 = m.ExcPending
	if v236 != 0 {
		goto L3
	} else {
		goto L45
	}
L40:
	;
	v219 = v211
	goto L41
L41:
	;
	v220 = *(*int32)(unsafe.Add(mBase, uint32(v219)+8))
	v221 = *(*int64)(unsafe.Add(mBase, uint32(v220)+16))
	v222 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v220)+28)))
	if l1 <= v221+v222 {
		goto L38
	} else {
		goto L43
	}
L42:
	;
	goto L39
L43:
	;
	v225 = *(*int32)(unsafe.Add(mBase, uint32(v219)+4))
	if v225 != 0 {
		v219 = v225
		goto L41
	} else {
		goto L44
	}
L44:
	;
	goto L42
L45:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L46:
	;
	v240 = *(*int32)(unsafe.Add(mBase, uint32(v219)+8))
	v241 = *(*int32)(unsafe.Add(mBase, uint32(v240)))
	*(*int32)(unsafe.Add(mBase, uint32(v240))) = v241 + int32(1)
	v245 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
	*(*int32)(unsafe.Add(mBase, uint32(v245)+184)) = v219
	v247 = *(*int64)(unsafe.Add(mBase, uint32(v240)+16))
	v248 = l1 - v247
	*(*uint32)(unsafe.Add(mBase, uint32(v245)+188)) = uint32(v248)
	v251 = *(*int32)(unsafe.Add(mBase, _consts[293]))
	v252 = *(*int64)(unsafe.Add(mBase, uint32(v251)+16))
	v257 = v252 - v98
	goto L5
}
func F_addReplySentinelDebugInfo(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v9 int64
	_ = v9
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v16 int64
	_ = v16
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v23 int64
	_ = v23
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v30 int64
	_ = v30
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v37 int64
	_ = v37
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v44 int64
	_ = v44
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v51 int64
	_ = v51
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v58 int64
	_ = v58
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v65 int64
	_ = v65
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v72 int64
	_ = v72
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v79 int64
	_ = v79
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v86 int64
	_ = v86
	var v88 int32
	_ = v88
	var v91 int32
	_ = v91
	var v93 int64
	_ = v93
	var v95 int32
	_ = v95
	var v98 int32
	_ = v98
	v3 = F_addReplyDeferredLen(m, l0)
	mBase = m.M
	v4 = m.ExcPending
	if v4 != 0 {
		return
	} else {
		F_addReplyBulkCString(m, l0, int32(_a2123))
		mBase = m.M
		v7 = m.ExcPending
		if v7 != 0 {
			return
		} else {
			v9 = *(*int64)(unsafe.Add(mBase, _consts[692]))
			F_addReplyBulkLongLong(m, l0, v9)
			mBase = m.M
			v11 = m.ExcPending
			if v11 != 0 {
				return
			} else {
				F_addReplyBulkCString(m, l0, int32(_a2124))
				mBase = m.M
				v14 = m.ExcPending
				if v14 != 0 {
					return
				} else {
					v16 = *(*int64)(unsafe.Add(mBase, _consts[693]))
					F_addReplyBulkLongLong(m, l0, v16)
					mBase = m.M
					v18 = m.ExcPending
					if v18 != 0 {
						return
					} else {
						F_addReplyBulkCString(m, l0, int32(_a2125))
						mBase = m.M
						v21 = m.ExcPending
						if v21 != 0 {
							return
						} else {
							v23 = *(*int64)(unsafe.Add(mBase, _consts[694]))
							F_addReplyBulkLongLong(m, l0, v23)
							mBase = m.M
							v25 = m.ExcPending
							if v25 != 0 {
								return
							} else {
								F_addReplyBulkCString(m, l0, int32(_a2126))
								mBase = m.M
								v28 = m.ExcPending
								if v28 != 0 {
									return
								} else {
									v30 = *(*int64)(unsafe.Add(mBase, _consts[691]))
									F_addReplyBulkLongLong(m, l0, v30)
									mBase = m.M
									v32 = m.ExcPending
									if v32 != 0 {
										return
									} else {
										F_addReplyBulkCString(m, l0, int32(_a2127))
										mBase = m.M
										v35 = m.ExcPending
										if v35 != 0 {
											return
										} else {
											v37 = *(*int64)(unsafe.Add(mBase, _consts[695]))
											F_addReplyBulkLongLong(m, l0, v37)
											mBase = m.M
											v39 = m.ExcPending
											if v39 != 0 {
												return
											} else {
												F_addReplyBulkCString(m, l0, int32(_a2128))
												mBase = m.M
												v42 = m.ExcPending
												if v42 != 0 {
													return
												} else {
													v44 = *(*int64)(unsafe.Add(mBase, _consts[696]))
													F_addReplyBulkLongLong(m, l0, v44)
													mBase = m.M
													v46 = m.ExcPending
													if v46 != 0 {
														return
													} else {
														F_addReplyBulkCString(m, l0, int32(_a2129))
														mBase = m.M
														v49 = m.ExcPending
														if v49 != 0 {
															return
														} else {
															v51 = *(*int64)(unsafe.Add(mBase, _consts[697]))
															F_addReplyBulkLongLong(m, l0, v51)
															mBase = m.M
															v53 = m.ExcPending
															if v53 != 0 {
																return
															} else {
																F_addReplyBulkCString(m, l0, int32(_a2130))
																mBase = m.M
																v56 = m.ExcPending
																if v56 != 0 {
																	return
																} else {
																	v58 = *(*int64)(unsafe.Add(mBase, _consts[698]))
																	F_addReplyBulkLongLong(m, l0, v58)
																	mBase = m.M
																	v60 = m.ExcPending
																	if v60 != 0 {
																		return
																	} else {
																		F_addReplyBulkCString(m, l0, int32(_a2131))
																		mBase = m.M
																		v63 = m.ExcPending
																		if v63 != 0 {
																			return
																		} else {
																			v65 = *(*int64)(unsafe.Add(mBase, _consts[699]))
																			F_addReplyBulkLongLong(m, l0, v65)
																			mBase = m.M
																			v67 = m.ExcPending
																			if v67 != 0 {
																				return
																			} else {
																				F_addReplyBulkCString(m, l0, int32(_a2132))
																				mBase = m.M
																				v70 = m.ExcPending
																				if v70 != 0 {
																					return
																				} else {
																					v72 = *(*int64)(unsafe.Add(mBase, _consts[700]))
																					F_addReplyBulkLongLong(m, l0, v72)
																					mBase = m.M
																					v74 = m.ExcPending
																					if v74 != 0 {
																						return
																					} else {
																						F_addReplyBulkCString(m, l0, int32(_a2133))
																						mBase = m.M
																						v77 = m.ExcPending
																						if v77 != 0 {
																							return
																						} else {
																							v79 = *(*int64)(unsafe.Add(mBase, _consts[701]))
																							F_addReplyBulkLongLong(m, l0, v79)
																							mBase = m.M
																							v81 = m.ExcPending
																							if v81 != 0 {
																								return
																							} else {
																								F_addReplyBulkCString(m, l0, int32(_a2134))
																								mBase = m.M
																								v84 = m.ExcPending
																								if v84 != 0 {
																									return
																								} else {
																									v86 = *(*int64)(unsafe.Add(mBase, _consts[702]))
																									F_addReplyBulkLongLong(m, l0, v86)
																									mBase = m.M
																									v88 = m.ExcPending
																									if v88 != 0 {
																										return
																									} else {
																										F_addReplyBulkCString(m, l0, int32(_a2135))
																										mBase = m.M
																										v91 = m.ExcPending
																										if v91 != 0 {
																											return
																										} else {
																											v93 = *(*int64)(unsafe.Add(mBase, _consts[680]))
																											F_addReplyBulkLongLong(m, l0, v93)
																											mBase = m.M
																											v95 = m.ExcPending
																											if v95 != 0 {
																												return
																											} else {
																												F_setDeferredMapLen(m, l0, v3, int32(13))
																												mBase = m.M
																												v98 = m.ExcPending
																												if v98 != 0 {
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
				}
			}
		}
	}
}
func F_addReplySlotStat(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v26 int32
	_ = v26
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v63 int32
	_ = v63
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v85 int64
	_ = v85
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v96 int64
	_ = v96
	var v98 int32
	_ = v98
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v107 int64
	_ = v107
	var v109 int32
	_ = v109
	F_addReplyArrayLen(m, l0, int32(2))
	mBase = m.M
	v5 = m.ExcPending
	if v5 != 0 {
		return
	} else {
		F_addReplyLongLong(m, l0, base.I64_extend_i32_s(l1))
		mBase = m.M
		v8 = m.ExcPending
		if v8 != 0 {
			return
		} else {
			v12 = *(*int32)(unsafe.Add(mBase, _consts[229]))
			if v12 != 0 {
				v13 = int32(4)
			} else {
				v13 = int32(1)
			}
			F_addReplyMapLen(m, l0, v13)
			mBase = m.M
			v15 = m.ExcPending
			if v15 != 0 {
				return
			} else {
				F_addReplyBulkCString(m, l0, int32(_a453))
				mBase = m.M
				v18 = m.ExcPending
				if v18 != 0 {
					return
				} else {
					v26 = *(*int32)(unsafe.Add(mBase, _consts[172]))
					if int32(1) <= v26 {
						v31 = *(*int32)(unsafe.Add(mBase, _consts[173]))
						v32 = int32(0)
						v35 = v26
						v36 = v32
						v37 = v31
						v38 = v32
						for {
							v41 = int32(0)
							v45 = *(*int32)(unsafe.Add(mBase, uint32(v37+v38<<(uint(int32(2))%32))))
							if v45 == v41 {
								v54 = v35
								v55 = v37
								v56 = v41
							} else {
								v48 = *(*int32)(unsafe.Add(mBase, uint32(v45)))
								v49 = F_kvstoreHashtableSize(m, v48, l1)
								mBase = m.M
								v50 = int32(_a20)
								v51 = *(*int32)(unsafe.Add(mBase, _consts[172]))
								v53 = *(*int32)(unsafe.Add(mBase, _consts[173]))
								v54 = v51
								v55 = v53
								v56 = v49
							}
							v57 = v56 + v36
							v59 = v38 + int32(1)
							if v59 < v54 {
								v35 = v54
								v36 = v57
								v37 = v55
								v38 = v59
								continue
							} else {
								break
							}
							break
						}
						v63 = v57
					} else {
						v63 = int32(0)
					}
					F_addReplyLongLong(m, l0, base.I64_extend_i32_u(v63))
					mBase = m.M
					v70 = m.ExcPending
					if v70 != 0 {
						return
					} else {
						v72 = *(*int32)(unsafe.Add(mBase, _consts[229]))
						if v72 == int32(0) {
							return
						} else {
							F_addReplyBulkCString(m, l0, int32(_a454))
							mBase = m.M
							v77 = m.ExcPending
							if v77 != 0 {
								return
							} else {
								v79 = *(*int32)(unsafe.Add(mBase, _consts[111]))
								v81 = l1 * int32(24)
								v85 = *(*int64)(unsafe.Add(mBase, uint32(v79+v81)+uint32(_consts[165])))
								F_addReplyLongLong(m, l0, v85)
								mBase = m.M
								v87 = m.ExcPending
								if v87 != 0 {
									return
								} else {
									F_addReplyBulkCString(m, l0, int32(_a455))
									mBase = m.M
									v90 = m.ExcPending
									if v90 != 0 {
										return
									} else {
										v92 = *(*int32)(unsafe.Add(mBase, _consts[111]))
										v96 = *(*int64)(unsafe.Add(mBase, uint32(v92+v81)+uint32(_consts[164])))
										F_addReplyLongLong(m, l0, v96)
										mBase = m.M
										v98 = m.ExcPending
										if v98 != 0 {
											return
										} else {
											F_addReplyBulkCString(m, l0, int32(_a456))
											mBase = m.M
											v101 = m.ExcPending
											if v101 != 0 {
												return
											} else {
												v103 = *(*int32)(unsafe.Add(mBase, _consts[111]))
												v107 = *(*int64)(unsafe.Add(mBase, uint32(v103+v81)+uint32(_consts[163])))
												F_addReplyLongLong(m, l0, v107)
												mBase = m.M
												v109 = m.ExcPending
												if v109 != 0 {
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
					}
				}
			}
		}
	}
}
func F_addReplyStatusFormat(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
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
	var v40 int32
	_ = v40
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	*(*int32)(unsafe.Add(mBase, uint32(v8)+12)) = l2
	v12 = F_sdsempty(m)
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return
	} else {
		v14 = *(*int32)(unsafe.Add(mBase, uint32(v8)+12))
		v15 = F_sdscatvprintf(m, v12, l1, v14)
		mBase = m.M
		v16 = m.ExcPending
		if v16 != 0 {
			return
		} else {
			v19 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15+int32(-1)))))
			switch v19 & int32(7) {
			case 0:
				v36 = int32(base.Ui32(v19) >> (uint(int32(3)) % 32))
			case 1:
				v26 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15+int32(-3)))))
				v36 = v26
			case 2:
				v29 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v15+int32(-5)))))
				v36 = v29
			case 3:
				v32 = *(*int32)(unsafe.Add(mBase, uint32(v15+int32(-9))))
				v36 = v32
			case 4:
				v35 = *(*int32)(unsafe.Add(mBase, uint32(v15+int32(-17))))
				v36 = v35
			default:
				v36 = int32(0)
			}
			F_addReplyStatusLength(m, l0, v15, v36)
			mBase = m.M
			v38 = m.ExcPending
			if v38 != 0 {
				return
			} else {
				F_sdsfree(m, v15)
				mBase = m.M
				v40 = m.ExcPending
				if v40 != 0 {
					return
				} else {
					m.G0 = v8 + int32(16)
					return
				}
			}
		}
	}
}
func F_addReplySubcommandSyntaxError(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
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
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(v9)))
	v11 = F_objectGetVal(m, v10)
	mBase = m.M
	v12 = F_sdsnew(m, v11)
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return
	} else {
		v19 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12+int32(-1)))))
		switch v19 & int32(7) {
		case 0:
			v36 = int32(base.Ui32(v19) >> (uint(int32(3)) % 32))
			if v36 == int32(0) {
			} else {
				v41 = int32(0)
				for {
					v44 = v12 + v41
					v45 = int32(*(*int8)(unsafe.Add(mBase, uint32(v44))))
					v46 = F_toupper(m, v45)
					mBase = m.M
					*(*uint8)(unsafe.Add(mBase, uint32(v44))) = uint8(v46)
					v49 = v41 + int32(1)
					if v49 != v36 {
						v41 = v49
						continue
					} else {
						break
					}
					break
				}
			}
		case 1:
			v26 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12+int32(-3)))))
			v36 = v26
			if v36 == int32(0) {
			} else {
				v41 = int32(0)
				for {
					v44 = v12 + v41
					v45 = int32(*(*int8)(unsafe.Add(mBase, uint32(v44))))
					v46 = F_toupper(m, v45)
					mBase = m.M
					*(*uint8)(unsafe.Add(mBase, uint32(v44))) = uint8(v46)
					v49 = v41 + int32(1)
					if v49 != v36 {
						v41 = v49
						continue
					} else {
						break
					}
					break
				}
			}
		case 2:
			v29 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v12+int32(-5)))))
			v36 = v29
			if v36 == int32(0) {
			} else {
				v41 = int32(0)
				for {
					v44 = v12 + v41
					v45 = int32(*(*int8)(unsafe.Add(mBase, uint32(v44))))
					v46 = F_toupper(m, v45)
					mBase = m.M
					*(*uint8)(unsafe.Add(mBase, uint32(v44))) = uint8(v46)
					v49 = v41 + int32(1)
					if v49 != v36 {
						v41 = v49
						continue
					} else {
						break
					}
					break
				}
			}
		case 3:
			v32 = *(*int32)(unsafe.Add(mBase, uint32(v12+int32(-9))))
			v36 = v32
			if v36 == int32(0) {
			} else {
				v41 = int32(0)
				for {
					v44 = v12 + v41
					v45 = int32(*(*int8)(unsafe.Add(mBase, uint32(v44))))
					v46 = F_toupper(m, v45)
					mBase = m.M
					*(*uint8)(unsafe.Add(mBase, uint32(v44))) = uint8(v46)
					v49 = v41 + int32(1)
					if v49 != v36 {
						v41 = v49
						continue
					} else {
						break
					}
					break
				}
			}
		case 4:
			v35 = *(*int32)(unsafe.Add(mBase, uint32(v12+int32(-17))))
			v36 = v35
			if v36 == int32(0) {
			} else {
				v41 = int32(0)
				for {
					v44 = v12 + v41
					v45 = int32(*(*int8)(unsafe.Add(mBase, uint32(v44))))
					v46 = F_toupper(m, v45)
					mBase = m.M
					*(*uint8)(unsafe.Add(mBase, uint32(v44))) = uint8(v46)
					v49 = v41 + int32(1)
					if v49 != v36 {
						v41 = v49
						continue
					} else {
						break
					}
					break
				}
			}
		default:
		}
		v55 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
		v56 = *(*int32)(unsafe.Add(mBase, uint32(v55)+4))
		v57 = F_objectGetVal(m, v56)
		mBase = m.M
		*(*int32)(unsafe.Add(mBase, uint32(v7)+4)) = v12
		*(*int32)(unsafe.Add(mBase, uint32(v7))) = v57
		F_addReplyErrorFormat(m, l0, int32(_a1660), v7)
		mBase = m.M
		v62 = m.ExcPending
		if v62 != 0 {
			return
		} else {
			F_sdsfree(m, v12)
			mBase = m.M
			v64 = m.ExcPending
			if v64 != 0 {
				return
			} else {
				m.G0 = v7 + int32(16)
				return
			}
		}
	}
}
func F_callReplyCreateError(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
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
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v44 int32
	_ = v44
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	v8 = m.G0
	v10 = v8 - int32(16)
	m.G0 = v10
	v12 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v12 != int32(45) {
		v15 = F_sdsempty(m)
		mBase = m.M
		v18 = m.ExcPending
		if v18 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v10))) = l0
			v21 = F_sdscatfmt(m, v15, int32(_a194), v10)
			mBase = m.M
			v22 = m.ExcPending
			if v22 != 0 {
				return int32(0)
			} else {
				F_sdsfree(m, l0)
				mBase = m.M
				v24 = m.ExcPending
				if v24 != 0 {
					return int32(0)
				} else {
					v25 = v21
					v26 = F_listCreate(m)
					mBase = m.M
					v27 = m.ExcPending
					if v27 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v26)+12)) = int32(3)
						v30 = F_sdsnew(m, v25)
						mBase = m.M
						v31 = m.ExcPending
						if v31 != 0 {
							return int32(0)
						} else {
							v32 = F_listAddNodeTail(m, v26, v30)
							mBase = m.M
							v33 = m.ExcPending
							if v33 != 0 {
								return int32(0)
							} else {
								v35 = F_valkey_malloc(m, int32(48))
								mBase = m.M
								v36 = m.ExcPending
								if v36 != 0 {
									return int32(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v35)+8)) = v25
									*(*int32)(unsafe.Add(mBase, uint32(v35)+4)) = v25
									*(*int32)(unsafe.Add(mBase, uint32(v35)+20)) = int32(1)
									v44 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v25+int32(-1)))))
									switch v44 & int32(7) {
									case 0:
										v61 = int32(base.Ui32(v44) >> (uint(int32(3)) % 32))
									case 1:
										v51 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v25+int32(-3)))))
										v61 = v51
									case 2:
										v54 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v25+int32(-5)))))
										v61 = v54
									case 3:
										v57 = *(*int32)(unsafe.Add(mBase, uint32(v25+int32(-9))))
										v61 = v57
									case 4:
										v60 = *(*int32)(unsafe.Add(mBase, uint32(v25+int32(-17))))
										v61 = v60
									default:
										v61 = int32(0)
									}
									*(*int32)(unsafe.Add(mBase, uint32(v35)+44)) = int32(0)
									*(*int32)(unsafe.Add(mBase, uint32(v35))) = l1
									*(*int32)(unsafe.Add(mBase, uint32(v35)+12)) = v61
									*(*int32)(unsafe.Add(mBase, uint32(v35)+40)) = v26
									m.G0 = v10 + int32(16)
									return v35
								}
							}
						}
					}
				}
			}
		}
	} else {
		v25 = l0
		v26 = F_listCreate(m)
		mBase = m.M
		v27 = m.ExcPending
		if v27 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v26)+12)) = int32(3)
			v30 = F_sdsnew(m, v25)
			mBase = m.M
			v31 = m.ExcPending
			if v31 != 0 {
				return int32(0)
			} else {
				v32 = F_listAddNodeTail(m, v26, v30)
				mBase = m.M
				v33 = m.ExcPending
				if v33 != 0 {
					return int32(0)
				} else {
					v35 = F_valkey_malloc(m, int32(48))
					mBase = m.M
					v36 = m.ExcPending
					if v36 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v35)+8)) = v25
						*(*int32)(unsafe.Add(mBase, uint32(v35)+4)) = v25
						*(*int32)(unsafe.Add(mBase, uint32(v35)+20)) = int32(1)
						v44 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v25+int32(-1)))))
						switch v44 & int32(7) {
						case 0:
							v61 = int32(base.Ui32(v44) >> (uint(int32(3)) % 32))
						case 1:
							v51 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v25+int32(-3)))))
							v61 = v51
						case 2:
							v54 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v25+int32(-5)))))
							v61 = v54
						case 3:
							v57 = *(*int32)(unsafe.Add(mBase, uint32(v25+int32(-9))))
							v61 = v57
						case 4:
							v60 = *(*int32)(unsafe.Add(mBase, uint32(v25+int32(-17))))
							v61 = v60
						default:
							v61 = int32(0)
						}
						*(*int32)(unsafe.Add(mBase, uint32(v35)+44)) = int32(0)
						*(*int32)(unsafe.Add(mBase, uint32(v35))) = l1
						*(*int32)(unsafe.Add(mBase, uint32(v35)+12)) = v61
						*(*int32)(unsafe.Add(mBase, uint32(v35)+40)) = v26
						m.G0 = v10 + int32(16)
						return v35
					}
				}
			}
		}
	}
}
func F_callReplyGetAttribute(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	return v2
}
func F_callReplyGetBigNumber(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v16 int64
	_ = v16
	var v21 int64
	_ = v21
	var v26 int64
	_ = v26
	var v31 int64
	_ = v31
	var v36 int64
	_ = v36
	var v41 int64
	_ = v41
	var v46 int64
	_ = v46
	var v50 int64
	_ = v50
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	v5 = m.G0
	v7 = v5 - int32(80)
	m.G0 = v7
	v9 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+20)))
	if v9&int32(2) != 0 {
		v64 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
		if v64 != int32(9) {
			v70 = int32(0)
		} else {
			v67 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
			*(*int32)(unsafe.Add(mBase, uint32(l1))) = v67
			v69 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
			v70 = v69
		}
		m.G0 = v7 + int32(80)
		return v70
	} else {
		v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		v15 = int32(0)
		v16 = *(*int64)(unsafe.Add(mBase, _consts[82]))
		*(*int64)(unsafe.Add(mBase, uint32(v7+int32(24)))) = v16
		v21 = *(*int64)(unsafe.Add(mBase, _consts[83]))
		*(*int64)(unsafe.Add(mBase, uint32(v7+int32(32)))) = v21
		v26 = *(*int64)(unsafe.Add(mBase, _consts[84]))
		*(*int64)(unsafe.Add(mBase, uint32(v7+int32(40)))) = v26
		v31 = *(*int64)(unsafe.Add(mBase, _consts[85]))
		*(*int64)(unsafe.Add(mBase, uint32(v7+int32(48)))) = v31
		v36 = *(*int64)(unsafe.Add(mBase, _consts[86]))
		*(*int64)(unsafe.Add(mBase, uint32(v7+int32(56)))) = v36
		v41 = *(*int64)(unsafe.Add(mBase, _consts[87]))
		*(*int64)(unsafe.Add(mBase, uint32(v7+int32(64)))) = v41
		v46 = *(*int64)(unsafe.Add(mBase, _consts[88]))
		*(*int64)(unsafe.Add(mBase, uint32(v7+int32(72)))) = v46
		*(*int32)(unsafe.Add(mBase, uint32(v7)+12)) = v12
		v50 = *(*int64)(unsafe.Add(mBase, _consts[89]))
		*(*int64)(unsafe.Add(mBase, uint32(v7)+16)) = v50
		v54 = F_parseReply(m, v7+int32(12), l0)
		mBase = m.M
		v57 = m.ExcPending
		if v57 != 0 {
			return int32(0)
		} else {
			v58 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
			*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v58 | int32(2)
			v64 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
			if v64 != int32(9) {
				v70 = int32(0)
			} else {
				v67 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
				*(*int32)(unsafe.Add(mBase, uint32(l1))) = v67
				v69 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
				v70 = v69
			}
			m.G0 = v7 + int32(80)
			return v70
		}
	}
}
func F_callReplyGetMapElement(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
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
	var v19 int32
	_ = v19
	var v20 int64
	_ = v20
	var v25 int64
	_ = v25
	var v30 int64
	_ = v30
	var v35 int64
	_ = v35
	var v40 int64
	_ = v40
	var v45 int64
	_ = v45
	var v50 int64
	_ = v50
	var v54 int64
	_ = v54
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v81 int32
	_ = v81
	var v85 int32
	_ = v85
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v95 int32
	_ = v95
	var v99 int32
	_ = v99
	var v103 int32
	_ = v103
	var v107 int32
	_ = v107
	v9 = m.G0
	v11 = v9 - int32(80)
	m.G0 = v11
	v13 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+20)))
	if v13&int32(2) != 0 {
		v67 = int32(-1)
		v68 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
		if v68 != int32(5) {
			v107 = v67
		} else {
			v71 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
			if base.Ui32(v71) <= base.Ui32(l1) {
				v107 = v67
			} else {
				if l2 == int32(0) {
				} else {
					v76 = int32(1)
					v77 = l1 << (uint(v76) % 32)
					if base.Ui32(v71<<(uint(v76)%32)) <= base.Ui32(v77) {
						v85 = int32(0)
					} else {
						v81 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
						v85 = v81 + v77*int32(48)
					}
					*(*int32)(unsafe.Add(mBase, uint32(l2))) = v85
				}
				if l3 != 0 {
					v90 = int32(0)
					v92 = int32(1)
					v95 = l1<<(uint(v92)%32) | v92
					if base.Ui32(v71<<(uint(v92)%32)) <= base.Ui32(v95) {
						v103 = v90
					} else {
						v99 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
						v103 = v99 + v95*int32(48)
					}
					*(*int32)(unsafe.Add(mBase, uint32(l3))) = v103
					v107 = v90
				} else {
					v107 = int32(0)
				}
			}
		}
		m.G0 = v11 + int32(80)
		return v107
	} else {
		v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		v19 = int32(0)
		v20 = *(*int64)(unsafe.Add(mBase, _consts[82]))
		*(*int64)(unsafe.Add(mBase, uint32(v11+int32(24)))) = v20
		v25 = *(*int64)(unsafe.Add(mBase, _consts[83]))
		*(*int64)(unsafe.Add(mBase, uint32(v11+int32(32)))) = v25
		v30 = *(*int64)(unsafe.Add(mBase, _consts[84]))
		*(*int64)(unsafe.Add(mBase, uint32(v11+int32(40)))) = v30
		v35 = *(*int64)(unsafe.Add(mBase, _consts[85]))
		*(*int64)(unsafe.Add(mBase, uint32(v11+int32(48)))) = v35
		v40 = *(*int64)(unsafe.Add(mBase, _consts[86]))
		*(*int64)(unsafe.Add(mBase, uint32(v11+int32(56)))) = v40
		v45 = *(*int64)(unsafe.Add(mBase, _consts[87]))
		*(*int64)(unsafe.Add(mBase, uint32(v11+int32(64)))) = v45
		v50 = *(*int64)(unsafe.Add(mBase, _consts[88]))
		*(*int64)(unsafe.Add(mBase, uint32(v11+int32(72)))) = v50
		*(*int32)(unsafe.Add(mBase, uint32(v11)+12)) = v16
		v54 = *(*int64)(unsafe.Add(mBase, _consts[89]))
		*(*int64)(unsafe.Add(mBase, uint32(v11)+16)) = v54
		v58 = F_parseReply(m, v11+int32(12), l0)
		mBase = m.M
		v61 = m.ExcPending
		if v61 != 0 {
			return int32(0)
		} else {
			v62 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
			*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v62 | int32(2)
			v67 = int32(-1)
			v68 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
			if v68 != int32(5) {
				v107 = v67
			} else {
				v71 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
				if base.Ui32(v71) <= base.Ui32(l1) {
					v107 = v67
				} else {
					if l2 == int32(0) {
					} else {
						v76 = int32(1)
						v77 = l1 << (uint(v76) % 32)
						if base.Ui32(v71<<(uint(v76)%32)) <= base.Ui32(v77) {
							v85 = int32(0)
						} else {
							v81 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
							v85 = v81 + v77*int32(48)
						}
						*(*int32)(unsafe.Add(mBase, uint32(l2))) = v85
					}
					if l3 != 0 {
						v90 = int32(0)
						v92 = int32(1)
						v95 = l1<<(uint(v92)%32) | v92
						if base.Ui32(v71<<(uint(v92)%32)) <= base.Ui32(v95) {
							v103 = v90
						} else {
							v99 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
							v103 = v99 + v95*int32(48)
						}
						*(*int32)(unsafe.Add(mBase, uint32(l3))) = v103
						v107 = v90
					} else {
						v107 = int32(0)
					}
				}
			}
			m.G0 = v11 + int32(80)
			return v107
		}
	}
}
func F_callReplyGetString(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v17 int64
	_ = v17
	var v22 int64
	_ = v22
	var v27 int64
	_ = v27
	var v32 int64
	_ = v32
	var v37 int64
	_ = v37
	var v42 int64
	_ = v42
	var v47 int64
	_ = v47
	var v51 int64
	_ = v51
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	v6 = m.G0
	v8 = v6 - int32(80)
	m.G0 = v8
	v10 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+20)))
	if v10&int32(2) != 0 {
		v64 = int32(0)
		v65 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
		if base.Ui32(int32(13)) < base.Ui32(v65) {
			v79 = v64
		} else {
			if int32(1)<<(uint(v65)%32)&int32(8195) == int32(0) {
				v79 = v64
			} else {
				if l1 == int32(0) {
				} else {
					v76 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
					*(*int32)(unsafe.Add(mBase, uint32(l1))) = v76
				}
				v78 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
				v79 = v78
			}
		}
		m.G0 = v8 + int32(80)
		return v79
	} else {
		v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		v16 = int32(0)
		v17 = *(*int64)(unsafe.Add(mBase, _consts[82]))
		*(*int64)(unsafe.Add(mBase, uint32(v8+int32(24)))) = v17
		v22 = *(*int64)(unsafe.Add(mBase, _consts[83]))
		*(*int64)(unsafe.Add(mBase, uint32(v8+int32(32)))) = v22
		v27 = *(*int64)(unsafe.Add(mBase, _consts[84]))
		*(*int64)(unsafe.Add(mBase, uint32(v8+int32(40)))) = v27
		v32 = *(*int64)(unsafe.Add(mBase, _consts[85]))
		*(*int64)(unsafe.Add(mBase, uint32(v8+int32(48)))) = v32
		v37 = *(*int64)(unsafe.Add(mBase, _consts[86]))
		*(*int64)(unsafe.Add(mBase, uint32(v8+int32(56)))) = v37
		v42 = *(*int64)(unsafe.Add(mBase, _consts[87]))
		*(*int64)(unsafe.Add(mBase, uint32(v8+int32(64)))) = v42
		v47 = *(*int64)(unsafe.Add(mBase, _consts[88]))
		*(*int64)(unsafe.Add(mBase, uint32(v8+int32(72)))) = v47
		*(*int32)(unsafe.Add(mBase, uint32(v8)+12)) = v13
		v51 = *(*int64)(unsafe.Add(mBase, _consts[89]))
		*(*int64)(unsafe.Add(mBase, uint32(v8)+16)) = v51
		v55 = F_parseReply(m, v8+int32(12), l0)
		mBase = m.M
		v58 = m.ExcPending
		if v58 != 0 {
			return int32(0)
		} else {
			v59 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
			*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v59 | int32(2)
			v64 = int32(0)
			v65 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
			if base.Ui32(int32(13)) < base.Ui32(v65) {
				v79 = v64
			} else {
				if int32(1)<<(uint(v65)%32)&int32(8195) == int32(0) {
					v79 = v64
				} else {
					if l1 == int32(0) {
					} else {
						v76 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
						*(*int32)(unsafe.Add(mBase, uint32(l1))) = v76
					}
					v78 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
					v79 = v78
				}
			}
			m.G0 = v8 + int32(80)
			return v79
		}
	}
}
func F_callReplyIsResp3(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	return v2 & int32(4)
}
func F_callReplyMap(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
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
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v52 int32
	_ = v52
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v82 int32
	_ = v82
	var v87 int32
	_ = v87
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	*(*int32)(unsafe.Add(mBase, uint32(l1)+24)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(l1)+16)) = int32(5)
	v14 = F_valkey_calloc(m, l2*int32(96))
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+32)) = v14
	v18 = l2 << (uint(int32(1)) % 32)
	if v18 == int32(0) {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+8)) = l3
	v98 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v99 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	*(*int32)(unsafe.Add(mBase, uint32(l1)+20)) = v99 | int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(l1)+12)) = v98 - l3
	return
L4:
	;
	v26 = v14
	v28 = int32(0)
	goto L5
L5:
	;
	v31 = v28 * int32(48)
	v33 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	*(*int32)(unsafe.Add(mBase, uint32(v26+v31))) = v33
	v35 = *(*int32)(unsafe.Add(mBase, uint32(l1)+32))
	v37 = F_parseReply(m, l0, v35+v31)
	mBase = m.M
	v38 = m.ExcPending
	if v38 != 0 {
		goto L1
	} else {
		goto L7
	}
L6:
	;
	goto L3
L7:
	;
	v39 = *(*int32)(unsafe.Add(mBase, uint32(l1)+32))
	v40 = v39 + v31
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v40)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v40)+20)) = v41 | int32(2)
	v45 = *(*int32)(unsafe.Add(mBase, uint32(l1)+32))
	v47 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v45+v31)+20)))
	if v47&int32(4) == int32(0) {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v58 = int32(48)
	v59 = (v28 | int32(1)) * v58
	v61 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	*(*int32)(unsafe.Add(mBase, uint32(v45+v59))) = v61
	v63 = *(*int32)(unsafe.Add(mBase, uint32(l1)+32))
	v67 = F_parseReply(m, l0, v63+v31+v58)
	mBase = m.M
	v68 = m.ExcPending
	if v68 != 0 {
		goto L1
	} else {
		goto L10
	}
L9:
	;
	v52 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	*(*int32)(unsafe.Add(mBase, uint32(l1)+20)) = v52 | int32(4)
	goto L8
L10:
	;
	v69 = *(*int32)(unsafe.Add(mBase, uint32(l1)+32))
	v70 = v69 + v59
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v70)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v70)+20)) = v71 | int32(2)
	v75 = *(*int32)(unsafe.Add(mBase, uint32(l1)+32))
	v77 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v75+v59)+20)))
	if v77&int32(4) == int32(0) {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v87 = v28 + int32(2)
	if base.Ui32(v87) < base.Ui32(v18) {
		v26 = v75
		v28 = v87
		goto L5
	} else {
		goto L13
	}
L12:
	;
	v82 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	*(*int32)(unsafe.Add(mBase, uint32(l1)+20)) = v82 | int32(4)
	goto L11
L13:
	;
	goto L6
}
func F_callReplyNullArray(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = l1
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v8&int32(8) != 0 {
		v11 = int32(14)
	} else {
		v11 = int32(4)
	}
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v11
	return
}
func F_freeCallReply(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v12 int32
	_ = v12
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v3&int32(1) == int32(0) {
		return
	} else {
		if v3&int32(2) == int32(0) {
			v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			F_sdsfree(m, v17)
			mBase = m.M
			v19 = m.ExcPending
			if v19 != 0 {
				return
			} else {
				v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
				if v20 == int32(0) {
					F_valkey_free(m, l0)
					mBase = m.M
					v27 = m.ExcPending
					if v27 != 0 {
						return
					} else {
						return
					}
				} else {
					F_listRelease(m, v20)
					mBase = m.M
					v24 = m.ExcPending
					if v24 != 0 {
						return
					} else {
						F_valkey_free(m, l0)
						mBase = m.M
						v27 = m.ExcPending
						if v27 != 0 {
							return
						} else {
							return
						}
					}
				}
			}
		} else {
			v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
			if v12 == int32(12) {
				F_valkey_free(m, l0)
				mBase = m.M
				v27 = m.ExcPending
				if v27 != 0 {
					return
				} else {
					return
				}
			} else {
				F_freeCallReplyInternal(m, l0)
				mBase = m.M
				v16 = m.ExcPending
				if v16 != 0 {
					return
				} else {
					v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
					F_sdsfree(m, v17)
					mBase = m.M
					v19 = m.ExcPending
					if v19 != 0 {
						return
					} else {
						v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
						if v20 == int32(0) {
							F_valkey_free(m, l0)
							mBase = m.M
							v27 = m.ExcPending
							if v27 != 0 {
								return
							} else {
								return
							}
						} else {
							F_listRelease(m, v20)
							mBase = m.M
							v24 = m.ExcPending
							if v24 != 0 {
								return
							} else {
								F_valkey_free(m, l0)
								mBase = m.M
								v27 = m.ExcPending
								if v27 != 0 {
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
	}
}
func F_invokeReplyHandlers(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
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
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v53 int32
	_ = v53
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v83 int32
	_ = v83
	var v91 int32
	_ = v91
	var v96 int32
	_ = v96
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v106 int32
	_ = v106
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
	var v121 int32
	_ = v121
	var v125 int32
	_ = v125
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v131 int32
	_ = v131
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
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v153 int32
	_ = v153
	var v155 int32
	_ = v155
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v168 int32
	_ = v168
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v179 int32
	_ = v179
	var v180 int64
	_ = v180
	var v185 int64
	_ = v185
	var v190 int64
	_ = v190
	var v195 int64
	_ = v195
	var v200 int64
	_ = v200
	var v205 int64
	_ = v205
	var v210 int64
	_ = v210
	var v214 int64
	_ = v214
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v223 int32
	_ = v223
	var v231 int32
	_ = v231
	v12 = m.G0
	v14 = v12 - int32(80)
	m.G0 = v14
	v16 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+200)))
	if v16&int32(16) != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	F__serverAssert(m, int32(_a197), int32(_a196), int32(818))
	mBase = m.M
	v231 = m.ExcPending
	if v231 != 0 {
		goto L19
	} else {
		goto L42
	}
L2:
	;
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l1)+132))
	v20 = *(*int32)(unsafe.Add(mBase, uint32(v19)+20))
	if v20 != 0 {
		goto L4
	} else {
		goto L5
	}
L3:
	;
	v168 = *(*int32)(unsafe.Add(mBase, uint32(l2)+92))
	if v168 == int32(0) {
		goto L34
	} else {
		goto L35
	}
L4:
	;
	v35 = v14 + int32(4)
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v19)))
	*(*int32)(unsafe.Add(mBase, uint32(v35)+4)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v35))) = v36
	goto L7
L5:
	;
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l1)+180))
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l1)+128))
	if base.Ui32(v22) <= base.Ui32(v21) {
		goto L4
	} else {
		goto L6
	}
L6:
	;
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l1)+124))
	v26 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v24+v21))) = uint8(v26)
	v28 = *(*int32)(unsafe.Add(mBase, uint32(l1)+180))
	*(*int32)(unsafe.Add(mBase, uint32(l1)+180)) = v26
	v31 = *(*int32)(unsafe.Add(mBase, uint32(l1)+124))
	v163 = v28
	v164 = v31
	v165 = int32(1)
	goto L3
L7:
	;
	v40 = *(*int32)(unsafe.Add(mBase, uint32(l1)+180))
	v42 = v14 + int32(4)
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v42)))
	if v44 == int32(0) {
		goto L10
	} else {
		goto L11
	}
L8:
	;
	v96 = int32(0)
	v100 = F_zmalloc_usable(m, v91+int32(1), v96)
	mBase = m.M
	v101 = m.ExcPending
	if v101 != 0 {
		goto L19
	} else {
		goto L20
	}
L9:
	;
	if v44 == int32(0) {
		v91 = v40
		goto L8
	} else {
		goto L12
	}
L10:
	;
	goto L9
L11:
	;
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v42)+4))
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v44+base.B2i32(v47 == int32(0))<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v42))) = v53
	goto L10
L12:
	;
	v62 = v44
	v63 = v40
	goto L13
L13:
	;
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v62)+8))
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v68)+4))
	v70 = v69 + v63
	v72 = v14 + int32(4)
	v74 = *(*int32)(unsafe.Add(mBase, uint32(v72)))
	if v74 == int32(0) {
		goto L16
	} else {
		goto L17
	}
L14:
	;
	v91 = v70
	goto L8
L15:
	;
	if v74 != 0 {
		v62 = v74
		v63 = v70
		goto L13
	} else {
		goto L18
	}
L16:
	;
	goto L15
L17:
	;
	v77 = *(*int32)(unsafe.Add(mBase, uint32(v72)+4))
	v83 = *(*int32)(unsafe.Add(mBase, uint32(v74+base.B2i32(v77 == int32(0))<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v72))) = v83
	goto L16
L18:
	;
	goto L14
L19:
	;
	return
L20:
	;
	v102 = *(*int32)(unsafe.Add(mBase, uint32(l1)+124))
	v103 = *(*int32)(unsafe.Add(mBase, uint32(l1)+180))
	if v103 == int32(0) {
		v107 = v100
		goto L22
	} else {
		goto L23
	}
L21:
	;
	v108 = *(*int32)(unsafe.Add(mBase, uint32(l1)+180))
	v109 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l1)+180)) = v109
	v111 = v107 + v108
	v112 = *(*int32)(unsafe.Add(mBase, uint32(l1)+132))
	v113 = *(*int32)(unsafe.Add(mBase, uint32(v112)+20))
	if v113 == v109 {
		v153 = v111
		goto L24
	} else {
		goto L25
	}
L22:
	;
	goto L21
L23:
	;
	v106 = F__emscripten_memcpy_bulkmem(m, v100, v102, v103)
	mBase = m.M
	v107 = v106
	goto L22
L24:
	;
	v155 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v153))) = uint8(v155)
	v163 = v91
	v164 = v100
	v165 = v96
	goto L3
L25:
	;
	v121 = v112
	v125 = v111
	goto L26
L26:
	;
	v127 = *(*int32)(unsafe.Add(mBase, uint32(v121)))
	v128 = *(*int32)(unsafe.Add(mBase, uint32(v127)+8))
	v131 = *(*int32)(unsafe.Add(mBase, uint32(v128)+4))
	if v131 == int32(0) {
		v135 = v125
		goto L29
	} else {
		goto L30
	}
L27:
	;
	v153 = v141
	goto L24
L28:
	;
	v136 = *(*int32)(unsafe.Add(mBase, uint32(v128)+4))
	v137 = *(*int32)(unsafe.Add(mBase, uint32(l1)+132))
	v138 = *(*int32)(unsafe.Add(mBase, uint32(v137)))
	F_listDelNode(m, v137, v138)
	mBase = m.M
	v140 = m.ExcPending
	if v140 != 0 {
		goto L19
	} else {
		goto L31
	}
L29:
	;
	goto L28
L30:
	;
	v134 = F__emscripten_memcpy_bulkmem(m, v125, v128+int32(13), v131)
	mBase = m.M
	v135 = v134
	goto L29
L31:
	;
	v141 = v135 + v136
	v142 = *(*int32)(unsafe.Add(mBase, uint32(l1)+132))
	v143 = *(*int32)(unsafe.Add(mBase, uint32(v142)+20))
	if v143 != 0 {
		v121 = v142
		v125 = v141
		goto L26
	} else {
		goto L32
	}
L32:
	;
	goto L27
L33:
	;
	if v165 != 0 {
		goto L39
	} else {
		goto L40
	}
L34:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+76)) = l3
	*(*int32)(unsafe.Add(mBase, uint32(v14)+72)) = l2
	v179 = int32(0)
	v180 = *(*int64)(unsafe.Add(mBase, _consts[90]))
	*(*int64)(unsafe.Add(mBase, uint32(v14+int32(16)))) = v180
	v185 = *(*int64)(unsafe.Add(mBase, _consts[91]))
	*(*int64)(unsafe.Add(mBase, uint32(v14+int32(24)))) = v185
	v190 = *(*int64)(unsafe.Add(mBase, _consts[92]))
	*(*int64)(unsafe.Add(mBase, uint32(v14+int32(32)))) = v190
	v195 = *(*int64)(unsafe.Add(mBase, _consts[93]))
	*(*int64)(unsafe.Add(mBase, uint32(v14+int32(40)))) = v195
	v200 = *(*int64)(unsafe.Add(mBase, _consts[94]))
	*(*int64)(unsafe.Add(mBase, uint32(v14+int32(48)))) = v200
	v205 = *(*int64)(unsafe.Add(mBase, _consts[95]))
	*(*int64)(unsafe.Add(mBase, uint32(v14+int32(56)))) = v205
	v210 = *(*int64)(unsafe.Add(mBase, _consts[96]))
	*(*int64)(unsafe.Add(mBase, uint32(v14+int32(64)))) = v210
	*(*int32)(unsafe.Add(mBase, uint32(v14)+4)) = v164
	v214 = *(*int64)(unsafe.Add(mBase, _consts[97]))
	*(*int64)(unsafe.Add(mBase, uint32(v14)+8)) = v214
	v220 = F_parseReply(m, v14+int32(4), v14+int32(72))
	mBase = m.M
	v221 = m.ExcPending
	if v221 != 0 {
		goto L19
	} else {
		goto L38
	}
L35:
	;
	v171 = m.T0[v168].(func(*base.Module, int32, int32, int32, int32) int32)(m, l3, l0, v164, v163)
	mBase = m.M
	v172 = m.ExcPending
	if v172 != 0 {
		goto L19
	} else {
		goto L36
	}
L36:
	;
	if v171 == int32(0) {
		goto L33
	} else {
		goto L37
	}
L37:
	;
	goto L34
L38:
	;
	goto L33
L39:
	;
	m.G0 = v14 + int32(80)
	return
L40:
	;
	F_valkey_free(m, v164)
	mBase = m.M
	v223 = m.ExcPending
	if v223 != 0 {
		goto L19
	} else {
		goto L41
	}
L41:
	;
	goto L39
L42:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_parseReply(m *base.Module, l0 int32, l1 int32) int32 {
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
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v53 int32
	_ = v53
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v84 int64
	_ = v84
	var v90 int32
	_ = v90
	var v92 int64
	_ = v92
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v106 int64
	_ = v106
	var v111 int64
	_ = v111
	var v115 int32
	_ = v115
	var v117 int64
	_ = v117
	var v119 int32
	_ = v119
	var v127 int64
	_ = v127
	var v151 int64
	_ = v151
	var v177 int64
	_ = v177
	var v180 int32
	_ = v180
	var v182 int32
	_ = v182
	var v186 int32
	_ = v186
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v192 int32
	_ = v192
	var v195 int32
	_ = v195
	var v197 int32
	_ = v197
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v204 int32
	_ = v204
	var v208 int32
	_ = v208
	var v210 int32
	_ = v210
	var v216 int32
	_ = v216
	var v218 int32
	_ = v218
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v225 int32
	_ = v225
	var v229 int32
	_ = v229
	var v231 int32
	_ = v231
	var v237 int32
	_ = v237
	var v239 int32
	_ = v239
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v246 int32
	_ = v246
	var v250 int32
	_ = v250
	var v256 int32
	_ = v256
	var v258 int32
	_ = v258
	var v271 int32
	_ = v271
	var v272 int32
	_ = v272
	var v276 int32
	_ = v276
	var v292 int32
	_ = v292
	var v294 int32
	_ = v294
	var v295 int32
	_ = v295
	var v296 int32
	_ = v296
	var v307 int64
	_ = v307
	var v313 int32
	_ = v313
	var v315 int64
	_ = v315
	var v317 int32
	_ = v317
	var v319 int32
	_ = v319
	var v329 int64
	_ = v329
	var v334 int64
	_ = v334
	var v338 int32
	_ = v338
	var v340 int64
	_ = v340
	var v342 int32
	_ = v342
	var v350 int64
	_ = v350
	var v374 int64
	_ = v374
	var v400 int64
	_ = v400
	var v401 int32
	_ = v401
	var v403 int32
	_ = v403
	var v405 int32
	_ = v405
	var v408 int32
	_ = v408
	var v409 int32
	_ = v409
	var v410 int32
	_ = v410
	var v412 int32
	_ = v412
	var v416 int32
	_ = v416
	var v419 int32
	_ = v419
	var v421 int32
	_ = v421
	var v434 int32
	_ = v434
	var v435 int32
	_ = v435
	var v439 int32
	_ = v439
	var v455 int32
	_ = v455
	var v457 int32
	_ = v457
	var v458 int32
	_ = v458
	var v459 int32
	_ = v459
	var v470 int64
	_ = v470
	var v476 int32
	_ = v476
	var v478 int64
	_ = v478
	var v480 int32
	_ = v480
	var v482 int32
	_ = v482
	var v492 int64
	_ = v492
	var v497 int64
	_ = v497
	var v501 int32
	_ = v501
	var v503 int64
	_ = v503
	var v505 int32
	_ = v505
	var v513 int64
	_ = v513
	var v537 int64
	_ = v537
	var v564 int32
	_ = v564
	var v566 int64
	_ = v566
	var v570 int32
	_ = v570
	var v572 int32
	_ = v572
	var v575 int32
	_ = v575
	var v577 int32
	_ = v577
	var v580 int32
	_ = v580
	var v581 int32
	_ = v581
	var v582 int32
	_ = v582
	var v584 int32
	_ = v584
	var v588 int32
	_ = v588
	var v591 int32
	_ = v591
	var v593 int32
	_ = v593
	var v606 int32
	_ = v606
	var v607 int32
	_ = v607
	var v611 int32
	_ = v611
	var v627 int32
	_ = v627
	var v629 int32
	_ = v629
	var v630 int32
	_ = v630
	var v631 int32
	_ = v631
	var v642 int64
	_ = v642
	var v648 int32
	_ = v648
	var v650 int64
	_ = v650
	var v652 int32
	_ = v652
	var v654 int32
	_ = v654
	var v664 int64
	_ = v664
	var v669 int64
	_ = v669
	var v673 int32
	_ = v673
	var v675 int64
	_ = v675
	var v677 int32
	_ = v677
	var v685 int64
	_ = v685
	var v709 int64
	_ = v709
	var v738 int32
	_ = v738
	var v739 int32
	_ = v739
	var v741 int32
	_ = v741
	var v744 int32
	_ = v744
	var v745 int32
	_ = v745
	var v746 int32
	_ = v746
	var v748 int32
	_ = v748
	var v752 int32
	_ = v752
	var v755 int32
	_ = v755
	var v757 int32
	_ = v757
	var v770 int32
	_ = v770
	var v771 int32
	_ = v771
	var v775 int32
	_ = v775
	var v791 int32
	_ = v791
	var v793 int32
	_ = v793
	var v794 int32
	_ = v794
	var v795 int32
	_ = v795
	var v806 int64
	_ = v806
	var v812 int32
	_ = v812
	var v814 int64
	_ = v814
	var v816 int32
	_ = v816
	var v818 int32
	_ = v818
	var v828 int64
	_ = v828
	var v833 int64
	_ = v833
	var v837 int32
	_ = v837
	var v839 int64
	_ = v839
	var v841 int32
	_ = v841
	var v849 int64
	_ = v849
	var v873 int64
	_ = v873
	var v902 int32
	_ = v902
	var v903 int32
	_ = v903
	var v905 int32
	_ = v905
	var v909 int32
	_ = v909
	var v910 int32
	_ = v910
	var v912 int32
	_ = v912
	var v916 int32
	_ = v916
	var v918 int32
	_ = v918
	var v920 int32
	_ = v920
	var v924 int32
	_ = v924
	var v926 int32
	_ = v926
	var v929 int32
	_ = v929
	var v930 int32
	_ = v930
	var v931 int32
	_ = v931
	var v933 int32
	_ = v933
	var v937 int32
	_ = v937
	var v939 int32
	_ = v939
	var v944 int32
	_ = v944
	var v947 int32
	_ = v947
	var v951 int32
	_ = v951
	var v953 int32
	_ = v953
	var v955 int32
	_ = v955
	var v961 int64
	_ = v961
	var v966 int64
	_ = v966
	var v974 int32
	_ = v974
	var v981 int32
	_ = v981
	var v987 float64
	_ = v987
	var v991 int32
	_ = v991
	var v992 int32
	_ = v992
	var v993 float64
	_ = v993
	var v995 int32
	_ = v995
	var v997 int32
	_ = v997
	var v1001 int32
	_ = v1001
	var v1002 int32
	_ = v1002
	var v1004 int32
	_ = v1004
	var v1008 int32
	_ = v1008
	var v1010 int32
	_ = v1010
	var v1013 int32
	_ = v1013
	var v1015 int32
	_ = v1015
	var v1018 int32
	_ = v1018
	var v1019 int32
	_ = v1019
	var v1020 int32
	_ = v1020
	var v1022 int32
	_ = v1022
	var v1026 int32
	_ = v1026
	var v1028 int32
	_ = v1028
	var v1034 int32
	_ = v1034
	var v1036 int32
	_ = v1036
	var v1039 int32
	_ = v1039
	var v1040 int32
	_ = v1040
	var v1041 int32
	_ = v1041
	var v1043 int32
	_ = v1043
	var v1047 int32
	_ = v1047
	var v1053 int32
	_ = v1053
	var v1055 int32
	_ = v1055
	var v1068 int32
	_ = v1068
	var v1069 int32
	_ = v1069
	var v1073 int32
	_ = v1073
	var v1089 int32
	_ = v1089
	var v1091 int32
	_ = v1091
	var v1092 int32
	_ = v1092
	var v1093 int32
	_ = v1093
	var v1104 int64
	_ = v1104
	var v1110 int32
	_ = v1110
	var v1112 int64
	_ = v1112
	var v1114 int32
	_ = v1114
	var v1116 int32
	_ = v1116
	var v1126 int64
	_ = v1126
	var v1131 int64
	_ = v1131
	var v1135 int32
	_ = v1135
	var v1137 int64
	_ = v1137
	var v1139 int32
	_ = v1139
	var v1147 int64
	_ = v1147
	var v1171 int64
	_ = v1171
	var v1197 int32
	_ = v1197
	var v1198 int32
	_ = v1198
	var v1201 int32
	_ = v1201
	var v1208 int32
	_ = v1208
	var v1210 int32
	_ = v1210
	var v1213 int32
	_ = v1213
	var v1214 int32
	_ = v1214
	var v1215 int32
	_ = v1215
	var v1217 int32
	_ = v1217
	var v1221 int32
	_ = v1221
	var v1224 int32
	_ = v1224
	var v1226 int32
	_ = v1226
	var v1239 int32
	_ = v1239
	var v1240 int32
	_ = v1240
	var v1244 int32
	_ = v1244
	var v1260 int32
	_ = v1260
	var v1262 int32
	_ = v1262
	var v1263 int32
	_ = v1263
	var v1264 int32
	_ = v1264
	var v1275 int64
	_ = v1275
	var v1281 int32
	_ = v1281
	var v1283 int64
	_ = v1283
	var v1285 int32
	_ = v1285
	var v1287 int32
	_ = v1287
	var v1297 int64
	_ = v1297
	var v1302 int64
	_ = v1302
	var v1306 int32
	_ = v1306
	var v1308 int64
	_ = v1308
	var v1310 int32
	_ = v1310
	var v1318 int64
	_ = v1318
	var v1342 int64
	_ = v1342
	var v1371 int32
	_ = v1371
	var v1372 int32
	_ = v1372
	var v1374 int32
	_ = v1374
	var v1376 int32
	_ = v1376
	var v1377 int32
	_ = v1377
	var v1381 int32
	_ = v1381
	var v1383 int32
	_ = v1383
	var v1393 int32
	_ = v1393
	var v1394 int32
	_ = v1394
	var v1395 int32
	_ = v1395
	var v1396 int32
	_ = v1396
	var v1397 int32
	_ = v1397
	var v1398 int32
	_ = v1398
	var v1399 int32
	_ = v1399
	v10 = m.G0
	v12 = v10 - int32(16)
	m.G0 = v12
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v15 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14))))
	switch v15 + int32(-35) {
	case 0:
		goto L8
	case 1:
		goto L15
	case 2:
		goto L9
	default:
		goto L2
	case 5:
		goto L5
	case 7:
		goto L11
	case 8:
		goto L14
	case 9:
		goto L7
	case 10:
		goto L13
	case 23:
		goto L12
	case 26:
		goto L4
	case 60:
		goto L6
	case 89:
		goto L3
	case 91:
		goto L10
	}
L1:
	;
	m.G0 = v12 + int32(16)
	return v1383
L2:
	;
	v1376 = int32(-1)
	v1377 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
	if v1377 == int32(0) {
		v1383 = v1376
		goto L1
	} else {
		goto L287
	}
L3:
	;
	v1213 = v14 + int32(1)
	v1214 = int32(13)
	v1215 = F___strchrnul(m, v1213, v1214)
	mBase = m.M
	v1217 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1215))))
	if v1217 == v1214 {
		goto L256
	} else {
		goto L257
	}
L4:
	;
	v1039 = v14 + int32(1)
	v1040 = int32(13)
	v1041 = F___strchrnul(m, v1039, v1040)
	mBase = m.M
	v1043 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1041))))
	if v1043 == v1040 {
		goto L224
	} else {
		goto L225
	}
L5:
	;
	v1018 = v14 + int32(1)
	v1019 = int32(13)
	v1020 = F___strchrnul(m, v1018, v1019)
	mBase = m.M
	v1022 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1020))))
	if v1022 == v1019 {
		goto L219
	} else {
		goto L220
	}
L6:
	;
	v1001 = int32(13)
	v1002 = F___strchrnul(m, v14+int32(1), v1001)
	mBase = m.M
	v1004 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1002))))
	if v1004 == v1001 {
		goto L214
	} else {
		goto L215
	}
L7:
	;
	v929 = v14 + int32(1)
	v930 = int32(13)
	v931 = F___strchrnul(m, v929, v930)
	mBase = m.M
	v933 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v931))))
	if v933 == v930 {
		goto L199
	} else {
		goto L200
	}
L8:
	;
	v909 = int32(13)
	v910 = F___strchrnul(m, v14+int32(1), v909)
	mBase = m.M
	v912 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v910))))
	if v912 == v909 {
		goto L194
	} else {
		goto L195
	}
L9:
	;
	v744 = v14 + int32(1)
	v745 = int32(13)
	v746 = F___strchrnul(m, v744, v745)
	mBase = m.M
	v748 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v746))))
	if v748 == v745 {
		goto L162
	} else {
		goto L163
	}
L10:
	;
	v580 = v14 + int32(1)
	v581 = int32(13)
	v582 = F___strchrnul(m, v580, v581)
	mBase = m.M
	v584 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v582))))
	if v584 == v581 {
		goto L130
	} else {
		goto L131
	}
L11:
	;
	v408 = v14 + int32(1)
	v409 = int32(13)
	v410 = F___strchrnul(m, v408, v409)
	mBase = m.M
	v412 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v410))))
	if v412 == v409 {
		goto L95
	} else {
		goto L96
	}
L12:
	;
	v242 = v14 + int32(1)
	v243 = int32(13)
	v244 = F___strchrnul(m, v242, v243)
	mBase = m.M
	v246 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v244))))
	if v246 == v243 {
		goto L63
	} else {
		goto L64
	}
L13:
	;
	v221 = v14 + int32(1)
	v222 = int32(13)
	v223 = F___strchrnul(m, v221, v222)
	mBase = m.M
	v225 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v223))))
	if v225 == v222 {
		goto L58
	} else {
		goto L59
	}
L14:
	;
	v200 = v14 + int32(1)
	v201 = int32(13)
	v202 = F___strchrnul(m, v200, v201)
	mBase = m.M
	v204 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v202))))
	if v204 == v201 {
		goto L53
	} else {
		goto L54
	}
L15:
	;
	v19 = v14 + int32(1)
	v20 = int32(13)
	v21 = F___strchrnul(m, v19, v20)
	mBase = m.M
	v23 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21))))
	if v23 == v20 {
		goto L17
	} else {
		goto L18
	}
L16:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v27 + int32(2)
	v33 = v27 + (v14 ^ int32(-1))
	v35 = v12 + int32(8)
	if base.Ui32(v33+int32(-21)) < base.Ui32(int32(-20)) {
		goto L21
	} else {
		goto L22
	}
L17:
	;
	v27 = v21
	goto L19
L18:
	;
	v27 = int32(0)
	goto L19
L19:
	;
	goto L16
L20:
	;
	v177 = *(*int64)(unsafe.Add(mBase, uint32(v12)+8))
	if v177 != int64(-1) {
		goto L47
	} else {
		goto L48
	}
L21:
	;
	goto L20
L22:
	;
	v48 = int32(1)
	v49 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19))))
	if v33 != v48 {
		goto L25
	} else {
		goto L26
	}
L23:
	;
	goto L21
L24:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v35))) = v151
	goto L23
L25:
	;
	if v49&int32(255) == int32(45) {
		goto L30
	} else {
		goto L31
	}
L26:
	;
	v53 = v49 + int32(-48)
	if base.Ui32(int32(9)) < base.Ui32(v53&int32(255)) {
		goto L21
	} else {
		goto L27
	}
L27:
	;
	if v35 == int32(0) {
		goto L23
	} else {
		goto L28
	}
L28:
	;
	v151 = base.I64_extend_i32_u(v53) & int64(255)
	goto L24
L29:
	;
	if base.Ui32(int32(8)) < base.Ui32((v72+int32(-49))&int32(255)) {
		goto L21
	} else {
		goto L32
	}
L30:
	;
	v1393 = int32(2)
	v69 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19)+1)))
	v71 = v1393
	v72 = v69
	v73 = v14 + v1393
	goto L29
L31:
	;
	v71 = v48
	v72 = v49
	v73 = v19
	goto L29
L32:
	;
	v84 = base.I64_extend_i32_u(v72+int32(-48)) & int64(255)
	if base.Ui32(v33) <= base.Ui32(v71) {
		v127 = v84
		goto L33
	} else {
		goto L34
	}
L33:
	;
	if v49&int32(255) != int32(45) {
		goto L41
	} else {
		goto L42
	}
L34:
	;
	v90 = v71
	v92 = v84
	v94 = v73
	goto L35
L35:
	;
	v96 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v94)+1)))
	if base.Ui32((v96+int32(-58))&int32(255)) < base.Ui32(int32(246)) {
		goto L21
	} else {
		goto L37
	}
L36:
	;
	v127 = v117
	goto L33
L37:
	;
	if base.Ui64(int64(1844674407370955161)) < base.Ui64(v92) {
		goto L21
	} else {
		goto L38
	}
L38:
	;
	v106 = v92 * int64(10)
	v111 = base.I64_extend_i32_u(v96+int32(-48)) & int64(255)
	if base.Ui64(v111^int64(-1)) < base.Ui64(v106) {
		goto L21
	} else {
		goto L39
	}
L39:
	;
	v115 = int32(1)
	v117 = v106 + v111
	v119 = v90 + v115
	if v119 != v33 {
		v90 = v119
		v92 = v117
		v94 = v94 + v115
		goto L35
	} else {
		goto L40
	}
L40:
	;
	goto L36
L41:
	;
	if v127 < int64(0) {
		goto L21
	} else {
		goto L45
	}
L42:
	;
	if base.Ui64(int64(-9223372036854775807-1)) < base.Ui64(v127) {
		goto L21
	} else {
		goto L43
	}
L43:
	;
	if v35 == int32(0) {
		goto L23
	} else {
		goto L44
	}
L44:
	;
	v151 = int64(0) - v127
	goto L24
L45:
	;
	if v35 == int32(0) {
		goto L23
	} else {
		goto L46
	}
L46:
	;
	v151 = v127
	goto L24
L47:
	;
	v188 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v189 = base.I32_wrap_i64(v177)
	v192 = v188 + v189 + int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v192
	v195 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	m.T0[v195].(func(*base.Module, int32, int32, int32, int32, int32))(m, l1, v188, v189, v14, v192-v14)
	mBase = m.M
	v197 = m.ExcPending
	if v197 != 0 {
		goto L49
	} else {
		goto L51
	}
L48:
	;
	v180 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v182 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	m.T0[v182].(func(*base.Module, int32, int32, int32))(m, l1, v14, v180-v14)
	mBase = m.M
	v186 = m.ExcPending
	if v186 != 0 {
		goto L49
	} else {
		goto L50
	}
L49:
	;
	return int32(0)
L50:
	;
	v1383 = int32(0)
	goto L1
L51:
	;
	v1383 = int32(0)
	goto L1
L52:
	;
	v210 = v208 + int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v210
	v216 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	m.T0[v216].(func(*base.Module, int32, int32, int32, int32, int32))(m, l1, v200, v208+(v14^int32(-1)), v14, v210-v14)
	mBase = m.M
	v218 = m.ExcPending
	if v218 != 0 {
		goto L49
	} else {
		goto L56
	}
L53:
	;
	v208 = v202
	goto L55
L54:
	;
	v208 = int32(0)
	goto L55
L55:
	;
	goto L52
L56:
	;
	v1383 = int32(0)
	goto L1
L57:
	;
	v231 = v229 + int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v231
	v237 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	m.T0[v237].(func(*base.Module, int32, int32, int32, int32, int32))(m, l1, v221, v229+(v14^int32(-1)), v14, v231-v14)
	mBase = m.M
	v239 = m.ExcPending
	if v239 != 0 {
		goto L49
	} else {
		goto L61
	}
L58:
	;
	v229 = v223
	goto L60
L59:
	;
	v229 = int32(0)
	goto L60
L60:
	;
	goto L57
L61:
	;
	v1383 = int32(0)
	goto L1
L62:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v250 + int32(2)
	v256 = v250 + (v14 ^ int32(-1))
	v258 = v12 + int32(8)
	if base.Ui32(v256+int32(-21)) < base.Ui32(int32(-20)) {
		goto L67
	} else {
		goto L68
	}
L63:
	;
	v250 = v244
	goto L65
L64:
	;
	v250 = int32(0)
	goto L65
L65:
	;
	goto L62
L66:
	;
	v400 = *(*int64)(unsafe.Add(mBase, uint32(v12)+8))
	v401 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v403 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	m.T0[v403].(func(*base.Module, int32, int64, int32, int32))(m, l1, v400, v14, v401-v14)
	mBase = m.M
	v405 = m.ExcPending
	if v405 != 0 {
		goto L49
	} else {
		goto L93
	}
L67:
	;
	goto L66
L68:
	;
	v271 = int32(1)
	v272 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v242))))
	if v256 != v271 {
		goto L71
	} else {
		goto L72
	}
L69:
	;
	goto L67
L70:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v258))) = v374
	goto L69
L71:
	;
	if v272&int32(255) == int32(45) {
		goto L76
	} else {
		goto L77
	}
L72:
	;
	v276 = v272 + int32(-48)
	if base.Ui32(int32(9)) < base.Ui32(v276&int32(255)) {
		goto L67
	} else {
		goto L73
	}
L73:
	;
	if v258 == int32(0) {
		goto L69
	} else {
		goto L74
	}
L74:
	;
	v374 = base.I64_extend_i32_u(v276) & int64(255)
	goto L70
L75:
	;
	if base.Ui32(int32(8)) < base.Ui32((v295+int32(-49))&int32(255)) {
		goto L67
	} else {
		goto L78
	}
L76:
	;
	v1394 = int32(2)
	v292 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v242)+1)))
	v294 = v1394
	v295 = v292
	v296 = v14 + v1394
	goto L75
L77:
	;
	v294 = v271
	v295 = v272
	v296 = v242
	goto L75
L78:
	;
	v307 = base.I64_extend_i32_u(v295+int32(-48)) & int64(255)
	if base.Ui32(v256) <= base.Ui32(v294) {
		v350 = v307
		goto L79
	} else {
		goto L80
	}
L79:
	;
	if v272&int32(255) != int32(45) {
		goto L87
	} else {
		goto L88
	}
L80:
	;
	v313 = v294
	v315 = v307
	v317 = v296
	goto L81
L81:
	;
	v319 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v317)+1)))
	if base.Ui32((v319+int32(-58))&int32(255)) < base.Ui32(int32(246)) {
		goto L67
	} else {
		goto L83
	}
L82:
	;
	v350 = v340
	goto L79
L83:
	;
	if base.Ui64(int64(1844674407370955161)) < base.Ui64(v315) {
		goto L67
	} else {
		goto L84
	}
L84:
	;
	v329 = v315 * int64(10)
	v334 = base.I64_extend_i32_u(v319+int32(-48)) & int64(255)
	if base.Ui64(v334^int64(-1)) < base.Ui64(v329) {
		goto L67
	} else {
		goto L85
	}
L85:
	;
	v338 = int32(1)
	v340 = v329 + v334
	v342 = v313 + v338
	if v342 != v256 {
		v313 = v342
		v315 = v340
		v317 = v317 + v338
		goto L81
	} else {
		goto L86
	}
L86:
	;
	goto L82
L87:
	;
	if v350 < int64(0) {
		goto L67
	} else {
		goto L91
	}
L88:
	;
	if base.Ui64(int64(-9223372036854775807-1)) < base.Ui64(v350) {
		goto L67
	} else {
		goto L89
	}
L89:
	;
	if v258 == int32(0) {
		goto L69
	} else {
		goto L90
	}
L90:
	;
	v374 = int64(0) - v350
	goto L70
L91:
	;
	if v258 == int32(0) {
		goto L69
	} else {
		goto L92
	}
L92:
	;
	v374 = v350
	goto L70
L93:
	;
	v1383 = int32(0)
	goto L1
L94:
	;
	v419 = v416 + (v14 ^ int32(-1))
	v421 = v12 + int32(8)
	if base.Ui32(v419+int32(-21)) < base.Ui32(int32(-20)) {
		goto L99
	} else {
		goto L100
	}
L95:
	;
	v416 = v410
	goto L97
L96:
	;
	v416 = int32(0)
	goto L97
L97:
	;
	goto L94
L98:
	;
	v564 = v416 + int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v564
	v566 = *(*int64)(unsafe.Add(mBase, uint32(v12)+8))
	if v566 != int64(-1) {
		goto L125
	} else {
		goto L126
	}
L99:
	;
	goto L98
L100:
	;
	v434 = int32(1)
	v435 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v408))))
	if v419 != v434 {
		goto L103
	} else {
		goto L104
	}
L101:
	;
	goto L99
L102:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v421))) = v537
	goto L101
L103:
	;
	if v435&int32(255) == int32(45) {
		goto L108
	} else {
		goto L109
	}
L104:
	;
	v439 = v435 + int32(-48)
	if base.Ui32(int32(9)) < base.Ui32(v439&int32(255)) {
		goto L99
	} else {
		goto L105
	}
L105:
	;
	if v421 == int32(0) {
		goto L101
	} else {
		goto L106
	}
L106:
	;
	v537 = base.I64_extend_i32_u(v439) & int64(255)
	goto L102
L107:
	;
	if base.Ui32(int32(8)) < base.Ui32((v458+int32(-49))&int32(255)) {
		goto L99
	} else {
		goto L110
	}
L108:
	;
	v1395 = int32(2)
	v455 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v408)+1)))
	v457 = v1395
	v458 = v455
	v459 = v14 + v1395
	goto L107
L109:
	;
	v457 = v434
	v458 = v435
	v459 = v408
	goto L107
L110:
	;
	v470 = base.I64_extend_i32_u(v458+int32(-48)) & int64(255)
	if base.Ui32(v419) <= base.Ui32(v457) {
		v513 = v470
		goto L111
	} else {
		goto L112
	}
L111:
	;
	if v435&int32(255) != int32(45) {
		goto L119
	} else {
		goto L120
	}
L112:
	;
	v476 = v457
	v478 = v470
	v480 = v459
	goto L113
L113:
	;
	v482 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v480)+1)))
	if base.Ui32((v482+int32(-58))&int32(255)) < base.Ui32(int32(246)) {
		goto L99
	} else {
		goto L115
	}
L114:
	;
	v513 = v503
	goto L111
L115:
	;
	if base.Ui64(int64(1844674407370955161)) < base.Ui64(v478) {
		goto L99
	} else {
		goto L116
	}
L116:
	;
	v492 = v478 * int64(10)
	v497 = base.I64_extend_i32_u(v482+int32(-48)) & int64(255)
	if base.Ui64(v497^int64(-1)) < base.Ui64(v492) {
		goto L99
	} else {
		goto L117
	}
L117:
	;
	v501 = int32(1)
	v503 = v492 + v497
	v505 = v476 + v501
	if v505 != v419 {
		v476 = v505
		v478 = v503
		v480 = v480 + v501
		goto L113
	} else {
		goto L118
	}
L118:
	;
	goto L114
L119:
	;
	if v513 < int64(0) {
		goto L99
	} else {
		goto L123
	}
L120:
	;
	if base.Ui64(int64(-9223372036854775807-1)) < base.Ui64(v513) {
		goto L99
	} else {
		goto L121
	}
L121:
	;
	if v421 == int32(0) {
		goto L101
	} else {
		goto L122
	}
L122:
	;
	v537 = int64(0) - v513
	goto L102
L123:
	;
	if v421 == int32(0) {
		goto L101
	} else {
		goto L124
	}
L124:
	;
	v537 = v513
	goto L102
L125:
	;
	v575 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	m.T0[v575].(func(*base.Module, int32, int32, int32, int32))(m, l0, l1, base.I32_wrap_i64(v566), v14)
	mBase = m.M
	v577 = m.ExcPending
	if v577 != 0 {
		goto L49
	} else {
		goto L128
	}
L126:
	;
	v570 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	m.T0[v570].(func(*base.Module, int32, int32, int32))(m, l1, v14, v564-v14)
	mBase = m.M
	v572 = m.ExcPending
	if v572 != 0 {
		goto L49
	} else {
		goto L127
	}
L127:
	;
	v1383 = int32(0)
	goto L1
L128:
	;
	v1383 = int32(0)
	goto L1
L129:
	;
	v591 = v588 + (v14 ^ int32(-1))
	v593 = v12 + int32(8)
	if base.Ui32(v591+int32(-21)) < base.Ui32(int32(-20)) {
		goto L134
	} else {
		goto L135
	}
L130:
	;
	v588 = v582
	goto L132
L131:
	;
	v588 = int32(0)
	goto L132
L132:
	;
	goto L129
L133:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v588 + int32(2)
	v738 = *(*int32)(unsafe.Add(mBase, uint32(v12)+8))
	v739 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	m.T0[v739].(func(*base.Module, int32, int32, int32, int32))(m, l0, l1, v738, v14)
	mBase = m.M
	v741 = m.ExcPending
	if v741 != 0 {
		goto L49
	} else {
		goto L160
	}
L134:
	;
	goto L133
L135:
	;
	v606 = int32(1)
	v607 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v580))))
	if v591 != v606 {
		goto L138
	} else {
		goto L139
	}
L136:
	;
	goto L134
L137:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v593))) = v709
	goto L136
L138:
	;
	if v607&int32(255) == int32(45) {
		goto L143
	} else {
		goto L144
	}
L139:
	;
	v611 = v607 + int32(-48)
	if base.Ui32(int32(9)) < base.Ui32(v611&int32(255)) {
		goto L134
	} else {
		goto L140
	}
L140:
	;
	if v593 == int32(0) {
		goto L136
	} else {
		goto L141
	}
L141:
	;
	v709 = base.I64_extend_i32_u(v611) & int64(255)
	goto L137
L142:
	;
	if base.Ui32(int32(8)) < base.Ui32((v630+int32(-49))&int32(255)) {
		goto L134
	} else {
		goto L145
	}
L143:
	;
	v1396 = int32(2)
	v627 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v580)+1)))
	v629 = v1396
	v630 = v627
	v631 = v14 + v1396
	goto L142
L144:
	;
	v629 = v606
	v630 = v607
	v631 = v580
	goto L142
L145:
	;
	v642 = base.I64_extend_i32_u(v630+int32(-48)) & int64(255)
	if base.Ui32(v591) <= base.Ui32(v629) {
		v685 = v642
		goto L146
	} else {
		goto L147
	}
L146:
	;
	if v607&int32(255) != int32(45) {
		goto L154
	} else {
		goto L155
	}
L147:
	;
	v648 = v629
	v650 = v642
	v652 = v631
	goto L148
L148:
	;
	v654 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v652)+1)))
	if base.Ui32((v654+int32(-58))&int32(255)) < base.Ui32(int32(246)) {
		goto L134
	} else {
		goto L150
	}
L149:
	;
	v685 = v675
	goto L146
L150:
	;
	if base.Ui64(int64(1844674407370955161)) < base.Ui64(v650) {
		goto L134
	} else {
		goto L151
	}
L151:
	;
	v664 = v650 * int64(10)
	v669 = base.I64_extend_i32_u(v654+int32(-48)) & int64(255)
	if base.Ui64(v669^int64(-1)) < base.Ui64(v664) {
		goto L134
	} else {
		goto L152
	}
L152:
	;
	v673 = int32(1)
	v675 = v664 + v669
	v677 = v648 + v673
	if v677 != v591 {
		v648 = v677
		v650 = v675
		v652 = v652 + v673
		goto L148
	} else {
		goto L153
	}
L153:
	;
	goto L149
L154:
	;
	if v685 < int64(0) {
		goto L134
	} else {
		goto L158
	}
L155:
	;
	if base.Ui64(int64(-9223372036854775807-1)) < base.Ui64(v685) {
		goto L134
	} else {
		goto L156
	}
L156:
	;
	if v593 == int32(0) {
		goto L136
	} else {
		goto L157
	}
L157:
	;
	v709 = int64(0) - v685
	goto L137
L158:
	;
	if v593 == int32(0) {
		goto L136
	} else {
		goto L159
	}
L159:
	;
	v709 = v685
	goto L137
L160:
	;
	v1383 = int32(0)
	goto L1
L161:
	;
	v755 = v752 + (v14 ^ int32(-1))
	v757 = v12 + int32(8)
	if base.Ui32(v755+int32(-21)) < base.Ui32(int32(-20)) {
		goto L166
	} else {
		goto L167
	}
L162:
	;
	v752 = v746
	goto L164
L163:
	;
	v752 = int32(0)
	goto L164
L164:
	;
	goto L161
L165:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v752 + int32(2)
	v902 = *(*int32)(unsafe.Add(mBase, uint32(v12)+8))
	v903 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	m.T0[v903].(func(*base.Module, int32, int32, int32, int32))(m, l0, l1, v902, v14)
	mBase = m.M
	v905 = m.ExcPending
	if v905 != 0 {
		goto L49
	} else {
		goto L192
	}
L166:
	;
	goto L165
L167:
	;
	v770 = int32(1)
	v771 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v744))))
	if v755 != v770 {
		goto L170
	} else {
		goto L171
	}
L168:
	;
	goto L166
L169:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v757))) = v873
	goto L168
L170:
	;
	if v771&int32(255) == int32(45) {
		goto L175
	} else {
		goto L176
	}
L171:
	;
	v775 = v771 + int32(-48)
	if base.Ui32(int32(9)) < base.Ui32(v775&int32(255)) {
		goto L166
	} else {
		goto L172
	}
L172:
	;
	if v757 == int32(0) {
		goto L168
	} else {
		goto L173
	}
L173:
	;
	v873 = base.I64_extend_i32_u(v775) & int64(255)
	goto L169
L174:
	;
	if base.Ui32(int32(8)) < base.Ui32((v794+int32(-49))&int32(255)) {
		goto L166
	} else {
		goto L177
	}
L175:
	;
	v1397 = int32(2)
	v791 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v744)+1)))
	v793 = v1397
	v794 = v791
	v795 = v14 + v1397
	goto L174
L176:
	;
	v793 = v770
	v794 = v771
	v795 = v744
	goto L174
L177:
	;
	v806 = base.I64_extend_i32_u(v794+int32(-48)) & int64(255)
	if base.Ui32(v755) <= base.Ui32(v793) {
		v849 = v806
		goto L178
	} else {
		goto L179
	}
L178:
	;
	if v771&int32(255) != int32(45) {
		goto L186
	} else {
		goto L187
	}
L179:
	;
	v812 = v793
	v814 = v806
	v816 = v795
	goto L180
L180:
	;
	v818 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v816)+1)))
	if base.Ui32((v818+int32(-58))&int32(255)) < base.Ui32(int32(246)) {
		goto L166
	} else {
		goto L182
	}
L181:
	;
	v849 = v839
	goto L178
L182:
	;
	if base.Ui64(int64(1844674407370955161)) < base.Ui64(v814) {
		goto L166
	} else {
		goto L183
	}
L183:
	;
	v828 = v814 * int64(10)
	v833 = base.I64_extend_i32_u(v818+int32(-48)) & int64(255)
	if base.Ui64(v833^int64(-1)) < base.Ui64(v828) {
		goto L166
	} else {
		goto L184
	}
L184:
	;
	v837 = int32(1)
	v839 = v828 + v833
	v841 = v812 + v837
	if v841 != v755 {
		v812 = v841
		v814 = v839
		v816 = v816 + v837
		goto L180
	} else {
		goto L185
	}
L185:
	;
	goto L181
L186:
	;
	if v849 < int64(0) {
		goto L166
	} else {
		goto L190
	}
L187:
	;
	if base.Ui64(int64(-9223372036854775807-1)) < base.Ui64(v849) {
		goto L166
	} else {
		goto L188
	}
L188:
	;
	if v757 == int32(0) {
		goto L168
	} else {
		goto L189
	}
L189:
	;
	v873 = int64(0) - v849
	goto L169
L190:
	;
	if v757 == int32(0) {
		goto L168
	} else {
		goto L191
	}
L191:
	;
	v873 = v849
	goto L169
L192:
	;
	v1383 = int32(0)
	goto L1
L193:
	;
	v918 = v916 + int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v918
	v920 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+1)))
	v924 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	m.T0[v924].(func(*base.Module, int32, int32, int32, int32))(m, l1, base.B2i32(v920 == int32(116)), v14, v918-v14)
	mBase = m.M
	v926 = m.ExcPending
	if v926 != 0 {
		goto L49
	} else {
		goto L197
	}
L194:
	;
	v916 = v910
	goto L196
L195:
	;
	v916 = int32(0)
	goto L196
L196:
	;
	goto L193
L197:
	;
	v1383 = int32(0)
	goto L1
L198:
	;
	v939 = v937 + int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v939
	v944 = v937 + (v14 ^ int32(-1))
	if base.Ui32(int32(5120)) < base.Ui32(v944) {
		v992 = v939
		v993 = float64(0)
		goto L202
	} else {
		goto L203
	}
L199:
	;
	v937 = v931
	goto L201
L200:
	;
	v937 = int32(0)
	goto L201
L201:
	;
	goto L198
L202:
	;
	v995 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	m.T0[v995].(func(*base.Module, int32, float64, int32, int32))(m, l1, v993, v14, v992-v14)
	mBase = m.M
	v997 = m.ExcPending
	if v997 != 0 {
		goto L49
	} else {
		goto L212
	}
L203:
	;
	v947 = int32(0)
	v951 = m.G0
	v953 = v951 - int32(32)
	m.G0 = v953
	v955 = F___errno_location(m)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v955))) = v947
	v961 = *(*int64)(unsafe.Add(mBase, _consts[648]))
	*(*int64)(unsafe.Add(mBase, uint32(v953+int32(8)))) = v961
	*(*int64)(unsafe.Add(mBase, uint32(v953)+24)) = int64(0)
	v966 = *(*int64)(unsafe.Add(mBase, _consts[649]))
	*(*int64)(unsafe.Add(mBase, uint32(v953))) = v966
	F_ffc_from_chars_double_options(m, v953+int32(16), v929, v929+v944, v953+int32(24), v953)
	mBase = m.M
	v974 = *(*int32)(unsafe.Add(mBase, uint32(v953)+20))
	if v974 == v947 {
		goto L205
	} else {
		goto L206
	}
L204:
	;
	v991 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v992 = v991
	v993 = v987
	goto L202
L205:
	;
	goto L210
L206:
	;
	if v974 == int32(2) {
		goto L207
	} else {
		goto L208
	}
L207:
	;
	v981 = int32(68)
	goto L209
L208:
	;
	v981 = int32(28)
	goto L209
L209:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v955))) = v981
	goto L205
L210:
	;
	v987 = *(*float64)(unsafe.Add(mBase, uint32(v953)+24))
	m.G0 = v953 + int32(32)
	goto L204
L212:
	;
	v1383 = int32(0)
	goto L1
L213:
	;
	v1010 = v1008 + int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v1010
	v1013 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	m.T0[v1013].(func(*base.Module, int32, int32, int32))(m, l1, v14, v1010-v14)
	mBase = m.M
	v1015 = m.ExcPending
	if v1015 != 0 {
		goto L49
	} else {
		goto L217
	}
L214:
	;
	v1008 = v1002
	goto L216
L215:
	;
	v1008 = int32(0)
	goto L216
L216:
	;
	goto L213
L217:
	;
	v1383 = int32(0)
	goto L1
L218:
	;
	v1028 = v1026 + int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v1028
	v1034 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	m.T0[v1034].(func(*base.Module, int32, int32, int32, int32, int32))(m, l1, v1018, v1026+(v14^int32(-1)), v14, v1028-v14)
	mBase = m.M
	v1036 = m.ExcPending
	if v1036 != 0 {
		goto L49
	} else {
		goto L222
	}
L219:
	;
	v1026 = v1020
	goto L221
L220:
	;
	v1026 = int32(0)
	goto L221
L221:
	;
	goto L218
L222:
	;
	v1383 = int32(0)
	goto L1
L223:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v1047 + int32(2)
	v1053 = v1047 + (v14 ^ int32(-1))
	v1055 = v12 + int32(8)
	if base.Ui32(v1053+int32(-21)) < base.Ui32(int32(-20)) {
		goto L228
	} else {
		goto L229
	}
L224:
	;
	v1047 = v1041
	goto L226
L225:
	;
	v1047 = int32(0)
	goto L226
L226:
	;
	goto L223
L227:
	;
	v1197 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1198 = *(*int32)(unsafe.Add(mBase, uint32(v12)+8))
	v1201 = v1197 + v1198 + int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v1201
	v1208 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	m.T0[v1208].(func(*base.Module, int32, int32, int32, int32, int32, int32))(m, l1, v1197, v1197+int32(4), v1198+int32(-4), v14, v1201-v14)
	mBase = m.M
	v1210 = m.ExcPending
	if v1210 != 0 {
		goto L49
	} else {
		goto L254
	}
L228:
	;
	goto L227
L229:
	;
	v1068 = int32(1)
	v1069 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1039))))
	if v1053 != v1068 {
		goto L232
	} else {
		goto L233
	}
L230:
	;
	goto L228
L231:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v1055))) = v1171
	goto L230
L232:
	;
	if v1069&int32(255) == int32(45) {
		goto L237
	} else {
		goto L238
	}
L233:
	;
	v1073 = v1069 + int32(-48)
	if base.Ui32(int32(9)) < base.Ui32(v1073&int32(255)) {
		goto L228
	} else {
		goto L234
	}
L234:
	;
	if v1055 == int32(0) {
		goto L230
	} else {
		goto L235
	}
L235:
	;
	v1171 = base.I64_extend_i32_u(v1073) & int64(255)
	goto L231
L236:
	;
	if base.Ui32(int32(8)) < base.Ui32((v1092+int32(-49))&int32(255)) {
		goto L228
	} else {
		goto L239
	}
L237:
	;
	v1398 = int32(2)
	v1089 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1039)+1)))
	v1091 = v1398
	v1092 = v1089
	v1093 = v14 + v1398
	goto L236
L238:
	;
	v1091 = v1068
	v1092 = v1069
	v1093 = v1039
	goto L236
L239:
	;
	v1104 = base.I64_extend_i32_u(v1092+int32(-48)) & int64(255)
	if base.Ui32(v1053) <= base.Ui32(v1091) {
		v1147 = v1104
		goto L240
	} else {
		goto L241
	}
L240:
	;
	if v1069&int32(255) != int32(45) {
		goto L248
	} else {
		goto L249
	}
L241:
	;
	v1110 = v1091
	v1112 = v1104
	v1114 = v1093
	goto L242
L242:
	;
	v1116 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1114)+1)))
	if base.Ui32((v1116+int32(-58))&int32(255)) < base.Ui32(int32(246)) {
		goto L228
	} else {
		goto L244
	}
L243:
	;
	v1147 = v1137
	goto L240
L244:
	;
	if base.Ui64(int64(1844674407370955161)) < base.Ui64(v1112) {
		goto L228
	} else {
		goto L245
	}
L245:
	;
	v1126 = v1112 * int64(10)
	v1131 = base.I64_extend_i32_u(v1116+int32(-48)) & int64(255)
	if base.Ui64(v1131^int64(-1)) < base.Ui64(v1126) {
		goto L228
	} else {
		goto L246
	}
L246:
	;
	v1135 = int32(1)
	v1137 = v1126 + v1131
	v1139 = v1110 + v1135
	if v1139 != v1053 {
		v1110 = v1139
		v1112 = v1137
		v1114 = v1114 + v1135
		goto L242
	} else {
		goto L247
	}
L247:
	;
	goto L243
L248:
	;
	if v1147 < int64(0) {
		goto L228
	} else {
		goto L252
	}
L249:
	;
	if base.Ui64(int64(-9223372036854775807-1)) < base.Ui64(v1147) {
		goto L228
	} else {
		goto L250
	}
L250:
	;
	if v1055 == int32(0) {
		goto L230
	} else {
		goto L251
	}
L251:
	;
	v1171 = int64(0) - v1147
	goto L231
L252:
	;
	if v1055 == int32(0) {
		goto L230
	} else {
		goto L253
	}
L253:
	;
	v1171 = v1147
	goto L231
L254:
	;
	v1383 = int32(0)
	goto L1
L255:
	;
	v1224 = v1221 + (v14 ^ int32(-1))
	v1226 = v12 + int32(8)
	if base.Ui32(v1224+int32(-21)) < base.Ui32(int32(-20)) {
		goto L260
	} else {
		goto L261
	}
L256:
	;
	v1221 = v1215
	goto L258
L257:
	;
	v1221 = int32(0)
	goto L258
L258:
	;
	goto L255
L259:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v1221 + int32(2)
	v1371 = *(*int32)(unsafe.Add(mBase, uint32(v12)+8))
	v1372 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	m.T0[v1372].(func(*base.Module, int32, int32, int32, int32))(m, l0, l1, v1371, v14)
	mBase = m.M
	v1374 = m.ExcPending
	if v1374 != 0 {
		goto L49
	} else {
		goto L286
	}
L260:
	;
	goto L259
L261:
	;
	v1239 = int32(1)
	v1240 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1213))))
	if v1224 != v1239 {
		goto L264
	} else {
		goto L265
	}
L262:
	;
	goto L260
L263:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v1226))) = v1342
	goto L262
L264:
	;
	if v1240&int32(255) == int32(45) {
		goto L269
	} else {
		goto L270
	}
L265:
	;
	v1244 = v1240 + int32(-48)
	if base.Ui32(int32(9)) < base.Ui32(v1244&int32(255)) {
		goto L260
	} else {
		goto L266
	}
L266:
	;
	if v1226 == int32(0) {
		goto L262
	} else {
		goto L267
	}
L267:
	;
	v1342 = base.I64_extend_i32_u(v1244) & int64(255)
	goto L263
L268:
	;
	if base.Ui32(int32(8)) < base.Ui32((v1263+int32(-49))&int32(255)) {
		goto L260
	} else {
		goto L271
	}
L269:
	;
	v1399 = int32(2)
	v1260 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1213)+1)))
	v1262 = v1399
	v1263 = v1260
	v1264 = v14 + v1399
	goto L268
L270:
	;
	v1262 = v1239
	v1263 = v1240
	v1264 = v1213
	goto L268
L271:
	;
	v1275 = base.I64_extend_i32_u(v1263+int32(-48)) & int64(255)
	if base.Ui32(v1224) <= base.Ui32(v1262) {
		v1318 = v1275
		goto L272
	} else {
		goto L273
	}
L272:
	;
	if v1240&int32(255) != int32(45) {
		goto L280
	} else {
		goto L281
	}
L273:
	;
	v1281 = v1262
	v1283 = v1275
	v1285 = v1264
	goto L274
L274:
	;
	v1287 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1285)+1)))
	if base.Ui32((v1287+int32(-58))&int32(255)) < base.Ui32(int32(246)) {
		goto L260
	} else {
		goto L276
	}
L275:
	;
	v1318 = v1308
	goto L272
L276:
	;
	if base.Ui64(int64(1844674407370955161)) < base.Ui64(v1283) {
		goto L260
	} else {
		goto L277
	}
L277:
	;
	v1297 = v1283 * int64(10)
	v1302 = base.I64_extend_i32_u(v1287+int32(-48)) & int64(255)
	if base.Ui64(v1302^int64(-1)) < base.Ui64(v1297) {
		goto L260
	} else {
		goto L278
	}
L278:
	;
	v1306 = int32(1)
	v1308 = v1297 + v1302
	v1310 = v1281 + v1306
	if v1310 != v1224 {
		v1281 = v1310
		v1283 = v1308
		v1285 = v1285 + v1306
		goto L274
	} else {
		goto L279
	}
L279:
	;
	goto L275
L280:
	;
	if v1318 < int64(0) {
		goto L260
	} else {
		goto L284
	}
L281:
	;
	if base.Ui64(int64(-9223372036854775807-1)) < base.Ui64(v1318) {
		goto L260
	} else {
		goto L282
	}
L282:
	;
	if v1226 == int32(0) {
		goto L262
	} else {
		goto L283
	}
L283:
	;
	v1342 = int64(0) - v1318
	goto L263
L284:
	;
	if v1226 == int32(0) {
		goto L262
	} else {
		goto L285
	}
L285:
	;
	v1342 = v1318
	goto L263
L286:
	;
	v1383 = int32(0)
	goto L1
L287:
	;
	m.T0[v1377].(func(*base.Module, int32))(m, l1)
	mBase = m.M
	v1381 = m.ExcPending
	if v1381 != 0 {
		goto L49
	} else {
		goto L288
	}
L288:
	;
	v1383 = v1376
	goto L1
}
func F_releaseReplyReferences(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v16 int32
	_ = v16
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
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v38 int32
	_ = v38
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v68 int32
	_ = v68
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+180))
	if v8 == int32(0) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+132))
	v21 = v6 + int32(8)
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v19)))
	*(*int32)(unsafe.Add(mBase, uint32(v21)+4)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v21))) = v22
	goto L6
L2:
	;
	v11 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+203)))
	if v11&int32(1) == int32(0) {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+124))
	F_releaseBufReferences(m, v16, v8, l0)
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	return
L5:
	;
	goto L1
L6:
	;
	v27 = v6 + int32(8)
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v27)))
	if v29 == int32(0) {
		goto L9
	} else {
		goto L10
	}
L7:
	;
	m.G0 = v6 + int32(16)
	return
L8:
	;
	if v29 == int32(0) {
		goto L7
	} else {
		goto L11
	}
L9:
	;
	goto L8
L10:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v27)+4))
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v29+base.B2i32(v32 == int32(0))<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v27))) = v38
	goto L9
L11:
	;
	v44 = v29
	goto L12
L12:
	;
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v44)+8))
	v46 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v45)+12)))
	if v46&int32(1) == int32(0) {
		goto L14
	} else {
		goto L15
	}
L13:
	;
	goto L7
L14:
	;
	v57 = v6 + int32(8)
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v57)))
	if v59 == int32(0) {
		goto L18
	} else {
		goto L19
	}
L15:
	;
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v45)+4))
	F_releaseBufReferences(m, v45+int32(13), v53, l0)
	mBase = m.M
	v55 = m.ExcPending
	if v55 != 0 {
		goto L4
	} else {
		goto L16
	}
L16:
	;
	goto L14
L17:
	;
	if v59 != 0 {
		v44 = v59
		goto L12
	} else {
		goto L20
	}
L18:
	;
	goto L17
L19:
	;
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v57)+4))
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v59+base.B2i32(v62 == int32(0))<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v57))) = v68
	goto L18
L20:
	;
	goto L13
}
func F_replyHandlersError(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(v7)+32))
	if v8 == int32(0) {
		return
	} else {
		v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		m.T0[v8].(func(*base.Module, int32, int32, int32))(m, v11, l1, l2)
		mBase = m.M
		v13 = m.ExcPending
		if v13 != 0 {
			return
		} else {
			return
		}
	}
}
func F_replyHandlersNull(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(v5)+8))
	if v6 == int32(0) {
		return
	} else {
		v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		m.T0[v6].(func(*base.Module, int32))(m, v9)
		mBase = m.M
		v11 = m.ExcPending
		if v11 != 0 {
			return
		} else {
			return
		}
	}
}
func F_replyHandlersNullArray(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(v5)+16))
	if v6 == int32(0) {
		return
	} else {
		v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		m.T0[v6].(func(*base.Module, int32))(m, v9)
		mBase = m.M
		v11 = m.ExcPending
		if v11 != 0 {
			return
		} else {
			return
		}
	}
}
func F_replyHandlersSet(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(v6)+76))
	if v7 == int32(0) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	if l2 == int32(0) {
		goto L5
	} else {
		goto L6
	}
L2:
	;
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	m.T0[v7].(func(*base.Module, int32, int32))(m, v10, l2)
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	return
L4:
	;
	goto L1
L5:
	;
	v31 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v31)+80))
	if v32 == int32(0) {
		goto L11
	} else {
		goto L12
	}
L6:
	;
	v20 = int32(0)
	goto L7
L7:
	;
	v21 = F_parseReply(m, l0, l1)
	mBase = m.M
	v22 = m.ExcPending
	if v22 != 0 {
		goto L3
	} else {
		goto L9
	}
L8:
	;
	goto L5
L9:
	;
	v24 = v20 + int32(1)
	if v24 != l2 {
		v20 = v24
		goto L7
	} else {
		goto L10
	}
L10:
	;
	goto L8
L11:
	;
	return
L12:
	;
	v35 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	m.T0[v32].(func(*base.Module, int32))(m, v35)
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		goto L3
	} else {
		goto L13
	}
L13:
	;
	goto L11
}
func F_replyToBlockedClientTimedOut(m *base.Module, l0 int32) {
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
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v18 int64
	_ = v18
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v34 int32
	_ = v34
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
	var v46 int64
	_ = v46
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v62 int32
	_ = v62
	var v67 int32
	_ = v67
	var v69 int64
	_ = v69
	var v70 int32
	_ = v70
	var v71 int64
	_ = v71
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v77 int64
	_ = v77
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v93 int32
	_ = v93
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
	var v105 int64
	_ = v105
	var v108 int32
	_ = v108
	var v111 int32
	_ = v111
	var v114 int32
	_ = v114
	var v121 int32
	_ = v121
	var v125 int32
	_ = v125
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v134 int32
	_ = v134
	var v138 int32
	_ = v138
	var v144 int32
	_ = v144
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v152 int32
	_ = v152
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+116))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(v10)))
	switch v11 + int32(-1) {
	case 0, 3, 4:
		F_addReplyNullArray(m, l0)
		mBase = m.M
		v147 = m.ExcPending
		if v147 != 0 {
			return
		} else {
			v148 = int32(0)
			F_updateStatsOnUnblock(m, l0, v148, v148, v148)
			mBase = m.M
			v152 = m.ExcPending
			if v152 != 0 {
				return
			} else {
				m.G0 = v8 + int32(16)
				return
			}
		}
	case 1:
		v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
		v15 = *(*int32)(unsafe.Add(mBase, uint32(v14)+48))
		if v15 != int32(19) {
			if v15 != int32(20) {
				if v15 != int32(21) {
					v128 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
					*(*int32)(unsafe.Add(mBase, uint32(v8))) = v128
					F__serverPanic_1(m, int32(_a184), int32(291), int32(_a189), v8)
					mBase = m.M
					v134 = m.ExcPending
					if v134 != 0 {
						return
					} else {
						F_abort(m)
						mBase = m.M
						base.Wasm_trap_unreachable()
						for {
						}
					}
				} else {
					v125 = *(*int32)(unsafe.Add(mBase, _consts[77]))
					F_addReplyErrorObject(m, l0, v125)
					mBase = m.M
					v127 = m.ExcPending
					if v127 != 0 {
						return
					} else {
						m.G0 = v8 + int32(16)
						return
					}
				}
			} else {
				F_addReplyArrayLen(m, l0, int32(2))
				mBase = m.M
				v67 = m.ExcPending
				if v67 != 0 {
					return
				} else {
					v69 = *(*int64)(unsafe.Add(mBase, _consts[49]))
					v70 = *(*int32)(unsafe.Add(mBase, uint32(l0)+116))
					v71 = *(*int64)(unsafe.Add(mBase, uint32(v70)+40))
					F_addReplyLongLong(m, l0, base.I64_extend_i32_u(base.B2i32(v71 <= v69)))
					mBase = m.M
					v75 = m.ExcPending
					if v75 != 0 {
						return
					} else {
						v76 = *(*int32)(unsafe.Add(mBase, uint32(l0)+116))
						v77 = *(*int64)(unsafe.Add(mBase, uint32(v76)+40))
						v78 = int32(0)
						v81 = m.G0
						v83 = v81 - int32(16)
						m.G0 = v83
						v86 = *(*int32)(unsafe.Add(mBase, _consts[78]))
						v88 = v83 + int32(8)
						F_listRewind(m, v86, v88)
						mBase = m.M
						v93 = F_listNext(m, v88)
						mBase = m.M
						if v93 == v78 {
							v114 = v78
						} else {
							v98 = v78
							v99 = v93
							for {
								v100 = *(*int32)(unsafe.Add(mBase, uint32(v99)+8))
								v101 = *(*int32)(unsafe.Add(mBase, uint32(v100)+104))
								v102 = *(*int32)(unsafe.Add(mBase, uint32(v101)))
								if v102 != int32(9) {
									v108 = v98
								} else {
									v105 = *(*int64)(unsafe.Add(mBase, uint32(v101)+72))
									v108 = v98 + base.B2i32(v77 <= v105)
								}
								v111 = F_listNext(m, v83+int32(8))
								mBase = m.M
								if v111 != 0 {
									v98 = v108
									v99 = v111
									continue
								} else {
									break
								}
								break
							}
							v114 = v108
						}
						m.G0 = v83 + int32(16)
						F_addReplyLongLong(m, l0, base.I64_extend_i32_s(v114))
						mBase = m.M
						v121 = m.ExcPending
						if v121 != 0 {
							return
						} else {
							m.G0 = v8 + int32(16)
							return
						}
					}
				}
			}
		} else {
			v18 = *(*int64)(unsafe.Add(mBase, uint32(v10)+40))
			v19 = int32(0)
			v22 = m.G0
			v24 = v22 - int32(16)
			m.G0 = v24
			v27 = *(*int32)(unsafe.Add(mBase, _consts[78]))
			v29 = v24 + int32(8)
			F_listRewind(m, v27, v29)
			mBase = m.M
			v34 = F_listNext(m, v29)
			mBase = m.M
			if v34 == v19 {
				v55 = v19
			} else {
				v39 = v19
				v40 = v34
				for {
					v41 = *(*int32)(unsafe.Add(mBase, uint32(v40)+8))
					v42 = *(*int32)(unsafe.Add(mBase, uint32(v41)+104))
					v43 = *(*int32)(unsafe.Add(mBase, uint32(v42)))
					if v43 != int32(9) {
						v49 = v39
					} else {
						v46 = *(*int64)(unsafe.Add(mBase, uint32(v42)+64))
						v49 = v39 + base.B2i32(v18 <= v46)
					}
					v52 = F_listNext(m, v24+int32(8))
					mBase = m.M
					if v52 != 0 {
						v39 = v49
						v40 = v52
						continue
					} else {
						break
					}
					break
				}
				v55 = v49
			}
			m.G0 = v24 + int32(16)
			F_addReplyLongLong(m, l0, base.I64_extend_i32_s(v55))
			mBase = m.M
			v62 = m.ExcPending
			if v62 != 0 {
				return
			} else {
				m.G0 = v8 + int32(16)
				return
			}
		}
	case 2:
		F_moduleBlockedClientTimedOut(m, l0, int32(0))
		mBase = m.M
		v138 = m.ExcPending
		if v138 != 0 {
			return
		} else {
			m.G0 = v8 + int32(16)
			return
		}
	default:
		F__serverPanic_1(m, int32(_a184), int32(296), int32(_a190), int32(0))
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
	}
}
func F_replyToClientsBlockedOnShutdown(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v31 int32
	_ = v31
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v65 int32
	_ = v65
	v3 = m.G0
	v5 = v3 - int32(16)
	m.G0 = v5
	v8 = *(*int32)(unsafe.Add(mBase, _consts[79]))
	if v8 == int32(0) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	m.G0 = v5 + int32(16)
	return
L2:
	;
	v12 = *(*int32)(unsafe.Add(mBase, _consts[26]))
	v14 = v5 + int32(8)
	v15 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
	*(*int32)(unsafe.Add(mBase, uint32(v14)+4)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v14))) = v15
	goto L3
L3:
	;
	v20 = v5 + int32(8)
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v20)))
	if v22 == int32(0) {
		goto L5
	} else {
		goto L6
	}
L4:
	;
	if v22 == int32(0) {
		goto L1
	} else {
		goto L7
	}
L5:
	;
	goto L4
L6:
	;
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v20)+4))
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v22+base.B2i32(v25 == int32(0))<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v20))) = v31
	goto L5
L7:
	;
	v36 = v22
	goto L8
L8:
	;
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v36)+8))
	v38 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v37)+200)))
	if v38&int32(16) == int32(0) {
		goto L10
	} else {
		goto L11
	}
L9:
	;
	goto L1
L10:
	;
	v54 = v5 + int32(8)
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v54)))
	if v56 == int32(0) {
		goto L17
	} else {
		goto L18
	}
L11:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v37)+116))
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v43)))
	if v44 != int32(7) {
		goto L10
	} else {
		goto L12
	}
L12:
	;
	F_addReplyError(m, v37, int32(_a191))
	mBase = m.M
	v49 = m.ExcPending
	if v49 != 0 {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	return
L14:
	;
	F_unblockClient(m, v37, int32(1))
	mBase = m.M
	v52 = m.ExcPending
	if v52 != 0 {
		goto L13
	} else {
		goto L15
	}
L15:
	;
	goto L10
L16:
	;
	if v56 != 0 {
		v36 = v56
		goto L8
	} else {
		goto L19
	}
L17:
	;
	goto L16
L18:
	;
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v54)+4))
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v56+base.B2i32(v59 == int32(0))<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v54))) = v65
	goto L17
L19:
	;
	goto L9
}
