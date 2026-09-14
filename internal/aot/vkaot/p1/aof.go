package p1

import (
	base "github.com/shibukawa/vkmem/internal/aot/vkaot/base"
	"unsafe"
)

func F_aofDelTempIncrAofFile(m *base.Module) {
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
	var v15 int32
	_ = v15
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	v4 = m.G0
	v6 = v4 - int32(32)
	m.G0 = v6
	v8 = F_sdsempty(m)
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(v6)+16)) = int32(_a_F_aofDelTempIncrAofFile_0)
		*(*int32)(unsafe.Add(mBase, uint32(v6)+24)) = int32(_a_F_aofDelTempIncrAofFile_1)
		v15 = *(*int32)(unsafe.Add(mBase, _c_F_aofDelTempIncrAofFile[0]))
		*(*int32)(unsafe.Add(mBase, uint32(v6)+20)) = v15
		v20 = F_sdscatprintf(m, v8, int32(_a_F_aofDelTempIncrAofFile_2), v6+int32(16))
		mBase = m.M
		v21 = m.ExcPending
		if v21 != 0 {
			return
		} else {
			v23 = *(*int32)(unsafe.Add(mBase, _c_F_aofDelTempIncrAofFile[1]))
			v24 = F_makePath(m, v23, v20)
			mBase = m.M
			v25 = m.ExcPending
			if v25 != 0 {
				return
			} else {
				v27 = *(*int32)(unsafe.Add(mBase, _c_F_aofDelTempIncrAofFile[2]))
				if int32(2) < v27 {
					v35 = F_bg_unlink(m, v24)
					mBase = m.M
					v36 = m.ExcPending
					if v36 != 0 {
						return
					} else {
						F_sdsfree(m, v24)
						mBase = m.M
						v38 = m.ExcPending
						if v38 != 0 {
							return
						} else {
							F_sdsfree(m, v20)
							mBase = m.M
							v40 = m.ExcPending
							if v40 != 0 {
								return
							} else {
								m.G0 = v6 + int32(32)
								return
							}
						}
					}
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v6))) = v20
					F__serverLog(m, int32(2), int32(_a_F_aofDelTempIncrAofFile_3), v6)
					mBase = m.M
					v34 = m.ExcPending
					if v34 != 0 {
						return
					} else {
						v35 = F_bg_unlink(m, v24)
						mBase = m.M
						v36 = m.ExcPending
						if v36 != 0 {
							return
						} else {
							F_sdsfree(m, v24)
							mBase = m.M
							v38 = m.ExcPending
							if v38 != 0 {
								return
							} else {
								F_sdsfree(m, v20)
								mBase = m.M
								v40 = m.ExcPending
								if v40 != 0 {
									return
								} else {
									m.G0 = v6 + int32(32)
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
func F_aofListFree(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v16 int32
	_ = v16
	if l0 == int32(0) {
		F__serverAssert(m, int32(_a_F_aofListFree_0), int32(_a_F_aofListFree_1), int32(105))
		mBase = m.M
		v16 = m.ExcPending
		if v16 != 0 {
			return
		} else {
			F_abort(m)
			mBase = m.M
			base.Wasm_trap_unreachable()
			for {
			}
		}
	} else {
		v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		if v5 == int32(0) {
			F_valkey_free(m, l0)
			mBase = m.M
			v11 = m.ExcPending
			if v11 != 0 {
				return
			} else {
				return
			}
		} else {
			F_sdsfree(m, v5)
			mBase = m.M
			v9 = m.ExcPending
			if v9 != 0 {
				return
			} else {
				F_valkey_free(m, l0)
				mBase = m.M
				v11 = m.ExcPending
				if v11 != 0 {
					return
				} else {
					return
				}
			}
		}
	}
}
func F_aofManifestDup(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v12 int64
	_ = v12
	var v14 int64
	_ = v14
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v28 int64
	_ = v28
	var v30 int32
	_ = v30
	var v34 int32
	_ = v34
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
	var v51 int32
	_ = v51
	var v57 int32
	_ = v57
	var v63 int32
	_ = v63
	if l0 == int32(0) {
		F__serverAssert(m, int32(_a_F_aofManifestDup_0), int32(_a_F_aofManifestDup_1), int32(392))
		mBase = m.M
		v51 = m.ExcPending
		if v51 != 0 {
			return int32(0)
		} else {
			F_abort(m)
			mBase = m.M
			base.Wasm_trap_unreachable()
			for {
			}
		}
	} else {
		v8 = F_valkey_calloc(m, int32(40))
		mBase = m.M
		v11 = m.ExcPending
		if v11 != 0 {
			return int32(0)
		} else {
			v12 = *(*int64)(unsafe.Add(mBase, uint32(l0)+16))
			*(*int64)(unsafe.Add(mBase, uint32(v8)+16)) = v12
			v14 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
			*(*int64)(unsafe.Add(mBase, uint32(v8)+24)) = v14
			v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
			*(*int32)(unsafe.Add(mBase, uint32(v8)+32)) = v16
			v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			if v18 == int32(0) {
				v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
				v35 = F_listDup(m, v34)
				mBase = m.M
				v36 = m.ExcPending
				if v36 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = v35
					v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
					v39 = F_listDup(m, v38)
					mBase = m.M
					v40 = m.ExcPending
					if v40 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v8)+8)) = v39
						if v35 == int32(0) {
							F__serverAssert(m, int32(_a_F_aofManifestDup_2), int32(_a_F_aofManifestDup_1), int32(405))
							mBase = m.M
							v57 = m.ExcPending
							if v57 != 0 {
								return int32(0)
							} else {
								F_abort(m)
								mBase = m.M
								base.Wasm_trap_unreachable()
								for {
								}
							}
						} else {
							if v39 == int32(0) {
								F__serverAssert(m, int32(_a_F_aofManifestDup_3), int32(_a_F_aofManifestDup_1), int32(406))
								mBase = m.M
								v63 = m.ExcPending
								if v63 != 0 {
									return int32(0)
								} else {
									F_abort(m)
									mBase = m.M
									base.Wasm_trap_unreachable()
									for {
									}
								}
							} else {
								return v8
							}
						}
					}
				}
			} else {
				v22 = F_valkey_calloc(m, int32(24))
				mBase = m.M
				v23 = m.ExcPending
				if v23 != 0 {
					return int32(0)
				} else {
					v24 = *(*int32)(unsafe.Add(mBase, uint32(v18)))
					v25 = F_sdsdup(m, v24)
					mBase = m.M
					v26 = m.ExcPending
					if v26 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v22))) = v25
						v28 = *(*int64)(unsafe.Add(mBase, uint32(v18)+8))
						*(*int64)(unsafe.Add(mBase, uint32(v22)+8)) = v28
						v30 = *(*int32)(unsafe.Add(mBase, uint32(v18)+16))
						*(*int32)(unsafe.Add(mBase, uint32(v22)+16)) = v30
						*(*int32)(unsafe.Add(mBase, uint32(v8))) = v22
						v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
						v35 = F_listDup(m, v34)
						mBase = m.M
						v36 = m.ExcPending
						if v36 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = v35
							v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
							v39 = F_listDup(m, v38)
							mBase = m.M
							v40 = m.ExcPending
							if v40 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v8)+8)) = v39
								if v35 == int32(0) {
									F__serverAssert(m, int32(_a_F_aofManifestDup_2), int32(_a_F_aofManifestDup_1), int32(405))
									mBase = m.M
									v57 = m.ExcPending
									if v57 != 0 {
										return int32(0)
									} else {
										F_abort(m)
										mBase = m.M
										base.Wasm_trap_unreachable()
										for {
										}
									}
								} else {
									if v39 == int32(0) {
										F__serverAssert(m, int32(_a_F_aofManifestDup_3), int32(_a_F_aofManifestDup_1), int32(406))
										mBase = m.M
										v63 = m.ExcPending
										if v63 != 0 {
											return int32(0)
										} else {
											F_abort(m)
											mBase = m.M
											base.Wasm_trap_unreachable()
											for {
											}
										}
									} else {
										return v8
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
func F_aofOpenIfNeededOnServerStart(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v17 int32
	_ = v17
	var v21 int32
	_ = v21
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v51 int32
	_ = v51
	var v59 int32
	_ = v59
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v72 int32
	_ = v72
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v103 int32
	_ = v103
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
	var v120 int32
	_ = v120
	var v123 int32
	_ = v123
	var v125 int32
	_ = v125
	var v129 int32
	_ = v129
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v142 int32
	_ = v142
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v155 int32
	_ = v155
	var v162 int64
	_ = v162
	var v163 int32
	_ = v163
	var v165 int32
	_ = v165
	var v168 int32
	_ = v168
	var v175 int32
	_ = v175
	var v179 int32
	_ = v179
	var v193 int32
	_ = v193
	var v199 int32
	_ = v199
	v8 = m.G0
	v10 = v8 - int32(80)
	m.G0 = v10
	v13 = *(*int32)(unsafe.Add(mBase, _c_F_aofOpenIfNeededOnServerStart[0]))
	if v13 != int32(1) {
		m.G0 = v10 + int32(80)
		return
	} else {
		v17 = *(*int32)(unsafe.Add(mBase, _c_F_aofOpenIfNeededOnServerStart[1]))
		if v17 == int32(0) {
			F__serverAssert(m, int32(_a_F_aofOpenIfNeededOnServerStart_0), int32(_a_F_aofOpenIfNeededOnServerStart_1), int32(705))
			mBase = m.M
			v193 = m.ExcPending
			if v193 != 0 {
				return
			} else {
				F_abort(m)
				mBase = m.M
				base.Wasm_trap_unreachable()
				for {
				}
			}
		} else {
			v21 = *(*int32)(unsafe.Add(mBase, _c_F_aofOpenIfNeededOnServerStart[2]))
			if v21 != int32(-1) {
				F__serverAssert(m, int32(_a_F_aofOpenIfNeededOnServerStart_2), int32(_a_F_aofOpenIfNeededOnServerStart_1), int32(706))
				mBase = m.M
				v199 = m.ExcPending
				if v199 != 0 {
					return
				} else {
					F_abort(m)
					mBase = m.M
					base.Wasm_trap_unreachable()
					for {
					}
				}
			} else {
				v25 = *(*int32)(unsafe.Add(mBase, _c_F_aofOpenIfNeededOnServerStart[3]))
				v26 = int32(0)
				v29 = m.G0
				v31 = v29 - int32(96)
				m.G0 = v31
				v35 = F_mkdir(m, v25, int32(493))
				mBase = m.M
				if v35 == v26 {
					v51 = v26
				} else {
					v38 = F___errno_location(m)
					mBase = m.M
					v39 = *(*int32)(unsafe.Add(mBase, uint32(v38)))
					if v39 != int32(20) {
						v51 = int32(-1)
					} else {
						v42 = F_stat(m, v25, v31)
						mBase = m.M
						if v42 != 0 {
							*(*int32)(unsafe.Add(mBase, uint32(v38))) = int32(54)
							v51 = int32(-1)
						} else {
							v43 = *(*int32)(unsafe.Add(mBase, uint32(v31)+4))
							if v43&int32(61440) == int32(16384) {
								v51 = v26
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v38))) = int32(54)
								v51 = int32(-1)
							}
						}
					}
				}
				m.G0 = v31 + int32(96)
				if v51 != int32(-1) {
					v77 = *(*int32)(unsafe.Add(mBase, _c_F_aofOpenIfNeededOnServerStart[1]))
					v78 = *(*int32)(unsafe.Add(mBase, uint32(v77)+4))
					v79 = *(*int32)(unsafe.Add(mBase, uint32(v78)+20))
					v80 = *(*int32)(unsafe.Add(mBase, uint32(v77)))
					if v80 != 0 {
						v107 = *(*int32)(unsafe.Add(mBase, _c_F_aofOpenIfNeededOnServerStart[1]))
						v108 = F_getLastIncrAofName(m, v107)
						mBase = m.M
						v109 = m.ExcPending
						if v109 != 0 {
							return
						} else {
							v111 = *(*int32)(unsafe.Add(mBase, _c_F_aofOpenIfNeededOnServerStart[3]))
							v112 = F_makePath(m, v111, v108)
							mBase = m.M
							v113 = m.ExcPending
							if v113 != 0 {
								return
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v10)+48)) = int32(438)
								v120 = F_open(m, v112, int32(1089), v10+int32(48))
								mBase = m.M
								*(*int32)(unsafe.Add(mBase, _c_F_aofOpenIfNeededOnServerStart[2])) = v120
								F_sdsfree(m, v112)
								mBase = m.M
								v123 = m.ExcPending
								if v123 != 0 {
									return
								} else {
									v125 = *(*int32)(unsafe.Add(mBase, _c_F_aofOpenIfNeededOnServerStart[2]))
									if v125 != int32(-1) {
										v146 = *(*int32)(unsafe.Add(mBase, _c_F_aofOpenIfNeededOnServerStart[1]))
										v147 = *(*int32)(unsafe.Add(mBase, uint32(v146)+32))
										if v147 == int32(0) {
											v162 = F_getAppendOnlyFileSize(m, v108, int32(0))
											mBase = m.M
											v163 = m.ExcPending
											if v163 != 0 {
												return
											} else {
												*(*int64)(unsafe.Add(mBase, _c_F_aofOpenIfNeededOnServerStart[4])) = v162
												v165 = int32(_a_F_aofOpenIfNeededOnServerStart_3)
												*(*int64)(unsafe.Add(mBase, _c_F_aofOpenIfNeededOnServerStart[5])) = v162
												v168 = *(*int32)(unsafe.Add(mBase, _c_F_aofOpenIfNeededOnServerStart[6]))
												if int32(2) < v168 {
													m.G0 = v10 + int32(80)
													return
												} else {
													*(*int32)(unsafe.Add(mBase, uint32(v10)+32)) = v108
													if v79 != 0 {
														v175 = int32(_a_F_aofOpenIfNeededOnServerStart_4)
													} else {
														v175 = int32(_a_F_aofOpenIfNeededOnServerStart_5)
													}
													F__serverLog(m, int32(2), v175, v10+int32(32))
													mBase = m.M
													v179 = m.ExcPending
													if v179 != 0 {
														return
													} else {
														m.G0 = v10 + int32(80)
														return
													}
												}
											}
										} else {
											v150 = F_getAofManifestAsString(m, v146)
											mBase = m.M
											v151 = m.ExcPending
											if v151 != 0 {
												return
											} else {
												v152 = F_writeAofManifestFile(m, v150)
												mBase = m.M
												v153 = m.ExcPending
												if v153 != 0 {
													return
												} else {
													F_sdsfree(m, v150)
													mBase = m.M
													v155 = m.ExcPending
													if v155 != 0 {
														return
													} else {
														if v152 != 0 {
															m.Env.Exit(m, int32(1))
															mBase = m.M
															base.Wasm_trap_unreachable()
															for {
															}
														} else {
															*(*int32)(unsafe.Add(mBase, uint32(v146)+32)) = int32(0)
															v162 = F_getAppendOnlyFileSize(m, v108, int32(0))
															mBase = m.M
															v163 = m.ExcPending
															if v163 != 0 {
																return
															} else {
																*(*int64)(unsafe.Add(mBase, _c_F_aofOpenIfNeededOnServerStart[4])) = v162
																v165 = int32(_a_F_aofOpenIfNeededOnServerStart_3)
																*(*int64)(unsafe.Add(mBase, _c_F_aofOpenIfNeededOnServerStart[5])) = v162
																v168 = *(*int32)(unsafe.Add(mBase, _c_F_aofOpenIfNeededOnServerStart[6]))
																if int32(2) < v168 {
																	m.G0 = v10 + int32(80)
																	return
																} else {
																	*(*int32)(unsafe.Add(mBase, uint32(v10)+32)) = v108
																	if v79 != 0 {
																		v175 = int32(_a_F_aofOpenIfNeededOnServerStart_4)
																	} else {
																		v175 = int32(_a_F_aofOpenIfNeededOnServerStart_5)
																	}
																	F__serverLog(m, int32(2), v175, v10+int32(32))
																	mBase = m.M
																	v179 = m.ExcPending
																	if v179 != 0 {
																		return
																	} else {
																		m.G0 = v10 + int32(80)
																		return
																	}
																}
															}
														}
													}
												}
											}
										}
									} else {
										v129 = *(*int32)(unsafe.Add(mBase, _c_F_aofOpenIfNeededOnServerStart[6]))
										if int32(3) < v129 {
											m.Env.Exit(m, int32(1))
											mBase = m.M
											base.Wasm_trap_unreachable()
											for {
											}
										} else {
											v133 = *(*int32)(unsafe.Add(mBase, _c_F_aofOpenIfNeededOnServerStart[7]))
											v134 = F___strerror_l(m, v133, v133)
											mBase = m.M
											*(*int32)(unsafe.Add(mBase, uint32(v10)+20)) = v134
											*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = v108
											F__serverLog(m, int32(3), int32(_a_F_aofOpenIfNeededOnServerStart_6), v10+int32(16))
											mBase = m.M
											v142 = m.ExcPending
											if v142 != 0 {
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
						if v79 != 0 {
							v107 = *(*int32)(unsafe.Add(mBase, _c_F_aofOpenIfNeededOnServerStart[1]))
							v108 = F_getLastIncrAofName(m, v107)
							mBase = m.M
							v109 = m.ExcPending
							if v109 != 0 {
								return
							} else {
								v111 = *(*int32)(unsafe.Add(mBase, _c_F_aofOpenIfNeededOnServerStart[3]))
								v112 = F_makePath(m, v111, v108)
								mBase = m.M
								v113 = m.ExcPending
								if v113 != 0 {
									return
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v10)+48)) = int32(438)
									v120 = F_open(m, v112, int32(1089), v10+int32(48))
									mBase = m.M
									*(*int32)(unsafe.Add(mBase, _c_F_aofOpenIfNeededOnServerStart[2])) = v120
									F_sdsfree(m, v112)
									mBase = m.M
									v123 = m.ExcPending
									if v123 != 0 {
										return
									} else {
										v125 = *(*int32)(unsafe.Add(mBase, _c_F_aofOpenIfNeededOnServerStart[2]))
										if v125 != int32(-1) {
											v146 = *(*int32)(unsafe.Add(mBase, _c_F_aofOpenIfNeededOnServerStart[1]))
											v147 = *(*int32)(unsafe.Add(mBase, uint32(v146)+32))
											if v147 == int32(0) {
												v162 = F_getAppendOnlyFileSize(m, v108, int32(0))
												mBase = m.M
												v163 = m.ExcPending
												if v163 != 0 {
													return
												} else {
													*(*int64)(unsafe.Add(mBase, _c_F_aofOpenIfNeededOnServerStart[4])) = v162
													v165 = int32(_a_F_aofOpenIfNeededOnServerStart_3)
													*(*int64)(unsafe.Add(mBase, _c_F_aofOpenIfNeededOnServerStart[5])) = v162
													v168 = *(*int32)(unsafe.Add(mBase, _c_F_aofOpenIfNeededOnServerStart[6]))
													if int32(2) < v168 {
														m.G0 = v10 + int32(80)
														return
													} else {
														*(*int32)(unsafe.Add(mBase, uint32(v10)+32)) = v108
														if v79 != 0 {
															v175 = int32(_a_F_aofOpenIfNeededOnServerStart_4)
														} else {
															v175 = int32(_a_F_aofOpenIfNeededOnServerStart_5)
														}
														F__serverLog(m, int32(2), v175, v10+int32(32))
														mBase = m.M
														v179 = m.ExcPending
														if v179 != 0 {
															return
														} else {
															m.G0 = v10 + int32(80)
															return
														}
													}
												}
											} else {
												v150 = F_getAofManifestAsString(m, v146)
												mBase = m.M
												v151 = m.ExcPending
												if v151 != 0 {
													return
												} else {
													v152 = F_writeAofManifestFile(m, v150)
													mBase = m.M
													v153 = m.ExcPending
													if v153 != 0 {
														return
													} else {
														F_sdsfree(m, v150)
														mBase = m.M
														v155 = m.ExcPending
														if v155 != 0 {
															return
														} else {
															if v152 != 0 {
																m.Env.Exit(m, int32(1))
																mBase = m.M
																base.Wasm_trap_unreachable()
																for {
																}
															} else {
																*(*int32)(unsafe.Add(mBase, uint32(v146)+32)) = int32(0)
																v162 = F_getAppendOnlyFileSize(m, v108, int32(0))
																mBase = m.M
																v163 = m.ExcPending
																if v163 != 0 {
																	return
																} else {
																	*(*int64)(unsafe.Add(mBase, _c_F_aofOpenIfNeededOnServerStart[4])) = v162
																	v165 = int32(_a_F_aofOpenIfNeededOnServerStart_3)
																	*(*int64)(unsafe.Add(mBase, _c_F_aofOpenIfNeededOnServerStart[5])) = v162
																	v168 = *(*int32)(unsafe.Add(mBase, _c_F_aofOpenIfNeededOnServerStart[6]))
																	if int32(2) < v168 {
																		m.G0 = v10 + int32(80)
																		return
																	} else {
																		*(*int32)(unsafe.Add(mBase, uint32(v10)+32)) = v108
																		if v79 != 0 {
																			v175 = int32(_a_F_aofOpenIfNeededOnServerStart_4)
																		} else {
																			v175 = int32(_a_F_aofOpenIfNeededOnServerStart_5)
																		}
																		F__serverLog(m, int32(2), v175, v10+int32(32))
																		mBase = m.M
																		v179 = m.ExcPending
																		if v179 != 0 {
																			return
																		} else {
																			m.G0 = v10 + int32(80)
																			return
																		}
																	}
																}
															}
														}
													}
												}
											}
										} else {
											v129 = *(*int32)(unsafe.Add(mBase, _c_F_aofOpenIfNeededOnServerStart[6]))
											if int32(3) < v129 {
												m.Env.Exit(m, int32(1))
												mBase = m.M
												base.Wasm_trap_unreachable()
												for {
												}
											} else {
												v133 = *(*int32)(unsafe.Add(mBase, _c_F_aofOpenIfNeededOnServerStart[7]))
												v134 = F___strerror_l(m, v133, v133)
												mBase = m.M
												*(*int32)(unsafe.Add(mBase, uint32(v10)+20)) = v134
												*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = v108
												F__serverLog(m, int32(3), int32(_a_F_aofOpenIfNeededOnServerStart_6), v10+int32(16))
												mBase = m.M
												v142 = m.ExcPending
												if v142 != 0 {
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
							v82 = *(*int32)(unsafe.Add(mBase, _c_F_aofOpenIfNeededOnServerStart[8]))
							v83 = F_getNewBaseFileNameAndMarkPreAsHistory(m, v77, v82)
							mBase = m.M
							v84 = m.ExcPending
							if v84 != 0 {
								return
							} else {
								v86 = *(*int32)(unsafe.Add(mBase, _c_F_aofOpenIfNeededOnServerStart[3]))
								v87 = F_makePath(m, v86, v83)
								mBase = m.M
								v88 = m.ExcPending
								if v88 != 0 {
									return
								} else {
									v89 = F_rewriteAppendOnlyFile(m, v87)
									mBase = m.M
									v90 = m.ExcPending
									if v90 != 0 {
										return
									} else {
										if v89 != 0 {
											m.Env.Exit(m, int32(1))
											mBase = m.M
											base.Wasm_trap_unreachable()
											for {
											}
										} else {
											F_sdsfree(m, v87)
											mBase = m.M
											v92 = m.ExcPending
											if v92 != 0 {
												return
											} else {
												v94 = *(*int32)(unsafe.Add(mBase, _c_F_aofOpenIfNeededOnServerStart[6]))
												if int32(2) < v94 {
													v107 = *(*int32)(unsafe.Add(mBase, _c_F_aofOpenIfNeededOnServerStart[1]))
													v108 = F_getLastIncrAofName(m, v107)
													mBase = m.M
													v109 = m.ExcPending
													if v109 != 0 {
														return
													} else {
														v111 = *(*int32)(unsafe.Add(mBase, _c_F_aofOpenIfNeededOnServerStart[3]))
														v112 = F_makePath(m, v111, v108)
														mBase = m.M
														v113 = m.ExcPending
														if v113 != 0 {
															return
														} else {
															*(*int32)(unsafe.Add(mBase, uint32(v10)+48)) = int32(438)
															v120 = F_open(m, v112, int32(1089), v10+int32(48))
															mBase = m.M
															*(*int32)(unsafe.Add(mBase, _c_F_aofOpenIfNeededOnServerStart[2])) = v120
															F_sdsfree(m, v112)
															mBase = m.M
															v123 = m.ExcPending
															if v123 != 0 {
																return
															} else {
																v125 = *(*int32)(unsafe.Add(mBase, _c_F_aofOpenIfNeededOnServerStart[2]))
																if v125 != int32(-1) {
																	v146 = *(*int32)(unsafe.Add(mBase, _c_F_aofOpenIfNeededOnServerStart[1]))
																	v147 = *(*int32)(unsafe.Add(mBase, uint32(v146)+32))
																	if v147 == int32(0) {
																		v162 = F_getAppendOnlyFileSize(m, v108, int32(0))
																		mBase = m.M
																		v163 = m.ExcPending
																		if v163 != 0 {
																			return
																		} else {
																			*(*int64)(unsafe.Add(mBase, _c_F_aofOpenIfNeededOnServerStart[4])) = v162
																			v165 = int32(_a_F_aofOpenIfNeededOnServerStart_3)
																			*(*int64)(unsafe.Add(mBase, _c_F_aofOpenIfNeededOnServerStart[5])) = v162
																			v168 = *(*int32)(unsafe.Add(mBase, _c_F_aofOpenIfNeededOnServerStart[6]))
																			if int32(2) < v168 {
																				m.G0 = v10 + int32(80)
																				return
																			} else {
																				*(*int32)(unsafe.Add(mBase, uint32(v10)+32)) = v108
																				if v79 != 0 {
																					v175 = int32(_a_F_aofOpenIfNeededOnServerStart_4)
																				} else {
																					v175 = int32(_a_F_aofOpenIfNeededOnServerStart_5)
																				}
																				F__serverLog(m, int32(2), v175, v10+int32(32))
																				mBase = m.M
																				v179 = m.ExcPending
																				if v179 != 0 {
																					return
																				} else {
																					m.G0 = v10 + int32(80)
																					return
																				}
																			}
																		}
																	} else {
																		v150 = F_getAofManifestAsString(m, v146)
																		mBase = m.M
																		v151 = m.ExcPending
																		if v151 != 0 {
																			return
																		} else {
																			v152 = F_writeAofManifestFile(m, v150)
																			mBase = m.M
																			v153 = m.ExcPending
																			if v153 != 0 {
																				return
																			} else {
																				F_sdsfree(m, v150)
																				mBase = m.M
																				v155 = m.ExcPending
																				if v155 != 0 {
																					return
																				} else {
																					if v152 != 0 {
																						m.Env.Exit(m, int32(1))
																						mBase = m.M
																						base.Wasm_trap_unreachable()
																						for {
																						}
																					} else {
																						*(*int32)(unsafe.Add(mBase, uint32(v146)+32)) = int32(0)
																						v162 = F_getAppendOnlyFileSize(m, v108, int32(0))
																						mBase = m.M
																						v163 = m.ExcPending
																						if v163 != 0 {
																							return
																						} else {
																							*(*int64)(unsafe.Add(mBase, _c_F_aofOpenIfNeededOnServerStart[4])) = v162
																							v165 = int32(_a_F_aofOpenIfNeededOnServerStart_3)
																							*(*int64)(unsafe.Add(mBase, _c_F_aofOpenIfNeededOnServerStart[5])) = v162
																							v168 = *(*int32)(unsafe.Add(mBase, _c_F_aofOpenIfNeededOnServerStart[6]))
																							if int32(2) < v168 {
																								m.G0 = v10 + int32(80)
																								return
																							} else {
																								*(*int32)(unsafe.Add(mBase, uint32(v10)+32)) = v108
																								if v79 != 0 {
																									v175 = int32(_a_F_aofOpenIfNeededOnServerStart_4)
																								} else {
																									v175 = int32(_a_F_aofOpenIfNeededOnServerStart_5)
																								}
																								F__serverLog(m, int32(2), v175, v10+int32(32))
																								mBase = m.M
																								v179 = m.ExcPending
																								if v179 != 0 {
																									return
																								} else {
																									m.G0 = v10 + int32(80)
																									return
																								}
																							}
																						}
																					}
																				}
																			}
																		}
																	}
																} else {
																	v129 = *(*int32)(unsafe.Add(mBase, _c_F_aofOpenIfNeededOnServerStart[6]))
																	if int32(3) < v129 {
																		m.Env.Exit(m, int32(1))
																		mBase = m.M
																		base.Wasm_trap_unreachable()
																		for {
																		}
																	} else {
																		v133 = *(*int32)(unsafe.Add(mBase, _c_F_aofOpenIfNeededOnServerStart[7]))
																		v134 = F___strerror_l(m, v133, v133)
																		mBase = m.M
																		*(*int32)(unsafe.Add(mBase, uint32(v10)+20)) = v134
																		*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = v108
																		F__serverLog(m, int32(3), int32(_a_F_aofOpenIfNeededOnServerStart_6), v10+int32(16))
																		mBase = m.M
																		v142 = m.ExcPending
																		if v142 != 0 {
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
													*(*int32)(unsafe.Add(mBase, uint32(v10)+64)) = v83
													F__serverLog(m, int32(2), int32(_a_F_aofOpenIfNeededOnServerStart_7), v10+int32(64))
													mBase = m.M
													v103 = m.ExcPending
													if v103 != 0 {
														return
													} else {
														v107 = *(*int32)(unsafe.Add(mBase, _c_F_aofOpenIfNeededOnServerStart[1]))
														v108 = F_getLastIncrAofName(m, v107)
														mBase = m.M
														v109 = m.ExcPending
														if v109 != 0 {
															return
														} else {
															v111 = *(*int32)(unsafe.Add(mBase, _c_F_aofOpenIfNeededOnServerStart[3]))
															v112 = F_makePath(m, v111, v108)
															mBase = m.M
															v113 = m.ExcPending
															if v113 != 0 {
																return
															} else {
																*(*int32)(unsafe.Add(mBase, uint32(v10)+48)) = int32(438)
																v120 = F_open(m, v112, int32(1089), v10+int32(48))
																mBase = m.M
																*(*int32)(unsafe.Add(mBase, _c_F_aofOpenIfNeededOnServerStart[2])) = v120
																F_sdsfree(m, v112)
																mBase = m.M
																v123 = m.ExcPending
																if v123 != 0 {
																	return
																} else {
																	v125 = *(*int32)(unsafe.Add(mBase, _c_F_aofOpenIfNeededOnServerStart[2]))
																	if v125 != int32(-1) {
																		v146 = *(*int32)(unsafe.Add(mBase, _c_F_aofOpenIfNeededOnServerStart[1]))
																		v147 = *(*int32)(unsafe.Add(mBase, uint32(v146)+32))
																		if v147 == int32(0) {
																			v162 = F_getAppendOnlyFileSize(m, v108, int32(0))
																			mBase = m.M
																			v163 = m.ExcPending
																			if v163 != 0 {
																				return
																			} else {
																				*(*int64)(unsafe.Add(mBase, _c_F_aofOpenIfNeededOnServerStart[4])) = v162
																				v165 = int32(_a_F_aofOpenIfNeededOnServerStart_3)
																				*(*int64)(unsafe.Add(mBase, _c_F_aofOpenIfNeededOnServerStart[5])) = v162
																				v168 = *(*int32)(unsafe.Add(mBase, _c_F_aofOpenIfNeededOnServerStart[6]))
																				if int32(2) < v168 {
																					m.G0 = v10 + int32(80)
																					return
																				} else {
																					*(*int32)(unsafe.Add(mBase, uint32(v10)+32)) = v108
																					if v79 != 0 {
																						v175 = int32(_a_F_aofOpenIfNeededOnServerStart_4)
																					} else {
																						v175 = int32(_a_F_aofOpenIfNeededOnServerStart_5)
																					}
																					F__serverLog(m, int32(2), v175, v10+int32(32))
																					mBase = m.M
																					v179 = m.ExcPending
																					if v179 != 0 {
																						return
																					} else {
																						m.G0 = v10 + int32(80)
																						return
																					}
																				}
																			}
																		} else {
																			v150 = F_getAofManifestAsString(m, v146)
																			mBase = m.M
																			v151 = m.ExcPending
																			if v151 != 0 {
																				return
																			} else {
																				v152 = F_writeAofManifestFile(m, v150)
																				mBase = m.M
																				v153 = m.ExcPending
																				if v153 != 0 {
																					return
																				} else {
																					F_sdsfree(m, v150)
																					mBase = m.M
																					v155 = m.ExcPending
																					if v155 != 0 {
																						return
																					} else {
																						if v152 != 0 {
																							m.Env.Exit(m, int32(1))
																							mBase = m.M
																							base.Wasm_trap_unreachable()
																							for {
																							}
																						} else {
																							*(*int32)(unsafe.Add(mBase, uint32(v146)+32)) = int32(0)
																							v162 = F_getAppendOnlyFileSize(m, v108, int32(0))
																							mBase = m.M
																							v163 = m.ExcPending
																							if v163 != 0 {
																								return
																							} else {
																								*(*int64)(unsafe.Add(mBase, _c_F_aofOpenIfNeededOnServerStart[4])) = v162
																								v165 = int32(_a_F_aofOpenIfNeededOnServerStart_3)
																								*(*int64)(unsafe.Add(mBase, _c_F_aofOpenIfNeededOnServerStart[5])) = v162
																								v168 = *(*int32)(unsafe.Add(mBase, _c_F_aofOpenIfNeededOnServerStart[6]))
																								if int32(2) < v168 {
																									m.G0 = v10 + int32(80)
																									return
																								} else {
																									*(*int32)(unsafe.Add(mBase, uint32(v10)+32)) = v108
																									if v79 != 0 {
																										v175 = int32(_a_F_aofOpenIfNeededOnServerStart_4)
																									} else {
																										v175 = int32(_a_F_aofOpenIfNeededOnServerStart_5)
																									}
																									F__serverLog(m, int32(2), v175, v10+int32(32))
																									mBase = m.M
																									v179 = m.ExcPending
																									if v179 != 0 {
																										return
																									} else {
																										m.G0 = v10 + int32(80)
																										return
																									}
																								}
																							}
																						}
																					}
																				}
																			}
																		}
																	} else {
																		v129 = *(*int32)(unsafe.Add(mBase, _c_F_aofOpenIfNeededOnServerStart[6]))
																		if int32(3) < v129 {
																			m.Env.Exit(m, int32(1))
																			mBase = m.M
																			base.Wasm_trap_unreachable()
																			for {
																			}
																		} else {
																			v133 = *(*int32)(unsafe.Add(mBase, _c_F_aofOpenIfNeededOnServerStart[7]))
																			v134 = F___strerror_l(m, v133, v133)
																			mBase = m.M
																			*(*int32)(unsafe.Add(mBase, uint32(v10)+20)) = v134
																			*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = v108
																			F__serverLog(m, int32(3), int32(_a_F_aofOpenIfNeededOnServerStart_6), v10+int32(16))
																			mBase = m.M
																			v142 = m.ExcPending
																			if v142 != 0 {
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
						}
					}
				} else {
					v59 = *(*int32)(unsafe.Add(mBase, _c_F_aofOpenIfNeededOnServerStart[6]))
					if int32(3) < v59 {
						m.Env.Exit(m, int32(1))
						mBase = m.M
						base.Wasm_trap_unreachable()
						for {
						}
					} else {
						v63 = *(*int32)(unsafe.Add(mBase, _c_F_aofOpenIfNeededOnServerStart[3]))
						v65 = *(*int32)(unsafe.Add(mBase, _c_F_aofOpenIfNeededOnServerStart[7]))
						v66 = F___strerror_l(m, v65, v65)
						mBase = m.M
						*(*int32)(unsafe.Add(mBase, uint32(v10)+4)) = v66
						*(*int32)(unsafe.Add(mBase, uint32(v10))) = v63
						F__serverLog(m, int32(3), int32(_a_F_aofOpenIfNeededOnServerStart_8), v10)
						mBase = m.M
						v72 = m.ExcPending
						if v72 != 0 {
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
func F_getAofManifestAsString(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
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
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v40 int32
	_ = v40
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
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v63 int32
	_ = v63
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v88 int32
	_ = v88
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v102 int32
	_ = v102
	var v105 int32
	_ = v105
	var v111 int32
	_ = v111
	var v115 int32
	_ = v115
	var v125 int32
	_ = v125
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	if l0 == int32(0) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	F__serverAssert(m, int32(_a_F_getAofManifestAsString_0), int32(_a_F_getAofManifestAsString_1), int32(191))
	mBase = m.M
	v125 = m.ExcPending
	if v125 != 0 {
		goto L3
	} else {
		goto L34
	}
L2:
	;
	v11 = F_sdsempty(m)
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	return int32(0)
L4:
	;
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v15 == int32(0) {
		v20 = v11
		goto L5
	} else {
		goto L6
	}
L5:
	;
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v23 = v7 + int32(8)
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v21)))
	*(*int32)(unsafe.Add(mBase, uint32(v23)+4)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v23))) = v24
	goto L8
L6:
	;
	v18 = F_aofInfoFormat(m, v11, v15)
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		goto L3
	} else {
		goto L7
	}
L7:
	;
	v20 = v18
	goto L5
L8:
	;
	v29 = v7 + int32(8)
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v29)))
	if v31 == int32(0) {
		goto L11
	} else {
		goto L12
	}
L9:
	;
	v69 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v71 = v7 + int32(8)
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v69)))
	*(*int32)(unsafe.Add(mBase, uint32(v71)+4)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v71))) = v72
	goto L21
L10:
	;
	if v31 == int32(0) {
		v67 = v20
		goto L9
	} else {
		goto L13
	}
L11:
	;
	goto L10
L12:
	;
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v29)+4))
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v31+base.B2i32(v34 == int32(0))<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v29))) = v40
	goto L11
L13:
	;
	v46 = v20
	v47 = v31
	goto L14
L14:
	;
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v47)+8))
	v49 = F_aofInfoFormat(m, v46, v48)
	mBase = m.M
	v50 = m.ExcPending
	if v50 != 0 {
		goto L3
	} else {
		goto L16
	}
L15:
	;
	v67 = v49
	goto L9
L16:
	;
	v52 = v7 + int32(8)
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v52)))
	if v54 == int32(0) {
		goto L18
	} else {
		goto L19
	}
L17:
	;
	if v54 != 0 {
		v46 = v49
		v47 = v54
		goto L14
	} else {
		goto L20
	}
L18:
	;
	goto L17
L19:
	;
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v52)+4))
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v54+base.B2i32(v57 == int32(0))<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v52))) = v63
	goto L18
