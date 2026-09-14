package p0

import (
	base "github.com/shibukawa/vkmem/internal/aot/vkaot/base"
	"unsafe"
)

func F_a_cas_3(m *base.Module) int32 {
	mBase := m.M
	_ = mBase
	var v1 int32
	_ = v1
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	v1 = int32(0)
	v4 = *(*int32)(unsafe.Add(mBase, _consts[1241]))
	if v4 != 0 {
		v6 = v4
	} else {
		v6 = int32(1073741823)
	}
	*(*int32)(unsafe.Add(mBase, _consts[1241])) = v6
	return v4
}
func F_access(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v13 int32
	_ = v13
	v5 = m.Env.X__syscall_faccessat(m, int32(-100), l0, l1, int32(0))
	mBase = m.M
	if base.Ui32(v5) < base.Ui32(int32(-4095)) {
		v13 = v5
	} else {
		v8 = F___errno_location(m)
		mBase = m.M
		*(*int32)(unsafe.Add(mBase, uint32(v8))) = int32(0) - v5
		v13 = int32(-1)
	}
	return v13
}
func F_acos(m *base.Module, l0 float64) float64 {
	var v6 int64
	_ = v6
	var v11 int32
	_ = v11
	var v22 float64
	_ = v22
	var v34 float64
	_ = v34
	var v78 float64
	_ = v78
	var v79 float64
	_ = v79
	var v115 float64
	_ = v115
	var v121 float64
	_ = v121
	var v122 float64
	_ = v122
	var v158 float64
	_ = v158
	var v164 float64
	_ = v164
	var v167 float64
	_ = v167
	v6 = base.I64_reinterpret_f64(l0)
	v11 = base.I32_wrap_i64(int64(base.Ui64(v6)>>(uint(int64(32))%64))) & int32(2147483647)
	if base.Ui32(v11) < base.Ui32(int32(1072693248)) {
		if base.Ui32(int32(1071644671)) < base.Ui32(v11) {
			if int64(-1) < v6 {
				v121 = base.F64_mul(base.F64_sub(float64(1), l0), float64(0.5))
				v122 = base.F64_sqrt(v121)
				v158 = base.F64_reinterpret_i64(base.I64_reinterpret_f64(v122) & int64(-4294967296))
				v164 = base.F64_add(base.F64_add(base.F64_mul(v122, base.F64_div(base.F64_mul(v121, base.F64_add(base.F64_mul(v121, base.F64_add(base.F64_mul(v121, base.F64_add(base.F64_mul(v121, base.F64_add(base.F64_mul(v121, base.F64_add(base.F64_mul(v121, float64(3.479331075960212e-05)), float64(0.0007915349942898145))), float64(-0.04005553450067941))), float64(0.20121253213486293))), float64(-0.3255658186224009))), float64(0.16666666666666666))), base.F64_add(base.F64_mul(v121, base.F64_add(base.F64_mul(v121, base.F64_add(base.F64_mul(v121, base.F64_add(base.F64_mul(v121, float64(0.07703815055590194)), float64(-0.6882839716054533))), float64(2.0209457602335057))), float64(-2.403394911734414))), float64(1)))), base.F64_div(base.F64_sub(v121, base.F64_mul(v158, v158)), base.F64_add(v122, v158))), v158)
				v167 = base.F64_add(v164, v164)
				return v167
			} else {
				v78 = base.F64_mul(base.F64_add(l0, float64(1)), float64(0.5))
				v79 = base.F64_sqrt(v78)
				v115 = base.F64_sub(float64(1.5707963267948966), base.F64_add(v79, base.F64_add(base.F64_mul(v79, base.F64_div(base.F64_mul(v78, base.F64_add(base.F64_mul(v78, base.F64_add(base.F64_mul(v78, base.F64_add(base.F64_mul(v78, base.F64_add(base.F64_mul(v78, base.F64_add(base.F64_mul(v78, float64(3.479331075960212e-05)), float64(0.0007915349942898145))), float64(-0.04005553450067941))), float64(0.20121253213486293))), float64(-0.3255658186224009))), float64(0.16666666666666666))), base.F64_add(base.F64_mul(v78, base.F64_add(base.F64_mul(v78, base.F64_add(base.F64_mul(v78, base.F64_add(base.F64_mul(v78, float64(0.07703815055590194)), float64(-0.6882839716054533))), float64(2.0209457602335057))), float64(-2.403394911734414))), float64(1)))), float64(-6.123233995736766e-17))))
				return base.F64_add(v115, v115)
			}
		} else {
			if base.Ui32(v11) < base.Ui32(int32(1012924417)) {
				v167 = float64(1.5707963267948966)
				return v167
			} else {
				v34 = base.F64_mul(l0, l0)
				return base.F64_add(base.F64_sub(base.F64_sub(float64(6.123233995736766e-17), base.F64_mul(l0, base.F64_div(base.F64_mul(v34, base.F64_add(base.F64_mul(v34, base.F64_add(base.F64_mul(v34, base.F64_add(base.F64_mul(v34, base.F64_add(base.F64_mul(v34, base.F64_add(base.F64_mul(v34, float64(3.479331075960212e-05)), float64(0.0007915349942898145))), float64(-0.04005553450067941))), float64(0.20121253213486293))), float64(-0.3255658186224009))), float64(0.16666666666666666))), base.F64_add(base.F64_mul(v34, base.F64_add(base.F64_mul(v34, base.F64_add(base.F64_mul(v34, base.F64_add(base.F64_mul(v34, float64(0.07703815055590194)), float64(-0.6882839716054533))), float64(2.0209457602335057))), float64(-2.403394911734414))), float64(1))))), l0), float64(1.5707963267948966))
			}
		}
	} else {
		if v11+int32(-1072693248)|base.I32_wrap_i64(v6) != 0 {
			return base.F64_div(float64(0), base.F64_sub(l0, l0))
		} else {
			if int64(-1) < v6 {
				v22 = float64(0)
			} else {
				v22 = float64(3.141592653589793)
			}
			return v22
		}
	}
}
func F_activeDefragStringOb(m *base.Module, l0 int32) int32 {
	return int32(0)
}
func F_activeExpireCycleTryExpire(m *base.Module, l0 int32, l1 int32, l2 int64, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v12 int32
	_ = v12
	var v22 int64
	_ = v22
	var v23 int64
	_ = v23
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v35 int32
	_ = v35
	var v43 int64
	_ = v43
	var v45 int32
	_ = v45
	var v49 int64
	_ = v49
	var v53 int64
	_ = v53
	var v56 int32
	_ = v56
	var v64 int64
	_ = v64
	var v68 int32
	_ = v68
	var v72 int32
	_ = v72
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v92 int32
	_ = v92
	var v95 int32
	_ = v95
	var v102 int32
	_ = v102
	var v105 int32
	_ = v105
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
	var v118 int32
	_ = v118
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v123 int32
	_ = v123
	var v128 int32
	_ = v128
	var v136 int32
	_ = v136
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v12&int32(1) == int32(0) {
		v23 = int64(-1)
	} else {
		v22 = *(*int64)(unsafe.Add(mBase, uint32(l1+(v12&int32(4)^int32(12)))))
		v23 = v22
	}
	if v23 <= int64(-1) {
		F__serverAssert(m, int32(_a573), int32(_a574), int32(68))
		mBase = m.M
		v136 = m.ExcPending
		if v136 != 0 {
			return int32(0)
		} else {
			F_abort(m)
			mBase = m.M
			base.Wasm_trap_unreachable()
			for {
			}
		}
	} else {
		if l2 <= v23 {
			v128 = int32(0)
			return v128
		} else {
			v28 = int32(1)
			v31 = int32(0)
			v35 = *(*int32)(unsafe.Add(mBase, _consts[275]))
			*(*int32)(unsafe.Add(mBase, _consts[275])) = v35 + v28
			if v35 != 0 {
			} else {
				v43 = F_ustime(m)
				mBase = m.M
				v45 = int32(0)
				*(*int64)(unsafe.Add(mBase, _consts[277])) = v43
				v49 = base.I64_div_s(v43, int64(1000))
				*(*int64)(unsafe.Add(mBase, _consts[32])) = v49
				v53 = base.I64_div_s(v43, int64(1000000))
				*(*int64)(unsafe.Add(mBase, _consts[37])) = v53
				v56 = *(*int32)(unsafe.Add(mBase, _consts[167]))
				F_lrulfu_updateClockAndPolicy(m, v49, int32(base.Ui32(v56&int32(2))>>(uint(int32(1))%32)))
				mBase = m.M
				v64 = *(*int64)(unsafe.Add(mBase, _consts[32]))
				*(*int64)(unsafe.Add(mBase, _consts[78])) = v64
			}
			v68 = int32(0)
			v72 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
			if v72&int32(2) == v68 {
				v92 = v68
			} else {
				v86 = l1 + (v72&int32(4) ^ int32(12)) + v72<<(uint(int32(3))%32)&int32(8)
				v87 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v86))))
				v92 = v86 + v87 + int32(1)
			}
			v95 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v92+int32(-1)))))
			switch v95 & int32(7) {
			case 0:
				v112 = int32(base.Ui32(v95) >> (uint(int32(3)) % 32))
			case 1:
				v102 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v92+int32(-3)))))
				v112 = v102
			case 2:
				v105 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v92+int32(-5)))))
				v112 = v105
			case 3:
				v108 = *(*int32)(unsafe.Add(mBase, uint32(v92+int32(-9))))
				v112 = v108
			case 4:
				v111 = *(*int32)(unsafe.Add(mBase, uint32(v92+int32(-17))))
				v112 = v111
			default:
				v112 = v68
			}
			v113 = F_createStringObject_1(m, v92, v112)
			mBase = m.M
			v116 = m.ExcPending
			if v116 != 0 {
				return int32(0)
			} else {
				F_deleteExpiredKeyAndPropagateWithDictIndex(m, l0, v113, l3)
				mBase = m.M
				v118 = m.ExcPending
				if v118 != 0 {
					return int32(0)
				} else {
					F_decrRefCount(m, v113)
					mBase = m.M
					v120 = m.ExcPending
					if v120 != 0 {
						return int32(0)
					} else {
						v121 = int32(0)
						v123 = *(*int32)(unsafe.Add(mBase, _consts[275]))
						*(*int32)(unsafe.Add(mBase, _consts[275])) = v123 + int32(-1)
						v128 = v28
						return v128
					}
				}
			}
		}
	}
}
func F_addAuthErrReply(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	v3 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+207)))
	if v3&int32(64) != 0 {
		return
	} else {
		if l1 != 0 {
			v7 = F_objectGetVal(m, l1)
			mBase = m.M
			v8 = v7
		} else {
			v8 = int32(_a48)
		}
		F_addReplyError(m, l0, v8)
		mBase = m.M
		v10 = m.ExcPending
		if v10 != 0 {
			return
		} else {
			return
		}
	}
}
func F_addWritePreparedReplyArrayLen(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v30 int64
	_ = v30
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v50 int64
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v71 int32
	_ = v71
	var v75 int32
	_ = v75
	var v80 int32
	_ = v80
	var v89 int32
	_ = v89
	v4 = m.G0
	v6 = v4 - int32(128)
	m.G0 = v6
	if l1 <= int32(-1) {
		F__serverAssert(m, int32(_a789), int32(_a774), int32(1431))
		mBase = m.M
		v89 = m.ExcPending
		if v89 != 0 {
			return
		} else {
			F_abort(m)
			mBase = m.M
			base.Wasm_trap_unreachable()
			for {
			}
		}
	} else {
		if base.Ui32(int32(31)) < base.Ui32(l1) {
			v25 = int32(42)
			*(*uint8)(unsafe.Add(mBase, uint32(v6))) = uint8(v25)
			v28 = v6 | int32(1)
			v30 = base.I64_extend_i32_u(l1)
			if v30 <= int64(-1) {
				v39 = int32(45)
				*(*uint8)(unsafe.Add(mBase, uint32(v28))) = uint8(v39)
				v43 = int32(1)
				v48 = v28 + v43
				v49 = int32(126)
				v50 = int64(0) - v30
				v51 = v43
			} else {
				v48 = v28
				v49 = int32(127)
				v50 = v30
				v51 = int32(0)
			}
			v52 = F_ull2string(m, v48, v49, v50)
			mBase = m.M
			if v52 == int32(0) {
				v71 = int32(0)
			} else {
				v71 = v52 + v51
			}
			v75 = int32(2573)
			*(*uint16)(unsafe.Add(mBase, uint32(v71+v6+int32(1)))) = uint16(v75)
			F__addReplyToBufferOrList(m, l0, v6, v71+int32(3))
			mBase = m.M
			v80 = m.ExcPending
			if v80 != 0 {
				return
			} else {
				m.G0 = v6 + int32(128)
				return
			}
		} else {
			v16 = *(*int32)(unsafe.Add(mBase, uint32(l1<<(uint(int32(2))%32))+uint32(_consts[424])))
			v17 = F_objectGetVal(m, v16)
			mBase = m.M
			if base.Ui32(l1) < base.Ui32(int32(10)) {
				v22 = int32(4)
			} else {
				v22 = int32(5)
			}
			F__addReplyToBufferOrList(m, l0, v17, v22)
			mBase = m.M
			v24 = m.ExcPending
			if v24 != 0 {
				return
			} else {
				m.G0 = v6 + int32(128)
				return
			}
		}
	}
}
func F_addWritePreparedReplyBulkLongLong(m *base.Module, l0 int32, l1 int64) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v17 int32
	_ = v17
	var v21 int32
	_ = v21
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int64
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	v4 = m.G0
	v5 = int32(64)
	v6 = v4 - v5
	m.G0 = v6
	if l1 <= int64(-1) {
		v17 = int32(45)
		*(*uint8)(unsafe.Add(mBase, uint32(v6))) = uint8(v17)
		v21 = int32(1)
		v26 = v6 + v21
		v27 = int32(63)
		v28 = int64(0) - l1
		v29 = v21
	} else {
		v26 = v6
		v27 = v5
		v28 = l1
		v29 = int32(0)
	}
	v30 = F_ull2string(m, v26, v27, v28)
	mBase = m.M
	if v30 == int32(0) {
		v49 = int32(0)
	} else {
		v49 = v30 + v29
	}
	F_addWritePreparedReplyBulkCBuffer(m, l0, v6, v49)
	mBase = m.M
	v51 = m.ExcPending
	if v51 != 0 {
		return
	} else {
		m.G0 = v6 + int32(64)
		return
	}
}
func F_adjust_assign(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
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
	var v11 int32
	_ = v11
	var v15 int32
	_ = v15
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	v5 = l1 - l2
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	switch v7 {
	case 0:
		if v5 < int32(1) {
			return
		} else {
			v24 = *(*int32)(unsafe.Add(mBase, uint32(v6)+36))
			F_luaK_reserveregs(m, v6, v5)
			mBase = m.M
			v26 = m.ExcPending
			if v26 != 0 {
				return
			} else {
				F_luaK_nil(m, v6, v24, v5)
				mBase = m.M
				v28 = m.ExcPending
				if v28 != 0 {
					return
				} else {
					return
				}
			}
		}
	default:
		F_luaK_exp2nextreg(m, v6, l3)
		mBase = m.M
		v21 = m.ExcPending
		if v21 != 0 {
			return
		} else {
			if v5 < int32(1) {
				return
			} else {
				v24 = *(*int32)(unsafe.Add(mBase, uint32(v6)+36))
				F_luaK_reserveregs(m, v6, v5)
				mBase = m.M
				v26 = m.ExcPending
				if v26 != 0 {
					return
				} else {
					F_luaK_nil(m, v6, v24, v5)
					mBase = m.M
					v28 = m.ExcPending
					if v28 != 0 {
						return
					} else {
						return
					}
				}
			}
		}
	case 13, 14:
		v8 = int32(-1)
		if v8 < v5 {
			v11 = v5
		} else {
			v11 = v8
		}
		F_luaK_setreturns(m, v6, l3, v11+int32(1))
		mBase = m.M
		v15 = m.ExcPending
		if v15 != 0 {
			return
		} else {
			if v5 < int32(1) {
				return
			} else {
				F_luaK_reserveregs(m, v6, v11)
				mBase = m.M
				v19 = m.ExcPending
				if v19 != 0 {
					return
				} else {
					return
				}
			}
		}
	}
}
func F_afterSleep(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v21 int32
	_ = v21
	var v22 int64
	_ = v22
	var v28 int64
	_ = v28
	var v29 int64
	_ = v29
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v42 int32
	_ = v42
	var v44 int64
	_ = v44
	var v47 int64
	_ = v47
	var v49 int64
	_ = v49
	var v52 int64
	_ = v52
	var v58 int32
	_ = v58
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v64 int64
	_ = v64
	var v73 int64
	_ = v73
	var v78 int32
	_ = v78
	var v79 int64
	_ = v79
	var v82 int64
	_ = v82
	var v83 int64
	_ = v83
	var v87 int64
	_ = v87
	var v90 int32
	_ = v90
	var v94 int32
	_ = v94
	var v99 int64
	_ = v99
	var v103 int64
	_ = v103
	var v108 int32
	_ = v108
	var v109 int64
	_ = v109
	var v115 int32
	_ = v115
	var v117 int32
	_ = v117
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v123 int64
	_ = v123
	var v126 int32
	_ = v126
	v7 = m.G0
	v9 = v7 - int32(64)
	m.G0 = v9
	v12 = *(*int32)(unsafe.Add(mBase, _consts[435]))
	if v12 != 0 {
		v78 = int32(0)
		v79 = F_ustime(m)
		mBase = m.M
		*(*int64)(unsafe.Add(mBase, _consts[277])) = v79
		v82 = int64(1000)
		v83 = base.I64_div_s(v79, v82)
		*(*int64)(unsafe.Add(mBase, _consts[32])) = v83
		v87 = base.I64_div_s(v79, int64(1000000))
		*(*int64)(unsafe.Add(mBase, _consts[37])) = v87
		v90 = *(*int32)(unsafe.Add(mBase, _consts[167]))
		v94 = int32(base.Ui32(v90&int32(2)) >> (uint(int32(1)) % 32))
		*(*uint8)(unsafe.Add(mBase, _consts[334])) = uint8(v94)
		v99 = base.I64_div_s(v83, int64(60000))
		*(*uint16)(unsafe.Add(mBase, _consts[335])) = uint16(v99)
		v103 = base.I64_div_s(v83, v82)
		*(*int32)(unsafe.Add(mBase, _consts[333])) = base.I32_wrap_i64(v103) & int32(16777215)
		v108 = int32(0)
		v109 = *(*int64)(unsafe.Add(mBase, _consts[37]))
		*(*int64)(unsafe.Add(mBase, uint32(v9)+8)) = v109
		v115 = F___localtime_r(m, v9+int32(8), v9+int32(20))
		mBase = m.M
		v117 = *(*int32)(unsafe.Add(mBase, uint32(v9)+52))
		*(*int32)(unsafe.Add(mBase, _consts[685])) = v117
		v120 = *(*int32)(unsafe.Add(mBase, _consts[435]))
		if v120 != 0 {
		} else {
			v121 = int32(0)
			v123 = *(*int64)(unsafe.Add(mBase, _consts[32]))
			*(*int64)(unsafe.Add(mBase, _consts[78])) = v123
		}
		F_IOThreadsAfterSleep(m, l1)
		mBase = m.M
		v126 = m.ExcPending
		if v126 != 0 {
			return
		} else {
			m.G0 = v9 + int32(64)
			return
		}
	} else {
		v15 = *(*int32)(unsafe.Add(mBase, _consts[177]))
		v16 = *(*int32)(unsafe.Add(mBase, uint32(v15)+16))
		v17 = *(*int32)(unsafe.Add(mBase, uint32(v15)+12))
		if v16+v17 == int32(0) {
			v62 = int32(0)
			v63 = *(*int32)(unsafe.Add(mBase, _consts[271]))
			v64 = m.T0[v63].(func(*base.Module) int64)(m)
			mBase = m.M
			*(*uint8)(unsafe.Add(mBase, _consts[688])) = uint8(base.B2i32(v62 < l1))
			*(*int64)(unsafe.Add(mBase, _consts[689])) = v64
			v73 = *(*int64)(unsafe.Add(mBase, _consts[690]))
			*(*int64)(unsafe.Add(mBase, _consts[691])) = v73
			v78 = int32(0)
			v79 = F_ustime(m)
			mBase = m.M
			*(*int64)(unsafe.Add(mBase, _consts[277])) = v79
			v82 = int64(1000)
			v83 = base.I64_div_s(v79, v82)
			*(*int64)(unsafe.Add(mBase, _consts[32])) = v83
			v87 = base.I64_div_s(v79, int64(1000000))
			*(*int64)(unsafe.Add(mBase, _consts[37])) = v87
			v90 = *(*int32)(unsafe.Add(mBase, _consts[167]))
			v94 = int32(base.Ui32(v90&int32(2)) >> (uint(int32(1)) % 32))
			*(*uint8)(unsafe.Add(mBase, _consts[334])) = uint8(v94)
			v99 = base.I64_div_s(v83, int64(60000))
			*(*uint16)(unsafe.Add(mBase, _consts[335])) = uint16(v99)
			v103 = base.I64_div_s(v83, v82)
			*(*int32)(unsafe.Add(mBase, _consts[333])) = base.I32_wrap_i64(v103) & int32(16777215)
			v108 = int32(0)
			v109 = *(*int64)(unsafe.Add(mBase, _consts[37]))
			*(*int64)(unsafe.Add(mBase, uint32(v9)+8)) = v109
			v115 = F___localtime_r(m, v9+int32(8), v9+int32(20))
			mBase = m.M
			v117 = *(*int32)(unsafe.Add(mBase, uint32(v9)+52))
			*(*int32)(unsafe.Add(mBase, _consts[685])) = v117
			v120 = *(*int32)(unsafe.Add(mBase, _consts[435]))
			if v120 != 0 {
			} else {
				v121 = int32(0)
				v123 = *(*int64)(unsafe.Add(mBase, _consts[32]))
				*(*int64)(unsafe.Add(mBase, _consts[78])) = v123
			}
			F_IOThreadsAfterSleep(m, l1)
			mBase = m.M
			v126 = m.ExcPending
			if v126 != 0 {
				return
			} else {
				m.G0 = v9 + int32(64)
				return
			}
		} else {
			v21 = int32(0)
			v22 = *(*int64)(unsafe.Add(mBase, _consts[270]))
			if base.B2i32(v22 == int64(0)) == v21 {
				v28 = F_ustime(m)
				mBase = m.M
				v29 = v28
			} else {
				v29 = int64(0)
			}
			*(*int32)(unsafe.Add(mBase, _consts[692])) = int32(1)
			v34 = F___pthread_mutex_lock(m, int32(_a1094))
			mBase = m.M
			v35 = int32(0)
			*(*int32)(unsafe.Add(mBase, _consts[692])) = v35
			F_moduleFireServerEvent(m, int64(15), int32(1), v35)
			mBase = m.M
			v42 = m.ExcPending
			if v42 != 0 {
				return
			} else {
				v44 = *(*int64)(unsafe.Add(mBase, _consts[270]))
				if v44 == int64(0) {
					v62 = int32(0)
					v63 = *(*int32)(unsafe.Add(mBase, _consts[271]))
					v64 = m.T0[v63].(func(*base.Module) int64)(m)
					mBase = m.M
					*(*uint8)(unsafe.Add(mBase, _consts[688])) = uint8(base.B2i32(v62 < l1))
					*(*int64)(unsafe.Add(mBase, _consts[689])) = v64
					v73 = *(*int64)(unsafe.Add(mBase, _consts[690]))
					*(*int64)(unsafe.Add(mBase, _consts[691])) = v73
					v78 = int32(0)
					v79 = F_ustime(m)
					mBase = m.M
					*(*int64)(unsafe.Add(mBase, _consts[277])) = v79
					v82 = int64(1000)
					v83 = base.I64_div_s(v79, v82)
					*(*int64)(unsafe.Add(mBase, _consts[32])) = v83
					v87 = base.I64_div_s(v79, int64(1000000))
					*(*int64)(unsafe.Add(mBase, _consts[37])) = v87
					v90 = *(*int32)(unsafe.Add(mBase, _consts[167]))
					v94 = int32(base.Ui32(v90&int32(2)) >> (uint(int32(1)) % 32))
					*(*uint8)(unsafe.Add(mBase, _consts[334])) = uint8(v94)
					v99 = base.I64_div_s(v83, int64(60000))
					*(*uint16)(unsafe.Add(mBase, _consts[335])) = uint16(v99)
					v103 = base.I64_div_s(v83, v82)
					*(*int32)(unsafe.Add(mBase, _consts[333])) = base.I32_wrap_i64(v103) & int32(16777215)
					v108 = int32(0)
					v109 = *(*int64)(unsafe.Add(mBase, _consts[37]))
					*(*int64)(unsafe.Add(mBase, uint32(v9)+8)) = v109
					v115 = F___localtime_r(m, v9+int32(8), v9+int32(20))
					mBase = m.M
					v117 = *(*int32)(unsafe.Add(mBase, uint32(v9)+52))
					*(*int32)(unsafe.Add(mBase, _consts[685])) = v117
					v120 = *(*int32)(unsafe.Add(mBase, _consts[435]))
					if v120 != 0 {
					} else {
						v121 = int32(0)
						v123 = *(*int64)(unsafe.Add(mBase, _consts[32]))
						*(*int64)(unsafe.Add(mBase, _consts[78])) = v123
					}
					F_IOThreadsAfterSleep(m, l1)
					mBase = m.M
					v126 = m.ExcPending
					if v126 != 0 {
						return
					} else {
						m.G0 = v9 + int32(64)
						return
					}
				} else {
					v47 = F_ustime(m)
					mBase = m.M
					v49 = *(*int64)(unsafe.Add(mBase, _consts[270]))
					if v49 == int64(0) {
						v62 = int32(0)
						v63 = *(*int32)(unsafe.Add(mBase, _consts[271]))
						v64 = m.T0[v63].(func(*base.Module) int64)(m)
						mBase = m.M
						*(*uint8)(unsafe.Add(mBase, _consts[688])) = uint8(base.B2i32(v62 < l1))
						*(*int64)(unsafe.Add(mBase, _consts[689])) = v64
						v73 = *(*int64)(unsafe.Add(mBase, _consts[690]))
						*(*int64)(unsafe.Add(mBase, _consts[691])) = v73
						v78 = int32(0)
						v79 = F_ustime(m)
						mBase = m.M
						*(*int64)(unsafe.Add(mBase, _consts[277])) = v79
						v82 = int64(1000)
						v83 = base.I64_div_s(v79, v82)
						*(*int64)(unsafe.Add(mBase, _consts[32])) = v83
						v87 = base.I64_div_s(v79, int64(1000000))
						*(*int64)(unsafe.Add(mBase, _consts[37])) = v87
						v90 = *(*int32)(unsafe.Add(mBase, _consts[167]))
						v94 = int32(base.Ui32(v90&int32(2)) >> (uint(int32(1)) % 32))
						*(*uint8)(unsafe.Add(mBase, _consts[334])) = uint8(v94)
						v99 = base.I64_div_s(v83, int64(60000))
						*(*uint16)(unsafe.Add(mBase, _consts[335])) = uint16(v99)
						v103 = base.I64_div_s(v83, v82)
						*(*int32)(unsafe.Add(mBase, _consts[333])) = base.I32_wrap_i64(v103) & int32(16777215)
						v108 = int32(0)
						v109 = *(*int64)(unsafe.Add(mBase, _consts[37]))
						*(*int64)(unsafe.Add(mBase, uint32(v9)+8)) = v109
						v115 = F___localtime_r(m, v9+int32(8), v9+int32(20))
						mBase = m.M
						v117 = *(*int32)(unsafe.Add(mBase, uint32(v9)+52))
						*(*int32)(unsafe.Add(mBase, _consts[685])) = v117
						v120 = *(*int32)(unsafe.Add(mBase, _consts[435]))
						if v120 != 0 {
						} else {
							v121 = int32(0)
							v123 = *(*int64)(unsafe.Add(mBase, _consts[32]))
							*(*int64)(unsafe.Add(mBase, _consts[78])) = v123
						}
						F_IOThreadsAfterSleep(m, l1)
						mBase = m.M
						v126 = m.ExcPending
						if v126 != 0 {
							return
						} else {
							m.G0 = v9 + int32(64)
							return
						}
					} else {
						v52 = v47 - v29
						if v52 < v49*int64(1000) {
							v62 = int32(0)
							v63 = *(*int32)(unsafe.Add(mBase, _consts[271]))
							v64 = m.T0[v63].(func(*base.Module) int64)(m)
							mBase = m.M
							*(*uint8)(unsafe.Add(mBase, _consts[688])) = uint8(base.B2i32(v62 < l1))
							*(*int64)(unsafe.Add(mBase, _consts[689])) = v64
							v73 = *(*int64)(unsafe.Add(mBase, _consts[690]))
							*(*int64)(unsafe.Add(mBase, _consts[691])) = v73
							v78 = int32(0)
							v79 = F_ustime(m)
							mBase = m.M
							*(*int64)(unsafe.Add(mBase, _consts[277])) = v79
							v82 = int64(1000)
							v83 = base.I64_div_s(v79, v82)
							*(*int64)(unsafe.Add(mBase, _consts[32])) = v83
							v87 = base.I64_div_s(v79, int64(1000000))
							*(*int64)(unsafe.Add(mBase, _consts[37])) = v87
							v90 = *(*int32)(unsafe.Add(mBase, _consts[167]))
							v94 = int32(base.Ui32(v90&int32(2)) >> (uint(int32(1)) % 32))
							*(*uint8)(unsafe.Add(mBase, _consts[334])) = uint8(v94)
							v99 = base.I64_div_s(v83, int64(60000))
							*(*uint16)(unsafe.Add(mBase, _consts[335])) = uint16(v99)
							v103 = base.I64_div_s(v83, v82)
							*(*int32)(unsafe.Add(mBase, _consts[333])) = base.I32_wrap_i64(v103) & int32(16777215)
							v108 = int32(0)
							v109 = *(*int64)(unsafe.Add(mBase, _consts[37]))
							*(*int64)(unsafe.Add(mBase, uint32(v9)+8)) = v109
							v115 = F___localtime_r(m, v9+int32(8), v9+int32(20))
							mBase = m.M
							v117 = *(*int32)(unsafe.Add(mBase, uint32(v9)+52))
							*(*int32)(unsafe.Add(mBase, _consts[685])) = v117
							v120 = *(*int32)(unsafe.Add(mBase, _consts[435]))
							if v120 != 0 {
							} else {
								v121 = int32(0)
								v123 = *(*int64)(unsafe.Add(mBase, _consts[32]))
								*(*int64)(unsafe.Add(mBase, _consts[78])) = v123
							}
							F_IOThreadsAfterSleep(m, l1)
							mBase = m.M
							v126 = m.ExcPending
							if v126 != 0 {
								return
							} else {
								m.G0 = v9 + int32(64)
								return
							}
						} else {
							F_latencyAddSample(m, int32(_a1250), v52)
							mBase = m.M
							v58 = m.ExcPending
							if v58 != 0 {
								return
							} else {
								v62 = int32(0)
								v63 = *(*int32)(unsafe.Add(mBase, _consts[271]))
								v64 = m.T0[v63].(func(*base.Module) int64)(m)
								mBase = m.M
								*(*uint8)(unsafe.Add(mBase, _consts[688])) = uint8(base.B2i32(v62 < l1))
								*(*int64)(unsafe.Add(mBase, _consts[689])) = v64
								v73 = *(*int64)(unsafe.Add(mBase, _consts[690]))
								*(*int64)(unsafe.Add(mBase, _consts[691])) = v73
								v78 = int32(0)
								v79 = F_ustime(m)
								mBase = m.M
								*(*int64)(unsafe.Add(mBase, _consts[277])) = v79
								v82 = int64(1000)
								v83 = base.I64_div_s(v79, v82)
								*(*int64)(unsafe.Add(mBase, _consts[32])) = v83
								v87 = base.I64_div_s(v79, int64(1000000))
								*(*int64)(unsafe.Add(mBase, _consts[37])) = v87
								v90 = *(*int32)(unsafe.Add(mBase, _consts[167]))
								v94 = int32(base.Ui32(v90&int32(2)) >> (uint(int32(1)) % 32))
								*(*uint8)(unsafe.Add(mBase, _consts[334])) = uint8(v94)
								v99 = base.I64_div_s(v83, int64(60000))
								*(*uint16)(unsafe.Add(mBase, _consts[335])) = uint16(v99)
								v103 = base.I64_div_s(v83, v82)
								*(*int32)(unsafe.Add(mBase, _consts[333])) = base.I32_wrap_i64(v103) & int32(16777215)
								v108 = int32(0)
								v109 = *(*int64)(unsafe.Add(mBase, _consts[37]))
								*(*int64)(unsafe.Add(mBase, uint32(v9)+8)) = v109
								v115 = F___localtime_r(m, v9+int32(8), v9+int32(20))
								mBase = m.M
								v117 = *(*int32)(unsafe.Add(mBase, uint32(v9)+52))
								*(*int32)(unsafe.Add(mBase, _consts[685])) = v117
								v120 = *(*int32)(unsafe.Add(mBase, _consts[435]))
								if v120 != 0 {
								} else {
									v121 = int32(0)
									v123 = *(*int64)(unsafe.Add(mBase, _consts[32]))
									*(*int64)(unsafe.Add(mBase, _consts[78])) = v123
								}
								F_IOThreadsAfterSleep(m, l1)
								mBase = m.M
								v126 = m.ExcPending
								if v126 != 0 {
									return
								} else {
									m.G0 = v9 + int32(64)
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
func F_allPersistenceDisabled(m *base.Module) int32 {
	mBase := m.M
	_ = mBase
	var v1 int32
	_ = v1
	var v2 int32
	_ = v2
	var v4 int32
	_ = v4
	v1 = int32(0)
	v2 = *(*int32)(unsafe.Add(mBase, _consts[174]))
	v4 = *(*int32)(unsafe.Add(mBase, _consts[27]))
	return base.B2i32(v2|v4 == v1)
}
func F_alsoPropagate(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) {
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
	var v23 int32
	_ = v23
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v37 int32
	_ = v37
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v88 int32
	_ = v88
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v100 int32
	_ = v100
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v113 int32
	_ = v113
	var v117 int32
	_ = v117
	var v120 int32
	_ = v120
	var v123 int32
	_ = v123
	var v125 int32
	_ = v125
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v139 int32
	_ = v139
	var v143 int32
	_ = v143
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v150 int32
	_ = v150
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v165 int32
	_ = v165
	if l3 == int32(0) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v12 = int32(0)
	v13 = *(*int32)(unsafe.Add(mBase, _consts[235]))
	if v13 == v12 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v17 = *(*int32)(unsafe.Add(mBase, _consts[116]))
	if v17 != 0 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v19 = l3 & int32(1)
	if v19 == int32(0) {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	v83 = int32(0)
	v85 = *(*int32)(unsafe.Add(mBase, _consts[88]))
	if v85 == v83 {
		v100 = l3
		goto L25
	} else {
		goto L26
	}
L6:
	;
	if l3&int32(2) == int32(0) {
		goto L1
	} else {
		goto L9
	}
L7:
	;
	v23 = *(*int32)(unsafe.Add(mBase, _consts[27]))
	if v23 != 0 {
		goto L5
	} else {
		goto L8
	}
L8:
	;
	goto L6
L9:
	;
	v28 = int32(0)
	v31 = m.G0
	v33 = v31 - int32(16)
	m.G0 = v33
	v37 = *(*int32)(unsafe.Add(mBase, _consts[63]))
	if v37 == v28 {
		v69 = v28
		goto L11
	} else {
		goto L12
	}
L10:
	;
	if v69 != 0 {
		goto L5
	} else {
		goto L21
	}
L11:
	;
	m.G0 = v33 + int32(16)
	goto L10
L12:
	;
	v41 = *(*int32)(unsafe.Add(mBase, _consts[90]))
	if v41 == int32(0) {
		v69 = v28
		goto L11
	} else {
		goto L13
	}
L13:
	;
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v41)+uint32(_consts[143])))
	v46 = v33 + int32(8)
	F_listRewind(m, v44, v46)
	mBase = m.M
	v48 = int32(0)
	v51 = F_listNext(m, v46)
	mBase = m.M
	if v51 == v48 {
		v69 = v48
		goto L11
	} else {
		goto L14
	}
L14:
	;
	v56 = v51
	goto L15
L15:
	;
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v56)+8))
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v57)))
	if v58 != 0 {
		goto L17
	} else {
		goto L18
	}
L16:
	;
	v69 = v48
	goto L11
L17:
	;
	v67 = F_listNext(m, v33+int32(8))
	mBase = m.M
	if v67 != 0 {
		v56 = v67
		goto L15
	} else {
		goto L20
	}
L18:
	;
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v57)+156))
	if base.Ui32(v59+int32(-18)) <= base.Ui32(int32(2)) {
		goto L17
	} else {
		goto L19
	}
