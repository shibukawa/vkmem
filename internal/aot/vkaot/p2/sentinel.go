package p2

import (
	base "github.com/shibukawa/vkmem/internal/aot/vkaot/base"
	"unsafe"
)

func F_createSentinelAddr(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v16 int32
	_ = v16
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v31 int32
	_ = v31
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
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v65 int32
	_ = v65
	v6 = m.G0
	v8 = v6 - int32(64)
	m.G0 = v8
	if base.Ui32(l1) < base.Ui32(int32(65536)) {
		v16 = int32(0)
		v21 = *(*int32)(unsafe.Add(mBase, _consts[729]))
		v24 = F_anetResolve(m, v16, l0, v8+int32(16), int32(46), base.B2i32(v21 == v16))
		mBase = m.M
		v27 = m.ExcPending
		if v27 != 0 {
			return int32(0)
		} else {
			if v24 != int32(-1) {
				v50 = F_valkey_malloc(m, int32(12))
				mBase = m.M
				v51 = m.ExcPending
				if v51 != 0 {
					return int32(0)
				} else {
					v52 = F_sdsnew(m, l0)
					mBase = m.M
					v53 = m.ExcPending
					if v53 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v50))) = v52
						v57 = F_sdsnew(m, v8+int32(16))
						mBase = m.M
						v58 = m.ExcPending
						if v58 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v50)+8)) = l1
							*(*int32)(unsafe.Add(mBase, uint32(v50)+4)) = v57
							v65 = v50
							m.G0 = v8 + int32(64)
							return v65
						}
					}
				}
			} else {
				v31 = *(*int32)(unsafe.Add(mBase, _consts[6]))
				if int32(3) < v31 {
					v39 = int32(0)
					if l2 == v39 {
						*(*int32)(unsafe.Add(mBase, _consts[5])) = int32(44)
						v65 = v39
						m.G0 = v8 + int32(64)
						return v65
					} else {
						v42 = int32(0)
						v43 = *(*int32)(unsafe.Add(mBase, _consts[729]))
						if v43 == v42 {
							*(*int32)(unsafe.Add(mBase, _consts[5])) = int32(44)
							v65 = v39
							m.G0 = v8 + int32(64)
							return v65
						} else {
							v46 = int32(0)
							*(*uint8)(unsafe.Add(mBase, uint32(v8)+16)) = uint8(v46)
							v50 = F_valkey_malloc(m, int32(12))
							mBase = m.M
							v51 = m.ExcPending
							if v51 != 0 {
								return int32(0)
							} else {
								v52 = F_sdsnew(m, l0)
								mBase = m.M
								v53 = m.ExcPending
								if v53 != 0 {
									return int32(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v50))) = v52
									v57 = F_sdsnew(m, v8+int32(16))
									mBase = m.M
									v58 = m.ExcPending
									if v58 != 0 {
										return int32(0)
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v50)+8)) = l1
										*(*int32)(unsafe.Add(mBase, uint32(v50)+4)) = v57
										v65 = v50
										m.G0 = v8 + int32(64)
										return v65
									}
								}
							}
						}
					}
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v8))) = l0
					F__serverLog(m, int32(3), int32(_a1340), v8)
					mBase = m.M
					v38 = m.ExcPending
					if v38 != 0 {
						return int32(0)
					} else {
						v39 = int32(0)
						if l2 == v39 {
							*(*int32)(unsafe.Add(mBase, _consts[5])) = int32(44)
							v65 = v39
							m.G0 = v8 + int32(64)
							return v65
						} else {
							v42 = int32(0)
							v43 = *(*int32)(unsafe.Add(mBase, _consts[729]))
							if v43 == v42 {
								*(*int32)(unsafe.Add(mBase, _consts[5])) = int32(44)
								v65 = v39
								m.G0 = v8 + int32(64)
								return v65
							} else {
								v46 = int32(0)
								*(*uint8)(unsafe.Add(mBase, uint32(v8)+16)) = uint8(v46)
								v50 = F_valkey_malloc(m, int32(12))
								mBase = m.M
								v51 = m.ExcPending
								if v51 != 0 {
									return int32(0)
								} else {
									v52 = F_sdsnew(m, l0)
									mBase = m.M
									v53 = m.ExcPending
									if v53 != 0 {
										return int32(0)
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v50))) = v52
										v57 = F_sdsnew(m, v8+int32(16))
										mBase = m.M
										v58 = m.ExcPending
										if v58 != 0 {
											return int32(0)
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(v50)+8)) = l1
											*(*int32)(unsafe.Add(mBase, uint32(v50)+4)) = v57
											v65 = v50
											m.G0 = v8 + int32(64)
											return v65
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
		*(*int32)(unsafe.Add(mBase, _consts[5])) = int32(28)
		v65 = int32(0)
		m.G0 = v8 + int32(64)
		return v65
	}
}
func F_createSentinelValkeyInstance(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32 {
	mBase := m.M
	_ = mBase
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v49 int32
	_ = v49
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
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v110 int64
	_ = v110
	var v128 int64
	_ = v128
	var v132 int64
	_ = v132
	var v134 int64
	_ = v134
	var v137 int64
	_ = v137
	var v139 int64
	_ = v139
	var v141 int64
	_ = v141
	var v152 int32
	_ = v152
	var v153 int64
	_ = v153
	var v154 int64
	_ = v154
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v190 int64
	_ = v190
	var v192 int32
	_ = v192
	var v210 int64
	_ = v210
	var v222 int32
	_ = v222
	var v226 int64
	_ = v226
	var v228 int64
	_ = v228
	var v232 int64
	_ = v232
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
	var v241 int32
	_ = v241
	var v252 int32
	_ = v252
	var v258 int32
	_ = v258
	v12 = m.G0
	v14 = v12 - int32(16)
	m.G0 = v14
	if l1&int32(7) == int32(0) {
		F__serverAssert(m, int32(_a1356), int32(_a1333), int32(1281))
		mBase = m.M
		v252 = m.ExcPending
		if v252 != 0 {
			return int32(0)
		} else {
			F_abort(m)
			mBase = m.M
			base.Wasm_trap_unreachable()
			for {
			}
		}
	} else {
		v21 = l1 & int32(1)
		if v21 != 0 {
			v24 = int32(0)
			v26 = F_createSentinelAddr(m, l2, l3, int32(1))
			mBase = m.M
			v29 = m.ExcPending
			if v29 != 0 {
				return int32(0)
			} else {
				if v26 == int32(0) {
					v241 = v24
					m.G0 = v14 + int32(16)
					return v241
				} else {
					if l1&int32(2) != 0 {
						v42 = int32(0)
						v43 = *(*int32)(unsafe.Add(mBase, _consts[734]))
						v49 = *(*int32)(unsafe.Add(mBase, uint32(v26+base.B2i32(v43 == v42)<<(uint(int32(2))%32))))
						v50 = int32(58)
						v51 = F___strchrnul(m, v49, v50)
						mBase = m.M
						v53 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v51))))
						if v53 == v50 {
							v57 = v51
						} else {
							v57 = v42
						}
						v58 = F_sdsempty(m)
						mBase = m.M
						v59 = m.ExcPending
						if v59 != 0 {
							return int32(0)
						} else {
							v60 = *(*int32)(unsafe.Add(mBase, uint32(v26)+8))
							*(*int32)(unsafe.Add(mBase, uint32(v14)+4)) = v60
							*(*int32)(unsafe.Add(mBase, uint32(v14))) = v49
							if v57 != 0 {
								v65 = int32(_a1004)
							} else {
								v65 = int32(_a1005)
							}
							v66 = F_sdscatprintf(m, v58, v65, v14)
							mBase = m.M
							v67 = m.ExcPending
							if v67 != 0 {
								return int32(0)
							} else {
								if v21 == int32(0) {
									v75 = *(*int32)(unsafe.Add(mBase, uint32(l5)+148))
									v77 = v66
									v78 = v75
								} else {
									v71 = v66
									v74 = *(*int32)(unsafe.Add(mBase, _consts[733]))
									v77 = v71
									v78 = v74
								}
								v80 = F_dictFind(m, v78, v77)
								mBase = m.M
								v81 = m.ExcPending
								if v81 != 0 {
									return int32(0)
								} else {
									if v80 == int32(0) {
										v98 = F_valkey_malloc(m, int32(296))
										mBase = m.M
										v99 = m.ExcPending
										if v99 != 0 {
											return int32(0)
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(v98)+24)) = v26
											*(*int64)(unsafe.Add(mBase, uint32(v98)+16)) = int64(0)
											*(*int32)(unsafe.Add(mBase, uint32(v98)+8)) = int32(0)
											*(*int32)(unsafe.Add(mBase, uint32(v98)+4)) = v77
											*(*int32)(unsafe.Add(mBase, uint32(v98))) = l1
											v108 = F_valkey_malloc(m, int32(88))
											mBase = m.M
											v109 = m.ExcPending
											if v109 != 0 {
												return int32(0)
											} else {
												v110 = int64(0)
												*(*int64)(unsafe.Add(mBase, uint32(v108)+80)) = v110
												*(*int32)(unsafe.Add(mBase, uint32(v108)+16)) = int32(0)
												*(*int64)(unsafe.Add(mBase, uint32(v108)+8)) = v110
												*(*int64)(unsafe.Add(mBase, uint32(v108))) = int64(4294967297)
												*(*int64)(unsafe.Add(mBase, uint32(v108)+24)) = v110
												*(*int64)(unsafe.Add(mBase, uint32(v108+int32(32)))) = v110
												*(*int64)(unsafe.Add(mBase, uint32(v108+int32(40)))) = v110
												v128 = F_mstime(m)
												mBase = m.M
												*(*int64)(unsafe.Add(mBase, uint32(v108)+64)) = v110
												*(*int64)(unsafe.Add(mBase, uint32(v108)+56)) = v128
												v132 = F_mstime(m)
												mBase = m.M
												*(*int64)(unsafe.Add(mBase, uint32(v108)+48)) = v132
												v134 = F_mstime(m)
												mBase = m.M
												*(*int64)(unsafe.Add(mBase, uint32(v108)+72)) = v134
												*(*int32)(unsafe.Add(mBase, uint32(v98)+28)) = v108
												v137 = F_mstime(m)
												mBase = m.M
												*(*int64)(unsafe.Add(mBase, uint32(v98)+32)) = v137
												v139 = F_mstime(m)
												mBase = m.M
												*(*int64)(unsafe.Add(mBase, uint32(v98)+40)) = v139
												v141 = F_mstime(m)
												mBase = m.M
												*(*int64)(unsafe.Add(mBase, uint32(v98)+56)) = v110
												*(*int64)(unsafe.Add(mBase, uint32(v98)+48)) = v141
												*(*int64)(unsafe.Add(mBase, uint32(v98+int32(64)))) = v110
												if l5 != 0 {
													v152 = l5 + int32(72)
												} else {
													v152 = int32(_a1357)
												}
												v153 = *(*int64)(unsafe.Add(mBase, uint32(v152)))
												v154 = int64(0)
												*(*int64)(unsafe.Add(mBase, uint32(v98)+160)) = v154
												*(*int64)(unsafe.Add(mBase, uint32(v98)+80)) = v154
												*(*int64)(unsafe.Add(mBase, uint32(v98)+72)) = v153
												*(*int64)(unsafe.Add(mBase, uint32(v98+int32(168)))) = v154
												*(*int64)(unsafe.Add(mBase, uint32(v98)+208)) = v154
												*(*int32)(unsafe.Add(mBase, uint32(v98)+204)) = int32(1)
												*(*int64)(unsafe.Add(mBase, uint32(v98)+196)) = v154
												*(*int64)(unsafe.Add(mBase, uint32(v98)+184)) = v154
												*(*int64)(unsafe.Add(mBase, uint32(v98)+176)) = int64(4294967396)
												v174 = F_dictCreate(m, int32(_a1358))
												mBase = m.M
												v175 = m.ExcPending
												if v175 != 0 {
													return int32(0)
												} else {
													*(*int32)(unsafe.Add(mBase, uint32(v98)+192)) = l5
													*(*int32)(unsafe.Add(mBase, uint32(v98)+156)) = int32(1)
													*(*int32)(unsafe.Add(mBase, uint32(v98)+152)) = l4
													*(*int32)(unsafe.Add(mBase, uint32(v98)+144)) = v174
													v182 = F_dictCreate(m, int32(_a1358))
													mBase = m.M
													v183 = m.ExcPending
													if v183 != 0 {
														return int32(0)
													} else {
														*(*int64)(unsafe.Add(mBase, uint32(v98)+96)) = int64(0)
														*(*int32)(unsafe.Add(mBase, uint32(v98)+148)) = v182
														v188 = F_dictCreate(m, int32(_a1359))
														mBase = m.M
														v189 = m.ExcPending
														if v189 != 0 {
															return int32(0)
														} else {
															v190 = int64(0)
															*(*int64)(unsafe.Add(mBase, uint32(v98)+224)) = v190
															v192 = int32(0)
															*(*int32)(unsafe.Add(mBase, uint32(v98)+216)) = v192
															*(*int32)(unsafe.Add(mBase, uint32(v98)+104)) = v188
															*(*int64)(unsafe.Add(mBase, uint32(v98+int32(232)))) = v190
															*(*int32)(unsafe.Add(mBase, uint32(v98+int32(240)))) = v192
															*(*int64)(unsafe.Add(mBase, uint32(v98)+248)) = v190
															*(*int64)(unsafe.Add(mBase, uint32(v98+int32(256)))) = v190
															v210 = *(*int64)(unsafe.Add(mBase, _consts[735]))
															*(*int64)(unsafe.Add(mBase, uint32(v98)+272)) = v190
															*(*int64)(unsafe.Add(mBase, uint32(v98)+264)) = v210
															*(*int64)(unsafe.Add(mBase, uint32(v98+int32(280)))) = v190
															*(*int64)(unsafe.Add(mBase, uint32(v98+int32(288)))) = v190
															v222 = *(*int32)(unsafe.Add(mBase, uint32(v98)))
															*(*int32)(unsafe.Add(mBase, uint32(v98)+108)) = v222 & int32(3)
															v226 = F_mstime(m)
															mBase = m.M
															*(*int64)(unsafe.Add(mBase, uint32(v98)+112)) = v226
															v228 = F_mstime(m)
															mBase = m.M
															*(*int32)(unsafe.Add(mBase, uint32(v98)+128)) = v192
															*(*int64)(unsafe.Add(mBase, uint32(v98)+120)) = v228
															v232 = F_mstime(m)
															mBase = m.M
															*(*int64)(unsafe.Add(mBase, uint32(v98)+136)) = v232
															v234 = *(*int32)(unsafe.Add(mBase, uint32(v98)+4))
															v235 = F_dictAdd(m, v78, v234, v98)
															mBase = m.M
															v236 = m.ExcPending
															if v236 != 0 {
																return int32(0)
															} else {
																v241 = v98
																m.G0 = v14 + int32(16)
																return v241
															}
														}
													}
												}
											}
										}
									} else {
										v84 = *(*int32)(unsafe.Add(mBase, uint32(v26)))
										F_sdsfree(m, v84)
										mBase = m.M
										v86 = m.ExcPending
										if v86 != 0 {
											return int32(0)
										} else {
											v87 = *(*int32)(unsafe.Add(mBase, uint32(v26)+4))
											F_sdsfree(m, v87)
											mBase = m.M
											v89 = m.ExcPending
											if v89 != 0 {
												return int32(0)
											} else {
												F_valkey_free(m, v26)
												mBase = m.M
												v91 = m.ExcPending
												if v91 != 0 {
													return int32(0)
												} else {
													F_sdsfree(m, v77)
													mBase = m.M
													v93 = m.ExcPending
													if v93 != 0 {
														return int32(0)
													} else {
														*(*int32)(unsafe.Add(mBase, _consts[5])) = int32(10)
														v241 = v24
														m.G0 = v14 + int32(16)
														return v241
													}
												}
											}
										}
									}
								}
							}
						}
					} else {
						v34 = F_sdsnew(m, l0)
						mBase = m.M
						v35 = m.ExcPending
						if v35 != 0 {
							return int32(0)
						} else {
							if v21 != 0 {
								v71 = v34
								v74 = *(*int32)(unsafe.Add(mBase, _consts[733]))
								v77 = v71
								v78 = v74
							} else {
								v36 = int32(0)
								if l1&int32(4) == v36 {
									v77 = v34
									v78 = v36
								} else {
									v41 = *(*int32)(unsafe.Add(mBase, uint32(l5)+144))
									v77 = v34
									v78 = v41
								}
							}
							v80 = F_dictFind(m, v78, v77)
							mBase = m.M
							v81 = m.ExcPending
							if v81 != 0 {
								return int32(0)
							} else {
								if v80 == int32(0) {
									v98 = F_valkey_malloc(m, int32(296))
									mBase = m.M
									v99 = m.ExcPending
									if v99 != 0 {
										return int32(0)
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v98)+24)) = v26
										*(*int64)(unsafe.Add(mBase, uint32(v98)+16)) = int64(0)
										*(*int32)(unsafe.Add(mBase, uint32(v98)+8)) = int32(0)
										*(*int32)(unsafe.Add(mBase, uint32(v98)+4)) = v77
										*(*int32)(unsafe.Add(mBase, uint32(v98))) = l1
										v108 = F_valkey_malloc(m, int32(88))
										mBase = m.M
										v109 = m.ExcPending
										if v109 != 0 {
											return int32(0)
										} else {
											v110 = int64(0)
											*(*int64)(unsafe.Add(mBase, uint32(v108)+80)) = v110
											*(*int32)(unsafe.Add(mBase, uint32(v108)+16)) = int32(0)
											*(*int64)(unsafe.Add(mBase, uint32(v108)+8)) = v110
											*(*int64)(unsafe.Add(mBase, uint32(v108))) = int64(4294967297)
											*(*int64)(unsafe.Add(mBase, uint32(v108)+24)) = v110
											*(*int64)(unsafe.Add(mBase, uint32(v108+int32(32)))) = v110
											*(*int64)(unsafe.Add(mBase, uint32(v108+int32(40)))) = v110
											v128 = F_mstime(m)
											mBase = m.M
											*(*int64)(unsafe.Add(mBase, uint32(v108)+64)) = v110
											*(*int64)(unsafe.Add(mBase, uint32(v108)+56)) = v128
											v132 = F_mstime(m)
											mBase = m.M
											*(*int64)(unsafe.Add(mBase, uint32(v108)+48)) = v132
											v134 = F_mstime(m)
											mBase = m.M
											*(*int64)(unsafe.Add(mBase, uint32(v108)+72)) = v134
											*(*int32)(unsafe.Add(mBase, uint32(v98)+28)) = v108
											v137 = F_mstime(m)
											mBase = m.M
											*(*int64)(unsafe.Add(mBase, uint32(v98)+32)) = v137
											v139 = F_mstime(m)
											mBase = m.M
											*(*int64)(unsafe.Add(mBase, uint32(v98)+40)) = v139
											v141 = F_mstime(m)
											mBase = m.M
											*(*int64)(unsafe.Add(mBase, uint32(v98)+56)) = v110
											*(*int64)(unsafe.Add(mBase, uint32(v98)+48)) = v141
											*(*int64)(unsafe.Add(mBase, uint32(v98+int32(64)))) = v110
											if l5 != 0 {
												v152 = l5 + int32(72)
											} else {
												v152 = int32(_a1357)
											}
											v153 = *(*int64)(unsafe.Add(mBase, uint32(v152)))
											v154 = int64(0)
											*(*int64)(unsafe.Add(mBase, uint32(v98)+160)) = v154
											*(*int64)(unsafe.Add(mBase, uint32(v98)+80)) = v154
											*(*int64)(unsafe.Add(mBase, uint32(v98)+72)) = v153
											*(*int64)(unsafe.Add(mBase, uint32(v98+int32(168)))) = v154
											*(*int64)(unsafe.Add(mBase, uint32(v98)+208)) = v154
											*(*int32)(unsafe.Add(mBase, uint32(v98)+204)) = int32(1)
											*(*int64)(unsafe.Add(mBase, uint32(v98)+196)) = v154
											*(*int64)(unsafe.Add(mBase, uint32(v98)+184)) = v154
											*(*int64)(unsafe.Add(mBase, uint32(v98)+176)) = int64(4294967396)
											v174 = F_dictCreate(m, int32(_a1358))
											mBase = m.M
											v175 = m.ExcPending
											if v175 != 0 {
												return int32(0)
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(v98)+192)) = l5
												*(*int32)(unsafe.Add(mBase, uint32(v98)+156)) = int32(1)
												*(*int32)(unsafe.Add(mBase, uint32(v98)+152)) = l4
												*(*int32)(unsafe.Add(mBase, uint32(v98)+144)) = v174
												v182 = F_dictCreate(m, int32(_a1358))
												mBase = m.M
												v183 = m.ExcPending
												if v183 != 0 {
													return int32(0)
												} else {
													*(*int64)(unsafe.Add(mBase, uint32(v98)+96)) = int64(0)
													*(*int32)(unsafe.Add(mBase, uint32(v98)+148)) = v182
													v188 = F_dictCreate(m, int32(_a1359))
													mBase = m.M
													v189 = m.ExcPending
													if v189 != 0 {
														return int32(0)
													} else {
														v190 = int64(0)
														*(*int64)(unsafe.Add(mBase, uint32(v98)+224)) = v190
														v192 = int32(0)
														*(*int32)(unsafe.Add(mBase, uint32(v98)+216)) = v192
														*(*int32)(unsafe.Add(mBase, uint32(v98)+104)) = v188
														*(*int64)(unsafe.Add(mBase, uint32(v98+int32(232)))) = v190
														*(*int32)(unsafe.Add(mBase, uint32(v98+int32(240)))) = v192
														*(*int64)(unsafe.Add(mBase, uint32(v98)+248)) = v190
														*(*int64)(unsafe.Add(mBase, uint32(v98+int32(256)))) = v190
														v210 = *(*int64)(unsafe.Add(mBase, _consts[735]))
														*(*int64)(unsafe.Add(mBase, uint32(v98)+272)) = v190
														*(*int64)(unsafe.Add(mBase, uint32(v98)+264)) = v210
														*(*int64)(unsafe.Add(mBase, uint32(v98+int32(280)))) = v190
														*(*int64)(unsafe.Add(mBase, uint32(v98+int32(288)))) = v190
														v222 = *(*int32)(unsafe.Add(mBase, uint32(v98)))
														*(*int32)(unsafe.Add(mBase, uint32(v98)+108)) = v222 & int32(3)
														v226 = F_mstime(m)
														mBase = m.M
														*(*int64)(unsafe.Add(mBase, uint32(v98)+112)) = v226
														v228 = F_mstime(m)
														mBase = m.M
														*(*int32)(unsafe.Add(mBase, uint32(v98)+128)) = v192
														*(*int64)(unsafe.Add(mBase, uint32(v98)+120)) = v228
														v232 = F_mstime(m)
														mBase = m.M
														*(*int64)(unsafe.Add(mBase, uint32(v98)+136)) = v232
														v234 = *(*int32)(unsafe.Add(mBase, uint32(v98)+4))
														v235 = F_dictAdd(m, v78, v234, v98)
														mBase = m.M
														v236 = m.ExcPending
														if v236 != 0 {
															return int32(0)
														} else {
															v241 = v98
															m.G0 = v14 + int32(16)
															return v241
														}
													}
												}
											}
										}
									}
								} else {
									v84 = *(*int32)(unsafe.Add(mBase, uint32(v26)))
									F_sdsfree(m, v84)
									mBase = m.M
									v86 = m.ExcPending
									if v86 != 0 {
										return int32(0)
									} else {
										v87 = *(*int32)(unsafe.Add(mBase, uint32(v26)+4))
										F_sdsfree(m, v87)
										mBase = m.M
										v89 = m.ExcPending
										if v89 != 0 {
											return int32(0)
										} else {
											F_valkey_free(m, v26)
											mBase = m.M
											v91 = m.ExcPending
											if v91 != 0 {
												return int32(0)
											} else {
												F_sdsfree(m, v77)
												mBase = m.M
												v93 = m.ExcPending
												if v93 != 0 {
													return int32(0)
												} else {
													*(*int32)(unsafe.Add(mBase, _consts[5])) = int32(10)
													v241 = v24
													m.G0 = v14 + int32(16)
													return v241
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
			if l5 == int32(0) {
				F__serverAssert(m, int32(_a1360), int32(_a1333), int32(1282))
				mBase = m.M
				v258 = m.ExcPending
				if v258 != 0 {
					return int32(0)
				} else {
					F_abort(m)
					mBase = m.M
					base.Wasm_trap_unreachable()
					for {
					}
				}
			} else {
				v24 = int32(0)
				v26 = F_createSentinelAddr(m, l2, l3, int32(1))
				mBase = m.M
				v29 = m.ExcPending
				if v29 != 0 {
					return int32(0)
				} else {
					if v26 == int32(0) {
						v241 = v24
						m.G0 = v14 + int32(16)
						return v241
					} else {
						if l1&int32(2) != 0 {
							v42 = int32(0)
							v43 = *(*int32)(unsafe.Add(mBase, _consts[734]))
							v49 = *(*int32)(unsafe.Add(mBase, uint32(v26+base.B2i32(v43 == v42)<<(uint(int32(2))%32))))
							v50 = int32(58)
							v51 = F___strchrnul(m, v49, v50)
							mBase = m.M
							v53 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v51))))
							if v53 == v50 {
								v57 = v51
							} else {
								v57 = v42
							}
							v58 = F_sdsempty(m)
							mBase = m.M
							v59 = m.ExcPending
							if v59 != 0 {
								return int32(0)
							} else {
								v60 = *(*int32)(unsafe.Add(mBase, uint32(v26)+8))
								*(*int32)(unsafe.Add(mBase, uint32(v14)+4)) = v60
								*(*int32)(unsafe.Add(mBase, uint32(v14))) = v49
								if v57 != 0 {
									v65 = int32(_a1004)
								} else {
									v65 = int32(_a1005)
								}
								v66 = F_sdscatprintf(m, v58, v65, v14)
								mBase = m.M
								v67 = m.ExcPending
								if v67 != 0 {
									return int32(0)
								} else {
									if v21 == int32(0) {
										v75 = *(*int32)(unsafe.Add(mBase, uint32(l5)+148))
										v77 = v66
										v78 = v75
									} else {
										v71 = v66
										v74 = *(*int32)(unsafe.Add(mBase, _consts[733]))
										v77 = v71
										v78 = v74
									}
									v80 = F_dictFind(m, v78, v77)
									mBase = m.M
									v81 = m.ExcPending
									if v81 != 0 {
										return int32(0)
									} else {
										if v80 == int32(0) {
											v98 = F_valkey_malloc(m, int32(296))
											mBase = m.M
											v99 = m.ExcPending
											if v99 != 0 {
												return int32(0)
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(v98)+24)) = v26
												*(*int64)(unsafe.Add(mBase, uint32(v98)+16)) = int64(0)
												*(*int32)(unsafe.Add(mBase, uint32(v98)+8)) = int32(0)
												*(*int32)(unsafe.Add(mBase, uint32(v98)+4)) = v77
												*(*int32)(unsafe.Add(mBase, uint32(v98))) = l1
												v108 = F_valkey_malloc(m, int32(88))
												mBase = m.M
												v109 = m.ExcPending
												if v109 != 0 {
													return int32(0)
												} else {
													v110 = int64(0)
													*(*int64)(unsafe.Add(mBase, uint32(v108)+80)) = v110
													*(*int32)(unsafe.Add(mBase, uint32(v108)+16)) = int32(0)
													*(*int64)(unsafe.Add(mBase, uint32(v108)+8)) = v110
													*(*int64)(unsafe.Add(mBase, uint32(v108))) = int64(4294967297)
													*(*int64)(unsafe.Add(mBase, uint32(v108)+24)) = v110
													*(*int64)(unsafe.Add(mBase, uint32(v108+int32(32)))) = v110
													*(*int64)(unsafe.Add(mBase, uint32(v108+int32(40)))) = v110
													v128 = F_mstime(m)
													mBase = m.M
													*(*int64)(unsafe.Add(mBase, uint32(v108)+64)) = v110
													*(*int64)(unsafe.Add(mBase, uint32(v108)+56)) = v128
													v132 = F_mstime(m)
													mBase = m.M
													*(*int64)(unsafe.Add(mBase, uint32(v108)+48)) = v132
													v134 = F_mstime(m)
													mBase = m.M
													*(*int64)(unsafe.Add(mBase, uint32(v108)+72)) = v134
													*(*int32)(unsafe.Add(mBase, uint32(v98)+28)) = v108
													v137 = F_mstime(m)
													mBase = m.M
													*(*int64)(unsafe.Add(mBase, uint32(v98)+32)) = v137
													v139 = F_mstime(m)
													mBase = m.M
													*(*int64)(unsafe.Add(mBase, uint32(v98)+40)) = v139
													v141 = F_mstime(m)
													mBase = m.M
													*(*int64)(unsafe.Add(mBase, uint32(v98)+56)) = v110
													*(*int64)(unsafe.Add(mBase, uint32(v98)+48)) = v141
													*(*int64)(unsafe.Add(mBase, uint32(v98+int32(64)))) = v110
													if l5 != 0 {
														v152 = l5 + int32(72)
													} else {
														v152 = int32(_a1357)
													}
													v153 = *(*int64)(unsafe.Add(mBase, uint32(v152)))
													v154 = int64(0)
													*(*int64)(unsafe.Add(mBase, uint32(v98)+160)) = v154
													*(*int64)(unsafe.Add(mBase, uint32(v98)+80)) = v154
													*(*int64)(unsafe.Add(mBase, uint32(v98)+72)) = v153
													*(*int64)(unsafe.Add(mBase, uint32(v98+int32(168)))) = v154
													*(*int64)(unsafe.Add(mBase, uint32(v98)+208)) = v154
													*(*int32)(unsafe.Add(mBase, uint32(v98)+204)) = int32(1)
													*(*int64)(unsafe.Add(mBase, uint32(v98)+196)) = v154
													*(*int64)(unsafe.Add(mBase, uint32(v98)+184)) = v154
													*(*int64)(unsafe.Add(mBase, uint32(v98)+176)) = int64(4294967396)
													v174 = F_dictCreate(m, int32(_a1358))
													mBase = m.M
													v175 = m.ExcPending
													if v175 != 0 {
														return int32(0)
													} else {
														*(*int32)(unsafe.Add(mBase, uint32(v98)+192)) = l5
														*(*int32)(unsafe.Add(mBase, uint32(v98)+156)) = int32(1)
														*(*int32)(unsafe.Add(mBase, uint32(v98)+152)) = l4
														*(*int32)(unsafe.Add(mBase, uint32(v98)+144)) = v174
														v182 = F_dictCreate(m, int32(_a1358))
														mBase = m.M
														v183 = m.ExcPending
														if v183 != 0 {
															return int32(0)
														} else {
															*(*int64)(unsafe.Add(mBase, uint32(v98)+96)) = int64(0)
															*(*int32)(unsafe.Add(mBase, uint32(v98)+148)) = v182
															v188 = F_dictCreate(m, int32(_a1359))
															mBase = m.M
															v189 = m.ExcPending
															if v189 != 0 {
																return int32(0)
															} else {
																v190 = int64(0)
																*(*int64)(unsafe.Add(mBase, uint32(v98)+224)) = v190
																v192 = int32(0)
																*(*int32)(unsafe.Add(mBase, uint32(v98)+216)) = v192
																*(*int32)(unsafe.Add(mBase, uint32(v98)+104)) = v188
																*(*int64)(unsafe.Add(mBase, uint32(v98+int32(232)))) = v190
																*(*int32)(unsafe.Add(mBase, uint32(v98+int32(240)))) = v192
																*(*int64)(unsafe.Add(mBase, uint32(v98)+248)) = v190
																*(*int64)(unsafe.Add(mBase, uint32(v98+int32(256)))) = v190
																v210 = *(*int64)(unsafe.Add(mBase, _consts[735]))
																*(*int64)(unsafe.Add(mBase, uint32(v98)+272)) = v190
																*(*int64)(unsafe.Add(mBase, uint32(v98)+264)) = v210
																*(*int64)(unsafe.Add(mBase, uint32(v98+int32(280)))) = v190
																*(*int64)(unsafe.Add(mBase, uint32(v98+int32(288)))) = v190
																v222 = *(*int32)(unsafe.Add(mBase, uint32(v98)))
																*(*int32)(unsafe.Add(mBase, uint32(v98)+108)) = v222 & int32(3)
																v226 = F_mstime(m)
																mBase = m.M
																*(*int64)(unsafe.Add(mBase, uint32(v98)+112)) = v226
																v228 = F_mstime(m)
																mBase = m.M
																*(*int32)(unsafe.Add(mBase, uint32(v98)+128)) = v192
																*(*int64)(unsafe.Add(mBase, uint32(v98)+120)) = v228
																v232 = F_mstime(m)
																mBase = m.M
																*(*int64)(unsafe.Add(mBase, uint32(v98)+136)) = v232
																v234 = *(*int32)(unsafe.Add(mBase, uint32(v98)+4))
																v235 = F_dictAdd(m, v78, v234, v98)
																mBase = m.M
																v236 = m.ExcPending
																if v236 != 0 {
																	return int32(0)
																} else {
																	v241 = v98
																	m.G0 = v14 + int32(16)
																	return v241
																}
															}
														}
													}
												}
											}
										} else {
											v84 = *(*int32)(unsafe.Add(mBase, uint32(v26)))
											F_sdsfree(m, v84)
											mBase = m.M
											v86 = m.ExcPending
											if v86 != 0 {
												return int32(0)
											} else {
												v87 = *(*int32)(unsafe.Add(mBase, uint32(v26)+4))
												F_sdsfree(m, v87)
												mBase = m.M
												v89 = m.ExcPending
												if v89 != 0 {
													return int32(0)
												} else {
													F_valkey_free(m, v26)
													mBase = m.M
													v91 = m.ExcPending
													if v91 != 0 {
														return int32(0)
													} else {
														F_sdsfree(m, v77)
														mBase = m.M
														v93 = m.ExcPending
														if v93 != 0 {
															return int32(0)
														} else {
															*(*int32)(unsafe.Add(mBase, _consts[5])) = int32(10)
															v241 = v24
															m.G0 = v14 + int32(16)
															return v241
														}
													}
												}
											}
										}
									}
								}
							}
						} else {
							v34 = F_sdsnew(m, l0)
							mBase = m.M
							v35 = m.ExcPending
							if v35 != 0 {
								return int32(0)
							} else {
								if v21 != 0 {
									v71 = v34
									v74 = *(*int32)(unsafe.Add(mBase, _consts[733]))
									v77 = v71
									v78 = v74
								} else {
									v36 = int32(0)
									if l1&int32(4) == v36 {
										v77 = v34
										v78 = v36
									} else {
										v41 = *(*int32)(unsafe.Add(mBase, uint32(l5)+144))
										v77 = v34
										v78 = v41
									}
								}
								v80 = F_dictFind(m, v78, v77)
								mBase = m.M
								v81 = m.ExcPending
								if v81 != 0 {
									return int32(0)
								} else {
									if v80 == int32(0) {
										v98 = F_valkey_malloc(m, int32(296))
										mBase = m.M
										v99 = m.ExcPending
										if v99 != 0 {
											return int32(0)
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(v98)+24)) = v26
											*(*int64)(unsafe.Add(mBase, uint32(v98)+16)) = int64(0)
											*(*int32)(unsafe.Add(mBase, uint32(v98)+8)) = int32(0)
											*(*int32)(unsafe.Add(mBase, uint32(v98)+4)) = v77
											*(*int32)(unsafe.Add(mBase, uint32(v98))) = l1
											v108 = F_valkey_malloc(m, int32(88))
											mBase = m.M
											v109 = m.ExcPending
											if v109 != 0 {
												return int32(0)
											} else {
												v110 = int64(0)
												*(*int64)(unsafe.Add(mBase, uint32(v108)+80)) = v110
												*(*int32)(unsafe.Add(mBase, uint32(v108)+16)) = int32(0)
												*(*int64)(unsafe.Add(mBase, uint32(v108)+8)) = v110
												*(*int64)(unsafe.Add(mBase, uint32(v108))) = int64(4294967297)
												*(*int64)(unsafe.Add(mBase, uint32(v108)+24)) = v110
												*(*int64)(unsafe.Add(mBase, uint32(v108+int32(32)))) = v110
												*(*int64)(unsafe.Add(mBase, uint32(v108+int32(40)))) = v110
												v128 = F_mstime(m)
												mBase = m.M
												*(*int64)(unsafe.Add(mBase, uint32(v108)+64)) = v110
												*(*int64)(unsafe.Add(mBase, uint32(v108)+56)) = v128
												v132 = F_mstime(m)
												mBase = m.M
												*(*int64)(unsafe.Add(mBase, uint32(v108)+48)) = v132
												v134 = F_mstime(m)
												mBase = m.M
												*(*int64)(unsafe.Add(mBase, uint32(v108)+72)) = v134
												*(*int32)(unsafe.Add(mBase, uint32(v98)+28)) = v108
												v137 = F_mstime(m)
												mBase = m.M
												*(*int64)(unsafe.Add(mBase, uint32(v98)+32)) = v137
												v139 = F_mstime(m)
												mBase = m.M
												*(*int64)(unsafe.Add(mBase, uint32(v98)+40)) = v139
												v141 = F_mstime(m)
												mBase = m.M
												*(*int64)(unsafe.Add(mBase, uint32(v98)+56)) = v110
												*(*int64)(unsafe.Add(mBase, uint32(v98)+48)) = v141
												*(*int64)(unsafe.Add(mBase, uint32(v98+int32(64)))) = v110
												if l5 != 0 {
													v152 = l5 + int32(72)
												} else {
													v152 = int32(_a1357)
												}
												v153 = *(*int64)(unsafe.Add(mBase, uint32(v152)))
												v154 = int64(0)
												*(*int64)(unsafe.Add(mBase, uint32(v98)+160)) = v154
												*(*int64)(unsafe.Add(mBase, uint32(v98)+80)) = v154
												*(*int64)(unsafe.Add(mBase, uint32(v98)+72)) = v153
												*(*int64)(unsafe.Add(mBase, uint32(v98+int32(168)))) = v154
												*(*int64)(unsafe.Add(mBase, uint32(v98)+208)) = v154
												*(*int32)(unsafe.Add(mBase, uint32(v98)+204)) = int32(1)
												*(*int64)(unsafe.Add(mBase, uint32(v98)+196)) = v154
												*(*int64)(unsafe.Add(mBase, uint32(v98)+184)) = v154
												*(*int64)(unsafe.Add(mBase, uint32(v98)+176)) = int64(4294967396)
												v174 = F_dictCreate(m, int32(_a1358))
												mBase = m.M
												v175 = m.ExcPending
												if v175 != 0 {
													return int32(0)
												} else {
													*(*int32)(unsafe.Add(mBase, uint32(v98)+192)) = l5
													*(*int32)(unsafe.Add(mBase, uint32(v98)+156)) = int32(1)
													*(*int32)(unsafe.Add(mBase, uint32(v98)+152)) = l4
													*(*int32)(unsafe.Add(mBase, uint32(v98)+144)) = v174
													v182 = F_dictCreate(m, int32(_a1358))
													mBase = m.M
													v183 = m.ExcPending
													if v183 != 0 {
														return int32(0)
													} else {
														*(*int64)(unsafe.Add(mBase, uint32(v98)+96)) = int64(0)
														*(*int32)(unsafe.Add(mBase, uint32(v98)+148)) = v182
														v188 = F_dictCreate(m, int32(_a1359))
														mBase = m.M
														v189 = m.ExcPending
														if v189 != 0 {
															return int32(0)
														} else {
															v190 = int64(0)
															*(*int64)(unsafe.Add(mBase, uint32(v98)+224)) = v190
															v192 = int32(0)
															*(*int32)(unsafe.Add(mBase, uint32(v98)+216)) = v192
															*(*int32)(unsafe.Add(mBase, uint32(v98)+104)) = v188
															*(*int64)(unsafe.Add(mBase, uint32(v98+int32(232)))) = v190
															*(*int32)(unsafe.Add(mBase, uint32(v98+int32(240)))) = v192
															*(*int64)(unsafe.Add(mBase, uint32(v98)+248)) = v190
															*(*int64)(unsafe.Add(mBase, uint32(v98+int32(256)))) = v190
															v210 = *(*int64)(unsafe.Add(mBase, _consts[735]))
															*(*int64)(unsafe.Add(mBase, uint32(v98)+272)) = v190
															*(*int64)(unsafe.Add(mBase, uint32(v98)+264)) = v210
															*(*int64)(unsafe.Add(mBase, uint32(v98+int32(280)))) = v190
															*(*int64)(unsafe.Add(mBase, uint32(v98+int32(288)))) = v190
															v222 = *(*int32)(unsafe.Add(mBase, uint32(v98)))
															*(*int32)(unsafe.Add(mBase, uint32(v98)+108)) = v222 & int32(3)
															v226 = F_mstime(m)
															mBase = m.M
															*(*int64)(unsafe.Add(mBase, uint32(v98)+112)) = v226
															v228 = F_mstime(m)
															mBase = m.M
															*(*int32)(unsafe.Add(mBase, uint32(v98)+128)) = v192
															*(*int64)(unsafe.Add(mBase, uint32(v98)+120)) = v228
															v232 = F_mstime(m)
															mBase = m.M
															*(*int64)(unsafe.Add(mBase, uint32(v98)+136)) = v232
															v234 = *(*int32)(unsafe.Add(mBase, uint32(v98)+4))
															v235 = F_dictAdd(m, v78, v234, v98)
															mBase = m.M
															v236 = m.ExcPending
															if v236 != 0 {
																return int32(0)
															} else {
																v241 = v98
																m.G0 = v14 + int32(16)
																return v241
															}
														}
													}
												}
											}
										}
									} else {
										v84 = *(*int32)(unsafe.Add(mBase, uint32(v26)))
										F_sdsfree(m, v84)
										mBase = m.M
										v86 = m.ExcPending
										if v86 != 0 {
											return int32(0)
										} else {
											v87 = *(*int32)(unsafe.Add(mBase, uint32(v26)+4))
											F_sdsfree(m, v87)
											mBase = m.M
											v89 = m.ExcPending
											if v89 != 0 {
												return int32(0)
											} else {
												F_valkey_free(m, v26)
												mBase = m.M
												v91 = m.ExcPending
												if v91 != 0 {
													return int32(0)
												} else {
													F_sdsfree(m, v77)
													mBase = m.M
													v93 = m.ExcPending
													if v93 != 0 {
														return int32(0)
													} else {
														*(*int32)(unsafe.Add(mBase, _consts[5])) = int32(10)
														v241 = v24
														m.G0 = v14 + int32(16)
														return v241
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
func F_freeSentinelLoadQueueEntry(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	F_sdsfreesplitres(m, v2, v3)
	mBase = m.M
	v5 = m.ExcPending
	if v5 != 0 {
		return
	} else {
		v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
		F_sdsfree(m, v6)
		mBase = m.M
		v8 = m.ExcPending
		if v8 != 0 {
			return
		} else {
			F_valkey_free(m, l0)
			mBase = m.M
			v10 = m.ExcPending
			if v10 != 0 {
				return
			} else {
				return
			}
		}
	}
}
func F_releaseSentinelValkeyInstance(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
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
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
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
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v56 int32
	_ = v56
	var v63 int32
	_ = v63
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+144))
	F_dictRelease(m, v3)
	mBase = m.M
	v5 = m.ExcPending
	if v5 != 0 {
		return
	} else {
		v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+148))
		F_dictRelease(m, v6)
		mBase = m.M
		v8 = m.ExcPending
		if v8 != 0 {
			return
		} else {
			v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
			v10 = F_releaseInstanceLink(m, v9, l0)
			mBase = m.M
			v11 = m.ExcPending
			if v11 != 0 {
				return
			} else {
				v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
				F_sdsfree(m, v12)
				mBase = m.M
				v14 = m.ExcPending
				if v14 != 0 {
					return
				} else {
					v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
					F_sdsfree(m, v15)
					mBase = m.M
					v17 = m.ExcPending
					if v17 != 0 {
						return
					} else {
						v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+284))
						F_sdsfree(m, v18)
						mBase = m.M
						v20 = m.ExcPending
						if v20 != 0 {
							return
						} else {
							v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+288))
							F_sdsfree(m, v21)
							mBase = m.M
							v23 = m.ExcPending
							if v23 != 0 {
								return
							} else {
								v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+196))
								F_sdsfree(m, v24)
								mBase = m.M
								v26 = m.ExcPending
								if v26 != 0 {
									return
								} else {
									v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)+216))
									F_sdsfree(m, v27)
									mBase = m.M
									v29 = m.ExcPending
									if v29 != 0 {
										return
									} else {
										v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)+160))
										F_sdsfree(m, v30)
										mBase = m.M
										v32 = m.ExcPending
										if v32 != 0 {
											return
										} else {
											v33 = *(*int32)(unsafe.Add(mBase, uint32(l0)+164))
											F_sdsfree(m, v33)
											mBase = m.M
											v35 = m.ExcPending
											if v35 != 0 {
												return
											} else {
												v36 = *(*int32)(unsafe.Add(mBase, uint32(l0)+292))
												F_sdsfree(m, v36)
												mBase = m.M
												v38 = m.ExcPending
												if v38 != 0 {
													return
												} else {
													v39 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
													v40 = *(*int32)(unsafe.Add(mBase, uint32(v39)))
													F_sdsfree(m, v40)
													mBase = m.M
													v42 = m.ExcPending
													if v42 != 0 {
														return
													} else {
														v43 = *(*int32)(unsafe.Add(mBase, uint32(v39)+4))
														F_sdsfree(m, v43)
														mBase = m.M
														v45 = m.ExcPending
														if v45 != 0 {
															return
														} else {
															F_valkey_free(m, v39)
															mBase = m.M
															v47 = m.ExcPending
															if v47 != 0 {
																return
															} else {
																v48 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
																F_dictRelease(m, v48)
																mBase = m.M
																v50 = m.ExcPending
																if v50 != 0 {
																	return
																} else {
																	v51 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
																	v52 = int32(130)
																	if v51&v52 != v52 {
																	} else {
																		v56 = *(*int32)(unsafe.Add(mBase, uint32(l0)+192))
																		if v56 == int32(0) {
																		} else {
																			*(*int32)(unsafe.Add(mBase, uint32(v56)+280)) = int32(0)
																		}
																	}
																	F_valkey_free(m, l0)
																	mBase = m.M
																	v63 = m.ExcPending
																	if v63 != 0 {
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
func F_sentinelCheckConfigFile(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v17 int32
	_ = v17
	var v21 int32
	_ = v21
	var v25 int32
	_ = v25
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v38 int32
	_ = v38
	v3 = m.G0
	v5 = v3 - int32(16)
	m.G0 = v5
	v8 = *(*int32)(unsafe.Add(mBase, _consts[392]))
	if v8 != 0 {
		v21 = F_access(m, v8, int32(2))
		mBase = m.M
		if v21 != int32(-1) {
			m.G0 = v5 + int32(16)
			return
		} else {
			v25 = *(*int32)(unsafe.Add(mBase, _consts[6]))
			if int32(3) < v25 {
				m.Env.Exit(m, int32(1))
				mBase = m.M
				base.Wasm_trap_unreachable()
				for {
				}
			} else {
				v29 = *(*int32)(unsafe.Add(mBase, _consts[392]))
				v31 = *(*int32)(unsafe.Add(mBase, _consts[5]))
				v32 = F___strerror_l(m, v31, v31)
				mBase = m.M
				*(*int32)(unsafe.Add(mBase, uint32(v5)+4)) = v32
				*(*int32)(unsafe.Add(mBase, uint32(v5))) = v29
				F__serverLog(m, int32(3), int32(_a1334), v5)
				mBase = m.M
				v38 = m.ExcPending
				if v38 != 0 {
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
	} else {
		v10 = *(*int32)(unsafe.Add(mBase, _consts[6]))
		if int32(3) < v10 {
			m.Env.Exit(m, int32(1))
			mBase = m.M
			base.Wasm_trap_unreachable()
			for {
			}
		} else {
			F__serverLog(m, int32(3), int32(_a1335), int32(0))
			mBase = m.M
			v17 = m.ExcPending
			if v17 != 0 {
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
func F_sentinelConfigGetCommand(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v56 int32
	_ = v56
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
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
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
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
	var v98 int32
	_ = v98
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v109 int32
	_ = v109
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v116 int32
	_ = v116
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v126 int32
	_ = v126
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v140 int32
	_ = v140
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v151 int32
	_ = v151
	var v153 int32
	_ = v153
	var v155 int32
	_ = v155
	var v157 int32
	_ = v157
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v164 int32
	_ = v164
	var v166 int32
	_ = v166
	var v168 int32
	_ = v168
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v182 int32
	_ = v182
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v193 int32
	_ = v193
	var v195 int64
	_ = v195
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
	var v207 int32
	_ = v207
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v221 int32
	_ = v221
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v232 int32
	_ = v232
	var v234 int32
	_ = v234
	var v236 int32
	_ = v236
	var v238 int32
	_ = v238
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v245 int32
	_ = v245
	var v247 int32
	_ = v247
	var v249 int32
	_ = v249
	var v252 int32
	_ = v252
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v263 int32
	_ = v263
	var v270 int32
	_ = v270
	var v271 int32
	_ = v271
	var v274 int32
	_ = v274
	var v276 int32
	_ = v276
	var v278 int32
	_ = v278
	var v280 int32
	_ = v280
	var v283 int32
	_ = v283
	var v284 int32
	_ = v284
	var v287 int32
	_ = v287
	var v289 int32
	_ = v289
	var v291 int32
	_ = v291
	var v294 int32
	_ = v294
	var v295 int32
	_ = v295
	var v296 int32
	_ = v296
	var v298 int32
	_ = v298
	var v299 int32
	_ = v299
	var v305 int32
	_ = v305
	var v312 int32
	_ = v312
	var v313 int32
	_ = v313
	var v316 int32
	_ = v316
	var v319 int32
	_ = v319
	var v326 int32
	_ = v326
	var v327 int32
	_ = v327
	var v329 int32
	_ = v329
	var v332 int32
	_ = v332
	var v333 int32
	_ = v333
	var v336 int32
	_ = v336
	var v340 int32
	_ = v340
	var v341 int32
	_ = v341
	var v346 int32
	_ = v346
	var v351 int32
	_ = v351
	var v353 int32
	_ = v353
	v8 = F_addReplyDeferredLen(m, l0)
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v11 = F_dictCreate(m, int32(_a492))
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v14 < int32(4) {
		v346 = int32(0)
		goto L4
	} else {
		goto L5
	}
L4:
	;
	F_dictRelease(m, v11)
	mBase = m.M
	v351 = m.ExcPending
	if v351 != 0 {
		goto L1
	} else {
		goto L90
	}
L5:
	;
	v22 = int32(0)
	v23 = int32(3)
	goto L6
L6:
	;
	v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v26+v23<<(uint(int32(2))%32))))
	v31 = F_objectGetVal(m, v30)
	mBase = m.M
	v33 = F_strcspn(m, v31, int32(_a493))
	mBase = m.M
	v34 = v31 + v33
	v36 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v34))))
	if v36 != 0 {
		goto L11
	} else {
		goto L12
	}
L7:
	;
	v346 = v336
	goto L4
L8:
	;
	v340 = v23 + int32(1)
	v341 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v340 < v341 {
		v22 = v336
		v23 = v340
		goto L6
	} else {
		goto L89
	}
L9:
	;
	v40 = int32(_a1450)
	v42 = int32(0)
	v45 = m.G0
	v46 = int32(16)
	v47 = v45 - v46
	m.G0 = v47
	v49 = F_strlen(m, v31)
	mBase = m.M
	v50 = F_strlen(m, v40)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v47)+12)) = v42
	v56 = F_stringmatchlen_impl(m, v31, v49, v40, v50, int32(1), v47+int32(12), v42)
	mBase = m.M
	m.G0 = v47 + v46
	goto L18
L10:
	;
	if v37 != 0 {
		goto L9
	} else {
		goto L14
	}
L11:
	;
	v37 = v34
	goto L13
L12:
	;
	v37 = int32(0)
	goto L13
L13:
	;
	goto L10
L14:
	;
	v38 = F_dictFind(m, v11, v31)
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L1
	} else {
		goto L15
	}
L15:
	;
	if v38 != 0 {
		v336 = v22
		goto L8
	} else {
		goto L16
	}
L16:
	;
	goto L9
L17:
	;
	v82 = int32(_a1452)
	v84 = int32(0)
	v87 = m.G0
	v88 = int32(16)
	v89 = v87 - v88
	m.G0 = v89
	v91 = F_strlen(m, v31)
	mBase = m.M
	v92 = F_strlen(m, v82)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v89)+12)) = v84
	v98 = F_stringmatchlen_impl(m, v31, v91, v82, v92, int32(1), v89+int32(12), v84)
	mBase = m.M
	m.G0 = v89 + v88
	goto L29
L18:
	;
	if v56 == int32(0) {
		v81 = v22
		goto L17
	} else {
		goto L19
	}
L19:
	;
	v63 = F_dictFind(m, v11, int32(_a1450))
	mBase = m.M
	v64 = m.ExcPending
	if v64 != 0 {
		goto L1
	} else {
		goto L20
	}
L20:
	;
	if v63 != 0 {
		v81 = v22
		goto L17
	} else {
		goto L21
	}
L21:
	;
	F_addReplyBulkCString(m, l0, int32(_a1450))
	mBase = m.M
	v67 = m.ExcPending
	if v67 != 0 {
		goto L1
	} else {
		goto L22
	}
L22:
	;
	v71 = *(*int32)(unsafe.Add(mBase, _consts[729]))
	if v71 != 0 {
		goto L23
	} else {
		goto L24
	}
L23:
	;
	v72 = int32(_a483)
	goto L25
L24:
	;
	v72 = int32(_a484)
	goto L25
L25:
	;
	F_addReplyBulkCString(m, l0, v72)
	mBase = m.M
	v74 = m.ExcPending
	if v74 != 0 {
		goto L1
	} else {
		goto L26
	}
L26:
	;
	v77 = F_dictAdd(m, v11, int32(_a1450), int32(0))
	mBase = m.M
	v78 = m.ExcPending
	if v78 != 0 {
		goto L1
	} else {
		goto L27
	}
L27:
	;
	v81 = v22 + int32(1)
	goto L17
L28:
	;
	v124 = int32(_a1447)
	v126 = int32(0)
	v129 = m.G0
	v130 = int32(16)
	v131 = v129 - v130
	m.G0 = v131
	v133 = F_strlen(m, v31)
	mBase = m.M
	v134 = F_strlen(m, v124)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v131)+12)) = v126
	v140 = F_stringmatchlen_impl(m, v31, v133, v124, v134, int32(1), v131+int32(12), v126)
	mBase = m.M
	m.G0 = v131 + v130
	goto L40
L29:
	;
	if v98 == int32(0) {
		v123 = v81
		goto L28
	} else {
		goto L30
	}
L30:
	;
	v105 = F_dictFind(m, v11, int32(_a1452))
	mBase = m.M
	v106 = m.ExcPending
	if v106 != 0 {
		goto L1
	} else {
		goto L31
	}
L31:
	;
	if v105 != 0 {
		v123 = v81
		goto L28
	} else {
		goto L32
	}
L32:
	;
	F_addReplyBulkCString(m, l0, int32(_a1452))
	mBase = m.M
	v109 = m.ExcPending
	if v109 != 0 {
		goto L1
	} else {
		goto L33
	}
L33:
	;
	v113 = *(*int32)(unsafe.Add(mBase, _consts[734]))
	if v113 != 0 {
		goto L34
	} else {
		goto L35
	}
L34:
	;
	v114 = int32(_a483)
	goto L36
L35:
	;
	v114 = int32(_a484)
	goto L36
L36:
	;
	F_addReplyBulkCString(m, l0, v114)
	mBase = m.M
	v116 = m.ExcPending
	if v116 != 0 {
		goto L1
	} else {
		goto L37
	}
L37:
	;
	v119 = F_dictAdd(m, v11, int32(_a1452), int32(0))
	mBase = m.M
	v120 = m.ExcPending
	if v120 != 0 {
		goto L1
	} else {
		goto L38
	}
L38:
	;
	v123 = v81 + int32(1)
	goto L28
L39:
	;
	v166 = int32(_a1451)
	v168 = int32(0)
	v171 = m.G0
	v172 = int32(16)
	v173 = v171 - v172
	m.G0 = v173
	v175 = F_strlen(m, v31)
	mBase = m.M
	v176 = F_strlen(m, v166)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v173)+12)) = v168
	v182 = F_stringmatchlen_impl(m, v31, v175, v166, v176, int32(1), v173+int32(12), v168)
	mBase = m.M
	m.G0 = v173 + v172
	goto L51
L40:
	;
	if v140 == int32(0) {
		v164 = v123
		goto L39
	} else {
		goto L41
	}
L41:
	;
	v147 = F_dictFind(m, v11, int32(_a1447))
	mBase = m.M
	v148 = m.ExcPending
	if v148 != 0 {
		goto L1
	} else {
		goto L42
	}
L42:
	;
	if v147 != 0 {
		v164 = v123
		goto L39
	} else {
		goto L43
	}
L43:
	;
	F_addReplyBulkCString(m, l0, int32(_a1447))
	mBase = m.M
	v151 = m.ExcPending
	if v151 != 0 {
		goto L1
	} else {
		goto L44
	}
L44:
	;
	v153 = *(*int32)(unsafe.Add(mBase, _consts[749]))
	if v153 != 0 {
		goto L45
	} else {
		goto L46
	}
L45:
	;
	v155 = v153
	goto L47
L46:
	;
	v155 = int32(_a139)
	goto L47
L47:
	;
	F_addReplyBulkCString(m, l0, v155)
	mBase = m.M
	v157 = m.ExcPending
	if v157 != 0 {
		goto L1
	} else {
		goto L48
	}
L48:
	;
	v160 = F_dictAdd(m, v11, int32(_a1447), int32(0))
	mBase = m.M
	v161 = m.ExcPending
	if v161 != 0 {
		goto L1
	} else {
		goto L49
	}
L49:
	;
	v164 = v123 + int32(1)
	goto L39
L50:
	;
	v205 = int32(_a1448)
	v207 = int32(0)
	v210 = m.G0
	v211 = int32(16)
	v212 = v210 - v211
	m.G0 = v212
	v214 = F_strlen(m, v31)
	mBase = m.M
	v215 = F_strlen(m, v205)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v212)+12)) = v207
	v221 = F_stringmatchlen_impl(m, v31, v214, v205, v215, int32(1), v212+int32(12), v207)
	mBase = m.M
	m.G0 = v212 + v211
	goto L59
L51:
	;
	if v182 == int32(0) {
		v204 = v164
		goto L50
	} else {
		goto L52
	}
L52:
	;
	v189 = F_dictFind(m, v11, int32(_a1451))
	mBase = m.M
	v190 = m.ExcPending
	if v190 != 0 {
		goto L1
	} else {
		goto L53
	}
L53:
	;
	if v189 != 0 {
		v204 = v164
		goto L50
	} else {
		goto L54
	}
L54:
	;
	F_addReplyBulkCString(m, l0, int32(_a1451))
	mBase = m.M
	v193 = m.ExcPending
	if v193 != 0 {
		goto L1
	} else {
		goto L55
	}
L55:
	;
	v195 = int64(*(*int32)(unsafe.Add(mBase, _consts[750])))
	F_addReplyBulkLongLong(m, l0, v195)
	mBase = m.M
	v197 = m.ExcPending
	if v197 != 0 {
		goto L1
	} else {
		goto L56
	}
L56:
	;
	v200 = F_dictAdd(m, v11, int32(_a1451), int32(0))
	mBase = m.M
	v201 = m.ExcPending
	if v201 != 0 {
		goto L1
	} else {
		goto L57
	}
L57:
	;
	v204 = v164 + int32(1)
	goto L50
L58:
	;
	v247 = int32(_a1449)
	v249 = int32(0)
	v252 = m.G0
	v253 = int32(16)
	v254 = v252 - v253
	m.G0 = v254
	v256 = F_strlen(m, v31)
	mBase = m.M
	v257 = F_strlen(m, v247)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v254)+12)) = v249
	v263 = F_stringmatchlen_impl(m, v31, v256, v247, v257, int32(1), v254+int32(12), v249)
	mBase = m.M
	m.G0 = v254 + v253
	goto L70
L59:
	;
	if v221 == int32(0) {
		v245 = v204
		goto L58
	} else {
		goto L60
	}
L60:
	;
	v228 = F_dictFind(m, v11, int32(_a1448))
	mBase = m.M
	v229 = m.ExcPending
	if v229 != 0 {
		goto L1
	} else {
		goto L61
	}
L61:
	;
	if v228 != 0 {
		v245 = v204
		goto L58
	} else {
		goto L62
	}
L62:
	;
	F_addReplyBulkCString(m, l0, int32(_a1448))
	mBase = m.M
	v232 = m.ExcPending
	if v232 != 0 {
		goto L1
	} else {
		goto L63
	}
L63:
	;
	v234 = *(*int32)(unsafe.Add(mBase, _consts[751]))
	if v234 != 0 {
		goto L64
	} else {
		goto L65
	}
L64:
	;
	v236 = v234
	goto L66
L65:
	;
	v236 = int32(_a139)
	goto L66
L66:
	;
	F_addReplyBulkCString(m, l0, v236)
	mBase = m.M
	v238 = m.ExcPending
	if v238 != 0 {
		goto L1
	} else {
		goto L67
	}
L67:
	;
	v241 = F_dictAdd(m, v11, int32(_a1448), int32(0))
	mBase = m.M
	v242 = m.ExcPending
	if v242 != 0 {
		goto L1
	} else {
		goto L68
	}
L68:
	;
	v245 = v204 + int32(1)
	goto L58
L69:
	;
	v289 = int32(_a1453)
	v291 = int32(0)
	v294 = m.G0
	v295 = int32(16)
	v296 = v294 - v295
	m.G0 = v296
	v298 = F_strlen(m, v31)
	mBase = m.M
	v299 = F_strlen(m, v289)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v296)+12)) = v291
	v305 = F_stringmatchlen_impl(m, v31, v298, v289, v299, int32(1), v296+int32(12), v291)
	mBase = m.M
	m.G0 = v296 + v295
	goto L80
L70:
	;
	if v263 == int32(0) {
		v287 = v245
		goto L69
	} else {
		goto L71
	}
L71:
	;
	v270 = F_dictFind(m, v11, int32(_a1449))
	mBase = m.M
	v271 = m.ExcPending
	if v271 != 0 {
		goto L1
	} else {
		goto L72
	}
L72:
	;
	if v270 != 0 {
		v287 = v245
		goto L69
	} else {
		goto L73
	}
L73:
	;
	F_addReplyBulkCString(m, l0, int32(_a1449))
	mBase = m.M
	v274 = m.ExcPending
	if v274 != 0 {
		goto L1
	} else {
		goto L74
	}
L74:
	;
	v276 = *(*int32)(unsafe.Add(mBase, _consts[752]))
	if v276 != 0 {
		goto L75
	} else {
		goto L76
	}
L75:
	;
	v278 = v276
	goto L77
L76:
	;
	v278 = int32(_a139)
	goto L77
L77:
	;
	F_addReplyBulkCString(m, l0, v278)
	mBase = m.M
	v280 = m.ExcPending
	if v280 != 0 {
		goto L1
	} else {
		goto L78
	}
L78:
	;
	v283 = F_dictAdd(m, v11, int32(_a1449), int32(0))
	mBase = m.M
	v284 = m.ExcPending
	if v284 != 0 {
		goto L1
	} else {
		goto L79
	}
L79:
	;
	v287 = v245 + int32(1)
	goto L69
L80:
	;
	if v305 == int32(0) {
		v336 = v287
		goto L8
	} else {
		goto L81
	}
L81:
	;
	v312 = F_dictFind(m, v11, int32(_a1453))
	mBase = m.M
	v313 = m.ExcPending
	if v313 != 0 {
		goto L1
	} else {
		goto L82
	}
L82:
	;
	if v312 != 0 {
		v336 = v287
		goto L8
	} else {
		goto L83
	}
L83:
	;
	F_addReplyBulkCString(m, l0, int32(_a1453))
	mBase = m.M
	v316 = m.ExcPending
	if v316 != 0 {
		goto L1
	} else {
		goto L84
	}
L84:
	;
	v319 = *(*int32)(unsafe.Add(mBase, _consts[6]))
	if base.Ui32(int32(4)) < base.Ui32(v319) {
		v327 = int32(_a288)
		goto L85
	} else {
		goto L86
	}
L85:
	;
	F_addReplyBulkCString(m, l0, v327)
	mBase = m.M
	v329 = m.ExcPending
	if v329 != 0 {
		goto L1
	} else {
		goto L87
	}
L86:
	;
	v326 = *(*int32)(unsafe.Add(mBase, uint32(v319<<(uint(int32(2))%32))+uint32(_consts[769])))
	v327 = v326
	goto L85
L87:
	;
	v332 = F_dictAdd(m, v11, int32(_a1453), int32(0))
	mBase = m.M
	v333 = m.ExcPending
	if v333 != 0 {
		goto L1
	} else {
		goto L88
	}
L88:
	;
	v336 = v287 + int32(1)
	goto L8
L89:
	;
	goto L7
L90:
	;
	F_setDeferredMapLen(m, l0, v8, v346)
	mBase = m.M
	v353 = m.ExcPending
	if v353 != 0 {
		goto L1
	} else {
		goto L91
	}
L91:
	;
	return
}
func F_sentinelConfigSetCommand(m *base.Module, l0 int32) {
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
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v89 int32
	_ = v89
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v101 int32
	_ = v101
	var v108 int32
	_ = v108
	var v113 int32
	_ = v113
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
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
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
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
	var v170 int32
	_ = v170
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v196 int32
	_ = v196
	var v198 int32
	_ = v198
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v208 int32
	_ = v208
	var v211 int32
	_ = v211
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v217 int32
	_ = v217
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
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v247 int32
	_ = v247
	var v249 int32
	_ = v249
	var v254 int32
	_ = v254
	var v255 int32
	_ = v255
	var v256 int32
	_ = v256
	var v259 int32
	_ = v259
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
	var v272 int32
	_ = v272
	var v274 int32
	_ = v274
	var v275 int32
	_ = v275
	var v276 int32
	_ = v276
	var v278 int32
	_ = v278
	var v279 int32
	_ = v279
	var v280 int32
	_ = v280
	var v281 int32
	_ = v281
	var v286 int32
	_ = v286
	var v287 int32
	_ = v287
	var v290 int32
	_ = v290
	var v291 int32
	_ = v291
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
	var v304 int64
	_ = v304
	var v307 int32
	_ = v307
	var v310 int32
	_ = v310
	var v312 int32
	_ = v312
	var v313 int32
	_ = v313
	var v314 int32
	_ = v314
	var v316 int32
	_ = v316
	var v320 int32
	_ = v320
	var v322 int32
	_ = v322
	var v323 int32
	_ = v323
	var v324 int32
	_ = v324
	var v326 int32
	_ = v326
	var v327 int32
	_ = v327
	var v328 int32
	_ = v328
	var v329 int32
	_ = v329
	var v334 int32
	_ = v334
	var v335 int32
	_ = v335
	var v338 int32
	_ = v338
	var v339 int32
	_ = v339
	var v342 int32
	_ = v342
	var v343 int32
	_ = v343
	var v344 int32
	_ = v344
	var v346 int32
	_ = v346
	var v347 int32
	_ = v347
	var v350 int32
	_ = v350
	var v352 int32
	_ = v352
	var v353 int32
	_ = v353
	var v354 int32
	_ = v354
	var v356 int32
	_ = v356
	var v360 int32
	_ = v360
	var v362 int32
	_ = v362
	var v363 int32
	_ = v363
	var v364 int32
	_ = v364
	var v366 int32
	_ = v366
	var v367 int32
	_ = v367
	var v368 int32
	_ = v368
	var v369 int32
	_ = v369
	var v374 int32
	_ = v374
	var v375 int32
	_ = v375
	var v378 int32
	_ = v378
	var v379 int32
	_ = v379
	var v382 int32
	_ = v382
	var v383 int32
	_ = v383
	var v384 int32
	_ = v384
	var v388 int32
	_ = v388
	var v389 int32
	_ = v389
	var v392 int32
	_ = v392
	var v394 int32
	_ = v394
	var v395 int32
	_ = v395
	var v396 int32
	_ = v396
	var v398 int32
	_ = v398
	var v402 int32
	_ = v402
	var v404 int32
	_ = v404
	var v405 int32
	_ = v405
	var v406 int32
	_ = v406
	var v408 int32
	_ = v408
	var v409 int32
	_ = v409
	var v410 int32
	_ = v410
	var v411 int32
	_ = v411
	var v416 int32
	_ = v416
	var v417 int32
	_ = v417
	var v420 int32
	_ = v420
	var v421 int32
	_ = v421
	var v424 int32
	_ = v424
	var v425 int32
	_ = v425
	var v426 int32
	_ = v426
	var v430 int32
	_ = v430
	var v431 int32
	_ = v431
	var v434 int32
	_ = v434
	var v436 int32
	_ = v436
	var v437 int32
	_ = v437
	var v438 int32
	_ = v438
	var v440 int32
	_ = v440
	var v444 int32
	_ = v444
	var v446 int32
	_ = v446
	var v447 int32
	_ = v447
	var v448 int32
	_ = v448
	var v450 int32
	_ = v450
	var v451 int32
	_ = v451
	var v452 int32
	_ = v452
	var v453 int32
	_ = v453
	var v458 int32
	_ = v458
	var v459 int32
	_ = v459
	var v462 int32
	_ = v462
	var v463 int32
	_ = v463
	var v466 int32
	_ = v466
	var v467 int32
	_ = v467
	var v468 int32
	_ = v468
	var v472 int32
	_ = v472
	var v473 int32
	_ = v473
	var v476 int32
	_ = v476
	var v478 int32
	_ = v478
	var v479 int32
	_ = v479
	var v480 int32
	_ = v480
	var v482 int32
	_ = v482
	var v486 int32
	_ = v486
	var v488 int32
	_ = v488
	var v489 int32
	_ = v489
	var v490 int32
	_ = v490
	var v492 int32
	_ = v492
	var v493 int32
	_ = v493
	var v494 int32
	_ = v494
	var v495 int32
	_ = v495
	var v500 int32
	_ = v500
	var v501 int32
	_ = v501
	var v504 int32
	_ = v504
	var v505 int32
	_ = v505
	var v508 int32
	_ = v508
	var v509 int32
	_ = v509
	var v510 int32
	_ = v510
	var v514 int32
	_ = v514
	var v515 int32
	_ = v515
	var v518 int32
	_ = v518
	var v520 int32
	_ = v520
	var v521 int32
	_ = v521
	var v522 int32
	_ = v522
	var v524 int32
	_ = v524
	var v528 int32
	_ = v528
	var v530 int32
	_ = v530
	var v531 int32
	_ = v531
	var v532 int32
	_ = v532
	var v534 int32
	_ = v534
	var v535 int32
	_ = v535
	var v536 int32
	_ = v536
	var v537 int32
	_ = v537
	var v542 int32
	_ = v542
	var v543 int32
	_ = v543
	var v546 int32
	_ = v546
	var v547 int32
	_ = v547
	var v550 int32
	_ = v550
	var v551 int32
	_ = v551
	var v552 int32
	_ = v552
	var v555 int32
	_ = v555
	var v556 int32
	_ = v556
	var v562 int32
	_ = v562
	var v564 int32
	_ = v564
	var v571 int32
	_ = v571
	var v578 int32
	_ = v578
	var v579 int32
	_ = v579
	var v582 int32
	_ = v582
	var v584 int32
	_ = v584
	var v588 int32
	_ = v588
	var v589 int32
	_ = v589
	var v590 int32
	_ = v590
	var v593 int32
	_ = v593
	var v595 int32
	_ = v595
	var v596 int32
	_ = v596
	var v597 int32
	_ = v597
	var v599 int32
	_ = v599
	var v603 int32
	_ = v603
	var v605 int32
	_ = v605
	var v606 int32
	_ = v606
	var v607 int32
	_ = v607
	var v609 int32
	_ = v609
	var v610 int32
	_ = v610
	var v611 int32
	_ = v611
	var v612 int32
	_ = v612
	var v617 int32
	_ = v617
	var v618 int32
	_ = v618
	var v621 int32
	_ = v621
	var v622 int32
	_ = v622
	var v625 int32
	_ = v625
	var v626 int32
	_ = v626
	var v627 int32
	_ = v627
	var v631 int32
	_ = v631
	var v634 int32
	_ = v634
	var v637 int32
	_ = v637
	var v639 int32
	_ = v639
	var v640 int32
	_ = v640
	var v641 int32
	_ = v641
	var v643 int32
	_ = v643
	var v647 int32
	_ = v647
	var v649 int32
	_ = v649
	var v650 int32
	_ = v650
	var v651 int32
	_ = v651
	var v653 int32
	_ = v653
	var v654 int32
	_ = v654
	var v655 int32
	_ = v655
	var v656 int32
	_ = v656
	var v661 int32
	_ = v661
	var v662 int32
	_ = v662
	var v665 int32
	_ = v665
	var v666 int32
	_ = v666
	var v669 int32
	_ = v669
	var v670 int32
	_ = v670
	var v671 int32
	_ = v671
	var v674 int32
	_ = v674
	var v676 int32
	_ = v676
	var v680 int32
	_ = v680
	var v681 int32
	_ = v681
	var v683 int32
	_ = v683
	var v688 int32
	_ = v688
	var v689 int32
	_ = v689
	var v690 int32
	_ = v690
	var v694 int32
	_ = v694
	var v696 int32
	_ = v696
	var v700 int32
	_ = v700
	var v701 int32
	_ = v701
	var v702 int32
	_ = v702
	var v705 int32
	_ = v705
	var v707 int32
	_ = v707
	var v708 int32
	_ = v708
	var v709 int32
	_ = v709
	var v711 int32
	_ = v711
	var v715 int32
	_ = v715
	var v717 int32
	_ = v717
	var v718 int32
	_ = v718
	var v719 int32
	_ = v719
	var v721 int32
	_ = v721
	var v722 int32
	_ = v722
	var v723 int32
	_ = v723
	var v724 int32
	_ = v724
	var v729 int32
	_ = v729
	var v730 int32
	_ = v730
	var v733 int32
	_ = v733
	var v734 int32
	_ = v734
	var v737 int32
	_ = v737
	var v738 int32
	_ = v738
	var v739 int32
	_ = v739
	var v744 int32
	_ = v744
	var v745 int32
	_ = v745
	var v748 int32
	_ = v748
	var v750 int32
	_ = v750
	var v751 int32
	_ = v751
	var v752 int32
	_ = v752
	var v754 int32
	_ = v754
	var v758 int32
	_ = v758
	var v760 int32
	_ = v760
	var v761 int32
	_ = v761
	var v762 int32
	_ = v762
	var v764 int32
	_ = v764
	var v765 int32
	_ = v765
	var v766 int32
	_ = v766
	var v767 int32
	_ = v767
	var v772 int32
	_ = v772
	var v773 int32
	_ = v773
	var v776 int32
	_ = v776
	var v777 int32
	_ = v777
	var v780 int32
	_ = v780
	var v781 int32
	_ = v781
	var v782 int32
	_ = v782
	var v787 int32
	_ = v787
	var v788 int32
	_ = v788
	var v791 int32
	_ = v791
	var v793 int32
	_ = v793
	var v794 int32
	_ = v794
	var v795 int32
	_ = v795
	var v797 int32
	_ = v797
	var v801 int32
	_ = v801
	var v803 int32
	_ = v803
	var v804 int32
	_ = v804
	var v805 int32
	_ = v805
	var v807 int32
	_ = v807
	var v808 int32
	_ = v808
	var v809 int32
	_ = v809
	var v810 int32
	_ = v810
	var v815 int32
	_ = v815
	var v816 int32
	_ = v816
	var v819 int32
	_ = v819
	var v820 int32
	_ = v820
	var v823 int32
	_ = v823
	var v824 int32
	_ = v824
	var v825 int32
	_ = v825
	var v830 int32
	_ = v830
	var v831 int32
	_ = v831
	var v834 int32
	_ = v834
	var v836 int32
	_ = v836
	var v837 int32
	_ = v837
	var v838 int32
	_ = v838
	var v840 int32
	_ = v840
	var v844 int32
	_ = v844
	var v846 int32
	_ = v846
	var v847 int32
	_ = v847
	var v848 int32
	_ = v848
	var v850 int32
	_ = v850
	var v851 int32
	_ = v851
	var v852 int32
	_ = v852
	var v853 int32
	_ = v853
	var v858 int32
	_ = v858
	var v859 int32
	_ = v859
	var v862 int32
	_ = v862
	var v863 int32
	_ = v863
	var v866 int32
	_ = v866
	var v867 int32
	_ = v867
	var v868 int32
	_ = v868
	var v873 int32
	_ = v873
	var v874 int32
	_ = v874
	var v877 int32
	_ = v877
	var v879 int32
	_ = v879
	var v880 int32
	_ = v880
	var v881 int32
	_ = v881
	var v883 int32
	_ = v883
	var v887 int32
	_ = v887
	var v889 int32
	_ = v889
	var v890 int32
	_ = v890
	var v891 int32
	_ = v891
	var v893 int32
	_ = v893
	var v894 int32
	_ = v894
	var v895 int32
	_ = v895
	var v896 int32
	_ = v896
	var v901 int32
	_ = v901
	var v902 int32
	_ = v902
	var v905 int32
	_ = v905
	var v906 int32
	_ = v906
	var v909 int32
	_ = v909
	var v910 int32
	_ = v910
	var v911 int32
	_ = v911
	var v916 int32
	_ = v916
	var v919 int32
	_ = v919
	var v921 int32
	_ = v921
	var v922 int32
	_ = v922
	var v923 int32
	_ = v923
	var v925 int32
	_ = v925
	var v929 int32
	_ = v929
	var v931 int32
	_ = v931
	var v932 int32
	_ = v932
	var v933 int32
	_ = v933
	var v935 int32
	_ = v935
	var v936 int32
	_ = v936
	var v937 int32
	_ = v937
	var v938 int32
	_ = v938
	var v943 int32
	_ = v943
	var v944 int32
	_ = v944
	var v947 int32
	_ = v947
	var v948 int32
	_ = v948
	var v951 int32
	_ = v951
	var v952 int32
	_ = v952
	var v953 int32
	_ = v953
	var v956 int32
	_ = v956
	var v958 int32
	_ = v958
	var v962 int32
	_ = v962
	var v963 int32
	_ = v963
	var v965 int32
	_ = v965
	var v970 int32
	_ = v970
	var v971 int32
	_ = v971
	var v972 int32
	_ = v972
	var v976 int32
	_ = v976
	var v979 int32
	_ = v979
	var v981 int32
	_ = v981
	var v982 int32
	_ = v982
	var v983 int32
	_ = v983
	var v985 int32
	_ = v985
	var v989 int32
	_ = v989
	var v991 int32
	_ = v991
	var v992 int32
	_ = v992
	var v993 int32
	_ = v993
	var v995 int32
	_ = v995
	var v996 int32
	_ = v996
	var v997 int32
	_ = v997
	var v998 int32
	_ = v998
	var v1003 int32
	_ = v1003
	var v1004 int32
	_ = v1004
	var v1007 int32
	_ = v1007
	var v1008 int32
	_ = v1008
	var v1011 int32
	_ = v1011
	var v1012 int32
	_ = v1012
	var v1013 int32
	_ = v1013
	var v1015 int32
	_ = v1015
	var v1017 int32
	_ = v1017
	var v1021 int32
	_ = v1021
	var v1022 int32
	_ = v1022
	var v1023 int32
	_ = v1023
	var v1027 int32
	_ = v1027
	var v1029 int32
	_ = v1029
	var v1030 int32
	_ = v1030
	var v1031 int32
	_ = v1031
	var v1033 int32
	_ = v1033
	var v1036 int32
	_ = v1036
	var v1038 int32
	_ = v1038
	var v1039 int32
	_ = v1039
	var v1040 int32
	_ = v1040
	var v1042 int32
	_ = v1042
	var v1046 int32
	_ = v1046
	var v1048 int32
	_ = v1048
	var v1049 int32
	_ = v1049
	var v1050 int32
	_ = v1050
	var v1052 int32
	_ = v1052
	var v1053 int32
	_ = v1053
	var v1054 int32
	_ = v1054
	var v1055 int32
	_ = v1055
	var v1060 int32
	_ = v1060
	var v1061 int32
	_ = v1061
	var v1064 int32
	_ = v1064
	var v1065 int32
	_ = v1065
	var v1068 int32
	_ = v1068
	var v1069 int32
	_ = v1069
	var v1070 int32
	_ = v1070
	var v1072 int32
	_ = v1072
	var v1074 int32
	_ = v1074
	var v1078 int32
	_ = v1078
	var v1081 int32
	_ = v1081
	var v1082 int32
	_ = v1082
	var v1084 int64
	_ = v1084
	var v1086 int32
	_ = v1086
	var v1089 int32
	_ = v1089
	var v1091 int32
	_ = v1091
	var v1092 int32
	_ = v1092
	var v1093 int32
	_ = v1093
	var v1095 int32
	_ = v1095
	var v1099 int32
	_ = v1099
	var v1101 int32
	_ = v1101
	var v1102 int32
	_ = v1102
	var v1103 int32
	_ = v1103
	var v1105 int32
	_ = v1105
	var v1106 int32
	_ = v1106
	var v1107 int32
	_ = v1107
	var v1108 int32
	_ = v1108
	var v1113 int32
	_ = v1113
	var v1114 int32
	_ = v1114
	var v1117 int32
	_ = v1117
	var v1118 int32
	_ = v1118
	var v1121 int32
	_ = v1121
	var v1122 int32
	_ = v1122
	var v1123 int32
	_ = v1123
	var v1125 int32
	_ = v1125
	var v1127 int32
	_ = v1127
	var v1131 int32
	_ = v1131
	var v1132 int32
	_ = v1132
	var v1134 int32
	_ = v1134
	var v1136 int32
	_ = v1136
	var v1137 int32
	_ = v1137
	var v1140 int32
	_ = v1140
	var v1147 int32
	_ = v1147
	var v1150 int32
	_ = v1150
	var v1153 int32
	_ = v1153
	var v1156 int32
	_ = v1156
	var v1157 int32
	_ = v1157
	var v1160 int32
	_ = v1160
	var v1161 int32
	_ = v1161
	var v1162 int32
	_ = v1162
	var v1163 int32
	_ = v1163
	var v1167 int32
	_ = v1167
	var v1170 int32
	_ = v1170
	var v1172 int32
	_ = v1172
	var v1173 int32
	_ = v1173
	var v1174 int32
	_ = v1174
	var v1176 int32
	_ = v1176
	var v1180 int32
	_ = v1180
	var v1182 int32
	_ = v1182
	var v1183 int32
	_ = v1183
	var v1184 int32
	_ = v1184
	var v1186 int32
	_ = v1186
	var v1187 int32
	_ = v1187
	var v1188 int32
	_ = v1188
	var v1189 int32
	_ = v1189
	var v1194 int32
	_ = v1194
	var v1195 int32
	_ = v1195
	var v1198 int32
	_ = v1198
	var v1199 int32
	_ = v1199
	var v1202 int32
	_ = v1202
	var v1203 int32
	_ = v1203
	var v1204 int32
	_ = v1204
	var v1206 int32
	_ = v1206
	var v1208 int32
	_ = v1208
	var v1212 int32
	_ = v1212
	var v1213 int32
	_ = v1213
	var v1215 int32
	_ = v1215
	var v1217 int32
	_ = v1217
	var v1218 int32
	_ = v1218
	var v1221 int32
	_ = v1221
	var v1228 int32
	_ = v1228
	var v1231 int32
	_ = v1231
	var v1234 int32
	_ = v1234
	var v1237 int32
	_ = v1237
	var v1238 int32
	_ = v1238
	var v1241 int32
	_ = v1241
	var v1242 int32
	_ = v1242
	var v1243 int32
	_ = v1243
	var v1244 int32
	_ = v1244
	var v1248 int32
	_ = v1248
	var v1254 int32
	_ = v1254
	var v1257 int32
	_ = v1257
	var v1260 int32
	_ = v1260
	var v1261 int32
	_ = v1261
	var v1264 int32
	_ = v1264
	var v1267 int32
	_ = v1267
	var v1268 int32
	_ = v1268
	var v1273 int32
	_ = v1273
	var v1284 int32
	_ = v1284
	var v1294 int32
	_ = v1294
	v9 = m.G0
	v11 = v9 - int32(64)
	m.G0 = v11
	v14 = *(*int32)(unsafe.Add(mBase, _consts[768]))
	if v14 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v93 = F_dictCreate(m, int32(_a1446))
	mBase = m.M
	v94 = m.ExcPending
	if v94 != 0 {
		goto L3
	} else {
		goto L39
	}
L2:
	;
	v17 = F_dictCreate(m, int32(_a1446))
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	return
L4:
	;
	*(*int32)(unsafe.Add(mBase, _consts[768])) = v17
	v21 = F_sdsnew(m, int32(_a1447))
	mBase = m.M
	v22 = m.ExcPending
	if v22 != 0 {
		goto L3
	} else {
		goto L6
	}
L5:
	;
	v31 = F_sdsnew(m, int32(_a1448))
	mBase = m.M
	v32 = m.ExcPending
	if v32 != 0 {
		goto L3
	} else {
		goto L11
	}
L6:
	;
	v24 = F_dictAdd(m, v17, v21, int32(0))
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		goto L3
	} else {
		goto L7
	}
L7:
	;
	if v24 != int32(1) {
		goto L5
	} else {
		goto L8
	}
L8:
	;
	F_sdsfree(m, v21)
	mBase = m.M
	v29 = m.ExcPending
	if v29 != 0 {
		goto L3
	} else {
		goto L9
	}
L9:
	;
	goto L5
L10:
	;
	v41 = F_sdsnew(m, int32(_a1449))
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L3
	} else {
		goto L16
	}
L11:
	;
	v34 = F_dictAdd(m, v17, v31, int32(0))
	mBase = m.M
	v35 = m.ExcPending
	if v35 != 0 {
		goto L3
	} else {
		goto L12
	}
L12:
	;
	if v34 != int32(1) {
		goto L10
	} else {
		goto L13
	}
L13:
	;
	F_sdsfree(m, v31)
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L3
	} else {
		goto L14
	}
L14:
	;
	goto L10
L15:
	;
	v51 = F_sdsnew(m, int32(_a1450))
	mBase = m.M
	v52 = m.ExcPending
	if v52 != 0 {
		goto L3
	} else {
		goto L21
	}
L16:
	;
	v44 = F_dictAdd(m, v17, v41, int32(0))
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L3
	} else {
		goto L17
	}
L17:
	;
	if v44 != int32(1) {
		goto L15
	} else {
		goto L18
	}
L18:
	;
	F_sdsfree(m, v41)
	mBase = m.M
	v49 = m.ExcPending
	if v49 != 0 {
		goto L3
	} else {
		goto L19
	}
L19:
	;
	goto L15
L20:
	;
	v61 = F_sdsnew(m, int32(_a1451))
	mBase = m.M
	v62 = m.ExcPending
	if v62 != 0 {
		goto L3
	} else {
		goto L26
	}
L21:
	;
	v54 = F_dictAdd(m, v17, v51, int32(0))
	mBase = m.M
	v55 = m.ExcPending
	if v55 != 0 {
		goto L3
	} else {
		goto L22
	}
L22:
	;
	if v54 != int32(1) {
		goto L20
	} else {
		goto L23
	}
L23:
	;
	F_sdsfree(m, v51)
	mBase = m.M
	v59 = m.ExcPending
	if v59 != 0 {
		goto L3
	} else {
		goto L24
	}
L24:
	;
	goto L20
L25:
	;
	v71 = F_sdsnew(m, int32(_a1452))
	mBase = m.M
	v72 = m.ExcPending
	if v72 != 0 {
		goto L3
	} else {
		goto L31
	}
L26:
	;
	v64 = F_dictAdd(m, v17, v61, int32(0))
	mBase = m.M
	v65 = m.ExcPending
	if v65 != 0 {
		goto L3
	} else {
		goto L27
	}
L27:
	;
	if v64 != int32(1) {
		goto L25
	} else {
		goto L28
	}
L28:
	;
	F_sdsfree(m, v61)
	mBase = m.M
	v69 = m.ExcPending
	if v69 != 0 {
		goto L3
	} else {
		goto L29
	}
L29:
	;
	goto L25
L30:
	;
	v81 = F_sdsnew(m, int32(_a1453))
	mBase = m.M
	v82 = m.ExcPending
	if v82 != 0 {
		goto L3
	} else {
		goto L35
	}
L31:
	;
	v74 = F_dictAdd(m, v17, v71, int32(0))
	mBase = m.M
	v75 = m.ExcPending
	if v75 != 0 {
		goto L3
	} else {
		goto L32
	}
L32:
	;
	if v74 != int32(1) {
		goto L30
	} else {
		goto L33
	}
L33:
	;
	F_sdsfree(m, v71)
	mBase = m.M
	v79 = m.ExcPending
	if v79 != 0 {
		goto L3
	} else {
		goto L34
	}
L34:
	;
	goto L30
L35:
	;
	v84 = F_dictAdd(m, v17, v81, int32(0))
	mBase = m.M
	v85 = m.ExcPending
	if v85 != 0 {
		goto L3
	} else {
		goto L36
	}
L36:
	;
	if v84 != int32(1) {
		goto L1
	} else {
		goto L37
	}
L37:
	;
	F_sdsfree(m, v81)
	mBase = m.M
	v89 = m.ExcPending
	if v89 != 0 {
		goto L3
	} else {
		goto L38
	}
L38:
	;
	goto L1
L39:
	;
	v97 = F_sentinelValidateArgs(m, l0, int32(3), int32(_a1454))
	mBase = m.M
	v98 = m.ExcPending
	if v98 != 0 {
		goto L3
	} else {
		goto L41
	}
L40:
	;
	F_dictRelease(m, v93)
	mBase = m.M
	v1294 = m.ExcPending
	if v1294 != 0 {
		goto L3
	} else {
		goto L428
	}
L41:
	;
	if v97 == int32(-1) {
		goto L40
	} else {
		goto L42
	}
L42:
	;
	v101 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v101 < int32(4) {
		goto L43
	} else {
		goto L44
	}
L43:
	;
	F_sentinelFlushConfigAndReply(m, l0)
	mBase = m.M
	v1284 = m.ExcPending
	if v1284 != 0 {
		goto L3
	} else {
		goto L427
	}
L44:
	;
	v108 = int32(3)
	goto L47
L45:
	;
	if v556 < int32(4) {
		goto L43
	} else {
		goto L205
	}
L46:
	;
	v564 = F_objectGetVal(m, v156)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v11)+36)) = v118
	*(*int32)(unsafe.Add(mBase, uint32(v11)+32)) = v564
	F_addReplyErrorFormat(m, l0, int32(_a1455), v11+int32(32))
	mBase = m.M
	v571 = m.ExcPending
	if v571 != 0 {
		goto L3
	} else {
		goto L204
	}
L47:
	;
	v113 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v117 = *(*int32)(unsafe.Add(mBase, uint32(v113+v108<<(uint(int32(2))%32))))
	v118 = F_objectGetVal(m, v117)
	mBase = m.M
	v120 = *(*int32)(unsafe.Add(mBase, _consts[768]))
	v121 = F_dictFind(m, v120, v118)
	mBase = m.M
	v122 = m.ExcPending
	if v122 != 0 {
		goto L3
	} else {
		goto L50
	}
L48:
	;
	F__serverAssert(m, int32(_a1456), int32(_a1333), int32(3155))
	mBase = m.M
	v562 = m.ExcPending
	if v562 != 0 {
		goto L3
	} else {
		goto L203
	}
L49:
	;
	v127 = F_dictFind(m, v93, v118)
	mBase = m.M
	v128 = m.ExcPending
	if v128 != 0 {
		goto L3
	} else {
		goto L54
	}
L50:
	;
	if v121 != 0 {
		goto L49
	} else {
		goto L51
	}
L51:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11))) = v118
	F_addReplyErrorFormat(m, l0, int32(_a1457), v11)
	mBase = m.M
	v126 = m.ExcPending
	if v126 != 0 {
		goto L3
	} else {
		goto L52
	}
L52:
	;
	goto L40
L53:
	;
	v137 = F_sdsnew(m, v118)
	mBase = m.M
	v138 = m.ExcPending
	if v138 != 0 {
		goto L3
	} else {
		goto L58
	}
L54:
	;
	if v127 == int32(0) {
		goto L53
	} else {
		goto L55
	}
L55:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+48)) = v118
	F_addReplyErrorFormat(m, l0, int32(_a1458), v11+int32(48))
	mBase = m.M
	v136 = m.ExcPending
	if v136 != 0 {
		goto L3
	} else {
		goto L56
	}
L56:
	;
	goto L40
L57:
	;
	goto L48
L58:
	;
	v140 = F_dictAdd(m, v93, v137, int32(0))
	mBase = m.M
	v141 = m.ExcPending
	if v141 != 0 {
		goto L3
	} else {
		goto L59
	}
L59:
	;
	if v140 != 0 {
		goto L57
	} else {
		goto L60
	}
L60:
	;
	v143 = v108 + int32(1)
	v144 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v143 != v144 {
		goto L61
	} else {
		goto L62
	}
L61:
	;
	v152 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v156 = *(*int32)(unsafe.Add(mBase, uint32(v152+v143<<(uint(int32(2))%32))))
	v157 = int32(_a1450)
	v160 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v118))))
	if v160 != 0 {
		goto L68
	} else {
		goto L69
	}
L62:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = v118
	F_addReplyErrorFormat(m, l0, int32(_a1459), v11+int32(16))
	mBase = m.M
	v151 = m.ExcPending
	if v151 != 0 {
		goto L3
	} else {
		goto L63
	}
L63:
	;
	goto L40
L64:
	;
	v555 = v108 + int32(2)
	v556 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v556 <= v555 {
		goto L45
	} else {
		goto L202
	}
L65:
	;
	v208 = int32(_a1452)
	v211 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v118))))
	if v211 != 0 {
		goto L89
	} else {
		goto L90
	}
L66:
	;
	if v192-v194 != 0 {
		goto L65
	} else {
		goto L78
	}
L67:
	;
	v192 = F_tolower(m, v188)
	mBase = m.M
	v193 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v189))))
	v194 = F_tolower(m, v193)
	mBase = m.M
	goto L66
L68:
	;
	v162 = v118
	v163 = v157
	v164 = v160
	goto L71
L69:
	;
	v188 = int32(0)
	v189 = v157
	goto L67
L70:
	;
	v188 = v185 & int32(255)
	v189 = v184
	goto L67
L71:
	;
	v166 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v163))))
	if v166 == int32(0) {
		v184 = v163
		v185 = v164
		goto L70
	} else {
		goto L73
	}
L72:
	;
	v184 = v178
	v185 = int32(0)
	goto L70
L73:
	;
	v170 = v164 & int32(255)
	if v170 == v166 {
		goto L74
	} else {
		goto L75
	}
L74:
	;
	v177 = int32(1)
	v178 = v163 + v177
	v179 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v162)+1)))
	if v179 != 0 {
		v162 = v162 + v177
		v163 = v178
		v164 = v179
		goto L71
	} else {
		goto L77
	}
L75:
	;
	v172 = F_tolower(m, v170)
	mBase = m.M
	v173 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v163))))
	v174 = F_tolower(m, v173)
	mBase = m.M
	if v172 == v174 {
		goto L74
	} else {
		goto L76
	}
L76:
	;
	v176 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v162))))
	v184 = v163
	v185 = v176
	goto L70
L77:
	;
	goto L72
L78:
	;
	v196 = F_objectGetVal(m, v156)
	mBase = m.M
	v198 = F_strcasecmp(m, v196, int32(_a483))
	mBase = m.M
	if v198 != 0 {
		goto L80
	} else {
		goto L81
	}
L79:
	;
	if v205 == int32(-1) {
		goto L46
	} else {
		goto L85
	}
L80:
	;
	v203 = F_strcasecmp(m, v196, int32(_a484))
	mBase = m.M
	if v203 != 0 {
		goto L82
	} else {
		goto L83
	}
L81:
	;
	v205 = int32(1)
	goto L79
L82:
	;
	v204 = int32(-1)
	goto L84
L83:
	;
	v204 = int32(0)
	goto L84
L84:
	;
	v205 = v204
	goto L79
L85:
	;
	goto L64
L86:
	;
	v259 = int32(_a1451)
	v262 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v118))))
	if v262 != 0 {
		goto L110
	} else {
		goto L111
	}
L87:
	;
	if v243-v245 != 0 {
		goto L86
	} else {
		goto L99
	}
L88:
	;
	v243 = F_tolower(m, v239)
	mBase = m.M
	v244 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v240))))
	v245 = F_tolower(m, v244)
	mBase = m.M
	goto L87
L89:
	;
	v213 = v118
	v214 = v208
	v215 = v211
	goto L92
L90:
	;
	v239 = int32(0)
	v240 = v208
	goto L88
L91:
	;
	v239 = v236 & int32(255)
	v240 = v235
	goto L88
L92:
	;
	v217 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v214))))
	if v217 == int32(0) {
		v235 = v214
		v236 = v215
		goto L91
	} else {
		goto L94
	}
L93:
	;
	v235 = v229
	v236 = int32(0)
	goto L91
L94:
	;
	v221 = v215 & int32(255)
	if v221 == v217 {
		goto L95
	} else {
		goto L96
	}
L95:
	;
	v228 = int32(1)
	v229 = v214 + v228
	v230 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v213)+1)))
	if v230 != 0 {
		v213 = v213 + v228
		v214 = v229
		v215 = v230
		goto L92
	} else {
		goto L98
	}
L96:
	;
	v223 = F_tolower(m, v221)
	mBase = m.M
	v224 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v214))))
	v225 = F_tolower(m, v224)
	mBase = m.M
	if v223 == v225 {
		goto L95
	} else {
		goto L97
	}
L97:
	;
	v227 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v213))))
	v235 = v214
	v236 = v227
	goto L91
L98:
	;
	goto L93
L99:
	;
	v247 = F_objectGetVal(m, v156)
	mBase = m.M
	v249 = F_strcasecmp(m, v247, int32(_a483))
	mBase = m.M
	if v249 != 0 {
		goto L101
	} else {
		goto L102
	}
L100:
	;
	if v256 == int32(-1) {
		goto L46
	} else {
		goto L106
	}
L101:
	;
	v254 = F_strcasecmp(m, v247, int32(_a484))
	mBase = m.M
	if v254 != 0 {
		goto L103
	} else {
		goto L104
	}
L102:
	;
	v256 = int32(1)
	goto L100
L103:
	;
	v255 = int32(-1)
	goto L105
L104:
	;
	v255 = int32(0)
	goto L105
L105:
	;
	v256 = v255
	goto L100
L106:
	;
	goto L64
L107:
	;
	v307 = int32(_a1453)
	v310 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v118))))
	if v310 != 0 {
		goto L126
	} else {
		goto L127
	}
L108:
	;
	if v294-v296 != 0 {
		goto L107
	} else {
		goto L120
	}
L109:
	;
	v294 = F_tolower(m, v290)
	mBase = m.M
	v295 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v291))))
	v296 = F_tolower(m, v295)
	mBase = m.M
	goto L108
L110:
	;
	v264 = v118
	v265 = v259
	v266 = v262
	goto L113
L111:
	;
	v290 = int32(0)
	v291 = v259
	goto L109
L112:
	;
	v290 = v287 & int32(255)
	v291 = v286
	goto L109
L113:
	;
	v268 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v265))))
	if v268 == int32(0) {
		v286 = v265
		v287 = v266
		goto L112
	} else {
		goto L115
	}
L114:
	;
	v286 = v280
	v287 = int32(0)
	goto L112
L115:
	;
	v272 = v266 & int32(255)
	if v272 == v268 {
		goto L116
	} else {
		goto L117
	}
L116:
	;
	v279 = int32(1)
	v280 = v265 + v279
	v281 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v264)+1)))
	if v281 != 0 {
		v264 = v264 + v279
		v265 = v280
		v266 = v281
		goto L113
	} else {
		goto L119
	}
L117:
	;
	v274 = F_tolower(m, v272)
	mBase = m.M
	v275 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v265))))
	v276 = F_tolower(m, v275)
	mBase = m.M
	if v274 == v276 {
		goto L116
	} else {
		goto L118
	}
L118:
	;
	v278 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v264))))
	v286 = v265
	v287 = v278
	goto L112
L119:
	;
	goto L114
L120:
	;
	v300 = F_getLongLongFromObject(m, v156, v11+int32(56))
	mBase = m.M
	v301 = m.ExcPending
	if v301 != 0 {
		goto L3
	} else {
		goto L121
	}
L121:
	;
	if v300 == int32(-1) {
		goto L46
	} else {
		goto L122
	}
L122:
	;
	v304 = *(*int64)(unsafe.Add(mBase, uint32(v11)+56))
	if base.Ui64(v304) <= base.Ui64(int64(65535)) {
		goto L64
	} else {
		goto L123
	}
L123:
	;
	goto L46
L124:
	;
	if v342-v344 != 0 {
		goto L64
	} else {
		goto L136
	}
L125:
	;
	v342 = F_tolower(m, v338)
	mBase = m.M
	v343 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v339))))
	v344 = F_tolower(m, v343)
	mBase = m.M
	goto L124
L126:
	;
	v312 = v118
	v313 = v307
	v314 = v310
	goto L129
L127:
	;
	v338 = int32(0)
	v339 = v307
	goto L125
L128:
	;
	v338 = v335 & int32(255)
	v339 = v334
	goto L125
L129:
	;
	v316 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v313))))
	if v316 == int32(0) {
		v334 = v313
		v335 = v314
		goto L128
	} else {
		goto L131
	}
L130:
	;
	v334 = v328
	v335 = int32(0)
	goto L128
L131:
	;
	v320 = v314 & int32(255)
	if v320 == v316 {
		goto L132
	} else {
		goto L133
	}
L132:
	;
	v327 = int32(1)
	v328 = v313 + v327
	v329 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v312)+1)))
	if v329 != 0 {
		v312 = v312 + v327
		v313 = v328
		v314 = v329
		goto L129
	} else {
		goto L135
	}
L133:
	;
	v322 = F_tolower(m, v320)
	mBase = m.M
	v323 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v313))))
	v324 = F_tolower(m, v323)
	mBase = m.M
	if v322 == v324 {
		goto L132
	} else {
		goto L134
	}
L134:
	;
	v326 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v312))))
	v334 = v313
	v335 = v326
	goto L128
L135:
	;
	goto L130
L136:
	;
	v346 = F_objectGetVal(m, v156)
	mBase = m.M
	v347 = int32(_a774)
	v350 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v346))))
	if v350 != 0 {
		goto L139
	} else {
		goto L140
	}
L137:
	;
	if v382-v384 == int32(0) {
		goto L64
	} else {
		goto L149
	}
L138:
	;
	v382 = F_tolower(m, v378)
	mBase = m.M
	v383 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v379))))
	v384 = F_tolower(m, v383)
	mBase = m.M
	goto L137
L139:
	;
	v352 = v346
	v353 = v347
	v354 = v350
	goto L142
L140:
	;
	v378 = int32(0)
	v379 = v347
	goto L138
L141:
	;
	v378 = v375 & int32(255)
	v379 = v374
	goto L138
L142:
	;
	v356 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v353))))
	if v356 == int32(0) {
		v374 = v353
		v375 = v354
		goto L141
	} else {
		goto L144
	}
L143:
	;
	v374 = v368
	v375 = int32(0)
	goto L141
L144:
	;
	v360 = v354 & int32(255)
	if v360 == v356 {
		goto L145
	} else {
		goto L146
	}
L145:
	;
	v367 = int32(1)
	v368 = v353 + v367
	v369 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v352)+1)))
	if v369 != 0 {
		v352 = v352 + v367
		v353 = v368
		v354 = v369
		goto L142
	} else {
		goto L148
	}
L146:
	;
	v362 = F_tolower(m, v360)
	mBase = m.M
	v363 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v353))))
	v364 = F_tolower(m, v363)
	mBase = m.M
	if v362 == v364 {
		goto L145
	} else {
		goto L147
	}
L147:
	;
	v366 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v352))))
	v374 = v353
	v375 = v366
	goto L141
L148:
	;
	goto L143
L149:
	;
	v388 = F_objectGetVal(m, v156)
	mBase = m.M
	v389 = int32(_a1460)
	v392 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v388))))
	if v392 != 0 {
		goto L152
	} else {
		goto L153
	}
L150:
	;
	if v424-v426 == int32(0) {
		goto L64
	} else {
		goto L162
	}
L151:
	;
	v424 = F_tolower(m, v420)
	mBase = m.M
	v425 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v421))))
	v426 = F_tolower(m, v425)
	mBase = m.M
	goto L150
L152:
	;
	v394 = v388
	v395 = v389
	v396 = v392
	goto L155
L153:
	;
	v420 = int32(0)
	v421 = v389
	goto L151
L154:
	;
	v420 = v417 & int32(255)
	v421 = v416
	goto L151
L155:
	;
	v398 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v395))))
	if v398 == int32(0) {
		v416 = v395
		v417 = v396
		goto L154
	} else {
		goto L157
	}
L156:
	;
	v416 = v410
	v417 = int32(0)
	goto L154
L157:
	;
	v402 = v396 & int32(255)
	if v402 == v398 {
		goto L158
	} else {
		goto L159
	}
L158:
	;
	v409 = int32(1)
	v410 = v395 + v409
	v411 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v394)+1)))
	if v411 != 0 {
		v394 = v394 + v409
		v395 = v410
		v396 = v411
		goto L155
	} else {
		goto L161
	}
L159:
	;
	v404 = F_tolower(m, v402)
	mBase = m.M
	v405 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v395))))
	v406 = F_tolower(m, v405)
	mBase = m.M
	if v404 == v406 {
		goto L158
	} else {
		goto L160
	}
L160:
	;
	v408 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v394))))
	v416 = v395
	v417 = v408
	goto L154
L161:
	;
	goto L156
L162:
	;
	v430 = F_objectGetVal(m, v156)
	mBase = m.M
	v431 = int32(_a1461)
	v434 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v430))))
	if v434 != 0 {
		goto L165
	} else {
		goto L166
	}
L163:
	;
	if v466-v468 == int32(0) {
		goto L64
	} else {
		goto L175
	}
L164:
	;
	v466 = F_tolower(m, v462)
	mBase = m.M
	v467 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v463))))
	v468 = F_tolower(m, v467)
	mBase = m.M
	goto L163
L165:
	;
	v436 = v430
	v437 = v431
	v438 = v434
	goto L168
L166:
	;
	v462 = int32(0)
	v463 = v431
	goto L164
L167:
	;
	v462 = v459 & int32(255)
	v463 = v458
	goto L164
L168:
	;
	v440 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v437))))
	if v440 == int32(0) {
		v458 = v437
		v459 = v438
		goto L167
	} else {
		goto L170
	}
L169:
	;
	v458 = v452
	v459 = int32(0)
	goto L167
L170:
	;
	v444 = v438 & int32(255)
	if v444 == v440 {
		goto L171
	} else {
		goto L172
	}
L171:
	;
	v451 = int32(1)
	v452 = v437 + v451
	v453 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v436)+1)))
	if v453 != 0 {
		v436 = v436 + v451
		v437 = v452
		v438 = v453
		goto L168
	} else {
		goto L174
	}
L172:
	;
	v446 = F_tolower(m, v444)
	mBase = m.M
	v447 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v437))))
	v448 = F_tolower(m, v447)
	mBase = m.M
	if v446 == v448 {
		goto L171
	} else {
		goto L173
	}
L173:
	;
	v450 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v436))))
	v458 = v437
	v459 = v450
	goto L167
L174:
	;
	goto L169
L175:
	;
	v472 = F_objectGetVal(m, v156)
	mBase = m.M
	v473 = int32(_a1462)
	v476 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v472))))
	if v476 != 0 {
		goto L178
	} else {
		goto L179
	}
L176:
	;
	if v508-v510 == int32(0) {
		goto L64
	} else {
		goto L188
	}
L177:
	;
	v508 = F_tolower(m, v504)
	mBase = m.M
	v509 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v505))))
	v510 = F_tolower(m, v509)
	mBase = m.M
	goto L176
L178:
	;
	v478 = v472
	v479 = v473
	v480 = v476
	goto L181
L179:
	;
	v504 = int32(0)
	v505 = v473
	goto L177
L180:
	;
	v504 = v501 & int32(255)
	v505 = v500
	goto L177
L181:
	;
	v482 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v479))))
	if v482 == int32(0) {
		v500 = v479
		v501 = v480
		goto L180
	} else {
		goto L183
	}
L182:
	;
	v500 = v494
	v501 = int32(0)
	goto L180
L183:
	;
	v486 = v480 & int32(255)
	if v486 == v482 {
		goto L184
	} else {
		goto L185
	}
L184:
	;
	v493 = int32(1)
	v494 = v479 + v493
	v495 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v478)+1)))
	if v495 != 0 {
		v478 = v478 + v493
		v479 = v494
		v480 = v495
		goto L181
	} else {
		goto L187
	}
L185:
	;
	v488 = F_tolower(m, v486)
	mBase = m.M
	v489 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v479))))
	v490 = F_tolower(m, v489)
	mBase = m.M
	if v488 == v490 {
		goto L184
	} else {
		goto L186
	}
L186:
	;
	v492 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v478))))
	v500 = v479
	v501 = v492
	goto L180
L187:
	;
	goto L182
L188:
	;
	v514 = F_objectGetVal(m, v156)
	mBase = m.M
	v515 = int32(_a1463)
	v518 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v514))))
	if v518 != 0 {
		goto L191
	} else {
		goto L192
	}
L189:
	;
	if v550-v552 != 0 {
		goto L46
	} else {
		goto L201
	}
L190:
	;
	v550 = F_tolower(m, v546)
	mBase = m.M
	v551 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v547))))
	v552 = F_tolower(m, v551)
	mBase = m.M
	goto L189
L191:
	;
	v520 = v514
	v521 = v515
	v522 = v518
	goto L194
L192:
	;
	v546 = int32(0)
	v547 = v515
	goto L190
L193:
	;
	v546 = v543 & int32(255)
	v547 = v542
	goto L190
L194:
	;
	v524 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v521))))
	if v524 == int32(0) {
		v542 = v521
		v543 = v522
		goto L193
	} else {
		goto L196
	}
L195:
	;
	v542 = v536
	v543 = int32(0)
	goto L193
L196:
	;
	v528 = v522 & int32(255)
	if v528 == v524 {
		goto L197
	} else {
		goto L198
	}
L197:
	;
	v535 = int32(1)
	v536 = v521 + v535
	v537 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v520)+1)))
	if v537 != 0 {
		v520 = v520 + v535
		v521 = v536
		v522 = v537
		goto L194
	} else {
		goto L200
	}
L198:
	;
	v530 = F_tolower(m, v528)
	mBase = m.M
	v531 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v521))))
	v532 = F_tolower(m, v531)
	mBase = m.M
	if v530 == v532 {
		goto L197
	} else {
		goto L199
	}
L199:
	;
	v534 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v520))))
	v542 = v521
	v543 = v534
	goto L193
L200:
	;
	goto L195
L201:
	;
	goto L64
L202:
	;
	v108 = v555
	goto L47
L203:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L204:
	;
	goto L40
L205:
	;
	v578 = v556
	v579 = int32(3)
	v582 = int32(0)
	goto L207
L206:
	;
	F__serverAssert(m, int32(_a195), int32(_a1333), int32(3222))
	mBase = m.M
	v1273 = m.ExcPending
	if v1273 != 0 {
		goto L3
	} else {
		goto L426
	}
L207:
	;
	v584 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v588 = *(*int32)(unsafe.Add(mBase, uint32(v584+v579<<(uint(int32(2))%32))))
	v589 = F_objectGetVal(m, v588)
	mBase = m.M
	v590 = int32(_a1453)
	v593 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v589))))
	if v593 != 0 {
		goto L211
	} else {
		goto L212
	}
L208:
	;
	F_sentinelFlushConfigAndReply(m, l0)
	mBase = m.M
	v1264 = m.ExcPending
	if v1264 != 0 {
		goto L3
	} else {
		goto L423
	}
L209:
	;
	v631 = base.B2i32(int32(-2) < v579-v578)
	if int32(-2) < v579-v578 {
		goto L224
	} else {
		goto L225
	}
L210:
	;
	v625 = F_tolower(m, v621)
	mBase = m.M
	v626 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v622))))
	v627 = F_tolower(m, v626)
	mBase = m.M
	goto L209
L211:
	;
	v595 = v589
	v596 = v590
	v597 = v593
	goto L214
L212:
	;
	v621 = int32(0)
	v622 = v590
	goto L210
L213:
	;
	v621 = v618 & int32(255)
	v622 = v617
	goto L210
L214:
	;
	v599 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v596))))
	if v599 == int32(0) {
		v617 = v596
		v618 = v597
		goto L213
	} else {
		goto L216
	}
L215:
	;
	v617 = v611
	v618 = int32(0)
	goto L213
L216:
	;
	v603 = v597 & int32(255)
	if v603 == v599 {
		goto L217
	} else {
		goto L218
	}
L217:
	;
	v610 = int32(1)
	v611 = v596 + v610
	v612 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v595)+1)))
	if v612 != 0 {
		v595 = v595 + v610
		v596 = v611
		v597 = v612
		goto L214
	} else {
		goto L220
	}
L218:
	;
	v605 = F_tolower(m, v603)
	mBase = m.M
	v606 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v596))))
	v607 = F_tolower(m, v606)
	mBase = m.M
	if v605 == v607 {
		goto L217
	} else {
		goto L219
	}
L219:
	;
	v609 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v595))))
	v617 = v596
	v618 = v609
	goto L213
L220:
	;
	goto L215
L221:
	;
	v1260 = v1254 + int32(1)
	v1261 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v1260 < v1261 {
		v578 = v1261
		v579 = v1260
		v582 = v1257
		goto L207
	} else {
		goto L422
	}
L222:
	;
	v916 = int32(_a1452)
	v919 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v589))))
	if v919 != 0 {
		goto L318
	} else {
		goto L319
	}
L223:
	;
	v694 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v696 = v579 + int32(1)
	v700 = *(*int32)(unsafe.Add(mBase, uint32(v694+v696<<(uint(int32(2))%32))))
	v701 = F_objectGetVal(m, v700)
	mBase = m.M
	v702 = int32(_a774)
	v705 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v701))))
	if v705 != 0 {
		goto L250
	} else {
		goto L251
	}
L224:
	;
	v634 = int32(_a1450)
	v637 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v589))))
	if v637 != 0 {
		goto L229
	} else {
		goto L230
	}
L225:
	;
	if v625-v627 == int32(0) {
		goto L223
	} else {
		goto L226
	}
L226:
	;
	goto L224
L227:
	;
	if int32(-2) < v579-v578 {
		goto L222
	} else {
		goto L239
	}
L228:
	;
	v669 = F_tolower(m, v665)
	mBase = m.M
	v670 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v666))))
	v671 = F_tolower(m, v670)
	mBase = m.M
	goto L227
L229:
	;
	v639 = v589
	v640 = v634
	v641 = v637
	goto L232
L230:
	;
	v665 = int32(0)
	v666 = v634
	goto L228
L231:
	;
	v665 = v662 & int32(255)
	v666 = v661
	goto L228
L232:
	;
	v643 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v640))))
	if v643 == int32(0) {
		v661 = v640
		v662 = v641
		goto L231
	} else {
		goto L234
	}
L233:
	;
	v661 = v655
	v662 = int32(0)
	goto L231
L234:
	;
	v647 = v641 & int32(255)
	if v647 == v643 {
		goto L235
	} else {
		goto L236
	}
L235:
	;
	v654 = int32(1)
	v655 = v640 + v654
	v656 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v639)+1)))
	if v656 != 0 {
		v639 = v639 + v654
		v640 = v655
		v641 = v656
		goto L232
	} else {
		goto L238
	}
L236:
	;
	v649 = F_tolower(m, v647)
	mBase = m.M
	v650 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v640))))
	v651 = F_tolower(m, v650)
	mBase = m.M
	if v649 == v651 {
		goto L235
	} else {
		goto L237
	}
L237:
	;
	v653 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v639))))
	v661 = v640
	v662 = v653
	goto L231
L238:
	;
	goto L233
L239:
	;
	if v669-v671 != 0 {
		goto L222
	} else {
		goto L240
	}
L240:
	;
	v674 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v676 = v579 + int32(1)
	v680 = *(*int32)(unsafe.Add(mBase, uint32(v674+v676<<(uint(int32(2))%32))))
	v681 = F_objectGetVal(m, v680)
	mBase = m.M
	v683 = F_strcasecmp(m, v681, int32(_a483))
	mBase = m.M
	if v683 != 0 {
		goto L242
	} else {
		goto L243
	}
L241:
	;
	*(*int32)(unsafe.Add(mBase, _consts[729])) = v690
	*(*int64)(unsafe.Add(mBase, uint32(v11)+56)) = base.I64_extend_i32_s(v690)
	v1254 = v676
	v1257 = v582
	goto L221
L242:
	;
	v688 = F_strcasecmp(m, v681, int32(_a484))
	mBase = m.M
	if v688 != 0 {
		goto L244
	} else {
		goto L245
	}
L243:
	;
	v690 = int32(1)
	goto L241
L244:
	;
	v689 = int32(-1)
	goto L246
L245:
	;
	v689 = int32(0)
	goto L246
L246:
	;
	v690 = v689
	goto L241
L247:
	;
	v744 = F_objectGetVal(m, v700)
	mBase = m.M
	v745 = int32(_a1460)
	v748 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v744))))
	if v748 != 0 {
		goto L264
	} else {
		goto L265
	}
L248:
	;
	if v737-v739 != 0 {
		goto L247
	} else {
		goto L260
	}
L249:
	;
	v737 = F_tolower(m, v733)
	mBase = m.M
	v738 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v734))))
	v739 = F_tolower(m, v738)
	mBase = m.M
	goto L248
L250:
	;
	v707 = v701
	v708 = v702
	v709 = v705
	goto L253
L251:
	;
	v733 = int32(0)
	v734 = v702
	goto L249
L252:
	;
	v733 = v730 & int32(255)
	v734 = v729
	goto L249
L253:
	;
	v711 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v708))))
	if v711 == int32(0) {
		v729 = v708
		v730 = v709
		goto L252
	} else {
		goto L255
	}
L254:
	;
	v729 = v723
	v730 = int32(0)
	goto L252
L255:
	;
	v715 = v709 & int32(255)
	if v715 == v711 {
		goto L256
	} else {
		goto L257
	}
L256:
	;
	v722 = int32(1)
	v723 = v708 + v722
	v724 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v707)+1)))
	if v724 != 0 {
		v707 = v707 + v722
		v708 = v723
		v709 = v724
		goto L253
	} else {
		goto L259
	}
L257:
	;
	v717 = F_tolower(m, v715)
	mBase = m.M
	v718 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v708))))
	v719 = F_tolower(m, v718)
	mBase = m.M
	if v717 == v719 {
		goto L256
	} else {
		goto L258
	}
L258:
	;
	v721 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v707))))
	v729 = v708
	v730 = v721
	goto L252
L259:
	;
	goto L254
L260:
	;
	*(*int32)(unsafe.Add(mBase, _consts[6])) = int32(0)
	v1254 = v696
	v1257 = v582
	goto L221
L261:
	;
	v787 = F_objectGetVal(m, v700)
	mBase = m.M
	v788 = int32(_a1461)
	v791 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v787))))
	if v791 != 0 {
		goto L278
	} else {
		goto L279
	}
L262:
	;
	if v780-v782 != 0 {
		goto L261
	} else {
		goto L274
	}
L263:
	;
	v780 = F_tolower(m, v776)
	mBase = m.M
	v781 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v777))))
	v782 = F_tolower(m, v781)
	mBase = m.M
	goto L262
L264:
	;
	v750 = v744
	v751 = v745
	v752 = v748
	goto L267
L265:
	;
	v776 = int32(0)
	v777 = v745
	goto L263
L266:
	;
	v776 = v773 & int32(255)
	v777 = v772
	goto L263
L267:
	;
	v754 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v751))))
	if v754 == int32(0) {
		v772 = v751
		v773 = v752
		goto L266
	} else {
		goto L269
	}
L268:
	;
	v772 = v766
	v773 = int32(0)
	goto L266
L269:
	;
	v758 = v752 & int32(255)
	if v758 == v754 {
		goto L270
	} else {
		goto L271
	}
L270:
	;
	v765 = int32(1)
	v766 = v751 + v765
	v767 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v750)+1)))
	if v767 != 0 {
		v750 = v750 + v765
		v751 = v766
		v752 = v767
		goto L267
	} else {
		goto L273
	}
L271:
	;
	v760 = F_tolower(m, v758)
	mBase = m.M
	v761 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v751))))
	v762 = F_tolower(m, v761)
	mBase = m.M
	if v760 == v762 {
		goto L270
	} else {
		goto L272
	}
L272:
	;
	v764 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v750))))
	v772 = v751
	v773 = v764
	goto L266
L273:
	;
	goto L268
L274:
	;
	*(*int32)(unsafe.Add(mBase, _consts[6])) = int32(1)
	v1254 = v696
	v1257 = v582
	goto L221
L275:
	;
	v830 = F_objectGetVal(m, v700)
	mBase = m.M
	v831 = int32(_a1462)
	v834 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v830))))
	if v834 != 0 {
		goto L292
	} else {
		goto L293
	}
L276:
	;
	if v823-v825 != 0 {
		goto L275
	} else {
		goto L288
	}
L277:
	;
	v823 = F_tolower(m, v819)
	mBase = m.M
	v824 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v820))))
	v825 = F_tolower(m, v824)
	mBase = m.M
	goto L276
L278:
	;
	v793 = v787
	v794 = v788
	v795 = v791
	goto L281
L279:
	;
	v819 = int32(0)
	v820 = v788
	goto L277
L280:
	;
	v819 = v816 & int32(255)
	v820 = v815
	goto L277
L281:
	;
	v797 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v794))))
	if v797 == int32(0) {
		v815 = v794
		v816 = v795
		goto L280
	} else {
		goto L283
	}
L282:
	;
	v815 = v809
	v816 = int32(0)
	goto L280
L283:
	;
	v801 = v795 & int32(255)
	if v801 == v797 {
		goto L284
	} else {
		goto L285
	}
L284:
	;
	v808 = int32(1)
	v809 = v794 + v808
	v810 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v793)+1)))
	if v810 != 0 {
		v793 = v793 + v808
		v794 = v809
		v795 = v810
		goto L281
	} else {
		goto L287
	}
L285:
	;
	v803 = F_tolower(m, v801)
	mBase = m.M
	v804 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v794))))
	v805 = F_tolower(m, v804)
	mBase = m.M
	if v803 == v805 {
		goto L284
	} else {
		goto L286
	}
L286:
	;
	v807 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v793))))
	v815 = v794
	v816 = v807
	goto L280
L287:
	;
	goto L282
L288:
	;
	*(*int32)(unsafe.Add(mBase, _consts[6])) = int32(2)
	v1254 = v696
	v1257 = v582
	goto L221
L289:
	;
	v873 = F_objectGetVal(m, v700)
	mBase = m.M
	v874 = int32(_a1463)
	v877 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v873))))
	if v877 != 0 {
		goto L305
	} else {
		goto L306
	}
L290:
	;
	if v866-v868 != 0 {
		goto L289
	} else {
		goto L302
	}
L291:
	;
	v866 = F_tolower(m, v862)
	mBase = m.M
	v867 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v863))))
	v868 = F_tolower(m, v867)
	mBase = m.M
	goto L290
L292:
	;
	v836 = v830
	v837 = v831
	v838 = v834
	goto L295
L293:
	;
	v862 = int32(0)
	v863 = v831
	goto L291
L294:
	;
	v862 = v859 & int32(255)
	v863 = v858
	goto L291
L295:
	;
	v840 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v837))))
	if v840 == int32(0) {
		v858 = v837
		v859 = v838
		goto L294
	} else {
		goto L297
	}
L296:
	;
	v858 = v852
	v859 = int32(0)
	goto L294
L297:
	;
	v844 = v838 & int32(255)
	if v844 == v840 {
		goto L298
	} else {
		goto L299
	}
L298:
	;
	v851 = int32(1)
	v852 = v837 + v851
	v853 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v836)+1)))
	if v853 != 0 {
		v836 = v836 + v851
		v837 = v852
		v838 = v853
		goto L295
	} else {
		goto L301
	}
L299:
	;
	v846 = F_tolower(m, v844)
	mBase = m.M
	v847 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v837))))
	v848 = F_tolower(m, v847)
	mBase = m.M
	if v846 == v848 {
		goto L298
	} else {
		goto L300
	}
L300:
	;
	v850 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v836))))
	v858 = v837
	v859 = v850
	goto L294
L301:
	;
	goto L296
L302:
	;
	*(*int32)(unsafe.Add(mBase, _consts[6])) = int32(3)
	v1254 = v696
	v1257 = v582
	goto L221
L303:
	;
	if v909-v911 != 0 {
		v1254 = v696
		v1257 = v582
		goto L221
	} else {
		goto L315
	}
L304:
	;
	v909 = F_tolower(m, v905)
	mBase = m.M
	v910 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v906))))
	v911 = F_tolower(m, v910)
	mBase = m.M
	goto L303
L305:
	;
	v879 = v873
	v880 = v874
	v881 = v877
	goto L308
L306:
	;
	v905 = int32(0)
	v906 = v874
	goto L304
L307:
	;
	v905 = v902 & int32(255)
	v906 = v901
	goto L304
L308:
	;
	v883 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v880))))
	if v883 == int32(0) {
		v901 = v880
		v902 = v881
		goto L307
	} else {
		goto L310
	}
L309:
	;
	v901 = v895
	v902 = int32(0)
	goto L307
L310:
	;
	v887 = v881 & int32(255)
	if v887 == v883 {
		goto L311
	} else {
		goto L312
	}
L311:
	;
	v894 = int32(1)
	v895 = v880 + v894
	v896 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v879)+1)))
	if v896 != 0 {
		v879 = v879 + v894
		v880 = v895
		v881 = v896
		goto L308
	} else {
		goto L314
	}
L312:
	;
	v889 = F_tolower(m, v887)
	mBase = m.M
	v890 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v880))))
	v891 = F_tolower(m, v890)
	mBase = m.M
	if v889 == v891 {
		goto L311
	} else {
		goto L313
	}
L313:
	;
	v893 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v879))))
	v901 = v880
	v902 = v893
	goto L307
L314:
	;
	goto L309
L315:
	;
	*(*int32)(unsafe.Add(mBase, _consts[6])) = int32(4)
	v1254 = v696
	v1257 = v582
	goto L221
L316:
	;
	if int32(-2) < v579-v578 {
		goto L328
	} else {
		goto L329
	}
L317:
	;
	v951 = F_tolower(m, v947)
	mBase = m.M
	v952 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v948))))
	v953 = F_tolower(m, v952)
	mBase = m.M
	goto L316
L318:
	;
	v921 = v589
	v922 = v916
	v923 = v919
	goto L321
L319:
	;
	v947 = int32(0)
	v948 = v916
	goto L317
L320:
	;
	v947 = v944 & int32(255)
	v948 = v943
	goto L317
L321:
	;
	v925 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v922))))
	if v925 == int32(0) {
		v943 = v922
		v944 = v923
		goto L320
	} else {
		goto L323
	}
L322:
	;
	v943 = v937
	v944 = int32(0)
	goto L320
L323:
	;
	v929 = v923 & int32(255)
	if v929 == v925 {
		goto L324
	} else {
		goto L325
	}
L324:
	;
	v936 = int32(1)
	v937 = v922 + v936
	v938 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v921)+1)))
	if v938 != 0 {
		v921 = v921 + v936
		v922 = v937
		v923 = v938
		goto L321
	} else {
		goto L327
	}
L325:
	;
	v931 = F_tolower(m, v929)
	mBase = m.M
	v932 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v922))))
	v933 = F_tolower(m, v932)
	mBase = m.M
	if v931 == v933 {
		goto L324
	} else {
		goto L326
	}
L326:
	;
	v935 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v921))))
	v943 = v922
	v944 = v935
	goto L320
L327:
	;
	goto L322
L328:
	;
	v976 = int32(_a1447)
	v979 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v589))))
	if v979 != 0 {
		goto L339
	} else {
		goto L340
	}
L329:
	;
	if v951-v953 != 0 {
		goto L328
	} else {
		goto L330
	}
L330:
	;
	v956 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v958 = v579 + int32(1)
	v962 = *(*int32)(unsafe.Add(mBase, uint32(v956+v958<<(uint(int32(2))%32))))
	v963 = F_objectGetVal(m, v962)
	mBase = m.M
	v965 = F_strcasecmp(m, v963, int32(_a483))
	mBase = m.M
	if v965 != 0 {
		goto L332
	} else {
		goto L333
	}
L331:
	;
	*(*int32)(unsafe.Add(mBase, _consts[734])) = v972
	*(*int64)(unsafe.Add(mBase, uint32(v11)+56)) = base.I64_extend_i32_s(v972)
	v1254 = v958
	v1257 = v582
	goto L221
L332:
	;
	v970 = F_strcasecmp(m, v963, int32(_a484))
	mBase = m.M
	if v970 != 0 {
		goto L334
	} else {
		goto L335
	}
L333:
	;
	v972 = int32(1)
	goto L331
L334:
	;
	v971 = int32(-1)
	goto L336
L335:
	;
	v971 = int32(0)
	goto L336
L336:
	;
	v972 = v971
	goto L331
L337:
	;
	if int32(-2) < v579-v578 {
		goto L349
	} else {
		goto L350
	}
L338:
	;
	v1011 = F_tolower(m, v1007)
	mBase = m.M
	v1012 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1008))))
	v1013 = F_tolower(m, v1012)
	mBase = m.M
	goto L337
L339:
	;
	v981 = v589
	v982 = v976
	v983 = v979
	goto L342
L340:
	;
	v1007 = int32(0)
	v1008 = v976
	goto L338
L341:
	;
	v1007 = v1004 & int32(255)
	v1008 = v1003
	goto L338
L342:
	;
	v985 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v982))))
	if v985 == int32(0) {
		v1003 = v982
		v1004 = v983
		goto L341
	} else {
		goto L344
	}
L343:
	;
	v1003 = v997
	v1004 = int32(0)
	goto L341
L344:
	;
	v989 = v983 & int32(255)
	if v989 == v985 {
		goto L345
	} else {
		goto L346
	}
L345:
	;
	v996 = int32(1)
	v997 = v982 + v996
	v998 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v981)+1)))
	if v998 != 0 {
		v981 = v981 + v996
		v982 = v997
		v983 = v998
		goto L342
	} else {
		goto L348
	}
L346:
	;
	v991 = F_tolower(m, v989)
	mBase = m.M
	v992 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v982))))
	v993 = F_tolower(m, v992)
	mBase = m.M
	if v991 == v993 {
		goto L345
	} else {
		goto L347
	}
L347:
	;
	v995 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v981))))
	v1003 = v982
	v1004 = v995
	goto L341
L348:
	;
	goto L343
L349:
	;
	v1033 = int32(_a1451)
	v1036 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v589))))
	if v1036 != 0 {
		goto L358
	} else {
		goto L359
	}
L350:
	;
	if v1011-v1013 != 0 {
		goto L349
	} else {
		goto L351
	}
L351:
	;
	v1015 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v1017 = v579 + int32(1)
	v1021 = *(*int32)(unsafe.Add(mBase, uint32(v1015+v1017<<(uint(int32(2))%32))))
	v1022 = int32(0)
	v1023 = *(*int32)(unsafe.Add(mBase, _consts[749]))
	if v1023 == v1022 {
		goto L352
	} else {
		goto L353
	}
L352:
	;
	v1029 = F_objectGetVal(m, v1021)
	mBase = m.M
	v1030 = F_sdsnew(m, v1029)
	mBase = m.M
	v1031 = m.ExcPending
	if v1031 != 0 {
		goto L3
	} else {
		goto L355
	}
L353:
	;
	F_sdsfree(m, v1023)
	mBase = m.M
	v1027 = m.ExcPending
	if v1027 != 0 {
		goto L3
	} else {
		goto L354
	}
L354:
	;
	goto L352
L355:
	;
	*(*int32)(unsafe.Add(mBase, _consts[749])) = v1030
	v1254 = v1017
	v1257 = v582
	goto L221
L356:
	;
	if int32(-2) < v579-v578 {
		goto L368
	} else {
		goto L369
	}
L357:
	;
	v1068 = F_tolower(m, v1064)
	mBase = m.M
	v1069 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1065))))
	v1070 = F_tolower(m, v1069)
	mBase = m.M
	goto L356
L358:
	;
	v1038 = v589
	v1039 = v1033
	v1040 = v1036
	goto L361
L359:
	;
	v1064 = int32(0)
	v1065 = v1033
	goto L357
L360:
	;
	v1064 = v1061 & int32(255)
	v1065 = v1060
	goto L357
L361:
	;
	v1042 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1039))))
	if v1042 == int32(0) {
		v1060 = v1039
		v1061 = v1040
		goto L360
	} else {
		goto L363
	}
L362:
	;
	v1060 = v1054
	v1061 = int32(0)
	goto L360
L363:
	;
	v1046 = v1040 & int32(255)
	if v1046 == v1042 {
		goto L364
	} else {
		goto L365
	}
L364:
	;
	v1053 = int32(1)
	v1054 = v1039 + v1053
	v1055 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1038)+1)))
	if v1055 != 0 {
		v1038 = v1038 + v1053
		v1039 = v1054
		v1040 = v1055
		goto L361
	} else {
		goto L367
	}
L365:
	;
	v1048 = F_tolower(m, v1046)
	mBase = m.M
	v1049 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1039))))
	v1050 = F_tolower(m, v1049)
	mBase = m.M
	if v1048 == v1050 {
		goto L364
	} else {
		goto L366
	}
L366:
	;
	v1052 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1038))))
	v1060 = v1039
	v1061 = v1052
	goto L360
L367:
	;
	goto L362
L368:
	;
	v1086 = int32(_a1448)
	v1089 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v589))))
	if v1089 != 0 {
		goto L374
	} else {
		goto L375
	}
L369:
	;
	if v1068-v1070 != 0 {
		goto L368
	} else {
		goto L370
	}
L370:
	;
	v1072 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v1074 = v579 + int32(1)
	v1078 = *(*int32)(unsafe.Add(mBase, uint32(v1072+v1074<<(uint(int32(2))%32))))
	v1081 = F_getLongLongFromObject(m, v1078, v11+int32(56))
	mBase = m.M
	v1082 = m.ExcPending
	if v1082 != 0 {
		goto L3
	} else {
		goto L371
	}
L371:
	;
	v1084 = *(*int64)(unsafe.Add(mBase, uint32(v11)+56))
	*(*uint32)(unsafe.Add(mBase, _consts[750])) = uint32(v1084)
	v1254 = v1074
	v1257 = v582
	goto L221
L372:
	;
	if int32(-2) < v579-v578 {
		goto L385
	} else {
		goto L386
	}
L373:
	;
	v1121 = F_tolower(m, v1117)
	mBase = m.M
	v1122 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1118))))
	v1123 = F_tolower(m, v1122)
	mBase = m.M
	goto L372
L374:
	;
	v1091 = v589
	v1092 = v1086
	v1093 = v1089
	goto L377
L375:
	;
	v1117 = int32(0)
	v1118 = v1086
	goto L373
L376:
	;
	v1117 = v1114 & int32(255)
	v1118 = v1113
	goto L373
L377:
	;
	v1095 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1092))))
	if v1095 == int32(0) {
		v1113 = v1092
		v1114 = v1093
		goto L376
	} else {
		goto L379
	}
L378:
	;
	v1113 = v1107
	v1114 = int32(0)
	goto L376
L379:
	;
	v1099 = v1093 & int32(255)
	if v1099 == v1095 {
		goto L380
	} else {
		goto L381
	}
L380:
	;
	v1106 = int32(1)
	v1107 = v1092 + v1106
	v1108 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1091)+1)))
	if v1108 != 0 {
		v1091 = v1091 + v1106
		v1092 = v1107
		v1093 = v1108
		goto L377
	} else {
		goto L383
	}
L381:
	;
	v1101 = F_tolower(m, v1099)
	mBase = m.M
	v1102 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1092))))
	v1103 = F_tolower(m, v1102)
	mBase = m.M
	if v1101 == v1103 {
		goto L380
	} else {
		goto L382
	}
L382:
	;
	v1105 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1091))))
	v1113 = v1092
	v1114 = v1105
	goto L376
L383:
	;
	goto L378
L384:
	;
	v1254 = v1248
	v1257 = int32(1)
	goto L221
L385:
	;
	v1167 = int32(_a1449)
	v1170 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v589))))
	if v1170 != 0 {
		goto L400
	} else {
		goto L401
	}
L386:
	;
	if v1121-v1123 != 0 {
		goto L385
	} else {
		goto L387
	}
L387:
	;
	v1125 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v1127 = v579 + int32(1)
	v1131 = *(*int32)(unsafe.Add(mBase, uint32(v1125+v1127<<(uint(int32(2))%32))))
	v1132 = int32(0)
	v1134 = *(*int32)(unsafe.Add(mBase, _consts[751]))
	F_sdsfree(m, v1134)
	mBase = m.M
	v1136 = m.ExcPending
	if v1136 != 0 {
		goto L3
	} else {
		goto L388
	}
L388:
	;
	v1137 = F_objectGetVal(m, v1131)
	mBase = m.M
	v1140 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1137+int32(-1)))))
	switch v1140 & int32(7) {
	case 0:
		goto L395
	case 1:
		goto L394
	case 2:
		goto L393
	case 3:
		goto L392
	case 4:
		goto L391
	default:
		v1163 = v1132
		goto L389
	}
L389:
	;
	*(*int32)(unsafe.Add(mBase, _consts[751])) = v1163
	v1248 = v1127
	goto L384
L390:
	;
	if v1157 == int32(0) {
		v1163 = v1132
		goto L389
	} else {
		goto L396
	}
L391:
	;
	v1156 = *(*int32)(unsafe.Add(mBase, uint32(v1137+int32(-17))))
	v1157 = v1156
	goto L390
L392:
	;
	v1153 = *(*int32)(unsafe.Add(mBase, uint32(v1137+int32(-9))))
	v1157 = v1153
	goto L390
L393:
	;
	v1150 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1137+int32(-5)))))
	v1157 = v1150
	goto L390
L394:
	;
	v1147 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1137+int32(-3)))))
	v1157 = v1147
	goto L390
L395:
	;
	v1157 = int32(base.Ui32(v1140) >> (uint(int32(3)) % 32))
	goto L390
L396:
	;
	v1160 = F_objectGetVal(m, v1131)
	mBase = m.M
	v1161 = F_sdsdup(m, v1160)
	mBase = m.M
	v1162 = m.ExcPending
	if v1162 != 0 {
		goto L3
	} else {
		goto L397
	}
L397:
	;
	v1163 = v1161
	goto L389
L398:
	;
	if int32(-2) < v579-v578 {
		goto L206
	} else {
		goto L410
	}
L399:
	;
	v1202 = F_tolower(m, v1198)
	mBase = m.M
	v1203 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1199))))
	v1204 = F_tolower(m, v1203)
	mBase = m.M
	goto L398
L400:
	;
	v1172 = v589
	v1173 = v1167
	v1174 = v1170
	goto L403
L401:
	;
	v1198 = int32(0)
	v1199 = v1167
	goto L399
L402:
	;
	v1198 = v1195 & int32(255)
	v1199 = v1194
	goto L399
L403:
	;
	v1176 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1173))))
	if v1176 == int32(0) {
		v1194 = v1173
		v1195 = v1174
		goto L402
	} else {
		goto L405
	}
L404:
	;
	v1194 = v1188
	v1195 = int32(0)
	goto L402
L405:
	;
	v1180 = v1174 & int32(255)
	if v1180 == v1176 {
		goto L406
	} else {
		goto L407
	}
L406:
	;
	v1187 = int32(1)
	v1188 = v1173 + v1187
	v1189 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1172)+1)))
	if v1189 != 0 {
		v1172 = v1172 + v1187
		v1173 = v1188
		v1174 = v1189
		goto L403
	} else {
		goto L409
	}
L407:
	;
	v1182 = F_tolower(m, v1180)
	mBase = m.M
	v1183 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1173))))
	v1184 = F_tolower(m, v1183)
	mBase = m.M
	if v1182 == v1184 {
		goto L406
	} else {
		goto L408
	}
L408:
	;
	v1186 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1172))))
	v1194 = v1173
	v1195 = v1186
	goto L402
L409:
	;
	goto L404
L410:
	;
	if v1202-v1204 != 0 {
		goto L206
	} else {
		goto L411
	}
L411:
	;
	v1206 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v1208 = v579 + int32(1)
	v1212 = *(*int32)(unsafe.Add(mBase, uint32(v1206+v1208<<(uint(int32(2))%32))))
	v1213 = int32(0)
	v1215 = *(*int32)(unsafe.Add(mBase, _consts[752]))
	F_sdsfree(m, v1215)
	mBase = m.M
	v1217 = m.ExcPending
	if v1217 != 0 {
		goto L3
	} else {
		goto L412
	}
L412:
	;
	v1218 = F_objectGetVal(m, v1212)
	mBase = m.M
	v1221 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1218+int32(-1)))))
	switch v1221 & int32(7) {
	case 0:
		goto L419
	case 1:
		goto L418
	case 2:
		goto L417
	case 3:
		goto L416
	case 4:
		goto L415
	default:
		v1244 = v1213
		goto L413
	}
L413:
	;
	*(*int32)(unsafe.Add(mBase, _consts[752])) = v1244
	v1248 = v1208
	goto L384
L414:
	;
	if v1238 == int32(0) {
		v1244 = v1213
		goto L413
	} else {
		goto L420
	}
L415:
	;
	v1237 = *(*int32)(unsafe.Add(mBase, uint32(v1218+int32(-17))))
	v1238 = v1237
	goto L414
L416:
	;
	v1234 = *(*int32)(unsafe.Add(mBase, uint32(v1218+int32(-9))))
	v1238 = v1234
	goto L414
L417:
	;
	v1231 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1218+int32(-5)))))
	v1238 = v1231
	goto L414
L418:
	;
	v1228 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1218+int32(-3)))))
	v1238 = v1228
	goto L414
L419:
	;
	v1238 = int32(base.Ui32(v1221) >> (uint(int32(3)) % 32))
	goto L414
L420:
	;
	v1241 = F_objectGetVal(m, v1212)
	mBase = m.M
	v1242 = F_sdsdup(m, v1241)
	mBase = m.M
	v1243 = m.ExcPending
	if v1243 != 0 {
		goto L3
	} else {
		goto L421
	}
L421:
	;
	v1244 = v1242
	goto L413
L422:
	;
	goto L208
L423:
	;
	if v1257 == int32(0) {
		goto L40
	} else {
		goto L424
	}
L424:
	;
	v1267 = F_sentinelDropConnections(m)
	mBase = m.M
	v1268 = m.ExcPending
	if v1268 != 0 {
		goto L3
	} else {
		goto L425
	}
L425:
	;
	goto L40
L426:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L427:
	;
	goto L40
L428:
	;
	m.G0 = v11 + int32(64)
	return
}
func F_sentinelFailoverReconfNextReplica(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v29 int32
	_ = v29
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v42 int64
	_ = v42
	var v43 int64
	_ = v43
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
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v80 int32
	_ = v80
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v100 int32
	_ = v100
	var v105 int32
	_ = v105
	var v109 int32
	_ = v109
	var v113 int32
	_ = v113
	var v120 int32
	_ = v120
	var v122 int32
	_ = v122
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v130 int32
	_ = v130
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v142 int32
	_ = v142
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v150 int32
	_ = v150
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v159 int32
	_ = v159
	var v161 int32
	_ = v161
	var v163 int64
	_ = v163
	var v164 int64
	_ = v164
	var v165 int64
	_ = v165
	var v166 int64
	_ = v166
	var v167 int64
	_ = v167
	var v168 int64
	_ = v168
	var v169 int64
	_ = v169
	var v171 int64
	_ = v171
	var v173 int64
	_ = v173
	var v175 int64
	_ = v175
	var v177 int64
	_ = v177
	var v179 int64
	_ = v179
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v186 int32
	_ = v186
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v193 int32
	_ = v193
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v201 int32
	_ = v201
	var v205 int32
	_ = v205
	var v207 int32
	_ = v207
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v221 int32
	_ = v221
	var v226 int32
	_ = v226
	var v230 int32
	_ = v230
	var v234 int32
	_ = v234
	var v239 int32
	_ = v239
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v248 int32
	_ = v248
	var v251 int32
	_ = v251
	var v262 int32
	_ = v262
	var v263 int32
	_ = v263
	var v266 int32
	_ = v266
	var v269 int32
	_ = v269
	var v270 int32
	_ = v270
	var v274 int32
	_ = v274
	var v278 int32
	_ = v278
	var v279 int32
	_ = v279
	var v280 int32
	_ = v280
	var v283 int32
	_ = v283
	var v285 int32
	_ = v285
	var v287 int64
	_ = v287
	var v288 int64
	_ = v288
	var v289 int64
	_ = v289
	var v290 int64
	_ = v290
	var v291 int64
	_ = v291
	var v292 int64
	_ = v292
	var v293 int64
	_ = v293
	var v295 int64
	_ = v295
	var v297 int64
	_ = v297
	var v299 int64
	_ = v299
	var v301 int64
	_ = v301
	var v303 int64
	_ = v303
	var v305 int32
	_ = v305
	var v306 int32
	_ = v306
	var v307 int32
	_ = v307
	var v310 int32
	_ = v310
	var v313 int32
	_ = v313
	var v314 int32
	_ = v314
	var v315 int32
	_ = v315
	var v317 int32
	_ = v317
	var v320 int32
	_ = v320
	var v321 int32
	_ = v321
	var v325 int32
	_ = v325
	var v329 int32
	_ = v329
	var v331 int32
	_ = v331
	var v338 int32
	_ = v338
	var v339 int32
	_ = v339
	var v340 int32
	_ = v340
	var v345 int32
	_ = v345
	var v350 int32
	_ = v350
	var v354 int32
	_ = v354
	var v358 int32
	_ = v358
	var v364 int32
	_ = v364
	var v365 int32
	_ = v365
	var v372 int64
	_ = v372
	var v373 int64
	_ = v373
	var v376 int64
	_ = v376
	var v378 int32
	_ = v378
	var v384 int32
	_ = v384
	var v385 int32
	_ = v385
	var v389 int32
	_ = v389
	var v391 int32
	_ = v391
	var v394 int32
	_ = v394
	var v395 int32
	_ = v395
	var v396 int32
	_ = v396
	var v397 int32
	_ = v397
	var v398 int32
	_ = v398
	var v399 int32
	_ = v399
	var v400 int32
	_ = v400
	var v404 int64
	_ = v404
	var v411 int32
	_ = v411
	var v414 int32
	_ = v414
	var v416 int32
	_ = v416
	var v424 int32
	_ = v424
	var v426 int32
	_ = v426
	v2 = int32(0)
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+148))
	v8 = F_dictGetIterator(m, v7)
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	F_dictReleaseIterator(m, v8)
	mBase = m.M
	v244 = m.ExcPending
	if v244 != 0 {
		goto L2
	} else {
		goto L61
	}
L2:
	;
	return
L3:
	;
	v17 = v8 + int32(20)
	v18 = *(*int32)(unsafe.Add(mBase, uint32(v8)+16))
	if v18 != 0 {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	if v113 == int32(0) {
		v239 = v2
		goto L1
	} else {
		goto L30
	}
L5:
	;
	v24 = v17
	v25 = v21
	goto L8
L6:
	;
	v21 = int32(1)
	goto L5
L7:
	;
	v21 = int32(0)
	goto L5
L8:
	;
	switch v25 {
	case 0:
		goto L13
	default:
		goto L12
	}
L10:
	;
	v25 = int32(0)
	goto L8
L11:
	;
	goto L4
L12:
	;
	v105 = *(*int32)(unsafe.Add(mBase, uint32(v24)))
	*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = v105
	if v105 == int32(0) {
		goto L10
	} else {
		goto L29
	}
L13:
	;
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v8)+4))
	if v29 != int32(-1) {
		v68 = v29
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v69 = int32(1)
	v70 = v68 + v69
	*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = v70
	v72 = int32(0)
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v8)))
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v8)+8))
	v80 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v75+v76+int32(26)))))
	if v80 == int32(255) {
		goto L23
	} else {
		goto L24
	}
L15:
	;
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v8)+8))
	if v33 != 0 {
		v68 = int32(-1)
		goto L14
	} else {
		goto L16
	}
L16:
	;
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v8)))
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v8)+12))
	if v35 == int32(0) {
		goto L18
	} else {
		goto L19
	}
L17:
	;
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v61)+20))
	if v62 != int32(-1) {
		goto L20
	} else {
		goto L21
	}
L18:
	;
	v42 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v34)+16)))
	v43 = int64(*(*int8)(unsafe.Add(mBase, uint32(v34)+27)))
	v44 = int64(*(*int32)(unsafe.Add(mBase, uint32(v34)+8)))
	v45 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v34)+12)))
	v46 = int64(*(*int8)(unsafe.Add(mBase, uint32(v34)+26)))
	v47 = int64(*(*int32)(unsafe.Add(mBase, uint32(v34)+4)))
	v48 = F_wangHash64(m, v47)
	mBase = m.M
	v50 = F_wangHash64(m, v46+v48)
	mBase = m.M
	v52 = F_wangHash64(m, v45+v50)
	mBase = m.M
	v54 = F_wangHash64(m, v44+v52)
	mBase = m.M
	v56 = F_wangHash64(m, v43+v54)
	mBase = m.M
	v58 = F_wangHash64(m, v42+v56)
	mBase = m.M
	*(*int64)(unsafe.Add(mBase, uint32(v8)+24)) = v58
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v8)))
	v61 = v60
	goto L17
L19:
	;
	v38 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v34)+24)))
	v40 = v38 + int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v34)+24)) = uint16(v40)
	v61 = v34
	goto L17
L20:
	;
	v68 = v62 + int32(-1)
	goto L14
L21:
	;
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v8)+4))
	v68 = v65
	goto L14
L22:
	;
	v95 = int32(2)
	v100 = *(*int32)(unsafe.Add(mBase, uint32(v75+v93<<(uint(v95)%32)+int32(4))))
	v24 = v100 + v94<<(uint(v95)%32)
	v25 = int32(1)
	goto L8
L23:
	;
	v84 = v72
	goto L25
L24:
	;
	v84 = v69 << (uint(v80) % 32)
	goto L25
L25:
	;
	if v70 < v84 {
		v93 = v76
		v94 = v70
		goto L22
	} else {
		goto L26
	}
L26:
	;
	if v76 != 0 {
		v113 = v72
		goto L11
	} else {
		goto L27
	}
L27:
	;
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v75)+20))
	if v86 == int32(-1) {
		v113 = v72
		goto L11
	} else {
		goto L28
	}
L28:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v8)+4)) = int64(4294967296)
	v93 = int32(1)
	v94 = int32(0)
	goto L22
L29:
	;
	v109 = *(*int32)(unsafe.Add(mBase, uint32(v105)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v17))) = v109
	v113 = v105
	goto L11
L30:
	;
	v120 = v2
	v122 = v113
	goto L31
L31:
	;
	v124 = *(*int32)(unsafe.Add(mBase, uint32(v122)+8))
	goto L33
L32:
	;
	v239 = v130
	goto L1
L33:
	;
	v125 = *(*int32)(unsafe.Add(mBase, uint32(v124)))
	v130 = v120 + base.B2i32(v125&int32(768) != int32(0))
	v138 = v8 + int32(20)
	v139 = *(*int32)(unsafe.Add(mBase, uint32(v8)+16))
	if v139 != 0 {
		goto L36
	} else {
		goto L37
	}
L34:
	;
	if v234 != 0 {
		v120 = v130
		v122 = v234
		goto L31
	} else {
		goto L60
	}
L35:
	;
	v145 = v138
	v146 = v142
	goto L38
L36:
	;
	v142 = int32(1)
	goto L35
L37:
	;
	v142 = int32(0)
	goto L35
L38:
	;
	switch v146 {
	case 0:
		goto L43
	default:
		goto L42
	}
L40:
	;
	v146 = int32(0)
	goto L38
L41:
	;
	goto L34
L42:
	;
	v226 = *(*int32)(unsafe.Add(mBase, uint32(v145)))
	*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = v226
	if v226 == int32(0) {
		goto L40
	} else {
		goto L59
	}
L43:
	;
	v150 = *(*int32)(unsafe.Add(mBase, uint32(v8)+4))
	if v150 != int32(-1) {
		v189 = v150
		goto L44
	} else {
		goto L45
	}
L44:
	;
	v190 = int32(1)
	v191 = v189 + v190
	*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = v191
	v193 = int32(0)
	v196 = *(*int32)(unsafe.Add(mBase, uint32(v8)))
	v197 = *(*int32)(unsafe.Add(mBase, uint32(v8)+8))
	v201 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v196+v197+int32(26)))))
	if v201 == int32(255) {
		goto L53
	} else {
		goto L54
	}
L45:
	;
	v154 = *(*int32)(unsafe.Add(mBase, uint32(v8)+8))
	if v154 != 0 {
		v189 = int32(-1)
		goto L44
	} else {
		goto L46
	}
L46:
	;
	v155 = *(*int32)(unsafe.Add(mBase, uint32(v8)))
	v156 = *(*int32)(unsafe.Add(mBase, uint32(v8)+12))
	if v156 == int32(0) {
		goto L48
	} else {
		goto L49
	}
L47:
	;
	v183 = *(*int32)(unsafe.Add(mBase, uint32(v182)+20))
	if v183 != int32(-1) {
		goto L50
	} else {
		goto L51
	}
L48:
	;
	v163 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v155)+16)))
	v164 = int64(*(*int8)(unsafe.Add(mBase, uint32(v155)+27)))
	v165 = int64(*(*int32)(unsafe.Add(mBase, uint32(v155)+8)))
	v166 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v155)+12)))
	v167 = int64(*(*int8)(unsafe.Add(mBase, uint32(v155)+26)))
	v168 = int64(*(*int32)(unsafe.Add(mBase, uint32(v155)+4)))
	v169 = F_wangHash64(m, v168)
	mBase = m.M
	v171 = F_wangHash64(m, v167+v169)
	mBase = m.M
	v173 = F_wangHash64(m, v166+v171)
	mBase = m.M
	v175 = F_wangHash64(m, v165+v173)
	mBase = m.M
	v177 = F_wangHash64(m, v164+v175)
	mBase = m.M
	v179 = F_wangHash64(m, v163+v177)
	mBase = m.M
	*(*int64)(unsafe.Add(mBase, uint32(v8)+24)) = v179
	v181 = *(*int32)(unsafe.Add(mBase, uint32(v8)))
	v182 = v181
	goto L47
L49:
	;
	v159 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v155)+24)))
	v161 = v159 + int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v155)+24)) = uint16(v161)
	v182 = v155
	goto L47
L50:
	;
	v189 = v183 + int32(-1)
	goto L44
L51:
	;
	v186 = *(*int32)(unsafe.Add(mBase, uint32(v8)+4))
	v189 = v186
	goto L44
L52:
	;
	v216 = int32(2)
	v221 = *(*int32)(unsafe.Add(mBase, uint32(v196+v214<<(uint(v216)%32)+int32(4))))
	v145 = v221 + v215<<(uint(v216)%32)
	v146 = int32(1)
	goto L38
L53:
	;
	v205 = v193
	goto L55
L54:
	;
	v205 = v190 << (uint(v201) % 32)
	goto L55
L55:
	;
	if v191 < v205 {
		v214 = v197
		v215 = v191
		goto L52
	} else {
		goto L56
	}
L56:
	;
	if v197 != 0 {
		v234 = v193
		goto L41
	} else {
		goto L57
	}
L57:
	;
	v207 = *(*int32)(unsafe.Add(mBase, uint32(v196)+20))
	if v207 == int32(-1) {
		v234 = v193
		goto L41
	} else {
		goto L58
	}
L58:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v8)+4)) = int64(4294967296)
	v214 = int32(1)
	v215 = int32(0)
	goto L52
L59:
	;
	v230 = *(*int32)(unsafe.Add(mBase, uint32(v226)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v138))) = v230
	v234 = v226
	goto L41
L60:
	;
	goto L32
L61:
	;
	v245 = *(*int32)(unsafe.Add(mBase, uint32(l0)+148))
	v246 = F_dictGetIterator(m, v245)
	mBase = m.M
	v247 = m.ExcPending
	if v247 != 0 {
		goto L2
	} else {
		goto L62
	}
L62:
	;
	v248 = *(*int32)(unsafe.Add(mBase, uint32(l0)+156))
	if v248 <= v239 {
		goto L63
	} else {
		goto L64
	}
L63:
	;
	F_dictReleaseIterator(m, v246)
	mBase = m.M
	v424 = m.ExcPending
	if v424 != 0 {
		goto L2
	} else {
		goto L108
	}
L64:
	;
	v251 = v239
	goto L65
L65:
	;
	v262 = v246 + int32(20)
	v263 = *(*int32)(unsafe.Add(mBase, uint32(v246)+16))
	if v263 != 0 {
		goto L69
	} else {
		goto L70
	}
L66:
	;
	goto L63
L67:
	;
	if v358 == int32(0) {
		goto L63
	} else {
		goto L93
	}
L68:
	;
	v269 = v262
	v270 = v266
	goto L71
L69:
	;
	v266 = int32(1)
	goto L68
L70:
	;
	v266 = int32(0)
	goto L68
L71:
	;
	switch v270 {
	case 0:
		goto L76
	default:
		goto L75
	}
L73:
	;
	v270 = int32(0)
	goto L71
L74:
	;
	goto L67
L75:
	;
	v350 = *(*int32)(unsafe.Add(mBase, uint32(v269)))
	*(*int32)(unsafe.Add(mBase, uint32(v246)+16)) = v350
	if v350 == int32(0) {
		goto L73
	} else {
		goto L92
	}
L76:
	;
	v274 = *(*int32)(unsafe.Add(mBase, uint32(v246)+4))
	if v274 != int32(-1) {
		v313 = v274
		goto L77
	} else {
		goto L78
	}
L77:
	;
	v314 = int32(1)
	v315 = v313 + v314
	*(*int32)(unsafe.Add(mBase, uint32(v246)+4)) = v315
	v317 = int32(0)
	v320 = *(*int32)(unsafe.Add(mBase, uint32(v246)))
	v321 = *(*int32)(unsafe.Add(mBase, uint32(v246)+8))
	v325 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v320+v321+int32(26)))))
	if v325 == int32(255) {
		goto L86
	} else {
		goto L87
	}
L78:
	;
	v278 = *(*int32)(unsafe.Add(mBase, uint32(v246)+8))
	if v278 != 0 {
		v313 = int32(-1)
		goto L77
	} else {
		goto L79
	}
L79:
	;
	v279 = *(*int32)(unsafe.Add(mBase, uint32(v246)))
	v280 = *(*int32)(unsafe.Add(mBase, uint32(v246)+12))
	if v280 == int32(0) {
		goto L81
	} else {
		goto L82
	}
L80:
	;
	v307 = *(*int32)(unsafe.Add(mBase, uint32(v306)+20))
	if v307 != int32(-1) {
		goto L83
	} else {
		goto L84
	}
L81:
	;
	v287 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v279)+16)))
	v288 = int64(*(*int8)(unsafe.Add(mBase, uint32(v279)+27)))
	v289 = int64(*(*int32)(unsafe.Add(mBase, uint32(v279)+8)))
	v290 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v279)+12)))
	v291 = int64(*(*int8)(unsafe.Add(mBase, uint32(v279)+26)))
	v292 = int64(*(*int32)(unsafe.Add(mBase, uint32(v279)+4)))
	v293 = F_wangHash64(m, v292)
	mBase = m.M
	v295 = F_wangHash64(m, v291+v293)
	mBase = m.M
	v297 = F_wangHash64(m, v290+v295)
	mBase = m.M
	v299 = F_wangHash64(m, v289+v297)
	mBase = m.M
	v301 = F_wangHash64(m, v288+v299)
	mBase = m.M
	v303 = F_wangHash64(m, v287+v301)
	mBase = m.M
	*(*int64)(unsafe.Add(mBase, uint32(v246)+24)) = v303
	v305 = *(*int32)(unsafe.Add(mBase, uint32(v246)))
	v306 = v305
	goto L80
L82:
	;
	v283 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v279)+24)))
	v285 = v283 + int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v279)+24)) = uint16(v285)
	v306 = v279
	goto L80
L83:
	;
	v313 = v307 + int32(-1)
	goto L77
L84:
	;
	v310 = *(*int32)(unsafe.Add(mBase, uint32(v246)+4))
	v313 = v310
	goto L77
L85:
	;
	v340 = int32(2)
	v345 = *(*int32)(unsafe.Add(mBase, uint32(v320+v338<<(uint(v340)%32)+int32(4))))
	v269 = v345 + v339<<(uint(v340)%32)
	v270 = int32(1)
	goto L71
L86:
	;
	v329 = v317
	goto L88
L87:
	;
	v329 = v314 << (uint(v325) % 32)
	goto L88
L88:
	;
	if v315 < v329 {
		v338 = v321
		v339 = v315
		goto L85
	} else {
		goto L89
	}
L89:
	;
	if v321 != 0 {
		v358 = v317
		goto L74
	} else {
		goto L90
	}
L90:
	;
	v331 = *(*int32)(unsafe.Add(mBase, uint32(v320)+20))
	if v331 == int32(-1) {
		v358 = v317
		goto L74
	} else {
		goto L91
	}
L91:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v246)+4)) = int64(4294967296)
	v338 = int32(1)
	v339 = int32(0)
	goto L85
L92:
	;
	v354 = *(*int32)(unsafe.Add(mBase, uint32(v350)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v262))) = v354
	v358 = v350
	goto L74
L93:
	;
	v364 = *(*int32)(unsafe.Add(mBase, uint32(v358)+8))
	goto L95
L94:
	;
	v416 = *(*int32)(unsafe.Add(mBase, uint32(l0)+156))
	if v414 < v416 {
		v251 = v414
		goto L65
	} else {
		goto L107
	}
L95:
	;
	v365 = *(*int32)(unsafe.Add(mBase, uint32(v364)))
	if v365&int32(1152) != 0 {
		v414 = v251
		goto L94
	} else {
		goto L96
	}
L96:
	;
	if v365&int32(256) == int32(0) {
		v391 = v365
		goto L97
	} else {
		goto L98
	}
L97:
	;
	if v391&int32(768) != 0 {
		v414 = v251
		goto L94
	} else {
		goto L102
	}
L98:
	;
	v372 = F_mstime(m)
	mBase = m.M
	v373 = *(*int64)(unsafe.Add(mBase, uint32(v364)+184))
	v376 = *(*int64)(unsafe.Add(mBase, _consts[775]))
	if v376 < v372-v373 {
		goto L99
	} else {
		goto L100
	}
L99:
	;
	F_sentinelEvent(m, int32(2), int32(_a1550), v364, int32(_a1362), int32(0))
	mBase = m.M
	v384 = m.ExcPending
	if v384 != 0 {
		goto L2
	} else {
		goto L101
	}
L100:
	;
	v378 = *(*int32)(unsafe.Add(mBase, uint32(v364)))
	v391 = v378
	goto L97
L101:
	;
	v385 = *(*int32)(unsafe.Add(mBase, uint32(v364)))
	v389 = v385&int32(-1281) | int32(1024)
	*(*int32)(unsafe.Add(mBase, uint32(v364))) = v389
	v391 = v389
	goto L97
L102:
	;
	v394 = *(*int32)(unsafe.Add(mBase, uint32(v364)+28))
	v395 = *(*int32)(unsafe.Add(mBase, uint32(v394)+4))
	if v395 != 0 {
		v414 = v251
		goto L94
	} else {
		goto L103
	}
L103:
	;
	v396 = *(*int32)(unsafe.Add(mBase, uint32(l0)+280))
	v397 = *(*int32)(unsafe.Add(mBase, uint32(v396)+24))
	v398 = F_sentinelSendReplicaOf(m, v364, v397)
	mBase = m.M
	v399 = m.ExcPending
	if v399 != 0 {
		goto L2
	} else {
		goto L104
	}
L104:
	;
	if v398 != 0 {
		v414 = v251
		goto L94
	} else {
		goto L105
	}
L105:
	;
	v400 = *(*int32)(unsafe.Add(mBase, uint32(v364)))
	*(*int32)(unsafe.Add(mBase, uint32(v364))) = v400 | int32(256)
	v404 = F_mstime(m)
	mBase = m.M
	*(*int64)(unsafe.Add(mBase, uint32(v364)+184)) = v404
	F_sentinelEvent(m, int32(2), int32(_a1551), v364, int32(_a1362), int32(0))
	mBase = m.M
	v411 = m.ExcPending
	if v411 != 0 {
		goto L2
	} else {
		goto L106
	}
L106:
	;
	v414 = v251 + int32(1)
	goto L94
L107:
	;
	goto L66
L108:
	;
	F_sentinelFailoverDetectEnd(m, l0)
	mBase = m.M
	v426 = m.ExcPending
	if v426 != 0 {
		goto L2
	} else {
		goto L109
	}
L109:
	;
	return
}
func F_sentinelFailoverSelectReplica(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v26 int64
	_ = v26
	var v33 int32
	_ = v33
	v3 = F_sentinelSelectReplica(m, l0)
	mBase = m.M
	v4 = m.ExcPending
	if v4 != 0 {
		return
	} else {
		if v3 != 0 {
			F_sentinelEvent(m, int32(3), int32(_a1547), v3, int32(_a1362), int32(0))
			mBase = m.M
			v18 = m.ExcPending
			if v18 != 0 {
				return
			} else {
				v19 = *(*int32)(unsafe.Add(mBase, uint32(v3)))
				*(*int32)(unsafe.Add(mBase, uint32(v3))) = v19 | int32(128)
				*(*int32)(unsafe.Add(mBase, uint32(l0)+240)) = int32(3)
				*(*int32)(unsafe.Add(mBase, uint32(l0)+280)) = v3
				v26 = F_mstime(m)
				mBase = m.M
				*(*int64)(unsafe.Add(mBase, uint32(l0)+248)) = v26
				F_sentinelEvent(m, int32(2), int32(_a1548), v3, int32(_a1362), int32(0))
				mBase = m.M
				v33 = m.ExcPending
				if v33 != 0 {
					return
				} else {
					return
				}
			}
		} else {
			F_sentinelEvent(m, int32(3), int32(_a1549), l0, int32(_a1362), int32(0))
			mBase = m.M
			v10 = m.ExcPending
			if v10 != 0 {
				return
			} else {
				F_sentinelAbortFailover(m, l0)
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
func F_sentinelGetLeader(m *base.Module, l0 int32, l1 int64) int32 {
	mBase := m.M
	_ = mBase
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
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
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v51 int32
	_ = v51
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v64 int64
	_ = v64
	var v65 int64
	_ = v65
	var v66 int64
	_ = v66
	var v67 int64
	_ = v67
	var v68 int64
	_ = v68
	var v69 int64
	_ = v69
	var v70 int64
	_ = v70
	var v72 int64
	_ = v72
	var v74 int64
	_ = v74
	var v76 int64
	_ = v76
	var v78 int64
	_ = v78
	var v80 int64
	_ = v80
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v102 int32
	_ = v102
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v122 int32
	_ = v122
	var v127 int32
	_ = v127
	var v131 int32
	_ = v131
	var v135 int32
	_ = v135
	var v145 int32
	_ = v145
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v157 int64
	_ = v157
	var v159 int64
	_ = v159
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v174 int32
	_ = v174
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v182 int32
	_ = v182
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v191 int32
	_ = v191
	var v193 int32
	_ = v193
	var v195 int64
	_ = v195
	var v196 int64
	_ = v196
	var v197 int64
	_ = v197
	var v198 int64
	_ = v198
	var v199 int64
	_ = v199
	var v200 int64
	_ = v200
	var v201 int64
	_ = v201
	var v203 int64
	_ = v203
	var v205 int64
	_ = v205
	var v207 int64
	_ = v207
	var v209 int64
	_ = v209
	var v211 int64
	_ = v211
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
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v225 int32
	_ = v225
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v233 int32
	_ = v233
	var v237 int32
	_ = v237
	var v239 int32
	_ = v239
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v248 int32
	_ = v248
	var v253 int32
	_ = v253
	var v258 int32
	_ = v258
	var v262 int32
	_ = v262
	var v266 int32
	_ = v266
	var v283 int32
	_ = v283
	var v284 int32
	_ = v284
	var v285 int32
	_ = v285
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
	var v305 int32
	_ = v305
	var v309 int32
	_ = v309
	var v310 int32
	_ = v310
	var v311 int32
	_ = v311
	var v314 int32
	_ = v314
	var v316 int32
	_ = v316
	var v318 int64
	_ = v318
	var v319 int64
	_ = v319
	var v320 int64
	_ = v320
	var v321 int64
	_ = v321
	var v322 int64
	_ = v322
	var v323 int64
	_ = v323
	var v324 int64
	_ = v324
	var v326 int64
	_ = v326
	var v328 int64
	_ = v328
	var v330 int64
	_ = v330
	var v332 int64
	_ = v332
	var v334 int64
	_ = v334
	var v336 int32
	_ = v336
	var v337 int32
	_ = v337
	var v338 int32
	_ = v338
	var v341 int32
	_ = v341
	var v344 int32
	_ = v344
	var v345 int32
	_ = v345
	var v346 int32
	_ = v346
	var v348 int32
	_ = v348
	var v351 int32
	_ = v351
	var v352 int32
	_ = v352
	var v356 int32
	_ = v356
	var v360 int32
	_ = v360
	var v362 int32
	_ = v362
	var v369 int32
	_ = v369
	var v370 int32
	_ = v370
	var v371 int32
	_ = v371
	var v376 int32
	_ = v376
	var v381 int32
	_ = v381
	var v385 int32
	_ = v385
	var v389 int32
	_ = v389
	var v401 int32
	_ = v401
	var v404 int32
	_ = v404
	var v406 int64
	_ = v406
	var v409 int64
	_ = v409
	var v411 int32
	_ = v411
	var v412 int32
	_ = v412
	var v413 int64
	_ = v413
	var v421 int32
	_ = v421
	var v422 int32
	_ = v422
	var v425 int32
	_ = v425
	var v428 int32
	_ = v428
	var v429 int32
	_ = v429
	var v433 int32
	_ = v433
	var v437 int32
	_ = v437
	var v438 int32
	_ = v438
	var v439 int32
	_ = v439
	var v442 int32
	_ = v442
	var v444 int32
	_ = v444
	var v446 int64
	_ = v446
	var v447 int64
	_ = v447
	var v448 int64
	_ = v448
	var v449 int64
	_ = v449
	var v450 int64
	_ = v450
	var v451 int64
	_ = v451
	var v452 int64
	_ = v452
	var v454 int64
	_ = v454
	var v456 int64
	_ = v456
	var v458 int64
	_ = v458
	var v460 int64
	_ = v460
	var v462 int64
	_ = v462
	var v464 int32
	_ = v464
	var v465 int32
	_ = v465
	var v466 int32
	_ = v466
	var v469 int32
	_ = v469
	var v472 int32
	_ = v472
	var v473 int32
	_ = v473
	var v474 int32
	_ = v474
	var v476 int32
	_ = v476
	var v479 int32
	_ = v479
	var v480 int32
	_ = v480
	var v484 int32
	_ = v484
	var v488 int32
	_ = v488
	var v490 int32
	_ = v490
	var v497 int32
	_ = v497
	var v498 int32
	_ = v498
	var v499 int32
	_ = v499
	var v504 int32
	_ = v504
	var v509 int32
	_ = v509
	var v513 int32
	_ = v513
	var v517 int32
	_ = v517
	var v522 int32
	_ = v522
	var v523 int32
	_ = v523
	var v531 int32
	_ = v531
	var v534 int32
	_ = v534
	var v546 int64
	_ = v546
	var v548 int32
	_ = v548
	var v558 int32
	_ = v558
	var v559 int32
	_ = v559
	var v560 int64
	_ = v560
	var v562 int32
	_ = v562
	var v565 int32
	_ = v565
	var v566 int32
	_ = v566
	var v569 int64
	_ = v569
	var v571 int32
	_ = v571
	var v572 int32
	_ = v572
	var v573 int64
	_ = v573
	var v574 int32
	_ = v574
	var v583 int64
	_ = v583
	var v584 int32
	_ = v584
	var v586 int64
	_ = v586
	var v588 int32
	_ = v588
	var v589 int32
	_ = v589
	var v591 int32
	_ = v591
	var v599 int64
	_ = v599
	var v601 int32
	_ = v601
	var v602 int32
	_ = v602
	var v604 int32
	_ = v604
	var v609 int32
	_ = v609
	var v611 int32
	_ = v611
	v13 = m.G0
	v15 = v13 - int32(16)
	m.G0 = v15
	v17 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v17&int32(80) == int32(0) {
		goto L4
	} else {
		goto L5
	}
L1:
	;
	v565 = F_sentinelVoteLeader(m, l0, l1, v558, v15+int32(8))
	mBase = m.M
	v566 = m.ExcPending
	if v566 != 0 {
		goto L6
	} else {
		goto L142
	}
L2:
	;
	v558 = int32(_a1336)
	v559 = int32(1)
	v560 = v546
	v562 = v548
	goto L1
L3:
	;
	F_dictReleaseIterator(m, v284)
	mBase = m.M
	v534 = m.ExcPending
	if v534 != 0 {
		goto L6
	} else {
		goto L137
	}
L4:
	;
	F__serverAssert(m, int32(_a1544), int32(_a1333), int32(4700))
	mBase = m.M
	v531 = m.ExcPending
	if v531 != 0 {
		goto L6
	} else {
		goto L136
	}
L5:
	;
	v23 = F_dictCreate(m, int32(_a1545))
	mBase = m.M
	v26 = m.ExcPending
	if v26 != 0 {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	return int32(0)
L7:
	;
	v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)+144))
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v27)+16))
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v27)+12))
	v30 = F_dictGetIterator(m, v27)
	mBase = m.M
	v31 = m.ExcPending
	if v31 != 0 {
		goto L6
	} else {
		goto L9
	}
L8:
	;
	F_dictReleaseIterator(m, v30)
	mBase = m.M
	v283 = m.ExcPending
	if v283 != 0 {
		goto L6
	} else {
		goto L71
	}
L9:
	;
	v39 = v30 + int32(20)
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v30)+16))
	if v40 != 0 {
		goto L12
	} else {
		goto L13
	}
L10:
	;
	if v135 == int32(0) {
		goto L8
	} else {
		goto L36
	}
L11:
	;
	v46 = v39
	v47 = v43
	goto L14
L12:
	;
	v43 = int32(1)
	goto L11
L13:
	;
	v43 = int32(0)
	goto L11
L14:
	;
	switch v47 {
	case 0:
		goto L19
	default:
		goto L18
	}
L16:
	;
	v47 = int32(0)
	goto L14
L17:
	;
	goto L10
L18:
	;
	v127 = *(*int32)(unsafe.Add(mBase, uint32(v46)))
	*(*int32)(unsafe.Add(mBase, uint32(v30)+16)) = v127
	if v127 == int32(0) {
		goto L16
	} else {
		goto L35
	}
L19:
	;
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v30)+4))
	if v51 != int32(-1) {
		v90 = v51
		goto L20
	} else {
		goto L21
	}
L20:
	;
	v91 = int32(1)
	v92 = v90 + v91
	*(*int32)(unsafe.Add(mBase, uint32(v30)+4)) = v92
	v94 = int32(0)
	v97 = *(*int32)(unsafe.Add(mBase, uint32(v30)))
	v98 = *(*int32)(unsafe.Add(mBase, uint32(v30)+8))
	v102 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v97+v98+int32(26)))))
	if v102 == int32(255) {
		goto L29
	} else {
		goto L30
	}
L21:
	;
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v30)+8))
	if v55 != 0 {
		v90 = int32(-1)
		goto L20
	} else {
		goto L22
	}
L22:
	;
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v30)))
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v30)+12))
	if v57 == int32(0) {
		goto L24
	} else {
		goto L25
	}
L23:
	;
	v84 = *(*int32)(unsafe.Add(mBase, uint32(v83)+20))
	if v84 != int32(-1) {
		goto L26
	} else {
		goto L27
	}
L24:
	;
	v64 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v56)+16)))
	v65 = int64(*(*int8)(unsafe.Add(mBase, uint32(v56)+27)))
	v66 = int64(*(*int32)(unsafe.Add(mBase, uint32(v56)+8)))
	v67 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v56)+12)))
	v68 = int64(*(*int8)(unsafe.Add(mBase, uint32(v56)+26)))
	v69 = int64(*(*int32)(unsafe.Add(mBase, uint32(v56)+4)))
	v70 = F_wangHash64(m, v69)
	mBase = m.M
	v72 = F_wangHash64(m, v68+v70)
	mBase = m.M
	v74 = F_wangHash64(m, v67+v72)
	mBase = m.M
	v76 = F_wangHash64(m, v66+v74)
	mBase = m.M
	v78 = F_wangHash64(m, v65+v76)
	mBase = m.M
	v80 = F_wangHash64(m, v64+v78)
	mBase = m.M
	*(*int64)(unsafe.Add(mBase, uint32(v30)+24)) = v80
	v82 = *(*int32)(unsafe.Add(mBase, uint32(v30)))
	v83 = v82
	goto L23
L25:
	;
	v60 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v56)+24)))
	v62 = v60 + int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v56)+24)) = uint16(v62)
	v83 = v56
	goto L23
L26:
	;
	v90 = v84 + int32(-1)
	goto L20
L27:
	;
	v87 = *(*int32)(unsafe.Add(mBase, uint32(v30)+4))
	v90 = v87
	goto L20
L28:
	;
	v117 = int32(2)
	v122 = *(*int32)(unsafe.Add(mBase, uint32(v97+v115<<(uint(v117)%32)+int32(4))))
	v46 = v122 + v116<<(uint(v117)%32)
	v47 = int32(1)
	goto L14
L29:
	;
	v106 = v94
	goto L31
L30:
	;
	v106 = v91 << (uint(v102) % 32)
	goto L31
L31:
	;
	if v92 < v106 {
		v115 = v98
		v116 = v92
		goto L28
	} else {
		goto L32
	}
L32:
	;
	if v98 != 0 {
		v135 = v94
		goto L17
	} else {
		goto L33
	}
L33:
	;
	v108 = *(*int32)(unsafe.Add(mBase, uint32(v97)+20))
	if v108 == int32(-1) {
		v135 = v94
		goto L17
	} else {
		goto L34
	}
L34:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v30)+4)) = int64(4294967296)
	v115 = int32(1)
	v116 = int32(0)
	goto L28
L35:
	;
	v131 = *(*int32)(unsafe.Add(mBase, uint32(v127)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v39))) = v131
	v135 = v127
	goto L17
L36:
	;
	v145 = v135
	goto L37
L37:
	;
	v153 = *(*int32)(unsafe.Add(mBase, uint32(v145)+8))
	goto L40
L38:
	;
	goto L8
L39:
	;
	v170 = v30 + int32(20)
	v171 = *(*int32)(unsafe.Add(mBase, uint32(v30)+16))
	if v171 != 0 {
		goto L46
	} else {
		goto L47
	}
L40:
	;
	v154 = *(*int32)(unsafe.Add(mBase, uint32(v153)+216))
	if v154 == int32(0) {
		goto L39
	} else {
		goto L41
	}
L41:
	;
	v157 = *(*int64)(unsafe.Add(mBase, uint32(v153)+224))
	v159 = *(*int64)(unsafe.Add(mBase, _consts[748]))
	if v157 != v159 {
		goto L39
	} else {
		goto L42
	}
L42:
	;
	v161 = F_sentinelLeaderIncr(m, v23, v154)
	mBase = m.M
	v162 = m.ExcPending
	if v162 != 0 {
		goto L6
	} else {
		goto L43
	}
L43:
	;
	goto L39
L44:
	;
	if v266 != 0 {
		v145 = v266
		goto L37
	} else {
		goto L70
	}
L45:
	;
	v177 = v170
	v178 = v174
	goto L48
L46:
	;
	v174 = int32(1)
	goto L45
L47:
	;
	v174 = int32(0)
	goto L45
L48:
	;
	switch v178 {
	case 0:
		goto L53
	default:
		goto L52
	}
L50:
	;
	v178 = int32(0)
	goto L48
L51:
	;
	goto L44
L52:
	;
	v258 = *(*int32)(unsafe.Add(mBase, uint32(v177)))
	*(*int32)(unsafe.Add(mBase, uint32(v30)+16)) = v258
	if v258 == int32(0) {
		goto L50
	} else {
		goto L69
	}
L53:
	;
	v182 = *(*int32)(unsafe.Add(mBase, uint32(v30)+4))
	if v182 != int32(-1) {
		v221 = v182
		goto L54
	} else {
		goto L55
	}
L54:
	;
	v222 = int32(1)
	v223 = v221 + v222
	*(*int32)(unsafe.Add(mBase, uint32(v30)+4)) = v223
	v225 = int32(0)
	v228 = *(*int32)(unsafe.Add(mBase, uint32(v30)))
	v229 = *(*int32)(unsafe.Add(mBase, uint32(v30)+8))
	v233 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v228+v229+int32(26)))))
	if v233 == int32(255) {
		goto L63
	} else {
		goto L64
	}
L55:
	;
	v186 = *(*int32)(unsafe.Add(mBase, uint32(v30)+8))
	if v186 != 0 {
		v221 = int32(-1)
		goto L54
	} else {
		goto L56
	}
L56:
	;
	v187 = *(*int32)(unsafe.Add(mBase, uint32(v30)))
	v188 = *(*int32)(unsafe.Add(mBase, uint32(v30)+12))
	if v188 == int32(0) {
		goto L58
	} else {
		goto L59
	}
L57:
	;
	v215 = *(*int32)(unsafe.Add(mBase, uint32(v214)+20))
	if v215 != int32(-1) {
		goto L60
	} else {
		goto L61
	}
L58:
	;
	v195 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v187)+16)))
	v196 = int64(*(*int8)(unsafe.Add(mBase, uint32(v187)+27)))
	v197 = int64(*(*int32)(unsafe.Add(mBase, uint32(v187)+8)))
	v198 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v187)+12)))
	v199 = int64(*(*int8)(unsafe.Add(mBase, uint32(v187)+26)))
	v200 = int64(*(*int32)(unsafe.Add(mBase, uint32(v187)+4)))
	v201 = F_wangHash64(m, v200)
	mBase = m.M
	v203 = F_wangHash64(m, v199+v201)
	mBase = m.M
	v205 = F_wangHash64(m, v198+v203)
	mBase = m.M
	v207 = F_wangHash64(m, v197+v205)
	mBase = m.M
	v209 = F_wangHash64(m, v196+v207)
	mBase = m.M
	v211 = F_wangHash64(m, v195+v209)
	mBase = m.M
	*(*int64)(unsafe.Add(mBase, uint32(v30)+24)) = v211
	v213 = *(*int32)(unsafe.Add(mBase, uint32(v30)))
	v214 = v213
	goto L57
L59:
	;
	v191 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v187)+24)))
	v193 = v191 + int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v187)+24)) = uint16(v193)
	v214 = v187
	goto L57
L60:
	;
	v221 = v215 + int32(-1)
	goto L54
L61:
	;
	v218 = *(*int32)(unsafe.Add(mBase, uint32(v30)+4))
	v221 = v218
	goto L54
L62:
	;
	v248 = int32(2)
	v253 = *(*int32)(unsafe.Add(mBase, uint32(v228+v246<<(uint(v248)%32)+int32(4))))
	v177 = v253 + v247<<(uint(v248)%32)
	v178 = int32(1)
	goto L48
L63:
	;
	v237 = v225
	goto L65
L64:
	;
	v237 = v222 << (uint(v233) % 32)
	goto L65
L65:
	;
	if v223 < v237 {
		v246 = v229
		v247 = v223
		goto L62
	} else {
		goto L66
	}
L66:
	;
	if v229 != 0 {
		v266 = v225
		goto L51
	} else {
		goto L67
	}
L67:
	;
	v239 = *(*int32)(unsafe.Add(mBase, uint32(v228)+20))
	if v239 == int32(-1) {
		v266 = v225
		goto L51
	} else {
		goto L68
	}
L68:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v30)+4)) = int64(4294967296)
	v246 = int32(1)
	v247 = int32(0)
	goto L62
L69:
	;
	v262 = *(*int32)(unsafe.Add(mBase, uint32(v258)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v170))) = v262
	v266 = v258
	goto L51
L70:
	;
	goto L38
L71:
	;
	v284 = F_dictGetIterator(m, v23)
	mBase = m.M
	v285 = m.ExcPending
	if v285 != 0 {
		goto L6
	} else {
		goto L72
	}
L72:
	;
	v293 = v284 + int32(20)
	v294 = *(*int32)(unsafe.Add(mBase, uint32(v284)+16))
	if v294 != 0 {
		goto L75
	} else {
		goto L76
	}
L73:
	;
	if v389 == int32(0) {
		goto L3
	} else {
		goto L99
	}
L74:
	;
	v300 = v293
	v301 = v297
	goto L77
L75:
	;
	v297 = int32(1)
	goto L74
L76:
	;
	v297 = int32(0)
	goto L74
L77:
	;
	switch v301 {
	case 0:
		goto L82
	default:
		goto L81
	}
L79:
	;
	v301 = int32(0)
	goto L77
L80:
	;
	goto L73
L81:
	;
	v381 = *(*int32)(unsafe.Add(mBase, uint32(v300)))
	*(*int32)(unsafe.Add(mBase, uint32(v284)+16)) = v381
	if v381 == int32(0) {
		goto L79
	} else {
		goto L98
	}
L82:
	;
	v305 = *(*int32)(unsafe.Add(mBase, uint32(v284)+4))
	if v305 != int32(-1) {
		v344 = v305
		goto L83
	} else {
		goto L84
	}
L83:
	;
	v345 = int32(1)
	v346 = v344 + v345
	*(*int32)(unsafe.Add(mBase, uint32(v284)+4)) = v346
	v348 = int32(0)
	v351 = *(*int32)(unsafe.Add(mBase, uint32(v284)))
	v352 = *(*int32)(unsafe.Add(mBase, uint32(v284)+8))
	v356 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v351+v352+int32(26)))))
	if v356 == int32(255) {
		goto L92
	} else {
		goto L93
	}
L84:
	;
	v309 = *(*int32)(unsafe.Add(mBase, uint32(v284)+8))
	if v309 != 0 {
		v344 = int32(-1)
		goto L83
	} else {
		goto L85
	}
L85:
	;
	v310 = *(*int32)(unsafe.Add(mBase, uint32(v284)))
	v311 = *(*int32)(unsafe.Add(mBase, uint32(v284)+12))
	if v311 == int32(0) {
		goto L87
	} else {
		goto L88
	}
L86:
	;
	v338 = *(*int32)(unsafe.Add(mBase, uint32(v337)+20))
	if v338 != int32(-1) {
		goto L89
	} else {
		goto L90
	}
L87:
	;
	v318 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v310)+16)))
	v319 = int64(*(*int8)(unsafe.Add(mBase, uint32(v310)+27)))
	v320 = int64(*(*int32)(unsafe.Add(mBase, uint32(v310)+8)))
	v321 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v310)+12)))
	v322 = int64(*(*int8)(unsafe.Add(mBase, uint32(v310)+26)))
	v323 = int64(*(*int32)(unsafe.Add(mBase, uint32(v310)+4)))
	v324 = F_wangHash64(m, v323)
	mBase = m.M
	v326 = F_wangHash64(m, v322+v324)
	mBase = m.M
	v328 = F_wangHash64(m, v321+v326)
	mBase = m.M
	v330 = F_wangHash64(m, v320+v328)
	mBase = m.M
	v332 = F_wangHash64(m, v319+v330)
	mBase = m.M
	v334 = F_wangHash64(m, v318+v332)
	mBase = m.M
	*(*int64)(unsafe.Add(mBase, uint32(v284)+24)) = v334
	v336 = *(*int32)(unsafe.Add(mBase, uint32(v284)))
	v337 = v336
	goto L86
L88:
	;
	v314 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v310)+24)))
	v316 = v314 + int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v310)+24)) = uint16(v316)
	v337 = v310
	goto L86
L89:
	;
	v344 = v338 + int32(-1)
	goto L83
L90:
	;
	v341 = *(*int32)(unsafe.Add(mBase, uint32(v284)+4))
	v344 = v341
	goto L83
L91:
	;
	v371 = int32(2)
	v376 = *(*int32)(unsafe.Add(mBase, uint32(v351+v369<<(uint(v371)%32)+int32(4))))
	v300 = v376 + v370<<(uint(v371)%32)
	v301 = int32(1)
	goto L77
L92:
	;
	v360 = v348
	goto L94
L93:
	;
	v360 = v345 << (uint(v356) % 32)
	goto L94
L94:
	;
	if v346 < v360 {
		v369 = v352
		v370 = v346
		goto L91
	} else {
		goto L95
	}
L95:
	;
	if v352 != 0 {
		v389 = v348
		goto L80
	} else {
		goto L96
	}
L96:
	;
	v362 = *(*int32)(unsafe.Add(mBase, uint32(v351)+20))
	if v362 == int32(-1) {
		v389 = v348
		goto L80
	} else {
		goto L97
	}
L97:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v284)+4)) = int64(4294967296)
	v369 = int32(1)
	v370 = int32(0)
	goto L91
L98:
	;
	v385 = *(*int32)(unsafe.Add(mBase, uint32(v381)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v293))) = v385
	v389 = v381
	goto L80
L99:
	;
	v401 = v389
	v404 = int32(0)
	v406 = int64(0)
	goto L100
L100:
	;
	v409 = *(*int64)(unsafe.Add(mBase, uint32(v401)+8))
	goto L103
L101:
	;
	F_dictReleaseIterator(m, v284)
	mBase = m.M
	v522 = m.ExcPending
	if v522 != 0 {
		goto L6
	} else {
		goto L133
	}
L102:
	;
	v421 = v284 + int32(20)
	v422 = *(*int32)(unsafe.Add(mBase, uint32(v284)+16))
	if v422 != 0 {
		goto L108
	} else {
		goto L109
	}
L103:
	;
	if base.Ui64(v409) <= base.Ui64(v406) {
		v412 = v404
		v413 = v406
		goto L102
	} else {
		goto L104
	}
L104:
	;
	v411 = *(*int32)(unsafe.Add(mBase, uint32(v401)))
	goto L105
L105:
	;
	v412 = v411
	v413 = v409
	goto L102
L106:
	;
	if v517 != 0 {
		v401 = v517
		v404 = v412
		v406 = v413
		goto L100
	} else {
		goto L132
	}
L107:
	;
	v428 = v421
	v429 = v425
	goto L110
L108:
	;
	v425 = int32(1)
	goto L107
L109:
	;
	v425 = int32(0)
	goto L107
L110:
	;
	switch v429 {
	case 0:
		goto L115
	default:
		goto L114
	}
L112:
	;
	v429 = int32(0)
	goto L110
L113:
	;
	goto L106
L114:
	;
	v509 = *(*int32)(unsafe.Add(mBase, uint32(v428)))
	*(*int32)(unsafe.Add(mBase, uint32(v284)+16)) = v509
	if v509 == int32(0) {
		goto L112
	} else {
		goto L131
	}
L115:
	;
	v433 = *(*int32)(unsafe.Add(mBase, uint32(v284)+4))
	if v433 != int32(-1) {
		v472 = v433
		goto L116
	} else {
		goto L117
	}
L116:
	;
	v473 = int32(1)
	v474 = v472 + v473
	*(*int32)(unsafe.Add(mBase, uint32(v284)+4)) = v474
	v476 = int32(0)
	v479 = *(*int32)(unsafe.Add(mBase, uint32(v284)))
	v480 = *(*int32)(unsafe.Add(mBase, uint32(v284)+8))
	v484 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v479+v480+int32(26)))))
	if v484 == int32(255) {
		goto L125
	} else {
		goto L126
	}
L117:
	;
	v437 = *(*int32)(unsafe.Add(mBase, uint32(v284)+8))
	if v437 != 0 {
		v472 = int32(-1)
		goto L116
	} else {
		goto L118
	}
L118:
	;
	v438 = *(*int32)(unsafe.Add(mBase, uint32(v284)))
	v439 = *(*int32)(unsafe.Add(mBase, uint32(v284)+12))
	if v439 == int32(0) {
		goto L120
	} else {
		goto L121
	}
L119:
	;
	v466 = *(*int32)(unsafe.Add(mBase, uint32(v465)+20))
	if v466 != int32(-1) {
		goto L122
	} else {
		goto L123
	}
L120:
	;
	v446 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v438)+16)))
	v447 = int64(*(*int8)(unsafe.Add(mBase, uint32(v438)+27)))
	v448 = int64(*(*int32)(unsafe.Add(mBase, uint32(v438)+8)))
	v449 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v438)+12)))
	v450 = int64(*(*int8)(unsafe.Add(mBase, uint32(v438)+26)))
	v451 = int64(*(*int32)(unsafe.Add(mBase, uint32(v438)+4)))
	v452 = F_wangHash64(m, v451)
	mBase = m.M
	v454 = F_wangHash64(m, v450+v452)
	mBase = m.M
	v456 = F_wangHash64(m, v449+v454)
	mBase = m.M
	v458 = F_wangHash64(m, v448+v456)
	mBase = m.M
	v460 = F_wangHash64(m, v447+v458)
	mBase = m.M
	v462 = F_wangHash64(m, v446+v460)
	mBase = m.M
	*(*int64)(unsafe.Add(mBase, uint32(v284)+24)) = v462
	v464 = *(*int32)(unsafe.Add(mBase, uint32(v284)))
	v465 = v464
	goto L119
L121:
	;
	v442 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v438)+24)))
	v444 = v442 + int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v438)+24)) = uint16(v444)
	v465 = v438
	goto L119
L122:
	;
	v472 = v466 + int32(-1)
	goto L116
L123:
	;
	v469 = *(*int32)(unsafe.Add(mBase, uint32(v284)+4))
	v472 = v469
	goto L116
L124:
	;
	v499 = int32(2)
	v504 = *(*int32)(unsafe.Add(mBase, uint32(v479+v497<<(uint(v499)%32)+int32(4))))
	v428 = v504 + v498<<(uint(v499)%32)
	v429 = int32(1)
	goto L110
L125:
	;
	v488 = v476
	goto L127
L126:
	;
	v488 = v473 << (uint(v484) % 32)
	goto L127
L127:
	;
	if v474 < v488 {
		v497 = v480
		v498 = v474
		goto L124
	} else {
		goto L128
	}
L128:
	;
	if v480 != 0 {
		v517 = v476
		goto L113
	} else {
		goto L129
	}
L129:
	;
	v490 = *(*int32)(unsafe.Add(mBase, uint32(v479)+20))
	if v490 == int32(-1) {
		v517 = v476
		goto L113
	} else {
		goto L130
	}
L130:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v284)+4)) = int64(4294967296)
	v497 = int32(1)
	v498 = int32(0)
	goto L124
L131:
	;
	v513 = *(*int32)(unsafe.Add(mBase, uint32(v509)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v421))) = v513
	v517 = v509
	goto L113
L132:
	;
	goto L101
L133:
	;
	v523 = int32(0)
	if v412 == v523 {
		goto L134
	} else {
		goto L135
	}
L134:
	;
	v546 = v413
	v548 = int32(0)
	goto L2
L135:
	;
	v558 = v412
	v559 = v523
	v560 = v413
	v562 = v412
	goto L1
L136:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L137:
	;
	v546 = int64(0)
	v548 = int32(0)
	goto L2
L138:
	;
	F_sdsfree(m, v565)
	mBase = m.M
	v609 = m.ExcPending
	if v609 != 0 {
		goto L6
	} else {
		goto L157
	}
L139:
	;
	v589 = int32(0)
	v591 = int32(1)
	if base.Ui64(v586) < base.Ui64(base.I64_extend_i32_u(int32(base.Ui32(v29+v28+v591)>>(uint(v591)%32))+v591)) {
		v604 = v589
		goto L138
	} else {
		goto L154
	}
L140:
	;
	if base.Ui64(v573) < base.Ui64(v560) {
		goto L148
	} else {
		goto L149
	}
L141:
	;
	if v559 == int32(0) {
		v586 = v560
		v588 = v562
		goto L139
	} else {
		goto L147
	}
L142:
	;
	if v565 == int32(0) {
		goto L141
	} else {
		goto L143
	}
L143:
	;
	v569 = *(*int64)(unsafe.Add(mBase, uint32(v15)+8))
	if v569 != l1 {
		goto L141
	} else {
		goto L144
	}
L144:
	;
	v571 = F_sentinelLeaderIncr(m, v23, v565)
	mBase = m.M
	v572 = m.ExcPending
	if v572 != 0 {
		goto L6
	} else {
		goto L145
	}
L145:
	;
	v573 = base.I64_extend_i32_s(v571)
	v574 = base.B2i32(base.Ui64(v573) <= base.Ui64(v560))
	if v574&v559 == int32(0) {
		goto L140
	} else {
		goto L146
	}
L146:
	;
	v604 = int32(0)
	goto L138
L147:
	;
	v604 = int32(0)
	goto L138
L148:
	;
	v583 = v560
	goto L150
L149:
	;
	v583 = v573
	goto L150
L150:
	;
	if base.Ui64(v573) <= base.Ui64(v560) {
		goto L151
	} else {
		goto L152
	}
L151:
	;
	v584 = v562
	goto L153
L152:
	;
	v584 = v565
	goto L153
L153:
	;
	v586 = v583
	v588 = v584
	goto L139
L154:
	;
	v599 = int64(*(*uint32)(unsafe.Add(mBase, uint32(l0)+152)))
	if base.Ui64(v586) < base.Ui64(v599) {
		v604 = v589
		goto L138
	} else {
		goto L155
	}
L155:
	;
	v601 = F_sdsnew(m, v588)
	mBase = m.M
	v602 = m.ExcPending
	if v602 != 0 {
		goto L6
	} else {
		goto L156
	}
L156:
	;
	v604 = v601
	goto L138
L157:
	;
	F_dictRelease(m, v23)
	mBase = m.M
	v611 = m.ExcPending
	if v611 != 0 {
		goto L6
	} else {
		goto L158
	}
L158:
	;
	m.G0 = v15 + int32(16)
	return v604
}
func F_sentinelGetPrimaryByName(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	v3 = F_sdsnew(m, l0)
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		v8 = *(*int32)(unsafe.Add(mBase, _consts[733]))
		v9 = F_dictFetchValue(m, v8, v3)
		mBase = m.M
		v10 = m.ExcPending
		if v10 != 0 {
			return int32(0)
		} else {
			F_sdsfree(m, v3)
			mBase = m.M
			v12 = m.ExcPending
			if v12 != 0 {
				return int32(0)
			} else {
				return v9
			}
		}
	}
}
func F_sentinelHandleDictOfValkeyInstances(m *base.Module, l0 int32) {
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
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
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
	var v54 int64
	_ = v54
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
	var v65 int64
	_ = v65
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v87 int32
	_ = v87
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v107 int32
	_ = v107
	var v112 int32
	_ = v112
	var v116 int32
	_ = v116
	var v120 int32
	_ = v120
	var v128 int32
	_ = v128
	var v138 int32
	_ = v138
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v146 int32
	_ = v146
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v163 int32
	_ = v163
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v204 int32
	_ = v204
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v212 int32
	_ = v212
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v221 int32
	_ = v221
	var v223 int32
	_ = v223
	var v225 int64
	_ = v225
	var v226 int64
	_ = v226
	var v227 int64
	_ = v227
	var v228 int64
	_ = v228
	var v229 int64
	_ = v229
	var v230 int64
	_ = v230
	var v231 int64
	_ = v231
	var v233 int64
	_ = v233
	var v235 int64
	_ = v235
	var v237 int64
	_ = v237
	var v239 int64
	_ = v239
	var v241 int64
	_ = v241
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v248 int32
	_ = v248
	var v251 int32
	_ = v251
	var v252 int32
	_ = v252
	var v253 int32
	_ = v253
	var v255 int32
	_ = v255
	var v258 int32
	_ = v258
	var v259 int32
	_ = v259
	var v263 int32
	_ = v263
	var v267 int32
	_ = v267
	var v269 int32
	_ = v269
	var v276 int32
	_ = v276
	var v277 int32
	_ = v277
	var v278 int32
	_ = v278
	var v283 int32
	_ = v283
	var v288 int32
	_ = v288
	var v292 int32
	_ = v292
	var v296 int32
	_ = v296
	var v311 int32
	_ = v311
	v11 = m.G0
	v13 = v11 - int32(32)
	m.G0 = v13
	v15 = F_dictGetIterator(m, l0)
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	F_dictReleaseIterator(m, v15)
	mBase = m.M
	v311 = m.ExcPending
	if v311 != 0 {
		goto L2
	} else {
		goto L72
	}
L2:
	;
	return
L3:
	;
	v24 = v15 + int32(20)
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v15)+16))
	if v25 != 0 {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	if v120 == int32(0) {
		goto L1
	} else {
		goto L30
	}
L5:
	;
	v31 = v24
	v32 = v28
	goto L8
L6:
	;
	v28 = int32(1)
	goto L5
L7:
	;
	v28 = int32(0)
	goto L5
L8:
	;
	switch v32 {
	case 0:
		goto L13
	default:
		goto L12
	}
L10:
	;
	v32 = int32(0)
	goto L8
L11:
	;
	goto L4
L12:
	;
	v112 = *(*int32)(unsafe.Add(mBase, uint32(v31)))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+16)) = v112
	if v112 == int32(0) {
		goto L10
	} else {
		goto L29
	}
L13:
	;
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v15)+4))
	if v36 != int32(-1) {
		v75 = v36
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v76 = int32(1)
	v77 = v75 + v76
	*(*int32)(unsafe.Add(mBase, uint32(v15)+4)) = v77
	v79 = int32(0)
	v82 = *(*int32)(unsafe.Add(mBase, uint32(v15)))
	v83 = *(*int32)(unsafe.Add(mBase, uint32(v15)+8))
	v87 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v82+v83+int32(26)))))
	if v87 == int32(255) {
		goto L23
	} else {
		goto L24
	}
L15:
	;
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v15)+8))
	if v40 != 0 {
		v75 = int32(-1)
		goto L14
	} else {
		goto L16
	}
L16:
	;
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v15)))
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v15)+12))
	if v42 == int32(0) {
		goto L18
	} else {
		goto L19
	}
L17:
	;
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v68)+20))
	if v69 != int32(-1) {
		goto L20
	} else {
		goto L21
	}
L18:
	;
	v49 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v41)+16)))
	v50 = int64(*(*int8)(unsafe.Add(mBase, uint32(v41)+27)))
	v51 = int64(*(*int32)(unsafe.Add(mBase, uint32(v41)+8)))
	v52 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v41)+12)))
	v53 = int64(*(*int8)(unsafe.Add(mBase, uint32(v41)+26)))
	v54 = int64(*(*int32)(unsafe.Add(mBase, uint32(v41)+4)))
	v55 = F_wangHash64(m, v54)
	mBase = m.M
	v57 = F_wangHash64(m, v53+v55)
	mBase = m.M
	v59 = F_wangHash64(m, v52+v57)
	mBase = m.M
	v61 = F_wangHash64(m, v51+v59)
	mBase = m.M
	v63 = F_wangHash64(m, v50+v61)
	mBase = m.M
	v65 = F_wangHash64(m, v49+v63)
	mBase = m.M
	*(*int64)(unsafe.Add(mBase, uint32(v15)+24)) = v65
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v15)))
	v68 = v67
	goto L17
L19:
	;
	v45 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v41)+24)))
	v47 = v45 + int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v41)+24)) = uint16(v47)
	v68 = v41
	goto L17
L20:
	;
	v75 = v69 + int32(-1)
	goto L14
L21:
	;
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v15)+4))
	v75 = v72
	goto L14
L22:
	;
	v102 = int32(2)
	v107 = *(*int32)(unsafe.Add(mBase, uint32(v82+v100<<(uint(v102)%32)+int32(4))))
	v31 = v107 + v101<<(uint(v102)%32)
	v32 = int32(1)
	goto L8
L23:
	;
	v91 = v79
	goto L25
L24:
	;
	v91 = v76 << (uint(v87) % 32)
	goto L25
L25:
	;
	if v77 < v91 {
		v100 = v83
		v101 = v77
		goto L22
	} else {
		goto L26
	}
L26:
	;
	if v83 != 0 {
		v120 = v79
		goto L11
	} else {
		goto L27
	}
L27:
	;
	v93 = *(*int32)(unsafe.Add(mBase, uint32(v82)+20))
	if v93 == int32(-1) {
		v120 = v79
		goto L11
	} else {
		goto L28
	}
L28:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v15)+4)) = int64(4294967296)
	v100 = int32(1)
	v101 = int32(0)
	goto L22
L29:
	;
	v116 = *(*int32)(unsafe.Add(mBase, uint32(v112)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v24))) = v116
	v120 = v112
	goto L11
L30:
	;
	v128 = v120
	goto L31
L31:
	;
	v138 = *(*int32)(unsafe.Add(mBase, uint32(v128)+8))
	goto L33
L32:
	;
	goto L1
L33:
	;
	F_sentinelHandleValkeyInstance(m, v138)
	mBase = m.M
	v140 = m.ExcPending
	if v140 != 0 {
		goto L2
	} else {
		goto L34
	}
L34:
	;
	v141 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v138))))
	if v141&int32(1) == int32(0) {
		goto L35
	} else {
		goto L36
	}
L35:
	;
	v200 = v15 + int32(20)
	v201 = *(*int32)(unsafe.Add(mBase, uint32(v15)+16))
	if v201 != 0 {
		goto L47
	} else {
		goto L48
	}
L36:
	;
	v146 = *(*int32)(unsafe.Add(mBase, uint32(v138)+148))
	F_sentinelHandleDictOfValkeyInstances(m, v146)
	mBase = m.M
	v148 = m.ExcPending
	if v148 != 0 {
		goto L2
	} else {
		goto L37
	}
L37:
	;
	v149 = *(*int32)(unsafe.Add(mBase, uint32(v138)+144))
	F_sentinelHandleDictOfValkeyInstances(m, v149)
	mBase = m.M
	v151 = m.ExcPending
	if v151 != 0 {
		goto L2
	} else {
		goto L38
	}
L38:
	;
	v152 = *(*int32)(unsafe.Add(mBase, uint32(v138)+240))
	if v152 != int32(6) {
		goto L35
	} else {
		goto L39
	}
L39:
	;
	v155 = *(*int32)(unsafe.Add(mBase, uint32(v138)+280))
	if v155 != 0 {
		goto L40
	} else {
		goto L41
	}
L40:
	;
	v156 = v155
	goto L42
L41:
	;
	v156 = v138
	goto L42
L42:
	;
	v157 = *(*int32)(unsafe.Add(mBase, uint32(v156)+24))
	v158 = int32(0)
	v159 = *(*int32)(unsafe.Add(mBase, _consts[734]))
	v163 = base.B2i32(v159 == v158) << (uint(int32(2)) % 32)
	v165 = *(*int32)(unsafe.Add(mBase, uint32(v157+v163)))
	v166 = *(*int32)(unsafe.Add(mBase, uint32(v138)+24))
	v168 = *(*int32)(unsafe.Add(mBase, uint32(v166+v163)))
	v169 = *(*int32)(unsafe.Add(mBase, uint32(v166)+8))
	v170 = *(*int32)(unsafe.Add(mBase, uint32(v138)+4))
	v171 = *(*int32)(unsafe.Add(mBase, uint32(v157)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v13+int32(16)))) = v171
	*(*int32)(unsafe.Add(mBase, uint32(v13))) = v170
	*(*int32)(unsafe.Add(mBase, uint32(v13)+4)) = v168
	*(*int32)(unsafe.Add(mBase, uint32(v13)+8)) = v169
	*(*int32)(unsafe.Add(mBase, uint32(v13)+12)) = v165
	F_sentinelEvent(m, int32(3), int32(_a1552), v138, int32(_a1553), v13)
	mBase = m.M
	v181 = m.ExcPending
	if v181 != 0 {
		goto L2
	} else {
		goto L43
	}
L43:
	;
	v182 = *(*int32)(unsafe.Add(mBase, uint32(v156)+24))
	v183 = *(*int32)(unsafe.Add(mBase, uint32(v182)))
	v184 = *(*int32)(unsafe.Add(mBase, uint32(v182)+8))
	v185 = F_sentinelResetPrimaryAndChangeAddress(m, v138, v183, v184)
	mBase = m.M
	v186 = m.ExcPending
	if v186 != 0 {
		goto L2
	} else {
		goto L44
	}
L44:
	;
	goto L35
L45:
	;
	if v296 != 0 {
		v128 = v296
		goto L31
	} else {
		goto L71
	}
L46:
	;
	v207 = v200
	v208 = v204
	goto L49
L47:
	;
	v204 = int32(1)
	goto L46
L48:
	;
	v204 = int32(0)
	goto L46
L49:
	;
	switch v208 {
	case 0:
		goto L54
	default:
		goto L53
	}
L51:
	;
	v208 = int32(0)
	goto L49
L52:
	;
	goto L45
L53:
	;
	v288 = *(*int32)(unsafe.Add(mBase, uint32(v207)))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+16)) = v288
	if v288 == int32(0) {
		goto L51
	} else {
		goto L70
	}
L54:
	;
	v212 = *(*int32)(unsafe.Add(mBase, uint32(v15)+4))
	if v212 != int32(-1) {
		v251 = v212
		goto L55
	} else {
		goto L56
	}
L55:
	;
	v252 = int32(1)
	v253 = v251 + v252
	*(*int32)(unsafe.Add(mBase, uint32(v15)+4)) = v253
	v255 = int32(0)
	v258 = *(*int32)(unsafe.Add(mBase, uint32(v15)))
	v259 = *(*int32)(unsafe.Add(mBase, uint32(v15)+8))
	v263 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v258+v259+int32(26)))))
	if v263 == int32(255) {
		goto L64
	} else {
		goto L65
	}
L56:
	;
	v216 = *(*int32)(unsafe.Add(mBase, uint32(v15)+8))
	if v216 != 0 {
		v251 = int32(-1)
		goto L55
	} else {
		goto L57
	}
L57:
	;
	v217 = *(*int32)(unsafe.Add(mBase, uint32(v15)))
	v218 = *(*int32)(unsafe.Add(mBase, uint32(v15)+12))
	if v218 == int32(0) {
		goto L59
	} else {
		goto L60
	}
L58:
	;
	v245 = *(*int32)(unsafe.Add(mBase, uint32(v244)+20))
	if v245 != int32(-1) {
		goto L61
	} else {
		goto L62
	}
L59:
	;
	v225 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v217)+16)))
	v226 = int64(*(*int8)(unsafe.Add(mBase, uint32(v217)+27)))
	v227 = int64(*(*int32)(unsafe.Add(mBase, uint32(v217)+8)))
	v228 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v217)+12)))
	v229 = int64(*(*int8)(unsafe.Add(mBase, uint32(v217)+26)))
	v230 = int64(*(*int32)(unsafe.Add(mBase, uint32(v217)+4)))
	v231 = F_wangHash64(m, v230)
	mBase = m.M
	v233 = F_wangHash64(m, v229+v231)
	mBase = m.M
	v235 = F_wangHash64(m, v228+v233)
	mBase = m.M
	v237 = F_wangHash64(m, v227+v235)
	mBase = m.M
	v239 = F_wangHash64(m, v226+v237)
	mBase = m.M
	v241 = F_wangHash64(m, v225+v239)
	mBase = m.M
	*(*int64)(unsafe.Add(mBase, uint32(v15)+24)) = v241
	v243 = *(*int32)(unsafe.Add(mBase, uint32(v15)))
	v244 = v243
	goto L58
L60:
	;
	v221 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v217)+24)))
	v223 = v221 + int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v217)+24)) = uint16(v223)
	v244 = v217
	goto L58
L61:
	;
	v251 = v245 + int32(-1)
	goto L55
L62:
	;
	v248 = *(*int32)(unsafe.Add(mBase, uint32(v15)+4))
	v251 = v248
	goto L55
L63:
	;
	v278 = int32(2)
	v283 = *(*int32)(unsafe.Add(mBase, uint32(v258+v276<<(uint(v278)%32)+int32(4))))
	v207 = v283 + v277<<(uint(v278)%32)
	v208 = int32(1)
	goto L49
L64:
	;
	v267 = v255
	goto L66
L65:
	;
	v267 = v252 << (uint(v263) % 32)
	goto L66
L66:
	;
	if v253 < v267 {
		v276 = v259
		v277 = v253
		goto L63
	} else {
		goto L67
	}
L67:
	;
	if v259 != 0 {
		v296 = v255
		goto L52
	} else {
		goto L68
	}
L68:
	;
	v269 = *(*int32)(unsafe.Add(mBase, uint32(v258)+20))
	if v269 == int32(-1) {
		v296 = v255
		goto L52
	} else {
		goto L69
	}
L69:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v15)+4)) = int64(4294967296)
	v276 = int32(1)
	v277 = int32(0)
	goto L63
L70:
	;
	v292 = *(*int32)(unsafe.Add(mBase, uint32(v288)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v200))) = v292
	v296 = v288
	goto L52
L71:
	;
	goto L32
L72:
	;
	m.G0 = v13 + int32(32)
	return
}
func F_sentinelInstanceMapCommand(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
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
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	v5 = F_sdsnew(m, l1)
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		return int32(0)
	} else {
		v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+192))
		if v9 != 0 {
			v10 = v9
		} else {
			v10 = l0
		}
		v11 = *(*int32)(unsafe.Add(mBase, uint32(v10)+104))
		v12 = F_dictFetchValue(m, v11, v5)
		mBase = m.M
		v13 = m.ExcPending
		if v13 != 0 {
			return int32(0)
		} else {
			F_sdsfree(m, v5)
			mBase = m.M
			v15 = m.ExcPending
			if v15 != 0 {
				return int32(0)
			} else {
				if v12 != 0 {
					v16 = v12
				} else {
					v16 = l1
				}
				return v16
			}
		}
	}
}
func F_sentinelIsRunning(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v10 int32
	_ = v10
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
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
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v62 int32
	_ = v62
	var v69 int32
	_ = v69
	var v73 int32
	_ = v73
	var v81 int32
	_ = v81
	var v85 int32
	_ = v85
	v4 = m.G0
	v6 = v4 - int32(32)
	m.G0 = v6
	v10 = int32(0)
	goto L3
L1:
	;
	F_sentinelGenerateInitialMonitorEvents(m)
	mBase = m.M
	v85 = m.ExcPending
	if v85 != 0 {
		goto L11
	} else {
		goto L24
	}
L2:
	;
	v73 = *(*int32)(unsafe.Add(mBase, _consts[6]))
	if int32(2) < v73 {
		goto L1
	} else {
		goto L22
	}
L3:
	;
	v14 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+uint32(_consts[724]))))
	if v14 != 0 {
		goto L2
	} else {
		goto L5
	}
L4:
	;
	F_getRandomHexChars(m, int32(_a1336), int32(40))
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
		goto L11
	} else {
		goto L12
	}
L5:
	;
	v17 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+uint32(_consts[725]))))
	if v17 != 0 {
		goto L2
	} else {
		goto L6
	}
L6:
	;
	v20 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+uint32(_consts[726]))))
	if v20 != 0 {
		goto L2
	} else {
		goto L7
	}
L7:
	;
	v23 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+uint32(_consts[727]))))
	if v23 != 0 {
		goto L2
	} else {
		goto L8
	}
L8:
	;
	v26 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+uint32(_consts[728]))))
	if v26 != 0 {
		goto L2
	} else {
		goto L9
	}
L9:
	;
	v28 = v10 + int32(5)
	if v28 != int32(40) {
		v10 = v28
		goto L3
	} else {
		goto L10
	}
L10:
	;
	goto L4
L11:
	;
	return
L12:
	;
	v35 = int32(_a44)
	v36 = *(*int32)(unsafe.Add(mBase, _consts[219]))
	*(*int32)(unsafe.Add(mBase, _consts[219])) = int32(10)
	v41 = *(*int32)(unsafe.Add(mBase, _consts[392]))
	v43 = F_rewriteConfig(m, v41, int32(0))
	mBase = m.M
	v44 = m.ExcPending
	if v44 != 0 {
		goto L11
	} else {
		goto L13
	}
L13:
	;
	v45 = int32(_a44)
	*(*int32)(unsafe.Add(mBase, _consts[219])) = v36
	v48 = *(*int32)(unsafe.Add(mBase, _consts[6]))
	if v43 != int32(-1) {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	if int32(2) < v48 {
		goto L1
	} else {
		goto L20
	}
L15:
	;
	if int32(3) < v48 {
		goto L1
	} else {
		goto L16
	}
L16:
	;
	goto L17
L17:
	;
	v54 = *(*int32)(unsafe.Add(mBase, _consts[5]))
	v55 = F___strerror_l(m, v54, v54)
	mBase = m.M
	goto L18
L18:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v6)+16)) = v55
	F__serverLog(m, int32(3), int32(_a1337), v6+int32(16))
	mBase = m.M
	v62 = m.ExcPending
	if v62 != 0 {
		goto L11
	} else {
		goto L19
	}
L19:
	;
	goto L2
L20:
	;
	F__serverLog(m, int32(2), int32(_a1338), int32(0))
	mBase = m.M
	v69 = m.ExcPending
	if v69 != 0 {
		goto L11
	} else {
		goto L21
	}
L21:
	;
	goto L2
L22:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v6))) = int32(_a1336)
	F__serverLog(m, int32(2), int32(_a1339), v6)
	mBase = m.M
	v81 = m.ExcPending
	if v81 != 0 {
		goto L11
	} else {
		goto L23
	}
L23:
	;
	goto L1
L24:
	;
	m.G0 = v6 + int32(32)
	return
}
func F_sentinelKillTimedoutScripts(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v1 int32
	_ = v1
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int64
	_ = v9
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v30 int32
	_ = v30
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v44 int64
	_ = v44
	var v47 int64
	_ = v47
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v77 int32
	_ = v77
	v1 = int32(0)
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v9 = F_mstime(m)
	mBase = m.M
	v11 = *(*int32)(unsafe.Add(mBase, _consts[730]))
	v13 = v7 + int32(8)
	v14 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
	*(*int32)(unsafe.Add(mBase, uint32(v13)+4)) = v1
	*(*int32)(unsafe.Add(mBase, uint32(v13))) = v14
	goto L1
L1:
	;
	v19 = v7 + int32(8)
	v21 = *(*int32)(unsafe.Add(mBase, uint32(v19)))
	if v21 == int32(0) {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	m.G0 = v7 + int32(16)
	return
L3:
	;
	if v21 == int32(0) {
		goto L2
	} else {
		goto L6
	}
L4:
	;
	goto L3
L5:
	;
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v19)+4))
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v21+base.B2i32(v24 == int32(0))<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v19))) = v30
	goto L4
L6:
	;
	v36 = v21
	goto L7
L7:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v36)+8))
	v39 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v38))))
	if v39&int32(1) == int32(0) {
		goto L9
	} else {
		goto L10
	}
L8:
	;
	goto L2
L9:
	;
	v66 = v7 + int32(8)
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v66)))
	if v68 == int32(0) {
		goto L16
	} else {
		goto L17
	}
L10:
	;
	v44 = *(*int64)(unsafe.Add(mBase, uint32(v38)+16))
	v47 = *(*int64)(unsafe.Add(mBase, _consts[732]))
	if v9-v44 <= v47 {
		goto L9
	} else {
		goto L11
	}
L11:
	;
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v38)+8))
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v49)))
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v38)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v7)+4)) = v51
	*(*int32)(unsafe.Add(mBase, uint32(v7))) = v50
	F_sentinelEvent(m, int32(3), int32(_a1345), int32(0), int32(_a1346), v7)
	mBase = m.M
	v59 = m.ExcPending
	if v59 != 0 {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	return
L13:
	;
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v38)+24))
	v62 = F_kill(m, v60, int32(9))
	mBase = m.M
	v63 = m.ExcPending
	if v63 != 0 {
		goto L12
	} else {
		goto L14
	}
L14:
	;
	goto L9
L15:
	;
	if v68 != 0 {
		v36 = v68
		goto L7
	} else {
		goto L18
	}
L16:
	;
	goto L15
L17:
	;
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v66)+4))
	v77 = *(*int32)(unsafe.Add(mBase, uint32(v68+base.B2i32(v71 == int32(0))<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v66))) = v77
	goto L16
L18:
	;
	goto L8
}
func F_sentinelLinkEstablishedCallback(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	if l1 == int32(0) {
	} else {
		v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+212))
		if v5 == int32(0) {
		} else {
			v8 = *(*int32)(unsafe.Add(mBase, uint32(v5)+16))
			if v8 != l0 {
				*(*int32)(unsafe.Add(mBase, uint32(v5)+12)) = int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v5)+16)) = int32(0)
			}
			*(*int32)(unsafe.Add(mBase, uint32(v5)+4)) = int32(1)
		}
	}
	return
}
func F_sentinelPendingScriptsCommand(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v37 int32
	_ = v37
	var v43 int32
	_ = v43
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v59 int32
	_ = v59
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v92 int32
	_ = v92
	var v102 int32
	_ = v102
	var v105 int32
	_ = v105
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
	var v113 int32
	_ = v113
	var v114 int64
	_ = v114
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v124 int32
	_ = v124
	var v125 int64
	_ = v125
	var v126 int64
	_ = v126
	var v128 int64
	_ = v128
	var v134 int64
	_ = v134
	var v136 int64
	_ = v136
	var v139 int32
	_ = v139
	var v140 int64
	_ = v140
	var v143 int64
	_ = v143
	var v144 int64
	_ = v144
	var v146 int32
	_ = v146
	var v149 int32
	_ = v149
	var v150 int64
	_ = v150
	var v152 int32
	_ = v152
	var v154 int32
	_ = v154
	var v156 int32
	_ = v156
	var v159 int32
	_ = v159
	var v165 int32
	_ = v165
	v8 = m.G0
	v10 = v8 - int32(16)
	m.G0 = v10
	v13 = *(*int32)(unsafe.Add(mBase, _consts[730]))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(v13)+20))
	F_addReplyArrayLen(m, l0, v14)
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v17 = int32(0)
	v18 = *(*int32)(unsafe.Add(mBase, _consts[730]))
	v20 = v10 + int32(8)
	v21 = *(*int32)(unsafe.Add(mBase, uint32(v18)))
	*(*int32)(unsafe.Add(mBase, uint32(v20)+4)) = v17
	*(*int32)(unsafe.Add(mBase, uint32(v20))) = v21
	goto L3
L3:
	;
	v26 = v10 + int32(8)
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v26)))
	if v28 == int32(0) {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	m.G0 = v10 + int32(16)
	return
L5:
	;
	if v28 == int32(0) {
		goto L4
	} else {
		goto L8
	}
L6:
	;
	goto L5
L7:
	;
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v26)+4))
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v28+base.B2i32(v31 == int32(0))<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v26))) = v37
	goto L6
L8:
	;
	v43 = v28
	goto L9
L9:
	;
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v43)+8))
	F_addReplyMapLen(m, l0, int32(5))
	mBase = m.M
	v51 = m.ExcPending
	if v51 != 0 {
		goto L1
	} else {
		goto L11
	}
L10:
	;
	goto L4
L11:
	;
	F_addReplyBulkCString(m, l0, int32(_a1347))
	mBase = m.M
	v54 = m.ExcPending
	if v54 != 0 {
		goto L1
	} else {
		goto L12
	}
L12:
	;
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v48)+8))
	v59 = int32(0)
	goto L13
L13:
	;
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v55+v59<<(uint(int32(2))%32))))
	if v69 != 0 {
		v59 = v59 + int32(1)
		goto L13
	} else {
		goto L15
	}
L14:
	;
	F_addReplyArrayLen(m, l0, v59)
	mBase = m.M
	v71 = m.ExcPending
	if v71 != 0 {
		goto L1
	} else {
		goto L16
	}
L15:
	;
	goto L14
L16:
	;
	v72 = int32(0)
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v48)+8))
	v74 = *(*int32)(unsafe.Add(mBase, uint32(v73)))
	if v74 == v72 {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	F_addReplyBulkCString(m, l0, int32(_a809))
	mBase = m.M
	v102 = m.ExcPending
	if v102 != 0 {
		goto L1
	} else {
		goto L23
	}
L18:
	;
	v79 = v72
	v82 = v74
	goto L19
L19:
	;
	F_addReplyBulkCString(m, l0, v82)
	mBase = m.M
	v85 = m.ExcPending
	if v85 != 0 {
		goto L1
	} else {
		goto L21
	}
L20:
	;
	goto L17
L21:
	;
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v48)+8))
	v88 = v79 + int32(1)
	v92 = *(*int32)(unsafe.Add(mBase, uint32(v86+v88<<(uint(int32(2))%32))))
	if v92 != 0 {
		v79 = v88
		v82 = v92
		goto L19
	} else {
		goto L22
	}
L22:
	;
	goto L20
L23:
	;
	v105 = *(*int32)(unsafe.Add(mBase, uint32(v48)))
	if v105&int32(1) != 0 {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	v108 = int32(_a1348)
	goto L26
L25:
	;
	v108 = int32(_a1349)
	goto L26
L26:
	;
	F_addReplyBulkCString(m, l0, v108)
	mBase = m.M
	v110 = m.ExcPending
	if v110 != 0 {
		goto L1
	} else {
		goto L27
	}
L27:
	;
	F_addReplyBulkCString(m, l0, int32(_a1350))
	mBase = m.M
	v113 = m.ExcPending
	if v113 != 0 {
		goto L1
	} else {
		goto L28
	}
L28:
	;
	v114 = int64(*(*int32)(unsafe.Add(mBase, uint32(v48)+24)))
	F_addReplyBulkLongLong(m, l0, v114)
	mBase = m.M
	v116 = m.ExcPending
	if v116 != 0 {
		goto L1
	} else {
		goto L29
	}
L29:
	;
	v117 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v48))))
	if v117&int32(1) == int32(0) {
		goto L31
	} else {
		goto L32
	}
L30:
	;
	F_addReplyBulkLongLong(m, l0, v144)
	mBase = m.M
	v146 = m.ExcPending
	if v146 != 0 {
		goto L1
	} else {
		goto L41
	}
L31:
	;
	v128 = *(*int64)(unsafe.Add(mBase, uint32(v48)+16))
	if base.B2i32(v128 == int64(0)) == int32(0) {
		goto L35
	} else {
		goto L36
	}
L32:
	;
	F_addReplyBulkCString(m, l0, int32(_a1351))
	mBase = m.M
	v124 = m.ExcPending
	if v124 != 0 {
		goto L1
	} else {
		goto L33
	}
L33:
	;
	v125 = F_mstime(m)
	mBase = m.M
	v126 = *(*int64)(unsafe.Add(mBase, uint32(v48)+16))
	v144 = v125 - v126
	goto L30
L34:
	;
	F_addReplyBulkCString(m, l0, int32(_a1352))
	mBase = m.M
	v139 = m.ExcPending
	if v139 != 0 {
		goto L1
	} else {
		goto L37
	}
L35:
	;
	v134 = F_mstime(m)
	mBase = m.M
	v136 = v128 - v134
	goto L34
L36:
	;
	v136 = int64(0)
	goto L34
L37:
	;
	v140 = int64(0)
	if v140 < v136 {
		goto L38
	} else {
		goto L39
	}
L38:
	;
	v143 = v136
	goto L40
L39:
	;
	v143 = v140
	goto L40
L40:
	;
	v144 = v143
	goto L30
L41:
	;
	F_addReplyBulkCString(m, l0, int32(_a1353))
	mBase = m.M
	v149 = m.ExcPending
	if v149 != 0 {
		goto L1
	} else {
		goto L42
	}
L42:
	;
	v150 = int64(*(*int32)(unsafe.Add(mBase, uint32(v48)+4)))
	F_addReplyBulkLongLong(m, l0, v150)
	mBase = m.M
	v152 = m.ExcPending
	if v152 != 0 {
		goto L1
	} else {
		goto L43
	}
L43:
	;
	v154 = v10 + int32(8)
	v156 = *(*int32)(unsafe.Add(mBase, uint32(v154)))
	if v156 == int32(0) {
		goto L45
	} else {
		goto L46
	}
L44:
	;
	if v156 != 0 {
		v43 = v156
		goto L9
	} else {
		goto L47
	}
L45:
	;
	goto L44
L46:
	;
	v159 = *(*int32)(unsafe.Add(mBase, uint32(v154)+4))
	v165 = *(*int32)(unsafe.Add(mBase, uint32(v156+base.B2i32(v159 == int32(0))<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v154))) = v165
	goto L45
L47:
	;
	goto L10
}
func F_sentinelPingReplyCallback(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v20 int32
	_ = v20
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v40 int32
	_ = v40
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v65 int32
	_ = v65
	var v75 int32
	_ = v75
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v89 int32
	_ = v89
	var v94 int32
	_ = v94
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v104 int32
	_ = v104
	var v106 int32
	_ = v106
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v114 int32
	_ = v114
	var v124 int32
	_ = v124
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
	var v138 int32
	_ = v138
	var v143 int32
	_ = v143
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v153 int32
	_ = v153
	var v155 int32
	_ = v155
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v163 int32
	_ = v163
	var v171 int64
	_ = v171
	var v175 int32
	_ = v175
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
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
	var v195 int32
	_ = v195
	var v200 int32
	_ = v200
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v210 int32
	_ = v210
	var v212 int32
	_ = v212
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v220 int32
	_ = v220
	var v230 int32
	_ = v230
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
	var v244 int32
	_ = v244
	var v249 int32
	_ = v249
	var v252 int32
	_ = v252
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v259 int32
	_ = v259
	var v261 int32
	_ = v261
	var v264 int32
	_ = v264
	var v265 int32
	_ = v265
	var v269 int32
	_ = v269
	var v277 int32
	_ = v277
	var v282 int32
	_ = v282
	var v283 int32
	_ = v283
	var v285 int32
	_ = v285
	var v286 int32
	_ = v286
	var v290 int32
	_ = v290
	var v291 int32
	_ = v291
	var v292 int32
	_ = v292
	var v293 int32
	_ = v293
	var v298 int32
	_ = v298
	var v301 int32
	_ = v301
	var v308 int64
	_ = v308
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	if l1 == int32(0) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	m.G0 = v9 + int32(16)
	return
L2:
	;
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+212))
	if v13 == int32(0) {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v16 = *(*int32)(unsafe.Add(mBase, uint32(v13)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v13)+8)) = v16 + int32(-1)
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	if base.Ui32(int32(1)) < base.Ui32(v20+int32(-5)) {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v308 = F_mstime(m)
	mBase = m.M
	*(*int64)(unsafe.Add(mBase, uint32(v13)+72)) = v308
	goto L1
L5:
	;
	v25 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
	v26 = int32(_a1413)
	goto L10
L6:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v301
	goto L4
L7:
	;
	v230 = int32(_a1414)
	goto L67
L8:
	;
	v171 = F_mstime(m)
	mBase = m.M
	*(*int64)(unsafe.Add(mBase, uint32(v13)+56)) = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v13)+48)) = v171
	v175 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	if v175&int32(8192) == int32(0) {
		goto L4
	} else {
		goto L51
	}
L9:
	;
	if v60-v65 == int32(0) {
		goto L8
	} else {
		goto L22
	}
L10:
	;
	v31 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v25))))
	if v31 != 0 {
		goto L13
	} else {
		goto L14
	}
L12:
	;
	v65 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v61))))
	goto L9
L13:
	;
	v33 = v25
	v34 = v26
	v35 = int32(4)
	v36 = v31
	goto L16
L14:
	;
	v60 = int32(0)
	v61 = v26
	goto L12
L15:
	;
	v60 = v57 & int32(255)
	v61 = v55
	goto L12
L16:
	;
	v40 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v34))))
	if v36&int32(255) != v40 {
		v55 = v34
		v57 = v36
		goto L15
	} else {
		goto L18
	}
L17:
	;
	v55 = v49
	v57 = int32(0)
	goto L15
L18:
	;
	if v40 == int32(0) {
		v55 = v34
		v57 = v36
		goto L15
	} else {
		goto L19
	}
L19:
	;
	v45 = v35 + int32(-1)
	if v45 == int32(0) {
		v55 = v34
		v57 = v36
		goto L15
	} else {
		goto L20
	}
L20:
	;
	v48 = int32(1)
	v49 = v34 + v48
	v50 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v33)+1)))
	if v50 != 0 {
		v33 = v33 + v48
		v34 = v49
		v35 = v45
		v36 = v50
		goto L16
	} else {
		goto L21
	}
L21:
	;
	goto L17
L22:
	;
	v75 = int32(_a1415)
	goto L24
L23:
	;
	if v109-v114 == int32(0) {
		goto L8
	} else {
		goto L36
	}
L24:
	;
	v80 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v25))))
	if v80 != 0 {
		goto L27
	} else {
		goto L28
	}
L26:
	;
	v114 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v110))))
	goto L23
L27:
	;
	v82 = v25
	v83 = v75
	v84 = int32(7)
	v85 = v80
	goto L30
L28:
	;
	v109 = int32(0)
	v110 = v75
	goto L26
L29:
	;
	v109 = v106 & int32(255)
	v110 = v104
	goto L26
L30:
	;
	v89 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v83))))
	if v85&int32(255) != v89 {
		v104 = v83
		v106 = v85
		goto L29
	} else {
		goto L32
	}
L31:
	;
	v104 = v98
	v106 = int32(0)
	goto L29
L32:
	;
	if v89 == int32(0) {
		v104 = v83
		v106 = v85
		goto L29
	} else {
		goto L33
	}
L33:
	;
	v94 = v84 + int32(-1)
	if v94 == int32(0) {
		v104 = v83
		v106 = v85
		goto L29
	} else {
		goto L34
	}
L34:
	;
	v97 = int32(1)
	v98 = v83 + v97
	v99 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v82)+1)))
	if v99 != 0 {
		v82 = v82 + v97
		v83 = v98
		v84 = v94
		v85 = v99
		goto L30
	} else {
		goto L35
	}
L35:
	;
	goto L31
L36:
	;
	v124 = int32(_a1416)
	goto L38
L37:
	;
	if v158-v163 != 0 {
		goto L7
	} else {
		goto L50
	}
L38:
	;
	v129 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v25))))
	if v129 != 0 {
		goto L41
	} else {
		goto L42
	}
L40:
	;
	v163 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v159))))
	goto L37
L41:
	;
	v131 = v25
	v132 = v124
	v133 = int32(10)
	v134 = v129
	goto L44
L42:
	;
	v158 = int32(0)
	v159 = v124
	goto L40
L43:
	;
	v158 = v155 & int32(255)
	v159 = v153
	goto L40
L44:
	;
	v138 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v132))))
	if v134&int32(255) != v138 {
		v153 = v132
		v155 = v134
		goto L43
	} else {
		goto L46
	}
L45:
	;
	v153 = v147
	v155 = int32(0)
	goto L43
L46:
	;
	if v138 == int32(0) {
		v153 = v132
		v155 = v134
		goto L43
	} else {
		goto L47
	}
L47:
	;
	v143 = v133 + int32(-1)
	if v143 == int32(0) {
		v153 = v132
		v155 = v134
		goto L43
	} else {
		goto L48
	}
L48:
	;
	v146 = int32(1)
	v147 = v132 + v146
	v148 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v131)+1)))
	if v148 != 0 {
		v131 = v131 + v146
		v132 = v147
		v133 = v143
		v134 = v148
		goto L44
	} else {
		goto L49
	}
L49:
	;
	goto L45
L50:
	;
	goto L8
L51:
	;
	v180 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
	v181 = int32(_a1413)
	goto L53
L52:
	;
	if v215-v220 != 0 {
		goto L4
	} else {
		goto L65
	}
L53:
	;
	v186 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v180))))
	if v186 != 0 {
		goto L56
	} else {
		goto L57
	}
L55:
	;
	v220 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v216))))
	goto L52
L56:
	;
	v188 = v180
	v189 = v181
	v190 = int32(4)
	v191 = v186
	goto L59
L57:
	;
	v215 = int32(0)
	v216 = v181
	goto L55
L58:
	;
	v215 = v212 & int32(255)
	v216 = v210
	goto L55
L59:
	;
	v195 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v189))))
	if v191&int32(255) != v195 {
		v210 = v189
		v212 = v191
		goto L58
	} else {
		goto L61
	}
L60:
	;
	v210 = v204
	v212 = int32(0)
	goto L58
L61:
	;
	if v195 == int32(0) {
		v210 = v189
		v212 = v191
		goto L58
	} else {
		goto L62
	}
L62:
	;
	v200 = v190 + int32(-1)
	if v200 == int32(0) {
		v210 = v189
		v212 = v191
		goto L58
	} else {
		goto L63
	}
L63:
	;
	v203 = int32(1)
	v204 = v189 + v203
	v205 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v188)+1)))
	if v205 != 0 {
		v188 = v188 + v203
		v189 = v204
		v190 = v200
		v191 = v205
		goto L59
	} else {
		goto L64
	}
L64:
	;
	goto L60
L65:
	;
	v301 = v175 & int32(-8193)
	goto L6
L66:
	;
	if v264-v269 != 0 {
		goto L4
	} else {
		goto L79
	}
L67:
	;
	v235 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v25))))
	if v235 != 0 {
		goto L70
	} else {
		goto L71
	}
L69:
	;
	v269 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v265))))
	goto L66
L70:
	;
	v237 = v25
	v238 = v230
	v239 = int32(4)
	v240 = v235
	goto L73
L71:
	;
	v264 = int32(0)
	v265 = v230
	goto L69
L72:
	;
	v264 = v261 & int32(255)
	v265 = v259
	goto L69
L73:
	;
	v244 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v238))))
	if v240&int32(255) != v244 {
		v259 = v238
		v261 = v240
		goto L72
	} else {
		goto L75
	}
L74:
	;
	v259 = v253
	v261 = int32(0)
	goto L72
L75:
	;
	if v244 == int32(0) {
		v259 = v238
		v261 = v240
		goto L72
	} else {
		goto L76
	}
L76:
	;
	v249 = v239 + int32(-1)
	if v249 == int32(0) {
		v259 = v238
		v261 = v240
		goto L72
	} else {
		goto L77
	}
L77:
	;
	v252 = int32(1)
	v253 = v238 + v252
	v254 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v237)+1)))
	if v254 != 0 {
		v237 = v237 + v252
		v238 = v253
		v239 = v249
		v240 = v254
		goto L73
	} else {
		goto L78
	}
L78:
	;
	goto L74
L79:
	;
	v277 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	if v277&int32(4104) != int32(8) {
		goto L4
	} else {
		goto L80
	}
L80:
	;
	v282 = *(*int32)(unsafe.Add(mBase, uint32(l2)+28))
	v283 = *(*int32)(unsafe.Add(mBase, uint32(v282)+12))
	v285 = F_sentinelInstanceMapCommand(m, l2, int32(_a1417))
	mBase = m.M
	v286 = m.ExcPending
	if v286 != 0 {
		goto L81
	} else {
		goto L82
	}
L81:
	;
	return
L82:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9))) = v285
	v290 = F_valkeyAsyncCommand(m, v283, int32(1004), l2, int32(_a1418), v9)
	mBase = m.M
	v291 = m.ExcPending
	if v291 != 0 {
		goto L81
	} else {
		goto L84
	}
L83:
	;
	v298 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	v301 = v298 | int32(4096)
	goto L6
L84:
	;
	if v290 != 0 {
		goto L83
	} else {
		goto L85
	}
L85:
	;
	v292 = *(*int32)(unsafe.Add(mBase, uint32(l2)+28))
	v293 = *(*int32)(unsafe.Add(mBase, uint32(v292)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v292)+8)) = v293 + int32(1)
	goto L83
}
func F_sentinelPublishCommand(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(v6)+4))
	v8 = F_objectGetVal(m, v7)
	mBase = m.M
	v9 = int32(_a1542)
	v12 = int32(*(*uint8)(unsafe.Add(mBase, _consts[779])))
	v13 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8))))
	if v13 == int32(0) {
		v36 = v12
		v37 = v13
		goto L3
	} else {
		goto L4
	}
L1:
	;
	v46 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v46)+8))
	v48 = F_objectGetVal(m, v47)
	mBase = m.M
	v50 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v50)+8))
	v52 = F_objectGetVal(m, v51)
	mBase = m.M
	v55 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v52+int32(-1)))))
	switch v55 & int32(7) {
	case 0:
		goto L18
	case 1:
		goto L17
	case 2:
		goto L16
	case 3:
		goto L15
	case 4:
		goto L14
	default:
		v72 = int32(0)
		goto L13
	}
L2:
	;
	if v37-v36&int32(255) == int32(0) {
		goto L1
	} else {
		goto L10
	}
L3:
	;
	goto L2
L4:
	;
	if v13 != v12&int32(255) {
		v36 = v12
		v37 = v13
		goto L3
	} else {
		goto L5
	}
L5:
	;
	v19 = v8
	v20 = v9
	goto L6
L6:
	;
	v23 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20)+1)))
	v24 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19)+1)))
	if v24 == int32(0) {
		v36 = v23
		v37 = v24
		goto L3
	} else {
		goto L8
	}
L7:
	;
	v36 = v23
	v37 = v24
	goto L3
L8:
	;
	v27 = int32(1)
	if v24 == v23&int32(255) {
		v19 = v19 + v27
		v20 = v20 + v27
		goto L6
	} else {
		goto L9
	}
L9:
	;
	goto L7
L10:
	;
	F_addReplyError(m, l0, int32(_a1543))
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	return
L12:
	;
	return
L13:
	;
	F_sentinelProcessHelloMessage(m, v48, v72)
	mBase = m.M
	v74 = m.ExcPending
	if v74 != 0 {
		goto L11
	} else {
		goto L19
	}
L14:
	;
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v52+int32(-17))))
	v72 = v71
	goto L13
L15:
	;
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v52+int32(-9))))
	v72 = v68
	goto L13
L16:
	;
	v65 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v52+int32(-5)))))
	v72 = v65
	goto L13
L17:
	;
	v62 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v52+int32(-3)))))
	v72 = v62
	goto L13
L18:
	;
	v72 = int32(base.Ui32(v55) >> (uint(int32(3)) % 32))
	goto L13
L19:
	;
	F_addReplyLongLong(m, l0, int64(1))
	mBase = m.M
	v77 = m.ExcPending
	if v77 != 0 {
		goto L11
	} else {
		goto L20
	}
L20:
	;
	return
}
func F_sentinelRefreshInstanceInfo(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v40 int32
	_ = v40
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v51 int32
	_ = v51
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v66 int32
	_ = v66
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v80 int32
	_ = v80
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v91 int32
	_ = v91
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v119 int32
	_ = v119
	var v122 int32
	_ = v122
	var v125 int32
	_ = v125
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v152 int32
	_ = v152
	var v154 int32
	_ = v154
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v177 int32
	_ = v177
	var v182 int32
	_ = v182
	var v197 int32
	_ = v197
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v207 int32
	_ = v207
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v216 int32
	_ = v216
	var v221 int32
	_ = v221
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v231 int32
	_ = v231
	var v233 int32
	_ = v233
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v241 int32
	_ = v241
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v262 int64
	_ = v262
	var v268 int64
	_ = v268
	var v270 int32
	_ = v270
	var v272 int32
	_ = v272
	var v275 int32
	_ = v275
	var v276 int32
	_ = v276
	var v280 int32
	_ = v280
	var v285 int32
	_ = v285
	var v292 int32
	_ = v292
	var v295 int32
	_ = v295
	var v298 int32
	_ = v298
	var v301 int32
	_ = v301
	var v302 int32
	_ = v302
	var v305 int32
	_ = v305
	var v306 int32
	_ = v306
	var v314 int32
	_ = v314
	var v315 int32
	_ = v315
	var v316 int32
	_ = v316
	var v319 int32
	_ = v319
	var v320 int32
	_ = v320
	var v322 int32
	_ = v322
	var v323 int32
	_ = v323
	var v325 int32
	_ = v325
	var v327 int32
	_ = v327
	var v337 int32
	_ = v337
	var v338 int32
	_ = v338
	var v339 int32
	_ = v339
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
	var v350 int32
	_ = v350
	var v355 int32
	_ = v355
	var v370 int32
	_ = v370
	var v371 int32
	_ = v371
	var v376 int32
	_ = v376
	var v379 int32
	_ = v379
	var v380 int32
	_ = v380
	var v381 int32
	_ = v381
	var v384 int32
	_ = v384
	var v385 int32
	_ = v385
	var v388 int32
	_ = v388
	var v389 int32
	_ = v389
	var v390 int32
	_ = v390
	var v393 int32
	_ = v393
	var v394 int32
	_ = v394
	var v395 int32
	_ = v395
	var v398 int32
	_ = v398
	var v399 int32
	_ = v399
	var v400 int32
	_ = v400
	var v401 int32
	_ = v401
	var v404 int32
	_ = v404
	var v405 int32
	_ = v405
	var v406 int32
	_ = v406
	var v408 int32
	_ = v408
	var v412 int32
	_ = v412
	var v416 int32
	_ = v416
	var v417 int32
	_ = v417
	var v418 int32
	_ = v418
	var v420 int32
	_ = v420
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
	var v434 int32
	_ = v434
	var v438 int32
	_ = v438
	var v439 int32
	_ = v439
	var v442 int32
	_ = v442
	var v443 int32
	_ = v443
	var v444 int32
	_ = v444
	var v447 int32
	_ = v447
	var v448 int32
	_ = v448
	var v451 int32
	_ = v451
	var v452 int32
	_ = v452
	var v453 int32
	_ = v453
	var v456 int32
	_ = v456
	var v457 int32
	_ = v457
	var v458 int32
	_ = v458
	var v461 int32
	_ = v461
	var v462 int32
	_ = v462
	var v463 int32
	_ = v463
	var v464 int32
	_ = v464
	var v467 int32
	_ = v467
	var v471 int32
	_ = v471
	var v473 int32
	_ = v473
	var v474 int32
	_ = v474
	var v475 int32
	_ = v475
	var v477 int32
	_ = v477
	var v481 int32
	_ = v481
	var v484 int32
	_ = v484
	var v486 int32
	_ = v486
	var v487 int32
	_ = v487
	var v489 int32
	_ = v489
	var v493 int32
	_ = v493
	var v496 int32
	_ = v496
	var v497 int32
	_ = v497
	var v498 int32
	_ = v498
	var v499 int32
	_ = v499
	var v501 int32
	_ = v501
	var v502 int32
	_ = v502
	var v507 int32
	_ = v507
	var v512 int32
	_ = v512
	var v513 int32
	_ = v513
	var v514 int32
	_ = v514
	var v515 int32
	_ = v515
	var v521 int32
	_ = v521
	var v522 int32
	_ = v522
	var v523 int32
	_ = v523
	var v524 int32
	_ = v524
	var v525 int32
	_ = v525
	var v528 int32
	_ = v528
	var v532 int32
	_ = v532
	var v533 int32
	_ = v533
	var v534 int32
	_ = v534
	var v536 int32
	_ = v536
	var v538 int32
	_ = v538
	var v539 int32
	_ = v539
	var v543 int32
	_ = v543
	var v546 int32
	_ = v546
	var v552 int32
	_ = v552
	var v553 int32
	_ = v553
	var v554 int32
	_ = v554
	var v560 int32
	_ = v560
	var v565 int32
	_ = v565
	var v566 int32
	_ = v566
	var v567 int32
	_ = v567
	var v568 int32
	_ = v568
	var v574 int32
	_ = v574
	var v575 int32
	_ = v575
	var v576 int32
	_ = v576
	var v577 int32
	_ = v577
	var v578 int32
	_ = v578
	var v581 int32
	_ = v581
	var v585 int32
	_ = v585
	var v586 int32
	_ = v586
	var v587 int32
	_ = v587
	var v589 int32
	_ = v589
	var v591 int32
	_ = v591
	var v592 int32
	_ = v592
	var v596 int32
	_ = v596
	var v599 int32
	_ = v599
	var v605 int32
	_ = v605
	var v606 int32
	_ = v606
	var v607 int32
	_ = v607
	var v608 int32
	_ = v608
	var v616 int32
	_ = v616
	var v617 int32
	_ = v617
	var v618 int32
	_ = v618
	var v623 int32
	_ = v623
	var v625 int32
	_ = v625
	var v626 int32
	_ = v626
	var v627 int32
	_ = v627
	var v630 int32
	_ = v630
	var v636 int32
	_ = v636
	var v637 int32
	_ = v637
	var v644 int32
	_ = v644
	var v651 int32
	_ = v651
	var v655 int32
	_ = v655
	var v662 int32
	_ = v662
	var v665 int32
	_ = v665
	var v668 int32
	_ = v668
	var v671 int32
	_ = v671
	var v672 int32
	_ = v672
	var v675 int32
	_ = v675
	var v676 int32
	_ = v676
	var v684 int32
	_ = v684
	var v685 int32
	_ = v685
	var v686 int32
	_ = v686
	var v689 int32
	_ = v689
	var v690 int32
	_ = v690
	var v692 int32
	_ = v692
	var v693 int32
	_ = v693
	var v695 int32
	_ = v695
	var v697 int32
	_ = v697
	var v707 int32
	_ = v707
	var v708 int32
	_ = v708
	var v709 int32
	_ = v709
	var v712 int32
	_ = v712
	var v713 int32
	_ = v713
	var v714 int32
	_ = v714
	var v717 int32
	_ = v717
	var v718 int32
	_ = v718
	var v720 int32
	_ = v720
	var v725 int32
	_ = v725
	var v740 int32
	_ = v740
	var v746 int64
	_ = v746
	var v750 int32
	_ = v750
	var v751 int32
	_ = v751
	var v754 int32
	_ = v754
	var v761 int32
	_ = v761
	var v764 int32
	_ = v764
	var v767 int32
	_ = v767
	var v770 int32
	_ = v770
	var v771 int32
	_ = v771
	var v774 int32
	_ = v774
	var v775 int32
	_ = v775
	var v783 int32
	_ = v783
	var v784 int32
	_ = v784
	var v785 int32
	_ = v785
	var v788 int32
	_ = v788
	var v789 int32
	_ = v789
	var v791 int32
	_ = v791
	var v792 int32
	_ = v792
	var v794 int32
	_ = v794
	var v796 int32
	_ = v796
	var v806 int32
	_ = v806
	var v807 int32
	_ = v807
	var v808 int32
	_ = v808
	var v811 int32
	_ = v811
	var v812 int32
	_ = v812
	var v813 int32
	_ = v813
	var v816 int32
	_ = v816
	var v817 int32
	_ = v817
	var v819 int32
	_ = v819
	var v824 int32
	_ = v824
	var v839 int32
	_ = v839
	var v847 int32
	_ = v847
	var v850 int32
	_ = v850
	var v853 int32
	_ = v853
	var v856 int32
	_ = v856
	var v857 int32
	_ = v857
	var v860 int32
	_ = v860
	var v861 int32
	_ = v861
	var v869 int32
	_ = v869
	var v870 int32
	_ = v870
	var v871 int32
	_ = v871
	var v874 int32
	_ = v874
	var v875 int32
	_ = v875
	var v877 int32
	_ = v877
	var v878 int32
	_ = v878
	var v880 int32
	_ = v880
	var v882 int32
	_ = v882
	var v892 int32
	_ = v892
	var v893 int32
	_ = v893
	var v894 int32
	_ = v894
	var v897 int32
	_ = v897
	var v898 int32
	_ = v898
	var v899 int32
	_ = v899
	var v902 int32
	_ = v902
	var v903 int32
	_ = v903
	var v905 int32
	_ = v905
	var v910 int32
	_ = v910
	var v925 int32
	_ = v925
	var v940 int32
	_ = v940
	var v943 int32
	_ = v943
	var v946 int32
	_ = v946
	var v949 int32
	_ = v949
	var v950 int32
	_ = v950
	var v953 int32
	_ = v953
	var v954 int32
	_ = v954
	var v962 int32
	_ = v962
	var v963 int32
	_ = v963
	var v964 int32
	_ = v964
	var v967 int32
	_ = v967
	var v968 int32
	_ = v968
	var v970 int32
	_ = v970
	var v971 int32
	_ = v971
	var v973 int32
	_ = v973
	var v975 int32
	_ = v975
	var v985 int32
	_ = v985
	var v986 int32
	_ = v986
	var v987 int32
	_ = v987
	var v990 int32
	_ = v990
	var v991 int32
	_ = v991
	var v992 int32
	_ = v992
	var v995 int32
	_ = v995
	var v996 int32
	_ = v996
	var v998 int32
	_ = v998
	var v1003 int32
	_ = v1003
	var v1018 int32
	_ = v1018
	var v1019 int32
	_ = v1019
	var v1023 int32
	_ = v1023
	var v1026 int32
	_ = v1026
	var v1028 int32
	_ = v1028
	var v1029 int32
	_ = v1029
	var v1030 int32
	_ = v1030
	var v1032 int32
	_ = v1032
	var v1036 int32
	_ = v1036
	var v1038 int32
	_ = v1038
	var v1039 int32
	_ = v1039
	var v1040 int32
	_ = v1040
	var v1042 int32
	_ = v1042
	var v1043 int32
	_ = v1043
	var v1044 int32
	_ = v1044
	var v1045 int32
	_ = v1045
	var v1050 int32
	_ = v1050
	var v1051 int32
	_ = v1051
	var v1054 int32
	_ = v1054
	var v1055 int32
	_ = v1055
	var v1058 int32
	_ = v1058
	var v1059 int32
	_ = v1059
	var v1060 int32
	_ = v1060
	var v1065 int32
	_ = v1065
	var v1068 int32
	_ = v1068
	var v1069 int32
	_ = v1069
	var v1071 int64
	_ = v1071
	var v1073 int32
	_ = v1073
	var v1074 int32
	_ = v1074
	var v1084 int32
	_ = v1084
	var v1087 int32
	_ = v1087
	var v1090 int32
	_ = v1090
	var v1093 int32
	_ = v1093
	var v1094 int32
	_ = v1094
	var v1097 int32
	_ = v1097
	var v1098 int32
	_ = v1098
	var v1106 int32
	_ = v1106
	var v1107 int32
	_ = v1107
	var v1108 int32
	_ = v1108
	var v1111 int32
	_ = v1111
	var v1112 int32
	_ = v1112
	var v1114 int32
	_ = v1114
	var v1115 int32
	_ = v1115
	var v1117 int32
	_ = v1117
	var v1119 int32
	_ = v1119
	var v1129 int32
	_ = v1129
	var v1130 int32
	_ = v1130
	var v1131 int32
	_ = v1131
	var v1134 int32
	_ = v1134
	var v1135 int32
	_ = v1135
	var v1136 int32
	_ = v1136
	var v1139 int32
	_ = v1139
	var v1140 int32
	_ = v1140
	var v1142 int32
	_ = v1142
	var v1147 int32
	_ = v1147
	var v1162 int32
	_ = v1162
	var v1163 int32
	_ = v1163
	var v1169 int32
	_ = v1169
	var v1174 int32
	_ = v1174
	var v1175 int32
	_ = v1175
	var v1176 int32
	_ = v1176
	var v1177 int32
	_ = v1177
	var v1183 int32
	_ = v1183
	var v1184 int32
	_ = v1184
	var v1185 int32
	_ = v1185
	var v1186 int32
	_ = v1186
	var v1187 int32
	_ = v1187
	var v1190 int32
	_ = v1190
	var v1194 int32
	_ = v1194
	var v1195 int32
	_ = v1195
	var v1196 int32
	_ = v1196
	var v1198 int32
	_ = v1198
	var v1200 int32
	_ = v1200
	var v1201 int32
	_ = v1201
	var v1205 int32
	_ = v1205
	var v1208 int32
	_ = v1208
	var v1214 int32
	_ = v1214
	var v1217 int64
	_ = v1217
	var v1219 int32
	_ = v1219
	var v1220 int32
	_ = v1220
	var v1230 int32
	_ = v1230
	var v1233 int32
	_ = v1233
	var v1236 int32
	_ = v1236
	var v1239 int32
	_ = v1239
	var v1240 int32
	_ = v1240
	var v1243 int32
	_ = v1243
	var v1244 int32
	_ = v1244
	var v1252 int32
	_ = v1252
	var v1253 int32
	_ = v1253
	var v1254 int32
	_ = v1254
	var v1257 int32
	_ = v1257
	var v1258 int32
	_ = v1258
	var v1260 int32
	_ = v1260
	var v1261 int32
	_ = v1261
	var v1263 int32
	_ = v1263
	var v1265 int32
	_ = v1265
	var v1275 int32
	_ = v1275
	var v1276 int32
	_ = v1276
	var v1277 int32
	_ = v1277
	var v1280 int32
	_ = v1280
	var v1281 int32
	_ = v1281
	var v1282 int32
	_ = v1282
	var v1285 int32
	_ = v1285
	var v1286 int32
	_ = v1286
	var v1288 int32
	_ = v1288
	var v1293 int32
	_ = v1293
	var v1308 int32
	_ = v1308
	var v1310 int32
	_ = v1310
	var v1311 int32
	_ = v1311
	var v1314 int32
	_ = v1314
	var v1316 int32
	_ = v1316
	var v1317 int32
	_ = v1317
	var v1318 int32
	_ = v1318
	var v1320 int32
	_ = v1320
	var v1324 int32
	_ = v1324
	var v1326 int32
	_ = v1326
	var v1327 int32
	_ = v1327
	var v1328 int32
	_ = v1328
	var v1330 int32
	_ = v1330
	var v1331 int32
	_ = v1331
	var v1332 int32
	_ = v1332
	var v1333 int32
	_ = v1333
	var v1338 int32
	_ = v1338
	var v1339 int32
	_ = v1339
	var v1342 int32
	_ = v1342
	var v1343 int32
	_ = v1343
	var v1346 int32
	_ = v1346
	var v1347 int32
	_ = v1347
	var v1348 int32
	_ = v1348
	var v1353 int32
	_ = v1353
	var v1354 int32
	_ = v1354
	var v1364 int32
	_ = v1364
	var v1367 int32
	_ = v1367
	var v1370 int32
	_ = v1370
	var v1373 int32
	_ = v1373
	var v1374 int32
	_ = v1374
	var v1377 int32
	_ = v1377
	var v1378 int32
	_ = v1378
	var v1386 int32
	_ = v1386
	var v1387 int32
	_ = v1387
	var v1388 int32
	_ = v1388
	var v1391 int32
	_ = v1391
	var v1392 int32
	_ = v1392
	var v1394 int32
	_ = v1394
	var v1395 int32
	_ = v1395
	var v1397 int32
	_ = v1397
	var v1399 int32
	_ = v1399
	var v1409 int32
	_ = v1409
	var v1410 int32
	_ = v1410
	var v1411 int32
	_ = v1411
	var v1414 int32
	_ = v1414
	var v1415 int32
	_ = v1415
	var v1416 int32
	_ = v1416
	var v1419 int32
	_ = v1419
	var v1420 int32
	_ = v1420
	var v1422 int32
	_ = v1422
	var v1427 int32
	_ = v1427
	var v1442 int32
	_ = v1442
	var v1448 int32
	_ = v1448
	var v1453 int32
	_ = v1453
	var v1454 int32
	_ = v1454
	var v1455 int32
	_ = v1455
	var v1456 int32
	_ = v1456
	var v1462 int32
	_ = v1462
	var v1463 int32
	_ = v1463
	var v1464 int32
	_ = v1464
	var v1465 int32
	_ = v1465
	var v1466 int32
	_ = v1466
	var v1469 int32
	_ = v1469
	var v1473 int32
	_ = v1473
	var v1474 int32
	_ = v1474
	var v1475 int32
	_ = v1475
	var v1477 int32
	_ = v1477
	var v1479 int32
	_ = v1479
	var v1480 int32
	_ = v1480
	var v1484 int32
	_ = v1484
	var v1487 int32
	_ = v1487
	var v1493 int32
	_ = v1493
	var v1495 int32
	_ = v1495
	var v1496 int32
	_ = v1496
	var v1506 int32
	_ = v1506
	var v1509 int32
	_ = v1509
	var v1512 int32
	_ = v1512
	var v1515 int32
	_ = v1515
	var v1516 int32
	_ = v1516
	var v1519 int32
	_ = v1519
	var v1520 int32
	_ = v1520
	var v1528 int32
	_ = v1528
	var v1529 int32
	_ = v1529
	var v1530 int32
	_ = v1530
	var v1533 int32
	_ = v1533
	var v1534 int32
	_ = v1534
	var v1536 int32
	_ = v1536
	var v1537 int32
	_ = v1537
	var v1539 int32
	_ = v1539
	var v1541 int32
	_ = v1541
	var v1551 int32
	_ = v1551
	var v1552 int32
	_ = v1552
	var v1553 int32
	_ = v1553
	var v1556 int32
	_ = v1556
	var v1557 int32
	_ = v1557
	var v1558 int32
	_ = v1558
	var v1561 int32
	_ = v1561
	var v1562 int32
	_ = v1562
	var v1564 int32
	_ = v1564
	var v1569 int32
	_ = v1569
	var v1584 int32
	_ = v1584
	var v1590 int64
	_ = v1590
	var v1592 int32
	_ = v1592
	var v1593 int32
	_ = v1593
	var v1595 int32
	_ = v1595
	var v1604 int32
	_ = v1604
	var v1607 int32
	_ = v1607
	var v1610 int32
	_ = v1610
	var v1613 int32
	_ = v1613
	var v1614 int32
	_ = v1614
	var v1617 int32
	_ = v1617
	var v1618 int32
	_ = v1618
	var v1626 int32
	_ = v1626
	var v1627 int32
	_ = v1627
	var v1628 int32
	_ = v1628
	var v1631 int32
	_ = v1631
	var v1632 int32
	_ = v1632
	var v1634 int32
	_ = v1634
	var v1635 int32
	_ = v1635
	var v1637 int32
	_ = v1637
	var v1639 int32
	_ = v1639
	var v1649 int32
	_ = v1649
	var v1650 int32
	_ = v1650
	var v1651 int32
	_ = v1651
	var v1654 int32
	_ = v1654
	var v1655 int32
	_ = v1655
	var v1656 int32
	_ = v1656
	var v1659 int32
	_ = v1659
	var v1660 int32
	_ = v1660
	var v1662 int32
	_ = v1662
	var v1667 int32
	_ = v1667
	var v1682 int32
	_ = v1682
	var v1688 int32
	_ = v1688
	var v1693 int32
	_ = v1693
	var v1694 int32
	_ = v1694
	var v1695 int32
	_ = v1695
	var v1696 int32
	_ = v1696
	var v1702 int32
	_ = v1702
	var v1703 int32
	_ = v1703
	var v1704 int32
	_ = v1704
	var v1705 int32
	_ = v1705
	var v1706 int32
	_ = v1706
	var v1709 int32
	_ = v1709
	var v1713 int32
	_ = v1713
	var v1714 int32
	_ = v1714
	var v1715 int32
	_ = v1715
	var v1717 int32
	_ = v1717
	var v1719 int32
	_ = v1719
	var v1720 int32
	_ = v1720
	var v1724 int32
	_ = v1724
	var v1727 int32
	_ = v1727
	var v1733 int32
	_ = v1733
	var v1735 int32
	_ = v1735
	var v1736 int32
	_ = v1736
	var v1737 int32
	_ = v1737
	var v1741 int32
	_ = v1741
	var v1748 int32
	_ = v1748
	var v1751 int32
	_ = v1751
	var v1754 int32
	_ = v1754
	var v1757 int32
	_ = v1757
	var v1758 int32
	_ = v1758
	var v1761 int32
	_ = v1761
	var v1762 int32
	_ = v1762
	var v1770 int32
	_ = v1770
	var v1771 int32
	_ = v1771
	var v1772 int32
	_ = v1772
	var v1775 int32
	_ = v1775
	var v1776 int32
	_ = v1776
	var v1778 int32
	_ = v1778
	var v1779 int32
	_ = v1779
	var v1781 int32
	_ = v1781
	var v1783 int32
	_ = v1783
	var v1793 int32
	_ = v1793
	var v1794 int32
	_ = v1794
	var v1795 int32
	_ = v1795
	var v1798 int32
	_ = v1798
	var v1799 int32
	_ = v1799
	var v1800 int32
	_ = v1800
	var v1803 int32
	_ = v1803
	var v1804 int32
	_ = v1804
	var v1806 int32
	_ = v1806
	var v1811 int32
	_ = v1811
	var v1826 int32
	_ = v1826
	var v1827 int32
	_ = v1827
	var v1830 int32
	_ = v1830
	var v1833 int32
	_ = v1833
	var v1836 int32
	_ = v1836
	var v1839 int32
	_ = v1839
	var v1840 int32
	_ = v1840
	var v1845 int32
	_ = v1845
	var v1846 int32
	_ = v1846
	var v1854 int32
	_ = v1854
	var v1855 int32
	_ = v1855
	var v1856 int32
	_ = v1856
	var v1859 int32
	_ = v1859
	var v1860 int32
	_ = v1860
	var v1862 int32
	_ = v1862
	var v1863 int32
	_ = v1863
	var v1865 int32
	_ = v1865
	var v1867 int32
	_ = v1867
	var v1877 int32
	_ = v1877
	var v1878 int32
	_ = v1878
	var v1879 int32
	_ = v1879
	var v1882 int32
	_ = v1882
	var v1883 int32
	_ = v1883
	var v1884 int32
	_ = v1884
	var v1887 int32
	_ = v1887
	var v1888 int32
	_ = v1888
	var v1890 int32
	_ = v1890
	var v1895 int32
	_ = v1895
	var v1910 int32
	_ = v1910
	var v1911 int32
	_ = v1911
	var v1912 int32
	_ = v1912
	var v1914 int32
	_ = v1914
	var v1917 int64
	_ = v1917
	var v1920 int32
	_ = v1920
	var v1925 int32
	_ = v1925
	var v1926 int32
	_ = v1926
	var v1933 int32
	_ = v1933
	var v1941 int64
	_ = v1941
	var v1943 int32
	_ = v1943
	var v1945 int32
	_ = v1945
	var v1946 int32
	_ = v1946
	var v1948 int64
	_ = v1948
	var v1953 int64
	_ = v1953
	var v1955 int32
	_ = v1955
	var v1960 int32
	_ = v1960
	var v1968 int32
	_ = v1968
	var v1973 int32
	_ = v1973
	var v1977 int32
	_ = v1977
	var v1978 int32
	_ = v1978
	var v1980 int32
	_ = v1980
	var v1989 int32
	_ = v1989
	var v1990 int32
	_ = v1990
	var v1995 int32
	_ = v1995
	var v2000 int64
	_ = v2000
	var v2002 int64
	_ = v2002
	var v2003 int32
	_ = v2003
	var v2005 int32
	_ = v2005
	var v2006 int32
	_ = v2006
	var v2011 int32
	_ = v2011
	var v2013 int32
	_ = v2013
	var v2014 int32
	_ = v2014
	var v2015 int32
	_ = v2015
	var v2018 int32
	_ = v2018
	var v2024 int32
	_ = v2024
	var v2025 int32
	_ = v2025
	var v2032 int32
	_ = v2032
	var v2039 int32
	_ = v2039
	var v2045 int32
	_ = v2045
	var v2047 int32
	_ = v2047
	var v2050 int32
	_ = v2050
	var v2051 int32
	_ = v2051
	var v2056 int32
	_ = v2056
	var v2057 int32
	_ = v2057
	var v2058 int32
	_ = v2058
	var v2059 int32
	_ = v2059
	var v2060 int32
	_ = v2060
	var v2061 int32
	_ = v2061
	var v2067 int32
	_ = v2067
	var v2068 int32
	_ = v2068
	var v2069 int32
	_ = v2069
	var v2072 int32
	_ = v2072
	var v2074 int32
	_ = v2074
	var v2076 int32
	_ = v2076
	var v2077 int64
	_ = v2077
	var v2086 int32
	_ = v2086
	var v2095 int32
	_ = v2095
	var v2096 int32
	_ = v2096
	var v2097 int64
	_ = v2097
	var v2099 int32
	_ = v2099
	var v2120 int32
	_ = v2120
	var v2122 int64
	_ = v2122
	var v2131 int32
	_ = v2131
	var v2140 int32
	_ = v2140
	var v2141 int32
	_ = v2141
	var v2142 int64
	_ = v2142
	var v2144 int32
	_ = v2144
	var v2164 int32
	_ = v2164
	var v2165 int32
	_ = v2165
	var v2169 int32
	_ = v2169
	var v2171 int32
	_ = v2171
	var v2172 int32
	_ = v2172
	var v2173 int32
	_ = v2173
	var v2177 int32
	_ = v2177
	var v2200 int32
	_ = v2200
	var v2201 int32
	_ = v2201
	var v2202 int32
	_ = v2202
	var v2207 int32
	_ = v2207
	var v2208 int32
	_ = v2208
	var v2209 int32
	_ = v2209
	var v2210 int32
	_ = v2210
	var v2217 int32
	_ = v2217
	var v2220 int32
	_ = v2220
	var v2221 int64
	_ = v2221
	var v2222 int64
	_ = v2222
	var v2223 int64
	_ = v2223
	var v2226 int64
	_ = v2226
	var v2231 int64
	_ = v2231
	var v2232 int64
	_ = v2232
	var v2233 int64
	_ = v2233
	var v2235 int64
	_ = v2235
	var v2238 int64
	_ = v2238
	var v2241 int64
	_ = v2241
	var v2242 int64
	_ = v2242
	var v2245 int32
	_ = v2245
	var v2246 int32
	_ = v2246
	var v2247 int32
	_ = v2247
	var v2248 int32
	_ = v2248
	var v2254 int32
	_ = v2254
	var v2259 int32
	_ = v2259
	var v2262 int32
	_ = v2262
	var v2263 int32
	_ = v2263
	var v2270 int32
	_ = v2270
	var v2273 int64
	_ = v2273
	var v2274 int64
	_ = v2274
	var v2275 int64
	_ = v2275
	var v2278 int64
	_ = v2278
	var v2282 int64
	_ = v2282
	var v2283 int64
	_ = v2283
	var v2285 int64
	_ = v2285
	var v2288 int64
	_ = v2288
	var v2291 int64
	_ = v2291
	var v2292 int64
	_ = v2292
	var v2295 int32
	_ = v2295
	var v2296 int32
	_ = v2296
	var v2297 int32
	_ = v2297
	var v2298 int32
	_ = v2298
	var v2304 int32
	_ = v2304
	var v2310 int32
	_ = v2310
	var v2317 int32
	_ = v2317
	var v2318 int32
	_ = v2318
	var v2319 int32
	_ = v2319
	var v2320 int32
	_ = v2320
	var v2322 int32
	_ = v2322
	var v2323 int32
	_ = v2323
	var v2328 int32
	_ = v2328
	var v2331 int32
	_ = v2331
	var v2332 int32
	_ = v2332
	var v2333 int32
	_ = v2333
	var v2334 int32
	_ = v2334
	var v2341 int32
	_ = v2341
	var v2342 int32
	_ = v2342
	var v2344 int32
	_ = v2344
	var v2347 int32
	_ = v2347
	var v2350 int32
	_ = v2350
	var v2352 int32
	_ = v2352
	var v2353 int32
	_ = v2353
	var v2354 int32
	_ = v2354
	var v2356 int32
	_ = v2356
	var v2360 int32
	_ = v2360
	var v2362 int32
	_ = v2362
	var v2363 int32
	_ = v2363
	var v2364 int32
	_ = v2364
	var v2366 int32
	_ = v2366
	var v2367 int32
	_ = v2367
	var v2368 int32
	_ = v2368
	var v2369 int32
	_ = v2369
	var v2374 int32
	_ = v2374
	var v2375 int32
	_ = v2375
	var v2378 int32
	_ = v2378
	var v2379 int32
	_ = v2379
	var v2382 int32
	_ = v2382
	var v2383 int32
	_ = v2383
	var v2384 int32
	_ = v2384
	var v2388 int32
	_ = v2388
	var v2389 int32
	_ = v2389
	var v2391 int32
	_ = v2391
	var v2398 int32
	_ = v2398
	var v2401 int64
	_ = v2401
	var v2402 int64
	_ = v2402
	var v2403 int64
	_ = v2403
	var v2406 int64
	_ = v2406
	var v2410 int64
	_ = v2410
	var v2411 int64
	_ = v2411
	var v2413 int64
	_ = v2413
	var v2416 int64
	_ = v2416
	var v2419 int64
	_ = v2419
	var v2420 int64
	_ = v2420
	var v2423 int32
	_ = v2423
	var v2424 int32
	_ = v2424
	var v2425 int32
	_ = v2425
	var v2426 int32
	_ = v2426
	var v2432 int32
	_ = v2432
	var v2441 int32
	_ = v2441
	var v2456 int32
	_ = v2456
	var v2459 int32
	_ = v2459
	var v2460 int32
	_ = v2460
	var v2461 int32
	_ = v2461
	var v2462 int32
	_ = v2462
	var v2467 int32
	_ = v2467
	var v2470 int32
	_ = v2470
	var v2471 int32
	_ = v2471
	var v2472 int32
	_ = v2472
	var v2473 int32
	_ = v2473
	var v2480 int32
	_ = v2480
	var v2481 int32
	_ = v2481
	var v2483 int32
	_ = v2483
	var v2486 int32
	_ = v2486
	var v2489 int32
	_ = v2489
	var v2491 int32
	_ = v2491
	var v2492 int32
	_ = v2492
	var v2493 int32
	_ = v2493
	var v2495 int32
	_ = v2495
	var v2499 int32
	_ = v2499
	var v2501 int32
	_ = v2501
	var v2502 int32
	_ = v2502
	var v2503 int32
	_ = v2503
	var v2505 int32
	_ = v2505
	var v2506 int32
	_ = v2506
	var v2507 int32
	_ = v2507
	var v2508 int32
	_ = v2508
	var v2513 int32
	_ = v2513
	var v2514 int32
	_ = v2514
	var v2517 int32
	_ = v2517
	var v2518 int32
	_ = v2518
	var v2521 int32
	_ = v2521
	var v2522 int32
	_ = v2522
	var v2523 int32
	_ = v2523
	var v2525 int32
	_ = v2525
	var v2526 int32
	_ = v2526
	var v2527 int32
	_ = v2527
	var v2528 int32
	_ = v2528
	var v2529 int32
	_ = v2529
	var v2531 int32
	_ = v2531
	var v2542 int32
	_ = v2542
	var v2546 int32
	_ = v2546
	var v2551 int32
	_ = v2551
	var v2562 int32
	_ = v2562
	var v2575 int32
	_ = v2575
	v14 = m.G0
	v16 = v14 - int32(160)
	m.G0 = v16
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+292))
	F_sdsfree(m, v18)
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v21 = F_sdsnew(m, l1)
	mBase = m.M
	v22 = m.ExcPending
	if v22 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l0)+168)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+292)) = v21
	if l1&int32(3) == int32(0) {
		v47 = l1
		goto L6
	} else {
		goto L7
	}
L4:
	;
	v85 = F_sdssplitlen(m, l1, v80, int32(_a132), int32(2), v16+int32(76))
	mBase = m.M
	v86 = m.ExcPending
	if v86 != 0 {
		goto L1
	} else {
		goto L20
	}
L5:
	;
	v80 = v72 - l1
	goto L4
L6:
	;
	v51 = v47
	goto L14
L7:
	;
	v33 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	if v33 != 0 {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v36 = l1
	goto L10
L9:
	;
	v80 = l1 - l1
	goto L4
L10:
	;
	v40 = v36 + int32(1)
	if v40&int32(3) == int32(0) {
		v47 = v40
		goto L6
	} else {
		goto L12
	}
L12:
	;
	v45 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v40))))
	if v45 != 0 {
		v36 = v40
		goto L10
	} else {
		goto L13
	}
L13:
	;
	v72 = v40
	goto L5
L14:
	;
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v51)))
	v60 = int32(-2139062144)
	if (int32(16843008)-v57|v57)&v60 == v60 {
		v51 = v51 + int32(4)
		goto L14
	} else {
		goto L16
	}
L15:
	;
	v66 = v51
	goto L17
L16:
	;
	goto L15
L17:
	;
	v70 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v66))))
	if v70 != 0 {
		v66 = v66 + int32(1)
		goto L17
	} else {
		goto L19
	}
L18:
	;
	v72 = v66
	goto L5
L19:
	;
	goto L18
L20:
	;
	v87 = *(*int32)(unsafe.Add(mBase, uint32(v16)+76))
	if int32(1) <= v87 {
		goto L22
	} else {
		goto L23
	}
L21:
	;
	v1941 = F_mstime(m)
	mBase = m.M
	*(*int64)(unsafe.Add(mBase, uint32(l0)+96)) = v1941
	v1943 = *(*int32)(unsafe.Add(mBase, uint32(v16)+76))
	F_sdsfreesplitres(m, v85, v1943)
	mBase = m.M
	v1945 = m.ExcPending
	if v1945 != 0 {
		goto L1
	} else {
		goto L563
	}
L22:
	;
	v91 = int32(0)
	v98 = v91
	v99 = v91
	goto L24
L23:
	;
	v1933 = int32(0)
	goto L21
L24:
	;
	v109 = *(*int32)(unsafe.Add(mBase, uint32(v85+v99<<(uint(int32(2))%32))))
	v110 = int32(-1)
	v111 = v109 + v110
	v112 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v111))))
	switch v112&int32(7) + v110 {
	case 0:
		goto L31
	case 1:
		goto L30
	case 2:
		goto L29
	case 3:
		goto L28
	default:
		goto L26
	}
L25:
	;
	v1933 = v1920
	goto L21
L26:
	;
	v280 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v280&int32(1) == int32(0) {
		goto L73
	} else {
		goto L74
	}
L27:
	;
	if base.Ui32(v129) < base.Ui32(int32(47)) {
		goto L26
	} else {
		goto L32
	}
L28:
	;
	v128 = *(*int32)(unsafe.Add(mBase, uint32(v109+int32(-17))))
	v129 = v128
	goto L27
L29:
	;
	v125 = *(*int32)(unsafe.Add(mBase, uint32(v109+int32(-9))))
	v129 = v125
	goto L27
L30:
	;
	v122 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v109+int32(-5)))))
	v129 = v122
	goto L27
L31:
	;
	v119 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v109+int32(-3)))))
	v129 = v119
	goto L27
L32:
	;
	v132 = int32(_a1419)
	v133 = int32(7)
	goto L37
L33:
	;
	if v197 != 0 {
		goto L26
	} else {
		goto L49
	}
L34:
	;
	v197 = int32(0)
	goto L33
L35:
	;
	v169 = v164
	v170 = v165
	v171 = v166
	goto L45
L36:
	;
	if v154 == int32(0) {
		goto L34
	} else {
		goto L43
	}
L37:
	;
	if (v132|v109)&int32(3) != 0 {
		v164 = v109
		v165 = v132
		v166 = v133
		goto L35
	} else {
		goto L38
	}
L38:
	;
	v141 = v109
	v142 = v132
	v143 = v133
	goto L39
L39:
	;
	v146 = *(*int32)(unsafe.Add(mBase, uint32(v141)))
	v147 = *(*int32)(unsafe.Add(mBase, uint32(v142)))
	if v146 != v147 {
		v164 = v141
		v165 = v142
		v166 = v143
		goto L35
	} else {
		goto L41
	}
L40:
	;
	goto L36
L41:
	;
	v149 = int32(4)
	v150 = v142 + v149
	v152 = v141 + v149
	v154 = v143 + int32(-4)
	if base.Ui32(int32(3)) < base.Ui32(v154) {
		v141 = v152
		v142 = v150
		v143 = v154
		goto L39
	} else {
		goto L42
	}
L42:
	;
	goto L40
L43:
	;
	v164 = v152
	v165 = v150
	v166 = v154
	goto L35
L44:
	;
	v197 = v174 - v175
	goto L33
L45:
	;
	v174 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v169))))
	v175 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v170))))
	if v174 != v175 {
		goto L44
	} else {
		goto L47
	}
L47:
	;
	v177 = int32(1)
	v182 = v171 + int32(-1)
	if v182 == int32(0) {
		goto L34
	} else {
		goto L48
	}
L48:
	;
	v169 = v169 + v177
	v170 = v170 + v177
	v171 = v182
	goto L45
L49:
	;
	v199 = v109 + int32(7)
	v200 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v200 == int32(0) {
		goto L50
	} else {
		goto L51
	}
L50:
	;
	v275 = F_sdsnewlen(m, v199, int32(40))
	mBase = m.M
	v276 = m.ExcPending
	if v276 != 0 {
		goto L1
	} else {
		goto L71
	}
L51:
	;
	goto L53
L52:
	;
	if v236-v241 == int32(0) {
		goto L26
	} else {
		goto L65
	}
L53:
	;
	v207 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v200))))
	if v207 != 0 {
		goto L56
	} else {
		goto L57
	}
L55:
	;
	v241 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v237))))
	goto L52
L56:
	;
	v209 = v200
	v210 = v199
	v211 = int32(40)
	v212 = v207
	goto L59
L57:
	;
	v236 = int32(0)
	v237 = v199
	goto L55
L58:
	;
	v236 = v233 & int32(255)
	v237 = v231
	goto L55
L59:
	;
	v216 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v210))))
	if v212&int32(255) != v216 {
		v231 = v210
		v233 = v212
		goto L58
	} else {
		goto L61
	}
L60:
	;
	v231 = v225
	v233 = int32(0)
	goto L58
L61:
	;
	if v216 == int32(0) {
		v231 = v210
		v233 = v212
		goto L58
	} else {
		goto L62
	}
L62:
	;
	v221 = v211 + int32(-1)
	if v221 == int32(0) {
		v231 = v210
		v233 = v212
		goto L58
	} else {
		goto L63
	}
L63:
	;
	v224 = int32(1)
	v225 = v210 + v224
	v226 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v209)+1)))
	if v226 != 0 {
		v209 = v209 + v224
		v210 = v225
		v211 = v221
		v212 = v226
		goto L59
	} else {
		goto L64
	}
L64:
	;
	goto L60
L65:
	;
	F_sentinelEvent(m, int32(2), int32(_a1420), l0, int32(_a1362), int32(0))
	mBase = m.M
	v256 = m.ExcPending
	if v256 != 0 {
		goto L1
	} else {
		goto L66
	}
L66:
	;
	v257 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v257&int32(1) == int32(0) {
		goto L67
	} else {
		goto L68
	}
L67:
	;
	v270 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	F_sdsfree(m, v270)
	mBase = m.M
	v272 = m.ExcPending
	if v272 != 0 {
		goto L1
	} else {
		goto L70
	}
L68:
	;
	v262 = *(*int64)(unsafe.Add(mBase, uint32(l0)+80))
	if v262 == int64(0) {
		goto L67
	} else {
		goto L69
	}
L69:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v257 | int32(8192)
	v268 = F_mstime(m)
	mBase = m.M
	*(*int64)(unsafe.Add(mBase, uint32(l0)+88)) = v268
	goto L67
L70:
	;
	goto L50
L71:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v275
	goto L26
L72:
	;
	v1925 = v99 + int32(1)
	v1926 = *(*int32)(unsafe.Add(mBase, uint32(v16)+76))
	if v1925 < v1926 {
		v98 = v1920
		v99 = v1925
		goto L24
	} else {
		goto L562
	}
L73:
	;
	v655 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v111))))
	switch v655&int32(7) + int32(-1) {
	case 0:
		goto L212
	case 1:
		goto L211
	case 2:
		goto L210
	case 3:
		goto L209
	default:
		v751 = v655
		goto L207
	}
L74:
	;
	v285 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v111))))
	switch v285 & int32(7) {
	case 0:
		goto L80
	case 1:
		goto L79
	case 2:
		goto L78
	case 3:
		goto L77
	case 4:
		goto L76
	default:
		goto L73
	}
L75:
	;
	if base.Ui32(v302) < base.Ui32(int32(7)) {
		goto L73
	} else {
		goto L81
	}
L76:
	;
	v301 = *(*int32)(unsafe.Add(mBase, uint32(v109+int32(-17))))
	v302 = v301
	goto L75
L77:
	;
	v298 = *(*int32)(unsafe.Add(mBase, uint32(v109+int32(-9))))
	v302 = v298
	goto L75
L78:
	;
	v295 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v109+int32(-5)))))
	v302 = v295
	goto L75
L79:
	;
	v292 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v109+int32(-3)))))
	v302 = v292
	goto L75
L80:
	;
	v302 = int32(base.Ui32(v285) >> (uint(int32(3)) % 32))
	goto L75
L81:
	;
	v305 = int32(_a505)
	v306 = int32(5)
	goto L86
L82:
	;
	if v370 != 0 {
		goto L73
	} else {
		goto L98
	}
L83:
	;
	v370 = int32(0)
	goto L82
L84:
	;
	v342 = v337
	v343 = v338
	v344 = v339
	goto L94
L85:
	;
	if v327 == int32(0) {
		goto L83
	} else {
		goto L92
	}
L86:
	;
	if (v305|v109)&int32(3) != 0 {
		v337 = v109
		v338 = v305
		v339 = v306
		goto L84
	} else {
		goto L87
	}
L87:
	;
	v314 = v109
	v315 = v305
	v316 = v306
	goto L88
L88:
	;
	v319 = *(*int32)(unsafe.Add(mBase, uint32(v314)))
	v320 = *(*int32)(unsafe.Add(mBase, uint32(v315)))
	if v319 != v320 {
		v337 = v314
		v338 = v315
		v339 = v316
		goto L84
	} else {
		goto L90
	}
L89:
	;
	goto L85
L90:
	;
	v322 = int32(4)
	v323 = v315 + v322
	v325 = v314 + v322
	v327 = v316 + int32(-4)
	if base.Ui32(int32(3)) < base.Ui32(v327) {
		v314 = v325
		v315 = v323
		v316 = v327
		goto L88
	} else {
		goto L91
	}
L91:
	;
	goto L89
L92:
	;
	v337 = v325
	v338 = v323
	v339 = v327
	goto L84
L93:
	;
	v370 = v347 - v348
	goto L82
L94:
	;
	v347 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v342))))
	v348 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v343))))
	if v347 != v348 {
		goto L93
	} else {
		goto L96
	}
L96:
	;
	v350 = int32(1)
	v355 = v344 + int32(-1)
	if v355 == int32(0) {
		goto L83
	} else {
		goto L97
	}
L97:
	;
	v342 = v342 + v350
	v343 = v343 + v350
	v344 = v355
	goto L94
L98:
	;
	v371 = int32(*(*int8)(unsafe.Add(mBase, uint32(v109)+5)))
	if base.Ui32(int32(9)) < base.Ui32(v371+int32(-48)) {
		goto L73
	} else {
		goto L99
	}
L99:
	;
	v376 = int32(_a1421)
	v379 = int32(*(*int8)(unsafe.Add(mBase, _consts[754])))
	if v379 != 0 {
		goto L104
	} else {
		goto L105
	}
L100:
	;
	v507 = v501
	goto L164
L101:
	;
	v499 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v498))) = uint8(v499)
	v501 = v496
	v502 = v497
	goto L100
L102:
	;
	v439 = int32(_a1422)
	v442 = int32(*(*int8)(unsafe.Add(mBase, _consts[755])))
	if v442 != 0 {
		goto L136
	} else {
		goto L137
	}
L103:
	;
	if v404 != 0 {
		goto L102
	} else {
		goto L119
	}
L104:
	;
	v380 = int32(0)
	v381 = F_strchr(m, v109, v379)
	mBase = m.M
	if v381 == v380 {
		v401 = v380
		goto L106
	} else {
		goto L107
	}
L105:
	;
	v404 = v109
	goto L103
L106:
	;
	v404 = v401
	goto L103
L107:
	;
	v384 = int32(*(*uint8)(unsafe.Add(mBase, _consts[756])))
	if v384 != 0 {
		goto L108
	} else {
		goto L109
	}
L108:
	;
	v385 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v381)+1)))
	if v385 == int32(0) {
		v401 = v380
		goto L106
	} else {
		goto L110
	}
L109:
	;
	v404 = v381
	goto L103
L110:
	;
	v388 = int32(*(*uint8)(unsafe.Add(mBase, _consts[757])))
	if v388 != 0 {
		goto L111
	} else {
		goto L112
	}
L111:
	;
	v390 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v381)+2)))
	if v390 == int32(0) {
		v401 = v380
		goto L106
	} else {
		goto L113
	}
L112:
	;
	v389 = F_twobyte_strstr(m, v381, v376)
	mBase = m.M
	v404 = v389
	goto L103
L113:
	;
	v393 = int32(*(*uint8)(unsafe.Add(mBase, _consts[758])))
	if v393 != 0 {
		goto L114
	} else {
		goto L115
	}
L114:
	;
	v395 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v381)+3)))
	if v395 == int32(0) {
		v401 = v380
		goto L106
	} else {
		goto L116
	}
L115:
	;
	v394 = F_threebyte_strstr(m, v381, v376)
	mBase = m.M
	v404 = v394
	goto L103
L116:
	;
	v398 = int32(*(*uint8)(unsafe.Add(mBase, _consts[759])))
	if v398 != 0 {
		goto L117
	} else {
		goto L118
	}
L117:
	;
	v400 = F_twoway_strstr(m, v381, v376)
	mBase = m.M
	v401 = v400
	goto L106
L118:
	;
	v399 = F_fourbyte_strstr(m, v381, v376)
	mBase = m.M
	v404 = v399
	goto L103
L119:
	;
	v405 = int32(58)
	v406 = F___strchrnul(m, v109, v405)
	mBase = m.M
	v408 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v406))))
	if v408 == v405 {
		goto L121
	} else {
		goto L122
	}
L120:
	;
	if v412 == int32(0) {
		v1920 = v98
		goto L72
	} else {
		goto L124
	}
L121:
	;
	v412 = v406
	goto L123
L122:
	;
	v412 = int32(0)
	goto L123
L123:
	;
	goto L120
L124:
	;
	v416 = v412 + int32(1)
	v417 = int32(44)
	v418 = F___strchrnul(m, v416, v417)
	mBase = m.M
	v420 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v418))))
	if v420 == v417 {
		goto L126
	} else {
		goto L127
	}
L125:
	;
	if v424 == int32(0) {
		v1920 = v98
		goto L72
	} else {
		goto L129
	}
L126:
	;
	v424 = v418
	goto L128
L127:
	;
	v424 = int32(0)
	goto L128
L128:
	;
	goto L125
L129:
	;
	v427 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v424))) = uint8(v427)
	v430 = v424 + int32(1)
	v431 = int32(44)
	v432 = F___strchrnul(m, v430, v431)
	mBase = m.M
	v434 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v432))))
	if v434 == v431 {
		goto L131
	} else {
		goto L132
	}
L130:
	;
	if v438 != 0 {
		v496 = v430
		v497 = v416
		v498 = v438
		goto L101
	} else {
		goto L134
	}
L131:
	;
	v438 = v432
	goto L133
L132:
	;
	v438 = v427
	goto L133
L133:
	;
	goto L130
L134:
	;
	v1920 = v98
	goto L72
L135:
	;
	if v467 == int32(0) {
		v1920 = v98
		goto L72
	} else {
		goto L151
	}
L136:
	;
	v443 = int32(0)
	v444 = F_strchr(m, v109, v442)
	mBase = m.M
	if v444 == v443 {
		v464 = v443
		goto L138
	} else {
		goto L139
	}
L137:
	;
	v467 = v109
	goto L135
L138:
	;
	v467 = v464
	goto L135
L139:
	;
	v447 = int32(*(*uint8)(unsafe.Add(mBase, _consts[760])))
	if v447 != 0 {
		goto L140
	} else {
		goto L141
	}
L140:
	;
	v448 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v444)+1)))
	if v448 == int32(0) {
		v464 = v443
		goto L138
	} else {
		goto L142
	}
L141:
	;
	v467 = v444
	goto L135
L142:
	;
	v451 = int32(*(*uint8)(unsafe.Add(mBase, _consts[761])))
	if v451 != 0 {
		goto L143
	} else {
		goto L144
	}
L143:
	;
	v453 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v444)+2)))
	if v453 == int32(0) {
		v464 = v443
		goto L138
	} else {
		goto L145
	}
L144:
	;
	v452 = F_twobyte_strstr(m, v444, v439)
	mBase = m.M
	v467 = v452
	goto L135
L145:
	;
	v456 = int32(*(*uint8)(unsafe.Add(mBase, _consts[762])))
	if v456 != 0 {
		goto L146
	} else {
		goto L147
	}
L146:
	;
	v458 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v444)+3)))
	if v458 == int32(0) {
		v464 = v443
		goto L138
	} else {
		goto L148
	}
L147:
	;
	v457 = F_threebyte_strstr(m, v444, v439)
	mBase = m.M
	v467 = v457
	goto L135
L148:
	;
	v461 = int32(*(*uint8)(unsafe.Add(mBase, _consts[763])))
	if v461 != 0 {
		goto L149
	} else {
		goto L150
	}
L149:
	;
	v463 = F_twoway_strstr(m, v444, v439)
	mBase = m.M
	v464 = v463
	goto L138
L150:
	;
	v462 = F_fourbyte_strstr(m, v444, v439)
	mBase = m.M
	v467 = v462
	goto L135
L151:
	;
	v471 = v467 + int32(5)
	v473 = v404 + int32(3)
	v474 = int32(44)
	v475 = F___strchrnul(m, v473, v474)
	mBase = m.M
	v477 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v475))))
	if v477 == v474 {
		goto L154
	} else {
		goto L155
	}
L152:
	;
	v486 = int32(44)
	v487 = F___strchrnul(m, v471, v486)
	mBase = m.M
	v489 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v487))))
	if v489 == v486 {
		goto L159
	} else {
		goto L160
	}
L153:
	;
	if v481 == int32(0) {
		goto L152
	} else {
		goto L157
	}
L154:
	;
	v481 = v475
	goto L156
L155:
	;
	v481 = int32(0)
	goto L156
L156:
	;
	goto L153
L157:
	;
	v484 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v481))) = uint8(v484)
	goto L152
L158:
	;
	if v493 == int32(0) {
		v501 = v471
		v502 = v473
		goto L100
	} else {
		goto L162
	}
L159:
	;
	v493 = v487
	goto L161
L160:
	;
	v493 = int32(0)
	goto L161
L161:
	;
	goto L158
L162:
	;
	v496 = v471
	v497 = v473
	v498 = v493
	goto L101
L163:
	;
	v553 = F_sentinelValkeyInstanceLookupReplica(m, l0, v502, v552)
	mBase = m.M
	v554 = m.ExcPending
	if v554 != 0 {
		goto L1
	} else {
		goto L178
	}
L164:
	;
	v512 = v507 + int32(1)
	v513 = int32(*(*int8)(unsafe.Add(mBase, uint32(v507))))
	v514 = F___isspace_1(m, v513)
	mBase = m.M
	if v514 != 0 {
		v507 = v512
		goto L164
	} else {
		goto L166
	}
L165:
	;
	v515 = int32(1)
	switch v513&int32(255) + int32(-43) {
	case 0:
		v521 = v515
		goto L168
	default:
		v523 = v507
		v524 = v513
		v525 = v515
		goto L167
	case 2:
		goto L169
	}
L166:
	;
	goto L165
L167:
	;
	v528 = v524 + int32(-48)
	if base.Ui32(int32(9)) < base.Ui32(v528) {
		v546 = int32(0)
		goto L170
	} else {
		goto L171
	}
L168:
	;
	v522 = int32(*(*int8)(unsafe.Add(mBase, uint32(v512))))
	v523 = v512
	v524 = v522
	v525 = v521
	goto L167
L169:
	;
	v521 = int32(0)
	goto L168
L170:
	;
	if v525 != 0 {
		goto L175
	} else {
		goto L176
	}
L171:
	;
	v532 = int32(0)
	v533 = v523
	v534 = v528
	goto L172
L172:
	;
	v536 = int32(10)
	v538 = v532*v536 - v534
	v539 = int32(*(*int8)(unsafe.Add(mBase, uint32(v533)+1)))
	v543 = v539 + int32(-48)
	if base.Ui32(v543) < base.Ui32(v536) {
		v532 = v538
		v533 = v533 + int32(1)
		v534 = v543
		goto L172
	} else {
		goto L174
	}
L173:
	;
	v546 = v538
	goto L170
L174:
	;
	goto L173
L175:
	;
	v552 = int32(0) - v546
	goto L177
L176:
	;
	v552 = v546
	goto L177
L177:
	;
	goto L163
L178:
	;
	if v553 != 0 {
		goto L73
	} else {
		goto L179
	}
L179:
	;
	v560 = v501
	goto L181
L180:
	;
	v606 = *(*int32)(unsafe.Add(mBase, uint32(l0)+152))
	v607 = F_createSentinelValkeyInstance(m, int32(0), int32(2), v502, v605, v606, l0)
	mBase = m.M
	v608 = m.ExcPending
	if v608 != 0 {
		goto L1
	} else {
		goto L195
	}
L181:
	;
	v565 = v560 + int32(1)
	v566 = int32(*(*int8)(unsafe.Add(mBase, uint32(v560))))
	v567 = F___isspace_1(m, v566)
	mBase = m.M
	if v567 != 0 {
		v560 = v565
		goto L181
	} else {
		goto L183
	}
L182:
	;
	v568 = int32(1)
	switch v566&int32(255) + int32(-43) {
	case 0:
		v574 = v568
		goto L185
	default:
		v576 = v560
		v577 = v566
		v578 = v568
		goto L184
	case 2:
		goto L186
	}
L183:
	;
	goto L182
L184:
	;
	v581 = v577 + int32(-48)
	if base.Ui32(int32(9)) < base.Ui32(v581) {
		v599 = int32(0)
		goto L187
	} else {
		goto L188
	}
L185:
	;
	v575 = int32(*(*int8)(unsafe.Add(mBase, uint32(v565))))
	v576 = v565
	v577 = v575
	v578 = v574
	goto L184
L186:
	;
	v574 = int32(0)
	goto L185
L187:
	;
	if v578 != 0 {
		goto L192
	} else {
		goto L193
	}
L188:
	;
	v585 = int32(0)
	v586 = v576
	v587 = v581
	goto L189
L189:
	;
	v589 = int32(10)
	v591 = v585*v589 - v587
	v592 = int32(*(*int8)(unsafe.Add(mBase, uint32(v586)+1)))
	v596 = v592 + int32(-48)
	if base.Ui32(v596) < base.Ui32(v589) {
		v585 = v591
		v586 = v586 + int32(1)
		v587 = v596
		goto L189
	} else {
		goto L191
	}
L190:
	;
	v599 = v591
	goto L187
L191:
	;
	goto L190
L192:
	;
	v605 = int32(0) - v599
	goto L194
L193:
	;
	v605 = v599
	goto L194
L194:
	;
	goto L180
L195:
	;
	if v607 == int32(0) {
		goto L73
	} else {
		goto L196
	}
L196:
	;
	F_sentinelEvent(m, int32(2), int32(_a1361), v607, int32(_a1362), int32(0))
	mBase = m.M
	v616 = m.ExcPending
	if v616 != 0 {
		goto L1
	} else {
		goto L197
	}
L197:
	;
	v617 = int32(_a44)
	v618 = *(*int32)(unsafe.Add(mBase, _consts[219]))
	*(*int32)(unsafe.Add(mBase, _consts[219])) = int32(10)
	v623 = *(*int32)(unsafe.Add(mBase, _consts[392]))
	v625 = F_rewriteConfig(m, v623, int32(0))
	mBase = m.M
	v626 = m.ExcPending
	if v626 != 0 {
		goto L1
	} else {
		goto L198
	}
L198:
	;
	v627 = int32(_a44)
	*(*int32)(unsafe.Add(mBase, _consts[219])) = v618
	v630 = *(*int32)(unsafe.Add(mBase, _consts[6]))
	if v625 != int32(-1) {
		goto L199
	} else {
		goto L200
	}
L199:
	;
	if int32(2) < v630 {
		goto L73
	} else {
		goto L205
	}
L200:
	;
	if int32(3) < v630 {
		goto L73
	} else {
		goto L201
	}
L201:
	;
	goto L202
L202:
	;
	v636 = *(*int32)(unsafe.Add(mBase, _consts[5]))
	v637 = F___strerror_l(m, v636, v636)
	mBase = m.M
	goto L203
L203:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+64)) = v637
	F__serverLog(m, int32(3), int32(_a1337), v16+int32(64))
	mBase = m.M
	v644 = m.ExcPending
	if v644 != 0 {
		goto L1
	} else {
		goto L204
	}
L204:
	;
	goto L73
L205:
	;
	F__serverLog(m, int32(2), int32(_a1338), int32(0))
	mBase = m.M
	v651 = m.ExcPending
	if v651 != 0 {
		goto L1
	} else {
		goto L206
	}
L206:
	;
	goto L73
L207:
	;
	v754 = v751 & int32(7)
	switch v754 {
	case 0:
		goto L240
	case 1:
		goto L239
	case 2:
		goto L238
	case 3:
		goto L237
	case 4:
		goto L236
	default:
		goto L234
	}
L208:
	;
	if base.Ui32(v672) < base.Ui32(int32(32)) {
		v751 = v655
		goto L207
	} else {
		goto L213
	}
L209:
	;
	v671 = *(*int32)(unsafe.Add(mBase, uint32(v109+int32(-17))))
	v672 = v671
	goto L208
L210:
	;
	v668 = *(*int32)(unsafe.Add(mBase, uint32(v109+int32(-9))))
	v672 = v668
	goto L208
L211:
	;
	v665 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v109+int32(-5)))))
	v672 = v665
	goto L208
L212:
	;
	v662 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v109+int32(-3)))))
	v672 = v662
	goto L208
L213:
	;
	v675 = int32(_a1423)
	v676 = int32(30)
	goto L218
L214:
	;
	if v740 != 0 {
		v751 = v655
		goto L207
	} else {
		goto L230
	}
L215:
	;
	v740 = int32(0)
	goto L214
L216:
	;
	v712 = v707
	v713 = v708
	v714 = v709
	goto L226
L217:
	;
	if v697 == int32(0) {
		goto L215
	} else {
		goto L224
	}
L218:
	;
	if (v675|v109)&int32(3) != 0 {
		v707 = v109
		v708 = v675
		v709 = v676
		goto L216
	} else {
		goto L219
	}
L219:
	;
	v684 = v109
	v685 = v675
	v686 = v676
	goto L220
L220:
	;
	v689 = *(*int32)(unsafe.Add(mBase, uint32(v684)))
	v690 = *(*int32)(unsafe.Add(mBase, uint32(v685)))
	if v689 != v690 {
		v707 = v684
		v708 = v685
		v709 = v686
		goto L216
	} else {
		goto L222
	}
L221:
	;
	goto L217
L222:
	;
	v692 = int32(4)
	v693 = v685 + v692
	v695 = v684 + v692
	v697 = v686 + int32(-4)
	if base.Ui32(int32(3)) < base.Ui32(v697) {
		v684 = v695
		v685 = v693
		v686 = v697
		goto L220
	} else {
		goto L223
	}
L223:
	;
	goto L221
L224:
	;
	v707 = v695
	v708 = v693
	v709 = v697
	goto L216
L225:
	;
	v740 = v717 - v718
	goto L214
L226:
	;
	v717 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v712))))
	v718 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v713))))
	if v717 != v718 {
		goto L225
	} else {
		goto L228
	}
L228:
	;
	v720 = int32(1)
	v725 = v714 + int32(-1)
	if v725 == int32(0) {
		goto L215
	} else {
		goto L229
	}
L229:
	;
	v712 = v712 + v720
	v713 = v713 + v720
	v714 = v725
	goto L226
L230:
	;
	v746 = F_strtox_2(m, v109+int32(31), int32(0), int32(10), int64(-9223372036854775807-1))
	mBase = m.M
	goto L231
L231:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l0)+168)) = v746 * int64(1000)
	v750 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v111))))
	v751 = v750
	goto L207
L232:
	;
	v1741 = v1736 & int32(7)
	switch v1741 {
	case 0:
		goto L516
	case 1:
		goto L515
	case 2:
		goto L514
	case 3:
		goto L513
	case 4:
		goto L512
	default:
		v1920 = v1737
		goto L72
	}
L233:
	;
	switch v754 {
	case 0:
		goto L292
	case 1:
		goto L291
	case 2:
		goto L290
	case 3:
		goto L289
	case 4:
		goto L288
	default:
		v1074 = v751
		goto L286
	}
L234:
	;
	if v98 != int32(2) {
		v1736 = v751
		v1737 = v98
		goto L232
	} else {
		goto L285
	}
L235:
	;
	if base.Ui32(v771) < base.Ui32(int32(11)) {
		goto L241
	} else {
		goto L242
	}
L236:
	;
	v770 = *(*int32)(unsafe.Add(mBase, uint32(v109+int32(-17))))
	v771 = v770
	goto L235
L237:
	;
	v767 = *(*int32)(unsafe.Add(mBase, uint32(v109+int32(-9))))
	v771 = v767
	goto L235
L238:
	;
	v764 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v109+int32(-5)))))
	v771 = v764
	goto L235
L239:
	;
	v761 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v109+int32(-3)))))
	v771 = v761
	goto L235
L240:
	;
	v771 = int32(base.Ui32(v751&int32(248)) >> (uint(int32(3)) % 32))
	goto L235
L241:
	;
	switch v754 {
	default:
		goto L265
	case 1:
		goto L264
	case 2:
		goto L263
	case 3:
		goto L262
	case 4:
		goto L261
	}
L242:
	;
	v774 = int32(_a1424)
	v775 = int32(11)
	goto L247
L243:
	;
	if v839 != 0 {
		goto L241
	} else {
		goto L259
	}
L244:
	;
	v839 = int32(0)
	goto L243
L245:
	;
	v811 = v806
	v812 = v807
	v813 = v808
	goto L255
L246:
	;
	if v796 == int32(0) {
		goto L244
	} else {
		goto L253
	}
L247:
	;
	if (v774|v109)&int32(3) != 0 {
		v806 = v109
		v807 = v774
		v808 = v775
		goto L245
	} else {
		goto L248
	}
L248:
	;
	v783 = v109
	v784 = v774
	v785 = v775
	goto L249
L249:
	;
	v788 = *(*int32)(unsafe.Add(mBase, uint32(v783)))
	v789 = *(*int32)(unsafe.Add(mBase, uint32(v784)))
	if v788 != v789 {
		v806 = v783
		v807 = v784
		v808 = v785
		goto L245
	} else {
		goto L251
	}
L250:
	;
	goto L246
L251:
	;
	v791 = int32(4)
	v792 = v784 + v791
	v794 = v783 + v791
	v796 = v785 + int32(-4)
	if base.Ui32(int32(3)) < base.Ui32(v796) {
		v783 = v794
		v784 = v792
		v785 = v796
		goto L249
	} else {
		goto L252
	}
L252:
	;
	goto L250
L253:
	;
	v806 = v794
	v807 = v792
	v808 = v796
	goto L245
L254:
	;
	v839 = v816 - v817
	goto L243
L255:
	;
	v816 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v811))))
	v817 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v812))))
	if v816 != v817 {
		goto L254
	} else {
		goto L257
	}
L257:
	;
	v819 = int32(1)
	v824 = v813 + int32(-1)
	if v824 == int32(0) {
		goto L244
	} else {
		goto L258
	}
L258:
	;
	v811 = v811 + v819
	v812 = v812 + v819
	v813 = v824
	goto L255
L259:
	;
	v1736 = v751
	v1737 = int32(1)
	goto L232
L260:
	;
	if base.Ui32(v857) < base.Ui32(int32(10)) {
		goto L234
	} else {
		goto L266
	}
L261:
	;
	v856 = *(*int32)(unsafe.Add(mBase, uint32(v109+int32(-17))))
	v857 = v856
	goto L260
L262:
	;
	v853 = *(*int32)(unsafe.Add(mBase, uint32(v109+int32(-9))))
	v857 = v853
	goto L260
L263:
	;
	v850 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v109+int32(-5)))))
	v857 = v850
	goto L260
L264:
	;
	v847 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v109+int32(-3)))))
	v857 = v847
	goto L260
L265:
	;
	v857 = int32(base.Ui32(v751&int32(248)) >> (uint(int32(3)) % 32))
	goto L260
L266:
	;
	v860 = int32(_a1425)
	v861 = int32(10)
	goto L271
L267:
	;
	if v925 == int32(0) {
		goto L233
	} else {
		goto L283
	}
L268:
	;
	v925 = int32(0)
	goto L267
L269:
	;
	v897 = v892
	v898 = v893
	v899 = v894
	goto L279
L270:
	;
	if v882 == int32(0) {
		goto L268
	} else {
		goto L277
	}
L271:
	;
	if (v860|v109)&int32(3) != 0 {
		v892 = v109
		v893 = v860
		v894 = v861
		goto L269
	} else {
		goto L272
	}
L272:
	;
	v869 = v109
	v870 = v860
	v871 = v861
	goto L273
L273:
	;
	v874 = *(*int32)(unsafe.Add(mBase, uint32(v869)))
	v875 = *(*int32)(unsafe.Add(mBase, uint32(v870)))
	if v874 != v875 {
		v892 = v869
		v893 = v870
		v894 = v871
		goto L269
	} else {
		goto L275
	}
L274:
	;
	goto L270
L275:
	;
	v877 = int32(4)
	v878 = v870 + v877
	v880 = v869 + v877
	v882 = v871 + int32(-4)
	if base.Ui32(int32(3)) < base.Ui32(v882) {
		v869 = v880
		v870 = v878
		v871 = v882
		goto L273
	} else {
		goto L276
	}
L276:
	;
	goto L274
L277:
	;
	v892 = v880
	v893 = v878
	v894 = v882
	goto L269
L278:
	;
	v925 = v902 - v903
	goto L267
L279:
	;
	v902 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v897))))
	v903 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v898))))
	if v902 != v903 {
		goto L278
	} else {
		goto L281
	}
L281:
	;
	v905 = int32(1)
	v910 = v899 + int32(-1)
	if v910 == int32(0) {
		goto L268
	} else {
		goto L282
	}
L282:
	;
	v897 = v897 + v905
	v898 = v898 + v905
	v899 = v910
	goto L279
L283:
	;
	if v98 == int32(2) {
		goto L233
	} else {
		goto L284
	}
L284:
	;
	v1736 = v751
	v1737 = v98
	goto L232
L285:
	;
	goto L233
L286:
	;
	switch v1074 & int32(7) {
	case 0:
		goto L334
	case 1:
		goto L333
	case 2:
		goto L332
	case 3:
		goto L331
	case 4:
		goto L330
	default:
		v1220 = v1074
		goto L328
	}
L287:
	;
	if base.Ui32(v950) < base.Ui32(int32(12)) {
		v1074 = v751
		goto L286
	} else {
		goto L293
	}
L288:
	;
	v949 = *(*int32)(unsafe.Add(mBase, uint32(v109+int32(-17))))
	v950 = v949
	goto L287
L289:
	;
	v946 = *(*int32)(unsafe.Add(mBase, uint32(v109+int32(-9))))
	v950 = v946
	goto L287
L290:
	;
	v943 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v109+int32(-5)))))
	v950 = v943
	goto L287
L291:
	;
	v940 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v109+int32(-3)))))
	v950 = v940
	goto L287
L292:
	;
	v950 = int32(base.Ui32(v751&int32(248)) >> (uint(int32(3)) % 32))
	goto L287
L293:
	;
	v953 = int32(_a1426)
	v954 = int32(12)
	goto L298
L294:
	;
	if v1018 != 0 {
		v1074 = v751
		goto L286
	} else {
		goto L310
	}
L295:
	;
	v1018 = int32(0)
	goto L294
L296:
	;
	v990 = v985
	v991 = v986
	v992 = v987
	goto L306
L297:
	;
	if v975 == int32(0) {
		goto L295
	} else {
		goto L304
	}
L298:
	;
	if (v953|v109)&int32(3) != 0 {
		v985 = v109
		v986 = v953
		v987 = v954
		goto L296
	} else {
		goto L299
	}
L299:
	;
	v962 = v109
	v963 = v953
	v964 = v954
	goto L300
L300:
	;
	v967 = *(*int32)(unsafe.Add(mBase, uint32(v962)))
	v968 = *(*int32)(unsafe.Add(mBase, uint32(v963)))
	if v967 != v968 {
		v985 = v962
		v986 = v963
		v987 = v964
		goto L296
	} else {
		goto L302
	}
L301:
	;
	goto L297
L302:
	;
	v970 = int32(4)
	v971 = v963 + v970
	v973 = v962 + v970
	v975 = v964 + int32(-4)
	if base.Ui32(int32(3)) < base.Ui32(v975) {
		v962 = v973
		v963 = v971
		v964 = v975
		goto L300
	} else {
		goto L303
	}
L303:
	;
	goto L301
L304:
	;
	v985 = v973
	v986 = v971
	v987 = v975
	goto L296
L305:
	;
	v1018 = v995 - v996
	goto L294
L306:
	;
	v995 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v990))))
	v996 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v991))))
	if v995 != v996 {
		goto L305
	} else {
		goto L308
	}
L308:
	;
	v998 = int32(1)
	v1003 = v992 + int32(-1)
	if v1003 == int32(0) {
		goto L295
	} else {
		goto L309
	}
L309:
	;
	v990 = v990 + v998
	v991 = v991 + v998
	v992 = v1003
	goto L306
L310:
	;
	v1019 = *(*int32)(unsafe.Add(mBase, uint32(l0)+196))
	if v1019 == int32(0) {
		goto L311
	} else {
		goto L312
	}
L311:
	;
	F_sdsfree(m, v1019)
	mBase = m.M
	v1065 = m.ExcPending
	if v1065 != 0 {
		goto L1
	} else {
		goto L326
	}
L312:
	;
	v1023 = v109 + int32(12)
	v1026 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1023))))
	if v1026 != 0 {
		goto L315
	} else {
		goto L316
	}
L313:
	;
	if v1058-v1060 == int32(0) {
		v1074 = v751
		goto L286
	} else {
		goto L325
	}
L314:
	;
	v1058 = F_tolower(m, v1054)
	mBase = m.M
	v1059 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1055))))
	v1060 = F_tolower(m, v1059)
	mBase = m.M
	goto L313
L315:
	;
	v1028 = v1023
	v1029 = v1019
	v1030 = v1026
	goto L318
L316:
	;
	v1054 = int32(0)
	v1055 = v1019
	goto L314
L317:
	;
	v1054 = v1051 & int32(255)
	v1055 = v1050
	goto L314
L318:
	;
	v1032 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1029))))
	if v1032 == int32(0) {
		v1050 = v1029
		v1051 = v1030
		goto L317
	} else {
		goto L320
	}
L319:
	;
	v1050 = v1044
	v1051 = int32(0)
	goto L317
L320:
	;
	v1036 = v1030 & int32(255)
	if v1036 == v1032 {
		goto L321
	} else {
		goto L322
	}
L321:
	;
	v1043 = int32(1)
	v1044 = v1029 + v1043
	v1045 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1028)+1)))
	if v1045 != 0 {
		v1028 = v1028 + v1043
		v1029 = v1044
		v1030 = v1045
		goto L318
	} else {
		goto L324
	}
L322:
	;
	v1038 = F_tolower(m, v1036)
	mBase = m.M
	v1039 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1029))))
	v1040 = F_tolower(m, v1039)
	mBase = m.M
	if v1038 == v1040 {
		goto L321
	} else {
		goto L323
	}
L323:
	;
	v1042 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1028))))
	v1050 = v1029
	v1051 = v1042
	goto L317
L324:
	;
	goto L319
L325:
	;
	goto L311
L326:
	;
	v1068 = F_sdsnew(m, v109+int32(12))
	mBase = m.M
	v1069 = m.ExcPending
	if v1069 != 0 {
		goto L1
	} else {
		goto L327
	}
L327:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+196)) = v1068
	v1071 = F_mstime(m)
	mBase = m.M
	*(*int64)(unsafe.Add(mBase, uint32(l0)+120)) = v1071
	v1073 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v111))))
	v1074 = v1073
	goto L286
L328:
	;
	switch v1220 & int32(7) {
	case 0:
		goto L375
	case 1:
		goto L374
	case 2:
		goto L373
	case 3:
		goto L372
	case 4:
		goto L371
	default:
		v1354 = v1220
		goto L369
	}
L329:
	;
	if base.Ui32(v1094) < base.Ui32(int32(12)) {
		v1220 = v1074
		goto L328
	} else {
		goto L335
	}
L330:
	;
	v1093 = *(*int32)(unsafe.Add(mBase, uint32(v109+int32(-17))))
	v1094 = v1093
	goto L329
L331:
	;
	v1090 = *(*int32)(unsafe.Add(mBase, uint32(v109+int32(-9))))
	v1094 = v1090
	goto L329
L332:
	;
	v1087 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v109+int32(-5)))))
	v1094 = v1087
	goto L329
L333:
	;
	v1084 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v109+int32(-3)))))
	v1094 = v1084
	goto L329
L334:
	;
	v1094 = int32(base.Ui32(v1074&int32(248)) >> (uint(int32(3)) % 32))
	goto L329
L335:
	;
	v1097 = int32(_a1427)
	v1098 = int32(12)
	goto L340
L336:
	;
	if v1162 != 0 {
		v1220 = v1074
		goto L328
	} else {
		goto L352
	}
L337:
	;
	v1162 = int32(0)
	goto L336
L338:
	;
	v1134 = v1129
	v1135 = v1130
	v1136 = v1131
	goto L348
L339:
	;
	if v1119 == int32(0) {
		goto L337
	} else {
		goto L346
	}
L340:
	;
	if (v1097|v109)&int32(3) != 0 {
		v1129 = v109
		v1130 = v1097
		v1131 = v1098
		goto L338
	} else {
		goto L341
	}
L341:
	;
	v1106 = v109
	v1107 = v1097
	v1108 = v1098
	goto L342
L342:
	;
	v1111 = *(*int32)(unsafe.Add(mBase, uint32(v1106)))
	v1112 = *(*int32)(unsafe.Add(mBase, uint32(v1107)))
	if v1111 != v1112 {
		v1129 = v1106
		v1130 = v1107
		v1131 = v1108
		goto L338
	} else {
		goto L344
	}
L343:
	;
	goto L339
L344:
	;
	v1114 = int32(4)
	v1115 = v1107 + v1114
	v1117 = v1106 + v1114
	v1119 = v1108 + int32(-4)
	if base.Ui32(int32(3)) < base.Ui32(v1119) {
		v1106 = v1117
		v1107 = v1115
		v1108 = v1119
		goto L342
	} else {
		goto L345
	}
L345:
	;
	goto L343
L346:
	;
	v1129 = v1117
	v1130 = v1115
	v1131 = v1119
	goto L338
L347:
	;
	v1162 = v1139 - v1140
	goto L336
L348:
	;
	v1139 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1134))))
	v1140 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1135))))
	if v1139 != v1140 {
		goto L347
	} else {
		goto L350
	}
L350:
	;
	v1142 = int32(1)
	v1147 = v1136 + int32(-1)
	if v1147 == int32(0) {
		goto L337
	} else {
		goto L351
	}
L351:
	;
	v1134 = v1134 + v1142
	v1135 = v1135 + v1142
	v1136 = v1147
	goto L348
L352:
	;
	v1163 = *(*int32)(unsafe.Add(mBase, uint32(l0)+200))
	v1169 = v109 + int32(12)
	goto L354
L353:
	;
	if v1163 == v1214 {
		v1220 = v1074
		goto L328
	} else {
		goto L368
	}
L354:
	;
	v1174 = v1169 + int32(1)
	v1175 = int32(*(*int8)(unsafe.Add(mBase, uint32(v1169))))
	v1176 = F___isspace_1(m, v1175)
	mBase = m.M
	if v1176 != 0 {
		v1169 = v1174
		goto L354
	} else {
		goto L356
	}
L355:
	;
	v1177 = int32(1)
	switch v1175&int32(255) + int32(-43) {
	case 0:
		v1183 = v1177
		goto L358
	default:
		v1185 = v1169
		v1186 = v1175
		v1187 = v1177
		goto L357
	case 2:
		goto L359
	}
L356:
	;
	goto L355
L357:
	;
	v1190 = v1186 + int32(-48)
	if base.Ui32(int32(9)) < base.Ui32(v1190) {
		v1208 = int32(0)
		goto L360
	} else {
		goto L361
	}
L358:
	;
	v1184 = int32(*(*int8)(unsafe.Add(mBase, uint32(v1174))))
	v1185 = v1174
	v1186 = v1184
	v1187 = v1183
	goto L357
L359:
	;
	v1183 = int32(0)
	goto L358
L360:
	;
	if v1187 != 0 {
		goto L365
	} else {
		goto L366
	}
L361:
	;
	v1194 = int32(0)
	v1195 = v1185
	v1196 = v1190
	goto L362
L362:
	;
	v1198 = int32(10)
	v1200 = v1194*v1198 - v1196
	v1201 = int32(*(*int8)(unsafe.Add(mBase, uint32(v1195)+1)))
	v1205 = v1201 + int32(-48)
	if base.Ui32(v1205) < base.Ui32(v1198) {
		v1194 = v1200
		v1195 = v1195 + int32(1)
		v1196 = v1205
		goto L362
	} else {
		goto L364
	}
L363:
	;
	v1208 = v1200
	goto L360
L364:
	;
	goto L363
L365:
	;
	v1214 = int32(0) - v1208
	goto L367
L366:
	;
	v1214 = v1208
	goto L367
L367:
	;
	goto L353
L368:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+200)) = v1214
	v1217 = F_mstime(m)
	mBase = m.M
	*(*int64)(unsafe.Add(mBase, uint32(l0)+120)) = v1217
	v1219 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v111))))
	v1220 = v1219
	goto L328
L369:
	;
	switch v1354 & int32(7) {
	case 0:
		goto L412
	case 1:
		goto L411
	case 2:
		goto L410
	case 3:
		goto L409
	case 4:
		goto L408
	default:
		v1496 = v1354
		goto L406
	}
L370:
	;
	if base.Ui32(v1240) < base.Ui32(int32(19)) {
		v1354 = v1220
		goto L369
	} else {
		goto L376
	}
L371:
	;
	v1239 = *(*int32)(unsafe.Add(mBase, uint32(v109+int32(-17))))
	v1240 = v1239
	goto L370
L372:
	;
	v1236 = *(*int32)(unsafe.Add(mBase, uint32(v109+int32(-9))))
	v1240 = v1236
	goto L370
L373:
	;
	v1233 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v109+int32(-5)))))
	v1240 = v1233
	goto L370
L374:
	;
	v1230 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v109+int32(-3)))))
	v1240 = v1230
	goto L370
L375:
	;
	v1240 = int32(base.Ui32(v1220&int32(248)) >> (uint(int32(3)) % 32))
	goto L370
L376:
	;
	v1243 = int32(_a1428)
	v1244 = int32(19)
	goto L381
L377:
	;
	if v1308 != 0 {
		v1354 = v1220
		goto L369
	} else {
		goto L393
	}
L378:
	;
	v1308 = int32(0)
	goto L377
L379:
	;
	v1280 = v1275
	v1281 = v1276
	v1282 = v1277
	goto L389
L380:
	;
	if v1265 == int32(0) {
		goto L378
	} else {
		goto L387
	}
L381:
	;
	if (v1243|v109)&int32(3) != 0 {
		v1275 = v109
		v1276 = v1243
		v1277 = v1244
		goto L379
	} else {
		goto L382
	}
L382:
	;
	v1252 = v109
	v1253 = v1243
	v1254 = v1244
	goto L383
L383:
	;
	v1257 = *(*int32)(unsafe.Add(mBase, uint32(v1252)))
	v1258 = *(*int32)(unsafe.Add(mBase, uint32(v1253)))
	if v1257 != v1258 {
		v1275 = v1252
		v1276 = v1253
		v1277 = v1254
		goto L379
	} else {
		goto L385
	}
L384:
	;
	goto L380
L385:
	;
	v1260 = int32(4)
	v1261 = v1253 + v1260
	v1263 = v1252 + v1260
	v1265 = v1254 + int32(-4)
	if base.Ui32(int32(3)) < base.Ui32(v1265) {
		v1252 = v1263
		v1253 = v1261
		v1254 = v1265
		goto L383
	} else {
		goto L386
	}
L386:
	;
	goto L384
L387:
	;
	v1275 = v1263
	v1276 = v1261
	v1277 = v1265
	goto L379
L388:
	;
	v1308 = v1285 - v1286
	goto L377
L389:
	;
	v1285 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1280))))
	v1286 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1281))))
	if v1285 != v1286 {
		goto L388
	} else {
		goto L391
	}
L391:
	;
	v1288 = int32(1)
	v1293 = v1282 + int32(-1)
	if v1293 == int32(0) {
		goto L378
	} else {
		goto L392
	}
L392:
	;
	v1280 = v1280 + v1288
	v1281 = v1281 + v1288
	v1282 = v1293
	goto L389
L393:
	;
	v1310 = v109 + int32(19)
	v1311 = int32(_a1429)
	v1314 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1310))))
	if v1314 != 0 {
		goto L396
	} else {
		goto L397
	}
L394:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+204)) = base.B2i32(v1346-v1348 != int32(0))
	v1353 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v111))))
	v1354 = v1353
	goto L369
L395:
	;
	v1346 = F_tolower(m, v1342)
	mBase = m.M
	v1347 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1343))))
	v1348 = F_tolower(m, v1347)
	mBase = m.M
	goto L394
L396:
	;
	v1316 = v1310
	v1317 = v1311
	v1318 = v1314
	goto L399
L397:
	;
	v1342 = int32(0)
	v1343 = v1311
	goto L395
L398:
	;
	v1342 = v1339 & int32(255)
	v1343 = v1338
	goto L395
L399:
	;
	v1320 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1317))))
	if v1320 == int32(0) {
		v1338 = v1317
		v1339 = v1318
		goto L398
	} else {
		goto L401
	}
L400:
	;
	v1338 = v1332
	v1339 = int32(0)
	goto L398
L401:
	;
	v1324 = v1318 & int32(255)
	if v1324 == v1320 {
		goto L402
	} else {
		goto L403
	}
L402:
	;
	v1331 = int32(1)
	v1332 = v1317 + v1331
	v1333 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1316)+1)))
	if v1333 != 0 {
		v1316 = v1316 + v1331
		v1317 = v1332
		v1318 = v1333
		goto L399
	} else {
		goto L405
	}
L403:
	;
	v1326 = F_tolower(m, v1324)
	mBase = m.M
	v1327 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1317))))
	v1328 = F_tolower(m, v1327)
	mBase = m.M
	if v1326 == v1328 {
		goto L402
	} else {
		goto L404
	}
L404:
	;
	v1330 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1316))))
	v1338 = v1317
	v1339 = v1330
	goto L398
L405:
	;
	goto L400
L406:
	;
	switch v1496 & int32(7) {
	case 0:
		goto L452
	case 1:
		goto L451
	case 2:
		goto L450
	case 3:
		goto L449
	case 4:
		goto L448
	default:
		v1593 = v1496
		goto L446
	}
L407:
	;
	if base.Ui32(v1374) < base.Ui32(int32(15)) {
		v1496 = v1354
		goto L406
	} else {
		goto L413
	}
L408:
	;
	v1373 = *(*int32)(unsafe.Add(mBase, uint32(v109+int32(-17))))
	v1374 = v1373
	goto L407
L409:
	;
	v1370 = *(*int32)(unsafe.Add(mBase, uint32(v109+int32(-9))))
	v1374 = v1370
	goto L407
L410:
	;
	v1367 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v109+int32(-5)))))
	v1374 = v1367
	goto L407
L411:
	;
	v1364 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v109+int32(-3)))))
	v1374 = v1364
	goto L407
L412:
	;
	v1374 = int32(base.Ui32(v1354&int32(248)) >> (uint(int32(3)) % 32))
	goto L407
L413:
	;
	v1377 = int32(_a1430)
	v1378 = int32(15)
	goto L418
L414:
	;
	if v1442 != 0 {
		v1496 = v1354
		goto L406
	} else {
		goto L430
	}
L415:
	;
	v1442 = int32(0)
	goto L414
L416:
	;
	v1414 = v1409
	v1415 = v1410
	v1416 = v1411
	goto L426
L417:
	;
	if v1399 == int32(0) {
		goto L415
	} else {
		goto L424
	}
L418:
	;
	if (v1377|v109)&int32(3) != 0 {
		v1409 = v109
		v1410 = v1377
		v1411 = v1378
		goto L416
	} else {
		goto L419
	}
L419:
	;
	v1386 = v109
	v1387 = v1377
	v1388 = v1378
	goto L420
L420:
	;
	v1391 = *(*int32)(unsafe.Add(mBase, uint32(v1386)))
	v1392 = *(*int32)(unsafe.Add(mBase, uint32(v1387)))
	if v1391 != v1392 {
		v1409 = v1386
		v1410 = v1387
		v1411 = v1388
		goto L416
	} else {
		goto L422
	}
L421:
	;
	goto L417
L422:
	;
	v1394 = int32(4)
	v1395 = v1387 + v1394
	v1397 = v1386 + v1394
	v1399 = v1388 + int32(-4)
	if base.Ui32(int32(3)) < base.Ui32(v1399) {
		v1386 = v1397
		v1387 = v1395
		v1388 = v1399
		goto L420
	} else {
		goto L423
	}
L423:
	;
	goto L421
L424:
	;
	v1409 = v1397
	v1410 = v1395
	v1411 = v1399
	goto L416
L425:
	;
	v1442 = v1419 - v1420
	goto L414
L426:
	;
	v1419 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1414))))
	v1420 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1415))))
	if v1419 != v1420 {
		goto L425
	} else {
		goto L428
	}
L428:
	;
	v1422 = int32(1)
	v1427 = v1416 + int32(-1)
	if v1427 == int32(0) {
		goto L415
	} else {
		goto L429
	}
L429:
	;
	v1414 = v1414 + v1422
	v1415 = v1415 + v1422
	v1416 = v1427
	goto L426
L430:
	;
	v1448 = v109 + int32(15)
	goto L432
L431:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+176)) = v1493
	v1495 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v111))))
	v1496 = v1495
	goto L406
L432:
	;
	v1453 = v1448 + int32(1)
	v1454 = int32(*(*int8)(unsafe.Add(mBase, uint32(v1448))))
	v1455 = F___isspace_1(m, v1454)
	mBase = m.M
	if v1455 != 0 {
		v1448 = v1453
		goto L432
	} else {
		goto L434
	}
L433:
	;
	v1456 = int32(1)
	switch v1454&int32(255) + int32(-43) {
	case 0:
		v1462 = v1456
		goto L436
	default:
		v1464 = v1448
		v1465 = v1454
		v1466 = v1456
		goto L435
	case 2:
		goto L437
	}
L434:
	;
	goto L433
L435:
	;
	v1469 = v1465 + int32(-48)
	if base.Ui32(int32(9)) < base.Ui32(v1469) {
		v1487 = int32(0)
		goto L438
	} else {
		goto L439
	}
L436:
	;
	v1463 = int32(*(*int8)(unsafe.Add(mBase, uint32(v1453))))
	v1464 = v1453
	v1465 = v1463
	v1466 = v1462
	goto L435
L437:
	;
	v1462 = int32(0)
	goto L436
L438:
	;
	if v1466 != 0 {
		goto L443
	} else {
		goto L444
	}
L439:
	;
	v1473 = int32(0)
	v1474 = v1464
	v1475 = v1469
	goto L440
L440:
	;
	v1477 = int32(10)
	v1479 = v1473*v1477 - v1475
	v1480 = int32(*(*int8)(unsafe.Add(mBase, uint32(v1474)+1)))
	v1484 = v1480 + int32(-48)
	if base.Ui32(v1484) < base.Ui32(v1477) {
		v1473 = v1479
		v1474 = v1474 + int32(1)
		v1475 = v1484
		goto L440
	} else {
		goto L442
	}
L441:
	;
	v1487 = v1479
	goto L438
L442:
	;
	goto L441
L443:
	;
	v1493 = int32(0) - v1487
	goto L445
L444:
	;
	v1493 = v1487
	goto L445
L445:
	;
	goto L431
L446:
	;
	v1595 = int32(2)
	switch v1593 & int32(7) {
	case 0:
		goto L477
	case 1:
		goto L476
	case 2:
		goto L475
	case 3:
		goto L474
	case 4:
		goto L473
	default:
		v1736 = v1593
		v1737 = v1595
		goto L232
	}
L447:
	;
	if base.Ui32(v1516) < base.Ui32(int32(18)) {
		v1593 = v1496
		goto L446
	} else {
		goto L453
	}
L448:
	;
	v1515 = *(*int32)(unsafe.Add(mBase, uint32(v109+int32(-17))))
	v1516 = v1515
	goto L447
L449:
	;
	v1512 = *(*int32)(unsafe.Add(mBase, uint32(v109+int32(-9))))
	v1516 = v1512
	goto L447
L450:
	;
	v1509 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v109+int32(-5)))))
	v1516 = v1509
	goto L447
L451:
	;
	v1506 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v109+int32(-3)))))
	v1516 = v1506
	goto L447
L452:
	;
	v1516 = int32(base.Ui32(v1496&int32(248)) >> (uint(int32(3)) % 32))
	goto L447
L453:
	;
	v1519 = int32(_a1431)
	v1520 = int32(18)
	goto L458
L454:
	;
	if v1584 != 0 {
		v1593 = v1496
		goto L446
	} else {
		goto L470
	}
L455:
	;
	v1584 = int32(0)
	goto L454
L456:
	;
	v1556 = v1551
	v1557 = v1552
	v1558 = v1553
	goto L466
L457:
	;
	if v1541 == int32(0) {
		goto L455
	} else {
		goto L464
	}
L458:
	;
	if (v1519|v109)&int32(3) != 0 {
		v1551 = v109
		v1552 = v1519
		v1553 = v1520
		goto L456
	} else {
		goto L459
	}
L459:
	;
	v1528 = v109
	v1529 = v1519
	v1530 = v1520
	goto L460
L460:
	;
	v1533 = *(*int32)(unsafe.Add(mBase, uint32(v1528)))
	v1534 = *(*int32)(unsafe.Add(mBase, uint32(v1529)))
	if v1533 != v1534 {
		v1551 = v1528
		v1552 = v1529
		v1553 = v1530
		goto L456
	} else {
		goto L462
	}
L461:
	;
	goto L457
L462:
	;
	v1536 = int32(4)
	v1537 = v1529 + v1536
	v1539 = v1528 + v1536
	v1541 = v1530 + int32(-4)
	if base.Ui32(int32(3)) < base.Ui32(v1541) {
		v1528 = v1539
		v1529 = v1537
		v1530 = v1541
		goto L460
	} else {
		goto L463
	}
L463:
	;
	goto L461
L464:
	;
	v1551 = v1539
	v1552 = v1537
	v1553 = v1541
	goto L456
L465:
	;
	v1584 = v1561 - v1562
	goto L454
L466:
	;
	v1561 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1556))))
	v1562 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1557))))
	if v1561 != v1562 {
		goto L465
	} else {
		goto L468
	}
L468:
	;
	v1564 = int32(1)
	v1569 = v1558 + int32(-1)
	if v1569 == int32(0) {
		goto L455
	} else {
		goto L469
	}
L469:
	;
	v1556 = v1556 + v1564
	v1557 = v1557 + v1564
	v1558 = v1569
	goto L466
L470:
	;
	v1590 = F_strtox_2(m, v109+int32(18), int32(0), int32(10), int64(-1))
	mBase = m.M
	goto L471
L471:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l0)+208)) = v1590
	v1592 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v111))))
	v1593 = v1592
	goto L446
L472:
	;
	if base.Ui32(v1614) < base.Ui32(int32(18)) {
		v1736 = v1593
		v1737 = v1595
		goto L232
	} else {
		goto L478
	}
L473:
	;
	v1613 = *(*int32)(unsafe.Add(mBase, uint32(v109+int32(-17))))
	v1614 = v1613
	goto L472
L474:
	;
	v1610 = *(*int32)(unsafe.Add(mBase, uint32(v109+int32(-9))))
	v1614 = v1610
	goto L472
L475:
	;
	v1607 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v109+int32(-5)))))
	v1614 = v1607
	goto L472
L476:
	;
	v1604 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v109+int32(-3)))))
	v1614 = v1604
	goto L472
L477:
	;
	v1614 = int32(base.Ui32(v1593&int32(248)) >> (uint(int32(3)) % 32))
	goto L472
L478:
	;
	v1617 = int32(_a1432)
	v1618 = int32(18)
	goto L483
L479:
	;
	if v1682 != 0 {
		v1736 = v1593
		v1737 = v1595
		goto L232
	} else {
		goto L495
	}
L480:
	;
	v1682 = int32(0)
	goto L479
L481:
	;
	v1654 = v1649
	v1655 = v1650
	v1656 = v1651
	goto L491
L482:
	;
	if v1639 == int32(0) {
		goto L480
	} else {
		goto L489
	}
L483:
	;
	if (v1617|v109)&int32(3) != 0 {
		v1649 = v109
		v1650 = v1617
		v1651 = v1618
		goto L481
	} else {
		goto L484
	}
L484:
	;
	v1626 = v109
	v1627 = v1617
	v1628 = v1618
	goto L485
L485:
	;
	v1631 = *(*int32)(unsafe.Add(mBase, uint32(v1626)))
	v1632 = *(*int32)(unsafe.Add(mBase, uint32(v1627)))
	if v1631 != v1632 {
		v1649 = v1626
		v1650 = v1627
		v1651 = v1628
		goto L481
	} else {
		goto L487
	}
L486:
	;
	goto L482
L487:
	;
	v1634 = int32(4)
	v1635 = v1627 + v1634
	v1637 = v1626 + v1634
	v1639 = v1628 + int32(-4)
	if base.Ui32(int32(3)) < base.Ui32(v1639) {
		v1626 = v1637
		v1627 = v1635
		v1628 = v1639
		goto L485
	} else {
		goto L488
	}
L488:
	;
	goto L486
L489:
	;
	v1649 = v1637
	v1650 = v1635
	v1651 = v1639
	goto L481
L490:
	;
	v1682 = v1659 - v1660
	goto L479
L491:
	;
	v1659 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1654))))
	v1660 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1655))))
	if v1659 != v1660 {
		goto L490
	} else {
		goto L493
	}
L493:
	;
	v1662 = int32(1)
	v1667 = v1656 + int32(-1)
	if v1667 == int32(0) {
		goto L480
	} else {
		goto L494
	}
L494:
	;
	v1654 = v1654 + v1662
	v1655 = v1655 + v1662
	v1656 = v1667
	goto L491
L495:
	;
	v1688 = v109 + int32(18)
	goto L497
L496:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+180)) = v1733
	v1735 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v111))))
	v1736 = v1735
	v1737 = v1595
	goto L232
L497:
	;
	v1693 = v1688 + int32(1)
	v1694 = int32(*(*int8)(unsafe.Add(mBase, uint32(v1688))))
	v1695 = F___isspace_1(m, v1694)
	mBase = m.M
	if v1695 != 0 {
		v1688 = v1693
		goto L497
	} else {
		goto L499
	}
L498:
	;
	v1696 = int32(1)
	switch v1694&int32(255) + int32(-43) {
	case 0:
		v1702 = v1696
		goto L501
	default:
		v1704 = v1688
		v1705 = v1694
		v1706 = v1696
		goto L500
	case 2:
		goto L502
	}
L499:
	;
	goto L498
L500:
	;
	v1709 = v1705 + int32(-48)
	if base.Ui32(int32(9)) < base.Ui32(v1709) {
		v1727 = int32(0)
		goto L503
	} else {
		goto L504
	}
L501:
	;
	v1703 = int32(*(*int8)(unsafe.Add(mBase, uint32(v1693))))
	v1704 = v1693
	v1705 = v1703
	v1706 = v1702
	goto L500
L502:
	;
	v1702 = int32(0)
	goto L501
L503:
	;
	if v1706 != 0 {
		goto L508
	} else {
		goto L509
	}
L504:
	;
	v1713 = int32(0)
	v1714 = v1704
	v1715 = v1709
	goto L505
L505:
	;
	v1717 = int32(10)
	v1719 = v1713*v1717 - v1715
	v1720 = int32(*(*int8)(unsafe.Add(mBase, uint32(v1714)+1)))
	v1724 = v1720 + int32(-48)
	if base.Ui32(v1724) < base.Ui32(v1717) {
		v1713 = v1719
		v1714 = v1714 + int32(1)
		v1715 = v1724
		goto L505
	} else {
		goto L507
	}
L506:
	;
	v1727 = v1719
	goto L503
L507:
	;
	goto L506
L508:
	;
	v1733 = int32(0) - v1727
	goto L510
L509:
	;
	v1733 = v1727
	goto L510
L510:
	;
	goto L496
L511:
	;
	if base.Ui32(v1758) < base.Ui32(int32(22)) {
		v1920 = v1737
		goto L72
	} else {
		goto L517
	}
L512:
	;
	v1757 = *(*int32)(unsafe.Add(mBase, uint32(v109+int32(-17))))
	v1758 = v1757
	goto L511
L513:
	;
	v1754 = *(*int32)(unsafe.Add(mBase, uint32(v109+int32(-9))))
	v1758 = v1754
	goto L511
L514:
	;
	v1751 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v109+int32(-5)))))
	v1758 = v1751
	goto L511
L515:
	;
	v1748 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v109+int32(-3)))))
	v1758 = v1748
	goto L511
L516:
	;
	v1758 = int32(base.Ui32(v1736&int32(248)) >> (uint(int32(3)) % 32))
	goto L511
L517:
	;
	v1761 = int32(_a1433)
	v1762 = int32(22)
	goto L522
L518:
	;
	if v1826 != 0 {
		v1920 = v1737
		goto L72
	} else {
		goto L534
	}
L519:
	;
	v1826 = int32(0)
	goto L518
L520:
	;
	v1798 = v1793
	v1799 = v1794
	v1800 = v1795
	goto L530
L521:
	;
	if v1783 == int32(0) {
		goto L519
	} else {
		goto L528
	}
L522:
	;
	if (v1761|v109)&int32(3) != 0 {
		v1793 = v109
		v1794 = v1761
		v1795 = v1762
		goto L520
	} else {
		goto L523
	}
L523:
	;
	v1770 = v109
	v1771 = v1761
	v1772 = v1762
	goto L524
L524:
	;
	v1775 = *(*int32)(unsafe.Add(mBase, uint32(v1770)))
	v1776 = *(*int32)(unsafe.Add(mBase, uint32(v1771)))
	if v1775 != v1776 {
		v1793 = v1770
		v1794 = v1771
		v1795 = v1772
		goto L520
	} else {
		goto L526
	}
L525:
	;
	goto L521
L526:
	;
	v1778 = int32(4)
	v1779 = v1771 + v1778
	v1781 = v1770 + v1778
	v1783 = v1772 + int32(-4)
	if base.Ui32(int32(3)) < base.Ui32(v1783) {
		v1770 = v1781
		v1771 = v1779
		v1772 = v1783
		goto L524
	} else {
		goto L527
	}
L527:
	;
	goto L525
L528:
	;
	v1793 = v1781
	v1794 = v1779
	v1795 = v1783
	goto L520
L529:
	;
	v1826 = v1803 - v1804
	goto L518
L530:
	;
	v1803 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1798))))
	v1804 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1799))))
	if v1803 != v1804 {
		goto L529
	} else {
		goto L532
	}
L532:
	;
	v1806 = int32(1)
	v1811 = v1800 + int32(-1)
	if v1811 == int32(0) {
		goto L519
	} else {
		goto L533
	}
L533:
	;
	v1798 = v1798 + v1806
	v1799 = v1799 + v1806
	v1800 = v1811
	goto L530
L534:
	;
	v1827 = int32(2)
	switch v1741 {
	default:
		v1912 = v1827
		goto L535
	case 1:
		goto L540
	case 2:
		goto L539
	case 3:
		goto L538
	case 4:
		goto L537
	}
L535:
	;
	v1914 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
	if v1912 == v1914 {
		v1920 = v1737
		goto L72
	} else {
		goto L561
	}
L536:
	;
	if base.Ui32(v1840) < base.Ui32(int32(33)) {
		v1912 = v1827
		goto L535
	} else {
		goto L541
	}
L537:
	;
	v1839 = *(*int32)(unsafe.Add(mBase, uint32(v109+int32(-17))))
	v1840 = v1839
	goto L536
L538:
	;
	v1836 = *(*int32)(unsafe.Add(mBase, uint32(v109+int32(-9))))
	v1840 = v1836
	goto L536
L539:
	;
	v1833 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v109+int32(-5)))))
	v1840 = v1833
	goto L536
L540:
	;
	v1830 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v109+int32(-3)))))
	v1840 = v1830
	goto L536
L541:
	;
	v1845 = int32(_a1434)
	v1846 = int32(33)
	goto L546
L542:
	;
	if v1910 != 0 {
		goto L558
	} else {
		goto L559
	}
L543:
	;
	v1910 = int32(0)
	goto L542
L544:
	;
	v1882 = v1877
	v1883 = v1878
	v1884 = v1879
	goto L554
L545:
	;
	if v1867 == int32(0) {
		goto L543
	} else {
		goto L552
	}
L546:
	;
	if (v1845|v109)&int32(3) != 0 {
		v1877 = v109
		v1878 = v1845
		v1879 = v1846
		goto L544
	} else {
		goto L547
	}
L547:
	;
	v1854 = v109
	v1855 = v1845
	v1856 = v1846
	goto L548
L548:
	;
	v1859 = *(*int32)(unsafe.Add(mBase, uint32(v1854)))
	v1860 = *(*int32)(unsafe.Add(mBase, uint32(v1855)))
	if v1859 != v1860 {
		v1877 = v1854
		v1878 = v1855
		v1879 = v1856
		goto L544
	} else {
		goto L550
	}
L549:
	;
	goto L545
L550:
	;
	v1862 = int32(4)
	v1863 = v1855 + v1862
	v1865 = v1854 + v1862
	v1867 = v1856 + int32(-4)
	if base.Ui32(int32(3)) < base.Ui32(v1867) {
		v1854 = v1865
		v1855 = v1863
		v1856 = v1867
		goto L548
	} else {
		goto L551
	}
L551:
	;
	goto L549
L552:
	;
	v1877 = v1865
	v1878 = v1863
	v1879 = v1867
	goto L544
L553:
	;
	v1910 = v1887 - v1888
	goto L542
L554:
	;
	v1887 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1882))))
	v1888 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1883))))
	if v1887 != v1888 {
		goto L553
	} else {
		goto L556
	}
L556:
	;
	v1890 = int32(1)
	v1895 = v1884 + int32(-1)
	if v1895 == int32(0) {
		goto L543
	} else {
		goto L557
	}
L557:
	;
	v1882 = v1882 + v1890
	v1883 = v1883 + v1890
	v1884 = v1895
	goto L554
L558:
	;
	v1911 = int32(2)
	goto L560
L559:
	;
	v1911 = int32(1)
	goto L560
L560:
	;
	v1912 = v1911
	goto L535
L561:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+128)) = v1912
	v1917 = F_mstime(m)
	mBase = m.M
	*(*int64)(unsafe.Add(mBase, uint32(l0)+136)) = v1917
	v1920 = v1737
	goto L72
L562:
	;
	goto L25
L563:
	;
	v1946 = *(*int32)(unsafe.Add(mBase, uint32(l0)+108))
	if v1933 == v1946 {
		goto L564
	} else {
		goto L565
	}
L564:
	;
	v1977 = *(*int32)(unsafe.Add(mBase, _consts[764]))
	if v1977 != 0 {
		goto L576
	} else {
		goto L577
	}
L565:
	;
	v1948 = F_mstime(m)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(l0)+108)) = v1933
	*(*int64)(unsafe.Add(mBase, uint32(l0)+112)) = v1948
	if v1933 != int32(2) {
		goto L566
	} else {
		goto L567
	}
L566:
	;
	v1955 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v1933 == int32(1) {
		goto L568
	} else {
		goto L569
	}
L567:
	;
	v1953 = F_mstime(m)
	mBase = m.M
	*(*int64)(unsafe.Add(mBase, uint32(l0)+120)) = v1953
	goto L566
L568:
	;
	v1960 = int32(_a349)
	goto L570
L569:
	;
	v1960 = int32(_a505)
	goto L570
L570:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+48)) = v1960
	if v1955&int32(3) == v1933 {
		goto L571
	} else {
		goto L572
	}
L571:
	;
	v1968 = int32(_a1435)
	goto L573
L572:
	;
	v1968 = int32(_a1436)
	goto L573
L573:
	;
	F_sentinelEvent(m, int32(1), v1968, l0, int32(_a1437), v16+int32(48))
	mBase = m.M
	v1973 = m.ExcPending
	if v1973 != 0 {
		goto L1
	} else {
		goto L574
	}
L574:
	;
	goto L564
L575:
	;
	F_sentinelSimFailureCrash(m)
	mBase = m.M
	v2575 = m.ExcPending
	if v2575 != 0 {
		goto L1
	} else {
		goto L726
	}
L576:
	;
	m.G0 = v16 + int32(160)
	return
L577:
	;
	v1978 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1980 = v1978 & int32(2)
	if v1980 == int32(0) {
		goto L579
	} else {
		goto L580
	}
L578:
	;
	v2441 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v2441&int32(768) == int32(0) {
		goto L576
	} else {
		goto L695
	}
L579:
	;
	if v1980 == int32(0) {
		goto L639
	} else {
		goto L640
	}
L580:
	;
	if v1933 != int32(1) {
		goto L579
	} else {
		goto L581
	}
L581:
	;
	if v1978&int32(128) == int32(0) {
		goto L582
	} else {
		goto L583
	}
L582:
	;
	v2209 = *(*int32)(unsafe.Add(mBase, uint32(l0)+192))
	v2210 = *(*int32)(unsafe.Add(mBase, uint32(v2209)))
	if v2210&int32(1) == int32(0) {
		goto L578
	} else {
		goto L625
	}
L583:
	;
	v1989 = *(*int32)(unsafe.Add(mBase, uint32(l0)+192))
	v1990 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1989))))
	if v1990&int32(64) == int32(0) {
		goto L578
	} else {
		goto L584
	}
L584:
	;
	v1995 = *(*int32)(unsafe.Add(mBase, uint32(v1989)+240))
	if v1995 != int32(4) {
		goto L578
	} else {
		goto L585
	}
L585:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1989)+240)) = int32(5)
	v2000 = *(*int64)(unsafe.Add(mBase, uint32(v1989)+232))
	*(*int64)(unsafe.Add(mBase, uint32(v1989)+16)) = v2000
	v2002 = F_mstime(m)
	mBase = m.M
	v2003 = *(*int32)(unsafe.Add(mBase, uint32(l0)+192))
	*(*int64)(unsafe.Add(mBase, uint32(v2003)+248)) = v2002
	v2005 = int32(_a44)
	v2006 = *(*int32)(unsafe.Add(mBase, _consts[219]))
	*(*int32)(unsafe.Add(mBase, _consts[219])) = int32(10)
	v2011 = *(*int32)(unsafe.Add(mBase, _consts[392]))
	v2013 = F_rewriteConfig(m, v2011, int32(0))
	mBase = m.M
	v2014 = m.ExcPending
	if v2014 != 0 {
		goto L1
	} else {
		goto L586
	}
L586:
	;
	v2015 = int32(_a44)
	*(*int32)(unsafe.Add(mBase, _consts[219])) = v2006
	v2018 = *(*int32)(unsafe.Add(mBase, _consts[6]))
	if v2013 != int32(-1) {
		goto L588
	} else {
		goto L589
	}
L587:
	;
	F_sentinelEvent(m, int32(3), int32(_a1438), l0, int32(_a1362), int32(0))
	mBase = m.M
	v2045 = m.ExcPending
	if v2045 != 0 {
		goto L1
	} else {
		goto L596
	}
L588:
	;
	if int32(2) < v2018 {
		goto L587
	} else {
		goto L594
	}
L589:
	;
	if int32(3) < v2018 {
		goto L587
	} else {
		goto L590
	}
L590:
	;
	goto L591
L591:
	;
	v2024 = *(*int32)(unsafe.Add(mBase, _consts[5]))
	v2025 = F___strerror_l(m, v2024, v2024)
	mBase = m.M
	goto L592
L592:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+32)) = v2025
	F__serverLog(m, int32(3), int32(_a1337), v16+int32(32))
	mBase = m.M
	v2032 = m.ExcPending
	if v2032 != 0 {
		goto L1
	} else {
		goto L593
	}
L593:
	;
	goto L587
L594:
	;
	F__serverLog(m, int32(2), int32(_a1338), int32(0))
	mBase = m.M
	v2039 = m.ExcPending
	if v2039 != 0 {
		goto L1
	} else {
		goto L595
	}
L595:
	;
	goto L587
L596:
	;
	v2047 = int32(*(*uint8)(unsafe.Add(mBase, _consts[765])))
	if v2047&int32(2) != 0 {
		goto L575
	} else {
		goto L597
	}
L597:
	;
	v2050 = *(*int32)(unsafe.Add(mBase, uint32(l0)+192))
	v2051 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2050)+1)))
	if v2051&int32(64) == int32(0) {
		v2061 = v2050
		goto L598
	} else {
		goto L599
	}
L598:
	;
	F_sentinelEvent(m, int32(3), int32(_a1439), v2061, int32(_a1362), int32(0))
	mBase = m.M
	v2067 = m.ExcPending
	if v2067 != 0 {
		goto L1
	} else {
		goto L602
	}
L599:
	;
	v2056 = F_sentinelKillClients(m, v2050)
	mBase = m.M
	v2057 = m.ExcPending
	if v2057 != 0 {
		goto L1
	} else {
		goto L600
	}
L600:
	;
	v2058 = F_sentinelKillClients(m, l0)
	mBase = m.M
	v2059 = m.ExcPending
	if v2059 != 0 {
		goto L1
	} else {
		goto L601
	}
L601:
	;
	v2060 = *(*int32)(unsafe.Add(mBase, uint32(l0)+192))
	v2061 = v2060
	goto L598
L602:
	;
	v2068 = *(*int32)(unsafe.Add(mBase, uint32(l0)+192))
	v2069 = *(*int32)(unsafe.Add(mBase, uint32(v2068)+288))
	if v2069 == int32(0) {
		v2202 = v2068
		goto L603
	} else {
		goto L604
	}
L603:
	;
	v2207 = F_sentinelForceHelloUpdateForPrimary(m, v2202)
	mBase = m.M
	v2208 = m.ExcPending
	if v2208 != 0 {
		goto L1
	} else {
		goto L624
	}
L604:
	;
	v2072 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v2074 = v16 + int32(112)
	v2076 = *(*int32)(unsafe.Add(mBase, uint32(v2068)+24))
	v2077 = int64(*(*int32)(unsafe.Add(mBase, uint32(v2076)+8)))
	if v2077 <= int64(-1) {
		goto L609
	} else {
		goto L610
	}
L605:
	;
	v2120 = v16 + int32(80)
	v2122 = int64(*(*int32)(unsafe.Add(mBase, uint32(v2072)+8)))
	if v2122 <= int64(-1) {
		goto L618
	} else {
		goto L619
	}
L606:
	;
	goto L605
L608:
	;
	v2099 = F_ull2string(m, v2095, v2096, v2097)
	mBase = m.M
	if v2099 == int32(0) {
		goto L606
	} else {
		goto L612
	}
L609:
	;
	goto L611
L610:
	;
	v2095 = v2074
	v2096 = int32(32)
	v2097 = v2077
	goto L608
L611:
	;
	v2086 = int32(45)
	*(*uint8)(unsafe.Add(mBase, uint32(v2074))) = uint8(v2086)
	v2095 = v16 + int32(113)
	v2096 = int32(31)
	v2097 = int64(0) - v2077
	goto L608
L612:
	;
	goto L605
L614:
	;
	v2164 = int32(0)
	v2165 = *(*int32)(unsafe.Add(mBase, _consts[734]))
	v2169 = base.B2i32(v2165 == v2164) << (uint(int32(2)) % 32)
	v2171 = *(*int32)(unsafe.Add(mBase, uint32(v2076+v2169)))
	v2172 = *(*int32)(unsafe.Add(mBase, uint32(v2068)+288))
	v2173 = *(*int32)(unsafe.Add(mBase, uint32(v2068)+4))
	v2177 = *(*int32)(unsafe.Add(mBase, uint32(v2072+v2169)))
	*(*int32)(unsafe.Add(mBase, uint32(v16+int32(20)))) = v2177
	*(*int32)(unsafe.Add(mBase, uint32(v16+int32(28)))) = v2164
	*(*int32)(unsafe.Add(mBase, uint32(v16))) = v2173
	*(*int32)(unsafe.Add(mBase, uint32(v16)+4)) = int32(_a1440)
	*(*int32)(unsafe.Add(mBase, uint32(v16)+8)) = int32(_a1441)
	*(*int32)(unsafe.Add(mBase, uint32(v16)+12)) = v2171
	*(*int32)(unsafe.Add(mBase, uint32(v16+int32(16)))) = v16 + int32(112)
	*(*int32)(unsafe.Add(mBase, uint32(v16+int32(24)))) = v16 + int32(80)
	F_sentinelScheduleScriptExecution(m, v2172, v16)
	mBase = m.M
	v2200 = m.ExcPending
	if v2200 != 0 {
		goto L1
	} else {
		goto L623
	}
L615:
	;
	goto L614
L617:
	;
	v2144 = F_ull2string(m, v2140, v2141, v2142)
	mBase = m.M
	if v2144 == int32(0) {
		goto L615
	} else {
		goto L621
	}
L618:
	;
	goto L620
L619:
	;
	v2140 = v2120
	v2141 = int32(32)
	v2142 = v2122
	goto L617
L620:
	;
	v2131 = int32(45)
	*(*uint8)(unsafe.Add(mBase, uint32(v2120))) = uint8(v2131)
	v2140 = v16 + int32(81)
	v2141 = int32(31)
	v2142 = int64(0) - v2122
	goto L617
L621:
	;
	goto L614
L623:
	;
	v2201 = *(*int32)(unsafe.Add(mBase, uint32(l0)+192))
	v2202 = v2201
	goto L603
L624:
	;
	goto L578
L625:
	;
	if v2210&int32(24) != 0 {
		goto L578
	} else {
		goto L626
	}
L626:
	;
	v2217 = *(*int32)(unsafe.Add(mBase, uint32(v2209)+108))
	if v2217 != int32(1) {
		goto L578
	} else {
		goto L627
	}
L627:
	;
	v2220 = int32(0)
	v2221 = *(*int64)(unsafe.Add(mBase, _consts[766]))
	v2222 = F_mstime(m)
	mBase = m.M
	v2223 = *(*int64)(unsafe.Add(mBase, uint32(v2209)+96))
	v2226 = *(*int64)(unsafe.Add(mBase, _consts[767]))
	if v2226<<(uint(int64(1))%64) <= v2222-v2223 {
		goto L578
	} else {
		goto L628
	}
L628:
	;
	v2231 = v2221 << (uint(int64(2)) % 64)
	v2232 = *(*int64)(unsafe.Add(mBase, uint32(l0)+64))
	v2233 = *(*int64)(unsafe.Add(mBase, uint32(l0)+56))
	if v2233 < v2232 {
		goto L630
	} else {
		goto L631
	}
L629:
	;
	v2241 = F_mstime(m)
	mBase = m.M
	v2242 = *(*int64)(unsafe.Add(mBase, uint32(l0)+112))
	if v2241-v2242 <= v2231 {
		goto L578
	} else {
		goto L635
	}
L630:
	;
	v2235 = v2232
	goto L632
L631:
	;
	v2235 = v2233
	goto L632
L632:
	;
	if v2235 == int64(0) {
		goto L629
	} else {
		goto L633
	}
L633:
	;
	v2238 = F_mstime(m)
	mBase = m.M
	if v2238-v2235 <= v2231 {
		goto L578
	} else {
		goto L634
	}
L634:
	;
	goto L629
L635:
	;
	v2245 = *(*int32)(unsafe.Add(mBase, uint32(l0)+192))
	v2246 = *(*int32)(unsafe.Add(mBase, uint32(v2245)+24))
	v2247 = F_sentinelSendReplicaOf(m, l0, v2246)
	mBase = m.M
	v2248 = m.ExcPending
	if v2248 != 0 {
		goto L1
	} else {
		goto L636
	}
L636:
	;
	if v2247 != 0 {
		goto L578
	} else {
		goto L637
	}
L637:
	;
	F_sentinelEvent(m, int32(2), int32(_a1442), l0, int32(_a1362), int32(0))
	mBase = m.M
	v2254 = m.ExcPending
	if v2254 != 0 {
		goto L1
	} else {
		goto L638
	}
L638:
	;
	goto L578
L639:
	;
	v2310 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v2310&int32(2) == int32(0) {
		goto L578
	} else {
		goto L657
	}
L640:
	;
	if v1933 != int32(2) {
		goto L639
	} else {
		goto L641
	}
L641:
	;
	v2259 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
	if v2259 != int32(2) {
		goto L639
	} else {
		goto L642
	}
L642:
	;
	v2262 = *(*int32)(unsafe.Add(mBase, uint32(l0)+192))
	v2263 = *(*int32)(unsafe.Add(mBase, uint32(v2262)))
	if v2263&int32(1) == int32(0) {
		goto L639
	} else {
		goto L643
	}
L643:
	;
	if v2263&int32(24) != 0 {
		goto L639
	} else {
		goto L644
	}
L644:
	;
	v2270 = *(*int32)(unsafe.Add(mBase, uint32(v2262)+108))
	if v2270 != int32(1) {
		goto L639
	} else {
		goto L645
	}
L645:
	;
	v2273 = *(*int64)(unsafe.Add(mBase, uint32(v2262)+264))
	v2274 = F_mstime(m)
	mBase = m.M
	v2275 = *(*int64)(unsafe.Add(mBase, uint32(v2262)+96))
	v2278 = *(*int64)(unsafe.Add(mBase, _consts[767]))
	if v2278<<(uint(int64(1))%64) <= v2274-v2275 {
		goto L639
	} else {
		goto L646
	}
L646:
	;
	v2282 = *(*int64)(unsafe.Add(mBase, uint32(l0)+64))
	v2283 = *(*int64)(unsafe.Add(mBase, uint32(l0)+56))
	if v2283 < v2282 {
		goto L648
	} else {
		goto L649
	}
L647:
	;
	v2291 = F_mstime(m)
	mBase = m.M
	v2292 = *(*int64)(unsafe.Add(mBase, uint32(l0)+136))
	if v2291-v2292 <= v2273 {
		goto L639
	} else {
		goto L653
	}
L648:
	;
	v2285 = v2282
	goto L650
L649:
	;
	v2285 = v2283
	goto L650
L650:
	;
	if v2285 == int64(0) {
		goto L647
	} else {
		goto L651
	}
L651:
	;
	v2288 = F_mstime(m)
	mBase = m.M
	if v2288-v2285 <= v2273 {
		goto L639
	} else {
		goto L652
	}
L652:
	;
	goto L647
L653:
	;
	v2295 = *(*int32)(unsafe.Add(mBase, uint32(l0)+192))
	v2296 = *(*int32)(unsafe.Add(mBase, uint32(v2295)+24))
	v2297 = F_sentinelSendReplicaOf(m, l0, v2296)
	mBase = m.M
	v2298 = m.ExcPending
	if v2298 != 0 {
		goto L1
	} else {
		goto L654
	}
L654:
	;
	if v2297 != 0 {
		goto L639
	} else {
		goto L655
	}
L655:
	;
	F_sentinelEvent(m, int32(2), int32(_a1443), l0, int32(_a1362), int32(0))
	mBase = m.M
	v2304 = m.ExcPending
	if v2304 != 0 {
		goto L1
	} else {
		goto L656
	}
L656:
	;
	goto L639
L657:
	;
	if v1933 != int32(2) {
		goto L578
	} else {
		goto L658
	}
L658:
	;
	v2317 = *(*int32)(unsafe.Add(mBase, uint32(l0)+200))
	v2318 = *(*int32)(unsafe.Add(mBase, uint32(l0)+192))
	v2319 = *(*int32)(unsafe.Add(mBase, uint32(v2318)+24))
	v2320 = *(*int32)(unsafe.Add(mBase, uint32(v2319)+8))
	if v2317 != v2320 {
		v2389 = v2318
		goto L659
	} else {
		goto L660
	}
L659:
	;
	v2391 = *(*int32)(unsafe.Add(mBase, uint32(v2389)))
	if v2391&int32(1) == int32(0) {
		goto L578
	} else {
		goto L681
	}
L660:
	;
	v2322 = int32(0)
	v2323 = *(*int32)(unsafe.Add(mBase, uint32(l0)+196))
	v2328 = *(*int32)(unsafe.Add(mBase, _consts[729]))
	v2331 = F_anetResolve(m, v2322, v2323, v16+int32(112), int32(46), base.B2i32(v2328 == v2322))
	mBase = m.M
	v2332 = m.ExcPending
	if v2332 != 0 {
		goto L1
	} else {
		goto L661
	}
L661:
	;
	v2333 = int32(0)
	v2334 = *(*int32)(unsafe.Add(mBase, _consts[729]))
	v2341 = base.B2i32(v2331 == int32(-1))
	if v2331 == int32(-1) {
		goto L662
	} else {
		goto L663
	}
L662:
	;
	v2342 = base.B2i32(v2334 == v2333) << (uint(int32(2)) % 32)
	goto L664
L663:
	;
	v2342 = int32(4)
	goto L664
L664:
	;
	v2344 = *(*int32)(unsafe.Add(mBase, uint32(v2319+v2342)))
	if v2331 == int32(-1) {
		goto L665
	} else {
		goto L666
	}
L665:
	;
	v2347 = v2323
	goto L667
L666:
	;
	v2347 = v16 + int32(112)
	goto L667
L667:
	;
	v2350 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2344))))
	if v2350 != 0 {
		goto L670
	} else {
		goto L671
	}
L668:
	;
	if v2382-v2384 == int32(0) {
		goto L578
	} else {
		goto L680
	}
L669:
	;
	v2382 = F_tolower(m, v2378)
	mBase = m.M
	v2383 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2379))))
	v2384 = F_tolower(m, v2383)
	mBase = m.M
	goto L668
L670:
	;
	v2352 = v2344
	v2353 = v2347
	v2354 = v2350
	goto L673
L671:
	;
	v2378 = int32(0)
	v2379 = v2347
	goto L669
L672:
	;
	v2378 = v2375 & int32(255)
	v2379 = v2374
	goto L669
L673:
	;
	v2356 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2353))))
	if v2356 == int32(0) {
		v2374 = v2353
		v2375 = v2354
		goto L672
	} else {
		goto L675
	}
L674:
	;
	v2374 = v2368
	v2375 = int32(0)
	goto L672
L675:
	;
	v2360 = v2354 & int32(255)
	if v2360 == v2356 {
		goto L676
	} else {
		goto L677
	}
L676:
	;
	v2367 = int32(1)
	v2368 = v2353 + v2367
	v2369 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2352)+1)))
	if v2369 != 0 {
		v2352 = v2352 + v2367
		v2353 = v2368
		v2354 = v2369
		goto L673
	} else {
		goto L679
	}
L677:
	;
	v2362 = F_tolower(m, v2360)
	mBase = m.M
	v2363 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2353))))
	v2364 = F_tolower(m, v2363)
	mBase = m.M
	if v2362 == v2364 {
		goto L676
	} else {
		goto L678
	}
L678:
	;
	v2366 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2352))))
	v2374 = v2353
	v2375 = v2366
	goto L672
L679:
	;
	goto L674
L680:
	;
	v2388 = *(*int32)(unsafe.Add(mBase, uint32(l0)+192))
	v2389 = v2388
	goto L659
L681:
	;
	if v2391&int32(24) != 0 {
		goto L578
	} else {
		goto L682
	}
L682:
	;
	v2398 = *(*int32)(unsafe.Add(mBase, uint32(v2389)+108))
	if v2398 != int32(1) {
		goto L578
	} else {
		goto L683
	}
L683:
	;
	v2401 = *(*int64)(unsafe.Add(mBase, uint32(v2389)+264))
	v2402 = F_mstime(m)
	mBase = m.M
	v2403 = *(*int64)(unsafe.Add(mBase, uint32(v2389)+96))
	v2406 = *(*int64)(unsafe.Add(mBase, _consts[767]))
	if v2406<<(uint(int64(1))%64) <= v2402-v2403 {
		goto L578
	} else {
		goto L684
	}
L684:
	;
	v2410 = *(*int64)(unsafe.Add(mBase, uint32(l0)+64))
	v2411 = *(*int64)(unsafe.Add(mBase, uint32(l0)+56))
	if v2411 < v2410 {
		goto L686
	} else {
		goto L687
	}
L685:
	;
	v2419 = F_mstime(m)
	mBase = m.M
	v2420 = *(*int64)(unsafe.Add(mBase, uint32(l0)+120))
	if v2419-v2420 <= v2401 {
		goto L578
	} else {
		goto L691
	}
L686:
	;
	v2413 = v2410
	goto L688
L687:
	;
	v2413 = v2411
	goto L688
L688:
	;
	if v2413 == int64(0) {
		goto L685
	} else {
		goto L689
	}
L689:
	;
	v2416 = F_mstime(m)
	mBase = m.M
	if v2416-v2413 <= v2401 {
		goto L578
	} else {
		goto L690
	}
L690:
	;
	goto L685
L691:
	;
	v2423 = *(*int32)(unsafe.Add(mBase, uint32(l0)+192))
	v2424 = *(*int32)(unsafe.Add(mBase, uint32(v2423)+24))
	v2425 = F_sentinelSendReplicaOf(m, l0, v2424)
	mBase = m.M
	v2426 = m.ExcPending
	if v2426 != 0 {
		goto L1
	} else {
		goto L692
	}
L692:
	;
	if v2425 != 0 {
		goto L578
	} else {
		goto L693
	}
L693:
	;
	F_sentinelEvent(m, int32(2), int32(_a1443), l0, int32(_a1362), int32(0))
	mBase = m.M
	v2432 = m.ExcPending
	if v2432 != 0 {
		goto L1
	} else {
		goto L694
	}
L694:
	;
	goto L578
L695:
	;
	if v2441&int32(2) == int32(0) {
		goto L576
	} else {
		goto L696
	}
L696:
	;
	if v1933 != int32(2) {
		goto L576
	} else {
		goto L697
	}
L697:
	;
	if v2441&int32(256) == int32(0) {
		goto L698
	} else {
		goto L699
	}
L698:
	;
	v2546 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v2546&int32(512) == int32(0) {
		goto L576
	} else {
		goto L723
	}
L699:
	;
	v2456 = *(*int32)(unsafe.Add(mBase, uint32(l0)+196))
	if v2456 == int32(0) {
		goto L698
	} else {
		goto L700
	}
L700:
	;
	v2459 = *(*int32)(unsafe.Add(mBase, uint32(l0)+192))
	v2460 = *(*int32)(unsafe.Add(mBase, uint32(v2459)+280))
	v2461 = *(*int32)(unsafe.Add(mBase, uint32(v2460)+24))
	v2462 = int32(0)
	v2467 = *(*int32)(unsafe.Add(mBase, _consts[729]))
	v2470 = F_anetResolve(m, v2462, v2456, v16+int32(112), int32(46), base.B2i32(v2467 == v2462))
	mBase = m.M
	v2471 = m.ExcPending
	if v2471 != 0 {
		goto L1
	} else {
		goto L701
	}
L701:
	;
	v2472 = int32(0)
	v2473 = *(*int32)(unsafe.Add(mBase, _consts[729]))
	v2480 = base.B2i32(v2470 == int32(-1))
	if v2470 == int32(-1) {
		goto L702
	} else {
		goto L703
	}
L702:
	;
	v2481 = base.B2i32(v2473 == v2472) << (uint(int32(2)) % 32)
	goto L704
L703:
	;
	v2481 = int32(4)
	goto L704
L704:
	;
	v2483 = *(*int32)(unsafe.Add(mBase, uint32(v2461+v2481)))
	if v2470 == int32(-1) {
		goto L705
	} else {
		goto L706
	}
L705:
	;
	v2486 = v2456
	goto L707
L706:
	;
	v2486 = v16 + int32(112)
	goto L707
L707:
	;
	v2489 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2483))))
	if v2489 != 0 {
		goto L710
	} else {
		goto L711
	}
L708:
	;
	if v2521-v2523 != 0 {
		goto L698
	} else {
		goto L720
	}
L709:
	;
	v2521 = F_tolower(m, v2517)
	mBase = m.M
	v2522 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2518))))
	v2523 = F_tolower(m, v2522)
	mBase = m.M
	goto L708
L710:
	;
	v2491 = v2483
	v2492 = v2486
	v2493 = v2489
	goto L713
L711:
	;
	v2517 = int32(0)
	v2518 = v2486
	goto L709
L712:
	;
	v2517 = v2514 & int32(255)
	v2518 = v2513
	goto L709
L713:
	;
	v2495 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2492))))
	if v2495 == int32(0) {
		v2513 = v2492
		v2514 = v2493
		goto L712
	} else {
		goto L715
	}
L714:
	;
	v2513 = v2507
	v2514 = int32(0)
	goto L712
L715:
	;
	v2499 = v2493 & int32(255)
	if v2499 == v2495 {
		goto L716
	} else {
		goto L717
	}
L716:
	;
	v2506 = int32(1)
	v2507 = v2492 + v2506
	v2508 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2491)+1)))
	if v2508 != 0 {
		v2491 = v2491 + v2506
		v2492 = v2507
		v2493 = v2508
		goto L713
	} else {
		goto L719
	}
L717:
	;
	v2501 = F_tolower(m, v2499)
	mBase = m.M
	v2502 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2492))))
	v2503 = F_tolower(m, v2502)
	mBase = m.M
	if v2501 == v2503 {
		goto L716
	} else {
		goto L718
	}
L718:
	;
	v2505 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2491))))
	v2513 = v2492
	v2514 = v2505
	goto L712
L719:
	;
	goto L714
L720:
	;
	v2525 = *(*int32)(unsafe.Add(mBase, uint32(l0)+200))
	v2526 = *(*int32)(unsafe.Add(mBase, uint32(l0)+192))
	v2527 = *(*int32)(unsafe.Add(mBase, uint32(v2526)+280))
	v2528 = *(*int32)(unsafe.Add(mBase, uint32(v2527)+24))
	v2529 = *(*int32)(unsafe.Add(mBase, uint32(v2528)+8))
	if v2525 != v2529 {
		goto L698
	} else {
		goto L721
	}
L721:
	;
	v2531 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v2531&int32(-769) | int32(512)
	F_sentinelEvent(m, int32(2), int32(_a1444), l0, int32(_a1362), int32(0))
	mBase = m.M
	v2542 = m.ExcPending
	if v2542 != 0 {
		goto L1
	} else {
		goto L722
	}
L722:
	;
	goto L698
L723:
	;
	v2551 = *(*int32)(unsafe.Add(mBase, uint32(l0)+204))
	if v2551 != 0 {
		goto L576
	} else {
		goto L724
	}
L724:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v2546&int32(-1537) | int32(1024)
	F_sentinelEvent(m, int32(2), int32(_a1445), l0, int32(_a1362), int32(0))
	mBase = m.M
	v2562 = m.ExcPending
	if v2562 != 0 {
		goto L1
	} else {
		goto L725
	}
L725:
	;
	goto L576
L726:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_sentinelResetPrimaryAndChangeAddress(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
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
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v52 int32
	_ = v52
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v65 int64
	_ = v65
	var v66 int64
	_ = v66
	var v67 int64
	_ = v67
	var v68 int64
	_ = v68
	var v69 int64
	_ = v69
	var v70 int64
	_ = v70
	var v71 int64
	_ = v71
	var v73 int64
	_ = v73
	var v75 int64
	_ = v75
	var v77 int64
	_ = v77
	var v79 int64
	_ = v79
	var v81 int64
	_ = v81
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v88 int32
	_ = v88
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v103 int32
	_ = v103
	var v107 int32
	_ = v107
	var v109 int32
	_ = v109
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v123 int32
	_ = v123
	var v128 int32
	_ = v128
	var v132 int32
	_ = v132
	var v136 int32
	_ = v136
	var v145 int32
	_ = v145
	var v147 int32
	_ = v147
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v175 int32
	_ = v175
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
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
	var v205 int32
	_ = v205
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
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
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v244 int32
	_ = v244
	var v253 int32
	_ = v253
	var v261 int32
	_ = v261
	var v262 int32
	_ = v262
	var v265 int32
	_ = v265
	var v268 int32
	_ = v268
	var v269 int32
	_ = v269
	var v273 int32
	_ = v273
	var v277 int32
	_ = v277
	var v278 int32
	_ = v278
	var v279 int32
	_ = v279
	var v282 int32
	_ = v282
	var v284 int32
	_ = v284
	var v286 int64
	_ = v286
	var v287 int64
	_ = v287
	var v288 int64
	_ = v288
	var v289 int64
	_ = v289
	var v290 int64
	_ = v290
	var v291 int64
	_ = v291
	var v292 int64
	_ = v292
	var v294 int64
	_ = v294
	var v296 int64
	_ = v296
	var v298 int64
	_ = v298
	var v300 int64
	_ = v300
	var v302 int64
	_ = v302
	var v304 int32
	_ = v304
	var v305 int32
	_ = v305
	var v306 int32
	_ = v306
	var v309 int32
	_ = v309
	var v312 int32
	_ = v312
	var v313 int32
	_ = v313
	var v314 int32
	_ = v314
	var v316 int32
	_ = v316
	var v319 int32
	_ = v319
	var v320 int32
	_ = v320
	var v324 int32
	_ = v324
	var v328 int32
	_ = v328
	var v330 int32
	_ = v330
	var v337 int32
	_ = v337
	var v338 int32
	_ = v338
	var v339 int32
	_ = v339
	var v344 int32
	_ = v344
	var v349 int32
	_ = v349
	var v353 int32
	_ = v353
	var v357 int32
	_ = v357
	var v365 int32
	_ = v365
	var v370 int32
	_ = v370
	var v371 int32
	_ = v371
	var v372 int32
	_ = v372
	var v373 int32
	_ = v373
	var v375 int32
	_ = v375
	var v376 int32
	_ = v376
	var v379 int32
	_ = v379
	var v380 int32
	_ = v380
	var v386 int32
	_ = v386
	var v387 int32
	_ = v387
	var v390 int32
	_ = v390
	var v391 int32
	_ = v391
	var v394 int32
	_ = v394
	var v403 int32
	_ = v403
	var v404 int32
	_ = v404
	var v410 int32
	_ = v410
	var v411 int32
	_ = v411
	var v414 int32
	_ = v414
	var v416 int32
	_ = v416
	var v417 int32
	_ = v417
	var v418 int32
	_ = v418
	var v420 int32
	_ = v420
	var v424 int32
	_ = v424
	var v426 int32
	_ = v426
	var v427 int32
	_ = v427
	var v428 int32
	_ = v428
	var v430 int32
	_ = v430
	var v431 int32
	_ = v431
	var v432 int32
	_ = v432
	var v433 int32
	_ = v433
	var v438 int32
	_ = v438
	var v439 int32
	_ = v439
	var v442 int32
	_ = v442
	var v443 int32
	_ = v443
	var v446 int32
	_ = v446
	var v447 int32
	_ = v447
	var v448 int32
	_ = v448
	var v453 int32
	_ = v453
	var v454 int32
	_ = v454
	var v455 int32
	_ = v455
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
	var v463 int32
	_ = v463
	var v472 int32
	_ = v472
	var v475 int32
	_ = v475
	var v476 int64
	_ = v476
	var v482 int32
	_ = v482
	var v488 int32
	_ = v488
	var v496 int32
	_ = v496
	var v500 int32
	_ = v500
	var v501 int32
	_ = v501
	var v502 int32
	_ = v502
	var v503 int32
	_ = v503
	var v504 int32
	_ = v504
	var v505 int32
	_ = v505
	var v506 int32
	_ = v506
	var v508 int32
	_ = v508
	var v509 int32
	_ = v509
	var v511 int32
	_ = v511
	var v513 int32
	_ = v513
	var v521 int32
	_ = v521
	var v523 int32
	_ = v523
	var v534 int32
	_ = v534
	var v535 int32
	_ = v535
	var v537 int32
	_ = v537
	var v538 int32
	_ = v538
	var v540 int32
	_ = v540
	var v542 int32
	_ = v542
	var v543 int32
	_ = v543
	var v544 int32
	_ = v544
	var v548 int32
	_ = v548
	var v550 int32
	_ = v550
	var v552 int32
	_ = v552
	var v553 int32
	_ = v553
	var v554 int32
	_ = v554
	var v557 int32
	_ = v557
	var v563 int32
	_ = v563
	var v564 int32
	_ = v564
	var v569 int32
	_ = v569
	var v572 int32
	_ = v572
	var v577 int32
	_ = v577
	var v580 int32
	_ = v580
	v4 = int32(0)
	v9 = m.G0
	v11 = v9 - int32(16)
	m.G0 = v11
	v15 = F_createSentinelAddr(m, l1, l2, v4)
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	m.G0 = v11 + int32(16)
	return v580
L2:
	;
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+148))
	v21 = *(*int32)(unsafe.Add(mBase, uint32(v20)+16))
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v20)+12))
	v28 = F_valkey_malloc(m, (v21+v22)<<(uint(int32(2))%32)+int32(4))
	mBase = m.M
	v29 = m.ExcPending
	if v29 != 0 {
		goto L3
	} else {
		goto L6
	}
L3:
	;
	return int32(0)
L4:
	;
	if v15 != 0 {
		goto L2
	} else {
		goto L5
	}
L5:
	;
	v580 = int32(-1)
	goto L1
L6:
	;
	v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)+148))
	v31 = F_dictGetIterator(m, v30)
	mBase = m.M
	v32 = m.ExcPending
	if v32 != 0 {
		goto L3
	} else {
		goto L8
	}
L7:
	;
	F_dictReleaseIterator(m, v31)
	mBase = m.M
	v370 = m.ExcPending
	if v370 != 0 {
		goto L3
	} else {
		goto L94
	}
L8:
	;
	v40 = v31 + int32(20)
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v31)+16))
	if v41 != 0 {
		goto L11
	} else {
		goto L12
	}
L9:
	;
	if v136 == int32(0) {
		v365 = v4
		goto L7
	} else {
		goto L35
	}
L10:
	;
	v47 = v40
	v48 = v44
	goto L13
L11:
	;
	v44 = int32(1)
	goto L10
L12:
	;
	v44 = int32(0)
	goto L10
L13:
	;
	switch v48 {
	case 0:
		goto L18
	default:
		goto L17
	}
L15:
	;
	v48 = int32(0)
	goto L13
L16:
	;
	goto L9
L17:
	;
	v128 = *(*int32)(unsafe.Add(mBase, uint32(v47)))
	*(*int32)(unsafe.Add(mBase, uint32(v31)+16)) = v128
	if v128 == int32(0) {
		goto L15
	} else {
		goto L34
	}
L18:
	;
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v31)+4))
	if v52 != int32(-1) {
		v91 = v52
		goto L19
	} else {
		goto L20
	}
L19:
	;
	v92 = int32(1)
	v93 = v91 + v92
	*(*int32)(unsafe.Add(mBase, uint32(v31)+4)) = v93
	v95 = int32(0)
	v98 = *(*int32)(unsafe.Add(mBase, uint32(v31)))
	v99 = *(*int32)(unsafe.Add(mBase, uint32(v31)+8))
	v103 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v98+v99+int32(26)))))
	if v103 == int32(255) {
		goto L28
	} else {
		goto L29
	}
L20:
	;
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v31)+8))
	if v56 != 0 {
		v91 = int32(-1)
		goto L19
	} else {
		goto L21
	}
L21:
	;
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v31)))
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v31)+12))
	if v58 == int32(0) {
		goto L23
	} else {
		goto L24
	}
L22:
	;
	v85 = *(*int32)(unsafe.Add(mBase, uint32(v84)+20))
	if v85 != int32(-1) {
		goto L25
	} else {
		goto L26
	}
L23:
	;
	v65 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v57)+16)))
	v66 = int64(*(*int8)(unsafe.Add(mBase, uint32(v57)+27)))
	v67 = int64(*(*int32)(unsafe.Add(mBase, uint32(v57)+8)))
	v68 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v57)+12)))
	v69 = int64(*(*int8)(unsafe.Add(mBase, uint32(v57)+26)))
	v70 = int64(*(*int32)(unsafe.Add(mBase, uint32(v57)+4)))
	v71 = F_wangHash64(m, v70)
	mBase = m.M
	v73 = F_wangHash64(m, v69+v71)
	mBase = m.M
	v75 = F_wangHash64(m, v68+v73)
	mBase = m.M
	v77 = F_wangHash64(m, v67+v75)
	mBase = m.M
	v79 = F_wangHash64(m, v66+v77)
	mBase = m.M
	v81 = F_wangHash64(m, v65+v79)
	mBase = m.M
	*(*int64)(unsafe.Add(mBase, uint32(v31)+24)) = v81
	v83 = *(*int32)(unsafe.Add(mBase, uint32(v31)))
	v84 = v83
	goto L22
L24:
	;
	v61 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v57)+24)))
	v63 = v61 + int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v57)+24)) = uint16(v63)
	v84 = v57
	goto L22
L25:
	;
	v91 = v85 + int32(-1)
	goto L19
L26:
	;
	v88 = *(*int32)(unsafe.Add(mBase, uint32(v31)+4))
	v91 = v88
	goto L19
L27:
	;
	v118 = int32(2)
	v123 = *(*int32)(unsafe.Add(mBase, uint32(v98+v116<<(uint(v118)%32)+int32(4))))
	v47 = v123 + v117<<(uint(v118)%32)
	v48 = int32(1)
	goto L13
L28:
	;
	v107 = v95
	goto L30
L29:
	;
	v107 = v92 << (uint(v103) % 32)
	goto L30
L30:
	;
	if v93 < v107 {
		v116 = v99
		v117 = v93
		goto L27
	} else {
		goto L31
	}
L31:
	;
	if v99 != 0 {
		v136 = v95
		goto L16
	} else {
		goto L32
	}
L32:
	;
	v109 = *(*int32)(unsafe.Add(mBase, uint32(v98)+20))
	if v109 == int32(-1) {
		v136 = v95
		goto L16
	} else {
		goto L33
	}
L33:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v31)+4)) = int64(4294967296)
	v116 = int32(1)
	v117 = int32(0)
	goto L27
L34:
	;
	v132 = *(*int32)(unsafe.Add(mBase, uint32(v128)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v40))) = v132
	v136 = v128
	goto L16
L35:
	;
	v145 = v136
	v147 = int32(0)
	goto L36
L36:
	;
	v151 = *(*int32)(unsafe.Add(mBase, uint32(v145)+8))
	goto L40
L37:
	;
	v365 = v253
	goto L7
L38:
	;
	v261 = v31 + int32(20)
	v262 = *(*int32)(unsafe.Add(mBase, uint32(v31)+16))
	if v262 != 0 {
		goto L69
	} else {
		goto L70
	}
L39:
	;
	v234 = F_valkey_malloc(m, int32(12))
	mBase = m.M
	v235 = m.ExcPending
	if v235 != 0 {
		goto L3
	} else {
		goto L64
	}
L40:
	;
	v152 = *(*int32)(unsafe.Add(mBase, uint32(v151)+24))
	v153 = *(*int32)(unsafe.Add(mBase, uint32(v152)+8))
	v154 = *(*int32)(unsafe.Add(mBase, uint32(v15)+8))
	if v153 != v154 {
		goto L39
	} else {
		goto L41
	}
L41:
	;
	v156 = *(*int32)(unsafe.Add(mBase, uint32(v152)+4))
	v157 = *(*int32)(unsafe.Add(mBase, uint32(v15)+4))
	v160 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v157))))
	v161 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v156))))
	if v161 == int32(0) {
		v184 = v160
		v185 = v161
		goto L43
	} else {
		goto L44
	}
L42:
	;
	if v185-v184&int32(255) == int32(0) {
		v253 = v147
		goto L38
	} else {
		goto L50
	}
L43:
	;
	goto L42
L44:
	;
	if v161 != v160&int32(255) {
		v184 = v160
		v185 = v161
		goto L43
	} else {
		goto L45
	}
L45:
	;
	v167 = v156
	v168 = v157
	goto L46
L46:
	;
	v171 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v168)+1)))
	v172 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v167)+1)))
	if v172 == int32(0) {
		v184 = v171
		v185 = v172
		goto L43
	} else {
		goto L48
	}
L47:
	;
	v184 = v171
	v185 = v172
	goto L43
L48:
	;
	v175 = int32(1)
	if v172 == v171&int32(255) {
		v167 = v167 + v175
		v168 = v168 + v175
		goto L46
	} else {
		goto L49
	}
L49:
	;
	goto L47
L50:
	;
	v191 = *(*int32)(unsafe.Add(mBase, uint32(v152)))
	v192 = *(*int32)(unsafe.Add(mBase, uint32(v15)))
	v195 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v191))))
	if v195 != 0 {
		goto L53
	} else {
		goto L54
	}
L51:
	;
	if v227-v229 == int32(0) {
		v253 = v147
		goto L38
	} else {
		goto L63
	}
L52:
	;
	v227 = F_tolower(m, v223)
	mBase = m.M
	v228 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v224))))
	v229 = F_tolower(m, v228)
	mBase = m.M
	goto L51
L53:
	;
	v197 = v191
	v198 = v192
	v199 = v195
	goto L56
L54:
	;
	v223 = int32(0)
	v224 = v192
	goto L52
L55:
	;
	v223 = v220 & int32(255)
	v224 = v219
	goto L52
L56:
	;
	v201 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v198))))
	if v201 == int32(0) {
		v219 = v198
		v220 = v199
		goto L55
	} else {
		goto L58
	}
L57:
	;
	v219 = v213
	v220 = int32(0)
	goto L55
L58:
	;
	v205 = v199 & int32(255)
	if v205 == v201 {
		goto L59
	} else {
		goto L60
	}
L59:
	;
	v212 = int32(1)
	v213 = v198 + v212
	v214 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v197)+1)))
	if v214 != 0 {
		v197 = v197 + v212
		v198 = v213
		v199 = v214
		goto L56
	} else {
		goto L62
	}
L60:
	;
	v207 = F_tolower(m, v205)
	mBase = m.M
	v208 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v198))))
	v209 = F_tolower(m, v208)
	mBase = m.M
	if v207 == v209 {
		goto L59
	} else {
		goto L61
	}
L61:
	;
	v211 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v197))))
	v219 = v198
	v220 = v211
	goto L55
L62:
	;
	goto L57
L63:
	;
	goto L39
L64:
	;
	v236 = *(*int32)(unsafe.Add(mBase, uint32(v152)))
	v237 = F_sdsnew(m, v236)
	mBase = m.M
	v238 = m.ExcPending
	if v238 != 0 {
		goto L3
	} else {
		goto L65
	}
L65:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v234))) = v237
	v240 = *(*int32)(unsafe.Add(mBase, uint32(v152)+4))
	v241 = F_sdsnew(m, v240)
	mBase = m.M
	v242 = m.ExcPending
	if v242 != 0 {
		goto L3
	} else {
		goto L66
	}
L66:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v234)+4)) = v241
	v244 = *(*int32)(unsafe.Add(mBase, uint32(v152)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v234)+8)) = v244
	*(*int32)(unsafe.Add(mBase, uint32(v28+v147<<(uint(int32(2))%32)))) = v234
	v253 = v147 + int32(1)
	goto L38
L67:
	;
	if v357 != 0 {
		v145 = v357
		v147 = v253
		goto L36
	} else {
		goto L93
	}
L68:
	;
	v268 = v261
	v269 = v265
	goto L71
L69:
	;
	v265 = int32(1)
	goto L68
L70:
	;
	v265 = int32(0)
	goto L68
L71:
	;
	switch v269 {
	case 0:
		goto L76
	default:
		goto L75
	}
L73:
	;
	v269 = int32(0)
	goto L71
L74:
	;
	goto L67
L75:
	;
	v349 = *(*int32)(unsafe.Add(mBase, uint32(v268)))
	*(*int32)(unsafe.Add(mBase, uint32(v31)+16)) = v349
	if v349 == int32(0) {
		goto L73
	} else {
		goto L92
	}
L76:
	;
	v273 = *(*int32)(unsafe.Add(mBase, uint32(v31)+4))
	if v273 != int32(-1) {
		v312 = v273
		goto L77
	} else {
		goto L78
	}
L77:
	;
	v313 = int32(1)
	v314 = v312 + v313
	*(*int32)(unsafe.Add(mBase, uint32(v31)+4)) = v314
	v316 = int32(0)
	v319 = *(*int32)(unsafe.Add(mBase, uint32(v31)))
	v320 = *(*int32)(unsafe.Add(mBase, uint32(v31)+8))
	v324 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v319+v320+int32(26)))))
	if v324 == int32(255) {
		goto L86
	} else {
		goto L87
	}
L78:
	;
	v277 = *(*int32)(unsafe.Add(mBase, uint32(v31)+8))
	if v277 != 0 {
		v312 = int32(-1)
		goto L77
	} else {
		goto L79
	}
L79:
	;
	v278 = *(*int32)(unsafe.Add(mBase, uint32(v31)))
	v279 = *(*int32)(unsafe.Add(mBase, uint32(v31)+12))
	if v279 == int32(0) {
		goto L81
	} else {
		goto L82
	}
L80:
	;
	v306 = *(*int32)(unsafe.Add(mBase, uint32(v305)+20))
	if v306 != int32(-1) {
		goto L83
	} else {
		goto L84
	}
L81:
	;
	v286 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v278)+16)))
	v287 = int64(*(*int8)(unsafe.Add(mBase, uint32(v278)+27)))
	v288 = int64(*(*int32)(unsafe.Add(mBase, uint32(v278)+8)))
	v289 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v278)+12)))
	v290 = int64(*(*int8)(unsafe.Add(mBase, uint32(v278)+26)))
	v291 = int64(*(*int32)(unsafe.Add(mBase, uint32(v278)+4)))
	v292 = F_wangHash64(m, v291)
	mBase = m.M
	v294 = F_wangHash64(m, v290+v292)
	mBase = m.M
	v296 = F_wangHash64(m, v289+v294)
	mBase = m.M
	v298 = F_wangHash64(m, v288+v296)
	mBase = m.M
	v300 = F_wangHash64(m, v287+v298)
	mBase = m.M
	v302 = F_wangHash64(m, v286+v300)
	mBase = m.M
	*(*int64)(unsafe.Add(mBase, uint32(v31)+24)) = v302
	v304 = *(*int32)(unsafe.Add(mBase, uint32(v31)))
	v305 = v304
	goto L80
L82:
	;
	v282 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v278)+24)))
	v284 = v282 + int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v278)+24)) = uint16(v284)
	v305 = v278
	goto L80
L83:
	;
	v312 = v306 + int32(-1)
	goto L77
L84:
	;
	v309 = *(*int32)(unsafe.Add(mBase, uint32(v31)+4))
	v312 = v309
	goto L77
L85:
	;
	v339 = int32(2)
	v344 = *(*int32)(unsafe.Add(mBase, uint32(v319+v337<<(uint(v339)%32)+int32(4))))
	v268 = v344 + v338<<(uint(v339)%32)
	v269 = int32(1)
	goto L71
L86:
	;
	v328 = v316
	goto L88
L87:
	;
	v328 = v313 << (uint(v324) % 32)
	goto L88
L88:
	;
	if v314 < v328 {
		v337 = v320
		v338 = v314
		goto L85
	} else {
		goto L89
	}
L89:
	;
	if v320 != 0 {
		v357 = v316
		goto L74
	} else {
		goto L90
	}
L90:
	;
	v330 = *(*int32)(unsafe.Add(mBase, uint32(v319)+20))
	if v330 == int32(-1) {
		v357 = v316
		goto L74
	} else {
		goto L91
	}
L91:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v31)+4)) = int64(4294967296)
	v337 = int32(1)
	v338 = int32(0)
	goto L85
L92:
	;
	v353 = *(*int32)(unsafe.Add(mBase, uint32(v349)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v261))) = v353
	v357 = v349
	goto L74
L93:
	;
	goto L37
L94:
	;
	v371 = *(*int32)(unsafe.Add(mBase, uint32(v15)+8))
	v372 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v373 = *(*int32)(unsafe.Add(mBase, uint32(v372)+8))
	if v371 != v373 {
		goto L96
	} else {
		goto L97
	}
L95:
	;
	F_sentinelResetPrimary(m, l0, int32(1))
	mBase = m.M
	v475 = m.ExcPending
	if v475 != 0 {
		goto L3
	} else {
		goto L123
	}
L96:
	;
	v453 = F_valkey_malloc(m, int32(12))
	mBase = m.M
	v454 = m.ExcPending
	if v454 != 0 {
		goto L3
	} else {
		goto L120
	}
L97:
	;
	v375 = *(*int32)(unsafe.Add(mBase, uint32(v15)+4))
	v376 = *(*int32)(unsafe.Add(mBase, uint32(v372)+4))
	v379 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v376))))
	v380 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v375))))
	if v380 == int32(0) {
		v403 = v379
		v404 = v380
		goto L99
	} else {
		goto L100
	}
L98:
	;
	if v404-v403&int32(255) == int32(0) {
		v472 = v365
		goto L95
	} else {
		goto L106
	}
L99:
	;
	goto L98
L100:
	;
	if v380 != v379&int32(255) {
		v403 = v379
		v404 = v380
		goto L99
	} else {
		goto L101
	}
L101:
	;
	v386 = v375
	v387 = v376
	goto L102
L102:
	;
	v390 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v387)+1)))
	v391 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v386)+1)))
	if v391 == int32(0) {
		v403 = v390
		v404 = v391
		goto L99
	} else {
		goto L104
	}
L103:
	;
	v403 = v390
	v404 = v391
	goto L99
L104:
	;
	v394 = int32(1)
	if v391 == v390&int32(255) {
		v386 = v386 + v394
		v387 = v387 + v394
		goto L102
	} else {
		goto L105
	}
L105:
	;
	goto L103
L106:
	;
	v410 = *(*int32)(unsafe.Add(mBase, uint32(v15)))
	v411 = *(*int32)(unsafe.Add(mBase, uint32(v372)))
	v414 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v410))))
	if v414 != 0 {
		goto L109
	} else {
		goto L110
	}
L107:
	;
	if v446-v448 == int32(0) {
		v472 = v365
		goto L95
	} else {
		goto L119
	}
L108:
	;
	v446 = F_tolower(m, v442)
	mBase = m.M
	v447 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v443))))
	v448 = F_tolower(m, v447)
	mBase = m.M
	goto L107
L109:
	;
	v416 = v410
	v417 = v411
	v418 = v414
	goto L112
L110:
	;
	v442 = int32(0)
	v443 = v411
	goto L108
L111:
	;
	v442 = v439 & int32(255)
	v443 = v438
	goto L108
L112:
	;
	v420 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v417))))
	if v420 == int32(0) {
		v438 = v417
		v439 = v418
		goto L111
	} else {
		goto L114
	}
L113:
	;
	v438 = v432
	v439 = int32(0)
	goto L111
L114:
	;
	v424 = v418 & int32(255)
	if v424 == v420 {
		goto L115
	} else {
		goto L116
	}
L115:
	;
	v431 = int32(1)
	v432 = v417 + v431
	v433 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v416)+1)))
	if v433 != 0 {
		v416 = v416 + v431
		v417 = v432
		v418 = v433
		goto L112
	} else {
		goto L118
	}
L116:
	;
	v426 = F_tolower(m, v424)
	mBase = m.M
	v427 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v417))))
	v428 = F_tolower(m, v427)
	mBase = m.M
	if v426 == v428 {
		goto L115
	} else {
		goto L117
	}
L117:
	;
	v430 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v416))))
	v438 = v417
	v439 = v430
	goto L111
L118:
	;
	goto L113
L119:
	;
	goto L96
L120:
	;
	v455 = *(*int32)(unsafe.Add(mBase, uint32(v372)))
	v456 = F_sdsnew(m, v455)
	mBase = m.M
	v457 = m.ExcPending
	if v457 != 0 {
		goto L3
	} else {
		goto L121
	}
L121:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v453))) = v456
	v459 = *(*int32)(unsafe.Add(mBase, uint32(v372)+4))
	v460 = F_sdsnew(m, v459)
	mBase = m.M
	v461 = m.ExcPending
	if v461 != 0 {
		goto L3
	} else {
		goto L122
	}
L122:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v453)+4)) = v460
	v463 = *(*int32)(unsafe.Add(mBase, uint32(v372)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v453)+8)) = v463
	*(*int32)(unsafe.Add(mBase, uint32(v28+v365<<(uint(int32(2))%32)))) = v453
	v472 = v365 + int32(1)
	goto L95
L123:
	;
	v476 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(l0)+56)) = v476
	*(*int64)(unsafe.Add(mBase, uint32(l0+int32(64)))) = v476
	v482 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = v15
	if v472 < int32(1) {
		goto L124
	} else {
		goto L125
	}
L124:
	;
	F_valkey_free(m, v28)
	mBase = m.M
	v534 = m.ExcPending
	if v534 != 0 {
		goto L3
	} else {
		goto L136
	}
L125:
	;
	v488 = int32(0)
	goto L126
L126:
	;
	v496 = int32(2)
	v500 = *(*int32)(unsafe.Add(mBase, uint32(v28+v488<<(uint(v496)%32))))
	v501 = *(*int32)(unsafe.Add(mBase, uint32(v500)))
	v502 = *(*int32)(unsafe.Add(mBase, uint32(v500)+8))
	v503 = *(*int32)(unsafe.Add(mBase, uint32(l0)+152))
	v504 = F_createSentinelValkeyInstance(m, int32(0), v496, v501, v502, v503, l0)
	mBase = m.M
	v505 = m.ExcPending
	if v505 != 0 {
		goto L3
	} else {
		goto L128
	}
L127:
	;
	goto L124
L128:
	;
	v506 = *(*int32)(unsafe.Add(mBase, uint32(v500)))
	F_sdsfree(m, v506)
	mBase = m.M
	v508 = m.ExcPending
	if v508 != 0 {
		goto L3
	} else {
		goto L129
	}
L129:
	;
	v509 = *(*int32)(unsafe.Add(mBase, uint32(v500)+4))
	F_sdsfree(m, v509)
	mBase = m.M
	v511 = m.ExcPending
	if v511 != 0 {
		goto L3
	} else {
		goto L130
	}
L130:
	;
	F_valkey_free(m, v500)
	mBase = m.M
	v513 = m.ExcPending
	if v513 != 0 {
		goto L3
	} else {
		goto L131
	}
L131:
	;
	if v504 == int32(0) {
		goto L132
	} else {
		goto L133
	}
L132:
	;
	v523 = v488 + int32(1)
	if v523 != v472 {
		v488 = v523
		goto L126
	} else {
		goto L135
	}
L133:
	;
	F_sentinelEvent(m, int32(2), int32(_a1361), v504, int32(_a1362), int32(0))
	mBase = m.M
	v521 = m.ExcPending
	if v521 != 0 {
		goto L3
	} else {
		goto L134
	}
L134:
	;
	goto L132
L135:
	;
	goto L127
L136:
	;
	v535 = *(*int32)(unsafe.Add(mBase, uint32(v482)))
	F_sdsfree(m, v535)
	mBase = m.M
	v537 = m.ExcPending
	if v537 != 0 {
		goto L3
	} else {
		goto L137
	}
L137:
	;
	v538 = *(*int32)(unsafe.Add(mBase, uint32(v482)+4))
	F_sdsfree(m, v538)
	mBase = m.M
	v540 = m.ExcPending
	if v540 != 0 {
		goto L3
	} else {
		goto L138
	}
L138:
	;
	F_valkey_free(m, v482)
	mBase = m.M
	v542 = m.ExcPending
	if v542 != 0 {
		goto L3
	} else {
		goto L139
	}
L139:
	;
	v543 = int32(_a44)
	v544 = *(*int32)(unsafe.Add(mBase, _consts[219]))
	*(*int32)(unsafe.Add(mBase, _consts[219])) = int32(10)
	v548 = int32(0)
	v550 = *(*int32)(unsafe.Add(mBase, _consts[392]))
	v552 = F_rewriteConfig(m, v550, v548)
	mBase = m.M
	v553 = m.ExcPending
	if v553 != 0 {
		goto L3
	} else {
		goto L140
	}
L140:
	;
	v554 = int32(_a44)
	*(*int32)(unsafe.Add(mBase, _consts[219])) = v544
	v557 = *(*int32)(unsafe.Add(mBase, _consts[6]))
	if v552 != int32(-1) {
		goto L141
	} else {
		goto L142
	}
L141:
	;
	if int32(2) < v557 {
		v580 = v548
		goto L1
	} else {
		goto L147
	}
L142:
	;
	if int32(3) < v557 {
		v580 = v548
		goto L1
	} else {
		goto L143
	}
L143:
	;
	goto L144
L144:
	;
	v563 = *(*int32)(unsafe.Add(mBase, _consts[5]))
	v564 = F___strerror_l(m, v563, v563)
	mBase = m.M
	goto L145
L145:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11))) = v564
	F__serverLog(m, int32(3), int32(_a1337), v11)
	mBase = m.M
	v569 = m.ExcPending
	if v569 != 0 {
		goto L3
	} else {
		goto L146
	}
L146:
	;
	v580 = v548
	goto L1
L147:
	;
	v572 = int32(0)
	F__serverLog(m, int32(2), int32(_a1338), v572)
	mBase = m.M
	v577 = m.ExcPending
	if v577 != 0 {
		goto L3
	} else {
		goto L148
	}
L148:
	;
	v580 = v572
	goto L1
}
func F_sentinelRunPendingScripts(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v1 int32
	_ = v1
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int64
	_ = v10
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v20 int32
	_ = v20
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v48 int64
	_ = v48
	var v52 int32
	_ = v52
	var v55 int64
	_ = v55
	var v57 int32
	_ = v57
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v87 int32
	_ = v87
	var v92 int32
	_ = v92
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v111 int32
	_ = v111
	var v115 int32
	_ = v115
	v1 = int32(0)
	v6 = m.G0
	v8 = v6 - int32(48)
	m.G0 = v8
	v10 = F_mstime(m)
	mBase = m.M
	v12 = *(*int32)(unsafe.Add(mBase, _consts[730]))
	v14 = v8 + int32(40)
	v15 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
	*(*int32)(unsafe.Add(mBase, uint32(v14)+4)) = v1
	*(*int32)(unsafe.Add(mBase, uint32(v14))) = v15
	goto L1
L1:
	;
	v20 = *(*int32)(unsafe.Add(mBase, _consts[731]))
	if int32(15) < v20 {
		goto L2
	} else {
		goto L3
	}
L2:
	;
	m.G0 = v8 + int32(48)
	return
L3:
	;
	goto L4
L4:
	;
	v29 = v8 + int32(40)
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v29)))
	if v31 == int32(0) {
		goto L7
	} else {
		goto L8
	}
L5:
	;
	goto L2
L6:
	;
	if v31 == int32(0) {
		goto L2
	} else {
		goto L9
	}
L7:
	;
	goto L6
L8:
	;
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v29)+4))
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v31+base.B2i32(v34 == int32(0))<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v29))) = v40
	goto L7
L9:
	;
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v31)+8))
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v44)))
	if v45&int32(1) != 0 {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	v115 = *(*int32)(unsafe.Add(mBase, _consts[731]))
	if v115 < int32(16) {
		goto L4
	} else {
		goto L25
	}
L11:
	;
	v48 = *(*int64)(unsafe.Add(mBase, uint32(v44)+16))
	if v48 == int64(0) {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	v52 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v44))) = v45 | v52
	v55 = F_mstime(m)
	mBase = m.M
	*(*int64)(unsafe.Add(mBase, uint32(v44)+16)) = v55
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v44)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v44)+4)) = v57 + v52
	v61 = F___errno_location(m)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v61))) = int32(52)
	v64 = int32(-1)
	goto L18
L13:
	;
	if v10 < v48 {
		goto L10
	} else {
		goto L14
	}
L14:
	;
	goto L12
L15:
	;
	v98 = int32(0)
	v100 = *(*int32)(unsafe.Add(mBase, _consts[731]))
	*(*int32)(unsafe.Add(mBase, _consts[731])) = v100 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v44)+24)) = v64
	*(*int32)(unsafe.Add(mBase, uint32(v8))) = v64
	F_sentinelEvent(m, v98, int32(_a1341), v98, int32(_a1342), v8)
	mBase = m.M
	v111 = m.ExcPending
	if v111 != 0 {
		goto L19
	} else {
		goto L24
	}
L16:
	;
	F_connTypeCleanupAll(m)
	mBase = m.M
	v87 = m.ExcPending
	if v87 != 0 {
		goto L19
	} else {
		goto L21
	}
L17:
	;
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v44)+8))
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v67)))
	*(*int64)(unsafe.Add(mBase, uint32(v8)+20)) = int64(99)
	*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = v68
	F_sentinelEvent(m, int32(3), int32(_a1343), int32(0), int32(_a1344), v8+int32(16))
	mBase = m.M
	v79 = m.ExcPending
	if v79 != 0 {
		goto L19
	} else {
		goto L20
	}
L18:
	;
	switch int32(0) {
	case 0:
		goto L17
	case 1:
		goto L16
	default:
		goto L15
	}
L19:
	;
	return
L20:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+24)) = int32(0)
	v82 = *(*int32)(unsafe.Add(mBase, uint32(v44)))
	*(*int32)(unsafe.Add(mBase, uint32(v44))) = v82 & int32(-2)
	goto L10
L21:
	;
	v92 = F___errno_location(m)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v92))) = int32(45)
	goto L22
L22:
	;
	F__Exit(m, int32(2))
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L24:
	;
	goto L10
L25:
	;
	goto L5
}
func F_sentinelSendAuthIfNeeded(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
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
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
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
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v80 int32
	_ = v80
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	v8 = m.G0
	v10 = v8 - int32(32)
	m.G0 = v10
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v12&int32(1) == int32(0) {
		if v12&int32(2) == int32(0) {
			if v12&int32(4) == int32(0) {
				m.G0 = v10 + int32(32)
				return
			} else {
				v30 = int32(0)
				v32 = *(*int32)(unsafe.Add(mBase, _consts[752]))
				if v32 == v30 {
					v38 = *(*int32)(unsafe.Add(mBase, _consts[753]))
					v39 = v38
					v40 = v30
				} else {
					v36 = *(*int32)(unsafe.Add(mBase, _consts[751]))
					v39 = v32
					v40 = v36
				}
				if v39 == int32(0) {
					if v39 == int32(0) {
						m.G0 = v10 + int32(32)
						return
					} else {
						if v40 == int32(0) {
							m.G0 = v10 + int32(32)
							return
						} else {
							v68 = F_sdsnew(m, int32(_a1035))
							mBase = m.M
							v69 = m.ExcPending
							if v69 != 0 {
								return
							} else {
								v70 = *(*int32)(unsafe.Add(mBase, uint32(l0)+192))
								if v70 != 0 {
									v71 = v70
								} else {
									v71 = l0
								}
								v72 = *(*int32)(unsafe.Add(mBase, uint32(v71)+104))
								v73 = F_dictFetchValue(m, v72, v68)
								mBase = m.M
								v74 = m.ExcPending
								if v74 != 0 {
									return
								} else {
									F_sdsfree(m, v68)
									mBase = m.M
									v76 = m.ExcPending
									if v76 != 0 {
										return
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v10)+24)) = v39
										*(*int32)(unsafe.Add(mBase, uint32(v10)+20)) = v40
										if v73 != 0 {
											v80 = v73
										} else {
											v80 = int32(_a1035)
										}
										*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = v80
										v86 = F_valkeyAsyncCommand(m, l1, int32(1004), l0, int32(_a1411), v10+int32(16))
										mBase = m.M
										v87 = m.ExcPending
										if v87 != 0 {
											return
										} else {
											if v86 != 0 {
											} else {
												v91 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
												v92 = *(*int32)(unsafe.Add(mBase, uint32(v91)+8))
												*(*int32)(unsafe.Add(mBase, uint32(v91)+8)) = v92 + int32(1)
											}
											m.G0 = v10 + int32(32)
											return
										}
									}
								}
							}
						}
					}
				} else {
					if v40 != 0 {
						if v39 == int32(0) {
							m.G0 = v10 + int32(32)
							return
						} else {
							if v40 == int32(0) {
								m.G0 = v10 + int32(32)
								return
							} else {
								v68 = F_sdsnew(m, int32(_a1035))
								mBase = m.M
								v69 = m.ExcPending
								if v69 != 0 {
									return
								} else {
									v70 = *(*int32)(unsafe.Add(mBase, uint32(l0)+192))
									if v70 != 0 {
										v71 = v70
									} else {
										v71 = l0
									}
									v72 = *(*int32)(unsafe.Add(mBase, uint32(v71)+104))
									v73 = F_dictFetchValue(m, v72, v68)
									mBase = m.M
									v74 = m.ExcPending
									if v74 != 0 {
										return
									} else {
										F_sdsfree(m, v68)
										mBase = m.M
										v76 = m.ExcPending
										if v76 != 0 {
											return
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(v10)+24)) = v39
											*(*int32)(unsafe.Add(mBase, uint32(v10)+20)) = v40
											if v73 != 0 {
												v80 = v73
											} else {
												v80 = int32(_a1035)
											}
											*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = v80
											v86 = F_valkeyAsyncCommand(m, l1, int32(1004), l0, int32(_a1411), v10+int32(16))
											mBase = m.M
											v87 = m.ExcPending
											if v87 != 0 {
												return
											} else {
												if v86 != 0 {
												} else {
													v91 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
													v92 = *(*int32)(unsafe.Add(mBase, uint32(v91)+8))
													*(*int32)(unsafe.Add(mBase, uint32(v91)+8)) = v92 + int32(1)
												}
												m.G0 = v10 + int32(32)
												return
											}
										}
									}
								}
							}
						}
					} else {
						v44 = F_sdsnew(m, int32(_a1035))
						mBase = m.M
						v45 = m.ExcPending
						if v45 != 0 {
							return
						} else {
							v46 = *(*int32)(unsafe.Add(mBase, uint32(l0)+192))
							if v46 != 0 {
								v47 = v46
							} else {
								v47 = l0
							}
							v48 = *(*int32)(unsafe.Add(mBase, uint32(v47)+104))
							v49 = F_dictFetchValue(m, v48, v44)
							mBase = m.M
							v50 = m.ExcPending
							if v50 != 0 {
								return
							} else {
								F_sdsfree(m, v44)
								mBase = m.M
								v52 = m.ExcPending
								if v52 != 0 {
									return
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v10)+4)) = v39
									if v49 != 0 {
										v55 = v49
									} else {
										v55 = int32(_a1035)
									}
									*(*int32)(unsafe.Add(mBase, uint32(v10))) = v55
									v59 = F_valkeyAsyncCommand(m, l1, int32(1004), l0, int32(_a1412), v10)
									mBase = m.M
									v60 = m.ExcPending
									if v60 != 0 {
										return
									} else {
										if v59 == int32(0) {
											v91 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
											v92 = *(*int32)(unsafe.Add(mBase, uint32(v91)+8))
											*(*int32)(unsafe.Add(mBase, uint32(v91)+8)) = v92 + int32(1)
										} else {
										}
										m.G0 = v10 + int32(32)
										return
									}
								}
							}
						}
					}
				}
			}
		} else {
			v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+192))
			v24 = *(*int32)(unsafe.Add(mBase, uint32(v23)+164))
			v25 = *(*int32)(unsafe.Add(mBase, uint32(v23)+160))
			v39 = v25
			v40 = v24
			if v39 == int32(0) {
				if v39 == int32(0) {
					m.G0 = v10 + int32(32)
					return
				} else {
					if v40 == int32(0) {
						m.G0 = v10 + int32(32)
						return
					} else {
						v68 = F_sdsnew(m, int32(_a1035))
						mBase = m.M
						v69 = m.ExcPending
						if v69 != 0 {
							return
						} else {
							v70 = *(*int32)(unsafe.Add(mBase, uint32(l0)+192))
							if v70 != 0 {
								v71 = v70
							} else {
								v71 = l0
							}
							v72 = *(*int32)(unsafe.Add(mBase, uint32(v71)+104))
							v73 = F_dictFetchValue(m, v72, v68)
							mBase = m.M
							v74 = m.ExcPending
							if v74 != 0 {
								return
							} else {
								F_sdsfree(m, v68)
								mBase = m.M
								v76 = m.ExcPending
								if v76 != 0 {
									return
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v10)+24)) = v39
									*(*int32)(unsafe.Add(mBase, uint32(v10)+20)) = v40
									if v73 != 0 {
										v80 = v73
									} else {
										v80 = int32(_a1035)
									}
									*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = v80
									v86 = F_valkeyAsyncCommand(m, l1, int32(1004), l0, int32(_a1411), v10+int32(16))
									mBase = m.M
									v87 = m.ExcPending
									if v87 != 0 {
										return
									} else {
										if v86 != 0 {
										} else {
											v91 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
											v92 = *(*int32)(unsafe.Add(mBase, uint32(v91)+8))
											*(*int32)(unsafe.Add(mBase, uint32(v91)+8)) = v92 + int32(1)
										}
										m.G0 = v10 + int32(32)
										return
									}
								}
							}
						}
					}
				}
			} else {
				if v40 != 0 {
					if v39 == int32(0) {
						m.G0 = v10 + int32(32)
						return
					} else {
						if v40 == int32(0) {
							m.G0 = v10 + int32(32)
							return
						} else {
							v68 = F_sdsnew(m, int32(_a1035))
							mBase = m.M
							v69 = m.ExcPending
							if v69 != 0 {
								return
							} else {
								v70 = *(*int32)(unsafe.Add(mBase, uint32(l0)+192))
								if v70 != 0 {
									v71 = v70
								} else {
									v71 = l0
								}
								v72 = *(*int32)(unsafe.Add(mBase, uint32(v71)+104))
								v73 = F_dictFetchValue(m, v72, v68)
								mBase = m.M
								v74 = m.ExcPending
								if v74 != 0 {
									return
								} else {
									F_sdsfree(m, v68)
									mBase = m.M
									v76 = m.ExcPending
									if v76 != 0 {
										return
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v10)+24)) = v39
										*(*int32)(unsafe.Add(mBase, uint32(v10)+20)) = v40
										if v73 != 0 {
											v80 = v73
										} else {
											v80 = int32(_a1035)
										}
										*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = v80
										v86 = F_valkeyAsyncCommand(m, l1, int32(1004), l0, int32(_a1411), v10+int32(16))
										mBase = m.M
										v87 = m.ExcPending
										if v87 != 0 {
											return
										} else {
											if v86 != 0 {
											} else {
												v91 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
												v92 = *(*int32)(unsafe.Add(mBase, uint32(v91)+8))
												*(*int32)(unsafe.Add(mBase, uint32(v91)+8)) = v92 + int32(1)
											}
											m.G0 = v10 + int32(32)
											return
										}
									}
								}
							}
						}
					}
				} else {
					v44 = F_sdsnew(m, int32(_a1035))
					mBase = m.M
					v45 = m.ExcPending
					if v45 != 0 {
						return
					} else {
						v46 = *(*int32)(unsafe.Add(mBase, uint32(l0)+192))
						if v46 != 0 {
							v47 = v46
						} else {
							v47 = l0
						}
						v48 = *(*int32)(unsafe.Add(mBase, uint32(v47)+104))
						v49 = F_dictFetchValue(m, v48, v44)
						mBase = m.M
						v50 = m.ExcPending
						if v50 != 0 {
							return
						} else {
							F_sdsfree(m, v44)
							mBase = m.M
							v52 = m.ExcPending
							if v52 != 0 {
								return
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v10)+4)) = v39
								if v49 != 0 {
									v55 = v49
								} else {
									v55 = int32(_a1035)
								}
								*(*int32)(unsafe.Add(mBase, uint32(v10))) = v55
								v59 = F_valkeyAsyncCommand(m, l1, int32(1004), l0, int32(_a1412), v10)
								mBase = m.M
								v60 = m.ExcPending
								if v60 != 0 {
									return
								} else {
									if v59 == int32(0) {
										v91 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
										v92 = *(*int32)(unsafe.Add(mBase, uint32(v91)+8))
										*(*int32)(unsafe.Add(mBase, uint32(v91)+8)) = v92 + int32(1)
									} else {
									}
									m.G0 = v10 + int32(32)
									return
								}
							}
						}
					}
				}
			}
		}
	} else {
		v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+164))
		v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+160))
		v39 = v18
		v40 = v17
		if v39 == int32(0) {
			if v39 == int32(0) {
				m.G0 = v10 + int32(32)
				return
			} else {
				if v40 == int32(0) {
					m.G0 = v10 + int32(32)
					return
				} else {
					v68 = F_sdsnew(m, int32(_a1035))
					mBase = m.M
					v69 = m.ExcPending
					if v69 != 0 {
						return
					} else {
						v70 = *(*int32)(unsafe.Add(mBase, uint32(l0)+192))
						if v70 != 0 {
							v71 = v70
						} else {
							v71 = l0
						}
						v72 = *(*int32)(unsafe.Add(mBase, uint32(v71)+104))
						v73 = F_dictFetchValue(m, v72, v68)
						mBase = m.M
						v74 = m.ExcPending
						if v74 != 0 {
							return
						} else {
							F_sdsfree(m, v68)
							mBase = m.M
							v76 = m.ExcPending
							if v76 != 0 {
								return
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v10)+24)) = v39
								*(*int32)(unsafe.Add(mBase, uint32(v10)+20)) = v40
								if v73 != 0 {
									v80 = v73
								} else {
									v80 = int32(_a1035)
								}
								*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = v80
								v86 = F_valkeyAsyncCommand(m, l1, int32(1004), l0, int32(_a1411), v10+int32(16))
								mBase = m.M
								v87 = m.ExcPending
								if v87 != 0 {
									return
								} else {
									if v86 != 0 {
									} else {
										v91 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
										v92 = *(*int32)(unsafe.Add(mBase, uint32(v91)+8))
										*(*int32)(unsafe.Add(mBase, uint32(v91)+8)) = v92 + int32(1)
									}
									m.G0 = v10 + int32(32)
									return
								}
							}
						}
					}
				}
			}
		} else {
			if v40 != 0 {
				if v39 == int32(0) {
					m.G0 = v10 + int32(32)
					return
				} else {
					if v40 == int32(0) {
						m.G0 = v10 + int32(32)
						return
					} else {
						v68 = F_sdsnew(m, int32(_a1035))
						mBase = m.M
						v69 = m.ExcPending
						if v69 != 0 {
							return
						} else {
							v70 = *(*int32)(unsafe.Add(mBase, uint32(l0)+192))
							if v70 != 0 {
								v71 = v70
							} else {
								v71 = l0
							}
							v72 = *(*int32)(unsafe.Add(mBase, uint32(v71)+104))
							v73 = F_dictFetchValue(m, v72, v68)
							mBase = m.M
							v74 = m.ExcPending
							if v74 != 0 {
								return
							} else {
								F_sdsfree(m, v68)
								mBase = m.M
								v76 = m.ExcPending
								if v76 != 0 {
									return
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v10)+24)) = v39
									*(*int32)(unsafe.Add(mBase, uint32(v10)+20)) = v40
									if v73 != 0 {
										v80 = v73
									} else {
										v80 = int32(_a1035)
									}
									*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = v80
									v86 = F_valkeyAsyncCommand(m, l1, int32(1004), l0, int32(_a1411), v10+int32(16))
									mBase = m.M
									v87 = m.ExcPending
									if v87 != 0 {
										return
									} else {
										if v86 != 0 {
										} else {
											v91 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
											v92 = *(*int32)(unsafe.Add(mBase, uint32(v91)+8))
											*(*int32)(unsafe.Add(mBase, uint32(v91)+8)) = v92 + int32(1)
										}
										m.G0 = v10 + int32(32)
										return
									}
								}
							}
						}
					}
				}
			} else {
				v44 = F_sdsnew(m, int32(_a1035))
				mBase = m.M
				v45 = m.ExcPending
				if v45 != 0 {
					return
				} else {
					v46 = *(*int32)(unsafe.Add(mBase, uint32(l0)+192))
					if v46 != 0 {
						v47 = v46
					} else {
						v47 = l0
					}
					v48 = *(*int32)(unsafe.Add(mBase, uint32(v47)+104))
					v49 = F_dictFetchValue(m, v48, v44)
					mBase = m.M
					v50 = m.ExcPending
					if v50 != 0 {
						return
					} else {
						F_sdsfree(m, v44)
						mBase = m.M
						v52 = m.ExcPending
						if v52 != 0 {
							return
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v10)+4)) = v39
							if v49 != 0 {
								v55 = v49
							} else {
								v55 = int32(_a1035)
							}
							*(*int32)(unsafe.Add(mBase, uint32(v10))) = v55
							v59 = F_valkeyAsyncCommand(m, l1, int32(1004), l0, int32(_a1412), v10)
							mBase = m.M
							v60 = m.ExcPending
							if v60 != 0 {
								return
							} else {
								if v59 == int32(0) {
									v91 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
									v92 = *(*int32)(unsafe.Add(mBase, uint32(v91)+8))
									*(*int32)(unsafe.Add(mBase, uint32(v91)+8)) = v92 + int32(1)
								} else {
								}
								m.G0 = v10 + int32(32)
								return
							}
						}
					}
				}
			}
		}
	}
}
func F_sentinelSetCommand(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
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
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
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
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v102 int32
	_ = v102
	var v106 int32
	_ = v106
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v113 int64
	_ = v113
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v122 int32
	_ = v122
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v128 int32
	_ = v128
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
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
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
	var v164 int32
	_ = v164
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v171 int64
	_ = v171
	var v175 int32
	_ = v175
	var v178 int32
	_ = v178
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v184 int32
	_ = v184
	var v188 int32
	_ = v188
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v214 int32
	_ = v214
	var v216 int32
	_ = v216
	var v220 int32
	_ = v220
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v227 int64
	_ = v227
	var v231 int32
	_ = v231
	var v234 int32
	_ = v234
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v240 int32
	_ = v240
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
	var v251 int32
	_ = v251
	var v252 int32
	_ = v252
	var v253 int32
	_ = v253
	var v258 int32
	_ = v258
	var v259 int32
	_ = v259
	var v262 int32
	_ = v262
	var v263 int32
	_ = v263
	var v266 int32
	_ = v266
	var v267 int32
	_ = v267
	var v268 int32
	_ = v268
	var v270 int32
	_ = v270
	var v272 int32
	_ = v272
	var v276 int32
	_ = v276
	var v277 int32
	_ = v277
	var v278 int32
	_ = v278
	var v279 int32
	_ = v279
	var v284 int32
	_ = v284
	var v286 int32
	_ = v286
	var v287 int32
	_ = v287
	var v294 int32
	_ = v294
	var v297 int32
	_ = v297
	var v300 int32
	_ = v300
	var v303 int32
	_ = v303
	var v304 int32
	_ = v304
	var v308 int32
	_ = v308
	var v313 int32
	_ = v313
	var v315 int32
	_ = v315
	var v317 int32
	_ = v317
	var v318 int32
	_ = v318
	var v319 int32
	_ = v319
	var v326 int32
	_ = v326
	var v329 int32
	_ = v329
	var v332 int32
	_ = v332
	var v335 int32
	_ = v335
	var v336 int32
	_ = v336
	var v339 int32
	_ = v339
	var v340 int32
	_ = v340
	var v341 int32
	_ = v341
	var v344 int32
	_ = v344
	var v347 int32
	_ = v347
	var v349 int32
	_ = v349
	var v350 int32
	_ = v350
	var v351 int32
	_ = v351
	var v353 int32
	_ = v353
	var v357 int32
	_ = v357
	var v359 int32
	_ = v359
	var v360 int32
	_ = v360
	var v361 int32
	_ = v361
	var v363 int32
	_ = v363
	var v364 int32
	_ = v364
	var v365 int32
	_ = v365
	var v366 int32
	_ = v366
	var v371 int32
	_ = v371
	var v372 int32
	_ = v372
	var v375 int32
	_ = v375
	var v376 int32
	_ = v376
	var v379 int32
	_ = v379
	var v380 int32
	_ = v380
	var v381 int32
	_ = v381
	var v383 int32
	_ = v383
	var v385 int32
	_ = v385
	var v389 int32
	_ = v389
	var v390 int32
	_ = v390
	var v391 int32
	_ = v391
	var v392 int32
	_ = v392
	var v397 int32
	_ = v397
	var v399 int32
	_ = v399
	var v400 int32
	_ = v400
	var v407 int32
	_ = v407
	var v410 int32
	_ = v410
	var v413 int32
	_ = v413
	var v416 int32
	_ = v416
	var v417 int32
	_ = v417
	var v421 int32
	_ = v421
	var v426 int32
	_ = v426
	var v428 int32
	_ = v428
	var v430 int32
	_ = v430
	var v431 int32
	_ = v431
	var v432 int32
	_ = v432
	var v439 int32
	_ = v439
	var v442 int32
	_ = v442
	var v445 int32
	_ = v445
	var v448 int32
	_ = v448
	var v449 int32
	_ = v449
	var v452 int32
	_ = v452
	var v453 int32
	_ = v453
	var v454 int32
	_ = v454
	var v457 int32
	_ = v457
	var v460 int32
	_ = v460
	var v462 int32
	_ = v462
	var v463 int32
	_ = v463
	var v464 int32
	_ = v464
	var v466 int32
	_ = v466
	var v470 int32
	_ = v470
	var v472 int32
	_ = v472
	var v473 int32
	_ = v473
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
	var v484 int32
	_ = v484
	var v485 int32
	_ = v485
	var v488 int32
	_ = v488
	var v489 int32
	_ = v489
	var v492 int32
	_ = v492
	var v493 int32
	_ = v493
	var v494 int32
	_ = v494
	var v496 int32
	_ = v496
	var v498 int32
	_ = v498
	var v502 int32
	_ = v502
	var v503 int32
	_ = v503
	var v504 int32
	_ = v504
	var v506 int32
	_ = v506
	var v507 int32
	_ = v507
	var v511 int32
	_ = v511
	var v518 int32
	_ = v518
	var v521 int32
	_ = v521
	var v524 int32
	_ = v524
	var v527 int32
	_ = v527
	var v528 int32
	_ = v528
	var v529 int32
	_ = v529
	var v532 int32
	_ = v532
	var v533 int32
	_ = v533
	var v534 int32
	_ = v534
	var v538 int32
	_ = v538
	var v539 int32
	_ = v539
	var v542 int32
	_ = v542
	var v544 int32
	_ = v544
	var v545 int32
	_ = v545
	var v546 int32
	_ = v546
	var v548 int32
	_ = v548
	var v552 int32
	_ = v552
	var v554 int32
	_ = v554
	var v555 int32
	_ = v555
	var v556 int32
	_ = v556
	var v558 int32
	_ = v558
	var v559 int32
	_ = v559
	var v560 int32
	_ = v560
	var v561 int32
	_ = v561
	var v566 int32
	_ = v566
	var v567 int32
	_ = v567
	var v570 int32
	_ = v570
	var v571 int32
	_ = v571
	var v574 int32
	_ = v574
	var v575 int32
	_ = v575
	var v576 int32
	_ = v576
	var v578 int32
	_ = v578
	var v580 int32
	_ = v580
	var v584 int32
	_ = v584
	var v585 int32
	_ = v585
	var v586 int32
	_ = v586
	var v588 int32
	_ = v588
	var v589 int32
	_ = v589
	var v592 int32
	_ = v592
	var v599 int32
	_ = v599
	var v602 int32
	_ = v602
	var v605 int32
	_ = v605
	var v608 int32
	_ = v608
	var v609 int32
	_ = v609
	var v612 int32
	_ = v612
	var v613 int32
	_ = v613
	var v614 int32
	_ = v614
	var v618 int32
	_ = v618
	var v619 int32
	_ = v619
	var v622 int32
	_ = v622
	var v624 int32
	_ = v624
	var v625 int32
	_ = v625
	var v626 int32
	_ = v626
	var v628 int32
	_ = v628
	var v632 int32
	_ = v632
	var v634 int32
	_ = v634
	var v635 int32
	_ = v635
	var v636 int32
	_ = v636
	var v638 int32
	_ = v638
	var v639 int32
	_ = v639
	var v640 int32
	_ = v640
	var v641 int32
	_ = v641
	var v646 int32
	_ = v646
	var v647 int32
	_ = v647
	var v650 int32
	_ = v650
	var v651 int32
	_ = v651
	var v654 int32
	_ = v654
	var v655 int32
	_ = v655
	var v656 int32
	_ = v656
	var v658 int32
	_ = v658
	var v660 int32
	_ = v660
	var v664 int32
	_ = v664
	var v667 int32
	_ = v667
	var v668 int32
	_ = v668
	var v671 int64
	_ = v671
	var v675 int32
	_ = v675
	var v678 int32
	_ = v678
	var v680 int32
	_ = v680
	var v681 int32
	_ = v681
	var v682 int32
	_ = v682
	var v684 int32
	_ = v684
	var v688 int32
	_ = v688
	var v690 int32
	_ = v690
	var v691 int32
	_ = v691
	var v692 int32
	_ = v692
	var v694 int32
	_ = v694
	var v695 int32
	_ = v695
	var v696 int32
	_ = v696
	var v697 int32
	_ = v697
	var v702 int32
	_ = v702
	var v703 int32
	_ = v703
	var v706 int32
	_ = v706
	var v707 int32
	_ = v707
	var v710 int32
	_ = v710
	var v711 int32
	_ = v711
	var v712 int32
	_ = v712
	var v716 int32
	_ = v716
	var v718 int32
	_ = v718
	var v719 int32
	_ = v719
	var v722 int32
	_ = v722
	var v723 int32
	_ = v723
	var v724 int32
	_ = v724
	var v726 int32
	_ = v726
	var v730 int32
	_ = v730
	var v731 int32
	_ = v731
	var v737 int32
	_ = v737
	var v744 int32
	_ = v744
	var v747 int32
	_ = v747
	var v750 int32
	_ = v750
	var v753 int32
	_ = v753
	var v754 int32
	_ = v754
	var v756 int32
	_ = v756
	var v764 int32
	_ = v764
	var v771 int32
	_ = v771
	var v774 int32
	_ = v774
	var v777 int32
	_ = v777
	var v780 int32
	_ = v780
	var v781 int32
	_ = v781
	var v783 int32
	_ = v783
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
	var v793 int32
	_ = v793
	var v794 int32
	_ = v794
	var v795 int32
	_ = v795
	var v796 int32
	_ = v796
	var v797 int32
	_ = v797
	var v798 int32
	_ = v798
	var v799 int32
	_ = v799
	var v800 int32
	_ = v800
	var v803 int32
	_ = v803
	var v805 int32
	_ = v805
	var v806 int32
	_ = v806
	var v807 int32
	_ = v807
	var v809 int32
	_ = v809
	var v813 int32
	_ = v813
	var v815 int32
	_ = v815
	var v816 int32
	_ = v816
	var v817 int32
	_ = v817
	var v819 int32
	_ = v819
	var v820 int32
	_ = v820
	var v821 int32
	_ = v821
	var v822 int32
	_ = v822
	var v827 int32
	_ = v827
	var v828 int32
	_ = v828
	var v831 int32
	_ = v831
	var v832 int32
	_ = v832
	var v835 int32
	_ = v835
	var v836 int32
	_ = v836
	var v837 int32
	_ = v837
	var v841 int32
	_ = v841
	var v844 int32
	_ = v844
	var v846 int32
	_ = v846
	var v847 int32
	_ = v847
	var v848 int32
	_ = v848
	var v850 int32
	_ = v850
	var v854 int32
	_ = v854
	var v856 int32
	_ = v856
	var v857 int32
	_ = v857
	var v858 int32
	_ = v858
	var v860 int32
	_ = v860
	var v861 int32
	_ = v861
	var v862 int32
	_ = v862
	var v863 int32
	_ = v863
	var v868 int32
	_ = v868
	var v869 int32
	_ = v869
	var v872 int32
	_ = v872
	var v873 int32
	_ = v873
	var v876 int32
	_ = v876
	var v877 int32
	_ = v877
	var v878 int32
	_ = v878
	var v883 int32
	_ = v883
	var v885 int32
	_ = v885
	var v889 int32
	_ = v889
	var v892 int32
	_ = v892
	var v893 int32
	_ = v893
	var v896 int64
	_ = v896
	var v901 int32
	_ = v901
	var v907 int32
	_ = v907
	var v908 int32
	_ = v908
	var v912 int32
	_ = v912
	var v914 int32
	_ = v914
	var v915 int32
	_ = v915
	var v920 int32
	_ = v920
	var v924 int32
	_ = v924
	var v925 int32
	_ = v925
	var v926 int32
	_ = v926
	var v935 int32
	_ = v935
	var v936 int32
	_ = v936
	var v940 int32
	_ = v940
	var v941 int32
	_ = v941
	var v942 int32
	_ = v942
	var v946 int32
	_ = v946
	var v947 int32
	_ = v947
	var v957 int32
	_ = v957
	var v963 int32
	_ = v963
	var v966 int32
	_ = v966
	var v969 int32
	_ = v969
	var v970 int32
	_ = v970
	var v973 int32
	_ = v973
	var v979 int32
	_ = v979
	var v986 int32
	_ = v986
	var v989 int32
	_ = v989
	var v992 int32
	_ = v992
	var v995 int32
	_ = v995
	var v996 int32
	_ = v996
	var v998 int32
	_ = v998
	var v999 int32
	_ = v999
	var v1001 int32
	_ = v1001
	var v1005 int32
	_ = v1005
	var v1009 int32
	_ = v1009
	var v1010 int32
	_ = v1010
	var v1017 int32
	_ = v1017
	var v1024 int32
	_ = v1024
	var v1033 int32
	_ = v1033
	var v1034 int32
	_ = v1034
	var v1039 int32
	_ = v1039
	var v1041 int32
	_ = v1041
	var v1042 int32
	_ = v1042
	var v1043 int32
	_ = v1043
	var v1046 int32
	_ = v1046
	var v1052 int32
	_ = v1052
	var v1053 int32
	_ = v1053
	var v1060 int32
	_ = v1060
	var v1067 int32
	_ = v1067
	v13 = m.G0
	v15 = v13 - int32(96)
	m.G0 = v15
	v18 = *(*int32)(unsafe.Add(mBase, _consts[733]))
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v20 = *(*int32)(unsafe.Add(mBase, uint32(v19)+8))
	v21 = F_objectGetVal(m, v20)
	mBase = m.M
	v22 = F_dictFetchValue(m, v18, v21)
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	m.G0 = v15 + int32(96)
	return
L2:
	;
	v29 = F_sentinelValidateArgs(m, l0, int32(3), int32(_a1526))
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		goto L3
	} else {
		goto L7
	}
L3:
	;
	return
L4:
	;
	if v22 != 0 {
		goto L2
	} else {
		goto L5
	}
L5:
	;
	F_addReplyError(m, l0, int32(_a1527))
	mBase = m.M
	v26 = m.ExcPending
	if v26 != 0 {
		goto L3
	} else {
		goto L6
	}
L6:
	;
	goto L1
L7:
	;
	if v29 == int32(-1) {
		goto L1
	} else {
		goto L8
	}
L8:
	;
	v33 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v33 < int32(4) {
		goto L1
	} else {
		goto L9
	}
L9:
	;
	v41 = v33
	v42 = int32(3)
	v43 = int32(0)
	goto L14
L10:
	;
	if v43 == int32(0) {
		goto L1
	} else {
		goto L312
	}
L11:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+80)) = v55
	F_addReplyErrorFormat(m, l0, int32(_a1528), v15+int32(80))
	mBase = m.M
	v1024 = m.ExcPending
	if v1024 != 0 {
		goto L3
	} else {
		goto L311
	}
L12:
	;
	v1005 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v1009 = *(*int32)(unsafe.Add(mBase, uint32(v1005+v1001<<(uint(int32(2))%32))))
	v1010 = F_objectGetVal(m, v1009)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v15)+68)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v15)+64)) = v1010
	F_addReplyErrorFormat(m, l0, int32(_a1529), v15+int32(64))
	mBase = m.M
	v1017 = m.ExcPending
	if v1017 != 0 {
		goto L3
	} else {
		goto L310
	}
L13:
	;
	v979 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v731+int32(-1)))))
	switch v979 & int32(7) {
	case 0:
		goto L306
	case 1:
		goto L305
	case 2:
		goto L304
	case 3:
		goto L303
	case 4:
		goto L302
	default:
		v996 = int32(0)
		goto L301
	}
L14:
	;
	v50 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v52 = v42 << (uint(int32(2)) % 32)
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v50+v52)))
	v55 = F_objectGetVal(m, v54)
	mBase = m.M
	v56 = int32(_a1470)
	v59 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v55))))
	if v59 != 0 {
		goto L18
	} else {
		goto L19
	}
L15:
	;
	F_sentinelFlushConfigAndReply(m, l0)
	mBase = m.M
	v973 = m.ExcPending
	if v973 != 0 {
		goto L3
	} else {
		goto L299
	}
L16:
	;
	v96 = v42 ^ int32(-1)
	v97 = v41 + v96
	v99 = base.B2i32(v97 < int32(1))
	if v97 < int32(1) {
		goto L30
	} else {
		goto L31
	}
L17:
	;
	v91 = F_tolower(m, v87)
	mBase = m.M
	v92 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v88))))
	v93 = F_tolower(m, v92)
	mBase = m.M
	goto L16
L18:
	;
	v61 = v55
	v62 = v56
	v63 = v59
	goto L21
L19:
	;
	v87 = int32(0)
	v88 = v56
	goto L17
L20:
	;
	v87 = v84 & int32(255)
	v88 = v83
	goto L17
L21:
	;
	v65 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v62))))
	if v65 == int32(0) {
		v83 = v62
		v84 = v63
		goto L20
	} else {
		goto L23
	}
L22:
	;
	v83 = v77
	v84 = int32(0)
	goto L20
L23:
	;
	v69 = v63 & int32(255)
	if v69 == v65 {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	v76 = int32(1)
	v77 = v62 + v76
	v78 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v61)+1)))
	if v78 != 0 {
		v61 = v61 + v76
		v62 = v77
		v63 = v78
		goto L21
	} else {
		goto L27
	}
L25:
	;
	v71 = F_tolower(m, v69)
	mBase = m.M
	v72 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v62))))
	v73 = F_tolower(m, v72)
	mBase = m.M
	if v71 == v73 {
		goto L24
	} else {
		goto L26
	}
L26:
	;
	v75 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v61))))
	v83 = v62
	v84 = v75
	goto L20
L27:
	;
	goto L22
L28:
	;
	v912 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v914 = *(*int32)(unsafe.Add(mBase, uint32(v912+v52)))
	v915 = F_objectGetVal(m, v914)
	mBase = m.M
	switch v96 + v907 {
	case 0:
		goto L292
	case 1:
		goto L291
	default:
		goto L290
	}
L29:
	;
	v907 = v901
	v908 = int32(1)
	goto L28
L30:
	;
	v119 = int32(_a1490)
	v122 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v55))))
	if v122 != 0 {
		goto L39
	} else {
		goto L40
	}
L31:
	;
	if v91-v93 != 0 {
		goto L30
	} else {
		goto L32
	}
L32:
	;
	v100 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v102 = v42 + int32(1)
	v106 = *(*int32)(unsafe.Add(mBase, uint32(v100+v102<<(uint(int32(2))%32))))
	v109 = F_getLongLongFromObject(m, v106, v15+int32(88))
	mBase = m.M
	v110 = m.ExcPending
	if v110 != 0 {
		goto L3
	} else {
		goto L33
	}
L33:
	;
	if v109 == int32(-1) {
		v1001 = v102
		goto L12
	} else {
		goto L34
	}
L34:
	;
	v113 = *(*int64)(unsafe.Add(mBase, uint32(v15)+88))
	if v113 < int64(1) {
		v1001 = v102
		goto L12
	} else {
		goto L35
	}
L35:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v22)+72)) = v113
	F_sentinelPropagateDownAfterPeriod(m, v22)
	mBase = m.M
	v118 = m.ExcPending
	if v118 != 0 {
		goto L3
	} else {
		goto L36
	}
L36:
	;
	v901 = v102
	goto L29
L37:
	;
	if v97 < int32(1) {
		goto L49
	} else {
		goto L50
	}
L38:
	;
	v154 = F_tolower(m, v150)
	mBase = m.M
	v155 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v151))))
	v156 = F_tolower(m, v155)
	mBase = m.M
	goto L37
L39:
	;
	v124 = v55
	v125 = v119
	v126 = v122
	goto L42
L40:
	;
	v150 = int32(0)
	v151 = v119
	goto L38
L41:
	;
	v150 = v147 & int32(255)
	v151 = v146
	goto L38
L42:
	;
	v128 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v125))))
	if v128 == int32(0) {
		v146 = v125
		v147 = v126
		goto L41
	} else {
		goto L44
	}
L43:
	;
	v146 = v140
	v147 = int32(0)
	goto L41
L44:
	;
	v132 = v126 & int32(255)
	if v132 == v128 {
		goto L45
	} else {
		goto L46
	}
L45:
	;
	v139 = int32(1)
	v140 = v125 + v139
	v141 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v124)+1)))
	if v141 != 0 {
		v124 = v124 + v139
		v125 = v140
		v126 = v141
		goto L42
	} else {
		goto L48
	}
L46:
	;
	v134 = F_tolower(m, v132)
	mBase = m.M
	v135 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v125))))
	v136 = F_tolower(m, v135)
	mBase = m.M
	if v134 == v136 {
		goto L45
	} else {
		goto L47
	}
L47:
	;
	v138 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v124))))
	v146 = v125
	v147 = v138
	goto L41
L48:
	;
	goto L43
L49:
	;
	v175 = int32(_a1491)
	v178 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v55))))
	if v178 != 0 {
		goto L57
	} else {
		goto L58
	}
L50:
	;
	if v154-v156 != 0 {
		goto L49
	} else {
		goto L51
	}
L51:
	;
	v158 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v160 = v42 + int32(1)
	v164 = *(*int32)(unsafe.Add(mBase, uint32(v158+v160<<(uint(int32(2))%32))))
	v167 = F_getLongLongFromObject(m, v164, v15+int32(88))
	mBase = m.M
	v168 = m.ExcPending
	if v168 != 0 {
		goto L3
	} else {
		goto L52
	}
L52:
	;
	if v167 == int32(-1) {
		v1001 = v160
		goto L12
	} else {
		goto L53
	}
L53:
	;
	v171 = *(*int64)(unsafe.Add(mBase, uint32(v15)+88))
	if v171 < int64(1) {
		v1001 = v160
		goto L12
	} else {
		goto L54
	}
L54:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v22)+264)) = v171
	v901 = v160
	goto L29
L55:
	;
	if v97 < int32(1) {
		goto L67
	} else {
		goto L68
	}
L56:
	;
	v210 = F_tolower(m, v206)
	mBase = m.M
	v211 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v207))))
	v212 = F_tolower(m, v211)
	mBase = m.M
	goto L55
L57:
	;
	v180 = v55
	v181 = v175
	v182 = v178
	goto L60
L58:
	;
	v206 = int32(0)
	v207 = v175
	goto L56
L59:
	;
	v206 = v203 & int32(255)
	v207 = v202
	goto L56
L60:
	;
	v184 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v181))))
	if v184 == int32(0) {
		v202 = v181
		v203 = v182
		goto L59
	} else {
		goto L62
	}
L61:
	;
	v202 = v196
	v203 = int32(0)
	goto L59
L62:
	;
	v188 = v182 & int32(255)
	if v188 == v184 {
		goto L63
	} else {
		goto L64
	}
L63:
	;
	v195 = int32(1)
	v196 = v181 + v195
	v197 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v180)+1)))
	if v197 != 0 {
		v180 = v180 + v195
		v181 = v196
		v182 = v197
		goto L60
	} else {
		goto L66
	}
L64:
	;
	v190 = F_tolower(m, v188)
	mBase = m.M
	v191 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v181))))
	v192 = F_tolower(m, v191)
	mBase = m.M
	if v190 == v192 {
		goto L63
	} else {
		goto L65
	}
L65:
	;
	v194 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v180))))
	v202 = v181
	v203 = v194
	goto L59
L66:
	;
	goto L61
L67:
	;
	v231 = int32(_a1492)
	v234 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v55))))
	if v234 != 0 {
		goto L75
	} else {
		goto L76
	}
L68:
	;
	if v210-v212 != 0 {
		goto L67
	} else {
		goto L69
	}
L69:
	;
	v214 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v216 = v42 + int32(1)
	v220 = *(*int32)(unsafe.Add(mBase, uint32(v214+v216<<(uint(int32(2))%32))))
	v223 = F_getLongLongFromObject(m, v220, v15+int32(88))
	mBase = m.M
	v224 = m.ExcPending
	if v224 != 0 {
		goto L3
	} else {
		goto L70
	}
L70:
	;
	if v223 == int32(-1) {
		v1001 = v216
		goto L12
	} else {
		goto L71
	}
L71:
	;
	v227 = *(*int64)(unsafe.Add(mBase, uint32(v15)+88))
	if v227 < int64(1) {
		v1001 = v216
		goto L12
	} else {
		goto L72
	}
L72:
	;
	*(*uint32)(unsafe.Add(mBase, uint32(v22)+156)) = uint32(v227)
	v901 = v216
	goto L29
L73:
	;
	if v97 < int32(1) {
		goto L85
	} else {
		goto L86
	}
L74:
	;
	v266 = F_tolower(m, v262)
	mBase = m.M
	v267 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v263))))
	v268 = F_tolower(m, v267)
	mBase = m.M
	goto L73
L75:
	;
	v236 = v55
	v237 = v231
	v238 = v234
	goto L78
L76:
	;
	v262 = int32(0)
	v263 = v231
	goto L74
L77:
	;
	v262 = v259 & int32(255)
	v263 = v258
	goto L74
L78:
	;
	v240 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v237))))
	if v240 == int32(0) {
		v258 = v237
		v259 = v238
		goto L77
	} else {
		goto L80
	}
L79:
	;
	v258 = v252
	v259 = int32(0)
	goto L77
L80:
	;
	v244 = v238 & int32(255)
	if v244 == v240 {
		goto L81
	} else {
		goto L82
	}
L81:
	;
	v251 = int32(1)
	v252 = v237 + v251
	v253 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v236)+1)))
	if v253 != 0 {
		v236 = v236 + v251
		v237 = v252
		v238 = v253
		goto L78
	} else {
		goto L84
	}
L82:
	;
	v246 = F_tolower(m, v244)
	mBase = m.M
	v247 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v237))))
	v248 = F_tolower(m, v247)
	mBase = m.M
	if v246 == v248 {
		goto L81
	} else {
		goto L83
	}
L83:
	;
	v250 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v236))))
	v258 = v237
	v259 = v250
	goto L77
L84:
	;
	goto L79
L85:
	;
	v344 = int32(_a1493)
	v347 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v55))))
	if v347 != 0 {
		goto L113
	} else {
		goto L114
	}
L86:
	;
	if v266-v268 != 0 {
		goto L85
	} else {
		goto L87
	}
L87:
	;
	v270 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v272 = v42 + int32(1)
	v276 = *(*int32)(unsafe.Add(mBase, uint32(v270+v272<<(uint(int32(2))%32))))
	v277 = F_objectGetVal(m, v276)
	mBase = m.M
	v278 = int32(0)
	v279 = *(*int32)(unsafe.Add(mBase, _consts[746]))
	if v279 == v278 {
		goto L88
	} else {
		goto L89
	}
L88:
	;
	v286 = v277 + int32(-1)
	v287 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v286))))
	switch v287 & int32(7) {
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
		goto L91
	}
L89:
	;
	F_addReplyError(m, l0, int32(_a1530))
	mBase = m.M
	v284 = m.ExcPending
	if v284 != 0 {
		goto L3
	} else {
		goto L90
	}
L90:
	;
	goto L10
L91:
	;
	v315 = *(*int32)(unsafe.Add(mBase, uint32(v22)+284))
	F_sdsfree(m, v315)
	mBase = m.M
	v317 = m.ExcPending
	if v317 != 0 {
		goto L3
	} else {
		goto L101
	}
L92:
	;
	if v304 == int32(0) {
		goto L91
	} else {
		goto L98
	}
L93:
	;
	v303 = *(*int32)(unsafe.Add(mBase, uint32(v277+int32(-17))))
	v304 = v303
	goto L92
L94:
	;
	v300 = *(*int32)(unsafe.Add(mBase, uint32(v277+int32(-9))))
	v304 = v300
	goto L92
L95:
	;
	v297 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v277+int32(-5)))))
	v304 = v297
	goto L92
L96:
	;
	v294 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v277+int32(-3)))))
	v304 = v294
	goto L92
L97:
	;
	v304 = int32(base.Ui32(v287) >> (uint(int32(3)) % 32))
	goto L92
L98:
	;
	v308 = F_access(m, v277, int32(1))
	mBase = m.M
	if v308 != int32(-1) {
		goto L91
	} else {
		goto L99
	}
L99:
	;
	F_addReplyError(m, l0, int32(_a1531))
	mBase = m.M
	v313 = m.ExcPending
	if v313 != 0 {
		goto L3
	} else {
		goto L100
	}
L100:
	;
	goto L10
L101:
	;
	v318 = int32(0)
	v319 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v286))))
	switch v319 & int32(7) {
	case 0:
		goto L108
	case 1:
		goto L107
	case 2:
		goto L106
	case 3:
		goto L105
	case 4:
		goto L104
	default:
		v341 = v318
		goto L102
	}
L102:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+284)) = v341
	v901 = v272
	goto L29
L103:
	;
	if v336 == int32(0) {
		v341 = v318
		goto L102
	} else {
		goto L109
	}
L104:
	;
	v335 = *(*int32)(unsafe.Add(mBase, uint32(v277+int32(-17))))
	v336 = v335
	goto L103
L105:
	;
	v332 = *(*int32)(unsafe.Add(mBase, uint32(v277+int32(-9))))
	v336 = v332
	goto L103
L106:
	;
	v329 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v277+int32(-5)))))
	v336 = v329
	goto L103
L107:
	;
	v326 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v277+int32(-3)))))
	v336 = v326
	goto L103
L108:
	;
	v336 = int32(base.Ui32(v319) >> (uint(int32(3)) % 32))
	goto L103
L109:
	;
	v339 = F_sdsdup(m, v277)
	mBase = m.M
	v340 = m.ExcPending
	if v340 != 0 {
		goto L3
	} else {
		goto L110
	}
L110:
	;
	v341 = v339
	goto L102
L111:
	;
	if v97 < int32(1) {
		goto L123
	} else {
		goto L124
	}
L112:
	;
	v379 = F_tolower(m, v375)
	mBase = m.M
	v380 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v376))))
	v381 = F_tolower(m, v380)
	mBase = m.M
	goto L111
L113:
	;
	v349 = v55
	v350 = v344
	v351 = v347
	goto L116
L114:
	;
	v375 = int32(0)
	v376 = v344
	goto L112
L115:
	;
	v375 = v372 & int32(255)
	v376 = v371
	goto L112
L116:
	;
	v353 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v350))))
	if v353 == int32(0) {
		v371 = v350
		v372 = v351
		goto L115
	} else {
		goto L118
	}
L117:
	;
	v371 = v365
	v372 = int32(0)
	goto L115
L118:
	;
	v357 = v351 & int32(255)
	if v357 == v353 {
		goto L119
	} else {
		goto L120
	}
L119:
	;
	v364 = int32(1)
	v365 = v350 + v364
	v366 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v349)+1)))
	if v366 != 0 {
		v349 = v349 + v364
		v350 = v365
		v351 = v366
		goto L116
	} else {
		goto L122
	}
L120:
	;
	v359 = F_tolower(m, v357)
	mBase = m.M
	v360 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v350))))
	v361 = F_tolower(m, v360)
	mBase = m.M
	if v359 == v361 {
		goto L119
	} else {
		goto L121
	}
L121:
	;
	v363 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v349))))
	v371 = v350
	v372 = v363
	goto L115
L122:
	;
	goto L117
L123:
	;
	v457 = int32(_a1532)
	v460 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v55))))
	if v460 != 0 {
		goto L151
	} else {
		goto L152
	}
L124:
	;
	if v379-v381 != 0 {
		goto L123
	} else {
		goto L125
	}
L125:
	;
	v383 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v385 = v42 + int32(1)
	v389 = *(*int32)(unsafe.Add(mBase, uint32(v383+v385<<(uint(int32(2))%32))))
	v390 = F_objectGetVal(m, v389)
	mBase = m.M
	v391 = int32(0)
	v392 = *(*int32)(unsafe.Add(mBase, _consts[746]))
	if v392 == v391 {
		goto L126
	} else {
		goto L127
	}
L126:
	;
	v399 = v390 + int32(-1)
	v400 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v399))))
	switch v400 & int32(7) {
	case 0:
		goto L135
	case 1:
		goto L134
	case 2:
		goto L133
	case 3:
		goto L132
	case 4:
		goto L131
	default:
		goto L129
	}
L127:
	;
	F_addReplyError(m, l0, int32(_a1530))
	mBase = m.M
	v397 = m.ExcPending
	if v397 != 0 {
		goto L3
	} else {
		goto L128
	}
L128:
	;
	goto L10
L129:
	;
	v428 = *(*int32)(unsafe.Add(mBase, uint32(v22)+288))
	F_sdsfree(m, v428)
	mBase = m.M
	v430 = m.ExcPending
	if v430 != 0 {
		goto L3
	} else {
		goto L139
	}
L130:
	;
	if v417 == int32(0) {
		goto L129
	} else {
		goto L136
	}
L131:
	;
	v416 = *(*int32)(unsafe.Add(mBase, uint32(v390+int32(-17))))
	v417 = v416
	goto L130
L132:
	;
	v413 = *(*int32)(unsafe.Add(mBase, uint32(v390+int32(-9))))
	v417 = v413
	goto L130
L133:
	;
	v410 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v390+int32(-5)))))
	v417 = v410
	goto L130
L134:
	;
	v407 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v390+int32(-3)))))
	v417 = v407
	goto L130
L135:
	;
	v417 = int32(base.Ui32(v400) >> (uint(int32(3)) % 32))
	goto L130
L136:
	;
	v421 = F_access(m, v390, int32(1))
	mBase = m.M
	if v421 != int32(-1) {
		goto L129
	} else {
		goto L137
	}
L137:
	;
	F_addReplyError(m, l0, int32(_a1533))
	mBase = m.M
	v426 = m.ExcPending
	if v426 != 0 {
		goto L3
	} else {
		goto L138
	}
L138:
	;
	goto L10
L139:
	;
	v431 = int32(0)
	v432 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v399))))
	switch v432 & int32(7) {
	case 0:
		goto L146
	case 1:
		goto L145
	case 2:
		goto L144
	case 3:
		goto L143
	case 4:
		goto L142
	default:
		v454 = v431
		goto L140
	}
L140:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+288)) = v454
	v901 = v385
	goto L29
L141:
	;
	if v449 == int32(0) {
		v454 = v431
		goto L140
	} else {
		goto L147
	}
L142:
	;
	v448 = *(*int32)(unsafe.Add(mBase, uint32(v390+int32(-17))))
	v449 = v448
	goto L141
L143:
	;
	v445 = *(*int32)(unsafe.Add(mBase, uint32(v390+int32(-9))))
	v449 = v445
	goto L141
L144:
	;
	v442 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v390+int32(-5)))))
	v449 = v442
	goto L141
L145:
	;
	v439 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v390+int32(-3)))))
	v449 = v439
	goto L141
L146:
	;
	v449 = int32(base.Ui32(v432) >> (uint(int32(3)) % 32))
	goto L141
L147:
	;
	v452 = F_sdsdup(m, v390)
	mBase = m.M
	v453 = m.ExcPending
	if v453 != 0 {
		goto L3
	} else {
		goto L148
	}
L148:
	;
	v454 = v452
	goto L140
L149:
	;
	if v97 < int32(1) {
		goto L161
	} else {
		goto L162
	}
L150:
	;
	v492 = F_tolower(m, v488)
	mBase = m.M
	v493 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v489))))
	v494 = F_tolower(m, v493)
	mBase = m.M
	goto L149
L151:
	;
	v462 = v55
	v463 = v457
	v464 = v460
	goto L154
L152:
	;
	v488 = int32(0)
	v489 = v457
	goto L150
L153:
	;
	v488 = v485 & int32(255)
	v489 = v484
	goto L150
L154:
	;
	v466 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v463))))
	if v466 == int32(0) {
		v484 = v463
		v485 = v464
		goto L153
	} else {
		goto L156
	}
L155:
	;
	v484 = v478
	v485 = int32(0)
	goto L153
L156:
	;
	v470 = v464 & int32(255)
	if v470 == v466 {
		goto L157
	} else {
		goto L158
	}
L157:
	;
	v477 = int32(1)
	v478 = v463 + v477
	v479 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v462)+1)))
	if v479 != 0 {
		v462 = v462 + v477
		v463 = v478
		v464 = v479
		goto L154
	} else {
		goto L160
	}
L158:
	;
	v472 = F_tolower(m, v470)
	mBase = m.M
	v473 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v463))))
	v474 = F_tolower(m, v473)
	mBase = m.M
	if v472 == v474 {
		goto L157
	} else {
		goto L159
	}
L159:
	;
	v476 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v462))))
	v484 = v463
	v485 = v476
	goto L153
L160:
	;
	goto L155
L161:
	;
	v539 = int32(_a1534)
	v542 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v55))))
	if v542 != 0 {
		goto L177
	} else {
		goto L178
	}
L162:
	;
	if v492-v494 != 0 {
		goto L161
	} else {
		goto L163
	}
L163:
	;
	v496 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v498 = v42 + int32(1)
	v502 = *(*int32)(unsafe.Add(mBase, uint32(v496+v498<<(uint(int32(2))%32))))
	v503 = F_objectGetVal(m, v502)
	mBase = m.M
	v504 = *(*int32)(unsafe.Add(mBase, uint32(v22)+160))
	F_sdsfree(m, v504)
	mBase = m.M
	v506 = m.ExcPending
	if v506 != 0 {
		goto L3
	} else {
		goto L164
	}
L164:
	;
	v507 = int32(0)
	v511 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v503+int32(-1)))))
	switch v511 & int32(7) {
	case 0:
		goto L171
	case 1:
		goto L170
	case 2:
		goto L169
	case 3:
		goto L168
	case 4:
		goto L167
	default:
		v534 = v507
		goto L165
	}
L165:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+160)) = v534
	F_dropInstanceConnections(m, v22)
	mBase = m.M
	v538 = m.ExcPending
	if v538 != 0 {
		goto L3
	} else {
		goto L174
	}
L166:
	;
	v529 = int32(0)
	if v528 == v529 {
		v534 = v529
		goto L165
	} else {
		goto L172
	}
L167:
	;
	v527 = *(*int32)(unsafe.Add(mBase, uint32(v503+int32(-17))))
	v528 = v527
	goto L166
L168:
	;
	v524 = *(*int32)(unsafe.Add(mBase, uint32(v503+int32(-9))))
	v528 = v524
	goto L166
L169:
	;
	v521 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v503+int32(-5)))))
	v528 = v521
	goto L166
L170:
	;
	v518 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v503+int32(-3)))))
	v528 = v518
	goto L166
L171:
	;
	v528 = int32(base.Ui32(v511) >> (uint(int32(3)) % 32))
	goto L166
L172:
	;
	v532 = F_sdsdup(m, v503)
	mBase = m.M
	v533 = m.ExcPending
	if v533 != 0 {
		goto L3
	} else {
		goto L173
	}
L173:
	;
	v534 = v532
	goto L165
L174:
	;
	v907 = v498
	v908 = v507
	goto L28
L175:
	;
	if v97 < int32(1) {
		goto L187
	} else {
		goto L188
	}
L176:
	;
	v574 = F_tolower(m, v570)
	mBase = m.M
	v575 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v571))))
	v576 = F_tolower(m, v575)
	mBase = m.M
	goto L175
L177:
	;
	v544 = v55
	v545 = v539
	v546 = v542
	goto L180
L178:
	;
	v570 = int32(0)
	v571 = v539
	goto L176
L179:
	;
	v570 = v567 & int32(255)
	v571 = v566
	goto L176
L180:
	;
	v548 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v545))))
	if v548 == int32(0) {
		v566 = v545
		v567 = v546
		goto L179
	} else {
		goto L182
	}
L181:
	;
	v566 = v560
	v567 = int32(0)
	goto L179
L182:
	;
	v552 = v546 & int32(255)
	if v552 == v548 {
		goto L183
	} else {
		goto L184
	}
L183:
	;
	v559 = int32(1)
	v560 = v545 + v559
	v561 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v544)+1)))
	if v561 != 0 {
		v544 = v544 + v559
		v545 = v560
		v546 = v561
		goto L180
	} else {
		goto L186
	}
L184:
	;
	v554 = F_tolower(m, v552)
	mBase = m.M
	v555 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v545))))
	v556 = F_tolower(m, v555)
	mBase = m.M
	if v554 == v556 {
		goto L183
	} else {
		goto L185
	}
L185:
	;
	v558 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v544))))
	v566 = v545
	v567 = v558
	goto L179
L186:
	;
	goto L181
L187:
	;
	v619 = int32(_a1489)
	v622 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v55))))
	if v622 != 0 {
		goto L203
	} else {
		goto L204
	}
L188:
	;
	if v574-v576 != 0 {
		goto L187
	} else {
		goto L189
	}
L189:
	;
	v578 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v580 = v42 + int32(1)
	v584 = *(*int32)(unsafe.Add(mBase, uint32(v578+v580<<(uint(int32(2))%32))))
	v585 = F_objectGetVal(m, v584)
	mBase = m.M
	v586 = *(*int32)(unsafe.Add(mBase, uint32(v22)+164))
	F_sdsfree(m, v586)
	mBase = m.M
	v588 = m.ExcPending
	if v588 != 0 {
		goto L3
	} else {
		goto L190
	}
L190:
	;
	v589 = int32(0)
	v592 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v585+int32(-1)))))
	switch v592 & int32(7) {
	case 0:
		goto L197
	case 1:
		goto L196
	case 2:
		goto L195
	case 3:
		goto L194
	case 4:
		goto L193
	default:
		v614 = v589
		goto L191
	}
L191:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+164)) = v614
	F_dropInstanceConnections(m, v22)
	mBase = m.M
	v618 = m.ExcPending
	if v618 != 0 {
		goto L3
	} else {
		goto L200
	}
L192:
	;
	if v609 == int32(0) {
		v614 = v589
		goto L191
	} else {
		goto L198
	}
L193:
	;
	v608 = *(*int32)(unsafe.Add(mBase, uint32(v585+int32(-17))))
	v609 = v608
	goto L192
L194:
	;
	v605 = *(*int32)(unsafe.Add(mBase, uint32(v585+int32(-9))))
	v609 = v605
	goto L192
L195:
	;
	v602 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v585+int32(-5)))))
	v609 = v602
	goto L192
L196:
	;
	v599 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v585+int32(-3)))))
	v609 = v599
	goto L192
L197:
	;
	v609 = int32(base.Ui32(v592) >> (uint(int32(3)) % 32))
	goto L192
L198:
	;
	v612 = F_sdsdup(m, v585)
	mBase = m.M
	v613 = m.ExcPending
	if v613 != 0 {
		goto L3
	} else {
		goto L199
	}
L199:
	;
	v614 = v612
	goto L191
L200:
	;
	v901 = v580
	goto L29
L201:
	;
	if v97 < int32(1) {
		goto L213
	} else {
		goto L214
	}
L202:
	;
	v654 = F_tolower(m, v650)
	mBase = m.M
	v655 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v651))))
	v656 = F_tolower(m, v655)
	mBase = m.M
	goto L201
L203:
	;
	v624 = v55
	v625 = v619
	v626 = v622
	goto L206
L204:
	;
	v650 = int32(0)
	v651 = v619
	goto L202
L205:
	;
	v650 = v647 & int32(255)
	v651 = v646
	goto L202
L206:
	;
	v628 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v625))))
	if v628 == int32(0) {
		v646 = v625
		v647 = v626
		goto L205
	} else {
		goto L208
	}
L207:
	;
	v646 = v640
	v647 = int32(0)
	goto L205
L208:
	;
	v632 = v626 & int32(255)
	if v632 == v628 {
		goto L209
	} else {
		goto L210
	}
L209:
	;
	v639 = int32(1)
	v640 = v625 + v639
	v641 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v624)+1)))
	if v641 != 0 {
		v624 = v624 + v639
		v625 = v640
		v626 = v641
		goto L206
	} else {
		goto L212
	}
L210:
	;
	v634 = F_tolower(m, v632)
	mBase = m.M
	v635 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v625))))
	v636 = F_tolower(m, v635)
	mBase = m.M
	if v634 == v636 {
		goto L209
	} else {
		goto L211
	}
L211:
	;
	v638 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v624))))
	v646 = v625
	v647 = v638
	goto L205
L212:
	;
	goto L207
L213:
	;
	v675 = int32(_a500)
	v678 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v55))))
	if v678 != 0 {
		goto L221
	} else {
		goto L222
	}
L214:
	;
	if v654-v656 != 0 {
		goto L213
	} else {
		goto L215
	}
L215:
	;
	v658 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v660 = v42 + int32(1)
	v664 = *(*int32)(unsafe.Add(mBase, uint32(v658+v660<<(uint(int32(2))%32))))
	v667 = F_getLongLongFromObject(m, v664, v15+int32(88))
	mBase = m.M
	v668 = m.ExcPending
	if v668 != 0 {
		goto L3
	} else {
		goto L216
	}
L216:
	;
	if v667 == int32(-1) {
		v1001 = v660
		goto L12
	} else {
		goto L217
	}
L217:
	;
	v671 = *(*int64)(unsafe.Add(mBase, uint32(v15)+88))
	if v671 < int64(1) {
		v1001 = v660
		goto L12
	} else {
		goto L218
	}
L218:
	;
	*(*uint32)(unsafe.Add(mBase, uint32(v22)+152)) = uint32(v671)
	v901 = v660
	goto L29
L219:
	;
	if v97 < int32(2) {
		goto L231
	} else {
		goto L232
	}
L220:
	;
	v710 = F_tolower(m, v706)
	mBase = m.M
	v711 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v707))))
	v712 = F_tolower(m, v711)
	mBase = m.M
	goto L219
L221:
	;
	v680 = v55
	v681 = v675
	v682 = v678
	goto L224
L222:
	;
	v706 = int32(0)
	v707 = v675
	goto L220
L223:
	;
	v706 = v703 & int32(255)
	v707 = v702
	goto L220
L224:
	;
	v684 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v681))))
	if v684 == int32(0) {
		v702 = v681
		v703 = v682
		goto L223
	} else {
		goto L226
	}
L225:
	;
	v702 = v696
	v703 = int32(0)
	goto L223
L226:
	;
	v688 = v682 & int32(255)
	if v688 == v684 {
		goto L227
	} else {
		goto L228
	}
L227:
	;
	v695 = int32(1)
	v696 = v681 + v695
	v697 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v680)+1)))
	if v697 != 0 {
		v680 = v680 + v695
		v681 = v696
		v682 = v697
		goto L224
	} else {
		goto L230
	}
L228:
	;
	v690 = F_tolower(m, v688)
	mBase = m.M
	v691 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v681))))
	v692 = F_tolower(m, v691)
	mBase = m.M
	if v690 == v692 {
		goto L227
	} else {
		goto L229
	}
L229:
	;
	v694 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v680))))
	v702 = v681
	v703 = v694
	goto L223
L230:
	;
	goto L225
L231:
	;
	v800 = int32(_a1535)
	v803 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v55))))
	if v803 != 0 {
		goto L260
	} else {
		goto L261
	}
L232:
	;
	if v710-v712 != 0 {
		goto L231
	} else {
		goto L233
	}
L233:
	;
	v716 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v718 = v42 + int32(1)
	v719 = int32(2)
	v722 = *(*int32)(unsafe.Add(mBase, uint32(v716+v718<<(uint(v719)%32))))
	v723 = F_objectGetVal(m, v722)
	mBase = m.M
	v724 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v726 = v42 + v719
	v730 = *(*int32)(unsafe.Add(mBase, uint32(v724+v726<<(uint(v719)%32))))
	v731 = F_objectGetVal(m, v730)
	mBase = m.M
	v737 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v723+int32(-1)))))
	switch v737 & int32(7) {
	case 0:
		goto L240
	case 1:
		goto L239
	case 2:
		goto L238
	case 3:
		goto L237
	case 4:
		goto L236
	default:
		v754 = int32(0)
		goto L235
	}
L234:
	;
	if v756 == int32(0) {
		goto L13
	} else {
		goto L241
	}
L235:
	;
	v756 = v754
	goto L234
L236:
	;
	v753 = *(*int32)(unsafe.Add(mBase, uint32(v723+int32(-17))))
	v754 = v753
	goto L235
L237:
	;
	v750 = *(*int32)(unsafe.Add(mBase, uint32(v723+int32(-9))))
	v756 = v750
	goto L234
L238:
	;
	v747 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v723+int32(-5)))))
	v756 = v747
	goto L234
L239:
	;
	v744 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v723+int32(-3)))))
	v756 = v744
	goto L234
L240:
	;
	v756 = int32(base.Ui32(v737) >> (uint(int32(3)) % 32))
	goto L234
L241:
	;
	v764 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v731+int32(-1)))))
	switch v764 & int32(7) {
	case 0:
		goto L248
	case 1:
		goto L247
	case 2:
		goto L246
	case 3:
		goto L245
	case 4:
		goto L244
	default:
		v781 = int32(0)
		goto L243
	}
L242:
	;
	if v783 == int32(0) {
		goto L13
	} else {
		goto L249
	}
L243:
	;
	v783 = v781
	goto L242
L244:
	;
	v780 = *(*int32)(unsafe.Add(mBase, uint32(v731+int32(-17))))
	v781 = v780
	goto L243
L245:
	;
	v777 = *(*int32)(unsafe.Add(mBase, uint32(v731+int32(-9))))
	v783 = v777
	goto L242
L246:
	;
	v774 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v731+int32(-5)))))
	v783 = v774
	goto L242
L247:
	;
	v771 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v731+int32(-3)))))
	v783 = v771
	goto L242
L248:
	;
	v783 = int32(base.Ui32(v764) >> (uint(int32(3)) % 32))
	goto L242
L249:
	;
	v786 = *(*int32)(unsafe.Add(mBase, uint32(v22)+104))
	v787 = F_dictDelete(m, v786, v723)
	mBase = m.M
	v788 = m.ExcPending
	if v788 != 0 {
		goto L3
	} else {
		goto L250
	}
L250:
	;
	v789 = int32(1)
	v790 = F_strcasecmp(m, v723, v731)
	mBase = m.M
	goto L251
L251:
	;
	if v790 == int32(0) {
		v907 = v726
		v908 = v789
		goto L28
	} else {
		goto L252
	}
L252:
	;
	v793 = F_sdsdup(m, v723)
	mBase = m.M
	v794 = m.ExcPending
	if v794 != 0 {
		goto L3
	} else {
		goto L253
	}
L253:
	;
	v795 = F_sdsdup(m, v731)
	mBase = m.M
	v796 = m.ExcPending
	if v796 != 0 {
		goto L3
	} else {
		goto L254
	}
L254:
	;
	v797 = *(*int32)(unsafe.Add(mBase, uint32(v22)+104))
	v798 = F_dictAdd(m, v797, v793, v795)
	mBase = m.M
	v799 = m.ExcPending
	if v799 != 0 {
		goto L3
	} else {
		goto L255
	}
L255:
	;
	v907 = v726
	v908 = v789
	goto L28
L256:
	;
	v883 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v885 = v42 + int32(1)
	v889 = *(*int32)(unsafe.Add(mBase, uint32(v883+v885<<(uint(int32(2))%32))))
	v892 = F_getLongLongFromObject(m, v889, v15+int32(88))
	mBase = m.M
	v893 = m.ExcPending
	if v893 != 0 {
		goto L3
	} else {
		goto L286
	}
L257:
	;
	if v97 < int32(1) {
		goto L11
	} else {
		goto L285
	}
L258:
	;
	if v835-v837 == int32(0) {
		goto L257
	} else {
		goto L270
	}
L259:
	;
	v835 = F_tolower(m, v831)
	mBase = m.M
	v836 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v832))))
	v837 = F_tolower(m, v836)
	mBase = m.M
	goto L258
L260:
	;
	v805 = v55
	v806 = v800
	v807 = v803
	goto L263
L261:
	;
	v831 = int32(0)
	v832 = v800
	goto L259
L262:
	;
	v831 = v828 & int32(255)
	v832 = v827
	goto L259
L263:
	;
	v809 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v806))))
	if v809 == int32(0) {
		v827 = v806
		v828 = v807
		goto L262
	} else {
		goto L265
	}
L264:
	;
	v827 = v821
	v828 = int32(0)
	goto L262
L265:
	;
	v813 = v807 & int32(255)
	if v813 == v809 {
		goto L266
	} else {
		goto L267
	}
L266:
	;
	v820 = int32(1)
	v821 = v806 + v820
	v822 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v805)+1)))
	if v822 != 0 {
		v805 = v805 + v820
		v806 = v821
		v807 = v822
		goto L263
	} else {
		goto L269
	}
L267:
	;
	v815 = F_tolower(m, v813)
	mBase = m.M
	v816 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v806))))
	v817 = F_tolower(m, v816)
	mBase = m.M
	if v815 == v817 {
		goto L266
	} else {
		goto L268
	}
L268:
	;
	v819 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v805))))
	v827 = v806
	v828 = v819
	goto L262
L269:
	;
	goto L264
L270:
	;
	v841 = int32(_a1536)
	v844 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v55))))
	if v844 != 0 {
		goto L273
	} else {
		goto L274
	}
L271:
	;
	if v97 < int32(1) {
		goto L11
	} else {
		goto L283
	}
L272:
	;
	v876 = F_tolower(m, v872)
	mBase = m.M
	v877 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v873))))
	v878 = F_tolower(m, v877)
	mBase = m.M
	goto L271
L273:
	;
	v846 = v55
	v847 = v841
	v848 = v844
	goto L276
L274:
	;
	v872 = int32(0)
	v873 = v841
	goto L272
L275:
	;
	v872 = v869 & int32(255)
	v873 = v868
	goto L272
L276:
	;
	v850 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v847))))
	if v850 == int32(0) {
		v868 = v847
		v869 = v848
		goto L275
	} else {
		goto L278
	}
L277:
	;
	v868 = v862
	v869 = int32(0)
	goto L275
L278:
	;
	v854 = v848 & int32(255)
	if v854 == v850 {
		goto L279
	} else {
		goto L280
	}
L279:
	;
	v861 = int32(1)
	v862 = v847 + v861
	v863 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v846)+1)))
	if v863 != 0 {
		v846 = v846 + v861
		v847 = v862
		v848 = v863
		goto L276
	} else {
		goto L282
	}
L280:
	;
	v856 = F_tolower(m, v854)
	mBase = m.M
	v857 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v847))))
	v858 = F_tolower(m, v857)
	mBase = m.M
	if v856 == v858 {
		goto L279
	} else {
		goto L281
	}
L281:
	;
	v860 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v846))))
	v868 = v847
	v869 = v860
	goto L275
L282:
	;
	goto L277
L283:
	;
	if v876-v878 == int32(0) {
		goto L256
	} else {
		goto L284
	}
L284:
	;
	goto L11
L285:
	;
	goto L256
L286:
	;
	if v892 == int32(-1) {
		v1001 = v885
		goto L12
	} else {
		goto L287
	}
L287:
	;
	v896 = *(*int64)(unsafe.Add(mBase, uint32(v15)+88))
	if v896 < int64(0) {
		v1001 = v885
		goto L12
	} else {
		goto L288
	}
L288:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v22)+80)) = v896
	v901 = v885
	goto L29
L289:
	;
	v966 = int32(1)
	v969 = v907 + v966
	v970 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v969 < v970 {
		v41 = v970
		v42 = v969
		v43 = v43 + v966
		goto L14
	} else {
		goto L298
	}
L290:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15))) = v915
	F_sentinelEvent(m, int32(3), int32(_a1537), v22, int32(_a1538), v15)
	mBase = m.M
	v963 = m.ExcPending
	if v963 != 0 {
		goto L3
	} else {
		goto L297
	}
L291:
	;
	v936 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v940 = *(*int32)(unsafe.Add(mBase, uint32(v936+v52+int32(4))))
	v941 = F_objectGetVal(m, v940)
	mBase = m.M
	v942 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v946 = *(*int32)(unsafe.Add(mBase, uint32(v942+v52+int32(8))))
	v947 = F_objectGetVal(m, v946)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v15)+40)) = v947
	*(*int32)(unsafe.Add(mBase, uint32(v15)+36)) = v941
	*(*int32)(unsafe.Add(mBase, uint32(v15)+32)) = v915
	F_sentinelEvent(m, int32(3), int32(_a1537), v22, int32(_a1539), v15+int32(32))
	mBase = m.M
	v957 = m.ExcPending
	if v957 != 0 {
		goto L3
	} else {
		goto L296
	}
L292:
	;
	if v908 == int32(0) {
		v926 = int32(_a1540)
		goto L293
	} else {
		goto L294
	}
L293:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+20)) = v926
	*(*int32)(unsafe.Add(mBase, uint32(v15)+16)) = v915
	F_sentinelEvent(m, int32(3), int32(_a1537), v22, int32(_a1541), v15+int32(16))
	mBase = m.M
	v935 = m.ExcPending
	if v935 != 0 {
		goto L3
	} else {
		goto L295
	}
L294:
	;
	v920 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v924 = *(*int32)(unsafe.Add(mBase, uint32(v920+v52+int32(4))))
	v925 = F_objectGetVal(m, v924)
	mBase = m.M
	v926 = v925
	goto L293
L295:
	;
	goto L289
L296:
	;
	goto L289
L297:
	;
	goto L289
L298:
	;
	goto L15
L299:
	;
	goto L1
L300:
	;
	if v998 != 0 {
		goto L307
	} else {
		goto L308
	}
L301:
	;
	v998 = v996
	goto L300
L302:
	;
	v995 = *(*int32)(unsafe.Add(mBase, uint32(v731+int32(-17))))
	v996 = v995
	goto L301
L303:
	;
	v992 = *(*int32)(unsafe.Add(mBase, uint32(v731+int32(-9))))
	v998 = v992
	goto L300
L304:
	;
	v989 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v731+int32(-5)))))
	v998 = v989
	goto L300
L305:
	;
	v986 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v731+int32(-3)))))
	v998 = v986
	goto L300
L306:
	;
	v998 = int32(base.Ui32(v979) >> (uint(int32(3)) % 32))
	goto L300
L307:
	;
	v999 = v718
	goto L309
L308:
	;
	v999 = v726
	goto L309
L309:
	;
	v1001 = v999
	goto L12
L310:
	;
	goto L10
L311:
	;
	goto L10
L312:
	;
	v1033 = int32(_a44)
	v1034 = *(*int32)(unsafe.Add(mBase, _consts[219]))
	*(*int32)(unsafe.Add(mBase, _consts[219])) = int32(10)
	v1039 = *(*int32)(unsafe.Add(mBase, _consts[392]))
	v1041 = F_rewriteConfig(m, v1039, int32(0))
	mBase = m.M
	v1042 = m.ExcPending
	if v1042 != 0 {
		goto L3
	} else {
		goto L313
	}
L313:
	;
	v1043 = int32(_a44)
	*(*int32)(unsafe.Add(mBase, _consts[219])) = v1034
	v1046 = *(*int32)(unsafe.Add(mBase, _consts[6]))
	if v1041 != int32(-1) {
		goto L314
	} else {
		goto L315
	}
L314:
	;
	if int32(2) < v1046 {
		goto L1
	} else {
		goto L320
	}
L315:
	;
	if int32(3) < v1046 {
		goto L1
	} else {
		goto L316
	}
L316:
	;
	goto L317
L317:
	;
	v1052 = *(*int32)(unsafe.Add(mBase, _consts[5]))
	v1053 = F___strerror_l(m, v1052, v1052)
	mBase = m.M
	goto L318
L318:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+48)) = v1053
	F__serverLog(m, int32(3), int32(_a1337), v15+int32(48))
	mBase = m.M
	v1060 = m.ExcPending
	if v1060 != 0 {
		goto L3
	} else {
		goto L319
	}
L319:
	;
	goto L1
L320:
	;
	F__serverLog(m, int32(2), int32(_a1338), int32(0))
	mBase = m.M
	v1067 = m.ExcPending
	if v1067 != 0 {
		goto L3
	} else {
		goto L321
	}
L321:
	;
	goto L1
}
func F_sentinelSetDebugConfigParameters(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
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
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
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
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v84 int64
	_ = v84
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v102 int32
	_ = v102
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
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
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v128 int32
	_ = v128
	var v130 int32
	_ = v130
	var v134 int32
	_ = v134
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v141 int64
	_ = v141
	var v146 int32
	_ = v146
	var v149 int32
	_ = v149
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v155 int32
	_ = v155
	var v159 int32
	_ = v159
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
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
	var v191 int32
	_ = v191
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v198 int64
	_ = v198
	var v203 int32
	_ = v203
	var v206 int32
	_ = v206
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v212 int32
	_ = v212
	var v216 int32
	_ = v216
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v242 int32
	_ = v242
	var v244 int32
	_ = v244
	var v248 int32
	_ = v248
	var v251 int32
	_ = v251
	var v252 int32
	_ = v252
	var v255 int64
	_ = v255
	var v260 int32
	_ = v260
	var v263 int32
	_ = v263
	var v265 int32
	_ = v265
	var v266 int32
	_ = v266
	var v267 int32
	_ = v267
	var v269 int32
	_ = v269
	var v273 int32
	_ = v273
	var v275 int32
	_ = v275
	var v276 int32
	_ = v276
	var v277 int32
	_ = v277
	var v279 int32
	_ = v279
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
	var v291 int32
	_ = v291
	var v292 int32
	_ = v292
	var v295 int32
	_ = v295
	var v296 int32
	_ = v296
	var v297 int32
	_ = v297
	var v299 int32
	_ = v299
	var v301 int32
	_ = v301
	var v305 int32
	_ = v305
	var v308 int32
	_ = v308
	var v309 int32
	_ = v309
	var v312 int64
	_ = v312
	var v317 int32
	_ = v317
	var v320 int32
	_ = v320
	var v322 int32
	_ = v322
	var v323 int32
	_ = v323
	var v324 int32
	_ = v324
	var v326 int32
	_ = v326
	var v330 int32
	_ = v330
	var v332 int32
	_ = v332
	var v333 int32
	_ = v333
	var v334 int32
	_ = v334
	var v336 int32
	_ = v336
	var v337 int32
	_ = v337
	var v338 int32
	_ = v338
	var v339 int32
	_ = v339
	var v344 int32
	_ = v344
	var v345 int32
	_ = v345
	var v348 int32
	_ = v348
	var v349 int32
	_ = v349
	var v352 int32
	_ = v352
	var v353 int32
	_ = v353
	var v354 int32
	_ = v354
	var v356 int32
	_ = v356
	var v358 int32
	_ = v358
	var v362 int32
	_ = v362
	var v365 int32
	_ = v365
	var v366 int32
	_ = v366
	var v369 int64
	_ = v369
	var v374 int32
	_ = v374
	var v377 int32
	_ = v377
	var v379 int32
	_ = v379
	var v380 int32
	_ = v380
	var v381 int32
	_ = v381
	var v383 int32
	_ = v383
	var v387 int32
	_ = v387
	var v389 int32
	_ = v389
	var v390 int32
	_ = v390
	var v391 int32
	_ = v391
	var v393 int32
	_ = v393
	var v394 int32
	_ = v394
	var v395 int32
	_ = v395
	var v396 int32
	_ = v396
	var v401 int32
	_ = v401
	var v402 int32
	_ = v402
	var v405 int32
	_ = v405
	var v406 int32
	_ = v406
	var v409 int32
	_ = v409
	var v410 int32
	_ = v410
	var v411 int32
	_ = v411
	var v413 int32
	_ = v413
	var v415 int32
	_ = v415
	var v419 int32
	_ = v419
	var v422 int32
	_ = v422
	var v423 int32
	_ = v423
	var v426 int64
	_ = v426
	var v431 int32
	_ = v431
	var v434 int32
	_ = v434
	var v436 int32
	_ = v436
	var v437 int32
	_ = v437
	var v438 int32
	_ = v438
	var v440 int32
	_ = v440
	var v444 int32
	_ = v444
	var v446 int32
	_ = v446
	var v447 int32
	_ = v447
	var v448 int32
	_ = v448
	var v450 int32
	_ = v450
	var v451 int32
	_ = v451
	var v452 int32
	_ = v452
	var v453 int32
	_ = v453
	var v458 int32
	_ = v458
	var v459 int32
	_ = v459
	var v462 int32
	_ = v462
	var v463 int32
	_ = v463
	var v466 int32
	_ = v466
	var v467 int32
	_ = v467
	var v468 int32
	_ = v468
	var v470 int32
	_ = v470
	var v472 int32
	_ = v472
	var v476 int32
	_ = v476
	var v479 int32
	_ = v479
	var v480 int32
	_ = v480
	var v483 int64
	_ = v483
	var v488 int32
	_ = v488
	var v491 int32
	_ = v491
	var v493 int32
	_ = v493
	var v494 int32
	_ = v494
	var v495 int32
	_ = v495
	var v497 int32
	_ = v497
	var v501 int32
	_ = v501
	var v503 int32
	_ = v503
	var v504 int32
	_ = v504
	var v505 int32
	_ = v505
	var v507 int32
	_ = v507
	var v508 int32
	_ = v508
	var v509 int32
	_ = v509
	var v510 int32
	_ = v510
	var v515 int32
	_ = v515
	var v516 int32
	_ = v516
	var v519 int32
	_ = v519
	var v520 int32
	_ = v520
	var v523 int32
	_ = v523
	var v524 int32
	_ = v524
	var v525 int32
	_ = v525
	var v527 int32
	_ = v527
	var v529 int32
	_ = v529
	var v533 int32
	_ = v533
	var v536 int32
	_ = v536
	var v537 int32
	_ = v537
	var v540 int64
	_ = v540
	var v545 int32
	_ = v545
	var v548 int32
	_ = v548
	var v550 int32
	_ = v550
	var v551 int32
	_ = v551
	var v552 int32
	_ = v552
	var v554 int32
	_ = v554
	var v558 int32
	_ = v558
	var v560 int32
	_ = v560
	var v561 int32
	_ = v561
	var v562 int32
	_ = v562
	var v564 int32
	_ = v564
	var v565 int32
	_ = v565
	var v566 int32
	_ = v566
	var v567 int32
	_ = v567
	var v572 int32
	_ = v572
	var v573 int32
	_ = v573
	var v576 int32
	_ = v576
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
	var v586 int32
	_ = v586
	var v590 int32
	_ = v590
	var v593 int32
	_ = v593
	var v594 int32
	_ = v594
	var v597 int64
	_ = v597
	var v602 int32
	_ = v602
	var v605 int32
	_ = v605
	var v607 int32
	_ = v607
	var v608 int32
	_ = v608
	var v609 int32
	_ = v609
	var v611 int32
	_ = v611
	var v615 int32
	_ = v615
	var v617 int32
	_ = v617
	var v618 int32
	_ = v618
	var v619 int32
	_ = v619
	var v621 int32
	_ = v621
	var v622 int32
	_ = v622
	var v623 int32
	_ = v623
	var v624 int32
	_ = v624
	var v629 int32
	_ = v629
	var v630 int32
	_ = v630
	var v633 int32
	_ = v633
	var v634 int32
	_ = v634
	var v637 int32
	_ = v637
	var v638 int32
	_ = v638
	var v639 int32
	_ = v639
	var v641 int32
	_ = v641
	var v643 int32
	_ = v643
	var v647 int32
	_ = v647
	var v650 int32
	_ = v650
	var v651 int32
	_ = v651
	var v654 int64
	_ = v654
	var v659 int32
	_ = v659
	var v662 int32
	_ = v662
	var v664 int32
	_ = v664
	var v665 int32
	_ = v665
	var v666 int32
	_ = v666
	var v668 int32
	_ = v668
	var v672 int32
	_ = v672
	var v674 int32
	_ = v674
	var v675 int32
	_ = v675
	var v676 int32
	_ = v676
	var v678 int32
	_ = v678
	var v679 int32
	_ = v679
	var v680 int32
	_ = v680
	var v681 int32
	_ = v681
	var v686 int32
	_ = v686
	var v687 int32
	_ = v687
	var v690 int32
	_ = v690
	var v691 int32
	_ = v691
	var v694 int32
	_ = v694
	var v695 int32
	_ = v695
	var v696 int32
	_ = v696
	var v698 int32
	_ = v698
	var v700 int32
	_ = v700
	var v704 int32
	_ = v704
	var v707 int32
	_ = v707
	var v708 int32
	_ = v708
	var v711 int64
	_ = v711
	var v716 int32
	_ = v716
	var v719 int32
	_ = v719
	var v721 int32
	_ = v721
	var v722 int32
	_ = v722
	var v723 int32
	_ = v723
	var v725 int32
	_ = v725
	var v729 int32
	_ = v729
	var v731 int32
	_ = v731
	var v732 int32
	_ = v732
	var v733 int32
	_ = v733
	var v735 int32
	_ = v735
	var v736 int32
	_ = v736
	var v737 int32
	_ = v737
	var v738 int32
	_ = v738
	var v743 int32
	_ = v743
	var v744 int32
	_ = v744
	var v747 int32
	_ = v747
	var v748 int32
	_ = v748
	var v751 int32
	_ = v751
	var v752 int32
	_ = v752
	var v753 int32
	_ = v753
	var v755 int32
	_ = v755
	var v757 int32
	_ = v757
	var v761 int32
	_ = v761
	var v764 int32
	_ = v764
	var v765 int32
	_ = v765
	var v768 int64
	_ = v768
	var v773 int32
	_ = v773
	var v777 int32
	_ = v777
	var v778 int32
	_ = v778
	var v788 int32
	_ = v788
	var v790 int32
	_ = v790
	var v791 int32
	_ = v791
	var v794 int32
	_ = v794
	var v798 int32
	_ = v798
	var v799 int32
	_ = v799
	var v804 int32
	_ = v804
	var v810 int32
	_ = v810
	v8 = m.G0
	v10 = v8 - int32(32)
	m.G0 = v10
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v12 < int32(3) {
		goto L4
	} else {
		goto L5
	}
L1:
	;
	m.G0 = v10 + int32(32)
	return
L2:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = v28
	F_addReplyErrorFormat(m, l0, int32(_a1511), v10+int32(16))
	mBase = m.M
	v810 = m.ExcPending
	if v810 != 0 {
		goto L24
	} else {
		goto L246
	}
L3:
	;
	v794 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v798 = *(*int32)(unsafe.Add(mBase, uint32(v794+v791<<(uint(int32(2))%32))))
	v799 = F_objectGetVal(m, v798)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v10)+4)) = v28
	*(*int32)(unsafe.Add(mBase, uint32(v10))) = v799
	F_addReplyErrorFormat(m, l0, int32(_a1512), v10)
	mBase = m.M
	v804 = m.ExcPending
	if v804 != 0 {
		goto L24
	} else {
		goto L245
	}
L4:
	;
	v788 = *(*int32)(unsafe.Add(mBase, _consts[84]))
	F_addReply(m, l0, v788)
	mBase = m.M
	v790 = m.ExcPending
	if v790 != 0 {
		goto L24
	} else {
		goto L244
	}
L5:
	;
	v18 = v12
	v19 = int32(2)
	goto L6
L6:
	;
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v23+v19<<(uint(int32(2))%32))))
	v28 = F_objectGetVal(m, v27)
	mBase = m.M
	v29 = int32(_a1513)
	v32 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28))))
	if v32 != 0 {
		goto L10
	} else {
		goto L11
	}
L7:
	;
	goto L4
L8:
	;
	v70 = base.B2i32(int32(-2) < v19-v18)
	if int32(-2) < v19-v18 {
		goto L21
	} else {
		goto L22
	}
L9:
	;
	v64 = F_tolower(m, v60)
	mBase = m.M
	v65 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v61))))
	v66 = F_tolower(m, v65)
	mBase = m.M
	goto L8
L10:
	;
	v34 = v28
	v35 = v29
	v36 = v32
	goto L13
L11:
	;
	v60 = int32(0)
	v61 = v29
	goto L9
L12:
	;
	v60 = v57 & int32(255)
	v61 = v56
	goto L9
L13:
	;
	v38 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v35))))
	if v38 == int32(0) {
		v56 = v35
		v57 = v36
		goto L12
	} else {
		goto L15
	}
L14:
	;
	v56 = v50
	v57 = int32(0)
	goto L12
L15:
	;
	v42 = v36 & int32(255)
	if v42 == v38 {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	v49 = int32(1)
	v50 = v35 + v49
	v51 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v34)+1)))
	if v51 != 0 {
		v34 = v34 + v49
		v35 = v50
		v36 = v51
		goto L13
	} else {
		goto L19
	}
L17:
	;
	v44 = F_tolower(m, v42)
	mBase = m.M
	v45 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v35))))
	v46 = F_tolower(m, v45)
	mBase = m.M
	if v44 == v46 {
		goto L16
	} else {
		goto L18
	}
L18:
	;
	v48 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v34))))
	v56 = v35
	v57 = v48
	goto L12
L19:
	;
	goto L14
L20:
	;
	v777 = v773 + int32(1)
	v778 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v777 < v778 {
		v18 = v778
		v19 = v777
		goto L6
	} else {
		goto L243
	}
L21:
	;
	v89 = int32(_a1514)
	v92 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28))))
	if v92 != 0 {
		goto L30
	} else {
		goto L31
	}
L22:
	;
	if v64-v66 != 0 {
		goto L21
	} else {
		goto L23
	}
L23:
	;
	v71 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v73 = v19 + int32(1)
	v77 = *(*int32)(unsafe.Add(mBase, uint32(v71+v73<<(uint(int32(2))%32))))
	v80 = F_getLongLongFromObject(m, v77, v10+int32(24))
	mBase = m.M
	v81 = m.ExcPending
	if v81 != 0 {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	return
L25:
	;
	if v80 == int32(-1) {
		v791 = v73
		goto L3
	} else {
		goto L26
	}
L26:
	;
	v84 = *(*int64)(unsafe.Add(mBase, uint32(v10)+24))
	if v84 < int64(1) {
		v791 = v73
		goto L3
	} else {
		goto L27
	}
L27:
	;
	*(*int64)(unsafe.Add(mBase, _consts[767])) = v84
	v773 = v73
	goto L20
L28:
	;
	if int32(-2) < v19-v18 {
		goto L40
	} else {
		goto L41
	}
L29:
	;
	v124 = F_tolower(m, v120)
	mBase = m.M
	v125 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v121))))
	v126 = F_tolower(m, v125)
	mBase = m.M
	goto L28
L30:
	;
	v94 = v28
	v95 = v89
	v96 = v92
	goto L33
L31:
	;
	v120 = int32(0)
	v121 = v89
	goto L29
L32:
	;
	v120 = v117 & int32(255)
	v121 = v116
	goto L29
L33:
	;
	v98 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v95))))
	if v98 == int32(0) {
		v116 = v95
		v117 = v96
		goto L32
	} else {
		goto L35
	}
L34:
	;
	v116 = v110
	v117 = int32(0)
	goto L32
L35:
	;
	v102 = v96 & int32(255)
	if v102 == v98 {
		goto L36
	} else {
		goto L37
	}
L36:
	;
	v109 = int32(1)
	v110 = v95 + v109
	v111 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v94)+1)))
	if v111 != 0 {
		v94 = v94 + v109
		v95 = v110
		v96 = v111
		goto L33
	} else {
		goto L39
	}
L37:
	;
	v104 = F_tolower(m, v102)
	mBase = m.M
	v105 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v95))))
	v106 = F_tolower(m, v105)
	mBase = m.M
	if v104 == v106 {
		goto L36
	} else {
		goto L38
	}
L38:
	;
	v108 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v94))))
	v116 = v95
	v117 = v108
	goto L32
L39:
	;
	goto L34
L40:
	;
	v146 = int32(_a1515)
	v149 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28))))
	if v149 != 0 {
		goto L48
	} else {
		goto L49
	}
L41:
	;
	if v124-v126 != 0 {
		goto L40
	} else {
		goto L42
	}
L42:
	;
	v128 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v130 = v19 + int32(1)
	v134 = *(*int32)(unsafe.Add(mBase, uint32(v128+v130<<(uint(int32(2))%32))))
	v137 = F_getLongLongFromObject(m, v134, v10+int32(24))
	mBase = m.M
	v138 = m.ExcPending
	if v138 != 0 {
		goto L24
	} else {
		goto L43
	}
L43:
	;
	if v137 == int32(-1) {
		v791 = v130
		goto L3
	} else {
		goto L44
	}
L44:
	;
	v141 = *(*int64)(unsafe.Add(mBase, uint32(v10)+24))
	if v141 < int64(1) {
		v791 = v130
		goto L3
	} else {
		goto L45
	}
L45:
	;
	*(*int64)(unsafe.Add(mBase, _consts[771])) = v141
	v773 = v130
	goto L20
L46:
	;
	if int32(-2) < v19-v18 {
		goto L58
	} else {
		goto L59
	}
L47:
	;
	v181 = F_tolower(m, v177)
	mBase = m.M
	v182 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v178))))
	v183 = F_tolower(m, v182)
	mBase = m.M
	goto L46
L48:
	;
	v151 = v28
	v152 = v146
	v153 = v149
	goto L51
L49:
	;
	v177 = int32(0)
	v178 = v146
	goto L47
L50:
	;
	v177 = v174 & int32(255)
	v178 = v173
	goto L47
L51:
	;
	v155 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v152))))
	if v155 == int32(0) {
		v173 = v152
		v174 = v153
		goto L50
	} else {
		goto L53
	}
L52:
	;
	v173 = v167
	v174 = int32(0)
	goto L50
L53:
	;
	v159 = v153 & int32(255)
	if v159 == v155 {
		goto L54
	} else {
		goto L55
	}
L54:
	;
	v166 = int32(1)
	v167 = v152 + v166
	v168 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v151)+1)))
	if v168 != 0 {
		v151 = v151 + v166
		v152 = v167
		v153 = v168
		goto L51
	} else {
		goto L57
	}
L55:
	;
	v161 = F_tolower(m, v159)
	mBase = m.M
	v162 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v152))))
	v163 = F_tolower(m, v162)
	mBase = m.M
	if v161 == v163 {
		goto L54
	} else {
		goto L56
	}
L56:
	;
	v165 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v151))))
	v173 = v152
	v174 = v165
	goto L50
L57:
	;
	goto L52
L58:
	;
	v203 = int32(_a1516)
	v206 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28))))
	if v206 != 0 {
		goto L66
	} else {
		goto L67
	}
L59:
	;
	if v181-v183 != 0 {
		goto L58
	} else {
		goto L60
	}
L60:
	;
	v185 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v187 = v19 + int32(1)
	v191 = *(*int32)(unsafe.Add(mBase, uint32(v185+v187<<(uint(int32(2))%32))))
	v194 = F_getLongLongFromObject(m, v191, v10+int32(24))
	mBase = m.M
	v195 = m.ExcPending
	if v195 != 0 {
		goto L24
	} else {
		goto L61
	}
L61:
	;
	if v194 == int32(-1) {
		v791 = v187
		goto L3
	} else {
		goto L62
	}
L62:
	;
	v198 = *(*int64)(unsafe.Add(mBase, uint32(v10)+24))
	if v198 < int64(1) {
		v791 = v187
		goto L3
	} else {
		goto L63
	}
L63:
	;
	*(*int64)(unsafe.Add(mBase, _consts[772])) = v198
	v773 = v187
	goto L20
L64:
	;
	if int32(-2) < v19-v18 {
		goto L76
	} else {
		goto L77
	}
L65:
	;
	v238 = F_tolower(m, v234)
	mBase = m.M
	v239 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v235))))
	v240 = F_tolower(m, v239)
	mBase = m.M
	goto L64
L66:
	;
	v208 = v28
	v209 = v203
	v210 = v206
	goto L69
L67:
	;
	v234 = int32(0)
	v235 = v203
	goto L65
L68:
	;
	v234 = v231 & int32(255)
	v235 = v230
	goto L65
L69:
	;
	v212 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v209))))
	if v212 == int32(0) {
		v230 = v209
		v231 = v210
		goto L68
	} else {
		goto L71
	}
L70:
	;
	v230 = v224
	v231 = int32(0)
	goto L68
L71:
	;
	v216 = v210 & int32(255)
	if v216 == v212 {
		goto L72
	} else {
		goto L73
	}
L72:
	;
	v223 = int32(1)
	v224 = v209 + v223
	v225 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v208)+1)))
	if v225 != 0 {
		v208 = v208 + v223
		v209 = v224
		v210 = v225
		goto L69
	} else {
		goto L75
	}
L73:
	;
	v218 = F_tolower(m, v216)
	mBase = m.M
	v219 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v209))))
	v220 = F_tolower(m, v219)
	mBase = m.M
	if v218 == v220 {
		goto L72
	} else {
		goto L74
	}
L74:
	;
	v222 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v208))))
	v230 = v209
	v231 = v222
	goto L68
L75:
	;
	goto L70
L76:
	;
	v260 = int32(_a1517)
	v263 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28))))
	if v263 != 0 {
		goto L84
	} else {
		goto L85
	}
L77:
	;
	if v238-v240 != 0 {
		goto L76
	} else {
		goto L78
	}
L78:
	;
	v242 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v244 = v19 + int32(1)
	v248 = *(*int32)(unsafe.Add(mBase, uint32(v242+v244<<(uint(int32(2))%32))))
	v251 = F_getLongLongFromObject(m, v248, v10+int32(24))
	mBase = m.M
	v252 = m.ExcPending
	if v252 != 0 {
		goto L24
	} else {
		goto L79
	}
L79:
	;
	if v251 == int32(-1) {
		v791 = v244
		goto L3
	} else {
		goto L80
	}
L80:
	;
	v255 = *(*int64)(unsafe.Add(mBase, uint32(v10)+24))
	if v255 < int64(1) {
		v791 = v244
		goto L3
	} else {
		goto L81
	}
L81:
	;
	*(*int64)(unsafe.Add(mBase, _consts[766])) = v255
	v773 = v244
	goto L20
L82:
	;
	if int32(-2) < v19-v18 {
		goto L94
	} else {
		goto L95
	}
L83:
	;
	v295 = F_tolower(m, v291)
	mBase = m.M
	v296 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v292))))
	v297 = F_tolower(m, v296)
	mBase = m.M
	goto L82
L84:
	;
	v265 = v28
	v266 = v260
	v267 = v263
	goto L87
L85:
	;
	v291 = int32(0)
	v292 = v260
	goto L83
L86:
	;
	v291 = v288 & int32(255)
	v292 = v287
	goto L83
L87:
	;
	v269 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v266))))
	if v269 == int32(0) {
		v287 = v266
		v288 = v267
		goto L86
	} else {
		goto L89
	}
L88:
	;
	v287 = v281
	v288 = int32(0)
	goto L86
L89:
	;
	v273 = v267 & int32(255)
	if v273 == v269 {
		goto L90
	} else {
		goto L91
	}
L90:
	;
	v280 = int32(1)
	v281 = v266 + v280
	v282 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v265)+1)))
	if v282 != 0 {
		v265 = v265 + v280
		v266 = v281
		v267 = v282
		goto L87
	} else {
		goto L93
	}
L91:
	;
	v275 = F_tolower(m, v273)
	mBase = m.M
	v276 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v266))))
	v277 = F_tolower(m, v276)
	mBase = m.M
	if v275 == v277 {
		goto L90
	} else {
		goto L92
	}
L92:
	;
	v279 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v265))))
	v287 = v266
	v288 = v279
	goto L86
L93:
	;
	goto L88
L94:
	;
	v317 = int32(_a1518)
	v320 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28))))
	if v320 != 0 {
		goto L102
	} else {
		goto L103
	}
L95:
	;
	if v295-v297 != 0 {
		goto L94
	} else {
		goto L96
	}
L96:
	;
	v299 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v301 = v19 + int32(1)
	v305 = *(*int32)(unsafe.Add(mBase, uint32(v299+v301<<(uint(int32(2))%32))))
	v308 = F_getLongLongFromObject(m, v305, v10+int32(24))
	mBase = m.M
	v309 = m.ExcPending
	if v309 != 0 {
		goto L24
	} else {
		goto L97
	}
L97:
	;
	if v308 == int32(-1) {
		v791 = v301
		goto L3
	} else {
		goto L98
	}
L98:
	;
	v312 = *(*int64)(unsafe.Add(mBase, uint32(v10)+24))
	if v312 < int64(1) {
		v791 = v301
		goto L3
	} else {
		goto L99
	}
L99:
	;
	*(*int64)(unsafe.Add(mBase, _consts[747])) = v312
	v773 = v301
	goto L20
L100:
	;
	if int32(-2) < v19-v18 {
		goto L112
	} else {
		goto L113
	}
L101:
	;
	v352 = F_tolower(m, v348)
	mBase = m.M
	v353 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v349))))
	v354 = F_tolower(m, v353)
	mBase = m.M
	goto L100
L102:
	;
	v322 = v28
	v323 = v317
	v324 = v320
	goto L105
L103:
	;
	v348 = int32(0)
	v349 = v317
	goto L101
L104:
	;
	v348 = v345 & int32(255)
	v349 = v344
	goto L101
L105:
	;
	v326 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v323))))
	if v326 == int32(0) {
		v344 = v323
		v345 = v324
		goto L104
	} else {
		goto L107
	}
L106:
	;
	v344 = v338
	v345 = int32(0)
	goto L104
L107:
	;
	v330 = v324 & int32(255)
	if v330 == v326 {
		goto L108
	} else {
		goto L109
	}
L108:
	;
	v337 = int32(1)
	v338 = v323 + v337
	v339 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v322)+1)))
	if v339 != 0 {
		v322 = v322 + v337
		v323 = v338
		v324 = v339
		goto L105
	} else {
		goto L111
	}
L109:
	;
	v332 = F_tolower(m, v330)
	mBase = m.M
	v333 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v323))))
	v334 = F_tolower(m, v333)
	mBase = m.M
	if v332 == v334 {
		goto L108
	} else {
		goto L110
	}
L110:
	;
	v336 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v322))))
	v344 = v323
	v345 = v336
	goto L104
L111:
	;
	goto L106
L112:
	;
	v374 = int32(_a1519)
	v377 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28))))
	if v377 != 0 {
		goto L120
	} else {
		goto L121
	}
L113:
	;
	if v352-v354 != 0 {
		goto L112
	} else {
		goto L114
	}
L114:
	;
	v356 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v358 = v19 + int32(1)
	v362 = *(*int32)(unsafe.Add(mBase, uint32(v356+v358<<(uint(int32(2))%32))))
	v365 = F_getLongLongFromObject(m, v362, v10+int32(24))
	mBase = m.M
	v366 = m.ExcPending
	if v366 != 0 {
		goto L24
	} else {
		goto L115
	}
L115:
	;
	if v365 == int32(-1) {
		v791 = v358
		goto L3
	} else {
		goto L116
	}
L116:
	;
	v369 = *(*int64)(unsafe.Add(mBase, uint32(v10)+24))
	if v369 < int64(1) {
		v791 = v358
		goto L3
	} else {
		goto L117
	}
L117:
	;
	*(*int64)(unsafe.Add(mBase, _consts[773])) = v369
	v773 = v358
	goto L20
L118:
	;
	if int32(-2) < v19-v18 {
		goto L130
	} else {
		goto L131
	}
L119:
	;
	v409 = F_tolower(m, v405)
	mBase = m.M
	v410 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v406))))
	v411 = F_tolower(m, v410)
	mBase = m.M
	goto L118
L120:
	;
	v379 = v28
	v380 = v374
	v381 = v377
	goto L123
L121:
	;
	v405 = int32(0)
	v406 = v374
	goto L119
L122:
	;
	v405 = v402 & int32(255)
	v406 = v401
	goto L119
L123:
	;
	v383 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v380))))
	if v383 == int32(0) {
		v401 = v380
		v402 = v381
		goto L122
	} else {
		goto L125
	}
L124:
	;
	v401 = v395
	v402 = int32(0)
	goto L122
L125:
	;
	v387 = v381 & int32(255)
	if v387 == v383 {
		goto L126
	} else {
		goto L127
	}
L126:
	;
	v394 = int32(1)
	v395 = v380 + v394
	v396 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v379)+1)))
	if v396 != 0 {
		v379 = v379 + v394
		v380 = v395
		v381 = v396
		goto L123
	} else {
		goto L129
	}
L127:
	;
	v389 = F_tolower(m, v387)
	mBase = m.M
	v390 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v380))))
	v391 = F_tolower(m, v390)
	mBase = m.M
	if v389 == v391 {
		goto L126
	} else {
		goto L128
	}
L128:
	;
	v393 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v379))))
	v401 = v380
	v402 = v393
	goto L122
L129:
	;
	goto L124
L130:
	;
	v431 = int32(_a1520)
	v434 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28))))
	if v434 != 0 {
		goto L138
	} else {
		goto L139
	}
L131:
	;
	if v409-v411 != 0 {
		goto L130
	} else {
		goto L132
	}
L132:
	;
	v413 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v415 = v19 + int32(1)
	v419 = *(*int32)(unsafe.Add(mBase, uint32(v413+v415<<(uint(int32(2))%32))))
	v422 = F_getLongLongFromObject(m, v419, v10+int32(24))
	mBase = m.M
	v423 = m.ExcPending
	if v423 != 0 {
		goto L24
	} else {
		goto L133
	}
L133:
	;
	if v422 == int32(-1) {
		v791 = v415
		goto L3
	} else {
		goto L134
	}
L134:
	;
	v426 = *(*int64)(unsafe.Add(mBase, uint32(v10)+24))
	if v426 < int64(1) {
		v791 = v415
		goto L3
	} else {
		goto L135
	}
L135:
	;
	*(*int64)(unsafe.Add(mBase, _consts[774])) = v426
	v773 = v415
	goto L20
L136:
	;
	if int32(-2) < v19-v18 {
		goto L148
	} else {
		goto L149
	}
L137:
	;
	v466 = F_tolower(m, v462)
	mBase = m.M
	v467 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v463))))
	v468 = F_tolower(m, v467)
	mBase = m.M
	goto L136
L138:
	;
	v436 = v28
	v437 = v431
	v438 = v434
	goto L141
L139:
	;
	v462 = int32(0)
	v463 = v431
	goto L137
L140:
	;
	v462 = v459 & int32(255)
	v463 = v458
	goto L137
L141:
	;
	v440 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v437))))
	if v440 == int32(0) {
		v458 = v437
		v459 = v438
		goto L140
	} else {
		goto L143
	}
L142:
	;
	v458 = v452
	v459 = int32(0)
	goto L140
L143:
	;
	v444 = v438 & int32(255)
	if v444 == v440 {
		goto L144
	} else {
		goto L145
	}
L144:
	;
	v451 = int32(1)
	v452 = v437 + v451
	v453 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v436)+1)))
	if v453 != 0 {
		v436 = v436 + v451
		v437 = v452
		v438 = v453
		goto L141
	} else {
		goto L147
	}
L145:
	;
	v446 = F_tolower(m, v444)
	mBase = m.M
	v447 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v437))))
	v448 = F_tolower(m, v447)
	mBase = m.M
	if v446 == v448 {
		goto L144
	} else {
		goto L146
	}
L146:
	;
	v450 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v436))))
	v458 = v437
	v459 = v450
	goto L140
L147:
	;
	goto L142
L148:
	;
	v488 = int32(_a1521)
	v491 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28))))
	if v491 != 0 {
		goto L156
	} else {
		goto L157
	}
L149:
	;
	if v466-v468 != 0 {
		goto L148
	} else {
		goto L150
	}
L150:
	;
	v470 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v472 = v19 + int32(1)
	v476 = *(*int32)(unsafe.Add(mBase, uint32(v470+v472<<(uint(int32(2))%32))))
	v479 = F_getLongLongFromObject(m, v476, v10+int32(24))
	mBase = m.M
	v480 = m.ExcPending
	if v480 != 0 {
		goto L24
	} else {
		goto L151
	}
L151:
	;
	if v479 == int32(-1) {
		v791 = v472
		goto L3
	} else {
		goto L152
	}
L152:
	;
	v483 = *(*int64)(unsafe.Add(mBase, uint32(v10)+24))
	if v483 < int64(1) {
		v791 = v472
		goto L3
	} else {
		goto L153
	}
L153:
	;
	*(*int64)(unsafe.Add(mBase, _consts[775])) = v483
	v773 = v472
	goto L20
L154:
	;
	if int32(-2) < v19-v18 {
		goto L166
	} else {
		goto L167
	}
L155:
	;
	v523 = F_tolower(m, v519)
	mBase = m.M
	v524 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v520))))
	v525 = F_tolower(m, v524)
	mBase = m.M
	goto L154
L156:
	;
	v493 = v28
	v494 = v488
	v495 = v491
	goto L159
L157:
	;
	v519 = int32(0)
	v520 = v488
	goto L155
L158:
	;
	v519 = v516 & int32(255)
	v520 = v515
	goto L155
L159:
	;
	v497 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v494))))
	if v497 == int32(0) {
		v515 = v494
		v516 = v495
		goto L158
	} else {
		goto L161
	}
L160:
	;
	v515 = v509
	v516 = int32(0)
	goto L158
L161:
	;
	v501 = v495 & int32(255)
	if v501 == v497 {
		goto L162
	} else {
		goto L163
	}
L162:
	;
	v508 = int32(1)
	v509 = v494 + v508
	v510 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v493)+1)))
	if v510 != 0 {
		v493 = v493 + v508
		v494 = v509
		v495 = v510
		goto L159
	} else {
		goto L165
	}
L163:
	;
	v503 = F_tolower(m, v501)
	mBase = m.M
	v504 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v494))))
	v505 = F_tolower(m, v504)
	mBase = m.M
	if v503 == v505 {
		goto L162
	} else {
		goto L164
	}
L164:
	;
	v507 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v493))))
	v515 = v494
	v516 = v507
	goto L158
L165:
	;
	goto L160
L166:
	;
	v545 = int32(_a1522)
	v548 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28))))
	if v548 != 0 {
		goto L174
	} else {
		goto L175
	}
L167:
	;
	if v523-v525 != 0 {
		goto L166
	} else {
		goto L168
	}
L168:
	;
	v527 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v529 = v19 + int32(1)
	v533 = *(*int32)(unsafe.Add(mBase, uint32(v527+v529<<(uint(int32(2))%32))))
	v536 = F_getLongLongFromObject(m, v533, v10+int32(24))
	mBase = m.M
	v537 = m.ExcPending
	if v537 != 0 {
		goto L24
	} else {
		goto L169
	}
L169:
	;
	if v536 == int32(-1) {
		v791 = v529
		goto L3
	} else {
		goto L170
	}
L170:
	;
	v540 = *(*int64)(unsafe.Add(mBase, uint32(v10)+24))
	if v540 < int64(1) {
		v791 = v529
		goto L3
	} else {
		goto L171
	}
L171:
	;
	*(*int64)(unsafe.Add(mBase, _consts[776])) = v540
	v773 = v529
	goto L20
L172:
	;
	if int32(-2) < v19-v18 {
		goto L184
	} else {
		goto L185
	}
L173:
	;
	v580 = F_tolower(m, v576)
	mBase = m.M
	v581 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v577))))
	v582 = F_tolower(m, v581)
	mBase = m.M
	goto L172
L174:
	;
	v550 = v28
	v551 = v545
	v552 = v548
	goto L177
L175:
	;
	v576 = int32(0)
	v577 = v545
	goto L173
L176:
	;
	v576 = v573 & int32(255)
	v577 = v572
	goto L173
L177:
	;
	v554 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v551))))
	if v554 == int32(0) {
		v572 = v551
		v573 = v552
		goto L176
	} else {
		goto L179
	}
L178:
	;
	v572 = v566
	v573 = int32(0)
	goto L176
L179:
	;
	v558 = v552 & int32(255)
	if v558 == v554 {
		goto L180
	} else {
		goto L181
	}
L180:
	;
	v565 = int32(1)
	v566 = v551 + v565
	v567 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v550)+1)))
	if v567 != 0 {
		v550 = v550 + v565
		v551 = v566
		v552 = v567
		goto L177
	} else {
		goto L183
	}
L181:
	;
	v560 = F_tolower(m, v558)
	mBase = m.M
	v561 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v551))))
	v562 = F_tolower(m, v561)
	mBase = m.M
	if v560 == v562 {
		goto L180
	} else {
		goto L182
	}
L182:
	;
	v564 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v550))))
	v572 = v551
	v573 = v564
	goto L176
L183:
	;
	goto L178
L184:
	;
	v602 = int32(_a1523)
	v605 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28))))
	if v605 != 0 {
		goto L192
	} else {
		goto L193
	}
L185:
	;
	if v580-v582 != 0 {
		goto L184
	} else {
		goto L186
	}
L186:
	;
	v584 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v586 = v19 + int32(1)
	v590 = *(*int32)(unsafe.Add(mBase, uint32(v584+v586<<(uint(int32(2))%32))))
	v593 = F_getLongLongFromObject(m, v590, v10+int32(24))
	mBase = m.M
	v594 = m.ExcPending
	if v594 != 0 {
		goto L24
	} else {
		goto L187
	}
L187:
	;
	if v593 == int32(-1) {
		v791 = v586
		goto L3
	} else {
		goto L188
	}
L188:
	;
	v597 = *(*int64)(unsafe.Add(mBase, uint32(v10)+24))
	if v597 < int64(1) {
		v791 = v586
		goto L3
	} else {
		goto L189
	}
L189:
	;
	*(*int64)(unsafe.Add(mBase, _consts[735])) = v597
	v773 = v586
	goto L20
L190:
	;
	if int32(-2) < v19-v18 {
		goto L202
	} else {
		goto L203
	}
L191:
	;
	v637 = F_tolower(m, v633)
	mBase = m.M
	v638 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v634))))
	v639 = F_tolower(m, v638)
	mBase = m.M
	goto L190
L192:
	;
	v607 = v28
	v608 = v602
	v609 = v605
	goto L195
L193:
	;
	v633 = int32(0)
	v634 = v602
	goto L191
L194:
	;
	v633 = v630 & int32(255)
	v634 = v629
	goto L191
L195:
	;
	v611 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v608))))
	if v611 == int32(0) {
		v629 = v608
		v630 = v609
		goto L194
	} else {
		goto L197
	}
L196:
	;
	v629 = v623
	v630 = int32(0)
	goto L194
L197:
	;
	v615 = v609 & int32(255)
	if v615 == v611 {
		goto L198
	} else {
		goto L199
	}
L198:
	;
	v622 = int32(1)
	v623 = v608 + v622
	v624 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v607)+1)))
	if v624 != 0 {
		v607 = v607 + v622
		v608 = v623
		v609 = v624
		goto L195
	} else {
		goto L201
	}
L199:
	;
	v617 = F_tolower(m, v615)
	mBase = m.M
	v618 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v608))))
	v619 = F_tolower(m, v618)
	mBase = m.M
	if v617 == v619 {
		goto L198
	} else {
		goto L200
	}
L200:
	;
	v621 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v607))))
	v629 = v608
	v630 = v621
	goto L194
L201:
	;
	goto L196
L202:
	;
	v659 = int32(_a1524)
	v662 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28))))
	if v662 != 0 {
		goto L210
	} else {
		goto L211
	}
L203:
	;
	if v637-v639 != 0 {
		goto L202
	} else {
		goto L204
	}
L204:
	;
	v641 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v643 = v19 + int32(1)
	v647 = *(*int32)(unsafe.Add(mBase, uint32(v641+v643<<(uint(int32(2))%32))))
	v650 = F_getLongLongFromObject(m, v647, v10+int32(24))
	mBase = m.M
	v651 = m.ExcPending
	if v651 != 0 {
		goto L24
	} else {
		goto L205
	}
L205:
	;
	if v650 == int32(-1) {
		v791 = v643
		goto L3
	} else {
		goto L206
	}
L206:
	;
	v654 = *(*int64)(unsafe.Add(mBase, uint32(v10)+24))
	if v654 < int64(1) {
		v791 = v643
		goto L3
	} else {
		goto L207
	}
L207:
	;
	*(*int64)(unsafe.Add(mBase, _consts[777])) = v654
	v773 = v643
	goto L20
L208:
	;
	if int32(-2) < v19-v18 {
		goto L220
	} else {
		goto L221
	}
L209:
	;
	v694 = F_tolower(m, v690)
	mBase = m.M
	v695 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v691))))
	v696 = F_tolower(m, v695)
	mBase = m.M
	goto L208
L210:
	;
	v664 = v28
	v665 = v659
	v666 = v662
	goto L213
L211:
	;
	v690 = int32(0)
	v691 = v659
	goto L209
L212:
	;
	v690 = v687 & int32(255)
	v691 = v686
	goto L209
L213:
	;
	v668 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v665))))
	if v668 == int32(0) {
		v686 = v665
		v687 = v666
		goto L212
	} else {
		goto L215
	}
L214:
	;
	v686 = v680
	v687 = int32(0)
	goto L212
L215:
	;
	v672 = v666 & int32(255)
	if v672 == v668 {
		goto L216
	} else {
		goto L217
	}
L216:
	;
	v679 = int32(1)
	v680 = v665 + v679
	v681 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v664)+1)))
	if v681 != 0 {
		v664 = v664 + v679
		v665 = v680
		v666 = v681
		goto L213
	} else {
		goto L219
	}
L217:
	;
	v674 = F_tolower(m, v672)
	mBase = m.M
	v675 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v665))))
	v676 = F_tolower(m, v675)
	mBase = m.M
	if v674 == v676 {
		goto L216
	} else {
		goto L218
	}
L218:
	;
	v678 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v664))))
	v686 = v665
	v687 = v678
	goto L212
L219:
	;
	goto L214
L220:
	;
	v716 = int32(_a1525)
	v719 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28))))
	if v719 != 0 {
		goto L228
	} else {
		goto L229
	}
L221:
	;
	if v694-v696 != 0 {
		goto L220
	} else {
		goto L222
	}
L222:
	;
	v698 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v700 = v19 + int32(1)
	v704 = *(*int32)(unsafe.Add(mBase, uint32(v698+v700<<(uint(int32(2))%32))))
	v707 = F_getLongLongFromObject(m, v704, v10+int32(24))
	mBase = m.M
	v708 = m.ExcPending
	if v708 != 0 {
		goto L24
	} else {
		goto L223
	}
L223:
	;
	if v707 == int32(-1) {
		v791 = v700
		goto L3
	} else {
		goto L224
	}
L224:
	;
	v711 = *(*int64)(unsafe.Add(mBase, uint32(v10)+24))
	if v711 < int64(1) {
		v791 = v700
		goto L3
	} else {
		goto L225
	}
L225:
	;
	*(*int64)(unsafe.Add(mBase, _consts[732])) = v711
	v773 = v700
	goto L20
L226:
	;
	if int32(-2) < v19-v18 {
		goto L2
	} else {
		goto L238
	}
L227:
	;
	v751 = F_tolower(m, v747)
	mBase = m.M
	v752 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v748))))
	v753 = F_tolower(m, v752)
	mBase = m.M
	goto L226
L228:
	;
	v721 = v28
	v722 = v716
	v723 = v719
	goto L231
L229:
	;
	v747 = int32(0)
	v748 = v716
	goto L227
L230:
	;
	v747 = v744 & int32(255)
	v748 = v743
	goto L227
L231:
	;
	v725 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v722))))
	if v725 == int32(0) {
		v743 = v722
		v744 = v723
		goto L230
	} else {
		goto L233
	}
L232:
	;
	v743 = v737
	v744 = int32(0)
	goto L230
L233:
	;
	v729 = v723 & int32(255)
	if v729 == v725 {
		goto L234
	} else {
		goto L235
	}
L234:
	;
	v736 = int32(1)
	v737 = v722 + v736
	v738 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v721)+1)))
	if v738 != 0 {
		v721 = v721 + v736
		v722 = v737
		v723 = v738
		goto L231
	} else {
		goto L237
	}
L235:
	;
	v731 = F_tolower(m, v729)
	mBase = m.M
	v732 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v722))))
	v733 = F_tolower(m, v732)
	mBase = m.M
	if v731 == v733 {
		goto L234
	} else {
		goto L236
	}
L236:
	;
	v735 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v721))))
	v743 = v722
	v744 = v735
	goto L230
L237:
	;
	goto L232
L238:
	;
	if v751-v753 != 0 {
		goto L2
	} else {
		goto L239
	}
L239:
	;
	v755 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v757 = v19 + int32(1)
	v761 = *(*int32)(unsafe.Add(mBase, uint32(v755+v757<<(uint(int32(2))%32))))
	v764 = F_getLongLongFromObject(m, v761, v10+int32(24))
	mBase = m.M
	v765 = m.ExcPending
	if v765 != 0 {
		goto L24
	} else {
		goto L240
	}
L240:
	;
	if v764 == int32(-1) {
		v791 = v757
		goto L3
	} else {
		goto L241
	}
L241:
	;
	v768 = *(*int64)(unsafe.Add(mBase, uint32(v10)+24))
	if v768 < int64(1) {
		v791 = v757
		goto L3
	} else {
		goto L242
	}
L242:
	;
	*(*int64)(unsafe.Add(mBase, _consts[778])) = v768
	v773 = v757
	goto L20
L243:
	;
	goto L7
L244:
	;
	goto L1
L245:
	;
	goto L1
L246:
	;
	goto L1
}
func F_sentinelStartFailoverIfNeeded(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v16 int64
	_ = v16
	var v17 int64
	_ = v17
	var v19 int64
	_ = v19
	var v21 int64
	_ = v21
	var v23 int64
	_ = v23
	var v27 int64
	_ = v27
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v40 int64
	_ = v40
	var v43 int32
	_ = v43
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	v2 = int32(0)
	v6 = m.G0
	v8 = v6 - int32(64)
	m.G0 = v8
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v11&int32(80) != int32(16) {
		v56 = v2
		m.G0 = v8 + int32(64)
		return v56
	} else {
		v16 = F_mstime(m)
		mBase = m.M
		v17 = *(*int64)(unsafe.Add(mBase, uint32(l0)+256))
		v19 = *(*int64)(unsafe.Add(mBase, uint32(l0)+264))
		v21 = v19 << (uint(int64(1)) % 64)
		if v21 <= v16-v17 {
			F_sentinelStartFailover(m, l0)
			mBase = m.M
			v54 = m.ExcPending
			if v54 != 0 {
				return int32(0)
			} else {
				v56 = int32(1)
				m.G0 = v8 + int32(64)
				return v56
			}
		} else {
			v23 = *(*int64)(unsafe.Add(mBase, uint32(l0)+272))
			if v23 == v17 {
				v56 = v2
				m.G0 = v8 + int32(64)
				return v56
			} else {
				v27 = base.I64_div_s(v21+v17, int64(1000))
				*(*int64)(unsafe.Add(mBase, uint32(v8)+56)) = v27
				v33 = F_ctime_r(m, v8+int32(56), v8+int32(16))
				mBase = m.M
				v36 = m.ExcPending
				if v36 != 0 {
					return int32(0)
				} else {
					v37 = int32(0)
					*(*uint8)(unsafe.Add(mBase, uint32(v8)+40)) = uint8(v37)
					v40 = *(*int64)(unsafe.Add(mBase, uint32(l0)+256))
					*(*int64)(unsafe.Add(mBase, uint32(l0)+272)) = v40
					v43 = *(*int32)(unsafe.Add(mBase, _consts[6]))
					if int32(2) < v43 {
						v56 = v37
						m.G0 = v8 + int32(64)
						return v56
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v8))) = v8 + int32(16)
						F__serverLog(m, int32(2), int32(_a1546), v8)
						mBase = m.M
						v52 = m.ExcPending
						if v52 != 0 {
							return int32(0)
						} else {
							v56 = v37
							m.G0 = v8 + int32(64)
							return v56
						}
					}
				}
			}
		}
	}
}
func F_sentinelTryConnectionSharing(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v61 int64
	_ = v61
	var v62 int64
	_ = v62
	var v63 int64
	_ = v63
	var v64 int64
	_ = v64
	var v65 int64
	_ = v65
	var v66 int64
	_ = v66
	var v67 int64
	_ = v67
	var v69 int64
	_ = v69
	var v71 int64
	_ = v71
	var v73 int64
	_ = v73
	var v75 int64
	_ = v75
	var v77 int64
	_ = v77
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v99 int32
	_ = v99
	var v103 int32
	_ = v103
	var v105 int32
	_ = v105
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v119 int32
	_ = v119
	var v124 int32
	_ = v124
	var v128 int32
	_ = v128
	var v132 int32
	_ = v132
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v150 int32
	_ = v150
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
	var v161 int32
	_ = v161
	var v164 int32
	_ = v164
	var v166 int32
	_ = v166
	var v174 int32
	_ = v174
	v5 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v5&int32(4) == int32(0) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	F__serverAssert(m, int32(_a1354), int32(_a1333), int32(1098))
	mBase = m.M
	v174 = m.ExcPending
	if v174 != 0 {
		goto L6
	} else {
		goto L45
	}
L2:
	;
	v10 = int32(-1)
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v11 == int32(0) {
		v166 = v10
		goto L3
	} else {
		goto L4
	}
L3:
	;
	return v166
L4:
	;
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v15 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
	if int32(1) < v15 {
		v166 = v10
		goto L3
	} else {
		goto L5
	}
L5:
	;
	v20 = *(*int32)(unsafe.Add(mBase, _consts[733]))
	v21 = F_dictGetIterator(m, v20)
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	return int32(0)
L7:
	;
	goto L9
L8:
	;
	F_dictReleaseIterator(m, v21)
	mBase = m.M
	v164 = m.ExcPending
	if v164 != 0 {
		goto L6
	} else {
		goto L44
	}
L9:
	;
	v36 = v21 + int32(20)
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v21)+16))
	if v37 != 0 {
		goto L13
	} else {
		goto L14
	}
L10:
	;
	v150 = int32(0)
	v151 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v153 = F_releaseInstanceLink(m, v151, v150)
	mBase = m.M
	v154 = m.ExcPending
	if v154 != 0 {
		goto L6
	} else {
		goto L43
	}
L11:
	;
	if v132 == int32(0) {
		v161 = int32(-1)
		goto L8
	} else {
		goto L37
	}
L12:
	;
	v43 = v36
	v44 = v40
	goto L15
L13:
	;
	v40 = int32(1)
	goto L12
L14:
	;
	v40 = int32(0)
	goto L12
L15:
	;
	switch v44 {
	case 0:
		goto L20
	default:
		goto L19
	}
L17:
	;
	v44 = int32(0)
	goto L15
L18:
	;
	goto L11
L19:
	;
	v124 = *(*int32)(unsafe.Add(mBase, uint32(v43)))
	*(*int32)(unsafe.Add(mBase, uint32(v21)+16)) = v124
	if v124 == int32(0) {
		goto L17
	} else {
		goto L36
	}
L20:
	;
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v21)+4))
	if v48 != int32(-1) {
		v87 = v48
		goto L21
	} else {
		goto L22
	}
L21:
	;
	v88 = int32(1)
	v89 = v87 + v88
	*(*int32)(unsafe.Add(mBase, uint32(v21)+4)) = v89
	v91 = int32(0)
	v94 = *(*int32)(unsafe.Add(mBase, uint32(v21)))
	v95 = *(*int32)(unsafe.Add(mBase, uint32(v21)+8))
	v99 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v94+v95+int32(26)))))
	if v99 == int32(255) {
		goto L30
	} else {
		goto L31
	}
L22:
	;
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v21)+8))
	if v52 != 0 {
		v87 = int32(-1)
		goto L21
	} else {
		goto L23
	}
L23:
	;
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v21)))
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v21)+12))
	if v54 == int32(0) {
		goto L25
	} else {
		goto L26
	}
L24:
	;
	v81 = *(*int32)(unsafe.Add(mBase, uint32(v80)+20))
	if v81 != int32(-1) {
		goto L27
	} else {
		goto L28
	}
L25:
	;
	v61 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v53)+16)))
	v62 = int64(*(*int8)(unsafe.Add(mBase, uint32(v53)+27)))
	v63 = int64(*(*int32)(unsafe.Add(mBase, uint32(v53)+8)))
	v64 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v53)+12)))
	v65 = int64(*(*int8)(unsafe.Add(mBase, uint32(v53)+26)))
	v66 = int64(*(*int32)(unsafe.Add(mBase, uint32(v53)+4)))
	v67 = F_wangHash64(m, v66)
	mBase = m.M
	v69 = F_wangHash64(m, v65+v67)
	mBase = m.M
	v71 = F_wangHash64(m, v64+v69)
	mBase = m.M
	v73 = F_wangHash64(m, v63+v71)
	mBase = m.M
	v75 = F_wangHash64(m, v62+v73)
	mBase = m.M
	v77 = F_wangHash64(m, v61+v75)
	mBase = m.M
	*(*int64)(unsafe.Add(mBase, uint32(v21)+24)) = v77
	v79 = *(*int32)(unsafe.Add(mBase, uint32(v21)))
	v80 = v79
	goto L24
L26:
	;
	v57 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v53)+24)))
	v59 = v57 + int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v53)+24)) = uint16(v59)
	v80 = v53
	goto L24
L27:
	;
	v87 = v81 + int32(-1)
	goto L21
L28:
	;
	v84 = *(*int32)(unsafe.Add(mBase, uint32(v21)+4))
	v87 = v84
	goto L21
L29:
	;
	v114 = int32(2)
	v119 = *(*int32)(unsafe.Add(mBase, uint32(v94+v112<<(uint(v114)%32)+int32(4))))
	v43 = v119 + v113<<(uint(v114)%32)
	v44 = int32(1)
	goto L15
L30:
	;
	v103 = v91
	goto L32
L31:
	;
	v103 = v88 << (uint(v99) % 32)
	goto L32
L32:
	;
	if v89 < v103 {
		v112 = v95
		v113 = v89
		goto L29
	} else {
		goto L33
	}
L33:
	;
	if v95 != 0 {
		v132 = v91
		goto L18
	} else {
		goto L34
	}
L34:
	;
	v105 = *(*int32)(unsafe.Add(mBase, uint32(v94)+20))
	if v105 == int32(-1) {
		v132 = v91
		goto L18
	} else {
		goto L35
	}
L35:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v21)+4)) = int64(4294967296)
	v112 = int32(1)
	v113 = int32(0)
	goto L29
L36:
	;
	v128 = *(*int32)(unsafe.Add(mBase, uint32(v124)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v36))) = v128
	v132 = v124
	goto L18
L37:
	;
	v138 = *(*int32)(unsafe.Add(mBase, uint32(v132)+8))
	goto L38
L38:
	;
	v139 = *(*int32)(unsafe.Add(mBase, uint32(l0)+192))
	if v138 == v139 {
		goto L9
	} else {
		goto L39
	}
L39:
	;
	v141 = *(*int32)(unsafe.Add(mBase, uint32(v138)+144))
	v142 = int32(0)
	v144 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v145 = F_getSentinelValkeyInstanceByAddrAndRunID(m, v141, v142, v142, v144)
	mBase = m.M
	v146 = m.ExcPending
	if v146 != 0 {
		goto L6
	} else {
		goto L40
	}
L40:
	;
	if v145 == int32(0) {
		goto L9
	} else {
		goto L41
	}
L41:
	;
	if v145 == l0 {
		goto L9
	} else {
		goto L42
	}
L42:
	;
	goto L10
L43:
	;
	v155 = *(*int32)(unsafe.Add(mBase, uint32(v145)+28))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v155
	v157 = *(*int32)(unsafe.Add(mBase, uint32(v155)))
	*(*int32)(unsafe.Add(mBase, uint32(v155))) = v157 + int32(1)
	v161 = v150
	goto L8
L44:
	;
	v166 = v161
	goto L3
L45:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_sentinelValkeyInstanceLookupReplica(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
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
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
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
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v74 int32
	_ = v74
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	v11 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v11&int32(1) == int32(0) {
		F__serverAssert(m, int32(_a1355), int32(_a1333), int32(1410))
		mBase = m.M
		v74 = m.ExcPending
		if v74 != 0 {
			return int32(0)
		} else {
			F_abort(m)
			mBase = m.M
			base.Wasm_trap_unreachable()
			for {
			}
		}
	} else {
		v16 = int32(0)
		v18 = F_createSentinelAddr(m, l1, l2, v16)
		mBase = m.M
		v21 = m.ExcPending
		if v21 != 0 {
			return int32(0)
		} else {
			if v18 == int32(0) {
				v64 = v16
				m.G0 = v9 + int32(16)
				return v64
			} else {
				v24 = int32(0)
				v25 = *(*int32)(unsafe.Add(mBase, _consts[734]))
				v31 = *(*int32)(unsafe.Add(mBase, uint32(v18+base.B2i32(v25 == v24)<<(uint(int32(2))%32))))
				v32 = int32(58)
				v33 = F___strchrnul(m, v31, v32)
				mBase = m.M
				v35 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v33))))
				if v35 == v32 {
					v39 = v33
				} else {
					v39 = v24
				}
				v40 = F_sdsempty(m)
				mBase = m.M
				v41 = m.ExcPending
				if v41 != 0 {
					return int32(0)
				} else {
					v42 = *(*int32)(unsafe.Add(mBase, uint32(v18)+8))
					*(*int32)(unsafe.Add(mBase, uint32(v9)+4)) = v42
					*(*int32)(unsafe.Add(mBase, uint32(v9))) = v31
					if v39 != 0 {
						v47 = int32(_a1004)
					} else {
						v47 = int32(_a1005)
					}
					v48 = F_sdscatprintf(m, v40, v47, v9)
					mBase = m.M
					v49 = m.ExcPending
					if v49 != 0 {
						return int32(0)
					} else {
						v50 = *(*int32)(unsafe.Add(mBase, uint32(v18)))
						F_sdsfree(m, v50)
						mBase = m.M
						v52 = m.ExcPending
						if v52 != 0 {
							return int32(0)
						} else {
							v53 = *(*int32)(unsafe.Add(mBase, uint32(v18)+4))
							F_sdsfree(m, v53)
							mBase = m.M
							v55 = m.ExcPending
							if v55 != 0 {
								return int32(0)
							} else {
								F_valkey_free(m, v18)
								mBase = m.M
								v57 = m.ExcPending
								if v57 != 0 {
									return int32(0)
								} else {
									v58 = *(*int32)(unsafe.Add(mBase, uint32(l0)+148))
									v59 = F_dictFetchValue(m, v58, v48)
									mBase = m.M
									v60 = m.ExcPending
									if v60 != 0 {
										return int32(0)
									} else {
										F_sdsfree(m, v48)
										mBase = m.M
										v62 = m.ExcPending
										if v62 != 0 {
											return int32(0)
										} else {
											v64 = v59
											m.G0 = v9 + int32(16)
											return v64
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
