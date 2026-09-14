package p1

import (
	base "github.com/shibukawa/vkmem/internal/aot/vkaot/base"
	"unsafe"
)

func F_entryFree(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v16 int32
	_ = v16
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	v5 = l0 + int32(-1)
	v6 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5))))
	if v6&int32(7) == int32(0) {
		v22 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0+int32(-1)))))
		v24 = v22 & int32(7)
		if base.Ui32(int32(4)) < base.Ui32(v24) {
			v32 = int32(0)
		} else {
			v31 = *(*int32)(unsafe.Add(mBase, uint32(v24<<(uint(int32(2))%32))+uint32(_consts[307])))
			v32 = v31
		}
		v36 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5))))
		if int32(base.Ui32(v36&int32(16))>>(uint(int32(4))%32)) != 0 {
			v41 = int32(-4)
		} else {
			v41 = int32(0)
		}
		v44 = v36 & int32(7)
		if v44 != 0 {
			v45 = v41
		} else {
			v45 = int32(0)
		}
		if int32(base.Ui32(v36&int32(8))>>(uint(int32(3))%32)) != 0 {
			v53 = int32(-8)
		} else {
			v53 = int32(0)
		}
		if v44 != 0 {
			v55 = v53
		} else {
			v55 = int32(0)
		}
		F_valkey_free(m, l0+v32+v45+v55)
		mBase = m.M
		v58 = m.ExcPending
		if v58 != 0 {
			return
		} else {
			return
		}
	} else {
		if v6&int32(16) == int32(0) {
			v22 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0+int32(-1)))))
			v24 = v22 & int32(7)
			if base.Ui32(int32(4)) < base.Ui32(v24) {
				v32 = int32(0)
			} else {
				v31 = *(*int32)(unsafe.Add(mBase, uint32(v24<<(uint(int32(2))%32))+uint32(_consts[307])))
				v32 = v31
			}
			v36 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5))))
			if int32(base.Ui32(v36&int32(16))>>(uint(int32(4))%32)) != 0 {
				v41 = int32(-4)
			} else {
				v41 = int32(0)
			}
			v44 = v36 & int32(7)
			if v44 != 0 {
				v45 = v41
			} else {
				v45 = int32(0)
			}
			if int32(base.Ui32(v36&int32(8))>>(uint(int32(3))%32)) != 0 {
				v53 = int32(-8)
			} else {
				v53 = int32(0)
			}
			if v44 != 0 {
				v55 = v53
			} else {
				v55 = int32(0)
			}
			F_valkey_free(m, l0+v32+v45+v55)
			mBase = m.M
			v58 = m.ExcPending
			if v58 != 0 {
				return
			} else {
				return
			}
		} else {
			F_entryFreeValuePtr(m, l0)
			mBase = m.M
			v16 = m.ExcPending
			if v16 != 0 {
				return
			} else {
				v22 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0+int32(-1)))))
				v24 = v22 & int32(7)
				if base.Ui32(int32(4)) < base.Ui32(v24) {
					v32 = int32(0)
				} else {
					v31 = *(*int32)(unsafe.Add(mBase, uint32(v24<<(uint(int32(2))%32))+uint32(_consts[307])))
					v32 = v31
				}
				v36 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5))))
				if int32(base.Ui32(v36&int32(16))>>(uint(int32(4))%32)) != 0 {
					v41 = int32(-4)
				} else {
					v41 = int32(0)
				}
				v44 = v36 & int32(7)
				if v44 != 0 {
					v45 = v41
				} else {
					v45 = int32(0)
				}
				if int32(base.Ui32(v36&int32(8))>>(uint(int32(3))%32)) != 0 {
					v53 = int32(-8)
				} else {
					v53 = int32(0)
				}
				if v44 != 0 {
					v55 = v53
				} else {
					v55 = int32(0)
				}
				F_valkey_free(m, l0+v32+v45+v55)
				mBase = m.M
				v58 = m.ExcPending
				if v58 != 0 {
					return
				} else {
					return
				}
			}
		}
	}
}
func F_entryHasExpiry(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	v4 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0+int32(-1)))))
	return base.B2i32(v4&int32(7) != int32(0)) & int32(base.Ui32(v4&int32(8))>>(uint(int32(3))%32))
}
func F_entryIsExpired(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v18 int32
	_ = v18
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v38 int32
	_ = v38
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v59 int64
	_ = v59
	var v65 int64
	_ = v65
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	v2 = int32(0)
	v8 = l0 + int32(-1)
	v9 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8))))
	if v9&int32(7) == v2 {
		v69 = v2
	} else {
		if v9&int32(8) == int32(0) {
			v69 = v2
		} else {
			v18 = int32(0)
			v24 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0+int32(-1)))))
			v26 = v24 & int32(7)
			if base.Ui32(int32(4)) < base.Ui32(v26) {
				v34 = v18
			} else {
				v33 = *(*int32)(unsafe.Add(mBase, uint32(v26<<(uint(int32(2))%32))+uint32(_consts[307])))
				v34 = v33
			}
			v38 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8))))
			if int32(base.Ui32(v38&int32(16))>>(uint(int32(4))%32)) != 0 {
				v43 = int32(-4)
			} else {
				v43 = int32(0)
			}
			v46 = v38 & int32(7)
			if v46 != 0 {
				v47 = v43
			} else {
				v47 = int32(0)
			}
			if int32(base.Ui32(v38&int32(8))>>(uint(int32(3))%32)) != 0 {
				v55 = int32(-8)
			} else {
				v55 = int32(0)
			}
			if v46 != 0 {
				v57 = v55
			} else {
				v57 = int32(0)
			}
			v59 = *(*int64)(unsafe.Add(mBase, uint32(l0+v34+v47+v57)))
			if v59 == int64(-1) {
				v69 = v18
			} else {
				if int64(0) <= v59 {
					v65 = F_commandTimeSnapshot(m)
					mBase = m.M
					v67 = base.B2i32(v59 < v65)
				} else {
					v67 = int32(0)
				}
				v69 = v67
			}
		}
	}
	return v69
}
func F_entryMemUsage(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v34 int32
	_ = v34
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v76 int32
	_ = v76
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	var v99 int32
	_ = v99
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v105 int32
	_ = v105
	var v108 int32
	_ = v108
	var v111 int32
	_ = v111
	var v120 int32
	_ = v120
	var v122 int32
	_ = v122
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v134 int32
	_ = v134
	var v137 int32
	_ = v137
	var v153 int32
	_ = v153
	var v156 int32
	_ = v156
	var v159 int32
	_ = v159
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v176 int32
	_ = v176
	var v186 int32
	_ = v186
	var v188 int32
	_ = v188
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v200 int32
	_ = v200
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v224 int32
	_ = v224
	var v231 int32
	_ = v231
	var v235 int32
	_ = v235
	var v239 int32
	_ = v239
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v248 int32
	_ = v248
	v7 = l0 + int32(-1)
	v8 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7))))
	v10 = v8 & int32(7)
	if v10 == int32(0) {
		switch v10 {
		case 0:
			v30 = int32(base.Ui32(v8) >> (uint(int32(3)) % 32))
		case 1:
			v20 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0+int32(-3)))))
			v30 = v20
		case 2:
			v23 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0+int32(-5)))))
			v30 = v23
		case 3:
			v26 = *(*int32)(unsafe.Add(mBase, uint32(l0+int32(-9))))
			v30 = v26
		case 4:
			v29 = *(*int32)(unsafe.Add(mBase, uint32(l0+int32(-17))))
			v30 = v29
		default:
			v30 = int32(0)
		}
		v34 = v10 & int32(7)
		if base.Ui32(int32(4)) < base.Ui32(v34) {
			v42 = int32(0)
		} else {
			v41 = *(*int32)(unsafe.Add(mBase, uint32(v34<<(uint(int32(2))%32))+uint32(_consts[309])))
			v42 = v41
		}
		v43 = v30 + v42
		v47 = v43 + int32(1)
		v48 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7))))
		if int32(base.Ui32(v48&int32(8))>>(uint(int32(3))%32)) != 0 {
			v53 = v43 + int32(9)
		} else {
			v53 = v47
		}
		if v48&int32(7) != 0 {
			v56 = v53
		} else {
			v56 = v47
		}
		v103 = v48
		v105 = v56
	} else {
		if v8&int32(16) != 0 {
			v62 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0+int32(-1)))))
			v64 = v62 & int32(7)
			if base.Ui32(int32(4)) < base.Ui32(v64) {
				v72 = int32(0)
			} else {
				v71 = *(*int32)(unsafe.Add(mBase, uint32(v64<<(uint(int32(2))%32))+uint32(_consts[307])))
				v72 = v71
			}
			v76 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7))))
			if int32(base.Ui32(v76&int32(16))>>(uint(int32(4))%32)) != 0 {
				v81 = int32(-4)
			} else {
				v81 = int32(0)
			}
			v84 = v76 & int32(7)
			if v84 != 0 {
				v85 = v81
			} else {
				v85 = int32(0)
			}
			if int32(base.Ui32(v76&int32(8))>>(uint(int32(3))%32)) != 0 {
				v93 = int32(-8)
			} else {
				v93 = int32(0)
			}
			if v84 != 0 {
				v95 = v93
			} else {
				v95 = int32(0)
			}
			v99 = *(*int32)(unsafe.Add(mBase, uint32(l0+v72+v85+v95+int32(-8))))
			v102 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7))))
			v103 = v102
			v105 = v99 & int32(2147483647)
		} else {
			switch v10 {
			case 0:
				v30 = int32(base.Ui32(v8) >> (uint(int32(3)) % 32))
			case 1:
				v20 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0+int32(-3)))))
				v30 = v20
			case 2:
				v23 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0+int32(-5)))))
				v30 = v23
			case 3:
				v26 = *(*int32)(unsafe.Add(mBase, uint32(l0+int32(-9))))
				v30 = v26
			case 4:
				v29 = *(*int32)(unsafe.Add(mBase, uint32(l0+int32(-17))))
				v30 = v29
			default:
				v30 = int32(0)
			}
			v34 = v10 & int32(7)
			if base.Ui32(int32(4)) < base.Ui32(v34) {
				v42 = int32(0)
			} else {
				v41 = *(*int32)(unsafe.Add(mBase, uint32(v34<<(uint(int32(2))%32))+uint32(_consts[309])))
				v42 = v41
			}
			v43 = v30 + v42
			v47 = v43 + int32(1)
			v48 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7))))
			if int32(base.Ui32(v48&int32(8))>>(uint(int32(3))%32)) != 0 {
				v53 = v43 + int32(9)
			} else {
				v53 = v47
			}
			if v48&int32(7) != 0 {
				v56 = v53
			} else {
				v56 = v47
			}
			v103 = v48
			v105 = v56
		}
	}
	v108 = v103 & int32(7)
	if v108 == int32(0) {
		if v108 == int32(0) {
			switch v108 {
			case 0:
				v163 = int32(base.Ui32(v103&int32(248)) >> (uint(int32(3)) % 32))
			case 1:
				v153 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0+int32(-3)))))
				v163 = v153
			case 2:
				v156 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0+int32(-5)))))
				v163 = v156
			case 3:
				v159 = *(*int32)(unsafe.Add(mBase, uint32(l0+int32(-9))))
				v163 = v159
			case 4:
				v162 = *(*int32)(unsafe.Add(mBase, uint32(l0+int32(-17))))
				v163 = v162
			default:
				v163 = int32(0)
			}
			v176 = *(*int32)(unsafe.Add(mBase, _consts[308]))
			v207 = l0 + v163 + v176 + int32(1)
		} else {
			if v103&int32(16) != 0 {
				v186 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0+int32(-1)))))
				v188 = v186 & int32(7)
				if base.Ui32(int32(4)) < base.Ui32(v188) {
					v196 = int32(0)
				} else {
					v195 = *(*int32)(unsafe.Add(mBase, uint32(v188<<(uint(int32(2))%32))+uint32(_consts[307])))
					v196 = v195
				}
				v200 = *(*int32)(unsafe.Add(mBase, uint32(l0+v196+int32(-4))))
				if v103&int32(32) == int32(0) {
					v207 = v200
				} else {
					if v200 != 0 {
						v206 = *(*int32)(unsafe.Add(mBase, uint32(v200)))
						v207 = v206
					} else {
						v207 = int32(0)
					}
				}
			} else {
				switch v108 {
				case 0:
					v163 = int32(base.Ui32(v103&int32(248)) >> (uint(int32(3)) % 32))
				case 1:
					v153 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0+int32(-3)))))
					v163 = v153
				case 2:
					v156 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0+int32(-5)))))
					v163 = v156
				case 3:
					v159 = *(*int32)(unsafe.Add(mBase, uint32(l0+int32(-9))))
					v163 = v159
				case 4:
					v162 = *(*int32)(unsafe.Add(mBase, uint32(l0+int32(-17))))
					v163 = v162
				default:
					v163 = int32(0)
				}
				v176 = *(*int32)(unsafe.Add(mBase, _consts[308]))
				v207 = l0 + v163 + v176 + int32(1)
			}
		}
		v215 = v207 + int32(-1)
		v216 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v215))))
		v218 = v216 & int32(7)
		switch v218 {
		case 0:
			v219 = F_zmalloc_usable_size(m, v215)
			mBase = m.M
			v248 = v219
		case 1:
			v224 = int32(4)
			switch v218 {
			case 0:
				v248 = v224 + int32(base.Ui32(v216)>>(uint(int32(3))%32))
			case 1:
				v231 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v207+int32(-2)))))
				v248 = v224 + v231
			case 2:
				v235 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v207+int32(-3)))))
				v248 = v224 + v235
			case 3:
				v239 = *(*int32)(unsafe.Add(mBase, uint32(v207+int32(-5))))
				v248 = v224 + v239
			case 4:
				v243 = *(*int32)(unsafe.Add(mBase, uint32(v207+int32(-9))))
				v244 = v243
				v248 = v224 + v244
			default:
				v244 = int32(0)
				v248 = v224 + v244
			}
		case 2:
			v224 = int32(6)
			switch v218 {
			case 0:
				v248 = v224 + int32(base.Ui32(v216)>>(uint(int32(3))%32))
			case 1:
				v231 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v207+int32(-2)))))
				v248 = v224 + v231
			case 2:
				v235 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v207+int32(-3)))))
				v248 = v224 + v235
			case 3:
				v239 = *(*int32)(unsafe.Add(mBase, uint32(v207+int32(-5))))
				v248 = v224 + v239
			case 4:
				v243 = *(*int32)(unsafe.Add(mBase, uint32(v207+int32(-9))))
				v244 = v243
				v248 = v224 + v244
			default:
				v244 = int32(0)
				v248 = v224 + v244
			}
		case 3:
			v224 = int32(10)
			switch v218 {
			case 0:
				v248 = v224 + int32(base.Ui32(v216)>>(uint(int32(3))%32))
			case 1:
				v231 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v207+int32(-2)))))
				v248 = v224 + v231
			case 2:
				v235 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v207+int32(-3)))))
				v248 = v224 + v235
			case 3:
				v239 = *(*int32)(unsafe.Add(mBase, uint32(v207+int32(-5))))
				v248 = v224 + v239
			case 4:
				v243 = *(*int32)(unsafe.Add(mBase, uint32(v207+int32(-9))))
				v244 = v243
				v248 = v224 + v244
			default:
				v244 = int32(0)
				v248 = v224 + v244
			}
		case 4:
			v224 = int32(18)
			switch v218 {
			case 0:
				v248 = v224 + int32(base.Ui32(v216)>>(uint(int32(3))%32))
			case 1:
				v231 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v207+int32(-2)))))
				v248 = v224 + v231
			case 2:
				v235 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v207+int32(-3)))))
				v248 = v224 + v235
			case 3:
				v239 = *(*int32)(unsafe.Add(mBase, uint32(v207+int32(-5))))
				v248 = v224 + v239
			case 4:
				v243 = *(*int32)(unsafe.Add(mBase, uint32(v207+int32(-9))))
				v244 = v243
				v248 = v224 + v244
			default:
				v244 = int32(0)
				v248 = v224 + v244
			}
		default:
			v224 = int32(1)
			switch v218 {
			case 0:
				v248 = v224 + int32(base.Ui32(v216)>>(uint(int32(3))%32))
			case 1:
				v231 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v207+int32(-2)))))
				v248 = v224 + v231
			case 2:
				v235 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v207+int32(-3)))))
				v248 = v224 + v235
			case 3:
				v239 = *(*int32)(unsafe.Add(mBase, uint32(v207+int32(-5))))
				v248 = v224 + v239
			case 4:
				v243 = *(*int32)(unsafe.Add(mBase, uint32(v207+int32(-9))))
				v244 = v243
				v248 = v224 + v244
			default:
				v244 = int32(0)
				v248 = v224 + v244
			}
		}
		return v248 + v105
	} else {
		v111 = int32(48)
		if v103&v111 != v111 {
			if v108 == int32(0) {
				switch v108 {
				case 0:
					v163 = int32(base.Ui32(v103&int32(248)) >> (uint(int32(3)) % 32))
				case 1:
					v153 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0+int32(-3)))))
					v163 = v153
				case 2:
					v156 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0+int32(-5)))))
					v163 = v156
				case 3:
					v159 = *(*int32)(unsafe.Add(mBase, uint32(l0+int32(-9))))
					v163 = v159
				case 4:
					v162 = *(*int32)(unsafe.Add(mBase, uint32(l0+int32(-17))))
					v163 = v162
				default:
					v163 = int32(0)
				}
				v176 = *(*int32)(unsafe.Add(mBase, _consts[308]))
				v207 = l0 + v163 + v176 + int32(1)
			} else {
				if v103&int32(16) != 0 {
					v186 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0+int32(-1)))))
					v188 = v186 & int32(7)
					if base.Ui32(int32(4)) < base.Ui32(v188) {
						v196 = int32(0)
					} else {
						v195 = *(*int32)(unsafe.Add(mBase, uint32(v188<<(uint(int32(2))%32))+uint32(_consts[307])))
						v196 = v195
					}
					v200 = *(*int32)(unsafe.Add(mBase, uint32(l0+v196+int32(-4))))
					if v103&int32(32) == int32(0) {
						v207 = v200
					} else {
						if v200 != 0 {
							v206 = *(*int32)(unsafe.Add(mBase, uint32(v200)))
							v207 = v206
						} else {
							v207 = int32(0)
						}
					}
				} else {
					switch v108 {
					case 0:
						v163 = int32(base.Ui32(v103&int32(248)) >> (uint(int32(3)) % 32))
					case 1:
						v153 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0+int32(-3)))))
						v163 = v153
					case 2:
						v156 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0+int32(-5)))))
						v163 = v156
					case 3:
						v159 = *(*int32)(unsafe.Add(mBase, uint32(l0+int32(-9))))
						v163 = v159
					case 4:
						v162 = *(*int32)(unsafe.Add(mBase, uint32(l0+int32(-17))))
						v163 = v162
					default:
						v163 = int32(0)
					}
					v176 = *(*int32)(unsafe.Add(mBase, _consts[308]))
					v207 = l0 + v163 + v176 + int32(1)
				}
			}
			v215 = v207 + int32(-1)
			v216 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v215))))
			v218 = v216 & int32(7)
			switch v218 {
			case 0:
				v219 = F_zmalloc_usable_size(m, v215)
				mBase = m.M
				v248 = v219
			case 1:
				v224 = int32(4)
				switch v218 {
				case 0:
					v248 = v224 + int32(base.Ui32(v216)>>(uint(int32(3))%32))
				case 1:
					v231 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v207+int32(-2)))))
					v248 = v224 + v231
				case 2:
					v235 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v207+int32(-3)))))
					v248 = v224 + v235
				case 3:
					v239 = *(*int32)(unsafe.Add(mBase, uint32(v207+int32(-5))))
					v248 = v224 + v239
				case 4:
					v243 = *(*int32)(unsafe.Add(mBase, uint32(v207+int32(-9))))
					v244 = v243
					v248 = v224 + v244
				default:
					v244 = int32(0)
					v248 = v224 + v244
				}
			case 2:
				v224 = int32(6)
				switch v218 {
				case 0:
					v248 = v224 + int32(base.Ui32(v216)>>(uint(int32(3))%32))
				case 1:
					v231 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v207+int32(-2)))))
					v248 = v224 + v231
				case 2:
					v235 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v207+int32(-3)))))
					v248 = v224 + v235
				case 3:
					v239 = *(*int32)(unsafe.Add(mBase, uint32(v207+int32(-5))))
					v248 = v224 + v239
				case 4:
					v243 = *(*int32)(unsafe.Add(mBase, uint32(v207+int32(-9))))
					v244 = v243
					v248 = v224 + v244
				default:
					v244 = int32(0)
					v248 = v224 + v244
				}
			case 3:
				v224 = int32(10)
				switch v218 {
				case 0:
					v248 = v224 + int32(base.Ui32(v216)>>(uint(int32(3))%32))
				case 1:
					v231 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v207+int32(-2)))))
					v248 = v224 + v231
				case 2:
					v235 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v207+int32(-3)))))
					v248 = v224 + v235
				case 3:
					v239 = *(*int32)(unsafe.Add(mBase, uint32(v207+int32(-5))))
					v248 = v224 + v239
				case 4:
					v243 = *(*int32)(unsafe.Add(mBase, uint32(v207+int32(-9))))
					v244 = v243
					v248 = v224 + v244
				default:
					v244 = int32(0)
					v248 = v224 + v244
				}
			case 4:
				v224 = int32(18)
				switch v218 {
				case 0:
					v248 = v224 + int32(base.Ui32(v216)>>(uint(int32(3))%32))
				case 1:
					v231 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v207+int32(-2)))))
					v248 = v224 + v231
				case 2:
					v235 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v207+int32(-3)))))
					v248 = v224 + v235
				case 3:
					v239 = *(*int32)(unsafe.Add(mBase, uint32(v207+int32(-5))))
					v248 = v224 + v239
				case 4:
					v243 = *(*int32)(unsafe.Add(mBase, uint32(v207+int32(-9))))
					v244 = v243
					v248 = v224 + v244
				default:
					v244 = int32(0)
					v248 = v224 + v244
				}
			default:
				v224 = int32(1)
				switch v218 {
				case 0:
					v248 = v224 + int32(base.Ui32(v216)>>(uint(int32(3))%32))
				case 1:
					v231 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v207+int32(-2)))))
					v248 = v224 + v231
				case 2:
					v235 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v207+int32(-3)))))
					v248 = v224 + v235
				case 3:
					v239 = *(*int32)(unsafe.Add(mBase, uint32(v207+int32(-5))))
					v248 = v224 + v239
				case 4:
					v243 = *(*int32)(unsafe.Add(mBase, uint32(v207+int32(-9))))
					v244 = v243
					v248 = v224 + v244
				default:
					v244 = int32(0)
					v248 = v224 + v244
				}
			}
			return v248 + v105
		} else {
			v120 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0+int32(-1)))))
			v122 = v120 & int32(7)
			if base.Ui32(int32(4)) < base.Ui32(v122) {
				v130 = int32(0)
			} else {
				v129 = *(*int32)(unsafe.Add(mBase, uint32(v122<<(uint(int32(2))%32))+uint32(_consts[307])))
				v130 = v129
			}
			v134 = *(*int32)(unsafe.Add(mBase, uint32(l0+v130+int32(-4))))
			v137 = *(*int32)(unsafe.Add(mBase, uint32(v134+int32(-8))))
			return v137&int32(2147483647) + v105
		}
	}
}
func F_entryUpdate(m *base.Module, l0 int32, l1 int32, l2 int64) int32 {
	mBase := m.M
	_ = mBase
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v34 int32
	_ = v34
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int64
	_ = v66
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v92 int32
	_ = v92
	var v97 int32
	_ = v97
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
	var v113 int64
	_ = v113
	var v114 int32
	_ = v114
	var v116 int64
	_ = v116
	var v117 int32
	_ = v117
	var v124 int32
	_ = v124
	var v132 int32
	_ = v132
	var v136 int32
	_ = v136
	var v140 int32
	_ = v140
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v148 int32
	_ = v148
	var v154 int32
	_ = v154
	var v156 int32
	_ = v156
	var v166 int32
	_ = v166
	var v169 int32
	_ = v169
	var v172 int32
	_ = v172
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v182 int32
	_ = v182
	var v185 int32
	_ = v185
	var v188 int32
	_ = v188
	var v196 int32
	_ = v196
	var v200 int32
	_ = v200
	var v204 int32
	_ = v204
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v211 int32
	_ = v211
	var v213 int32
	_ = v213
	var v218 int32
	_ = v218
	var v222 int32
	_ = v222
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v231 int32
	_ = v231
	var v238 int32
	_ = v238
	var v241 int32
	_ = v241
	var v244 int32
	_ = v244
	var v247 int32
	_ = v247
	var v248 int32
	_ = v248
	var v250 int32
	_ = v250
	var v258 int32
	_ = v258
	var v259 int32
	_ = v259
	var v260 int32
	_ = v260
	var v261 int32
	_ = v261
	var v273 int32
	_ = v273
	var v276 int32
	_ = v276
	var v279 int32
	_ = v279
	var v282 int32
	_ = v282
	var v283 int32
	_ = v283
	var v284 int32
	_ = v284
	var v296 int32
	_ = v296
	var v297 int32
	_ = v297
	var v299 int32
	_ = v299
	var v301 int32
	_ = v301
	var v302 int32
	_ = v302
	var v306 int32
	_ = v306
	var v313 int32
	_ = v313
	var v314 int32
	_ = v314
	var v316 int32
	_ = v316
	var v317 int32
	_ = v317
	var v319 int32
	_ = v319
	var v323 int32
	_ = v323
	var v336 int32
	_ = v336
	var v340 int32
	_ = v340
	var v341 int32
	_ = v341
	var v342 int32
	_ = v342
	var v345 int32
	_ = v345
	var v362 int32
	_ = v362
	var v364 int32
	_ = v364
	var v368 int32
	_ = v368
	var v369 int32
	_ = v369
	var v370 int32
	_ = v370
	var v372 int32
	_ = v372
	var v373 int32
	_ = v373
	var v374 int32
	_ = v374
	var v388 int32
	_ = v388
	var v389 int32
	_ = v389
	var v391 int32
	_ = v391
	var v401 int32
	_ = v401
	var v404 int32
	_ = v404
	var v407 int32
	_ = v407
	var v410 int32
	_ = v410
	var v411 int32
	_ = v411
	var v412 int32
	_ = v412
	var v413 int32
	_ = v413
	var v417 int32
	_ = v417
	var v418 int32
	_ = v418
	var v423 int32
	_ = v423
	var v426 int32
	_ = v426
	var v427 int32
	_ = v427
	var v430 int32
	_ = v430
	var v435 int32
	_ = v435
	var v438 int32
	_ = v438
	var v439 int32
	_ = v439
	var v447 int32
	_ = v447
	var v449 int32
	_ = v449
	var v451 int32
	_ = v451
	var v452 int32
	_ = v452
	var v453 int32
	_ = v453
	var v455 int32
	_ = v455
	var v458 int32
	_ = v458
	var v461 int32
	_ = v461
	var v465 int32
	_ = v465
	var v468 int32
	_ = v468
	var v469 int32
	_ = v469
	var v482 int32
	_ = v482
	var v485 int32
	_ = v485
	var v488 int32
	_ = v488
	var v491 int32
	_ = v491
	var v492 int32
	_ = v492
	var v494 int32
	_ = v494
	var v495 int32
	_ = v495
	var v499 int32
	_ = v499
	var v502 int32
	_ = v502
	var v508 int32
	_ = v508
	var v509 int32
	_ = v509
	var v511 int32
	_ = v511
	var v515 int32
	_ = v515
	var v516 int32
	_ = v516
	var v517 int64
	_ = v517
	var v521 int32
	_ = v521
	var v526 int32
	_ = v526
	var v533 int32
	_ = v533
	var v534 int32
	_ = v534
	var v536 int32
	_ = v536
	var v537 int32
	_ = v537
	var v540 int32
	_ = v540
	var v546 int32
	_ = v546
	var v548 int32
	_ = v548
	var v566 int32
	_ = v566
	var v580 int32
	_ = v580
	var v582 int32
	_ = v582
	var v589 int32
	_ = v589
	var v590 int32
	_ = v590
	var v594 int32
	_ = v594
	var v599 int32
	_ = v599
	var v602 int32
	_ = v602
	var v603 int32
	_ = v603
	var v611 int32
	_ = v611
	var v613 int32
	_ = v613
	var v620 int32
	_ = v620
	var v622 int32
	_ = v622
	var v627 int32
	_ = v627
	var v633 int32
	_ = v633
	var v636 int32
	_ = v636
	var v639 int32
	_ = v639
	var v642 int32
	_ = v642
	var v643 int32
	_ = v643
	var v655 int32
	_ = v655
	var v668 int32
	_ = v668
	var v671 int32
	_ = v671
	var v674 int32
	_ = v674
	var v681 int32
	_ = v681
	var v684 int32
	_ = v684
	var v687 int32
	_ = v687
	var v690 int32
	_ = v690
	var v691 int32
	_ = v691
	var v692 int32
	_ = v692
	var v697 int32
	_ = v697
	var v699 int32
	_ = v699
	var v706 int32
	_ = v706
	var v707 int32
	_ = v707
	var v715 int32
	_ = v715
	var v722 int32
	_ = v722
	var v725 int32
	_ = v725
	var v728 int32
	_ = v728
	var v731 int32
	_ = v731
	var v732 int32
	_ = v732
	var v734 int32
	_ = v734
	var v735 int32
	_ = v735
	var v737 int32
	_ = v737
	var v739 int32
	_ = v739
	var v740 int32
	_ = v740
	var v742 int32
	_ = v742
	var v745 int32
	_ = v745
	var v750 int32
	_ = v750
	var v752 int32
	_ = v752
	var v764 int32
	_ = v764
	var v766 int32
	_ = v766
	var v773 int32
	_ = v773
	var v774 int32
	_ = v774
	var v784 int32
	_ = v784
	var v786 int32
	_ = v786
	var v793 int32
	_ = v793
	var v794 int32
	_ = v794
	var v800 int32
	_ = v800
	var v805 int32
	_ = v805
	var v806 int32
	_ = v806
	var v807 int32
	_ = v807
	var v817 int32
	_ = v817
	var v823 int32
	_ = v823
	var v825 int32
	_ = v825
	var v832 int32
	_ = v832
	var v833 int32
	_ = v833
	var v837 int32
	_ = v837
	var v842 int32
	_ = v842
	var v845 int32
	_ = v845
	var v846 int32
	_ = v846
	var v854 int32
	_ = v854
	var v856 int32
	_ = v856
	var v859 int32
	_ = v859
	var v860 int32
	_ = v860
	var v869 int32
	_ = v869
	var v874 int32
	_ = v874
	var v876 int32
	_ = v876
	var v890 int32
	_ = v890
	var v898 int32
	_ = v898
	var v902 int32
	_ = v902
	var v925 int32
	_ = v925
	var v931 int32
	_ = v931
	var v937 int32
	_ = v937
	var v943 int32
	_ = v943
	v18 = m.G0
	v20 = v18 - int32(16)
	m.G0 = v20
	v23 = l0 + int32(-1)
	v24 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23))))
	v26 = v24 & int32(7)
	if v26 == int32(0) {
		goto L8
	} else {
		goto L9
	}