L19:
	;
	v69 = int32(1)
	goto L11
L20:
	;
	goto L16
L21:
	;
	v75 = *(*int32)(unsafe.Add(mBase, _consts[64]))
	if v75 != 0 {
		goto L1
	} else {
		goto L22
	}
L22:
	;
	v77 = *(*int32)(unsafe.Add(mBase, _consts[448]))
	if v77 != 0 {
		goto L5
	} else {
		goto L23
	}
L23:
	;
	v78 = int32(0)
	v79 = *(*int32)(unsafe.Add(mBase, _consts[188]))
	v80 = *(*int32)(unsafe.Add(mBase, uint32(v79)+20))
	if v80 == v78 {
		goto L1
	} else {
		goto L24
	}
L24:
	;
	goto L5
L25:
	;
	v103 = F_valkey_malloc(m, l2<<(uint(int32(2))%32))
	mBase = m.M
	v104 = m.ExcPending
	if v104 != 0 {
		goto L29
	} else {
		goto L30
	}
L26:
	;
	v88 = *(*int32)(unsafe.Add(mBase, uint32(v85)+216))
	if v88 == int32(0) {
		v100 = l3
		goto L25
	} else {
		goto L27
	}
L27:
	;
	v91 = int32(0)
	v92 = *(*int32)(unsafe.Add(mBase, _consts[27]))
	if v19&base.B2i32(v92 != v91) == v91 {
		goto L1
	} else {
		goto L28
	}
