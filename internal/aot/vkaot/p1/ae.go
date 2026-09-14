package p1

import (
	base "github.com/shibukawa/vkmem/internal/aot/vkaot/base"
	"unsafe"
)

func F_aeCreateEventLoop(m *base.Module, l0 int32) int32 {
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
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v38 int64
	_ = v38
	var v42 int32
	_ = v42
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v71 int32
	_ = v71
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v82 int32
	_ = v82
	var v86 int32
	_ = v86
	var v89 int32
	_ = v89
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v128 int32
	_ = v128
	var v132 int32
	_ = v132
	var v141 int32
	_ = v141
	var v143 int32
	_ = v143
	var v154 int32
	_ = v154
	var v157 int32
	_ = v157
	var v168 int32
	_ = v168
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v173 int32
	_ = v173
	var v175 int32
	_ = v175
	var v178 int32
	_ = v178
	v9 = m.G0
	v11 = v9 - int32(16)
	m.G0 = v11
	v13 = F_monotonicInit(m)
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		return int32(0)
	} else {
		v17 = int32(0)
		v19 = F_valkey_malloc(m, int32(80))
		mBase = m.M
		v20 = m.ExcPending
		if v20 != 0 {
			return int32(0)
		} else {
			if v19 == int32(0) {
				v178 = v17
				m.G0 = v11 + int32(16)
				return v178
			} else {
				v25 = F_valkey_malloc(m, l0<<(uint(int32(4))%32))
				mBase = m.M
				v26 = m.ExcPending
				if v26 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v19)+16)) = v25
					v30 = F_valkey_malloc(m, l0<<(uint(int32(3))%32))
					mBase = m.M
					v31 = m.ExcPending
					if v31 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v19)+20)) = v30
						if v25 == int32(0) {
							v168 = *(*int32)(unsafe.Add(mBase, uint32(v19)+16))
							F_valkey_free(m, v168)
							mBase = m.M
							v170 = m.ExcPending
							if v170 != 0 {
								return int32(0)
							} else {
								v171 = *(*int32)(unsafe.Add(mBase, uint32(v19)+20))
								F_valkey_free(m, v171)
								mBase = m.M
								v173 = m.ExcPending
								if v173 != 0 {
									return int32(0)
								} else {
									F_valkey_free(m, v19)
									mBase = m.M
									v175 = m.ExcPending
									if v175 != 0 {
										return int32(0)
									} else {
										v178 = v17
										m.G0 = v11 + int32(16)
										return v178
									}
								}
							}
						} else {
							if v30 == int32(0) {
								v168 = *(*int32)(unsafe.Add(mBase, uint32(v19)+16))
								F_valkey_free(m, v168)
								mBase = m.M
								v170 = m.ExcPending
								if v170 != 0 {
									return int32(0)
								} else {
									v171 = *(*int32)(unsafe.Add(mBase, uint32(v19)+20))
									F_valkey_free(m, v171)
									mBase = m.M
									v173 = m.ExcPending
									if v173 != 0 {
										return int32(0)
									} else {
										F_valkey_free(m, v19)
										mBase = m.M
										v175 = m.ExcPending
										if v175 != 0 {
											return int32(0)
										} else {
											v178 = v17
											m.G0 = v11 + int32(16)
											return v178
										}
									}
								}
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v19)+4)) = l0
								v38 = int64(0)
								*(*int64)(unsafe.Add(mBase, uint32(v19)+24)) = v38
								*(*int64)(unsafe.Add(mBase, uint32(v19)+8)) = int64(1)
								v42 = int32(0)
								*(*int32)(unsafe.Add(mBase, uint32(v19)+72)) = v42
								*(*int32)(unsafe.Add(mBase, uint32(v19)+44)) = v42
								*(*int64)(unsafe.Add(mBase, uint32(v19)+36)) = v38
								*(*int32)(unsafe.Add(mBase, uint32(v19))) = int32(-1)
								v63 = F_valkey_malloc(m, int32(512))
								mBase = m.M
								v64 = m.ExcPending
								if v64 != 0 {
									return int32(0)
								} else {
									if v63 == int32(0) {
										v168 = *(*int32)(unsafe.Add(mBase, uint32(v19)+16))
										F_valkey_free(m, v168)
										mBase = m.M
										v170 = m.ExcPending
										if v170 != 0 {
											return int32(0)
										} else {
											v171 = *(*int32)(unsafe.Add(mBase, uint32(v19)+20))
											F_valkey_free(m, v171)
											mBase = m.M
											v173 = m.ExcPending
											if v173 != 0 {
												return int32(0)
											} else {
												F_valkey_free(m, v19)
												mBase = m.M
												v175 = m.ExcPending
												if v175 != 0 {
													return int32(0)
												} else {
													v178 = v17
													m.G0 = v11 + int32(16)
													return v178
												}
											}
										}
									} else {
										v67 = int32(0)
										v71 = F__emscripten_memset_bulkmem(m, v63, base.I32_extend8_s(v67), int32(256))
										mBase = m.M
										*(*int32)(unsafe.Add(mBase, uint32(v19)+32)) = v71
										if l0 < int32(1) {
										} else {
											v76 = l0 & int32(7)
											v77 = *(*int32)(unsafe.Add(mBase, uint32(v19)+16))
											if base.Ui32(l0) < base.Ui32(int32(8)) {
												v132 = v67
											} else {
												v82 = int32(0)
												v86 = v82
												v89 = v82
												for {
													v94 = v77 + v86<<(uint(int32(4))%32)
													v95 = int32(0)
													*(*int32)(unsafe.Add(mBase, uint32(v94))) = v95
													*(*int32)(unsafe.Add(mBase, uint32(v94+int32(16)))) = v95
													*(*int32)(unsafe.Add(mBase, uint32(v94+int32(32)))) = v95
													*(*int32)(unsafe.Add(mBase, uint32(v94+int32(48)))) = v95
													*(*int32)(unsafe.Add(mBase, uint32(v94+int32(64)))) = v95
													*(*int32)(unsafe.Add(mBase, uint32(v94+int32(80)))) = v95
													*(*int32)(unsafe.Add(mBase, uint32(v94+int32(96)))) = v95
													*(*int32)(unsafe.Add(mBase, uint32(v94+int32(112)))) = v95
													v125 = int32(8)
													v126 = v86 + v125
													v128 = v89 + v125
													if v128 != l0&int32(2147483640) {
														v86 = v126
														v89 = v128
														continue
													} else {
														break
													}
													break
												}
												v132 = v126
											}
											if v76 == int32(0) {
											} else {
												v141 = int32(0)
												v143 = v132
												for {
													*(*int32)(unsafe.Add(mBase, uint32(v77+v143<<(uint(int32(4))%32)))) = int32(0)
													v154 = int32(1)
													v157 = v141 + v154
													if v157 != v76 {
														v141 = v157
														v143 = v143 + v154
														continue
													} else {
														break
													}
													break
												}
											}
										}
										v178 = v19
										m.G0 = v11 + int32(16)
										return v178
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
func F_aeCreateFileEvent(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v19 int32
	_ = v19
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v60 int32
	_ = v60
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v80 int32
	_ = v80
	v11 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+72)))
	if v11&int32(32) == int32(0) {
	} else {
	}
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if l1 < v19 {
		v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
		v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
		v30 = l2 & int32(1)
		if v30 == int32(0) {
		} else {
			v37 = v28 + int32(base.Ui32(l1)>>(uint(int32(3))%32))&int32(536870908)
			v38 = *(*int32)(unsafe.Add(mBase, uint32(v37)))
			*(*int32)(unsafe.Add(mBase, uint32(v37))) = v38 | int32(1)<<(uint(l1)%32)
		}
		v44 = v27 + l1<<(uint(int32(4))%32)
		v46 = l2 & int32(2)
		if v46 == int32(0) {
		} else {
			v53 = v28 + int32(base.Ui32(l1)>>(uint(int32(3))%32))&int32(536870908)
			v54 = *(*int32)(unsafe.Add(mBase, uint32(v53)+128))
			*(*int32)(unsafe.Add(mBase, uint32(v53)+128)) = v54 | int32(1)<<(uint(l1)%32)
		}
		v60 = *(*int32)(unsafe.Add(mBase, uint32(v44)))
		*(*int32)(unsafe.Add(mBase, uint32(v44))) = v60 | l2
		if v30 == int32(0) {
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v44)+4)) = l3
		}
		if v46 == int32(0) {
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v44)+8)) = l3
		}
		*(*int32)(unsafe.Add(mBase, uint32(v44)+12)) = l4
		v70 = int32(0)
		v71 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		if l1 <= v71 {
			v74 = v70
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(l0))) = l1
			v74 = v70
		}
	} else {
		*(*int32)(unsafe.Add(mBase, _consts[18])) = int32(68)
		v74 = int32(-1)
	}
	v80 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+72)))
	if v80&int32(32) == int32(0) {
	} else {
	}
	return v74
}
func F_aeGetFileClientData(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
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
	var v16 int32
	_ = v16
	v3 = int32(0)
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v5 <= l1 {
		v16 = v3
	} else {
		v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
		v10 = v7 + l1<<(uint(int32(4))%32)
		v11 = *(*int32)(unsafe.Add(mBase, uint32(v10)))
		if v11 == int32(0) {
			v16 = v3
		} else {
			v14 = *(*int32)(unsafe.Add(mBase, uint32(v10)+12))
			v16 = v14
		}
	}
	return v16
}
func F_aeMain(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = int32(0)
	goto L1
L1:
	;
	v6 = F_aeProcessEvents(m, l0, int32(27))
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		goto L3
	} else {
		goto L4
	}
