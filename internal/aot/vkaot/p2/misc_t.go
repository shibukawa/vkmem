package p2

import (
	base "github.com/shibukawa/vkmem/internal/aot/vkaot/base"
	"unsafe"
)

func F_TerminateModuleForkChild(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v69 int32
	_ = v69
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v9 = int32(-1)
	v11 = *(*int32)(unsafe.Add(mBase, _consts[66]))
	if v11 != int32(4) {
		v69 = v9
		m.G0 = v7 + int32(16)
		return v69
	} else {
		v15 = *(*int32)(unsafe.Add(mBase, _consts[45]))
		if v15 != l0 {
			v69 = v9
			m.G0 = v7 + int32(16)
			return v69
		} else {
			v18 = *(*int32)(unsafe.Add(mBase, _consts[6]))
			if int32(1) < v18 {
				v30 = l0
				v32 = F_kill(m, v30, int32(10))
				mBase = m.M
				v33 = m.ExcPending
				if v33 != 0 {
					return int32(0)
				} else {
					if l1 == int32(0) {
					} else {
						if v32 == int32(-1) {
						} else {
							v39 = *(*int32)(unsafe.Add(mBase, _consts[45]))
							v40 = v39
							for {
								v46 = int32(0)
								v48 = F___syscall_wait4(m, v40, v7+int32(12), v46, v46)
								mBase = m.M
								v49 = F___syscall_ret(m, v48)
								mBase = m.M
								v51 = *(*int32)(unsafe.Add(mBase, _consts[45]))
								if v49 != v51 {
									v40 = v51
									continue
								} else {
									break
								}
								break
							}
						}
					}
					F_resetChildState(m)
					mBase = m.M
					v58 = m.ExcPending
					if v58 != 0 {
						return int32(0)
					} else {
						v59 = int32(0)
						*(*int32)(unsafe.Add(mBase, _consts[580])) = v59
						*(*int32)(unsafe.Add(mBase, _consts[581])) = v59
						v69 = v59
						m.G0 = v7 + int32(16)
						return v69
					}
				}
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v7))) = l0
				F__serverLog(m, int32(1), int32(_a952), v7)
				mBase = m.M
				v27 = m.ExcPending
				if v27 != 0 {
					return int32(0)
				} else {
					v29 = *(*int32)(unsafe.Add(mBase, _consts[45]))
					v30 = v29
					v32 = F_kill(m, v30, int32(10))
					mBase = m.M
					v33 = m.ExcPending
					if v33 != 0 {
						return int32(0)
					} else {
						if l1 == int32(0) {
						} else {
							if v32 == int32(-1) {
							} else {
								v39 = *(*int32)(unsafe.Add(mBase, _consts[45]))
								v40 = v39
								for {
									v46 = int32(0)
									v48 = F___syscall_wait4(m, v40, v7+int32(12), v46, v46)
									mBase = m.M
									v49 = F___syscall_ret(m, v48)
									mBase = m.M
									v51 = *(*int32)(unsafe.Add(mBase, _consts[45]))
									if v49 != v51 {
										v40 = v51
										continue
									} else {
										break
									}
									break
								}
							}
						}
						F_resetChildState(m)
						mBase = m.M
						v58 = m.ExcPending
						if v58 != 0 {
							return int32(0)
						} else {
							v59 = int32(0)
							*(*int32)(unsafe.Add(mBase, _consts[580])) = v59
							*(*int32)(unsafe.Add(mBase, _consts[581])) = v59
							v69 = v59
							m.G0 = v7 + int32(16)
							return v69
						}
					}
				}
			}
		}
	}
}
func F___tan(m *base.Module, l0 float64, l1 float64, l2 int32) float64 {
	var v5 int32
	_ = v5
	var v10 int64
	_ = v10
	var v14 int32
	_ = v14
	var v23 int32
	_ = v23
	var v24 float64
	_ = v24
	var v28 float64
	_ = v28
	var v29 float64
	_ = v29
	var v30 int32
	_ = v30
	var v31 float64
	_ = v31
	var v32 float64
	_ = v32
	var v35 float64
	_ = v35
	var v74 float64
	_ = v74
	var v75 float64
	_ = v75
	var v76 int32
	_ = v76
	var v80 float64
	_ = v80
	var v85 float64
	_ = v85
	var v87 float64
	_ = v87
	var v91 float64
	_ = v91
	var v96 float64
	_ = v96
	var v98 int64
	_ = v98
	var v100 float64
	_ = v100
	var v104 float64
	_ = v104
	var v115 float64
	_ = v115
	v5 = int32(0)
	v10 = base.I64_reinterpret_f64(l0)
	v14 = base.B2i32(base.Ui64(v10&int64(9223372002495037440)) < base.Ui64(int64(4604249089280835585)))
	if v14 == v5 {
		v23 = base.B2i32(int64(-1) < v10)
		if int64(-1) < v10 {
			v24 = l1
		} else {
			v24 = base.F64_neg(l1)
		}
		v28 = base.F64_add(base.F64_sub(float64(0.7853981633974483), base.F64_abs(l0)), base.F64_sub(float64(3.061616997868383e-17), v24))
		v29 = float64(0)
		v30 = v23
	} else {
		v28 = l0
		v29 = l1
		v30 = v5
	}
	v31 = base.F64_mul(v28, v28)
	v32 = base.F64_mul(v28, v31)
	v35 = base.F64_mul(v31, v31)
	v74 = base.F64_add(base.F64_mul(v32, float64(0.3333333333333341)), base.F64_add(base.F64_mul(v31, base.F64_add(base.F64_mul(v32, base.F64_add(base.F64_add(base.F64_mul(v35, base.F64_add(base.F64_mul(v35, base.F64_add(base.F64_mul(v35, base.F64_add(base.F64_mul(v35, base.F64_add(base.F64_mul(v35, float64(-1.8558637485527546e-05)), float64(7.817944429395571e-05))), float64(0.0005880412408202641))), float64(0.0035920791075913124))), float64(0.021869488294859542))), float64(0.13333333333320124)), base.F64_mul(v31, base.F64_add(base.F64_mul(v35, base.F64_add(base.F64_mul(v35, base.F64_add(base.F64_mul(v35, base.F64_add(base.F64_mul(v35, base.F64_add(base.F64_mul(v35, float64(2.590730518636337e-05)), float64(7.140724913826082e-05))), float64(0.0002464631348184699))), float64(0.0014562094543252903))), float64(0.0088632398235993))), float64(0.05396825397622605))))), v29)), v29))
	v75 = base.F64_add(v28, v74)
	if base.Ui64(v10&int64(9223372002495037440)) < base.Ui64(int64(4604249089280835585)) {
		if l2 == int32(0) {
			v115 = v75
		} else {
			v96 = base.F64_div(float64(-1), v75)
			v98 = int64(-4294967296)
			v100 = base.F64_reinterpret_i64(base.I64_reinterpret_f64(v96) & v98)
			v104 = base.F64_reinterpret_i64(base.I64_reinterpret_f64(v75) & v98)
			v115 = base.F64_add(base.F64_mul(v96, base.F64_add(base.F64_mul(v100, base.F64_sub(v74, base.F64_sub(v104, v28))), base.F64_add(base.F64_mul(v100, v104), float64(1)))), v100)
		}
		return v115
	} else {
		v76 = int32(1)
		v80 = base.F64_convert_i32_s(v76 - l2<<(uint(v76)%32))
		v85 = base.F64_add(v28, base.F64_sub(v74, base.F64_div(base.F64_mul(v75, v75), base.F64_add(v75, v80))))
		v87 = base.F64_sub(v80, base.F64_add(v85, v85))
		if v30&v76 != 0 {
			v91 = v87
		} else {
			v91 = base.F64_neg(v87)
		}
		return v91
	}
}
func F___tm_to_tzname(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	F_do_tzset(m)
	mBase = m.M
	return v2
}
func F___trunctfsf2(m *base.Module, l0 int64, l1 int64) float32 {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v15 int64
	_ = v15
	var v19 int64
	_ = v19
	var v20 int32
	_ = v20
	var v27 int32
	_ = v27
	var v31 int64
	_ = v31
	var v32 int64
	_ = v32
	var v36 int32
	_ = v36
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v85 int64
	_ = v85
	var v87 int32
	_ = v87
	var v104 int64
	_ = v104
	var v108 int64
	_ = v108
	var v109 int64
	_ = v109
	var v129 int64
	_ = v129
	var v133 int64
	_ = v133
	var v134 int64
	_ = v134
	var v140 int64
	_ = v140
	var v143 int32
	_ = v143
	var v144 int64
	_ = v144
	var v146 int64
	_ = v146
	var v151 int64
	_ = v151
	var v153 int64
	_ = v153
	var v157 int64
	_ = v157
	var v161 int64
	_ = v161
	var v162 int64
	_ = v162
	var v166 int32
	_ = v166
	var v177 int32
	_ = v177
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	v10 = m.G0
	v12 = v10 - int32(32)
	m.G0 = v12
	v15 = l1 & int64(281474976710655)
	v19 = int64(base.Ui64(l1)>>(uint(int64(48))%64)) & int64(32767)
	v20 = base.I32_wrap_i64(v19)
	if base.Ui32(int32(253)) < base.Ui32(v20+int32(-16257)) {
		if l0|v15 == int64(0) {
			if base.Ui32(v20) <= base.Ui32(int32(16510)) {
				v74 = base.B2i32(v19 == int64(0))
				if v19 == int64(0) {
					v75 = int32(16256)
				} else {
					v75 = int32(16257)
				}
				v76 = v75 - v20
				if v76 <= int32(112) {
					v82 = v12 + int32(16)
					if v19 == int64(0) {
						v85 = v15
					} else {
						v85 = v15 | int64(281474976710656)
					}
					v87 = int32(128) - v76
					if v87&int32(64) == int32(0) {
						if v87 == int32(0) {
							v108 = l0
							v109 = v85
						} else {
							v104 = base.I64_extend_i32_u(v87)
							v108 = l0 << (uint(v104) % 64)
							v109 = int64(base.Ui64(l0)>>(uint(base.I64_extend_i32_u(int32(64)-v87))%64)) | v85<<(uint(v104)%64)
						}
					} else {
						v108 = int64(0)
						v109 = l0 << (uint(base.I64_extend_i32_u(v87+int32(-64))) % 64)
					}
					*(*int64)(unsafe.Add(mBase, uint32(v82))) = v108
					*(*int64)(unsafe.Add(mBase, uint32(v82)+8)) = v109
					if v76&int32(64) == int32(0) {
						if v76 == int32(0) {
							v133 = l0
							v134 = v85
						} else {
							v129 = base.I64_extend_i32_u(v76)
							v133 = v85<<(uint(base.I64_extend_i32_u(int32(64)-v76))%64) | int64(base.Ui64(l0)>>(uint(v129)%64))
							v134 = int64(base.Ui64(v85) >> (uint(v129) % 64))
						}
					} else {
						v133 = int64(base.Ui64(v85) >> (uint(base.I64_extend_i32_u(v76+int32(-64))) % 64))
						v134 = int64(0)
					}
					*(*int64)(unsafe.Add(mBase, uint32(v12))) = v133
					*(*int64)(unsafe.Add(mBase, uint32(v12)+8)) = v134
					v140 = *(*int64)(unsafe.Add(mBase, uint32(v12+int32(8))))
					v143 = base.I32_wrap_i64(int64(base.Ui64(v140) >> (uint(int64(25)) % 64)))
					v144 = *(*int64)(unsafe.Add(mBase, uint32(v12)))
					v146 = *(*int64)(unsafe.Add(mBase, uint32(v12)+16))
					v151 = *(*int64)(unsafe.Add(mBase, uint32(v12+int32(24))))
					v153 = int64(0)
					v157 = v144 | base.I64_extend_i32_u(base.B2i32(v75 != v20)&base.B2i32(v146|v151 != v153))
					v161 = v140 & int64(33554431)
					v162 = int64(16777216)
					if v161 == v162 {
						v166 = base.B2i32(v157 == v153)
					} else {
						v166 = base.B2i32(base.Ui64(v161) < base.Ui64(v162))
					}
					if v166 != 0 {
						if v157|(v161^int64(16777216)) != int64(0) {
							v177 = v143
						} else {
							v177 = v143&int32(1) + v143
						}
					} else {
						v177 = v143 + int32(1)
					}
					v181 = base.B2i32(base.Ui32(int32(8388607)) < base.Ui32(v177))
					if base.Ui32(int32(8388607)) < base.Ui32(v177) {
						v182 = v177 ^ int32(_a0)
					} else {
						v182 = v177
					}
					v185 = v181
					v186 = v182
				} else {
					v79 = int32(0)
					v185 = v79
					v186 = v79
				}
			} else {
				v185 = int32(255)
				v186 = int32(0)
			}
		} else {
			if v19 != int64(32767) {
				if base.Ui32(v20) <= base.Ui32(int32(16510)) {
					v74 = base.B2i32(v19 == int64(0))
					if v19 == int64(0) {
						v75 = int32(16256)
					} else {
						v75 = int32(16257)
					}
					v76 = v75 - v20
					if v76 <= int32(112) {
						v82 = v12 + int32(16)
						if v19 == int64(0) {
							v85 = v15
						} else {
							v85 = v15 | int64(281474976710656)
						}
						v87 = int32(128) - v76
						if v87&int32(64) == int32(0) {
							if v87 == int32(0) {
								v108 = l0
								v109 = v85
							} else {
								v104 = base.I64_extend_i32_u(v87)
								v108 = l0 << (uint(v104) % 64)
								v109 = int64(base.Ui64(l0)>>(uint(base.I64_extend_i32_u(int32(64)-v87))%64)) | v85<<(uint(v104)%64)
							}
						} else {
							v108 = int64(0)
							v109 = l0 << (uint(base.I64_extend_i32_u(v87+int32(-64))) % 64)
						}
						*(*int64)(unsafe.Add(mBase, uint32(v82))) = v108
						*(*int64)(unsafe.Add(mBase, uint32(v82)+8)) = v109
						if v76&int32(64) == int32(0) {
							if v76 == int32(0) {
								v133 = l0
								v134 = v85
							} else {
								v129 = base.I64_extend_i32_u(v76)
								v133 = v85<<(uint(base.I64_extend_i32_u(int32(64)-v76))%64) | int64(base.Ui64(l0)>>(uint(v129)%64))
								v134 = int64(base.Ui64(v85) >> (uint(v129) % 64))
							}
						} else {
							v133 = int64(base.Ui64(v85) >> (uint(base.I64_extend_i32_u(v76+int32(-64))) % 64))
							v134 = int64(0)
						}
						*(*int64)(unsafe.Add(mBase, uint32(v12))) = v133
						*(*int64)(unsafe.Add(mBase, uint32(v12)+8)) = v134
						v140 = *(*int64)(unsafe.Add(mBase, uint32(v12+int32(8))))
						v143 = base.I32_wrap_i64(int64(base.Ui64(v140) >> (uint(int64(25)) % 64)))
						v144 = *(*int64)(unsafe.Add(mBase, uint32(v12)))
						v146 = *(*int64)(unsafe.Add(mBase, uint32(v12)+16))
						v151 = *(*int64)(unsafe.Add(mBase, uint32(v12+int32(24))))
						v153 = int64(0)
						v157 = v144 | base.I64_extend_i32_u(base.B2i32(v75 != v20)&base.B2i32(v146|v151 != v153))
						v161 = v140 & int64(33554431)
						v162 = int64(16777216)
						if v161 == v162 {
							v166 = base.B2i32(v157 == v153)
						} else {
							v166 = base.B2i32(base.Ui64(v161) < base.Ui64(v162))
						}
						if v166 != 0 {
							if v157|(v161^int64(16777216)) != int64(0) {
								v177 = v143
							} else {
								v177 = v143&int32(1) + v143
							}
						} else {
							v177 = v143 + int32(1)
						}
						v181 = base.B2i32(base.Ui32(int32(8388607)) < base.Ui32(v177))
						if base.Ui32(int32(8388607)) < base.Ui32(v177) {
							v182 = v177 ^ int32(_a0)
						} else {
							v182 = v177
						}
						v185 = v181
						v186 = v182
					} else {
						v79 = int32(0)
						v185 = v79
						v186 = v79
					}
				} else {
					v185 = int32(255)
					v186 = int32(0)
				}
			} else {
				v185 = int32(255)
				v186 = base.I32_wrap_i64(int64(base.Ui64(v15)>>(uint(int64(25))%64))) | int32(4194304)
			}
		}
	} else {
		v27 = base.I32_wrap_i64(int64(base.Ui64(v15) >> (uint(int64(25)) % 64)))
		v31 = l1 & int64(33554431)
		v32 = int64(16777216)
		if v31 == v32 {
			v36 = base.B2i32(l0 == int64(0))
		} else {
			v36 = base.B2i32(base.Ui64(v31) < base.Ui64(v32))
		}
		if v36 != 0 {
			if l0|(v31^int64(16777216)) != int64(0) {
				v47 = v27
			} else {
				v47 = v27&int32(1) + v27
			}
		} else {
			v47 = v27 + int32(1)
		}
		v50 = base.B2i32(base.Ui32(int32(8388607)) < base.Ui32(v47))
		if base.Ui32(int32(8388607)) < base.Ui32(v47) {
			v51 = int32(0)
		} else {
			v51 = v47
		}
		if base.Ui32(int32(8388607)) < base.Ui32(v47) {
			v54 = int32(-16255)
		} else {
			v54 = int32(-16256)
		}
		v185 = v54 + v20
		v186 = v51
	}
	m.G0 = v12 + int32(32)
	return base.F32_reinterpret_i32(v185<<(uint(int32(23))%32) | base.I32_wrap_i64(int64(base.Ui64(l1)>>(uint(int64(32))%64)))&int32(-2147483648) | v186)
}
func F___tzset(m *base.Module) {
	F_do_tzset(m)
	return
}
func F_tinsert(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v29 int32
	_ = v29
	var v33 int32
	_ = v33
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v92 int64
	_ = v92
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v102 int32
	_ = v102
	var v110 int32
	_ = v110
	var v114 int32
	_ = v114
	var v117 int32
	_ = v117
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	F_luaL_checktype(m, l0, int32(1), int32(5))
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v12 = F_lua_objlen(m, l0, int32(1))
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v15 = v12 + int32(1)
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	goto L7
L4:
	;
	v117 = m.G3
	v121 = F_luaL_error(m, l0, v117+int32(_a2098), int32(0))
	mBase = m.M
	v122 = m.ExcPending
	if v122 != 0 {
		goto L1
	} else {
		goto L32
	}
L5:
	;
	F_lua_rawseti(m, l0, int32(1), v110)
	mBase = m.M
	v114 = m.ExcPending
	if v114 != 0 {
		goto L1
	} else {
		goto L31
	}
L6:
	;
	v24 = F_luaL_checkinteger(m, l0, int32(2))
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		goto L1
	} else {
		goto L9
	}