L28:
	;
	v100 = l3 & int32(-3)
	goto L25
L29:
	;
	return
L30:
	;
	if l2 < int32(1) {
		goto L31
	} else {
		goto L32
	}
L31:
	;
	v136 = int32(0)
	v137 = *(*int32)(unsafe.Add(mBase, _consts[842]))
	v139 = *(*int32)(unsafe.Add(mBase, _consts[272]))
	if v139 != 0 {
		goto L39
	} else {
		goto L40
	}
L32:
	;
	v113 = v83
	goto L33
L33:
	;
	v117 = v113 << (uint(int32(2)) % 32)
	v120 = *(*int32)(unsafe.Add(mBase, uint32(l1+v117)))
	*(*int32)(unsafe.Add(mBase, uint32(v103+v117))) = v120
	F_incrRefCount(m, v120)
	mBase = m.M
	v123 = m.ExcPending
	if v123 != 0 {
		goto L29
	} else {
		goto L35
	}
L34:
	;
	goto L31
L35:
	;
	v125 = v113 + int32(1)
	if v125 != l2 {
		v113 = v125
		goto L33
	} else {
		goto L36
	}
L36:
	;
	goto L34
L37:
	;
	v165 = v161 + v160*int32(20)
	*(*int32)(unsafe.Add(mBase, uint32(v165)+16)) = l4
	*(*int32)(unsafe.Add(mBase, uint32(v165)+12)) = v100
	*(*int32)(unsafe.Add(mBase, uint32(v165)+4)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(v165))) = v103
	*(*int32)(unsafe.Add(mBase, uint32(v165)+8)) = l0
	*(*int32)(unsafe.Add(mBase, _consts[272])) = v160 + int32(1)
	goto L1
