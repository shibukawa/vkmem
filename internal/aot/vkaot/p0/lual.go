package p0

import (
	base "github.com/shibukawa/vkmem/internal/aot/vkaot/base"
	"unsafe"
)

func F_luaL_addlstring(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
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
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v31 int32
	_ = v31
	if l2 == int32(0) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v12 = l1
	v13 = l2
	goto L3
L3:
	;
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if base.Ui32(v17) < base.Ui32(l0+int32(1036)) {
		v22 = v17
		goto L5
	} else {
		goto L6
	}
L4:
	;
	goto L1
L5:
	;
	v23 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12))))
	v24 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v22 + v24
	*(*uint8)(unsafe.Add(mBase, uint32(v22))) = uint8(v23)
	v31 = v13 + int32(-1)
	if v31 != 0 {
		v12 = v12 + v24
		v13 = v31
		goto L3
	} else {
		goto L9
	}
L6:
	;
	v19 = F_luaL_prepbuffer(m, l0)
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	return
L8:
	;
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v22 = v21
	goto L5
L9:
	;
	goto L4
}
func F_luaL_addvalue(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v54 int32
	_ = v54
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v91 int32
	_ = v91
	var v94 int32
	_ = v94
	var v130 int32
	_ = v130
	var v133 int32
	_ = v133
	var v137 int32
	_ = v137
	var v138 int64
	_ = v138
	var v142 int32
	_ = v142
	var v145 int32
	_ = v145
	var v147 int32
	_ = v147
	var v150 int64
	_ = v150
	var v152 int32
	_ = v152
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v160 int32
	_ = v160
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v167 int32
	_ = v167
	var v170 int32
	_ = v170
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
	var v185 int32
	_ = v185
	var v191 int32
	_ = v191
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	v10 = m.G0
	v12 = v10 - int32(16)
	m.G0 = v12
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v18 = F_lua_tolstring(m, v14, int32(-1), v12+int32(12))
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v20 = *(*int32)(unsafe.Add(mBase, uint32(v12)+12))
	v22 = l0 + int32(12)
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if base.Ui32(v22-v23+int32(1024)) < base.Ui32(v20) {
		goto L4
	} else {
		goto L5
	}
L3:
	;
	m.G0 = v12 + int32(16)
	return
L4:
	;
	if v23 == v22 {
		goto L17
	} else {
		goto L18
	}
L5:
	;
	if v20 == int32(0) {
		goto L7
	} else {
		goto L8
	}
L6:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v12)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v32 + v33
	goto L11
L7:
	;
	goto L6
L8:
	;
	v30 = F__emscripten_memcpy_bulkmem(m, v23, v18, v20)
	mBase = m.M
	goto L7
L9:
	;
	goto L3
L10:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+8)) = v54 + int32(-16)
	goto L9
L11:
	;
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v14)+8))
	goto L10
L17:
	;
	v154 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v155 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v154 + v155
	if v154 < v155 {
		goto L3
	} else {
		goto L41
	}
L18:
	;
	v65 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	F_lua_pushlstring(m, v65, v22, v23-v22)
	mBase = m.M
	v68 = m.ExcPending
	if v68 != 0 {
		goto L1
	} else {
		goto L19
	}
L19:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v22
	v70 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v70 + int32(1)
	goto L22
L20:
	;
	goto L17
L21:
	;
	v130 = *(*int32)(unsafe.Add(mBase, uint32(v14)+8))
	if base.Ui32(v130) <= base.Ui32(v94) {
		v147 = v130
		goto L36
	} else {
		goto L37
	}
L22:
	;
	goto L28
L28:
	;
	v91 = *(*int32)(unsafe.Add(mBase, uint32(v14)+8))
	v94 = v91 + int32(-32)
	goto L21
L36:
	;
	v150 = *(*int64)(unsafe.Add(mBase, uint32(v147)))
	*(*int64)(unsafe.Add(mBase, uint32(v94))) = v150
	v152 = *(*int32)(unsafe.Add(mBase, uint32(v147)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v94)+8)) = v152
	goto L20
L37:
	;
	v133 = v130
	goto L38
L38:
	;
	v137 = v133 + int32(-16)
	v138 = *(*int64)(unsafe.Add(mBase, uint32(v137)))
	*(*int64)(unsafe.Add(mBase, uint32(v133))) = v138
	v142 = *(*int32)(unsafe.Add(mBase, uint32(v133+int32(-8))))
	*(*int32)(unsafe.Add(mBase, uint32(v133)+8)) = v142
	if base.Ui32(v94) < base.Ui32(v137) {
		v133 = v137
		goto L38
	} else {
		goto L40
	}
L39:
	;
	v145 = *(*int32)(unsafe.Add(mBase, uint32(v14)+8))
	v147 = v145
	goto L36
L40:
	;
	goto L39
L41:
	;
	v160 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v162 = F_lua_objlen(m, v160, int32(-1))
	mBase = m.M
	v163 = m.ExcPending
	if v163 != 0 {
		goto L1
	} else {
		goto L42
	}
L42:
	;
	v167 = v162
	v170 = int32(1)
	goto L44
L43:
	;
	if v185 != 0 {
		goto L49
	} else {
		goto L50
	}
L44:
	;
	v176 = F_lua_objlen(m, v160, v170^int32(-1))
	mBase = m.M
	v177 = m.ExcPending
	if v177 != 0 {
		goto L1
	} else {
		goto L46
	}
L45:
	;
	goto L43
L46:
	;
	v178 = int32(1)
	v179 = v170 + v178
	v180 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v185 = base.B2i32(int32(8) < v180-v170) | base.B2i32(base.Ui32(v176) < base.Ui32(v167))
	if v185 != v178 {
		goto L43
	} else {
		goto L47
	}
L47:
	;
	if v179 < v180 {
		v167 = v176 + v167
		v170 = v179
		goto L44
	} else {
		goto L48
	}
L48:
	;
	goto L45
L49:
	;
	v191 = v179
	goto L51
L50:
	;
	v191 = v170
	goto L51
L51:
	;
	F_lua_concat(m, v160, v191)
	mBase = m.M
	v193 = m.ExcPending
	if v193 != 0 {
		goto L1
	} else {
		goto L52
	}
