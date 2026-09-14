package p1

import (
	base "github.com/shibukawa/valkeymem/internal/aot/vkaot/base"
	"unsafe"
)

func F___toread(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v23 int32
	_ = v23
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+72))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+72)) = v4 + int32(-1) | v4
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v9 == v10 {
		v19 = int32(0)
		*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v19
		*(*int64)(unsafe.Add(mBase, uint32(l0)+16)) = int64(0)
		v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		if v23&int32(4) == v19 {
			v33 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
			v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
			v35 = v33 + v34
			*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v35
			*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v35
			return v23 << (uint(int32(27)) % 32) >> (uint(int32(31)) % 32)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(l0))) = v23 | int32(32)
			return int32(-1)
		}
	} else {
		v12 = int32(0)
		v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
		v15 = m.T0[v14].(func(*base.Module, int32, int32, int32) int32)(m, l0, v12, v12)
		mBase = m.M
		v18 = m.ExcPending
		if v18 != 0 {
			return int32(0)
		} else {
			v19 = int32(0)
			*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v19
			*(*int64)(unsafe.Add(mBase, uint32(l0)+16)) = int64(0)
			v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			if v23&int32(4) == v19 {
				v33 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
				v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
				v35 = v33 + v34
				*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v35
				*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v35
				return v23 << (uint(int32(27)) % 32) >> (uint(int32(31)) % 32)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(l0))) = v23 | int32(32)
				return int32(-1)
			}
		}
	}
}
func F___towrite(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v8 int32
	_ = v8
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+72))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+72)) = v3 + int32(-1) | v3
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v8&int32(8) == int32(0) {
		*(*int64)(unsafe.Add(mBase, uint32(l0)+4)) = int64(0)
		v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
		*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v20
		*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v20
		v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
		*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v20 + v23
		return int32(0)
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(l0))) = v8 | int32(32)
		return int32(-1)
	}
}
func F_tconcat(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v32 int32
	_ = v32
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	var v86 int32
	_ = v86
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v97 int32
	_ = v97
	var v108 int32
	_ = v108
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v125 int32
	_ = v125
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v170 int64
	_ = v170
	var v172 int32
	_ = v172
	var v174 int32
	_ = v174
	var v192 int32
	_ = v192
	var v195 int32
	_ = v195
	var v229 int32
	_ = v229
	var v232 int32
	_ = v232
	var v239 int32
	_ = v239
	var v254 int32
	_ = v254
	var v257 int32
	_ = v257
	var v291 int32
	_ = v291
	var v294 int32
	_ = v294
	var v297 int32
	_ = v297
	var v299 int32
	_ = v299
	var v304 int32
	_ = v304
	var v308 int32
	_ = v308
	var v309 int32
	_ = v309
	var v312 int32
	_ = v312
	var v317 int32
	_ = v317
	var v318 int32
	_ = v318
	var v323 int32
	_ = v323
	var v326 int32
	_ = v326
	var v328 int32
	_ = v328
	var v330 int32
	_ = v330
	var v344 int32
	_ = v344
	var v345 int32
	_ = v345
	var v350 int32
	_ = v350
	var v351 int32
	_ = v351
	var v353 int32
	_ = v353
	var v395 int32
	_ = v395
	var v396 int32
	_ = v396
	var v397 int32
	_ = v397
	var v398 int64
	_ = v398
	var v400 int32
	_ = v400
	var v402 int32
	_ = v402
	var v420 int32
	_ = v420
	var v423 int32
	_ = v423
	var v457 int32
	_ = v457
	var v460 int32
	_ = v460
	var v467 int32
	_ = v467
	var v482 int32
	_ = v482
	var v485 int32
	_ = v485
	var v519 int32
	_ = v519
	var v522 int32
	_ = v522
	var v525 int32
	_ = v525
	var v527 int32
	_ = v527
	var v532 int32
	_ = v532
	var v536 int32
	_ = v536
	var v537 int32
	_ = v537
	var v540 int32
	_ = v540
	var v543 int32
	_ = v543
	var v544 int32
	_ = v544
	var v549 int32
	_ = v549
	var v559 int32
	_ = v559
	v7 = m.G0
	v9 = v7 - int32(1072)
	m.G0 = v9
	v12 = m.G3
	v17 = F_luaL_optlstring(m, l0, int32(2), v12+int32(_a320), v9+int32(32))
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	F_luaL_checktype(m, l0, int32(1), int32(5))
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v27 = F_luaL_optinteger(m, l0, int32(3), int32(1))
	mBase = m.M
	v28 = m.ExcPending
	if v28 != 0 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	goto L10
L5:
	;
	v97 = v9 + int32(36)
	*(*int32)(unsafe.Add(mBase, uint32(v97)+4)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v97)+8)) = l0
	*(*int32)(unsafe.Add(mBase, uint32(v97))) = v9 + int32(48)
	goto L25
L6:
	;
	v93 = F_luaL_checkinteger(m, l0, int32(4))
	mBase = m.M
	v94 = m.ExcPending
	if v94 != 0 {
		goto L1
	} else {
		goto L24
	}
L7:
	;
	if int32(0) < v86 {
		goto L6
	} else {
		goto L22
	}
L8:
	;
	v80 = m.G398
	if v37 != v80 {
		goto L20
	} else {
		goto L21
	}
L10:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v37 = v32 + int32(48)
	v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if base.Ui32(v37) < base.Ui32(v38) {
		goto L8
	} else {
		goto L11
	}
L11:
	;
	v86 = int32(-1)
	goto L7
L20:
	;
	v83 = *(*int32)(unsafe.Add(mBase, uint32(v37)+8))
	v86 = v83
	goto L7
L21:
	;
	v86 = int32(-1)
	goto L7
L22:
	;
	v90 = F_lua_objlen(m, l0, int32(1))
	mBase = m.M
	v91 = m.ExcPending
	if v91 != 0 {
		goto L1
	} else {
		goto L23
	}
L23:
	;
	v95 = v90
	goto L5
L24:
	;
	v95 = v93
	goto L5
L25:
	;
	if v95 <= v27 {
		goto L28
	} else {
		goto L29
	}
L26:
	;
	F_luaL_pushresult(m, v9+int32(36))
	mBase = m.M
	v559 = m.ExcPending
	if v559 != 0 {
		goto L1
	} else {
		goto L141
	}
L27:
	;
	goto L91
L28:
	;
	if v27 != v95 {
		goto L26
	} else {
		goto L87
	}
L29:
	;
	v108 = v27
	goto L30
L30:
	;
	goto L35
L32:
	;
	goto L51
