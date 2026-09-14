package p1

import (
	base "github.com/shibukawa/vkmem/internal/aot/vkaot/base"
	"unsafe"
)

func F_freeRdbProfile(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
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
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v32 int32
	_ = v32
	if l1 == int32(0) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	F_valkey_free(m, l0)
	mBase = m.M
	v32 = m.ExcPending
	if v32 != 0 {
		goto L5
	} else {
		goto L10
	}
L2:
	;
	v10 = int32(0)
	goto L3
L3:
	;
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0+v10<<(uint(int32(2))%32))))
	v16 = *(*int32)(unsafe.Add(mBase, uint32(v15)+40))
	F_hdr_close(m, v16)
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		goto L5
	} else {
		goto L6
	}
L4:
	;
	goto L1
L5:
	;
	return
L6:
	;
	v19 = *(*int32)(unsafe.Add(mBase, uint32(v15)+44))
	F_hdr_close(m, v19)
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		goto L5
	} else {
		goto L7
	}
L7:
	;
	F_valkey_free(m, v15)
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		goto L5
	} else {
		goto L8
	}
L8:
	;
	v25 = v10 + int32(1)
	if v25 != l1 {
		v10 = v25
		goto L3
	} else {
		goto L9
	}
L9:
	;
	goto L4
L10:
	;
	return
}
func F_rdbCheckHandleCrash(m *base.Module, l0 int32, l1 int32, l2 int32) {
	var v7 int32
	_ = v7
	F_rdbCheckError(m, int32(_a_F_rdbCheckHandleCrash_0), int32(0))
	v7 = m.ExcPending
	if v7 != 0 {
		return
	} else {
		m.Env.Exit(m, int32(1))
		base.Wasm_trap_unreachable()
		for {
		}
	}
}
func F_rdbFunctionLoad(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v25 int32
	_ = v25
	var v31 int32
	_ = v31
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
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v59 int32
	_ = v59
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	v6 = int32(0)
	v8 = m.G0
	v10 = v8 - int32(16)
	m.G0 = v10
	*(*int32)(unsafe.Add(mBase, uint32(v10)+12)) = v6
	v17 = F_rdbGenericLoadStringObject(m, l0, int32(4), v6)
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		return int32(0)
	} else {
		if v17 == int32(0) {
			v40 = F_sdsnew(m, int32(_a_F_rdbFunctionLoad_0))
			mBase = m.M
			v41 = m.ExcPending
			if v41 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v10)+12)) = v40
				v51 = v40
				v52 = int32(-1)
				if v51 == int32(0) {
					m.G0 = v10 + int32(16)
					return v52
				} else {
					if l4 == int32(0) {
						v59 = *(*int32)(unsafe.Add(mBase, _c_F_rdbFunctionLoad[0]))
						if int32(3) < v59 {
							v68 = v51
							F_sdsfree(m, v68)
							mBase = m.M
							v70 = m.ExcPending
							if v70 != 0 {
								return int32(0)
							} else {
								m.G0 = v10 + int32(16)
								return v52
							}
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v10))) = v51
							F__serverLog(m, int32(3), int32(_a_F_rdbFunctionLoad_1), v10)
							mBase = m.M
							v66 = m.ExcPending
							if v66 != 0 {
								return int32(0)
							} else {
								v67 = *(*int32)(unsafe.Add(mBase, uint32(v10)+12))
								v68 = v67
								F_sdsfree(m, v68)
								mBase = m.M
								v70 = m.ExcPending
								if v70 != 0 {
									return int32(0)
								} else {
									m.G0 = v10 + int32(16)
									return v52
								}
							}
						}
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(l4))) = v51
						m.G0 = v10 + int32(16)
						return v52
					}
				}
			}
		} else {
			if l2 == int32(0) {
				v47 = v6
				F_sdsfree(m, v17)
				mBase = m.M
				v49 = m.ExcPending
				if v49 != 0 {
					return int32(0)
				} else {
					v50 = *(*int32)(unsafe.Add(mBase, uint32(v10)+12))
					v51 = v50
					v52 = v47
					if v51 == int32(0) {
						m.G0 = v10 + int32(16)
						return v52
					} else {
						if l4 == int32(0) {
							v59 = *(*int32)(unsafe.Add(mBase, _c_F_rdbFunctionLoad[0]))
							if int32(3) < v59 {
								v68 = v51
								F_sdsfree(m, v68)
								mBase = m.M
								v70 = m.ExcPending
								if v70 != 0 {
									return int32(0)
								} else {
									m.G0 = v10 + int32(16)
									return v52
								}
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v10))) = v51
								F__serverLog(m, int32(3), int32(_a_F_rdbFunctionLoad_1), v10)
								mBase = m.M
								v66 = m.ExcPending
								if v66 != 0 {
									return int32(0)
								} else {
									v67 = *(*int32)(unsafe.Add(mBase, uint32(v10)+12))
									v68 = v67
									F_sdsfree(m, v68)
									mBase = m.M
									v70 = m.ExcPending
									if v70 != 0 {
										return int32(0)
									} else {
										m.G0 = v10 + int32(16)
										return v52
									}
								}
							}
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(l4))) = v51
							m.G0 = v10 + int32(16)
							return v52
						}
					}
				}
			} else {
				v25 = int32(0)
				v31 = F_functionsCreateWithLibraryCtx(m, v17, l3&int32(4), v10+int32(12), l2, v25)
				mBase = m.M
				v32 = m.ExcPending
				if v32 != 0 {
					return int32(0)
				} else {
					if v31 != 0 {
						F_sdsfree(m, v31)
						mBase = m.M
						v45 = m.ExcPending
						if v45 != 0 {
							return int32(0)
						} else {
							v47 = v25
							F_sdsfree(m, v17)
							mBase = m.M
							v49 = m.ExcPending
							if v49 != 0 {
								return int32(0)
							} else {
								v50 = *(*int32)(unsafe.Add(mBase, uint32(v10)+12))
								v51 = v50
								v52 = v47
								if v51 == int32(0) {
									m.G0 = v10 + int32(16)
									return v52
								} else {
									if l4 == int32(0) {
										v59 = *(*int32)(unsafe.Add(mBase, _c_F_rdbFunctionLoad[0]))
										if int32(3) < v59 {
											v68 = v51
											F_sdsfree(m, v68)
											mBase = m.M
											v70 = m.ExcPending
											if v70 != 0 {
												return int32(0)
											} else {
												m.G0 = v10 + int32(16)
												return v52
											}
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(v10))) = v51
											F__serverLog(m, int32(3), int32(_a_F_rdbFunctionLoad_1), v10)
											mBase = m.M
											v66 = m.ExcPending
											if v66 != 0 {
												return int32(0)
											} else {
												v67 = *(*int32)(unsafe.Add(mBase, uint32(v10)+12))
												v68 = v67
												F_sdsfree(m, v68)
												mBase = m.M
												v70 = m.ExcPending
												if v70 != 0 {
													return int32(0)
												} else {
													m.G0 = v10 + int32(16)
													return v52
												}
											}
										}
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(l4))) = v51
										m.G0 = v10 + int32(16)
										return v52
									}
								}
							}
						}
					} else {
						v33 = int32(-1)
						v34 = *(*int32)(unsafe.Add(mBase, uint32(v10)+12))
						if v34 != 0 {
							v47 = v33
							F_sdsfree(m, v17)
							mBase = m.M
							v49 = m.ExcPending
							if v49 != 0 {
								return int32(0)
							} else {
								v50 = *(*int32)(unsafe.Add(mBase, uint32(v10)+12))
								v51 = v50
								v52 = v47
								if v51 == int32(0) {
									m.G0 = v10 + int32(16)
									return v52
								} else {
									if l4 == int32(0) {
										v59 = *(*int32)(unsafe.Add(mBase, _c_F_rdbFunctionLoad[0]))
										if int32(3) < v59 {
											v68 = v51
											F_sdsfree(m, v68)
											mBase = m.M
											v70 = m.ExcPending
											if v70 != 0 {
												return int32(0)
											} else {
												m.G0 = v10 + int32(16)
												return v52
											}
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(v10))) = v51
											F__serverLog(m, int32(3), int32(_a_F_rdbFunctionLoad_1), v10)
											mBase = m.M
											v66 = m.ExcPending
											if v66 != 0 {
												return int32(0)
											} else {
												v67 = *(*int32)(unsafe.Add(mBase, uint32(v10)+12))
												v68 = v67
												F_sdsfree(m, v68)
												mBase = m.M
												v70 = m.ExcPending
												if v70 != 0 {
													return int32(0)
												} else {
													m.G0 = v10 + int32(16)
													return v52
												}
											}
										}
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(l4))) = v51
										m.G0 = v10 + int32(16)
										return v52
									}
								}
							}
						} else {
							v36 = F_sdsnew(m, int32(_a_F_rdbFunctionLoad_2))
							mBase = m.M
							v37 = m.ExcPending
							if v37 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v10)+12)) = v36
								v47 = v33
								F_sdsfree(m, v17)
								mBase = m.M
								v49 = m.ExcPending
								if v49 != 0 {
									return int32(0)
								} else {
									v50 = *(*int32)(unsafe.Add(mBase, uint32(v10)+12))
									v51 = v50
									v52 = v47
									if v51 == int32(0) {
										m.G0 = v10 + int32(16)
										return v52
									} else {
										if l4 == int32(0) {
											v59 = *(*int32)(unsafe.Add(mBase, _c_F_rdbFunctionLoad[0]))
											if int32(3) < v59 {
												v68 = v51
												F_sdsfree(m, v68)
												mBase = m.M
												v70 = m.ExcPending
												if v70 != 0 {
													return int32(0)
												} else {
													m.G0 = v10 + int32(16)
													return v52
												}
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(v10))) = v51
												F__serverLog(m, int32(3), int32(_a_F_rdbFunctionLoad_1), v10)
												mBase = m.M
												v66 = m.ExcPending
												if v66 != 0 {
													return int32(0)
												} else {
													v67 = *(*int32)(unsafe.Add(mBase, uint32(v10)+12))
													v68 = v67
													F_sdsfree(m, v68)
													mBase = m.M
													v70 = m.ExcPending
													if v70 != 0 {
														return int32(0)
													} else {
														m.G0 = v10 + int32(16)
														return v52
													}
												}
											}
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(l4))) = v51
											m.G0 = v10 + int32(16)
											return v52
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
func F_rdbLoad(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
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
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v50 int32
	_ = v50
	var v54 int32
	_ = v54
	var v60 int32
	_ = v60
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
	var v72 int64
	_ = v72
	var v74 int64
	_ = v74
	var v88 int64
	_ = v88
	var v95 int32
	_ = v95
	var v100 int32
	_ = v100
	var v104 int32
	_ = v104
	var v105 int64
	_ = v105
	var v109 int64
	_ = v109
	var v113 int64
	_ = v113
	var v116 int32
	_ = v116
	var v124 int64
	_ = v124
	var v129 int32
	_ = v129
	var v131 int32
	_ = v131
	var v136 int32
	_ = v136
	var v139 int32
	_ = v139
	var v144 int32
	_ = v144
	var v145 int64
	_ = v145
	var v154 int32
	_ = v154
	var v157 int32
	_ = v157
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v163 int32
	_ = v163
	var v166 int32
	_ = v166
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v187 int32
	_ = v187
	var v189 int32
	_ = v189
	var v194 int32
	_ = v194
	var v200 int32
	_ = v200
	var v203 int32
	_ = v203
	var v207 int32
	_ = v207
	var v210 int32
	_ = v210
	var v213 int32
	_ = v213
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	v9 = m.G0
	v11 = v9 - int32(192)
	m.G0 = v11
	v14 = F_fopen(m, l0, int32(_a_F_rdbLoad_0))
	mBase = m.M
	if v14 != 0 {
		v37 = *(*int32)(unsafe.Add(mBase, uint32(v14)+76))
		if int32(-1) < v37 {
			v41 = F___lockfile(m, v14)
			mBase = m.M
			v42 = *(*int32)(unsafe.Add(mBase, uint32(v14)+60))
			if v41 == int32(0) {
				v46 = v42
			} else {
				F___unlockfile(m, v14)
				mBase = m.M
				v46 = v42
			}
		} else {
			v40 = *(*int32)(unsafe.Add(mBase, uint32(v14)+60))
			v46 = v40
		}
		if int32(-1) < v46 {
			v54 = v46
		} else {
			v50 = F___errno_location(m)
			mBase = m.M
			*(*int32)(unsafe.Add(mBase, uint32(v50))) = int32(8)
			v54 = int32(-1)
		}
		if int32(-1) < v54 {
			v63 = F___fstatat(m, v54, int32(_a_F_rdbLoad_1), v11+int32(8), int32(4096))
			mBase = m.M
			v64 = v63
		} else {
			v60 = F___syscall_ret(m, int32(-8))
			mBase = m.M
			v64 = v60
		}
		v65 = int32(0)
		*(*int32)(unsafe.Add(mBase, _c_F_rdbLoad[0])) = l0
		v67 = int32(_a_F_rdbLoad_2)
		*(*int32)(unsafe.Add(mBase, _c_F_rdbLoad[1])) = int32(1)
		v70 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v11)+32)))
		v72 = F___time(m, v65)
		mBase = m.M
		v74 = int64(0)
		*(*int64)(unsafe.Add(mBase, _c_F_rdbLoad[2])) = v74
		*(*int64)(unsafe.Add(mBase, _c_F_rdbLoad[3])) = v72
		*(*int64)(unsafe.Add(mBase, _c_F_rdbLoad[4])) = v74
		*(*int64)(unsafe.Add(mBase, _c_F_rdbLoad[5])) = v74
		if v64 == int32(-1) {
			v88 = v74
		} else {
			v88 = v70
		}
		*(*int64)(unsafe.Add(mBase, _c_F_rdbLoad[6])) = v88
		*(*int64)(unsafe.Add(mBase, _c_F_rdbLoad[7])) = int64(0)
		v95 = int32(0)
		v100 = *(*int32)(unsafe.Add(mBase, _c_F_rdbLoad[8]))
		*(*int32)(unsafe.Add(mBase, _c_F_rdbLoad[8])) = v100 + int32(1)
		if v100 != 0 {
		} else {
			v104 = int32(0)
			v105 = F_ustime(m)
			mBase = m.M
			*(*int64)(unsafe.Add(mBase, _c_F_rdbLoad[9])) = v105
			v109 = base.I64_div_s(v105, int64(1000))
			*(*int64)(unsafe.Add(mBase, _c_F_rdbLoad[10])) = v109
			v113 = base.I64_div_s(v105, int64(1000000))
			*(*int64)(unsafe.Add(mBase, _c_F_rdbLoad[11])) = v113
			v116 = *(*int32)(unsafe.Add(mBase, _c_F_rdbLoad[12]))
			F_lrulfu_updateClockAndPolicy(m, v109, int32(base.Ui32(v116&int32(2))>>(uint(int32(1))%32)))
			mBase = m.M
			v124 = *(*int64)(unsafe.Add(mBase, _c_F_rdbLoad[10]))
			*(*int64)(unsafe.Add(mBase, _c_F_rdbLoad[13])) = v124
		}
		F_clusterCleanSlotImportsBeforeLoad(m)
		mBase = m.M
		v129 = m.ExcPending
		if v129 != 0 {
			return int32(0)
		} else {
			v131 = int32(1)
			if l2&v131 != 0 {
				v136 = v131
			} else {
				v136 = l2 & int32(2)
			}
			F_moduleFireServerEvent(m, int64(3), v136, int32(0))
			mBase = m.M
			v139 = m.ExcPending
			if v139 != 0 {
				return int32(0)
			} else {
				v144 = F___memcpy(m, v11+int32(104), int32(_a_F_rdbLoad_3), int32(80))
				mBase = m.M
				v145 = int64(0)
				*(*int64)(unsafe.Add(mBase, uint32(v144)+56)) = v145
				*(*int32)(unsafe.Add(mBase, uint32(v144)+48)) = v14
				*(*int64)(unsafe.Add(mBase, uint32(v144+int32(64)))) = v145
				v154 = int32(0)
				*(*uint8)(unsafe.Add(mBase, uint32(v144+int32(72)))) = uint8(v154)
				v157 = *(*int32)(unsafe.Add(mBase, _c_F_rdbLoad[14]))
				*(*int32)(unsafe.Add(mBase, uint32(v11)+188)) = v157
				v159 = int32(_a_F_rdbLoad_2)
				v160 = *(*int32)(unsafe.Add(mBase, _c_F_rdbLoad[15]))
				*(*int32)(unsafe.Add(mBase, uint32(v11)+184)) = v160
				v163 = *(*int32)(unsafe.Add(mBase, _c_F_rdbLoad[16]))
				v166 = v11 + int32(104)
				*(*int32)(unsafe.Add(mBase, _c_F_rdbLoad[16])) = v166
				v172 = F_rdbLoadRioWithLoadingCtx(m, v166, l2, l1, v11+int32(184))
				mBase = m.M
				v173 = m.ExcPending
				if v173 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, _c_F_rdbLoad[16])) = v163
					v176 = F_fclose(m, v14)
					mBase = m.M
					v177 = m.ExcPending
					if v177 != 0 {
						return int32(0)
					} else {
						v178 = int32(_a_F_rdbLoad_2)
						v179 = int32(0)
						*(*int32)(unsafe.Add(mBase, _c_F_rdbLoad[1])) = v179
						*(*int32)(unsafe.Add(mBase, _c_F_rdbLoad[17])) = v179
						v187 = *(*int32)(unsafe.Add(mBase, _c_F_rdbLoad[8]))
						v189 = v187 + int32(-1)
						*(*int32)(unsafe.Add(mBase, _c_F_rdbLoad[8])) = v189
						if v189 != 0 {
						} else {
							*(*int64)(unsafe.Add(mBase, _c_F_rdbLoad[13])) = int64(0)
						}
						v194 = int32(0)
						*(*int32)(unsafe.Add(mBase, _c_F_rdbLoad[0])) = v194
						if v172 != 0 {
							v200 = int32(4)
						} else {
							v200 = int32(3)
						}
						F_moduleFireServerEvent(m, int64(3), v200, int32(0))
						mBase = m.M
						v203 = m.ExcPending
						if v203 != 0 {
							return int32(0)
						} else {
							if v172|l2&int32(16) != 0 {
								v218 = v172
								m.G0 = v11 + int32(192)
								return v218
							} else {
								v207 = int32(0)
								v210 = F_open(m, l0, v207, v207)
								mBase = m.M
								if v210 < v207 {
									v218 = v207
									m.G0 = v11 + int32(192)
									return v218
								} else {
									v213 = int32(0)
									F_bioCreateCloseJob(m, v210, v213, int32(1))
									mBase = m.M
									v217 = m.ExcPending
									if v217 != 0 {
										return int32(0)
									} else {
										v218 = v213
										m.G0 = v11 + int32(192)
										return v218
									}
								}
							}
						}
					}
				}
			}
		}
	} else {
		v17 = *(*int32)(unsafe.Add(mBase, _c_F_rdbLoad[18]))
		if v17 == int32(44) {
			v218 = int32(1)
			m.G0 = v11 + int32(192)
			return v218
		} else {
			v20 = int32(3)
			v22 = *(*int32)(unsafe.Add(mBase, _c_F_rdbLoad[19]))
			if v20 < v22 {
				v218 = v20
				m.G0 = v11 + int32(192)
				return v218
			} else {
				v25 = F___strerror_l(m, v17, v17)
				mBase = m.M
				*(*int32)(unsafe.Add(mBase, uint32(v11)+4)) = v25
				*(*int32)(unsafe.Add(mBase, uint32(v11))) = l0
				v28 = int32(3)
				F__serverLog(m, v28, int32(_a_F_rdbLoad_4), v11)
				mBase = m.M
				v34 = m.ExcPending
				if v34 != 0 {
					return int32(0)
				} else {
					v218 = v28
					m.G0 = v11 + int32(192)
					return v218
				}
			}
		}
	}
}
func F_rdbLoadBinaryDoubleValue(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
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
	var v24 int32
	_ = v24
	var v25 int64
	_ = v25
	var v31 int32
	_ = v31
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	v7 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
	if v7&int32(5) != 0 {
		v44 = int32(-1)
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return v44
L2:
	;
	v12 = l1
	v13 = int32(8)
	goto L3
L3:
	;
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	if base.Ui32(v16) < base.Ui32(v13) {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	v44 = int32(0)
	goto L1
L5:
	;
	v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v31 == int32(0) {
		goto L15
	} else {
		goto L16
	}
L6:
	;
	v18 = v16
	goto L8
L7:
	;
	v18 = v13
	goto L8
L8:
	;
	if v16 != 0 {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	v19 = v18
	goto L11
L10:
	;
	v19 = v13
	goto L11
L11:
	;
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v21 = m.T0[v20].(func(*base.Module, int32, int32, int32) int32)(m, l0, v12, v19)
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	return int32(0)
L13:
	;
	if v21 != 0 {
		goto L5
	} else {
		goto L14
	}
L14:
	;
	v25 = *(*int64)(unsafe.Add(mBase, uint32(l0)+32))
	*(*int64)(unsafe.Add(mBase, uint32(l0)+32)) = v25 | int64(1)
	return int32(-1)
L15:
	;
	v36 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = v36 + v19
	v40 = v13 - v19
	if v40 != 0 {
		v12 = v12 + v19
		v13 = v40
		goto L3
	} else {
		goto L18
	}
L16:
	;
	m.T0[v31].(func(*base.Module, int32, int32, int32))(m, l0, v12, v19)
	mBase = m.M
	v35 = m.ExcPending
	if v35 != 0 {
		goto L12
	} else {
		goto L17
	}
L17:
	;
	goto L15
L18:
	;
	goto L4
}
func F_rdbLoadIntegerObject(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
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
	var v23 int32
	_ = v23
	var v28 int32
	_ = v28
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v41 int64
	_ = v41
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v54 int32
	_ = v54
	var v55 int64
	_ = v55
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v64 int32
	_ = v64
	var v69 int32
	_ = v69
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
	var v79 int32
	_ = v79
	var v80 int64
	_ = v80
	var v84 int32
	_ = v84
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v93 int32
	_ = v93
	var v94 int64
	_ = v94
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v103 int32
	_ = v103
	var v108 int32
	_ = v108
	var v112 int32
	_ = v112
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v119 int64
	_ = v119
	var v123 int32
	_ = v123
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v132 int32
	_ = v132
	var v133 int64
	_ = v133
	var v143 int64
	_ = v143
	var v151 int32
	_ = v151
	var v161 int32
	_ = v161
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v172 int64
	_ = v172
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v193 int32
	_ = v193
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v210 int32
	_ = v210
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v225 int32
	_ = v225
	var v232 int32
	_ = v232
	v11 = m.G0
	v13 = v11 - int32(48)
	m.G0 = v13
	switch l1 {
	case 0:
		goto L6
	case 1:
		goto L5
	case 2:
		goto L4
	default:
		goto L2
	}
L1:
	;
	m.G0 = v13 + int32(48)
	return v232
L2:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13))) = l1
	F_rdbReportError(m, int32(1), int32(366), int32(_a_F_rdbLoadIntegerObject_0), v13)
	mBase = m.M
	v225 = m.ExcPending
	if v225 != 0 {
		goto L17
	} else {
		goto L81
	}
L3:
	;
	if l2&int32(6) == int32(0) {
		goto L56
	} else {
		goto L57
	}
L4:
	;
	v95 = int32(0)
	v96 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
	if v96&int32(5) != 0 {
		v232 = v95
		goto L1
	} else {
		goto L40
	}
L5:
	;
	v56 = int32(0)
	v57 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
	if v57&int32(5) != 0 {
		v232 = v56
		goto L1
	} else {
		goto L24
	}
L6:
	;
	v15 = int32(0)
	v16 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
	if v16&int32(5) != 0 {
		v232 = v15
		goto L1
	} else {
		goto L7
	}
L7:
	;
	v23 = int32(1)
	v28 = v13 + int32(44)
	goto L8
L8:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	if base.Ui32(v32) < base.Ui32(v23) {
		goto L11
	} else {
		goto L12
	}
L9:
	;
	v55 = int64(*(*int8)(unsafe.Add(mBase, uint32(v13)+44)))
	v143 = v55
	goto L3
L10:
	;
	v45 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v45 == int32(0) {
		goto L20
	} else {
		goto L21
	}
L11:
	;
	v34 = v32
	goto L13
L12:
	;
	v34 = v23
	goto L13
L13:
	;
	if v32 != 0 {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v35 = v34
	goto L16
L15:
	;
	v35 = v23
	goto L16
L16:
	;
	v36 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v37 = m.T0[v36].(func(*base.Module, int32, int32, int32) int32)(m, l0, v28, v35)
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	return int32(0)
L18:
	;
	if v37 != 0 {
		goto L10
	} else {
		goto L19
	}
L19:
	;
	v41 = *(*int64)(unsafe.Add(mBase, uint32(l0)+32))
	*(*int64)(unsafe.Add(mBase, uint32(l0)+32)) = v41 | int64(1)
	v232 = v15
	goto L1
L20:
	;
	v50 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = v50 + v35
	v54 = v23 - v35
	if v54 != 0 {
		v23 = v54
		v28 = v28 + v35
		goto L8
	} else {
		goto L23
	}
L21:
	;
	m.T0[v45].(func(*base.Module, int32, int32, int32))(m, l0, v28, v35)
	mBase = m.M
	v49 = m.ExcPending
	if v49 != 0 {
		goto L17
	} else {
		goto L22
	}
L22:
	;
	goto L20
L23:
	;
	goto L9
L24:
	;
	v64 = int32(2)
	v69 = v13 + int32(44)
	goto L25
L25:
	;
	v73 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	if base.Ui32(v73) < base.Ui32(v64) {
		goto L28
	} else {
		goto L29
	}
L26:
	;
	v94 = int64(*(*int16)(unsafe.Add(mBase, uint32(v13)+44)))
	v143 = v94
	goto L3
L27:
	;
	v84 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v84 == int32(0) {
		goto L36
	} else {
		goto L37
	}
L28:
	;
	v75 = v73
	goto L30
L29:
	;
	v75 = v64
	goto L30
L30:
	;
	if v73 != 0 {
		goto L31
	} else {
		goto L32
	}
L31:
	;
	v76 = v75
	goto L33
L32:
	;
	v76 = v64
	goto L33
L33:
	;
	v77 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v78 = m.T0[v77].(func(*base.Module, int32, int32, int32) int32)(m, l0, v69, v76)
	mBase = m.M
	v79 = m.ExcPending
	if v79 != 0 {
		goto L17
	} else {
		goto L34
	}
L34:
	;
	if v78 != 0 {
		goto L27
	} else {
		goto L35
	}
L35:
	;
	v80 = *(*int64)(unsafe.Add(mBase, uint32(l0)+32))
	*(*int64)(unsafe.Add(mBase, uint32(l0)+32)) = v80 | int64(1)
	v232 = v56
	goto L1
L36:
	;
	v89 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = v89 + v76
	v93 = v64 - v76
	if v93 != 0 {
		v64 = v93
		v69 = v69 + v76
		goto L25
	} else {
		goto L39
	}
L37:
	;
	m.T0[v84].(func(*base.Module, int32, int32, int32))(m, l0, v69, v76)
	mBase = m.M
	v88 = m.ExcPending
	if v88 != 0 {
		goto L17
	} else {
		goto L38
	}
L38:
	;
	goto L36
L39:
	;
	goto L26
L40:
	;
	v103 = int32(4)
	v108 = v13 + int32(44)
	goto L41
L41:
	;
	v112 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	if base.Ui32(v112) < base.Ui32(v103) {
		goto L44
	} else {
		goto L45
	}
L42:
	;
	v133 = int64(*(*int32)(unsafe.Add(mBase, uint32(v13)+44)))
	v143 = v133
	goto L3
L43:
	;
	v123 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v123 == int32(0) {
		goto L52
	} else {
		goto L53
	}
L44:
	;
	v114 = v112
	goto L46
L45:
	;
	v114 = v103
	goto L46
L46:
	;
	if v112 != 0 {
		goto L47
	} else {
		goto L48
	}
L47:
	;
	v115 = v114
	goto L49
L48:
	;
	v115 = v103
	goto L49
L49:
	;
	v116 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v117 = m.T0[v116].(func(*base.Module, int32, int32, int32) int32)(m, l0, v108, v115)
	mBase = m.M
	v118 = m.ExcPending
	if v118 != 0 {
		goto L17
	} else {
		goto L50
	}
L50:
	;
	if v117 != 0 {
		goto L43
	} else {
		goto L51
	}
L51:
	;
	v119 = *(*int64)(unsafe.Add(mBase, uint32(l0)+32))
	*(*int64)(unsafe.Add(mBase, uint32(l0)+32)) = v119 | int64(1)
	v232 = v95
	goto L1
L52:
	;
	v128 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = v128 + v115
	v132 = v103 - v115
	if v132 != 0 {
		v103 = v132
		v108 = v108 + v115
		goto L41
	} else {
		goto L55
	}
L53:
	;
	m.T0[v123].(func(*base.Module, int32, int32, int32))(m, l0, v108, v115)
	mBase = m.M
	v127 = m.ExcPending
	if v127 != 0 {
		goto L17
	} else {
		goto L54
	}
L54:
	;
	goto L52
L55:
	;
	goto L42
L56:
	;
	if l2&int32(1) == int32(0) {
		goto L77
	} else {
		goto L78
	}
L57:
	;
	v151 = v13 + int32(16)
	if v143 <= int64(-1) {
		goto L62
	} else {
		goto L63
	}
L58:
	;
	if l3 == int32(0) {
		goto L67
	} else {
		goto L68
	}
L59:
	;
	v193 = int32(0)
	goto L58
L61:
	;
	v174 = F_ull2string(m, v170, v171, v172)
	mBase = m.M
	if v174 == int32(0) {
		goto L59
	} else {
		goto L65
	}
L62:
	;
	goto L64
L63:
	;
	v170 = v151
	v171 = int32(21)
	v172 = v143
	v173 = int32(0)
	goto L61
L64:
	;
	v161 = int32(45)
	*(*uint8)(unsafe.Add(mBase, uint32(v151))) = uint8(v161)
	v170 = v13 + int32(17)
	v171 = int32(20)
	v172 = int64(0) - v143
	v173 = int32(1)
	goto L61
L65:
	;
	v193 = v174 + v173
	goto L58
L67:
	;
	if l2&int32(2) == int32(0) {
		goto L70
	} else {
		goto L71
	}
L68:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v193
	goto L67
L69:
	;
	if v193 == int32(0) {
		goto L75
	} else {
		goto L76
	}
L70:
	;
	v202 = *(*int32)(unsafe.Add(mBase, _c_F_rdbLoadIntegerObject[0]))
	v203 = F_sdsnewlen(m, v202, v193)
	mBase = m.M
	v204 = m.ExcPending
	if v204 != 0 {
		goto L17
	} else {
		goto L73
	}
L71:
	;
	v199 = F_valkey_malloc(m, v193)
	mBase = m.M
	v200 = m.ExcPending
	if v200 != 0 {
		goto L17
	} else {
		goto L72
	}
L72:
	;
	v205 = v199
	goto L69
L73:
	;
	v205 = v203
	goto L69
L74:
	;
	v232 = v205
	goto L1
L75:
	;
	goto L74
L76:
	;
	v210 = F__emscripten_memcpy_bulkmem(m, v205, v13+int32(16), v193)
	mBase = m.M
	goto L75
L77:
	;
	v218 = F_createStringObjectFromLongLongWithSds(m, v143)
	mBase = m.M
	v219 = m.ExcPending
	if v219 != 0 {
		goto L17
	} else {
		goto L80
	}
L78:
	;
	v216 = F_createStringObjectFromLongLongForValue(m, v143)
	mBase = m.M
	v217 = m.ExcPending
	if v217 != 0 {
		goto L17
	} else {
		goto L79
	}
L79:
	;
	v232 = v216
	goto L1
L80:
	;
	v232 = v218
	goto L1
L81:
	;
	v232 = int32(0)
	goto L1
}
func F_rdbLoadLenByRef(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v31 int32
	_ = v31
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
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v45 int64
	_ = v45
	var v49 int32
	_ = v49
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v60 int64
	_ = v60
	var v62 int32
	_ = v62
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v104 int64
	_ = v104
	var v108 int32
	_ = v108
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v117 int32
	_ = v117
	var v118 int64
	_ = v118
	var v123 int64
	_ = v123
	var v129 int32
	_ = v129
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v145 int32
	_ = v145
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v152 int64
	_ = v152
	var v156 int32
	_ = v156
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v171 int32
	_ = v171
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v187 int32
	_ = v187
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v194 int64
	_ = v194
	var v198 int32
	_ = v198
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v207 int32
	_ = v207
	var v208 int64
	_ = v208
	var v209 int64
	_ = v209
	var v211 int64
	_ = v211
	var v213 int64
	_ = v213
	var v216 int64
	_ = v216
	var v218 int64
	_ = v218
	var v220 int64
	_ = v220
	var v222 int64
	_ = v222
	var v252 int32
	_ = v252
	var v257 int32
	_ = v257
	v11 = m.G0
	v13 = v11 - int32(32)
	m.G0 = v13
	if l1 == int32(0) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v19 = int32(-1)
	v20 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
	if v20&int32(5) != 0 {
		v257 = v19
		goto L3
	} else {
		goto L4
	}
L2:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = int32(0)
	goto L1
L3:
	;
	m.G0 = v13 + int32(32)
	return v257
L4:
	;
	v31 = v13 + int32(30)
	v32 = int32(1)
	goto L5
L5:
	;
	v36 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	if base.Ui32(v36) < base.Ui32(v32) {
		goto L8
	} else {
		goto L9
	}
L6:
	;
	v59 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+30)))
	v60 = base.I64_extend_i32_u(v59)
	v62 = int32(base.Ui32(v59) >> (uint(int32(6)) % 32))
	if v62 != int32(3) {
		goto L21
	} else {
		goto L22
	}
