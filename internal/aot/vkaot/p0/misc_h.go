package p0

import (
	base "github.com/shibukawa/vkmem/internal/aot/vkaot/base"
	"unsafe"
)

func F_handleBlockedClientsTimeout(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v12 int64
	_ = v12
	var v15 int64
	_ = v15
	var v17 int32
	_ = v17
	var v23 int64
	_ = v23
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
	var v49 int32
	_ = v49
	var v50 int64
	_ = v50
	var v51 int64
	_ = v51
	var v53 int64
	_ = v53
	var v55 int64
	_ = v55
	var v58 int64
	_ = v58
	var v60 int64
	_ = v60
	var v62 int64
	_ = v62
	var v64 int64
	_ = v64
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v92 int32
	_ = v92
	var v97 int32
	_ = v97
	var v98 int64
	_ = v98
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v114 int32
	_ = v114
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v126 int32
	_ = v126
	v6 = m.G0
	v8 = v6 - int32(304)
	m.G0 = v8
	v11 = *(*int32)(unsafe.Add(mBase, _consts[883]))
	v12 = *(*int64)(unsafe.Add(mBase, uint32(v11)+8))
	goto L2
L1:
	;
	m.G0 = v8 + int32(304)
	return
L2:
	;
	if v12 == int64(0) {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v15 = F_mstime(m)
	mBase = m.M
	v17 = *(*int32)(unsafe.Add(mBase, _consts[883]))
	*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = v17
	*(*int32)(unsafe.Add(mBase, uint32(v8))) = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v8)+20)) = int32(128)
	v23 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v8)+12)) = v23
	*(*int64)(unsafe.Add(mBase, uint32(v8)+296)) = v23
	*(*int64)(unsafe.Add(mBase, uint32(v8)+160)) = int64(137438953472)
	*(*int32)(unsafe.Add(mBase, uint32(v8)+8)) = v8 + int32(24)
	*(*int32)(unsafe.Add(mBase, uint32(v8)+156)) = v8 + int32(168)
	goto L4
L4:
	;
	v36 = int32(0)
	v38 = F_raxSeek(m, v8, int32(_a4), v36, v36)
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	return
L6:
	;
	v40 = F_raxNext(m, v8)
	mBase = m.M
	v41 = m.ExcPending
	if v41 != 0 {
		goto L5
	} else {
		goto L8
	}
L7:
	;
	F_raxStop(m, v8)
	mBase = m.M
	v126 = m.ExcPending
	if v126 != 0 {
		goto L5
	} else {
		goto L22
	}
L8:
	;
	if v40 == int32(0) {
		goto L7
	} else {
		goto L9
	}
L9:
	;
	goto L10
L10:
	;
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v8)+8))
	v50 = *(*int64)(unsafe.Add(mBase, uint32(v49)))
	v51 = int64(56)
	v53 = int64(65280)
	v55 = int64(40)
	v58 = int64(16711680)
	v60 = int64(24)
	v62 = int64(4278190080)
	v64 = int64(8)
	if base.Ui64(v15) <= base.Ui64(v50<<(uint(v51)%64)|v50&v53<<(uint(v55)%64)|(v50&v58<<(uint(v60)%64)|v50&v62<<(uint(v64)%64))|(int64(base.Ui64(v50)>>(uint(v64)%64))&v62|int64(base.Ui64(v50)>>(uint(v60)%64))&v58|(int64(base.Ui64(v50)>>(uint(v55)%64))&v53|int64(base.Ui64(v50)>>(uint(v51)%64))))) {
		goto L7
	} else {
		goto L12
	}
L11:
	;
	goto L7
L12:
	;
	v87 = *(*int32)(unsafe.Add(mBase, uint32(v49)+8))
	v88 = *(*int32)(unsafe.Add(mBase, uint32(v87)+204))
	*(*int32)(unsafe.Add(mBase, uint32(v87)+204)) = v88 & int32(-513)
	v92 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v87)+200)))
	if v92&int32(16) == int32(0) {
		v105 = v49
		goto L13
	} else {
		goto L14
	}
L13:
	;
	v108 = *(*int32)(unsafe.Add(mBase, _consts[883]))
	v109 = *(*int32)(unsafe.Add(mBase, uint32(v8)+16))
	v111 = F_raxRemove(m, v108, v105, v109, int32(0))
	mBase = m.M
	v112 = m.ExcPending
	if v112 != 0 {
		goto L5
	} else {
		goto L18
	}
L14:
	;
	v97 = *(*int32)(unsafe.Add(mBase, uint32(v87)+116))
	v98 = *(*int64)(unsafe.Add(mBase, uint32(v97)+8))
	if v98 == int64(0) {
		v105 = v49
		goto L13
	} else {
		goto L15
	}
L15:
	;
	if v15 <= v98 {
		v105 = v49
		goto L13
	} else {
		goto L16
	}
L16:
	;
	F_unblockClientOnTimeout(m, v87)
	mBase = m.M
	v103 = m.ExcPending
	if v103 != 0 {
		goto L5
	} else {
		goto L17
	}
L17:
	;
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v8)+8))
	v105 = v104
	goto L13
L18:
	;
	v114 = int32(0)
	v116 = F_raxSeek(m, v8, int32(_a4), v114, v114)
	mBase = m.M
	v117 = m.ExcPending
	if v117 != 0 {
		goto L5
	} else {
		goto L19
	}
L19:
	;
	v118 = F_raxNext(m, v8)
	mBase = m.M
	v119 = m.ExcPending
	if v119 != 0 {
		goto L5
	} else {
		goto L20
	}
L20:
	;
	if v118 != 0 {
		goto L10
	} else {
		goto L21
	}
L21:
	;
	goto L11
