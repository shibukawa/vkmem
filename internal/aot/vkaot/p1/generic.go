package p1

import (
	base "github.com/shibukawa/vkmem/internal/aot/vkaot/base"
	"unsafe"
)

func F_evalGenericCommand(m *base.Module, l0 int32, l1 int32) {
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
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int64
	_ = v18
	var v19 int32
	_ = v19
	var v26 int32
	_ = v26
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v39 int32
	_ = v39
	var v42 int64
	_ = v42
	var v50 int64
	_ = v50
	var v58 int64
	_ = v58
	var v60 int64
	_ = v60
	var v64 int64
	_ = v64
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v83 int32
	_ = v83
	var v91 int32
	_ = v91
	var v100 int32
	_ = v100
	var v103 int32
	_ = v103
	var v106 int32
	_ = v106
	var v115 int32
	_ = v115
	var v118 int32
	_ = v118
	var v121 int32
	_ = v121
	var v123 int32
	_ = v123
	var v127 int32
	_ = v127
	var v134 int32
	_ = v134
	var v137 int32
	_ = v137
	var v140 int32
	_ = v140
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v146 int32
	_ = v146
	var v154 int32
	_ = v154
	var v160 int32
	_ = v160
	var v162 int32
	_ = v162
	var v166 int32
	_ = v166
	var v171 int32
	_ = v171
	var v177 int32
	_ = v177
	var v180 int32
	_ = v180
	var v183 int32
	_ = v183
	var v196 int32
	_ = v196
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v204 int32
	_ = v204
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v215 int32
	_ = v215
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v228 int64
	_ = v228
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v245 int32
	_ = v245
	var v247 int32
	_ = v247
	var v248 int32
	_ = v248
	var v252 int32
	_ = v252
	var v257 int32
	_ = v257
	var v259 int32
	_ = v259
	var v260 int32
	_ = v260
	var v264 int32
	_ = v264
	var v266 int32
	_ = v266
	var v268 int32
	_ = v268
	var v269 int32
	_ = v269
	var v272 int32
	_ = v272
	var v275 int32
	_ = v275
	var v278 int32
	_ = v278
	var v283 int32
	_ = v283
	var v300 int32
	_ = v300
	v7 = m.G0
	v9 = v7 - int32(96)
	m.G0 = v9
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(v11)+8))
	v16 = F_getLongLongFromObjectOrReply(m, l0, v12, v9+int32(40), int32(0))
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		return
	} else {
		if v16 != 0 {
			m.G0 = v9 + int32(96)
			return
		} else {
			v18 = *(*int64)(unsafe.Add(mBase, uint32(v9)+40))
			v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
			if v18 <= base.I64_extend_i32_s(v19+int32(-3)) {
				if int64(-1) < v18 {
					v32 = *(*int32)(unsafe.Add(mBase, uint32(l0)+324))
					if v32 == int32(0) {
						v68 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
						v69 = *(*int32)(unsafe.Add(mBase, uint32(v68)+4))
						v70 = F_objectGetVal(m, v69)
						mBase = m.M
						v72 = v9 + int32(48)
						v76 = m.G0
						v78 = v76 - int32(112)
						m.G0 = v78
						if l1 == int32(0) {
							v123 = int32(0)
							v127 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v70+int32(-1)))))
							switch v127 & int32(7) {
							case 0:
								v144 = int32(base.Ui32(v127) >> (uint(int32(3)) % 32))
							case 1:
								v134 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v70+int32(-3)))))
								v144 = v134
							case 2:
								v137 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v70+int32(-5)))))
								v144 = v137
							case 3:
								v140 = *(*int32)(unsafe.Add(mBase, uint32(v70+int32(-9))))
								v144 = v140
							case 4:
								v143 = *(*int32)(unsafe.Add(mBase, uint32(v70+int32(-17))))
								v144 = v143
							default:
								v144 = v123
							}
							v146 = v78 + int32(20)
							F_SHA1Init(m, v146)
							mBase = m.M
							F_SHA1Update(m, v146, v70, v144)
							mBase = m.M
							F_SHA1Final(m, v78, v146)
							mBase = m.M
							v154 = v123
							for {
								v160 = int32(1)
								v162 = v72 + v154<<(uint(v160)%32)
								v166 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v78+v154))))
								v171 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v166&int32(15))+uint32(_c_F_evalGenericCommand[0]))))
								*(*uint8)(unsafe.Add(mBase, uint32(v162+v160))) = uint8(v171)
								v177 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v166)>>(uint(int32(4))%32)))+uint32(_c_F_evalGenericCommand[0]))))
								*(*uint8)(unsafe.Add(mBase, uint32(v162))) = uint8(v177)
								v180 = v154 + v160
								if v180 != int32(20) {
									v154 = v180
									continue
								} else {
									break
								}
								break
							}
							v183 = int32(0)
							*(*uint8)(unsafe.Add(mBase, uint32(v72)+40)) = uint8(v183)
						} else {
							v83 = int32(0)
							for {
								v91 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v70+v83))))
								if base.Ui32((v91+int32(-65))&int32(255)) < base.Ui32(int32(26)) {
									v100 = v91 + int32(32)
								} else {
									v100 = v91
								}
								*(*uint8)(unsafe.Add(mBase, uint32(v72+v83))) = uint8(v100)
								v103 = v83 | int32(1)
								v106 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v70+v103))))
								if base.Ui32((v106+int32(-65))&int32(255)) < base.Ui32(int32(26)) {
									v115 = v106 + int32(32)
								} else {
									v115 = v106
								}
								*(*uint8)(unsafe.Add(mBase, uint32(v72+v103))) = uint8(v115)
								v118 = v83 + int32(2)
								if v118 != int32(40) {
									v83 = v118
									continue
								} else {
									break
								}
								break
							}
							v121 = int32(0)
							*(*uint8)(unsafe.Add(mBase, uint32(v72)+40)) = uint8(v121)
						}
						m.G0 = v78 + int32(112)
					} else {
						v39 = *(*int32)(unsafe.Add(mBase, uint32(v32)))
						v42 = *(*int64)(unsafe.Add(mBase, uint32(v39+int32(32))))
						*(*int64)(unsafe.Add(mBase, uint32(v9+int32(80)))) = v42
						v50 = *(*int64)(unsafe.Add(mBase, uint32(v39+int32(24))))
						*(*int64)(unsafe.Add(mBase, uint32(v9+int32(72)))) = v50
						v58 = *(*int64)(unsafe.Add(mBase, uint32(v39+int32(16))))
						*(*int64)(unsafe.Add(mBase, uint32(v9+int32(64)))) = v58
						v60 = *(*int64)(unsafe.Add(mBase, uint32(v39)))
						*(*int64)(unsafe.Add(mBase, uint32(v9)+48)) = v60
						v64 = *(*int64)(unsafe.Add(mBase, uint32(v39+int32(8))))
						*(*int64)(unsafe.Add(mBase, uint32(v9)+56)) = v64
						v66 = int32(0)
						*(*uint8)(unsafe.Add(mBase, uint32(v9)+88)) = uint8(v66)
					}
					v196 = *(*int32)(unsafe.Add(mBase, _c_F_evalGenericCommand[1]))
					v199 = F_dictFind(m, v196, v9+int32(48))
					mBase = m.M
					v200 = m.ExcPending
					if v200 != 0 {
						return
					} else {
						if l1 == int32(0) {
							if v199 != 0 {
								v223 = v199
								v224 = *(*int32)(unsafe.Add(mBase, uint32(v223)+8))
								v225 = *(*int32)(unsafe.Add(mBase, uint32(v224)+4))
								v228 = *(*int64)(unsafe.Add(mBase, uint32(v224)+16))
								v229 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
								v230 = *(*int32)(unsafe.Add(mBase, uint32(v229)+48))
								v236 = F_scriptPrepareForRun(m, v9, v225, l0, v9+int32(48), v228, base.B2i32(v230 == int32(294))|base.B2i32(v230 == int32(293)))
								mBase = m.M
								v237 = m.ExcPending
								if v237 != 0 {
									return
								} else {
									if v236 != 0 {
										m.G0 = v9 + int32(96)
										return
									} else {
										v238 = *(*int32)(unsafe.Add(mBase, uint32(v9)+16))
										*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = v238 | int32(128)
										v242 = *(*int32)(unsafe.Add(mBase, uint32(v224)+4))
										v243 = *(*int32)(unsafe.Add(mBase, uint32(v224)))
										v245 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
										v247 = v245 + int32(12)
										v248 = *(*int32)(unsafe.Add(mBase, uint32(v9)+40))
										v252 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
										F_scriptingEngineCallFunction(m, v242, v9, l0, v243, int32(0), v247, v248, v247+v248<<(uint(int32(2))%32), v252-v248+int32(-3))
										mBase = m.M
										v257 = m.ExcPending
										if v257 != 0 {
											return
										} else {
											F_scriptResetRun(m, v9)
											mBase = m.M
											v259 = m.ExcPending
											if v259 != 0 {
												return
											} else {
												v260 = *(*int32)(unsafe.Add(mBase, uint32(v224)+24))
												if v260 == int32(0) {
													m.G0 = v9 + int32(96)
													return
												} else {
													v264 = *(*int32)(unsafe.Add(mBase, _c_F_evalGenericCommand[2]))
													F_listUnlinkNode(m, v264, v260)
													mBase = m.M
													v266 = m.ExcPending
													if v266 != 0 {
														return
													} else {
														v268 = *(*int32)(unsafe.Add(mBase, _c_F_evalGenericCommand[2]))
														v269 = *(*int32)(unsafe.Add(mBase, uint32(v224)+24))
														v272 = *(*int32)(unsafe.Add(mBase, uint32(v268)+20))
														if v272 != 0 {
															v278 = *(*int32)(unsafe.Add(mBase, uint32(v268)+4))
															*(*int32)(unsafe.Add(mBase, uint32(v269)+4)) = int32(0)
															*(*int32)(unsafe.Add(mBase, uint32(v278)+4)) = v269
															*(*int32)(unsafe.Add(mBase, uint32(v268)+4)) = v269
															v283 = v278
														} else {
															*(*int32)(unsafe.Add(mBase, uint32(v268))) = v269
															*(*int32)(unsafe.Add(mBase, uint32(v268)+4)) = v269
															v275 = int32(0)
															*(*int32)(unsafe.Add(mBase, uint32(v269)+4)) = v275
															v283 = v275
														}
														*(*int32)(unsafe.Add(mBase, uint32(v269))) = v283
														*(*int32)(unsafe.Add(mBase, uint32(v268)+20)) = v272 + int32(1)
														m.G0 = v9 + int32(96)
														return
													}
												}
											}
										}
									}
								}
							} else {
								v207 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
								v208 = *(*int32)(unsafe.Add(mBase, uint32(v207)+4))
								*(*int32)(unsafe.Add(mBase, uint32(v9))) = v9 + int32(48)
								v212 = F_evalRegisterNewScript(m, l0, v208, v9)
								mBase = m.M
								v213 = m.ExcPending
								if v213 != 0 {
									return
								} else {
									if v212 != 0 {
										m.G0 = v9 + int32(96)
										return
									} else {
										v215 = *(*int32)(unsafe.Add(mBase, _c_F_evalGenericCommand[1]))
										v218 = F_dictFind(m, v215, v9+int32(48))
										mBase = m.M
										v219 = m.ExcPending
										if v219 != 0 {
											return
										} else {
											if v218 == int32(0) {
												F__serverAssert(m, int32(_a_F_evalGenericCommand_0), int32(_a_F_evalGenericCommand_1), int32(509))
												mBase = m.M
												v300 = m.ExcPending
												if v300 != 0 {
													return
												} else {
													F_abort(m)
													mBase = m.M
													base.Wasm_trap_unreachable()
													for {
													}
												}
											} else {
												v223 = v218
												v224 = *(*int32)(unsafe.Add(mBase, uint32(v223)+8))
												v225 = *(*int32)(unsafe.Add(mBase, uint32(v224)+4))
												v228 = *(*int64)(unsafe.Add(mBase, uint32(v224)+16))
												v229 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
												v230 = *(*int32)(unsafe.Add(mBase, uint32(v229)+48))
												v236 = F_scriptPrepareForRun(m, v9, v225, l0, v9+int32(48), v228, base.B2i32(v230 == int32(294))|base.B2i32(v230 == int32(293)))
												mBase = m.M
												v237 = m.ExcPending
												if v237 != 0 {
													return
												} else {
													if v236 != 0 {
														m.G0 = v9 + int32(96)
														return
													} else {
														v238 = *(*int32)(unsafe.Add(mBase, uint32(v9)+16))
														*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = v238 | int32(128)
														v242 = *(*int32)(unsafe.Add(mBase, uint32(v224)+4))
														v243 = *(*int32)(unsafe.Add(mBase, uint32(v224)))
														v245 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
														v247 = v245 + int32(12)
														v248 = *(*int32)(unsafe.Add(mBase, uint32(v9)+40))
														v252 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
														F_scriptingEngineCallFunction(m, v242, v9, l0, v243, int32(0), v247, v248, v247+v248<<(uint(int32(2))%32), v252-v248+int32(-3))
														mBase = m.M
														v257 = m.ExcPending
														if v257 != 0 {
															return
														} else {
															F_scriptResetRun(m, v9)
															mBase = m.M
															v259 = m.ExcPending
															if v259 != 0 {
																return
															} else {
																v260 = *(*int32)(unsafe.Add(mBase, uint32(v224)+24))
																if v260 == int32(0) {
																	m.G0 = v9 + int32(96)
																	return
																} else {
																	v264 = *(*int32)(unsafe.Add(mBase, _c_F_evalGenericCommand[2]))
																	F_listUnlinkNode(m, v264, v260)
																	mBase = m.M
																	v266 = m.ExcPending
																	if v266 != 0 {
																		return
																	} else {
																		v268 = *(*int32)(unsafe.Add(mBase, _c_F_evalGenericCommand[2]))
																		v269 = *(*int32)(unsafe.Add(mBase, uint32(v224)+24))
																		v272 = *(*int32)(unsafe.Add(mBase, uint32(v268)+20))
																		if v272 != 0 {
																			v278 = *(*int32)(unsafe.Add(mBase, uint32(v268)+4))
																			*(*int32)(unsafe.Add(mBase, uint32(v269)+4)) = int32(0)
																			*(*int32)(unsafe.Add(mBase, uint32(v278)+4)) = v269
																			*(*int32)(unsafe.Add(mBase, uint32(v268)+4)) = v269
																			v283 = v278
																		} else {
																			*(*int32)(unsafe.Add(mBase, uint32(v268))) = v269
																			*(*int32)(unsafe.Add(mBase, uint32(v268)+4)) = v269
																			v275 = int32(0)
																			*(*int32)(unsafe.Add(mBase, uint32(v269)+4)) = v275
																			v283 = v275
																		}
																		*(*int32)(unsafe.Add(mBase, uint32(v269))) = v283
																		*(*int32)(unsafe.Add(mBase, uint32(v268)+20)) = v272 + int32(1)
																		m.G0 = v9 + int32(96)
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
						} else {
							if v199 != 0 {
								if v199 != 0 {
									v223 = v199
									v224 = *(*int32)(unsafe.Add(mBase, uint32(v223)+8))
									v225 = *(*int32)(unsafe.Add(mBase, uint32(v224)+4))
									v228 = *(*int64)(unsafe.Add(mBase, uint32(v224)+16))
									v229 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
									v230 = *(*int32)(unsafe.Add(mBase, uint32(v229)+48))
									v236 = F_scriptPrepareForRun(m, v9, v225, l0, v9+int32(48), v228, base.B2i32(v230 == int32(294))|base.B2i32(v230 == int32(293)))
									mBase = m.M
									v237 = m.ExcPending
									if v237 != 0 {
										return
									} else {
										if v236 != 0 {
											m.G0 = v9 + int32(96)
											return
										} else {
											v238 = *(*int32)(unsafe.Add(mBase, uint32(v9)+16))
											*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = v238 | int32(128)
											v242 = *(*int32)(unsafe.Add(mBase, uint32(v224)+4))
											v243 = *(*int32)(unsafe.Add(mBase, uint32(v224)))
											v245 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
											v247 = v245 + int32(12)
											v248 = *(*int32)(unsafe.Add(mBase, uint32(v9)+40))
											v252 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
											F_scriptingEngineCallFunction(m, v242, v9, l0, v243, int32(0), v247, v248, v247+v248<<(uint(int32(2))%32), v252-v248+int32(-3))
											mBase = m.M
											v257 = m.ExcPending
											if v257 != 0 {
												return
											} else {
												F_scriptResetRun(m, v9)
												mBase = m.M
												v259 = m.ExcPending
												if v259 != 0 {
													return
												} else {
													v260 = *(*int32)(unsafe.Add(mBase, uint32(v224)+24))
													if v260 == int32(0) {
														m.G0 = v9 + int32(96)
														return
													} else {
														v264 = *(*int32)(unsafe.Add(mBase, _c_F_evalGenericCommand[2]))
														F_listUnlinkNode(m, v264, v260)
														mBase = m.M
														v266 = m.ExcPending
														if v266 != 0 {
															return
														} else {
															v268 = *(*int32)(unsafe.Add(mBase, _c_F_evalGenericCommand[2]))
															v269 = *(*int32)(unsafe.Add(mBase, uint32(v224)+24))
															v272 = *(*int32)(unsafe.Add(mBase, uint32(v268)+20))
															if v272 != 0 {
																v278 = *(*int32)(unsafe.Add(mBase, uint32(v268)+4))
																*(*int32)(unsafe.Add(mBase, uint32(v269)+4)) = int32(0)
																*(*int32)(unsafe.Add(mBase, uint32(v278)+4)) = v269
																*(*int32)(unsafe.Add(mBase, uint32(v268)+4)) = v269
																v283 = v278
															} else {
																*(*int32)(unsafe.Add(mBase, uint32(v268))) = v269
																*(*int32)(unsafe.Add(mBase, uint32(v268)+4)) = v269
																v275 = int32(0)
																*(*int32)(unsafe.Add(mBase, uint32(v269)+4)) = v275
																v283 = v275
															}
															*(*int32)(unsafe.Add(mBase, uint32(v269))) = v283
															*(*int32)(unsafe.Add(mBase, uint32(v268)+20)) = v272 + int32(1)
															m.G0 = v9 + int32(96)
															return
														}
													}
												}
											}
										}
									}
								} else {
									v207 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
									v208 = *(*int32)(unsafe.Add(mBase, uint32(v207)+4))
									*(*int32)(unsafe.Add(mBase, uint32(v9))) = v9 + int32(48)
									v212 = F_evalRegisterNewScript(m, l0, v208, v9)
									mBase = m.M
									v213 = m.ExcPending
									if v213 != 0 {
										return
									} else {
										if v212 != 0 {
											m.G0 = v9 + int32(96)
											return
										} else {
											v215 = *(*int32)(unsafe.Add(mBase, _c_F_evalGenericCommand[1]))
											v218 = F_dictFind(m, v215, v9+int32(48))
											mBase = m.M
											v219 = m.ExcPending
											if v219 != 0 {
												return
											} else {
												if v218 == int32(0) {
													F__serverAssert(m, int32(_a_F_evalGenericCommand_0), int32(_a_F_evalGenericCommand_1), int32(509))
													mBase = m.M
													v300 = m.ExcPending
													if v300 != 0 {
														return
													} else {
														F_abort(m)
														mBase = m.M
														base.Wasm_trap_unreachable()
														for {
														}
													}
												} else {
													v223 = v218
													v224 = *(*int32)(unsafe.Add(mBase, uint32(v223)+8))
													v225 = *(*int32)(unsafe.Add(mBase, uint32(v224)+4))
													v228 = *(*int64)(unsafe.Add(mBase, uint32(v224)+16))
													v229 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
													v230 = *(*int32)(unsafe.Add(mBase, uint32(v229)+48))
													v236 = F_scriptPrepareForRun(m, v9, v225, l0, v9+int32(48), v228, base.B2i32(v230 == int32(294))|base.B2i32(v230 == int32(293)))
													mBase = m.M
													v237 = m.ExcPending
													if v237 != 0 {
														return
													} else {
														if v236 != 0 {
															m.G0 = v9 + int32(96)
															return
														} else {
															v238 = *(*int32)(unsafe.Add(mBase, uint32(v9)+16))
															*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = v238 | int32(128)
															v242 = *(*int32)(unsafe.Add(mBase, uint32(v224)+4))
															v243 = *(*int32)(unsafe.Add(mBase, uint32(v224)))
															v245 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
															v247 = v245 + int32(12)
															v248 = *(*int32)(unsafe.Add(mBase, uint32(v9)+40))
															v252 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
															F_scriptingEngineCallFunction(m, v242, v9, l0, v243, int32(0), v247, v248, v247+v248<<(uint(int32(2))%32), v252-v248+int32(-3))
															mBase = m.M
															v257 = m.ExcPending
															if v257 != 0 {
																return
															} else {
																F_scriptResetRun(m, v9)
																mBase = m.M
																v259 = m.ExcPending
																if v259 != 0 {
																	return
																} else {
																	v260 = *(*int32)(unsafe.Add(mBase, uint32(v224)+24))
																	if v260 == int32(0) {
																		m.G0 = v9 + int32(96)
																		return
																	} else {
																		v264 = *(*int32)(unsafe.Add(mBase, _c_F_evalGenericCommand[2]))
																		F_listUnlinkNode(m, v264, v260)
																		mBase = m.M
																		v266 = m.ExcPending
																		if v266 != 0 {
																			return
																		} else {
																			v268 = *(*int32)(unsafe.Add(mBase, _c_F_evalGenericCommand[2]))
																			v269 = *(*int32)(unsafe.Add(mBase, uint32(v224)+24))
																			v272 = *(*int32)(unsafe.Add(mBase, uint32(v268)+20))
																			if v272 != 0 {
																				v278 = *(*int32)(unsafe.Add(mBase, uint32(v268)+4))
																				*(*int32)(unsafe.Add(mBase, uint32(v269)+4)) = int32(0)
																				*(*int32)(unsafe.Add(mBase, uint32(v278)+4)) = v269
																				*(*int32)(unsafe.Add(mBase, uint32(v268)+4)) = v269
																				v283 = v278
																			} else {
																				*(*int32)(unsafe.Add(mBase, uint32(v268))) = v269
																				*(*int32)(unsafe.Add(mBase, uint32(v268)+4)) = v269
																				v275 = int32(0)
																				*(*int32)(unsafe.Add(mBase, uint32(v269)+4)) = v275
																				v283 = v275
																			}
																			*(*int32)(unsafe.Add(mBase, uint32(v269))) = v283
																			*(*int32)(unsafe.Add(mBase, uint32(v268)+20)) = v272 + int32(1)
																			m.G0 = v9 + int32(96)
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
							} else {
								v204 = *(*int32)(unsafe.Add(mBase, _c_F_evalGenericCommand[3]))
								F_addReplyErrorObject(m, l0, v204)
								mBase = m.M
								v206 = m.ExcPending
								if v206 != 0 {
									return
								} else {
									m.G0 = v9 + int32(96)
									return
								}
							}
						}
					}
				} else {
					F_addReplyError(m, l0, int32(_a_F_evalGenericCommand_2))
					mBase = m.M
					v31 = m.ExcPending
					if v31 != 0 {
						return
					} else {
						m.G0 = v9 + int32(96)
						return
					}
				}
			} else {
				F_addReplyError(m, l0, int32(_a_F_evalGenericCommand_3))
				mBase = m.M
				v26 = m.ExcPending
				if v26 != 0 {
					return
				} else {
					m.G0 = v9 + int32(96)
					return
				}
			}
		}
	}
}
func F_genericHgetallCommand(m *base.Module, l0 int32, l1 int32) {
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
	var v21 int32
	_ = v21
	var v27 int32
	_ = v27
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
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v63 int32
	_ = v63
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v83 int32
	_ = v83
	var v86 int32
	_ = v86
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v97 int32
	_ = v97
	var v101 int32
	_ = v101
	var v104 int32
	_ = v104
	var v108 int32
	_ = v108
	var v112 int32
	_ = v112
	var v116 int32
	_ = v116
	var v118 int32
	_ = v118
	var v120 int32
	_ = v120
	v9 = m.G0
	v11 = v9 - int32(688)
	m.G0 = v11
	v13 = int32(3)
	v14 = l1 & v13
	if v14 == v13 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v28)+4))
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v27)))
	v31 = F_lookupKeyReadOrReply(m, l0, v29, v30)
	mBase = m.M
	v32 = m.ExcPending
	if v32 != 0 {
		goto L5
	} else {
		goto L6
	}
L2:
	;
	v21 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+224)))
	v27 = v21<<(uint(int32(2))%32) + int32(_a_F_genericHgetallCommand_0)
	goto L1
L3:
	;
	v27 = int32(_a_F_genericHgetallCommand_1)
	goto L1
L4:
	;
	m.G0 = v11 + int32(688)
	return
L5:
	;
	return
L6:
	;
	if v31 == int32(0) {
		goto L4
	} else {
		goto L7
	}
L7:
	;
	v36 = F_checkType(m, l0, v31, int32(4))
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		goto L5
	} else {
		goto L8
	}
L8:
	;
	if v36 != 0 {
		goto L4
	} else {
		goto L9
	}
L9:
	;
	v38 = F_prepareClientForFutureWrites(m, l0)
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L5
	} else {
		goto L10
	}
L10:
	;
	if v38 == int32(0) {
		goto L4
	} else {
		goto L11
	}
L11:
	;
	v42 = F_addReplyDeferredLen(m, l0)
	mBase = m.M
	v43 = m.ExcPending
	if v43 != 0 {
		goto L5
	} else {
		goto L12
	}
L12:
	;
	F_hashTypeInitIterator(m, v31, v11+int32(8))
	mBase = m.M
	v47 = m.ExcPending
	if v47 != 0 {
		goto L5
	} else {
		goto L13
	}
L13:
	;
	v48 = int32(0)
	v51 = F_hashTypeNext(m, v11+int32(8))
	mBase = m.M
	v52 = m.ExcPending
	if v52 != 0 {
		goto L5
	} else {
		goto L15
	}
L14:
	;
	v101 = *(*int32)(unsafe.Add(mBase, uint32(v11)+12))
	if v101 != int32(2) {
		goto L27
	} else {
		goto L28
	}
L15:
	;
	if v51 == int32(-1) {
		v97 = v48
		goto L14
	} else {
		goto L16
	}
L16:
	;
	v63 = v48
	goto L17
L17:
	;
	if l1&int32(1) == int32(0) {
		v76 = v63
		goto L19
	} else {
		goto L20
	}
L18:
	;
	v97 = v86
	goto L14
L19:
	;
	if l1&int32(2) == int32(0) {
		v86 = v76
		goto L22
	} else {
		goto L23
	}
L20:
	;
	F_addHashIteratorCursorToReply(m, v38, v11+int32(8), int32(1))
	mBase = m.M
	v73 = m.ExcPending
	if v73 != 0 {
		goto L5
	} else {
		goto L21
	}
L21:
	;
	v76 = v63 + int32(1)
	goto L19
L22:
	;
	v89 = F_hashTypeNext(m, v11+int32(8))
	mBase = m.M
	v90 = m.ExcPending
	if v90 != 0 {
		goto L5
	} else {
		goto L25
	}
L23:
	;
	F_addHashIteratorCursorToReply(m, v38, v11+int32(8), int32(2))
	mBase = m.M
	v83 = m.ExcPending
	if v83 != 0 {
		goto L5
	} else {
		goto L24
	}
L24:
	;
	v86 = v76 + int32(1)
	goto L22
L25:
	;
	if v89 != int32(-1) {
		v63 = v86
		goto L17
	} else {
		goto L26
	}
L26:
	;
	goto L18
L27:
	;
	if v14 != int32(3) {
		goto L33
	} else {
		goto L34
	}
L28:
	;
	v104 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+16)))
	if v104 != 0 {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	F_vsetResetIterator(m, v11+int32(80))
	mBase = m.M
	v112 = m.ExcPending
	if v112 != 0 {
		goto L5
	} else {
		goto L32
	}
L30:
	;
	F_hashtableCleanupIterator(m, v11+int32(32))
	mBase = m.M
	v108 = m.ExcPending
	if v108 != 0 {
		goto L5
	} else {
		goto L31
	}
L31:
	;
	goto L27
L32:
	;
	goto L27
L33:
	;
	F_setDeferredArrayLen(m, l0, v42, v97)
	mBase = m.M
	v120 = m.ExcPending
	if v120 != 0 {
		goto L5
	} else {
		goto L36
	}
L34:
	;
	v116 = base.I32_div_s(v97, int32(2))
	F_setDeferredMapLen(m, l0, v42, v116)
	mBase = m.M
	v118 = m.ExcPending
	if v118 != 0 {
		goto L5
	} else {
		goto L35
	}
L35:
	;
	goto L4
L36:
	;
	goto L4
}
func F_genericZrangebylexCommand(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) {
	mBase := m.M
	_ = mBase
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v52 int32
	_ = v52
	var v70 int32
	_ = v70
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v100 int32
	_ = v100
	var v102 int32
	_ = v102
	var v107 int32
	_ = v107
	var v114 int32
	_ = v114
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v130 int32
	_ = v130
	var v132 int32
	_ = v132
	var v134 int32
	_ = v134
	var v140 int64
	_ = v140
	var v145 int64
	_ = v145
	var v153 int32
	_ = v153
	var v160 int32
	_ = v160
	var v166 float64
	_ = v166
	var v170 int64
	_ = v170
	var v172 float64
	_ = v172
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v188 int64
	_ = v188
	var v189 int32
	_ = v189
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v195 int32
	_ = v195
	var v203 int32
	_ = v203
	var v209 int32
	_ = v209
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v220 int32
	_ = v220
	var v222 int32
	_ = v222
	var v227 int32
	_ = v227
	var v233 int32
	_ = v233
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v251 int32
	_ = v251
	var v253 int32
	_ = v253
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v262 int32
	_ = v262
	var v265 int32
	_ = v265
	var v272 int32
	_ = v272
	var v275 int32
	_ = v275
	var v280 int32
	_ = v280
	var v282 int32
	_ = v282
	var v289 int32
	_ = v289
	var v296 int32
	_ = v296
	var v299 int32
	_ = v299
	var v302 int32
	_ = v302
	var v305 int32
	_ = v305
	var v306 int32
	_ = v306
	var v309 int32
	_ = v309
	var v316 int32
	_ = v316
	var v319 int32
	_ = v319
	var v322 int32
	_ = v322
	var v325 int32
	_ = v325
	var v326 int32
	_ = v326
	var v327 int32
	_ = v327
	var v328 int32
	_ = v328
	var v329 int32
	_ = v329
	var v332 int32
	_ = v332
	var v334 int32
	_ = v334
	var v335 int32
	_ = v335
	var v340 int32
	_ = v340
	var v343 int32
	_ = v343
	var v347 int32
	_ = v347
	var v354 int32
	_ = v354
	var v361 int32
	_ = v361
	var v364 int32
	_ = v364
	var v367 int32
	_ = v367
	var v370 int32
	_ = v370
	var v371 int32
	_ = v371
	var v374 int32
	_ = v374
	var v381 int32
	_ = v381
	var v384 int32
	_ = v384
	var v387 int32
	_ = v387
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
	var v397 int32
	_ = v397
	var v402 int32
	_ = v402
	var v405 int32
	_ = v405
	var v409 int32
	_ = v409
	var v416 int32
	_ = v416
	var v423 int32
	_ = v423
	var v426 int32
	_ = v426
	var v429 int32
	_ = v429
	var v432 int32
	_ = v432
	var v433 int32
	_ = v433
	var v436 int32
	_ = v436
	var v443 int32
	_ = v443
	var v446 int32
	_ = v446
	var v449 int32
	_ = v449
	var v452 int32
	_ = v452
	var v453 int32
	_ = v453
	var v454 int32
	_ = v454
	var v455 int32
	_ = v455
	var v456 int32
	_ = v456
	var v459 int32
	_ = v459
	var v462 int32
	_ = v462
	var v471 int32
	_ = v471
	var v478 int32
	_ = v478
	var v481 int32
	_ = v481
	var v484 int32
	_ = v484
	var v487 int32
	_ = v487
	var v488 int32
	_ = v488
	var v492 int32
	_ = v492
	var v493 float64
	_ = v493
	var v494 int32
	_ = v494
	var v496 int32
	_ = v496
	var v498 int32
	_ = v498
	var v508 int32
	_ = v508
	var v514 int32
	_ = v514
	var v516 int32
	_ = v516
	var v524 int32
	_ = v524
	var v531 int32
	_ = v531
	v16 = m.G0
	v18 = v16 - int32(48)
	m.G0 = v18
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	m.T0[v21].(func(*base.Module, int32, int32))(m, l0, int32(-1))
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	switch int32(base.Ui32(v24)>>(uint(int32(4))%32))&int32(15) + int32(-7) {
	case 0:
		goto L6
	default:
		goto L3
	case 4:
		goto L7
	}
L3:
	;
	F__serverPanic_1(m, int32(_a_F_genericZrangebylexCommand_0), int32(3585), int32(_a_F_genericZrangebylexCommand_1), int32(0))
	mBase = m.M
	v531 = m.ExcPending
	if v531 != 0 {
		goto L1
	} else {
		goto L173
	}
L4:
	;
	F__serverAssert(m, int32(_a_F_genericZrangebylexCommand_2), int32(_a_F_genericZrangebylexCommand_0), int32(857))
	mBase = m.M
	v524 = m.ExcPending
	if v524 != 0 {
		goto L1
	} else {
		goto L172
	}
L5:
	;
	v514 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	m.T0[v514].(func(*base.Module, int32, int32))(m, l0, v508)
	mBase = m.M
	v516 = m.ExcPending
	if v516 != 0 {
		goto L1
	} else {
		goto L171
	}
L6:
	;
	v215 = int32(0)
	v216 = F_objectGetVal(m, l2)
	mBase = m.M
	v217 = *(*int32)(unsafe.Add(mBase, uint32(v216)+4))
	if l6 != 0 {
		goto L64
	} else {
		goto L65
	}
L7:
	;
	v31 = F_objectGetVal(m, l2)
	mBase = m.M
	if l6 == int32(0) {
		goto L9
	} else {
		goto L10
	}
L8:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+28)) = v38
	v40 = int32(0)
	if v38 == v40 {
		v508 = v40
		goto L5
	} else {
		goto L13
	}
L9:
	;
	v36 = F_zzlFirstInLexRange(m, v31, l1)
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		goto L1
	} else {
		goto L12
	}
L10:
	;
	v34 = F_zzlLastInLexRange(m, v31, l1)
	mBase = m.M
	v35 = m.ExcPending
	if v35 != 0 {
		goto L1
	} else {
		goto L11
	}
L11:
	;
	v38 = v34
	goto L8
L12:
	;
	v38 = v36
	goto L8
L13:
	;
	v43 = F_lpNext(m, v31, v38)
	mBase = m.M
	v44 = m.ExcPending
	if v44 != 0 {
		goto L1
	} else {
		goto L14
	}
L14:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+24)) = v43
	if l4 == int32(0) {
		v84 = v38
		goto L15
	} else {
		goto L16
	}
