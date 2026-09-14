package p2

import (
	base "github.com/shibukawa/vkmem/internal/aot/vkaot/base"
	"unsafe"
)

func F_AddReplyFromClient(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v15 int32
	_ = v15
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
	var v28 int32
	_ = v28
	var v32 int32
	_ = v32
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
	var v40 int32
	_ = v40
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
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
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v92 int64
	_ = v92
	var v93 int64
	_ = v93
	var v96 int32
	_ = v96
	var v100 int32
	_ = v100
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v124 int32
	_ = v124
	var v130 int32
	_ = v130
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l1)+200))
	if v10&int32(1024) == int32(0) {
		if v10&int32(16777216) != 0 {
			F__serverAssert(m, int32(_a_F_AddReplyFromClient_0), int32(_a_F_AddReplyFromClient_1), int32(1742))
			mBase = m.M
			v130 = m.ExcPending
			if v130 != 0 {
				return
			} else {
				F_abort(m)
				mBase = m.M
				base.Wasm_trap_unreachable()
				for {
				}
			}
		} else {
			v52 = *(*int32)(unsafe.Add(mBase, uint32(l1)+180))
			v53 = *(*int32)(unsafe.Add(mBase, uint32(l1)+124))
			v54 = F_prepareClientToWrite(m, l0)
			mBase = m.M
			v55 = m.ExcPending
			if v55 != 0 {
				return
			} else {
				if v54 != 0 {
					v58 = F_prepareClientToWrite(m, l0)
					mBase = m.M
					v59 = m.ExcPending
					if v59 != 0 {
						return
					} else {
						if v58 != 0 {
							m.G0 = v8 + int32(16)
							return
						} else {
							v60 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+200)))
							if v60&int32(64) != 0 {
								m.G0 = v8 + int32(16)
								return
							} else {
								v63 = *(*int32)(unsafe.Add(mBase, uint32(l1)+132))
								v64 = *(*int32)(unsafe.Add(mBase, uint32(v63)+20))
								if v64 == int32(0) {
								} else {
									v67 = *(*int32)(unsafe.Add(mBase, uint32(l0)+132))
									v71 = *(*int32)(unsafe.Add(mBase, uint32(v63)+20))
									if v71 == int32(0) {
									} else {
										v74 = *(*int32)(unsafe.Add(mBase, uint32(v63)))
										v75 = *(*int32)(unsafe.Add(mBase, uint32(v67)+4))
										*(*int32)(unsafe.Add(mBase, uint32(v74))) = v75
										if v75 == int32(0) {
											*(*int32)(unsafe.Add(mBase, uint32(v67))) = v74
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(v75)+4)) = v74
										}
										v81 = *(*int32)(unsafe.Add(mBase, uint32(v63)+4))
										*(*int32)(unsafe.Add(mBase, uint32(v67)+4)) = v81
										v83 = *(*int32)(unsafe.Add(mBase, uint32(v67)+20))
										*(*int32)(unsafe.Add(mBase, uint32(v67)+20)) = v83 + v71
										*(*int32)(unsafe.Add(mBase, uint32(v63)+20)) = int32(0)
										*(*int64)(unsafe.Add(mBase, uint32(v63))) = int64(0)
									}
								}
								v92 = *(*int64)(unsafe.Add(mBase, uint32(l0)+160))
								v93 = *(*int64)(unsafe.Add(mBase, uint32(l1)+160))
								*(*int64)(unsafe.Add(mBase, uint32(l0)+160)) = v92 + v93
								v96 = int32(0)
								*(*int32)(unsafe.Add(mBase, uint32(l1)+180)) = v96
								*(*int64)(unsafe.Add(mBase, uint32(l1)+160)) = int64(0)
								v100 = *(*int32)(unsafe.Add(mBase, uint32(l1)+344))
								if v100 == v96 {
									v111 = F_closeClientOnOutputBufferLimitReached(m, l0, int32(1))
									mBase = m.M
									v112 = m.ExcPending
									if v112 != 0 {
										return
									} else {
										m.G0 = v8 + int32(16)
										return
									}
								} else {
									F_deferredAfterErrorReply(m, l0, v100)
									mBase = m.M
									v104 = m.ExcPending
									if v104 != 0 {
										return
									} else {
										v105 = *(*int32)(unsafe.Add(mBase, uint32(l1)+344))
										F_listRelease(m, v105)
										mBase = m.M
										v107 = m.ExcPending
										if v107 != 0 {
											return
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(l1)+344)) = int32(0)
											v111 = F_closeClientOnOutputBufferLimitReached(m, l0, int32(1))
											mBase = m.M
											v112 = m.ExcPending
											if v112 != 0 {
												return
											} else {
												m.G0 = v8 + int32(16)
												return
											}
										}
									}
								}
							}
						}
					}
				} else {
					F__addReplyToBufferOrList(m, l0, v53, v52)
					mBase = m.M
					v57 = m.ExcPending
					if v57 != 0 {
						return
					} else {
						v58 = F_prepareClientToWrite(m, l0)
						mBase = m.M
						v59 = m.ExcPending
						if v59 != 0 {
							return
						} else {
							if v58 != 0 {
								m.G0 = v8 + int32(16)
								return
							} else {
								v60 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+200)))
								if v60&int32(64) != 0 {
									m.G0 = v8 + int32(16)
									return
								} else {
									v63 = *(*int32)(unsafe.Add(mBase, uint32(l1)+132))
									v64 = *(*int32)(unsafe.Add(mBase, uint32(v63)+20))
									if v64 == int32(0) {
									} else {
										v67 = *(*int32)(unsafe.Add(mBase, uint32(l0)+132))
										v71 = *(*int32)(unsafe.Add(mBase, uint32(v63)+20))
										if v71 == int32(0) {
										} else {
											v74 = *(*int32)(unsafe.Add(mBase, uint32(v63)))
											v75 = *(*int32)(unsafe.Add(mBase, uint32(v67)+4))
											*(*int32)(unsafe.Add(mBase, uint32(v74))) = v75
											if v75 == int32(0) {
												*(*int32)(unsafe.Add(mBase, uint32(v67))) = v74
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(v75)+4)) = v74
											}
											v81 = *(*int32)(unsafe.Add(mBase, uint32(v63)+4))
											*(*int32)(unsafe.Add(mBase, uint32(v67)+4)) = v81
											v83 = *(*int32)(unsafe.Add(mBase, uint32(v67)+20))
											*(*int32)(unsafe.Add(mBase, uint32(v67)+20)) = v83 + v71
											*(*int32)(unsafe.Add(mBase, uint32(v63)+20)) = int32(0)
											*(*int64)(unsafe.Add(mBase, uint32(v63))) = int64(0)
										}
									}
									v92 = *(*int64)(unsafe.Add(mBase, uint32(l0)+160))
									v93 = *(*int64)(unsafe.Add(mBase, uint32(l1)+160))
									*(*int64)(unsafe.Add(mBase, uint32(l0)+160)) = v92 + v93
									v96 = int32(0)
									*(*int32)(unsafe.Add(mBase, uint32(l1)+180)) = v96
									*(*int64)(unsafe.Add(mBase, uint32(l1)+160)) = int64(0)
									v100 = *(*int32)(unsafe.Add(mBase, uint32(l1)+344))
									if v100 == v96 {
										v111 = F_closeClientOnOutputBufferLimitReached(m, l0, int32(1))
										mBase = m.M
										v112 = m.ExcPending
										if v112 != 0 {
											return
										} else {
											m.G0 = v8 + int32(16)
											return
										}
									} else {
										F_deferredAfterErrorReply(m, l0, v100)
										mBase = m.M
										v104 = m.ExcPending
										if v104 != 0 {
											return
										} else {
											v105 = *(*int32)(unsafe.Add(mBase, uint32(l1)+344))
											F_listRelease(m, v105)
											mBase = m.M
											v107 = m.ExcPending
											if v107 != 0 {
												return
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(l1)+344)) = int32(0)
												v111 = F_closeClientOnOutputBufferLimitReached(m, l0, int32(1))
												mBase = m.M
												v112 = m.ExcPending
												if v112 != 0 {
													return
												} else {
													m.G0 = v8 + int32(16)
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
		v15 = F_sdsempty(m)
		mBase = m.M
		v16 = m.ExcPending
		if v16 != 0 {
			return
		} else {
			v18 = *(*int32)(unsafe.Add(mBase, _c_F_AddReplyFromClient[0]))
			v19 = F_catClientInfoString(m, v15, l0, v18)
			mBase = m.M
			v20 = m.ExcPending
			if v20 != 0 {
				return
			} else {
				v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+200))
				if v21&int32(1280) != 0 {
					v40 = *(*int32)(unsafe.Add(mBase, _c_F_AddReplyFromClient[1]))
					if int32(3) < v40 {
						F_sdsfree(m, v19)
						mBase = m.M
						v49 = m.ExcPending
						if v49 != 0 {
							return
						} else {
							m.G0 = v8 + int32(16)
							return
						}
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v8))) = v19
						F__serverLog(m, int32(3), int32(_a_F_AddReplyFromClient_2), v8)
						mBase = m.M
						v47 = m.ExcPending
						if v47 != 0 {
							return
						} else {
							F_sdsfree(m, v19)
							mBase = m.M
							v49 = m.ExcPending
							if v49 != 0 {
								return
							} else {
								m.G0 = v8 + int32(16)
								return
							}
						}
					}
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(l0)+200)) = v21 | int32(1024)
					v28 = *(*int32)(unsafe.Add(mBase, _c_F_AddReplyFromClient[2]))
					if v28 == int32(0) {
						v36 = *(*int32)(unsafe.Add(mBase, _c_F_AddReplyFromClient[3]))
						v37 = F_listAddNodeTail(m, v36, l0)
						mBase = m.M
						v38 = m.ExcPending
						if v38 != 0 {
							return
						} else {
							v40 = *(*int32)(unsafe.Add(mBase, _c_F_AddReplyFromClient[1]))
							if int32(3) < v40 {
								F_sdsfree(m, v19)
								mBase = m.M
								v49 = m.ExcPending
								if v49 != 0 {
									return
								} else {
									m.G0 = v8 + int32(16)
									return
								}
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v8))) = v19
								F__serverLog(m, int32(3), int32(_a_F_AddReplyFromClient_2), v8)
								mBase = m.M
								v47 = m.ExcPending
								if v47 != 0 {
									return
								} else {
									F_sdsfree(m, v19)
									mBase = m.M
									v49 = m.ExcPending
									if v49 != 0 {
										return
									} else {
										m.G0 = v8 + int32(16)
										return
									}
								}
							}
						}
					} else {
						v32 = *(*int32)(unsafe.Add(mBase, _c_F_AddReplyFromClient[3]))
						v33 = F_listSearchKey(m, v32, l0)
						mBase = m.M
						v34 = m.ExcPending
						if v34 != 0 {
							return
						} else {
							if v33 != 0 {
								F__serverAssertWithInfo(m, l0, int32(0), int32(_a_F_AddReplyFromClient_3), int32(_a_F_AddReplyFromClient_1), int32(2277))
								mBase = m.M
								v124 = m.ExcPending
								if v124 != 0 {
									return
								} else {
									F_abort(m)
									mBase = m.M
									base.Wasm_trap_unreachable()
									for {
									}
								}
							} else {
								v36 = *(*int32)(unsafe.Add(mBase, _c_F_AddReplyFromClient[3]))
								v37 = F_listAddNodeTail(m, v36, l0)
								mBase = m.M
								v38 = m.ExcPending
								if v38 != 0 {
									return
								} else {
									v40 = *(*int32)(unsafe.Add(mBase, _c_F_AddReplyFromClient[1]))
									if int32(3) < v40 {
										F_sdsfree(m, v19)
										mBase = m.M
										v49 = m.ExcPending
										if v49 != 0 {
											return
										} else {
											m.G0 = v8 + int32(16)
											return
										}
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v8))) = v19
										F__serverLog(m, int32(3), int32(_a_F_AddReplyFromClient_2), v8)
										mBase = m.M
										v47 = m.ExcPending
										if v47 != 0 {
											return
										} else {
											F_sdsfree(m, v19)
											mBase = m.M
											v49 = m.ExcPending
											if v49 != 0 {
												return
											} else {
												m.G0 = v8 + int32(16)
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
func F__addReplyToBuffer(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v9 int32
	_ = v9
	var v14 int32
	_ = v14
	var v17 int64
	_ = v17
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v48 int32
	_ = v48
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v72 int32
	_ = v72
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v143 int32
	_ = v143
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v149 int64
	_ = v149
	var v151 int32
	_ = v151
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v171 int32
	_ = v171
	var v175 int32
	_ = v175
	var v177 int32
	_ = v177
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v194 int32
	_ = v194
	var v199 int32
	_ = v199
	var v201 int32
	_ = v201
	var v204 int32
	_ = v204
	var v206 int32
	_ = v206
	var v208 int32
	_ = v208
	var v210 int32
	_ = v210
	var v215 int32
	_ = v215
	var v217 int32
	_ = v217
	var v220 int32
	_ = v220
	var v222 int32
	_ = v222
	var v224 int32
	_ = v224
	var v231 int32
	_ = v231
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v251 int32
	_ = v251
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v256 int32
	_ = v256
	var v260 int32
	_ = v260
	v4 = int32(0)
	if l2 == v4 {
		v260 = v4
	} else {
		v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+180))
		if v9 != 0 {
		} else {
			v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+204))
			if v14&int32(268435456) != 0 {
				v120 = int32(0)
			} else {
				v17 = *(*int64)(unsafe.Add(mBase, uint32(l0)+384))
				if v17 != int64(-1) {
					v120 = int32(0)
				} else {
					if v14&int32(131072) == int32(0) {
						v48 = *(*int32)(unsafe.Add(mBase, uint32(l0)+200))
						if v48&int32(1) != 0 {
							v120 = int32(0)
						} else {
							if v48&int32(2) == int32(0) {
								if v48&int32(262144) != 0 {
									v67 = int32(_a_F__addReplyToBuffer_0)
									v68 = *(*int32)(unsafe.Add(mBase, _c_F__addReplyToBuffer[0]))
									v72 = *(*int32)(unsafe.Add(mBase, _c_F__addReplyToBuffer[1]))
									v120 = base.B2i32(v68 != int32(0)) & base.B2i32(v68 <= v72)
								} else {
									v61 = *(*int32)(unsafe.Add(mBase, uint32(l0)+216))
									if v61 == int32(0) {
										v67 = int32(_a_F__addReplyToBuffer_0)
										v68 = *(*int32)(unsafe.Add(mBase, _c_F__addReplyToBuffer[0]))
										v72 = *(*int32)(unsafe.Add(mBase, _c_F__addReplyToBuffer[1]))
										v120 = base.B2i32(v68 != int32(0)) & base.B2i32(v68 <= v72)
									} else {
										v64 = F_isImportSlotMigrationJob(m, v61)
										mBase = m.M
										v120 = int32(0)
									}
								}
							} else {
								if v48&int32(4) == int32(0) {
									v120 = int32(0)
								} else {
									if v48&int32(262144) != 0 {
										v67 = int32(_a_F__addReplyToBuffer_0)
										v68 = *(*int32)(unsafe.Add(mBase, _c_F__addReplyToBuffer[0]))
										v72 = *(*int32)(unsafe.Add(mBase, _c_F__addReplyToBuffer[1]))
										v120 = base.B2i32(v68 != int32(0)) & base.B2i32(v68 <= v72)
									} else {
										v61 = *(*int32)(unsafe.Add(mBase, uint32(l0)+216))
										if v61 == int32(0) {
											v67 = int32(_a_F__addReplyToBuffer_0)
											v68 = *(*int32)(unsafe.Add(mBase, _c_F__addReplyToBuffer[0]))
											v72 = *(*int32)(unsafe.Add(mBase, _c_F__addReplyToBuffer[1]))
											v120 = base.B2i32(v68 != int32(0)) & base.B2i32(v68 <= v72)
										} else {
											v64 = F_isImportSlotMigrationJob(m, v61)
											mBase = m.M
											v120 = int32(0)
										}
									}
								}
							}
						}
					} else {
						v25 = *(*int32)(unsafe.Add(mBase, _c_F__addReplyToBuffer[2]))
						if l0 != v25 {
							v48 = *(*int32)(unsafe.Add(mBase, uint32(l0)+200))
							if v48&int32(1) != 0 {
								v120 = int32(0)
							} else {
								if v48&int32(2) == int32(0) {
									if v48&int32(262144) != 0 {
										v67 = int32(_a_F__addReplyToBuffer_0)
										v68 = *(*int32)(unsafe.Add(mBase, _c_F__addReplyToBuffer[0]))
										v72 = *(*int32)(unsafe.Add(mBase, _c_F__addReplyToBuffer[1]))
										v120 = base.B2i32(v68 != int32(0)) & base.B2i32(v68 <= v72)
									} else {
										v61 = *(*int32)(unsafe.Add(mBase, uint32(l0)+216))
										if v61 == int32(0) {
											v67 = int32(_a_F__addReplyToBuffer_0)
											v68 = *(*int32)(unsafe.Add(mBase, _c_F__addReplyToBuffer[0]))
											v72 = *(*int32)(unsafe.Add(mBase, _c_F__addReplyToBuffer[1]))
											v120 = base.B2i32(v68 != int32(0)) & base.B2i32(v68 <= v72)
										} else {
											v64 = F_isImportSlotMigrationJob(m, v61)
											mBase = m.M
											v120 = int32(0)
										}
									}
								} else {
									if v48&int32(4) == int32(0) {
										v120 = int32(0)
									} else {
										if v48&int32(262144) != 0 {
											v67 = int32(_a_F__addReplyToBuffer_0)
											v68 = *(*int32)(unsafe.Add(mBase, _c_F__addReplyToBuffer[0]))
											v72 = *(*int32)(unsafe.Add(mBase, _c_F__addReplyToBuffer[1]))
											v120 = base.B2i32(v68 != int32(0)) & base.B2i32(v68 <= v72)
										} else {
											v61 = *(*int32)(unsafe.Add(mBase, uint32(l0)+216))
											if v61 == int32(0) {
												v67 = int32(_a_F__addReplyToBuffer_0)
												v68 = *(*int32)(unsafe.Add(mBase, _c_F__addReplyToBuffer[0]))
												v72 = *(*int32)(unsafe.Add(mBase, _c_F__addReplyToBuffer[1]))
												v120 = base.B2i32(v68 != int32(0)) & base.B2i32(v68 <= v72)
											} else {
												v64 = F_isImportSlotMigrationJob(m, v61)
												mBase = m.M
												v120 = int32(0)
											}
										}
									}
								}
							}
						} else {
							v28 = *(*int32)(unsafe.Add(mBase, _c_F__addReplyToBuffer[3]))
							if v28 == int32(0) {
								v48 = *(*int32)(unsafe.Add(mBase, uint32(l0)+200))
								if v48&int32(1) != 0 {
									v120 = int32(0)
								} else {
									if v48&int32(2) == int32(0) {
										if v48&int32(262144) != 0 {
											v67 = int32(_a_F__addReplyToBuffer_0)
											v68 = *(*int32)(unsafe.Add(mBase, _c_F__addReplyToBuffer[0]))
											v72 = *(*int32)(unsafe.Add(mBase, _c_F__addReplyToBuffer[1]))
											v120 = base.B2i32(v68 != int32(0)) & base.B2i32(v68 <= v72)
										} else {
											v61 = *(*int32)(unsafe.Add(mBase, uint32(l0)+216))
											if v61 == int32(0) {
												v67 = int32(_a_F__addReplyToBuffer_0)
												v68 = *(*int32)(unsafe.Add(mBase, _c_F__addReplyToBuffer[0]))
												v72 = *(*int32)(unsafe.Add(mBase, _c_F__addReplyToBuffer[1]))
												v120 = base.B2i32(v68 != int32(0)) & base.B2i32(v68 <= v72)
											} else {
												v64 = F_isImportSlotMigrationJob(m, v61)
												mBase = m.M
												v120 = int32(0)
											}
										}
									} else {
										if v48&int32(4) == int32(0) {
											v120 = int32(0)
										} else {
											if v48&int32(262144) != 0 {
												v67 = int32(_a_F__addReplyToBuffer_0)
												v68 = *(*int32)(unsafe.Add(mBase, _c_F__addReplyToBuffer[0]))
												v72 = *(*int32)(unsafe.Add(mBase, _c_F__addReplyToBuffer[1]))
												v120 = base.B2i32(v68 != int32(0)) & base.B2i32(v68 <= v72)
											} else {
												v61 = *(*int32)(unsafe.Add(mBase, uint32(l0)+216))
												if v61 == int32(0) {
													v67 = int32(_a_F__addReplyToBuffer_0)
													v68 = *(*int32)(unsafe.Add(mBase, _c_F__addReplyToBuffer[0]))
													v72 = *(*int32)(unsafe.Add(mBase, _c_F__addReplyToBuffer[1]))
													v120 = base.B2i32(v68 != int32(0)) & base.B2i32(v68 <= v72)
												} else {
													v64 = F_isImportSlotMigrationJob(m, v61)
													mBase = m.M
													v120 = int32(0)
												}
											}
										}
									}
								}
							} else {
								v31 = *(*int32)(unsafe.Add(mBase, uint32(v28)+68))
								if v31 == int32(0) {
									v120 = int32(0)
								} else {
									v34 = *(*int32)(unsafe.Add(mBase, uint32(v31)+48))
									if v34 == int32(288) {
										v48 = *(*int32)(unsafe.Add(mBase, uint32(l0)+200))
										if v48&int32(1) != 0 {
											v120 = int32(0)
										} else {
											if v48&int32(2) == int32(0) {
												if v48&int32(262144) != 0 {
													v67 = int32(_a_F__addReplyToBuffer_0)
													v68 = *(*int32)(unsafe.Add(mBase, _c_F__addReplyToBuffer[0]))
													v72 = *(*int32)(unsafe.Add(mBase, _c_F__addReplyToBuffer[1]))
													v120 = base.B2i32(v68 != int32(0)) & base.B2i32(v68 <= v72)
												} else {
													v61 = *(*int32)(unsafe.Add(mBase, uint32(l0)+216))
													if v61 == int32(0) {
														v67 = int32(_a_F__addReplyToBuffer_0)
														v68 = *(*int32)(unsafe.Add(mBase, _c_F__addReplyToBuffer[0]))
														v72 = *(*int32)(unsafe.Add(mBase, _c_F__addReplyToBuffer[1]))
														v120 = base.B2i32(v68 != int32(0)) & base.B2i32(v68 <= v72)
													} else {
														v64 = F_isImportSlotMigrationJob(m, v61)
														mBase = m.M
														v120 = int32(0)
													}
												}
											} else {
												if v48&int32(4) == int32(0) {
													v120 = int32(0)
												} else {
													if v48&int32(262144) != 0 {
														v67 = int32(_a_F__addReplyToBuffer_0)
														v68 = *(*int32)(unsafe.Add(mBase, _c_F__addReplyToBuffer[0]))
														v72 = *(*int32)(unsafe.Add(mBase, _c_F__addReplyToBuffer[1]))
														v120 = base.B2i32(v68 != int32(0)) & base.B2i32(v68 <= v72)
													} else {
														v61 = *(*int32)(unsafe.Add(mBase, uint32(l0)+216))
														if v61 == int32(0) {
															v67 = int32(_a_F__addReplyToBuffer_0)
															v68 = *(*int32)(unsafe.Add(mBase, _c_F__addReplyToBuffer[0]))
															v72 = *(*int32)(unsafe.Add(mBase, _c_F__addReplyToBuffer[1]))
															v120 = base.B2i32(v68 != int32(0)) & base.B2i32(v68 <= v72)
														} else {
															v64 = F_isImportSlotMigrationJob(m, v61)
															mBase = m.M
															v120 = int32(0)
														}
													}
												}
											}
										}
									} else {
										if v34 == int32(286) {
											v48 = *(*int32)(unsafe.Add(mBase, uint32(l0)+200))
											if v48&int32(1) != 0 {
												v120 = int32(0)
											} else {
												if v48&int32(2) == int32(0) {
													if v48&int32(262144) != 0 {
														v67 = int32(_a_F__addReplyToBuffer_0)
														v68 = *(*int32)(unsafe.Add(mBase, _c_F__addReplyToBuffer[0]))
														v72 = *(*int32)(unsafe.Add(mBase, _c_F__addReplyToBuffer[1]))
														v120 = base.B2i32(v68 != int32(0)) & base.B2i32(v68 <= v72)
													} else {
														v61 = *(*int32)(unsafe.Add(mBase, uint32(l0)+216))
														if v61 == int32(0) {
															v67 = int32(_a_F__addReplyToBuffer_0)
															v68 = *(*int32)(unsafe.Add(mBase, _c_F__addReplyToBuffer[0]))
															v72 = *(*int32)(unsafe.Add(mBase, _c_F__addReplyToBuffer[1]))
															v120 = base.B2i32(v68 != int32(0)) & base.B2i32(v68 <= v72)
														} else {
															v64 = F_isImportSlotMigrationJob(m, v61)
															mBase = m.M
															v120 = int32(0)
														}
													}
												} else {
													if v48&int32(4) == int32(0) {
														v120 = int32(0)
													} else {
														if v48&int32(262144) != 0 {
															v67 = int32(_a_F__addReplyToBuffer_0)
															v68 = *(*int32)(unsafe.Add(mBase, _c_F__addReplyToBuffer[0]))
															v72 = *(*int32)(unsafe.Add(mBase, _c_F__addReplyToBuffer[1]))
															v120 = base.B2i32(v68 != int32(0)) & base.B2i32(v68 <= v72)
														} else {
															v61 = *(*int32)(unsafe.Add(mBase, uint32(l0)+216))
															if v61 == int32(0) {
																v67 = int32(_a_F__addReplyToBuffer_0)
																v68 = *(*int32)(unsafe.Add(mBase, _c_F__addReplyToBuffer[0]))
																v72 = *(*int32)(unsafe.Add(mBase, _c_F__addReplyToBuffer[1]))
																v120 = base.B2i32(v68 != int32(0)) & base.B2i32(v68 <= v72)
															} else {
																v64 = F_isImportSlotMigrationJob(m, v61)
																mBase = m.M
																v120 = int32(0)
															}
														}
													}
												}
											}
										} else {
											if v34 == int32(284) {
												v48 = *(*int32)(unsafe.Add(mBase, uint32(l0)+200))
												if v48&int32(1) != 0 {
													v120 = int32(0)
												} else {
													if v48&int32(2) == int32(0) {
														if v48&int32(262144) != 0 {
															v67 = int32(_a_F__addReplyToBuffer_0)
															v68 = *(*int32)(unsafe.Add(mBase, _c_F__addReplyToBuffer[0]))
															v72 = *(*int32)(unsafe.Add(mBase, _c_F__addReplyToBuffer[1]))
															v120 = base.B2i32(v68 != int32(0)) & base.B2i32(v68 <= v72)
														} else {
															v61 = *(*int32)(unsafe.Add(mBase, uint32(l0)+216))
															if v61 == int32(0) {
																v67 = int32(_a_F__addReplyToBuffer_0)
																v68 = *(*int32)(unsafe.Add(mBase, _c_F__addReplyToBuffer[0]))
																v72 = *(*int32)(unsafe.Add(mBase, _c_F__addReplyToBuffer[1]))
																v120 = base.B2i32(v68 != int32(0)) & base.B2i32(v68 <= v72)
															} else {
																v64 = F_isImportSlotMigrationJob(m, v61)
																mBase = m.M
																v120 = int32(0)
															}
														}
													} else {
														if v48&int32(4) == int32(0) {
															v120 = int32(0)
														} else {
															if v48&int32(262144) != 0 {
																v67 = int32(_a_F__addReplyToBuffer_0)
																v68 = *(*int32)(unsafe.Add(mBase, _c_F__addReplyToBuffer[0]))
																v72 = *(*int32)(unsafe.Add(mBase, _c_F__addReplyToBuffer[1]))
																v120 = base.B2i32(v68 != int32(0)) & base.B2i32(v68 <= v72)
															} else {
																v61 = *(*int32)(unsafe.Add(mBase, uint32(l0)+216))
																if v61 == int32(0) {
																	v67 = int32(_a_F__addReplyToBuffer_0)
																	v68 = *(*int32)(unsafe.Add(mBase, _c_F__addReplyToBuffer[0]))
																	v72 = *(*int32)(unsafe.Add(mBase, _c_F__addReplyToBuffer[1]))
																	v120 = base.B2i32(v68 != int32(0)) & base.B2i32(v68 <= v72)
																} else {
																	v64 = F_isImportSlotMigrationJob(m, v61)
																	mBase = m.M
																	v120 = int32(0)
																}
															}
														}
													}
												}
											} else {
												if v34 == int32(282) {
													v48 = *(*int32)(unsafe.Add(mBase, uint32(l0)+200))
													if v48&int32(1) != 0 {
														v120 = int32(0)
													} else {
														if v48&int32(2) == int32(0) {
															if v48&int32(262144) != 0 {
																v67 = int32(_a_F__addReplyToBuffer_0)
																v68 = *(*int32)(unsafe.Add(mBase, _c_F__addReplyToBuffer[0]))
																v72 = *(*int32)(unsafe.Add(mBase, _c_F__addReplyToBuffer[1]))
																v120 = base.B2i32(v68 != int32(0)) & base.B2i32(v68 <= v72)
															} else {
																v61 = *(*int32)(unsafe.Add(mBase, uint32(l0)+216))
																if v61 == int32(0) {
																	v67 = int32(_a_F__addReplyToBuffer_0)
																	v68 = *(*int32)(unsafe.Add(mBase, _c_F__addReplyToBuffer[0]))
																	v72 = *(*int32)(unsafe.Add(mBase, _c_F__addReplyToBuffer[1]))
																	v120 = base.B2i32(v68 != int32(0)) & base.B2i32(v68 <= v72)
																} else {
																	v64 = F_isImportSlotMigrationJob(m, v61)
																	mBase = m.M
																	v120 = int32(0)
																}
															}
														} else {
															if v48&int32(4) == int32(0) {
																v120 = int32(0)
															} else {
																if v48&int32(262144) != 0 {
																	v67 = int32(_a_F__addReplyToBuffer_0)
																	v68 = *(*int32)(unsafe.Add(mBase, _c_F__addReplyToBuffer[0]))
																	v72 = *(*int32)(unsafe.Add(mBase, _c_F__addReplyToBuffer[1]))
																	v120 = base.B2i32(v68 != int32(0)) & base.B2i32(v68 <= v72)
																} else {
																	v61 = *(*int32)(unsafe.Add(mBase, uint32(l0)+216))
																	if v61 == int32(0) {
																		v67 = int32(_a_F__addReplyToBuffer_0)
																		v68 = *(*int32)(unsafe.Add(mBase, _c_F__addReplyToBuffer[0]))
																		v72 = *(*int32)(unsafe.Add(mBase, _c_F__addReplyToBuffer[1]))
																		v120 = base.B2i32(v68 != int32(0)) & base.B2i32(v68 <= v72)
																	} else {
																		v64 = F_isImportSlotMigrationJob(m, v61)
																		mBase = m.M
																		v120 = int32(0)
																	}
																}
															}
														}
													}
												} else {
													if v34 == int32(287) {
														v48 = *(*int32)(unsafe.Add(mBase, uint32(l0)+200))
														if v48&int32(1) != 0 {
															v120 = int32(0)
														} else {
															if v48&int32(2) == int32(0) {
																if v48&int32(262144) != 0 {
																	v67 = int32(_a_F__addReplyToBuffer_0)
																	v68 = *(*int32)(unsafe.Add(mBase, _c_F__addReplyToBuffer[0]))
																	v72 = *(*int32)(unsafe.Add(mBase, _c_F__addReplyToBuffer[1]))
																	v120 = base.B2i32(v68 != int32(0)) & base.B2i32(v68 <= v72)
																} else {
																	v61 = *(*int32)(unsafe.Add(mBase, uint32(l0)+216))
																	if v61 == int32(0) {
																		v67 = int32(_a_F__addReplyToBuffer_0)
																		v68 = *(*int32)(unsafe.Add(mBase, _c_F__addReplyToBuffer[0]))
																		v72 = *(*int32)(unsafe.Add(mBase, _c_F__addReplyToBuffer[1]))
																		v120 = base.B2i32(v68 != int32(0)) & base.B2i32(v68 <= v72)
																	} else {
																		v64 = F_isImportSlotMigrationJob(m, v61)
																		mBase = m.M
																		v120 = int32(0)
																	}
																}
															} else {
																if v48&int32(4) == int32(0) {
																	v120 = int32(0)
																} else {
																	if v48&int32(262144) != 0 {
																		v67 = int32(_a_F__addReplyToBuffer_0)
																		v68 = *(*int32)(unsafe.Add(mBase, _c_F__addReplyToBuffer[0]))
																		v72 = *(*int32)(unsafe.Add(mBase, _c_F__addReplyToBuffer[1]))
																		v120 = base.B2i32(v68 != int32(0)) & base.B2i32(v68 <= v72)
																	} else {
																		v61 = *(*int32)(unsafe.Add(mBase, uint32(l0)+216))
																		if v61 == int32(0) {
																			v67 = int32(_a_F__addReplyToBuffer_0)
																			v68 = *(*int32)(unsafe.Add(mBase, _c_F__addReplyToBuffer[0]))
																			v72 = *(*int32)(unsafe.Add(mBase, _c_F__addReplyToBuffer[1]))
																			v120 = base.B2i32(v68 != int32(0)) & base.B2i32(v68 <= v72)
																		} else {
																			v64 = F_isImportSlotMigrationJob(m, v61)
																			mBase = m.M
																			v120 = int32(0)
																		}
																	}
																}
															}
														}
													} else {
														if v34 != int32(289) {
															v120 = int32(0)
														} else {
															v48 = *(*int32)(unsafe.Add(mBase, uint32(l0)+200))
															if v48&int32(1) != 0 {
																v120 = int32(0)
															} else {
																if v48&int32(2) == int32(0) {
																	if v48&int32(262144) != 0 {
																		v67 = int32(_a_F__addReplyToBuffer_0)
																		v68 = *(*int32)(unsafe.Add(mBase, _c_F__addReplyToBuffer[0]))
																		v72 = *(*int32)(unsafe.Add(mBase, _c_F__addReplyToBuffer[1]))
																		v120 = base.B2i32(v68 != int32(0)) & base.B2i32(v68 <= v72)
																	} else {
																		v61 = *(*int32)(unsafe.Add(mBase, uint32(l0)+216))
																		if v61 == int32(0) {
																			v67 = int32(_a_F__addReplyToBuffer_0)
																			v68 = *(*int32)(unsafe.Add(mBase, _c_F__addReplyToBuffer[0]))
																			v72 = *(*int32)(unsafe.Add(mBase, _c_F__addReplyToBuffer[1]))
																			v120 = base.B2i32(v68 != int32(0)) & base.B2i32(v68 <= v72)
																		} else {
																			v64 = F_isImportSlotMigrationJob(m, v61)
																			mBase = m.M
																			v120 = int32(0)
																		}
																	}
																} else {
																	if v48&int32(4) == int32(0) {
																		v120 = int32(0)
																	} else {
																		if v48&int32(262144) != 0 {
																			v67 = int32(_a_F__addReplyToBuffer_0)
																			v68 = *(*int32)(unsafe.Add(mBase, _c_F__addReplyToBuffer[0]))
																			v72 = *(*int32)(unsafe.Add(mBase, _c_F__addReplyToBuffer[1]))
																			v120 = base.B2i32(v68 != int32(0)) & base.B2i32(v68 <= v72)
																		} else {
																			v61 = *(*int32)(unsafe.Add(mBase, uint32(l0)+216))
																			if v61 == int32(0) {
																				v67 = int32(_a_F__addReplyToBuffer_0)
																				v68 = *(*int32)(unsafe.Add(mBase, _c_F__addReplyToBuffer[0]))
																				v72 = *(*int32)(unsafe.Add(mBase, _c_F__addReplyToBuffer[1]))
																				v120 = base.B2i32(v68 != int32(0)) & base.B2i32(v68 <= v72)
																			} else {
																				v64 = F_isImportSlotMigrationJob(m, v61)
																				mBase = m.M
																				v120 = int32(0)
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
			v121 = *(*int32)(unsafe.Add(mBase, uint32(l0)+200))
			*(*int32)(unsafe.Add(mBase, uint32(l0)+200)) = v121&int32(-16777217) | v120<<(uint(int32(24))%32)
		}
		v130 = *(*int32)(unsafe.Add(mBase, _c_F__addReplyToBuffer[4]))
		if v130 != 0 {
			v260 = v4
		} else {
			v131 = *(*int32)(unsafe.Add(mBase, uint32(l0)+132))
			v132 = *(*int32)(unsafe.Add(mBase, uint32(v131)+20))
			if v132 != 0 {
				v260 = v4
			} else {
				v133 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
				v134 = *(*int32)(unsafe.Add(mBase, uint32(l0)+180))
				v135 = v133 - v134
				v136 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+203)))
				if v136&int32(1) != 0 {
					v141 = *(*int32)(unsafe.Add(mBase, uint32(l0)+124))
					v143 = l0 + int32(180)
					v145 = l0 + int32(184)
					v146 = int32(0)
					v147 = *(*int32)(unsafe.Add(mBase, uint32(l0)+292))
					v149 = *(*int64)(unsafe.Add(mBase, _c_F__addReplyToBuffer[5]))
					v151 = base.B2i32(v149 != int64(-1))
					if base.Ui32(v135) < base.Ui32(int32(1)) {
						v231 = v146
						v242 = v231
					} else {
						v164 = F_clusterSlotStatsEnabled(m, v147)
						mBase = m.M
						if v164 != 0 {
							v165 = v147
						} else {
							v165 = int32(-1)
						}
						if base.Ui32(v135) < base.Ui32(l2) {
							v167 = v135
						} else {
							v167 = l2
						}
						v168 = *(*int32)(unsafe.Add(mBase, uint32(v145)))
						if v168 == int32(0) {
							if base.Ui32(v135) < base.Ui32(int32(13)) {
								v231 = v146
							} else {
								v191 = *(*int32)(unsafe.Add(mBase, uint32(v143)))
								v192 = v141 + v191
								*(*int32)(unsafe.Add(mBase, uint32(v145))) = v192
								v194 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v192)+10)))
								v199 = v194&int32(254) | int32(0)
								*(*uint8)(unsafe.Add(mBase, uint32(v192)+10)) = uint8(v199)
								v201 = *(*int32)(unsafe.Add(mBase, uint32(v145)))
								*(*uint16)(unsafe.Add(mBase, uint32(v201)+8)) = uint16(v165)
								v204 = v135 + int32(-12)
								if base.Ui32(v204) < base.Ui32(l2) {
									v206 = v204
								} else {
									v206 = v167
								}
								*(*int32)(unsafe.Add(mBase, uint32(v201))) = v206
								v208 = int32(0)
								*(*int32)(unsafe.Add(mBase, uint32(v201)+4)) = v208
								v210 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v201)+10)))
								v215 = v210&int32(253) | v151<<(uint(int32(1))%32)
								*(*uint8)(unsafe.Add(mBase, uint32(v201)+10)) = uint8(v215)
								v217 = *(*int32)(unsafe.Add(mBase, uint32(v145)))
								*(*uint8)(unsafe.Add(mBase, uint32(v217)+11)) = uint8(v208)
								v220 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v217)+10)))
								v222 = v220 & int32(3)
								*(*uint8)(unsafe.Add(mBase, uint32(v217)+10)) = uint8(v222)
								v224 = *(*int32)(unsafe.Add(mBase, uint32(v143)))
								*(*int32)(unsafe.Add(mBase, uint32(v143))) = v224 + int32(12)
								v231 = v206
							}
							v242 = v231
						} else {
							v171 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v168)+10)))
							if v171&int32(1) != v146 {
								if base.Ui32(v135) < base.Ui32(int32(13)) {
									v231 = v146
								} else {
									v191 = *(*int32)(unsafe.Add(mBase, uint32(v143)))
									v192 = v141 + v191
									*(*int32)(unsafe.Add(mBase, uint32(v145))) = v192
									v194 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v192)+10)))
									v199 = v194&int32(254) | int32(0)
									*(*uint8)(unsafe.Add(mBase, uint32(v192)+10)) = uint8(v199)
									v201 = *(*int32)(unsafe.Add(mBase, uint32(v145)))
									*(*uint16)(unsafe.Add(mBase, uint32(v201)+8)) = uint16(v165)
									v204 = v135 + int32(-12)
									if base.Ui32(v204) < base.Ui32(l2) {
										v206 = v204
									} else {
										v206 = v167
									}
									*(*int32)(unsafe.Add(mBase, uint32(v201))) = v206
									v208 = int32(0)
									*(*int32)(unsafe.Add(mBase, uint32(v201)+4)) = v208
									v210 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v201)+10)))
									v215 = v210&int32(253) | v151<<(uint(int32(1))%32)
									*(*uint8)(unsafe.Add(mBase, uint32(v201)+10)) = uint8(v215)
									v217 = *(*int32)(unsafe.Add(mBase, uint32(v145)))
									*(*uint8)(unsafe.Add(mBase, uint32(v217)+11)) = uint8(v208)
									v220 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v217)+10)))
									v222 = v220 & int32(3)
									*(*uint8)(unsafe.Add(mBase, uint32(v217)+10)) = uint8(v222)
									v224 = *(*int32)(unsafe.Add(mBase, uint32(v143)))
									*(*int32)(unsafe.Add(mBase, uint32(v143))) = v224 + int32(12)
									v231 = v206
								}
								v242 = v231
							} else {
								v175 = int32(*(*int16)(unsafe.Add(mBase, uint32(v168)+8)))
								if v165 != v175 {
									if base.Ui32(v135) < base.Ui32(int32(13)) {
										v231 = v146
									} else {
										v191 = *(*int32)(unsafe.Add(mBase, uint32(v143)))
										v192 = v141 + v191
										*(*int32)(unsafe.Add(mBase, uint32(v145))) = v192
										v194 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v192)+10)))
										v199 = v194&int32(254) | int32(0)
										*(*uint8)(unsafe.Add(mBase, uint32(v192)+10)) = uint8(v199)
										v201 = *(*int32)(unsafe.Add(mBase, uint32(v145)))
										*(*uint16)(unsafe.Add(mBase, uint32(v201)+8)) = uint16(v165)
										v204 = v135 + int32(-12)
										if base.Ui32(v204) < base.Ui32(l2) {
											v206 = v204
										} else {
											v206 = v167
										}
										*(*int32)(unsafe.Add(mBase, uint32(v201))) = v206
										v208 = int32(0)
										*(*int32)(unsafe.Add(mBase, uint32(v201)+4)) = v208
										v210 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v201)+10)))
										v215 = v210&int32(253) | v151<<(uint(int32(1))%32)
										*(*uint8)(unsafe.Add(mBase, uint32(v201)+10)) = uint8(v215)
										v217 = *(*int32)(unsafe.Add(mBase, uint32(v145)))
										*(*uint8)(unsafe.Add(mBase, uint32(v217)+11)) = uint8(v208)
										v220 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v217)+10)))
										v222 = v220 & int32(3)
										*(*uint8)(unsafe.Add(mBase, uint32(v217)+10)) = uint8(v222)
										v224 = *(*int32)(unsafe.Add(mBase, uint32(v143)))
										*(*int32)(unsafe.Add(mBase, uint32(v143))) = v224 + int32(12)
										v231 = v206
									}
									v242 = v231
								} else {
									v177 = int32(1)
									if v151 != int32(base.Ui32(v171)>>(uint(v177)%32))&v177 {
										if base.Ui32(v135) < base.Ui32(int32(13)) {
											v231 = v146
										} else {
											v191 = *(*int32)(unsafe.Add(mBase, uint32(v143)))
											v192 = v141 + v191
											*(*int32)(unsafe.Add(mBase, uint32(v145))) = v192
											v194 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v192)+10)))
											v199 = v194&int32(254) | int32(0)
											*(*uint8)(unsafe.Add(mBase, uint32(v192)+10)) = uint8(v199)
											v201 = *(*int32)(unsafe.Add(mBase, uint32(v145)))
											*(*uint16)(unsafe.Add(mBase, uint32(v201)+8)) = uint16(v165)
											v204 = v135 + int32(-12)
											if base.Ui32(v204) < base.Ui32(l2) {
												v206 = v204
											} else {
												v206 = v167
											}
											*(*int32)(unsafe.Add(mBase, uint32(v201))) = v206
											v208 = int32(0)
											*(*int32)(unsafe.Add(mBase, uint32(v201)+4)) = v208
											v210 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v201)+10)))
											v215 = v210&int32(253) | v151<<(uint(int32(1))%32)
											*(*uint8)(unsafe.Add(mBase, uint32(v201)+10)) = uint8(v215)
											v217 = *(*int32)(unsafe.Add(mBase, uint32(v145)))
											*(*uint8)(unsafe.Add(mBase, uint32(v217)+11)) = uint8(v208)
											v220 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v217)+10)))
											v222 = v220 & int32(3)
											*(*uint8)(unsafe.Add(mBase, uint32(v217)+10)) = uint8(v222)
											v224 = *(*int32)(unsafe.Add(mBase, uint32(v143)))
											*(*int32)(unsafe.Add(mBase, uint32(v143))) = v224 + int32(12)
											v231 = v206
										}
										v242 = v231
									} else {
										v182 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v168)+11)))
										if v182 != 0 {
											if base.Ui32(v135) < base.Ui32(int32(13)) {
												v231 = v146
											} else {
												v191 = *(*int32)(unsafe.Add(mBase, uint32(v143)))
												v192 = v141 + v191
												*(*int32)(unsafe.Add(mBase, uint32(v145))) = v192
												v194 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v192)+10)))
												v199 = v194&int32(254) | int32(0)
												*(*uint8)(unsafe.Add(mBase, uint32(v192)+10)) = uint8(v199)
												v201 = *(*int32)(unsafe.Add(mBase, uint32(v145)))
												*(*uint16)(unsafe.Add(mBase, uint32(v201)+8)) = uint16(v165)
												v204 = v135 + int32(-12)
												if base.Ui32(v204) < base.Ui32(l2) {
													v206 = v204
												} else {
													v206 = v167
												}
												*(*int32)(unsafe.Add(mBase, uint32(v201))) = v206
												v208 = int32(0)
												*(*int32)(unsafe.Add(mBase, uint32(v201)+4)) = v208
												v210 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v201)+10)))
												v215 = v210&int32(253) | v151<<(uint(int32(1))%32)
												*(*uint8)(unsafe.Add(mBase, uint32(v201)+10)) = uint8(v215)
												v217 = *(*int32)(unsafe.Add(mBase, uint32(v145)))
												*(*uint8)(unsafe.Add(mBase, uint32(v217)+11)) = uint8(v208)
												v220 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v217)+10)))
												v222 = v220 & int32(3)
												*(*uint8)(unsafe.Add(mBase, uint32(v217)+10)) = uint8(v222)
												v224 = *(*int32)(unsafe.Add(mBase, uint32(v143)))
												*(*int32)(unsafe.Add(mBase, uint32(v143))) = v224 + int32(12)
												v231 = v206
											}
											v242 = v231
										} else {
											v183 = *(*int32)(unsafe.Add(mBase, uint32(v145)))
											v184 = *(*int32)(unsafe.Add(mBase, uint32(v183)))
											*(*int32)(unsafe.Add(mBase, uint32(v183))) = v184 + v167
											v242 = v167
										}
									}
								}
							}
						}
					}
					v243 = v242
				} else {
					if base.Ui32(v135) < base.Ui32(l2) {
						v140 = v135
					} else {
						v140 = l2
					}
					v243 = v140
				}
				if v243 == int32(0) {
					v260 = v4
				} else {
					v246 = *(*int32)(unsafe.Add(mBase, uint32(l0)+124))
					v247 = *(*int32)(unsafe.Add(mBase, uint32(l0)+180))
					if v243 == int32(0) {
					} else {
						v251 = F__emscripten_memcpy_bulkmem(m, v246+v247, l1, v243)
						mBase = m.M
					}
					v253 = *(*int32)(unsafe.Add(mBase, uint32(l0)+180))
					v254 = v253 + v243
					*(*int32)(unsafe.Add(mBase, uint32(l0)+180)) = v254
					v256 = *(*int32)(unsafe.Add(mBase, uint32(l0)+276))
					if base.Ui32(v254) <= base.Ui32(v256) {
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(l0)+276)) = v254
					}
					v260 = v243
				}
			}
		}
	}
	return v260
}
func F_addReplyBool(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
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
	v3 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+224)))
	if v3 != int32(2) {
		v14 = F_prepareClientToWrite(m, l0)
		mBase = m.M
		v15 = m.ExcPending
		if v15 != 0 {
			return
		} else {
			if v14 != 0 {
				return
			} else {
				if l1 != 0 {
					v18 = int32(_a_F_addReplyBool_0)
				} else {
					v18 = int32(_a_F_addReplyBool_1)
				}
				F__addReplyToBufferOrList(m, l0, v18, int32(4))
				mBase = m.M
				v21 = m.ExcPending
				if v21 != 0 {
					return
				} else {
					return
				}
			}
		}
	} else {
		if l1 != 0 {
			v9 = int32(16)
		} else {
			v9 = int32(12)
		}
		v11 = *(*int32)(unsafe.Add(mBase, uint32(v9)+uint32(_c_F_addReplyBool[0])))
		F_addReply(m, l0, v11)
		mBase = m.M
		v13 = m.ExcPending
		if v13 != 0 {
			return
		} else {
			return
		}
	}
}
func F_addReplyBulkCString(m *base.Module, l0 int32, l1 int32) {
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
	var v73 int32
	_ = v73
	if l1 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	return
L2:
	;
	if l1&int32(3) == int32(0) {
		v38 = l1
		goto L14
	} else {
		goto L15
	}
L3:
	;
	v4 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+224)))
	v5 = F_prepareClientToWrite(m, l0)
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	return
L5:
	;
	if v4 != int32(2) {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	if v5 != 0 {
		goto L1
	} else {
		goto L10
	}
L7:
	;
	if v5 != 0 {
		goto L1
	} else {
		goto L8
	}
L8:
	;
	F__addReplyToBufferOrList(m, l0, int32(_a_F_addReplyBulkCString_0), int32(5))
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		goto L4
	} else {
		goto L9
	}
