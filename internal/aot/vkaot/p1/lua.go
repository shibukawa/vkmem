package p1

import (
	base "github.com/shibukawa/vkmem/internal/aot/vkaot/base"
	"unsafe"
)

func F_initializeLuaState(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v17 int32
	_ = v17
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v55 int32
	_ = v55
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	var v97 int32
	_ = v97
	var v121 int32
	_ = v121
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
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
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v154 int32
	_ = v154
	var v155 int64
	_ = v155
	var v157 int32
	_ = v157
	var v159 int32
	_ = v159
	var v164 int32
	_ = v164
	var v166 int32
	_ = v166
	var v185 int32
	_ = v185
	var v196 int32
	_ = v196
	var v220 int32
	_ = v220
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v253 int32
	_ = v253
	var v254 int64
	_ = v254
	var v256 int32
	_ = v256
	var v258 int32
	_ = v258
	var v263 int32
	_ = v263
	var v265 int32
	_ = v265
	var v284 int32
	_ = v284
	var v295 int32
	_ = v295
	var v297 int32
	_ = v297
	var v298 int32
	_ = v298
	var v304 int32
	_ = v304
	var v305 int32
	_ = v305
	var v307 int32
	_ = v307
	v5 = F_luaL_newstate(m)
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return
	} else {
		if l1 != 0 {
			if l1 != int32(1) {
				v298 = m.G3
				v304 = m.G8
				v305 = *(*int32)(unsafe.Add(mBase, uint32(v304)))
				m.T0[v305].(func(*base.Module, int32, int32, int32))(m, v298+int32(_a2609), v298+int32(_a2610), int32(180))
				mBase = m.M
				v307 = m.ExcPending
				if v307 != 0 {
					return
				} else {
					m.Env.Exit(m, int32(1))
					mBase = m.M
					base.Wasm_trap_unreachable()
					for {
					}
				}
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v5
				F_luaRegisterServerAPI(m, l0, v5)
				mBase = m.M
				v12 = m.ExcPending
				if v12 != 0 {
					return
				} else {
					v13 = m.G3
					F_lua_pushstring(m, v5, v13+int32(_a2611))
					mBase = m.M
					v17 = m.ExcPending
					if v17 != 0 {
						return
					} else {
						v23 = F_luaL_loadbuffer(m, v5, v13+int32(_a2612), int32(355), v13+int32(_a2613))
						mBase = m.M
						v24 = m.ExcPending
						if v24 != 0 {
							return
						} else {
							v25 = int32(0)
							v28 = F_lua_pcall(m, v5, v25, int32(1), v25)
							mBase = m.M
							v29 = m.ExcPending
							if v29 != 0 {
								return
							} else {
								F_lua_settable(m, v5, int32(-10000))
								mBase = m.M
								v32 = m.ExcPending
								if v32 != 0 {
									return
								} else {
									if l1 != 0 {
										switch int32(0) {
										case 0:
											v250 = v5 + int32(72)
										case 1:
											v223 = *(*int32)(unsafe.Add(mBase, uint32(v5)+20))
											v224 = *(*int32)(unsafe.Add(mBase, uint32(v223)+4))
											v225 = *(*int32)(unsafe.Add(mBase, uint32(v224)))
											v226 = *(*int32)(unsafe.Add(mBase, uint32(v225)+12))
											*(*int32)(unsafe.Add(mBase, uint32(v5)+96)) = int32(5)
											*(*int32)(unsafe.Add(mBase, uint32(v5)+88)) = v226
											v250 = v5 + int32(88)
										case 2:
											v220 = *(*int32)(unsafe.Add(mBase, uint32(v5)+16))
											v250 = v220 + int32(96)
										default:
											v234 = *(*int32)(unsafe.Add(mBase, uint32(v5)+20))
											v235 = *(*int32)(unsafe.Add(mBase, uint32(v234)+4))
											v236 = *(*int32)(unsafe.Add(mBase, uint32(v235)))
											v237 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v236)+7)))
											v238 = m.G398
											if base.Ui32(v237) < base.Ui32(int32(0)) {
												v249 = v238
											} else {
												v249 = v236 + int32(8)
											}
											v250 = v249
										}
										v253 = *(*int32)(unsafe.Add(mBase, uint32(v5)+8))
										v254 = *(*int64)(unsafe.Add(mBase, uint32(v250)))
										*(*int64)(unsafe.Add(mBase, uint32(v253))) = v254
										v256 = *(*int32)(unsafe.Add(mBase, uint32(v250)+8))
										*(*int32)(unsafe.Add(mBase, uint32(v253)+8)) = v256
										v258 = *(*int32)(unsafe.Add(mBase, uint32(v5)+8))
										*(*int32)(unsafe.Add(mBase, uint32(v5)+8)) = v258 + int32(16)
										F_luaSetErrorMetatable(m, v5)
										mBase = m.M
										v263 = m.ExcPending
										if v263 != 0 {
											return
										} else {
											F_luaSetTableProtectionRecursively(m, v5)
											mBase = m.M
											v265 = m.ExcPending
											if v265 != 0 {
												return
											} else {
												v284 = *(*int32)(unsafe.Add(mBase, uint32(v5)+8))
												*(*int32)(unsafe.Add(mBase, uint32(v5)+8)) = v284 + int32(-16)
												F_luaSetTableProtectionForBasicTypes(m, v5)
												mBase = m.M
												v295 = m.ExcPending
												if v295 != 0 {
													return
												} else {
													F_luaFunctionInitializeLuaState(m, l0, v5)
													mBase = m.M
													v297 = m.ExcPending
													if v297 != 0 {
														return
													} else {
														return
													}
												}
											}
										}
									} else {
										v34 = m.G3
										v36 = v34 + int32(_a2614)
										F_lua_getfield(m, v5, int32(-10002), v36)
										mBase = m.M
										v38 = m.ExcPending
										if v38 != 0 {
											return
										} else {
											F_lua_pushstring(m, v5, v34+int32(_a2615))
											mBase = m.M
											v42 = m.ExcPending
											if v42 != 0 {
												return
											} else {
												v43 = m.G5
												F_lua_pushcclosure(m, v5, v43+int32(1169), int32(0))
												mBase = m.M
												v48 = m.ExcPending
												if v48 != 0 {
													return
												} else {
													F_lua_settable(m, v5, int32(-3))
													mBase = m.M
													v51 = m.ExcPending
													if v51 != 0 {
														return
													} else {
														F_lua_pushstring(m, v5, v34+int32(_a2616))
														mBase = m.M
														v55 = m.ExcPending
														if v55 != 0 {
															return
														} else {
															F_lua_pushcclosure(m, v5, v43+int32(1170), int32(0))
															mBase = m.M
															v60 = m.ExcPending
															if v60 != 0 {
																return
															} else {
																F_lua_settable(m, v5, int32(-3))
																mBase = m.M
																v63 = m.ExcPending
																if v63 != 0 {
																	return
																} else {
																	F_lua_pushstring(m, v5, v34+int32(_a2617))
																	mBase = m.M
																	v67 = m.ExcPending
																	if v67 != 0 {
																		return
																	} else {
																		v68 = m.G384
																		F_lua_pushcclosure(m, v5, v68, int32(0))
																		mBase = m.M
																		v71 = m.ExcPending
																		if v71 != 0 {
																			return
																		} else {
																			F_lua_settable(m, v5, int32(-3))
																			mBase = m.M
																			v74 = m.ExcPending
																			if v74 != 0 {
																				return
																			} else {
																				F_lua_setfield(m, v5, int32(-10002), v36)
																				mBase = m.M
																				v77 = m.ExcPending
																				if v77 != 0 {
																					return
																				} else {
																					F_lua_pushstring(m, v5, v34+int32(_a2611))
																					mBase = m.M
																					v81 = m.ExcPending
																					if v81 != 0 {
																						return
																					} else {
																						F_lua_gettable(m, v5, int32(-10000))
																						mBase = m.M
																						v84 = m.ExcPending
																						if v84 != 0 {
																							return
																						} else {
																							v87 = v34 + int32(_a2618)
																							F_lua_setfield(m, v5, int32(-10002), v87)
																							mBase = m.M
																							v89 = m.ExcPending
																							if v89 != 0 {
																								return
																							} else {
																								F_lua_getfield(m, v5, int32(-10002), v87)
																								mBase = m.M
																								v92 = m.ExcPending
																								if v92 != 0 {
																									return
																								} else {
																									F_lua_setfield(m, v5, int32(-10002), v34+int32(_a2619))
																									mBase = m.M
																									v97 = m.ExcPending
																									if v97 != 0 {
																										return
																									} else {
																										switch int32(0) {
																										case 0:
																											v151 = v5 + int32(72)
																										case 1:
																											v124 = *(*int32)(unsafe.Add(mBase, uint32(v5)+20))
																											v125 = *(*int32)(unsafe.Add(mBase, uint32(v124)+4))
																											v126 = *(*int32)(unsafe.Add(mBase, uint32(v125)))
																											v127 = *(*int32)(unsafe.Add(mBase, uint32(v126)+12))
																											*(*int32)(unsafe.Add(mBase, uint32(v5)+96)) = int32(5)
																											*(*int32)(unsafe.Add(mBase, uint32(v5)+88)) = v127
																											v151 = v5 + int32(88)
																										case 2:
																											v121 = *(*int32)(unsafe.Add(mBase, uint32(v5)+16))
																											v151 = v121 + int32(96)
																										default:
																											v135 = *(*int32)(unsafe.Add(mBase, uint32(v5)+20))
																											v136 = *(*int32)(unsafe.Add(mBase, uint32(v135)+4))
																											v137 = *(*int32)(unsafe.Add(mBase, uint32(v136)))
																											v138 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v137)+7)))
																											v139 = m.G398
																											if base.Ui32(v138) < base.Ui32(int32(0)) {
																												v150 = v139
																											} else {
																												v150 = v137 + int32(8)
																											}
																											v151 = v150
																										}
																										v154 = *(*int32)(unsafe.Add(mBase, uint32(v5)+8))
																										v155 = *(*int64)(unsafe.Add(mBase, uint32(v151)))
																										*(*int64)(unsafe.Add(mBase, uint32(v154))) = v155
																										v157 = *(*int32)(unsafe.Add(mBase, uint32(v151)+8))
																										*(*int32)(unsafe.Add(mBase, uint32(v154)+8)) = v157
																										v159 = *(*int32)(unsafe.Add(mBase, uint32(v5)+8))
																										*(*int32)(unsafe.Add(mBase, uint32(v5)+8)) = v159 + int32(16)
																										F_luaSetErrorMetatable(m, v5)
																										mBase = m.M
																										v164 = m.ExcPending
																										if v164 != 0 {
																											return
																										} else {
																											F_luaSetTableProtectionRecursively(m, v5)
																											mBase = m.M
																											v166 = m.ExcPending
																											if v166 != 0 {
																												return
																											} else {
																												v185 = *(*int32)(unsafe.Add(mBase, uint32(v5)+8))
																												*(*int32)(unsafe.Add(mBase, uint32(v5)+8)) = v185 + int32(-16)
																												F_luaSetTableProtectionForBasicTypes(m, v5)
																												mBase = m.M
																												v196 = m.ExcPending
																												if v196 != 0 {
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
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(l0))) = v5
			F_luaRegisterServerAPI(m, l0, v5)
			mBase = m.M
			v12 = m.ExcPending
			if v12 != 0 {
				return
			} else {
				v13 = m.G3
				F_lua_pushstring(m, v5, v13+int32(_a2611))
				mBase = m.M
				v17 = m.ExcPending
				if v17 != 0 {
					return
				} else {
					v23 = F_luaL_loadbuffer(m, v5, v13+int32(_a2612), int32(355), v13+int32(_a2613))
					mBase = m.M
					v24 = m.ExcPending
					if v24 != 0 {
						return
					} else {
						v25 = int32(0)
						v28 = F_lua_pcall(m, v5, v25, int32(1), v25)
						mBase = m.M
						v29 = m.ExcPending
						if v29 != 0 {
							return
						} else {
							F_lua_settable(m, v5, int32(-10000))
							mBase = m.M
							v32 = m.ExcPending
							if v32 != 0 {
								return
							} else {
								if l1 != 0 {
									switch int32(0) {
									case 0:
										v250 = v5 + int32(72)
									case 1:
										v223 = *(*int32)(unsafe.Add(mBase, uint32(v5)+20))
										v224 = *(*int32)(unsafe.Add(mBase, uint32(v223)+4))
										v225 = *(*int32)(unsafe.Add(mBase, uint32(v224)))
										v226 = *(*int32)(unsafe.Add(mBase, uint32(v225)+12))
										*(*int32)(unsafe.Add(mBase, uint32(v5)+96)) = int32(5)
										*(*int32)(unsafe.Add(mBase, uint32(v5)+88)) = v226
										v250 = v5 + int32(88)
									case 2:
										v220 = *(*int32)(unsafe.Add(mBase, uint32(v5)+16))
										v250 = v220 + int32(96)
									default:
										v234 = *(*int32)(unsafe.Add(mBase, uint32(v5)+20))
										v235 = *(*int32)(unsafe.Add(mBase, uint32(v234)+4))
										v236 = *(*int32)(unsafe.Add(mBase, uint32(v235)))
										v237 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v236)+7)))
										v238 = m.G398
										if base.Ui32(v237) < base.Ui32(int32(0)) {
											v249 = v238
										} else {
											v249 = v236 + int32(8)
										}
										v250 = v249
									}
									v253 = *(*int32)(unsafe.Add(mBase, uint32(v5)+8))
									v254 = *(*int64)(unsafe.Add(mBase, uint32(v250)))
									*(*int64)(unsafe.Add(mBase, uint32(v253))) = v254
									v256 = *(*int32)(unsafe.Add(mBase, uint32(v250)+8))
									*(*int32)(unsafe.Add(mBase, uint32(v253)+8)) = v256
									v258 = *(*int32)(unsafe.Add(mBase, uint32(v5)+8))
									*(*int32)(unsafe.Add(mBase, uint32(v5)+8)) = v258 + int32(16)
									F_luaSetErrorMetatable(m, v5)
									mBase = m.M
									v263 = m.ExcPending
									if v263 != 0 {
										return
									} else {
										F_luaSetTableProtectionRecursively(m, v5)
										mBase = m.M
										v265 = m.ExcPending
										if v265 != 0 {
											return
										} else {
											v284 = *(*int32)(unsafe.Add(mBase, uint32(v5)+8))
											*(*int32)(unsafe.Add(mBase, uint32(v5)+8)) = v284 + int32(-16)
											F_luaSetTableProtectionForBasicTypes(m, v5)
											mBase = m.M
											v295 = m.ExcPending
											if v295 != 0 {
												return
											} else {
												F_luaFunctionInitializeLuaState(m, l0, v5)
												mBase = m.M
												v297 = m.ExcPending
												if v297 != 0 {
													return
												} else {
													return
												}
											}
										}
									}
								} else {
									v34 = m.G3
									v36 = v34 + int32(_a2614)
									F_lua_getfield(m, v5, int32(-10002), v36)
									mBase = m.M
									v38 = m.ExcPending
									if v38 != 0 {
										return
									} else {
										F_lua_pushstring(m, v5, v34+int32(_a2615))
										mBase = m.M
										v42 = m.ExcPending
										if v42 != 0 {
											return
										} else {
											v43 = m.G5
											F_lua_pushcclosure(m, v5, v43+int32(1169), int32(0))
											mBase = m.M
											v48 = m.ExcPending
											if v48 != 0 {
												return
											} else {
												F_lua_settable(m, v5, int32(-3))
												mBase = m.M
												v51 = m.ExcPending
												if v51 != 0 {
													return
												} else {
													F_lua_pushstring(m, v5, v34+int32(_a2616))
													mBase = m.M
													v55 = m.ExcPending
													if v55 != 0 {
														return
													} else {
														F_lua_pushcclosure(m, v5, v43+int32(1170), int32(0))
														mBase = m.M
														v60 = m.ExcPending
														if v60 != 0 {
															return
														} else {
															F_lua_settable(m, v5, int32(-3))
															mBase = m.M
															v63 = m.ExcPending
															if v63 != 0 {
																return
															} else {
																F_lua_pushstring(m, v5, v34+int32(_a2617))
																mBase = m.M
																v67 = m.ExcPending
																if v67 != 0 {
																	return
																} else {
																	v68 = m.G384
																	F_lua_pushcclosure(m, v5, v68, int32(0))
																	mBase = m.M
																	v71 = m.ExcPending
																	if v71 != 0 {
																		return
																	} else {
																		F_lua_settable(m, v5, int32(-3))
																		mBase = m.M
																		v74 = m.ExcPending
																		if v74 != 0 {
																			return
																		} else {
																			F_lua_setfield(m, v5, int32(-10002), v36)
																			mBase = m.M
																			v77 = m.ExcPending
																			if v77 != 0 {
																				return
																			} else {
																				F_lua_pushstring(m, v5, v34+int32(_a2611))
																				mBase = m.M
																				v81 = m.ExcPending
																				if v81 != 0 {
																					return
																				} else {
																					F_lua_gettable(m, v5, int32(-10000))
																					mBase = m.M
																					v84 = m.ExcPending
																					if v84 != 0 {
																						return
																					} else {
																						v87 = v34 + int32(_a2618)
																						F_lua_setfield(m, v5, int32(-10002), v87)
																						mBase = m.M
																						v89 = m.ExcPending
																						if v89 != 0 {
																							return
																						} else {
																							F_lua_getfield(m, v5, int32(-10002), v87)
																							mBase = m.M
																							v92 = m.ExcPending
																							if v92 != 0 {
																								return
																							} else {
																								F_lua_setfield(m, v5, int32(-10002), v34+int32(_a2619))
																								mBase = m.M
																								v97 = m.ExcPending
																								if v97 != 0 {
																									return
																								} else {
																									switch int32(0) {
																									case 0:
																										v151 = v5 + int32(72)
																									case 1:
																										v124 = *(*int32)(unsafe.Add(mBase, uint32(v5)+20))
																										v125 = *(*int32)(unsafe.Add(mBase, uint32(v124)+4))
																										v126 = *(*int32)(unsafe.Add(mBase, uint32(v125)))
																										v127 = *(*int32)(unsafe.Add(mBase, uint32(v126)+12))
																										*(*int32)(unsafe.Add(mBase, uint32(v5)+96)) = int32(5)
																										*(*int32)(unsafe.Add(mBase, uint32(v5)+88)) = v127
																										v151 = v5 + int32(88)
																									case 2:
																										v121 = *(*int32)(unsafe.Add(mBase, uint32(v5)+16))
																										v151 = v121 + int32(96)
																									default:
																										v135 = *(*int32)(unsafe.Add(mBase, uint32(v5)+20))
																										v136 = *(*int32)(unsafe.Add(mBase, uint32(v135)+4))
																										v137 = *(*int32)(unsafe.Add(mBase, uint32(v136)))
																										v138 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v137)+7)))
																										v139 = m.G398
																										if base.Ui32(v138) < base.Ui32(int32(0)) {
																											v150 = v139
																										} else {
																											v150 = v137 + int32(8)
																										}
																										v151 = v150
																									}
																									v154 = *(*int32)(unsafe.Add(mBase, uint32(v5)+8))
																									v155 = *(*int64)(unsafe.Add(mBase, uint32(v151)))
																									*(*int64)(unsafe.Add(mBase, uint32(v154))) = v155
																									v157 = *(*int32)(unsafe.Add(mBase, uint32(v151)+8))
																									*(*int32)(unsafe.Add(mBase, uint32(v154)+8)) = v157
																									v159 = *(*int32)(unsafe.Add(mBase, uint32(v5)+8))
																									*(*int32)(unsafe.Add(mBase, uint32(v5)+8)) = v159 + int32(16)
																									F_luaSetErrorMetatable(m, v5)
																									mBase = m.M
																									v164 = m.ExcPending
																									if v164 != 0 {
																										return
																									} else {
																										F_luaSetTableProtectionRecursively(m, v5)
																										mBase = m.M
																										v166 = m.ExcPending
																										if v166 != 0 {
																											return
																										} else {
																											v185 = *(*int32)(unsafe.Add(mBase, uint32(v5)+8))
																											*(*int32)(unsafe.Add(mBase, uint32(v5)+8)) = v185 + int32(-16)
																											F_luaSetTableProtectionForBasicTypes(m, v5)
																											mBase = m.M
																											v196 = m.ExcPending
																											if v196 != 0 {
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
func F_luaEngineDebuggerEnable(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	if l2 != 0 {
		v21 = int32(0)
		return v21
	} else {
		v9 = m.G6
		*(*int32)(unsafe.Add(mBase, uint32(v9))) = int32(1)
		*(*int32)(unsafe.Add(mBase, uint32(v9)+268)) = int32(0)
		*(*int64)(unsafe.Add(mBase, uint32(v9)+260)) = int64(4294967296)
		F_ldbGenerateDebuggerCommandsArray(m, l3, l4)
		mBase = m.M
		v19 = m.ExcPending
		if v19 != 0 {
			return int32(0)
		} else {
			v21 = int32(1)
			return v21
		}
	}
}
func F_luaEngineDebuggerEnd(m *base.Module, l0 int32, l1 int32, l2 int32) {
	var v5 int32
	_ = v5
	F_ldbEnd(m)
	v5 = m.ExcPending
	if v5 != 0 {
		return
	} else {
		return
	}
}
func F_luaEngineDebuggerStart(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	var v6 int32
	_ = v6
	F_ldbStart(m, l3)
	v6 = m.ExcPending
	if v6 != 0 {
		return
	} else {
		return
	}
}
func F_luaEngineFreeFunction(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v12 int32
	_ = v12
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
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	if base.Ui32(int32(2)) <= base.Ui32(l2) {
		v44 = m.G3
		v50 = m.G8
		v51 = *(*int32)(unsafe.Add(mBase, uint32(v50)))
		m.T0[v51].(func(*base.Module, int32, int32, int32))(m, v44+int32(_a2022), v44+int32(_a2610), int32(409))
		mBase = m.M
		v53 = m.ExcPending
		if v53 != 0 {
			return
		} else {
			m.Env.Exit(m, int32(1))
			mBase = m.M
			base.Wasm_trap_unreachable()
			for {
			}
		}
	} else {
		v7 = int32(0)
		v12 = *(*int32)(unsafe.Add(mBase, uint32(l1+base.B2i32(l2 != v7)<<(uint(int32(2))%32))))
		if v12 == v7 {
			v56 = m.G3
			v62 = m.G8
			v63 = *(*int32)(unsafe.Add(mBase, uint32(v62)))
			m.T0[v63].(func(*base.Module, int32, int32, int32))(m, v56+int32(_a2316), v56+int32(_a2610), int32(413))
			mBase = m.M
			v65 = m.ExcPending
			if v65 != 0 {
				return
			} else {
				m.Env.Exit(m, int32(1))
				mBase = m.M
				base.Wasm_trap_unreachable()
				for {
				}
			}
		} else {
			v15 = *(*int32)(unsafe.Add(mBase, uint32(l3)+12))
			v16 = *(*int32)(unsafe.Add(mBase, uint32(v15)))
			if v12 != v16 {
				v22 = m.G11
				v23 = *(*int32)(unsafe.Add(mBase, uint32(v22)))
				m.T0[v23].(func(*base.Module, int32))(m, v15)
				mBase = m.M
				v25 = m.ExcPending
				if v25 != 0 {
					return
				} else {
					v26 = *(*int32)(unsafe.Add(mBase, uint32(l3)+8))
					if v26 == int32(0) {
						v33 = *(*int32)(unsafe.Add(mBase, uint32(l3)+16))
						if v33 == int32(0) {
							v40 = m.G11
							v41 = *(*int32)(unsafe.Add(mBase, uint32(v40)))
							m.T0[v41].(func(*base.Module, int32))(m, l3)
							mBase = m.M
							v43 = m.ExcPending
							if v43 != 0 {
								return
							} else {
								return
							}
						} else {
							v36 = m.G11
							v37 = *(*int32)(unsafe.Add(mBase, uint32(v36)))
							m.T0[v37].(func(*base.Module, int32))(m, v33)
							mBase = m.M
							v39 = m.ExcPending
							if v39 != 0 {
								return
							} else {
								v40 = m.G11
								v41 = *(*int32)(unsafe.Add(mBase, uint32(v40)))
								m.T0[v41].(func(*base.Module, int32))(m, l3)
								mBase = m.M
								v43 = m.ExcPending
								if v43 != 0 {
									return
								} else {
									return
								}
							}
						}
					} else {
						v29 = m.G11
						v30 = *(*int32)(unsafe.Add(mBase, uint32(v29)))
						m.T0[v30].(func(*base.Module, int32))(m, v26)
						mBase = m.M
						v32 = m.ExcPending
						if v32 != 0 {
							return
						} else {
							v33 = *(*int32)(unsafe.Add(mBase, uint32(l3)+16))
							if v33 == int32(0) {
								v40 = m.G11
								v41 = *(*int32)(unsafe.Add(mBase, uint32(v40)))
								m.T0[v41].(func(*base.Module, int32))(m, l3)
								mBase = m.M
								v43 = m.ExcPending
								if v43 != 0 {
									return
								} else {
									return
								}
							} else {
								v36 = m.G11
								v37 = *(*int32)(unsafe.Add(mBase, uint32(v36)))
								m.T0[v37].(func(*base.Module, int32))(m, v33)
								mBase = m.M
								v39 = m.ExcPending
								if v39 != 0 {
									return
								} else {
									v40 = m.G11
									v41 = *(*int32)(unsafe.Add(mBase, uint32(v40)))
									m.T0[v41].(func(*base.Module, int32))(m, l3)
									mBase = m.M
									v43 = m.ExcPending
									if v43 != 0 {
										return
									} else {
										return
									}
								}
							}
						}
					}
				}
			} else {
				v19 = *(*int32)(unsafe.Add(mBase, uint32(v15)+4))
				F_luaL_unref(m, v12, int32(-10000), v19)
				mBase = m.M
				v21 = m.ExcPending
				if v21 != 0 {
					return
				} else {
					v22 = m.G11
					v23 = *(*int32)(unsafe.Add(mBase, uint32(v22)))
					m.T0[v23].(func(*base.Module, int32))(m, v15)
					mBase = m.M
					v25 = m.ExcPending
					if v25 != 0 {
						return
					} else {
						v26 = *(*int32)(unsafe.Add(mBase, uint32(l3)+8))
						if v26 == int32(0) {
							v33 = *(*int32)(unsafe.Add(mBase, uint32(l3)+16))
							if v33 == int32(0) {
								v40 = m.G11
								v41 = *(*int32)(unsafe.Add(mBase, uint32(v40)))
								m.T0[v41].(func(*base.Module, int32))(m, l3)
								mBase = m.M
								v43 = m.ExcPending
								if v43 != 0 {
									return
								} else {
									return
								}
							} else {
								v36 = m.G11
								v37 = *(*int32)(unsafe.Add(mBase, uint32(v36)))
								m.T0[v37].(func(*base.Module, int32))(m, v33)
								mBase = m.M
								v39 = m.ExcPending
								if v39 != 0 {
									return
								} else {
									v40 = m.G11
									v41 = *(*int32)(unsafe.Add(mBase, uint32(v40)))
									m.T0[v41].(func(*base.Module, int32))(m, l3)
									mBase = m.M
									v43 = m.ExcPending
									if v43 != 0 {
										return
									} else {
										return
									}
								}
							}
						} else {
							v29 = m.G11
							v30 = *(*int32)(unsafe.Add(mBase, uint32(v29)))
							m.T0[v30].(func(*base.Module, int32))(m, v26)
							mBase = m.M
							v32 = m.ExcPending
							if v32 != 0 {
								return
							} else {
								v33 = *(*int32)(unsafe.Add(mBase, uint32(l3)+16))
								if v33 == int32(0) {
									v40 = m.G11
									v41 = *(*int32)(unsafe.Add(mBase, uint32(v40)))
									m.T0[v41].(func(*base.Module, int32))(m, l3)
									mBase = m.M
									v43 = m.ExcPending
									if v43 != 0 {
										return
									} else {
										return
									}
								} else {
									v36 = m.G11
									v37 = *(*int32)(unsafe.Add(mBase, uint32(v36)))
									m.T0[v37].(func(*base.Module, int32))(m, v33)
									mBase = m.M
									v39 = m.ExcPending
									if v39 != 0 {
										return
									} else {
										v40 = m.G11
										v41 = *(*int32)(unsafe.Add(mBase, uint32(v40)))
										m.T0[v41].(func(*base.Module, int32))(m, l3)
										mBase = m.M
										v43 = m.ExcPending
										if v43 != 0 {
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
func F_luaError(m *base.Module, l0 int32) int32 {
	var v2 int32
	_ = v2
	var v5 int32
	_ = v5
	v2 = F_lua_error(m, l0)
	v5 = m.ExcPending
	if v5 != 0 {
		return int32(0)
	} else {
		return v2
	}
}
func F_luaFunctionInitializeLuaState(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v53 int32
	_ = v53
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v110 int32
	_ = v110
	var v111 int64
	_ = v111
	var v113 int32
	_ = v113
	var v115 int32
	_ = v115
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v131 int32
	_ = v131
	var v155 int32
	_ = v155
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v188 int32
	_ = v188
	var v189 int64
	_ = v189
	var v191 int32
	_ = v191
	var v193 int32
	_ = v193
	var v201 int32
	_ = v201
	var v219 int32
	_ = v219
	var v258 int32
	_ = v258
	var v261 int32
	_ = v261
	var v262 int32
	_ = v262
	var v280 int32
	_ = v280
	var v319 int32
	_ = v319
	var v323 int32
	_ = v323
	var v325 int32
	_ = v325
	v3 = int32(0)
	F_lua_createtable(m, l1, v3, v3)
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return
	} else {
		v8 = int32(0)
		F_lua_createtable(m, l1, v8, v8)
		mBase = m.M
		v11 = m.ExcPending
		if v11 != 0 {
			return
		} else {
			v12 = m.G3
			F_lua_pushstring(m, l1, v12+int32(_a2622))
			mBase = m.M
			v16 = m.ExcPending
			if v16 != 0 {
				return
			} else {
				v17 = m.G5
				F_lua_pushcclosure(m, l1, v17+int32(1184), int32(0))
				mBase = m.M
				v22 = m.ExcPending
				if v22 != 0 {
					return
				} else {
					F_lua_settable(m, l1, int32(-3))
					mBase = m.M
					v25 = m.ExcPending
					if v25 != 0 {
						return
					} else {
						F_luaRegisterLogFunction(m, l1)
						mBase = m.M
						v27 = m.ExcPending
						if v27 != 0 {
							return
						} else {
							F_luaRegisterVersion(m, l0, l1)
							mBase = m.M
							v29 = m.ExcPending
							if v29 != 0 {
								return
							} else {
								F_luaSetErrorMetatable(m, l1)
								mBase = m.M
								v31 = m.ExcPending
								if v31 != 0 {
									return
								} else {
									v34 = v12 + int32(_a2614)
									F_lua_setfield(m, l1, int32(-2), v34)
									mBase = m.M
									v36 = m.ExcPending
									if v36 != 0 {
										return
									} else {
										F_lua_getfield(m, l1, int32(-1), v34)
										mBase = m.M
										v39 = m.ExcPending
										if v39 != 0 {
											return
										} else {
											F_lua_setfield(m, l1, int32(-2), v12+int32(_a2623))
											mBase = m.M
											v44 = m.ExcPending
											if v44 != 0 {
												return
											} else {
												F_luaSetErrorMetatable(m, l1)
												mBase = m.M
												v46 = m.ExcPending
												if v46 != 0 {
													return
												} else {
													F_luaSetTableProtectionRecursively(m, l1)
													mBase = m.M
													v48 = m.ExcPending
													if v48 != 0 {
														return
													} else {
														F_lua_setfield(m, l1, int32(-10000), v12+int32(_a2624))
														mBase = m.M
														v53 = m.ExcPending
														if v53 != 0 {
															return
														} else {
															switch int32(0) {
															case 0:
																v107 = l1 + int32(72)
															case 1:
																v80 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
																v81 = *(*int32)(unsafe.Add(mBase, uint32(v80)+4))
																v82 = *(*int32)(unsafe.Add(mBase, uint32(v81)))
																v83 = *(*int32)(unsafe.Add(mBase, uint32(v82)+12))
																*(*int32)(unsafe.Add(mBase, uint32(l1)+96)) = int32(5)
																*(*int32)(unsafe.Add(mBase, uint32(l1)+88)) = v83
																v107 = l1 + int32(88)
															case 2:
																v77 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
																v107 = v77 + int32(96)
															default:
																v91 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
																v92 = *(*int32)(unsafe.Add(mBase, uint32(v91)+4))
																v93 = *(*int32)(unsafe.Add(mBase, uint32(v92)))
																v94 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v93)+7)))
																v95 = m.G398
																if base.Ui32(v94) < base.Ui32(int32(0)) {
																	v106 = v95
																} else {
																	v106 = v93 + int32(8)
																}
																v107 = v106
															}
															v110 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
															v111 = *(*int64)(unsafe.Add(mBase, uint32(v107)))
															*(*int64)(unsafe.Add(mBase, uint32(v110))) = v111
															v113 = *(*int32)(unsafe.Add(mBase, uint32(v107)+8))
															*(*int32)(unsafe.Add(mBase, uint32(v110)+8)) = v113
															v115 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
															*(*int32)(unsafe.Add(mBase, uint32(l1)+8)) = v115 + int32(16)
															F_lua_setfield(m, l1, int32(-10000), v12+int32(_a2625))
															mBase = m.M
															v123 = m.ExcPending
															if v123 != 0 {
																return
															} else {
																v124 = int32(0)
																F_lua_createtable(m, l1, v124, v124)
																mBase = m.M
																v127 = m.ExcPending
																if v127 != 0 {
																	return
																} else {
																	v128 = int32(0)
																	F_lua_createtable(m, l1, v128, v128)
																	mBase = m.M
																	v131 = m.ExcPending
																	if v131 != 0 {
																		return
																	} else {
																		switch int32(0) {
																		case 0:
																			v185 = l1 + int32(72)
																		case 1:
																			v158 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
																			v159 = *(*int32)(unsafe.Add(mBase, uint32(v158)+4))
																			v160 = *(*int32)(unsafe.Add(mBase, uint32(v159)))
																			v161 = *(*int32)(unsafe.Add(mBase, uint32(v160)+12))
																			*(*int32)(unsafe.Add(mBase, uint32(l1)+96)) = int32(5)
																			*(*int32)(unsafe.Add(mBase, uint32(l1)+88)) = v161
																			v185 = l1 + int32(88)
																		case 2:
																			v155 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
																			v185 = v155 + int32(96)
																		default:
																			v169 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
																			v170 = *(*int32)(unsafe.Add(mBase, uint32(v169)+4))
																			v171 = *(*int32)(unsafe.Add(mBase, uint32(v170)))
																			v172 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v171)+7)))
																			v173 = m.G398
																			if base.Ui32(v172) < base.Ui32(int32(0)) {
																				v184 = v173
																			} else {
																				v184 = v171 + int32(8)
																			}
																			v185 = v184
																		}
																		v188 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
																		v189 = *(*int64)(unsafe.Add(mBase, uint32(v185)))
																		*(*int64)(unsafe.Add(mBase, uint32(v188))) = v189
																		v191 = *(*int32)(unsafe.Add(mBase, uint32(v185)+8))
																		*(*int32)(unsafe.Add(mBase, uint32(v188)+8)) = v191
																		v193 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
																		*(*int32)(unsafe.Add(mBase, uint32(l1)+8)) = v193 + int32(16)
																		F_lua_setfield(m, l1, int32(-2), v12+int32(_a2626))
																		mBase = m.M
																		v201 = m.ExcPending
																		if v201 != 0 {
																			return
																		} else {
																			v219 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
																			v258 = *(*int32)(unsafe.Add(mBase, uint32(v219+int32(-16))))
																			*(*int32)(unsafe.Add(mBase, uint32(v258)+8)) = int32(1)
																			v261 = F_lua_setmetatable(m, l1, int32(-2))
																			mBase = m.M
																			v262 = m.ExcPending
																			if v262 != 0 {
																				return
																			} else {
																				v280 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
																				v319 = *(*int32)(unsafe.Add(mBase, uint32(v280+int32(-16))))
																				*(*int32)(unsafe.Add(mBase, uint32(v319)+8)) = int32(1)
																				F_lua_replace(m, l1, int32(-10002))
																				mBase = m.M
																				v323 = m.ExcPending
																				if v323 != 0 {
																					return
																				} else {
																					F_luaSetTableProtectionForBasicTypes(m, l1)
																					mBase = m.M
																					v325 = m.ExcPending
																					if v325 != 0 {
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
func F_luaGetFromRegistry(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v83 int32
	_ = v83
	var v86 int32
	_ = v86
	var v120 int32
	_ = v120
	var v123 int32
	_ = v123
	var v126 int32
	_ = v126
	var v146 int32
	_ = v146
	var v149 int32
	_ = v149
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v188 int32
	_ = v188
	var v203 int32
	_ = v203
	var v206 int32
	_ = v206
	var v242 int32
	_ = v242
	var v245 int32
	_ = v245
	var v248 int32
	_ = v248
	var v252 int32
	_ = v252
	var v257 int32
	_ = v257
	var v260 int32
	_ = v260
	var v279 int32
	_ = v279
	var v290 int32
	_ = v290
	var v296 int32
	_ = v296
	var v297 int32
	_ = v297
	var v299 int32
	_ = v299
	var v302 int32
	_ = v302
	var v308 int32
	_ = v308
	var v309 int32
	_ = v309
	var v311 int32
	_ = v311
	F_lua_pushstring(m, l0, l1)
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		F_lua_gettable(m, l0, int32(-10000))
		mBase = m.M
		v9 = m.ExcPending
		if v9 != 0 {
			return int32(0)
		} else {
			v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			v27 = v24 + int32(-16)
			v61 = m.G398
			if v27 != v61 {
				v64 = *(*int32)(unsafe.Add(mBase, uint32(v27)+8))
				v67 = v64
			} else {
				v67 = int32(-1)
			}
			if v67 != 0 {
				v83 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				v86 = v83 + int32(-16)
				v120 = m.G398
				if v86 != v120 {
					v123 = *(*int32)(unsafe.Add(mBase, uint32(v86)+8))
					v126 = v123
				} else {
					v126 = int32(-1)
				}
				if v126 != int32(2) {
					v290 = m.G3
					v296 = m.G8
					v297 = *(*int32)(unsafe.Add(mBase, uint32(v296)))
					m.T0[v297].(func(*base.Module, int32, int32, int32))(m, v290+int32(_a2627), v290+int32(_a2628), int32(210))
					mBase = m.M
					v299 = m.ExcPending
					if v299 != 0 {
						return int32(0)
					} else {
						m.Env.Exit(m, int32(1))
						mBase = m.M
						base.Wasm_trap_unreachable()
						for {
						}
					}
				} else {
					v146 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
					v149 = v146 + int32(-16)
					v184 = int32(0)
					v185 = *(*int32)(unsafe.Add(mBase, uint32(v149)+8))
					switch v185 + int32(-2) {
					case 0, 5:
						v203 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
						v206 = v203 + int32(-16)
						v242 = *(*int32)(unsafe.Add(mBase, uint32(v206)+8))
						switch v242 + int32(-2) {
						case 0:
							v248 = *(*int32)(unsafe.Add(mBase, uint32(v206)))
							v252 = v248
							v257 = v252
						default:
							v252 = v184
							v257 = v252
						case 5:
							v245 = *(*int32)(unsafe.Add(mBase, uint32(v206)))
							v257 = v245 + int32(24)
						}
					default:
						v252 = v184
						v257 = v252
					case 3, 4, 6:
						v188 = *(*int32)(unsafe.Add(mBase, uint32(v149)))
						v257 = v188
					}
					if v257 == int32(0) {
						v302 = m.G3
						v308 = m.G8
						v309 = *(*int32)(unsafe.Add(mBase, uint32(v308)))
						m.T0[v309].(func(*base.Module, int32, int32, int32))(m, v302+int32(_a2629), v302+int32(_a2628), int32(213))
						mBase = m.M
						v311 = m.ExcPending
						if v311 != 0 {
							return int32(0)
						} else {
							m.Env.Exit(m, int32(1))
							mBase = m.M
							base.Wasm_trap_unreachable()
							for {
							}
						}
					} else {
						v260 = v257
						v279 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
						*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v279 + int32(-16)
						return v260
					}
				}
			} else {
				v260 = int32(0)
				v279 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v279 + int32(-16)
				return v260
			}
		}
	}
}
func F_luaPushErrorBuff(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v25 int32
	_ = v25
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
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v66 int32
	_ = v66
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v77 int32
	_ = v77
	var v83 int32
	_ = v83
	var v86 int32
	_ = v86
	var v92 int32
	_ = v92
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v106 int32
	_ = v106
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
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v118 int32
	_ = v118
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v126 int32
	_ = v126
	var v134 int32
	_ = v134
	var v137 int32
	_ = v137
	var v141 int32
	_ = v141
	var v146 int32
	_ = v146
	var v148 int32
	_ = v148
	var v152 int32
	_ = v152
	var v158 int32
	_ = v158
	var v161 int32
	_ = v161
	var v167 int32
	_ = v167
	var v171 int32
	_ = v171
	var v173 int32
	_ = v173
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
	var v187 int32
	_ = v187
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v199 int32
	_ = v199
	var v202 int32
	_ = v202
	var v206 int32
	_ = v206
	var v211 int32
	_ = v211
	var v213 int32
	_ = v213
	var v217 int32
	_ = v217
	var v223 int32
	_ = v223
	var v226 int32
	_ = v226
	var v232 int32
	_ = v232
	var v236 int32
	_ = v236
	var v238 int32
	_ = v238
	var v246 int32
	_ = v246
	var v249 int32
	_ = v249
	var v254 int32
	_ = v254
	var v259 int32
	_ = v259
	var v269 int32
	_ = v269
	var v274 int32
	_ = v274
	var v284 int32
	_ = v284
	var v288 int32
	_ = v288
	var v298 int32
	_ = v298
	var v304 int32
	_ = v304
	var v310 int32
	_ = v310
	var v315 int32
	_ = v315
	var v322 int32
	_ = v322
	var v325 int32
	_ = v325
	var v332 int32
	_ = v332
	var v333 int32
	_ = v333
	var v334 int32
	_ = v334
	var v338 int32
	_ = v338
	var v340 int32
	_ = v340
	var v341 int32
	_ = v341
	var v343 int32
	_ = v343
	var v345 int32
	_ = v345
	var v356 int32
	_ = v356
	var v362 int32
	_ = v362
	var v363 int32
	_ = v363
	var v365 int32
	_ = v365
	var v371 int32
	_ = v371
	var v378 int32
	_ = v378
	var v382 int32
	_ = v382
	var v385 int32
	_ = v385
	var v391 int32
	_ = v391
	var v398 int32
	_ = v398
	var v402 int32
	_ = v402
	var v405 int32
	_ = v405
	var v408 int32
	_ = v408
	var v409 int32
	_ = v409
	var v410 int32
	_ = v410
	var v415 int32
	_ = v415
	var v416 int32
	_ = v416
	var v417 int32
	_ = v417
	var v419 int32
	_ = v419
	var v421 int32
	_ = v421
	var v422 int32
	_ = v422
	var v424 int32
	_ = v424
	var v426 int32
	_ = v426
	var v430 int32
	_ = v430
	var v431 int32
	_ = v431
	var v432 int32
	_ = v432
	var v437 int32
	_ = v437
	var v438 int32
	_ = v438
	var v439 int32
	_ = v439
	var v441 int32
	_ = v441
	var v443 int32
	_ = v443
	var v448 int32
	_ = v448
	var v461 int32
	_ = v461
	var v467 int32
	_ = v467
	var v472 int32
	_ = v472
	var v473 int32
	_ = v473
	var v481 int32
	_ = v481
	var v484 int32
	_ = v484
	var v488 int32
	_ = v488
	var v493 int32
	_ = v493
	var v495 int32
	_ = v495
	var v499 int32
	_ = v499
	var v505 int32
	_ = v505
	var v508 int32
	_ = v508
	var v514 int32
	_ = v514
	var v518 int32
	_ = v518
	var v520 int32
	_ = v520
	var v528 int32
	_ = v528
	var v530 int32
	_ = v530
	var v531 int32
	_ = v531
	var v532 int32
	_ = v532
	var v533 int32
	_ = v533
	var v534 int32
	_ = v534
	var v537 int32
	_ = v537
	var v538 int32
	_ = v538
	var v546 int32
	_ = v546
	var v549 int32
	_ = v549
	var v553 int32
	_ = v553
	var v558 int32
	_ = v558
	var v560 int32
	_ = v560
	var v564 int32
	_ = v564
	var v570 int32
	_ = v570
	var v573 int32
	_ = v573
	var v579 int32
	_ = v579
	var v583 int32
	_ = v583
	var v585 int32
	_ = v585
	var v593 int32
	_ = v593
	var v596 int32
	_ = v596
	var v601 int32
	_ = v601
	var v606 int32
	_ = v606
	var v616 int32
	_ = v616
	var v621 int32
	_ = v621
	var v631 int32
	_ = v631
	var v635 int32
	_ = v635
	var v645 int32
	_ = v645
	var v651 int32
	_ = v651
	var v657 int32
	_ = v657
	var v662 int32
	_ = v662
	var v669 int32
	_ = v669
	var v672 int32
	_ = v672
	var v679 int32
	_ = v679
	var v680 int32
	_ = v680
	var v681 int32
	_ = v681
	var v685 int32
	_ = v685
	var v687 int32
	_ = v687
	var v688 int32
	_ = v688
	var v690 int32
	_ = v690
	var v692 int32
	_ = v692
	var v703 int32
	_ = v703
	var v709 int32
	_ = v709
	var v710 int32
	_ = v710
	var v712 int32
	_ = v712
	var v718 int32
	_ = v718
	var v725 int32
	_ = v725
	var v729 int32
	_ = v729
	var v732 int32
	_ = v732
	var v738 int32
	_ = v738
	var v745 int32
	_ = v745
	var v749 int32
	_ = v749
	var v752 int32
	_ = v752
	var v755 int32
	_ = v755
	var v756 int32
	_ = v756
	var v757 int32
	_ = v757
	var v762 int32
	_ = v762
	var v763 int32
	_ = v763
	var v764 int32
	_ = v764
	var v766 int32
	_ = v766
	var v768 int32
	_ = v768
	var v769 int32
	_ = v769
	var v771 int32
	_ = v771
	var v773 int32
	_ = v773
	var v777 int32
	_ = v777
	var v778 int32
	_ = v778
	var v779 int32
	_ = v779
	var v784 int32
	_ = v784
	var v785 int32
	_ = v785
	var v786 int32
	_ = v786
	var v788 int32
	_ = v788
	var v790 int32
	_ = v790
	var v795 int32
	_ = v795
	var v808 int32
	_ = v808
	var v811 int32
	_ = v811
	var v816 int32
	_ = v816
	var v817 int32
	_ = v817
	var v821 int32
	_ = v821
	var v823 int32
	_ = v823
	var v826 int32
	_ = v826
	var v829 int32
	_ = v829
	var v830 int32
	_ = v830
	var v834 int32
	_ = v834
	var v836 int32
	_ = v836
	var v839 int32
	_ = v839
	var v840 int32
	_ = v840
	var v841 int32
	_ = v841
	var v843 int32
	_ = v843
	var v844 int32
	_ = v844
	var v846 int32
	_ = v846
	v3 = int32(0)
	v9 = m.G0
	v11 = v9 - int32(64)
	m.G0 = v11
	v14 = m.G6
	v15 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
	v18 = *(*int32)(unsafe.Add(mBase, uint32(v14)+264))
	goto L2
L1:
	;
	v39 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	if v39 != int32(45) {
		goto L9
	} else {
		goto L10
	}
L2:
	;
	if base.B2i32(v15 != v3)&base.B2i32(v18 != v3) == int32(0) {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+48)) = l1
	v25 = m.G3
	v30 = F_lm_asprintf(m, v25+int32(_a2630), v11+int32(48))
	mBase = m.M
	v31 = m.ExcPending
	if v31 != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	return
L5:
	;
	F_ldbLogCString(m, v30)
	mBase = m.M
	v33 = m.ExcPending
	if v33 != 0 {
		goto L4
	} else {
		goto L6
	}
L6:
	;
	v34 = m.G11
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v34)))
	m.T0[v35].(func(*base.Module, int32))(m, v30)
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		goto L4
	} else {
		goto L7
	}
L7:
	;
	goto L1
L8:
	;
	v826 = int32(0)
	F_lua_createtable(m, l0, v826, v826)
	mBase = m.M
	v829 = m.ExcPending
	if v829 != 0 {
		goto L4
	} else {
		goto L226
	}
L9:
	;
	if l1&int32(3) == int32(0) {
		v495 = l1
		goto L134
	} else {
		goto L135
	}
L10:
	;
	v42 = int32(32)
	v43 = F___strchrnul(m, l1, v42)
	mBase = m.M
	v45 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v43))))
	if v45 == v42 {
		goto L13
	} else {
		goto L14
	}
L11:
	;
	v123 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v49))) = uint8(v123)
	v126 = v49 + int32(1)
	if v126&int32(3) == v123 {
		v148 = v126
		goto L40
	} else {
		goto L41
	}
L12:
	;
	if v49 != 0 {
		goto L11
	} else {
		goto L16
	}
L13:
	;
	v49 = v43
	goto L15
L14:
	;
	v49 = int32(0)
	goto L15
L15:
	;
	goto L12
L16:
	;
	v51 = l1 + int32(1)
	if v51&int32(3) == int32(0) {
		v73 = v51
		goto L19
	} else {
		goto L20
	}
L17:
	;
	v108 = v106 + int32(1)
	v109 = m.G22
	v110 = *(*int32)(unsafe.Add(mBase, uint32(v109)))
	v111 = m.T0[v110].(func(*base.Module, int32) int32)(m, v108)
	mBase = m.M
	v112 = m.ExcPending
	if v112 != 0 {
		goto L4
	} else {
		goto L33
	}
L18:
	;
	v106 = v98 - v51
	goto L17
L19:
	;
	v77 = v73
	goto L27
L20:
	;
	v59 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v51))))
	if v59 != 0 {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	v62 = v51
	goto L23
L22:
	;
	v106 = v51 - v51
	goto L17
L23:
	;
	v66 = v62 + int32(1)
	if v66&int32(3) == int32(0) {
		v73 = v66
		goto L19
	} else {
		goto L25
	}
L25:
	;
	v71 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v66))))
	if v71 != 0 {
		v62 = v66
		goto L23
	} else {
		goto L26
	}
L26:
	;
	v98 = v66
	goto L18
L27:
	;
	v83 = *(*int32)(unsafe.Add(mBase, uint32(v77)))
	v86 = int32(-2139062144)
	if (int32(16843008)-v83|v83)&v86 == v86 {
		v77 = v77 + int32(4)
		goto L27
	} else {
		goto L29
	}
L28:
	;
	v92 = v77
	goto L30
L29:
	;
	goto L28
L30:
	;
	v96 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v92))))
	if v96 != 0 {
		v92 = v92 + int32(1)
		goto L30
	} else {
		goto L32
	}
L31:
	;
	v98 = v92
	goto L18
L32:
	;
	goto L31
L33:
	;
	if v108 == int32(0) {
		v116 = v111
		goto L35
	} else {
		goto L36
	}
L34:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11))) = v116
	v118 = m.G3
	v121 = F_lm_asprintf(m, v118+int32(_a2631), v11)
	mBase = m.M
	v122 = m.ExcPending
	if v122 != 0 {
		goto L4
	} else {
		goto L37
	}
L35:
	;
	goto L34
L36:
	;
	v115 = F__emscripten_memcpy_bulkmem(m, v111, v51, v108)
	mBase = m.M
	v116 = v115
	goto L35
L37:
	;
	v821 = v121
	v823 = v111
	goto L8
L38:
	;
	v183 = v181 + int32(1)
	v184 = m.G22
	v185 = *(*int32)(unsafe.Add(mBase, uint32(v184)))
	v186 = m.T0[v185].(func(*base.Module, int32) int32)(m, v183)
	mBase = m.M
	v187 = m.ExcPending
	if v187 != 0 {
		goto L4
	} else {
		goto L54
	}
L39:
	;
	v181 = v173 - v126
	goto L38
L40:
	;
	v152 = v148
	goto L48
L41:
	;
	v134 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v126))))
	if v134 != 0 {
		goto L42
	} else {
		goto L43
	}
L42:
	;
	v137 = v126
	goto L44
L43:
	;
	v181 = v126 - v126
	goto L38
L44:
	;
	v141 = v137 + int32(1)
	if v141&int32(3) == int32(0) {
		v148 = v141
		goto L40
	} else {
		goto L46
	}
L46:
	;
	v146 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v141))))
	if v146 != 0 {
		v137 = v141
		goto L44
	} else {
		goto L47
	}
L47:
	;
	v173 = v141
	goto L39
L48:
	;
	v158 = *(*int32)(unsafe.Add(mBase, uint32(v152)))
	v161 = int32(-2139062144)
	if (int32(16843008)-v158|v158)&v161 == v161 {
		v152 = v152 + int32(4)
		goto L48
	} else {
		goto L50
	}
L49:
	;
	v167 = v152
	goto L51
L50:
	;
	goto L49
L51:
	;
	v171 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v167))))
	if v171 != 0 {
		v167 = v167 + int32(1)
		goto L51
	} else {
		goto L53
	}
L52:
	;
	v173 = v167
	goto L39
L53:
	;
	goto L52
L54:
	;
	if v183 == int32(0) {
		v191 = v186
		goto L56
	} else {
		goto L57
	}
L55:
	;
	if v191&int32(3) == int32(0) {
		v213 = v191
		goto L61
	} else {
		goto L62
	}
L56:
	;
	goto L55
L57:
	;
	v190 = F__emscripten_memcpy_bulkmem(m, v186, v126, v183)
	mBase = m.M
	v191 = v190
	goto L56
L58:
	;
	if base.Ui32(v249) <= base.Ui32(v274) {
		v304 = v249
		goto L81
	} else {
		goto L82
	}
L59:
	;
	v249 = v191 + v246 + int32(-1)
	if base.Ui32(v249) < base.Ui32(v191) {
		v274 = v191
		goto L58
	} else {
		goto L75
	}
L60:
	;
	v246 = v238 - v191
	goto L59
L61:
	;
	v217 = v213
	goto L69
L62:
	;
	v199 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v191))))
	if v199 != 0 {
		goto L63
	} else {
		goto L64
	}
L63:
	;
	v202 = v191
	goto L65
L64:
	;
	v246 = v191 - v191
	goto L59
L65:
	;
	v206 = v202 + int32(1)
	if v206&int32(3) == int32(0) {
		v213 = v206
		goto L61
	} else {
		goto L67
	}
L67:
	;
	v211 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v206))))
	if v211 != 0 {
		v202 = v206
		goto L65
	} else {
		goto L68
	}
L68:
	;
	v238 = v206
	goto L60
L69:
	;
	v223 = *(*int32)(unsafe.Add(mBase, uint32(v217)))
	v226 = int32(-2139062144)
	if (int32(16843008)-v223|v223)&v226 == v226 {
		v217 = v217 + int32(4)
		goto L69
	} else {
		goto L71
	}
L70:
	;
	v232 = v217
	goto L72
L71:
	;
	goto L70
L72:
	;
	v236 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v232))))
	if v236 != 0 {
		v232 = v232 + int32(1)
		goto L72
	} else {
		goto L74
	}
L73:
	;
	v238 = v232
	goto L60
L74:
	;
	goto L73
L75:
	;
	v254 = v191
	goto L76
L76:
	;
	v259 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v254))))
	if base.Ui32(int32(15)) < base.Ui32(v259) {
		v274 = v254
		goto L58
	} else {
		goto L78
	}
L77:
	;
	v274 = v269
	goto L58
L78:
	;
	if int32(1)<<(uint(v259)%32)&int32(9217) == int32(0) {
		v274 = v254
		goto L58
	} else {
		goto L79
	}
L79:
	;
	v269 = v254 + int32(1)
	if base.Ui32(v269) <= base.Ui32(v249) {
		v254 = v269
		goto L76
	} else {
		goto L80
	}
L80:
	;
	goto L77
L81:
	;
	v310 = v304 - v274 + int32(1)
	if v191 == v274 {
		goto L88
	} else {
		goto L89
	}
L82:
	;
	v284 = v249
	goto L83
L83:
	;
	v288 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v284))))
	if base.Ui32(int32(15)) < base.Ui32(v288) {
		v304 = v284
		goto L81
	} else {
		goto L85
	}
L84:
	;
	v304 = v274
	goto L81
L85:
	;
	if int32(1)<<(uint(v288)%32)&int32(9217) == int32(0) {
		v304 = v284
		goto L81
	} else {
		goto L86
	}
L86:
	;
	v298 = v284 + int32(-1)
	if base.Ui32(v274) < base.Ui32(v298) {
		v284 = v298
		goto L83
	} else {
		goto L87
	}
L87:
	;
	goto L84
L88:
	;
	v461 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v191+v310))) = uint8(v461)
	*(*int32)(unsafe.Add(mBase, uint32(v11)+20)) = v191
	*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = l1 + int32(1)
	v467 = m.G3
	v472 = F_lm_asprintf(m, v467+int32(_a518), v11+int32(16))
	mBase = m.M
	v473 = m.ExcPending
	if v473 != 0 {
		goto L4
	} else {
		goto L131
	}
