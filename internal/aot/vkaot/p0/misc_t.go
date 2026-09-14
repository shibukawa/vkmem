package p0

import (
	base "github.com/shibukawa/vkmem/internal/aot/vkaot/base"
	"math"
	"unsafe"
)

func F_ThreadsManager_init(m *base.Module) {
	return
}
func F___time(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v4 float64
	_ = v4
	var v6 float64
	_ = v6
	var v12 int64
	_ = v12
	var v14 int64
	_ = v14
	v4 = m.Env.Emscripten_date_now(m)
	mBase = m.M
	v6 = base.F64_div(v4, float64(1000))
	if base.F64_lt(base.F64_abs(v6), float64(9.223372036854776e+18)) == int32(0) {
		v14 = int64(-9223372036854775807 - 1)
	} else {
		v12 = base.I64_trunc_f64_s(v6)
		v14 = v12
	}
	if l0 == int32(0) {
	} else {
		*(*int64)(unsafe.Add(mBase, uint32(l0))) = v14
	}
	return v14
}
func F___tm_to_secs(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v14 int64
	_ = v14
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v27 int32
	_ = v27
	var v33 int64
	_ = v33
	var v34 int32
	_ = v34
	var v38 int32
	_ = v38
	var v49 int32
	_ = v49
	var v53 int32
	_ = v53
	var v57 int32
	_ = v57
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v78 int64
	_ = v78
	var v79 int64
	_ = v79
	var v80 int64
	_ = v80
	var v83 int64
	_ = v83
	var v89 int32
	_ = v89
	var v94 int32
	_ = v94
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
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
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v155 int64
	_ = v155
	var v156 int32
	_ = v156
	var v162 int32
	_ = v162
	var v165 int32
	_ = v165
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v170 int64
	_ = v170
	var v171 int64
	_ = v171
	var v172 int64
	_ = v172
	v10 = m.G0
	v12 = v10 - int32(16)
	m.G0 = v12
	v14 = int64(*(*int32)(unsafe.Add(mBase, uint32(l0)+20)))
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if base.Ui32(v15) < base.Ui32(int32(12)) {
		v33 = v14
		v34 = v15
	} else {
		v18 = int32(12)
		v19 = base.I32_div_s(v15, v18)
		v22 = v15 - v19*v18
		if v22 < int32(0) {
			v27 = v22 + v18
		} else {
			v27 = v22
		}
		v33 = base.I64_extend_i32_s(v19+v22>>(uint(int32(31))%32)) + v14
		v34 = v27
	}
	v38 = v12 + int32(12)
	if base.Ui64(int64(136)) < base.Ui64(v33+int64(-2)) {
		v78 = v33 + int64(-100)
		v79 = int64(400)
		v80 = base.I64_div_s(v78, v79)
		v83 = v78 - v80*v79
		v89 = base.I32_wrap_i64(v83)
		if v83 < int64(0) {
			v94 = v89 + int32(400)
		} else {
			v94 = v89
		}
		if v94 != 0 {
			if v94 < int32(200) {
				v110 = base.B2i32(int32(99) < v94)
				if int32(99) < v94 {
					v111 = v94 + int32(-100)
				} else {
					v111 = v94
				}
				v112 = v111
				v113 = v110
			} else {
				if base.Ui32(v94) < base.Ui32(int32(300)) {
					v112 = v94 + int32(-200)
					v113 = int32(2)
				} else {
					v112 = v94 + int32(-300)
					v113 = int32(3)
				}
			}
			if v112 != 0 {
				v119 = int32(base.Ui32(v112) >> (uint(int32(2)) % 32))
				v122 = int32(0)
				v123 = base.B2i32(v112&int32(3) == v122)
				if v38 == v122 {
					v130 = v123
					v131 = v113
					v132 = v119
				} else {
					v126 = v123
					v127 = v113
					v128 = v119
					*(*int32)(unsafe.Add(mBase, uint32(v38))) = v126
					v130 = v126
					v131 = v127
					v132 = v128
				}
			} else {
				v115 = int32(0)
				v116 = v113
				v117 = int32(0)
				if v38 != 0 {
					v126 = v115
					v127 = v116
					v128 = v117
					*(*int32)(unsafe.Add(mBase, uint32(v38))) = v126
					v130 = v126
					v131 = v127
					v132 = v128
				} else {
					v130 = v115
					v131 = v116
					v132 = v117
				}
			}
		} else {
			v115 = int32(1)
			v116 = int32(0)
			v117 = int32(0)
			if v38 != 0 {
				v126 = v115
				v127 = v116
				v128 = v117
				*(*int32)(unsafe.Add(mBase, uint32(v38))) = v126
				v130 = v126
				v131 = v127
				v132 = v128
			} else {
				v130 = v115
				v131 = v116
				v132 = v117
			}
		}
		v155 = v78*int64(31536000) + base.I64_extend_i32_s(v132+(v131*int32(24)+(base.I32_wrap_i64(v83>>(uint(int64(63))%64))+base.I32_wrap_i64(v80))*int32(97))-v130)*int64(86400) + int64(946771200)
	} else {
		v49 = base.I32_wrap_i64(v33)
		v53 = (v49 + int32(-68)) >> (uint(int32(2)) % 32)
		if v49&int32(3) != 0 {
			if v38 == int32(0) {
				v67 = v53
			} else {
				v64 = v53
				v65 = int32(0)
				*(*int32)(unsafe.Add(mBase, uint32(v38))) = v65
				v67 = v64
			}
		} else {
			v57 = v53 + int32(-1)
			if v38 == int32(0) {
				v67 = v57
			} else {
				v64 = v57
				v65 = int32(1)
				*(*int32)(unsafe.Add(mBase, uint32(v38))) = v65
				v67 = v64
			}
		}
		v155 = base.I64_extend_i32_s(v49*int32(31536000) + v67*int32(86400) + int32(2087447296))
	}
	v156 = *(*int32)(unsafe.Add(mBase, uint32(v12)+12))
	v162 = *(*int32)(unsafe.Add(mBase, uint32(v34<<(uint(int32(2))%32))+uint32(_c_F___tm_to_secs[0])))
	if v156 != 0 {
		v165 = v162 + int32(86400)
	} else {
		v165 = v162
	}
	if int32(1) < v34 {
		v168 = v165
	} else {
		v168 = v162
	}
	v169 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v170 = int64(*(*int32)(unsafe.Add(mBase, uint32(l0)+8)))
	v171 = int64(*(*int32)(unsafe.Add(mBase, uint32(l0)+4)))
	v172 = int64(*(*int32)(unsafe.Add(mBase, uint32(l0))))
	m.G0 = v12 + int32(16)
	return v172 + (v155 + base.I64_extend_i32_s(v168) + base.I64_extend_i32_s(v169+int32(-1))*int64(86400) + v170*int64(3600) + v171*int64(60))
}
func F___trunctfdf2(m *base.Module, l0 int64, l1 int64) float64 {
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
	var v29 int64
	_ = v29
	var v34 int64
	_ = v34
	var v44 int64
	_ = v44
	var v47 int32
	_ = v47
	var v48 int64
	_ = v48
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v76 int64
	_ = v76
	var v79 int32
	_ = v79
	var v82 int64
	_ = v82
	var v84 int32
	_ = v84
	var v101 int64
	_ = v101
	var v105 int64
	_ = v105
	var v106 int64
	_ = v106
	var v126 int64
	_ = v126
	var v130 int64
	_ = v130
	var v131 int64
	_ = v131
	var v135 int64
	_ = v135
	var v140 int64
	_ = v140
	var v143 int64
	_ = v143
	var v147 int64
	_ = v147
	var v152 int64
	_ = v152
	var v158 int64
	_ = v158
	var v168 int64
	_ = v168
	var v172 int32
	_ = v172
	var v173 int64
	_ = v173
	var v175 int64
	_ = v175
	var v176 int64
	_ = v176
	v10 = m.G0
	v12 = v10 - int32(32)
	m.G0 = v12
	v15 = l1 & int64(281474976710655)
	v19 = int64(base.Ui64(l1)>>(uint(int64(48))%64)) & int64(32767)
	v20 = base.I32_wrap_i64(v19)
	if base.Ui32(int32(2045)) < base.Ui32(v20+int32(-15361)) {
		if l0|v15 == int64(0) {
			if base.Ui32(v20) <= base.Ui32(int32(17406)) {
				v71 = base.B2i32(v19 == int64(0))
				if v19 == int64(0) {
					v72 = int32(15360)
				} else {
					v72 = int32(15361)
				}
				v73 = v72 - v20
				if v73 <= int32(112) {
					v79 = v12 + int32(16)
					if v19 == int64(0) {
						v82 = v15
					} else {
						v82 = v15 | int64(281474976710656)
					}
					v84 = int32(128) - v73
					if v84&int32(64) == int32(0) {
						if v84 == int32(0) {
							v105 = l0
							v106 = v82
						} else {
							v101 = base.I64_extend_i32_u(v84)
							v105 = l0 << (uint(v101) % 64)
							v106 = int64(base.Ui64(l0)>>(uint(base.I64_extend_i32_u(int32(64)-v84))%64)) | v82<<(uint(v101)%64)
						}
					} else {
						v105 = int64(0)
						v106 = l0 << (uint(base.I64_extend_i32_u(v84+int32(-64))) % 64)
					}
					*(*int64)(unsafe.Add(mBase, uint32(v79))) = v105
					*(*int64)(unsafe.Add(mBase, uint32(v79)+8)) = v106
					if v73&int32(64) == int32(0) {
						if v73 == int32(0) {
							v130 = l0
							v131 = v82
						} else {
							v126 = base.I64_extend_i32_u(v73)
							v130 = v82<<(uint(base.I64_extend_i32_u(int32(64)-v73))%64) | int64(base.Ui64(l0)>>(uint(v126)%64))
							v131 = int64(base.Ui64(v82) >> (uint(v126) % 64))
						}
					} else {
						v130 = int64(base.Ui64(v82) >> (uint(base.I64_extend_i32_u(v73+int32(-64))) % 64))
						v131 = int64(0)
					}
					*(*int64)(unsafe.Add(mBase, uint32(v12))) = v130
					*(*int64)(unsafe.Add(mBase, uint32(v12)+8)) = v131
					v135 = *(*int64)(unsafe.Add(mBase, uint32(v12)))
					v140 = *(*int64)(unsafe.Add(mBase, uint32(v12+int32(8))))
					v143 = int64(base.Ui64(v135)>>(uint(int64(60))%64)) | v140<<(uint(int64(4))%64)
					v147 = *(*int64)(unsafe.Add(mBase, uint32(v12)+16))
					v152 = *(*int64)(unsafe.Add(mBase, uint32(v12+int32(24))))
					v158 = v135&int64(1152921504606846975) | base.I64_extend_i32_u(base.B2i32(v72 != v20)&base.B2i32(v147|v152 != int64(0)))
					if base.Ui64(v158) < base.Ui64(int64(576460752303423489)) {
						if v158 != int64(576460752303423488) {
							v168 = v143
						} else {
							v168 = v143&int64(1) + v143
						}
					} else {
						v168 = v143 + int64(1)
					}
					v172 = base.B2i32(base.Ui64(int64(4503599627370495)) < base.Ui64(v168))
					if base.Ui64(int64(4503599627370495)) < base.Ui64(v168) {
						v173 = v168 ^ int64(4503599627370496)
					} else {
						v173 = v168
					}
					v175 = v173
					v176 = base.I64_extend_i32_u(v172)
				} else {
					v76 = int64(0)
					v175 = v76
					v176 = v76
				}
			} else {
				v175 = int64(0)
				v176 = int64(2047)
			}
		} else {
			if v19 != int64(32767) {
				if base.Ui32(v20) <= base.Ui32(int32(17406)) {
					v71 = base.B2i32(v19 == int64(0))
					if v19 == int64(0) {
						v72 = int32(15360)
					} else {
						v72 = int32(15361)
					}
					v73 = v72 - v20
					if v73 <= int32(112) {
						v79 = v12 + int32(16)
						if v19 == int64(0) {
							v82 = v15
						} else {
							v82 = v15 | int64(281474976710656)
						}
						v84 = int32(128) - v73
						if v84&int32(64) == int32(0) {
							if v84 == int32(0) {
								v105 = l0
								v106 = v82
							} else {
								v101 = base.I64_extend_i32_u(v84)
								v105 = l0 << (uint(v101) % 64)
								v106 = int64(base.Ui64(l0)>>(uint(base.I64_extend_i32_u(int32(64)-v84))%64)) | v82<<(uint(v101)%64)
							}
						} else {
							v105 = int64(0)
							v106 = l0 << (uint(base.I64_extend_i32_u(v84+int32(-64))) % 64)
						}
						*(*int64)(unsafe.Add(mBase, uint32(v79))) = v105
						*(*int64)(unsafe.Add(mBase, uint32(v79)+8)) = v106
						if v73&int32(64) == int32(0) {
							if v73 == int32(0) {
								v130 = l0
								v131 = v82
							} else {
								v126 = base.I64_extend_i32_u(v73)
								v130 = v82<<(uint(base.I64_extend_i32_u(int32(64)-v73))%64) | int64(base.Ui64(l0)>>(uint(v126)%64))
								v131 = int64(base.Ui64(v82) >> (uint(v126) % 64))
							}
						} else {
							v130 = int64(base.Ui64(v82) >> (uint(base.I64_extend_i32_u(v73+int32(-64))) % 64))
							v131 = int64(0)
						}
						*(*int64)(unsafe.Add(mBase, uint32(v12))) = v130
						*(*int64)(unsafe.Add(mBase, uint32(v12)+8)) = v131
						v135 = *(*int64)(unsafe.Add(mBase, uint32(v12)))
						v140 = *(*int64)(unsafe.Add(mBase, uint32(v12+int32(8))))
						v143 = int64(base.Ui64(v135)>>(uint(int64(60))%64)) | v140<<(uint(int64(4))%64)
						v147 = *(*int64)(unsafe.Add(mBase, uint32(v12)+16))
						v152 = *(*int64)(unsafe.Add(mBase, uint32(v12+int32(24))))
						v158 = v135&int64(1152921504606846975) | base.I64_extend_i32_u(base.B2i32(v72 != v20)&base.B2i32(v147|v152 != int64(0)))
						if base.Ui64(v158) < base.Ui64(int64(576460752303423489)) {
							if v158 != int64(576460752303423488) {
								v168 = v143
							} else {
								v168 = v143&int64(1) + v143
							}
						} else {
							v168 = v143 + int64(1)
						}
						v172 = base.B2i32(base.Ui64(int64(4503599627370495)) < base.Ui64(v168))
						if base.Ui64(int64(4503599627370495)) < base.Ui64(v168) {
							v173 = v168 ^ int64(4503599627370496)
						} else {
							v173 = v168
						}
						v175 = v173
						v176 = base.I64_extend_i32_u(v172)
					} else {
						v76 = int64(0)
						v175 = v76
						v176 = v76
					}
				} else {
					v175 = int64(0)
					v176 = int64(2047)
				}
			} else {
				v175 = int64(base.Ui64(l0)>>(uint(int64(60))%64)) | v15<<(uint(int64(4))%64) | int64(2251799813685248)
				v176 = int64(2047)
			}
		}
	} else {
		v29 = int64(base.Ui64(l0)>>(uint(int64(60))%64)) | v15<<(uint(int64(4))%64)
		v34 = l0 & int64(1152921504606846975)
		if base.Ui64(v34) < base.Ui64(int64(576460752303423489)) {
			if v34 != int64(576460752303423488) {
				v44 = v29
			} else {
				v44 = v29&int64(1) + v29
			}
		} else {
			v44 = v29 + int64(1)
		}
		v47 = base.B2i32(base.Ui64(int64(4503599627370495)) < base.Ui64(v44))
		if base.Ui64(int64(4503599627370495)) < base.Ui64(v44) {
			v48 = int64(0)
		} else {
			v48 = v44
		}
		v175 = v48
		v176 = base.I64_extend_i32_u(v47) + base.I64_extend_i32_u(v20+int32(-15360))
	}
	m.G0 = v12 + int32(32)
	return base.F64_reinterpret_i64(v176<<(uint(int64(52))%64) | l1&int64(-9223372036854775807-1) | v175)
}
func F_table_is_an_array(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v63 int32
	_ = v63
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v124 int32
	_ = v124
	var v127 int32
	_ = v127
	var v130 int32
	_ = v130
	var v134 float64
	_ = v134
	var v135 int32
	_ = v135
	var v136 int64
	_ = v136
	var v138 int64
	_ = v138
	var v139 int64
	_ = v139
	var v146 int32
	_ = v146
	var v163 int32
	_ = v163
	var v165 int32
	_ = v165
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v179 int32
	_ = v179
	var v183 int32
	_ = v183
	var v187 int32
	_ = v187
	var v189 int32
	_ = v189
	var v196 int32
	_ = v196
	var v201 float64
	_ = v201
	var v203 float64
	_ = v203
	var v209 int32
	_ = v209
	var v211 int32
	_ = v211
	var v213 int32
	_ = v213
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v218 int32
	_ = v218
	var v220 int32
	_ = v220
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v233 int32
	_ = v233
	var v237 int32
	_ = v237
	var v241 int32
	_ = v241
	var v243 int32
	_ = v243
	var v250 int32
	_ = v250
	v10 = m.G3
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v15 = (v11 - v12) >> (uint(int32(4)) % 32)
	goto L1
L1:
	;
	F_luaL_checkstack(m, l0, int32(2), v10+int32(_a_F_table_is_an_array_0))
	mBase = m.M
	v22 = m.ExcPending
	if v22 != 0 {
		goto L2
	} else {
		goto L3
	}
L2:
	;
	return int32(0)
L3:
	;
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v24)+8)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v24 + int32(16)
	goto L4