L7:
	;
	switch (v16-v17)>>(uint(int32(4))%32) + int32(-2) {
	case 0:
		v110 = v15
		goto L5
	case 1:
		goto L6
	default:
		goto L4
	}
L8:
	;
	v110 = v24
	goto L5
L9:
	;
	if v12 < v24 {
		goto L8
	} else {
		goto L10
	}
L10:
	;
	v29 = v15
	goto L11
L11:
	;
	v33 = v29 + int32(-1)
	goto L16
L12:
	;
	goto L8
L13:
	;
	F_lua_rawseti(m, l0, int32(1), v29)
	mBase = m.M
	v102 = m.ExcPending
	if v102 != 0 {
		goto L1
	} else {
		goto L29
	}
L14:
	;
	v89 = *(*int32)(unsafe.Add(mBase, uint32(v47)))
	v90 = F_luaH_getnum(m, v89, v33)
	mBase = m.M
	v91 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v92 = *(*int64)(unsafe.Add(mBase, uint32(v90)))
	*(*int64)(unsafe.Add(mBase, uint32(v91))) = v92
	v94 = *(*int32)(unsafe.Add(mBase, uint32(v90)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v91)+8)) = v94
	v96 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v96 + int32(16)
	goto L13
L16:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v39 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v44 = v39 + int32(0)
	v45 = m.G398
	if base.Ui32(v44) < base.Ui32(v38) {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	v47 = v44
	goto L19
L18:
	;
	v47 = v45
	goto L19
L19:
	;
	goto L14
L29:
	;
	if v24 < v33 {
		v29 = v33
		goto L11
	} else {
		goto L30
	}
L30:
	;
	goto L12
L31:
	;
	return int32(0)
L32:
	;
	return v121
}
func F_tolower(m *base.Module, l0 int32) int32 {
	var v8 int32
	_ = v8
	if base.Ui32(l0+int32(-65)) < base.Ui32(int32(26)) {
		v8 = l0 | int32(32)
	} else {
		v8 = l0
	}
	return v8
}
func F_top12_2(m *base.Module, l0 float64) int32 {
	return base.I32_wrap_i64(int64(base.Ui64(base.I64_reinterpret_f64(l0)) >> (uint(int64(52)) % 64)))
}
func F_touchCommand(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v32 int64
	_ = v32
	var v36 int32
	_ = v36
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v6 < int32(2) {
		v32 = int64(0)
		goto L1
	} else {
		goto L2
	}
L1:
	;
	F_addReplyLongLong(m, l0, v32)
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
		goto L5
	} else {
		goto L8
	}