L89:
	;
	if v191 == v274 {
		goto L91
	} else {
		goto L92
	}
L90:
	;
	goto L88
L91:
	;
	goto L90
L92:
	;
	v315 = v310 + v191
	if base.Ui32(int32(0)-v310<<(uint(int32(1))%32)) < base.Ui32(v274-v315) {
		goto L93
	} else {
		goto L94
	}
L93:
	;
	v325 = (v274 ^ v191) & int32(3)
	if base.Ui32(v274) <= base.Ui32(v191) {
		goto L97
	} else {
		goto L98
	}
L94:
	;
	v322 = F___memcpy(m, v191, v274, v310)
	mBase = m.M
	goto L90
L95:
	;
	if v431 == int32(0) {
		goto L91
	} else {
		goto L127
	}
L96:
	;
	if base.Ui32(v409) <= base.Ui32(int32(3)) {
		v430 = v408
		v431 = v409
		v432 = v410
		goto L95
	} else {
		goto L123
	}
L97:
	;
	if v325 != 0 {
		v391 = v310
		goto L107
	} else {
		goto L108
	}
L98:
	;
	if v325 == int32(0) {
		goto L99
	} else {
		goto L100
	}
L99:
	;
	if v191&int32(3) != 0 {
		goto L101
	} else {
		goto L102
	}