L1:
	;
	F__serverAssert(m, int32(_a604), int32(_a601), int32(489))
	mBase = m.M
	v943 = m.ExcPending
	if v943 != 0 {
		goto L16
	} else {
		goto L284
	}
L2:
	;
	F__serverAssert(m, int32(_a605), int32(_a601), int32(487))
	mBase = m.M
	v937 = m.ExcPending
	if v937 != 0 {
		goto L16
	} else {
		goto L283
	}
L3:
	;
	F__serverAssert(m, int32(_a600), int32(_a601), int32(135))
	mBase = m.M
	v931 = m.ExcPending
	if v931 != 0 {
		goto L16
	} else {
		goto L282
	}
L4:
	;
	F__serverAssert(m, int32(_a607), int32(_a601), int32(448))
	mBase = m.M
	v925 = m.ExcPending
	if v925 != 0 {
		goto L16
	} else {
		goto L281
	}
L5:
	;
	m.G0 = v20 + int32(16)
	return v902
L6:
	;
	v66 = int64(-1)
	if v26 == int32(0) {
		v114 = v24
		v116 = v66
		goto L18
	} else {
		goto L19
	}
L7:
	;
	v34 = base.B2i32(l1 == int32(0))
	if l1 != 0 {
		v65 = v34
		goto L6
	} else {
		goto L11
	}