L2:
	;
	v13 = int32(1)
	v14 = int32(0)
	goto L3
L3:
	;
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v20 = *(*int32)(unsafe.Add(mBase, uint32(v16+v13<<(uint(int32(2))%32))))
	v21 = F_lookupKeyRead(m, v15, v20)
	mBase = m.M
	v22 = m.ExcPending
	if v22 != 0 {
		goto L5
	} else {
		goto L6
	}
L4:
	;
	v32 = base.I64_extend_i32_u(v25)
	goto L1
L5:
	;
	return
L6:
	;
	v25 = v14 + base.B2i32(v21 != int32(0))
	v27 = v13 + int32(1)
	v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v27 < v28 {
		v13 = v27
		v14 = v25
		goto L3
	} else {
		goto L7
	}
L7:
	;
	goto L4
L8:
	;
	return
}
func F_toupper(m *base.Module, l0 int32) int32 {
	var v8 int32
	_ = v8
	if base.Ui32(l0+int32(-97)) < base.Ui32(int32(26)) {
		v8 = l0 & int32(95)
	} else {
		v8 = l0
	}
	return v8
}
func F_trackBufReferences(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
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
	var v74 int64
	_ = v74
	var v75 int32
	_ = v75
	var v80 int64
	_ = v80
	var v81 int32
	_ = v81
	var v105 int32
	_ = v105
	var v111 int32
	_ = v111
	var v119 int32
	_ = v119
	var v125 int32
	_ = v125
	var v128 int32
	_ = v128
	var v132 int64
	_ = v132
	var v134 int32
	_ = v134
	var v141 int32
	_ = v141
	var v144 int32
	_ = v144
	var v146 int32
	_ = v146
	var v152 int32
	_ = v152
	var v158 int32
	_ = v158
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v179 int32
	_ = v179
	var v185 int32
	_ = v185
	if l2 == int32(0) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	F__serverAssert(m, int32(_a1007), int32(_a977), int32(2951))
	mBase = m.M
	v185 = m.ExcPending
	if v185 != 0 {
		goto L57
	} else {
		goto L59
	}
L2:
	;
	F__serverAssert(m, int32(_a1010), int32(_a977), int32(2918))
	mBase = m.M
	v179 = m.ExcPending
	if v179 != 0 {
		goto L57
	} else {
		goto L58
	}
L3:
	;
	v13 = l0 + l1
	v14 = l0
	goto L4
L4:
	;
	v25 = v14 + int32(12)
	v26 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+10)))
	if v26&int32(1) == int32(0) {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	if v172 != v13 {
		goto L1
	} else {
		goto L56
	}
L6:
	;
	v171 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
	v172 = v25 + v171
	if base.Ui32(v172) < base.Ui32(v13) {
		v14 = v172
		goto L4
	} else {
		goto L55
	}
L7:
	;
	v31 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+11)))
	if v31 != 0 {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v33 = v31
	goto L10
L9:
	;
	v33 = int32(1)
	goto L10
L10:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v14)+11)) = uint8(v33)
	if v31 != 0 {
		goto L6
	} else {
		goto L11
	}