L9:
	;
	return
L10:
	;
	F__addReplyToBufferOrList(m, l0, int32(_a_F_addReplyBulkCString_1), int32(3))
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		goto L4
	} else {
		goto L11
	}
L11:
	;
	return
L12:
	;
	F_addReplyBulkCBuffer(m, l0, l1, v71)
	mBase = m.M
	v73 = m.ExcPending
	if v73 != 0 {
		goto L4
	} else {
		goto L28
	}
L13:
	;
	v71 = v63 - l1
	goto L12
L14:
	;
	v42 = v38
	goto L22
L15:
	;
	v24 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	if v24 != 0 {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	v27 = l1
	goto L18
L17:
	;
	v71 = l1 - l1
	goto L12
L18:
	;
	v31 = v27 + int32(1)
	if v31&int32(3) == int32(0) {
		v38 = v31
		goto L14
	} else {
		goto L20
	}
L20:
	;
	v36 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v31))))
	if v36 != 0 {
		v27 = v31
		goto L18
	} else {
		goto L21
	}
L21:
	;
	v63 = v31
	goto L13
L22:
	;
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v42)))
	v51 = int32(-2139062144)
	if (int32(16843008)-v48|v48)&v51 == v51 {
		v42 = v42 + int32(4)
		goto L22
	} else {
		goto L24
	}