L33:
	;
	v167 = *(*int32)(unsafe.Add(mBase, uint32(v125)))
	v168 = F_luaH_getnum(m, v167, v108)
	mBase = m.M
	v169 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v170 = *(*int64)(unsafe.Add(mBase, uint32(v168)))
	*(*int64)(unsafe.Add(mBase, uint32(v169))) = v170
	v172 = *(*int32)(unsafe.Add(mBase, uint32(v168)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v169)+8)) = v172
	v174 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v174 + int32(16)
	goto L32
L35:
	;
	v116 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v117 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v122 = v117 + int32(0)
	v123 = m.G398
	if base.Ui32(v122) < base.Ui32(v116) {
		goto L36
	} else {
		goto L37
	}
L36:
	;
	v125 = v122
	goto L38
L37:
	;
	v125 = v123
	goto L38
L38:
	;
	goto L33
L48:
	;
	F_luaL_addvalue(m, v9+int32(36))
	mBase = m.M
	v323 = m.ExcPending
	if v323 != 0 {
		goto L1
	} else {
		goto L84
	}
L49:
	;
	if v239 != 0 {
		goto L48
	} else {
		goto L64
	}
L50:
	;
	v229 = m.G398
	if v195 != v229 {
		goto L62
	} else {
		goto L63
	}
L51:
	;
	goto L55
L55:
	;
	v192 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v195 = v192 + int32(-16)
	goto L50
L62:
	;
	v232 = *(*int32)(unsafe.Add(mBase, uint32(v195)+8))
	v239 = base.B2i32(base.Ui32(v232+int32(-3)) < base.Ui32(int32(2)))
	goto L49
L63:
	;
	v239 = int32(0)
	goto L49
L64:
	;
	goto L67
L65:
	;
	v299 = m.G3
	if v297 != int32(-1) {
		goto L81
	} else {
		goto L82
	}
L66:
	;
	v291 = m.G398
	if v257 != v291 {
		goto L78
	} else {
		goto L79
	}
L67:
	;
	goto L71
L71:
	;
	v254 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v257 = v254 + int32(-16)
	goto L66
L78:
	;
	v294 = *(*int32)(unsafe.Add(mBase, uint32(v257)+8))
	v297 = v294
	goto L65
L79:
	;
	v297 = int32(-1)
	goto L65
L80:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9)+20)) = v108
	*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = v309
	v312 = m.G3
	v317 = F_luaL_error(m, l0, v312+int32(_a2715), v9+int32(16))
	mBase = m.M
	v318 = m.ExcPending
	if v318 != 0 {
		goto L1
	} else {
		goto L83
	}
L81:
	;
	v304 = m.G399
	v308 = *(*int32)(unsafe.Add(mBase, uint32(v304+v297<<(uint(int32(2))%32))))
	v309 = v308
	goto L80
L82:
	;
	v309 = v299 + int32(_a2694)
	goto L80
L83:
	;
	goto L48
L84:
	;
	v326 = *(*int32)(unsafe.Add(mBase, uint32(v9)+32))
	F_luaL_addlstring(m, v9+int32(36), v17, v326)
	mBase = m.M
	v328 = m.ExcPending
	if v328 != 0 {
		goto L1
	} else {
		goto L85
	}
L85:
	;
	v330 = v108 + int32(1)
	if v330 != v95 {
		v108 = v330
		goto L30
	} else {
		goto L86
	}
L86:
	;
	goto L27
L87:
	;
	goto L27
L88:
	;
	goto L107
L89:
	;
	v395 = *(*int32)(unsafe.Add(mBase, uint32(v353)))
	v396 = F_luaH_getnum(m, v395, v95)
	mBase = m.M
	v397 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v398 = *(*int64)(unsafe.Add(mBase, uint32(v396)))
	*(*int64)(unsafe.Add(mBase, uint32(v397))) = v398
	v400 = *(*int32)(unsafe.Add(mBase, uint32(v396)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v397)+8)) = v400
	v402 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v402 + int32(16)
	goto L88
L91:
	;
	v344 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v345 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v350 = v345 + int32(0)
	v351 = m.G398
	if base.Ui32(v350) < base.Ui32(v344) {
		goto L92
	} else {
		goto L93
	}
L92:
	;
	v353 = v350
	goto L94
L93:
	;
	v353 = v351
	goto L94
L94:
	;
	goto L89
L104:
	;
	F_luaL_addvalue(m, v9+int32(36))
	mBase = m.M
	v549 = m.ExcPending
	if v549 != 0 {
		goto L1
	} else {
		goto L140
	}
L105:
	;
	if v467 != 0 {
		goto L104
	} else {
		goto L120
	}
L106:
	;
	v457 = m.G398
	if v423 != v457 {
		goto L118
	} else {
		goto L119
	}
L107:
	;
	goto L111
L111:
	;
	v420 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v423 = v420 + int32(-16)
	goto L106
L118:
	;
	v460 = *(*int32)(unsafe.Add(mBase, uint32(v423)+8))
	v467 = base.B2i32(base.Ui32(v460+int32(-3)) < base.Ui32(int32(2)))
	goto L105
L119:
	;
	v467 = int32(0)
	goto L105
L120:
	;
	goto L123
L121:
	;
	v527 = m.G3
	if v525 != int32(-1) {
		goto L137
	} else {
		goto L138
	}
L122:
	;
	v519 = m.G398
	if v485 != v519 {
		goto L134
	} else {
		goto L135
	}
L123:
	;
	goto L127
L127:
	;
	v482 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v485 = v482 + int32(-16)
	goto L122
L134:
	;
	v522 = *(*int32)(unsafe.Add(mBase, uint32(v485)+8))
	v525 = v522
	goto L121
L135:
	;
	v525 = int32(-1)
	goto L121
L136:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9)+4)) = v95
	*(*int32)(unsafe.Add(mBase, uint32(v9))) = v537
	v540 = m.G3
	v543 = F_luaL_error(m, l0, v540+int32(_a2715), v9)
	mBase = m.M
	v544 = m.ExcPending
	if v544 != 0 {
		goto L1
	} else {
		goto L139
	}
L137:
	;
	v532 = m.G399
	v536 = *(*int32)(unsafe.Add(mBase, uint32(v532+v525<<(uint(int32(2))%32))))
	v537 = v536
	goto L136
L138:
	;
	v537 = v527 + int32(_a2694)
	goto L136
L139:
	;
	goto L104
L140:
	;
	goto L26
