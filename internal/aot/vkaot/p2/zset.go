package p2

import (
	base "github.com/shibukawa/valkeymem/internal/aot/vkaot/base"
	"unsafe"
)

func F_createZsetObject(m *base.Module) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v30 int32
	_ = v30
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	v9 = F_valkey_malloc(m, int32(8))
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		return int32(0)
	} else {
		v14 = F_hashtableCreate(m, int32(_a1057))
		mBase = m.M
		v15 = m.ExcPending
		if v15 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v9))) = v14
			v17 = F_zslCreate(m)
			mBase = m.M
			v18 = m.ExcPending
			if v18 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v9)+4)) = v17
				*(*int32)(unsafe.Add(mBase, uint32(v6)+12)) = int32(0)
				v22 = int32(12)
				v25 = F_zmalloc_usable(m, v22, v6+v22)
				mBase = m.M
				v26 = m.ExcPending
				if v26 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v25)+8)) = v9
					*(*int64)(unsafe.Add(mBase, uint32(v25))) = int64(34359738371)
					v30 = *(*int32)(unsafe.Add(mBase, uint32(v25)))
					*(*int32)(unsafe.Add(mBase, uint32(v25))) = v30&int32(-241) | int32(112)
					m.G0 = v6 + int32(16)
					return v25
				}
			}
		}
	}
}
func F_freeZsetObject(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v12 int32
	_ = v12
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v32 int32
	_ = v32
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v53 int32
	_ = v53
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v104 int32
	_ = v104
	var v111 int32
	_ = v111
	var v114 int32
	_ = v114
	var v117 int32
	_ = v117
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v125 int32
	_ = v125
	var v140 int32
	_ = v140
	var v144 int32
	_ = v144
	var v150 int32
	_ = v150
	var v152 int32
	_ = v152
	var v154 int32
	_ = v154
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	switch int32(base.Ui32(v5)>>(uint(int32(4))%32))&int32(15) + int32(-7) {
	case 0:
		v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		if v12&int32(4) == int32(0) {
			v71 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			v72 = v71
		} else {
			if v12&int32(1) != 0 {
				v21 = int32(16)
			} else {
				v21 = int32(8)
			}
			v22 = l0 + v21
			if v12&int32(2) == int32(0) {
				v53 = v22
			} else {
				v27 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22))))
				v28 = v22 + v27
				v32 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28))))
				switch v32 & int32(7) {
				case 0:
					v49 = int32(base.Ui32(v32) >> (uint(int32(3)) % 32))
				case 1:
					v39 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28+int32(-2)))))
					v49 = v39
				case 2:
					v42 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v28+int32(-4)))))
					v49 = v42
				case 3:
					v45 = *(*int32)(unsafe.Add(mBase, uint32(v28+int32(-8))))
					v49 = v45
				case 4:
					v48 = *(*int32)(unsafe.Add(mBase, uint32(v28+int32(-16))))
					v49 = v48
				default:
					v49 = int32(0)
				}
				v53 = v28 + int32(1) + v49 + int32(1)
			}
			v68 = *(*int32)(unsafe.Add(mBase, _consts[601]))
			v72 = v53 + v68
		}
		v76 = *(*int32)(unsafe.Add(mBase, uint32(v72)))
		F_hashtableRelease(m, v76)
		mBase = m.M
		v78 = m.ExcPending
		if v78 != 0 {
			return
		} else {
			v79 = *(*int32)(unsafe.Add(mBase, uint32(v72)+4))
			F_zslFree(m, v79)
			mBase = m.M
			v81 = m.ExcPending
			if v81 != 0 {
				return
			} else {
				F_valkey_free(m, v72)
				mBase = m.M
				v83 = m.ExcPending
				if v83 != 0 {
					return
				} else {
					return
				}
			}
		}
	default:
		F__serverPanic_1(m, int32(_a1051), int32(590), int32(_a1054), int32(0))
		mBase = m.M
		v150 = m.ExcPending
		if v150 != 0 {
			return
		} else {
			F_abort(m)
			mBase = m.M
			base.Wasm_trap_unreachable()
			for {
			}
		}
	case 4:
		v84 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		if v84&int32(4) == int32(0) {
			v152 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			F_valkey_free(m, v152)
			mBase = m.M
			v154 = m.ExcPending
			if v154 != 0 {
				return
			} else {
				return
			}
		} else {
			if v84&int32(1) != 0 {
				v93 = int32(16)
			} else {
				v93 = int32(8)
			}
			v94 = l0 + v93
			if v84&int32(2) == int32(0) {
				v125 = v94
			} else {
				v99 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v94))))
				v100 = v94 + v99
				v104 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v100))))
				switch v104 & int32(7) {
				case 0:
					v121 = int32(base.Ui32(v104) >> (uint(int32(3)) % 32))
				case 1:
					v111 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v100+int32(-2)))))
					v121 = v111
				case 2:
					v114 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v100+int32(-4)))))
					v121 = v114
				case 3:
					v117 = *(*int32)(unsafe.Add(mBase, uint32(v100+int32(-8))))
					v121 = v117
				case 4:
					v120 = *(*int32)(unsafe.Add(mBase, uint32(v100+int32(-16))))
					v121 = v120
				default:
					v121 = int32(0)
				}
				v125 = v100 + int32(1) + v121 + int32(1)
			}
			v140 = *(*int32)(unsafe.Add(mBase, _consts[601]))
			F_valkey_free(m, v125+v140)
			mBase = m.M
			v144 = m.ExcPending
			if v144 != 0 {
				return
			} else {
				return
			}
		}
	}
}
func F_zsetConvert(m *base.Module, l0 int32, l1 int32) {
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	v3 = F_zsetLength(m, l0)
	v4 = m.ExcPending
	if v4 != 0 {
		return
	} else {
		F_zsetConvertAndExpand(m, l0, l1, v3)
		v6 = m.ExcPending
		if v6 != 0 {
			return
		} else {
			return
		}
	}
}
func F_zsetConvertToListpackIfNeeded(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v4&int32(240) == int32(176) {
		return
	} else {
		v9 = F_objectGetVal(m, l0)
		mBase = m.M
		v10 = *(*int32)(unsafe.Add(mBase, uint32(v9)+4))
		v11 = *(*int32)(unsafe.Add(mBase, uint32(v10)))
		v13 = *(*int32)(unsafe.Add(mBase, _consts[942]))
		if base.Ui32(v13) < base.Ui32(v11) {
			return
		} else {
			v16 = *(*int32)(unsafe.Add(mBase, _consts[943]))
			if base.Ui32(v16) < base.Ui32(l1) {
				return
			} else {
				if base.B2i32(base.Ui32(int32(0)+l2) < base.Ui32(int32(1073741825))) == int32(0) {
					return
				} else {
					v28 = F_zsetLength(m, l0)
					mBase = m.M
					v29 = m.ExcPending
					if v29 != 0 {
						return
					} else {
						F_zsetConvertAndExpand(m, l0, int32(11), v28)
						mBase = m.M
						v31 = m.ExcPending
						if v31 != 0 {
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
func F_zsetDel(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	switch int32(base.Ui32(v9)>>(uint(int32(4))%32))&int32(15) + int32(-7) {
	case 0:
		v42 = F_objectGetVal(m, l0)
		mBase = m.M
		v43 = F_zsetRemoveFromSkiplist(m, v42, l1)
		mBase = m.M
		v44 = m.ExcPending
		if v44 != 0 {
			return int32(0)
		} else {
			if v43 != 0 {
				v48 = int32(1)
			} else {
				v48 = int32(0)
			}
			m.G0 = v7 + int32(16)
			return v48
		}
	default:
		F__serverPanic_1(m, int32(_a1723), int32(1641), int32(_a1054), int32(0))
		mBase = m.M
		v39 = m.ExcPending
		if v39 != 0 {
			return int32(0)
		} else {
			F_abort(m)
			mBase = m.M
			base.Wasm_trap_unreachable()
			for {
			}
		}
	case 4:
		v16 = F_objectGetVal(m, l0)
		mBase = m.M
		v18 = F_zzlFind(m, v16, l1, int32(0))
		mBase = m.M
		v21 = m.ExcPending
		if v21 != 0 {
			return int32(0)
		} else {
			if v18 == int32(0) {
				v48 = int32(0)
				m.G0 = v7 + int32(16)
				return v48
			} else {
				v24 = F_objectGetVal(m, l0)
				mBase = m.M
				*(*int32)(unsafe.Add(mBase, uint32(v7)+12)) = v18
				v29 = F_lpDeleteRangeWithEntry(m, v24, v7+int32(12), int32(2))
				mBase = m.M
				v30 = m.ExcPending
				if v30 != 0 {
					return int32(0)
				} else {
					F_objectSetVal(m, l0, v29)
					mBase = m.M
					v32 = m.ExcPending
					if v32 != 0 {
						return int32(0)
					} else {
						v48 = int32(1)
						m.G0 = v7 + int32(16)
						return v48
					}
				}
			}
		}
	}
}
func F_zsetDup(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v20 int32
	_ = v20
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
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v54 float64
	_ = v54
	var v55 int32
	_ = v55
	var v65 int32
	_ = v65
	var v74 int64
	_ = v74
	var v78 int64
	_ = v78
	var v79 int64
	_ = v79
	var v87 int64
	_ = v87
	var v89 int64
	_ = v89
	var v93 int32
	_ = v93
	var v97 int64
	_ = v97
	var v100 int64
	_ = v100
	var v102 int64
	_ = v102
	var v105 int64
	_ = v105
	var v116 int64
	_ = v116
	var v119 int64
	_ = v119
	var v130 int64
	_ = v130
	var v133 int64
	_ = v133
	var v146 int64
	_ = v146
	var v153 int64
	_ = v153
	var v158 int32
	_ = v158
	var v161 int64
	_ = v161
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v177 int64
	_ = v177
	var v185 int64
	_ = v185
	var v188 int64
	_ = v188
	var v200 int64
	_ = v200
	var v202 int32
	_ = v202
	var v204 int64
	_ = v204
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v221 int64
	_ = v221
	var v229 int64
	_ = v229
	var v232 int64
	_ = v232
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v246 int64
	_ = v246
	var v254 int64
	_ = v254
	var v256 int64
	_ = v256
	var v261 int64
	_ = v261
	var v270 int32
	_ = v270
	var v271 int64
	_ = v271
	var v282 int64
	_ = v282
	var v287 int64
	_ = v287
	var v292 int64
	_ = v292
	var v295 int64
	_ = v295
	var v298 int32
	_ = v298
	var v304 int32
	_ = v304
	var v308 int32
	_ = v308
	var v309 int32
	_ = v309
	var v310 int32
	_ = v310
	var v311 int32
	_ = v311
	var v312 int32
	_ = v312
	var v313 int32
	_ = v313
	var v314 int32
	_ = v314
	var v316 int32
	_ = v316
	var v322 int32
	_ = v322
	var v325 int32
	_ = v325
	var v326 int32
	_ = v326
	var v327 int32
	_ = v327
	var v328 int32
	_ = v328
	var v331 int32
	_ = v331
	var v332 int32
	_ = v332
	var v333 int32
	_ = v333
	var v334 int32
	_ = v334
	var v335 int32
	_ = v335
	var v343 int32
	_ = v343
	var v354 int32
	_ = v354
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v9&int32(15) != int32(3) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	F__serverAssert(m, int32(_a1725), int32(_a1723), int32(1718))
	mBase = m.M
	v354 = m.ExcPending
	if v354 != 0 {
		goto L7
	} else {
		goto L45
	}
L2:
	;
	switch int32(base.Ui32(v9)>>(uint(int32(4))%32))&int32(15) + int32(-7) {
	case 0:
		goto L6
	default:
		goto L5
	case 4:
		goto L4
	}
L3:
	;
	return v343
L4:
	;
	v325 = F_objectGetVal(m, l0)
	mBase = m.M
	v326 = *(*int32)(unsafe.Add(mBase, uint32(v325)))
	goto L39
L5:
	;
	F__serverPanic_1(m, int32(_a1723), int32(1752), int32(_a1054), int32(0))
	mBase = m.M
	v322 = m.ExcPending
	if v322 != 0 {
		goto L7
	} else {
		goto L38
	}
L6:
	;
	v20 = F_createZsetObject(m)
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	return int32(0)
L8:
	;
	v24 = F_objectGetVal(m, l0)
	mBase = m.M
	v25 = F_objectGetVal(m, v20)
	mBase = m.M
	v26 = *(*int32)(unsafe.Add(mBase, uint32(v25)))
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v24)))
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v27)+20))
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v27)+16))
	goto L9