L4:
	;
	v30 = int32(0)
	v32 = F_lua_next(m, l0, int32(-2))
	mBase = m.M
	v33 = m.ExcPending
	if v33 != 0 {
		goto L2
	} else {
		goto L7
	}
L5:
	;
	if v15 < int32(0) {
		goto L61
	} else {
		goto L62
	}
L6:
	;
	v37 = v30
	v39 = int32(0)
	goto L9
L7:
	;
	if v32 != 0 {
		goto L6
	} else {
		goto L8
	}
L8:
	;
	v218 = v30
	v220 = int32(0)
	goto L5
L9:
	;
	goto L13
L10:
	;
	v218 = v213
	v220 = v211
	goto L5
L11:
	;
	goto L23
L12:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v63 + int32(-16)
	goto L11
L13:
	;
	v63 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	goto L12
L19:
	;
	v201 = base.F64_convert_i32_s(v39)
	if base.F64_gt(v134, v201) != 0 {
		goto L53
	} else {
		goto L54
	}
L20:
	;
	if v15 < int32(0) {
		goto L45
	} else {
		goto L46
	}
L21:
	;
	if v130 != int32(3) {
		goto L20
	} else {
		goto L36
	}
L22:
	;
	v124 = m.G398
	if v90 != v124 {
		goto L34
	} else {
		goto L35
	}
L23:
	;
	goto L27
L27:
	;
	v87 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v90 = v87 + int32(-16)
	goto L22
L34:
	;
	v127 = *(*int32)(unsafe.Add(mBase, uint32(v90)+8))
	v130 = v127
	goto L21
L35:
	;
	v130 = int32(-1)
	goto L21
L36:
	;
	v134 = F_lua_tonumber(m, l0, int32(-1))
	mBase = m.M
	v135 = m.ExcPending
	if v135 != 0 {
		goto L2
	} else {
		goto L37
	}
L37:
	;
	v136 = base.I64_reinterpret_f64(v134)
	v138 = v136 & int64(9223372036854775807)
	v139 = int64(0)
	v146 = base.B2i32(v136 < v139)
	if base.B2i32(v138 == v139)|base.B2i32(base.Ui64(v138+int64(-1)) < base.Ui64(int64(4503599627370495)))&v146|base.B2i32(v138 == int64(9218868437227405312))|base.B2i32(base.Ui64(v138+int64(-4503599627370496)) < base.Ui64(int64(9214364837600034816)))&v146 != 0 {
		goto L20
	} else {
		goto L38
	}
L38:
	;
	if base.F64_lt(base.F64_abs(v134), float64(2.147483648e+09)) == int32(0) {
		goto L40
	} else {
		goto L41
	}
L39:
	;
	if base.F64_eq(v134, base.F64_convert_i32_s(v165)) != 0 {
		goto L19
	} else {
		goto L42
	}
L40:
	;
	v165 = int32(-2147483648)
	goto L39
L41:
	;
	v163 = base.I32_trunc_f64_s(v134)
	v165 = v163
	goto L39
L42:
	;
	goto L20
L43:
	;
	return int32(0)
L44:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v196
	goto L43
L45:
	;
	v189 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v196 = v189 + v15<<(uint(int32(4))%32) + int32(16)
	goto L44
L46:
	;
	v175 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v176 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v179 = v176 + v15<<(uint(int32(4))%32)
	if base.Ui32(v179) <= base.Ui32(v175) {
		v196 = v179
		goto L44
	} else {
		goto L47
	}
L47:
	;
	v183 = v175
	goto L48
L48:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v183)+8)) = int32(0)
	v187 = v183 + int32(16)
	if base.Ui32(v187) < base.Ui32(v179) {
		v183 = v187
		goto L48
	} else {
		goto L50
	}
L50:
	;
	v196 = v179
	goto L44
L51:
	;
	v213 = v37 + int32(1)
	v215 = F_lua_next(m, l0, int32(-2))
	mBase = m.M
	v216 = m.ExcPending
	if v216 != 0 {
		goto L2
	} else {
		goto L57
	}
L52:
	;
	v211 = int32(-2147483648)
	goto L51
L53:
	;
	v203 = v134
	goto L55
L54:
	;
	v203 = v201
	goto L55
L55:
	;
	if base.F64_lt(base.F64_abs(v203), float64(2.147483648e+09)) == int32(0) {
		goto L52
	} else {
		goto L56
	}
L56:
	;
	v209 = base.I32_trunc_f64_s(v203)
	v211 = v209
	goto L51
L57:
	;
	if v215 != 0 {
		v37 = v213
		v39 = v211
		goto L9
	} else {
		goto L58
	}
L58:
	;
	goto L10
L59:
	;
	return base.B2i32(v220 == v218)
L60:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v250
	goto L59
L61:
	;
	v243 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v250 = v243 + v15<<(uint(int32(4))%32) + int32(16)
	goto L60
L62:
	;
	v229 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v230 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v233 = v230 + v15<<(uint(int32(4))%32)
	if base.Ui32(v233) <= base.Ui32(v229) {
		v250 = v233
		goto L60
	} else {
		goto L63
	}
L63:
	;
	v237 = v229
	goto L64
L64:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v237)+8)) = int32(0)
	v241 = v237 + int32(16)
	if base.Ui32(v241) < base.Ui32(v233) {
		v237 = v241
		goto L64
	} else {
		goto L66
	}
L66:
	;
	v250 = v233
	goto L60
}
func F_tan(m *base.Module, l0 float64) float64 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v13 int32
	_ = v13
	var v18 float64
	_ = v18
	var v19 int32
	_ = v19
	var v26 int64
	_ = v26
	var v30 int32
	_ = v30
	var v39 int32
	_ = v39
	var v40 float64
	_ = v40
	var v44 float64
	_ = v44
	var v45 float64
	_ = v45
	var v46 int32
	_ = v46
	var v47 float64
	_ = v47
	var v48 float64
	_ = v48
	var v51 float64
	_ = v51
	var v90 float64
	_ = v90
	var v91 float64
	_ = v91
	var v92 int32
	_ = v92
	var v96 float64
	_ = v96
	var v101 float64
	_ = v101
	var v103 float64
	_ = v103
	var v107 float64
	_ = v107
	var v133 float64
	_ = v133
	var v137 int32
	_ = v137
	var v138 float64
	_ = v138
	var v139 float64
	_ = v139
	var v141 int32
	_ = v141
	var v143 int32
	_ = v143
	var v148 int64
	_ = v148
	var v152 int32
	_ = v152
	var v161 int32
	_ = v161
	var v162 float64
	_ = v162
	var v166 float64
	_ = v166
	var v167 float64
	_ = v167
	var v168 int32
	_ = v168
	var v169 float64
	_ = v169
	var v170 float64
	_ = v170
	var v173 float64
	_ = v173
	var v212 float64
	_ = v212
	var v213 float64
	_ = v213
	var v214 int32
	_ = v214
	var v218 float64
	_ = v218
	var v223 float64
	_ = v223
	var v225 float64
	_ = v225
	var v229 float64
	_ = v229
	var v233 float64
	_ = v233
	var v235 int64
	_ = v235
	var v237 float64
	_ = v237
	var v241 float64
	_ = v241
	var v252 float64
	_ = v252
	var v255 float64
	_ = v255
	var v256 float64
	_ = v256
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	v13 = base.I32_wrap_i64(int64(base.Ui64(base.I64_reinterpret_f64(l0))>>(uint(int64(32))%64))) & int32(2147483647)
	if base.Ui32(int32(1072243195)) < base.Ui32(v13) {
		if base.Ui32(v13) < base.Ui32(int32(2146435072)) {
			v137 = F___rem_pio2(m, l0, v6)
			mBase = m.M
			v138 = *(*float64)(unsafe.Add(mBase, uint32(v6)))
			v139 = *(*float64)(unsafe.Add(mBase, uint32(v6)+8))
			v141 = v137 & int32(1)
			v143 = int32(0)
			v148 = base.I64_reinterpret_f64(v138)
			v152 = base.B2i32(base.Ui64(v148&int64(9223372002495037440)) < base.Ui64(int64(4604249089280835585)))
			if v152 == v143 {
				v161 = base.B2i32(int64(-1) < v148)
				if int64(-1) < v148 {
					v162 = v139
				} else {
					v162 = base.F64_neg(v139)
				}
				v166 = base.F64_add(base.F64_sub(float64(0.7853981633974483), base.F64_abs(v138)), base.F64_sub(float64(3.061616997868383e-17), v162))
				v167 = float64(0)
				v168 = v161
			} else {
				v166 = v138
				v167 = v139
				v168 = v143
			}
			v169 = base.F64_mul(v166, v166)
			v170 = base.F64_mul(v166, v169)
			v173 = base.F64_mul(v169, v169)
			v212 = base.F64_add(base.F64_mul(v170, float64(0.3333333333333341)), base.F64_add(base.F64_mul(v169, base.F64_add(base.F64_mul(v170, base.F64_add(base.F64_add(base.F64_mul(v173, base.F64_add(base.F64_mul(v173, base.F64_add(base.F64_mul(v173, base.F64_add(base.F64_mul(v173, base.F64_add(base.F64_mul(v173, float64(-1.8558637485527546e-05)), float64(7.817944429395571e-05))), float64(0.0005880412408202641))), float64(0.0035920791075913124))), float64(0.021869488294859542))), float64(0.13333333333320124)), base.F64_mul(v169, base.F64_add(base.F64_mul(v173, base.F64_add(base.F64_mul(v173, base.F64_add(base.F64_mul(v173, base.F64_add(base.F64_mul(v173, base.F64_add(base.F64_mul(v173, float64(2.590730518636337e-05)), float64(7.140724913826082e-05))), float64(0.0002464631348184699))), float64(0.0014562094543252903))), float64(0.0088632398235993))), float64(0.05396825397622605))))), v167)), v167))
			v213 = base.F64_add(v166, v212)
			if base.Ui64(v148&int64(9223372002495037440)) < base.Ui64(int64(4604249089280835585)) {
				if v141 == int32(0) {
					v252 = v213
				} else {
					v233 = base.F64_div(float64(-1), v213)
					v235 = int64(-4294967296)
					v237 = base.F64_reinterpret_i64(base.I64_reinterpret_f64(v233) & v235)
					v241 = base.F64_reinterpret_i64(base.I64_reinterpret_f64(v213) & v235)
					v252 = base.F64_add(base.F64_mul(v233, base.F64_add(base.F64_mul(v237, base.F64_sub(v212, base.F64_sub(v241, v166))), base.F64_add(base.F64_mul(v237, v241), float64(1)))), v237)
				}
				v255 = v252
			} else {
				v214 = int32(1)
				v218 = base.F64_convert_i32_s(v214 - v141<<(uint(v214)%32))
				v223 = base.F64_add(v166, base.F64_sub(v212, base.F64_div(base.F64_mul(v213, v213), base.F64_add(v213, v218))))
				v225 = base.F64_sub(v218, base.F64_add(v223, v223))
				if v168&v214 != 0 {
					v229 = v225
				} else {
					v229 = base.F64_neg(v225)
				}
				v255 = v229
			}
			v256 = v255
		} else {
			v256 = base.F64_sub(l0, l0)
		}
	} else {
		if base.Ui32(v13) < base.Ui32(int32(1044381696)) {
			v256 = l0
		} else {
			v18 = float64(0)
			v19 = int32(0)
			v26 = base.I64_reinterpret_f64(l0)
			v30 = base.B2i32(base.Ui64(v26&int64(9223372002495037440)) < base.Ui64(int64(4604249089280835585)))
			if v30 == v19 {
				v39 = base.B2i32(int64(-1) < v26)
				if int64(-1) < v26 {
					v40 = v18
				} else {
					v40 = base.F64_neg(v18)
				}
				v44 = base.F64_add(base.F64_sub(float64(0.7853981633974483), base.F64_abs(l0)), base.F64_sub(float64(3.061616997868383e-17), v40))
				v45 = float64(0)
				v46 = v39
			} else {
				v44 = l0
				v45 = v18
				v46 = v19
			}
			v47 = base.F64_mul(v44, v44)
			v48 = base.F64_mul(v44, v47)
			v51 = base.F64_mul(v47, v47)
			v90 = base.F64_add(base.F64_mul(v48, float64(0.3333333333333341)), base.F64_add(base.F64_mul(v47, base.F64_add(base.F64_mul(v48, base.F64_add(base.F64_add(base.F64_mul(v51, base.F64_add(base.F64_mul(v51, base.F64_add(base.F64_mul(v51, base.F64_add(base.F64_mul(v51, base.F64_add(base.F64_mul(v51, float64(-1.8558637485527546e-05)), float64(7.817944429395571e-05))), float64(0.0005880412408202641))), float64(0.0035920791075913124))), float64(0.021869488294859542))), float64(0.13333333333320124)), base.F64_mul(v47, base.F64_add(base.F64_mul(v51, base.F64_add(base.F64_mul(v51, base.F64_add(base.F64_mul(v51, base.F64_add(base.F64_mul(v51, base.F64_add(base.F64_mul(v51, float64(2.590730518636337e-05)), float64(7.140724913826082e-05))), float64(0.0002464631348184699))), float64(0.0014562094543252903))), float64(0.0088632398235993))), float64(0.05396825397622605))))), v45)), v45))
			v91 = base.F64_add(v44, v90)
			if base.Ui64(v26&int64(9223372002495037440)) < base.Ui64(int64(4604249089280835585)) {
				v133 = v91
			} else {
				v92 = int32(1)
				v96 = base.F64_convert_i32_s(v92)
				v101 = base.F64_add(v44, base.F64_sub(v90, base.F64_div(base.F64_mul(v91, v91), base.F64_add(v91, v96))))
				v103 = base.F64_sub(v96, base.F64_add(v101, v101))
				if v46&v92 != 0 {
					v107 = v103
				} else {
					v107 = base.F64_neg(v103)
				}
				v133 = v107
			}
			v256 = v133
		}
	}
	m.G0 = v6 + int32(16)
	return v256
}
func F_tanh(m *base.Module, l0 float64) float64 {
	var v4 float64
	_ = v4
	var v5 int64
	_ = v5
	var v15 float64
	_ = v15
	var v17 float64
	_ = v17
	var v25 float64
	_ = v25
	var v33 float64
	_ = v33
	var v38 float64
	_ = v38
	var v43 float64
	_ = v43
	v4 = base.F64_abs(l0)
	v5 = base.I64_reinterpret_f64(v4)
	if base.Ui64(v5) < base.Ui64(int64(4603122931675955200)) {
		if base.Ui64(v5) < base.Ui64(int64(4598272728187797504)) {
			if base.Ui64(v5) < base.Ui64(int64(4503599627370496)) {
				v38 = v4
			} else {
				v33 = F_expm1(m, base.F64_mul(v4, float64(-2)))
				v38 = base.F64_div(base.F64_neg(v33), base.F64_add(v33, float64(2)))
			}
		} else {
			v25 = F_expm1(m, base.F64_add(v4, v4))
			v38 = base.F64_div(v25, base.F64_add(v25, float64(2)))
		}
	} else {
		if base.Ui64(v5) < base.Ui64(int64(4626322721511309312)) {
			v15 = float64(2)
			v17 = F_expm1(m, base.F64_add(v4, v4))
			v38 = base.F64_sub(float64(1), base.F64_div(v15, base.F64_add(v17, v15)))
		} else {
			v38 = base.F64_add(base.F64_div(math.Float64frombits(uint64(0x8000000000000000)), v4), float64(1))
		}
	}
	if base.I64_reinterpret_f64(l0) < int64(0) {
		v43 = base.F64_neg(v38)
	} else {
		v43 = v38
	}
	return v43
}
func F_threebyte_strstr(m *base.Module, l0 int32, l1 int32) int32 {
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
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
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
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
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
	var v59 int32
	_ = v59
	var v64 int32
	_ = v64
	v3 = int32(0)
	v8 = l0 + int32(2)
	v9 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+2)))
	if v9 == v3 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	if v59 != 0 {
		goto L9
	} else {
		goto L10
	}