L141:
	;
	m.G0 = v9 + int32(1072)
	return int32(1)
}
func F_timestampIsExpired(m *base.Module, l0 int64) int32 {
	mBase := m.M
	_ = mBase
	var v7 int64
	_ = v7
	if int64(0) <= l0 {
		v7 = *(*int64)(unsafe.Add(mBase, _consts[12]))
		return base.B2i32(l0 < v7)
	} else {
		return int32(0)
	}
}
func F_tlsResetCertInfo(m *base.Module) {
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
	var v16 int32
	_ = v16
	var v20 int32
	_ = v20
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
	var v40 int32
	_ = v40
	v3 = *(*int32)(unsafe.Add(mBase, _consts[139]))
	if v3 != 0 {
		return
	} else {
		v5 = *(*int32)(unsafe.Add(mBase, _consts[170]))
		if v5 != 0 {
			return
		} else {
			v7 = *(*int32)(unsafe.Add(mBase, _consts[107]))
			if v7 != 0 {
				return
			} else {
				v8 = int32(_a20)
				*(*int64)(unsafe.Add(mBase, _consts[879])) = int64(0)
				v12 = *(*int32)(unsafe.Add(mBase, _consts[880]))
				if v12 == int32(0) {
					v20 = int32(_a20)
					*(*int64)(unsafe.Add(mBase, _consts[881])) = int64(0)
					v24 = *(*int32)(unsafe.Add(mBase, _consts[882]))
					if v24 == int32(0) {
						v32 = int32(_a20)
						*(*int64)(unsafe.Add(mBase, _consts[883])) = int64(0)
						v36 = *(*int32)(unsafe.Add(mBase, _consts[884]))
						if v36 == int32(0) {
							return
						} else {
							F_sdsfree(m, v36)
							mBase = m.M
							v40 = m.ExcPending
							if v40 != 0 {
								return
							} else {
								*(*int32)(unsafe.Add(mBase, _consts[884])) = int32(0)
								return
							}
						}
					} else {
						F_sdsfree(m, v24)
						mBase = m.M
						v28 = m.ExcPending
						if v28 != 0 {
							return
						} else {
							*(*int32)(unsafe.Add(mBase, _consts[882])) = int32(0)
							v32 = int32(_a20)
							*(*int64)(unsafe.Add(mBase, _consts[883])) = int64(0)
							v36 = *(*int32)(unsafe.Add(mBase, _consts[884]))
							if v36 == int32(0) {
								return
							} else {
								F_sdsfree(m, v36)
								mBase = m.M
								v40 = m.ExcPending
								if v40 != 0 {
									return
								} else {
									*(*int32)(unsafe.Add(mBase, _consts[884])) = int32(0)
									return
								}
							}
						}
					}
				} else {
					F_sdsfree(m, v12)
					mBase = m.M
					v16 = m.ExcPending
					if v16 != 0 {
						return
					} else {
						*(*int32)(unsafe.Add(mBase, _consts[880])) = int32(0)
						v20 = int32(_a20)
						*(*int64)(unsafe.Add(mBase, _consts[881])) = int64(0)
						v24 = *(*int32)(unsafe.Add(mBase, _consts[882]))
						if v24 == int32(0) {
							v32 = int32(_a20)
							*(*int64)(unsafe.Add(mBase, _consts[883])) = int64(0)
							v36 = *(*int32)(unsafe.Add(mBase, _consts[884]))
							if v36 == int32(0) {
								return
							} else {
								F_sdsfree(m, v36)
								mBase = m.M
								v40 = m.ExcPending
								if v40 != 0 {
									return
								} else {
									*(*int32)(unsafe.Add(mBase, _consts[884])) = int32(0)
									return
								}
							}
						} else {
							F_sdsfree(m, v24)
							mBase = m.M
							v28 = m.ExcPending
							if v28 != 0 {
								return
							} else {
								*(*int32)(unsafe.Add(mBase, _consts[882])) = int32(0)
								v32 = int32(_a20)
								*(*int64)(unsafe.Add(mBase, _consts[883])) = int64(0)
								v36 = *(*int32)(unsafe.Add(mBase, _consts[884]))
								if v36 == int32(0) {
									return
								} else {
									F_sdsfree(m, v36)
									mBase = m.M
									v40 = m.ExcPending
									if v40 != 0 {
										return
									} else {
										*(*int32)(unsafe.Add(mBase, _consts[884])) = int32(0)
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
func F_top12_1(m *base.Module, l0 float64) int32 {
	return base.I32_wrap_i64(int64(base.Ui64(base.I64_reinterpret_f64(l0)) >> (uint(int64(52)) % 64)))
}
func F_touchWatchedKey(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
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
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
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
	var v65 int32
	_ = v65
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v86 int32
	_ = v86
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(v9)+12))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(v9)+16))
	if v10 == int32(0)-v12 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	m.G0 = v7 + int32(16)
	return
L2:
	;
	v15 = F_dictFetchValue(m, v9, l1)
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	return
L4:
	;
	if v15 == int32(0) {
		goto L1
	} else {
		goto L5
	}
L5:
	;
	v20 = v7 + int32(8)
	v21 = *(*int32)(unsafe.Add(mBase, uint32(v15)))
	*(*int32)(unsafe.Add(mBase, uint32(v20)+4)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v20))) = v21
	goto L6
L6:
	;
	v26 = v7 + int32(8)
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v26)))
	if v28 == int32(0) {
		goto L8
	} else {
		goto L9
	}
L7:
	;
	if v28 == int32(0) {
		goto L1
	} else {
		goto L10
	}
L8:
	;
	goto L7
L9:
	;
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v26)+4))
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v28+base.B2i32(v31 == int32(0))<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v26))) = v37
	goto L8
L10:
	;
	v44 = v28
	goto L11
L11:
	;
	v45 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v44)+24)))
	if v45&int32(1) == int32(0) {
		goto L14
	} else {
		goto L15
	}
L12:
	;
	goto L1
L13:
	;
	v75 = v7 + int32(8)
	v77 = *(*int32)(unsafe.Add(mBase, uint32(v75)))
	if v77 == int32(0) {
		goto L24
	} else {
		goto L25
	}
L14:
	;
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v44)+20))
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v64)+200))
	*(*int32)(unsafe.Add(mBase, uint32(v64)+200)) = v65 | int32(32)
	F_resetClientMultiState(m, v64)
	mBase = m.M
	v70 = m.ExcPending
	if v70 != 0 {
		goto L3
	} else {
		goto L21
	}
L15:
	;
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v44)+16))
	if l0 != v50 {
		goto L1
	} else {
		goto L16
	}
L16:
	;
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v44)+12))
	v53 = F_equalStringObjects(m, l1, v52)
	mBase = m.M
	v54 = m.ExcPending
	if v54 != 0 {
		goto L3
	} else {
		goto L17
	}
L17:
	;
	if v53 == int32(0) {
		goto L1
	} else {
		goto L18
	}
L18:
	;
	v57 = F_objectGetVal(m, l1)
	mBase = m.M
	v58 = F_dbFind(m, l0, v57)
	mBase = m.M
	v59 = m.ExcPending
	if v59 != 0 {
		goto L3
	} else {
		goto L19
	}