L2:
	;
	return
L3:
	;
	return
L4:
	;
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v8 == int32(0) {
		goto L1
	} else {
		goto L5
	}
L5:
	;
	goto L2
}
func F_aeProcessEvents(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v31 int32
	_ = v31
	var v39 int32
	_ = v39
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v65 int32
	_ = v65
	var v68 int64
	_ = v68
	var v71 int64
	_ = v71
	var v72 int64
	_ = v72
	var v76 int64
	_ = v76
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v97 int32
	_ = v97
	var v101 int32
	_ = v101
	var v111 int32
	_ = v111
	var v122 int32
	_ = v122
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v134 int32
	_ = v134
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v158 int32
	_ = v158
	var v160 int32
	_ = v160
	var v162 int32
	_ = v162
	var v164 int32
	_ = v164
	var v168 int32
	_ = v168
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v188 int32
	_ = v188
	var v194 int32
	_ = v194
	var v199 int32
	_ = v199
	var v214 int64
	_ = v214
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v218 int32
	_ = v218
	var v219 int64
	_ = v219
	var v223 int32
	_ = v223
	var v230 int64
	_ = v230
	var v235 int32
	_ = v235
	var v238 int64
	_ = v238
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v245 int32
	_ = v245
	var v249 int32
	_ = v249
	var v251 int32
	_ = v251
	var v255 int32
	_ = v255
	var v258 int32
	_ = v258
	var v260 int32
	_ = v260
	var v262 int32
	_ = v262
	var v263 int64
	_ = v263
	var v264 int64
	_ = v264
	var v266 int32
	_ = v266
	var v268 int64
	_ = v268
	var v270 int32
	_ = v270
	var v274 int32
	_ = v274
	var v275 int32
	_ = v275
	var v276 int64
	_ = v276
	var v277 int32
	_ = v277
	var v278 int32
	_ = v278
	var v283 int32
	_ = v283
	var v285 int32
	_ = v285
	var v286 int64
	_ = v286
	var v296 int64
	_ = v296
	var v297 int32
	_ = v297
	var v298 int32
	_ = v298
	var v299 int32
	_ = v299
	var v302 int64
	_ = v302
	var v305 int32
	_ = v305
	var v319 int32
	_ = v319
	var v326 int32
	_ = v326
	v17 = m.G0
	v19 = v17 - int32(16)
	m.G0 = v19
	if l1&int32(3) != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v19 + int32(16)
	return v326
L2:
	;
	v24 = int32(2)
	v25 = l1 & v24
	if l1&int32(6) == v24 {
		goto L5
	} else {
		goto L6
	}
L3:
	;
	v326 = int32(0)
	goto L1
L4:
	;
	if v25 == int32(0) {
		v326 = v199
		goto L1
	} else {
		goto L57
	}
L5:
	;
	if l1&int32(8) == int32(0) {
		goto L8
	} else {
		goto L9
	}
L6:
	;
	v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v31 == int32(-1) {
		v199 = int32(0)
		goto L4
	} else {
		goto L7
	}
L7:
	;
	goto L5
L8:
	;
	v49 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	if v49 == int32(0) {
		goto L14
	} else {
		goto L15
	}
L9:
	;
	v39 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	if v39 == int32(0) {
		goto L8
	} else {
		goto L10
	}
L10:
	;
	m.T0[v39].(func(*base.Module, int32))(m, l0)
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	return int32(0)
L12:
	;
	goto L8
L13:
	;
	v90 = int32(0)
	if l1&int32(1) != 0 {
		goto L26
	} else {
		goto L27
	}
L14:
	;
	if l1&int32(4) != 0 {
		goto L20
	} else {
		goto L21
	}
L15:
	;
	v52 = m.T0[v49].(func(*base.Module, int32) int32)(m, l0)
	mBase = m.M
	v53 = m.ExcPending
	if v53 != 0 {
		goto L11
	} else {
		goto L16
	}
L16:
	;
	v87 = v52
	goto L13
L17:
	;
	v84 = F_aeApiPoll(m, l0, v81)
	mBase = m.M
	v85 = m.ExcPending
	if v85 != 0 {
		goto L11
	} else {
		goto L25
	}
L18:
	;
	v81 = v19
	goto L17
L19:
	;
	v65 = int32(0)
	if v25 == v65 {
		v81 = v65
		goto L17
	} else {
		goto L23
	}
L20:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v19))) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v19)+8)) = int32(0)
	goto L18