L22:
	;
	goto L1
}
func F_handleReadResult(m *base.Module, l0 int32) int32 {
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
	var v16 int32
	_ = v16
	var v18 int64
	_ = v18
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v54 int32
	_ = v54
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v96 int32
	_ = v96
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v108 int64
	_ = v108
	var v110 int64
	_ = v110
	var v111 int64
	_ = v111
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v119 int32
	_ = v119
	var v121 int32
	_ = v121
	var v123 int32
	_ = v123
	var v126 int32
	_ = v126
	var v127 int64
	_ = v127
	var v130 int32
	_ = v130
	var v131 int64
	_ = v131
	var v134 int32
	_ = v134
	var v147 int32
	_ = v147
	var v153 int64
	_ = v153
	var v155 int64
	_ = v155
	var v156 int32
	_ = v156
	var v158 int64
	_ = v158
	var v161 int32
	_ = v161
	var v163 int64
	_ = v163
	var v170 int32
	_ = v170
	var v177 int32
	_ = v177
	var v184 int32
	_ = v184
	var v186 int32
	_ = v186
	var v188 int64
	_ = v188
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v200 int32
	_ = v200
	var v203 int32
	_ = v203
	v2 = int32(0)
	v6 = m.G0
	v8 = v6 - int32(32)
	m.G0 = v8
	v11 = *(*int32)(unsafe.Add(mBase, _consts[299]))
	if base.B2i32(v11 == v2) == int32(0) {
		F__serverAssert(m, int32(_a777), int32(_a774), int32(3169))
		mBase = m.M
		v170 = m.ExcPending
		if v170 != 0 {
			return int32(0)
		} else {
			F_abort(m)
			mBase = m.M
			base.Wasm_trap_unreachable()
			for {
			}
		}
	} else {
		v16 = int32(_a69)
		v18 = *(*int64)(unsafe.Add(mBase, _consts[409]))
		*(*int64)(unsafe.Add(mBase, _consts[409])) = v18 + int64(1)
		v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+284))
		if int32(0) < v22 {
			v108 = *(*int64)(unsafe.Add(mBase, _consts[37]))
			*(*int64)(unsafe.Add(mBase, uint32(l0)+88)) = v108
			v110 = *(*int64)(unsafe.Add(mBase, uint32(l0)+232))
			v111 = int64(*(*int32)(unsafe.Add(mBase, uint32(l0)+284)))
			*(*int64)(unsafe.Add(mBase, uint32(l0)+232)) = v110 + v111
			v115 = int32(1)
			v116 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+200)))
			if v116&v115 != 0 {
				v123 = v115
				v126 = v123
			} else {
				v119 = *(*int32)(unsafe.Add(mBase, uint32(l0)+216))
				if v119 != 0 {
					v121 = F_isImportSlotMigrationJob(m, v119)
					mBase = m.M
					v123 = v121
					v126 = v123
				} else {
					v126 = int32(0)
				}
			}
			v127 = int64(*(*int32)(unsafe.Add(mBase, uint32(l0)+284)))
			if v126 == int32(0) {
				v161 = int32(_a69)
				v163 = *(*int64)(unsafe.Add(mBase, _consts[410]))
				*(*int64)(unsafe.Add(mBase, _consts[410])) = v163 + v127
			} else {
				v130 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
				v131 = *(*int64)(unsafe.Add(mBase, uint32(v130)+40))
				*(*int64)(unsafe.Add(mBase, uint32(v130)+40)) = v131 + v127
				v134 = *(*int32)(unsafe.Add(mBase, uint32(l0)+200))
				if v134&int32(1) != 0 {
					v186 = int32(_a69)
					v188 = *(*int64)(unsafe.Add(mBase, _consts[411]))
					*(*int64)(unsafe.Add(mBase, _consts[411])) = v188 + v127
				} else {
					if v134&int32(2) == int32(0) {
						if v134&int32(262144) != 0 {
							v155 = v127
						} else {
							v147 = *(*int32)(unsafe.Add(mBase, uint32(l0)+216))
							if v147 == int32(0) {
								v155 = v127
							} else {
								v153 = int64(*(*int32)(unsafe.Add(mBase, uint32(l0)+284)))
								v155 = v153
							}
						}
					} else {
						if v134&int32(262148) == int32(4) {
							v147 = *(*int32)(unsafe.Add(mBase, uint32(l0)+216))
							if v147 == int32(0) {
								v155 = v127
							} else {
								v153 = int64(*(*int32)(unsafe.Add(mBase, uint32(l0)+284)))
								v155 = v153
							}
						} else {
							v155 = v127
						}
					}
					v156 = int32(_a69)
					v158 = *(*int64)(unsafe.Add(mBase, _consts[412]))
					*(*int64)(unsafe.Add(mBase, _consts[412])) = v158 + v155
				}
			}
			v193 = int32(0)
			v194 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+288)))
			if v194&int32(1) == v193 {
				v203 = v193
				m.G0 = v8 + int32(32)
				return v203
			} else {
				F_handleQbLimitReached(m, l0)
				mBase = m.M
				v200 = m.ExcPending
				if v200 != 0 {
					return int32(0)
				} else {
					v203 = int32(-1)
					m.G0 = v8 + int32(32)
					return v203
				}
			}
		} else {
			v25 = int32(-1)
			switch v22 + int32(1) {
			case 0:
				v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				v29 = *(*int32)(unsafe.Add(mBase, uint32(v28)+4))
				if v29 == int32(3) {
					v203 = v25
					m.G0 = v8 + int32(32)
					return v203
				} else {
					v33 = *(*int32)(unsafe.Add(mBase, _consts[15]))
					if int32(1) < v33 {
						v47 = *(*int32)(unsafe.Add(mBase, uint32(l0)+200))
						if v47&int32(1280) != 0 {
							v203 = v25
							m.G0 = v8 + int32(32)
							return v203
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(l0)+200)) = v47 | int32(1024)
							v54 = *(*int32)(unsafe.Add(mBase, _consts[122]))
							if v54 == int32(0) {
								v62 = *(*int32)(unsafe.Add(mBase, _consts[413]))
								v63 = F_listAddNodeTail(m, v62, l0)
								mBase = m.M
								v64 = m.ExcPending
								if v64 != 0 {
									return int32(0)
								} else {
									v203 = v25
									m.G0 = v8 + int32(32)
									return v203
								}
							} else {
								v58 = *(*int32)(unsafe.Add(mBase, _consts[413]))
								v59 = F_listSearchKey(m, v58, l0)
								mBase = m.M
								v60 = m.ExcPending
								if v60 != 0 {
									return int32(0)
								} else {
									if v59 != 0 {
										F__serverAssertWithInfo(m, l0, int32(0), int32(_a778), int32(_a774), int32(2277))
										mBase = m.M
										v177 = m.ExcPending
										if v177 != 0 {
											return int32(0)
										} else {
											F_abort(m)
											mBase = m.M
											base.Wasm_trap_unreachable()
											for {
											}
										}
									} else {
										v62 = *(*int32)(unsafe.Add(mBase, _consts[413]))
										v63 = F_listAddNodeTail(m, v62, l0)
										mBase = m.M
										v64 = m.ExcPending
										if v64 != 0 {
											return int32(0)
										} else {
											v203 = v25
											m.G0 = v8 + int32(32)
											return v203
										}
									}
								}
							}
						}
					} else {
						v36 = *(*int32)(unsafe.Add(mBase, uint32(v28)))
						v37 = *(*int32)(unsafe.Add(mBase, uint32(v36)+88))
						v38 = m.T0[v37].(func(*base.Module, int32) int32)(m, v28)
						mBase = m.M
						v41 = m.ExcPending
						if v41 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v8))) = v38
							F__serverLog(m, int32(1), int32(_a779), v8)
							mBase = m.M
							v46 = m.ExcPending
							if v46 != 0 {
								return int32(0)
							} else {
								v47 = *(*int32)(unsafe.Add(mBase, uint32(l0)+200))
								if v47&int32(1280) != 0 {
									v203 = v25
									m.G0 = v8 + int32(32)
									return v203
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(l0)+200)) = v47 | int32(1024)
									v54 = *(*int32)(unsafe.Add(mBase, _consts[122]))
									if v54 == int32(0) {
										v62 = *(*int32)(unsafe.Add(mBase, _consts[413]))
										v63 = F_listAddNodeTail(m, v62, l0)
										mBase = m.M
										v64 = m.ExcPending
										if v64 != 0 {
											return int32(0)
										} else {
											v203 = v25
											m.G0 = v8 + int32(32)
											return v203
										}
									} else {
										v58 = *(*int32)(unsafe.Add(mBase, _consts[413]))
										v59 = F_listSearchKey(m, v58, l0)
										mBase = m.M
										v60 = m.ExcPending
										if v60 != 0 {
											return int32(0)
										} else {
											if v59 != 0 {
												F__serverAssertWithInfo(m, l0, int32(0), int32(_a778), int32(_a774), int32(2277))
												mBase = m.M
												v177 = m.ExcPending
												if v177 != 0 {
													return int32(0)
												} else {
													F_abort(m)
													mBase = m.M
													base.Wasm_trap_unreachable()
													for {
													}
												}
											} else {
												v62 = *(*int32)(unsafe.Add(mBase, _consts[413]))
												v63 = F_listAddNodeTail(m, v62, l0)
												mBase = m.M
												v64 = m.ExcPending
												if v64 != 0 {
													return int32(0)
												} else {
													v203 = v25
													m.G0 = v8 + int32(32)
													return v203
												}
											}
										}
									}
								}
							}
						}
					}
				}
			case 1:
				v66 = *(*int32)(unsafe.Add(mBase, _consts[15]))
				if int32(1) < v66 {
					v89 = *(*int32)(unsafe.Add(mBase, uint32(l0)+200))
					if v89&int32(1280) != 0 {
						v203 = v25
						m.G0 = v8 + int32(32)
						return v203
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(l0)+200)) = v89 | int32(1024)
						v96 = *(*int32)(unsafe.Add(mBase, _consts[122]))
						if v96 == int32(0) {
							v104 = *(*int32)(unsafe.Add(mBase, _consts[413]))
							v105 = F_listAddNodeTail(m, v104, l0)
							mBase = m.M
							v106 = m.ExcPending
							if v106 != 0 {
								return int32(0)
							} else {
								v203 = v25
								m.G0 = v8 + int32(32)
								return v203
							}
						} else {
							v100 = *(*int32)(unsafe.Add(mBase, _consts[413]))
							v101 = F_listSearchKey(m, v100, l0)
							mBase = m.M
							v102 = m.ExcPending
							if v102 != 0 {
								return int32(0)
							} else {
								if v101 != 0 {
									F__serverAssertWithInfo(m, l0, int32(0), int32(_a778), int32(_a774), int32(2277))
									mBase = m.M
									v184 = m.ExcPending
									if v184 != 0 {
										return int32(0)
									} else {
										F_abort(m)
										mBase = m.M
										base.Wasm_trap_unreachable()
										for {
										}
									}
								} else {
									v104 = *(*int32)(unsafe.Add(mBase, _consts[413]))
									v105 = F_listAddNodeTail(m, v104, l0)
									mBase = m.M
									v106 = m.ExcPending
									if v106 != 0 {
										return int32(0)
									} else {
										v203 = v25
										m.G0 = v8 + int32(32)
										return v203
									}
								}
							}
						}
					}
				} else {
					v69 = F_sdsempty(m)
					mBase = m.M
					v70 = m.ExcPending
					if v70 != 0 {
						return int32(0)
					} else {
						v72 = *(*int32)(unsafe.Add(mBase, _consts[151]))
						v73 = F_catClientInfoString(m, v69, l0, v72)
						mBase = m.M
						v74 = m.ExcPending
						if v74 != 0 {
							return int32(0)
						} else {
							v76 = *(*int32)(unsafe.Add(mBase, _consts[15]))
							if int32(1) < v76 {
								F_sdsfree(m, v73)
								mBase = m.M
								v87 = m.ExcPending
								if v87 != 0 {
									return int32(0)
								} else {
									v89 = *(*int32)(unsafe.Add(mBase, uint32(l0)+200))
									if v89&int32(1280) != 0 {
										v203 = v25
										m.G0 = v8 + int32(32)
										return v203
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(l0)+200)) = v89 | int32(1024)
										v96 = *(*int32)(unsafe.Add(mBase, _consts[122]))
										if v96 == int32(0) {
											v104 = *(*int32)(unsafe.Add(mBase, _consts[413]))
											v105 = F_listAddNodeTail(m, v104, l0)
											mBase = m.M
											v106 = m.ExcPending
											if v106 != 0 {
												return int32(0)
											} else {
												v203 = v25
												m.G0 = v8 + int32(32)
												return v203
											}
										} else {
											v100 = *(*int32)(unsafe.Add(mBase, _consts[413]))
											v101 = F_listSearchKey(m, v100, l0)
											mBase = m.M
											v102 = m.ExcPending
											if v102 != 0 {
												return int32(0)
											} else {
												if v101 != 0 {
													F__serverAssertWithInfo(m, l0, int32(0), int32(_a778), int32(_a774), int32(2277))
													mBase = m.M
													v184 = m.ExcPending
													if v184 != 0 {
														return int32(0)
													} else {
														F_abort(m)
														mBase = m.M
														base.Wasm_trap_unreachable()
														for {
														}
													}
												} else {
													v104 = *(*int32)(unsafe.Add(mBase, _consts[413]))
													v105 = F_listAddNodeTail(m, v104, l0)
													mBase = m.M
													v106 = m.ExcPending
													if v106 != 0 {
														return int32(0)
													} else {
														v203 = v25
														m.G0 = v8 + int32(32)
														return v203
													}
												}
											}
										}
									}
								}
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = v73
								F__serverLog(m, int32(1), int32(_a780), v8+int32(16))
								mBase = m.M
								v85 = m.ExcPending
								if v85 != 0 {
									return int32(0)
								} else {
									F_sdsfree(m, v73)
									mBase = m.M
									v87 = m.ExcPending
									if v87 != 0 {
										return int32(0)
									} else {
										v89 = *(*int32)(unsafe.Add(mBase, uint32(l0)+200))
										if v89&int32(1280) != 0 {
											v203 = v25
											m.G0 = v8 + int32(32)
											return v203
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(l0)+200)) = v89 | int32(1024)
											v96 = *(*int32)(unsafe.Add(mBase, _consts[122]))
											if v96 == int32(0) {
												v104 = *(*int32)(unsafe.Add(mBase, _consts[413]))
												v105 = F_listAddNodeTail(m, v104, l0)
												mBase = m.M
												v106 = m.ExcPending
												if v106 != 0 {
													return int32(0)
												} else {
													v203 = v25
													m.G0 = v8 + int32(32)
													return v203
												}
											} else {
												v100 = *(*int32)(unsafe.Add(mBase, _consts[413]))
												v101 = F_listSearchKey(m, v100, l0)
												mBase = m.M
												v102 = m.ExcPending
												if v102 != 0 {
													return int32(0)
												} else {
													if v101 != 0 {
														F__serverAssertWithInfo(m, l0, int32(0), int32(_a778), int32(_a774), int32(2277))
														mBase = m.M
														v184 = m.ExcPending
														if v184 != 0 {
															return int32(0)
														} else {
															F_abort(m)
															mBase = m.M
															base.Wasm_trap_unreachable()
															for {
															}
														}
													} else {
														v104 = *(*int32)(unsafe.Add(mBase, _consts[413]))
														v105 = F_listAddNodeTail(m, v104, l0)
														mBase = m.M
														v106 = m.ExcPending
														if v106 != 0 {
															return int32(0)
														} else {
															v203 = v25
															m.G0 = v8 + int32(32)
															return v203
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
			default:
				v203 = v25
				m.G0 = v8 + int32(32)
				return v203
			}
		}
	}
}
func F_hdelCommand(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
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
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
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
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v104 int32
	_ = v104
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v125 int64
	_ = v125
	var v126 int64
	_ = v126
	var v134 int64
	_ = v134
	var v136 int32
	_ = v136
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(v7)+4))
	v10 = *(*int32)(unsafe.Add(mBase, _consts[233]))
	v11 = F_lookupKeyWriteOrReply(m, l0, v8, v10)
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	return
L2:
	;
	return
L3:
	;
	if v11 == int32(0) {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v16 = F_checkType(m, l0, v11, int32(4))
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		goto L2
	} else {
		goto L5
	}
L5:
	;
	if v16 != 0 {
		goto L1
	} else {
		goto L6
	}
L6:
	;
	v18 = F_hashTypeHasVolatileFields(m, v11)
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		goto L2
	} else {
		goto L7
	}
L7:
	;
	v20 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
	if v20&int32(240) != int32(32) {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v30 = int32(0)
	v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v31 < int32(3) {
		v75 = v30
		goto L13
	} else {
		goto L14
	}
L9:
	;
	v25 = F_objectGetVal(m, v11)
	mBase = m.M
	v26 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v25)+28)))
	v28 = v26 + int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v25)+28)) = uint16(v28)
	goto L10
L10:
	;
	goto L8
L11:
	;
	F_addReplyLongLong(m, l0, v134)
	mBase = m.M
	v136 = m.ExcPending
	if v136 != 0 {
		goto L2
	} else {
		goto L41
	}
L12:
	;
	v100 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v101 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v102 = *(*int32)(unsafe.Add(mBase, uint32(v101)+4))
	F_signalModifiedKey(m, l0, v100, v102)
	mBase = m.M
	v104 = m.ExcPending
	if v104 != 0 {
		goto L2
	} else {
		goto L36
	}
L13:
	;
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
	if v78&int32(240) != int32(32) {
		goto L28
	} else {
		goto L29
	}
L14:
	;
	v38 = v30
	v39 = int32(2)
	goto L15
L15:
	;
	v41 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v41+v39<<(uint(int32(2))%32))))
	v46 = F_objectGetVal(m, v45)
	mBase = m.M
	v47 = F_hashTypeDelete(m, v11, v46)
	mBase = m.M
	v48 = m.ExcPending
	if v48 != 0 {
		goto L2
	} else {
		goto L18
	}
L16:
	;
	v75 = v67
	goto L13
L17:
	;
	v69 = v39 + int32(1)
	v70 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v69 < v70 {
		v38 = v67
		v39 = v69
		goto L15
	} else {
		goto L27
	}
L18:
	;
	if v47 == int32(0) {
		v67 = v38
		goto L17
	} else {
		goto L19
	}
L19:
	;
	v52 = v38 + int32(1)
	v53 = F_hashTypeLength(m, v11)
	mBase = m.M
	v54 = m.ExcPending
	if v54 != 0 {
		goto L2
	} else {
		goto L20
	}
L20:
	;
	if v53 != 0 {
		v67 = v52
		goto L17
	} else {
		goto L21
	}
L21:
	;
	if v18 == int32(0) {
		goto L22
	} else {
		goto L23
	}
L22:
	;
	v60 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v61 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v61)+4))
	v63 = F_dbDelete(m, v60, v62)
	mBase = m.M
	v64 = m.ExcPending
	if v64 != 0 {
		goto L2
	} else {
		goto L25
	}
L23:
	;
	v57 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	F_dbUntrackKeyWithVolatileItems(m, v57, v11)
	mBase = m.M
	v59 = m.ExcPending
	if v59 != 0 {
		goto L2
	} else {
		goto L24
	}
L24:
	;
	goto L22
L25:
	;
	if v52 != 0 {
		v97 = v52
		v98 = int32(1)
		goto L12
	} else {
		goto L26
	}
L26:
	;
	v134 = int64(0)
	goto L11
L27:
	;
	goto L16