L7:
	;
	v49 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v49 == int32(0) {
		goto L17
	} else {
		goto L18
	}
L8:
	;
	v38 = v36
	goto L10
L9:
	;
	v38 = v32
	goto L10
L10:
	;
	if v36 != 0 {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v39 = v38
	goto L13
L12:
	;
	v39 = v32
	goto L13
L13:
	;
	v40 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v41 = m.T0[v40].(func(*base.Module, int32, int32, int32) int32)(m, l0, v31, v39)
	mBase = m.M
	v44 = m.ExcPending
	if v44 != 0 {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	return int32(0)
L15:
	;
	if v41 != 0 {
		goto L7
	} else {
		goto L16
	}
L16:
	;
	v45 = *(*int64)(unsafe.Add(mBase, uint32(l0)+32))
	*(*int64)(unsafe.Add(mBase, uint32(l0)+32)) = v45 | int64(1)
	v257 = v19
	goto L3
L17:
	;
	v54 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = v54 + v39
	v58 = v32 - v39
	if v58 != 0 {
		v31 = v31 + v39
		v32 = v58
		goto L5
	} else {
		goto L20
	}
L18:
	;
	m.T0[v49].(func(*base.Module, int32, int32, int32))(m, l0, v31, v39)
	mBase = m.M
	v53 = m.ExcPending
	if v53 != 0 {
		goto L14
	} else {
		goto L19
	}
L19:
	;
	goto L17
L20:
	;
	goto L6
L21:
	;
	if base.Ui32(int32(63)) < base.Ui32(v59) {
		goto L25
	} else {
		goto L26
	}
L22:
	;
	if l1 == int32(0) {
		goto L23
	} else {
		goto L24
	}
L23:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l2))) = v60 & int64(63)
	v257 = int32(0)
	goto L3
L24:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = int32(1)
	goto L23
L25:
	;
	v77 = int32(1)
	if v62 != v77 {
		goto L27
	} else {
		goto L28
	}
L26:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l2))) = v60
	v257 = int32(0)
	goto L3
L27:
	;
	switch v59 + int32(-128) {
	case 0:
		goto L47
	case 1:
		goto L46
	default:
		goto L45
	}
L28:
	;
	v80 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
	if v80&int32(5) != 0 {
		v257 = v19
		goto L3
	} else {
		goto L29
	}
L29:
	;
	v92 = v13 + int32(31)
	v93 = v77
	goto L30
L30:
	;
	v97 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	if base.Ui32(v97) < base.Ui32(v93) {
		goto L33
	} else {
		goto L34
	}
L31:
	;
	v118 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v13)+30)))
	v123 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v13)+31)))
	*(*int64)(unsafe.Add(mBase, uint32(l2))) = v118&int64(63)<<(uint(int64(8))%64) | v123
	v257 = int32(0)
	goto L3
L32:
	;
	v108 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v108 == int32(0) {
		goto L41
	} else {
		goto L42
	}
L33:
	;
	v99 = v97
	goto L35
L34:
	;
	v99 = v93
	goto L35
L35:
	;
	if v97 != 0 {
		goto L36
	} else {
		goto L37
	}
L36:
	;
	v100 = v99
	goto L38
L37:
	;
	v100 = v93
	goto L38
L38:
	;
	v101 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v102 = m.T0[v101].(func(*base.Module, int32, int32, int32) int32)(m, l0, v92, v100)
	mBase = m.M
	v103 = m.ExcPending
	if v103 != 0 {
		goto L14
	} else {
		goto L39
	}
