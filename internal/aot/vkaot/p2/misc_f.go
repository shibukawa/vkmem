package p2

import (
	base "github.com/shibukawa/vkmem/internal/aot/vkaot/base"
	"math"
	"unsafe"
)

func F___fe_getround(m *base.Module) int32 {
	return int32(0)
}
func F___floatditf(m *base.Module, l0 int32, l1 int64) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v14 int64
	_ = v14
	var v17 int64
	_ = v17
	var v19 int64
	_ = v19
	var v20 int64
	_ = v20
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v41 int64
	_ = v41
	var v45 int64
	_ = v45
	var v46 int64
	_ = v46
	var v52 int64
	_ = v52
	var v64 int64
	_ = v64
	var v65 int64
	_ = v65
	var v66 int64
	_ = v66
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	if base.B2i32(l1 == int64(0)) == int32(0) {
		v17 = l1 >> (uint(int64(63)) % 64)
		v19 = l1 ^ v17 - v17
		v20 = int64(0)
		v22 = base.I32_wrap_i64(base.I64_clz(v19))
		v24 = v22 + int32(49)
		if v24&int32(64) == int32(0) {
			if v24 == int32(0) {
				v45 = v19
				v46 = v20
			} else {
				v41 = base.I64_extend_i32_u(v24)
				v45 = v19 << (uint(v41) % 64)
				v46 = int64(base.Ui64(v19)>>(uint(base.I64_extend_i32_u(int32(64)-v24))%64)) | v20<<(uint(v41)%64)
			}
		} else {
			v45 = int64(0)
			v46 = v19 << (uint(base.I64_extend_i32_u(v22+int32(-15))) % 64)
		}
		*(*int64)(unsafe.Add(mBase, uint32(v8))) = v45
		*(*int64)(unsafe.Add(mBase, uint32(v8)+8)) = v46
		v52 = *(*int64)(unsafe.Add(mBase, uint32(v8+int32(8))))
		v64 = *(*int64)(unsafe.Add(mBase, uint32(v8)))
		v65 = v64
		v66 = v52 ^ int64(281474976710656) + base.I64_extend_i32_u(int32(16446)-v22)<<(uint(int64(48))%64) | l1&int64(-9223372036854775807-1)
	} else {
		v14 = int64(0)
		v65 = v14
		v66 = v14
	}
	*(*int64)(unsafe.Add(mBase, uint32(l0))) = v65
	*(*int64)(unsafe.Add(mBase, uint32(l0)+8)) = v66
	m.G0 = v8 + int32(16)
	return
}
func F___fseeko(m *base.Module, l0 int32, l1 int64, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
	if int32(-1) < v5 {
		v14 = F___fseeko_unlocked(m, l0, l1, l2)
		mBase = m.M
		v15 = m.ExcPending
		if v15 != 0 {
			return int32(0)
		} else {
			return v14
		}
	} else {
		v8 = F___fseeko_unlocked(m, l0, l1, l2)
		mBase = m.M
		v11 = m.ExcPending
		if v11 != 0 {
			return int32(0)
		} else {
			return v8
		}
	}
}
func F___fstatat(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v37 int32
	_ = v37
	var v42 int32
	_ = v42
	if l0 < int32(0) {
		if l0 == int32(-100) {
			if l3 == int32(256) {
				v32 = m.Env.X__syscall_lstat64(m, l1, l2)
				mBase = m.M
				v33 = v32
			} else {
				if l3 != 0 {
					v30 = m.Env.X__syscall_newfstatat(m, l0, l1, l2, l3)
					mBase = m.M
					v33 = v30
				} else {
					v28 = m.Env.X__syscall_stat64(m, l1, l2)
					mBase = m.M
					v33 = v28
				}
			}
		} else {
			v14 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
			if l3 != 0 {
				if l3 != int32(256) {
					v30 = m.Env.X__syscall_newfstatat(m, l0, l1, l2, l3)
					mBase = m.M
					v33 = v30
				} else {
					if v14&int32(255) != int32(47) {
						v30 = m.Env.X__syscall_newfstatat(m, l0, l1, l2, l3)
						mBase = m.M
						v33 = v30
					} else {
						v32 = m.Env.X__syscall_lstat64(m, l1, l2)
						mBase = m.M
						v33 = v32
					}
				}
			} else {
				if v14&int32(255) == int32(47) {
					v28 = m.Env.X__syscall_stat64(m, l1, l2)
					mBase = m.M
					v33 = v28
				} else {
					if l3 != int32(256) {
						v30 = m.Env.X__syscall_newfstatat(m, l0, l1, l2, l3)
						mBase = m.M
						v33 = v30
					} else {
						if v14&int32(255) != int32(47) {
							v30 = m.Env.X__syscall_newfstatat(m, l0, l1, l2, l3)
							mBase = m.M
							v33 = v30
						} else {
							v32 = m.Env.X__syscall_lstat64(m, l1, l2)
							mBase = m.M
							v33 = v32
						}
					}
				}
			}
		}
	} else {
		if l3 != int32(4096) {
			if l0 == int32(-100) {
				if l3 == int32(256) {
					v32 = m.Env.X__syscall_lstat64(m, l1, l2)
					mBase = m.M
					v33 = v32
				} else {
					if l3 != 0 {
						v30 = m.Env.X__syscall_newfstatat(m, l0, l1, l2, l3)
						mBase = m.M
						v33 = v30
					} else {
						v28 = m.Env.X__syscall_stat64(m, l1, l2)
						mBase = m.M
						v33 = v28
					}
				}
			} else {
				v14 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
				if l3 != 0 {
					if l3 != int32(256) {
						v30 = m.Env.X__syscall_newfstatat(m, l0, l1, l2, l3)
						mBase = m.M
						v33 = v30
					} else {
						if v14&int32(255) != int32(47) {
							v30 = m.Env.X__syscall_newfstatat(m, l0, l1, l2, l3)
							mBase = m.M
							v33 = v30
						} else {
							v32 = m.Env.X__syscall_lstat64(m, l1, l2)
							mBase = m.M
							v33 = v32
						}
					}
				} else {
					if v14&int32(255) == int32(47) {
						v28 = m.Env.X__syscall_stat64(m, l1, l2)
						mBase = m.M
						v33 = v28
					} else {
						if l3 != int32(256) {
							v30 = m.Env.X__syscall_newfstatat(m, l0, l1, l2, l3)
							mBase = m.M
							v33 = v30
						} else {
							if v14&int32(255) != int32(47) {
								v30 = m.Env.X__syscall_newfstatat(m, l0, l1, l2, l3)
								mBase = m.M
								v33 = v30
							} else {
								v32 = m.Env.X__syscall_lstat64(m, l1, l2)
								mBase = m.M
								v33 = v32
							}
						}
					}
				}
			}
		} else {
			v10 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
			if v10 != 0 {
				v30 = m.Env.X__syscall_newfstatat(m, l0, l1, l2, l3)
				mBase = m.M
				v33 = v30
			} else {
				v11 = m.Env.X__syscall_fstat64(m, l0, l2)
				mBase = m.M
				v33 = v11
			}
		}
	}
	if base.Ui32(v33) < base.Ui32(int32(-4095)) {
		v42 = v33
	} else {
		v37 = F___errno_location(m)
		mBase = m.M
		*(*int32)(unsafe.Add(mBase, uint32(v37))) = int32(0) - v33
		v42 = int32(-1)
	}
	return v42
}
func F___ftello(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v7 int64
	_ = v7
	var v10 int32
	_ = v10
	var v13 int64
	_ = v13
	var v14 int32
	_ = v14
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
	if int32(-1) < v4 {
		v13 = F___ftello_unlocked(m, l0)
		mBase = m.M
		v14 = m.ExcPending
		if v14 != 0 {
			return int64(0)
		} else {
			return v13
		}
	} else {
		v7 = F___ftello_unlocked(m, l0)
		mBase = m.M
		v10 = m.ExcPending
		if v10 != 0 {
			return int64(0)
		} else {
			return v7
		}
	}
}
func F___funcs_on_exit(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v16 int32
	_ = v16
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
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
	var v53 int32
	_ = v53
	goto L1
L1:
	;
	v5 = int32(0)
	v6 = *(*int32)(unsafe.Add(mBase, _consts[1106]))
	if v6 == v5 {
		goto L2
	} else {
		goto L3
	}
L2:
	;
	return
L3:
	;
	v10 = *(*int32)(unsafe.Add(mBase, _consts[1107]))
	v11 = v6
	v12 = v10
	goto L4
L4:
	;
	v16 = v12 + int32(-1)
	*(*int32)(unsafe.Add(mBase, _consts[1107])) = v16
	if v12 < int32(1) {
		v48 = v11
		goto L6
	} else {
		goto L7
	}
L5:
	;
	goto L2
L6:
	;
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v48)))
	v52 = int32(32)
	v53 = int32(0)
	*(*int32)(unsafe.Add(mBase, _consts[1107])) = v52
	*(*int32)(unsafe.Add(mBase, _consts[1106])) = v51
	if v51 != 0 {
		v11 = v51
		v12 = v52
		goto L4
	} else {
		goto L15
	}
L7:
	;
	v22 = v16
	goto L8
L8:
	;
	v24 = *(*int32)(unsafe.Add(mBase, _consts[1106]))
	v27 = v24 + v22<<(uint(int32(2))%32)
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v27+int32(132))))
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v27+int32(4))))
	goto L10
L9:
	;
	v47 = *(*int32)(unsafe.Add(mBase, _consts[1106]))
	v48 = v47
	goto L6
L10:
	;
	m.T0[v33].(func(*base.Module, int32))(m, v30)
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	return
L12:
	;
	goto L13
L13:
	;
	v38 = int32(0)
	v40 = *(*int32)(unsafe.Add(mBase, _consts[1107]))
	v42 = v40 + int32(-1)
	*(*int32)(unsafe.Add(mBase, _consts[1107])) = v42
	if v38 < v40 {
		v22 = v42
		goto L8
	} else {
		goto L14
	}
L14:
	;
	goto L9
L15:
	;
	goto L5
}
func F_fabsl(m *base.Module, l0 int32, l1 int64, l2 int64) {
	mBase := m.M
	_ = mBase
	*(*int64)(unsafe.Add(mBase, uint32(l0)+8)) = l2 & int64(9223372036854775807)
	*(*int64)(unsafe.Add(mBase, uint32(l0))) = l1
	return
}
func F_failoverCommand(m *base.Module, l0 int32) {
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
	var v17 int32
	_ = v17
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
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v74 int32
	_ = v74
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v104 int32
	_ = v104
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
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
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v139 int32
	_ = v139
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v150 int32
	_ = v150
	var v152 int32
	_ = v152
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
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v165 int32
	_ = v165
	var v169 int32
	_ = v169
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
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
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
	var v199 int32
	_ = v199
	var v203 int32
	_ = v203
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v216 int32
	_ = v216
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v223 int32
	_ = v223
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v229 int32
	_ = v229
	var v233 int32
	_ = v233
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v247 int32
	_ = v247
	var v248 int32
	_ = v248
	var v251 int32
	_ = v251
	var v252 int32
	_ = v252
	var v255 int32
	_ = v255
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v264 int32
	_ = v264
	var v266 int32
	_ = v266
	var v267 int32
	_ = v267
	var v268 int32
	_ = v268
	var v269 int32
	_ = v269
	var v273 int32
	_ = v273
	var v274 int32
	_ = v274
	var v279 int32
	_ = v279
	var v280 int32
	_ = v280
	var v285 int32
	_ = v285
	var v290 int32
	_ = v290
	var v292 int32
	_ = v292
	var v297 int32
	_ = v297
	var v299 int32
	_ = v299
	var v300 int32
	_ = v300
	var v303 int32
	_ = v303
	var v306 int32
	_ = v306
	var v311 int32
	_ = v311
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
	var v321 int32
	_ = v321
	var v326 int32
	_ = v326
	var v328 int32
	_ = v328
	var v329 int32
	_ = v329
	var v331 int32
	_ = v331
	var v332 int32
	_ = v332
	var v335 int32
	_ = v335
	var v343 int32
	_ = v343
	var v345 int32
	_ = v345
	var v352 int32
	_ = v352
	var v355 int64
	_ = v355
	var v356 int32
	_ = v356
	var v363 int32
	_ = v363
	var v372 int32
	_ = v372
	var v374 int32
	_ = v374
	var v376 int32
	_ = v376
	var v378 int32
	_ = v378
	var v381 int32
	_ = v381
	var v384 int32
	_ = v384
	var v386 int32
	_ = v386
	var v388 int32
	_ = v388
	v9 = m.G0
	v11 = v9 - int32(16)
	m.G0 = v11
	v13 = F_clusterAllowFailoverCmd(m, l0)
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v11 + int32(16)
	return
L2:
	;
	return
L3:
	;
	if v13 == int32(0) {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v17 != int32(2) {
		v65 = v17
		goto L6
	} else {
		goto L7
	}
L5:
	;
	v378 = *(*int32)(unsafe.Add(mBase, _consts[658]))
	if v378 != 0 {
		goto L115
	} else {
		goto L116
	}
L6:
	;
	v66 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v11)+12)) = v66
	*(*int32)(unsafe.Add(mBase, uint32(v11)+8)) = v66
	if v65 < int32(2) {
		v279 = v66
		v280 = v66
		goto L21
	} else {
		goto L22
	}
L7:
	;
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v21 = *(*int32)(unsafe.Add(mBase, uint32(v20)+4))
	v22 = F_objectGetVal(m, v21)
	mBase = m.M
	v23 = int32(_a1295)
	v26 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22))))
	if v26 != 0 {
		goto L10
	} else {
		goto L11
	}
L8:
	;
	if v58-v60 == int32(0) {
		goto L5
	} else {
		goto L20
	}
L9:
	;
	v58 = F_tolower(m, v54)
	mBase = m.M
	v59 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v55))))
	v60 = F_tolower(m, v59)
	mBase = m.M
	goto L8
L10:
	;
	v28 = v22
	v29 = v23
	v30 = v26
	goto L13
L11:
	;
	v54 = int32(0)
	v55 = v23
	goto L9
L12:
	;
	v54 = v51 & int32(255)
	v55 = v50
	goto L9
L13:
	;
	v32 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v29))))
	if v32 == int32(0) {
		v50 = v29
		v51 = v30
		goto L12
	} else {
		goto L15
	}
L14:
	;
	v50 = v44
	v51 = int32(0)
	goto L12
L15:
	;
	v36 = v30 & int32(255)
	if v36 == v32 {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	v43 = int32(1)
	v44 = v29 + v43
	v45 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28)+1)))
	if v45 != 0 {
		v28 = v28 + v43
		v29 = v44
		v30 = v45
		goto L13
	} else {
		goto L19
	}
L17:
	;
	v38 = F_tolower(m, v36)
	mBase = m.M
	v39 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v29))))
	v40 = F_tolower(m, v39)
	mBase = m.M
	if v38 == v40 {
		goto L16
	} else {
		goto L18
	}
L18:
	;
	v42 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28))))
	v50 = v29
	v51 = v42
	goto L12
L19:
	;
	goto L14
L20:
	;
	v64 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v65 = v64
	goto L6
L21:
	;
	v285 = *(*int32)(unsafe.Add(mBase, _consts[658]))
	if v285 == int32(0) {
		goto L79
	} else {
		goto L80
	}
L22:
	;
	v74 = int32(0)
	v79 = int32(1)
	v80 = v74
	v81 = v74
	goto L23
L23:
	;
	v85 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v87 = v79 << (uint(int32(2)) % 32)
	v89 = *(*int32)(unsafe.Add(mBase, uint32(v85+v87)))
	v90 = F_objectGetVal(m, v89)
	mBase = m.M
	v91 = int32(_a1028)
	v94 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v90))))
	if v94 != 0 {
		goto L29
	} else {
		goto L30
	}
L24:
	;
	v279 = v268
	v280 = v269
	goto L21
L25:
	;
	v273 = v267 + int32(1)
	v274 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v273 < v274 {
		v79 = v273
		v80 = v268
		v81 = v269
		goto L23
	} else {
		goto L78
	}
L26:
	;
	v152 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v154 = *(*int32)(unsafe.Add(mBase, uint32(v152+v87)))
	v155 = F_objectGetVal(m, v154)
	mBase = m.M
	v156 = int32(_a332)
	v159 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v155))))
	if v159 != 0 {
		goto L49
	} else {
		goto L50
	}
L27:
	;
	if v126-v128 != 0 {
		goto L26
	} else {
		goto L39
	}
L28:
	;
	v126 = F_tolower(m, v122)
	mBase = m.M
	v127 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v123))))
	v128 = F_tolower(m, v127)
	mBase = m.M
	goto L27
L29:
	;
	v96 = v90
	v97 = v91
	v98 = v94
	goto L32
L30:
	;
	v122 = int32(0)
	v123 = v91
	goto L28
L31:
	;
	v122 = v119 & int32(255)
	v123 = v118
	goto L28
L32:
	;
	v100 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v97))))
	if v100 == int32(0) {
		v118 = v97
		v119 = v98
		goto L31
	} else {
		goto L34
	}
L33:
	;
	v118 = v112
	v119 = int32(0)
	goto L31
L34:
	;
	v104 = v98 & int32(255)
	if v104 == v100 {
		goto L35
	} else {
		goto L36
	}
L35:
	;
	v111 = int32(1)
	v112 = v97 + v111
	v113 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v96)+1)))
	if v113 != 0 {
		v96 = v96 + v111
		v97 = v112
		v98 = v113
		goto L32
	} else {
		goto L38
	}
L36:
	;
	v106 = F_tolower(m, v104)
	mBase = m.M
	v107 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v97))))
	v108 = F_tolower(m, v107)
	mBase = m.M
	if v106 == v108 {
		goto L35
	} else {
		goto L37
	}
L37:
	;
	v110 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v96))))
	v118 = v97
	v119 = v110
	goto L31
L38:
	;
	goto L33
L39:
	;
	v131 = v79 + int32(1)
	v132 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v132 <= v131 {
		goto L26
	} else {
		goto L40
	}
L40:
	;
	v134 = *(*int32)(unsafe.Add(mBase, uint32(v11)+12))
	if v134 != 0 {
		goto L26
	} else {
		goto L41
	}
L41:
	;
	v135 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v139 = *(*int32)(unsafe.Add(mBase, uint32(v135+v131<<(uint(int32(2))%32))))
	v143 = F_getLongFromObjectOrReply(m, l0, v139, v11+int32(12), int32(0))
	mBase = m.M
	v144 = m.ExcPending
	if v144 != 0 {
		goto L2
	} else {
		goto L42
	}
L42:
	;
	if v143 != 0 {
		goto L1
	} else {
		goto L43
	}
L43:
	;
	v145 = *(*int32)(unsafe.Add(mBase, uint32(v11)+12))
	if int32(0) < v145 {
		v267 = v131
		v268 = v80
		v269 = v81
		goto L25
	} else {
		goto L44
	}
L44:
	;
	F_addReplyError(m, l0, int32(_a1296))
	mBase = m.M
	v150 = m.ExcPending
	if v150 != 0 {
		goto L2
	} else {
		goto L45
	}
L45:
	;
	goto L1
L46:
	;
	v216 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v218 = *(*int32)(unsafe.Add(mBase, uint32(v216+v87)))
	v219 = F_objectGetVal(m, v218)
	mBase = m.M
	v220 = int32(_a409)
	v223 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v219))))
	if v223 != 0 {
		goto L66
	} else {
		goto L67
	}
L47:
	;
	if v191-v193 != 0 {
		goto L46
	} else {
		goto L59
	}
L48:
	;
	v191 = F_tolower(m, v187)
	mBase = m.M
	v192 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v188))))
	v193 = F_tolower(m, v192)
	mBase = m.M
	goto L47
L49:
	;
	v161 = v155
	v162 = v156
	v163 = v159
	goto L52
L50:
	;
	v187 = int32(0)
	v188 = v156
	goto L48
L51:
	;
	v187 = v184 & int32(255)
	v188 = v183
	goto L48
L52:
	;
	v165 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v162))))
	if v165 == int32(0) {
		v183 = v162
		v184 = v163
		goto L51
	} else {
		goto L54
	}
L53:
	;
	v183 = v177
	v184 = int32(0)
	goto L51
L54:
	;
	v169 = v163 & int32(255)
	if v169 == v165 {
		goto L55
	} else {
		goto L56
	}
L55:
	;
	v176 = int32(1)
	v177 = v162 + v176
	v178 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v161)+1)))
	if v178 != 0 {
		v161 = v161 + v176
		v162 = v177
		v163 = v178
		goto L52
	} else {
		goto L58
	}
L56:
	;
	v171 = F_tolower(m, v169)
	mBase = m.M
	v172 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v162))))
	v173 = F_tolower(m, v172)
	mBase = m.M
	if v171 == v173 {
		goto L55
	} else {
		goto L57
	}
L57:
	;
	v175 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v161))))
	v183 = v162
	v184 = v175
	goto L51
L58:
	;
	goto L53
L59:
	;
	v196 = v79 + int32(2)
	v197 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v197 <= v196 {
		goto L46
	} else {
		goto L60
	}
L60:
	;
	if v80 != 0 {
		goto L46
	} else {
		goto L61
	}
L61:
	;
	v199 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v203 = *(*int32)(unsafe.Add(mBase, uint32(v199+v196<<(uint(int32(2))%32))))
	v207 = F_getLongFromObjectOrReply(m, l0, v203, v11+int32(8), int32(0))
	mBase = m.M
	v208 = m.ExcPending
	if v208 != 0 {
		goto L2
	} else {
		goto L62
	}
L62:
	;
	if v207 != 0 {
		goto L1
	} else {
		goto L63
	}
L63:
	;
	v209 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v213 = *(*int32)(unsafe.Add(mBase, uint32(v209+v87+int32(4))))
	v214 = F_objectGetVal(m, v213)
	mBase = m.M
	v267 = v196
	v268 = v214
	v269 = v81
	goto L25
L64:
	;
	if v255-v257|v81 == int32(0) {
		v267 = v79
		v268 = v80
		v269 = int32(1)
		goto L25
	} else {
		goto L76
	}
L65:
	;
	v255 = F_tolower(m, v251)
	mBase = m.M
	v256 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v252))))
	v257 = F_tolower(m, v256)
	mBase = m.M
	goto L64
L66:
	;
	v225 = v219
	v226 = v220
	v227 = v223
	goto L69
L67:
	;
	v251 = int32(0)
	v252 = v220
	goto L65
L68:
	;
	v251 = v248 & int32(255)
	v252 = v247
	goto L65
L69:
	;
	v229 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v226))))
	if v229 == int32(0) {
		v247 = v226
		v248 = v227
		goto L68
	} else {
		goto L71
	}
L70:
	;
	v247 = v241
	v248 = int32(0)
	goto L68
L71:
	;
	v233 = v227 & int32(255)
	if v233 == v229 {
		goto L72
	} else {
		goto L73
	}
L72:
	;
	v240 = int32(1)
	v241 = v226 + v240
	v242 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v225)+1)))
	if v242 != 0 {
		v225 = v225 + v240
		v226 = v241
		v227 = v242
		goto L69
	} else {
		goto L75
	}
L73:
	;
	v235 = F_tolower(m, v233)
	mBase = m.M
	v236 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v226))))
	v237 = F_tolower(m, v236)
	mBase = m.M
	if v235 == v237 {
		goto L72
	} else {
		goto L74
	}
L74:
	;
	v239 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v225))))
	v247 = v226
	v248 = v239
	goto L68
L75:
	;
	goto L70
L76:
	;
	v264 = *(*int32)(unsafe.Add(mBase, _consts[86]))
	F_addReplyErrorObject(m, l0, v264)
	mBase = m.M
	v266 = m.ExcPending
	if v266 != 0 {
		goto L2
	} else {
		goto L77
	}
L77:
	;
	goto L1
L78:
	;
	goto L24
L79:
	;
	v292 = *(*int32)(unsafe.Add(mBase, _consts[130]))
	if v292 == int32(0) {
		goto L82
	} else {
		goto L83
	}
L80:
	;
	F_addReplyError(m, l0, int32(_a1297))
	mBase = m.M
	v290 = m.ExcPending
	if v290 != 0 {
		goto L2
	} else {
		goto L81
	}
L81:
	;
	goto L1
L82:
	;
	v299 = *(*int32)(unsafe.Add(mBase, _consts[158]))
	v300 = *(*int32)(unsafe.Add(mBase, uint32(v299)+20))
	if v300 != 0 {
		goto L85
	} else {
		goto L86
	}
L83:
	;
	F_addReplyError(m, l0, int32(_a1298))
	mBase = m.M
	v297 = m.ExcPending
	if v297 != 0 {
		goto L2
	} else {
		goto L84
	}
L84:
	;
	goto L1
L85:
	;
	if v280 == int32(0) {
		goto L91
	} else {
		goto L92
	}
L86:
	;
	F_addReplyError(m, l0, int32(_a1299))
	mBase = m.M
	v303 = m.ExcPending
	if v303 != 0 {
		goto L2
	} else {
		goto L87
	}
L87:
	;
	goto L1
L88:
	;
	v355 = *(*int64)(unsafe.Add(mBase, _consts[98]))
	goto L110
L89:
	;
	v345 = *(*int32)(unsafe.Add(mBase, _consts[6]))
	if int32(2) < v345 {
		goto L88
	} else {
		goto L108
	}
L90:
	;
	v314 = *(*int32)(unsafe.Add(mBase, uint32(v11)+8))
	v315 = F_findReplica(m, v279, v314)
	mBase = m.M
	v316 = m.ExcPending
	if v316 != 0 {
		goto L2
	} else {
		goto L99
	}
L91:
	;
	if v279 == int32(0) {
		goto L89
	} else {
		goto L97
	}
L92:
	;
	v306 = *(*int32)(unsafe.Add(mBase, uint32(v11)+12))
	if v306 == int32(0) {
		goto L93
	} else {
		goto L94
	}
L93:
	;
	F_addReplyError(m, l0, int32(_a1300))
	mBase = m.M
	v311 = m.ExcPending
	if v311 != 0 {
		goto L2
	} else {
		goto L96
	}
L94:
	;
	if v279 != 0 {
		goto L90
	} else {
		goto L95
	}
L95:
	;
	goto L93
L96:
	;
	goto L1
L97:
	;
	goto L90
L98:
	;
	v320 = *(*int32)(unsafe.Add(mBase, uint32(v315)+104))
	v321 = *(*int32)(unsafe.Add(mBase, uint32(v320)))
	if v321 == int32(9) {
		goto L102
	} else {
		goto L103
	}
L99:
	;
	if v315 != 0 {
		goto L98
	} else {
		goto L100
	}
L100:
	;
	F_addReplyError(m, l0, int32(_a1301))
	mBase = m.M
	v319 = m.ExcPending
	if v319 != 0 {
		goto L2
	} else {
		goto L101
	}
L101:
	;
	goto L1
L102:
	;
	v328 = F_zstrdup(m, v279)
	mBase = m.M
	v329 = m.ExcPending
	if v329 != 0 {
		goto L2
	} else {
		goto L105
	}
L103:
	;
	F_addReplyError(m, l0, int32(_a1302))
	mBase = m.M
	v326 = m.ExcPending
	if v326 != 0 {
		goto L2
	} else {
		goto L104
	}
L104:
	;
	goto L1
L105:
	;
	*(*int32)(unsafe.Add(mBase, _consts[713])) = v328
	v331 = int32(_a44)
	v332 = *(*int32)(unsafe.Add(mBase, uint32(v11)+8))
	*(*int32)(unsafe.Add(mBase, _consts[716])) = v332
	v335 = *(*int32)(unsafe.Add(mBase, _consts[6]))
	if int32(2) < v335 {
		goto L88
	} else {
		goto L106
	}
L106:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+4)) = v332
	*(*int32)(unsafe.Add(mBase, uint32(v11))) = v279
	F__serverLog(m, int32(2), int32(_a1303), v11)
	mBase = m.M
	v343 = m.ExcPending
	if v343 != 0 {
		goto L2
	} else {
		goto L107
	}
L107:
	;
	goto L88
L108:
	;
	F__serverLog(m, int32(2), int32(_a1304), int32(0))
	mBase = m.M
	v352 = m.ExcPending
	if v352 != 0 {
		goto L2
	} else {
		goto L109
	}
L109:
	;
	goto L88
L110:
	;
	v356 = *(*int32)(unsafe.Add(mBase, uint32(v11)+12))
	if v356 == int32(0) {
		goto L111
	} else {
		goto L112
	}
L111:
	;
	v363 = int32(_a44)
	*(*int32)(unsafe.Add(mBase, _consts[658])) = int32(1)
	*(*int32)(unsafe.Add(mBase, _consts[714])) = v280
	F_pauseActions(m, int32(2), int64(9223372036854775807), int32(29))
	mBase = m.M
	v372 = m.ExcPending
	if v372 != 0 {
		goto L2
	} else {
		goto L113
	}
L112:
	;
	*(*int64)(unsafe.Add(mBase, _consts[715])) = v355 + base.I64_extend_i32_s(v356)
	goto L111
L113:
	;
	v374 = *(*int32)(unsafe.Add(mBase, _consts[84]))
	F_addReply(m, l0, v374)
	mBase = m.M
	v376 = m.ExcPending
	if v376 != 0 {
		goto L2
	} else {
		goto L114
	}
L114:
	;
	goto L1
L115:
	;
	F_abortFailover(m, int32(_a1305))
	mBase = m.M
	v384 = m.ExcPending
	if v384 != 0 {
		goto L2
	} else {
		goto L118
	}
L116:
	;
	F_addReplyError(m, l0, int32(_a1306))
	mBase = m.M
	v381 = m.ExcPending
	if v381 != 0 {
		goto L2
	} else {
		goto L117
	}
L117:
	;
	goto L1
L118:
	;
	v386 = *(*int32)(unsafe.Add(mBase, _consts[84]))
	F_addReply(m, l0, v386)
	mBase = m.M
	v388 = m.ExcPending
	if v388 != 0 {
		goto L2
	} else {
		goto L119
	}
L119:
	;
	goto L1
}
func F_fcallGetCommandFlags(m *base.Module, l0 int32, l1 int64) int64 {
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
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int64
	_ = v19
	var v28 int64
	_ = v28
	var v35 int64
	_ = v35
	var v39 int64
	_ = v39
	v5 = *(*int32)(unsafe.Add(mBase, _consts[423]))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(v5)+4))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(v7)+4))
	v9 = F_objectGetVal(m, v8)
	mBase = m.M
	v10 = F_dictFind(m, v6, v9)
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return int64(0)
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(l0)+324)) = v10
		if v10 == int32(0) {
			v39 = l1
		} else {
			v17 = *(*int32)(unsafe.Add(mBase, uint32(v10)+8))
			v18 = *(*int32)(unsafe.Add(mBase, uint32(v17)))
			v19 = *(*int64)(unsafe.Add(mBase, uint32(v18)+24))
			v28 = l1 & int64(-66566)
			if v19&int64(3) == int64(0) {
				v35 = v28 | int64(4)
			} else {
				v35 = v28
			}
			v39 = v19<<(uint(int64(8))%64)&int64(1024) | v19&int64(1) | v35 ^ int64(1)
		}
		return v39
	}
}
func F_fcallroCommand(m *base.Module, l0 int32) {
	var v4 int32
	_ = v4
	F_fcallCommandGeneric(m, l0, int32(1))
	v4 = m.ExcPending
	if v4 != 0 {
		return
	} else {
		return
	}
}
func F_feedAppendOnlyFile(m *base.Module, l0 int32, l1 int32, l2 int32) {
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
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
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
	var v38 int32
	_ = v38
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v65 int32
	_ = v65
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v76 int32
	_ = v76
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
	var v91 int32
	_ = v91
	var v95 int32
	_ = v95
	var v97 int32
	_ = v97
	var v105 int32
	_ = v105
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v119 int32
	_ = v119
	var v125 int32
	_ = v125
	var v129 int32
	_ = v129
	var v133 int32
	_ = v133
	var v140 int32
	_ = v140
	var v143 int32
	_ = v143
	var v146 int32
	_ = v146
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v159 int32
	_ = v159
	var v167 int32
	_ = v167
	v7 = m.G0
	v9 = v7 - int32(96)
	m.G0 = v9
	v11 = F_sdsempty(m)
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
	if l0 == int32(-1) {
		goto L4
	} else {
		goto L5
	}
L3:
	;
	F__serverAssert(m, int32(_a135), int32(_a85), int32(1449))
	mBase = m.M
	v167 = m.ExcPending
	if v167 != 0 {
		goto L1
	} else {
		goto L49
	}
L4:
	;
	v21 = *(*int32)(unsafe.Add(mBase, _consts[64]))
	if v21 == int32(0) {
		v33 = v11
		goto L8
	} else {
		goto L9
	}
L5:
	;
	if l0 < int32(0) {
		goto L3
	} else {
		goto L6
	}
L6:
	;
	v18 = *(*int32)(unsafe.Add(mBase, _consts[65]))
	if v18 <= l0 {
		goto L3
	} else {
		goto L7
	}
L7:
	;
	goto L4
L8:
	;
	if l0 == int32(-1) {
		v115 = v33
		goto L14
	} else {
		goto L15
	}
L9:
	;
	v25 = F_genAofTimestampAnnotationIfNeeded(m, int32(0))
	mBase = m.M
	v26 = m.ExcPending
	if v26 != 0 {
		goto L1
	} else {
		goto L10
	}
L10:
	;
	if v25 == int32(0) {
		v33 = v11
		goto L8
	} else {
		goto L11
	}
L11:
	;
	v29 = F_sdscatsds(m, v11, v25)
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		goto L1
	} else {
		goto L12
	}