L38:
	;
	v147 = int32(0)
	*(*int32)(unsafe.Add(mBase, _consts[842])) = v146
	v150 = *(*int32)(unsafe.Add(mBase, _consts[843]))
	if v137 == v146 {
		v160 = v139
		v161 = v150
		goto L37
	} else {
		goto L43
	}
L39:
	;
	if v137 <= v139 {
		goto L41
	} else {
		goto L42
	}
L40:
	;
	v146 = int32(16)
	goto L38
L41:
	;
	v146 = v137 << (uint(int32(1)) % 32)
	goto L38
L42:
	;
	v143 = *(*int32)(unsafe.Add(mBase, _consts[843]))
	v160 = v139
	v161 = v143
	goto L37
L43:
	;
	v155 = F_valkey_realloc(m, v150, v146*int32(20))
	mBase = m.M
	v156 = m.ExcPending
	if v156 != 0 {
		goto L29
	} else {
		goto L44
	}
L44:
	;
	*(*int32)(unsafe.Add(mBase, _consts[843])) = v155
	v159 = *(*int32)(unsafe.Add(mBase, _consts[272]))
	v160 = v159
	v161 = v155
	goto L37
}
func F_analyzeLatencyForEvent(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int64
	_ = v23
	var v37 int32
	_ = v37
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v52 int32
	_ = v52
	var v53 int64
	_ = v53
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v59 int64
	_ = v59
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v84 int64
	_ = v84
	var v89 int64
	_ = v89
	var v91 int64
	_ = v91
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v95 int64
	_ = v95
	var v99 int32
	_ = v99
	var v105 int64
	_ = v105
	var v109 int64
	_ = v109
	var v110 int64
	_ = v110
	var v113 int64
	_ = v113
	var v119 int32
	_ = v119
	var v131 int64
	_ = v131
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v143 int32
	_ = v143
	var v148 int64
	_ = v148
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v159 int32
	_ = v159
	var v164 int64
	_ = v164
	var v166 int32
	_ = v166
	var v169 int32
	_ = v169
	var v173 int64
	_ = v173
	v16 = *(*int32)(unsafe.Add(mBase, _consts[322]))
	v17 = F_dictFetchValue(m, v16, l0)
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		return
	} else {
		if v17 == int32(0) {
			v22 = int32(0)
		} else {
			v21 = *(*int32)(unsafe.Add(mBase, uint32(v17)+4))
			v22 = v21
		}
		v23 = int64(0)
		*(*int64)(unsafe.Add(mBase, uint32(l1)+4)) = v23
		*(*int32)(unsafe.Add(mBase, uint32(l1))) = v22
		*(*int64)(unsafe.Add(mBase, uint32(l1+int32(12)))) = v23
		*(*int64)(unsafe.Add(mBase, uint32(l1+int32(20)))) = v23
		v37 = int32(0)
		*(*int32)(unsafe.Add(mBase, uint32(l1+int32(28)))) = v37
		if v17 == v37 {
		} else {
			v42 = l1 + int32(4)
			v44 = v17 + int32(16)
			v45 = int32(0)
			v52 = v45
			v53 = v23
			v56 = v45
			v57 = v45
			v58 = v45
			v59 = int64(0)
			for {
				v65 = v44 + v52<<(uint(int32(3))%32)
				v66 = *(*int32)(unsafe.Add(mBase, uint32(v65)))
				if v66 == int32(0) {
					v91 = v53
					v92 = v56
					v93 = v57
					v94 = v58
					v95 = v59
				} else {
					v70 = v58 + int32(1)
					*(*int32)(unsafe.Add(mBase, uint32(l1)+20)) = v70
					v72 = *(*int32)(unsafe.Add(mBase, uint32(v65)+4))
					if v58 != 0 {
						if base.Ui32(v57) <= base.Ui32(v72) {
							v77 = v57
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(l1)+8)) = v72
							v77 = v72
						}
						if base.Ui32(v72) <= base.Ui32(v56) {
							v80 = v56
							v81 = v77
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(l1)+12)) = v72
							v80 = v72
							v81 = v77
						}
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(l1)+8)) = v72
						*(*int32)(unsafe.Add(mBase, uint32(l1)+12)) = v72
						v80 = v72
						v81 = v72
					}
					v84 = base.I64_extend_i32_s(v66)
					if v53 == int64(0) {
						*(*int64)(unsafe.Add(mBase, uint32(l1)+24)) = v84
						v89 = v84
					} else {
						if v53 <= v84 {
							v89 = v53
						} else {
							*(*int64)(unsafe.Add(mBase, uint32(l1)+24)) = v84
							v89 = v84
						}
					}
					v91 = v89
					v92 = v80
					v93 = v81
					v94 = v70
					v95 = v59 + base.I64_extend_i32_u(v72)
				}
				v99 = v52 + int32(1)
				if v99 != int32(160) {
					v52 = v99
					v53 = v91
					v56 = v92
					v57 = v93
					v58 = v94
					v59 = v95
					continue
				} else {
					break
				}
				break
			}
			if v94 == int32(0) {
			} else {
				v105 = base.I64_div_u_s(v95, base.I64_extend_i32_u(v94))
				*(*uint32)(unsafe.Add(mBase, uint32(l1)+4)) = uint32(v105)
				v109 = F___time(m, int32(0))
				mBase = m.M
				v110 = *(*int64)(unsafe.Add(mBase, uint32(l1)+24))
				if v109 == v110 {
					v113 = int64(1)
				} else {
					v113 = v109 - v110
				}
				*(*int64)(unsafe.Add(mBase, uint32(l1)+24)) = v113
			}
			v119 = int32(0)
			v131 = int64(0)
			for {
				v134 = v44 + v119<<(uint(int32(3))%32)
				v135 = *(*int32)(unsafe.Add(mBase, uint32(v134)))
				if v135 == int32(0) {
					v148 = v131
				} else {
					v138 = *(*int32)(unsafe.Add(mBase, uint32(v42)))
					v139 = *(*int32)(unsafe.Add(mBase, uint32(v134)+4))
					if base.Ui32(v139) < base.Ui32(v138) {
						v143 = v138 - v139
					} else {
						v143 = v139 - v138
					}
					v148 = base.I64_extend_i32_u(v143) + v131
				}
				v150 = v134 + int32(8)
				v151 = *(*int32)(unsafe.Add(mBase, uint32(v150)))
				if v151 == int32(0) {
					v164 = v148
				} else {
					v154 = *(*int32)(unsafe.Add(mBase, uint32(v42)))
					v155 = *(*int32)(unsafe.Add(mBase, uint32(v150)+4))
					if base.Ui32(v155) < base.Ui32(v154) {
						v159 = v154 - v155
					} else {
						v159 = v155 - v154
					}
					v164 = base.I64_extend_i32_u(v159) + v148
				}
				v166 = v119 + int32(2)
				if v166 != int32(160) {
					v119 = v166
					v131 = v164
					continue
				} else {
					break
				}
				break
			}
			v169 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
			if v169 == int32(0) {
			} else {
				v173 = base.I64_div_u_s(v164, base.I64_extend_i32_u(v169))
				*(*uint32)(unsafe.Add(mBase, uint32(l1)+16)) = uint32(v173)
			}
		}
		return
	}
}
func F_append(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	v10 = F_emscripten_builtin_malloc(m, l2+int32(6))
	mBase = m.M
	if v10 != 0 {
		v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		*(*int32)(unsafe.Add(mBase, uint32(v13))) = v10
		v15 = int32(0)
		*(*int32)(unsafe.Add(mBase, uint32(v10))) = v15
		v18 = v10 + int32(4)
		v20 = l2 + int32(1)
		if v20 == v15 {
		} else {
			v23 = F__emscripten_memcpy_bulkmem(m, v18, l1, v20)
			mBase = m.M
		}
		if l2 == int32(0) {
		} else {
			if l3 == int32(0) {
			} else {
				v32 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1+l2+int32(-1)))))
				if v32 == int32(47) {
				} else {
					v36 = int32(47)
					*(*uint8)(unsafe.Add(mBase, uint32(v18+l2))) = uint8(v36)
					v39 = int32(0)
					*(*uint8)(unsafe.Add(mBase, uint32(v18+v20))) = uint8(v39)
				}
			}
		}
		*(*int32)(unsafe.Add(mBase, uint32(l0))) = v10
		return int32(0)
	} else {
		return int32(-1)
	}
}
func F_applyBind(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
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
	var v32 int32
	_ = v32
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v53 int32
	_ = v53
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v83 int32
	_ = v83
	var v91 int32
	_ = v91
	var v97 int32
	_ = v97
	v8 = F_listenerByType(m, int32(0))
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return int32(0)
	} else {
		v13 = F_listenerByType(m, int32(2))
		mBase = m.M
		v14 = m.ExcPending
		if v14 != 0 {
			return int32(0)
		} else {
			if v8 == int32(0) {
				F__serverAssert(m, int32(_a459), int32(_a392), int32(2730))
				mBase = m.M
				v91 = m.ExcPending
				if v91 != 0 {
					return int32(0)
				} else {
					F_abort(m)
					mBase = m.M
					base.Wasm_trap_unreachable()
					for {
					}
				}
			} else {
				v17 = int32(_a69)
				*(*int32)(unsafe.Add(mBase, uint32(v8)+68)) = int32(_a452)
				v22 = *(*int32)(unsafe.Add(mBase, _consts[192]))
				*(*int32)(unsafe.Add(mBase, uint32(v8)+72)) = v22
				v25 = *(*int32)(unsafe.Add(mBase, _consts[103]))
				*(*int32)(unsafe.Add(mBase, uint32(v8)+76)) = v25
				v28 = F_connectionByType(m, int32(0))
				mBase = m.M
				v29 = m.ExcPending
				if v29 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v8)+80)) = v28
					v31 = F_changeListener(m, v8)
					mBase = m.M
					v32 = m.ExcPending
					if v32 != 0 {
						return int32(0)
					} else {
						if v31 != int32(-1) {
							v41 = int32(1)
							v43 = *(*int32)(unsafe.Add(mBase, _consts[105]))
							if v43 == int32(0) {
								v83 = v41
								return v83
							} else {
								if v13 == int32(0) {
									F__serverAssert(m, int32(_a460), int32(_a392), int32(2742))
									mBase = m.M
									v97 = m.ExcPending
									if v97 != 0 {
										return int32(0)
									} else {
										F_abort(m)
										mBase = m.M
										base.Wasm_trap_unreachable()
										for {
										}
									}
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v13)+68)) = int32(_a452)
									v53 = *(*int32)(unsafe.Add(mBase, _consts[192]))
									*(*int32)(unsafe.Add(mBase, uint32(v13)+76)) = v43
									*(*int32)(unsafe.Add(mBase, uint32(v13)+72)) = v53
									v57 = F_connectionByType(m, int32(2))
									mBase = m.M
									v58 = m.ExcPending
									if v58 != 0 {
										return int32(0)
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v13)+80)) = v57
										v60 = F_changeListener(m, v13)
										mBase = m.M
										v61 = m.ExcPending
										if v61 != 0 {
											return int32(0)
										} else {
											if v60 != int32(-1) {
												v83 = v41
												return v83
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(_a461)
												v66 = *(*int32)(unsafe.Add(mBase, uint32(v8)+64))
												if v66 == int32(0) {
													v83 = int32(0)
													return v83
												} else {
													v69 = v8
													v73 = *(*int32)(unsafe.Add(mBase, uint32(v69)+80))
													v74 = *(*int32)(unsafe.Add(mBase, uint32(v73)+36))
													m.T0[v74].(func(*base.Module, int32))(m, v69)
													mBase = m.M
													v76 = m.ExcPending
													if v76 != 0 {
														return int32(0)
													} else {
														v83 = int32(0)
														return v83
													}
												}
											}
										}
									}
								}
							}
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(_a461)
							v37 = int32(0)
							if v13 == v37 {
								v83 = v37
								return v83
							} else {
								v40 = *(*int32)(unsafe.Add(mBase, uint32(v13)+64))
								if v40 != 0 {
									v69 = v13
									v73 = *(*int32)(unsafe.Add(mBase, uint32(v69)+80))
									v74 = *(*int32)(unsafe.Add(mBase, uint32(v73)+36))
									m.T0[v74].(func(*base.Module, int32))(m, v69)
									mBase = m.M
									v76 = m.ExcPending
									if v76 != 0 {
										return int32(0)
									} else {
										v83 = int32(0)
										return v83
									}
								} else {
									v83 = v37
									return v83
								}
							}
						}
					}
				}
			}
		}
	}
}
func F_applyRdmaBind(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v15 int32
	_ = v15
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
	var v30 int32
	_ = v30
	var v35 int32
	_ = v35
	v5 = F_listenerByType(m, int32(3))
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		return int32(0)
	} else {
		if v5 != 0 {
			v10 = int32(_a69)
			*(*int32)(unsafe.Add(mBase, uint32(v5)+68)) = int32(_a462)
			v15 = *(*int32)(unsafe.Add(mBase, _consts[205]))
			*(*int32)(unsafe.Add(mBase, uint32(v5)+72)) = v15
			v18 = *(*int32)(unsafe.Add(mBase, _consts[206]))
			*(*int32)(unsafe.Add(mBase, uint32(v5)+76)) = v18
			v21 = F_connectionByType(m, int32(3))
			mBase = m.M
			v22 = m.ExcPending
			if v22 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v5)+80)) = v21
				v25 = F_changeListener(m, v5)
				mBase = m.M
				v26 = m.ExcPending
				if v26 != 0 {
					return int32(0)
				} else {
					if v25 != int32(-1) {
						v35 = int32(1)
					} else {
						v30 = int32(_a463)
						*(*int32)(unsafe.Add(mBase, uint32(l0))) = v30
						v35 = int32(0)
					}
					return v35
				}
			}
		} else {
			v30 = int32(_a464)
			*(*int32)(unsafe.Add(mBase, uint32(l0))) = v30
			v35 = int32(0)
			return v35
		}
	}
}
func F_applyTLSPort(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
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
	var v50 int32
	_ = v50
	var v57 int32
	_ = v57
	v7 = F_connectionTypeTls(m)
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		return int32(0)
	} else {
		v11 = *(*int32)(unsafe.Add(mBase, uint32(v7)+12))
		v12 = m.T0[v11].(func(*base.Module, int32, int32) int32)(m, int32(_a450), int32(0))
		mBase = m.M
		v13 = m.ExcPending
		if v13 != 0 {
			return int32(0)
		} else {
			if v12 != int32(-1) {
				v21 = F_listenerByType(m, int32(2))
				mBase = m.M
				v22 = m.ExcPending
				if v22 != 0 {
					return int32(0)
				} else {
					if v21 == int32(0) {
						F__serverAssert(m, int32(_a451), int32(_a392), int32(2827))
						mBase = m.M
						v57 = m.ExcPending
						if v57 != 0 {
							return int32(0)
						} else {
							F_abort(m)
							mBase = m.M
							base.Wasm_trap_unreachable()
							for {
							}
						}
					} else {
						v25 = int32(_a69)
						*(*int32)(unsafe.Add(mBase, uint32(v21)+68)) = int32(_a452)
						v30 = *(*int32)(unsafe.Add(mBase, _consts[192]))
						*(*int32)(unsafe.Add(mBase, uint32(v21)+72)) = v30
						v33 = *(*int32)(unsafe.Add(mBase, _consts[105]))
						*(*int32)(unsafe.Add(mBase, uint32(v21)+76)) = v33
						v36 = F_connectionByType(m, int32(2))
						mBase = m.M
						v37 = m.ExcPending
						if v37 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v21)+80)) = v36
							F_clusterUpdateMyselfAnnouncedPorts(m)
							mBase = m.M
							v40 = m.ExcPending
							if v40 != 0 {
								return int32(0)
							} else {
								v41 = F_changeListener(m, v21)
								mBase = m.M
								v42 = m.ExcPending
								if v42 != 0 {
									return int32(0)
								} else {
									if v41 != int32(-1) {
										F_tlsResetCertInfo(m)
										mBase = m.M
										v50 = m.ExcPending
										if v50 != 0 {
											return int32(0)
										} else {
											return int32(1)
										}
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(_a453)
										return int32(0)
									}
								}
							}
						}
					}
				}
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(_a454)
				return int32(0)
			}
		}
	}
}
func F_applyWatchdogPeriod(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v12 int64
	_ = v12
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
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
	var v47 int32
	_ = v47
	v4 = m.G0
	v6 = v4 - int32(32)
	m.G0 = v6
	v9 = *(*int32)(unsafe.Add(mBase, _consts[241]))
	if v9 != 0 {
		v23 = *(*int32)(unsafe.Add(mBase, _consts[149]))
		v24 = base.I32_div_s(int32(1000), v23)
		v26 = v24 << (uint(int32(1)) % 32)
		if v26 <= v9 {
			v30 = v9
		} else {
			*(*int32)(unsafe.Add(mBase, _consts[241])) = v26
			v30 = v26
		}
		v31 = int32(0)
		*(*int32)(unsafe.Add(mBase, uint32(v6)+8)) = v31
		*(*int64)(unsafe.Add(mBase, uint32(v6))) = int64(0)
		v35 = int32(1000)
		v36 = base.I32_div_s(v30, v35)
		*(*int64)(unsafe.Add(mBase, uint32(v6)+16)) = base.I64_extend_i32_s(v36)
		*(*int32)(unsafe.Add(mBase, uint32(v6)+24)) = (v30 - v36*v35) * v35
		v47 = F_setitimer(m, v31, v6, v31)
		mBase = m.M
	} else {
		v10 = int32(0)
		*(*int32)(unsafe.Add(mBase, uint32(v6)+24)) = v10
		v12 = int64(0)
		*(*int64)(unsafe.Add(mBase, uint32(v6)+16)) = v12
		*(*int32)(unsafe.Add(mBase, uint32(v6)+8)) = v10
		*(*int64)(unsafe.Add(mBase, uint32(v6))) = v12
		v20 = F_setitimer(m, v10, v6, v10)
		mBase = m.M
	}
	m.G0 = v6 + int32(32)
	return
}
func F_auxresume(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v12 int32
	_ = v12
	var v20 int32
	_ = v20
	var v52 int32
	_ = v52
	var v56 int32
	_ = v56
	var v67 int32
	_ = v67
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	var v98 int32
	_ = v98
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v111 int32
	_ = v111
	var v118 int32
	_ = v118
	var v125 int32
	_ = v125
	var v131 int32
	_ = v131
	var v133 int32
	_ = v133
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v144 int64
	_ = v144
	var v146 int32
	_ = v146
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v153 int32
	_ = v153
	var v156 int64
	_ = v156
	var v160 int32
	_ = v160
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v165 int32
	_ = v165
	var v171 int32
	_ = v171
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v185 int32
	_ = v185
	var v186 int64
	_ = v186
	var v188 int32
	_ = v188
	var v199 int32
	_ = v199
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v209 int32
	_ = v209
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v227 int32
	_ = v227
	var v234 int32
	_ = v234
	var v241 int32
	_ = v241
	var v247 int32
	_ = v247
	var v249 int32
	_ = v249
	var v252 int32
	_ = v252
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v258 int32
	_ = v258
	var v259 int32
	_ = v259
	var v260 int64
	_ = v260
	var v262 int32
	_ = v262
	var v264 int32
	_ = v264
	var v265 int32
	_ = v265
	var v269 int32
	_ = v269
	var v272 int64
	_ = v272
	var v276 int32
	_ = v276
	var v278 int32
	_ = v278
	var v279 int32
	_ = v279
	var v281 int32
	_ = v281
	var v287 int32
	_ = v287
	var v294 int32
	_ = v294
	var v295 int32
	_ = v295
	var v301 int32
	_ = v301
	var v302 int64
	_ = v302
	var v304 int32
	_ = v304
	var v323 int32
	_ = v323
	var v390 int32
	_ = v390
	var v391 int32
	_ = v391
	var v397 int32
	_ = v397
	var v398 int64
	_ = v398
	var v400 int32
	_ = v400
	var v414 int32
	_ = v414
	v6 = m.G0
	v8 = v6 - int32(112)
	m.G0 = v8
	if l0 != l1 {
		v12 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+6)))
		switch v12 {
		case 0:
			v20 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
			v52 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
			if base.Ui32(v20) <= base.Ui32(v52) {
				v67 = int32(0)
			} else {
				v56 = base.I32_div_s(v20-v52, int32(24))
				*(*int32)(unsafe.Add(mBase, uint32(v8+int32(12))+96)) = v56
				v67 = int32(1)
			}
			if int32(0) < v67 {
				v79 = int32(2)
			} else {
				v72 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
				v73 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
				if (v72-v73)>>(uint(int32(4))%32) != 0 {
					v77 = int32(1)
				} else {
					v77 = int32(3)
				}
				v79 = v77
			}
		case 1:
			v79 = v12
		default:
			v79 = int32(3)
		}
	} else {
		v79 = int32(0)
	}
	v80 = F_lua_checkstack(m, l1, l2)
	mBase = m.M
	v83 = m.ExcPending
	if v83 != 0 {
		return int32(0)
	} else {
		if v80 != 0 {
			if v79 == int32(1) {
				if l0 == l1 {
				} else {
					v111 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
					*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v111 - l2<<(uint(int32(4))%32)
					if l2 < int32(1) {
					} else {
						v118 = int32(1)
						if l2 == v118 {
							v171 = int32(0)
						} else {
							v125 = int32(0)
							v131 = v125
							v133 = v125
							for {
								v136 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
								v137 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
								v138 = int32(16)
								*(*int32)(unsafe.Add(mBase, uint32(l1)+8)) = v137 + v138
								v142 = v131 << (uint(int32(4)) % 32)
								v143 = v136 + v142
								v144 = *(*int64)(unsafe.Add(mBase, uint32(v143)))
								*(*int64)(unsafe.Add(mBase, uint32(v137))) = v144
								v146 = *(*int32)(unsafe.Add(mBase, uint32(v143)+8))
								*(*int32)(unsafe.Add(mBase, uint32(v137)+8)) = v146
								v148 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
								v149 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
								*(*int32)(unsafe.Add(mBase, uint32(l1)+8)) = v149 + v138
								v153 = v148 + v142
								v156 = *(*int64)(unsafe.Add(mBase, uint32(v153+v138)))
								*(*int64)(unsafe.Add(mBase, uint32(v149))) = v156
								v160 = *(*int32)(unsafe.Add(mBase, uint32(v153+int32(24))))
								*(*int32)(unsafe.Add(mBase, uint32(v149)+8)) = v160
								v162 = int32(2)
								v163 = v131 + v162
								v165 = v133 + v162
								if v165 != l2&int32(2147483646) {
									v131 = v163
									v133 = v165
									continue
								} else {
									break
								}
								break
							}
							v171 = v163
						}
						if l2&v118 == int32(0) {
						} else {
							v178 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
							v179 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
							*(*int32)(unsafe.Add(mBase, uint32(l1)+8)) = v179 + int32(16)
							v185 = v178 + v171<<(uint(int32(4))%32)
							v186 = *(*int64)(unsafe.Add(mBase, uint32(v185)))
							*(*int64)(unsafe.Add(mBase, uint32(v179))) = v186
							v188 = *(*int32)(unsafe.Add(mBase, uint32(v185)+8))
							*(*int32)(unsafe.Add(mBase, uint32(v179)+8)) = v188
						}
					}
				}
				v199 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+52)))
				*(*uint16)(unsafe.Add(mBase, uint32(l1)+52)) = uint16(v199)
				v201 = F_lua_resume(m, l1, l2)
				mBase = m.M
				v202 = m.ExcPending
				if v202 != 0 {
					return int32(0)
				} else {
					if base.Ui32(int32(1)) < base.Ui32(v201) {
						if l1 == l0 {
						} else {
							v323 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
							*(*int32)(unsafe.Add(mBase, uint32(l1)+8)) = v323 - int32(16)
							v390 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
							v391 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
							*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v391 + int32(16)
							v397 = v390 + int32(0)
							v398 = *(*int64)(unsafe.Add(mBase, uint32(v397)))
							*(*int64)(unsafe.Add(mBase, uint32(v391))) = v398
							v400 = *(*int32)(unsafe.Add(mBase, uint32(v397)+8))
							*(*int32)(unsafe.Add(mBase, uint32(v391)+8)) = v400
						}
						v414 = int32(-1)
						m.G0 = v8 + int32(112)
						return v414
					} else {
						v205 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
						v206 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
						v209 = (v205 - v206) >> (uint(int32(4)) % 32)
						v212 = F_lua_checkstack(m, l0, v209+int32(1))
						mBase = m.M
						v213 = m.ExcPending
						if v213 != 0 {
							return int32(0)
						} else {
							if v212 != 0 {
								if l1 == l0 {
								} else {
									v227 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
									*(*int32)(unsafe.Add(mBase, uint32(l1)+8)) = v227 - v209<<(uint(int32(4))%32)
									if v209 < int32(1) {
									} else {
										v234 = int32(1)
										if v209 == v234 {
											v287 = int32(0)
										} else {
											v241 = int32(0)
											v247 = v241
											v249 = v241
											for {
												v252 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
												v253 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
												v254 = int32(16)
												*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v253 + v254
												v258 = v247 << (uint(int32(4)) % 32)
												v259 = v252 + v258
												v260 = *(*int64)(unsafe.Add(mBase, uint32(v259)))
												*(*int64)(unsafe.Add(mBase, uint32(v253))) = v260
												v262 = *(*int32)(unsafe.Add(mBase, uint32(v259)+8))
												*(*int32)(unsafe.Add(mBase, uint32(v253)+8)) = v262
												v264 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
												v265 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
												*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v265 + v254
												v269 = v264 + v258
												v272 = *(*int64)(unsafe.Add(mBase, uint32(v269+v254)))
												*(*int64)(unsafe.Add(mBase, uint32(v265))) = v272
												v276 = *(*int32)(unsafe.Add(mBase, uint32(v269+int32(24))))
												*(*int32)(unsafe.Add(mBase, uint32(v265)+8)) = v276
												v278 = int32(2)
												v279 = v247 + v278
												v281 = v249 + v278
												if v281 != v209&int32(2147483646) {
													v247 = v279
													v249 = v281
													continue
												} else {
													break
												}
												break
											}
											v287 = v279
										}
										if v209&v234 == int32(0) {
										} else {
											v294 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
											v295 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
											*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v295 + int32(16)
											v301 = v294 + v287<<(uint(int32(4))%32)
											v302 = *(*int64)(unsafe.Add(mBase, uint32(v301)))
											*(*int64)(unsafe.Add(mBase, uint32(v295))) = v302
											v304 = *(*int32)(unsafe.Add(mBase, uint32(v301)+8))
											*(*int32)(unsafe.Add(mBase, uint32(v295)+8)) = v304
										}
									}
								}
								v414 = v209
								m.G0 = v8 + int32(112)
								return v414
							} else {
								v214 = m.G3
								v218 = F_luaL_error(m, l0, v214+int32(_a2310), int32(0))
								mBase = m.M
								v219 = m.ExcPending
								if v219 != 0 {
									return int32(0)
								} else {
									if l1 == l0 {
									} else {
										v227 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
										*(*int32)(unsafe.Add(mBase, uint32(l1)+8)) = v227 - v209<<(uint(int32(4))%32)
										if v209 < int32(1) {
										} else {
											v234 = int32(1)
											if v209 == v234 {
												v287 = int32(0)
											} else {
												v241 = int32(0)
												v247 = v241
												v249 = v241
												for {
													v252 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
													v253 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
													v254 = int32(16)
													*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v253 + v254
													v258 = v247 << (uint(int32(4)) % 32)
													v259 = v252 + v258
													v260 = *(*int64)(unsafe.Add(mBase, uint32(v259)))
													*(*int64)(unsafe.Add(mBase, uint32(v253))) = v260
													v262 = *(*int32)(unsafe.Add(mBase, uint32(v259)+8))
													*(*int32)(unsafe.Add(mBase, uint32(v253)+8)) = v262
													v264 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
													v265 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
													*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v265 + v254
													v269 = v264 + v258
													v272 = *(*int64)(unsafe.Add(mBase, uint32(v269+v254)))
													*(*int64)(unsafe.Add(mBase, uint32(v265))) = v272
													v276 = *(*int32)(unsafe.Add(mBase, uint32(v269+int32(24))))
													*(*int32)(unsafe.Add(mBase, uint32(v265)+8)) = v276
													v278 = int32(2)
													v279 = v247 + v278
													v281 = v249 + v278
													if v281 != v209&int32(2147483646) {
														v247 = v279
														v249 = v281
														continue
													} else {
														break
													}
													break
												}
												v287 = v279
											}
											if v209&v234 == int32(0) {
											} else {
												v294 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
												v295 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
												*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v295 + int32(16)
												v301 = v294 + v287<<(uint(int32(4))%32)
												v302 = *(*int64)(unsafe.Add(mBase, uint32(v301)))
												*(*int64)(unsafe.Add(mBase, uint32(v295))) = v302
												v304 = *(*int32)(unsafe.Add(mBase, uint32(v301)+8))
												*(*int32)(unsafe.Add(mBase, uint32(v295)+8)) = v304
											}
										}
									}
									v414 = v209
									m.G0 = v8 + int32(112)
									return v414
								}
							}
						}
					}
				}
			} else {
				v92 = m.G3
				v98 = *(*int32)(unsafe.Add(mBase, uint32(v92+int32(_a2311)+v79<<(uint(int32(2))%32))))
				*(*int32)(unsafe.Add(mBase, uint32(v8))) = v98
				v102 = F_lua_pushfstring(m, l0, v92+int32(_a2312), v8)
				mBase = m.M
				v103 = m.ExcPending
				if v103 != 0 {
					return int32(0)
				} else {
					v414 = int32(-1)
					m.G0 = v8 + int32(112)
					return v414
				}
			}
		} else {
			v84 = m.G3
			v88 = F_luaL_error(m, l0, v84+int32(_a2313), int32(0))
			mBase = m.M
			v89 = m.ExcPending
			if v89 != 0 {
				return int32(0)
			} else {
				if v79 == int32(1) {
					if l0 == l1 {
					} else {
						v111 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
						*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v111 - l2<<(uint(int32(4))%32)
						if l2 < int32(1) {
						} else {
							v118 = int32(1)
							if l2 == v118 {
								v171 = int32(0)
							} else {
								v125 = int32(0)
								v131 = v125
								v133 = v125
								for {
									v136 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
									v137 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
									v138 = int32(16)
									*(*int32)(unsafe.Add(mBase, uint32(l1)+8)) = v137 + v138
									v142 = v131 << (uint(int32(4)) % 32)
									v143 = v136 + v142
									v144 = *(*int64)(unsafe.Add(mBase, uint32(v143)))
									*(*int64)(unsafe.Add(mBase, uint32(v137))) = v144
									v146 = *(*int32)(unsafe.Add(mBase, uint32(v143)+8))
									*(*int32)(unsafe.Add(mBase, uint32(v137)+8)) = v146
									v148 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
									v149 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
									*(*int32)(unsafe.Add(mBase, uint32(l1)+8)) = v149 + v138
									v153 = v148 + v142
									v156 = *(*int64)(unsafe.Add(mBase, uint32(v153+v138)))
									*(*int64)(unsafe.Add(mBase, uint32(v149))) = v156
									v160 = *(*int32)(unsafe.Add(mBase, uint32(v153+int32(24))))
									*(*int32)(unsafe.Add(mBase, uint32(v149)+8)) = v160
									v162 = int32(2)
									v163 = v131 + v162
									v165 = v133 + v162
									if v165 != l2&int32(2147483646) {
										v131 = v163
										v133 = v165
										continue
									} else {
										break
									}
									break
								}
								v171 = v163
							}
							if l2&v118 == int32(0) {
							} else {
								v178 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
								v179 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
								*(*int32)(unsafe.Add(mBase, uint32(l1)+8)) = v179 + int32(16)
								v185 = v178 + v171<<(uint(int32(4))%32)
								v186 = *(*int64)(unsafe.Add(mBase, uint32(v185)))
								*(*int64)(unsafe.Add(mBase, uint32(v179))) = v186
								v188 = *(*int32)(unsafe.Add(mBase, uint32(v185)+8))
								*(*int32)(unsafe.Add(mBase, uint32(v179)+8)) = v188
							}
						}
					}
					v199 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+52)))
					*(*uint16)(unsafe.Add(mBase, uint32(l1)+52)) = uint16(v199)
					v201 = F_lua_resume(m, l1, l2)
					mBase = m.M
					v202 = m.ExcPending
					if v202 != 0 {
						return int32(0)
					} else {
						if base.Ui32(int32(1)) < base.Ui32(v201) {
							if l1 == l0 {
							} else {
								v323 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
								*(*int32)(unsafe.Add(mBase, uint32(l1)+8)) = v323 - int32(16)
								v390 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
								v391 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
								*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v391 + int32(16)
								v397 = v390 + int32(0)
								v398 = *(*int64)(unsafe.Add(mBase, uint32(v397)))
								*(*int64)(unsafe.Add(mBase, uint32(v391))) = v398
								v400 = *(*int32)(unsafe.Add(mBase, uint32(v397)+8))
								*(*int32)(unsafe.Add(mBase, uint32(v391)+8)) = v400
							}
							v414 = int32(-1)
							m.G0 = v8 + int32(112)
							return v414
						} else {
							v205 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
							v206 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
							v209 = (v205 - v206) >> (uint(int32(4)) % 32)
							v212 = F_lua_checkstack(m, l0, v209+int32(1))
							mBase = m.M
							v213 = m.ExcPending
							if v213 != 0 {
								return int32(0)
							} else {
								if v212 != 0 {
									if l1 == l0 {
									} else {
										v227 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
										*(*int32)(unsafe.Add(mBase, uint32(l1)+8)) = v227 - v209<<(uint(int32(4))%32)
										if v209 < int32(1) {
										} else {
											v234 = int32(1)
											if v209 == v234 {
												v287 = int32(0)
											} else {
												v241 = int32(0)
												v247 = v241
												v249 = v241
												for {
													v252 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
													v253 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
													v254 = int32(16)
													*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v253 + v254
													v258 = v247 << (uint(int32(4)) % 32)
													v259 = v252 + v258
													v260 = *(*int64)(unsafe.Add(mBase, uint32(v259)))
													*(*int64)(unsafe.Add(mBase, uint32(v253))) = v260
													v262 = *(*int32)(unsafe.Add(mBase, uint32(v259)+8))
													*(*int32)(unsafe.Add(mBase, uint32(v253)+8)) = v262
													v264 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
													v265 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
													*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v265 + v254
													v269 = v264 + v258
													v272 = *(*int64)(unsafe.Add(mBase, uint32(v269+v254)))
													*(*int64)(unsafe.Add(mBase, uint32(v265))) = v272
													v276 = *(*int32)(unsafe.Add(mBase, uint32(v269+int32(24))))
													*(*int32)(unsafe.Add(mBase, uint32(v265)+8)) = v276
													v278 = int32(2)
													v279 = v247 + v278
													v281 = v249 + v278
													if v281 != v209&int32(2147483646) {
														v247 = v279
														v249 = v281
														continue
													} else {
														break
													}
													break
												}
												v287 = v279
											}
											if v209&v234 == int32(0) {
											} else {
												v294 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
												v295 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
												*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v295 + int32(16)
												v301 = v294 + v287<<(uint(int32(4))%32)
												v302 = *(*int64)(unsafe.Add(mBase, uint32(v301)))
												*(*int64)(unsafe.Add(mBase, uint32(v295))) = v302
												v304 = *(*int32)(unsafe.Add(mBase, uint32(v301)+8))
												*(*int32)(unsafe.Add(mBase, uint32(v295)+8)) = v304
											}
										}
									}
									v414 = v209
									m.G0 = v8 + int32(112)
									return v414
								} else {
									v214 = m.G3
									v218 = F_luaL_error(m, l0, v214+int32(_a2310), int32(0))
									mBase = m.M
									v219 = m.ExcPending
									if v219 != 0 {
										return int32(0)
									} else {
										if l1 == l0 {
										} else {
											v227 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
											*(*int32)(unsafe.Add(mBase, uint32(l1)+8)) = v227 - v209<<(uint(int32(4))%32)
											if v209 < int32(1) {
											} else {
												v234 = int32(1)
												if v209 == v234 {
													v287 = int32(0)
												} else {
													v241 = int32(0)
													v247 = v241
													v249 = v241
													for {
														v252 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
														v253 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
														v254 = int32(16)
														*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v253 + v254
														v258 = v247 << (uint(int32(4)) % 32)
														v259 = v252 + v258
														v260 = *(*int64)(unsafe.Add(mBase, uint32(v259)))
														*(*int64)(unsafe.Add(mBase, uint32(v253))) = v260
														v262 = *(*int32)(unsafe.Add(mBase, uint32(v259)+8))
														*(*int32)(unsafe.Add(mBase, uint32(v253)+8)) = v262
														v264 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
														v265 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
														*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v265 + v254
														v269 = v264 + v258
														v272 = *(*int64)(unsafe.Add(mBase, uint32(v269+v254)))
														*(*int64)(unsafe.Add(mBase, uint32(v265))) = v272
														v276 = *(*int32)(unsafe.Add(mBase, uint32(v269+int32(24))))
														*(*int32)(unsafe.Add(mBase, uint32(v265)+8)) = v276
														v278 = int32(2)
														v279 = v247 + v278
														v281 = v249 + v278
														if v281 != v209&int32(2147483646) {
															v247 = v279
															v249 = v281
															continue
														} else {
															break
														}
														break
													}
													v287 = v279
												}
												if v209&v234 == int32(0) {
												} else {
													v294 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
													v295 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
													*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v295 + int32(16)
													v301 = v294 + v287<<(uint(int32(4))%32)
													v302 = *(*int64)(unsafe.Add(mBase, uint32(v301)))
													*(*int64)(unsafe.Add(mBase, uint32(v295))) = v302
													v304 = *(*int32)(unsafe.Add(mBase, uint32(v301)+8))
													*(*int32)(unsafe.Add(mBase, uint32(v295)+8)) = v304
												}
											}
										}
										v414 = v209
										m.G0 = v8 + int32(112)
										return v414
									}
								}
							}
						}
					}
				} else {
					v92 = m.G3
					v98 = *(*int32)(unsafe.Add(mBase, uint32(v92+int32(_a2311)+v79<<(uint(int32(2))%32))))
					*(*int32)(unsafe.Add(mBase, uint32(v8))) = v98
					v102 = F_lua_pushfstring(m, l0, v92+int32(_a2312), v8)
					mBase = m.M
					v103 = m.ExcPending
					if v103 != 0 {
						return int32(0)
					} else {
						v414 = int32(-1)
						m.G0 = v8 + int32(112)
						return v414
					}
				}
			}
		}
	}
}
func F_auxsort(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v81 int64
	_ = v81
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v148 int64
	_ = v148
	var v150 int32
	_ = v150
	var v152 int32
	_ = v152
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v164 int32
	_ = v164
	var v167 int32
	_ = v167
	var v186 int32
	_ = v186
	var v196 int32
	_ = v196
	var v202 int32
	_ = v202
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v216 int32
	_ = v216
	var v258 int32
	_ = v258
	var v259 int32
	_ = v259
	var v260 int32
	_ = v260
	var v261 int64
	_ = v261
	var v263 int32
	_ = v263
	var v265 int32
	_ = v265
	var v274 int32
	_ = v274
	var v275 int32
	_ = v275
	var v280 int32
	_ = v280
	var v281 int32
	_ = v281
	var v283 int32
	_ = v283
	var v325 int32
	_ = v325
	var v326 int32
	_ = v326
	var v327 int32
	_ = v327
	var v328 int64
	_ = v328
	var v330 int32
	_ = v330
	var v332 int32
	_ = v332
	var v338 int32
	_ = v338
	var v339 int32
	_ = v339
	var v344 int32
	_ = v344
	var v347 int32
	_ = v347
	var v366 int32
	_ = v366
	var v381 int32
	_ = v381
	var v382 int32
	_ = v382
	var v387 int32
	_ = v387
	var v388 int32
	_ = v388
	var v390 int32
	_ = v390
	var v432 int32
	_ = v432
	var v433 int32
	_ = v433
	var v434 int32
	_ = v434
	var v435 int64
	_ = v435
	var v437 int32
	_ = v437
	var v439 int32
	_ = v439
	var v445 int32
	_ = v445
	var v446 int32
	_ = v446
	var v451 int32
	_ = v451
	var v454 int32
	_ = v454
	var v473 int32
	_ = v473
	var v490 int32
	_ = v490
	var v491 int32
	_ = v491
	var v496 int32
	_ = v496
	var v497 int32
	_ = v497
	var v499 int32
	_ = v499
	var v541 int32
	_ = v541
	var v542 int32
	_ = v542
	var v543 int32
	_ = v543
	var v544 int64
	_ = v544
	var v546 int32
	_ = v546
	var v548 int32
	_ = v548
	var v569 int32
	_ = v569
	var v572 int32
	_ = v572
	var v608 int32
	_ = v608
	var v609 int64
	_ = v609
	var v611 int32
	_ = v611
	var v613 int32
	_ = v613
	var v619 int32
	_ = v619
	var v624 int32
	_ = v624
	var v625 int32
	_ = v625
	var v630 int32
	_ = v630
	var v631 int32
	_ = v631
	var v633 int32
	_ = v633
	var v675 int32
	_ = v675
	var v676 int32
	_ = v676
	var v677 int32
	_ = v677
	var v678 int64
	_ = v678
	var v680 int32
	_ = v680
	var v682 int32
	_ = v682
	var v689 int32
	_ = v689
	var v690 int32
	_ = v690
	var v692 int32
	_ = v692
	var v698 int32
	_ = v698
	var v701 int32
	_ = v701
	var v704 int32
	_ = v704
	var v709 int32
	_ = v709
	var v710 int32
	_ = v710
	var v715 int32
	_ = v715
	var v716 int32
	_ = v716
	var v718 int32
	_ = v718
	var v760 int32
	_ = v760
	var v761 int32
	_ = v761
	var v762 int32
	_ = v762
	var v763 int64
	_ = v763
	var v765 int32
	_ = v765
	var v767 int32
	_ = v767
	var v773 int32
	_ = v773
	var v774 int32
	_ = v774
	var v780 int32
	_ = v780
	var __phi780 int32
	_ = __phi780
	var v781 int32
	_ = v781
	var __phi781 int32
	_ = __phi781
	var v788 int32
	_ = v788
	var v792 int32
	_ = v792
	var v793 int32
	_ = v793
	var v812 int32
	_ = v812
	var v824 int32
	_ = v824
	var v829 int32
	_ = v829
	var v830 int32
	_ = v830
	var v835 int32
	_ = v835
	var v836 int32
	_ = v836
	var v838 int32
	_ = v838
	var v880 int32
	_ = v880
	var v881 int32
	_ = v881
	var v882 int32
	_ = v882
	var v883 int64
	_ = v883
	var v885 int32
	_ = v885
	var v887 int32
	_ = v887
	var v893 int32
	_ = v893
	var v894 int32
	_ = v894
	var v898 int32
	_ = v898
	var v899 int32
	_ = v899
	var v907 int32
	_ = v907
	var v912 int32
	_ = v912
	var v913 int32
	_ = v913
	var v918 int32
	_ = v918
	var v919 int32
	_ = v919
	var v921 int32
	_ = v921
	var v963 int32
	_ = v963
	var v964 int32
	_ = v964
	var v965 int32
	_ = v965
	var v966 int64
	_ = v966
	var v968 int32
	_ = v968
	var v970 int32
	_ = v970
	var v976 int32
	_ = v976
	var v977 int32
	_ = v977
	var v986 int32
	_ = v986
	var __phi986 int32
	_ = __phi986
	var v987 int32
	_ = v987
	var __phi987 int32
	_ = __phi987
	var v991 int32
	_ = v991
	var v995 int32
	_ = v995
	var v996 int32
	_ = v996
	var v1015 int32
	_ = v1015
	var v1027 int32
	_ = v1027
	var v1032 int32
	_ = v1032
	var v1033 int32
	_ = v1033
	var v1038 int32
	_ = v1038
	var v1039 int32
	_ = v1039
	var v1041 int32
	_ = v1041
	var v1083 int32
	_ = v1083
	var v1084 int32
	_ = v1084
	var v1085 int32
	_ = v1085
	var v1086 int64
	_ = v1086
	var v1088 int32
	_ = v1088
	var v1090 int32
	_ = v1090
	var v1096 int32
	_ = v1096
	var v1097 int32
	_ = v1097
	var v1104 int32
	_ = v1104
	var v1105 int32
	_ = v1105
	var v1127 int32
	_ = v1127
	var v1142 int32
	_ = v1142
	var v1143 int32
	_ = v1143
	var v1148 int32
	_ = v1148
	var v1149 int32
	_ = v1149
	var v1151 int32
	_ = v1151
	var v1193 int32
	_ = v1193
	var v1194 int32
	_ = v1194
	var v1195 int32
	_ = v1195
	var v1196 int64
	_ = v1196
	var v1198 int32
	_ = v1198
	var v1200 int32
	_ = v1200
	var v1209 int32
	_ = v1209
	var v1210 int32
	_ = v1210
	var v1215 int32
	_ = v1215
	var v1216 int32
	_ = v1216
	var v1218 int32
	_ = v1218
	var v1260 int32
	_ = v1260
	var v1261 int32
	_ = v1261
	var v1262 int32
	_ = v1262
	var v1263 int64
	_ = v1263
	var v1265 int32
	_ = v1265
	var v1267 int32
	_ = v1267
	var v1273 int32
	_ = v1273
	var v1276 int32
	_ = v1276
	var v1278 int32
	_ = v1278
	var v1281 int32
	_ = v1281
	var v1282 int32
	_ = v1282
	var v1283 int32
	_ = v1283
	var v1285 int32
	_ = v1285
	var v1286 int32
	_ = v1286
	var v1287 int32
	_ = v1287
	if l2 <= l1 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v13 = l1
	v14 = l2
	goto L3
L3:
	;
	goto L8
L4:
	;
	goto L1
L5:
	;
	goto L24
L6:
	;
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v36)))
	v79 = F_luaH_getnum(m, v78, v13)
	mBase = m.M
	v80 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v81 = *(*int64)(unsafe.Add(mBase, uint32(v79)))
	*(*int64)(unsafe.Add(mBase, uint32(v80))) = v81
	v83 = *(*int32)(unsafe.Add(mBase, uint32(v79)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v80)+8)) = v83
	v85 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v85 + int32(16)
	goto L5