L20:
	;
	goto L15
L21:
	;
	v77 = v7 + int32(8)
	v79 = *(*int32)(unsafe.Add(mBase, uint32(v77)))
	if v79 == int32(0) {
		goto L24
	} else {
		goto L25
	}
L22:
	;
	m.G0 = v7 + int32(16)
	return v115
L23:
	;
	if v79 == int32(0) {
		v115 = v67
		goto L22
	} else {
		goto L26
	}
L24:
	;
	goto L23
L25:
	;
	v82 = *(*int32)(unsafe.Add(mBase, uint32(v77)+4))
	v88 = *(*int32)(unsafe.Add(mBase, uint32(v79+base.B2i32(v82 == int32(0))<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v77))) = v88
	goto L24
L26:
	;
	v94 = v67
	v95 = v79
	goto L27
L27:
	;
	v96 = *(*int32)(unsafe.Add(mBase, uint32(v95)+8))
	v97 = F_aofInfoFormat(m, v94, v96)
	mBase = m.M
	v98 = m.ExcPending
	if v98 != 0 {
		goto L3
	} else {
		goto L29
	}
L28:
	;
	v115 = v97
	goto L22
L29:
	;
	v100 = v7 + int32(8)
	v102 = *(*int32)(unsafe.Add(mBase, uint32(v100)))
	if v102 == int32(0) {
		goto L31
	} else {
		goto L32
	}
L30:
	;
	if v102 != 0 {
		v94 = v97
		v95 = v102
		goto L27
	} else {
		goto L33
	}
L31:
	;
	goto L30
L32:
	;
	v105 = *(*int32)(unsafe.Add(mBase, uint32(v100)+4))
	v111 = *(*int32)(unsafe.Add(mBase, uint32(v102+base.B2i32(v105 == int32(0))<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v100))) = v111
	goto L31
L33:
	;
	goto L28
L34:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_updateAofAutoGCEnabled(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	v3 = *(*int32)(unsafe.Add(mBase, _c_F_updateAofAutoGCEnabled[0]))
	if v3 != 0 {
		return int32(1)
	} else {
		v4 = F_aofDelHistoryFiles(m)
		mBase = m.M
		v7 = m.ExcPending
		if v7 != 0 {
			return int32(0)
		} else {
			return int32(1)
		}
	}
}