L100:
	;
	v430 = v274
	v431 = v310
	v432 = v191
	goto L95
L101:
	;
	v332 = v274
	v333 = v310
	v334 = v191
	goto L103
L102:
	;
	v408 = v274
	v409 = v310
	v410 = v191
	goto L96
L103:
	;
	if v333 == int32(0) {
		goto L91
	} else {
		goto L105
	}
L105:
	;
	v338 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v332))))
	*(*uint8)(unsafe.Add(mBase, uint32(v334))) = uint8(v338)
	v340 = int32(1)
	v341 = v332 + v340
	v343 = v333 + int32(-1)
	v345 = v334 + v340
	if v345&int32(3) == int32(0) {
		v408 = v341
		v409 = v343
		v410 = v345
		goto L96
	} else {
		goto L106
	}
L106:
	;
	v332 = v341
	v333 = v343
	v334 = v345
	goto L103
L107:
	;
	if v391 == int32(0) {
		goto L91
	} else {
		goto L119
	}
L108:
	;
	if v315&int32(3) == int32(0) {
		v371 = v310
		goto L109
	} else {
		goto L110
	}
L109:
	;
	if base.Ui32(v371) <= base.Ui32(int32(3)) {
		v391 = v371
		goto L107
	} else {
		goto L115
	}
L110:
	;
	v356 = v310
	goto L111
L111:
	;
	if v356 == int32(0) {
		goto L91
	} else {
		goto L113
	}
L112:
	;
	v371 = v362
	goto L109
L113:
	;
	v362 = v356 + int32(-1)
	v363 = v191 + v362
	v365 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v274+v362))))
	*(*uint8)(unsafe.Add(mBase, uint32(v363))) = uint8(v365)
	if v363&int32(3) != 0 {
		v356 = v362
		goto L111
	} else {
		goto L114
	}
L114:
	;
	goto L112
L115:
	;
	v378 = v371
	goto L116
L116:
	;
	v382 = v378 + int32(-4)
	v385 = *(*int32)(unsafe.Add(mBase, uint32(v274+v382)))
	*(*int32)(unsafe.Add(mBase, uint32(v191+v382))) = v385
	if base.Ui32(int32(3)) < base.Ui32(v382) {
		v378 = v382
		goto L116
	} else {
		goto L118
	}
L117:
	;
	v391 = v382
	goto L107
L118:
	;
	goto L117
L119:
	;
	v398 = v391
	goto L120
L120:
	;
	v402 = v398 + int32(-1)
	v405 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v274+v402))))
	*(*uint8)(unsafe.Add(mBase, uint32(v191+v402))) = uint8(v405)
	if v402 != 0 {
		v398 = v402
		goto L120
	} else {
		goto L122
	}
L122:
	;
	goto L91
L123:
	;
	v415 = v408
	v416 = v409
	v417 = v410
	goto L124
L124:
	;
	v419 = *(*int32)(unsafe.Add(mBase, uint32(v415)))
	*(*int32)(unsafe.Add(mBase, uint32(v417))) = v419
	v421 = int32(4)
	v422 = v415 + v421
	v424 = v417 + v421
	v426 = v416 + int32(-4)
	if base.Ui32(int32(3)) < base.Ui32(v426) {
		v415 = v422
		v416 = v426
		v417 = v424
		goto L124
	} else {
		goto L126
	}
L125:
	;
	v430 = v422
	v431 = v426
	v432 = v424
	goto L95
L126:
	;
	goto L125
L127:
	;
	v437 = v430
	v438 = v431
	v439 = v432
	goto L128
L128:
	;
	v441 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v437))))
	*(*uint8)(unsafe.Add(mBase, uint32(v439))) = uint8(v441)
	v443 = int32(1)
	v448 = v438 + int32(-1)
	if v448 != 0 {
		v437 = v437 + v443
		v438 = v448
		v439 = v439 + v443
		goto L128
	} else {
		goto L130
	}
L129:
	;
	goto L91
L130:
	;
	goto L129
L131:
	;
	v821 = v472
	v823 = v186
	goto L8
L132:
	;
	v530 = v528 + int32(1)
	v531 = m.G22
	v532 = *(*int32)(unsafe.Add(mBase, uint32(v531)))
	v533 = m.T0[v532].(func(*base.Module, int32) int32)(m, v530)
	mBase = m.M
	v534 = m.ExcPending
	if v534 != 0 {
		goto L4
	} else {
		goto L148
	}
L133:
	;
	v528 = v520 - l1
	goto L132
L134:
	;
	v499 = v495
	goto L142
L135:
	;
	v481 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	if v481 != 0 {
		goto L136
	} else {
		goto L137
	}
L136:
	;
	v484 = l1
	goto L138
L137:
	;
	v528 = l1 - l1
	goto L132
L138:
	;
	v488 = v484 + int32(1)
	if v488&int32(3) == int32(0) {
		v495 = v488
		goto L134
	} else {
		goto L140
	}
L140:
	;
	v493 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v488))))
	if v493 != 0 {
		v484 = v488
		goto L138
	} else {
		goto L141
	}
L141:
	;
	v520 = v488
	goto L133
L142:
	;
	v505 = *(*int32)(unsafe.Add(mBase, uint32(v499)))
	v508 = int32(-2139062144)
	if (int32(16843008)-v505|v505)&v508 == v508 {
		v499 = v499 + int32(4)
		goto L142
	} else {
		goto L144
	}
L143:
	;
	v514 = v499
	goto L145
L144:
	;
	goto L143
L145:
	;
	v518 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v514))))
	if v518 != 0 {
		v514 = v514 + int32(1)
		goto L145
	} else {
		goto L147
	}
L146:
	;
	v520 = v514
	goto L133
L147:
	;
	goto L146
L148:
	;
	if v530 == int32(0) {
		v538 = v533
		goto L150
	} else {
		goto L151
	}
L149:
	;
	if v538&int32(3) == int32(0) {
		v560 = v538
		goto L155
	} else {
		goto L156
	}
L150:
	;
	goto L149
L151:
	;
	v537 = F__emscripten_memcpy_bulkmem(m, v533, l1, v530)
	mBase = m.M
	v538 = v537
	goto L150
L152:
	;
	if base.Ui32(v596) <= base.Ui32(v621) {
		v651 = v596
		goto L175
	} else {
		goto L176
	}
L153:
	;
	v596 = v538 + v593 + int32(-1)
	if base.Ui32(v596) < base.Ui32(v538) {
		v621 = v538
		goto L152
	} else {
		goto L169
	}
L154:
	;
	v593 = v585 - v538
	goto L153
L155:
	;
	v564 = v560
	goto L163
L156:
	;
	v546 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v538))))
	if v546 != 0 {
		goto L157
	} else {
		goto L158
	}
L157:
	;
	v549 = v538
	goto L159
L158:
	;
	v593 = v538 - v538
	goto L153
L159:
	;
	v553 = v549 + int32(1)
	if v553&int32(3) == int32(0) {
		v560 = v553
		goto L155
	} else {
		goto L161
	}
L161:
	;
	v558 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v553))))
	if v558 != 0 {
		v549 = v553
		goto L159
	} else {
		goto L162
	}
L162:
	;
	v585 = v553
	goto L154
L163:
	;
	v570 = *(*int32)(unsafe.Add(mBase, uint32(v564)))
	v573 = int32(-2139062144)
	if (int32(16843008)-v570|v570)&v573 == v573 {
		v564 = v564 + int32(4)
		goto L163
	} else {
		goto L165
	}
L164:
	;
	v579 = v564
	goto L166
L165:
	;
	goto L164
L166:
	;
	v583 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v579))))
	if v583 != 0 {
		v579 = v579 + int32(1)
		goto L166
	} else {
		goto L168
	}
L167:
	;
	v585 = v579
	goto L154
L168:
	;
	goto L167
L169:
	;
	v601 = v538
	goto L170
L170:
	;
	v606 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v601))))
	if base.Ui32(int32(15)) < base.Ui32(v606) {
		v621 = v601
		goto L152
	} else {
		goto L172
	}
L171:
	;
	v621 = v616
	goto L152
L172:
	;
	if int32(1)<<(uint(v606)%32)&int32(9217) == int32(0) {
		v621 = v601
		goto L152
	} else {
		goto L173
	}
L173:
	;
	v616 = v601 + int32(1)
	if base.Ui32(v616) <= base.Ui32(v596) {
		v601 = v616
		goto L170
	} else {
		goto L174
	}
L174:
	;
	goto L171
L175:
	;
	v657 = v651 - v621 + int32(1)
	if v538 == v621 {
		goto L182
	} else {
		goto L183
	}
L176:
	;
	v631 = v596
	goto L177
L177:
	;
	v635 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v631))))
	if base.Ui32(int32(15)) < base.Ui32(v635) {
		v651 = v631
		goto L175
	} else {
		goto L179
	}
L178:
	;
	v651 = v621
	goto L175
L179:
	;
	if int32(1)<<(uint(v635)%32)&int32(9217) == int32(0) {
		v651 = v631
		goto L175
	} else {
		goto L180
	}
L180:
	;
	v645 = v631 + int32(-1)
	if base.Ui32(v621) < base.Ui32(v645) {
		v631 = v645
		goto L177
	} else {
		goto L181
	}
L181:
	;
	goto L178
L182:
	;
	v808 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v538+v657))) = uint8(v808)
	*(*int32)(unsafe.Add(mBase, uint32(v11)+32)) = v538
	v811 = m.G3
	v816 = F_lm_asprintf(m, v811+int32(_a2631), v11+int32(32))
	mBase = m.M
	v817 = m.ExcPending
	if v817 != 0 {
		goto L4
	} else {
		goto L225
	}
L183:
	;
	if v538 == v621 {
		goto L185
	} else {
		goto L186
	}
L184:
	;
	goto L182
L185:
	;
	goto L184
L186:
	;
	v662 = v657 + v538
	if base.Ui32(int32(0)-v657<<(uint(int32(1))%32)) < base.Ui32(v621-v662) {
		goto L187
	} else {
		goto L188
	}
L187:
	;
	v672 = (v621 ^ v538) & int32(3)
	if base.Ui32(v621) <= base.Ui32(v538) {
		goto L191
	} else {
		goto L192
	}
L188:
	;
	v669 = F___memcpy(m, v538, v621, v657)
	mBase = m.M
	goto L184
L189:
	;
	if v778 == int32(0) {
		goto L185
	} else {
		goto L221
	}
L190:
	;
	if base.Ui32(v756) <= base.Ui32(int32(3)) {
		v777 = v755
		v778 = v756
		v779 = v757
		goto L189
	} else {
		goto L217
	}
L191:
	;
	if v672 != 0 {
		v738 = v657
		goto L201
	} else {
		goto L202
	}
L192:
	;
	if v672 == int32(0) {
		goto L193
	} else {
		goto L194
	}
L193:
	;
	if v538&int32(3) != 0 {
		goto L195
	} else {
		goto L196
	}
L194:
	;
	v777 = v621
	v778 = v657
	v779 = v538
	goto L189
L195:
	;
	v679 = v621
	v680 = v657
	v681 = v538
	goto L197
L196:
	;
	v755 = v621
	v756 = v657
	v757 = v538
	goto L190
L197:
	;
	if v680 == int32(0) {
		goto L185
	} else {
		goto L199
	}
L199:
	;
	v685 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v679))))
	*(*uint8)(unsafe.Add(mBase, uint32(v681))) = uint8(v685)
	v687 = int32(1)
	v688 = v679 + v687
	v690 = v680 + int32(-1)
	v692 = v681 + v687
	if v692&int32(3) == int32(0) {
		v755 = v688
		v756 = v690
		v757 = v692
		goto L190
	} else {
		goto L200
	}
L200:
	;
	v679 = v688
	v680 = v690
	v681 = v692
	goto L197
L201:
	;
	if v738 == int32(0) {
		goto L185
	} else {
		goto L213
	}
L202:
	;
	if v662&int32(3) == int32(0) {
		v718 = v657
		goto L203
	} else {
		goto L204
	}
L203:
	;
	if base.Ui32(v718) <= base.Ui32(int32(3)) {
		v738 = v718
		goto L201
	} else {
		goto L209
	}
L204:
	;
	v703 = v657
	goto L205
L205:
	;
	if v703 == int32(0) {
		goto L185
	} else {
		goto L207
	}
L206:
	;
	v718 = v709
	goto L203
L207:
	;
	v709 = v703 + int32(-1)
	v710 = v538 + v709
	v712 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v621+v709))))
	*(*uint8)(unsafe.Add(mBase, uint32(v710))) = uint8(v712)
	if v710&int32(3) != 0 {
		v703 = v709
		goto L205
	} else {
		goto L208
	}
L208:
	;
	goto L206
L209:
	;
	v725 = v718
	goto L210
L210:
	;
	v729 = v725 + int32(-4)
	v732 = *(*int32)(unsafe.Add(mBase, uint32(v621+v729)))
	*(*int32)(unsafe.Add(mBase, uint32(v538+v729))) = v732
	if base.Ui32(int32(3)) < base.Ui32(v729) {
		v725 = v729
		goto L210
	} else {
		goto L212
	}
L211:
	;
	v738 = v729
	goto L201
L212:
	;
	goto L211
L213:
	;
	v745 = v738
	goto L214
L214:
	;
	v749 = v745 + int32(-1)
	v752 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v621+v749))))
	*(*uint8)(unsafe.Add(mBase, uint32(v538+v749))) = uint8(v752)
	if v749 != 0 {
		v745 = v749
		goto L214
	} else {
		goto L216
	}
L216:
	;
	goto L185
L217:
	;
	v762 = v755
	v763 = v756
	v764 = v757
	goto L218
L218:
	;
	v766 = *(*int32)(unsafe.Add(mBase, uint32(v762)))
	*(*int32)(unsafe.Add(mBase, uint32(v764))) = v766
	v768 = int32(4)
	v769 = v762 + v768
	v771 = v764 + v768
	v773 = v763 + int32(-4)
	if base.Ui32(int32(3)) < base.Ui32(v773) {
		v762 = v769
		v763 = v773
		v764 = v771
		goto L218
	} else {
		goto L220
	}
L219:
	;
	v777 = v769
	v778 = v773
	v779 = v771
	goto L189
L220:
	;
	goto L219
L221:
	;
	v784 = v777
	v785 = v778
	v786 = v779
	goto L222
L222:
	;
	v788 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v784))))
	*(*uint8)(unsafe.Add(mBase, uint32(v786))) = uint8(v788)
	v790 = int32(1)
	v795 = v785 + int32(-1)
	if v795 != 0 {
		v784 = v784 + v790
		v785 = v795
		v786 = v786 + v790
		goto L222
	} else {
		goto L224
	}
L223:
	;
	goto L185
L224:
	;
	goto L223
L225:
	;
	v821 = v816
	v823 = v533
	goto L8
L226:
	;
	v830 = m.G3
	F_lua_pushstring(m, l0, v830+int32(_a2632))
	mBase = m.M
	v834 = m.ExcPending
	if v834 != 0 {
		goto L4
	} else {
		goto L227
	}
L227:
	;
	F_lua_pushstring(m, l0, v821)
	mBase = m.M
	v836 = m.ExcPending
	if v836 != 0 {
		goto L4
	} else {
		goto L228
	}
L228:
	;
	F_lua_settable(m, l0, int32(-3))
	mBase = m.M
	v839 = m.ExcPending
	if v839 != 0 {
		goto L4
	} else {
		goto L229
	}
L229:
	;
	v840 = m.G11
	v841 = *(*int32)(unsafe.Add(mBase, uint32(v840)))
	m.T0[v841].(func(*base.Module, int32))(m, v823)
	mBase = m.M
	v843 = m.ExcPending
	if v843 != 0 {
		goto L4
	} else {
		goto L230
	}
L230:
	;
	v844 = *(*int32)(unsafe.Add(mBase, uint32(v840)))
	m.T0[v844].(func(*base.Module, int32))(m, v821)
	mBase = m.M
	v846 = m.ExcPending
	if v846 != 0 {
		goto L4
	} else {
		goto L231
	}