L21:
	;
	v56 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+72)))
	if v56&int32(4) == int32(0) {
		goto L19
	} else {
		goto L22
	}
L22:
	;
	goto L20
L23:
	;
	v68 = F_usUntilEarliestTimer(m, l0)
	mBase = m.M
	if v68 < int64(0) {
		v81 = v65
		goto L17
	} else {
		goto L24
	}
L24:
	;
	v71 = int64(1000000)
	v72 = base.I64_div_u_s(v68, v71)
	*(*int64)(unsafe.Add(mBase, uint32(v19))) = v72
	v76 = v68 - v72*v71
	*(*uint32)(unsafe.Add(mBase, uint32(v19)+8)) = uint32(v76)
	goto L18
L25:
	;
	v87 = v84
	goto L13
L26:
	;
	v92 = v87
	goto L28
L27:
	;
	v92 = v90
	goto L28
L28:
	;
	if l1&int32(16) == int32(0) {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	if v92 < int32(1) {
		v199 = v90
		goto L4
	} else {
		goto L33
	}
L30:
	;
	v97 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	if v97 == int32(0) {
		goto L29
	} else {
		goto L31
	}
L31:
	;
	m.T0[v97].(func(*base.Module, int32, int32))(m, l0, v92)
	mBase = m.M
	v101 = m.ExcPending
	if v101 != 0 {
		goto L11
	} else {
		goto L32
	}
L32:
	;
	goto L29
L33:
	;
	v111 = int32(0)
	goto L34
L34:
	;
	v122 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v125 = v122 + v111<<(uint(int32(3))%32)
	v126 = *(*int32)(unsafe.Add(mBase, uint32(v125)+4))
	v127 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v128 = *(*int32)(unsafe.Add(mBase, uint32(v125)))
	v129 = int32(4)
	v130 = v128 << (uint(v129) % 32)
	v131 = v127 + v130
	v132 = *(*int32)(unsafe.Add(mBase, uint32(v131)))
	v134 = v132 & v129
	if v134 != 0 {
		goto L40
	} else {
		goto L41
	}
L35:
	;
	v199 = v87
	goto L4
L36:
	;
	v194 = v111 + int32(1)
	if v194 != v92 {
		v111 = v194
		goto L34
	} else {
		goto L56
	}
L37:
	;
	if v134 == int32(0) {
		goto L36
	} else {
		goto L49
	}
L38:
	;
	v162 = *(*int32)(unsafe.Add(mBase, uint32(v158+v130)+12))
	m.T0[v160].(func(*base.Module, int32, int32, int32, int32))(m, l0, v128, v162, v126)
	mBase = m.M
	v164 = m.ExcPending
	if v164 != 0 {
		goto L11
	} else {
		goto L48
	}
L39:
	;
	v143 = *(*int32)(unsafe.Add(mBase, uint32(v131)+12))
	v144 = *(*int32)(unsafe.Add(mBase, uint32(v131)+4))
	m.T0[v144].(func(*base.Module, int32, int32, int32, int32))(m, l0, v128, v143, v126)
	mBase = m.M
	v146 = m.ExcPending
	if v146 != 0 {
		goto L11
	} else {
		goto L45
	}
L40:
	;
	if v126&v132&int32(2) != 0 {
		goto L43
	} else {
		goto L44
	}
L41:
	;
	if v126&v132&int32(1) != 0 {
		goto L39
	} else {
		goto L42
	}
L42:
	;
	goto L40
L43:
	;
	v142 = *(*int32)(unsafe.Add(mBase, uint32(v131)+8))
	v158 = v127
	v160 = v142
	goto L38
L44:
	;
	v168 = int32(1)
	goto L37
L45:
	;
	v147 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v148 = v147 + v130
	v149 = *(*int32)(unsafe.Add(mBase, uint32(v148)))
	if v126&v149&int32(2) == int32(0) {
		goto L36
	} else {
		goto L46
	}
L46:
	;
	v155 = *(*int32)(unsafe.Add(mBase, uint32(v148)+8))
	v156 = *(*int32)(unsafe.Add(mBase, uint32(v148)+4))
	if v155 == v156 {
		goto L36
	} else {
		goto L47
	}
L47:
	;
	v158 = v147
	v160 = v155
	goto L38
L48:
	;
	v168 = int32(0)
	goto L37
L49:
	;
	v171 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v172 = v171 + v130
	v173 = *(*int32)(unsafe.Add(mBase, uint32(v172)))
	if v126&v173&int32(1) == int32(0) {
		goto L36
	} else {
		goto L50
	}
L50:
	;
	if v168 == int32(0) {
		goto L52
	} else {
		goto L53
	}
L51:
	;
	v186 = *(*int32)(unsafe.Add(mBase, uint32(v172)+12))
	m.T0[v185].(func(*base.Module, int32, int32, int32, int32))(m, l0, v128, v186, v126)
	mBase = m.M
	v188 = m.ExcPending
	if v188 != 0 {
		goto L11
	} else {
		goto L55
	}
L52:
	;
	v182 = *(*int32)(unsafe.Add(mBase, uint32(v172)+8))
	v183 = *(*int32)(unsafe.Add(mBase, uint32(v172)+4))
	if v182 == v183 {
		goto L36
	} else {
		goto L54
	}
L53:
	;
	v181 = *(*int32)(unsafe.Add(mBase, uint32(v172)+4))
	v185 = v181
	goto L51
L54:
	;
	v185 = v183
	goto L51
L55:
	;
	goto L36
L56:
	;
	goto L35
L57:
	;
	v214 = *(*int64)(unsafe.Add(mBase, uint32(l0)+8))
	v215 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v216 = int32(0)
	v218 = *(*int32)(unsafe.Add(mBase, _consts[34]))
	v219 = m.T0[v218].(func(*base.Module) int64)(m)
	mBase = m.M
	if v215 == v216 {
		v319 = v216
		goto L58
	} else {
		goto L59
	}
L58:
	;
	v326 = v319 + v199
	goto L1
L59:
	;
	v223 = v215
	v230 = v219
	v235 = v216
	goto L60
L60:
	;
	v238 = *(*int64)(unsafe.Add(mBase, uint32(v223)))
	if v238 != int64(-1) {
		goto L63
	} else {
		goto L64
	}
L61:
	;
	v319 = v305
	goto L58
L62:
	;
	if v299 != 0 {
		v223 = v299
		v230 = v302
		v235 = v305
		goto L60
	} else {
		goto L82
	}
L63:
	;
	if v214 <= v238 {
		v296 = v230
		v297 = v235
		goto L76
	} else {
		goto L77
	}
L64:
	;
	v241 = *(*int32)(unsafe.Add(mBase, uint32(v223)+32))
	v242 = *(*int32)(unsafe.Add(mBase, uint32(v223)+36))
	if v242 == int32(0) {
		goto L65
	} else {
		goto L66
	}
L65:
	;
	v245 = *(*int32)(unsafe.Add(mBase, uint32(v223)+28))
	if v245 == int32(0) {
		goto L68
	} else {
		goto L69
	}
L66:
	;
	v299 = v241
	v302 = v230
	v305 = v235
	goto L62
L67:
	;
	if v251 == int32(0) {
		goto L70
	} else {
		goto L71
	}
L68:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = v241
	v251 = v241
	goto L67
L69:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v245)+32)) = v241
	v249 = *(*int32)(unsafe.Add(mBase, uint32(v223)+32))
	v251 = v249
	goto L67