L19:
	;
	if v58 != 0 {
		goto L1
	} else {
		goto L20
	}
L20:
	;
	v60 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v44)+24)))
	v62 = v60 & int32(254)
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+24)) = uint8(v62)
	goto L13
L21:
	;
	F_unwatchAllKeys(m, v64)
	mBase = m.M
	v72 = m.ExcPending
	if v72 != 0 {
		goto L3
	} else {
		goto L22
	}
L22:
	;
	goto L13
L23:
	;
	if v77 != 0 {
		v44 = v77
		goto L11
	} else {
		goto L26
	}
L24:
	;
	goto L23
L25:
	;
	v80 = *(*int32)(unsafe.Add(mBase, uint32(v75)+4))
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v77+base.B2i32(v80 == int32(0))<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v75))) = v86
	goto L24
L26:
	;
	goto L12
}
func F_towupper(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v21 int32
	_ = v21
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v36 int32
	_ = v36
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v63 int32
	_ = v63
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
	var v129 int32
	_ = v129
	v2 = int32(1)
	if base.Ui32(int32(131071)) < base.Ui32(l0) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	return v129
L2:
	;
	v129 = l0
	goto L1
L3:
	;
	v12 = int32(255)
	v13 = l0 & v12
	v14 = int32(3)
	v15 = base.I32_div_u_s(v13, v14)
	v21 = int32(2)
	v25 = *(*int32)(unsafe.Add(mBase, uint32((l0-v15*v14)&v12<<(uint(v21)%32))+uint32(_consts[1061])))
	v26 = int32(8)
	v27 = int32(base.Ui32(l0) >> (uint(v26) % 32))
	v30 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v27)+uint32(_consts[1062]))))
	v36 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v30*int32(86)+v15)+uint32(_consts[1062]))))
	v41 = base.I32_rem_u_s(int32(base.Ui32(v25*v36)>>(uint(int32(11))%32)), int32(6))
	v44 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v27)+uint32(_consts[1063]))))
	v50 = *(*int32)(unsafe.Add(mBase, uint32((v41+v44)<<(uint(v21)%32))+uint32(_consts[1064])))
	v52 = v50 >> (uint(v26) % 32)
	v54 = v50 & v12
	if base.Ui32(int32(1)) < base.Ui32(v54) {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v63 = v52 & int32(255)
	if v63 == int32(0) {
		goto L2
	} else {
		goto L6
	}
L5:
	;
	v129 = v52&(int32(0)-(v54^v2)) + l0
	goto L1
L6:
	;
	v71 = v63
	v72 = int32(base.Ui32(v52) >> (uint(int32(8)) % 32))
	goto L7
L7:
	;
	v77 = int32(1)
	v78 = int32(base.Ui32(v71) >> (uint(v77) % 32))
	v79 = v78 + v72
	v81 = v79 << (uint(v77) % 32)
	v84 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v81)+uint32(_consts[1065]))))
	if v13 != v84 {
		goto L9
	} else {
		goto L10
	}
L8:
	;
	goto L2
L9:
	;
	v107 = base.B2i32(base.Ui32(v13) < base.Ui32(v84))
	if base.Ui32(v13) < base.Ui32(v84) {
		goto L16
	} else {
		goto L17
	}
L10:
	;
	v86 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v81)+uint32(_consts[1066]))))
	v91 = *(*int32)(unsafe.Add(mBase, uint32(v86<<(uint(int32(2))%32))+uint32(_consts[1064])))
	v93 = v91 & int32(255)
	if base.Ui32(int32(1)) < base.Ui32(v93) {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	goto L13
L12:
	;
	v129 = v91>>(uint(int32(8))%32)&(int32(0)-(v93^v2)) + l0
	goto L1
L13:
	;
	goto L15
L15:
	;
	v129 = int32(-1) + l0
	goto L1
L16:
	;
	v108 = v72
	goto L18
L17:
	;
	v108 = v79
	goto L18
L18:
	;
	if base.Ui32(v13) < base.Ui32(v84) {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	v110 = v78
	goto L21
L20:
	;
	v110 = v71 - v78
	goto L21
L21:
	;
	if v110 != 0 {
		v71 = v110
		v72 = v108
		goto L7
	} else {
		goto L22
	}
L22:
	;
	goto L8
}
func F_traceCommandHandler(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v18 int32
	_ = v18
	var v25 int32
	_ = v25
	var v57 int32
	_ = v57
	var v61 int32
	_ = v61
	var v72 int32
	_ = v72
	var v76 int32
	_ = v76
	var v84 int32
	_ = v84
	var v89 int32
	_ = v89
	var v94 int32
	_ = v94
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v127 int32
	_ = v127
	var v130 int32
	_ = v130
	var v133 int32
	_ = v133
	var v136 int32
	_ = v136
	var v139 int32
	_ = v139
	var v141 int32
	_ = v141
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
	var v158 int32
	_ = v158
	var v161 int32
	_ = v161
	var v165 int32
	_ = v165
	var v170 int32
	_ = v170
	var v172 int32
	_ = v172
	var v176 int32
	_ = v176
	var v182 int32
	_ = v182
	var v185 int32
	_ = v185
	var v191 int32
	_ = v191
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
	var v209 int32
	_ = v209
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v215 int32
	_ = v215
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v222 int32
	_ = v222
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
	var v245 int32
	_ = v245
	var v254 int32
	_ = v254
	var v255 int32
	_ = v255
	var v259 int32
	_ = v259
	var v261 int32
	_ = v261
	var v263 int32
	_ = v263
	var v275 int32
	_ = v275
	var v281 int32
	_ = v281
	var v286 int32
	_ = v286
	var v287 int32
	_ = v287
	var v289 int32
	_ = v289
	var v291 int32
	_ = v291
	var v292 int32
	_ = v292
	var v293 int32
	_ = v293
	var v294 int32
	_ = v294
	var v298 int32
	_ = v298
	var v299 int32
	_ = v299
	var v301 int32
	_ = v301
	var v303 int32
	_ = v303
	var v315 int32
	_ = v315
	var v316 int32
	_ = v316
	var v322 int32
	_ = v322
	var v325 int32
	_ = v325
	var v328 int32
	_ = v328
	var v330 int32
	_ = v330
	var v335 int32
	_ = v335
	var v336 int32
	_ = v336
	var v337 int32
	_ = v337
	var v338 int32
	_ = v338
	var v339 int32
	_ = v339
	var v341 int32
	_ = v341
	var v343 int32
	_ = v343
	var v347 int32
	_ = v347
	var v349 int32
	_ = v349
	var v354 int32
	_ = v354
	var v358 int32
	_ = v358
	var v360 int32
	_ = v360
	var v369 int32
	_ = v369
	var v370 int32
	_ = v370
	var v371 int32
	_ = v371
	var v372 int32
	_ = v372
	var v373 int32
	_ = v373
	var v378 int32
	_ = v378
	var v379 int32
	_ = v379
	var v381 int32
	_ = v381
	var v383 int32
	_ = v383
	var v395 int32
	_ = v395
	var v396 int32
	_ = v396
	var v398 int32
	_ = v398
	var v404 int32
	_ = v404
	var v410 int32
	_ = v410
	var v411 int32
	_ = v411
	var v413 int32
	_ = v413
	v12 = m.G0
	v14 = v12 - int32(128)
	m.G0 = v14
	if l2 == int32(0) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v404 = m.G3
	v410 = m.G8
	v411 = *(*int32)(unsafe.Add(mBase, uint32(v410)))
	m.T0[v411].(func(*base.Module, int32, int32, int32))(m, v404+int32(_a2587), v404+int32(_a2588), int32(606))
	mBase = m.M
	v413 = m.ExcPending
	if v413 != 0 {
		goto L21
	} else {
		goto L100
	}
L2:
	;
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v18)+20))
	goto L7