L52:
	;
	v194 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v194 - v191 + int32(1)
	goto L3
}
func F_luaL_buffinit(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l1)+8)) = l0
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = l1 + int32(12)
	return
}
func F_luaL_callmeta(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v17 int32
	_ = v17
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
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
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v96 int32
	_ = v96
	var v102 int32
	_ = v102
	var v108 int32
	_ = v108
	var v126 int32
	_ = v126
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v173 int64
	_ = v173
	var v177 int32
	_ = v177
	var v193 int32
	_ = v193
	var v196 int32
	_ = v196
	var v230 int32
	_ = v230
	var v233 int32
	_ = v233
	var v236 int32
	_ = v236
	var v255 int32
	_ = v255
	var v284 int32
	_ = v284
	var v324 int32
	_ = v324
	var v325 int32
	_ = v325
	var v328 int32
	_ = v328
	var __phi328 int32
	_ = __phi328
	var v329 int32
	_ = v329
	var __phi329 int32
	_ = __phi329
	var v331 int64
	_ = v331
	var v333 int32
	_ = v333
	var v336 int32
	_ = v336
	var v337 int32
	_ = v337
	var v342 int32
	_ = v342
	var v350 int32
	_ = v350
	var v351 int32
	_ = v351
	var v356 int32
	_ = v356
	var v357 int32
	_ = v357
	var v359 int32
	_ = v359
	var v362 int32
	_ = v362
	var v368 int32
	_ = v368
	var v371 int32
	_ = v371
	var v372 int32
	_ = v372
	var v373 int32
	_ = v373
	var v374 int32
	_ = v374
	var v382 int32
	_ = v382
	var v383 int32
	_ = v383
	var v384 int32
	_ = v384
	var v385 int32
	_ = v385
	var v386 int32
	_ = v386
	var v397 int32
	_ = v397
	var v398 int32
	_ = v398
	var v401 int32
	_ = v401
	var v402 int64
	_ = v402
	var v404 int32
	_ = v404
	var v406 int32
	_ = v406
	var v410 int32
	_ = v410
	var v414 int32
	_ = v414
	var v415 int32
	_ = v415
	if base.Ui32(l1+int32(-1)) < base.Ui32(int32(-10000)) {
		v17 = l1
	} else {
		v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
		v17 = l1 + (v9-v10)>>(uint(int32(4))%32) + int32(1)
	}
	if v17 < int32(1) {
		if v17 < int32(-9999) {
			switch v17 + int32(10002) {
			case 0:
				v71 = l0 + int32(72)
			case 1:
				v44 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
				v45 = *(*int32)(unsafe.Add(mBase, uint32(v44)+4))
				v46 = *(*int32)(unsafe.Add(mBase, uint32(v45)))
				v47 = *(*int32)(unsafe.Add(mBase, uint32(v46)+12))
				*(*int32)(unsafe.Add(mBase, uint32(l0)+96)) = int32(5)
				*(*int32)(unsafe.Add(mBase, uint32(l0)+88)) = v47
				v71 = l0 + int32(88)
			case 2:
				v41 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
				v71 = v41 + int32(96)
			default:
				v55 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
				v56 = *(*int32)(unsafe.Add(mBase, uint32(v55)+4))
				v57 = *(*int32)(unsafe.Add(mBase, uint32(v56)))
				v58 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v57)+7)))
				v59 = m.G398
				if base.Ui32(v58) < base.Ui32(int32(-10002)-v17) {
					v70 = v59
				} else {
					v70 = v57 + (int32(-10003)-v17)<<(uint(int32(4))%32) + int32(24)
				}
				v71 = v70
			}
		} else {
			v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			v71 = v35 + v17<<(uint(int32(4))%32)
		}
	} else {
		v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
		v29 = v24 + v17<<(uint(int32(4))%32) + int32(-16)
		v30 = m.G398
		if base.Ui32(v29) < base.Ui32(v23) {
			v32 = v29
		} else {
			v32 = v30
		}
		v71 = v32
	}
	v74 = *(*int32)(unsafe.Add(mBase, uint32(v71)+8))
	switch v74 + int32(-5) {
	case 0:
		v77 = *(*int32)(unsafe.Add(mBase, uint32(v71)))
		v89 = v77 + int32(16)
	default:
		v83 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
		v89 = v83 + v74<<(uint(int32(2))%32) + int32(152)
	case 2:
		v80 = *(*int32)(unsafe.Add(mBase, uint32(v71)))
		v89 = v80 + int32(8)
	}
	v90 = *(*int32)(unsafe.Add(mBase, uint32(v89)))
	if v90 != 0 {
		v92 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		*(*int32)(unsafe.Add(mBase, uint32(v92)+8)) = int32(5)
		*(*int32)(unsafe.Add(mBase, uint32(v92))) = v90
		v96 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v96 + int32(16)
		v102 = int32(1)
	} else {
		v102 = int32(0)
	}
	if v102 == int32(0) {
		v415 = int32(0)
		return v415
	} else {
		F_lua_pushstring(m, l0, l2)
		mBase = m.M
		v108 = m.ExcPending
		if v108 != 0 {
			return int32(0)
		} else {
			v126 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			v165 = *(*int32)(unsafe.Add(mBase, uint32(v126+int32(-32))))
			v166 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			v167 = int32(-16)
			v169 = F_luaH_get(m, v165, v166+v167)
			mBase = m.M
			v170 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			v173 = *(*int64)(unsafe.Add(mBase, uint32(v169)))
			*(*int64)(unsafe.Add(mBase, uint32(v170+v167))) = v173
			v177 = *(*int32)(unsafe.Add(mBase, uint32(v169)+8))
			*(*int32)(unsafe.Add(mBase, uint32(v170+int32(-8)))) = v177
			v193 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			v196 = v193 + int32(-16)
			v230 = m.G398
			if v196 != v230 {
				v233 = *(*int32)(unsafe.Add(mBase, uint32(v196)+8))
				v236 = v233
			} else {
				v236 = int32(-1)
			}
			if v236 != 0 {
				v284 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				v324 = v284 + int32(-16)
				v325 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				if base.Ui32(v325) <= base.Ui32(v324) {
					v342 = v325
				} else {
					__phi328 = v284 + int32(-32)
					__phi329 = v324
					v328 = __phi328
					v329 = __phi329
					for {
						v331 = *(*int64)(unsafe.Add(mBase, uint32(v328)+16))
						*(*int64)(unsafe.Add(mBase, uint32(v328))) = v331
						v333 = *(*int32)(unsafe.Add(mBase, uint32(v328)+24))
						*(*int32)(unsafe.Add(mBase, uint32(v328)+8)) = v333
						v336 = v329 + int32(16)
						v337 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
						if base.Ui32(v336) < base.Ui32(v337) {
							__phi328 = v329
							__phi329 = v336
							v328 = __phi328
							v329 = __phi329
							continue
						} else {
							break
						}
						break
					}
					v342 = v337
				}
				*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v342 + int32(-16)
				if v17 < int32(1) {
					if v17 < int32(-9999) {
						switch v17 + int32(10002) {
						case 0:
							v398 = l0 + int32(72)
						case 1:
							v371 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
							v372 = *(*int32)(unsafe.Add(mBase, uint32(v371)+4))
							v373 = *(*int32)(unsafe.Add(mBase, uint32(v372)))
							v374 = *(*int32)(unsafe.Add(mBase, uint32(v373)+12))
							*(*int32)(unsafe.Add(mBase, uint32(l0)+96)) = int32(5)
							*(*int32)(unsafe.Add(mBase, uint32(l0)+88)) = v374
							v398 = l0 + int32(88)
						case 2:
							v368 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
							v398 = v368 + int32(96)
						default:
							v382 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
							v383 = *(*int32)(unsafe.Add(mBase, uint32(v382)+4))
							v384 = *(*int32)(unsafe.Add(mBase, uint32(v383)))
							v385 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v384)+7)))
							v386 = m.G398
							if base.Ui32(v385) < base.Ui32(int32(-10002)-v17) {
								v397 = v386
							} else {
								v397 = v384 + (int32(-10003)-v17)<<(uint(int32(4))%32) + int32(24)
							}
							v398 = v397
						}
					} else {
						v362 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
						v398 = v362 + v17<<(uint(int32(4))%32)
					}
				} else {
					v350 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
					v351 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
					v356 = v351 + v17<<(uint(int32(4))%32) + int32(-16)
					v357 = m.G398
					if base.Ui32(v356) < base.Ui32(v350) {
						v359 = v356
					} else {
						v359 = v357
					}
					v398 = v359
				}
				v401 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				v402 = *(*int64)(unsafe.Add(mBase, uint32(v398)))
				*(*int64)(unsafe.Add(mBase, uint32(v401))) = v402
				v404 = *(*int32)(unsafe.Add(mBase, uint32(v398)+8))
				*(*int32)(unsafe.Add(mBase, uint32(v401)+8)) = v404
				v406 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v406 + int32(16)
				v410 = int32(1)
				F_lua_call(m, l0, v410, v410)
				mBase = m.M
				v414 = m.ExcPending
				if v414 != 0 {
					return int32(0)
				} else {
					v415 = v410
					return v415
				}
			} else {
				v255 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v255 + int32(-32)
				return int32(0)
			}
		}
	}
}
func F_luaL_checkinteger(m *base.Module, l0 int32, l1 int32) int32 {
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
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	var v31 int32
	_ = v31
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v42 int32
	_ = v42
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v74 int32
	_ = v74
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v92 int32
	_ = v92
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v100 int32
	_ = v100
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v10 = F_lua_tointeger(m, l0, l1)
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return int32(0)
	} else {
		if v10 != 0 {
			m.G0 = v8 + int32(16)
			return v10
		} else {
			v14 = F_lua_isnumber(m, l0, l1)
			mBase = m.M
			v15 = m.ExcPending
			if v15 != 0 {
				return int32(0)
			} else {
				if v14 != 0 {
					m.G0 = v8 + int32(16)
					return v10
				} else {
					v23 = m.G399
					v27 = *(*int32)(unsafe.Add(mBase, uint32(v23+int32(12))))
					if l1 < int32(1) {
						if l1 < int32(-9999) {
							switch l1 + int32(10002) {
							case 0:
								v78 = l0 + int32(72)
								v79 = m.G398
								if v78 != v79 {
									v82 = *(*int32)(unsafe.Add(mBase, uint32(v78)+8))
									v85 = v82
								} else {
									v85 = int32(-1)
								}
							case 1:
								v48 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
								v49 = *(*int32)(unsafe.Add(mBase, uint32(v48)+4))
								v50 = *(*int32)(unsafe.Add(mBase, uint32(v49)))
								v51 = *(*int32)(unsafe.Add(mBase, uint32(v50)+12))
								*(*int32)(unsafe.Add(mBase, uint32(l0)+96)) = int32(5)
								*(*int32)(unsafe.Add(mBase, uint32(l0)+88)) = v51
								v78 = l0 + int32(88)
								v79 = m.G398
								if v78 != v79 {
									v82 = *(*int32)(unsafe.Add(mBase, uint32(v78)+8))
									v85 = v82
								} else {
									v85 = int32(-1)
								}
							case 2:
								v74 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
								v78 = v74 + int32(96)
								v79 = m.G398
								if v78 != v79 {
									v82 = *(*int32)(unsafe.Add(mBase, uint32(v78)+8))
									v85 = v82
								} else {
									v85 = int32(-1)
								}
							default:
								v61 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
								v62 = *(*int32)(unsafe.Add(mBase, uint32(v61)+4))
								v63 = *(*int32)(unsafe.Add(mBase, uint32(v62)))
								v64 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v63)+7)))
								if base.Ui32(v64) < base.Ui32(int32(-10002)-l1) {
									v85 = int32(-1)
								} else {
									v78 = v63 + (int32(-10003)-l1)<<(uint(int32(4))%32) + int32(24)
									v79 = m.G398
									if v78 != v79 {
										v82 = *(*int32)(unsafe.Add(mBase, uint32(v78)+8))
										v85 = v82
									} else {
										v85 = int32(-1)
									}
								}
							}
						} else {
							v42 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
							v78 = v42 + l1<<(uint(int32(4))%32)
							v79 = m.G398
							if v78 != v79 {
								v82 = *(*int32)(unsafe.Add(mBase, uint32(v78)+8))
								v85 = v82
							} else {
								v85 = int32(-1)
							}
						}
					} else {
						v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
						v36 = v31 + l1<<(uint(int32(4))%32) + int32(-16)
						v37 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
						if base.Ui32(v36) < base.Ui32(v37) {
							v78 = v36
							v79 = m.G398
							if v78 != v79 {
								v82 = *(*int32)(unsafe.Add(mBase, uint32(v78)+8))
								v85 = v82
							} else {
								v85 = int32(-1)
							}
						} else {
							v85 = int32(-1)
						}
					}
					v87 = m.G3
					if v85 != int32(-1) {
						v92 = m.G399
						v96 = *(*int32)(unsafe.Add(mBase, uint32(v92+v85<<(uint(int32(2))%32))))
						v97 = v96
					} else {
						v97 = v87 + int32(_a2282)
					}
					*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = v97
					*(*int32)(unsafe.Add(mBase, uint32(v8))) = v27
					v100 = m.G3
					v103 = F_lua_pushfstring(m, l0, v100+int32(_a2283), v8)
					mBase = m.M
					v104 = m.ExcPending
					if v104 != 0 {
						return int32(0)
					} else {
						v105 = F_luaL_argerror(m, l0, l1, v103)
						mBase = m.M
						v106 = m.ExcPending
						if v106 != 0 {
							return int32(0)
						} else {
							m.G0 = v8 + int32(16)
							return v10
						}
					}
				}
			}
		}
	}
}
func F_luaL_checkstack(m *base.Module, l0 int32, l1 int32, l2 int32) {
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
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v9 = F_lua_checkstack(m, l0, l1)
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		return
	} else {
		if v9 != 0 {
			m.G0 = v7 + int32(16)
			return
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v7))) = l2
			v12 = m.G3
			v15 = F_luaL_error(m, l0, v12+int32(_a2281), v7)
			mBase = m.M
			v16 = m.ExcPending
			if v16 != 0 {
				return
			} else {
				m.G0 = v7 + int32(16)
				return
			}
		}
	}
}
func F_luaL_checktype(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v11 int32
	_ = v11
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
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
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v54 int32
	_ = v54
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v73 int32
	_ = v73
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v92 int32
	_ = v92
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v124 int32
	_ = v124
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v132 int32
	_ = v132
	var v135 int32
	_ = v135
	var v137 int32
	_ = v137
	var v142 int32
	_ = v142
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
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	if l1 < int32(1) {
		if l1 < int32(-9999) {
			switch l1 + int32(10002) {
			case 0:
				v58 = l0 + int32(72)
				v59 = m.G398
				if v58 != v59 {
					v62 = *(*int32)(unsafe.Add(mBase, uint32(v58)+8))
					v65 = v62
				} else {
					v65 = int32(-1)
				}
			case 1:
				v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
				v29 = *(*int32)(unsafe.Add(mBase, uint32(v28)+4))
				v30 = *(*int32)(unsafe.Add(mBase, uint32(v29)))
				v31 = *(*int32)(unsafe.Add(mBase, uint32(v30)+12))
				*(*int32)(unsafe.Add(mBase, uint32(l0)+96)) = int32(5)
				*(*int32)(unsafe.Add(mBase, uint32(l0)+88)) = v31
				v58 = l0 + int32(88)
				v59 = m.G398
				if v58 != v59 {
					v62 = *(*int32)(unsafe.Add(mBase, uint32(v58)+8))
					v65 = v62
				} else {
					v65 = int32(-1)
				}
			case 2:
				v54 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
				v58 = v54 + int32(96)
				v59 = m.G398
				if v58 != v59 {
					v62 = *(*int32)(unsafe.Add(mBase, uint32(v58)+8))
					v65 = v62
				} else {
					v65 = int32(-1)
				}
			default:
				v41 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
				v42 = *(*int32)(unsafe.Add(mBase, uint32(v41)+4))
				v43 = *(*int32)(unsafe.Add(mBase, uint32(v42)))
				v44 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v43)+7)))
				if base.Ui32(v44) < base.Ui32(int32(-10002)-l1) {
					v65 = int32(-1)
				} else {
					v58 = v43 + (int32(-10003)-l1)<<(uint(int32(4))%32) + int32(24)
					v59 = m.G398
					if v58 != v59 {
						v62 = *(*int32)(unsafe.Add(mBase, uint32(v58)+8))
						v65 = v62
					} else {
						v65 = int32(-1)
					}
				}
			}
		} else {
			v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			v58 = v22 + l1<<(uint(int32(4))%32)
			v59 = m.G398
			if v58 != v59 {
				v62 = *(*int32)(unsafe.Add(mBase, uint32(v58)+8))
				v65 = v62
			} else {
				v65 = int32(-1)
			}
		}
	} else {
		v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
		v16 = v11 + l1<<(uint(int32(4))%32) + int32(-16)
		v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		if base.Ui32(v16) < base.Ui32(v17) {
			v58 = v16
			v59 = m.G398
			if v58 != v59 {
				v62 = *(*int32)(unsafe.Add(mBase, uint32(v58)+8))
				v65 = v62
			} else {
				v65 = int32(-1)
			}
		} else {
			v65 = int32(-1)
		}
	}
	if v65 == l2 {
		m.G0 = v7 + int32(16)
		return
	} else {
		v68 = m.G3
		if l2 != int32(-1) {
			v73 = m.G399
			v77 = *(*int32)(unsafe.Add(mBase, uint32(v73+l2<<(uint(int32(2))%32))))
			v78 = v77
		} else {
			v78 = v68 + int32(_a2282)
		}
		if l1 < int32(1) {
			if l1 < int32(-9999) {
				switch l1 + int32(10002) {
				case 0:
					v128 = l0 + int32(72)
					v129 = m.G398
					if v128 != v129 {
						v132 = *(*int32)(unsafe.Add(mBase, uint32(v128)+8))
						v135 = v132
					} else {
						v135 = int32(-1)
					}
				case 1:
					v98 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
					v99 = *(*int32)(unsafe.Add(mBase, uint32(v98)+4))
					v100 = *(*int32)(unsafe.Add(mBase, uint32(v99)))
					v101 = *(*int32)(unsafe.Add(mBase, uint32(v100)+12))
					*(*int32)(unsafe.Add(mBase, uint32(l0)+96)) = int32(5)
					*(*int32)(unsafe.Add(mBase, uint32(l0)+88)) = v101
					v128 = l0 + int32(88)
					v129 = m.G398
					if v128 != v129 {
						v132 = *(*int32)(unsafe.Add(mBase, uint32(v128)+8))
						v135 = v132
					} else {
						v135 = int32(-1)
					}
				case 2:
					v124 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
					v128 = v124 + int32(96)
					v129 = m.G398
					if v128 != v129 {
						v132 = *(*int32)(unsafe.Add(mBase, uint32(v128)+8))
						v135 = v132
					} else {
						v135 = int32(-1)
					}
				default:
					v111 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
					v112 = *(*int32)(unsafe.Add(mBase, uint32(v111)+4))
					v113 = *(*int32)(unsafe.Add(mBase, uint32(v112)))
					v114 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v113)+7)))
					if base.Ui32(v114) < base.Ui32(int32(-10002)-l1) {
						v135 = int32(-1)
					} else {
						v128 = v113 + (int32(-10003)-l1)<<(uint(int32(4))%32) + int32(24)
						v129 = m.G398
						if v128 != v129 {
							v132 = *(*int32)(unsafe.Add(mBase, uint32(v128)+8))
							v135 = v132
						} else {
							v135 = int32(-1)
						}
					}
				}
			} else {
				v92 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				v128 = v92 + l1<<(uint(int32(4))%32)
				v129 = m.G398
				if v128 != v129 {
					v132 = *(*int32)(unsafe.Add(mBase, uint32(v128)+8))
					v135 = v132
				} else {
					v135 = int32(-1)
				}
			}
		} else {
			v81 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
			v86 = v81 + l1<<(uint(int32(4))%32) + int32(-16)
			v87 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			if base.Ui32(v86) < base.Ui32(v87) {
				v128 = v86
				v129 = m.G398
				if v128 != v129 {
					v132 = *(*int32)(unsafe.Add(mBase, uint32(v128)+8))
					v135 = v132
				} else {
					v135 = int32(-1)
				}
			} else {
				v135 = int32(-1)
			}
		}
		v137 = m.G3
		if v135 != int32(-1) {
			v142 = m.G399
			v146 = *(*int32)(unsafe.Add(mBase, uint32(v142+v135<<(uint(int32(2))%32))))
			v147 = v146
		} else {
			v147 = v137 + int32(_a2282)
		}
		*(*int32)(unsafe.Add(mBase, uint32(v7)+4)) = v147
		*(*int32)(unsafe.Add(mBase, uint32(v7))) = v78
		v150 = m.G3
		v153 = F_lua_pushfstring(m, l0, v150+int32(_a2283), v7)
		mBase = m.M
		v154 = m.ExcPending
		if v154 != 0 {
			return
		} else {
			v155 = F_luaL_argerror(m, l0, l1, v153)
			mBase = m.M
			v156 = m.ExcPending
			if v156 != 0 {
				return
			} else {
				m.G0 = v7 + int32(16)
				return
			}
		}
	}
}
func F_luaL_loadbuffer(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v12 int32
	_ = v12
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	*(*int32)(unsafe.Add(mBase, uint32(v8)+12)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(v8)+8)) = l1
	v12 = m.G5
	v17 = F_lua_load(m, l0, v12+int32(1235), v8+int32(8), l3)
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		return int32(0)
	} else {
		m.G0 = v8 + int32(16)
		return v17
	}
}
func F_luaL_loadfile(m *base.Module, l0 int32, l1 int32) int32 {
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
	var v16 int32
	_ = v16
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v100 int32
	_ = v100
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v131 int32
	_ = v131
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
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v165 int32
	_ = v165
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v174 int32
	_ = v174
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v193 int32
	_ = v193
	var v197 int32
	_ = v197
	var v201 int32
	_ = v201
	var v203 int32
	_ = v203
	var v210 int32
	_ = v210
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v220 int32
	_ = v220
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v250 int32
	_ = v250
	var v251 int32
	_ = v251
	var v253 int32
	_ = v253
	var v256 int32
	_ = v256
	var v262 int32
	_ = v262
	var v265 int32
	_ = v265
	var v266 int32
	_ = v266
	var v267 int32
	_ = v267
	var v268 int32
	_ = v268
	var v276 int32
	_ = v276
	var v277 int32
	_ = v277
	var v278 int32
	_ = v278
	var v279 int32
	_ = v279
	var v280 int32
	_ = v280
	var v291 int32
	_ = v291
	var v292 int32
	_ = v292
	var v296 int32
	_ = v296
	var v297 int32
	_ = v297
	var v300 int32
	_ = v300
	var __phi300 int32
	_ = __phi300
	var v301 int32
	_ = v301
	var __phi301 int32
	_ = __phi301
	var v303 int64
	_ = v303
	var v305 int32
	_ = v305
	var v308 int32
	_ = v308
	var v309 int32
	_ = v309
	var v314 int32
	_ = v314
	var v323 int32
	_ = v323
	var v324 int32
	_ = v324
	var v329 int32
	_ = v329
	var v330 int32
	_ = v330
	var v332 int32
	_ = v332
	var v335 int32
	_ = v335
	var v341 int32
	_ = v341
	var v344 int32
	_ = v344
	var v345 int32
	_ = v345
	var v346 int32
	_ = v346
	var v347 int32
	_ = v347
	var v355 int32
	_ = v355
	var v356 int32
	_ = v356
	var v357 int32
	_ = v357
	var v358 int32
	_ = v358
	var v359 int32
	_ = v359
	var v370 int32
	_ = v370
	var v371 int32
	_ = v371
	var v375 int32
	_ = v375
	var v376 int32
	_ = v376
	var v379 int32
	_ = v379
	var __phi379 int32
	_ = __phi379
	var v380 int32
	_ = v380
	var __phi380 int32
	_ = __phi380
	var v382 int64
	_ = v382
	var v384 int32
	_ = v384
	var v387 int32
	_ = v387
	var v388 int32
	_ = v388
	var v393 int32
	_ = v393
	var v400 int32
	_ = v400
	v8 = m.G0
	v10 = v8 - int32(1088)
	m.G0 = v10
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v16 = (v12 - v13) >> (uint(int32(4)) % 32)
	goto L1
L1:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+56)) = int32(0)
	v20 = v16 + int32(1)
	if l1 != 0 {
		goto L6
	} else {
		goto L7
	}