L39:
	;
	if v102 != 0 {
		goto L32
	} else {
		goto L40
	}
L40:
	;
	v104 = *(*int64)(unsafe.Add(mBase, uint32(l0)+32))
	*(*int64)(unsafe.Add(mBase, uint32(l0)+32)) = v104 | int64(1)
	v257 = v19
	goto L3
L41:
	;
	v113 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = v113 + v100
	v117 = v93 - v100
	if v117 != 0 {
		v92 = v92 + v100
		v93 = v117
		goto L30
	} else {
		goto L44
	}
L42:
	;
	m.T0[v108].(func(*base.Module, int32, int32, int32))(m, l0, v92, v100)
	mBase = m.M
	v112 = m.ExcPending
	if v112 != 0 {
		goto L14
	} else {
		goto L43
	}
L43:
	;
	goto L41
L44:
	;
	goto L31
L45:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13))) = int32(2)
	F_rdbReportError(m, int32(1), int32(299), int32(_a_F_rdbLoadLenByRef_0), v13)
	mBase = m.M
	v252 = m.ExcPending
	if v252 != 0 {
		goto L14
	} else {
		goto L81
	}
L46:
	;
	v171 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
	if v171&int32(5) != 0 {
		v257 = v19
		goto L3
	} else {
		goto L65
	}
L47:
	;
	v129 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
	if v129&int32(5) != 0 {
		v257 = v19
		goto L3
	} else {
		goto L48
	}
L48:
	;
	v140 = v13 + int32(16)
	v141 = int32(4)
	goto L49
L49:
	;
	v145 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	if base.Ui32(v145) < base.Ui32(v141) {
		goto L52
	} else {
		goto L53
	}
L50:
	;
	v166 = *(*int32)(unsafe.Add(mBase, uint32(v13)+16))
	v167 = F___bswap_32_2(m, v166)
	mBase = m.M
	goto L64
L51:
	;
	v156 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v156 == int32(0) {
		goto L60
	} else {
		goto L61
	}
L52:
	;
	v147 = v145
	goto L54
L53:
	;
	v147 = v141
	goto L54
L54:
	;
	if v145 != 0 {
		goto L55
	} else {
		goto L56
	}
L55:
	;
	v148 = v147
	goto L57
L56:
	;
	v148 = v141
	goto L57
L57:
	;
	v149 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v150 = m.T0[v149].(func(*base.Module, int32, int32, int32) int32)(m, l0, v140, v148)
	mBase = m.M
	v151 = m.ExcPending
	if v151 != 0 {
		goto L14
	} else {
		goto L58
	}
L58:
	;
	if v150 != 0 {
		goto L51
	} else {
		goto L59
	}
L59:
	;
	v152 = *(*int64)(unsafe.Add(mBase, uint32(l0)+32))
	*(*int64)(unsafe.Add(mBase, uint32(l0)+32)) = v152 | int64(1)
	v257 = v19
	goto L3
L60:
	;
	v161 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = v161 + v148
	v165 = v141 - v148
	if v165 != 0 {
		v140 = v140 + v148
		v141 = v165
		goto L49
	} else {
		goto L63
	}
L61:
	;
	m.T0[v156].(func(*base.Module, int32, int32, int32))(m, l0, v140, v148)
	mBase = m.M
	v160 = m.ExcPending
	if v160 != 0 {
		goto L14
	} else {
		goto L62
	}
L62:
	;
	goto L60
L63:
	;
	goto L50
L64:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l2))) = base.I64_extend_i32_u(v167)
	v257 = int32(0)
	goto L3
L65:
	;
	v182 = v13 + int32(16)
	v183 = int32(8)
	goto L66
L66:
	;
	v187 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	if base.Ui32(v187) < base.Ui32(v183) {
		goto L69
	} else {
		goto L70
	}
L67:
	;
	v208 = *(*int64)(unsafe.Add(mBase, uint32(v13)+16))
	v209 = int64(56)
	v211 = int64(65280)
	v213 = int64(40)
	v216 = int64(16711680)
	v218 = int64(24)
	v220 = int64(4278190080)
	v222 = int64(8)
	*(*int64)(unsafe.Add(mBase, uint32(l2))) = v208<<(uint(v209)%64) | v208&v211<<(uint(v213)%64) | (v208&v216<<(uint(v218)%64) | v208&v220<<(uint(v222)%64)) | (int64(base.Ui64(v208)>>(uint(v222)%64))&v220 | int64(base.Ui64(v208)>>(uint(v218)%64))&v216 | (int64(base.Ui64(v208)>>(uint(v213)%64))&v211 | int64(base.Ui64(v208)>>(uint(v209)%64))))
	v257 = int32(0)
	goto L3
L68:
	;
	v198 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v198 == int32(0) {
		goto L77
	} else {
		goto L78
	}
L69:
	;
	v189 = v187
	goto L71
L70:
	;
	v189 = v183
	goto L71
L71:
	;
	if v187 != 0 {
		goto L72
	} else {
		goto L73
	}
L72:
	;
	v190 = v189
	goto L74
L73:
	;
	v190 = v183
	goto L74
L74:
	;
	v191 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v192 = m.T0[v191].(func(*base.Module, int32, int32, int32) int32)(m, l0, v182, v190)
	mBase = m.M
	v193 = m.ExcPending
	if v193 != 0 {
		goto L14
	} else {
		goto L75
	}
L75:
	;
	if v192 != 0 {
		goto L68
	} else {
		goto L76
	}
L76:
	;
	v194 = *(*int64)(unsafe.Add(mBase, uint32(l0)+32))
	*(*int64)(unsafe.Add(mBase, uint32(l0)+32)) = v194 | int64(1)
	v257 = v19
	goto L3
L77:
	;
	v203 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = v203 + v190
	v207 = v183 - v190
	if v207 != 0 {
		v182 = v182 + v190
		v183 = v207
		goto L66
	} else {
		goto L80
	}
L78:
	;
	m.T0[v198].(func(*base.Module, int32, int32, int32))(m, l0, v182, v190)
	mBase = m.M
	v202 = m.ExcPending
	if v202 != 0 {
		goto L14
	} else {
		goto L79
	}
L79:
	;
	goto L77
L80:
	;
	goto L67
L81:
	;
	v257 = v19
	goto L3
}
func F_rdbLoadLzfStringObject(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v29 int64
	_ = v29
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v40 int64
	_ = v40
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
	var v56 int32
	_ = v56
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v91 int32
	_ = v91
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v103 int64
	_ = v103
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
	var v114 int32
	_ = v114
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v125 int32
	_ = v125
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v133 int32
	_ = v133
	var v136 int32
	_ = v136
	var v138 int32
	_ = v138
	var v144 int32
	_ = v144
	var v147 int32
	_ = v147
	var v149 int32
	_ = v149
	var v153 int32
	_ = v153
	var v156 int32
	_ = v156
	var v163 int32
	_ = v163
	var v167 int32
	_ = v167
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v176 int32
	_ = v176
	var v181 int64
	_ = v181
	var v184 int32
	_ = v184
	var v186 int32
	_ = v186
	var v188 int32
	_ = v188
	var v194 int32
	_ = v194
	var v199 int32
	_ = v199
	var v206 int32
	_ = v206
	var v215 int32
	_ = v215
	var v218 int32
	_ = v218
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v225 int64
	_ = v225
	var v229 int32
	_ = v229
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
	var v238 int32
	_ = v238
	var v254 int32
	_ = v254
	var v262 int32
	_ = v262
	var v264 int32
	_ = v264
	var v270 int32
	_ = v270
	var v271 int32
	_ = v271
	var v285 int32
	_ = v285
	var v290 int32
	_ = v290
	var v294 int32
	_ = v294
	var v296 int32
	_ = v296
	var v301 int32
	_ = v301
	v4 = int32(0)
	v15 = m.G0
	v17 = v15 - int32(32)
	m.G0 = v17
	v23 = F_rdbLoadLenByRef(m, l0, v4, v17+int32(24))
	mBase = m.M
	v26 = m.ExcPending
	if v26 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v17 + int32(32)
	return v301
L2:
	;
	return int32(0)
L3:
	;
	if v23 == int32(-1) {
		v301 = v4
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v29 = *(*int64)(unsafe.Add(mBase, uint32(v17)+24))
	if v29 == int64(-1) {
		v301 = v4
		goto L1
	} else {
		goto L5
	}
L5:
	;
	v32 = int32(0)
	v36 = F_rdbLoadLenByRef(m, l0, v32, v17+int32(24))
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		goto L2
	} else {
		goto L6
	}
L6:
	;
	if v36 == int32(-1) {
		v301 = v32
		goto L1
	} else {
		goto L7
	}
L7:
	;
	v40 = *(*int64)(unsafe.Add(mBase, uint32(v17)+24))
	if v40 == int64(-1) {
		v301 = v32
		goto L1
	} else {
		goto L8
	}
L8:
	;
	v44 = l1 & int32(2)
	v45 = base.I32_wrap_i64(v29)
	v46 = int32(0)
	if base.Ui32(int32(2147483646)) < base.Ui32(v45) {
		v91 = v46
		goto L13
	} else {
		goto L14
	}
L9:
	;
	F_valkey_free(m, v91)
	mBase = m.M
	v290 = m.ExcPending
	if v290 != 0 {
		goto L2
	} else {
		goto L88
	}
L10:
	;
	v285 = int32(0)
	goto L9
L11:
	;
	if v44 == int32(0) {
		goto L35
	} else {
		goto L36
	}
L12:
	;
	if v91 != 0 {
		goto L11
	} else {
		goto L24
	}
L13:
	;
	goto L12
L14:
	;
	if v45 != 0 {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	v53 = v45
	goto L17
L16:
	;
	v53 = int32(4)
	goto L17
L17:
	;
	v55 = v53 + int32(8)
	v56 = F_emscripten_builtin_malloc(m, v55)
	mBase = m.M
	if v56 == int32(0) {
		v91 = v46
		goto L13
	} else {
		goto L18
	}
L18:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v56))) = v53
	v61 = *(*int32)(unsafe.Add(mBase, _c_F_rdbLoadLzfStringObject[0]))
	if v61 != int32(-1) {
		v72 = v61
		goto L19
	} else {
		goto L20
	}
L19:
	;
	if v72 < int32(260) {
		goto L22
	} else {
		goto L23
	}
L20:
	;
	v64 = int32(0)
	v66 = *(*int32)(unsafe.Add(mBase, _c_F_rdbLoadLzfStringObject[1]))
	*(*int32)(unsafe.Add(mBase, _c_F_rdbLoadLzfStringObject[0])) = v66
	*(*int32)(unsafe.Add(mBase, _c_F_rdbLoadLzfStringObject[1])) = v66 + int32(1)
	v72 = v66
	goto L19
L21:
	;
	v91 = v56 + int32(8)
	goto L13
L22:
	;
	v81 = v72 << (uint(int32(2)) % 32)
	v84 = *(*int32)(unsafe.Add(mBase, uint32(v81)+uint32(_c_F_rdbLoadLzfStringObject[2])))
	*(*int32)(unsafe.Add(mBase, uint32(v81)+uint32(_c_F_rdbLoadLzfStringObject[2]))) = v84 + v55
	goto L21
L23:
	;
	v75 = int32(0)
	v77 = *(*int32)(unsafe.Add(mBase, _c_F_rdbLoadLzfStringObject[3]))
	*(*int32)(unsafe.Add(mBase, _c_F_rdbLoadLzfStringObject[3])) = v77 + v55
	goto L21
L24:
	;
	v95 = *(*int32)(unsafe.Add(mBase, _c_F_rdbLoadLzfStringObject[4]))
	if v95 != 0 {
		goto L26
	} else {
		goto L27
	}
L25:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v17))) = v29
	F__serverLog(m, v110, int32(_a_F_rdbLoadLzfStringObject_0), v17)
	mBase = m.M
	v114 = m.ExcPending
	if v114 != 0 {
		goto L2
	} else {
		goto L33
	}
L26:
	;
	v103 = *(*int64)(unsafe.Add(mBase, uint32(v95)))
	if v103 == int64(-1) {
		goto L29
	} else {
		goto L30
	}
L27:
	;
	v96 = int32(3)
	v98 = *(*int32)(unsafe.Add(mBase, _c_F_rdbLoadLzfStringObject[5]))
	if v98 <= v96 {
		v110 = v96
		goto L25
	} else {
		goto L28
	}
L28:
	;
	goto L10
L29:
	;
	v106 = int32(3)
	goto L31
L30:
	;
	v106 = int32(1)
	goto L31
L31:
	;
	v108 = *(*int32)(unsafe.Add(mBase, _c_F_rdbLoadLzfStringObject[5]))
	if v106 < v108 {
		goto L10
	} else {
		goto L32
	}
L32:
	;
	v110 = v106
	goto L25
L33:
	;
	goto L10
L34:
	;
	if v171 != 0 {
		goto L50
	} else {
		goto L51
	}
L35:
	;
	v167 = *(*int32)(unsafe.Add(mBase, _c_F_rdbLoadLzfStringObject[6]))
	v169 = F_sdstrynewlen(m, v167, base.I32_wrap_i64(v40))
	mBase = m.M
	v170 = m.ExcPending
	if v170 != 0 {
		goto L2
	} else {
		goto L49
	}
L36:
	;
	v117 = base.I32_wrap_i64(v40)
	v118 = int32(0)
	if base.Ui32(int32(2147483646)) < base.Ui32(v117) {
		v163 = v118
		goto L38
	} else {
		goto L39
	}
L37:
	;
	v171 = v163
	goto L34
L38:
	;
	goto L37
L39:
	;
	if v117 != 0 {
		goto L40
	} else {
		goto L41
	}
L40:
	;
	v125 = v117
	goto L42
L41:
	;
	v125 = int32(4)
	goto L42
L42:
	;
	v127 = v125 + int32(8)
	v128 = F_emscripten_builtin_malloc(m, v127)
	mBase = m.M
	if v128 == int32(0) {
		v163 = v118
		goto L38
	} else {
		goto L43
	}
L43:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v128))) = v125
	v133 = *(*int32)(unsafe.Add(mBase, _c_F_rdbLoadLzfStringObject[0]))
	if v133 != int32(-1) {
		v144 = v133
		goto L44
	} else {
		goto L45
	}
L44:
	;
	if v144 < int32(260) {
		goto L47
	} else {
		goto L48
	}
L45:
	;
	v136 = int32(0)
	v138 = *(*int32)(unsafe.Add(mBase, _c_F_rdbLoadLzfStringObject[1]))
	*(*int32)(unsafe.Add(mBase, _c_F_rdbLoadLzfStringObject[0])) = v138
	*(*int32)(unsafe.Add(mBase, _c_F_rdbLoadLzfStringObject[1])) = v138 + int32(1)
	v144 = v138
	goto L44
L46:
	;
	v163 = v128 + int32(8)
	goto L38
L47:
	;
	v153 = v144 << (uint(int32(2)) % 32)
	v156 = *(*int32)(unsafe.Add(mBase, uint32(v153)+uint32(_c_F_rdbLoadLzfStringObject[2])))
	*(*int32)(unsafe.Add(mBase, uint32(v153)+uint32(_c_F_rdbLoadLzfStringObject[2]))) = v156 + v127
	goto L46
L48:
	;
	v147 = int32(0)
	v149 = *(*int32)(unsafe.Add(mBase, _c_F_rdbLoadLzfStringObject[3]))
	*(*int32)(unsafe.Add(mBase, _c_F_rdbLoadLzfStringObject[3])) = v149 + v127
	goto L46
L49:
	;
	v171 = v169
	goto L34
L50:
	;
	if l2 == int32(0) {
		goto L61
	} else {
		goto L62
	}
L51:
	;
	v173 = *(*int32)(unsafe.Add(mBase, _c_F_rdbLoadLzfStringObject[4]))
	if v173 != 0 {
		goto L53
	} else {
		goto L54
	}
L52:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v17)+16)) = v40
	F__serverLog(m, v188, int32(_a_F_rdbLoadLzfStringObject_0), v17+int32(16))
	mBase = m.M
	v194 = m.ExcPending
	if v194 != 0 {
		goto L2
	} else {
		goto L60
	}
L53:
	;
	v181 = *(*int64)(unsafe.Add(mBase, uint32(v173)))
	if v181 == int64(-1) {
		goto L56
	} else {
		goto L57
	}
L54:
	;
	v174 = int32(3)
	v176 = *(*int32)(unsafe.Add(mBase, _c_F_rdbLoadLzfStringObject[5]))
	if v174 < v176 {
		goto L10
	} else {
		goto L55
	}
L55:
	;
	v188 = v174
	goto L52
L56:
	;
	v184 = int32(3)
	goto L58
L57:
	;
	v184 = int32(1)
	goto L58
L58:
	;
	v186 = *(*int32)(unsafe.Add(mBase, _c_F_rdbLoadLzfStringObject[5]))
	if v184 < v186 {
		goto L10
	} else {
		goto L59
	}
L59:
	;
	v188 = v184
	goto L52
L60:
	;
	v285 = int32(0)
	goto L9
L61:
	;
	v199 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
	if v199&int32(5) != 0 {
		v285 = v171
		goto L9
	} else {
		goto L63
	}
L62:
	;
	*(*uint32)(unsafe.Add(mBase, uint32(l2))) = uint32(v40)
	goto L61
L63:
	;
	if v45 == int32(0) {
		goto L64
	} else {
		goto L65
	}
L64:
	;
	v254 = F_lzf_decompress(m, v91, v45, v171, base.I32_wrap_i64(v40))
	mBase = m.M
	if v40 == base.I64_extend_i32_u(v254) {
		goto L81
	} else {
		goto L82
	}
L65:
	;
	v206 = v45
	v215 = v91
	goto L66
L66:
	;
	v218 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	if base.Ui32(v218) < base.Ui32(v206) {
		goto L69
	} else {
		goto L70
	}
L67:
	;
	goto L64
L68:
	;
	v229 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v229 == int32(0) {
		goto L77
	} else {
		goto L78
	}
L69:
	;
	v220 = v218
	goto L71
L70:
	;
	v220 = v206
	goto L71
L71:
	;
	if v218 != 0 {
		goto L72
	} else {
		goto L73
	}
L72:
	;
	v221 = v220
	goto L74
L73:
	;
	v221 = v206
	goto L74
L74:
	;
	v222 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v223 = m.T0[v222].(func(*base.Module, int32, int32, int32) int32)(m, l0, v215, v221)
	mBase = m.M
	v224 = m.ExcPending
	if v224 != 0 {
		goto L2
	} else {
		goto L75
	}
L75:
	;
	if v223 != 0 {
		goto L68
	} else {
		goto L76
	}