L231:
	;
	m.G0 = v11 + int32(64)
	return
}
func F_luaRedisPcall(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v10 int32
	_ = v10
	var v16 int32
	_ = v16
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v83 int32
	_ = v83
	var v84 int64
	_ = v84
	var v88 int32
	_ = v88
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v96 int64
	_ = v96
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v104 int32
	_ = v104
	var v107 int32
	_ = v107
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v124 int32
	_ = v124
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v171 int32
	_ = v171
	var __phi171 int32
	_ = __phi171
	var v172 int32
	_ = v172
	var __phi172 int32
	_ = __phi172
	var v174 int64
	_ = v174
	var v176 int32
	_ = v176
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v185 int32
	_ = v185
	var v203 int32
	_ = v203
	var v206 int32
	_ = v206
	var v240 int32
	_ = v240
	var v243 int32
	_ = v243
	var v246 int32
	_ = v246
	var v250 int32
	_ = v250
	var v254 int32
	_ = v254
	var v269 int32
	_ = v269
	var v272 int32
	_ = v272
	var v306 int32
	_ = v306
	var v309 int32
	_ = v309
	var v316 int32
	_ = v316
	var v321 int32
	_ = v321
	var v324 int32
	_ = v324
	var v330 int32
	_ = v330
	var v339 int32
	_ = v339
	var v340 int32
	_ = v340
	var v345 int32
	_ = v345
	var v346 int32
	_ = v346
	var v348 int32
	_ = v348
	var v390 int32
	_ = v390
	var v393 int32
	_ = v393
	var v397 int32
	_ = v397
	var v398 int64
	_ = v398
	var v402 int32
	_ = v402
	var v405 int32
	_ = v405
	var v407 int32
	_ = v407
	var v410 int64
	_ = v410
	var v412 int32
	_ = v412
	var v414 int32
	_ = v414
	var v415 int32
	_ = v415
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v10)+8)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v10))) = int32(1)
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v16 + int32(16)
	v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v31 = v26 + int32(0)
	v32 = m.G398
	if base.Ui32(v31) < base.Ui32(v25) {
		v34 = v31
	} else {
		v34 = v32
	}
	v76 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if base.Ui32(v76) <= base.Ui32(v34) {
		v93 = v76
	} else {
		v79 = v76
		for {
			v83 = v79 + int32(-16)
			v84 = *(*int64)(unsafe.Add(mBase, uint32(v83)))
			*(*int64)(unsafe.Add(mBase, uint32(v79))) = v84
			v88 = *(*int32)(unsafe.Add(mBase, uint32(v79+int32(-8))))
			*(*int32)(unsafe.Add(mBase, uint32(v79)+8)) = v88
			if base.Ui32(v34) < base.Ui32(v83) {
				v79 = v83
				continue
			} else {
				break
			}
			break
		}
		v91 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		v93 = v91
	}
	v96 = *(*int64)(unsafe.Add(mBase, uint32(v93)))
	*(*int64)(unsafe.Add(mBase, uint32(v34))) = v96
	v98 = *(*int32)(unsafe.Add(mBase, uint32(v93)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v34)+8)) = v98
	v100 = int32(-1)
	v104 = F_lua_pcall(m, l0, (v3-v4)>>(uint(int32(4))%32)+v100, v100, int32(0))
	mBase = m.M
	v107 = m.ExcPending
	if v107 != 0 {
		return int32(0)
	} else {
		if v104 == int32(0) {
			v414 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			v415 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
			return (v414 - v415) >> (uint(int32(4)) % 32)
		} else {
			v115 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			v116 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
			v121 = v116 + int32(0)
			v122 = m.G398
			if base.Ui32(v121) < base.Ui32(v115) {
				v124 = v121
			} else {
				v124 = v122
			}
			v167 = v124 + int32(16)
			v168 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			if base.Ui32(v168) <= base.Ui32(v167) {
				v185 = v168
			} else {
				__phi171 = v124
				__phi172 = v167
				v171 = __phi171
				v172 = __phi172
				for {
					v174 = *(*int64)(unsafe.Add(mBase, uint32(v171)+16))
					*(*int64)(unsafe.Add(mBase, uint32(v171))) = v174
					v176 = *(*int32)(unsafe.Add(mBase, uint32(v171)+24))
					*(*int32)(unsafe.Add(mBase, uint32(v171)+8)) = v176
					v179 = v172 + int32(16)
					v180 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
					if base.Ui32(v179) < base.Ui32(v180) {
						__phi171 = v172
						__phi172 = v179
						v171 = __phi171
						v172 = __phi172
						continue
					} else {
						break
					}
					break
				}
				v185 = v180
			}
			*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v185 + int32(-16)
			v203 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			v206 = v203 + int32(-16)
			v240 = m.G398
			if v206 != v240 {
				v243 = *(*int32)(unsafe.Add(mBase, uint32(v206)+8))
				v246 = v243
			} else {
				v246 = int32(-1)
			}
			if v246 != int32(5) {
				v324 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				*(*int32)(unsafe.Add(mBase, uint32(v324)+8)) = int32(1)
				*(*int32)(unsafe.Add(mBase, uint32(v324))) = int32(0)
				v330 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v330 + int32(16)
				v339 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				v340 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
				v345 = v340 + int32(0)
				v346 = m.G398
				if base.Ui32(v345) < base.Ui32(v339) {
					v348 = v345
				} else {
					v348 = v346
				}
				v390 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				if base.Ui32(v390) <= base.Ui32(v348) {
					v407 = v390
				} else {
					v393 = v390
					for {
						v397 = v393 + int32(-16)
						v398 = *(*int64)(unsafe.Add(mBase, uint32(v397)))
						*(*int64)(unsafe.Add(mBase, uint32(v393))) = v398
						v402 = *(*int32)(unsafe.Add(mBase, uint32(v393+int32(-8))))
						*(*int32)(unsafe.Add(mBase, uint32(v393)+8)) = v402
						if base.Ui32(v348) < base.Ui32(v397) {
							v393 = v397
							continue
						} else {
							break
						}
						break
					}
					v405 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
					v407 = v405
				}
				v410 = *(*int64)(unsafe.Add(mBase, uint32(v407)))
				*(*int64)(unsafe.Add(mBase, uint32(v348))) = v410
				v412 = *(*int32)(unsafe.Add(mBase, uint32(v407)+8))
				*(*int32)(unsafe.Add(mBase, uint32(v348)+8)) = v412
				v414 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				v415 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
				return (v414 - v415) >> (uint(int32(4)) % 32)
			} else {
				v250 = m.G3
				F_lua_getfield(m, l0, int32(-1), v250+int32(_a2632))
				mBase = m.M
				v254 = m.ExcPending
				if v254 != 0 {
					return int32(0)
				} else {
					v269 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
					v272 = v269 + int32(-16)
					v306 = m.G398
					if v272 != v306 {
						v309 = *(*int32)(unsafe.Add(mBase, uint32(v272)+8))
						v316 = base.B2i32(base.Ui32(v309+int32(-3)) < base.Ui32(int32(2)))
					} else {
						v316 = int32(0)
					}
					if v316 == int32(0) {
						v324 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
						*(*int32)(unsafe.Add(mBase, uint32(v324)+8)) = int32(1)
						*(*int32)(unsafe.Add(mBase, uint32(v324))) = int32(0)
						v330 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
						*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v330 + int32(16)
						v339 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
						v340 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
						v345 = v340 + int32(0)
						v346 = m.G398
						if base.Ui32(v345) < base.Ui32(v339) {
							v348 = v345
						} else {
							v348 = v346
						}
						v390 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
						if base.Ui32(v390) <= base.Ui32(v348) {
							v407 = v390
						} else {
							v393 = v390
							for {
								v397 = v393 + int32(-16)
								v398 = *(*int64)(unsafe.Add(mBase, uint32(v397)))
								*(*int64)(unsafe.Add(mBase, uint32(v393))) = v398
								v402 = *(*int32)(unsafe.Add(mBase, uint32(v393+int32(-8))))
								*(*int32)(unsafe.Add(mBase, uint32(v393)+8)) = v402
								if base.Ui32(v348) < base.Ui32(v397) {
									v393 = v397
									continue
								} else {
									break
								}
								break
							}
							v405 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
							v407 = v405
						}
						v410 = *(*int64)(unsafe.Add(mBase, uint32(v407)))
						*(*int64)(unsafe.Add(mBase, uint32(v348))) = v410
						v412 = *(*int32)(unsafe.Add(mBase, uint32(v407)+8))
						*(*int32)(unsafe.Add(mBase, uint32(v348)+8)) = v412
						v414 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
						v415 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
						return (v414 - v415) >> (uint(int32(4)) % 32)
					} else {
						F_lua_replace(m, l0, int32(-2))
						mBase = m.M
						v321 = m.ExcPending
						if v321 != 0 {
							return int32(0)
						} else {
							v324 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
							*(*int32)(unsafe.Add(mBase, uint32(v324)+8)) = int32(1)
							*(*int32)(unsafe.Add(mBase, uint32(v324))) = int32(0)
							v330 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
							*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v330 + int32(16)
							v339 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
							v340 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
							v345 = v340 + int32(0)
							v346 = m.G398
							if base.Ui32(v345) < base.Ui32(v339) {
								v348 = v345
							} else {
								v348 = v346
							}
							v390 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
							if base.Ui32(v390) <= base.Ui32(v348) {
								v407 = v390
							} else {
								v393 = v390
								for {
									v397 = v393 + int32(-16)
									v398 = *(*int64)(unsafe.Add(mBase, uint32(v397)))
									*(*int64)(unsafe.Add(mBase, uint32(v393))) = v398
									v402 = *(*int32)(unsafe.Add(mBase, uint32(v393+int32(-8))))
									*(*int32)(unsafe.Add(mBase, uint32(v393)+8)) = v402
									if base.Ui32(v348) < base.Ui32(v397) {
										v393 = v397
										continue
									} else {
										break
									}
									break
								}
								v405 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
								v407 = v405
							}
							v410 = *(*int64)(unsafe.Add(mBase, uint32(v407)))
							*(*int64)(unsafe.Add(mBase, uint32(v348))) = v410
							v412 = *(*int32)(unsafe.Add(mBase, uint32(v407)+8))
							*(*int32)(unsafe.Add(mBase, uint32(v348)+8)) = v412
							v414 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
							v415 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
							return (v414 - v415) >> (uint(int32(4)) % 32)
						}
					}
				}
			}
		}
	}
}
func F_luaRegisterLogFunction(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	var v33 int32
	_ = v33
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v50 int32
	_ = v50
	var v54 int32
	_ = v54
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
	var v78 int32
	_ = v78
	var v84 int32
	_ = v84
	v3 = m.G3
	F_lua_pushstring(m, l0, v3+int32(_a80))
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return
	} else {
		v8 = m.G5
		F_lua_pushcclosure(m, l0, v8+int32(1186), int32(0))
		mBase = m.M
		v13 = m.ExcPending
		if v13 != 0 {
			return
		} else {
			F_lua_settable(m, l0, int32(-3))
			mBase = m.M
			v16 = m.ExcPending
			if v16 != 0 {
				return
			} else {
				F_lua_pushstring(m, l0, v3+int32(_a2633))
				mBase = m.M
				v20 = m.ExcPending
				if v20 != 0 {
					return
				} else {
					v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
					*(*int32)(unsafe.Add(mBase, uint32(v23)+8)) = int32(3)
					*(*float64)(unsafe.Add(mBase, uint32(v23))) = float64(0)
					v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
					*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v27 + int32(16)
					F_lua_settable(m, l0, int32(-3))
					mBase = m.M
					v33 = m.ExcPending
					if v33 != 0 {
						return
					} else {
						F_lua_pushstring(m, l0, v3+int32(_a2634))
						mBase = m.M
						v37 = m.ExcPending
						if v37 != 0 {
							return
						} else {
							v40 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
							*(*int32)(unsafe.Add(mBase, uint32(v40)+8)) = int32(3)
							*(*float64)(unsafe.Add(mBase, uint32(v40))) = float64(1)
							v44 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
							*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v44 + int32(16)
							F_lua_settable(m, l0, int32(-3))
							mBase = m.M
							v50 = m.ExcPending
							if v50 != 0 {
								return
							} else {
								F_lua_pushstring(m, l0, v3+int32(_a2635))
								mBase = m.M
								v54 = m.ExcPending
								if v54 != 0 {
									return
								} else {
									v57 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
									*(*int32)(unsafe.Add(mBase, uint32(v57)+8)) = int32(3)
									*(*float64)(unsafe.Add(mBase, uint32(v57))) = float64(2)
									v61 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
									*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v61 + int32(16)
									F_lua_settable(m, l0, int32(-3))
									mBase = m.M
									v67 = m.ExcPending
									if v67 != 0 {
										return
									} else {
										F_lua_pushstring(m, l0, v3+int32(_a2636))
										mBase = m.M
										v71 = m.ExcPending
										if v71 != 0 {
											return
										} else {
											v74 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
											*(*int32)(unsafe.Add(mBase, uint32(v74)+8)) = int32(3)
											*(*float64)(unsafe.Add(mBase, uint32(v74))) = float64(3)
											v78 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
											*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v78 + int32(16)
											F_lua_settable(m, l0, int32(-3))
											mBase = m.M
											v84 = m.ExcPending
											if v84 != 0 {
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
func F_luaReplyToServerReply(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v102 int32
	_ = v102
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v143 int32
	_ = v143
	var v146 int32
	_ = v146
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v187 int32
	_ = v187
	var v189 int32
	_ = v189
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v219 int32
	_ = v219
	var v222 int32
	_ = v222
	var v258 int32
	_ = v258
	var v259 int32
	_ = v259
	var v263 int32
	_ = v263
	var v265 int32
	_ = v265
	var v266 int32
	_ = v266
	var v267 int32
	_ = v267
	var v268 int32
	_ = v268
	var v269 int32
	_ = v269
	var v271 float64
	_ = v271
	var v272 int32
	_ = v272
	var v278 int64
	_ = v278
	var v279 int32
	_ = v279
	var v280 int32
	_ = v280
	var v282 int32
	_ = v282
	var v283 int32
	_ = v283
	var v284 int32
	_ = v284
	var v288 int32
	_ = v288
	var v306 int32
	_ = v306
	var v345 int32
	_ = v345
	var v346 int32
	_ = v346
	var v347 int32
	_ = v347
	var v349 int32
	_ = v349
	var v350 int32
	_ = v350
	var v353 int64
	_ = v353
	var v357 int32
	_ = v357
	var v373 int32
	_ = v373
	var v376 int32
	_ = v376
	var v410 int32
	_ = v410
	var v413 int32
	_ = v413
	var v416 int32
	_ = v416
	var v435 int32
	_ = v435
	var v449 int64
	_ = v449
	var v456 int32
	_ = v456
	var v457 int32
	_ = v457
	var v459 int32
	_ = v459
	var v460 int32
	_ = v460
	var v461 int32
	_ = v461
	var v466 int32
	_ = v466
	var v467 int32
	_ = v467
	var v468 int32
	_ = v468
	var v471 int32
	_ = v471
	var v472 int32
	_ = v472
	var v474 int32
	_ = v474
	var v475 int32
	_ = v475
	var v478 int32
	_ = v478
	var v479 int32
	_ = v479
	var v481 int32
	_ = v481
	var v482 int32
	_ = v482
	var v485 int32
	_ = v485
	var v486 int32
	_ = v486
	var v488 int32
	_ = v488
	var v507 int32
	_ = v507
	var v517 int32
	_ = v517
	var v521 int32
	_ = v521
	var v539 int32
	_ = v539
	var v578 int32
	_ = v578
	var v579 int32
	_ = v579
	var v580 int32
	_ = v580
	var v582 int32
	_ = v582
	var v583 int32
	_ = v583
	var v586 int64
	_ = v586
	var v590 int32
	_ = v590
	var v606 int32
	_ = v606
	var v609 int32
	_ = v609
	var v643 int32
	_ = v643
	var v646 int32
	_ = v646
	var v649 int32
	_ = v649
	var v652 int32
	_ = v652
	var v655 int32
	_ = v655
	var v656 int32
	_ = v656
	var v658 int32
	_ = v658
	var v659 int32
	_ = v659
	var v662 int32
	_ = v662
	var v663 int32
	_ = v663
	var v664 int32
	_ = v664
	var v665 int32
	_ = v665
	var v666 int32
	_ = v666
	var v668 int32
	_ = v668
	var v677 int32
	_ = v677
	var v680 int32
	_ = v680
	var v684 int32
	_ = v684
	var v689 int32
	_ = v689
	var v691 int32
	_ = v691
	var v695 int32
	_ = v695
	var v701 int32
	_ = v701
	var v704 int32
	_ = v704
	var v710 int32
	_ = v710
	var v714 int32
	_ = v714
	var v716 int32
	_ = v716
	var v724 int32
	_ = v724
	var v727 int32
	_ = v727
	var v733 int32
	_ = v733
	var v739 int32
	_ = v739
	var v745 int32
	_ = v745
	var v747 int32
	_ = v747
	var v748 int32
	_ = v748
	var v752 int32
	_ = v752
	var v753 int32
	_ = v753
	var v757 int32
	_ = v757
	var v762 int32
	_ = v762
	var v763 int32
	_ = v763
	var v767 int32
	_ = v767
	var v768 int32
	_ = v768
	var v772 int32
	_ = v772
	var v775 int32
	_ = v775
	var v776 int32
	_ = v776
	var v778 int32
	_ = v778
	var v784 int32
	_ = v784
	var v794 int32
	_ = v794
	var v795 int32
	_ = v795
	var v799 int32
	_ = v799
	var v800 int32
	_ = v800
	var v804 int32
	_ = v804
	var v817 int32
	_ = v817
	var v818 int32
	_ = v818
	var v819 int32
	_ = v819
	var v820 int32
	_ = v820
	var v821 int32
	_ = v821
	var v822 int32
	_ = v822
	var v824 int32
	_ = v824
	var v843 int32
	_ = v843
	var v871 int32
	_ = v871
	var v881 int32
	_ = v881
	var v885 int32
	_ = v885
	var v903 int32
	_ = v903
	var v942 int32
	_ = v942
	var v943 int32
	_ = v943
	var v944 int32
	_ = v944
	var v946 int32
	_ = v946
	var v947 int32
	_ = v947
	var v950 int64
	_ = v950
	var v954 int32
	_ = v954
	var v970 int32
	_ = v970
	var v973 int32
	_ = v973
	var v1007 int32
	_ = v1007
	var v1010 int32
	_ = v1010
	var v1013 int32
	_ = v1013
	var v1016 int32
	_ = v1016
	var v1017 int32
	_ = v1017
	var v1019 float64
	_ = v1019
	var v1020 int32
	_ = v1020
	var v1021 int32
	_ = v1021
	var v1022 int32
	_ = v1022
	var v1041 int32
	_ = v1041
	var v1069 int32
	_ = v1069
	var v1079 int32
	_ = v1079
	var v1083 int32
	_ = v1083
	var v1101 int32
	_ = v1101
	var v1140 int32
	_ = v1140
	var v1141 int32
	_ = v1141
	var v1142 int32
	_ = v1142
	var v1144 int32
	_ = v1144
	var v1145 int32
	_ = v1145
	var v1148 int64
	_ = v1148
	var v1152 int32
	_ = v1152
	var v1168 int32
	_ = v1168
	var v1171 int32
	_ = v1171
	var v1205 int32
	_ = v1205
	var v1208 int32
	_ = v1208
	var v1211 int32
	_ = v1211
	var v1214 int32
	_ = v1214
	var v1215 int32
	_ = v1215
	var v1216 int32
	_ = v1216
	var v1217 int32
	_ = v1217
	var v1228 int32
	_ = v1228
	var v1237 int32
	_ = v1237
	var v1244 int32
	_ = v1244
	var v1245 int32
	_ = v1245
	var v1254 int32
	_ = v1254
	var v1258 int32
	_ = v1258
	var v1261 int32
	_ = v1261
	var v1264 int32
	_ = v1264
	var v1276 int32
	_ = v1276
	var v1294 int32
	_ = v1294
	var v1297 int32
	_ = v1297
	var v1301 int32
	_ = v1301
	var v1306 int32
	_ = v1306
	var v1308 int32
	_ = v1308
	var v1312 int32
	_ = v1312
	var v1318 int32
	_ = v1318
	var v1321 int32
	_ = v1321
	var v1327 int32
	_ = v1327
	var v1331 int32
	_ = v1331
	var v1333 int32
	_ = v1333
	var v1341 int32
	_ = v1341
	var v1342 int32
	_ = v1342
	var v1343 int32
	_ = v1343
	var v1344 int32
	_ = v1344
	var v1345 int32
	_ = v1345
	var v1346 int32
	_ = v1346
	var v1348 int32
	_ = v1348
	var v1367 int32
	_ = v1367
	var v1395 int32
	_ = v1395
	var v1405 int32
	_ = v1405
	var v1409 int32
	_ = v1409
	var v1427 int32
	_ = v1427
	var v1466 int32
	_ = v1466
	var v1467 int32
	_ = v1467
	var v1468 int32
	_ = v1468
	var v1470 int32
	_ = v1470
	var v1471 int32
	_ = v1471
	var v1474 int64
	_ = v1474
	var v1478 int32
	_ = v1478
	var v1494 int32
	_ = v1494
	var v1497 int32
	_ = v1497
	var v1531 int32
	_ = v1531
	var v1534 int32
	_ = v1534
	var v1537 int32
	_ = v1537
	var v1540 int32
	_ = v1540
	var v1544 int32
	_ = v1544
	var v1562 int32
	_ = v1562
	var v1601 int32
	_ = v1601
	var v1602 int32
	_ = v1602
	var v1603 int32
	_ = v1603
	var v1605 int32
	_ = v1605
	var v1606 int32
	_ = v1606
	var v1609 int64
	_ = v1609
	var v1613 int32
	_ = v1613
	var v1629 int32
	_ = v1629
	var v1632 int32
	_ = v1632
	var v1666 int32
	_ = v1666
	var v1669 int32
	_ = v1669
	var v1672 int32
	_ = v1672
	var v1675 int32
	_ = v1675
	var v1678 int32
	_ = v1678
	var v1679 int32
	_ = v1679
	var v1683 int32
	_ = v1683
	var v1701 int32
	_ = v1701
	var v1740 int32
	_ = v1740
	var v1741 int32
	_ = v1741
	var v1742 int32
	_ = v1742
	var v1744 int32
	_ = v1744
	var v1745 int32
	_ = v1745
	var v1748 int64
	_ = v1748
	var v1752 int32
	_ = v1752
	var v1768 int32
	_ = v1768
	var v1771 int32
	_ = v1771
	var v1805 int32
	_ = v1805
	var v1808 int32
	_ = v1808
	var v1811 int32
	_ = v1811
	var v1814 int32
	_ = v1814
	var v1818 int32
	_ = v1818
	var v1819 int32
	_ = v1819
	var v1820 int32
	_ = v1820
	var v1821 int32
	_ = v1821
	var v1822 int32
	_ = v1822
	var v1823 int32
	_ = v1823
	var v1842 int32
	_ = v1842
	var v1870 int32
	_ = v1870
	var v1900 int32
	_ = v1900
	var v1930 int32
	_ = v1930
	var v1940 int32
	_ = v1940
	var v1944 int32
	_ = v1944
	var v1962 int32
	_ = v1962
	var v2001 int32
	_ = v2001
	var v2002 int32
	_ = v2002
	var v2003 int32
	_ = v2003
	var v2005 int32
	_ = v2005
	var v2006 int32
	_ = v2006
	var v2009 int64
	_ = v2009
	var v2013 int32
	_ = v2013
	var v2029 int32
	_ = v2029
	var v2032 int32
	_ = v2032
	var v2066 int32
	_ = v2066
	var v2069 int32
	_ = v2069
	var v2072 int32
	_ = v2072
	var v2076 int32
	_ = v2076
	var v2077 int32
	_ = v2077
	var v2078 int32
	_ = v2078
	var v2079 int32
	_ = v2079
	var v2081 int32
	_ = v2081
	var v2087 int32
	_ = v2087
	var v2089 int32
	_ = v2089
	var v2090 int32
	_ = v2090
	var v2097 int32
	_ = v2097
	var v2121 int32
	_ = v2121
	var v2124 int32
	_ = v2124
	var v2160 int32
	_ = v2160
	var v2161 int64
	_ = v2161
	var v2163 int32
	_ = v2163
	var v2165 int32
	_ = v2165
	var v2170 int32
	_ = v2170
	var v2172 int32
	_ = v2172
	var v2174 int32
	_ = v2174
	var v2176 int32
	_ = v2176
	var v2177 int32
	_ = v2177
	var v2182 int32
	_ = v2182
	var v2189 int32
	_ = v2189
	var v2190 int32
	_ = v2190
	var v2192 int32
	_ = v2192
	var v2211 int32
	_ = v2211
	var v2239 int32
	_ = v2239
	var v2249 int32
	_ = v2249
	var v2253 int32
	_ = v2253
	var v2271 int32
	_ = v2271
	var v2310 int32
	_ = v2310
	var v2311 int32
	_ = v2311
	var v2312 int32
	_ = v2312
	var v2314 int32
	_ = v2314
	var v2315 int32
	_ = v2315
	var v2318 int64
	_ = v2318
	var v2322 int32
	_ = v2322
	var v2338 int32
	_ = v2338
	var v2341 int32
	_ = v2341
	var v2375 int32
	_ = v2375
	var v2378 int32
	_ = v2378
	var v2381 int32
	_ = v2381
	var v2385 int32
	_ = v2385
	var v2386 int32
	_ = v2386
	var v2387 int32
	_ = v2387
	var v2388 int32
	_ = v2388
	var v2390 int32
	_ = v2390
	var v2396 int32
	_ = v2396
	var v2398 int32
	_ = v2398
	var v2399 int32
	_ = v2399
	var v2406 int32
	_ = v2406
	var v2431 int32
	_ = v2431
	var v2458 int32
	_ = v2458
	var v2461 int32
	_ = v2461
	var v2497 int32
	_ = v2497
	var v2498 int64
	_ = v2498
	var v2500 int32
	_ = v2500
	var v2502 int32
	_ = v2502
	var v2507 int32
	_ = v2507
	var v2509 int32
	_ = v2509
	var v2511 int32
	_ = v2511
	var v2512 int32
	_ = v2512
	var v2517 int32
	_ = v2517
	var v2524 int32
	_ = v2524
	var v2525 int32
	_ = v2525
	var v2527 int32
	_ = v2527
	var v2546 int32
	_ = v2546
	var v2574 int32
	_ = v2574
	var v2585 int32
	_ = v2585
	var v2586 int32
	_ = v2586
	var v2587 int32
	_ = v2587
	var v2588 int32
	_ = v2588
	var v2591 int32
	_ = v2591
	var v2595 int32
	_ = v2595
	var v2616 int32
	_ = v2616
	var v2655 int32
	_ = v2655
	var v2656 int32
	_ = v2656
	var v2657 int32
	_ = v2657
	var v2659 int32
	_ = v2659
	var v2660 int32
	_ = v2660
	var v2663 int64
	_ = v2663
	var v2667 int32
	_ = v2667
	var v2669 int32
	_ = v2669
	var v2684 int32
	_ = v2684
	var v2687 int32
	_ = v2687
	var v2721 int32
	_ = v2721
	var v2724 int32
	_ = v2724
	var v2727 int32
	_ = v2727
	var v2735 int32
	_ = v2735
	var v2737 int32
	_ = v2737
	var v2743 int32
	_ = v2743
	var v2745 int32
	_ = v2745
	var v2748 int32
	_ = v2748
	var v2752 int32
	_ = v2752
	var v2773 int32
	_ = v2773
	var v2812 int32
	_ = v2812
	var v2813 int32
	_ = v2813
	var v2814 int32
	_ = v2814
	var v2816 int32
	_ = v2816
	var v2817 int32
	_ = v2817
	var v2820 int64
	_ = v2820
	var v2824 int32
	_ = v2824
	var v2827 int32
	_ = v2827
	var v2842 int32
	_ = v2842
	var v2845 int32
	_ = v2845
	var v2879 int32
	_ = v2879
	var v2882 int32
	_ = v2882
	var v2885 int32
	_ = v2885
	var v2890 int32
	_ = v2890
	var v2915 int32
	_ = v2915
	var v2925 int32
	_ = v2925
	var v2926 int32
	_ = v2926
	var v2928 int32
	_ = v2928
	var v2929 int32
	_ = v2929
	var v2930 int32
	_ = v2930
	var v2931 int32
	_ = v2931
	var v2932 int32
	_ = v2932
	var v2962 int32
	_ = v2962
	v12 = m.G0
	v14 = v12 - int32(32)
	m.G0 = v14
	goto L3
L1:
	;
	v75 = F_lua_checkstack(m, l2, int32(4))
	mBase = m.M
	v76 = m.ExcPending
	if v76 != 0 {
		goto L18
	} else {
		goto L19
	}
L2:
	;
	v67 = m.G398
	if v33 != v67 {
		goto L14
	} else {
		goto L15
	}
L3:
	;
	goto L7
L7:
	;
	v30 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	v33 = v30 + int32(-16)
	goto L2
L14:
	;
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v33)+8))
	v73 = v70
	goto L1
L15:
	;
	v73 = int32(-1)
	goto L1
L16:
	;
	m.G0 = v14 + int32(32)
	return
L17:
	;
	switch v73 + int32(-1) {
	case 0:
		goto L34
	default:
		goto L31
	case 2:
		goto L33
	case 3:
		goto L35
	case 4:
		goto L32
	}
L18:
	;
	return
L19:
	;
	if v75 != 0 {
		goto L17
	} else {
		goto L20
	}
L20:
	;
	v77 = m.G3
	v80 = m.G39
	v81 = *(*int32)(unsafe.Add(mBase, uint32(v80)))
	v82 = m.T0[v81].(func(*base.Module, int32, int32) int32)(m, l0, v77+int32(_a2638))
	mBase = m.M
	v83 = m.ExcPending
	if v83 != 0 {
		goto L18
	} else {
		goto L21
	}
L21:
	;
	goto L24
L22:
	;
	goto L16
L23:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2)+8)) = v102 + int32(-16)
	goto L22
L24:
	;
	v102 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	goto L23
L30:
	;
	goto L755
L31:
	;
	v2929 = m.G60
	v2930 = *(*int32)(unsafe.Add(mBase, uint32(v2929)))
	v2931 = m.T0[v2930].(func(*base.Module, int32) int32)(m, l0)
	mBase = m.M
	v2932 = m.ExcPending
	if v2932 != 0 {
		goto L18
	} else {
		goto L752
	}
L32:
	;
	v284 = m.G3
	F_lua_pushstring(m, l2, v284+int32(_a2632))
	mBase = m.M
	v288 = m.ExcPending
	if v288 != 0 {
		goto L18
	} else {
		goto L89
	}
L33:
	;
	v268 = m.G38
	v269 = *(*int32)(unsafe.Add(mBase, uint32(v268)))
	v271 = F_lua_tonumber(m, l2, int32(-1))
	mBase = m.M
	v272 = m.ExcPending
	if v272 != 0 {
		goto L18
	} else {
		goto L85
	}
L34:
	;
	if l1 != int32(2) {
		goto L39
	} else {
		goto L40
	}
L35:
	;
	v114 = m.G54
	v115 = *(*int32)(unsafe.Add(mBase, uint32(v114)))
	v118 = F_lua_tolstring(m, l2, int32(-1), int32(0))
	mBase = m.M
	v119 = m.ExcPending
	if v119 != 0 {
		goto L18
	} else {
		goto L36
	}
L36:
	;
	v121 = F_lua_objlen(m, l2, int32(-1))
	mBase = m.M
	v122 = m.ExcPending
	if v122 != 0 {
		goto L18
	} else {
		goto L37
	}
L37:
	;
	v123 = m.T0[v115].(func(*base.Module, int32, int32, int32) int32)(m, l0, v118, v121)
	mBase = m.M
	v124 = m.ExcPending
	if v124 != 0 {
		goto L18
	} else {
		goto L38
	}
L38:
	;
	goto L30
L39:
	;
	v201 = m.G61
	v202 = *(*int32)(unsafe.Add(mBase, uint32(v201)))
	goto L66
L40:
	;
	goto L44
L41:
	;
	v197 = m.G60
	v198 = *(*int32)(unsafe.Add(mBase, uint32(v197)))
	v199 = m.T0[v198].(func(*base.Module, int32) int32)(m, l0)
	mBase = m.M
	v200 = m.ExcPending
	if v200 != 0 {
		goto L18
	} else {
		goto L63
	}
L42:
	;
	if v189 == int32(0) {
		goto L41
	} else {
		goto L61
	}
L43:
	;
	v182 = *(*int32)(unsafe.Add(mBase, uint32(v146)+8))
	switch v182 {
	case 0:
		v187 = v182
		goto L58
	case 1:
		goto L60
	default:
		goto L59
	}
L44:
	;
	goto L50
L50:
	;
	v143 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	v146 = v143 + int32(-16)
	goto L43
L58:
	;
	v189 = v187
	goto L42
L59:
	;
	v187 = int32(1)
	goto L58
L60:
	;
	v183 = *(*int32)(unsafe.Add(mBase, uint32(v146)))
	v189 = base.B2i32(v183 != int32(0))
	goto L42
L61:
	;
	v193 = m.G38
	v194 = *(*int32)(unsafe.Add(mBase, uint32(v193)))
	v195 = m.T0[v194].(func(*base.Module, int32, int64) int32)(m, l0, int64(1))
	mBase = m.M
	v196 = m.ExcPending
	if v196 != 0 {
		goto L18
	} else {
		goto L62
	}
L62:
	;
	goto L30
L63:
	;
	goto L30
L64:
	;
	v266 = m.T0[v202].(func(*base.Module, int32, int32) int32)(m, l0, v265)
	mBase = m.M
	v267 = m.ExcPending
	if v267 != 0 {
		goto L18
	} else {
		goto L83
	}
L65:
	;
	v258 = *(*int32)(unsafe.Add(mBase, uint32(v222)+8))
	switch v258 {
	case 0:
		v263 = v258
		goto L80
	case 1:
		goto L82
	default:
		goto L81
	}
L66:
	;
	goto L72
L72:
	;
	v219 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	v222 = v219 + int32(-16)
	goto L65
L80:
	;
	v265 = v263
	goto L64
L81:
	;
	v263 = int32(1)
	goto L80
L82:
	;
	v259 = *(*int32)(unsafe.Add(mBase, uint32(v222)))
	v265 = base.B2i32(v259 != int32(0))
	goto L64
L83:
	;
	goto L30
L84:
	;
	v282 = m.T0[v269].(func(*base.Module, int32, int64) int32)(m, l0, int64(-9223372036854775807-1))
	mBase = m.M
	v283 = m.ExcPending
	if v283 != 0 {
		goto L18
	} else {
		goto L88
	}
L85:
	;
	if base.F64_lt(base.F64_abs(v271), float64(9.223372036854776e+18)) == int32(0) {
		goto L84
	} else {
		goto L86
	}
L86:
	;
	v278 = base.I64_trunc_f64_s(v271)
	v279 = m.T0[v269].(func(*base.Module, int32, int64) int32)(m, l0, v278)
	mBase = m.M
	v280 = m.ExcPending
	if v280 != 0 {
		goto L18
	} else {
		goto L87
	}
L87:
	;
	goto L30
L88:
	;
	goto L30
L89:
	;
	goto L92
L90:
	;
	goto L108
L91:
	;
	v345 = *(*int32)(unsafe.Add(mBase, uint32(v306+int32(-32))))
	v346 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	v347 = int32(-16)
	v349 = F_luaH_get(m, v345, v346+v347)
	mBase = m.M
	v350 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	v353 = *(*int64)(unsafe.Add(mBase, uint32(v349)))
	*(*int64)(unsafe.Add(mBase, uint32(v350+v347))) = v353
	v357 = *(*int32)(unsafe.Add(mBase, uint32(v349)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v350+int32(-8)))) = v357
	goto L90
L92:
	;
	goto L98
L98:
	;
	v306 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	goto L91
L106:
	;
	goto L123
L107:
	;
	v410 = m.G398
	if v376 != v410 {
		goto L119
	} else {
		goto L120
	}
L108:
	;
	goto L112
L112:
	;
	v373 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	v376 = v373 + int32(-16)
	goto L107
L119:
	;
	v413 = *(*int32)(unsafe.Add(mBase, uint32(v376)+8))
	v416 = v413
	goto L106
L120:
	;
	v416 = int32(-1)
	goto L106
L121:
	;
	if v416 != int32(4) {
		goto L129
	} else {
		goto L130
	}
L122:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2)+8)) = v435 + int32(-16)
	goto L121
L123:
	;
	v435 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	goto L122
L129:
	;
	v517 = m.G3
	F_lua_pushstring(m, l2, v517+int32(_a2639))
	mBase = m.M
	v521 = m.ExcPending
	if v521 != 0 {
		goto L18
	} else {
		goto L150
	}
L130:
	;
	v449 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v14+int32(24)))) = v449
	*(*int64)(unsafe.Add(mBase, uint32(v14)+16)) = v449
	F_luaExtractErrorInformation(m, l2, v14+int32(16))
	mBase = m.M
	v456 = m.ExcPending
	if v456 != 0 {
		goto L18
	} else {
		goto L131
	}