L2:
	;
	m.G0 = v10 + int32(1088)
	return v400
L3:
	;
	if v20 < int32(1) {
		goto L88
	} else {
		goto L89
	}
L4:
	;
	if v20 < int32(1) {
		goto L67
	} else {
		goto L68
	}
L5:
	;
	v65 = F_getc(m, v64)
	mBase = m.M
	v66 = m.ExcPending
	if v66 != 0 {
		goto L8
	} else {
		goto L18
	}
L6:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+48)) = l1
	v33 = m.G3
	v38 = F_lua_pushfstring(m, l0, v33+int32(_a2284), v10+int32(48))
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L8
	} else {
		goto L10
	}
L7:
	;
	v21 = m.G3
	F_lua_pushlstring(m, l0, v21+int32(_a2285), int32(6))
	mBase = m.M
	v28 = m.ExcPending
	if v28 != 0 {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	return int32(0)
L9:
	;
	v29 = m.G402
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v29)))
	*(*int32)(unsafe.Add(mBase, uint32(v10)+60)) = v30
	v64 = v30
	goto L5
L10:
	;
	v42 = F_fopen(m, l1, v33+int32(_a178))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v10)+60)) = v42
	if v42 != 0 {
		v64 = v42
		goto L5
	} else {
		goto L11
	}