L3:
	;
	v395 = m.G14
	v396 = *(*int32)(unsafe.Add(mBase, uint32(v395)))
	m.T0[v396].(func(*base.Module))(m)
	mBase = m.M
	v398 = m.ExcPending
	if v398 != 0 {
		goto L21
	} else {
		goto L99
	}
L4:
	;
	v370 = m.G3
	v371 = m.G13
	v372 = *(*int32)(unsafe.Add(mBase, uint32(v371)))
	v373 = m.G12
	v378 = m.T0[v372].(func(*base.Module, int32, int32, int32) int32)(m, int32(0), v370+int32(_a2596), int32(33))
	mBase = m.M
	v379 = m.ExcPending
	if v379 != 0 {
		goto L21
	} else {
		goto L97
	}
L5:
	;
	if v72 == int32(0) {
		goto L4
	} else {
		goto L18
	}
L6:
	;
	goto L5
L7:
	;
	goto L16
L15:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14+int32(28))+96)) = v61
	v72 = int32(1)
	goto L6
L16:
	;
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v18)+40))
	if base.Ui32(v25) <= base.Ui32(v57) {
		v72 = int32(0)
		goto L6
	} else {
		goto L17
	}
L17:
	;
	v61 = base.I32_div_s(v25-v57, int32(24))
	goto L15
L18:
	;
	v76 = v14 + int32(64)
	v84 = int32(0)
	goto L19
L19:
	;
	v89 = m.G3
	v94 = F_lua_getinfo(m, v18, v89+int32(_a2597), v14+int32(28))
	mBase = m.M
	v97 = m.ExcPending
	if v97 != 0 {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	return int32(0)
L22:
	;
	v99 = v89 + int32(_a2598)
	v102 = int32(*(*int8)(unsafe.Add(mBase, uint32(v89)+uint32(_consts[979]))))
	if v102 != 0 {
		goto L25
	} else {
		goto L26
	}
L23:
	;
	v315 = int32(1)
	v316 = v84 + v315
	v322 = *(*int32)(unsafe.Add(mBase, uint32(v18)+20))
	if v316 < v315 {
		v347 = v316
		v349 = v322
		goto L85
	} else {
		goto L86
	}
L24:
	;
	if v127 == int32(0) {
		goto L23
	} else {
		goto L40
	}
L25:
	;
	v103 = int32(0)
	v104 = F_strchr(m, v76, v102)
	mBase = m.M
	if v104 == v103 {
		v124 = v103
		goto L27
	} else {
		goto L28
	}
L26:
	;
	v127 = v76
	goto L24
L27:
	;
	v127 = v124
	goto L24
L28:
	;
	v107 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v89)+uint32(_consts[980]))))
	if v107 != 0 {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	v108 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v104)+1)))
	if v108 == int32(0) {
		v124 = v103
		goto L27
	} else {
		goto L31
	}
L30:
	;
	v127 = v104
	goto L24
L31:
	;
	v111 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v89)+uint32(_consts[981]))))
	if v111 != 0 {
		goto L32
	} else {
		goto L33
	}
L32:
	;
	v113 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v104)+2)))
	if v113 == int32(0) {
		v124 = v103
		goto L27
	} else {
		goto L34
	}
L33:
	;
	v112 = F_twobyte_strstr(m, v104, v99)
	mBase = m.M
	v127 = v112
	goto L24
L34:
	;
	v116 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v89)+uint32(_consts[982]))))
	if v116 != 0 {
		goto L35
	} else {
		goto L36
	}
L35:
	;
	v118 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v104)+3)))
	if v118 == int32(0) {
		v124 = v103
		goto L27
	} else {
		goto L37
	}
L36:
	;
	v117 = F_threebyte_strstr(m, v104, v99)
	mBase = m.M
	v127 = v117
	goto L24
L37:
	;
	v121 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v89)+uint32(_consts[983]))))
	if v121 != 0 {
		goto L38
	} else {
		goto L39
	}
L38:
	;
	v123 = F_twoway_strstr(m, v104, v99)
	mBase = m.M
	v124 = v123
	goto L27
L39:
	;
	v122 = F_fourbyte_strstr(m, v104, v99)
	mBase = m.M
	v127 = v122
	goto L24
L40:
	;
	v130 = m.G3
	if v84 != 0 {
		goto L41
	} else {
		goto L42
	}
L41:
	;
	v133 = int32(_a2599)
	goto L43
L42:
	;
	v133 = int32(_a2600)
	goto L43
L43:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+16)) = v130 + v133
	v136 = *(*int32)(unsafe.Add(mBase, uint32(v14)+32))
	if v136 != 0 {
		goto L44
	} else {
		goto L45
	}
L44:
	;
	v139 = v136
	goto L46
L45:
	;
	v139 = v130 + int32(_a2601)
	goto L46
L46:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+20)) = v139
	v141 = m.G13
	v146 = F_lm_asprintf(m, v130+int32(_a2602), v14+int32(16))
	mBase = m.M
	v147 = m.ExcPending
	if v147 != 0 {
		goto L21
	} else {
		goto L47
	}
L47:
	;
	v148 = *(*int32)(unsafe.Add(mBase, uint32(v141)))
	v149 = m.G12
	v150 = int32(0)
	if v146&int32(3) == v150 {
		v172 = v146
		goto L50
	} else {
		goto L51
	}