L9:
	;
	v31 = F_hashtableExpand(m, v26, v28+v29)
	mBase = m.M
	v32 = m.ExcPending
	if v32 != 0 {
		goto L7
	} else {
		goto L10
	}
L10:
	;
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v24)+4))
	v34 = F_zsetLength(m, l0)
	mBase = m.M
	v35 = m.ExcPending
	if v35 != 0 {
		goto L7
	} else {
		goto L11
	}
L11:
	;
	if v34 == int32(0) {
		v343 = v20
		goto L3
	} else {
		goto L12
	}
L12:
	;
	v38 = v34
	v39 = v33
	goto L13
L13:
	;
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v39)+8))
	v48 = v46 + int32(16)
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v48)))
	v52 = v48 + v49<<(uint(int32(3))%32)
	v53 = int32(*(*int8)(unsafe.Add(mBase, uint32(v52))))
	v54 = *(*float64)(unsafe.Add(mBase, uint32(v46)))
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v25)+4))
	v65 = *(*int32)(unsafe.Add(mBase, _consts[587]))
	if int32(311) < v65 {
		goto L17
	} else {
		goto L18
	}
L15:
	;
	v298 = int32(1)
	if v295 == int64(0) {
		goto L31
	} else {
		goto L32
	}
L16:
	;
	*(*int32)(unsafe.Add(mBase, _consts[587])) = v270
	v282 = int64(base.Ui64(v271)>>(uint(int64(29))%64))&int64(22906492245) ^ v271
	v287 = v282<<(uint(int64(17))%64)&int64(8202884508482404352) ^ v282
	v292 = v287<<(uint(int64(37))%64)&int64(-2270628950310912) ^ v287
	v295 = int64(base.Ui64(v292)>>(uint(int64(43))%64)) ^ v292
	goto L15
L17:
	;
	if v65 == int32(313) {
		goto L20
	} else {
		goto L21
	}
L18:
	;
	v74 = *(*int64)(unsafe.Add(mBase, uint32(v65<<(uint(int32(3))%32))+uint32(_consts[586])))
	v270 = v65 + int32(1)
	v271 = v74
	goto L16
L19:
	;
	v158 = int32(0)
	v161 = v153
	goto L25
L20:
	;
	v79 = int64(5489)
	*(*int64)(unsafe.Add(mBase, _consts[586])) = v79
	v87 = int64(1)
	v89 = v79
	goto L22