L11:
	;
	goto L12
L12:
	;
	v45 = *(*int32)(unsafe.Add(mBase, _consts[9]))
	v46 = F___strerror_l(m, v45, v45)
	mBase = m.M
	goto L13
L13:
	;
	v48 = F_lua_tolstring(m, l0, v20, int32(0))
	mBase = m.M
	v49 = m.ExcPending
	if v49 != 0 {
		goto L8
	} else {
		goto L14
	}
L14:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+40)) = v46
	v51 = m.G3
	*(*int32)(unsafe.Add(mBase, uint32(v10)+32)) = v51 + int32(_a2286)
	*(*int32)(unsafe.Add(mBase, uint32(v10)+36)) = v48 + int32(1)
	v62 = F_lua_pushfstring(m, l0, v51+int32(_a2287), v10+int32(32))
	mBase = m.M
	v63 = m.ExcPending
	if v63 != 0 {
		goto L8
	} else {
		goto L15
	}
L15:
	;
	goto L4
L16:
	;
	v149 = F_ungetc(m, v147, v145)
	mBase = m.M
	v150 = m.ExcPending
	if v150 != 0 {
		goto L8
	} else {
		goto L41
	}
L17:
	;
	v94 = int32(0)
	v95 = base.B2i32(l1 != v94)
	if l1 == v94 {
		v145 = v64
		v147 = v92
		v148 = v95
		goto L16
	} else {
		goto L26
	}
