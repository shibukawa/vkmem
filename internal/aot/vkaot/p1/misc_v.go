package p1

import (
	base "github.com/shibukawa/vkmem/internal/aot/vkaot/base"
	"unsafe"
)

func F_ValkeyModule_OnUnload_lua(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
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
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
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
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	v4 = m.G3
	v7 = m.G380
	v8 = *(*int32)(unsafe.Add(mBase, uint32(v7)))
	v9 = m.T0[v8].(func(*base.Module, int32, int32) int32)(m, l0, v4+int32(_a22))
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		return int32(0)
	} else {
		if v9 == int32(0) {
			v27 = m.G3
			v30 = *(*int32)(unsafe.Add(mBase, uint32(v27)+uint32(_consts[984])))
			v31 = *(*int32)(unsafe.Add(mBase, uint32(v30)))
			F_lua_close(m, v31)
			mBase = m.M
			v33 = m.ExcPending
			if v33 != 0 {
				return int32(0)
			} else {
				v34 = *(*int32)(unsafe.Add(mBase, uint32(v30)+4))
				F_lua_close(m, v34)
				mBase = m.M
				v36 = m.ExcPending
				if v36 != 0 {
					return int32(0)
				} else {
					v37 = *(*int32)(unsafe.Add(mBase, uint32(v30)+8))
					v38 = m.G11
					v39 = *(*int32)(unsafe.Add(mBase, uint32(v38)))
					m.T0[v39].(func(*base.Module, int32))(m, v37)
					mBase = m.M
					v41 = m.ExcPending
					if v41 != 0 {
						return int32(0)
					} else {
						v42 = *(*int32)(unsafe.Add(mBase, uint32(v30)+16))
						v43 = *(*int32)(unsafe.Add(mBase, uint32(v38)))
						m.T0[v43].(func(*base.Module, int32))(m, v42)
						mBase = m.M
						v45 = m.ExcPending
						if v45 != 0 {
							return int32(0)
						} else {
							v46 = *(*int32)(unsafe.Add(mBase, uint32(v30)+20))
							v47 = *(*int32)(unsafe.Add(mBase, uint32(v38)))
							m.T0[v47].(func(*base.Module, int32))(m, v46)
							mBase = m.M
							v49 = m.ExcPending
							if v49 != 0 {
								return int32(0)
							} else {
								v50 = *(*int32)(unsafe.Add(mBase, uint32(v38)))
								m.T0[v50].(func(*base.Module, int32))(m, v30)
								mBase = m.M
								v52 = m.ExcPending
								if v52 != 0 {
									return int32(0)
								} else {
									v53 = int32(0)
									*(*int32)(unsafe.Add(mBase, uint32(v27)+uint32(_consts[984]))) = v53
									return v53
								}
							}
						}
					}
				}
			}
		} else {
			v15 = m.G3
			v21 = m.G10
			v22 = *(*int32)(unsafe.Add(mBase, uint32(v21)))
			m.T0[v22].(func(*base.Module, int32, int32, int32, int32))(m, l0, v15+int32(_a2620), v15+int32(_a2621), int32(0))
			mBase = m.M
			v24 = m.ExcPending
			if v24 != 0 {
				return int32(0)
			} else {
				return int32(1)
			}
		}
	}
}
func F__vsyslog(m *base.Module, l0 int32, l1 int32, l2 int32) {
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
	var v35 int64
	_ = v35
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v70 int32
	_ = v70
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v111 int32
	_ = v111
	var v117 int32
	_ = v117
	var v123 int32
	_ = v123
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v135 int32
	_ = v135
	var v138 int32
	_ = v138
	var v140 int32
	_ = v140
	var v149 int32
	_ = v149
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v156 int32
	_ = v156
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v166 int32
	_ = v166
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v181 int32
	_ = v181
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v202 int32
	_ = v202
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	v9 = m.G0
	v11 = v9 - int32(1168)
	m.G0 = v11
	v14 = int32(9116376)
	v15 = *(*int32)(unsafe.Add(mBase, _consts[18]))
	v17 = *(*int32)(unsafe.Add(mBase, _consts[1060]))
	if int32(-1) < v17 {
	} else {
		v20 = int32(0)
		v25 = F_socket(m, int32(1), int32(524290), v20)
		mBase = m.M
		*(*int32)(unsafe.Add(mBase, _consts[1060])) = v25
		if v25 < v20 {
		} else {
			v31 = F_connect(m, v25, int32(_a2772), int32(12))
			mBase = m.M
		}
	}
	v32 = int32(0)
	v33 = *(*int32)(unsafe.Add(mBase, _consts[1058]))
	v35 = F___time(m, v32)
	mBase = m.M
	*(*int64)(unsafe.Add(mBase, uint32(v11)+1144)) = v35
	v40 = v11 + int32(1100)
	v41 = F___gmtime_r(m, v11+int32(1144), v40)
	mBase = m.M
	v49 = F___strftime_l(m, v11+int32(1152), int32(16), int32(_a2773), v40, int32(_a2774))
	mBase = m.M
	v50 = m.ExcPending
	if v50 != 0 {
		return
	} else {
		if l0&int32(1016) != 0 {
			v54 = int32(0)
		} else {
			v54 = v33
		}
		v56 = int32(0)
		v57 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1059])))
		if v57&int32(1) == v56 {
			v63 = int32(0)
		} else {
			v62 = F___syscall_getpid(m)
			mBase = m.M
			v63 = v62
		}
		*(*int32)(unsafe.Add(mBase, uint32(v11+int32(52)))) = v63
		v70 = base.B2i32(v63 == int32(0))
		*(*int32)(unsafe.Add(mBase, uint32(v11+int32(56)))) = v70 + int32(_a2775)
		*(*int32)(unsafe.Add(mBase, uint32(v11+int32(48)))) = v70 + int32(_a2776)
		*(*int32)(unsafe.Add(mBase, uint32(v11)+44)) = int32(9128384)
		*(*int32)(unsafe.Add(mBase, uint32(v11)+32)) = v54 | l0
		*(*int32)(unsafe.Add(mBase, uint32(v11)+40)) = v11 + int32(60)
		*(*int32)(unsafe.Add(mBase, uint32(v11)+36)) = v11 + int32(1152)
		v96 = F_snprintf(m, v11+int32(64), int32(1024), int32(_a2777), v11+int32(32))
		mBase = m.M
		v97 = m.ExcPending
		if v97 != 0 {
			return
		} else {
			*(*int32)(unsafe.Add(mBase, _consts[18])) = v15
			v103 = int32(1024) - v96
			v104 = F_vsnprintf(m, v96+(v11+int32(64)), v103, l1, l2)
			mBase = m.M
			v105 = m.ExcPending
			if v105 != 0 {
				return
			} else {
				if v104 < int32(0) {
					m.G0 = v11 + int32(1168)
					return
				} else {
					if base.Ui32(v104) < base.Ui32(v103) {
						v111 = v104 + v96
					} else {
						v111 = int32(1023)
					}
					v117 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v111+(v11+int32(64))+int32(-1)))))
					if v117 == int32(10) {
						v127 = v111
					} else {
						v123 = int32(10)
						*(*uint8)(unsafe.Add(mBase, uint32(v11+int32(64)+v111))) = uint8(v123)
						v127 = v111 + int32(1)
					}
					v128 = int32(0)
					v129 = *(*int32)(unsafe.Add(mBase, _consts[1060]))
					v135 = F_sendto(m, v129, v11+int32(64), v127, v128, v128, v128)
					mBase = m.M
					if int32(-1) < v135 {
						v196 = int32(0)
						v197 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1059])))
						if v197&int32(32) == v196 {
							m.G0 = v11 + int32(1168)
							return
						} else {
							v202 = *(*int32)(unsafe.Add(mBase, uint32(v11)+60))
							*(*int32)(unsafe.Add(mBase, uint32(v11))) = v127 - v202
							*(*int32)(unsafe.Add(mBase, uint32(v11)+4)) = v202 + (v11 + int32(64))
							v211 = F_dprintf(m, int32(2), int32(_a2778), v11)
							mBase = m.M
							v212 = m.ExcPending
							if v212 != 0 {
								return
							} else {
								m.G0 = v11 + int32(1168)
								return
							}
						}
					} else {
						v138 = *(*int32)(unsafe.Add(mBase, _consts[18]))
						v140 = int32(1)
						if base.Ui32(v138+int32(-14)) < base.Ui32(int32(2)) {
							v149 = v140
						} else {
							if v138 == int32(53) {
								v149 = v140
							} else {
								v149 = base.B2i32(v138 == int32(64))
							}
						}
						if v149 == int32(0) {
							v169 = int32(0)
							v170 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1059])))
							if v170&int32(2) == v169 {
								v196 = int32(0)
								v197 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1059])))
								if v197&int32(32) == v196 {
									m.G0 = v11 + int32(1168)
									return
								} else {
									v202 = *(*int32)(unsafe.Add(mBase, uint32(v11)+60))
									*(*int32)(unsafe.Add(mBase, uint32(v11))) = v127 - v202
									*(*int32)(unsafe.Add(mBase, uint32(v11)+4)) = v202 + (v11 + int32(64))
									v211 = F_dprintf(m, int32(2), int32(_a2778), v11)
									mBase = m.M
									v212 = m.ExcPending
									if v212 != 0 {
										return
									} else {
										m.G0 = v11 + int32(1168)
										return
									}
								}
							} else {
								v177 = int32(0)
								v178 = F_open(m, int32(_a2779), int32(524545), v177)
								mBase = m.M
								if v178 < v177 {
									v196 = int32(0)
									v197 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1059])))
									if v197&int32(32) == v196 {
										m.G0 = v11 + int32(1168)
										return
									} else {
										v202 = *(*int32)(unsafe.Add(mBase, uint32(v11)+60))
										*(*int32)(unsafe.Add(mBase, uint32(v11))) = v127 - v202
										*(*int32)(unsafe.Add(mBase, uint32(v11)+4)) = v202 + (v11 + int32(64))
										v211 = F_dprintf(m, int32(2), int32(_a2778), v11)
										mBase = m.M
										v212 = m.ExcPending
										if v212 != 0 {
											return
										} else {
											m.G0 = v11 + int32(1168)
											return
										}
									}
								} else {
									v181 = *(*int32)(unsafe.Add(mBase, uint32(v11)+60))
									*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = v127 - v181
									*(*int32)(unsafe.Add(mBase, uint32(v11)+20)) = v181 + (v11 + int32(64))
									v191 = F_dprintf(m, v178, int32(_a2778), v11+int32(16))
									mBase = m.M
									v192 = m.ExcPending
									if v192 != 0 {
										return
									} else {
										v193 = F_close(m, v178)
										mBase = m.M
										v196 = int32(0)
										v197 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1059])))
										if v197&int32(32) == v196 {
											m.G0 = v11 + int32(1168)
											return
										} else {
											v202 = *(*int32)(unsafe.Add(mBase, uint32(v11)+60))
											*(*int32)(unsafe.Add(mBase, uint32(v11))) = v127 - v202
											*(*int32)(unsafe.Add(mBase, uint32(v11)+4)) = v202 + (v11 + int32(64))
											v211 = F_dprintf(m, int32(2), int32(_a2778), v11)
											mBase = m.M
											v212 = m.ExcPending
											if v212 != 0 {
												return
											} else {
												m.G0 = v11 + int32(1168)
												return
											}
										}
									}
								}
							}
						} else {
							v152 = int32(0)
							v153 = *(*int32)(unsafe.Add(mBase, _consts[1060]))
							v156 = F_connect(m, v153, int32(_a2772), int32(12))
							mBase = m.M
							if v156 < v152 {
								v169 = int32(0)
								v170 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1059])))
								if v170&int32(2) == v169 {
									v196 = int32(0)
									v197 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1059])))
									if v197&int32(32) == v196 {
										m.G0 = v11 + int32(1168)
										return
									} else {
										v202 = *(*int32)(unsafe.Add(mBase, uint32(v11)+60))
										*(*int32)(unsafe.Add(mBase, uint32(v11))) = v127 - v202
										*(*int32)(unsafe.Add(mBase, uint32(v11)+4)) = v202 + (v11 + int32(64))
										v211 = F_dprintf(m, int32(2), int32(_a2778), v11)
										mBase = m.M
										v212 = m.ExcPending
										if v212 != 0 {
											return
										} else {
											m.G0 = v11 + int32(1168)
											return
										}
									}
								} else {
									v177 = int32(0)
									v178 = F_open(m, int32(_a2779), int32(524545), v177)
									mBase = m.M
									if v178 < v177 {
										v196 = int32(0)
										v197 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1059])))
										if v197&int32(32) == v196 {
											m.G0 = v11 + int32(1168)
											return
										} else {
											v202 = *(*int32)(unsafe.Add(mBase, uint32(v11)+60))
											*(*int32)(unsafe.Add(mBase, uint32(v11))) = v127 - v202
											*(*int32)(unsafe.Add(mBase, uint32(v11)+4)) = v202 + (v11 + int32(64))
											v211 = F_dprintf(m, int32(2), int32(_a2778), v11)
											mBase = m.M
											v212 = m.ExcPending
											if v212 != 0 {
												return
											} else {
												m.G0 = v11 + int32(1168)
												return
											}
										}
									} else {
										v181 = *(*int32)(unsafe.Add(mBase, uint32(v11)+60))
										*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = v127 - v181
										*(*int32)(unsafe.Add(mBase, uint32(v11)+20)) = v181 + (v11 + int32(64))
										v191 = F_dprintf(m, v178, int32(_a2778), v11+int32(16))
										mBase = m.M
										v192 = m.ExcPending
										if v192 != 0 {
											return
										} else {
											v193 = F_close(m, v178)
											mBase = m.M
											v196 = int32(0)
											v197 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1059])))
											if v197&int32(32) == v196 {
												m.G0 = v11 + int32(1168)
												return
											} else {
												v202 = *(*int32)(unsafe.Add(mBase, uint32(v11)+60))
												*(*int32)(unsafe.Add(mBase, uint32(v11))) = v127 - v202
												*(*int32)(unsafe.Add(mBase, uint32(v11)+4)) = v202 + (v11 + int32(64))
												v211 = F_dprintf(m, int32(2), int32(_a2778), v11)
												mBase = m.M
												v212 = m.ExcPending
												if v212 != 0 {
													return
												} else {
													m.G0 = v11 + int32(1168)
													return
												}
											}
										}
									}
								}
							} else {
								v159 = int32(0)
								v160 = *(*int32)(unsafe.Add(mBase, _consts[1060]))
								v166 = F_sendto(m, v160, v11+int32(64), v127, v159, v159, v159)
								mBase = m.M
								if int32(-1) < v166 {
									v196 = int32(0)
									v197 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1059])))
									if v197&int32(32) == v196 {
										m.G0 = v11 + int32(1168)
										return
									} else {
										v202 = *(*int32)(unsafe.Add(mBase, uint32(v11)+60))
										*(*int32)(unsafe.Add(mBase, uint32(v11))) = v127 - v202
										*(*int32)(unsafe.Add(mBase, uint32(v11)+4)) = v202 + (v11 + int32(64))
										v211 = F_dprintf(m, int32(2), int32(_a2778), v11)
										mBase = m.M
										v212 = m.ExcPending
										if v212 != 0 {
											return
										} else {
											m.G0 = v11 + int32(1168)
											return
										}
									}
								} else {
									v169 = int32(0)
									v170 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1059])))
									if v170&int32(2) == v169 {
										v196 = int32(0)
										v197 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1059])))
										if v197&int32(32) == v196 {
											m.G0 = v11 + int32(1168)
											return
										} else {
											v202 = *(*int32)(unsafe.Add(mBase, uint32(v11)+60))
											*(*int32)(unsafe.Add(mBase, uint32(v11))) = v127 - v202
											*(*int32)(unsafe.Add(mBase, uint32(v11)+4)) = v202 + (v11 + int32(64))
											v211 = F_dprintf(m, int32(2), int32(_a2778), v11)
											mBase = m.M
											v212 = m.ExcPending
											if v212 != 0 {
												return
											} else {
												m.G0 = v11 + int32(1168)
												return
											}
										}
									} else {
										v177 = int32(0)
										v178 = F_open(m, int32(_a2779), int32(524545), v177)
										mBase = m.M
										if v178 < v177 {
											v196 = int32(0)
											v197 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1059])))
											if v197&int32(32) == v196 {
												m.G0 = v11 + int32(1168)
												return
											} else {
												v202 = *(*int32)(unsafe.Add(mBase, uint32(v11)+60))
												*(*int32)(unsafe.Add(mBase, uint32(v11))) = v127 - v202
												*(*int32)(unsafe.Add(mBase, uint32(v11)+4)) = v202 + (v11 + int32(64))
												v211 = F_dprintf(m, int32(2), int32(_a2778), v11)
												mBase = m.M
												v212 = m.ExcPending
												if v212 != 0 {
													return
												} else {
													m.G0 = v11 + int32(1168)
													return
												}
											}
										} else {
											v181 = *(*int32)(unsafe.Add(mBase, uint32(v11)+60))
											*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = v127 - v181
											*(*int32)(unsafe.Add(mBase, uint32(v11)+20)) = v181 + (v11 + int32(64))
											v191 = F_dprintf(m, v178, int32(_a2778), v11+int32(16))
											mBase = m.M
											v192 = m.ExcPending
											if v192 != 0 {
												return
											} else {
												v193 = F_close(m, v178)
												mBase = m.M
												v196 = int32(0)
												v197 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1059])))
												if v197&int32(32) == v196 {
													m.G0 = v11 + int32(1168)
													return
												} else {
													v202 = *(*int32)(unsafe.Add(mBase, uint32(v11)+60))
													*(*int32)(unsafe.Add(mBase, uint32(v11))) = v127 - v202
													*(*int32)(unsafe.Add(mBase, uint32(v11)+4)) = v202 + (v11 + int32(64))
													v211 = F_dprintf(m, int32(2), int32(_a2778), v11)
													mBase = m.M
													v212 = m.ExcPending
													if v212 != 0 {
														return
													} else {
														m.G0 = v11 + int32(1168)
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
func F_vectorPush(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v5 != v6 {
		v21 = v4
		v22 = v5
		*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v22 + int32(1)
		v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
		return v21 + v26*v22
	} else {
		v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
		if v5 != 0 {
			v12 = v5 << (uint(int32(1)) % 32)
		} else {
			v12 = int32(8)
		}
		v14 = F_valkey_realloc(m, v4, v8*v12)
		mBase = m.M
		v17 = m.ExcPending
		if v17 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v12
			*(*int32)(unsafe.Add(mBase, uint32(l0))) = v14
			v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			v21 = v14
			v22 = v20
			*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v22 + int32(1)
			v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
			return v21 + v26*v22
		}
	}
}
func F_vfiprintf(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	v4 = int32(0)
	v6 = F___vfprintf_internal(m, l0, l1, l2, v4, v4)
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		return v6
	}
}
func F_vkmem_call_sighandler(m *base.Module, l0 int32, l1 int32) {
	var v4 int32
	_ = v4
	m.T0[l0].(func(*base.Module, int32))(m, l1)
	v4 = m.ExcPending
	if v4 != 0 {
		return
	} else {
		return
	}
}
func F_vkmem_dlclose(m *base.Module, l0 int32) int32 {
	return int32(0)
}
func F_vkmem_dlopen(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	if l0 != 0 {
		v5 = int32(0)
		*(*int32)(unsafe.Add(mBase, _consts[0])) = int32(_a0)
		return v5
	} else {
		return int32(_a1)
	}
}
func F_vkmem_main(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v23 int32
	_ = v23
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
	var v37 int32
	_ = v37
	var v42 int32
	_ = v42
	var v48 int32
	_ = v48
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v84 int32
	_ = v84
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v91 int32
	_ = v91
	var v96 int32
	_ = v96
	var v103 int32
	_ = v103
	var v106 int32
	_ = v106
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
	var v129 int32
	_ = v129
	var v132 int32
	_ = v132
	var v136 int32
	_ = v136
	var v141 int32
	_ = v141
	var v143 int32
	_ = v143
	var v147 int32
	_ = v147
	var v153 int32
	_ = v153
	var v156 int32
	_ = v156
	var v162 int32
	_ = v162
	var v166 int32
	_ = v166
	var v168 int32
	_ = v168
	var v176 int32
	_ = v176
	var v178 int32
	_ = v178
	var v181 int32
	_ = v181
	var v195 int32
	_ = v195
	var v198 int32
	_ = v198
	v3 = int32(0)
	if l1 < int32(1) {
		v96 = v3
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v103 = v96 << (uint(int32(2)) % 32)
	v106 = F_emscripten_builtin_malloc(m, v103+int32(4))
	mBase = m.M
	if v96 == int32(0) {
		goto L12
	} else {
		goto L13
	}
L2:
	;
	v15 = l1 & int32(3)
	v16 = int32(0)
	if base.Ui32(l1) < base.Ui32(int32(4)) {
		v66 = v16
		v69 = v16
		goto L3
	} else {
		goto L4
	}
L3:
	;
	if v15 == int32(0) {
		v96 = v66
		goto L1
	} else {
		goto L8
	}
L4:
	;
	v23 = int32(0)
	v29 = v23
	v32 = v23
	v34 = v23
	goto L5
L5:
	;
	v35 = l0 + v32
	v36 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v35))))
	v37 = int32(0)
	v42 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v35+int32(1)))))
	v48 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v35+int32(2)))))
	v54 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v35+int32(3)))))
	v57 = v29 + base.B2i32(v36 == v37) + base.B2i32(v42 == v37) + base.B2i32(v48 == v37) + base.B2i32(v54 == v37)
	v58 = int32(4)
	v59 = v32 + v58
	v61 = v34 + v58
	if v61 != l1&int32(2147483644) {
		v29 = v57
		v32 = v59
		v34 = v61
		goto L5
	} else {
		goto L7
	}