L28:
	;
	if v75 != 0 {
		goto L31
	} else {
		goto L32
	}
L29:
	;
	v83 = F_objectGetVal(m, v11)
	mBase = m.M
	F_hashtableResumeAutoShrink(m, v83)
	mBase = m.M
	v85 = m.ExcPending
	if v85 != 0 {
		goto L2
	} else {
		goto L30
	}
L30:
	;
	goto L28
L31:
	;
	v87 = int32(0)
	v88 = F_hashTypeHasVolatileFields(m, v11)
	mBase = m.M
	v89 = m.ExcPending
	if v89 != 0 {
		goto L2
	} else {
		goto L33
	}
L32:
	;
	v134 = int64(0)
	goto L11
L33:
	;
	if v18 == v88 {
		v97 = v75
		v98 = v87
		goto L12
	} else {
		goto L34
	}
L34:
	;
	v91 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	F_dbUpdateObjectWithVolatileItemsTracking(m, v91, v11)
	mBase = m.M
	v93 = m.ExcPending
	if v93 != 0 {
		goto L2
	} else {
		goto L35
	}
L35:
	;
	v97 = v75
	v98 = v87
	goto L12
L36:
	;
	v107 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v108 = *(*int32)(unsafe.Add(mBase, uint32(v107)+4))
	v109 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v110 = *(*int32)(unsafe.Add(mBase, uint32(v109)+28))
	F_notifyKeyspaceEvent(m, int32(64), int32(_a1489), v108, v110)
	mBase = m.M
	v112 = m.ExcPending
	if v112 != 0 {
		goto L2
	} else {
		goto L37
	}
L37:
	;
	if v98 == int32(0) {
		goto L38
	} else {
		goto L39
	}
L38:
	;
	v123 = int32(_a69)
	v125 = *(*int64)(unsafe.Add(mBase, _consts[60]))
	v126 = base.I64_extend_i32_s(v97)
	*(*int64)(unsafe.Add(mBase, _consts[60])) = v125 + v126
	v134 = v126
	goto L11
L39:
	;
	v117 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v118 = *(*int32)(unsafe.Add(mBase, uint32(v117)+4))
	v119 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v120 = *(*int32)(unsafe.Add(mBase, uint32(v119)+28))
	F_notifyKeyspaceEvent(m, int32(4), int32(_a132), v118, v120)
	mBase = m.M
	v122 = m.ExcPending
	if v122 != 0 {
		goto L2
	} else {
		goto L40
	}
L40:
	;
	goto L38
L41:
	;
	goto L1
}
func F_hdr_close(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	if l0 == int32(0) {
		return
	} else {
		v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
		F_valkey_free(m, v4)
		mBase = m.M
		v6 = m.ExcPending
		if v6 != 0 {
			return
		} else {
			F_valkey_free(m, l0)
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
func F_hexpireGenericCommand(m *base.Module, l0 int32, l1 int64, l2 int32) {
	mBase := m.M
	_ = mBase
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v28 int32
	_ = v28
	var v36 int32
	_ = v36
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
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v101 int32
	_ = v101
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v107 int64
	_ = v107
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v136 int64
	_ = v136
	var v139 int32
	_ = v139
	var v143 int32
	_ = v143
	var v151 int32
	_ = v151
	var v153 int64
	_ = v153
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v162 int32
	_ = v162
	var v166 int32
	_ = v166
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v170 int64
	_ = v170
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v178 int32
	_ = v178
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v187 int32
	_ = v187
	var v189 int32
	_ = v189
	var v192 int32
	_ = v192
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v199 int32
	_ = v199
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v209 int32
	_ = v209
	var v211 int32
	_ = v211
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v217 int64
	_ = v217
	var v221 int32
	_ = v221
	var v225 int32
	_ = v225
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v234 int32
	_ = v234
	var v235 int64
	_ = v235
	var v237 int64
	_ = v237
	var v244 int32
	_ = v244
	var v250 int32
	_ = v250
	var v251 int32
	_ = v251
	var v252 int32
	_ = v252
	var v258 int32
	_ = v258
	var v259 int32
	_ = v259
	var v261 int32
	_ = v261
	var v263 int32
	_ = v263
	var v267 int32
	_ = v267
	var v271 int32
	_ = v271
	var v272 int32
	_ = v272
	var v277 int32
	_ = v277
	var v279 int32
	_ = v279
	var v280 int32
	_ = v280
	var v284 int64
	_ = v284
	var v285 int32
	_ = v285
	var v286 int32
	_ = v286
	var v288 int32
	_ = v288
	var v290 int32
	_ = v290
	var v292 int32
	_ = v292
	var v294 int32
	_ = v294
	var v295 int32
	_ = v295
	var v296 int32
	_ = v296
	var v297 int32
	_ = v297
	var v299 int32
	_ = v299
	var v302 int32
	_ = v302
	var v304 int64
	_ = v304
	var v309 int32
	_ = v309
	var v310 int32
	_ = v310
	var v311 int32
	_ = v311
	var v313 int32
	_ = v313
	var v314 int32
	_ = v314
	var v315 int32
	_ = v315
	var v316 int32
	_ = v316
	var v317 int32
	_ = v317
	var v318 int32
	_ = v318
	var v319 int32
	_ = v319
	var v320 int32
	_ = v320
	var v323 int32
	_ = v323
	var v324 int32
	_ = v324
	var v325 int32
	_ = v325
	var v326 int32
	_ = v326
	var v328 int32
	_ = v328
	var v333 int32
	_ = v333
	var v335 int32
	_ = v335
	var v336 int32
	_ = v336
	var v358 int32
	_ = v358
	v17 = m.G0
	v19 = v17 - int32(32)
	m.G0 = v19
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v21)+8))
	v23 = *(*int32)(unsafe.Add(mBase, uint32(v21)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v19)+20)) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v19)+8)) = int64(0)
	v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v28 < int32(5) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v19 + int32(32)
	return
L2:
	;
	F_addReplyError(m, l0, int32(_a1492))
	mBase = m.M
	v358 = m.ExcPending
	if v358 != 0 {
		goto L20
	} else {
		goto L81
	}
L3:
	;
	v36 = int32(3)
	goto L4
L4:
	;
	v48 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v48+v36<<(uint(int32(2))%32))))
	v53 = F_objectGetVal(m, v52)
	mBase = m.M
	v54 = int32(_a1493)
	v57 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v53))))
	if v57 != 0 {
		goto L9
	} else {
		goto L10
	}
L5:
	;
	goto L2
L6:
	;
	v335 = v36 + int32(1)
	v336 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v335 < v336+int32(-1) {
		v36 = v335
		goto L4
	} else {
		goto L80
	}
L7:
	;
	if v89-v91 != 0 {
		goto L6
	} else {
		goto L19
	}
L8:
	;
	v89 = F_tolower(m, v85)
	mBase = m.M
	v90 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v86))))
	v91 = F_tolower(m, v90)
	mBase = m.M
	goto L7
L9:
	;
	v59 = v53
	v60 = v54
	v61 = v57
	goto L12
L10:
	;
	v85 = int32(0)
	v86 = v54
	goto L8
L11:
	;
	v85 = v82 & int32(255)
	v86 = v81
	goto L8
L12:
	;
	v63 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v60))))
	if v63 == int32(0) {
		v81 = v60
		v82 = v61
		goto L11
	} else {
		goto L14
	}
L13:
	;
	v81 = v75
	v82 = int32(0)
	goto L11
L14:
	;
	v67 = v61 & int32(255)
	if v67 == v63 {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	v74 = int32(1)
	v75 = v60 + v74
	v76 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v59)+1)))
	if v76 != 0 {
		v59 = v59 + v74
		v60 = v75
		v61 = v76
		goto L12
	} else {
		goto L18
	}
L16:
	;
	v69 = F_tolower(m, v67)
	mBase = m.M
	v70 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v60))))
	v71 = F_tolower(m, v70)
	mBase = m.M
	if v69 == v71 {
		goto L15
	} else {
		goto L17
	}
L17:
	;
	v73 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v59))))
	v81 = v60
	v82 = v73
	goto L11
L18:
	;
	goto L13
L19:
	;
	v95 = F_parseExtendedExpireArgumentsOrReply(m, l0, v19+int32(20), v36)
	mBase = m.M
	v96 = m.ExcPending
	if v96 != 0 {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	return
L21:
	;
	if v95 != 0 {
		goto L1
	} else {
		goto L22
	}
L22:
	;
	v97 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v101 = *(*int32)(unsafe.Add(mBase, uint32(v97+v36<<(uint(int32(2))%32))+4))
	v105 = F_getLongLongFromObjectOrReply(m, l0, v101, v19+int32(8), int32(0))
	mBase = m.M
	v106 = m.ExcPending
	if v106 != 0 {
		goto L20
	} else {
		goto L23
	}
L23:
	;
	if v105 != 0 {
		goto L1
	} else {
		goto L24
	}
L24:
	;
	v107 = *(*int64)(unsafe.Add(mBase, uint32(v19)+8))
	if v107 == int64(0) {
		goto L2
	} else {
		goto L25
	}
L25:
	;
	v110 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v112 = v36 + int32(2)
	if v107 != base.I64_extend_i32_s(v110-v112) {
		goto L2
	} else {
		goto L26
	}
L26:
	;
	v118 = F_convertExpireArgumentToUnixTime(m, l0, v22, l1, l2, v19+int32(24))
	mBase = m.M
	v119 = m.ExcPending
	if v119 != 0 {
		goto L20
	} else {
		goto L27
	}
L27:
	;
	if v118 == int32(-1) {
		goto L1
	} else {
		goto L28
	}
L28:
	;
	v122 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v123 = F_lookupKeyWrite(m, v122, v23)
	mBase = m.M
	v124 = m.ExcPending
	if v124 != 0 {
		goto L20
	} else {
		goto L29
	}
L29:
	;
	v126 = F_checkType(m, l0, v123, int32(4))
	mBase = m.M
	v127 = m.ExcPending
	if v127 != 0 {
		goto L20
	} else {
		goto L30
	}
L30:
	;
	if v126 != 0 {
		goto L1
	} else {
		goto L31
	}
L31:
	;
	v128 = F_hashTypeHasVolatileFields(m, v123)
	mBase = m.M
	v129 = m.ExcPending
	if v129 != 0 {
		goto L20
	} else {
		goto L32
	}
L32:
	;
	F_initDeferredReplyBuffer(m, l0)
	mBase = m.M
	v131 = m.ExcPending
	if v131 != 0 {
		goto L20
	} else {
		goto L33
	}
L33:
	;
	v132 = *(*int32)(unsafe.Add(mBase, uint32(v19)+8))
	F_addReplyArrayLen(m, l0, v132)
	mBase = m.M
	v134 = m.ExcPending
	if v134 != 0 {
		goto L20
	} else {
		goto L34
	}
L34:
	;
	v135 = int32(0)
	v136 = *(*int64)(unsafe.Add(mBase, uint32(v19)+8))
	if int64(1) <= v136 {
		goto L36
	} else {
		goto L37
	}
L35:
	;
	if v250|v252 == int32(0) {
		goto L52
	} else {
		goto L53
	}
L36:
	;
	v143 = int32(0)
	v151 = v143
	v153 = int64(0)
	v157 = v135
	v158 = v143
	v159 = v143
	goto L38
L37:
	;
	v139 = int32(0)
	v244 = v139
	v250 = v135
	v251 = v139
	v252 = v139
	goto L35
L38:
	;
	v162 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v166 = (v112 + base.I32_wrap_i64(v153)) << (uint(int32(2)) % 32)
	v168 = *(*int32)(unsafe.Add(mBase, uint32(v162+v166)))
	v169 = F_objectGetVal(m, v168)
	mBase = m.M
	v170 = *(*int64)(unsafe.Add(mBase, uint32(v19)+24))
	v171 = *(*int32)(unsafe.Add(mBase, uint32(v19)+20))
	v172 = F_hashTypeSetExpire(m, v123, v169, v170, v171)
	mBase = m.M
	v173 = m.ExcPending
	if v173 != 0 {
		goto L20
	} else {
		goto L43
	}
L39:
	;
	v244 = v225
	v250 = v227
	v251 = v228
	v252 = v229
	goto L35
L40:
	;
	F_addReplyLongLong(m, l0, base.I64_extend_i32_s(v172))
	mBase = m.M
	v234 = m.ExcPending
	if v234 != 0 {
		goto L20
	} else {
		goto L50
	}
L41:
	;
	if v158 != 0 {
		v202 = v151
		v203 = v158
		goto L44
	} else {
		goto L45
	}
L42:
	;
	v225 = v151
	v227 = v157
	v228 = v158
	v229 = v159 + int32(1)
	goto L40
L43:
	;
	switch v172 + int32(-1) {
	case 0:
		goto L42
	case 1:
		goto L41
	default:
		v225 = v151
		v227 = v157
		v228 = v158
		v229 = v159
		goto L40
	}
L44:
	;
	v209 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v211 = *(*int32)(unsafe.Add(mBase, uint32(v209+v166)))
	*(*int32)(unsafe.Add(mBase, uint32(v203+v202<<(uint(int32(2))%32)))) = v211
	F_incrRefCount(m, v211)
	mBase = m.M
	v214 = m.ExcPending
	if v214 != 0 {
		goto L20
	} else {
		goto L49
	}
L45:
	;
	v178 = *(*int32)(unsafe.Add(mBase, uint32(v19)+8))
	v183 = F_valkey_malloc(m, v178<<(uint(int32(2))%32)+int32(12))
	mBase = m.M
	v184 = m.ExcPending
	if v184 != 0 {
		goto L20
	} else {
		goto L46
	}
L46:
	;
	v187 = v183 + v151<<(uint(int32(2))%32)
	v189 = *(*int32)(unsafe.Add(mBase, _consts[236]))
	*(*int32)(unsafe.Add(mBase, uint32(v187))) = v189
	F_incrRefCount(m, v189)
	mBase = m.M
	v192 = m.ExcPending
	if v192 != 0 {
		goto L20
	} else {
		goto L47
	}
L47:
	;
	v195 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v196 = *(*int32)(unsafe.Add(mBase, uint32(v195)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v187+int32(4)))) = v196
	F_incrRefCount(m, v196)
	mBase = m.M
	v199 = m.ExcPending
	if v199 != 0 {
		goto L20
	} else {
		goto L48
	}