L12:
	;
	F_sdsfree(m, v25)
	mBase = m.M
	v32 = m.ExcPending
	if v32 != 0 {
		goto L1
	} else {
		goto L13
	}
L13:
	;
	v33 = v29
	goto L8
L14:
	;
	v116 = F_catAppendOnlyGenericCommand(m, v115, l2, l1)
	mBase = m.M
	v117 = m.ExcPending
	if v117 != 0 {
		goto L1
	} else {
		goto L35
	}
L15:
	;
	v38 = *(*int32)(unsafe.Add(mBase, _consts[51]))
	if l0 == v38 {
		v115 = v33
		goto L14
	} else {
		goto L16
	}
L16:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = l0
	v47 = F_snprintf(m, v9+int32(32), int32(64), int32(_a77), v9+int32(16))
	mBase = m.M
	v48 = m.ExcPending
	if v48 != 0 {
		goto L1
	} else {
		goto L17
	}
L17:
	;
	v50 = v9 + int32(32)
	if v50&int32(3) == int32(0) {
		v72 = v50
		goto L20
	} else {
		goto L21
	}
L18:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9))) = v105
	*(*int32)(unsafe.Add(mBase, uint32(v9)+4)) = v9 + int32(32)
	v111 = F_sdscatprintf(m, v33, int32(_a136), v9)
	mBase = m.M
	v112 = m.ExcPending
	if v112 != 0 {
		goto L1
	} else {
		goto L34
	}
L19:
	;
	v105 = v97 - v50
	goto L18
L20:
	;
	v76 = v72
	goto L28
L21:
	;
	v58 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v50))))
	if v58 != 0 {
		goto L22
	} else {
		goto L23
	}
L22:
	;
	v61 = v50
	goto L24
L23:
	;
	v105 = v50 - v50
	goto L18
L24:
	;
	v65 = v61 + int32(1)
	if v65&int32(3) == int32(0) {
		v72 = v65
		goto L20
	} else {
		goto L26
	}
L26:
	;
	v70 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v65))))
	if v70 != 0 {
		v61 = v65
		goto L24
	} else {
		goto L27
	}
L27:
	;
	v97 = v65
	goto L19
L28:
	;
	v82 = *(*int32)(unsafe.Add(mBase, uint32(v76)))
	v85 = int32(-2139062144)
	if (int32(16843008)-v82|v82)&v85 == v85 {
		v76 = v76 + int32(4)
		goto L28
	} else {
		goto L30
	}
L29:
	;
	v91 = v76
	goto L31
L30:
	;
	goto L29
L31:
	;
	v95 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v91))))
	if v95 != 0 {
		v91 = v91 + int32(1)
		goto L31
	} else {
		goto L33
	}
L32:
	;
	v97 = v91
	goto L19
L33:
	;
	goto L32
L34:
	;
	*(*int32)(unsafe.Add(mBase, _consts[51])) = l0
	v115 = v111
	goto L14
L35:
	;
	v119 = *(*int32)(unsafe.Add(mBase, _consts[28]))
	if v119 == int32(1) {
		goto L37
	} else {
		goto L38
	}
L36:
	;
	F_sdsfree(m, v116)
	mBase = m.M
	v159 = m.ExcPending
	if v159 != 0 {
		goto L1
	} else {
		goto L48
	}
L37:
	;
	v129 = *(*int32)(unsafe.Add(mBase, _consts[34]))
	v133 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v116+int32(-1)))))
	switch v133 & int32(7) {
	case 0:
		goto L46
	case 1:
		goto L45
	case 2:
		goto L44
	case 3:
		goto L43
	case 4:
		goto L42
	default:
		v150 = int32(0)
		goto L41
	}
L38:
	;
	if v119 != int32(2) {
		goto L36
	} else {
		goto L39
	}
L39:
	;
	v125 = *(*int32)(unsafe.Add(mBase, _consts[66]))
	if v125 != int32(2) {
		goto L36
	} else {
		goto L40
	}
L40:
	;
	goto L37
L41:
	;
	v152 = F_sdscatlen(m, v129, v116, v150)
	mBase = m.M
	v153 = m.ExcPending
	if v153 != 0 {
		goto L1
	} else {
		goto L47
	}
L42:
	;
	v149 = *(*int32)(unsafe.Add(mBase, uint32(v116+int32(-17))))
	v150 = v149
	goto L41
L43:
	;
	v146 = *(*int32)(unsafe.Add(mBase, uint32(v116+int32(-9))))
	v150 = v146
	goto L41
L44:
	;
	v143 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v116+int32(-5)))))
	v150 = v143
	goto L41
L45:
	;
	v140 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v116+int32(-3)))))
	v150 = v140
	goto L41
L46:
	;
	v150 = int32(base.Ui32(v133) >> (uint(int32(3)) % 32))
	goto L41
L47:
	;
	*(*int32)(unsafe.Add(mBase, _consts[34])) = v152
	goto L36
L48:
	;
	m.G0 = v9 + int32(96)
	return
L49:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_feedReplicationBuffer(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v59 int32
	_ = v59
	var v61 int64
	_ = v61
	var v62 int64
	_ = v62
	var v66 int32
	_ = v66
	var v67 int64
	_ = v67
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v86 int32
	_ = v86
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v94 int32
	_ = v94
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v107 int32
	_ = v107
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v122 int32
	_ = v122
	var v125 int32
	_ = v125
	var v128 int64
	_ = v128
	var v129 int64
	_ = v129
	var v132 int32
	_ = v132
	var v134 int64
	_ = v134
	var v143 int32
	_ = v143
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v152 int32
	_ = v152
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v165 int64
	_ = v165
	var v166 int64
	_ = v166
	var v170 int32
	_ = v170
	var v171 int64
	_ = v171
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v180 int32
	_ = v180
	var v181 int64
	_ = v181
	var v185 int32
	_ = v185
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
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v212 int32
	_ = v212
	var v215 int32
	_ = v215
	var v221 int32
	_ = v221
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v229 int32
	_ = v229
	var v232 int32
	_ = v232
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v259 int32
	_ = v259
	var v260 int32
	_ = v260
	var v264 int32
	_ = v264
	var v273 int32
	_ = v273
	var v274 int32
	_ = v274
	var v275 int32
	_ = v275
	var v277 int32
	_ = v277
	var v281 int32
	_ = v281
	var v282 int64
	_ = v282
	var v283 int64
	_ = v283
	var v285 int64
	_ = v285
	var v287 int64
	_ = v287
	var v290 int64
	_ = v290
	var v292 int64
	_ = v292
	var v294 int64
	_ = v294
	var v296 int64
	_ = v296
	var v319 int32
	_ = v319
	var v320 int32
	_ = v320
	var v324 int32
	_ = v324
	var v325 int32
	_ = v325
	var v327 int32
	_ = v327
	var v333 int32
	_ = v333
	var v354 int32
	_ = v354
	v11 = m.G0
	v13 = v11 - int32(16)
	m.G0 = v13
	v16 = *(*int32)(unsafe.Add(mBase, _consts[370]))
	if v16 == int32(0) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	F__serverAssert(m, int32(_a1192), int32(_a1190), int32(541))
	mBase = m.M
	v354 = m.ExcPending
	if v354 != 0 {
		goto L4
	} else {
		goto L76
	}
L2:
	;
	m.G0 = v13 + int32(16)
	return
L3:
	;
	F_clusterSlotStatsIncrNetworkBytesOutForReplication(m, base.I64_extend_i32_u(l1))
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	return
L5:
	;
	if l1 == int32(0) {
		goto L2
	} else {
		goto L6
	}
L6:
	;
	v24 = l0
	v25 = l1
	goto L7
L7:
	;
	v35 = *(*int32)(unsafe.Add(mBase, _consts[217]))
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v35)+4))
	if v36 == int32(0) {
		goto L11
	} else {
		goto L12
	}
L8:
	;
	goto L2
L9:
	;
	v196 = *(*int32)(unsafe.Add(mBase, _consts[158]))
	v197 = *(*int32)(unsafe.Add(mBase, uint32(v196)))
	*(*int32)(unsafe.Add(mBase, uint32(v13)+4)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v13))) = v197
	goto L48
L10:
	;
	v98 = int32(16384)
	if base.Ui32(v98) < base.Ui32(v90) {
		goto L25
	} else {
		goto L26
	}
L11:
	;
	v86 = int32(0)
	v89 = v24
	v90 = v25
	v91 = v86
	v94 = v86
	v97 = int32(1)
	goto L10
L12:
	;
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v36)+8))
	if v39 == int32(0) {
		goto L11
	} else {
		goto L13
	}
L13:
	;
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v39)+24))
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v39)+28))
	if base.Ui32(v43) < base.Ui32(v42) {
		goto L16
	} else {
		goto L17
	}
L14:
	;
	v186 = v70
	v187 = int32(0)
	v188 = v36
	v191 = v43
	v193 = v71
	goto L9
L15:
	;
	v89 = v75
	v90 = v76
	v91 = v77
	v94 = v80
	v97 = int32(0)
	goto L10
L16:
	;
	v50 = v42 - v43
	if base.Ui32(v50) < base.Ui32(v25) {
		goto L18
	} else {
		goto L19
	}
L17:
	;
	v45 = int32(0)
	v75 = v24
	v76 = v25
	v77 = v45
	v80 = v45
	goto L15
L18:
	;
	v52 = v50
	goto L20
L19:
	;
	v52 = v25
	goto L20
L20:
	;
	if v52 == int32(0) {
		goto L22
	} else {
		goto L23
	}
L21:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v39)+28)) = v52 + v43
	v59 = int32(_a44)
	v61 = *(*int64)(unsafe.Add(mBase, _consts[40]))
	v62 = base.I64_extend_i32_u(v52)
	*(*int64)(unsafe.Add(mBase, _consts[40])) = v61 + v62
	v66 = *(*int32)(unsafe.Add(mBase, _consts[370]))
	v67 = *(*int64)(unsafe.Add(mBase, uint32(v66)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v66)+16)) = v67 + v62
	v70 = v24 + v52
	v71 = int32(0)
	v72 = v25 - v52
	if v72 == v71 {
		goto L14
	} else {
		goto L24
	}
L22:
	;
	goto L21
L23:
	;
	v55 = F__emscripten_memcpy_bulkmem(m, v39+v43+int32(32), v24, v52)
	mBase = m.M
	goto L22
L24:
	;
	v75 = v70
	v76 = v72
	v77 = v36
	v80 = v43
	goto L15
L25:
	;
	v101 = v90
	goto L27
L26:
	;
	v101 = v98
	goto L27
L27:
	;
	v103 = *(*int32)(unsafe.Add(mBase, _consts[287]))
	v104 = int32(262159)
	if base.Ui32(v104) < base.Ui32(v103) {
		goto L28
	} else {
		goto L29
	}
L28:
	;
	v107 = v103
	goto L30
L29:
	;
	v107 = v104
	goto L30
L30:
	;
	v109 = int32(base.Ui32(v107) >> (uint(int32(4)) % 32))
	if base.Ui32(v101) < base.Ui32(v109) {
		goto L31
	} else {
		goto L32
	}
L31:
	;
	v111 = v101
	goto L33
L32:
	;
	v111 = v109
	goto L33
L33:
	;
	v116 = F_zmalloc_usable(m, v111+int32(32), v13+int32(8))
	mBase = m.M
	v117 = m.ExcPending
	if v117 != 0 {
		goto L4
	} else {
		goto L34
	}
L34:
	;
	v118 = *(*int32)(unsafe.Add(mBase, uint32(v13)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v116))) = int32(0)
	v122 = v118 + int32(-32)
	*(*int32)(unsafe.Add(mBase, uint32(v116)+24)) = v122
	if base.Ui32(v122) < base.Ui32(v90) {
		goto L35
	} else {
		goto L36
	}
L35:
	;
	v125 = v122
	goto L37
L36:
	;
	v125 = v90
	goto L37
L37:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v116)+28)) = v125
	v128 = *(*int64)(unsafe.Add(mBase, _consts[40]))
	v129 = int64(1)
	*(*int64)(unsafe.Add(mBase, uint32(v116)+16)) = v128 + v129
	v132 = int32(0)
	v134 = *(*int64)(unsafe.Add(mBase, _consts[653]))
	*(*int64)(unsafe.Add(mBase, _consts[653])) = v134 + v129
	*(*int64)(unsafe.Add(mBase, uint32(v116)+8)) = v134
	if v125 == v132 {
		goto L39
	} else {
		goto L40
	}
L38:
	;
	v146 = *(*int32)(unsafe.Add(mBase, _consts[217]))
	v147 = F_listAddNodeTail(m, v146, v116)
	mBase = m.M
	v148 = m.ExcPending
	if v148 != 0 {
		goto L4
	} else {
		goto L41
	}
L39:
	;
	goto L38
L40:
	;
	v143 = F__emscripten_memcpy_bulkmem(m, v116+int32(32), v89, v125)
	mBase = m.M
	goto L39
L41:
	;
	v149 = int32(_a44)
	v150 = *(*int32)(unsafe.Add(mBase, uint32(v13)+8))
	v152 = *(*int32)(unsafe.Add(mBase, _consts[288]))
	*(*int32)(unsafe.Add(mBase, _consts[288])) = v150 + v152 + int32(12)
	if v91 != 0 {
		v161 = v91
		v162 = v94
		goto L42
	} else {
		goto L43
	}
L42:
	;
	v163 = int32(_a44)
	v165 = *(*int64)(unsafe.Add(mBase, _consts[40]))
	v166 = base.I64_extend_i32_u(v125)
	*(*int64)(unsafe.Add(mBase, _consts[40])) = v165 + v166
	v170 = *(*int32)(unsafe.Add(mBase, _consts[370]))
	v171 = *(*int64)(unsafe.Add(mBase, uint32(v170)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v170)+16)) = v171 + v166
	v174 = v90 - v125
	v175 = v89 + v125
	v176 = int32(1)
	if v97 == int32(0) {
		v186 = v175
		v187 = v174
		v188 = v161
		v191 = v162
		v193 = v176
		goto L9
	} else {
		goto L44
	}
L43:
	;
	v158 = *(*int32)(unsafe.Add(mBase, _consts[217]))
	v159 = *(*int32)(unsafe.Add(mBase, uint32(v158)+4))
	v161 = v159
	v162 = int32(0)
	goto L42
L44:
	;
	v180 = *(*int32)(unsafe.Add(mBase, _consts[654]))
	v181 = *(*int64)(unsafe.Add(mBase, uint32(v180)+8))
	goto L45
L45:
	;
	if v181 == int64(0) {
		v186 = v175
		v187 = v174
		v188 = v161
		v191 = v162
		v193 = v176
		goto L9
	} else {
		goto L46
	}
L46:
	;
	F_backfillRdbReplicasToPsyncWait(m)
	mBase = m.M
	v185 = m.ExcPending
	if v185 != 0 {
		goto L4
	} else {
		goto L47
	}
L47:
	;
	v186 = v175
	v187 = v174
	v188 = v161
	v191 = v162
	v193 = v176
	goto L9
L48:
	;
	goto L50
L49:
	;
	v256 = *(*int32)(unsafe.Add(mBase, _consts[370]))
	v257 = *(*int32)(unsafe.Add(mBase, uint32(v256)))
	if v257 != 0 {
		goto L66
	} else {
		goto L67
	}
L50:
	;
	v212 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
	if v212 == int32(0) {
		goto L53
	} else {
		goto L54
	}
L52:
	;
	if v212 == int32(0) {
		goto L49
	} else {
		goto L55
	}
L53:
	;
	goto L52
L54:
	;
	v215 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
	v221 = *(*int32)(unsafe.Add(mBase, uint32(v212+base.B2i32(v215 == int32(0))<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v13))) = v221
	goto L53
L55:
	;
	v225 = *(*int32)(unsafe.Add(mBase, uint32(v212)+8))
	v226 = *(*int32)(unsafe.Add(mBase, uint32(v225)+204))
	if v226&int32(8192) != 0 {
		goto L57
	} else {
		goto L58
	}
L56:
	;
	v241 = *(*int32)(unsafe.Add(mBase, uint32(v240)+184))
	if v241 != 0 {
		goto L62
	} else {
		goto L63
	}
L57:
	;
	if v226&int32(33554432) == int32(0) {
		goto L50
	} else {
		goto L61
	}
L58:
	;
	v229 = *(*int32)(unsafe.Add(mBase, uint32(v225)+104))
	if v226&int32(33554432) != 0 {
		v240 = v229
		goto L56
	} else {
		goto L59
	}
L59:
	;
	v232 = *(*int32)(unsafe.Add(mBase, uint32(v229)))
	if v232 == int32(6) {
		goto L50
	} else {
		goto L60
	}
L60:
	;
	v240 = v229
	goto L56
L61:
	;
	v239 = *(*int32)(unsafe.Add(mBase, uint32(v225)+104))
	v240 = v239
	goto L56
L62:
	;
	if v193 == int32(0) {
		goto L50
	} else {
		goto L64
	}
L63:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v240)+188)) = v191
	*(*int32)(unsafe.Add(mBase, uint32(v240)+184)) = v188
	v244 = *(*int32)(unsafe.Add(mBase, uint32(v188)+8))
	v245 = *(*int32)(unsafe.Add(mBase, uint32(v244)))
	*(*int32)(unsafe.Add(mBase, uint32(v244))) = v245 + int32(1)
	goto L62
L64:
	;
	v253 = F_closeClientOnOutputBufferLimitReached(m, v225, int32(1))
	mBase = m.M
	v254 = m.ExcPending
	if v254 != 0 {
		goto L4
	} else {
		goto L65
	}
L65:
	;
	goto L50
L66:
	;
	if v193 == int32(0) {
		goto L69
	} else {
		goto L70
	}
L67:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v256))) = v188
	v259 = *(*int32)(unsafe.Add(mBase, uint32(v188)+8))
	v260 = *(*int32)(unsafe.Add(mBase, uint32(v259)))
	*(*int32)(unsafe.Add(mBase, uint32(v259))) = v260 + int32(1)
	v264 = int32(0)
	if base.B2i32(v191 == v264)&v193 == v264 {
		goto L1
	} else {
		goto L68
	}
L68:
	;
	goto L66
L69:
	;
	if v187 != 0 {
		v24 = v186
		v25 = v187
		goto L7
	} else {
		goto L75
	}
L70:
	;
	v273 = *(*int32)(unsafe.Add(mBase, _consts[217]))
	v274 = *(*int32)(unsafe.Add(mBase, uint32(v273)+4))
	v275 = *(*int32)(unsafe.Add(mBase, uint32(v256)+4))
	v277 = v275 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v256)+4)) = v277
	if base.Ui32(v277) < base.Ui32(int32(64)) {
		goto L71
	} else {
		goto L72
	}
L71:
	;
	F_incrementalTrimReplicationBacklog(m, int32(64))
	mBase = m.M
	v333 = m.ExcPending
	if v333 != 0 {
		goto L4
	} else {
		goto L74
	}
L72:
	;
	v281 = *(*int32)(unsafe.Add(mBase, uint32(v274)+8))
	v282 = *(*int64)(unsafe.Add(mBase, uint32(v281)+16))
	v283 = int64(56)
	v285 = int64(65280)
	v287 = int64(40)
	v290 = int64(16711680)
	v292 = int64(24)
	v294 = int64(4278190080)
	v296 = int64(8)
	*(*int64)(unsafe.Add(mBase, uint32(v13)+8)) = v282<<(uint(v283)%64) | v282&v285<<(uint(v287)%64) | (v282&v290<<(uint(v292)%64) | v282&v294<<(uint(v296)%64)) | (int64(base.Ui64(v282)>>(uint(v296)%64))&v294 | int64(base.Ui64(v282)>>(uint(v292)%64))&v290 | (int64(base.Ui64(v282)>>(uint(v287)%64))&v285 | int64(base.Ui64(v282)>>(uint(v283)%64))))
	v319 = *(*int32)(unsafe.Add(mBase, uint32(v256)+8))
	v320 = int32(8)
	v324 = F_raxInsert(m, v319, v13+v320, v320, v274, int32(0))
	mBase = m.M
	v325 = m.ExcPending
	if v325 != 0 {
		goto L4
	} else {
		goto L73
	}
L73:
	;
	v327 = *(*int32)(unsafe.Add(mBase, _consts[370]))
	*(*int32)(unsafe.Add(mBase, uint32(v327)+4)) = int32(0)
	goto L71
L74:
	;
	goto L69
L75:
	;
	goto L8
L76:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_ferror(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
	if int32(-1) < v4 {
		v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		v12 = v9
	} else {
		v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		v12 = v7
	}
	return int32(base.Ui32(v12)>>(uint(int32(5))%32)) & int32(1)
}
func F_ffc_bigint_long_mul(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v10 int64
	_ = v10
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v36 int32
	_ = v36
	var v40 int32
	_ = v40
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
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v57 int64
	_ = v57
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v79 int64
	_ = v79
	var v81 int32
	_ = v81
	var v96 int32
	_ = v96
	var v97 int64
	_ = v97
	var v99 int64
	_ = v99
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v103 int64
	_ = v103
	var v105 int64
	_ = v105
	var v107 int64
	_ = v107
	var v110 int32
	_ = v110
	var v111 int64
	_ = v111
	var v115 int64
	_ = v115
	var v118 int32
	_ = v118
	var v119 int64
	_ = v119
	var v123 int64
	_ = v123
	var v126 int64
	_ = v126
	var v128 int32
	_ = v128
	var v130 int32
	_ = v130
	var v133 int32
	_ = v133
	var v143 int64
	_ = v143
	var v146 int64
	_ = v146
	var v161 int32
	_ = v161
	var v170 int32
	_ = v170
	var v171 int64
	_ = v171
	var v188 int32
	_ = v188
	var v189 int64
	_ = v189
	var v191 int64
	_ = v191
	var v193 int32
	_ = v193
	var v196 int64
	_ = v196
	var v198 int32
	_ = v198
	var v211 int64
	_ = v211
	var v214 int64
	_ = v214
	var v237 int32
	_ = v237
	var v242 int32
	_ = v242
	var v253 int64
	_ = v253
	var v267 int32
	_ = v267
	var v270 int32
	_ = v270
	var v278 int32
	_ = v278
	var v287 int32
	_ = v287
	var v298 int64
	_ = v298
	var v299 int32
	_ = v299
	var v313 int32
	_ = v313
	var v316 int32
	_ = v316
	var v322 int32
	_ = v322
	var v324 int32
	_ = v324
	var v325 int32
	_ = v325
	var v327 int32
	_ = v327
	var v330 int32
	_ = v330
	var v331 int64
	_ = v331
	var v333 int32
	_ = v333
	var v340 int32
	_ = v340
	var v343 int32
	_ = v343
	var v353 int64
	_ = v353
	var v355 int32
	_ = v355
	var v372 int32
	_ = v372
	var v373 int64
	_ = v373
	var v375 int64
	_ = v375
	var v377 int32
	_ = v377
	var v378 int32
	_ = v378
	var v379 int64
	_ = v379
	var v381 int64
	_ = v381
	var v383 int64
	_ = v383
	var v386 int32
	_ = v386
	var v387 int64
	_ = v387
	var v391 int64
	_ = v391
	var v394 int32
	_ = v394
	var v395 int64
	_ = v395
	var v399 int64
	_ = v399
	var v402 int64
	_ = v402
	var v404 int32
	_ = v404
	var v406 int32
	_ = v406
	var v409 int32
	_ = v409
	var v419 int64
	_ = v419
	var v422 int64
	_ = v422
	var v437 int32
	_ = v437
	var v446 int32
	_ = v446
	var v447 int64
	_ = v447
	var v466 int32
	_ = v466
	var v467 int64
	_ = v467
	var v469 int64
	_ = v469
	var v471 int32
	_ = v471
	var v474 int64
	_ = v474
	var v476 int32
	_ = v476
	var v489 int64
	_ = v489
	var v492 int64
	_ = v492
	var v531 int64
	_ = v531
	var v538 int32
	_ = v538
	var v543 int32
	_ = v543
	var v544 int32
	_ = v544
	var v546 int32
	_ = v546
	var v550 int32
	_ = v550
	var v554 int32
	_ = v554
	var v555 int32
	_ = v555
	var v559 int32
	_ = v559
	var v562 int32
	_ = v562
	var v564 int32
	_ = v564
	var v571 int32
	_ = v571
	var v576 int32
	_ = v576
	var v580 int32
	_ = v580
	var v582 int32
	_ = v582
	var v588 int32
	_ = v588
	var v592 int32
	_ = v592
	var v596 int32
	_ = v596
	var v599 int32
	_ = v599
	var v618 int32
	_ = v618
	var v620 int32
	_ = v620
	var v621 int32
	_ = v621
	var v627 int32
	_ = v627
	var v628 int32
	_ = v628
	var v629 int32
	_ = v629
	var v630 int32
	_ = v630
	var v633 int32
	_ = v633
	var v635 int32
	_ = v635
	var v636 int32
	_ = v636
	var v638 int32
	_ = v638
	var v640 int32
	_ = v640
	var v641 int32
	_ = v641
	var v647 int32
	_ = v647
	var v648 int32
	_ = v648
	var v650 int32
	_ = v650
	var v655 int32
	_ = v655
	var v658 int32
	_ = v658
	var v664 int32
	_ = v664
	var v665 int32
	_ = v665
	var v666 int32
	_ = v666
	var v668 int32
	_ = v668
	var v671 int32
	_ = v671
	var v675 int32
	_ = v675
	var v699 int32
	_ = v699
	var v701 int32
	_ = v701
	var v702 int32
	_ = v702
	var v708 int32
	_ = v708
	var v709 int32
	_ = v709
	var v711 int32
	_ = v711
	var v712 int32
	_ = v712
	var v720 int32
	_ = v720
	var v727 int32
	_ = v727
	var v729 int32
	_ = v729
	var v731 int32
	_ = v731
	var v733 int32
	_ = v733
	var v761 int32
	_ = v761
	var v762 int32
	_ = v762
	var v763 int32
	_ = v763
	var v764 int32
	_ = v764
	var v775 int32
	_ = v775
	var v778 int32
	_ = v778
	var v783 int32
	_ = v783
	var v794 int64
	_ = v794
	var v807 int32
	_ = v807
	var v812 int32
	_ = v812
	var v835 int32
	_ = v835
	var v843 int32
	_ = v843
	var v866 int32
	_ = v866
	var v871 int32
	_ = v871
	var v873 int32
	_ = v873
	var v881 int32
	_ = v881
	var v883 int32
	_ = v883
	var v915 int32
	_ = v915
	v3 = int32(0)
	v10 = int64(0)
	v27 = m.G0
	v29 = v27 - int32(1008)
	m.G0 = v29
	v31 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+500)))
	v36 = F__emscripten_memset_bulkmem(m, v29, base.I32_extend8_s(v3), int32(504))
	mBase = m.M
	goto L1
L1:
	;
	if base.Ui32(int32(125)) < base.Ui32(v31) {
		v48 = v3
		goto L2
	} else {
		goto L3
	}
L2:
	;
	v50 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v50 == int32(0) {
		v812 = v31
		goto L9
	} else {
		goto L10
	}
L3:
	;
	v40 = v31 << (uint(int32(2)) % 32)
	if v40 == int32(0) {
		v44 = v36
		goto L5
	} else {
		goto L6
	}
L4:
	;
	v45 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v44)+500)))
	v46 = v45 + v31
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+500)) = uint16(v46)
	v48 = v46
	goto L2
L5:
	;
	goto L4
L6:
	;
	v43 = F__emscripten_memcpy_bulkmem(m, v36, l0, v40)
	mBase = m.M
	v44 = v43
	goto L5
L7:
	;
	m.G0 = v36 + int32(1008)
	return v915
L8:
	;
	v915 = int32(0)
	goto L7
L9:
	;
	v835 = int32(1)
	if v812&int32(65535) == int32(0) {
		v915 = v835
		goto L7
	} else {
		goto L93
	}
L10:
	;
	v53 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	if v31 != 0 {
		goto L12
	} else {
		goto L13
	}
L11:
	;
	if v50 == int32(1) {
		v812 = v242
		goto L9
	} else {
		goto L27
	}
L12:
	;
	v56 = v31 & int32(3)
	v57 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v53))))
	if base.Ui32(int32(4)) <= base.Ui32(v31) {
		goto L15
	} else {
		goto L16
	}
L13:
	;
	v242 = int32(0)
	v253 = v10
	goto L11
L14:
	;
	if v56 == int32(0) {
		v211 = v143
		v214 = v146
		goto L20
	} else {
		goto L21
	}
L15:
	;
	v66 = int32(0)
	v69 = v66
	v79 = int64(0)
	v81 = v66
	goto L17
L16:
	;
	v133 = int32(0)
	v143 = int64(0)
	v146 = v10
	goto L14
L17:
	;
	v96 = l0 + v69<<(uint(int32(2))%32)
	v97 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v96))))
	v99 = v97*v57 + v79
	*(*uint32)(unsafe.Add(mBase, uint32(v96))) = uint32(v99)
	v101 = int32(4)
	v102 = v96 + v101
	v103 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v102))))
	v105 = int64(32)
	v107 = v103*v57 + int64(base.Ui64(v99)>>(uint(v105)%64))
	*(*uint32)(unsafe.Add(mBase, uint32(v102))) = uint32(v107)
	v110 = v96 + int32(8)
	v111 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v110))))
	v115 = v111*v57 + int64(base.Ui64(v107)>>(uint(v105)%64))
	*(*uint32)(unsafe.Add(mBase, uint32(v110))) = uint32(v115)
	v118 = v96 + int32(12)
	v119 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v118))))
	v123 = v119*v57 + int64(base.Ui64(v115)>>(uint(v105)%64))
	*(*uint32)(unsafe.Add(mBase, uint32(v118))) = uint32(v123)
	v126 = int64(base.Ui64(v123) >> (uint(v105) % 64))
	v128 = v69 + v101
	v130 = v81 + v101
	if v130 != v31&int32(65532) {
		v69 = v128
		v79 = v126
		v81 = v130
		goto L17
	} else {
		goto L19
	}
L18:
	;
	v133 = v128
	v143 = v126
	v146 = v123
	goto L14
L19:
	;
	goto L18
L20:
	;
	if base.Ui64(v214) < base.Ui64(int64(4294967296)) {
		v242 = v31
		v253 = v214
		goto L11
	} else {
		goto L25
	}
L21:
	;
	v161 = v133
	v170 = int32(0)
	v171 = v143
	goto L22
L22:
	;
	v188 = l0 + v161<<(uint(int32(2))%32)
	v189 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v188))))
	v191 = v189*v57 + v171
	*(*uint32)(unsafe.Add(mBase, uint32(v188))) = uint32(v191)
	v193 = int32(1)
	v196 = int64(base.Ui64(v191) >> (uint(int64(32)) % 64))
	v198 = v170 + v193
	if v198 != v56 {
		v161 = v161 + v193
		v170 = v198
		v171 = v196
		goto L22
	} else {
		goto L24
	}
L23:
	;
	v211 = v196
	v214 = v191
	goto L20
L24:
	;
	goto L23
L25:
	;
	if base.Ui32(int32(124)) < base.Ui32(v31) {
		v915 = int32(0)
		goto L7
	} else {
		goto L26
	}
L26:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0+v31<<(uint(int32(2))%32)))) = base.I32_wrap_i64(v211)
	v237 = v31 + int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(l0)+500)) = uint16(v237)
	v242 = v237
	v253 = v214
	goto L11
L27:
	;
	v267 = int32(2)
	if base.Ui32(v267) < base.Ui32(v50) {
		goto L28
	} else {
		goto L29
	}
L28:
	;
	v270 = v50
	goto L30
L29:
	;
	v270 = v267
	goto L30
L30:
	;
	v278 = v48 << (uint(int32(2)) % 32) & int32(65532)
	v287 = v242
	v298 = v253
	v299 = int32(1)
	goto L31
L31:
	;
	v313 = *(*int32)(unsafe.Add(mBase, uint32(v53+v299<<(uint(int32(2))%32))))
	if v313 == int32(0) {
		v783 = v287
		v794 = v298
		goto L33
	} else {
		goto L34
	}
L32:
	;
	v812 = v783
	goto L9
L33:
	;
	v807 = v299 + int32(1)
	if v807 != v270 {
		v287 = v783
		v298 = v794
		v299 = v807
		goto L31
	} else {
		goto L92
	}
L34:
	;
	v316 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v36)+1004)) = uint16(v316)
	if base.Ui32(int32(125)) < base.Ui32(v48&int32(65535)) {
		goto L8
	} else {
		goto L35
	}