L76:
	;
	v225 = *(*int64)(unsafe.Add(mBase, uint32(l0)+32))
	*(*int64)(unsafe.Add(mBase, uint32(l0)+32)) = v225 | int64(1)
	v285 = v171
	goto L9
L77:
	;
	v234 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = v234 + v221
	v238 = v206 - v221
	if v238 != 0 {
		v206 = v238
		v215 = v215 + v221
		goto L66
	} else {
		goto L80
	}
L78:
	;
	m.T0[v229].(func(*base.Module, int32, int32, int32))(m, l0, v215, v221)
	mBase = m.M
	v233 = m.ExcPending
	if v233 != 0 {
		goto L2
	} else {
		goto L79
	}
L79:
	;
	goto L77
L80:
	;
	goto L67
L81:
	;
	F_valkey_free(m, v91)
	mBase = m.M
	v264 = m.ExcPending
	if v264 != 0 {
		goto L2
	} else {
		goto L84
	}
L82:
	;
	F_rdbReportError(m, int32(1), int32(474), int32(_a_F_rdbLoadLzfStringObject_1), int32(0))
	mBase = m.M
	v262 = m.ExcPending
	if v262 != 0 {
		goto L2
	} else {
		goto L83
	}
L83:
	;
	v285 = v171
	goto L9
L84:
	;
	if l1&int32(6) == int32(0) {
		goto L85
	} else {
		goto L86
	}
L85:
	;
	v270 = F_createObject(m, int32(0), v171)
	mBase = m.M
	v271 = m.ExcPending
	if v271 != 0 {
		goto L2
	} else {
		goto L87
	}
L86:
	;
	v301 = v171
	goto L1
L87:
	;
	v301 = v270
	goto L1
L88:
	;
	if v44 == int32(0) {
		goto L89
	} else {
		goto L90
	}
L89:
	;
	F_sdsfree(m, v285)
	mBase = m.M
	v296 = m.ExcPending
	if v296 != 0 {
		goto L2
	} else {
		goto L92
	}
L90:
	;
	F_valkey_free(m, v285)
	mBase = m.M
	v294 = m.ExcPending
	if v294 != 0 {
		goto L2
	} else {
		goto L91
	}
L91:
	;
	v301 = v32
	goto L1
L92:
	;
	v301 = v32
	goto L1
}
func F_rdbLoadObjectType(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v32 int64
	_ = v32
	var v36 int32
	_ = v36
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v65 int32
	_ = v65
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	v11 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
	if v11&int32(5) != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v9 + int32(16)
	return v65
L2:
	;
	v65 = int32(-1)
	goto L1
L3:
	;
	v19 = v9 + int32(15)
	v20 = int32(1)
	goto L4
L4:
	;
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	if base.Ui32(v23) < base.Ui32(v20) {
		goto L7
	} else {
		goto L8
	}
L5:
	;
	v46 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+15)))
	if base.Ui32((v46+int32(-9))&int32(255)) < base.Ui32(int32(14)) {
		goto L20
	} else {
		goto L21
	}
L6:
	;
	v36 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v36 == int32(0) {
		goto L16
	} else {
		goto L17
	}
L7:
	;
	v25 = v23
	goto L9
L8:
	;
	v25 = v20
	goto L9
L9:
	;
	if v23 != 0 {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	v26 = v25
	goto L12
L11:
	;
	v26 = v20
	goto L12
L12:
	;
	v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v28 = m.T0[v27].(func(*base.Module, int32, int32, int32) int32)(m, l0, v19, v26)
	mBase = m.M
	v31 = m.ExcPending
	if v31 != 0 {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	return int32(0)
L14:
	;
	if v28 != 0 {
		goto L6
	} else {
		goto L15
	}
L15:
	;
	v32 = *(*int64)(unsafe.Add(mBase, uint32(l0)+32))
	*(*int64)(unsafe.Add(mBase, uint32(l0)+32)) = v32 | int64(1)
	goto L2
L16:
	;
	v41 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = v41 + v26
	v45 = v20 - v26
	if v45 != 0 {
		v19 = v19 + v26
		v20 = v45
		goto L4
	} else {
		goto L19
	}
L17:
	;
	m.T0[v36].(func(*base.Module, int32, int32, int32))(m, l0, v19, v26)
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
		goto L13
	} else {
		goto L18
	}
L18:
	;
	goto L16
L19:
	;
	goto L5
L20:
	;
	v54 = v46
	goto L22
L21:
	;
	v54 = int32(-1)
	goto L22
L22:
	;
	if base.Ui32(v46) < base.Ui32(int32(8)) {
		goto L23
	} else {
		goto L24
	}
L23:
	;
	v57 = v46
	goto L25
L24:
	;
	v57 = v54
	goto L25
L25:
	;
	v65 = v57
	goto L1
}
func F_rdbLoadRio(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
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
	var v14 int32
	_ = v14
	var v19 int32
	_ = v19
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	v12 = *(*int32)(unsafe.Add(mBase, _c_F_rdbLoadRio[0]))
	v13 = int32(_a_F_rdbLoadRio_0)
	v14 = *(*int32)(unsafe.Add(mBase, _c_F_rdbLoadRio[1]))
	*(*int32)(unsafe.Add(mBase, _c_F_rdbLoadRio[1])) = l0
	*(*int32)(unsafe.Add(mBase, uint32(v9)+12)) = v12
	v19 = *(*int32)(unsafe.Add(mBase, _c_F_rdbLoadRio[2]))
	*(*int32)(unsafe.Add(mBase, uint32(v9)+8)) = v19
	v23 = F_rdbLoadRioWithLoadingCtx(m, l0, l1, l2, v9+int32(8))
	mBase = m.M
	v26 = m.ExcPending
	if v26 != 0 {
		return int32(0)
	} else {
		*(*int32)(unsafe.Add(mBase, _c_F_rdbLoadRio[1])) = v14
		m.G0 = v9 + int32(16)
		return v23
	}
}
func F_rdbLoadRioWithLoadingCtxScopedRdb(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	v6 = int32(_a_F_rdbLoadRioWithLoadingCtxScopedRdb_0)
	v7 = *(*int32)(unsafe.Add(mBase, _c_F_rdbLoadRioWithLoadingCtxScopedRdb[0]))
	*(*int32)(unsafe.Add(mBase, _c_F_rdbLoadRioWithLoadingCtxScopedRdb[0])) = l0
	v10 = F_rdbLoadRioWithLoadingCtx(m, l0, l1, l2, l3)
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return int32(0)
	} else {
		*(*int32)(unsafe.Add(mBase, _c_F_rdbLoadRioWithLoadingCtxScopedRdb[0])) = v7
		return v10
	}
}
func F_rdbLoadType(m *base.Module, l0 int32) int32 {
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
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
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
	var v35 int64
	_ = v35
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	v8 = m.G0
	v10 = v8 - int32(16)
	m.G0 = v10
	v12 = int32(-1)
	v13 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
	if v13&int32(5) != 0 {
		v52 = v12
		goto L1
	} else {
		goto L2
	}
L1:
	;
	m.G0 = v10 + int32(16)
	return v52
L2:
	;
	v22 = v10 + int32(15)
	v23 = int32(1)
	goto L3
L3:
	;
	v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	if base.Ui32(v26) < base.Ui32(v23) {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	v49 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+15)))
	v52 = v49
	goto L1
L5:
	;
	v39 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v39 == int32(0) {
		goto L15
	} else {
		goto L16
	}
L6:
	;
	v28 = v26
	goto L8
L7:
	;
	v28 = v23
	goto L8
L8:
	;
	if v26 != 0 {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	v29 = v28
	goto L11
L10:
	;
	v29 = v23
	goto L11
L11:
	;
	v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v31 = m.T0[v30].(func(*base.Module, int32, int32, int32) int32)(m, l0, v22, v29)
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	return int32(0)
L13:
	;
	if v31 != 0 {
		goto L5
	} else {
		goto L14
	}
L14:
	;
	v35 = *(*int64)(unsafe.Add(mBase, uint32(l0)+32))
	*(*int64)(unsafe.Add(mBase, uint32(l0)+32)) = v35 | int64(1)
	v52 = v12
	goto L1
L15:
	;
	v44 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = v44 + v29
	v48 = v23 - v29
	if v48 != 0 {
		v22 = v22 + v29
		v23 = v48
		goto L3
	} else {
		goto L18
	}
L16:
	;
	m.T0[v39].(func(*base.Module, int32, int32, int32))(m, l0, v22, v29)
	mBase = m.M
	v43 = m.ExcPending
	if v43 != 0 {
		goto L12
	} else {
		goto L17
	}
L17:
	;
	goto L15
L18:
	;
	goto L4
}
func F_rdbPipeWriteHandlerConnRemoved(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
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
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v37 int32
	_ = v37
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v2 == int32(0) {
		return
	} else {
		v5 = int32(0)
		v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		v8 = *(*int32)(unsafe.Add(mBase, uint32(v7)+80))
		v9 = m.T0[v8].(func(*base.Module, int32, int32, int32) int32)(m, l0, v5, v5)
		mBase = m.M
		v10 = m.ExcPending
		if v10 != 0 {
			return
		} else {
			v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
			v12 = *(*int32)(unsafe.Add(mBase, uint32(v11)+104))
			*(*int64)(unsafe.Add(mBase, uint32(v12)+88)) = int64(0)
			v15 = int32(_a_F_rdbPipeWriteHandlerConnRemoved_0)
			v17 = *(*int32)(unsafe.Add(mBase, _c_F_rdbPipeWriteHandlerConnRemoved[0]))
			v19 = v17 + int32(-1)
			*(*int32)(unsafe.Add(mBase, _c_F_rdbPipeWriteHandlerConnRemoved[0])) = v19
			if v19 != 0 {
				return
			} else {
				v21 = int32(_a_F_rdbPipeWriteHandlerConnRemoved_0)
				v22 = *(*int32)(unsafe.Add(mBase, _c_F_rdbPipeWriteHandlerConnRemoved[1]))
				v24 = *(*int32)(unsafe.Add(mBase, _c_F_rdbPipeWriteHandlerConnRemoved[2]))
				v28 = F_aeCreateFileEvent(m, v22, v24, int32(1), int32(969), int32(0))
				mBase = m.M
				v29 = m.ExcPending
				if v29 != 0 {
					return
				} else {
					if v28 != int32(-1) {
						return
					} else {
						F__serverPanic_1(m, int32(_a_F_rdbPipeWriteHandlerConnRemoved_1), int32(1765), int32(_a_F_rdbPipeWriteHandlerConnRemoved_2), int32(0))
						mBase = m.M
						v37 = m.ExcPending
						if v37 != 0 {
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
func F_rdbPopulateSaveInfo(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v6 int64
	_ = v6
	var v11 int64
	_ = v11
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
	var v39 int64
	_ = v39
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	v2 = int32(0)
	v6 = *(*int64)(unsafe.Add(mBase, _c_F_rdbPopulateSaveInfo[0]))
	*(*int64)(unsafe.Add(mBase, uint32(l0+int32(56)))) = v6
	v11 = *(*int64)(unsafe.Add(mBase, _c_F_rdbPopulateSaveInfo[1]))
	*(*int64)(unsafe.Add(mBase, uint32(l0+int32(48)))) = v11
	v16 = *(*int64)(unsafe.Add(mBase, _c_F_rdbPopulateSaveInfo[2]))
	*(*int64)(unsafe.Add(mBase, uint32(l0+int32(40)))) = v16
	v21 = *(*int64)(unsafe.Add(mBase, _c_F_rdbPopulateSaveInfo[3]))
	*(*int64)(unsafe.Add(mBase, uint32(l0+int32(32)))) = v21
	v26 = *(*int64)(unsafe.Add(mBase, _c_F_rdbPopulateSaveInfo[4]))
	*(*int64)(unsafe.Add(mBase, uint32(l0+int32(24)))) = v26
	v31 = *(*int64)(unsafe.Add(mBase, _c_F_rdbPopulateSaveInfo[5]))
	*(*int64)(unsafe.Add(mBase, uint32(l0+int32(16)))) = v31
	v36 = *(*int64)(unsafe.Add(mBase, _c_F_rdbPopulateSaveInfo[6]))
	*(*int64)(unsafe.Add(mBase, uint32(l0+int32(8)))) = v36
	v39 = *(*int64)(unsafe.Add(mBase, _c_F_rdbPopulateSaveInfo[7]))
	*(*int64)(unsafe.Add(mBase, uint32(l0))) = v39
	v42 = *(*int32)(unsafe.Add(mBase, _c_F_rdbPopulateSaveInfo[8]))
	if v42 != 0 {
		v56 = *(*int32)(unsafe.Add(mBase, _c_F_rdbPopulateSaveInfo[9]))
		if v56 == int32(0) {
			v64 = *(*int32)(unsafe.Add(mBase, _c_F_rdbPopulateSaveInfo[10]))
			if v64 != 0 {
				v67 = *(*int32)(unsafe.Add(mBase, uint32(v64)+96))
				v68 = *(*int32)(unsafe.Add(mBase, uint32(v67)+28))
				*(*int32)(unsafe.Add(mBase, uint32(l0))) = v68
				return l0
			} else {
				return int32(0)
			}
		} else {
			v59 = *(*int32)(unsafe.Add(mBase, uint32(v56)+96))
			v60 = *(*int32)(unsafe.Add(mBase, uint32(v59)+28))
			*(*int32)(unsafe.Add(mBase, uint32(l0))) = v60
			return l0
		}
	} else {
		v44 = *(*int32)(unsafe.Add(mBase, _c_F_rdbPopulateSaveInfo[11]))
		if v44 == int32(0) {
			v56 = *(*int32)(unsafe.Add(mBase, _c_F_rdbPopulateSaveInfo[9]))
			if v56 == int32(0) {
				v64 = *(*int32)(unsafe.Add(mBase, _c_F_rdbPopulateSaveInfo[10]))
				if v64 != 0 {
					v67 = *(*int32)(unsafe.Add(mBase, uint32(v64)+96))
					v68 = *(*int32)(unsafe.Add(mBase, uint32(v67)+28))
					*(*int32)(unsafe.Add(mBase, uint32(l0))) = v68
					return l0
				} else {
					return int32(0)
				}
			} else {
				v59 = *(*int32)(unsafe.Add(mBase, uint32(v56)+96))
				v60 = *(*int32)(unsafe.Add(mBase, uint32(v59)+28))
				*(*int32)(unsafe.Add(mBase, uint32(l0))) = v60
				return l0
			}
		} else {
			v49 = *(*int32)(unsafe.Add(mBase, _c_F_rdbPopulateSaveInfo[12]))
			if v49 == int32(-1) {
				v52 = int32(0)
			} else {
				v52 = v49
			}
			*(*int32)(unsafe.Add(mBase, uint32(l0))) = v52
			return l0
		}
	}
}
func F_rdbRemoveTempFile(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v9 int64
	_ = v9
	var v18 int32
	_ = v18
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int64
	_ = v29
	var v31 int32
	_ = v31
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v74 int32
	_ = v74
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
	var v114 int32
	_ = v114
	var v120 int32
	_ = v120
	var v124 int32
	_ = v124
	var v128 int32
	_ = v128
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v136 int32
	_ = v136
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v141 int32
	_ = v141
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v157 int32
	_ = v157
	var v161 int32
	_ = v161
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v187 int32
	_ = v187
	var v189 int32
	_ = v189
	var v191 int32
	_ = v191
	var v197 int32
	_ = v197
	var v201 int32
	_ = v201
	var v205 int32
	_ = v205
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v213 int32
	_ = v213
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v218 int32
	_ = v218
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v234 int32
	_ = v234
	var v238 int32
	_ = v238
	var v253 int32
	_ = v253
	var v256 int32
	_ = v256
	var v259 int32
	_ = v259
	var v262 int32
	_ = v262
	var v263 int32
	_ = v263
	v4 = m.G0
	v6 = v4 - int32(288)
	m.G0 = v6
	v9 = base.I64_extend_i32_s(l0)
	if v9 <= int64(-1) {
		goto L5
	} else {
		goto L6
	}
L1:
	;
	goto L13
L2:
	;
	goto L1
L4:
	;
	v31 = F_ull2string(m, v27, v28, v29)
	mBase = m.M
	if v31 == int32(0) {
		goto L2
	} else {
		goto L8
	}
L5:
	;
	goto L7
L6:
	;
	v27 = v6
	v28 = int32(32)
	v29 = v9
	goto L4
L7:
	;
	v18 = int32(45)
	*(*uint8)(unsafe.Add(mBase, uint32(v6))) = uint8(v18)
	v27 = v6 + int32(1)
	v28 = int32(31)
	v29 = int64(0) - v9
	goto L4
L8:
	;
	goto L1
L10:
	;
	v98 = v6 + int32(32)
	v99 = int32(256)
	goto L24
L11:
	;
	goto L10
L12:
	;
	v84 = v63
	goto L19
L13:
	;
	v59 = v6 + int32(32)
	v61 = int32(256)
	v63 = int32(_a_F_rdbRemoveTempFile_0)
	goto L15
L14:
	;
	v74 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v59))) = uint8(v74)
	goto L12
L15:
	;
	v65 = v61 + int32(-1)
	if v65 == int32(0) {
		goto L14
	} else {
		goto L17
	}
L17:
	;
	v68 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v63))))
	*(*uint8)(unsafe.Add(mBase, uint32(v59))) = uint8(v68)
	v70 = int32(1)
	if v68 != 0 {
		v59 = v59 + v70
		v61 = v65
		v63 = v63 + v70
		goto L15
	} else {
		goto L18
	}
L18:
	;
	goto L11
L19:
	;
	v86 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v84))))
	if v86 != 0 {
		v84 = v84 + int32(1)
		goto L19
	} else {
		goto L21
	}
L20:
	;
	goto L11
L21:
	;
	goto L20
L22:
	;
	v174 = v6 + int32(32)
	v175 = int32(_a_F_rdbRemoveTempFile_1)
	v176 = int32(256)
	goto L42
L23:
	;
	v128 = v124 - v98
	if v99 == v128 {
		goto L31
	} else {
		goto L32
	}
L24:
	;
	v110 = v98
	v112 = v99
	goto L25
L25:
	;
	v114 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v110))))
	if v114 == int32(0) {
		v124 = v110
		goto L23
	} else {
		goto L27
	}
L26:
	;
	v124 = v6 + int32(288)
	goto L23
L27:
	;
	v120 = v112 + int32(-1)
	if v120 != 0 {
		v110 = v110 + int32(1)
		v112 = v120
		goto L25
	} else {
		goto L28
	}
L28:
	;
	goto L26
L29:
	;
	v161 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v157))) = uint8(v161)
	goto L22