L48:
	;
	v206 = m.T0[v148].(func(*base.Module, int32, int32, int32) int32)(m, v150, v146, v205)
	mBase = m.M
	v207 = m.ExcPending
	if v207 != 0 {
		goto L21
	} else {
		goto L64
	}
L49:
	;
	v205 = v197 - v146
	goto L48
L50:
	;
	v176 = v172
	goto L58
L51:
	;
	v158 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v146))))
	if v158 != 0 {
		goto L52
	} else {
		goto L53
	}
L52:
	;
	v161 = v146
	goto L54
L53:
	;
	v205 = v146 - v146
	goto L48
L54:
	;
	v165 = v161 + int32(1)
	if v165&int32(3) == int32(0) {
		v172 = v165
		goto L50
	} else {
		goto L56
	}
L56:
	;
	v170 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v165))))
	if v170 != 0 {
		v161 = v165
		goto L54
	} else {
		goto L57
	}
L57:
	;
	v197 = v165
	goto L49
L58:
	;
	v182 = *(*int32)(unsafe.Add(mBase, uint32(v176)))
	v185 = int32(-2139062144)
	if (int32(16843008)-v182|v182)&v185 == v185 {
		v176 = v176 + int32(4)
		goto L58
	} else {
		goto L60
	}
L59:
	;
	v191 = v176
	goto L61
L60:
	;
	goto L59
L61:
	;
	v195 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v191))))
	if v195 != 0 {
		v191 = v191 + int32(1)
		goto L61
	} else {
		goto L63
	}
L62:
	;
	v197 = v191
	goto L49
L63:
	;
	goto L62
L64:
	;
	v209 = *(*int32)(unsafe.Add(mBase, uint32(v149)))
	m.T0[v209].(func(*base.Module, int32, int32))(m, v206, int32(0))
	mBase = m.M
	v211 = m.ExcPending
	if v211 != 0 {
		goto L21
	} else {
		goto L65
	}
L65:
	;
	v212 = m.G11
	v213 = *(*int32)(unsafe.Add(mBase, uint32(v212)))
	m.T0[v213].(func(*base.Module, int32))(m, v146)
	mBase = m.M
	v215 = m.ExcPending
	if v215 != 0 {
		goto L21
	} else {
		goto L66
	}
L66:
	;
	v217 = v130 + int32(_a2603)
	v218 = *(*int32)(unsafe.Add(mBase, uint32(v14)+48))
	v219 = m.G6
	if v218 < int32(1) {
		v232 = v217
		goto L67
	} else {
		goto L68
	}
L67:
	;
	v233 = m.G3
	v234 = int32(0)
	v235 = m.G6
	v236 = *(*int32)(unsafe.Add(mBase, uint32(v235)+260))
	if v234 < v236 {
		goto L71
	} else {
		goto L72
	}
L68:
	;
	v222 = *(*int32)(unsafe.Add(mBase, uint32(v219)+276))
	if v222 < v218 {
		v232 = v217
		goto L67
	} else {
		goto L69
	}
L69:
	;
	v224 = m.G6
	v225 = *(*int32)(unsafe.Add(mBase, uint32(v224)+272))
	v231 = *(*int32)(unsafe.Add(mBase, uint32(v225+v218<<(uint(int32(2))%32)+int32(-4))))
	v232 = v231
	goto L67
L70:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+8)) = v232
	*(*int32)(unsafe.Add(mBase, uint32(v14)+4)) = v218
	v286 = m.G6
	v287 = *(*int32)(unsafe.Add(mBase, uint32(v286)+280))
	if v287 == v218 {
		goto L78
	} else {
		goto L79
	}
L71:
	;
	v245 = v234
	goto L74
L72:
	;
	v275 = v233 + int32(_a2604)
	v281 = v233 + int32(_a2605)
	goto L70
L73:
	;
	v275 = v254 + int32(_a2606)
	v281 = v254 + int32(_a2607)
	goto L70
L74:
	;
	v254 = m.G3
	v255 = m.G6
	v259 = *(*int32)(unsafe.Add(mBase, uint32(v255+v245<<(uint(int32(2))%32))+4))
	if v259 == v218 {
		goto L73
	} else {
		goto L76
	}
L75:
	;
	v275 = v261 + int32(_a2604)
	v281 = v261 + int32(_a2605)
	goto L70
L76:
	;
	v261 = m.G3
	v263 = v245 + int32(1)
	if v263 != v236 {
		v245 = v263
		goto L74
	} else {
		goto L77
	}
L77:
	;
	goto L75
L78:
	;
	v289 = v275
	goto L80
L79:
	;
	v289 = v281
	goto L80
L80:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14))) = v289
	v291 = m.G3
	v292 = m.G15
	v293 = *(*int32)(unsafe.Add(mBase, uint32(v292)))
	v294 = m.G12
	v298 = m.T0[v293].(func(*base.Module, int32, int32, int32) int32)(m, int32(0), v291+int32(_a2608), v14)
	mBase = m.M
	v299 = m.ExcPending
	if v299 != 0 {
		goto L21
	} else {
		goto L81
	}
L81:
	;
	v301 = *(*int32)(unsafe.Add(mBase, uint32(v294)))
	m.T0[v301].(func(*base.Module, int32, int32))(m, v298, int32(0))
	mBase = m.M
	v303 = m.ExcPending
	if v303 != 0 {
		goto L21
	} else {
		goto L82
	}
L82:
	;
	goto L23
L83:
	;
	if v369 != 0 {
		v84 = v316
		goto L19
	} else {
		goto L96
	}
L84:
	;
	goto L83
L85:
	;
	if v347 != 0 {
		v360 = int32(0)
		goto L93
	} else {
		goto L94
	}
L86:
	;
	v325 = *(*int32)(unsafe.Add(mBase, uint32(v18)+40))
	v328 = v316
	v330 = v322
	goto L87
L87:
	;
	if base.Ui32(v330) <= base.Ui32(v325) {
		v369 = int32(0)
		goto L84
	} else {
		goto L89
	}
L88:
	;
	v347 = v341
	v349 = v343
	goto L85
L89:
	;
	v335 = v328 + int32(-1)
	v336 = *(*int32)(unsafe.Add(mBase, uint32(v330)+4))
	v337 = *(*int32)(unsafe.Add(mBase, uint32(v336)))
	v338 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v337)+6)))
	if v338 != 0 {
		v341 = v335
		goto L90
	} else {
		goto L91
	}
L90:
	;
	v343 = v330 + int32(-24)
	if int32(0) < v341 {
		v328 = v341
		v330 = v343
		goto L87
	} else {
		goto L92
	}
L91:
	;
	v339 = *(*int32)(unsafe.Add(mBase, uint32(v330)+20))
	v341 = v335 - v339
	goto L90
L92:
	;
	goto L88