L70:
	;
	v255 = *(*int32)(unsafe.Add(mBase, uint32(v223)+20))
	if v255 == int32(0) {
		v264 = v230
		goto L72
	} else {
		goto L73
	}
L71:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v251)+28)) = v245
	goto L70
L72:
	;
	F_valkey_free(m, v223)
	mBase = m.M
	v266 = m.ExcPending
	if v266 != 0 {
		goto L11
	} else {
		goto L75
	}
L73:
	;
	v258 = *(*int32)(unsafe.Add(mBase, uint32(v223)+24))
	m.T0[v255].(func(*base.Module, int32, int32))(m, l0, v258)
	mBase = m.M
	v260 = m.ExcPending
	if v260 != 0 {
		goto L11
	} else {
		goto L74
	}
L74:
	;
	v262 = *(*int32)(unsafe.Add(mBase, _consts[34]))
	v263 = m.T0[v262].(func(*base.Module) int64)(m)
	mBase = m.M
	v264 = v263
	goto L72
L75:
	;
	v299 = v241
	v302 = v264
	v305 = v235
	goto L62
L76:
	;
	v298 = *(*int32)(unsafe.Add(mBase, uint32(v223)+32))
	v299 = v298
	v302 = v296
	v305 = v297
	goto L62
L77:
	;
	v268 = *(*int64)(unsafe.Add(mBase, uint32(v223)+8))
	if base.Ui64(v230) < base.Ui64(v268) {
		v296 = v230
		v297 = v235
		goto L76
	} else {
		goto L78
	}