L2:
	;
	v56 = v8
	v59 = base.B2i32(v9 != v3)
	goto L1
L3:
	;
	v14 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+1)))
	v15 = int32(16)
	v17 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	v18 = int32(24)
	v21 = int32(8)
	v23 = v14<<(uint(v15)%32) | v17<<(uint(v18)%32) | v9<<(uint(v21)%32)
	v24 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+1)))
	v27 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	v31 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+2)))
	v34 = v24<<(uint(v15)%32) | v27<<(uint(v18)%32) | v31<<(uint(v21)%32)
	if v23 == v34 {
		goto L2
	} else {
		goto L4
	}
L4:
	;
	v38 = v8
	v39 = v23
	goto L5
L5:
	;
	v43 = v38 + int32(1)
	v44 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v38)+1)))
	v45 = int32(0)
	v46 = base.B2i32(v44 != v45)
	if v44 == v45 {
		v56 = v43
		v59 = v46
		goto L1
	} else {
		goto L7
	}
L7:
	;
	v51 = (v39 | v44) << (uint(int32(8)) % 32)
	if v51 != v34 {
		v38 = v43
		v39 = v51
		goto L5
	} else {
		goto L8
	}
L8:
	;
	v56 = v43
	v59 = v46
	goto L1
L9:
	;
	v64 = v56 + int32(-2)
	goto L11
L10:
	;
	v64 = int32(0)
	goto L11
L11:
	;
	return v64
}
func F_timeCommand(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int64
	_ = v6
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v10 int64
	_ = v10
	var v14 int64
	_ = v14
	var v17 int32
	_ = v17
	F_addReplyArrayLen(m, l0, int32(2))
	mBase = m.M
	v4 = m.ExcPending
	if v4 != 0 {
		return
	} else {
		v6 = *(*int64)(unsafe.Add(mBase, _c_F_timeCommand[0]))
		F_addReplyBulkLongLong(m, l0, v6)
		mBase = m.M
		v8 = m.ExcPending
		if v8 != 0 {
			return
		} else {
			v9 = int32(0)
			v10 = *(*int64)(unsafe.Add(mBase, _c_F_timeCommand[0]))
			v14 = *(*int64)(unsafe.Add(mBase, _c_F_timeCommand[1]))
			F_addReplyBulkLongLong(m, l0, v10*int64(-1000000)+v14)
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
func F_top16(m *base.Module, l0 float64) int32 {
	return base.I32_wrap_i64(int64(base.Ui64(base.I64_reinterpret_f64(l0)) >> (uint(int64(48)) % 64)))
}
func F_touchAllWatchedKeysInDb(m *base.Module, l0 int32, l1 int32) {
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
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v53 int64
	_ = v53
	var v54 int64
	_ = v54
	var v55 int64
	_ = v55
	var v56 int64
	_ = v56
	var v57 int64
	_ = v57
	var v58 int64
	_ = v58
	var v59 int64
	_ = v59
	var v61 int64
	_ = v61
	var v63 int64
	_ = v63
	var v65 int64
	_ = v65
	var v67 int64
	_ = v67
	var v69 int64
	_ = v69
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v91 int32
	_ = v91
	var v95 int32
	_ = v95
	var v97 int32
	_ = v97
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v111 int32
	_ = v111
	var v116 int32
	_ = v116
	var v120 int32
	_ = v120
	var v124 int32
	_ = v124
	var v133 int32
	_ = v133
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v149 int32
	_ = v149
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v159 int32
	_ = v159
	var v161 int32
	_ = v161
	var v164 int32
	_ = v164
	var v170 int32
	_ = v170
	var v177 int32
	_ = v177
	var v182 int32
	_ = v182
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v195 int32
	_ = v195
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v201 int32
	_ = v201
	var v203 int32
	_ = v203
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v213 int32
	_ = v213
	var v217 int32
	_ = v217
	var v219 int32
	_ = v219
	var v222 int32
	_ = v222
	var v228 int32
	_ = v228
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v249 int32
	_ = v249
	var v252 int32
	_ = v252
	var v253 int32
	_ = v253
	var v257 int32
	_ = v257
	var v261 int32
	_ = v261
	var v262 int32
	_ = v262
	var v263 int32
	_ = v263
	var v266 int32
	_ = v266
	var v268 int32
	_ = v268
	var v270 int64
	_ = v270
	var v271 int64
	_ = v271
	var v272 int64
	_ = v272
	var v273 int64
	_ = v273
	var v274 int64
	_ = v274
	var v275 int64
	_ = v275
	var v276 int64
	_ = v276
	var v278 int64
	_ = v278
	var v280 int64
	_ = v280
	var v282 int64
	_ = v282
	var v284 int64
	_ = v284
	var v286 int64
	_ = v286
	var v288 int32
	_ = v288
	var v289 int32
	_ = v289
	var v290 int32
	_ = v290
	var v293 int32
	_ = v293
	var v296 int32
	_ = v296
	var v297 int32
	_ = v297
	var v298 int32
	_ = v298
	var v300 int32
	_ = v300
	var v303 int32
	_ = v303
	var v304 int32
	_ = v304
	var v308 int32
	_ = v308
	var v312 int32
	_ = v312
	var v314 int32
	_ = v314
	var v321 int32
	_ = v321
	var v322 int32
	_ = v322
	var v323 int32
	_ = v323
	var v328 int32
	_ = v328
	var v333 int32
	_ = v333
	var v337 int32
	_ = v337
	var v341 int32
	_ = v341
	var v354 int32
	_ = v354
	v9 = m.G0
	v11 = v9 - int32(16)
	m.G0 = v11
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(v13)+12))
	v16 = *(*int32)(unsafe.Add(mBase, uint32(v13)+16))
	if v14 == int32(0)-v16 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	m.G0 = v11 + int32(16)
	return
L2:
	;
	v19 = F_dictGetSafeIterator(m, v13)
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		goto L4
	} else {
		goto L5
	}
L3:
	;
	F_dictReleaseIterator(m, v19)
	mBase = m.M
	v354 = m.ExcPending
	if v354 != 0 {
		goto L4
	} else {
		goto L98
	}
L4:
	;
	return
L5:
	;
	v28 = v19 + int32(20)
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v19)+16))
	if v29 != 0 {
		goto L8
	} else {
		goto L9
	}
L6:
	;
	if v124 == int32(0) {
		goto L3
	} else {
		goto L32
	}
L7:
	;
	v35 = v28
	v36 = v32
	goto L10
L8:
	;
	v32 = int32(1)
	goto L7
L9:
	;
	v32 = int32(0)
	goto L7
L10:
	;
	switch v36 {
	case 0:
		goto L15
	default:
		goto L14
	}
L12:
	;
	v36 = int32(0)
	goto L10
L13:
	;
	goto L6
L14:
	;
	v116 = *(*int32)(unsafe.Add(mBase, uint32(v35)))
	*(*int32)(unsafe.Add(mBase, uint32(v19)+16)) = v116
	if v116 == int32(0) {
		goto L12
	} else {
		goto L31
	}
L15:
	;
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v19)+4))
	if v40 != int32(-1) {
		v79 = v40
		goto L16
	} else {
		goto L17
	}
L16:
	;
	v80 = int32(1)
	v81 = v79 + v80
	*(*int32)(unsafe.Add(mBase, uint32(v19)+4)) = v81
	v83 = int32(0)
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v19)))
	v87 = *(*int32)(unsafe.Add(mBase, uint32(v19)+8))
	v91 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v86+v87+int32(26)))))
	if v91 == int32(255) {
		goto L25
	} else {
		goto L26
	}
L17:
	;
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v19)+8))
	if v44 != 0 {
		v79 = int32(-1)
		goto L16
	} else {
		goto L18
	}
L18:
	;
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v19)))
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v19)+12))
	if v46 == int32(0) {
		goto L20
	} else {
		goto L21
	}
L19:
	;
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v72)+20))
	if v73 != int32(-1) {
		goto L22
	} else {
		goto L23
	}
L20:
	;
	v53 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v45)+16)))
	v54 = int64(*(*int8)(unsafe.Add(mBase, uint32(v45)+27)))
	v55 = int64(*(*int32)(unsafe.Add(mBase, uint32(v45)+8)))
	v56 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v45)+12)))
	v57 = int64(*(*int8)(unsafe.Add(mBase, uint32(v45)+26)))
	v58 = int64(*(*int32)(unsafe.Add(mBase, uint32(v45)+4)))
	v59 = F_wangHash64(m, v58)
	mBase = m.M
	v61 = F_wangHash64(m, v57+v59)
	mBase = m.M
	v63 = F_wangHash64(m, v56+v61)
	mBase = m.M
	v65 = F_wangHash64(m, v55+v63)
	mBase = m.M
	v67 = F_wangHash64(m, v54+v65)
	mBase = m.M
	v69 = F_wangHash64(m, v53+v67)
	mBase = m.M
	*(*int64)(unsafe.Add(mBase, uint32(v19)+24)) = v69
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v19)))
	v72 = v71
	goto L19
L21:
	;
	v49 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v45)+24)))
	v51 = v49 + int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v45)+24)) = uint16(v51)
	v72 = v45
	goto L19
L22:
	;
	v79 = v73 + int32(-1)
	goto L16
L23:
	;
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v19)+4))
	v79 = v76
	goto L16
L24:
	;
	v106 = int32(2)
	v111 = *(*int32)(unsafe.Add(mBase, uint32(v86+v104<<(uint(v106)%32)+int32(4))))
	v35 = v111 + v105<<(uint(v106)%32)
	v36 = int32(1)
	goto L10
L25:
	;
	v95 = v83
	goto L27
L26:
	;
	v95 = v80 << (uint(v91) % 32)
	goto L27
L27:
	;
	if v81 < v95 {
		v104 = v87
		v105 = v81
		goto L24
	} else {
		goto L28
	}
L28:
	;
	if v87 != 0 {
		v124 = v83
		goto L13
	} else {
		goto L29
	}
L29:
	;
	v97 = *(*int32)(unsafe.Add(mBase, uint32(v86)+20))
	if v97 == int32(-1) {
		v124 = v83
		goto L13
	} else {
		goto L30
	}
L30:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v19)+4)) = int64(4294967296)
	v104 = int32(1)
	v105 = int32(0)
	goto L24
L31:
	;
	v120 = *(*int32)(unsafe.Add(mBase, uint32(v116)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v28))) = v120
	v124 = v116
	goto L13
L32:
	;
	v133 = v124
	goto L33
L33:
	;
	v138 = *(*int32)(unsafe.Add(mBase, uint32(v133)))
	goto L37
L34:
	;
	goto L3
L35:
	;
	v245 = v19 + int32(20)
	v246 = *(*int32)(unsafe.Add(mBase, uint32(v19)+16))
	if v246 != 0 {
		goto L73
	} else {
		goto L74
	}
L36:
	;
	v149 = *(*int32)(unsafe.Add(mBase, uint32(v133)+8))
	goto L43
L37:
	;
	v139 = F_objectGetVal(m, v138)
	mBase = m.M
	v140 = F_dbFind(m, l0, v139)
	mBase = m.M
	v141 = m.ExcPending
	if v141 != 0 {
		goto L4
	} else {
		goto L38
	}
L38:
	;
	if v140 != 0 {
		goto L36
	} else {
		goto L39
	}
L39:
	;
	if l1 == int32(0) {
		goto L35
	} else {
		goto L40
	}
L40:
	;
	v144 = F_objectGetVal(m, v138)
	mBase = m.M
	v145 = F_dbFind(m, l1, v144)
	mBase = m.M
	v146 = m.ExcPending
	if v146 != 0 {
		goto L4
	} else {
		goto L41
	}
L41:
	;
	if v145 == int32(0) {
		goto L35
	} else {
		goto L42
	}
L42:
	;
	goto L36
L43:
	;
	if v149 == int32(0) {
		goto L35
	} else {
		goto L44
	}
L44:
	;
	v153 = v11 + int32(8)
	v154 = *(*int32)(unsafe.Add(mBase, uint32(v149)))
	*(*int32)(unsafe.Add(mBase, uint32(v153)+4)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v153))) = v154
	goto L45
L45:
	;
	v159 = v11 + int32(8)
	v161 = *(*int32)(unsafe.Add(mBase, uint32(v159)))
	if v161 == int32(0) {
		goto L47
	} else {
		goto L48
	}
L46:
	;
	if v161 == int32(0) {
		goto L35
	} else {
		goto L49
	}
L47:
	;
	goto L46
L48:
	;
	v164 = *(*int32)(unsafe.Add(mBase, uint32(v159)+4))
	v170 = *(*int32)(unsafe.Add(mBase, uint32(v161+base.B2i32(v164 == int32(0))<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v159))) = v170
	goto L47
L49:
	;
	v177 = v161
	goto L50
L50:
	;
	v182 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v177)+24)))
	if v182&int32(1) == int32(0) {
		goto L55
	} else {
		goto L56
	}
L51:
	;
	goto L35
L52:
	;
	v217 = v11 + int32(8)
	v219 = *(*int32)(unsafe.Add(mBase, uint32(v217)))
	if v219 == int32(0) {
		goto L68
	} else {
		goto L69
	}
L53:
	;
	v207 = *(*int32)(unsafe.Add(mBase, uint32(v177)+20))
	v208 = *(*int32)(unsafe.Add(mBase, uint32(v207)+200))
	*(*int32)(unsafe.Add(mBase, uint32(v207)+200)) = v208 | int32(32)
	F_resetClientMultiState(m, v207)
	mBase = m.M
	v213 = m.ExcPending
	if v213 != 0 {
		goto L4
	} else {
		goto L66
	}
L54:
	;
	v205 = F_keyIsExpired(m, l1, v138)
	mBase = m.M
	v206 = m.ExcPending
	if v206 != 0 {
		goto L4
	} else {
		goto L64
	}
L55:
	;
	if v140 != 0 {
		goto L53
	} else {
		goto L61
	}
L56:
	;
	if l1 == int32(0) {
		v193 = v182
		goto L57
	} else {
		goto L58
	}
L57:
	;
	v195 = v193 & int32(254)
	*(*uint8)(unsafe.Add(mBase, uint32(v177)+24)) = uint8(v195)
	goto L52
L58:
	;
	v189 = F_objectGetVal(m, v138)
	mBase = m.M
	v190 = F_dbFind(m, l1, v189)
	mBase = m.M
	v191 = m.ExcPending
	if v191 != 0 {
		goto L4
	} else {
		goto L59
	}
L59:
	;
	if v190 != 0 {
		goto L54
	} else {
		goto L60
	}
L60:
	;
	v192 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v177)+24)))
	v193 = v192
	goto L57
L61:
	;
	v197 = F_keyIsExpired(m, l1, v138)
	mBase = m.M
	v198 = m.ExcPending
	if v198 != 0 {
		goto L4
	} else {
		goto L62
	}
L62:
	;
	if v197 == int32(0) {
		goto L53
	} else {
		goto L63
	}
L63:
	;
	v201 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v177)+24)))
	v203 = v201 | int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v177)+24)) = uint8(v203)
	goto L52
L64:
	;
	if v205 != 0 {
		goto L52
	} else {
		goto L65
	}
L65:
	;
	goto L53
L66:
	;
	goto L52
L67:
	;
	if v219 != 0 {
		v177 = v219
		goto L50
	} else {
		goto L70
	}
L68:
	;
	goto L67
L69:
	;
	v222 = *(*int32)(unsafe.Add(mBase, uint32(v217)+4))
	v228 = *(*int32)(unsafe.Add(mBase, uint32(v219+base.B2i32(v222 == int32(0))<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v217))) = v228
	goto L68
L70:
	;
	goto L51
L71:
	;
	if v341 != 0 {
		v133 = v341
		goto L33
	} else {
		goto L97
	}
L72:
	;
	v252 = v245
	v253 = v249
	goto L75
L73:
	;
	v249 = int32(1)
	goto L72
L74:
	;
	v249 = int32(0)
	goto L72
L75:
	;
	switch v253 {
	case 0:
		goto L80
	default:
		goto L79
	}
L77:
	;
	v253 = int32(0)
	goto L75
L78:
	;
	goto L71
