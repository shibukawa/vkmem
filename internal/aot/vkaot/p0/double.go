package p0

import (
	base "github.com/shibukawa/vkmem/internal/aot/vkaot/base"
	"math"
	"unsafe"
)

func F___DOUBLE_BITS_1(m *base.Module, l0 float64) int64 {
	return base.I64_reinterpret_f64(l0)
}
func F_doubleCallback(m *base.Module, l0 int32, l1 float64) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v54 int32
	_ = v54
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v7 < int32(0) {
		v31 = v6
	} else {
		v12 = l0 + v7<<(uint(int32(2))%32)
		v13 = *(*int32)(unsafe.Add(mBase, uint32(v12)+16))
		v14 = int32(1)
		v15 = v13 + v14
		*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = v15
		v17 = *(*int32)(unsafe.Add(mBase, uint32(v12)+1040))
		if v17 != v14 {
			v31 = v6
		} else {
			v22 = *(*int32)(unsafe.Add(mBase, uint32(v6)+8))
			*(*int32)(unsafe.Add(mBase, uint32(v22)+8)) = int32(3)
			*(*float64)(unsafe.Add(mBase, uint32(v22))) = base.F64_convert_i32_u(v15)
			v26 = *(*int32)(unsafe.Add(mBase, uint32(v6)+8))
			*(*int32)(unsafe.Add(mBase, uint32(v6)+8)) = v26 + int32(16)
			v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			v31 = v30
		}
	}
	v35 = F_lua_checkstack(m, v31, int32(3))
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
		return
	} else {
		if v35 != 0 {
			v40 = int32(0)
			F_lua_createtable(m, v31, v40, v40)
			mBase = m.M
			v43 = m.ExcPending
			if v43 != 0 {
				return
			} else {
				v44 = m.G3
				F_lua_pushstring(m, v31, v44+int32(_a2237))
				mBase = m.M
				v48 = m.ExcPending
				if v48 != 0 {
					return
				} else {
					v50 = *(*int32)(unsafe.Add(mBase, uint32(v31)+8))
					*(*int32)(unsafe.Add(mBase, uint32(v50)+8)) = int32(3)
					*(*float64)(unsafe.Add(mBase, uint32(v50))) = l1
					v54 = *(*int32)(unsafe.Add(mBase, uint32(v31)+8))
					*(*int32)(unsafe.Add(mBase, uint32(v31)+8)) = v54 + int32(16)
					F_lua_settable(m, v31, int32(-3))
					mBase = m.M
					v60 = m.ExcPending
					if v60 != 0 {
						return
					} else {
						F_processCollectionElementEnd(m, l0)
						mBase = m.M
						v62 = m.ExcPending
						if v62 != 0 {
							return
						} else {
							return
						}
					}
				}
			}
		} else {
			F__serverPanic_2(m, int32(1094))
			mBase = m.M
			v39 = m.ExcPending
			if v39 != 0 {
				return
			} else {
				base.Wasm_trap_unreachable()
				for {
				}
			}
		}
	}
}
func F_getDoubleFromObject(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v21 int32
	_ = v21
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v41 int32
	_ = v41
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v63 int32
	_ = v63
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v105 int32
	_ = v105
	var v112 int32
	_ = v112
	var v115 int32
	_ = v115
	var v118 int32
	_ = v118
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v126 int32
	_ = v126
	var v141 int32
	_ = v141
	var v144 int32
	_ = v144
	var v147 int32
	_ = v147
	var v149 int32
	_ = v149
	var v153 int32
	_ = v153
	var v160 int32
	_ = v160
	var v163 int32
	_ = v163
	var v166 int32
	_ = v166
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v178 int32
	_ = v178
	var v180 int32
	_ = v180
	var v182 int32
	_ = v182
	var v187 float64
	_ = v187
	var v191 int32
	_ = v191
	var v198 int32
	_ = v198
	var v201 int32
	_ = v201
	var v208 int32
	_ = v208
	var v211 int64
	_ = v211
	var v212 int64
	_ = v212
	var v218 int32
	_ = v218
	var v224 int32
	_ = v224
	var v227 int32
	_ = v227
	var v235 float64
	_ = v235
	var v236 int32
	_ = v236
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v251 int32
	_ = v251
	var v252 int32
	_ = v252
	var v256 int32
	_ = v256
	var v263 int32
	_ = v263
	var v266 int32
	_ = v266
	var v269 int32
	_ = v269
	var v272 int32
	_ = v272
	var v273 int32
	_ = v273
	var v277 int32
	_ = v277
	var v292 int32
	_ = v292
	var v303 int32
	_ = v303
	var v305 int32
	_ = v305
	var v308 float64
	_ = v308
	var v315 int32
	_ = v315
	var v330 int32
	_ = v330
	v9 = m.G0
	v11 = v9 - int32(16)
	m.G0 = v11
	if l0 != 0 {
		v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		if v14&int32(15) != 0 {
			F__serverAssertWithInfo(m, int32(0), l0, int32(_a479), int32(_a838), int32(1032))
			mBase = m.M
			v330 = m.ExcPending
			if v330 != 0 {
				return int32(0)
			} else {
				F_abort(m)
				mBase = m.M
				base.Wasm_trap_unreachable()
				for {
				}
			}
		} else {
			switch int32(base.Ui32(v14)>>(uint(int32(4))%32)) & int32(15) {
			case 0, 8:
				v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
				if v21&int32(4) == int32(0) {
					v81 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
					v82 = v21
					v84 = v81
				} else {
					if v21&int32(1) != 0 {
						v30 = int32(16)
					} else {
						v30 = int32(8)
					}
					v31 = l0 + v30
					if v21&int32(2) == int32(0) {
						v63 = v31
					} else {
						v36 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v31))))
						v37 = v31 + v36
						v41 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v37))))
						switch v41 & int32(7) {
						case 0:
							v58 = int32(base.Ui32(v41) >> (uint(int32(3)) % 32))
						case 1:
							v48 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v37+int32(-2)))))
							v58 = v48
						case 2:
							v51 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v37+int32(-4)))))
							v58 = v51
						case 3:
							v54 = *(*int32)(unsafe.Add(mBase, uint32(v37+int32(-8))))
							v58 = v54
						case 4:
							v57 = *(*int32)(unsafe.Add(mBase, uint32(v37+int32(-16))))
							v58 = v57
						default:
							v58 = int32(0)
						}
						v63 = v37 + int32(1) + v58 + int32(1)
					}
					v77 = *(*int32)(unsafe.Add(mBase, _consts[249]))
					v80 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
					v82 = v80
					v84 = v63 + v77
				}
				if v82&int32(4) == int32(0) {
					v144 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
					v147 = v144
				} else {
					if v82&int32(1) != 0 {
						v94 = int32(16)
					} else {
						v94 = int32(8)
					}
					v95 = l0 + v94
					if v82&int32(2) == int32(0) {
						v126 = v95
					} else {
						v100 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v95))))
						v101 = v95 + v100
						v105 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v101))))
						switch v105 & int32(7) {
						case 0:
							v122 = int32(base.Ui32(v105) >> (uint(int32(3)) % 32))
						case 1:
							v112 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v101+int32(-2)))))
							v122 = v112
						case 2:
							v115 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v101+int32(-4)))))
							v122 = v115
						case 3:
							v118 = *(*int32)(unsafe.Add(mBase, uint32(v101+int32(-8))))
							v122 = v118
						case 4:
							v121 = *(*int32)(unsafe.Add(mBase, uint32(v101+int32(-16))))
							v122 = v121
						default:
							v122 = int32(0)
						}
						v126 = v101 + int32(1) + v122 + int32(1)
					}
					v141 = *(*int32)(unsafe.Add(mBase, _consts[249]))
					v147 = v126 + v141
				}
				v149 = int32(-1)
				v153 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v147+v149))))
				switch v153 & int32(7) {
				case 0:
					v170 = int32(base.Ui32(v153) >> (uint(int32(3)) % 32))
				case 1:
					v160 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v147+int32(-3)))))
					v170 = v160
				case 2:
					v163 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v147+int32(-5)))))
					v170 = v163
				case 3:
					v166 = *(*int32)(unsafe.Add(mBase, uint32(v147+int32(-9))))
					v170 = v166
				case 4:
					v169 = *(*int32)(unsafe.Add(mBase, uint32(v147+int32(-17))))
					v170 = v169
				default:
					v170 = int32(0)
				}
				v172 = v11 + int32(8)
				v173 = int32(0)
				v178 = m.G0
				v180 = v178 - int32(16)
				m.G0 = v180
				v182 = F___errno_location(m)
				mBase = m.M
				*(*int32)(unsafe.Add(mBase, uint32(v182))) = v173
				v187 = F_valkey_strtod_n(m, v84, v170, v180+int32(12))
				mBase = m.M
				*(*float64)(unsafe.Add(mBase, uint32(v172))) = v187
				if v170 == v173 {
					v224 = int32(0)
					*(*int32)(unsafe.Add(mBase, uint32(v182))) = v224
					v227 = v224
				} else {
					v191 = int32(*(*int8)(unsafe.Add(mBase, uint32(v84))))
					if v191 == int32(32) {
						v224 = int32(0)
						*(*int32)(unsafe.Add(mBase, uint32(v182))) = v224
						v227 = v224
					} else {
						if base.Ui32(int32(-6)) < base.Ui32(v191+int32(-14)) {
							v224 = int32(0)
							*(*int32)(unsafe.Add(mBase, uint32(v182))) = v224
							v227 = v224
						} else {
							v198 = *(*int32)(unsafe.Add(mBase, uint32(v180)+12))
							if v198-v84 != v170 {
								v224 = int32(0)
								*(*int32)(unsafe.Add(mBase, uint32(v182))) = v224
								v227 = v224
							} else {
								v201 = *(*int32)(unsafe.Add(mBase, uint32(v182)))
								if v201 == int32(68) {
									if base.F64_eq(base.F64_abs(v187), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
										v224 = int32(0)
										*(*int32)(unsafe.Add(mBase, uint32(v182))) = v224
										v227 = v224
									} else {
										v208 = F___fpclassify(m, v187)
										mBase = m.M
										if v208 == int32(2) {
											v224 = int32(0)
											*(*int32)(unsafe.Add(mBase, uint32(v182))) = v224
											v227 = v224
										} else {
											v211 = *(*int64)(unsafe.Add(mBase, uint32(v172)))
											v212 = v211
											if base.Ui64(int64(9218868437227405312)) < base.Ui64(v212&int64(9223372036854775807)) {
												v224 = int32(0)
												*(*int32)(unsafe.Add(mBase, uint32(v182))) = v224
												v227 = v224
											} else {
												v218 = *(*int32)(unsafe.Add(mBase, uint32(v182)))
												if v218 != int32(28) {
													v227 = int32(1)
												} else {
													v224 = int32(0)
													*(*int32)(unsafe.Add(mBase, uint32(v182))) = v224
													v227 = v224
												}
											}
										}
									}
								} else {
									v212 = base.I64_reinterpret_f64(v187)
									if base.Ui64(int64(9218868437227405312)) < base.Ui64(v212&int64(9223372036854775807)) {
										v224 = int32(0)
										*(*int32)(unsafe.Add(mBase, uint32(v182))) = v224
										v227 = v224
									} else {
										v218 = *(*int32)(unsafe.Add(mBase, uint32(v182)))
										if v218 != int32(28) {
											v227 = int32(1)
										} else {
											v224 = int32(0)
											*(*int32)(unsafe.Add(mBase, uint32(v182))) = v224
											v227 = v224
										}
									}
								}
							}
						}
					}
				}
				m.G0 = v180 + int32(16)
				if v227 == int32(0) {
					v315 = v149
				} else {
					v235 = *(*float64)(unsafe.Add(mBase, uint32(v11)+8))
					v308 = v235
					*(*float64)(unsafe.Add(mBase, uint32(l1))) = v308
					v315 = int32(0)
				}
				m.G0 = v11 + int32(16)
				return v315
			case 1:
				v236 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
				if v236&int32(4) == int32(0) {
					v305 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
					v308 = base.F64_convert_i32_s(v305)
				} else {
					if v236&int32(1) != 0 {
						v245 = int32(16)
					} else {
						v245 = int32(8)
					}
					v246 = l0 + v245
					if v236&int32(2) == int32(0) {
						v277 = v246
					} else {
						v251 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v246))))
						v252 = v246 + v251
						v256 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v252))))
						switch v256 & int32(7) {
						case 0:
							v273 = int32(base.Ui32(v256) >> (uint(int32(3)) % 32))
						case 1:
							v263 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v252+int32(-2)))))
							v273 = v263
						case 2:
							v266 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v252+int32(-4)))))
							v273 = v266
						case 3:
							v269 = *(*int32)(unsafe.Add(mBase, uint32(v252+int32(-8))))
							v273 = v269
						case 4:
							v272 = *(*int32)(unsafe.Add(mBase, uint32(v252+int32(-16))))
							v273 = v272
						default:
							v273 = int32(0)
						}
						v277 = v252 + int32(1) + v273 + int32(1)
					}
					v292 = *(*int32)(unsafe.Add(mBase, _consts[249]))
					v308 = base.F64_convert_i32_s(v277 + v292)
				}
				*(*float64)(unsafe.Add(mBase, uint32(l1))) = v308
				v315 = int32(0)
				m.G0 = v11 + int32(16)
				return v315
			default:
				F__serverPanic_1(m, int32(_a838), int32(1038), int32(_a848), int32(0))
				mBase = m.M
				v303 = m.ExcPending
				if v303 != 0 {
					return int32(0)
				} else {
					F_abort(m)
					mBase = m.M
					base.Wasm_trap_unreachable()
					for {
					}
				}
			}
		}
	} else {
		v308 = float64(0)
		*(*float64)(unsafe.Add(mBase, uint32(l1))) = v308
		v315 = int32(0)
		m.G0 = v11 + int32(16)
		return v315
	}
}