L8:
	;
	v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v33 = v28 + int32(0)
	v34 = m.G398
	if base.Ui32(v33) < base.Ui32(v27) {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	v36 = v33
	goto L11
L10:
	;
	v36 = v34
	goto L11
L11:
	;
	goto L6
L21:
	;
	v158 = F_sort_comp(m, l0, int32(-1), int32(-2))
	mBase = m.M
	v159 = m.ExcPending
	if v159 != 0 {
		goto L39
	} else {
		goto L40
	}
L22:
	;
	v145 = *(*int32)(unsafe.Add(mBase, uint32(v103)))
	v146 = F_luaH_getnum(m, v145, v14)
	mBase = m.M
	v147 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v148 = *(*int64)(unsafe.Add(mBase, uint32(v146)))
	*(*int64)(unsafe.Add(mBase, uint32(v147))) = v148
	v150 = *(*int32)(unsafe.Add(mBase, uint32(v146)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v147)+8)) = v150
	v152 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v152 + int32(16)
	goto L21
L24:
	;
	v94 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v95 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v100 = v95 + int32(0)
	v101 = m.G398
	if base.Ui32(v100) < base.Ui32(v94) {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	v103 = v100
	goto L27
L26:
	;
	v103 = v101
	goto L27
L27:
	;
	goto L22
L37:
	;
	v196 = v14 - v13
	if v196 == int32(1) {
		goto L1
	} else {
		goto L52
	}
L38:
	;
	goto L46
L39:
	;
	return
L40:
	;
	if v158 == int32(0) {
		goto L38
	} else {
		goto L41
	}
L41:
	;
	F_lua_rawseti(m, l0, int32(1), v13)
	mBase = m.M
	v164 = m.ExcPending
	if v164 != 0 {
		goto L39
	} else {
		goto L42
	}
L42:
	;
	F_lua_rawseti(m, l0, int32(1), v14)
	mBase = m.M
	v167 = m.ExcPending
	if v167 != 0 {
		goto L39
	} else {
		goto L43
	}
L43:
	;
	goto L37
L44:
	;
	goto L37
L45:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v186 + int32(-32)
	goto L44
L46:
	;
	v186 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	goto L45
L52:
	;
	v202 = base.I32_div_s(v14+v13, int32(2))
	goto L56
L53:
	;
	goto L72
L54:
	;
	v258 = *(*int32)(unsafe.Add(mBase, uint32(v216)))
	v259 = F_luaH_getnum(m, v258, v202)
	mBase = m.M
	v260 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v261 = *(*int64)(unsafe.Add(mBase, uint32(v259)))
	*(*int64)(unsafe.Add(mBase, uint32(v260))) = v261
	v263 = *(*int32)(unsafe.Add(mBase, uint32(v259)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v260)+8)) = v263
	v265 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v265 + int32(16)
	goto L53