L48:
	;
	v202 = v151 + int32(2)
	v203 = v183
	goto L44
L49:
	;
	v215 = int32(_a69)
	v217 = *(*int64)(unsafe.Add(mBase, _consts[971]))
	*(*int64)(unsafe.Add(mBase, _consts[971])) = v217 + int64(1)
	v221 = int32(1)
	v225 = v202 + v221
	v227 = v157 + v221
	v228 = v203
	v229 = v159
	goto L40
L50:
	;
	v235 = *(*int64)(unsafe.Add(mBase, uint32(v19)+8))
	v237 = v153 + int64(1)
	if v237 < v235 {
		v151 = v225
		v153 = v237
		v157 = v227
		v158 = v228
		v159 = v229
		goto L38
	} else {
		goto L51
	}
L51:
	;
	goto L39
L52:
	;
	F_commitDeferredReplyBuffer(m, l0, int32(1))
	mBase = m.M
	v333 = m.ExcPending
	if v333 != 0 {
		goto L20
	} else {
		goto L79
	}
L53:
	;
	v258 = F_hashTypeHasVolatileFields(m, v123)
	mBase = m.M
	v259 = m.ExcPending
	if v259 != 0 {
		goto L20
	} else {
		goto L55
	}
L54:
	;
	if v250 == int32(0) {
		goto L60
	} else {
		goto L61
	}
L55:
	;
	if v128 == v258 {
		goto L54
	} else {
		goto L56
	}
L56:
	;
	v261 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	F_dbUpdateObjectWithVolatileItemsTracking(m, v261, v123)
	mBase = m.M
	v263 = m.ExcPending
	if v263 != 0 {
		goto L20
	} else {
		goto L57
	}
L57:
	;
	goto L54
L58:
	;
	v302 = int32(_a69)
	v304 = *(*int64)(unsafe.Add(mBase, _consts[60]))
	*(*int64)(unsafe.Add(mBase, _consts[60])) = v304 + base.I64_extend_i32_s(v250+v252)
	v309 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v310 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v311 = *(*int32)(unsafe.Add(mBase, uint32(v310)+4))
	F_signalModifiedKey(m, l0, v309, v311)
	mBase = m.M
	v313 = m.ExcPending
	if v313 != 0 {
		goto L20
	} else {
		goto L74
	}
L59:
	;
	v294 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v295 = *(*int32)(unsafe.Add(mBase, uint32(v294)+4))
	v296 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v297 = *(*int32)(unsafe.Add(mBase, uint32(v296)+28))
	F_notifyKeyspaceEvent(m, int32(64), v292, v295, v297)
	mBase = m.M
	v299 = m.ExcPending
	if v299 != 0 {
		goto L20
	} else {
		goto L73
	}
L60:
	;
	if v252 == int32(0) {
		goto L58
	} else {
		goto L63
	}
L61:
	;
	F_replaceClientCommandVector(m, l0, v244, v251)
	mBase = m.M
	v267 = m.ExcPending
	if v267 != 0 {
		goto L20
	} else {
		goto L62
	}
L62:
	;
	v292 = int32(_a1491)
	goto L59
L63:
	;
	v271 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
	v272 = *(*int32)(unsafe.Add(mBase, uint32(v271)+48))
	if v272 == int32(243) {
		goto L64
	} else {
		goto L65
	}
L64:
	;
	v280 = int32(_a1495)
	if l1 != int64(0) {
		goto L67
	} else {
		goto L68
	}
L65:
	;
	v277 = *(*int32)(unsafe.Add(mBase, _consts[742]))
	F_rewriteClientCommandArgument(m, l0, int32(0), v277)
	mBase = m.M
	v279 = m.ExcPending
	if v279 != 0 {
		goto L20
	} else {
		goto L66
	}
L66:
	;
	goto L64
L67:
	;
	v284 = *(*int64)(unsafe.Add(mBase, uint32(v19)+24))
	v285 = F_createStringObjectFromLongLong(m, v284)
	mBase = m.M
	v286 = m.ExcPending
	if v286 != 0 {
		goto L20
	} else {
		goto L70
	}
L68:
	;
	if l2 != 0 {
		v292 = v280
		goto L59
	} else {
		goto L69
	}
L69:
	;
	goto L67
L70:
	;
	F_rewriteClientCommandArgument(m, l0, int32(2), v285)
	mBase = m.M
	v288 = m.ExcPending
	if v288 != 0 {
		goto L20
	} else {
		goto L71
	}
L71:
	;
	F_decrRefCount(m, v285)
	mBase = m.M
	v290 = m.ExcPending
	if v290 != 0 {
		goto L20
	} else {
		goto L72
	}
L72:
	;
	v292 = v280
	goto L59
L73:
	;
	goto L58
L74:
	;
	v314 = F_hashTypeLength(m, v123)
	mBase = m.M
	v315 = m.ExcPending
	if v315 != 0 {
		goto L20
	} else {
		goto L75
	}
L75:
	;
	if v314 != 0 {
		goto L52
	} else {
		goto L76
	}
L76:
	;
	v316 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v317 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v318 = *(*int32)(unsafe.Add(mBase, uint32(v317)+4))
	v319 = F_dbDelete(m, v316, v318)
	mBase = m.M
	v320 = m.ExcPending
	if v320 != 0 {
		goto L20
	} else {
		goto L77
	}
L77:
	;
	v323 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v324 = *(*int32)(unsafe.Add(mBase, uint32(v323)+4))
	v325 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v326 = *(*int32)(unsafe.Add(mBase, uint32(v325)+28))
	F_notifyKeyspaceEvent(m, int32(4), int32(_a132), v324, v326)
	mBase = m.M
	v328 = m.ExcPending
	if v328 != 0 {
		goto L20
	} else {
		goto L78
	}
L78:
	;
	goto L52
L79:
	;
	goto L1
L80:
	;
	goto L5
L81:
	;
	goto L1
}
func F_hgetallCommand(m *base.Module, l0 int32) {
	var v4 int32
	_ = v4
	F_genericHgetallCommand(m, l0, int32(3))
	v4 = m.ExcPending
	if v4 != 0 {
		return
	} else {
		return
	}
}
func F_hgetexCommand(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v31 int64
	_ = v31
	var v36 int32
	_ = v36
	var v42 int32
	_ = v42
	var v57 int32
	_ = v57
	var v61 int32
	_ = v61
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
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v76 int32
	_ = v76
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
	var v85 int32
	_ = v85
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v102 int32
	_ = v102
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v119 int32
	_ = v119
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v125 int64
	_ = v125
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v138 int32
	_ = v138
	var v139 int64
	_ = v139
	var v153 int32
	_ = v153
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
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
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v191 int32
	_ = v191
	var v195 int32
	_ = v195
	var v198 int32
	_ = v198
	var v205 int64
	_ = v205
	var v206 int32
	_ = v206
	var v207 int64
	_ = v207
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v214 int64
	_ = v214
	var v218 int32
	_ = v218
	var v222 int32
	_ = v222
	var v224 int32
	_ = v224
	var v225 int64
	_ = v225
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v230 int32
	_ = v230
	var v232 int32
	_ = v232
	var v236 int32
	_ = v236
	var v241 int32
	_ = v241
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v251 int32
	_ = v251
	var v254 int32
	_ = v254
	var v255 int32
	_ = v255
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v260 int32
	_ = v260
	var v261 int32
	_ = v261
	var v263 int32
	_ = v263
	var v265 int32
	_ = v265
	var v266 int32
	_ = v266
	var v269 int32
	_ = v269
	var v270 int32
	_ = v270
	var v275 int32
	_ = v275
	var v276 int32
	_ = v276
	var v281 int32
	_ = v281
	var v282 int32
	_ = v282
	var v284 int32
	_ = v284
	var v286 int32
	_ = v286
	var v287 int32
	_ = v287
	var v290 int32
	_ = v290
	var v291 int32
	_ = v291
	var v297 int32
	_ = v297
	var v298 int32
	_ = v298
	var v302 int32
	_ = v302
	var v306 int32
	_ = v306
	var v309 int32
	_ = v309
	var v317 int32
	_ = v317
	var v319 int32
	_ = v319
	var v320 int32
	_ = v320
	var v321 int32
	_ = v321
	var v322 int32
	_ = v322
	var v324 int32
	_ = v324
	var v330 int32
	_ = v330
	var v340 int32
	_ = v340
	var v341 int32
	_ = v341
	var v345 int32
	_ = v345
	var v347 int32
	_ = v347
	var v349 int32
	_ = v349
	var v350 int32
	_ = v350
	var v352 int32
	_ = v352
	var v355 int32
	_ = v355
	var v357 int32
	_ = v357
	var v358 int32
	_ = v358
	var v359 int32
	_ = v359
	var v360 int32
	_ = v360
	var v363 int32
	_ = v363
	var v365 int64
	_ = v365
	var v369 int32
	_ = v369
	var v371 int32
	_ = v371
	var v372 int32
	_ = v372
	var v373 int64
	_ = v373
	var v375 int32
	_ = v375
	var v376 int32
	_ = v376
	var v381 int32
	_ = v381
	var v383 int32
	_ = v383
	var v384 int32
	_ = v384
	var v385 int32
	_ = v385
	var v386 int32
	_ = v386
	var v392 int32
	_ = v392
	var v394 int32
	_ = v394
	var v397 int32
	_ = v397
	var v398 int32
	_ = v398
	var v403 int32
	_ = v403
	var v404 int32
	_ = v404
	var v406 int32
	_ = v406
	var v407 int32
	_ = v407
	var v411 int32
	_ = v411
	var v418 int64
	_ = v418
	var v419 int32
	_ = v419
	var v420 int32
	_ = v420
	var v423 int32
	_ = v423
	var v424 int32
	_ = v424
	var v431 int32
	_ = v431
	var v432 int32
	_ = v432
	var v435 int32
	_ = v435
	var v436 int32
	_ = v436
	var v438 int32
	_ = v438
	var v443 int32
	_ = v443
	var v444 int32
	_ = v444
	var v445 int32
	_ = v445
	var v446 int32
	_ = v446
	var v447 int32
	_ = v447
	var v448 int32
	_ = v448
	var v450 int32
	_ = v450
	var v454 int32
	_ = v454
	var v458 int32
	_ = v458
	var v459 int32
	_ = v459
	var v461 int64
	_ = v461
	var v465 int32
	_ = v465
	var v466 int32
	_ = v466
	var v467 int32
	_ = v467
	var v469 int32
	_ = v469
	var v470 int32
	_ = v470
	var v471 int32
	_ = v471
	var v473 int32
	_ = v473
	var v475 int32
	_ = v475
	var v476 int32
	_ = v476
	var v477 int32
	_ = v477
	var v478 int32
	_ = v478
	var v479 int32
	_ = v479
	var v480 int32
	_ = v480
	var v481 int32
	_ = v481
	var v482 int32
	_ = v482
	var v485 int32
	_ = v485
	var v486 int32
	_ = v486
	var v487 int32
	_ = v487
	var v488 int32
	_ = v488
	var v490 int32
	_ = v490
	var v510 int32
	_ = v510
	var v511 int32
	_ = v511
	var v513 int32
	_ = v513
	var v517 int32
	_ = v517
	var v537 int32
	_ = v537
	v2 = int32(0)
	v18 = m.G0
	v20 = v18 - int32(32)
	m.G0 = v20
	*(*int32)(unsafe.Add(mBase, uint32(v20)+28)) = v2
	*(*int32)(unsafe.Add(mBase, uint32(v20)+24)) = v2
	*(*int32)(unsafe.Add(mBase, uint32(v20)+20)) = v2
	*(*int32)(unsafe.Add(mBase, uint32(v20)+16)) = v2
	v31 = int64(-1)
	*(*int64)(unsafe.Add(mBase, uint32(v20)+8)) = v31
	*(*int64)(unsafe.Add(mBase, uint32(v20))) = v31
	v36 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v36 < int32(2) {
		v138 = v2
		v139 = v31
		goto L4
	} else {
		goto L5
	}
L1:
	;
	m.G0 = v20 + int32(32)
	return
L2:
	;
	v177 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v178 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v179 = *(*int32)(unsafe.Add(mBase, uint32(v178)+4))
	v180 = F_lookupKeyRead(m, v177, v179)
	mBase = m.M
	v181 = m.ExcPending
	if v181 != 0 {
		goto L22
	} else {
		goto L31
	}
L3:
	;
	F_addReplyError(m, l0, int32(_a1492))
	mBase = m.M
	v176 = m.ExcPending
	if v176 != 0 {
		goto L22
	} else {
		goto L30
	}
L4:
	;
	v153 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v139 == base.I64_extend_i32_s(v153-v138) {
		goto L2
	} else {
		goto L29
	}
L5:
	;
	v42 = int32(0)
	goto L6
L6:
	;
	v57 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v57+v42<<(uint(int32(2))%32))))
	v62 = F_objectGetVal(m, v61)
	mBase = m.M
	v63 = int32(_a1493)
	v66 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v62))))
	if v66 != 0 {
		goto L11
	} else {
		goto L12
	}