L131:
	;
	v457 = *(*int32)(unsafe.Add(mBase, uint32(v14)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v14))) = v457
	v459 = m.G3
	v460 = m.G41
	v461 = *(*int32)(unsafe.Add(mBase, uint32(v14)+28))
	v466 = *(*int32)(unsafe.Add(mBase, uint32(v460)))
	v467 = m.T0[v466].(func(*base.Module, int32, int32, int32, int32) int32)(m, l0, base.B2i32(v461 == int32(0)), v459+int32(_a16), v14)
	mBase = m.M
	v468 = m.ExcPending
	if v468 != 0 {
		goto L18
	} else {
		goto L132
	}
L132:
	;
	if v457 == int32(0) {
		goto L133
	} else {
		goto L134
	}
L133:
	;
	v475 = *(*int32)(unsafe.Add(mBase, uint32(v14)+20))
	if v475 == int32(0) {
		goto L136
	} else {
		goto L137
	}
L134:
	;
	v471 = m.G11
	v472 = *(*int32)(unsafe.Add(mBase, uint32(v471)))
	m.T0[v472].(func(*base.Module, int32))(m, v457)
	mBase = m.M
	v474 = m.ExcPending
	if v474 != 0 {
		goto L18
	} else {
		goto L135
	}
L135:
	;
	goto L133
L136:
	;
	v482 = *(*int32)(unsafe.Add(mBase, uint32(v14)+24))
	if v482 == int32(0) {
		goto L139
	} else {
		goto L140
	}
L137:
	;
	v478 = m.G11
	v479 = *(*int32)(unsafe.Add(mBase, uint32(v478)))
	m.T0[v479].(func(*base.Module, int32))(m, v475)
	mBase = m.M
	v481 = m.ExcPending
	if v481 != 0 {
		goto L18
	} else {
		goto L138
	}
L138:
	;
	goto L136
L139:
	;
	goto L144
L140:
	;
	v485 = m.G11
	v486 = *(*int32)(unsafe.Add(mBase, uint32(v485)))
	m.T0[v486].(func(*base.Module, int32))(m, v482)
	mBase = m.M
	v488 = m.ExcPending
	if v488 != 0 {
		goto L18
	} else {
		goto L141
	}
L141:
	;
	goto L139
L142:
	;
	goto L16
L143:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2)+8)) = v507 + int32(-16)
	goto L142
L144:
	;
	v507 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	goto L143
L150:
	;
	goto L153
L151:
	;
	goto L170
L152:
	;
	v578 = *(*int32)(unsafe.Add(mBase, uint32(v539+int32(-32))))
	v579 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	v580 = int32(-16)
	v582 = F_luaH_get(m, v578, v579+v580)
	mBase = m.M
	v583 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	v586 = *(*int64)(unsafe.Add(mBase, uint32(v582)))
	*(*int64)(unsafe.Add(mBase, uint32(v583+v580))) = v586
	v590 = *(*int32)(unsafe.Add(mBase, uint32(v582)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v583+int32(-8)))) = v590
	goto L151
L153:
	;
	goto L159
L159:
	;
	v539 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	goto L152
L167:
	;
	goto L232
L168:
	;
	if v649 != int32(4) {
		goto L167
	} else {
		goto L183
	}
L169:
	;
	v643 = m.G398
	if v609 != v643 {
		goto L181
	} else {
		goto L182
	}
L170:
	;
	goto L174
L174:
	;
	v606 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	v609 = v606 + int32(-16)
	goto L169
L181:
	;
	v646 = *(*int32)(unsafe.Add(mBase, uint32(v609)+8))
	v649 = v646
	goto L168
L182:
	;
	v649 = int32(-1)
	goto L168
L183:
	;
	v652 = int32(0)
	v655 = F_lua_tolstring(m, l2, int32(-1), v652)
	mBase = m.M
	v656 = m.ExcPending
	if v656 != 0 {
		goto L18
	} else {
		goto L184
	}
L184:
	;
	v658 = F_lua_objlen(m, l2, int32(-1))
	mBase = m.M
	v659 = m.ExcPending
	if v659 != 0 {
		goto L18
	} else {
		goto L185
	}
L185:
	;
	v662 = m.G22
	v663 = *(*int32)(unsafe.Add(mBase, uint32(v662)))
	v664 = m.T0[v663].(func(*base.Module, int32) int32)(m, v658+int32(1))
	mBase = m.M
	v665 = m.ExcPending
	if v665 != 0 {
		goto L18
	} else {
		goto L186
	}
L186:
	;
	v666 = F___stpncpy(m, v664, v655, v658)
	mBase = m.M
	goto L187
L187:
	;
	v668 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v658+v664))) = uint8(v668)
	if v664&int32(3) == v668 {
		v691 = v664
		goto L191
	} else {
		goto L192
	}
L188:
	;
	v817 = m.G42
	v818 = *(*int32)(unsafe.Add(mBase, uint32(v817)))
	v819 = m.T0[v818].(func(*base.Module, int32, int32) int32)(m, l0, v664)
	mBase = m.M
	v820 = m.ExcPending
	if v820 != 0 {
		goto L18
	} else {
		goto L220
	}
L189:
	;
	if v724 == int32(0) {
		goto L188
	} else {
		goto L205
	}
L190:
	;
	v724 = v716 - v664
	goto L189
L191:
	;
	v695 = v691
	goto L199
L192:
	;
	v677 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v664))))
	if v677 != 0 {
		goto L193
	} else {
		goto L194
	}
L193:
	;
	v680 = v664
	goto L195
L194:
	;
	v724 = v664 - v664
	goto L189
L195:
	;
	v684 = v680 + int32(1)
	if v684&int32(3) == int32(0) {
		v691 = v684
		goto L191
	} else {
		goto L197
	}
L197:
	;
	v689 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v684))))
	if v689 != 0 {
		v680 = v684
		goto L195
	} else {
		goto L198
	}
L198:
	;
	v716 = v684
	goto L190
L199:
	;
	v701 = *(*int32)(unsafe.Add(mBase, uint32(v695)))
	v704 = int32(-2139062144)
	if (int32(16843008)-v701|v701)&v704 == v704 {
		v695 = v695 + int32(4)
		goto L199
	} else {
		goto L201
	}
L200:
	;
	v710 = v695
	goto L202
L201:
	;
	goto L200
L202:
	;
	v714 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v710))))
	if v714 != 0 {
		v710 = v710 + int32(1)
		goto L202
	} else {
		goto L204
	}
L203:
	;
	v716 = v710
	goto L190
L204:
	;
	goto L203
L205:
	;
	v727 = int32(1)
	if v724 == v727 {
		v784 = v652
		goto L206
	} else {
		goto L207
	}
L206:
	;
	if v724&v727 == int32(0) {
		goto L188
	} else {
		goto L217
	}
L207:
	;
	v733 = int32(0)
	v739 = v733
	v745 = v733
	goto L208
L208:
	;
	v747 = v664 + v739
	v748 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v747))))
	switch v748 + int32(-10) {
	case 0:
		goto L212
	default:
		goto L210
	case 3:
		v752 = int32(0)
		goto L211
	}
L209:
	;
	v784 = v776
	goto L206
L210:
	;
	v762 = v747 + int32(1)
	v763 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v762))))
	switch v763 + int32(-10) {
	case 0:
		goto L215
	default:
		goto L213
	case 3:
		v767 = int32(0)
		goto L214
	}
L211:
	;
	v753 = m.G3
	v757 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v753+int32(_a2640)+v752))))
	*(*uint8)(unsafe.Add(mBase, uint32(v747))) = uint8(v757)
	goto L210
L212:
	;
	v752 = int32(1)
	goto L211
L213:
	;
	v775 = int32(2)
	v776 = v739 + v775
	v778 = v745 + v775
	if v778 != v724&int32(-2) {
		v739 = v776
		v745 = v778
		goto L208
	} else {
		goto L216
	}
L214:
	;
	v768 = m.G3
	v772 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v768+int32(_a2640)+v767))))
	*(*uint8)(unsafe.Add(mBase, uint32(v762))) = uint8(v772)
	goto L213
L215:
	;
	v767 = int32(1)
	goto L214
L216:
	;
	goto L209
L217:
	;
	v794 = v664 + v784
	v795 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v794))))
	switch v795 + int32(-10) {
	case 0:
		goto L219
	default:
		goto L188
	case 3:
		v799 = int32(0)
		goto L218
	}
L218:
	;
	v800 = m.G3
	v804 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v800+int32(_a2640)+v799))))
	*(*uint8)(unsafe.Add(mBase, uint32(v794))) = uint8(v804)
	goto L188
L219:
	;
	v799 = int32(1)
	goto L218
L220:
	;
	v821 = m.G11
	v822 = *(*int32)(unsafe.Add(mBase, uint32(v821)))
	m.T0[v822].(func(*base.Module, int32))(m, v664)
	mBase = m.M
	v824 = m.ExcPending
	if v824 != 0 {
		goto L18
	} else {
		goto L221
	}
L221:
	;
	goto L224
L222:
	;
	goto L16
L223:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2)+8)) = v843 + int32(-32)
	goto L222
L224:
	;
	v843 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	goto L223
L230:
	;
	v881 = m.G3
	F_lua_pushstring(m, l2, v881+int32(_a2641))
	mBase = m.M
	v885 = m.ExcPending
	if v885 != 0 {
		goto L18
	} else {
		goto L238
	}
L231:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2)+8)) = v871 + int32(-16)
	goto L230
L232:
	;
	v871 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	goto L231
L238:
	;
	goto L241
L239:
	;
	goto L258
L240:
	;
	v942 = *(*int32)(unsafe.Add(mBase, uint32(v903+int32(-32))))
	v943 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	v944 = int32(-16)
	v946 = F_luaH_get(m, v942, v943+v944)
	mBase = m.M
	v947 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	v950 = *(*int64)(unsafe.Add(mBase, uint32(v946)))
	*(*int64)(unsafe.Add(mBase, uint32(v947+v944))) = v950
	v954 = *(*int32)(unsafe.Add(mBase, uint32(v946)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v947+int32(-8)))) = v954
	goto L239
L241:
	;
	goto L247
L247:
	;
	v903 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	goto L240
L255:
	;
	goto L284
L256:
	;
	if v1013 != int32(3) {
		goto L255
	} else {
		goto L271
	}
L257:
	;
	v1007 = m.G398
	if v973 != v1007 {
		goto L269
	} else {
		goto L270
	}
L258:
	;
	goto L262
L262:
	;
	v970 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	v973 = v970 + int32(-16)
	goto L257
L269:
	;
	v1010 = *(*int32)(unsafe.Add(mBase, uint32(v973)+8))
	v1013 = v1010
	goto L256
L270:
	;
	v1013 = int32(-1)
	goto L256
L271:
	;
	v1016 = m.G63
	v1017 = *(*int32)(unsafe.Add(mBase, uint32(v1016)))
	v1019 = F_lua_tonumber(m, l2, int32(-1))
	mBase = m.M
	v1020 = m.ExcPending
	if v1020 != 0 {
		goto L18
	} else {
		goto L272
	}
L272:
	;
	v1021 = m.T0[v1017].(func(*base.Module, int32, float64) int32)(m, l0, v1019)
	mBase = m.M
	v1022 = m.ExcPending
	if v1022 != 0 {
		goto L18
	} else {
		goto L273
	}
L273:
	;
	goto L276
L274:
	;
	goto L16
L275:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2)+8)) = v1041 + int32(-32)
	goto L274
L276:
	;
	v1041 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	goto L275
L282:
	;
	v1079 = m.G3
	F_lua_pushstring(m, l2, v1079+int32(_a2642))
	mBase = m.M
	v1083 = m.ExcPending
	if v1083 != 0 {
		goto L18
	} else {
		goto L290
	}
L283:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2)+8)) = v1069 + int32(-16)
	goto L282
L284:
	;
	v1069 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	goto L283
L290:
	;
	goto L293
L291:
	;
	goto L310
L292:
	;
	v1140 = *(*int32)(unsafe.Add(mBase, uint32(v1101+int32(-32))))
	v1141 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	v1142 = int32(-16)
	v1144 = F_luaH_get(m, v1140, v1141+v1142)
	mBase = m.M
	v1145 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	v1148 = *(*int64)(unsafe.Add(mBase, uint32(v1144)))
	*(*int64)(unsafe.Add(mBase, uint32(v1145+v1142))) = v1148
	v1152 = *(*int32)(unsafe.Add(mBase, uint32(v1144)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v1145+int32(-8)))) = v1152
	goto L291
L293:
	;
	goto L299
L299:
	;
	v1101 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	goto L292
L307:
	;
	goto L366
L308:
	;
	if v1211 != int32(4) {
		goto L307
	} else {
		goto L323
	}
L309:
	;
	v1205 = m.G398
	if v1171 != v1205 {
		goto L321
	} else {
		goto L322
	}
L310:
	;
	goto L314
L314:
	;
	v1168 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	v1171 = v1168 + int32(-16)
	goto L309
L321:
	;
	v1208 = *(*int32)(unsafe.Add(mBase, uint32(v1171)+8))
	v1211 = v1208
	goto L308
L322:
	;
	v1211 = int32(-1)
	goto L308
L323:
	;
	v1214 = m.G3
	v1215 = m.G64
	v1216 = F_copy_string_from_lua_stack(m, l2)
	mBase = m.M
	v1217 = m.ExcPending
	if v1217 != 0 {
		goto L18
	} else {
		goto L324
	}
L324:
	;
	v1228 = F_strlen(m, v1216)
	mBase = m.M
	if v1228 == int32(0) {
		goto L326
	} else {
		goto L327
	}
L325:
	;
	if v1216&int32(3) == int32(0) {
		v1308 = v1216
		goto L340
	} else {
		goto L341
	}
L326:
	;
	goto L325
L327:
	;
	v1237 = int32(0)
	goto L328
L328:
	;
	goto L331
L329:
	;
	goto L326
L330:
	;
	v1276 = v1237 + int32(1)
	if v1276 != v1228 {
		v1237 = v1276
		goto L328
	} else {
		goto L337
	}
L331:
	;
	v1244 = v1216 + v1237
	v1245 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1244))))
	v1254 = int32(0)
	goto L332
L332:
	;
	v1258 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1214+int32(_a823)+v1254))))
	if v1245&int32(255) != v1258 {
		goto L334
	} else {
		goto L335
	}
L333:
	;
	goto L330
L334:
	;
	v1264 = v1254 + int32(1)
	if v1264 != int32(2) {
		v1254 = v1264
		goto L332
	} else {
		goto L336
	}
L335:
	;
	v1261 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1214+int32(_a2640)+v1254))))
	*(*uint8)(unsafe.Add(mBase, uint32(v1244))) = uint8(v1261)
	goto L330
L336:
	;
	goto L333
L337:
	;
	goto L329
L338:
	;
	v1342 = *(*int32)(unsafe.Add(mBase, uint32(v1215)))
	v1343 = m.T0[v1342].(func(*base.Module, int32, int32, int32) int32)(m, l0, v1216, v1341)
	mBase = m.M
	v1344 = m.ExcPending
	if v1344 != 0 {
		goto L18
	} else {
		goto L354
	}
L339:
	;
	v1341 = v1333 - v1216
	goto L338
L340:
	;
	v1312 = v1308
	goto L348
L341:
	;
	v1294 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1216))))
	if v1294 != 0 {
		goto L342
	} else {
		goto L343
	}
L342:
	;
	v1297 = v1216
	goto L344
L343:
	;
	v1341 = v1216 - v1216
	goto L338
L344:
	;
	v1301 = v1297 + int32(1)
	if v1301&int32(3) == int32(0) {
		v1308 = v1301
		goto L340
	} else {
		goto L346
	}
L346:
	;
	v1306 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1301))))
	if v1306 != 0 {
		v1297 = v1301
		goto L344
	} else {
		goto L347
	}
L347:
	;
	v1333 = v1301
	goto L339
L348:
	;
	v1318 = *(*int32)(unsafe.Add(mBase, uint32(v1312)))
	v1321 = int32(-2139062144)
	if (int32(16843008)-v1318|v1318)&v1321 == v1321 {
		v1312 = v1312 + int32(4)
		goto L348
	} else {
		goto L350
	}
L349:
	;
	v1327 = v1312
	goto L351
L350:
	;
	goto L349
L351:
	;
	v1331 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1327))))
	if v1331 != 0 {
		v1327 = v1327 + int32(1)
		goto L351
	} else {
		goto L353
	}
L352:
	;
	v1333 = v1327
	goto L339
L353:
	;
	goto L352
L354:
	;
	v1345 = m.G11
	v1346 = *(*int32)(unsafe.Add(mBase, uint32(v1345)))
	m.T0[v1346].(func(*base.Module, int32))(m, v1216)
	mBase = m.M
	v1348 = m.ExcPending
	if v1348 != 0 {
		goto L18
	} else {
		goto L355
	}
L355:
	;
	goto L358
L356:
	;
	goto L16
L357:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2)+8)) = v1367 + int32(-32)
	goto L356
L358:
	;
	v1367 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	goto L357
L364:
	;
	v1405 = m.G3
	F_lua_pushstring(m, l2, v1405+int32(_a2643))
	mBase = m.M
	v1409 = m.ExcPending
	if v1409 != 0 {
		goto L18
	} else {
		goto L372
	}
L365:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2)+8)) = v1395 + int32(-16)
	goto L364
L366:
	;
	v1395 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	goto L365
L372:
	;
	goto L375
L373:
	;
	goto L392
L374:
	;
	v1466 = *(*int32)(unsafe.Add(mBase, uint32(v1427+int32(-32))))
	v1467 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	v1468 = int32(-16)
	v1470 = F_luaH_get(m, v1466, v1467+v1468)
	mBase = m.M
	v1471 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	v1474 = *(*int64)(unsafe.Add(mBase, uint32(v1470)))
	*(*int64)(unsafe.Add(mBase, uint32(v1471+v1468))) = v1474
	v1478 = *(*int32)(unsafe.Add(mBase, uint32(v1470)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v1471+int32(-8)))) = v1478
	goto L373
L375:
	;
	goto L381
L381:
	;
	v1427 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	goto L374
L389:
	;
	goto L503
L390:
	;
	if v1537 != int32(5) {
		goto L389
	} else {
		goto L405
	}
L391:
	;
	v1531 = m.G398
	if v1497 != v1531 {
		goto L403
	} else {
		goto L404
	}
L392:
	;
	goto L396
L396:
	;
	v1494 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	v1497 = v1494 + int32(-16)
	goto L391
L403:
	;
	v1534 = *(*int32)(unsafe.Add(mBase, uint32(v1497)+8))
	v1537 = v1534
	goto L390
L404:
	;
	v1537 = int32(-1)
	goto L390
L405:
	;
	v1540 = m.G3
	F_lua_pushstring(m, l2, v1540+int32(_a2644))
	mBase = m.M
	v1544 = m.ExcPending
	if v1544 != 0 {
		goto L18
	} else {
		goto L406
	}
L406:
	;
	goto L409
L407:
	;
	goto L426
L408:
	;
	v1601 = *(*int32)(unsafe.Add(mBase, uint32(v1562+int32(-32))))
	v1602 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	v1603 = int32(-16)
	v1605 = F_luaH_get(m, v1601, v1602+v1603)
	mBase = m.M
	v1606 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	v1609 = *(*int64)(unsafe.Add(mBase, uint32(v1605)))
	*(*int64)(unsafe.Add(mBase, uint32(v1606+v1603))) = v1609
	v1613 = *(*int32)(unsafe.Add(mBase, uint32(v1605)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v1606+int32(-8)))) = v1613
	goto L407
L409:
	;
	goto L415
L415:
	;
	v1562 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	goto L408
L423:
	;
	goto L495
L424:
	;
	if v1672 != int32(4) {
		goto L423
	} else {
		goto L439
	}
L425:
	;
	v1666 = m.G398
	if v1632 != v1666 {
		goto L437
	} else {
		goto L438
	}
L426:
	;
	goto L430
L430:
	;
	v1629 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	v1632 = v1629 + int32(-16)
	goto L425
L437:
	;
	v1669 = *(*int32)(unsafe.Add(mBase, uint32(v1632)+8))
	v1672 = v1669
	goto L424
L438:
	;
	v1672 = int32(-1)
	goto L424
L439:
	;
	v1675 = m.G3
	v1678 = F_lua_tolstring(m, l2, int32(-1), int32(0))
	mBase = m.M
	v1679 = m.ExcPending
	if v1679 != 0 {
		goto L18
	} else {
		goto L440
	}
L440:
	;
	F_lua_pushstring(m, l2, v1675+int32(_a2645))
	mBase = m.M
	v1683 = m.ExcPending
	if v1683 != 0 {
		goto L18
	} else {
		goto L441
	}
L441:
	;
	goto L444
L442:
	;
	goto L461
L443:
	;
	v1740 = *(*int32)(unsafe.Add(mBase, uint32(v1701+int32(-48))))
	v1741 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	v1742 = int32(-16)
	v1744 = F_luaH_get(m, v1740, v1741+v1742)
	mBase = m.M
	v1745 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	v1748 = *(*int64)(unsafe.Add(mBase, uint32(v1744)))
	*(*int64)(unsafe.Add(mBase, uint32(v1745+v1742))) = v1748
	v1752 = *(*int32)(unsafe.Add(mBase, uint32(v1744)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v1745+int32(-8)))) = v1752
	goto L442
L444:
	;
	goto L450