L56:
	;
	v207 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v208 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v213 = v208 + int32(0)
	v214 = m.G398
	if base.Ui32(v213) < base.Ui32(v207) {
		goto L57
	} else {
		goto L58
	}
L57:
	;
	v216 = v213
	goto L59
L58:
	;
	v216 = v214
	goto L59
L59:
	;
	goto L54
L69:
	;
	v338 = F_sort_comp(m, l0, int32(-2), int32(-1))
	mBase = m.M
	v339 = m.ExcPending
	if v339 != 0 {
		goto L39
	} else {
		goto L87
	}
L70:
	;
	v325 = *(*int32)(unsafe.Add(mBase, uint32(v283)))
	v326 = F_luaH_getnum(m, v325, v13)
	mBase = m.M
	v327 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v328 = *(*int64)(unsafe.Add(mBase, uint32(v326)))
	*(*int64)(unsafe.Add(mBase, uint32(v327))) = v328
	v330 = *(*int32)(unsafe.Add(mBase, uint32(v326)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v327)+8)) = v330
	v332 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v332 + int32(16)
	goto L69
L72:
	;
	v274 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v275 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v280 = v275 + int32(0)
	v281 = m.G398
	if base.Ui32(v280) < base.Ui32(v274) {
		goto L73
	} else {
		goto L74
	}