L7:
	;
	v138 = v131
	v139 = v31
	goto L4
L8:
	;
	v131 = v42 + int32(1)
	v132 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v131 < v132+int32(-1) {
		v42 = v131
		goto L6
	} else {
		goto L28
	}
L9:
	;
	if v98-v100 != 0 {
		goto L8
	} else {
		goto L21
	}
L10:
	;
	v98 = F_tolower(m, v94)
	mBase = m.M
	v99 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v95))))
	v100 = F_tolower(m, v99)
	mBase = m.M
	goto L9
L11:
	;
	v68 = v62
	v69 = v63
	v70 = v66
	goto L14
L12:
	;
	v94 = int32(0)
	v95 = v63
	goto L10
L13:
	;
	v94 = v91 & int32(255)
	v95 = v90
	goto L10
L14:
	;
	v72 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v69))))
	if v72 == int32(0) {
		v90 = v69
		v91 = v70
		goto L13
	} else {
		goto L16
	}
L15:
	;
	v90 = v84
	v91 = int32(0)
	goto L13
L16:
	;
	v76 = v70 & int32(255)
	if v76 == v72 {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	v83 = int32(1)
	v84 = v69 + v83
	v85 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v68)+1)))
	if v85 != 0 {
		v68 = v68 + v83
		v69 = v84
		v70 = v85
		goto L14
	} else {
		goto L20
	}
L18:
	;
	v78 = F_tolower(m, v76)
	mBase = m.M
	v79 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v69))))
	v80 = F_tolower(m, v79)
	mBase = m.M
	if v78 == v80 {
		goto L17
	} else {
		goto L19
	}
L19:
	;
	v82 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v68))))
	v90 = v69
	v91 = v82
	goto L13
L20:
	;
	goto L15
L21:
	;
	v102 = int32(2)
	v113 = F_parseExtendedCommandArgumentsOrReply(m, l0, v102, v102, v42, v20+int32(16), v20+int32(20), int32(0), v20+int32(28), v20+int32(24))
	mBase = m.M
	v114 = m.ExcPending
	if v114 != 0 {
		goto L22
	} else {
		goto L23
	}
L22:
	;
	return
L23:
	;
	if v113 != 0 {
		goto L1
	} else {
		goto L24
	}
L24:
	;
	v115 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v119 = *(*int32)(unsafe.Add(mBase, uint32(v115+v42<<(uint(int32(2))%32))+4))
	v123 = F_getLongLongFromObjectOrReply(m, l0, v119, v20+int32(8), int32(0))
	mBase = m.M
	v124 = m.ExcPending
	if v124 != 0 {
		goto L22
	} else {
		goto L25
	}
L25:
	;
	if v123 != 0 {
		goto L1
	} else {
		goto L26
	}
L26:
	;
	v125 = *(*int64)(unsafe.Add(mBase, uint32(v20)+8))
	if v125 == int64(0) {
		goto L3
	} else {
		goto L27
	}
L27:
	;
	v138 = v42 + int32(2)
	v139 = v125
	goto L4
L28:
	;
	goto L7
L29:
	;
	goto L3
L30:
	;
	goto L1
L31:
	;
	v183 = F_checkType(m, l0, v180, int32(4))
	mBase = m.M
	v184 = m.ExcPending
	if v184 != 0 {
		goto L22
	} else {
		goto L32
	}
L32:
	;
	if v183 != 0 {
		goto L1
	} else {
		goto L33
	}
L33:
	;
	v185 = F_hashTypeHasVolatileFields(m, v180)
	mBase = m.M
	v186 = m.ExcPending
	if v186 != 0 {
		goto L22
	} else {
		goto L34
	}
L34:
	;
	v187 = int32(0)
	v188 = int32(1)
	v189 = *(*int32)(unsafe.Add(mBase, uint32(v20)+16))
	v191 = v189 & int32(256)
	if v191 == v187 {
		goto L37
	} else {
		goto L38
	}
L35:
	;
	F_initDeferredReplyBuffer(m, l0)
	mBase = m.M
	v260 = m.ExcPending
	if v260 != 0 {
		goto L22
	} else {
		goto L55
	}
L36:
	;
	v251 = int32(0)
	v254 = v246
	v255 = v247
	v256 = v251
	v257 = v251
	goto L35
L37:
	;
	v195 = int32(0)
	v198 = *(*int32)(unsafe.Add(mBase, uint32(v20)+28))
	if v198 == v195 {
		v254 = v188
		v255 = v195
		v256 = v195
		v257 = v195
		goto L35
	} else {
		goto L39
	}
L38:
	;
	v246 = v188
	v247 = int32(0)
	goto L36
L39:
	;
	if v189&int32(192) != 0 {
		v207 = int64(0)
		v208 = v198
		goto L40
	} else {
		goto L41
	}
L40:
	;
	v209 = *(*int32)(unsafe.Add(mBase, uint32(v20)+20))
	v210 = F_convertExpireArgumentToUnixTime(m, l0, v208, v207, v209, v20)
	mBase = m.M
	v211 = m.ExcPending
	if v211 != 0 {
		goto L22
	} else {
		goto L43
	}
L41:
	;
	v205 = *(*int64)(unsafe.Add(mBase, _consts[78]))
	goto L42
L42:
	;
	v206 = *(*int32)(unsafe.Add(mBase, uint32(v20)+28))
	v207 = v205
	v208 = v206
	goto L40
L43:
	;
	if v210 == int32(-1) {
		goto L1
	} else {
		goto L44
	}
L44:
	;
	v214 = *(*int64)(unsafe.Add(mBase, uint32(v20)))
	v218 = *(*int32)(unsafe.Add(mBase, _consts[88]))
	if v218 == int32(0) {
		goto L48
	} else {
		goto L49
	}
L45:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v20))) = int64(0)
	v241 = int32(1)
	v254 = v241
	v255 = int32(0)
	v256 = v241
	v257 = v241
	goto L35
L46:
	;
	if v236 != 0 {
		goto L45
	} else {
		goto L54
	}
L47:
	;
	goto L46
L48:
	;
	v224 = int32(0)
	v225 = F_commandTimeSnapshot(m)
	mBase = m.M
	if v225 < v214 {
		v236 = v224
		goto L47
	} else {
		goto L51
	}
L49:
	;
	v222 = *(*int32)(unsafe.Add(mBase, uint32(v218)+216))
	if v222 != 0 {
		v236 = int32(0)
		goto L47
	} else {
		goto L50
	}
L50:
	;
	goto L48
L51:
	;
	v227 = int32(_a69)
	v228 = *(*int32)(unsafe.Add(mBase, _consts[64]))
	v230 = *(*int32)(unsafe.Add(mBase, _consts[116]))
	if v230 != 0 {
		v236 = v224
		goto L47
	} else {
		goto L52
	}
L52:
	;
	if v228 != 0 {
		v236 = v224
		goto L47
	} else {
		goto L53
	}
L53:
	;
	v232 = *(*int32)(unsafe.Add(mBase, _consts[218]))
	v236 = base.B2i32(v232 == int32(0))
	goto L47
L54:
	;
	v246 = int32(0)
	v247 = int32(1)
	goto L36
L55:
	;
	v261 = *(*int32)(unsafe.Add(mBase, uint32(v20)+8))
	F_addReplyArrayLen(m, l0, v261)
	mBase = m.M
	v263 = m.ExcPending
	if v263 != 0 {
		goto L22
	} else {
		goto L56
	}
L56:
	;
	v265 = v191 | v257 | v255
	if v265 != 0 {
		goto L58
	} else {
		goto L59
	}
L57:
	;
	v322 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v322 <= v138 {
		goto L72
	} else {
		goto L73
	}
L58:
	;
	v269 = int32(2)
	v270 = *(*int32)(unsafe.Add(mBase, uint32(v20)+8))
	v275 = F_valkey_malloc(m, v270<<(uint(v269)%32)+int32(20))
	mBase = m.M
	v276 = m.ExcPending
	if v276 != 0 {
		goto L22
	} else {
		goto L60
	}
L59:
	;
	v266 = int32(-1)
	v317 = v187
	v319 = int32(0)
	v320 = v266
	v321 = v266
	goto L57
L60:
	;
	if v191 != 0 {
		goto L61
	} else {
		goto L62
	}
L61:
	;
	v281 = int32(320)
	goto L63
L62:
	;
	v281 = int32(316)
	goto L63
L63:
	;
	if v256 != 0 {
		goto L64
	} else {
		goto L65
	}
L64:
	;
	v282 = int32(312)
	goto L66
L65:
	;
	v282 = v281
	goto L66
L66:
	;
	v284 = *(*int32)(unsafe.Add(mBase, uint32(v282)+uint32(_consts[77])))
	*(*int32)(unsafe.Add(mBase, uint32(v275))) = v284
	v286 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v287 = *(*int32)(unsafe.Add(mBase, uint32(v286)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v275)+4)) = v287
	F_incrRefCount(m, v287)
	mBase = m.M
	v290 = m.ExcPending
	if v290 != 0 {
		goto L22
	} else {
		goto L67
	}
L67:
	;
	v291 = int32(-1)
	if v254 != 0 {
		v297 = v269
		v298 = v291
		goto L68
	} else {
		goto L69
	}
L68:
	;
	if v255|v191 == int32(0) {
		v317 = v275
		v319 = v297
		v320 = v298
		v321 = v291
		goto L57
	} else {
		goto L70
	}
L69:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v275)+8)) = int32(0)
	v297 = int32(3)
	v298 = int32(2)
	goto L68
L70:
	;
	v302 = int32(2)
	v306 = *(*int32)(unsafe.Add(mBase, _consts[773]))
	*(*int32)(unsafe.Add(mBase, uint32(v275+v297<<(uint(v302)%32)))) = v306
	v309 = v297 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v275+v309<<(uint(v302)%32)))) = int32(0)
	v317 = v275
	v319 = v297 + v302
	v320 = v298
	v321 = v309
	goto L57
L71:
	;
	F_commitDeferredReplyBuffer(m, l0, int32(1))
	mBase = m.M
	v537 = m.ExcPending
	if v537 != 0 {
		goto L22
	} else {
		goto L129
	}
L72:
	;
	if v265 == int32(0) {
		goto L124
	} else {
		goto L125
	}
L73:
	;
	v324 = int32(0)
	v330 = v138
	v340 = v319
	v341 = v324
	goto L74
L74:
	;
	v345 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v347 = v330 << (uint(int32(2)) % 32)
	v349 = *(*int32)(unsafe.Add(mBase, uint32(v345+v347)))
	v350 = F_objectGetVal(m, v349)
	mBase = m.M
	F_addHashFieldToReply(m, l0, v180, v350)
	mBase = m.M
	v352 = m.ExcPending
	if v352 != 0 {
		goto L22
	} else {
		goto L76
	}