L18:
	;
	if v65 != int32(35) {
		v92 = v65
		goto L17
	} else {
		goto L19
	}
L19:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+56)) = int32(1)
	goto L22
L20:
	;
	v85 = F_getc(m, v64)
	mBase = m.M
	v86 = m.ExcPending
	if v86 != 0 {
		goto L8
	} else {
		goto L25
	}
L21:
	;
	v145 = v64
	v147 = int32(-1)
	v148 = base.B2i32(l1 != int32(0))
	goto L16
L22:
	;
	v78 = F_getc(m, v64)
	mBase = m.M
	v79 = m.ExcPending
	if v79 != 0 {
		goto L8
	} else {
		goto L24
	}
L24:
	;
	switch v78 + int32(1) {
	case 0:
		goto L21
	default:
		goto L22
	case 11:
		goto L20
	}
L25:
	;
	v92 = v85
	goto L17
L26:
	;
	if v92 != int32(27) {
		v145 = v64
		v147 = v92
		v148 = v95
		goto L16
	} else {
		goto L27
	}
L27:
	;
	v100 = m.G3
	v103 = F_freopen(m, l1, v100+int32(_a2288), v64)
	mBase = m.M
	v104 = m.ExcPending
	if v104 != 0 {
		goto L8
	} else {
		goto L28
	}
L28:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+60)) = v103
	if v103 == int32(0) {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	goto L37
L30:
	;
	goto L32
L31:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+56)) = int32(0)
	v145 = v103
	v147 = v115
	v148 = int32(1)
	goto L16
L32:
	;
	v115 = F_getc(m, v103)
	mBase = m.M
	v116 = m.ExcPending
	if v116 != 0 {
		goto L8
	} else {
		goto L34
	}
L33:
	;
	goto L31