L79:
	;
	v333 = *(*int32)(unsafe.Add(mBase, uint32(v252)))
	*(*int32)(unsafe.Add(mBase, uint32(v19)+16)) = v333
	if v333 == int32(0) {
		goto L77
	} else {
		goto L96
	}
L80:
	;
	v257 = *(*int32)(unsafe.Add(mBase, uint32(v19)+4))
	if v257 != int32(-1) {
		v296 = v257
		goto L81
	} else {
		goto L82
	}
L81:
	;
	v297 = int32(1)
	v298 = v296 + v297
	*(*int32)(unsafe.Add(mBase, uint32(v19)+4)) = v298
	v300 = int32(0)
	v303 = *(*int32)(unsafe.Add(mBase, uint32(v19)))
	v304 = *(*int32)(unsafe.Add(mBase, uint32(v19)+8))
	v308 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v303+v304+int32(26)))))
	if v308 == int32(255) {
		goto L90
	} else {
		goto L91
	}
L82:
	;
	v261 = *(*int32)(unsafe.Add(mBase, uint32(v19)+8))
	if v261 != 0 {
		v296 = int32(-1)
		goto L81
	} else {
		goto L83
	}
L83:
	;
	v262 = *(*int32)(unsafe.Add(mBase, uint32(v19)))
	v263 = *(*int32)(unsafe.Add(mBase, uint32(v19)+12))
	if v263 == int32(0) {
		goto L85
	} else {
		goto L86
	}
L84:
	;
	v290 = *(*int32)(unsafe.Add(mBase, uint32(v289)+20))
	if v290 != int32(-1) {
		goto L87
	} else {
		goto L88
	}
L85:
	;
	v270 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v262)+16)))
	v271 = int64(*(*int8)(unsafe.Add(mBase, uint32(v262)+27)))
	v272 = int64(*(*int32)(unsafe.Add(mBase, uint32(v262)+8)))
	v273 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v262)+12)))
	v274 = int64(*(*int8)(unsafe.Add(mBase, uint32(v262)+26)))
	v275 = int64(*(*int32)(unsafe.Add(mBase, uint32(v262)+4)))
	v276 = F_wangHash64(m, v275)
	mBase = m.M
	v278 = F_wangHash64(m, v274+v276)
	mBase = m.M
	v280 = F_wangHash64(m, v273+v278)
	mBase = m.M
	v282 = F_wangHash64(m, v272+v280)
	mBase = m.M
	v284 = F_wangHash64(m, v271+v282)
	mBase = m.M
	v286 = F_wangHash64(m, v270+v284)
	mBase = m.M
	*(*int64)(unsafe.Add(mBase, uint32(v19)+24)) = v286
	v288 = *(*int32)(unsafe.Add(mBase, uint32(v19)))
	v289 = v288
	goto L84
L86:
	;
	v266 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v262)+24)))
	v268 = v266 + int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v262)+24)) = uint16(v268)
	v289 = v262
	goto L84
L87:
	;
	v296 = v290 + int32(-1)
	goto L81
L88:
	;
	v293 = *(*int32)(unsafe.Add(mBase, uint32(v19)+4))
	v296 = v293
	goto L81
L89:
	;
	v323 = int32(2)
	v328 = *(*int32)(unsafe.Add(mBase, uint32(v303+v321<<(uint(v323)%32)+int32(4))))
	v252 = v328 + v322<<(uint(v323)%32)
	v253 = int32(1)
	goto L75
L90:
	;
	v312 = v300
	goto L92
L91:
	;
	v312 = v297 << (uint(v308) % 32)
	goto L92
L92:
	;
	if v298 < v312 {
		v321 = v304
		v322 = v298
		goto L89
	} else {
		goto L93
	}
L93:
	;
	if v304 != 0 {
		v341 = v300
		goto L78
	} else {
		goto L94
	}
L94:
	;
	v314 = *(*int32)(unsafe.Add(mBase, uint32(v303)+20))
	if v314 == int32(-1) {
		v341 = v300
		goto L78
	} else {
		goto L95
	}
L95:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v19)+4)) = int64(4294967296)
	v321 = int32(1)
	v322 = int32(0)
	goto L89
L96:
	;
	v337 = *(*int32)(unsafe.Add(mBase, uint32(v333)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v245))) = v337
	v341 = v333
	goto L78
L97:
	;
	goto L34
L98:
	;
	goto L1
}
func F_towlower(m *base.Module, l0 int32) int32 {
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
	v2 = int32(0)
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
	v25 = *(*int32)(unsafe.Add(mBase, uint32((l0-v15*v14)&v12<<(uint(v21)%32))+uint32(_c_F_towlower[0])))
	v26 = int32(8)
	v27 = int32(base.Ui32(l0) >> (uint(v26) % 32))
	v30 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v27)+uint32(_c_F_towlower[1]))))
	v36 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v30*int32(86)+v15)+uint32(_c_F_towlower[1]))))
	v41 = base.I32_rem_u_s(int32(base.Ui32(v25*v36)>>(uint(int32(11))%32)), int32(6))
	v44 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v27)+uint32(_c_F_towlower[2]))))
	v50 = *(*int32)(unsafe.Add(mBase, uint32((v41+v44)<<(uint(v21)%32))+uint32(_c_F_towlower[3])))
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
	v84 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v81)+uint32(_c_F_towlower[4]))))
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
	v86 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v81)+uint32(_c_F_towlower[5]))))
	v91 = *(*int32)(unsafe.Add(mBase, uint32(v86<<(uint(int32(2))%32))+uint32(_c_F_towlower[3])))
	v93 = v91 & int32(255)
	if base.Ui32(int32(1)) < base.Ui32(v93) {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	goto L14
L12:
	;
	v129 = v91>>(uint(int32(8))%32)&(int32(0)-(v93^v2)) + l0
	goto L1
L14:
	;
	goto L15
L15:
	;
	v129 = int32(1) + l0
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
func F_trinkle(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32) {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v27 int32
	_ = v27
	var v33 int32
	_ = v33
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v59 int32
	_ = v59
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v85 int32
	_ = v85
	var v89 int32
	_ = v89
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
	var v104 int32
	_ = v104
	var v107 int32
	_ = v107
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v140 int32
	_ = v140
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v166 int32
	_ = v166
	var v168 int32
	_ = v168
	var v174 int32
	_ = v174
	var v178 int32
	_ = v178
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v200 int32
	_ = v200
	var v202 int32
	_ = v202
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v207 int32
	_ = v207
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v217 int32
	_ = v217
	var v230 int32
	_ = v230
	v9 = int32(0)
	v15 = m.G0
	v17 = v15 - int32(240)
	m.G0 = v17
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
	*(*int32)(unsafe.Add(mBase, uint32(v17)+232)) = v19
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l4)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v17))) = l0
	*(*int32)(unsafe.Add(mBase, uint32(v17)+236)) = v21
	v27 = base.B2i32(l6 == v9)
	if v19 == int32(1) {
		goto L5
	} else {
		goto L6
	}
L1:
	;
	m.G0 = v17 + int32(240)
	return
L2:
	;
	v166 = m.G0
	v168 = v166 - int32(256)
	m.G0 = v168
	if v153 < int32(2) {
		goto L35
	} else {
		goto L36
	}
L3:
	;
	if v140&int32(1) == int32(0) {
		goto L1
	} else {
		goto L33
	}
L4:
	;
	v39 = l5
	v40 = v33
	v43 = l0
	v45 = v27
	goto L8
L5:
	;
	if v21 != 0 {
		v33 = int32(1)
		goto L4
	} else {
		goto L7
	}
L6:
	;
	v33 = int32(1)
	goto L4
L7:
	;
	v133 = l0
	v134 = l5
	v135 = int32(1)
	v140 = v27
	goto L3
L8:
	;
	v50 = l7 + v39<<(uint(int32(2))%32)
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v50)))
	v52 = v43 - v51
	v53 = m.T0[l2].(func(*base.Module, int32, int32, int32) int32)(m, v52, l0, l3)
	mBase = m.M
	v54 = m.ExcPending
	if v54 != 0 {
		goto L11
	} else {
		goto L12
	}
L9:
	;
	v151 = v43
	v152 = v39
	v153 = v40
	goto L2
L10:
	;
	v59 = int32(1)
	if (v45^int32(-1)|base.B2i32(v39 < int32(2)))&v59 != 0 {
		goto L15
	} else {
		goto L16
	}
L11:
	;
	return
L12:
	;
	if int32(1) <= v53 {
		goto L10
	} else {
		goto L13
	}
L13:
	;
	v133 = v43
	v134 = v39
	v135 = v40
	v140 = v45
	goto L3
L14:
	;
	goto L9
L15:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17+v40<<(uint(int32(2))%32)))) = v52
	v85 = v17 + int32(232)
	v89 = *(*int32)(unsafe.Add(mBase, uint32(v85)))
	v92 = F___builtin_ctz(m, v89+int32(-1))
	mBase = m.M
	if v92 != 0 {
		v100 = v92
		goto L22
	} else {
		goto L23
	}
L16:
	;
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v50+int32(-8))))
	v68 = v43 + (v9 - l1)
	v69 = m.T0[l2].(func(*base.Module, int32, int32, int32) int32)(m, v68, v52, l3)
	mBase = m.M
	v70 = m.ExcPending
	if v70 != 0 {
		goto L11
	} else {
		goto L17
	}
L17:
	;
	if int32(-1) < v69 {
		goto L14
	} else {
		goto L18
	}
L18:
	;
	v74 = m.T0[l2].(func(*base.Module, int32, int32, int32) int32)(m, v68-v67, v52, l3)
	mBase = m.M
	v75 = m.ExcPending
	if v75 != 0 {
		goto L11
	} else {
		goto L19
	}
L19:
	;
	if int32(-1) < v74 {
		goto L14
	} else {
		goto L20
	}
L20:
	;
	goto L15
L21:
	;
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v85)+4))
	if base.Ui32(int32(31)) < base.Ui32(v100) {
		goto L29
	} else {
		goto L30
	}
L22:
	;
	goto L21
L23:
	;
	v93 = *(*int32)(unsafe.Add(mBase, uint32(v85)+4))
	v94 = F___builtin_ctz(m, v93)
	mBase = m.M
	if v94 != 0 {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	v98 = v94 | int32(32)
	goto L26
L25:
	;
	v98 = int32(0)
	goto L26
L26:
	;
	v100 = v98
	goto L22
L27:
	;
	v122 = int32(1)
	v123 = v40 + v122
	v124 = v100 + v39
	v125 = *(*int32)(unsafe.Add(mBase, uint32(v17)+236))
	v126 = *(*int32)(unsafe.Add(mBase, uint32(v17)+232))
	if v126 != v122 {
		v39 = v124
		v40 = v123
		v43 = v52
		v45 = v59
		goto L8
	} else {
		goto L31
	}
L28:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v85)+4)) = int32(base.Ui32(v113) >> (uint(v111) % 32))
	*(*int32)(unsafe.Add(mBase, uint32(v85))) = v113<<(uint(int32(32)-v111)%32) | int32(base.Ui32(v112)>>(uint(v111)%32))
	goto L27
L29:
	;
	v111 = v100 + int32(-32)
	v112 = v104
	v113 = int32(0)
	goto L28
L30:
	;
	v107 = *(*int32)(unsafe.Add(mBase, uint32(v85)))
	v111 = v100
	v112 = v107
	v113 = v104
	goto L28
L31:
	;
	if v125 != 0 {
		v39 = v124
		v40 = v123
		v43 = v52
		v45 = v59
		goto L8
	} else {
		goto L32
	}
L32:
	;
	v151 = v52
	v152 = v124
	v153 = v123
	goto L2
L33:
	;
	v151 = v133
	v152 = v134
	v153 = v135
	goto L2
L34:
	;
	F_sift(m, v151, l1, l2, l3, v152, l7)
	mBase = m.M
	v230 = m.ExcPending
	if v230 != 0 {
		goto L11
	} else {
		goto L47
	}
L35:
	;
	m.G0 = v168 + int32(256)
	goto L34
L36:
	;
	v174 = v17 + v153<<(uint(int32(2))%32)
	*(*int32)(unsafe.Add(mBase, uint32(v174))) = v168
	if l1 == int32(0) {
		goto L35
	} else {
		goto L37
	}
L37:
	;
	v178 = l1
	goto L38
L38:
	;
	v186 = *(*int32)(unsafe.Add(mBase, uint32(v174)))
	v187 = *(*int32)(unsafe.Add(mBase, uint32(v17)))
	v188 = int32(256)
	if base.Ui32(v178) < base.Ui32(v188) {
		goto L40
	} else {
		goto L41
	}
L39:
	;
	goto L35
L40:
	;
	v191 = v178
	goto L42
L41:
	;
	v191 = v188
	goto L42
L42:
	;
	v192 = F___memcpy(m, v186, v187, v191)
	mBase = m.M
	v200 = int32(0)
	goto L43
L43:
	;
	v202 = int32(2)
	v204 = v17 + v200<<(uint(v202)%32)
	v205 = *(*int32)(unsafe.Add(mBase, uint32(v204)))
	v207 = v200 + int32(1)
	v211 = *(*int32)(unsafe.Add(mBase, uint32(v17+v207<<(uint(v202)%32))))
	v212 = F___memcpy(m, v205, v211, v191)
	mBase = m.M
	v213 = *(*int32)(unsafe.Add(mBase, uint32(v204)))
	*(*int32)(unsafe.Add(mBase, uint32(v204))) = v213 + v191
	if v207 != v153 {
		v200 = v207
		goto L43
	} else {
		goto L45
	}
L44:
	;
	v217 = v178 - v191
	if v217 != 0 {
		v178 = v217
		goto L38
	} else {
		goto L46
	}
L45:
	;
	goto L44
L46:
	;
	goto L39
L47:
	;
	goto L1
}
func F_tryOffloadFreeArgvToIOThreads(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v20 int64
	_ = v20
	var v24 int64
	_ = v24
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v37 int32
	_ = v37
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
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v95 int32
	_ = v95
	var v98 int32
	_ = v98
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v115 int32
	_ = v115
	var v127 int32
	_ = v127
	var v129 int32
	_ = v129
	var v136 int32
	_ = v136
	v9 = int32(-1)
	if l1 == int32(0) {
		v136 = v9
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return v136
L2:
	;
	v13 = *(*int32)(unsafe.Add(mBase, _c_F_tryOffloadFreeArgvToIOThreads[0]))
	if v13 < int32(2) {
		v136 = v9
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v16 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+225)))
	if v16 == int32(0) {
		goto L5
	} else {
		goto L6
	}
L4:
	;
	v30 = v28 * int32(192)
	v33 = int32(0)
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v30)+uint32(_c_F_tryOffloadFreeArgvToIOThreads[1])))
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v30)+uint32(_c_F_tryOffloadFreeArgvToIOThreads[2])))
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v30)+uint32(_c_F_tryOffloadFreeArgvToIOThreads[3])))
	if base.Ui32(v37-v38) < base.Ui32(v40) {
		v53 = v33
		goto L9
	} else {
		goto L10
	}
L5:
	;
	v20 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	v24 = base.I64_rem_u_s(v20, base.I64_extend_i32_u(v13+int32(-1)))
	v28 = base.I32_wrap_i64(v24) + int32(1)
	goto L4
L6:
	;
	if base.Ui32(v16) < base.Ui32(v13) {
		v28 = v16
		goto L4
	} else {
		goto L7
	}
L7:
	;
	goto L5
L8:
	;
	if v53 != 0 {
		v136 = v9
		goto L1
	} else {
		goto L13
	}
L9:
	;
	goto L8
L10:
	;
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v30)+uint32(_c_F_tryOffloadFreeArgvToIOThreads[4])))
	*(*int32)(unsafe.Add(mBase, uint32(v30)+uint32(_c_F_tryOffloadFreeArgvToIOThreads[2]))) = v42
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v30)+uint32(_c_F_tryOffloadFreeArgvToIOThreads[3])))
	if base.Ui32(v37-v42) < base.Ui32(v45) {
		v53 = v33
		goto L9
	} else {
		goto L11
	}
L11:
	;
	v47 = int32(1)
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v30)+uint32(_c_F_tryOffloadFreeArgvToIOThreads[1])))
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v30)+uint32(_c_F_tryOffloadFreeArgvToIOThreads[5])))
	if v48 == v49 {
		v53 = v47
		goto L9
	} else {
		goto L12
	}
L12:
	;
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v30)+uint32(_c_F_tryOffloadFreeArgvToIOThreads[1])))
	*(*int32)(unsafe.Add(mBase, uint32(v30)+uint32(_c_F_tryOffloadFreeArgvToIOThreads[5]))) = v51
	v53 = v47
	goto L9
L13:
	;
	v55 = int32(0)
	if l1 <= v55 {
		goto L15
	} else {
		goto L16
	}