L21:
	;
	v78 = *(*int64)(unsafe.Add(mBase, _consts[586]))
	v153 = v78
	goto L19
L22:
	;
	v93 = int32(3)
	v97 = int64(62)
	v100 = int64(6364136223846793005)
	v102 = (int64(base.Ui64(v89)>>(uint(v97)%64))^v89)*v100 + v87
	*(*int64)(unsafe.Add(mBase, uint32(base.I32_wrap_i64(v87)<<(uint(v93)%32))+uint32(_consts[586]))) = v102
	v105 = v87 + int64(1)
	v116 = (int64(base.Ui64(v102)>>(uint(v97)%64))^v102)*v100 + v105
	*(*int64)(unsafe.Add(mBase, uint32(base.I32_wrap_i64(v105)<<(uint(v93)%32))+uint32(_consts[586]))) = v116
	v119 = v87 + int64(2)
	v130 = (int64(base.Ui64(v116)>>(uint(v97)%64))^v116)*v100 + v119
	*(*int64)(unsafe.Add(mBase, uint32(base.I32_wrap_i64(v119)<<(uint(v93)%32))+uint32(_consts[586]))) = v130
	v133 = v87 + int64(3)
	if v133 == int64(312) {
		v153 = v79
		goto L19
	} else {
		goto L24
	}
L24:
	;
	v146 = (int64(base.Ui64(v130)>>(uint(int64(62))%64))^v130)*int64(6364136223846793005) + v133
	*(*int64)(unsafe.Add(mBase, uint32(base.I32_wrap_i64(v133)<<(uint(int32(3))%32))+uint32(_consts[586]))) = v146
	v87 = v87 + int64(4)
	v89 = v146
	goto L22
L25:
	;
	v167 = int32(3)
	v168 = v158 << (uint(v167) % 32)
	v171 = int32(1)
	v172 = v158 + v171
	v177 = *(*int64)(unsafe.Add(mBase, uint32(v172<<(uint(v167)%32))+uint32(_consts[586])))
	v185 = *(*int64)(unsafe.Add(mBase, uint32(base.I32_wrap_i64(v177)&v171<<(uint(v167)%32))+uint32(_consts[944])))
	v188 = *(*int64)(unsafe.Add(mBase, uint32(v168)+uint32(_consts[945])))
	*(*int64)(unsafe.Add(mBase, uint32(v168)+uint32(_consts[586]))) = v185 ^ v188 ^ int64(base.Ui64(v161&int64(-2147483648)|v177&int64(2147483646))>>(uint(int64(1))%64))
	if v172 != int32(156) {
		v158 = v172
		v161 = v177
		goto L25
	} else {
		goto L27
	}
L26:
	;
	v200 = *(*int64)(unsafe.Add(mBase, _consts[945]))
	v202 = int32(156)
	v204 = v200
	goto L28
L27:
	;
	goto L26
L28:
	;
	v211 = int32(3)
	v212 = v202 << (uint(v211) % 32)
	v215 = int32(1)
	v216 = v202 + v215
	v221 = *(*int64)(unsafe.Add(mBase, uint32(v216<<(uint(v211)%32))+uint32(_consts[586])))
	v229 = *(*int64)(unsafe.Add(mBase, uint32(base.I32_wrap_i64(v221)&v215<<(uint(v211)%32))+uint32(_consts[944])))
	v232 = *(*int64)(unsafe.Add(mBase, uint32(v212)+uint32(_consts[946])))
	*(*int64)(unsafe.Add(mBase, uint32(v212)+uint32(_consts[586]))) = v229 ^ v232 ^ int64(base.Ui64(v204&int64(-2147483648)|v221&int64(2147483646))>>(uint(int64(1))%64))
	if v216 != int32(311) {
		v202 = v216
		v204 = v221
		goto L28
	} else {
		goto L30
	}
L29:
	;
	v243 = int32(1)
	v244 = int32(0)
	v246 = *(*int64)(unsafe.Add(mBase, _consts[586]))
	v254 = *(*int64)(unsafe.Add(mBase, uint32(base.I32_wrap_i64(v246)&v243<<(uint(int32(3))%32))+uint32(_consts[944])))
	v256 = *(*int64)(unsafe.Add(mBase, _consts[947]))
	v261 = *(*int64)(unsafe.Add(mBase, _consts[948]))
	*(*int64)(unsafe.Add(mBase, _consts[948])) = v254 ^ v256 ^ int64(base.Ui64(v246&int64(2147483646)|v261&int64(-2147483648))>>(uint(int64(1))%64))
	v270 = v243
	v271 = v246
	goto L16
L30:
	;
	goto L29
L31:
	;
	v304 = int32(32)
	goto L33
L32:
	;
	v304 = int32(base.Ui32(base.I32_wrap_i64(base.I64_clz(v295)))>>(uint(v298)%32)) + v298
	goto L33
L33:
	;
	v308 = F_zslCreateNode(m, v304, v54, v52+v53+int32(1))
	mBase = m.M
	v309 = m.ExcPending
	if v309 != 0 {
		goto L7
	} else {
		goto L34
	}
L34:
	;
	v310 = F_zslInsertNode(m, v55, v308)
	mBase = m.M
	v311 = m.ExcPending
	if v311 != 0 {
		goto L7
	} else {
		goto L35
	}
L35:
	;
	v312 = *(*int32)(unsafe.Add(mBase, uint32(v25)))
	v313 = F_hashtableAdd(m, v312, v310)
	mBase = m.M
	v314 = m.ExcPending
	if v314 != 0 {
		goto L7
	} else {
		goto L36
	}
L36:
	;
	v316 = v38 + int32(-1)
	if v316 != 0 {
		v38 = v316
		v39 = v46
		goto L13
	} else {
		goto L37
	}
L37:
	;
	v343 = v20
	goto L3
L38:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L39:
	;
	v327 = F_valkey_malloc(m, v326)
	mBase = m.M
	v328 = m.ExcPending
	if v328 != 0 {
		goto L7
	} else {
		goto L40
	}
L40:
	;
	if v326 == int32(0) {
		v332 = v327
		goto L42
	} else {
		goto L43
	}
L41:
	;
	v333 = F_createObject(m, int32(3), v332)
	mBase = m.M
	v334 = m.ExcPending
	if v334 != 0 {
		goto L7
	} else {
		goto L44
	}
L42:
	;
	goto L41
L43:
	;
	v331 = F__emscripten_memcpy_bulkmem(m, v327, v325, v326)
	mBase = m.M
	v332 = v331
	goto L42
L44:
	;
	v335 = *(*int32)(unsafe.Add(mBase, uint32(v333)))
	*(*int32)(unsafe.Add(mBase, uint32(v333))) = v335&int32(-241) | int32(176)
	v343 = v333
	goto L3