L11:
	;
	v35 = int32(0)
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
	if v36 == v35 {
		v152 = v35
		goto L12
	} else {
		goto L13
	}
L12:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+4)) = v152
	v158 = *(*int32)(unsafe.Add(mBase, uint32(l2)+272))
	*(*int32)(unsafe.Add(mBase, uint32(l2)+272)) = v158 + v152
	goto L6
L13:
	;
	v44 = v35
	v45 = v25
	v46 = v36
	goto L14
L14:
	;
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v45)+4))
	v53 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v50+int32(-1)))))
	switch v53 & int32(7) {
	case 0:
		goto L21
	case 1:
		goto L20
	case 2:
		goto L19
	case 3:
		goto L18
	case 4:
		goto L17
	default:
		v70 = int32(0)
		goto L16
	}
L15:
	;
	v152 = v144
	goto L12
L16:
	;
	v74 = base.I64_extend_i32_u(v70)
	v75 = int32(0)
	if base.Ui64(v74) < base.Ui64(int64(10)) {
		v134 = v75
		goto L23
	} else {
		goto L24
	}
L17:
	;
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v50+int32(-17))))
	v70 = v69
	goto L16
L18:
	;
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v50+int32(-9))))
	v70 = v66
	goto L16
L19:
	;
	v63 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v50+int32(-5)))))
	v70 = v63
	goto L16
L20:
	;
	v60 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v50+int32(-3)))))
	v70 = v60
	goto L16
L21:
	;
	v70 = int32(base.Ui32(v53) >> (uint(int32(3)) % 32))
	goto L16
L22:
	;
	v144 = v44 + v70 + v141 + int32(5)
	v146 = v46 + int32(-8)
	if v146 != 0 {
		v44 = v144
		v45 = v45 + int32(8)
		v46 = v146
		goto L14
	} else {
		goto L54
	}
L23:
	;
	v141 = int32(1) + v134
	goto L22
L24:
	;
	v80 = v74
	v81 = v75
	goto L25
L25:
	;
	if base.Ui64(int64(100)) <= base.Ui64(v80) {
		goto L27
	} else {
		goto L28
	}
L26:
	;
	v134 = v128
	goto L23
L27:
	;
	if base.Ui64(int64(1000)) <= base.Ui64(v80) {
		goto L29
	} else {
		goto L30
	}
L28:
	;
	v141 = int32(2) + v81
	goto L22
L29:
	;
	if base.Ui64(int64(999999999999)) < base.Ui64(v80) {
		goto L31
	} else {
		goto L32
	}