L34:
	;
	if v115 == int32(27) {
		goto L31
	} else {
		goto L35
	}
L35:
	;
	if v115 != int32(-1) {
		goto L32
	} else {
		goto L36
	}
L36:
	;
	goto L33
L37:
	;
	v125 = *(*int32)(unsafe.Add(mBase, _consts[9]))
	v126 = F___strerror_l(m, v125, v125)
	mBase = m.M
	goto L38
L38:
	;
	v128 = F_lua_tolstring(m, l0, v20, int32(0))
	mBase = m.M
	v129 = m.ExcPending
	if v129 != 0 {
		goto L8
	} else {
		goto L39
	}
L39:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+8)) = v126
	v131 = m.G3
	*(*int32)(unsafe.Add(mBase, uint32(v10))) = v131 + int32(_a2289)
	*(*int32)(unsafe.Add(mBase, uint32(v10)+4)) = v128 + int32(1)
	v140 = F_lua_pushfstring(m, l0, v131+int32(_a2287), v10)
	mBase = m.M
	v141 = m.ExcPending
	if v141 != 0 {
		goto L8
	} else {
		goto L40
	}
L40:
	;
	goto L4
L41:
	;
	v151 = m.G5
	v158 = F_lua_tolstring(m, l0, int32(-1), int32(0))
	mBase = m.M
	v159 = m.ExcPending
	if v159 != 0 {
		goto L8
	} else {
		goto L42
	}
L42:
	;
	v160 = F_lua_load(m, l0, v151+int32(1234), v10+int32(56), v158)
	mBase = m.M
	v161 = m.ExcPending
	if v161 != 0 {
		goto L8
	} else {
		goto L43
	}
L43:
	;
	v162 = *(*int32)(unsafe.Add(mBase, uint32(v10)+60))
	v165 = *(*int32)(unsafe.Add(mBase, uint32(v162)+76))
	if int32(-1) < v165 {
		goto L46
	} else {
		goto L47
	}
L44:
	;
	if v148 == int32(0) {
		goto L49
	} else {
		goto L50
	}
L45:
	;
	goto L44
L46:
	;
	v169 = F___lockfile(m, v162)
	mBase = m.M
	v170 = *(*int32)(unsafe.Add(mBase, uint32(v162)))
	if v169 == int32(0) {
		v174 = v170
		goto L45
	} else {
		goto L48
	}
L47:
	;
	v168 = *(*int32)(unsafe.Add(mBase, uint32(v162)))
	v174 = v168
	goto L45
L48:
	;
	F___unlockfile(m, v162)
	mBase = m.M
	v174 = v170
	goto L45
L49:
	;
	if int32(base.Ui32(v174)>>(uint(int32(5))%32))&int32(1) == int32(0) {
		goto L3
	} else {
		goto L52
	}
L50:
	;
	v182 = F_fclose(m, v162)
	mBase = m.M
	v183 = m.ExcPending
	if v183 != 0 {
		goto L8
	} else {
		goto L51
	}
L51:
	;
	goto L49
L52:
	;
	if v20 < int32(0) {
		goto L55
	} else {
		goto L56
	}
L53:
	;
	goto L61
L54:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v210
	goto L53
L55:
	;
	v203 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v210 = v203 + v20<<(uint(int32(4))%32) + int32(16)
	goto L54
L56:
	;
	v189 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v190 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v193 = v190 + v20<<(uint(int32(4))%32)
	if base.Ui32(v193) <= base.Ui32(v189) {
		v210 = v193
		goto L54
	} else {
		goto L57
	}
L57:
	;
	v197 = v189
	goto L58
L58:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v197)+8)) = int32(0)
	v201 = v197 + int32(16)
	if base.Ui32(v201) < base.Ui32(v193) {
		v197 = v201
		goto L58
	} else {
		goto L60
	}
L60:
	;
	v210 = v193
	goto L54
L61:
	;
	v214 = *(*int32)(unsafe.Add(mBase, _consts[9]))
	v215 = F___strerror_l(m, v214, v214)
	mBase = m.M
	goto L62
L62:
	;
	v217 = F_lua_tolstring(m, l0, v20, int32(0))
	mBase = m.M
	v218 = m.ExcPending
	if v218 != 0 {
		goto L8
	} else {
		goto L63
	}
L63:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+24)) = v215
	v220 = m.G3
	*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = v220 + int32(_a2290)
	*(*int32)(unsafe.Add(mBase, uint32(v10)+20)) = v217 + int32(1)
	v231 = F_lua_pushfstring(m, l0, v220+int32(_a2287), v10+int32(16))
	mBase = m.M
	v232 = m.ExcPending
	if v232 != 0 {
		goto L8
	} else {
		goto L64
	}
L64:
	;
	goto L4
L65:
	;
	v400 = int32(6)
	goto L2
L66:
	;
	v296 = v292 + int32(16)
	v297 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if base.Ui32(v297) <= base.Ui32(v296) {
		v314 = v297
		goto L81
	} else {
		goto L82
	}
L67:
	;
	if v20 < int32(-9999) {
		goto L72
	} else {
		goto L73
	}
L68:
	;
	v244 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v245 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v250 = v245 + v20<<(uint(int32(4))%32) + int32(-16)
	v251 = m.G398
	if base.Ui32(v250) < base.Ui32(v244) {
		goto L69
	} else {
		goto L70
	}
L69:
	;
	v253 = v250
	goto L71
L70:
	;
	v253 = v251
	goto L71
L71:
	;
	v292 = v253
	goto L66
L72:
	;
	switch v16 + int32(10003) {
	case 0:
		goto L75
	case 1:
		goto L76
	case 2:
		goto L77
	default:
		goto L74
	}
L73:
	;
	v256 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v292 = v256 + v20<<(uint(int32(4))%32)
	goto L66
L74:
	;
	v276 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v277 = *(*int32)(unsafe.Add(mBase, uint32(v276)+4))
	v278 = *(*int32)(unsafe.Add(mBase, uint32(v277)))
	v279 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v278)+7)))
	v280 = m.G398
	if base.Ui32(v279) < base.Ui32(int32(-10002)-v20) {
		goto L78
	} else {
		goto L79
	}
L75:
	;
	v292 = l0 + int32(72)
	goto L66
L76:
	;
	v265 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v266 = *(*int32)(unsafe.Add(mBase, uint32(v265)+4))
	v267 = *(*int32)(unsafe.Add(mBase, uint32(v266)))
	v268 = *(*int32)(unsafe.Add(mBase, uint32(v267)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+96)) = int32(5)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+88)) = v268
	v292 = l0 + int32(88)
	goto L66
L77:
	;
	v262 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v292 = v262 + int32(96)
	goto L66
L78:
	;
	v291 = v280
	goto L80
L79:
	;
	v291 = v278 + (int32(-10003)-v20)<<(uint(int32(4))%32) + int32(24)
	goto L80
L80:
	;
	v292 = v291
	goto L66
L81:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v314 + int32(-16)
	goto L65
L82:
	;
	__phi300 = v292
	__phi301 = v296
	v300 = __phi300
	v301 = __phi301
	goto L83
