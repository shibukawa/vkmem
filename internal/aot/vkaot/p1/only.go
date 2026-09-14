package p1

import (
	base "github.com/shibukawa/vkmem/internal/aot/vkaot/base"
	"unsafe"
)

func F_getAppendOnlyFileSize(m *base.Module, l0 int32, l1 int32) int64 {
	mBase := m.M
	_ = mBase
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
	var v20 int64
	_ = v20
	var v26 int64
	_ = v26
	var v27 int64
	_ = v27
	var v32 int32
	_ = v32
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v45 int64
	_ = v45
	var v47 int32
	_ = v47
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v58 int32
	_ = v58
	var v63 int64
	_ = v63
	var v64 int64
	_ = v64
	var v66 int64
	_ = v66
	var v69 int64
	_ = v69
	var v71 int64
	_ = v71
	var v74 int64
	_ = v74
	var v80 int32
	_ = v80
	var v85 int32
	_ = v85
	v9 = m.G0
	v11 = v9 - int32(112)
	m.G0 = v11
	v14 = *(*int32)(unsafe.Add(mBase, _c_F_getAppendOnlyFileSize[0]))
	v15 = F_makePath(m, v14, l0)
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		return int64(0)
	} else {
		v20 = *(*int64)(unsafe.Add(mBase, _c_F_getAppendOnlyFileSize[1]))
		if base.B2i32(v20 == int64(0)) == int32(0) {
			v26 = F_ustime(m)
			mBase = m.M
			v27 = v26
		} else {
			v27 = int64(0)
		}
		v32 = F___fstatat(m, int32(-100), v15, v11+int32(16), int32(0))
		mBase = m.M
		if v32 != int32(-1) {
			if l1 == int32(0) {
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(l1))) = int32(0)
			}
			v63 = *(*int64)(unsafe.Add(mBase, uint32(v11)+40))
			v64 = v63
			v66 = *(*int64)(unsafe.Add(mBase, _c_F_getAppendOnlyFileSize[1]))
			if v66 == int64(0) {
				F_sdsfree(m, v15)
				mBase = m.M
				v85 = m.ExcPending
				if v85 != 0 {
					return int64(0)
				} else {
					m.G0 = v11 + int32(112)
					return v64
				}
			} else {
				v69 = F_ustime(m)
				mBase = m.M
				v71 = *(*int64)(unsafe.Add(mBase, _c_F_getAppendOnlyFileSize[1]))
				if v71 == int64(0) {
					F_sdsfree(m, v15)
					mBase = m.M
					v85 = m.ExcPending
					if v85 != 0 {
						return int64(0)
					} else {
						m.G0 = v11 + int32(112)
						return v64
					}
				} else {
					v74 = v69 - v27
					if v74 < v71*int64(1000) {
						F_sdsfree(m, v15)
						mBase = m.M
						v85 = m.ExcPending
						if v85 != 0 {
							return int64(0)
						} else {
							m.G0 = v11 + int32(112)
							return v64
						}
					} else {
						F_latencyAddSample(m, int32(_a_F_getAppendOnlyFileSize_0), v74)
						mBase = m.M
						v80 = m.ExcPending
						if v80 != 0 {
							return int64(0)
						} else {
							F_sdsfree(m, v15)
							mBase = m.M
							v85 = m.ExcPending
							if v85 != 0 {
								return int64(0)
							} else {
								m.G0 = v11 + int32(112)
								return v64
							}
						}
					}
				}
			}
		} else {
			if l1 == int32(0) {
			} else {
				v40 = *(*int32)(unsafe.Add(mBase, _c_F_getAppendOnlyFileSize[2]))
				if v40 == int32(44) {
					v43 = int32(1)
				} else {
					v43 = int32(3)
				}
				*(*int32)(unsafe.Add(mBase, uint32(l1))) = v43
			}
			v45 = int64(0)
			v47 = *(*int32)(unsafe.Add(mBase, _c_F_getAppendOnlyFileSize[3]))
			if int32(3) < v47 {
				v64 = v45
				v66 = *(*int64)(unsafe.Add(mBase, _c_F_getAppendOnlyFileSize[1]))
				if v66 == int64(0) {
					F_sdsfree(m, v15)
					mBase = m.M
					v85 = m.ExcPending
					if v85 != 0 {
						return int64(0)
					} else {
						m.G0 = v11 + int32(112)
						return v64
					}
				} else {
					v69 = F_ustime(m)
					mBase = m.M
					v71 = *(*int64)(unsafe.Add(mBase, _c_F_getAppendOnlyFileSize[1]))
					if v71 == int64(0) {
						F_sdsfree(m, v15)
						mBase = m.M
						v85 = m.ExcPending
						if v85 != 0 {
							return int64(0)
						} else {
							m.G0 = v11 + int32(112)
							return v64
						}
					} else {
						v74 = v69 - v27
						if v74 < v71*int64(1000) {
							F_sdsfree(m, v15)
							mBase = m.M
							v85 = m.ExcPending
							if v85 != 0 {
								return int64(0)
							} else {
								m.G0 = v11 + int32(112)
								return v64
							}
						} else {
							F_latencyAddSample(m, int32(_a_F_getAppendOnlyFileSize_0), v74)
							mBase = m.M
							v80 = m.ExcPending
							if v80 != 0 {
								return int64(0)
							} else {
								F_sdsfree(m, v15)
								mBase = m.M
								v85 = m.ExcPending
								if v85 != 0 {
									return int64(0)
								} else {
									m.G0 = v11 + int32(112)
									return v64
								}
							}
						}
					}
				}
			} else {
				v51 = *(*int32)(unsafe.Add(mBase, _c_F_getAppendOnlyFileSize[2]))
				v52 = F___strerror_l(m, v51, v51)
				mBase = m.M
				*(*int32)(unsafe.Add(mBase, uint32(v11)+4)) = v52
				*(*int32)(unsafe.Add(mBase, uint32(v11))) = l0
				F__serverLog(m, int32(3), int32(_a_F_getAppendOnlyFileSize_1), v11)
				mBase = m.M
				v58 = m.ExcPending
				if v58 != 0 {
					return int64(0)
				} else {
					v64 = v45
					v66 = *(*int64)(unsafe.Add(mBase, _c_F_getAppendOnlyFileSize[1]))
					if v66 == int64(0) {
						F_sdsfree(m, v15)
						mBase = m.M
						v85 = m.ExcPending
						if v85 != 0 {
							return int64(0)
						} else {
							m.G0 = v11 + int32(112)
							return v64
						}
					} else {
						v69 = F_ustime(m)
						mBase = m.M
						v71 = *(*int64)(unsafe.Add(mBase, _c_F_getAppendOnlyFileSize[1]))
						if v71 == int64(0) {
							F_sdsfree(m, v15)
							mBase = m.M
							v85 = m.ExcPending
							if v85 != 0 {
								return int64(0)
							} else {
								m.G0 = v11 + int32(112)
								return v64
							}
						} else {
							v74 = v69 - v27
							if v74 < v71*int64(1000) {
								F_sdsfree(m, v15)
								mBase = m.M
								v85 = m.ExcPending
								if v85 != 0 {
									return int64(0)
								} else {
									m.G0 = v11 + int32(112)
									return v64
								}
							} else {
								F_latencyAddSample(m, int32(_a_F_getAppendOnlyFileSize_0), v74)
								mBase = m.M
								v80 = m.ExcPending
								if v80 != 0 {
									return int64(0)
								} else {
									F_sdsfree(m, v15)
									mBase = m.M
									v85 = m.ExcPending
									if v85 != 0 {
										return int64(0)
									} else {
										m.G0 = v11 + int32(112)
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
func F_rewriteAppendOnlyFile(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v37 int32
	_ = v37
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
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v73 int32
	_ = v73
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v95 int32
	_ = v95
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v107 int32
	_ = v107
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v116 int32
	_ = v116
	var v120 int32
	_ = v120
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v129 int32
	_ = v129
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v138 int32
	_ = v138
	var v142 int32
	_ = v142
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v170 int32
	_ = v170
	var v174 int32
	_ = v174
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v186 int32
	_ = v186
	var v189 int32
	_ = v189
	var v192 int32
	_ = v192
	var v196 int32
	_ = v196
	var v198 int32
	_ = v198
	var v200 int32
	_ = v200
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v212 int32
	_ = v212
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v219 int32
	_ = v219
	var v222 int32
	_ = v222
	var v224 int32
	_ = v224
	v5 = m.G0
	v7 = v5 - int32(416)
	m.G0 = v7
	v9 = F___syscall_getpid(m)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v7)+64)) = v9
	v17 = F_snprintf(m, v7+int32(80), int32(256), int32(_a_F_rewriteAppendOnlyFile_0), v7+int32(64))
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		return int32(0)
	} else {
		v24 = F_fopen(m, v7+int32(80), int32(_a_F_rewriteAppendOnlyFile_1))
		mBase = m.M
		if v24 != 0 {
			v42 = F___memcpy(m, v7+int32(336), int32(_a_F_rewriteAppendOnlyFile_2), int32(80))
			mBase = m.M
			v43 = int64(0)
			*(*int64)(unsafe.Add(mBase, uint32(v42)+56)) = v43
			*(*int32)(unsafe.Add(mBase, uint32(v42)+48)) = v24
			*(*int64)(unsafe.Add(mBase, uint32(v42+int32(64)))) = v43
			v52 = int32(0)
			*(*uint8)(unsafe.Add(mBase, uint32(v42+int32(72)))) = uint8(v52)
			v55 = *(*int32)(unsafe.Add(mBase, _c_F_rewriteAppendOnlyFile[0]))
			if v55 == int32(0) {
			} else {
				v59 = v7 + int32(336)
				v61 = *(*int32)(unsafe.Add(mBase, uint32(v59)+4))
				if v61 != int32(980) {
				} else {
					*(*int64)(unsafe.Add(mBase, uint32(v59)+64)) = int64(4194304)
				}
				v66 = v7 + int32(336)
				v68 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v66)+72)))
				v73 = v68&int32(254) | int32(1)
				*(*uint8)(unsafe.Add(mBase, uint32(v66)+72)) = uint8(v73)
			}
			F_startSaving(m, int32(1))
			mBase = m.M
			v77 = m.ExcPending
			if v77 != 0 {
				return int32(0)
			} else {
				v79 = *(*int32)(unsafe.Add(mBase, _c_F_rewriteAppendOnlyFile[1]))
				if v79 == int32(0) {
					v99 = F_rewriteAppendOnlyFileRio(m, v7+int32(336))
					mBase = m.M
					v100 = m.ExcPending
					if v100 != 0 {
						return int32(0)
					} else {
						if v99 == int32(-1) {
							v198 = v24
							v200 = *(*int32)(unsafe.Add(mBase, _c_F_rewriteAppendOnlyFile[2]))
							if int32(3) < v200 {
								if v198 == int32(0) {
									v219 = F_unlink(m, v7+int32(80))
									mBase = m.M
									F_stopSaving(m, int32(0))
									mBase = m.M
									v222 = m.ExcPending
									if v222 != 0 {
										return int32(0)
									} else {
										v224 = int32(-1)
										m.G0 = v7 + int32(416)
										return v224
									}
								} else {
									v215 = F_fclose(m, v198)
									mBase = m.M
									v216 = m.ExcPending
									if v216 != 0 {
										return int32(0)
									} else {
										v219 = F_unlink(m, v7+int32(80))
										mBase = m.M
										F_stopSaving(m, int32(0))
										mBase = m.M
										v222 = m.ExcPending
										if v222 != 0 {
											return int32(0)
										} else {
											v224 = int32(-1)
											m.G0 = v7 + int32(416)
											return v224
										}
									}
								}
							} else {
								v204 = *(*int32)(unsafe.Add(mBase, _c_F_rewriteAppendOnlyFile[3]))
								v205 = F___strerror_l(m, v204, v204)
								mBase = m.M
								*(*int32)(unsafe.Add(mBase, uint32(v7)+16)) = v205
								F__serverLog(m, int32(3), int32(_a_F_rewriteAppendOnlyFile_3), v7+int32(16))
								mBase = m.M
								v212 = m.ExcPending
								if v212 != 0 {
									return int32(0)
								} else {
									if v198 == int32(0) {
										v219 = F_unlink(m, v7+int32(80))
										mBase = m.M
										F_stopSaving(m, int32(0))
										mBase = m.M
										v222 = m.ExcPending
										if v222 != 0 {
											return int32(0)
										} else {
											v224 = int32(-1)
											m.G0 = v7 + int32(416)
											return v224
										}
									} else {
										v215 = F_fclose(m, v198)
										mBase = m.M
										v216 = m.ExcPending
										if v216 != 0 {
											return int32(0)
										} else {
											v219 = F_unlink(m, v7+int32(80))
											mBase = m.M
											F_stopSaving(m, int32(0))
											mBase = m.M
											v222 = m.ExcPending
											if v222 != 0 {
												return int32(0)
											} else {
												v224 = int32(-1)
												m.G0 = v7 + int32(416)
												return v224
											}
										}
									}
								}
							}
						} else {
							v103 = F_fflush(m, v24)
							mBase = m.M
							v104 = m.ExcPending
							if v104 != 0 {
								return int32(0)
							} else {
								if v103 != 0 {
									v198 = v24
									v200 = *(*int32)(unsafe.Add(mBase, _c_F_rewriteAppendOnlyFile[2]))
									if int32(3) < v200 {
										if v198 == int32(0) {
											v219 = F_unlink(m, v7+int32(80))
											mBase = m.M
											F_stopSaving(m, int32(0))
											mBase = m.M
											v222 = m.ExcPending
											if v222 != 0 {
												return int32(0)
											} else {
												v224 = int32(-1)
												m.G0 = v7 + int32(416)
												return v224
											}
										} else {
											v215 = F_fclose(m, v198)
											mBase = m.M
											v216 = m.ExcPending
											if v216 != 0 {
												return int32(0)
											} else {
												v219 = F_unlink(m, v7+int32(80))
												mBase = m.M
												F_stopSaving(m, int32(0))
												mBase = m.M
												v222 = m.ExcPending
												if v222 != 0 {
													return int32(0)
												} else {
													v224 = int32(-1)
													m.G0 = v7 + int32(416)
													return v224
												}
											}
										}
									} else {
										v204 = *(*int32)(unsafe.Add(mBase, _c_F_rewriteAppendOnlyFile[3]))
										v205 = F___strerror_l(m, v204, v204)
										mBase = m.M
										*(*int32)(unsafe.Add(mBase, uint32(v7)+16)) = v205
										F__serverLog(m, int32(3), int32(_a_F_rewriteAppendOnlyFile_3), v7+int32(16))
										mBase = m.M
										v212 = m.ExcPending
										if v212 != 0 {
											return int32(0)
										} else {
											if v198 == int32(0) {
												v219 = F_unlink(m, v7+int32(80))
												mBase = m.M
												F_stopSaving(m, int32(0))
												mBase = m.M
												v222 = m.ExcPending
												if v222 != 0 {
													return int32(0)
												} else {
													v224 = int32(-1)
													m.G0 = v7 + int32(416)
													return v224
												}
											} else {
												v215 = F_fclose(m, v198)
												mBase = m.M
												v216 = m.ExcPending
												if v216 != 0 {
													return int32(0)
												} else {
													v219 = F_unlink(m, v7+int32(80))
													mBase = m.M
													F_stopSaving(m, int32(0))
													mBase = m.M
													v222 = m.ExcPending
													if v222 != 0 {
														return int32(0)
													} else {
														v224 = int32(-1)
														m.G0 = v7 + int32(416)
														return v224
													}
												}
											}
										}
									}
								} else {
									v107 = *(*int32)(unsafe.Add(mBase, uint32(v24)+76))
									if int32(-1) < v107 {
										v111 = F___lockfile(m, v24)
										mBase = m.M
										v112 = *(*int32)(unsafe.Add(mBase, uint32(v24)+60))
										if v111 == int32(0) {
											v116 = v112
										} else {
											F___unlockfile(m, v24)
											mBase = m.M
											v116 = v112
										}
									} else {
										v110 = *(*int32)(unsafe.Add(mBase, uint32(v24)+60))
										v116 = v110
									}
									if int32(-1) < v116 {
										v124 = v116
									} else {
										v120 = F___errno_location(m)
										mBase = m.M
										*(*int32)(unsafe.Add(mBase, uint32(v120))) = int32(8)
										v124 = int32(-1)
									}
									v125 = F_fsync(m, v124)
									mBase = m.M
									if v125 != 0 {
										v198 = v24
										v200 = *(*int32)(unsafe.Add(mBase, _c_F_rewriteAppendOnlyFile[2]))
										if int32(3) < v200 {
											if v198 == int32(0) {
												v219 = F_unlink(m, v7+int32(80))
												mBase = m.M
												F_stopSaving(m, int32(0))
												mBase = m.M
												v222 = m.ExcPending
												if v222 != 0 {
													return int32(0)
												} else {
													v224 = int32(-1)
													m.G0 = v7 + int32(416)
													return v224
												}
											} else {
												v215 = F_fclose(m, v198)
												mBase = m.M
												v216 = m.ExcPending
												if v216 != 0 {
													return int32(0)
												} else {
													v219 = F_unlink(m, v7+int32(80))
													mBase = m.M
													F_stopSaving(m, int32(0))
													mBase = m.M
													v222 = m.ExcPending
													if v222 != 0 {
														return int32(0)
													} else {
														v224 = int32(-1)
														m.G0 = v7 + int32(416)
														return v224
													}
												}
											}
										} else {
											v204 = *(*int32)(unsafe.Add(mBase, _c_F_rewriteAppendOnlyFile[3]))
											v205 = F___strerror_l(m, v204, v204)
											mBase = m.M
											*(*int32)(unsafe.Add(mBase, uint32(v7)+16)) = v205
											F__serverLog(m, int32(3), int32(_a_F_rewriteAppendOnlyFile_3), v7+int32(16))
											mBase = m.M
											v212 = m.ExcPending
											if v212 != 0 {
												return int32(0)
											} else {
												if v198 == int32(0) {
													v219 = F_unlink(m, v7+int32(80))
													mBase = m.M
													F_stopSaving(m, int32(0))
													mBase = m.M
													v222 = m.ExcPending
													if v222 != 0 {
														return int32(0)
													} else {
														v224 = int32(-1)
														m.G0 = v7 + int32(416)
														return v224
													}
												} else {
													v215 = F_fclose(m, v198)
													mBase = m.M
													v216 = m.ExcPending
													if v216 != 0 {
														return int32(0)
													} else {
														v219 = F_unlink(m, v7+int32(80))
														mBase = m.M
														F_stopSaving(m, int32(0))
														mBase = m.M
														v222 = m.ExcPending
														if v222 != 0 {
															return int32(0)
														} else {
															v224 = int32(-1)
															m.G0 = v7 + int32(416)
															return v224
														}
													}
												}
											}
										}
									} else {
										v129 = *(*int32)(unsafe.Add(mBase, uint32(v24)+76))
										if int32(-1) < v129 {
											v133 = F___lockfile(m, v24)
											mBase = m.M
											v134 = *(*int32)(unsafe.Add(mBase, uint32(v24)+60))
											if v133 == int32(0) {
												v138 = v134
											} else {
												F___unlockfile(m, v24)
												mBase = m.M
												v138 = v134
											}
										} else {
											v132 = *(*int32)(unsafe.Add(mBase, uint32(v24)+60))
											v138 = v132
										}
										if int32(-1) < v138 {
										} else {
											v142 = F___errno_location(m)
											mBase = m.M
											*(*int32)(unsafe.Add(mBase, uint32(v142))) = int32(8)
										}
										v166 = F_fclose(m, v24)
										mBase = m.M
										v167 = m.ExcPending
										if v167 != 0 {
											return int32(0)
										} else {
											if v166 != 0 {
												v198 = int32(0)
												v200 = *(*int32)(unsafe.Add(mBase, _c_F_rewriteAppendOnlyFile[2]))
												if int32(3) < v200 {
													if v198 == int32(0) {
														v219 = F_unlink(m, v7+int32(80))
														mBase = m.M
														F_stopSaving(m, int32(0))
														mBase = m.M
														v222 = m.ExcPending
														if v222 != 0 {
															return int32(0)
														} else {
															v224 = int32(-1)
															m.G0 = v7 + int32(416)
															return v224
														}
													} else {
														v215 = F_fclose(m, v198)
														mBase = m.M
														v216 = m.ExcPending
														if v216 != 0 {
															return int32(0)
														} else {
															v219 = F_unlink(m, v7+int32(80))
															mBase = m.M
															F_stopSaving(m, int32(0))
															mBase = m.M
															v222 = m.ExcPending
															if v222 != 0 {
																return int32(0)
															} else {
																v224 = int32(-1)
																m.G0 = v7 + int32(416)
																return v224
															}
														}
													}
												} else {
													v204 = *(*int32)(unsafe.Add(mBase, _c_F_rewriteAppendOnlyFile[3]))
													v205 = F___strerror_l(m, v204, v204)
													mBase = m.M
													*(*int32)(unsafe.Add(mBase, uint32(v7)+16)) = v205
													F__serverLog(m, int32(3), int32(_a_F_rewriteAppendOnlyFile_3), v7+int32(16))
													mBase = m.M
													v212 = m.ExcPending
													if v212 != 0 {
														return int32(0)
													} else {
														if v198 == int32(0) {
															v219 = F_unlink(m, v7+int32(80))
															mBase = m.M
															F_stopSaving(m, int32(0))
															mBase = m.M
															v222 = m.ExcPending
															if v222 != 0 {
																return int32(0)
															} else {
																v224 = int32(-1)
																m.G0 = v7 + int32(416)
																return v224
															}
														} else {
															v215 = F_fclose(m, v198)
															mBase = m.M
															v216 = m.ExcPending
															if v216 != 0 {
																return int32(0)
															} else {
																v219 = F_unlink(m, v7+int32(80))
																mBase = m.M
																F_stopSaving(m, int32(0))
																mBase = m.M
																v222 = m.ExcPending
																if v222 != 0 {
																	return int32(0)
																} else {
																	v224 = int32(-1)
																	m.G0 = v7 + int32(416)
																	return v224
																}
															}
														}
													}
												}
											} else {
												v170 = F_rename(m, v7+int32(80), l0)
												mBase = m.M
												if v170 != int32(-1) {
													F_stopSaving(m, int32(1))
													mBase = m.M
													v196 = m.ExcPending
													if v196 != 0 {
														return int32(0)
													} else {
														v224 = int32(0)
														m.G0 = v7 + int32(416)
														return v224
													}
												} else {
													v174 = *(*int32)(unsafe.Add(mBase, _c_F_rewriteAppendOnlyFile[2]))
													if int32(3) < v174 {
														v189 = F_unlink(m, v7+int32(80))
														mBase = m.M
														F_stopSaving(m, int32(0))
														mBase = m.M
														v192 = m.ExcPending
														if v192 != 0 {
															return int32(0)
														} else {
															v224 = int32(-1)
															m.G0 = v7 + int32(416)
															return v224
														}
													} else {
														v178 = *(*int32)(unsafe.Add(mBase, _c_F_rewriteAppendOnlyFile[3]))
														v179 = F___strerror_l(m, v178, v178)
														mBase = m.M
														*(*int32)(unsafe.Add(mBase, uint32(v7)+32)) = v179
														F__serverLog(m, int32(3), int32(_a_F_rewriteAppendOnlyFile_4), v7+int32(32))
														mBase = m.M
														v186 = m.ExcPending
														if v186 != 0 {
															return int32(0)
														} else {
															v189 = F_unlink(m, v7+int32(80))
															mBase = m.M
															F_stopSaving(m, int32(0))
															mBase = m.M
															v192 = m.ExcPending
															if v192 != 0 {
																return int32(0)
															} else {
																v224 = int32(-1)
																m.G0 = v7 + int32(416)
																return v224
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
					v82 = int32(0)
					v90 = F_rdbSaveRio(m, v82, int32(80), v7+int32(336), v7+int32(76), int32(1), v82)
					mBase = m.M
					v91 = m.ExcPending
					if v91 != 0 {
						return int32(0)
					} else {
						if v90 != int32(-1) {
							v103 = F_fflush(m, v24)
							mBase = m.M
							v104 = m.ExcPending
							if v104 != 0 {
								return int32(0)
							} else {
								if v103 != 0 {
									v198 = v24
									v200 = *(*int32)(unsafe.Add(mBase, _c_F_rewriteAppendOnlyFile[2]))
									if int32(3) < v200 {
										if v198 == int32(0) {
											v219 = F_unlink(m, v7+int32(80))
											mBase = m.M
											F_stopSaving(m, int32(0))
											mBase = m.M
											v222 = m.ExcPending
											if v222 != 0 {
												return int32(0)
											} else {
												v224 = int32(-1)
												m.G0 = v7 + int32(416)
												return v224
											}
										} else {
											v215 = F_fclose(m, v198)
											mBase = m.M
											v216 = m.ExcPending
											if v216 != 0 {
												return int32(0)
											} else {
												v219 = F_unlink(m, v7+int32(80))
												mBase = m.M
												F_stopSaving(m, int32(0))
												mBase = m.M
												v222 = m.ExcPending
												if v222 != 0 {
													return int32(0)
												} else {
													v224 = int32(-1)
													m.G0 = v7 + int32(416)
													return v224
												}
											}
										}
									} else {
										v204 = *(*int32)(unsafe.Add(mBase, _c_F_rewriteAppendOnlyFile[3]))
										v205 = F___strerror_l(m, v204, v204)
										mBase = m.M
										*(*int32)(unsafe.Add(mBase, uint32(v7)+16)) = v205
										F__serverLog(m, int32(3), int32(_a_F_rewriteAppendOnlyFile_3), v7+int32(16))
										mBase = m.M
										v212 = m.ExcPending
										if v212 != 0 {
											return int32(0)
										} else {
											if v198 == int32(0) {
												v219 = F_unlink(m, v7+int32(80))
												mBase = m.M
												F_stopSaving(m, int32(0))
												mBase = m.M
												v222 = m.ExcPending
												if v222 != 0 {
													return int32(0)
												} else {
													v224 = int32(-1)
													m.G0 = v7 + int32(416)
													return v224
												}
											} else {
												v215 = F_fclose(m, v198)
												mBase = m.M
												v216 = m.ExcPending
												if v216 != 0 {
													return int32(0)
												} else {
													v219 = F_unlink(m, v7+int32(80))
													mBase = m.M
													F_stopSaving(m, int32(0))
													mBase = m.M
													v222 = m.ExcPending
													if v222 != 0 {
														return int32(0)
													} else {
														v224 = int32(-1)
														m.G0 = v7 + int32(416)
														return v224
													}
												}
											}
										}
									}
								} else {
									v107 = *(*int32)(unsafe.Add(mBase, uint32(v24)+76))
									if int32(-1) < v107 {
										v111 = F___lockfile(m, v24)
										mBase = m.M
										v112 = *(*int32)(unsafe.Add(mBase, uint32(v24)+60))
										if v111 == int32(0) {
											v116 = v112
										} else {
											F___unlockfile(m, v24)
											mBase = m.M
											v116 = v112
										}
									} else {
										v110 = *(*int32)(unsafe.Add(mBase, uint32(v24)+60))
										v116 = v110
									}
									if int32(-1) < v116 {
										v124 = v116
									} else {
										v120 = F___errno_location(m)
										mBase = m.M
										*(*int32)(unsafe.Add(mBase, uint32(v120))) = int32(8)
										v124 = int32(-1)
									}
									v125 = F_fsync(m, v124)
									mBase = m.M
									if v125 != 0 {
										v198 = v24
										v200 = *(*int32)(unsafe.Add(mBase, _c_F_rewriteAppendOnlyFile[2]))
										if int32(3) < v200 {
											if v198 == int32(0) {
												v219 = F_unlink(m, v7+int32(80))
												mBase = m.M
												F_stopSaving(m, int32(0))
												mBase = m.M
												v222 = m.ExcPending
												if v222 != 0 {
													return int32(0)
												} else {
													v224 = int32(-1)
													m.G0 = v7 + int32(416)
													return v224
												}
											} else {
												v215 = F_fclose(m, v198)
												mBase = m.M
												v216 = m.ExcPending
												if v216 != 0 {
													return int32(0)
												} else {
													v219 = F_unlink(m, v7+int32(80))
													mBase = m.M
													F_stopSaving(m, int32(0))
													mBase = m.M
													v222 = m.ExcPending
													if v222 != 0 {
														return int32(0)
													} else {
														v224 = int32(-1)
														m.G0 = v7 + int32(416)
														return v224
													}
												}
											}
										} else {
											v204 = *(*int32)(unsafe.Add(mBase, _c_F_rewriteAppendOnlyFile[3]))
											v205 = F___strerror_l(m, v204, v204)
											mBase = m.M
											*(*int32)(unsafe.Add(mBase, uint32(v7)+16)) = v205
											F__serverLog(m, int32(3), int32(_a_F_rewriteAppendOnlyFile_3), v7+int32(16))
											mBase = m.M
											v212 = m.ExcPending
											if v212 != 0 {
												return int32(0)
											} else {
												if v198 == int32(0) {
													v219 = F_unlink(m, v7+int32(80))
													mBase = m.M
													F_stopSaving(m, int32(0))
													mBase = m.M
													v222 = m.ExcPending
													if v222 != 0 {
														return int32(0)
													} else {
														v224 = int32(-1)
														m.G0 = v7 + int32(416)
														return v224
													}
												} else {
													v215 = F_fclose(m, v198)
													mBase = m.M
													v216 = m.ExcPending
													if v216 != 0 {
														return int32(0)
													} else {
														v219 = F_unlink(m, v7+int32(80))
														mBase = m.M
														F_stopSaving(m, int32(0))
														mBase = m.M
														v222 = m.ExcPending
														if v222 != 0 {
															return int32(0)
														} else {
															v224 = int32(-1)
															m.G0 = v7 + int32(416)
															return v224
														}
													}
												}
											}
										}
									} else {
										v129 = *(*int32)(unsafe.Add(mBase, uint32(v24)+76))
										if int32(-1) < v129 {
											v133 = F___lockfile(m, v24)
											mBase = m.M
											v134 = *(*int32)(unsafe.Add(mBase, uint32(v24)+60))
											if v133 == int32(0) {
												v138 = v134
											} else {
												F___unlockfile(m, v24)
												mBase = m.M
												v138 = v134
											}
										} else {
											v132 = *(*int32)(unsafe.Add(mBase, uint32(v24)+60))
											v138 = v132
										}
										if int32(-1) < v138 {
										} else {
											v142 = F___errno_location(m)
											mBase = m.M
											*(*int32)(unsafe.Add(mBase, uint32(v142))) = int32(8)
										}
										v166 = F_fclose(m, v24)
										mBase = m.M
										v167 = m.ExcPending
										if v167 != 0 {
											return int32(0)
										} else {
											if v166 != 0 {
												v198 = int32(0)
												v200 = *(*int32)(unsafe.Add(mBase, _c_F_rewriteAppendOnlyFile[2]))
												if int32(3) < v200 {
													if v198 == int32(0) {
														v219 = F_unlink(m, v7+int32(80))
														mBase = m.M
														F_stopSaving(m, int32(0))
														mBase = m.M
														v222 = m.ExcPending
														if v222 != 0 {
															return int32(0)
														} else {
															v224 = int32(-1)
															m.G0 = v7 + int32(416)
															return v224
														}
													} else {
														v215 = F_fclose(m, v198)
														mBase = m.M
														v216 = m.ExcPending
														if v216 != 0 {
															return int32(0)
														} else {
															v219 = F_unlink(m, v7+int32(80))
															mBase = m.M
															F_stopSaving(m, int32(0))
															mBase = m.M
															v222 = m.ExcPending
															if v222 != 0 {
																return int32(0)
															} else {
																v224 = int32(-1)
																m.G0 = v7 + int32(416)
																return v224
															}
														}
													}
												} else {
													v204 = *(*int32)(unsafe.Add(mBase, _c_F_rewriteAppendOnlyFile[3]))
													v205 = F___strerror_l(m, v204, v204)
													mBase = m.M
													*(*int32)(unsafe.Add(mBase, uint32(v7)+16)) = v205
													F__serverLog(m, int32(3), int32(_a_F_rewriteAppendOnlyFile_3), v7+int32(16))
													mBase = m.M
													v212 = m.ExcPending
													if v212 != 0 {
														return int32(0)
													} else {
														if v198 == int32(0) {
															v219 = F_unlink(m, v7+int32(80))
															mBase = m.M
															F_stopSaving(m, int32(0))
															mBase = m.M
															v222 = m.ExcPending
															if v222 != 0 {
																return int32(0)
															} else {
																v224 = int32(-1)
																m.G0 = v7 + int32(416)
																return v224
															}
														} else {
															v215 = F_fclose(m, v198)
															mBase = m.M
															v216 = m.ExcPending
															if v216 != 0 {
																return int32(0)
															} else {
																v219 = F_unlink(m, v7+int32(80))
																mBase = m.M
																F_stopSaving(m, int32(0))
																mBase = m.M
																v222 = m.ExcPending
																if v222 != 0 {
																	return int32(0)
																} else {
																	v224 = int32(-1)
																	m.G0 = v7 + int32(416)
																	return v224
																}
															}
														}
													}
												}
											} else {
												v170 = F_rename(m, v7+int32(80), l0)
												mBase = m.M
												if v170 != int32(-1) {
													F_stopSaving(m, int32(1))
													mBase = m.M
													v196 = m.ExcPending
													if v196 != 0 {
														return int32(0)
													} else {
														v224 = int32(0)
														m.G0 = v7 + int32(416)
														return v224
													}
												} else {
													v174 = *(*int32)(unsafe.Add(mBase, _c_F_rewriteAppendOnlyFile[2]))
													if int32(3) < v174 {
														v189 = F_unlink(m, v7+int32(80))
														mBase = m.M
														F_stopSaving(m, int32(0))
														mBase = m.M
														v192 = m.ExcPending
														if v192 != 0 {
															return int32(0)
														} else {
															v224 = int32(-1)
															m.G0 = v7 + int32(416)
															return v224
														}
													} else {
														v178 = *(*int32)(unsafe.Add(mBase, _c_F_rewriteAppendOnlyFile[3]))
														v179 = F___strerror_l(m, v178, v178)
														mBase = m.M
														*(*int32)(unsafe.Add(mBase, uint32(v7)+32)) = v179
														F__serverLog(m, int32(3), int32(_a_F_rewriteAppendOnlyFile_4), v7+int32(32))
														mBase = m.M
														v186 = m.ExcPending
														if v186 != 0 {
															return int32(0)
														} else {
															v189 = F_unlink(m, v7+int32(80))
															mBase = m.M
															F_stopSaving(m, int32(0))
															mBase = m.M
															v192 = m.ExcPending
															if v192 != 0 {
																return int32(0)
															} else {
																v224 = int32(-1)
																m.G0 = v7 + int32(416)
																return v224
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
							v95 = *(*int32)(unsafe.Add(mBase, uint32(v7)+76))
							*(*int32)(unsafe.Add(mBase, _c_F_rewriteAppendOnlyFile[3])) = v95
							v198 = v24
							v200 = *(*int32)(unsafe.Add(mBase, _c_F_rewriteAppendOnlyFile[2]))
							if int32(3) < v200 {
								if v198 == int32(0) {
									v219 = F_unlink(m, v7+int32(80))
									mBase = m.M
									F_stopSaving(m, int32(0))
									mBase = m.M
									v222 = m.ExcPending
									if v222 != 0 {
										return int32(0)
									} else {
										v224 = int32(-1)
										m.G0 = v7 + int32(416)
										return v224
									}
								} else {
									v215 = F_fclose(m, v198)
									mBase = m.M
									v216 = m.ExcPending
									if v216 != 0 {
										return int32(0)
									} else {
										v219 = F_unlink(m, v7+int32(80))
										mBase = m.M
										F_stopSaving(m, int32(0))
										mBase = m.M
										v222 = m.ExcPending
										if v222 != 0 {
											return int32(0)
										} else {
											v224 = int32(-1)
											m.G0 = v7 + int32(416)
											return v224
										}
									}
								}
							} else {
								v204 = *(*int32)(unsafe.Add(mBase, _c_F_rewriteAppendOnlyFile[3]))
								v205 = F___strerror_l(m, v204, v204)
								mBase = m.M
								*(*int32)(unsafe.Add(mBase, uint32(v7)+16)) = v205
								F__serverLog(m, int32(3), int32(_a_F_rewriteAppendOnlyFile_3), v7+int32(16))
								mBase = m.M
								v212 = m.ExcPending
								if v212 != 0 {
									return int32(0)
								} else {
									if v198 == int32(0) {
										v219 = F_unlink(m, v7+int32(80))
										mBase = m.M
										F_stopSaving(m, int32(0))
										mBase = m.M
										v222 = m.ExcPending
										if v222 != 0 {
											return int32(0)
										} else {
											v224 = int32(-1)
											m.G0 = v7 + int32(416)
											return v224
										}
									} else {
										v215 = F_fclose(m, v198)
										mBase = m.M
										v216 = m.ExcPending
										if v216 != 0 {
											return int32(0)
										} else {
											v219 = F_unlink(m, v7+int32(80))
											mBase = m.M
											F_stopSaving(m, int32(0))
											mBase = m.M
											v222 = m.ExcPending
											if v222 != 0 {
												return int32(0)
											} else {
												v224 = int32(-1)
												m.G0 = v7 + int32(416)
												return v224
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
			v25 = int32(-1)
			v27 = *(*int32)(unsafe.Add(mBase, _c_F_rewriteAppendOnlyFile[2]))
			if int32(3) < v27 {
				v224 = v25
				m.G0 = v7 + int32(416)
				return v224
			} else {
				v31 = *(*int32)(unsafe.Add(mBase, _c_F_rewriteAppendOnlyFile[3]))
				v32 = F___strerror_l(m, v31, v31)
				mBase = m.M
				*(*int32)(unsafe.Add(mBase, uint32(v7))) = v32
				F__serverLog(m, int32(3), int32(_a_F_rewriteAppendOnlyFile_5), v7)
				mBase = m.M
				v37 = m.ExcPending
				if v37 != 0 {
					return int32(0)
				} else {
					v224 = v25
					m.G0 = v7 + int32(416)
					return v224
				}
			}
		}
	}
}