L35:
	;
	if v278 == int32(0) {
		goto L37
	} else {
		goto L38
	}
L36:
	;
	v324 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v36)+1004)))
	v325 = v324 + v48
	v327 = v325 & int32(65535)
	if v327 != 0 {
		goto L40
	} else {
		goto L41
	}
L37:
	;
	goto L36
L38:
	;
	v322 = F__emscripten_memcpy_bulkmem(m, v36+int32(504), v36, v278)
	mBase = m.M
	goto L37
L39:
	;
	v543 = int32(65535)
	v544 = v538 & v543
	v546 = v287 & v543
	if base.Ui32(v546) < base.Ui32(v299) {
		goto L55
	} else {
		goto L56
	}
L40:
	;
	v330 = v327 & int32(3)
	v331 = base.I64_extend_i32_u(v313)
	v333 = int32(0)
	if base.Ui32(v327) < base.Ui32(int32(4)) {
		v409 = v333
		v419 = int64(0)
		v422 = v298
		goto L42
	} else {
		goto L43
	}
L41:
	;
	v531 = v298
	v538 = int32(0)
	goto L39
L42:
	;
	if v330 == int32(0) {
		v489 = v419
		v492 = v422
		goto L47
	} else {
		goto L48
	}
L43:
	;
	v340 = int32(0)
	v343 = v340
	v353 = int64(0)
	v355 = v340
	goto L44
L44:
	;
	v372 = v36 + int32(504) + v343<<(uint(int32(2))%32)
	v373 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v372))))
	v375 = v373*v331 + v353
	*(*uint32)(unsafe.Add(mBase, uint32(v372))) = uint32(v375)
	v377 = int32(4)
	v378 = v372 + v377
	v379 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v378))))
	v381 = int64(32)
	v383 = v379*v331 + int64(base.Ui64(v375)>>(uint(v381)%64))
	*(*uint32)(unsafe.Add(mBase, uint32(v378))) = uint32(v383)
	v386 = v372 + int32(8)
	v387 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v386))))
	v391 = v387*v331 + int64(base.Ui64(v383)>>(uint(v381)%64))
	*(*uint32)(unsafe.Add(mBase, uint32(v386))) = uint32(v391)
	v394 = v372 + int32(12)
	v395 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v394))))
	v399 = v395*v331 + int64(base.Ui64(v391)>>(uint(v381)%64))
	*(*uint32)(unsafe.Add(mBase, uint32(v394))) = uint32(v399)
	v402 = int64(base.Ui64(v399) >> (uint(v381) % 64))
	v404 = v343 + v377
	v406 = v355 + v377
	if v406 != v327&int32(65532) {
		v343 = v404
		v353 = v402
		v355 = v406
		goto L44
	} else {
		goto L46
	}
L45:
	;
	v409 = v404
	v419 = v402
	v422 = v399
	goto L42
L46:
	;
	goto L45
L47:
	;
	if base.Ui64(v492) < base.Ui64(int64(4294967296)) {
		v531 = v492
		v538 = v325
		goto L39
	} else {
		goto L52
	}
L48:
	;
	v437 = v409
	v446 = v333
	v447 = v419
	goto L49
L49:
	;
	v466 = v36 + int32(504) + v437<<(uint(int32(2))%32)
	v467 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v466))))
	v469 = v467*v331 + v447
	*(*uint32)(unsafe.Add(mBase, uint32(v466))) = uint32(v469)
	v471 = int32(1)
	v474 = int64(base.Ui64(v469) >> (uint(int64(32)) % 64))
	v476 = v446 + v471
	if v476 != v330 {
		v437 = v437 + v471
		v446 = v476
		v447 = v474
		goto L49
	} else {
		goto L51
	}
L50:
	;
	v489 = v474
	v492 = v469
	goto L47
L51:
	;
	goto L50
L52:
	;
	if base.Ui32(int32(124)) < base.Ui32(v327) {
		goto L8
	} else {
		goto L53
	}
L53:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v36+int32(504)+v327<<(uint(int32(2))%32)))) = base.I32_wrap_i64(v489)
	v531 = v492
	v538 = v325 + int32(1)
	goto L39
L54:
	;
	if v544 == int32(0) {
		v783 = v576
		v794 = v531
		goto L33
	} else {
		goto L65
	}
L55:
	;
	v550 = v299 + v544
	if base.Ui32(int32(125)) < base.Ui32(v550) {
		goto L8
	} else {
		goto L58
	}
L56:
	;
	if base.Ui32(v544) <= base.Ui32(v546-v299) {
		v576 = v287
		goto L54
	} else {
		goto L57
	}
L57:
	;
	goto L55
L58:
	;
	if base.Ui32(v550) <= base.Ui32(v546) {
		goto L59
	} else {
		goto L60
	}
L59:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(l0)+500)) = uint16(v550)
	v576 = v550
	goto L54
L60:
	;
	v554 = int32(2)
	v555 = v546 << (uint(v554) % 32)
	v559 = l0 + int32(4) + v555
	v562 = v550<<(uint(v554)%32) + l0
	if base.Ui32(v562) < base.Ui32(v559) {
		goto L61
	} else {
		goto L62
	}
L61:
	;
	v564 = v559
	goto L63
L62:
	;
	v564 = v562
	goto L63
L63:
	;
	v571 = F__emscripten_memset_bulkmem(m, l0+v555, base.I32_extend8_s(int32(0)), (l0^int32(-1)-v555+v564)&int32(-4)+int32(4))
	mBase = m.M
	goto L64
L64:
	;
	goto L59
L65:
	;
	v580 = int32(1)
	v582 = int32(0)
	if v544 == v580 {
		v671 = v582
		v675 = v582
		goto L66
	} else {
		goto L67
	}
L66:
	;
	if v544&v580 == int32(0) {
		v720 = v675
		goto L77
	} else {
		goto L78
	}
L67:
	;
	v588 = int32(0)
	v592 = v588
	v596 = v588
	v599 = v588
	goto L68
L68:
	;
	v618 = int32(2)
	v620 = l0 + (v592+v299)<<(uint(v618)%32)
	v621 = *(*int32)(unsafe.Add(mBase, uint32(v620)))
	v627 = *(*int32)(unsafe.Add(mBase, uint32(v36+int32(504)+v592<<(uint(v618)%32))))
	v628 = v621 + v627
	v629 = int32(1)
	v630 = v628 + v629
	if v596&v629 != 0 {
		goto L70
	} else {
		goto L71
	}
L69:
	;
	v671 = v666
	v675 = v664
	goto L66
L70:
	;
	v633 = v630
	goto L72
L71:
	;
	v633 = v628
	goto L72
L72:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v620))) = v633
	v635 = int32(1)
	v636 = v592 | v635
	v638 = int32(2)
	v640 = l0 + (v636+v299)<<(uint(v638)%32)
	v641 = *(*int32)(unsafe.Add(mBase, uint32(v640)))
	v647 = *(*int32)(unsafe.Add(mBase, uint32(v36+int32(504)+v636<<(uint(v638)%32))))
	v648 = v641 + v647
	v650 = v648 + v635
	v655 = base.B2i32(base.Ui32(v628) < base.Ui32(v621)) | v596&base.B2i32(v630 == int32(0))
	if v655&v635 != 0 {
		goto L73
	} else {
		goto L74
	}
L73:
	;
	v658 = v650
	goto L75
L74:
	;
	v658 = v648
	goto L75
L75:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v640))) = v658
	v664 = base.B2i32(base.Ui32(v648) < base.Ui32(v641)) | v655&base.B2i32(v650 == int32(0))
	v665 = int32(2)
	v666 = v592 + v665
	v668 = v599 + v665
	if v668 != v544&int32(65534) {
		v592 = v666
		v596 = v664
		v599 = v668
		goto L68
	} else {
		goto L76
	}
L76:
	;
	goto L69
L77:
	;
	if v720&int32(1) == int32(0) {
		v783 = v576
		v794 = v531
		goto L33
	} else {
		goto L82
	}
L78:
	;
	v699 = int32(2)
	v701 = l0 + (v671+v299)<<(uint(v699)%32)
	v702 = *(*int32)(unsafe.Add(mBase, uint32(v701)))
	v708 = *(*int32)(unsafe.Add(mBase, uint32(v36+int32(504)+v671<<(uint(v699)%32))))
	v709 = v702 + v708
	v711 = v709 + int32(1)
	if v675 != 0 {
		goto L79
	} else {
		goto L80
	}
L79:
	;
	v712 = v711
	goto L81
L80:
	;
	v712 = v709
	goto L81
L81:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v701))) = v712
	v720 = base.B2i32(base.Ui32(v709) < base.Ui32(v702)) | v675&base.B2i32(v711 == int32(0))
	goto L77
L82:
	;
	v727 = v299 + v544
	v729 = v576 & int32(65535)
	if base.Ui32(v729) < base.Ui32(v727) {
		goto L83
	} else {
		goto L84
	}
L83:
	;
	v731 = v727
	goto L85
L84:
	;
	v731 = v729
	goto L85
L85:
	;
	v733 = v727
	goto L87
L86:
	;
	if base.Ui32(int32(124)) < base.Ui32(v729) {
		goto L8
	} else {
		goto L91
	}
L87:
	;
	if v733 == v731 {
		goto L86
	} else {
		goto L89
	}
L89:
	;
	v761 = l0 + v733<<(uint(int32(2))%32)
	v762 = *(*int32)(unsafe.Add(mBase, uint32(v761)))
	v763 = int32(1)
	v764 = v762 + v763
	*(*int32)(unsafe.Add(mBase, uint32(v761))) = v764
	if v764 == int32(0) {
		v733 = v733 + v763
		goto L87
	} else {
		goto L90
	}
L90:
	;
	v783 = v576
	v794 = v531
	goto L33
L91:
	;
	v775 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0+v729<<(uint(int32(2))%32)))) = v775
	v778 = v576 + v775
	*(*uint16)(unsafe.Add(mBase, uint32(l0)+500)) = uint16(v778)
	v783 = v778
	v794 = v531
	goto L33
L92:
	;
	goto L32
L93:
	;
	v843 = v812
	goto L94
L94:
	;
	v866 = int32(504)
	goto L98
L96:
	;
	v873 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v36)+1004)))
	v881 = *(*int32)(unsafe.Add(mBase, uint32(v873<<(uint(int32(2))%32)+(v36+int32(504))+int32(-4))))
	if v881 != 0 {
		v915 = v835
		goto L7
	} else {
		goto L99
	}
L97:
	;
	goto L96
L98:
	;
	v871 = F__emscripten_memcpy_bulkmem(m, v36+v866, l0, v866)
	mBase = m.M
	goto L97
L99:
	;
	v883 = v843 + int32(-1)
	*(*uint16)(unsafe.Add(mBase, uint32(l0)+500)) = uint16(v883)
	if v883&int32(65535) != 0 {
		v843 = v883
		goto L94
	} else {
		goto L100
	}
L100:
	;
	v915 = v835
	goto L7
}
func F_ffc_from_chars_double_options(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) {
	mBase := m.M
	_ = mBase
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v33 int64
	_ = v33
	var v40 int32
	_ = v40
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v108 int64
	_ = v108
	var v109 int64
	_ = v109
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v113 int32
	_ = v113
	var v123 int32
	_ = v123
	var v125 int32
	_ = v125
	var v127 int32
	_ = v127
	var v134 int32
	_ = v134
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v152 int32
	_ = v152
	var v163 int64
	_ = v163
	var v176 int32
	_ = v176
	var v178 int32
	_ = v178
	var v186 int64
	_ = v186
	var v188 int32
	_ = v188
	var v190 int32
	_ = v190
	var v191 int64
	_ = v191
	var v198 int32
	_ = v198
	var v208 int32
	_ = v208
	var v209 int64
	_ = v209
	var v222 int32
	_ = v222
	var v224 int32
	_ = v224
	var v232 int64
	_ = v232
	var v234 int32
	_ = v234
	var v239 int32
	_ = v239
	var v257 int32
	_ = v257
	var v258 int64
	_ = v258
	var v271 int64
	_ = v271
	var v275 int64
	_ = v275
	var v287 int64
	_ = v287
	var v290 int64
	_ = v290
	var v303 int64
	_ = v303
	var v305 int32
	_ = v305
	var v323 int32
	_ = v323
	var v324 int64
	_ = v324
	var v354 int32
	_ = v354
	var v355 int64
	_ = v355
	var v368 int32
	_ = v368
	var v370 int32
	_ = v370
	var v378 int64
	_ = v378
	var v380 int32
	_ = v380
	var v396 int32
	_ = v396
	var v397 int64
	_ = v397
	var v412 int64
	_ = v412
	var v415 int32
	_ = v415
	var v425 int64
	_ = v425
	var v432 int32
	_ = v432
	var v433 int64
	_ = v433
	var v434 int32
	_ = v434
	var v437 int64
	_ = v437
	var v438 int32
	_ = v438
	var v439 int32
	_ = v439
	var v458 int32
	_ = v458
	var v469 int32
	_ = v469
	var v471 int32
	_ = v471
	var v486 int32
	_ = v486
	var v492 int32
	_ = v492
	var v493 int32
	_ = v493
	var v495 int32
	_ = v495
	var v498 int32
	_ = v498
	var v504 int32
	_ = v504
	var v505 int32
	_ = v505
	var v508 int32
	_ = v508
	var v520 int64
	_ = v520
	var v530 int32
	_ = v530
	var v542 int32
	_ = v542
	var v544 int32
	_ = v544
	var v555 int64
	_ = v555
	var v557 int32
	_ = v557
	var v559 int64
	_ = v559
	var v560 int32
	_ = v560
	var v564 int64
	_ = v564
	var v619 int32
	_ = v619
	var v620 int32
	_ = v620
	var v623 int64
	_ = v623
	var v626 int64
	_ = v626
	var v631 int64
	_ = v631
	var v637 int32
	_ = v637
	var v643 float64
	_ = v643
	var v645 int32
	_ = v645
	var v649 int32
	_ = v649
	var v651 int32
	_ = v651
	var v655 int32
	_ = v655
	var v658 int32
	_ = v658
	var v690 int32
	_ = v690
	var v702 int32
	_ = v702
	var v705 int32
	_ = v705
	var v707 int32
	_ = v707
	var v708 int32
	_ = v708
	var v712 int32
	_ = v712
	var v742 int32
	_ = v742
	var v745 int64
	_ = v745
	var v748 int64
	_ = v748
	var v757 int32
	_ = v757
	var v764 float64
	_ = v764
	var v779 int64
	_ = v779
	var v780 int64
	_ = v780
	var v787 int32
	_ = v787
	var v807 int32
	_ = v807
	var v825 int64
	_ = v825
	var v834 int32
	_ = v834
	var v836 int32
	_ = v836
	var v839 int64
	_ = v839
	var v841 int32
	_ = v841
	var v844 int64
	_ = v844
	var v876 int64
	_ = v876
	var v880 int32
	_ = v880
	var v893 int64
	_ = v893
	var v907 int32
	_ = v907
	var v910 int64
	_ = v910
	var v913 int64
	_ = v913
	var v935 int64
	_ = v935
	var v951 int32
	_ = v951
	var v964 int64
	_ = v964
	var v978 int32
	_ = v978
	var v981 int64
	_ = v981
	var v984 int64
	_ = v984
	var v989 int32
	_ = v989
	var v1003 int64
	_ = v1003
	var v1017 int32
	_ = v1017
	var v1020 int32
	_ = v1020
	var v1031 int64
	_ = v1031
	var v1050 int32
	_ = v1050
	var v1055 int64
	_ = v1055
	var v1063 int64
	_ = v1063
	var v1076 int64
	_ = v1076
	var v1083 int32
	_ = v1083
	var v1086 int32
	_ = v1086
	var v1100 float32
	_ = v1100
	var v1101 float32
	_ = v1101
	var v1108 int32
	_ = v1108
	var v1109 float64
	_ = v1109
	var v1116 float64
	_ = v1116
	var v1122 float64
	_ = v1122
	var v1124 float64
	_ = v1124
	var v1128 float64
	_ = v1128
	var v1131 int32
	_ = v1131
	var v1136 int64
	_ = v1136
	var v1144 float64
	_ = v1144
	var v1149 float64
	_ = v1149
	var v1151 float64
	_ = v1151
	var v1155 float64
	_ = v1155
	var v1157 int32
	_ = v1157
	var v1168 int32
	_ = v1168
	var v1169 int32
	_ = v1169
	var v1174 int64
	_ = v1174
	var v1175 int64
	_ = v1175
	var v1176 int64
	_ = v1176
	var v1177 int64
	_ = v1177
	var v1183 int64
	_ = v1183
	var v1184 int64
	_ = v1184
	var v1186 int64
	_ = v1186
	var v1189 int64
	_ = v1189
	var v1190 int64
	_ = v1190
	var v1192 int64
	_ = v1192
	var v1193 int64
	_ = v1193
	var v1197 int64
	_ = v1197
	var v1204 int64
	_ = v1204
	var v1216 int64
	_ = v1216
	var v1221 int64
	_ = v1221
	var v1222 int64
	_ = v1222
	var v1227 int32
	_ = v1227
	var v1234 int64
	_ = v1234
	var v1235 int64
	_ = v1235
	var v1241 int64
	_ = v1241
	var v1242 int64
	_ = v1242
	var v1244 int64
	_ = v1244
	var v1247 int64
	_ = v1247
	var v1248 int64
	_ = v1248
	var v1250 int64
	_ = v1250
	var v1251 int64
	_ = v1251
	var v1255 int64
	_ = v1255
	var v1262 int64
	_ = v1262
	var v1277 int64
	_ = v1277
	var v1278 int64
	_ = v1278
	var v1283 int64
	_ = v1283
	var v1284 int64
	_ = v1284
	var v1286 int64
	_ = v1286
	var v1288 int64
	_ = v1288
	var v1289 int64
	_ = v1289
	var v1296 int32
	_ = v1296
	var v1298 int32
	_ = v1298
	var v1308 int64
	_ = v1308
	var v1309 int64
	_ = v1309
	var v1311 int64
	_ = v1311
	var v1320 int64
	_ = v1320
	var v1325 int64
	_ = v1325
	var v1328 int64
	_ = v1328
	var v1333 int64
	_ = v1333
	var v1336 int64
	_ = v1336
	var v1342 int32
	_ = v1342
	var v1351 int32
	_ = v1351
	var v1352 int64
	_ = v1352
	var v1353 int32
	_ = v1353
	var v1355 int32
	_ = v1355
	var v1356 int32
	_ = v1356
	var v1358 int64
	_ = v1358
	var v1359 int32
	_ = v1359
	var v1361 int64
	_ = v1361
	var v1369 int32
	_ = v1369
	var v1370 int64
	_ = v1370
	var v1372 int64
	_ = v1372
	var v1381 int32
	_ = v1381
	var v1382 int32
	_ = v1382
	var v1387 int64
	_ = v1387
	var v1388 int64
	_ = v1388
	var v1389 int64
	_ = v1389
	var v1390 int64
	_ = v1390
	var v1396 int64
	_ = v1396
	var v1397 int64
	_ = v1397
	var v1399 int64
	_ = v1399
	var v1402 int64
	_ = v1402
	var v1403 int64
	_ = v1403
	var v1405 int64
	_ = v1405
	var v1406 int64
	_ = v1406
	var v1410 int64
	_ = v1410
	var v1417 int64
	_ = v1417
	var v1429 int64
	_ = v1429
	var v1434 int64
	_ = v1434
	var v1435 int64
	_ = v1435
	var v1440 int32
	_ = v1440
	var v1447 int64
	_ = v1447
	var v1448 int64
	_ = v1448
	var v1454 int64
	_ = v1454
	var v1455 int64
	_ = v1455
	var v1457 int64
	_ = v1457
	var v1460 int64
	_ = v1460
	var v1461 int64
	_ = v1461
	var v1463 int64
	_ = v1463
	var v1464 int64
	_ = v1464
	var v1468 int64
	_ = v1468
	var v1475 int64
	_ = v1475
	var v1490 int64
	_ = v1490
	var v1491 int64
	_ = v1491
	var v1496 int64
	_ = v1496
	var v1497 int64
	_ = v1497
	var v1499 int64
	_ = v1499
	var v1501 int64
	_ = v1501
	var v1502 int64
	_ = v1502
	var v1511 int32
	_ = v1511
	var v1520 int64
	_ = v1520
	var v1521 int64
	_ = v1521
	var v1523 int64
	_ = v1523
	var v1534 int64
	_ = v1534
	var v1539 int64
	_ = v1539
	var v1542 int64
	_ = v1542
	var v1547 int64
	_ = v1547
	var v1549 int64
	_ = v1549
	var v1552 int32
	_ = v1552
	var v1553 int32
	_ = v1553
	var v1554 int32
	_ = v1554
	var v1555 int32
	_ = v1555
	var v1557 int32
	_ = v1557
	var v1558 int32
	_ = v1558
	var v1567 int64
	_ = v1567
	var v1569 int64
	_ = v1569
	var v1570 int32
	_ = v1570
	var v1573 int64
	_ = v1573
	var v1581 int32
	_ = v1581
	var v1582 int32
	_ = v1582
	var v1587 int64
	_ = v1587
	var v1588 int64
	_ = v1588
	var v1589 int64
	_ = v1589
	var v1590 int64
	_ = v1590
	var v1596 int64
	_ = v1596
	var v1597 int64
	_ = v1597
	var v1599 int64
	_ = v1599
	var v1602 int64
	_ = v1602
	var v1603 int64
	_ = v1603
	var v1605 int64
	_ = v1605
	var v1606 int64
	_ = v1606
	var v1610 int64
	_ = v1610
	var v1617 int64
	_ = v1617
	var v1633 int64
	_ = v1633
	var v1634 int64
	_ = v1634
	var v1638 int64
	_ = v1638
	var v1640 int32
	_ = v1640
	var v1647 int64
	_ = v1647
	var v1648 int64
	_ = v1648
	var v1654 int64
	_ = v1654
	var v1655 int64
	_ = v1655
	var v1657 int64
	_ = v1657
	var v1660 int64
	_ = v1660
	var v1661 int64
	_ = v1661
	var v1663 int64
	_ = v1663
	var v1664 int64
	_ = v1664
	var v1668 int64
	_ = v1668
	var v1675 int64
	_ = v1675
	var v1690 int64
	_ = v1690
	var v1696 int64
	_ = v1696
	var v1702 int32
	_ = v1702
	var v1704 int64
	_ = v1704
	var v1713 int32
	_ = v1713
	var v1725 int64
	_ = v1725
	var v1735 int64
	_ = v1735
	var v1745 int64
	_ = v1745
	var v1755 int64
	_ = v1755
	var v1765 int64
	_ = v1765
	var v1772 int32
	_ = v1772
	var v1773 int64
	_ = v1773
	var v1775 int64
	_ = v1775
	var v1777 int64
	_ = v1777
	var v1789 int64
	_ = v1789
	var v1791 int64
	_ = v1791
	var v1794 int32
	_ = v1794
	var v1797 int64
	_ = v1797
	var v1815 int64
	_ = v1815
	var v1880 float64
	_ = v1880
	v29 = m.G0
	v31 = v29 - int32(240)
	m.G0 = v31
	v33 = *(*int64)(unsafe.Add(mBase, uint32(l4)))
	if v33&int64(256) == int64(0) {
		v77 = l1
		goto L1
	} else {
		goto L2
	}
L1:
	;
	if v77 != l2 {
		goto L10
	} else {
		goto L11
	}
L2:
	;
	if l1 == l2 {
		v77 = l1
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v40 = l1
	goto L4
L4:
	;
	v67 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v40))))
	v70 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v67)+uint32(_consts[994]))))
	if v70 != int32(1) {
		v77 = v40
		goto L1
	} else {
		goto L6
	}
L5:
	;
	v77 = l2
	goto L1
L6:
	;
	v74 = v40 + int32(1)
	if v74 != l2 {
		v40 = v74
		goto L4
	} else {
		goto L7
	}
L7:
	;
	goto L5
L8:
	;
	*(*float64)(unsafe.Add(mBase, uint32(l3))) = v1880
	m.G0 = v31 + int32(240)
	return
L9:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v77
	v1880 = float64(0)
	goto L8
L10:
	;
	v108 = v33 & int64(32)
	v109 = *(*int64)(unsafe.Add(mBase, uint32(l4)+8))
	v110 = base.I32_wrap_i64(v109)
	v111 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v77))))
	v113 = base.B2i32(v111 == int32(45))
	if v111 == int32(45) {
		goto L16
	} else {
		goto L17
	}
L11:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = int32(1)
	goto L9
L12:
	;
	if int64(20) <= v437 {
		goto L133
	} else {
		goto L134
	}
L13:
	;
	v779 = int64(0)
	v780 = v425
	v787 = v432
	goto L12
L14:
	;
	if v33&int64(16) == int64(0) {
		goto L96
	} else {
		goto L97
	}
L15:
	;
	v142 = base.B2i32(v141 == l2)
	if v142 == int32(0) {
		goto L27
	} else {
		goto L28
	}
L16:
	;
	v123 = v77 + int32(1)
	if v123 == l2 {
		goto L14
	} else {
		goto L20
	}
L17:
	;
	if v33&int64(160) != int64(128) {
		v140 = v111
		v141 = v77
		goto L15
	} else {
		goto L18
	}
L18:
	;
	if v111 != int32(43) {
		v140 = v111
		v141 = v77
		goto L15
	} else {
		goto L19
	}
L19:
	;
	goto L16
L20:
	;
	v125 = int32(*(*int8)(unsafe.Add(mBase, uint32(v123))))
	v127 = v125 + int32(-48)
	if v108 == int64(0) {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	if base.Ui32(v127) < base.Ui32(int32(10)) {
		v140 = v125
		v141 = v123
		goto L15
	} else {
		goto L24
	}
L22:
	;
	if base.Ui32(v127) < base.Ui32(int32(10)) {
		v140 = v125
		v141 = v123
		goto L15
	} else {
		goto L23
	}
L23:
	;
	goto L14
L24:
	;
	v134 = int32(255)
	if v125&v134 != v110&v134 {
		goto L14
	} else {
		goto L25
	}
L25:
	;
	v140 = v125
	v141 = v123
	goto L15
L26:
	;
	v222 = v198 - v141
	v224 = base.B2i32(v108 == int64(0))
	if v108 == int64(0) {
		goto L34
	} else {
		goto L35
	}
L27:
	;
	v152 = v141
	v163 = int64(0)
	goto L30
L28:
	;
	v198 = v141
	v208 = int32(1)
	v209 = int64(0)
	goto L26
L29:
	;
	v198 = v190
	v208 = base.B2i32(base.Ui32(v178) < base.Ui32(int32(10)))
	v209 = v191
	goto L26
L30:
	;
	v176 = int32(*(*int8)(unsafe.Add(mBase, uint32(v152))))
	v178 = v176 + int32(-48)
	if base.Ui32(int32(9)) < base.Ui32(v178) {
		v190 = v152
		v191 = v163
		goto L29
	} else {
		goto L32
	}
L31:
	;
	v190 = l2
	v191 = v186
	goto L29
L32:
	;
	v186 = v163*int64(10) + base.I64_extend_i32_s(v176) + int64(-48)
	v188 = v152 + int32(1)
	if v188 != l2 {
		v152 = v188
		v163 = v186
		goto L30
	} else {
		goto L33
	}
L33:
	;
	goto L31
L34:
	;
	v232 = base.I64_extend_i32_s(v222)
	if v208 != 0 {
		goto L40
	} else {
		goto L41
	}
L35:
	;
	if v198 == v141 {
		goto L14
	} else {
		goto L36
	}
L36:
	;
	if v222 < int32(2) {
		goto L34
	} else {
		goto L37
	}
L37:
	;
	if v140&int32(255) == int32(48) {
		goto L14
	} else {
		goto L38
	}
L38:
	;
	goto L34
L39:
	;
	if v108 == int64(0) {
		goto L57
	} else {
		goto L58
	}
L40:
	;
	v415 = int32(0)
	v425 = int64(0)
	v432 = v198
	v433 = v209
	v434 = v415
	v437 = v232
	v438 = v415
	v439 = v415
	goto L39
L41:
	;
	v234 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v198))))
	if v234 != v110&int32(255) {
		goto L40
	} else {
		goto L42
	}
L42:
	;
	v239 = v198 + int32(1)
	if int32(8) <= l2-v239 {
		goto L44
	} else {
		goto L45
	}
L43:
	;
	if v323 == l2 {
		v396 = v323
		v397 = v324
		goto L50
	} else {
		goto L51
	}
L44:
	;
	v257 = v239
	v258 = v209
	goto L46
L45:
	;
	v323 = v239
	v324 = v209
	goto L43
L46:
	;
	v271 = *(*int64)(unsafe.Add(mBase, uint32(v257)))
	v275 = v271 + int64(-3472328296227680304)
	if base.B2i32((v271+int64(5063812098665367110)|v275)&int64(-9187201950435737472) == int64(0)) == int32(0) {
		v323 = v257
		v324 = v258
		goto L43
	} else {
		goto L48
	}
L47:
	;
	v323 = v305
	v324 = v303
	goto L43
L48:
	;
	v287 = v275*int64(10) + int64(base.Ui64(v275)>>(uint(int64(8))%64))
	v290 = int64(1095216660735)
	v303 = int64(base.Ui64(int64(base.Ui64(v287)>>(uint(int64(16))%64))&v290*int64(42949672960001)+v287&v290*int64(4294967296000100))>>(uint(int64(32))%64)) + v258*int64(100000000)
	v305 = v257 + int32(8)
	if int32(7) < l2-v305 {
		v257 = v305
		v258 = v303
		goto L46
	} else {
		goto L49
	}
L49:
	;
	goto L47
L50:
	;
	v412 = base.I64_extend_i32_s(v239 - v396)
	v425 = v412
	v432 = v396
	v433 = v397
	v434 = int32(1)
	v437 = v232 - v412
	v438 = v239
	v439 = v396 - v239
	goto L39
L51:
	;
	v354 = v323
	v355 = v324
	goto L52
L52:
	;
	v368 = int32(*(*int8)(unsafe.Add(mBase, uint32(v354))))
	v370 = v368 + int32(-48)
	if base.Ui32(int32(9)) < base.Ui32(v370) {
		v396 = v354
		v397 = v355
		goto L50
	} else {
		goto L54
	}
L53:
	;
	v396 = v323 + (l2 - v323)
	v397 = v378
	goto L50
L54:
	;
	v378 = v355*int64(10) + base.I64_extend_i32_u(v370)&int64(255)
	v380 = v354 + int32(1)
	if v380 != l2 {
		v354 = v380
		v355 = v378
		goto L52
	} else {
		goto L55
	}
L55:
	;
	goto L53
L56:
	;
	if v33&int64(1) == int64(0) {
		goto L62
	} else {
		goto L63
	}
L57:
	;
	if v437 == int64(0) {
		goto L14
	} else {
		goto L60
	}
L58:
	;
	if v434&base.B2i32(v425 == int64(0)) == int32(0) {
		goto L56
	} else {
		goto L59
	}
L59:
	;
	goto L14
L60:
	;
	goto L56
L61:
	;
	switch v486 + int32(-68) {
	case 0, 1, 32, 33:
		goto L75
	default:
		v492 = v432
		goto L74
	}
L62:
	;
	if v33&int64(64) == int64(0) {
		goto L66
	} else {
		goto L67
	}
L63:
	;
	if v432 == l2 {
		goto L62
	} else {
		goto L64
	}
L64:
	;
	v458 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v432))))
	if v458|int32(32) == int32(101) {
		v486 = v458
		goto L61
	} else {
		goto L65
	}
L65:
	;
	goto L62
L66:
	;
	if v33&int64(5) == int64(1) {
		goto L14
	} else {
		goto L73
	}
L67:
	;
	if v432 == l2 {
		goto L66
	} else {
		goto L68
	}
L68:
	;
	v469 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v432))))
	v471 = v469 + int32(-43)
	if base.Ui32(int32(25)) < base.Ui32(v471) {
		goto L69
	} else {
		goto L70
	}
L69:
	;
	if v469 == int32(100) {
		v486 = v469
		goto L61
	} else {
		goto L72
	}
L70:
	;
	if int32(1)<<(uint(v471)%32)&int32(33554437) != 0 {
		v486 = v469
		goto L61
	} else {
		goto L71
	}
L71:
	;
	goto L69
L72:
	;
	goto L66
L73:
	;
	goto L13
L74:
	;
	v493 = int32(0)
	if v492 == l2 {
		v504 = v493
		v505 = v492
		goto L76
	} else {
		goto L77
	}
L75:
	;
	v492 = v432 + int32(1)
	goto L74
L76:
	;
	if v505 == l2 {
		goto L80
	} else {
		goto L81
	}
L77:
	;
	v495 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v492))))
	if v495 != int32(45) {
		goto L78
	} else {
		goto L79
	}
L78:
	;
	v504 = v493
	v505 = v492 + base.B2i32(v495 == int32(43))
	goto L76
L79:
	;
	v498 = int32(1)
	v504 = v498
	v505 = v492 + v498
	goto L76