L14:
	;
	v98 = int32(2)
	v101 = *(*int32)(unsafe.Add(mBase, uint32(l2+v80<<(uint(v98)%32))))
	v102 = *(*int32)(unsafe.Add(mBase, uint32(v101)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v101)+4)) = v102 & int32(7)
	v111 = *(*int32)(unsafe.Add(mBase, uint32(v30)+uint32(_c_F_tryOffloadFreeArgvToIOThreads[6])))
	v112 = *(*int32)(unsafe.Add(mBase, uint32(v30)+uint32(_c_F_tryOffloadFreeArgvToIOThreads[3])))
	v115 = *(*int32)(unsafe.Add(mBase, uint32(v30)+uint32(_c_F_tryOffloadFreeArgvToIOThreads[1])))
	*(*int32)(unsafe.Add(mBase, uint32(v111+(v112+int32(-1))&v115<<(uint(v98)%32)))) = l2 | v98
	*(*int32)(unsafe.Add(mBase, uint32(v30)+uint32(_c_F_tryOffloadFreeArgvToIOThreads[1]))) = v115 + int32(1)
	goto L27
L15:
	;
	F_valkey_free(m, l2)
	mBase = m.M
	v95 = m.ExcPending
	if v95 != 0 {
		goto L21
	} else {
		goto L25
	}
L16:
	;
	v62 = v55
	v64 = int32(-1)
	goto L17
L17:
	;
	v69 = l2 + v62<<(uint(int32(2))%32)
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v69)))
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v70)+4))
	if base.Ui32(v71) < base.Ui32(int32(16)) {
		v80 = v62
		goto L19
	} else {
		goto L20
	}
L18:
	;
	if v80 != int32(-1) {
		goto L14
	} else {
		goto L24
	}
L19:
	;
	v82 = v62 + int32(1)
	if v82 != l1 {
		v62 = v82
		v64 = v80
		goto L17
	} else {
		goto L23
	}
L20:
	;
	F_decrRefCount(m, v70)
	mBase = m.M
	v77 = m.ExcPending
	if v77 != 0 {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	return int32(0)
L22:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v69))) = int32(0)
	v80 = v64
	goto L19
L23:
	;
	goto L18
L24:
	;
	goto L15
L25:
	;
	return int32(0)
L26:
	;
	v127 = int32(0)
	v129 = *(*int32)(unsafe.Add(mBase, _c_F_tryOffloadFreeArgvToIOThreads[7]))
	*(*int32)(unsafe.Add(mBase, _c_F_tryOffloadFreeArgvToIOThreads[7])) = v129 + int32(1)
	v136 = int32(0)
	goto L1
L27:
	;
	goto L26
}
func F_trySendPollJobToIOThreads(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v8 int64
	_ = v8
	var v10 int64
	_ = v10
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v87 int32
	_ = v87
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
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
	var v110 int32
	_ = v110
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v118 int32
	_ = v118
	var v120 int32
	_ = v120
	var v130 int32
	_ = v130
	var v132 int32
	_ = v132
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v141 int32
	_ = v141
	var v148 int32
	_ = v148
	var v155 int32
	_ = v155
	var v158 int32
	_ = v158
	var v160 int32
	_ = v160
	v4 = *(*int32)(unsafe.Add(mBase, _c_F_trySendPollJobToIOThreads[0]))
	if v4 < int32(2) {
		return
	} else {
		v7 = int32(_a_F_trySendPollJobToIOThreads_0)
		v8 = *(*int64)(unsafe.Add(mBase, _c_F_trySendPollJobToIOThreads[1]))
		v10 = *(*int64)(unsafe.Add(mBase, _c_F_trySendPollJobToIOThreads[2]))
		if base.I32_wrap_i64(v8+v10) == int32(0) {
			return
		} else {
			v16 = *(*int32)(unsafe.Add(mBase, _c_F_trySendPollJobToIOThreads[3]))
			if v16 != 0 {
				return
			} else {
				v17 = int32(_a_F_trySendPollJobToIOThreads_0)
				*(*int32)(unsafe.Add(mBase, _c_F_trySendPollJobToIOThreads[3])) = int32(1)
				v21 = *(*int32)(unsafe.Add(mBase, _c_F_trySendPollJobToIOThreads[4]))
				v23 = *(*int32)(unsafe.Add(mBase, uint32(v21)+72))
				*(*int32)(unsafe.Add(mBase, uint32(v21)+72)) = v23&int32(-33) | int32(32)
				v33 = v21 | int32(4)
				v35 = *(*int32)(unsafe.Add(mBase, _c_F_trySendPollJobToIOThreads[0]))
				if int32(9) < v35 {
					v77 = int32(0)
					v79 = *(*int32)(unsafe.Add(mBase, _c_F_trySendPollJobToIOThreads[5]))
					v82 = base.I32_rem_s(v79, v35+int32(-1))
					v84 = v82 + int32(1)
					*(*int32)(unsafe.Add(mBase, _c_F_trySendPollJobToIOThreads[5])) = v84
					v87 = v84 * int32(192)
					v94 = *(*int32)(unsafe.Add(mBase, uint32(v87)+uint32(_c_F_trySendPollJobToIOThreads[6])))
					v95 = *(*int32)(unsafe.Add(mBase, uint32(v87)+uint32(_c_F_trySendPollJobToIOThreads[7])))
					v97 = *(*int32)(unsafe.Add(mBase, uint32(v87)+uint32(_c_F_trySendPollJobToIOThreads[8])))
					if base.Ui32(v94-v95) < base.Ui32(v97) {
						v110 = v77
					} else {
						v99 = *(*int32)(unsafe.Add(mBase, uint32(v87)+uint32(_c_F_trySendPollJobToIOThreads[9])))
						*(*int32)(unsafe.Add(mBase, uint32(v87)+uint32(_c_F_trySendPollJobToIOThreads[7]))) = v99
						v102 = *(*int32)(unsafe.Add(mBase, uint32(v87)+uint32(_c_F_trySendPollJobToIOThreads[8])))
						if base.Ui32(v94-v99) < base.Ui32(v102) {
							v110 = v77
						} else {
							v104 = int32(1)
							v105 = *(*int32)(unsafe.Add(mBase, uint32(v87)+uint32(_c_F_trySendPollJobToIOThreads[6])))
							v106 = *(*int32)(unsafe.Add(mBase, uint32(v87)+uint32(_c_F_trySendPollJobToIOThreads[10])))
							if v105 == v106 {
								v110 = v104
							} else {
								v108 = *(*int32)(unsafe.Add(mBase, uint32(v87)+uint32(_c_F_trySendPollJobToIOThreads[6])))
								*(*int32)(unsafe.Add(mBase, uint32(v87)+uint32(_c_F_trySendPollJobToIOThreads[10]))) = v108
								v110 = v104
							}
						}
					}
					if v110 == int32(0) {
						v130 = *(*int32)(unsafe.Add(mBase, _c_F_trySendPollJobToIOThreads[5]))
						v132 = v130 * int32(192)
						v137 = *(*int32)(unsafe.Add(mBase, uint32(v132)+uint32(_c_F_trySendPollJobToIOThreads[11])))
						v138 = *(*int32)(unsafe.Add(mBase, uint32(v132)+uint32(_c_F_trySendPollJobToIOThreads[8])))
						v141 = *(*int32)(unsafe.Add(mBase, uint32(v132)+uint32(_c_F_trySendPollJobToIOThreads[6])))
						*(*int32)(unsafe.Add(mBase, uint32(v137+(v138+int32(-1))&v141<<(uint(int32(2))%32)))) = v33
						v148 = v141 + int32(1)
						*(*int32)(unsafe.Add(mBase, uint32(v132)+uint32(_c_F_trySendPollJobToIOThreads[6]))) = v148
						*(*int32)(unsafe.Add(mBase, uint32(v132)+uint32(_c_F_trySendPollJobToIOThreads[10]))) = v148
						v155 = *(*int32)(unsafe.Add(mBase, _c_F_trySendPollJobToIOThreads[4]))
						*(*int32)(unsafe.Add(mBase, uint32(v155)+44)) = int32(542)
						v158 = int32(0)
						v160 = *(*int32)(unsafe.Add(mBase, _c_F_trySendPollJobToIOThreads[12]))
						*(*int32)(unsafe.Add(mBase, _c_F_trySendPollJobToIOThreads[12])) = v160 + int32(1)
						return
					} else {
						v114 = int32(_a_F_trySendPollJobToIOThreads_0)
						v115 = int32(0)
						*(*int32)(unsafe.Add(mBase, _c_F_trySendPollJobToIOThreads[3])) = v115
						v118 = *(*int32)(unsafe.Add(mBase, _c_F_trySendPollJobToIOThreads[4]))
						v120 = *(*int32)(unsafe.Add(mBase, uint32(v118)+72))
						*(*int32)(unsafe.Add(mBase, uint32(v118)+72)) = v120&int32(-33) | v115
						return
					}
				} else {
					v38 = int32(_a_F_trySendPollJobToIOThreads_1)
					v42 = *(*int32)(unsafe.Add(mBase, _c_F_trySendPollJobToIOThreads[13]))
					v43 = *(*int32)(unsafe.Add(mBase, _c_F_trySendPollJobToIOThreads[14]))
					v46 = *(*int32)(unsafe.Add(mBase, _c_F_trySendPollJobToIOThreads[15]))
					v50 = v42 + (v43+int32(-1))&v46<<(uint(int32(6))%32)
					v51 = *(*int32)(unsafe.Add(mBase, uint32(v50)))
					if v51 != v46 {
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v50)+4)) = v33
						v54 = int32(1)
						*(*int32)(unsafe.Add(mBase, uint32(v50))) = v51 + v54
						v57 = *(*int32)(unsafe.Add(mBase, _c_F_trySendPollJobToIOThreads[15]))
						*(*int32)(unsafe.Add(mBase, _c_F_trySendPollJobToIOThreads[15])) = v57 + v54
					}
					if v51 == v46 {
						v155 = *(*int32)(unsafe.Add(mBase, _c_F_trySendPollJobToIOThreads[4]))
						*(*int32)(unsafe.Add(mBase, uint32(v155)+44)) = int32(542)
						v158 = int32(0)
						v160 = *(*int32)(unsafe.Add(mBase, _c_F_trySendPollJobToIOThreads[12]))
						*(*int32)(unsafe.Add(mBase, _c_F_trySendPollJobToIOThreads[12])) = v160 + int32(1)
						return
					} else {
						v62 = int32(_a_F_trySendPollJobToIOThreads_0)
						v63 = int32(0)
						*(*int32)(unsafe.Add(mBase, _c_F_trySendPollJobToIOThreads[3])) = v63
						v66 = *(*int32)(unsafe.Add(mBase, _c_F_trySendPollJobToIOThreads[4]))
						v68 = *(*int32)(unsafe.Add(mBase, uint32(v66)+72))
						*(*int32)(unsafe.Add(mBase, uint32(v66)+72)) = v68&int32(-33) | v63
						return
					}
				}
			}
		}
	}
}
func F_try_realloc_chunk(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v25 int32
	_ = v25
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v91 int32
	_ = v91
	var v96 int32
	_ = v96
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v107 int32
	_ = v107
	var __phi107 int32
	_ = __phi107
	var v108 int32
	_ = v108
	var __phi108 int32
	_ = __phi108
	var v114 int32
	_ = v114
	var v117 int32
	_ = v117
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v140 int32
	_ = v140
	var v148 int32
	_ = v148
	var v150 int32
	_ = v150
	var v153 int32
	_ = v153
	var v156 int32
	_ = v156
	var v158 int32
	_ = v158
	var v163 int32
	_ = v163
	var v170 int32
	_ = v170
	var v175 int32
	_ = v175
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v189 int32
	_ = v189
	var v193 int32
	_ = v193
	var v195 int32
	_ = v195
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v206 int32
	_ = v206
	var v208 int32
	_ = v208
	var v215 int32
	_ = v215
	var v217 int32
	_ = v217
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v235 int32
	_ = v235
	var v237 int32
	_ = v237
	var v239 int32
	_ = v239
	var v248 int32
	_ = v248
	var v250 int32
	_ = v250
	var v253 int32
	_ = v253
	var v258 int32
	_ = v258
	var v263 int32
	_ = v263
	var v264 int32
	_ = v264
	var v269 int32
	_ = v269
	var __phi269 int32
	_ = __phi269
	var v270 int32
	_ = v270
	var __phi270 int32
	_ = __phi270
	var v276 int32
	_ = v276
	var v279 int32
	_ = v279
	var v294 int32
	_ = v294
	var v302 int32
	_ = v302
	var v304 int32
	_ = v304
	var v307 int32
	_ = v307
	var v310 int32
	_ = v310
	var v312 int32
	_ = v312
	var v317 int32
	_ = v317
	var v324 int32
	_ = v324
	var v329 int32
	_ = v329
	var v352 int32
	_ = v352
	var v363 int32
	_ = v363
	var v365 int32
	_ = v365
	var v367 int32
	_ = v367
	var v371 int32
	_ = v371
	var v376 int32
	_ = v376
	var v377 int32
	_ = v377
	var v388 int32
	_ = v388
	var v391 int32
	_ = v391
	var v398 int32
	_ = v398
	var v403 int32
	_ = v403
	var v407 int32
	_ = v407
	var v409 int32
	_ = v409
	var v423 int32
	_ = v423
	var v425 int32
	_ = v425
	var v429 int32
	_ = v429
	var v431 int32
	_ = v431
	var v435 int32
	_ = v435
	var v445 int32
	_ = v445
	var v446 int32
	_ = v446
	var v462 int32
	_ = v462
	var v487 int32
	_ = v487
	var v489 int32
	_ = v489
	var v492 int32
	_ = v492
	var v493 int32
	_ = v493
	var v495 int32
	_ = v495
	var v501 int32
	_ = v501
	var v502 int32
	_ = v502
	var v506 int32
	_ = v506
	var v511 int32
	_ = v511
	var v513 int32
	_ = v513
	var v515 int32
	_ = v515
	var v516 int32
	_ = v516
	var v518 int32
	_ = v518
	var v521 int32
	_ = v521
	var v527 int32
	_ = v527
	var v531 int32
	_ = v531
	var v533 int32
	_ = v533
	var v537 int32
	_ = v537
	var v543 int32
	_ = v543
	var v544 int32
	_ = v544
	var v548 int32
	_ = v548
	var v550 int32
	_ = v550
	var v552 int32
	_ = v552
	var v553 int32
	_ = v553
	var v557 int32
	_ = v557
	var v558 int32
	_ = v558
	var v563 int32
	_ = v563
	var v565 int32
	_ = v565
	var v566 int32
	_ = v566
	var v569 int32
	_ = v569
	var v571 int32
	_ = v571
	var v573 int32
	_ = v573
	var v582 int32
	_ = v582
	var v584 int32
	_ = v584
	var v587 int32
	_ = v587
	var v592 int32
	_ = v592
	var v597 int32
	_ = v597
	var v598 int32
	_ = v598
	var v603 int32
	_ = v603
	var __phi603 int32
	_ = __phi603
	var v605 int32
	_ = v605
	var __phi605 int32
	_ = __phi605
	var v612 int32
	_ = v612
	var v615 int32
	_ = v615
	var v622 int32
	_ = v622
	var v632 int32
	_ = v632
	var v634 int32
	_ = v634
	var v637 int32
	_ = v637
	var v640 int32
	_ = v640
	var v642 int32
	_ = v642
	var v647 int32
	_ = v647
	var v654 int32
	_ = v654
	var v659 int32
	_ = v659
	var v677 int32
	_ = v677
	var v683 int32
	_ = v683
	var v684 int32
	_ = v684
	var v688 int32
	_ = v688
	var v694 int32
	_ = v694
	var v698 int32
	_ = v698
	var v699 int32
	_ = v699
	var v710 int32
	_ = v710
	var v711 int32
	_ = v711
	var v718 int32
	_ = v718
	var v719 int32
	_ = v719
	var v720 int32
	_ = v720
	var v722 int32
	_ = v722
	var v724 int32
	_ = v724
	var v727 int32
	_ = v727
	var v729 int32
	_ = v729
	var v731 int32
	_ = v731
	var v738 int32
	_ = v738
	var v740 int32
	_ = v740
	var v743 int32
	_ = v743
	var v748 int32
	_ = v748
	var v753 int32
	_ = v753
	var v754 int32
	_ = v754
	var v759 int32
	_ = v759
	var __phi759 int32
	_ = __phi759
	var v760 int32
	_ = v760
	var __phi760 int32
	_ = __phi760
	var v766 int32
	_ = v766
	var v769 int32
	_ = v769
	var v772 int32
	_ = v772
	var v773 int32
	_ = v773
	var v792 int32
	_ = v792
	var v800 int32
	_ = v800
	var v802 int32
	_ = v802
	var v805 int32
	_ = v805
	var v808 int32
	_ = v808
	var v810 int32
	_ = v810
	var v815 int32
	_ = v815
	var v822 int32
	_ = v822
	var v827 int32
	_ = v827
	var v832 int32
	_ = v832
	var v833 int32
	_ = v833
	var v841 int32
	_ = v841
	var v845 int32
	_ = v845
	var v847 int32
	_ = v847
	var v851 int32
	_ = v851
	var v852 int32
	_ = v852
	var v858 int32
	_ = v858
	var v860 int32
	_ = v860
	var v867 int32
	_ = v867
	var v869 int32
	_ = v869
	var v873 int32
	_ = v873
	var v874 int32
	_ = v874
	var v883 int32
	_ = v883
	var v884 int32
	_ = v884
	var v887 int32
	_ = v887
	var v889 int32
	_ = v889
	var v891 int32
	_ = v891
	var v900 int32
	_ = v900
	var v902 int32
	_ = v902
	var v905 int32
	_ = v905
	var v910 int32
	_ = v910
	var v915 int32
	_ = v915
	var v916 int32
	_ = v916
	var v921 int32
	_ = v921
	var __phi921 int32
	_ = __phi921
	var v922 int32
	_ = v922
	var __phi922 int32
	_ = __phi922
	var v928 int32
	_ = v928
	var v931 int32
	_ = v931
	var v946 int32
	_ = v946
	var v954 int32
	_ = v954
	var v956 int32
	_ = v956
	var v959 int32
	_ = v959
	var v962 int32
	_ = v962
	var v964 int32
	_ = v964
	var v969 int32
	_ = v969
	var v976 int32
	_ = v976
	var v981 int32
	_ = v981
	var v1004 int32
	_ = v1004
	var v1015 int32
	_ = v1015
	var v1017 int32
	_ = v1017
	var v1019 int32
	_ = v1019
	var v1023 int32
	_ = v1023
	var v1028 int32
	_ = v1028
	var v1029 int32
	_ = v1029
	var v1040 int32
	_ = v1040
	var v1043 int32
	_ = v1043
	var v1050 int32
	_ = v1050
	var v1055 int32
	_ = v1055
	var v1059 int32
	_ = v1059
	var v1061 int32
	_ = v1061
	var v1075 int32
	_ = v1075
	var v1077 int32
	_ = v1077
	var v1081 int32
	_ = v1081
	var v1083 int32
	_ = v1083
	var v1087 int32
	_ = v1087
	var v1097 int32
	_ = v1097
	var v1098 int32
	_ = v1098
	var v1114 int32
	_ = v1114
	var v1154 int32
	_ = v1154
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v14 = v12 & int32(-8)
	if v12&int32(3) != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	return v1154
L2:
	;
	v32 = l0 + v14
	if base.Ui32(v14) < base.Ui32(l1) {
		goto L9
	} else {
		goto L10
	}
L3:
	;
	if base.Ui32(l1) < base.Ui32(int32(256)) {
		v1154 = int32(0)
		goto L1
	} else {
		goto L4
	}
L4:
	;
	if base.Ui32(v14) < base.Ui32(l1+int32(4)) {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	return int32(0)
L6:
	;
	v25 = *(*int32)(unsafe.Add(mBase, _c_F_try_realloc_chunk[0]))
	if base.Ui32(v14-l1) <= base.Ui32(v25<<(uint(int32(1))%32)) {
		v1154 = l0
		goto L1
	} else {
		goto L7
	}
L7:
	;
	goto L5
L8:
	;
	v1154 = l0
	goto L1
L9:
	;
	v487 = int32(0)
	v489 = *(*int32)(unsafe.Add(mBase, _c_F_try_realloc_chunk[1]))
	if v32 != v489 {
		goto L104
	} else {
		goto L105
	}
L10:
	;
	v34 = v14 - l1
	if base.Ui32(v34) < base.Ui32(int32(16)) {
		goto L8
	} else {
		goto L11
	}
L11:
	;
	v37 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = l1 | v12&v37 | int32(2)
	v43 = l0 + l1
	*(*int32)(unsafe.Add(mBase, uint32(v43)+4)) = v34 | int32(3)
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v32)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v32)+4)) = v47 | v37
	v58 = v43 + v34
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v43)+4))
	if v59&v37 != 0 {
		v180 = v43
		v181 = v34
		goto L14
	} else {
		goto L15
	}