L15:
	;
	v100 = v84
	v102 = l5
	v107 = int32(0)
	goto L26
L16:
	;
	v52 = l4
	goto L17
L17:
	;
	if l6 == int32(0) {
		goto L20
	} else {
		goto L21
	}
L18:
	;
	v84 = v77
	goto L15
L19:
	;
	v77 = *(*int32)(unsafe.Add(mBase, uint32(v18)+28))
	if v77 == int32(0) {
		v508 = v40
		goto L5
	} else {
		goto L24
	}
L20:
	;
	F_zzlNext(m, v31, v18+int32(28), v18+int32(24))
	mBase = m.M
	v76 = m.ExcPending
	if v76 != 0 {
		goto L1
	} else {
		goto L23
	}
L21:
	;
	F_zzlPrev(m, v31, v18+int32(28), v18+int32(24))
	mBase = m.M
	v70 = m.ExcPending
	if v70 != 0 {
		goto L1
	} else {
		goto L22
	}
L22:
	;
	goto L19
L23:
	;
	goto L19
L24:
	;
	v81 = v52 + int32(-1)
	if v81 != 0 {
		v52 = v81
		goto L17
	} else {
		goto L25
	}
L25:
	;
	goto L18
L26:
	;
	if v102 != 0 {
		goto L28
	} else {
		goto L29
	}