L80:
	;
	if base.B2i32(v33&int64(4) == int64(0)) == int32(0) {
		goto L13
	} else {
		goto L95
	}
L81:
	;
	v508 = int32(*(*int8)(unsafe.Add(mBase, uint32(v505))))
	if base.Ui32(int32(9)) < base.Ui32(v508+int32(-48)) {
		goto L80
	} else {
		goto L82
	}
L82:
	;
	v520 = int64(0)
	v530 = v505
	goto L84
L83:
	;
	if v504 != 0 {
		goto L92
	} else {
		goto L93
	}
L84:
	;
	v542 = int32(*(*int8)(unsafe.Add(mBase, uint32(v530))))
	v544 = v542 + int32(-48)
	if base.Ui32(v544) <= base.Ui32(int32(9)) {
		goto L86
	} else {
		goto L87
	}
L85:
	;
	v559 = v555
	v560 = l2
	goto L83
L86:
	;
	if v520 < int64(268435456) {
		goto L88
	} else {
		goto L89
	}
L87:
	;
	v559 = v520
	v560 = v530
	goto L83
L88:
	;
	v555 = v520*int64(10) + base.I64_extend_i32_u(v544)&int64(255)
	goto L90
L89:
	;
	v555 = v520
	goto L90
L90:
	;
	v557 = v530 + int32(1)
	if v557 != l2 {
		v520 = v555
		v530 = v557
		goto L84
	} else {
		goto L91
	}
L91:
	;
	goto L85
L92:
	;
	v564 = int64(0) - v559
	goto L94
L93:
	;
	v564 = v559
	goto L94
L94:
	;
	v779 = v564
	v780 = v564 + v425
	v787 = v560
	goto L12
L95:
	;
	goto L14
L96:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v77
	if v111 == int32(45) {
		goto L99
	} else {
		goto L100
	}
L97:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = int32(1)
	goto L9
L98:
	;
	v620 = l2 - v619
	if v620 < int32(3) {
		goto L103
	} else {
		goto L104
	}
L99:
	;
	v619 = v77 + int32(1)
	goto L98
L100:
	;
	if v33&int64(128) == int64(0) {
		v619 = v77
		goto L98
	} else {
		goto L101
	}
L101:
	;
	if v111 != int32(43) {
		v619 = v77
		goto L98
	} else {
		goto L102
	}
L102:
	;
	goto L99
L103:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = int32(1)
	v1880 = float64(0)
	goto L8
L104:
	;
	v623 = int64(*(*uint16)(unsafe.Add(mBase, uint32(v619))))
	v626 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v619+int32(2)))))
	v631 = v623 | v626<<(uint(int64(16))%64) | int64(2314885530818453536)
	if v631 == int64(2314885530823061097) {
		goto L105
	} else {
		goto L106
	}
L105:
	;
	v742 = int32(8)
	if base.Ui32(v620) < base.Ui32(v742) {
		goto L126
	} else {
		goto L127
	}
L106:
	;
	if v631 != int64(2314885530823582062) {
		goto L103
	} else {
		goto L107
	}
L107:
	;
	v637 = v619 + int32(3)
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v637
	if v111 == int32(45) {
		goto L108
	} else {
		goto L109
	}
L108:
	;
	v643 = math.Float64frombits(uint64(0xfff8000000000000))
	goto L110
L109:
	;
	v643 = math.Float64frombits(uint64(0x7ff8000000000000))
	goto L110
L110:
	;
	if v637 == l2 {
		v1880 = v643
		goto L8
	} else {
		goto L111
	}
L111:
	;
	v645 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v619)+3)))
	if v645 != int32(40) {
		v1880 = v643
		goto L8
	} else {
		goto L112
	}
L112:
	;
	v649 = v619 + int32(4)
	if v649 == l2 {
		v1880 = v643
		goto L8
	} else {
		goto L113
	}
L113:
	;
	v651 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v649))))
	if v651 == int32(41) {
		v712 = v649
		goto L114
	} else {
		goto L115
	}
L114:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v712 + int32(1)
	v1880 = v643
	goto L8
L115:
	;
	v655 = v649
	v658 = v651
	goto L116
L116:
	;
	if base.Ui32((v658&int32(223)+int32(-65))&int32(255)) < base.Ui32(int32(26)) {
		goto L119
	} else {
		goto L120
	}
L117:
	;
	v712 = v707
	goto L114
L118:
	;
	v708 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v707))))
	if v708 != int32(41) {
		v655 = v707
		v658 = v708
		goto L116
	} else {
		goto L124
	}
L119:
	;
	v705 = v655 + int32(1)
	if v705 == l2 {
		v1880 = v643
		goto L8
	} else {
		goto L123
	}
L120:
	;
	v690 = int32(255)
	if base.B2i32(v658&v690 != int32(95))&base.B2i32(base.Ui32((v658+int32(-58))&v690) < base.Ui32(int32(246))) != 0 {
		v1880 = v643
		goto L8
	} else {
		goto L121
	}
L121:
	;
	v702 = v655 + int32(1)
	if v702 != l2 {
		v707 = v702
		goto L118
	} else {
		goto L122
	}
L122:
	;
	v1880 = v643
	goto L8
L123:
	;
	v707 = v705
	goto L118
L124:
	;
	goto L117
L125:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v619 + v757
	if v111 == int32(45) {
		goto L129
	} else {
		goto L130
	}
L126:
	;
	v757 = int32(3)
	goto L125
L127:
	;
	v745 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v619)+3)))
	v748 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v619+int32(7)))))
	if (v745|v748<<(uint(int64(32))%64))&int64(961533698015) == int64(383666179657) {
		v757 = v742
		goto L125
	} else {
		goto L128
	}
L128:
	;
	goto L126
L129:
	;
	v764 = math.Float64frombits(uint64(0xfff0000000000000))
	goto L131
L130:
	;
	v764 = math.Float64frombits(uint64(0x7ff0000000000000))
	goto L131
L131:
	;
	v1880 = v764
	goto L8
L132:
	;
	v1076 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v31)+200)) = v1076
	*(*int32)(unsafe.Add(mBase, uint32(v31)+196)) = v439
	*(*int32)(unsafe.Add(mBase, uint32(v31)+192)) = v438
	*(*int32)(unsafe.Add(mBase, uint32(v31)+188)) = v222
	*(*int32)(unsafe.Add(mBase, uint32(v31)+184)) = v141
	v1083 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v31)+183)) = uint8(v1083)
	*(*uint8)(unsafe.Add(mBase, uint32(v31)+182)) = uint8(v1050)
	v1086 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v31)+181)) = uint8(v1086)
	*(*uint8)(unsafe.Add(mBase, uint32(v31)+180)) = uint8(v113)
	*(*int32)(unsafe.Add(mBase, uint32(v31)+176)) = v787
	*(*int64)(unsafe.Add(mBase, uint32(v31)+168)) = v1063
	*(*int64)(unsafe.Add(mBase, uint32(v31)+160)) = v1055
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v787
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1083
	if v1050 != 0 {
		goto L163
	} else {
		goto L164
	}
L133:
	;
	if v141 == l2 {
		goto L135
	} else {
		goto L136
	}
L134:
	;
	v1050 = int32(0)
	v1055 = v780
	v1063 = v433
	goto L132
L135:
	;
	v876 = int64(0)
	if v141 == v198 {
		v935 = v876
		goto L146
	} else {
		goto L147
	}
L136:
	;
	v807 = v141
	v825 = v437
	goto L138
L137:
	;
	if v844 < int64(20) {
		v1050 = int32(0)
		v1055 = v780
		v1063 = v433
		goto L132
	} else {
		goto L144
	}
L138:
	;
	v834 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v807))))
	v836 = base.B2i32(v834 == int32(48))
	if v834 == int32(48) {
		goto L140
	} else {
		goto L141
	}
L139:
	;
	v844 = v839
	goto L137
L140:
	;
	v839 = v825 - base.I64_extend_i32_u(v836)
	v841 = v807 + int32(1)
	if v841 != l2 {
		v807 = v841
		v825 = v839
		goto L138
	} else {
		goto L143
	}
L141:
	;
	if v834 != v110&int32(255) {
		v844 = v825
		goto L137
	} else {
		goto L142
	}
L142:
	;
	goto L140
L143:
	;
	goto L139
L144:
	;
	goto L135
L145:
	;
	v1050 = int32(1)
	v1055 = v779 + base.I64_extend_i32_s(v1020-v1017)
	v1063 = v1031
	goto L132
L146:
	;
	if v439 != 0 {
		goto L155
	} else {
		goto L156
	}
L147:
	;
	v880 = v141
	v893 = v876
	goto L149
L148:
	;
	if base.Ui64(int64(999999999999999999)) < base.Ui64(v913) {
		v1017 = v907
		v1020 = v198
		v1031 = v913
		goto L145
	} else {
		goto L153
	}
L149:
	;
	v907 = v880 + int32(1)
	v910 = int64(*(*int8)(unsafe.Add(mBase, uint32(v880))))
	v913 = v893*int64(10) + v910 + int64(-48)
	if base.Ui64(int64(999999999999999999)) < base.Ui64(v913) {
		goto L148
	} else {
		goto L151
	}
L150:
	;
	goto L148
L151:
	;
	if v907 != v198 {
		v880 = v907
		v893 = v913
		goto L149
	} else {
		goto L152
	}
L152:
	;
	goto L150
L153:
	;
	v935 = v913
	goto L146
L154:
	;
	v1017 = v989
	v1020 = v438
	v1031 = v1003
	goto L145
L155:
	;
	v951 = v438
	v964 = v935
	goto L157
L156:
	;
	v989 = v438
	v1003 = v935
	goto L154
L157:
	;
	v978 = v951 + int32(1)
	v981 = int64(*(*int8)(unsafe.Add(mBase, uint32(v951))))
	v984 = v964*int64(10) + v981 + int64(-48)
	if base.Ui64(int64(999999999999999999)) < base.Ui64(v984) {
		v989 = v978
		v1003 = v984
		goto L154
	} else {
		goto L159
	}
L158:
	;
	v989 = v978
	v1003 = v984
	goto L154
L159:
	;
	if v978 != v438+v439 {
		v951 = v978
		v964 = v984
		goto L157
	} else {
		goto L160
	}
L160:
	;
	goto L158
L161:
	;
	if v1050 == int32(0) {
		v1794 = v1359
		v1797 = v1361
		goto L217
	} else {
		goto L218
	}
L162:
	;
	v1168 = v31 + int32(144)
	v1169 = base.I32_wrap_i64(v1055)
	v1174 = *(*int64)(unsafe.Add(mBase, uint32(v1169<<(uint(int32(4))%32))+uint32(_consts[995])))
	v1175 = int64(0)
	v1176 = base.I64_clz(v1063)
	v1177 = v1063 << (uint(v1176) % 64)
	v1183 = int64(32)
	v1184 = int64(base.Ui64(v1177) >> (uint(v1183) % 64))
	v1186 = int64(base.Ui64(v1174) >> (uint(v1183) % 64))
	v1189 = int64(4294967295)
	v1190 = v1177 & v1189
	v1192 = v1174 & v1189
	v1193 = v1190 * v1192
	v1197 = int64(base.Ui64(v1193)>>(uint(v1183)%64)) + v1190*v1186
	v1204 = v1197&v1189 + v1184*v1192
	*(*int64)(unsafe.Add(mBase, uint32(v1168)+8)) = v1175*v1174 + v1175*v1177 + v1184*v1186 + int64(base.Ui64(v1197)>>(uint(v1183)%64)) + int64(base.Ui64(v1204)>>(uint(v1183)%64))
	*(*int64)(unsafe.Add(mBase, uint32(v1168))) = v1204<<(uint(v1183)%64) | v1193&v1189
	goto L188
L163:
	;
	v1157 = int32(0)
	if v1063 == int64(0) {
		v1359 = v1157
		v1361 = v1076
		goto L161
	} else {
		goto L185
	}
L164:
	;
	if base.Ui64(v1055+int64(-23)) < base.Ui64(int64(-45)) {
		goto L163
	} else {
		goto L165
	}
L165:
	;
	v1100 = *(*float32)(unsafe.Add(mBase, _consts[996]))
	v1101 = float32(1)
	if base.F32_ne(base.F32_add(v1100, v1101), base.F32_sub(v1101, v1100)) != 0 {
		goto L166
	} else {
		goto L167
	}
L166:
	;
	if v1055 < int64(0) {
		goto L163
	} else {
		goto L175
	}
L167:
	;
	if base.Ui64(int64(9007199254740992)) < base.Ui64(v1063) {
		goto L162
	} else {
		goto L168
	}
L168:
	;
	v1108 = base.I32_wrap_i64(v1055)
	v1109 = base.F64_convert_i64_u(v1063)
	if int64(-1) < v1055 {
		goto L170
	} else {
		goto L171
	}
L169:
	;
	if v111 == int32(45) {
		goto L172
	} else {
		goto L173
	}
L170:
	;
	v1122 = *(*float64)(unsafe.Add(mBase, uint32(v1108<<(uint(int32(3))%32))+uint32(_consts[997])))
	v1124 = base.F64_mul(v1122, v1109)
	goto L169
L171:
	;
	v1116 = *(*float64)(unsafe.Add(mBase, uint32(int32(_a1859)-v1108<<(uint(int32(3))%32))))
	v1124 = base.F64_div(v1109, v1116)
	goto L169
L172:
	;
	v1128 = base.F64_neg(v1124)
	goto L174
L173:
	;
	v1128 = v1124
	goto L174
L174:
	;
	v1880 = v1128
	goto L8
L175:
	;
	v1131 = base.I32_wrap_i64(v1055)
	v1136 = *(*int64)(unsafe.Add(mBase, uint32(v1131<<(uint(int32(3))%32))+uint32(_consts[998])))
	if base.Ui64(v1136) < base.Ui64(v1063) {
		goto L162
	} else {
		goto L176
	}
L176:
	;
	if v1063 != int64(0) {
		goto L177
	} else {
		goto L178
	}
L177:
	;
	v1149 = *(*float64)(unsafe.Add(mBase, uint32(v1131<<(uint(int32(3))%32))+uint32(_consts[997])))
	v1151 = base.F64_mul(v1149, base.F64_convert_i64_u(v1063))
	if v111 == int32(45) {
		goto L182
	} else {
		goto L183
	}
L178:
	;
	if v111 == int32(45) {
		goto L179
	} else {
		goto L180
	}
L179:
	;
	v1144 = math.Float64frombits(uint64(0x8000000000000000))
	goto L181
L180:
	;
	v1144 = float64(0)
	goto L181
L181:
	;
	v1880 = v1144
	goto L8
L182:
	;
	v1155 = base.F64_neg(v1151)
	goto L184
L183:
	;
	v1155 = v1151
	goto L184
L184:
	;
	v1880 = v1155
	goto L8
L185:
	;
	if v1055 < int64(-342) {
		v1359 = v1157
		v1361 = v1076
		goto L161
	} else {
		goto L186
	}
L186:
	;
	if v1055 <= int64(308) {
		goto L162
	} else {
		goto L187
	}
L187:
	;
	v1359 = int32(2047)
	v1361 = v1076
	goto L161
L188:
	;
	v1216 = *(*int64)(unsafe.Add(mBase, uint32(v31)+144))
	v1221 = *(*int64)(unsafe.Add(mBase, uint32(v31+int32(152))))
	v1222 = int64(511)
	if v1221&v1222 != v1222 {
		v1283 = v1221
		v1284 = v1216
		goto L189
	} else {
		goto L190
	}
L189:
	;
	v1286 = int64(base.Ui64(v1283) >> (uint(int64(63)) % 64))
	v1288 = v1286 + int64(9)
	v1289 = int64(base.Ui64(v1283) >> (uint(v1288) % 64))
	v1296 = v1169*int32(217706)>>(uint(int32(16))%32) - base.I32_wrap_i64(v1176) + base.I32_wrap_i64(v1286)
	v1298 = v1296 + int32(63)
	if int32(-1023) < v1298 {
		goto L192
	} else {
		goto L193
	}
L190:
	;
	v1227 = v31 + int32(128)
	v1234 = *(*int64)(unsafe.Add(mBase, uint32(v1169<<(uint(int32(1))%32)<<(uint(int32(3))%32))+uint32(_consts[999])))
	v1235 = int64(0)
	v1241 = int64(32)
	v1242 = int64(base.Ui64(v1177) >> (uint(v1241) % 64))
	v1244 = int64(base.Ui64(v1234) >> (uint(v1241) % 64))
	v1247 = int64(4294967295)
	v1248 = v1177 & v1247
	v1250 = v1234 & v1247
	v1251 = v1248 * v1250
	v1255 = int64(base.Ui64(v1251)>>(uint(v1241)%64)) + v1248*v1244
	v1262 = v1255&v1247 + v1242*v1250
	*(*int64)(unsafe.Add(mBase, uint32(v1227)+8)) = v1235*v1234 + v1235*v1177 + v1242*v1244 + int64(base.Ui64(v1255)>>(uint(v1241)%64)) + int64(base.Ui64(v1262)>>(uint(v1241)%64))
	*(*int64)(unsafe.Add(mBase, uint32(v1227))) = v1262<<(uint(v1241)%64) | v1251&v1247
	goto L191
L191:
	;
	v1277 = *(*int64)(unsafe.Add(mBase, uint32(v31+int32(136))))
	v1278 = v1277 + v1216
	v1283 = base.I64_extend_i32_u(base.B2i32(base.Ui64(v1278) < base.Ui64(v1277))) + v1221
	v1284 = v1278
	goto L189
L192:
	;
	if v1289<<(uint(v1288)%64) == v1283 {
		goto L198
	} else {
		goto L199
	}
L193:
	;
	if base.Ui32(int32(-1085)) <= base.Ui32(v1298) {
		goto L194
	} else {
		goto L195
	}
L194:
	;
	v1308 = int64(base.Ui64(v1289) >> (uint(base.I64_extend_i32_u(int32(-1022)-v1298)) % 64))
	v1309 = int64(1)
	v1311 = v1308&v1309 + v1308
	v1359 = base.B2i32(base.Ui64(int64(9007199254740991)) < base.Ui64(v1311))
	v1361 = int64(base.Ui64(v1311) >> (uint(v1309) % 64))
	goto L161
L195:
	;
	v1359 = int32(0)
	v1361 = int64(0)
	goto L161
L196:
	;
	v1353 = int32(2047)
	v1355 = base.B2i32(base.Ui32(v1351) < base.Ui32(v1353))
	if base.Ui32(v1351) < base.Ui32(v1353) {
		goto L211
	} else {
		goto L212
	}
L197:
	;
	v1351 = v1296 + int32(1086)
	v1352 = int64(base.Ui64(v1336)>>(uint(int64(1))%64)) & int64(4503599627370495)
	goto L196
L198:
	;
	v1320 = v1289 & int64(72057594037927932)
	goto L200
L199:
	;
	v1320 = v1289
	goto L200
L200:
	;
	if v1289&int64(3) == int64(1) {
		goto L201
	} else {
		goto L202
	}
L201:
	;
	v1325 = v1320
	goto L203
L202:
	;
	v1325 = v1289
	goto L203
L203:
	;
	if base.Ui64(v1284) < base.Ui64(int64(2)) {
		goto L204
	} else {
		goto L205
	}
L204:
	;
	v1328 = v1325
	goto L206
L205:
	;
	v1328 = v1289
	goto L206
L206:
	;
	if base.Ui64(v1055+int64(4)) < base.Ui64(int64(28)) {
		goto L207
	} else {
		goto L208
	}
L207:
	;
	v1333 = v1328
	goto L209
L208:
	;
	v1333 = v1289
	goto L209
L209:
	;
	v1336 = v1333&int64(1) + v1333
	if base.Ui64(v1336) < base.Ui64(int64(18014398509481984)) {
		goto L197
	} else {
		goto L210
	}
L210:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v31)+224)) = int64(4503599627370496)
	v1342 = v1296 + int32(1087)
	*(*int32)(unsafe.Add(mBase, uint32(v31)+232)) = v1342
	v1351 = v1342
	v1352 = int64(0)
	goto L196
L211:
	;
	v1356 = v1351
	goto L213
L212:
	;
	v1356 = v1353
	goto L213
L213:
	;
	if base.Ui32(v1351) < base.Ui32(v1353) {
		goto L214
	} else {
		goto L215
	}
L214:
	;
	v1358 = v1352
	goto L216
L215:
	;
	v1358 = int64(0)
	goto L216
L216:
	;
	v1359 = v1356
	v1361 = v1358
	goto L161
L217:
	;
	v1815 = int64(0)
	if base.B2i32(v1797 == v1815)&base.B2i32(v1794 == int32(0))&base.B2i32(v1063 != v1815) != 0 {
		goto L265
	} else {
		goto L266
	}
L218:
	;
	v1369 = int32(0)
	v1370 = int64(0)
	v1372 = v1063 + int64(1)
	if v1372 == v1370 {
		v1570 = v1369
		v1573 = v1370
		goto L219
	} else {
		goto L220
	}
L219:
	;
	if v1361 != v1573 {
		goto L256
	} else {
		goto L257
	}
L220:
	;
	if v1055 < int64(-342) {
		v1570 = v1369
		v1573 = v1370
		goto L219
	} else {
		goto L221
	}
L221:
	;
	if v1055 <= int64(308) {
		goto L222
	} else {
		goto L223
	}
L222:
	;
	v1381 = v31 + int32(112)
	v1382 = base.I32_wrap_i64(v1055)
	v1387 = *(*int64)(unsafe.Add(mBase, uint32(v1382<<(uint(int32(4))%32))+uint32(_consts[995])))
	v1388 = int64(0)
	v1389 = base.I64_clz(v1372)
	v1390 = v1372 << (uint(v1389) % 64)
	v1396 = int64(32)
	v1397 = int64(base.Ui64(v1390) >> (uint(v1396) % 64))
	v1399 = int64(base.Ui64(v1387) >> (uint(v1396) % 64))
	v1402 = int64(4294967295)
	v1403 = v1390 & v1402
	v1405 = v1387 & v1402
	v1406 = v1403 * v1405
	v1410 = int64(base.Ui64(v1406)>>(uint(v1396)%64)) + v1403*v1399
	v1417 = v1410&v1402 + v1397*v1405
	*(*int64)(unsafe.Add(mBase, uint32(v1381)+8)) = v1388*v1387 + v1388*v1390 + v1397*v1399 + int64(base.Ui64(v1410)>>(uint(v1396)%64)) + int64(base.Ui64(v1417)>>(uint(v1396)%64))
	*(*int64)(unsafe.Add(mBase, uint32(v1381))) = v1417<<(uint(v1396)%64) | v1406&v1402
	goto L224
L223:
	;
	v1570 = int32(2047)
	v1573 = v1370
	goto L219
L224:
	;
	v1429 = *(*int64)(unsafe.Add(mBase, uint32(v31)+112))
	v1434 = *(*int64)(unsafe.Add(mBase, uint32(v31+int32(120))))
	v1435 = int64(511)
	if v1434&v1435 != v1435 {
		v1496 = v1429
		v1497 = v1434
		goto L225
	} else {
		goto L226
	}
L225:
	;
	v1499 = int64(base.Ui64(v1497) >> (uint(int64(63)) % 64))
	v1501 = v1499 + int64(9)
	v1502 = int64(base.Ui64(v1497) >> (uint(v1501) % 64))
	v1511 = v1382*int32(217706)>>(uint(int32(16))%32) - base.I32_wrap_i64(v1389) + base.I32_wrap_i64(v1499) + int32(63)
	if int32(-1023) < v1511 {
		goto L228
	} else {
		goto L229
	}
L226:
	;
	v1440 = v31 + int32(96)
	v1447 = *(*int64)(unsafe.Add(mBase, uint32(v1382<<(uint(int32(1))%32)<<(uint(int32(3))%32))+uint32(_consts[999])))
	v1448 = int64(0)
	v1454 = int64(32)
	v1455 = int64(base.Ui64(v1390) >> (uint(v1454) % 64))
	v1457 = int64(base.Ui64(v1447) >> (uint(v1454) % 64))
	v1460 = int64(4294967295)
	v1461 = v1390 & v1460
	v1463 = v1447 & v1460
	v1464 = v1461 * v1463
	v1468 = int64(base.Ui64(v1464)>>(uint(v1454)%64)) + v1461*v1457
	v1475 = v1468&v1460 + v1455*v1463
	*(*int64)(unsafe.Add(mBase, uint32(v1440)+8)) = v1448*v1447 + v1448*v1390 + v1455*v1457 + int64(base.Ui64(v1468)>>(uint(v1454)%64)) + int64(base.Ui64(v1475)>>(uint(v1454)%64))
	*(*int64)(unsafe.Add(mBase, uint32(v1440))) = v1475<<(uint(v1454)%64) | v1464&v1460
	goto L227
L227:
	;
	v1490 = *(*int64)(unsafe.Add(mBase, uint32(v31+int32(104))))
	v1491 = v1490 + v1429
	v1496 = v1491
	v1497 = base.I64_extend_i32_u(base.B2i32(base.Ui64(v1491) < base.Ui64(v1490))) + v1434
	goto L225
L228:
	;
	if v1502<<(uint(v1501)%64) == v1497 {
		goto L232
	} else {
		goto L233
	}
L229:
	;
	if base.Ui32(int32(-1085)) <= base.Ui32(v1511) {
		goto L230
	} else {
		goto L231
	}
L230:
	;
	v1520 = int64(base.Ui64(v1502) >> (uint(base.I64_extend_i32_u(int32(-1022)-v1511)) % 64))
	v1521 = int64(1)
	v1523 = v1520&v1521 + v1520
	v1570 = base.B2i32(base.Ui64(int64(9007199254740991)) < base.Ui64(v1523))
	v1573 = int64(base.Ui64(v1523) >> (uint(v1521) % 64))
	goto L219
L231:
	;
	v1570 = v1369
	v1573 = int64(0)
	goto L219
L232:
	;
	v1534 = v1502 & int64(72057594037927932)
	goto L234
L233:
	;
	v1534 = v1502
	goto L234
L234:
	;
	if v1502&int64(3) == int64(1) {
		goto L235
	} else {
		goto L236
	}
L235:
	;
	v1539 = v1534
	goto L237
L236:
	;
	v1539 = v1502
	goto L237
L237:
	;
	if base.Ui64(v1496) < base.Ui64(int64(2)) {
		goto L238
	} else {
		goto L239
	}
L238:
	;
	v1542 = v1539
	goto L240
L239:
	;
	v1542 = v1502
	goto L240
L240:
	;
	if base.Ui64(v1055+int64(4)) < base.Ui64(int64(28)) {
		goto L241
	} else {
		goto L242
	}
L241:
	;
	v1547 = v1542
	goto L243
L242:
	;
	v1547 = v1502
	goto L243
L243:
	;
	v1549 = v1547 & int64(1)
	v1552 = base.B2i32(base.Ui64(v1549+v1547) < base.Ui64(int64(18014398509481984)))
	if base.Ui64(v1549+v1547) < base.Ui64(int64(18014398509481984)) {
		goto L244
	} else {
		goto L245
	}
L244:
	;
	v1553 = int32(1023)
	goto L246
L245:
	;
	v1553 = int32(1024)
	goto L246
L246:
	;
	v1554 = v1553 + v1511
	v1555 = int32(2047)
	v1557 = base.B2i32(base.Ui32(v1554) < base.Ui32(v1555))
	if base.Ui32(v1554) < base.Ui32(v1555) {
		goto L247
	} else {
		goto L248
	}
L247:
	;
	v1558 = v1554
	goto L249
L248:
	;
	v1558 = v1555
	goto L249
L249:
	;
	if base.Ui32(v1554) < base.Ui32(v1555) {
		goto L250
	} else {
		goto L251
	}
L250:
	;
	v1567 = (v1549&v1547 + int64(base.Ui64(v1549^v1547)>>(uint(int64(1))%64))) & int64(67553994410557439)
	goto L252
L251:
	;
	v1567 = int64(0)
	goto L252
L252:
	;
	if base.Ui64(v1549+v1547) < base.Ui64(int64(18014398509481984)) {
		goto L253
	} else {
		goto L254
	}
L253:
	;
	v1569 = v1567
	goto L255
L254:
	;
	v1569 = int64(0)
	goto L255
L255:
	;
	v1570 = v1558
	v1573 = v1569
	goto L219
L256:
	;
	v1581 = v31 + int32(80)
	v1582 = base.I32_wrap_i64(v1055)
	v1587 = *(*int64)(unsafe.Add(mBase, uint32(v1582<<(uint(int32(4))%32))+uint32(_consts[995])))
	v1588 = int64(0)
	v1589 = base.I64_clz(v1063)
	v1590 = v1063 << (uint(v1589) % 64)
	v1596 = int64(32)
	v1597 = int64(base.Ui64(v1590) >> (uint(v1596) % 64))
	v1599 = int64(base.Ui64(v1587) >> (uint(v1596) % 64))
	v1602 = int64(4294967295)
	v1603 = v1590 & v1602
	v1605 = v1587 & v1602
	v1606 = v1603 * v1605
	v1610 = int64(base.Ui64(v1606)>>(uint(v1596)%64)) + v1603*v1599
	v1617 = v1610&v1602 + v1597*v1605
	*(*int64)(unsafe.Add(mBase, uint32(v1581)+8)) = v1588*v1587 + v1588*v1590 + v1597*v1599 + int64(base.Ui64(v1610)>>(uint(v1596)%64)) + int64(base.Ui64(v1617)>>(uint(v1596)%64))
	*(*int64)(unsafe.Add(mBase, uint32(v1581))) = v1617<<(uint(v1596)%64) | v1606&v1602
	goto L259
L257:
	;
	if v1359 == v1570 {
		v1794 = v1359
		v1797 = v1361
		goto L217
	} else {
		goto L258
	}
L258:
	;
	goto L256
L259:
	;
	v1633 = *(*int64)(unsafe.Add(mBase, uint32(v31+int32(88))))
	v1634 = int64(511)
	if v1633&v1634 != v1634 {
		v1696 = v1633
		goto L260
	} else {
		goto L261
	}
L260:
	;
	v1702 = base.I32_wrap_i64(int64(base.Ui64(v1696)>>(uint(int64(63))%64))) ^ int32(1)
	v1704 = v1696 << (uint(base.I64_extend_i32_u(v1702)) % 64)
	*(*int64)(unsafe.Add(mBase, uint32(v31)+224)) = v1704
	v1713 = v1582*int32(217706)>>(uint(int32(16))%32) - (v1702 + base.I32_wrap_i64(v1589)) + int32(-31692)
	*(*int32)(unsafe.Add(mBase, uint32(v31)+232)) = v1713
	if int32(-1) < v1713 {
		v1794 = v1713
		v1797 = v1704
		goto L217
	} else {
		goto L263
	}
L261:
	;
	v1638 = *(*int64)(unsafe.Add(mBase, uint32(v31)+80))
	v1640 = v31 + int32(64)
	v1647 = *(*int64)(unsafe.Add(mBase, uint32(v1582<<(uint(int32(1))%32)<<(uint(int32(3))%32))+uint32(_consts[999])))
	v1648 = int64(0)
	v1654 = int64(32)
	v1655 = int64(base.Ui64(v1590) >> (uint(v1654) % 64))
	v1657 = int64(base.Ui64(v1647) >> (uint(v1654) % 64))
	v1660 = int64(4294967295)
	v1661 = v1590 & v1660
	v1663 = v1647 & v1660
	v1664 = v1661 * v1663
	v1668 = int64(base.Ui64(v1664)>>(uint(v1654)%64)) + v1661*v1657
	v1675 = v1668&v1660 + v1655*v1663
	*(*int64)(unsafe.Add(mBase, uint32(v1640)+8)) = v1648*v1647 + v1648*v1590 + v1655*v1657 + int64(base.Ui64(v1668)>>(uint(v1654)%64)) + int64(base.Ui64(v1675)>>(uint(v1654)%64))
	*(*int64)(unsafe.Add(mBase, uint32(v1640))) = v1675<<(uint(v1654)%64) | v1664&v1660
	goto L262
L262:
	;
	v1690 = *(*int64)(unsafe.Add(mBase, uint32(v31+int32(72))))
	v1696 = base.I64_extend_i32_u(base.B2i32(base.Ui64(v1690^int64(-1)) < base.Ui64(v1638))) + v1633
	goto L260