L6:
	;
	v66 = v57
	v69 = v59
	goto L3
L7:
	;
	goto L6
L8:
	;
	v77 = v66
	v79 = v16
	v80 = v69
	goto L9
L9:
	;
	v84 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0+v80))))
	v87 = v77 + base.B2i32(v84 == int32(0))
	v88 = int32(1)
	v91 = v79 + v88
	if v91 != v15 {
		v77 = v87
		v79 = v91
		v80 = v80 + v88
		goto L9
	} else {
		goto L11
	}
L10:
	;
	v96 = v87
	goto L1
L11:
	;
	goto L10
L12:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v106+v103))) = int32(0)
	v195 = F_main(m, v96, v106)
	mBase = m.M
	v198 = m.ExcPending
	if v198 != 0 {
		goto L33
	} else {
		goto L34
	}
L13:
	;
	v109 = l0
	v111 = v3
	goto L14
L14:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v106+v111<<(uint(int32(2))%32)))) = v109
	if v109&int32(3) == int32(0) {
		v143 = v109
		goto L18
	} else {
		goto L19
	}
L15:
	;
	goto L12
L16:
	;
	v178 = int32(1)
	v181 = v111 + v178
	if v181 != v96 {
		v109 = v109 + v176 + v178
		v111 = v181
		goto L14
	} else {
		goto L32
	}