L8:
	;
	v65 = base.B2i32(l1 == int32(0))
	goto L6
L9:
	;
	if v24&int32(16) != 0 {
		goto L7
	} else {
		goto L10
	}
L10:
	;
	goto L8
L11:
	;
	if v24&int32(32) == int32(0) {
		v65 = v34
		goto L6
	} else {
		goto L12
	}
L12:
	;
	v44 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0+int32(-1)))))
	v46 = v44 & int32(7)
	if base.Ui32(int32(4)) < base.Ui32(v46) {
		v54 = int32(0)
		goto L14
	} else {
		goto L15
	}
L13:
	;
	v58 = *(*int32)(unsafe.Add(mBase, uint32(l0+v54+int32(-4))))
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v58)))
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v58)+4))
	v61 = F_entryUpdateAsStringRef(m, l0, v59, v60, l2)
	mBase = m.M
	v64 = m.ExcPending
	if v64 != 0 {
		goto L16
	} else {
		goto L17
	}
L14:
	;
	goto L13
L15:
	;
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v46<<(uint(int32(2))%32))+uint32(_consts[307])))
	v54 = v53
	goto L14
L16:
	;
	return int32(0)
L17:
	;
	v902 = v61
	goto L5
L18:
	;
	v117 = base.B2i32(l2 == v116)
	if v65&v117 != 0 {
		v902 = l0
		goto L5
	} else {
		goto L36
	}
L19:
	;
	if v24&int32(8) == int32(0) {
		v114 = v24
		v116 = v66
		goto L18
	} else {
		goto L20
	}
L20:
	;
	v78 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0+int32(-1)))))
	v80 = v78 & int32(7)
	if base.Ui32(int32(4)) < base.Ui32(v80) {
		v88 = int32(0)
		goto L22
	} else {
		goto L23
	}
L21:
	;
	v92 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23))))
	if int32(base.Ui32(v92&int32(16))>>(uint(int32(4))%32)) != 0 {
		goto L24
	} else {
		goto L25
	}
L22:
	;
	goto L21
L23:
	;
	v87 = *(*int32)(unsafe.Add(mBase, uint32(v80<<(uint(int32(2))%32))+uint32(_consts[307])))
	v88 = v87
	goto L22
L24:
	;
	v97 = int32(-4)
	goto L26
L25:
	;
	v97 = int32(0)
	goto L26
L26:
	;
	v100 = v92 & int32(7)
	if v100 != 0 {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	v101 = v97
	goto L29
L28:
	;
	v101 = int32(0)
	goto L29
L29:
	;
	if int32(base.Ui32(v92&int32(8))>>(uint(int32(3))%32)) != 0 {
		goto L30
	} else {
		goto L31
	}
L30:
	;
	v109 = int32(-8)
	goto L32
L31:
	;
	v109 = int32(0)
	goto L32
L32:
	;
	if v100 != 0 {
		goto L33
	} else {
		goto L34
	}
L33:
	;
	v111 = v109
	goto L35
L34:
	;
	v111 = int32(0)
	goto L35
L35:
	;
	v113 = *(*int64)(unsafe.Add(mBase, uint32(l0+v88+v101+v111)))
	v114 = v92
	v116 = v113
	goto L18
L36:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+12)) = int32(-1)
	if v65 != 0 {
		goto L38
	} else {
		goto L39
	}