L30:
	;
	v141 = int32(3) + v81
	goto L22
L31:
	;
	v128 = v81 + int32(12)
	v132 = base.I64_div_u_s(v80, int64(1000000000000))
	if base.Ui64(int64(9999999999999)) < base.Ui64(v80) {
		v80 = v132
		v81 = v128
		goto L25
	} else {
		goto L53
	}
L32:
	;
	if base.Ui64(int64(99999999)) < base.Ui64(v80) {
		goto L33
	} else {
		goto L34
	}
L33:
	;
	if base.Ui64(int64(9999999999)) < base.Ui64(v80) {
		goto L45
	} else {
		goto L46
	}
L34:
	;
	if base.Ui64(int64(999999)) < base.Ui64(v80) {
		goto L35
	} else {
		goto L36
	}
L35:
	;
	if base.Ui64(int64(9999999)) < base.Ui64(v80) {
		goto L42
	} else {
		goto L43
	}
L36:
	;
	if base.Ui64(int64(10000)) <= base.Ui64(v80) {
		goto L37
	} else {
		goto L38
	}
L37:
	;
	if base.Ui64(int64(99999)) < base.Ui64(v80) {
		goto L39
	} else {
		goto L40
	}
L38:
	;
	v141 = int32(4) + v81
	goto L22
L39:
	;
	v105 = int32(6)
	goto L41
L40:
	;
	v105 = int32(5)
	goto L41
L41:
	;
	v141 = v105 + v81
	goto L22
L42:
	;
	v111 = int32(8)
	goto L44
L43:
	;
	v111 = int32(7)
	goto L44
L44:
	;
	v141 = v111 + v81
	goto L22
L45:
	;
	if base.Ui64(int64(99999999999)) < base.Ui64(v80) {
		goto L50
	} else {
		goto L51
	}
L46:
	;
	if base.Ui64(int64(999999999)) < base.Ui64(v80) {
		goto L47
	} else {
		goto L48
	}
L47:
	;
	v119 = int32(10)
	goto L49
L48:
	;
	v119 = int32(9)
	goto L49
L49:
	;
	v141 = v119 + v81
	goto L22
L50:
	;
	v125 = int32(12)
	goto L52
L51:
	;
	v125 = int32(11)
	goto L52
L52:
	;
	v141 = v125 + v81
	goto L22
L53:
	;
	goto L26
L54:
	;
	goto L15
L55:
	;
	goto L5
L56:
	;
	return
L57:
	;
	return
L58:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L59:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_trackInstantaneousMetric(m *base.Module, l0 int32, l1 int64, l2 int64, l3 int64) {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v11 int64
	_ = v11
	var v14 int64
	_ = v14
	var v18 int64
	_ = v18
	var v21 int64
	_ = v21
	var v22 int64
	_ = v22
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	var v33 int32
	_ = v33
	v8 = l0 * int32(152)
	v11 = *(*int64)(unsafe.Add(mBase, uint32(v8)+uint32(_consts[799])))
	if v11 < int64(1) {
	} else {
		v14 = l2 - v11
		if int64(1) <= v14 {
			v18 = *(*int64)(unsafe.Add(mBase, uint32(v8)+uint32(_consts[800])))
			v21 = base.I64_div_s((l1-v18)*l3, v14)
			v22 = v21
		} else {
			v22 = int64(0)
		}
		v23 = *(*int32)(unsafe.Add(mBase, uint32(v8)+uint32(_consts[801])))
		v27 = int32(16)
		*(*int64)(unsafe.Add(mBase, uint32(v8+int32(_a528)+v23<<(uint(int32(3))%32)+v27))) = v22
		v33 = base.I32_rem_s(v23+int32(1), v27)
		*(*int32)(unsafe.Add(mBase, uint32(v8)+uint32(_consts[801]))) = v33
	}
	*(*int64)(unsafe.Add(mBase, uint32(v8)+uint32(_consts[800]))) = l1
	*(*int64)(unsafe.Add(mBase, uint32(v8)+uint32(_consts[799]))) = l2
	return
}
func F_tremove(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v82 int64
	_ = v82
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v94 int32
	_ = v94
	var v98 int32
	_ = v98
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v157 int64
	_ = v157
	var v159 int32
	_ = v159
	var v161 int32
	_ = v161
	var v167 int32
	_ = v167
	var v175 int32
	_ = v175
	var v183 int32
	_ = v183
	var v185 int32
	_ = v185
	F_luaL_checktype(m, l0, int32(1), int32(5))
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v12 = int32(0)
	v15 = F_lua_objlen(m, l0, int32(1))
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		goto L1
	} else {
		goto L4
	}
L3:
	;
	return v185
L4:
	;
	v17 = F_luaL_optinteger(m, l0, int32(2), v15)
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		goto L1
	} else {
		goto L5
	}
L5:
	;
	if v17 < int32(1) {
		v185 = v12
		goto L3
	} else {
		goto L6
	}
L6:
	;
	if v15 < v17 {
		v185 = v12
		goto L3
	} else {
		goto L7
	}
L7:
	;
	goto L11
L8:
	;
	if v15 <= v17 {
		goto L24
	} else {
		goto L25
	}
L9:
	;
	v79 = *(*int32)(unsafe.Add(mBase, uint32(v37)))
	v80 = F_luaH_getnum(m, v79, v17)
	mBase = m.M
	v81 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v82 = *(*int64)(unsafe.Add(mBase, uint32(v80)))
	*(*int64)(unsafe.Add(mBase, uint32(v81))) = v82
	v84 = *(*int32)(unsafe.Add(mBase, uint32(v80)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v81)+8)) = v84
	v86 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v86 + int32(16)
	goto L8
L11:
	;
	v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v34 = v29 + int32(0)
	v35 = m.G398
	if base.Ui32(v34) < base.Ui32(v28) {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	v37 = v34
	goto L14
L13:
	;
	v37 = v35
	goto L14
L14:
	;
	goto L9
L24:
	;
	v175 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v175)+8)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v175 + int32(16)
	goto L46
L25:
	;
	v94 = v17
	goto L26
L26:
	;
	v98 = v94 + int32(1)
	goto L31
L27:
	;
	goto L24
L28:
	;
	F_lua_rawseti(m, l0, int32(1), v94)
	mBase = m.M
	v167 = m.ExcPending
	if v167 != 0 {
		goto L1
	} else {
		goto L44
	}
L29:
	;
	v154 = *(*int32)(unsafe.Add(mBase, uint32(v112)))
	v155 = F_luaH_getnum(m, v154, v98)
	mBase = m.M
	v156 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v157 = *(*int64)(unsafe.Add(mBase, uint32(v155)))
	*(*int64)(unsafe.Add(mBase, uint32(v156))) = v157
	v159 = *(*int32)(unsafe.Add(mBase, uint32(v155)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v156)+8)) = v159
	v161 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v161 + int32(16)
	goto L28