L45:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_zsetFreeLexRange(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v24 int32
	_ = v24
	v5 = int32(_a578)
	v6 = *(*int32)(unsafe.Add(mBase, _consts[940]))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v9 = *(*int32)(unsafe.Add(mBase, _consts[941]))
	if v7 == v9 {
		v18 = v6
		v19 = v9
		v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		if v20 == v19 {
			return
		} else {
			if v20 == v18 {
				return
			} else {
				F_sdsfree(m, v20)
				mBase = m.M
				v24 = m.ExcPending
				if v24 != 0 {
					return
				} else {
					return
				}
			}
		}
	} else {
		if v7 == v6 {
			v18 = v6
			v19 = v9
			v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			if v20 == v19 {
				return
			} else {
				if v20 == v18 {
					return
				} else {
					F_sdsfree(m, v20)
					mBase = m.M
					v24 = m.ExcPending
					if v24 != 0 {
						return
					} else {
						return
					}
				}
			}
		} else {
			F_sdsfree(m, v7)
			mBase = m.M
			v13 = m.ExcPending
			if v13 != 0 {
				return
			} else {
				v14 = int32(_a578)
				v15 = *(*int32)(unsafe.Add(mBase, _consts[940]))
				v17 = *(*int32)(unsafe.Add(mBase, _consts[941]))
				v18 = v15
				v19 = v17
				v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
				if v20 == v19 {
					return
				} else {
					if v20 == v18 {
						return
					} else {
						F_sdsfree(m, v20)
						mBase = m.M
						v24 = m.ExcPending
						if v24 != 0 {
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
func F_zsetInitScoreRange(m *base.Module, l0 int32, l1 float64, l2 float64, l3 int32, l4 int32, l5 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v18 int32
	_ = v18
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v74 float64
	_ = v74
	var v75 float64
	_ = v75
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v83 float64
	_ = v83
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v93 float64
	_ = v93
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v102 int32
	_ = v102
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v114 int32
	_ = v114
	var __phi114 int32
	_ = __phi114
	var v126 int32
	_ = v126
	var __phi126 int32
	_ = __phi126
	var v129 int32
	_ = v129
	var __phi129 int32
	_ = __phi129
	var v133 float64
	_ = v133
	var v136 int32
	_ = v136
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v149 int32
	_ = v149
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v473 int32
	_ = v473
	var v479 int32
	_ = v479
	var v488 int32
	_ = v488
	var v493 int32
	_ = v493
	var v495 int32
	_ = v495
	var v505 int32
	_ = v505
	var __phi505 int32
	_ = __phi505
	var v514 int32
	_ = v514
	var __phi514 int32
	_ = __phi514
	var v516 int32
	_ = v516
	var __phi516 int32
	_ = __phi516
	var v518 float64
	_ = v518
	var v521 int32
	_ = v521
	var v528 int32
	_ = v528
	var v529 int32
	_ = v529
	var v530 int32
	_ = v530
	var v532 int32
	_ = v532
	var v534 int32
	_ = v534
	var v549 int32
	_ = v549
	var v558 int32
	_ = v558
	var v573 int32
	_ = v573
	var v577 int32
	_ = v577
	var v579 int32
	_ = v579
	var v647 int32
	_ = v647
	var v653 int32
	_ = v653
	var v666 int32
	_ = v666
	var v668 int32
	_ = v668
	var v805 float64
	_ = v805
	var v808 int32
	_ = v808
	var v828 int32
	_ = v828
	var v878 int32
	_ = v878
	var v895 int32
	_ = v895
	var v913 float64
	_ = v913
	var v914 float64
	_ = v914
	var v917 int32
	_ = v917
	var v918 int32
	_ = v918
	var v919 int32
	_ = v919
	var v922 float64
	_ = v922
	var v925 int32
	_ = v925
	var v926 int32
	_ = v926
	var v929 int32
	_ = v929
	var v932 float64
	_ = v932
	var v935 int32
	_ = v935
	var v936 int32
	_ = v936
	var v941 int32
	_ = v941
	var v945 int32
	_ = v945
	var v947 int32
	_ = v947
	var v953 int32
	_ = v953
	var __phi953 int32
	_ = __phi953
	var v965 int32
	_ = v965
	var __phi965 int32
	_ = __phi965
	var v968 int32
	_ = v968
	var __phi968 int32
	_ = __phi968
	var v972 float64
	_ = v972
	var v975 int32
	_ = v975
	var v982 int32
	_ = v982
	var v983 int32
	_ = v983
	var v984 int32
	_ = v984
	var v988 int32
	_ = v988
	var v1002 int32
	_ = v1002
	var v1003 int32
	_ = v1003
	var v1011 int32
	_ = v1011
	var v1015 int32
	_ = v1015
	var v1030 int32
	_ = v1030
	var v1031 int32
	_ = v1031
	var v1035 int32
	_ = v1035
	var v1037 int32
	_ = v1037
	var v1041 int32
	_ = v1041
	var v1051 int32
	_ = v1051
	var __phi1051 int32
	_ = __phi1051
	var v1060 int32
	_ = v1060
	var __phi1060 int32
	_ = __phi1060
	var v1062 int32
	_ = v1062
	var __phi1062 int32
	_ = __phi1062
	var v1064 float64
	_ = v1064
	var v1067 int32
	_ = v1067
	var v1070 int32
	_ = v1070
	var v1076 int32
	_ = v1076
	var v1077 int32
	_ = v1077
	var v1078 int32
	_ = v1078
	var v1082 int32
	_ = v1082
	var v1084 int32
	_ = v1084
	var v1099 int32
	_ = v1099
	var v1106 int32
	_ = v1106
	var v1121 int32
	_ = v1121
	var v1278 float64
	_ = v1278
	var v1281 int32
	_ = v1281
	var v1717 int32
	_ = v1717
	var v1738 int32
	_ = v1738
	var v1740 int32
	_ = v1740
	var v1741 int32
	_ = v1741
	var v1742 int32
	_ = v1742
	var v1744 int32
	_ = v1744
	var v1747 int32
	_ = v1747
	v9 = int32(1)
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v10 == int32(0) {
		v1747 = v9
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return v1747
L2:
	;
	v13 = *(*int32)(unsafe.Add(mBase, uint32(v10)))
	if v13&int32(15) != int32(3) {
		v1747 = v9
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v18 != int32(1) {
		v28 = v10
		goto L4
	} else {
		goto L5
	}
L4:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = int32(2)
	*(*int64)(unsafe.Add(mBase, uint32(l0)+80)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+52)) = l4
	*(*int32)(unsafe.Add(mBase, uint32(l0)+48)) = l3
	*(*float64)(unsafe.Add(mBase, uint32(l0)+40)) = l2
	*(*float64)(unsafe.Add(mBase, uint32(l0)+32)) = l1
	v38 = l0 + int32(32)
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v28)))
	switch int32(base.Ui32(v39)>>(uint(int32(4))%32))&int32(15) + int32(-7) {
	case 0:
		goto L11
	default:
		goto L10
	case 4:
		goto L12
	}
L5:
	;
	F_zsetFreeLexRange(m, l0+int32(56))
	mBase = m.M
	v26 = m.ExcPending
	if v26 != 0 {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	return int32(0)
L7:
	;
	v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v28 = v27
	goto L4
L8:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+80)) = v1742
	v1744 = int32(0)
	if v1742 != 0 {
		v1747 = v1744
		goto L1
	} else {
		goto L303
	}
L9:
	;
	v1740 = F_zzlLastInRange(m, v46, v38)
	mBase = m.M
	v1741 = m.ExcPending
	if v1741 != 0 {
		goto L6
	} else {
		goto L302
	}
L10:
	;
	F__serverPanic_1(m, int32(_a917), int32(5134), int32(_a927), int32(0))
	mBase = m.M
	v1738 = m.ExcPending
	if v1738 != 0 {
		goto L6
	} else {
		goto L301
	}
L11:
	;
	v51 = F_objectGetVal(m, v28)
	mBase = m.M
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v51)+4))
	if l5 == int32(0) {
		goto L15
	} else {
		goto L16
	}
L12:
	;
	v46 = F_objectGetVal(m, v28)
	mBase = m.M
	if l5 == int32(0) {
		goto L9
	} else {
		goto L13
	}
L13:
	;
	v49 = F_zzlFirstInRange(m, v46, v38)
	mBase = m.M
	v50 = m.ExcPending
	if v50 != 0 {
		goto L6
	} else {
		goto L14
	}
L14:
	;
	v1742 = v49
	goto L8
L15:
	;
	v895 = int32(0)
	v913 = *(*float64)(unsafe.Add(mBase, uint32(v38)))
	v914 = *(*float64)(unsafe.Add(mBase, uint32(v38)+8))
	if base.F64_gt(v913, v914) != 0 {
		v1717 = v895
		goto L160
	} else {
		goto L161
	}
L16:
	;
	v55 = int32(0)
	v74 = *(*float64)(unsafe.Add(mBase, uint32(v38)))
	v75 = *(*float64)(unsafe.Add(mBase, uint32(v38)+8))
	if base.F64_gt(v74, v75) != 0 {
		v878 = v55
		goto L18
	} else {
		goto L19
	}
L17:
	;
	v1742 = v878
	goto L8
L18:
	;
	goto L17
L19:
	;
	if base.F64_ne(v74, v75) != 0 {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	v80 = *(*int32)(unsafe.Add(mBase, uint32(v52)+8))
	if v80 == int32(0) {
		v878 = v55
		goto L18
	} else {
		goto L24
	}
L21:
	;
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v38)+16))
	if v78 != 0 {
		v878 = v55
		goto L18
	} else {
		goto L22
	}