L12:
	;
	goto L8
L13:
	;
	goto L12
L14:
	;
	v189 = *(*int32)(unsafe.Add(mBase, uint32(v58)+4))
	if v189&int32(2) != 0 {
		goto L52
	} else {
		goto L53
	}
L15:
	;
	if v59&int32(2) == int32(0) {
		goto L13
	} else {
		goto L16
	}
L16:
	;
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v43)))
	v67 = v66 + v34
	v68 = v43 - v66
	v70 = *(*int32)(unsafe.Add(mBase, _c_F_try_realloc_chunk[2]))
	if v68 == v70 {
		goto L20
	} else {
		goto L21
	}
L17:
	;
	if v86 == int32(0) {
		v180 = v68
		v181 = v67
		goto L14
	} else {
		goto L36
	}
L18:
	;
	v140 = int32(0)
	goto L17
L19:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v75)+12)) = v72
	*(*int32)(unsafe.Add(mBase, uint32(v72)+8)) = v75
	v180 = v68
	v181 = v67
	goto L14
L20:
	;
	v120 = *(*int32)(unsafe.Add(mBase, uint32(v58)+4))
	v121 = int32(3)
	if v120&v121 != v121 {
		v180 = v68
		v181 = v67
		goto L14
	} else {
		goto L35
	}
L21:
	;
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v68)+12))
	if base.Ui32(int32(255)) < base.Ui32(v66) {
		goto L22
	} else {
		goto L23
	}
L22:
	;
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v68)+24))
	if v72 == v68 {
		goto L25
	} else {
		goto L26
	}
L23:
	;
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v68)+8))
	if v72 != v75 {
		goto L19
	} else {
		goto L24
	}
L24:
	;
	v77 = int32(0)
	v79 = *(*int32)(unsafe.Add(mBase, _c_F_try_realloc_chunk[3]))
	*(*int32)(unsafe.Add(mBase, _c_F_try_realloc_chunk[3])) = v79 & base.I32_rotl(int32(-2), int32(base.Ui32(v66)>>(uint(int32(3))%32)))
	v180 = v68
	v181 = v67
	goto L14
L25:
	;
	v91 = *(*int32)(unsafe.Add(mBase, uint32(v68)+20))
	if v91 == int32(0) {
		goto L28
	} else {
		goto L29
	}
L26:
	;
	v88 = *(*int32)(unsafe.Add(mBase, uint32(v68)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v88)+12)) = v72
	*(*int32)(unsafe.Add(mBase, uint32(v72)+8)) = v88
	v140 = v72
	goto L17
L27:
	;
	__phi107 = v101
	__phi108 = v102
	v107 = __phi107
	v108 = __phi108
	goto L31
L28:
	;
	v96 = *(*int32)(unsafe.Add(mBase, uint32(v68)+16))
	if v96 == int32(0) {
		goto L18
	} else {
		goto L30
	}
L29:
	;
	v101 = v91
	v102 = v68 + int32(20)
	goto L27
L30:
	;
	v101 = v96
	v102 = v68 + int32(16)
	goto L27
L31:
	;
	v114 = *(*int32)(unsafe.Add(mBase, uint32(v107)+20))
	if v114 != 0 {
		__phi107 = v114
		__phi108 = v107 + int32(20)
		v107 = __phi107
		v108 = __phi108
		goto L31
	} else {
		goto L33
	}
L32:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v108))) = int32(0)
	v140 = v107
	goto L17
L33:
	;
	v117 = *(*int32)(unsafe.Add(mBase, uint32(v107)+16))
	if v117 != 0 {
		__phi107 = v117
		__phi108 = v107 + int32(16)
		v107 = __phi107
		v108 = __phi108
		goto L31
	} else {
		goto L34
	}
L34:
	;
	goto L32
L35:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_try_realloc_chunk[4])) = v67
	*(*int32)(unsafe.Add(mBase, uint32(v58)+4)) = v120 & int32(-2)
	*(*int32)(unsafe.Add(mBase, uint32(v68)+4)) = v67 | int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v58))) = v67
	goto L12
L36:
	;
	v148 = *(*int32)(unsafe.Add(mBase, uint32(v68)+28))
	v150 = v148 << (uint(int32(2)) % 32)
	v153 = *(*int32)(unsafe.Add(mBase, uint32(v150)+uint32(_c_F_try_realloc_chunk[5])))
	if v68 != v153 {
		goto L38
	} else {
		goto L39
	}
L37:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v140)+24)) = v86
	v170 = *(*int32)(unsafe.Add(mBase, uint32(v68)+16))
	if v170 == int32(0) {
		goto L45
	} else {
		goto L46
	}
L38:
	;
	v163 = *(*int32)(unsafe.Add(mBase, uint32(v86)+16))
	if v163 != v68 {
		goto L42
	} else {
		goto L43
	}
L39:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v150)+uint32(_c_F_try_realloc_chunk[5]))) = v140
	if v140 != 0 {
		goto L37
	} else {
		goto L40
	}
L40:
	;
	v156 = int32(0)
	v158 = *(*int32)(unsafe.Add(mBase, _c_F_try_realloc_chunk[6]))
	*(*int32)(unsafe.Add(mBase, _c_F_try_realloc_chunk[6])) = v158 & base.I32_rotl(int32(-2), v148)
	v180 = v68
	v181 = v67
	goto L14
L41:
	;
	if v140 == int32(0) {
		v180 = v68
		v181 = v67
		goto L14
	} else {
		goto L44
	}
L42:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v86)+20)) = v140
	goto L41
L43:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v86)+16)) = v140
	goto L41
L44:
	;
	goto L37
L45:
	;
	v175 = *(*int32)(unsafe.Add(mBase, uint32(v68)+20))
	if v175 == int32(0) {
		v180 = v68
		v181 = v67
		goto L14
	} else {
		goto L47
	}
L46:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v140)+16)) = v170
	*(*int32)(unsafe.Add(mBase, uint32(v170)+24)) = v140
	goto L45
L47:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v140)+20)) = v175
	*(*int32)(unsafe.Add(mBase, uint32(v175)+24)) = v140
	v180 = v68
	v181 = v67
	goto L14
L48:
	;
	if base.Ui32(int32(255)) < base.Ui32(v352) {
		goto L86
	} else {
		goto L87
	}
L49:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v180)+4)) = v231 | int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v180+v231))) = v231
	if v180 != v215 {
		v352 = v231
		goto L48
	} else {
		goto L85
	}
L50:
	;
	if v248 == int32(0) {
		goto L49
	} else {
		goto L73
	}
L51:
	;
	v294 = int32(0)
	goto L50
L52:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v58)+4)) = v189 & int32(-2)
	*(*int32)(unsafe.Add(mBase, uint32(v180)+4)) = v181 | int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v180+v181))) = v181
	v352 = v181
	goto L48
L53:
	;
	v193 = *(*int32)(unsafe.Add(mBase, _c_F_try_realloc_chunk[1]))
	if v58 != v193 {
		goto L54
	} else {
		goto L55
	}
L54:
	;
	v215 = *(*int32)(unsafe.Add(mBase, _c_F_try_realloc_chunk[2]))
	if v58 != v215 {
		goto L57
	} else {
		goto L58
	}
L55:
	;
	v195 = int32(0)
	*(*int32)(unsafe.Add(mBase, _c_F_try_realloc_chunk[1])) = v180
	v199 = *(*int32)(unsafe.Add(mBase, _c_F_try_realloc_chunk[7]))
	v200 = v199 + v181
	*(*int32)(unsafe.Add(mBase, _c_F_try_realloc_chunk[7])) = v200
	*(*int32)(unsafe.Add(mBase, uint32(v180)+4)) = v200 | int32(1)
	v206 = *(*int32)(unsafe.Add(mBase, _c_F_try_realloc_chunk[2]))
	if v180 != v206 {
		goto L13
	} else {
		goto L56
	}
L56:
	;
	v208 = int32(0)
	*(*int32)(unsafe.Add(mBase, _c_F_try_realloc_chunk[4])) = v208
	*(*int32)(unsafe.Add(mBase, _c_F_try_realloc_chunk[2])) = v208
	goto L12
L57:
	;
	v231 = v189&int32(-8) + v181
	v232 = *(*int32)(unsafe.Add(mBase, uint32(v58)+12))
	if base.Ui32(int32(255)) < base.Ui32(v189) {
		goto L59
	} else {
		goto L60
	}
L58:
	;
	v217 = int32(0)
	*(*int32)(unsafe.Add(mBase, _c_F_try_realloc_chunk[2])) = v180
	v221 = *(*int32)(unsafe.Add(mBase, _c_F_try_realloc_chunk[4]))
	v222 = v221 + v181
	*(*int32)(unsafe.Add(mBase, _c_F_try_realloc_chunk[4])) = v222
	*(*int32)(unsafe.Add(mBase, uint32(v180)+4)) = v222 | int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v180+v222))) = v222
	goto L12
L59:
	;
	v248 = *(*int32)(unsafe.Add(mBase, uint32(v58)+24))
	if v232 == v58 {
		goto L63
	} else {
		goto L64
	}
L60:
	;
	v235 = *(*int32)(unsafe.Add(mBase, uint32(v58)+8))
	if v232 != v235 {
		goto L61
	} else {
		goto L62
	}
L61:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v235)+12)) = v232
	*(*int32)(unsafe.Add(mBase, uint32(v232)+8)) = v235
	goto L49
L62:
	;
	v237 = int32(0)
	v239 = *(*int32)(unsafe.Add(mBase, _c_F_try_realloc_chunk[3]))
	*(*int32)(unsafe.Add(mBase, _c_F_try_realloc_chunk[3])) = v239 & base.I32_rotl(int32(-2), int32(base.Ui32(v189)>>(uint(int32(3))%32)))
	goto L49
L63:
	;
	v253 = *(*int32)(unsafe.Add(mBase, uint32(v58)+20))
	if v253 == int32(0) {
		goto L66
	} else {
		goto L67
	}
L64:
	;
	v250 = *(*int32)(unsafe.Add(mBase, uint32(v58)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v250)+12)) = v232
	*(*int32)(unsafe.Add(mBase, uint32(v232)+8)) = v250
	v294 = v232
	goto L50
L65:
	;
	__phi269 = v263
	__phi270 = v264
	v269 = __phi269
	v270 = __phi270
	goto L69
L66:
	;
	v258 = *(*int32)(unsafe.Add(mBase, uint32(v58)+16))
	if v258 == int32(0) {
		goto L51
	} else {
		goto L68
	}
L67:
	;
	v263 = v253
	v264 = v58 + int32(20)
	goto L65
L68:
	;
	v263 = v258
	v264 = v58 + int32(16)
	goto L65
L69:
	;
	v276 = *(*int32)(unsafe.Add(mBase, uint32(v269)+20))
	if v276 != 0 {
		__phi269 = v276
		__phi270 = v269 + int32(20)
		v269 = __phi269
		v270 = __phi270
		goto L69
	} else {
		goto L71
	}
L70:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v270))) = int32(0)
	v294 = v269
	goto L50
L71:
	;
	v279 = *(*int32)(unsafe.Add(mBase, uint32(v269)+16))
	if v279 != 0 {
		__phi269 = v279
		__phi270 = v269 + int32(16)
		v269 = __phi269
		v270 = __phi270
		goto L69
	} else {
		goto L72
	}
L72:
	;
	goto L70
L73:
	;
	v302 = *(*int32)(unsafe.Add(mBase, uint32(v58)+28))
	v304 = v302 << (uint(int32(2)) % 32)
	v307 = *(*int32)(unsafe.Add(mBase, uint32(v304)+uint32(_c_F_try_realloc_chunk[5])))
	if v58 != v307 {
		goto L75
	} else {
		goto L76
	}
L74:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v294)+24)) = v248
	v324 = *(*int32)(unsafe.Add(mBase, uint32(v58)+16))
	if v324 == int32(0) {
		goto L82
	} else {
		goto L83
	}
L75:
	;
	v317 = *(*int32)(unsafe.Add(mBase, uint32(v248)+16))
	if v317 != v58 {
		goto L79
	} else {
		goto L80
	}
L76:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v304)+uint32(_c_F_try_realloc_chunk[5]))) = v294
	if v294 != 0 {
		goto L74
	} else {
		goto L77
	}