L78:
	;
	v270 = *(*int32)(unsafe.Add(mBase, uint32(v223)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v223)+36)) = v270 + int32(1)
	v274 = *(*int32)(unsafe.Add(mBase, uint32(v223)+24))
	v275 = *(*int32)(unsafe.Add(mBase, uint32(v223)+16))
	v276 = m.T0[v275].(func(*base.Module, int32, int64, int32) int64)(m, l0, v238, v274)
	mBase = m.M
	v277 = m.ExcPending
	if v277 != 0 {
		goto L11
	} else {
		goto L79
	}
L79:
	;
	v278 = *(*int32)(unsafe.Add(mBase, uint32(v223)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v223)+36)) = v278 + int32(-1)
	v283 = v235 + int32(1)
	v285 = *(*int32)(unsafe.Add(mBase, _consts[34]))
	v286 = m.T0[v285].(func(*base.Module) int64)(m)
	mBase = m.M
	if v276 == int64(-1) {
		goto L80
	} else {
		goto L81
	}
L80:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v223))) = int64(-1)
	v296 = v286
	v297 = v283
	goto L76
L81:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v223)+8)) = v286 + v276*int64(1000)
	v296 = v286
	v297 = v283
	goto L76
L82:
	;
	goto L61
}
func F_aeWait(m *base.Module, l0 int32, l1 int32, l2 int64) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v40 int32
	_ = v40
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	*(*int64)(unsafe.Add(mBase, uint32(v7)+8)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v7)+8)) = l0
	if l1&int32(1) == int32(0) {
		v20 = int32(4)
	} else {
		v17 = int32(1)
		*(*uint16)(unsafe.Add(mBase, uint32(v7)+12)) = uint16(v17)
		v20 = int32(5)
	}
	if l1&int32(2) == int32(0) {
	} else {
		*(*uint16)(unsafe.Add(mBase, uint32(v7)+12)) = uint16(v20)
	}
	v28 = int32(1)
	v30 = F_poll(m, v7+int32(8), v28, base.I32_wrap_i64(l2))
	mBase = m.M
	if v30 != v28 {
		v47 = v30
	} else {
		v33 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v7)+14)))
		v34 = int32(1)
		v36 = int32(2)
		v40 = int32(base.Ui32(v33)>>(uint(v34)%32))&v36 | v33&v34
		if v33&int32(24) != 0 {
			v45 = v40 | v36
		} else {
			v45 = v40
		}
		v47 = v45
	}
	m.G0 = v7 + int32(16)
	return v47
}