L31:
	;
	v103 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v104 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v109 = v104 + int32(0)
	v110 = m.G398
	if base.Ui32(v109) < base.Ui32(v103) {
		goto L32
	} else {
		goto L33
	}
L32:
	;
	v112 = v109
	goto L34
L33:
	;
	v112 = v110
	goto L34
L34:
	;
	goto L29
L44:
	;
	if v98 != v15 {
		v94 = v98
		goto L26
	} else {
		goto L45
	}
L45:
	;
	goto L27
L46:
	;
	F_lua_rawseti(m, l0, int32(1), v15)
	mBase = m.M
	v183 = m.ExcPending
	if v183 != 0 {
		goto L1
	} else {
		goto L47
	}
L47:
	;
	v185 = int32(1)
	goto L3
}
func F_tryOffloadFreeObjToIOThreads(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v49 int64
	_ = v49
	var v53 int32
	_ = v53
	v3 = int32(-1)
	v5 = *(*int32)(unsafe.Add(mBase, _consts[431]))
	if v5 < int32(2) {
		v53 = v3
	} else {
		v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		if base.Ui32(int32(15)) < base.Ui32(v8) {
			v53 = v3
		} else {
			v11 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
			if v11 != 0 {
				v53 = v3
			} else {
				v12 = int32(_a847)
				v18 = *(*int32)(unsafe.Add(mBase, _consts[484]))
				v19 = *(*int32)(unsafe.Add(mBase, _consts[485]))
				v22 = *(*int32)(unsafe.Add(mBase, _consts[439]))
				v26 = v18 + (v19+int32(-1))&v22<<(uint(int32(6))%32)
				v27 = *(*int32)(unsafe.Add(mBase, uint32(v26)))
				if v27 != v22 {
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v26)+4)) = l0 | int32(3)
					v30 = int32(1)
					*(*int32)(unsafe.Add(mBase, uint32(v26))) = v27 + v30
					v33 = *(*int32)(unsafe.Add(mBase, _consts[439]))
					*(*int32)(unsafe.Add(mBase, _consts[439])) = v33 + v30
				}
				if base.B2i32(v27 == v22) == int32(0) {
					v53 = v3
				} else {
					v40 = int32(0)
					v43 = *(*int32)(unsafe.Add(mBase, _consts[434]))
					*(*int32)(unsafe.Add(mBase, _consts[434])) = v43 + int32(1)
					v47 = int32(_a44)
					v49 = *(*int64)(unsafe.Add(mBase, _consts[309]))
					*(*int64)(unsafe.Add(mBase, _consts[309])) = v49 + int64(1)
					v53 = v40
				}
			}
		}
	}
	return v53
}
func F_twoway_strstr(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v21 int64
	_ = v21
	var v32 int32
	_ = v32
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v50 int32
	_ = v50
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v97 int32
	_ = v97
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v115 int32
	_ = v115
	var v117 int32
	_ = v117
	var v127 int32
	_ = v127
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v139 int32
	_ = v139
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
	var v158 int32
	_ = v158
	var v160 int32
	_ = v160
	var v170 int32
	_ = v170
	var v174 int32
	_ = v174
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
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v194 int32
	_ = v194
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v203 int32
	_ = v203
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v222 int32
	_ = v222
	var v224 int32
	_ = v224
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
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
	var v247 int32
	_ = v247
	var v252 int32
	_ = v252
	var v267 int32
	_ = v267
	var v272 int32
	_ = v272
	var v274 int32
	_ = v274
	var v280 int32
	_ = v280
	var v281 int32
	_ = v281
	var v283 int32
	_ = v283
	var v285 int32
	_ = v285
	var v289 int32
	_ = v289
	var v291 int32
	_ = v291
	var v301 int32
	_ = v301
	var v306 int32
	_ = v306
	var v315 int32
	_ = v315
	var v317 int32
	_ = v317
	var v320 int32
	_ = v320
	var v323 int32
	_ = v323
	var v324 int32
	_ = v324
	var v325 int32
	_ = v325
	var v327 int32
	_ = v327
	var v332 int32
	_ = v332
	var v334 int32
	_ = v334
	var v335 int32
	_ = v335
	var v339 int32
	_ = v339
	var v349 int32
	_ = v349
	var v351 int32
	_ = v351
	var v355 int32
	_ = v355
	var v356 int32
	_ = v356
	var v359 int32
	_ = v359
	var v364 int32
	_ = v364
	var v366 int32
	_ = v366
	var v369 int32
	_ = v369
	var v371 int32
	_ = v371
	var v376 int32
	_ = v376
	var v378 int32
	_ = v378
	var v383 int32
	_ = v383
	var v385 int32
	_ = v385
	var v388 int32
	_ = v388
	var v393 int32
	_ = v393
	var v405 int32
	_ = v405
	var v407 int32
	_ = v407
	var v412 int32
	_ = v412
	var v415 int32
	_ = v415
	var v418 int32
	_ = v418
	var v421 int32
	_ = v421
	var v427 int32
	_ = v427
	var v436 int32
	_ = v436
	var v438 int32
	_ = v438
	var v440 int32
	_ = v440
	var v444 int32
	_ = v444
	var v446 int32
	_ = v446
	var v454 int32
	_ = v454
	var v455 int32
	_ = v455
	var v466 int32
	_ = v466
	var v469 int32
	_ = v469
	var v471 int32
	_ = v471
	var v490 int32
	_ = v490
	var v502 int32
	_ = v502
	var v504 int32
	_ = v504
	var v506 int32
	_ = v506
	var v518 int32
	_ = v518
	v3 = int32(0)
	v15 = m.G0
	v17 = v15 - int32(1056)
	m.G0 = v17
	v21 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v17+int32(1048)))) = v21
	*(*int64)(unsafe.Add(mBase, uint32(v17+int32(1040)))) = v21
	*(*int64)(unsafe.Add(mBase, uint32(v17)+1032)) = v21
	*(*int64)(unsafe.Add(mBase, uint32(v17)+1024)) = v21
	v32 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	if v32 != 0 {
		goto L6
	} else {
		goto L7
	}
L1:
	;
	m.G0 = v17 + int32(1056)
	return v518
L2:
	;
	v194 = int32(1)
	v198 = base.B2i32(base.Ui32(v185+v194) < base.Ui32(v187+v194))
	if base.Ui32(v185+v194) < base.Ui32(v187+v194) {
		goto L35
	} else {
		goto L36
	}
L3:
	;
	v97 = int32(1)
	v103 = v97
	v104 = v76
	v105 = v75
	v108 = int32(0)
	v109 = v97
	goto L13
L4:
	;
	v518 = int32(0)
	goto L1
L5:
	;
	v183 = v82
	v185 = v84
	v186 = v85
	v187 = int32(-1)
	v188 = int32(1)
	goto L2