L23:
	;
	v57 = v42
	goto L25
L24:
	;
	goto L23
L25:
	;
	v61 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v57))))
	if v61 != 0 {
		v57 = v57 + int32(1)
		goto L25
	} else {
		goto L27
	}
L26:
	;
	v63 = v57
	goto L13
L27:
	;
	goto L26
L28:
	;
	goto L1
}
func F_addReplyCommandCategories(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v19 int64
	_ = v19
	var v22 int32
	_ = v22
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v30 int64
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int64
	_ = v33
	var v37 int32
	_ = v37
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v53 int64
	_ = v53
	var v61 int32
	_ = v61
	var v68 int32
	_ = v68
	v10 = m.G0
	v12 = v10 - int32(16)
	m.G0 = v12
	v15 = F_addReplyDeferredLen(m, l0)
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
	v18 = *(*int32)(unsafe.Add(mBase, _c_F_addReplyCommandCategories[0]))
	v19 = *(*int64)(unsafe.Add(mBase, uint32(v18)+8))
	if v19 == int64(0) {
		v61 = int32(0)
		goto L3
	} else {
		goto L4
	}
L3:
	;
	F_setDeferredSetLen(m, l0, v15, v61)
	mBase = m.M
	v68 = m.ExcPending
	if v68 != 0 {
		goto L1
	} else {
		goto L11
	}
L4:
	;
	v22 = int32(0)
	v27 = v22
	v29 = v18
	v30 = v19
	v31 = v18
	v32 = v22
	goto L5
L5:
	;
	v33 = *(*int64)(unsafe.Add(mBase, uint32(l1)+64))
	if v33&v30 == int64(0) {
		v46 = v27
		v47 = v29
		goto L7
	} else {
		goto L8
	}
L6:
	;
	v61 = v46
	goto L3
L7:
	;
	v49 = v32 + int32(1)
	v52 = v47 + v49<<(uint(int32(4))%32)
	v53 = *(*int64)(unsafe.Add(mBase, uint32(v52)+8))
	if base.B2i32(v53 == int64(0)) == int32(0) {
		v27 = v46
		v29 = v47
		v30 = v53
		v31 = v52
		v32 = v49
		goto L5
	} else {
		goto L10
	}
L8:
	;
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v31)))
	*(*int32)(unsafe.Add(mBase, uint32(v12))) = v37
	F_addReplyStatusFormat(m, l0, int32(_a_F_addReplyCommandCategories_0), v12)
	mBase = m.M
	v41 = m.ExcPending
	if v41 != 0 {
		goto L1
	} else {
		goto L9
	}
L9:
	;
	v45 = *(*int32)(unsafe.Add(mBase, _c_F_addReplyCommandCategories[0]))
	v46 = v27 + int32(1)
	v47 = v45
	goto L7
L10:
	;
	goto L6
L11:
	;
	m.G0 = v12 + int32(16)
	return
}
func F_addReplyCommandKeySpecs(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
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
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
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
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v52 int64
	_ = v52
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v83 int64
	_ = v83
	var v85 int32
	_ = v85
	var v88 int32
	_ = v88
	var v91 int32
	_ = v91
	var v94 int32
	_ = v94
	var v97 int32
	_ = v97
	var v100 int32
	_ = v100
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v114 int64
	_ = v114
	var v116 int32
	_ = v116
	var v122 int32
	_ = v122
	var v126 int32
	_ = v126
	var v129 int32
	_ = v129
	var v132 int32
	_ = v132
	var v135 int32
	_ = v135
	var v138 int32
	_ = v138
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v149 int32
	_ = v149
	var v152 int32
	_ = v152
	var v155 int32
	_ = v155
	var v158 int32
	_ = v158
	var v161 int32
	_ = v161
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v167 int64
	_ = v167
	var v169 int32
	_ = v169
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v175 int64
	_ = v175
	var v177 int32
	_ = v177
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v183 int64
	_ = v183
	var v185 int32
	_ = v185
	var v188 int32
	_ = v188
	var v191 int32
	_ = v191
	var v194 int32
	_ = v194
	var v197 int32
	_ = v197
	var v200 int32
	_ = v200
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v206 int64
	_ = v206
	var v208 int32
	_ = v208
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v214 int64
	_ = v214
	var v216 int32
	_ = v216
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v222 int64
	_ = v222
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v233 int32
	_ = v233
	var v237 int32
	_ = v237
	var v240 int32
	_ = v240
	var v243 int32
	_ = v243
	var v246 int32
	_ = v246
	var v249 int32
	_ = v249
	var v251 int32
	_ = v251
	var v252 int32
	_ = v252
	v7 = m.G0
	v9 = v7 - int32(32)
	m.G0 = v9
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l1)+80))
	F_addReplySetLen(m, l0, v11)
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l1)+80))
	if v14 < int32(1) {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	m.G0 = v9 + int32(32)
	return
L4:
	;
	v21 = int32(0)
	goto L5
L5:
	;
	v26 = *(*int32)(unsafe.Add(mBase, uint32(l1)+76))
	v28 = v21 * int32(48)
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v26+v28)))
	if v30 != 0 {
		goto L7
	} else {
		goto L8
	}
L6:
	;
	goto L3
L7:
	;
	v31 = int32(4)
	goto L9
L8:
	;
	v31 = int32(3)
	goto L9
L9:
	;
	F_addReplyMapLen(m, l0, v31)
	mBase = m.M
	v33 = m.ExcPending
	if v33 != 0 {
		goto L1
	} else {
		goto L10
	}
L10:
	;
	v34 = *(*int32)(unsafe.Add(mBase, uint32(l1)+76))
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v34+v28)))
	if v36 == int32(0) {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	F_addReplyBulkCString(m, l0, int32(_a_F_addReplyCommandKeySpecs_0))
	mBase = m.M
	v49 = m.ExcPending
	if v49 != 0 {
		goto L1
	} else {
		goto L15
	}
L12:
	;
	F_addReplyBulkCString(m, l0, int32(_a_F_addReplyCommandKeySpecs_1))
	mBase = m.M
	v41 = m.ExcPending
	if v41 != 0 {
		goto L1
	} else {
		goto L13
	}
L13:
	;
	v42 = *(*int32)(unsafe.Add(mBase, uint32(l1)+76))
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v42+v28)))
	F_addReplyBulkCString(m, l0, v44)
	mBase = m.M
	v46 = m.ExcPending
	if v46 != 0 {
		goto L1
	} else {
		goto L14
	}
L14:
	;
	goto L11
L15:
	;
	v50 = *(*int32)(unsafe.Add(mBase, uint32(l1)+76))
	v52 = *(*int64)(unsafe.Add(mBase, uint32(v50+v28)+8))
	F_addReplyFlagsForKeyArgs(m, l0, v52)
	mBase = m.M
	v54 = m.ExcPending
	if v54 != 0 {
		goto L1
	} else {
		goto L16
	}
L16:
	;
	F_addReplyBulkCString(m, l0, int32(_a_F_addReplyCommandKeySpecs_2))
	mBase = m.M
	v57 = m.ExcPending
	if v57 != 0 {
		goto L1
	} else {
		goto L17
	}
L17:
	;
	v58 = *(*int32)(unsafe.Add(mBase, uint32(l1)+76))
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v58+v28)+16))
	switch v60 + int32(-1) {
	case 0:
		goto L19
	case 1:
		goto L22
	case 2:
		goto L21
	default:
		goto L20
	}
L18:
	;
	F_addReplyBulkCString(m, l0, int32(_a_F_addReplyCommandKeySpecs_3))
	mBase = m.M
	v141 = m.ExcPending
	if v141 != 0 {
		goto L1
	} else {
		goto L45
	}
L19:
	;
	F_addReplyMapLen(m, l0, int32(2))
	mBase = m.M
	v126 = m.ExcPending
	if v126 != 0 {
		goto L1
	} else {
		goto L40
	}
L20:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9))) = v60
	F__serverPanic_1(m, int32(_a_F_addReplyCommandKeySpecs_4), int32(5281), int32(_a_F_addReplyCommandKeySpecs_5), v9)
	mBase = m.M
	v122 = m.ExcPending
	if v122 != 0 {
		goto L1
	} else {
		goto L39
	}
L21:
	;
	F_addReplyMapLen(m, l0, int32(2))
	mBase = m.M
	v88 = m.ExcPending
	if v88 != 0 {
		goto L1
	} else {
		goto L30
	}
L22:
	;
	F_addReplyMapLen(m, l0, int32(2))
	mBase = m.M
	v65 = m.ExcPending
	if v65 != 0 {
		goto L1
	} else {
		goto L23
	}
L23:
	;
	F_addReplyBulkCString(m, l0, int32(_a_F_addReplyCommandKeySpecs_6))
	mBase = m.M
	v68 = m.ExcPending
	if v68 != 0 {
		goto L1
	} else {
		goto L24
	}
L24:
	;
	F_addReplyBulkCString(m, l0, int32(_a_F_addReplyCommandKeySpecs_7))
	mBase = m.M
	v71 = m.ExcPending
	if v71 != 0 {
		goto L1
	} else {
		goto L25
	}
L25:
	;
	F_addReplyBulkCString(m, l0, int32(_a_F_addReplyCommandKeySpecs_8))
	mBase = m.M
	v74 = m.ExcPending
	if v74 != 0 {
		goto L1
	} else {
		goto L26
	}
L26:
	;
	F_addReplyMapLen(m, l0, int32(1))
	mBase = m.M
	v77 = m.ExcPending
	if v77 != 0 {
		goto L1
	} else {
		goto L27
	}
L27:
	;
	F_addReplyBulkCString(m, l0, int32(_a_F_addReplyCommandKeySpecs_7))
	mBase = m.M
	v80 = m.ExcPending
	if v80 != 0 {
		goto L1
	} else {
		goto L28
	}
L28:
	;
	v81 = *(*int32)(unsafe.Add(mBase, uint32(l1)+76))
	v83 = int64(*(*int32)(unsafe.Add(mBase, uint32(v81+v28)+20)))
	F_addReplyLongLong(m, l0, v83)
	mBase = m.M
	v85 = m.ExcPending
	if v85 != 0 {
		goto L1
	} else {
		goto L29
	}
L29:
	;
	goto L18
L30:
	;
	F_addReplyBulkCString(m, l0, int32(_a_F_addReplyCommandKeySpecs_6))
	mBase = m.M
	v91 = m.ExcPending
	if v91 != 0 {
		goto L1
	} else {
		goto L31
	}
L31:
	;
	F_addReplyBulkCString(m, l0, int32(_a_F_addReplyCommandKeySpecs_9))
	mBase = m.M
	v94 = m.ExcPending
	if v94 != 0 {
		goto L1
	} else {
		goto L32
	}
L32:
	;
	F_addReplyBulkCString(m, l0, int32(_a_F_addReplyCommandKeySpecs_8))
	mBase = m.M
	v97 = m.ExcPending
	if v97 != 0 {
		goto L1
	} else {
		goto L33
	}
L33:
	;
	F_addReplyMapLen(m, l0, int32(2))
	mBase = m.M
	v100 = m.ExcPending
	if v100 != 0 {
		goto L1
	} else {
		goto L34
	}
L34:
	;
	F_addReplyBulkCString(m, l0, int32(_a_F_addReplyCommandKeySpecs_9))
	mBase = m.M
	v103 = m.ExcPending
	if v103 != 0 {
		goto L1
	} else {
		goto L35
	}
L35:
	;
	v104 = *(*int32)(unsafe.Add(mBase, uint32(l1)+76))
	v106 = *(*int32)(unsafe.Add(mBase, uint32(v104+v28)+20))
	F_addReplyBulkCString(m, l0, v106)
	mBase = m.M
	v108 = m.ExcPending
	if v108 != 0 {
		goto L1
	} else {
		goto L36
	}
L36:
	;
	F_addReplyBulkCString(m, l0, int32(_a_F_addReplyCommandKeySpecs_10))
	mBase = m.M
	v111 = m.ExcPending
	if v111 != 0 {
		goto L1
	} else {
		goto L37
	}
L37:
	;
	v112 = *(*int32)(unsafe.Add(mBase, uint32(l1)+76))
	v114 = int64(*(*int32)(unsafe.Add(mBase, uint32(v112+v28)+24)))
	F_addReplyLongLong(m, l0, v114)
	mBase = m.M
	v116 = m.ExcPending
	if v116 != 0 {
		goto L1
	} else {
		goto L38
	}
L38:
	;
	goto L18
L39:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L40:
	;
	F_addReplyBulkCString(m, l0, int32(_a_F_addReplyCommandKeySpecs_6))
	mBase = m.M
	v129 = m.ExcPending
	if v129 != 0 {
		goto L1
	} else {
		goto L41
	}
L41:
	;
	F_addReplyBulkCString(m, l0, int32(_a_F_addReplyCommandKeySpecs_11))
	mBase = m.M
	v132 = m.ExcPending
	if v132 != 0 {
		goto L1
	} else {
		goto L42
	}
L42:
	;
	F_addReplyBulkCString(m, l0, int32(_a_F_addReplyCommandKeySpecs_8))
	mBase = m.M
	v135 = m.ExcPending
	if v135 != 0 {
		goto L1
	} else {
		goto L43
	}
L43:
	;
	F_addReplyMapLen(m, l0, int32(0))
	mBase = m.M
	v138 = m.ExcPending
	if v138 != 0 {
		goto L1
	} else {
		goto L44
	}
L44:
	;
	goto L18
L45:
	;
	v142 = *(*int32)(unsafe.Add(mBase, uint32(l1)+76))
	v143 = v142 + v28
	v144 = *(*int32)(unsafe.Add(mBase, uint32(v143)+28))
	switch v144 + int32(-1) {
	case 0:
		goto L47
	case 1:
		goto L50
	case 2:
		goto L49
	default:
		goto L48
	}
L46:
	;
	v251 = v21 + int32(1)
	v252 = *(*int32)(unsafe.Add(mBase, uint32(l1)+80))
	if v251 < v252 {
		v21 = v251
		goto L5
	} else {
		goto L79
	}
L47:
	;
	F_addReplyMapLen(m, l0, int32(2))
	mBase = m.M
	v237 = m.ExcPending
	if v237 != 0 {
		goto L1
	} else {
		goto L74
	}