L37:
	;
	switch v261 & int32(7) {
	case 0:
		goto L83
	case 1:
		goto L82
	case 2:
		goto L81
	case 3:
		goto L80
	case 4:
		goto L79
	default:
		v283 = int32(0)
		goto L78
	}
L38:
	;
	v148 = v20 + int32(12)
	v154 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0+int32(-1)))))
	v156 = v154 & int32(7)
	if v156 == int32(0) {
		goto L49
	} else {
		goto L50
	}
L39:
	;
	v124 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1+int32(-1)))))
	switch v124 & int32(7) {
	case 0:
		goto L45
	case 1:
		goto L44
	case 2:
		goto L43
	case 3:
		goto L42
	case 4:
		goto L41
	default:
		v145 = int32(0)
		goto L40
	}
L40:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+12)) = v145
	v260 = l1
	v261 = v114
	goto L37
L41:
	;
	v144 = *(*int32)(unsafe.Add(mBase, uint32(l1+int32(-17))))
	v145 = v144
	goto L40
L42:
	;
	v140 = *(*int32)(unsafe.Add(mBase, uint32(l1+int32(-9))))
	*(*int32)(unsafe.Add(mBase, uint32(v20)+12)) = v140
	v260 = l1
	v261 = v114
	goto L37
L43:
	;
	v136 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1+int32(-5)))))
	*(*int32)(unsafe.Add(mBase, uint32(v20)+12)) = v136
	v260 = l1
	v261 = v114
	goto L37
L44:
	;
	v132 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1+int32(-3)))))
	*(*int32)(unsafe.Add(mBase, uint32(v20)+12)) = v132
	v260 = l1
	v261 = v114
	goto L37
L45:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+12)) = int32(base.Ui32(v124) >> (uint(int32(3)) % 32))
	v260 = l1
	v261 = v114
	goto L37
L46:
	;
	v259 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23))))
	v260 = v258
	v261 = v259
	goto L37
L47:
	;
	v258 = v250
	goto L46
L48:
	;
	v211 = F_sdsAllocPtr(m, l0)
	mBase = m.M
	v213 = v211 + int32(-4)
	if v154&int32(32) == int32(0) {
		goto L65
	} else {
		goto L66
	}
L49:
	;
	switch v156 {
	case 0:
		goto L57
	case 1:
		goto L56
	case 2:
		goto L55
	case 3:
		goto L54
	case 4:
		goto L53
	default:
		v176 = int32(0)
		goto L52
	}
L50:
	;
	if v154&int32(16) != 0 {
		goto L48
	} else {
		goto L51
	}
L51:
	;
	goto L49
L52:
	;
	v178 = int32(1)
	v179 = F_sdsHdrSize(m, v178)
	mBase = m.M
	v180 = l0 + v176 + v179
	v182 = v180 + v178
	if v148 == int32(0) {
		v250 = v182
		goto L47
	} else {
		goto L58
	}
L53:
	;
	v175 = *(*int32)(unsafe.Add(mBase, uint32(l0+int32(-17))))
	v176 = v175
	goto L52
L54:
	;
	v172 = *(*int32)(unsafe.Add(mBase, uint32(l0+int32(-9))))
	v176 = v172
	goto L52
L55:
	;
	v169 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0+int32(-5)))))
	v176 = v169
	goto L52
L56:
	;
	v166 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0+int32(-3)))))
	v176 = v166
	goto L52
L57:
	;
	v176 = int32(base.Ui32(v154) >> (uint(int32(3)) % 32))
	goto L52
L58:
	;
	v185 = int32(0)
	v188 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v180+v185))))
	switch v188 & int32(7) {
	case 0:
		goto L64
	case 1:
		goto L63
	case 2:
		goto L62
	case 3:
		goto L61
	case 4:
		goto L60
	default:
		v209 = v185
		goto L59
	}
L59:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v148))) = v209
	v258 = v182
	goto L46
L60:
	;
	v208 = *(*int32)(unsafe.Add(mBase, uint32(v180+int32(-16))))
	v209 = v208
	goto L59
L61:
	;
	v204 = *(*int32)(unsafe.Add(mBase, uint32(v180+int32(-8))))
	*(*int32)(unsafe.Add(mBase, uint32(v148))) = v204
	v258 = v182
	goto L46
L62:
	;
	v200 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v180+int32(-4)))))
	*(*int32)(unsafe.Add(mBase, uint32(v148))) = v200
	v258 = v182
	goto L46
L63:
	;
	v196 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v180+int32(-2)))))
	*(*int32)(unsafe.Add(mBase, uint32(v148))) = v196
	v258 = v182
	goto L46
L64:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v148))) = int32(base.Ui32(v188) >> (uint(int32(3)) % 32))
	v258 = v182
	goto L46
L65:
	;
	v225 = *(*int32)(unsafe.Add(mBase, uint32(v213)))
	if v148 == int32(0) {
		v250 = v225
		goto L47
	} else {
		goto L71
	}
L66:
	;
	v218 = *(*int32)(unsafe.Add(mBase, uint32(v213)))
	if v218 != 0 {
		goto L67
	} else {
		goto L68
	}
L67:
	;
	if v148 == int32(0) {
		goto L69
	} else {
		goto L70
	}
L68:
	;
	v258 = int32(0)
	goto L46
L69:
	;
	v224 = *(*int32)(unsafe.Add(mBase, uint32(v218)))
	v258 = v224
	goto L46
L70:
	;
	v222 = *(*int32)(unsafe.Add(mBase, uint32(v218)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v148))) = v222
	goto L69
L71:
	;
	v231 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v225+int32(-1)))))
	switch v231 & int32(7) {
	case 0:
		goto L77
	case 1:
		goto L76
	case 2:
		goto L75
	case 3:
		goto L74
	case 4:
		goto L73
	default:
		v248 = int32(0)
		goto L72
	}
L72:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v148))) = v248
	v250 = v225
	goto L47
L73:
	;
	v247 = *(*int32)(unsafe.Add(mBase, uint32(v225+int32(-17))))
	v248 = v247
	goto L72
L74:
	;
	v244 = *(*int32)(unsafe.Add(mBase, uint32(v225+int32(-9))))
	v248 = v244
	goto L72
L75:
	;
	v241 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v225+int32(-5)))))
	v248 = v241
	goto L72
L76:
	;
	v238 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v225+int32(-3)))))
	v248 = v238
	goto L72
L77:
	;
	v248 = int32(base.Ui32(v231) >> (uint(int32(3)) % 32))
	goto L72
L78:
	;
	v284 = *(*int32)(unsafe.Add(mBase, uint32(v20)+12))
	if base.Ui32(int32(32)) <= base.Ui32(v283) {
		goto L85
	} else {
		goto L86
	}
L79:
	;
	v282 = *(*int32)(unsafe.Add(mBase, uint32(l0+int32(-17))))
	v283 = v282
	goto L78
L80:
	;
	v279 = *(*int32)(unsafe.Add(mBase, uint32(l0+int32(-9))))
	v283 = v279
	goto L78
L81:
	;
	v276 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0+int32(-5)))))
	v283 = v276
	goto L78
L82:
	;
	v273 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0+int32(-3)))))
	v283 = v273
	goto L78
L83:
	;
	v283 = int32(base.Ui32(v261&int32(248)) >> (uint(int32(3)) % 32))
	goto L78
L84:
	;
	if v297 != 0 {
		goto L92
	} else {
		goto L93
	}
L85:
	;
	if base.Ui32(int32(253)) <= base.Ui32(v283) {
		goto L87
	} else {
		goto L88
	}
L86:
	;
	v297 = int32(0)
	goto L84
L87:
	;
	if base.Ui32(v283) < base.Ui32(int32(65531)) {
		goto L89
	} else {
		goto L90
	}
L88:
	;
	v297 = int32(1)
	goto L84
L89:
	;
	v296 = int32(2)
	goto L91
L90:
	;
	v296 = int32(3)
	goto L91
L91:
	;
	v297 = v296
	goto L84
L92:
	;
	v299 = v297
	goto L94
L93:
	;
	v299 = int32(1)
	goto L94
L94:
	;
	v301 = base.B2i32(l2 != int64(-1))
	if l2 != int64(-1) {
		goto L95
	} else {
		goto L96
	}
L95:
	;
	v302 = v299
	goto L97
L96:
	;
	v302 = v297
	goto L97
L97:
	;
	v306 = v302 & int32(7)
	if base.Ui32(int32(4)) < base.Ui32(v306) {
		v314 = int32(0)
		goto L99
	} else {
		goto L100
	}
L98:
	;
	v316 = v283 + int32(1)
	v317 = v314 + v316
	v319 = v301 << (uint(int32(3)) % 32)
	if v284 != int32(-1) {
		goto L102
	} else {
		goto L103
	}
L99:
	;
	goto L98
L100:
	;
	v313 = *(*int32)(unsafe.Add(mBase, uint32(v306<<(uint(int32(2))%32))+uint32(_consts[309])))
	v314 = v313
	goto L99
L101:
	;
	v374 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23))))
	if v374&int32(7) == int32(0) {
		goto L114
	} else {
		goto L115
	}
L102:
	;
	goto L106
L103:
	;
	v323 = int32(0)
	v368 = v323
	v369 = v317
	v370 = v317 + v319
	v372 = v302
	v373 = v323
	goto L101
L104:
	;
	v340 = v284 + v336 + int32(1)
	v341 = v317 + v319
	v342 = v340 + v341
	if base.Ui32(v342) < base.Ui32(int32(129)) {
		v368 = int32(1)
		v369 = v317
		v370 = v342
		v372 = v302
		v373 = v340
		goto L101
	} else {
		goto L107
	}
L105:
	;
	goto L104
L106:
	;
	v336 = *(*int32)(unsafe.Add(mBase, _consts[308]))
	goto L105
L107:
	;
	v345 = int32(0)
	if v302 == v345 {
		goto L108
	} else {
		goto L109
	}
L108:
	;
	goto L112
L109:
	;
	v368 = v345
	v369 = v317
	v370 = v341 + int32(4)
	v372 = v302
	v373 = v340
	goto L101
L110:
	;
	v364 = v362 + v316
	v368 = v345
	v369 = v364
	v370 = v364 + v319 + int32(4)
	v372 = int32(1)
	v373 = v340
	goto L101
L111:
	;
	goto L110
L112:
	;
	v362 = *(*int32)(unsafe.Add(mBase, _consts[308]))
	goto L111
L113:
	;
	if l2 == v116 {
		goto L170
	} else {
		goto L171
	}
L114:
	;
	v388 = l0 + int32(-1)
	v389 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v388))))
	v391 = v389 & int32(7)
	if v391 == int32(0) {
		goto L120
	} else {
		goto L121
	}
L115:
	;
	if v374&int32(16) != 0 {
		v516 = int32(0)
		goto L113
	} else {
		goto L116
	}
L116:
	;
	goto L114
L117:
	;
	v516 = v515
	goto L113
L118:
	;
	v458 = v453 & int32(7)
	if v458 == int32(0) {
		goto L147
	} else {
		goto L148
	}
L119:
	;
	v427 = F_sdsAllocPtr(m, l0)
	mBase = m.M
	v430 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v388))))
	if int32(base.Ui32(v430&int32(16))>>(uint(int32(4))%32)) != 0 {
		goto L135
	} else {
		goto L136
	}
L120:
	;
	switch v391 {
	case 0:
		goto L128
	case 1:
		goto L127
	case 2:
		goto L126
	case 3:
		goto L125
	case 4:
		goto L124
	default:
		v411 = int32(0)
		goto L123
	}
L121:
	;
	if v389&int32(16) != 0 {
		goto L119
	} else {
		goto L122
	}