L28:
	;
	if l3 != 0 {
		goto L31
	} else {
		goto L32
	}
L29:
	;
	v508 = l5
	goto L5
L30:
	;
	if l6 == int32(0) {
		goto L46
	} else {
		goto L47
	}
L31:
	;
	v114 = *(*int32)(unsafe.Add(mBase, uint32(v18)+24))
	if v114 == int32(0) {
		goto L4
	} else {
		goto L33
	}
L32:
	;
	v172 = float64(0)
	goto L30
L33:
	;
	v121 = F_lpGetValue(m, v114, v18+int32(44), v18+int32(32))
	mBase = m.M
	v122 = m.ExcPending
	if v122 != 0 {
		goto L1
	} else {
		goto L35
	}
L34:
	;
	v170 = *(*int64)(unsafe.Add(mBase, uint32(v18)+32))
	v172 = base.F64_convert_i64_s(v170)
	goto L30
L35:
	;
	if v121 == int32(0) {
		goto L34
	} else {
		goto L36
	}
L36:
	;
	v125 = *(*int32)(unsafe.Add(mBase, uint32(v18)+44))
	v126 = int32(0)
	v130 = m.G0
	v132 = v130 - int32(32)
	m.G0 = v132
	v134 = F___errno_location(m)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v134))) = v126
	v140 = *(*int64)(unsafe.Add(mBase, _c_F_genericZrangebylexCommand[0]))
	*(*int64)(unsafe.Add(mBase, uint32(v132+int32(8)))) = v140
	*(*int64)(unsafe.Add(mBase, uint32(v132)+24)) = int64(0)
	v145 = *(*int64)(unsafe.Add(mBase, _c_F_genericZrangebylexCommand[1]))
	*(*int64)(unsafe.Add(mBase, uint32(v132))) = v145
	F_ffc_from_chars_double_options(m, v132+int32(16), v121, v121+v125, v132+int32(24), v132)
	mBase = m.M
	v153 = *(*int32)(unsafe.Add(mBase, uint32(v132)+20))
	if v153 == v126 {
		goto L38
	} else {
		goto L39
	}