L48:
	;
	v225 = *(*int32)(unsafe.Add(mBase, uint32(v143)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = v225
	F__serverPanic_1(m, int32(_a_F_addReplyCommandKeySpecs_4), int32(5322), int32(_a_F_addReplyCommandKeySpecs_12), v9+int32(16))
	mBase = m.M
	v233 = m.ExcPending
	if v233 != 0 {
		goto L1
	} else {
		goto L73
	}
L49:
	;
	F_addReplyMapLen(m, l0, int32(2))
	mBase = m.M
	v188 = m.ExcPending
	if v188 != 0 {
		goto L1
	} else {
		goto L62
	}
L50:
	;
	F_addReplyMapLen(m, l0, int32(2))
	mBase = m.M
	v149 = m.ExcPending
	if v149 != 0 {
		goto L1
	} else {
		goto L51
	}
L51:
	;
	F_addReplyBulkCString(m, l0, int32(_a_F_addReplyCommandKeySpecs_6))
	mBase = m.M
	v152 = m.ExcPending
	if v152 != 0 {
		goto L1
	} else {
		goto L52
	}
L52:
	;
	F_addReplyBulkCString(m, l0, int32(_a_F_addReplyCommandKeySpecs_13))
	mBase = m.M
	v155 = m.ExcPending
	if v155 != 0 {
		goto L1
	} else {
		goto L53
	}
L53:
	;
	F_addReplyBulkCString(m, l0, int32(_a_F_addReplyCommandKeySpecs_8))
	mBase = m.M
	v158 = m.ExcPending
	if v158 != 0 {
		goto L1
	} else {
		goto L54
	}
L54:
	;
	F_addReplyMapLen(m, l0, int32(3))
	mBase = m.M
	v161 = m.ExcPending
	if v161 != 0 {
		goto L1
	} else {
		goto L55
	}
L55:
	;
	F_addReplyBulkCString(m, l0, int32(_a_F_addReplyCommandKeySpecs_14))
	mBase = m.M
	v164 = m.ExcPending
	if v164 != 0 {
		goto L1
	} else {
		goto L56
	}
L56:
	;
	v165 = *(*int32)(unsafe.Add(mBase, uint32(l1)+76))
	v167 = int64(*(*int32)(unsafe.Add(mBase, uint32(v165+v28)+32)))
	F_addReplyLongLong(m, l0, v167)
	mBase = m.M
	v169 = m.ExcPending
	if v169 != 0 {
		goto L1
	} else {
		goto L57
	}
L57:
	;
	F_addReplyBulkCString(m, l0, int32(_a_F_addReplyCommandKeySpecs_15))
	mBase = m.M
	v172 = m.ExcPending
	if v172 != 0 {
		goto L1
	} else {
		goto L58
	}
L58:
	;
	v173 = *(*int32)(unsafe.Add(mBase, uint32(l1)+76))
	v175 = int64(*(*int32)(unsafe.Add(mBase, uint32(v173+v28)+36)))
	F_addReplyLongLong(m, l0, v175)
	mBase = m.M
	v177 = m.ExcPending
	if v177 != 0 {
		goto L1
	} else {
		goto L59
	}
L59:
	;
	F_addReplyBulkCString(m, l0, int32(_a_F_addReplyCommandKeySpecs_16))
	mBase = m.M
	v180 = m.ExcPending
	if v180 != 0 {
		goto L1
	} else {
		goto L60
	}
L60:
	;
	v181 = *(*int32)(unsafe.Add(mBase, uint32(l1)+76))
	v183 = int64(*(*int32)(unsafe.Add(mBase, uint32(v181+v28)+40)))
	F_addReplyLongLong(m, l0, v183)
	mBase = m.M
	v185 = m.ExcPending
	if v185 != 0 {
		goto L1
	} else {
		goto L61
	}
L61:
	;
	goto L46
L62:
	;
	F_addReplyBulkCString(m, l0, int32(_a_F_addReplyCommandKeySpecs_6))
	mBase = m.M
	v191 = m.ExcPending
	if v191 != 0 {
		goto L1
	} else {
		goto L63
	}
L63:
	;
	F_addReplyBulkCString(m, l0, int32(_a_F_addReplyCommandKeySpecs_17))
	mBase = m.M
	v194 = m.ExcPending
	if v194 != 0 {
		goto L1
	} else {
		goto L64
	}
L64:
	;
	F_addReplyBulkCString(m, l0, int32(_a_F_addReplyCommandKeySpecs_8))
	mBase = m.M
	v197 = m.ExcPending
	if v197 != 0 {
		goto L1
	} else {
		goto L65
	}
L65:
	;
	F_addReplyMapLen(m, l0, int32(3))
	mBase = m.M
	v200 = m.ExcPending
	if v200 != 0 {
		goto L1
	} else {
		goto L66
	}
L66:
	;
	F_addReplyBulkCString(m, l0, int32(_a_F_addReplyCommandKeySpecs_18))
	mBase = m.M
	v203 = m.ExcPending
	if v203 != 0 {
		goto L1
	} else {
		goto L67
	}
L67:
	;
	v204 = *(*int32)(unsafe.Add(mBase, uint32(l1)+76))
	v206 = int64(*(*int32)(unsafe.Add(mBase, uint32(v204+v28)+32)))
	F_addReplyLongLong(m, l0, v206)
	mBase = m.M
	v208 = m.ExcPending
	if v208 != 0 {
		goto L1
	} else {
		goto L68
	}
L68:
	;
	F_addReplyBulkCString(m, l0, int32(_a_F_addReplyCommandKeySpecs_19))
	mBase = m.M
	v211 = m.ExcPending
	if v211 != 0 {
		goto L1
	} else {
		goto L69
	}
L69:
	;
	v212 = *(*int32)(unsafe.Add(mBase, uint32(l1)+76))
	v214 = int64(*(*int32)(unsafe.Add(mBase, uint32(v212+v28)+36)))
	F_addReplyLongLong(m, l0, v214)
	mBase = m.M
	v216 = m.ExcPending
	if v216 != 0 {
		goto L1
	} else {
		goto L70
	}
L70:
	;
	F_addReplyBulkCString(m, l0, int32(_a_F_addReplyCommandKeySpecs_15))
	mBase = m.M
	v219 = m.ExcPending
	if v219 != 0 {
		goto L1
	} else {
		goto L71
	}
L71:
	;
	v220 = *(*int32)(unsafe.Add(mBase, uint32(l1)+76))
	v222 = int64(*(*int32)(unsafe.Add(mBase, uint32(v220+v28)+40)))
	F_addReplyLongLong(m, l0, v222)
	mBase = m.M
	v224 = m.ExcPending
	if v224 != 0 {
		goto L1
	} else {
		goto L72
	}
L72:
	;
	goto L46
L73:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L74:
	;
	F_addReplyBulkCString(m, l0, int32(_a_F_addReplyCommandKeySpecs_6))
	mBase = m.M
	v240 = m.ExcPending
	if v240 != 0 {
		goto L1
	} else {
		goto L75
	}
L75:
	;
	F_addReplyBulkCString(m, l0, int32(_a_F_addReplyCommandKeySpecs_11))
	mBase = m.M
	v243 = m.ExcPending
	if v243 != 0 {
		goto L1
	} else {
		goto L76
	}
L76:
	;
	F_addReplyBulkCString(m, l0, int32(_a_F_addReplyCommandKeySpecs_8))
	mBase = m.M
	v246 = m.ExcPending
	if v246 != 0 {
		goto L1
	} else {
		goto L77
	}
L77:
	;
	F_addReplyMapLen(m, l0, int32(0))
	mBase = m.M
	v249 = m.ExcPending
	if v249 != 0 {
		goto L1
	} else {
		goto L78
	}
L78:
	;
	goto L46
L79:
	;
	goto L6
}
func F_addReplyCommandSubCommands(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v47 int32
	_ = v47
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v64 int32
	_ = v64
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v95 int32
	_ = v95
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v109 int32
	_ = v109
	v9 = m.G0
	v11 = v9 - int32(64)
	m.G0 = v11
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l1)+200))
	if v13 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v11 + int32(64)
	return
L2:
	;
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v13)+20))
	v23 = *(*int32)(unsafe.Add(mBase, uint32(v13)+16))
	v24 = v22 + v23
	goto L9
L3:
	;
	if l3 == int32(0) {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	F_addReplyArrayLen(m, l0, int32(0))
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		goto L6
	} else {
		goto L8
	}
L5:
	;
	F_addReplyMapLen(m, l0, int32(0))
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	return
L7:
	;
	goto L1
L8:
	;
	goto L1
L9:
	;
	if l3 == int32(0) {
		goto L11
	} else {
		goto L12
	}
L10:
	;
	v31 = *(*int32)(unsafe.Add(mBase, uint32(l1)+200))
	v32 = int32(1)
	v33 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v11)+14)) = uint8(v33)
	*(*int32)(unsafe.Add(mBase, uint32(v11))) = v31
	*(*int32)(unsafe.Add(mBase, uint32(v11)+24)) = v33
	*(*uint8)(unsafe.Add(mBase, uint32(v11)+15)) = uint8(v32)
	*(*int32)(unsafe.Add(mBase, uint32(v11)+8)) = int32(-1)
	if v31 == v33 {
		goto L16
	} else {
		goto L17
	}
L11:
	;
	F_addReplyArrayLen(m, l0, v24)
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		goto L6
	} else {
		goto L14
	}
L12:
	;
	F_addReplyMapLen(m, l0, v24)
	mBase = m.M
	v28 = m.ExcPending
	if v28 != 0 {
		goto L6
	} else {
		goto L13
	}
L13:
	;
	goto L10
L14:
	;
	goto L10
L15:
	;
	v52 = F_hashtableNext(m, v11, v11+int32(60))
	mBase = m.M
	v53 = m.ExcPending
	if v53 != 0 {
		goto L6
	} else {
		goto L20
	}
L16:
	;
	goto L15
L17:
	;
	goto L18
L18:
	;
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v31)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+24)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v31)+40)) = v11
	goto L16
L19:
	;
	F_hashtableCleanupIterator(m, v11)
	mBase = m.M
	v109 = m.ExcPending
	if v109 != 0 {
		goto L6
	} else {
		goto L36
	}
L20:
	;
	if v52 == int32(0) {
		goto L19
	} else {
		goto L21
	}
L21:
	;
	goto L22
L22:
	;
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v11)+60))
	if l3 == int32(0) {
		goto L24
	} else {
		goto L25
	}
L23:
	;
	goto L19
L24:
	;
	m.T0[l2].(func(*base.Module, int32, int32))(m, l0, v64)
	mBase = m.M
	v95 = m.ExcPending
	if v95 != 0 {
		goto L6
	} else {
		goto L33
	}
L25:
	;
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v64)+140))
	v71 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v68+int32(-1)))))
	switch v71 & int32(7) {
	case 0:
		goto L31
	case 1:
		goto L30
	case 2:
		goto L29
	case 3:
		goto L28
	case 4:
		goto L27
	default:
		v88 = int32(0)
		goto L26
	}
L26:
	;
	F_addReplyBulkCBuffer(m, l0, v68, v88)
	mBase = m.M
	v90 = m.ExcPending
	if v90 != 0 {
		goto L6
	} else {
		goto L32
	}
L27:
	;
	v87 = *(*int32)(unsafe.Add(mBase, uint32(v68+int32(-17))))
	v88 = v87
	goto L26
L28:
	;
	v84 = *(*int32)(unsafe.Add(mBase, uint32(v68+int32(-9))))
	v88 = v84
	goto L26
L29:
	;
	v81 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v68+int32(-5)))))
	v88 = v81
	goto L26
L30:
	;
	v78 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v68+int32(-3)))))
	v88 = v78
	goto L26
L31:
	;
	v88 = int32(base.Ui32(v71) >> (uint(int32(3)) % 32))
	goto L26
L32:
	;
	goto L24
L33:
	;
	v98 = F_hashtableNext(m, v11, v11+int32(60))
	mBase = m.M
	v99 = m.ExcPending
	if v99 != 0 {
		goto L6
	} else {
		goto L34
	}
L34:
	;
	if v98 != 0 {
		goto L22
	} else {
		goto L35
	}
L35:
	;
	goto L23
L36:
	;
	goto L1
}
func F_addReplyDictOfValkeyInstances(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v30 int32
	_ = v30
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
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
	var v49 int64
	_ = v49
	var v51 int64
	_ = v51
	var v53 int64
	_ = v53
	var v55 int64
	_ = v55
	var v57 int64
	_ = v57
	var v59 int64
	_ = v59
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v81 int32
	_ = v81
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v101 int32
	_ = v101
	var v106 int32
	_ = v106
	var v110 int32
	_ = v110
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
	var v131 int32
	_ = v131
	var v135 int32
	_ = v135
	var v138 int32
	_ = v138
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v150 int32
	_ = v150
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v158 int32
	_ = v158
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v167 int32
	_ = v167
	var v169 int32
	_ = v169
	var v171 int64
	_ = v171
	var v172 int64
	_ = v172
	var v173 int64
	_ = v173
	var v174 int64
	_ = v174
	var v175 int64
	_ = v175
	var v176 int64
	_ = v176
	var v177 int64
	_ = v177
	var v179 int64
	_ = v179
	var v181 int64
	_ = v181
	var v183 int64
	_ = v183
	var v185 int64
	_ = v185
	var v187 int64
	_ = v187
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v194 int32
	_ = v194
	var v197 int32
	_ = v197
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
	var v209 int32
	_ = v209
	var v213 int32
	_ = v213
	var v215 int32
	_ = v215
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v229 int32
	_ = v229
	var v234 int32
	_ = v234
	var v238 int32
	_ = v238
	var v242 int32
	_ = v242
	var v248 int32
	_ = v248
	var v252 int32
	_ = v252
	var v254 int32
	_ = v254
	v3 = int32(0)
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
	v9 = F_dictGetIterator(m, l1)
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		goto L1
	} else {
		goto L4
	}
L3:
	;
	F_dictReleaseIterator(m, v9)
	mBase = m.M
	v252 = m.ExcPending
	if v252 != 0 {
		goto L1
	} else {
		goto L67
	}
L4:
	;
	v18 = v9 + int32(20)
	v19 = *(*int32)(unsafe.Add(mBase, uint32(v9)+16))
	if v19 != 0 {
		goto L7
	} else {
		goto L8
	}
L5:
	;
	if v114 == int32(0) {
		v248 = v3
		goto L3
	} else {
		goto L31
	}
L6:
	;
	v25 = v18
	v26 = v22
	goto L9
L7:
	;
	v22 = int32(1)
	goto L6
L8:
	;
	v22 = int32(0)
	goto L6
L9:
	;
	switch v26 {
	case 0:
		goto L14
	default:
		goto L13
	}
L11:
	;
	v26 = int32(0)
	goto L9
L12:
	;
	goto L5
L13:
	;
	v106 = *(*int32)(unsafe.Add(mBase, uint32(v25)))
	*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = v106
	if v106 == int32(0) {
		goto L11
	} else {
		goto L30
	}
L14:
	;
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v9)+4))
	if v30 != int32(-1) {
		v69 = v30
		goto L15
	} else {
		goto L16
	}
L15:
	;
	v70 = int32(1)
	v71 = v69 + v70
	*(*int32)(unsafe.Add(mBase, uint32(v9)+4)) = v71
	v73 = int32(0)
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v9)))
	v77 = *(*int32)(unsafe.Add(mBase, uint32(v9)+8))
	v81 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v76+v77+int32(26)))))
	if v81 == int32(255) {
		goto L24
	} else {
		goto L25
	}
L16:
	;
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v9)+8))
	if v34 != 0 {
		v69 = int32(-1)
		goto L15
	} else {
		goto L17
	}
L17:
	;
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v9)))
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v9)+12))
	if v36 == int32(0) {
		goto L19
	} else {
		goto L20
	}
L18:
	;
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v62)+20))
	if v63 != int32(-1) {
		goto L21
	} else {
		goto L22
	}
L19:
	;
	v43 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v35)+16)))
	v44 = int64(*(*int8)(unsafe.Add(mBase, uint32(v35)+27)))
	v45 = int64(*(*int32)(unsafe.Add(mBase, uint32(v35)+8)))
	v46 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v35)+12)))
	v47 = int64(*(*int8)(unsafe.Add(mBase, uint32(v35)+26)))
	v48 = int64(*(*int32)(unsafe.Add(mBase, uint32(v35)+4)))
	v49 = F_wangHash64(m, v48)
	mBase = m.M
	v51 = F_wangHash64(m, v47+v49)
	mBase = m.M
	v53 = F_wangHash64(m, v46+v51)
	mBase = m.M
	v55 = F_wangHash64(m, v45+v53)
	mBase = m.M
	v57 = F_wangHash64(m, v44+v55)
	mBase = m.M
	v59 = F_wangHash64(m, v43+v57)
	mBase = m.M
	*(*int64)(unsafe.Add(mBase, uint32(v9)+24)) = v59
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v9)))
	v62 = v61
	goto L18
L20:
	;
	v39 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v35)+24)))
	v41 = v39 + int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v35)+24)) = uint16(v41)
	v62 = v35
	goto L18
L21:
	;
	v69 = v63 + int32(-1)
	goto L15
L22:
	;
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v9)+4))
	v69 = v66
	goto L15
L23:
	;
	v96 = int32(2)
	v101 = *(*int32)(unsafe.Add(mBase, uint32(v76+v94<<(uint(v96)%32)+int32(4))))
	v25 = v101 + v95<<(uint(v96)%32)
	v26 = int32(1)
	goto L9
L24:
	;
	v85 = v73
	goto L26
L25:
	;
	v85 = v70 << (uint(v81) % 32)
	goto L26
L26:
	;
	if v71 < v85 {
		v94 = v77
		v95 = v71
		goto L23
	} else {
		goto L27
	}
L27:
	;
	if v77 != 0 {
		v114 = v73
		goto L12
	} else {
		goto L28
	}
L28:
	;
	v87 = *(*int32)(unsafe.Add(mBase, uint32(v76)+20))
	if v87 == int32(-1) {
		v114 = v73
		goto L12
	} else {
		goto L29
	}
L29:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v9)+4)) = int64(4294967296)
	v94 = int32(1)
	v95 = int32(0)
	goto L23
L30:
	;
	v110 = *(*int32)(unsafe.Add(mBase, uint32(v106)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v18))) = v110
	v114 = v106
	goto L12
L31:
	;
	v121 = v114
	v122 = v3
	goto L32
L32:
	;
	v125 = *(*int32)(unsafe.Add(mBase, uint32(v121)+8))
	goto L36
L33:
	;
	v248 = v138
	goto L3
L34:
	;
	v146 = v9 + int32(20)
	v147 = *(*int32)(unsafe.Add(mBase, uint32(v9)+16))
	if v147 != 0 {
		goto L42
	} else {
		goto L43
	}
L35:
	;
	F_addReplySentinelValkeyInstance(m, l0, v125)
	mBase = m.M
	v135 = m.ExcPending
	if v135 != 0 {
		goto L1
	} else {
		goto L39
	}
L36:
	;
	v126 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v125))))
	if v126&int32(2) == int32(0) {
		goto L35
	} else {
		goto L37
	}
L37:
	;
	v131 = *(*int32)(unsafe.Add(mBase, uint32(v125)+180))
	if v131 == int32(0) {
		v138 = v122
		goto L34
	} else {
		goto L38
	}
L38:
	;
	goto L35
L39:
	;
	v138 = v122 + int32(1)
	goto L34
L40:
	;
	if v242 != 0 {
		v121 = v242
		v122 = v138
		goto L32
	} else {
		goto L66
	}
L41:
	;
	v153 = v146
	v154 = v150
	goto L44
L42:
	;
	v150 = int32(1)
	goto L41
L43:
	;
	v150 = int32(0)
	goto L41
L44:
	;
	switch v154 {
	case 0:
		goto L49
	default:
		goto L48
	}
L46:
	;
	v154 = int32(0)
	goto L44
L47:
	;
	goto L40
L48:
	;
	v234 = *(*int32)(unsafe.Add(mBase, uint32(v153)))
	*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = v234
	if v234 == int32(0) {
		goto L46
	} else {
		goto L65
	}
L49:
	;
	v158 = *(*int32)(unsafe.Add(mBase, uint32(v9)+4))
	if v158 != int32(-1) {
		v197 = v158
		goto L50
	} else {
		goto L51
	}
L50:
	;
	v198 = int32(1)
	v199 = v197 + v198
	*(*int32)(unsafe.Add(mBase, uint32(v9)+4)) = v199
	v201 = int32(0)
	v204 = *(*int32)(unsafe.Add(mBase, uint32(v9)))
	v205 = *(*int32)(unsafe.Add(mBase, uint32(v9)+8))
	v209 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v204+v205+int32(26)))))
	if v209 == int32(255) {
		goto L59
	} else {
		goto L60
	}
L51:
	;
	v162 = *(*int32)(unsafe.Add(mBase, uint32(v9)+8))
	if v162 != 0 {
		v197 = int32(-1)
		goto L50
	} else {
		goto L52
	}
L52:
	;
	v163 = *(*int32)(unsafe.Add(mBase, uint32(v9)))
	v164 = *(*int32)(unsafe.Add(mBase, uint32(v9)+12))
	if v164 == int32(0) {
		goto L54
	} else {
		goto L55
	}
L53:
	;
	v191 = *(*int32)(unsafe.Add(mBase, uint32(v190)+20))
	if v191 != int32(-1) {
		goto L56
	} else {
		goto L57
	}
L54:
	;
	v171 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v163)+16)))
	v172 = int64(*(*int8)(unsafe.Add(mBase, uint32(v163)+27)))
	v173 = int64(*(*int32)(unsafe.Add(mBase, uint32(v163)+8)))
	v174 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v163)+12)))
	v175 = int64(*(*int8)(unsafe.Add(mBase, uint32(v163)+26)))
	v176 = int64(*(*int32)(unsafe.Add(mBase, uint32(v163)+4)))
	v177 = F_wangHash64(m, v176)
	mBase = m.M
	v179 = F_wangHash64(m, v175+v177)
	mBase = m.M
	v181 = F_wangHash64(m, v174+v179)
	mBase = m.M
	v183 = F_wangHash64(m, v173+v181)
	mBase = m.M
	v185 = F_wangHash64(m, v172+v183)
	mBase = m.M
	v187 = F_wangHash64(m, v171+v185)
	mBase = m.M
	*(*int64)(unsafe.Add(mBase, uint32(v9)+24)) = v187
	v189 = *(*int32)(unsafe.Add(mBase, uint32(v9)))
	v190 = v189
	goto L53
L55:
	;
	v167 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v163)+24)))
	v169 = v167 + int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v163)+24)) = uint16(v169)
	v190 = v163
	goto L53
L56:
	;
	v197 = v191 + int32(-1)
	goto L50
L57:
	;
	v194 = *(*int32)(unsafe.Add(mBase, uint32(v9)+4))
	v197 = v194
	goto L50
L58:
	;
	v224 = int32(2)
	v229 = *(*int32)(unsafe.Add(mBase, uint32(v204+v222<<(uint(v224)%32)+int32(4))))
	v153 = v229 + v223<<(uint(v224)%32)
	v154 = int32(1)
	goto L44
L59:
	;
	v213 = v201
	goto L61
L60:
	;
	v213 = v198 << (uint(v209) % 32)
	goto L61
L61:
	;
	if v199 < v213 {
		v222 = v205
		v223 = v199
		goto L58
	} else {
		goto L62
	}
L62:
	;
	if v205 != 0 {
		v242 = v201
		goto L47
	} else {
		goto L63
	}
L63:
	;
	v215 = *(*int32)(unsafe.Add(mBase, uint32(v204)+20))
	if v215 == int32(-1) {
		v242 = v201
		goto L47
	} else {
		goto L64
	}