L6:
	;
	v38 = v3
	v39 = v32
	goto L8
L7:
	;
	v82 = v3
	v84 = int32(-1)
	v85 = int32(1)
	goto L5
L8:
	;
	v50 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0+v38))))
	if v50 == int32(0) {
		goto L4
	} else {
		goto L10
	}
L9:
	;
	v75 = int32(1)
	v76 = int32(-1)
	if base.Ui32(v75) < base.Ui32(v59) {
		goto L3
	} else {
		goto L12
	}
L10:
	;
	v58 = int32(1)
	v59 = v38 + v58
	*(*int32)(unsafe.Add(mBase, uint32(v17+v39&int32(255)<<(uint(int32(2))%32)))) = v59
	v67 = v17 + int32(1024) + int32(base.Ui32(v39)>>(uint(int32(3))%32))&int32(28)
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v67)))
	*(*int32)(unsafe.Add(mBase, uint32(v67))) = v68 | v58<<(uint(v39)%32)
	v74 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1+v59))))
	if v74 != 0 {
		v38 = v59
		v39 = v74
		goto L8
	} else {
		goto L11
	}
L11:
	;
	goto L9
L12:
	;
	v82 = v59
	v84 = v76
	v85 = v75
	goto L5
L13:
	;
	v115 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1+v104+v103))))
	v117 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1+v105))))
	if v115 != v117 {
		goto L16
	} else {
		goto L17
	}
L14:
	;
	v139 = int32(1)
	v146 = v139
	v148 = int32(0)
	v149 = int32(-1)
	v150 = v139
	v151 = v139
	goto L23
L15:
	;
	v135 = v131 + v133
	if base.Ui32(v135) < base.Ui32(v59) {
		v103 = v131
		v104 = v132
		v105 = v135
		v108 = v133
		v109 = v134
		goto L13
	} else {
		goto L22
	}
L16:
	;
	if base.Ui32(v115) <= base.Ui32(v117) {
		goto L20
	} else {
		goto L21
	}
L17:
	;
	if v103 != v109 {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	v131 = v103 + int32(1)
	v132 = v104
	v133 = v108
	v134 = v109
	goto L15
L19:
	;
	v131 = int32(1)
	v132 = v104
	v133 = v109 + v108
	v134 = v109
	goto L15
L20:
	;
	v127 = int32(1)
	v131 = v127
	v132 = v108
	v133 = v108 + v127
	v134 = v127
	goto L15
L21:
	;
	v131 = int32(1)
	v132 = v104
	v133 = v105
	v134 = v105 - v104
	goto L15
L22:
	;
	goto L14
L23:
	;
	v158 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1+v149+v146))))
	v160 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1+v151))))
	if v158 != v160 {
		goto L26
	} else {
		goto L27
	}
L24:
	;
	v183 = v59
	v185 = v132
	v186 = v134
	v187 = v176
	v188 = v177
	goto L2
L25:
	;
	v178 = v174 + v175
	if base.Ui32(v178) < base.Ui32(v59) {
		v146 = v174
		v148 = v175
		v149 = v176
		v150 = v177
		v151 = v178
		goto L23
	} else {
		goto L32
	}
L26:
	;
	if base.Ui32(v160) <= base.Ui32(v158) {
		goto L30
	} else {
		goto L31
	}
L27:
	;
	if v146 != v150 {
		goto L28
	} else {
		goto L29
	}
L28:
	;
	v174 = v146 + int32(1)
	v175 = v148
	v176 = v149
	v177 = v150
	goto L25
L29:
	;
	v174 = int32(1)
	v175 = v150 + v148
	v176 = v149
	v177 = v150
	goto L25
L30:
	;
	v170 = int32(1)
	v174 = v170
	v175 = v148 + v170
	v176 = v148
	v177 = v170
	goto L25
L31:
	;
	v174 = int32(1)
	v175 = v151
	v176 = v149
	v177 = v151 - v149
	goto L25
L32:
	;
	goto L24
L33:
	;
	v283 = v183 | int32(63)
	v285 = l0
	v289 = int32(0)
	v291 = l0
	goto L61
L34:
	;
	v280 = v199
	v281 = v183 - v199
	goto L33
L35:
	;
	v199 = v188
	goto L37
L36:
	;
	v199 = v186
	goto L37
L37:
	;
	v200 = l1 + v199
	if base.Ui32(v185+v194) < base.Ui32(v187+v194) {
		goto L38
	} else {
		goto L39
	}
L38:
	;
	v201 = v187
	goto L40
L39:
	;
	v201 = v185
	goto L40
L40:
	;
	v203 = v201 + int32(1)
	if base.Ui32(v203) < base.Ui32(int32(4)) {
		v227 = l1
		v228 = v200
		v229 = v203
		goto L44
	} else {
		goto L45
	}
L41:
	;
	if v267 == int32(0) {
		goto L34
	} else {
		goto L57
	}
L42:
	;
	v267 = int32(0)
	goto L41
L43:
	;
	v239 = v234
	v240 = v235
	v241 = v236
	goto L53
L44:
	;
	if v229 == int32(0) {
		goto L42
	} else {
		goto L51
	}
L45:
	;
	if (v200|l1)&int32(3) != 0 {
		v234 = l1
		v235 = v200
		v236 = v203
		goto L43
	} else {
		goto L46
	}
L46:
	;
	v211 = l1
	v212 = v200
	v213 = v203
	goto L47
L47:
	;
	v216 = *(*int32)(unsafe.Add(mBase, uint32(v211)))
	v217 = *(*int32)(unsafe.Add(mBase, uint32(v212)))
	if v216 != v217 {
		v234 = v211
		v235 = v212
		v236 = v213
		goto L43
	} else {
		goto L49
	}
L48:
	;
	v227 = v222
	v228 = v220
	v229 = v224
	goto L44
L49:
	;
	v219 = int32(4)
	v220 = v212 + v219
	v222 = v211 + v219
	v224 = v213 + int32(-4)
	if base.Ui32(int32(3)) < base.Ui32(v224) {
		v211 = v222
		v212 = v220
		v213 = v224
		goto L47
	} else {
		goto L50
	}
L50:
	;
	goto L48
L51:
	;
	v234 = v227
	v235 = v228
	v236 = v229
	goto L43
L52:
	;
	v267 = v244 - v245
	goto L41
L53:
	;
	v244 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v239))))
	v245 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v240))))
	if v244 != v245 {
		goto L52
	} else {
		goto L55
	}
L55:
	;
	v247 = int32(1)
	v252 = v241 + int32(-1)
	if v252 == int32(0) {
		goto L42
	} else {
		goto L56
	}
L56:
	;
	v239 = v239 + v247
	v240 = v240 + v247
	v241 = v252
	goto L53