L122:
	;
	goto L120
L123:
	;
	v412 = F_sdsHdrSize(m, v391)
	mBase = m.M
	v413 = v411 + v412
	v417 = v413 + int32(1)
	v418 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v388))))
	if int32(base.Ui32(v418&int32(8))>>(uint(int32(3))%32)) != 0 {
		goto L129
	} else {
		goto L130
	}
L124:
	;
	v410 = *(*int32)(unsafe.Add(mBase, uint32(l0+int32(-17))))
	v411 = v410
	goto L123
L125:
	;
	v407 = *(*int32)(unsafe.Add(mBase, uint32(l0+int32(-9))))
	v411 = v407
	goto L123
L126:
	;
	v404 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0+int32(-5)))))
	v411 = v404
	goto L123
L127:
	;
	v401 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0+int32(-3)))))
	v411 = v401
	goto L123
L128:
	;
	v411 = int32(base.Ui32(v389) >> (uint(int32(3)) % 32))
	goto L123
L129:
	;
	v423 = v413 + int32(9)
	goto L131
L130:
	;
	v423 = v417
	goto L131
L131:
	;
	if v418&int32(7) != 0 {
		goto L132
	} else {
		goto L133
	}
L132:
	;
	v426 = v423
	goto L134
L133:
	;
	v426 = v417
	goto L134
L134:
	;
	v453 = v418
	v455 = v426
	goto L118
L135:
	;
	v435 = int32(-4)
	goto L137
L136:
	;
	v435 = int32(0)
	goto L137
L137:
	;
	v438 = v430 & int32(7)
	if v438 != 0 {
		goto L138
	} else {
		goto L139
	}
L138:
	;
	v439 = v435
	goto L140
L139:
	;
	v439 = int32(0)
	goto L140
L140:
	;
	if int32(base.Ui32(v430&int32(8))>>(uint(int32(3))%32)) != 0 {
		goto L141
	} else {
		goto L142
	}
L141:
	;
	v447 = int32(-8)
	goto L143
L142:
	;
	v447 = int32(0)
	goto L143
L143:
	;
	if v438 != 0 {
		goto L144
	} else {
		goto L145
	}
L144:
	;
	v449 = v447
	goto L146
L145:
	;
	v449 = int32(0)
	goto L146
L146:
	;
	v451 = F_zmalloc_usable_size(m, v427+v439+v449)
	mBase = m.M
	v452 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v388))))
	v453 = v452
	v455 = v451
	goto L118
L147:
	;
	if v458 == int32(0) {
		goto L152
	} else {
		goto L153
	}
L148:
	;
	v461 = int32(48)
	if v453&v461 != v461 {
		goto L147
	} else {
		goto L149
	}
L149:
	;
	v465 = F_sdsAllocPtr(m, l0)
	mBase = m.M
	v468 = *(*int32)(unsafe.Add(mBase, uint32(v465+int32(-4))))
	v469 = F_zmalloc_usable_size(m, v468)
	mBase = m.M
	v515 = v469 + v455
	goto L117
L150:
	;
	v511 = F_sdsAllocSize(m, v509)
	mBase = m.M
	v515 = v511 + v455
	goto L117
L151:
	;
	v499 = F_sdsAllocPtr(m, l0)
	mBase = m.M
	v502 = *(*int32)(unsafe.Add(mBase, uint32(v499+int32(-4))))
	if v453&int32(32) == int32(0) {
		v509 = v502
		goto L150
	} else {
		goto L161
	}
L152:
	;
	switch v458 {
	case 0:
		goto L160
	case 1:
		goto L159
	case 2:
		goto L158
	case 3:
		goto L157
	case 4:
		goto L156
	default:
		v492 = int32(0)
		goto L155
	}
L153:
	;
	if v453&int32(16) != 0 {
		goto L151
	} else {
		goto L154
	}
L154:
	;
	goto L152
L155:
	;
	v494 = int32(1)
	v495 = F_sdsHdrSize(m, v494)
	mBase = m.M
	v509 = l0 + v492 + v495 + v494
	goto L150
L156:
	;
	v491 = *(*int32)(unsafe.Add(mBase, uint32(l0+int32(-17))))
	v492 = v491
	goto L155
L157:
	;
	v488 = *(*int32)(unsafe.Add(mBase, uint32(l0+int32(-9))))
	v492 = v488
	goto L155
L158:
	;
	v485 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0+int32(-5)))))
	v492 = v485
	goto L155
L159:
	;
	v482 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0+int32(-3)))))
	v492 = v482
	goto L155
L160:
	;
	v492 = int32(base.Ui32(v453&int32(248)) >> (uint(int32(3)) % 32))
	goto L155
L161:
	;
	if v502 != 0 {
		goto L162
	} else {
		goto L163
	}
L162:
	;
	v508 = *(*int32)(unsafe.Add(mBase, uint32(v502)))
	v509 = v508
	goto L150
L163:
	;
	v509 = int32(0)
	goto L150
L164:
	;
	v869 = *(*int32)(unsafe.Add(mBase, _consts[113]))
	if v869 == int32(0) {
		goto L272
	} else {
		goto L273
	}
L165:
	;
	v805 = F_entryConstruct(m, v370, l0, v800, int32(0), l2, v368, v372, v319, v373, v369)
	mBase = m.M
	v806 = m.ExcPending
	if v806 != 0 {
		goto L16
	} else {
		goto L251
	}
L166:
	;
	v784 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0+int32(-1)))))
	v786 = v784 & int32(7)
	if base.Ui32(int32(4)) < base.Ui32(v786) {
		v794 = int32(0)
		goto L249
	} else {
		goto L250
	}
L167:
	;
	v620 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23))))
	v622 = v620 & int32(7)
	if v622 == int32(0) {
		goto L207
	} else {
		goto L208
	}
L168:
	;
	v566 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23))))
	if v566&int32(7) == int32(0) {
		goto L4
	} else {
		goto L188
	}
L169:
	;
	v537 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23))))
	v540 = int32(0)
	v546 = base.B2i32(v537&int32(7) == v540) | base.B2i32(v537&int32(16) == v540)
	v548 = v536 | (v368 ^ v546)
	if (v65|v548)&int32(1) != 0 {
		goto L180
	} else {
		goto L181
	}
L170:
	;
	if v65 != 0 {
		v860 = l0
		goto L164
	} else {
		goto L178
	}
L171:
	;
	v517 = int64(-1)
	v521 = base.B2i32(l2 == v517) | base.B2i32(v116 == v517)
	if v65 == int32(0) {
		v536 = v521
		goto L169
	} else {
		goto L172
	}
L172:
	;
	if v521 == int32(0) {
		goto L168
	} else {
		goto L173
	}
L173:
	;
	v526 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23))))
	if v526&int32(7) == int32(0) {
		goto L174
	} else {
		goto L175
	}
L174:
	;
	v533 = F_sdsdup(m, v260)
	mBase = m.M
	v534 = m.ExcPending
	if v534 != 0 {
		goto L16
	} else {
		goto L177
	}
L175:
	;
	if v526&int32(16) != 0 {
		goto L166
	} else {
		goto L176
	}
L176:
	;
	goto L174
L177:
	;
	v800 = v533
	goto L165
L178:
	;
	v536 = int32(0)
	goto L169
L179:
	;
	if l2 == v116 {
		goto L167
	} else {
		goto L187
	}
L180:
	;
	if v548&int32(1) != 0 {
		v800 = v260
		goto L165
	} else {
		goto L186
	}
L181:
	;
	if v546 == int32(0) {
		goto L179
	} else {
		goto L182
	}
L182:
	;
	if base.Ui32(int32(128)) < base.Ui32(v370) {
		v800 = v260
		goto L165
	} else {
		goto L183
	}
L183:
	;
	if base.Ui32(v516) < base.Ui32(v370) {
		v800 = v260
		goto L165
	} else {
		goto L184
	}
L184:
	;
	if base.Ui32(int32(base.Ui32(v516*int32(3))>>(uint(int32(2))%32))) <= base.Ui32(v370) {
		goto L179
	} else {
		goto L185
	}
L185:
	;
	v800 = v260
	goto L165
L186:
	;
	goto L179
L187:
	;
	goto L168
L188:
	;
	if v566&int32(8) == int32(0) {
		goto L4
	} else {
		goto L189
	}
L189:
	;
	v580 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0+int32(-1)))))
	v582 = v580 & int32(7)
	if base.Ui32(int32(4)) < base.Ui32(v582) {
		v590 = int32(0)
		goto L191
	} else {
		goto L192
	}
L190:
	;
	v594 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23))))
	if int32(base.Ui32(v594&int32(16))>>(uint(int32(4))%32)) != 0 {
		goto L193
	} else {
		goto L194
	}
L191:
	;
	goto L190
L192:
	;
	v589 = *(*int32)(unsafe.Add(mBase, uint32(v582<<(uint(int32(2))%32))+uint32(_consts[307])))
	v590 = v589
	goto L191
L193:
	;
	v599 = int32(-4)
	goto L195
L194:
	;
	v599 = int32(0)
	goto L195
L195:
	;
	v602 = v594 & int32(7)
	if v602 != 0 {
		goto L196
	} else {
		goto L197
	}
L196:
	;
	v603 = v599
	goto L198
L197:
	;
	v603 = int32(0)
	goto L198
L198:
	;
	if int32(base.Ui32(v594&int32(8))>>(uint(int32(3))%32)) != 0 {
		goto L199
	} else {
		goto L200
	}
L199:
	;
	v611 = int32(-8)
	goto L201
L200:
	;
	v611 = int32(0)
	goto L201
L201:
	;
	if v602 != 0 {
		goto L202
	} else {
		goto L203
	}
L202:
	;
	v613 = v611
	goto L204
L203:
	;
	v613 = int32(0)
	goto L204
L204:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l0+v590+v603+v613))) = l2
	if v65 != 0 {
		v860 = l0
		goto L164
	} else {
		goto L205
	}
L205:
	;
	goto L167
L206:
	;
	F_entryFreeValuePtr(m, l0)
	mBase = m.M
	v739 = m.ExcPending
	if v739 != 0 {
		goto L16
	} else {
		goto L239
	}
L207:
	;
	v627 = int32(0)
	switch v622 {
	case 0:
		goto L215
	case 1:
		goto L214
	case 2:
		goto L213
	case 3:
		goto L212
	case 4:
		goto L211
	default:
		v643 = v627
		goto L210
	}
L208:
	;
	if v620&int32(16) != 0 {
		goto L206
	} else {
		goto L209
	}
L209:
	;
	goto L207
L210:
	;
	goto L218
L211:
	;
	v642 = *(*int32)(unsafe.Add(mBase, uint32(l0+int32(-17))))
	v643 = v642
	goto L210
L212:
	;
	v639 = *(*int32)(unsafe.Add(mBase, uint32(l0+int32(-9))))
	v643 = v639
	goto L210
L213:
	;
	v636 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0+int32(-5)))))
	v643 = v636
	goto L210
L214:
	;
	v633 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0+int32(-3)))))
	v643 = v633
	goto L210
L215:
	;
	v643 = int32(base.Ui32(v620) >> (uint(int32(3)) % 32))
	goto L210
L216:
	;
	goto L221
L217:
	;
	goto L216
L218:
	;
	v655 = *(*int32)(unsafe.Add(mBase, _consts[308]))
	goto L217
L219:
	;
	v671 = v655 + (l0 + v643)
	v674 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v671))))
	switch v674 & int32(7) {
	case 0:
		goto L227
	case 1:
		goto L226
	case 2:
		goto L225
	case 3:
		goto L224
	case 4:
		goto L223
	default:
		v691 = v627
		goto L222
	}
L220:
	;
	goto L219
L221:
	;
	v668 = *(*int32)(unsafe.Add(mBase, _consts[308]))
	goto L220
L222:
	;
	v692 = int32(0)
	v697 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v671+v692))))
	v699 = v697 & int32(7)
	if base.Ui32(int32(4)) < base.Ui32(v699) {
		v707 = v692
		goto L229
	} else {
		goto L230
	}