L75:
	;
	if v404 == int32(0) {
		goto L72
	} else {
		goto L92
	}
L76:
	;
	if base.B2i32(v180 != v324)&v256 == int32(0) {
		goto L79
	} else {
		goto L80
	}
L77:
	;
	v406 = v330 + int32(1)
	v407 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v406 < v407 {
		v330 = v406
		v340 = v403
		v341 = v404
		goto L74
	} else {
		goto L91
	}
L78:
	;
	v392 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v394 = *(*int32)(unsafe.Add(mBase, uint32(v392+v347)))
	*(*int32)(unsafe.Add(mBase, uint32(v317+v340<<(uint(int32(2))%32)))) = v394
	F_incrRefCount(m, v394)
	mBase = m.M
	v397 = m.ExcPending
	if v397 != 0 {
		goto L22
	} else {
		goto L90
	}
L79:
	;
	if v254 != 0 {
		goto L83
	} else {
		goto L84
	}
L80:
	;
	v355 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v357 = *(*int32)(unsafe.Add(mBase, uint32(v355+v347)))
	v358 = F_objectGetVal(m, v357)
	mBase = m.M
	v359 = F_hashTypeDelete(m, v180, v358)
	mBase = m.M
	v360 = m.ExcPending
	if v360 != 0 {
		goto L22
	} else {
		goto L81
	}
L81:
	;
	if v359 == int32(0) {
		v403 = v340
		v404 = v341
		goto L77
	} else {
		goto L82
	}
L82:
	;
	v363 = int32(_a69)
	v365 = *(*int64)(unsafe.Add(mBase, _consts[971]))
	*(*int64)(unsafe.Add(mBase, _consts[971])) = v365 + int64(1)
	goto L78
L83:
	;
	if v191 == int32(0) {
		v403 = v340
		v404 = v341
		goto L77
	} else {
		goto L87
	}
L84:
	;
	v369 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v371 = *(*int32)(unsafe.Add(mBase, uint32(v369+v347)))
	v372 = F_objectGetVal(m, v371)
	mBase = m.M
	v373 = *(*int64)(unsafe.Add(mBase, uint32(v20)))
	v375 = F_hashTypeSetExpire(m, v180, v372, v373, int32(0))
	mBase = m.M
	v376 = m.ExcPending
	if v376 != 0 {
		goto L22
	} else {
		goto L85
	}
L85:
	;
	if v375 == int32(1) {
		goto L78
	} else {
		goto L86
	}
L86:
	;
	v403 = v340
	v404 = v341
	goto L77
L87:
	;
	v381 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v383 = *(*int32)(unsafe.Add(mBase, uint32(v381+v347)))
	v384 = F_objectGetVal(m, v383)
	mBase = m.M
	v385 = F_hashTypePersist(m, v180, v384)
	mBase = m.M
	v386 = m.ExcPending
	if v386 != 0 {
		goto L22
	} else {
		goto L88
	}
L88:
	;
	if v385 != int32(1) {
		v403 = v340
		v404 = v341
		goto L77
	} else {
		goto L89
	}
L89:
	;
	goto L78
L90:
	;
	v398 = int32(1)
	v403 = v340 + v398
	v404 = v341 + v398
	goto L77
L91:
	;
	goto L75
L92:
	;
	v411 = int32(0)
	if v320 < int32(1) {
		v424 = v411
		goto L93
	} else {
		goto L94
	}
L93:
	;
	if v321 < int32(1) {
		v436 = v411
		goto L97
	} else {
		goto L98
	}
L94:
	;
	v418 = *(*int64)(unsafe.Add(mBase, uint32(v20)))
	v419 = F_createStringObjectFromLongLong(m, v418)
	mBase = m.M
	v420 = m.ExcPending
	if v420 != 0 {
		goto L22
	} else {
		goto L95
	}
L95:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v317+v320<<(uint(int32(2))%32)))) = v419
	F_incrRefCount(m, v419)
	mBase = m.M
	v423 = m.ExcPending
	if v423 != 0 {
		goto L22
	} else {
		goto L96
	}
L96:
	;
	v424 = v419
	goto L93
L97:
	;
	F_replaceClientCommandVector(m, l0, v403, v317)
	mBase = m.M
	v438 = m.ExcPending
	if v438 != 0 {
		goto L22
	} else {
		goto L101
	}
L98:
	;
	v431 = F_createStringObjectFromLongLong(m, base.I64_extend_i32_s(v404))
	mBase = m.M
	v432 = m.ExcPending
	if v432 != 0 {
		goto L22
	} else {
		goto L99
	}
L99:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v317+v321<<(uint(int32(2))%32)))) = v431
	F_incrRefCount(m, v431)
	mBase = m.M
	v435 = m.ExcPending
	if v435 != 0 {
		goto L22
	} else {
		goto L100
	}
L100:
	;
	v436 = v431
	goto L97
L101:
	;
	if v254 != 0 {
		goto L102
	} else {
		goto L103
	}
L102:
	;
	v443 = int32(_a1494)
	goto L104
L103:
	;
	v443 = int32(_a1495)
	goto L104
L104:
	;
	if v256 != 0 {
		goto L105
	} else {
		goto L106
	}
L105:
	;
	v444 = int32(_a1491)
	goto L107
L106:
	;
	v444 = v443
	goto L107
L107:
	;
	v445 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v446 = *(*int32)(unsafe.Add(mBase, uint32(v445)+4))
	v447 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v448 = *(*int32)(unsafe.Add(mBase, uint32(v447)+28))
	F_notifyKeyspaceEvent(m, int32(64), v444, v446, v448)
	mBase = m.M
	v450 = m.ExcPending
	if v450 != 0 {
		goto L22
	} else {
		goto L108
	}
L108:
	;
	if v424 == int32(0) {
		goto L109
	} else {
		goto L110
	}
L109:
	;
	if v436 == int32(0) {
		goto L112
	} else {
		goto L113
	}
L110:
	;
	F_decrRefCount(m, v424)
	mBase = m.M
	v454 = m.ExcPending
	if v454 != 0 {
		goto L22
	} else {
		goto L111
	}
L111:
	;
	goto L109
L112:
	;
	v459 = int32(_a69)
	v461 = *(*int64)(unsafe.Add(mBase, _consts[60]))
	*(*int64)(unsafe.Add(mBase, _consts[60])) = v461 + base.I64_extend_i32_s(v404)
	v465 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v466 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v467 = *(*int32)(unsafe.Add(mBase, uint32(v466)+4))
	F_signalModifiedKey(m, l0, v465, v467)
	mBase = m.M
	v469 = m.ExcPending
	if v469 != 0 {
		goto L22
	} else {
		goto L115
	}
L113:
	;
	F_decrRefCount(m, v436)
	mBase = m.M
	v458 = m.ExcPending
	if v458 != 0 {
		goto L22
	} else {
		goto L114
	}
L114:
	;
	goto L112
L115:
	;
	v470 = F_hashTypeHasVolatileFields(m, v180)
	mBase = m.M
	v471 = m.ExcPending
	if v471 != 0 {
		goto L22
	} else {
		goto L117
	}
L116:
	;
	v476 = F_hashTypeLength(m, v180)
	mBase = m.M
	v477 = m.ExcPending
	if v477 != 0 {
		goto L22
	} else {
		goto L120
	}
L117:
	;
	if v185 == v470 {
		goto L116
	} else {
		goto L118
	}
L118:
	;
	v473 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	F_dbUpdateObjectWithVolatileItemsTracking(m, v473, v180)
	mBase = m.M
	v475 = m.ExcPending
	if v475 != 0 {
		goto L22
	} else {
		goto L119
	}
L119:
	;
	goto L116
L120:
	;
	if v476 != 0 {
		goto L71
	} else {
		goto L121
	}
L121:
	;
	v478 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v479 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v480 = *(*int32)(unsafe.Add(mBase, uint32(v479)+4))
	v481 = F_dbDelete(m, v478, v480)
	mBase = m.M
	v482 = m.ExcPending
	if v482 != 0 {
		goto L22
	} else {
		goto L122
	}
L122:
	;
	v485 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v486 = *(*int32)(unsafe.Add(mBase, uint32(v485)+4))
	v487 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v488 = *(*int32)(unsafe.Add(mBase, uint32(v487)+28))
	F_notifyKeyspaceEvent(m, int32(4), int32(_a132), v486, v488)
	mBase = m.M
	v490 = m.ExcPending
	if v490 != 0 {
		goto L22
	} else {
		goto L123
	}
L123:
	;
	goto L71
L124:
	;
	if v317 == int32(0) {
		goto L71
	} else {
		goto L127
	}
L125:
	;
	v510 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v511 = *(*int32)(unsafe.Add(mBase, uint32(v510)+4))
	F_decrRefCount(m, v511)
	mBase = m.M
	v513 = m.ExcPending
	if v513 != 0 {
		goto L22
	} else {
		goto L126
	}
L126:
	;
	goto L124
L127:
	;
	F_valkey_free(m, v317)
	mBase = m.M
	v517 = m.ExcPending
	if v517 != 0 {
		goto L22
	} else {
		goto L128
	}
L128:
	;
	goto L71
L129:
	;
	goto L1
}
func F_hllDenseRegHisto(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
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
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v225 int32
	_ = v225
	v18 = l0
	v20 = int32(0)
	for {
		v34 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18)+11)))
		v35 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18)+10)))
		v36 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18)+9)))
		v37 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18)+8)))
		v38 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18)+7)))
		v39 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18)+6)))
		v40 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18)+5)))
		v41 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18)+4)))
		v42 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18)+3)))
		v43 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18)+2)))
		v44 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18)+1)))
		v45 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18))))
		v46 = int32(63)
		v48 = int32(2)
		v50 = l1 + v45&v46<<(uint(v48)%32)
		v51 = *(*int32)(unsafe.Add(mBase, uint32(v50)))
		v52 = int32(1)
		*(*int32)(unsafe.Add(mBase, uint32(v50))) = v51 + v52
		v57 = int32(60)
		v59 = int32(6)
		v64 = l1 + (v44<<(uint(v48)%32)&v57|int32(base.Ui32(v45)>>(uint(v59)%32)))<<(uint(v48)%32)
		v65 = *(*int32)(unsafe.Add(mBase, uint32(v64)))
		*(*int32)(unsafe.Add(mBase, uint32(v64))) = v65 + v52
		v69 = int32(4)
		v71 = int32(48)
		v78 = l1 + (v43<<(uint(v69)%32)&v71|int32(base.Ui32(v44)>>(uint(v69)%32)))<<(uint(v48)%32)
		v79 = *(*int32)(unsafe.Add(mBase, uint32(v78)))
		*(*int32)(unsafe.Add(mBase, uint32(v78))) = v79 + v52
		v83 = int32(252)
		v85 = l1 + v43&v83
		v86 = *(*int32)(unsafe.Add(mBase, uint32(v85)))
		*(*int32)(unsafe.Add(mBase, uint32(v85))) = v86 + v52
		v94 = l1 + v42&v46<<(uint(v48)%32)
		v95 = *(*int32)(unsafe.Add(mBase, uint32(v94)))
		*(*int32)(unsafe.Add(mBase, uint32(v94))) = v95 + v52
		v108 = l1 + (v41<<(uint(v48)%32)&v57|int32(base.Ui32(v42)>>(uint(v59)%32)))<<(uint(v48)%32)
		v109 = *(*int32)(unsafe.Add(mBase, uint32(v108)))
		*(*int32)(unsafe.Add(mBase, uint32(v108))) = v109 + v52
		v122 = l1 + (v40<<(uint(v69)%32)&v71|int32(base.Ui32(v41)>>(uint(v69)%32)))<<(uint(v48)%32)
		v123 = *(*int32)(unsafe.Add(mBase, uint32(v122)))
		*(*int32)(unsafe.Add(mBase, uint32(v122))) = v123 + v52
		v129 = l1 + v40&v83
		v130 = *(*int32)(unsafe.Add(mBase, uint32(v129)))
		*(*int32)(unsafe.Add(mBase, uint32(v129))) = v130 + v52
		v138 = l1 + v39&v46<<(uint(v48)%32)
		v139 = *(*int32)(unsafe.Add(mBase, uint32(v138)))
		*(*int32)(unsafe.Add(mBase, uint32(v138))) = v139 + v52
		v152 = l1 + (v38<<(uint(v48)%32)&v57|int32(base.Ui32(v39)>>(uint(v59)%32)))<<(uint(v48)%32)
		v153 = *(*int32)(unsafe.Add(mBase, uint32(v152)))
		*(*int32)(unsafe.Add(mBase, uint32(v152))) = v153 + v52
		v166 = l1 + (v37<<(uint(v69)%32)&v71|int32(base.Ui32(v38)>>(uint(v69)%32)))<<(uint(v48)%32)
		v167 = *(*int32)(unsafe.Add(mBase, uint32(v166)))
		*(*int32)(unsafe.Add(mBase, uint32(v166))) = v167 + v52
		v173 = l1 + v37&v83
		v174 = *(*int32)(unsafe.Add(mBase, uint32(v173)))
		*(*int32)(unsafe.Add(mBase, uint32(v173))) = v174 + v52
		v182 = l1 + v36&v46<<(uint(v48)%32)
		v183 = *(*int32)(unsafe.Add(mBase, uint32(v182)))
		*(*int32)(unsafe.Add(mBase, uint32(v182))) = v183 + v52
		v196 = l1 + (v35<<(uint(v48)%32)&v57|int32(base.Ui32(v36)>>(uint(v59)%32)))<<(uint(v48)%32)
		v197 = *(*int32)(unsafe.Add(mBase, uint32(v196)))
		*(*int32)(unsafe.Add(mBase, uint32(v196))) = v197 + v52
		v210 = l1 + (v34<<(uint(v69)%32)&v71|int32(base.Ui32(v35)>>(uint(v69)%32)))<<(uint(v48)%32)
		v211 = *(*int32)(unsafe.Add(mBase, uint32(v210)))
		*(*int32)(unsafe.Add(mBase, uint32(v210))) = v211 + v52
		v217 = l1 + v34&v83
		v218 = *(*int32)(unsafe.Add(mBase, uint32(v217)))
		*(*int32)(unsafe.Add(mBase, uint32(v217))) = v218 + v52
		v225 = v20 + v52
		if v225 != int32(1024) {
			v18 = v18 + int32(12)
			v20 = v225
			continue
		} else {
			break
		}
		break
	}
	return
}
func F_hllMerge(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
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
	var v34 int32
	_ = v34
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	var v96 int32
	_ = v96
	var v105 int32
	_ = v105
	var v108 int32
	_ = v108
	var v112 int32
	_ = v112
	var v121 int32
	_ = v121
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v128 int32
	_ = v128
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v145 int32
	_ = v145
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v160 int32
	_ = v160
	var v180 int32
	_ = v180
	v11 = F_objectGetVal(m, l1)
	mBase = m.M
	v12 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+4)))
	if v12 != 0 {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	return v180
L2:
	;
	v180 = int32(0)
	goto L1
L3:
	;
	v49 = int32(-1)
	v50 = F_objectGetVal(m, l1)
	mBase = m.M
	v51 = F_objectGetVal(m, l1)
	mBase = m.M
	v54 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v51+v49))))
	switch v54 & int32(7) {
	case 0:
		goto L15
	case 1:
		goto L14
	case 2:
		goto L13
	case 3:
		goto L12
	case 4:
		goto L11
	default:
		v180 = v49
		goto L1
	}