L93:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14+int32(28))+96)) = v360
	v369 = int32(1)
	goto L84
L94:
	;
	v354 = *(*int32)(unsafe.Add(mBase, uint32(v18)+40))
	if base.Ui32(v349) <= base.Ui32(v354) {
		v369 = int32(0)
		goto L84
	} else {
		goto L95
	}
L95:
	;
	v358 = base.I32_div_s(v349-v354, int32(24))
	v360 = v358
	goto L93
L96:
	;
	goto L3
L97:
	;
	v381 = *(*int32)(unsafe.Add(mBase, uint32(v373)))
	m.T0[v381].(func(*base.Module, int32, int32))(m, v378, int32(0))
	mBase = m.M
	v383 = m.ExcPending
	if v383 != 0 {
		goto L21
	} else {
		goto L98
	}
L98:
	;
	goto L3
L99:
	;
	m.G0 = v14 + int32(128)
	return int32(1)
L100:
	;
	m.Env.Exit(m, int32(1))
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_truncate(m *base.Module, l0 int32, l1 int64) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v11 int32
	_ = v11
	v3 = m.Env.X__syscall_truncate64(m, l0, l1)
	mBase = m.M
	if base.Ui32(v3) < base.Ui32(int32(-4095)) {
		v11 = v3
	} else {
		v6 = F___errno_location(m)
		mBase = m.M
		*(*int32)(unsafe.Add(mBase, uint32(v6))) = int32(0) - v3
		v11 = int32(-1)
	}
	return v11
}
func F_tryReadBulkPayloadMetadata(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32 {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v31 int32
	_ = v31
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
	var v45 int64
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v52 int32
	_ = v52
	var v54 int64
	_ = v54
	var v57 int32
	_ = v57
	var v59 int64
	_ = v59
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v100 int32
	_ = v100
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v109 int32
	_ = v109
	var v114 int32
	_ = v114
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v124 int32
	_ = v124
	var v126 int32
	_ = v126
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v134 int32
	_ = v134
	var v143 int32
	_ = v143
	var v151 int32
	_ = v151
	var v154 int32
	_ = v154
	var v158 int32
	_ = v158
	var v163 int32
	_ = v163
	var v165 int32
	_ = v165
	var v169 int32
	_ = v169
	var v175 int32
	_ = v175
	var v178 int32
	_ = v178
	var v184 int32
	_ = v184
	var v188 int32
	_ = v188
	var v190 int32
	_ = v190
	var v198 int32
	_ = v198
	var v201 int64
	_ = v201
	var v203 int32
	_ = v203
	var v207 int64
	_ = v207
	var v209 int32
	_ = v209
	var v213 int64
	_ = v213
	var v215 int32
	_ = v215
	var v219 int64
	_ = v219
	var v221 int32
	_ = v221
	var v225 int64
	_ = v225
	var v227 int64
	_ = v227
	var v249 int32
	_ = v249
	var v254 int64
	_ = v254
	var v258 int64
	_ = v258
	var v261 int32
	_ = v261
	var v263 int64
	_ = v263
	var v268 int32
	_ = v268
	v11 = m.G0
	v13 = v11 - int32(48)
	m.G0 = v13
	v15 = int32(-1)
	v18 = *(*int32)(unsafe.Add(mBase, _consts[599]))
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v23 = *(*int32)(unsafe.Add(mBase, uint32(v22)+100))
	v24 = m.T0[v23].(func(*base.Module, int32, int32, int32, int64) int32)(m, l0, l1, int32(1024), base.I64_extend_i32_s(v18*int32(1000)))
	mBase = m.M
	v27 = m.ExcPending
	if v27 != 0 {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	m.G0 = v13 + int32(48)
	return v268
L2:
	;
	v45 = base.I64_extend_i32_s(v24 + int32(1))
	v46 = int32(0)
	v47 = *(*int32)(unsafe.Add(mBase, _consts[600]))
	goto L11
L3:
	;
	return int32(0)
L4:
	;
	if v24 != int32(-1) {
		goto L2
	} else {
		goto L5
	}
L5:
	;
	v31 = *(*int32)(unsafe.Add(mBase, _consts[28]))
	if int32(3) < v31 {
		v268 = v15
		goto L1
	} else {
		goto L6
	}
L6:
	;
	v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v34)+88))
	v36 = m.T0[v35].(func(*base.Module, int32) int32)(m, l0)
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		goto L3
	} else {
		goto L7
	}
L7:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13))) = v36
	F__serverLog(m, int32(3), int32(_a1951), v13)
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L3
	} else {
		goto L8
	}
L8:
	;
	v268 = v15
	goto L1
L9:
	;
	v62 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	switch v62 + int32(-36) {
	case 0:
		goto L14
	case 1, 2, 3, 4, 5, 6, 7, 8:
		goto L15
	case 9:
		goto L17
	default:
		goto L16
	}
L10:
	;
	v57 = int32(_a20)
	v59 = *(*int64)(unsafe.Add(mBase, _consts[601]))
	*(*int64)(unsafe.Add(mBase, _consts[601])) = v59 + v45
	goto L9
L11:
	;
	if base.B2i32(v47 != v46) == int32(0) {
		goto L10
	} else {
		goto L12
	}
L12:
	;
	v52 = int32(_a20)
	v54 = *(*int64)(unsafe.Add(mBase, _consts[602]))
	*(*int64)(unsafe.Add(mBase, _consts[602])) = v54 + v45
	goto L9
L13:
	;
	v261 = int32(_a20)
	v263 = *(*int64)(unsafe.Add(mBase, _consts[109]))
	*(*int64)(unsafe.Add(mBase, _consts[603])) = v263
	v268 = int32(-2)
	goto L1
L14:
	;
	v94 = l1 + int32(1)
	v95 = int32(_a1952)
	goto L26
L15:
	;
	v81 = int32(-1)
	v83 = *(*int32)(unsafe.Add(mBase, _consts[28]))
	if int32(3) < v83 {
		v268 = v81
		goto L1
	} else {
		goto L21
	}
L16:
	;
	if v62 == int32(0) {
		goto L13
	} else {
		goto L20
	}
L17:
	;
	v65 = int32(-1)
	v67 = *(*int32)(unsafe.Add(mBase, _consts[28]))
	if int32(3) < v67 {
		v268 = v65
		goto L1
	} else {
		goto L18
	}
L18:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+32)) = l1 + int32(1)
	F__serverLog(m, int32(3), int32(_a1953), v13+int32(32))
	mBase = m.M
	v78 = m.ExcPending
	if v78 != 0 {
		goto L3
	} else {
		goto L19
	}
L19:
	;
	v268 = v65
	goto L1
L20:
	;
	goto L15