L22:
	;
	v79 = *(*int32)(unsafe.Add(mBase, uint32(v38)+20))
	if v79 != 0 {
		v878 = v55
		goto L18
	} else {
		goto L23
	}
L23:
	;
	goto L20
L24:
	;
	v83 = *(*float64)(unsafe.Add(mBase, uint32(v80)))
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v38)+16))
	if v86 != 0 {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	v87 = base.F64_gt(v83, v74)
	goto L27
L26:
	;
	v87 = base.F64_ge(v83, v74)
	goto L27
L27:
	;
	if v87 != int32(1) {
		v878 = v55
		goto L18
	} else {
		goto L28
	}
L28:
	;
	v90 = *(*int32)(unsafe.Add(mBase, uint32(v52)+12))
	if v90 == int32(0) {
		v878 = v55
		goto L18
	} else {
		goto L29
	}
L29:
	;
	v93 = *(*float64)(unsafe.Add(mBase, uint32(v90)))
	v96 = *(*int32)(unsafe.Add(mBase, uint32(v38)+20))
	if v96 != 0 {
		goto L30
	} else {
		goto L31
	}
L30:
	;
	v97 = base.F64_lt(v93, v75)
	goto L32
L31:
	;
	v97 = base.F64_le(v93, v75)
	goto L32
L32:
	;
	if v97 != int32(1) {
		v878 = v55
		goto L18
	} else {
		goto L33
	}
L33:
	;
	v102 = *(*int32)(unsafe.Add(mBase, uint32(v52)+16))
	v106 = (v102 + int32(-1)) << (uint(int32(3)) % 32)
	v108 = *(*int32)(unsafe.Add(mBase, uint32(v52+int32(12)+v106)))
	if v108 != 0 {
		goto L35
	} else {
		goto L36
	}
L34:
	;
	goto L49
L35:
	;
	__phi114 = v108
	__phi126 = int32(0)
	__phi129 = v52
	v114 = __phi114
	v126 = __phi126
	v129 = __phi129
	goto L37
L36:
	;
	v163 = int32(0)
	v164 = v52
	goto L34
L37:
	;
	v133 = *(*float64)(unsafe.Add(mBase, uint32(v114)))
	if v86 != 0 {
		goto L40
	} else {
		goto L41
	}
L38:
	;
	v163 = v145
	v164 = v114
	goto L34
L39:
	;
	if v102 < int32(2) {
		v144 = int32(1)
		goto L44
	} else {
		goto L45
	}
L40:
	;
	v136 = base.F64_gt(v133, v74)
	goto L42
L41:
	;
	v136 = base.F64_ge(v133, v74)
	goto L42
L42:
	;
	if v136 == int32(0) {
		goto L39
	} else {
		goto L43
	}
L43:
	;
	v163 = v126
	v164 = v129
	goto L34
L44:
	;
	v145 = v144 + v126
	v149 = *(*int32)(unsafe.Add(mBase, uint32(v114+v106+int32(12))))
	if v149 != 0 {
		__phi114 = v149
		__phi126 = v145
		__phi129 = v114
		v114 = __phi114
		v126 = __phi126
		v129 = __phi129
		goto L37
	} else {
		goto L46
	}
L45:
	;
	v143 = *(*int32)(unsafe.Add(mBase, uint32(v129+v106+int32(16))))
	v144 = v143
	goto L44
L46:
	;
	goto L38
L49:
	;
	if v102 < int32(2) {
		v558 = v164
		v573 = v163
		goto L103
	} else {
		goto L104
	}
L103:
	;
	v577 = int32(0)
	v579 = *(*int32)(unsafe.Add(mBase, uint32(v52)))
	if base.Ui32(v579) <= base.Ui32(v573+v55) {
		v878 = v577
		goto L18
	} else {
		goto L121
	}
L104:
	;
	v473 = v164
	v479 = v102 + int32(-2)
	v488 = v163
	goto L105
L105:
	;
	v493 = v479 << (uint(int32(3)) % 32)
	v495 = *(*int32)(unsafe.Add(mBase, uint32(v473+v493)+12))
	if v495 == int32(0) {
		v534 = v473
		v549 = v488
		goto L107
	} else {
		goto L108
	}
L106:
	;
	v558 = v534
	v573 = v549
	goto L103
L107:
	;
	if int32(0) < v479 {
		v473 = v534
		v479 = v479 + int32(-1)
		v488 = v549
		goto L105
	} else {
		goto L120
	}
L108:
	;
	__phi505 = v495
	__phi514 = v488
	__phi516 = v473
	v505 = __phi505
	v514 = __phi514
	v516 = __phi516
	goto L109
L109:
	;
	v518 = *(*float64)(unsafe.Add(mBase, uint32(v505)))
	if v86 != 0 {
		goto L112
	} else {
		goto L113
	}
L110:
	;
	v534 = v505
	v549 = v530
	goto L107
L111:
	;
	if v479 != 0 {
		goto L117
	} else {
		goto L118
	}
L112:
	;
	v521 = base.F64_gt(v518, v74)
	goto L114
L113:
	;
	v521 = base.F64_ge(v518, v74)
	goto L114
L114:
	;
	if v521 == int32(0) {
		goto L111
	} else {
		goto L115
	}
L115:
	;
	v534 = v516
	v549 = v514
	goto L107
L116:
	;
	v530 = v529 + v514
	v532 = *(*int32)(unsafe.Add(mBase, uint32(v505+v493)+12))
	if v532 != 0 {
		__phi505 = v532
		__phi514 = v530
		__phi516 = v505
		v505 = __phi505
		v514 = __phi514
		v516 = __phi516
		goto L109
	} else {
		goto L119
	}
L117:
	;
	v528 = *(*int32)(unsafe.Add(mBase, uint32(v516+v493+int32(16))))
	v529 = v528
	goto L116
L118:
	;
	v529 = int32(1)
	goto L116
L119:
	;
	goto L110
L120:
	;
	goto L106
L121:
	;
	goto L125
L122:
	;
	v878 = v828
	goto L18
L123:
	;
	v805 = *(*float64)(unsafe.Add(mBase, uint32(v666)))
	if v96 != 0 {
		goto L154
	} else {
		goto L155
	}
L125:
	;
	goto L126
L126:
	;
	goto L132
L131:
	;
	if v666 != 0 {
		goto L123
	} else {
		goto L136
	}
L132:
	;
	v647 = v558
	v653 = int32(0)
	goto L133
L133:
	;
	v666 = *(*int32)(unsafe.Add(mBase, uint32(v647)+12))
	v668 = v653 + int32(1)
	if v668 != int32(1) {
		v647 = v666
		v653 = v668
		goto L133
	} else {
		goto L135
	}
L134:
	;
	goto L131
L135:
	;
	goto L134
L136:
	;
	v828 = int32(0)
	goto L122
L154:
	;
	v808 = base.F64_lt(v805, v75)
	goto L156
L155:
	;
	v808 = base.F64_le(v805, v75)
	goto L156
L156:
	;
	if v808 != int32(1) {
		v878 = v577
		goto L18
	} else {
		goto L157
	}
L157:
	;
	v828 = v666
	goto L122
L159:
	;
	v1742 = v1717
	goto L8
L160:
	;
	goto L159
L161:
	;
	if base.F64_ne(v913, v914) != 0 {
		goto L162
	} else {
		goto L163
	}
L162:
	;
	v919 = *(*int32)(unsafe.Add(mBase, uint32(v52)+8))
	if v919 == int32(0) {
		v1717 = v895
		goto L160
	} else {
		goto L166
	}
L163:
	;
	v917 = *(*int32)(unsafe.Add(mBase, uint32(v38)+16))
	if v917 != 0 {
		v1717 = v895
		goto L160
	} else {
		goto L164
	}
L164:
	;
	v918 = *(*int32)(unsafe.Add(mBase, uint32(v38)+20))
	if v918 != 0 {
		v1717 = v895
		goto L160
	} else {
		goto L165
	}
L165:
	;
	goto L162
L166:
	;
	v922 = *(*float64)(unsafe.Add(mBase, uint32(v919)))
	v925 = *(*int32)(unsafe.Add(mBase, uint32(v38)+16))
	if v925 != 0 {
		goto L167
	} else {
		goto L168
	}
L167:
	;
	v926 = base.F64_gt(v922, v913)
	goto L169
L168:
	;
	v926 = base.F64_ge(v922, v913)
	goto L169
L169:
	;
	if v926 != int32(1) {
		v1717 = v895
		goto L160
	} else {
		goto L170
	}
L170:
	;
	v929 = *(*int32)(unsafe.Add(mBase, uint32(v52)+12))
	if v929 == int32(0) {
		v1717 = v895
		goto L160
	} else {
		goto L171
	}
L171:
	;
	v932 = *(*float64)(unsafe.Add(mBase, uint32(v929)))
	v935 = *(*int32)(unsafe.Add(mBase, uint32(v38)+20))
	if v935 != 0 {
		goto L172
	} else {
		goto L173
	}
L172:
	;
	v936 = base.F64_lt(v932, v914)
	goto L174
L173:
	;
	v936 = base.F64_le(v932, v914)
	goto L174
L174:
	;
	if v936 != int32(1) {
		v1717 = v895
		goto L160
	} else {
		goto L175
	}
L175:
	;
	v941 = *(*int32)(unsafe.Add(mBase, uint32(v52)+16))
	v945 = (v941 + int32(-1)) << (uint(int32(3)) % 32)
	v947 = *(*int32)(unsafe.Add(mBase, uint32(v52+int32(12)+v945)))
	if v947 != 0 {
		goto L177
	} else {
		goto L178
	}
L176:
	;
	goto L192
L177:
	;
	__phi953 = v947
	__phi965 = int32(0)
	__phi968 = v52
	v953 = __phi953
	v965 = __phi965
	v968 = __phi968
	goto L179
L178:
	;
	v1002 = int32(0)
	v1003 = v52
	goto L176
L179:
	;
	v972 = *(*float64)(unsafe.Add(mBase, uint32(v953)))
	if v925 != 0 {
		goto L182
	} else {
		goto L183
	}
L180:
	;
	v1002 = v984
	v1003 = v953
	goto L176
L181:
	;
	if v941 < int32(2) {
		v983 = int32(1)
		goto L186
	} else {
		goto L187
	}
L182:
	;
	v975 = base.F64_gt(v972, v913)
	goto L184
L183:
	;
	v975 = base.F64_ge(v972, v913)
	goto L184
L184:
	;
	if v975 == int32(0) {
		goto L181
	} else {
		goto L185
	}
L185:
	;
	v1002 = v965
	v1003 = v968
	goto L176
L186:
	;
	v984 = v983 + v965
	v988 = *(*int32)(unsafe.Add(mBase, uint32(v953+v945+int32(12))))
	if v988 != 0 {
		__phi953 = v988
		__phi965 = v984
		__phi968 = v953
		v953 = __phi953
		v965 = __phi965
		v968 = __phi968
		goto L179
	} else {
		goto L188
	}
L187:
	;
	v982 = *(*int32)(unsafe.Add(mBase, uint32(v968+v945+int32(16))))
	v983 = v982
	goto L186
L188:
	;
	goto L180
L189:
	;
	v1717 = v1106
	goto L160
L192:
	;
	v1011 = int32(0)
	if v941 <= v1011 {
		v1106 = v1003
		v1121 = v1002
		goto L193
	} else {
		goto L194
	}
L193:
	;
	if v1121 < int32(1) {
		v1717 = v1011
		goto L160
	} else {
		goto L210
	}
L194:
	;
	v1015 = v1003
	v1030 = v1002
	v1031 = v941
	goto L195
L195:
	;
	v1035 = v1031 + int32(-1)
	v1037 = v1035 << (uint(int32(3)) % 32)
	v1041 = *(*int32)(unsafe.Add(mBase, uint32(v1015+v1037+int32(12))))
	if v1041 == int32(0) {
		v1084 = v1015
		v1099 = v1030
		goto L197
	} else {
		goto L198
	}
L196:
	;
	v1106 = v1084
	v1121 = v1099
	goto L193
L197:
	;
	if int32(1) < v1031 {
		v1015 = v1084
		v1030 = v1099
		v1031 = v1035
		goto L195
	} else {
		goto L209
	}
L198:
	;
	__phi1051 = v1041
	__phi1060 = v1030
	__phi1062 = v1015
	v1051 = __phi1051
	v1060 = __phi1060
	v1062 = __phi1062
	goto L199
L199:
	;
	v1064 = *(*float64)(unsafe.Add(mBase, uint32(v1051)))
	if v935 != 0 {
		goto L202
	} else {
		goto L203
	}
L200:
	;
	v1084 = v1051
	v1099 = v1078
	goto L197
L201:
	;
	v1070 = int32(1)
	if v1031 == v1070 {
		v1077 = v1070
		goto L206
	} else {
		goto L207
	}
L202:
	;
	v1067 = base.F64_lt(v1064, v914)
	goto L204
L203:
	;
	v1067 = base.F64_le(v1064, v914)
	goto L204
L204:
	;
	if v1067 == int32(1) {
		goto L201
	} else {
		goto L205
	}
L205:
	;
	v1084 = v1062
	v1099 = v1060
	goto L197
L206:
	;
	v1078 = v1077 + v1060
	v1082 = *(*int32)(unsafe.Add(mBase, uint32(v1051+v1037+int32(12))))
	if v1082 != 0 {
		__phi1051 = v1082
		__phi1060 = v1078
		__phi1062 = v1051
		v1051 = __phi1051
		v1060 = __phi1060
		v1062 = __phi1062
		goto L199
	} else {
		goto L208
	}
L207:
	;
	v1076 = *(*int32)(unsafe.Add(mBase, uint32(v1062+v1037+int32(16))))
	v1077 = v1076
	goto L206
L208:
	;
	goto L200
L209:
	;
	goto L196
L210:
	;
	goto L214
L211:
	;
	goto L189
L212:
	;
	v1278 = *(*float64)(unsafe.Add(mBase, uint32(v1106)))
	if v925 != 0 {
		goto L240
	} else {
		goto L241
	}
L214:
	;
	goto L212
L240:
	;
	v1281 = base.F64_gt(v1278, v913)
	goto L242
L241:
	;
	v1281 = base.F64_ge(v1278, v913)
	goto L242
L242:
	;
	if v1281 != int32(1) {
		v1717 = v1011
		goto L160
	} else {
		goto L243
	}
L243:
	;
	goto L211
L301:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L302:
	;
	v1742 = v1740
	goto L8
L303:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+84)) = int32(1)
	v1747 = v1744
	goto L1
}
func F_zsetParseLexRange(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v12 int32
	_ = v12
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v37 int32
	_ = v37
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
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v57 int32
	_ = v57
	var v61 int32
	_ = v61
	v6 = int32(-1)
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v7&int32(240) == int32(16) {
		v61 = v6
		return v61
	} else {
		v12 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
		if v12&int32(240) == int32(16) {
			v61 = v6
			return v61
		} else {
			*(*int64)(unsafe.Add(mBase, uint32(l2))) = int64(0)
			v20 = l2 + int32(4)
			v23 = F_zslParseLexRangeItem(m, l0, l2, l2+int32(8))
			mBase = m.M
			v26 = m.ExcPending
			if v26 != 0 {
				return int32(0)
			} else {
				if v23 == int32(-1) {
					v37 = int32(_a578)
					v38 = *(*int32)(unsafe.Add(mBase, _consts[940]))
					v39 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
					v41 = *(*int32)(unsafe.Add(mBase, _consts[941]))
					if v39 == v41 {
						v50 = v38
						v51 = v41
						v52 = int32(-1)
						v53 = *(*int32)(unsafe.Add(mBase, uint32(v20)))
						if v53 == v51 {
							v61 = v52
							return v61
						} else {
							if v53 == v50 {
								v61 = v52
								return v61
							} else {
								F_sdsfree(m, v53)
								mBase = m.M
								v57 = m.ExcPending
								if v57 != 0 {
									return int32(0)
								} else {
									v61 = v52
									return v61
								}
							}
						}
					} else {
						if v39 == v38 {
							v50 = v38
							v51 = v41
							v52 = int32(-1)
							v53 = *(*int32)(unsafe.Add(mBase, uint32(v20)))
							if v53 == v51 {
								v61 = v52
								return v61
							} else {
								if v53 == v50 {
									v61 = v52
									return v61
								} else {
									F_sdsfree(m, v53)
									mBase = m.M
									v57 = m.ExcPending
									if v57 != 0 {
										return int32(0)
									} else {
										v61 = v52
										return v61
									}
								}
							}
						} else {
							F_sdsfree(m, v39)
							mBase = m.M
							v45 = m.ExcPending
							if v45 != 0 {
								return int32(0)
							} else {
								v46 = int32(_a578)
								v47 = *(*int32)(unsafe.Add(mBase, _consts[940]))
								v49 = *(*int32)(unsafe.Add(mBase, _consts[941]))
								v50 = v47
								v51 = v49
								v52 = int32(-1)
								v53 = *(*int32)(unsafe.Add(mBase, uint32(v20)))
								if v53 == v51 {
									v61 = v52
									return v61
								} else {
									if v53 == v50 {
										v61 = v52
										return v61
									} else {
										F_sdsfree(m, v53)
										mBase = m.M
										v57 = m.ExcPending
										if v57 != 0 {
											return int32(0)
										} else {
											v61 = v52
											return v61
										}
									}
								}
							}
						}
					}
				} else {
					v32 = F_zslParseLexRangeItem(m, l1, v20, l2+int32(12))
					mBase = m.M
					v33 = m.ExcPending
					if v33 != 0 {
						return int32(0)
					} else {
						if v32 != int32(-1) {
							v61 = int32(0)
							return v61
						} else {
							v37 = int32(_a578)
							v38 = *(*int32)(unsafe.Add(mBase, _consts[940]))
							v39 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
							v41 = *(*int32)(unsafe.Add(mBase, _consts[941]))
							if v39 == v41 {
								v50 = v38
								v51 = v41
								v52 = int32(-1)
								v53 = *(*int32)(unsafe.Add(mBase, uint32(v20)))
								if v53 == v51 {
									v61 = v52
									return v61
								} else {
									if v53 == v50 {
										v61 = v52
										return v61
									} else {
										F_sdsfree(m, v53)
										mBase = m.M
										v57 = m.ExcPending
										if v57 != 0 {
											return int32(0)
										} else {
											v61 = v52
											return v61
										}
									}
								}
							} else {
								if v39 == v38 {
									v50 = v38
									v51 = v41
									v52 = int32(-1)
									v53 = *(*int32)(unsafe.Add(mBase, uint32(v20)))
									if v53 == v51 {
										v61 = v52
										return v61
									} else {
										if v53 == v50 {
											v61 = v52
											return v61
										} else {
											F_sdsfree(m, v53)
											mBase = m.M
											v57 = m.ExcPending
											if v57 != 0 {
												return int32(0)
											} else {
												v61 = v52
												return v61
											}
										}
									}
								} else {
									F_sdsfree(m, v39)
									mBase = m.M
									v45 = m.ExcPending
									if v45 != 0 {
										return int32(0)
									} else {
										v46 = int32(_a578)
										v47 = *(*int32)(unsafe.Add(mBase, _consts[940]))
										v49 = *(*int32)(unsafe.Add(mBase, _consts[941]))
										v50 = v47
										v51 = v49
										v52 = int32(-1)
										v53 = *(*int32)(unsafe.Add(mBase, uint32(v20)))
										if v53 == v51 {
											v61 = v52
											return v61
										} else {
											if v53 == v50 {
												v61 = v52
												return v61
											} else {
												F_sdsfree(m, v53)
												mBase = m.M
												v57 = m.ExcPending
												if v57 != 0 {
													return int32(0)
												} else {
													v61 = v52
													return v61
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
func F_zsetScore(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
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
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v33 float64
	_ = v33
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v10 = int32(-1)
	if l0 == int32(0) {
		v48 = v10
		m.G0 = v8 + int32(16)
		return v48
	} else {
		if l1 == int32(0) {
			v48 = v10
			m.G0 = v8 + int32(16)
			return v48
		} else {
			v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			switch int32(base.Ui32(v15)>>(uint(int32(4))%32))&int32(15) + int32(-7) {
			case 0:
				v22 = F_objectGetVal(m, l0)
				mBase = m.M
				v23 = *(*int32)(unsafe.Add(mBase, uint32(v22)))
				v26 = F_hashtableFind(m, v23, l1, v8+int32(12))
				mBase = m.M
				v29 = m.ExcPending
				if v29 != 0 {
					return int32(0)
				} else {
					if v26 == int32(0) {
						v48 = v10
					} else {
						v32 = *(*int32)(unsafe.Add(mBase, uint32(v8)+12))
						v33 = *(*float64)(unsafe.Add(mBase, uint32(v32)))
						*(*float64)(unsafe.Add(mBase, uint32(l2))) = v33
						v48 = int32(0)
					}
					m.G0 = v8 + int32(16)
					return v48
				}
			default:
				F__serverPanic_1(m, int32(_a1723), int32(1432), int32(_a1054), int32(0))
				mBase = m.M
				v40 = m.ExcPending
				if v40 != 0 {
					return int32(0)
				} else {
					F_abort(m)
					mBase = m.M
					base.Wasm_trap_unreachable()
					for {
					}
				}
			case 4:
				v42 = F_objectGetVal(m, l0)
				mBase = m.M
				v43 = F_zzlFind(m, v42, l1, l2)
				mBase = m.M
				v44 = m.ExcPending
				if v44 != 0 {
					return int32(0)
				} else {
					if v43 == int32(0) {
						v48 = v10
					} else {
						v48 = int32(0)
					}
					m.G0 = v8 + int32(16)
					return v48
				}
			}
		}
	}
}
func F_zsetTypeCreate(m *base.Module, l0 int32, l1 int32) int32 {
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
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	v4 = *(*int32)(unsafe.Add(mBase, _consts[942]))
	if base.Ui32(v4) < base.Ui32(l0) {
		v14 = F_createZsetObject(m)
		mBase = m.M
		v15 = m.ExcPending
		if v15 != 0 {
			return int32(0)
		} else {
			v16 = F_objectGetVal(m, v14)
			mBase = m.M
			v17 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
			v18 = F_hashtableExpand(m, v17, l0)
			mBase = m.M
			v19 = m.ExcPending
			if v19 != 0 {
				return int32(0)
			} else {
				return v14
			}
		}
	} else {
		v7 = *(*int32)(unsafe.Add(mBase, _consts[943]))
		if base.Ui32(v7) < base.Ui32(l1) {
			v14 = F_createZsetObject(m)
			mBase = m.M
			v15 = m.ExcPending
			if v15 != 0 {
				return int32(0)
			} else {
				v16 = F_objectGetVal(m, v14)
				mBase = m.M
				v17 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
				v18 = F_hashtableExpand(m, v17, l0)
				mBase = m.M
				v19 = m.ExcPending
				if v19 != 0 {
					return int32(0)
				} else {
					return v14
				}
			}
		} else {
			v9 = F_createZsetListpackObject(m)
			mBase = m.M
			v12 = m.ExcPending
			if v12 != 0 {
				return int32(0)
			} else {
				return v9
			}
		}
	}
}
func F_zsetTypeMaybeConvert(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v17 int32
	_ = v17
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v4&int32(240) != int32(176) {
		return
	} else {
		v10 = *(*int32)(unsafe.Add(mBase, _consts[942]))
		if base.Ui32(v10) < base.Ui32(l1) {
			F_zsetConvertAndExpand(m, l0, int32(7), l1)
			mBase = m.M
			v17 = m.ExcPending
			if v17 != 0 {
				return
			} else {
				return
			}
		} else {
			v13 = *(*int32)(unsafe.Add(mBase, _consts[943]))
			if base.Ui32(l2) <= base.Ui32(v13) {
				return
			} else {
				F_zsetConvertAndExpand(m, l0, int32(7), l1)
				mBase = m.M
				v17 = m.ExcPending
				if v17 != 0 {
					return
				} else {
					return
				}
			}
		}
	}
}
func F_zsetTypeRandomElement(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
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
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v37 int32
	_ = v37
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v58 float64
	_ = v58
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v83 int64
	_ = v83
	var v88 int64
	_ = v88
	var v96 int32
	_ = v96
	var v103 int32
	_ = v103
	var v109 float64
	_ = v109
	var v113 int64
	_ = v113
	var v115 float64
	_ = v115
	var v130 int32
	_ = v130
	v9 = m.G0
	v11 = v9 - int32(16)
	m.G0 = v11
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	switch int32(base.Ui32(v13)>>(uint(int32(4))%32))&int32(15) + int32(-7) {
	case 0:
		v20 = F_objectGetVal(m, l0)
		mBase = m.M
		v21 = *(*int32)(unsafe.Add(mBase, uint32(v20)))
		v22 = F_hashtableFairRandomEntry(m, v21, v11)
		mBase = m.M
		v23 = m.ExcPending
		if v23 != 0 {
			return
		} else {
			v24 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
			v26 = v24 + int32(16)
			v27 = *(*int32)(unsafe.Add(mBase, uint32(v26)))
			v30 = v26 + v27<<(uint(int32(3))%32)
			v31 = int32(*(*int8)(unsafe.Add(mBase, uint32(v30))))
			v32 = v30 + v31
			*(*int32)(unsafe.Add(mBase, uint32(l2))) = v32 + int32(1)
			v37 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v32))))
			switch v37 & int32(7) {
			case 0:
				v54 = int32(base.Ui32(v37) >> (uint(int32(3)) % 32))
			case 1:
				v44 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v32+int32(-2)))))
				v54 = v44
			case 2:
				v47 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v32+int32(-4)))))
				v54 = v47
			case 3:
				v50 = *(*int32)(unsafe.Add(mBase, uint32(v32+int32(-8))))
				v54 = v50
			case 4:
				v53 = *(*int32)(unsafe.Add(mBase, uint32(v32+int32(-16))))
				v54 = v53
			default:
				v54 = int32(0)
			}
			*(*int32)(unsafe.Add(mBase, uint32(l2)+4)) = v54
			if l3 == int32(0) {
			} else {
				v58 = *(*float64)(unsafe.Add(mBase, uint32(v24)))
				*(*float64)(unsafe.Add(mBase, uint32(l3))) = v58
			}
			m.G0 = v11 + int32(16)
			return
		}
	default:
		F__serverPanic_1(m, int32(_a1723), int32(1796), int32(_a1744), int32(0))
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
	case 4:
		v60 = F_objectGetVal(m, l0)
		mBase = m.M
		F_lpRandomPair(m, v60, l1, l2, v11)
		mBase = m.M
		v62 = m.ExcPending
		if v62 != 0 {
			return
		} else {
			if l3 == int32(0) {
			} else {
				v65 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
				if v65 == int32(0) {
					v113 = *(*int64)(unsafe.Add(mBase, uint32(v11)+8))
					v115 = base.F64_convert_i64_s(v113)
				} else {
					v68 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
					v69 = int32(0)
					v73 = m.G0
					v75 = v73 - int32(32)
					m.G0 = v75
					v77 = F___errno_location(m)
					mBase = m.M
					*(*int32)(unsafe.Add(mBase, uint32(v77))) = v69
					v83 = *(*int64)(unsafe.Add(mBase, _consts[378]))
					*(*int64)(unsafe.Add(mBase, uint32(v75+int32(8)))) = v83
					*(*int64)(unsafe.Add(mBase, uint32(v75)+24)) = int64(0)
					v88 = *(*int64)(unsafe.Add(mBase, _consts[379]))
					*(*int64)(unsafe.Add(mBase, uint32(v75))) = v88
					F_ffc_from_chars_double_options(m, v75+int32(16), v65, v65+v68, v75+int32(24), v75)
					mBase = m.M
					v96 = *(*int32)(unsafe.Add(mBase, uint32(v75)+20))
					if v96 == v69 {
					} else {
						if v96 == int32(2) {
							v103 = int32(68)
						} else {
							v103 = int32(28)
						}
						*(*int32)(unsafe.Add(mBase, uint32(v77))) = v103
					}
					v109 = *(*float64)(unsafe.Add(mBase, uint32(v75)+24))
					m.G0 = v75 + int32(32)
					v115 = v109
				}
				*(*float64)(unsafe.Add(mBase, uint32(l3))) = v115
			}
			m.G0 = v11 + int32(16)
			return
		}
	}
}