L37:
	;
	v172 = v166
	goto L30
L38:
	;
	goto L43
L39:
	;
	if v153 == int32(2) {
		goto L40
	} else {
		goto L41
	}
L40:
	;
	v160 = int32(68)
	goto L42
L41:
	;
	v160 = int32(28)
	goto L42
L42:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v134))) = v160
	goto L38
L43:
	;
	v166 = *(*float64)(unsafe.Add(mBase, uint32(v132)+24))
	m.G0 = v132 + int32(32)
	goto L37
L45:
	;
	v186 = F_lpGetValue(m, v100, v18+int32(20), v18+int32(8))
	mBase = m.M
	v187 = m.ExcPending
	if v187 != 0 {
		goto L1
	} else {
		goto L54
	}
L46:
	;
	v178 = F_zzlLexValueLteMax(m, v100, l1)
	mBase = m.M
	v179 = m.ExcPending
	if v179 != 0 {
		goto L1
	} else {
		goto L50
	}
L47:
	;
	v176 = F_zzlLexValueGteMin(m, v100, l1)
	mBase = m.M
	v177 = m.ExcPending
	if v177 != 0 {
		goto L1
	} else {
		goto L48
	}
L48:
	;
	if v176 != 0 {
		goto L45
	} else {
		goto L49
	}
L49:
	;
	v508 = v107
	goto L5
L50:
	;
	if v178 == int32(0) {
		v508 = v107
		goto L5
	} else {
		goto L51
	}
L51:
	;
	goto L45
L52:
	;
	if l6 == int32(0) {
		goto L59
	} else {
		goto L60
	}
L53:
	;
	v192 = *(*int32)(unsafe.Add(mBase, uint32(v18)+20))
	v193 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	m.T0[v193].(func(*base.Module, int32, int32, int32, float64))(m, l0, v186, v192, v172)
	mBase = m.M
	v195 = m.ExcPending
	if v195 != 0 {
		goto L1
	} else {
		goto L57
	}
L54:
	;
	if v186 != 0 {
		goto L53
	} else {
		goto L55
	}
L55:
	;
	v188 = *(*int64)(unsafe.Add(mBase, uint32(v18)+8))
	v189 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	m.T0[v189].(func(*base.Module, int32, int64, float64))(m, l0, v188, v172)
	mBase = m.M
	v191 = m.ExcPending
	if v191 != 0 {
		goto L1
	} else {
		goto L56
	}