L223:
	;
	v690 = *(*int32)(unsafe.Add(mBase, uint32(v671+int32(-8))))
	v691 = v690
	goto L222
L224:
	;
	v687 = *(*int32)(unsafe.Add(mBase, uint32(v671+int32(-4))))
	v691 = v687
	goto L222
L225:
	;
	v684 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v671+int32(-2)))))
	v691 = v684
	goto L222
L226:
	;
	v681 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v671+int32(-1)))))
	v691 = v681
	goto L222
L227:
	;
	v691 = int32(base.Ui32(v674) >> (uint(int32(3)) % 32))
	goto L222
L228:
	;
	v715 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v260+int32(-1)))))
	switch v715 & int32(7) {
	case 0:
		goto L236
	case 1:
		goto L235
	case 2:
		goto L234
	case 3:
		goto L233
	case 4:
		goto L232
	default:
		v732 = int32(0)
		goto L231
	}
L229:
	;
	goto L228
L230:
	;
	v706 = *(*int32)(unsafe.Add(mBase, uint32(v699<<(uint(int32(2))%32))+uint32(_consts[307])))
	v707 = v706
	goto L229
L231:
	;
	v734 = F_sdswrite(m, v671+int32(1)+v707, v668+v691+int32(1), int32(1), v260, v732)
	mBase = m.M
	v735 = m.ExcPending
	if v735 != 0 {
		goto L16
	} else {
		goto L237
	}
L232:
	;
	v731 = *(*int32)(unsafe.Add(mBase, uint32(v260+int32(-17))))
	v732 = v731
	goto L231
L233:
	;
	v728 = *(*int32)(unsafe.Add(mBase, uint32(v260+int32(-9))))
	v732 = v728
	goto L231
L234:
	;
	v725 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v260+int32(-5)))))
	v732 = v725
	goto L231
L235:
	;
	v722 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v260+int32(-3)))))
	v732 = v722
	goto L231
L236:
	;
	v732 = int32(base.Ui32(v715) >> (uint(int32(3)) % 32))
	goto L231
L237:
	;
	F_sdsfree(m, v260)
	mBase = m.M
	v737 = m.ExcPending
	if v737 != 0 {
		goto L16
	} else {
		goto L238
	}
L238:
	;
	v860 = l0
	goto L164
L239:
	;
	v740 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23))))
	v742 = v740 & int32(7)
	if v742 == int32(0) {
		v752 = v740
		goto L240
	} else {
		goto L241
	}
L240:
	;
	if v742 == int32(0) {
		goto L3
	} else {
		goto L243
	}
L241:
	;
	v745 = int32(48)
	if v740&v745 != v745 {
		v752 = v740
		goto L240
	} else {
		goto L242
	}
L242:
	;
	v750 = v740 & int32(223)
	*(*uint8)(unsafe.Add(mBase, uint32(v23))) = uint8(v750)
	v752 = v750
	goto L240
L243:
	;
	if v752&int32(16) == int32(0) {
		goto L3
	} else {
		goto L244
	}
L244:
	;
	v764 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0+int32(-1)))))
	v766 = v764 & int32(7)
	if base.Ui32(int32(4)) < base.Ui32(v766) {
		v774 = int32(0)
		goto L246
	} else {
		goto L247
	}
L245:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0+v774+int32(-4)))) = v260
	v860 = l0
	goto L164
L246:
	;
	goto L245
L247:
	;
	v773 = *(*int32)(unsafe.Add(mBase, uint32(v766<<(uint(int32(2))%32))+uint32(_consts[307])))
	v774 = v773
	goto L246
L248:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0+v794+int32(-4)))) = int32(0)
	v800 = v260
	goto L165
L249:
	;
	goto L248
L250:
	;
	v793 = *(*int32)(unsafe.Add(mBase, uint32(v786<<(uint(int32(2))%32))+uint32(_consts[307])))
	v794 = v793
	goto L249
L251:
	;
	v807 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23))))
	if v807&int32(7) == int32(0) {
		goto L252
	} else {
		goto L253
	}
L252:
	;
	v823 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0+int32(-1)))))
	v825 = v823 & int32(7)
	if base.Ui32(int32(4)) < base.Ui32(v825) {
		v833 = int32(0)
		goto L257
	} else {
		goto L258
	}
L253:
	;
	if v807&int32(16) == int32(0) {
		goto L252
	} else {
		goto L254
	}
L254:
	;
	F_entryFreeValuePtr(m, l0)
	mBase = m.M
	v817 = m.ExcPending
	if v817 != 0 {
		goto L16
	} else {
		goto L255
	}
L255:
	;
	goto L252
L256:
	;
	v837 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23))))
	if int32(base.Ui32(v837&int32(16))>>(uint(int32(4))%32)) != 0 {
		goto L259
	} else {
		goto L260
	}
L257:
	;
	goto L256
L258:
	;
	v832 = *(*int32)(unsafe.Add(mBase, uint32(v825<<(uint(int32(2))%32))+uint32(_consts[307])))
	v833 = v832
	goto L257
L259:
	;
	v842 = int32(-4)
	goto L261
L260:
	;
	v842 = int32(0)
	goto L261
L261:
	;
	v845 = v837 & int32(7)
	if v845 != 0 {
		goto L262
	} else {
		goto L263
	}
L262:
	;
	v846 = v842
	goto L264
L263:
	;
	v846 = int32(0)
	goto L264
L264:
	;
	if int32(base.Ui32(v837&int32(8))>>(uint(int32(3))%32)) != 0 {
		goto L265
	} else {
		goto L266
	}
L265:
	;
	v854 = int32(-8)
	goto L267
L266:
	;
	v854 = int32(0)
	goto L267
L267:
	;
	if v845 != 0 {
		goto L268
	} else {
		goto L269
	}
L268:
	;
	v856 = v854
	goto L270
L269:
	;
	v856 = int32(0)
	goto L270
L270:
	;
	F_valkey_free(m, l0+v833+v846+v856)
	mBase = m.M
	v859 = m.ExcPending
	if v859 != 0 {
		goto L16
	} else {
		goto L271
	}
L271:
	;
	v860 = v805
	goto L164
L272:
	;
	if v860 == int32(0) {
		goto L1
	} else {
		goto L280
	}
L273:
	;
	v874 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v860+int32(-1)))))
	v876 = v874 & int32(7)
	if v368 == base.B2i32(v876 != int32(0))&int32(base.Ui32(v874&int32(16))>>(uint(int32(4))%32)) {
		goto L2
	} else {
		goto L274
	}
L274:
	;
	if v876 != 0 {
		goto L275
	} else {
		goto L276
	}
L275:
	;
	v890 = int32(base.Ui32(v874)>>(uint(int32(3))%32)) & int32(1)
	goto L277
L276:
	;
	v890 = int32(0)
	goto L277
L277:
	;
	if v890 == base.B2i32(l2 != int64(-1)) {
		v902 = v860
		goto L5
	} else {
		goto L278
	}
L278:
	;
	F__serverAssert(m, int32(_a606), int32(_a601), int32(488))
	mBase = m.M
	v898 = m.ExcPending
	if v898 != 0 {
		goto L16
	} else {
		goto L279
	}
L279:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L280:
	;
	v902 = v860
	goto L5
