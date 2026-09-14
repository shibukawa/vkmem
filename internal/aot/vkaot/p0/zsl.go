package p0

import (
	base "github.com/shibukawa/vkmem/internal/aot/vkaot/base"
	"unsafe"
)

func F_zslCreateNode(m *base.Module, l0 int32, l1 float64, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v12 int32
	_ = v12
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v72 int32
	_ = v72
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	v12 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2+int32(-1)))))
	switch v12 & int32(7) {
	case 0:
		v29 = int32(base.Ui32(v12) >> (uint(int32(3)) % 32))
	case 1:
		v19 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2+int32(-3)))))
		v29 = v19
	case 2:
		v22 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l2+int32(-5)))))
		v29 = v22
	case 3:
		v25 = *(*int32)(unsafe.Add(mBase, uint32(l2+int32(-9))))
		v29 = v25
	case 4:
		v28 = *(*int32)(unsafe.Add(mBase, uint32(l2+int32(-17))))
		v29 = v28
	default:
		v29 = int32(0)
	}
	v31 = l0 << (uint(int32(3)) % 32)
	if base.Ui32(int32(32)) <= base.Ui32(v29) {
		if base.Ui32(int32(253)) <= base.Ui32(v29) {
			if base.Ui32(v29) < base.Ui32(int32(65531)) {
				v42 = int32(2)
			} else {
				v42 = int32(3)
			}
			v43 = v42
		} else {
			v43 = int32(1)
		}
	} else {
		v43 = int32(0)
	}
	v47 = v43 & int32(7)
	if base.Ui32(int32(4)) < base.Ui32(v47) {
		v55 = int32(0)
	} else {
		v54 = *(*int32)(unsafe.Add(mBase, uint32(v47<<(uint(int32(2))%32))+uint32(_consts[251])))
		v55 = v54
	}
	v58 = v29 + v55 + int32(1)
	v62 = F_valkey_malloc(m, v31+v58+int32(17))
	mBase = m.M
	v65 = m.ExcPending
	if v65 != 0 {
		return int32(0)
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(v62)+16)) = l0
		*(*float64)(unsafe.Add(mBase, uint32(v62))) = l1
		v68 = v62 + v31
		v72 = v43 & int32(7)
		if base.Ui32(int32(4)) < base.Ui32(v72) {
			v80 = int32(0)
		} else {
			v79 = *(*int32)(unsafe.Add(mBase, uint32(v72<<(uint(int32(2))%32))+uint32(_consts[251])))
			v80 = v79
		}
		*(*uint8)(unsafe.Add(mBase, uint32(v68)+16)) = uint8(v80)
		v84 = F_sdswrite(m, v68+int32(17), v58, v43, l2, v29)
		mBase = m.M
		v85 = m.ExcPending
		if v85 != 0 {
			return int32(0)
		} else {
			return v62
		}
	}
}
func F_zslGetAllocSize(m *base.Module) int32 {
	return int32(272)
}
func F_zslInsert(m *base.Module, l0 int32, l1 float64, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v14 int32
	_ = v14
	var v23 int64
	_ = v23
	var v27 int64
	_ = v27
	var v28 int64
	_ = v28
	var v36 int64
	_ = v36
	var v38 int64
	_ = v38
	var v42 int32
	_ = v42
	var v46 int64
	_ = v46
	var v49 int64
	_ = v49
	var v51 int64
	_ = v51
	var v54 int64
	_ = v54
	var v65 int64
	_ = v65
	var v68 int64
	_ = v68
	var v79 int64
	_ = v79
	var v82 int64
	_ = v82
	var v95 int64
	_ = v95
	var v102 int64
	_ = v102
	var v107 int32
	_ = v107
	var v110 int64
	_ = v110
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v126 int64
	_ = v126
	var v134 int64
	_ = v134
	var v137 int64
	_ = v137
	var v149 int64
	_ = v149
	var v151 int32
	_ = v151
	var v153 int64
	_ = v153
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v170 int64
	_ = v170
	var v178 int64
	_ = v178
	var v181 int64
	_ = v181
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v195 int64
	_ = v195
	var v203 int64
	_ = v203
	var v205 int64
	_ = v205
	var v210 int64
	_ = v210
	var v219 int32
	_ = v219
	var v220 int64
	_ = v220
	var v231 int64
	_ = v231
	var v236 int64
	_ = v236
	var v241 int64
	_ = v241
	var v244 int64
	_ = v244
	var v247 int32
	_ = v247
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v257 int32
	_ = v257
	var v258 int32
	_ = v258
	var v259 int32
	_ = v259
	v14 = *(*int32)(unsafe.Add(mBase, _consts[1094]))
	if int32(311) < v14 {
		if v14 == int32(313) {
			v28 = int64(5489)
			*(*int64)(unsafe.Add(mBase, _consts[1095])) = v28
			v36 = int64(1)
			v38 = v28
			for {
				v42 = int32(3)
				v46 = int64(62)
				v49 = int64(6364136223846793005)
				v51 = (int64(base.Ui64(v38)>>(uint(v46)%64))^v38)*v49 + v36
				*(*int64)(unsafe.Add(mBase, uint32(base.I32_wrap_i64(v36)<<(uint(v42)%32))+uint32(_consts[1095]))) = v51
				v54 = v36 + int64(1)
				v65 = (int64(base.Ui64(v51)>>(uint(v46)%64))^v51)*v49 + v54
				*(*int64)(unsafe.Add(mBase, uint32(base.I32_wrap_i64(v54)<<(uint(v42)%32))+uint32(_consts[1095]))) = v65
				v68 = v36 + int64(2)
				v79 = (int64(base.Ui64(v65)>>(uint(v46)%64))^v65)*v49 + v68
				*(*int64)(unsafe.Add(mBase, uint32(base.I32_wrap_i64(v68)<<(uint(v42)%32))+uint32(_consts[1095]))) = v79
				v82 = v36 + int64(3)
				if v82 == int64(312) {
					v102 = v28
					break
				} else {
					v95 = (int64(base.Ui64(v79)>>(uint(int64(62))%64))^v79)*int64(6364136223846793005) + v82
					*(*int64)(unsafe.Add(mBase, uint32(base.I32_wrap_i64(v82)<<(uint(int32(3))%32))+uint32(_consts[1095]))) = v95
					v36 = v36 + int64(4)
					v38 = v95
					continue
				}
				break
			}
		} else {
			v27 = *(*int64)(unsafe.Add(mBase, _consts[1095]))
			v102 = v27
		}
		v107 = int32(0)
		v110 = v102
		for {
			v116 = int32(3)
			v117 = v107 << (uint(v116) % 32)
			v120 = int32(1)
			v121 = v107 + v120
			v126 = *(*int64)(unsafe.Add(mBase, uint32(v121<<(uint(v116)%32))+uint32(_consts[1095])))
			v134 = *(*int64)(unsafe.Add(mBase, uint32(base.I32_wrap_i64(v126)&v120<<(uint(v116)%32))+uint32(_consts[1096])))
			v137 = *(*int64)(unsafe.Add(mBase, uint32(v117)+uint32(_consts[1097])))
			*(*int64)(unsafe.Add(mBase, uint32(v117)+uint32(_consts[1095]))) = v134 ^ v137 ^ int64(base.Ui64(v110&int64(-2147483648)|v126&int64(2147483646))>>(uint(int64(1))%64))
			if v121 != int32(156) {
				v107 = v121
				v110 = v126
				continue
			} else {
				break
			}
			break
		}
		v149 = *(*int64)(unsafe.Add(mBase, _consts[1097]))
		v151 = int32(156)
		v153 = v149
		for {
			v160 = int32(3)
			v161 = v151 << (uint(v160) % 32)
			v164 = int32(1)
			v165 = v151 + v164
			v170 = *(*int64)(unsafe.Add(mBase, uint32(v165<<(uint(v160)%32))+uint32(_consts[1095])))
			v178 = *(*int64)(unsafe.Add(mBase, uint32(base.I32_wrap_i64(v170)&v164<<(uint(v160)%32))+uint32(_consts[1096])))
			v181 = *(*int64)(unsafe.Add(mBase, uint32(v161)+uint32(_consts[1098])))
			*(*int64)(unsafe.Add(mBase, uint32(v161)+uint32(_consts[1095]))) = v178 ^ v181 ^ int64(base.Ui64(v153&int64(-2147483648)|v170&int64(2147483646))>>(uint(int64(1))%64))
			if v165 != int32(311) {
				v151 = v165
				v153 = v170
				continue
			} else {
				break
			}
			break
		}
		v192 = int32(1)
		v193 = int32(0)
		v195 = *(*int64)(unsafe.Add(mBase, _consts[1095]))
		v203 = *(*int64)(unsafe.Add(mBase, uint32(base.I32_wrap_i64(v195)&v192<<(uint(int32(3))%32))+uint32(_consts[1096])))
		v205 = *(*int64)(unsafe.Add(mBase, _consts[1099]))
		v210 = *(*int64)(unsafe.Add(mBase, _consts[1100]))
		*(*int64)(unsafe.Add(mBase, _consts[1100])) = v203 ^ v205 ^ int64(base.Ui64(v195&int64(2147483646)|v210&int64(-2147483648))>>(uint(int64(1))%64))
		v219 = v192
		v220 = v195
	} else {
		v23 = *(*int64)(unsafe.Add(mBase, uint32(v14<<(uint(int32(3))%32))+uint32(_consts[1095])))
		v219 = v14 + int32(1)
		v220 = v23
	}
	*(*int32)(unsafe.Add(mBase, _consts[1094])) = v219
	v231 = int64(base.Ui64(v220)>>(uint(int64(29))%64))&int64(22906492245) ^ v220
	v236 = v231<<(uint(int64(17))%64)&int64(8202884508482404352) ^ v231
	v241 = v236<<(uint(int64(37))%64)&int64(-2270628950310912) ^ v236
	v244 = int64(base.Ui64(v241)>>(uint(int64(43))%64)) ^ v241
	v247 = int32(1)
	if v244 == int64(0) {
		v253 = int32(32)
	} else {
		v253 = int32(base.Ui32(base.I32_wrap_i64(base.I64_clz(v244)))>>(uint(v247)%32)) + v247
	}
	v254 = F_zslCreateNode(m, v253, l1, l2)
	mBase = m.M
	v257 = m.ExcPending
	if v257 != 0 {
		return int32(0)
	} else {
		v258 = F_zslInsertNode(m, l0, v254)
		mBase = m.M
		v259 = m.ExcPending
		if v259 != 0 {
			return int32(0)
		} else {
			return v258
		}
	}
}
func F_zslLexValueGteMin(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v48 int32
	_ = v48
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
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
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	if v7 == int32(0) {
		if l0 != v6 {
			v25 = int32(0)
			v27 = *(*int32)(unsafe.Add(mBase, _consts[376]))
			if l0 == v27 {
				v93 = v25
				return v93
			} else {
				v30 = *(*int32)(unsafe.Add(mBase, _consts[377]))
				if v6 == v30 {
					v93 = v25
					return v93
				} else {
					if v6 != v27 {
						if l0 == v30 {
							v93 = int32(1)
						} else {
							v38 = int32(-1)
							v41 = int32(0)
							v48 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0+int32(-1)))))
							switch v48 & int32(7) {
							case 0:
								v65 = int32(base.Ui32(v48) >> (uint(int32(3)) % 32))
							case 1:
								v55 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0+int32(-3)))))
								v65 = v55
							case 2:
								v58 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0+int32(-5)))))
								v65 = v58
							case 3:
								v61 = *(*int32)(unsafe.Add(mBase, uint32(l0+int32(-9))))
								v65 = v61
							case 4:
								v64 = *(*int32)(unsafe.Add(mBase, uint32(l0+int32(-17))))
								v65 = v64
							default:
								v65 = v41
							}
							v68 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6+int32(-1)))))
							switch v68 & int32(7) {
							case 0:
								v85 = int32(base.Ui32(v68) >> (uint(int32(3)) % 32))
							case 1:
								v75 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6+int32(-3)))))
								v85 = v75
							case 2:
								v78 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v6+int32(-5)))))
								v85 = v78
							case 3:
								v81 = *(*int32)(unsafe.Add(mBase, uint32(v6+int32(-9))))
								v85 = v81
							case 4:
								v84 = *(*int32)(unsafe.Add(mBase, uint32(v6+int32(-17))))
								v85 = v84
							default:
								v85 = v41
							}
							v86 = base.B2i32(base.Ui32(v65) < base.Ui32(v85))
							if base.Ui32(v65) < base.Ui32(v85) {
								v87 = v65
							} else {
								v87 = v85
							}
							v88 = F_memcmp(m, l0, v6, v87)
							mBase = m.M
							if v88 != 0 {
								v91 = v88
							} else {
								v91 = base.B2i32(base.Ui32(v85) < base.Ui32(v65)) - v86
							}
							v93 = base.B2i32(v38 < v91)
						}
						return v93
					} else {
						return int32(1)
					}
				}
			}
		} else {
			return int32(1)
		}
	} else {
		v10 = int32(0)
		if l0 == v6 {
			v93 = v10
		} else {
			v13 = *(*int32)(unsafe.Add(mBase, _consts[376]))
			if l0 == v13 {
				v93 = v10
			} else {
				v16 = *(*int32)(unsafe.Add(mBase, _consts[377]))
				if v6 == v16 {
					v93 = v10
				} else {
					v18 = int32(1)
					if v6 == v13 {
						v93 = v18
					} else {
						if l0 == v16 {
							v93 = v18
						} else {
							v38 = int32(0)
							v41 = int32(0)
							v48 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0+int32(-1)))))
							switch v48 & int32(7) {
							case 0:
								v65 = int32(base.Ui32(v48) >> (uint(int32(3)) % 32))
							case 1:
								v55 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0+int32(-3)))))
								v65 = v55
							case 2:
								v58 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0+int32(-5)))))
								v65 = v58
							case 3:
								v61 = *(*int32)(unsafe.Add(mBase, uint32(l0+int32(-9))))
								v65 = v61
							case 4:
								v64 = *(*int32)(unsafe.Add(mBase, uint32(l0+int32(-17))))
								v65 = v64
							default:
								v65 = v41
							}
							v68 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6+int32(-1)))))
							switch v68 & int32(7) {
							case 0:
								v85 = int32(base.Ui32(v68) >> (uint(int32(3)) % 32))
							case 1:
								v75 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6+int32(-3)))))
								v85 = v75
							case 2:
								v78 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v6+int32(-5)))))
								v85 = v78
							case 3:
								v81 = *(*int32)(unsafe.Add(mBase, uint32(v6+int32(-9))))
								v85 = v81
							case 4:
								v84 = *(*int32)(unsafe.Add(mBase, uint32(v6+int32(-17))))
								v85 = v84
							default:
								v85 = v41
							}
							v86 = base.B2i32(base.Ui32(v65) < base.Ui32(v85))
							if base.Ui32(v65) < base.Ui32(v85) {
								v87 = v65
							} else {
								v87 = v85
							}
							v88 = F_memcmp(m, l0, v6, v87)
							mBase = m.M
							if v88 != 0 {
								v91 = v88
							} else {
								v91 = base.B2i32(base.Ui32(v85) < base.Ui32(v65)) - v86
							}
							v93 = base.B2i32(v38 < v91)
						}
					}
				}
			}
		}
		return v93
	}
}