L30:
	;
	v136 = v130
	v138 = v99 + (v128 ^ int32(-1))
	v139 = v124
	v141 = v6
	goto L34
L31:
	;
	v131 = F_strlen(m, v6)
	mBase = m.M
	goto L22
L32:
	;
	v130 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6))))
	if v130 != 0 {
		goto L30
	} else {
		goto L33
	}
L33:
	;
	v157 = v124
	goto L29
L34:
	;
	if v138 != 0 {
		goto L37
	} else {
		goto L38
	}
L35:
	;
	v157 = v150
	goto L29
L36:
	;
	v151 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v141)+1)))
	if v151 != 0 {
		v136 = v151
		v138 = v149
		v139 = v150
		v141 = v141 + int32(1)
		goto L34
	} else {
		goto L39
	}
L37:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v139))) = uint8(v136)
	v149 = v138 + int32(-1)
	v150 = v139 + int32(1)
	goto L36
L38:
	;
	v149 = int32(0)
	v150 = v139
	goto L36
L39:
	;
	goto L35
L40:
	;
	if l1 == int32(0) {
		goto L59
	} else {
		goto L60
	}
L41:
	;
	v205 = v201 - v174
	if v176 == v205 {
		goto L49
	} else {
		goto L50
	}
L42:
	;
	v187 = v174
	v189 = v176
	goto L43
L43:
	;
	v191 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v187))))
	if v191 == int32(0) {
		v201 = v187
		goto L41
	} else {
		goto L45
	}
L44:
	;
	v201 = v6 + int32(288)
	goto L41
L45:
	;
	v197 = v189 + int32(-1)
	if v197 != 0 {
		v187 = v187 + int32(1)
		v189 = v197
		goto L43
	} else {
		goto L46
	}
L46:
	;
	goto L44
L47:
	;
	v238 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v234))) = uint8(v238)
	goto L40
L48:
	;
	v213 = v207
	v215 = v176 + (v205 ^ int32(-1))
	v216 = v201
	v218 = v175
	goto L52
L49:
	;
	v208 = F_strlen(m, v175)
	mBase = m.M
	goto L40
L50:
	;
	v207 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_rdbRemoveTempFile[0])))
	if v207 != 0 {
		goto L48
	} else {
		goto L51
	}
L51:
	;
	v234 = v201
	goto L47
L52:
	;
	if v215 != 0 {
		goto L55
	} else {
		goto L56
	}
L53:
	;
	v234 = v227
	goto L47
L54:
	;
	v228 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v218)+1)))
	if v228 != 0 {
		v213 = v228
		v215 = v226
		v216 = v227
		v218 = v218 + int32(1)
		goto L52
	} else {
		goto L57
	}
L55:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v216))) = uint8(v213)
	v226 = v215 + int32(-1)
	v227 = v216 + int32(1)
	goto L54
L56:
	;
	v226 = int32(0)
	v227 = v216
	goto L54
L57:
	;
	goto L53
L58:
	;
	m.G0 = v6 + int32(288)
	return
L59:
	;
	v262 = F_bg_unlink(m, v6+int32(32))
	mBase = m.M
	v263 = m.ExcPending
	if v263 != 0 {
		goto L61
	} else {
		goto L62
	}
L60:
	;
	v253 = v6 + int32(32)
	v256 = F_open(m, v253, int32(2048), int32(0))
	mBase = m.M
	v259 = F_unlink(m, v253)
	mBase = m.M
	goto L58
L61:
	;
	return
L62:
	;
	goto L58
}
func F_rdbSave(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
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
	var v25 int32
	_ = v25
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
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
	var v49 int32
	_ = v49
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v66 int32
	_ = v66
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v89 int32
	_ = v89
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v118 int32
	_ = v118
	var v123 int32
	_ = v123
	var v125 int32
	_ = v125
	var v132 int32
	_ = v132
	var v135 int32
	_ = v135
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v147 int32
	_ = v147
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v158 int32
	_ = v158
	var v160 int64
	_ = v160
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v175 int32
	_ = v175
	v9 = m.G0
	v11 = v9 - int32(4400)
	m.G0 = v11
	v16 = F___syscall_getpid(m)
	mBase = m.M
	v18 = *(*int32)(unsafe.Add(mBase, _c_F_rdbSave[0]))
	v19 = base.B2i32(v16 == v18)
	if v16 == v18 {
		v20 = int32(5)
	} else {
		v20 = int32(1)
	}
	v21 = int32(1)
	if l3&v21 != 0 {
		v25 = v20
	} else {
		v25 = v19 << (uint(v21) % 32)
	}
	F_moduleFireServerEvent(m, int64(1), v25, int32(0))
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		return int32(0)
	} else {
		v31 = F___syscall_getpid(m)
		mBase = m.M
		*(*int32)(unsafe.Add(mBase, uint32(v11)+32)) = v31
		v39 = F_snprintf(m, v11+int32(4144), int32(256), int32(_a_F_rdbSave_0), v11+int32(32))
		mBase = m.M
		v40 = m.ExcPending
		if v40 != 0 {
			return int32(0)
		} else {
			v42 = int32(4)
			v45 = F_rdbSaveInternal(m, l0, v11+int32(4144), l2, l3)
			mBase = m.M
			v46 = m.ExcPending
			if v46 != 0 {
				return int32(0)
			} else {
				if v45 != 0 {
					v169 = int32(-1)
					v170 = v42
					F_moduleFireServerEvent(m, int64(1), v170, int32(0))
					mBase = m.M
					v175 = m.ExcPending
					if v175 != 0 {
						return int32(0)
					} else {
						m.G0 = v11 + int32(4400)
						return v169
					}
				} else {
					v49 = F_rename(m, v11+int32(4144), l1)
					mBase = m.M
					if v49 != int32(-1) {
						v82 = m.G0
						v84 = v82 - int32(4112)
						m.G0 = v84
						v86 = F_strlen(m, l1)
						mBase = m.M
						if base.Ui32(v86) < base.Ui32(int32(4097)) {
							v95 = F___memcpy(m, v84, l1, v86+int32(1))
							mBase = m.M
							v96 = F_dirname(m, v95)
							mBase = m.M
							v97 = int32(0)
							v99 = F_open(m, v96, v97, v97)
							mBase = m.M
							if v99 != int32(-1) {
								v109 = F_fsync(m, v99)
								mBase = m.M
								if v109 != int32(-1) {
									v123 = F_close(m, v99)
									mBase = m.M
									v125 = int32(0)
								} else {
									v112 = F___errno_location(m)
									mBase = m.M
									v113 = *(*int32)(unsafe.Add(mBase, uint32(v112)))
									if v113 == int32(8) {
										v123 = F_close(m, v99)
										mBase = m.M
										v125 = int32(0)
									} else {
										if v113 == int32(28) {
											v123 = F_close(m, v99)
											mBase = m.M
											v125 = int32(0)
										} else {
											v118 = F_close(m, v99)
											mBase = m.M
											*(*int32)(unsafe.Add(mBase, uint32(v112))) = v113
											v125 = int32(-1)
										}
									}
								}
							} else {
								v104 = F___errno_location(m)
								mBase = m.M
								v105 = *(*int32)(unsafe.Add(mBase, uint32(v104)))
								if v105 != int32(31) {
									v108 = int32(-1)
								} else {
									v108 = int32(0)
								}
								v125 = v108
							}
						} else {
							v89 = F___errno_location(m)
							mBase = m.M
							*(*int32)(unsafe.Add(mBase, uint32(v89))) = int32(37)
							v125 = int32(-1)
						}
						m.G0 = v84 + int32(4112)
						v132 = *(*int32)(unsafe.Add(mBase, _c_F_rdbSave[1]))
						if v125 == int32(0) {
							if int32(2) < v132 {
								v155 = int32(_a_F_rdbSave_1)
								*(*int64)(unsafe.Add(mBase, _c_F_rdbSave[2])) = int64(0)
								v158 = int32(0)
								v160 = F___time(m, v158)
								mBase = m.M
								*(*int32)(unsafe.Add(mBase, _c_F_rdbSave[3])) = v158
								*(*int64)(unsafe.Add(mBase, _c_F_rdbSave[4])) = v160
								v169 = v158
								v170 = int32(3)
								F_moduleFireServerEvent(m, int64(1), v170, int32(0))
								mBase = m.M
								v175 = m.ExcPending
								if v175 != 0 {
									return int32(0)
								} else {
									m.G0 = v11 + int32(4400)
									return v169
								}
							} else {
								F__serverLog(m, int32(2), int32(_a_F_rdbSave_2), int32(0))
								mBase = m.M
								v154 = m.ExcPending
								if v154 != 0 {
									return int32(0)
								} else {
									v155 = int32(_a_F_rdbSave_1)
									*(*int64)(unsafe.Add(mBase, _c_F_rdbSave[2])) = int64(0)
									v158 = int32(0)
									v160 = F___time(m, v158)
									mBase = m.M
									*(*int32)(unsafe.Add(mBase, _c_F_rdbSave[3])) = v158
									*(*int64)(unsafe.Add(mBase, _c_F_rdbSave[4])) = v160
									v169 = v158
									v170 = int32(3)
									F_moduleFireServerEvent(m, int64(1), v170, int32(0))
									mBase = m.M
									v175 = m.ExcPending
									if v175 != 0 {
										return int32(0)
									} else {
										m.G0 = v11 + int32(4400)
										return v169
									}
								}
							}
						} else {
							v135 = int32(-1)
							if int32(3) < v132 {
								v169 = v135
								v170 = v42
								F_moduleFireServerEvent(m, int64(1), v170, int32(0))
								mBase = m.M
								v175 = m.ExcPending
								if v175 != 0 {
									return int32(0)
								} else {
									m.G0 = v11 + int32(4400)
									return v169
								}
							} else {
								v139 = *(*int32)(unsafe.Add(mBase, _c_F_rdbSave[5]))
								v140 = F___strerror_l(m, v139, v139)
								mBase = m.M
								*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = v140
								F__serverLog(m, int32(3), int32(_a_F_rdbSave_3), v11+int32(16))
								mBase = m.M
								v147 = m.ExcPending
								if v147 != 0 {
									return int32(0)
								} else {
									v169 = v135
									v170 = v42
									F_moduleFireServerEvent(m, int64(1), v170, int32(0))
									mBase = m.M
									v175 = m.ExcPending
									if v175 != 0 {
										return int32(0)
									} else {
										m.G0 = v11 + int32(4400)
										return v169
									}
								}
							}
						}
					} else {
						v53 = *(*int32)(unsafe.Add(mBase, _c_F_rdbSave[5]))
						v54 = F___strerror_l(m, v53, v53)
						mBase = m.M
						v58 = F_getcwd(m, v11+int32(48), int32(4096))
						mBase = m.M
						v60 = *(*int32)(unsafe.Add(mBase, _c_F_rdbSave[1]))
						if int32(3) < v60 {
							v77 = F_unlink(m, v11+int32(4144))
							mBase = m.M
							v169 = int32(-1)
							v170 = v42
							F_moduleFireServerEvent(m, int64(1), v170, int32(0))
							mBase = m.M
							v175 = m.ExcPending
							if v175 != 0 {
								return int32(0)
							} else {
								m.G0 = v11 + int32(4400)
								return v169
							}
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v11)+12)) = v54
							*(*int32)(unsafe.Add(mBase, uint32(v11)+4)) = l1
							if v58 != 0 {
								v66 = v58
							} else {
								v66 = int32(_a_F_rdbSave_4)
							}
							*(*int32)(unsafe.Add(mBase, uint32(v11)+8)) = v66
							*(*int32)(unsafe.Add(mBase, uint32(v11))) = v11 + int32(4144)
							F__serverLog(m, int32(3), int32(_a_F_rdbSave_5), v11)
							mBase = m.M
							v74 = m.ExcPending
							if v74 != 0 {
								return int32(0)
							} else {
								v77 = F_unlink(m, v11+int32(4144))
								mBase = m.M
								v169 = int32(-1)
								v170 = v42
								F_moduleFireServerEvent(m, int64(1), v170, int32(0))
								mBase = m.M
								v175 = m.ExcPending
								if v175 != 0 {
									return int32(0)
								} else {
									m.G0 = v11 + int32(4400)
									return v169
								}
							}
						}
					}
				}
			}
		}
	}
}
func F_rdbSaveBackground(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v19 int64
	_ = v19
	var v23 int64
	_ = v23
	var v29 int64
	_ = v29
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
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
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v81 int32
	_ = v81
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
	var v103 int32
	_ = v103
	var v105 int32
	_ = v105
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v115 int64
	_ = v115
	var v116 int32
	_ = v116
	var v121 int32
	_ = v121
	v8 = m.G0
	v10 = v8 - int32(32)
	m.G0 = v10
	v12 = int32(-1)
	v14 = *(*int32)(unsafe.Add(mBase, _c_F_rdbSaveBackground[0]))
	if v14 != v12 {
		v121 = v12
		m.G0 = v10 + int32(32)
		return v121
	} else {
		v17 = int32(_a_F_rdbSaveBackground_0)
		v19 = *(*int64)(unsafe.Add(mBase, _c_F_rdbSaveBackground[1]))
		*(*int64)(unsafe.Add(mBase, _c_F_rdbSaveBackground[2])) = v19
		v23 = *(*int64)(unsafe.Add(mBase, _c_F_rdbSaveBackground[3]))
		*(*int64)(unsafe.Add(mBase, _c_F_rdbSaveBackground[3])) = v23 + int64(1)
		v29 = F___time(m, int32(0))
		mBase = m.M
		*(*int64)(unsafe.Add(mBase, _c_F_rdbSaveBackground[4])) = v29
		v32 = F_serverFork(m, int32(1))
		mBase = m.M
		v35 = m.ExcPending
		if v35 != 0 {
			return int32(0)
		} else {
			switch v32 + int32(1) {
			case 0:
				v86 = int32(-1)
				v87 = int32(_a_F_rdbSaveBackground_0)
				*(*int32)(unsafe.Add(mBase, _c_F_rdbSaveBackground[5])) = v86
				v91 = *(*int32)(unsafe.Add(mBase, _c_F_rdbSaveBackground[6]))
				if int32(3) < v91 {
					v121 = v86
					m.G0 = v10 + int32(32)
					return v121
				} else {
					v95 = *(*int32)(unsafe.Add(mBase, _c_F_rdbSaveBackground[7]))
					v96 = F___strerror_l(m, v95, v95)
					mBase = m.M
					*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = v96
					F__serverLog(m, int32(3), int32(_a_F_rdbSaveBackground_1), v10+int32(16))
					mBase = m.M
					v103 = m.ExcPending
					if v103 != 0 {
						return int32(0)
					} else {
						v121 = v86
						m.G0 = v10 + int32(32)
						return v121
					}
				}
			case 1:
				v41 = *(*int32)(unsafe.Add(mBase, _c_F_rdbSaveBackground[8]))
				v42 = *(*int32)(unsafe.Add(mBase, uint32(v41)))
				v43 = int32(_a_F_rdbSaveBackground_2)
				v46 = int32(*(*int8)(unsafe.Add(mBase, _c_F_rdbSaveBackground[9])))
				if v46 != 0 {
					v48 = F_strchr(m, v42, v46)
					mBase = m.M
					if v48 == int32(0) {
					} else {
						v51 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_rdbSaveBackground[10])))
						if v51 != 0 {
							v52 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v48)+1)))
							if v52 == int32(0) {
							} else {
								v55 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_rdbSaveBackground[11])))
								if v55 != 0 {
									v57 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v48)+2)))
									if v57 == int32(0) {
									} else {
										v60 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_rdbSaveBackground[12])))
										if v60 != 0 {
											v62 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v48)+3)))
											if v62 == int32(0) {
											} else {
												v65 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_rdbSaveBackground[13])))
												if v65 != 0 {
													v67 = F_twoway_strstr(m, v48, v43)
													mBase = m.M
												} else {
													v66 = F_fourbyte_strstr(m, v48, v43)
													mBase = m.M
												}
											}
										} else {
											v61 = F_threebyte_strstr(m, v48, v43)
											mBase = m.M
										}
									}
								} else {
									v56 = F_twobyte_strstr(m, v48, v43)
									mBase = m.M
								}
							}
						} else {
						}
					}
				} else {
				}
				v76 = F_rdbSave(m, l0, l1, l2, l3)
				mBase = m.M
				v77 = m.ExcPending
				if v77 != 0 {
					return int32(0)
				} else {
					if v76 != 0 {
						F__exit(m, base.B2i32(v76 != int32(0)))
						mBase = m.M
						base.Wasm_trap_unreachable()
						for {
						}
					} else {
						F_sendChildCowInfo(m, int32(2), int32(_a_F_rdbSaveBackground_3))
						mBase = m.M
						v81 = m.ExcPending
						if v81 != 0 {
							return int32(0)
						} else {
							F__exit(m, base.B2i32(v76 != int32(0)))
							mBase = m.M
							base.Wasm_trap_unreachable()
							for {
							}
						}
					}
				}
			default:
				v105 = *(*int32)(unsafe.Add(mBase, _c_F_rdbSaveBackground[6]))
				if int32(2) < v105 {
					v113 = int32(0)
					v115 = F___time(m, v113)
					mBase = m.M
					v116 = int32(_a_F_rdbSaveBackground_0)
					*(*int32)(unsafe.Add(mBase, _c_F_rdbSaveBackground[14])) = int32(1)
					*(*int64)(unsafe.Add(mBase, _c_F_rdbSaveBackground[15])) = v115
					v121 = v113
					m.G0 = v10 + int32(32)
					return v121
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v10))) = v32
					F__serverLog(m, int32(2), int32(_a_F_rdbSaveBackground_4), v10)
					mBase = m.M
					v112 = m.ExcPending
					if v112 != 0 {
						return int32(0)
					} else {
						v113 = int32(0)
						v115 = F___time(m, v113)
						mBase = m.M
						v116 = int32(_a_F_rdbSaveBackground_0)
						*(*int32)(unsafe.Add(mBase, _c_F_rdbSaveBackground[14])) = int32(1)
						*(*int64)(unsafe.Add(mBase, _c_F_rdbSaveBackground[15])) = v115
						v121 = v113
						m.G0 = v10 + int32(32)
						return v121
					}
				}
			}
		}
	}
}
func F_rdbSaveInfoAuxFields(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v47 int32
	_ = v47
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v75 int32
	_ = v75
	var v76 int64
	_ = v76
	var v85 int32
	_ = v85
	var v89 int32
	_ = v89
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v96 int64
	_ = v96
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v126 int32
	_ = v126
	var v135 int32
	_ = v135
	var v139 int32
	_ = v139
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v146 int32
	_ = v146
	var v153 int32
	_ = v153
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v160 int32
	_ = v160
	var v163 int32
	_ = v163
	var v166 int32
	_ = v166
	var v169 int32
	_ = v169
	var v172 int32
	_ = v172
	var v175 int32
	_ = v175
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v183 int32
	_ = v183
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v197 int32
	_ = v197
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v210 int32
	_ = v210
	var v212 int32
	_ = v212
	var v219 int64
	_ = v219
	var v228 int32
	_ = v228
	var v232 int32
	_ = v232
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v239 int64
	_ = v239
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v260 int32
	_ = v260
	var v261 int32
	_ = v261
	var v262 int32
	_ = v262
	var v271 int64
	_ = v271
	var v280 int32
	_ = v280
	var v284 int32
	_ = v284
	var v289 int32
	_ = v289
	var v290 int32
	_ = v290
	var v291 int64
	_ = v291
	var v292 int32
	_ = v292
	var v293 int32
	_ = v293
	var v312 int32
	_ = v312
	var v313 int32
	_ = v313
	var v314 int32
	_ = v314
	var v319 int32
	_ = v319
	var v348 int32
	_ = v348
	var v354 int32
	_ = v354
	var v357 int32
	_ = v357
	var v363 int32
	_ = v363
	var v367 int32
	_ = v367
	var v378 int32
	_ = v378
	var v379 int32
	_ = v379
	var v387 int64
	_ = v387
	var v396 int32
	_ = v396
	var v400 int32
	_ = v400
	var v405 int32
	_ = v405
	var v406 int32
	_ = v406
	var v407 int64
	_ = v407
	var v408 int32
	_ = v408
	var v409 int32
	_ = v409
	var v428 int32
	_ = v428
	var v429 int32
	_ = v429
	var v430 int32
	_ = v430
	var v441 int64
	_ = v441
	var v450 int32
	_ = v450
	var v454 int32
	_ = v454
	var v459 int32
	_ = v459
	var v460 int32
	_ = v460
	var v461 int64
	_ = v461
	var v462 int32
	_ = v462
	var v463 int32
	_ = v463
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
	var v510 int32
	_ = v510
	var v511 int32
	_ = v511
	var v514 int32
	_ = v514
	var v517 int32
	_ = v517
	var v518 int32
	_ = v518
	var v522 int32
	_ = v522
	var v526 int32
	_ = v526
	var v527 int32
	_ = v527
	var v528 int32
	_ = v528
	var v531 int32
	_ = v531
	var v533 int32
	_ = v533
	var v535 int64
	_ = v535
	var v536 int64
	_ = v536
	var v537 int64
	_ = v537
	var v538 int64
	_ = v538
	var v539 int64
	_ = v539
	var v540 int64
	_ = v540
	var v541 int64
	_ = v541
	var v543 int64
	_ = v543
	var v545 int64
	_ = v545
	var v547 int64
	_ = v547
	var v549 int64
	_ = v549
	var v551 int64
	_ = v551
	var v553 int32
	_ = v553
	var v554 int32
	_ = v554
	var v555 int32
	_ = v555
	var v558 int32
	_ = v558
	var v561 int32
	_ = v561
	var v562 int32
	_ = v562
	var v563 int32
	_ = v563
	var v565 int32
	_ = v565
	var v568 int32
	_ = v568
	var v569 int32
	_ = v569
	var v573 int32
	_ = v573
	var v577 int32
	_ = v577
	var v579 int32
	_ = v579
	var v586 int32
	_ = v586
	var v587 int32
	_ = v587
	var v588 int32
	_ = v588
	var v593 int32
	_ = v593
	var v598 int32
	_ = v598
	var v602 int32
	_ = v602
	var v606 int32
	_ = v606
	var v612 int32
	_ = v612
	var v613 int32
	_ = v613
	var v614 int32
	_ = v614
	var v615 int32
	_ = v615
	var v618 int32
	_ = v618
	var v626 int32
	_ = v626
	var v629 int32
	_ = v629
	var v633 int32
	_ = v633
	var v638 int32
	_ = v638
	var v640 int32
	_ = v640
	var v644 int32
	_ = v644
	var v650 int32
	_ = v650
	var v653 int32
	_ = v653
	var v659 int32
	_ = v659
	var v663 int32
	_ = v663
	var v665 int32
	_ = v665
	var v673 int32
	_ = v673
	var v681 int32
	_ = v681
	var v684 int32
	_ = v684
	var v688 int32
	_ = v688
	var v693 int32
	_ = v693
	var v695 int32
	_ = v695
	var v699 int32
	_ = v699
	var v705 int32
	_ = v705
	var v708 int32
	_ = v708
	var v714 int32
	_ = v714
	var v718 int32
	_ = v718
	var v720 int32
	_ = v720
	var v728 int32
	_ = v728
	var v729 int32
	_ = v729
	var v730 int32
	_ = v730
	var v732 int32
	_ = v732
	var v733 int32
	_ = v733
	var v746 int32
	_ = v746
	v6 = m.G0
	v8 = v6 - int32(32)
	m.G0 = v8
	v15 = F_rdbSaveAuxField(m, l0, int32(_a_F_rdbSaveInfoAuxFields_0), int32(10), int32(_a_F_rdbSaveInfoAuxFields_1), int32(5))
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v8 + int32(32)
	return v746
L2:
	;
	return int32(0)
L3:
	;
	if v15 == int32(-1) {
		v746 = int32(-1)
		goto L1
	} else {
		goto L4
	}
L4:
	;
	goto L10
L5:
	;
	v67 = F_rdbSaveAuxField(m, l0, int32(_a_F_rdbSaveInfoAuxFields_2), int32(10), v8, v66)
	mBase = m.M
	v68 = m.ExcPending
	if v68 != 0 {
		goto L2
	} else {
		goto L14
	}
L6:
	;
	v66 = int32(0)
	goto L5
L8:
	;
	v47 = F_ull2string(m, v8, int32(21), int64(32))
	mBase = m.M
	if v47 == int32(0) {
		goto L6
	} else {
		goto L12
	}
L10:
	;
	goto L8
L12:
	;
	v66 = v47 + int32(0)
	goto L5
L14:
	;
	if v67 == int32(-1) {
		v746 = int32(-1)
		goto L1
	} else {
		goto L15
	}
L15:
	;
	v75 = int32(0)
	v76 = F___time(m, v75)
	mBase = m.M
	if v76 <= int64(-1) {
		goto L20
	} else {
		goto L21
	}
L16:
	;
	v118 = F_rdbSaveAuxField(m, l0, int32(_a_F_rdbSaveInfoAuxFields_3), int32(5), v8, v117)
	mBase = m.M
	v119 = m.ExcPending
	if v119 != 0 {
		goto L2
	} else {
		goto L25
	}
L17:
	;
	v117 = v75
	goto L16
L19:
	;
	v98 = F_ull2string(m, v94, v95, v96)
	mBase = m.M
	if v98 == int32(0) {
		goto L17
	} else {
		goto L23
	}
L20:
	;
	goto L22
L21:
	;
	v94 = v8
	v95 = int32(21)
	v96 = v76
	v97 = int32(0)
	goto L19
L22:
	;
	v85 = int32(45)
	*(*uint8)(unsafe.Add(mBase, uint32(v8))) = uint8(v85)
	v89 = int32(1)
	v94 = v8 + v89
	v95 = int32(20)
	v96 = int64(0) - v76
	v97 = v89
	goto L19
L23:
	;
	v117 = v98 + v97
	goto L16
L25:
	;
	if v118 == int32(-1) {
		v746 = int32(-1)
		goto L1
	} else {
		goto L26
	}
L26:
	;
	v126 = int32(0)
	v135 = *(*int32)(unsafe.Add(mBase, _c_F_rdbSaveInfoAuxFields[0]))
	if v135 < int32(261) {
		goto L30
	} else {
		goto L31
	}
L27:
	;
	v219 = base.I64_extend_i32_u(v212)
	if v219 <= int64(-1) {
		goto L47
	} else {
		goto L48
	}
L28:
	;
	goto L27
L29:
	;
	v146 = v144 & int32(3)
	if base.Ui32(int32(4)) <= base.Ui32(v144) {
		goto L34
	} else {
		goto L35
	}
L30:
	;
	if v135 < int32(1) {
		v212 = v126
		goto L28
	} else {
		goto L32
	}
L31:
	;
	v139 = *(*int32)(unsafe.Add(mBase, _c_F_rdbSaveInfoAuxFields[1]))
	v143 = v139
	v144 = int32(260)
	goto L29
L32:
	;
	v143 = v126
	v144 = v135
	goto L29
L33:
	;
	if v146 == int32(0) {
		v212 = v185
		goto L28
	} else {
		goto L39
	}
L34:
	;
	v153 = int32(0)
	v155 = v143
	v156 = v153
	v160 = v153
	goto L36
L35:
	;
	v185 = v143
	v186 = int32(0)
	goto L33
L36:
	;
	v163 = v156 << (uint(int32(2)) % 32)
	v166 = *(*int32)(unsafe.Add(mBase, uint32(v163)+uint32(_c_F_rdbSaveInfoAuxFields[2])))
	v169 = *(*int32)(unsafe.Add(mBase, uint32(v163)+uint32(_c_F_rdbSaveInfoAuxFields[3])))
	v172 = *(*int32)(unsafe.Add(mBase, uint32(v163)+uint32(_c_F_rdbSaveInfoAuxFields[4])))
	v175 = *(*int32)(unsafe.Add(mBase, uint32(v163)+uint32(_c_F_rdbSaveInfoAuxFields[5])))
	v179 = v166 + (v169 + (v172 + (v175 + v155)))
	v180 = int32(4)
	v181 = v156 + v180
	v183 = v160 + v180
	if v183 != v144&int32(2147483644) {
		v155 = v179
		v156 = v181
		v160 = v183
		goto L36
	} else {
		goto L38
	}
L37:
	;
	v185 = v179
	v186 = v181
	goto L33
L38:
	;
	goto L37
L39:
	;
	v194 = v185
	v195 = v186
	v197 = int32(0)
	goto L40
L40:
	;
	v205 = *(*int32)(unsafe.Add(mBase, uint32(v195<<(uint(int32(2))%32))+uint32(_c_F_rdbSaveInfoAuxFields[5])))
	v206 = v205 + v194
	v207 = int32(1)
	v210 = v197 + v207
	if v210 != v146 {
		v194 = v206
		v195 = v195 + v207
		v197 = v210
		goto L40
	} else {
		goto L42
	}
L41:
	;
	v212 = v206
	goto L28
L42:
	;
	goto L41
L43:
	;
	v261 = F_rdbSaveAuxField(m, l0, int32(_a_F_rdbSaveInfoAuxFields_4), int32(8), v8, v260)
	mBase = m.M
	v262 = m.ExcPending
	if v262 != 0 {
		goto L2
	} else {
		goto L52
	}
L44:
	;
	v260 = int32(0)
	goto L43
L46:
	;
	v241 = F_ull2string(m, v237, v238, v239)
	mBase = m.M
	if v241 == int32(0) {
		goto L44
	} else {
		goto L50
	}
L47:
	;
	goto L49
L48:
	;
	v237 = v8
	v238 = int32(21)
	v239 = v219
	v240 = int32(0)
	goto L46
L49:
	;
	v228 = int32(45)
	*(*uint8)(unsafe.Add(mBase, uint32(v8))) = uint8(v228)
	v232 = int32(1)
	v237 = v8 + v232
	v238 = int32(20)
	v239 = int64(0) - v219
	v240 = v232
	goto L46
L50:
	;
	v260 = v241 + v240
	goto L43
L52:
	;
	if v261 == int32(-1) {
		v746 = int32(-1)
		goto L1
	} else {
		goto L53
	}
L53:
	;
	if l2 == int32(0) {
		goto L54
	} else {
		goto L55
	}
L54:
	;
	v441 = base.I64_extend_i32_u(l1 & int32(1))
	if v441 <= int64(-1) {
		goto L100
	} else {
		goto L101
	}
L55:
	;
	v271 = int64(*(*int32)(unsafe.Add(mBase, uint32(l2))))
	if v271 <= int64(-1) {
		goto L60
	} else {
		goto L61
	}
L56:
	;
	v313 = F_rdbSaveAuxField(m, l0, int32(_a_F_rdbSaveInfoAuxFields_5), int32(14), v8, v312)
	mBase = m.M
	v314 = m.ExcPending
	if v314 != 0 {
		goto L2
	} else {
		goto L65
	}
L57:
	;
	v312 = int32(0)
	goto L56
L59:
	;
	v293 = F_ull2string(m, v289, v290, v291)
	mBase = m.M
	if v293 == int32(0) {
		goto L57
	} else {
		goto L63
	}
L60:
	;
	goto L62
L61:
	;
	v289 = v8
	v290 = int32(21)
	v291 = v271
	v292 = int32(0)
	goto L59
L62:
	;
	v280 = int32(45)
	*(*uint8)(unsafe.Add(mBase, uint32(v8))) = uint8(v280)
	v284 = int32(1)
	v289 = v8 + v284
	v290 = int32(20)
	v291 = int64(0) - v271
	v292 = v284
	goto L59
L63:
	;
	v312 = v293 + v292
	goto L56
L65:
	;
	if v313 == int32(-1) {
		v746 = int32(-1)
		goto L1
	} else {
		goto L66
	}
L66:
	;
	v319 = int32(_a_F_rdbSaveInfoAuxFields_6)
	goto L69
L67:
	;
	v378 = F_rdbSaveAuxField(m, l0, int32(_a_F_rdbSaveInfoAuxFields_7), int32(7), v319, v363-v319)
	mBase = m.M
	v379 = m.ExcPending
	if v379 != 0 {
		goto L2
	} else {
		goto L83
	}
L68:
	;
	goto L67
L69:
	;
	v348 = v319
	goto L77
L77:
	;
	v354 = *(*int32)(unsafe.Add(mBase, uint32(v348)))
	v357 = int32(-2139062144)
	if (int32(16843008)-v354|v354)&v357 == v357 {
		v348 = v348 + int32(4)
		goto L77
	} else {
		goto L79
	}
L78:
	;
	v363 = v348
	goto L80
L79:
	;
	goto L78
L80:
	;
	v367 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v363))))
	if v367 != 0 {
		v363 = v363 + int32(1)
		goto L80
	} else {
		goto L82
	}