L281:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L282:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L283:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L284:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_entryUpdateAsStringRef(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int64) int32 {
	mBase := m.M
	_ = mBase
	var v9 int64
	_ = v9
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v40 int32
	_ = v40
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v61 int64
	_ = v61
	var v62 int64
	_ = v62
	var v63 int32
	_ = v63
	var v67 int64
	_ = v67
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v87 int32
	_ = v87
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v110 int32
	_ = v110
	var v114 int32
	_ = v114
	var v117 int32
	_ = v117
	var v120 int32
	_ = v120
	var v134 int32
	_ = v134
	var v136 int32
	_ = v136
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v150 int32
	_ = v150
	var v152 int32
	_ = v152
	var v158 int32
	_ = v158
	var v164 int32
	_ = v164
	var v178 int32
	_ = v178
	var v180 int32
	_ = v180
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v192 int32
	_ = v192
	var v197 int32
	_ = v197
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v209 int32
	_ = v209
	var v211 int32
	_ = v211
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v222 int32
	_ = v222
	var v229 int32
	_ = v229
	var v232 int32
	_ = v232
	var v235 int32
	_ = v235
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v251 int32
	_ = v251
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
	var v267 int32
	_ = v267
	var v268 int32
	_ = v268
	var v269 int32
	_ = v269
	var v279 int32
	_ = v279
	var v285 int32
	_ = v285
	var v287 int32
	_ = v287
	var v294 int32
	_ = v294
	var v295 int32
	_ = v295
	var v299 int32
	_ = v299
	var v304 int32
	_ = v304
	var v307 int32
	_ = v307
	var v308 int32
	_ = v308
	var v316 int32
	_ = v316
	var v318 int32
	_ = v318
	var v321 int32
	_ = v321
	var v323 int32
	_ = v323
	var v324 int32
	_ = v324
	var v330 int32
	_ = v330
	var v332 int32
	_ = v332
	var v342 int32
	_ = v342
	var v348 int32
	_ = v348
	var v354 int32
	_ = v354
	v9 = int64(-1)
	v11 = l0 + int32(-1)
	v12 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11))))
	if v12&int32(7) == int32(0) {
		v62 = v9
		v63 = v12
	} else {
		if v12&int32(8) == int32(0) {
			v62 = v9
			v63 = v12
		} else {
			v26 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0+int32(-1)))))
			v28 = v26 & int32(7)
			if base.Ui32(int32(4)) < base.Ui32(v28) {
				v36 = int32(0)
			} else {
				v35 = *(*int32)(unsafe.Add(mBase, uint32(v28<<(uint(int32(2))%32))+uint32(_consts[307])))
				v36 = v35
			}
			v40 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11))))
			if int32(base.Ui32(v40&int32(16))>>(uint(int32(4))%32)) != 0 {
				v45 = int32(-4)
			} else {
				v45 = int32(0)
			}
			v48 = v40 & int32(7)
			if v48 != 0 {
				v49 = v45
			} else {
				v49 = int32(0)
			}
			if int32(base.Ui32(v40&int32(8))>>(uint(int32(3))%32)) != 0 {
				v57 = int32(-8)
			} else {
				v57 = int32(0)
			}
			if v48 != 0 {
				v59 = v57
			} else {
				v59 = int32(0)
			}
			v61 = *(*int64)(unsafe.Add(mBase, uint32(l0+v36+v49+v59)))
			v62 = v61
			v63 = v40
		}
	}
	if l3 == v62 {
		v72 = int32(0)
	} else {
		v67 = int64(-1)
		v72 = base.B2i32(l3 == v67) | base.B2i32(v62 == v67)
	}
	if v72 != 0 {
		v217 = F_valkey_malloc(m, int32(8))
		mBase = m.M
		v218 = m.ExcPending
		if v218 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v217)+4)) = l2
			*(*int32)(unsafe.Add(mBase, uint32(v217))) = l1
			v222 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11))))
			switch v222 & int32(7) {
			case 0:
				v239 = int32(base.Ui32(v222) >> (uint(int32(3)) % 32))
			case 1:
				v229 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0+int32(-3)))))
				v239 = v229
			case 2:
				v232 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0+int32(-5)))))
				v239 = v232
			case 3:
				v235 = *(*int32)(unsafe.Add(mBase, uint32(l0+int32(-9))))
				v239 = v235
			case 4:
				v238 = *(*int32)(unsafe.Add(mBase, uint32(l0+int32(-17))))
				v239 = v238
			default:
				v239 = int32(0)
			}
			v251 = *(*int32)(unsafe.Add(mBase, _consts[308]))
			v254 = int32(1)
			v255 = v239 + v251 + v254
			v259 = base.B2i32(l3 != int64(-1)) << (uint(int32(3)) % 32)
			v261 = int32(4)
			v263 = int32(0)
			v267 = F_entryConstruct(m, v255+v259+v261, l0, v263, v217, l3, v263, v254, v259, v261, v255)
			mBase = m.M
			v268 = m.ExcPending
			if v268 != 0 {
				return int32(0)
			} else {
				v269 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11))))
				if v269&int32(7) == int32(0) {
					v285 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0+int32(-1)))))
					v287 = v285 & int32(7)
					if base.Ui32(int32(4)) < base.Ui32(v287) {
						v295 = int32(0)
					} else {
						v294 = *(*int32)(unsafe.Add(mBase, uint32(v287<<(uint(int32(2))%32))+uint32(_consts[307])))
						v295 = v294
					}
					v299 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11))))
					if int32(base.Ui32(v299&int32(16))>>(uint(int32(4))%32)) != 0 {
						v304 = int32(-4)
					} else {
						v304 = int32(0)
					}
					v307 = v299 & int32(7)
					if v307 != 0 {
						v308 = v304
					} else {
						v308 = int32(0)
					}
					if int32(base.Ui32(v299&int32(8))>>(uint(int32(3))%32)) != 0 {
						v316 = int32(-8)
					} else {
						v316 = int32(0)
					}
					if v307 != 0 {
						v318 = v316
					} else {
						v318 = int32(0)
					}
					F_valkey_free(m, l0+v295+v308+v318)
					mBase = m.M
					v321 = m.ExcPending
					if v321 != 0 {
						return int32(0)
					} else {
						v323 = v267 + int32(-1)
						v324 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v323))))
						if v324&int32(7) == int32(0) {
						} else {
							v330 = v324 | int32(32)
							*(*uint8)(unsafe.Add(mBase, uint32(v323))) = uint8(v330)
						}
						v332 = v267
						return v332
					}
				} else {
					if v269&int32(16) == int32(0) {
						v285 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0+int32(-1)))))
						v287 = v285 & int32(7)
						if base.Ui32(int32(4)) < base.Ui32(v287) {
							v295 = int32(0)
						} else {
							v294 = *(*int32)(unsafe.Add(mBase, uint32(v287<<(uint(int32(2))%32))+uint32(_consts[307])))
							v295 = v294
						}
						v299 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11))))
						if int32(base.Ui32(v299&int32(16))>>(uint(int32(4))%32)) != 0 {
							v304 = int32(-4)
						} else {
							v304 = int32(0)
						}
						v307 = v299 & int32(7)
						if v307 != 0 {
							v308 = v304
						} else {
							v308 = int32(0)
						}
						if int32(base.Ui32(v299&int32(8))>>(uint(int32(3))%32)) != 0 {
							v316 = int32(-8)
						} else {
							v316 = int32(0)
						}
						if v307 != 0 {
							v318 = v316
						} else {
							v318 = int32(0)
						}
						F_valkey_free(m, l0+v295+v308+v318)
						mBase = m.M
						v321 = m.ExcPending
						if v321 != 0 {
							return int32(0)
						} else {
							v323 = v267 + int32(-1)
							v324 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v323))))
							if v324&int32(7) == int32(0) {
							} else {
								v330 = v324 | int32(32)
								*(*uint8)(unsafe.Add(mBase, uint32(v323))) = uint8(v330)
							}
							v332 = v267
							return v332
						}
					} else {
						F_entryFreeValuePtr(m, l0)
						mBase = m.M
						v279 = m.ExcPending
						if v279 != 0 {
							return int32(0)
						} else {
							v285 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0+int32(-1)))))
							v287 = v285 & int32(7)
							if base.Ui32(int32(4)) < base.Ui32(v287) {
								v295 = int32(0)
							} else {
								v294 = *(*int32)(unsafe.Add(mBase, uint32(v287<<(uint(int32(2))%32))+uint32(_consts[307])))
								v295 = v294
							}
							v299 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11))))
							if int32(base.Ui32(v299&int32(16))>>(uint(int32(4))%32)) != 0 {
								v304 = int32(-4)
							} else {
								v304 = int32(0)
							}
							v307 = v299 & int32(7)
							if v307 != 0 {
								v308 = v304
							} else {
								v308 = int32(0)
							}
							if int32(base.Ui32(v299&int32(8))>>(uint(int32(3))%32)) != 0 {
								v316 = int32(-8)
							} else {
								v316 = int32(0)
							}
							if v307 != 0 {
								v318 = v316
							} else {
								v318 = int32(0)
							}
							F_valkey_free(m, l0+v295+v308+v318)
							mBase = m.M
							v321 = m.ExcPending
							if v321 != 0 {
								return int32(0)
							} else {
								v323 = v267 + int32(-1)
								v324 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v323))))
								if v324&int32(7) == int32(0) {
								} else {
									v330 = v324 | int32(32)
									*(*uint8)(unsafe.Add(mBase, uint32(v323))) = uint8(v330)
								}
								v332 = v267
								return v332
							}
						}
					}
				}
			}
		}
	} else {
		v74 = v63 & int32(7)
		if v74 == int32(0) {
			v217 = F_valkey_malloc(m, int32(8))
			mBase = m.M
			v218 = m.ExcPending
			if v218 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v217)+4)) = l2
				*(*int32)(unsafe.Add(mBase, uint32(v217))) = l1
				v222 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11))))
				switch v222 & int32(7) {
				case 0:
					v239 = int32(base.Ui32(v222) >> (uint(int32(3)) % 32))
				case 1:
					v229 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0+int32(-3)))))
					v239 = v229
				case 2:
					v232 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0+int32(-5)))))
					v239 = v232
				case 3:
					v235 = *(*int32)(unsafe.Add(mBase, uint32(l0+int32(-9))))
					v239 = v235
				case 4:
					v238 = *(*int32)(unsafe.Add(mBase, uint32(l0+int32(-17))))
					v239 = v238
				default:
					v239 = int32(0)
				}
				v251 = *(*int32)(unsafe.Add(mBase, _consts[308]))
				v254 = int32(1)
				v255 = v239 + v251 + v254
				v259 = base.B2i32(l3 != int64(-1)) << (uint(int32(3)) % 32)
				v261 = int32(4)
				v263 = int32(0)
				v267 = F_entryConstruct(m, v255+v259+v261, l0, v263, v217, l3, v263, v254, v259, v261, v255)
				mBase = m.M
				v268 = m.ExcPending
				if v268 != 0 {
					return int32(0)
				} else {
					v269 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11))))
					if v269&int32(7) == int32(0) {
						v285 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0+int32(-1)))))
						v287 = v285 & int32(7)
						if base.Ui32(int32(4)) < base.Ui32(v287) {
							v295 = int32(0)
						} else {
							v294 = *(*int32)(unsafe.Add(mBase, uint32(v287<<(uint(int32(2))%32))+uint32(_consts[307])))
							v295 = v294
						}
						v299 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11))))
						if int32(base.Ui32(v299&int32(16))>>(uint(int32(4))%32)) != 0 {
							v304 = int32(-4)
						} else {
							v304 = int32(0)
						}
						v307 = v299 & int32(7)
						if v307 != 0 {
							v308 = v304
						} else {
							v308 = int32(0)
						}
						if int32(base.Ui32(v299&int32(8))>>(uint(int32(3))%32)) != 0 {
							v316 = int32(-8)
						} else {
							v316 = int32(0)
						}
						if v307 != 0 {
							v318 = v316
						} else {
							v318 = int32(0)
						}
						F_valkey_free(m, l0+v295+v308+v318)
						mBase = m.M
						v321 = m.ExcPending
						if v321 != 0 {
							return int32(0)
						} else {
							v323 = v267 + int32(-1)
							v324 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v323))))
							if v324&int32(7) == int32(0) {
							} else {
								v330 = v324 | int32(32)
								*(*uint8)(unsafe.Add(mBase, uint32(v323))) = uint8(v330)
							}
							v332 = v267
							return v332
						}
					} else {
						if v269&int32(16) == int32(0) {
							v285 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0+int32(-1)))))
							v287 = v285 & int32(7)
							if base.Ui32(int32(4)) < base.Ui32(v287) {
								v295 = int32(0)
							} else {
								v294 = *(*int32)(unsafe.Add(mBase, uint32(v287<<(uint(int32(2))%32))+uint32(_consts[307])))
								v295 = v294
							}
							v299 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11))))
							if int32(base.Ui32(v299&int32(16))>>(uint(int32(4))%32)) != 0 {
								v304 = int32(-4)
							} else {
								v304 = int32(0)
							}
							v307 = v299 & int32(7)
							if v307 != 0 {
								v308 = v304
							} else {
								v308 = int32(0)
							}
							if int32(base.Ui32(v299&int32(8))>>(uint(int32(3))%32)) != 0 {
								v316 = int32(-8)
							} else {
								v316 = int32(0)
							}
							if v307 != 0 {
								v318 = v316
							} else {
								v318 = int32(0)
							}
							F_valkey_free(m, l0+v295+v308+v318)
							mBase = m.M
							v321 = m.ExcPending
							if v321 != 0 {
								return int32(0)
							} else {
								v323 = v267 + int32(-1)
								v324 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v323))))
								if v324&int32(7) == int32(0) {
								} else {
									v330 = v324 | int32(32)
									*(*uint8)(unsafe.Add(mBase, uint32(v323))) = uint8(v330)
								}
								v332 = v267
								return v332
							}
						} else {
							F_entryFreeValuePtr(m, l0)
							mBase = m.M
							v279 = m.ExcPending
							if v279 != 0 {
								return int32(0)
							} else {
								v285 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0+int32(-1)))))
								v287 = v285 & int32(7)
								if base.Ui32(int32(4)) < base.Ui32(v287) {
									v295 = int32(0)
								} else {
									v294 = *(*int32)(unsafe.Add(mBase, uint32(v287<<(uint(int32(2))%32))+uint32(_consts[307])))
									v295 = v294
								}
								v299 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11))))
								if int32(base.Ui32(v299&int32(16))>>(uint(int32(4))%32)) != 0 {
									v304 = int32(-4)
								} else {
									v304 = int32(0)
								}
								v307 = v299 & int32(7)
								if v307 != 0 {
									v308 = v304
								} else {
									v308 = int32(0)
								}
								if int32(base.Ui32(v299&int32(8))>>(uint(int32(3))%32)) != 0 {
									v316 = int32(-8)
								} else {
									v316 = int32(0)
								}
								if v307 != 0 {
									v318 = v316
								} else {
									v318 = int32(0)
								}
								F_valkey_free(m, l0+v295+v308+v318)
								mBase = m.M
								v321 = m.ExcPending
								if v321 != 0 {
									return int32(0)
								} else {
									v323 = v267 + int32(-1)
									v324 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v323))))
									if v324&int32(7) == int32(0) {
									} else {
										v330 = v324 | int32(32)
										*(*uint8)(unsafe.Add(mBase, uint32(v323))) = uint8(v330)
									}
									v332 = v267
									return v332
								}
							}
						}
					}
				}
			}
		} else {
			if v63&int32(16) == int32(0) {
				v217 = F_valkey_malloc(m, int32(8))
				mBase = m.M
				v218 = m.ExcPending
				if v218 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v217)+4)) = l2
					*(*int32)(unsafe.Add(mBase, uint32(v217))) = l1
					v222 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11))))
					switch v222 & int32(7) {
					case 0:
						v239 = int32(base.Ui32(v222) >> (uint(int32(3)) % 32))
					case 1:
						v229 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0+int32(-3)))))
						v239 = v229
					case 2:
						v232 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0+int32(-5)))))
						v239 = v232
					case 3:
						v235 = *(*int32)(unsafe.Add(mBase, uint32(l0+int32(-9))))
						v239 = v235
					case 4:
						v238 = *(*int32)(unsafe.Add(mBase, uint32(l0+int32(-17))))
						v239 = v238
					default:
						v239 = int32(0)
					}
					v251 = *(*int32)(unsafe.Add(mBase, _consts[308]))
					v254 = int32(1)
					v255 = v239 + v251 + v254
					v259 = base.B2i32(l3 != int64(-1)) << (uint(int32(3)) % 32)
					v261 = int32(4)
					v263 = int32(0)
					v267 = F_entryConstruct(m, v255+v259+v261, l0, v263, v217, l3, v263, v254, v259, v261, v255)
					mBase = m.M
					v268 = m.ExcPending
					if v268 != 0 {
						return int32(0)
					} else {
						v269 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11))))
						if v269&int32(7) == int32(0) {
							v285 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0+int32(-1)))))
							v287 = v285 & int32(7)
							if base.Ui32(int32(4)) < base.Ui32(v287) {
								v295 = int32(0)
							} else {
								v294 = *(*int32)(unsafe.Add(mBase, uint32(v287<<(uint(int32(2))%32))+uint32(_consts[307])))
								v295 = v294
							}
							v299 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11))))
							if int32(base.Ui32(v299&int32(16))>>(uint(int32(4))%32)) != 0 {
								v304 = int32(-4)
							} else {
								v304 = int32(0)
							}
							v307 = v299 & int32(7)
							if v307 != 0 {
								v308 = v304
							} else {
								v308 = int32(0)
							}
							if int32(base.Ui32(v299&int32(8))>>(uint(int32(3))%32)) != 0 {
								v316 = int32(-8)
							} else {
								v316 = int32(0)
							}
							if v307 != 0 {
								v318 = v316
							} else {
								v318 = int32(0)
							}
							F_valkey_free(m, l0+v295+v308+v318)
							mBase = m.M
							v321 = m.ExcPending
							if v321 != 0 {
								return int32(0)
							} else {
								v323 = v267 + int32(-1)
								v324 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v323))))
								if v324&int32(7) == int32(0) {
								} else {
									v330 = v324 | int32(32)
									*(*uint8)(unsafe.Add(mBase, uint32(v323))) = uint8(v330)
								}
								v332 = v267
								return v332
							}
						} else {
							if v269&int32(16) == int32(0) {
								v285 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0+int32(-1)))))
								v287 = v285 & int32(7)
								if base.Ui32(int32(4)) < base.Ui32(v287) {
									v295 = int32(0)
								} else {
									v294 = *(*int32)(unsafe.Add(mBase, uint32(v287<<(uint(int32(2))%32))+uint32(_consts[307])))
									v295 = v294
								}
								v299 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11))))
								if int32(base.Ui32(v299&int32(16))>>(uint(int32(4))%32)) != 0 {
									v304 = int32(-4)
								} else {
									v304 = int32(0)
								}
								v307 = v299 & int32(7)
								if v307 != 0 {
									v308 = v304
								} else {
									v308 = int32(0)
								}
								if int32(base.Ui32(v299&int32(8))>>(uint(int32(3))%32)) != 0 {
									v316 = int32(-8)
								} else {
									v316 = int32(0)
								}
								if v307 != 0 {
									v318 = v316
								} else {
									v318 = int32(0)
								}
								F_valkey_free(m, l0+v295+v308+v318)
								mBase = m.M
								v321 = m.ExcPending
								if v321 != 0 {
									return int32(0)
								} else {
									v323 = v267 + int32(-1)
									v324 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v323))))
									if v324&int32(7) == int32(0) {
									} else {
										v330 = v324 | int32(32)
										*(*uint8)(unsafe.Add(mBase, uint32(v323))) = uint8(v330)
									}
									v332 = v267
									return v332
								}
							} else {
								F_entryFreeValuePtr(m, l0)
								mBase = m.M
								v279 = m.ExcPending
								if v279 != 0 {
									return int32(0)
								} else {
									v285 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0+int32(-1)))))
									v287 = v285 & int32(7)
									if base.Ui32(int32(4)) < base.Ui32(v287) {
										v295 = int32(0)
									} else {
										v294 = *(*int32)(unsafe.Add(mBase, uint32(v287<<(uint(int32(2))%32))+uint32(_consts[307])))
										v295 = v294
									}
									v299 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11))))
									if int32(base.Ui32(v299&int32(16))>>(uint(int32(4))%32)) != 0 {
										v304 = int32(-4)
									} else {
										v304 = int32(0)
									}
									v307 = v299 & int32(7)
									if v307 != 0 {
										v308 = v304
									} else {
										v308 = int32(0)
									}
									if int32(base.Ui32(v299&int32(8))>>(uint(int32(3))%32)) != 0 {
										v316 = int32(-8)
									} else {
										v316 = int32(0)
									}
									if v307 != 0 {
										v318 = v316
									} else {
										v318 = int32(0)
									}
									F_valkey_free(m, l0+v295+v308+v318)
									mBase = m.M
									v321 = m.ExcPending
									if v321 != 0 {
										return int32(0)
									} else {
										v323 = v267 + int32(-1)
										v324 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v323))))
										if v324&int32(7) == int32(0) {
										} else {
											v330 = v324 | int32(32)
											*(*uint8)(unsafe.Add(mBase, uint32(v323))) = uint8(v330)
										}
										v332 = v267
										return v332
									}
								}
							}
						}
					}
				}
			} else {
				if v63&int32(32) == int32(0) {
					v114 = F_valkey_malloc(m, int32(8))
					mBase = m.M
					v117 = m.ExcPending
					if v117 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v114)+4)) = l2
						*(*int32)(unsafe.Add(mBase, uint32(v114))) = l1
						v120 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11))))
						if v120&int32(7) == int32(0) {
							F__serverAssert(m, int32(_a600), int32(_a601), int32(135))
							mBase = m.M
							v348 = m.ExcPending
							if v348 != 0 {
								return int32(0)
							} else {
								F_abort(m)
								mBase = m.M
								base.Wasm_trap_unreachable()
								for {
								}
							}
						} else {
							if v120&int32(16) == int32(0) {
								F__serverAssert(m, int32(_a600), int32(_a601), int32(135))
								mBase = m.M
								v348 = m.ExcPending
								if v348 != 0 {
									return int32(0)
								} else {
									F_abort(m)
									mBase = m.M
									base.Wasm_trap_unreachable()
									for {
									}
								}
							} else {
								v134 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0+int32(-1)))))
								v136 = v134 & int32(7)
								if base.Ui32(int32(4)) < base.Ui32(v136) {
									v144 = int32(0)
								} else {
									v143 = *(*int32)(unsafe.Add(mBase, uint32(v136<<(uint(int32(2))%32))+uint32(_consts[307])))
									v144 = v143
								}
								v147 = l0 + v144 + int32(-4)
								v148 = *(*int32)(unsafe.Add(mBase, uint32(v147)))
								F_sdsfree(m, v148)
								mBase = m.M
								v150 = m.ExcPending
								if v150 != 0 {
									return int32(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v147))) = v114
									v152 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11))))
									if v152&int32(7) == int32(0) {
									} else {
										v158 = v152 | int32(32)
										*(*uint8)(unsafe.Add(mBase, uint32(v11))) = uint8(v158)
									}
									if l3 == int64(-1) {
										v332 = l0
										return v332
									} else {
										v164 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11))))
										if v164&int32(7) == int32(0) {
											F__serverAssert(m, int32(_a602), int32(_a601), int32(204))
											mBase = m.M
											v354 = m.ExcPending
											if v354 != 0 {
												return int32(0)
											} else {
												F_abort(m)
												mBase = m.M
												base.Wasm_trap_unreachable()
												for {
												}
											}
										} else {
											if v164&int32(8) == int32(0) {
												F__serverAssert(m, int32(_a602), int32(_a601), int32(204))
												mBase = m.M
												v354 = m.ExcPending
												if v354 != 0 {
													return int32(0)
												} else {
													F_abort(m)
													mBase = m.M
													base.Wasm_trap_unreachable()
													for {
													}
												}
											} else {
												v178 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0+int32(-1)))))
												v180 = v178 & int32(7)
												if base.Ui32(int32(4)) < base.Ui32(v180) {
													v188 = int32(0)
												} else {
													v187 = *(*int32)(unsafe.Add(mBase, uint32(v180<<(uint(int32(2))%32))+uint32(_consts[307])))
													v188 = v187
												}
												v192 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11))))
												if int32(base.Ui32(v192&int32(16))>>(uint(int32(4))%32)) != 0 {
													v197 = int32(-4)
												} else {
													v197 = int32(0)
												}
												v200 = v192 & int32(7)
												if v200 != 0 {
													v201 = v197
												} else {
													v201 = int32(0)
												}
												if int32(base.Ui32(v192&int32(8))>>(uint(int32(3))%32)) != 0 {
													v209 = int32(-8)
												} else {
													v209 = int32(0)
												}
												if v200 != 0 {
													v211 = v209
												} else {
													v211 = int32(0)
												}
												*(*int64)(unsafe.Add(mBase, uint32(l0+v188+v201+v211))) = l3
												return l0
											}
										}
									}
								}
							}
						}
					}
				} else {
					if v74 == int32(0) {
						F__serverAssert(m, int32(_a603), int32(_a601), int32(146))
						mBase = m.M
						v342 = m.ExcPending
						if v342 != 0 {
							return int32(0)
						} else {
							F_abort(m)
							mBase = m.M
							base.Wasm_trap_unreachable()
							for {
							}
						}
					} else {
						v87 = int32(48)
						if v63&v87 != v87 {
							F__serverAssert(m, int32(_a603), int32(_a601), int32(146))
							mBase = m.M
							v342 = m.ExcPending
							if v342 != 0 {
								return int32(0)
							} else {
								F_abort(m)
								mBase = m.M
								base.Wasm_trap_unreachable()
								for {
								}
							}
						} else {
							v96 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0+int32(-1)))))
							v98 = v96 & int32(7)
							if base.Ui32(int32(4)) < base.Ui32(v98) {
								v106 = int32(0)
							} else {
								v105 = *(*int32)(unsafe.Add(mBase, uint32(v98<<(uint(int32(2))%32))+uint32(_consts[307])))
								v106 = v105
							}
							v110 = *(*int32)(unsafe.Add(mBase, uint32(l0+v106+int32(-4))))
							*(*int32)(unsafe.Add(mBase, uint32(v110)+4)) = l2
							*(*int32)(unsafe.Add(mBase, uint32(v110))) = l1
							if l3 == int64(-1) {
								v332 = l0
								return v332
							} else {
								v164 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11))))
								if v164&int32(7) == int32(0) {
									F__serverAssert(m, int32(_a602), int32(_a601), int32(204))
									mBase = m.M
									v354 = m.ExcPending
									if v354 != 0 {
										return int32(0)
									} else {
										F_abort(m)
										mBase = m.M
										base.Wasm_trap_unreachable()
										for {
										}
									}
								} else {
									if v164&int32(8) == int32(0) {
										F__serverAssert(m, int32(_a602), int32(_a601), int32(204))
										mBase = m.M
										v354 = m.ExcPending
										if v354 != 0 {
											return int32(0)
										} else {
											F_abort(m)
											mBase = m.M
											base.Wasm_trap_unreachable()
											for {
											}
										}
									} else {
										v178 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0+int32(-1)))))
										v180 = v178 & int32(7)
										if base.Ui32(int32(4)) < base.Ui32(v180) {
											v188 = int32(0)
										} else {
											v187 = *(*int32)(unsafe.Add(mBase, uint32(v180<<(uint(int32(2))%32))+uint32(_consts[307])))
											v188 = v187
										}
										v192 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11))))
										if int32(base.Ui32(v192&int32(16))>>(uint(int32(4))%32)) != 0 {
											v197 = int32(-4)
										} else {
											v197 = int32(0)
										}
										v200 = v192 & int32(7)
										if v200 != 0 {
											v201 = v197
										} else {
											v201 = int32(0)
										}
										if int32(base.Ui32(v192&int32(8))>>(uint(int32(3))%32)) != 0 {
											v209 = int32(-8)
										} else {
											v209 = int32(0)
										}
										if v200 != 0 {
											v211 = v209
										} else {
											v211 = int32(0)
										}
										*(*int64)(unsafe.Add(mBase, uint32(l0+v188+v201+v211))) = l3
										return l0
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