L56:
	;
	goto L52
L57:
	;
	goto L52
L58:
	;
	v213 = v107 + int32(1)
	v214 = *(*int32)(unsafe.Add(mBase, uint32(v18)+28))
	if v214 != 0 {
		v100 = v214
		v102 = v102 + int32(-1)
		v107 = v213
		goto L26
	} else {
		goto L63
	}
L59:
	;
	F_zzlNext(m, v31, v18+int32(28), v18+int32(24))
	mBase = m.M
	v209 = m.ExcPending
	if v209 != 0 {
		goto L1
	} else {
		goto L62
	}
L60:
	;
	F_zzlPrev(m, v31, v18+int32(28), v18+int32(24))
	mBase = m.M
	v203 = m.ExcPending
	if v203 != 0 {
		goto L1
	} else {
		goto L61
	}
L61:
	;
	goto L58
L62:
	;
	goto L58
L63:
	;
	v508 = v213
	goto L5
L64:
	;
	v220 = int32(-1)
	goto L66
L65:
	;
	v220 = v215
	goto L66
L66:
	;
	v222 = F_zslNthInLexRange(m, v217, l1, l4^v220)
	mBase = m.M
	if v222 == int32(0) {
		v508 = v215
		goto L5
	} else {
		goto L67
	}
L67:
	;
	if l6 != 0 {
		goto L68
	} else {
		goto L69
	}
L68:
	;
	v227 = int32(8)
	goto L70
L69:
	;
	v227 = int32(12)
	goto L70
L70:
	;
	v233 = v222
	v237 = l5
	v238 = int32(0)
	goto L71
L71:
	;
	if v237 != 0 {
		goto L73
	} else {
		goto L74
	}
L72:
	;
	v508 = v492
	goto L5
L73:
	;
	v245 = v233 + int32(16)
	v246 = *(*int32)(unsafe.Add(mBase, uint32(v245)))
	v249 = v245 + v246<<(uint(int32(3))%32)
	v250 = int32(*(*int8)(unsafe.Add(mBase, uint32(v249))))
	v251 = v249 + v250
	v253 = v251 + int32(1)
	if l6 == int32(0) {
		goto L76
	} else {
		goto L77
	}
L74:
	;
	v508 = l5
	goto L5
L75:
	;
	v471 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v251))))
	switch v471 & int32(7) {
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
		v488 = int32(0)
		goto L163
	}
L76:
	;
	v334 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v335 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	if v335 == int32(0) {
		goto L112
	} else {
		goto L113
	}
L77:
	;
	v256 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v257 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	if v257 == int32(0) {
		goto L79
	} else {
		goto L80
	}
L78:
	;
	v282 = int32(0)
	v289 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v251+v282))))
	switch v289 & int32(7) {
	case 0:
		goto L97
	case 1:
		goto L96
	case 2:
		goto L95
	case 3:
		goto L94
	case 4:
		goto L93
	default:
		v306 = v282
		goto L92
	}
L79:
	;
	if v253 == v256 {
		goto L75
	} else {
		goto L86
	}
L80:
	;
	if v253 == v256 {
		v508 = v238
		goto L5
	} else {
		goto L81
	}
L81:
	;
	v262 = *(*int32)(unsafe.Add(mBase, _c_F_genericZrangebylexCommand[2]))
	if v253 == v262 {
		v508 = v238
		goto L5
	} else {
		goto L82
	}
L82:
	;
	v265 = *(*int32)(unsafe.Add(mBase, _c_F_genericZrangebylexCommand[3]))
	if v256 == v265 {
		v508 = v238
		goto L5
	} else {
		goto L83
	}
L83:
	;
	if v256 == v262 {
		goto L75
	} else {
		goto L84
	}
L84:
	;
	if v253 == v265 {
		goto L75
	} else {
		goto L85
	}
L85:
	;
	v280 = int32(0)
	goto L78
L86:
	;
	v272 = *(*int32)(unsafe.Add(mBase, _c_F_genericZrangebylexCommand[2]))
	if v253 == v272 {
		v508 = v238
		goto L5
	} else {
		goto L87
	}
L87:
	;
	v275 = *(*int32)(unsafe.Add(mBase, _c_F_genericZrangebylexCommand[3]))
	if v256 == v275 {
		v508 = v238
		goto L5
	} else {
		goto L88
	}
L88:
	;
	if v256 == v272 {
		goto L75
	} else {
		goto L89
	}
L89:
	;
	if v253 == v275 {
		goto L75
	} else {
		goto L90
	}
L90:
	;
	v280 = int32(-1)
	goto L78
L91:
	;
	if v280 < v332 {
		goto L75
	} else {
		goto L110
	}
L92:
	;
	v309 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v256+int32(-1)))))
	switch v309 & int32(7) {
	case 0:
		goto L103
	case 1:
		goto L102
	case 2:
		goto L101
	case 3:
		goto L100
	case 4:
		goto L99
	default:
		v326 = v282
		goto L98
	}
L93:
	;
	v305 = *(*int32)(unsafe.Add(mBase, uint32(v251+int32(-16))))
	v306 = v305
	goto L92
L94:
	;
	v302 = *(*int32)(unsafe.Add(mBase, uint32(v251+int32(-8))))
	v306 = v302
	goto L92
L95:
	;
	v299 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v251+int32(-4)))))
	v306 = v299
	goto L92
L96:
	;
	v296 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v251+int32(-2)))))
	v306 = v296
	goto L92
L97:
	;
	v306 = int32(base.Ui32(v289) >> (uint(int32(3)) % 32))
	goto L92
L98:
	;
	v327 = base.B2i32(base.Ui32(v306) < base.Ui32(v326))
	if base.Ui32(v306) < base.Ui32(v326) {
		goto L104
	} else {
		goto L105
	}
L99:
	;
	v325 = *(*int32)(unsafe.Add(mBase, uint32(v256+int32(-17))))
	v326 = v325
	goto L98
L100:
	;
	v322 = *(*int32)(unsafe.Add(mBase, uint32(v256+int32(-9))))
	v326 = v322
	goto L98
L101:
	;
	v319 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v256+int32(-5)))))
	v326 = v319
	goto L98
L102:
	;
	v316 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v256+int32(-3)))))
	v326 = v316
	goto L98
L103:
	;
	v326 = int32(base.Ui32(v309) >> (uint(int32(3)) % 32))
	goto L98
L104:
	;
	v328 = v306
	goto L106
L105:
	;
	v328 = v326
	goto L106
L106:
	;
	v329 = F_memcmp(m, v253, v256, v328)
	mBase = m.M
	if v329 != 0 {
		goto L107
	} else {
		goto L108
	}
L107:
	;
	v332 = v329
	goto L109
L108:
	;
	v332 = base.B2i32(base.Ui32(v326) < base.Ui32(v306)) - v327
	goto L109
L109:
	;
	goto L91
L110:
	;
	v508 = v238
	goto L5
L111:
	;
	if v462 == int32(0) {
		v508 = v238
		goto L5
	} else {
		goto L162
	}
L112:
	;
	if v253 == v334 {
		goto L75
	} else {
		goto L138
	}
L113:
	;
	if v253 == v334 {
		v508 = v238
		goto L5
	} else {
		goto L114
	}
L114:
	;
	v340 = *(*int32)(unsafe.Add(mBase, _c_F_genericZrangebylexCommand[2]))
	if v253 == v340 {
		goto L75
	} else {
		goto L115
	}
L115:
	;
	v343 = *(*int32)(unsafe.Add(mBase, _c_F_genericZrangebylexCommand[3]))
	if v334 == v343 {
		goto L75
	} else {
		goto L116
	}
L116:
	;
	if v334 == v340 {
		v508 = v238
		goto L5
	} else {
		goto L117
	}
L117:
	;
	if v253 == v343 {
		v508 = v238
		goto L5
	} else {
		goto L118
	}
L118:
	;
	v347 = int32(0)
	v354 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v251+v347))))
	switch v354 & int32(7) {
	case 0:
		goto L125
	case 1:
		goto L124
	case 2:
		goto L123
	case 3:
		goto L122
	case 4:
		goto L121
	default:
		v371 = v347
		goto L120
	}
L119:
	;
	v462 = int32(base.Ui32(v397) >> (uint(int32(31)) % 32))
	goto L111
L120:
	;
	v374 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v334+int32(-1)))))
	switch v374 & int32(7) {
	case 0:
		goto L131
	case 1:
		goto L130
	case 2:
		goto L129
	case 3:
		goto L128
	case 4:
		goto L127
	default:
		v391 = v347
		goto L126
	}
L121:
	;
	v370 = *(*int32)(unsafe.Add(mBase, uint32(v251+int32(-16))))
	v371 = v370
	goto L120
L122:
	;
	v367 = *(*int32)(unsafe.Add(mBase, uint32(v251+int32(-8))))
	v371 = v367
	goto L120
L123:
	;
	v364 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v251+int32(-4)))))
	v371 = v364
	goto L120
L124:
	;
	v361 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v251+int32(-2)))))
	v371 = v361
	goto L120
L125:
	;
	v371 = int32(base.Ui32(v354) >> (uint(int32(3)) % 32))
	goto L120
L126:
	;
	v392 = base.B2i32(base.Ui32(v371) < base.Ui32(v391))
	if base.Ui32(v371) < base.Ui32(v391) {
		goto L132
	} else {
		goto L133
	}
L127:
	;
	v390 = *(*int32)(unsafe.Add(mBase, uint32(v334+int32(-17))))
	v391 = v390
	goto L126
L128:
	;
	v387 = *(*int32)(unsafe.Add(mBase, uint32(v334+int32(-9))))
	v391 = v387
	goto L126
L129:
	;
	v384 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v334+int32(-5)))))
	v391 = v384
	goto L126
L130:
	;
	v381 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v334+int32(-3)))))
	v391 = v381
	goto L126
L131:
	;
	v391 = int32(base.Ui32(v374) >> (uint(int32(3)) % 32))
	goto L126
L132:
	;
	v393 = v371
	goto L134
L133:
	;
	v393 = v391
	goto L134
L134:
	;
	v394 = F_memcmp(m, v253, v334, v393)
	mBase = m.M
	if v394 != 0 {
		goto L135
	} else {
		goto L136
	}