L64:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v9)+4)) = int64(4294967296)
	v222 = int32(1)
	v223 = int32(0)
	goto L58
L65:
	;
	v238 = *(*int32)(unsafe.Add(mBase, uint32(v234)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v146))) = v238
	v242 = v234
	goto L47
L66:
	;
	goto L33
L67:
	;
	F_setDeferredArrayLen(m, l0, v7, v248)
	mBase = m.M
	v254 = m.ExcPending
	if v254 != 0 {
		goto L1
	} else {
		goto L68
	}
L68:
	;
	return
}
func F_addReplyError(m *base.Module, l0 int32, l1 int32) {
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
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v74 int32
	_ = v74
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v85 int32
	_ = v85
	var v91 int32
	_ = v91
	var v94 int32
	_ = v94
	var v100 int32
	_ = v100
	var v104 int32
	_ = v104
	var v106 int32
	_ = v106
	var v114 int32
	_ = v114
	var v117 int32
	_ = v117
	if l1&int32(3) == int32(0) {
		v24 = l1
		goto L3
	} else {
		goto L4
	}
L1:
	;
	F_addReplyErrorLength(m, l0, l1, v57)
	mBase = m.M
	v59 = m.ExcPending
	if v59 != 0 {
		goto L17
	} else {
		goto L18
	}
L2:
	;
	v57 = v49 - l1
	goto L1
L3:
	;
	v28 = v24
	goto L11
L4:
	;
	v10 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	if v10 != 0 {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	v13 = l1
	goto L7
L6:
	;
	v57 = l1 - l1
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
	return
L18:
	;
	if l1&int32(3) == int32(0) {
		v81 = l1
		goto L21
	} else {
		goto L22
	}
L19:
	;
	F_afterErrorReply(m, l0, l1, v114, int32(0))
	mBase = m.M
	v117 = m.ExcPending
	if v117 != 0 {
		goto L17
	} else {
		goto L35
	}
L20:
	;
	v114 = v106 - l1
	goto L19
L21:
	;
	v85 = v81
	goto L29
L22:
	;
	v67 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	if v67 != 0 {
		goto L23
	} else {
		goto L24
	}
L23:
	;
	v70 = l1
	goto L25
L24:
	;
	v114 = l1 - l1
	goto L19
L25:
	;
	v74 = v70 + int32(1)
	if v74&int32(3) == int32(0) {
		v81 = v74
		goto L21
	} else {
		goto L27
	}
L27:
	;
	v79 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v74))))
	if v79 != 0 {
		v70 = v74
		goto L25
	} else {
		goto L28
	}
L28:
	;
	v106 = v74
	goto L20
L29:
	;
	v91 = *(*int32)(unsafe.Add(mBase, uint32(v85)))
	v94 = int32(-2139062144)
	if (int32(16843008)-v91|v91)&v94 == v94 {
		v85 = v85 + int32(4)
		goto L29
	} else {
		goto L31
	}
L30:
	;
	v100 = v85
	goto L32
L31:
	;
	goto L30
L32:
	;
	v104 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v100))))
	if v104 != 0 {
		v100 = v100 + int32(1)
		goto L32
	} else {
		goto L34
	}
L33:
	;
	v106 = v100
	goto L20
L34:
	;
	goto L33
L35:
	;
	return
}
func F_addReplyErrorExpireTime(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v12 int32
	_ = v12
	v3 = m.G0
	v5 = v3 - int32(16)
	m.G0 = v5
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(v7)+140))
	*(*int32)(unsafe.Add(mBase, uint32(v5))) = v8
	F_addReplyErrorFormat(m, l0, int32(_a_F_addReplyErrorExpireTime_0), v5)
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		return
	} else {
		m.G0 = v5 + int32(16)
		return
	}
}
func F_addReplyErrorSdsSafe(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v13 int32
	_ = v13
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v40 int32
	_ = v40
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v54 int32
	_ = v54
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v78 int32
	_ = v78
	var v91 int32
	_ = v91
	v13 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1+int32(-1)))))
	switch v13 & int32(7) {
	case 0:
		goto L8
	case 1:
		goto L7
	case 2:
		goto L6
	case 3:
		goto L5
	case 4:
		goto L4
	default:
		goto L2
	}
L1:
	;
	F_addReplyErrorSdsEx(m, l0, l1, int32(0))
	mBase = m.M
	v91 = m.ExcPending
	if v91 != 0 {
		goto L20
	} else {
		goto L21
	}
L2:
	;
	goto L1
L3:
	;
	if v30 == int32(0) {
		goto L2
	} else {
		goto L9
	}
L4:
	;
	v29 = *(*int32)(unsafe.Add(mBase, uint32(l1+int32(-17))))
	v30 = v29
	goto L3
L5:
	;
	v26 = *(*int32)(unsafe.Add(mBase, uint32(l1+int32(-9))))
	v30 = v26
	goto L3
L6:
	;
	v23 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1+int32(-5)))))
	v30 = v23
	goto L3
L7:
	;
	v20 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1+int32(-3)))))
	v30 = v20
	goto L3
L8:
	;
	v30 = int32(base.Ui32(v13) >> (uint(int32(3)) % 32))
	goto L3
L9:
	;
	v40 = int32(0)
	goto L10
L10:
	;
	goto L13
L11:
	;
	goto L2
L12:
	;
	v78 = v40 + int32(1)
	if v78 != v30 {
		v40 = v78
		goto L10
	} else {
		goto L19
	}
L13:
	;
	v46 = l1 + v40
	v47 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v46))))
	v54 = int32(0)
	goto L14
L14:
	;
	v60 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v54)+uint32(_c_F_addReplyErrorSdsSafe[0]))))
	if v47&int32(255) != v60 {
		goto L16
	} else {
		goto L17
	}
L15:
	;
	goto L12
L16:
	;
	v66 = v54 + int32(1)
	if v66 != int32(2) {
		v54 = v66
		goto L14
	} else {
		goto L18
	}
L17:
	;
	v63 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v54)+uint32(_c_F_addReplyErrorSdsSafe[1]))))
	*(*uint8)(unsafe.Add(mBase, uint32(v46))) = uint8(v63)
	goto L12
L18:
	;
	goto L15
L19:
	;
	goto L11
L20:
	;
	return
L21:
	;
	return
}
func F_addReplyLongLong(m *base.Module, l0 int32, l1 int64) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v13 int32
	_ = v13
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
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int64
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v67 int32
	_ = v67
	var v71 int32
	_ = v71
	var v76 int32
	_ = v76
	v5 = m.G0
	v7 = v5 - int32(128)
	m.G0 = v7
	if base.Ui64(int64(1)) < base.Ui64(l1) {
		v20 = F_prepareClientToWrite(m, l0)
		mBase = m.M
		v21 = m.ExcPending
		if v21 != 0 {
			return
		} else {
			if v20 != 0 {
				m.G0 = v7 + int32(128)
				return
			} else {
				v22 = int32(58)
				*(*uint8)(unsafe.Add(mBase, uint32(v7))) = uint8(v22)
				v25 = v7 | int32(1)
				if l1 <= int64(-1) {
					v35 = int32(45)
					*(*uint8)(unsafe.Add(mBase, uint32(v25))) = uint8(v35)
					v39 = int32(1)
					v44 = v25 + v39
					v45 = int32(126)
					v46 = int64(0) - l1
					v47 = v39
				} else {
					v44 = v25
					v45 = int32(127)
					v46 = l1
					v47 = int32(0)
				}
				v48 = F_ull2string(m, v44, v45, v46)
				mBase = m.M
				if v48 == int32(0) {
					v67 = int32(0)
				} else {
					v67 = v48 + v47
				}
				v71 = int32(2573)
				*(*uint16)(unsafe.Add(mBase, uint32(v67+v7+int32(1)))) = uint16(v71)
				F__addReplyToBufferOrList(m, l0, v7, v67+int32(3))
				mBase = m.M
				v76 = m.ExcPending
				if v76 != 0 {
					return
				} else {
					m.G0 = v7 + int32(128)
					return
				}
			}
		}
	} else {
		switch base.I32_wrap_i64(l1) {
		default:
			v13 = *(*int32)(unsafe.Add(mBase, _c_F_addReplyLongLong[0]))
			F_addReply(m, l0, v13)
			mBase = m.M
			v15 = m.ExcPending
			if v15 != 0 {
				return
			} else {
				m.G0 = v7 + int32(128)
				return
			}
		case 1:
			v17 = *(*int32)(unsafe.Add(mBase, _c_F_addReplyLongLong[1]))
			F_addReply(m, l0, v17)
			mBase = m.M
			v19 = m.ExcPending
			if v19 != 0 {
				return
			} else {
				m.G0 = v7 + int32(128)
				return
			}
		}
	}
}
func F_addReplyPubsubPatMessage(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v10 int32
	_ = v10
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+204))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+204)) = v6 | int32(131072)
	v10 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+224)))
	if v10 != int32(2) {
		F_addReplyPushLen(m, l0, int32(4))
		mBase = m.M
		v19 = m.ExcPending
		if v19 != 0 {
			return
		} else {
			v21 = *(*int32)(unsafe.Add(mBase, _c_F_addReplyPubsubPatMessage[0]))
			F_addReply(m, l0, v21)
			mBase = m.M
			v23 = m.ExcPending
			if v23 != 0 {
				return
			} else {
				F_addReplyBulk(m, l0, l1)
				mBase = m.M
				v25 = m.ExcPending
				if v25 != 0 {
					return
				} else {
					F_addReplyBulk(m, l0, l2)
					mBase = m.M
					v27 = m.ExcPending
					if v27 != 0 {
						return
					} else {
						F_addReplyBulk(m, l0, l3)
						mBase = m.M
						v29 = m.ExcPending
						if v29 != 0 {
							return
						} else {
							if v6&int32(131072) != 0 {
							} else {
								v32 = *(*int32)(unsafe.Add(mBase, uint32(l0)+204))
								*(*int32)(unsafe.Add(mBase, uint32(l0)+204)) = v32 & int32(-131073)
							}
							return
						}
					}
				}
			}
		}
	} else {
		v14 = *(*int32)(unsafe.Add(mBase, _c_F_addReplyPubsubPatMessage[1]))
		F_addReply(m, l0, v14)
		mBase = m.M
		v16 = m.ExcPending
		if v16 != 0 {
			return
		} else {
			v21 = *(*int32)(unsafe.Add(mBase, _c_F_addReplyPubsubPatMessage[0]))
			F_addReply(m, l0, v21)
			mBase = m.M
			v23 = m.ExcPending
			if v23 != 0 {
				return
			} else {
				F_addReplyBulk(m, l0, l1)
				mBase = m.M
				v25 = m.ExcPending
				if v25 != 0 {
					return
				} else {
					F_addReplyBulk(m, l0, l2)
					mBase = m.M
					v27 = m.ExcPending
					if v27 != 0 {
						return
					} else {
						F_addReplyBulk(m, l0, l3)
						mBase = m.M
						v29 = m.ExcPending
						if v29 != 0 {
							return
						} else {
							if v6&int32(131072) != 0 {
							} else {
								v32 = *(*int32)(unsafe.Add(mBase, uint32(l0)+204))
								*(*int32)(unsafe.Add(mBase, uint32(l0)+204)) = v32 & int32(-131073)
							}
							return
						}
					}
				}
			}
		}
	}
}
func F_addReplyPubsubPatSubscribed(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v8 int32
	_ = v8
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+204))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+204)) = v4 | int32(131072)
	v8 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+224)))
	if v8 != int32(2) {
		F_addReplyPushLen(m, l0, int32(3))
		mBase = m.M
		v17 = m.ExcPending
		if v17 != 0 {
			return
		} else {
			v19 = *(*int32)(unsafe.Add(mBase, _c_F_addReplyPubsubPatSubscribed[0]))
			F_addReply(m, l0, v19)
			mBase = m.M
			v21 = m.ExcPending
			if v21 != 0 {
				return
			} else {
				F_addReplyBulk(m, l0, l1)
				mBase = m.M
				v23 = m.ExcPending
				if v23 != 0 {
					return
				} else {
					v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+100))
					v25 = *(*int32)(unsafe.Add(mBase, uint32(v24)))
					v26 = *(*int32)(unsafe.Add(mBase, uint32(v25)+20))
					v27 = *(*int32)(unsafe.Add(mBase, uint32(v25)+16))
					v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)+100))
					v30 = *(*int32)(unsafe.Add(mBase, uint32(v29)+4))
					v31 = *(*int32)(unsafe.Add(mBase, uint32(v30)+20))
					v32 = *(*int32)(unsafe.Add(mBase, uint32(v30)+16))
					F_addReplyLongLong(m, l0, base.I64_extend_i32_s(v26+v27+(v31+v32)))
					mBase = m.M
					v37 = m.ExcPending
					if v37 != 0 {
						return
					} else {
						if v4&int32(131072) != 0 {
						} else {
							v40 = *(*int32)(unsafe.Add(mBase, uint32(l0)+204))
							*(*int32)(unsafe.Add(mBase, uint32(l0)+204)) = v40 & int32(-131073)
						}
						return
					}
				}
			}
		}
	} else {
		v12 = *(*int32)(unsafe.Add(mBase, _c_F_addReplyPubsubPatSubscribed[1]))
		F_addReply(m, l0, v12)
		mBase = m.M
		v14 = m.ExcPending
		if v14 != 0 {
			return
		} else {
			v19 = *(*int32)(unsafe.Add(mBase, _c_F_addReplyPubsubPatSubscribed[0]))
			F_addReply(m, l0, v19)
			mBase = m.M
			v21 = m.ExcPending
			if v21 != 0 {
				return
			} else {
				F_addReplyBulk(m, l0, l1)
				mBase = m.M
				v23 = m.ExcPending
				if v23 != 0 {
					return
				} else {
					v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+100))
					v25 = *(*int32)(unsafe.Add(mBase, uint32(v24)))
					v26 = *(*int32)(unsafe.Add(mBase, uint32(v25)+20))
					v27 = *(*int32)(unsafe.Add(mBase, uint32(v25)+16))
					v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)+100))
					v30 = *(*int32)(unsafe.Add(mBase, uint32(v29)+4))
					v31 = *(*int32)(unsafe.Add(mBase, uint32(v30)+20))
					v32 = *(*int32)(unsafe.Add(mBase, uint32(v30)+16))
					F_addReplyLongLong(m, l0, base.I64_extend_i32_s(v26+v27+(v31+v32)))
					mBase = m.M
					v37 = m.ExcPending
					if v37 != 0 {
						return
					} else {
						if v4&int32(131072) != 0 {
						} else {
							v40 = *(*int32)(unsafe.Add(mBase, uint32(l0)+204))
							*(*int32)(unsafe.Add(mBase, uint32(l0)+204)) = v40 & int32(-131073)
						}
						return
					}
				}
			}
		}
	}
}
func F_addReplyPubsubPatUnsubscribed(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v8 int32
	_ = v8
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v25 int32
	_ = v25
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
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+204))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+204)) = v4 | int32(131072)
	v8 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+224)))
	if v8 != int32(2) {
		F_addReplyPushLen(m, l0, int32(3))
		mBase = m.M
		v17 = m.ExcPending
		if v17 != 0 {
			return
		} else {
			v19 = *(*int32)(unsafe.Add(mBase, _c_F_addReplyPubsubPatUnsubscribed[0]))
			F_addReply(m, l0, v19)
			mBase = m.M
			v21 = m.ExcPending
			if v21 != 0 {
				return
			} else {
				if l1 == int32(0) {
					F_addReplyNull(m, l0)
					mBase = m.M
					v27 = m.ExcPending
					if v27 != 0 {
						return
					} else {
						v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+100))
						v29 = *(*int32)(unsafe.Add(mBase, uint32(v28)))
						v30 = *(*int32)(unsafe.Add(mBase, uint32(v29)+20))
						v31 = *(*int32)(unsafe.Add(mBase, uint32(v29)+16))
						v33 = *(*int32)(unsafe.Add(mBase, uint32(l0)+100))
						v34 = *(*int32)(unsafe.Add(mBase, uint32(v33)+4))
						v35 = *(*int32)(unsafe.Add(mBase, uint32(v34)+20))
						v36 = *(*int32)(unsafe.Add(mBase, uint32(v34)+16))
						F_addReplyLongLong(m, l0, base.I64_extend_i32_s(v30+v31+(v35+v36)))
						mBase = m.M
						v41 = m.ExcPending
						if v41 != 0 {
							return
						} else {
							if v4&int32(131072) != 0 {
							} else {
								v44 = *(*int32)(unsafe.Add(mBase, uint32(l0)+204))
								*(*int32)(unsafe.Add(mBase, uint32(l0)+204)) = v44 & int32(-131073)
							}
							return
						}
					}
				} else {
					F_addReplyBulk(m, l0, l1)
					mBase = m.M
					v25 = m.ExcPending
					if v25 != 0 {
						return
					} else {
						v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+100))
						v29 = *(*int32)(unsafe.Add(mBase, uint32(v28)))
						v30 = *(*int32)(unsafe.Add(mBase, uint32(v29)+20))
						v31 = *(*int32)(unsafe.Add(mBase, uint32(v29)+16))
						v33 = *(*int32)(unsafe.Add(mBase, uint32(l0)+100))
						v34 = *(*int32)(unsafe.Add(mBase, uint32(v33)+4))
						v35 = *(*int32)(unsafe.Add(mBase, uint32(v34)+20))
						v36 = *(*int32)(unsafe.Add(mBase, uint32(v34)+16))
						F_addReplyLongLong(m, l0, base.I64_extend_i32_s(v30+v31+(v35+v36)))
						mBase = m.M
						v41 = m.ExcPending
						if v41 != 0 {
							return
						} else {
							if v4&int32(131072) != 0 {
							} else {
								v44 = *(*int32)(unsafe.Add(mBase, uint32(l0)+204))
								*(*int32)(unsafe.Add(mBase, uint32(l0)+204)) = v44 & int32(-131073)
							}
							return
						}
					}
				}
			}
		}
	} else {
		v12 = *(*int32)(unsafe.Add(mBase, _c_F_addReplyPubsubPatUnsubscribed[1]))
		F_addReply(m, l0, v12)
		mBase = m.M
		v14 = m.ExcPending
		if v14 != 0 {
			return
		} else {
			v19 = *(*int32)(unsafe.Add(mBase, _c_F_addReplyPubsubPatUnsubscribed[0]))
			F_addReply(m, l0, v19)
			mBase = m.M
			v21 = m.ExcPending
			if v21 != 0 {
				return
			} else {
				if l1 == int32(0) {
					F_addReplyNull(m, l0)
					mBase = m.M
					v27 = m.ExcPending
					if v27 != 0 {
						return
					} else {
						v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+100))
						v29 = *(*int32)(unsafe.Add(mBase, uint32(v28)))
						v30 = *(*int32)(unsafe.Add(mBase, uint32(v29)+20))
						v31 = *(*int32)(unsafe.Add(mBase, uint32(v29)+16))
						v33 = *(*int32)(unsafe.Add(mBase, uint32(l0)+100))
						v34 = *(*int32)(unsafe.Add(mBase, uint32(v33)+4))
						v35 = *(*int32)(unsafe.Add(mBase, uint32(v34)+20))
						v36 = *(*int32)(unsafe.Add(mBase, uint32(v34)+16))
						F_addReplyLongLong(m, l0, base.I64_extend_i32_s(v30+v31+(v35+v36)))
						mBase = m.M
						v41 = m.ExcPending
						if v41 != 0 {
							return
						} else {
							if v4&int32(131072) != 0 {
							} else {
								v44 = *(*int32)(unsafe.Add(mBase, uint32(l0)+204))
								*(*int32)(unsafe.Add(mBase, uint32(l0)+204)) = v44 & int32(-131073)
							}
							return
						}
					}
				} else {
					F_addReplyBulk(m, l0, l1)
					mBase = m.M
					v25 = m.ExcPending
					if v25 != 0 {
						return
					} else {
						v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+100))
						v29 = *(*int32)(unsafe.Add(mBase, uint32(v28)))
						v30 = *(*int32)(unsafe.Add(mBase, uint32(v29)+20))
						v31 = *(*int32)(unsafe.Add(mBase, uint32(v29)+16))
						v33 = *(*int32)(unsafe.Add(mBase, uint32(l0)+100))
						v34 = *(*int32)(unsafe.Add(mBase, uint32(v33)+4))
						v35 = *(*int32)(unsafe.Add(mBase, uint32(v34)+20))
						v36 = *(*int32)(unsafe.Add(mBase, uint32(v34)+16))
						F_addReplyLongLong(m, l0, base.I64_extend_i32_s(v30+v31+(v35+v36)))
						mBase = m.M
						v41 = m.ExcPending
						if v41 != 0 {
							return
						} else {
							if v4&int32(131072) != 0 {
							} else {
								v44 = *(*int32)(unsafe.Add(mBase, uint32(l0)+204))
								*(*int32)(unsafe.Add(mBase, uint32(l0)+204)) = v44 & int32(-131073)
							}
							return
						}
					}
				}
			}
		}
	}
}
func F_addReplyPubsubUnsubscribed(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v9 int32
	_ = v9
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
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+204))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+204)) = v5 | int32(131072)
	v9 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+224)))
	if v9 != int32(2) {
		F_addReplyPushLen(m, l0, int32(3))
		mBase = m.M
		v18 = m.ExcPending
		if v18 != 0 {
			return
		} else {
			v19 = *(*int32)(unsafe.Add(mBase, uint32(l2)+20))
			v20 = *(*int32)(unsafe.Add(mBase, uint32(v19)))
			F_addReply(m, l0, v20)
			mBase = m.M
			v22 = m.ExcPending
			if v22 != 0 {
				return
			} else {
				if l1 == int32(0) {
					F_addReplyNull(m, l0)
					mBase = m.M
					v28 = m.ExcPending
					if v28 != 0 {
						return
					} else {
						v29 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
						v30 = m.T0[v29].(func(*base.Module, int32) int32)(m, l0)
						mBase = m.M
						v31 = m.ExcPending
						if v31 != 0 {
							return
						} else {
							F_addReplyLongLong(m, l0, base.I64_extend_i32_s(v30))
							mBase = m.M
							v34 = m.ExcPending
							if v34 != 0 {
								return
							} else {
								if v5&int32(131072) != 0 {
								} else {
									v37 = *(*int32)(unsafe.Add(mBase, uint32(l0)+204))
									*(*int32)(unsafe.Add(mBase, uint32(l0)+204)) = v37 & int32(-131073)
								}
								return
							}
						}
					}
				} else {
					F_addReplyBulk(m, l0, l1)
					mBase = m.M
					v26 = m.ExcPending
					if v26 != 0 {
						return
					} else {
						v29 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
						v30 = m.T0[v29].(func(*base.Module, int32) int32)(m, l0)
						mBase = m.M
						v31 = m.ExcPending
						if v31 != 0 {
							return
						} else {
							F_addReplyLongLong(m, l0, base.I64_extend_i32_s(v30))
							mBase = m.M
							v34 = m.ExcPending
							if v34 != 0 {
								return
							} else {
								if v5&int32(131072) != 0 {
								} else {
									v37 = *(*int32)(unsafe.Add(mBase, uint32(l0)+204))
									*(*int32)(unsafe.Add(mBase, uint32(l0)+204)) = v37 & int32(-131073)
								}
								return
							}
						}
					}
				}
			}
		}
	} else {
		v13 = *(*int32)(unsafe.Add(mBase, _c_F_addReplyPubsubUnsubscribed[0]))
		F_addReply(m, l0, v13)
		mBase = m.M
		v15 = m.ExcPending
		if v15 != 0 {
			return
		} else {
			v19 = *(*int32)(unsafe.Add(mBase, uint32(l2)+20))
			v20 = *(*int32)(unsafe.Add(mBase, uint32(v19)))
			F_addReply(m, l0, v20)
			mBase = m.M
			v22 = m.ExcPending
			if v22 != 0 {
				return
			} else {
				if l1 == int32(0) {
					F_addReplyNull(m, l0)
					mBase = m.M
					v28 = m.ExcPending
					if v28 != 0 {
						return
					} else {
						v29 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
						v30 = m.T0[v29].(func(*base.Module, int32) int32)(m, l0)
						mBase = m.M
						v31 = m.ExcPending
						if v31 != 0 {
							return
						} else {
							F_addReplyLongLong(m, l0, base.I64_extend_i32_s(v30))
							mBase = m.M
							v34 = m.ExcPending
							if v34 != 0 {
								return
							} else {
								if v5&int32(131072) != 0 {
								} else {
									v37 = *(*int32)(unsafe.Add(mBase, uint32(l0)+204))
									*(*int32)(unsafe.Add(mBase, uint32(l0)+204)) = v37 & int32(-131073)
								}
								return
							}
						}
					}
				} else {
					F_addReplyBulk(m, l0, l1)
					mBase = m.M
					v26 = m.ExcPending
					if v26 != 0 {
						return
					} else {
						v29 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
						v30 = m.T0[v29].(func(*base.Module, int32) int32)(m, l0)
						mBase = m.M
						v31 = m.ExcPending
						if v31 != 0 {
							return
						} else {
							F_addReplyLongLong(m, l0, base.I64_extend_i32_s(v30))
							mBase = m.M
							v34 = m.ExcPending
							if v34 != 0 {
								return
							} else {
								if v5&int32(131072) != 0 {
								} else {
									v37 = *(*int32)(unsafe.Add(mBase, uint32(l0)+204))
									*(*int32)(unsafe.Add(mBase, uint32(l0)+204)) = v37 & int32(-131073)
								}
								return
							}
						}
					}
				}
			}
		}
	}
}
func F_addReplySentinelValkeyInstance(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
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
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int64
	_ = v35
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
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
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
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
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v207 int32
	_ = v207
	var v214 int32
	_ = v214
	var v217 int32
	_ = v217
	var v220 int32
	_ = v220
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v239 int32
	_ = v239
	var v244 int32
	_ = v244
	var v247 int32
	_ = v247
	var v250 int32
	_ = v250
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v260 int32
	_ = v260
	var v264 int32
	_ = v264
	var v267 int32
	_ = v267
	var v272 int32
	_ = v272
	var v275 int32
	_ = v275
	var v278 int32
	_ = v278
	var v281 int32
	_ = v281
	var v282 int32
	_ = v282
	var v283 int32
	_ = v283
	var v285 int32
	_ = v285
	var v286 int32
	_ = v286
	var v287 int32
	_ = v287
	var v289 int32
	_ = v289
	var v291 int32
	_ = v291
	var v293 int32
	_ = v293
	var v297 int32
	_ = v297
	var v299 int32
	_ = v299
	var v302 int32
	_ = v302
	var v331 int32
	_ = v331
	var v333 int32
	_ = v333
	var v336 int32
	_ = v336
	var v337 int32
	_ = v337
	var v338 int64
	_ = v338
	var v340 int32
	_ = v340
	var v343 int32
	_ = v343
	var v344 int32
	_ = v344
	var v345 int64
	_ = v345
	var v347 int32
	_ = v347
	var v349 int32
	_ = v349
	var v356 int32
	_ = v356
	var v358 int32
	_ = v358
	var v365 int32
	_ = v365
	var v366 int32
	_ = v366
	var v368 int32
	_ = v368
	var v370 int32
	_ = v370
	var v374 int32
	_ = v374
	var v375 int32
	_ = v375
	var v376 int64
	_ = v376
	var v382 int64
	_ = v382
	var v383 int32
	_ = v383
	var v384 int64
	_ = v384
	var v386 int64
	_ = v386
	var v388 int32
	_ = v388
	var v391 int32
	_ = v391
	var v392 int64
	_ = v392
	var v393 int32
	_ = v393
	var v394 int64
	_ = v394
	var v397 int32
	_ = v397
	var v400 int32
	_ = v400
	var v401 int64
	_ = v401
	var v402 int32
	_ = v402
	var v403 int64
	_ = v403
	var v406 int32
	_ = v406
	var v407 int32
	_ = v407
	var v414 int32
	_ = v414
	var v415 int64
	_ = v415
	var v416 int64
	_ = v416
	var v419 int32
	_ = v419
	var v422 int32
	_ = v422
	var v423 int32
	_ = v423
	var v424 int32
	_ = v424
	var v431 int32
	_ = v431
	var v432 int64
	_ = v432
	var v433 int64
	_ = v433
	var v436 int32
	_ = v436
	var v439 int32
	_ = v439
	var v442 int32
	_ = v442
	var v443 int64
	_ = v443
	var v445 int32
	_ = v445
	var v446 int32
	_ = v446
	var v453 int32
	_ = v453
	var v454 int64
	_ = v454
	var v460 int64
	_ = v460
	var v461 int64
	_ = v461
	var v463 int64
	_ = v463
	var v465 int32
	_ = v465
	var v468 int32
	_ = v468
	var v471 int32
	_ = v471
	var v474 int32
	_ = v474
	var v476 int32
	_ = v476
	var v479 int32
	_ = v479
	var v480 int64
	_ = v480
	var v481 int64
	_ = v481
	var v484 int32
	_ = v484
	var v487 int32
	_ = v487
	var v488 int32
	_ = v488
	var v489 int32
	_ = v489
	var v497 int32
	_ = v497
	var v498 int64
	_ = v498
	var v500 int32
	_ = v500
	var v503 int32
	_ = v503
	var v504 int32
	_ = v504
	var v505 int32
	_ = v505
	var v506 int32
	_ = v506
	var v510 int32
	_ = v510
	var v513 int32
	_ = v513
	var v514 int32
	_ = v514
	var v515 int32
	_ = v515
	var v516 int32
	_ = v516
	var v520 int32
	_ = v520
	var v523 int32
	_ = v523
	var v524 int64
	_ = v524
	var v526 int32
	_ = v526
	var v529 int32
	_ = v529
	var v530 int64
	_ = v530
	var v532 int32
	_ = v532
	var v535 int32
	_ = v535
	var v536 int64
	_ = v536
	var v538 int32
	_ = v538
	var v539 int32
	_ = v539
	var v544 int32
	_ = v544
	var v545 int32
	_ = v545
	var v547 int32
	_ = v547
	var v550 int32
	_ = v550
	var v551 int32
	_ = v551
	var v556 int32
	_ = v556
	var v557 int32
	_ = v557
	var v559 int32
	_ = v559
	var v562 int32
	_ = v562
	var v564 int32
	_ = v564
	var v571 int32
	_ = v571
	var v572 int64
	_ = v572
	var v574 int32
	_ = v574
	var v577 int32
	_ = v577
	var v580 int32
	_ = v580
	var v581 int32
	_ = v581
	var v583 int32
	_ = v583
	var v586 int32
	_ = v586
	var v587 int32
	_ = v587
	var v589 int32
	_ = v589
	var v591 int32
	_ = v591
	var v594 int32
	_ = v594
	var v595 int64
	_ = v595
	var v597 int32
	_ = v597
	var v600 int32
	_ = v600
	var v601 int64
	_ = v601
	var v603 int32
	_ = v603
	var v606 int32
	_ = v606
	var v607 int64
	_ = v607
	var v609 int32
	_ = v609
	var v612 int32
	_ = v612
	var v613 int64
	_ = v613
	var v615 int32
	_ = v615
	var v618 int32
	_ = v618
	var v619 int32
	_ = v619
	var v620 int32
	_ = v620
	var v627 int32
	_ = v627
	var v628 int64
	_ = v628
	var v629 int64
	_ = v629
	var v632 int32
	_ = v632
	var v635 int32
	_ = v635
	var v636 int32
	_ = v636
	var v638 int32
	_ = v638
	var v640 int32
	_ = v640
	var v643 int32
	_ = v643
	var v644 int64
	_ = v644
	var v646 int32
	_ = v646
	var v649 int32
	_ = v649
	var v652 int32
	_ = v652
	v7 = F_sdsempty(m)
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
	v9 = F_addReplyDeferredLen(m, l0)
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	F_addReplyBulkCString(m, l0, int32(_a_F_addReplySentinelValkeyInstance_0))
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	F_addReplyBulkCString(m, l0, v14)
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		goto L1
	} else {
		goto L5
	}