L73:
	;
	v283 = v280
	goto L75
L74:
	;
	v283 = v281
	goto L75
L75:
	;
	goto L70
L85:
	;
	if v196 == int32(2) {
		goto L1
	} else {
		goto L128
	}
L86:
	;
	goto L93
L87:
	;
	if v338 == int32(0) {
		goto L86
	} else {
		goto L88
	}
L88:
	;
	F_lua_rawseti(m, l0, int32(1), v202)
	mBase = m.M
	v344 = m.ExcPending
	if v344 != 0 {
		goto L39
	} else {
		goto L89
	}
L89:
	;
	F_lua_rawseti(m, l0, int32(1), v13)
	mBase = m.M
	v347 = m.ExcPending
	if v347 != 0 {
		goto L39
	} else {
		goto L90
	}
L90:
	;
	goto L85
L91:
	;
	goto L102
L92:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v366 + int32(-16)
	goto L91
L93:
	;
	v366 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	goto L92
L99:
	;
	v445 = F_sort_comp(m, l0, int32(-1), int32(-2))
	mBase = m.M
	v446 = m.ExcPending
	if v446 != 0 {
		goto L39
	} else {
		goto L116
	}
L100:
	;
	v432 = *(*int32)(unsafe.Add(mBase, uint32(v390)))
	v433 = F_luaH_getnum(m, v432, v14)
	mBase = m.M
	v434 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v435 = *(*int64)(unsafe.Add(mBase, uint32(v433)))
	*(*int64)(unsafe.Add(mBase, uint32(v434))) = v435
	v437 = *(*int32)(unsafe.Add(mBase, uint32(v433)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v434)+8)) = v437
	v439 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v439 + int32(16)
	goto L99
L102:
	;
	v381 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v382 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v387 = v382 + int32(0)
	v388 = m.G398
	if base.Ui32(v387) < base.Ui32(v381) {
		goto L103
	} else {
		goto L104
	}
L103:
	;
	v390 = v387
	goto L105
L104:
	;
	v390 = v388
	goto L105
L105:
	;
	goto L100
L115:
	;
	goto L122
L116:
	;
	if v445 == int32(0) {
		goto L115
	} else {
		goto L117
	}
L117:
	;
	F_lua_rawseti(m, l0, int32(1), v202)
	mBase = m.M
	v451 = m.ExcPending
	if v451 != 0 {
		goto L39
	} else {
		goto L118
	}
L118:
	;
	F_lua_rawseti(m, l0, int32(1), v14)
	mBase = m.M
	v454 = m.ExcPending
	if v454 != 0 {
		goto L39
	} else {
		goto L119
	}
L119:
	;
	goto L85
L120:
	;
	goto L85
L121:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v473 + int32(-32)
	goto L120
L122:
	;
	v473 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	goto L121
L128:
	;
	goto L132
L129:
	;
	goto L147