L4:
	;
	v17 = int32(0)
	goto L5
L5:
	;
	v26 = int32(6)
	v27 = v17 * v26
	v30 = v11 + int32(16) + int32(base.Ui32(v27)>>(uint(int32(3))%32))
	v31 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v30)+1)))
	v34 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v30))))
	v40 = int32(base.Ui32(v31<<(uint(int32(8))%32)|v34)>>(uint(v27&v26)%32)) & int32(63)
	v41 = l0 + v17
	v42 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v41))))
	if base.Ui32(v40) <= base.Ui32(v42) {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	v46 = v17 + int32(1)
	if v46 != int32(16384) {
		v17 = v46
		goto L5
	} else {
		goto L9
	}
L8:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v41))) = uint8(v40)
	goto L7
L9:
	;
	goto L2
L10:
	;
	if base.Ui32(v71) < base.Ui32(int32(17)) {
		v180 = v49
		goto L1
	} else {
		goto L16
	}
L11:
	;
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v51+int32(-17))))
	v71 = v70
	goto L10
L12:
	;
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v51+int32(-9))))
	v71 = v67
	goto L10
L13:
	;
	v64 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v51+int32(-5)))))
	v71 = v64
	goto L10
L14:
	;
	v61 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v51+int32(-3)))))
	v71 = v61
	goto L10
L15:
	;
	v71 = int32(base.Ui32(v54) >> (uint(int32(3)) % 32))
	goto L10
L16:
	;
	v79 = v50 + int32(16)
	v80 = int32(0)
	goto L17
L17:
	;
	v88 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v79))))
	v90 = v88 & int32(192)
	if v90 == int32(64) {
		goto L21
	} else {
		goto L22
	}
L18:
	;
	if v154 != int32(16384) {
		v180 = v49
		goto L1
	} else {
		goto L42
	}
L19:
	;
	v160 = v79 + v155
	if base.Ui32(v160) < base.Ui32(v50+v71) {
		v79 = v160
		v80 = v154
		goto L17
	} else {
		goto L41
	}
L20:
	;
	v112 = v88 & int32(3)
	if int32(16384) < v112+v80+int32(1) {
		v180 = v49
		goto L1
	} else {
		goto L26
	}
L21:
	;
	v105 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v79)+1)))
	v108 = v80 + v88<<(uint(int32(8))%32)&int32(16128) + v105 + int32(1)
	if v108 <= int32(16384) {
		v154 = v108
		v155 = int32(2)
		goto L19
	} else {
		goto L25
	}
L22:
	;
	if v90 != 0 {
		goto L20
	} else {
		goto L23
	}
L23:
	;
	v93 = int32(1)
	v96 = v80 + v88 + v93
	if v96 <= int32(16384) {
		v154 = v96
		v155 = v93
		goto L19
	} else {
		goto L24
	}
L24:
	;
	v180 = v49
	goto L1
L25:
	;
	v180 = v49
	goto L1
L26:
	;
	v121 = int32(base.Ui32(v88)>>(uint(int32(2))%32)) & int32(31)
	v123 = v121 + int32(1)
	v124 = l0 + v80
	v125 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v124))))
	if base.Ui32(v121) < base.Ui32(v125) {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	v128 = int32(1)
	v130 = v80 + v128
	if v112 != 0 {
		goto L29
	} else {
		goto L30
	}
L28:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v124))) = uint8(v123)
	goto L27
L29:
	;
	v131 = l0 + v130
	v132 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v131))))
	if base.Ui32(v121) < base.Ui32(v132) {
		goto L31
	} else {
		goto L32
	}
L30:
	;
	v154 = v130
	v155 = v128
	goto L19
L31:
	;
	v136 = v80 + int32(2)
	v137 = int32(1)
	if v112 != v137 {
		goto L33
	} else {
		goto L34
	}
L32:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v131))) = uint8(v123)
	goto L31
L33:
	;
	v140 = l0 + v136
	v141 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v140))))
	if base.Ui32(v121) < base.Ui32(v141) {
		goto L35
	} else {
		goto L36
	}
L34:
	;
	v154 = v136
	v155 = v137
	goto L19
L35:
	;
	v145 = v80 + int32(3)
	if v112 != int32(2) {
		goto L37
	} else {
		goto L38
	}
L36:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v140))) = uint8(v123)
	goto L35
L37:
	;
	v148 = l0 + v145
	v149 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v148))))
	if base.Ui32(v121) < base.Ui32(v149) {
		goto L39
	} else {
		goto L40
	}
L38:
	;
	v154 = v145
	v155 = v137
	goto L19
L39:
	;
	v154 = v80 + int32(4)
	v155 = v137
	goto L19
L40:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v148))) = uint8(v123)
	goto L39
L41:
	;
	goto L18
L42:
	;
	goto L2
}
func F_hpexpireCommand(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v3 int64
	_ = v3
	var v6 int32
	_ = v6
	v3 = *(*int64)(unsafe.Add(mBase, _consts[78]))
	F_hexpireGenericCommand(m, l0, v3, int32(1))
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return
	} else {
		return
	}
}
func F_hpttlCommand(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v3 int64
	_ = v3
	var v6 int32
	_ = v6
	v3 = *(*int64)(unsafe.Add(mBase, _consts[78]))
	F_httlGenericCommand(m, l0, v3, int32(1))
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return
	} else {
		return
	}
}
func F_hrandfieldReplyWithListpack(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v16 int32
	_ = v16
	var v22 int32
	_ = v22
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v37 int64
	_ = v37
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v49 int64
	_ = v49
	var v51 int32
	_ = v51
	var v55 int32
	_ = v55
	if l1 == int32(0) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v16 = int32(0)
	goto L3
L3:
	;
	if l3 == int32(0) {
		goto L5
	} else {
		goto L6
	}
L4:
	;
	goto L1
L5:
	;
	v29 = v16 << (uint(int32(4)) % 32)
	v30 = l2 + v29
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v30)))
	if v31 == int32(0) {
		goto L11
	} else {
		goto L12
	}
L6:
	;
	v22 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+224)))
	if base.Ui32(v22) < base.Ui32(int32(3)) {
		goto L5
	} else {
		goto L7
	}
L7:
	;
	F_addWritePreparedReplyArrayLen(m, l0, int32(2))
	mBase = m.M
	v27 = m.ExcPending
	if v27 != 0 {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	return
L9:
	;
	goto L5
L10:
	;
	if l3 == int32(0) {
		goto L15
	} else {
		goto L16
	}
L11:
	;
	v37 = *(*int64)(unsafe.Add(mBase, uint32(v30)+8))
	F_addWritePreparedReplyBulkLongLong(m, l0, v37)
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L8
	} else {
		goto L14
	}
L12:
	;
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v30)+4))
	F_addWritePreparedReplyBulkCBuffer(m, l0, v31, v34)
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
		goto L8
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
	v55 = v16 + int32(1)
	if v55 != l1 {
		v16 = v55
		goto L3
	} else {
		goto L21
	}
L16:
	;
	v42 = l3 + v29
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v42)))
	if v43 == int32(0) {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	v49 = *(*int64)(unsafe.Add(mBase, uint32(v42)+8))
	F_addWritePreparedReplyBulkLongLong(m, l0, v49)
	mBase = m.M
	v51 = m.ExcPending
	if v51 != 0 {
		goto L8
	} else {
		goto L20
	}
L18:
	;
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v42)+4))
	F_addWritePreparedReplyBulkCBuffer(m, l0, v43, v46)
	mBase = m.M
	v48 = m.ExcPending
	if v48 != 0 {
		goto L8
	} else {
		goto L19
	}
L19:
	;
	goto L15
L20:
	;
	goto L15
L21:
	;
	goto L4
}
func F_hsetCommand(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
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
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
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
	var v59 int32
	_ = v59
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v101 int64
	_ = v101
	var v102 int32
	_ = v102
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v115 int32
	_ = v115
	var v118 int32
	_ = v118
	var v120 int64
	_ = v120
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
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
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v142 int64
	_ = v142
	var v143 int32
	_ = v143
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
	var v160 int32
	_ = v160
	var v162 int32
	_ = v162
	var v164 int32
	_ = v164
	v10 = m.G0
	v12 = v10 - int32(16)
	m.G0 = v12
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v14&int32(-2147483647) != int32(1) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v12 + int32(16)
	return
L2:
	;
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v23 = *(*int32)(unsafe.Add(mBase, uint32(v22)+4))
	v24 = F_lookupKeyWrite(m, v21, v23)
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		goto L4
	} else {
		goto L6
	}
L3:
	;
	F_addReplyErrorArity(m, l0)
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
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
	v27 = F_checkType(m, l0, v24, int32(4))
	mBase = m.M
	v28 = m.ExcPending
	if v28 != 0 {
		goto L4
	} else {
		goto L7
	}
L7:
	;
	if v27 != 0 {
		goto L1
	} else {
		goto L8
	}
L8:
	;
	if v24 != 0 {
		v40 = v24
		goto L9
	} else {
		goto L10
	}
L9:
	;
	v41 = int32(2)
	v42 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v44 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	F_hashTypeTryConversion(m, v40, v42, v41, v44+int32(-1))
	mBase = m.M
	v48 = m.ExcPending
	if v48 != 0 {
		goto L4
	} else {
		goto L14
	}
L10:
	;
	v29 = F_createHashObject(m)
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		goto L4
	} else {
		goto L11
	}
L11:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+12)) = v29
	v32 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	F_dbAdd(m, v32, v23, v12+int32(12))
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
		goto L4
	} else {
		goto L12
	}
L12:
	;
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v12)+12))
	if v37 == int32(0) {
		goto L1
	} else {
		goto L13
	}
L13:
	;
	v40 = v37
	goto L9