L5:
	;
	F_addReplyBulkCString(m, l0, int32(_a_F_addReplySentinelValkeyInstance_1))
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		goto L1
	} else {
		goto L6
	}
L6:
	;
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
	v21 = int32(0)
	v22 = *(*int32)(unsafe.Add(mBase, _c_F_addReplySentinelValkeyInstance[0]))
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v20+base.B2i32(v22 == v21)<<(uint(int32(2))%32))))
	F_addReplyBulkCString(m, l0, v28)
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		goto L1
	} else {
		goto L7
	}
L7:
	;
	F_addReplyBulkCString(m, l0, int32(_a_F_addReplySentinelValkeyInstance_2))
	mBase = m.M
	v33 = m.ExcPending
	if v33 != 0 {
		goto L1
	} else {
		goto L8
	}
L8:
	;
	v34 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
	v35 = int64(*(*int32)(unsafe.Add(mBase, uint32(v34)+8)))
	F_addReplyBulkLongLong(m, l0, v35)
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		goto L1
	} else {
		goto L9
	}
L9:
	;
	F_addReplyBulkCString(m, l0, int32(_a_F_addReplySentinelValkeyInstance_3))
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
		goto L1
	} else {
		goto L10
	}
L10:
	;
	v41 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	if v41 != 0 {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v43 = v41
	goto L13
L12:
	;
	v43 = int32(_a_F_addReplySentinelValkeyInstance_4)
	goto L13
L13:
	;
	F_addReplyBulkCString(m, l0, v43)
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L1
	} else {
		goto L14
	}
L14:
	;
	F_addReplyBulkCString(m, l0, int32(_a_F_addReplySentinelValkeyInstance_5))
	mBase = m.M
	v48 = m.ExcPending
	if v48 != 0 {
		goto L1
	} else {
		goto L15
	}
L15:
	;
	v49 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	if v49&int32(8) == int32(0) {
		v58 = v7
		v59 = v49
		goto L16
	} else {
		goto L17
	}
L16:
	;
	if v59&int32(16) == int32(0) {
		v68 = v58
		v69 = v59
		goto L19
	} else {
		goto L20
	}
L17:
	;
	v55 = F_sdscat(m, v7, int32(_a_F_addReplySentinelValkeyInstance_38))
	mBase = m.M
	v56 = m.ExcPending
	if v56 != 0 {
		goto L1
	} else {
		goto L18
	}
L18:
	;
	v57 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v58 = v55
	v59 = v57
	goto L16
L19:
	;
	if v69&int32(1) == int32(0) {
		v78 = v68
		v79 = v69
		goto L22
	} else {
		goto L23
	}
L20:
	;
	v65 = F_sdscat(m, v58, int32(_a_F_addReplySentinelValkeyInstance_39))
	mBase = m.M
	v66 = m.ExcPending
	if v66 != 0 {
		goto L1
	} else {
		goto L21
	}
L21:
	;
	v67 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v68 = v65
	v69 = v67
	goto L19
L22:
	;
	if v79&int32(2) == int32(0) {
		v88 = v78
		v89 = v79
		goto L25
	} else {
		goto L26
	}
L23:
	;
	v75 = F_sdscat(m, v68, int32(_a_F_addReplySentinelValkeyInstance_40))
	mBase = m.M
	v76 = m.ExcPending
	if v76 != 0 {
		goto L1
	} else {
		goto L24
	}
L24:
	;
	v77 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v78 = v75
	v79 = v77
	goto L22
L25:
	;
	if v89&int32(4) == int32(0) {
		v97 = v88
		goto L28
	} else {
		goto L29
	}
L26:
	;
	v85 = F_sdscat(m, v78, int32(_a_F_addReplySentinelValkeyInstance_41))
	mBase = m.M
	v86 = m.ExcPending
	if v86 != 0 {
		goto L1
	} else {
		goto L27
	}
L27:
	;
	v87 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v88 = v85
	v89 = v87
	goto L25
L28:
	;
	v98 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
	v99 = *(*int32)(unsafe.Add(mBase, uint32(v98)+4))
	if v99 == int32(0) {
		v105 = v97
		goto L31
	} else {
		goto L32
	}
L29:
	;
	v95 = F_sdscat(m, v88, int32(_a_F_addReplySentinelValkeyInstance_42))
	mBase = m.M
	v96 = m.ExcPending
	if v96 != 0 {
		goto L1
	} else {
		goto L30
	}
L30:
	;
	v97 = v95
	goto L28
L31:
	;
	v106 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	if v106&int32(32) == int32(0) {
		v115 = v105
		v116 = v106
		goto L34
	} else {
		goto L35
	}
L32:
	;
	v103 = F_sdscat(m, v97, int32(_a_F_addReplySentinelValkeyInstance_43))
	mBase = m.M
	v104 = m.ExcPending
	if v104 != 0 {
		goto L1
	} else {
		goto L33
	}
L33:
	;
	v105 = v103
	goto L31
L34:
	;
	if v116&int32(64) == int32(0) {
		v125 = v115
		v126 = v116
		goto L37
	} else {
		goto L38
	}
L35:
	;
	v112 = F_sdscat(m, v105, int32(_a_F_addReplySentinelValkeyInstance_44))
	mBase = m.M
	v113 = m.ExcPending
	if v113 != 0 {
		goto L1
	} else {
		goto L36
	}
L36:
	;
	v114 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v115 = v112
	v116 = v114
	goto L34
L37:
	;
	if v126&int32(128) == int32(0) {
		v135 = v125
		v136 = v126
		goto L40
	} else {
		goto L41
	}
L38:
	;
	v122 = F_sdscat(m, v115, int32(_a_F_addReplySentinelValkeyInstance_45))
	mBase = m.M
	v123 = m.ExcPending
	if v123 != 0 {
		goto L1
	} else {
		goto L39
	}
L39:
	;
	v124 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v125 = v122
	v126 = v124
	goto L37
L40:
	;
	if v136&int32(256) == int32(0) {
		v145 = v135
		v146 = v136
		goto L43
	} else {
		goto L44
	}
L41:
	;
	v132 = F_sdscat(m, v125, int32(_a_F_addReplySentinelValkeyInstance_46))
	mBase = m.M
	v133 = m.ExcPending
	if v133 != 0 {
		goto L1
	} else {
		goto L42
	}
L42:
	;
	v134 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v135 = v132
	v136 = v134
	goto L40
L43:
	;
	if v146&int32(512) == int32(0) {
		v155 = v145
		v156 = v146
		goto L46
	} else {
		goto L47
	}
L44:
	;
	v142 = F_sdscat(m, v135, int32(_a_F_addReplySentinelValkeyInstance_47))
	mBase = m.M
	v143 = m.ExcPending
	if v143 != 0 {
		goto L1
	} else {
		goto L45
	}
L45:
	;
	v144 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v145 = v142
	v146 = v144
	goto L43
L46:
	;
	if v156&int32(1024) == int32(0) {
		v165 = v155
		v166 = v156
		goto L49
	} else {
		goto L50
	}
L47:
	;
	v152 = F_sdscat(m, v145, int32(_a_F_addReplySentinelValkeyInstance_48))
	mBase = m.M
	v153 = m.ExcPending
	if v153 != 0 {
		goto L1
	} else {
		goto L48
	}
L48:
	;
	v154 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v155 = v152
	v156 = v154
	goto L46
L49:
	;
	if v166&int32(2048) == int32(0) {
		v175 = v165
		v176 = v166
		goto L52
	} else {
		goto L53
	}
L50:
	;
	v162 = F_sdscat(m, v155, int32(_a_F_addReplySentinelValkeyInstance_49))
	mBase = m.M
	v163 = m.ExcPending
	if v163 != 0 {
		goto L1
	} else {
		goto L51
	}
L51:
	;
	v164 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v165 = v162
	v166 = v164
	goto L49
L52:
	;
	if v176&int32(16384) == int32(0) {
		v185 = v175
		v186 = v176
		goto L55
	} else {
		goto L56
	}
L53:
	;
	v172 = F_sdscat(m, v165, int32(_a_F_addReplySentinelValkeyInstance_50))
	mBase = m.M
	v173 = m.ExcPending
	if v173 != 0 {
		goto L1
	} else {
		goto L54
	}
L54:
	;
	v174 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v175 = v172
	v176 = v174
	goto L52
L55:
	;
	if v186&int32(4096) == int32(0) {
		v195 = v185
		v196 = v186
		goto L58
	} else {
		goto L59
	}
L56:
	;
	v182 = F_sdscat(m, v175, int32(_a_F_addReplySentinelValkeyInstance_51))
	mBase = m.M
	v183 = m.ExcPending
	if v183 != 0 {
		goto L1
	} else {
		goto L57
	}
L57:
	;
	v184 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v185 = v182
	v186 = v184
	goto L55
L58:
	;
	if v196&int32(8192) == int32(0) {
		v204 = v195
		goto L61
	} else {
		goto L62
	}
L59:
	;
	v192 = F_sdscat(m, v185, int32(_a_F_addReplySentinelValkeyInstance_52))
	mBase = m.M
	v193 = m.ExcPending
	if v193 != 0 {
		goto L1
	} else {
		goto L60
	}
L60:
	;
	v194 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v195 = v192
	v196 = v194
	goto L58
L61:
	;
	v207 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v204+int32(-1)))))
	switch v207 & int32(7) {
	case 0:
		goto L70
	case 1:
		goto L69
	case 2:
		goto L68
	case 3:
		goto L67
	case 4:
		goto L66
	default:
		goto L64
	}
L62:
	;
	v202 = F_sdscat(m, v195, int32(_a_F_addReplySentinelValkeyInstance_53))
	mBase = m.M
	v203 = m.ExcPending
	if v203 != 0 {
		goto L1
	} else {
		goto L63
	}
L63:
	;
	v204 = v202
	goto L61
L64:
	;
	F_addReplyBulkCString(m, l0, v204)
	mBase = m.M
	v331 = m.ExcPending
	if v331 != 0 {
		goto L1
	} else {
		goto L106
	}
L65:
	;
	if v224 == int32(0) {
		goto L64
	} else {
		goto L71
	}
L66:
	;
	v223 = *(*int32)(unsafe.Add(mBase, uint32(v204+int32(-17))))
	v224 = v223
	goto L65
L67:
	;
	v220 = *(*int32)(unsafe.Add(mBase, uint32(v204+int32(-9))))
	v224 = v220
	goto L65
L68:
	;
	v217 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v204+int32(-5)))))
	v224 = v217
	goto L65
L69:
	;
	v214 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v204+int32(-3)))))
	v224 = v214
	goto L65
L70:
	;
	v224 = int32(base.Ui32(v207) >> (uint(int32(3)) % 32))
	goto L65
L71:
	;
	v236 = v204 + int32(-1)
	v237 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v236))))
	v239 = v237 & int32(7)
	switch v239 {
	case 0:
		goto L79
	case 1:
		goto L78
	case 2:
		goto L77
	case 3:
		goto L76
	case 4:
		goto L75
	default:
		goto L73
	}
L72:
	;
	goto L64
L73:
	;
	goto L72
L74:
	;
	if v254 == int32(0) {
		goto L73
	} else {
		goto L80
	}
L75:
	;
	v253 = *(*int32)(unsafe.Add(mBase, uint32(v204+int32(-17))))
	v254 = v253
	goto L74
L76:
	;
	v250 = *(*int32)(unsafe.Add(mBase, uint32(v204+int32(-9))))
	v254 = v250
	goto L74
L77:
	;
	v247 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v204+int32(-5)))))
	v254 = v247
	goto L74
L78:
	;
	v244 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v204+int32(-3)))))
	v254 = v244
	goto L74
L79:
	;
	v254 = int32(base.Ui32(v237) >> (uint(int32(3)) % 32))
	goto L74
L80:
	;
	v260 = int32(-1)&v254 + int32(-2)
	v264 = int32(0)&v254 + int32(0)
	v267 = v260 - v264 + int32(1)
	switch v239 {
	default:
		goto L86
	case 1:
		goto L85
	case 2:
		goto L84
	case 3:
		goto L83
	case 4:
		goto L82
	}
L81:
	;
	v283 = int32(0)
	v285 = base.B2i32(base.Ui32(v264) < base.Ui32(v282))
	if base.Ui32(v264) < base.Ui32(v282) {
		goto L88
	} else {
		goto L89
	}
L82:
	;
	v281 = *(*int32)(unsafe.Add(mBase, uint32(v204+int32(-17))))
	v282 = v281
	goto L81
L83:
	;
	v278 = *(*int32)(unsafe.Add(mBase, uint32(v204+int32(-9))))
	v282 = v278
	goto L81
L84:
	;
	v275 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v204+int32(-5)))))
	v282 = v275
	goto L81
L85:
	;
	v272 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v204+int32(-3)))))
	v282 = v272
	goto L81
L86:
	;
	v282 = int32(base.Ui32(v237) >> (uint(int32(3)) % 32))
	goto L81