L130:
	;
	v541 = *(*int32)(unsafe.Add(mBase, uint32(v499)))
	v542 = F_luaH_getnum(m, v541, v202)
	mBase = m.M
	v543 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v544 = *(*int64)(unsafe.Add(mBase, uint32(v542)))
	*(*int64)(unsafe.Add(mBase, uint32(v543))) = v544
	v546 = *(*int32)(unsafe.Add(mBase, uint32(v542)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v543)+8)) = v546
	v548 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v548 + int32(16)
	goto L129
L132:
	;
	v490 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v491 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v496 = v491 + int32(0)
	v497 = m.G398
	if base.Ui32(v496) < base.Ui32(v490) {
		goto L133
	} else {
		goto L134
	}
L133:
	;
	v499 = v496
	goto L135
L134:
	;
	v499 = v497
	goto L135
L135:
	;
	goto L130
L145:
	;
	v619 = v14 + int32(-1)
	goto L164
L146:
	;
	v608 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v609 = *(*int64)(unsafe.Add(mBase, uint32(v572)))
	*(*int64)(unsafe.Add(mBase, uint32(v608))) = v609
	v611 = *(*int32)(unsafe.Add(mBase, uint32(v572)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v608)+8)) = v611
	v613 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v613 + int32(16)
	goto L145
L147:
	;
	goto L153
L153:
	;
	v569 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v572 = v569 + int32(-16)
	goto L146
L161:
	;
	v689 = v13
	v690 = v202
	v692 = v619
	goto L177
L162:
	;
	v675 = *(*int32)(unsafe.Add(mBase, uint32(v633)))
	v676 = F_luaH_getnum(m, v675, v619)
	mBase = m.M
	v677 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v678 = *(*int64)(unsafe.Add(mBase, uint32(v676)))
	*(*int64)(unsafe.Add(mBase, uint32(v677))) = v678
	v680 = *(*int32)(unsafe.Add(mBase, uint32(v676)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v677)+8)) = v680
	v682 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v682 + int32(16)
	goto L161
L164:
	;
	v624 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v625 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v630 = v625 + int32(0)
	v631 = m.G398
	if base.Ui32(v630) < base.Ui32(v624) {
		goto L165
	} else {
		goto L166
	}
L165:
	;
	v633 = v630
	goto L167
L166:
	;
	v633 = v631
	goto L167
L167:
	;
	goto L162
L177:
	;
	F_lua_rawseti(m, l0, int32(1), v690)
	mBase = m.M
	v698 = m.ExcPending
	if v698 != 0 {
		goto L39
	} else {
		goto L179
	}
L178:
	;
	goto L284
L179:
	;
	F_lua_rawseti(m, l0, int32(1), v692)
	mBase = m.M
	v701 = m.ExcPending
	if v701 != 0 {
		goto L39
	} else {
		goto L180
	}
L180:
	;
	v704 = v689 + int32(1)
	goto L184
L181:
	;
	v773 = F_sort_comp(m, l0, int32(-1), int32(-2))
	mBase = m.M
	v774 = m.ExcPending
	if v774 != 0 {
		goto L39
	} else {
		goto L198
	}
L182:
	;
	v760 = *(*int32)(unsafe.Add(mBase, uint32(v718)))
	v761 = F_luaH_getnum(m, v760, v704)
	mBase = m.M
	v762 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v763 = *(*int64)(unsafe.Add(mBase, uint32(v761)))
	*(*int64)(unsafe.Add(mBase, uint32(v762))) = v763
	v765 = *(*int32)(unsafe.Add(mBase, uint32(v761)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v762)+8)) = v765
	v767 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v767 + int32(16)
	goto L181
L184:
	;
	v709 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v710 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v715 = v710 + int32(0)
	v716 = m.G398
	if base.Ui32(v715) < base.Ui32(v709) {
		goto L185
	} else {
		goto L186
	}
L185:
	;
	v718 = v715
	goto L187
L186:
	;
	v718 = v716
	goto L187
L187:
	;
	goto L182
L197:
	;
	v907 = v692 + int32(-1)
	goto L234
L198:
	;
	if v773 == int32(0) {
		v898 = v689
		v899 = v704
		goto L197
	} else {
		goto L199
	}
L199:
	;
	__phi780 = v689
	__phi781 = v704
	v780 = __phi780
	v781 = __phi781
	goto L200
L200:
	;
	if v780 < v14 {
		goto L202
	} else {
		goto L203
	}
L201:
	;
	v898 = v781
	v899 = v824
	goto L197
L202:
	;
	goto L207
L203:
	;
	v788 = m.G3
	v792 = F_luaL_error(m, l0, v788+int32(_a2321), int32(0))
	mBase = m.M
	v793 = m.ExcPending
	if v793 != 0 {
		goto L39
	} else {
		goto L204
	}
L204:
	;
	goto L202
L205:
	;
	v824 = v781 + int32(1)
	goto L216
L206:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v812 + int32(-16)
	goto L205
L207:
	;
	v812 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	goto L206
L213:
	;
	v893 = F_sort_comp(m, l0, int32(-1), int32(-2))
	mBase = m.M
	v894 = m.ExcPending
	if v894 != 0 {
		goto L39
	} else {
		goto L229
	}
L214:
	;
	v880 = *(*int32)(unsafe.Add(mBase, uint32(v838)))
	v881 = F_luaH_getnum(m, v880, v824)
	mBase = m.M
	v882 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v883 = *(*int64)(unsafe.Add(mBase, uint32(v881)))
	*(*int64)(unsafe.Add(mBase, uint32(v882))) = v883
	v885 = *(*int32)(unsafe.Add(mBase, uint32(v881)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v882)+8)) = v885
	v887 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v887 + int32(16)
	goto L213
L216:
	;
	v829 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v830 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v835 = v830 + int32(0)
	v836 = m.G398
	if base.Ui32(v835) < base.Ui32(v829) {
		goto L217
	} else {
		goto L218
	}
L217:
	;
	v838 = v835
	goto L219
L218:
	;
	v838 = v836
	goto L219
L219:
	;
	goto L214
L229:
	;
	if v893 != 0 {
		__phi780 = v781
		__phi781 = v824
		v780 = __phi780
		v781 = __phi781
		goto L200
	} else {
		goto L230
	}
L230:
	;
	goto L201
L231:
	;
	v976 = F_sort_comp(m, l0, int32(-3), int32(-1))
	mBase = m.M
	v977 = m.ExcPending
	if v977 != 0 {
		goto L39
	} else {
		goto L248
	}
L232:
	;
	v963 = *(*int32)(unsafe.Add(mBase, uint32(v921)))
	v964 = F_luaH_getnum(m, v963, v907)
	mBase = m.M
	v965 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v966 = *(*int64)(unsafe.Add(mBase, uint32(v964)))
	*(*int64)(unsafe.Add(mBase, uint32(v965))) = v966
	v968 = *(*int32)(unsafe.Add(mBase, uint32(v964)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v965)+8)) = v968
	v970 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v970 + int32(16)
	goto L231
L234:
	;
	v912 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v913 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v918 = v913 + int32(0)
	v919 = m.G398
	if base.Ui32(v918) < base.Ui32(v912) {
		goto L235
	} else {
		goto L236
	}
L235:
	;
	v921 = v918
	goto L237
L236:
	;
	v921 = v919
	goto L237
L237:
	;
	goto L232
L247:
	;
	if v899 < v1104 {
		v689 = v899
		v690 = v899
		v692 = v1105
		goto L177
	} else {
		goto L281
	}
L248:
	;
	if v976 == int32(0) {
		v1104 = v692
		v1105 = v907
		goto L247
	} else {
		goto L249
	}
L249:
	;
	__phi986 = v692
	__phi987 = v907
	v986 = __phi986
	v987 = __phi987
	goto L250
L250:
	;
	if v13 < v986 {
		goto L252
	} else {
		goto L253
	}
L251:
	;
	v1104 = v987
	v1105 = v1027
	goto L247
L252:
	;
	goto L257
L253:
	;
	v991 = m.G3
	v995 = F_luaL_error(m, l0, v991+int32(_a2321), int32(0))
	mBase = m.M
	v996 = m.ExcPending
	if v996 != 0 {
		goto L39
	} else {
		goto L254
	}
L254:
	;
	goto L252
L255:
	;
	v1027 = v987 + int32(-1)
	goto L266
L256:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v1015 + int32(-16)
	goto L255
L257:
	;
	v1015 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	goto L256
L263:
	;
	v1096 = F_sort_comp(m, l0, int32(-3), int32(-1))
	mBase = m.M
	v1097 = m.ExcPending
	if v1097 != 0 {
		goto L39
	} else {
		goto L279
	}
L264:
	;
	v1083 = *(*int32)(unsafe.Add(mBase, uint32(v1041)))
	v1084 = F_luaH_getnum(m, v1083, v1027)
	mBase = m.M
	v1085 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1086 = *(*int64)(unsafe.Add(mBase, uint32(v1084)))
	*(*int64)(unsafe.Add(mBase, uint32(v1085))) = v1086
	v1088 = *(*int32)(unsafe.Add(mBase, uint32(v1084)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v1085)+8)) = v1088
	v1090 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v1090 + int32(16)
	goto L263
L266:
	;
	v1032 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1033 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v1038 = v1033 + int32(0)
	v1039 = m.G398
	if base.Ui32(v1038) < base.Ui32(v1032) {
		goto L267
	} else {
		goto L268
	}
L267:
	;
	v1041 = v1038
	goto L269
L268:
	;
	v1041 = v1039
	goto L269
L269:
	;
	goto L264
L279:
	;
	if v1096 != 0 {
		__phi986 = v987
		__phi987 = v1027
		v986 = __phi986
		v987 = __phi987
		goto L250
	} else {
		goto L280
	}
L280:
	;
	goto L251
L281:
	;
	goto L178
L282:
	;
	goto L293
L283:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v1127 + int32(-48)
	goto L282
L284:
	;
	v1127 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	goto L283
L290:
	;
	goto L309
L291:
	;
	v1193 = *(*int32)(unsafe.Add(mBase, uint32(v1151)))
	v1194 = F_luaH_getnum(m, v1193, v619)
	mBase = m.M
	v1195 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1196 = *(*int64)(unsafe.Add(mBase, uint32(v1194)))
	*(*int64)(unsafe.Add(mBase, uint32(v1195))) = v1196
	v1198 = *(*int32)(unsafe.Add(mBase, uint32(v1194)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v1195)+8)) = v1198
	v1200 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v1200 + int32(16)
	goto L290
L293:
	;
	v1142 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1143 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v1148 = v1143 + int32(0)
	v1149 = m.G398
	if base.Ui32(v1148) < base.Ui32(v1142) {
		goto L294
	} else {
		goto L295
	}
L294:
	;
	v1151 = v1148
	goto L296
L295:
	;
	v1151 = v1149
	goto L296
L296:
	;
	goto L291
L306:
	;
	F_lua_rawseti(m, l0, int32(1), v619)
	mBase = m.M
	v1273 = m.ExcPending
	if v1273 != 0 {
		goto L39
	} else {
		goto L322
	}
L307:
	;
	v1260 = *(*int32)(unsafe.Add(mBase, uint32(v1218)))
	v1261 = F_luaH_getnum(m, v1260, v899)
	mBase = m.M
	v1262 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1263 = *(*int64)(unsafe.Add(mBase, uint32(v1261)))
	*(*int64)(unsafe.Add(mBase, uint32(v1262))) = v1263
	v1265 = *(*int32)(unsafe.Add(mBase, uint32(v1261)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v1262)+8)) = v1265
	v1267 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v1267 + int32(16)
	goto L306
L309:
	;
	v1209 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1210 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v1215 = v1210 + int32(0)
	v1216 = m.G398
	if base.Ui32(v1215) < base.Ui32(v1209) {
		goto L310
	} else {
		goto L311
	}
L310:
	;
	v1218 = v1215
	goto L312
L311:
	;
	v1218 = v1216
	goto L312
L312:
	;
	goto L307
L322:
	;
	F_lua_rawseti(m, l0, int32(1), v899)
	mBase = m.M
	v1276 = m.ExcPending
	if v1276 != 0 {
		goto L39
	} else {
		goto L323
	}
L323:
	;
	v1278 = v898 + int32(2)
	v1281 = base.B2i32(v899-v13 < v14-v899)
	if v899-v13 < v14-v899 {
		goto L324
	} else {
		goto L325
	}
L324:
	;
	v1282 = v13
	goto L326
L325:
	;
	v1282 = v1278
	goto L326
L326:
	;
	if v899-v13 < v14-v899 {
		goto L327
	} else {
		goto L328
	}
L327:
	;
	v1283 = v898
	goto L329
L328:
	;
	v1283 = v14
	goto L329
L329:
	;
	F_auxsort(m, l0, v1282, v1283)
	mBase = m.M
	v1285 = m.ExcPending
	if v1285 != 0 {
		goto L39
	} else {
		goto L330
	}
L330:
	;
	if v899-v13 < v14-v899 {
		goto L331
	} else {
		goto L332
	}
L331:
	;
	v1286 = v1278
	goto L333
L332:
	;
	v1286 = v13
	goto L333
L333:
	;
	if v899-v13 < v14-v899 {
		goto L334
	} else {
		goto L335
	}
L334:
	;
	v1287 = v14
	goto L336
L335:
	;
	v1287 = v898
	goto L336
L336:
	;
	if v1286 < v1287 {
		v13 = v1286
		v14 = v1287
		goto L3
	} else {
		goto L337
	}
L337:
	;
	goto L4
}