L17:
	;
	v176 = v168 - v109
	goto L16
L18:
	;
	v147 = v143
	goto L26
L19:
	;
	v129 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v109))))
	if v129 != 0 {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	v132 = v109
	goto L22
L21:
	;
	v176 = v109 - v109
	goto L16
L22:
	;
	v136 = v132 + int32(1)
	if v136&int32(3) == int32(0) {
		v143 = v136
		goto L18
	} else {
		goto L24
	}
L24:
	;
	v141 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v136))))
	if v141 != 0 {
		v132 = v136
		goto L22
	} else {
		goto L25
	}
L25:
	;
	v168 = v136
	goto L17
L26:
	;
	v153 = *(*int32)(unsafe.Add(mBase, uint32(v147)))
	v156 = int32(-2139062144)
	if (int32(16843008)-v153|v153)&v156 == v156 {
		v147 = v147 + int32(4)
		goto L26
	} else {
		goto L28
	}
L27:
	;
	v162 = v147
	goto L29
L28:
	;
	goto L27
L29:
	;
	v166 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v162))))
	if v166 != 0 {
		v162 = v162 + int32(1)
		goto L29
	} else {
		goto L31
	}
L30:
	;
	v168 = v162
	goto L17
L31:
	;
	goto L30
L32:
	;
	goto L15