L263:
	;
	v1725 = *(*int64)(unsafe.Add(mBase, uint32(v31+int32(168))))
	*(*int64)(unsafe.Add(mBase, uint32(v31+int32(24)))) = v1725
	v1735 = *(*int64)(unsafe.Add(mBase, uint32(v31+int32(176))))
	*(*int64)(unsafe.Add(mBase, uint32(v31+int32(32)))) = v1735
	v1745 = *(*int64)(unsafe.Add(mBase, uint32(v31+int32(184))))
	*(*int64)(unsafe.Add(mBase, uint32(v31+int32(40)))) = v1745
	v1755 = *(*int64)(unsafe.Add(mBase, uint32(v31+int32(192))))
	*(*int64)(unsafe.Add(mBase, uint32(v31+int32(48)))) = v1755
	v1765 = *(*int64)(unsafe.Add(mBase, uint32(v31+int32(200))))
	*(*int64)(unsafe.Add(mBase, uint32(v31+int32(56)))) = v1765
	v1772 = v31 + int32(232)
	v1773 = *(*int64)(unsafe.Add(mBase, uint32(v1772)))
	*(*int64)(unsafe.Add(mBase, uint32(v31+int32(8)))) = v1773
	v1775 = *(*int64)(unsafe.Add(mBase, uint32(v31)+160))
	*(*int64)(unsafe.Add(mBase, uint32(v31)+16)) = v1775
	v1777 = *(*int64)(unsafe.Add(mBase, uint32(v31)+224))
	*(*int64)(unsafe.Add(mBase, uint32(v31))) = v1777
	F_ffc_digit_comp(m, v31+int32(208), v31+int32(16), v31, int32(1))
	mBase = m.M
	v1789 = *(*int64)(unsafe.Add(mBase, uint32(v31+int32(216))))
	*(*int64)(unsafe.Add(mBase, uint32(v1772))) = v1789
	v1791 = *(*int64)(unsafe.Add(mBase, uint32(v31)+208))
	*(*int64)(unsafe.Add(mBase, uint32(v31)+224)) = v1791
	v1794 = base.I32_wrap_i64(v1789)
	v1797 = v1791
	goto L217
L264:
	;
	v1880 = base.F64_reinterpret_i64(base.I64_extend_i32_u(base.B2i32(v111 == int32(45)))<<(uint(int64(63))%64) | base.I64_extend_i32_u(v1794)<<(uint(int64(52))%64) | v1797)
	goto L8
L265:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = int32(2)
	goto L264
L266:
	;
	if v1794 != int32(2047) {
		goto L264
	} else {
		goto L267
	}
L267:
	;
	goto L265
}
func F_field(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
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
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
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
	var v45 int32
	_ = v45
	v6 = m.G0
	v8 = v6 - int32(32)
	m.G0 = v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v11 = F_luaK_exp2anyreg(m, v10, l1)
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		return
	} else {
		F_luaX_next(m, l0)
		mBase = m.M
		v14 = m.ExcPending
		if v14 != 0 {
			return
		} else {
			v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
			if v15 == int32(285) {
				v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
				F_luaX_next(m, l0)
				mBase = m.M
				v33 = m.ExcPending
				if v33 != 0 {
					return
				} else {
					v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
					v35 = F_luaK_stringK(m, v34, v31)
					mBase = m.M
					v36 = m.ExcPending
					if v36 != 0 {
						return
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = v35
						*(*int32)(unsafe.Add(mBase, uint32(v8)+8)) = int32(4)
						*(*int64)(unsafe.Add(mBase, uint32(v8)+24)) = int64(-1)
						F_luaK_indexed(m, v10, l1, v8+int32(8))
						mBase = m.M
						v45 = m.ExcPending
						if v45 != 0 {
							return
						} else {
							m.G0 = v8 + int32(32)
							return
						}
					}
				}
			} else {
				v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
				v20 = F_luaX_token2str(m, l0, int32(285))
				mBase = m.M
				v21 = m.ExcPending
				if v21 != 0 {
					return
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v8))) = v20
					v23 = m.G3
					v26 = F_luaO_pushfstring(m, v18, v23+int32(_a2038), v8)
					mBase = m.M
					v27 = m.ExcPending
					if v27 != 0 {
						return
					} else {
						F_luaX_syntaxerror(m, l0, v26)
						mBase = m.M
						v29 = m.ExcPending
						if v29 != 0 {
							return
						} else {
							v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
							F_luaX_next(m, l0)
							mBase = m.M
							v33 = m.ExcPending
							if v33 != 0 {
								return
							} else {
								v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
								v35 = F_luaK_stringK(m, v34, v31)
								mBase = m.M
								v36 = m.ExcPending
								if v36 != 0 {
									return
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = v35
									*(*int32)(unsafe.Add(mBase, uint32(v8)+8)) = int32(4)
									*(*int64)(unsafe.Add(mBase, uint32(v8)+24)) = int64(-1)
									F_luaK_indexed(m, v10, l1, v8+int32(8))
									mBase = m.M
									v45 = m.ExcPending
									if v45 != 0 {
										return
									} else {
										m.G0 = v8 + int32(32)
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
func F_fieldExpireScanCallback(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v12 int64
	_ = v12
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
	var v25 int32
	_ = v25
	var v33 int32
	_ = v33
	var v39 int32
	_ = v39
	if l1 == int32(0) {
		F__serverAssert(m, int32(_a782), int32(_a783), int32(168))
		mBase = m.M
		v33 = m.ExcPending
		if v33 != 0 {
			return
		} else {
			F_abort(m)
			mBase = m.M
			base.Wasm_trap_unreachable()
			for {
			}
		}
	} else {
		v6 = F_hashTypeHasVolatileFields(m, l1)
		mBase = m.M
		v7 = m.ExcPending
		if v7 != 0 {
			return
		} else {
			if v6 == int32(0) {
				F__serverAssert(m, int32(_a784), int32(_a783), int32(169))
				mBase = m.M
				v39 = m.ExcPending
				if v39 != 0 {
					return
				} else {
					F_abort(m)
					mBase = m.M
					base.Wasm_trap_unreachable()
					for {
					}
				}
			} else {
				v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
				v12 = *(*int64)(unsafe.Add(mBase, _consts[35]))
				v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
				v14 = F_dbReclaimExpiredFields(m, l1, v10, v12, v13, l2)
				mBase = m.M
				v15 = m.ExcPending
				if v15 != 0 {
					return
				} else {
					if v14 == int32(0) {
					} else {
						v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
						*(*uint8)(unsafe.Add(mBase, uint32(l0)+40)) = uint8(base.B2i32(v14 == v18))
						v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
						*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v21 + int32(1)
					}
					v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
					*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v25 + int32(1)
					return
				}
			}
		}
	}
}
func F_fifoCreate(m *base.Module) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	v3 = F_valkey_malloc(m, int32(12))
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(v3)+8)) = int32(0)
		*(*int64)(unsafe.Add(mBase, uint32(v3))) = int64(0)
		return v3
	}
}
func F_fileExist(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	v4 = m.G0
	v6 = v4 - int32(96)
	m.G0 = v6
	v10 = F___fstatat(m, int32(-100), l0, v6, int32(0))
	mBase = m.M
	v11 = *(*int32)(unsafe.Add(mBase, uint32(v6)+4))
	m.G0 = v6 + int32(96)
	return base.B2i32(v10 == int32(0)) & base.B2i32(v11&int32(61440) == int32(32768))
}
func F_fileIsRDB(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	var v32 int32
	_ = v32
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v45 int64
	_ = v45
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v109 int32
	_ = v109
	var v114 int32
	_ = v114
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
	var v136 int32
	_ = v136
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	v5 = m.G0
	v7 = v5 - int32(128)
	m.G0 = v7
	v10 = F_fopen(m, l0, int32(_a86))
	mBase = m.M
	if v10 == int32(0) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v7)+16)) = l0
	v155 = F_iprintf(m, int32(_a1780), v7+int32(16))
	mBase = m.M
	v156 = m.ExcPending
	if v156 != 0 {
		goto L19
	} else {
		goto L47
	}
L2:
	;
	goto L44
L3:
	;
	v15 = *(*int32)(unsafe.Add(mBase, uint32(v10)+76))
	if int32(-1) < v15 {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	if int32(-1) < v32 {
		goto L12
	} else {
		goto L13
	}
L5:
	;
	if int32(-1) < v24 {
		v32 = v24
		goto L9
	} else {
		goto L10
	}
L6:
	;
	v19 = F___lockfile(m, v10)
	mBase = m.M
	v20 = *(*int32)(unsafe.Add(mBase, uint32(v10)+60))
	if v19 == int32(0) {
		v24 = v20
		goto L5
	} else {
		goto L8
	}
L7:
	;
	v18 = *(*int32)(unsafe.Add(mBase, uint32(v10)+60))
	v24 = v18
	goto L5
L8:
	;
	F___unlockfile(m, v10)
	mBase = m.M
	v24 = v20
	goto L5
L9:
	;
	goto L4
L10:
	;
	v28 = F___errno_location(m)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v28))) = int32(8)
	v32 = int32(-1)
	goto L9
L11:
	;
	if v42 == int32(-1) {
		goto L1
	} else {
		goto L14
	}
L12:
	;
	v41 = F___fstatat(m, v32, int32(_a139), v7+int32(32), int32(4096))
	mBase = m.M
	v42 = v41
	goto L11
L13:
	;
	v38 = F___syscall_ret(m, int32(-8))
	mBase = m.M
	v42 = v38
	goto L11
L14:
	;
	v45 = *(*int64)(unsafe.Add(mBase, uint32(v7)+56))
	if v45 != int64(0) {
		goto L17
	} else {
		goto L18
	}
L15:
	;
	m.G0 = v7 + int32(128)
	return v136
L16:
	;
	v136 = int32(0)
	goto L15
L17:
	;
	if v45 < int64(8) {
		goto L21
	} else {
		goto L22
	}
L18:
	;
	v48 = F_fclose(m, v10)
	mBase = m.M
	v51 = m.ExcPending
	if v51 != 0 {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	return int32(0)
L20:
	;
	goto L16
L21:
	;
	v133 = F_fclose(m, v10)
	mBase = m.M
	v134 = m.ExcPending
	if v134 != 0 {
		goto L19
	} else {
		goto L43
	}
L22:
	;
	v58 = F_fread(m, v7+int32(27), int32(5), int32(1), v10)
	mBase = m.M
	v59 = m.ExcPending
	if v59 != 0 {
		goto L19
	} else {
		goto L23
	}
L23:
	;
	if v58 != int32(1) {
		goto L21
	} else {
		goto L24
	}
L24:
	;
	v63 = v7 + int32(27)
	v64 = int32(_a614)
	v65 = int32(5)
	goto L29
L25:
	;
	if v129 != 0 {
		goto L21
	} else {
		goto L41
	}
L26:
	;
	v129 = int32(0)
	goto L25
L27:
	;
	v101 = v96
	v102 = v97
	v103 = v98
	goto L37
L28:
	;
	if v86 == int32(0) {
		goto L26
	} else {
		goto L35
	}
L29:
	;
	if (v64|v63)&int32(3) != 0 {
		v96 = v63
		v97 = v64
		v98 = v65
		goto L27
	} else {
		goto L30
	}
L30:
	;
	v73 = v63
	v74 = v64
	v75 = v65
	goto L31
L31:
	;
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v73)))
	v79 = *(*int32)(unsafe.Add(mBase, uint32(v74)))
	if v78 != v79 {
		v96 = v73
		v97 = v74
		v98 = v75
		goto L27
	} else {
		goto L33
	}
L32:
	;
	goto L28
L33:
	;
	v81 = int32(4)
	v82 = v74 + v81
	v84 = v73 + v81
	v86 = v75 + int32(-4)
	if base.Ui32(int32(3)) < base.Ui32(v86) {
		v73 = v84
		v74 = v82
		v75 = v86
		goto L31
	} else {
		goto L34
	}
L34:
	;
	goto L32
L35:
	;
	v96 = v84
	v97 = v82
	v98 = v86
	goto L27
L36:
	;
	v129 = v106 - v107
	goto L25
L37:
	;
	v106 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v101))))
	v107 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v102))))
	if v106 != v107 {
		goto L36
	} else {
		goto L39
	}
L39:
	;
	v109 = int32(1)
	v114 = v103 + int32(-1)
	if v114 == int32(0) {
		goto L26
	} else {
		goto L40
	}
L40:
	;
	v101 = v101 + v109
	v102 = v102 + v109
	v103 = v114
	goto L37
L41:
	;
	v130 = F_fclose(m, v10)
	mBase = m.M
	v131 = m.ExcPending
	if v131 != 0 {
		goto L19
	} else {
		goto L42
	}
L42:
	;
	v136 = int32(1)
	goto L15
L43:
	;
	goto L16
L44:
	;
	v142 = *(*int32)(unsafe.Add(mBase, _consts[5]))
	v143 = F___strerror_l(m, v142, v142)
	mBase = m.M
	goto L45
L45:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v7)+4)) = v143
	*(*int32)(unsafe.Add(mBase, uint32(v7))) = l0
	v147 = F_iprintf(m, int32(_a1781), v7)
	mBase = m.M
	v148 = m.ExcPending
	if v148 != 0 {
		goto L19
	} else {
		goto L46
	}
L46:
	;
	m.Env.Exit(m, int32(1))
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L47:
	;
	v157 = F_fclose(m, v10)
	mBase = m.M
	v158 = m.ExcPending
	if v158 != 0 {
		goto L19
	} else {
		goto L48
	}
L48:
	;
	m.Env.Exit(m, int32(1))
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_fillBucketHole(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v14 int32
	_ = v14
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
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v51 int32
	_ = v51
	var v57 int32
	_ = v57
	var v62 int32
	_ = v62
	var v68 int32
	_ = v68
	v9 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1))))
	if v9&int32(1) == int32(0) {
		F__serverAssert(m, int32(_a832), int32(_a827), int32(1002))
		mBase = m.M
		v62 = m.ExcPending
		if v62 != 0 {
			return
		} else {
			F_abort(m)
			mBase = m.M
			base.Wasm_trap_unreachable()
			for {
			}
		}
	} else {
		v14 = int32(1)
		if int32(base.Ui32(int32(base.Ui32(v9)>>(uint(v14)%32))&int32(4095))>>(uint(l2)%32))&v14 != 0 {
			F__serverAssert(m, int32(_a832), int32(_a827), int32(1002))
			mBase = m.M
			v62 = m.ExcPending
			if v62 != 0 {
				return
			} else {
				F_abort(m)
				mBase = m.M
				base.Wasm_trap_unreachable()
				for {
				}
			}
		} else {
			v25 = l1
			for {
				v29 = *(*int32)(unsafe.Add(mBase, uint32(v25)+60))
				v30 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v29))))
				if v30&int32(1) != 0 {
					v25 = v29
					continue
				} else {
					break
				}
				break
			}
			v36 = int32(base.Ui32(v30)>>(uint(int32(1))%32)) & int32(4095)
			if v36 != 0 {
				v38 = base.I32_ctz(v36)
				if int32(base.Ui32(v36)>>(uint(v38)%32))&int32(1) == int32(0) {
					F__serverAssert(m, int32(_a833), int32(_a827), int32(1014))
					mBase = m.M
					v68 = m.ExcPending
					if v68 != 0 {
						return
					} else {
						F_abort(m)
						mBase = m.M
						base.Wasm_trap_unreachable()
						for {
						}
					}
				} else {
					F_moveEntry(m, l1, l2, v29, v38)
					mBase = m.M
					v45 = m.ExcPending
					if v45 != 0 {
						return
					} else {
						v46 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v29))))
						v51 = int32(base.Ui32(v46)>>(uint(int32(1))%32)) & int32(4095)
						if v51&(v51+int32(-1)) != 0 {
							return
						} else {
							F_pruneLastBucket(m, l0, v25, v29, l3)
							mBase = m.M
							v57 = m.ExcPending
							if v57 != 0 {
								return
							} else {
								return
							}
						}
					}
				}
			} else {
				v51 = int32(0)
				if v51&(v51+int32(-1)) != 0 {
					return
				} else {
					F_pruneLastBucket(m, l0, v25, v29, l3)
					mBase = m.M
					v57 = m.ExcPending
					if v57 != 0 {
						return
					} else {
						return
					}
				}
			}
		}
	}
}
func F_filterInvalidLogfmtChar(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v14 int32
	_ = v14
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
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
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v68 int32
	_ = v68
	if l1 != int32(1024) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	F__serverAssert(m, int32(_a1554), int32(_a1555), int32(163))
	mBase = m.M
	v68 = m.ExcPending
	if v68 != 0 {
		goto L18
	} else {
		goto L19
	}
L2:
	;
	if l2 == int32(0) {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	return
L4:
	;
	v14 = int32(0)
	goto L6
L5:
	;
	v55 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l0+v53))) = uint8(v55)
	goto L3
L6:
	;
	v20 = int32(39)
	v23 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2+v14))))
	switch v23 {
	case 0:
		v53 = v14
		goto L5
	default:
		goto L9
	case 10, 13:
		goto L10
	case 34:
		v25 = v20
		goto L8
	}
L7:
	;
	v53 = v45
	goto L5
L8:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(l0+v14))) = uint8(v25)
	v29 = v14 + int32(1)
	v31 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2+v29))))
	switch v31 {
	case 0:
		v53 = v29
		goto L5
	default:
		goto L12
	case 10, 13:
		goto L13
	case 34:
		v33 = v20
		goto L11
	}
L9:
	;
	v25 = v23
	goto L8
L10:
	;
	v25 = int32(32)
	goto L8
L11:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(l0+v29))) = uint8(v33)
	v38 = v14 + int32(2)
	v40 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2+v38))))
	switch v40 {
	case 0:
		v53 = v38
		goto L5
	default:
		goto L15
	case 10, 13:
		goto L16
	case 34:
		v42 = int32(39)
		goto L14
	}
L12:
	;
	v33 = v31
	goto L11
L13:
	;
	v33 = int32(32)
	goto L11
L14:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(l0+v38))) = uint8(v42)
	v45 = int32(1023)
	v47 = v14 + int32(3)
	if v47 != v45 {
		v14 = v47
		goto L6
	} else {
		goto L17
	}
L15:
	;
	v42 = v40
	goto L14
L16:
	;
	v42 = int32(32)
	goto L14
L17:
	;
	goto L7
L18:
	;
	return
L19:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_findBucket_1(m *base.Module, l0 int32, l1 int64, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
	var v55 int32
	_ = v55
	var v62 int32
	_ = v62
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v92 int32
	_ = v92
	var v98 int32
	_ = v98
	var v117 int32
	_ = v117
	var v127 int32
	_ = v127
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v150 int32
	_ = v150
	var v160 int32
	_ = v160
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v201 int32
	_ = v201
	var v203 int32
	_ = v203
	var v208 int32
	_ = v208
	var v230 int32
	_ = v230
	var v233 int32
	_ = v233
	var v239 int32
	_ = v239
	v6 = int32(0)
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v23 == v6-v25 {
		v239 = v6
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return v239
L2:
	;
	v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v30 == int32(-1) {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v55 = int32(1)
	v62 = int32(0)
	goto L9
L4:
	;
	v33 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+26)))
	if v33 != 0 {
		goto L3
	} else {
		goto L5
	}
L5:
	;
	v35 = *(*int32)(unsafe.Add(mBase, _consts[427]))
	if v35 != 0 {
		goto L3
	} else {
		goto L6
	}
L6:
	;
	F_rehashStep(m, l0)
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	return int32(0)
L8:
	;
	goto L3
L9:
	;
	v72 = v62 << (uint(int32(2)) % 32)
	v74 = *(*int32)(unsafe.Add(mBase, uint32(l0+int32(16)+v72)))
	if v74 == int32(0) {
		goto L11
	} else {
		goto L12
	}
L10:
	;
	v239 = v233
	goto L1
L11:
	;
	v230 = int32(1)
	v233 = int32(0)
	if v55&v230 != 0 {
		v55 = v233
		v62 = v230
		goto L9
	} else {
		goto L44
	}
L12:
	;
	v78 = int32(-1)
	v80 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0+int32(24)+v62))))
	if v80 == int32(255) {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	v86 = int32(0)
	goto L15
L14:
	;
	v86 = v78<<(uint(v80)%32) ^ v78
	goto L15
L15:
	;
	v87 = v86 & base.I32_wrap_i64(l1)
	if v55&int32(1) == int32(0) {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	v98 = *(*int32)(unsafe.Add(mBase, uint32(l0+int32(8)+v72)))
	v117 = v98 + v87<<(uint(int32(6))%32)
	goto L20
L17:
	;
	v92 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v92 < int32(0) {
		goto L16
	} else {
		goto L18
	}
L18:
	;
	if base.Ui32(v87) < base.Ui32(v92) {
		goto L11
	} else {
		goto L19
	}
L19:
	;
	goto L16
L20:
	;
	v127 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v117))))
	v142 = int32(0)
	v143 = v127
	goto L22
L21:
	;
	goto L11
L22:
	;
	v150 = int32(1)
	if int32(base.Ui32(int32(base.Ui32(v143)>>(uint(v150)%32))&int32(4095))>>(uint(v142)%32))&v150 == int32(0) {
		goto L24
	} else {
		goto L25
	}
L23:
	;
	if v203 == int32(0) {
		goto L11
	} else {
		goto L42
	}
L24:
	;
	v198 = int32(1)
	v199 = v142 + v198
	v201 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v117))))
	v203 = v201 & v198
	if base.Ui32(v199) < base.Ui32(int32(12)-v203) {
		v142 = v199
		v143 = v201
		goto L22
	} else {
		goto L41
	}
L25:
	;
	v160 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v117+int32(2)+v142))))
	if v160 != base.I32_wrap_i64(int64(base.Ui64(l1)>>(uint(int64(56))%64))) {
		goto L24
	} else {
		goto L26
	}
L26:
	;
	v165 = *(*int32)(unsafe.Add(mBase, uint32(v117+int32(16)+v142<<(uint(int32(2))%32))))
	v166 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v167 = *(*int32)(unsafe.Add(mBase, uint32(v166)))
	if v167 == int32(0) {
		v173 = v165
		v174 = v166
		goto L27
	} else {
		goto L28
	}
L27:
	;
	v175 = *(*int32)(unsafe.Add(mBase, uint32(v174)+8))
	if v175 == int32(0) {
		goto L31
	} else {
		goto L32
	}
L28:
	;
	v170 = m.T0[v167].(func(*base.Module, int32) int32)(m, v165)
	mBase = m.M
	v171 = m.ExcPending
	if v171 != 0 {
		goto L7
	} else {
		goto L29
	}
L29:
	;
	v172 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v173 = v170
	v174 = v172
	goto L27
L30:
	;
	if v181 != 0 {
		goto L24
	} else {
		goto L34
	}
L31:
	;
	v181 = base.B2i32(l2 != v173)
	goto L30
L32:
	;
	v178 = m.T0[v175].(func(*base.Module, int32, int32) int32)(m, l2, v173)
	mBase = m.M
	v179 = m.ExcPending
	if v179 != 0 {
		goto L7
	} else {
		goto L33
	}
L33:
	;
	v181 = v178
	goto L30
L34:
	;
	v182 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v183 = *(*int32)(unsafe.Add(mBase, uint32(v182)+12))
	if v183 == int32(0) {
		goto L35
	} else {
		goto L36
	}
L35:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v142
	if l4 != 0 {
		goto L39
	} else {
		goto L40
	}
L36:
	;
	v186 = m.T0[v183].(func(*base.Module, int32, int32) int32)(m, l0, v165)
	mBase = m.M
	v187 = m.ExcPending
	if v187 != 0 {
		goto L7
	} else {
		goto L37
	}
L37:
	;
	if v186 == int32(0) {
		goto L24
	} else {
		goto L38
	}
L38:
	;
	goto L35
L39:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l4))) = v62
	return v117
L40:
	;
	return v117
L41:
	;
	goto L23
L42:
	;
	v208 = *(*int32)(unsafe.Add(mBase, uint32(v117)+60))
	if v208 != 0 {
		v117 = v208
		goto L20
	} else {
		goto L43
	}
L43:
	;
	goto L21
L44:
	;
	goto L10
}
func F_fixedpoint_d2string(m *base.Module, l0 int32, l1 int32, l2 float64, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v21 int32
	_ = v21
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v39 float64
	_ = v39
	var v41 float64
	_ = v41
	var v47 int64
	_ = v47
	var v49 int64
	_ = v49
	var v53 int32
	_ = v53
	var v57 int32
	_ = v57
	var v62 int32
	_ = v62
	var v63 int64
	_ = v63
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v73 int32
	_ = v73
	var v78 int64
	_ = v78
	var v99 int32
	_ = v99
	var v104 int32
	_ = v104
	var v111 int32
	_ = v111
	var v116 int32
	_ = v116
	var v118 int32
	_ = v118
	var v122 int64
	_ = v122
	var v127 int32
	_ = v127
	var v131 int32
	_ = v131
	var v134 int32
	_ = v134
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v149 int32
	_ = v149
	var v153 int32
	_ = v153
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v171 int32
	_ = v171
	var v173 int64
	_ = v173
	var v179 int32
	_ = v179
	var v182 int64
	_ = v182
	var v183 int64
	_ = v183
	var v189 int32
	_ = v189
	var v192 int32
	_ = v192
	var v196 int32
	_ = v196
	var v201 int32
	_ = v201
	var v203 int32
	_ = v203
	var v209 int32
	_ = v209
	var v215 int64
	_ = v215
	var v218 int32
	_ = v218
	var v224 int32
	_ = v224
	var v226 int32
	_ = v226
	var v231 int32
	_ = v231
	var v234 int32
	_ = v234
	var v238 int32
	_ = v238
	var v243 int32
	_ = v243
	var v247 int32
	_ = v247
	if base.Ui32(l3+int32(-18)) < base.Ui32(int32(-17)) {
		v136 = l0
		v137 = l1
		goto L3
	} else {
		goto L4
	}
L1:
	;
	v247 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v136))) = uint8(v247)
	return v247
L2:
	;
	v149 = v134 - l3
	if int32(0) < v149 {
		v156 = v149
		goto L48
	} else {
		goto L49
	}
L3:
	;
	if v137 != 0 {
		goto L1
	} else {
		goto L47
	}
L4:
	;
	if l1 < l3+int32(3) {
		v136 = l0
		v137 = l1
		goto L3
	} else {
		goto L5
	}
L5:
	;
	if base.F64_ne(l2, float64(0)) != 0 {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	v34 = int32(0)
	v39 = *(*float64)(unsafe.Add(mBase, uint32(l3<<(uint(int32(3))%32))+uint32(_consts[956])))
	v41 = F_rint(m, base.F64_mul(l2, v39))
	mBase = m.M
	if base.F64_lt(base.F64_abs(v41), float64(9.223372036854776e+18)) == v34 {
		goto L12
	} else {
		goto L13
	}
L7:
	;
	v21 = int32(11824)
	*(*uint16)(unsafe.Add(mBase, uint32(l0))) = uint16(v21)
	v27 = F__emscripten_memset_bulkmem(m, l0+int32(2), base.I32_extend8_s(int32(48)), l3)
	mBase = m.M
	goto L8
L8:
	;
	v29 = l3 + int32(2)
	v31 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l0+v29))) = uint8(v31)
	return v29
L9:
	;
	v66 = int32(1)
	if base.Ui64(v63) < base.Ui64(int64(10)) {
		v127 = v34
		v131 = v66
		goto L15
	} else {
		goto L16
	}
L10:
	;
	v53 = int32(45)
	*(*uint8)(unsafe.Add(mBase, uint32(l0))) = uint8(v53)
	v57 = int32(1)
	v62 = l1 + int32(-1)
	v63 = int64(0) - v49
	v64 = v57
	v65 = l0 + v57
	goto L9
L11:
	;
	if v49 <= int64(-1) {
		goto L10
	} else {
		goto L14
	}
L12:
	;
	v49 = int64(-9223372036854775807 - 1)
	goto L11
L13:
	;
	v47 = base.I64_trunc_f64_s(v41)
	v49 = v47
	goto L11
L14:
	;
	v62 = l1
	v63 = v49
	v64 = int32(0)
	v65 = l0
	goto L9
L15:
	;
	v134 = v131 + v127
	if base.Ui32(v134) < base.Ui32(v62) {
		goto L2
	} else {
		goto L46
	}
L16:
	;
	v73 = v34
	v78 = v63
	goto L17
L17:
	;
	if base.Ui64(int64(100)) <= base.Ui64(v78) {
		goto L19
	} else {
		goto L20
	}
L18:
	;
	v127 = v118
	v131 = v66
	goto L15
L19:
	;
	if base.Ui64(int64(1000)) <= base.Ui64(v78) {
		goto L21
	} else {
		goto L22
	}
L20:
	;
	v127 = v73
	v131 = int32(2)
	goto L15
L21:
	;
	if base.Ui64(int64(999999999999)) < base.Ui64(v78) {
		goto L23
	} else {
		goto L24
	}
L22:
	;
	v127 = v73
	v131 = int32(3)
	goto L15
L23:
	;
	v118 = v73 + int32(12)
	v122 = base.I64_div_u_s(v78, int64(1000000000000))
	if base.Ui64(int64(9999999999999)) < base.Ui64(v78) {
		v73 = v118
		v78 = v122
		goto L17
	} else {
		goto L45
	}
L24:
	;
	if base.Ui64(int64(99999999)) < base.Ui64(v78) {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	if base.Ui64(int64(9999999999)) < base.Ui64(v78) {
		goto L37
	} else {
		goto L38
	}
L26:
	;
	if base.Ui64(int64(999999)) < base.Ui64(v78) {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	if base.Ui64(int64(9999999)) < base.Ui64(v78) {
		goto L34
	} else {
		goto L35
	}
L28:
	;
	if base.Ui64(int64(10000)) <= base.Ui64(v78) {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	if base.Ui64(int64(99999)) < base.Ui64(v78) {
		goto L31
	} else {
		goto L32
	}
L30:
	;
	v127 = v73
	v131 = int32(4)
	goto L15
L31:
	;
	v99 = int32(6)
	goto L33
L32:
	;
	v99 = int32(5)
	goto L33
L33:
	;
	v127 = v73
	v131 = v99
	goto L15
L34:
	;
	v104 = int32(8)
	goto L36
L35:
	;
	v104 = int32(7)
	goto L36
L36:
	;
	v127 = v73
	v131 = v104
	goto L15
L37:
	;
	if base.Ui64(int64(99999999999)) < base.Ui64(v78) {
		goto L42
	} else {
		goto L43
	}
L38:
	;
	if base.Ui64(int64(999999999)) < base.Ui64(v78) {
		goto L39
	} else {
		goto L40
	}
L39:
	;
	v111 = int32(10)
	goto L41
L40:
	;
	v111 = int32(9)
	goto L41
L41:
	;
	v127 = v73
	v131 = v111
	goto L15
L42:
	;
	v116 = int32(12)
	goto L44
L43:
	;
	v116 = int32(11)
	goto L44
L44:
	;
	v127 = v73
	v131 = v116
	goto L15
L45:
	;
	goto L18
L46:
	;
	v136 = v65
	v137 = v62
	goto L3
L47:
	;
	return int32(0)
L48:
	;
	v157 = v65 + v156
	v158 = int32(46)
	*(*uint8)(unsafe.Add(mBase, uint32(v157))) = uint8(v158)
	v164 = F__emscripten_memset_bulkmem(m, v157+int32(1), base.I32_extend8_s(int32(48)), l3)
	mBase = m.M
	goto L50
L49:
	;
	v153 = int32(48)
	*(*uint8)(unsafe.Add(mBase, uint32(l0+v64))) = uint8(v153)
	v156 = int32(1)
	goto L48
L50:
	;
	v165 = v156 + l3
	if base.Ui64(int64(100)) <= base.Ui64(v63) {
		goto L52
	} else {
		goto L53
	}
L51:
	;
	v218 = v165 + int32(1)
	if base.Ui64(int64(9)) < base.Ui64(v215) {
		goto L61
	} else {
		goto L62
	}
L52:
	;
	v171 = v165
	v173 = v63
	goto L54
L53:
	;
	v209 = v165
	v215 = v63
	goto L51
L54:
	;
	v179 = v65 + v171
	v182 = int64(100)
	v183 = base.I64_div_u_s(v173, v182)
	v189 = base.I32_wrap_i64(v173-v183*v182) << (uint(int32(1)) % 32)
	v192 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v189)+uint32(_consts[723]))))
	*(*uint8)(unsafe.Add(mBase, uint32(v179+int32(-1)))) = uint8(v192)
	v196 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v189)+uint32(_consts[957]))))
	*(*uint8)(unsafe.Add(mBase, uint32(v179))) = uint8(v196)
	v201 = v171 + int32(-2)
	if v201 == v156 {
		goto L56
	} else {
		goto L57
	}
L55:
	;
	v209 = v203
	v215 = v183
	goto L51
L56:
	;
	v203 = v171 + int32(-3)
	goto L58
L57:
	;
	v203 = v201
	goto L58
L58:
	;
	if base.Ui64(int64(9999)) < base.Ui64(v173) {
		v171 = v203
		v173 = v183
		goto L54
	} else {
		goto L59
	}
L59:
	;
	goto L55
L60:
	;
	v243 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v65+v218))) = uint8(v243)
	return v218 + v64
L61:
	;
	v226 = v65 + v209
	v231 = base.I32_wrap_i64(v215) << (uint(int32(1)) % 32)
	v234 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v231)+uint32(_consts[723]))))
	*(*uint8)(unsafe.Add(mBase, uint32(v226+int32(-1)))) = uint8(v234)
	v238 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v231)+uint32(_consts[957]))))
	*(*uint8)(unsafe.Add(mBase, uint32(v226))) = uint8(v238)
	goto L60