L81:
	;
	goto L68
L82:
	;
	goto L81
L83:
	;
	if v378 == int32(-1) {
		v746 = int32(-1)
		goto L1
	} else {
		goto L84
	}
L84:
	;
	v387 = *(*int64)(unsafe.Add(mBase, _c_F_rdbSaveInfoAuxFields[6]))
	if v387 <= int64(-1) {
		goto L89
	} else {
		goto L90
	}
L85:
	;
	v429 = F_rdbSaveAuxField(m, l0, int32(_a_F_rdbSaveInfoAuxFields_8), int32(11), v8, v428)
	mBase = m.M
	v430 = m.ExcPending
	if v430 != 0 {
		goto L2
	} else {
		goto L94
	}
L86:
	;
	v428 = int32(0)
	goto L85
L88:
	;
	v409 = F_ull2string(m, v405, v406, v407)
	mBase = m.M
	if v409 == int32(0) {
		goto L86
	} else {
		goto L92
	}
L89:
	;
	goto L91
L90:
	;
	v405 = v8
	v406 = int32(21)
	v407 = v387
	v408 = int32(0)
	goto L88
L91:
	;
	v396 = int32(45)
	*(*uint8)(unsafe.Add(mBase, uint32(v8))) = uint8(v396)
	v400 = int32(1)
	v405 = v8 + v400
	v406 = int32(20)
	v407 = int64(0) - v387
	v408 = v400
	goto L88
L92:
	;
	v428 = v409 + v408
	goto L85
L94:
	;
	if v429 == int32(-1) {
		v746 = int32(-1)
		goto L1
	} else {
		goto L95
	}
L95:
	;
	goto L54
L96:
	;
	v483 = F_rdbSaveAuxField(m, l0, int32(_a_F_rdbSaveInfoAuxFields_9), int32(8), v8, v482)
	mBase = m.M
	v484 = m.ExcPending
	if v484 != 0 {
		goto L2
	} else {
		goto L105
	}
L97:
	;
	v482 = int32(0)
	goto L96
L99:
	;
	v463 = F_ull2string(m, v459, v460, v461)
	mBase = m.M
	if v463 == int32(0) {
		goto L97
	} else {
		goto L103
	}
L100:
	;
	goto L102
L101:
	;
	v459 = v8
	v460 = int32(21)
	v461 = v441
	v462 = int32(0)
	goto L99
L102:
	;
	v450 = int32(45)
	*(*uint8)(unsafe.Add(mBase, uint32(v8))) = uint8(v450)
	v454 = int32(1)
	v459 = v8 + v454
	v460 = int32(20)
	v461 = int64(0) - v441
	v462 = v454
	goto L99
L103:
	;
	v482 = v463 + v462
	goto L96
L105:
	;
	if v483 == int32(-1) {
		v746 = int32(-1)
		goto L1
	} else {
		goto L106
	}
L106:
	;
	v487 = int32(0)
	v488 = *(*int32)(unsafe.Add(mBase, _c_F_rdbSaveInfoAuxFields[7]))
	if v488 == v487 {
		goto L107
	} else {
		goto L108
	}
L107:
	;
	v746 = int32(1)
	goto L1
L108:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8))) = v488
	*(*int32)(unsafe.Add(mBase, uint32(v8)+20)) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v8)+12)) = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v8)+4)) = int64(4294967295)
	goto L109
L109:
	;
	goto L110
L110:
	;
	v510 = v8 + int32(20)
	v511 = *(*int32)(unsafe.Add(mBase, uint32(v8)+16))
	if v511 != 0 {
		goto L114
	} else {
		goto L115
	}
L112:
	;
	if v606 == int32(0) {
		goto L107
	} else {
		goto L138
	}
L113:
	;
	v517 = v510
	v518 = v514
	goto L116
L114:
	;
	v514 = int32(1)
	goto L113
L115:
	;
	v514 = int32(0)
	goto L113
L116:
	;
	switch v518 {
	case 0:
		goto L121
	default:
		goto L120
	}
L118:
	;
	v518 = int32(0)
	goto L116
L119:
	;
	goto L112
L120:
	;
	v598 = *(*int32)(unsafe.Add(mBase, uint32(v517)))
	*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = v598
	if v598 == int32(0) {
		goto L118
	} else {
		goto L137
	}
L121:
	;
	v522 = *(*int32)(unsafe.Add(mBase, uint32(v8)+4))
	if v522 != int32(-1) {
		v561 = v522
		goto L122
	} else {
		goto L123
	}
L122:
	;
	v562 = int32(1)
	v563 = v561 + v562
	*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = v563
	v565 = int32(0)
	v568 = *(*int32)(unsafe.Add(mBase, uint32(v8)))
	v569 = *(*int32)(unsafe.Add(mBase, uint32(v8)+8))
	v573 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v568+v569+int32(26)))))
	if v573 == int32(255) {
		goto L131
	} else {
		goto L132
	}