L33:
	;
	return int32(0)
L34:
	;
	return v195
}
func F_vsiprintf(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	v5 = F_vsniprintf(m, l0, int32(2147483647), l1, l2)
	v8 = m.ExcPending
	if v8 != 0 {
		return int32(0)
	} else {
		return v5
	}
}
func F_vsniprintf(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v17 int32
	_ = v17
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
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v55 int32
	_ = v55
	v7 = m.G0
	v9 = v7 - int32(160)
	m.G0 = v9
	v17 = F__emscripten_memcpy_bulkmem(m, v9+int32(8), int32(_a2789), int32(144))
	mBase = m.M
	if int32(0) < l1 {
		v24 = l0
		v25 = l1
		*(*int32)(unsafe.Add(mBase, uint32(v9)+52)) = v24
		*(*int32)(unsafe.Add(mBase, uint32(v9)+28)) = v24
		v29 = int32(-2) - v24
		if base.Ui32(v25) < base.Ui32(v29) {
			v31 = v25
		} else {
			v31 = v29
		}
		*(*int32)(unsafe.Add(mBase, uint32(v9)+56)) = v31
		v33 = v24 + v31
		*(*int32)(unsafe.Add(mBase, uint32(v9)+36)) = v33
		*(*int32)(unsafe.Add(mBase, uint32(v9)+24)) = v33
		v38 = F_vfiprintf(m, v9+int32(8), l2, l3)
		mBase = m.M
		v41 = m.ExcPending
		if v41 != 0 {
			return int32(0)
		} else {
			if v24 == int32(-2) {
				v55 = v38
			} else {
				v44 = *(*int32)(unsafe.Add(mBase, uint32(v9)+28))
				v45 = *(*int32)(unsafe.Add(mBase, uint32(v9)+24))
				v48 = int32(0)
				*(*uint8)(unsafe.Add(mBase, uint32(v44-base.B2i32(v44 == v45)))) = uint8(v48)
				v55 = v38
			}
			m.G0 = v9 + int32(160)
			return v55
		}
	} else {
		if l1 != 0 {
			*(*int32)(unsafe.Add(mBase, _consts[18])) = int32(61)
			v55 = int32(-1)
			m.G0 = v9 + int32(160)
			return v55
		} else {
			v24 = v9 + int32(159)
			v25 = int32(1)
			*(*int32)(unsafe.Add(mBase, uint32(v9)+52)) = v24
			*(*int32)(unsafe.Add(mBase, uint32(v9)+28)) = v24
			v29 = int32(-2) - v24
			if base.Ui32(v25) < base.Ui32(v29) {
				v31 = v25
			} else {
				v31 = v29
			}
			*(*int32)(unsafe.Add(mBase, uint32(v9)+56)) = v31
			v33 = v24 + v31
			*(*int32)(unsafe.Add(mBase, uint32(v9)+36)) = v33
			*(*int32)(unsafe.Add(mBase, uint32(v9)+24)) = v33
			v38 = F_vfiprintf(m, v9+int32(8), l2, l3)
			mBase = m.M
			v41 = m.ExcPending
			if v41 != 0 {
				return int32(0)
			} else {
				if v24 == int32(-2) {
					v55 = v38
				} else {
					v44 = *(*int32)(unsafe.Add(mBase, uint32(v9)+28))
					v45 = *(*int32)(unsafe.Add(mBase, uint32(v9)+24))
					v48 = int32(0)
					*(*uint8)(unsafe.Add(mBase, uint32(v44-base.B2i32(v44 == v45)))) = uint8(v48)
					v55 = v38
				}
				m.G0 = v9 + int32(160)
				return v55
			}
		}
	}
}
func F_vsscanf(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
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
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	v5 = m.G0
	v6 = int32(144)
	v7 = v5 - v6
	m.G0 = v7
	v12 = F__emscripten_memset_bulkmem(m, v7, base.I32_extend8_s(int32(0)), v6)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v12)+76)) = int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(v12)+44)) = l0
	*(*int32)(unsafe.Add(mBase, uint32(v12)+32)) = int32(1390)
	*(*int32)(unsafe.Add(mBase, uint32(v12)+84)) = l0
	v19 = F_vfscanf(m, v12, l1, l2)
	mBase = m.M
	v22 = m.ExcPending
	if v22 != 0 {
		return int32(0)
	} else {
		m.G0 = v12 + int32(144)
		return v19
	}
}
