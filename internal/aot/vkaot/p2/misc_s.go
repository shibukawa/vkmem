package p2

import (
	base "github.com/shibukawa/vkmem/internal/aot/vkaot/base"
	"math"
	"unsafe"
)

func F___shgetc(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int64
	_ = v7
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v12 int64
	_ = v12
	var v13 int64
	_ = v13
	var v17 int32
	_ = v17
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
	var v37 int64
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int64
	_ = v40
	var v43 int64
	_ = v43
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	v7 = *(*int64)(unsafe.Add(mBase, uint32(l0)+120))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	v12 = v7 + base.I64_extend_i32_s(v8-v9)
	v13 = *(*int64)(unsafe.Add(mBase, uint32(l0)+112))
	if v13 == int64(0) {
		v17 = F___uflow(m, l0)
		mBase = m.M
		v20 = m.ExcPending
		if v20 != 0 {
			return int32(0)
		} else {
			if int32(-1) < v17 {
				v37 = v12 + int64(1)
				v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
				v39 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				v40 = *(*int64)(unsafe.Add(mBase, uint32(l0)+112))
				if v40 == int64(0) {
					v50 = v39
				} else {
					v43 = v40 - v37
					if base.I64_extend_i32_s(v39-v38) <= v43 {
						v50 = v39
					} else {
						v50 = v38 + base.I32_wrap_i64(v43)
					}
				}
				*(*int32)(unsafe.Add(mBase, uint32(l0)+104)) = v50
				v52 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
				*(*int64)(unsafe.Add(mBase, uint32(l0)+120)) = v37 + base.I64_extend_i32_s(v52-v38)
				if base.Ui32(v52) < base.Ui32(v38) {
				} else {
					*(*uint8)(unsafe.Add(mBase, uint32(v38+int32(-1)))) = uint8(v17)
				}
				return v17
			} else {
				v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
				v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
				v25 = v23
				v26 = v24
				*(*int64)(unsafe.Add(mBase, uint32(l0)+112)) = int64(-1)
				*(*int32)(unsafe.Add(mBase, uint32(l0)+104)) = v25
				*(*int64)(unsafe.Add(mBase, uint32(l0)+120)) = v12 + base.I64_extend_i32_s(v26-v25)
				return int32(-1)
			}
		}
	} else {
		if v13 <= v12 {
			v25 = v8
			v26 = v9
			*(*int64)(unsafe.Add(mBase, uint32(l0)+112)) = int64(-1)
			*(*int32)(unsafe.Add(mBase, uint32(l0)+104)) = v25
			*(*int64)(unsafe.Add(mBase, uint32(l0)+120)) = v12 + base.I64_extend_i32_s(v26-v25)
			return int32(-1)
		} else {
			v17 = F___uflow(m, l0)
			mBase = m.M
			v20 = m.ExcPending
			if v20 != 0 {
				return int32(0)
			} else {
				if int32(-1) < v17 {
					v37 = v12 + int64(1)
					v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
					v39 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
					v40 = *(*int64)(unsafe.Add(mBase, uint32(l0)+112))
					if v40 == int64(0) {
						v50 = v39
					} else {
						v43 = v40 - v37
						if base.I64_extend_i32_s(v39-v38) <= v43 {
							v50 = v39
						} else {
							v50 = v38 + base.I32_wrap_i64(v43)
						}
					}
					*(*int32)(unsafe.Add(mBase, uint32(l0)+104)) = v50
					v52 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
					*(*int64)(unsafe.Add(mBase, uint32(l0)+120)) = v37 + base.I64_extend_i32_s(v52-v38)
					if base.Ui32(v52) < base.Ui32(v38) {
					} else {
						*(*uint8)(unsafe.Add(mBase, uint32(v38+int32(-1)))) = uint8(v17)
					}
					return v17
				} else {
					v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
					v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
					v25 = v23
					v26 = v24
					*(*int64)(unsafe.Add(mBase, uint32(l0)+112)) = int64(-1)
					*(*int32)(unsafe.Add(mBase, uint32(l0)+104)) = v25
					*(*int64)(unsafe.Add(mBase, uint32(l0)+120)) = v12 + base.I64_extend_i32_s(v26-v25)
					return int32(-1)
				}
			}
		}
	}
}
func F___small_sprintf(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	*(*int32)(unsafe.Add(mBase, uint32(v7)+12)) = l2
	v10 = F___small_vsprintf(m, l0, l1, l2)
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return int32(0)
	} else {
		m.G0 = v7 + int32(16)
		return v10
	}
}
func F___srandom(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v25 int32
	_ = v25
	var v30 int32
	_ = v30
	var v33 int64
	_ = v33
	var v40 int64
	_ = v40
	var v42 int64
	_ = v42
	var v45 int32
	_ = v45
	var v51 int32
	_ = v51
	v6 = *(*int32)(unsafe.Add(mBase, _consts[184]))
	if v6 != 0 {
		v11 = int32(3)
		if v6 == int32(7) {
			v16 = v11
		} else {
			v16 = int32(1)
		}
		if v6 == int32(31) {
			v19 = v11
		} else {
			v19 = v16
		}
		*(*int32)(unsafe.Add(mBase, _consts[185])) = v19
		v21 = int32(0)
		*(*int32)(unsafe.Add(mBase, _consts[186])) = v21
		v25 = *(*int32)(unsafe.Add(mBase, _consts[183]))
		if v6 < int32(1) {
		} else {
			v30 = int32(0)
			v33 = base.I64_extend_i32_u(l0)
			for {
				v40 = v33*int64(6364136223846793005) + int64(1)
				v42 = int64(base.Ui64(v40) >> (uint(int64(32)) % 64))
				*(*uint32)(unsafe.Add(mBase, uint32(v25+v30<<(uint(int32(2))%32)))) = uint32(v42)
				v45 = v30 + int32(1)
				if v45 != v6 {
					v30 = v45
					v33 = v40
					continue
				} else {
					break
				}
				break
			}
		}
		v51 = *(*int32)(unsafe.Add(mBase, uint32(v25)))
		*(*int32)(unsafe.Add(mBase, uint32(v25))) = v51 | int32(1)
		return
	} else {
		v8 = *(*int32)(unsafe.Add(mBase, _consts[183]))
		*(*int32)(unsafe.Add(mBase, uint32(v8))) = l0
		return
	}
}
func F___stdio_close(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	v3 = m.Wasi_snapshot_preview1.Fd_close(m, v2)
	mBase = m.M
	if v3 != 0 {
		v5 = F___errno_location(m)
		mBase = m.M
		*(*int32)(unsafe.Add(mBase, uint32(v5))) = v3
		v8 = int32(-1)
	} else {
		v8 = int32(0)
	}
	return v8
}
func F___subtf3(m *base.Module, l0 int32, l1 int64, l2 int64, l3 int64, l4 int64) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v14 int64
	_ = v14
	var v17 int64
	_ = v17
	v7 = m.G0
	v8 = int32(16)
	v9 = v7 - v8
	m.G0 = v9
	F___addtf3(m, v9, l1, l2, l3, l4^int64(-9223372036854775807-1))
	mBase = m.M
	v14 = *(*int64)(unsafe.Add(mBase, uint32(v9)))
	v17 = *(*int64)(unsafe.Add(mBase, uint32(v9+int32(8))))
	*(*int64)(unsafe.Add(mBase, uint32(l0)+8)) = v17
	*(*int64)(unsafe.Add(mBase, uint32(l0))) = v14
	m.G0 = v9 + v8
	return
}
func F__sdsnewlen(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v95 int32
	_ = v95
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	var v106 int32
	_ = v106
	var v109 int32
	_ = v109
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v134 int32
	_ = v134
	v9 = m.G0
	v11 = v9 - int32(16)
	m.G0 = v11
	if base.Ui32(int32(32)) <= base.Ui32(l1) {
		if base.Ui32(int32(253)) <= base.Ui32(l1) {
			if base.Ui32(l1) < base.Ui32(int32(65531)) {
				v23 = int32(2)
			} else {
				v23 = int32(3)
			}
			v24 = v23
		} else {
			v24 = int32(1)
		}
	} else {
		v24 = int32(0)
	}
	if v24 != 0 {
		v26 = v24
	} else {
		v26 = int32(1)
	}
	if l1 != 0 {
		v27 = v24
	} else {
		v27 = v26
	}
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v27<<(uint(int32(2))%32))+uint32(_consts[721])))
	v35 = l1 + v32 + int32(1)
	if base.Ui32(v35) <= base.Ui32(l1) {
		F__serverAssert(m, int32(_a1320), int32(_a1321), int32(100))
		mBase = m.M
		v134 = m.ExcPending
		if v134 != 0 {
			return int32(0)
		} else {
			F_abort(m)
			mBase = m.M
			base.Wasm_trap_unreachable()
			for {
			}
		}
	} else {
		if l2 == int32(0) {
			v95 = F_zmalloc_usable(m, v35, v11+int32(12))
			mBase = m.M
			v98 = m.ExcPending
			if v98 != 0 {
				return int32(0)
			} else {
				v99 = v95
				if v99 != 0 {
					v101 = *(*int32)(unsafe.Add(mBase, uint32(v11)+12))
					switch v27 + int32(-1) {
					case 0:
						v106 = int32(255)
						v109 = v101 + (v32 ^ int32(-1))
						if base.Ui32(v109) <= base.Ui32(v106) {
							v116 = v27
						} else {
							if base.Ui32(v109) < base.Ui32(int32(65531)) {
								v115 = int32(2)
							} else {
								v115 = int32(3)
							}
							v116 = v115
						}
					case 1:
						v106 = int32(65535)
						v109 = v101 + (v32 ^ int32(-1))
						if base.Ui32(v109) <= base.Ui32(v106) {
							v116 = v27
						} else {
							if base.Ui32(v109) < base.Ui32(int32(65531)) {
								v115 = int32(2)
							} else {
								v115 = int32(3)
							}
							v116 = v115
						}
					default:
						v116 = v27
					}
					v119 = F_sdswrite(m, v99, v101, v116, l0, l1)
					mBase = m.M
					v120 = m.ExcPending
					if v120 != 0 {
						return int32(0)
					} else {
						v121 = v119
						m.G0 = v11 + int32(16)
						return v121
					}
				} else {
					v121 = int32(0)
					m.G0 = v11 + int32(16)
					return v121
				}
			}
		} else {
			v40 = v11 + int32(12)
			v41 = int32(0)
			if base.Ui32(int32(2147483646)) < base.Ui32(v35) {
				v87 = v41
				v88 = v41
			} else {
				if v35 != 0 {
					v49 = v35
				} else {
					v49 = int32(4)
				}
				v51 = v49 + int32(8)
				v52 = F_emscripten_builtin_malloc(m, v51)
				mBase = m.M
				if v52 != 0 {
					*(*int32)(unsafe.Add(mBase, uint32(v52))) = v49
					v57 = *(*int32)(unsafe.Add(mBase, _consts[411]))
					if v57 != int32(-1) {
						v68 = v57
					} else {
						v60 = int32(0)
						v62 = *(*int32)(unsafe.Add(mBase, _consts[281]))
						*(*int32)(unsafe.Add(mBase, _consts[411])) = v62
						*(*int32)(unsafe.Add(mBase, _consts[281])) = v62 + int32(1)
						v68 = v62
					}
					if v68 < int32(260) {
						v77 = v68 << (uint(int32(2)) % 32)
						v80 = *(*int32)(unsafe.Add(mBase, uint32(v77)+uint32(_consts[285])))
						*(*int32)(unsafe.Add(mBase, uint32(v77)+uint32(_consts[285]))) = v80 + v51
					} else {
						v71 = int32(0)
						v73 = *(*int32)(unsafe.Add(mBase, _consts[286]))
						*(*int32)(unsafe.Add(mBase, _consts[286])) = v73 + v51
					}
					v87 = v49
					v88 = v52 + int32(8)
				} else {
					v53 = int32(0)
					v87 = v53
					v88 = v53
				}
			}
			if v40 == int32(0) {
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v40))) = v87
			}
			v99 = v88
			if v99 != 0 {
				v101 = *(*int32)(unsafe.Add(mBase, uint32(v11)+12))
				switch v27 + int32(-1) {
				case 0:
					v106 = int32(255)
					v109 = v101 + (v32 ^ int32(-1))
					if base.Ui32(v109) <= base.Ui32(v106) {
						v116 = v27
					} else {
						if base.Ui32(v109) < base.Ui32(int32(65531)) {
							v115 = int32(2)
						} else {
							v115 = int32(3)
						}
						v116 = v115
					}
				case 1:
					v106 = int32(65535)
					v109 = v101 + (v32 ^ int32(-1))
					if base.Ui32(v109) <= base.Ui32(v106) {
						v116 = v27
					} else {
						if base.Ui32(v109) < base.Ui32(int32(65531)) {
							v115 = int32(2)
						} else {
							v115 = int32(3)
						}
						v116 = v115
					}
				default:
					v116 = v27
				}
				v119 = F_sdswrite(m, v99, v101, v116, l0, l1)
				mBase = m.M
				v120 = m.ExcPending
				if v120 != 0 {
					return int32(0)
				} else {
					v121 = v119
					m.G0 = v11 + int32(16)
					return v121
				}
			} else {
				v121 = int32(0)
				m.G0 = v11 + int32(16)
				return v121
			}
		}
	}
}
func F_scalbn(m *base.Module, l0 float64, l1 int32) float64 {
	var v6 float64
	_ = v6
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v22 float64
	_ = v22
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v35 float64
	_ = v35
	var v36 int32
	_ = v36
	if l1 < int32(1024) {
		if int32(-1023) < l1 {
			v35 = l0
			v36 = l1
		} else {
			v22 = base.F64_mul(l0, float64(2.004168360008973e-292))
			if base.Ui32(l1) <= base.Ui32(int32(-1992)) {
				v29 = int32(-2960)
				if base.Ui32(v29) < base.Ui32(l1) {
					v32 = l1
				} else {
					v32 = v29
				}
				v35 = base.F64_mul(v22, float64(2.004168360008973e-292))
				v36 = v32 + int32(1938)
			} else {
				v35 = v22
				v36 = l1 + int32(969)
			}
		}
	} else {
		v6 = base.F64_mul(l0, float64(8.98846567431158e+307))
		if base.Ui32(int32(2047)) <= base.Ui32(l1) {
			v13 = int32(3069)
			if base.Ui32(l1) < base.Ui32(v13) {
				v16 = l1
			} else {
				v16 = v13
			}
			v35 = base.F64_mul(v6, float64(8.98846567431158e+307))
			v36 = v16 + int32(-2046)
		} else {
			v35 = v6
			v36 = l1 + int32(-1023)
		}
	}
	return base.F64_mul(v35, base.F64_reinterpret_i64(base.I64_extend_i32_u(v36+int32(1023))<<(uint(int64(52))%64)))
}
func F_scanexp(m *base.Module, l0 int32, l1 int32) int64 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v40 int64
	_ = v40
	var v43 int32
	_ = v43
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
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
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v89 int64
	_ = v89
	var v94 int32
	_ = v94
	var v98 int64
	_ = v98
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v114 int64
	_ = v114
	var v116 int32
	_ = v116
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v150 int64
	_ = v150
	var v151 int64
	_ = v151
	var v154 int32
	_ = v154
	var v160 int64
	_ = v160
	var v165 int64
	_ = v165
	var v168 int32
	_ = v168
	var v180 int64
	_ = v180
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
	if v8 == v9 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	switch v19 + int32(-43) {
	case 0, 2:
		goto L10
	default:
		goto L9
	}
L2:
	;
	v15 = F___shgetc(m, l0)
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		goto L4
	} else {
		goto L5
	}
L3:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v8 + int32(1)
	v14 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8))))
	v19 = v14
	goto L1
L4:
	;
	return int64(0)
L5:
	;
	v19 = v15
	goto L1
L6:
	;
	return v180
L7:
	;
	v165 = *(*int64)(unsafe.Add(mBase, uint32(l0)+112))
	if v165 < int64(0) {
		v180 = int64(-9223372036854775807 - 1)
		goto L6
	} else {
		goto L53
	}
L8:
	;
	if base.Ui32(v52) < base.Ui32(int32(-10)) {
		goto L7
	} else {
		goto L18
	}
L9:
	;
	v50 = v19
	v51 = int32(0)
	v52 = v19 + int32(-58)
	goto L8
L10:
	;
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
	if v22 == v23 {
		goto L12
	} else {
		goto L13
	}
L11:
	;
	v33 = base.B2i32(v19 == int32(45))
	v35 = v31 + int32(-58)
	if l1 == int32(0) {
		v50 = v31
		v51 = v33
		v52 = v35
		goto L8
	} else {
		goto L15
	}
L12:
	;
	v29 = F___shgetc(m, l0)
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		goto L4
	} else {
		goto L14
	}
L13:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v22 + int32(1)
	v28 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22))))
	v31 = v28
	goto L11
L14:
	;
	v31 = v29
	goto L11
L15:
	;
	if base.Ui32(int32(-11)) < base.Ui32(v35) {
		v50 = v31
		v51 = v33
		v52 = v35
		goto L8
	} else {
		goto L16
	}
L16:
	;
	v40 = *(*int64)(unsafe.Add(mBase, uint32(l0)+112))
	if v40 < int64(0) {
		goto L7
	} else {
		goto L17
	}
L17:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v43 + int32(-1)
	goto L7
L18:
	;
	if base.Ui32(int32(10)) <= base.Ui32(v50+int32(-48)) {
		v150 = int64(0)
		goto L19
	} else {
		goto L20
	}
L19:
	;
	v151 = *(*int64)(unsafe.Add(mBase, uint32(l0)+112))
	if v151 < int64(0) {
		goto L48
	} else {
		goto L49
	}
L20:
	;
	v63 = v50
	v64 = int32(0)
	goto L21
L21:
	;
	v71 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v72 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
	if v71 == v72 {
		goto L24
	} else {
		goto L25
	}
L22:
	;
	v89 = base.I64_extend_i32_s(v82)
	if base.Ui32(int32(10)) <= base.Ui32(v84) {
		v150 = v89
		goto L19
	} else {
		goto L30
	}
L23:
	;
	v81 = int32(-48)
	v82 = v63 + v64*int32(10) + v81
	v84 = v80 + v81
	if base.Ui32(int32(9)) < base.Ui32(v84) {
		goto L27
	} else {
		goto L28
	}
L24:
	;
	v78 = F___shgetc(m, l0)
	mBase = m.M
	v79 = m.ExcPending
	if v79 != 0 {
		goto L4
	} else {
		goto L26
	}
L25:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v71 + int32(1)
	v77 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v71))))
	v80 = v77
	goto L23
L26:
	;
	v80 = v78
	goto L23
L27:
	;
	goto L22
L28:
	;
	if v82 < int32(214748364) {
		v63 = v80
		v64 = v82
		goto L21
	} else {
		goto L29
	}
L29:
	;
	goto L27
L30:
	;
	v94 = v80
	v98 = v89
	goto L31
L31:
	;
	v103 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v104 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
	if v103 == v104 {
		goto L34
	} else {
		goto L35
	}
L32:
	;
	if base.Ui32(int32(10)) <= base.Ui32(v116) {
		v150 = v114
		goto L19
	} else {
		goto L40
	}
L33:
	;
	v114 = base.I64_extend_i32_u(v94) + v98*int64(10) + int64(-48)
	v116 = v112 + int32(-48)
	if base.Ui32(int32(9)) < base.Ui32(v116) {
		goto L37
	} else {
		goto L38
	}
L34:
	;
	v110 = F___shgetc(m, l0)
	mBase = m.M
	v111 = m.ExcPending
	if v111 != 0 {
		goto L4
	} else {
		goto L36
	}
L35:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v103 + int32(1)
	v109 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v103))))
	v112 = v109
	goto L33
L36:
	;
	v112 = v110
	goto L33
L37:
	;
	goto L32
L38:
	;
	if v114 < int64(92233720368547758) {
		v94 = v112
		v98 = v114
		goto L31
	} else {
		goto L39
	}
L39:
	;
	goto L37
L40:
	;
	goto L41
L41:
	;
	v130 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v131 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
	if v130 == v131 {
		goto L44
	} else {
		goto L45
	}
L42:
	;
	v150 = v114
	goto L19
L43:
	;
	if base.Ui32(v139+int32(-48)) < base.Ui32(int32(10)) {
		goto L41
	} else {
		goto L47
	}
L44:
	;
	v137 = F___shgetc(m, l0)
	mBase = m.M
	v138 = m.ExcPending
	if v138 != 0 {
		goto L4
	} else {
		goto L46
	}
L45:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v130 + int32(1)
	v136 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v130))))
	v139 = v136
	goto L43
L46:
	;
	v139 = v137
	goto L43
L47:
	;
	goto L42
L48:
	;
	if v51 != 0 {
		goto L50
	} else {
		goto L51
	}
L49:
	;
	v154 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v154 + int32(-1)
	goto L48
L50:
	;
	v160 = int64(0) - v150
	goto L52
L51:
	;
	v160 = v150
	goto L52
L52:
	;
	v180 = v160
	goto L6
L53:
	;
	v168 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v168 + int32(-1)
	return int64(-9223372036854775807 - 1)
}
func F_sdiffCommand(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v5 int32
	_ = v5
	var v11 int32
	_ = v11
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	F_sunionDiffGenericCommand(m, l0, v2+int32(4), v5+int32(-1), int32(0), int32(1))
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return
	} else {
		return
	}
}
func F_sdscatfmt(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v16 int32
	_ = v16
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v59 int32
	_ = v59
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v74 int32
	_ = v74
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	var v95 int32
	_ = v95
	var v97 int32
	_ = v97
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v104 int32
	_ = v104
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
	var v117 int32
	_ = v117
	var v120 int32
	_ = v120
	var v124 int32
	_ = v124
	var v127 int32
	_ = v127
	var v131 int32
	_ = v131
	var v134 int32
	_ = v134
	var v138 int64
	_ = v138
	var v141 int64
	_ = v141
	var v144 int32
	_ = v144
	var v146 int32
	_ = v146
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v168 int32
	_ = v168
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v174 int32
	_ = v174
	var v178 int32
	_ = v178
	var v188 int32
	_ = v188
	var v191 int32
	_ = v191
	var v195 int32
	_ = v195
	var v200 int32
	_ = v200
	var v202 int32
	_ = v202
	var v206 int32
	_ = v206
	var v212 int32
	_ = v212
	var v215 int32
	_ = v215
	var v221 int32
	_ = v221
	var v225 int32
	_ = v225
	var v227 int32
	_ = v227
	var v235 int32
	_ = v235
	var v239 int32
	_ = v239
	var v246 int32
	_ = v246
	var v249 int32
	_ = v249
	var v252 int32
	_ = v252
	var v255 int32
	_ = v255
	var v256 int32
	_ = v256
	var v259 int32
	_ = v259
	var v261 int32
	_ = v261
	var v268 int32
	_ = v268
	var v271 int32
	_ = v271
	var v275 int32
	_ = v275
	var v278 int32
	_ = v278
	var v282 int32
	_ = v282
	var v285 int32
	_ = v285
	var v289 int64
	_ = v289
	var v292 int64
	_ = v292
	var v295 int32
	_ = v295
	var v298 int32
	_ = v298
	var v299 int32
	_ = v299
	var v300 int32
	_ = v300
	var v304 int32
	_ = v304
	var v307 int32
	_ = v307
	var v308 int32
	_ = v308
	var v315 int32
	_ = v315
	var v319 int32
	_ = v319
	var v320 int32
	_ = v320
	var v321 int32
	_ = v321
	var v325 int32
	_ = v325
	var v326 int32
	_ = v326
	var v327 int32
	_ = v327
	var v331 int32
	_ = v331
	var v332 int32
	_ = v332
	var v337 int32
	_ = v337
	var v338 int64
	_ = v338
	var v344 int32
	_ = v344
	var v348 int64
	_ = v348
	var v349 int32
	_ = v349
	var v353 int32
	_ = v353
	var v357 int64
	_ = v357
	var v359 int64
	_ = v359
	var v369 int32
	_ = v369
	var v373 int32
	_ = v373
	var v378 int32
	_ = v378
	var v379 int32
	_ = v379
	var v380 int64
	_ = v380
	var v381 int32
	_ = v381
	var v382 int32
	_ = v382
	var v401 int32
	_ = v401
	var v403 int32
	_ = v403
	var v405 int32
	_ = v405
	var v412 int32
	_ = v412
	var v415 int32
	_ = v415
	var v419 int32
	_ = v419
	var v422 int32
	_ = v422
	var v426 int32
	_ = v426
	var v429 int32
	_ = v429
	var v433 int64
	_ = v433
	var v436 int64
	_ = v436
	var v439 int32
	_ = v439
	var v442 int32
	_ = v442
	var v443 int32
	_ = v443
	var v444 int32
	_ = v444
	var v448 int32
	_ = v448
	var v451 int32
	_ = v451
	var v452 int32
	_ = v452
	var v459 int32
	_ = v459
	var v463 int32
	_ = v463
	var v464 int32
	_ = v464
	var v465 int32
	_ = v465
	var v469 int32
	_ = v469
	var v470 int32
	_ = v470
	var v471 int32
	_ = v471
	var v475 int32
	_ = v475
	var v476 int32
	_ = v476
	var v481 int32
	_ = v481
	var v482 int64
	_ = v482
	var v488 int32
	_ = v488
	var v492 int64
	_ = v492
	var v493 int32
	_ = v493
	var v497 int32
	_ = v497
	var v501 int64
	_ = v501
	var v503 int64
	_ = v503
	var v505 int32
	_ = v505
	var v509 int32
	_ = v509
	var v517 int32
	_ = v517
	var v518 int64
	_ = v518
	var v539 int32
	_ = v539
	var v544 int32
	_ = v544
	var v551 int32
	_ = v551
	var v556 int32
	_ = v556
	var v558 int32
	_ = v558
	var v562 int64
	_ = v562
	var v566 int32
	_ = v566
	var v567 int32
	_ = v567
	var v570 int32
	_ = v570
	var v573 int32
	_ = v573
	var v576 int32
	_ = v576
	var v583 int64
	_ = v583
	var v585 int32
	_ = v585
	var v589 int64
	_ = v589
	var v590 int64
	_ = v590
	var v599 int32
	_ = v599
	var v602 int32
	_ = v602
	var v609 int32
	_ = v609
	var v610 int64
	_ = v610
	var v612 int32
	_ = v612
	var v617 int32
	_ = v617
	var v626 int32
	_ = v626
	var v630 int32
	_ = v630
	var v640 int32
	_ = v640
	var v642 int32
	_ = v642
	var v644 int32
	_ = v644
	var v651 int32
	_ = v651
	var v654 int32
	_ = v654
	var v658 int32
	_ = v658
	var v661 int32
	_ = v661
	var v665 int32
	_ = v665
	var v668 int32
	_ = v668
	var v672 int64
	_ = v672
	var v675 int64
	_ = v675
	var v678 int32
	_ = v678
	var v681 int32
	_ = v681
	var v682 int32
	_ = v682
	var v683 int32
	_ = v683
	var v687 int32
	_ = v687
	var v690 int32
	_ = v690
	var v691 int32
	_ = v691
	var v698 int32
	_ = v698
	var v702 int32
	_ = v702
	var v703 int32
	_ = v703
	var v704 int32
	_ = v704
	var v708 int32
	_ = v708
	var v709 int32
	_ = v709
	var v710 int32
	_ = v710
	var v714 int32
	_ = v714
	var v715 int32
	_ = v715
	var v720 int32
	_ = v720
	var v721 int64
	_ = v721
	var v730 int32
	_ = v730
	var v733 int32
	_ = v733
	var v734 int32
	_ = v734
	var v736 int32
	_ = v736
	var v739 int32
	_ = v739
	var v740 int32
	_ = v740
	var v742 int32
	_ = v742
	var v745 int32
	_ = v745
	var v746 int32
	_ = v746
	var v751 int32
	_ = v751
	var v752 int64
	_ = v752
	var v759 int32
	_ = v759
	var v761 int32
	_ = v761
	var v762 int32
	_ = v762
	var v768 int32
	_ = v768
	var v771 int32
	_ = v771
	var v772 int32
	_ = v772
	var v774 int32
	_ = v774
	var v777 int32
	_ = v777
	var v778 int32
	_ = v778
	var v780 int32
	_ = v780
	var v783 int32
	_ = v783
	var v784 int32
	_ = v784
	var v789 int32
	_ = v789
	var v790 int64
	_ = v790
	var v795 int32
	_ = v795
	var v798 int32
	_ = v798
	var v800 int32
	_ = v800
	var v801 int32
	_ = v801
	var v807 int32
	_ = v807
	var v808 int32
	_ = v808
	var v812 int32
	_ = v812
	var v817 int32
	_ = v817
	v9 = m.G0
	v11 = v9 - int32(32)
	m.G0 = v11
	v16 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0+int32(-1)))))
	switch v16 & int32(7) {
	case 0:
		goto L6
	case 1:
		goto L5
	case 2:
		goto L4
	case 3:
		goto L3
	case 4:
		goto L2
	default:
		v33 = int32(0)
		goto L1
	}
L1:
	;
	if l1&int32(3) == int32(0) {
		v55 = l1
		goto L9
	} else {
		goto L10
	}
L2:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(l0+int32(-17))))
	v33 = v32
	goto L1
L3:
	;
	v29 = *(*int32)(unsafe.Add(mBase, uint32(l0+int32(-9))))
	v33 = v29
	goto L1
L4:
	;
	v26 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0+int32(-5)))))
	v33 = v26
	goto L1
L5:
	;
	v23 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0+int32(-3)))))
	v33 = v23
	goto L1
L6:
	;
	v33 = int32(base.Ui32(v16) >> (uint(int32(3)) % 32))
	goto L1
L7:
	;
	v89 = int32(1)
	v92 = F__sdsMakeRoomFor(m, l0, v88<<(uint(v89)%32), v89)
	mBase = m.M
	v95 = m.ExcPending
	if v95 != 0 {
		goto L23
	} else {
		goto L24
	}
L8:
	;
	v88 = v80 - l1
	goto L7
L9:
	;
	v59 = v55
	goto L17
L10:
	;
	v41 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	if v41 != 0 {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v44 = l1
	goto L13
L12:
	;
	v88 = l1 - l1
	goto L7
L13:
	;
	v48 = v44 + int32(1)
	if v48&int32(3) == int32(0) {
		v55 = v48
		goto L9
	} else {
		goto L15
	}
L15:
	;
	v53 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v48))))
	if v53 != 0 {
		v44 = v48
		goto L13
	} else {
		goto L16
	}
L16:
	;
	v80 = v48
	goto L8
L17:
	;
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v59)))
	v68 = int32(-2139062144)
	if (int32(16843008)-v65|v65)&v68 == v68 {
		v59 = v59 + int32(4)
		goto L17
	} else {
		goto L19
	}
L18:
	;
	v74 = v59
	goto L20
L19:
	;
	goto L18
L20:
	;
	v78 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v74))))
	if v78 != 0 {
		v74 = v74 + int32(1)
		goto L20
	} else {
		goto L22
	}
L21:
	;
	v80 = v74
	goto L8
L22:
	;
	goto L21
L23:
	;
	return int32(0)
L24:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+28)) = l2
	v97 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	if v97 == int32(0) {
		v808 = v92
		v812 = v33
		goto L25
	} else {
		goto L26
	}
L25:
	;
	v817 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v808+v812))) = uint8(v817)
	m.G0 = v11 + int32(32)
	return v808
L26:
	;
	v100 = v92
	v101 = l1
	v102 = v97
	v104 = v33
	goto L27
L27:
	;
	v108 = int32(-1)
	v110 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v100+v108))))
	switch v110&int32(7) + v108 {
	case 0:
		goto L35
	case 1:
		goto L34
	case 2:
		goto L33
	case 3:
		goto L32
	default:
		goto L30
	}
L28:
	;
	v808 = v798
	v812 = v801
	goto L25
L29:
	;
	if v152&int32(255) != int32(37) {
		goto L40
	} else {
		goto L41
	}
L30:
	;
	v146 = int32(1)
	v148 = F__sdsMakeRoomFor(m, v100, v146, v146)
	mBase = m.M
	v149 = m.ExcPending
	if v149 != 0 {
		goto L23
	} else {
		goto L37
	}
L31:
	;
	if v144 != 0 {
		v151 = v100
		v152 = v102
		goto L29
	} else {
		goto L36
	}
L32:
	;
	v138 = *(*int64)(unsafe.Add(mBase, uint32(v100+int32(-9))))
	v141 = *(*int64)(unsafe.Add(mBase, uint32(v100+int32(-17))))
	v144 = base.I32_wrap_i64(v138 - v141)
	goto L31
L33:
	;
	v131 = *(*int32)(unsafe.Add(mBase, uint32(v100+int32(-5))))
	v134 = *(*int32)(unsafe.Add(mBase, uint32(v100+int32(-9))))
	v144 = v131 - v134
	goto L31
L34:
	;
	v124 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v100+int32(-3)))))
	v127 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v100+int32(-5)))))
	v144 = v124 - v127
	goto L31
L35:
	;
	v117 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v100+int32(-2)))))
	v120 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v100+int32(-3)))))
	v144 = v117 - v120
	goto L31
L36:
	;
	goto L30
L37:
	;
	v150 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v101))))
	v151 = v148
	v152 = v150
	goto L29
L38:
	;
	v807 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v800)+1)))
	if v807 != 0 {
		v100 = v798
		v101 = v800 + int32(1)
		v102 = v807
		v104 = v801
		goto L27
	} else {
		goto L190
	}
L39:
	;
	v798 = v151
	v800 = v101
	v801 = v795
	goto L38
L40:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v151+v104))) = uint8(v152)
	v759 = v104 + int32(1)
	v761 = v151 + int32(-1)
	v762 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v761))))
	switch v762 & int32(7) {
	case 0:
		goto L189
	case 1:
		goto L188
	case 2:
		goto L187
	case 3:
		goto L186
	case 4:
		goto L185
	default:
		v798 = v151
		v800 = v101
		v801 = v759
		goto L38
	}
L41:
	;
	v159 = v101 + int32(1)
	v160 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v101)+1)))
	switch v160 + int32(-73) {
	case 0:
		goto L51
	case 1, 2, 3, 4, 5, 6, 7, 8, 9, 11, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 33, 34, 35, 36, 37, 38, 39, 40, 41, 43:
		goto L54
	case 10, 42:
		goto L53
	case 12:
		goto L48
	case 32:
		goto L52
	case 44:
		goto L49
	default:
		goto L55
	}
L42:
	;
	v751 = v151 + int32(-17)
	v752 = *(*int64)(unsafe.Add(mBase, uint32(v751)))
	*(*int64)(unsafe.Add(mBase, uint32(v751))) = v752 + int64(1)
	v798 = v151
	v800 = v159
	v801 = v168
	goto L38
L43:
	;
	v745 = v151 + int32(-9)
	v746 = *(*int32)(unsafe.Add(mBase, uint32(v745)))
	*(*int32)(unsafe.Add(mBase, uint32(v745))) = v746 + int32(1)
	v798 = v151
	v800 = v159
	v801 = v168
	goto L38
L44:
	;
	v739 = v151 + int32(-5)
	v740 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v739))))
	v742 = v740 + int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v739))) = uint16(v742)
	v798 = v151
	v800 = v159
	v801 = v168
	goto L38
L45:
	;
	v733 = v151 + int32(-3)
	v734 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v733))))
	v736 = v734 + int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v733))) = uint8(v736)
	v798 = v151
	v800 = v159
	v801 = v168
	goto L38
L46:
	;
	v730 = v171&int32(248) + int32(8)
	*(*uint8)(unsafe.Add(mBase, uint32(v170))) = uint8(v730)
	v798 = v151
	v800 = v159
	v801 = v168
	goto L38
L47:
	;
	v505 = int32(0)
	v509 = int32(1)
	if base.Ui64(v503) < base.Ui64(int64(10)) {
		v566 = v509
		v567 = v505
		goto L125
	} else {
		goto L126
	}
L48:
	;
	v493 = *(*int32)(unsafe.Add(mBase, uint32(v11)+28))
	v497 = (v493 + int32(7)) & int32(-8)
	*(*int32)(unsafe.Add(mBase, uint32(v11)+28)) = v497 + int32(8)
	v501 = *(*int64)(unsafe.Add(mBase, uint32(v497)))
	v503 = v501
	goto L47
L49:
	;
	v488 = *(*int32)(unsafe.Add(mBase, uint32(v11)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+28)) = v488 + int32(4)
	v492 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v488))))
	v503 = v492
	goto L47
L50:
	;
	if v359 <= int64(-1) {
		goto L102
	} else {
		goto L103
	}
L51:
	;
	v349 = *(*int32)(unsafe.Add(mBase, uint32(v11)+28))
	v353 = (v349 + int32(7)) & int32(-8)
	*(*int32)(unsafe.Add(mBase, uint32(v11)+28)) = v353 + int32(8)
	v357 = *(*int64)(unsafe.Add(mBase, uint32(v353)))
	v359 = v357
	goto L50
L52:
	;
	v344 = *(*int32)(unsafe.Add(mBase, uint32(v11)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+28)) = v344 + int32(4)
	v348 = int64(*(*int32)(unsafe.Add(mBase, uint32(v344))))
	v359 = v348
	goto L50
L53:
	;
	v174 = *(*int32)(unsafe.Add(mBase, uint32(v11)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+28)) = v174 + int32(4)
	v178 = *(*int32)(unsafe.Add(mBase, uint32(v174)))
	if v160 != int32(115) {
		goto L58
	} else {
		goto L59
	}
L54:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v151+v104))) = uint8(v160)
	v168 = v104 + int32(1)
	v170 = v151 + int32(-1)
	v171 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v170))))
	switch v171 & int32(7) {
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
		v798 = v151
		v800 = v159
		v801 = v168
		goto L38
	}
L55:
	;
	if v160 == int32(0) {
		v795 = v104
		goto L39
	} else {
		goto L56
	}
L56:
	;
	goto L54
L57:
	;
	v259 = int32(-1)
	v261 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v151+v259))))
	switch v261&int32(7) + v259 {
	case 0:
		goto L85
	case 1:
		goto L84
	case 2:
		goto L83
	case 3:
		goto L82
	default:
		v295 = int32(0)
		goto L81
	}
L58:
	;
	v239 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v178+int32(-1)))))
	switch v239 & int32(7) {
	case 0:
		goto L80
	case 1:
		goto L79
	case 2:
		goto L78
	case 3:
		goto L77
	case 4:
		goto L76
	default:
		v256 = int32(0)
		goto L57
	}
L59:
	;
	if v178&int32(3) == int32(0) {
		v202 = v178
		goto L62
	} else {
		goto L63
	}
L60:
	;
	v256 = v235
	goto L57
L61:
	;
	v235 = v227 - v178
	goto L60
L62:
	;
	v206 = v202
	goto L70
L63:
	;
	v188 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v178))))
	if v188 != 0 {
		goto L64
	} else {
		goto L65
	}
L64:
	;
	v191 = v178
	goto L66
L65:
	;
	v235 = v178 - v178
	goto L60
L66:
	;
	v195 = v191 + int32(1)
	if v195&int32(3) == int32(0) {
		v202 = v195
		goto L62
	} else {
		goto L68
	}
L68:
	;
	v200 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v195))))
	if v200 != 0 {
		v191 = v195
		goto L66
	} else {
		goto L69
	}
L69:
	;
	v227 = v195
	goto L61
L70:
	;
	v212 = *(*int32)(unsafe.Add(mBase, uint32(v206)))
	v215 = int32(-2139062144)
	if (int32(16843008)-v212|v212)&v215 == v215 {
		v206 = v206 + int32(4)
		goto L70
	} else {
		goto L72
	}
L71:
	;
	v221 = v206
	goto L73
L72:
	;
	goto L71
L73:
	;
	v225 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v221))))
	if v225 != 0 {
		v221 = v221 + int32(1)
		goto L73
	} else {
		goto L75
	}
L74:
	;
	v227 = v221
	goto L61
L75:
	;
	goto L74
L76:
	;
	v255 = *(*int32)(unsafe.Add(mBase, uint32(v178+int32(-17))))
	v256 = v255
	goto L57
L77:
	;
	v252 = *(*int32)(unsafe.Add(mBase, uint32(v178+int32(-9))))
	v256 = v252
	goto L57
L78:
	;
	v249 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v178+int32(-5)))))
	v256 = v249
	goto L57
L79:
	;
	v246 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v178+int32(-3)))))
	v256 = v246
	goto L57
L80:
	;
	v256 = int32(base.Ui32(v239) >> (uint(int32(3)) % 32))
	goto L57
L81:
	;
	if base.Ui32(v256) <= base.Ui32(v295) {
		v300 = v151
		goto L86
	} else {
		goto L87
	}
L82:
	;
	v289 = *(*int64)(unsafe.Add(mBase, uint32(v151+int32(-9))))
	v292 = *(*int64)(unsafe.Add(mBase, uint32(v151+int32(-17))))
	v295 = base.I32_wrap_i64(v289 - v292)
	goto L81
L83:
	;
	v282 = *(*int32)(unsafe.Add(mBase, uint32(v151+int32(-5))))
	v285 = *(*int32)(unsafe.Add(mBase, uint32(v151+int32(-9))))
	v295 = v282 - v285
	goto L81
L84:
	;
	v275 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v151+int32(-3)))))
	v278 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v151+int32(-5)))))
	v295 = v275 - v278
	goto L81
L85:
	;
	v268 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v151+int32(-2)))))
	v271 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v151+int32(-3)))))
	v295 = v268 - v271
	goto L81
L86:
	;
	if v256 == int32(0) {
		goto L90
	} else {
		goto L91
	}
L87:
	;
	v298 = F__sdsMakeRoomFor(m, v151, v256, int32(1))
	mBase = m.M
	v299 = m.ExcPending
	if v299 != 0 {
		goto L23
	} else {
		goto L88
	}
L88:
	;
	v300 = v298
	goto L86
L89:
	;
	v307 = v300 + int32(-1)
	v308 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v307))))
	switch v308 & int32(7) {
	case 0:
		goto L97
	case 1:
		goto L96
	case 2:
		goto L95
	case 3:
		goto L94
	case 4:
		goto L93
	default:
		goto L92
	}
L90:
	;
	goto L89
L91:
	;
	v304 = F__emscripten_memcpy_bulkmem(m, v300+v104, v178, v256)
	mBase = m.M
	goto L90
L92:
	;
	v798 = v300
	v800 = v159
	v801 = v256 + v104
	goto L38
L93:
	;
	v337 = v300 + int32(-17)
	v338 = *(*int64)(unsafe.Add(mBase, uint32(v337)))
	*(*int64)(unsafe.Add(mBase, uint32(v337))) = v338 + base.I64_extend_i32_u(v256)
	goto L92
L94:
	;
	v331 = v300 + int32(-9)
	v332 = *(*int32)(unsafe.Add(mBase, uint32(v331)))
	*(*int32)(unsafe.Add(mBase, uint32(v331))) = v332 + v256
	v798 = v300
	v800 = v159
	v801 = v256 + v104
	goto L38
L95:
	;
	v325 = v300 + int32(-5)
	v326 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v325))))
	v327 = v326 + v256
	*(*uint16)(unsafe.Add(mBase, uint32(v325))) = uint16(v327)
	v798 = v300
	v800 = v159
	v801 = v256 + v104
	goto L38
L96:
	;
	v319 = v300 + int32(-3)
	v320 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v319))))
	v321 = v320 + v256
	*(*uint8)(unsafe.Add(mBase, uint32(v319))) = uint8(v321)
	v798 = v300
	v800 = v159
	v801 = v256 + v104
	goto L38
L97:
	;
	v315 = (v308 + v256<<(uint(int32(3))%32)) & int32(248)
	*(*uint8)(unsafe.Add(mBase, uint32(v307))) = uint8(v315)
	v798 = v300
	v800 = v159
	v801 = v256 + v104
	goto L38
L98:
	;
	v403 = int32(-1)
	v405 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v151+v403))))
	switch v405&int32(7) + v403 {
	case 0:
		goto L111
	case 1:
		goto L110
	case 2:
		goto L109
	case 3:
		goto L108
	default:
		v439 = int32(0)
		goto L107
	}
L99:
	;
	v401 = int32(0)
	goto L98
L101:
	;
	v382 = F_ull2string(m, v378, v379, v380)
	mBase = m.M
	if v382 == int32(0) {
		goto L99
	} else {
		goto L105
	}
L102:
	;
	goto L104
L103:
	;
	v378 = v11
	v379 = int32(21)
	v380 = v359
	v381 = int32(0)
	goto L101
L104:
	;
	v369 = int32(45)
	*(*uint8)(unsafe.Add(mBase, uint32(v11))) = uint8(v369)
	v373 = int32(1)
	v378 = v11 + v373
	v379 = int32(20)
	v380 = int64(0) - v359
	v381 = v373
	goto L101
L105:
	;
	v401 = v382 + v381
	goto L98
L107:
	;
	if base.Ui32(v401) <= base.Ui32(v439) {
		v444 = v151
		goto L112
	} else {
		goto L113
	}
L108:
	;
	v433 = *(*int64)(unsafe.Add(mBase, uint32(v151+int32(-9))))
	v436 = *(*int64)(unsafe.Add(mBase, uint32(v151+int32(-17))))
	v439 = base.I32_wrap_i64(v433 - v436)
	goto L107
L109:
	;
	v426 = *(*int32)(unsafe.Add(mBase, uint32(v151+int32(-5))))
	v429 = *(*int32)(unsafe.Add(mBase, uint32(v151+int32(-9))))
	v439 = v426 - v429
	goto L107
L110:
	;
	v419 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v151+int32(-3)))))
	v422 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v151+int32(-5)))))
	v439 = v419 - v422
	goto L107
L111:
	;
	v412 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v151+int32(-2)))))
	v415 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v151+int32(-3)))))
	v439 = v412 - v415
	goto L107
L112:
	;
	if v401 == int32(0) {
		goto L116
	} else {
		goto L117
	}
L113:
	;
	v442 = F__sdsMakeRoomFor(m, v151, v401, int32(1))
	mBase = m.M
	v443 = m.ExcPending
	if v443 != 0 {
		goto L23
	} else {
		goto L114
	}
L114:
	;
	v444 = v442
	goto L112
L115:
	;
	v451 = v444 + int32(-1)
	v452 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v451))))
	switch v452 & int32(7) {
	case 0:
		goto L123
	case 1:
		goto L122
	case 2:
		goto L121
	case 3:
		goto L120
	case 4:
		goto L119
	default:
		goto L118
	}
L116:
	;
	goto L115
L117:
	;
	v448 = F__emscripten_memcpy_bulkmem(m, v444+v104, v11, v401)
	mBase = m.M
	goto L116
L118:
	;
	v798 = v444
	v800 = v159
	v801 = v401 + v104
	goto L38
L119:
	;
	v481 = v444 + int32(-17)
	v482 = *(*int64)(unsafe.Add(mBase, uint32(v481)))
	*(*int64)(unsafe.Add(mBase, uint32(v481))) = v482 + base.I64_extend_i32_u(v401)
	goto L118
L120:
	;
	v475 = v444 + int32(-9)
	v476 = *(*int32)(unsafe.Add(mBase, uint32(v475)))
	*(*int32)(unsafe.Add(mBase, uint32(v475))) = v476 + v401
	v798 = v444
	v800 = v159
	v801 = v401 + v104
	goto L38
L121:
	;
	v469 = v444 + int32(-5)
	v470 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v469))))
	v471 = v470 + v401
	*(*uint16)(unsafe.Add(mBase, uint32(v469))) = uint16(v471)
	v798 = v444
	v800 = v159
	v801 = v401 + v104
	goto L38
L122:
	;
	v463 = v444 + int32(-3)
	v464 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v463))))
	v465 = v464 + v401
	*(*uint8)(unsafe.Add(mBase, uint32(v463))) = uint8(v465)
	v798 = v444
	v800 = v159
	v801 = v401 + v104
	goto L38
L123:
	;
	v459 = (v452 + v401<<(uint(int32(3))%32)) & int32(248)
	*(*uint8)(unsafe.Add(mBase, uint32(v451))) = uint8(v459)
	v798 = v444
	v800 = v159
	v801 = v401 + v104
	goto L38
L124:
	;
	v642 = int32(-1)
	v644 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v151+v642))))
	switch v644&int32(7) + v642 {
	case 0:
		goto L172
	case 1:
		goto L171
	case 2:
		goto L170
	case 3:
		goto L169
	default:
		v678 = int32(0)
		goto L168
	}
L125:
	;
	v570 = v566 + v567
	if base.Ui32(int32(21)) <= base.Ui32(v570) {
		goto L156
	} else {
		goto L157
	}
L126:
	;
	v517 = v505
	v518 = v503
	goto L127
L127:
	;
	if base.Ui64(int64(100)) <= base.Ui64(v518) {
		goto L129
	} else {
		goto L130
	}
L128:
	;
	v566 = v509
	v567 = v558
	goto L125
L129:
	;
	if base.Ui64(int64(1000)) <= base.Ui64(v518) {
		goto L131
	} else {
		goto L132
	}
L130:
	;
	v566 = int32(2)
	v567 = v517
	goto L125
L131:
	;
	if base.Ui64(int64(999999999999)) < base.Ui64(v518) {
		goto L133
	} else {
		goto L134
	}
L132:
	;
	v566 = int32(3)
	v567 = v517
	goto L125
L133:
	;
	v558 = v517 + int32(12)
	v562 = base.I64_div_u_s(v518, int64(1000000000000))
	if base.Ui64(int64(9999999999999)) < base.Ui64(v518) {
		v517 = v558
		v518 = v562
		goto L127
	} else {
		goto L155
	}
L134:
	;
	if base.Ui64(int64(99999999)) < base.Ui64(v518) {
		goto L135
	} else {
		goto L136
	}
L135:
	;
	if base.Ui64(int64(9999999999)) < base.Ui64(v518) {
		goto L147
	} else {
		goto L148
	}
L136:
	;
	if base.Ui64(int64(999999)) < base.Ui64(v518) {
		goto L137
	} else {
		goto L138
	}
L137:
	;
	if base.Ui64(int64(9999999)) < base.Ui64(v518) {
		goto L144
	} else {
		goto L145
	}
L138:
	;
	if base.Ui64(int64(10000)) <= base.Ui64(v518) {
		goto L139
	} else {
		goto L140
	}
L139:
	;
	if base.Ui64(int64(99999)) < base.Ui64(v518) {
		goto L141
	} else {
		goto L142
	}
L140:
	;
	v566 = int32(4)
	v567 = v517
	goto L125
L141:
	;
	v539 = int32(6)
	goto L143
L142:
	;
	v539 = int32(5)
	goto L143
L143:
	;
	v566 = v539
	v567 = v517
	goto L125
L144:
	;
	v544 = int32(8)
	goto L146
L145:
	;
	v544 = int32(7)
	goto L146
L146:
	;
	v566 = v544
	v567 = v517
	goto L125
L147:
	;
	if base.Ui64(int64(99999999999)) < base.Ui64(v518) {
		goto L152
	} else {
		goto L153
	}
L148:
	;
	if base.Ui64(int64(999999999)) < base.Ui64(v518) {
		goto L149
	} else {
		goto L150
	}
L149:
	;
	v551 = int32(10)
	goto L151
L150:
	;
	v551 = int32(9)
	goto L151
L151:
	;
	v566 = v551
	v567 = v517
	goto L125
L152:
	;
	v556 = int32(12)
	goto L154
L153:
	;
	v556 = int32(11)
	goto L154
L154:
	;
	v566 = v556
	v567 = v517
	goto L125
L155:
	;
	goto L128
L156:
	;
	goto L167
L157:
	;
	v573 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v11+v570))) = uint8(v573)
	v576 = v570 + int32(-1)
	if base.Ui64(int64(100)) <= base.Ui64(v503) {
		goto L159
	} else {
		goto L160
	}
L158:
	;
	v612 = v11 + v609
	if base.Ui64(int64(9)) < base.Ui64(v610) {
		goto L164
	} else {
		goto L165
	}
L159:
	;
	v583 = v503
	v585 = v576
	goto L161
L160:
	;
	v609 = v576
	v610 = v503
	goto L158
L161:
	;
	v589 = int64(100)
	v590 = base.I64_div_u_s(v583, v589)
	v599 = int32(*(*uint16)(unsafe.Add(mBase, uint32(base.I32_wrap_i64(v583-v590*v589)<<(uint(int32(1))%32))+uint32(_consts[723]))))
	*(*uint16)(unsafe.Add(mBase, uint32(v11+int32(-1)+v585))) = uint16(v599)
	v602 = v585 + int32(-2)
	if base.Ui64(int64(9999)) < base.Ui64(v583) {
		v583 = v590
		v585 = v602
		goto L161
	} else {
		goto L163
	}
L162:
	;
	v609 = v602
	v610 = v590
	goto L158
L163:
	;
	goto L162
L164:
	;
	v626 = int32(*(*uint16)(unsafe.Add(mBase, uint32(base.I32_wrap_i64(v610)<<(uint(int32(1))%32))+uint32(_consts[723]))))
	*(*uint16)(unsafe.Add(mBase, uint32(v612+int32(-1)))) = uint16(v626)
	v640 = v570
	goto L124
L165:
	;
	v617 = base.I32_wrap_i64(v610) | int32(48)
	*(*uint8)(unsafe.Add(mBase, uint32(v612))) = uint8(v617)
	v640 = v570
	goto L124
L166:
	;
	v640 = int32(0)
	goto L124
L167:
	;
	v630 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v11))) = uint8(v630)
	goto L166
L168:
	;
	if base.Ui32(v640) <= base.Ui32(v678) {
		v683 = v151
		goto L173
	} else {
		goto L174
	}
L169:
	;
	v672 = *(*int64)(unsafe.Add(mBase, uint32(v151+int32(-9))))
	v675 = *(*int64)(unsafe.Add(mBase, uint32(v151+int32(-17))))
	v678 = base.I32_wrap_i64(v672 - v675)
	goto L168
L170:
	;
	v665 = *(*int32)(unsafe.Add(mBase, uint32(v151+int32(-5))))
	v668 = *(*int32)(unsafe.Add(mBase, uint32(v151+int32(-9))))
	v678 = v665 - v668
	goto L168
L171:
	;
	v658 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v151+int32(-3)))))
	v661 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v151+int32(-5)))))
	v678 = v658 - v661
	goto L168
L172:
	;
	v651 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v151+int32(-2)))))
	v654 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v151+int32(-3)))))
	v678 = v651 - v654
	goto L168
L173:
	;
	if v640 == int32(0) {
		goto L177
	} else {
		goto L178
	}
L174:
	;
	v681 = F__sdsMakeRoomFor(m, v151, v640, int32(1))
	mBase = m.M
	v682 = m.ExcPending
	if v682 != 0 {
		goto L23
	} else {
		goto L175
	}
L175:
	;
	v683 = v681
	goto L173
L176:
	;
	v690 = v683 + int32(-1)
	v691 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v690))))
	switch v691 & int32(7) {
	case 0:
		goto L184
	case 1:
		goto L183
	case 2:
		goto L182
	case 3:
		goto L181
	case 4:
		goto L180
	default:
		goto L179
	}
L177:
	;
	goto L176
L178:
	;
	v687 = F__emscripten_memcpy_bulkmem(m, v683+v104, v11, v640)
	mBase = m.M
	goto L177
L179:
	;
	v798 = v683
	v800 = v159
	v801 = v640 + v104
	goto L38
L180:
	;
	v720 = v683 + int32(-17)
	v721 = *(*int64)(unsafe.Add(mBase, uint32(v720)))
	*(*int64)(unsafe.Add(mBase, uint32(v720))) = v721 + base.I64_extend_i32_u(v640)
	goto L179
L181:
	;
	v714 = v683 + int32(-9)
	v715 = *(*int32)(unsafe.Add(mBase, uint32(v714)))
	*(*int32)(unsafe.Add(mBase, uint32(v714))) = v715 + v640
	v798 = v683
	v800 = v159
	v801 = v640 + v104
	goto L38
L182:
	;
	v708 = v683 + int32(-5)
	v709 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v708))))
	v710 = v709 + v640
	*(*uint16)(unsafe.Add(mBase, uint32(v708))) = uint16(v710)
	v798 = v683
	v800 = v159
	v801 = v640 + v104
	goto L38
L183:
	;
	v702 = v683 + int32(-3)
	v703 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v702))))
	v704 = v703 + v640
	*(*uint8)(unsafe.Add(mBase, uint32(v702))) = uint8(v704)
	v798 = v683
	v800 = v159
	v801 = v640 + v104
	goto L38
L184:
	;
	v698 = (v691 + v640<<(uint(int32(3))%32)) & int32(248)
	*(*uint8)(unsafe.Add(mBase, uint32(v690))) = uint8(v698)
	v798 = v683
	v800 = v159
	v801 = v640 + v104
	goto L38
L185:
	;
	v789 = v151 + int32(-17)
	v790 = *(*int64)(unsafe.Add(mBase, uint32(v789)))
	*(*int64)(unsafe.Add(mBase, uint32(v789))) = v790 + int64(1)
	v795 = v759
	goto L39
L186:
	;
	v783 = v151 + int32(-9)
	v784 = *(*int32)(unsafe.Add(mBase, uint32(v783)))
	*(*int32)(unsafe.Add(mBase, uint32(v783))) = v784 + int32(1)
	v795 = v759
	goto L39
L187:
	;
	v777 = v151 + int32(-5)
	v778 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v777))))
	v780 = v778 + int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v777))) = uint16(v780)
	v795 = v759
	goto L39
L188:
	;
	v771 = v151 + int32(-3)
	v772 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v771))))
	v774 = v772 + int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v771))) = uint8(v774)
	v795 = v759
	goto L39
L189:
	;
	v768 = v762&int32(248) + int32(8)
	*(*uint8)(unsafe.Add(mBase, uint32(v761))) = uint8(v768)
	v795 = v759
	goto L39
L190:
	;
	goto L28
}
func F_sdscatlen(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v62 int32
	_ = v62
	v9 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0+int32(-1)))))
	switch v9 & int32(7) {
	case 0:
		v26 = int32(base.Ui32(v9) >> (uint(int32(3)) % 32))
	case 1:
		v16 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0+int32(-3)))))
		v26 = v16
	case 2:
		v19 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0+int32(-5)))))
		v26 = v19
	case 3:
		v22 = *(*int32)(unsafe.Add(mBase, uint32(l0+int32(-9))))
		v26 = v22
	case 4:
		v25 = *(*int32)(unsafe.Add(mBase, uint32(l0+int32(-17))))
		v26 = v25
	default:
		v26 = int32(0)
	}
	v28 = F__sdsMakeRoomFor(m, l0, l2, int32(1))
	mBase = m.M
	v31 = m.ExcPending
	if v31 != 0 {
		return int32(0)
	} else {
		if v28 == int32(0) {
		} else {
			if l2 == int32(0) {
			} else {
				v37 = F__emscripten_memcpy_bulkmem(m, v28+v26, l1, l2)
				mBase = m.M
			}
			v39 = v26 + l2
			v41 = v28 + int32(-1)
			v42 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v41))))
			switch v42 & int32(7) {
			case 0:
				v46 = v39 << (uint(int32(3)) % 32)
				*(*uint8)(unsafe.Add(mBase, uint32(v41))) = uint8(v46)
			case 1:
				*(*uint8)(unsafe.Add(mBase, uint32(v28+int32(-3)))) = uint8(v39)
			case 2:
				*(*uint16)(unsafe.Add(mBase, uint32(v28+int32(-5)))) = uint16(v39)
			case 3:
				*(*int32)(unsafe.Add(mBase, uint32(v28+int32(-9)))) = v39
			case 4:
				*(*int64)(unsafe.Add(mBase, uint32(v28+int32(-17)))) = base.I64_extend_i32_u(v39)
			default:
			}
			v62 = int32(0)
			*(*uint8)(unsafe.Add(mBase, uint32(v28+v39))) = uint8(v62)
		}
		return v28
	}
}
func F_sdscatsds(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v10 int32
	_ = v10
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v67 int32
	_ = v67
	var v83 int32
	_ = v83
	v3 = int32(0)
	v10 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1+int32(-1)))))
	switch v10 & int32(7) {
	case 0:
		v27 = int32(base.Ui32(v10) >> (uint(int32(3)) % 32))
	case 1:
		v17 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1+int32(-3)))))
		v27 = v17
	case 2:
		v20 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1+int32(-5)))))
		v27 = v20
	case 3:
		v23 = *(*int32)(unsafe.Add(mBase, uint32(l1+int32(-9))))
		v27 = v23
	case 4:
		v26 = *(*int32)(unsafe.Add(mBase, uint32(l1+int32(-17))))
		v27 = v26
	default:
		v27 = v3
	}
	v30 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0+int32(-1)))))
	switch v30 & int32(7) {
	case 0:
		v47 = int32(base.Ui32(v30) >> (uint(int32(3)) % 32))
	case 1:
		v37 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0+int32(-3)))))
		v47 = v37
	case 2:
		v40 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0+int32(-5)))))
		v47 = v40
	case 3:
		v43 = *(*int32)(unsafe.Add(mBase, uint32(l0+int32(-9))))
		v47 = v43
	case 4:
		v46 = *(*int32)(unsafe.Add(mBase, uint32(l0+int32(-17))))
		v47 = v46
	default:
		v47 = v3
	}
	v49 = F__sdsMakeRoomFor(m, l0, v27, int32(1))
	mBase = m.M
	v52 = m.ExcPending
	if v52 != 0 {
		return int32(0)
	} else {
		if v49 == int32(0) {
		} else {
			if v27 == int32(0) {
			} else {
				v58 = F__emscripten_memcpy_bulkmem(m, v49+v47, l1, v27)
				mBase = m.M
			}
			v60 = v47 + v27
			v62 = v49 + int32(-1)
			v63 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v62))))
			switch v63 & int32(7) {
			case 0:
				v67 = v60 << (uint(int32(3)) % 32)
				*(*uint8)(unsafe.Add(mBase, uint32(v62))) = uint8(v67)
			case 1:
				*(*uint8)(unsafe.Add(mBase, uint32(v49+int32(-3)))) = uint8(v60)
			case 2:
				*(*uint16)(unsafe.Add(mBase, uint32(v49+int32(-5)))) = uint16(v60)
			case 3:
				*(*int32)(unsafe.Add(mBase, uint32(v49+int32(-9)))) = v60
			case 4:
				*(*int64)(unsafe.Add(mBase, uint32(v49+int32(-17)))) = base.I64_extend_i32_u(v60)
			default:
			}
			v83 = int32(0)
			*(*uint8)(unsafe.Add(mBase, uint32(v49+v60))) = uint8(v83)
		}
		return v49
	}
}
func F_sdsclear(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v12 int32
	_ = v12
	var v16 int32
	_ = v16
	var v26 int32
	_ = v26
	v4 = l0 + int32(-1)
	v5 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4))))
	switch v5 & int32(7) {
	case 0:
		v8 = int32(0)
		*(*uint8)(unsafe.Add(mBase, uint32(v4))) = uint8(v8)
	case 1:
		v12 = int32(0)
		*(*uint8)(unsafe.Add(mBase, uint32(l0+int32(-3)))) = uint8(v12)
	case 2:
		v16 = int32(0)
		*(*uint16)(unsafe.Add(mBase, uint32(l0+int32(-5)))) = uint16(v16)
	case 3:
		*(*int32)(unsafe.Add(mBase, uint32(l0+int32(-9)))) = int32(0)
	case 4:
		*(*int64)(unsafe.Add(mBase, uint32(l0+int32(-17)))) = int64(0)
	default:
	}
	v26 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l0))) = uint8(v26)
	return
}
func F_sdscmp(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v10 int32
	_ = v10
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v98 int32
	_ = v98
	var v113 int32
	_ = v113
	var v116 int32
	_ = v116
	v3 = int32(0)
	v10 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0+int32(-1)))))
	switch v10 & int32(7) {
	case 0:
		goto L6
	case 1:
		goto L5
	case 2:
		goto L4
	case 3:
		goto L3
	case 4:
		goto L2
	default:
		v27 = v3
		goto L1
	}
L1:
	;
	v30 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1+int32(-1)))))
	switch v30 & int32(7) {
	case 0:
		goto L12
	case 1:
		goto L11
	case 2:
		goto L10
	case 3:
		goto L9
	case 4:
		goto L8
	default:
		v47 = v3
		goto L7
	}
L2:
	;
	v26 = *(*int32)(unsafe.Add(mBase, uint32(l0+int32(-17))))
	v27 = v26
	goto L1
L3:
	;
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l0+int32(-9))))
	v27 = v23
	goto L1
L4:
	;
	v20 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0+int32(-5)))))
	v27 = v20
	goto L1
L5:
	;
	v17 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0+int32(-3)))))
	v27 = v17
	goto L1
L6:
	;
	v27 = int32(base.Ui32(v10) >> (uint(int32(3)) % 32))
	goto L1
L7:
	;
	v48 = base.B2i32(base.Ui32(v27) < base.Ui32(v47))
	if base.Ui32(v27) < base.Ui32(v47) {
		goto L13
	} else {
		goto L14
	}
L8:
	;
	v46 = *(*int32)(unsafe.Add(mBase, uint32(l1+int32(-17))))
	v47 = v46
	goto L7
L9:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(l1+int32(-9))))
	v47 = v43
	goto L7
L10:
	;
	v40 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1+int32(-5)))))
	v47 = v40
	goto L7
L11:
	;
	v37 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1+int32(-3)))))
	v47 = v37
	goto L7
L12:
	;
	v47 = int32(base.Ui32(v30) >> (uint(int32(3)) % 32))
	goto L7
L13:
	;
	v49 = v27
	goto L15
L14:
	;
	v49 = v47
	goto L15
L15:
	;
	if base.Ui32(v49) < base.Ui32(int32(4)) {
		v73 = l0
		v74 = l1
		v75 = v49
		goto L19
	} else {
		goto L20
	}
L16:
	;
	if v113 != 0 {
		goto L32
	} else {
		goto L33
	}
L17:
	;
	v113 = int32(0)
	goto L16
L18:
	;
	v85 = v80
	v86 = v81
	v87 = v82
	goto L28
L19:
	;
	if v75 == int32(0) {
		goto L17
	} else {
		goto L26
	}
L20:
	;
	if (l1|l0)&int32(3) != 0 {
		v80 = l0
		v81 = l1
		v82 = v49
		goto L18
	} else {
		goto L21
	}
L21:
	;
	v57 = l0
	v58 = l1
	v59 = v49
	goto L22
L22:
	;
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v57)))
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v58)))
	if v62 != v63 {
		v80 = v57
		v81 = v58
		v82 = v59
		goto L18
	} else {
		goto L24
	}
L23:
	;
	v73 = v68
	v74 = v66
	v75 = v70
	goto L19
L24:
	;
	v65 = int32(4)
	v66 = v58 + v65
	v68 = v57 + v65
	v70 = v59 + int32(-4)
	if base.Ui32(int32(3)) < base.Ui32(v70) {
		v57 = v68
		v58 = v66
		v59 = v70
		goto L22
	} else {
		goto L25
	}
L25:
	;
	goto L23
L26:
	;
	v80 = v73
	v81 = v74
	v82 = v75
	goto L18
L27:
	;
	v113 = v90 - v91
	goto L16
L28:
	;
	v90 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v85))))
	v91 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v86))))
	if v90 != v91 {
		goto L27
	} else {
		goto L30
	}
L30:
	;
	v93 = int32(1)
	v98 = v87 + int32(-1)
	if v98 == int32(0) {
		goto L17
	} else {
		goto L31
	}
L31:
	;
	v85 = v85 + v93
	v86 = v86 + v93
	v87 = v98
	goto L28
L32:
	;
	v116 = v113
	goto L34
L33:
	;
	v116 = base.B2i32(base.Ui32(v47) < base.Ui32(v27)) - v48
	goto L34
L34:
	;
	return v116
}
func F_sdsmapchars(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
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
	var v39 int32
	_ = v39
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v53 int32
	_ = v53
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v77 int32
	_ = v77
	v12 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0+int32(-1)))))
	switch v12 & int32(7) {
	case 0:
		goto L7
	case 1:
		goto L6
	case 2:
		goto L5
	case 3:
		goto L4
	case 4:
		goto L3
	default:
		goto L1
	}
L1:
	;
	return l0
L2:
	;
	if v29 == int32(0) {
		goto L1
	} else {
		goto L8
	}
L3:
	;
	v28 = *(*int32)(unsafe.Add(mBase, uint32(l0+int32(-17))))
	v29 = v28
	goto L2
L4:
	;
	v25 = *(*int32)(unsafe.Add(mBase, uint32(l0+int32(-9))))
	v29 = v25
	goto L2
L5:
	;
	v22 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0+int32(-5)))))
	v29 = v22
	goto L2
L6:
	;
	v19 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0+int32(-3)))))
	v29 = v19
	goto L2
L7:
	;
	v29 = int32(base.Ui32(v12) >> (uint(int32(3)) % 32))
	goto L2
L8:
	;
	v39 = int32(0)
	goto L9
L9:
	;
	if l3 == int32(0) {
		goto L11
	} else {
		goto L12
	}
L10:
	;
	goto L1
L11:
	;
	v77 = v39 + int32(1)
	if v77 != v29 {
		v39 = v77
		goto L9
	} else {
		goto L18
	}
L12:
	;
	v45 = l0 + v39
	v46 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v45))))
	v53 = int32(0)
	goto L13
L13:
	;
	v59 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1+v53))))
	if v46&int32(255) != v59 {
		goto L15
	} else {
		goto L16
	}
L14:
	;
	goto L11
L15:
	;
	v65 = v53 + int32(1)
	if v65 != l3 {
		v53 = v65
		goto L13
	} else {
		goto L17
	}
L16:
	;
	v62 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2+v53))))
	*(*uint8)(unsafe.Add(mBase, uint32(v45))) = uint8(v62)
	goto L11
L17:
	;
	goto L14
L18:
	;
	goto L10
}
func F_sdsnew(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v20 int32
	_ = v20
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v31 int32
	_ = v31
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v46 int32
	_ = v46
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	v2 = int32(0)
	if l0 == v2 {
		v61 = v2
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v63 = F__sdsnewlen(m, l0, v61, int32(0))
	mBase = m.M
	v66 = m.ExcPending
	if v66 != 0 {
		goto L19
	} else {
		goto L20
	}
L2:
	;
	if l0&int32(3) == int32(0) {
		v27 = l0
		goto L5
	} else {
		goto L6
	}
L3:
	;
	v61 = v60
	goto L1
L4:
	;
	v60 = v52 - l0
	goto L3
L5:
	;
	v31 = v27
	goto L13
L6:
	;
	v13 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v13 != 0 {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	v16 = l0
	goto L9
L8:
	;
	v60 = l0 - l0
	goto L3
L9:
	;
	v20 = v16 + int32(1)
	if v20&int32(3) == int32(0) {
		v27 = v20
		goto L5
	} else {
		goto L11
	}
L11:
	;
	v25 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20))))
	if v25 != 0 {
		v16 = v20
		goto L9
	} else {
		goto L12
	}
L12:
	;
	v52 = v20
	goto L4
L13:
	;
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v31)))
	v40 = int32(-2139062144)
	if (int32(16843008)-v37|v37)&v40 == v40 {
		v31 = v31 + int32(4)
		goto L13
	} else {
		goto L15
	}
L14:
	;
	v46 = v31
	goto L16
L15:
	;
	goto L14
L16:
	;
	v50 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v46))))
	if v50 != 0 {
		v46 = v46 + int32(1)
		goto L16
	} else {
		goto L18
	}
L17:
	;
	v52 = v46
	goto L4
L18:
	;
	goto L17
L19:
	;
	return int32(0)
L20:
	;
	return v63
}
func F_sdsnsplitargs(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	v5 = F_sdsnsplitargs_internal(m, l0, l0+l1, l2)
	v8 = m.ExcPending
	if v8 != 0 {
		return int32(0)
	} else {
		return v5
	}
}
func F_sdsparsearg(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v69 int32
	_ = v69
	var v103 int32
	_ = v103
	var v105 int32
	_ = v105
	var v127 int32
	_ = v127
	var v129 int32
	_ = v129
	var v136 int32
	_ = v136
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v157 int32
	_ = v157
	var v159 int32
	_ = v159
	var v165 int32
	_ = v165
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v177 int32
	_ = v177
	var v183 int32
	_ = v183
	var v187 int32
	_ = v187
	var v190 int32
	_ = v190
	var v192 int32
	_ = v192
	var v194 int32
	_ = v194
	var v198 int32
	_ = v198
	var v202 int32
	_ = v202
	var v204 int32
	_ = v204
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v221 int32
	_ = v221
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v232 int32
	_ = v232
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v245 int32
	_ = v245
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v255 int32
	_ = v255
	var v257 int32
	_ = v257
	var v258 int32
	_ = v258
	var v262 int32
	_ = v262
	var v263 int32
	_ = v263
	var v264 int32
	_ = v264
	var v265 int32
	_ = v265
	var v268 int32
	_ = v268
	var v281 int32
	_ = v281
	if l1 == int32(0) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v20 = int32(0)
	v21 = base.B2i32(l1 == v20)
	v27 = l3
	v29 = v20
	v30 = v20
	v31 = l0
	goto L4
L2:
	;
	if base.Ui32(l0) < base.Ui32(l1) {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	return int32(0)
L4:
	;
	if v30 == int32(0) {
		goto L12
	} else {
		goto L13
	}
L5:
	;
	return v265 - l0
L6:
	;
	if l1 == int32(0) {
		v281 = v268
		goto L88
	} else {
		goto L89
	}
L7:
	;
	v262 = v253
	v263 = v254
	v264 = v255
	v265 = v257
	v268 = v258
	goto L6
L8:
	;
	if l2 == int32(0) {
		goto L84
	} else {
		goto L85
	}
L9:
	;
	if v232 == int32(-1) {
		v253 = v27
		v254 = v225
		v255 = v226
		v257 = v228
		v258 = v229
		goto L7
	} else {
		goto L83
	}
L10:
	;
	v221 = int32(1)
	v225 = v29
	v226 = v221
	v228 = v31 + v221
	v229 = int32(0)
	v232 = base.I32_extend8_s(v40)
	goto L9
L11:
	;
	v235 = v29
	v236 = v147
	v238 = v136
	v239 = v146
	v240 = int32(13)
	goto L8
L12:
	;
	v159 = int32(*(*int8)(unsafe.Add(mBase, uint32(v31))))
	if v29 == int32(0) {
		goto L69
	} else {
		goto L70
	}
L13:
	;
	v40 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v31))))
	if v40 != 0 {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v44 = v31 + int32(4)
	if v21|base.B2i32(base.Ui32(v44) < base.Ui32(l1)) != int32(1) {
		goto L16
	} else {
		goto L17
	}
L15:
	;
	return int32(0)
L16:
	;
	v136 = v31 + int32(2)
	if v21|base.B2i32(base.Ui32(v136) < base.Ui32(l1)) != int32(1) {
		goto L60
	} else {
		goto L61
	}
L17:
	;
	if v40 != int32(92) {
		goto L16
	} else {
		goto L18
	}
L18:
	;
	v51 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v31)+1)))
	if v51 != int32(120) {
		goto L16
	} else {
		goto L19
	}
L19:
	;
	v54 = int32(*(*int8)(unsafe.Add(mBase, uint32(v31)+2)))
	if base.Ui32(int32(245)) < base.Ui32((v54+int32(-58))&int32(255)) {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	v69 = int32(*(*int8)(unsafe.Add(mBase, uint32(v31)+3)))
	if base.Ui32(int32(245)) < base.Ui32((v69+int32(-58))&int32(255)) {
		goto L23
	} else {
		goto L24
	}
L21:
	;
	if base.Ui32((v54&int32(-33)+int32(-71))&int32(255)) < base.Ui32(int32(250)) {
		goto L16
	} else {
		goto L22
	}
L22:
	;
	goto L20
L23:
	;
	switch v54 + int32(-49) {
	case 0:
		v103 = int32(1)
		goto L27
	case 1:
		goto L42
	case 2:
		goto L41
	case 3:
		goto L40
	case 4:
		goto L39
	case 5:
		goto L38
	case 6:
		goto L37
	case 7:
		goto L36
	case 8:
		goto L35
	default:
		goto L28
	case 16, 48:
		goto L34
	case 17, 49:
		goto L33
	case 18, 50:
		goto L32
	case 19, 51:
		goto L31
	case 20, 52:
		goto L30
	case 21, 53:
		goto L29
	}
L24:
	;
	if base.Ui32((v69&int32(-33)+int32(-71))&int32(255)) < base.Ui32(int32(250)) {
		goto L16
	} else {
		goto L25
	}
L25:
	;
	goto L23
L26:
	;
	switch v69 + int32(-49) {
	case 0:
		v127 = int32(1)
		goto L44
	case 1:
		goto L59
	case 2:
		goto L58
	case 3:
		goto L57
	case 4:
		goto L56
	case 5:
		goto L55
	case 6:
		goto L54
	case 7:
		goto L53
	case 8:
		goto L52
	default:
		goto L45
	case 16, 48:
		goto L51
	case 17, 49:
		goto L50
	case 18, 50:
		goto L49
	case 19, 51:
		goto L48
	case 20, 52:
		goto L47
	case 21, 53:
		goto L46
	}
L27:
	;
	v105 = v103
	goto L26
L28:
	;
	v103 = int32(0)
	goto L27
L29:
	;
	v105 = int32(15)
	goto L26
L30:
	;
	v105 = int32(14)
	goto L26
L31:
	;
	v105 = int32(13)
	goto L26
L32:
	;
	v105 = int32(12)
	goto L26
L33:
	;
	v105 = int32(11)
	goto L26
L34:
	;
	v105 = int32(10)
	goto L26
L35:
	;
	v105 = int32(9)
	goto L26
L36:
	;
	v105 = int32(8)
	goto L26
L37:
	;
	v105 = int32(7)
	goto L26
L38:
	;
	v105 = int32(6)
	goto L26
L39:
	;
	v105 = int32(5)
	goto L26
L40:
	;
	v105 = int32(4)
	goto L26
L41:
	;
	v105 = int32(3)
	goto L26
L42:
	;
	v105 = int32(2)
	goto L26
L43:
	;
	v235 = v29
	v236 = int32(1)
	v238 = v44
	v239 = int32(0)
	v240 = v105<<(uint(int32(4))%32) | v129
	goto L8
L44:
	;
	v129 = v127
	goto L43
L45:
	;
	v127 = int32(0)
	goto L44
L46:
	;
	v129 = int32(15)
	goto L43
L47:
	;
	v129 = int32(14)
	goto L43
L48:
	;
	v129 = int32(13)
	goto L43
L49:
	;
	v129 = int32(12)
	goto L43
L50:
	;
	v129 = int32(11)
	goto L43
L51:
	;
	v129 = int32(10)
	goto L43
L52:
	;
	v129 = int32(9)
	goto L43
L53:
	;
	v129 = int32(8)
	goto L43
L54:
	;
	v129 = int32(7)
	goto L43
L55:
	;
	v129 = int32(6)
	goto L43
L56:
	;
	v129 = int32(5)
	goto L43
L57:
	;
	v129 = int32(4)
	goto L43
L58:
	;
	v129 = int32(3)
	goto L43
L59:
	;
	v129 = int32(2)
	goto L43
L60:
	;
	if v40 != int32(34) {
		goto L10
	} else {
		goto L68
	}
L61:
	;
	if v40 != int32(92) {
		goto L60
	} else {
		goto L62
	}
L62:
	;
	v143 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v31)+1)))
	v144 = base.I32_extend8_s(v143)
	v146 = int32(0)
	v147 = int32(1)
	switch v143 + int32(-97) {
	case 0:
		goto L63
	case 1:
		goto L64
	case 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 14, 15, 16, 18:
		v225 = v29
		v226 = v147
		v228 = v136
		v229 = v146
		v232 = v144
		goto L9
	case 13:
		v235 = v29
		v236 = v147
		v238 = v136
		v239 = v146
		v240 = int32(10)
		goto L8
	case 17:
		goto L11
	case 19:
		goto L65
	default:
		goto L66
	}
L63:
	;
	v235 = v29
	v236 = v147
	v238 = v136
	v239 = v146
	v240 = int32(7)
	goto L8
L64:
	;
	v235 = v29
	v236 = v147
	v238 = v136
	v239 = v146
	v240 = int32(8)
	goto L8
L65:
	;
	v235 = v29
	v236 = v147
	v238 = v136
	v239 = v146
	v240 = int32(9)
	goto L8
L66:
	;
	if v143 != 0 {
		v225 = v29
		v226 = v147
		v228 = v136
		v229 = v146
		v232 = v144
		goto L9
	} else {
		goto L67
	}
L67:
	;
	goto L10
L68:
	;
	v157 = int32(0)
	v262 = v27
	v263 = v29
	v264 = v157
	v265 = v31 + int32(1)
	v268 = v157
	goto L6
L69:
	;
	v192 = int32(-1)
	v194 = int32(0)
	switch v159 & int32(255) {
	case 0, 9, 10, 13, 32:
		v207 = v194
		v208 = v194
		v209 = int32(1)
		v210 = v192
		goto L79
	default:
		goto L80
	case 34:
		goto L82
	case 39:
		goto L81
	}
L70:
	;
	if v159 != 0 {
		goto L71
	} else {
		goto L72
	}
L71:
	;
	v165 = v31 + int32(2)
	if v21|base.B2i32(base.Ui32(v165) < base.Ui32(l1)) != int32(1) {
		goto L74
	} else {
		goto L75
	}
L72:
	;
	return int32(0)
L73:
	;
	v187 = int32(1)
	v190 = int32(0)
	v225 = v187
	v226 = v190
	v228 = v31 + v187
	v229 = v190
	v232 = v159
	goto L9
L74:
	;
	if v159 != int32(39) {
		goto L73
	} else {
		goto L78
	}
L75:
	;
	if v159 != int32(92) {
		goto L74
	} else {
		goto L76
	}
L76:
	;
	v172 = int32(39)
	v173 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v31)+1)))
	if v173 != v172 {
		goto L73
	} else {
		goto L77
	}
L77:
	;
	v177 = int32(0)
	v235 = int32(1)
	v236 = v177
	v238 = v165
	v239 = v177
	v240 = v172
	goto L8
L78:
	;
	v183 = int32(0)
	v262 = v27
	v263 = v183
	v264 = v183
	v265 = v31 + int32(1)
	v268 = v183
	goto L6
L79:
	;
	v225 = v207
	v226 = v208
	v228 = v31 + base.B2i32(v159 != int32(0))
	v229 = v209
	v232 = v210
	goto L9
L80:
	;
	v204 = int32(0)
	v207 = v204
	v208 = v204
	v209 = v204
	v210 = v159
	goto L79
L81:
	;
	v202 = int32(0)
	v207 = int32(1)
	v208 = v202
	v209 = v202
	v210 = v192
	goto L79
L82:
	;
	v198 = int32(0)
	v207 = v198
	v208 = int32(1)
	v209 = v198
	v210 = v192
	goto L79
L83:
	;
	v235 = v225
	v236 = v226
	v238 = v228
	v239 = v229
	v240 = v232
	goto L8
L84:
	;
	if v27 != 0 {
		goto L86
	} else {
		goto L87
	}
L85:
	;
	v245 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v245 + int32(1)
	goto L84
L86:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v27))) = uint8(v240)
	v253 = v27 + int32(1)
	v254 = v235
	v255 = v236
	v257 = v238
	v258 = v239
	goto L7
L87:
	;
	v253 = int32(0)
	v254 = v235
	v255 = v236
	v257 = v238
	v258 = v239
	goto L7
L88:
	;
	if v281 == int32(0) {
		v27 = v262
		v29 = v263
		v30 = v264
		v31 = v265
		goto L4
	} else {
		goto L92
	}
L89:
	;
	if base.Ui32(v265) < base.Ui32(l1) {
		v281 = v268
		goto L88
	} else {
		goto L90
	}
L90:
	;
	if v264|v263 == int32(0) {
		v281 = int32(1)
		goto L88
	} else {
		goto L91
	}
L91:
	;
	return int32(0)
L92:
	;
	goto L5
}
func F_sdsrange(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
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
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v75 int32
	_ = v75
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v105 int32
	_ = v105
	var v116 int32
	_ = v116
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v125 int32
	_ = v125
	var v131 int32
	_ = v131
	var v138 int32
	_ = v138
	var v142 int32
	_ = v142
	var v145 int32
	_ = v145
	var v151 int32
	_ = v151
	var v158 int32
	_ = v158
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
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v179 int32
	_ = v179
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v184 int32
	_ = v184
	var v186 int32
	_ = v186
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v197 int32
	_ = v197
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
	var v221 int32
	_ = v221
	var v224 int32
	_ = v224
	v11 = l0 + int32(-1)
	v12 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11))))
	v14 = v12 & int32(7)
	switch v14 {
	case 0:
		goto L7
	case 1:
		goto L6
	case 2:
		goto L5
	case 3:
		goto L4
	case 4:
		goto L3
	default:
		goto L1
	}
L1:
	;
	return
L2:
	;
	if v29 == int32(0) {
		goto L1
	} else {
		goto L8
	}
L3:
	;
	v28 = *(*int32)(unsafe.Add(mBase, uint32(l0+int32(-17))))
	v29 = v28
	goto L2
L4:
	;
	v25 = *(*int32)(unsafe.Add(mBase, uint32(l0+int32(-9))))
	v29 = v25
	goto L2
L5:
	;
	v22 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0+int32(-5)))))
	v29 = v22
	goto L2
L6:
	;
	v19 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0+int32(-3)))))
	v29 = v19
	goto L2
L7:
	;
	v29 = int32(base.Ui32(v12) >> (uint(int32(3)) % 32))
	goto L2
L8:
	;
	v32 = int32(31)
	v35 = l2>>(uint(v32)%32)&v29 + l2
	v39 = l1>>(uint(v32)%32)&v29 + l1
	v42 = v35 - v39 + int32(1)
	switch v14 {
	default:
		goto L14
	case 1:
		goto L13
	case 2:
		goto L12
	case 3:
		goto L11
	case 4:
		goto L10
	}
L9:
	;
	v58 = int32(0)
	v60 = base.B2i32(base.Ui32(v39) < base.Ui32(v57))
	if base.Ui32(v39) < base.Ui32(v57) {
		goto L16
	} else {
		goto L17
	}
L10:
	;
	v56 = *(*int32)(unsafe.Add(mBase, uint32(l0+int32(-17))))
	v57 = v56
	goto L9
L11:
	;
	v53 = *(*int32)(unsafe.Add(mBase, uint32(l0+int32(-9))))
	v57 = v53
	goto L9
L12:
	;
	v50 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0+int32(-5)))))
	v57 = v50
	goto L9
L13:
	;
	v47 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0+int32(-3)))))
	v57 = v47
	goto L9
L14:
	;
	v57 = int32(base.Ui32(v12) >> (uint(int32(3)) % 32))
	goto L9
L15:
	;
	v221 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l0+v68))) = uint8(v221)
	switch v14 {
	default:
		goto L74
	case 1:
		goto L73
	case 2:
		goto L72
	case 3:
		goto L71
	case 4:
		goto L70
	}
L16:
	;
	v61 = v39
	goto L18
L17:
	;
	v61 = v58
	goto L18
L18:
	;
	v62 = v57 - v61
	if base.Ui32(v42) < base.Ui32(v62) {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	v64 = v42
	goto L21
L20:
	;
	v64 = v62
	goto L21
L21:
	;
	if v35 < v39 {
		goto L22
	} else {
		goto L23
	}
L22:
	;
	v66 = v58
	goto L24
L23:
	;
	v66 = v64
	goto L24
L24:
	;
	if base.Ui32(v39) < base.Ui32(v57) {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	v68 = v66
	goto L27
L26:
	;
	v68 = int32(0)
	goto L27
L27:
	;
	if v68 == int32(0) {
		goto L15
	} else {
		goto L28
	}
L28:
	;
	v71 = l0 + v61
	if l0 == v71 {
		goto L30
	} else {
		goto L31
	}
L29:
	;
	goto L15
L30:
	;
	goto L29
L31:
	;
	v75 = v68 + l0
	if base.Ui32(int32(0)-v68<<(uint(int32(1))%32)) < base.Ui32(v71-v75) {
		goto L32
	} else {
		goto L33
	}
L32:
	;
	v85 = (v71 ^ l0) & int32(3)
	if base.Ui32(v71) <= base.Ui32(l0) {
		goto L36
	} else {
		goto L37
	}
L33:
	;
	v82 = F___memcpy(m, l0, v71, v68)
	mBase = m.M
	goto L29
L34:
	;
	if v191 == int32(0) {
		goto L30
	} else {
		goto L66
	}
L35:
	;
	if base.Ui32(v169) <= base.Ui32(int32(3)) {
		v190 = v168
		v191 = v169
		v192 = v170
		goto L34
	} else {
		goto L62
	}
L36:
	;
	if v85 != 0 {
		v151 = v68
		goto L46
	} else {
		goto L47
	}
L37:
	;
	if v85 == int32(0) {
		goto L38
	} else {
		goto L39
	}
L38:
	;
	if l0&int32(3) != 0 {
		goto L40
	} else {
		goto L41
	}
L39:
	;
	v190 = v71
	v191 = v68
	v192 = l0
	goto L34
L40:
	;
	v92 = v71
	v93 = v68
	v94 = l0
	goto L42
L41:
	;
	v168 = v71
	v169 = v68
	v170 = l0
	goto L35
L42:
	;
	if v93 == int32(0) {
		goto L30
	} else {
		goto L44
	}
L44:
	;
	v98 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v92))))
	*(*uint8)(unsafe.Add(mBase, uint32(v94))) = uint8(v98)
	v100 = int32(1)
	v101 = v92 + v100
	v103 = v93 + int32(-1)
	v105 = v94 + v100
	if v105&int32(3) == int32(0) {
		v168 = v101
		v169 = v103
		v170 = v105
		goto L35
	} else {
		goto L45
	}
L45:
	;
	v92 = v101
	v93 = v103
	v94 = v105
	goto L42
L46:
	;
	if v151 == int32(0) {
		goto L30
	} else {
		goto L58
	}
L47:
	;
	if v75&int32(3) == int32(0) {
		v131 = v68
		goto L48
	} else {
		goto L49
	}
L48:
	;
	if base.Ui32(v131) <= base.Ui32(int32(3)) {
		v151 = v131
		goto L46
	} else {
		goto L54
	}
L49:
	;
	v116 = v68
	goto L50
L50:
	;
	if v116 == int32(0) {
		goto L30
	} else {
		goto L52
	}
L51:
	;
	v131 = v122
	goto L48
L52:
	;
	v122 = v116 + int32(-1)
	v123 = l0 + v122
	v125 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v71+v122))))
	*(*uint8)(unsafe.Add(mBase, uint32(v123))) = uint8(v125)
	if v123&int32(3) != 0 {
		v116 = v122
		goto L50
	} else {
		goto L53
	}
L53:
	;
	goto L51
L54:
	;
	v138 = v131
	goto L55
L55:
	;
	v142 = v138 + int32(-4)
	v145 = *(*int32)(unsafe.Add(mBase, uint32(v71+v142)))
	*(*int32)(unsafe.Add(mBase, uint32(l0+v142))) = v145
	if base.Ui32(int32(3)) < base.Ui32(v142) {
		v138 = v142
		goto L55
	} else {
		goto L57
	}
L56:
	;
	v151 = v142
	goto L46
L57:
	;
	goto L56
L58:
	;
	v158 = v151
	goto L59
L59:
	;
	v162 = v158 + int32(-1)
	v165 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v71+v162))))
	*(*uint8)(unsafe.Add(mBase, uint32(l0+v162))) = uint8(v165)
	if v162 != 0 {
		v158 = v162
		goto L59
	} else {
		goto L61
	}
L61:
	;
	goto L30
L62:
	;
	v175 = v168
	v176 = v169
	v177 = v170
	goto L63
L63:
	;
	v179 = *(*int32)(unsafe.Add(mBase, uint32(v175)))
	*(*int32)(unsafe.Add(mBase, uint32(v177))) = v179
	v181 = int32(4)
	v182 = v175 + v181
	v184 = v177 + v181
	v186 = v176 + int32(-4)
	if base.Ui32(int32(3)) < base.Ui32(v186) {
		v175 = v182
		v176 = v186
		v177 = v184
		goto L63
	} else {
		goto L65
	}
L64:
	;
	v190 = v182
	v191 = v186
	v192 = v184
	goto L34
L65:
	;
	goto L64
L66:
	;
	v197 = v190
	v198 = v191
	v199 = v192
	goto L67
L67:
	;
	v201 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v197))))
	*(*uint8)(unsafe.Add(mBase, uint32(v199))) = uint8(v201)
	v203 = int32(1)
	v208 = v198 + int32(-1)
	if v208 != 0 {
		v197 = v197 + v203
		v198 = v208
		v199 = v199 + v203
		goto L67
	} else {
		goto L69
	}
L68:
	;
	goto L30
L69:
	;
	goto L68
L70:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l0+int32(-17)))) = base.I64_extend_i32_u(v68)
	goto L1
L71:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0+int32(-9)))) = v68
	return
L72:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(l0+int32(-5)))) = uint16(v68)
	return
L73:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(l0+int32(-3)))) = uint8(v68)
	return
L74:
	;
	v224 = v68 << (uint(int32(3)) % 32)
	*(*uint8)(unsafe.Add(mBase, uint32(v11))) = uint8(v224)
	return
}
func F_sdstemplate(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v50 int32
	_ = v50
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v73 int32
	_ = v73
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	var v86 int32
	_ = v86
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v108 int32
	_ = v108
	var v124 int32
	_ = v124
	var v128 int32
	_ = v128
	var v132 int32
	_ = v132
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v142 int32
	_ = v142
	var v146 int32
	_ = v146
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v162 int32
	_ = v162
	var v165 int32
	_ = v165
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v187 int32
	_ = v187
	var v190 int32
	_ = v190
	var v193 int32
	_ = v193
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v204 int32
	_ = v204
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v220 int32
	_ = v220
	var v223 int32
	_ = v223
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v245 int32
	_ = v245
	var v248 int32
	_ = v248
	var v251 int32
	_ = v251
	var v254 int32
	_ = v254
	var v255 int32
	_ = v255
	var v257 int32
	_ = v257
	var v258 int32
	_ = v258
	var v262 int32
	_ = v262
	var v273 int32
	_ = v273
	var v278 int32
	_ = v278
	var v279 int32
	_ = v279
	var v281 int32
	_ = v281
	var v284 int32
	_ = v284
	var v299 int32
	_ = v299
	var v300 int32
	_ = v300
	var v306 int32
	_ = v306
	var v309 int32
	_ = v309
	var v312 int32
	_ = v312
	var v315 int32
	_ = v315
	var v316 int32
	_ = v316
	var v318 int32
	_ = v318
	var v319 int32
	_ = v319
	var v322 int32
	_ = v322
	var v326 int32
	_ = v326
	var v327 int32
	_ = v327
	var v328 int32
	_ = v328
	var v330 int32
	_ = v330
	var v337 int32
	_ = v337
	var v343 int32
	_ = v343
	v13 = m.G0
	v15 = v13 - int32(16)
	m.G0 = v15
	v20 = F_zmalloc_usable(m, int32(4), v15+int32(12))
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	v42 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v42 == int32(0) {
		v343 = v41
		goto L13
	} else {
		goto L14
	}
L2:
	;
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v15)+12))
	v29 = v25 + int32(-4)
	if base.Ui32(v29) < base.Ui32(int32(65531)) {
		goto L6
	} else {
		goto L7
	}
L3:
	;
	return int32(0)
L4:
	;
	if v20 != 0 {
		goto L2
	} else {
		goto L5
	}
L5:
	;
	v41 = int32(0)
	goto L1
L6:
	;
	v32 = int32(2)
	goto L8
L7:
	;
	v32 = int32(3)
	goto L8
L8:
	;
	if base.Ui32(int32(255)) < base.Ui32(v29) {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	v36 = v32
	goto L11
L10:
	;
	v36 = int32(1)
	goto L11
L11:
	;
	v39 = F_sdswrite(m, v20, v25, v36, int32(_a139), int32(0))
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
		goto L3
	} else {
		goto L12
	}
L12:
	;
	v41 = v39
	goto L1
L13:
	;
	m.G0 = v15 + int32(16)
	return v343
L14:
	;
	v45 = l0
	v50 = v41
	goto L15
L15:
	;
	v57 = int32(123)
	v58 = F___strchrnul(m, v45, v57)
	mBase = m.M
	v60 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v58))))
	if v60 == v57 {
		goto L19
	} else {
		goto L20
	}
L16:
	;
	v343 = v330
	goto L13
L17:
	;
	if base.Ui32(v45) < base.Ui32(v64) {
		goto L25
	} else {
		goto L26
	}
L18:
	;
	if v64 != 0 {
		goto L17
	} else {
		goto L22
	}
L19:
	;
	v64 = v58
	goto L21
L20:
	;
	v64 = int32(0)
	goto L21
L21:
	;
	goto L18
L22:
	;
	v65 = F_sdscat(m, v50, v45)
	mBase = m.M
	v66 = m.ExcPending
	if v66 != 0 {
		goto L3
	} else {
		goto L23
	}
L23:
	;
	v343 = v65
	goto L13
L24:
	;
	v132 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v64)+1)))
	if v132 == int32(123) {
		goto L45
	} else {
		goto L46
	}
L25:
	;
	v68 = v64 - v45
	v69 = int32(0)
	v73 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v50+int32(-1)))))
	switch v73 & int32(7) {
	case 0:
		goto L32
	case 1:
		goto L31
	case 2:
		goto L30
	case 3:
		goto L29
	case 4:
		goto L28
	default:
		v90 = v69
		goto L27
	}
L26:
	;
	v128 = v50
	goto L24
L27:
	;
	v92 = F__sdsMakeRoomFor(m, v50, v68, int32(1))
	mBase = m.M
	v93 = m.ExcPending
	if v93 != 0 {
		goto L3
	} else {
		goto L33
	}
L28:
	;
	v89 = *(*int32)(unsafe.Add(mBase, uint32(v50+int32(-17))))
	v90 = v89
	goto L27
L29:
	;
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v50+int32(-9))))
	v90 = v86
	goto L27
L30:
	;
	v83 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v50+int32(-5)))))
	v90 = v83
	goto L27
L31:
	;
	v80 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v50+int32(-3)))))
	v90 = v80
	goto L27
L32:
	;
	v90 = int32(base.Ui32(v73) >> (uint(int32(3)) % 32))
	goto L27
L33:
	;
	if v92 == int32(0) {
		v128 = v69
		goto L24
	} else {
		goto L34
	}
L34:
	;
	if v68 == int32(0) {
		goto L36
	} else {
		goto L37
	}
L35:
	;
	v101 = v90 + v68
	v103 = v92 + int32(-1)
	v104 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v103))))
	switch v104 & int32(7) {
	case 0:
		goto L43
	case 1:
		goto L42
	case 2:
		goto L41
	case 3:
		goto L40
	case 4:
		goto L39
	default:
		goto L38
	}
L36:
	;
	goto L35
L37:
	;
	v99 = F__emscripten_memcpy_bulkmem(m, v92+v90, v45, v68)
	mBase = m.M
	goto L36
L38:
	;
	v124 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v92+v101))) = uint8(v124)
	v128 = v92
	goto L24
L39:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v92+int32(-17)))) = base.I64_extend_i32_u(v101)
	goto L38
L40:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v92+int32(-9)))) = v101
	goto L38
L41:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v92+int32(-5)))) = uint16(v101)
	goto L38
L42:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v92+int32(-3)))) = uint8(v101)
	goto L38
L43:
	;
	v108 = v101 << (uint(int32(3)) % 32)
	*(*uint8)(unsafe.Add(mBase, uint32(v103))) = uint8(v108)
	goto L38
L44:
	;
	v337 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v328))))
	if v337 != 0 {
		v45 = v328
		v50 = v330
		goto L15
	} else {
		goto L107
	}
L45:
	;
	v326 = F_sdscat(m, v128, int32(_a1331))
	mBase = m.M
	v327 = m.ExcPending
	if v327 != 0 {
		goto L3
	} else {
		goto L106
	}
L46:
	;
	if v132 == int32(0) {
		goto L47
	} else {
		goto L48
	}
L47:
	;
	v273 = int32(0)
	if v128 == v273 {
		v343 = v273
		goto L13
	} else {
		goto L90
	}
L48:
	;
	v138 = v64 + int32(1)
	v139 = int32(125)
	v140 = F___strchrnul(m, v138, v139)
	mBase = m.M
	v142 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v140))))
	if v142 == v139 {
		goto L50
	} else {
		goto L51
	}
L49:
	;
	if v146 == int32(0) {
		goto L47
	} else {
		goto L53
	}
L50:
	;
	v146 = v140
	goto L52
L51:
	;
	v146 = int32(0)
	goto L52
L52:
	;
	goto L49
L53:
	;
	v151 = F__sdsnewlen(m, v138, v146-v138, int32(0))
	mBase = m.M
	v152 = m.ExcPending
	if v152 != 0 {
		goto L3
	} else {
		goto L54
	}
L54:
	;
	v153 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, v151, l2)
	mBase = m.M
	v154 = m.ExcPending
	if v154 != 0 {
		goto L3
	} else {
		goto L55
	}
L55:
	;
	if v151 == int32(0) {
		goto L56
	} else {
		goto L57
	}
L56:
	;
	if v153 == int32(0) {
		goto L47
	} else {
		goto L73
	}
L57:
	;
	v159 = v151 + int32(-1)
	v160 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v159))))
	v162 = v160 & int32(7)
	switch v162 {
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
		v180 = v151
		v181 = int32(1)
		goto L59
	}
L58:
	;
	F_zfree_with_size(m, v200, v199)
	mBase = m.M
	v204 = m.ExcPending
	if v204 != 0 {
		goto L3
	} else {
		goto L72
	}
L59:
	;
	switch v162 {
	case 0:
		goto L71
	case 1:
		goto L70
	case 2:
		goto L69
	case 3:
		goto L68
	case 4:
		goto L67
	default:
		v197 = int32(0)
		goto L66
	}
L60:
	;
	v180 = v151 + int32(-17)
	v181 = int32(18)
	goto L59
L61:
	;
	v180 = v151 + int32(-9)
	v181 = int32(10)
	goto L59
L62:
	;
	v180 = v151 + int32(-5)
	v181 = int32(6)
	goto L59
L63:
	;
	v180 = v151 + int32(-3)
	v181 = int32(4)
	goto L59
L64:
	;
	v165 = *(*int32)(unsafe.Add(mBase, uint32(v151+int32(-9))))
	goto L65
L65:
	;
	v199 = v165 & int32(2147483647)
	v200 = v159
	goto L58
L66:
	;
	v199 = v197 + v181
	v200 = v180
	goto L58
L67:
	;
	v196 = *(*int32)(unsafe.Add(mBase, uint32(v151+int32(-9))))
	v197 = v196
	goto L66
L68:
	;
	v193 = *(*int32)(unsafe.Add(mBase, uint32(v151+int32(-5))))
	v197 = v193
	goto L66
L69:
	;
	v190 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v151+int32(-3)))))
	v197 = v190
	goto L66
L70:
	;
	v187 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v151+int32(-2)))))
	v197 = v187
	goto L66
L71:
	;
	v197 = int32(base.Ui32(v160) >> (uint(int32(3)) % 32))
	goto L66
L72:
	;
	goto L56
L73:
	;
	v213 = F_sdscat(m, v128, v153)
	mBase = m.M
	v214 = m.ExcPending
	if v214 != 0 {
		goto L3
	} else {
		goto L74
	}
L74:
	;
	v217 = v153 + int32(-1)
	v218 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v217))))
	v220 = v218 & int32(7)
	switch v220 {
	case 0:
		goto L81
	case 1:
		goto L80
	case 2:
		goto L79
	case 3:
		goto L78
	case 4:
		goto L77
	default:
		v238 = v153
		v239 = int32(1)
		goto L76
	}
L75:
	;
	F_zfree_with_size(m, v258, v257)
	mBase = m.M
	v262 = m.ExcPending
	if v262 != 0 {
		goto L3
	} else {
		goto L89
	}
L76:
	;
	switch v220 {
	case 0:
		goto L88
	case 1:
		goto L87
	case 2:
		goto L86
	case 3:
		goto L85
	case 4:
		goto L84
	default:
		v255 = int32(0)
		goto L83
	}
L77:
	;
	v238 = v153 + int32(-17)
	v239 = int32(18)
	goto L76
L78:
	;
	v238 = v153 + int32(-9)
	v239 = int32(10)
	goto L76
L79:
	;
	v238 = v153 + int32(-5)
	v239 = int32(6)
	goto L76
L80:
	;
	v238 = v153 + int32(-3)
	v239 = int32(4)
	goto L76
L81:
	;
	v223 = *(*int32)(unsafe.Add(mBase, uint32(v153+int32(-9))))
	goto L82
L82:
	;
	v257 = v223 & int32(2147483647)
	v258 = v217
	goto L75
L83:
	;
	v257 = v255 + v239
	v258 = v238
	goto L75
L84:
	;
	v254 = *(*int32)(unsafe.Add(mBase, uint32(v153+int32(-9))))
	v255 = v254
	goto L83
L85:
	;
	v251 = *(*int32)(unsafe.Add(mBase, uint32(v153+int32(-5))))
	v255 = v251
	goto L83
L86:
	;
	v248 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v153+int32(-3)))))
	v255 = v248
	goto L83
L87:
	;
	v245 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v153+int32(-2)))))
	v255 = v245
	goto L83
L88:
	;
	v255 = int32(base.Ui32(v218) >> (uint(int32(3)) % 32))
	goto L83
L89:
	;
	v328 = v146 + int32(1)
	v330 = v213
	goto L44
L90:
	;
	v278 = v128 + int32(-1)
	v279 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v278))))
	v281 = v279 & int32(7)
	switch v281 {
	case 0:
		goto L97
	case 1:
		goto L96
	case 2:
		goto L95
	case 3:
		goto L94
	case 4:
		goto L93
	default:
		v299 = v128
		v300 = int32(1)
		goto L92
	}
L91:
	;
	F_zfree_with_size(m, v318, v319)
	mBase = m.M
	v322 = m.ExcPending
	if v322 != 0 {
		goto L3
	} else {
		goto L105
	}
L92:
	;
	switch v281 {
	case 0:
		goto L104
	case 1:
		goto L103
	case 2:
		goto L102
	case 3:
		goto L101
	case 4:
		goto L100
	default:
		v316 = int32(0)
		goto L99
	}
L93:
	;
	v299 = v128 + int32(-17)
	v300 = int32(18)
	goto L92
L94:
	;
	v299 = v128 + int32(-9)
	v300 = int32(10)
	goto L92
L95:
	;
	v299 = v128 + int32(-5)
	v300 = int32(6)
	goto L92
L96:
	;
	v299 = v128 + int32(-3)
	v300 = int32(4)
	goto L92
L97:
	;
	v284 = *(*int32)(unsafe.Add(mBase, uint32(v128+int32(-9))))
	goto L98
L98:
	;
	v318 = v278
	v319 = v284 & int32(2147483647)
	goto L91
L99:
	;
	v318 = v299
	v319 = v316 + v300
	goto L91
L100:
	;
	v315 = *(*int32)(unsafe.Add(mBase, uint32(v128+int32(-9))))
	v316 = v315
	goto L99
L101:
	;
	v312 = *(*int32)(unsafe.Add(mBase, uint32(v128+int32(-5))))
	v316 = v312
	goto L99
L102:
	;
	v309 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v128+int32(-3)))))
	v316 = v309
	goto L99
L103:
	;
	v306 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v128+int32(-2)))))
	v316 = v306
	goto L99
L104:
	;
	v316 = int32(base.Ui32(v279) >> (uint(int32(3)) % 32))
	goto L99
L105:
	;
	v343 = v273
	goto L13
L106:
	;
	v328 = v64 + int32(2)
	v330 = v326
	goto L44
L107:
	;
	goto L16
}
func F_sdstolower(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	v7 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0+int32(-1)))))
	switch v7 & int32(7) {
	case 0:
		v24 = int32(base.Ui32(v7) >> (uint(int32(3)) % 32))
		if v24 == int32(0) {
		} else {
			v29 = int32(0)
			for {
				v32 = l0 + v29
				v33 = int32(*(*int8)(unsafe.Add(mBase, uint32(v32))))
				if base.Ui32(v33+int32(-65)) < base.Ui32(int32(26)) {
					v40 = v33 | int32(32)
				} else {
					v40 = v33
				}
				*(*uint8)(unsafe.Add(mBase, uint32(v32))) = uint8(v40)
				v43 = v29 + int32(1)
				if v43 != v24 {
					v29 = v43
					continue
				} else {
					break
				}
				break
			}
		}
	case 1:
		v14 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0+int32(-3)))))
		v24 = v14
		if v24 == int32(0) {
		} else {
			v29 = int32(0)
			for {
				v32 = l0 + v29
				v33 = int32(*(*int8)(unsafe.Add(mBase, uint32(v32))))
				if base.Ui32(v33+int32(-65)) < base.Ui32(int32(26)) {
					v40 = v33 | int32(32)
				} else {
					v40 = v33
				}
				*(*uint8)(unsafe.Add(mBase, uint32(v32))) = uint8(v40)
				v43 = v29 + int32(1)
				if v43 != v24 {
					v29 = v43
					continue
				} else {
					break
				}
				break
			}
		}
	case 2:
		v17 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0+int32(-5)))))
		v24 = v17
		if v24 == int32(0) {
		} else {
			v29 = int32(0)
			for {
				v32 = l0 + v29
				v33 = int32(*(*int8)(unsafe.Add(mBase, uint32(v32))))
				if base.Ui32(v33+int32(-65)) < base.Ui32(int32(26)) {
					v40 = v33 | int32(32)
				} else {
					v40 = v33
				}
				*(*uint8)(unsafe.Add(mBase, uint32(v32))) = uint8(v40)
				v43 = v29 + int32(1)
				if v43 != v24 {
					v29 = v43
					continue
				} else {
					break
				}
				break
			}
		}
	case 3:
		v20 = *(*int32)(unsafe.Add(mBase, uint32(l0+int32(-9))))
		v24 = v20
		if v24 == int32(0) {
		} else {
			v29 = int32(0)
			for {
				v32 = l0 + v29
				v33 = int32(*(*int8)(unsafe.Add(mBase, uint32(v32))))
				if base.Ui32(v33+int32(-65)) < base.Ui32(int32(26)) {
					v40 = v33 | int32(32)
				} else {
					v40 = v33
				}
				*(*uint8)(unsafe.Add(mBase, uint32(v32))) = uint8(v40)
				v43 = v29 + int32(1)
				if v43 != v24 {
					v29 = v43
					continue
				} else {
					break
				}
				break
			}
		}
	case 4:
		v23 = *(*int32)(unsafe.Add(mBase, uint32(l0+int32(-17))))
		v24 = v23
		if v24 == int32(0) {
		} else {
			v29 = int32(0)
			for {
				v32 = l0 + v29
				v33 = int32(*(*int8)(unsafe.Add(mBase, uint32(v32))))
				if base.Ui32(v33+int32(-65)) < base.Ui32(int32(26)) {
					v40 = v33 | int32(32)
				} else {
					v40 = v33
				}
				*(*uint8)(unsafe.Add(mBase, uint32(v32))) = uint8(v40)
				v43 = v29 + int32(1)
				if v43 != v24 {
					v29 = v43
					continue
				} else {
					break
				}
				break
			}
		}
	default:
	}
	return
}
func F_sdstoupper(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	v7 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0+int32(-1)))))
	switch v7 & int32(7) {
	case 0:
		v24 = int32(base.Ui32(v7) >> (uint(int32(3)) % 32))
		if v24 == int32(0) {
		} else {
			v29 = int32(0)
			for {
				v32 = l0 + v29
				v33 = int32(*(*int8)(unsafe.Add(mBase, uint32(v32))))
				if base.Ui32(v33+int32(-97)) < base.Ui32(int32(26)) {
					v40 = v33 & int32(95)
				} else {
					v40 = v33
				}
				*(*uint8)(unsafe.Add(mBase, uint32(v32))) = uint8(v40)
				v43 = v29 + int32(1)
				if v43 != v24 {
					v29 = v43
					continue
				} else {
					break
				}
				break
			}
		}
	case 1:
		v14 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0+int32(-3)))))
		v24 = v14
		if v24 == int32(0) {
		} else {
			v29 = int32(0)
			for {
				v32 = l0 + v29
				v33 = int32(*(*int8)(unsafe.Add(mBase, uint32(v32))))
				if base.Ui32(v33+int32(-97)) < base.Ui32(int32(26)) {
					v40 = v33 & int32(95)
				} else {
					v40 = v33
				}
				*(*uint8)(unsafe.Add(mBase, uint32(v32))) = uint8(v40)
				v43 = v29 + int32(1)
				if v43 != v24 {
					v29 = v43
					continue
				} else {
					break
				}
				break
			}
		}
	case 2:
		v17 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0+int32(-5)))))
		v24 = v17
		if v24 == int32(0) {
		} else {
			v29 = int32(0)
			for {
				v32 = l0 + v29
				v33 = int32(*(*int8)(unsafe.Add(mBase, uint32(v32))))
				if base.Ui32(v33+int32(-97)) < base.Ui32(int32(26)) {
					v40 = v33 & int32(95)
				} else {
					v40 = v33
				}
				*(*uint8)(unsafe.Add(mBase, uint32(v32))) = uint8(v40)
				v43 = v29 + int32(1)
				if v43 != v24 {
					v29 = v43
					continue
				} else {
					break
				}
				break
			}
		}
	case 3:
		v20 = *(*int32)(unsafe.Add(mBase, uint32(l0+int32(-9))))
		v24 = v20
		if v24 == int32(0) {
		} else {
			v29 = int32(0)
			for {
				v32 = l0 + v29
				v33 = int32(*(*int8)(unsafe.Add(mBase, uint32(v32))))
				if base.Ui32(v33+int32(-97)) < base.Ui32(int32(26)) {
					v40 = v33 & int32(95)
				} else {
					v40 = v33
				}
				*(*uint8)(unsafe.Add(mBase, uint32(v32))) = uint8(v40)
				v43 = v29 + int32(1)
				if v43 != v24 {
					v29 = v43
					continue
				} else {
					break
				}
				break
			}
		}
	case 4:
		v23 = *(*int32)(unsafe.Add(mBase, uint32(l0+int32(-17))))
		v24 = v23
		if v24 == int32(0) {
		} else {
			v29 = int32(0)
			for {
				v32 = l0 + v29
				v33 = int32(*(*int8)(unsafe.Add(mBase, uint32(v32))))
				if base.Ui32(v33+int32(-97)) < base.Ui32(int32(26)) {
					v40 = v33 & int32(95)
				} else {
					v40 = v33
				}
				*(*uint8)(unsafe.Add(mBase, uint32(v32))) = uint8(v40)
				v43 = v29 + int32(1)
				if v43 != v24 {
					v29 = v43
					continue
				} else {
					break
				}
				break
			}
		}
	default:
	}
	return
}
func F_sdsupdatelen(m *base.Module, l0 int32) {
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
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v65 int32
	_ = v65
	if l0&int32(3) == int32(0) {
		v25 = l0
		goto L3
	} else {
		goto L4
	}
L1:
	;
	v60 = l0 + int32(-1)
	v61 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v60))))
	switch v61 & int32(7) {
	case 0:
		goto L22
	case 1:
		goto L21
	case 2:
		goto L20
	case 3:
		goto L19
	case 4:
		goto L18
	default:
		goto L17
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
	return
L18:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l0+int32(-17)))) = base.I64_extend_i32_u(v58)
	goto L17
L19:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0+int32(-9)))) = v58
	return
L20:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(l0+int32(-5)))) = uint16(v58)
	return
L21:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(l0+int32(-3)))) = uint8(v58)
	return
L22:
	;
	v65 = v58 << (uint(int32(3)) % 32)
	*(*uint8)(unsafe.Add(mBase, uint32(v60))) = uint8(v65)
	return
}
func F_sdswrite(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v10 int32
	_ = v10
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v72 int32
	_ = v72
	var v78 int32
	_ = v78
	var v82 int32
	_ = v82
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v99 int32
	_ = v99
	var v105 int32
	_ = v105
	var v111 int32
	_ = v111
	v5 = l4
	v10 = l2 & int32(7)
	if base.Ui32(int32(4)) < base.Ui32(v10) {
		v18 = int32(0)
	} else {
		v17 = *(*int32)(unsafe.Add(mBase, uint32(v10<<(uint(int32(2))%32))+uint32(_consts[602])))
		v18 = v17
	}
	if base.Ui32(l1) < base.Ui32(v5+v18+int32(1)) {
		F__serverAssert(m, int32(_a1322), int32(_a1321), int32(115))
		mBase = m.M
		v99 = m.ExcPending
		if v99 != 0 {
			return int32(0)
		} else {
			F_abort(m)
			mBase = m.M
			base.Wasm_trap_unreachable()
			for {
			}
		}
	} else {
		if base.Ui32(int32(4)) < base.Ui32(v10) {
			v31 = int32(0)
		} else {
			v30 = *(*int32)(unsafe.Add(mBase, uint32(v10<<(uint(int32(2))%32))+uint32(_consts[602])))
			v31 = v30
		}
		v32 = l0 + v31
		v35 = l1 + (v31 ^ int32(-1))
		switch l2 {
		case 0:
			v72 = v5 << (uint(int32(3)) % 32)
			*(*uint8)(unsafe.Add(mBase, uint32(v32+int32(-1)))) = uint8(v72)
			v78 = *(*int32)(unsafe.Add(mBase, _consts[72]))
			if l3 == v78 {
			} else {
				if l3 != 0 {
					if v5 == int32(0) {
					} else {
						if v5 == int32(0) {
						} else {
							v87 = F__emscripten_memcpy_bulkmem(m, v32, l3, v5)
							mBase = m.M
						}
					}
				} else {
					v82 = F__emscripten_memset_bulkmem(m, v32, base.I32_extend8_s(int32(0)), v5)
					mBase = m.M
				}
			}
			v90 = int32(0)
			*(*uint8)(unsafe.Add(mBase, uint32(v32+v5))) = uint8(v90)
			return v32
		case 1:
			*(*uint8)(unsafe.Add(mBase, uint32(v32+int32(-3)))) = uint8(v5)
			if base.Ui32(int32(256)) <= base.Ui32(v35) {
				F__serverAssert(m, int32(_a1323), int32(_a1321), int32(129))
				mBase = m.M
				v105 = m.ExcPending
				if v105 != 0 {
					return int32(0)
				} else {
					F_abort(m)
					mBase = m.M
					base.Wasm_trap_unreachable()
					for {
					}
				}
			} else {
				*(*uint8)(unsafe.Add(mBase, uint32(v32+int32(-2)))) = uint8(v35)
				v72 = int32(1)
				*(*uint8)(unsafe.Add(mBase, uint32(v32+int32(-1)))) = uint8(v72)
				v78 = *(*int32)(unsafe.Add(mBase, _consts[72]))
				if l3 == v78 {
				} else {
					if l3 != 0 {
						if v5 == int32(0) {
						} else {
							if v5 == int32(0) {
							} else {
								v87 = F__emscripten_memcpy_bulkmem(m, v32, l3, v5)
								mBase = m.M
							}
						}
					} else {
						v82 = F__emscripten_memset_bulkmem(m, v32, base.I32_extend8_s(int32(0)), v5)
						mBase = m.M
					}
				}
				v90 = int32(0)
				*(*uint8)(unsafe.Add(mBase, uint32(v32+v5))) = uint8(v90)
				return v32
			}
		case 2:
			*(*uint16)(unsafe.Add(mBase, uint32(v32+int32(-5)))) = uint16(v5)
			if base.Ui32(int32(65536)) <= base.Ui32(v35) {
				F__serverAssert(m, int32(_a1323), int32(_a1321), int32(137))
				mBase = m.M
				v111 = m.ExcPending
				if v111 != 0 {
					return int32(0)
				} else {
					F_abort(m)
					mBase = m.M
					base.Wasm_trap_unreachable()
					for {
					}
				}
			} else {
				*(*uint16)(unsafe.Add(mBase, uint32(v32+int32(-3)))) = uint16(v35)
				v72 = int32(2)
				*(*uint8)(unsafe.Add(mBase, uint32(v32+int32(-1)))) = uint8(v72)
				v78 = *(*int32)(unsafe.Add(mBase, _consts[72]))
				if l3 == v78 {
				} else {
					if l3 != 0 {
						if v5 == int32(0) {
						} else {
							if v5 == int32(0) {
							} else {
								v87 = F__emscripten_memcpy_bulkmem(m, v32, l3, v5)
								mBase = m.M
							}
						}
					} else {
						v82 = F__emscripten_memset_bulkmem(m, v32, base.I32_extend8_s(int32(0)), v5)
						mBase = m.M
					}
				}
				v90 = int32(0)
				*(*uint8)(unsafe.Add(mBase, uint32(v32+v5))) = uint8(v90)
				return v32
			}
		case 3:
			*(*int32)(unsafe.Add(mBase, uint32(v32+int32(-5)))) = v35
			*(*int32)(unsafe.Add(mBase, uint32(v32+int32(-9)))) = v5
			v72 = int32(3)
			*(*uint8)(unsafe.Add(mBase, uint32(v32+int32(-1)))) = uint8(v72)
			v78 = *(*int32)(unsafe.Add(mBase, _consts[72]))
			if l3 == v78 {
			} else {
				if l3 != 0 {
					if v5 == int32(0) {
					} else {
						if v5 == int32(0) {
						} else {
							v87 = F__emscripten_memcpy_bulkmem(m, v32, l3, v5)
							mBase = m.M
						}
					}
				} else {
					v82 = F__emscripten_memset_bulkmem(m, v32, base.I32_extend8_s(int32(0)), v5)
					mBase = m.M
				}
			}
			v90 = int32(0)
			*(*uint8)(unsafe.Add(mBase, uint32(v32+v5))) = uint8(v90)
			return v32
		case 4:
			*(*int64)(unsafe.Add(mBase, uint32(v32+int32(-9)))) = base.I64_extend_i32_u(v35)
			*(*int64)(unsafe.Add(mBase, uint32(v32+int32(-17)))) = base.I64_extend_i32_u(v5)
			v72 = int32(4)
			*(*uint8)(unsafe.Add(mBase, uint32(v32+int32(-1)))) = uint8(v72)
			v78 = *(*int32)(unsafe.Add(mBase, _consts[72]))
			if l3 == v78 {
			} else {
				if l3 != 0 {
					if v5 == int32(0) {
					} else {
						if v5 == int32(0) {
						} else {
							v87 = F__emscripten_memcpy_bulkmem(m, v32, l3, v5)
							mBase = m.M
						}
					}
				} else {
					v82 = F__emscripten_memset_bulkmem(m, v32, base.I32_extend8_s(int32(0)), v5)
					mBase = m.M
				}
			}
			v90 = int32(0)
			*(*uint8)(unsafe.Add(mBase, uint32(v32+v5))) = uint8(v90)
			return v32
		default:
			v78 = *(*int32)(unsafe.Add(mBase, _consts[72]))
			if l3 == v78 {
			} else {
				if l3 != 0 {
					if v5 == int32(0) {
					} else {
						if v5 == int32(0) {
						} else {
							v87 = F__emscripten_memcpy_bulkmem(m, v32, l3, v5)
							mBase = m.M
						}
					}
				} else {
					v82 = F__emscripten_memset_bulkmem(m, v32, base.I32_extend8_s(int32(0)), v5)
					mBase = m.M
				}
			}
			v90 = int32(0)
			*(*uint8)(unsafe.Add(mBase, uint32(v32+v5))) = uint8(v90)
			return v32
		}
	}
}
func F_searchPreMonitorCfgName(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
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
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v32 int32
	_ = v32
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
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
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
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v102 int32
	_ = v102
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v131 int32
	_ = v131
	var v134 int32
	_ = v134
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v140 int32
	_ = v140
	var v144 int32
	_ = v144
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v173 int32
	_ = v173
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
	var v186 int32
	_ = v186
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v215 int32
	_ = v215
	var v218 int32
	_ = v218
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v224 int32
	_ = v224
	var v228 int32
	_ = v228
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v250 int32
	_ = v250
	var v251 int32
	_ = v251
	var v252 int32
	_ = v252
	var v257 int32
	_ = v257
	var v260 int32
	_ = v260
	var v262 int32
	_ = v262
	var v263 int32
	_ = v263
	var v264 int32
	_ = v264
	var v266 int32
	_ = v266
	var v270 int32
	_ = v270
	var v272 int32
	_ = v272
	var v273 int32
	_ = v273
	var v274 int32
	_ = v274
	var v276 int32
	_ = v276
	var v277 int32
	_ = v277
	var v278 int32
	_ = v278
	var v279 int32
	_ = v279
	var v284 int32
	_ = v284
	var v285 int32
	_ = v285
	var v288 int32
	_ = v288
	var v289 int32
	_ = v289
	var v292 int32
	_ = v292
	var v293 int32
	_ = v293
	var v294 int32
	_ = v294
	var v299 int32
	_ = v299
	var v302 int32
	_ = v302
	var v304 int32
	_ = v304
	var v305 int32
	_ = v305
	var v306 int32
	_ = v306
	var v308 int32
	_ = v308
	var v312 int32
	_ = v312
	var v314 int32
	_ = v314
	var v315 int32
	_ = v315
	var v316 int32
	_ = v316
	var v318 int32
	_ = v318
	var v319 int32
	_ = v319
	var v320 int32
	_ = v320
	var v321 int32
	_ = v321
	var v326 int32
	_ = v326
	var v327 int32
	_ = v327
	var v330 int32
	_ = v330
	var v331 int32
	_ = v331
	var v334 int32
	_ = v334
	var v335 int32
	_ = v335
	var v336 int32
	_ = v336
	var v341 int32
	_ = v341
	var v344 int32
	_ = v344
	var v346 int32
	_ = v346
	var v347 int32
	_ = v347
	var v348 int32
	_ = v348
	var v350 int32
	_ = v350
	var v354 int32
	_ = v354
	var v356 int32
	_ = v356
	var v357 int32
	_ = v357
	var v358 int32
	_ = v358
	var v360 int32
	_ = v360
	var v361 int32
	_ = v361
	var v362 int32
	_ = v362
	var v363 int32
	_ = v363
	var v368 int32
	_ = v368
	var v369 int32
	_ = v369
	var v372 int32
	_ = v372
	var v373 int32
	_ = v373
	var v376 int32
	_ = v376
	var v377 int32
	_ = v377
	var v378 int32
	_ = v378
	var v382 int32
	_ = v382
	v3 = int32(1)
	v5 = *(*int32)(unsafe.Add(mBase, _consts[736]))
	v8 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5))))
	if v8 != 0 {
		goto L4
	} else {
		goto L5
	}
L1:
	;
	return v382
L2:
	;
	if v40-v42 == int32(0) {
		v382 = v3
		goto L1
	} else {
		goto L14
	}
L3:
	;
	v40 = F_tolower(m, v36)
	mBase = m.M
	v41 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v37))))
	v42 = F_tolower(m, v41)
	mBase = m.M
	goto L2
L4:
	;
	v10 = v5
	v11 = l0
	v12 = v8
	goto L7
L5:
	;
	v36 = int32(0)
	v37 = l0
	goto L3
L6:
	;
	v36 = v33 & int32(255)
	v37 = v32
	goto L3
L7:
	;
	v14 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11))))
	if v14 == int32(0) {
		v32 = v11
		v33 = v12
		goto L6
	} else {
		goto L9
	}
L8:
	;
	v32 = v26
	v33 = int32(0)
	goto L6
L9:
	;
	v18 = v12 & int32(255)
	if v18 == v14 {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	v25 = int32(1)
	v26 = v11 + v25
	v27 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+1)))
	if v27 != 0 {
		v10 = v10 + v25
		v11 = v26
		v12 = v27
		goto L7
	} else {
		goto L13
	}
L11:
	;
	v20 = F_tolower(m, v18)
	mBase = m.M
	v21 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11))))
	v22 = F_tolower(m, v21)
	mBase = m.M
	if v20 == v22 {
		goto L10
	} else {
		goto L12
	}
L12:
	;
	v24 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10))))
	v32 = v11
	v33 = v24
	goto L6
L13:
	;
	goto L8
L14:
	;
	v47 = *(*int32)(unsafe.Add(mBase, _consts[737]))
	v50 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v47))))
	if v50 != 0 {
		goto L17
	} else {
		goto L18
	}
L15:
	;
	if v82-v84 == int32(0) {
		v382 = v3
		goto L1
	} else {
		goto L27
	}
L16:
	;
	v82 = F_tolower(m, v78)
	mBase = m.M
	v83 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v79))))
	v84 = F_tolower(m, v83)
	mBase = m.M
	goto L15
L17:
	;
	v52 = v47
	v53 = l0
	v54 = v50
	goto L20
L18:
	;
	v78 = int32(0)
	v79 = l0
	goto L16
L19:
	;
	v78 = v75 & int32(255)
	v79 = v74
	goto L16
L20:
	;
	v56 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v53))))
	if v56 == int32(0) {
		v74 = v53
		v75 = v54
		goto L19
	} else {
		goto L22
	}
L21:
	;
	v74 = v68
	v75 = int32(0)
	goto L19
L22:
	;
	v60 = v54 & int32(255)
	if v60 == v56 {
		goto L23
	} else {
		goto L24
	}
L23:
	;
	v67 = int32(1)
	v68 = v53 + v67
	v69 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v52)+1)))
	if v69 != 0 {
		v52 = v52 + v67
		v53 = v68
		v54 = v69
		goto L20
	} else {
		goto L26
	}
L24:
	;
	v62 = F_tolower(m, v60)
	mBase = m.M
	v63 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v53))))
	v64 = F_tolower(m, v63)
	mBase = m.M
	if v62 == v64 {
		goto L23
	} else {
		goto L25
	}
L25:
	;
	v66 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v52))))
	v74 = v53
	v75 = v66
	goto L19
L26:
	;
	goto L21
L27:
	;
	v89 = *(*int32)(unsafe.Add(mBase, _consts[738]))
	v92 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v89))))
	if v92 != 0 {
		goto L30
	} else {
		goto L31
	}
L28:
	;
	if v124-v126 == int32(0) {
		v382 = v3
		goto L1
	} else {
		goto L40
	}
L29:
	;
	v124 = F_tolower(m, v120)
	mBase = m.M
	v125 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v121))))
	v126 = F_tolower(m, v125)
	mBase = m.M
	goto L28
L30:
	;
	v94 = v89
	v95 = l0
	v96 = v92
	goto L33
L31:
	;
	v120 = int32(0)
	v121 = l0
	goto L29
L32:
	;
	v120 = v117 & int32(255)
	v121 = v116
	goto L29
L33:
	;
	v98 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v95))))
	if v98 == int32(0) {
		v116 = v95
		v117 = v96
		goto L32
	} else {
		goto L35
	}
L34:
	;
	v116 = v110
	v117 = int32(0)
	goto L32
L35:
	;
	v102 = v96 & int32(255)
	if v102 == v98 {
		goto L36
	} else {
		goto L37
	}
L36:
	;
	v109 = int32(1)
	v110 = v95 + v109
	v111 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v94)+1)))
	if v111 != 0 {
		v94 = v94 + v109
		v95 = v110
		v96 = v111
		goto L33
	} else {
		goto L39
	}
L37:
	;
	v104 = F_tolower(m, v102)
	mBase = m.M
	v105 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v95))))
	v106 = F_tolower(m, v105)
	mBase = m.M
	if v104 == v106 {
		goto L36
	} else {
		goto L38
	}
L38:
	;
	v108 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v94))))
	v116 = v95
	v117 = v108
	goto L32
L39:
	;
	goto L34
L40:
	;
	v131 = *(*int32)(unsafe.Add(mBase, _consts[739]))
	v134 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v131))))
	if v134 != 0 {
		goto L43
	} else {
		goto L44
	}
L41:
	;
	if v166-v168 == int32(0) {
		v382 = v3
		goto L1
	} else {
		goto L53
	}
L42:
	;
	v166 = F_tolower(m, v162)
	mBase = m.M
	v167 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v163))))
	v168 = F_tolower(m, v167)
	mBase = m.M
	goto L41
L43:
	;
	v136 = v131
	v137 = l0
	v138 = v134
	goto L46
L44:
	;
	v162 = int32(0)
	v163 = l0
	goto L42
L45:
	;
	v162 = v159 & int32(255)
	v163 = v158
	goto L42
L46:
	;
	v140 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v137))))
	if v140 == int32(0) {
		v158 = v137
		v159 = v138
		goto L45
	} else {
		goto L48
	}
L47:
	;
	v158 = v152
	v159 = int32(0)
	goto L45
L48:
	;
	v144 = v138 & int32(255)
	if v144 == v140 {
		goto L49
	} else {
		goto L50
	}
L49:
	;
	v151 = int32(1)
	v152 = v137 + v151
	v153 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v136)+1)))
	if v153 != 0 {
		v136 = v136 + v151
		v137 = v152
		v138 = v153
		goto L46
	} else {
		goto L52
	}
L50:
	;
	v146 = F_tolower(m, v144)
	mBase = m.M
	v147 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v137))))
	v148 = F_tolower(m, v147)
	mBase = m.M
	if v146 == v148 {
		goto L49
	} else {
		goto L51
	}
L51:
	;
	v150 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v136))))
	v158 = v137
	v159 = v150
	goto L45
L52:
	;
	goto L47
L53:
	;
	v173 = *(*int32)(unsafe.Add(mBase, _consts[740]))
	v176 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v173))))
	if v176 != 0 {
		goto L56
	} else {
		goto L57
	}
L54:
	;
	if v208-v210 == int32(0) {
		v382 = v3
		goto L1
	} else {
		goto L66
	}
L55:
	;
	v208 = F_tolower(m, v204)
	mBase = m.M
	v209 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v205))))
	v210 = F_tolower(m, v209)
	mBase = m.M
	goto L54
L56:
	;
	v178 = v173
	v179 = l0
	v180 = v176
	goto L59
L57:
	;
	v204 = int32(0)
	v205 = l0
	goto L55
L58:
	;
	v204 = v201 & int32(255)
	v205 = v200
	goto L55
L59:
	;
	v182 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v179))))
	if v182 == int32(0) {
		v200 = v179
		v201 = v180
		goto L58
	} else {
		goto L61
	}
L60:
	;
	v200 = v194
	v201 = int32(0)
	goto L58
L61:
	;
	v186 = v180 & int32(255)
	if v186 == v182 {
		goto L62
	} else {
		goto L63
	}
L62:
	;
	v193 = int32(1)
	v194 = v179 + v193
	v195 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v178)+1)))
	if v195 != 0 {
		v178 = v178 + v193
		v179 = v194
		v180 = v195
		goto L59
	} else {
		goto L65
	}
L63:
	;
	v188 = F_tolower(m, v186)
	mBase = m.M
	v189 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v179))))
	v190 = F_tolower(m, v189)
	mBase = m.M
	if v188 == v190 {
		goto L62
	} else {
		goto L64
	}
L64:
	;
	v192 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v178))))
	v200 = v179
	v201 = v192
	goto L58
L65:
	;
	goto L60
L66:
	;
	v215 = *(*int32)(unsafe.Add(mBase, _consts[741]))
	v218 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v215))))
	if v218 != 0 {
		goto L69
	} else {
		goto L70
	}
L67:
	;
	if v250-v252 == int32(0) {
		v382 = v3
		goto L1
	} else {
		goto L79
	}
L68:
	;
	v250 = F_tolower(m, v246)
	mBase = m.M
	v251 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v247))))
	v252 = F_tolower(m, v251)
	mBase = m.M
	goto L67
L69:
	;
	v220 = v215
	v221 = l0
	v222 = v218
	goto L72
L70:
	;
	v246 = int32(0)
	v247 = l0
	goto L68
L71:
	;
	v246 = v243 & int32(255)
	v247 = v242
	goto L68
L72:
	;
	v224 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v221))))
	if v224 == int32(0) {
		v242 = v221
		v243 = v222
		goto L71
	} else {
		goto L74
	}
L73:
	;
	v242 = v236
	v243 = int32(0)
	goto L71
L74:
	;
	v228 = v222 & int32(255)
	if v228 == v224 {
		goto L75
	} else {
		goto L76
	}
L75:
	;
	v235 = int32(1)
	v236 = v221 + v235
	v237 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v220)+1)))
	if v237 != 0 {
		v220 = v220 + v235
		v221 = v236
		v222 = v237
		goto L72
	} else {
		goto L78
	}
L76:
	;
	v230 = F_tolower(m, v228)
	mBase = m.M
	v231 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v221))))
	v232 = F_tolower(m, v231)
	mBase = m.M
	if v230 == v232 {
		goto L75
	} else {
		goto L77
	}
L77:
	;
	v234 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v220))))
	v242 = v221
	v243 = v234
	goto L71
L78:
	;
	goto L73
L79:
	;
	v257 = *(*int32)(unsafe.Add(mBase, _consts[742]))
	v260 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v257))))
	if v260 != 0 {
		goto L82
	} else {
		goto L83
	}
L80:
	;
	if v292-v294 == int32(0) {
		v382 = v3
		goto L1
	} else {
		goto L92
	}
L81:
	;
	v292 = F_tolower(m, v288)
	mBase = m.M
	v293 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v289))))
	v294 = F_tolower(m, v293)
	mBase = m.M
	goto L80
L82:
	;
	v262 = v257
	v263 = l0
	v264 = v260
	goto L85
L83:
	;
	v288 = int32(0)
	v289 = l0
	goto L81
L84:
	;
	v288 = v285 & int32(255)
	v289 = v284
	goto L81
L85:
	;
	v266 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v263))))
	if v266 == int32(0) {
		v284 = v263
		v285 = v264
		goto L84
	} else {
		goto L87
	}
L86:
	;
	v284 = v278
	v285 = int32(0)
	goto L84
L87:
	;
	v270 = v264 & int32(255)
	if v270 == v266 {
		goto L88
	} else {
		goto L89
	}
L88:
	;
	v277 = int32(1)
	v278 = v263 + v277
	v279 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v262)+1)))
	if v279 != 0 {
		v262 = v262 + v277
		v263 = v278
		v264 = v279
		goto L85
	} else {
		goto L91
	}
L89:
	;
	v272 = F_tolower(m, v270)
	mBase = m.M
	v273 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v263))))
	v274 = F_tolower(m, v273)
	mBase = m.M
	if v272 == v274 {
		goto L88
	} else {
		goto L90
	}
L90:
	;
	v276 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v262))))
	v284 = v263
	v285 = v276
	goto L84
L91:
	;
	goto L86
L92:
	;
	v299 = *(*int32)(unsafe.Add(mBase, _consts[743]))
	v302 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v299))))
	if v302 != 0 {
		goto L95
	} else {
		goto L96
	}
L93:
	;
	if v334-v336 == int32(0) {
		v382 = v3
		goto L1
	} else {
		goto L105
	}
L94:
	;
	v334 = F_tolower(m, v330)
	mBase = m.M
	v335 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v331))))
	v336 = F_tolower(m, v335)
	mBase = m.M
	goto L93
L95:
	;
	v304 = v299
	v305 = l0
	v306 = v302
	goto L98
L96:
	;
	v330 = int32(0)
	v331 = l0
	goto L94
L97:
	;
	v330 = v327 & int32(255)
	v331 = v326
	goto L94
L98:
	;
	v308 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v305))))
	if v308 == int32(0) {
		v326 = v305
		v327 = v306
		goto L97
	} else {
		goto L100
	}
L99:
	;
	v326 = v320
	v327 = int32(0)
	goto L97
L100:
	;
	v312 = v306 & int32(255)
	if v312 == v308 {
		goto L101
	} else {
		goto L102
	}
L101:
	;
	v319 = int32(1)
	v320 = v305 + v319
	v321 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v304)+1)))
	if v321 != 0 {
		v304 = v304 + v319
		v305 = v320
		v306 = v321
		goto L98
	} else {
		goto L104
	}
L102:
	;
	v314 = F_tolower(m, v312)
	mBase = m.M
	v315 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v305))))
	v316 = F_tolower(m, v315)
	mBase = m.M
	if v314 == v316 {
		goto L101
	} else {
		goto L103
	}
L103:
	;
	v318 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v304))))
	v326 = v305
	v327 = v318
	goto L97
L104:
	;
	goto L99
L105:
	;
	v341 = *(*int32)(unsafe.Add(mBase, _consts[744]))
	v344 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v341))))
	if v344 != 0 {
		goto L108
	} else {
		goto L109
	}
L106:
	;
	v382 = base.B2i32(v376-v378 == int32(0))
	goto L1
L107:
	;
	v376 = F_tolower(m, v372)
	mBase = m.M
	v377 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v373))))
	v378 = F_tolower(m, v377)
	mBase = m.M
	goto L106
L108:
	;
	v346 = v341
	v347 = l0
	v348 = v344
	goto L111
L109:
	;
	v372 = int32(0)
	v373 = l0
	goto L107
L110:
	;
	v372 = v369 & int32(255)
	v373 = v368
	goto L107
L111:
	;
	v350 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v347))))
	if v350 == int32(0) {
		v368 = v347
		v369 = v348
		goto L110
	} else {
		goto L113
	}
L112:
	;
	v368 = v362
	v369 = int32(0)
	goto L110
L113:
	;
	v354 = v348 & int32(255)
	if v354 == v350 {
		goto L114
	} else {
		goto L115
	}
L114:
	;
	v361 = int32(1)
	v362 = v347 + v361
	v363 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v346)+1)))
	if v363 != 0 {
		v346 = v346 + v361
		v347 = v362
		v348 = v363
		goto L111
	} else {
		goto L117
	}
L115:
	;
	v356 = F_tolower(m, v354)
	mBase = m.M
	v357 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v347))))
	v358 = F_tolower(m, v357)
	mBase = m.M
	if v356 == v358 {
		goto L114
	} else {
		goto L116
	}
L116:
	;
	v360 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v346))))
	v368 = v347
	v369 = v360
	goto L110
L117:
	;
	goto L112
}
func F_securityWarningCommand(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v12 int64
	_ = v12
	var v14 int64
	_ = v14
	var v15 int64
	_ = v15
	var v17 int64
	_ = v17
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v60 int32
	_ = v60
	var v70 int32
	_ = v70
	var v76 int32
	_ = v76
	var v83 int32
	_ = v83
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v102 int32
	_ = v102
	v2 = int32(0)
	v7 = m.G0
	v9 = v7 - int32(96)
	m.G0 = v9
	v12 = F___time(m, v2)
	mBase = m.M
	v14 = *(*int64)(unsafe.Add(mBase, _consts[600]))
	v15 = v12 - v14
	v17 = v15 >> (uint(int64(63)) % 64)
	if base.Ui64(v15^v17-v17) < base.Ui64(int64(61)) {
		v76 = *(*int32)(unsafe.Add(mBase, uint32(l0)+200))
		if v76&int32(1280) != 0 {
			m.G0 = v9 + int32(96)
			return
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(l0)+200)) = v76 | int32(1024)
			v83 = *(*int32)(unsafe.Add(mBase, _consts[88]))
			if v83 == int32(0) {
				v91 = *(*int32)(unsafe.Add(mBase, _consts[593]))
				v92 = F_listAddNodeTail(m, v91, l0)
				mBase = m.M
				v93 = m.ExcPending
				if v93 != 0 {
					return
				} else {
					m.G0 = v9 + int32(96)
					return
				}
			} else {
				v87 = *(*int32)(unsafe.Add(mBase, _consts[593]))
				v88 = F_listSearchKey(m, v87, l0)
				mBase = m.M
				v89 = m.ExcPending
				if v89 != 0 {
					return
				} else {
					if v88 != 0 {
						F__serverAssertWithInfo(m, l0, int32(0), int32(_a998), int32(_a977), int32(2277))
						mBase = m.M
						v102 = m.ExcPending
						if v102 != 0 {
							return
						} else {
							F_abort(m)
							mBase = m.M
							base.Wasm_trap_unreachable()
							for {
							}
						}
					} else {
						v91 = *(*int32)(unsafe.Add(mBase, _consts[593]))
						v92 = F_listAddNodeTail(m, v91, l0)
						mBase = m.M
						v93 = m.ExcPending
						if v93 != 0 {
							return
						} else {
							m.G0 = v9 + int32(96)
							return
						}
					}
				}
			}
		}
	} else {
		v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		if v22 == int32(0) {
			v41 = *(*int32)(unsafe.Add(mBase, _consts[6]))
			if int32(3) < v41 {
				*(*int64)(unsafe.Add(mBase, _consts[600])) = v12
				v76 = *(*int32)(unsafe.Add(mBase, uint32(l0)+200))
				if v76&int32(1280) != 0 {
					m.G0 = v9 + int32(96)
					return
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(l0)+200)) = v76 | int32(1024)
					v83 = *(*int32)(unsafe.Add(mBase, _consts[88]))
					if v83 == int32(0) {
						v91 = *(*int32)(unsafe.Add(mBase, _consts[593]))
						v92 = F_listAddNodeTail(m, v91, l0)
						mBase = m.M
						v93 = m.ExcPending
						if v93 != 0 {
							return
						} else {
							m.G0 = v9 + int32(96)
							return
						}
					} else {
						v87 = *(*int32)(unsafe.Add(mBase, _consts[593]))
						v88 = F_listSearchKey(m, v87, l0)
						mBase = m.M
						v89 = m.ExcPending
						if v89 != 0 {
							return
						} else {
							if v88 != 0 {
								F__serverAssertWithInfo(m, l0, int32(0), int32(_a998), int32(_a977), int32(2277))
								mBase = m.M
								v102 = m.ExcPending
								if v102 != 0 {
									return
								} else {
									F_abort(m)
									mBase = m.M
									base.Wasm_trap_unreachable()
									for {
									}
								}
							} else {
								v91 = *(*int32)(unsafe.Add(mBase, _consts[593]))
								v92 = F_listAddNodeTail(m, v91, l0)
								mBase = m.M
								v93 = m.ExcPending
								if v93 != 0 {
									return
								} else {
									m.G0 = v9 + int32(96)
									return
								}
							}
						}
					}
				}
			} else {
				v44 = int32(_a46)
				*(*int32)(unsafe.Add(mBase, uint32(v9)+4)) = v44
				*(*int32)(unsafe.Add(mBase, uint32(v9))) = v44
				F__serverLog(m, int32(3), int32(_a1048), v9)
				mBase = m.M
				v51 = m.ExcPending
				if v51 != 0 {
					return
				} else {
					*(*int64)(unsafe.Add(mBase, _consts[600])) = v12
					v76 = *(*int32)(unsafe.Add(mBase, uint32(l0)+200))
					if v76&int32(1280) != 0 {
						m.G0 = v9 + int32(96)
						return
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(l0)+200)) = v76 | int32(1024)
						v83 = *(*int32)(unsafe.Add(mBase, _consts[88]))
						if v83 == int32(0) {
							v91 = *(*int32)(unsafe.Add(mBase, _consts[593]))
							v92 = F_listAddNodeTail(m, v91, l0)
							mBase = m.M
							v93 = m.ExcPending
							if v93 != 0 {
								return
							} else {
								m.G0 = v9 + int32(96)
								return
							}
						} else {
							v87 = *(*int32)(unsafe.Add(mBase, _consts[593]))
							v88 = F_listSearchKey(m, v87, l0)
							mBase = m.M
							v89 = m.ExcPending
							if v89 != 0 {
								return
							} else {
								if v88 != 0 {
									F__serverAssertWithInfo(m, l0, int32(0), int32(_a998), int32(_a977), int32(2277))
									mBase = m.M
									v102 = m.ExcPending
									if v102 != 0 {
										return
									} else {
										F_abort(m)
										mBase = m.M
										base.Wasm_trap_unreachable()
										for {
										}
									}
								} else {
									v91 = *(*int32)(unsafe.Add(mBase, _consts[593]))
									v92 = F_listAddNodeTail(m, v91, l0)
									mBase = m.M
									v93 = m.ExcPending
									if v93 != 0 {
										return
									} else {
										m.G0 = v9 + int32(96)
										return
									}
								}
							}
						}
					}
				}
			}
		} else {
			v25 = *(*int32)(unsafe.Add(mBase, uint32(v22)))
			v26 = *(*int32)(unsafe.Add(mBase, uint32(v25)+24))
			if v26 == int32(0) {
				v41 = *(*int32)(unsafe.Add(mBase, _consts[6]))
				if int32(3) < v41 {
					*(*int64)(unsafe.Add(mBase, _consts[600])) = v12
					v76 = *(*int32)(unsafe.Add(mBase, uint32(l0)+200))
					if v76&int32(1280) != 0 {
						m.G0 = v9 + int32(96)
						return
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(l0)+200)) = v76 | int32(1024)
						v83 = *(*int32)(unsafe.Add(mBase, _consts[88]))
						if v83 == int32(0) {
							v91 = *(*int32)(unsafe.Add(mBase, _consts[593]))
							v92 = F_listAddNodeTail(m, v91, l0)
							mBase = m.M
							v93 = m.ExcPending
							if v93 != 0 {
								return
							} else {
								m.G0 = v9 + int32(96)
								return
							}
						} else {
							v87 = *(*int32)(unsafe.Add(mBase, _consts[593]))
							v88 = F_listSearchKey(m, v87, l0)
							mBase = m.M
							v89 = m.ExcPending
							if v89 != 0 {
								return
							} else {
								if v88 != 0 {
									F__serverAssertWithInfo(m, l0, int32(0), int32(_a998), int32(_a977), int32(2277))
									mBase = m.M
									v102 = m.ExcPending
									if v102 != 0 {
										return
									} else {
										F_abort(m)
										mBase = m.M
										base.Wasm_trap_unreachable()
										for {
										}
									}
								} else {
									v91 = *(*int32)(unsafe.Add(mBase, _consts[593]))
									v92 = F_listAddNodeTail(m, v91, l0)
									mBase = m.M
									v93 = m.ExcPending
									if v93 != 0 {
										return
									} else {
										m.G0 = v9 + int32(96)
										return
									}
								}
							}
						}
					}
				} else {
					v44 = int32(_a46)
					*(*int32)(unsafe.Add(mBase, uint32(v9)+4)) = v44
					*(*int32)(unsafe.Add(mBase, uint32(v9))) = v44
					F__serverLog(m, int32(3), int32(_a1048), v9)
					mBase = m.M
					v51 = m.ExcPending
					if v51 != 0 {
						return
					} else {
						*(*int64)(unsafe.Add(mBase, _consts[600])) = v12
						v76 = *(*int32)(unsafe.Add(mBase, uint32(l0)+200))
						if v76&int32(1280) != 0 {
							m.G0 = v9 + int32(96)
							return
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(l0)+200)) = v76 | int32(1024)
							v83 = *(*int32)(unsafe.Add(mBase, _consts[88]))
							if v83 == int32(0) {
								v91 = *(*int32)(unsafe.Add(mBase, _consts[593]))
								v92 = F_listAddNodeTail(m, v91, l0)
								mBase = m.M
								v93 = m.ExcPending
								if v93 != 0 {
									return
								} else {
									m.G0 = v9 + int32(96)
									return
								}
							} else {
								v87 = *(*int32)(unsafe.Add(mBase, _consts[593]))
								v88 = F_listSearchKey(m, v87, l0)
								mBase = m.M
								v89 = m.ExcPending
								if v89 != 0 {
									return
								} else {
									if v88 != 0 {
										F__serverAssertWithInfo(m, l0, int32(0), int32(_a998), int32(_a977), int32(2277))
										mBase = m.M
										v102 = m.ExcPending
										if v102 != 0 {
											return
										} else {
											F_abort(m)
											mBase = m.M
											base.Wasm_trap_unreachable()
											for {
											}
										}
									} else {
										v91 = *(*int32)(unsafe.Add(mBase, _consts[593]))
										v92 = F_listAddNodeTail(m, v91, l0)
										mBase = m.M
										v93 = m.ExcPending
										if v93 != 0 {
											return
										} else {
											m.G0 = v9 + int32(96)
											return
										}
									}
								}
							}
						}
					}
				}
			} else {
				v35 = m.T0[v26].(func(*base.Module, int32, int32, int32, int32, int32) int32)(m, v22, v9+int32(48), int32(46), v9+int32(44), int32(1))
				mBase = m.M
				v36 = m.ExcPending
				if v36 != 0 {
					return
				} else {
					if v35 != int32(-1) {
						v53 = *(*int32)(unsafe.Add(mBase, _consts[6]))
						if int32(3) < v53 {
							*(*int64)(unsafe.Add(mBase, _consts[600])) = v12
							v76 = *(*int32)(unsafe.Add(mBase, uint32(l0)+200))
							if v76&int32(1280) != 0 {
								m.G0 = v9 + int32(96)
								return
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(l0)+200)) = v76 | int32(1024)
								v83 = *(*int32)(unsafe.Add(mBase, _consts[88]))
								if v83 == int32(0) {
									v91 = *(*int32)(unsafe.Add(mBase, _consts[593]))
									v92 = F_listAddNodeTail(m, v91, l0)
									mBase = m.M
									v93 = m.ExcPending
									if v93 != 0 {
										return
									} else {
										m.G0 = v9 + int32(96)
										return
									}
								} else {
									v87 = *(*int32)(unsafe.Add(mBase, _consts[593]))
									v88 = F_listSearchKey(m, v87, l0)
									mBase = m.M
									v89 = m.ExcPending
									if v89 != 0 {
										return
									} else {
										if v88 != 0 {
											F__serverAssertWithInfo(m, l0, int32(0), int32(_a998), int32(_a977), int32(2277))
											mBase = m.M
											v102 = m.ExcPending
											if v102 != 0 {
												return
											} else {
												F_abort(m)
												mBase = m.M
												base.Wasm_trap_unreachable()
												for {
												}
											}
										} else {
											v91 = *(*int32)(unsafe.Add(mBase, _consts[593]))
											v92 = F_listAddNodeTail(m, v91, l0)
											mBase = m.M
											v93 = m.ExcPending
											if v93 != 0 {
												return
											} else {
												m.G0 = v9 + int32(96)
												return
											}
										}
									}
								}
							}
						} else {
							v56 = int32(_a46)
							*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = v56
							*(*int32)(unsafe.Add(mBase, uint32(v9)+20)) = v56
							v60 = *(*int32)(unsafe.Add(mBase, uint32(v9)+44))
							*(*int32)(unsafe.Add(mBase, uint32(v9)+28)) = v60
							*(*int32)(unsafe.Add(mBase, uint32(v9)+24)) = v9 + int32(48)
							F__serverLog(m, int32(3), int32(_a1049), v9+int32(16))
							mBase = m.M
							v70 = m.ExcPending
							if v70 != 0 {
								return
							} else {
								*(*int64)(unsafe.Add(mBase, _consts[600])) = v12
								v76 = *(*int32)(unsafe.Add(mBase, uint32(l0)+200))
								if v76&int32(1280) != 0 {
									m.G0 = v9 + int32(96)
									return
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(l0)+200)) = v76 | int32(1024)
									v83 = *(*int32)(unsafe.Add(mBase, _consts[88]))
									if v83 == int32(0) {
										v91 = *(*int32)(unsafe.Add(mBase, _consts[593]))
										v92 = F_listAddNodeTail(m, v91, l0)
										mBase = m.M
										v93 = m.ExcPending
										if v93 != 0 {
											return
										} else {
											m.G0 = v9 + int32(96)
											return
										}
									} else {
										v87 = *(*int32)(unsafe.Add(mBase, _consts[593]))
										v88 = F_listSearchKey(m, v87, l0)
										mBase = m.M
										v89 = m.ExcPending
										if v89 != 0 {
											return
										} else {
											if v88 != 0 {
												F__serverAssertWithInfo(m, l0, int32(0), int32(_a998), int32(_a977), int32(2277))
												mBase = m.M
												v102 = m.ExcPending
												if v102 != 0 {
													return
												} else {
													F_abort(m)
													mBase = m.M
													base.Wasm_trap_unreachable()
													for {
													}
												}
											} else {
												v91 = *(*int32)(unsafe.Add(mBase, _consts[593]))
												v92 = F_listAddNodeTail(m, v91, l0)
												mBase = m.M
												v93 = m.ExcPending
												if v93 != 0 {
													return
												} else {
													m.G0 = v9 + int32(96)
													return
												}
											}
										}
									}
								}
							}
						}
					} else {
						v41 = *(*int32)(unsafe.Add(mBase, _consts[6]))
						if int32(3) < v41 {
							*(*int64)(unsafe.Add(mBase, _consts[600])) = v12
							v76 = *(*int32)(unsafe.Add(mBase, uint32(l0)+200))
							if v76&int32(1280) != 0 {
								m.G0 = v9 + int32(96)
								return
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(l0)+200)) = v76 | int32(1024)
								v83 = *(*int32)(unsafe.Add(mBase, _consts[88]))
								if v83 == int32(0) {
									v91 = *(*int32)(unsafe.Add(mBase, _consts[593]))
									v92 = F_listAddNodeTail(m, v91, l0)
									mBase = m.M
									v93 = m.ExcPending
									if v93 != 0 {
										return
									} else {
										m.G0 = v9 + int32(96)
										return
									}
								} else {
									v87 = *(*int32)(unsafe.Add(mBase, _consts[593]))
									v88 = F_listSearchKey(m, v87, l0)
									mBase = m.M
									v89 = m.ExcPending
									if v89 != 0 {
										return
									} else {
										if v88 != 0 {
											F__serverAssertWithInfo(m, l0, int32(0), int32(_a998), int32(_a977), int32(2277))
											mBase = m.M
											v102 = m.ExcPending
											if v102 != 0 {
												return
											} else {
												F_abort(m)
												mBase = m.M
												base.Wasm_trap_unreachable()
												for {
												}
											}
										} else {
											v91 = *(*int32)(unsafe.Add(mBase, _consts[593]))
											v92 = F_listAddNodeTail(m, v91, l0)
											mBase = m.M
											v93 = m.ExcPending
											if v93 != 0 {
												return
											} else {
												m.G0 = v9 + int32(96)
												return
											}
										}
									}
								}
							}
						} else {
							v44 = int32(_a46)
							*(*int32)(unsafe.Add(mBase, uint32(v9)+4)) = v44
							*(*int32)(unsafe.Add(mBase, uint32(v9))) = v44
							F__serverLog(m, int32(3), int32(_a1048), v9)
							mBase = m.M
							v51 = m.ExcPending
							if v51 != 0 {
								return
							} else {
								*(*int64)(unsafe.Add(mBase, _consts[600])) = v12
								v76 = *(*int32)(unsafe.Add(mBase, uint32(l0)+200))
								if v76&int32(1280) != 0 {
									m.G0 = v9 + int32(96)
									return
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(l0)+200)) = v76 | int32(1024)
									v83 = *(*int32)(unsafe.Add(mBase, _consts[88]))
									if v83 == int32(0) {
										v91 = *(*int32)(unsafe.Add(mBase, _consts[593]))
										v92 = F_listAddNodeTail(m, v91, l0)
										mBase = m.M
										v93 = m.ExcPending
										if v93 != 0 {
											return
										} else {
											m.G0 = v9 + int32(96)
											return
										}
									} else {
										v87 = *(*int32)(unsafe.Add(mBase, _consts[593]))
										v88 = F_listSearchKey(m, v87, l0)
										mBase = m.M
										v89 = m.ExcPending
										if v89 != 0 {
											return
										} else {
											if v88 != 0 {
												F__serverAssertWithInfo(m, l0, int32(0), int32(_a998), int32(_a977), int32(2277))
												mBase = m.M
												v102 = m.ExcPending
												if v102 != 0 {
													return
												} else {
													F_abort(m)
													mBase = m.M
													base.Wasm_trap_unreachable()
													for {
													}
												}
											} else {
												v91 = *(*int32)(unsafe.Add(mBase, _consts[593]))
												v92 = F_listAddNodeTail(m, v91, l0)
												mBase = m.M
												v93 = m.ExcPending
												if v93 != 0 {
													return
												} else {
													m.G0 = v9 + int32(96)
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
		}
	}
}
func F_setProtocolError(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v50 int64
	_ = v50
	var v52 int32
	_ = v52
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v97 int32
	_ = v97
	var v103 int32
	_ = v103
	var v109 int32
	_ = v109
	var v115 int32
	_ = v115
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v130 int32
	_ = v130
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v149 int32
	_ = v149
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v169 int32
	_ = v169
	var v171 int32
	_ = v171
	var v185 int32
	_ = v185
	var v188 int32
	_ = v188
	var v191 int32
	_ = v191
	var v193 int32
	_ = v193
	var v195 int32
	_ = v195
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v201 int32
	_ = v201
	var v210 int32
	_ = v210
	var v212 int32
	_ = v212
	var v223 int32
	_ = v223
	var v227 int32
	_ = v227
	v11 = m.G0
	v13 = v11 - int32(320)
	m.G0 = v13
	v16 = *(*int32)(unsafe.Add(mBase, _consts[6]))
	if v16 < int32(2) {
		v34 = F_sdsempty(m)
		mBase = m.M
		v35 = m.ExcPending
		if v35 != 0 {
			return
		} else {
			v37 = *(*int32)(unsafe.Add(mBase, _consts[209]))
			v38 = F_catClientInfoString(m, v34, l1, v37)
			mBase = m.M
			v39 = m.ExcPending
			if v39 != 0 {
				return
			} else {
				v40 = int32(0)
				*(*uint8)(unsafe.Add(mBase, uint32(v13)+64)) = uint8(v40)
				v43 = *(*int32)(unsafe.Add(mBase, _consts[209]))
				if v43 == v40 {
					v52 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
					if v52 == int32(0) {
						v149 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+64)))
						if v149 == int32(0) {
						} else {
							v158 = v13 + int32(64)
							v159 = v149
							for {
								if base.Ui32(base.I32_extend8_s(v159)+int32(-32)) < base.Ui32(int32(95)) {
								} else {
									v169 = int32(46)
									*(*uint8)(unsafe.Add(mBase, uint32(v158))) = uint8(v169)
								}
								v171 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v158)+1)))
								if v171 != 0 {
									v158 = v158 + int32(1)
									v159 = v171
									continue
								} else {
									break
								}
								break
							}
						}
						v185 = int32(1)
						v188 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+200)))
						if v188&v185 != 0 {
							v195 = v185
							v198 = v195
						} else {
							v191 = *(*int32)(unsafe.Add(mBase, uint32(l1)+216))
							if v191 != 0 {
								v193 = F_isImportSlotMigrationJob(m, v191)
								mBase = m.M
								v195 = v193
								v198 = v195
							} else {
								v198 = int32(0)
							}
						}
						if v198 != 0 {
							v199 = int32(3)
						} else {
							v199 = v185
						}
						v201 = *(*int32)(unsafe.Add(mBase, _consts[6]))
						if v199 < v201 {
							F_sdsfree(m, v38)
							mBase = m.M
							v212 = m.ExcPending
							if v212 != 0 {
								return
							} else {
								v223 = *(*int32)(unsafe.Add(mBase, uint32(l1)+200))
								*(*int32)(unsafe.Add(mBase, uint32(l1)+200)) = v223 | int32(64)
								v227 = *(*int32)(unsafe.Add(mBase, uint32(l1)+204))
								*(*int32)(unsafe.Add(mBase, uint32(l1)+204)) = v227 | int32(1024)
								m.G0 = v13 + int32(320)
								return
							}
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v13)+4)) = v38
							*(*int32)(unsafe.Add(mBase, uint32(v13))) = l0
							*(*int32)(unsafe.Add(mBase, uint32(v13)+8)) = v13 + int32(64)
							F__serverLog(m, v199, int32(_a1016), v13)
							mBase = m.M
							v210 = m.ExcPending
							if v210 != 0 {
								return
							} else {
								F_sdsfree(m, v38)
								mBase = m.M
								v212 = m.ExcPending
								if v212 != 0 {
									return
								} else {
									v223 = *(*int32)(unsafe.Add(mBase, uint32(l1)+200))
									*(*int32)(unsafe.Add(mBase, uint32(l1)+200)) = v223 | int32(64)
									v227 = *(*int32)(unsafe.Add(mBase, uint32(l1)+204))
									*(*int32)(unsafe.Add(mBase, uint32(l1)+204)) = v227 | int32(1024)
									m.G0 = v13 + int32(320)
									return
								}
							}
						}
					} else {
						v58 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v52+int32(-1)))))
						v60 = v58 & int32(7)
						switch v60 {
						case 0:
							v75 = int32(base.Ui32(v58) >> (uint(int32(3)) % 32))
						case 1:
							v65 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v52+int32(-3)))))
							v75 = v65
						case 2:
							v68 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v52+int32(-5)))))
							v75 = v68
						case 3:
							v71 = *(*int32)(unsafe.Add(mBase, uint32(v52+int32(-9))))
							v75 = v71
						case 4:
							v74 = *(*int32)(unsafe.Add(mBase, uint32(v52+int32(-17))))
							v75 = v74
						default:
							v75 = int32(0)
						}
						v76 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
						v77 = v52 + v76
						if base.Ui32(int32(127)) < base.Ui32(v75-v76) {
							switch v60 {
							case 0:
								v91 = int32(base.Ui32(v58) >> (uint(int32(3)) % 32))
								v122 = v91
								v123 = v91 - v76 + int32(-128)
							case 1:
								v97 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v52+int32(-3)))))
								v122 = v97
								v123 = v97 - v76 + int32(-128)
							case 2:
								v103 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v52+int32(-5)))))
								v122 = v103
								v123 = v103 - v76 + int32(-128)
							case 3:
								v109 = *(*int32)(unsafe.Add(mBase, uint32(v52+int32(-9))))
								v122 = v109
								v123 = v109 - v76 + int32(-128)
							case 4:
								v115 = *(*int32)(unsafe.Add(mBase, uint32(v52+int32(-17))))
								v122 = v115
								v123 = v115 - v76 + int32(-128)
							default:
								v122 = int32(0)
								v123 = int32(-128) - v76
							}
							*(*int32)(unsafe.Add(mBase, uint32(v13+int32(48)))) = v52 + v122 + int32(-64)
							v130 = int32(64)
							*(*int32)(unsafe.Add(mBase, uint32(v13)+44)) = v130
							*(*int32)(unsafe.Add(mBase, uint32(v13)+40)) = v123
							*(*int32)(unsafe.Add(mBase, uint32(v13)+36)) = v77
							*(*int32)(unsafe.Add(mBase, uint32(v13)+32)) = v130
							v142 = F_snprintf(m, v13+v130, int32(256), int32(_a1017), v13+int32(32))
							mBase = m.M
							v143 = m.ExcPending
							if v143 != 0 {
								return
							} else {
								v149 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+64)))
								if v149 == int32(0) {
								} else {
									v158 = v13 + int32(64)
									v159 = v149
									for {
										if base.Ui32(base.I32_extend8_s(v159)+int32(-32)) < base.Ui32(int32(95)) {
										} else {
											v169 = int32(46)
											*(*uint8)(unsafe.Add(mBase, uint32(v158))) = uint8(v169)
										}
										v171 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v158)+1)))
										if v171 != 0 {
											v158 = v158 + int32(1)
											v159 = v171
											continue
										} else {
											break
										}
										break
									}
								}
								v185 = int32(1)
								v188 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+200)))
								if v188&v185 != 0 {
									v195 = v185
									v198 = v195
								} else {
									v191 = *(*int32)(unsafe.Add(mBase, uint32(l1)+216))
									if v191 != 0 {
										v193 = F_isImportSlotMigrationJob(m, v191)
										mBase = m.M
										v195 = v193
										v198 = v195
									} else {
										v198 = int32(0)
									}
								}
								if v198 != 0 {
									v199 = int32(3)
								} else {
									v199 = v185
								}
								v201 = *(*int32)(unsafe.Add(mBase, _consts[6]))
								if v199 < v201 {
									F_sdsfree(m, v38)
									mBase = m.M
									v212 = m.ExcPending
									if v212 != 0 {
										return
									} else {
										v223 = *(*int32)(unsafe.Add(mBase, uint32(l1)+200))
										*(*int32)(unsafe.Add(mBase, uint32(l1)+200)) = v223 | int32(64)
										v227 = *(*int32)(unsafe.Add(mBase, uint32(l1)+204))
										*(*int32)(unsafe.Add(mBase, uint32(l1)+204)) = v227 | int32(1024)
										m.G0 = v13 + int32(320)
										return
									}
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v13)+4)) = v38
									*(*int32)(unsafe.Add(mBase, uint32(v13))) = l0
									*(*int32)(unsafe.Add(mBase, uint32(v13)+8)) = v13 + int32(64)
									F__serverLog(m, v199, int32(_a1016), v13)
									mBase = m.M
									v210 = m.ExcPending
									if v210 != 0 {
										return
									} else {
										F_sdsfree(m, v38)
										mBase = m.M
										v212 = m.ExcPending
										if v212 != 0 {
											return
										} else {
											v223 = *(*int32)(unsafe.Add(mBase, uint32(l1)+200))
											*(*int32)(unsafe.Add(mBase, uint32(l1)+200)) = v223 | int32(64)
											v227 = *(*int32)(unsafe.Add(mBase, uint32(l1)+204))
											*(*int32)(unsafe.Add(mBase, uint32(l1)+204)) = v227 | int32(1024)
											m.G0 = v13 + int32(320)
											return
										}
									}
								}
							}
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v13)+16)) = v77
							v88 = F_snprintf(m, v13+int32(64), int32(256), int32(_a1018), v13+int32(16))
							mBase = m.M
							v89 = m.ExcPending
							if v89 != 0 {
								return
							} else {
								v149 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+64)))
								if v149 == int32(0) {
								} else {
									v158 = v13 + int32(64)
									v159 = v149
									for {
										if base.Ui32(base.I32_extend8_s(v159)+int32(-32)) < base.Ui32(int32(95)) {
										} else {
											v169 = int32(46)
											*(*uint8)(unsafe.Add(mBase, uint32(v158))) = uint8(v169)
										}
										v171 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v158)+1)))
										if v171 != 0 {
											v158 = v158 + int32(1)
											v159 = v171
											continue
										} else {
											break
										}
										break
									}
								}
								v185 = int32(1)
								v188 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+200)))
								if v188&v185 != 0 {
									v195 = v185
									v198 = v195
								} else {
									v191 = *(*int32)(unsafe.Add(mBase, uint32(l1)+216))
									if v191 != 0 {
										v193 = F_isImportSlotMigrationJob(m, v191)
										mBase = m.M
										v195 = v193
										v198 = v195
									} else {
										v198 = int32(0)
									}
								}
								if v198 != 0 {
									v199 = int32(3)
								} else {
									v199 = v185
								}
								v201 = *(*int32)(unsafe.Add(mBase, _consts[6]))
								if v199 < v201 {
									F_sdsfree(m, v38)
									mBase = m.M
									v212 = m.ExcPending
									if v212 != 0 {
										return
									} else {
										v223 = *(*int32)(unsafe.Add(mBase, uint32(l1)+200))
										*(*int32)(unsafe.Add(mBase, uint32(l1)+200)) = v223 | int32(64)
										v227 = *(*int32)(unsafe.Add(mBase, uint32(l1)+204))
										*(*int32)(unsafe.Add(mBase, uint32(l1)+204)) = v227 | int32(1024)
										m.G0 = v13 + int32(320)
										return
									}
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v13)+4)) = v38
									*(*int32)(unsafe.Add(mBase, uint32(v13))) = l0
									*(*int32)(unsafe.Add(mBase, uint32(v13)+8)) = v13 + int32(64)
									F__serverLog(m, v199, int32(_a1016), v13)
									mBase = m.M
									v210 = m.ExcPending
									if v210 != 0 {
										return
									} else {
										F_sdsfree(m, v38)
										mBase = m.M
										v212 = m.ExcPending
										if v212 != 0 {
											return
										} else {
											v223 = *(*int32)(unsafe.Add(mBase, uint32(l1)+200))
											*(*int32)(unsafe.Add(mBase, uint32(l1)+200)) = v223 | int32(64)
											v227 = *(*int32)(unsafe.Add(mBase, uint32(l1)+204))
											*(*int32)(unsafe.Add(mBase, uint32(l1)+204)) = v227 | int32(1024)
											m.G0 = v13 + int32(320)
											return
										}
									}
								}
							}
						}
					}
				} else {
					v46 = int32(0)
					v47 = *(*int32)(unsafe.Add(mBase, _consts[595]))
					*(*int32)(unsafe.Add(mBase, uint32(v13)+71)) = v47
					v50 = *(*int64)(unsafe.Add(mBase, _consts[596]))
					*(*int64)(unsafe.Add(mBase, uint32(v13)+64)) = v50
					v185 = int32(1)
					v188 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+200)))
					if v188&v185 != 0 {
						v195 = v185
						v198 = v195
					} else {
						v191 = *(*int32)(unsafe.Add(mBase, uint32(l1)+216))
						if v191 != 0 {
							v193 = F_isImportSlotMigrationJob(m, v191)
							mBase = m.M
							v195 = v193
							v198 = v195
						} else {
							v198 = int32(0)
						}
					}
					if v198 != 0 {
						v199 = int32(3)
					} else {
						v199 = v185
					}
					v201 = *(*int32)(unsafe.Add(mBase, _consts[6]))
					if v199 < v201 {
						F_sdsfree(m, v38)
						mBase = m.M
						v212 = m.ExcPending
						if v212 != 0 {
							return
						} else {
							v223 = *(*int32)(unsafe.Add(mBase, uint32(l1)+200))
							*(*int32)(unsafe.Add(mBase, uint32(l1)+200)) = v223 | int32(64)
							v227 = *(*int32)(unsafe.Add(mBase, uint32(l1)+204))
							*(*int32)(unsafe.Add(mBase, uint32(l1)+204)) = v227 | int32(1024)
							m.G0 = v13 + int32(320)
							return
						}
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v13)+4)) = v38
						*(*int32)(unsafe.Add(mBase, uint32(v13))) = l0
						*(*int32)(unsafe.Add(mBase, uint32(v13)+8)) = v13 + int32(64)
						F__serverLog(m, v199, int32(_a1016), v13)
						mBase = m.M
						v210 = m.ExcPending
						if v210 != 0 {
							return
						} else {
							F_sdsfree(m, v38)
							mBase = m.M
							v212 = m.ExcPending
							if v212 != 0 {
								return
							} else {
								v223 = *(*int32)(unsafe.Add(mBase, uint32(l1)+200))
								*(*int32)(unsafe.Add(mBase, uint32(l1)+200)) = v223 | int32(64)
								v227 = *(*int32)(unsafe.Add(mBase, uint32(l1)+204))
								*(*int32)(unsafe.Add(mBase, uint32(l1)+204)) = v227 | int32(1024)
								m.G0 = v13 + int32(320)
								return
							}
						}
					}
				}
			}
		}
	} else {
		v20 = int32(1)
		v21 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+200)))
		if v21&v20 != 0 {
			v28 = v20
			v31 = v28
		} else {
			v24 = *(*int32)(unsafe.Add(mBase, uint32(l1)+216))
			if v24 != 0 {
				v26 = F_isImportSlotMigrationJob(m, v24)
				mBase = m.M
				v28 = v26
				v31 = v28
			} else {
				v31 = int32(0)
			}
		}
		if v31 == int32(0) {
			v223 = *(*int32)(unsafe.Add(mBase, uint32(l1)+200))
			*(*int32)(unsafe.Add(mBase, uint32(l1)+200)) = v223 | int32(64)
			v227 = *(*int32)(unsafe.Add(mBase, uint32(l1)+204))
			*(*int32)(unsafe.Add(mBase, uint32(l1)+204)) = v227 | int32(1024)
			m.G0 = v13 + int32(320)
			return
		} else {
			v34 = F_sdsempty(m)
			mBase = m.M
			v35 = m.ExcPending
			if v35 != 0 {
				return
			} else {
				v37 = *(*int32)(unsafe.Add(mBase, _consts[209]))
				v38 = F_catClientInfoString(m, v34, l1, v37)
				mBase = m.M
				v39 = m.ExcPending
				if v39 != 0 {
					return
				} else {
					v40 = int32(0)
					*(*uint8)(unsafe.Add(mBase, uint32(v13)+64)) = uint8(v40)
					v43 = *(*int32)(unsafe.Add(mBase, _consts[209]))
					if v43 == v40 {
						v52 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
						if v52 == int32(0) {
							v149 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+64)))
							if v149 == int32(0) {
							} else {
								v158 = v13 + int32(64)
								v159 = v149
								for {
									if base.Ui32(base.I32_extend8_s(v159)+int32(-32)) < base.Ui32(int32(95)) {
									} else {
										v169 = int32(46)
										*(*uint8)(unsafe.Add(mBase, uint32(v158))) = uint8(v169)
									}
									v171 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v158)+1)))
									if v171 != 0 {
										v158 = v158 + int32(1)
										v159 = v171
										continue
									} else {
										break
									}
									break
								}
							}
							v185 = int32(1)
							v188 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+200)))
							if v188&v185 != 0 {
								v195 = v185
								v198 = v195
							} else {
								v191 = *(*int32)(unsafe.Add(mBase, uint32(l1)+216))
								if v191 != 0 {
									v193 = F_isImportSlotMigrationJob(m, v191)
									mBase = m.M
									v195 = v193
									v198 = v195
								} else {
									v198 = int32(0)
								}
							}
							if v198 != 0 {
								v199 = int32(3)
							} else {
								v199 = v185
							}
							v201 = *(*int32)(unsafe.Add(mBase, _consts[6]))
							if v199 < v201 {
								F_sdsfree(m, v38)
								mBase = m.M
								v212 = m.ExcPending
								if v212 != 0 {
									return
								} else {
									v223 = *(*int32)(unsafe.Add(mBase, uint32(l1)+200))
									*(*int32)(unsafe.Add(mBase, uint32(l1)+200)) = v223 | int32(64)
									v227 = *(*int32)(unsafe.Add(mBase, uint32(l1)+204))
									*(*int32)(unsafe.Add(mBase, uint32(l1)+204)) = v227 | int32(1024)
									m.G0 = v13 + int32(320)
									return
								}
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v13)+4)) = v38
								*(*int32)(unsafe.Add(mBase, uint32(v13))) = l0
								*(*int32)(unsafe.Add(mBase, uint32(v13)+8)) = v13 + int32(64)
								F__serverLog(m, v199, int32(_a1016), v13)
								mBase = m.M
								v210 = m.ExcPending
								if v210 != 0 {
									return
								} else {
									F_sdsfree(m, v38)
									mBase = m.M
									v212 = m.ExcPending
									if v212 != 0 {
										return
									} else {
										v223 = *(*int32)(unsafe.Add(mBase, uint32(l1)+200))
										*(*int32)(unsafe.Add(mBase, uint32(l1)+200)) = v223 | int32(64)
										v227 = *(*int32)(unsafe.Add(mBase, uint32(l1)+204))
										*(*int32)(unsafe.Add(mBase, uint32(l1)+204)) = v227 | int32(1024)
										m.G0 = v13 + int32(320)
										return
									}
								}
							}
						} else {
							v58 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v52+int32(-1)))))
							v60 = v58 & int32(7)
							switch v60 {
							case 0:
								v75 = int32(base.Ui32(v58) >> (uint(int32(3)) % 32))
							case 1:
								v65 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v52+int32(-3)))))
								v75 = v65
							case 2:
								v68 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v52+int32(-5)))))
								v75 = v68
							case 3:
								v71 = *(*int32)(unsafe.Add(mBase, uint32(v52+int32(-9))))
								v75 = v71
							case 4:
								v74 = *(*int32)(unsafe.Add(mBase, uint32(v52+int32(-17))))
								v75 = v74
							default:
								v75 = int32(0)
							}
							v76 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
							v77 = v52 + v76
							if base.Ui32(int32(127)) < base.Ui32(v75-v76) {
								switch v60 {
								case 0:
									v91 = int32(base.Ui32(v58) >> (uint(int32(3)) % 32))
									v122 = v91
									v123 = v91 - v76 + int32(-128)
								case 1:
									v97 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v52+int32(-3)))))
									v122 = v97
									v123 = v97 - v76 + int32(-128)
								case 2:
									v103 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v52+int32(-5)))))
									v122 = v103
									v123 = v103 - v76 + int32(-128)
								case 3:
									v109 = *(*int32)(unsafe.Add(mBase, uint32(v52+int32(-9))))
									v122 = v109
									v123 = v109 - v76 + int32(-128)
								case 4:
									v115 = *(*int32)(unsafe.Add(mBase, uint32(v52+int32(-17))))
									v122 = v115
									v123 = v115 - v76 + int32(-128)
								default:
									v122 = int32(0)
									v123 = int32(-128) - v76
								}
								*(*int32)(unsafe.Add(mBase, uint32(v13+int32(48)))) = v52 + v122 + int32(-64)
								v130 = int32(64)
								*(*int32)(unsafe.Add(mBase, uint32(v13)+44)) = v130
								*(*int32)(unsafe.Add(mBase, uint32(v13)+40)) = v123
								*(*int32)(unsafe.Add(mBase, uint32(v13)+36)) = v77
								*(*int32)(unsafe.Add(mBase, uint32(v13)+32)) = v130
								v142 = F_snprintf(m, v13+v130, int32(256), int32(_a1017), v13+int32(32))
								mBase = m.M
								v143 = m.ExcPending
								if v143 != 0 {
									return
								} else {
									v149 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+64)))
									if v149 == int32(0) {
									} else {
										v158 = v13 + int32(64)
										v159 = v149
										for {
											if base.Ui32(base.I32_extend8_s(v159)+int32(-32)) < base.Ui32(int32(95)) {
											} else {
												v169 = int32(46)
												*(*uint8)(unsafe.Add(mBase, uint32(v158))) = uint8(v169)
											}
											v171 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v158)+1)))
											if v171 != 0 {
												v158 = v158 + int32(1)
												v159 = v171
												continue
											} else {
												break
											}
											break
										}
									}
									v185 = int32(1)
									v188 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+200)))
									if v188&v185 != 0 {
										v195 = v185
										v198 = v195
									} else {
										v191 = *(*int32)(unsafe.Add(mBase, uint32(l1)+216))
										if v191 != 0 {
											v193 = F_isImportSlotMigrationJob(m, v191)
											mBase = m.M
											v195 = v193
											v198 = v195
										} else {
											v198 = int32(0)
										}
									}
									if v198 != 0 {
										v199 = int32(3)
									} else {
										v199 = v185
									}
									v201 = *(*int32)(unsafe.Add(mBase, _consts[6]))
									if v199 < v201 {
										F_sdsfree(m, v38)
										mBase = m.M
										v212 = m.ExcPending
										if v212 != 0 {
											return
										} else {
											v223 = *(*int32)(unsafe.Add(mBase, uint32(l1)+200))
											*(*int32)(unsafe.Add(mBase, uint32(l1)+200)) = v223 | int32(64)
											v227 = *(*int32)(unsafe.Add(mBase, uint32(l1)+204))
											*(*int32)(unsafe.Add(mBase, uint32(l1)+204)) = v227 | int32(1024)
											m.G0 = v13 + int32(320)
											return
										}
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v13)+4)) = v38
										*(*int32)(unsafe.Add(mBase, uint32(v13))) = l0
										*(*int32)(unsafe.Add(mBase, uint32(v13)+8)) = v13 + int32(64)
										F__serverLog(m, v199, int32(_a1016), v13)
										mBase = m.M
										v210 = m.ExcPending
										if v210 != 0 {
											return
										} else {
											F_sdsfree(m, v38)
											mBase = m.M
											v212 = m.ExcPending
											if v212 != 0 {
												return
											} else {
												v223 = *(*int32)(unsafe.Add(mBase, uint32(l1)+200))
												*(*int32)(unsafe.Add(mBase, uint32(l1)+200)) = v223 | int32(64)
												v227 = *(*int32)(unsafe.Add(mBase, uint32(l1)+204))
												*(*int32)(unsafe.Add(mBase, uint32(l1)+204)) = v227 | int32(1024)
												m.G0 = v13 + int32(320)
												return
											}
										}
									}
								}
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v13)+16)) = v77
								v88 = F_snprintf(m, v13+int32(64), int32(256), int32(_a1018), v13+int32(16))
								mBase = m.M
								v89 = m.ExcPending
								if v89 != 0 {
									return
								} else {
									v149 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+64)))
									if v149 == int32(0) {
									} else {
										v158 = v13 + int32(64)
										v159 = v149
										for {
											if base.Ui32(base.I32_extend8_s(v159)+int32(-32)) < base.Ui32(int32(95)) {
											} else {
												v169 = int32(46)
												*(*uint8)(unsafe.Add(mBase, uint32(v158))) = uint8(v169)
											}
											v171 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v158)+1)))
											if v171 != 0 {
												v158 = v158 + int32(1)
												v159 = v171
												continue
											} else {
												break
											}
											break
										}
									}
									v185 = int32(1)
									v188 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+200)))
									if v188&v185 != 0 {
										v195 = v185
										v198 = v195
									} else {
										v191 = *(*int32)(unsafe.Add(mBase, uint32(l1)+216))
										if v191 != 0 {
											v193 = F_isImportSlotMigrationJob(m, v191)
											mBase = m.M
											v195 = v193
											v198 = v195
										} else {
											v198 = int32(0)
										}
									}
									if v198 != 0 {
										v199 = int32(3)
									} else {
										v199 = v185
									}
									v201 = *(*int32)(unsafe.Add(mBase, _consts[6]))
									if v199 < v201 {
										F_sdsfree(m, v38)
										mBase = m.M
										v212 = m.ExcPending
										if v212 != 0 {
											return
										} else {
											v223 = *(*int32)(unsafe.Add(mBase, uint32(l1)+200))
											*(*int32)(unsafe.Add(mBase, uint32(l1)+200)) = v223 | int32(64)
											v227 = *(*int32)(unsafe.Add(mBase, uint32(l1)+204))
											*(*int32)(unsafe.Add(mBase, uint32(l1)+204)) = v227 | int32(1024)
											m.G0 = v13 + int32(320)
											return
										}
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v13)+4)) = v38
										*(*int32)(unsafe.Add(mBase, uint32(v13))) = l0
										*(*int32)(unsafe.Add(mBase, uint32(v13)+8)) = v13 + int32(64)
										F__serverLog(m, v199, int32(_a1016), v13)
										mBase = m.M
										v210 = m.ExcPending
										if v210 != 0 {
											return
										} else {
											F_sdsfree(m, v38)
											mBase = m.M
											v212 = m.ExcPending
											if v212 != 0 {
												return
											} else {
												v223 = *(*int32)(unsafe.Add(mBase, uint32(l1)+200))
												*(*int32)(unsafe.Add(mBase, uint32(l1)+200)) = v223 | int32(64)
												v227 = *(*int32)(unsafe.Add(mBase, uint32(l1)+204))
												*(*int32)(unsafe.Add(mBase, uint32(l1)+204)) = v227 | int32(1024)
												m.G0 = v13 + int32(320)
												return
											}
										}
									}
								}
							}
						}
					} else {
						v46 = int32(0)
						v47 = *(*int32)(unsafe.Add(mBase, _consts[595]))
						*(*int32)(unsafe.Add(mBase, uint32(v13)+71)) = v47
						v50 = *(*int64)(unsafe.Add(mBase, _consts[596]))
						*(*int64)(unsafe.Add(mBase, uint32(v13)+64)) = v50
						v185 = int32(1)
						v188 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+200)))
						if v188&v185 != 0 {
							v195 = v185
							v198 = v195
						} else {
							v191 = *(*int32)(unsafe.Add(mBase, uint32(l1)+216))
							if v191 != 0 {
								v193 = F_isImportSlotMigrationJob(m, v191)
								mBase = m.M
								v195 = v193
								v198 = v195
							} else {
								v198 = int32(0)
							}
						}
						if v198 != 0 {
							v199 = int32(3)
						} else {
							v199 = v185
						}
						v201 = *(*int32)(unsafe.Add(mBase, _consts[6]))
						if v199 < v201 {
							F_sdsfree(m, v38)
							mBase = m.M
							v212 = m.ExcPending
							if v212 != 0 {
								return
							} else {
								v223 = *(*int32)(unsafe.Add(mBase, uint32(l1)+200))
								*(*int32)(unsafe.Add(mBase, uint32(l1)+200)) = v223 | int32(64)
								v227 = *(*int32)(unsafe.Add(mBase, uint32(l1)+204))
								*(*int32)(unsafe.Add(mBase, uint32(l1)+204)) = v227 | int32(1024)
								m.G0 = v13 + int32(320)
								return
							}
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v13)+4)) = v38
							*(*int32)(unsafe.Add(mBase, uint32(v13))) = l0
							*(*int32)(unsafe.Add(mBase, uint32(v13)+8)) = v13 + int32(64)
							F__serverLog(m, v199, int32(_a1016), v13)
							mBase = m.M
							v210 = m.ExcPending
							if v210 != 0 {
								return
							} else {
								F_sdsfree(m, v38)
								mBase = m.M
								v212 = m.ExcPending
								if v212 != 0 {
									return
								} else {
									v223 = *(*int32)(unsafe.Add(mBase, uint32(l1)+200))
									*(*int32)(unsafe.Add(mBase, uint32(l1)+200)) = v223 | int32(64)
									v227 = *(*int32)(unsafe.Add(mBase, uint32(l1)+204))
									*(*int32)(unsafe.Add(mBase, uint32(l1)+204)) = v227 | int32(1024)
									m.G0 = v13 + int32(320)
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
func F_setbitCommand(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v35 int32
	_ = v35
	var v36 int64
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v60 int32
	_ = v60
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
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
	var v85 int64
	_ = v85
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	v2 = int32(0)
	v11 = m.G0
	v13 = v11 - int32(16)
	m.G0 = v13
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v16 = *(*int32)(unsafe.Add(mBase, uint32(v15)+8))
	v21 = F_getBitOffsetFromArgument(m, l0, v16, v13+int32(8), v2, v2)
	mBase = m.M
	v22 = m.ExcPending
	if v22 != 0 {
		return
	} else {
		if v21 != 0 {
			m.G0 = v13 + int32(16)
			return
		} else {
			v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
			v24 = *(*int32)(unsafe.Add(mBase, uint32(v23)+12))
			v28 = F_getLongFromObjectOrReply(m, l0, v24, v13+int32(4), int32(_a183))
			mBase = m.M
			v29 = m.ExcPending
			if v29 != 0 {
				return
			} else {
				if v28 != 0 {
					m.G0 = v13 + int32(16)
					return
				} else {
					v30 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
					if base.Ui32(v30) < base.Ui32(int32(2)) {
						v36 = *(*int64)(unsafe.Add(mBase, uint32(v13)+8))
						v37 = F_lookupStringForBitCommand(m, l0, v36, v13)
						mBase = m.M
						v38 = m.ExcPending
						if v38 != 0 {
							return
						} else {
							if v37 == int32(0) {
								m.G0 = v13 + int32(16)
								return
							} else {
								v46 = (base.I32_wrap_i64(v36) ^ int32(-1)) & int32(7)
								v47 = int32(1) << (uint(v46) % 32)
								v48 = F_objectGetVal(m, v37)
								mBase = m.M
								v51 = base.I32_wrap_i64(int64(base.Ui64(v36) >> (uint(int64(3)) % 64)))
								v53 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v48+v51))))
								v54 = v47 & v53
								v55 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
								v56 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
								if v56 != 0 {
									v60 = F_objectGetVal(m, v37)
									mBase = m.M
									v68 = v55&int32(1)<<(uint(v46)%32) | v53&(v47^int32(-1))
									*(*uint8)(unsafe.Add(mBase, uint32(v60+v51))) = uint8(v68)
									v70 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
									v71 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
									v72 = *(*int32)(unsafe.Add(mBase, uint32(v71)+4))
									F_signalModifiedKey(m, l0, v70, v72)
									mBase = m.M
									v74 = m.ExcPending
									if v74 != 0 {
										return
									} else {
										v77 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
										v78 = *(*int32)(unsafe.Add(mBase, uint32(v77)+4))
										v79 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
										v80 = *(*int32)(unsafe.Add(mBase, uint32(v79)+28))
										F_notifyKeyspaceEvent(m, int32(8), int32(_a184), v78, v80)
										mBase = m.M
										v82 = m.ExcPending
										if v82 != 0 {
											return
										} else {
											v83 = int32(_a44)
											v85 = *(*int64)(unsafe.Add(mBase, _consts[83]))
											*(*int64)(unsafe.Add(mBase, _consts[83])) = v85 + int64(1)
											if v54 != 0 {
												v92 = int32(16)
											} else {
												v92 = int32(12)
											}
											v94 = *(*int32)(unsafe.Add(mBase, uint32(v92)+uint32(_consts[84])))
											F_addReply(m, l0, v94)
											mBase = m.M
											v96 = m.ExcPending
											if v96 != 0 {
												return
											} else {
												m.G0 = v13 + int32(16)
												return
											}
										}
									}
								} else {
									if v55 == base.B2i32(v54 != int32(0)) {
										if v54 != 0 {
											v92 = int32(16)
										} else {
											v92 = int32(12)
										}
										v94 = *(*int32)(unsafe.Add(mBase, uint32(v92)+uint32(_consts[84])))
										F_addReply(m, l0, v94)
										mBase = m.M
										v96 = m.ExcPending
										if v96 != 0 {
											return
										} else {
											m.G0 = v13 + int32(16)
											return
										}
									} else {
										v60 = F_objectGetVal(m, v37)
										mBase = m.M
										v68 = v55&int32(1)<<(uint(v46)%32) | v53&(v47^int32(-1))
										*(*uint8)(unsafe.Add(mBase, uint32(v60+v51))) = uint8(v68)
										v70 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
										v71 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
										v72 = *(*int32)(unsafe.Add(mBase, uint32(v71)+4))
										F_signalModifiedKey(m, l0, v70, v72)
										mBase = m.M
										v74 = m.ExcPending
										if v74 != 0 {
											return
										} else {
											v77 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
											v78 = *(*int32)(unsafe.Add(mBase, uint32(v77)+4))
											v79 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
											v80 = *(*int32)(unsafe.Add(mBase, uint32(v79)+28))
											F_notifyKeyspaceEvent(m, int32(8), int32(_a184), v78, v80)
											mBase = m.M
											v82 = m.ExcPending
											if v82 != 0 {
												return
											} else {
												v83 = int32(_a44)
												v85 = *(*int64)(unsafe.Add(mBase, _consts[83]))
												*(*int64)(unsafe.Add(mBase, _consts[83])) = v85 + int64(1)
												if v54 != 0 {
													v92 = int32(16)
												} else {
													v92 = int32(12)
												}
												v94 = *(*int32)(unsafe.Add(mBase, uint32(v92)+uint32(_consts[84])))
												F_addReply(m, l0, v94)
												mBase = m.M
												v96 = m.ExcPending
												if v96 != 0 {
													return
												} else {
													m.G0 = v13 + int32(16)
													return
												}
											}
										}
									}
								}
							}
						}
					} else {
						F_addReplyError(m, l0, int32(_a183))
						mBase = m.M
						v35 = m.ExcPending
						if v35 != 0 {
							return
						} else {
							m.G0 = v13 + int32(16)
							return
						}
					}
				}
			}
		}
	}
}
func F_setitimer(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v10 float64
	_ = v10
	var v13 float64
	_ = v13
	var v17 int32
	_ = v17
	var v20 float64
	_ = v20
	var v21 float64
	_ = v21
	var v25 float64
	_ = v25
	var v27 float64
	_ = v27
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v38 float64
	_ = v38
	var v44 int64
	_ = v44
	var v46 int64
	_ = v46
	var v50 float64
	_ = v50
	var v52 float64
	_ = v52
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v63 float64
	_ = v63
	var v69 int64
	_ = v69
	var v71 int64
	_ = v71
	var v73 int32
	_ = v73
	var v74 int64
	_ = v74
	var v80 int32
	_ = v80
	var v83 int64
	_ = v83
	var v84 int64
	_ = v84
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v94 int64
	_ = v94
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v108 int32
	_ = v108
	var v111 int64
	_ = v111
	var v121 float64
	_ = v121
	var v123 int32
	_ = v123
	var v126 int32
	_ = v126
	if int32(2) < l0 {
		v126 = int32(28)
	} else {
		v10 = m.Env.Emscripten_get_now(m)
		mBase = m.M
		if l2 == int32(0) {
		} else {
			v13 = float64(0)
			v17 = l0 << (uint(int32(3)) % 32)
			v20 = *(*float64)(unsafe.Add(mBase, uint32(v17)+uint32(_consts[1081])))
			v21 = base.F64_sub(v20, v10)
			if base.F64_gt(v21, v13) != 0 {
				v25 = v21
			} else {
				v25 = v13
			}
			v27 = base.F64_mul(v25, float64(1000))
			if base.F64_lt(base.F64_abs(v27), float64(2.147483648e+09)) == int32(0) {
				v35 = int32(-2147483648)
			} else {
				v33 = base.I32_trunc_f64_s(v27)
				v35 = v33
			}
			*(*int32)(unsafe.Add(mBase, uint32(l2)+24)) = v35
			v38 = base.F64_div(v25, float64(1000))
			if base.F64_lt(base.F64_abs(v38), float64(9.223372036854776e+18)) == int32(0) {
				v46 = int64(-9223372036854775807 - 1)
			} else {
				v44 = base.I64_trunc_f64_s(v38)
				v46 = v44
			}
			*(*int64)(unsafe.Add(mBase, uint32(l2)+16)) = v46
			v50 = *(*float64)(unsafe.Add(mBase, uint32(v17)+uint32(_consts[1082])))
			v52 = base.F64_mul(v50, float64(1000))
			if base.F64_lt(base.F64_abs(v52), float64(2.147483648e+09)) == int32(0) {
				v60 = int32(-2147483648)
			} else {
				v58 = base.I32_trunc_f64_s(v52)
				v60 = v58
			}
			*(*int32)(unsafe.Add(mBase, uint32(l2)+8)) = v60
			v63 = base.F64_div(v50, float64(1000))
			if base.F64_lt(base.F64_abs(v63), float64(9.223372036854776e+18)) == int32(0) {
				v71 = int64(-9223372036854775807 - 1)
			} else {
				v69 = base.I64_trunc_f64_s(v63)
				v71 = v69
			}
			*(*int64)(unsafe.Add(mBase, uint32(l2))) = v71
		}
		v73 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
		v74 = *(*int64)(unsafe.Add(mBase, uint32(l1)+16))
		if v74 != int64(0) {
			v80 = l0 << (uint(int32(3)) % 32)
			v83 = int64(1000)
			v84 = v74 * v83
			v87 = int32(1000)
			v88 = base.I32_div_s(v73, v87)
			*(*float64)(unsafe.Add(mBase, uint32(v80)+uint32(_consts[1081]))) = base.F64_add(base.F64_add(v10, base.F64_convert_i64_s(v84)), base.F64_convert_i32_s(v88))
			v94 = *(*int64)(unsafe.Add(mBase, uint32(l1)))
			v97 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
			v99 = base.I32_div_s(v97, v87)
			*(*float64)(unsafe.Add(mBase, uint32(v80)+uint32(_consts[1082]))) = base.F64_convert_i64_s(v94*v83 + base.I64_extend_i32_s(v99))
			v121 = base.F64_convert_i64_s(v84 + base.I64_extend_i32_s(v88))
		} else {
			if v73 == int32(0) {
				v108 = l0 << (uint(int32(3)) % 32)
				v111 = int64(0)
				*(*int64)(unsafe.Add(mBase, uint32(v108)+uint32(_consts[1081]))) = v111
				*(*int64)(unsafe.Add(mBase, uint32(v108)+uint32(_consts[1082]))) = v111
				v121 = float64(0)
			} else {
				v80 = l0 << (uint(int32(3)) % 32)
				v83 = int64(1000)
				v84 = v74 * v83
				v87 = int32(1000)
				v88 = base.I32_div_s(v73, v87)
				*(*float64)(unsafe.Add(mBase, uint32(v80)+uint32(_consts[1081]))) = base.F64_add(base.F64_add(v10, base.F64_convert_i64_s(v84)), base.F64_convert_i32_s(v88))
				v94 = *(*int64)(unsafe.Add(mBase, uint32(l1)))
				v97 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
				v99 = base.I32_div_s(v97, v87)
				*(*float64)(unsafe.Add(mBase, uint32(v80)+uint32(_consts[1082]))) = base.F64_convert_i64_s(v94*v83 + base.I64_extend_i32_s(v99))
				v121 = base.F64_convert_i64_s(v84 + base.I64_extend_i32_s(v88))
			}
		}
		v123 = m.Env.X_setitimer_js(m, l0, v121)
		mBase = m.M
		v126 = v123
	}
	return v126
}
func F_setnodevector(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v67 int32
	_ = v67
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v77 int32
	_ = v77
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v88 int32
	_ = v88
	var v92 int32
	_ = v92
	var v97 int32
	_ = v97
	var v100 int32
	_ = v100
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v104 int64
	_ = v104
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v141 int32
	_ = v141
	var v145 int32
	_ = v145
	var v159 int32
	_ = v159
	var v162 int32
	_ = v162
	var v167 int32
	_ = v167
	var v170 int32
	_ = v170
	var v175 int32
	_ = v175
	var v178 int32
	_ = v178
	var v190 int32
	_ = v190
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	if l2 != 0 {
		v17 = int32(-1)
		v18 = l2 + v17
		if base.Ui32(int32(256)) <= base.Ui32(v18) {
			v25 = v18
			v26 = v17
			for {
				v29 = int32(8)
				v30 = v26 + v29
				v34 = int32(base.Ui32(v25) >> (uint(v29) % 32))
				if base.Ui32(int32(65535)) < base.Ui32(v25) {
					v25 = v34
					v26 = v30
					continue
				} else {
					break
				}
				break
			}
			v36 = v30
			v37 = v34
		} else {
			v36 = v17
			v37 = v18
		}
		v39 = m.G3
		v43 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v39+int32(_a2057)+v37))))
		v44 = v36 + v43
		v46 = v44 + int32(1)
		if int32(25) < v44 {
			v53 = m.G3
			F_luaG_runerror(m, l0, v53+int32(_a2058), int32(0))
			mBase = m.M
			v58 = m.ExcPending
			if v58 != 0 {
				return
			} else {
				v59 = F_luaM_toobig(m, l0)
				mBase = m.M
				v60 = m.ExcPending
				if v60 != 0 {
					return
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(l1)+24)) = v59
					if v46 == int32(31) {
						v193 = v59
						v194 = v46
						v195 = int32(-2147483648)
					} else {
						v77 = l1 + int32(24)
						v81 = int32(1) << (uint(v46) % 32)
						v82 = int32(0)
						if base.Ui32(v46) < base.Ui32(int32(2)) {
							v145 = v82
						} else {
							v88 = int32(0)
							v92 = v88
							v97 = v88
							for {
								v100 = *(*int32)(unsafe.Add(mBase, uint32(v77)))
								v102 = v92 << (uint(int32(5)) % 32)
								v103 = v100 + v102
								v104 = int64(0)
								*(*int64)(unsafe.Add(mBase, uint32(v103)+24)) = v104
								v106 = int32(0)
								*(*int32)(unsafe.Add(mBase, uint32(v103)+8)) = v106
								v108 = *(*int32)(unsafe.Add(mBase, uint32(v77)))
								v109 = v108 + v102
								*(*int64)(unsafe.Add(mBase, uint32(v109+int32(56)))) = v104
								*(*int32)(unsafe.Add(mBase, uint32(v109+int32(40)))) = v106
								v118 = *(*int32)(unsafe.Add(mBase, uint32(v77)))
								v119 = v118 + v102
								*(*int64)(unsafe.Add(mBase, uint32(v119+int32(88)))) = v104
								*(*int32)(unsafe.Add(mBase, uint32(v119+int32(72)))) = v106
								v128 = *(*int32)(unsafe.Add(mBase, uint32(v77)))
								v129 = v128 + v102
								*(*int64)(unsafe.Add(mBase, uint32(v129+int32(120)))) = v104
								*(*int32)(unsafe.Add(mBase, uint32(v129+int32(104)))) = v106
								v138 = int32(4)
								v139 = v92 + v138
								v141 = v97 + v138
								if v141 != v81&int32(-4) {
									v92 = v139
									v97 = v141
									continue
								} else {
									break
								}
								break
							}
							v145 = v139
						}
						if base.Ui32(int32(1)) < base.Ui32(v46) {
						} else {
							v159 = v145
							v162 = v82
							for {
								v167 = *(*int32)(unsafe.Add(mBase, uint32(v77)))
								v170 = v167 + v159<<(uint(int32(5))%32)
								*(*int64)(unsafe.Add(mBase, uint32(v170)+24)) = int64(0)
								*(*int32)(unsafe.Add(mBase, uint32(v170)+8)) = int32(0)
								v175 = int32(1)
								v178 = v162 + v175
								if v178 != v81&int32(3) {
									v159 = v159 + v175
									v162 = v178
									continue
								} else {
									break
								}
								break
							}
						}
						v190 = *(*int32)(unsafe.Add(mBase, uint32(v77)))
						v193 = v190
						v194 = v46
						v195 = v81
					}
					*(*uint8)(unsafe.Add(mBase, uint32(l1)+12)) = uint8(v194)
					*(*int32)(unsafe.Add(mBase, uint32(l1)+28)) = v193 + v195<<(uint(int32(5))%32)
					return
				}
			}
		} else {
			if base.Ui32(v46) < base.Ui32(int32(27)) {
				v67 = int32(0)
				v71 = F_luaM_realloc_(m, l0, v67, v67, int32(32)<<(uint(v46)%32))
				mBase = m.M
				v72 = m.ExcPending
				if v72 != 0 {
					return
				} else {
					v73 = v71
					*(*int32)(unsafe.Add(mBase, uint32(l1)+24)) = v73
					v77 = l1 + int32(24)
					v81 = int32(1) << (uint(v46) % 32)
					v82 = int32(0)
					if base.Ui32(v46) < base.Ui32(int32(2)) {
						v145 = v82
					} else {
						v88 = int32(0)
						v92 = v88
						v97 = v88
						for {
							v100 = *(*int32)(unsafe.Add(mBase, uint32(v77)))
							v102 = v92 << (uint(int32(5)) % 32)
							v103 = v100 + v102
							v104 = int64(0)
							*(*int64)(unsafe.Add(mBase, uint32(v103)+24)) = v104
							v106 = int32(0)
							*(*int32)(unsafe.Add(mBase, uint32(v103)+8)) = v106
							v108 = *(*int32)(unsafe.Add(mBase, uint32(v77)))
							v109 = v108 + v102
							*(*int64)(unsafe.Add(mBase, uint32(v109+int32(56)))) = v104
							*(*int32)(unsafe.Add(mBase, uint32(v109+int32(40)))) = v106
							v118 = *(*int32)(unsafe.Add(mBase, uint32(v77)))
							v119 = v118 + v102
							*(*int64)(unsafe.Add(mBase, uint32(v119+int32(88)))) = v104
							*(*int32)(unsafe.Add(mBase, uint32(v119+int32(72)))) = v106
							v128 = *(*int32)(unsafe.Add(mBase, uint32(v77)))
							v129 = v128 + v102
							*(*int64)(unsafe.Add(mBase, uint32(v129+int32(120)))) = v104
							*(*int32)(unsafe.Add(mBase, uint32(v129+int32(104)))) = v106
							v138 = int32(4)
							v139 = v92 + v138
							v141 = v97 + v138
							if v141 != v81&int32(-4) {
								v92 = v139
								v97 = v141
								continue
							} else {
								break
							}
							break
						}
						v145 = v139
					}
					if base.Ui32(int32(1)) < base.Ui32(v46) {
					} else {
						v159 = v145
						v162 = v82
						for {
							v167 = *(*int32)(unsafe.Add(mBase, uint32(v77)))
							v170 = v167 + v159<<(uint(int32(5))%32)
							*(*int64)(unsafe.Add(mBase, uint32(v170)+24)) = int64(0)
							*(*int32)(unsafe.Add(mBase, uint32(v170)+8)) = int32(0)
							v175 = int32(1)
							v178 = v162 + v175
							if v178 != v81&int32(3) {
								v159 = v159 + v175
								v162 = v178
								continue
							} else {
								break
							}
							break
						}
					}
					v190 = *(*int32)(unsafe.Add(mBase, uint32(v77)))
					v193 = v190
					v194 = v46
					v195 = v81
					*(*uint8)(unsafe.Add(mBase, uint32(l1)+12)) = uint8(v194)
					*(*int32)(unsafe.Add(mBase, uint32(l1)+28)) = v193 + v195<<(uint(int32(5))%32)
					return
				}
			} else {
				v51 = F_luaM_toobig(m, l0)
				mBase = m.M
				v52 = m.ExcPending
				if v52 != 0 {
					return
				} else {
					v73 = v51
					*(*int32)(unsafe.Add(mBase, uint32(l1)+24)) = v73
					v77 = l1 + int32(24)
					v81 = int32(1) << (uint(v46) % 32)
					v82 = int32(0)
					if base.Ui32(v46) < base.Ui32(int32(2)) {
						v145 = v82
					} else {
						v88 = int32(0)
						v92 = v88
						v97 = v88
						for {
							v100 = *(*int32)(unsafe.Add(mBase, uint32(v77)))
							v102 = v92 << (uint(int32(5)) % 32)
							v103 = v100 + v102
							v104 = int64(0)
							*(*int64)(unsafe.Add(mBase, uint32(v103)+24)) = v104
							v106 = int32(0)
							*(*int32)(unsafe.Add(mBase, uint32(v103)+8)) = v106
							v108 = *(*int32)(unsafe.Add(mBase, uint32(v77)))
							v109 = v108 + v102
							*(*int64)(unsafe.Add(mBase, uint32(v109+int32(56)))) = v104
							*(*int32)(unsafe.Add(mBase, uint32(v109+int32(40)))) = v106
							v118 = *(*int32)(unsafe.Add(mBase, uint32(v77)))
							v119 = v118 + v102
							*(*int64)(unsafe.Add(mBase, uint32(v119+int32(88)))) = v104
							*(*int32)(unsafe.Add(mBase, uint32(v119+int32(72)))) = v106
							v128 = *(*int32)(unsafe.Add(mBase, uint32(v77)))
							v129 = v128 + v102
							*(*int64)(unsafe.Add(mBase, uint32(v129+int32(120)))) = v104
							*(*int32)(unsafe.Add(mBase, uint32(v129+int32(104)))) = v106
							v138 = int32(4)
							v139 = v92 + v138
							v141 = v97 + v138
							if v141 != v81&int32(-4) {
								v92 = v139
								v97 = v141
								continue
							} else {
								break
							}
							break
						}
						v145 = v139
					}
					if base.Ui32(int32(1)) < base.Ui32(v46) {
					} else {
						v159 = v145
						v162 = v82
						for {
							v167 = *(*int32)(unsafe.Add(mBase, uint32(v77)))
							v170 = v167 + v159<<(uint(int32(5))%32)
							*(*int64)(unsafe.Add(mBase, uint32(v170)+24)) = int64(0)
							*(*int32)(unsafe.Add(mBase, uint32(v170)+8)) = int32(0)
							v175 = int32(1)
							v178 = v162 + v175
							if v178 != v81&int32(3) {
								v159 = v159 + v175
								v162 = v178
								continue
							} else {
								break
							}
							break
						}
					}
					v190 = *(*int32)(unsafe.Add(mBase, uint32(v77)))
					v193 = v190
					v194 = v46
					v195 = v81
					*(*uint8)(unsafe.Add(mBase, uint32(l1)+12)) = uint8(v194)
					*(*int32)(unsafe.Add(mBase, uint32(l1)+28)) = v193 + v195<<(uint(int32(5))%32)
					return
				}
			}
		}
	} else {
		v11 = m.G3
		v13 = v11 + int32(_a2056)
		*(*int32)(unsafe.Add(mBase, uint32(l1)+24)) = v13
		v15 = int32(0)
		v193 = v13
		v194 = v15
		v195 = v15
		*(*uint8)(unsafe.Add(mBase, uint32(l1)+12)) = uint8(v194)
		*(*int32)(unsafe.Add(mBase, uint32(l1)+28)) = v193 + v195<<(uint(int32(5))%32)
		return
	}
}
func F_setrlimit(m *base.Module, l0 int32, l1 int32) int32 {
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	m.G0 = v9 + int32(16)
	return int32(0)
}
func F_setsid(m *base.Module) int32 {
	return int32(0)
}
func F_setupDebugSigHandlers(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v28 int32
	_ = v28
	var v34 int32
	_ = v34
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v46 int32
	_ = v46
	var v51 int32
	_ = v51
	var v56 int32
	_ = v56
	var v61 int32
	_ = v61
	var v74 int32
	_ = v74
	var v78 int32
	_ = v78
	var v101 int32
	_ = v101
	v2 = m.G0
	v3 = int32(144)
	v4 = v2 - v3
	m.G0 = v4
	v7 = m.G0
	v9 = v7 - v3
	m.G0 = v9
	v12 = *(*int32)(unsafe.Add(mBase, _consts[406]))
	if v12 != 0 {
	} else {
		v13 = int32(_a761)
		v14 = F_pthread_mutexattr_init(m, v13)
		mBase = m.M
		v17 = F_pthread_mutexattr_settype(m, v13, int32(2))
		mBase = m.M
		v20 = F_pthread_mutex_init(m, int32(_a762), v13)
		mBase = m.M
		*(*int32)(unsafe.Add(mBase, _consts[406])) = int32(1)
	}
	v28 = F_sigemptyset(m, v9+int32(8))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v9)+4)) = int32(519)
	*(*int32)(unsafe.Add(mBase, uint32(v9)+136)) = int32(1073741828)
	v34 = *(*int32)(unsafe.Add(mBase, _consts[407]))
	if v34 == int32(0) {
	} else {
		v38 = int32(4)
		v39 = v9 + v38
		v40 = int32(0)
		v41 = F___sigaction(m, int32(11), v39, v40)
		mBase = m.M
		v46 = F___sigaction(m, int32(7), v39, v40)
		mBase = m.M
		v51 = F___sigaction(m, int32(8), v39, v40)
		mBase = m.M
		v56 = F___sigaction(m, v38, v39, v40)
		mBase = m.M
		v61 = F___sigaction(m, int32(6), v39, v40)
		mBase = m.M
	}
	m.G0 = v9 + int32(144)
	*(*int64)(unsafe.Add(mBase, uint32(v4+int32(8)))) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v4)+4)) = int32(518)
	v74 = int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v4)+136)) = v74
	v78 = v4 + v74
	if v78 == int32(0) {
	} else {
		v101 = F___memcpy(m, int32(9119144), v78, int32(140))
		mBase = m.M
	}
	m.G0 = v4 + int32(144)
	return
}
func F_sha256_update(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+80))
	m.Env.Vkmem_hash_update(m, v4, l1, l2)
	mBase = m.M
	return
}
func F_showLatestBacklog(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
	var v33 int32
	_ = v33
	var v34 int64
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int64
	_ = v37
	var v40 int64
	_ = v40
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v77 int32
	_ = v77
	var v81 int32
	_ = v81
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	v8 = m.G0
	v10 = v8 - int32(16)
	m.G0 = v10
	v13 = *(*int32)(unsafe.Add(mBase, _consts[370]))
	if v13 == int32(0) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	m.G0 = v10 + int32(16)
	return
L2:
	;
	v17 = *(*int32)(unsafe.Add(mBase, _consts[217]))
	v18 = *(*int32)(unsafe.Add(mBase, uint32(v17)+20))
	if v18 == int32(0) {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v22 = *(*int32)(unsafe.Add(mBase, _consts[209]))
	if v22 == int32(0) {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v34 = *(*int64)(unsafe.Add(mBase, uint32(v13)+16))
	v35 = F_sdsempty(m)
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
		goto L7
	} else {
		goto L9
	}
L5:
	;
	v26 = *(*int32)(unsafe.Add(mBase, _consts[6]))
	if int32(2) < v26 {
		goto L1
	} else {
		goto L6
	}
L6:
	;
	F__serverLog(m, int32(2), int32(_a1193), int32(0))
	mBase = m.M
	v33 = m.ExcPending
	if v33 != 0 {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	return
L8:
	;
	goto L1
L9:
	;
	v37 = int64(256)
	if v34 < v37 {
		goto L12
	} else {
		goto L13
	}
L10:
	;
	v81 = *(*int32)(unsafe.Add(mBase, _consts[6]))
	if int32(2) < v81 {
		goto L29
	} else {
		goto L30
	}
L11:
	;
	v43 = *(*int32)(unsafe.Add(mBase, _consts[217]))
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v43)+4))
	if v44 != 0 {
		goto L16
	} else {
		goto L17
	}
L12:
	;
	v40 = v34
	goto L14
L13:
	;
	v40 = v37
	goto L14
L14:
	;
	v41 = base.I32_wrap_i64(v40)
	if v41 != 0 {
		goto L11
	} else {
		goto L15
	}
L15:
	;
	v77 = v35
	goto L10
L16:
	;
	v46 = v41
	v48 = v35
	v50 = v44
	goto L18
L17:
	;
	v77 = v35
	goto L10
L18:
	;
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v50)+8))
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v52)+28))
	v54 = F_sdsempty(m)
	mBase = m.M
	v55 = m.ExcPending
	if v55 != 0 {
		goto L7
	} else {
		goto L20
	}
L19:
	;
	v77 = v65
	goto L10
L20:
	;
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v52)+28))
	if base.Ui32(v53) < base.Ui32(v46) {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	v59 = v53
	goto L23
L22:
	;
	v59 = v46
	goto L23
L23:
	;
	v63 = F_sdscatrepr(m, v54, v52+v56-v59+int32(32), v59)
	mBase = m.M
	v64 = m.ExcPending
	if v64 != 0 {
		goto L7
	} else {
		goto L24
	}
L24:
	;
	v65 = F_sdscatsds(m, v63, v48)
	mBase = m.M
	v66 = m.ExcPending
	if v66 != 0 {
		goto L7
	} else {
		goto L25
	}
L25:
	;
	F_sdsfree(m, v48)
	mBase = m.M
	v68 = m.ExcPending
	if v68 != 0 {
		goto L7
	} else {
		goto L26
	}
L26:
	;
	v69 = v46 - v59
	if v69 == int32(0) {
		v77 = v65
		goto L10
	} else {
		goto L27
	}
L27:
	;
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v50)))
	if v72 != 0 {
		v46 = v69
		v48 = v65
		v50 = v72
		goto L18
	} else {
		goto L28
	}
L28:
	;
	goto L19
L29:
	;
	F_sdsfree(m, v77)
	mBase = m.M
	v90 = m.ExcPending
	if v90 != 0 {
		goto L7
	} else {
		goto L32
	}
L30:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10))) = v77
	F__serverLog(m, int32(2), int32(_a1194), v10)
	mBase = m.M
	v88 = m.ExcPending
	if v88 != 0 {
		goto L7
	} else {
		goto L31
	}
L31:
	;
	goto L29
L32:
	;
	goto L1
}
func F_shutdown(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v15 int32
	_ = v15
	v3 = int32(0)
	v7 = F___syscall_shutdown(m, l0, l1, v3, v3, v3, v3)
	mBase = m.M
	if base.Ui32(v7) < base.Ui32(int32(-4095)) {
		v15 = v7
	} else {
		v10 = F___errno_location(m)
		mBase = m.M
		*(*int32)(unsafe.Add(mBase, uint32(v10))) = int32(0) - v7
		v15 = int32(-1)
	}
	return v15
}
func F_sigShutdownHandler(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v10 int32
	_ = v10
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v29 int32
	_ = v29
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v62 int32
	_ = v62
	var v66 int32
	_ = v66
	var v69 int64
	_ = v69
	var v70 int32
	_ = v70
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v90 int32
	_ = v90
	var v95 int32
	_ = v95
	var v102 int32
	_ = v102
	var v104 int32
	_ = v104
	var v107 int32
	_ = v107
	var v109 int32
	_ = v109
	var v112 int32
	_ = v112
	var v114 int32
	_ = v114
	var v117 int32
	_ = v117
	var v127 int32
	_ = v127
	var v130 int32
	_ = v130
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v137 int32
	_ = v137
	var v139 int32
	_ = v139
	var v144 int32
	_ = v144
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v152 int32
	_ = v152
	var v159 int32
	_ = v159
	var v162 int32
	_ = v162
	var v170 int32
	_ = v170
	var v172 int32
	_ = v172
	var v174 int32
	_ = v174
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v185 int32
	_ = v185
	var v189 int32
	_ = v189
	var v192 int64
	_ = v192
	var v193 int32
	_ = v193
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v204 int32
	_ = v204
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v213 int32
	_ = v213
	var v218 int32
	_ = v218
	var v224 int32
	_ = v224
	v4 = *(*int32)(unsafe.Add(mBase, _consts[841]))
	if l0 != int32(2) {
		if l0 == int32(15) {
			v127 = int32(_a1602)
		} else {
			v127 = int32(_a1603)
		}
		if l0 == int32(2) {
			v130 = int32(_a1604)
		} else {
			v130 = v127
		}
		v132 = *(*int32)(unsafe.Add(mBase, _consts[131]))
		if v132 != 0 {
			v133 = int32(_a1605)
		} else {
			v133 = v130
		}
		v137 = m.G0
		v139 = v137 - int32(80)
		m.G0 = v139
		v144 = *(*int32)(unsafe.Add(mBase, _consts[6]))
		if int32(3) < v144 {
		} else {
			v147 = *(*int32)(unsafe.Add(mBase, _consts[780]))
			v148 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v147))))
			if v148&int32(255) != 0 {
				if v148&int32(255) != 0 {
					*(*int32)(unsafe.Add(mBase, uint32(v139))) = int32(420)
					v159 = F_open(m, v147, int32(1089), v139)
					mBase = m.M
					if v159 == int32(-1) {
					} else {
						v162 = v159
						v170 = v139 + int32(16)
						v172 = F_getpid(m)
						mBase = m.M
						v174 = F_ll2string(m, v170, int32(64), base.I64_extend_i32_s(v172))
						mBase = m.M
						v179 = F_strlen(m, v170)
						mBase = m.M
						v180 = F_write(m, v162, v170, v179)
						mBase = m.M
						if v180 == int32(-1) {
						} else {
							v185 = F_write(m, v162, int32(_a1566), int32(17))
							mBase = m.M
							if v185 == int32(-1) {
							} else {
								v189 = v139 + int32(16)
								v192 = F___time(m, int32(0))
								mBase = m.M
								v193 = F_ll2string(m, v189, int32(64), v192)
								mBase = m.M
								v198 = F_strlen(m, v189)
								mBase = m.M
								v199 = F_write(m, v162, v189, v198)
								mBase = m.M
								if v199 == int32(-1) {
								} else {
									v204 = F_write(m, v162, int32(_a1567), int32(2))
									mBase = m.M
									if v204 == int32(-1) {
									} else {
										v207 = F_strlen(m, v133)
										mBase = m.M
										v208 = F_write(m, v162, v133, v207)
										mBase = m.M
										if v208 == int32(-1) {
										} else {
											v213 = F_write(m, v162, int32(_a247), int32(1))
											mBase = m.M
										}
									}
								}
							}
						}
						if v148&int32(255) == int32(0) {
						} else {
							v218 = F_close(m, v162)
							mBase = m.M
						}
					}
				} else {
					v162 = int32(1)
					v170 = v139 + int32(16)
					v172 = F_getpid(m)
					mBase = m.M
					v174 = F_ll2string(m, v170, int32(64), base.I64_extend_i32_s(v172))
					mBase = m.M
					v179 = F_strlen(m, v170)
					mBase = m.M
					v180 = F_write(m, v162, v170, v179)
					mBase = m.M
					if v180 == int32(-1) {
					} else {
						v185 = F_write(m, v162, int32(_a1566), int32(17))
						mBase = m.M
						if v185 == int32(-1) {
						} else {
							v189 = v139 + int32(16)
							v192 = F___time(m, int32(0))
							mBase = m.M
							v193 = F_ll2string(m, v189, int32(64), v192)
							mBase = m.M
							v198 = F_strlen(m, v189)
							mBase = m.M
							v199 = F_write(m, v162, v189, v198)
							mBase = m.M
							if v199 == int32(-1) {
							} else {
								v204 = F_write(m, v162, int32(_a1567), int32(2))
								mBase = m.M
								if v204 == int32(-1) {
								} else {
									v207 = F_strlen(m, v133)
									mBase = m.M
									v208 = F_write(m, v162, v133, v207)
									mBase = m.M
									if v208 == int32(-1) {
									} else {
										v213 = F_write(m, v162, int32(_a247), int32(1))
										mBase = m.M
									}
								}
							}
						}
					}
					if v148&int32(255) == int32(0) {
					} else {
						v218 = F_close(m, v162)
						mBase = m.M
					}
				}
			} else {
				v152 = *(*int32)(unsafe.Add(mBase, _consts[355]))
				if v152 != 0 {
				} else {
					if v148&int32(255) != 0 {
						*(*int32)(unsafe.Add(mBase, uint32(v139))) = int32(420)
						v159 = F_open(m, v147, int32(1089), v139)
						mBase = m.M
						if v159 == int32(-1) {
						} else {
							v162 = v159
							v170 = v139 + int32(16)
							v172 = F_getpid(m)
							mBase = m.M
							v174 = F_ll2string(m, v170, int32(64), base.I64_extend_i32_s(v172))
							mBase = m.M
							v179 = F_strlen(m, v170)
							mBase = m.M
							v180 = F_write(m, v162, v170, v179)
							mBase = m.M
							if v180 == int32(-1) {
							} else {
								v185 = F_write(m, v162, int32(_a1566), int32(17))
								mBase = m.M
								if v185 == int32(-1) {
								} else {
									v189 = v139 + int32(16)
									v192 = F___time(m, int32(0))
									mBase = m.M
									v193 = F_ll2string(m, v189, int32(64), v192)
									mBase = m.M
									v198 = F_strlen(m, v189)
									mBase = m.M
									v199 = F_write(m, v162, v189, v198)
									mBase = m.M
									if v199 == int32(-1) {
									} else {
										v204 = F_write(m, v162, int32(_a1567), int32(2))
										mBase = m.M
										if v204 == int32(-1) {
										} else {
											v207 = F_strlen(m, v133)
											mBase = m.M
											v208 = F_write(m, v162, v133, v207)
											mBase = m.M
											if v208 == int32(-1) {
											} else {
												v213 = F_write(m, v162, int32(_a247), int32(1))
												mBase = m.M
											}
										}
									}
								}
							}
							if v148&int32(255) == int32(0) {
							} else {
								v218 = F_close(m, v162)
								mBase = m.M
							}
						}
					} else {
						v162 = int32(1)
						v170 = v139 + int32(16)
						v172 = F_getpid(m)
						mBase = m.M
						v174 = F_ll2string(m, v170, int32(64), base.I64_extend_i32_s(v172))
						mBase = m.M
						v179 = F_strlen(m, v170)
						mBase = m.M
						v180 = F_write(m, v162, v170, v179)
						mBase = m.M
						if v180 == int32(-1) {
						} else {
							v185 = F_write(m, v162, int32(_a1566), int32(17))
							mBase = m.M
							if v185 == int32(-1) {
							} else {
								v189 = v139 + int32(16)
								v192 = F___time(m, int32(0))
								mBase = m.M
								v193 = F_ll2string(m, v189, int32(64), v192)
								mBase = m.M
								v198 = F_strlen(m, v189)
								mBase = m.M
								v199 = F_write(m, v162, v189, v198)
								mBase = m.M
								if v199 == int32(-1) {
								} else {
									v204 = F_write(m, v162, int32(_a1567), int32(2))
									mBase = m.M
									if v204 == int32(-1) {
									} else {
										v207 = F_strlen(m, v133)
										mBase = m.M
										v208 = F_write(m, v162, v133, v207)
										mBase = m.M
										if v208 == int32(-1) {
										} else {
											v213 = F_write(m, v162, int32(_a247), int32(1))
											mBase = m.M
										}
									}
								}
							}
						}
						if v148&int32(255) == int32(0) {
						} else {
							v218 = F_close(m, v162)
							mBase = m.M
						}
					}
				}
			}
		}
		m.G0 = v139 + int32(80)
		v224 = int32(0)
		*(*int32)(unsafe.Add(mBase, _consts[844])) = l0
		*(*int32)(unsafe.Add(mBase, _consts[841])) = int32(1)
		return
	} else {
		if v4 == int32(0) {
			if l0 == int32(15) {
				v127 = int32(_a1602)
			} else {
				v127 = int32(_a1603)
			}
			if l0 == int32(2) {
				v130 = int32(_a1604)
			} else {
				v130 = v127
			}
			v132 = *(*int32)(unsafe.Add(mBase, _consts[131]))
			if v132 != 0 {
				v133 = int32(_a1605)
			} else {
				v133 = v130
			}
			v137 = m.G0
			v139 = v137 - int32(80)
			m.G0 = v139
			v144 = *(*int32)(unsafe.Add(mBase, _consts[6]))
			if int32(3) < v144 {
			} else {
				v147 = *(*int32)(unsafe.Add(mBase, _consts[780]))
				v148 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v147))))
				if v148&int32(255) != 0 {
					if v148&int32(255) != 0 {
						*(*int32)(unsafe.Add(mBase, uint32(v139))) = int32(420)
						v159 = F_open(m, v147, int32(1089), v139)
						mBase = m.M
						if v159 == int32(-1) {
						} else {
							v162 = v159
							v170 = v139 + int32(16)
							v172 = F_getpid(m)
							mBase = m.M
							v174 = F_ll2string(m, v170, int32(64), base.I64_extend_i32_s(v172))
							mBase = m.M
							v179 = F_strlen(m, v170)
							mBase = m.M
							v180 = F_write(m, v162, v170, v179)
							mBase = m.M
							if v180 == int32(-1) {
							} else {
								v185 = F_write(m, v162, int32(_a1566), int32(17))
								mBase = m.M
								if v185 == int32(-1) {
								} else {
									v189 = v139 + int32(16)
									v192 = F___time(m, int32(0))
									mBase = m.M
									v193 = F_ll2string(m, v189, int32(64), v192)
									mBase = m.M
									v198 = F_strlen(m, v189)
									mBase = m.M
									v199 = F_write(m, v162, v189, v198)
									mBase = m.M
									if v199 == int32(-1) {
									} else {
										v204 = F_write(m, v162, int32(_a1567), int32(2))
										mBase = m.M
										if v204 == int32(-1) {
										} else {
											v207 = F_strlen(m, v133)
											mBase = m.M
											v208 = F_write(m, v162, v133, v207)
											mBase = m.M
											if v208 == int32(-1) {
											} else {
												v213 = F_write(m, v162, int32(_a247), int32(1))
												mBase = m.M
											}
										}
									}
								}
							}
							if v148&int32(255) == int32(0) {
							} else {
								v218 = F_close(m, v162)
								mBase = m.M
							}
						}
					} else {
						v162 = int32(1)
						v170 = v139 + int32(16)
						v172 = F_getpid(m)
						mBase = m.M
						v174 = F_ll2string(m, v170, int32(64), base.I64_extend_i32_s(v172))
						mBase = m.M
						v179 = F_strlen(m, v170)
						mBase = m.M
						v180 = F_write(m, v162, v170, v179)
						mBase = m.M
						if v180 == int32(-1) {
						} else {
							v185 = F_write(m, v162, int32(_a1566), int32(17))
							mBase = m.M
							if v185 == int32(-1) {
							} else {
								v189 = v139 + int32(16)
								v192 = F___time(m, int32(0))
								mBase = m.M
								v193 = F_ll2string(m, v189, int32(64), v192)
								mBase = m.M
								v198 = F_strlen(m, v189)
								mBase = m.M
								v199 = F_write(m, v162, v189, v198)
								mBase = m.M
								if v199 == int32(-1) {
								} else {
									v204 = F_write(m, v162, int32(_a1567), int32(2))
									mBase = m.M
									if v204 == int32(-1) {
									} else {
										v207 = F_strlen(m, v133)
										mBase = m.M
										v208 = F_write(m, v162, v133, v207)
										mBase = m.M
										if v208 == int32(-1) {
										} else {
											v213 = F_write(m, v162, int32(_a247), int32(1))
											mBase = m.M
										}
									}
								}
							}
						}
						if v148&int32(255) == int32(0) {
						} else {
							v218 = F_close(m, v162)
							mBase = m.M
						}
					}
				} else {
					v152 = *(*int32)(unsafe.Add(mBase, _consts[355]))
					if v152 != 0 {
					} else {
						if v148&int32(255) != 0 {
							*(*int32)(unsafe.Add(mBase, uint32(v139))) = int32(420)
							v159 = F_open(m, v147, int32(1089), v139)
							mBase = m.M
							if v159 == int32(-1) {
							} else {
								v162 = v159
								v170 = v139 + int32(16)
								v172 = F_getpid(m)
								mBase = m.M
								v174 = F_ll2string(m, v170, int32(64), base.I64_extend_i32_s(v172))
								mBase = m.M
								v179 = F_strlen(m, v170)
								mBase = m.M
								v180 = F_write(m, v162, v170, v179)
								mBase = m.M
								if v180 == int32(-1) {
								} else {
									v185 = F_write(m, v162, int32(_a1566), int32(17))
									mBase = m.M
									if v185 == int32(-1) {
									} else {
										v189 = v139 + int32(16)
										v192 = F___time(m, int32(0))
										mBase = m.M
										v193 = F_ll2string(m, v189, int32(64), v192)
										mBase = m.M
										v198 = F_strlen(m, v189)
										mBase = m.M
										v199 = F_write(m, v162, v189, v198)
										mBase = m.M
										if v199 == int32(-1) {
										} else {
											v204 = F_write(m, v162, int32(_a1567), int32(2))
											mBase = m.M
											if v204 == int32(-1) {
											} else {
												v207 = F_strlen(m, v133)
												mBase = m.M
												v208 = F_write(m, v162, v133, v207)
												mBase = m.M
												if v208 == int32(-1) {
												} else {
													v213 = F_write(m, v162, int32(_a247), int32(1))
													mBase = m.M
												}
											}
										}
									}
								}
								if v148&int32(255) == int32(0) {
								} else {
									v218 = F_close(m, v162)
									mBase = m.M
								}
							}
						} else {
							v162 = int32(1)
							v170 = v139 + int32(16)
							v172 = F_getpid(m)
							mBase = m.M
							v174 = F_ll2string(m, v170, int32(64), base.I64_extend_i32_s(v172))
							mBase = m.M
							v179 = F_strlen(m, v170)
							mBase = m.M
							v180 = F_write(m, v162, v170, v179)
							mBase = m.M
							if v180 == int32(-1) {
							} else {
								v185 = F_write(m, v162, int32(_a1566), int32(17))
								mBase = m.M
								if v185 == int32(-1) {
								} else {
									v189 = v139 + int32(16)
									v192 = F___time(m, int32(0))
									mBase = m.M
									v193 = F_ll2string(m, v189, int32(64), v192)
									mBase = m.M
									v198 = F_strlen(m, v189)
									mBase = m.M
									v199 = F_write(m, v162, v189, v198)
									mBase = m.M
									if v199 == int32(-1) {
									} else {
										v204 = F_write(m, v162, int32(_a1567), int32(2))
										mBase = m.M
										if v204 == int32(-1) {
										} else {
											v207 = F_strlen(m, v133)
											mBase = m.M
											v208 = F_write(m, v162, v133, v207)
											mBase = m.M
											if v208 == int32(-1) {
											} else {
												v213 = F_write(m, v162, int32(_a247), int32(1))
												mBase = m.M
											}
										}
									}
								}
							}
							if v148&int32(255) == int32(0) {
							} else {
								v218 = F_close(m, v162)
								mBase = m.M
							}
						}
					}
				}
			}
			m.G0 = v139 + int32(80)
			v224 = int32(0)
			*(*int32)(unsafe.Add(mBase, _consts[844])) = l0
			*(*int32)(unsafe.Add(mBase, _consts[841])) = int32(1)
			return
		} else {
			v10 = int32(_a1606)
			v14 = m.G0
			v16 = v14 - int32(80)
			m.G0 = v16
			v21 = *(*int32)(unsafe.Add(mBase, _consts[6]))
			if int32(3) < v21 {
			} else {
				v24 = *(*int32)(unsafe.Add(mBase, _consts[780]))
				v25 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24))))
				if v25&int32(255) != 0 {
					if v25&int32(255) != 0 {
						*(*int32)(unsafe.Add(mBase, uint32(v16))) = int32(420)
						v36 = F_open(m, v24, int32(1089), v16)
						mBase = m.M
						if v36 == int32(-1) {
						} else {
							v39 = v36
							v47 = v16 + int32(16)
							v49 = F_getpid(m)
							mBase = m.M
							v51 = F_ll2string(m, v47, int32(64), base.I64_extend_i32_s(v49))
							mBase = m.M
							v56 = F_strlen(m, v47)
							mBase = m.M
							v57 = F_write(m, v39, v47, v56)
							mBase = m.M
							if v57 == int32(-1) {
							} else {
								v62 = F_write(m, v39, int32(_a1566), int32(17))
								mBase = m.M
								if v62 == int32(-1) {
								} else {
									v66 = v16 + int32(16)
									v69 = F___time(m, int32(0))
									mBase = m.M
									v70 = F_ll2string(m, v66, int32(64), v69)
									mBase = m.M
									v75 = F_strlen(m, v66)
									mBase = m.M
									v76 = F_write(m, v39, v66, v75)
									mBase = m.M
									if v76 == int32(-1) {
									} else {
										v81 = F_write(m, v39, int32(_a1567), int32(2))
										mBase = m.M
										if v81 == int32(-1) {
										} else {
											v84 = F_strlen(m, v10)
											mBase = m.M
											v85 = F_write(m, v39, v10, v84)
											mBase = m.M
											if v85 == int32(-1) {
											} else {
												v90 = F_write(m, v39, int32(_a247), int32(1))
												mBase = m.M
											}
										}
									}
								}
							}
							if v25&int32(255) == int32(0) {
							} else {
								v95 = F_close(m, v39)
								mBase = m.M
							}
						}
					} else {
						v39 = int32(1)
						v47 = v16 + int32(16)
						v49 = F_getpid(m)
						mBase = m.M
						v51 = F_ll2string(m, v47, int32(64), base.I64_extend_i32_s(v49))
						mBase = m.M
						v56 = F_strlen(m, v47)
						mBase = m.M
						v57 = F_write(m, v39, v47, v56)
						mBase = m.M
						if v57 == int32(-1) {
						} else {
							v62 = F_write(m, v39, int32(_a1566), int32(17))
							mBase = m.M
							if v62 == int32(-1) {
							} else {
								v66 = v16 + int32(16)
								v69 = F___time(m, int32(0))
								mBase = m.M
								v70 = F_ll2string(m, v66, int32(64), v69)
								mBase = m.M
								v75 = F_strlen(m, v66)
								mBase = m.M
								v76 = F_write(m, v39, v66, v75)
								mBase = m.M
								if v76 == int32(-1) {
								} else {
									v81 = F_write(m, v39, int32(_a1567), int32(2))
									mBase = m.M
									if v81 == int32(-1) {
									} else {
										v84 = F_strlen(m, v10)
										mBase = m.M
										v85 = F_write(m, v39, v10, v84)
										mBase = m.M
										if v85 == int32(-1) {
										} else {
											v90 = F_write(m, v39, int32(_a247), int32(1))
											mBase = m.M
										}
									}
								}
							}
						}
						if v25&int32(255) == int32(0) {
						} else {
							v95 = F_close(m, v39)
							mBase = m.M
						}
					}
				} else {
					v29 = *(*int32)(unsafe.Add(mBase, _consts[355]))
					if v29 != 0 {
					} else {
						if v25&int32(255) != 0 {
							*(*int32)(unsafe.Add(mBase, uint32(v16))) = int32(420)
							v36 = F_open(m, v24, int32(1089), v16)
							mBase = m.M
							if v36 == int32(-1) {
							} else {
								v39 = v36
								v47 = v16 + int32(16)
								v49 = F_getpid(m)
								mBase = m.M
								v51 = F_ll2string(m, v47, int32(64), base.I64_extend_i32_s(v49))
								mBase = m.M
								v56 = F_strlen(m, v47)
								mBase = m.M
								v57 = F_write(m, v39, v47, v56)
								mBase = m.M
								if v57 == int32(-1) {
								} else {
									v62 = F_write(m, v39, int32(_a1566), int32(17))
									mBase = m.M
									if v62 == int32(-1) {
									} else {
										v66 = v16 + int32(16)
										v69 = F___time(m, int32(0))
										mBase = m.M
										v70 = F_ll2string(m, v66, int32(64), v69)
										mBase = m.M
										v75 = F_strlen(m, v66)
										mBase = m.M
										v76 = F_write(m, v39, v66, v75)
										mBase = m.M
										if v76 == int32(-1) {
										} else {
											v81 = F_write(m, v39, int32(_a1567), int32(2))
											mBase = m.M
											if v81 == int32(-1) {
											} else {
												v84 = F_strlen(m, v10)
												mBase = m.M
												v85 = F_write(m, v39, v10, v84)
												mBase = m.M
												if v85 == int32(-1) {
												} else {
													v90 = F_write(m, v39, int32(_a247), int32(1))
													mBase = m.M
												}
											}
										}
									}
								}
								if v25&int32(255) == int32(0) {
								} else {
									v95 = F_close(m, v39)
									mBase = m.M
								}
							}
						} else {
							v39 = int32(1)
							v47 = v16 + int32(16)
							v49 = F_getpid(m)
							mBase = m.M
							v51 = F_ll2string(m, v47, int32(64), base.I64_extend_i32_s(v49))
							mBase = m.M
							v56 = F_strlen(m, v47)
							mBase = m.M
							v57 = F_write(m, v39, v47, v56)
							mBase = m.M
							if v57 == int32(-1) {
							} else {
								v62 = F_write(m, v39, int32(_a1566), int32(17))
								mBase = m.M
								if v62 == int32(-1) {
								} else {
									v66 = v16 + int32(16)
									v69 = F___time(m, int32(0))
									mBase = m.M
									v70 = F_ll2string(m, v66, int32(64), v69)
									mBase = m.M
									v75 = F_strlen(m, v66)
									mBase = m.M
									v76 = F_write(m, v39, v66, v75)
									mBase = m.M
									if v76 == int32(-1) {
									} else {
										v81 = F_write(m, v39, int32(_a1567), int32(2))
										mBase = m.M
										if v81 == int32(-1) {
										} else {
											v84 = F_strlen(m, v10)
											mBase = m.M
											v85 = F_write(m, v39, v10, v84)
											mBase = m.M
											if v85 == int32(-1) {
											} else {
												v90 = F_write(m, v39, int32(_a247), int32(1))
												mBase = m.M
											}
										}
									}
								}
							}
							if v25&int32(255) == int32(0) {
							} else {
								v95 = F_close(m, v39)
								mBase = m.M
							}
						}
					}
				}
			}
			m.G0 = v16 + int32(80)
			v102 = *(*int32)(unsafe.Add(mBase, _consts[400]))
			if v102 != 0 {
				m.Env.Exit(m, int32(1))
				mBase = m.M
				base.Wasm_trap_unreachable()
				for {
				}
			} else {
				v104 = *(*int32)(unsafe.Add(mBase, _consts[635]))
				F_rdbRemoveTempFile(m, v104, int32(1))
				mBase = m.M
				v107 = m.ExcPending
				if v107 != 0 {
					return
				} else {
					v109 = *(*int32)(unsafe.Add(mBase, _consts[45]))
					F_rdbRemoveTempFile(m, v109, int32(1))
					mBase = m.M
					v112 = m.ExcPending
					if v112 != 0 {
						return
					} else {
						v114 = *(*int32)(unsafe.Add(mBase, _consts[45]))
						F_aofRemoveTempFile(m, v114, int32(1))
						mBase = m.M
						v117 = m.ExcPending
						if v117 != 0 {
							return
						} else {
							m.Env.Exit(m, int32(1))
							mBase = m.M
							base.Wasm_trap_unreachable()
							for {
							}
						}
					}
				}
			}
		}
	}
}
func F_sigaddset(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	v5 = l1 + int32(-1)
	if base.Ui32(int32(63)) < base.Ui32(v5) {
		*(*int32)(unsafe.Add(mBase, _consts[5])) = int32(28)
		return int32(-1)
	} else {
		if base.Ui32(int32(2)) < base.Ui32(l1+int32(-32)) {
			v21 = l0 + int32(base.Ui32(v5)>>(uint(int32(3))%32))&int32(536870908)
			v22 = *(*int32)(unsafe.Add(mBase, uint32(v21)))
			*(*int32)(unsafe.Add(mBase, uint32(v21))) = v22 | int32(1)<<(uint(v5)%32)
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, _consts[5])) = int32(28)
			return int32(-1)
		}
	}
}
func F_sigismember(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v14 int32
	_ = v14
	var v18 int32
	_ = v18
	v6 = l1 + int32(-1)
	if base.Ui32(int32(63)) < base.Ui32(v6) {
		v18 = int32(0)
	} else {
		v14 = *(*int32)(unsafe.Add(mBase, uint32(l0+int32(base.Ui32(v6)>>(uint(int32(3))%32))&int32(536870908))))
		v18 = int32(base.Ui32(v14)>>(uint(v6)%32)) & int32(1)
	}
	return v18
}
func F_sin(m *base.Module, l0 float64) float64 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v14 int32
	_ = v14
	var v24 float64
	_ = v24
	var v60 int32
	_ = v60
	var v61 float64
	_ = v61
	var v62 float64
	_ = v62
	var v69 float64
	_ = v69
	var v85 float64
	_ = v85
	var v105 float64
	_ = v105
	var v106 float64
	_ = v106
	var v108 float64
	_ = v108
	var v109 float64
	_ = v109
	var v121 float64
	_ = v121
	var v141 float64
	_ = v141
	var v157 float64
	_ = v157
	var v178 float64
	_ = v178
	var v179 float64
	_ = v179
	var v181 float64
	_ = v181
	var v182 float64
	_ = v182
	var v194 float64
	_ = v194
	var v211 float64
	_ = v211
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v14 = base.I32_wrap_i64(int64(base.Ui64(base.I64_reinterpret_f64(l0))>>(uint(int64(32))%64))) & int32(2147483647)
	if base.Ui32(int32(1072243195)) < base.Ui32(v14) {
		if base.Ui32(v14) < base.Ui32(int32(2146435072)) {
			v60 = F___rem_pio2(m, l0, v7)
			mBase = m.M
			v61 = *(*float64)(unsafe.Add(mBase, uint32(v7)+8))
			v62 = *(*float64)(unsafe.Add(mBase, uint32(v7)))
			switch v60 & int32(3) {
			default:
				v69 = base.F64_mul(v62, v62)
				v85 = base.F64_mul(v62, v69)
				v211 = base.F64_sub(v62, base.F64_add(base.F64_sub(base.F64_mul(v69, base.F64_sub(base.F64_mul(v61, float64(0.5)), base.F64_mul(v85, base.F64_add(base.F64_mul(base.F64_mul(v69, base.F64_mul(v69, v69)), base.F64_add(base.F64_mul(v69, float64(1.58969099521155e-10)), float64(-2.5050760253406863e-08))), base.F64_add(base.F64_mul(v69, base.F64_add(base.F64_mul(v69, float64(2.7557313707070068e-06)), float64(-0.0001984126982985795))), float64(0.00833333333332249)))))), v61), base.F64_mul(v85, float64(0.16666666666666632))))
			case 1:
				v105 = float64(1)
				v106 = base.F64_mul(v62, v62)
				v108 = base.F64_mul(v106, float64(0.5))
				v109 = base.F64_sub(v105, v108)
				v121 = base.F64_mul(v106, v106)
				v211 = base.F64_add(v109, base.F64_add(base.F64_sub(base.F64_sub(v105, v109), v108), base.F64_sub(base.F64_mul(v106, base.F64_add(base.F64_mul(v106, base.F64_add(base.F64_mul(v106, base.F64_add(base.F64_mul(v106, float64(2.480158728947673e-05)), float64(-0.001388888888887411))), float64(0.0416666666666666))), base.F64_mul(base.F64_mul(v121, v121), base.F64_add(base.F64_mul(v106, base.F64_add(base.F64_mul(v106, float64(-1.1359647557788195e-11)), float64(2.087572321298175e-09))), float64(-2.7557314351390663e-07))))), base.F64_mul(v62, v61))))
			case 2:
				v141 = base.F64_mul(v62, v62)
				v157 = base.F64_mul(v62, v141)
				v211 = base.F64_neg(base.F64_sub(v62, base.F64_add(base.F64_sub(base.F64_mul(v141, base.F64_sub(base.F64_mul(v61, float64(0.5)), base.F64_mul(v157, base.F64_add(base.F64_mul(base.F64_mul(v141, base.F64_mul(v141, v141)), base.F64_add(base.F64_mul(v141, float64(1.58969099521155e-10)), float64(-2.5050760253406863e-08))), base.F64_add(base.F64_mul(v141, base.F64_add(base.F64_mul(v141, float64(2.7557313707070068e-06)), float64(-0.0001984126982985795))), float64(0.00833333333332249)))))), v61), base.F64_mul(v157, float64(0.16666666666666632)))))
			case 3:
				v178 = float64(1)
				v179 = base.F64_mul(v62, v62)
				v181 = base.F64_mul(v179, float64(0.5))
				v182 = base.F64_sub(v178, v181)
				v194 = base.F64_mul(v179, v179)
				v211 = base.F64_neg(base.F64_add(v182, base.F64_add(base.F64_sub(base.F64_sub(v178, v182), v181), base.F64_sub(base.F64_mul(v179, base.F64_add(base.F64_mul(v179, base.F64_add(base.F64_mul(v179, base.F64_add(base.F64_mul(v179, float64(2.480158728947673e-05)), float64(-0.001388888888887411))), float64(0.0416666666666666))), base.F64_mul(base.F64_mul(v194, v194), base.F64_add(base.F64_mul(v179, base.F64_add(base.F64_mul(v179, float64(-1.1359647557788195e-11)), float64(2.087572321298175e-09))), float64(-2.7557314351390663e-07))))), base.F64_mul(v62, v61)))))
			}
		} else {
			v211 = base.F64_sub(l0, l0)
		}
	} else {
		if base.Ui32(v14) < base.Ui32(int32(1045430272)) {
			v211 = l0
		} else {
			v24 = base.F64_mul(l0, l0)
			v211 = base.F64_add(base.F64_mul(base.F64_mul(l0, v24), base.F64_add(base.F64_mul(v24, base.F64_add(base.F64_mul(base.F64_mul(v24, base.F64_mul(v24, v24)), base.F64_add(base.F64_mul(v24, float64(1.58969099521155e-10)), float64(-2.5050760253406863e-08))), base.F64_add(base.F64_mul(v24, base.F64_add(base.F64_mul(v24, float64(2.7557313707070068e-06)), float64(-0.0001984126982985795))), float64(0.00833333333332249)))), float64(-0.16666666666666632))), l0)
		}
	}
	m.G0 = v7 + int32(16)
	return v211
}
func F_singlestep(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v54 int32
	_ = v54
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v95 int32
	_ = v95
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
	var v114 int32
	_ = v114
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v141 int32
	_ = v141
	var v143 int32
	_ = v143
	var v145 int32
	_ = v145
	var v147 int32
	_ = v147
	var v149 int32
	_ = v149
	var v151 int32
	_ = v151
	var v161 int32
	_ = v161
	var v163 int32
	_ = v163
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v193 int32
	_ = v193
	var v200 int32
	_ = v200
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v206 int32
	_ = v206
	var v208 int32
	_ = v208
	var v214 int32
	_ = v214
	var v216 int32
	_ = v216
	var v219 int32
	_ = v219
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v246 int32
	_ = v246
	var v248 int32
	_ = v248
	var v250 int32
	_ = v250
	var v252 int32
	_ = v252
	var v254 int32
	_ = v254
	var v256 int32
	_ = v256
	var v266 int32
	_ = v266
	var v269 int32
	_ = v269
	var v275 int32
	_ = v275
	var v278 int32
	_ = v278
	var v284 int32
	_ = v284
	var v287 int32
	_ = v287
	var v293 int32
	_ = v293
	var v296 int32
	_ = v296
	var v302 int32
	_ = v302
	var v305 int32
	_ = v305
	var v311 int32
	_ = v311
	var v314 int32
	_ = v314
	var v320 int32
	_ = v320
	var v323 int32
	_ = v323
	var v329 int32
	_ = v329
	var v332 int32
	_ = v332
	var v338 int32
	_ = v338
	var v341 int32
	_ = v341
	var v347 int32
	_ = v347
	var v360 int32
	_ = v360
	var v361 int32
	_ = v361
	var v362 int32
	_ = v362
	var v373 int32
	_ = v373
	var v374 int32
	_ = v374
	var v389 int32
	_ = v389
	var v390 int32
	_ = v390
	var v391 int32
	_ = v391
	var v409 int32
	_ = v409
	var v410 int32
	_ = v410
	var v411 int32
	_ = v411
	var v417 int32
	_ = v417
	var v418 int32
	_ = v418
	var v419 int32
	_ = v419
	var v422 int32
	_ = v422
	var v430 int32
	_ = v430
	var v433 int32
	_ = v433
	var v437 int32
	_ = v437
	var v438 int32
	_ = v438
	var v439 int32
	_ = v439
	var v440 int32
	_ = v440
	var v441 int32
	_ = v441
	var v444 int32
	_ = v444
	var v447 int32
	_ = v447
	var v449 int32
	_ = v449
	var v450 int32
	_ = v450
	var v455 int32
	_ = v455
	var v457 int32
	_ = v457
	var v459 int32
	_ = v459
	var v462 int32
	_ = v462
	var v463 int32
	_ = v463
	var v466 int32
	_ = v466
	var v475 int32
	_ = v475
	var v476 int32
	_ = v476
	var v480 int32
	_ = v480
	var v489 int32
	_ = v489
	var v490 int32
	_ = v490
	var v493 int32
	_ = v493
	var v496 int32
	_ = v496
	var v500 int32
	_ = v500
	var v502 int32
	_ = v502
	var v503 int32
	_ = v503
	var v506 int32
	_ = v506
	var v508 int32
	_ = v508
	var v514 int32
	_ = v514
	var v516 int32
	_ = v516
	var v519 int32
	_ = v519
	var v525 int32
	_ = v525
	var v526 int32
	_ = v526
	var v529 int32
	_ = v529
	var v530 int32
	_ = v530
	var v533 int32
	_ = v533
	var v534 int32
	_ = v534
	var v540 int32
	_ = v540
	var v541 int32
	_ = v541
	var v546 int32
	_ = v546
	var v548 int32
	_ = v548
	var v550 int32
	_ = v550
	var v552 int32
	_ = v552
	var v554 int32
	_ = v554
	var v556 int32
	_ = v556
	var v565 int32
	_ = v565
	var v577 int32
	_ = v577
	var v578 int32
	_ = v578
	var v587 int32
	_ = v587
	var v591 int32
	_ = v591
	var v592 int32
	_ = v592
	var v593 int32
	_ = v593
	var v594 int32
	_ = v594
	var v601 int32
	_ = v601
	var v605 int32
	_ = v605
	var v612 int32
	_ = v612
	var v618 int32
	_ = v618
	var v623 int32
	_ = v623
	var v627 int32
	_ = v627
	var v636 int32
	_ = v636
	var v638 int32
	_ = v638
	var v639 int32
	_ = v639
	var v641 int32
	_ = v641
	var v642 int32
	_ = v642
	var v645 int32
	_ = v645
	var v646 int32
	_ = v646
	var v650 int32
	_ = v650
	var v674 int32
	_ = v674
	var v675 int32
	_ = v675
	var v679 int32
	_ = v679
	var v689 int32
	_ = v689
	var v692 int32
	_ = v692
	var v693 int32
	_ = v693
	var v696 int32
	_ = v696
	var v699 int32
	_ = v699
	var v700 int32
	_ = v700
	var v704 int32
	_ = v704
	var v706 int32
	_ = v706
	var v709 int32
	_ = v709
	var v714 int32
	_ = v714
	var v715 int32
	_ = v715
	var v719 int32
	_ = v719
	var v735 int32
	_ = v735
	var v746 int32
	_ = v746
	var v757 int32
	_ = v757
	var v759 int32
	_ = v759
	var v764 int32
	_ = v764
	var v766 int32
	_ = v766
	var v768 int32
	_ = v768
	var v774 int32
	_ = v774
	var v778 int32
	_ = v778
	var v779 int32
	_ = v779
	var v784 int32
	_ = v784
	var v785 int32
	_ = v785
	var v786 int32
	_ = v786
	var v787 int32
	_ = v787
	var v789 int32
	_ = v789
	var v791 int32
	_ = v791
	var v793 int32
	_ = v793
	var v798 int32
	_ = v798
	var v799 int32
	_ = v799
	var v801 int32
	_ = v801
	var v802 int32
	_ = v802
	var v804 int32
	_ = v804
	var v805 int32
	_ = v805
	var v806 int32
	_ = v806
	var v808 int32
	_ = v808
	var v811 int32
	_ = v811
	var v816 int32
	_ = v816
	var v817 int32
	_ = v817
	var v820 int32
	_ = v820
	var v822 int32
	_ = v822
	var v823 int32
	_ = v823
	var v824 int32
	_ = v824
	var v829 int32
	_ = v829
	var v835 int32
	_ = v835
	var v837 int32
	_ = v837
	var v842 int32
	_ = v842
	var v846 int32
	_ = v846
	var v848 int32
	_ = v848
	var v856 int32
	_ = v856
	var v862 int32
	_ = v862
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v13 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+21)))
	switch v13 {
	case 0:
		goto L6
	case 1:
		goto L5
	case 2:
		goto L4
	case 3:
		goto L3
	case 4:
		goto L2
	default:
		v862 = int32(0)
		goto L1
	}
L1:
	;
	return v862
L2:
	;
	v842 = *(*int32)(unsafe.Add(mBase, uint32(v12)+48))
	if v842 == int32(0) {
		goto L216
	} else {
		goto L217
	}
L3:
	;
	v798 = *(*int32)(unsafe.Add(mBase, uint32(v12)+68))
	v799 = *(*int32)(unsafe.Add(mBase, uint32(v12)+32))
	v801 = F_sweeplist(m, l0, v799, int32(40))
	mBase = m.M
	v802 = m.ExcPending
	if v802 != 0 {
		goto L18
	} else {
		goto L206
	}
L4:
	;
	v774 = *(*int32)(unsafe.Add(mBase, uint32(v12)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v12)+24)) = v774 + int32(1)
	v778 = *(*int32)(unsafe.Add(mBase, uint32(v12)+68))
	v779 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
	v784 = F_sweeplist(m, l0, v779+v774<<(uint(int32(2))%32), int32(-3))
	mBase = m.M
	v785 = m.ExcPending
	if v785 != 0 {
		goto L18
	} else {
		goto L203
	}
L5:
	;
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v12)+36))
	if v58 == int32(0) {
		goto L16
	} else {
		goto L17
	}
L6:
	;
	v14 = int32(0)
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v16)+44)) = v14
	*(*int64)(unsafe.Add(mBase, uint32(v16)+36)) = int64(0)
	v21 = *(*int32)(unsafe.Add(mBase, uint32(v16)+112))
	v22 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21)+5)))
	if v22&int32(3) == v14 {
		v29 = v21
		goto L8
	} else {
		goto L9
	}
L7:
	;
	return int32(0)
L8:
	;
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v29)+80))
	if v30 < int32(4) {
		goto L10
	} else {
		goto L11
	}
L9:
	;
	F_reallymarkobject(m, v16, v21)
	mBase = m.M
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v16)+112))
	v29 = v28
	goto L8
L10:
	;
	v41 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v41)+104))
	if v42 < int32(4) {
		goto L13
	} else {
		goto L14
	}
L11:
	;
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v29)+72))
	v34 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v33)+5)))
	if v34&int32(3) == int32(0) {
		goto L10
	} else {
		goto L12
	}
L12:
	;
	F_reallymarkobject(m, v16, v33)
	mBase = m.M
	goto L10
L13:
	;
	F_markmt(m, v16)
	mBase = m.M
	v54 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v16)+21)) = uint8(v54)
	goto L7
L14:
	;
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v41)+96))
	v46 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v45)+5)))
	if v46&int32(3) == int32(0) {
		goto L13
	} else {
		goto L15
	}
L15:
	;
	F_reallymarkobject(m, v16, v45)
	mBase = m.M
	goto L13
L16:
	;
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v12)+140))
	v68 = v12 + int32(120)
	if v66 == v68 {
		goto L20
	} else {
		goto L21
	}
L17:
	;
	v61 = F_propagatemark(m, v12)
	mBase = m.M
	v64 = m.ExcPending
	if v64 != 0 {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	return int32(0)
L19:
	;
	return v61
L20:
	;
	v189 = *(*int32)(unsafe.Add(mBase, uint32(v12)+44))
	v190 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v12)+44)) = v190
	*(*int32)(unsafe.Add(mBase, uint32(v12)+36)) = v189
	v193 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+5)))
	if v193&int32(3) == v190 {
		goto L54
	} else {
		goto L55
	}
L21:
	;
	v71 = v66
	goto L22
L22:
	;
	v80 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v71)+5)))
	if v80&int32(7) != 0 {
		goto L24
	} else {
		goto L25
	}
L23:
	;
	v163 = *(*int32)(unsafe.Add(mBase, uint32(v12)+36))
	if v163 == int32(0) {
		goto L20
	} else {
		goto L49
	}
L24:
	;
	v161 = *(*int32)(unsafe.Add(mBase, uint32(v71)+20))
	if v161 != v68 {
		v71 = v161
		goto L22
	} else {
		goto L48
	}
L25:
	;
	v83 = *(*int32)(unsafe.Add(mBase, uint32(v71)+8))
	v84 = *(*int32)(unsafe.Add(mBase, uint32(v83)+8))
	if v84 < int32(4) {
		goto L24
	} else {
		goto L26
	}
L26:
	;
	v87 = *(*int32)(unsafe.Add(mBase, uint32(v83)))
	v88 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v87)+5)))
	if v88&int32(3) == int32(0) {
		goto L24
	} else {
		goto L27
	}
L27:
	;
	v95 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v87)+5)))
	v97 = v87
	v98 = v95
	goto L36
L28:
	;
	goto L24
L29:
	;
	goto L28
L30:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+36)) = v97
	goto L29
L31:
	;
	v151 = *(*int32)(unsafe.Add(mBase, uint32(v12)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v97)+68)) = v151
	goto L30
L32:
	;
	v149 = *(*int32)(unsafe.Add(mBase, uint32(v12)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v97)+108)) = v149
	goto L30
L33:
	;
	v147 = *(*int32)(unsafe.Add(mBase, uint32(v12)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v97)+32)) = v147
	goto L30
L34:
	;
	v145 = *(*int32)(unsafe.Add(mBase, uint32(v12)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v97)+8)) = v145
	goto L30
L35:
	;
	v124 = *(*int32)(unsafe.Add(mBase, uint32(v97)+8))
	v125 = *(*int32)(unsafe.Add(mBase, uint32(v124)+8))
	if v125 < int32(4) {
		v136 = v124
		goto L44
	} else {
		goto L45
	}
L36:
	;
	v101 = v98 & int32(252)
	*(*uint8)(unsafe.Add(mBase, uint32(v97)+5)) = uint8(v101)
	v103 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v97)+4)))
	if v103 == int32(7) {
		goto L38
	} else {
		goto L39
	}
L38:
	;
	v109 = v101 | int32(4)
	*(*uint8)(unsafe.Add(mBase, uint32(v97)+5)) = uint8(v109)
	v111 = *(*int32)(unsafe.Add(mBase, uint32(v97)+8))
	if v111 == int32(0) {
		goto L40
	} else {
		goto L41
	}
L39:
	;
	switch v103 + int32(-5) {
	case 0:
		goto L33
	case 1:
		goto L34
	default:
		goto L29
	case 3:
		goto L32
	case 4:
		goto L31
	case 5:
		goto L35
	}
L40:
	;
	v120 = *(*int32)(unsafe.Add(mBase, uint32(v97)+12))
	v121 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v120)+5)))
	if v121&int32(3) != 0 {
		v97 = v120
		v98 = v121
		goto L36
	} else {
		goto L43
	}
L41:
	;
	v114 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v111)+5)))
	if v114&int32(3) == int32(0) {
		goto L40
	} else {
		goto L42
	}
L42:
	;
	F_reallymarkobject(m, v12, v111)
	mBase = m.M
	goto L40
L43:
	;
	goto L29
L44:
	;
	if v136 != v97+int32(16) {
		goto L29
	} else {
		goto L47
	}
L45:
	;
	v128 = *(*int32)(unsafe.Add(mBase, uint32(v124)))
	v129 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v128)+5)))
	if v129&int32(3) == int32(0) {
		v136 = v124
		goto L44
	} else {
		goto L46
	}
L46:
	;
	F_reallymarkobject(m, v12, v128)
	mBase = m.M
	v135 = *(*int32)(unsafe.Add(mBase, uint32(v97)+8))
	v136 = v135
	goto L44
L47:
	;
	v141 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v97)+5)))
	v143 = v141 | int32(4)
	*(*uint8)(unsafe.Add(mBase, uint32(v97)+5)) = uint8(v143)
	goto L28
L48:
	;
	goto L23
L49:
	;
	goto L50
L50:
	;
	v176 = F_propagatemark(m, v12)
	mBase = m.M
	v177 = m.ExcPending
	if v177 != 0 {
		goto L18
	} else {
		goto L52
	}
L51:
	;
	goto L20
L52:
	;
	v178 = *(*int32)(unsafe.Add(mBase, uint32(v12)+36))
	if v178 != 0 {
		goto L50
	} else {
		goto L53
	}
L53:
	;
	goto L51
L54:
	;
	v266 = *(*int32)(unsafe.Add(mBase, uint32(v12)+152))
	if v266 == int32(0) {
		goto L77
	} else {
		goto L78
	}
L55:
	;
	v200 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+5)))
	v202 = l0
	v203 = v200
	goto L64
L56:
	;
	goto L54
L57:
	;
	goto L56
L58:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+36)) = v202
	goto L57
L59:
	;
	v256 = *(*int32)(unsafe.Add(mBase, uint32(v12)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v202)+68)) = v256
	goto L58
L60:
	;
	v254 = *(*int32)(unsafe.Add(mBase, uint32(v12)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v202)+108)) = v254
	goto L58
L61:
	;
	v252 = *(*int32)(unsafe.Add(mBase, uint32(v12)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v202)+32)) = v252
	goto L58
L62:
	;
	v250 = *(*int32)(unsafe.Add(mBase, uint32(v12)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v202)+8)) = v250
	goto L58
L63:
	;
	v229 = *(*int32)(unsafe.Add(mBase, uint32(v202)+8))
	v230 = *(*int32)(unsafe.Add(mBase, uint32(v229)+8))
	if v230 < int32(4) {
		v241 = v229
		goto L72
	} else {
		goto L73
	}
L64:
	;
	v206 = v203 & int32(252)
	*(*uint8)(unsafe.Add(mBase, uint32(v202)+5)) = uint8(v206)
	v208 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v202)+4)))
	if v208 == int32(7) {
		goto L66
	} else {
		goto L67
	}
L66:
	;
	v214 = v206 | int32(4)
	*(*uint8)(unsafe.Add(mBase, uint32(v202)+5)) = uint8(v214)
	v216 = *(*int32)(unsafe.Add(mBase, uint32(v202)+8))
	if v216 == int32(0) {
		goto L68
	} else {
		goto L69
	}
L67:
	;
	switch v208 + int32(-5) {
	case 0:
		goto L61
	case 1:
		goto L62
	default:
		goto L57
	case 3:
		goto L60
	case 4:
		goto L59
	case 5:
		goto L63
	}
L68:
	;
	v225 = *(*int32)(unsafe.Add(mBase, uint32(v202)+12))
	v226 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v225)+5)))
	if v226&int32(3) != 0 {
		v202 = v225
		v203 = v226
		goto L64
	} else {
		goto L71
	}
L69:
	;
	v219 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v216)+5)))
	if v219&int32(3) == int32(0) {
		goto L68
	} else {
		goto L70
	}
L70:
	;
	F_reallymarkobject(m, v12, v216)
	mBase = m.M
	goto L68
L71:
	;
	goto L57
L72:
	;
	if v241 != v202+int32(16) {
		goto L57
	} else {
		goto L75
	}
L73:
	;
	v233 = *(*int32)(unsafe.Add(mBase, uint32(v229)))
	v234 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v233)+5)))
	if v234&int32(3) == int32(0) {
		v241 = v229
		goto L72
	} else {
		goto L74
	}
L74:
	;
	F_reallymarkobject(m, v12, v233)
	mBase = m.M
	v240 = *(*int32)(unsafe.Add(mBase, uint32(v202)+8))
	v241 = v240
	goto L72
L75:
	;
	v246 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v202)+5)))
	v248 = v246 | int32(4)
	*(*uint8)(unsafe.Add(mBase, uint32(v202)+5)) = uint8(v248)
	goto L56
L76:
	;
	v347 = *(*int32)(unsafe.Add(mBase, uint32(v12)+36))
	if v347 == int32(0) {
		goto L104
	} else {
		goto L105
	}
L77:
	;
	v275 = *(*int32)(unsafe.Add(mBase, uint32(v12)+156))
	if v275 == int32(0) {
		goto L80
	} else {
		goto L81
	}
L78:
	;
	v269 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v266)+5)))
	if v269&int32(3) == int32(0) {
		goto L77
	} else {
		goto L79
	}
L79:
	;
	F_reallymarkobject(m, v12, v266)
	mBase = m.M
	goto L77
L80:
	;
	v284 = *(*int32)(unsafe.Add(mBase, uint32(v12)+160))
	if v284 == int32(0) {
		goto L83
	} else {
		goto L84
	}
L81:
	;
	v278 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v275)+5)))
	if v278&int32(3) == int32(0) {
		goto L80
	} else {
		goto L82
	}
L82:
	;
	F_reallymarkobject(m, v12, v275)
	mBase = m.M
	goto L80
L83:
	;
	v293 = *(*int32)(unsafe.Add(mBase, uint32(v12)+164))
	if v293 == int32(0) {
		goto L86
	} else {
		goto L87
	}
L84:
	;
	v287 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v284)+5)))
	if v287&int32(3) == int32(0) {
		goto L83
	} else {
		goto L85
	}
L85:
	;
	F_reallymarkobject(m, v12, v284)
	mBase = m.M
	goto L83
L86:
	;
	v302 = *(*int32)(unsafe.Add(mBase, uint32(v12)+168))
	if v302 == int32(0) {
		goto L89
	} else {
		goto L90
	}
L87:
	;
	v296 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v293)+5)))
	if v296&int32(3) == int32(0) {
		goto L86
	} else {
		goto L88
	}
L88:
	;
	F_reallymarkobject(m, v12, v293)
	mBase = m.M
	goto L86
L89:
	;
	v311 = *(*int32)(unsafe.Add(mBase, uint32(v12)+172))
	if v311 == int32(0) {
		goto L92
	} else {
		goto L93
	}
L90:
	;
	v305 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v302)+5)))
	if v305&int32(3) == int32(0) {
		goto L89
	} else {
		goto L91
	}
L91:
	;
	F_reallymarkobject(m, v12, v302)
	mBase = m.M
	goto L89
L92:
	;
	v320 = *(*int32)(unsafe.Add(mBase, uint32(v12)+176))
	if v320 == int32(0) {
		goto L95
	} else {
		goto L96
	}
L93:
	;
	v314 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v311)+5)))
	if v314&int32(3) == int32(0) {
		goto L92
	} else {
		goto L94
	}
L94:
	;
	F_reallymarkobject(m, v12, v311)
	mBase = m.M
	goto L92
L95:
	;
	v329 = *(*int32)(unsafe.Add(mBase, uint32(v12)+180))
	if v329 == int32(0) {
		goto L98
	} else {
		goto L99
	}
L96:
	;
	v323 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v320)+5)))
	if v323&int32(3) == int32(0) {
		goto L95
	} else {
		goto L97
	}
L97:
	;
	F_reallymarkobject(m, v12, v320)
	mBase = m.M
	goto L95
L98:
	;
	v338 = *(*int32)(unsafe.Add(mBase, uint32(v12)+184))
	if v338 == int32(0) {
		goto L101
	} else {
		goto L102
	}
L99:
	;
	v332 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v329)+5)))
	if v332&int32(3) == int32(0) {
		goto L98
	} else {
		goto L100
	}
L100:
	;
	F_reallymarkobject(m, v12, v329)
	mBase = m.M
	goto L98
L101:
	;
	goto L76
L102:
	;
	v341 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v338)+5)))
	if v341&int32(3) == int32(0) {
		goto L101
	} else {
		goto L103
	}
L103:
	;
	F_reallymarkobject(m, v12, v338)
	mBase = m.M
	goto L101
L104:
	;
	v373 = *(*int32)(unsafe.Add(mBase, uint32(v12)+40))
	v374 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v12)+40)) = v374
	*(*int32)(unsafe.Add(mBase, uint32(v12)+36)) = v373
	if v373 == v374 {
		goto L110
	} else {
		goto L111
	}
L105:
	;
	goto L106
L106:
	;
	v360 = F_propagatemark(m, v12)
	mBase = m.M
	v361 = m.ExcPending
	if v361 != 0 {
		goto L18
	} else {
		goto L108
	}
L107:
	;
	goto L104
L108:
	;
	v362 = *(*int32)(unsafe.Add(mBase, uint32(v12)+36))
	if v362 != 0 {
		goto L106
	} else {
		goto L109
	}
L109:
	;
	goto L107
L110:
	;
	v409 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v410 = *(*int32)(unsafe.Add(mBase, uint32(v409)+112))
	v411 = *(*int32)(unsafe.Add(mBase, uint32(v410)))
	if v411 != 0 {
		goto L117
	} else {
		goto L118
	}
L111:
	;
	goto L112
L112:
	;
	v389 = F_propagatemark(m, v12)
	mBase = m.M
	v390 = m.ExcPending
	if v390 != 0 {
		goto L18
	} else {
		goto L114
	}
L113:
	;
	goto L110
L114:
	;
	v391 = *(*int32)(unsafe.Add(mBase, uint32(v12)+36))
	if v391 != 0 {
		goto L112
	} else {
		goto L115
	}
L115:
	;
	goto L113
L116:
	;
	v476 = *(*int32)(unsafe.Add(mBase, uint32(v12)+48))
	if v476 == int32(0) {
		goto L135
	} else {
		goto L136
	}
L117:
	;
	v417 = v410
	v418 = v411
	v419 = int32(0)
	goto L119
L118:
	;
	v475 = int32(0)
	goto L116
L119:
	;
	v422 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v418)+5)))
	if v422&int32(8) == int32(0) {
		goto L122
	} else {
		goto L123
	}
L120:
	;
	v475 = v463
	goto L116
L121:
	;
	v466 = *(*int32)(unsafe.Add(mBase, uint32(v462)))
	if v466 != 0 {
		v417 = v462
		v418 = v466
		v419 = v463
		goto L119
	} else {
		goto L134
	}
L122:
	;
	if v422&int32(3)|int32(0) != 0 {
		goto L124
	} else {
		goto L125
	}
L123:
	;
	v462 = v418
	v463 = v419
	goto L121
L124:
	;
	v430 = *(*int32)(unsafe.Add(mBase, uint32(v418)+8))
	if v430 == int32(0) {
		v441 = v422
		goto L127
	} else {
		goto L128
	}
L125:
	;
	v462 = v418
	v463 = v419
	goto L121
L126:
	;
	v447 = v440 | int32(8)
	*(*uint8)(unsafe.Add(mBase, uint32(v418)+5)) = uint8(v447)
	v449 = *(*int32)(unsafe.Add(mBase, uint32(v418)+16))
	v450 = *(*int32)(unsafe.Add(mBase, uint32(v418)))
	*(*int32)(unsafe.Add(mBase, uint32(v417))) = v450
	v455 = *(*int32)(unsafe.Add(mBase, uint32(v409)+48))
	if v455 != 0 {
		goto L132
	} else {
		goto L133
	}
L127:
	;
	v444 = v441 | int32(8)
	*(*uint8)(unsafe.Add(mBase, uint32(v418)+5)) = uint8(v444)
	v462 = v418
	v463 = v419
	goto L121
L128:
	;
	v433 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v430)+6)))
	if v433&int32(4) != 0 {
		v441 = v422
		goto L127
	} else {
		goto L129
	}
L129:
	;
	v437 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v438 = *(*int32)(unsafe.Add(mBase, uint32(v437)+196))
	v439 = F_luaT_gettm(m, v430, int32(2), v438)
	mBase = m.M
	v440 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v418)+5)))
	if v439 != 0 {
		goto L126
	} else {
		goto L130
	}
L130:
	;
	v441 = v440
	goto L127
L131:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v409)+48)) = v418
	v462 = v417
	v463 = v419 + v449 + int32(24)
	goto L121
L132:
	;
	v457 = *(*int32)(unsafe.Add(mBase, uint32(v455)))
	*(*int32)(unsafe.Add(mBase, uint32(v418))) = v457
	v459 = *(*int32)(unsafe.Add(mBase, uint32(v409)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v459))) = v418
	goto L131
L133:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v418))) = v418
	goto L131
L134:
	;
	goto L120
L135:
	;
	v577 = int32(0)
	v578 = *(*int32)(unsafe.Add(mBase, uint32(v12)+36))
	if v578 == v577 {
		v601 = v577
		goto L160
	} else {
		goto L161
	}
L136:
	;
	v480 = v476
	goto L137
L137:
	;
	v489 = *(*int32)(unsafe.Add(mBase, uint32(v480)))
	v490 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+20)))
	v493 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v489)+5)))
	v496 = v490&int32(3) | v493&int32(248)
	*(*uint8)(unsafe.Add(mBase, uint32(v489)+5)) = uint8(v496)
	v500 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v489)+5)))
	v502 = v489
	v503 = v500
	goto L147
L138:
	;
	goto L135
L139:
	;
	v565 = *(*int32)(unsafe.Add(mBase, uint32(v12)+48))
	if v489 != v565 {
		v480 = v489
		goto L137
	} else {
		goto L159
	}
L140:
	;
	goto L139
L141:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+36)) = v502
	goto L140
L142:
	;
	v556 = *(*int32)(unsafe.Add(mBase, uint32(v12)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v502)+68)) = v556
	goto L141
L143:
	;
	v554 = *(*int32)(unsafe.Add(mBase, uint32(v12)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v502)+108)) = v554
	goto L141
L144:
	;
	v552 = *(*int32)(unsafe.Add(mBase, uint32(v12)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v502)+32)) = v552
	goto L141
L145:
	;
	v550 = *(*int32)(unsafe.Add(mBase, uint32(v12)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v502)+8)) = v550
	goto L141
L146:
	;
	v529 = *(*int32)(unsafe.Add(mBase, uint32(v502)+8))
	v530 = *(*int32)(unsafe.Add(mBase, uint32(v529)+8))
	if v530 < int32(4) {
		v541 = v529
		goto L155
	} else {
		goto L156
	}
L147:
	;
	v506 = v503 & int32(252)
	*(*uint8)(unsafe.Add(mBase, uint32(v502)+5)) = uint8(v506)
	v508 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v502)+4)))
	if v508 == int32(7) {
		goto L149
	} else {
		goto L150
	}
L149:
	;
	v514 = v506 | int32(4)
	*(*uint8)(unsafe.Add(mBase, uint32(v502)+5)) = uint8(v514)
	v516 = *(*int32)(unsafe.Add(mBase, uint32(v502)+8))
	if v516 == int32(0) {
		goto L151
	} else {
		goto L152
	}
L150:
	;
	switch v508 + int32(-5) {
	case 0:
		goto L144
	case 1:
		goto L145
	default:
		goto L140
	case 3:
		goto L143
	case 4:
		goto L142
	case 5:
		goto L146
	}
L151:
	;
	v525 = *(*int32)(unsafe.Add(mBase, uint32(v502)+12))
	v526 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v525)+5)))
	if v526&int32(3) != 0 {
		v502 = v525
		v503 = v526
		goto L147
	} else {
		goto L154
	}
L152:
	;
	v519 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v516)+5)))
	if v519&int32(3) == int32(0) {
		goto L151
	} else {
		goto L153
	}
L153:
	;
	F_reallymarkobject(m, v12, v516)
	mBase = m.M
	goto L151
L154:
	;
	goto L140
L155:
	;
	if v541 != v502+int32(16) {
		goto L140
	} else {
		goto L158
	}
L156:
	;
	v533 = *(*int32)(unsafe.Add(mBase, uint32(v529)))
	v534 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v533)+5)))
	if v534&int32(3) == int32(0) {
		v541 = v529
		goto L155
	} else {
		goto L157
	}
L157:
	;
	F_reallymarkobject(m, v12, v533)
	mBase = m.M
	v540 = *(*int32)(unsafe.Add(mBase, uint32(v502)+8))
	v541 = v540
	goto L155
L158:
	;
	v546 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v502)+5)))
	v548 = v546 | int32(4)
	*(*uint8)(unsafe.Add(mBase, uint32(v502)+5)) = uint8(v548)
	goto L139
L159:
	;
	goto L138
L160:
	;
	v605 = *(*int32)(unsafe.Add(mBase, uint32(v12)+44))
	if v605 == int32(0) {
		goto L166
	} else {
		goto L167
	}
L161:
	;
	v587 = v577
	goto L162
L162:
	;
	v591 = F_propagatemark(m, v12)
	mBase = m.M
	v592 = m.ExcPending
	if v592 != 0 {
		goto L18
	} else {
		goto L164
	}
L163:
	;
	v601 = v593
	goto L160
L164:
	;
	v593 = v591 + v587
	v594 = *(*int32)(unsafe.Add(mBase, uint32(v12)+36))
	if v594 != 0 {
		v587 = v593
		goto L162
	} else {
		goto L165
	}
L165:
	;
	goto L163
L166:
	;
	v757 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v12)+24)) = v757
	v759 = int32(2)
	*(*uint8)(unsafe.Add(mBase, uint32(v12)+21)) = uint8(v759)
	*(*int32)(unsafe.Add(mBase, uint32(v12)+32)) = v12 + int32(28)
	v764 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+20)))
	v766 = v764 ^ int32(3)
	*(*uint8)(unsafe.Add(mBase, uint32(v12)+20)) = uint8(v766)
	v768 = *(*int32)(unsafe.Add(mBase, uint32(v12)+68))
	*(*int32)(unsafe.Add(mBase, uint32(v12)+72)) = v768 - (v601 + v475)
	return v757
L167:
	;
	v612 = v605
	goto L168
L168:
	;
	v618 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v612)+5)))
	if v618&int32(16) == int32(0) {
		goto L170
	} else {
		goto L171
	}
L169:
	;
	goto L166
L170:
	;
	v674 = int32(-1)
	v675 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v612)+12)))
	v679 = v674<<(uint(v675)%32) ^ v674
	goto L184
L171:
	;
	v623 = *(*int32)(unsafe.Add(mBase, uint32(v612)+36))
	if v623 == int32(0) {
		goto L170
	} else {
		goto L172
	}
L172:
	;
	v627 = v623
	goto L173
L173:
	;
	v636 = *(*int32)(unsafe.Add(mBase, uint32(v612)+20))
	v638 = v627 + int32(-1)
	v639 = int32(4)
	v641 = v636 + v638<<(uint(v639)%32)
	v642 = *(*int32)(unsafe.Add(mBase, uint32(v641)+8))
	if v642 < v639 {
		goto L175
	} else {
		goto L176
	}
L174:
	;
	goto L170
L175:
	;
	if v638 != 0 {
		v627 = v638
		goto L173
	} else {
		goto L183
	}
L176:
	;
	v645 = *(*int32)(unsafe.Add(mBase, uint32(v641)))
	v646 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v645)+5)))
	if v642 != int32(4) {
		goto L177
	} else {
		goto L178
	}
L177:
	;
	if v646&int32(3) != 0 {
		goto L179
	} else {
		goto L180
	}
L178:
	;
	v650 = v646 & int32(252)
	*(*uint8)(unsafe.Add(mBase, uint32(v645)+5)) = uint8(v650)
	goto L175
L179:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v641)+8)) = int32(0)
	goto L175
L180:
	;
	if v642 != int32(7) {
		goto L175
	} else {
		goto L181
	}
L181:
	;
	if v646&int32(8) == int32(0) {
		goto L175
	} else {
		goto L182
	}
L182:
	;
	goto L179
L183:
	;
	goto L174
L184:
	;
	v689 = *(*int32)(unsafe.Add(mBase, uint32(v612)+24))
	v692 = v689 + v679<<(uint(int32(5))%32)
	v693 = *(*int32)(unsafe.Add(mBase, uint32(v692)+8))
	if v693 == int32(0) {
		goto L186
	} else {
		goto L187
	}
L185:
	;
	v746 = *(*int32)(unsafe.Add(mBase, uint32(v612)+32))
	if v746 != 0 {
		v612 = v746
		goto L168
	} else {
		goto L202
	}
L186:
	;
	if v679 != 0 {
		v679 = v679 + int32(-1)
		goto L184
	} else {
		goto L201
	}
L187:
	;
	v696 = *(*int32)(unsafe.Add(mBase, uint32(v692)+24))
	if v696 < int32(4) {
		v709 = v693
		goto L189
	} else {
		goto L190
	}
L188:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v692)+8)) = int32(0)
	v735 = *(*int32)(unsafe.Add(mBase, uint32(v692)+24))
	if v735 < int32(4) {
		goto L186
	} else {
		goto L200
	}
L189:
	;
	if v709 < int32(4) {
		goto L186
	} else {
		goto L194
	}
L190:
	;
	v699 = *(*int32)(unsafe.Add(mBase, uint32(v692)+16))
	v700 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v699)+5)))
	if v696 != int32(4) {
		goto L191
	} else {
		goto L192
	}
L191:
	;
	if v700&int32(3) != 0 {
		goto L188
	} else {
		goto L193
	}
L192:
	;
	v704 = v700 & int32(252)
	*(*uint8)(unsafe.Add(mBase, uint32(v699)+5)) = uint8(v704)
	v706 = *(*int32)(unsafe.Add(mBase, uint32(v692)+8))
	v709 = v706
	goto L189
L193:
	;
	v709 = v693
	goto L189
L194:
	;
	v714 = *(*int32)(unsafe.Add(mBase, uint32(v692)))
	v715 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v714)+5)))
	if v709 != int32(4) {
		goto L195
	} else {
		goto L196
	}
L195:
	;
	if v715&int32(3) != 0 {
		goto L188
	} else {
		goto L197
	}
L196:
	;
	v719 = v715 & int32(252)
	*(*uint8)(unsafe.Add(mBase, uint32(v714)+5)) = uint8(v719)
	goto L186
L197:
	;
	if v709 != int32(7) {
		goto L186
	} else {
		goto L198
	}
L198:
	;
	if v715&int32(8) == int32(0) {
		goto L186
	} else {
		goto L199
	}
L199:
	;
	goto L188
L200:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v692)+24)) = int32(11)
	goto L186
L201:
	;
	goto L185
L202:
	;
	goto L169
L203:
	;
	v786 = *(*int32)(unsafe.Add(mBase, uint32(v12)+24))
	v787 = *(*int32)(unsafe.Add(mBase, uint32(v12)+8))
	if v786 < v787 {
		goto L204
	} else {
		goto L205
	}
L204:
	;
	v791 = *(*int32)(unsafe.Add(mBase, uint32(v12)+68))
	v793 = *(*int32)(unsafe.Add(mBase, uint32(v12)+72))
	*(*int32)(unsafe.Add(mBase, uint32(v12)+72)) = v791 - v778 + v793
	return int32(10)
L205:
	;
	v789 = int32(3)
	*(*uint8)(unsafe.Add(mBase, uint32(v12)+21)) = uint8(v789)
	goto L204
L206:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+32)) = v801
	v804 = *(*int32)(unsafe.Add(mBase, uint32(v801)))
	if v804 != 0 {
		goto L207
	} else {
		goto L208
	}
L207:
	;
	v835 = *(*int32)(unsafe.Add(mBase, uint32(v12)+68))
	v837 = *(*int32)(unsafe.Add(mBase, uint32(v12)+72))
	*(*int32)(unsafe.Add(mBase, uint32(v12)+72)) = v835 - v798 + v837
	return int32(400)
L208:
	;
	v805 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v806 = *(*int32)(unsafe.Add(mBase, uint32(v805)+8))
	v808 = base.I32_div_s(v806, int32(4))
	if v806 < int32(65) {
		goto L209
	} else {
		goto L210
	}
L209:
	;
	v817 = *(*int32)(unsafe.Add(mBase, uint32(v805)+60))
	if base.Ui32(v817) < base.Ui32(int32(65)) {
		goto L213
	} else {
		goto L214
	}
L210:
	;
	v811 = *(*int32)(unsafe.Add(mBase, uint32(v805)+4))
	if base.Ui32(v808) <= base.Ui32(v811) {
		goto L209
	} else {
		goto L211
	}
L211:
	;
	F_luaS_resize(m, l0, int32(base.Ui32(v806)>>(uint(int32(1))%32)))
	mBase = m.M
	v816 = m.ExcPending
	if v816 != 0 {
		goto L18
	} else {
		goto L212
	}
L212:
	;
	goto L209
L213:
	;
	v829 = int32(4)
	*(*uint8)(unsafe.Add(mBase, uint32(v12)+21)) = uint8(v829)
	goto L207
L214:
	;
	v820 = *(*int32)(unsafe.Add(mBase, uint32(v805)+52))
	v822 = int32(base.Ui32(v817) >> (uint(int32(1)) % 32))
	v823 = F_luaM_realloc_(m, l0, v820, v817, v822)
	mBase = m.M
	v824 = m.ExcPending
	if v824 != 0 {
		goto L18
	} else {
		goto L215
	}
L215:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v805)+60)) = v822
	*(*int32)(unsafe.Add(mBase, uint32(v805)+52)) = v823
	goto L213
L216:
	;
	v856 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v12)+76)) = v856
	*(*uint8)(unsafe.Add(mBase, uint32(v12)+21)) = uint8(v856)
	v862 = v856
	goto L1
L217:
	;
	F_GCTM(m, l0)
	mBase = m.M
	v846 = m.ExcPending
	if v846 != 0 {
		goto L18
	} else {
		goto L218
	}
L218:
	;
	v848 = *(*int32)(unsafe.Add(mBase, uint32(v12)+72))
	if base.Ui32(v848) < base.Ui32(int32(101)) {
		v862 = int32(100)
		goto L1
	} else {
		goto L219
	}
L219:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+72)) = v848 + int32(-100)
	return int32(100)
}
func F_singlevaraux(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v19 int32
	_ = v19
	var v26 int32
	_ = v26
	var v33 int32
	_ = v33
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v62 int32
	_ = v62
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
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
	var v93 int32
	_ = v93
	var v98 int32
	_ = v98
	var v106 int32
	_ = v106
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v121 int32
	_ = v121
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v129 int32
	_ = v129
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v140 int32
	_ = v140
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v188 int32
	_ = v188
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v202 int32
	_ = v202
	var v207 int32
	_ = v207
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v218 int32
	_ = v218
	var v221 int32
	_ = v221
	var v224 int32
	_ = v224
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v230 int32
	_ = v230
	var v233 int32
	_ = v233
	var v237 int32
	_ = v237
	var v240 int32
	_ = v240
	var v248 int32
	_ = v248
	var v254 int32
	_ = v254
	var v263 int32
	_ = v263
	v13 = m.G0
	v15 = v13 - int32(32)
	m.G0 = v15
	if l0 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v15 + int32(32)
	return v263
L2:
	;
	v26 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+50)))
	v33 = v26
	goto L5
L3:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2)+8)) = int32(255)
	v19 = int32(8)
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v19
	*(*int64)(unsafe.Add(mBase, uint32(l2)+16)) = int64(-1)
	v263 = v19
	goto L1
L4:
	;
	v82 = int32(0)
	v83 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v85 = F_singlevaraux(m, v83, l1, l2, v82)
	mBase = m.M
	v88 = m.ExcPending
	if v88 != 0 {
		goto L14
	} else {
		goto L15
	}
L5:
	;
	if v33 < int32(1) {
		goto L4
	} else {
		goto L7
	}
L6:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2)+8)) = v44
	v55 = int32(6)
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v55
	*(*int64)(unsafe.Add(mBase, uint32(l2)+16)) = int64(-1)
	if l3 != 0 {
		v263 = v55
		goto L1
	} else {
		goto L9
	}
L7:
	;
	v41 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v41)+24))
	v44 = v33 + int32(-1)
	v48 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0+int32(172)+v44<<(uint(int32(1))%32)))))
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v42+v48*int32(12))))
	if l1 != v52 {
		v33 = v44
		goto L5
	} else {
		goto L8
	}
L8:
	;
	goto L6
L9:
	;
	v62 = l0 + int32(20)
	goto L10
L10:
	;
	v74 = *(*int32)(unsafe.Add(mBase, uint32(v62)))
	if v74 == int32(0) {
		v263 = v55
		goto L1
	} else {
		goto L12
	}
L11:
	;
	v79 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v74)+9)) = uint8(v79)
	v263 = v55
	goto L1
L12:
	;
	v77 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v74)+8)))
	if base.Ui32(v44) < base.Ui32(v77) {
		v62 = v74
		goto L10
	} else {
		goto L13
	}
L13:
	;
	goto L11
L14:
	;
	return int32(0)
L15:
	;
	if v85 == int32(8) {
		v263 = int32(8)
		goto L1
	} else {
		goto L16
	}
L16:
	;
	v91 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v92 = *(*int32)(unsafe.Add(mBase, uint32(v91)+36))
	v93 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v91)+72)))
	if v93 == int32(0) {
		v164 = v92
		v165 = v82
		goto L18
	} else {
		goto L19
	}
L17:
	;
	v254 = int32(7)
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v254
	*(*int32)(unsafe.Add(mBase, uint32(l2)+8)) = v248
	v263 = v254
	goto L1
L18:
	;
	if v164 <= v165 {
		goto L34
	} else {
		goto L35
	}
L19:
	;
	v98 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	v106 = int32(0)
	goto L20
L20:
	;
	v114 = l0 + int32(51) + v106<<(uint(int32(1))%32)
	v115 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v114))))
	if v98 != v115 {
		goto L22
	} else {
		goto L23
	}
L21:
	;
	if base.Ui32(v93) < base.Ui32(int32(60)) {
		v164 = v92
		v165 = v93
		goto L18
	} else {
		goto L26
	}
L22:
	;
	v121 = v106 + int32(1)
	if v121 != v93 {
		v106 = v121
		goto L20
	} else {
		goto L25
	}
L23:
	;
	v117 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	v118 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v114)+1)))
	if v117 == v118 {
		v248 = v106
		goto L17
	} else {
		goto L24
	}
L24:
	;
	goto L22
L25:
	;
	goto L21
L26:
	;
	v125 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v126 = *(*int32)(unsafe.Add(mBase, uint32(v91)+60))
	if v126 != 0 {
		goto L28
	} else {
		goto L29
	}
L27:
	;
	v152 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	F_luaX_lexerror(m, v152, v151, int32(0))
	mBase = m.M
	v155 = m.ExcPending
	if v155 != 0 {
		goto L14
	} else {
		goto L32
	}
L28:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+20)) = int32(60)
	*(*int32)(unsafe.Add(mBase, uint32(v15)+16)) = v126
	v140 = m.G3
	*(*int32)(unsafe.Add(mBase, uint32(v15)+24)) = v140 + int32(_a2055)
	v148 = F_luaO_pushfstring(m, v125, v140+int32(_a2051), v15+int32(16))
	mBase = m.M
	v149 = m.ExcPending
	if v149 != 0 {
		goto L14
	} else {
		goto L31
	}
L29:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15))) = int32(60)
	v129 = m.G3
	*(*int32)(unsafe.Add(mBase, uint32(v15)+4)) = v129 + int32(_a2055)
	v135 = F_luaO_pushfstring(m, v125, v129+int32(_a2052), v15)
	mBase = m.M
	v136 = m.ExcPending
	if v136 != 0 {
		goto L14
	} else {
		goto L30
	}
L30:
	;
	v151 = v135
	goto L27
L31:
	;
	v151 = v148
	goto L27
L32:
	;
	v156 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v91)+72)))
	v157 = *(*int32)(unsafe.Add(mBase, uint32(v91)+36))
	v164 = v157
	v165 = v156
	goto L18
L33:
	;
	if v186 <= v92 {
		goto L37
	} else {
		goto L38
	}
L34:
	;
	v172 = m.G3
	v173 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v174 = *(*int32)(unsafe.Add(mBase, uint32(v91)+28))
	v181 = F_luaM_growaux_(m, v173, v174, v91+int32(36), int32(4), int32(2147483645), v172+int32(_a139))
	mBase = m.M
	v182 = m.ExcPending
	if v182 != 0 {
		goto L14
	} else {
		goto L36
	}
L35:
	;
	v171 = *(*int32)(unsafe.Add(mBase, uint32(v91)+28))
	v185 = v171
	v186 = v164
	goto L33
L36:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v91)+28)) = v181
	v184 = *(*int32)(unsafe.Add(mBase, uint32(v91)+36))
	v185 = v181
	v186 = v184
	goto L33
L37:
	;
	v197 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v91)+72)))
	*(*int32)(unsafe.Add(mBase, uint32(v185+v197<<(uint(int32(2))%32)))) = l1
	v202 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+5)))
	if v202&int32(3) == int32(0) {
		v227 = v197
		goto L40
	} else {
		goto L41
	}
L38:
	;
	v188 = int32(2)
	v196 = F__emscripten_memset_bulkmem(m, v185+v92<<(uint(v188)%32), base.I32_extend8_s(int32(0)), (v186-v92)<<(uint(v188)%32))
	mBase = m.M
	goto L39
L39:
	;
	goto L37
L40:
	;
	v228 = int32(1)
	v230 = l0 + v227<<(uint(v228)%32)
	v233 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	*(*uint8)(unsafe.Add(mBase, uint32(v230+int32(51)))) = uint8(v233)
	v237 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	*(*uint8)(unsafe.Add(mBase, uint32(v230+int32(52)))) = uint8(v237)
	v240 = v227 + v228
	*(*uint8)(unsafe.Add(mBase, uint32(v91)+72)) = uint8(v240)
	v248 = v227
	goto L17
L41:
	;
	v207 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v91)+5)))
	if v207&int32(4) == int32(0) {
		v227 = v197
		goto L40
	} else {
		goto L42
	}
L42:
	;
	v212 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v213 = *(*int32)(unsafe.Add(mBase, uint32(v212)+16))
	v214 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v213)+21)))
	if v214 != int32(1) {
		goto L44
	} else {
		goto L45
	}
L43:
	;
	v226 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v91)+72)))
	v227 = v226
	goto L40
L44:
	;
	v218 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v213)+20)))
	v221 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v91)+5)))
	v224 = v218&int32(3) | v221&int32(248)
	*(*uint8)(unsafe.Add(mBase, uint32(v91)+5)) = uint8(v224)
	goto L43
L45:
	;
	F_reallymarkobject(m, v213, l1)
	mBase = m.M
	goto L43
}
func F_sinh(m *base.Module, l0 float64) float64 {
	var v6 float64
	_ = v6
	var v7 float64
	_ = v7
	var v8 int64
	_ = v8
	var v11 float64
	_ = v11
	var v31 float64
	_ = v31
	var v35 float64
	_ = v35
	var v39 float64
	_ = v39
	v6 = base.F64_copysign(float64(0.5), l0)
	v7 = base.F64_abs(l0)
	v8 = base.I64_reinterpret_f64(v7)
	if base.Ui64(int64(4649454526309335039)) < base.Ui64(v8) {
		v31 = float64(2.247116418577895e+307)
		v35 = F_exp(m, base.F64_add(v7, float64(-1416.0996898839683)))
		v39 = base.F64_mul(base.F64_mul(base.F64_mul(base.F64_add(v6, v6), v31), v35), v31)
		return v39
	} else {
		v11 = F_expm1(m, v7)
		if base.Ui64(int64(4607182418800017407)) < base.Ui64(v8) {
			return base.F64_mul(v6, base.F64_add(v11, base.F64_div(v11, base.F64_add(v11, float64(1)))))
		} else {
			if base.Ui64(v8) < base.Ui64(int64(4490088828488384512)) {
				v39 = l0
				return v39
			} else {
				return base.F64_mul(v6, base.F64_sub(base.F64_add(v11, v11), base.F64_div(base.F64_mul(v11, v11), base.F64_add(v11, float64(1)))))
			}
		}
	}
}
func F_sinterGenericCommand(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) {
	mBase := m.M
	_ = mBase
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v71 int32
	_ = v71
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v95 int64
	_ = v95
	var v100 int32
	_ = v100
	var v102 int32
	_ = v102
	var v107 int32
	_ = v107
	var v109 int32
	_ = v109
	var v113 int32
	_ = v113
	var v115 int32
	_ = v115
	var v138 int32
	_ = v138
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v159 int32
	_ = v159
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v193 int32
	_ = v193
	var v206 int32
	_ = v206
	var v216 int32
	_ = v216
	var v219 int32
	_ = v219
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v234 int32
	_ = v234
	var v248 int32
	_ = v248
	var v249 int32
	_ = v249
	var v251 int64
	_ = v251
	var v252 int32
	_ = v252
	var v253 int32
	_ = v253
	var v257 int32
	_ = v257
	var v267 int32
	_ = v267
	var v301 int32
	_ = v301
	var v303 int32
	_ = v303
	var v306 int32
	_ = v306
	var v308 int32
	_ = v308
	var v311 int64
	_ = v311
	var v313 int32
	_ = v313
	var v318 int32
	_ = v318
	var v319 int32
	_ = v319
	var v320 int32
	_ = v320
	var v324 int32
	_ = v324
	var v325 int32
	_ = v325
	var v337 int32
	_ = v337
	var v338 int32
	_ = v338
	var v342 int32
	_ = v342
	var v358 int32
	_ = v358
	var v360 int32
	_ = v360
	var v361 int32
	_ = v361
	var v362 int32
	_ = v362
	var v373 int64
	_ = v373
	var v379 int32
	_ = v379
	var v381 int64
	_ = v381
	var v383 int32
	_ = v383
	var v385 int32
	_ = v385
	var v395 int64
	_ = v395
	var v400 int64
	_ = v400
	var v404 int32
	_ = v404
	var v406 int64
	_ = v406
	var v408 int32
	_ = v408
	var v416 int64
	_ = v416
	var v440 int64
	_ = v440
	var v459 int32
	_ = v459
	var v468 int32
	_ = v468
	var v469 int32
	_ = v469
	var v476 int32
	_ = v476
	var v479 int32
	_ = v479
	var v480 int32
	_ = v480
	var v481 int32
	_ = v481
	var v482 int64
	_ = v482
	var v485 int32
	_ = v485
	var v486 int32
	_ = v486
	var v492 int32
	_ = v492
	var v505 int32
	_ = v505
	var v512 int32
	_ = v512
	var v513 int32
	_ = v513
	var v521 int32
	_ = v521
	var v534 int32
	_ = v534
	var v542 int32
	_ = v542
	var v545 int32
	_ = v545
	var v556 int32
	_ = v556
	var v559 int32
	_ = v559
	var v561 int32
	_ = v561
	var v563 int32
	_ = v563
	var v568 int32
	_ = v568
	var v571 int32
	_ = v571
	var v572 int32
	_ = v572
	var v575 int32
	_ = v575
	var v581 int32
	_ = v581
	var v582 int32
	_ = v582
	var v584 int32
	_ = v584
	var v585 int32
	_ = v585
	var v588 int32
	_ = v588
	var v590 int32
	_ = v590
	var v591 int32
	_ = v591
	var v592 int32
	_ = v592
	var v593 int32
	_ = v593
	var v594 int32
	_ = v594
	var v601 int32
	_ = v601
	var v602 int32
	_ = v602
	var v613 int32
	_ = v613
	var v624 int32
	_ = v624
	var v627 int32
	_ = v627
	var v629 int32
	_ = v629
	var v630 int32
	_ = v630
	var v642 int32
	_ = v642
	var v643 int32
	_ = v643
	var v647 int32
	_ = v647
	var v663 int32
	_ = v663
	var v665 int32
	_ = v665
	var v666 int32
	_ = v666
	var v667 int32
	_ = v667
	var v678 int64
	_ = v678
	var v684 int32
	_ = v684
	var v686 int64
	_ = v686
	var v688 int32
	_ = v688
	var v690 int32
	_ = v690
	var v700 int64
	_ = v700
	var v705 int64
	_ = v705
	var v709 int32
	_ = v709
	var v711 int64
	_ = v711
	var v713 int32
	_ = v713
	var v721 int64
	_ = v721
	var v745 int64
	_ = v745
	var v764 int32
	_ = v764
	var v773 int32
	_ = v773
	var v775 int64
	_ = v775
	var v778 int32
	_ = v778
	var v779 int32
	_ = v779
	var v780 int32
	_ = v780
	var v789 int32
	_ = v789
	var v790 int32
	_ = v790
	var v801 int32
	_ = v801
	var v812 int32
	_ = v812
	var v815 int32
	_ = v815
	var v817 int32
	_ = v817
	var v819 int32
	_ = v819
	var v821 int32
	_ = v821
	var v823 int32
	_ = v823
	var v824 int32
	_ = v824
	var v828 int32
	_ = v828
	var v841 int32
	_ = v841
	var v853 int32
	_ = v853
	var v854 int32
	_ = v854
	var v855 int32
	_ = v855
	var v857 int32
	_ = v857
	var v858 int32
	_ = v858
	var v863 int32
	_ = v863
	var v866 int32
	_ = v866
	var v867 int32
	_ = v867
	var v869 int32
	_ = v869
	var v870 int32
	_ = v870
	var v872 int64
	_ = v872
	var v876 int32
	_ = v876
	var v877 int32
	_ = v877
	var v878 int32
	_ = v878
	var v881 int32
	_ = v881
	var v882 int32
	_ = v882
	var v883 int32
	_ = v883
	var v884 int32
	_ = v884
	var v887 int32
	_ = v887
	var v889 int64
	_ = v889
	var v893 int32
	_ = v893
	var v895 int32
	_ = v895
	var v898 int32
	_ = v898
	var v899 int32
	_ = v899
	var v901 int32
	_ = v901
	var v903 int32
	_ = v903
	var v905 int32
	_ = v905
	var v907 int32
	_ = v907
	var v909 int32
	_ = v909
	var v930 int32
	_ = v930
	var v957 int32
	_ = v957
	var v963 int32
	_ = v963
	v20 = m.G0
	v22 = v20 - int32(48)
	m.G0 = v22
	v26 = F_valkey_malloc(m, l2<<(uint(int32(2))%32))
	mBase = m.M
	v27 = m.ExcPending
	if v27 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v28 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v22)+24)) = v28
	if l2 == v28 {
		goto L7
	} else {
		goto L8
	}
L3:
	;
	F__serverAssert(m, int32(_a1688), int32(_a1683), int32(105))
	mBase = m.M
	v963 = m.ExcPending
	if v963 != 0 {
		goto L1
	} else {
		goto L200
	}
L4:
	;
	F__serverAssert(m, int32(_a1689), int32(_a1683), int32(101))
	mBase = m.M
	v957 = m.ExcPending
	if v957 != 0 {
		goto L1
	} else {
		goto L199
	}
L5:
	;
	m.G0 = v22 + int32(48)
	return
L6:
	;
	F_valkey_free(m, v26)
	mBase = m.M
	v930 = m.ExcPending
	if v930 != 0 {
		goto L1
	} else {
		goto L198
	}
L7:
	;
	F_qsort(m, v26, l2, int32(4), int32(1091))
	mBase = m.M
	v138 = m.ExcPending
	if v138 != 0 {
		goto L1
	} else {
		goto L32
	}
L8:
	;
	v42 = v28
	v43 = int32(0)
	goto L9
L9:
	;
	v53 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v55 = v42 << (uint(int32(2)) % 32)
	v57 = *(*int32)(unsafe.Add(mBase, uint32(l1+v55)))
	v58 = F_lookupKeyRead(m, v53, v57)
	mBase = m.M
	v59 = m.ExcPending
	if v59 != 0 {
		goto L1
	} else {
		goto L13
	}
L10:
	;
	if v66 < int32(1) {
		goto L7
	} else {
		goto L18
	}
L11:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26+v55))) = v67
	v71 = v42 + int32(1)
	if v71 != l2 {
		v42 = v71
		v43 = v66
		goto L9
	} else {
		goto L17
	}
L12:
	;
	v64 = F_checkType(m, l0, v58, int32(2))
	mBase = m.M
	v65 = m.ExcPending
	if v65 != 0 {
		goto L1
	} else {
		goto L15
	}
L13:
	;
	if v58 != 0 {
		goto L12
	} else {
		goto L14
	}
L14:
	;
	v66 = v43 + int32(1)
	v67 = int32(0)
	goto L11
L15:
	;
	if v64 != 0 {
		goto L6
	} else {
		goto L16
	}
L16:
	;
	v66 = v43
	v67 = v58
	goto L11
L17:
	;
	goto L10
L18:
	;
	F_valkey_free(m, v26)
	mBase = m.M
	v76 = m.ExcPending
	if v76 != 0 {
		goto L1
	} else {
		goto L19
	}
L19:
	;
	if l3 == int32(0) {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	if l4 == int32(0) {
		goto L28
	} else {
		goto L29
	}
L21:
	;
	v79 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v80 = F_dbDelete(m, v79, l3)
	mBase = m.M
	v81 = m.ExcPending
	if v81 != 0 {
		goto L1
	} else {
		goto L23
	}
L22:
	;
	v100 = *(*int32)(unsafe.Add(mBase, _consts[85]))
	F_addReply(m, l0, v100)
	mBase = m.M
	v102 = m.ExcPending
	if v102 != 0 {
		goto L1
	} else {
		goto L27
	}
L23:
	;
	if v80 == int32(0) {
		goto L22
	} else {
		goto L24
	}
L24:
	;
	v84 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	F_signalModifiedKey(m, l0, v84, l3)
	mBase = m.M
	v86 = m.ExcPending
	if v86 != 0 {
		goto L1
	} else {
		goto L25
	}
L25:
	;
	v89 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v90 = *(*int32)(unsafe.Add(mBase, uint32(v89)+28))
	F_notifyKeyspaceEvent(m, int32(4), int32(_a213), l3, v90)
	mBase = m.M
	v92 = m.ExcPending
	if v92 != 0 {
		goto L1
	} else {
		goto L26
	}
L26:
	;
	v93 = int32(_a44)
	v95 = *(*int64)(unsafe.Add(mBase, _consts[83]))
	*(*int64)(unsafe.Add(mBase, _consts[83])) = v95 + int64(1)
	goto L22
L27:
	;
	goto L5
L28:
	;
	v109 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+224)))
	v113 = *(*int32)(unsafe.Add(mBase, uint32(v109<<(uint(int32(2))%32))+uint32(_consts[929])))
	F_addReply(m, l0, v113)
	mBase = m.M
	v115 = m.ExcPending
	if v115 != 0 {
		goto L1
	} else {
		goto L31
	}
L29:
	;
	F_addReplyLongLong(m, l0, int64(0))
	mBase = m.M
	v107 = m.ExcPending
	if v107 != 0 {
		goto L1
	} else {
		goto L30
	}
L30:
	;
	goto L5
L31:
	;
	goto L5
L32:
	;
	if l3 == int32(0) {
		goto L35
	} else {
		goto L36
	}
L33:
	;
	v180 = *(*int32)(unsafe.Add(mBase, uint32(v26)))
	v181 = F_setTypeInitIterator(m, v180)
	mBase = m.M
	v182 = m.ExcPending
	if v182 != 0 {
		goto L1
	} else {
		goto L50
	}
L34:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+24)) = v174
	v178 = v174
	v179 = int32(0)
	goto L33
L35:
	;
	v168 = int32(0)
	if l4 == v168 {
		goto L45
	} else {
		goto L46
	}
L36:
	;
	v141 = *(*int32)(unsafe.Add(mBase, uint32(v26)))
	v142 = *(*int32)(unsafe.Add(mBase, uint32(v141)))
	switch int32(base.Ui32(v142)>>(uint(int32(4))%32))&int32(15) + int32(-6) {
	case 0:
		goto L39
	default:
		goto L37
	case 5:
		goto L38
	}
L37:
	;
	v166 = F_createSetListpackObject(m)
	mBase = m.M
	v167 = m.ExcPending
	if v167 != 0 {
		goto L1
	} else {
		goto L44
	}
L38:
	;
	v152 = F_objectGetVal(m, v141)
	mBase = m.M
	v153 = *(*int32)(unsafe.Add(mBase, uint32(v152)))
	goto L41
L39:
	;
	v149 = F_createIntsetObject(m)
	mBase = m.M
	v150 = m.ExcPending
	if v150 != 0 {
		goto L1
	} else {
		goto L40
	}
L40:
	;
	v174 = v149
	goto L34
L41:
	;
	v154 = F_lpNew(m, v153)
	mBase = m.M
	v155 = m.ExcPending
	if v155 != 0 {
		goto L1
	} else {
		goto L42
	}
L42:
	;
	v156 = F_createObject(m, int32(2), v154)
	mBase = m.M
	v157 = m.ExcPending
	if v157 != 0 {
		goto L1
	} else {
		goto L43
	}
L43:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+24)) = v156
	v159 = *(*int32)(unsafe.Add(mBase, uint32(v156)))
	*(*int32)(unsafe.Add(mBase, uint32(v156))) = v159&int32(-241) | int32(176)
	v178 = v156
	v179 = int32(0)
	goto L33
L44:
	;
	v174 = v166
	goto L34
L45:
	;
	v172 = F_addReplyDeferredLen(m, l0)
	mBase = m.M
	v173 = m.ExcPending
	if v173 != 0 {
		goto L1
	} else {
		goto L47
	}
L46:
	;
	v178 = v168
	v179 = int32(0)
	goto L33
L47:
	;
	v178 = v168
	v179 = v172
	goto L33
L48:
	;
	v556 = *(*int32)(unsafe.Add(mBase, uint32(v181)+4))
	if v556 != int32(2) {
		goto L115
	} else {
		goto L116
	}
L49:
	;
	v206 = int32(0)
	v216 = v189
	v219 = int32(1)
	goto L54
L50:
	;
	v189 = F_setTypeNext(m, v181, v22+int32(20), v22+int32(16), v22+int32(8))
	mBase = m.M
	v190 = m.ExcPending
	if v190 != 0 {
		goto L1
	} else {
		goto L51
	}
L51:
	;
	if v189 != int32(-1) {
		goto L49
	} else {
		goto L52
	}
L52:
	;
	v193 = int32(0)
	v542 = v193
	v545 = v193
	goto L48
L53:
	;
	v542 = v521
	v545 = base.B2i32(v534 == int32(0))
	goto L48
L54:
	;
	if base.Ui32(l2) < base.Ui32(int32(2)) {
		v267 = int32(1)
		goto L58
	} else {
		goto L59
	}
L55:
	;
	v521 = v492
	v534 = v505
	goto L53
L56:
	;
	v512 = F_setTypeNext(m, v181, v22+int32(20), v22+int32(16), v22+int32(8))
	mBase = m.M
	v513 = m.ExcPending
	if v513 != 0 {
		goto L1
	} else {
		goto L113
	}
L57:
	;
	if l4 == int32(0) {
		goto L68
	} else {
		goto L69
	}
L58:
	;
	if v267 != l2 {
		v492 = v206
		v505 = v219
		goto L56
	} else {
		goto L67
	}
L59:
	;
	v224 = *(*int32)(unsafe.Add(mBase, uint32(v22)+16))
	v225 = *(*int32)(unsafe.Add(mBase, uint32(v22)+20))
	v234 = int32(1)
	goto L60
L60:
	;
	v248 = *(*int32)(unsafe.Add(mBase, uint32(v26+v234<<(uint(int32(2))%32))))
	v249 = *(*int32)(unsafe.Add(mBase, uint32(v26)))
	if v248 == v249 {
		goto L62
	} else {
		goto L63
	}
L62:
	;
	v257 = v234 + int32(1)
	if v257 != l2 {
		v234 = v257
		goto L60
	} else {
		goto L66
	}
L63:
	;
	v251 = *(*int64)(unsafe.Add(mBase, uint32(v22)+8))
	v252 = F_setTypeIsMemberAux(m, v248, v225, v224, v251, base.B2i32(v216 == int32(2)))
	mBase = m.M
	v253 = m.ExcPending
	if v253 != 0 {
		goto L1
	} else {
		goto L64
	}
L64:
	;
	if v252 == int32(0) {
		v267 = v234
		goto L58
	} else {
		goto L65
	}
L65:
	;
	goto L62
L66:
	;
	goto L57
L67:
	;
	goto L57
L68:
	;
	v303 = *(*int32)(unsafe.Add(mBase, uint32(v22)+20))
	if l3 != 0 {
		goto L71
	} else {
		goto L72
	}
L69:
	;
	v301 = v206 + int32(1)
	if base.Ui32(v301) <= base.Ui32(l5+int32(-1)) {
		v492 = v301
		v505 = v219
		goto L56
	} else {
		goto L70
	}
L70:
	;
	v521 = v301
	v534 = v219
	goto L53
L71:
	;
	if v303 == int32(0) {
		goto L79
	} else {
		goto L80
	}
L72:
	;
	if v303 == int32(0) {
		goto L73
	} else {
		goto L74
	}
L73:
	;
	v311 = *(*int64)(unsafe.Add(mBase, uint32(v22)+8))
	F_addReplyBulkLongLong(m, l0, v311)
	mBase = m.M
	v313 = m.ExcPending
	if v313 != 0 {
		goto L1
	} else {
		goto L76
	}
L74:
	;
	v306 = *(*int32)(unsafe.Add(mBase, uint32(v22)+16))
	F_addReplyBulkCBuffer(m, l0, v303, v306)
	mBase = m.M
	v308 = m.ExcPending
	if v308 != 0 {
		goto L1
	} else {
		goto L75
	}
L75:
	;
	v492 = v206 + int32(1)
	v505 = v219
	goto L56
L76:
	;
	v492 = v206 + int32(1)
	v505 = v219
	goto L56
L77:
	;
	v482 = *(*int64)(unsafe.Add(mBase, uint32(v22)+8))
	v485 = F_setTypeAddAux(m, v178, v479, v480, v482, base.B2i32(v216 == int32(2)))
	mBase = m.M
	v486 = m.ExcPending
	if v486 != 0 {
		goto L1
	} else {
		goto L112
	}
L78:
	;
	v319 = int32(0)
	v320 = *(*int32)(unsafe.Add(mBase, uint32(v22)+16))
	if v216 != int32(2) {
		v479 = v303
		v480 = v320
		v481 = v319
		goto L77
	} else {
		goto L82
	}
L79:
	;
	v318 = *(*int32)(unsafe.Add(mBase, uint32(v22)+16))
	v479 = v303
	v480 = v318
	v481 = v219
	goto L77
L80:
	;
	if v219 != 0 {
		goto L78
	} else {
		goto L81
	}
L81:
	;
	goto L79
L82:
	;
	v324 = v22 + int32(8)
	v325 = int32(0)
	if base.Ui32(v320+int32(-21)) < base.Ui32(int32(-20)) {
		v459 = v325
		goto L84
	} else {
		goto L85
	}
L83:
	;
	if v459 == int32(0) {
		v479 = v303
		v480 = v320
		v481 = v319
		goto L77
	} else {
		goto L110
	}
L84:
	;
	goto L83
L85:
	;
	v337 = int32(1)
	v338 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v303))))
	if v320 != v337 {
		goto L88
	} else {
		goto L89
	}
L86:
	;
	v459 = int32(1)
	goto L84
L87:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v324))) = v440
	goto L86
L88:
	;
	if v338&int32(255) == int32(45) {
		goto L93
	} else {
		goto L94
	}
L89:
	;
	v342 = v338 + int32(-48)
	if base.Ui32(int32(9)) < base.Ui32(v342&int32(255)) {
		v459 = v325
		goto L84
	} else {
		goto L90
	}
L90:
	;
	if v324 == int32(0) {
		goto L86
	} else {
		goto L91
	}
L91:
	;
	v440 = base.I64_extend_i32_u(v342) & int64(255)
	goto L87
L92:
	;
	if base.Ui32(int32(8)) < base.Ui32((v361+int32(-49))&int32(255)) {
		v459 = v325
		goto L84
	} else {
		goto L95
	}
L93:
	;
	v358 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v303)+1)))
	v360 = int32(2)
	v361 = v358
	v362 = v303 + int32(1)
	goto L92
L94:
	;
	v360 = v337
	v361 = v338
	v362 = v303
	goto L92
L95:
	;
	v373 = base.I64_extend_i32_u(v361+int32(-48)) & int64(255)
	if base.Ui32(v320) <= base.Ui32(v360) {
		v416 = v373
		goto L96
	} else {
		goto L97
	}
L96:
	;
	if v338&int32(255) != int32(45) {
		goto L104
	} else {
		goto L105
	}
L97:
	;
	v379 = v360
	v381 = v373
	v383 = v362
	goto L98
L98:
	;
	v385 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v383)+1)))
	if base.Ui32((v385+int32(-58))&int32(255)) < base.Ui32(int32(246)) {
		v459 = v325
		goto L84
	} else {
		goto L100
	}
L99:
	;
	v416 = v406
	goto L96
L100:
	;
	if base.Ui64(int64(1844674407370955161)) < base.Ui64(v381) {
		v459 = v325
		goto L84
	} else {
		goto L101
	}
L101:
	;
	v395 = v381 * int64(10)
	v400 = base.I64_extend_i32_u(v385+int32(-48)) & int64(255)
	if base.Ui64(v400^int64(-1)) < base.Ui64(v395) {
		v459 = v325
		goto L84
	} else {
		goto L102
	}
L102:
	;
	v404 = int32(1)
	v406 = v395 + v400
	v408 = v379 + v404
	if v408 != v320 {
		v379 = v408
		v381 = v406
		v383 = v383 + v404
		goto L98
	} else {
		goto L103
	}
L103:
	;
	goto L99
L104:
	;
	if v416 < int64(0) {
		v459 = v325
		goto L84
	} else {
		goto L108
	}
L105:
	;
	if base.Ui64(int64(-9223372036854775807-1)) < base.Ui64(v416) {
		v459 = v325
		goto L84
	} else {
		goto L106
	}
L106:
	;
	if v324 == int32(0) {
		goto L86
	} else {
		goto L107
	}
L107:
	;
	v440 = int64(0) - v416
	goto L87
L108:
	;
	if v324 == int32(0) {
		goto L86
	} else {
		goto L109
	}
L109:
	;
	v440 = v416
	goto L87
L110:
	;
	v468 = int32(1)
	v469 = *(*int32)(unsafe.Add(mBase, uint32(v178)))
	switch int32(base.Ui32(v469)>>(uint(int32(4))%32))&int32(15) + int32(-6) {
	case 0, 5:
		goto L111
	default:
		v479 = v303
		v480 = v320
		v481 = v468
		goto L77
	}
L111:
	;
	v476 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v22)+20)) = v476
	v479 = v476
	v480 = v320
	v481 = v468
	goto L77
L112:
	;
	v492 = v206
	v505 = v481
	goto L56
L113:
	;
	if v512 != int32(-1) {
		v206 = v492
		v216 = v512
		v219 = v505
		goto L54
	} else {
		goto L114
	}
L114:
	;
	goto L55
L115:
	;
	F_valkey_free(m, v181)
	mBase = m.M
	v563 = m.ExcPending
	if v563 != 0 {
		goto L1
	} else {
		goto L118
	}
L116:
	;
	v559 = *(*int32)(unsafe.Add(mBase, uint32(v181)+12))
	F_hashtableReleaseIterator(m, v559)
	mBase = m.M
	v561 = m.ExcPending
	if v561 != 0 {
		goto L1
	} else {
		goto L117
	}
L117:
	;
	goto L115
L118:
	;
	if l4 == int32(0) {
		goto L119
	} else {
		goto L120
	}
L119:
	;
	if l3 == int32(0) {
		goto L122
	} else {
		goto L123
	}
L120:
	;
	F_addReplyLongLong(m, l0, base.I64_extend_i32_u(v542))
	mBase = m.M
	v568 = m.ExcPending
	if v568 != 0 {
		goto L1
	} else {
		goto L121
	}
L121:
	;
	goto L6
L122:
	;
	F_setDeferredSetLen(m, l0, v179, v542)
	mBase = m.M
	v909 = m.ExcPending
	if v909 != 0 {
		goto L1
	} else {
		goto L197
	}
L123:
	;
	v571 = F_setTypeSize(m, v178)
	mBase = m.M
	v572 = m.ExcPending
	if v572 != 0 {
		goto L1
	} else {
		goto L125
	}
L124:
	;
	v882 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v883 = F_dbDelete(m, v882, l3)
	mBase = m.M
	v884 = m.ExcPending
	if v884 != 0 {
		goto L1
	} else {
		goto L191
	}
L125:
	;
	if v571 == int32(0) {
		goto L124
	} else {
		goto L126
	}
L126:
	;
	v575 = *(*int32)(unsafe.Add(mBase, uint32(v178)))
	if v545|base.B2i32(v575&int32(240) == int32(96)) != 0 {
		v841 = v575
		goto L127
	} else {
		goto L128
	}
L127:
	;
	if v841&int32(240) != int32(176) {
		goto L182
	} else {
		goto L183
	}
L128:
	;
	v581 = F_setTypeSize(m, v178)
	mBase = m.M
	v582 = m.ExcPending
	if v582 != 0 {
		goto L1
	} else {
		goto L130
	}
L129:
	;
	v591 = F_intsetNew(m)
	mBase = m.M
	v592 = m.ExcPending
	if v592 != 0 {
		goto L1
	} else {
		goto L135
	}
L130:
	;
	v584 = *(*int32)(unsafe.Add(mBase, _consts[930]))
	v585 = int32(1073741824)
	if base.Ui32(v584) < base.Ui32(v585) {
		goto L131
	} else {
		goto L132
	}
L131:
	;
	v588 = v584
	goto L133
L132:
	;
	v588 = v585
	goto L133
L133:
	;
	if base.Ui32(v581) <= base.Ui32(v588) {
		goto L129
	} else {
		goto L134
	}
L134:
	;
	v590 = *(*int32)(unsafe.Add(mBase, uint32(v178)))
	v841 = v590
	goto L127
L135:
	;
	v593 = F_setTypeInitIterator(m, v178)
	mBase = m.M
	v594 = m.ExcPending
	if v594 != 0 {
		goto L1
	} else {
		goto L137
	}
L136:
	;
	v812 = *(*int32)(unsafe.Add(mBase, uint32(v593)+4))
	if v812 != int32(2) {
		goto L176
	} else {
		goto L177
	}
L137:
	;
	v601 = F_setTypeNext(m, v593, v22+int32(44), v22+int32(40), v22+int32(32))
	mBase = m.M
	v602 = m.ExcPending
	if v602 != 0 {
		goto L1
	} else {
		goto L138
	}
L138:
	;
	if v601 == int32(-1) {
		v801 = v591
		goto L136
	} else {
		goto L139
	}
L139:
	;
	v613 = v591
	goto L140
L140:
	;
	v624 = *(*int32)(unsafe.Add(mBase, uint32(v22)+44))
	if v624 == int32(0) {
		goto L142
	} else {
		goto L143
	}
L141:
	;
	v801 = v778
	goto L136
L142:
	;
	v773 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v22)+31)) = uint8(v773)
	v775 = *(*int64)(unsafe.Add(mBase, uint32(v22)+32))
	v778 = F_intsetAdd(m, v613, v775, v22+int32(31))
	mBase = m.M
	v779 = m.ExcPending
	if v779 != 0 {
		goto L1
	} else {
		goto L172
	}
L143:
	;
	v627 = *(*int32)(unsafe.Add(mBase, uint32(v22)+40))
	v629 = v22 + int32(32)
	v630 = int32(0)
	if base.Ui32(v627+int32(-21)) < base.Ui32(int32(-20)) {
		v764 = v630
		goto L145
	} else {
		goto L146
	}
L144:
	;
	if v764 == int32(0) {
		goto L4
	} else {
		goto L171
	}
L145:
	;
	goto L144
L146:
	;
	v642 = int32(1)
	v643 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v624))))
	if v627 != v642 {
		goto L149
	} else {
		goto L150
	}
L147:
	;
	v764 = int32(1)
	goto L145
L148:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v629))) = v745
	goto L147
L149:
	;
	if v643&int32(255) == int32(45) {
		goto L154
	} else {
		goto L155
	}
L150:
	;
	v647 = v643 + int32(-48)
	if base.Ui32(int32(9)) < base.Ui32(v647&int32(255)) {
		v764 = v630
		goto L145
	} else {
		goto L151
	}
L151:
	;
	if v629 == int32(0) {
		goto L147
	} else {
		goto L152
	}
L152:
	;
	v745 = base.I64_extend_i32_u(v647) & int64(255)
	goto L148
L153:
	;
	if base.Ui32(int32(8)) < base.Ui32((v666+int32(-49))&int32(255)) {
		v764 = v630
		goto L145
	} else {
		goto L156
	}
L154:
	;
	v663 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v624)+1)))
	v665 = int32(2)
	v666 = v663
	v667 = v624 + int32(1)
	goto L153
L155:
	;
	v665 = v642
	v666 = v643
	v667 = v624
	goto L153
L156:
	;
	v678 = base.I64_extend_i32_u(v666+int32(-48)) & int64(255)
	if base.Ui32(v627) <= base.Ui32(v665) {
		v721 = v678
		goto L157
	} else {
		goto L158
	}
L157:
	;
	if v643&int32(255) != int32(45) {
		goto L165
	} else {
		goto L166
	}
L158:
	;
	v684 = v665
	v686 = v678
	v688 = v667
	goto L159
L159:
	;
	v690 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v688)+1)))
	if base.Ui32((v690+int32(-58))&int32(255)) < base.Ui32(int32(246)) {
		v764 = v630
		goto L145
	} else {
		goto L161
	}
L160:
	;
	v721 = v711
	goto L157
L161:
	;
	if base.Ui64(int64(1844674407370955161)) < base.Ui64(v686) {
		v764 = v630
		goto L145
	} else {
		goto L162
	}
L162:
	;
	v700 = v686 * int64(10)
	v705 = base.I64_extend_i32_u(v690+int32(-48)) & int64(255)
	if base.Ui64(v705^int64(-1)) < base.Ui64(v700) {
		v764 = v630
		goto L145
	} else {
		goto L163
	}
L163:
	;
	v709 = int32(1)
	v711 = v700 + v705
	v713 = v684 + v709
	if v713 != v627 {
		v684 = v713
		v686 = v711
		v688 = v688 + v709
		goto L159
	} else {
		goto L164
	}
L164:
	;
	goto L160
L165:
	;
	if v721 < int64(0) {
		v764 = v630
		goto L145
	} else {
		goto L169
	}
L166:
	;
	if base.Ui64(int64(-9223372036854775807-1)) < base.Ui64(v721) {
		v764 = v630
		goto L145
	} else {
		goto L167
	}
L167:
	;
	if v629 == int32(0) {
		goto L147
	} else {
		goto L168
	}
L168:
	;
	v745 = int64(0) - v721
	goto L148
L169:
	;
	if v629 == int32(0) {
		goto L147
	} else {
		goto L170
	}
L170:
	;
	v745 = v721
	goto L148
L171:
	;
	goto L142
L172:
	;
	v780 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22)+31)))
	if v780 == int32(0) {
		goto L3
	} else {
		goto L173
	}
L173:
	;
	v789 = F_setTypeNext(m, v593, v22+int32(44), v22+int32(40), v22+int32(32))
	mBase = m.M
	v790 = m.ExcPending
	if v790 != 0 {
		goto L1
	} else {
		goto L174
	}
L174:
	;
	if v789 != int32(-1) {
		v613 = v778
		goto L140
	} else {
		goto L175
	}
L175:
	;
	goto L141
L176:
	;
	F_valkey_free(m, v593)
	mBase = m.M
	v819 = m.ExcPending
	if v819 != 0 {
		goto L1
	} else {
		goto L179
	}
L177:
	;
	v815 = *(*int32)(unsafe.Add(mBase, uint32(v593)+12))
	F_hashtableReleaseIterator(m, v815)
	mBase = m.M
	v817 = m.ExcPending
	if v817 != 0 {
		goto L1
	} else {
		goto L178
	}
L178:
	;
	goto L176
L179:
	;
	F_freeSetObject(m, v178)
	mBase = m.M
	v821 = m.ExcPending
	if v821 != 0 {
		goto L1
	} else {
		goto L180
	}
L180:
	;
	F_objectSetVal(m, v178, v801)
	mBase = m.M
	v823 = m.ExcPending
	if v823 != 0 {
		goto L1
	} else {
		goto L181
	}
L181:
	;
	v824 = *(*int32)(unsafe.Add(mBase, uint32(v178)))
	v828 = v824&int32(-241) | int32(96)
	*(*int32)(unsafe.Add(mBase, uint32(v178))) = v828
	v841 = v828
	goto L127
L182:
	;
	v858 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	F_setKey(m, l0, v858, l3, v22+int32(24), int32(0))
	mBase = m.M
	v863 = m.ExcPending
	if v863 != 0 {
		goto L1
	} else {
		goto L186
	}
L183:
	;
	v853 = F_objectGetVal(m, v178)
	mBase = m.M
	v854 = F_lpShrinkToFit(m, v853)
	mBase = m.M
	v855 = m.ExcPending
	if v855 != 0 {
		goto L1
	} else {
		goto L184
	}
L184:
	;
	F_objectSetVal(m, v178, v854)
	mBase = m.M
	v857 = m.ExcPending
	if v857 != 0 {
		goto L1
	} else {
		goto L185
	}
L185:
	;
	goto L182
L186:
	;
	v866 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v867 = *(*int32)(unsafe.Add(mBase, uint32(v866)+28))
	F_notifyKeyspaceEvent(m, int32(32), int32(_a1690), l3, v867)
	mBase = m.M
	v869 = m.ExcPending
	if v869 != 0 {
		goto L1
	} else {
		goto L187
	}
L187:
	;
	v870 = int32(_a44)
	v872 = *(*int64)(unsafe.Add(mBase, _consts[83]))
	*(*int64)(unsafe.Add(mBase, _consts[83])) = v872 + int64(1)
	v876 = *(*int32)(unsafe.Add(mBase, uint32(v22)+24))
	v877 = F_setTypeSize(m, v876)
	mBase = m.M
	v878 = m.ExcPending
	if v878 != 0 {
		goto L1
	} else {
		goto L188
	}
L188:
	;
	F_addReplyLongLong(m, l0, base.I64_extend_i32_u(v877))
	mBase = m.M
	v881 = m.ExcPending
	if v881 != 0 {
		goto L1
	} else {
		goto L189
	}
L189:
	;
	goto L6
L190:
	;
	v903 = *(*int32)(unsafe.Add(mBase, _consts[85]))
	F_addReply(m, l0, v903)
	mBase = m.M
	v905 = m.ExcPending
	if v905 != 0 {
		goto L1
	} else {
		goto L195
	}
L191:
	;
	if v883 == int32(0) {
		goto L190
	} else {
		goto L192
	}
L192:
	;
	v887 = int32(_a44)
	v889 = *(*int64)(unsafe.Add(mBase, _consts[83]))
	*(*int64)(unsafe.Add(mBase, _consts[83])) = v889 + int64(1)
	v893 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	F_signalModifiedKey(m, l0, v893, l3)
	mBase = m.M
	v895 = m.ExcPending
	if v895 != 0 {
		goto L1
	} else {
		goto L193
	}
L193:
	;
	v898 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v899 = *(*int32)(unsafe.Add(mBase, uint32(v898)+28))
	F_notifyKeyspaceEvent(m, int32(4), int32(_a213), l3, v899)
	mBase = m.M
	v901 = m.ExcPending
	if v901 != 0 {
		goto L1
	} else {
		goto L194
	}
L194:
	;
	goto L190
L195:
	;
	F_decrRefCount(m, v178)
	mBase = m.M
	v907 = m.ExcPending
	if v907 != 0 {
		goto L1
	} else {
		goto L196
	}
L196:
	;
	goto L6
L197:
	;
	goto L6
L198:
	;
	goto L5
L199:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L200:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_sintercardGetKeys(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	v6 = int32(1)
	v9 = F_genericGetKeys(m, int32(0), v6, int32(2), v6, l1, l2, l3)
	v12 = m.ExcPending
	if v12 != 0 {
		return int32(0)
	} else {
		return v9
	}
}
func F_siphash_nocase(m *base.Module, l0 int32, l1 int32, l2 int32) int64 {
	mBase := m.M
	_ = mBase
	var v13 int64
	_ = v13
	var v15 int64
	_ = v15
	var v16 int64
	_ = v16
	var v18 int64
	_ = v18
	var v20 int64
	_ = v20
	var v22 int64
	_ = v22
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v32 int64
	_ = v32
	var v33 int64
	_ = v33
	var v34 int64
	_ = v34
	var v35 int64
	_ = v35
	var v41 int32
	_ = v41
	var v50 int32
	_ = v50
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v101 int32
	_ = v101
	var v105 int32
	_ = v105
	var v114 int32
	_ = v114
	var v119 int32
	_ = v119
	var v128 int32
	_ = v128
	var v133 int32
	_ = v133
	var v142 int32
	_ = v142
	var v146 int64
	_ = v146
	var v147 int64
	_ = v147
	var v150 int64
	_ = v150
	var v151 int64
	_ = v151
	var v154 int64
	_ = v154
	var v155 int64
	_ = v155
	var v157 int64
	_ = v157
	var v158 int64
	_ = v158
	var v159 int64
	_ = v159
	var v162 int64
	_ = v162
	var v163 int64
	_ = v163
	var v165 int64
	_ = v165
	var v168 int64
	_ = v168
	var v170 int32
	_ = v170
	var v172 int32
	_ = v172
	var v175 int64
	_ = v175
	var v176 int64
	_ = v176
	var v177 int64
	_ = v177
	var v178 int64
	_ = v178
	var v185 int64
	_ = v185
	var v186 int32
	_ = v186
	var v195 int32
	_ = v195
	var v201 int64
	_ = v201
	var v202 int32
	_ = v202
	var v211 int32
	_ = v211
	var v217 int64
	_ = v217
	var v218 int32
	_ = v218
	var v227 int32
	_ = v227
	var v233 int64
	_ = v233
	var v234 int32
	_ = v234
	var v236 int32
	_ = v236
	var v245 int32
	_ = v245
	var v250 int64
	_ = v250
	var v251 int32
	_ = v251
	var v253 int32
	_ = v253
	var v262 int32
	_ = v262
	var v267 int64
	_ = v267
	var v268 int32
	_ = v268
	var v270 int32
	_ = v270
	var v279 int32
	_ = v279
	var v284 int64
	_ = v284
	var v285 int32
	_ = v285
	var v294 int32
	_ = v294
	var v300 int64
	_ = v300
	var v301 int64
	_ = v301
	var v302 int64
	_ = v302
	var v304 int64
	_ = v304
	var v305 int64
	_ = v305
	var v306 int64
	_ = v306
	var v308 int64
	_ = v308
	var v309 int64
	_ = v309
	var v311 int64
	_ = v311
	var v312 int64
	_ = v312
	var v315 int64
	_ = v315
	var v317 int64
	_ = v317
	var v318 int64
	_ = v318
	var v323 int64
	_ = v323
	var v324 int64
	_ = v324
	var v328 int64
	_ = v328
	var v330 int64
	_ = v330
	var v331 int64
	_ = v331
	var v334 int64
	_ = v334
	var v335 int64
	_ = v335
	var v340 int64
	_ = v340
	var v341 int64
	_ = v341
	var v344 int64
	_ = v344
	var v350 int64
	_ = v350
	var v354 int64
	_ = v354
	var v358 int64
	_ = v358
	v13 = *(*int64)(unsafe.Add(mBase, uint32(l2)))
	v15 = v13 ^ int64(8317987319222330741)
	v16 = *(*int64)(unsafe.Add(mBase, uint32(l2)+8))
	v18 = v16 ^ int64(7237128888997146477)
	v20 = v13 ^ int64(7816392313619706465)
	v22 = v16 ^ int64(8387220255154660723)
	v26 = l1 & int32(7)
	v27 = l0 + l1 - v26
	if l0 == v27 {
		v172 = l0
		v175 = v20
		v176 = v15
		v177 = v22
		v178 = v18
	} else {
		v29 = l0
		v32 = v20
		v33 = v15
		v34 = v22
		v35 = v18
		for {
			v41 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v29)+4)))
			if base.Ui32((v41+int32(-65))&int32(255)) < base.Ui32(int32(26)) {
				v50 = v41 | int32(32)
			} else {
				v50 = v41
			}
			v54 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v29)+1)))
			v56 = v54 << (uint(int32(8)) % 32)
			if base.Ui32((v54+int32(-65))&int32(255)) < base.Ui32(int32(26)) {
				v65 = v56 | int32(8192)
			} else {
				v65 = v56
			}
			v66 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v29))))
			if base.Ui32((v66+int32(-65))&int32(255)) < base.Ui32(int32(26)) {
				v75 = v66 | int32(32)
			} else {
				v75 = v66
			}
			v77 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v29)+2)))
			v79 = v77 << (uint(int32(16)) % 32)
			if base.Ui32((v77+int32(-65))&int32(255)) < base.Ui32(int32(26)) {
				v88 = v79 | int32(2097152)
			} else {
				v88 = v79
			}
			v90 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v29)+3)))
			v92 = v90 << (uint(int32(24)) % 32)
			if base.Ui32((v90+int32(-65))&int32(255)) < base.Ui32(int32(26)) {
				v101 = v92 | int32(536870912)
			} else {
				v101 = v92
			}
			v105 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v29)+5)))
			if base.Ui32((v105+int32(-65))&int32(255)) < base.Ui32(int32(26)) {
				v114 = v105 | int32(32)
			} else {
				v114 = v105
			}
			v119 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v29)+6)))
			if base.Ui32((v119+int32(-65))&int32(255)) < base.Ui32(int32(26)) {
				v128 = v119 | int32(32)
			} else {
				v128 = v119
			}
			v133 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v29)+7)))
			if base.Ui32((v133+int32(-65))&int32(255)) < base.Ui32(int32(26)) {
				v142 = v133 | int32(32)
			} else {
				v142 = v133
			}
			v146 = base.I64_extend_i32_u(v50)<<(uint(int64(32))%64) | base.I64_extend_i32_u(v65|v75|v88|v101) | base.I64_extend_i32_u(v114)<<(uint(int64(40))%64) | base.I64_extend_i32_u(v128)<<(uint(int64(48))%64) | base.I64_extend_i32_u(v142)<<(uint(int64(56))%64)
			v147 = v146 ^ v34
			v150 = v147 + v32
			v151 = base.I64_rotl(v147, int64(16)) ^ v150
			v154 = v33 + v35
			v155 = int64(32)
			v157 = v151 + base.I64_rotl(v154, v155)
			v158 = base.I64_rotl(v151, int64(21)) ^ v157
			v159 = v157 ^ v146
			v162 = v154 ^ base.I64_rotl(v35, int64(13))
			v163 = v150 + v162
			v165 = base.I64_rotl(v163, v155)
			v168 = v163 ^ base.I64_rotl(v162, int64(17))
			v170 = v29 + int32(8)
			if v170 != v27 {
				v29 = v170
				v32 = v165
				v33 = v159
				v34 = v158
				v35 = v168
				continue
			} else {
				break
			}
			break
		}
		v172 = v27
		v175 = v165
		v176 = v159
		v177 = v158
		v178 = v168
	}
	v185 = base.I64_extend_i32_u(l1) << (uint(int64(56)) % 64)
	switch v26 {
	default:
		v300 = v185
	case 1:
		v284 = v185
		v285 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v172))))
		if base.Ui32((v285+int32(-65))&int32(255)) < base.Ui32(int32(26)) {
			v294 = v285 | int32(32)
		} else {
			v294 = v285
		}
		v300 = v284 | base.I64_extend_i32_u(v294)
	case 2:
		v267 = v185
		v268 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v172)+1)))
		v270 = v268 << (uint(int32(8)) % 32)
		if base.Ui32((v268+int32(-65))&int32(255)) < base.Ui32(int32(26)) {
			v279 = v270 | int32(8192)
		} else {
			v279 = v270
		}
		v284 = v267 | base.I64_extend_i32_u(v279)
		v285 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v172))))
		if base.Ui32((v285+int32(-65))&int32(255)) < base.Ui32(int32(26)) {
			v294 = v285 | int32(32)
		} else {
			v294 = v285
		}
		v300 = v284 | base.I64_extend_i32_u(v294)
	case 3:
		v250 = v185
		v251 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v172)+2)))
		v253 = v251 << (uint(int32(16)) % 32)
		if base.Ui32((v251+int32(-65))&int32(255)) < base.Ui32(int32(26)) {
			v262 = v253 | int32(2097152)
		} else {
			v262 = v253
		}
		v267 = v250 | base.I64_extend_i32_u(v262)
		v268 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v172)+1)))
		v270 = v268 << (uint(int32(8)) % 32)
		if base.Ui32((v268+int32(-65))&int32(255)) < base.Ui32(int32(26)) {
			v279 = v270 | int32(8192)
		} else {
			v279 = v270
		}
		v284 = v267 | base.I64_extend_i32_u(v279)
		v285 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v172))))
		if base.Ui32((v285+int32(-65))&int32(255)) < base.Ui32(int32(26)) {
			v294 = v285 | int32(32)
		} else {
			v294 = v285
		}
		v300 = v284 | base.I64_extend_i32_u(v294)
	case 4:
		v233 = v185
		v234 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v172)+3)))
		v236 = v234 << (uint(int32(24)) % 32)
		if base.Ui32((v234+int32(-65))&int32(255)) < base.Ui32(int32(26)) {
			v245 = v236 | int32(536870912)
		} else {
			v245 = v236
		}
		v250 = v233 | base.I64_extend_i32_u(v245)
		v251 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v172)+2)))
		v253 = v251 << (uint(int32(16)) % 32)
		if base.Ui32((v251+int32(-65))&int32(255)) < base.Ui32(int32(26)) {
			v262 = v253 | int32(2097152)
		} else {
			v262 = v253
		}
		v267 = v250 | base.I64_extend_i32_u(v262)
		v268 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v172)+1)))
		v270 = v268 << (uint(int32(8)) % 32)
		if base.Ui32((v268+int32(-65))&int32(255)) < base.Ui32(int32(26)) {
			v279 = v270 | int32(8192)
		} else {
			v279 = v270
		}
		v284 = v267 | base.I64_extend_i32_u(v279)
		v285 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v172))))
		if base.Ui32((v285+int32(-65))&int32(255)) < base.Ui32(int32(26)) {
			v294 = v285 | int32(32)
		} else {
			v294 = v285
		}
		v300 = v284 | base.I64_extend_i32_u(v294)
	case 5:
		v217 = v185
		v218 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v172)+4)))
		if base.Ui32((v218+int32(-65))&int32(255)) < base.Ui32(int32(26)) {
			v227 = v218 | int32(32)
		} else {
			v227 = v218
		}
		v233 = base.I64_extend_i32_u(v227)<<(uint(int64(32))%64) | v217
		v234 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v172)+3)))
		v236 = v234 << (uint(int32(24)) % 32)
		if base.Ui32((v234+int32(-65))&int32(255)) < base.Ui32(int32(26)) {
			v245 = v236 | int32(536870912)
		} else {
			v245 = v236
		}
		v250 = v233 | base.I64_extend_i32_u(v245)
		v251 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v172)+2)))
		v253 = v251 << (uint(int32(16)) % 32)
		if base.Ui32((v251+int32(-65))&int32(255)) < base.Ui32(int32(26)) {
			v262 = v253 | int32(2097152)
		} else {
			v262 = v253
		}
		v267 = v250 | base.I64_extend_i32_u(v262)
		v268 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v172)+1)))
		v270 = v268 << (uint(int32(8)) % 32)
		if base.Ui32((v268+int32(-65))&int32(255)) < base.Ui32(int32(26)) {
			v279 = v270 | int32(8192)
		} else {
			v279 = v270
		}
		v284 = v267 | base.I64_extend_i32_u(v279)
		v285 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v172))))
		if base.Ui32((v285+int32(-65))&int32(255)) < base.Ui32(int32(26)) {
			v294 = v285 | int32(32)
		} else {
			v294 = v285
		}
		v300 = v284 | base.I64_extend_i32_u(v294)
	case 6:
		v201 = v185
		v202 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v172)+5)))
		if base.Ui32((v202+int32(-65))&int32(255)) < base.Ui32(int32(26)) {
			v211 = v202 | int32(32)
		} else {
			v211 = v202
		}
		v217 = base.I64_extend_i32_u(v211)<<(uint(int64(40))%64) | v201
		v218 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v172)+4)))
		if base.Ui32((v218+int32(-65))&int32(255)) < base.Ui32(int32(26)) {
			v227 = v218 | int32(32)
		} else {
			v227 = v218
		}
		v233 = base.I64_extend_i32_u(v227)<<(uint(int64(32))%64) | v217
		v234 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v172)+3)))
		v236 = v234 << (uint(int32(24)) % 32)
		if base.Ui32((v234+int32(-65))&int32(255)) < base.Ui32(int32(26)) {
			v245 = v236 | int32(536870912)
		} else {
			v245 = v236
		}
		v250 = v233 | base.I64_extend_i32_u(v245)
		v251 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v172)+2)))
		v253 = v251 << (uint(int32(16)) % 32)
		if base.Ui32((v251+int32(-65))&int32(255)) < base.Ui32(int32(26)) {
			v262 = v253 | int32(2097152)
		} else {
			v262 = v253
		}
		v267 = v250 | base.I64_extend_i32_u(v262)
		v268 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v172)+1)))
		v270 = v268 << (uint(int32(8)) % 32)
		if base.Ui32((v268+int32(-65))&int32(255)) < base.Ui32(int32(26)) {
			v279 = v270 | int32(8192)
		} else {
			v279 = v270
		}
		v284 = v267 | base.I64_extend_i32_u(v279)
		v285 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v172))))
		if base.Ui32((v285+int32(-65))&int32(255)) < base.Ui32(int32(26)) {
			v294 = v285 | int32(32)
		} else {
			v294 = v285
		}
		v300 = v284 | base.I64_extend_i32_u(v294)
	case 7:
		v186 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v172)+6)))
		if base.Ui32((v186+int32(-65))&int32(255)) < base.Ui32(int32(26)) {
			v195 = v186 | int32(32)
		} else {
			v195 = v186
		}
		v201 = base.I64_extend_i32_u(v195)<<(uint(int64(48))%64) | v185
		v202 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v172)+5)))
		if base.Ui32((v202+int32(-65))&int32(255)) < base.Ui32(int32(26)) {
			v211 = v202 | int32(32)
		} else {
			v211 = v202
		}
		v217 = base.I64_extend_i32_u(v211)<<(uint(int64(40))%64) | v201
		v218 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v172)+4)))
		if base.Ui32((v218+int32(-65))&int32(255)) < base.Ui32(int32(26)) {
			v227 = v218 | int32(32)
		} else {
			v227 = v218
		}
		v233 = base.I64_extend_i32_u(v227)<<(uint(int64(32))%64) | v217
		v234 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v172)+3)))
		v236 = v234 << (uint(int32(24)) % 32)
		if base.Ui32((v234+int32(-65))&int32(255)) < base.Ui32(int32(26)) {
			v245 = v236 | int32(536870912)
		} else {
			v245 = v236
		}
		v250 = v233 | base.I64_extend_i32_u(v245)
		v251 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v172)+2)))
		v253 = v251 << (uint(int32(16)) % 32)
		if base.Ui32((v251+int32(-65))&int32(255)) < base.Ui32(int32(26)) {
			v262 = v253 | int32(2097152)
		} else {
			v262 = v253
		}
		v267 = v250 | base.I64_extend_i32_u(v262)
		v268 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v172)+1)))
		v270 = v268 << (uint(int32(8)) % 32)
		if base.Ui32((v268+int32(-65))&int32(255)) < base.Ui32(int32(26)) {
			v279 = v270 | int32(8192)
		} else {
			v279 = v270
		}
		v284 = v267 | base.I64_extend_i32_u(v279)
		v285 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v172))))
		if base.Ui32((v285+int32(-65))&int32(255)) < base.Ui32(int32(26)) {
			v294 = v285 | int32(32)
		} else {
			v294 = v285
		}
		v300 = v284 | base.I64_extend_i32_u(v294)
	}
	v301 = v300 ^ v177
	v302 = int64(16)
	v304 = v301 + v175
	v305 = base.I64_rotl(v301, v302) ^ v304
	v306 = int64(21)
	v308 = v176 + v178
	v309 = int64(32)
	v311 = v305 + base.I64_rotl(v308, v309)
	v312 = base.I64_rotl(v305, v306) ^ v311
	v315 = int64(13)
	v317 = v308 ^ base.I64_rotl(v178, v315)
	v318 = v304 + v317
	v323 = base.I64_rotl(v318, v309) ^ int64(255) + v312
	v324 = base.I64_rotl(v312, v302) ^ v323
	v328 = int64(17)
	v330 = v318 ^ base.I64_rotl(v317, v328)
	v331 = v311 ^ v300 + v330
	v334 = base.I64_rotl(v331, v309) + v324
	v335 = base.I64_rotl(v324, v306) ^ v334
	v340 = v331 ^ base.I64_rotl(v330, v315)
	v341 = v340 + v323
	v344 = base.I64_rotl(v341, v309) + v335
	v350 = base.I64_rotl(v340, v328) ^ v341
	v354 = base.I64_rotl(v350, v315) ^ (v350 + v334)
	v358 = v354 + v344
	return base.I64_rotl(base.I64_rotl(v335, v302)^v344, v306) ^ base.I64_rotl(v354, v328) ^ base.I64_rotl(v358, v309) ^ v358
}
func F_sleep(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v21 int32
	_ = v21
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	*(*int64)(unsafe.Add(mBase, uint32(v6)+8)) = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v6))) = base.I64_extend_i32_u(l0)
	v12 = F_nanosleep(m, v6, v6)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		return int32(0)
	} else {
		v16 = *(*int32)(unsafe.Add(mBase, uint32(v6)))
		m.G0 = v6 + int32(16)
		if v12 != 0 {
			v21 = v16
		} else {
			v21 = int32(0)
		}
		return v21
	}
}
func F_smismemberCommand(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v36 int32
	_ = v36
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(v10)+4))
	v12 = F_lookupKeyRead(m, v9, v11)
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	return
L2:
	;
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	F_addReplyArrayLen(m, l0, v19+int32(-2))
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		goto L3
	} else {
		goto L8
	}
L3:
	;
	return
L4:
	;
	if v12 == int32(0) {
		goto L2
	} else {
		goto L5
	}
L5:
	;
	v17 = F_checkType(m, l0, v12, int32(2))
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		goto L3
	} else {
		goto L6
	}
L6:
	;
	if v17 != 0 {
		goto L1
	} else {
		goto L7
	}
L7:
	;
	goto L2
L8:
	;
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v24 < int32(3) {
		goto L1
	} else {
		goto L9
	}
L9:
	;
	v36 = int32(2)
	goto L10
L10:
	;
	if v12 == int32(0) {
		goto L13
	} else {
		goto L14
	}
L11:
	;
	goto L1
L12:
	;
	v81 = *(*int32)(unsafe.Add(mBase, uint32(v80)))
	F_addReply(m, l0, v81)
	mBase = m.M
	v83 = m.ExcPending
	if v83 != 0 {
		goto L3
	} else {
		goto L23
	}
L13:
	;
	v80 = int32(_a579)
	goto L12
L14:
	;
	v45 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v45+v36<<(uint(int32(2))%32))))
	v50 = F_objectGetVal(m, v49)
	mBase = m.M
	v53 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v50+int32(-1)))))
	switch v53 & int32(7) {
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
		v70 = int32(0)
		goto L15
	}
L15:
	;
	v73 = F_setTypeIsMemberAux(m, v12, v50, v70, int64(0), int32(1))
	mBase = m.M
	v74 = m.ExcPending
	if v74 != 0 {
		goto L3
	} else {
		goto L21
	}
L16:
	;
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v50+int32(-17))))
	v70 = v69
	goto L15
L17:
	;
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v50+int32(-9))))
	v70 = v66
	goto L15
L18:
	;
	v63 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v50+int32(-5)))))
	v70 = v63
	goto L15
L19:
	;
	v60 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v50+int32(-3)))))
	v70 = v60
	goto L15
L20:
	;
	v70 = int32(base.Ui32(v53) >> (uint(int32(3)) % 32))
	goto L15
L21:
	;
	if v73 != 0 {
		v80 = int32(_a577)
		goto L12
	} else {
		goto L22
	}
L22:
	;
	goto L13
L23:
	;
	v85 = v36 + int32(1)
	v86 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v85 < v86 {
		v36 = v85
		goto L10
	} else {
		goto L24
	}
L24:
	;
	goto L11
}
func F_smoveCommand(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
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
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v138 int64
	_ = v138
	var v142 int32
	_ = v142
	var v144 int32
	_ = v144
	var v147 int32
	_ = v147
	var v154 int32
	_ = v154
	var v157 int32
	_ = v157
	var v160 int32
	_ = v160
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v171 int32
	_ = v171
	var v173 int64
	_ = v173
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v181 int32
	_ = v181
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v189 int32
	_ = v189
	var v191 int32
	_ = v191
	var v193 int32
	_ = v193
	v10 = m.G0
	v12 = v10 - int32(16)
	m.G0 = v12
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v16 = *(*int32)(unsafe.Add(mBase, uint32(v15)+4))
	v17 = F_lookupKeyWrite(m, v14, v16)
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		return
	} else {
		v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
		v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
		v21 = *(*int32)(unsafe.Add(mBase, uint32(v20)+8))
		v22 = F_lookupKeyWrite(m, v19, v21)
		mBase = m.M
		v23 = m.ExcPending
		if v23 != 0 {
			return
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v12)+12)) = v22
			if v17 != 0 {
				v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
				v30 = *(*int32)(unsafe.Add(mBase, uint32(v29)+12))
				v32 = F_checkType(m, l0, v17, int32(2))
				mBase = m.M
				v33 = m.ExcPending
				if v33 != 0 {
					return
				} else {
					if v32 != 0 {
						m.G0 = v12 + int32(16)
						return
					} else {
						v35 = F_checkType(m, l0, v22, int32(2))
						mBase = m.M
						v36 = m.ExcPending
						if v36 != 0 {
							return
						} else {
							if v35 != 0 {
								m.G0 = v12 + int32(16)
								return
							} else {
								v37 = F_objectGetVal(m, v30)
								mBase = m.M
								v40 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v37+int32(-1)))))
								v42 = v40 & int32(7)
								if v17 != v22 {
									switch v42 {
									case 0:
										v87 = int32(base.Ui32(v40) >> (uint(int32(3)) % 32))
									case 1:
										v77 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v37+int32(-3)))))
										v87 = v77
									case 2:
										v80 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v37+int32(-5)))))
										v87 = v80
									case 3:
										v83 = *(*int32)(unsafe.Add(mBase, uint32(v37+int32(-9))))
										v87 = v83
									case 4:
										v86 = *(*int32)(unsafe.Add(mBase, uint32(v37+int32(-17))))
										v87 = v86
									default:
										v87 = int32(0)
									}
									v90 = F_setTypeRemoveAux(m, v17, v37, v87, int64(0), int32(1))
									mBase = m.M
									v91 = m.ExcPending
									if v91 != 0 {
										return
									} else {
										if v90 != 0 {
											v98 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
											v99 = *(*int32)(unsafe.Add(mBase, uint32(v98)+4))
											v100 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
											v101 = *(*int32)(unsafe.Add(mBase, uint32(v100)+28))
											F_notifyKeyspaceEvent(m, int32(32), int32(_a1686), v99, v101)
											mBase = m.M
											v103 = m.ExcPending
											if v103 != 0 {
												return
											} else {
												v104 = F_setTypeSize(m, v17)
												mBase = m.M
												v105 = m.ExcPending
												if v105 != 0 {
													return
												} else {
													if v104 != 0 {
														if v22 != 0 {
															v131 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
															v132 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
															v133 = *(*int32)(unsafe.Add(mBase, uint32(v132)+4))
															F_signalModifiedKey(m, l0, v131, v133)
															mBase = m.M
															v135 = m.ExcPending
															if v135 != 0 {
																return
															} else {
																v136 = int32(_a44)
																v138 = *(*int64)(unsafe.Add(mBase, _consts[83]))
																*(*int64)(unsafe.Add(mBase, _consts[83])) = v138 + int64(1)
																v142 = *(*int32)(unsafe.Add(mBase, uint32(v12)+12))
																v144 = F_objectGetVal(m, v30)
																mBase = m.M
																v147 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v144+int32(-1)))))
																switch v147 & int32(7) {
																case 0:
																	v164 = int32(base.Ui32(v147) >> (uint(int32(3)) % 32))
																case 1:
																	v154 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v144+int32(-3)))))
																	v164 = v154
																case 2:
																	v157 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v144+int32(-5)))))
																	v164 = v157
																case 3:
																	v160 = *(*int32)(unsafe.Add(mBase, uint32(v144+int32(-9))))
																	v164 = v160
																case 4:
																	v163 = *(*int32)(unsafe.Add(mBase, uint32(v144+int32(-17))))
																	v164 = v163
																default:
																	v164 = int32(0)
																}
																v167 = F_setTypeAddAux(m, v142, v144, v164, int64(0), int32(1))
																mBase = m.M
																v168 = m.ExcPending
																if v168 != 0 {
																	return
																} else {
																	if v167 == int32(0) {
																		v191 = *(*int32)(unsafe.Add(mBase, _consts[347]))
																		F_addReply(m, l0, v191)
																		mBase = m.M
																		v193 = m.ExcPending
																		if v193 != 0 {
																			return
																		} else {
																			m.G0 = v12 + int32(16)
																			return
																		}
																	} else {
																		v171 = int32(_a44)
																		v173 = *(*int64)(unsafe.Add(mBase, _consts[83]))
																		*(*int64)(unsafe.Add(mBase, _consts[83])) = v173 + int64(1)
																		v177 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
																		v178 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
																		v179 = *(*int32)(unsafe.Add(mBase, uint32(v178)+8))
																		F_signalModifiedKey(m, l0, v177, v179)
																		mBase = m.M
																		v181 = m.ExcPending
																		if v181 != 0 {
																			return
																		} else {
																			v184 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
																			v185 = *(*int32)(unsafe.Add(mBase, uint32(v184)+8))
																			v186 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
																			v187 = *(*int32)(unsafe.Add(mBase, uint32(v186)+28))
																			F_notifyKeyspaceEvent(m, int32(32), int32(_a1687), v185, v187)
																			mBase = m.M
																			v189 = m.ExcPending
																			if v189 != 0 {
																				return
																			} else {
																				v191 = *(*int32)(unsafe.Add(mBase, _consts[347]))
																				F_addReply(m, l0, v191)
																				mBase = m.M
																				v193 = m.ExcPending
																				if v193 != 0 {
																					return
																				} else {
																					m.G0 = v12 + int32(16)
																					return
																				}
																			}
																		}
																	}
																}
															}
														} else {
															v119 = F_objectGetVal(m, v30)
															mBase = m.M
															v121 = F_setTypeCreate(m, v119, int32(1))
															mBase = m.M
															v122 = m.ExcPending
															if v122 != 0 {
																return
															} else {
																*(*int32)(unsafe.Add(mBase, uint32(v12)+12)) = v121
																v124 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
																v125 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
																v126 = *(*int32)(unsafe.Add(mBase, uint32(v125)+8))
																F_dbAdd(m, v124, v126, v12+int32(12))
																mBase = m.M
																v130 = m.ExcPending
																if v130 != 0 {
																	return
																} else {
																	v131 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
																	v132 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
																	v133 = *(*int32)(unsafe.Add(mBase, uint32(v132)+4))
																	F_signalModifiedKey(m, l0, v131, v133)
																	mBase = m.M
																	v135 = m.ExcPending
																	if v135 != 0 {
																		return
																	} else {
																		v136 = int32(_a44)
																		v138 = *(*int64)(unsafe.Add(mBase, _consts[83]))
																		*(*int64)(unsafe.Add(mBase, _consts[83])) = v138 + int64(1)
																		v142 = *(*int32)(unsafe.Add(mBase, uint32(v12)+12))
																		v144 = F_objectGetVal(m, v30)
																		mBase = m.M
																		v147 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v144+int32(-1)))))
																		switch v147 & int32(7) {
																		case 0:
																			v164 = int32(base.Ui32(v147) >> (uint(int32(3)) % 32))
																		case 1:
																			v154 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v144+int32(-3)))))
																			v164 = v154
																		case 2:
																			v157 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v144+int32(-5)))))
																			v164 = v157
																		case 3:
																			v160 = *(*int32)(unsafe.Add(mBase, uint32(v144+int32(-9))))
																			v164 = v160
																		case 4:
																			v163 = *(*int32)(unsafe.Add(mBase, uint32(v144+int32(-17))))
																			v164 = v163
																		default:
																			v164 = int32(0)
																		}
																		v167 = F_setTypeAddAux(m, v142, v144, v164, int64(0), int32(1))
																		mBase = m.M
																		v168 = m.ExcPending
																		if v168 != 0 {
																			return
																		} else {
																			if v167 == int32(0) {
																				v191 = *(*int32)(unsafe.Add(mBase, _consts[347]))
																				F_addReply(m, l0, v191)
																				mBase = m.M
																				v193 = m.ExcPending
																				if v193 != 0 {
																					return
																				} else {
																					m.G0 = v12 + int32(16)
																					return
																				}
																			} else {
																				v171 = int32(_a44)
																				v173 = *(*int64)(unsafe.Add(mBase, _consts[83]))
																				*(*int64)(unsafe.Add(mBase, _consts[83])) = v173 + int64(1)
																				v177 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
																				v178 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
																				v179 = *(*int32)(unsafe.Add(mBase, uint32(v178)+8))
																				F_signalModifiedKey(m, l0, v177, v179)
																				mBase = m.M
																				v181 = m.ExcPending
																				if v181 != 0 {
																					return
																				} else {
																					v184 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
																					v185 = *(*int32)(unsafe.Add(mBase, uint32(v184)+8))
																					v186 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
																					v187 = *(*int32)(unsafe.Add(mBase, uint32(v186)+28))
																					F_notifyKeyspaceEvent(m, int32(32), int32(_a1687), v185, v187)
																					mBase = m.M
																					v189 = m.ExcPending
																					if v189 != 0 {
																						return
																					} else {
																						v191 = *(*int32)(unsafe.Add(mBase, _consts[347]))
																						F_addReply(m, l0, v191)
																						mBase = m.M
																						v193 = m.ExcPending
																						if v193 != 0 {
																							return
																						} else {
																							m.G0 = v12 + int32(16)
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
													} else {
														v106 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
														v107 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
														v108 = *(*int32)(unsafe.Add(mBase, uint32(v107)+4))
														v109 = F_dbDelete(m, v106, v108)
														mBase = m.M
														v110 = m.ExcPending
														if v110 != 0 {
															return
														} else {
															v113 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
															v114 = *(*int32)(unsafe.Add(mBase, uint32(v113)+4))
															v115 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
															v116 = *(*int32)(unsafe.Add(mBase, uint32(v115)+28))
															F_notifyKeyspaceEvent(m, int32(4), int32(_a213), v114, v116)
															mBase = m.M
															v118 = m.ExcPending
															if v118 != 0 {
																return
															} else {
																if v22 != 0 {
																	v131 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
																	v132 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
																	v133 = *(*int32)(unsafe.Add(mBase, uint32(v132)+4))
																	F_signalModifiedKey(m, l0, v131, v133)
																	mBase = m.M
																	v135 = m.ExcPending
																	if v135 != 0 {
																		return
																	} else {
																		v136 = int32(_a44)
																		v138 = *(*int64)(unsafe.Add(mBase, _consts[83]))
																		*(*int64)(unsafe.Add(mBase, _consts[83])) = v138 + int64(1)
																		v142 = *(*int32)(unsafe.Add(mBase, uint32(v12)+12))
																		v144 = F_objectGetVal(m, v30)
																		mBase = m.M
																		v147 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v144+int32(-1)))))
																		switch v147 & int32(7) {
																		case 0:
																			v164 = int32(base.Ui32(v147) >> (uint(int32(3)) % 32))
																		case 1:
																			v154 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v144+int32(-3)))))
																			v164 = v154
																		case 2:
																			v157 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v144+int32(-5)))))
																			v164 = v157
																		case 3:
																			v160 = *(*int32)(unsafe.Add(mBase, uint32(v144+int32(-9))))
																			v164 = v160
																		case 4:
																			v163 = *(*int32)(unsafe.Add(mBase, uint32(v144+int32(-17))))
																			v164 = v163
																		default:
																			v164 = int32(0)
																		}
																		v167 = F_setTypeAddAux(m, v142, v144, v164, int64(0), int32(1))
																		mBase = m.M
																		v168 = m.ExcPending
																		if v168 != 0 {
																			return
																		} else {
																			if v167 == int32(0) {
																				v191 = *(*int32)(unsafe.Add(mBase, _consts[347]))
																				F_addReply(m, l0, v191)
																				mBase = m.M
																				v193 = m.ExcPending
																				if v193 != 0 {
																					return
																				} else {
																					m.G0 = v12 + int32(16)
																					return
																				}
																			} else {
																				v171 = int32(_a44)
																				v173 = *(*int64)(unsafe.Add(mBase, _consts[83]))
																				*(*int64)(unsafe.Add(mBase, _consts[83])) = v173 + int64(1)
																				v177 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
																				v178 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
																				v179 = *(*int32)(unsafe.Add(mBase, uint32(v178)+8))
																				F_signalModifiedKey(m, l0, v177, v179)
																				mBase = m.M
																				v181 = m.ExcPending
																				if v181 != 0 {
																					return
																				} else {
																					v184 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
																					v185 = *(*int32)(unsafe.Add(mBase, uint32(v184)+8))
																					v186 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
																					v187 = *(*int32)(unsafe.Add(mBase, uint32(v186)+28))
																					F_notifyKeyspaceEvent(m, int32(32), int32(_a1687), v185, v187)
																					mBase = m.M
																					v189 = m.ExcPending
																					if v189 != 0 {
																						return
																					} else {
																						v191 = *(*int32)(unsafe.Add(mBase, _consts[347]))
																						F_addReply(m, l0, v191)
																						mBase = m.M
																						v193 = m.ExcPending
																						if v193 != 0 {
																							return
																						} else {
																							m.G0 = v12 + int32(16)
																							return
																						}
																					}
																				}
																			}
																		}
																	}
																} else {
																	v119 = F_objectGetVal(m, v30)
																	mBase = m.M
																	v121 = F_setTypeCreate(m, v119, int32(1))
																	mBase = m.M
																	v122 = m.ExcPending
																	if v122 != 0 {
																		return
																	} else {
																		*(*int32)(unsafe.Add(mBase, uint32(v12)+12)) = v121
																		v124 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
																		v125 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
																		v126 = *(*int32)(unsafe.Add(mBase, uint32(v125)+8))
																		F_dbAdd(m, v124, v126, v12+int32(12))
																		mBase = m.M
																		v130 = m.ExcPending
																		if v130 != 0 {
																			return
																		} else {
																			v131 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
																			v132 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
																			v133 = *(*int32)(unsafe.Add(mBase, uint32(v132)+4))
																			F_signalModifiedKey(m, l0, v131, v133)
																			mBase = m.M
																			v135 = m.ExcPending
																			if v135 != 0 {
																				return
																			} else {
																				v136 = int32(_a44)
																				v138 = *(*int64)(unsafe.Add(mBase, _consts[83]))
																				*(*int64)(unsafe.Add(mBase, _consts[83])) = v138 + int64(1)
																				v142 = *(*int32)(unsafe.Add(mBase, uint32(v12)+12))
																				v144 = F_objectGetVal(m, v30)
																				mBase = m.M
																				v147 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v144+int32(-1)))))
																				switch v147 & int32(7) {
																				case 0:
																					v164 = int32(base.Ui32(v147) >> (uint(int32(3)) % 32))
																				case 1:
																					v154 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v144+int32(-3)))))
																					v164 = v154
																				case 2:
																					v157 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v144+int32(-5)))))
																					v164 = v157
																				case 3:
																					v160 = *(*int32)(unsafe.Add(mBase, uint32(v144+int32(-9))))
																					v164 = v160
																				case 4:
																					v163 = *(*int32)(unsafe.Add(mBase, uint32(v144+int32(-17))))
																					v164 = v163
																				default:
																					v164 = int32(0)
																				}
																				v167 = F_setTypeAddAux(m, v142, v144, v164, int64(0), int32(1))
																				mBase = m.M
																				v168 = m.ExcPending
																				if v168 != 0 {
																					return
																				} else {
																					if v167 == int32(0) {
																						v191 = *(*int32)(unsafe.Add(mBase, _consts[347]))
																						F_addReply(m, l0, v191)
																						mBase = m.M
																						v193 = m.ExcPending
																						if v193 != 0 {
																							return
																						} else {
																							m.G0 = v12 + int32(16)
																							return
																						}
																					} else {
																						v171 = int32(_a44)
																						v173 = *(*int64)(unsafe.Add(mBase, _consts[83]))
																						*(*int64)(unsafe.Add(mBase, _consts[83])) = v173 + int64(1)
																						v177 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
																						v178 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
																						v179 = *(*int32)(unsafe.Add(mBase, uint32(v178)+8))
																						F_signalModifiedKey(m, l0, v177, v179)
																						mBase = m.M
																						v181 = m.ExcPending
																						if v181 != 0 {
																							return
																						} else {
																							v184 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
																							v185 = *(*int32)(unsafe.Add(mBase, uint32(v184)+8))
																							v186 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
																							v187 = *(*int32)(unsafe.Add(mBase, uint32(v186)+28))
																							F_notifyKeyspaceEvent(m, int32(32), int32(_a1687), v185, v187)
																							mBase = m.M
																							v189 = m.ExcPending
																							if v189 != 0 {
																								return
																							} else {
																								v191 = *(*int32)(unsafe.Add(mBase, _consts[347]))
																								F_addReply(m, l0, v191)
																								mBase = m.M
																								v193 = m.ExcPending
																								if v193 != 0 {
																									return
																								} else {
																									m.G0 = v12 + int32(16)
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
														}
													}
												}
											}
										} else {
											v93 = *(*int32)(unsafe.Add(mBase, _consts[85]))
											F_addReply(m, l0, v93)
											mBase = m.M
											v95 = m.ExcPending
											if v95 != 0 {
												return
											} else {
												m.G0 = v12 + int32(16)
												return
											}
										}
									}
								} else {
									switch v42 {
									case 0:
										v59 = int32(base.Ui32(v40) >> (uint(int32(3)) % 32))
									case 1:
										v49 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v37+int32(-3)))))
										v59 = v49
									case 2:
										v52 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v37+int32(-5)))))
										v59 = v52
									case 3:
										v55 = *(*int32)(unsafe.Add(mBase, uint32(v37+int32(-9))))
										v59 = v55
									case 4:
										v58 = *(*int32)(unsafe.Add(mBase, uint32(v37+int32(-17))))
										v59 = v58
									default:
										v59 = int32(0)
									}
									v65 = F_setTypeIsMemberAux(m, v17, v37, v59, int64(0), int32(1))
									mBase = m.M
									v66 = m.ExcPending
									if v66 != 0 {
										return
									} else {
										if v65 != 0 {
											v67 = int32(16)
										} else {
											v67 = int32(12)
										}
										v69 = *(*int32)(unsafe.Add(mBase, uint32(v67)+uint32(_consts[84])))
										F_addReply(m, l0, v69)
										mBase = m.M
										v71 = m.ExcPending
										if v71 != 0 {
											return
										} else {
											m.G0 = v12 + int32(16)
											return
										}
									}
								}
							}
						}
					}
				}
			} else {
				v26 = *(*int32)(unsafe.Add(mBase, _consts[85]))
				F_addReply(m, l0, v26)
				mBase = m.M
				v28 = m.ExcPending
				if v28 != 0 {
					return
				} else {
					m.G0 = v12 + int32(16)
					return
				}
			}
		}
	}
}
func F_sn_write(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+84))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(v9)))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(v9)+4))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v14 = v12 - v13
	if base.Ui32(v11) < base.Ui32(v14) {
		v16 = v11
	} else {
		v16 = v14
	}
	if v16 == int32(0) {
		v29 = v10
		v30 = v11
	} else {
		if v16 == int32(0) {
		} else {
			v21 = F__emscripten_memcpy_bulkmem(m, v10, v13, v16)
			mBase = m.M
		}
		v23 = *(*int32)(unsafe.Add(mBase, uint32(v9)))
		v24 = v23 + v16
		*(*int32)(unsafe.Add(mBase, uint32(v9))) = v24
		v26 = *(*int32)(unsafe.Add(mBase, uint32(v9)+4))
		v27 = v26 - v16
		*(*int32)(unsafe.Add(mBase, uint32(v9)+4)) = v27
		v29 = v24
		v30 = v27
	}
	if base.Ui32(v30) < base.Ui32(l2) {
		v32 = v30
	} else {
		v32 = l2
	}
	if v32 == int32(0) {
		v45 = v29
	} else {
		if v32 == int32(0) {
		} else {
			v37 = F__emscripten_memcpy_bulkmem(m, v29, l1, v32)
			mBase = m.M
		}
		v39 = *(*int32)(unsafe.Add(mBase, uint32(v9)))
		v40 = v39 + v32
		*(*int32)(unsafe.Add(mBase, uint32(v9))) = v40
		v42 = *(*int32)(unsafe.Add(mBase, uint32(v9)+4))
		*(*int32)(unsafe.Add(mBase, uint32(v9)+4)) = v42 - v32
		v45 = v40
	}
	v46 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v45))) = uint8(v46)
	v48 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v48
	return l2
}
func F_snprintf(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	*(*int32)(unsafe.Add(mBase, uint32(v8)+12)) = l3
	v11 = F_vsnprintf(m, l0, l1, l2, l3)
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		return int32(0)
	} else {
		m.G0 = v8 + int32(16)
		return v11
	}
}
func F_sort_2(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v7 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4))))
	v8 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3))))
	if v8 == int32(0) {
		v31 = v7
		v32 = v8
		goto L2
	} else {
		goto L3
	}
L1:
	;
	return v32 - v31&int32(255)
L2:
	;
	goto L1
L3:
	;
	if v8 != v7&int32(255) {
		v31 = v7
		v32 = v8
		goto L2
	} else {
		goto L4
	}
L4:
	;
	v14 = v3
	v15 = v4
	goto L5
L5:
	;
	v18 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15)+1)))
	v19 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+1)))
	if v19 == int32(0) {
		v31 = v18
		v32 = v19
		goto L2
	} else {
		goto L7
	}
L6:
	;
	v31 = v18
	v32 = v19
	goto L2
L7:
	;
	v22 = int32(1)
	if v19 == v18&int32(255) {
		v14 = v14 + v22
		v15 = v15 + v22
		goto L5
	} else {
		goto L8
	}
L8:
	;
	goto L6
}
func F_sortroCommand(m *base.Module, l0 int32) {
	var v4 int32
	_ = v4
	F_sortCommandGeneric(m, l0, int32(1))
	v4 = m.ExcPending
	if v4 != 0 {
		return
	} else {
		return
	}
}
func F_spmcEnqueue(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+132))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
	v14 = v6 + (v7+int32(-1))&v10<<(uint(int32(6))%32)
	v15 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
	if v15 != v10 {
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(v14)+4)) = l1
		v18 = int32(1)
		*(*int32)(unsafe.Add(mBase, uint32(v14))) = v15 + v18
		v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
		*(*int32)(unsafe.Add(mBase, uint32(l0)+64)) = v21 + v18
	}
	return base.B2i32(v15 == v10)
}
func F_spmcInit(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v16 int32
	_ = v16
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v45 int32
	_ = v45
	if base.I32_popcnt(l1) != int32(1) {
		F__serverAssert(m, int32(_a1881), int32(_a1882), int32(111))
		mBase = m.M
		v45 = m.ExcPending
		if v45 != 0 {
			return
		} else {
			F_abort(m)
			mBase = m.M
			base.Wasm_trap_unreachable()
			for {
			}
		}
	} else {
		v10 = F_valkey_malloc_cache_aligned(m, l1<<(uint(int32(6))%32))
		mBase = m.M
		v11 = m.ExcPending
		if v11 != 0 {
			return
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(l0)+132)) = l1
			*(*int32)(unsafe.Add(mBase, uint32(l0)+128)) = v10
			*(*int64)(unsafe.Add(mBase, uint32(l0)+64)) = int64(0)
			v16 = int32(0)
			*(*int32)(unsafe.Add(mBase, uint32(l0))) = v16
			if l1 == v16 {
			} else {
				v23 = v10
				v24 = v16
				for {
					v26 = v24 << (uint(int32(6)) % 32)
					*(*int32)(unsafe.Add(mBase, uint32(v23+v26))) = v24
					v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
					*(*int32)(unsafe.Add(mBase, uint32(v29+v26)+4)) = int32(0)
					v34 = v24 + int32(1)
					v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)+132))
					if base.Ui32(v34) < base.Ui32(v35) {
						v23 = v29
						v24 = v34
						continue
					} else {
						break
					}
					break
				}
			}
			return
		}
	}
}
func F_sscanCommand(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int64
	_ = v28
	var v30 int32
	_ = v30
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v9 = *(*int32)(unsafe.Add(mBase, uint32(v8)+8))
	v10 = F_objectGetVal(m, v9)
	mBase = m.M
	v13 = F_parseScanCursorOrReply(m, l0, v10, v6+int32(8))
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		return
	} else {
		if v13 == int32(-1) {
			m.G0 = v6 + int32(16)
			return
		} else {
			v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
			v18 = *(*int32)(unsafe.Add(mBase, uint32(v17)+4))
			v20 = *(*int32)(unsafe.Add(mBase, _consts[931]))
			v21 = F_lookupKeyReadOrReply(m, l0, v18, v20)
			mBase = m.M
			v22 = m.ExcPending
			if v22 != 0 {
				return
			} else {
				if v21 == int32(0) {
					m.G0 = v6 + int32(16)
					return
				} else {
					v26 = F_checkType(m, l0, v21, int32(2))
					mBase = m.M
					v27 = m.ExcPending
					if v27 != 0 {
						return
					} else {
						if v26 != 0 {
							m.G0 = v6 + int32(16)
							return
						} else {
							v28 = *(*int64)(unsafe.Add(mBase, uint32(v6)+8))
							F_scanGenericCommand(m, l0, v21, v28)
							mBase = m.M
							v30 = m.ExcPending
							if v30 != 0 {
								return
							} else {
								m.G0 = v6 + int32(16)
								return
							}
						}
					}
				}
			}
		}
	}
}
func F_sscanf(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	*(*int32)(unsafe.Add(mBase, uint32(v7)+12)) = l2
	v10 = F_vsscanf(m, l0, l1, l2)
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return int32(0)
	} else {
		m.G0 = v7 + int32(16)
		return v10
	}
}
func F_startBgsaveForReplication(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v36 int32
	_ = v36
	var v41 int32
	_ = v41
	var v47 int32
	_ = v47
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v63 int64
	_ = v63
	var v68 int64
	_ = v68
	var v73 int64
	_ = v73
	var v78 int64
	_ = v78
	var v83 int64
	_ = v83
	var v88 int64
	_ = v88
	var v93 int64
	_ = v93
	var v96 int64
	_ = v96
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	var v106 int32
	_ = v106
	var v109 int32
	_ = v109
	var v112 int32
	_ = v112
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v119 int32
	_ = v119
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v125 int32
	_ = v125
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v133 int32
	_ = v133
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v139 int32
	_ = v139
	var v143 int32
	_ = v143
	var v148 int32
	_ = v148
	var v151 int32
	_ = v151
	var v156 int32
	_ = v156
	var v163 int32
	_ = v163
	var v168 int32
	_ = v168
	var v175 int32
	_ = v175
	var v178 int32
	_ = v178
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v185 int32
	_ = v185
	var v187 int32
	_ = v187
	var v189 int32
	_ = v189
	var v192 int32
	_ = v192
	var v198 int32
	_ = v198
	var v206 int32
	_ = v206
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v218 int32
	_ = v218
	var v223 int32
	_ = v223
	var v225 int32
	_ = v225
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v234 int32
	_ = v234
	var v236 int32
	_ = v236
	var v239 int32
	_ = v239
	var v245 int32
	_ = v245
	var v251 int32
	_ = v251
	var v253 int32
	_ = v253
	var v255 int32
	_ = v255
	var v257 int32
	_ = v257
	var v258 int32
	_ = v258
	var v263 int32
	_ = v263
	var v265 int32
	_ = v265
	var v268 int32
	_ = v268
	var v274 int32
	_ = v274
	var v281 int32
	_ = v281
	var v290 int32
	_ = v290
	var v291 int32
	_ = v291
	var v292 int32
	_ = v292
	var v295 int32
	_ = v295
	var v297 int32
	_ = v297
	var v301 int32
	_ = v301
	var v308 int32
	_ = v308
	var v309 int32
	_ = v309
	var v311 int32
	_ = v311
	var v313 int32
	_ = v313
	var v314 int64
	_ = v314
	var v321 int32
	_ = v321
	var v330 int32
	_ = v330
	var v331 int32
	_ = v331
	var v332 int32
	_ = v332
	var v335 int32
	_ = v335
	var v336 int32
	_ = v336
	var v337 int32
	_ = v337
	var v338 int32
	_ = v338
	var v341 int32
	_ = v341
	var v346 int32
	_ = v346
	var v348 int32
	_ = v348
	var v351 int32
	_ = v351
	var v357 int32
	_ = v357
	var v364 int32
	_ = v364
	v10 = m.G0
	v12 = v10 - int32(224)
	m.G0 = v12
	if l0&int32(1) != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	if v30 == int32(0) {
		goto L7
	} else {
		goto L8
	}
L2:
	;
	v22 = *(*int32)(unsafe.Add(mBase, _consts[655]))
	v24 = l1 & int32(3)
	v29 = base.B2i32(l2 != int32(80)) | base.B2i32(v22|v24 != int32(0))
	v30 = v24
	goto L1
L3:
	;
	v29 = int32(0)
	v30 = l1 & int32(3)
	goto L1
L4:
	;
	m.G0 = v12 + int32(224)
	return v364
L5:
	;
	if v29 != 0 {
		v364 = v253
		goto L4
	} else {
		goto L67
	}
L6:
	;
	F__serverAssert(m, int32(_a1195), int32(_a1190), int32(1006))
	mBase = m.M
	v251 = m.ExcPending
	if v251 != 0 {
		goto L18
	} else {
		goto L66
	}
L7:
	;
	v36 = *(*int32)(unsafe.Add(mBase, _consts[6]))
	if int32(2) < v36 {
		goto L10
	} else {
		goto L11
	}
L8:
	;
	if v29 == int32(0) {
		goto L6
	} else {
		goto L9
	}
L9:
	;
	goto L7
L10:
	;
	v58 = v12 + int32(24)
	v59 = int32(0)
	v63 = *(*int64)(unsafe.Add(mBase, _consts[361]))
	*(*int64)(unsafe.Add(mBase, uint32(v12+int32(80)))) = v63
	v68 = *(*int64)(unsafe.Add(mBase, _consts[362]))
	*(*int64)(unsafe.Add(mBase, uint32(v12+int32(72)))) = v68
	v73 = *(*int64)(unsafe.Add(mBase, _consts[363]))
	*(*int64)(unsafe.Add(mBase, uint32(v12+int32(64)))) = v73
	v78 = *(*int64)(unsafe.Add(mBase, _consts[364]))
	*(*int64)(unsafe.Add(mBase, uint32(v12+int32(56)))) = v78
	v83 = *(*int64)(unsafe.Add(mBase, _consts[365]))
	*(*int64)(unsafe.Add(mBase, uint32(v12+int32(48)))) = v83
	v88 = *(*int64)(unsafe.Add(mBase, _consts[366]))
	*(*int64)(unsafe.Add(mBase, uint32(v12+int32(40)))) = v88
	v93 = *(*int64)(unsafe.Add(mBase, _consts[367]))
	*(*int64)(unsafe.Add(mBase, uint32(v12+int32(32)))) = v93
	v96 = *(*int64)(unsafe.Add(mBase, _consts[368]))
	*(*int64)(unsafe.Add(mBase, uint32(v58))) = v96
	v99 = *(*int32)(unsafe.Add(mBase, _consts[130]))
	if v99 != 0 {
		goto L25
	} else {
		goto L26
	}
L11:
	;
	if v29 != 0 {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	v41 = int32(_a1196)
	goto L14
L13:
	;
	v41 = int32(_a1197)
	goto L14
L14:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = v41
	if l1&int32(4) != 0 {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	v47 = int32(_a1198)
	goto L17
L16:
	;
	v47 = int32(_a1199)
	goto L17
L17:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+20)) = v47
	F__serverLog(m, int32(2), int32(_a1200), v12+int32(16))
	mBase = m.M
	v56 = m.ExcPending
	if v56 != 0 {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	return int32(0)
L19:
	;
	goto L10
L20:
	;
	v178 = *(*int32)(unsafe.Add(mBase, _consts[158]))
	v180 = v12 + int32(88)
	v181 = *(*int32)(unsafe.Add(mBase, uint32(v178)))
	*(*int32)(unsafe.Add(mBase, uint32(v180)+4)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v180))) = v181
	goto L51
L21:
	;
	v168 = *(*int32)(unsafe.Add(mBase, _consts[6]))
	if int32(3) < v168 {
		goto L20
	} else {
		goto L49
	}
L22:
	;
	if v137 != int32(-1) {
		v253 = v137
		goto L5
	} else {
		goto L48
	}
L23:
	;
	v156 = *(*int32)(unsafe.Add(mBase, _consts[6]))
	if int32(3) < v156 {
		goto L20
	} else {
		goto L46
	}
L24:
	;
	if v125 == int32(0) {
		goto L23
	} else {
		goto L35
	}
L25:
	;
	v112 = *(*int32)(unsafe.Add(mBase, _consts[188]))
	if v112 == int32(0) {
		goto L31
	} else {
		goto L32
	}
L26:
	;
	v101 = *(*int32)(unsafe.Add(mBase, _consts[370]))
	if v101 == int32(0) {
		goto L25
	} else {
		goto L27
	}
L27:
	;
	v106 = *(*int32)(unsafe.Add(mBase, _consts[371]))
	if v106 == int32(-1) {
		goto L28
	} else {
		goto L29
	}
L28:
	;
	v109 = int32(0)
	goto L30
L29:
	;
	v109 = v106
	goto L30
L30:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v58))) = v109
	v125 = v58
	goto L24
L31:
	;
	v119 = *(*int32)(unsafe.Add(mBase, _consts[189]))
	if v119 != 0 {
		goto L33
	} else {
		goto L34
	}
L32:
	;
	v115 = *(*int32)(unsafe.Add(mBase, uint32(v112)+96))
	v116 = *(*int32)(unsafe.Add(mBase, uint32(v115)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v58))) = v116
	v125 = v58
	goto L24
L33:
	;
	v121 = *(*int32)(unsafe.Add(mBase, uint32(v119)+96))
	v122 = *(*int32)(unsafe.Add(mBase, uint32(v121)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v58))) = v122
	v125 = v58
	goto L24
L34:
	;
	v125 = int32(0)
	goto L24
L35:
	;
	if v29 == int32(0) {
		goto L37
	} else {
		goto L38
	}
L36:
	;
	v139 = *(*int32)(unsafe.Add(mBase, _consts[397]))
	if v139 == int32(0) {
		goto L41
	} else {
		goto L42
	}
L37:
	;
	v133 = *(*int32)(unsafe.Add(mBase, _consts[369]))
	v135 = F_rdbSaveBackground(m, l1, v133, v125, int32(18))
	mBase = m.M
	v136 = m.ExcPending
	if v136 != 0 {
		goto L18
	} else {
		goto L40
	}
L38:
	;
	v130 = F_rdbSaveToReplicasSockets(m, l1, l2, v125)
	mBase = m.M
	v131 = m.ExcPending
	if v131 != 0 {
		goto L18
	} else {
		goto L39
	}
L39:
	;
	v137 = v130
	goto L36
L40:
	;
	v137 = v135
	goto L36
L41:
	;
	if v29|base.B2i32(v137 != int32(0)) != 0 {
		goto L22
	} else {
		goto L44
	}
L42:
	;
	F_debugPauseProcess(m)
	mBase = m.M
	v143 = m.ExcPending
	if v143 != 0 {
		goto L18
	} else {
		goto L43
	}
L43:
	;
	goto L41
L44:
	;
	v148 = *(*int32)(unsafe.Add(mBase, _consts[656]))
	if v148 == int32(0) {
		goto L22
	} else {
		goto L45
	}
L45:
	;
	v151 = int32(0)
	*(*int32)(unsafe.Add(mBase, _consts[657])) = int32(1)
	v253 = v151
	goto L5
L46:
	;
	F__serverLog(m, int32(3), int32(_a1201), int32(0))
	mBase = m.M
	v163 = m.ExcPending
	if v163 != 0 {
		goto L18
	} else {
		goto L47
	}
L47:
	;
	goto L21
L48:
	;
	goto L21
L49:
	;
	F__serverLog(m, int32(3), int32(_a1202), int32(0))
	mBase = m.M
	v175 = m.ExcPending
	if v175 != 0 {
		goto L18
	} else {
		goto L50
	}
L50:
	;
	goto L20
L51:
	;
	v185 = int32(-1)
	v187 = v12 + int32(88)
	v189 = *(*int32)(unsafe.Add(mBase, uint32(v187)))
	if v189 == int32(0) {
		goto L53
	} else {
		goto L54
	}
L52:
	;
	if v189 == int32(0) {
		v364 = v185
		goto L4
	} else {
		goto L55
	}
L53:
	;
	goto L52
L54:
	;
	v192 = *(*int32)(unsafe.Add(mBase, uint32(v187)+4))
	v198 = *(*int32)(unsafe.Add(mBase, uint32(v189+base.B2i32(v192 == int32(0))<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v187))) = v198
	goto L53
L55:
	;
	v206 = v189
	goto L56
L56:
	;
	v211 = *(*int32)(unsafe.Add(mBase, uint32(v206)+8))
	v212 = *(*int32)(unsafe.Add(mBase, uint32(v211)+104))
	v213 = *(*int32)(unsafe.Add(mBase, uint32(v212)))
	if v213 != int32(6) {
		goto L58
	} else {
		goto L59
	}
L58:
	;
	v234 = v12 + int32(88)
	v236 = *(*int32)(unsafe.Add(mBase, uint32(v234)))
	if v236 == int32(0) {
		goto L63
	} else {
		goto L64
	}
L59:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v212))) = int32(0)
	v218 = *(*int32)(unsafe.Add(mBase, uint32(v211)+200))
	*(*int32)(unsafe.Add(mBase, uint32(v211)+200)) = v218 & int32(-3)
	v223 = *(*int32)(unsafe.Add(mBase, _consts[158]))
	F_listDelNode(m, v223, v206)
	mBase = m.M
	v225 = m.ExcPending
	if v225 != 0 {
		goto L18
	} else {
		goto L60
	}
L60:
	;
	F_addReplyError(m, v211, int32(_a1203))
	mBase = m.M
	v228 = m.ExcPending
	if v228 != 0 {
		goto L18
	} else {
		goto L61
	}
L61:
	;
	v229 = *(*int32)(unsafe.Add(mBase, uint32(v211)+200))
	*(*int32)(unsafe.Add(mBase, uint32(v211)+200)) = v229 | int32(64)
	goto L58
L62:
	;
	if v236 != 0 {
		v206 = v236
		goto L56
	} else {
		goto L65
	}
L63:
	;
	goto L62
L64:
	;
	v239 = *(*int32)(unsafe.Add(mBase, uint32(v234)+4))
	v245 = *(*int32)(unsafe.Add(mBase, uint32(v236+base.B2i32(v239 == int32(0))<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v234))) = v245
	goto L63
L65:
	;
	v364 = v185
	goto L4
L66:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L67:
	;
	v255 = *(*int32)(unsafe.Add(mBase, _consts[158]))
	v257 = v12 + int32(88)
	v258 = *(*int32)(unsafe.Add(mBase, uint32(v255)))
	*(*int32)(unsafe.Add(mBase, uint32(v257)+4)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v257))) = v258
	goto L68
L68:
	;
	v263 = v12 + int32(88)
	v265 = *(*int32)(unsafe.Add(mBase, uint32(v263)))
	if v265 == int32(0) {
		goto L70
	} else {
		goto L71
	}
L69:
	;
	if v265 == int32(0) {
		v364 = v253
		goto L4
	} else {
		goto L72
	}
L70:
	;
	goto L69
L71:
	;
	v268 = *(*int32)(unsafe.Add(mBase, uint32(v263)+4))
	v274 = *(*int32)(unsafe.Add(mBase, uint32(v265+base.B2i32(v268 == int32(0))<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v263))) = v274
	goto L70
L72:
	;
	v281 = v265
	goto L73
L73:
	;
	v290 = *(*int32)(unsafe.Add(mBase, uint32(v281)+8))
	v291 = *(*int32)(unsafe.Add(mBase, uint32(v290)+104))
	v292 = *(*int32)(unsafe.Add(mBase, uint32(v291)))
	if v292 != int32(6) {
		goto L75
	} else {
		goto L76
	}
L74:
	;
	v364 = v253
	goto L4
L75:
	;
	v346 = v12 + int32(88)
	v348 = *(*int32)(unsafe.Add(mBase, uint32(v346)))
	if v348 == int32(0) {
		goto L93
	} else {
		goto L94
	}
L76:
	;
	v295 = int32(*(*int16)(unsafe.Add(mBase, uint32(v291)+162)))
	if l1 != v295 {
		goto L75
	} else {
		goto L77
	}
L77:
	;
	v297 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v291)+160)))
	if v297&int32(1) != 0 {
		goto L79
	} else {
		goto L80
	}
L78:
	;
	if v311 != l2 {
		goto L75
	} else {
		goto L86
	}
L79:
	;
	v301 = *(*int32)(unsafe.Add(mBase, uint32(v291)+156))
	if v301 <= int32(589823) {
		goto L83
	} else {
		goto L84
	}
L80:
	;
	v311 = int32(80)
	goto L78
L81:
	;
	v311 = int32(11)
	goto L78
L82:
	;
	v309 = *(*int32)(unsafe.Add(mBase, uint32(v308)))
	v311 = v309
	goto L78
L83:
	;
	if v301 < int32(459264) {
		goto L81
	} else {
		goto L85
	}
L84:
	;
	v308 = int32(_a1204)
	goto L82
L85:
	;
	v308 = int32(_a1205)
	goto L82
L86:
	;
	v313 = int32(_a44)
	v314 = *(*int64)(unsafe.Add(mBase, _consts[40]))
	*(*int32)(unsafe.Add(mBase, uint32(v291))) = int32(7)
	*(*int64)(unsafe.Add(mBase, uint32(v291)+96)) = v314
	*(*int32)(unsafe.Add(mBase, _consts[371])) = int32(-1)
	v321 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v290)+202)))
	if v321&int32(1) != 0 {
		goto L75
	} else {
		goto L87
	}
L87:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v12)+8)) = v314
	*(*int32)(unsafe.Add(mBase, uint32(v12))) = int32(_a946)
	v330 = F_snprintf(m, v12+int32(96), int32(128), int32(_a1206), v12)
	mBase = m.M
	v331 = m.ExcPending
	if v331 != 0 {
		goto L18
	} else {
		goto L88
	}
L88:
	;
	v332 = *(*int32)(unsafe.Add(mBase, uint32(v290)+8))
	v335 = *(*int32)(unsafe.Add(mBase, uint32(v332)))
	v336 = *(*int32)(unsafe.Add(mBase, uint32(v335)+68))
	v337 = m.T0[v336].(func(*base.Module, int32, int32, int32) int32)(m, v332, v12+int32(96), v330)
	mBase = m.M
	v338 = m.ExcPending
	if v338 != 0 {
		goto L18
	} else {
		goto L89
	}
L89:
	;
	if v330 == v337 {
		goto L75
	} else {
		goto L90
	}
L90:
	;
	F_freeClientAsync(m, v290)
	mBase = m.M
	v341 = m.ExcPending
	if v341 != 0 {
		goto L18
	} else {
		goto L91
	}
L91:
	;
	goto L75
L92:
	;
	if v348 != 0 {
		v281 = v348
		goto L73
	} else {
		goto L95
	}
L93:
	;
	goto L92
L94:
	;
	v351 = *(*int32)(unsafe.Add(mBase, uint32(v346)+4))
	v357 = *(*int32)(unsafe.Add(mBase, uint32(v348+base.B2i32(v351 == int32(0))<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v346))) = v357
	goto L93
L95:
	;
	goto L74
}
func F_startLoadingFile(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v4 int64
	_ = v4
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v11 int64
	_ = v11
	var v36 int32
	_ = v36
	var v40 int32
	_ = v40
	var v41 int64
	_ = v41
	var v45 int64
	_ = v45
	var v49 int64
	_ = v49
	var v52 int32
	_ = v52
	var v60 int64
	_ = v60
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	v4 = int64(0)
	v5 = int32(_a44)
	v6 = int32(1)
	*(*int32)(unsafe.Add(mBase, _consts[131])) = v6
	v8 = int32(0)
	*(*int32)(unsafe.Add(mBase, _consts[74])) = l1
	v11 = F___time(m, v8)
	mBase = m.M
	*(*int64)(unsafe.Add(mBase, _consts[70])) = v4
	*(*int64)(unsafe.Add(mBase, _consts[636])) = v11
	*(*int64)(unsafe.Add(mBase, _consts[637])) = v4
	*(*int64)(unsafe.Add(mBase, _consts[638])) = v4
	*(*int64)(unsafe.Add(mBase, _consts[639])) = base.I64_extend_i32_u(l0)
	*(*int64)(unsafe.Add(mBase, _consts[640])) = v4
	v36 = *(*int32)(unsafe.Add(mBase, _consts[527]))
	*(*int32)(unsafe.Add(mBase, _consts[527])) = v36 + v6
	if v36 != 0 {
	} else {
		v40 = int32(0)
		v41 = F_ustime(m)
		mBase = m.M
		*(*int64)(unsafe.Add(mBase, _consts[96])) = v41
		v45 = base.I64_div_s(v41, int64(1000))
		*(*int64)(unsafe.Add(mBase, _consts[35])) = v45
		v49 = base.I64_div_s(v41, int64(1000000))
		*(*int64)(unsafe.Add(mBase, _consts[47])) = v49
		v52 = *(*int32)(unsafe.Add(mBase, _consts[97]))
		F_lrulfu_updateClockAndPolicy(m, v45, int32(base.Ui32(v52&int32(2))>>(uint(int32(1))%32)))
		mBase = m.M
		v60 = *(*int64)(unsafe.Add(mBase, _consts[35]))
		*(*int64)(unsafe.Add(mBase, _consts[528])) = v60
	}
	F_clusterCleanSlotImportsBeforeLoad(m)
	mBase = m.M
	v65 = m.ExcPending
	if v65 != 0 {
		return
	} else {
		v67 = int32(1)
		if l2&v67 != 0 {
			v72 = v67
		} else {
			v72 = l2 & int32(2)
		}
		F_moduleFireServerEvent(m, int64(3), v72, int32(0))
		mBase = m.M
		v75 = m.ExcPending
		if v75 != 0 {
			return
		} else {
			return
		}
	}
}
func F_startSaving(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	v6 = F___syscall_getpid(m)
	mBase = m.M
	v8 = *(*int32)(unsafe.Add(mBase, _consts[635]))
	v9 = base.B2i32(v6 == v8)
	if v6 == v8 {
		v10 = int32(5)
	} else {
		v10 = int32(1)
	}
	v11 = int32(1)
	if l0&v11 != 0 {
		v15 = v10
	} else {
		v15 = v9 << (uint(v11) % 32)
	}
	F_moduleFireServerEvent(m, int64(1), v15, int32(0))
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		return
	} else {
		return
	}
}
func F_stepCommandHandler(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	v4 = m.G6
	*(*int32)(unsafe.Add(mBase, uint32(v4)+264)) = int32(1)
	return int32(0)
}
func F_stopAppendOnly(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v11 int32
	_ = v11
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v36 int64
	_ = v36
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v55 int32
	_ = v55
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v64 int64
	_ = v64
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v94 int32
	_ = v94
	v2 = m.G0
	v4 = v2 - int32(16)
	m.G0 = v4
	v7 = *(*int32)(unsafe.Add(mBase, _consts[28]))
	if v7 == int32(0) {
		F__serverAssert(m, int32(_a111), int32(_a85), int32(934))
		mBase = m.M
		v94 = m.ExcPending
		if v94 != 0 {
			return
		} else {
			F_abort(m)
			mBase = m.M
			base.Wasm_trap_unreachable()
			for {
			}
		}
	} else {
		v11 = *(*int32)(unsafe.Add(mBase, _consts[25]))
		if v11 == int32(-1) {
			v41 = int32(_a44)
			v42 = int32(0)
			*(*int32)(unsafe.Add(mBase, _consts[28])) = v42
			*(*int64)(unsafe.Add(mBase, _consts[25])) = int64(-1)
			v48 = *(*int32)(unsafe.Add(mBase, _consts[29]))
			if v48 == v42 {
				v63 = int32(_a44)
				v64 = int64(0)
				*(*int64)(unsafe.Add(mBase, _consts[30])) = v64
				*(*int64)(unsafe.Add(mBase, _consts[31])) = int64(-1)
				*(*int64)(unsafe.Add(mBase, _consts[32])) = v64
				*(*int64)(unsafe.Add(mBase, _consts[33])) = v64
				F_killAppendOnlyChild(m)
				mBase = m.M
				v78 = m.ExcPending
				if v78 != 0 {
					return
				} else {
					v80 = *(*int32)(unsafe.Add(mBase, _consts[34]))
					F_sdsfree(m, v80)
					mBase = m.M
					v82 = m.ExcPending
					if v82 != 0 {
						return
					} else {
						v84 = F_sdsempty(m)
						mBase = m.M
						v85 = m.ExcPending
						if v85 != 0 {
							return
						} else {
							*(*int32)(unsafe.Add(mBase, _consts[34])) = v84
							m.G0 = v4 + int32(16)
							return
						}
					}
				}
			} else {
				v51 = int32(_a44)
				*(*int32)(unsafe.Add(mBase, _consts[29])) = int32(0)
				v55 = *(*int32)(unsafe.Add(mBase, _consts[6]))
				if int32(2) < v55 {
					v63 = int32(_a44)
					v64 = int64(0)
					*(*int64)(unsafe.Add(mBase, _consts[30])) = v64
					*(*int64)(unsafe.Add(mBase, _consts[31])) = int64(-1)
					*(*int64)(unsafe.Add(mBase, _consts[32])) = v64
					*(*int64)(unsafe.Add(mBase, _consts[33])) = v64
					F_killAppendOnlyChild(m)
					mBase = m.M
					v78 = m.ExcPending
					if v78 != 0 {
						return
					} else {
						v80 = *(*int32)(unsafe.Add(mBase, _consts[34]))
						F_sdsfree(m, v80)
						mBase = m.M
						v82 = m.ExcPending
						if v82 != 0 {
							return
						} else {
							v84 = F_sdsempty(m)
							mBase = m.M
							v85 = m.ExcPending
							if v85 != 0 {
								return
							} else {
								*(*int32)(unsafe.Add(mBase, _consts[34])) = v84
								m.G0 = v4 + int32(16)
								return
							}
						}
					}
				} else {
					F__serverLog(m, int32(2), int32(_a112), int32(0))
					mBase = m.M
					v62 = m.ExcPending
					if v62 != 0 {
						return
					} else {
						v63 = int32(_a44)
						v64 = int64(0)
						*(*int64)(unsafe.Add(mBase, _consts[30])) = v64
						*(*int64)(unsafe.Add(mBase, _consts[31])) = int64(-1)
						*(*int64)(unsafe.Add(mBase, _consts[32])) = v64
						*(*int64)(unsafe.Add(mBase, _consts[33])) = v64
						F_killAppendOnlyChild(m)
						mBase = m.M
						v78 = m.ExcPending
						if v78 != 0 {
							return
						} else {
							v80 = *(*int32)(unsafe.Add(mBase, _consts[34]))
							F_sdsfree(m, v80)
							mBase = m.M
							v82 = m.ExcPending
							if v82 != 0 {
								return
							} else {
								v84 = F_sdsempty(m)
								mBase = m.M
								v85 = m.ExcPending
								if v85 != 0 {
									return
								} else {
									*(*int32)(unsafe.Add(mBase, _consts[34])) = v84
									m.G0 = v4 + int32(16)
									return
								}
							}
						}
					}
				}
			}
		} else {
			F_flushAppendOnlyFile(m, int32(1))
			mBase = m.M
			v16 = m.ExcPending
			if v16 != 0 {
				return
			} else {
				v18 = *(*int32)(unsafe.Add(mBase, _consts[25]))
				v19 = F_fsync(m, v18)
				mBase = m.M
				if v19 != int32(-1) {
					v34 = int32(_a44)
					v36 = *(*int64)(unsafe.Add(mBase, _consts[35]))
					*(*int64)(unsafe.Add(mBase, _consts[36])) = v36
					v39 = *(*int32)(unsafe.Add(mBase, _consts[25]))
					v40 = F_close(m, v39)
					mBase = m.M
					v41 = int32(_a44)
					v42 = int32(0)
					*(*int32)(unsafe.Add(mBase, _consts[28])) = v42
					*(*int64)(unsafe.Add(mBase, _consts[25])) = int64(-1)
					v48 = *(*int32)(unsafe.Add(mBase, _consts[29]))
					if v48 == v42 {
						v63 = int32(_a44)
						v64 = int64(0)
						*(*int64)(unsafe.Add(mBase, _consts[30])) = v64
						*(*int64)(unsafe.Add(mBase, _consts[31])) = int64(-1)
						*(*int64)(unsafe.Add(mBase, _consts[32])) = v64
						*(*int64)(unsafe.Add(mBase, _consts[33])) = v64
						F_killAppendOnlyChild(m)
						mBase = m.M
						v78 = m.ExcPending
						if v78 != 0 {
							return
						} else {
							v80 = *(*int32)(unsafe.Add(mBase, _consts[34]))
							F_sdsfree(m, v80)
							mBase = m.M
							v82 = m.ExcPending
							if v82 != 0 {
								return
							} else {
								v84 = F_sdsempty(m)
								mBase = m.M
								v85 = m.ExcPending
								if v85 != 0 {
									return
								} else {
									*(*int32)(unsafe.Add(mBase, _consts[34])) = v84
									m.G0 = v4 + int32(16)
									return
								}
							}
						}
					} else {
						v51 = int32(_a44)
						*(*int32)(unsafe.Add(mBase, _consts[29])) = int32(0)
						v55 = *(*int32)(unsafe.Add(mBase, _consts[6]))
						if int32(2) < v55 {
							v63 = int32(_a44)
							v64 = int64(0)
							*(*int64)(unsafe.Add(mBase, _consts[30])) = v64
							*(*int64)(unsafe.Add(mBase, _consts[31])) = int64(-1)
							*(*int64)(unsafe.Add(mBase, _consts[32])) = v64
							*(*int64)(unsafe.Add(mBase, _consts[33])) = v64
							F_killAppendOnlyChild(m)
							mBase = m.M
							v78 = m.ExcPending
							if v78 != 0 {
								return
							} else {
								v80 = *(*int32)(unsafe.Add(mBase, _consts[34]))
								F_sdsfree(m, v80)
								mBase = m.M
								v82 = m.ExcPending
								if v82 != 0 {
									return
								} else {
									v84 = F_sdsempty(m)
									mBase = m.M
									v85 = m.ExcPending
									if v85 != 0 {
										return
									} else {
										*(*int32)(unsafe.Add(mBase, _consts[34])) = v84
										m.G0 = v4 + int32(16)
										return
									}
								}
							}
						} else {
							F__serverLog(m, int32(2), int32(_a112), int32(0))
							mBase = m.M
							v62 = m.ExcPending
							if v62 != 0 {
								return
							} else {
								v63 = int32(_a44)
								v64 = int64(0)
								*(*int64)(unsafe.Add(mBase, _consts[30])) = v64
								*(*int64)(unsafe.Add(mBase, _consts[31])) = int64(-1)
								*(*int64)(unsafe.Add(mBase, _consts[32])) = v64
								*(*int64)(unsafe.Add(mBase, _consts[33])) = v64
								F_killAppendOnlyChild(m)
								mBase = m.M
								v78 = m.ExcPending
								if v78 != 0 {
									return
								} else {
									v80 = *(*int32)(unsafe.Add(mBase, _consts[34]))
									F_sdsfree(m, v80)
									mBase = m.M
									v82 = m.ExcPending
									if v82 != 0 {
										return
									} else {
										v84 = F_sdsempty(m)
										mBase = m.M
										v85 = m.ExcPending
										if v85 != 0 {
											return
										} else {
											*(*int32)(unsafe.Add(mBase, _consts[34])) = v84
											m.G0 = v4 + int32(16)
											return
										}
									}
								}
							}
						}
					}
				} else {
					v23 = *(*int32)(unsafe.Add(mBase, _consts[6]))
					if int32(3) < v23 {
						v39 = *(*int32)(unsafe.Add(mBase, _consts[25]))
						v40 = F_close(m, v39)
						mBase = m.M
						v41 = int32(_a44)
						v42 = int32(0)
						*(*int32)(unsafe.Add(mBase, _consts[28])) = v42
						*(*int64)(unsafe.Add(mBase, _consts[25])) = int64(-1)
						v48 = *(*int32)(unsafe.Add(mBase, _consts[29]))
						if v48 == v42 {
							v63 = int32(_a44)
							v64 = int64(0)
							*(*int64)(unsafe.Add(mBase, _consts[30])) = v64
							*(*int64)(unsafe.Add(mBase, _consts[31])) = int64(-1)
							*(*int64)(unsafe.Add(mBase, _consts[32])) = v64
							*(*int64)(unsafe.Add(mBase, _consts[33])) = v64
							F_killAppendOnlyChild(m)
							mBase = m.M
							v78 = m.ExcPending
							if v78 != 0 {
								return
							} else {
								v80 = *(*int32)(unsafe.Add(mBase, _consts[34]))
								F_sdsfree(m, v80)
								mBase = m.M
								v82 = m.ExcPending
								if v82 != 0 {
									return
								} else {
									v84 = F_sdsempty(m)
									mBase = m.M
									v85 = m.ExcPending
									if v85 != 0 {
										return
									} else {
										*(*int32)(unsafe.Add(mBase, _consts[34])) = v84
										m.G0 = v4 + int32(16)
										return
									}
								}
							}
						} else {
							v51 = int32(_a44)
							*(*int32)(unsafe.Add(mBase, _consts[29])) = int32(0)
							v55 = *(*int32)(unsafe.Add(mBase, _consts[6]))
							if int32(2) < v55 {
								v63 = int32(_a44)
								v64 = int64(0)
								*(*int64)(unsafe.Add(mBase, _consts[30])) = v64
								*(*int64)(unsafe.Add(mBase, _consts[31])) = int64(-1)
								*(*int64)(unsafe.Add(mBase, _consts[32])) = v64
								*(*int64)(unsafe.Add(mBase, _consts[33])) = v64
								F_killAppendOnlyChild(m)
								mBase = m.M
								v78 = m.ExcPending
								if v78 != 0 {
									return
								} else {
									v80 = *(*int32)(unsafe.Add(mBase, _consts[34]))
									F_sdsfree(m, v80)
									mBase = m.M
									v82 = m.ExcPending
									if v82 != 0 {
										return
									} else {
										v84 = F_sdsempty(m)
										mBase = m.M
										v85 = m.ExcPending
										if v85 != 0 {
											return
										} else {
											*(*int32)(unsafe.Add(mBase, _consts[34])) = v84
											m.G0 = v4 + int32(16)
											return
										}
									}
								}
							} else {
								F__serverLog(m, int32(2), int32(_a112), int32(0))
								mBase = m.M
								v62 = m.ExcPending
								if v62 != 0 {
									return
								} else {
									v63 = int32(_a44)
									v64 = int64(0)
									*(*int64)(unsafe.Add(mBase, _consts[30])) = v64
									*(*int64)(unsafe.Add(mBase, _consts[31])) = int64(-1)
									*(*int64)(unsafe.Add(mBase, _consts[32])) = v64
									*(*int64)(unsafe.Add(mBase, _consts[33])) = v64
									F_killAppendOnlyChild(m)
									mBase = m.M
									v78 = m.ExcPending
									if v78 != 0 {
										return
									} else {
										v80 = *(*int32)(unsafe.Add(mBase, _consts[34]))
										F_sdsfree(m, v80)
										mBase = m.M
										v82 = m.ExcPending
										if v82 != 0 {
											return
										} else {
											v84 = F_sdsempty(m)
											mBase = m.M
											v85 = m.ExcPending
											if v85 != 0 {
												return
											} else {
												*(*int32)(unsafe.Add(mBase, _consts[34])) = v84
												m.G0 = v4 + int32(16)
												return
											}
										}
									}
								}
							}
						}
					} else {
						v27 = *(*int32)(unsafe.Add(mBase, _consts[5]))
						v28 = F___strerror_l(m, v27, v27)
						mBase = m.M
						*(*int32)(unsafe.Add(mBase, uint32(v4))) = v28
						F__serverLog(m, int32(3), int32(_a113), v4)
						mBase = m.M
						v33 = m.ExcPending
						if v33 != 0 {
							return
						} else {
							v39 = *(*int32)(unsafe.Add(mBase, _consts[25]))
							v40 = F_close(m, v39)
							mBase = m.M
							v41 = int32(_a44)
							v42 = int32(0)
							*(*int32)(unsafe.Add(mBase, _consts[28])) = v42
							*(*int64)(unsafe.Add(mBase, _consts[25])) = int64(-1)
							v48 = *(*int32)(unsafe.Add(mBase, _consts[29]))
							if v48 == v42 {
								v63 = int32(_a44)
								v64 = int64(0)
								*(*int64)(unsafe.Add(mBase, _consts[30])) = v64
								*(*int64)(unsafe.Add(mBase, _consts[31])) = int64(-1)
								*(*int64)(unsafe.Add(mBase, _consts[32])) = v64
								*(*int64)(unsafe.Add(mBase, _consts[33])) = v64
								F_killAppendOnlyChild(m)
								mBase = m.M
								v78 = m.ExcPending
								if v78 != 0 {
									return
								} else {
									v80 = *(*int32)(unsafe.Add(mBase, _consts[34]))
									F_sdsfree(m, v80)
									mBase = m.M
									v82 = m.ExcPending
									if v82 != 0 {
										return
									} else {
										v84 = F_sdsempty(m)
										mBase = m.M
										v85 = m.ExcPending
										if v85 != 0 {
											return
										} else {
											*(*int32)(unsafe.Add(mBase, _consts[34])) = v84
											m.G0 = v4 + int32(16)
											return
										}
									}
								}
							} else {
								v51 = int32(_a44)
								*(*int32)(unsafe.Add(mBase, _consts[29])) = int32(0)
								v55 = *(*int32)(unsafe.Add(mBase, _consts[6]))
								if int32(2) < v55 {
									v63 = int32(_a44)
									v64 = int64(0)
									*(*int64)(unsafe.Add(mBase, _consts[30])) = v64
									*(*int64)(unsafe.Add(mBase, _consts[31])) = int64(-1)
									*(*int64)(unsafe.Add(mBase, _consts[32])) = v64
									*(*int64)(unsafe.Add(mBase, _consts[33])) = v64
									F_killAppendOnlyChild(m)
									mBase = m.M
									v78 = m.ExcPending
									if v78 != 0 {
										return
									} else {
										v80 = *(*int32)(unsafe.Add(mBase, _consts[34]))
										F_sdsfree(m, v80)
										mBase = m.M
										v82 = m.ExcPending
										if v82 != 0 {
											return
										} else {
											v84 = F_sdsempty(m)
											mBase = m.M
											v85 = m.ExcPending
											if v85 != 0 {
												return
											} else {
												*(*int32)(unsafe.Add(mBase, _consts[34])) = v84
												m.G0 = v4 + int32(16)
												return
											}
										}
									}
								} else {
									F__serverLog(m, int32(2), int32(_a112), int32(0))
									mBase = m.M
									v62 = m.ExcPending
									if v62 != 0 {
										return
									} else {
										v63 = int32(_a44)
										v64 = int64(0)
										*(*int64)(unsafe.Add(mBase, _consts[30])) = v64
										*(*int64)(unsafe.Add(mBase, _consts[31])) = int64(-1)
										*(*int64)(unsafe.Add(mBase, _consts[32])) = v64
										*(*int64)(unsafe.Add(mBase, _consts[33])) = v64
										F_killAppendOnlyChild(m)
										mBase = m.M
										v78 = m.ExcPending
										if v78 != 0 {
											return
										} else {
											v80 = *(*int32)(unsafe.Add(mBase, _consts[34]))
											F_sdsfree(m, v80)
											mBase = m.M
											v82 = m.ExcPending
											if v82 != 0 {
												return
											} else {
												v84 = F_sdsempty(m)
												mBase = m.M
												v85 = m.ExcPending
												if v85 != 0 {
													return
												} else {
													*(*int32)(unsafe.Add(mBase, _consts[34])) = v84
													m.G0 = v4 + int32(16)
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
		}
	}
}
func F_strbuf_resize(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v37 int32
	_ = v37
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v95 int32
	_ = v95
	var v98 int32
	_ = v98
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v108 int32
	_ = v108
	var v111 int32
	_ = v111
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	v7 = m.G0
	v9 = v7 - int32(48)
	m.G0 = v9
	v12 = l1 + int32(1)
	switch v12 {
	case 0:
		*(*int32)(unsafe.Add(mBase, uint32(v9)+32)) = int32(-1)
		v21 = m.G3
		F_die(m, v21+int32(_a2104), v9+int32(32))
		mBase = m.M
		v27 = m.ExcPending
		if v27 != 0 {
			return
		} else {
			base.Wasm_trap_unreachable()
			for {
			}
		}
	case 1:
		v13 = m.G3
		F_die(m, v13+int32(_a2105), int32(0))
		mBase = m.M
		v18 = m.ExcPending
		if v18 != 0 {
			return
		} else {
			base.Wasm_trap_unreachable()
			for {
			}
		}
	default:
		v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		if base.Ui32(int32(2147483645)) < base.Ui32(l1) {
			v44 = v12
		} else {
			if base.Ui32(v12) < base.Ui32(v28) {
				v44 = v12
			} else {
				v37 = v28
				for {
					if base.Ui32(v37) <= base.Ui32(l1) {
						v37 = v37 << (uint(int32(1)) % 32)
						continue
					} else {
						break
					}
					break
				}
				v44 = v37
			}
		}
		v47 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
		if v47 < int32(2) {
			*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v44
			v64 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			if v64 != 0 {
				if base.Ui32(v44) < base.Ui32(int32(-64)) {
					v74 = int32(-8)
					v77 = int32(11)
					if base.Ui32(v44) < base.Ui32(v77) {
						v83 = int32(16)
					} else {
						v83 = (v44 + v77) & v74
					}
					v84 = F_try_realloc_chunk(m, v64+v74, v83)
					mBase = m.M
					if v84 == int32(0) {
						v89 = F_emscripten_builtin_malloc(m, v44)
						mBase = m.M
						if v89 != 0 {
							v91 = int32(-4)
							v95 = *(*int32)(unsafe.Add(mBase, uint32(v64+v91)))
							if v95&int32(3) != 0 {
								v98 = v91
							} else {
								v98 = int32(-8)
							}
							v101 = v98 + v95&int32(-8)
							if base.Ui32(v101) < base.Ui32(v44) {
								v103 = v101
							} else {
								v103 = v44
							}
							v104 = F___memcpy(m, v89, v64, v103)
							mBase = m.M
							F_emscripten_builtin_free(m, v64)
							mBase = m.M
							v108 = v89
						} else {
							v108 = int32(0)
						}
					} else {
						v108 = v84 + int32(8)
					}
				} else {
					v70 = F___errno_location(m)
					mBase = m.M
					*(*int32)(unsafe.Add(mBase, uint32(v70))) = int32(48)
					v108 = int32(0)
				}
			} else {
				v67 = F_emscripten_builtin_malloc(m, v44)
				mBase = m.M
				v108 = v67
			}
			*(*int32)(unsafe.Add(mBase, uint32(l0))) = v108
			if v108 != 0 {
				v116 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
				*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v116 + int32(1)
				m.G0 = v9 + int32(48)
				return
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v9))) = l1
				v111 = m.G3
				F_die(m, v111+int32(_a2106), v9)
				mBase = m.M
				v115 = m.ExcPending
				if v115 != 0 {
					return
				} else {
					base.Wasm_trap_unreachable()
					for {
					}
				}
			}
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v9)+24)) = v44
			*(*int32)(unsafe.Add(mBase, uint32(v9)+20)) = v28
			*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = l0
			v53 = m.G3
			v54 = m.G397
			v55 = *(*int32)(unsafe.Add(mBase, uint32(v54)))
			v60 = F_fiprintf(m, v55, v53+int32(_a2107), v9+int32(16))
			mBase = m.M
			v61 = m.ExcPending
			if v61 != 0 {
				return
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v44
				v64 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
				if v64 != 0 {
					if base.Ui32(v44) < base.Ui32(int32(-64)) {
						v74 = int32(-8)
						v77 = int32(11)
						if base.Ui32(v44) < base.Ui32(v77) {
							v83 = int32(16)
						} else {
							v83 = (v44 + v77) & v74
						}
						v84 = F_try_realloc_chunk(m, v64+v74, v83)
						mBase = m.M
						if v84 == int32(0) {
							v89 = F_emscripten_builtin_malloc(m, v44)
							mBase = m.M
							if v89 != 0 {
								v91 = int32(-4)
								v95 = *(*int32)(unsafe.Add(mBase, uint32(v64+v91)))
								if v95&int32(3) != 0 {
									v98 = v91
								} else {
									v98 = int32(-8)
								}
								v101 = v98 + v95&int32(-8)
								if base.Ui32(v101) < base.Ui32(v44) {
									v103 = v101
								} else {
									v103 = v44
								}
								v104 = F___memcpy(m, v89, v64, v103)
								mBase = m.M
								F_emscripten_builtin_free(m, v64)
								mBase = m.M
								v108 = v89
							} else {
								v108 = int32(0)
							}
						} else {
							v108 = v84 + int32(8)
						}
					} else {
						v70 = F___errno_location(m)
						mBase = m.M
						*(*int32)(unsafe.Add(mBase, uint32(v70))) = int32(48)
						v108 = int32(0)
					}
				} else {
					v67 = F_emscripten_builtin_malloc(m, v44)
					mBase = m.M
					v108 = v67
				}
				*(*int32)(unsafe.Add(mBase, uint32(l0))) = v108
				if v108 != 0 {
					v116 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
					*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v116 + int32(1)
					m.G0 = v9 + int32(48)
					return
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v9))) = l1
					v111 = m.G3
					F_die(m, v111+int32(_a2106), v9)
					mBase = m.M
					v115 = m.ExcPending
					if v115 != 0 {
						return
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		}
	}
}
func F_strchr(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v16 int32
	_ = v16
	var v21 int32
	_ = v21
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v63 int32
	_ = v63
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v78 int32
	_ = v78
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v92 int32
	_ = v92
	var v102 int32
	_ = v102
	var v104 int32
	_ = v104
	var v108 int32
	_ = v108
	v7 = l1 & int32(255)
	if v7 == int32(0) {
		goto L5
	} else {
		goto L6
	}
L1:
	;
	v104 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v102))))
	if v104 == l1&int32(255) {
		goto L23
	} else {
		goto L24
	}
L2:
	;
	v102 = v92
	goto L1
L3:
	;
	v83 = v78
	goto L19
L4:
	;
	v78 = v69
	goto L3
L5:
	;
	v67 = F_strlen(m, l0)
	mBase = m.M
	v102 = l0 + v67
	goto L1
L6:
	;
	if l0&int32(3) == int32(0) {
		v29 = l0
		goto L7
	} else {
		goto L8
	}
L7:
	;
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v29)))
	v38 = int32(-2139062144)
	if (int32(16843008)-v35|v35)&v38 != v38 {
		v69 = v29
		goto L4
	} else {
		goto L14
	}
L8:
	;
	v16 = l0
	goto L9
L9:
	;
	v21 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16))))
	if v21 == int32(0) {
		v92 = v16
		goto L2
	} else {
		goto L11
	}
L10:
	;
	v29 = v26
	goto L7
L11:
	;
	if v21 == l1&int32(255) {
		v92 = v16
		goto L2
	} else {
		goto L12
	}
L12:
	;
	v26 = v16 + int32(1)
	if v26&int32(3) != 0 {
		v16 = v26
		goto L9
	} else {
		goto L13
	}
L13:
	;
	goto L10
L14:
	;
	v44 = v29
	v47 = v35
	goto L15
L15:
	;
	v50 = v47 ^ v7*int32(16843009)
	v53 = int32(-2139062144)
	if (int32(16843008)-v50|v50)&v53 != v53 {
		v69 = v44
		goto L4
	} else {
		goto L17
	}
L17:
	;
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v44)+4))
	v59 = v44 + int32(4)
	v63 = int32(-2139062144)
	if (v57|(int32(16843008)-v57))&v63 == v63 {
		v44 = v59
		v47 = v57
		goto L15
	} else {
		goto L18
	}
L18:
	;
	v78 = v59
	goto L3
L19:
	;
	v84 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v83))))
	if v84 == int32(0) {
		v92 = v83
		goto L2
	} else {
		goto L21
	}
L20:
	;
	v92 = v83
	goto L2
L21:
	;
	if v84 != l1&int32(255) {
		v83 = v83 + int32(1)
		goto L19
	} else {
		goto L22
	}
L22:
	;
	goto L20
L23:
	;
	v108 = v102
	goto L25
L24:
	;
	v108 = int32(0)
	goto L25
L25:
	;
	return v108
}
func F_strcoll(m *base.Module, l0 int32, l1 int32) int32 {
	var v4 int32
	_ = v4
	v4 = F_strcmp(m, l0, l1)
	return v4
}
func F_strdup(m *base.Module, l0 int32) int32 {
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
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	if l0&int32(3) == int32(0) {
		v25 = l0
		goto L4
	} else {
		goto L5
	}
L1:
	;
	if v60 == int32(0) {
		v67 = v61
		goto L20
	} else {
		goto L21
	}
L2:
	;
	v60 = v58 + int32(1)
	v61 = F_emscripten_builtin_malloc(m, v60)
	mBase = m.M
	if v61 != 0 {
		goto L1
	} else {
		goto L18
	}
L3:
	;
	v58 = v50 - l0
	goto L2
L4:
	;
	v29 = v25
	goto L12
L5:
	;
	v11 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v11 != 0 {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	v14 = l0
	goto L8
L7:
	;
	v58 = l0 - l0
	goto L2
L8:
	;
	v18 = v14 + int32(1)
	if v18&int32(3) == int32(0) {
		v25 = v18
		goto L4
	} else {
		goto L10
	}
L10:
	;
	v23 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18))))
	if v23 != 0 {
		v14 = v18
		goto L8
	} else {
		goto L11
	}
L11:
	;
	v50 = v18
	goto L3
L12:
	;
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v29)))
	v38 = int32(-2139062144)
	if (int32(16843008)-v35|v35)&v38 == v38 {
		v29 = v29 + int32(4)
		goto L12
	} else {
		goto L14
	}
L13:
	;
	v44 = v29
	goto L15
L14:
	;
	goto L13
L15:
	;
	v48 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v44))))
	if v48 != 0 {
		v44 = v44 + int32(1)
		goto L15
	} else {
		goto L17
	}
L16:
	;
	v50 = v44
	goto L3
L17:
	;
	goto L16
L18:
	;
	return int32(0)
L19:
	;
	return v67
L20:
	;
	goto L19
L21:
	;
	v66 = F__emscripten_memcpy_bulkmem(m, v61, l0, v60)
	mBase = m.M
	v67 = v66
	goto L20
}
func F_strftime(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	v6 = *(*int32)(unsafe.Add(mBase, _consts[0]))
	v7 = F___strftime_l(m, l0, l1, l2, l3, v6)
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		return int32(0)
	} else {
		return v7
	}
}
func F_string2d(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
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
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v31 int64
	_ = v31
	var v36 int64
	_ = v36
	var v44 int32
	_ = v44
	var v51 int32
	_ = v51
	var v55 int32
	_ = v55
	var v57 float64
	_ = v57
	var v64 int32
	_ = v64
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v84 int64
	_ = v84
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v97 int32
	_ = v97
	var v102 int32
	_ = v102
	var v104 int32
	_ = v104
	var v107 int64
	_ = v107
	var v108 int64
	_ = v108
	var v114 int32
	_ = v114
	var v120 int32
	_ = v120
	var v123 int32
	_ = v123
	v9 = m.G0
	v11 = v9 - int32(16)
	m.G0 = v11
	v13 = int32(9116376)
	v14 = int32(0)
	*(*int32)(unsafe.Add(mBase, _consts[5])) = v14
	v17 = v11 + int32(12)
	v21 = m.G0
	v23 = v21 - int32(32)
	m.G0 = v23
	v25 = F___errno_location(m)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v25))) = v14
	v31 = *(*int64)(unsafe.Add(mBase, _consts[378]))
	*(*int64)(unsafe.Add(mBase, uint32(v23+int32(8)))) = v31
	*(*int64)(unsafe.Add(mBase, uint32(v23)+24)) = int64(0)
	v36 = *(*int64)(unsafe.Add(mBase, _consts[379]))
	*(*int64)(unsafe.Add(mBase, uint32(v23))) = v36
	F_ffc_from_chars_double_options(m, v23+int32(16), l0, l0+l1, v23+int32(24), v23)
	mBase = m.M
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v23)+20))
	if v44 == v14 {
	} else {
		if v44 == int32(2) {
			v51 = int32(68)
		} else {
			v51 = int32(28)
		}
		*(*int32)(unsafe.Add(mBase, uint32(v25))) = v51
	}
	if v17 == int32(0) {
	} else {
		v55 = *(*int32)(unsafe.Add(mBase, uint32(v23)+16))
		*(*int32)(unsafe.Add(mBase, uint32(v17))) = v55
	}
	v57 = *(*float64)(unsafe.Add(mBase, uint32(v23)+24))
	m.G0 = v23 + int32(32)
	*(*float64)(unsafe.Add(mBase, uint32(l2))) = v57
	if l1 == int32(0) {
		v120 = int32(0)
		*(*int32)(unsafe.Add(mBase, _consts[5])) = v120
		v123 = v120
	} else {
		v64 = int32(*(*int8)(unsafe.Add(mBase, uint32(l0))))
		if v64 == int32(32) {
			v120 = int32(0)
			*(*int32)(unsafe.Add(mBase, _consts[5])) = v120
			v123 = v120
		} else {
			if base.Ui32(int32(-6)) < base.Ui32(v64+int32(-14)) {
				v120 = int32(0)
				*(*int32)(unsafe.Add(mBase, _consts[5])) = v120
				v123 = v120
			} else {
				v71 = *(*int32)(unsafe.Add(mBase, uint32(v11)+12))
				if v71-l0 != l1 {
					v120 = int32(0)
					*(*int32)(unsafe.Add(mBase, _consts[5])) = v120
					v123 = v120
				} else {
					v74 = *(*int32)(unsafe.Add(mBase, _consts[5]))
					if v74 == int32(68) {
						if base.F64_eq(base.F64_abs(v57), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
							v120 = int32(0)
							*(*int32)(unsafe.Add(mBase, _consts[5])) = v120
							v123 = v120
						} else {
							v84 = base.I64_reinterpret_f64(v57)
							v88 = int32(2047)
							v89 = base.I32_wrap_i64(int64(base.Ui64(v84)>>(uint(int64(52))%64))) & v88
							if v89 == v88 {
								v102 = base.B2i32(v84&int64(4503599627370495) == int64(0))
								v104 = v102
							} else {
								if v89 != 0 {
									v102 = int32(4)
									v104 = v102
								} else {
									if base.F64_eq(v57, float64(0)) != 0 {
										v97 = int32(2)
									} else {
										v97 = int32(3)
									}
									v104 = v97
								}
							}
							if v104 == int32(2) {
								v120 = int32(0)
								*(*int32)(unsafe.Add(mBase, _consts[5])) = v120
								v123 = v120
							} else {
								v107 = *(*int64)(unsafe.Add(mBase, uint32(l2)))
								v108 = v107
								if base.Ui64(int64(9218868437227405312)) < base.Ui64(v108&int64(9223372036854775807)) {
									v120 = int32(0)
									*(*int32)(unsafe.Add(mBase, _consts[5])) = v120
									v123 = v120
								} else {
									v114 = *(*int32)(unsafe.Add(mBase, _consts[5]))
									if v114 != int32(28) {
										v123 = int32(1)
									} else {
										v120 = int32(0)
										*(*int32)(unsafe.Add(mBase, _consts[5])) = v120
										v123 = v120
									}
								}
							}
						}
					} else {
						v108 = base.I64_reinterpret_f64(v57)
						if base.Ui64(int64(9218868437227405312)) < base.Ui64(v108&int64(9223372036854775807)) {
							v120 = int32(0)
							*(*int32)(unsafe.Add(mBase, _consts[5])) = v120
							v123 = v120
						} else {
							v114 = *(*int32)(unsafe.Add(mBase, _consts[5]))
							if v114 != int32(28) {
								v123 = int32(1)
							} else {
								v120 = int32(0)
								*(*int32)(unsafe.Add(mBase, _consts[5])) = v120
								v123 = v120
							}
						}
					}
				}
			}
		}
	}
	m.G0 = v11 + int32(16)
	return v123
}
func F_string2ld(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v30 int32
	_ = v30
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v58 int64
	_ = v58
	var v59 int64
	_ = v59
	var v60 int32
	_ = v60
	var v74 int64
	_ = v74
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v88 int32
	_ = v88
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v105 int64
	_ = v105
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v119 int32
	_ = v119
	var v123 int32
	_ = v123
	var v125 int32
	_ = v125
	var v128 int32
	_ = v128
	var v135 int32
	_ = v135
	v9 = m.G0
	v11 = v9 - int32(5152)
	m.G0 = v11
	if base.Ui32(l1+int32(-5120)) < base.Ui32(int32(-5119)) {
		v135 = int32(0)
		m.G0 = v11 + int32(5152)
		return v135
	} else {
		if l1 == int32(0) {
		} else {
			v22 = F__emscripten_memcpy_bulkmem(m, v11+int32(32), l0, l1)
			mBase = m.M
		}
		v24 = int32(0)
		*(*uint8)(unsafe.Add(mBase, uint32(v11+int32(32)+l1))) = uint8(v24)
		v30 = int32(9116376)
		*(*int32)(unsafe.Add(mBase, _consts[5])) = int32(0)
		F_strtold(m, v11+int32(8), v11+int32(32), v11+int32(28))
		mBase = m.M
		v42 = m.ExcPending
		if v42 != 0 {
			return int32(0)
		} else {
			v43 = int32(*(*int8)(unsafe.Add(mBase, uint32(v11)+32)))
			if v43 == int32(32) {
				v135 = v24
			} else {
				if base.Ui32(int32(-6)) < base.Ui32(v43+int32(-14)) {
					v135 = v24
				} else {
					v50 = *(*int32)(unsafe.Add(mBase, uint32(v11)+28))
					v51 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v50))))
					if v51 != 0 {
						v135 = v24
					} else {
						if v50-(v11+int32(32)) != l1 {
							v135 = v24
						} else {
							v58 = *(*int64)(unsafe.Add(mBase, uint32(v11+int32(16))))
							v59 = *(*int64)(unsafe.Add(mBase, uint32(v11)+8))
							v60 = *(*int32)(unsafe.Add(mBase, _consts[5]))
							if v60 != int32(68) {
								v98 = v60
								if v98 == int32(28) {
									v135 = v24
								} else {
									v105 = v58 & int64(281474976710655)
									v109 = int32(32767)
									v110 = base.I32_wrap_i64(int64(base.Ui64(v58)>>(uint(int64(48))%64))) & v109
									if v110 == v109 {
										v123 = base.B2i32(v105|v59 == int64(0))
										v125 = v123
									} else {
										if v110 != 0 {
											v123 = int32(4)
											v125 = v123
										} else {
											if v105|v59 == int64(0) {
												v119 = int32(2)
											} else {
												v119 = int32(3)
											}
											v125 = v119
										}
									}
									if v125 == int32(0) {
										v135 = v24
									} else {
										v128 = int32(1)
										if l2 == int32(0) {
											v135 = v128
										} else {
											*(*int64)(unsafe.Add(mBase, uint32(l2))) = v59
											*(*int64)(unsafe.Add(mBase, uint32(l2)+8)) = v58
											v135 = v128
										}
									}
								}
							} else {
								if v59|(v58&int64(9223372036854775807)^int64(9223090561878065152)) == int64(0) {
									v135 = v24
								} else {
									v74 = v58 & int64(281474976710655)
									v78 = int32(32767)
									v79 = base.I32_wrap_i64(int64(base.Ui64(v58)>>(uint(int64(48))%64))) & v78
									if v79 == v78 {
										v92 = base.B2i32(v74|v59 == int64(0))
										v94 = v92
									} else {
										if v79 != 0 {
											v92 = int32(4)
											v94 = v92
										} else {
											if v74|v59 == int64(0) {
												v88 = int32(2)
											} else {
												v88 = int32(3)
											}
											v94 = v88
										}
									}
									if v94 == int32(2) {
										v135 = v24
									} else {
										v97 = *(*int32)(unsafe.Add(mBase, _consts[5]))
										v98 = v97
										if v98 == int32(28) {
											v135 = v24
										} else {
											v105 = v58 & int64(281474976710655)
											v109 = int32(32767)
											v110 = base.I32_wrap_i64(int64(base.Ui64(v58)>>(uint(int64(48))%64))) & v109
											if v110 == v109 {
												v123 = base.B2i32(v105|v59 == int64(0))
												v125 = v123
											} else {
												if v110 != 0 {
													v123 = int32(4)
													v125 = v123
												} else {
													if v105|v59 == int64(0) {
														v119 = int32(2)
													} else {
														v119 = int32(3)
													}
													v125 = v119
												}
											}
											if v125 == int32(0) {
												v135 = v24
											} else {
												v128 = int32(1)
												if l2 == int32(0) {
													v135 = v128
												} else {
													*(*int64)(unsafe.Add(mBase, uint32(l2))) = v59
													*(*int64)(unsafe.Add(mBase, uint32(l2)+8)) = v58
													v135 = v128
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
			m.G0 = v11 + int32(5152)
			return v135
		}
	}
}
func F_string2ll(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v21 int32
	_ = v21
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v52 int64
	_ = v52
	var v58 int32
	_ = v58
	var v60 int64
	_ = v60
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v74 int64
	_ = v74
	var v79 int64
	_ = v79
	var v83 int32
	_ = v83
	var v85 int64
	_ = v85
	var v87 int32
	_ = v87
	var v95 int64
	_ = v95
	var v119 int64
	_ = v119
	var v138 int32
	_ = v138
	v4 = int32(0)
	if base.Ui32(l1+int32(-21)) < base.Ui32(int32(-20)) {
		v138 = v4
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return v138
L2:
	;
	v16 = int32(1)
	v17 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if l1 != v16 {
		goto L5
	} else {
		goto L6
	}
L3:
	;
	v138 = int32(1)
	goto L1
L4:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l2))) = v119
	goto L3
L5:
	;
	if v17&int32(255) == int32(45) {
		goto L10
	} else {
		goto L11
	}
L6:
	;
	v21 = v17 + int32(-48)
	if base.Ui32(int32(9)) < base.Ui32(v21&int32(255)) {
		v138 = v4
		goto L1
	} else {
		goto L7
	}
L7:
	;
	if l2 == int32(0) {
		goto L3
	} else {
		goto L8
	}
L8:
	;
	v119 = base.I64_extend_i32_u(v21) & int64(255)
	goto L4
L9:
	;
	if base.Ui32(int32(8)) < base.Ui32((v40+int32(-49))&int32(255)) {
		v138 = v4
		goto L1
	} else {
		goto L12
	}
L10:
	;
	v37 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+1)))
	v39 = int32(2)
	v40 = v37
	v41 = l0 + int32(1)
	goto L9
L11:
	;
	v39 = v16
	v40 = v17
	v41 = l0
	goto L9
L12:
	;
	v52 = base.I64_extend_i32_u(v40+int32(-48)) & int64(255)
	if base.Ui32(l1) <= base.Ui32(v39) {
		v95 = v52
		goto L13
	} else {
		goto L14
	}
L13:
	;
	if v17&int32(255) != int32(45) {
		goto L21
	} else {
		goto L22
	}
L14:
	;
	v58 = v39
	v60 = v52
	v62 = v41
	goto L15
L15:
	;
	v64 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v62)+1)))
	if base.Ui32((v64+int32(-58))&int32(255)) < base.Ui32(int32(246)) {
		v138 = v4
		goto L1
	} else {
		goto L17
	}
L16:
	;
	v95 = v85
	goto L13
L17:
	;
	if base.Ui64(int64(1844674407370955161)) < base.Ui64(v60) {
		v138 = v4
		goto L1
	} else {
		goto L18
	}
L18:
	;
	v74 = v60 * int64(10)
	v79 = base.I64_extend_i32_u(v64+int32(-48)) & int64(255)
	if base.Ui64(v79^int64(-1)) < base.Ui64(v74) {
		v138 = v4
		goto L1
	} else {
		goto L19
	}
L19:
	;
	v83 = int32(1)
	v85 = v74 + v79
	v87 = v58 + v83
	if v87 != l1 {
		v58 = v87
		v60 = v85
		v62 = v62 + v83
		goto L15
	} else {
		goto L20
	}
L20:
	;
	goto L16
L21:
	;
	if v95 < int64(0) {
		v138 = v4
		goto L1
	} else {
		goto L25
	}
L22:
	;
	if base.Ui64(int64(-9223372036854775807-1)) < base.Ui64(v95) {
		v138 = v4
		goto L1
	} else {
		goto L23
	}
L23:
	;
	if l2 == int32(0) {
		goto L3
	} else {
		goto L24
	}
L24:
	;
	v119 = int64(0) - v95
	goto L4
L25:
	;
	if l2 == int32(0) {
		goto L3
	} else {
		goto L26
	}
L26:
	;
	v119 = v95
	goto L4
}
func F_stringmatchlen(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v16 int32
	_ = v16
	v6 = int32(0)
	v7 = m.G0
	v8 = int32(16)
	v9 = v7 - v8
	m.G0 = v9
	*(*int32)(unsafe.Add(mBase, uint32(v9)+12)) = v6
	v16 = F_stringmatchlen_impl(m, l0, l1, l2, l3, l4, v9+int32(12), v6)
	mBase = m.M
	m.G0 = v9 + v8
	return v16
}
func F_strlenCommand(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v4 = *(*int32)(unsafe.Add(mBase, uint32(v3)+4))
	v6 = *(*int32)(unsafe.Add(mBase, _consts[85]))
	v7 = F_lookupKeyReadOrReply(m, l0, v4, v6)
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		return
	} else {
		if v7 == int32(0) {
			return
		} else {
			v12 = F_checkType(m, l0, v7, int32(0))
			mBase = m.M
			v13 = m.ExcPending
			if v13 != 0 {
				return
			} else {
				if v12 != 0 {
					return
				} else {
					v14 = F_stringObjectLen(m, v7)
					mBase = m.M
					v15 = m.ExcPending
					if v15 != 0 {
						return
					} else {
						F_addReplyLongLong(m, l0, base.I64_extend_i32_u(v14))
						mBase = m.M
						v18 = m.ExcPending
						if v18 != 0 {
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
func F_strstr(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
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
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
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
	var v100 int32
	_ = v100
	var v104 int32
	_ = v104
	var v107 int32
	_ = v107
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v124 int32
	_ = v124
	var v129 int32
	_ = v129
	var v132 int32
	_ = v132
	var v137 int32
	_ = v137
	var v139 int32
	_ = v139
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v154 int32
	_ = v154
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v169 int32
	_ = v169
	var v181 int32
	_ = v181
	var v185 int32
	_ = v185
	var v188 int32
	_ = v188
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v198 int32
	_ = v198
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v212 int32
	_ = v212
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	v5 = int32(*(*int8)(unsafe.Add(mBase, uint32(l1))))
	if v5 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v7 = int32(0)
	v8 = F___strchrnul(m, l0, v5)
	mBase = m.M
	v10 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8))))
	if v10 == v5&int32(255) {
		goto L5
	} else {
		goto L6
	}
L2:
	;
	return l0
L3:
	;
	return v215
L4:
	;
	if v14 == int32(0) {
		v215 = v7
		goto L3
	} else {
		goto L8
	}
L5:
	;
	v14 = v8
	goto L7
L6:
	;
	v14 = v7
	goto L7
L7:
	;
	goto L4
L8:
	;
	v17 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+1)))
	if v17 != 0 {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	v19 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+1)))
	if v19 == int32(0) {
		v215 = v7
		goto L3
	} else {
		goto L11
	}
L10:
	;
	return v14
L11:
	;
	v22 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+2)))
	if v22 != 0 {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	v72 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+2)))
	if v72 == int32(0) {
		v215 = v7
		goto L3
	} else {
		goto L25
	}
L13:
	;
	v23 = int32(0)
	v27 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+1)))
	v29 = base.B2i32(v27 != v23)
	if v27 == v23 {
		v63 = v14
		v66 = v29
		goto L15
	} else {
		goto L16
	}
L14:
	;
	return v70
L15:
	;
	if v66 != 0 {
		goto L22
	} else {
		goto L23
	}
L16:
	;
	v32 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14))))
	v33 = int32(8)
	v35 = v32<<(uint(v33)%32) | v27
	v36 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	v39 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+1)))
	v40 = v36<<(uint(v33)%32) | v39
	if v35 == v40 {
		v63 = v14
		v66 = v29
		goto L15
	} else {
		goto L17
	}
L17:
	;
	v45 = v14 + int32(1)
	v48 = v35
	goto L18
L18:
	;
	v50 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v45)+1)))
	v51 = int32(0)
	v52 = base.B2i32(v50 != v51)
	if v50 == v51 {
		v63 = v45
		v66 = v52
		goto L15
	} else {
		goto L20
	}
L19:
	;
	v63 = v45
	v66 = v52
	goto L15
L20:
	;
	v61 = v48<<(uint(int32(8))%32)&int32(65280) | v50
	if v61 != v40 {
		v45 = v45 + int32(1)
		v48 = v61
		goto L18
	} else {
		goto L21
	}
L21:
	;
	goto L19
L22:
	;
	v70 = v63
	goto L24
L23:
	;
	v70 = int32(0)
	goto L24
L24:
	;
	goto L14
L25:
	;
	v75 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+3)))
	if v75 != 0 {
		goto L26
	} else {
		goto L27
	}
L26:
	;
	v139 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+3)))
	if v139 == int32(0) {
		v215 = v7
		goto L3
	} else {
		goto L40
	}
L27:
	;
	v76 = int32(0)
	v81 = v14 + int32(2)
	v82 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+2)))
	if v82 == v76 {
		goto L30
	} else {
		goto L31
	}
L28:
	;
	return v137
L29:
	;
	if v132 != 0 {
		goto L37
	} else {
		goto L38
	}
L30:
	;
	v129 = v81
	v132 = base.B2i32(v82 != v76)
	goto L29
L31:
	;
	v87 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+1)))
	v88 = int32(16)
	v90 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14))))
	v91 = int32(24)
	v94 = int32(8)
	v96 = v87<<(uint(v88)%32) | v90<<(uint(v91)%32) | v82<<(uint(v94)%32)
	v97 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+1)))
	v100 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	v104 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+2)))
	v107 = v97<<(uint(v88)%32) | v100<<(uint(v91)%32) | v104<<(uint(v94)%32)
	if v96 == v107 {
		goto L30
	} else {
		goto L32
	}
L32:
	;
	v111 = v81
	v112 = v96
	goto L33
L33:
	;
	v116 = v111 + int32(1)
	v117 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v111)+1)))
	v118 = int32(0)
	v119 = base.B2i32(v117 != v118)
	if v117 == v118 {
		v129 = v116
		v132 = v119
		goto L29
	} else {
		goto L35
	}
L35:
	;
	v124 = (v112 | v117) << (uint(int32(8)) % 32)
	if v124 != v107 {
		v111 = v116
		v112 = v124
		goto L33
	} else {
		goto L36
	}
L36:
	;
	v129 = v116
	v132 = v119
	goto L29
L37:
	;
	v137 = v129 + int32(-2)
	goto L39
L38:
	;
	v137 = int32(0)
	goto L39
L39:
	;
	goto L28
L40:
	;
	v142 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+4)))
	if v142 != 0 {
		goto L41
	} else {
		goto L42
	}
L41:
	;
	v214 = F_twoway_strstr(m, v14, l1)
	mBase = m.M
	v215 = v214
	goto L3
L42:
	;
	v143 = int32(0)
	v148 = v14 + int32(3)
	v149 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+3)))
	if v149 == v143 {
		goto L45
	} else {
		goto L46
	}
L43:
	;
	return v212
L44:
	;
	if v207 != 0 {
		goto L52
	} else {
		goto L53
	}
L45:
	;
	v206 = v148
	v207 = base.B2i32(v149 != v143)
	goto L44
L46:
	;
	v154 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+1)))
	v157 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14))))
	v158 = int32(24)
	v161 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+2)))
	v162 = int32(8)
	v165 = v154<<(uint(int32(16))%32) | v157<<(uint(v158)%32) | v161<<(uint(v162)%32) | v149
	v166 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v169 = int32(65280)
	v181 = v166<<(uint(v158)%32) | v166&v169<<(uint(v162)%32) | (int32(base.Ui32(v166)>>(uint(v162)%32))&v169 | int32(base.Ui32(v166)>>(uint(v158)%32)))
	if v165 == v181 {
		goto L45
	} else {
		goto L47
	}
L47:
	;
	v185 = v148
	v188 = v165
	goto L48
L48:
	;
	v190 = v185 + int32(1)
	v191 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v185)+1)))
	v192 = int32(0)
	v193 = base.B2i32(v191 != v192)
	if v191 == v192 {
		v206 = v190
		v207 = v193
		goto L44
	} else {
		goto L50
	}
L50:
	;
	v198 = v188<<(uint(int32(8))%32) | v191
	if v198 != v181 {
		v185 = v190
		v188 = v198
		goto L48
	} else {
		goto L51
	}
L51:
	;
	v206 = v190
	v207 = v193
	goto L44
L52:
	;
	v212 = v206 + int32(-3)
	goto L54
L53:
	;
	v212 = int32(0)
	goto L54
L54:
	;
	goto L43
}
func F_strtol(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	var v5 int64
	_ = v5
	v5 = F_strtox_2(m, l0, l1, l2, int64(2147483648))
	return base.I32_wrap_i64(v5)
}
func F_strtold(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v12 int32
	_ = v12
	var v13 int64
	_ = v13
	var v16 int64
	_ = v16
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	F_strtox_1(m, v8, l1, l2, int32(2))
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		return
	} else {
		v13 = *(*int64)(unsafe.Add(mBase, uint32(v8)))
		v16 = *(*int64)(unsafe.Add(mBase, uint32(v8+int32(8))))
		*(*int64)(unsafe.Add(mBase, uint32(l0)+8)) = v16
		*(*int64)(unsafe.Add(mBase, uint32(l0))) = v13
		m.G0 = v8 + int32(16)
		return
	}
}
func F_strtoull(m *base.Module, l0 int32, l1 int32, l2 int32) int64 {
	var v5 int64
	_ = v5
	v5 = F_strtox_2(m, l0, l1, l2, int64(-1))
	return v5
}
func F_sumEngineUsedMemory(m *base.Module, l0 int32, l1 int32) {
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
	var v12 int32
	_ = v12
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	F_scriptingEngineCallGetMemoryInfo(m, v6, l0, int32(2))
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		return
	} else {
		v11 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
		v12 = *(*int32)(unsafe.Add(mBase, uint32(v6)+8))
		*(*int32)(unsafe.Add(mBase, uint32(l1))) = v11 + v12
		m.G0 = v6 + int32(16)
		return
	}
}
func F_sunsubscribeCommand(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
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
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v38 int32
	_ = v38
	var v52 int32
	_ = v52
	var v57 int32
	_ = v57
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v66 int64
	_ = v66
	var v69 int64
	_ = v69
	var v72 int64
	_ = v72
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v91 int64
	_ = v91
	var v96 int64
	_ = v96
	var v99 int64
	_ = v99
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
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
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v130 int32
	_ = v130
	var v138 int32
	_ = v138
	var v140 int32
	_ = v140
	v8 = m.G0
	v10 = v8 - int32(64)
	m.G0 = v10
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+100))
	if v12 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v38 == int32(1) {
		goto L9
	} else {
		goto L10
	}
L2:
	;
	v14 = F_valkey_malloc(m, int32(32))
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	return
L4:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+100)) = v14
	v18 = F_hashtableCreate(m, int32(_a1106))
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		goto L3
	} else {
		goto L5
	}
L5:
	;
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+100))
	*(*int32)(unsafe.Add(mBase, uint32(v20))) = v18
	v23 = F_hashtableCreate(m, int32(_a1106))
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		goto L3
	} else {
		goto L6
	}
L6:
	;
	v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+100))
	*(*int32)(unsafe.Add(mBase, uint32(v25)+4)) = v23
	v28 = F_hashtableCreate(m, int32(_a1106))
	mBase = m.M
	v29 = m.ExcPending
	if v29 != 0 {
		goto L3
	} else {
		goto L7
	}
L7:
	;
	v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)+100))
	*(*int32)(unsafe.Add(mBase, uint32(v30)+24)) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v30)+16)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v30)+8)) = v28
	goto L1
L8:
	;
	v111 = *(*int32)(unsafe.Add(mBase, uint32(l0)+100))
	v112 = *(*int32)(unsafe.Add(mBase, uint32(v111)))
	v113 = *(*int32)(unsafe.Add(mBase, uint32(v112)+20))
	v114 = *(*int32)(unsafe.Add(mBase, uint32(v112)+16))
	goto L18
L9:
	;
	v85 = int32(0)
	v86 = *(*int32)(unsafe.Add(mBase, _consts[619]))
	*(*int32)(unsafe.Add(mBase, uint32(v10+int32(24)))) = v86
	v91 = *(*int64)(unsafe.Add(mBase, _consts[620]))
	*(*int64)(unsafe.Add(mBase, uint32(v10+int32(16)))) = v91
	v96 = *(*int64)(unsafe.Add(mBase, _consts[621]))
	*(*int64)(unsafe.Add(mBase, uint32(v10+int32(8)))) = v96
	v99 = *(*int64)(unsafe.Add(mBase, _consts[622]))
	*(*int64)(unsafe.Add(mBase, uint32(v10))) = v99
	v102 = F_pubsubUnsubscribeAllChannelsInternal(m, l0, int32(1), v10)
	mBase = m.M
	v103 = m.ExcPending
	if v103 != 0 {
		goto L3
	} else {
		goto L16
	}
L10:
	;
	if v38 < int32(2) {
		goto L8
	} else {
		goto L11
	}
L11:
	;
	v52 = int32(1)
	goto L12
L12:
	;
	v57 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v57+v52<<(uint(int32(2))%32))))
	v62 = int32(0)
	v63 = *(*int32)(unsafe.Add(mBase, _consts[619]))
	*(*int32)(unsafe.Add(mBase, uint32(v10+int32(56)))) = v63
	v66 = *(*int64)(unsafe.Add(mBase, _consts[620]))
	*(*int64)(unsafe.Add(mBase, uint32(v10+int32(48)))) = v66
	v69 = *(*int64)(unsafe.Add(mBase, _consts[621]))
	*(*int64)(unsafe.Add(mBase, uint32(v10+int32(40)))) = v69
	v72 = *(*int64)(unsafe.Add(mBase, _consts[622]))
	*(*int64)(unsafe.Add(mBase, uint32(v10)+32)) = v72
	v77 = F_pubsubUnsubscribeChannel(m, l0, v61, int32(1), v10+int32(32))
	mBase = m.M
	v78 = m.ExcPending
	if v78 != 0 {
		goto L3
	} else {
		goto L14
	}
L14:
	;
	v80 = v52 + int32(1)
	v81 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v80 < v81 {
		v52 = v80
		goto L12
	} else {
		goto L15
	}
L15:
	;
	goto L8
L16:
	;
	goto L8
L17:
	;
	m.G0 = v10 + int32(64)
	return
L18:
	;
	v116 = *(*int32)(unsafe.Add(mBase, uint32(l0)+100))
	v117 = *(*int32)(unsafe.Add(mBase, uint32(v116)+4))
	v118 = *(*int32)(unsafe.Add(mBase, uint32(v117)+20))
	v119 = *(*int32)(unsafe.Add(mBase, uint32(v117)+16))
	goto L19
L19:
	;
	v123 = *(*int32)(unsafe.Add(mBase, uint32(l0)+100))
	v124 = *(*int32)(unsafe.Add(mBase, uint32(v123)+8))
	v125 = *(*int32)(unsafe.Add(mBase, uint32(v124)+20))
	v126 = *(*int32)(unsafe.Add(mBase, uint32(v124)+16))
	goto L20
L20:
	;
	if v113+v114+(v118+v119) != int32(0)-(v125+v126) {
		goto L17
	} else {
		goto L21
	}
L21:
	;
	v130 = *(*int32)(unsafe.Add(mBase, uint32(l0)+200))
	if v130&int32(262144) == int32(0) {
		goto L17
	} else {
		goto L22
	}
L22:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+200)) = v130 & int32(-262145)
	v138 = int32(_a44)
	v140 = *(*int32)(unsafe.Add(mBase, _consts[623]))
	*(*int32)(unsafe.Add(mBase, _consts[623])) = v140 + int32(-1)
	goto L17
}
func F_sweeplist(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
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
	var v24 int32
	_ = v24
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v42 int32
	_ = v42
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v72 int32
	_ = v72
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	if v8 == int32(0) {
		v90 = l1
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return v90
L2:
	;
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v12 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+20)))
	v16 = l1
	v17 = l2
	v18 = v8
	goto L3
L3:
	;
	if v17 == int32(0) {
		v90 = v16
		goto L1
	} else {
		goto L5
	}
L4:
	;
	v90 = v84
	goto L1
L5:
	;
	v24 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18)+4)))
	if v24 != int32(8) {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	v34 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18)+5)))
	if (v34^int32(3))&(v12^int32(3))&int32(255) == int32(0) {
		goto L11
	} else {
		goto L12
	}
L7:
	;
	v30 = F_sweeplist(m, l0, v18+int32(104), int32(-3))
	mBase = m.M
	v33 = m.ExcPending
	if v33 != 0 {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	return int32(0)
L9:
	;
	goto L6
L10:
	;
	v88 = *(*int32)(unsafe.Add(mBase, uint32(v84)))
	if v88 != 0 {
		v16 = v84
		v17 = v17 + int32(-1)
		v18 = v88
		goto L3
	} else {
		goto L29
	}
L11:
	;
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v18)))
	*(*int32)(unsafe.Add(mBase, uint32(v16))) = v49
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v11)+28))
	if v18 != v51 {
		goto L13
	} else {
		goto L14
	}
L12:
	;
	v42 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+20)))
	v47 = v42&int32(3) | v34&int32(248)
	*(*uint8)(unsafe.Add(mBase, uint32(v18)+5)) = uint8(v47)
	v84 = v18
	goto L10
L13:
	;
	v54 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18)+4)))
	switch v54 + int32(-4) {
	case 0:
		goto L16
	case 1:
		goto L18
	case 2:
		goto L20
	case 3:
		goto L15
	case 4:
		goto L17
	case 5:
		goto L21
	case 6:
		goto L19
	default:
		v84 = v16
		goto L10
	}
L14:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+28)) = v49
	goto L13
L15:
	;
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v18)+16))
	v82 = F_luaM_realloc_(m, l0, v18, v78+int32(24), int32(0))
	mBase = m.M
	v83 = m.ExcPending
	if v83 != 0 {
		goto L8
	} else {
		goto L28
	}
L16:
	;
	v67 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v67)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v67)+4)) = v68 + int32(-1)
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v18)+12))
	v76 = F_luaM_realloc_(m, l0, v18, v72+int32(17), int32(0))
	mBase = m.M
	v77 = m.ExcPending
	if v77 != 0 {
		goto L8
	} else {
		goto L27
	}
L17:
	;
	F_luaE_freethread(m, l0, v18)
	mBase = m.M
	v66 = m.ExcPending
	if v66 != 0 {
		goto L8
	} else {
		goto L26
	}
L18:
	;
	F_luaH_free(m, l0, v18)
	mBase = m.M
	v64 = m.ExcPending
	if v64 != 0 {
		goto L8
	} else {
		goto L25
	}
L19:
	;
	F_luaF_freeupval(m, l0, v18)
	mBase = m.M
	v62 = m.ExcPending
	if v62 != 0 {
		goto L8
	} else {
		goto L24
	}
L20:
	;
	F_luaF_freeclosure(m, l0, v18)
	mBase = m.M
	v60 = m.ExcPending
	if v60 != 0 {
		goto L8
	} else {
		goto L23
	}
L21:
	;
	F_luaF_freeproto(m, l0, v18)
	mBase = m.M
	v58 = m.ExcPending
	if v58 != 0 {
		goto L8
	} else {
		goto L22
	}
L22:
	;
	v84 = v16
	goto L10
L23:
	;
	v84 = v16
	goto L10
L24:
	;
	v84 = v16
	goto L10
L25:
	;
	v84 = v16
	goto L10
L26:
	;
	v84 = v16
	goto L10
L27:
	;
	v84 = v16
	goto L10
L28:
	;
	v84 = v16
	goto L10
L29:
	;
	goto L4
}