L62:
	;
	v224 = base.I32_wrap_i64(v215) | int32(48)
	*(*uint8)(unsafe.Add(mBase, uint32(v65+v209))) = uint8(v224)
	goto L60
}
func F_floor(m *base.Module, l0 float64) float64 {
	return base.F64_floor(l0)
}
func F_fmodl(m *base.Module, l0 int32, l1 int64, l2 int64, l3 int64, l4 int64) {
	mBase := m.M
	_ = mBase
	var v9 int64
	_ = v9
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v26 int64
	_ = v26
	var v27 int64
	_ = v27
	var v31 int32
	_ = v31
	var v55 int32
	_ = v55
	var v67 int32
	_ = v67
	var v76 int32
	_ = v76
	var v80 int32
	_ = v80
	var v87 int64
	_ = v87
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v101 int32
	_ = v101
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v122 int64
	_ = v122
	var v127 int64
	_ = v127
	var v131 int64
	_ = v131
	var v132 int64
	_ = v132
	var v133 int64
	_ = v133
	var v134 int64
	_ = v134
	var v136 int64
	_ = v136
	var v140 int32
	_ = v140
	var v144 int64
	_ = v144
	var v145 int64
	_ = v145
	var v149 int32
	_ = v149
	var v153 int64
	_ = v153
	var v154 int64
	_ = v154
	var v158 int32
	_ = v158
	var v173 int32
	_ = v173
	var v185 int32
	_ = v185
	var v194 int32
	_ = v194
	var v198 int32
	_ = v198
	var v204 int32
	_ = v204
	var v208 int64
	_ = v208
	var v209 int64
	_ = v209
	var v213 int32
	_ = v213
	var v217 int64
	_ = v217
	var v218 int64
	_ = v218
	var v222 int32
	_ = v222
	var v237 int32
	_ = v237
	var v249 int32
	_ = v249
	var v258 int32
	_ = v258
	var v262 int32
	_ = v262
	var v267 int64
	_ = v267
	var v272 int64
	_ = v272
	var v273 int64
	_ = v273
	var v278 int32
	_ = v278
	var v288 int64
	_ = v288
	var v294 int64
	_ = v294
	var v295 int64
	_ = v295
	var v296 int32
	_ = v296
	var v297 int64
	_ = v297
	var v305 int64
	_ = v305
	var v311 int64
	_ = v311
	var v312 int64
	_ = v312
	var v313 int64
	_ = v313
	var v314 int32
	_ = v314
	var v315 int64
	_ = v315
	var v317 int64
	_ = v317
	var v318 int64
	_ = v318
	var v322 int64
	_ = v322
	var v328 int64
	_ = v328
	var v331 int32
	_ = v331
	var v332 int64
	_ = v332
	var v339 int64
	_ = v339
	var v342 int64
	_ = v342
	var v348 int64
	_ = v348
	var v353 int64
	_ = v353
	var v354 int64
	_ = v354
	var v365 int64
	_ = v365
	var v366 int64
	_ = v366
	var v368 int64
	_ = v368
	var v370 int32
	_ = v370
	var v376 int64
	_ = v376
	var v379 int32
	_ = v379
	var v380 int64
	_ = v380
	var v387 int64
	_ = v387
	var v390 int64
	_ = v390
	var v396 int64
	_ = v396
	var v401 int64
	_ = v401
	var v402 int64
	_ = v402
	var v403 int64
	_ = v403
	var v404 int64
	_ = v404
	var v411 int64
	_ = v411
	var v414 int32
	_ = v414
	var v416 int64
	_ = v416
	var v422 int32
	_ = v422
	var v423 int64
	_ = v423
	var v424 int64
	_ = v424
	var v427 int64
	_ = v427
	var v434 int64
	_ = v434
	var v437 int32
	_ = v437
	var v439 int64
	_ = v439
	var v443 int32
	_ = v443
	var v462 int64
	_ = v462
	var v463 int64
	_ = v463
	var v473 int64
	_ = v473
	var v475 int64
	_ = v475
	v9 = int64(0)
	v13 = m.G0
	v15 = v13 - int32(128)
	m.G0 = v15
	v26 = l4 & int64(9223372036854775807)
	v27 = int64(9223090561878065152)
	if v26 == v27 {
		goto L6
	} else {
		goto L7
	}
L1:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l0))) = v475
	*(*int64)(unsafe.Add(mBase, uint32(l0)+8)) = v473
	m.G0 = v15 + int32(128)
	return
L2:
	;
	v133 = int64(9223372036854775807)
	v134 = l2 & v133
	v136 = l4 & v133
	v140 = int32(1)
	v144 = v134 & v133
	v145 = int64(9223090561878065152)
	if v144 == v145 {
		goto L42
	} else {
		goto L43
	}
L3:
	;
	F___multf3(m, v15+int32(16), l1, l2, l3, l4)
	mBase = m.M
	v122 = *(*int64)(unsafe.Add(mBase, uint32(v15)+16))
	v127 = *(*int64)(unsafe.Add(mBase, uint32(v15+int32(24))))
	F___divtf3(m, v15, v122, v127, v122, v127)
	mBase = m.M
	v131 = *(*int64)(unsafe.Add(mBase, uint32(v15+int32(8))))
	v132 = *(*int64)(unsafe.Add(mBase, uint32(v15)))
	v473 = v131
	v475 = v132
	goto L1
L4:
	;
	if v80 == int32(0) {
		goto L3
	} else {
		goto L28
	}
L5:
	;
	v80 = v76
	goto L4
L6:
	;
	v31 = base.B2i32(l3 != v9)
	goto L8
L7:
	;
	v31 = base.B2i32(base.Ui64(v27) < base.Ui64(v26))
	goto L8
L8:
	;
	if v31 != 0 {
		v76 = int32(1)
		goto L5
	} else {
		goto L9
	}
L9:
	;
	goto L11
L11:
	;
	goto L12
L12:
	;
	goto L13
L13:
	;
	if base.B2i32(v9|l3|(int64(0)|v26) == int64(0)) == int32(0) {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	if v9&l4 < int64(0) {
		goto L16
	} else {
		goto L17
	}
L15:
	;
	v80 = int32(0)
	goto L4
L16:
	;
	if l4 == v9 {
		goto L24
	} else {
		goto L25
	}
L17:
	;
	if l4 == v9 {
		goto L19
	} else {
		goto L20
	}
L18:
	;
	v80 = base.B2i32(l3^v9|(l4^v9) != int64(0))
	goto L4
L19:
	;
	v55 = base.B2i32(base.Ui64(l3) < base.Ui64(v9))
	goto L21
L20:
	;
	v55 = base.B2i32(l4 < v9)
	goto L21
L21:
	;
	if v55 == int32(0) {
		goto L18
	} else {
		goto L22
	}
L22:
	;
	v80 = int32(-1)
	goto L4
L23:
	;
	v76 = base.B2i32(l3^v9|(l4^v9) != int64(0))
	goto L5
L24:
	;
	v67 = base.B2i32(base.Ui64(v9) < base.Ui64(l3))
	goto L26
L25:
	;
	v67 = base.B2i32(v9 < l4)
	goto L26
L26:
	;
	if v67 == int32(0) {
		goto L23
	} else {
		goto L27
	}
L27:
	;
	v80 = int32(-1)
	goto L4
L28:
	;
	v87 = l4 & int64(281474976710655)
	v91 = int32(32767)
	v92 = base.I32_wrap_i64(int64(base.Ui64(l4)>>(uint(int64(48))%64))) & v91
	if v92 == v91 {
		goto L31
	} else {
		goto L32
	}
L29:
	;
	if v107 == int32(0) {
		goto L3
	} else {
		goto L37
	}
L30:
	;
	v107 = v105
	goto L29
L31:
	;
	v105 = base.B2i32(v87|l3 == int64(0))
	goto L30
L32:
	;
	if v92 != 0 {
		v105 = int32(4)
		goto L30
	} else {
		goto L33
	}
L33:
	;
	if v87|l3 == int64(0) {
		goto L34
	} else {
		goto L35
	}
L34:
	;
	v101 = int32(2)
	goto L36
L35:
	;
	v101 = int32(3)
	goto L36
L36:
	;
	v107 = v101
	goto L29
L37:
	;
	v112 = base.I32_wrap_i64(int64(base.Ui64(l2) >> (uint(int64(48)) % 64)))
	v113 = int32(32767)
	v114 = v112 & v113
	if v114 != v113 {
		goto L2
	} else {
		goto L38
	}
L38:
	;
	goto L3
L39:
	;
	v278 = base.I32_wrap_i64(int64(base.Ui64(l4)>>(uint(int64(48))%64))) & int32(32767)
	if v114 == int32(0) {
		goto L92
	} else {
		goto L93
	}
L40:
	;
	if int32(0) < v198 {
		goto L39
	} else {
		goto L64
	}
L41:
	;
	v198 = v194
	goto L40
L42:
	;
	v149 = base.B2i32(l1 != int64(0))
	goto L44
L43:
	;
	v149 = base.B2i32(base.Ui64(v145) < base.Ui64(v144))
	goto L44
L44:
	;
	if v149 != 0 {
		v194 = v140
		goto L41
	} else {
		goto L45
	}
L45:
	;
	v153 = v136 & int64(9223372036854775807)
	v154 = int64(9223090561878065152)
	if v153 == v154 {
		goto L46
	} else {
		goto L47
	}
L46:
	;
	v158 = base.B2i32(l3 != int64(0))
	goto L48
L47:
	;
	v158 = base.B2i32(base.Ui64(v154) < base.Ui64(v153))
	goto L48
L48:
	;
	if v158 != 0 {
		v194 = v140
		goto L41
	} else {
		goto L49
	}
L49:
	;
	if base.B2i32(l3|l1|(v153|v144) == int64(0)) == int32(0) {
		goto L50
	} else {
		goto L51
	}
L50:
	;
	if v136&v134 < int64(0) {
		goto L52
	} else {
		goto L53
	}
L51:
	;
	v198 = int32(0)
	goto L40
L52:
	;
	if v134 == v136 {
		goto L60
	} else {
		goto L61
	}
L53:
	;
	if v134 == v136 {
		goto L55
	} else {
		goto L56
	}
L54:
	;
	v198 = base.B2i32(l1^l3|(v134^v136) != int64(0))
	goto L40
L55:
	;
	v173 = base.B2i32(base.Ui64(l1) < base.Ui64(l3))
	goto L57
L56:
	;
	v173 = base.B2i32(v134 < v136)
	goto L57
L57:
	;
	if v173 == int32(0) {
		goto L54
	} else {
		goto L58
	}
L58:
	;
	v198 = int32(-1)
	goto L40
L59:
	;
	v194 = base.B2i32(l1^l3|(v134^v136) != int64(0))
	goto L41
L60:
	;
	v185 = base.B2i32(base.Ui64(l3) < base.Ui64(l1))
	goto L62
L61:
	;
	v185 = base.B2i32(v136 < v134)
	goto L62
L62:
	;
	if v185 == int32(0) {
		goto L59
	} else {
		goto L63
	}
L63:
	;
	v198 = int32(-1)
	goto L40
L64:
	;
	v204 = int32(1)
	v208 = v134 & int64(9223372036854775807)
	v209 = int64(9223090561878065152)
	if v208 == v209 {
		goto L68
	} else {
		goto L69
	}
L65:
	;
	v267 = int64(0)
	F___multf3(m, v15+int32(112), l1, l2, v267, v267)
	mBase = m.M
	v272 = *(*int64)(unsafe.Add(mBase, uint32(v15+int32(120))))
	v273 = *(*int64)(unsafe.Add(mBase, uint32(v15)+112))
	v473 = v272
	v475 = v273
	goto L1
L66:
	;
	if v262 == int32(0) {
		goto L65
	} else {
		goto L90
	}
L67:
	;
	v262 = v258
	goto L66
L68:
	;
	v213 = base.B2i32(l1 != int64(0))
	goto L70
L69:
	;
	v213 = base.B2i32(base.Ui64(v209) < base.Ui64(v208))
	goto L70
L70:
	;
	if v213 != 0 {
		v258 = v204
		goto L67
	} else {
		goto L71
	}
L71:
	;
	v217 = v136 & int64(9223372036854775807)
	v218 = int64(9223090561878065152)
	if v217 == v218 {
		goto L72
	} else {
		goto L73
	}
L72:
	;
	v222 = base.B2i32(l3 != int64(0))
	goto L74
L73:
	;
	v222 = base.B2i32(base.Ui64(v218) < base.Ui64(v217))
	goto L74
L74:
	;
	if v222 != 0 {
		v258 = v204
		goto L67
	} else {
		goto L75
	}
L75:
	;
	if base.B2i32(l3|l1|(v217|v208) == int64(0)) == int32(0) {
		goto L76
	} else {
		goto L77
	}
L76:
	;
	if v136&v134 < int64(0) {
		goto L78
	} else {
		goto L79
	}
L77:
	;
	v262 = int32(0)
	goto L66
L78:
	;
	if v134 == v136 {
		goto L86
	} else {
		goto L87
	}
L79:
	;
	if v134 == v136 {
		goto L81
	} else {
		goto L82
	}
L80:
	;
	v262 = base.B2i32(l1^l3|(v134^v136) != int64(0))
	goto L66
L81:
	;
	v237 = base.B2i32(base.Ui64(l1) < base.Ui64(l3))
	goto L83
L82:
	;
	v237 = base.B2i32(v134 < v136)
	goto L83
L83:
	;
	if v237 == int32(0) {
		goto L80
	} else {
		goto L84
	}
L84:
	;
	v262 = int32(-1)
	goto L66
L85:
	;
	v258 = base.B2i32(l1^l3|(v134^v136) != int64(0))
	goto L67
L86:
	;
	v249 = base.B2i32(base.Ui64(l3) < base.Ui64(l1))
	goto L88
L87:
	;
	v249 = base.B2i32(v136 < v134)
	goto L88
L88:
	;
	if v249 == int32(0) {
		goto L85
	} else {
		goto L89
	}
L89:
	;
	v262 = int32(-1)
	goto L66
L90:
	;
	v473 = l2
	v475 = l1
	goto L1
L91:
	;
	if v278 != 0 {
		v312 = l3
		v313 = v136
		v314 = v278
		goto L94
	} else {
		goto L95
	}
L92:
	;
	F___multf3(m, v15+int32(96), l1, v134, int64(0), int64(4645181540655955968))
	mBase = m.M
	v288 = *(*int64)(unsafe.Add(mBase, uint32(v15+int32(104))))
	v294 = *(*int64)(unsafe.Add(mBase, uint32(v15)+96))
	v295 = v294
	v296 = base.I32_wrap_i64(int64(base.Ui64(v288)>>(uint(int64(48))%64))) + int32(-120)
	v297 = v288
	goto L91
L93:
	;
	v295 = l1
	v296 = v114
	v297 = v134
	goto L91
L94:
	;
	v315 = int64(281474976710655)
	v317 = int64(281474976710656)
	v318 = v313&v315 | v317
	v322 = v297&v315 | v317
	if v296 <= v314 {
		v376 = v295
		v379 = v296
		v380 = v322
		goto L96
	} else {
		goto L97
	}
L95:
	;
	F___multf3(m, v15+int32(80), l3, v136, int64(0), int64(4645181540655955968))
	mBase = m.M
	v305 = *(*int64)(unsafe.Add(mBase, uint32(v15+int32(88))))
	v311 = *(*int64)(unsafe.Add(mBase, uint32(v15)+80))
	v312 = v311
	v313 = v305
	v314 = base.I32_wrap_i64(int64(base.Ui64(v305)>>(uint(int64(48))%64))) + int32(-120)
	goto L94
L96:
	;
	v387 = v380 - v318 - base.I64_extend_i32_u(base.B2i32(base.Ui64(v376) < base.Ui64(v312)))
	if int64(0) <= v387 {
		goto L107
	} else {
		goto L108
	}
L97:
	;
	v328 = v295
	v331 = v296
	v332 = v322
	goto L98
L98:
	;
	v339 = v332 - v318 - base.I64_extend_i32_u(base.B2i32(base.Ui64(v328) < base.Ui64(v312)))
	if v339 < int64(0) {
		goto L101
	} else {
		goto L102
	}
L99:
	;
	v376 = v368
	v379 = v314
	v380 = v366
	goto L96
L100:
	;
	v368 = v365 << (uint(int64(1)) % 64)
	v370 = v331 + int32(-1)
	if v314 < v370 {
		v328 = v368
		v331 = v370
		v332 = v366
		goto L98
	} else {
		goto L105
	}
L101:
	;
	v365 = v328
	v366 = v332<<(uint(int64(1))%64) | int64(base.Ui64(v328)>>(uint(int64(63))%64))
	goto L100
L102:
	;
	v342 = v328 - v312
	if v339|v342 != int64(0) {
		goto L103
	} else {
		goto L104
	}
L103:
	;
	v365 = v342
	v366 = v339<<(uint(int64(1))%64) | int64(base.Ui64(v342)>>(uint(int64(63))%64))
	goto L100
L104:
	;
	v348 = int64(0)
	F___multf3(m, v15+int32(32), l1, l2, v348, v348)
	mBase = m.M
	v353 = *(*int64)(unsafe.Add(mBase, uint32(v15+int32(40))))
	v354 = *(*int64)(unsafe.Add(mBase, uint32(v15)+32))
	v473 = v353
	v475 = v354
	goto L1
L105:
	;
	goto L99
L106:
	;
	if base.Ui64(int64(281474976710655)) < base.Ui64(v404) {
		v434 = v403
		v437 = v379
		v439 = v404
		goto L110
	} else {
		goto L111
	}
L107:
	;
	v390 = v376 - v312
	if v387|v390 != int64(0) {
		v403 = v390
		v404 = v387
		goto L106
	} else {
		goto L109
	}
L108:
	;
	v403 = v376
	v404 = v380
	goto L106
L109:
	;
	v396 = int64(0)
	F___multf3(m, v15+int32(48), l1, l2, v396, v396)
	mBase = m.M
	v401 = *(*int64)(unsafe.Add(mBase, uint32(v15+int32(56))))
	v402 = *(*int64)(unsafe.Add(mBase, uint32(v15)+48))
	v473 = v401
	v475 = v402
	goto L1
L110:
	;
	v443 = v112 & int32(32768)
	if int32(0) < v437 {
		goto L115
	} else {
		goto L116
	}
L111:
	;
	v411 = v403
	v414 = v379
	v416 = v404
	goto L112
L112:
	;
	v422 = v414 + int32(-1)
	v423 = int64(1)
	v424 = v411 << (uint(v423) % 64)
	v427 = int64(base.Ui64(v411)>>(uint(int64(63))%64)) | v416<<(uint(v423)%64)
	if base.Ui64(v427) < base.Ui64(int64(281474976710656)) {
		v411 = v424
		v414 = v422
		v416 = v427
		goto L112
	} else {
		goto L114
	}
L113:
	;
	v434 = v424
	v437 = v422
	v439 = v427
	goto L110
L114:
	;
	goto L113
L115:
	;
	v473 = v439&int64(281474976710655) | base.I64_extend_i32_u(v437|v443)<<(uint(int64(48))%64)
	v475 = v434
	goto L1
L116:
	;
	F___multf3(m, v15+int32(64), v434, v439&int64(281474976710655)|base.I64_extend_i32_u(v437+int32(120)|v443)<<(uint(int64(48))%64), int64(0), int64(4577627546245398528))
	mBase = m.M
	v462 = *(*int64)(unsafe.Add(mBase, uint32(v15+int32(72))))
	v463 = *(*int64)(unsafe.Add(mBase, uint32(v15)+64))
	v473 = v462
	v475 = v463
	goto L1
}
func F_fmt_o(m *base.Module, l0 int64, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int64
	_ = v6
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v15 int32
	_ = v15
	var v22 int32
	_ = v22
	if l0 == int64(0) {
		v22 = l1
	} else {
		v6 = l0
		v7 = l1
		for {
			v10 = v7 + int32(-1)
			v15 = base.I32_wrap_i64(v6)&int32(7) | int32(48)
			*(*uint8)(unsafe.Add(mBase, uint32(v10))) = uint8(v15)
			if base.Ui64(int64(7)) < base.Ui64(v6) {
				v6 = int64(base.Ui64(v6) >> (uint(int64(3)) % 64))
				v7 = v10
				continue
			} else {
				break
			}
			break
		}
		v22 = v10
	}
	return v22
}
func F_fmt_u(m *base.Module, l0 int64, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int64
	_ = v9
	var v10 int32
	_ = v10
	var v16 int32
	_ = v16
	var v17 int64
	_ = v17
	var v18 int64
	_ = v18
	var v24 int32
	_ = v24
	var v29 int32
	_ = v29
	var v30 int64
	_ = v30
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v51 int32
	_ = v51
	var v56 int32
	_ = v56
	if base.Ui64(int64(4294967296)) <= base.Ui64(l0) {
		v9 = l0
		v10 = l1
		for {
			v16 = v10 + int32(-1)
			v17 = int64(10)
			v18 = base.I64_div_u_s(v9, v17)
			v24 = base.I32_wrap_i64(v9-v18*v17) | int32(48)
			*(*uint8)(unsafe.Add(mBase, uint32(v16))) = uint8(v24)
			if base.Ui64(int64(42949672959)) < base.Ui64(v9) {
				v9 = v18
				v10 = v16
				continue
			} else {
				break
			}
			break
		}
		v29 = v16
		v30 = v18
	} else {
		v29 = l1
		v30 = l0
	}
	if v30 == int64(0) {
		v56 = v29
	} else {
		v38 = v29
		v40 = base.I32_wrap_i64(v30)
		for {
			v44 = v38 + int32(-1)
			v45 = int32(10)
			v46 = base.I32_div_u_s(v40, v45)
			v51 = v40 - v46*v45 | int32(48)
			*(*uint8)(unsafe.Add(mBase, uint32(v44))) = uint8(v51)
			if base.Ui32(int32(9)) < base.Ui32(v40) {
				v38 = v44
				v40 = v46
				continue
			} else {
				break
			}
			break
		}
		v56 = v44
	}
	return v56
}
func F_fourbyte_strstr(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
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
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v58 int32
	_ = v58
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v72 int32
	_ = v72
	v3 = int32(0)
	v8 = l0 + int32(3)
	v9 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+3)))
	if v9 == v3 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	if v67 != 0 {
		goto L9
	} else {
		goto L10
	}
L2:
	;
	v66 = v8
	v67 = base.B2i32(v9 != v3)
	goto L1
L3:
	;
	v14 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+1)))
	v17 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	v18 = int32(24)
	v21 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+2)))
	v22 = int32(8)
	v25 = v14<<(uint(int32(16))%32) | v17<<(uint(v18)%32) | v21<<(uint(v22)%32) | v9
	v26 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v29 = int32(65280)
	v41 = v26<<(uint(v18)%32) | v26&v29<<(uint(v22)%32) | (int32(base.Ui32(v26)>>(uint(v22)%32))&v29 | int32(base.Ui32(v26)>>(uint(v18)%32)))
	if v25 == v41 {
		goto L2
	} else {
		goto L4
	}
L4:
	;
	v45 = v8
	v48 = v25
	goto L5
L5:
	;
	v50 = v45 + int32(1)
	v51 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v45)+1)))
	v52 = int32(0)
	v53 = base.B2i32(v51 != v52)
	if v51 == v52 {
		v66 = v50
		v67 = v53
		goto L1
	} else {
		goto L7
	}
L7:
	;
	v58 = v48<<(uint(int32(8))%32) | v51
	if v58 != v41 {
		v45 = v50
		v48 = v58
		goto L5
	} else {
		goto L8
	}
L8:
	;
	v66 = v50
	v67 = v53
	goto L1
L9:
	;
	v72 = v66 + int32(-3)
	goto L11
L10:
	;
	v72 = int32(0)
	goto L11
L11:
	;
	return v72
}
func F_fp_force_eval_2(m *base.Module, l0 float64) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	v2 = m.G0
	*(*float64)(unsafe.Add(mBase, uint32(v2-int32(16))+8)) = l0
	return
}
func F_fpconv_dtoa(m *base.Module, l0 float64, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v29 int64
	_ = v29
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v43 int64
	_ = v43
	var v44 int64
	_ = v44
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v61 int64
	_ = v61
	var v66 int32
	_ = v66
	var v67 int64
	_ = v67
	var v68 int64
	_ = v68
	var v71 int64
	_ = v71
	var v79 int64
	_ = v79
	var v87 int64
	_ = v87
	var v88 int32
	_ = v88
	var v93 int32
	_ = v93
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v109 float64
	_ = v109
	var v115 int32
	_ = v115
	var v117 int32
	_ = v117
	var v121 int64
	_ = v121
	var v128 int32
	_ = v128
	var v130 int32
	_ = v130
	var v152 int32
	_ = v152
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v168 int64
	_ = v168
	var v169 int64
	_ = v169
	var v170 int64
	_ = v170
	var v172 int64
	_ = v172
	var v174 int64
	_ = v174
	var v176 int64
	_ = v176
	var v177 int64
	_ = v177
	var v178 int64
	_ = v178
	var v182 int64
	_ = v182
	var v183 int64
	_ = v183
	var v184 int64
	_ = v184
	var v203 int64
	_ = v203
	var v204 int64
	_ = v204
	var v206 int64
	_ = v206
	var v207 int64
	_ = v207
	var v208 int64
	_ = v208
	var v209 int64
	_ = v209
	var v214 int64
	_ = v214
	var v215 int64
	_ = v215
	var v216 int64
	_ = v216
	var v222 int64
	_ = v222
	var v229 int64
	_ = v229
	var v234 int64
	_ = v234
	var v236 int64
	_ = v236
	var v239 int64
	_ = v239
	var v241 int64
	_ = v241
	var v242 int64
	_ = v242
	var v248 int64
	_ = v248
	var v249 int64
	_ = v249
	var v269 int64
	_ = v269
	var v274 int64
	_ = v274
	var v275 int64
	_ = v275
	var v277 int64
	_ = v277
	var v278 int64
	_ = v278
	var v280 int32
	_ = v280
	var v288 int32
	_ = v288
	var v292 int32
	_ = v292
	var v297 int64
	_ = v297
	var v305 int32
	_ = v305
	var v308 int32
	_ = v308
	var v314 int64
	_ = v314
	var v315 int64
	_ = v315
	var v317 int64
	_ = v317
	var v320 int32
	_ = v320
	var v327 int32
	_ = v327
	var v331 int32
	_ = v331
	var v333 int32
	_ = v333
	var v335 int64
	_ = v335
	var v337 int64
	_ = v337
	var v343 int32
	_ = v343
	var v358 int64
	_ = v358
	var v359 int32
	_ = v359
	var v362 int32
	_ = v362
	var v364 int32
	_ = v364
	var v367 int64
	_ = v367
	var v368 int32
	_ = v368
	var v370 int32
	_ = v370
	var v375 int64
	_ = v375
	var v377 int32
	_ = v377
	var v378 int32
	_ = v378
	var v384 int32
	_ = v384
	var v388 int32
	_ = v388
	var v390 int32
	_ = v390
	var v392 int64
	_ = v392
	var v393 int64
	_ = v393
	var v394 int32
	_ = v394
	var v395 int32
	_ = v395
	var v398 int64
	_ = v398
	var v399 int64
	_ = v399
	var v401 int32
	_ = v401
	var v414 int64
	_ = v414
	var v425 int64
	_ = v425
	var v430 int32
	_ = v430
	var v432 int32
	_ = v432
	var v447 int32
	_ = v447
	var v450 int32
	_ = v450
	var v458 int32
	_ = v458
	var v462 int32
	_ = v462
	var v465 int64
	_ = v465
	var v469 int32
	_ = v469
	var v482 int64
	_ = v482
	var v493 int64
	_ = v493
	var v498 int32
	_ = v498
	var v500 int32
	_ = v500
	var v513 int32
	_ = v513
	var v516 int32
	_ = v516
	var v528 int32
	_ = v528
	var v529 int32
	_ = v529
	var v535 int32
	_ = v535
	var v543 int32
	_ = v543
	var v544 int32
	_ = v544
	var v548 int32
	_ = v548
	var v557 int32
	_ = v557
	var v565 int32
	_ = v565
	var v569 int32
	_ = v569
	var v575 int32
	_ = v575
	var v576 int32
	_ = v576
	var v577 int32
	_ = v577
	var v578 int32
	_ = v578
	var v583 int32
	_ = v583
	var v584 int32
	_ = v584
	var v587 int32
	_ = v587
	var v591 int32
	_ = v591
	var v597 int32
	_ = v597
	var v599 int32
	_ = v599
	var v603 int32
	_ = v603
	var v610 int32
	_ = v610
	var v613 int32
	_ = v613
	var v617 int32
	_ = v617
	var v618 int32
	_ = v618
	var v619 int32
	_ = v619
	var v621 int32
	_ = v621
	var v628 int32
	_ = v628
	var v631 int32
	_ = v631
	var v638 int32
	_ = v638
	var v640 int32
	_ = v640
	var v643 int32
	_ = v643
	var v646 int32
	_ = v646
	var v649 int32
	_ = v649
	var v650 int32
	_ = v650
	var v654 int32
	_ = v654
	var v656 int32
	_ = v656
	var v664 int32
	_ = v664
	var v668 int32
	_ = v668
	var v669 int32
	_ = v669
	var v674 int32
	_ = v674
	var v676 int32
	_ = v676
	var v681 int32
	_ = v681
	v24 = m.G0
	v26 = v24 - int32(32)
	m.G0 = v26
	v29 = base.I64_reinterpret_f64(l0)
	if int64(-1) < v29 {
		v35 = int32(0)
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v36 = l1 + v35
	if base.F64_ne(l0, float64(0)) != 0 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	v32 = int32(45)
	*(*uint8)(unsafe.Add(mBase, uint32(l1))) = uint8(v32)
	v35 = int32(1)
	goto L1
L3:
	;
	m.G0 = v26 + int32(32)
	return v681 + v35
L4:
	;
	v43 = v29 & int64(4503599627370495)
	v44 = int64(9218868437227405312)
	if v29&v44 != v44 {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	v39 = int32(48)
	*(*uint8)(unsafe.Add(mBase, uint32(v36))) = uint8(v39)
	v681 = int32(1)
	goto L3
L6:
	;
	v61 = v43 | int64(4503599627370496)
	v66 = base.I32_wrap_i64(int64(base.Ui64(v29)>>(uint(int64(52))%64))) & int32(2047)
	if v66 != 0 {
		goto L10
	} else {
		goto L11
	}
L7:
	;
	if v43 == int64(0) {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v55 = int32(102)
	*(*uint8)(unsafe.Add(mBase, uint32(v36)+2)) = uint8(v55)
	v57 = int32(28265)
	*(*uint16)(unsafe.Add(mBase, uint32(v36))) = uint16(v57)
	v681 = int32(3)
	goto L3
L9:
	;
	v50 = int32(110)
	*(*uint8)(unsafe.Add(mBase, uint32(v36)+2)) = uint8(v50)
	v52 = int32(24942)
	*(*uint16)(unsafe.Add(mBase, uint32(v36))) = uint16(v52)
	v681 = int32(3)
	goto L3
L10:
	;
	v67 = v61
	goto L12
L11:
	;
	v67 = v43
	goto L12
L12:
	;
	v68 = int64(1)
	v71 = v67<<(uint(v68)%64) | v68
	if v66 == int32(0) {
		goto L14
	} else {
		goto L15
	}
L13:
	;
	if v66 != 0 {
		goto L16
	} else {
		goto L17
	}
L14:
	;
	v79 = base.I64_clz(v71 & int64(9007199254740991))
	v87 = v71 << (uint(int64(53)-(v79^int64(63))) % 64)
	v88 = int32(-1064) - base.I32_wrap_i64(v79)
	goto L13
L15:
	;
	v87 = v71
	v88 = v66 + int32(-1075)
	goto L13
L16:
	;
	v93 = v66 + int32(-1064)
	goto L18
L17:
	;
	v93 = int32(-1063)
	goto L18
L18:
	;
	v97 = base.B2i32(v67 == int64(4503599627370496))
	if v67 == int64(4503599627370496) {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	v98 = int32(-2)
	goto L21
L20:
	;
	v98 = int32(-1)
	goto L21
L21:
	;
	v109 = base.F64_mul(base.F64_convert_i32_s(int32(-76)-v88), float64(0.30102999566398114))
	if base.F64_lt(base.F64_abs(v109), float64(2.147483648e+09)) == int32(0) {
		goto L23
	} else {
		goto L24
	}
L22:
	;
	if v67 == int64(4503599627370496) {
		goto L25
	} else {
		goto L26
	}
L23:
	;
	v117 = int32(-2147483648)
	goto L22
L24:
	;
	v115 = base.I32_trunc_f64_s(v109)
	v117 = v115
	goto L22
L25:
	;
	v121 = int64(2)
	goto L27
L26:
	;
	v121 = int64(1)
	goto L27
L27:
	;
	v128 = base.I32_div_s(v117+int32(348), int32(8))
	v130 = v128
	goto L28
L28:
	;
	v152 = m.G3
	v157 = v152 + int32(_a1919) + v130<<(uint(int32(4))%32)
	v158 = *(*int32)(unsafe.Add(mBase, uint32(v157)+8))
	v159 = v88 + int32(-11) + v158
	if int32(-124) <= v159 {
		goto L30
	} else {
		goto L31
	}
L29:
	;
	v168 = *(*int64)(unsafe.Add(mBase, uint32(v157)))
	v169 = int64(32)
	v170 = int64(base.Ui64(v168) >> (uint(v169) % 64))
	v172 = v87 << (uint(int64(10)) % 64)
	v174 = int64(base.Ui64(v172) >> (uint(v169) % 64))
	v176 = int64(4294967295)
	v177 = v168 & v176
	v178 = v177 * v174
	v182 = int64(4294966272)
	v183 = v172 & v182
	v184 = v170 * v183
	v203 = v170*v174 + int64(base.Ui64(v178)>>(uint(v169)%64)) + int64(base.Ui64(v184)>>(uint(v169)%64)) + int64(base.Ui64(v184&v182+v178&v176+int64(base.Ui64(v177*v183)>>(uint(v169)%64))+int64(2147483648))>>(uint(v169)%64)) + int64(-1)
	if v66 != 0 {
		goto L34
	} else {
		goto L35
	}
L30:
	;
	if v159 < int32(-95) {
		goto L32
	} else {
		goto L33
	}
L31:
	;
	v130 = int32(1) + v130
	goto L28
L32:
	;
	goto L29
L33:
	;
	v130 = int32(-1) + v130
	goto L28
L34:
	;
	v204 = v61
	goto L36
L35:
	;
	v204 = v43 << (uint(int64(52)-(base.I64_clz(v43)^int64(63))) % 64)
	goto L36
L36:
	;
	v206 = v204 << (uint(int64(11)) % 64)
	v207 = int64(32)
	v208 = int64(base.Ui64(v206) >> (uint(v207) % 64))
	v209 = v177 * v208
	v214 = int64(4294965248)
	v215 = v206 & v214
	v216 = v170 * v215
	v222 = int64(4294967295)
	v229 = int64(2147483648)
	v234 = v203 - (int64(base.Ui64(v209)>>(uint(v207)%64)) + v170*v208 + int64(base.Ui64(v216)>>(uint(v207)%64)) + int64(base.Ui64(v216&v214+v209&v222+int64(base.Ui64(v177*v215)>>(uint(v207)%64))+v229)>>(uint(v207)%64)))
	v236 = int64(-1)
	v239 = (v67<<(uint(v121)%64) + v236) << (uint(base.I64_extend_i32_u(v93+v98-v88)) % 64)
	v241 = int64(base.Ui64(v239) >> (uint(v207) % 64))
	v242 = v177 * v241
	v248 = v239 & v222
	v249 = v170 * v248
	v269 = v203 + (int64(base.Ui64(v242)>>(uint(v207)%64)) + v170*v241 + int64(base.Ui64(v249)>>(uint(v207)%64)) + int64(base.Ui64(v249&v222+v242&v222+int64(base.Ui64(v177*v248)>>(uint(v207)%64))+v229)>>(uint(v207)%64)) ^ v236)
	v274 = base.I64_extend_i32_u(int32(-53) - (v88 + v158))
	v275 = int64(1) << (uint(v274) % 64)
	v277 = v275 + v236
	v278 = v203 & v277
	v280 = m.G3
	v288 = int32(348) - v130<<(uint(int32(3))%32)
	v292 = int32(10)
	v297 = int64(base.Ui64(v203) >> (uint(v274) % 64))
	v305 = int32(0)
	v308 = v280 + int32(_a1920)
	goto L39
L37:
	;
	v528 = int32(1)
	v529 = v516 + v513
	if v529 < v528 {
		goto L81
	} else {
		goto L82
	}
L38:
	;
	v462 = v333 + v288
	if base.Ui64(v337) < base.Ui64(v234) {
		goto L69
	} else {
		goto L70
	}
L39:
	;
	v314 = *(*int64)(unsafe.Add(mBase, uint32(v308)))
	v315 = base.I64_div_u_s(v297, v314)
	v317 = v315 & int64(4294967295)
	if v317 != int64(0) {
		goto L42
	} else {
		goto L43
	}
L40:
	;
	v343 = m.G3
	v358 = v269
	v359 = v288
	v362 = v88
	v364 = v331
	v367 = v278
	v368 = int32(0)
	v370 = v343 + int32(_a1921)
	goto L47
L41:
	;
	v333 = v292 + int32(-1)
	v335 = v297 - v317*v314
	v337 = v335<<(uint(v274)%64) + v278
	if base.Ui64(v337) <= base.Ui64(v269) {
		goto L38
	} else {
		goto L45
	}
L42:
	;
	v327 = base.I32_wrap_i64(v315) + int32(48)
	*(*uint8)(unsafe.Add(mBase, uint32(v26+v305))) = uint8(v327)
	v331 = v305 + int32(1)
	goto L41
L43:
	;
	v320 = int32(0)
	if v305 == v320 {
		v331 = v320
		goto L41
	} else {
		goto L44
	}
L44:
	;
	goto L42
L45:
	;
	if base.Ui32(int32(1)) < base.Ui32(v292) {
		v292 = v333
		v297 = v335
		v305 = v331
		v308 = v308 + int32(8)
		goto L39
	} else {
		goto L46
	}
L46:
	;
	goto L40
L47:
	;
	v375 = v367 * int64(10)
	v377 = base.I32_wrap_i64(int64(base.Ui64(v375) >> (uint(v274) % 64)))
	if v377 != 0 {
		goto L50
	} else {
		goto L51
	}
L49:
	;
	v390 = v368 + int32(-1)
	v392 = v358 * int64(10)
	v393 = v375 & v277
	v394 = base.B2i32(base.Ui64(v392) <= base.Ui64(v393))
	if base.Ui64(v392) <= base.Ui64(v393) {
		goto L54
	} else {
		goto L55
	}
L50:
	;
	v384 = v377 + int32(48)
	*(*uint8)(unsafe.Add(mBase, uint32(v26+v364))) = uint8(v384)
	v388 = v364 + int32(1)
	goto L49
L51:
	;
	v378 = int32(0)
	if v364 == v378 {
		v388 = v378
		goto L49
	} else {
		goto L52
	}
L52:
	;
	goto L50
L53:
	;
	if base.Ui64(v392) <= base.Ui64(v393) {
		v358 = v392
		v359 = v447
		v362 = v450
		v364 = v388
		v367 = v393
		v368 = v390
		v370 = v458
		goto L47
	} else {
		goto L68
	}
L54:
	;
	v447 = v359
	v450 = v362
	v458 = v370 + int32(-8)
	goto L53
L55:
	;
	v395 = v390 + v359
	if base.Ui64(v275) <= base.Ui64(v392-v393) {
		goto L56
	} else {
		goto L57
	}
L56:
	;
	v398 = *(*int64)(unsafe.Add(mBase, uint32(v370)))
	v399 = v398 * v234
	if base.Ui64(v393) < base.Ui64(v399) {
		goto L58
	} else {
		goto L59
	}
L57:
	;
	v513 = v395
	v516 = v388
	goto L37
L58:
	;
	v401 = v26 + int32(-1) + v388
	v414 = v393
	goto L60
L59:
	;
	v513 = v395
	v516 = v388
	goto L37
L60:
	;
	v425 = v414 + v275
	if base.Ui64(v425) < base.Ui64(v399) {
		goto L62
	} else {
		goto L63
	}
L61:
	;
	v447 = v395
	v450 = v388
	v458 = v370
	goto L53
L62:
	;
	v430 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v401))))
	v432 = v430 + int32(-1)
	*(*uint8)(unsafe.Add(mBase, uint32(v401))) = uint8(v432)
	if base.Ui64(v425) < base.Ui64(v399) {
		goto L65
	} else {
		goto L66
	}
L63:
	;
	if base.Ui64(v425-v399) < base.Ui64(v399-v414) {
		goto L62
	} else {
		goto L64
	}
L64:
	;
	v447 = v395
	v450 = v388
	v458 = v370
	goto L53
L65:
	;
	if base.Ui64(v275) <= base.Ui64(v392-v425) {
		v414 = v425
		goto L60
	} else {
		goto L67
	}
L66:
	;
	v447 = v395
	v450 = v388
	v458 = v370
	goto L53
L67:
	;
	goto L61
L68:
	;
	v513 = v447
	v516 = v450
	goto L37
L69:
	;
	v465 = v314 << (uint(v274) % 64)
	if base.Ui64(v465) <= base.Ui64(v269-v337) {
		goto L71
	} else {
		goto L72
	}
L70:
	;
	v513 = v462
	v516 = v331
	goto L37
L71:
	;
	v469 = v26 + v331 + int32(-1)
	v482 = v337
	goto L73
L72:
	;
	v513 = v462
	v516 = v331
	goto L37
L73:
	;
	v493 = v482 + v465
	if base.Ui64(v493) < base.Ui64(v234) {
		goto L75
	} else {
		goto L76
	}
L74:
	;
	v513 = v462
	v516 = v331
	goto L37
L75:
	;
	v498 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v469))))
	v500 = v498 + int32(-1)
	*(*uint8)(unsafe.Add(mBase, uint32(v469))) = uint8(v500)
	if base.Ui64(v493) < base.Ui64(v234) {
		goto L78
	} else {
		goto L79
	}