L21:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+16)) = l1
	F__serverLog(m, int32(3), int32(_a1954), v13+int32(16))
	mBase = m.M
	v92 = m.ExcPending
	if v92 != 0 {
		goto L3
	} else {
		goto L22
	}
L22:
	;
	v268 = v81
	goto L1
L23:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l5))) = v258
	v268 = int32(0)
	goto L1
L24:
	;
	v249 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l4))) = v249
	v254 = F_strtox_2(m, v94, v249, int32(10), int64(2147483648))
	mBase = m.M
	goto L56
L25:
	;
	if v129-v134 != 0 {
		goto L24
	} else {
		goto L38
	}
L26:
	;
	v100 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v94))))
	if v100 != 0 {
		goto L29
	} else {
		goto L30
	}
L28:
	;
	v134 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v130))))
	goto L25
L29:
	;
	v102 = v94
	v103 = v95
	v104 = int32(4)
	v105 = v100
	goto L32
L30:
	;
	v129 = int32(0)
	v130 = v95
	goto L28
L31:
	;
	v129 = v126 & int32(255)
	v130 = v124
	goto L28
L32:
	;
	v109 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v103))))
	if v105&int32(255) != v109 {
		v124 = v103
		v126 = v105
		goto L31
	} else {
		goto L34
	}
L33:
	;
	v124 = v118
	v126 = int32(0)
	goto L31
L34:
	;
	if v109 == int32(0) {
		v124 = v103
		v126 = v105
		goto L31
	} else {
		goto L35
	}
L35:
	;
	v114 = v104 + int32(-1)
	if v114 == int32(0) {
		v124 = v103
		v126 = v105
		goto L31
	} else {
		goto L36
	}
L36:
	;
	v117 = int32(1)
	v118 = v103 + v117
	v119 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v102)+1)))
	if v119 != 0 {
		v102 = v102 + v117
		v103 = v118
		v104 = v114
		v105 = v119
		goto L32
	} else {
		goto L37
	}
L37:
	;
	goto L33
L38:
	;
	v143 = l1 + int32(5)
	if v143&int32(3) == int32(0) {
		v165 = v143
		goto L41
	} else {
		goto L42
	}
L39:
	;
	if base.Ui32(v198) < base.Ui32(int32(40)) {
		goto L24
	} else {
		goto L55
	}
L40:
	;
	v198 = v190 - v143
	goto L39
L41:
	;
	v169 = v165
	goto L49
L42:
	;
	v151 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v143))))
	if v151 != 0 {
		goto L43
	} else {
		goto L44
	}
L43:
	;
	v154 = v143
	goto L45
L44:
	;
	v198 = v143 - v143
	goto L39
L45:
	;
	v158 = v154 + int32(1)
	if v158&int32(3) == int32(0) {
		v165 = v158
		goto L41
	} else {
		goto L47
	}
L47:
	;
	v163 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v158))))
	if v163 != 0 {
		v154 = v158
		goto L45
	} else {
		goto L48
	}
L48:
	;
	v190 = v158
	goto L40
L49:
	;
	v175 = *(*int32)(unsafe.Add(mBase, uint32(v169)))
	v178 = int32(-2139062144)
	if (int32(16843008)-v175|v175)&v178 == v178 {
		v169 = v169 + int32(4)
		goto L49
	} else {
		goto L51
	}
L50:
	;
	v184 = v169
	goto L52
L51:
	;
	goto L50
L52:
	;
	v188 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v184))))
	if v188 != 0 {
		v184 = v184 + int32(1)
		goto L52
	} else {
		goto L54
	}
L53:
	;
	v190 = v184
	goto L40
L54:
	;
	goto L53
L55:
	;
	v201 = *(*int64)(unsafe.Add(mBase, uint32(v143)))
	*(*int64)(unsafe.Add(mBase, uint32(l2))) = v201
	v203 = int32(32)
	v207 = *(*int64)(unsafe.Add(mBase, uint32(l1+int32(37))))
	*(*int64)(unsafe.Add(mBase, uint32(l2+v203))) = v207
	v209 = int32(24)
	v213 = *(*int64)(unsafe.Add(mBase, uint32(l1+int32(29))))
	*(*int64)(unsafe.Add(mBase, uint32(l2+v209))) = v213
	v215 = int32(16)
	v219 = *(*int64)(unsafe.Add(mBase, uint32(l1+int32(21))))
	*(*int64)(unsafe.Add(mBase, uint32(l2+v215))) = v219
	v221 = int32(8)
	v225 = *(*int64)(unsafe.Add(mBase, uint32(l1+int32(13))))
	*(*int64)(unsafe.Add(mBase, uint32(l2+v221))) = v225
	v227 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(l3+v203))) = v227
	*(*int64)(unsafe.Add(mBase, uint32(l3+v209))) = v227
	*(*int64)(unsafe.Add(mBase, uint32(l3+v215))) = v227
	*(*int64)(unsafe.Add(mBase, uint32(l3+v221))) = v227
	*(*int64)(unsafe.Add(mBase, uint32(l3))) = v227
	*(*int32)(unsafe.Add(mBase, uint32(l4))) = int32(1)
	v258 = v227
	goto L23
L56:
	;
	v258 = base.I64_extend_i32_s(base.I32_wrap_i64(v254))
	goto L23
}
func F_twobyte_strstr(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v50 int32
	_ = v50
	v3 = int32(0)
	v7 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+1)))
	v9 = base.B2i32(v7 != v3)
	if v7 == v3 {
		v43 = l0
		v46 = v9
		goto L1
	} else {
		goto L2
	}
L1:
	;
	if v46 != 0 {
		goto L8
	} else {
		goto L9
	}
L2:
	;
	v12 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	v13 = int32(8)
	v15 = v12<<(uint(v13)%32) | v7
	v16 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	v19 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+1)))
	v20 = v16<<(uint(v13)%32) | v19
	if v15 == v20 {
		v43 = l0
		v46 = v9
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v25 = l0 + int32(1)
	v28 = v15
	goto L4
L4:
	;
	v30 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v25)+1)))
	v31 = int32(0)
	v32 = base.B2i32(v30 != v31)
	if v30 == v31 {
		v43 = v25
		v46 = v32
		goto L1
	} else {
		goto L6
	}
L5:
	;
	v43 = v25
	v46 = v32
	goto L1
L6:
	;
	v41 = v28<<(uint(int32(8))%32)&int32(65280) | v30
	if v41 != v20 {
		v25 = v25 + int32(1)
		v28 = v41
		goto L4
	} else {
		goto L7
	}
L7:
	;
	goto L5
L8:
	;
	v50 = v43
	goto L10
L9:
	;
	v50 = int32(0)
	goto L10
L10:
	;
	return v50
}