L123:
	;
	v526 = *(*int32)(unsafe.Add(mBase, uint32(v8)+8))
	if v526 != 0 {
		v561 = int32(-1)
		goto L122
	} else {
		goto L124
	}
L124:
	;
	v527 = *(*int32)(unsafe.Add(mBase, uint32(v8)))
	v528 = *(*int32)(unsafe.Add(mBase, uint32(v8)+12))
	if v528 == int32(0) {
		goto L126
	} else {
		goto L127
	}
L125:
	;
	v555 = *(*int32)(unsafe.Add(mBase, uint32(v554)+20))
	if v555 != int32(-1) {
		goto L128
	} else {
		goto L129
	}
L126:
	;
	v535 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v527)+16)))
	v536 = int64(*(*int8)(unsafe.Add(mBase, uint32(v527)+27)))
	v537 = int64(*(*int32)(unsafe.Add(mBase, uint32(v527)+8)))
	v538 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v527)+12)))
	v539 = int64(*(*int8)(unsafe.Add(mBase, uint32(v527)+26)))
	v540 = int64(*(*int32)(unsafe.Add(mBase, uint32(v527)+4)))
	v541 = F_wangHash64(m, v540)
	mBase = m.M
	v543 = F_wangHash64(m, v539+v541)
	mBase = m.M
	v545 = F_wangHash64(m, v538+v543)
	mBase = m.M
	v547 = F_wangHash64(m, v537+v545)
	mBase = m.M
	v549 = F_wangHash64(m, v536+v547)
	mBase = m.M
	v551 = F_wangHash64(m, v535+v549)
	mBase = m.M
	*(*int64)(unsafe.Add(mBase, uint32(v8)+24)) = v551
	v553 = *(*int32)(unsafe.Add(mBase, uint32(v8)))
	v554 = v553
	goto L125
L127:
	;
	v531 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v527)+24)))
	v533 = v531 + int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v527)+24)) = uint16(v533)
	v554 = v527
	goto L125
L128:
	;
	v561 = v555 + int32(-1)
	goto L122
L129:
	;
	v558 = *(*int32)(unsafe.Add(mBase, uint32(v8)+4))
	v561 = v558
	goto L122
L130:
	;
	v588 = int32(2)
	v593 = *(*int32)(unsafe.Add(mBase, uint32(v568+v586<<(uint(v588)%32)+int32(4))))
	v517 = v593 + v587<<(uint(v588)%32)
	v518 = int32(1)
	goto L116
L131:
	;
	v577 = v565
	goto L133
L132:
	;
	v577 = v562 << (uint(v573) % 32)
	goto L133
L133:
	;
	if v563 < v577 {
		v586 = v569
		v587 = v563
		goto L130
	} else {
		goto L134
	}
L134:
	;
	if v569 != 0 {
		v606 = v565
		goto L119
	} else {
		goto L135
	}
L135:
	;
	v579 = *(*int32)(unsafe.Add(mBase, uint32(v568)+20))
	if v579 == int32(-1) {
		v606 = v565
		goto L119
	} else {
		goto L136
	}
L136:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v8)+4)) = int64(4294967296)
	v586 = int32(1)
	v587 = int32(0)
	goto L130
L137:
	;
	v602 = *(*int32)(unsafe.Add(mBase, uint32(v598)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v510))) = v602
	v606 = v598
	goto L119
L138:
	;
	v612 = *(*int32)(unsafe.Add(mBase, uint32(v606)+8))
	goto L139
L139:
	;
	v613 = *(*int32)(unsafe.Add(mBase, uint32(v612)))
	v614 = m.T0[v613].(func(*base.Module, int32) int32)(m, l1)
	mBase = m.M
	v615 = m.ExcPending
	if v615 != 0 {
		goto L2
	} else {
		goto L140
	}
L140:
	;
	if v614 == int32(0) {
		goto L110
	} else {
		goto L141
	}
L141:
	;
	v618 = *(*int32)(unsafe.Add(mBase, uint32(v606)))
	goto L142
L142:
	;
	if v618&int32(3) == int32(0) {
		v640 = v618
		goto L145
	} else {
		goto L146
	}
L143:
	;
	if v614&int32(3) == int32(0) {
		v695 = v614
		goto L161
	} else {
		goto L162
	}
L144:
	;
	v673 = v665 - v618
	goto L143
L145:
	;
	v644 = v640
	goto L153
L146:
	;
	v626 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v618))))
	if v626 != 0 {
		goto L147
	} else {
		goto L148
	}
L147:
	;
	v629 = v618
	goto L149
L148:
	;
	v673 = v618 - v618
	goto L143
L149:
	;
	v633 = v629 + int32(1)
	if v633&int32(3) == int32(0) {
		v640 = v633
		goto L145
	} else {
		goto L151
	}
L151:
	;
	v638 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v633))))
	if v638 != 0 {
		v629 = v633
		goto L149
	} else {
		goto L152
	}
L152:
	;
	v665 = v633
	goto L144
L153:
	;
	v650 = *(*int32)(unsafe.Add(mBase, uint32(v644)))
	v653 = int32(-2139062144)
	if (int32(16843008)-v650|v650)&v653 == v653 {
		v644 = v644 + int32(4)
		goto L153
	} else {
		goto L155
	}
L154:
	;
	v659 = v644
	goto L156
L155:
	;
	goto L154
L156:
	;
	v663 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v659))))
	if v663 != 0 {
		v659 = v659 + int32(1)
		goto L156
	} else {
		goto L158
	}
L157:
	;
	v665 = v659
	goto L144
L158:
	;
	goto L157
L159:
	;
	v729 = F_rdbSaveAuxField(m, l0, v618, v673, v614, v728)
	mBase = m.M
	v730 = m.ExcPending
	if v730 != 0 {
		goto L2
	} else {
		goto L175
	}
L160:
	;
	v728 = v720 - v614
	goto L159
L161:
	;
	v699 = v695
	goto L169
L162:
	;
	v681 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v614))))
	if v681 != 0 {
		goto L163
	} else {
		goto L164
	}
L163:
	;
	v684 = v614
	goto L165
L164:
	;
	v728 = v614 - v614
	goto L159
L165:
	;
	v688 = v684 + int32(1)
	if v688&int32(3) == int32(0) {
		v695 = v688
		goto L161
	} else {
		goto L167
	}
L167:
	;
	v693 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v688))))
	if v693 != 0 {
		v684 = v688
		goto L165
	} else {
		goto L168
	}
L168:
	;
	v720 = v688
	goto L160
L169:
	;
	v705 = *(*int32)(unsafe.Add(mBase, uint32(v699)))
	v708 = int32(-2139062144)
	if (int32(16843008)-v705|v705)&v708 == v708 {
		v699 = v699 + int32(4)
		goto L169
	} else {
		goto L171
	}
L170:
	;
	v714 = v699
	goto L172
L171:
	;
	goto L170
L172:
	;
	v718 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v714))))
	if v718 != 0 {
		v714 = v714 + int32(1)
		goto L172
	} else {
		goto L174
	}
L173:
	;
	v720 = v714
	goto L160
L174:
	;
	goto L173
L175:
	;
	F_sdsfree(m, v614)
	mBase = m.M
	v732 = m.ExcPending
	if v732 != 0 {
		goto L2
	} else {
		goto L176
	}
L176:
	;
	v733 = int32(-1)
	if v729 != v733 {
		goto L110
	} else {
		goto L177
	}
L177:
	;
	v746 = v733
	goto L1
}
func F_rdbSaveInternal(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v29 int32
	_ = v29
	var v36 int32
	_ = v36
	var v42 int32
	_ = v42
	var v43 int64
	_ = v43
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v75 int32
	_ = v75
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v97 int32
	_ = v97
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v106 int32
	_ = v106
	var v110 int32
	_ = v110
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v123 int32
	_ = v123
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v132 int32
	_ = v132
	var v136 int32
	_ = v136
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v167 int32
	_ = v167
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v173 int32
	_ = v173
	var v176 int32
	_ = v176
	var v184 int32
	_ = v184
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v198 int32
	_ = v198
	v7 = m.G0
	v9 = v7 - int32(4144)
	m.G0 = v9
	*(*int32)(unsafe.Add(mBase, uint32(v9)+44)) = int32(0)
	v14 = F_fopen(m, l1, int32(_a_F_rdbSaveInternal_0))
	mBase = m.M
	if v14 != 0 {
		v42 = F___memcpy(m, v9+int32(48), int32(_a_F_rdbSaveInternal_1), int32(80))
		mBase = m.M
		v43 = int64(0)
		*(*int64)(unsafe.Add(mBase, uint32(v42)+56)) = v43
		*(*int32)(unsafe.Add(mBase, uint32(v42)+48)) = v14
		*(*int64)(unsafe.Add(mBase, uint32(v42+int32(64)))) = v43
		v52 = int32(0)
		*(*uint8)(unsafe.Add(mBase, uint32(v42+int32(72)))) = uint8(v52)
		v55 = *(*int32)(unsafe.Add(mBase, _c_F_rdbSaveInternal[0]))
		if v55 == int32(0) {
		} else {
			v59 = v9 + int32(48)
			v61 = *(*int32)(unsafe.Add(mBase, uint32(v59)+4))
			if v61 != int32(980) {
			} else {
				*(*int64)(unsafe.Add(mBase, uint32(v59)+64)) = int64(4194304)
			}
			if l3&int32(16) != 0 {
			} else {
				v68 = v9 + int32(48)
				v70 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v68)+72)))
				v75 = v70&int32(254) | int32(1)
				*(*uint8)(unsafe.Add(mBase, uint32(v68)+72)) = uint8(v75)
			}
		}
		v82 = F_rdbSaveRio(m, l0, int32(80), v9+int32(48), v9+int32(44), l3, l2)
		mBase = m.M
		v83 = m.ExcPending
		if v83 != 0 {
			return int32(0)
		} else {
			if v82 != int32(-1) {
				v90 = F_fflush(m, v14)
				mBase = m.M
				v91 = m.ExcPending
				if v91 != 0 {
					return int32(0)
				} else {
					if v90 == int32(0) {
						v97 = *(*int32)(unsafe.Add(mBase, uint32(v14)+76))
						if int32(-1) < v97 {
							v101 = F___lockfile(m, v14)
							mBase = m.M
							v102 = *(*int32)(unsafe.Add(mBase, uint32(v14)+60))
							if v101 == int32(0) {
								v106 = v102
							} else {
								F___unlockfile(m, v14)
								mBase = m.M
								v106 = v102
							}
						} else {
							v100 = *(*int32)(unsafe.Add(mBase, uint32(v14)+60))
							v106 = v100
						}
						if int32(-1) < v106 {
							v114 = v106
						} else {
							v110 = F___errno_location(m)
							mBase = m.M
							*(*int32)(unsafe.Add(mBase, uint32(v110))) = int32(8)
							v114 = int32(-1)
						}
						v115 = F_fsync(m, v114)
						mBase = m.M
						if v115 == int32(0) {
							if l3&int32(16) != 0 {
							} else {
								v123 = *(*int32)(unsafe.Add(mBase, uint32(v14)+76))
								if int32(-1) < v123 {
									v127 = F___lockfile(m, v14)
									mBase = m.M
									v128 = *(*int32)(unsafe.Add(mBase, uint32(v14)+60))
									if v127 == int32(0) {
										v132 = v128
									} else {
										F___unlockfile(m, v14)
										mBase = m.M
										v132 = v128
									}
								} else {
									v126 = *(*int32)(unsafe.Add(mBase, uint32(v14)+60))
									v132 = v126
								}
								if int32(-1) < v132 {
								} else {
									v136 = F___errno_location(m)
									mBase = m.M
									*(*int32)(unsafe.Add(mBase, uint32(v136))) = int32(8)
								}
							}
							v161 = F_fclose(m, v14)
							mBase = m.M
							v162 = m.ExcPending
							if v162 != 0 {
								return int32(0)
							} else {
								if v161 == int32(0) {
									v198 = int32(0)
									m.G0 = v9 + int32(4144)
									return v198
								} else {
									v167 = int32(_a_F_rdbSaveInternal_2)
									v169 = int32(0)
									v170 = int32(9116376)
									v171 = *(*int32)(unsafe.Add(mBase, _c_F_rdbSaveInternal[1]))
									v173 = *(*int32)(unsafe.Add(mBase, _c_F_rdbSaveInternal[2]))
									if int32(3) < v173 {
										if v169 == int32(0) {
											v189 = F_unlink(m, l1)
											mBase = m.M
											*(*int32)(unsafe.Add(mBase, _c_F_rdbSaveInternal[1])) = v171
											v198 = int32(-1)
											m.G0 = v9 + int32(4144)
											return v198
										} else {
											v187 = F_fclose(m, v169)
											mBase = m.M
											v188 = m.ExcPending
											if v188 != 0 {
												return int32(0)
											} else {
												v189 = F_unlink(m, l1)
												mBase = m.M
												*(*int32)(unsafe.Add(mBase, _c_F_rdbSaveInternal[1])) = v171
												v198 = int32(-1)
												m.G0 = v9 + int32(4144)
												return v198
											}
										}
									} else {
										v176 = F___strerror_l(m, v171, v171)
										mBase = m.M
										*(*int32)(unsafe.Add(mBase, uint32(v9)+20)) = v176
										*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = v167
										F__serverLog(m, int32(3), int32(_a_F_rdbSaveInternal_3), v9+int32(16))
										mBase = m.M
										v184 = m.ExcPending
										if v184 != 0 {
											return int32(0)
										} else {
											if v169 == int32(0) {
												v189 = F_unlink(m, l1)
												mBase = m.M
												*(*int32)(unsafe.Add(mBase, _c_F_rdbSaveInternal[1])) = v171
												v198 = int32(-1)
												m.G0 = v9 + int32(4144)
												return v198
											} else {
												v187 = F_fclose(m, v169)
												mBase = m.M
												v188 = m.ExcPending
												if v188 != 0 {
													return int32(0)
												} else {
													v189 = F_unlink(m, l1)
													mBase = m.M
													*(*int32)(unsafe.Add(mBase, _c_F_rdbSaveInternal[1])) = v171
													v198 = int32(-1)
													m.G0 = v9 + int32(4144)
													return v198
												}
											}
										}
									}
								}
							}
						} else {
							v167 = int32(_a_F_rdbSaveInternal_4)
							v169 = v14
							v170 = int32(9116376)
							v171 = *(*int32)(unsafe.Add(mBase, _c_F_rdbSaveInternal[1]))
							v173 = *(*int32)(unsafe.Add(mBase, _c_F_rdbSaveInternal[2]))
							if int32(3) < v173 {
								if v169 == int32(0) {
									v189 = F_unlink(m, l1)
									mBase = m.M
									*(*int32)(unsafe.Add(mBase, _c_F_rdbSaveInternal[1])) = v171
									v198 = int32(-1)
									m.G0 = v9 + int32(4144)
									return v198
								} else {
									v187 = F_fclose(m, v169)
									mBase = m.M
									v188 = m.ExcPending
									if v188 != 0 {
										return int32(0)
									} else {
										v189 = F_unlink(m, l1)
										mBase = m.M
										*(*int32)(unsafe.Add(mBase, _c_F_rdbSaveInternal[1])) = v171
										v198 = int32(-1)
										m.G0 = v9 + int32(4144)
										return v198
									}
								}
							} else {
								v176 = F___strerror_l(m, v171, v171)
								mBase = m.M
								*(*int32)(unsafe.Add(mBase, uint32(v9)+20)) = v176
								*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = v167
								F__serverLog(m, int32(3), int32(_a_F_rdbSaveInternal_3), v9+int32(16))
								mBase = m.M
								v184 = m.ExcPending
								if v184 != 0 {
									return int32(0)
								} else {
									if v169 == int32(0) {
										v189 = F_unlink(m, l1)
										mBase = m.M
										*(*int32)(unsafe.Add(mBase, _c_F_rdbSaveInternal[1])) = v171
										v198 = int32(-1)
										m.G0 = v9 + int32(4144)
										return v198
									} else {
										v187 = F_fclose(m, v169)
										mBase = m.M
										v188 = m.ExcPending
										if v188 != 0 {
											return int32(0)
										} else {
											v189 = F_unlink(m, l1)
											mBase = m.M
											*(*int32)(unsafe.Add(mBase, _c_F_rdbSaveInternal[1])) = v171
											v198 = int32(-1)
											m.G0 = v9 + int32(4144)
											return v198
										}
									}
								}
							}
						}
					} else {
						v167 = int32(_a_F_rdbSaveInternal_5)
						v169 = v14
						v170 = int32(9116376)
						v171 = *(*int32)(unsafe.Add(mBase, _c_F_rdbSaveInternal[1]))
						v173 = *(*int32)(unsafe.Add(mBase, _c_F_rdbSaveInternal[2]))
						if int32(3) < v173 {
							if v169 == int32(0) {
								v189 = F_unlink(m, l1)
								mBase = m.M
								*(*int32)(unsafe.Add(mBase, _c_F_rdbSaveInternal[1])) = v171
								v198 = int32(-1)
								m.G0 = v9 + int32(4144)
								return v198
							} else {
								v187 = F_fclose(m, v169)
								mBase = m.M
								v188 = m.ExcPending
								if v188 != 0 {
									return int32(0)
								} else {
									v189 = F_unlink(m, l1)
									mBase = m.M
									*(*int32)(unsafe.Add(mBase, _c_F_rdbSaveInternal[1])) = v171
									v198 = int32(-1)
									m.G0 = v9 + int32(4144)
									return v198
								}
							}
						} else {
							v176 = F___strerror_l(m, v171, v171)
							mBase = m.M
							*(*int32)(unsafe.Add(mBase, uint32(v9)+20)) = v176
							*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = v167
							F__serverLog(m, int32(3), int32(_a_F_rdbSaveInternal_3), v9+int32(16))
							mBase = m.M
							v184 = m.ExcPending
							if v184 != 0 {
								return int32(0)
							} else {
								if v169 == int32(0) {
									v189 = F_unlink(m, l1)
									mBase = m.M
									*(*int32)(unsafe.Add(mBase, _c_F_rdbSaveInternal[1])) = v171
									v198 = int32(-1)
									m.G0 = v9 + int32(4144)
									return v198
								} else {
									v187 = F_fclose(m, v169)
									mBase = m.M
									v188 = m.ExcPending
									if v188 != 0 {
										return int32(0)
									} else {
										v189 = F_unlink(m, l1)
										mBase = m.M
										*(*int32)(unsafe.Add(mBase, _c_F_rdbSaveInternal[1])) = v171
										v198 = int32(-1)
										m.G0 = v9 + int32(4144)
										return v198
									}
								}
							}
						}
					}
				}
			} else {
				v87 = *(*int32)(unsafe.Add(mBase, uint32(v9)+44))
				*(*int32)(unsafe.Add(mBase, _c_F_rdbSaveInternal[1])) = v87
				v167 = int32(_a_F_rdbSaveInternal_6)
				v169 = v14
				v170 = int32(9116376)
				v171 = *(*int32)(unsafe.Add(mBase, _c_F_rdbSaveInternal[1]))
				v173 = *(*int32)(unsafe.Add(mBase, _c_F_rdbSaveInternal[2]))
				if int32(3) < v173 {
					if v169 == int32(0) {
						v189 = F_unlink(m, l1)
						mBase = m.M
						*(*int32)(unsafe.Add(mBase, _c_F_rdbSaveInternal[1])) = v171
						v198 = int32(-1)
						m.G0 = v9 + int32(4144)
						return v198
					} else {
						v187 = F_fclose(m, v169)
						mBase = m.M
						v188 = m.ExcPending
						if v188 != 0 {
							return int32(0)
						} else {
							v189 = F_unlink(m, l1)
							mBase = m.M
							*(*int32)(unsafe.Add(mBase, _c_F_rdbSaveInternal[1])) = v171
							v198 = int32(-1)
							m.G0 = v9 + int32(4144)
							return v198
						}
					}
				} else {
					v176 = F___strerror_l(m, v171, v171)
					mBase = m.M
					*(*int32)(unsafe.Add(mBase, uint32(v9)+20)) = v176
					*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = v167
					F__serverLog(m, int32(3), int32(_a_F_rdbSaveInternal_3), v9+int32(16))
					mBase = m.M
					v184 = m.ExcPending
					if v184 != 0 {
						return int32(0)
					} else {
						if v169 == int32(0) {
							v189 = F_unlink(m, l1)
							mBase = m.M
							*(*int32)(unsafe.Add(mBase, _c_F_rdbSaveInternal[1])) = v171
							v198 = int32(-1)
							m.G0 = v9 + int32(4144)
							return v198
						} else {
							v187 = F_fclose(m, v169)
							mBase = m.M
							v188 = m.ExcPending
							if v188 != 0 {
								return int32(0)
							} else {
								v189 = F_unlink(m, l1)
								mBase = m.M
								*(*int32)(unsafe.Add(mBase, _c_F_rdbSaveInternal[1])) = v171
								v198 = int32(-1)
								m.G0 = v9 + int32(4144)
								return v198
							}
						}
					}
				}
			}
		}
	} else {
		v15 = int32(9116376)
		v16 = *(*int32)(unsafe.Add(mBase, _c_F_rdbSaveInternal[1]))
		v17 = F___strerror_l(m, v16, v16)
		mBase = m.M
		v21 = F_getcwd(m, v9+int32(48), int32(4096))
		mBase = m.M
		v23 = *(*int32)(unsafe.Add(mBase, _c_F_rdbSaveInternal[2]))
		if int32(3) < v23 {
			*(*int32)(unsafe.Add(mBase, _c_F_rdbSaveInternal[1])) = v16
			v198 = int32(-1)
			m.G0 = v9 + int32(4144)
			return v198
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v9)+8)) = v17
			*(*int32)(unsafe.Add(mBase, uint32(v9))) = l1
			if v21 != 0 {
				v29 = v21
			} else {
				v29 = int32(_a_F_rdbSaveInternal_7)
			}
			*(*int32)(unsafe.Add(mBase, uint32(v9)+4)) = v29
			F__serverLog(m, int32(3), int32(_a_F_rdbSaveInternal_8), v9)
			mBase = m.M
			v36 = m.ExcPending
			if v36 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, _c_F_rdbSaveInternal[1])) = v16
				v198 = int32(-1)
				m.G0 = v9 + int32(4144)
				return v198
			}
		}
	}
}
func F_rdbSaveRioWithEOFMark(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v34 int32
	_ = v34
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
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
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v96 int32
	_ = v96
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v116 int32
	_ = v116
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v130 int32
	_ = v130
	var v134 int32
	_ = v134
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v140 int32
	_ = v140
	var v148 int32
	_ = v148
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v168 int32
	_ = v168
	var v174 int32
	_ = v174
	var v185 int64
	_ = v185
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v203 int32
	_ = v203
	var v205 int32
	_ = v205
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v220 int32
	_ = v220
	v11 = m.G0
	v13 = v11 - int32(48)
	m.G0 = v13
	v16 = F___syscall_getpid(m)
	mBase = m.M
	goto L1
L1:
	;
	v17 = int32(0)
	v18 = *(*int32)(unsafe.Add(mBase, _c_F_rdbSaveRioWithEOFMark[0]))
	F_moduleFireServerEvent(m, int64(1), base.B2i32(v16 == v18)<<(uint(int32(1))%32), v17)
	mBase = m.M
	v26 = m.ExcPending
	if v26 != 0 {
		goto L2
	} else {
		goto L3
	}
L2:
	;
	return int32(0)
L3:
	;
	F_getRandomHexChars(m, v13, int32(40))
	mBase = m.M
	v29 = m.ExcPending
	if v29 != 0 {
		goto L2
	} else {
		goto L4
	}
L4:
	;
	if l3 == int32(0) {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	v34 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+32)))
	if v34&int32(6) != 0 {
		goto L8
	} else {
		goto L9
	}