L450:
	;
	v1701 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	goto L443
L458:
	;
	goto L487
L459:
	;
	if v1811 != int32(4) {
		goto L458
	} else {
		goto L474
	}
L460:
	;
	v1805 = m.G398
	if v1771 != v1805 {
		goto L472
	} else {
		goto L473
	}
L461:
	;
	goto L465
L465:
	;
	v1768 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	v1771 = v1768 + int32(-16)
	goto L460
L472:
	;
	v1808 = *(*int32)(unsafe.Add(mBase, uint32(v1771)+8))
	v1811 = v1808
	goto L459
L473:
	;
	v1811 = int32(-1)
	goto L459
L474:
	;
	v1814 = m.G59
	v1818 = F_lua_tolstring(m, l2, int32(-1), v14+int32(16))
	mBase = m.M
	v1819 = m.ExcPending
	if v1819 != 0 {
		goto L18
	} else {
		goto L475
	}
L475:
	;
	v1820 = *(*int32)(unsafe.Add(mBase, uint32(v14)+16))
	v1821 = *(*int32)(unsafe.Add(mBase, uint32(v1814)))
	v1822 = m.T0[v1821].(func(*base.Module, int32, int32, int32, int32) int32)(m, l0, v1818, v1820, v1678)
	mBase = m.M
	v1823 = m.ExcPending
	if v1823 != 0 {
		goto L18
	} else {
		goto L476
	}
L476:
	;
	goto L479
L477:
	;
	goto L16
L478:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2)+8)) = v1842 + int32(-64)
	goto L477
L479:
	;
	v1842 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	goto L478
L485:
	;
	goto L423
L486:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2)+8)) = v1870 + int32(-16)
	goto L485
L487:
	;
	v1870 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	goto L486
L493:
	;
	goto L389
L494:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2)+8)) = v1900 + int32(-16)
	goto L493
L495:
	;
	v1900 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	goto L494
L501:
	;
	v1940 = m.G3
	F_lua_pushstring(m, l2, v1940+int32(_a2646))
	mBase = m.M
	v1944 = m.ExcPending
	if v1944 != 0 {
		goto L18
	} else {
		goto L509
	}
L502:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2)+8)) = v1930 + int32(-16)
	goto L501
L503:
	;
	v1930 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	goto L502
L509:
	;
	goto L512
L510:
	;
	goto L529
L511:
	;
	v2001 = *(*int32)(unsafe.Add(mBase, uint32(v1962+int32(-32))))
	v2002 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	v2003 = int32(-16)
	v2005 = F_luaH_get(m, v2001, v2002+v2003)
	mBase = m.M
	v2006 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	v2009 = *(*int64)(unsafe.Add(mBase, uint32(v2005)))
	*(*int64)(unsafe.Add(mBase, uint32(v2006+v2003))) = v2009
	v2013 = *(*int32)(unsafe.Add(mBase, uint32(v2005)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v2006+int32(-8)))) = v2013
	goto L510
L512:
	;
	goto L518
L518:
	;
	v1962 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	goto L511
L526:
	;
	goto L581
L527:
	;
	if v2072 != int32(5) {
		goto L526
	} else {
		goto L542
	}
L528:
	;
	v2066 = m.G398
	if v2032 != v2066 {
		goto L540
	} else {
		goto L541
	}
L529:
	;
	goto L533
L533:
	;
	v2029 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	v2032 = v2029 + int32(-16)
	goto L528
L540:
	;
	v2069 = *(*int32)(unsafe.Add(mBase, uint32(v2032)+8))
	v2072 = v2069
	goto L527
L541:
	;
	v2072 = int32(-1)
	goto L527
L542:
	;
	v2076 = m.G44
	v2077 = *(*int32)(unsafe.Add(mBase, uint32(v2076)))
	v2078 = m.T0[v2077].(func(*base.Module, int32, int32) int32)(m, l0, int32(-1))
	mBase = m.M
	v2079 = m.ExcPending
	if v2079 != 0 {
		goto L18
	} else {
		goto L543
	}
L543:
	;
	v2081 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v2081)+8)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l2)+8)) = v2081 + int32(16)
	goto L544
L544:
	;
	v2087 = int32(0)
	v2089 = F_lua_next(m, l2, int32(-2))
	mBase = m.M
	v2090 = m.ExcPending
	if v2090 != 0 {
		goto L18
	} else {
		goto L546
	}
L545:
	;
	v2189 = m.G50
	v2190 = *(*int32)(unsafe.Add(mBase, uint32(v2189)))
	m.T0[v2190].(func(*base.Module, int32, int32))(m, l0, v2182)
	mBase = m.M
	v2192 = m.ExcPending
	if v2192 != 0 {
		goto L18
	} else {
		goto L570
	}
L546:
	;
	if v2089 == int32(0) {
		v2182 = v2087
		goto L545
	} else {
		goto L547
	}
L547:
	;
	v2097 = v2087
	goto L548
L548:
	;
	goto L552
L549:
	;
	v2182 = v2174
	goto L545
L550:
	;
	F_luaReplyToServerReply(m, l0, l1, l2)
	mBase = m.M
	v2170 = m.ExcPending
	if v2170 != 0 {
		goto L18
	} else {
		goto L566
	}
L551:
	;
	v2160 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	v2161 = *(*int64)(unsafe.Add(mBase, uint32(v2124)))
	*(*int64)(unsafe.Add(mBase, uint32(v2160))) = v2161
	v2163 = *(*int32)(unsafe.Add(mBase, uint32(v2124)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v2160)+8)) = v2163
	v2165 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l2)+8)) = v2165 + int32(16)
	goto L550
L552:
	;
	goto L558
L558:
	;
	v2121 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	v2124 = v2121 + int32(-32)
	goto L551
L566:
	;
	F_luaReplyToServerReply(m, l0, l1, l2)
	mBase = m.M
	v2172 = m.ExcPending
	if v2172 != 0 {
		goto L18
	} else {
		goto L567
	}
L567:
	;
	v2174 = v2097 + int32(1)
	v2176 = F_lua_next(m, l2, int32(-2))
	mBase = m.M
	v2177 = m.ExcPending
	if v2177 != 0 {
		goto L18
	} else {
		goto L568
	}
L568:
	;
	if v2176 != 0 {
		v2097 = v2174
		goto L548
	} else {
		goto L569
	}
L569:
	;
	goto L549
L570:
	;
	goto L573
L571:
	;
	goto L16
L572:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2)+8)) = v2211 + int32(-32)
	goto L571
L573:
	;
	v2211 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	goto L572
L579:
	;
	v2249 = m.G3
	F_lua_pushstring(m, l2, v2249+int32(_a2432))
	mBase = m.M
	v2253 = m.ExcPending
	if v2253 != 0 {
		goto L18
	} else {
		goto L587
	}
L580:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2)+8)) = v2239 + int32(-16)
	goto L579
L581:
	;
	v2239 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	goto L580
L587:
	;
	goto L590
L588:
	;
	goto L607
L589:
	;
	v2310 = *(*int32)(unsafe.Add(mBase, uint32(v2271+int32(-32))))
	v2311 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	v2312 = int32(-16)
	v2314 = F_luaH_get(m, v2310, v2311+v2312)
	mBase = m.M
	v2315 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	v2318 = *(*int64)(unsafe.Add(mBase, uint32(v2314)))
	*(*int64)(unsafe.Add(mBase, uint32(v2315+v2312))) = v2318
	v2322 = *(*int32)(unsafe.Add(mBase, uint32(v2314)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v2315+int32(-8)))) = v2322
	goto L588
L590:
	;
	goto L596
L596:
	;
	v2271 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	goto L589
L604:
	;
	goto L666
L605:
	;
	if v2381 != int32(5) {
		goto L604
	} else {
		goto L620
	}
L606:
	;
	v2375 = m.G398
	if v2341 != v2375 {
		goto L618
	} else {
		goto L619
	}
L607:
	;
	goto L611
L611:
	;
	v2338 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	v2341 = v2338 + int32(-16)
	goto L606
L618:
	;
	v2378 = *(*int32)(unsafe.Add(mBase, uint32(v2341)+8))
	v2381 = v2378
	goto L605
L619:
	;
	v2381 = int32(-1)
	goto L605
L620:
	;
	v2385 = m.G45
	v2386 = *(*int32)(unsafe.Add(mBase, uint32(v2385)))
	v2387 = m.T0[v2386].(func(*base.Module, int32, int32) int32)(m, l0, int32(-1))
	mBase = m.M
	v2388 = m.ExcPending
	if v2388 != 0 {
		goto L18
	} else {
		goto L621
	}
L621:
	;
	v2390 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v2390)+8)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l2)+8)) = v2390 + int32(16)
	goto L622
L622:
	;
	v2396 = int32(0)
	v2398 = F_lua_next(m, l2, int32(-2))
	mBase = m.M
	v2399 = m.ExcPending
	if v2399 != 0 {
		goto L18
	} else {
		goto L624
	}
L623:
	;
	v2524 = m.G51
	v2525 = *(*int32)(unsafe.Add(mBase, uint32(v2524)))
	m.T0[v2525].(func(*base.Module, int32, int32))(m, l0, v2517)
	mBase = m.M
	v2527 = m.ExcPending
	if v2527 != 0 {
		goto L18
	} else {
		goto L655
	}
L624:
	;
	if v2398 == int32(0) {
		v2517 = v2396
		goto L623
	} else {
		goto L625
	}
L625:
	;
	v2406 = v2396
	goto L626
L626:
	;
	goto L630
L627:
	;
	v2517 = v2509
	goto L623
L628:
	;
	goto L638
L629:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2)+8)) = v2431 + int32(-16)
	goto L628
L630:
	;
	v2431 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	goto L629
L636:
	;
	F_luaReplyToServerReply(m, l0, l1, l2)
	mBase = m.M
	v2507 = m.ExcPending
	if v2507 != 0 {
		goto L18
	} else {
		goto L652
	}
L637:
	;
	v2497 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	v2498 = *(*int64)(unsafe.Add(mBase, uint32(v2461)))
	*(*int64)(unsafe.Add(mBase, uint32(v2497))) = v2498
	v2500 = *(*int32)(unsafe.Add(mBase, uint32(v2461)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v2497)+8)) = v2500
	v2502 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l2)+8)) = v2502 + int32(16)
	goto L636
L638:
	;
	goto L644
L644:
	;
	v2458 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	v2461 = v2458 + int32(-16)
	goto L637
L652:
	;
	v2509 = v2406 + int32(1)
	v2511 = F_lua_next(m, l2, int32(-2))
	mBase = m.M
	v2512 = m.ExcPending
	if v2512 != 0 {
		goto L18
	} else {
		goto L653
	}
L653:
	;
	if v2511 != 0 {
		v2406 = v2509
		goto L626
	} else {
		goto L654
	}
L654:
	;
	goto L627
L655:
	;
	goto L658
L656:
	;
	goto L16
L657:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2)+8)) = v2546 + int32(-32)
	goto L656
L658:
	;
	v2546 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	goto L657
L664:
	;
	v2585 = m.G43
	v2586 = *(*int32)(unsafe.Add(mBase, uint32(v2585)))
	v2587 = m.T0[v2586].(func(*base.Module, int32, int32) int32)(m, l0, int32(-1))
	mBase = m.M
	v2588 = m.ExcPending
	if v2588 != 0 {
		goto L18
	} else {
		goto L672
	}
L665:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2)+8)) = v2574 + int32(-16)
	goto L664
L666:
	;
	v2574 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	goto L665
L672:
	;
	v2591 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v2591)+8)) = int32(3)
	*(*float64)(unsafe.Add(mBase, uint32(v2591))) = float64(1)
	v2595 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l2)+8)) = v2595 + int32(16)
	goto L673
L673:
	;
	goto L676
L674:
	;
	v2669 = int32(0)
	goto L693
L675:
	;
	v2655 = *(*int32)(unsafe.Add(mBase, uint32(v2616+int32(-32))))
	v2656 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	v2657 = int32(-16)
	v2659 = F_luaH_get(m, v2655, v2656+v2657)
	mBase = m.M
	v2660 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	v2663 = *(*int64)(unsafe.Add(mBase, uint32(v2659)))
	*(*int64)(unsafe.Add(mBase, uint32(v2660+v2657))) = v2663
	v2667 = *(*int32)(unsafe.Add(mBase, uint32(v2659)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v2660+int32(-8)))) = v2667
	goto L674
L676:
	;
	goto L682
L682:
	;
	v2616 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	goto L675
L690:
	;
	goto L745
L691:
	;
	if v2727 == int32(0) {
		v2890 = v2669
		goto L690
	} else {
		goto L706
	}
L692:
	;
	v2721 = m.G398
	if v2687 != v2721 {
		goto L704
	} else {
		goto L705
	}
L693:
	;
	goto L697
L697:
	;
	v2684 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	v2687 = v2684 + int32(-16)
	goto L692
L704:
	;
	v2724 = *(*int32)(unsafe.Add(mBase, uint32(v2687)+8))
	v2727 = v2724
	goto L691
L705:
	;
	v2727 = int32(-1)
	goto L691
L706:
	;
	v2735 = v2669
	v2737 = int32(1)
	goto L707
L707:
	;
	F_luaReplyToServerReply(m, l0, l1, l2)
	mBase = m.M
	v2743 = m.ExcPending
	if v2743 != 0 {
		goto L18
	} else {
		goto L709
	}
L708:
	;
	v2890 = v2827
	goto L690
L709:
	;
	v2745 = v2737 + int32(1)
	v2748 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v2748)+8)) = int32(3)
	*(*float64)(unsafe.Add(mBase, uint32(v2748))) = base.F64_convert_i32_u(v2745)
	v2752 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l2)+8)) = v2752 + int32(16)
	goto L710
L710:
	;
	goto L713
L711:
	;
	v2827 = v2735 + int32(1)
	goto L729
L712:
	;
	v2812 = *(*int32)(unsafe.Add(mBase, uint32(v2773+int32(-32))))
	v2813 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	v2814 = int32(-16)
	v2816 = F_luaH_get(m, v2812, v2813+v2814)
	mBase = m.M
	v2817 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	v2820 = *(*int64)(unsafe.Add(mBase, uint32(v2816)))
	*(*int64)(unsafe.Add(mBase, uint32(v2817+v2814))) = v2820
	v2824 = *(*int32)(unsafe.Add(mBase, uint32(v2816)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v2817+int32(-8)))) = v2824
	goto L711
L713:
	;
	goto L719
L719:
	;
	v2773 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	goto L712
L727:
	;
	if v2885 != 0 {
		v2735 = v2827
		v2737 = v2745
		goto L707
	} else {
		goto L742
	}
L728:
	;
	v2879 = m.G398
	if v2845 != v2879 {
		goto L740
	} else {
		goto L741
	}
L729:
	;
	goto L733
L733:
	;
	v2842 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	v2845 = v2842 + int32(-16)
	goto L728
L740:
	;
	v2882 = *(*int32)(unsafe.Add(mBase, uint32(v2845)+8))
	v2885 = v2882
	goto L727
L741:
	;
	v2885 = int32(-1)
	goto L727
L742:
	;
	goto L708
L743:
	;
	v2925 = m.G49
	v2926 = *(*int32)(unsafe.Add(mBase, uint32(v2925)))
	m.T0[v2926].(func(*base.Module, int32, int32))(m, l0, v2890)
	mBase = m.M
	v2928 = m.ExcPending
	if v2928 != 0 {
		goto L18
	} else {
		goto L751
	}
L744:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2)+8)) = v2915 + int32(-16)
	goto L743
L745:
	;
	v2915 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	goto L744
L751:
	;
	goto L30
L752:
	;
	goto L30
L753:
	;
	goto L16
L754:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2)+8)) = v2962 + int32(-16)
	goto L753
L755:
	;
	v2962 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	goto L754
}
func F_luaServerBreakpointCommand(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v18 int32
	_ = v18
	v3 = m.G6
	v4 = *(*int32)(unsafe.Add(mBase, uint32(v3)))
	if v4 != 0 {
		v6 = int32(1)
		v8 = m.G6
		*(*int32)(unsafe.Add(mBase, uint32(v8)+268)) = v6
		v10 = v6
	} else {
		v10 = int32(0)
	}
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v12)+8)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v12))) = base.B2i32(v10 != int32(0))
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v18 + int32(16)
	return int32(1)
}
func F_luaServerReplicateCommandsCommand(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v10 int32
	_ = v10
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v4)+8)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v4))) = int32(1)
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v10 + int32(16)
	return int32(1)
}
func F_luaSetErrorMetatable(m *base.Module, l0 int32) {
	var v2 int32
	_ = v2
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	v2 = int32(0)
	F_lua_createtable(m, l0, v2, v2)
	v5 = m.ExcPending
	if v5 != 0 {
		return
	} else {
		v6 = m.G5
		F_lua_pushcclosure(m, l0, v6+int32(1185), int32(0))
		v11 = m.ExcPending
		if v11 != 0 {
			return
		} else {
			v13 = m.G3
			F_lua_setfield(m, l0, int32(-2), v13+int32(_a2626))
			v17 = m.ExcPending
			if v17 != 0 {
				return
			} else {
				v19 = F_lua_setmetatable(m, l0, int32(-2))
				v20 = m.ExcPending
				if v20 != 0 {
					return
				} else {
					return
				}
			}
		}
	}
}
func F_lua_close(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v72 int32
	_ = v72
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v4 = *(*int32)(unsafe.Add(mBase, uint32(v3)+112))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(v4)+32))
	F_luaF_close(m, v4, v5)
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
	v15 = *(*int32)(unsafe.Add(mBase, uint32(v4)+16))
	v16 = *(*int32)(unsafe.Add(mBase, uint32(v15)+112))
	v17 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
	if v17 != 0 {
		goto L4
	} else {
		goto L5
	}
L3:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v4)+116)) = int32(0)
	goto L22
L4:
	;
	v23 = v16
	v24 = v17
	goto L6
L5:
	;
	goto L3
L6:
	;
	v28 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24)+5)))
	if v28&int32(8) == int32(0) {
		goto L9
	} else {
		goto L10
	}
L7:
	;
	goto L3
L8:
	;
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v68)))
	if v72 != 0 {
		v23 = v68
		v24 = v72
		goto L6
	} else {
		goto L21
	}
L9:
	;
	if v28&int32(3)|int32(1) != 0 {
		goto L11
	} else {
		goto L12
	}
L10:
	;
	v68 = v24
	goto L8
L11:
	;
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v24)+8))
	if v36 == int32(0) {
		v47 = v28
		goto L14
	} else {
		goto L15
	}
L12:
	;
	v68 = v24
	goto L8
L13:
	;
	v53 = v46 | int32(8)
	*(*uint8)(unsafe.Add(mBase, uint32(v24)+5)) = uint8(v53)
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v24)))
	*(*int32)(unsafe.Add(mBase, uint32(v23))) = v56
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v15)+48))
	if v61 != 0 {
		goto L19
	} else {
		goto L20
	}
L14:
	;
	v50 = v47 | int32(8)
	*(*uint8)(unsafe.Add(mBase, uint32(v24)+5)) = uint8(v50)
	v68 = v24
	goto L8
L15:
	;
	v39 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v36)+6)))
	if v39&int32(4) != 0 {
		v47 = v28
		goto L14
	} else {
		goto L16
	}
L16:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v4)+16))
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v43)+196))
	v45 = F_luaT_gettm(m, v36, int32(2), v44)
	mBase = m.M
	v46 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24)+5)))
	if v45 != 0 {
		goto L13
	} else {
		goto L17
	}
L17:
	;
	v47 = v46
	goto L14
L18:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+48)) = v24
	v68 = v23
	goto L8
L19:
	;
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v61)))
	*(*int32)(unsafe.Add(mBase, uint32(v24))) = v63
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v15)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v65))) = v24
	goto L18
L20:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24))) = v24
	goto L18
L21:
	;
	goto L7
L22:
	;
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v4)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v4)+20)) = v86
	v88 = *(*int32)(unsafe.Add(mBase, uint32(v86)))
	*(*int32)(unsafe.Add(mBase, uint32(v4)+8)) = v88
	*(*int32)(unsafe.Add(mBase, uint32(v4)+12)) = v88
	v91 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v4)+52)) = v91
	v93 = m.G5
	v97 = F_luaD_rawrunprotected(m, v4, v93+int32(1233), v91)
	mBase = m.M
	v98 = m.ExcPending
	if v98 != 0 {
		goto L1
	} else {
		goto L24
	}
L23:
	;
	F_close_state(m, v4)
	mBase = m.M
	v100 = m.ExcPending
	if v100 != 0 {
		goto L1
	} else {
		goto L26
	}
L24:
	;
	if v97 != 0 {
		goto L22
	} else {
		goto L25
	}
L25:
	;
	goto L23
L26:
	;
	return
}
func F_lua_error(m *base.Module, l0 int32) int32 {
	var v5 int32
	_ = v5
	F_luaG_errormsg(m, l0)
	v5 = m.ExcPending
	if v5 != 0 {
		return int32(0)
	} else {
		return int32(0)
	}
}
func F_lua_gc(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v12 int32
	_ = v12
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v61 int32
	_ = v61
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	switch l1 {
	case 0:
		goto L9
	case 1:
		goto L8
	case 2:
		goto L7
	case 3:
		goto L6
	case 4:
		goto L5
	case 5:
		goto L4
	case 6:
		goto L3
	case 7:
		goto L2
	default:
		v61 = int32(-1)
		goto L1
	}
L1:
	;
	return v61
L2:
	;
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v6)+84))
	*(*int32)(unsafe.Add(mBase, uint32(v6)+84)) = l2
	v61 = v55
	goto L1
L3:
	;
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v6)+80))
	*(*int32)(unsafe.Add(mBase, uint32(v6)+80)) = l2
	return v52
L4:
	;
	v30 = int32(0)
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v6)+68))
	v35 = v32 - l2<<(uint(int32(10))%32)
	if base.Ui32(v32) < base.Ui32(v35) {
		goto L12
	} else {
		goto L13
	}
L5:
	;
	v26 = *(*int32)(unsafe.Add(mBase, uint32(v6)+68))
	return v26 & int32(1023)
L6:
	;
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v6)+68))
	return int32(base.Ui32(v22) >> (uint(int32(10)) % 32))
L7:
	;
	F_luaC_fullgc(m, l0)
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		goto L10
	} else {
		goto L11
	}
L8:
	;
	v12 = *(*int32)(unsafe.Add(mBase, uint32(v6)+68))
	*(*int32)(unsafe.Add(mBase, uint32(v6)+64)) = v12
	return int32(0)
L9:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v6)+64)) = int32(-3)
	return int32(0)
L10:
	;
	return int32(0)
L11:
	;
	return int32(0)
L12:
	;
	v37 = v30
	goto L14
L13:
	;
	v37 = v35
	goto L14
L14:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v6)+64)) = v37
	goto L15
L15:
	;
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v6)+64))
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v6)+68))
	if base.Ui32(v45) < base.Ui32(v44) {
		v61 = v30
		goto L1
	} else {
		goto L17
	}
L16:
	;
	return int32(1)
L17:
	;
	F_luaC_step(m, l0)
	mBase = m.M
	v48 = m.ExcPending
	if v48 != 0 {
		goto L10
	} else {
		goto L18
	}
L18:
	;
	v49 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6)+21)))
	if v49 != 0 {
		goto L15
	} else {
		goto L19
	}
L19:
	;
	goto L16
}
func F_lua_getallocf(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if l1 == int32(0) {
	} else {
		v6 = *(*int32)(unsafe.Add(mBase, uint32(v3)+16))
		*(*int32)(unsafe.Add(mBase, uint32(l1))) = v6
	}
	v8 = *(*int32)(unsafe.Add(mBase, uint32(v3)+12))
	return v8
}
func F_lua_gethookcount(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	return v2
}
func F_lua_gethookmask(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	v2 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+56)))
	return v2
}
func F_lua_getstack(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v54 int32
	_ = v54
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if l1 < int32(1) {
		v32 = l1
		v34 = v7
		goto L2
	} else {
		goto L3
	}
L1:
	;
	return v54
L2:
	;
	if v32 != 0 {
		v45 = int32(0)
		goto L10
	} else {
		goto L11
	}
L3:
	;
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v13 = l1
	v15 = v7
	goto L4
L4:
	;
	if base.Ui32(v15) <= base.Ui32(v10) {
		v54 = int32(0)
		goto L1
	} else {
		goto L6
	}
L5:
	;
	v32 = v26
	v34 = v28
	goto L2
L6:
	;
	v20 = v13 + int32(-1)
	v21 = *(*int32)(unsafe.Add(mBase, uint32(v15)+4))
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v21)))
	v23 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22)+6)))
	if v23 != 0 {
		v26 = v20
		goto L7
	} else {
		goto L8
	}
L7:
	;
	v28 = v15 + int32(-24)
	if int32(0) < v26 {
		v13 = v26
		v15 = v28
		goto L4
	} else {
		goto L9
	}
L8:
	;
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v15)+20))
	v26 = v20 - v24
	goto L7
L9:
	;
	goto L5
L10:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2)+96)) = v45
	v54 = int32(1)
	goto L1
L11:
	;
	v39 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	if base.Ui32(v34) <= base.Ui32(v39) {
		v54 = int32(0)
		goto L1
	} else {
		goto L12
	}