L87:
	;
	v299 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v204+v293))) = uint8(v299)
	switch v239 {
	default:
		goto L105
	case 1:
		goto L104
	case 2:
		goto L103
	case 3:
		goto L102
	case 4:
		goto L101
	}
L88:
	;
	v286 = v264
	goto L90
L89:
	;
	v286 = v283
	goto L90
L90:
	;
	v287 = v282 - v286
	if base.Ui32(v267) < base.Ui32(v287) {
		goto L91
	} else {
		goto L92
	}
L91:
	;
	v289 = v267
	goto L93
L92:
	;
	v289 = v287
	goto L93
L93:
	;
	if v260 < v264 {
		goto L94
	} else {
		goto L95
	}
L94:
	;
	v291 = v283
	goto L96
L95:
	;
	v291 = v289
	goto L96
L96:
	;
	if base.Ui32(v264) < base.Ui32(v282) {
		goto L97
	} else {
		goto L98
	}
L97:
	;
	v293 = v291
	goto L99
L98:
	;
	v293 = int32(0)
	goto L99
L99:
	;
	if v293 == int32(0) {
		goto L87
	} else {
		goto L100
	}
L100:
	;
	v297 = F_memmove(m, v204, v204+v286, v293)
	mBase = m.M
	goto L87
L101:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v204+int32(-17)))) = base.I64_extend_i32_u(v293)
	goto L73
L102:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v204+int32(-9)))) = v293
	goto L72
L103:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v204+int32(-5)))) = uint16(v293)
	goto L72
L104:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v204+int32(-3)))) = uint8(v293)
	goto L72
L105:
	;
	v302 = v293 << (uint(int32(3)) % 32)
	*(*uint8)(unsafe.Add(mBase, uint32(v236))) = uint8(v302)
	goto L72
L106:
	;
	F_sdsfree(m, v204)
	mBase = m.M
	v333 = m.ExcPending
	if v333 != 0 {
		goto L1
	} else {
		goto L107
	}
L107:
	;
	F_addReplyBulkCString(m, l0, int32(_a_F_addReplySentinelValkeyInstance_6))
	mBase = m.M
	v336 = m.ExcPending
	if v336 != 0 {
		goto L1
	} else {
		goto L108
	}
L108:
	;
	v337 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
	v338 = int64(*(*int32)(unsafe.Add(mBase, uint32(v337)+8)))
	F_addReplyBulkLongLong(m, l0, v338)
	mBase = m.M
	v340 = m.ExcPending
	if v340 != 0 {
		goto L1
	} else {
		goto L109
	}
L109:
	;
	F_addReplyBulkCString(m, l0, int32(_a_F_addReplySentinelValkeyInstance_7))
	mBase = m.M
	v343 = m.ExcPending
	if v343 != 0 {
		goto L1
	} else {
		goto L110
	}
L110:
	;
	v344 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
	v345 = int64(*(*int32)(unsafe.Add(mBase, uint32(v344))))
	F_addReplyBulkLongLong(m, l0, v345)
	mBase = m.M
	v347 = m.ExcPending
	if v347 != 0 {
		goto L1
	} else {
		goto L111
	}
L111:
	;
	v349 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	if v349&int32(64) == int32(0) {
		v370 = int32(10)
		goto L112
	} else {
		goto L113
	}
L112:
	;
	F_addReplyBulkCString(m, l0, int32(_a_F_addReplySentinelValkeyInstance_8))
	mBase = m.M
	v374 = m.ExcPending
	if v374 != 0 {
		goto L1
	} else {
		goto L118
	}
L113:
	;
	F_addReplyBulkCString(m, l0, int32(_a_F_addReplySentinelValkeyInstance_54))
	mBase = m.M
	v356 = m.ExcPending
	if v356 != 0 {
		goto L1
	} else {
		goto L114
	}
L114:
	;
	v358 = *(*int32)(unsafe.Add(mBase, uint32(l1)+240))
	if base.Ui32(int32(6)) < base.Ui32(v358) {
		v366 = int32(_a_F_addReplySentinelValkeyInstance_55)
		goto L115
	} else {
		goto L116
	}
L115:
	;
	F_addReplyBulkCString(m, l0, v366)
	mBase = m.M
	v368 = m.ExcPending
	if v368 != 0 {
		goto L1
	} else {
		goto L117
	}
L116:
	;
	v365 = *(*int32)(unsafe.Add(mBase, uint32(v358<<(uint(int32(2))%32))+uint32(_c_F_addReplySentinelValkeyInstance[1])))
	v366 = v365
	goto L115
L117:
	;
	v370 = int32(11)
	goto L112
L118:
	;
	v375 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
	v376 = *(*int64)(unsafe.Add(mBase, uint32(v375)+56))
	if base.B2i32(v376 == int64(0)) == int32(0) {
		goto L120
	} else {
		goto L121
	}
L119:
	;
	F_addReplyBulkLongLong(m, l0, v386)
	mBase = m.M
	v388 = m.ExcPending
	if v388 != 0 {
		goto L1
	} else {
		goto L122
	}
L120:
	;
	v382 = F_mstime(m)
	mBase = m.M
	v383 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
	v384 = *(*int64)(unsafe.Add(mBase, uint32(v383)+56))
	v386 = v382 - v384
	goto L119
L121:
	;
	v386 = int64(0)
	goto L119
L122:
	;
	F_addReplyBulkCString(m, l0, int32(_a_F_addReplySentinelValkeyInstance_9))
	mBase = m.M
	v391 = m.ExcPending
	if v391 != 0 {
		goto L1
	} else {
		goto L123
	}
L123:
	;
	v392 = F_mstime(m)
	mBase = m.M
	v393 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
	v394 = *(*int64)(unsafe.Add(mBase, uint32(v393)+48))
	F_addReplyBulkLongLong(m, l0, v392-v394)
	mBase = m.M
	v397 = m.ExcPending
	if v397 != 0 {
		goto L1
	} else {
		goto L124
	}
L124:
	;
	F_addReplyBulkCString(m, l0, int32(_a_F_addReplySentinelValkeyInstance_10))
	mBase = m.M
	v400 = m.ExcPending
	if v400 != 0 {
		goto L1
	} else {
		goto L125
	}
L125:
	;
	v401 = F_mstime(m)
	mBase = m.M
	v402 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
	v403 = *(*int64)(unsafe.Add(mBase, uint32(v402)+72))
	F_addReplyBulkLongLong(m, l0, v401-v403)
	mBase = m.M
	v406 = m.ExcPending
	if v406 != 0 {
		goto L1
	} else {
		goto L126
	}
L126:
	;
	v407 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	if v407&int32(8) == int32(0) {
		v423 = v370
		v424 = v407
		goto L127
	} else {
		goto L128
	}
L127:
	;
	if v424&int32(16) == int32(0) {
		v439 = v423
		goto L131
	} else {
		goto L132
	}
L128:
	;
	F_addReplyBulkCString(m, l0, int32(_a_F_addReplySentinelValkeyInstance_56))
	mBase = m.M
	v414 = m.ExcPending
	if v414 != 0 {
		goto L1
	} else {
		goto L129
	}
L129:
	;
	v415 = F_mstime(m)
	mBase = m.M
	v416 = *(*int64)(unsafe.Add(mBase, uint32(l1)+56))
	F_addReplyBulkLongLong(m, l0, v415-v416)
	mBase = m.M
	v419 = m.ExcPending
	if v419 != 0 {
		goto L1
	} else {
		goto L130
	}
L130:
	;
	v422 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v423 = v370 + int32(1)
	v424 = v422
	goto L127
L131:
	;
	F_addReplyBulkCString(m, l0, int32(_a_F_addReplySentinelValkeyInstance_11))
	mBase = m.M
	v442 = m.ExcPending
	if v442 != 0 {
		goto L1
	} else {
		goto L135
	}
L132:
	;
	F_addReplyBulkCString(m, l0, int32(_a_F_addReplySentinelValkeyInstance_57))
	mBase = m.M
	v431 = m.ExcPending
	if v431 != 0 {
		goto L1
	} else {
		goto L133
	}
L133:
	;
	v432 = F_mstime(m)
	mBase = m.M
	v433 = *(*int64)(unsafe.Add(mBase, uint32(l1)+64))
	F_addReplyBulkLongLong(m, l0, v432-v433)
	mBase = m.M
	v436 = m.ExcPending
	if v436 != 0 {
		goto L1
	} else {
		goto L134
	}
L134:
	;
	v439 = v423 + int32(1)
	goto L131
L135:
	;
	v443 = *(*int64)(unsafe.Add(mBase, uint32(l1)+72))
	F_addReplyBulkLongLong(m, l0, v443)
	mBase = m.M
	v445 = m.ExcPending
	if v445 != 0 {
		goto L1
	} else {
		goto L136
	}
L136:
	;
	v446 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	if v446&int32(3) != 0 {
		goto L138
	} else {
		goto L139
	}
L137:
	;
	if v489&int32(1) == int32(0) {
		v562 = v488
		goto L152
	} else {
		goto L153
	}
L138:
	;
	F_addReplyBulkCString(m, l0, int32(_a_F_addReplySentinelValkeyInstance_12))
	mBase = m.M
	v453 = m.ExcPending
	if v453 != 0 {
		goto L1
	} else {
		goto L140
	}
L139:
	;
	v488 = v439 + int32(1)
	v489 = v446
	goto L137
L140:
	;
	v454 = *(*int64)(unsafe.Add(mBase, uint32(l1)+96))
	if base.B2i32(v454 == int64(0)) == int32(0) {
		goto L142
	} else {
		goto L143
	}
L141:
	;
	F_addReplyBulkLongLong(m, l0, v463)
	mBase = m.M
	v465 = m.ExcPending
	if v465 != 0 {
		goto L1
	} else {
		goto L144
	}
L142:
	;
	v460 = F_mstime(m)
	mBase = m.M
	v461 = *(*int64)(unsafe.Add(mBase, uint32(l1)+96))
	v463 = v460 - v461
	goto L141
L143:
	;
	v463 = int64(0)
	goto L141
L144:
	;
	F_addReplyBulkCString(m, l0, int32(_a_F_addReplySentinelValkeyInstance_13))
	mBase = m.M
	v468 = m.ExcPending
	if v468 != 0 {
		goto L1
	} else {
		goto L145
	}
L145:
	;
	v471 = *(*int32)(unsafe.Add(mBase, uint32(l1)+108))
	if v471 == int32(1) {
		goto L146
	} else {
		goto L147
	}
L146:
	;
	v474 = int32(_a_F_addReplySentinelValkeyInstance_14)
	goto L148
L147:
	;
	v474 = int32(_a_F_addReplySentinelValkeyInstance_15)
	goto L148
L148:
	;
	F_addReplyBulkCString(m, l0, v474)
	mBase = m.M
	v476 = m.ExcPending
	if v476 != 0 {
		goto L1
	} else {
		goto L149
	}
L149:
	;
	F_addReplyBulkCString(m, l0, int32(_a_F_addReplySentinelValkeyInstance_16))
	mBase = m.M
	v479 = m.ExcPending
	if v479 != 0 {
		goto L1
	} else {
		goto L150
	}
L150:
	;
	v480 = F_mstime(m)
	mBase = m.M
	v481 = *(*int64)(unsafe.Add(mBase, uint32(l1)+112))
	F_addReplyBulkLongLong(m, l0, v480-v481)
	mBase = m.M
	v484 = m.ExcPending
	if v484 != 0 {
		goto L1
	} else {
		goto L151
	}
L151:
	;
	v487 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v488 = v439 + int32(4)
	v489 = v487
	goto L137
L152:
	;
	v564 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	if v564&int32(2) == int32(0) {
		v619 = v562
		v620 = v564
		goto L174
	} else {
		goto L175
	}
L153:
	;
	F_addReplyBulkCString(m, l0, int32(_a_F_addReplySentinelValkeyInstance_30))
	mBase = m.M
	v497 = m.ExcPending
	if v497 != 0 {
		goto L1
	} else {
		goto L154
	}
L154:
	;
	v498 = *(*int64)(unsafe.Add(mBase, uint32(l1)+16))
	F_addReplyBulkLongLong(m, l0, v498)
	mBase = m.M
	v500 = m.ExcPending
	if v500 != 0 {
		goto L1
	} else {
		goto L155
	}
L155:
	;
	F_addReplyBulkCString(m, l0, int32(_a_F_addReplySentinelValkeyInstance_31))
	mBase = m.M
	v503 = m.ExcPending
	if v503 != 0 {
		goto L1
	} else {
		goto L156
	}
L156:
	;
	v504 = *(*int32)(unsafe.Add(mBase, uint32(l1)+148))
	v505 = *(*int32)(unsafe.Add(mBase, uint32(v504)+16))
	v506 = *(*int32)(unsafe.Add(mBase, uint32(v504)+12))
	F_addReplyBulkLongLong(m, l0, base.I64_extend_i32_u(v505+v506))
	mBase = m.M
	v510 = m.ExcPending
	if v510 != 0 {
		goto L1
	} else {
		goto L157
	}
L157:
	;
	F_addReplyBulkCString(m, l0, int32(_a_F_addReplySentinelValkeyInstance_32))
	mBase = m.M
	v513 = m.ExcPending
	if v513 != 0 {
		goto L1
	} else {
		goto L158
	}
L158:
	;
	v514 = *(*int32)(unsafe.Add(mBase, uint32(l1)+144))
	v515 = *(*int32)(unsafe.Add(mBase, uint32(v514)+16))
	v516 = *(*int32)(unsafe.Add(mBase, uint32(v514)+12))
	F_addReplyBulkLongLong(m, l0, base.I64_extend_i32_u(v515+v516))
	mBase = m.M
	v520 = m.ExcPending
	if v520 != 0 {
		goto L1
	} else {
		goto L159
	}
L159:
	;
	F_addReplyBulkCString(m, l0, int32(_a_F_addReplySentinelValkeyInstance_33))
	mBase = m.M
	v523 = m.ExcPending
	if v523 != 0 {
		goto L1
	} else {
		goto L160
	}
L160:
	;
	v524 = int64(*(*uint32)(unsafe.Add(mBase, uint32(l1)+152)))
	F_addReplyBulkLongLong(m, l0, v524)
	mBase = m.M
	v526 = m.ExcPending
	if v526 != 0 {
		goto L1
	} else {
		goto L161
	}
L161:
	;
	F_addReplyBulkCString(m, l0, int32(_a_F_addReplySentinelValkeyInstance_34))
	mBase = m.M
	v529 = m.ExcPending
	if v529 != 0 {
		goto L1
	} else {
		goto L162
	}
L162:
	;
	v530 = *(*int64)(unsafe.Add(mBase, uint32(l1)+264))
	F_addReplyBulkLongLong(m, l0, v530)
	mBase = m.M
	v532 = m.ExcPending
	if v532 != 0 {
		goto L1
	} else {
		goto L163
	}
L163:
	;
	F_addReplyBulkCString(m, l0, int32(_a_F_addReplySentinelValkeyInstance_35))
	mBase = m.M
	v535 = m.ExcPending
	if v535 != 0 {
		goto L1
	} else {
		goto L164
	}
L164:
	;
	v536 = int64(*(*int32)(unsafe.Add(mBase, uint32(l1)+156)))
	F_addReplyBulkLongLong(m, l0, v536)
	mBase = m.M
	v538 = m.ExcPending
	if v538 != 0 {
		goto L1
	} else {
		goto L165
	}
L165:
	;
	v539 = *(*int32)(unsafe.Add(mBase, uint32(l1)+284))
	if v539 != 0 {
		goto L167
	} else {
		goto L168
	}
L166:
	;
	v551 = *(*int32)(unsafe.Add(mBase, uint32(l1)+288))
	if v551 == int32(0) {
		v562 = v550
		goto L152
	} else {
		goto L171
	}
L167:
	;
	F_addReplyBulkCString(m, l0, int32(_a_F_addReplySentinelValkeyInstance_36))
	mBase = m.M
	v544 = m.ExcPending
	if v544 != 0 {
		goto L1
	} else {
		goto L169
	}
L168:
	;
	v550 = v488 + int32(6)
	goto L166
L169:
	;
	v545 = *(*int32)(unsafe.Add(mBase, uint32(l1)+284))
	F_addReplyBulkCString(m, l0, v545)
	mBase = m.M
	v547 = m.ExcPending
	if v547 != 0 {
		goto L1
	} else {
		goto L170
	}
L170:
	;
	v550 = v488 + int32(7)
	goto L166
L171:
	;
	F_addReplyBulkCString(m, l0, int32(_a_F_addReplySentinelValkeyInstance_37))
	mBase = m.M
	v556 = m.ExcPending
	if v556 != 0 {
		goto L1
	} else {
		goto L172
	}
L172:
	;
	v557 = *(*int32)(unsafe.Add(mBase, uint32(l1)+288))
	F_addReplyBulkCString(m, l0, v557)
	mBase = m.M
	v559 = m.ExcPending
	if v559 != 0 {
		goto L1
	} else {
		goto L173
	}
L173:
	;
	v562 = v550 + int32(1)
	goto L152
L174:
	;
	if v620&int32(4) == int32(0) {
		v649 = v619
		goto L196
	} else {
		goto L197
	}
L175:
	;
	F_addReplyBulkCString(m, l0, int32(_a_F_addReplySentinelValkeyInstance_21))
	mBase = m.M
	v571 = m.ExcPending
	if v571 != 0 {
		goto L1
	} else {
		goto L176
	}
L176:
	;
	v572 = *(*int64)(unsafe.Add(mBase, uint32(l1)+168))
	F_addReplyBulkLongLong(m, l0, v572)
	mBase = m.M
	v574 = m.ExcPending
	if v574 != 0 {
		goto L1
	} else {
		goto L177
	}
L177:
	;
	F_addReplyBulkCString(m, l0, int32(_a_F_addReplySentinelValkeyInstance_22))
	mBase = m.M
	v577 = m.ExcPending
	if v577 != 0 {
		goto L1
	} else {
		goto L178
	}
L178:
	;
	v580 = *(*int32)(unsafe.Add(mBase, uint32(l1)+204))
	if v580 != 0 {
		goto L179
	} else {
		goto L180
	}
L179:
	;
	v581 = int32(_a_F_addReplySentinelValkeyInstance_23)
	goto L181
L180:
	;
	v581 = int32(_a_F_addReplySentinelValkeyInstance_24)
	goto L181
L181:
	;
	F_addReplyBulkCString(m, l0, v581)
	mBase = m.M
	v583 = m.ExcPending
	if v583 != 0 {
		goto L1
	} else {
		goto L182
	}
L182:
	;
	F_addReplyBulkCString(m, l0, int32(_a_F_addReplySentinelValkeyInstance_25))
	mBase = m.M
	v586 = m.ExcPending
	if v586 != 0 {
		goto L1
	} else {
		goto L183
	}
L183:
	;
	v587 = *(*int32)(unsafe.Add(mBase, uint32(l1)+196))
	if v587 != 0 {
		goto L184
	} else {
		goto L185
	}
L184:
	;
	v589 = v587
	goto L186
L185:
	;
	v589 = int32(_a_F_addReplySentinelValkeyInstance_19)
	goto L186
L186:
	;
	F_addReplyBulkCString(m, l0, v589)
	mBase = m.M
	v591 = m.ExcPending
	if v591 != 0 {
		goto L1
	} else {
		goto L187
	}
L187:
	;
	F_addReplyBulkCString(m, l0, int32(_a_F_addReplySentinelValkeyInstance_26))
	mBase = m.M
	v594 = m.ExcPending
	if v594 != 0 {
		goto L1
	} else {
		goto L188
	}
L188:
	;
	v595 = int64(*(*int32)(unsafe.Add(mBase, uint32(l1)+200)))
	F_addReplyBulkLongLong(m, l0, v595)
	mBase = m.M
	v597 = m.ExcPending
	if v597 != 0 {
		goto L1
	} else {
		goto L189
	}
L189:
	;
	F_addReplyBulkCString(m, l0, int32(_a_F_addReplySentinelValkeyInstance_27))
	mBase = m.M
	v600 = m.ExcPending
	if v600 != 0 {
		goto L1
	} else {
		goto L190
	}
L190:
	;
	v601 = int64(*(*int32)(unsafe.Add(mBase, uint32(l1)+176)))
	F_addReplyBulkLongLong(m, l0, v601)
	mBase = m.M
	v603 = m.ExcPending
	if v603 != 0 {
		goto L1
	} else {
		goto L191
	}
L191:
	;
	F_addReplyBulkCString(m, l0, int32(_a_F_addReplySentinelValkeyInstance_28))
	mBase = m.M
	v606 = m.ExcPending
	if v606 != 0 {
		goto L1
	} else {
		goto L192
	}
L192:
	;
	v607 = *(*int64)(unsafe.Add(mBase, uint32(l1)+208))
	F_addReplyBulkLongLong(m, l0, v607)
	mBase = m.M
	v609 = m.ExcPending
	if v609 != 0 {
		goto L1
	} else {
		goto L193
	}
L193:
	;
	F_addReplyBulkCString(m, l0, int32(_a_F_addReplySentinelValkeyInstance_29))
	mBase = m.M
	v612 = m.ExcPending
	if v612 != 0 {
		goto L1
	} else {
		goto L194
	}
L194:
	;
	v613 = int64(*(*int32)(unsafe.Add(mBase, uint32(l1)+180)))
	F_addReplyBulkLongLong(m, l0, v613)
	mBase = m.M
	v615 = m.ExcPending
	if v615 != 0 {
		goto L1
	} else {
		goto L195
	}
L195:
	;
	v618 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v619 = v562 + int32(7)
	v620 = v618
	goto L174
L196:
	;
	F_setDeferredMapLen(m, l0, v9, v649)
	mBase = m.M
	v652 = m.ExcPending
	if v652 != 0 {
		goto L1
	} else {
		goto L207
	}
L197:
	;
	F_addReplyBulkCString(m, l0, int32(_a_F_addReplySentinelValkeyInstance_17))
	mBase = m.M
	v627 = m.ExcPending
	if v627 != 0 {
		goto L1
	} else {
		goto L198
	}
L198:
	;
	v628 = F_mstime(m)
	mBase = m.M
	v629 = *(*int64)(unsafe.Add(mBase, uint32(l1)+40))
	F_addReplyBulkLongLong(m, l0, v628-v629)
	mBase = m.M
	v632 = m.ExcPending
	if v632 != 0 {
		goto L1
	} else {
		goto L199
	}