L76:
	;
	if base.Ui64(v493-v234) < base.Ui64(v234-v482) {
		goto L75
	} else {
		goto L77
	}
L77:
	;
	v513 = v462
	v516 = v331
	goto L37
L78:
	;
	if base.Ui64(v465) <= base.Ui64(v269-v493) {
		v482 = v493
		goto L73
	} else {
		goto L80
	}
L79:
	;
	v513 = v462
	v516 = v331
	goto L37
L80:
	;
	goto L74
L81:
	;
	v535 = v528 - v529
	goto L83
L82:
	;
	v535 = v529 + int32(-1)
	goto L83
L83:
	;
	if v513 < int32(0) {
		goto L84
	} else {
		goto L85
	}
L84:
	;
	if int32(-1) < v513 {
		goto L91
	} else {
		goto L92
	}
L85:
	;
	if v516+int32(7) <= v535 {
		goto L84
	} else {
		goto L86
	}
L86:
	;
	if v516 == int32(0) {
		v544 = v36
		goto L88
	} else {
		goto L89
	}
L87:
	;
	v548 = F__emscripten_memset_bulkmem(m, v544+v516, base.I32_extend8_s(int32(48)), v513)
	mBase = m.M
	goto L90
L88:
	;
	goto L87
L89:
	;
	v543 = F__emscripten_memcpy_bulkmem(m, v36, v26, v516)
	mBase = m.M
	v544 = v543
	goto L88
L90:
	;
	v681 = v529
	goto L3
L91:
	;
	v591 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v26))))
	*(*uint8)(unsafe.Add(mBase, uint32(v36))) = uint8(v591)
	if v29 < int64(0) {
		goto L108
	} else {
		goto L109
	}
L92:
	;
	if base.Ui32(int32(-7)) < base.Ui32(v513) {
		goto L93
	} else {
		goto L94
	}
L93:
	;
	if int32(0) < v529 {
		goto L96
	} else {
		goto L97
	}
L94:
	;
	if int32(3) < v535 {
		goto L91
	} else {
		goto L95
	}
L95:
	;
	goto L93
L96:
	;
	if v529 == int32(0) {
		v576 = v36
		goto L103
	} else {
		goto L104
	}
L97:
	;
	v557 = int32(11824)
	*(*uint16)(unsafe.Add(mBase, uint32(v36))) = uint16(v557)
	v565 = F__emscripten_memset_bulkmem(m, v36+int32(2), base.I32_extend8_s(int32(48)), int32(0)-v529)
	mBase = m.M
	goto L98
L98:
	;
	if v516 == int32(0) {
		goto L100
	} else {
		goto L101
	}
L99:
	;
	v681 = int32(2) - v513
	goto L3
L100:
	;
	goto L99
L101:
	;
	v569 = F__emscripten_memcpy_bulkmem(m, v565-v529, v26, v516)
	mBase = m.M
	goto L100
L102:
	;
	v577 = v576 + v529
	v578 = int32(46)
	*(*uint8)(unsafe.Add(mBase, uint32(v577))) = uint8(v578)
	v583 = int32(0)
	v584 = v583 - v513
	if v584 == v583 {
		goto L106
	} else {
		goto L107
	}
L103:
	;
	goto L102
L104:
	;
	v575 = F__emscripten_memcpy_bulkmem(m, v36, v26, v529)
	mBase = m.M
	v576 = v575
	goto L103
L105:
	;
	v681 = v516 + int32(1)
	goto L3
L106:
	;
	goto L105
L107:
	;
	v587 = F__emscripten_memcpy_bulkmem(m, v577+int32(1), v26+v529, v584)
	mBase = m.M
	goto L106
L108:
	;
	v597 = int32(17)
	goto L110
L109:
	;
	v597 = int32(18)
	goto L110
L110:
	;
	if v516 < v597 {
		goto L111
	} else {
		goto L112
	}
L111:
	;
	v599 = v516
	goto L113
L112:
	;
	v599 = v597
	goto L113
L113:
	;
	if v516 < int32(2) {
		v617 = int32(1)
		goto L114
	} else {
		goto L115
	}
L114:
	;
	v618 = v36 + v617
	v619 = int32(101)
	*(*uint8)(unsafe.Add(mBase, uint32(v618))) = uint8(v619)
	v621 = int32(1)
	if v599+v513 < v621 {
		goto L119
	} else {
		goto L120
	}
L115:
	;
	v603 = int32(46)
	*(*uint8)(unsafe.Add(mBase, uint32(v36)+1)) = uint8(v603)
	v610 = v599 + int32(-1)
	if v610 == int32(0) {
		goto L117
	} else {
		goto L118
	}
L116:
	;
	v617 = v599 + int32(1)
	goto L114
L117:
	;
	goto L116
L118:
	;
	v613 = F__emscripten_memcpy_bulkmem(m, v36+int32(2), v26|int32(1), v610)
	mBase = m.M
	goto L117
L119:
	;
	v628 = int32(45)
	goto L121
L120:
	;
	v628 = int32(43)
	goto L121
L121:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v618+v621))) = uint8(v628)
	v631 = v617 + int32(2)
	if int32(99) < v535 {
		goto L125
	} else {
		goto L126
	}
L122:
	;
	v674 = base.I32_rem_s(v669, int32(10))
	v676 = v674 + int32(48)
	*(*uint8)(unsafe.Add(mBase, uint32(v36+v668))) = uint8(v676)
	v681 = v668 + int32(1)
	goto L3
L123:
	;
	v664 = int32(48)
	*(*uint8)(unsafe.Add(mBase, uint32(v36+v643))) = uint8(v664)
	v668 = v617 + int32(4)
	v669 = v646
	goto L122
L124:
	;
	v654 = base.I32_div_u_s(v650, int32(10))
	v656 = v654 + int32(48)
	*(*uint8)(unsafe.Add(mBase, uint32(v36+v649))) = uint8(v656)
	v668 = v649 + int32(1)
	v669 = v654*int32(-10) + v650
	goto L122
L125:
	;
	v638 = base.I32_div_u_s(v535, int32(100))
	v640 = v638 + int32(48)
	*(*uint8)(unsafe.Add(mBase, uint32(v36+v631))) = uint8(v640)
	v643 = v617 + int32(3)
	v646 = v638*int32(-100) + v535
	if v646 < int32(10) {
		goto L123
	} else {
		goto L128
	}
L126:
	;
	if int32(9) < v535 {
		v649 = v631
		v650 = v535
		goto L124
	} else {
		goto L127
	}
L127:
	;
	v668 = v631
	v669 = v535
	goto L122
L128:
	;
	v649 = v643
	v650 = v646
	goto L124
}
func F_fpconv_strtod(m *base.Module, l0 int32, l1 int32) float64 {
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
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	var v35 int32
	_ = v35
	var v56 int32
	_ = v56
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v79 int32
	_ = v79
	var v85 float64
	_ = v85
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v122 int32
	_ = v122
	var v124 int32
	_ = v124
	var v126 int32
	_ = v126
	var v128 int32
	_ = v128
	var v131 int32
	_ = v131
	var v133 int32
	_ = v133
	var v135 int32
	_ = v135
	var v142 int32
	_ = v142
	var v144 int32
	_ = v144
	var v147 int32
	_ = v147
	var v152 int32
	_ = v152
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v163 int32
	_ = v163
	var __phi163 int32
	_ = __phi163
	var v164 int32
	_ = v164
	var __phi164 int32
	_ = __phi164
	var v170 int32
	_ = v170
	var v173 int32
	_ = v173
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v195 int32
	_ = v195
	var v204 int32
	_ = v204
	var v206 int32
	_ = v206
	var v209 int32
	_ = v209
	var v212 int32
	_ = v212
	var v214 int32
	_ = v214
	var v219 int32
	_ = v219
	var v226 int32
	_ = v226
	var v231 int32
	_ = v231
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v246 int32
	_ = v246
	var v254 int32
	_ = v254
	var v256 int32
	_ = v256
	var v260 int32
	_ = v260
	var v261 int32
	_ = v261
	var v267 int32
	_ = v267
	var v269 int32
	_ = v269
	var v276 int32
	_ = v276
	var v278 int32
	_ = v278
	var v282 int32
	_ = v282
	var v283 int32
	_ = v283
	var v292 int32
	_ = v292
	var v293 int32
	_ = v293
	var v296 int32
	_ = v296
	var v298 int32
	_ = v298
	var v300 int32
	_ = v300
	var v309 int32
	_ = v309
	var v311 int32
	_ = v311
	var v314 int32
	_ = v314
	var v319 int32
	_ = v319
	var v324 int32
	_ = v324
	var v325 int32
	_ = v325
	var v330 int32
	_ = v330
	var __phi330 int32
	_ = __phi330
	var v331 int32
	_ = v331
	var __phi331 int32
	_ = __phi331
	var v337 int32
	_ = v337
	var v340 int32
	_ = v340
	var v354 int32
	_ = v354
	var v363 int32
	_ = v363
	var v365 int32
	_ = v365
	var v368 int32
	_ = v368
	var v371 int32
	_ = v371
	var v373 int32
	_ = v373
	var v378 int32
	_ = v378
	var v385 int32
	_ = v385
	var v390 int32
	_ = v390
	var v412 int32
	_ = v412
	var v424 int32
	_ = v424
	var v426 int32
	_ = v426
	var v428 int32
	_ = v428
	var v432 int32
	_ = v432
	var v437 int32
	_ = v437
	var v438 int32
	_ = v438
	var v449 int32
	_ = v449
	var v452 int32
	_ = v452
	var v459 int32
	_ = v459
	var v464 int32
	_ = v464
	var v468 int32
	_ = v468
	var v470 int32
	_ = v470
	var v485 int32
	_ = v485
	var v487 int32
	_ = v487
	var v490 int32
	_ = v490
	var v493 int32
	_ = v493
	var v497 int32
	_ = v497
	var v507 int32
	_ = v507
	var v508 int32
	_ = v508
	var v514 int32
	_ = v514
	var v516 int32
	_ = v516
	var v519 int32
	_ = v519
	var v523 int32
	_ = v523
	var v529 int32
	_ = v529
	var v531 int32
	_ = v531
	var v532 int32
	_ = v532
	var v533 int32
	_ = v533
	var v534 int32
	_ = v534
	var v543 int32
	_ = v543
	var v545 int32
	_ = v545
	var v546 int32
	_ = v546
	var v547 int32
	_ = v547
	var v549 int32
	_ = v549
	var v569 float64
	_ = v569
	var v570 int32
	_ = v570
	var v571 int32
	_ = v571
	var v576 int32
	_ = v576
	var v577 int32
	_ = v577
	var v578 int32
	_ = v578
	var v579 int32
	_ = v579
	var v588 float64
	_ = v588
	v9 = m.G0
	v11 = v9 - int32(48)
	m.G0 = v11
	v13 = m.G3
	v16 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+uint32(_consts[1018]))))
	if v16 == int32(46) {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	m.G0 = v11 + int32(48)
	return v588
L2:
	;
	v571 = m.G3
	v576 = m.G397
	v577 = *(*int32)(unsafe.Add(mBase, uint32(v576)))
	v578 = F_fwrite(m, v571+int32(_a1898), int32(13), int32(1), v577)
	mBase = m.M
	v579 = m.ExcPending
	if v579 != 0 {
		goto L28
	} else {
		goto L132
	}
L3:
	;
	v569 = F_strtod(m, l0, l1)
	mBase = m.M
	v570 = m.ExcPending
	if v570 != 0 {
		goto L28
	} else {
		goto L131
	}
L4:
	;
	v23 = l0
	goto L6
L5:
	;
	if v23 != l0 {
		goto L14
	} else {
		goto L15
	}
L6:
	;
	v27 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23))))
	if base.Ui32((v27+int32(-48))&int32(255)) < base.Ui32(int32(10)) {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v23 = v23 + int32(1)
	goto L6
L9:
	;
	v35 = v27 + int32(-43)
	if base.Ui32(int32(3)) < base.Ui32(v35) {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	if base.Ui32((v27|int32(32)+int32(-122))&int32(255)) < base.Ui32(int32(231)) {
		goto L5
	} else {
		goto L13
	}
L11:
	;
	if v35 != int32(1) {
		goto L8
	} else {
		goto L12
	}
L12:
	;
	goto L10
L13:
	;
	goto L8
L14:
	;
	v56 = v23 - l0
	if v56 < int32(32) {
		v64 = v11 + int32(16)
		goto L16
	} else {
		goto L17
	}
L15:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = l0
	v588 = float64(0)
	goto L1
L16:
	;
	if v56 == int32(0) {
		v68 = v64
		goto L20
	} else {
		goto L21
	}
L17:
	;
	v61 = F_emscripten_builtin_malloc(m, v56+int32(1))
	mBase = m.M
	if v61 == int32(0) {
		goto L2
	} else {
		goto L18
	}
L18:
	;
	v64 = v61
	goto L16
L19:
	;
	v70 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v68+v56))) = uint8(v70)
	v72 = int32(46)
	v73 = F___strchrnul(m, v68, v72)
	mBase = m.M
	v75 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v73))))
	if v75 == v72 {
		goto L24
	} else {
		goto L25
	}
L20:
	;
	goto L19
L21:
	;
	v67 = F__emscripten_memcpy_bulkmem(m, v64, l0, v56)
	mBase = m.M
	v68 = v67
	goto L20
L22:
	;
	v85 = F_strtod(m, v68, v11+int32(12))
	mBase = m.M
	v88 = m.ExcPending
	if v88 != 0 {
		goto L28
	} else {
		goto L29
	}
L23:
	;
	if v79 == int32(0) {
		goto L22
	} else {
		goto L27
	}
L24:
	;
	v79 = v73
	goto L26
L25:
	;
	v79 = v70
	goto L26
L26:
	;
	goto L23
L27:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v79))) = uint8(v16)
	goto L22
L28:
	;
	return float64(0)
L29:
	;
	v89 = *(*int32)(unsafe.Add(mBase, uint32(v11)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = l0 + (v89 - v68)
	if v56 < int32(32) {
		v588 = v85
		goto L1
	} else {
		goto L30
	}
L30:
	;
	if v68 == int32(0) {
		goto L32
	} else {
		goto L33
	}
L31:
	;
	v588 = v85
	goto L1
L32:
	;
	goto L31
L33:
	;
	v105 = int32(-8)
	v106 = v68 + v105
	v109 = *(*int32)(unsafe.Add(mBase, uint32(v68+int32(-4))))
	v111 = v109 & v105
	v112 = v106 + v111
	if v109&int32(1) != 0 {
		v236 = v111
		v237 = v106
		goto L34
	} else {
		goto L35
	}
L34:
	;
	if base.Ui32(v112) <= base.Ui32(v237) {
		goto L32
	} else {
		goto L69
	}
L35:
	;
	if v109&int32(2) == int32(0) {
		goto L32
	} else {
		goto L36
	}
L36:
	;
	v119 = *(*int32)(unsafe.Add(mBase, uint32(v106)))
	v120 = v106 - v119
	v122 = *(*int32)(unsafe.Add(mBase, _consts[515]))
	if base.Ui32(v120) < base.Ui32(v122) {
		goto L32
	} else {
		goto L37
	}
L37:
	;
	v124 = v119 + v111
	v126 = *(*int32)(unsafe.Add(mBase, _consts[516]))
	if v120 == v126 {
		goto L41
	} else {
		goto L42
	}
L38:
	;
	if v142 == int32(0) {
		v236 = v124
		v237 = v120
		goto L34
	} else {
		goto L57
	}
L39:
	;
	v195 = int32(0)
	goto L38
L40:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v131)+12)) = v128
	*(*int32)(unsafe.Add(mBase, uint32(v128)+8)) = v131
	v236 = v124
	v237 = v120
	goto L34
L41:
	;
	v176 = *(*int32)(unsafe.Add(mBase, uint32(v112)+4))
	v177 = int32(3)
	if v176&v177 != v177 {
		v236 = v124
		v237 = v120
		goto L34
	} else {
		goto L56
	}
L42:
	;
	v128 = *(*int32)(unsafe.Add(mBase, uint32(v120)+12))
	if base.Ui32(int32(255)) < base.Ui32(v119) {
		goto L43
	} else {
		goto L44
	}
L43:
	;
	v142 = *(*int32)(unsafe.Add(mBase, uint32(v120)+24))
	if v128 == v120 {
		goto L46
	} else {
		goto L47
	}
L44:
	;
	v131 = *(*int32)(unsafe.Add(mBase, uint32(v120)+8))
	if v128 != v131 {
		goto L40
	} else {
		goto L45
	}
L45:
	;
	v133 = int32(0)
	v135 = *(*int32)(unsafe.Add(mBase, _consts[517]))
	*(*int32)(unsafe.Add(mBase, _consts[517])) = v135 & base.I32_rotl(int32(-2), int32(base.Ui32(v119)>>(uint(int32(3))%32)))
	v236 = v124
	v237 = v120
	goto L34
L46:
	;
	v147 = *(*int32)(unsafe.Add(mBase, uint32(v120)+20))
	if v147 == int32(0) {
		goto L49
	} else {
		goto L50
	}
L47:
	;
	v144 = *(*int32)(unsafe.Add(mBase, uint32(v120)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v144)+12)) = v128
	*(*int32)(unsafe.Add(mBase, uint32(v128)+8)) = v144
	v195 = v128
	goto L38
L48:
	;
	__phi163 = v157
	__phi164 = v158
	v163 = __phi163
	v164 = __phi164
	goto L52
L49:
	;
	v152 = *(*int32)(unsafe.Add(mBase, uint32(v120)+16))
	if v152 == int32(0) {
		goto L39
	} else {
		goto L51
	}
L50:
	;
	v157 = v147
	v158 = v120 + int32(20)
	goto L48
L51:
	;
	v157 = v152
	v158 = v120 + int32(16)
	goto L48
L52:
	;
	v170 = *(*int32)(unsafe.Add(mBase, uint32(v163)+20))
	if v170 != 0 {
		__phi163 = v170
		__phi164 = v163 + int32(20)
		v163 = __phi163
		v164 = __phi164
		goto L52
	} else {
		goto L54
	}
L53:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v164))) = int32(0)
	v195 = v163
	goto L38
L54:
	;
	v173 = *(*int32)(unsafe.Add(mBase, uint32(v163)+16))
	if v173 != 0 {
		__phi163 = v173
		__phi164 = v163 + int32(16)
		v163 = __phi163
		v164 = __phi164
		goto L52
	} else {
		goto L55
	}
L55:
	;
	goto L53
L56:
	;
	*(*int32)(unsafe.Add(mBase, _consts[518])) = v124
	*(*int32)(unsafe.Add(mBase, uint32(v112)+4)) = v176 & int32(-2)
	*(*int32)(unsafe.Add(mBase, uint32(v120)+4)) = v124 | int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v112))) = v124
	goto L31
L57:
	;
	v204 = *(*int32)(unsafe.Add(mBase, uint32(v120)+28))
	v206 = v204 << (uint(int32(2)) % 32)
	v209 = *(*int32)(unsafe.Add(mBase, uint32(v206)+uint32(_consts[519])))
	if v120 != v209 {
		goto L59
	} else {
		goto L60
	}
L58:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v195)+24)) = v142
	v226 = *(*int32)(unsafe.Add(mBase, uint32(v120)+16))
	if v226 == int32(0) {
		goto L66
	} else {
		goto L67
	}
L59:
	;
	v219 = *(*int32)(unsafe.Add(mBase, uint32(v142)+16))
	if v219 != v120 {
		goto L63
	} else {
		goto L64
	}
L60:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v206)+uint32(_consts[519]))) = v195
	if v195 != 0 {
		goto L58
	} else {
		goto L61
	}
L61:
	;
	v212 = int32(0)
	v214 = *(*int32)(unsafe.Add(mBase, _consts[520]))
	*(*int32)(unsafe.Add(mBase, _consts[520])) = v214 & base.I32_rotl(int32(-2), v204)
	v236 = v124
	v237 = v120
	goto L34
L62:
	;
	if v195 == int32(0) {
		v236 = v124
		v237 = v120
		goto L34
	} else {
		goto L65
	}
L63:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v142)+20)) = v195
	goto L62
L64:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v142)+16)) = v195
	goto L62
L65:
	;
	goto L58
L66:
	;
	v231 = *(*int32)(unsafe.Add(mBase, uint32(v120)+20))
	if v231 == int32(0) {
		v236 = v124
		v237 = v120
		goto L34
	} else {
		goto L68
	}
L67:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v195)+16)) = v226
	*(*int32)(unsafe.Add(mBase, uint32(v226)+24)) = v195
	goto L66
L68:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v195)+20)) = v231
	*(*int32)(unsafe.Add(mBase, uint32(v231)+24)) = v195
	v236 = v124
	v237 = v120
	goto L34
L69:
	;
	v246 = *(*int32)(unsafe.Add(mBase, uint32(v112)+4))
	if v246&int32(1) == int32(0) {
		goto L32
	} else {
		goto L70
	}
L70:
	;
	if v246&int32(2) != 0 {
		goto L75
	} else {
		goto L76
	}
L71:
	;
	if base.Ui32(int32(255)) < base.Ui32(v412) {
		goto L109
	} else {
		goto L110
	}
L72:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v237)+4)) = v292 | int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v237+v292))) = v292
	if v237 != v276 {
		v412 = v292
		goto L71
	} else {
		goto L108
	}
L73:
	;
	if v309 == int32(0) {
		goto L72
	} else {
		goto L96
	}
L74:
	;
	v354 = int32(0)
	goto L73
L75:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v112)+4)) = v246 & int32(-2)
	*(*int32)(unsafe.Add(mBase, uint32(v237)+4)) = v236 | int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v237+v236))) = v236
	v412 = v236
	goto L71
L76:
	;
	v254 = *(*int32)(unsafe.Add(mBase, _consts[521]))
	if v112 != v254 {
		goto L77
	} else {
		goto L78
	}
L77:
	;
	v276 = *(*int32)(unsafe.Add(mBase, _consts[516]))
	if v112 != v276 {
		goto L80
	} else {
		goto L81
	}
L78:
	;
	v256 = int32(0)
	*(*int32)(unsafe.Add(mBase, _consts[521])) = v237
	v260 = *(*int32)(unsafe.Add(mBase, _consts[522]))
	v261 = v260 + v236
	*(*int32)(unsafe.Add(mBase, _consts[522])) = v261
	*(*int32)(unsafe.Add(mBase, uint32(v237)+4)) = v261 | int32(1)
	v267 = *(*int32)(unsafe.Add(mBase, _consts[516]))
	if v237 != v267 {
		goto L32
	} else {
		goto L79
	}
L79:
	;
	v269 = int32(0)
	*(*int32)(unsafe.Add(mBase, _consts[518])) = v269
	*(*int32)(unsafe.Add(mBase, _consts[516])) = v269
	goto L31
L80:
	;
	v292 = v246&int32(-8) + v236
	v293 = *(*int32)(unsafe.Add(mBase, uint32(v112)+12))
	if base.Ui32(int32(255)) < base.Ui32(v246) {
		goto L82
	} else {
		goto L83
	}
L81:
	;
	v278 = int32(0)
	*(*int32)(unsafe.Add(mBase, _consts[516])) = v237
	v282 = *(*int32)(unsafe.Add(mBase, _consts[518]))
	v283 = v282 + v236
	*(*int32)(unsafe.Add(mBase, _consts[518])) = v283
	*(*int32)(unsafe.Add(mBase, uint32(v237)+4)) = v283 | int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v237+v283))) = v283
	goto L31
L82:
	;
	v309 = *(*int32)(unsafe.Add(mBase, uint32(v112)+24))
	if v293 == v112 {
		goto L86
	} else {
		goto L87
	}
L83:
	;
	v296 = *(*int32)(unsafe.Add(mBase, uint32(v112)+8))
	if v293 != v296 {
		goto L84
	} else {
		goto L85
	}