L12:
	;
	v43 = base.I32_div_s(v34-v39, int32(24))
	v45 = v43
	goto L10
}
func F_lua_gettop(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	return (v2 - v3) >> (uint(int32(4)) % 32)
}
func F_lua_insert(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
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
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v65 int32
	_ = v65
	var v66 int64
	_ = v66
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v78 int64
	_ = v78
	var v80 int32
	_ = v80
	if l1 < int32(1) {
		if l1 < int32(-9999) {
			switch l1 + int32(10002) {
			case 0:
				v57 = l0 + int32(72)
			case 1:
				v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
				v29 = *(*int32)(unsafe.Add(mBase, uint32(v28)+4))
				v30 = *(*int32)(unsafe.Add(mBase, uint32(v29)))
				v31 = *(*int32)(unsafe.Add(mBase, uint32(v30)+12))
				*(*int32)(unsafe.Add(mBase, uint32(l0)+96)) = int32(5)
				*(*int32)(unsafe.Add(mBase, uint32(l0)+88)) = v31
				v57 = l0 + int32(88)
			case 2:
				v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
				v57 = v25 + int32(96)
			default:
				v39 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
				v40 = *(*int32)(unsafe.Add(mBase, uint32(v39)+4))
				v41 = *(*int32)(unsafe.Add(mBase, uint32(v40)))
				v42 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v41)+7)))
				v43 = m.G398
				if base.Ui32(v42) < base.Ui32(int32(-10002)-l1) {
					v54 = v43
				} else {
					v54 = v41 + (int32(-10003)-l1)<<(uint(int32(4))%32) + int32(24)
				}
				v57 = v54
			}
		} else {
			v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			v57 = v19 + l1<<(uint(int32(4))%32)
		}
	} else {
		v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
		v13 = v8 + l1<<(uint(int32(4))%32) + int32(-16)
		v14 = m.G398
		if base.Ui32(v13) < base.Ui32(v7) {
			v16 = v13
		} else {
			v16 = v14
		}
		v57 = v16
	}
	v58 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if base.Ui32(v58) <= base.Ui32(v57) {
		v75 = v58
	} else {
		v61 = v58
		for {
			v65 = v61 + int32(-16)
			v66 = *(*int64)(unsafe.Add(mBase, uint32(v65)))
			*(*int64)(unsafe.Add(mBase, uint32(v61))) = v66
			v70 = *(*int32)(unsafe.Add(mBase, uint32(v61+int32(-8))))
			*(*int32)(unsafe.Add(mBase, uint32(v61)+8)) = v70
			if base.Ui32(v57) < base.Ui32(v65) {
				v61 = v65
				continue
			} else {
				break
			}
			break
		}
		v73 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		v75 = v73
	}
	v78 = *(*int64)(unsafe.Add(mBase, uint32(v75)))
	*(*int64)(unsafe.Add(mBase, uint32(v57))) = v78
	v80 = *(*int32)(unsafe.Add(mBase, uint32(v75)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v57)+8)) = v80
	return
}
func F_lua_lessthan(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v26 int32
	_ = v26
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
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v115 int32
	_ = v115
	var v117 int32
	_ = v117
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	if l1 < int32(1) {
		if l1 < int32(-9999) {
			switch l1 + int32(10002) {
			case 0:
				v56 = l0 + int32(72)
			case 1:
				v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
				v30 = *(*int32)(unsafe.Add(mBase, uint32(v29)+4))
				v31 = *(*int32)(unsafe.Add(mBase, uint32(v30)))
				v32 = *(*int32)(unsafe.Add(mBase, uint32(v31)+12))
				*(*int32)(unsafe.Add(mBase, uint32(l0)+96)) = int32(5)
				*(*int32)(unsafe.Add(mBase, uint32(l0)+88)) = v32
				v56 = l0 + int32(88)
			case 2:
				v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
				v56 = v26 + int32(96)
			default:
				v40 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
				v41 = *(*int32)(unsafe.Add(mBase, uint32(v40)+4))
				v42 = *(*int32)(unsafe.Add(mBase, uint32(v41)))
				v43 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v42)+7)))
				v44 = m.G398
				if base.Ui32(v43) < base.Ui32(int32(-10002)-l1) {
					v55 = v44
				} else {
					v55 = v42 + (int32(-10003)-l1)<<(uint(int32(4))%32) + int32(24)
				}
				v56 = v55
			}
		} else {
			v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			v56 = v20 + l1<<(uint(int32(4))%32)
		}
	} else {
		v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
		v14 = v9 + l1<<(uint(int32(4))%32) + int32(-16)
		v15 = m.G398
		if base.Ui32(v14) < base.Ui32(v8) {
			v17 = v14
		} else {
			v17 = v15
		}
		v56 = v17
	}
	if l2 < int32(1) {
		if l2 < int32(-9999) {
			switch l2 + int32(10002) {
			case 0:
				v110 = l0 + int32(72)
			case 1:
				v82 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
				v83 = *(*int32)(unsafe.Add(mBase, uint32(v82)+4))
				v84 = *(*int32)(unsafe.Add(mBase, uint32(v83)))
				v85 = *(*int32)(unsafe.Add(mBase, uint32(v84)+12))
				*(*int32)(unsafe.Add(mBase, uint32(l0)+96)) = int32(5)
				*(*int32)(unsafe.Add(mBase, uint32(l0)+88)) = v85
				v110 = l0 + int32(88)
			case 2:
				v79 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
				v110 = v79 + int32(96)
			default:
				v93 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
				v94 = *(*int32)(unsafe.Add(mBase, uint32(v93)+4))
				v95 = *(*int32)(unsafe.Add(mBase, uint32(v94)))
				v96 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v95)+7)))
				v97 = m.G398
				if base.Ui32(v96) < base.Ui32(int32(-10002)-l2) {
					v108 = v97
				} else {
					v108 = v95 + (int32(-10003)-l2)<<(uint(int32(4))%32) + int32(24)
				}
				v110 = v108
			}
		} else {
			v73 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			v110 = v73 + l2<<(uint(int32(4))%32)
		}
	} else {
		v61 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		v62 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
		v67 = v62 + l2<<(uint(int32(4))%32) + int32(-16)
		v68 = m.G398
		if base.Ui32(v67) < base.Ui32(v61) {
			v70 = v67
		} else {
			v70 = v68
		}
		v110 = v70
	}
	v112 = int32(0)
	v113 = m.G398
	if v56 == v113 {
		v121 = v112
		return v121
	} else {
		v115 = m.G398
		if v110 == v115 {
			v121 = v112
			return v121
		} else {
			v117 = F_luaV_lessthan(m, l0, v56, v110)
			mBase = m.M
			v120 = m.ExcPending
			if v120 != 0 {
				return int32(0)
			} else {
				v121 = v117
				return v121
			}
		}
	}
}
func F_lua_newstate(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v28 int64
	_ = v28
	var v30 int32
	_ = v30
	var v34 int32
	_ = v34
	var v50 int32
	_ = v50
	var v104 int32
	_ = v104
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v112 int32
	_ = v112
	var v114 int32
	_ = v114
	v3 = int32(0)
	v9 = m.T0[l0].(func(*base.Module, int32, int32, int32, int32) int32)(m, l1, v3, v3, int32(376))
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		return int32(0)
	} else {
		if v9 == int32(0) {
			v114 = v3
			return v114
		} else {
			v15 = int32(33)
			*(*uint8)(unsafe.Add(mBase, uint32(v9)+140)) = uint8(v15)
			v17 = int32(8)
			*(*uint8)(unsafe.Add(mBase, uint32(v9)+4)) = uint8(v17)
			v19 = int32(0)
			*(*int32)(unsafe.Add(mBase, uint32(v9))) = v19
			*(*int32)(unsafe.Add(mBase, uint32(v9)+32)) = v19
			*(*int32)(unsafe.Add(mBase, uint32(v9)+68)) = v19
			*(*int32)(unsafe.Add(mBase, uint32(v9)+104)) = v19
			v28 = int64(0)
			*(*int64)(unsafe.Add(mBase, uint32(v9)+60)) = v28
			v30 = int32(256)
			*(*uint16)(unsafe.Add(mBase, uint32(v9)+56)) = uint16(v30)
			*(*int64)(unsafe.Add(mBase, uint32(v9)+40)) = v28
			v34 = int32(97)
			*(*uint16)(unsafe.Add(mBase, uint32(v9)+5)) = uint16(v34)
			*(*int64)(unsafe.Add(mBase, uint32(v9)+112)) = v28
			*(*int64)(unsafe.Add(mBase, uint32(v9)+20)) = v28
			*(*int32)(unsafe.Add(mBase, uint32(v9)+80)) = v19
			*(*int32)(unsafe.Add(mBase, uint32(v9)+136)) = l1
			*(*int32)(unsafe.Add(mBase, uint32(v9)+132)) = l0
			*(*int64)(unsafe.Add(mBase, uint32(v9)+48)) = v28
			*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = v9 + int32(120)
			v50 = v9 + int32(240)
			*(*int32)(unsafe.Add(mBase, uint32(v9)+260)) = v50
			*(*int32)(unsafe.Add(mBase, uint32(v9)+256)) = v50
			*(*int32)(unsafe.Add(mBase, uint32(v9)+232)) = v9
			*(*int32)(unsafe.Add(mBase, uint32(v9)+128)) = v19
			*(*int64)(unsafe.Add(mBase, uint32(v9)+184)) = int64(1614907703296)
			*(*int64)(unsafe.Add(mBase, uint32(v9)+120)) = v28
			*(*int32)(unsafe.Add(mBase, uint32(v9)+224)) = v19
			*(*int32)(unsafe.Add(mBase, uint32(v9)+172)) = v19
			*(*int32)(unsafe.Add(mBase, uint32(v9)+180)) = v19
			*(*int32)(unsafe.Add(mBase, uint32(v9)+208)) = v19
			*(*uint8)(unsafe.Add(mBase, uint32(v9)+141)) = uint8(v19)
			*(*int32)(unsafe.Add(mBase, uint32(v9)+144)) = v19
			*(*int32)(unsafe.Add(mBase, uint32(v9)+152)) = v9 + int32(148)
			*(*int64)(unsafe.Add(mBase, uint32(v9)+156)) = v28
			*(*int64)(unsafe.Add(mBase, uint32(v9+int32(164)))) = v28
			*(*int64)(unsafe.Add(mBase, uint32(v9)+200)) = int64(858993459400)
			*(*int32)(unsafe.Add(mBase, uint32(v9)+196)) = v19
			*(*int32)(unsafe.Add(mBase, uint32(v9+int32(304)))) = v19
			*(*int64)(unsafe.Add(mBase, uint32(v9+int32(296)))) = v28
			*(*int64)(unsafe.Add(mBase, uint32(v9+int32(288)))) = v28
			*(*int64)(unsafe.Add(mBase, uint32(v9+int32(280)))) = v28
			*(*int64)(unsafe.Add(mBase, uint32(v9)+272)) = v28
			*(*int32)(unsafe.Add(mBase, uint32(v9)+148)) = v9
			v104 = m.G5
			v108 = F_luaD_rawrunprotected(m, v9, v104+int32(1232), v19)
			mBase = m.M
			v109 = m.ExcPending
			if v109 != 0 {
				return int32(0)
			} else {
				if v108 != 0 {
					F_close_state(m, v9)
					mBase = m.M
					v112 = m.ExcPending
					if v112 != 0 {
						return int32(0)
					} else {
						v114 = v19
						return v114
					}
				} else {
					return v9
				}
			}
		}
	}
}
func F_lua_newthread(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v18 int32
	_ = v18
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(v4)+68))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(v4)+64))
	if base.Ui32(v5) < base.Ui32(v6) {
		v12 = F_luaE_newthread(m, l0)
		mBase = m.M
		v13 = m.ExcPending
		if v13 != 0 {
			return int32(0)
		} else {
			v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			*(*int32)(unsafe.Add(mBase, uint32(v14)+8)) = int32(8)
			*(*int32)(unsafe.Add(mBase, uint32(v14))) = v12
			v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v18 + int32(16)
			return v12
		}
	} else {
		F_luaC_step(m, l0)
		mBase = m.M
		v11 = m.ExcPending
		if v11 != 0 {
			return int32(0)
		} else {
			v12 = F_luaE_newthread(m, l0)
			mBase = m.M
			v13 = m.ExcPending
			if v13 != 0 {
				return int32(0)
			} else {
				v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				*(*int32)(unsafe.Add(mBase, uint32(v14)+8)) = int32(8)
				*(*int32)(unsafe.Add(mBase, uint32(v14))) = v12
				v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v18 + int32(16)
				return v12
			}
		}
	}
}
func F_lua_pcall(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v80 int32
	_ = v80
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	v9 = m.G0
	v11 = v9 - int32(16)
	m.G0 = v11
	if l3 != 0 {
		if l3 < int32(1) {
			if base.Ui32(l3) < base.Ui32(int32(-9999)) {
				switch l3 + int32(10002) {
				case 0:
					v66 = l0 + int32(72)
				case 1:
					v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
					v39 = *(*int32)(unsafe.Add(mBase, uint32(v38)+4))
					v40 = *(*int32)(unsafe.Add(mBase, uint32(v39)))
					v41 = *(*int32)(unsafe.Add(mBase, uint32(v40)+12))
					*(*int32)(unsafe.Add(mBase, uint32(l0)+96)) = int32(5)
					*(*int32)(unsafe.Add(mBase, uint32(l0)+88)) = v41
					v66 = l0 + int32(88)
				case 2:
					v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
					v66 = v35 + int32(96)
				default:
					v49 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
					v50 = *(*int32)(unsafe.Add(mBase, uint32(v49)+4))
					v51 = *(*int32)(unsafe.Add(mBase, uint32(v50)))
					v52 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v51)+7)))
					v53 = m.G398
					if base.Ui32(v52) < base.Ui32(int32(-10002)-l3) {
						v64 = v53
					} else {
						v64 = v51 + (int32(-10003)-l3)<<(uint(int32(4))%32) + int32(24)
					}
					v66 = v64
				}
			} else {
				v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				v66 = v29 + l3<<(uint(int32(4))%32)
			}
		} else {
			v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
			v23 = v18 + l3<<(uint(int32(4))%32) + int32(-16)
			v24 = m.G398
			if base.Ui32(v23) < base.Ui32(v17) {
				v26 = v23
			} else {
				v26 = v24
			}
			v66 = v26
		}
		v68 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
		v70 = v68
		v71 = v66 - v68
	} else {
		v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
		v70 = v13
		v71 = int32(0)
	}
	v73 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+12)) = l2
	v75 = m.G5
	v80 = v73 + (l1^int32(-1))<<(uint(int32(4))%32)
	*(*int32)(unsafe.Add(mBase, uint32(v11)+8)) = v80
	v87 = F_luaD_pcall(m, l0, v75+int32(1229), v11+int32(8), v80-v70, v71)
	mBase = m.M
	v90 = m.ExcPending
	if v90 != 0 {
		return int32(0)
	} else {
		if l2 != int32(-1) {
		} else {
			v93 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			v94 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
			v95 = *(*int32)(unsafe.Add(mBase, uint32(v94)+8))
			if base.Ui32(v93) < base.Ui32(v95) {
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v94)+8)) = v93
			}
		}
		m.G0 = v11 + int32(16)
		return v87
	}
}
func F_lua_pushboolean(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v10 int32
	_ = v10
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v4)+8)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v4))) = base.B2i32(l1 != int32(0))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v10 + int32(16)
	return
}
func F_lua_pushinteger(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v9 int32
	_ = v9
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v4)+8)) = int32(3)
	*(*float64)(unsafe.Add(mBase, uint32(v4))) = base.F64_convert_i32_s(l1)
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v9 + int32(16)
	return
}
func F_lua_pushlightuserdata(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v8 int32
	_ = v8
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v4)+8)) = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v4))) = l1
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v8 + int32(16)
	return
}
func F_lua_pushlstring(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v17 int32
	_ = v17
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(v5)+68))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(v5)+64))
	if base.Ui32(v6) < base.Ui32(v7) {
		v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		v12 = F_luaS_newlstr(m, l0, l1, l2)
		mBase = m.M
		v13 = m.ExcPending
		if v13 != 0 {
			return
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v11)+8)) = int32(4)
			*(*int32)(unsafe.Add(mBase, uint32(v11))) = v12
			v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v17 + int32(16)
			return
		}
	} else {
		F_luaC_step(m, l0)
		mBase = m.M
		v10 = m.ExcPending
		if v10 != 0 {
			return
		} else {
			v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			v12 = F_luaS_newlstr(m, l0, l1, l2)
			mBase = m.M
			v13 = m.ExcPending
			if v13 != 0 {
				return
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v11)+8)) = int32(4)
				*(*int32)(unsafe.Add(mBase, uint32(v11))) = v12
				v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v17 + int32(16)
				return
			}
		}
	}
}
func F_lua_pushnil(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v3)+8)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v3 + int32(16)
	return
}
func F_lua_pushstring(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v25 int32
	_ = v25
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v51 int32
	_ = v51
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v78 int32
	_ = v78
	if l1 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	if l1&int32(3) == int32(0) {
		v32 = l1
		goto L5
	} else {
		goto L6
	}
L2:
	;
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v5)+8)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v5 + int32(16)
	return
L3:
	;
	v66 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v66)+68))
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v66)+64))
	if base.Ui32(v67) < base.Ui32(v68) {
		goto L19
	} else {
		goto L20
	}
L4:
	;
	v65 = v57 - l1
	goto L3
L5:
	;
	v36 = v32
	goto L13
L6:
	;
	v18 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	if v18 != 0 {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	v21 = l1
	goto L9
L8:
	;
	v65 = l1 - l1
	goto L3
L9:
	;
	v25 = v21 + int32(1)
	if v25&int32(3) == int32(0) {
		v32 = v25
		goto L5
	} else {
		goto L11
	}
L11:
	;
	v30 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v25))))
	if v30 != 0 {
		v21 = v25
		goto L9
	} else {
		goto L12
	}
L12:
	;
	v57 = v25
	goto L4
L13:
	;
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v36)))
	v45 = int32(-2139062144)
	if (int32(16843008)-v42|v42)&v45 == v45 {
		v36 = v36 + int32(4)
		goto L13
	} else {
		goto L15
	}
L14:
	;
	v51 = v36
	goto L16
L15:
	;
	goto L14
L16:
	;
	v55 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v51))))
	if v55 != 0 {
		v51 = v51 + int32(1)
		goto L16
	} else {
		goto L18
	}
L17:
	;
	v57 = v51
	goto L4
L18:
	;
	goto L17
L19:
	;
	v72 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v73 = F_luaS_newlstr(m, l0, l1, v65)
	mBase = m.M
	v74 = m.ExcPending
	if v74 != 0 {
		goto L21
	} else {
		goto L23
	}