L199:
	;
	F_addReplyBulkCString(m, l0, int32(_a_F_addReplySentinelValkeyInstance_18))
	mBase = m.M
	v635 = m.ExcPending
	if v635 != 0 {
		goto L1
	} else {
		goto L200
	}
L200:
	;
	v636 = *(*int32)(unsafe.Add(mBase, uint32(l1)+216))
	if v636 != 0 {
		goto L201
	} else {
		goto L202
	}
L201:
	;
	v638 = v636
	goto L203
L202:
	;
	v638 = int32(_a_F_addReplySentinelValkeyInstance_19)
	goto L203
L203:
	;
	F_addReplyBulkCString(m, l0, v638)
	mBase = m.M
	v640 = m.ExcPending
	if v640 != 0 {
		goto L1
	} else {
		goto L204
	}
L204:
	;
	F_addReplyBulkCString(m, l0, int32(_a_F_addReplySentinelValkeyInstance_20))
	mBase = m.M
	v643 = m.ExcPending
	if v643 != 0 {
		goto L1
	} else {
		goto L205
	}
L205:
	;
	v644 = *(*int64)(unsafe.Add(mBase, uint32(l1)+224))
	F_addReplyBulkLongLong(m, l0, v644)
	mBase = m.M
	v646 = m.ExcPending
	if v646 != 0 {
		goto L1
	} else {
		goto L206
	}
L206:
	;
	v649 = v619 + int32(3)
	goto L196
L207:
	;
	return
}
func F_addReplyStatus(m *base.Module, l0 int32, l1 int32) {
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
	if l1&int32(3) == int32(0) {
		v24 = l1
		goto L3
	} else {
		goto L4
	}
L1:
	;
	F_addReplyStatusLength(m, l0, l1, v57)
	mBase = m.M
	v59 = m.ExcPending
	if v59 != 0 {
		goto L17
	} else {
		goto L18
	}
L2:
	;
	v57 = v49 - l1
	goto L1
L3:
	;
	v28 = v24
	goto L11
L4:
	;
	v10 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	if v10 != 0 {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	v13 = l1
	goto L7
L6:
	;
	v57 = l1 - l1
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
	return
L18:
	;
	return
}
func F_callReplyArray(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
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
	var v37 int32
	_ = v37
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v48 int32
	_ = v48
	var v53 int32
	_ = v53
	var v63 int32
	_ = v63
	*(*int32)(unsafe.Add(mBase, uint32(l1)+24)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(l1)+16)) = int32(3)
	v13 = F_valkey_calloc(m, l2*int32(48))
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+32)) = v13
	if l2 == int32(0) {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+8)) = l3
	v63 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(l1)+12)) = v63 - l3
	return
L4:
	;
	v23 = v13
	v24 = int32(0)
	goto L5
L5:
	;
	v27 = v24 * int32(48)
	v29 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	*(*int32)(unsafe.Add(mBase, uint32(v23+v27))) = v29
	v31 = *(*int32)(unsafe.Add(mBase, uint32(l1)+32))
	v33 = F_parseReply(m, l0, v31+v27)
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
		goto L1
	} else {
		goto L7
	}
L6:
	;
	goto L3
L7:
	;
	v35 = *(*int32)(unsafe.Add(mBase, uint32(l1)+32))
	v36 = v35 + v27
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v36)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v36)+20)) = v37 | int32(2)
	v41 = *(*int32)(unsafe.Add(mBase, uint32(l1)+32))
	v43 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v41+v27)+20)))
	if v43&int32(4) == int32(0) {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v53 = v24 + int32(1)
	if v53 != l2 {
		v23 = v41
		v24 = v53
		goto L5
	} else {
		goto L10
	}
L9:
	;
	v48 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	*(*int32)(unsafe.Add(mBase, uint32(l1)+20)) = v48 | int32(4)
	goto L8
L10:
	;
	goto L6
}
func F_callReplyAttribute(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
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
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v58 int32
	_ = v58
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v88 int32
	_ = v88
	var v93 int32
	_ = v93
	var v105 int32
	_ = v105
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v113 int32
	_ = v113
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	v11 = F_valkey_calloc(m, int32(48))
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
	*(*int32)(unsafe.Add(mBase, uint32(l1)+44)) = v11
	*(*int32)(unsafe.Add(mBase, uint32(v11)+24)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = int32(11)
	v19 = F_valkey_calloc(m, l2*int32(96))
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+32)) = v19
	v23 = l2 << (uint(int32(1)) % 32)
	if v23 == int32(0) {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+8)) = l3
	v105 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+12)) = v105 - l3
	v108 = *(*int32)(unsafe.Add(mBase, uint32(l1)+44))
	v109 = *(*int32)(unsafe.Add(mBase, uint32(v108)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v108)+20)) = v109 | int32(6)
	v113 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	*(*int32)(unsafe.Add(mBase, uint32(v108))) = v113
	v115 = F_parseReply(m, l0, l1)
	mBase = m.M
	v116 = m.ExcPending
	if v116 != 0 {
		goto L1
	} else {
		goto L15
	}
L5:
	;
	v32 = v19
	v34 = int32(0)
	goto L6
L6:
	;
	v37 = v34 * int32(48)
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
	*(*int32)(unsafe.Add(mBase, uint32(v32+v37))) = v39
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v11)+32))
	v43 = F_parseReply(m, l0, v41+v37)
	mBase = m.M
	v44 = m.ExcPending
	if v44 != 0 {
		goto L1
	} else {
		goto L8
	}
L7:
	;
	goto L4
L8:
	;
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v11)+32))
	v46 = v45 + v37
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v46)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v46)+20)) = v47 | int32(2)
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v11)+32))
	v53 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v51+v37)+20)))
	if v53&int32(4) == int32(0) {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	v64 = int32(48)
	v65 = (v34 | int32(1)) * v64
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
	*(*int32)(unsafe.Add(mBase, uint32(v51+v65))) = v67
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v11)+32))
	v73 = F_parseReply(m, l0, v69+v37+v64)
	mBase = m.M
	v74 = m.ExcPending
	if v74 != 0 {
		goto L1
	} else {
		goto L11
	}
L10:
	;
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v11)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+20)) = v58 | int32(4)
	goto L9
L11:
	;
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v11)+32))
	v76 = v75 + v65
	v77 = *(*int32)(unsafe.Add(mBase, uint32(v76)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v76)+20)) = v77 | int32(2)
	v81 = *(*int32)(unsafe.Add(mBase, uint32(v11)+32))
	v83 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v81+v65)+20)))
	if v83&int32(4) == int32(0) {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	v93 = v34 + int32(2)
	if base.Ui32(v93) < base.Ui32(v23) {
		v32 = v81
		v34 = v93
		goto L6
	} else {
		goto L14
	}
L13:
	;
	v88 = *(*int32)(unsafe.Add(mBase, uint32(v11)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+20)) = v88 | int32(4)
	goto L12
L14:
	;
	goto L7
L15:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+8)) = l3
	v118 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v119 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	*(*int32)(unsafe.Add(mBase, uint32(l1)+20)) = v119 | int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(l1)+12)) = v118 - l3
	return
}
func F_callReplyBool(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = l3
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = int32(7)
	*(*int64)(unsafe.Add(mBase, uint32(l0)+32)) = base.I64_extend_i32_s(l1)
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v11 | int32(4)
	return
}
func F_callReplyBulkString(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) {
	mBase := m.M
	_ = mBase
	*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = l4
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = l3
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = int32(0)
	return
}
func F_callReplyCreate(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
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
	v8 = F_valkey_malloc(m, int32(48))
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return int32(0)
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(v8)+8)) = l0
		*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = l0
		*(*int32)(unsafe.Add(mBase, uint32(v8)+20)) = int32(1)
		v19 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0+int32(-1)))))
		switch v19 & int32(7) {
		case 0:
			v36 = int32(base.Ui32(v19) >> (uint(int32(3)) % 32))
		case 1:
			v26 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0+int32(-3)))))
			v36 = v26
		case 2:
			v29 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0+int32(-5)))))
			v36 = v29
		case 3:
			v32 = *(*int32)(unsafe.Add(mBase, uint32(l0+int32(-9))))
			v36 = v32
		case 4:
			v35 = *(*int32)(unsafe.Add(mBase, uint32(l0+int32(-17))))
			v36 = v35
		default:
			v36 = int32(0)
		}
		*(*int32)(unsafe.Add(mBase, uint32(v8)+44)) = int32(0)
		*(*int32)(unsafe.Add(mBase, uint32(v8))) = l2
		*(*int32)(unsafe.Add(mBase, uint32(v8)+12)) = v36
		*(*int32)(unsafe.Add(mBase, uint32(v8)+40)) = l1
		return v8
	}
}
func F_callReplyDouble(m *base.Module, l0 int32, l1 float64, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = l3
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = int32(8)
	*(*float64)(unsafe.Add(mBase, uint32(l0)+32)) = l1
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v10 | int32(4)
	return
}
func F_callReplyGetLongLong(m *base.Module, l0 int32) int64 {
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
	var v67 int64
	_ = v67
	var v68 int64
	_ = v68
	v5 = m.G0
	v7 = v5 - int32(80)
	m.G0 = v7
	v9 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+20)))
	if v9&int32(2) != 0 {
		v64 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
		if v64 != int32(2) {
			v68 = int64(-9223372036854775807 - 1)
		} else {
			v67 = *(*int64)(unsafe.Add(mBase, uint32(l0)+32))
			v68 = v67
		}
		m.G0 = v7 + int32(80)
		return v68
	} else {
		v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		v15 = int32(0)
		v16 = *(*int64)(unsafe.Add(mBase, _c_F_callReplyGetLongLong[0]))
		*(*int64)(unsafe.Add(mBase, uint32(v7+int32(24)))) = v16
		v21 = *(*int64)(unsafe.Add(mBase, _c_F_callReplyGetLongLong[1]))
		*(*int64)(unsafe.Add(mBase, uint32(v7+int32(32)))) = v21
		v26 = *(*int64)(unsafe.Add(mBase, _c_F_callReplyGetLongLong[2]))
		*(*int64)(unsafe.Add(mBase, uint32(v7+int32(40)))) = v26
		v31 = *(*int64)(unsafe.Add(mBase, _c_F_callReplyGetLongLong[3]))
		*(*int64)(unsafe.Add(mBase, uint32(v7+int32(48)))) = v31
		v36 = *(*int64)(unsafe.Add(mBase, _c_F_callReplyGetLongLong[4]))
		*(*int64)(unsafe.Add(mBase, uint32(v7+int32(56)))) = v36
		v41 = *(*int64)(unsafe.Add(mBase, _c_F_callReplyGetLongLong[5]))
		*(*int64)(unsafe.Add(mBase, uint32(v7+int32(64)))) = v41
		v46 = *(*int64)(unsafe.Add(mBase, _c_F_callReplyGetLongLong[6]))
		*(*int64)(unsafe.Add(mBase, uint32(v7+int32(72)))) = v46
		*(*int32)(unsafe.Add(mBase, uint32(v7)+12)) = v12
		v50 = *(*int64)(unsafe.Add(mBase, _c_F_callReplyGetLongLong[7]))
		*(*int64)(unsafe.Add(mBase, uint32(v7)+16)) = v50
		v54 = F_parseReply(m, v7+int32(12), l0)
		mBase = m.M
		v57 = m.ExcPending
		if v57 != 0 {
			return int64(0)
		} else {
			v58 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
			*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v58 | int32(2)
			v64 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
			if v64 != int32(2) {
				v68 = int64(-9223372036854775807 - 1)
			} else {
				v67 = *(*int64)(unsafe.Add(mBase, uint32(l0)+32))
				v68 = v67
			}
			m.G0 = v7 + int32(80)
			return v68
		}
	}
}
func F_callReplyGetProto(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v3
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	return v5
}
func F_callReplyGetVerbatim(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
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
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	v6 = m.G0
	v8 = v6 - int32(80)
	m.G0 = v8
	v10 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+20)))
	if v10&int32(2) != 0 {
		v65 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
		if v65 != int32(10) {
			v75 = int32(0)
		} else {
			v68 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
			*(*int32)(unsafe.Add(mBase, uint32(l1))) = v68
			if l2 == int32(0) {
			} else {
				v72 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
				*(*int32)(unsafe.Add(mBase, uint32(l2))) = v72
			}
			v74 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
			v75 = v74
		}
		m.G0 = v8 + int32(80)
		return v75
	} else {
		v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		v16 = int32(0)
		v17 = *(*int64)(unsafe.Add(mBase, _c_F_callReplyGetVerbatim[0]))
		*(*int64)(unsafe.Add(mBase, uint32(v8+int32(24)))) = v17
		v22 = *(*int64)(unsafe.Add(mBase, _c_F_callReplyGetVerbatim[1]))
		*(*int64)(unsafe.Add(mBase, uint32(v8+int32(32)))) = v22
		v27 = *(*int64)(unsafe.Add(mBase, _c_F_callReplyGetVerbatim[2]))
		*(*int64)(unsafe.Add(mBase, uint32(v8+int32(40)))) = v27
		v32 = *(*int64)(unsafe.Add(mBase, _c_F_callReplyGetVerbatim[3]))
		*(*int64)(unsafe.Add(mBase, uint32(v8+int32(48)))) = v32
		v37 = *(*int64)(unsafe.Add(mBase, _c_F_callReplyGetVerbatim[4]))
		*(*int64)(unsafe.Add(mBase, uint32(v8+int32(56)))) = v37
		v42 = *(*int64)(unsafe.Add(mBase, _c_F_callReplyGetVerbatim[5]))
		*(*int64)(unsafe.Add(mBase, uint32(v8+int32(64)))) = v42
		v47 = *(*int64)(unsafe.Add(mBase, _c_F_callReplyGetVerbatim[6]))
		*(*int64)(unsafe.Add(mBase, uint32(v8+int32(72)))) = v47
		*(*int32)(unsafe.Add(mBase, uint32(v8)+12)) = v13
		v51 = *(*int64)(unsafe.Add(mBase, _c_F_callReplyGetVerbatim[7]))
		*(*int64)(unsafe.Add(mBase, uint32(v8)+16)) = v51
		v55 = F_parseReply(m, v8+int32(12), l0)
		mBase = m.M
		v58 = m.ExcPending
		if v58 != 0 {
			return int32(0)
		} else {
			v59 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
			*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v59 | int32(2)
			v65 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
			if v65 != int32(10) {
				v75 = int32(0)
			} else {
				v68 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
				*(*int32)(unsafe.Add(mBase, uint32(l1))) = v68
				if l2 == int32(0) {
				} else {
					v72 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
					*(*int32)(unsafe.Add(mBase, uint32(l2))) = v72
				}
				v74 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
				v75 = v74
			}
			m.G0 = v8 + int32(80)
			return v75
		}
	}
}
func F_callReplyLong(m *base.Module, l0 int32, l1 int64, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	*(*int64)(unsafe.Add(mBase, uint32(l0)+32)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = l3
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = int32(2)
	return
}
func F_callReplyNull(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = l1
	v6 = int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v6
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v8 | v6
	return
}
func F_freeReplyObject(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v9 int32
	_ = v9
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v29 int32
	_ = v29
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
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	if l0 == int32(0) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if base.Ui32(int32(14)) < base.Ui32(v5) {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v46 = m.G4
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v46)+16))
	m.T0[v47].(func(*base.Module, int32))(m, l0)
	mBase = m.M
	v49 = m.ExcPending
	if v49 != 0 {
		goto L13
	} else {
		goto L17
	}
L4:
	;
	v9 = int32(1) << (uint(v5) % 32)
	if v9&int32(24802) != 0 {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	v40 = m.G4
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v40)+16))
	m.T0[v41].(func(*base.Module, int32))(m, v39)
	mBase = m.M
	v43 = m.ExcPending
	if v43 != 0 {
		goto L13
	} else {
		goto L16
	}
L6:
	;
	v37 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v39 = v37
	goto L5
L7:
	;
	if v9&int32(7684) == int32(0) {
		goto L3
	} else {
		goto L8
	}
L8:
	;
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	if v16 == int32(0) {
		goto L3
	} else {
		goto L9
	}
L9:
	;
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	if v19 == int32(0) {
		v39 = v16
		goto L5
	} else {
		goto L10
	}
L10:
	;
	v24 = int32(0)
	goto L11
L11:
	;
	v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v25+v24<<(uint(int32(2))%32))))
	F_freeReplyObject(m, v29)
	mBase = m.M
	v31 = m.ExcPending
	if v31 != 0 {
		goto L13
	} else {
		goto L14
	}
L12:
	;
	v36 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v39 = v36
	goto L5
L13:
	;
	return
L14:
	;
	v33 = v24 + int32(1)
	v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	if base.Ui32(v33) < base.Ui32(v34) {
		v24 = v33
		goto L11
	} else {
		goto L15
	}
L15:
	;
	goto L12
L16:
	;
	goto L3
L17:
	;
	goto L1
}
func F_replyEngineStats(m *base.Module, l0 int32, l1 int32) {
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
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v20 int64
	_ = v20
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v26 int64
	_ = v26
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	F_addReplyBulkCString(m, l1, v3)
	mBase = m.M
	v5 = m.ExcPending
	if v5 != 0 {
		return
	} else {
		F_addReplyMapLen(m, l1, int32(2))
		mBase = m.M
		v8 = m.ExcPending
		if v8 != 0 {
			return
		} else {
			v10 = *(*int32)(unsafe.Add(mBase, _c_F_replyEngineStats[0]))
			v11 = *(*int32)(unsafe.Add(mBase, uint32(v10)+12))
			v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			v13 = F_dictFetchValue(m, v11, v12)
			mBase = m.M
			v14 = m.ExcPending
			if v14 != 0 {
				return
			} else {
				F_addReplyBulkCString(m, l1, int32(_a_F_replyEngineStats_0))
				mBase = m.M
				v17 = m.ExcPending
				if v17 != 0 {
					return
				} else {
					if v13 == int32(0) {
						F_addReplyLongLong(m, l1, int64(0))
						mBase = m.M
						v31 = m.ExcPending
						if v31 != 0 {
							return
						} else {
							F_addReplyBulkCString(m, l1, int32(_a_F_replyEngineStats_1))
							mBase = m.M
							v34 = m.ExcPending
							if v34 != 0 {
								return
							} else {
								F_addReplyLongLong(m, l1, int64(0))
								mBase = m.M
								v37 = m.ExcPending
								if v37 != 0 {
									return
								} else {
									return
								}
							}
						}
					} else {
						v20 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v13))))
						F_addReplyLongLong(m, l1, v20)
						mBase = m.M
						v22 = m.ExcPending
						if v22 != 0 {
							return
						} else {
							F_addReplyBulkCString(m, l1, int32(_a_F_replyEngineStats_1))
							mBase = m.M
							v25 = m.ExcPending
							if v25 != 0 {
								return
							} else {
								v26 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v13)+4)))
								F_addReplyLongLong(m, l1, v26)
								mBase = m.M
								v28 = m.ExcPending
								if v28 != 0 {
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
func F_replyHandlersArray(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
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
	v7 = *(*int32)(unsafe.Add(mBase, uint32(v6)+60))
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
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v31)+64))
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
func F_replyHandlersAttribute(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
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
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(v6)+52))
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
	v33 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v33)+56))
	if v34 == int32(0) {
		goto L12
	} else {
		goto L13
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
	v23 = F_parseReply(m, l0, l1)
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		goto L3
	} else {
		goto L10
	}
L10:
	;
	v26 = v20 + int32(1)
	if v26 != l2 {
		v20 = v26
		goto L7
	} else {
		goto L11
	}
L11:
	;
	goto L8
L12:
	;
	v40 = F_parseReply(m, l0, l1)
	mBase = m.M
	v41 = m.ExcPending
	if v41 != 0 {
		goto L3
	} else {
		goto L15
	}
L13:
	;
	v37 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	m.T0[v34].(func(*base.Module, int32))(m, v37)
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L3
	} else {
		goto L14
	}
L14:
	;
	goto L12
L15:
	;
	return
}
func F_replyHandlersBigNumber(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) {
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
	v8 = *(*int32)(unsafe.Add(mBase, uint32(v7)+44))
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
func F_replyHandlersBulkString(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) {
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
	v8 = *(*int32)(unsafe.Add(mBase, uint32(v7)+20))
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
func F_replyHandlersDouble(m *base.Module, l0 int32, l1 float64, l2 int32, l3 int32) {
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
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(v6)+40))
	if v7 == int32(0) {
		return
	} else {
		v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		m.T0[v7].(func(*base.Module, int32, float64))(m, v10, l1)
		mBase = m.M
		v12 = m.ExcPending
		if v12 != 0 {
			return
		} else {
			return
		}
	}
}
func F_replyHandlersLong(m *base.Module, l0 int32, l1 int64, l2 int32, l3 int32) {
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
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(v6)+36))
	if v7 == int32(0) {
		return
	} else {
		v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		m.T0[v7].(func(*base.Module, int32, int64))(m, v10, l1)
		mBase = m.M
		v12 = m.ExcPending
		if v12 != 0 {
			return
		} else {
			return
		}
	}
}
func F_replyHandlersNullBulkString(m *base.Module, l0 int32, l1 int32, l2 int32) {
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
	v6 = *(*int32)(unsafe.Add(mBase, uint32(v5)+12))
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
func F_replyHandlersParseError(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v4 = *(*int32)(unsafe.Add(mBase, uint32(v3)+84))
	if v4 == int32(0) {
		return
	} else {
		v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		m.T0[v4].(func(*base.Module, int32))(m, v7)
		mBase = m.M
		v9 = m.ExcPending
		if v9 != 0 {
			return
		} else {
			return
		}
	}
}
func F_replyHandlersVerbatimString(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v9 = *(*int32)(unsafe.Add(mBase, uint32(v8)+28))
	if v9 == int32(0) {
		return
	} else {
		v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		m.T0[v9].(func(*base.Module, int32, int32, int32, int32))(m, v12, l2, l3, l1)
		mBase = m.M
		v14 = m.ExcPending
		if v14 != 0 {
			return
		} else {
			return
		}
	}
}
func F_sendReplyToClient(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v3 = F_trySendWriteToIOThreads(m, v2)
	mBase = m.M
	v4 = m.ExcPending
	if v4 != 0 {
		return
	} else {
		if v3 == int32(0) {
			return
		} else {
			v7 = F_writeToClient(m, v2)
			mBase = m.M
			v8 = m.ExcPending
			if v8 != 0 {
				return
			} else {
				return
			}
		}
	}
}