L77:
	;
	v310 = int32(0)
	v312 = *(*int32)(unsafe.Add(mBase, _c_F_try_realloc_chunk[6]))
	*(*int32)(unsafe.Add(mBase, _c_F_try_realloc_chunk[6])) = v312 & base.I32_rotl(int32(-2), v302)
	goto L49
L78:
	;
	if v294 == int32(0) {
		goto L49
	} else {
		goto L81
	}
L79:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v248)+20)) = v294
	goto L78
L80:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v248)+16)) = v294
	goto L78
L81:
	;
	goto L74
L82:
	;
	v329 = *(*int32)(unsafe.Add(mBase, uint32(v58)+20))
	if v329 == int32(0) {
		goto L49
	} else {
		goto L84
	}
L83:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v294)+16)) = v324
	*(*int32)(unsafe.Add(mBase, uint32(v324)+24)) = v294
	goto L82
L84:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v294)+20)) = v329
	*(*int32)(unsafe.Add(mBase, uint32(v329)+24)) = v294
	goto L49
L85:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_try_realloc_chunk[4])) = v231
	goto L12
L86:
	;
	if base.Ui32(int32(16777215)) < base.Ui32(v352) {
		v398 = int32(31)
		goto L91
	} else {
		goto L92
	}
L87:
	;
	v363 = v352 & int32(-8)
	v365 = v363 + int32(9128464)
	v367 = *(*int32)(unsafe.Add(mBase, _c_F_try_realloc_chunk[3]))
	v371 = int32(1) << (uint(int32(base.Ui32(v352)>>(uint(int32(3))%32))) % 32)
	if v367&v371 != 0 {
		goto L89
	} else {
		goto L90
	}
L88:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v363)+uint32(_c_F_try_realloc_chunk[8]))) = v180
	*(*int32)(unsafe.Add(mBase, uint32(v377)+12)) = v180
	*(*int32)(unsafe.Add(mBase, uint32(v180)+12)) = v365
	*(*int32)(unsafe.Add(mBase, uint32(v180)+8)) = v377
	goto L12
L89:
	;
	v376 = *(*int32)(unsafe.Add(mBase, uint32(v363)+uint32(_c_F_try_realloc_chunk[8])))
	v377 = v376
	goto L88
L90:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_try_realloc_chunk[3])) = v367 | v371
	v377 = v365
	goto L88
L91:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v180)+28)) = v398
	*(*int64)(unsafe.Add(mBase, uint32(v180)+16)) = int64(0)
	v403 = v398 << (uint(int32(2)) % 32)
	v407 = *(*int32)(unsafe.Add(mBase, _c_F_try_realloc_chunk[6]))
	v409 = int32(1) << (uint(v398) % 32)
	if v407&v409 != 0 {
		goto L95
	} else {
		goto L96
	}
L92:
	;
	v388 = base.I32_clz(int32(base.Ui32(v352) >> (uint(int32(8)) % 32)))
	v391 = int32(1)
	v398 = int32(base.Ui32(v352)>>(uint(int32(38)-v388)%32))&v391 - v388<<(uint(v391)%32) + int32(62)
	goto L91
L93:
	;
	v462 = *(*int32)(unsafe.Add(mBase, uint32(v431)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v462)+12)) = v180
	*(*int32)(unsafe.Add(mBase, uint32(v431)+8)) = v180
	*(*int32)(unsafe.Add(mBase, uint32(v180)+24)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v180)+12)) = v431
	*(*int32)(unsafe.Add(mBase, uint32(v180)+8)) = v462
	goto L13
L94:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v180)+12)) = v180
	*(*int32)(unsafe.Add(mBase, uint32(v180)+8)) = v180
	goto L12
L95:
	;
	if v398 == int32(31) {
		goto L97
	} else {
		goto L98
	}
L96:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_try_realloc_chunk[6])) = v407 | v409
	*(*int32)(unsafe.Add(mBase, uint32(v403)+uint32(_c_F_try_realloc_chunk[5]))) = v180
	*(*int32)(unsafe.Add(mBase, uint32(v180)+24)) = v403 + int32(9128728)
	goto L94
L97:
	;
	v423 = int32(0)
	goto L99
L98:
	;
	v423 = int32(25) - int32(base.Ui32(v398)>>(uint(int32(1))%32))
	goto L99
L99:
	;
	v425 = *(*int32)(unsafe.Add(mBase, uint32(v403)+uint32(_c_F_try_realloc_chunk[5])))
	v429 = v352 << (uint(v423) % 32)
	v431 = v425
	goto L100
L100:
	;
	v435 = *(*int32)(unsafe.Add(mBase, uint32(v431)+4))
	if v435&int32(-8) == v352 {
		goto L93
	} else {
		goto L102
	}
L101:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v445+int32(16)))) = v180
	*(*int32)(unsafe.Add(mBase, uint32(v180)+24)) = v431
	goto L94
L102:
	;
	v445 = v431 + int32(base.Ui32(v429)>>(uint(int32(29))%32))&int32(4)
	v446 = *(*int32)(unsafe.Add(mBase, uint32(v445)+16))
	if v446 != 0 {
		v429 = v429 << (uint(int32(1)) % 32)
		v431 = v446
		goto L100
	} else {
		goto L103
	}
L103:
	;
	goto L101
L104:
	;
	v511 = *(*int32)(unsafe.Add(mBase, _c_F_try_realloc_chunk[2]))
	if v32 != v511 {
		goto L107
	} else {
		goto L108
	}
L105:
	;
	v492 = *(*int32)(unsafe.Add(mBase, _c_F_try_realloc_chunk[7]))
	v493 = v492 + v14
	if base.Ui32(v493) <= base.Ui32(l1) {
		v1154 = v487
		goto L1
	} else {
		goto L106
	}
L106:
	;
	v495 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = l1 | v12&v495 | int32(2)
	v501 = l0 + l1
	v502 = v493 - l1
	*(*int32)(unsafe.Add(mBase, uint32(v501)+4)) = v502 | v495
	v506 = int32(0)
	*(*int32)(unsafe.Add(mBase, _c_F_try_realloc_chunk[7])) = v502
	*(*int32)(unsafe.Add(mBase, _c_F_try_realloc_chunk[1])) = v501
	goto L8
L107:
	;
	v557 = int32(0)
	v558 = *(*int32)(unsafe.Add(mBase, uint32(v32)+4))
	if v558&int32(2) != 0 {
		v1154 = v557
		goto L1
	} else {
		goto L113
	}
L108:
	;
	v513 = int32(0)
	v515 = *(*int32)(unsafe.Add(mBase, _c_F_try_realloc_chunk[4]))
	v516 = v515 + v14
	if base.Ui32(v516) < base.Ui32(l1) {
		v1154 = v513
		goto L1
	} else {
		goto L109
	}
L109:
	;
	v518 = v516 - l1
	if base.Ui32(v518) < base.Ui32(int32(16)) {
		goto L111
	} else {
		goto L112
	}
L110:
	;
	v553 = int32(0)
	*(*int32)(unsafe.Add(mBase, _c_F_try_realloc_chunk[2])) = v550
	*(*int32)(unsafe.Add(mBase, _c_F_try_realloc_chunk[4])) = v552
	goto L8
L111:
	;
	v537 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v12&v537 | v516 | int32(2)
	v543 = l0 + v516
	v544 = *(*int32)(unsafe.Add(mBase, uint32(v543)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v543)+4)) = v544 | v537
	v548 = int32(0)
	v550 = v548
	v552 = v548
	goto L110
L112:
	;
	v521 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = l1 | v12&v521 | int32(2)
	v527 = l0 + l1
	*(*int32)(unsafe.Add(mBase, uint32(v527)+4)) = v518 | v521
	v531 = l0 + v516
	*(*int32)(unsafe.Add(mBase, uint32(v531))) = v518
	v533 = *(*int32)(unsafe.Add(mBase, uint32(v531)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v531)+4)) = v533 & int32(-2)
	v550 = v527
	v552 = v518
	goto L110
L113:
	;
	v563 = v558&int32(-8) + v14
	if base.Ui32(v563) < base.Ui32(l1) {
		v1154 = v557
		goto L1
	} else {
		goto L114
	}
L114:
	;
	v565 = v563 - l1
	v566 = *(*int32)(unsafe.Add(mBase, uint32(v32)+12))
	if base.Ui32(int32(255)) < base.Ui32(v558) {
		goto L116
	} else {
		goto L117
	}
L115:
	;
	if base.Ui32(int32(15)) < base.Ui32(v565) {
		goto L144
	} else {
		goto L145
	}
L116:
	;
	v582 = *(*int32)(unsafe.Add(mBase, uint32(v32)+24))
	if v566 == v32 {
		goto L121
	} else {
		goto L122
	}
L117:
	;
	v569 = *(*int32)(unsafe.Add(mBase, uint32(v32)+8))
	if v566 != v569 {
		goto L118
	} else {
		goto L119
	}
L118:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v569)+12)) = v566
	*(*int32)(unsafe.Add(mBase, uint32(v566)+8)) = v569
	goto L115
L119:
	;
	v571 = int32(0)
	v573 = *(*int32)(unsafe.Add(mBase, _c_F_try_realloc_chunk[3]))
	*(*int32)(unsafe.Add(mBase, _c_F_try_realloc_chunk[3])) = v573 & base.I32_rotl(int32(-2), int32(base.Ui32(v558)>>(uint(int32(3))%32)))
	goto L115
L120:
	;
	if v582 == int32(0) {
		goto L115
	} else {
		goto L132
	}
L121:
	;
	v587 = *(*int32)(unsafe.Add(mBase, uint32(v32)+20))
	if v587 == int32(0) {
		goto L125
	} else {
		goto L126
	}
L122:
	;
	v584 = *(*int32)(unsafe.Add(mBase, uint32(v32)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v584)+12)) = v566
	*(*int32)(unsafe.Add(mBase, uint32(v566)+8)) = v584
	v622 = v566
	goto L120
L123:
	;
	v622 = int32(0)
	goto L120
L124:
	;
	__phi603 = v597
	__phi605 = v598
	v603 = __phi603
	v605 = __phi605
	goto L128
L125:
	;
	v592 = *(*int32)(unsafe.Add(mBase, uint32(v32)+16))
	if v592 == int32(0) {
		goto L123
	} else {
		goto L127
	}
L126:
	;
	v597 = v587
	v598 = v32 + int32(20)
	goto L124
L127:
	;
	v597 = v592
	v598 = v32 + int32(16)
	goto L124
L128:
	;
	v612 = *(*int32)(unsafe.Add(mBase, uint32(v603)+20))
	if v612 != 0 {
		__phi603 = v612
		__phi605 = v603 + int32(20)
		v603 = __phi603
		v605 = __phi605
		goto L128
	} else {
		goto L130
	}
L129:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v605))) = int32(0)
	v622 = v603
	goto L120
L130:
	;
	v615 = *(*int32)(unsafe.Add(mBase, uint32(v603)+16))
	if v615 != 0 {
		__phi603 = v615
		__phi605 = v603 + int32(16)
		v603 = __phi603
		v605 = __phi605
		goto L128
	} else {
		goto L131
	}
L131:
	;
	goto L129
L132:
	;
	v632 = *(*int32)(unsafe.Add(mBase, uint32(v32)+28))
	v634 = v632 << (uint(int32(2)) % 32)
	v637 = *(*int32)(unsafe.Add(mBase, uint32(v634)+uint32(_c_F_try_realloc_chunk[5])))
	if v32 != v637 {
		goto L134
	} else {
		goto L135
	}
L133:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v622)+24)) = v582
	v654 = *(*int32)(unsafe.Add(mBase, uint32(v32)+16))
	if v654 == int32(0) {
		goto L141
	} else {
		goto L142
	}
L134:
	;
	v647 = *(*int32)(unsafe.Add(mBase, uint32(v582)+16))
	if v647 != v32 {
		goto L138
	} else {
		goto L139
	}
L135:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v634)+uint32(_c_F_try_realloc_chunk[5]))) = v622
	if v622 != 0 {
		goto L133
	} else {
		goto L136
	}
L136:
	;
	v640 = int32(0)
	v642 = *(*int32)(unsafe.Add(mBase, _c_F_try_realloc_chunk[6]))
	*(*int32)(unsafe.Add(mBase, _c_F_try_realloc_chunk[6])) = v642 & base.I32_rotl(int32(-2), v632)
	goto L115
L137:
	;
	if v622 == int32(0) {
		goto L115
	} else {
		goto L140
	}
L138:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v582)+20)) = v622
	goto L137
L139:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v582)+16)) = v622
	goto L137
L140:
	;
	goto L133
L141:
	;
	v659 = *(*int32)(unsafe.Add(mBase, uint32(v32)+20))
	if v659 == int32(0) {
		goto L115
	} else {
		goto L143
	}
L142:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v622)+16)) = v654
	*(*int32)(unsafe.Add(mBase, uint32(v654)+24)) = v622
	goto L141
L143:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v622)+20)) = v659
	*(*int32)(unsafe.Add(mBase, uint32(v659)+24)) = v622
	goto L115
L144:
	;
	v688 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = l1 | v12&v688 | int32(2)
	v694 = l0 + l1
	*(*int32)(unsafe.Add(mBase, uint32(v694)+4)) = v565 | int32(3)
	v698 = l0 + v563
	v699 = *(*int32)(unsafe.Add(mBase, uint32(v698)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v698)+4)) = v699 | v688
	v710 = v694 + v565
	v711 = *(*int32)(unsafe.Add(mBase, uint32(v694)+4))
	if v711&v688 != 0 {
		v832 = v694
		v833 = v565
		goto L148
	} else {
		goto L149
	}
L145:
	;
	v677 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v12&v677 | v563 | int32(2)
	v683 = l0 + v563
	v684 = *(*int32)(unsafe.Add(mBase, uint32(v683)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v683)+4)) = v684 | v677
	goto L8
L146:
	;
	goto L8
L147:
	;
	goto L146
L148:
	;
	v841 = *(*int32)(unsafe.Add(mBase, uint32(v710)+4))
	if v841&int32(2) != 0 {
		goto L186
	} else {
		goto L187
	}
L149:
	;
	if v711&int32(2) == int32(0) {
		goto L147
	} else {
		goto L150
	}
L150:
	;
	v718 = *(*int32)(unsafe.Add(mBase, uint32(v694)))
	v719 = v718 + v565
	v720 = v694 - v718
	v722 = *(*int32)(unsafe.Add(mBase, _c_F_try_realloc_chunk[2]))
	if v720 == v722 {
		goto L154
	} else {
		goto L155
	}
L151:
	;
	if v738 == int32(0) {
		v832 = v720
		v833 = v719
		goto L148
	} else {
		goto L170
	}
L152:
	;
	v792 = int32(0)
	goto L151
L153:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v727)+12)) = v724
	*(*int32)(unsafe.Add(mBase, uint32(v724)+8)) = v727
	v832 = v720
	v833 = v719
	goto L148
L154:
	;
	v772 = *(*int32)(unsafe.Add(mBase, uint32(v710)+4))
	v773 = int32(3)
	if v772&v773 != v773 {
		v832 = v720
		v833 = v719
		goto L148
	} else {
		goto L169
	}
L155:
	;
	v724 = *(*int32)(unsafe.Add(mBase, uint32(v720)+12))
	if base.Ui32(int32(255)) < base.Ui32(v718) {
		goto L156
	} else {
		goto L157
	}
L156:
	;
	v738 = *(*int32)(unsafe.Add(mBase, uint32(v720)+24))
	if v724 == v720 {
		goto L159
	} else {
		goto L160
	}
L157:
	;
	v727 = *(*int32)(unsafe.Add(mBase, uint32(v720)+8))
	if v724 != v727 {
		goto L153
	} else {
		goto L158
	}
L158:
	;
	v729 = int32(0)
	v731 = *(*int32)(unsafe.Add(mBase, _c_F_try_realloc_chunk[3]))
	*(*int32)(unsafe.Add(mBase, _c_F_try_realloc_chunk[3])) = v731 & base.I32_rotl(int32(-2), int32(base.Ui32(v718)>>(uint(int32(3))%32)))
	v832 = v720
	v833 = v719
	goto L148
L159:
	;
	v743 = *(*int32)(unsafe.Add(mBase, uint32(v720)+20))
	if v743 == int32(0) {
		goto L162
	} else {
		goto L163
	}
L160:
	;
	v740 = *(*int32)(unsafe.Add(mBase, uint32(v720)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v740)+12)) = v724
	*(*int32)(unsafe.Add(mBase, uint32(v724)+8)) = v740
	v792 = v724
	goto L151
L161:
	;
	__phi759 = v753
	__phi760 = v754
	v759 = __phi759
	v760 = __phi760
	goto L165