L20:
	;
	F_luaC_step(m, l0)
	mBase = m.M
	v71 = m.ExcPending
	if v71 != 0 {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	return
L22:
	;
	goto L19
L23:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v72)+8)) = int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v72))) = v73
	v78 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v78 + int32(16)
	return
}
func F_lua_pushvfstring(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(v5)+68))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(v5)+64))
	if base.Ui32(v6) < base.Ui32(v7) {
		v13 = F_luaO_pushvfstring(m, l0, l1, l2)
		mBase = m.M
		v14 = m.ExcPending
		if v14 != 0 {
			return int32(0)
		} else {
			return v13
		}
	} else {
		F_luaC_step(m, l0)
		mBase = m.M
		v12 = m.ExcPending
		if v12 != 0 {
			return int32(0)
		} else {
			v13 = F_luaO_pushvfstring(m, l0, l1, l2)
			mBase = m.M
			v14 = m.ExcPending
			if v14 != 0 {
				return int32(0)
			} else {
				return v13
			}
		}
	}
}
func F_lua_rawequal(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v26 int32
	_ = v26
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
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v115 int32
	_ = v115
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v124 float64
	_ = v124
	var v125 float64
	_ = v125
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v136 int32
	_ = v136
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	if l1 < int32(1) {
		if l1 < int32(-9999) {
			switch l1 + int32(10002) {
			case 0:
				v56 = l0 + int32(72)
			case 1:
				v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
				v30 = *(*int32)(unsafe.Add(mBase, uint32(v29)+4))
				v31 = *(*int32)(unsafe.Add(mBase, uint32(v30)))
				v32 = *(*int32)(unsafe.Add(mBase, uint32(v31)+12))
				*(*int32)(unsafe.Add(mBase, uint32(l0)+96)) = int32(5)
				*(*int32)(unsafe.Add(mBase, uint32(l0)+88)) = v32
				v56 = l0 + int32(88)
			case 2:
				v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
				v56 = v26 + int32(96)
			default:
				v40 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
				v41 = *(*int32)(unsafe.Add(mBase, uint32(v40)+4))
				v42 = *(*int32)(unsafe.Add(mBase, uint32(v41)))
				v43 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v42)+7)))
				v44 = m.G398
				if base.Ui32(v43) < base.Ui32(int32(-10002)-l1) {
					v55 = v44
				} else {
					v55 = v42 + (int32(-10003)-l1)<<(uint(int32(4))%32) + int32(24)
				}
				v56 = v55
			}
		} else {
			v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			v56 = v20 + l1<<(uint(int32(4))%32)
		}
	} else {
		v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
		v14 = v9 + l1<<(uint(int32(4))%32) + int32(-16)
		v15 = m.G398
		if base.Ui32(v14) < base.Ui32(v8) {
			v17 = v14
		} else {
			v17 = v15
		}
		v56 = v17
	}
	if l2 < int32(1) {
		if l2 < int32(-9999) {
			switch l2 + int32(10002) {
			case 0:
				v110 = l0 + int32(72)
			case 1:
				v82 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
				v83 = *(*int32)(unsafe.Add(mBase, uint32(v82)+4))
				v84 = *(*int32)(unsafe.Add(mBase, uint32(v83)))
				v85 = *(*int32)(unsafe.Add(mBase, uint32(v84)+12))
				*(*int32)(unsafe.Add(mBase, uint32(l0)+96)) = int32(5)
				*(*int32)(unsafe.Add(mBase, uint32(l0)+88)) = v85
				v110 = l0 + int32(88)
			case 2:
				v79 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
				v110 = v79 + int32(96)
			default:
				v93 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
				v94 = *(*int32)(unsafe.Add(mBase, uint32(v93)+4))
				v95 = *(*int32)(unsafe.Add(mBase, uint32(v94)))
				v96 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v95)+7)))
				v97 = m.G398
				if base.Ui32(v96) < base.Ui32(int32(-10002)-l2) {
					v108 = v97
				} else {
					v108 = v95 + (int32(-10003)-l2)<<(uint(int32(4))%32) + int32(24)
				}
				v110 = v108
			}
		} else {
			v73 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			v110 = v73 + l2<<(uint(int32(4))%32)
		}
	} else {
		v61 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		v62 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
		v67 = v62 + l2<<(uint(int32(4))%32) + int32(-16)
		v68 = m.G398
		if base.Ui32(v67) < base.Ui32(v61) {
			v70 = v67
		} else {
			v70 = v68
		}
		v110 = v70
	}
	v112 = int32(0)
	v113 = m.G398
	if v56 == v113 {
		v139 = v112
	} else {
		v115 = m.G398
		if v110 == v115 {
			v139 = v112
		} else {
			v119 = *(*int32)(unsafe.Add(mBase, uint32(v56)+8))
			v120 = *(*int32)(unsafe.Add(mBase, uint32(v110)+8))
			if v119 == v120 {
				switch v119 {
				case 0:
					v136 = int32(1)
					v138 = v136
				case 1:
					v127 = *(*int32)(unsafe.Add(mBase, uint32(v56)))
					v128 = *(*int32)(unsafe.Add(mBase, uint32(v110)))
					v138 = base.B2i32(v127 == v128)
				case 2:
					v130 = *(*int32)(unsafe.Add(mBase, uint32(v56)))
					v131 = *(*int32)(unsafe.Add(mBase, uint32(v110)))
					v138 = base.B2i32(v130 == v131)
				case 3:
					v124 = *(*float64)(unsafe.Add(mBase, uint32(v56)))
					v125 = *(*float64)(unsafe.Add(mBase, uint32(v110)))
					v138 = base.F64_eq(v124, v125)
				default:
					v133 = *(*int32)(unsafe.Add(mBase, uint32(v56)))
					v134 = *(*int32)(unsafe.Add(mBase, uint32(v110)))
					v136 = base.B2i32(v133 == v134)
					v138 = v136
				}
			} else {
				v138 = int32(0)
			}
			v139 = v138
		}
	}
	return v139
}
func F_lua_remove(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
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
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var __phi63 int32
	_ = __phi63
	var v64 int32
	_ = v64
	var __phi64 int32
	_ = __phi64
	var v66 int64
	_ = v66
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v77 int32
	_ = v77
	if l1 < int32(1) {
		if l1 < int32(-9999) {
			switch l1 + int32(10002) {
			case 0:
				v55 = l0 + int32(72)
			case 1:
				v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
				v29 = *(*int32)(unsafe.Add(mBase, uint32(v28)+4))
				v30 = *(*int32)(unsafe.Add(mBase, uint32(v29)))
				v31 = *(*int32)(unsafe.Add(mBase, uint32(v30)+12))
				*(*int32)(unsafe.Add(mBase, uint32(l0)+96)) = int32(5)
				*(*int32)(unsafe.Add(mBase, uint32(l0)+88)) = v31
				v55 = l0 + int32(88)
			case 2:
				v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
				v55 = v25 + int32(96)
			default:
				v39 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
				v40 = *(*int32)(unsafe.Add(mBase, uint32(v39)+4))
				v41 = *(*int32)(unsafe.Add(mBase, uint32(v40)))
				v42 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v41)+7)))
				v43 = m.G398
				if base.Ui32(v42) < base.Ui32(int32(-10002)-l1) {
					v54 = v43
				} else {
					v54 = v41 + (int32(-10003)-l1)<<(uint(int32(4))%32) + int32(24)
				}
				v55 = v54
			}
		} else {
			v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			v55 = v19 + l1<<(uint(int32(4))%32)
		}
	} else {
		v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
		v13 = v8 + l1<<(uint(int32(4))%32) + int32(-16)
		v14 = m.G398
		if base.Ui32(v13) < base.Ui32(v7) {
			v16 = v13
		} else {
			v16 = v14
		}
		v55 = v16
	}
	v59 = v55 + int32(16)
	v60 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if base.Ui32(v60) <= base.Ui32(v59) {
		v77 = v60
	} else {
		__phi63 = v55
		__phi64 = v59
		v63 = __phi63
		v64 = __phi64
		for {
			v66 = *(*int64)(unsafe.Add(mBase, uint32(v63)+16))
			*(*int64)(unsafe.Add(mBase, uint32(v63))) = v66
			v68 = *(*int32)(unsafe.Add(mBase, uint32(v63)+24))
			*(*int32)(unsafe.Add(mBase, uint32(v63)+8)) = v68
			v71 = v64 + int32(16)
			v72 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			if base.Ui32(v71) < base.Ui32(v72) {
				__phi63 = v64
				__phi64 = v71
				v63 = __phi63
				v64 = __phi64
				continue
			} else {
				break
			}
			break
		}
		v77 = v72
	}
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v77 + int32(-16)
	return
}
func F_lua_resume(m *base.Module, l0 int32, l1 int32) int32 {
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
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v85 int32
	_ = v85
	var v88 int32
	_ = v88
	var v91 int32
	_ = v91
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v100 int32
	_ = v100
	var v103 int32
	_ = v103
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v111 int32
	_ = v111
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v121 int64
	_ = v121
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
	var v135 int32
	_ = v135
	var v138 int32
	_ = v138
	var v140 int32
	_ = v140
	v5 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+6)))
	switch v5 {
	case 0:
		v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
		v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
		if v7 == v8 {
			v47 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+52)))
			if base.Ui32(v47) < base.Ui32(int32(200)) {
				v85 = v47 + int32(1)
				*(*uint16)(unsafe.Add(mBase, uint32(l0)+54)) = uint16(v85)
				*(*uint16)(unsafe.Add(mBase, uint32(l0)+52)) = uint16(v85)
				v88 = m.G5
				v91 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				v95 = F_luaD_rawrunprotected(m, l0, v88+int32(1230), v91-l1<<(uint(int32(4))%32))
				mBase = m.M
				v96 = m.ExcPending
				if v96 != 0 {
					return int32(0)
				} else {
					if v95 == int32(0) {
						v134 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+6)))
						v135 = v134
						v138 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+52)))
						v140 = v138 + int32(-1)
						*(*uint16)(unsafe.Add(mBase, uint32(l0)+52)) = uint16(v140)
						return v135
					} else {
						*(*uint8)(unsafe.Add(mBase, uint32(l0)+6)) = uint8(v95)
						v100 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
						switch v95 + int32(-2) {
						case 0, 1:
							v121 = *(*int64)(unsafe.Add(mBase, uint32(v100+int32(-16))))
							*(*int64)(unsafe.Add(mBase, uint32(v100))) = v121
							v125 = *(*int32)(unsafe.Add(mBase, uint32(v100+int32(-8))))
							v126 = v125
							*(*int32)(unsafe.Add(mBase, uint32(v100)+8)) = v126
							v130 = v100 + int32(16)
							*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v130
							v132 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
							*(*int32)(unsafe.Add(mBase, uint32(v132)+8)) = v130
							v135 = v95
							v138 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+52)))
							v140 = v138 + int32(-1)
							*(*uint16)(unsafe.Add(mBase, uint32(l0)+52)) = uint16(v140)
							return v135
						case 2:
							v103 = m.G3
							v107 = F_luaS_newlstr(m, l0, v103+int32(_a2661), int32(17))
							mBase = m.M
							v108 = m.ExcPending
							if v108 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v100))) = v107
								v126 = int32(4)
								*(*int32)(unsafe.Add(mBase, uint32(v100)+8)) = v126
								v130 = v100 + int32(16)
								*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v130
								v132 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
								*(*int32)(unsafe.Add(mBase, uint32(v132)+8)) = v130
								v135 = v95
								v138 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+52)))
								v140 = v138 + int32(-1)
								*(*uint16)(unsafe.Add(mBase, uint32(l0)+52)) = uint16(v140)
								return v135
							}
						case 3:
							v111 = m.G3
							v115 = F_luaS_newlstr(m, l0, v111+int32(_a2662), int32(23))
							mBase = m.M
							v116 = m.ExcPending
							if v116 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v100))) = v115
								v126 = int32(4)
								*(*int32)(unsafe.Add(mBase, uint32(v100)+8)) = v126
								v130 = v100 + int32(16)
								*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v130
								v132 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
								*(*int32)(unsafe.Add(mBase, uint32(v132)+8)) = v130
								v135 = v95
								v138 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+52)))
								v140 = v138 + int32(-1)
								*(*uint16)(unsafe.Add(mBase, uint32(l0)+52)) = uint16(v140)
								return v135
							}
						default:
							v130 = v100 + int32(16)
							*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v130
							v132 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
							*(*int32)(unsafe.Add(mBase, uint32(v132)+8)) = v130
							v135 = v95
							v138 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+52)))
							v140 = v138 + int32(-1)
							*(*uint16)(unsafe.Add(mBase, uint32(l0)+52)) = uint16(v140)
							return v135
						}
					}
				}
			} else {
				v50 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
				v51 = *(*int32)(unsafe.Add(mBase, uint32(v50)))
				*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v51
				v53 = m.G3
				v57 = F_luaS_newlstr(m, l0, v53+int32(_a2660), int32(16))
				mBase = m.M
				v58 = m.ExcPending
				if v58 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v51)+8)) = int32(4)
					*(*int32)(unsafe.Add(mBase, uint32(v51))) = v57
					v62 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
					v63 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
					if int32(16) < v62-v63 {
						v78 = v63
						*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v78 + int32(16)
						return int32(2)
					} else {
						v67 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
						v68 = int32(1)
						if v67 < v68 {
							v74 = v67 + v68
						} else {
							v74 = v67 << (uint(v68) % 32)
						}
						F_luaD_reallocstack(m, l0, v74)
						mBase = m.M
						v76 = m.ExcPending
						if v76 != 0 {
							return int32(0)
						} else {
							v77 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
							v78 = v77
							*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v78 + int32(16)
							return int32(2)
						}
					}
				}
			}
		} else {
			v10 = v7
			v11 = *(*int32)(unsafe.Add(mBase, uint32(v10)))
			*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v11
			v13 = m.G3
			v17 = F_luaS_newlstr(m, l0, v13+int32(_a2663), int32(37))
			mBase = m.M
			v20 = m.ExcPending
			if v20 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v11)+8)) = int32(4)
				*(*int32)(unsafe.Add(mBase, uint32(v11))) = v17
				v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
				v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				if int32(16) < v24-v25 {
					v40 = v25
					*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v40 + int32(16)
					return int32(2)
				} else {
					v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
					v30 = int32(1)
					if v29 < v30 {
						v36 = v29 + v30
					} else {
						v36 = v29 << (uint(v30) % 32)
					}
					F_luaD_reallocstack(m, l0, v36)
					mBase = m.M
					v38 = m.ExcPending
					if v38 != 0 {
						return int32(0)
					} else {
						v39 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
						v40 = v39
						*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v40 + int32(16)
						return int32(2)
					}
				}
			}
		}
	case 1:
		v47 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+52)))
		if base.Ui32(v47) < base.Ui32(int32(200)) {
			v85 = v47 + int32(1)
			*(*uint16)(unsafe.Add(mBase, uint32(l0)+54)) = uint16(v85)
			*(*uint16)(unsafe.Add(mBase, uint32(l0)+52)) = uint16(v85)
			v88 = m.G5
			v91 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			v95 = F_luaD_rawrunprotected(m, l0, v88+int32(1230), v91-l1<<(uint(int32(4))%32))
			mBase = m.M
			v96 = m.ExcPending
			if v96 != 0 {
				return int32(0)
			} else {
				if v95 == int32(0) {
					v134 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+6)))
					v135 = v134
					v138 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+52)))
					v140 = v138 + int32(-1)
					*(*uint16)(unsafe.Add(mBase, uint32(l0)+52)) = uint16(v140)
					return v135
				} else {
					*(*uint8)(unsafe.Add(mBase, uint32(l0)+6)) = uint8(v95)
					v100 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
					switch v95 + int32(-2) {
					case 0, 1:
						v121 = *(*int64)(unsafe.Add(mBase, uint32(v100+int32(-16))))
						*(*int64)(unsafe.Add(mBase, uint32(v100))) = v121
						v125 = *(*int32)(unsafe.Add(mBase, uint32(v100+int32(-8))))
						v126 = v125
						*(*int32)(unsafe.Add(mBase, uint32(v100)+8)) = v126
						v130 = v100 + int32(16)
						*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v130
						v132 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
						*(*int32)(unsafe.Add(mBase, uint32(v132)+8)) = v130
						v135 = v95
						v138 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+52)))
						v140 = v138 + int32(-1)
						*(*uint16)(unsafe.Add(mBase, uint32(l0)+52)) = uint16(v140)
						return v135
					case 2:
						v103 = m.G3
						v107 = F_luaS_newlstr(m, l0, v103+int32(_a2661), int32(17))
						mBase = m.M
						v108 = m.ExcPending
						if v108 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v100))) = v107
							v126 = int32(4)
							*(*int32)(unsafe.Add(mBase, uint32(v100)+8)) = v126
							v130 = v100 + int32(16)
							*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v130
							v132 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
							*(*int32)(unsafe.Add(mBase, uint32(v132)+8)) = v130
							v135 = v95
							v138 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+52)))
							v140 = v138 + int32(-1)
							*(*uint16)(unsafe.Add(mBase, uint32(l0)+52)) = uint16(v140)
							return v135
						}
					case 3:
						v111 = m.G3
						v115 = F_luaS_newlstr(m, l0, v111+int32(_a2662), int32(23))
						mBase = m.M
						v116 = m.ExcPending
						if v116 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v100))) = v115
							v126 = int32(4)
							*(*int32)(unsafe.Add(mBase, uint32(v100)+8)) = v126
							v130 = v100 + int32(16)
							*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v130
							v132 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
							*(*int32)(unsafe.Add(mBase, uint32(v132)+8)) = v130
							v135 = v95
							v138 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+52)))
							v140 = v138 + int32(-1)
							*(*uint16)(unsafe.Add(mBase, uint32(l0)+52)) = uint16(v140)
							return v135
						}
					default:
						v130 = v100 + int32(16)
						*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v130
						v132 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
						*(*int32)(unsafe.Add(mBase, uint32(v132)+8)) = v130
						v135 = v95
						v138 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+52)))
						v140 = v138 + int32(-1)
						*(*uint16)(unsafe.Add(mBase, uint32(l0)+52)) = uint16(v140)
						return v135
					}
				}
			}
		} else {
			v50 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
			v51 = *(*int32)(unsafe.Add(mBase, uint32(v50)))
			*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v51
			v53 = m.G3
			v57 = F_luaS_newlstr(m, l0, v53+int32(_a2660), int32(16))
			mBase = m.M
			v58 = m.ExcPending
			if v58 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v51)+8)) = int32(4)
				*(*int32)(unsafe.Add(mBase, uint32(v51))) = v57
				v62 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
				v63 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				if int32(16) < v62-v63 {
					v78 = v63
					*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v78 + int32(16)
					return int32(2)
				} else {
					v67 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
					v68 = int32(1)
					if v67 < v68 {
						v74 = v67 + v68
					} else {
						v74 = v67 << (uint(v68) % 32)
					}
					F_luaD_reallocstack(m, l0, v74)
					mBase = m.M
					v76 = m.ExcPending
					if v76 != 0 {
						return int32(0)
					} else {
						v77 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
						v78 = v77
						*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v78 + int32(16)
						return int32(2)
					}
				}
			}
		}
	default:
		v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
		v10 = v6
		v11 = *(*int32)(unsafe.Add(mBase, uint32(v10)))
		*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v11
		v13 = m.G3
		v17 = F_luaS_newlstr(m, l0, v13+int32(_a2663), int32(37))
		mBase = m.M
		v20 = m.ExcPending
		if v20 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v11)+8)) = int32(4)
			*(*int32)(unsafe.Add(mBase, uint32(v11))) = v17
			v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
			v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			if int32(16) < v24-v25 {
				v40 = v25
				*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v40 + int32(16)
				return int32(2)
			} else {
				v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
				v30 = int32(1)
				if v29 < v30 {
					v36 = v29 + v30
				} else {
					v36 = v29 << (uint(v30) % 32)
				}
				F_luaD_reallocstack(m, l0, v36)
				mBase = m.M
				v38 = m.ExcPending
				if v38 != 0 {
					return int32(0)
				} else {
					v39 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
					v40 = v39
					*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v40 + int32(16)
					return int32(2)
				}
			}
		}
	}
}
func F_lua_setfenv(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
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
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v105 int32
	_ = v105
	var v108 int32
	_ = v108
	var v111 int32
	_ = v111
	var v114 int32
	_ = v114
	var v116 int32
	_ = v116
	if l1 < int32(1) {
		if l1 < int32(-9999) {
			switch l1 + int32(10002) {
			case 0:
				v55 = l0 + int32(72)
			case 1:
				v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
				v29 = *(*int32)(unsafe.Add(mBase, uint32(v28)+4))
				v30 = *(*int32)(unsafe.Add(mBase, uint32(v29)))
				v31 = *(*int32)(unsafe.Add(mBase, uint32(v30)+12))
				*(*int32)(unsafe.Add(mBase, uint32(l0)+96)) = int32(5)
				*(*int32)(unsafe.Add(mBase, uint32(l0)+88)) = v31
				v55 = l0 + int32(88)
			case 2:
				v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
				v55 = v25 + int32(96)
			default:
				v39 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
				v40 = *(*int32)(unsafe.Add(mBase, uint32(v39)+4))
				v41 = *(*int32)(unsafe.Add(mBase, uint32(v40)))
				v42 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v41)+7)))
				v43 = m.G398
				if base.Ui32(v42) < base.Ui32(int32(-10002)-l1) {
					v54 = v43
				} else {
					v54 = v41 + (int32(-10003)-l1)<<(uint(int32(4))%32) + int32(24)
				}
				v55 = v54
			}
		} else {
			v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			v55 = v19 + l1<<(uint(int32(4))%32)
		}
	} else {
		v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
		v13 = v8 + l1<<(uint(int32(4))%32) + int32(-16)
		v14 = m.G398
		if base.Ui32(v13) < base.Ui32(v7) {
			v16 = v13
		} else {
			v16 = v14
		}
		v55 = v16
	}
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v55)+8))
	switch v59 + int32(-6) {
	case 0:
		v62 = *(*int32)(unsafe.Add(mBase, uint32(v55)))
		v63 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		v66 = *(*int32)(unsafe.Add(mBase, uint32(v63+int32(-16))))
		*(*int32)(unsafe.Add(mBase, uint32(v62)+12)) = v66
		v84 = int32(1)
		v85 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		v88 = *(*int32)(unsafe.Add(mBase, uint32(v85+int32(-16))))
		v89 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v88)+5)))
		if v89&int32(3) == int32(0) {
			v114 = v84
		} else {
			v94 = *(*int32)(unsafe.Add(mBase, uint32(v55)))
			v95 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v94)+5)))
			if v95&int32(4) == int32(0) {
				v114 = v84
			} else {
				v100 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
				v101 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v100)+21)))
				if v101 != int32(1) {
					v105 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v100)+20)))
					v108 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v94)+5)))
					v111 = v105&int32(3) | v108&int32(248)
					*(*uint8)(unsafe.Add(mBase, uint32(v94)+5)) = uint8(v111)
				} else {
					F_reallymarkobject(m, v100, v88)
					mBase = m.M
				}
				v114 = v84
			}
		}
	case 1:
		v68 = *(*int32)(unsafe.Add(mBase, uint32(v55)))
		v69 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		v72 = *(*int32)(unsafe.Add(mBase, uint32(v69+int32(-16))))
		*(*int32)(unsafe.Add(mBase, uint32(v68)+12)) = v72
		v84 = int32(1)
		v85 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		v88 = *(*int32)(unsafe.Add(mBase, uint32(v85+int32(-16))))
		v89 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v88)+5)))
		if v89&int32(3) == int32(0) {
			v114 = v84
		} else {
			v94 = *(*int32)(unsafe.Add(mBase, uint32(v55)))
			v95 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v94)+5)))
			if v95&int32(4) == int32(0) {
				v114 = v84
			} else {
				v100 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
				v101 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v100)+21)))
				if v101 != int32(1) {
					v105 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v100)+20)))
					v108 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v94)+5)))
					v111 = v105&int32(3) | v108&int32(248)
					*(*uint8)(unsafe.Add(mBase, uint32(v94)+5)) = uint8(v111)
				} else {
					F_reallymarkobject(m, v100, v88)
					mBase = m.M
				}
				v114 = v84
			}
		}
	case 2:
		v74 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		v77 = *(*int32)(unsafe.Add(mBase, uint32(v74+int32(-16))))
		v78 = *(*int32)(unsafe.Add(mBase, uint32(v55)))
		*(*int32)(unsafe.Add(mBase, uint32(v78)+80)) = int32(5)
		*(*int32)(unsafe.Add(mBase, uint32(v78)+72)) = v77
		v84 = int32(1)
		v85 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		v88 = *(*int32)(unsafe.Add(mBase, uint32(v85+int32(-16))))
		v89 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v88)+5)))
		if v89&int32(3) == int32(0) {
			v114 = v84
		} else {
			v94 = *(*int32)(unsafe.Add(mBase, uint32(v55)))
			v95 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v94)+5)))
			if v95&int32(4) == int32(0) {
				v114 = v84
			} else {
				v100 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
				v101 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v100)+21)))
				if v101 != int32(1) {
					v105 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v100)+20)))
					v108 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v94)+5)))
					v111 = v105&int32(3) | v108&int32(248)
					*(*uint8)(unsafe.Add(mBase, uint32(v94)+5)) = uint8(v111)
				} else {
					F_reallymarkobject(m, v100, v88)
					mBase = m.M
				}
				v114 = v84
			}
		}
	default:
		v114 = int32(0)
	}
	v116 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v116 + int32(-16)
	return v114
}
func F_lua_setmetatable(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
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
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v87 int32
	_ = v87
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v106 int32
	_ = v106
	var v110 int32
	_ = v110
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v126 int32
	_ = v126
	var v129 int32
	_ = v129
	var v132 int32
	_ = v132
	var v134 int32
	_ = v134
	var v144 int32
	_ = v144
	if l1 < int32(1) {
		if l1 < int32(-9999) {
			switch l1 + int32(10002) {
			case 0:
				v56 = l0 + int32(72)
			case 1:
				v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
				v29 = *(*int32)(unsafe.Add(mBase, uint32(v28)+4))
				v30 = *(*int32)(unsafe.Add(mBase, uint32(v29)))
				v31 = *(*int32)(unsafe.Add(mBase, uint32(v30)+12))
				*(*int32)(unsafe.Add(mBase, uint32(l0)+96)) = int32(5)
				*(*int32)(unsafe.Add(mBase, uint32(l0)+88)) = v31
				v56 = l0 + int32(88)
			case 2:
				v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
				v56 = v25 + int32(96)
			default:
				v39 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
				v40 = *(*int32)(unsafe.Add(mBase, uint32(v39)+4))
				v41 = *(*int32)(unsafe.Add(mBase, uint32(v40)))
				v42 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v41)+7)))
				v43 = m.G398
				if base.Ui32(v42) < base.Ui32(int32(-10002)-l1) {
					v54 = v43
				} else {
					v54 = v41 + (int32(-10003)-l1)<<(uint(int32(4))%32) + int32(24)
				}
				v56 = v54
			}
		} else {
			v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			v56 = v19 + l1<<(uint(int32(4))%32)
		}
	} else {
		v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
		v13 = v8 + l1<<(uint(int32(4))%32) + int32(-16)
		v14 = m.G398
		if base.Ui32(v13) < base.Ui32(v7) {
			v16 = v13
		} else {
			v16 = v14
		}
		v56 = v16
	}
	v58 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v58+int32(-8))))
	if v61 != 0 {
		v65 = *(*int32)(unsafe.Add(mBase, uint32(v58+int32(-16))))
		v66 = v65
	} else {
		v66 = int32(0)
	}
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v56)+8))
	switch v67 + int32(-5) {
	case 0:
		v70 = *(*int32)(unsafe.Add(mBase, uint32(v56)))
		v71 = *(*int32)(unsafe.Add(mBase, uint32(v70)+8))
		if v71 == int32(0) {
			v83 = v70
			*(*int32)(unsafe.Add(mBase, uint32(v83)+16)) = v66
			if v66 == int32(0) {
			} else {
				v87 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v66)+5)))
				if v87&int32(3) == int32(0) {
				} else {
					v92 = *(*int32)(unsafe.Add(mBase, uint32(v56)))
					v93 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v92)+5)))
					if v93&int32(4) == int32(0) {
					} else {
						v98 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
						v99 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v92)+5)))
						v101 = v99 & int32(251)
						*(*uint8)(unsafe.Add(mBase, uint32(v92)+5)) = uint8(v101)
						v103 = *(*int32)(unsafe.Add(mBase, uint32(v98)+40))
						*(*int32)(unsafe.Add(mBase, uint32(v92)+32)) = v103
						*(*int32)(unsafe.Add(mBase, uint32(v98)+40)) = v92
					}
				}
			}
			v144 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v144 + int32(-16)
			return int32(1)
		} else {
			v74 = m.G3
			F_luaG_runerror(m, l0, v74+int32(_a2647), int32(0))
			mBase = m.M
			v81 = m.ExcPending
			if v81 != 0 {
				return int32(0)
			} else {
				v82 = *(*int32)(unsafe.Add(mBase, uint32(v56)))
				v83 = v82
				*(*int32)(unsafe.Add(mBase, uint32(v83)+16)) = v66
				if v66 == int32(0) {
				} else {
					v87 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v66)+5)))
					if v87&int32(3) == int32(0) {
					} else {
						v92 = *(*int32)(unsafe.Add(mBase, uint32(v56)))
						v93 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v92)+5)))
						if v93&int32(4) == int32(0) {
						} else {
							v98 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
							v99 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v92)+5)))
							v101 = v99 & int32(251)
							*(*uint8)(unsafe.Add(mBase, uint32(v92)+5)) = uint8(v101)
							v103 = *(*int32)(unsafe.Add(mBase, uint32(v98)+40))
							*(*int32)(unsafe.Add(mBase, uint32(v92)+32)) = v103
							*(*int32)(unsafe.Add(mBase, uint32(v98)+40)) = v92
						}
					}
				}
				v144 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v144 + int32(-16)
				return int32(1)
			}
		}
	default:
		v134 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
		*(*int32)(unsafe.Add(mBase, uint32(v134+v67<<(uint(int32(2))%32)+int32(152)))) = v66
		v144 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v144 + int32(-16)
		return int32(1)
	case 2:
		v106 = *(*int32)(unsafe.Add(mBase, uint32(v56)))
		*(*int32)(unsafe.Add(mBase, uint32(v106)+8)) = v66
		if v66 == int32(0) {
		} else {
			v110 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v66)+5)))
			if v110&int32(3) == int32(0) {
			} else {
				v115 = *(*int32)(unsafe.Add(mBase, uint32(v56)))
				v116 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v115)+5)))
				if v116&int32(4) == int32(0) {
				} else {
					v121 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
					v122 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v121)+21)))
					if v122 != int32(1) {
						v126 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v121)+20)))
						v129 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v115)+5)))
						v132 = v126&int32(3) | v129&int32(248)
						*(*uint8)(unsafe.Add(mBase, uint32(v115)+5)) = uint8(v132)
					} else {
						F_reallymarkobject(m, v121, v66)
						mBase = m.M
					}
				}
			}
		}
		v144 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v144 + int32(-16)
		return int32(1)
	}
}
func F_lua_tonumber(m *base.Module, l0 int32, l1 int32) float64 {
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
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v30 int32
	_ = v30
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
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v72 float64
	_ = v72
	var v74 float64
	_ = v74
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	if l1 < int32(1) {
		if l1 < int32(-9999) {
			switch l1 + int32(10002) {
			case 0:
				v61 = l0 + int32(72)
			case 1:
				v33 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
				v34 = *(*int32)(unsafe.Add(mBase, uint32(v33)+4))
				v35 = *(*int32)(unsafe.Add(mBase, uint32(v34)))
				v36 = *(*int32)(unsafe.Add(mBase, uint32(v35)+12))
				*(*int32)(unsafe.Add(mBase, uint32(l0)+96)) = int32(5)
				*(*int32)(unsafe.Add(mBase, uint32(l0)+88)) = v36
				v61 = l0 + int32(88)
			case 2:
				v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
				v61 = v30 + int32(96)
			default:
				v44 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
				v45 = *(*int32)(unsafe.Add(mBase, uint32(v44)+4))
				v46 = *(*int32)(unsafe.Add(mBase, uint32(v45)))
				v47 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v46)+7)))
				v48 = m.G398
				if base.Ui32(v47) < base.Ui32(int32(-10002)-l1) {
					v59 = v48
				} else {
					v59 = v46 + (int32(-10003)-l1)<<(uint(int32(4))%32) + int32(24)
				}
				v61 = v59
			}
		} else {
			v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			v61 = v24 + l1<<(uint(int32(4))%32)
		}
	} else {
		v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
		v18 = v13 + l1<<(uint(int32(4))%32) + int32(-16)
		v19 = m.G398
		if base.Ui32(v18) < base.Ui32(v12) {
			v21 = v18
		} else {
			v21 = v19
		}
		v61 = v21
	}
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v61)+8))
	if v63 == int32(3) {
		v71 = v61
		v72 = *(*float64)(unsafe.Add(mBase, uint32(v71)))
		v74 = v72
		m.G0 = v8 + int32(16)
		return v74
	} else {
		v66 = F_luaV_tonumber(m, v61, v8)
		mBase = m.M
		v69 = m.ExcPending
		if v69 != 0 {
			return float64(0)
		} else {
			if v66 != 0 {
				v71 = v66
				v72 = *(*float64)(unsafe.Add(mBase, uint32(v71)))
				v74 = v72
			} else {
				v74 = float64(0)
			}
			m.G0 = v8 + int32(16)
			return v74
		}
	}
}
func F_lua_xmove(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v18 int32
	_ = v18
	var v25 int32
	_ = v25
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
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int64
	_ = v44
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v53 int32
	_ = v53
	var v56 int64
	_ = v56
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v71 int32
	_ = v71
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v85 int32
	_ = v85
	var v86 int64
	_ = v86
	var v88 int32
	_ = v88
	if l0 == l1 {
	} else {
		v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v11 - l2<<(uint(int32(4))%32)
		if l2 < int32(1) {
		} else {
			v18 = int32(1)
			if l2 == v18 {
				v71 = int32(0)
			} else {
				v25 = int32(0)
				v31 = v25
				v33 = v25
				for {
					v36 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
					v37 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
					v38 = int32(16)
					*(*int32)(unsafe.Add(mBase, uint32(l1)+8)) = v37 + v38
					v42 = v31 << (uint(int32(4)) % 32)
					v43 = v36 + v42
					v44 = *(*int64)(unsafe.Add(mBase, uint32(v43)))
					*(*int64)(unsafe.Add(mBase, uint32(v37))) = v44
					v46 = *(*int32)(unsafe.Add(mBase, uint32(v43)+8))
					*(*int32)(unsafe.Add(mBase, uint32(v37)+8)) = v46
					v48 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
					v49 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
					*(*int32)(unsafe.Add(mBase, uint32(l1)+8)) = v49 + v38
					v53 = v48 + v42
					v56 = *(*int64)(unsafe.Add(mBase, uint32(v53+v38)))
					*(*int64)(unsafe.Add(mBase, uint32(v49))) = v56
					v60 = *(*int32)(unsafe.Add(mBase, uint32(v53+int32(24))))
					*(*int32)(unsafe.Add(mBase, uint32(v49)+8)) = v60
					v62 = int32(2)
					v63 = v31 + v62
					v65 = v33 + v62
					if v65 != l2&int32(2147483646) {
						v31 = v63
						v33 = v65
						continue
					} else {
						break
					}
					break
				}
				v71 = v63
			}
			if l2&v18 == int32(0) {
			} else {
				v78 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				v79 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
				*(*int32)(unsafe.Add(mBase, uint32(l1)+8)) = v79 + int32(16)
				v85 = v78 + v71<<(uint(int32(4))%32)
				v86 = *(*int64)(unsafe.Add(mBase, uint32(v85)))
				*(*int64)(unsafe.Add(mBase, uint32(v79))) = v86
				v88 = *(*int32)(unsafe.Add(mBase, uint32(v85)+8))
				*(*int32)(unsafe.Add(mBase, uint32(v79)+8)) = v88
			}
		}
	}
	return
}