L6:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = int32(0)
	goto L5
L7:
	;
	F_moduleFireServerEvent(m, int64(1), v214, int32(0))
	mBase = m.M
	v220 = m.ExcPending
	if v220 != 0 {
		goto L2
	} else {
		goto L75
	}
L8:
	;
	v199 = int32(-1)
	v200 = int32(4)
	if l3 == int32(0) {
		v214 = v200
		v215 = v199
		goto L7
	} else {
		goto L72
	}
L9:
	;
	v45 = int32(_a_F_rdbSaveRioWithEOFMark_0)
	v46 = int32(5)
	goto L11
L10:
	;
	v185 = *(*int64)(unsafe.Add(mBase, uint32(l2)+32))
	*(*int64)(unsafe.Add(mBase, uint32(l2)+32)) = v185 | int64(2)
	goto L8
L11:
	;
	v49 = *(*int32)(unsafe.Add(mBase, uint32(l2)+44))
	if base.Ui32(v49) < base.Ui32(v46) {
		goto L13
	} else {
		goto L14
	}
L12:
	;
	v68 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+32)))
	if v68&int32(6) != 0 {
		goto L8
	} else {
		goto L25
	}
L13:
	;
	v51 = v49
	goto L15
L14:
	;
	v51 = v46
	goto L15
L15:
	;
	if v49 != 0 {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	v52 = v51
	goto L18
L17:
	;
	v52 = v46
	goto L18
L18:
	;
	v53 = *(*int32)(unsafe.Add(mBase, uint32(l2)+16))
	if v53 == int32(0) {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	v58 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	v59 = m.T0[v58].(func(*base.Module, int32, int32, int32) int32)(m, l2, v45, v52)
	mBase = m.M
	v60 = m.ExcPending
	if v60 != 0 {
		goto L2
	} else {
		goto L22
	}
L20:
	;
	m.T0[v53].(func(*base.Module, int32, int32, int32))(m, l2, v45, v52)
	mBase = m.M
	v57 = m.ExcPending
	if v57 != 0 {
		goto L2
	} else {
		goto L21
	}
L21:
	;
	goto L19
L22:
	;
	if v59 == int32(0) {
		goto L10
	} else {
		goto L23
	}
L23:
	;
	v63 = *(*int32)(unsafe.Add(mBase, uint32(l2)+40))
	*(*int32)(unsafe.Add(mBase, uint32(l2)+40)) = v63 + v52
	v67 = v46 - v52
	if v67 != 0 {
		v45 = v45 + v52
		v46 = v67
		goto L11
	} else {
		goto L24
	}
L24:
	;
	goto L12
L25:
	;
	v78 = v13
	v79 = int32(40)
	goto L26
L26:
	;
	v82 = *(*int32)(unsafe.Add(mBase, uint32(l2)+44))
	if base.Ui32(v82) < base.Ui32(v79) {
		goto L28
	} else {
		goto L29
	}
L27:
	;
	v101 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+32)))
	if v101&int32(6) != 0 {
		goto L8
	} else {
		goto L40
	}
L28:
	;
	v84 = v82
	goto L30
L29:
	;
	v84 = v79
	goto L30
L30:
	;
	if v82 != 0 {
		goto L31
	} else {
		goto L32
	}
L31:
	;
	v85 = v84
	goto L33
L32:
	;
	v85 = v79
	goto L33
L33:
	;
	v86 = *(*int32)(unsafe.Add(mBase, uint32(l2)+16))
	if v86 == int32(0) {
		goto L34
	} else {
		goto L35
	}
L34:
	;
	v91 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	v92 = m.T0[v91].(func(*base.Module, int32, int32, int32) int32)(m, l2, v78, v85)
	mBase = m.M
	v93 = m.ExcPending
	if v93 != 0 {
		goto L2
	} else {
		goto L37
	}
L35:
	;
	m.T0[v86].(func(*base.Module, int32, int32, int32))(m, l2, v78, v85)
	mBase = m.M
	v90 = m.ExcPending
	if v90 != 0 {
		goto L2
	} else {
		goto L36
	}
L36:
	;
	goto L34
L37:
	;
	if v92 == int32(0) {
		goto L10
	} else {
		goto L38
	}
L38:
	;
	v96 = *(*int32)(unsafe.Add(mBase, uint32(l2)+40))
	*(*int32)(unsafe.Add(mBase, uint32(l2)+40)) = v96 + v85
	v100 = v79 - v85
	if v100 != 0 {
		v78 = v78 + v85
		v79 = v100
		goto L26
	} else {
		goto L39
	}
L39:
	;
	goto L27
L40:
	;
	v112 = int32(_a_F_rdbSaveRioWithEOFMark_1)
	v113 = int32(2)
	goto L41
L41:
	;
	v116 = *(*int32)(unsafe.Add(mBase, uint32(l2)+44))
	if base.Ui32(v116) < base.Ui32(v113) {
		goto L43
	} else {
		goto L44
	}
L42:
	;
	v136 = F_rdbSaveRio(m, l0, l1, l2, l3, int32(2), l4)
	mBase = m.M
	v137 = m.ExcPending
	if v137 != 0 {
		goto L2
	} else {
		goto L55
	}
L43:
	;
	v118 = v116
	goto L45
L44:
	;
	v118 = v113
	goto L45
L45:
	;
	if v116 != 0 {
		goto L46
	} else {
		goto L47
	}
L46:
	;
	v119 = v118
	goto L48
L47:
	;
	v119 = v113
	goto L48
L48:
	;
	v120 = *(*int32)(unsafe.Add(mBase, uint32(l2)+16))
	if v120 == int32(0) {
		goto L49
	} else {
		goto L50
	}
L49:
	;
	v125 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	v126 = m.T0[v125].(func(*base.Module, int32, int32, int32) int32)(m, l2, v112, v119)
	mBase = m.M
	v127 = m.ExcPending
	if v127 != 0 {
		goto L2
	} else {
		goto L52
	}
L50:
	;
	m.T0[v120].(func(*base.Module, int32, int32, int32))(m, l2, v112, v119)
	mBase = m.M
	v124 = m.ExcPending
	if v124 != 0 {
		goto L2
	} else {
		goto L51
	}
L51:
	;
	goto L49
L52:
	;
	if v126 == int32(0) {
		goto L10
	} else {
		goto L53
	}
L53:
	;
	v130 = *(*int32)(unsafe.Add(mBase, uint32(l2)+40))
	*(*int32)(unsafe.Add(mBase, uint32(l2)+40)) = v130 + v119
	v134 = v113 - v119
	if v134 != 0 {
		v112 = v112 + v119
		v113 = v134
		goto L41
	} else {
		goto L54
	}
L54:
	;
	goto L42
L55:
	;
	if v136 == int32(-1) {
		goto L8
	} else {
		goto L56
	}
L56:
	;
	v140 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+32)))
	if v140&int32(6) != 0 {
		goto L8
	} else {
		goto L57
	}
L57:
	;
	v148 = v13
	v153 = int32(40)
	goto L58
L58:
	;
	v154 = *(*int32)(unsafe.Add(mBase, uint32(l2)+44))
	if base.Ui32(v154) < base.Ui32(v153) {
		goto L60
	} else {
		goto L61
	}
L60:
	;
	v156 = v154
	goto L62
L61:
	;
	v156 = v153
	goto L62
L62:
	;
	if v154 != 0 {
		goto L63
	} else {
		goto L64
	}
L63:
	;
	v157 = v156
	goto L65
L64:
	;
	v157 = v153
	goto L65
L65:
	;
	v158 = *(*int32)(unsafe.Add(mBase, uint32(l2)+16))
	if v158 == int32(0) {
		goto L66
	} else {
		goto L67
	}
L66:
	;
	v163 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	v164 = m.T0[v163].(func(*base.Module, int32, int32, int32) int32)(m, l2, v148, v157)
	mBase = m.M
	v165 = m.ExcPending
	if v165 != 0 {
		goto L2
	} else {
		goto L69
	}
L67:
	;
	m.T0[v158].(func(*base.Module, int32, int32, int32))(m, l2, v148, v157)
	mBase = m.M
	v162 = m.ExcPending
	if v162 != 0 {
		goto L2
	} else {
		goto L68
	}
L68:
	;
	goto L66
L69:
	;
	if v164 == int32(0) {
		goto L10
	} else {
		goto L70
	}
L70:
	;
	v168 = *(*int32)(unsafe.Add(mBase, uint32(l2)+40))
	*(*int32)(unsafe.Add(mBase, uint32(l2)+40)) = v168 + v157
	v174 = v153 - v157
	if v174 != 0 {
		v148 = v148 + v157
		v153 = v174
		goto L58
	} else {
		goto L71
	}
L71:
	;
	v214 = int32(3)
	v215 = int32(0)
	goto L7
L72:
	;
	v203 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	if v203 != 0 {
		v214 = v200
		v215 = v199
		goto L7
	} else {
		goto L73
	}
L73:
	;
	goto L74
L74:
	;
	v205 = *(*int32)(unsafe.Add(mBase, _c_F_rdbSaveRioWithEOFMark[1]))
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v205
	v214 = v200
	v215 = v199
	goto L7
L75:
	;
	m.G0 = v13 + int32(48)
	return v215
}
func F_removeRDBUsedToSyncReplicas(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
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
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v51 int32
	_ = v51
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v63 int32
	_ = v63
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v75 int32
	_ = v75
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v97 int32
	_ = v97
	var v100 int32
	_ = v100
	v5 = m.G0
	v7 = v5 - int32(112)
	m.G0 = v7
	v10 = *(*int32)(unsafe.Add(mBase, _c_F_removeRDBUsedToSyncReplicas[0]))
	if v10 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v7 + int32(112)
	return
L2:
	;
	v14 = int32(0)
	v15 = *(*int32)(unsafe.Add(mBase, _c_F_removeRDBUsedToSyncReplicas[1]))
	v17 = *(*int32)(unsafe.Add(mBase, _c_F_removeRDBUsedToSyncReplicas[2]))
	goto L4
L3:
	;
	v11 = int32(0)
	*(*int32)(unsafe.Add(mBase, _c_F_removeRDBUsedToSyncReplicas[3])) = v11
	goto L1
L4:
	;
	if base.B2i32(v15|v17 == v14) == int32(0) {
		goto L1
	} else {
		goto L5
	}
L5:
	;
	v23 = int32(0)
	v24 = *(*int32)(unsafe.Add(mBase, _c_F_removeRDBUsedToSyncReplicas[3]))
	if v24 == v23 {
		goto L1
	} else {
		goto L6
	}
L6:
	;
	v28 = *(*int32)(unsafe.Add(mBase, _c_F_removeRDBUsedToSyncReplicas[4]))
	v30 = v7 + int32(104)
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v28)))
	*(*int32)(unsafe.Add(mBase, uint32(v30)+4)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v30))) = v31
	goto L7
L7:
	;
	goto L9
L8:
	;
	v63 = *(*int32)(unsafe.Add(mBase, _c_F_removeRDBUsedToSyncReplicas[5]))
	v68 = F___fstatat(m, int32(-100), v63, v7+int32(8), int32(256))
	mBase = m.M
	goto L16
L9:
	;
	v40 = v7 + int32(104)
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v40)))
	if v42 == int32(0) {
		goto L12
	} else {
		goto L13
	}
L11:
	;
	if v42 == int32(0) {
		goto L8
	} else {
		goto L14
	}
L12:
	;
	goto L11
L13:
	;
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v40)+4))
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v42+base.B2i32(v45 == int32(0))<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v40))) = v51
	goto L12
L14:
	;
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v42)+8))
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v55)+104))
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v56)))
	if base.Ui32(int32(3)) <= base.Ui32(v57+int32(-6)) {
		goto L9
	} else {
		goto L15
	}
L15:
	;
	goto L1
L16:
	;
	if v68 == int32(-1) {
		goto L1
	} else {
		goto L17
	}
L17:
	;
	v71 = int32(0)
	*(*int32)(unsafe.Add(mBase, _c_F_removeRDBUsedToSyncReplicas[3])) = v71
	v75 = *(*int32)(unsafe.Add(mBase, _c_F_removeRDBUsedToSyncReplicas[6]))
	if int32(2) < v75 {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	v84 = *(*int32)(unsafe.Add(mBase, _c_F_removeRDBUsedToSyncReplicas[5]))
	v87 = F_open(m, v84, int32(2048), int32(0))
	mBase = m.M
	v88 = F_unlink(m, v84)
	mBase = m.M
	if v87 == int32(-1) {
		goto L1
	} else {
		goto L22
	}
L19:
	;
	F__serverLog(m, int32(2), int32(_a_F_removeRDBUsedToSyncReplicas_0), int32(0))
	mBase = m.M
	v82 = m.ExcPending
	if v82 != 0 {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	return
L21:
	;
	goto L18
L22:
	;
	if v88 != int32(-1) {
		goto L23
	} else {
		goto L24
	}
L23:
	;
	v97 = int32(0)
	F_bioCreateCloseJob(m, v87, v97, v97)
	mBase = m.M
	v100 = m.ExcPending
	if v100 != 0 {
		goto L20
	} else {
		goto L26
	}
L24:
	;
	v93 = int32(9116376)
	goto L25
L25:
	;
	v94 = *(*int32)(unsafe.Add(mBase, _c_F_removeRDBUsedToSyncReplicas[7]))
	v95 = F_close(m, v87)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, _c_F_removeRDBUsedToSyncReplicas[7])) = v94
	goto L1
L26:
	;
	goto L1
}