L162:
	;
	v748 = *(*int32)(unsafe.Add(mBase, uint32(v720)+16))
	if v748 == int32(0) {
		goto L152
	} else {
		goto L164
	}
L163:
	;
	v753 = v743
	v754 = v720 + int32(20)
	goto L161
L164:
	;
	v753 = v748
	v754 = v720 + int32(16)
	goto L161
L165:
	;
	v766 = *(*int32)(unsafe.Add(mBase, uint32(v759)+20))
	if v766 != 0 {
		__phi759 = v766
		__phi760 = v759 + int32(20)
		v759 = __phi759
		v760 = __phi760
		goto L165
	} else {
		goto L167
	}
L166:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v760))) = int32(0)
	v792 = v759
	goto L151
L167:
	;
	v769 = *(*int32)(unsafe.Add(mBase, uint32(v759)+16))
	if v769 != 0 {
		__phi759 = v769
		__phi760 = v759 + int32(16)
		v759 = __phi759
		v760 = __phi760
		goto L165
	} else {
		goto L168
	}
L168:
	;
	goto L166
L169:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_try_realloc_chunk[4])) = v719
	*(*int32)(unsafe.Add(mBase, uint32(v710)+4)) = v772 & int32(-2)
	*(*int32)(unsafe.Add(mBase, uint32(v720)+4)) = v719 | int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v710))) = v719
	goto L146
L170:
	;
	v800 = *(*int32)(unsafe.Add(mBase, uint32(v720)+28))
	v802 = v800 << (uint(int32(2)) % 32)
	v805 = *(*int32)(unsafe.Add(mBase, uint32(v802)+uint32(_c_F_try_realloc_chunk[5])))
	if v720 != v805 {
		goto L172
	} else {
		goto L173
	}
L171:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v792)+24)) = v738
	v822 = *(*int32)(unsafe.Add(mBase, uint32(v720)+16))
	if v822 == int32(0) {
		goto L179
	} else {
		goto L180
	}
L172:
	;
	v815 = *(*int32)(unsafe.Add(mBase, uint32(v738)+16))
	if v815 != v720 {
		goto L176
	} else {
		goto L177
	}
L173:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v802)+uint32(_c_F_try_realloc_chunk[5]))) = v792
	if v792 != 0 {
		goto L171
	} else {
		goto L174
	}
L174:
	;
	v808 = int32(0)
	v810 = *(*int32)(unsafe.Add(mBase, _c_F_try_realloc_chunk[6]))
	*(*int32)(unsafe.Add(mBase, _c_F_try_realloc_chunk[6])) = v810 & base.I32_rotl(int32(-2), v800)
	v832 = v720
	v833 = v719
	goto L148
L175:
	;
	if v792 == int32(0) {
		v832 = v720
		v833 = v719
		goto L148
	} else {
		goto L178
	}
L176:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v738)+20)) = v792
	goto L175
L177:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v738)+16)) = v792
	goto L175
L178:
	;
	goto L171
L179:
	;
	v827 = *(*int32)(unsafe.Add(mBase, uint32(v720)+20))
	if v827 == int32(0) {
		v832 = v720
		v833 = v719
		goto L148
	} else {
		goto L181
	}
L180:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v792)+16)) = v822
	*(*int32)(unsafe.Add(mBase, uint32(v822)+24)) = v792
	goto L179
L181:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v792)+20)) = v827
	*(*int32)(unsafe.Add(mBase, uint32(v827)+24)) = v792
	v832 = v720
	v833 = v719
	goto L148
L182:
	;
	if base.Ui32(int32(255)) < base.Ui32(v1004) {
		goto L220
	} else {
		goto L221
	}
L183:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v832)+4)) = v883 | int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v832+v883))) = v883
	if v832 != v867 {
		v1004 = v883
		goto L182
	} else {
		goto L219
	}
L184:
	;
	if v900 == int32(0) {
		goto L183
	} else {
		goto L207
	}
L185:
	;
	v946 = int32(0)
	goto L184
L186:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v710)+4)) = v841 & int32(-2)
	*(*int32)(unsafe.Add(mBase, uint32(v832)+4)) = v833 | int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v832+v833))) = v833
	v1004 = v833
	goto L182
L187:
	;
	v845 = *(*int32)(unsafe.Add(mBase, _c_F_try_realloc_chunk[1]))
	if v710 != v845 {
		goto L188
	} else {
		goto L189
	}
L188:
	;
	v867 = *(*int32)(unsafe.Add(mBase, _c_F_try_realloc_chunk[2]))
	if v710 != v867 {
		goto L191
	} else {
		goto L192
	}
L189:
	;
	v847 = int32(0)
	*(*int32)(unsafe.Add(mBase, _c_F_try_realloc_chunk[1])) = v832
	v851 = *(*int32)(unsafe.Add(mBase, _c_F_try_realloc_chunk[7]))
	v852 = v851 + v833
	*(*int32)(unsafe.Add(mBase, _c_F_try_realloc_chunk[7])) = v852
	*(*int32)(unsafe.Add(mBase, uint32(v832)+4)) = v852 | int32(1)
	v858 = *(*int32)(unsafe.Add(mBase, _c_F_try_realloc_chunk[2]))
	if v832 != v858 {
		goto L147
	} else {
		goto L190
	}
L190:
	;
	v860 = int32(0)
	*(*int32)(unsafe.Add(mBase, _c_F_try_realloc_chunk[4])) = v860
	*(*int32)(unsafe.Add(mBase, _c_F_try_realloc_chunk[2])) = v860
	goto L146
L191:
	;
	v883 = v841&int32(-8) + v833
	v884 = *(*int32)(unsafe.Add(mBase, uint32(v710)+12))
	if base.Ui32(int32(255)) < base.Ui32(v841) {
		goto L193
	} else {
		goto L194
	}
L192:
	;
	v869 = int32(0)
	*(*int32)(unsafe.Add(mBase, _c_F_try_realloc_chunk[2])) = v832
	v873 = *(*int32)(unsafe.Add(mBase, _c_F_try_realloc_chunk[4]))
	v874 = v873 + v833
	*(*int32)(unsafe.Add(mBase, _c_F_try_realloc_chunk[4])) = v874
	*(*int32)(unsafe.Add(mBase, uint32(v832)+4)) = v874 | int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v832+v874))) = v874
	goto L146
L193:
	;
	v900 = *(*int32)(unsafe.Add(mBase, uint32(v710)+24))
	if v884 == v710 {
		goto L197
	} else {
		goto L198
	}
L194:
	;
	v887 = *(*int32)(unsafe.Add(mBase, uint32(v710)+8))
	if v884 != v887 {
		goto L195
	} else {
		goto L196
	}
L195:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v887)+12)) = v884
	*(*int32)(unsafe.Add(mBase, uint32(v884)+8)) = v887
	goto L183
L196:
	;
	v889 = int32(0)
	v891 = *(*int32)(unsafe.Add(mBase, _c_F_try_realloc_chunk[3]))
	*(*int32)(unsafe.Add(mBase, _c_F_try_realloc_chunk[3])) = v891 & base.I32_rotl(int32(-2), int32(base.Ui32(v841)>>(uint(int32(3))%32)))
	goto L183
L197:
	;
	v905 = *(*int32)(unsafe.Add(mBase, uint32(v710)+20))
	if v905 == int32(0) {
		goto L200
	} else {
		goto L201
	}
L198:
	;
	v902 = *(*int32)(unsafe.Add(mBase, uint32(v710)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v902)+12)) = v884
	*(*int32)(unsafe.Add(mBase, uint32(v884)+8)) = v902
	v946 = v884
	goto L184
L199:
	;
	__phi921 = v915
	__phi922 = v916
	v921 = __phi921
	v922 = __phi922
	goto L203
L200:
	;
	v910 = *(*int32)(unsafe.Add(mBase, uint32(v710)+16))
	if v910 == int32(0) {
		goto L185
	} else {
		goto L202
	}
L201:
	;
	v915 = v905
	v916 = v710 + int32(20)
	goto L199
L202:
	;
	v915 = v910
	v916 = v710 + int32(16)
	goto L199
L203:
	;
	v928 = *(*int32)(unsafe.Add(mBase, uint32(v921)+20))
	if v928 != 0 {
		__phi921 = v928
		__phi922 = v921 + int32(20)
		v921 = __phi921
		v922 = __phi922
		goto L203
	} else {
		goto L205
	}
L204:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v922))) = int32(0)
	v946 = v921
	goto L184
L205:
	;
	v931 = *(*int32)(unsafe.Add(mBase, uint32(v921)+16))
	if v931 != 0 {
		__phi921 = v931
		__phi922 = v921 + int32(16)
		v921 = __phi921
		v922 = __phi922
		goto L203
	} else {
		goto L206
	}
L206:
	;
	goto L204
L207:
	;
	v954 = *(*int32)(unsafe.Add(mBase, uint32(v710)+28))
	v956 = v954 << (uint(int32(2)) % 32)
	v959 = *(*int32)(unsafe.Add(mBase, uint32(v956)+uint32(_c_F_try_realloc_chunk[5])))
	if v710 != v959 {
		goto L209
	} else {
		goto L210
	}
L208:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v946)+24)) = v900
	v976 = *(*int32)(unsafe.Add(mBase, uint32(v710)+16))
	if v976 == int32(0) {
		goto L216
	} else {
		goto L217
	}
L209:
	;
	v969 = *(*int32)(unsafe.Add(mBase, uint32(v900)+16))
	if v969 != v710 {
		goto L213
	} else {
		goto L214
	}
L210:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v956)+uint32(_c_F_try_realloc_chunk[5]))) = v946
	if v946 != 0 {
		goto L208
	} else {
		goto L211
	}
L211:
	;
	v962 = int32(0)
	v964 = *(*int32)(unsafe.Add(mBase, _c_F_try_realloc_chunk[6]))
	*(*int32)(unsafe.Add(mBase, _c_F_try_realloc_chunk[6])) = v964 & base.I32_rotl(int32(-2), v954)
	goto L183
L212:
	;
	if v946 == int32(0) {
		goto L183
	} else {
		goto L215
	}
L213:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v900)+20)) = v946
	goto L212
L214:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v900)+16)) = v946
	goto L212
L215:
	;
	goto L208
L216:
	;
	v981 = *(*int32)(unsafe.Add(mBase, uint32(v710)+20))
	if v981 == int32(0) {
		goto L183
	} else {
		goto L218
	}
L217:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v946)+16)) = v976
	*(*int32)(unsafe.Add(mBase, uint32(v976)+24)) = v946
	goto L216
L218:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v946)+20)) = v981
	*(*int32)(unsafe.Add(mBase, uint32(v981)+24)) = v946
	goto L183
L219:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_try_realloc_chunk[4])) = v883
	goto L146
L220:
	;
	if base.Ui32(int32(16777215)) < base.Ui32(v1004) {
		v1050 = int32(31)
		goto L225
	} else {
		goto L226
	}
L221:
	;
	v1015 = v1004 & int32(-8)
	v1017 = v1015 + int32(9128464)
	v1019 = *(*int32)(unsafe.Add(mBase, _c_F_try_realloc_chunk[3]))
	v1023 = int32(1) << (uint(int32(base.Ui32(v1004)>>(uint(int32(3))%32))) % 32)
	if v1019&v1023 != 0 {
		goto L223
	} else {
		goto L224
	}
L222:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1015)+uint32(_c_F_try_realloc_chunk[8]))) = v832
	*(*int32)(unsafe.Add(mBase, uint32(v1029)+12)) = v832
	*(*int32)(unsafe.Add(mBase, uint32(v832)+12)) = v1017
	*(*int32)(unsafe.Add(mBase, uint32(v832)+8)) = v1029
	goto L146
L223:
	;
	v1028 = *(*int32)(unsafe.Add(mBase, uint32(v1015)+uint32(_c_F_try_realloc_chunk[8])))
	v1029 = v1028
	goto L222
L224:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_try_realloc_chunk[3])) = v1019 | v1023
	v1029 = v1017
	goto L222
L225:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v832)+28)) = v1050
	*(*int64)(unsafe.Add(mBase, uint32(v832)+16)) = int64(0)
	v1055 = v1050 << (uint(int32(2)) % 32)
	v1059 = *(*int32)(unsafe.Add(mBase, _c_F_try_realloc_chunk[6]))
	v1061 = int32(1) << (uint(v1050) % 32)
	if v1059&v1061 != 0 {
		goto L229
	} else {
		goto L230
	}
L226:
	;
	v1040 = base.I32_clz(int32(base.Ui32(v1004) >> (uint(int32(8)) % 32)))
	v1043 = int32(1)
	v1050 = int32(base.Ui32(v1004)>>(uint(int32(38)-v1040)%32))&v1043 - v1040<<(uint(v1043)%32) + int32(62)
	goto L225
L227:
	;
	v1114 = *(*int32)(unsafe.Add(mBase, uint32(v1083)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v1114)+12)) = v832
	*(*int32)(unsafe.Add(mBase, uint32(v1083)+8)) = v832
	*(*int32)(unsafe.Add(mBase, uint32(v832)+24)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v832)+12)) = v1083
	*(*int32)(unsafe.Add(mBase, uint32(v832)+8)) = v1114
	goto L147
L228:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v832)+12)) = v832
	*(*int32)(unsafe.Add(mBase, uint32(v832)+8)) = v832
	goto L146
L229:
	;
	if v1050 == int32(31) {
		goto L231
	} else {
		goto L232
	}
L230:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_try_realloc_chunk[6])) = v1059 | v1061
	*(*int32)(unsafe.Add(mBase, uint32(v1055)+uint32(_c_F_try_realloc_chunk[5]))) = v832
	*(*int32)(unsafe.Add(mBase, uint32(v832)+24)) = v1055 + int32(9128728)
	goto L228
L231:
	;
	v1075 = int32(0)
	goto L233
L232:
	;
	v1075 = int32(25) - int32(base.Ui32(v1050)>>(uint(int32(1))%32))
	goto L233
L233:
	;
	v1077 = *(*int32)(unsafe.Add(mBase, uint32(v1055)+uint32(_c_F_try_realloc_chunk[5])))
	v1081 = v1004 << (uint(v1075) % 32)
	v1083 = v1077
	goto L234
L234:
	;
	v1087 = *(*int32)(unsafe.Add(mBase, uint32(v1083)+4))
	if v1087&int32(-8) == v1004 {
		goto L227
	} else {
		goto L236
	}
L235:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1097+int32(16)))) = v832
	*(*int32)(unsafe.Add(mBase, uint32(v832)+24)) = v1083
	goto L228
L236:
	;
	v1097 = v1083 + int32(base.Ui32(v1081)>>(uint(int32(29))%32))&int32(4)
	v1098 = *(*int32)(unsafe.Add(mBase, uint32(v1097)+16))
	if v1098 != 0 {
		v1081 = v1081 << (uint(int32(1)) % 32)
		v1083 = v1098
		goto L234
	} else {
		goto L237
	}
L237:
	;
	goto L235
}
func F_ttlCommand(m *base.Module, l0 int32) {
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
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v14 int64
	_ = v14
	var v18 int32
	_ = v18
	var v28 int64
	_ = v28
	var v29 int64
	_ = v29
	var v33 int64
	_ = v33
	var v34 int64
	_ = v34
	var v35 int64
	_ = v35
	var v38 int64
	_ = v38
	var v42 int64
	_ = v42
	var v43 int64
	_ = v43
	var v45 int32
	_ = v45
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(v6)+4))
	v9 = F_lookupKeyReadWithFlags(m, v5, v7, int32(1))
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		return
	} else {
		if v9 != 0 {
			v14 = int64(-1)
			v18 = *(*int32)(unsafe.Add(mBase, uint32(v9)+4))
			if v18&int32(1) == int32(0) {
				v29 = v14
			} else {
				v28 = *(*int64)(unsafe.Add(mBase, uint32(v9+(v18&int32(4)^int32(12)))))
				v29 = v28
			}
			if v29 == int64(-1) {
				v43 = v14
			} else {
				v33 = *(*int64)(unsafe.Add(mBase, _c_F_ttlCommand[0]))
				v34 = v29 - v33
				v35 = int64(0)
				if v35 < v34 {
					v38 = v34
				} else {
					v38 = v35
				}
				v42 = base.I64_div_u_s(v38+int64(500), int64(1000))
				v43 = v42
			}
			F_addReplyLongLong(m, l0, v43)
			mBase = m.M
			v45 = m.ExcPending
			if v45 != 0 {
				return
			} else {
				return
			}
		} else {
			F_addReplyLongLong(m, l0, int64(-2))
			mBase = m.M
			v13 = m.ExcPending
			if v13 != 0 {
				return
			} else {
				return
			}
		}
	}
}