L135:
	;
	v397 = v394
	goto L137
L136:
	;
	v397 = base.B2i32(base.Ui32(v391) < base.Ui32(v371)) - v392
	goto L137
L137:
	;
	goto L119
L138:
	;
	v402 = *(*int32)(unsafe.Add(mBase, _c_F_genericZrangebylexCommand[2]))
	if v253 == v402 {
		goto L75
	} else {
		goto L139
	}
L139:
	;
	v405 = *(*int32)(unsafe.Add(mBase, _c_F_genericZrangebylexCommand[3]))
	if v334 == v405 {
		goto L75
	} else {
		goto L140
	}
L140:
	;
	if v334 == v402 {
		v508 = v238
		goto L5
	} else {
		goto L141
	}
L141:
	;
	if v253 == v405 {
		v508 = v238
		goto L5
	} else {
		goto L142
	}
L142:
	;
	v409 = int32(0)
	v416 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v251+v409))))
	switch v416 & int32(7) {
	case 0:
		goto L149
	case 1:
		goto L148
	case 2:
		goto L147
	case 3:
		goto L146
	case 4:
		goto L145
	default:
		v433 = v409
		goto L144
	}
L143:
	;
	v462 = base.B2i32(v459 < int32(1))
	goto L111
L144:
	;
	v436 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v334+int32(-1)))))
	switch v436 & int32(7) {
	case 0:
		goto L155
	case 1:
		goto L154
	case 2:
		goto L153
	case 3:
		goto L152
	case 4:
		goto L151
	default:
		v453 = v409
		goto L150
	}
L145:
	;
	v432 = *(*int32)(unsafe.Add(mBase, uint32(v251+int32(-16))))
	v433 = v432
	goto L144
L146:
	;
	v429 = *(*int32)(unsafe.Add(mBase, uint32(v251+int32(-8))))
	v433 = v429
	goto L144
L147:
	;
	v426 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v251+int32(-4)))))
	v433 = v426
	goto L144
L148:
	;
	v423 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v251+int32(-2)))))
	v433 = v423
	goto L144
L149:
	;
	v433 = int32(base.Ui32(v416) >> (uint(int32(3)) % 32))
	goto L144
L150:
	;
	v454 = base.B2i32(base.Ui32(v433) < base.Ui32(v453))
	if base.Ui32(v433) < base.Ui32(v453) {
		goto L156
	} else {
		goto L157
	}
L151:
	;
	v452 = *(*int32)(unsafe.Add(mBase, uint32(v334+int32(-17))))
	v453 = v452
	goto L150
L152:
	;
	v449 = *(*int32)(unsafe.Add(mBase, uint32(v334+int32(-9))))
	v453 = v449
	goto L150
L153:
	;
	v446 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v334+int32(-5)))))
	v453 = v446
	goto L150
L154:
	;
	v443 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v334+int32(-3)))))
	v453 = v443
	goto L150
L155:
	;
	v453 = int32(base.Ui32(v436) >> (uint(int32(3)) % 32))
	goto L150
L156:
	;
	v455 = v433
	goto L158
L157:
	;
	v455 = v453
	goto L158
L158:
	;
	v456 = F_memcmp(m, v253, v334, v455)
	mBase = m.M
	if v456 != 0 {
		goto L159
	} else {
		goto L160
	}
L159:
	;
	v459 = v456
	goto L161
L160:
	;
	v459 = base.B2i32(base.Ui32(v453) < base.Ui32(v433)) - v454
	goto L161
L161:
	;
	goto L143
L162:
	;
	goto L75
L163:
	;
	v492 = v238 + int32(1)
	v493 = *(*float64)(unsafe.Add(mBase, uint32(v233)))
	v494 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	m.T0[v494].(func(*base.Module, int32, int32, int32, float64))(m, l0, v253, v488, v493)
	mBase = m.M
	v496 = m.ExcPending
	if v496 != 0 {
		goto L1
	} else {
		goto L169
	}
L164:
	;
	v487 = *(*int32)(unsafe.Add(mBase, uint32(v251+int32(-16))))
	v488 = v487
	goto L163
L165:
	;
	v484 = *(*int32)(unsafe.Add(mBase, uint32(v251+int32(-8))))
	v488 = v484
	goto L163
L166:
	;
	v481 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v251+int32(-4)))))
	v488 = v481
	goto L163
L167:
	;
	v478 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v251+int32(-2)))))
	v488 = v478
	goto L163
L168:
	;
	v488 = int32(base.Ui32(v471) >> (uint(int32(3)) % 32))
	goto L163
L169:
	;
	v498 = *(*int32)(unsafe.Add(mBase, uint32(v233+v227)))
	if v498 != 0 {
		v233 = v498
		v237 = v237 + int32(-1)
		v238 = v492
		goto L71
	} else {
		goto L170
	}
L170:
	;
	goto L72
L171:
	;
	m.G0 = v18 + int32(48)
	return
L172:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L173:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_scanGenericCommand(m *base.Module, l0 int32, l1 int32, l2 int64) {
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v11 int32
	_ = v11
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v21 int32
	_ = v21
	v5 = m.G0
	v7 = v5 - int32(48)
	m.G0 = v7
	if l1 != 0 {
		v11 = int32(3)
	} else {
		v11 = int32(2)
	}
	v15 = F_parseScanOptionsOrReply(m, l0, l1, v11, int32(0), v7+int32(8))
	v16 = m.ExcPending
	if v16 != 0 {
		return
	} else {
		if v15 != 0 {
			m.G0 = v7 + int32(48)
			return
		} else {
			F_scanGenericCommandWithOptions(m, l0, l1, l2, v7+int32(8), int32(0))
			v21 = m.ExcPending
			if v21 != 0 {
				return
			} else {
				m.G0 = v7 + int32(48)
				return
			}
		}
	}
}
func F_setGenericCommand(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32) {
	mBase := m.M
	_ = mBase
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int64
	_ = v27
	var v40 int64
	_ = v40
	var v41 int64
	_ = v41
	var v42 int64
	_ = v42
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
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
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v91 int32
	_ = v91
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v102 int32
	_ = v102
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v121 int32
	_ = v121
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v128 int32
	_ = v128
	var v131 int64
	_ = v131
	var v135 int32
	_ = v135
	var v139 int32
	_ = v139
	var v141 int32
	_ = v141
	var v142 int64
	_ = v142
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v147 int32
	_ = v147
	var v149 int32
	_ = v149
	var v153 int32
	_ = v153
	var v159 int32
	_ = v159
	var v161 int32
	_ = v161
	var v163 int32
	_ = v163
	var v166 int32
	_ = v166
	var v170 int32
	_ = v170
	var v173 int32
	_ = v173
	var v178 int32
	_ = v178
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v185 int32
	_ = v185
	var v188 int32
	_ = v188
	var v189 int64
	_ = v189
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v193 int32
	_ = v193
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
	var v213 int32
	_ = v213
	var v216 int32
	_ = v216
	var v218 int32
	_ = v218
	var v220 int64
	_ = v220
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v229 int32
	_ = v229
	var v234 int64
	_ = v234
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v241 int32
	_ = v241
	var v250 int32
	_ = v250
	var v252 int32
	_ = v252
	var v255 int32
	_ = v255
	var v260 int32
	_ = v260
	var v262 int32
	_ = v262
	var v263 int32
	_ = v263
	var v266 int32
	_ = v266
	var v270 int32
	_ = v270
	var v272 int32
	_ = v272
	var v276 int32
	_ = v276
	var v277 int32
	_ = v277
	var v279 int32
	_ = v279
	var v284 int32
	_ = v284
	var v285 int32
	_ = v285
	var v287 int32
	_ = v287
	var v288 int32
	_ = v288
	var v293 int32
	_ = v293
	var v294 int32
	_ = v294
	var v295 int32
	_ = v295
	var v296 int32
	_ = v296
	var v300 int32
	_ = v300
	var v301 int32
	_ = v301
	var v302 int32
	_ = v302
	var v303 int32
	_ = v303
	var v304 int32
	_ = v304
	var v307 int32
	_ = v307
	var v308 int32
	_ = v308
	var v311 int32
	_ = v311
	var v312 int32
	_ = v312
	var v313 int32
	_ = v313
	var v314 int32
	_ = v314
	var v315 int32
	_ = v315
	var v318 int32
	_ = v318
	var v319 int32
	_ = v319
	var v320 int32
	_ = v320
	var v324 int32
	_ = v324
	var v325 int32
	_ = v325
	var v326 int32
	_ = v326
	var v327 int32
	_ = v327
	var v328 int32
	_ = v328
	var v331 int32
	_ = v331
	var v332 int32
	_ = v332
	var v338 int32
	_ = v338
	var v341 int32
	_ = v341
	var v349 int32
	_ = v349
	var v351 int32
	_ = v351
	var v353 int32
	_ = v353
	var v354 int32
	_ = v354
	var v355 int32
	_ = v355
	var v360 int32
	_ = v360
	var v365 int32
	_ = v365
	var v370 int32
	_ = v370
	var v376 int32
	_ = v376
	var v378 int32
	_ = v378
	var v381 int32
	_ = v381
	var v385 int32
	_ = v385
	var v387 int32
	_ = v387
	var v388 int32
	_ = v388
	var v394 int32
	_ = v394
	var v403 int32
	_ = v403
	var v418 int32
	_ = v418
	var v438 int32
	_ = v438
	v13 = m.G0
	v15 = v13 - int32(48)
	m.G0 = v15
	*(*int32)(unsafe.Add(mBase, uint32(v15)+44)) = l3
	*(*int64)(unsafe.Add(mBase, uint32(v15)+32)) = int64(0)
	if l4 == int32(0) {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	F__serverAssert(m, int32(_a_F_setGenericCommand_0), int32(_a_F_setGenericCommand_1), int32(172))
	mBase = m.M
	v438 = m.ExcPending
	if v438 != 0 {
		goto L5
	} else {
		goto L146
	}
L2:
	;
	m.G0 = v15 + int32(48)
	return
L3:
	;
	v51 = l1 & int32(32)
	if v51 == int32(0) {
		goto L20
	} else {
		goto L21
	}
L4:
	;
	v25 = F_getLongLongFromObjectOrReply(m, l0, l4, v15+int32(32), int32(0))
	mBase = m.M
	v26 = m.ExcPending
	if v26 != 0 {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	return
L6:
	;
	if v25 != 0 {
		goto L2
	} else {
		goto L7
	}
L7:
	;
	v27 = *(*int64)(unsafe.Add(mBase, uint32(v15)+32))
	if v27 < int64(1) {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	F_addReplyErrorExpireTime(m, l0)
	mBase = m.M
	v48 = m.ExcPending
	if v48 != 0 {
		goto L5
	} else {
		goto L18
	}
L9:
	;
	if l5 != 0 {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	if l5 != 0 {
		goto L13
	} else {
		goto L14
	}
L11:
	;
	if base.Ui64(int64(9223372036854775)) < base.Ui64(v27) {
		goto L8
	} else {
		goto L12
	}
L12:
	;
	goto L10
L13:
	;
	if l1&int32(12) == int32(0) {
		goto L3
	} else {
		goto L15
	}
L14:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v15)+32)) = v27 * int64(1000)
	goto L13
L15:
	;
	v40 = *(*int64)(unsafe.Add(mBase, _c_F_setGenericCommand[0]))
	goto L16
L16:
	;
	v41 = *(*int64)(unsafe.Add(mBase, uint32(v15)+32))
	v42 = v40 + v41
	*(*int64)(unsafe.Add(mBase, uint32(v15)+32)) = v42
	if int64(0) < v42 {
		goto L3
	} else {
		goto L17
	}
L17:
	;
	goto L8
L18:
	;
	goto L2
L19:
	;
	F_commitDeferredReplyBuffer(m, l0, int32(1))
	mBase = m.M
	v418 = m.ExcPending
	if v418 != 0 {
		goto L5
	} else {
		goto L145
	}
L20:
	;
	v74 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v75 = F_lookupKeyWrite(m, v74, l2)
	mBase = m.M
	v76 = m.ExcPending
	if v76 != 0 {
		goto L5
	} else {
		goto L28
	}
L21:
	;
	F_initDeferredReplyBuffer(m, l0)
	mBase = m.M
	v55 = m.ExcPending
	if v55 != 0 {
		goto L5
	} else {
		goto L22
	}
L22:
	;
	v56 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v56)+4))
	v59 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+224)))
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v59<<(uint(int32(2))%32))+uint32(_c_F_setGenericCommand[1])))
	v64 = F_lookupKeyReadOrReply(m, l0, v57, v63)
	mBase = m.M
	v65 = m.ExcPending
	if v65 != 0 {
		goto L5
	} else {
		goto L23
	}