L14:
	;
	v49 = F_hashTypeHasVolatileFields(m, v40)
	mBase = m.M
	v50 = m.ExcPending
	if v50 != 0 {
		goto L4
	} else {
		goto L15
	}
L15:
	;
	v52 = int32(0)
	v53 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v53 < int32(3) {
		v101 = int64(0)
		v102 = v52
		goto L16
	} else {
		goto L17
	}
L16:
	;
	v105 = F_hashTypeHasVolatileFields(m, v40)
	mBase = m.M
	v106 = m.ExcPending
	if v106 != 0 {
		goto L4
	} else {
		goto L23
	}
L17:
	;
	v59 = v41
	v63 = v52
	v64 = int32(0)
	goto L18
L18:
	;
	v66 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v12)+11)) = uint8(v66)
	v68 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v70 = v59 << (uint(int32(2)) % 32)
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v68+v70)))
	v73 = F_objectGetVal(m, v72)
	mBase = m.M
	v74 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v74+v70+int32(4))))
	v79 = F_objectGetVal(m, v78)
	mBase = m.M
	v84 = F_hashTypeSet(m, v40, v73, v79, int64(-1), v66, v12+int32(11))
	mBase = m.M
	v85 = m.ExcPending
	if v85 != 0 {
		goto L4
	} else {
		goto L20
	}
L19:
	;
	v101 = base.I64_extend_i32_u(v88)
	v102 = v90
	goto L16
L20:
	;
	v88 = v84 ^ int32(1) + v64
	v89 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+11)))
	v90 = v63 + v89
	v92 = v59 + int32(2)
	v93 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v92 < v93 {
		v59 = v92
		v63 = v90
		v64 = v88
		goto L18
	} else {
		goto L21
	}
L21:
	;
	goto L19
L22:
	;
	v111 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v112 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v113 = *(*int32)(unsafe.Add(mBase, uint32(v112)+4))
	F_signalModifiedKey(m, l0, v111, v113)
	mBase = m.M
	v115 = m.ExcPending
	if v115 != 0 {
		goto L4
	} else {
		goto L26
	}
L23:
	;
	if v49 == v105 {
		goto L22
	} else {
		goto L24
	}
L24:
	;
	v108 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	F_dbUpdateObjectWithVolatileItemsTracking(m, v108, v40)
	mBase = m.M
	v110 = m.ExcPending
	if v110 != 0 {
		goto L4
	} else {
		goto L25
	}
L25:
	;
	goto L22
L26:
	;
	if v102 == int32(0) {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	v134 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v135 = *(*int32)(unsafe.Add(mBase, uint32(v134)+4))
	v136 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v137 = *(*int32)(unsafe.Add(mBase, uint32(v136)+28))
	F_notifyKeyspaceEvent(m, int32(64), int32(_a1490), v135, v137)
	mBase = m.M
	v139 = m.ExcPending
	if v139 != 0 {
		goto L4
	} else {
		goto L30
	}
L28:
	;
	v118 = int32(_a69)
	v120 = *(*int64)(unsafe.Add(mBase, _consts[971]))
	*(*int64)(unsafe.Add(mBase, _consts[971])) = v120 + base.I64_extend_i32_u(v102)
	v126 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v127 = *(*int32)(unsafe.Add(mBase, uint32(v126)+4))
	v128 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v129 = *(*int32)(unsafe.Add(mBase, uint32(v128)+28))
	F_notifyKeyspaceEvent(m, int32(64), int32(_a1491), v127, v129)
	mBase = m.M
	v131 = m.ExcPending
	if v131 != 0 {
		goto L4
	} else {
		goto L29
	}
L29:
	;
	goto L27
L30:
	;
	v140 = int32(_a69)
	v142 = *(*int64)(unsafe.Add(mBase, _consts[60]))
	v143 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v147 = base.I32_div_s(v143+int32(-2), int32(2))
	*(*int64)(unsafe.Add(mBase, _consts[60])) = v142 + base.I64_extend_i32_s(v147)
	v151 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v152 = *(*int32)(unsafe.Add(mBase, uint32(v151)))
	v153 = F_objectGetVal(m, v152)
	mBase = m.M
	v154 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v153)+1)))
	if v154|int32(32) != int32(115) {
		goto L31
	} else {
		goto L32
	}
L31:
	;
	v162 = *(*int32)(unsafe.Add(mBase, _consts[77]))
	F_addReply(m, l0, v162)
	mBase = m.M
	v164 = m.ExcPending
	if v164 != 0 {
		goto L4
	} else {
		goto L34
	}
L32:
	;
	F_addReplyLongLong(m, l0, v101)
	mBase = m.M
	v160 = m.ExcPending
	if v160 != 0 {
		goto L4
	} else {
		goto L33
	}
L33:
	;
	goto L1
L34:
	;
	goto L1
}
func F_htonl(m *base.Module, l0 int32) int32 {
	var v2 int32
	_ = v2
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	v2 = int32(24)
	v4 = int32(65280)
	v6 = int32(8)
	return l0<<(uint(v2)%32) | l0&v4<<(uint(v6)%32) | (int32(base.Ui32(l0)>>(uint(v6)%32))&v4 | int32(base.Ui32(l0)>>(uint(v2)%32)))
}
func F_htons(m *base.Module, l0 int32) int32 {
	var v2 int32
	_ = v2
	v2 = int32(8)
	return (l0<<(uint(v2)%32) | int32(base.Ui32(l0)>>(uint(v2)%32))) & int32(65535)
}
func F_httlGenericCommand(m *base.Module, l0 int32, l1 int64, l2 int32) {
	mBase := m.M
	_ = mBase
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
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
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
	var v42 int32
	_ = v42
	var v47 int32
	_ = v47
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
	var v62 int32
	_ = v62
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
	var v72 int64
	_ = v72
	var v75 int32
	_ = v75
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v95 int64
	_ = v95
	var v107 int64
	_ = v107
	var v109 int64
	_ = v109
	var v112 int32
	_ = v112
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v123 int64
	_ = v123
	var v124 int64
	_ = v124
	var v127 int64
	_ = v127
	var v128 int64
	_ = v128
	var v131 int64
	_ = v131
	var v136 int64
	_ = v136
	var v137 int64
	_ = v137
	var v140 int32
	_ = v140
	var v141 int64
	_ = v141
	var v143 int64
	_ = v143
	v9 = m.G0
	v11 = v9 - int32(16)
	m.G0 = v11
	*(*int64)(unsafe.Add(mBase, uint32(v11)+8)) = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v11))) = int64(-2)
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v18 = *(*int32)(unsafe.Add(mBase, uint32(v17)+8))
	v19 = F_objectGetVal(m, v18)
	mBase = m.M
	v20 = int32(_a1493)
	v23 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19))))
	if v23 != 0 {
		goto L5
	} else {
		goto L6
	}
L1:
	;
	m.G0 = v11 + int32(16)
	return
L2:
	;
	v65 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v65)+12))
	v70 = F_getLongLongFromObjectOrReply(m, l0, v66, v11+int32(8), int32(0))
	mBase = m.M
	v71 = m.ExcPending
	if v71 != 0 {
		goto L16
	} else {
		goto L18
	}
L3:
	;
	if v55-v57 == int32(0) {
		goto L2
	} else {
		goto L15
	}
L4:
	;
	v55 = F_tolower(m, v51)
	mBase = m.M
	v56 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v52))))
	v57 = F_tolower(m, v56)
	mBase = m.M
	goto L3
L5:
	;
	v25 = v19
	v26 = v20
	v27 = v23
	goto L8
L6:
	;
	v51 = int32(0)
	v52 = v20
	goto L4
L7:
	;
	v51 = v48 & int32(255)
	v52 = v47
	goto L4
L8:
	;
	v29 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v26))))
	if v29 == int32(0) {
		v47 = v26
		v48 = v27
		goto L7
	} else {
		goto L10
	}
L9:
	;
	v47 = v41
	v48 = int32(0)
	goto L7
L10:
	;
	v33 = v27 & int32(255)
	if v33 == v29 {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v40 = int32(1)
	v41 = v26 + v40
	v42 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v25)+1)))
	if v42 != 0 {
		v25 = v25 + v40
		v26 = v41
		v27 = v42
		goto L8
	} else {
		goto L14
	}
L12:
	;
	v35 = F_tolower(m, v33)
	mBase = m.M
	v36 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v26))))
	v37 = F_tolower(m, v36)
	mBase = m.M
	if v35 == v37 {
		goto L11
	} else {
		goto L13
	}
L13:
	;
	v39 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v25))))
	v47 = v26
	v48 = v39
	goto L7
L14:
	;
	goto L9
L15:
	;
	v62 = *(*int32)(unsafe.Add(mBase, _consts[59]))
	F_addReplyErrorObject(m, l0, v62)
	mBase = m.M
	v64 = m.ExcPending
	if v64 != 0 {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	return
L17:
	;
	goto L1
L18:
	;
	if v70 != 0 {
		goto L1
	} else {
		goto L19
	}
L19:
	;
	v72 = *(*int64)(unsafe.Add(mBase, uint32(v11)+8))
	if v72 == int64(0) {
		goto L21
	} else {
		goto L22
	}
L20:
	;
	v84 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v85 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v85)+4))
	v87 = F_lookupKeyRead(m, v84, v86)
	mBase = m.M
	v88 = m.ExcPending
	if v88 != 0 {
		goto L16
	} else {
		goto L25
	}
L21:
	;
	v81 = *(*int32)(unsafe.Add(mBase, _consts[59]))
	F_addReplyErrorObject(m, l0, v81)
	mBase = m.M
	v83 = m.ExcPending
	if v83 != 0 {
		goto L16
	} else {
		goto L24
	}
L22:
	;
	v75 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v72 == base.I64_extend_i32_s(v75+int32(-4)) {
		goto L20
	} else {
		goto L23
	}
L23:
	;
	goto L21
L24:
	;
	goto L1
L25:
	;
	v90 = F_checkType(m, l0, v87, int32(4))
	mBase = m.M
	v91 = m.ExcPending
	if v91 != 0 {
		goto L16
	} else {
		goto L26
	}
L26:
	;
	if v90 != 0 {
		goto L1
	} else {
		goto L27
	}
L27:
	;
	v92 = *(*int32)(unsafe.Add(mBase, uint32(v11)+8))
	F_addReplyArrayLen(m, l0, v92)
	mBase = m.M
	v94 = m.ExcPending
	if v94 != 0 {
		goto L16
	} else {
		goto L28
	}
L28:
	;
	v95 = *(*int64)(unsafe.Add(mBase, uint32(v11)+8))
	if v95 < int64(1) {
		goto L1
	} else {
		goto L29
	}
L29:
	;
	v107 = int64(0)
	goto L30
L30:
	;
	v109 = int64(-2)
	if v87 == int32(0) {
		v137 = v109
		goto L32
	} else {
		goto L33
	}
L31:
	;
	goto L1
L32:
	;
	F_addReplyLongLong(m, l0, v137)
	mBase = m.M
	v140 = m.ExcPending
	if v140 != 0 {
		goto L16
	} else {
		goto L41
	}
L33:
	;
	v112 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v117 = *(*int32)(unsafe.Add(mBase, uint32(v112+base.I32_wrap_i64(v107)<<(uint(int32(2))%32))+16))
	v118 = F_objectGetVal(m, v117)
	mBase = m.M
	v119 = F_hashTypeGetExpiry(m, v87, v118, v11)
	mBase = m.M
	v120 = m.ExcPending
	if v120 != 0 {
		goto L16
	} else {
		goto L34
	}
L34:
	;
	if v119 == int32(-1) {
		v137 = v109
		goto L32
	} else {
		goto L35
	}
L35:
	;
	v123 = int64(-1)
	v124 = *(*int64)(unsafe.Add(mBase, uint32(v11)))
	if v124 == v123 {
		v137 = v123
		goto L32
	} else {
		goto L36
	}
L36:
	;
	v127 = v124 - l1
	v128 = int64(0)
	if v128 < v127 {
		goto L37
	} else {
		goto L38
	}
L37:
	;
	v131 = v127
	goto L39
L38:
	;
	v131 = v128
	goto L39
L39:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v11))) = v131
	if l2 == int32(1) {
		v137 = v131
		goto L32
	} else {
		goto L40
	}
L40:
	;
	v136 = base.I64_div_u_s(v131+int64(500), int64(1000))
	v137 = v136
	goto L32
L41:
	;
	v141 = *(*int64)(unsafe.Add(mBase, uint32(v11)+8))
	v143 = v107 + int64(1)
	if v143 < v141 {
		v107 = v143
		goto L30
	} else {
		goto L42
	}
L42:
	;
	goto L31
}