L83:
	;
	v303 = *(*int64)(unsafe.Add(mBase, uint32(v300)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v300))) = v303
	v305 = *(*int32)(unsafe.Add(mBase, uint32(v300)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v300)+8)) = v305
	v308 = v301 + int32(16)
	v309 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if base.Ui32(v308) < base.Ui32(v309) {
		__phi300 = v301
		__phi301 = v308
		v300 = __phi300
		v301 = __phi301
		goto L83
	} else {
		goto L85
	}
L84:
	;
	v314 = v309
	goto L81
L85:
	;
	goto L84
L86:
	;
	v400 = v160
	goto L2
L87:
	;
	v375 = v371 + int32(16)
	v376 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if base.Ui32(v376) <= base.Ui32(v375) {
		v393 = v376
		goto L102
	} else {
		goto L103
	}
L88:
	;
	if v20 < int32(-9999) {
		goto L93
	} else {
		goto L94
	}
L89:
	;
	v323 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v324 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v329 = v324 + v20<<(uint(int32(4))%32) + int32(-16)
	v330 = m.G398
	if base.Ui32(v329) < base.Ui32(v323) {
		goto L90
	} else {
		goto L91
	}
L90:
	;
	v332 = v329
	goto L92
L91:
	;
	v332 = v330
	goto L92
L92:
	;
	v371 = v332
	goto L87
L93:
	;
	switch v16 + int32(10003) {
	case 0:
		goto L96
	case 1:
		goto L97
	case 2:
		goto L98
	default:
		goto L95
	}
L94:
	;
	v335 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v371 = v335 + v20<<(uint(int32(4))%32)
	goto L87
L95:
	;
	v355 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v356 = *(*int32)(unsafe.Add(mBase, uint32(v355)+4))
	v357 = *(*int32)(unsafe.Add(mBase, uint32(v356)))
	v358 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v357)+7)))
	v359 = m.G398
	if base.Ui32(v358) < base.Ui32(int32(-10002)-v20) {
		goto L99
	} else {
		goto L100
	}
L96:
	;
	v371 = l0 + int32(72)
	goto L87
L97:
	;
	v344 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v345 = *(*int32)(unsafe.Add(mBase, uint32(v344)+4))
	v346 = *(*int32)(unsafe.Add(mBase, uint32(v345)))
	v347 = *(*int32)(unsafe.Add(mBase, uint32(v346)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+96)) = int32(5)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+88)) = v347
	v371 = l0 + int32(88)
	goto L87
L98:
	;
	v341 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v371 = v341 + int32(96)
	goto L87
L99:
	;
	v370 = v359
	goto L101
L100:
	;
	v370 = v357 + (int32(-10003)-v20)<<(uint(int32(4))%32) + int32(24)
	goto L101
L101:
	;
	v371 = v370
	goto L87
L102:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v393 + int32(-16)
	goto L86
L103:
	;
	__phi379 = v371
	__phi380 = v375
	v379 = __phi379
	v380 = __phi380
	goto L104
L104:
	;
	v382 = *(*int64)(unsafe.Add(mBase, uint32(v379)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v379))) = v382
	v384 = *(*int32)(unsafe.Add(mBase, uint32(v379)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v379)+8)) = v384
	v387 = v380 + int32(16)
	v388 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if base.Ui32(v387) < base.Ui32(v388) {
		__phi379 = v380
		__phi380 = v387
		v379 = __phi379
		v380 = __phi380
		goto L104
	} else {
		goto L106
	}
L105:
	;
	v393 = v388
	goto L102
L106:
	;
	goto L105
}
func F_luaL_newstate(m *base.Module) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v16 int32
	_ = v16
	v2 = m.G5
	v6 = F_lua_newstate(m, v2+int32(1236), int32(0))
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		if v6 == int32(0) {
		} else {
			v12 = m.G5
			v16 = *(*int32)(unsafe.Add(mBase, uint32(v6)+16))
			*(*int32)(unsafe.Add(mBase, uint32(v16)+88)) = v12 + int32(1237)
		}
		return v6
	}
}
func F_luaL_optinteger(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v17 int32
	_ = v17
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v49 int32
	_ = v49
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	if l1 < int32(1) {
		if l1 < int32(-9999) {
			switch l1 + int32(10002) {
			case 0:
				v53 = l0 + int32(72)
				v54 = m.G398
				if v53 != v54 {
					v57 = *(*int32)(unsafe.Add(mBase, uint32(v53)+8))
					v60 = v57
				} else {
					v60 = int32(-1)
				}
			case 1:
				v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
				v24 = *(*int32)(unsafe.Add(mBase, uint32(v23)+4))
				v25 = *(*int32)(unsafe.Add(mBase, uint32(v24)))
				v26 = *(*int32)(unsafe.Add(mBase, uint32(v25)+12))
				*(*int32)(unsafe.Add(mBase, uint32(l0)+96)) = int32(5)
				*(*int32)(unsafe.Add(mBase, uint32(l0)+88)) = v26
				v53 = l0 + int32(88)
				v54 = m.G398
				if v53 != v54 {
					v57 = *(*int32)(unsafe.Add(mBase, uint32(v53)+8))
					v60 = v57
				} else {
					v60 = int32(-1)
				}
			case 2:
				v49 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
				v53 = v49 + int32(96)
				v54 = m.G398
				if v53 != v54 {
					v57 = *(*int32)(unsafe.Add(mBase, uint32(v53)+8))
					v60 = v57
				} else {
					v60 = int32(-1)
				}
			default:
				v36 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
				v37 = *(*int32)(unsafe.Add(mBase, uint32(v36)+4))
				v38 = *(*int32)(unsafe.Add(mBase, uint32(v37)))
				v39 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v38)+7)))
				if base.Ui32(v39) < base.Ui32(int32(-10002)-l1) {
					v60 = int32(-1)
				} else {
					v53 = v38 + (int32(-10003)-l1)<<(uint(int32(4))%32) + int32(24)
					v54 = m.G398
					if v53 != v54 {
						v57 = *(*int32)(unsafe.Add(mBase, uint32(v53)+8))
						v60 = v57
					} else {
						v60 = int32(-1)
					}
				}
			}
		} else {
			v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			v53 = v17 + l1<<(uint(int32(4))%32)
			v54 = m.G398
			if v53 != v54 {
				v57 = *(*int32)(unsafe.Add(mBase, uint32(v53)+8))
				v60 = v57
			} else {
				v60 = int32(-1)
			}
		}
	} else {
		v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
		v11 = v6 + l1<<(uint(int32(4))%32) + int32(-16)
		v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		if base.Ui32(v11) < base.Ui32(v12) {
			v53 = v11
			v54 = m.G398
			if v53 != v54 {
				v57 = *(*int32)(unsafe.Add(mBase, uint32(v53)+8))
				v60 = v57
			} else {
				v60 = int32(-1)
			}
		} else {
			v60 = int32(-1)
		}
	}
	if v60 < int32(1) {
		v67 = l2
		return v67
	} else {
		v63 = F_luaL_checkinteger(m, l0, l1)
		mBase = m.M
		v66 = m.ExcPending
		if v66 != 0 {
			return int32(0)
		} else {
			v67 = v63
			return v67
		}
	}
}
func F_luaL_optlstring(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v18 int32
	_ = v18
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v50 int32
	_ = v50
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v83 int32
	_ = v83
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v94 int32
	_ = v94
	var v100 int32
	_ = v100
	var v103 int32
	_ = v103
	var v109 int32
	_ = v109
	var v113 int32
	_ = v113
	var v115 int32
	_ = v115
	var v123 int32
	_ = v123
	var v126 int32
	_ = v126
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	if l1 < int32(1) {
		goto L5
	} else {
		goto L6
	}
L1:
	;
	return v130
L2:
	;
	v126 = F_luaL_checklstring(m, l0, l1, l3)
	mBase = m.M
	v129 = m.ExcPending
	if v129 != 0 {
		goto L38
	} else {
		goto L39
	}
L3:
	;
	if int32(0) < v61 {
		goto L2
	} else {
		goto L18
	}
L4:
	;
	v55 = m.G398
	if v54 != v55 {
		goto L16
	} else {
		goto L17
	}
L5:
	;
	if l1 < int32(-9999) {
		goto L8
	} else {
		goto L9
	}
L6:
	;
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v12 = v7 + l1<<(uint(int32(4))%32) + int32(-16)
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if base.Ui32(v12) < base.Ui32(v13) {
		v54 = v12
		goto L4
	} else {
		goto L7
	}
L7:
	;
	v61 = int32(-1)
	goto L3
L8:
	;
	switch l1 + int32(10002) {
	case 0:
		goto L12
	case 1:
		goto L13
	case 2:
		goto L10
	default:
		goto L11
	}
L9:
	;
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v54 = v18 + l1<<(uint(int32(4))%32)
	goto L4
L10:
	;
	v50 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v54 = v50 + int32(96)
	goto L4
L11:
	;
	v37 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v37)+4))
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v38)))
	v40 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v39)+7)))
	if base.Ui32(v40) < base.Ui32(int32(-10002)-l1) {
		goto L14
	} else {
		goto L15
	}