L23:
	;
	if v64 == int32(0) {
		goto L20
	} else {
		goto L24
	}
L24:
	;
	v69 = F_checkType(m, l0, v64, int32(0))
	mBase = m.M
	v70 = m.ExcPending
	if v70 != 0 {
		goto L5
	} else {
		goto L25
	}
L25:
	;
	if v69 != 0 {
		goto L19
	} else {
		goto L26
	}
L26:
	;
	F_addReplyBulk(m, l0, v64)
	mBase = m.M
	v72 = m.ExcPending
	if v72 != 0 {
		goto L5
	} else {
		goto L27
	}
L27:
	;
	goto L20
L28:
	;
	v78 = l1 & int32(512)
	if v78 == int32(0) {
		goto L31
	} else {
		goto L32
	}
L29:
	;
	if l1&int32(1) == int32(0) {
		goto L52
	} else {
		goto L53
	}
L30:
	;
	v110 = F_compareStringObjects(m, v75, l8)
	mBase = m.M
	v111 = m.ExcPending
	if v111 != 0 {
		goto L5
	} else {
		goto L48
	}
L31:
	;
	if v78 == int32(0) {
		goto L29
	} else {
		goto L42
	}
L32:
	;
	if v75 == int32(0) {
		goto L31
	} else {
		goto L33
	}
L33:
	;
	if v51 != 0 {
		goto L30
	} else {
		goto L34
	}
L34:
	;
	v84 = F_checkType(m, l0, v75, int32(0))
	mBase = m.M
	v85 = m.ExcPending
	if v85 != 0 {
		goto L5
	} else {
		goto L35
	}
L35:
	;
	if v84 != 0 {
		goto L19
	} else {
		goto L36
	}
L36:
	;
	v86 = F_compareStringObjects(m, v75, l8)
	mBase = m.M
	v87 = m.ExcPending
	if v87 != 0 {
		goto L5
	} else {
		goto L37
	}
L37:
	;
	if v86 == int32(0) {
		goto L29
	} else {
		goto L38
	}
L38:
	;
	if l7 != 0 {
		v96 = l7
		goto L39
	} else {
		goto L40
	}
L39:
	;
	F_addReply(m, l0, v96)
	mBase = m.M
	v98 = m.ExcPending
	if v98 != 0 {
		goto L5
	} else {
		goto L41
	}
L40:
	;
	v91 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+224)))
	v95 = *(*int32)(unsafe.Add(mBase, uint32(v91<<(uint(int32(2))%32))+uint32(_c_F_setGenericCommand[1])))
	v96 = v95
	goto L39
L41:
	;
	goto L19
L42:
	;
	if v75 != 0 {
		goto L29
	} else {
		goto L43
	}
L43:
	;
	if v51 != 0 {
		goto L19
	} else {
		goto L44
	}
L44:
	;
	if l7 != 0 {
		v107 = l7
		goto L45
	} else {
		goto L46
	}
L45:
	;
	F_addReply(m, l0, v107)
	mBase = m.M
	v109 = m.ExcPending
	if v109 != 0 {
		goto L5
	} else {
		goto L47
	}
L46:
	;
	v102 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+224)))
	v106 = *(*int32)(unsafe.Add(mBase, uint32(v102<<(uint(int32(2))%32))+uint32(_c_F_setGenericCommand[1])))
	v107 = v106
	goto L45
L47:
	;
	goto L19
L48:
	;
	if v110 != 0 {
		goto L19
	} else {
		goto L49
	}
L49:
	;
	goto L29
L50:
	;
	if l4 == int32(0) {
		goto L61
	} else {
		goto L62
	}
L51:
	;
	if v51 != 0 {
		goto L19
	} else {
		goto L57
	}
L52:
	;
	if l1&int32(2) == int32(0) {
		goto L50
	} else {
		goto L55
	}
L53:
	;
	if v75 != 0 {
		goto L51
	} else {
		goto L54
	}
L54:
	;
	goto L52
L55:
	;
	if v75 != 0 {
		goto L50
	} else {
		goto L56
	}
L56:
	;
	goto L51
L57:
	;
	if l7 != 0 {
		v126 = l7
		goto L58
	} else {
		goto L59
	}
L58:
	;
	F_addReply(m, l0, v126)
	mBase = m.M
	v128 = m.ExcPending
	if v128 != 0 {
		goto L5
	} else {
		goto L60
	}
L59:
	;
	v121 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+224)))
	v125 = *(*int32)(unsafe.Add(mBase, uint32(v121<<(uint(int32(2))%32))+uint32(_c_F_setGenericCommand[1])))
	v126 = v125
	goto L58
L60:
	;
	goto L19
L61:
	;
	if v75 != 0 {
		goto L77
	} else {
		goto L78
	}
L62:
	;
	v131 = *(*int64)(unsafe.Add(mBase, uint32(v15)+32))
	v135 = *(*int32)(unsafe.Add(mBase, _c_F_setGenericCommand[2]))
	if v135 == int32(0) {
		goto L65
	} else {
		goto L66
	}
L63:
	;
	if v153 == int32(0) {
		goto L61
	} else {
		goto L71
	}
L64:
	;
	goto L63
L65:
	;
	v141 = int32(0)
	v142 = F_commandTimeSnapshot(m)
	mBase = m.M
	if v142 < v131 {
		v153 = v141
		goto L64
	} else {
		goto L68
	}
L66:
	;
	v139 = *(*int32)(unsafe.Add(mBase, uint32(v135)+216))
	if v139 != 0 {
		v153 = int32(0)
		goto L64
	} else {
		goto L67
	}
L67:
	;
	goto L65
L68:
	;
	v144 = int32(_a_F_setGenericCommand_2)
	v145 = *(*int32)(unsafe.Add(mBase, _c_F_setGenericCommand[3]))
	v147 = *(*int32)(unsafe.Add(mBase, _c_F_setGenericCommand[4]))
	if v147 != 0 {
		v153 = v141
		goto L64
	} else {
		goto L69
	}
L69:
	;
	if v145 != 0 {
		v153 = v141
		goto L64
	} else {
		goto L70
	}
L70:
	;
	v149 = *(*int32)(unsafe.Add(mBase, _c_F_setGenericCommand[5]))
	v153 = base.B2i32(v149 == int32(0))
	goto L64
L71:
	;
	if v75 == int32(0) {
		goto L72
	} else {
		goto L73
	}
L72:
	;
	if v51 != 0 {
		goto L19
	} else {
		goto L75
	}
L73:
	;
	F_deleteExpiredKeyFromOverwriteAndPropagate(m, l0, l2)
	mBase = m.M
	v159 = m.ExcPending
	if v159 != 0 {
		goto L5
	} else {
		goto L74
	}
L74:
	;
	goto L72
L75:
	;
	v161 = *(*int32)(unsafe.Add(mBase, _c_F_setGenericCommand[6]))
	F_addReply(m, l0, v161)
	mBase = m.M
	v163 = m.ExcPending
	if v163 != 0 {
		goto L5
	} else {
		goto L76
	}
L76:
	;
	goto L19
L77:
	;
	v166 = int32(4)
	goto L79
L78:
	;
	v166 = int32(8)
	goto L79
L79:
	;
	v170 = int32(0)
	v173 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+208)))
	if v173&int32(1) == v170 {
		goto L80
	} else {
		goto L81
	}
L80:
	;
	v181 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	F_setKey(m, l0, v181, l2, v15+int32(44), v166|base.B2i32(l1&int32(16)|l4 != v170))
	mBase = m.M
	v185 = m.ExcPending
	if v185 != 0 {
		goto L5
	} else {
		goto L83
	}
L81:
	;
	v178 = *(*int32)(unsafe.Add(mBase, uint32(v15)+44))
	F_incrRefCount(m, v178)
	mBase = m.M
	v180 = m.ExcPending
	if v180 != 0 {
		goto L5
	} else {
		goto L82
	}
L82:
	;
	goto L80
L83:
	;
	if l4 == int32(0) {
		goto L84
	} else {
		goto L85
	}