L57:
	;
	v272 = v183 + (v201 ^ int32(-1))
	if base.Ui32(v272) < base.Ui32(v201) {
		goto L58
	} else {
		goto L59
	}
L58:
	;
	v274 = v201
	goto L60
L59:
	;
	v274 = v272
	goto L60
L60:
	;
	v280 = v274 + int32(1)
	v281 = int32(0)
	goto L33
L61:
	;
	if base.Ui32(v183) <= base.Ui32(v285-v291) {
		v412 = v285
		goto L63
	} else {
		goto L64
	}
L63:
	;
	v415 = int32(0)
	v418 = v291 + v183
	v421 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v418+int32(-1)))))
	v427 = *(*int32)(unsafe.Add(mBase, uint32(v17+int32(1024)+int32(base.Ui32(v421)>>(uint(int32(3))%32))&int32(28))))
	if int32(base.Ui32(v427)>>(uint(v421)%32))&int32(1) == v415 {
		v285 = v412
		v289 = v415
		v291 = v418
		goto L61
	} else {
		goto L95
	}
L64:
	;
	v301 = int32(0)
	v306 = base.B2i32(v283 != v301)
	if v285&int32(3) == v301 {
		v332 = v285
		v334 = v283
		v335 = v306
		goto L68
	} else {
		goto L69
	}
L65:
	;
	if v405 != 0 {
		goto L90
	} else {
		goto L91
	}
L66:
	;
	v405 = int32(0)
	goto L65
L67:
	;
	v383 = v376
	v385 = v378
	goto L85
L68:
	;
	if v335 == int32(0) {
		goto L66
	} else {
		goto L76
	}
L69:
	;
	if v283 == int32(0) {
		v332 = v285
		v334 = v283
		v335 = v306
		goto L68
	} else {
		goto L70
	}
L70:
	;
	v315 = v285
	v317 = v283
	goto L71
L71:
	;
	v320 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v315))))
	if v320 == int32(0) {
		v376 = v315
		v378 = v317
		goto L67
	} else {
		goto L73
	}
L72:
	;
	v332 = v327
	v334 = v323
	v335 = v325
	goto L68
L73:
	;
	v323 = v317 + int32(-1)
	v324 = int32(0)
	v325 = base.B2i32(v323 != v324)
	v327 = v315 + int32(1)
	if v327&int32(3) == v324 {
		v332 = v327
		v334 = v323
		v335 = v325
		goto L68
	} else {
		goto L74
	}
L74:
	;
	if v323 != 0 {
		v315 = v327
		v317 = v323
		goto L71
	} else {
		goto L75
	}
L75:
	;
	goto L72
L76:
	;
	v339 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v332))))
	if v339 == int32(0) {
		v369 = v332
		v371 = v334
		goto L77
	} else {
		goto L78
	}
L77:
	;
	if v371 == int32(0) {
		goto L66
	} else {
		goto L84
	}
L78:
	;
	if base.Ui32(v334) < base.Ui32(int32(4)) {
		v369 = v332
		v371 = v334
		goto L77
	} else {
		goto L79
	}
L79:
	;
	v349 = v332
	v351 = v334
	goto L80
L80:
	;
	v355 = *(*int32)(unsafe.Add(mBase, uint32(v349)))
	v356 = v355 ^ int32(0)
	v359 = int32(-2139062144)
	if (int32(16843008)-v356|v356)&v359 != v359 {
		v376 = v349
		v378 = v351
		goto L67
	} else {
		goto L82
	}
L81:
	;
	v369 = v364
	v371 = v366
	goto L77
L82:
	;
	v364 = v349 + int32(4)
	v366 = v351 + int32(-4)
	if base.Ui32(int32(3)) < base.Ui32(v366) {
		v349 = v364
		v351 = v366
		goto L80
	} else {
		goto L83
	}
L83:
	;
	goto L81
L84:
	;
	v376 = v369
	v378 = v371
	goto L67
L85:
	;
	v388 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v383))))
	if v388 != int32(0) {
		goto L87
	} else {
		goto L88
	}
L86:
	;
	goto L66
L87:
	;
	v393 = v385 + int32(-1)
	if v393 != 0 {
		v383 = v383 + int32(1)
		v385 = v393
		goto L85
	} else {
		goto L89
	}
L88:
	;
	v405 = v383
	goto L65
L89:
	;
	goto L86
L90:
	;
	v407 = v405
	goto L92
L91:
	;
	v407 = v285 + v283
	goto L92
L92:
	;
	if v405 == int32(0) {
		v412 = v407
		goto L63
	} else {
		goto L93
	}
L93:
	;
	if base.Ui32(v405-v291) < base.Ui32(v183) {
		v518 = v301
		goto L1
	} else {
		goto L94
	}
L94:
	;
	v412 = v407
	goto L63
L95:
	;
	v436 = *(*int32)(unsafe.Add(mBase, uint32(v17+v421<<(uint(int32(2))%32))))
	if v183 == v436 {
		goto L96
	} else {
		goto L97
	}
L96:
	;
	if base.Ui32(v289) < base.Ui32(v203) {
		goto L103
	} else {
		goto L104
	}
L97:
	;
	v438 = v183 - v436
	if base.Ui32(v289) < base.Ui32(v438) {
		goto L98
	} else {
		goto L99
	}
L98:
	;
	v440 = v438
	goto L100
L99:
	;
	v440 = v289
	goto L100
L100:
	;
	v285 = v412
	v289 = int32(0)
	v291 = v291 + v440
	goto L61
L101:
	;
	v285 = v412
	v289 = int32(0)
	v291 = v291 + (v455 - v201)
	goto L61
L102:
	;
	v490 = v203
	goto L111
L103:
	;
	v444 = v203
	goto L105
L104:
	;
	v444 = v289
	goto L105
L105:
	;
	v446 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1+v444))))
	if v446 == int32(0) {
		goto L102
	} else {
		goto L106
	}
L106:
	;
	v454 = v446
	v455 = v444
	goto L107
L107:
	;
	v466 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v291+v455))))
	if v454&int32(255) != v466 {
		goto L101
	} else {
		goto L109
	}
L108:
	;
	goto L102
L109:
	;
	v469 = v455 + int32(1)
	v471 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1+v469))))
	if v471 != 0 {
		v454 = v471
		v455 = v469
		goto L107
	} else {
		goto L110
	}
L110:
	;
	goto L108
L111:
	;
	if base.Ui32(v289) < base.Ui32(v490) {
		goto L113
	} else {
		goto L114
	}
L112:
	;
	v285 = v412
	v289 = v281
	v291 = v291 + v280
	goto L61
L113:
	;
	v502 = v490 + int32(-1)
	v504 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1+v502))))
	v506 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v291+v502))))
	if v504 == v506 {
		v490 = v502
		goto L111
	} else {
		goto L115
	}
L114:
	;
	v518 = v291
	goto L1
L115:
	;
	goto L112
}