L12:
	;
	v54 = l0 + int32(72)
	goto L4
L13:
	;
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v24)+4))
	v26 = *(*int32)(unsafe.Add(mBase, uint32(v25)))
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v26)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+96)) = int32(5)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+88)) = v27
	v54 = l0 + int32(88)
	goto L4
L14:
	;
	v61 = int32(-1)
	goto L3
L15:
	;
	v54 = v39 + (int32(-10003)-l1)<<(uint(int32(4))%32) + int32(24)
	goto L4
L16:
	;
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v54)+8))
	v61 = v58
	goto L3
L17:
	;
	v61 = int32(-1)
	goto L3
L18:
	;
	if l3 == int32(0) {
		v130 = l2
		goto L1
	} else {
		goto L19
	}
L19:
	;
	if l2 != 0 {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	if l2&int32(3) == int32(0) {
		v90 = l2
		goto L24
	} else {
		goto L25
	}
L21:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = int32(0)
	return l2
L22:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v123
	return l2
L23:
	;
	v123 = v115 - l2
	goto L22
L24:
	;
	v94 = v90
	goto L32
L25:
	;
	v76 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2))))
	if v76 != 0 {
		goto L26
	} else {
		goto L27
	}
L26:
	;
	v79 = l2
	goto L28
L27:
	;
	v123 = l2 - l2
	goto L22
L28:
	;
	v83 = v79 + int32(1)
	if v83&int32(3) == int32(0) {
		v90 = v83
		goto L24
	} else {
		goto L30
	}
L30:
	;
	v88 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v83))))
	if v88 != 0 {
		v79 = v83
		goto L28
	} else {
		goto L31
	}
L31:
	;
	v115 = v83
	goto L23
L32:
	;
	v100 = *(*int32)(unsafe.Add(mBase, uint32(v94)))
	v103 = int32(-2139062144)
	if (int32(16843008)-v100|v100)&v103 == v103 {
		v94 = v94 + int32(4)
		goto L32
	} else {
		goto L34
	}
L33:
	;
	v109 = v94
	goto L35
L34:
	;
	goto L33
L35:
	;
	v113 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v109))))
	if v113 != 0 {
		v109 = v109 + int32(1)
		goto L35
	} else {
		goto L37
	}
L36:
	;
	v115 = v109
	goto L23
L37:
	;
	goto L36
L38:
	;
	return int32(0)
L39:
	;
	v130 = v126
	goto L1
}
func F_luaL_where(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
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
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v88 int32
	_ = v88
	v4 = m.G0
	v6 = v4 - int32(112)
	m.G0 = v6
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if l1 < int32(1) {
		v38 = l1
		v40 = v13
		goto L5
	} else {
		goto L6
	}
L1:
	;
	m.G0 = v6 + int32(112)
	return
L2:
	;
	v83 = m.G3
	F_lua_pushlstring(m, l0, v83+int32(_a188), int32(0))
	mBase = m.M
	v88 = m.ExcPending
	if v88 != 0 {
		goto L17
	} else {
		goto L21
	}
L3:
	;
	if v60 == int32(0) {
		goto L2
	} else {
		goto L16
	}
L4:
	;
	goto L3
L5:
	;
	if v38 != 0 {
		v51 = int32(0)
		goto L13
	} else {
		goto L14
	}
L6:
	;
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v19 = l1
	v21 = v13
	goto L7
L7:
	;
	if base.Ui32(v21) <= base.Ui32(v16) {
		v60 = int32(0)
		goto L4
	} else {
		goto L9
	}
L8:
	;
	v38 = v32
	v40 = v34
	goto L5
L9:
	;
	v26 = v19 + int32(-1)
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v21)+4))
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v27)))
	v29 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28)+6)))
	if v29 != 0 {
		v32 = v26
		goto L10
	} else {
		goto L11
	}
L10:
	;
	v34 = v21 + int32(-24)
	if int32(0) < v32 {
		v19 = v32
		v21 = v34
		goto L7
	} else {
		goto L12
	}
L11:
	;
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v21)+20))
	v32 = v26 - v30
	goto L10
L12:
	;
	goto L8
L13:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v6+int32(12))+96)) = v51
	v60 = int32(1)
	goto L4
L14:
	;
	v45 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	if base.Ui32(v40) <= base.Ui32(v45) {
		v60 = int32(0)
		goto L4
	} else {
		goto L15
	}
L15:
	;
	v49 = base.I32_div_s(v40-v45, int32(24))
	v51 = v49
	goto L13
L16:
	;
	v63 = m.G3
	v68 = F_lua_getinfo(m, l0, v63+int32(_a2279), v6+int32(12))
	mBase = m.M
	v69 = m.ExcPending
	if v69 != 0 {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	return
L18:
	;
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v6)+32))
	if v70 < int32(1) {
		goto L2
	} else {
		goto L19
	}
L19:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v6)+4)) = v70
	*(*int32)(unsafe.Add(mBase, uint32(v6))) = v6 + int32(48)
	v77 = m.G3
	v80 = F_lua_pushfstring(m, l0, v77+int32(_a2280), v6)
	mBase = m.M
	v81 = m.ExcPending
	if v81 != 0 {
		goto L17
	} else {
		goto L20
	}
L20:
	;
	goto L1
L21:
	;
	goto L1
}