L84:
	;
	v193 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+208)))
	if v193&int32(1) == int32(0) {
		goto L88
	} else {
		goto L89
	}
L85:
	;
	v188 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v189 = *(*int64)(unsafe.Add(mBase, uint32(v15)+32))
	v190 = F_setExpire(m, l0, v188, l2, v189)
	mBase = m.M
	v191 = m.ExcPending
	if v191 != 0 {
		goto L5
	} else {
		goto L86
	}
L86:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+44)) = v190
	goto L84
L87:
	;
	v218 = int32(_a_F_setGenericCommand_2)
	v220 = *(*int64)(unsafe.Add(mBase, _c_F_setGenericCommand[7]))
	*(*int64)(unsafe.Add(mBase, _c_F_setGenericCommand[7])) = v220 + int64(1)
	v226 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v227 = *(*int32)(unsafe.Add(mBase, uint32(v226)+28))
	F_notifyKeyspaceEvent(m, int32(8), int32(_a_F_setGenericCommand_3), l2, v227)
	mBase = m.M
	v229 = m.ExcPending
	if v229 != 0 {
		goto L5
	} else {
		goto L98
	}
L88:
	;
	v206 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if l1&int32(1024) != 0 {
		goto L94
	} else {
		goto L95
	}
L89:
	;
	if l1&int32(1024) != 0 {
		goto L90
	} else {
		goto L91
	}
L90:
	;
	v202 = int32(3)
	goto L92
L91:
	;
	v202 = int32(2)
	goto L92
L92:
	;
	v203 = *(*int32)(unsafe.Add(mBase, uint32(v15)+44))
	F_rewriteClientCommandArgument(m, l0, v202, v203)
	mBase = m.M
	v205 = m.ExcPending
	if v205 != 0 {
		goto L5
	} else {
		goto L93
	}
L93:
	;
	goto L87
L94:
	;
	v211 = int32(12)
	goto L96
L95:
	;
	v211 = int32(8)
	goto L96
L96:
	;
	v213 = *(*int32)(unsafe.Add(mBase, uint32(v15)+44))
	*(*int32)(unsafe.Add(mBase, uint32(v206+v211))) = v213
	F_incrRefCount(m, v213)
	mBase = m.M
	v216 = m.ExcPending
	if v216 != 0 {
		goto L5
	} else {
		goto L97
	}
L97:
	;
	goto L87
L98:
	;
	if l4 == int32(0) {
		goto L101
	} else {
		goto L102
	}
L99:
	;
	v288 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v293 = F_valkey_malloc(m, v288<<(uint(int32(2))%32)+int32(-4))
	mBase = m.M
	v294 = m.ExcPending
	if v294 != 0 {
		goto L5
	} else {
		goto L122
	}
L100:
	;
	v284 = *(*int32)(unsafe.Add(mBase, _c_F_setGenericCommand[6]))
	if l6 != 0 {
		goto L118
	} else {
		goto L119
	}
L101:
	;
	if v51 != 0 {
		goto L99
	} else {
		goto L117
	}
L102:
	;
	if l1&int32(128) != 0 {
		goto L103
	} else {
		goto L104
	}
L103:
	;
	v276 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v277 = *(*int32)(unsafe.Add(mBase, uint32(v276)+28))
	F_notifyKeyspaceEvent(m, int32(4), int32(_a_F_setGenericCommand_4), l2, v277)
	mBase = m.M
	v279 = m.ExcPending
	if v279 != 0 {
		goto L5
	} else {
		goto L115
	}
L104:
	;
	v234 = *(*int64)(unsafe.Add(mBase, uint32(v15)+32))
	v235 = F_createStringObjectFromLongLong(m, v234)
	mBase = m.M
	v236 = m.ExcPending
	if v236 != 0 {
		goto L5
	} else {
		goto L105
	}
L105:
	;
	v237 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
	v238 = *(*int32)(unsafe.Add(mBase, uint32(v237)+48))
	if v238 != int32(406) {
		goto L107
	} else {
		goto L108
	}
L106:
	;
	F_decrRefCount(m, v235)
	mBase = m.M
	v272 = m.ExcPending
	if v272 != 0 {
		goto L5
	} else {
		goto L114
	}
L107:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15+int32(16)))) = v235
	*(*int32)(unsafe.Add(mBase, uint32(v15)+4)) = l2
	v260 = *(*int32)(unsafe.Add(mBase, uint32(v15)+44))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+8)) = v260
	v262 = int32(_a_F_setGenericCommand_5)
	v263 = *(*int32)(unsafe.Add(mBase, _c_F_setGenericCommand[8]))
	*(*int32)(unsafe.Add(mBase, uint32(v15))) = v263
	v266 = *(*int32)(unsafe.Add(mBase, _c_F_setGenericCommand[9]))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+12)) = v266
	F_rewriteClientCommandVector(m, l0, int32(5), v15)
	mBase = m.M
	v270 = m.ExcPending
	if v270 != 0 {
		goto L5
	} else {
		goto L113
	}
L108:
	;
	v241 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v241 != int32(5) {
		goto L107
	} else {
		goto L109
	}
L109:
	;
	if l1&int32(76) == int32(0) {
		goto L1
	} else {
		goto L110
	}
L110:
	;
	v250 = *(*int32)(unsafe.Add(mBase, _c_F_setGenericCommand[9]))
	F_rewriteClientCommandArgument(m, l0, int32(3), v250)
	mBase = m.M
	v252 = m.ExcPending
	if v252 != 0 {
		goto L5
	} else {
		goto L111
	}
L111:
	;
	F_rewriteClientCommandArgument(m, l0, int32(4), v235)
	mBase = m.M
	v255 = m.ExcPending
	if v255 != 0 {
		goto L5
	} else {
		goto L112
	}
L112:
	;
	goto L106
L113:
	;
	goto L106
L114:
	;
	goto L103
L115:
	;
	if v51 == int32(0) {
		goto L100
	} else {
		goto L116
	}
L116:
	;
	goto L19
L117:
	;
	goto L100
L118:
	;
	v285 = l6
	goto L120
L119:
	;
	v285 = v284
	goto L120
L120:
	;
	F_addReply(m, l0, v285)
	mBase = m.M
	v287 = m.ExcPending
	if v287 != 0 {
		goto L5
	} else {
		goto L121
	}
L121:
	;
	goto L19
L122:
	;
	v295 = int32(1)
	v296 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v295 <= v296 {
		goto L124
	} else {
		goto L125
	}
L123:
	;
	F_replaceClientCommandVector(m, l0, v394, v293)
	mBase = m.M
	v403 = m.ExcPending
	if v403 != 0 {
		goto L5
	} else {
		goto L144
	}
L124:
	;
	v300 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v301 = *(*int32)(unsafe.Add(mBase, uint32(v300)))
	v302 = F_objectGetVal(m, v301)
	mBase = m.M
	v303 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v304 = *(*int32)(unsafe.Add(mBase, uint32(v303)))
	*(*int32)(unsafe.Add(mBase, uint32(v293))) = v304
	F_incrRefCount(m, v304)
	mBase = m.M
	v307 = m.ExcPending
	if v307 != 0 {
		goto L5
	} else {
		goto L126
	}
L125:
	;
	v394 = int32(0)
	goto L123
L126:
	;
	v308 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v308 < int32(2) {
		v394 = v295
		goto L123
	} else {
		goto L127
	}
L127:
	;
	v311 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v312 = *(*int32)(unsafe.Add(mBase, uint32(v311)+4))
	v313 = F_objectGetVal(m, v312)
	mBase = m.M
	v314 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v315 = *(*int32)(unsafe.Add(mBase, uint32(v314)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v293)+4)) = v315
	F_incrRefCount(m, v315)
	mBase = m.M
	v318 = m.ExcPending
	if v318 != 0 {
		goto L5
	} else {
		goto L128
	}
L128:
	;
	v319 = int32(3)
	v320 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v319 <= v320 {
		goto L129
	} else {
		goto L130
	}
L129:
	;
	v324 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v325 = *(*int32)(unsafe.Add(mBase, uint32(v324)+8))
	v326 = F_objectGetVal(m, v325)
	mBase = m.M
	v327 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v328 = *(*int32)(unsafe.Add(mBase, uint32(v327)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v293)+8)) = v328
	F_incrRefCount(m, v328)
	mBase = m.M
	v331 = m.ExcPending
	if v331 != 0 {
		goto L5
	} else {
		goto L131
	}
L130:
	;
	v394 = int32(2)
	goto L123
L131:
	;
	v332 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if int32(4) <= v332 {
		goto L132
	} else {
		goto L133
	}
L132:
	;
	v338 = v319
	v341 = int32(3)
	goto L134
L133:
	;
	v394 = int32(3)
	goto L123
L134:
	;
	v349 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v351 = v338 << (uint(int32(2)) % 32)
	v353 = *(*int32)(unsafe.Add(mBase, uint32(v349+v351)))
	v354 = F_objectGetVal(m, v353)
	mBase = m.M
	v355 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v354))))
	if v355|int32(32) != int32(103) {
		goto L137
	} else {
		goto L138
	}
L135:
	;
	v394 = v385
	goto L123
L136:
	;
	v387 = v338 + int32(1)
	v388 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v387 < v388 {
		v338 = v387
		v341 = v385
		goto L134
	} else {
		goto L143
	}
L137:
	;
	v376 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v378 = *(*int32)(unsafe.Add(mBase, uint32(v376+v351)))
	*(*int32)(unsafe.Add(mBase, uint32(v293+v341<<(uint(int32(2))%32)))) = v378
	F_incrRefCount(m, v378)
	mBase = m.M
	v381 = m.ExcPending
	if v381 != 0 {
		goto L5
	} else {
		goto L142
	}
L138:
	;
	v360 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v354)+1)))
	if v360|int32(32) != int32(101) {
		goto L137
	} else {
		goto L139
	}
L139:
	;
	v365 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v354)+2)))
	if v365|int32(32) != int32(116) {
		goto L137
	} else {
		goto L140
	}
L140:
	;
	v370 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v354)+3)))
	if v370 == int32(0) {
		v385 = v341
		goto L136
	} else {
		goto L141
	}
L141:
	;
	goto L137
L142:
	;
	v385 = v341 + int32(1)
	goto L136
L143:
	;
	goto L135
L144:
	;
	goto L19
L145:
	;
	goto L2
L146:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
}