L84:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v296)+12)) = v293
	*(*int32)(unsafe.Add(mBase, uint32(v293)+8)) = v296
	goto L72
L85:
	;
	v298 = int32(0)
	v300 = *(*int32)(unsafe.Add(mBase, _consts[517]))
	*(*int32)(unsafe.Add(mBase, _consts[517])) = v300 & base.I32_rotl(int32(-2), int32(base.Ui32(v246)>>(uint(int32(3))%32)))
	goto L72
L86:
	;
	v314 = *(*int32)(unsafe.Add(mBase, uint32(v112)+20))
	if v314 == int32(0) {
		goto L89
	} else {
		goto L90
	}
L87:
	;
	v311 = *(*int32)(unsafe.Add(mBase, uint32(v112)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v311)+12)) = v293
	*(*int32)(unsafe.Add(mBase, uint32(v293)+8)) = v311
	v354 = v293
	goto L73
L88:
	;
	__phi330 = v324
	__phi331 = v325
	v330 = __phi330
	v331 = __phi331
	goto L92
L89:
	;
	v319 = *(*int32)(unsafe.Add(mBase, uint32(v112)+16))
	if v319 == int32(0) {
		goto L74
	} else {
		goto L91
	}
L90:
	;
	v324 = v314
	v325 = v112 + int32(20)
	goto L88
L91:
	;
	v324 = v319
	v325 = v112 + int32(16)
	goto L88
L92:
	;
	v337 = *(*int32)(unsafe.Add(mBase, uint32(v330)+20))
	if v337 != 0 {
		__phi330 = v337
		__phi331 = v330 + int32(20)
		v330 = __phi330
		v331 = __phi331
		goto L92
	} else {
		goto L94
	}
L93:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v331))) = int32(0)
	v354 = v330
	goto L73
L94:
	;
	v340 = *(*int32)(unsafe.Add(mBase, uint32(v330)+16))
	if v340 != 0 {
		__phi330 = v340
		__phi331 = v330 + int32(16)
		v330 = __phi330
		v331 = __phi331
		goto L92
	} else {
		goto L95
	}
L95:
	;
	goto L93
L96:
	;
	v363 = *(*int32)(unsafe.Add(mBase, uint32(v112)+28))
	v365 = v363 << (uint(int32(2)) % 32)
	v368 = *(*int32)(unsafe.Add(mBase, uint32(v365)+uint32(_consts[519])))
	if v112 != v368 {
		goto L98
	} else {
		goto L99
	}
L97:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v354)+24)) = v309
	v385 = *(*int32)(unsafe.Add(mBase, uint32(v112)+16))
	if v385 == int32(0) {
		goto L105
	} else {
		goto L106
	}
L98:
	;
	v378 = *(*int32)(unsafe.Add(mBase, uint32(v309)+16))
	if v378 != v112 {
		goto L102
	} else {
		goto L103
	}
L99:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v365)+uint32(_consts[519]))) = v354
	if v354 != 0 {
		goto L97
	} else {
		goto L100
	}
L100:
	;
	v371 = int32(0)
	v373 = *(*int32)(unsafe.Add(mBase, _consts[520]))
	*(*int32)(unsafe.Add(mBase, _consts[520])) = v373 & base.I32_rotl(int32(-2), v363)
	goto L72
L101:
	;
	if v354 == int32(0) {
		goto L72
	} else {
		goto L104
	}
L102:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v309)+20)) = v354
	goto L101
L103:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v309)+16)) = v354
	goto L101
L104:
	;
	goto L97
L105:
	;
	v390 = *(*int32)(unsafe.Add(mBase, uint32(v112)+20))
	if v390 == int32(0) {
		goto L72
	} else {
		goto L107
	}
L106:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v354)+16)) = v385
	*(*int32)(unsafe.Add(mBase, uint32(v385)+24)) = v354
	goto L105
L107:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v354)+20)) = v390
	*(*int32)(unsafe.Add(mBase, uint32(v390)+24)) = v354
	goto L72
L108:
	;
	*(*int32)(unsafe.Add(mBase, _consts[518])) = v292
	goto L31
L109:
	;
	if base.Ui32(int32(16777215)) < base.Ui32(v412) {
		v459 = int32(31)
		goto L114
	} else {
		goto L115
	}
L110:
	;
	v424 = v412 & int32(-8)
	v426 = v424 + int32(9128464)
	v428 = *(*int32)(unsafe.Add(mBase, _consts[517]))
	v432 = int32(1) << (uint(int32(base.Ui32(v412)>>(uint(int32(3))%32))) % 32)
	if v428&v432 != 0 {
		goto L112
	} else {
		goto L113
	}
L111:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v424)+uint32(_consts[523]))) = v237
	*(*int32)(unsafe.Add(mBase, uint32(v438)+12)) = v237
	*(*int32)(unsafe.Add(mBase, uint32(v237)+12)) = v426
	*(*int32)(unsafe.Add(mBase, uint32(v237)+8)) = v438
	goto L31
L112:
	;
	v437 = *(*int32)(unsafe.Add(mBase, uint32(v424)+uint32(_consts[523])))
	v438 = v437
	goto L111
L113:
	;
	*(*int32)(unsafe.Add(mBase, _consts[517])) = v428 | v432
	v438 = v426
	goto L111
L114:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v237)+28)) = v459
	*(*int64)(unsafe.Add(mBase, uint32(v237)+16)) = int64(0)
	v464 = v459 << (uint(int32(2)) % 32)
	v468 = *(*int32)(unsafe.Add(mBase, _consts[520]))
	v470 = int32(1) << (uint(v459) % 32)
	if v468&v470 != 0 {
		goto L119
	} else {
		goto L120
	}
L115:
	;
	v449 = base.I32_clz(int32(base.Ui32(v412) >> (uint(int32(8)) % 32)))
	v452 = int32(1)
	v459 = int32(base.Ui32(v412)>>(uint(int32(38)-v449)%32))&v452 - v449<<(uint(v452)%32) + int32(62)
	goto L114
L116:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v237+v531))) = v534
	*(*int32)(unsafe.Add(mBase, uint32(v237)+12)) = v533
	*(*int32)(unsafe.Add(mBase, uint32(v237+v529))) = v532
	v543 = int32(0)
	v545 = *(*int32)(unsafe.Add(mBase, _consts[524]))
	v546 = int32(-1)
	v547 = v545 + v546
	if v547 != 0 {
		goto L128
	} else {
		goto L129
	}
L117:
	;
	v523 = *(*int32)(unsafe.Add(mBase, uint32(v493)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v523)+12)) = v237
	*(*int32)(unsafe.Add(mBase, uint32(v493)+8)) = v237
	v529 = int32(24)
	v531 = int32(8)
	v532 = int32(0)
	v533 = v493
	v534 = v523
	goto L116
L118:
	;
	v529 = v514
	v531 = v516
	v532 = v237
	v533 = v237
	v534 = v519
	goto L116
L119:
	;
	if v459 == int32(31) {
		goto L121
	} else {
		goto L122
	}
L120:
	;
	*(*int32)(unsafe.Add(mBase, _consts[520])) = v468 | v470
	*(*int32)(unsafe.Add(mBase, uint32(v464)+uint32(_consts[519]))) = v237
	v514 = int32(8)
	v516 = int32(24)
	v519 = v464 + int32(9128728)
	goto L118
L121:
	;
	v485 = int32(0)
	goto L123
L122:
	;
	v485 = int32(25) - int32(base.Ui32(v459)>>(uint(int32(1))%32))
	goto L123
L123:
	;
	v487 = *(*int32)(unsafe.Add(mBase, uint32(v464)+uint32(_consts[519])))
	v490 = v412 << (uint(v485) % 32)
	v493 = v487
	goto L124
L124:
	;
	v497 = *(*int32)(unsafe.Add(mBase, uint32(v493)+4))
	if v497&int32(-8) == v412 {
		goto L117
	} else {
		goto L126
	}
L125:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v507+int32(16)))) = v237
	v514 = int32(8)
	v516 = int32(24)
	v519 = v493
	goto L118
L126:
	;
	v507 = v493 + int32(base.Ui32(v490)>>(uint(int32(29))%32))&int32(4)
	v508 = *(*int32)(unsafe.Add(mBase, uint32(v507)+16))
	if v508 != 0 {
		v490 = v490 << (uint(int32(1)) % 32)
		v493 = v508
		goto L124
	} else {
		goto L127
	}
L127:
	;
	goto L125
L128:
	;
	v549 = v547
	goto L130
L129:
	;
	v549 = v546
	goto L130
L130:
	;
	*(*int32)(unsafe.Add(mBase, _consts[524])) = v549
	goto L32
L131:
	;
	v588 = v569
	goto L1
L132:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_fputs(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v18 int32
	_ = v18
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v29 int32
	_ = v29
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v58 int32
	_ = v58
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	if l0&int32(3) == int32(0) {
		v25 = l0
		goto L3
	} else {
		goto L4
	}
L1:
	;
	v62 = F_fwrite(m, l0, int32(1), v58, l1)
	mBase = m.M
	v65 = m.ExcPending
	if v65 != 0 {
		goto L17
	} else {
		goto L18
	}
L2:
	;
	v58 = v50 - l0
	goto L1
L3:
	;
	v29 = v25
	goto L11
L4:
	;
	v11 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v11 != 0 {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	v14 = l0
	goto L7
L6:
	;
	v58 = l0 - l0
	goto L1
L7:
	;
	v18 = v14 + int32(1)
	if v18&int32(3) == int32(0) {
		v25 = v18
		goto L3
	} else {
		goto L9
	}
L9:
	;
	v23 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18))))
	if v23 != 0 {
		v14 = v18
		goto L7
	} else {
		goto L10
	}
L10:
	;
	v50 = v18
	goto L2
L11:
	;
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v29)))
	v38 = int32(-2139062144)
	if (int32(16843008)-v35|v35)&v38 == v38 {
		v29 = v29 + int32(4)
		goto L11
	} else {
		goto L13
	}
L12:
	;
	v44 = v29
	goto L14
L13:
	;
	goto L12
L14:
	;
	v48 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v44))))
	if v48 != 0 {
		v44 = v44 + int32(1)
		goto L14
	} else {
		goto L16
	}
L15:
	;
	v50 = v44
	goto L2
L16:
	;
	goto L15
L17:
	;
	return int32(0)
L18:
	;
	if v58 != v62 {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	v67 = int32(-1)
	goto L21
L20:
	;
	v67 = int32(0)
	goto L21
L21:
	;
	return v67
}
func F_freeEvalScriptsAsync(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
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
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v12 = v10 + v11
	if base.Ui32(v12) < base.Ui32(int32(65)) {
		F_freeEvalScripts(m, l0, l1, l2)
		mBase = m.M
		v28 = m.ExcPending
		if v28 != 0 {
			return
		} else {
			m.G0 = v8 + int32(16)
			return
		}
	} else {
		v15 = int32(0)
		v17 = *(*int32)(unsafe.Add(mBase, _consts[503]))
		*(*int32)(unsafe.Add(mBase, _consts[503])) = v17 + v12
		*(*int32)(unsafe.Add(mBase, uint32(v8))) = l0
		*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = l1
		*(*int32)(unsafe.Add(mBase, uint32(v8)+8)) = l2
		F_bioCreateLazyFreeJob(m, int32(555), int32(3), v8)
		mBase = m.M
		v26 = m.ExcPending
		if v26 != 0 {
			return
		} else {
			m.G0 = v8 + int32(16)
			return
		}
	}
}
func F_freeMemoryGetNotCountedMemory(m *base.Module) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int64
	_ = v6
	var v8 int32
	_ = v8
	var v13 int64
	_ = v13
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v29 int32
	_ = v29
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v45 int32
	_ = v45
	var v52 int32
	_ = v52
	var v56 int32
	_ = v56
	var v60 int32
	_ = v60
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
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v81 int32
	_ = v81
	var v85 int32
	_ = v85
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v95 int32
	_ = v95
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v111 int32
	_ = v111
	var v113 int32
	_ = v113
	var v120 int32
	_ = v120
	var v123 int32
	_ = v123
	var v125 int32
	_ = v125
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v131 int32
	_ = v131
	var v136 int32
	_ = v136
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v147 int32
	_ = v147
	var v149 int32
	_ = v149
	var v153 int32
	_ = v153
	var v155 int32
	_ = v155
	var v161 int32
	_ = v161
	v5 = int32(_a44)
	v6 = *(*int64)(unsafe.Add(mBase, _consts[287]))
	v8 = *(*int32)(unsafe.Add(mBase, _consts[288]))
	if base.I64_extend_i32_u(v8) <= v6 {
		v23 = int32(0)
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v25 = *(*int32)(unsafe.Add(mBase, _consts[28]))
	if v25 == int32(0) {
		v71 = v23
		goto L6
	} else {
		goto L7
	}
L2:
	;
	v13 = base.I64_div_s(v6, int64(16384))
	v20 = v8 - base.I32_wrap_i64(v6+v13*int64(44)) + int32(-44)
	if base.Ui32(v8) < base.Ui32(v20) {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v22 = int32(0)
	goto L5
L4:
	;
	v22 = v20
	goto L5
L5:
	;
	v23 = v22
	goto L1
L6:
	;
	v72 = int32(0)
	v75 = m.G0
	v77 = v75 - int32(16)
	m.G0 = v77
	v81 = *(*int32)(unsafe.Add(mBase, _consts[139]))
	if v81 == v72 {
		v113 = v72
		goto L23
	} else {
		goto L24
	}
L7:
	;
	v29 = *(*int32)(unsafe.Add(mBase, _consts[34]))
	v36 = v29 + int32(-1)
	v37 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v36))))
	v39 = v37 & int32(7)
	switch v39 {
	case 0:
		goto L14
	case 1:
		v45 = int32(4)
		goto L9
	case 2:
		goto L13
	case 3:
		goto L12
	case 4:
		goto L11
	default:
		goto L10
	}
L8:
	;
	v71 = v69 + v23
	goto L6
L9:
	;
	switch v39 {
	case 0:
		goto L20
	case 1:
		goto L19
	case 2:
		goto L18
	case 3:
		goto L17
	case 4:
		goto L16
	default:
		v65 = int32(0)
		goto L15
	}
L10:
	;
	v45 = int32(1)
	goto L9
L11:
	;
	v45 = int32(18)
	goto L9
L12:
	;
	v45 = int32(10)
	goto L9
L13:
	;
	v45 = int32(6)
	goto L9
L14:
	;
	v40 = F_zmalloc_usable_size(m, v36)
	mBase = m.M
	v69 = v40
	goto L8
L15:
	;
	v69 = v45 + v65
	goto L8
L16:
	;
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v29+int32(-9))))
	v65 = v64
	goto L15
L17:
	;
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v29+int32(-5))))
	v69 = v45 + v60
	goto L8
L18:
	;
	v56 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v29+int32(-3)))))
	v69 = v45 + v56
	goto L8
L19:
	;
	v52 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v29+int32(-2)))))
	v69 = v45 + v52
	goto L8
L20:
	;
	v69 = v45 + int32(base.Ui32(v37)>>(uint(int32(3))%32))
	goto L8
L21:
	;
	return v161
L22:
	;
	if v113 == int32(0) {
		v161 = v71
		goto L21
	} else {
		goto L33
	}
L23:
	;
	m.G0 = v77 + int32(16)
	goto L22
L24:
	;
	v85 = *(*int32)(unsafe.Add(mBase, _consts[136]))
	if v85 == int32(0) {
		v113 = v72
		goto L23
	} else {
		goto L25
	}
L25:
	;
	v88 = *(*int32)(unsafe.Add(mBase, uint32(v85)+uint32(_consts[212])))
	v90 = v77 + int32(8)
	F_listRewind(m, v88, v90)
	mBase = m.M
	v92 = int32(0)
	v95 = F_listNext(m, v90)
	mBase = m.M
	if v95 == v92 {
		v113 = v92
		goto L23
	} else {
		goto L26
	}
L26:
	;
	v100 = v95
	goto L27
L27:
	;
	v101 = *(*int32)(unsafe.Add(mBase, uint32(v100)+8))
	v102 = *(*int32)(unsafe.Add(mBase, uint32(v101)))
	if v102 != 0 {
		goto L29
	} else {
		goto L30
	}
L28:
	;
	v113 = v92
	goto L23
L29:
	;
	v111 = F_listNext(m, v77+int32(8))
	mBase = m.M
	if v111 != 0 {
		v100 = v111
		goto L27
	} else {
		goto L32
	}
L30:
	;
	v103 = *(*int32)(unsafe.Add(mBase, uint32(v101)+156))
	if base.Ui32(v103+int32(-18)) <= base.Ui32(int32(2)) {
		goto L29
	} else {
		goto L31
	}
L31:
	;
	v113 = int32(1)
	goto L23
L32:
	;
	goto L28
L33:
	;
	v120 = int32(0)
	v123 = m.G0
	v125 = v123 - int32(16)
	m.G0 = v125
	v128 = *(*int32)(unsafe.Add(mBase, _consts[136]))
	v129 = *(*int32)(unsafe.Add(mBase, uint32(v128)+uint32(_consts[212])))
	v131 = v125 + int32(8)
	F_listRewind(m, v129, v131)
	mBase = m.M
	v136 = F_listNext(m, v131)
	mBase = m.M
	if v136 == v120 {
		v155 = v120
		goto L35
	} else {
		goto L36
	}
L34:
	;
	v161 = v155 + v71
	goto L21
L35:
	;
	m.G0 = v125 + int32(16)
	goto L34
L36:
	;
	v140 = v120
	v141 = v136
	goto L37
L37:
	;
	v142 = *(*int32)(unsafe.Add(mBase, uint32(v141)+8))
	v143 = *(*int32)(unsafe.Add(mBase, uint32(v142)))
	if v143 != 0 {
		v149 = v140
		goto L39
	} else {
		goto L40
	}
L38:
	;
	v155 = v149
	goto L35
L39:
	;
	v153 = F_listNext(m, v125+int32(8))
	mBase = m.M
	if v153 != 0 {
		v140 = v149
		v141 = v153
		goto L37
	} else {
		goto L42
	}
L40:
	;
	v144 = *(*int32)(unsafe.Add(mBase, uint32(v142)+152))
	if v144 == int32(0) {
		v149 = v140
		goto L39
	} else {
		goto L41
	}
L41:
	;
	v147 = F_getClientOutputBufferMemoryUsage(m, v144)
	mBase = m.M
	v149 = v147 + v140
	goto L39
L42:
	;
	goto L38
}
func F_freeMemoryOverheadData(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+124))
	F_valkey_free(m, v2)
	mBase = m.M
	v4 = m.ExcPending
	if v4 != 0 {
		return
	} else {
		F_valkey_free(m, l0)
		mBase = m.M
		v6 = m.ExcPending
		if v6 != 0 {
			return
		} else {
			return
		}
	}
}
func F_freeSparklineSequence(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v3 < int32(1) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	F_valkey_free(m, v22)
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		goto L5
	} else {
		goto L8
	}
L2:
	;
	v8 = int32(0)
	goto L3
L3:
	;
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v13 = *(*int32)(unsafe.Add(mBase, uint32(v9+v8<<(uint(int32(4))%32))+8))
	F_valkey_free(m, v13)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
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
	v17 = v8 + int32(1)
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v17 < v18 {
		v8 = v17
		goto L3
	} else {
		goto L7
	}
L7:
	;
	goto L4
L8:
	;
	F_valkey_free(m, l0)
	mBase = m.M
	v26 = m.ExcPending
	if v26 != 0 {
		goto L5
	} else {
		goto L9
	}
L9:
	;
	return
}
func F_frexp(m *base.Module, l0 float64, l1 int32) float64 {
	mBase := m.M
	_ = mBase
	var v5 int64
	_ = v5
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v17 float64
	_ = v17
	var v20 int64
	_ = v20
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v33 float64
	_ = v33
	var v34 int32
	_ = v34
	var v37 float64
	_ = v37
	var v38 int32
	_ = v38
	var v48 float64
	_ = v48
	var v51 float64
	_ = v51
	var v52 int32
	_ = v52
	var v55 float64
	_ = v55
	var v56 int32
	_ = v56
	var v67 float64
	_ = v67
	v5 = base.I64_reinterpret_f64(l0)
	v9 = int32(2047)
	v10 = base.I32_wrap_i64(int64(base.Ui64(v5)>>(uint(int64(52))%64))) & v9
	if v10 == v9 {
		v67 = l0
		return v67
	} else {
		if v10 != 0 {
			*(*int32)(unsafe.Add(mBase, uint32(l1))) = v10 + int32(-1022)
			v67 = base.F64_reinterpret_i64(v5&int64(-9218868437227405313) | int64(4602678819172646912))
			return v67
		} else {
			if base.F64_ne(l0, float64(0)) != 0 {
				v17 = base.F64_mul(l0, float64(1.8446744073709552e+19))
				v20 = base.I64_reinterpret_f64(v17)
				v24 = int32(2047)
				v25 = base.I32_wrap_i64(int64(base.Ui64(v20)>>(uint(int64(52))%64))) & v24
				if v25 == v24 {
					v48 = v17
					v51 = v48
				} else {
					if v25 != 0 {
						*(*int32)(unsafe.Add(mBase, uint32(l1))) = v25 + int32(-1022)
						v48 = base.F64_reinterpret_i64(v20&int64(-9218868437227405313) | int64(4602678819172646912))
						v51 = v48
					} else {
						if base.F64_ne(v17, float64(0)) != 0 {
							v33 = F_frexp(m, base.F64_mul(v17, float64(1.8446744073709552e+19)), l1)
							mBase = m.M
							v34 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
							v37 = v33
							v38 = v34 + int32(-64)
						} else {
							v37 = v17
							v38 = int32(0)
						}
						*(*int32)(unsafe.Add(mBase, uint32(l1))) = v38
						v51 = v37
					}
				}
				v52 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
				v55 = v51
				v56 = v52 + int32(-64)
			} else {
				v55 = l0
				v56 = int32(0)
			}
			*(*int32)(unsafe.Add(mBase, uint32(l1))) = v56
			return v55
		}
	}
}
func F_fsyncFileDir(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v23 int32
	_ = v23
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v34 int32
	_ = v34
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v49 int32
	_ = v49
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v63 int32
	_ = v63
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v94 int32
	_ = v94
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	var v105 int32
	_ = v105
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v117 int32
	_ = v117
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v133 int32
	_ = v133
	var v139 int32
	_ = v139
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v152 int32
	_ = v152
	var v157 int32
	_ = v157
	var v159 int32
	_ = v159
	v5 = m.G0
	v7 = v5 - int32(4112)
	m.G0 = v7
	if l0&int32(3) == int32(0) {
		v30 = l0
		goto L5
	} else {
		goto L6
	}
L1:
	;
	m.G0 = v7 + int32(4112)
	return v159
L2:
	;
	v71 = v63 + int32(1)
	if v71 == int32(0) {
		v75 = v7
		goto L23
	} else {
		goto L24
	}
L3:
	;
	if base.Ui32(v63) < base.Ui32(int32(4097)) {
		goto L2
	} else {
		goto L19
	}
L4:
	;
	v63 = v55 - l0
	goto L3
L5:
	;
	v34 = v30
	goto L13
L6:
	;
	v16 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v16 != 0 {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	v19 = l0
	goto L9
L8:
	;
	v63 = l0 - l0
	goto L3
L9:
	;
	v23 = v19 + int32(1)
	if v23&int32(3) == int32(0) {
		v30 = v23
		goto L5
	} else {
		goto L11
	}
L11:
	;
	v28 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23))))
	if v28 != 0 {
		v19 = v23
		goto L9
	} else {
		goto L12
	}
L12:
	;
	v55 = v23
	goto L4
L13:
	;
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v34)))
	v43 = int32(-2139062144)
	if (int32(16843008)-v40|v40)&v43 == v43 {
		v34 = v34 + int32(4)
		goto L13
	} else {
		goto L15
	}
L14:
	;
	v49 = v34
	goto L16
L15:
	;
	goto L14
L16:
	;
	v53 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v49))))
	if v53 != 0 {
		v49 = v49 + int32(1)
		goto L16
	} else {
		goto L18
	}
L17:
	;
	v55 = v49
	goto L4
L18:
	;
	goto L17
L19:
	;
	goto L20
L20:
	;
	*(*int32)(unsafe.Add(mBase, _consts[5])) = int32(37)
	v159 = int32(-1)
	goto L1
L21:
	;
	v143 = F_fsync(m, v133)
	mBase = m.M
	if v143 != int32(-1) {
		goto L49
	} else {
		goto L50
	}
L22:
	;
	if v75 == int32(0) {
		goto L26
	} else {
		goto L27
	}
L23:
	;
	goto L22
L24:
	;
	v74 = F__emscripten_memcpy_bulkmem(m, v7, l0, v71)
	mBase = m.M
	v75 = v74
	goto L23
L25:
	;
	v131 = int32(0)
	v133 = F_open(m, v130, v131, v131)
	mBase = m.M
	if v133 != int32(-1) {
		goto L21
	} else {
		goto L44
	}
L26:
	;
	v130 = int32(_a955)
	goto L25
L27:
	;
	v80 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v75))))
	if v80 == int32(0) {
		goto L26
	} else {
		goto L28
	}
L28:
	;
	v83 = F_strlen(m, v75)
	mBase = m.M
	v85 = v83
	goto L31
L29:
	;
	v130 = int32(_a1760)
	goto L25
L30:
	;
	v105 = v99
	goto L40
L31:
	;
	v88 = v85 + int32(-1)
	v90 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v75+v88))))
	if v90 == int32(47) {
		goto L33
	} else {
		goto L34
	}
L33:
	;
	if v88 != 0 {
		v85 = v88
		goto L31
	} else {
		goto L39
	}
L34:
	;
	v94 = v88
	goto L35
L35:
	;
	if v94 == int32(0) {
		goto L26
	} else {
		goto L37
	}
L36:
	;
	goto L30
L37:
	;
	v99 = v94 + int32(-1)
	v101 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v75+v99))))
	if v101 != int32(47) {
		v94 = v99
		goto L35
	} else {
		goto L38
	}
L38:
	;
	goto L36
L39:
	;
	goto L29
L40:
	;
	if v105 == int32(0) {
		goto L29
	} else {
		goto L42
	}
L41:
	;
	v117 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v111+int32(1)))) = uint8(v117)
	v130 = v75
	goto L25
L42:
	;
	v110 = v105 + int32(-1)
	v111 = v75 + v110
	v112 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v111))))
	if v112 == int32(47) {
		v105 = v110
		goto L40
	} else {
		goto L43
	}
L43:
	;
	goto L41
L44:
	;
	goto L45
L45:
	;
	v139 = *(*int32)(unsafe.Add(mBase, _consts[5]))
	if v139 != int32(31) {
		goto L46
	} else {
		goto L47
	}
L46:
	;
	v142 = int32(-1)
	goto L48
L47:
	;
	v142 = int32(0)
	goto L48
L48:
	;
	v159 = v142
	goto L1
L49:
	;
	v157 = F_close(m, v133)
	mBase = m.M
	v159 = int32(0)
	goto L1
L50:
	;
	v146 = int32(9116376)
	goto L51
L51:
	;
	v147 = *(*int32)(unsafe.Add(mBase, _consts[5]))
	if v147 == int32(8) {
		goto L49
	} else {
		goto L52
	}
L52:
	;
	if v147 == int32(28) {
		goto L49
	} else {
		goto L53
	}
L53:
	;
	v152 = F_close(m, v133)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, _consts[5])) = v147
	v159 = int32(-1)
	goto L1
}
func F_funcargs(m *base.Module, l0 int32, l1 int32) {
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
	var v14 int32
	_ = v14
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v56 int32
	_ = v56
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v76 int32
	_ = v76
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v92 int32
	_ = v92
	var v96 int32
	_ = v96
	var v104 int32
	_ = v104
	var v106 int32
	_ = v106
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v142 int32
	_ = v142
	var v144 int32
	_ = v144
	var v149 int32
	_ = v149
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	v8 = m.G0
	v10 = v8 - int32(32)
	m.G0 = v10
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v14 == int32(286) {
		goto L5
	} else {
		goto L6
	}
L1:
	;
	m.G0 = v10 + int32(32)
	return
L2:
	;
	v153 = F_luaK_codeABC(m, v12, int32(28), v144, v149, int32(2))
	mBase = m.M
	v154 = m.ExcPending
	if v154 != 0 {
		goto L14
	} else {
		goto L38
	}
L3:
	;
	v142 = *(*int32)(unsafe.Add(mBase, uint32(v12)+36))
	v144 = v135
	v149 = v142 - v135
	goto L2
L4:
	;
	F_luaK_exp2nextreg(m, v12, v10+int32(8))
	mBase = m.M
	v134 = m.ExcPending
	if v134 != 0 {
		goto L14
	} else {
		goto L37
	}
L5:
	;
	v113 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v114 = F_luaK_stringK(m, v12, v113)
	mBase = m.M
	v115 = m.ExcPending
	if v115 != 0 {
		goto L14
	} else {
		goto L35
	}
L6:
	;
	if v14 == int32(123) {
		goto L8
	} else {
		goto L9
	}
L7:
	;
	v104 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v106 = *(*int32)(unsafe.Add(mBase, uint32(v10)+8))
	if base.Ui32(v106+int32(-13)) < base.Ui32(int32(2)) {
		v144 = v104
		v149 = int32(0)
		goto L2
	} else {
		goto L33
	}
L8:
	;
	F_constructor(m, l0, v10+int32(8))
	mBase = m.M
	v96 = m.ExcPending
	if v96 != 0 {
		goto L14
	} else {
		goto L32
	}
L9:
	;
	if v14 != int32(40) {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	v88 = m.G3
	F_luaX_syntaxerror(m, l0, v88+int32(_a2053))
	mBase = m.M
	v92 = m.ExcPending
	if v92 != 0 {
		goto L14
	} else {
		goto L31
	}
L11:
	;
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v13 == v21 {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	F_luaX_next(m, l0)
	mBase = m.M
	v29 = m.ExcPending
	if v29 != 0 {
		goto L14
	} else {
		goto L16
	}
L13:
	;
	v23 = m.G3
	F_luaX_syntaxerror(m, l0, v23+int32(_a2054))
	mBase = m.M
	v27 = m.ExcPending
	if v27 != 0 {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	return
L15:
	;
	goto L12
L16:
	;
	v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v30 != int32(41) {
		goto L18
	} else {
		goto L19
	}
L17:
	;
	F_check_match(m, l0, int32(41), int32(40), v13)
	mBase = m.M
	v87 = m.ExcPending
	if v87 != 0 {
		goto L14
	} else {
		goto L30
	}
L18:
	;
	v38 = F_subexpr(m, l0, v10+int32(8), int32(0))
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L14
	} else {
		goto L20
	}
L19:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+8)) = int32(0)
	goto L17
L20:
	;
	v40 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v40 != int32(44) {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	F_luaK_setreturns(m, v12, v10+int32(8), int32(-1))
	mBase = m.M
	v76 = m.ExcPending
	if v76 != 0 {
		goto L14
	} else {
		goto L29
	}
L22:
	;
	goto L23
L23:
	;
	F_luaX_next(m, l0)
	mBase = m.M
	v51 = m.ExcPending
	if v51 != 0 {
		goto L14
	} else {
		goto L25
	}
L24:
	;
	goto L21
L25:
	;
	v52 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	F_luaK_exp2nextreg(m, v52, v10+int32(8))
	mBase = m.M
	v56 = m.ExcPending
	if v56 != 0 {
		goto L14
	} else {
		goto L26
	}
L26:
	;
	v60 = F_subexpr(m, l0, v10+int32(8), int32(0))
	mBase = m.M
	v61 = m.ExcPending
	if v61 != 0 {
		goto L14
	} else {
		goto L27
	}
L27:
	;
	v62 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v62 == int32(44) {
		goto L23
	} else {
		goto L28
	}
L28:
	;
	goto L24
L29:
	;
	goto L17
L30:
	;
	goto L7
L31:
	;
	goto L1
L32:
	;
	goto L7
L33:
	;
	if v106 == int32(0) {
		v135 = v104
		goto L3
	} else {
		goto L34
	}
L34:
	;
	v124 = v104
	goto L4
L35:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = v114
	*(*int32)(unsafe.Add(mBase, uint32(v10)+8)) = int32(4)
	*(*int64)(unsafe.Add(mBase, uint32(v10)+24)) = int64(-1)
	F_luaX_next(m, l0)
	mBase = m.M
	v122 = m.ExcPending
	if v122 != 0 {
		goto L14
	} else {
		goto L36
	}
L36:
	;
	v123 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v124 = v123
	goto L4
L37:
	;
	v135 = v124
	goto L3
L38:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+8)) = v153
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = int32(13)
	*(*int64)(unsafe.Add(mBase, uint32(l1)+16)) = int64(-1)
	v160 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
	v161 = *(*int32)(unsafe.Add(mBase, uint32(v160)+20))
	v162 = *(*int32)(unsafe.Add(mBase, uint32(v12)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v161+v162<<(uint(int32(2))%32)+int32(-4)))) = v13
	goto L39
L39:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+36)) = v144 + int32(1)
	goto L1
}
