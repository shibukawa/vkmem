package p1

import (
	base "github.com/shibukawa/vkmem/internal/aot/vkaot/base"
	"unsafe"
)

func F_GCTM(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
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
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v57 int32
	_ = v57
	var v58 int64
	_ = v58
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v66 int32
	_ = v66
	var v72 int32
	_ = v72
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v9 = *(*int32)(unsafe.Add(mBase, uint32(v8)+48))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(v9)))
	if v10 != v9 {
		v14 = *(*int32)(unsafe.Add(mBase, uint32(v10)))
		*(*int32)(unsafe.Add(mBase, uint32(v9))) = v14
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(v8)+48)) = int32(0)
	}
	v16 = *(*int32)(unsafe.Add(mBase, uint32(v8)+112))
	v17 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
	*(*int32)(unsafe.Add(mBase, uint32(v10))) = v17
	v19 = *(*int32)(unsafe.Add(mBase, uint32(v8)+112))
	*(*int32)(unsafe.Add(mBase, uint32(v19))) = v10
	v21 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+20)))
	v24 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+5)))
	v27 = v21&int32(3) | v24&int32(248)
	*(*uint8)(unsafe.Add(mBase, uint32(v10)+5)) = uint8(v27)
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v10)+8))
	if v29 == int32(0) {
		return
	} else {
		v32 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v29)+6)))
		if v32&int32(4) != 0 {
			return
		} else {
			v36 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
			v37 = *(*int32)(unsafe.Add(mBase, uint32(v36)+196))
			v38 = F_luaH_getstr(m, v29, v37)
			mBase = m.M
			v39 = *(*int32)(unsafe.Add(mBase, uint32(v38)+8))
			if v39 != 0 {
				v46 = v38
			} else {
				v40 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v29)+6)))
				v43 = v40 | int32(4)
				*(*uint8)(unsafe.Add(mBase, uint32(v29)+6)) = uint8(v43)
				v46 = int32(0)
			}
			if v46 == int32(0) {
				return
			} else {
				v49 = *(*int32)(unsafe.Add(mBase, uint32(v8)+64))
				v50 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+57)))
				v51 = int32(0)
				*(*uint8)(unsafe.Add(mBase, uint32(l0)+57)) = uint8(v51)
				v53 = *(*int32)(unsafe.Add(mBase, uint32(v8)+68))
				*(*int32)(unsafe.Add(mBase, uint32(v8)+64)) = v53 << (uint(int32(1)) % 32)
				v57 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				v58 = *(*int64)(unsafe.Add(mBase, uint32(v46)))
				*(*int64)(unsafe.Add(mBase, uint32(v57))) = v58
				v60 = *(*int32)(unsafe.Add(mBase, uint32(v46)+8))
				*(*int32)(unsafe.Add(mBase, uint32(v57)+8)) = v60
				v62 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				*(*int32)(unsafe.Add(mBase, uint32(v62)+24)) = int32(7)
				*(*int32)(unsafe.Add(mBase, uint32(v62)+16)) = v10
				v66 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v66 + int32(32)
				F_luaD_call(m, l0, v66, v51)
				mBase = m.M
				v72 = m.ExcPending
				if v72 != 0 {
					return
				} else {
					*(*uint8)(unsafe.Add(mBase, uint32(l0)+57)) = uint8(v50)
					*(*int32)(unsafe.Add(mBase, uint32(v8)+64)) = v49
					return
				}
			}
		}
	}
}
func F___gettimeofday(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 float64
	_ = v7
	var v9 float64
	_ = v9
	var v15 int64
	_ = v15
	var v17 int64
	_ = v17
	var v24 float64
	_ = v24
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	v7 = m.Env.Emscripten_date_now(m)
	mBase = m.M
	v9 = base.F64_div(v7, float64(1000))
	if base.F64_lt(base.F64_abs(v9), float64(9.223372036854776e+18)) == int32(0) {
		v17 = int64(-9223372036854775807 - 1)
	} else {
		v15 = base.I64_trunc_f64_s(v9)
		v17 = v15
	}
	*(*int64)(unsafe.Add(mBase, uint32(l0))) = v17
	v24 = base.F64_mul(base.F64_sub(v7, base.F64_convert_i64_s(v17*int64(1000))), float64(1000))
	if base.F64_lt(base.F64_abs(v24), float64(2.147483648e+09)) == int32(0) {
		v32 = int32(-2147483648)
	} else {
		v30 = base.I32_trunc_f64_s(v24)
		v32 = v30
	}
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v32
	return int32(0)
}
func F___gmtime_r(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v8 int64
	_ = v8
	v3 = int32(9116396)
	F___lock(m, v3)
	mBase = m.M
	F_do_tzset(m)
	mBase = m.M
	F___unlock(m, v3)
	mBase = m.M
	v8 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	m.Env.X_gmtime_js(m, v8, l1)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(l1)+40)) = int32(_a2768)
	*(*int64)(unsafe.Add(mBase, uint32(l1)+32)) = int64(0)
	return l1
}
func F_gai_strerror(m *base.Module, l0 int32) int32 {
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
	var v16 int32
	_ = v16
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	v4 = int32(_a116)
	v6 = l0 + int32(1)
	if v6 == int32(0) {
		v26 = v4
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v28 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v26))))
	return v26 + base.B2i32(v28 == int32(0))
L2:
	;
	v10 = v4
	v11 = v6
	goto L3
L3:
	;
	v12 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10))))
	if v12 == int32(0) {
		v26 = v10
		goto L1
	} else {
		goto L5
	}
L4:
	;
	v26 = v22
	goto L1
L5:
	;
	v16 = v10
	goto L6
L6:
	;
	v20 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16)+1)))
	if v20 != 0 {
		v16 = v16 + int32(1)
		goto L6
	} else {
		goto L8
	}
L7:
	;
	v22 = v16 + int32(2)
	v24 = v11 + int32(1)
	if v24 != 0 {
		v10 = v22
		v11 = v24
		goto L3
	} else {
		goto L9
	}
L8:
	;
	goto L7
L9:
	;
	goto L4
}
func F_genrand64_int64(m *base.Module) int64 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v18 int64
	_ = v18
	var v22 int64
	_ = v22
	var v23 int64
	_ = v23
	var v31 int64
	_ = v31
	var v33 int64
	_ = v33
	var v37 int32
	_ = v37
	var v41 int64
	_ = v41
	var v44 int64
	_ = v44
	var v46 int64
	_ = v46
	var v49 int64
	_ = v49
	var v60 int64
	_ = v60
	var v63 int64
	_ = v63
	var v74 int64
	_ = v74
	var v77 int64
	_ = v77
	var v90 int64
	_ = v90
	var v97 int64
	_ = v97
	var v102 int32
	_ = v102
	var v105 int64
	_ = v105
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v121 int64
	_ = v121
	var v129 int64
	_ = v129
	var v132 int64
	_ = v132
	var v144 int64
	_ = v144
	var v146 int32
	_ = v146
	var v148 int64
	_ = v148
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v165 int64
	_ = v165
	var v173 int64
	_ = v173
	var v176 int64
	_ = v176
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v190 int64
	_ = v190
	var v198 int64
	_ = v198
	var v200 int64
	_ = v200
	var v205 int64
	_ = v205
	var v214 int32
	_ = v214
	var v215 int64
	_ = v215
	var v226 int64
	_ = v226
	var v231 int64
	_ = v231
	var v236 int64
	_ = v236
	v9 = *(*int32)(unsafe.Add(mBase, _consts[466]))
	if int32(311) < v9 {
		if v9 == int32(313) {
			v23 = int64(5489)
			*(*int64)(unsafe.Add(mBase, _consts[467])) = v23
			v31 = int64(1)
			v33 = v23
			for {
				v37 = int32(3)
				v41 = int64(62)
				v44 = int64(6364136223846793005)
				v46 = (int64(base.Ui64(v33)>>(uint(v41)%64))^v33)*v44 + v31
				*(*int64)(unsafe.Add(mBase, uint32(base.I32_wrap_i64(v31)<<(uint(v37)%32))+uint32(_consts[467]))) = v46
				v49 = v31 + int64(1)
				v60 = (int64(base.Ui64(v46)>>(uint(v41)%64))^v46)*v44 + v49
				*(*int64)(unsafe.Add(mBase, uint32(base.I32_wrap_i64(v49)<<(uint(v37)%32))+uint32(_consts[467]))) = v60
				v63 = v31 + int64(2)
				v74 = (int64(base.Ui64(v60)>>(uint(v41)%64))^v60)*v44 + v63
				*(*int64)(unsafe.Add(mBase, uint32(base.I32_wrap_i64(v63)<<(uint(v37)%32))+uint32(_consts[467]))) = v74
				v77 = v31 + int64(3)
				if v77 == int64(312) {
					v97 = v23
					break
				} else {
					v90 = (int64(base.Ui64(v74)>>(uint(int64(62))%64))^v74)*int64(6364136223846793005) + v77
					*(*int64)(unsafe.Add(mBase, uint32(base.I32_wrap_i64(v77)<<(uint(int32(3))%32))+uint32(_consts[467]))) = v90
					v31 = v31 + int64(4)
					v33 = v90
					continue
				}
				break
			}
		} else {
			v22 = *(*int64)(unsafe.Add(mBase, _consts[467]))
			v97 = v22
		}
		v102 = int32(0)
		v105 = v97
		for {
			v111 = int32(3)
			v112 = v102 << (uint(v111) % 32)
			v115 = int32(1)
			v116 = v102 + v115
			v121 = *(*int64)(unsafe.Add(mBase, uint32(v116<<(uint(v111)%32))+uint32(_consts[467])))
			v129 = *(*int64)(unsafe.Add(mBase, uint32(base.I32_wrap_i64(v121)&v115<<(uint(v111)%32))+uint32(_consts[468])))
			v132 = *(*int64)(unsafe.Add(mBase, uint32(v112)+uint32(_consts[469])))
			*(*int64)(unsafe.Add(mBase, uint32(v112)+uint32(_consts[467]))) = v129 ^ v132 ^ int64(base.Ui64(v105&int64(-2147483648)|v121&int64(2147483646))>>(uint(int64(1))%64))
			if v116 != int32(156) {
				v102 = v116
				v105 = v121
				continue
			} else {
				break
			}
			break
		}
		v144 = *(*int64)(unsafe.Add(mBase, _consts[469]))
		v146 = int32(156)
		v148 = v144
		for {
			v155 = int32(3)
			v156 = v146 << (uint(v155) % 32)
			v159 = int32(1)
			v160 = v146 + v159
			v165 = *(*int64)(unsafe.Add(mBase, uint32(v160<<(uint(v155)%32))+uint32(_consts[467])))
			v173 = *(*int64)(unsafe.Add(mBase, uint32(base.I32_wrap_i64(v165)&v159<<(uint(v155)%32))+uint32(_consts[468])))
			v176 = *(*int64)(unsafe.Add(mBase, uint32(v156)+uint32(_consts[470])))
			*(*int64)(unsafe.Add(mBase, uint32(v156)+uint32(_consts[467]))) = v173 ^ v176 ^ int64(base.Ui64(v148&int64(-2147483648)|v165&int64(2147483646))>>(uint(int64(1))%64))
			if v160 != int32(311) {
				v146 = v160
				v148 = v165
				continue
			} else {
				break
			}
			break
		}
		v187 = int32(1)
		v188 = int32(0)
		v190 = *(*int64)(unsafe.Add(mBase, _consts[467]))
		v198 = *(*int64)(unsafe.Add(mBase, uint32(base.I32_wrap_i64(v190)&v187<<(uint(int32(3))%32))+uint32(_consts[468])))
		v200 = *(*int64)(unsafe.Add(mBase, _consts[471]))
		v205 = *(*int64)(unsafe.Add(mBase, _consts[472]))
		*(*int64)(unsafe.Add(mBase, _consts[472])) = v198 ^ v200 ^ int64(base.Ui64(v190&int64(2147483646)|v205&int64(-2147483648))>>(uint(int64(1))%64))
		v214 = v187
		v215 = v190
	} else {
		v18 = *(*int64)(unsafe.Add(mBase, uint32(v9<<(uint(int32(3))%32))+uint32(_consts[467])))
		v214 = v9 + int32(1)
		v215 = v18
	}
	*(*int32)(unsafe.Add(mBase, _consts[466])) = v214
	v226 = int64(base.Ui64(v215)>>(uint(int64(29))%64))&int64(22906492245) ^ v215
	v231 = v226<<(uint(int64(17))%64)&int64(8202884508482404352) ^ v226
	v236 = v231<<(uint(int64(37))%64)&int64(-2270628950310912) ^ v231
	return int64(base.Ui64(v236)>>(uint(int64(43))%64)) ^ v236
}
func F_geoArrayCleanup(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v22 int32
	_ = v22
	var v27 int32
	_ = v27
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v3 == int32(0) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v22 == l0+int32(16) {
		goto L8
	} else {
		goto L9
	}
L2:
	;
	v8 = int32(0)
	goto L3
L3:
	;
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v13 = *(*int32)(unsafe.Add(mBase, uint32(v9+v8*int32(40))+32))
	F_sdsfree(m, v13)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		goto L5
	} else {
		goto L6
	}
L4:
	;
	goto L1
L5:
	;
	return
L6:
	;
	v17 = v8 + int32(1)
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if base.Ui32(v17) < base.Ui32(v18) {
		v8 = v17
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
L9:
	;
	F_valkey_free(m, v22)
	mBase = m.M
	v27 = m.ExcPending
	if v27 != 0 {
		goto L5
	} else {
		goto L10
	}
L10:
	;
	goto L8
}
func F_geodistCommand(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 float64
	_ = v16
	var v17 int32
	_ = v17
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v27 float64
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
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
	var v62 int32
	_ = v62
	var v63 float64
	_ = v63
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v72 int32
	_ = v72
	var v80 int64
	_ = v80
	var v89 int64
	_ = v89
	var v91 int64
	_ = v91
	var v94 int32
	_ = v94
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	var v107 int64
	_ = v107
	var v109 int64
	_ = v109
	var v111 int32
	_ = v111
	var v117 float64
	_ = v117
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v126 int32
	_ = v126
	var v132 int64
	_ = v132
	var v141 int64
	_ = v141
	var v143 int64
	_ = v143
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v153 int32
	_ = v153
	var v157 int64
	_ = v157
	var v159 int64
	_ = v159
	var v161 int32
	_ = v161
	var v169 int32
	_ = v169
	var v170 float64
	_ = v170
	var v171 float64
	_ = v171
	var v172 float64
	_ = v172
	var v173 float64
	_ = v173
	var v175 float64
	_ = v175
	var v182 float64
	_ = v182
	var v188 float64
	_ = v188
	var v196 float64
	_ = v196
	var v197 float64
	_ = v197
	var v198 float64
	_ = v198
	var v200 float64
	_ = v200
	var v201 float64
	_ = v201
	var v205 float64
	_ = v205
	var v212 float64
	_ = v212
	var v219 float64
	_ = v219
	var v222 int32
	_ = v222
	v7 = m.G0
	v9 = v7 - int32(96)
	m.G0 = v9
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v11 != int32(5) {
		if v11 < int32(6) {
			v27 = float64(1)
			v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
			v29 = *(*int32)(unsafe.Add(mBase, uint32(v28)+4))
			v31 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+224)))
			v35 = *(*int32)(unsafe.Add(mBase, uint32(v31<<(uint(int32(2))%32))+uint32(_consts[339])))
			v36 = F_lookupKeyReadOrReply(m, l0, v29, v35)
			mBase = m.M
			v37 = m.ExcPending
			if v37 != 0 {
				return
			} else {
				if v36 == int32(0) {
					m.G0 = v9 + int32(96)
					return
				} else {
					v41 = F_checkType(m, l0, v36, int32(3))
					mBase = m.M
					v42 = m.ExcPending
					if v42 != 0 {
						return
					} else {
						if v41 != 0 {
							m.G0 = v9 + int32(96)
							return
						} else {
							v43 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
							v44 = *(*int32)(unsafe.Add(mBase, uint32(v43)+8))
							v45 = F_objectGetVal(m, v44)
							mBase = m.M
							v48 = F_zsetScore(m, v36, v45, v9+int32(72))
							mBase = m.M
							v49 = m.ExcPending
							if v49 != 0 {
								return
							} else {
								if v48 == int32(-1) {
									F_addReplyNull(m, l0)
									mBase = m.M
									v62 = m.ExcPending
									if v62 != 0 {
										return
									} else {
										m.G0 = v9 + int32(96)
										return
									}
								} else {
									v52 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
									v53 = *(*int32)(unsafe.Add(mBase, uint32(v52)+12))
									v54 = F_objectGetVal(m, v53)
									mBase = m.M
									v57 = F_zsetScore(m, v36, v54, v9+int32(64))
									mBase = m.M
									v58 = m.ExcPending
									if v58 != 0 {
										return
									} else {
										if v57 != int32(-1) {
											v63 = *(*float64)(unsafe.Add(mBase, uint32(v9)+72))
											v67 = v9 + int32(88)
											v68 = int32(26)
											*(*uint8)(unsafe.Add(mBase, uint32(v67))) = uint8(v68)
											v72 = int32(0)
											*(*int32)(unsafe.Add(mBase, uint32(v9+int32(92)))) = v72
											*(*int32)(unsafe.Add(mBase, uint32(v9)+89)) = v72
											v80 = *(*int64)(unsafe.Add(mBase, uint32(v67)))
											*(*int64)(unsafe.Add(mBase, uint32(v9+int32(24)))) = v80
											if base.F64_lt(v63, float64(1.8446744073709552e+19))&base.F64_ge(v63, float64(0)) == v72 {
												v91 = int64(0)
											} else {
												v89 = base.I64_trunc_f64_u(v63)
												v91 = v89
											}
											*(*int64)(unsafe.Add(mBase, uint32(v9)+80)) = v91
											*(*int64)(unsafe.Add(mBase, uint32(v9)+16)) = v91
											v94 = int32(16)
											v99 = m.G0
											v101 = v99 - v94
											m.G0 = v101
											v107 = *(*int64)(unsafe.Add(mBase, uint32(v9+int32(24))))
											*(*int64)(unsafe.Add(mBase, uint32(v101+int32(8)))) = v107
											v109 = *(*int64)(unsafe.Add(mBase, uint32(v9+v94)))
											*(*int64)(unsafe.Add(mBase, uint32(v101))) = v109
											v111 = F_geohashDecodeToLongLatType(m, v101, v9+int32(32))
											mBase = m.M
											m.G0 = v101 + v94
											if v111 == int32(0) {
												F_addReplyNull(m, l0)
												mBase = m.M
												v169 = m.ExcPending
												if v169 != 0 {
													return
												} else {
													m.G0 = v9 + int32(96)
													return
												}
											} else {
												v117 = *(*float64)(unsafe.Add(mBase, uint32(v9)+64))
												v121 = v9 + int32(88)
												v122 = int32(26)
												*(*uint8)(unsafe.Add(mBase, uint32(v121))) = uint8(v122)
												v126 = int32(0)
												*(*int32)(unsafe.Add(mBase, uint32(v9+int32(92)))) = v126
												*(*int32)(unsafe.Add(mBase, uint32(v9)+89)) = v126
												v132 = *(*int64)(unsafe.Add(mBase, uint32(v121)))
												*(*int64)(unsafe.Add(mBase, uint32(v9+int32(8)))) = v132
												if base.F64_lt(v117, float64(1.8446744073709552e+19))&base.F64_ge(v117, float64(0)) == v126 {
													v143 = int64(0)
												} else {
													v141 = base.I64_trunc_f64_u(v117)
													v143 = v141
												}
												*(*int64)(unsafe.Add(mBase, uint32(v9)+80)) = v143
												*(*int64)(unsafe.Add(mBase, uint32(v9))) = v143
												v149 = m.G0
												v150 = int32(16)
												v151 = v149 - v150
												m.G0 = v151
												v153 = int32(8)
												v157 = *(*int64)(unsafe.Add(mBase, uint32(v9+v153)))
												*(*int64)(unsafe.Add(mBase, uint32(v151+v153))) = v157
												v159 = *(*int64)(unsafe.Add(mBase, uint32(v9)))
												*(*int64)(unsafe.Add(mBase, uint32(v151))) = v159
												v161 = F_geohashDecodeToLongLatType(m, v151, v9+int32(48))
												mBase = m.M
												m.G0 = v151 + v150
												if v161 != 0 {
													v170 = *(*float64)(unsafe.Add(mBase, uint32(v9)+32))
													v171 = *(*float64)(unsafe.Add(mBase, uint32(v9)+40))
													v172 = *(*float64)(unsafe.Add(mBase, uint32(v9)+48))
													v173 = *(*float64)(unsafe.Add(mBase, uint32(v9)+56))
													v175 = float64(0.017453292519943295)
													v182 = F_sin(m, base.F64_mul(base.F64_sub(base.F64_mul(v172, v175), base.F64_mul(v170, v175)), float64(0.5)))
													mBase = m.M
													if base.F64_le(base.F64_abs(v182), float64(1e-15)) == int32(0) {
														v196 = float64(0.017453292519943295)
														v197 = base.F64_mul(v171, v196)
														v198 = F_cos(m, v197)
														mBase = m.M
														v200 = base.F64_mul(v173, v196)
														v201 = F_cos(m, v200)
														mBase = m.M
														v205 = F_sin(m, base.F64_mul(base.F64_sub(v200, v197), float64(0.5)))
														mBase = m.M
														v212 = F_asin(m, base.F64_sqrt(base.F64_add(base.F64_mul(v205, v205), base.F64_mul(v182, base.F64_mul(base.F64_mul(v198, v201), v182)))))
														mBase = m.M
														v219 = base.F64_mul(v212, float64(1.2745595121712e+07))
													} else {
														v188 = float64(0.017453292519943295)
														v219 = base.F64_mul(base.F64_abs(base.F64_sub(base.F64_mul(v173, v188), base.F64_mul(v171, v188))), float64(6.372797560856e+06))
													}
													F_addReplyDoubleDistance(m, l0, base.F64_div(v219, v27))
													mBase = m.M
													v222 = m.ExcPending
													if v222 != 0 {
														return
													} else {
														m.G0 = v9 + int32(96)
														return
													}
												} else {
													F_addReplyNull(m, l0)
													mBase = m.M
													v169 = m.ExcPending
													if v169 != 0 {
														return
													} else {
														m.G0 = v9 + int32(96)
														return
													}
												}
											}
										} else {
											F_addReplyNull(m, l0)
											mBase = m.M
											v62 = m.ExcPending
											if v62 != 0 {
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
		} else {
			v24 = *(*int32)(unsafe.Add(mBase, _consts[33]))
			F_addReplyErrorObject(m, l0, v24)
			mBase = m.M
			v26 = m.ExcPending
			if v26 != 0 {
				return
			} else {
				m.G0 = v9 + int32(96)
				return
			}
		}
	} else {
		v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
		v15 = *(*int32)(unsafe.Add(mBase, uint32(v14)+16))
		v16 = F_extractUnitOrReply(m, l0, v15)
		mBase = m.M
		v17 = m.ExcPending
		if v17 != 0 {
			return
		} else {
			if base.F64_lt(v16, float64(0)) != 0 {
				m.G0 = v9 + int32(96)
				return
			} else {
				v27 = v16
				v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
				v29 = *(*int32)(unsafe.Add(mBase, uint32(v28)+4))
				v31 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+224)))
				v35 = *(*int32)(unsafe.Add(mBase, uint32(v31<<(uint(int32(2))%32))+uint32(_consts[339])))
				v36 = F_lookupKeyReadOrReply(m, l0, v29, v35)
				mBase = m.M
				v37 = m.ExcPending
				if v37 != 0 {
					return
				} else {
					if v36 == int32(0) {
						m.G0 = v9 + int32(96)
						return
					} else {
						v41 = F_checkType(m, l0, v36, int32(3))
						mBase = m.M
						v42 = m.ExcPending
						if v42 != 0 {
							return
						} else {
							if v41 != 0 {
								m.G0 = v9 + int32(96)
								return
							} else {
								v43 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
								v44 = *(*int32)(unsafe.Add(mBase, uint32(v43)+8))
								v45 = F_objectGetVal(m, v44)
								mBase = m.M
								v48 = F_zsetScore(m, v36, v45, v9+int32(72))
								mBase = m.M
								v49 = m.ExcPending
								if v49 != 0 {
									return
								} else {
									if v48 == int32(-1) {
										F_addReplyNull(m, l0)
										mBase = m.M
										v62 = m.ExcPending
										if v62 != 0 {
											return
										} else {
											m.G0 = v9 + int32(96)
											return
										}
									} else {
										v52 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
										v53 = *(*int32)(unsafe.Add(mBase, uint32(v52)+12))
										v54 = F_objectGetVal(m, v53)
										mBase = m.M
										v57 = F_zsetScore(m, v36, v54, v9+int32(64))
										mBase = m.M
										v58 = m.ExcPending
										if v58 != 0 {
											return
										} else {
											if v57 != int32(-1) {
												v63 = *(*float64)(unsafe.Add(mBase, uint32(v9)+72))
												v67 = v9 + int32(88)
												v68 = int32(26)
												*(*uint8)(unsafe.Add(mBase, uint32(v67))) = uint8(v68)
												v72 = int32(0)
												*(*int32)(unsafe.Add(mBase, uint32(v9+int32(92)))) = v72
												*(*int32)(unsafe.Add(mBase, uint32(v9)+89)) = v72
												v80 = *(*int64)(unsafe.Add(mBase, uint32(v67)))
												*(*int64)(unsafe.Add(mBase, uint32(v9+int32(24)))) = v80
												if base.F64_lt(v63, float64(1.8446744073709552e+19))&base.F64_ge(v63, float64(0)) == v72 {
													v91 = int64(0)
												} else {
													v89 = base.I64_trunc_f64_u(v63)
													v91 = v89
												}
												*(*int64)(unsafe.Add(mBase, uint32(v9)+80)) = v91
												*(*int64)(unsafe.Add(mBase, uint32(v9)+16)) = v91
												v94 = int32(16)
												v99 = m.G0
												v101 = v99 - v94
												m.G0 = v101
												v107 = *(*int64)(unsafe.Add(mBase, uint32(v9+int32(24))))
												*(*int64)(unsafe.Add(mBase, uint32(v101+int32(8)))) = v107
												v109 = *(*int64)(unsafe.Add(mBase, uint32(v9+v94)))
												*(*int64)(unsafe.Add(mBase, uint32(v101))) = v109
												v111 = F_geohashDecodeToLongLatType(m, v101, v9+int32(32))
												mBase = m.M
												m.G0 = v101 + v94
												if v111 == int32(0) {
													F_addReplyNull(m, l0)
													mBase = m.M
													v169 = m.ExcPending
													if v169 != 0 {
														return
													} else {
														m.G0 = v9 + int32(96)
														return
													}
												} else {
													v117 = *(*float64)(unsafe.Add(mBase, uint32(v9)+64))
													v121 = v9 + int32(88)
													v122 = int32(26)
													*(*uint8)(unsafe.Add(mBase, uint32(v121))) = uint8(v122)
													v126 = int32(0)
													*(*int32)(unsafe.Add(mBase, uint32(v9+int32(92)))) = v126
													*(*int32)(unsafe.Add(mBase, uint32(v9)+89)) = v126
													v132 = *(*int64)(unsafe.Add(mBase, uint32(v121)))
													*(*int64)(unsafe.Add(mBase, uint32(v9+int32(8)))) = v132
													if base.F64_lt(v117, float64(1.8446744073709552e+19))&base.F64_ge(v117, float64(0)) == v126 {
														v143 = int64(0)
													} else {
														v141 = base.I64_trunc_f64_u(v117)
														v143 = v141
													}
													*(*int64)(unsafe.Add(mBase, uint32(v9)+80)) = v143
													*(*int64)(unsafe.Add(mBase, uint32(v9))) = v143
													v149 = m.G0
													v150 = int32(16)
													v151 = v149 - v150
													m.G0 = v151
													v153 = int32(8)
													v157 = *(*int64)(unsafe.Add(mBase, uint32(v9+v153)))
													*(*int64)(unsafe.Add(mBase, uint32(v151+v153))) = v157
													v159 = *(*int64)(unsafe.Add(mBase, uint32(v9)))
													*(*int64)(unsafe.Add(mBase, uint32(v151))) = v159
													v161 = F_geohashDecodeToLongLatType(m, v151, v9+int32(48))
													mBase = m.M
													m.G0 = v151 + v150
													if v161 != 0 {
														v170 = *(*float64)(unsafe.Add(mBase, uint32(v9)+32))
														v171 = *(*float64)(unsafe.Add(mBase, uint32(v9)+40))
														v172 = *(*float64)(unsafe.Add(mBase, uint32(v9)+48))
														v173 = *(*float64)(unsafe.Add(mBase, uint32(v9)+56))
														v175 = float64(0.017453292519943295)
														v182 = F_sin(m, base.F64_mul(base.F64_sub(base.F64_mul(v172, v175), base.F64_mul(v170, v175)), float64(0.5)))
														mBase = m.M
														if base.F64_le(base.F64_abs(v182), float64(1e-15)) == int32(0) {
															v196 = float64(0.017453292519943295)
															v197 = base.F64_mul(v171, v196)
															v198 = F_cos(m, v197)
															mBase = m.M
															v200 = base.F64_mul(v173, v196)
															v201 = F_cos(m, v200)
															mBase = m.M
															v205 = F_sin(m, base.F64_mul(base.F64_sub(v200, v197), float64(0.5)))
															mBase = m.M
															v212 = F_asin(m, base.F64_sqrt(base.F64_add(base.F64_mul(v205, v205), base.F64_mul(v182, base.F64_mul(base.F64_mul(v198, v201), v182)))))
															mBase = m.M
															v219 = base.F64_mul(v212, float64(1.2745595121712e+07))
														} else {
															v188 = float64(0.017453292519943295)
															v219 = base.F64_mul(base.F64_abs(base.F64_sub(base.F64_mul(v173, v188), base.F64_mul(v171, v188))), float64(6.372797560856e+06))
														}
														F_addReplyDoubleDistance(m, l0, base.F64_div(v219, v27))
														mBase = m.M
														v222 = m.ExcPending
														if v222 != 0 {
															return
														} else {
															m.G0 = v9 + int32(96)
															return
														}
													} else {
														F_addReplyNull(m, l0)
														mBase = m.M
														v169 = m.ExcPending
														if v169 != 0 {
															return
														} else {
															m.G0 = v9 + int32(96)
															return
														}
													}
												}
											} else {
												F_addReplyNull(m, l0)
												mBase = m.M
												v62 = m.ExcPending
												if v62 != 0 {
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
func F_geoposCommand(m *base.Module, l0 int32) {
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
	var v40 int32
	_ = v40
	var v46 int32
	_ = v46
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v59 int32
	_ = v59
	var v60 float64
	_ = v60
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v74 int64
	_ = v74
	var v76 int64
	_ = v76
	var v78 int32
	_ = v78
	var v85 int32
	_ = v85
	var v86 int64
	_ = v86
	var v88 int64
	_ = v88
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v103 int64
	_ = v103
	var v105 int64
	_ = v105
	var v107 int32
	_ = v107
	var v112 int32
	_ = v112
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v118 float64
	_ = v118
	var v125 int32
	_ = v125
	var v127 int32
	_ = v127
	var v129 int64
	_ = v129
	var v131 int64
	_ = v131
	var v135 int64
	_ = v135
	var v155 int64
	_ = v155
	var v169 int32
	_ = v169
	var v178 int64
	_ = v178
	var v181 int64
	_ = v181
	var v182 int64
	_ = v182
	var v183 int64
	_ = v183
	var v184 int64
	_ = v184
	var v197 int64
	_ = v197
	var v202 int64
	_ = v202
	var v204 int32
	_ = v204
	var v205 float64
	_ = v205
	var v212 int32
	_ = v212
	var v214 int32
	_ = v214
	var v216 int64
	_ = v216
	var v218 int64
	_ = v218
	var v222 int64
	_ = v222
	var v242 int64
	_ = v242
	var v256 int32
	_ = v256
	var v265 int64
	_ = v265
	var v268 int64
	_ = v268
	var v269 int64
	_ = v269
	var v270 int64
	_ = v270
	var v271 int64
	_ = v271
	var v284 int64
	_ = v284
	var v287 int64
	_ = v287
	var v289 int32
	_ = v289
	var v294 int32
	_ = v294
	var v295 int32
	_ = v295
	v10 = m.G0
	v12 = v10 - int32(96)
	m.G0 = v12
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v16 = *(*int32)(unsafe.Add(mBase, uint32(v15)+4))
	v17 = F_lookupKeyRead(m, v14, v16)
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v12 + int32(96)
	return
L2:
	;
	return
L3:
	;
	v20 = F_checkType(m, l0, v17, int32(3))
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		goto L2
	} else {
		goto L4
	}
L4:
	;
	if v20 != 0 {
		goto L1
	} else {
		goto L5
	}
L5:
	;
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	F_addReplyArrayLen(m, l0, v22+int32(-2))
	mBase = m.M
	v26 = m.ExcPending
	if v26 != 0 {
		goto L2
	} else {
		goto L6
	}
L6:
	;
	v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v27 < int32(3) {
		goto L1
	} else {
		goto L7
	}
L7:
	;
	v40 = int32(2)
	goto L8
L8:
	;
	if v17 == int32(0) {
		goto L12
	} else {
		goto L13
	}
L9:
	;
	goto L1
L10:
	;
	v294 = v40 + int32(1)
	v295 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v294 < v295 {
		v40 = v294
		goto L8
	} else {
		goto L49
	}
L11:
	;
	v60 = *(*float64)(unsafe.Add(mBase, uint32(v12)+72))
	v64 = v12 + int32(88)
	v65 = int32(26)
	*(*uint8)(unsafe.Add(mBase, uint32(v64))) = uint8(v65)
	if base.F64_lt(v60, float64(1.8446744073709552e+19))&base.F64_ge(v60, float64(0)) == int32(0) {
		goto L18
	} else {
		goto L19
	}
L12:
	;
	F_addReplyNullArray(m, l0)
	mBase = m.M
	v59 = m.ExcPending
	if v59 != 0 {
		goto L2
	} else {
		goto L16
	}
L13:
	;
	v46 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v46+v40<<(uint(int32(2))%32))))
	v51 = F_objectGetVal(m, v50)
	mBase = m.M
	v54 = F_zsetScore(m, v17, v51, v12+int32(72))
	mBase = m.M
	v55 = m.ExcPending
	if v55 != 0 {
		goto L2
	} else {
		goto L14
	}
L14:
	;
	if v54 != int32(-1) {
		goto L11
	} else {
		goto L15
	}
L15:
	;
	goto L12
L16:
	;
	goto L10
L17:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v12)+80)) = v76
	v78 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v12+int32(92)))) = v78
	*(*int32)(unsafe.Add(mBase, uint32(v12+int32(89)))) = v78
	v85 = v12 + int32(40)
	v86 = *(*int64)(unsafe.Add(mBase, uint32(v64)))
	*(*int64)(unsafe.Add(mBase, uint32(v85))) = v86
	v88 = *(*int64)(unsafe.Add(mBase, uint32(v12)+80))
	*(*int64)(unsafe.Add(mBase, uint32(v12)+32)) = v88
	v95 = m.G0
	v96 = int32(16)
	v97 = v95 - v96
	m.G0 = v97
	v103 = *(*int64)(unsafe.Add(mBase, uint32(v85)))
	*(*int64)(unsafe.Add(mBase, uint32(v97+int32(8)))) = v103
	v105 = *(*int64)(unsafe.Add(mBase, uint32(v12+int32(32))))
	*(*int64)(unsafe.Add(mBase, uint32(v97))) = v105
	v107 = F_geohashDecodeToLongLatType(m, v97, v12+int32(48))
	mBase = m.M
	m.G0 = v97 + v96
	goto L21
L18:
	;
	v76 = int64(0)
	goto L17
L19:
	;
	v74 = base.I64_trunc_f64_u(v60)
	v76 = v74
	goto L17
L20:
	;
	F_addReplyArrayLen(m, l0, int32(2))
	mBase = m.M
	v115 = m.ExcPending
	if v115 != 0 {
		goto L2
	} else {
		goto L24
	}
L21:
	;
	if v107 != 0 {
		goto L20
	} else {
		goto L22
	}
L22:
	;
	F_addReplyNullArray(m, l0)
	mBase = m.M
	v112 = m.ExcPending
	if v112 != 0 {
		goto L2
	} else {
		goto L23
	}
L23:
	;
	goto L10
L24:
	;
	v116 = int32(16)
	v117 = v12 + v116
	v118 = *(*float64)(unsafe.Add(mBase, uint32(v12)+48))
	v125 = m.G0
	v127 = v125 - v116
	m.G0 = v127
	v129 = base.I64_reinterpret_f64(v118)
	v131 = v129 & int64(4503599627370495)
	v135 = int64(base.Ui64(v129)>>(uint(int64(52))%64)) & int64(2047)
	if v135 == int64(0) {
		goto L27
	} else {
		goto L28
	}
L25:
	;
	v197 = *(*int64)(unsafe.Add(mBase, uint32(v12)+16))
	v202 = *(*int64)(unsafe.Add(mBase, uint32(v12+int32(24))))
	F_addReplyHumanLongDouble(m, l0, v197, v202)
	mBase = m.M
	v204 = m.ExcPending
	if v204 != 0 {
		goto L2
	} else {
		goto L36
	}
L26:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v117))) = v182
	*(*int64)(unsafe.Add(mBase, uint32(v117)+8)) = v183<<(uint(int64(48))%64) | v129&int64(-9223372036854775807-1) | v184
	m.G0 = v127 + int32(16)
	goto L25
L27:
	;
	if base.B2i32(v131 == int64(0)) == int32(0) {
		goto L31
	} else {
		goto L32
	}
L28:
	;
	if v135 == int64(2047) {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	v182 = v131 << (uint(int64(60)) % 64)
	v183 = int64(32767)
	v184 = int64(base.Ui64(v131) >> (uint(int64(4)) % 64))
	goto L26
L30:
	;
	v182 = v131 << (uint(int64(60)) % 64)
	v183 = v135 + int64(15360)
	v184 = int64(base.Ui64(v131) >> (uint(int64(4)) % 64))
	goto L26
L31:
	;
	if base.Ui64(v131) < base.Ui64(int64(4294967296)) {
		goto L33
	} else {
		goto L34
	}
L32:
	;
	v155 = int64(0)
	v182 = v155
	v183 = v155
	v184 = v155
	goto L26
L33:
	;
	v169 = base.I32_clz(base.I32_wrap_i64(v129)) | int32(32)
	goto L35
L34:
	;
	v169 = base.I32_clz(base.I32_wrap_i64(int64(base.Ui64(v131) >> (uint(int64(32)) % 64))))
	goto L35
L35:
	;
	F___ashlti3(m, v127, v131, int64(0), v169+int32(49))
	mBase = m.M
	v178 = *(*int64)(unsafe.Add(mBase, uint32(v127+int32(8))))
	v181 = *(*int64)(unsafe.Add(mBase, uint32(v127)))
	v182 = v181
	v183 = base.I64_extend_i32_u(int32(15372) - v169)
	v184 = v178 ^ int64(281474976710656)
	goto L26
L36:
	;
	v205 = *(*float64)(unsafe.Add(mBase, uint32(v12)+56))
	v212 = m.G0
	v214 = v212 - int32(16)
	m.G0 = v214
	v216 = base.I64_reinterpret_f64(v205)
	v218 = v216 & int64(4503599627370495)
	v222 = int64(base.Ui64(v216)>>(uint(int64(52))%64)) & int64(2047)
	if v222 == int64(0) {
		goto L39
	} else {
		goto L40
	}
L37:
	;
	v284 = *(*int64)(unsafe.Add(mBase, uint32(v12)))
	v287 = *(*int64)(unsafe.Add(mBase, uint32(v12+int32(8))))
	F_addReplyHumanLongDouble(m, l0, v284, v287)
	mBase = m.M
	v289 = m.ExcPending
	if v289 != 0 {
		goto L2
	} else {
		goto L48
	}
L38:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v12))) = v269
	*(*int64)(unsafe.Add(mBase, uint32(v12)+8)) = v270<<(uint(int64(48))%64) | v216&int64(-9223372036854775807-1) | v271
	m.G0 = v214 + int32(16)
	goto L37
L39:
	;
	if base.B2i32(v218 == int64(0)) == int32(0) {
		goto L43
	} else {
		goto L44
	}
L40:
	;
	if v222 == int64(2047) {
		goto L41
	} else {
		goto L42
	}
L41:
	;
	v269 = v218 << (uint(int64(60)) % 64)
	v270 = int64(32767)
	v271 = int64(base.Ui64(v218) >> (uint(int64(4)) % 64))
	goto L38
L42:
	;
	v269 = v218 << (uint(int64(60)) % 64)
	v270 = v222 + int64(15360)
	v271 = int64(base.Ui64(v218) >> (uint(int64(4)) % 64))
	goto L38
L43:
	;
	if base.Ui64(v218) < base.Ui64(int64(4294967296)) {
		goto L45
	} else {
		goto L46
	}
L44:
	;
	v242 = int64(0)
	v269 = v242
	v270 = v242
	v271 = v242
	goto L38
L45:
	;
	v256 = base.I32_clz(base.I32_wrap_i64(v216)) | int32(32)
	goto L47
L46:
	;
	v256 = base.I32_clz(base.I32_wrap_i64(int64(base.Ui64(v218) >> (uint(int64(32)) % 64))))
	goto L47
L47:
	;
	F___ashlti3(m, v214, v218, int64(0), v256+int32(49))
	mBase = m.M
	v265 = *(*int64)(unsafe.Add(mBase, uint32(v214+int32(8))))
	v268 = *(*int64)(unsafe.Add(mBase, uint32(v214)))
	v269 = v268
	v270 = base.I64_extend_i32_u(int32(15372) - v256)
	v271 = v265 ^ int64(281474976710656)
	goto L38
L48:
	;
	goto L10
L49:
	;
	goto L9
}
func F_geosearchstoreCommand(m *base.Module, l0 int32) {
	var v5 int32
	_ = v5
	F_georadiusGeneric(m, l0, int32(2), int32(24))
	v5 = m.ExcPending
	if v5 != 0 {
		return
	} else {
		return
	}
}
func F_getConnectionTypeName(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	if base.Ui32(int32(3)) < base.Ui32(l0) {
		v11 = int32(_a265)
	} else {
		v10 = *(*int32)(unsafe.Add(mBase, uint32(l0<<(uint(int32(2))%32))+uint32(_consts[769])))
		v11 = v10
	}
	return v11
}
func F_getMaxmemoryState(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v18 int32
	_ = v18
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	var v106 int64
	_ = v106
	var v109 int32
	_ = v109
	var v112 int32
	_ = v112
	var v116 int64
	_ = v116
	var v121 int32
	_ = v121
	var v122 int64
	_ = v122
	var v124 int32
	_ = v124
	var v129 int64
	_ = v129
	var v136 int32
	_ = v136
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v141 int32
	_ = v141
	var v145 int32
	_ = v145
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v161 int32
	_ = v161
	var v168 int32
	_ = v168
	var v172 int32
	_ = v172
	var v176 int32
	_ = v176
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v185 int32
	_ = v185
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v191 int32
	_ = v191
	var v193 int32
	_ = v193
	var v197 int32
	_ = v197
	var v201 int32
	_ = v201
	var v204 int32
	_ = v204
	var v206 int32
	_ = v206
	var v208 int32
	_ = v208
	var v211 int32
	_ = v211
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v227 int32
	_ = v227
	var v229 int32
	_ = v229
	var v236 int32
	_ = v236
	var v239 int32
	_ = v239
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
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v258 int32
	_ = v258
	var v259 int32
	_ = v259
	var v260 int32
	_ = v260
	var v263 int32
	_ = v263
	var v265 int32
	_ = v265
	var v269 int32
	_ = v269
	var v271 int32
	_ = v271
	var v277 int32
	_ = v277
	var v278 int32
	_ = v278
	var v280 int32
	_ = v280
	var v282 int32
	_ = v282
	var v284 int64
	_ = v284
	var v297 int32
	_ = v297
	var v306 int32
	_ = v306
	v5 = int32(0)
	v18 = *(*int32)(unsafe.Add(mBase, _consts[315]))
	if v18 < int32(261) {
		goto L4
	} else {
		goto L5
	}
L1:
	;
	if l0 == int32(0) {
		goto L17
	} else {
		goto L18
	}
L2:
	;
	goto L1
L3:
	;
	v29 = v27 & int32(3)
	if base.Ui32(int32(4)) <= base.Ui32(v27) {
		goto L8
	} else {
		goto L9
	}
L4:
	;
	if v18 < int32(1) {
		v95 = v5
		goto L2
	} else {
		goto L6
	}
L5:
	;
	v22 = *(*int32)(unsafe.Add(mBase, _consts[316]))
	v26 = v22
	v27 = int32(260)
	goto L3
L6:
	;
	v26 = v5
	v27 = v18
	goto L3
L7:
	;
	if v29 == int32(0) {
		v95 = v68
		goto L2
	} else {
		goto L13
	}
L8:
	;
	v36 = int32(0)
	v38 = v26
	v39 = v36
	v43 = v36
	goto L10
L9:
	;
	v68 = v26
	v69 = int32(0)
	goto L7
L10:
	;
	v46 = v39 << (uint(int32(2)) % 32)
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v46)+uint32(_consts[317])))
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v46)+uint32(_consts[318])))
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v46)+uint32(_consts[319])))
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v46)+uint32(_consts[320])))
	v62 = v49 + (v52 + (v55 + (v58 + v38)))
	v63 = int32(4)
	v64 = v39 + v63
	v66 = v43 + v63
	if v66 != v27&int32(2147483644) {
		v38 = v62
		v39 = v64
		v43 = v66
		goto L10
	} else {
		goto L12
	}
L11:
	;
	v68 = v62
	v69 = v64
	goto L7
L12:
	;
	goto L11
L13:
	;
	v77 = v68
	v78 = v69
	v80 = int32(0)
	goto L14
L14:
	;
	v88 = *(*int32)(unsafe.Add(mBase, uint32(v78<<(uint(int32(2))%32))+uint32(_consts[320])))
	v89 = v88 + v77
	v90 = int32(1)
	v93 = v80 + v90
	if v93 != v29 {
		v77 = v89
		v78 = v78 + v90
		v80 = v93
		goto L14
	} else {
		goto L16
	}
L15:
	;
	v95 = v89
	goto L2
L16:
	;
	goto L15
L17:
	;
	v106 = *(*int64)(unsafe.Add(mBase, _consts[321]))
	if v106 != int64(0) {
		goto L20
	} else {
		goto L21
	}
L18:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v95
	goto L17
L19:
	;
	return v306
L20:
	;
	v116 = base.I64_extend_i32_u(v95)
	if l3 != 0 {
		goto L23
	} else {
		goto L24
	}
L21:
	;
	v109 = int32(0)
	if l3 == v109 {
		v306 = v109
		goto L19
	} else {
		goto L22
	}
L22:
	;
	v112 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v112
	return v112
L23:
	;
	v121 = int32(_a20)
	v122 = *(*int64)(unsafe.Add(mBase, _consts[322]))
	v124 = *(*int32)(unsafe.Add(mBase, _consts[323]))
	if base.I64_extend_i32_u(v124) <= v122 {
		v139 = int32(0)
		goto L26
	} else {
		goto L27
	}
L24:
	;
	if base.Ui64(v106) < base.Ui64(v116) {
		goto L23
	} else {
		goto L25
	}
L25:
	;
	return int32(0)
L26:
	;
	v141 = *(*int32)(unsafe.Add(mBase, _consts[38]))
	if v141 == int32(0) {
		v187 = v139
		goto L31
	} else {
		goto L32
	}
L27:
	;
	v129 = base.I64_div_s(v122, int64(16384))
	v136 = v124 - base.I32_wrap_i64(v122+v129*int64(44)) + int32(-44)
	if base.Ui32(v124) < base.Ui32(v136) {
		goto L28
	} else {
		goto L29
	}
L28:
	;
	v138 = int32(0)
	goto L30
L29:
	;
	v138 = v136
	goto L30
L30:
	;
	v139 = v138
	goto L26
L31:
	;
	v188 = int32(0)
	v191 = m.G0
	v193 = v191 - int32(16)
	m.G0 = v193
	v197 = *(*int32)(unsafe.Add(mBase, _consts[110]))
	if v197 == v188 {
		v229 = v188
		goto L48
	} else {
		goto L49
	}
L32:
	;
	v145 = *(*int32)(unsafe.Add(mBase, _consts[66]))
	v152 = v145 + int32(-1)
	v153 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v152))))
	v155 = v153 & int32(7)
	switch v155 {
	case 0:
		goto L39
	case 1:
		v161 = int32(4)
		goto L34
	case 2:
		goto L38
	case 3:
		goto L37
	case 4:
		goto L36
	default:
		goto L35
	}
L33:
	;
	v187 = v185 + v139
	goto L31
L34:
	;
	switch v155 {
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
		v181 = int32(0)
		goto L40
	}
L35:
	;
	v161 = int32(1)
	goto L34
L36:
	;
	v161 = int32(18)
	goto L34
L37:
	;
	v161 = int32(10)
	goto L34
L38:
	;
	v161 = int32(6)
	goto L34
L39:
	;
	v156 = F_zmalloc_usable_size(m, v152)
	mBase = m.M
	v185 = v156
	goto L33
L40:
	;
	v185 = v161 + v181
	goto L33
L41:
	;
	v180 = *(*int32)(unsafe.Add(mBase, uint32(v145+int32(-9))))
	v181 = v180
	goto L40
L42:
	;
	v176 = *(*int32)(unsafe.Add(mBase, uint32(v145+int32(-5))))
	v185 = v161 + v176
	goto L33
L43:
	;
	v172 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v145+int32(-3)))))
	v185 = v161 + v172
	goto L33
L44:
	;
	v168 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v145+int32(-2)))))
	v185 = v161 + v168
	goto L33
L45:
	;
	v185 = v161 + int32(base.Ui32(v153)>>(uint(int32(3))%32))
	goto L33
L46:
	;
	v278 = int32(0)
	v280 = v95 - v277
	if base.Ui32(v95) < base.Ui32(v280) {
		goto L68
	} else {
		goto L69
	}
L47:
	;
	if v229 == int32(0) {
		v277 = v187
		goto L46
	} else {
		goto L58
	}
L48:
	;
	m.G0 = v193 + int32(16)
	goto L47
L49:
	;
	v201 = *(*int32)(unsafe.Add(mBase, _consts[111]))
	if v201 == int32(0) {
		v229 = v188
		goto L48
	} else {
		goto L50
	}
L50:
	;
	v204 = *(*int32)(unsafe.Add(mBase, uint32(v201)+uint32(_consts[171])))
	v206 = v193 + int32(8)
	F_listRewind(m, v204, v206)
	mBase = m.M
	v208 = int32(0)
	v211 = F_listNext(m, v206)
	mBase = m.M
	if v211 == v208 {
		v229 = v208
		goto L48
	} else {
		goto L51
	}
L51:
	;
	v216 = v211
	goto L52
L52:
	;
	v217 = *(*int32)(unsafe.Add(mBase, uint32(v216)+8))
	v218 = *(*int32)(unsafe.Add(mBase, uint32(v217)))
	if v218 != 0 {
		goto L54
	} else {
		goto L55
	}
L53:
	;
	v229 = v208
	goto L48
L54:
	;
	v227 = F_listNext(m, v193+int32(8))
	mBase = m.M
	if v227 != 0 {
		v216 = v227
		goto L52
	} else {
		goto L57
	}
L55:
	;
	v219 = *(*int32)(unsafe.Add(mBase, uint32(v217)+156))
	if base.Ui32(v219+int32(-18)) <= base.Ui32(int32(2)) {
		goto L54
	} else {
		goto L56
	}
L56:
	;
	v229 = int32(1)
	goto L48
L57:
	;
	goto L53
L58:
	;
	v236 = int32(0)
	v239 = m.G0
	v241 = v239 - int32(16)
	m.G0 = v241
	v244 = *(*int32)(unsafe.Add(mBase, _consts[111]))
	v245 = *(*int32)(unsafe.Add(mBase, uint32(v244)+uint32(_consts[171])))
	v247 = v241 + int32(8)
	F_listRewind(m, v245, v247)
	mBase = m.M
	v252 = F_listNext(m, v247)
	mBase = m.M
	if v252 == v236 {
		v271 = v236
		goto L60
	} else {
		goto L61
	}
L59:
	;
	v277 = v271 + v187
	goto L46
L60:
	;
	m.G0 = v241 + int32(16)
	goto L59
L61:
	;
	v256 = v236
	v257 = v252
	goto L62
L62:
	;
	v258 = *(*int32)(unsafe.Add(mBase, uint32(v257)+8))
	v259 = *(*int32)(unsafe.Add(mBase, uint32(v258)))
	if v259 != 0 {
		v265 = v256
		goto L64
	} else {
		goto L65
	}
L63:
	;
	v271 = v265
	goto L60
L64:
	;
	v269 = F_listNext(m, v241+int32(8))
	mBase = m.M
	if v269 != 0 {
		v256 = v265
		v257 = v269
		goto L62
	} else {
		goto L67
	}
L65:
	;
	v260 = *(*int32)(unsafe.Add(mBase, uint32(v258)+152))
	if v260 == int32(0) {
		v265 = v256
		goto L64
	} else {
		goto L66
	}
L66:
	;
	v263 = F_getClientOutputBufferMemoryUsage(m, v260)
	mBase = m.M
	v265 = v263 + v256
	goto L64
L67:
	;
	goto L63
L68:
	;
	v282 = v278
	goto L70
L69:
	;
	v282 = v280
	goto L70
L70:
	;
	v284 = *(*int64)(unsafe.Add(mBase, _consts[321]))
	if l3 == int32(0) {
		goto L71
	} else {
		goto L72
	}
L71:
	;
	if base.Ui64(v116) <= base.Ui64(v284) {
		v306 = v278
		goto L19
	} else {
		goto L73
	}
L72:
	;
	*(*float32)(unsafe.Add(mBase, uint32(l3))) = base.F32_div(base.F32_convert_i32_u(v282), base.F32_convert_i64_u(v284))
	goto L71
L73:
	;
	if base.Ui64(base.I64_extend_i32_u(v282)) <= base.Ui64(v284) {
		v306 = v278
		goto L19
	} else {
		goto L74
	}
L74:
	;
	if l1 == int32(0) {
		goto L75
	} else {
		goto L76
	}
L75:
	;
	v297 = int32(-1)
	if l2 == int32(0) {
		v306 = v297
		goto L19
	} else {
		goto L77
	}
L76:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v282
	goto L75
L77:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v282 - base.I32_wrap_i64(v284)
	v306 = v297
	goto L19
}
func F_getMemoryDoctorReport(m *base.Module) int32 {
	mBase := m.M
	_ = mBase
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v26 float32
	_ = v26
	var v28 int32
	_ = v28
	var v30 float32
	_ = v30
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 float32
	_ = v44
	var v50 int32
	_ = v50
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v60 float32
	_ = v60
	var v65 int32
	_ = v65
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v74 float32
	_ = v74
	var v80 int32
	_ = v80
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v97 int32
	_ = v97
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v104 int32
	_ = v104
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v174 int32
	_ = v174
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
	var v189 int32
	_ = v189
	var v191 int32
	_ = v191
	var v193 int32
	_ = v193
	v12 = m.G0
	v14 = v12 - int32(16)
	m.G0 = v14
	v16 = F_getMemoryOverheadData(m)
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		goto L4
	} else {
		goto L5
	}
L1:
	;
	v189 = *(*int32)(unsafe.Add(mBase, uint32(v16)+124))
	F_valkey_free(m, v189)
	mBase = m.M
	v191 = m.ExcPending
	if v191 != 0 {
		goto L4
	} else {
		goto L56
	}
L2:
	;
	v130 = F_sdsnew(m, int32(_a1764))
	mBase = m.M
	v131 = m.ExcPending
	if v131 != 0 {
		goto L4
	} else {
		goto L30
	}
L3:
	;
	v127 = F_sdsnew(m, int32(_a1763))
	mBase = m.M
	v128 = m.ExcPending
	if v128 != 0 {
		goto L4
	} else {
		goto L29
	}
L4:
	;
	return int32(0)
L5:
	;
	v20 = *(*int32)(unsafe.Add(mBase, uint32(v16)+4))
	if base.Ui32(v20) < base.Ui32(int32(5242880)) {
		goto L3
	} else {
		goto L6
	}
L6:
	;
	v23 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
	v26 = base.F32_div(base.F32_convert_i32_u(v23), base.F32_convert_i32_u(v20))
	v28 = base.F32_gt(v26, float32(1.5))
	v30 = *(*float32)(unsafe.Add(mBase, uint32(v16)+76))
	if base.F64_gt(base.F64_promote_f32(v30), float64(1.4)) != 0 {
		goto L8
	} else {
		goto L9
	}
L7:
	;
	v44 = *(*float32)(unsafe.Add(mBase, uint32(v16)+84))
	if base.F64_gt(base.F64_promote_f32(v44), float64(1.1)) == int32(0) {
		v56 = v42
		v57 = int32(1)
		goto L16
	} else {
		goto L17
	}
L8:
	;
	if v28 != 0 {
		goto L10
	} else {
		goto L11
	}
L9:
	;
	v42 = v28
	v43 = int32(1)
	goto L7
L10:
	;
	v37 = int32(2)
	goto L12
L11:
	;
	v37 = int32(1)
	goto L12
L12:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v16)+80))
	v40 = base.B2i32(v38 < int32(10485761))
	if v38 < int32(10485761) {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	v41 = v28
	goto L15
L14:
	;
	v41 = v37
	goto L15
L15:
	;
	v42 = v41
	v43 = v40
	goto L7
L16:
	;
	v60 = *(*float32)(unsafe.Add(mBase, uint32(v16)+92))
	if base.F64_gt(base.F64_promote_f32(v60), float64(1.1)) != 0 {
		goto L19
	} else {
		goto L20
	}
L17:
	;
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v16)+88))
	v56 = v42 + base.B2i32(int32(10485760) < v50)
	v57 = base.B2i32(v50 < int32(10485761))
	goto L16
L18:
	;
	v74 = *(*float32)(unsafe.Add(mBase, uint32(v16)+100))
	if base.F64_gt(base.F64_promote_f32(v74), float64(1.1)) == int32(0) {
		v86 = v71
		v88 = int32(1)
		goto L21
	} else {
		goto L22
	}
L19:
	;
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v16)+96))
	v71 = v56 + base.B2i32(int32(10485760) < v65)
	v73 = base.B2i32(v65 < int32(10485761))
	goto L18
L20:
	;
	v71 = v56
	v73 = int32(1)
	goto L18
L21:
	;
	v89 = *(*int32)(unsafe.Add(mBase, uint32(v16)+24))
	v90 = int32(_a20)
	v91 = *(*int32)(unsafe.Add(mBase, _consts[26]))
	v92 = *(*int32)(unsafe.Add(mBase, uint32(v91)+20))
	v94 = *(*int32)(unsafe.Add(mBase, _consts[78]))
	v95 = *(*int32)(unsafe.Add(mBase, uint32(v94)+20))
	v97 = base.I32_div_u_s(v89, v92-v95)
	v100 = v86 + base.B2i32(base.Ui32(int32(204800)) < base.Ui32(v97))
	v101 = int32(1)
	if v95 < v101 {
		v110 = v100
		v111 = v101
		goto L23
	} else {
		goto L24
	}
L22:
	;
	v80 = *(*int32)(unsafe.Add(mBase, uint32(v16)+104))
	v86 = v71 + base.B2i32(base.Ui32(int32(10485760)) < base.Ui32(v80))
	v88 = base.B2i32(base.Ui32(v80) < base.Ui32(int32(10485761)))
	goto L21
L23:
	;
	v114 = *(*int32)(unsafe.Add(mBase, _consts[311]))
	goto L25
L24:
	;
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v16)+20))
	v110 = v100 + base.B2i32(base.Ui32(int32(10485760)) < base.Ui32(v104))
	v111 = base.B2i32(base.Ui32(v104) < base.Ui32(int32(10485761)))
	goto L23
L25:
	;
	v115 = *(*int32)(unsafe.Add(mBase, uint32(v114)+12))
	v117 = *(*int32)(unsafe.Add(mBase, _consts[311]))
	goto L26
L26:
	;
	v118 = *(*int32)(unsafe.Add(mBase, uint32(v117)+16))
	v119 = v115 + v118
	if v110|base.B2i32(base.Ui32(int32(1000)) < base.Ui32(v119)) != 0 {
		goto L2
	} else {
		goto L27
	}
L27:
	;
	v124 = F_sdsnew(m, int32(_a1772))
	mBase = m.M
	v125 = m.ExcPending
	if v125 != 0 {
		goto L4
	} else {
		goto L28
	}
L28:
	;
	v180 = v124
	goto L1
L29:
	;
	v180 = v127
	goto L1
L30:
	;
	if base.F32_gt(v26, float32(1.5)) == int32(0) {
		v139 = v130
		goto L31
	} else {
		goto L32
	}
L31:
	;
	if v43 != 0 {
		v145 = v139
		goto L34
	} else {
		goto L35
	}
L32:
	;
	v137 = F_sdscat(m, v130, int32(_a1773))
	mBase = m.M
	v138 = m.ExcPending
	if v138 != 0 {
		goto L4
	} else {
		goto L33
	}
L33:
	;
	v139 = v137
	goto L31
L34:
	;
	if v57 != 0 {
		v150 = v145
		goto L37
	} else {
		goto L38
	}
L35:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14))) = int32(_a1774)
	v143 = F_sdscatprintf(m, v139, int32(_a1775), v14)
	mBase = m.M
	v144 = m.ExcPending
	if v144 != 0 {
		goto L4
	} else {
		goto L36
	}
L36:
	;
	v145 = v143
	goto L34
L37:
	;
	if v73 != 0 {
		v155 = v150
		goto L40
	} else {
		goto L41
	}
L38:
	;
	v148 = F_sdscatprintf(m, v145, int32(_a1771), int32(0))
	mBase = m.M
	v149 = m.ExcPending
	if v149 != 0 {
		goto L4
	} else {
		goto L39
	}
L39:
	;
	v150 = v148
	goto L37
L40:
	;
	if v88 != 0 {
		v160 = v155
		goto L43
	} else {
		goto L44
	}
L41:
	;
	v153 = F_sdscatprintf(m, v150, int32(_a1770), int32(0))
	mBase = m.M
	v154 = m.ExcPending
	if v154 != 0 {
		goto L4
	} else {
		goto L42
	}
L42:
	;
	v155 = v153
	goto L40
L43:
	;
	if v111 != 0 {
		v164 = v160
		goto L46
	} else {
		goto L47
	}
L44:
	;
	v158 = F_sdscatprintf(m, v155, int32(_a1769), int32(0))
	mBase = m.M
	v159 = m.ExcPending
	if v159 != 0 {
		goto L4
	} else {
		goto L45
	}
L45:
	;
	v160 = v158
	goto L43
L46:
	;
	if base.Ui32(v97) < base.Ui32(int32(204801)) {
		v170 = v164
		goto L49
	} else {
		goto L50
	}
L47:
	;
	v162 = F_sdscat(m, v160, int32(_a1768))
	mBase = m.M
	v163 = m.ExcPending
	if v163 != 0 {
		goto L4
	} else {
		goto L48
	}
L48:
	;
	v164 = v162
	goto L46
L49:
	;
	if base.Ui32(v119) < base.Ui32(int32(1001)) {
		v176 = v170
		goto L52
	} else {
		goto L53
	}
L50:
	;
	v168 = F_sdscat(m, v164, int32(_a1767))
	mBase = m.M
	v169 = m.ExcPending
	if v169 != 0 {
		goto L4
	} else {
		goto L51
	}
L51:
	;
	v170 = v168
	goto L49
L52:
	;
	v178 = F_sdscat(m, v176, int32(_a1765))
	mBase = m.M
	v179 = m.ExcPending
	if v179 != 0 {
		goto L4
	} else {
		goto L55
	}
L53:
	;
	v174 = F_sdscat(m, v170, int32(_a1766))
	mBase = m.M
	v175 = m.ExcPending
	if v175 != 0 {
		goto L4
	} else {
		goto L54
	}
L54:
	;
	v176 = v174
	goto L52
L55:
	;
	v180 = v178
	goto L1
L56:
	;
	F_valkey_free(m, v16)
	mBase = m.M
	v193 = m.ExcPending
	if v193 != 0 {
		goto L4
	} else {
		goto L57
	}
L57:
	;
	m.G0 = v14 + int32(16)
	return v180
}
func F_getMyClusterNode(m *base.Module) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	v2 = *(*int32)(unsafe.Add(mBase, _consts[111]))
	v3 = *(*int32)(unsafe.Add(mBase, uint32(v2)))
	return v3
}
func F_getNewBaseFileNameAndMarkPreAsHistory(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
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
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v34 int64
	_ = v34
	var v36 int64
	_ = v36
	var v42 int32
	_ = v42
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v54 int64
	_ = v54
	var v69 int32
	_ = v69
	var v75 int32
	_ = v75
	v8 = m.G0
	v10 = v8 - int32(32)
	m.G0 = v10
	if l0 == int32(0) {
		F__serverAssert(m, int32(_a124), int32(_a123), int32(428))
		mBase = m.M
		v69 = m.ExcPending
		if v69 != 0 {
			return int32(0)
		} else {
			F_abort(m)
			mBase = m.M
			base.Wasm_trap_unreachable()
			for {
			}
		}
	} else {
		v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		if v14 == int32(0) {
			v28 = F_valkey_calloc(m, int32(24))
			mBase = m.M
			v29 = m.ExcPending
			if v29 != 0 {
				return int32(0)
			} else {
				v30 = F_sdsempty(m)
				mBase = m.M
				v31 = m.ExcPending
				if v31 != 0 {
					return int32(0)
				} else {
					v33 = *(*int32)(unsafe.Add(mBase, _consts[36]))
					v34 = *(*int64)(unsafe.Add(mBase, uint32(l0)+16))
					v36 = v34 + int64(1)
					*(*int64)(unsafe.Add(mBase, uint32(l0)+16)) = v36
					if l1 != 0 {
						v42 = int32(_a128)
					} else {
						v42 = int32(_a129)
					}
					*(*int32)(unsafe.Add(mBase, uint32(v10+int32(20)))) = v42
					*(*int32)(unsafe.Add(mBase, uint32(v10+int32(16)))) = int32(_a130)
					*(*int64)(unsafe.Add(mBase, uint32(v10)+8)) = v36
					*(*int32)(unsafe.Add(mBase, uint32(v10))) = v33
					v51 = F_sdscatprintf(m, v30, int32(_a131), v10)
					mBase = m.M
					v52 = m.ExcPending
					if v52 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v28))) = v51
						v54 = *(*int64)(unsafe.Add(mBase, uint32(l0)+16))
						*(*int32)(unsafe.Add(mBase, uint32(v28)+16)) = int32(98)
						*(*int64)(unsafe.Add(mBase, uint32(v28)+8)) = v54
						*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = int32(1)
						*(*int32)(unsafe.Add(mBase, uint32(l0))) = v28
						m.G0 = v10 + int32(32)
						return v51
					}
				}
			}
		} else {
			v17 = *(*int32)(unsafe.Add(mBase, uint32(v14)+16))
			if v17 != int32(98) {
				F__serverAssert(m, int32(_a132), int32(_a123), int32(430))
				mBase = m.M
				v75 = m.ExcPending
				if v75 != 0 {
					return int32(0)
				} else {
					F_abort(m)
					mBase = m.M
					base.Wasm_trap_unreachable()
					for {
					}
				}
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v14)+16)) = int32(104)
				v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				v23 = F_listAddNodeHead(m, v22, v14)
				mBase = m.M
				v26 = m.ExcPending
				if v26 != 0 {
					return int32(0)
				} else {
					v28 = F_valkey_calloc(m, int32(24))
					mBase = m.M
					v29 = m.ExcPending
					if v29 != 0 {
						return int32(0)
					} else {
						v30 = F_sdsempty(m)
						mBase = m.M
						v31 = m.ExcPending
						if v31 != 0 {
							return int32(0)
						} else {
							v33 = *(*int32)(unsafe.Add(mBase, _consts[36]))
							v34 = *(*int64)(unsafe.Add(mBase, uint32(l0)+16))
							v36 = v34 + int64(1)
							*(*int64)(unsafe.Add(mBase, uint32(l0)+16)) = v36
							if l1 != 0 {
								v42 = int32(_a128)
							} else {
								v42 = int32(_a129)
							}
							*(*int32)(unsafe.Add(mBase, uint32(v10+int32(20)))) = v42
							*(*int32)(unsafe.Add(mBase, uint32(v10+int32(16)))) = int32(_a130)
							*(*int64)(unsafe.Add(mBase, uint32(v10)+8)) = v36
							*(*int32)(unsafe.Add(mBase, uint32(v10))) = v33
							v51 = F_sdscatprintf(m, v30, int32(_a131), v10)
							mBase = m.M
							v52 = m.ExcPending
							if v52 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v28))) = v51
								v54 = *(*int64)(unsafe.Add(mBase, uint32(l0)+16))
								*(*int32)(unsafe.Add(mBase, uint32(v28)+16)) = int32(98)
								*(*int64)(unsafe.Add(mBase, uint32(v28)+8)) = v54
								*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = int32(1)
								*(*int32)(unsafe.Add(mBase, uint32(l0))) = v28
								m.G0 = v10 + int32(32)
								return v51
							}
						}
					}
				}
			}
		}
	}
}
func F_getPausedActionTimeout(m *base.Module, l0 int32, l1 int32) int64 {
	mBase := m.M
	_ = mBase
	var v3 int64
	_ = v3
	var v8 int32
	_ = v8
	var v9 int64
	_ = v9
	var v12 int32
	_ = v12
	var v17 int64
	_ = v17
	var v18 int64
	_ = v18
	var v23 int64
	_ = v23
	var v26 int32
	_ = v26
	var v31 int64
	_ = v31
	var v32 int64
	_ = v32
	var v36 int64
	_ = v36
	var v39 int32
	_ = v39
	var v44 int64
	_ = v44
	var v45 int64
	_ = v45
	var v49 int64
	_ = v49
	var v52 int32
	_ = v52
	var v57 int64
	_ = v57
	var v58 int64
	_ = v58
	var v63 int64
	_ = v63
	v3 = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = int32(4)
	v8 = int32(_a20)
	v9 = *(*int64)(unsafe.Add(mBase, _consts[54]))
	v12 = *(*int32)(unsafe.Add(mBase, _consts[499]))
	if v12&l0 == int32(0) {
		v23 = v3
	} else {
		v17 = *(*int64)(unsafe.Add(mBase, _consts[497]))
		v18 = v17 - v9
		if v18 < int64(1) {
			v23 = v3
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(l1))) = int32(0)
			v23 = v18
		}
	}
	v26 = *(*int32)(unsafe.Add(mBase, _consts[501]))
	if v26&l0 == int32(0) {
		v36 = v23
	} else {
		v31 = *(*int64)(unsafe.Add(mBase, _consts[498]))
		v32 = v31 - v9
		if v32 <= v23 {
			v36 = v23
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(l1))) = int32(1)
			v36 = v32
		}
	}
	v39 = *(*int32)(unsafe.Add(mBase, _consts[503]))
	if v39&l0 == int32(0) {
		v49 = v36
	} else {
		v44 = *(*int64)(unsafe.Add(mBase, _consts[500]))
		v45 = v44 - v9
		if v45 <= v36 {
			v49 = v36
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(l1))) = int32(2)
			v49 = v45
		}
	}
	v52 = *(*int32)(unsafe.Add(mBase, _consts[228]))
	if v52&l0 == int32(0) {
		v63 = v49
	} else {
		v57 = *(*int64)(unsafe.Add(mBase, _consts[502]))
		v58 = v57 - v9
		if v58 <= v49 {
			v63 = v49
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(l1))) = int32(3)
			v63 = v58
		}
	}
	return v63
}
func F_getPausedActionsWithPurpose(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0<<(uint(int32(4))%32))+uint32(_consts[499])))
	return v6
}
func F_getRandomBytes(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v49 int32
	_ = v49
	var v50 int64
	_ = v50
	var v53 int64
	_ = v53
	var v56 int64
	_ = v56
	var v59 int64
	_ = v59
	var v62 int64
	_ = v62
	var v65 int64
	_ = v65
	var v68 int64
	_ = v68
	var v71 int64
	_ = v71
	var v83 int32
	_ = v83
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v97 int32
	_ = v97
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v109 int32
	_ = v109
	var v112 int32
	_ = v112
	var v126 int32
	_ = v126
	var v128 int64
	_ = v128
	var v131 int64
	_ = v131
	var v134 int64
	_ = v134
	var v137 int64
	_ = v137
	var v140 int64
	_ = v140
	var v143 int64
	_ = v143
	var v146 int64
	_ = v146
	var v149 int64
	_ = v149
	var v160 int32
	_ = v160
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v174 int32
	_ = v174
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v180 int32
	_ = v180
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v186 int32
	_ = v186
	var v189 int32
	_ = v189
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v204 int32
	_ = v204
	var v206 int64
	_ = v206
	var v215 int32
	_ = v215
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v221 int32
	_ = v221
	v13 = m.G0
	v15 = v13 - int32(208)
	m.G0 = v15
	v18 = int32(*(*uint8)(unsafe.Add(mBase, _consts[890])))
	if v18 != 0 {
		if l1 == int32(0) {
		} else {
			v24 = v15 + int32(168)
			v26 = v15 + int32(160)
			v28 = v15 + int32(152)
			v32 = v15 + int32(144)
			v34 = v15 + int32(136)
			v36 = v15 + int32(128)
			v37 = l0
			v38 = l1
			for {
				v49 = int32(0)
				v50 = *(*int64)(unsafe.Add(mBase, _consts[891]))
				*(*int64)(unsafe.Add(mBase, uint32(v24))) = v50
				v53 = *(*int64)(unsafe.Add(mBase, _consts[892]))
				*(*int64)(unsafe.Add(mBase, uint32(v26))) = v53
				v56 = *(*int64)(unsafe.Add(mBase, _consts[893]))
				*(*int64)(unsafe.Add(mBase, uint32(v28))) = v56
				v59 = *(*int64)(unsafe.Add(mBase, _consts[894]))
				*(*int64)(unsafe.Add(mBase, uint32(v32))) = v59
				v62 = *(*int64)(unsafe.Add(mBase, _consts[895]))
				*(*int64)(unsafe.Add(mBase, uint32(v34))) = v62
				v65 = *(*int64)(unsafe.Add(mBase, _consts[896]))
				*(*int64)(unsafe.Add(mBase, uint32(v36))) = v65
				v68 = *(*int64)(unsafe.Add(mBase, _consts[897]))
				*(*int64)(unsafe.Add(mBase, uint32(v15)+120)) = v68
				v71 = *(*int64)(unsafe.Add(mBase, _consts[898]))
				*(*int64)(unsafe.Add(mBase, uint32(v15)+112)) = v71
				v83 = v49
				for {
					v88 = v15 + int32(112) + v83
					v89 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v88))))
					v90 = int32(54)
					v91 = v89 ^ v90
					*(*uint8)(unsafe.Add(mBase, uint32(v88))) = uint8(v91)
					v94 = v88 + int32(1)
					v95 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v94))))
					v97 = v95 ^ v90
					*(*uint8)(unsafe.Add(mBase, uint32(v94))) = uint8(v97)
					v100 = v88 + int32(2)
					v101 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v100))))
					v103 = v101 ^ v90
					*(*uint8)(unsafe.Add(mBase, uint32(v100))) = uint8(v103)
					v106 = v88 + int32(3)
					v107 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v106))))
					v109 = v107 ^ v90
					*(*uint8)(unsafe.Add(mBase, uint32(v106))) = uint8(v109)
					v112 = v83 + int32(4)
					if v112 != int32(64) {
						v83 = v112
						continue
					} else {
						break
					}
					break
				}
				F_sha256_init(m, v15)
				mBase = m.M
				F_sha256_update(m, v15, v15+int32(112), int32(64))
				mBase = m.M
				F_sha256_update(m, v15, int32(_a2468), int32(8))
				mBase = m.M
				F_sha256_final(m, v15, v15+int32(176))
				mBase = m.M
				v126 = int32(0)
				v128 = *(*int64)(unsafe.Add(mBase, _consts[891]))
				*(*int64)(unsafe.Add(mBase, uint32(v24))) = v128
				v131 = *(*int64)(unsafe.Add(mBase, _consts[892]))
				*(*int64)(unsafe.Add(mBase, uint32(v26))) = v131
				v134 = *(*int64)(unsafe.Add(mBase, _consts[893]))
				*(*int64)(unsafe.Add(mBase, uint32(v28))) = v134
				v137 = *(*int64)(unsafe.Add(mBase, _consts[894]))
				*(*int64)(unsafe.Add(mBase, uint32(v32))) = v137
				v140 = *(*int64)(unsafe.Add(mBase, _consts[895]))
				*(*int64)(unsafe.Add(mBase, uint32(v34))) = v140
				v143 = *(*int64)(unsafe.Add(mBase, _consts[896]))
				*(*int64)(unsafe.Add(mBase, uint32(v36))) = v143
				v146 = *(*int64)(unsafe.Add(mBase, _consts[897]))
				*(*int64)(unsafe.Add(mBase, uint32(v15)+120)) = v146
				v149 = *(*int64)(unsafe.Add(mBase, _consts[898]))
				*(*int64)(unsafe.Add(mBase, uint32(v15)+112)) = v149
				v160 = v126
				for {
					v165 = v15 + int32(112) + v160
					v166 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v165))))
					v167 = int32(92)
					v168 = v166 ^ v167
					*(*uint8)(unsafe.Add(mBase, uint32(v165))) = uint8(v168)
					v171 = v165 + int32(1)
					v172 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v171))))
					v174 = v172 ^ v167
					*(*uint8)(unsafe.Add(mBase, uint32(v171))) = uint8(v174)
					v177 = v165 + int32(2)
					v178 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v177))))
					v180 = v178 ^ v167
					*(*uint8)(unsafe.Add(mBase, uint32(v177))) = uint8(v180)
					v183 = v165 + int32(3)
					v184 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v183))))
					v186 = v184 ^ v167
					*(*uint8)(unsafe.Add(mBase, uint32(v183))) = uint8(v186)
					v189 = v160 + int32(4)
					if v189 != int32(64) {
						v160 = v189
						continue
					} else {
						break
					}
					break
				}
				F_sha256_init(m, v15)
				mBase = m.M
				F_sha256_update(m, v15, v15+int32(112), int32(64))
				mBase = m.M
				v198 = v15 + int32(176)
				v199 = int32(32)
				F_sha256_update(m, v15, v198, v199)
				mBase = m.M
				F_sha256_final(m, v15, v198)
				mBase = m.M
				v204 = int32(0)
				v206 = *(*int64)(unsafe.Add(mBase, _consts[899]))
				*(*int64)(unsafe.Add(mBase, _consts[899])) = v206 + int64(1)
				if base.Ui32(v38) < base.Ui32(v199) {
					v215 = v38
				} else {
					v215 = v199
				}
				if v215 == int32(0) {
					v219 = v37
				} else {
					v218 = F__emscripten_memcpy_bulkmem(m, v37, v198, v215)
					mBase = m.M
					v219 = v218
				}
				v221 = v38 - v215
				if v221 != 0 {
					v37 = v219 + v215
					v38 = v221
					continue
				} else {
					break
				}
				break
			}
		}
		m.G0 = v15 + int32(208)
		return
	} else {
		F_initializeRandomSeed(m)
		mBase = m.M
		v20 = m.ExcPending
		if v20 != 0 {
			return
		} else {
			if l1 == int32(0) {
			} else {
				v24 = v15 + int32(168)
				v26 = v15 + int32(160)
				v28 = v15 + int32(152)
				v32 = v15 + int32(144)
				v34 = v15 + int32(136)
				v36 = v15 + int32(128)
				v37 = l0
				v38 = l1
				for {
					v49 = int32(0)
					v50 = *(*int64)(unsafe.Add(mBase, _consts[891]))
					*(*int64)(unsafe.Add(mBase, uint32(v24))) = v50
					v53 = *(*int64)(unsafe.Add(mBase, _consts[892]))
					*(*int64)(unsafe.Add(mBase, uint32(v26))) = v53
					v56 = *(*int64)(unsafe.Add(mBase, _consts[893]))
					*(*int64)(unsafe.Add(mBase, uint32(v28))) = v56
					v59 = *(*int64)(unsafe.Add(mBase, _consts[894]))
					*(*int64)(unsafe.Add(mBase, uint32(v32))) = v59
					v62 = *(*int64)(unsafe.Add(mBase, _consts[895]))
					*(*int64)(unsafe.Add(mBase, uint32(v34))) = v62
					v65 = *(*int64)(unsafe.Add(mBase, _consts[896]))
					*(*int64)(unsafe.Add(mBase, uint32(v36))) = v65
					v68 = *(*int64)(unsafe.Add(mBase, _consts[897]))
					*(*int64)(unsafe.Add(mBase, uint32(v15)+120)) = v68
					v71 = *(*int64)(unsafe.Add(mBase, _consts[898]))
					*(*int64)(unsafe.Add(mBase, uint32(v15)+112)) = v71
					v83 = v49
					for {
						v88 = v15 + int32(112) + v83
						v89 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v88))))
						v90 = int32(54)
						v91 = v89 ^ v90
						*(*uint8)(unsafe.Add(mBase, uint32(v88))) = uint8(v91)
						v94 = v88 + int32(1)
						v95 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v94))))
						v97 = v95 ^ v90
						*(*uint8)(unsafe.Add(mBase, uint32(v94))) = uint8(v97)
						v100 = v88 + int32(2)
						v101 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v100))))
						v103 = v101 ^ v90
						*(*uint8)(unsafe.Add(mBase, uint32(v100))) = uint8(v103)
						v106 = v88 + int32(3)
						v107 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v106))))
						v109 = v107 ^ v90
						*(*uint8)(unsafe.Add(mBase, uint32(v106))) = uint8(v109)
						v112 = v83 + int32(4)
						if v112 != int32(64) {
							v83 = v112
							continue
						} else {
							break
						}
						break
					}
					F_sha256_init(m, v15)
					mBase = m.M
					F_sha256_update(m, v15, v15+int32(112), int32(64))
					mBase = m.M
					F_sha256_update(m, v15, int32(_a2468), int32(8))
					mBase = m.M
					F_sha256_final(m, v15, v15+int32(176))
					mBase = m.M
					v126 = int32(0)
					v128 = *(*int64)(unsafe.Add(mBase, _consts[891]))
					*(*int64)(unsafe.Add(mBase, uint32(v24))) = v128
					v131 = *(*int64)(unsafe.Add(mBase, _consts[892]))
					*(*int64)(unsafe.Add(mBase, uint32(v26))) = v131
					v134 = *(*int64)(unsafe.Add(mBase, _consts[893]))
					*(*int64)(unsafe.Add(mBase, uint32(v28))) = v134
					v137 = *(*int64)(unsafe.Add(mBase, _consts[894]))
					*(*int64)(unsafe.Add(mBase, uint32(v32))) = v137
					v140 = *(*int64)(unsafe.Add(mBase, _consts[895]))
					*(*int64)(unsafe.Add(mBase, uint32(v34))) = v140
					v143 = *(*int64)(unsafe.Add(mBase, _consts[896]))
					*(*int64)(unsafe.Add(mBase, uint32(v36))) = v143
					v146 = *(*int64)(unsafe.Add(mBase, _consts[897]))
					*(*int64)(unsafe.Add(mBase, uint32(v15)+120)) = v146
					v149 = *(*int64)(unsafe.Add(mBase, _consts[898]))
					*(*int64)(unsafe.Add(mBase, uint32(v15)+112)) = v149
					v160 = v126
					for {
						v165 = v15 + int32(112) + v160
						v166 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v165))))
						v167 = int32(92)
						v168 = v166 ^ v167
						*(*uint8)(unsafe.Add(mBase, uint32(v165))) = uint8(v168)
						v171 = v165 + int32(1)
						v172 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v171))))
						v174 = v172 ^ v167
						*(*uint8)(unsafe.Add(mBase, uint32(v171))) = uint8(v174)
						v177 = v165 + int32(2)
						v178 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v177))))
						v180 = v178 ^ v167
						*(*uint8)(unsafe.Add(mBase, uint32(v177))) = uint8(v180)
						v183 = v165 + int32(3)
						v184 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v183))))
						v186 = v184 ^ v167
						*(*uint8)(unsafe.Add(mBase, uint32(v183))) = uint8(v186)
						v189 = v160 + int32(4)
						if v189 != int32(64) {
							v160 = v189
							continue
						} else {
							break
						}
						break
					}
					F_sha256_init(m, v15)
					mBase = m.M
					F_sha256_update(m, v15, v15+int32(112), int32(64))
					mBase = m.M
					v198 = v15 + int32(176)
					v199 = int32(32)
					F_sha256_update(m, v15, v198, v199)
					mBase = m.M
					F_sha256_final(m, v15, v198)
					mBase = m.M
					v204 = int32(0)
					v206 = *(*int64)(unsafe.Add(mBase, _consts[899]))
					*(*int64)(unsafe.Add(mBase, _consts[899])) = v206 + int64(1)
					if base.Ui32(v38) < base.Ui32(v199) {
						v215 = v38
					} else {
						v215 = v199
					}
					if v215 == int32(0) {
						v219 = v37
					} else {
						v218 = F__emscripten_memcpy_bulkmem(m, v37, v198, v215)
						mBase = m.M
						v219 = v218
					}
					v221 = v38 - v215
					if v221 != 0 {
						v37 = v219 + v215
						v38 = v221
						continue
					} else {
						break
					}
					break
				}
			}
			m.G0 = v15 + int32(208)
			return
		}
	}
}
func F_getS(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v11 int32
	_ = v11
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v5 != 0 {
		*(*int32)(unsafe.Add(mBase, uint32(l2))) = v5
		*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = int32(0)
		v11 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
		return v11
	} else {
		return int32(0)
	}
}
func F_getSafeInfoString(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v20 int32
	_ = v20
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v36 int32
	_ = v36
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v54 int32
	_ = v54
	var v71 int32
	_ = v71
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v104 int32
	_ = v104
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v121 int32
	_ = v121
	var v125 int32
	_ = v125
	var v128 int32
	_ = v128
	var v131 int32
	_ = v131
	var v143 int32
	_ = v143
	var v154 int32
	_ = v154
	v4 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v4
	if l1 == v4 {
		goto L4
	} else {
		goto L5
	}
L1:
	;
	return v154
L2:
	;
	if v71 == int32(0) {
		v154 = l0
		goto L1
	} else {
		goto L15
	}
L3:
	;
	goto L2
L4:
	;
	v71 = int32(0)
	goto L3
L5:
	;
	v20 = int32(0)
	goto L6
L6:
	;
	goto L9
L7:
	;
	goto L4
L8:
	;
	v54 = v20 + int32(1)
	if v54 != l1 {
		v20 = v54
		goto L6
	} else {
		goto L14
	}
L9:
	;
	v27 = l0 + v20
	v28 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v27))))
	v36 = int32(0)
	goto L10
L10:
	;
	v40 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v36)+uint32(_consts[789]))))
	if v28&int32(255) == v40 {
		v71 = v27
		goto L3
	} else {
		goto L12
	}
L11:
	;
	goto L8
L12:
	;
	v43 = v36 + int32(1)
	if v43 != int32(4) {
		v36 = v43
		goto L10
	} else {
		goto L13
	}
L13:
	;
	goto L11
L14:
	;
	goto L7
L15:
	;
	v77 = F_valkey_malloc(m, l1+int32(1))
	mBase = m.M
	v80 = m.ExcPending
	if v80 != 0 {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	return int32(0)
L17:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v77
	if l1 == int32(0) {
		v85 = v77
		goto L19
	} else {
		goto L20
	}
L18:
	;
	v87 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v85+l1))) = uint8(v87)
	if l1 == v87 {
		goto L22
	} else {
		goto L23
	}
L19:
	;
	goto L18
L20:
	;
	v84 = F__emscripten_memcpy_bulkmem(m, v77, l0, l1)
	mBase = m.M
	v85 = v84
	goto L19
L21:
	;
	v154 = v85
	goto L1
L22:
	;
	goto L21
L23:
	;
	v104 = int32(0)
	goto L24
L24:
	;
	goto L27
L25:
	;
	goto L22
L26:
	;
	v143 = v104 + int32(1)
	if v143 != l1 {
		v104 = v143
		goto L24
	} else {
		goto L33
	}
L27:
	;
	v111 = v85 + v104
	v112 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v111))))
	v121 = int32(0)
	goto L28
L28:
	;
	v125 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v121)+uint32(_consts[789]))))
	if v112&int32(255) != v125 {
		goto L30
	} else {
		goto L31
	}
L29:
	;
	goto L26
L30:
	;
	v131 = v121 + int32(1)
	if v131 != int32(4) {
		v121 = v131
		goto L28
	} else {
		goto L32
	}
L31:
	;
	v128 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v121)+uint32(_consts[790]))))
	*(*uint8)(unsafe.Add(mBase, uint32(v111))) = uint8(v128)
	goto L26
L32:
	;
	goto L29
L33:
	;
	goto L25
}
func F_getTimeZone(m *base.Module) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	v2 = *(*int32)(unsafe.Add(mBase, _consts[900]))
	return v2
}
func F_getUpcomingChannelList(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v34 int32
	_ = v34
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v67 int32
	_ = v67
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v88 int32
	_ = v88
	var v91 int32
	_ = v91
	var v97 int32
	_ = v97
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	var v110 int32
	_ = v110
	var v116 int32
	_ = v116
	var v123 int32
	_ = v123
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v131 int32
	_ = v131
	var v133 int32
	_ = v133
	var v136 int32
	_ = v136
	var v142 int32
	_ = v142
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
	var v166 int32
	_ = v166
	var v169 int32
	_ = v169
	var v175 int32
	_ = v175
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v183 int32
	_ = v183
	var v185 int32
	_ = v185
	var v188 int32
	_ = v188
	var v194 int32
	_ = v194
	var v197 int32
	_ = v197
	var v199 int32
	_ = v199
	var v202 int32
	_ = v202
	var v208 int32
	_ = v208
	var v216 int32
	_ = v216
	var v227 int32
	_ = v227
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v12 = v8 + int32(8)
	v13 = *(*int32)(unsafe.Add(mBase, uint32(v10)))
	*(*int32)(unsafe.Add(mBase, uint32(v12)+4)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v12))) = v13
	goto L1
L1:
	;
	goto L5
L2:
	;
	m.G0 = v8 + int32(16)
	return v227
L3:
	;
	v227 = int32(0)
	goto L2
L4:
	;
	v44 = F_listCreate(m)
	mBase = m.M
	v47 = m.ExcPending
	if v47 != 0 {
		goto L12
	} else {
		goto L13
	}
L5:
	;
	v23 = v8 + int32(8)
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v23)))
	if v25 == int32(0) {
		goto L8
	} else {
		goto L9
	}
L7:
	;
	if v25 == int32(0) {
		goto L4
	} else {
		goto L10
	}
L8:
	;
	goto L7
L9:
	;
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v23)+4))
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v25+base.B2i32(v28 == int32(0))<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v23))) = v34
	goto L8
L10:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v25)+8))
	v39 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v38))))
	if v39&int32(8) == int32(0) {
		goto L5
	} else {
		goto L11
	}
L11:
	;
	goto L3
L12:
	;
	return int32(0)
L13:
	;
	v48 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v50 = v8 + int32(8)
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v48)))
	*(*int32)(unsafe.Add(mBase, uint32(v50)+4)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v50))) = v51
	goto L14
L14:
	;
	v56 = v8 + int32(8)
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v56)))
	if v58 == int32(0) {
		goto L17
	} else {
		goto L18
	}
L15:
	;
	v123 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v125 = v8 + int32(8)
	v126 = *(*int32)(unsafe.Add(mBase, uint32(v123)))
	*(*int32)(unsafe.Add(mBase, uint32(v125)+4)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v125))) = v126
	goto L35
L16:
	;
	if v58 == int32(0) {
		goto L15
	} else {
		goto L19
	}
L17:
	;
	goto L16
L18:
	;
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v56)+4))
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v58+base.B2i32(v61 == int32(0))<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v56))) = v67
	goto L17
L19:
	;
	v74 = v58
	goto L20
L20:
	;
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v74)+8))
	v77 = *(*int32)(unsafe.Add(mBase, uint32(v76)+144))
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v77)))
	*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v8))) = v78
	goto L22
L21:
	;
	goto L15
L22:
	;
	goto L24
L23:
	;
	v105 = v8 + int32(8)
	v107 = *(*int32)(unsafe.Add(mBase, uint32(v105)))
	if v107 == int32(0) {
		goto L32
	} else {
		goto L33
	}
L24:
	;
	v88 = *(*int32)(unsafe.Add(mBase, uint32(v8)))
	if v88 == int32(0) {
		goto L27
	} else {
		goto L28
	}
L26:
	;
	if v88 == int32(0) {
		goto L23
	} else {
		goto L29
	}
L27:
	;
	goto L26
L28:
	;
	v91 = *(*int32)(unsafe.Add(mBase, uint32(v8)+4))
	v97 = *(*int32)(unsafe.Add(mBase, uint32(v88+base.B2i32(v91 == int32(0))<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v8))) = v97
	goto L27
L29:
	;
	v101 = *(*int32)(unsafe.Add(mBase, uint32(v88)+8))
	v102 = F_listAddNodeTail(m, v44, v101)
	mBase = m.M
	v103 = m.ExcPending
	if v103 != 0 {
		goto L12
	} else {
		goto L30
	}
L30:
	;
	goto L24
L31:
	;
	if v107 != 0 {
		v74 = v107
		goto L20
	} else {
		goto L34
	}
L32:
	;
	goto L31
L33:
	;
	v110 = *(*int32)(unsafe.Add(mBase, uint32(v105)+4))
	v116 = *(*int32)(unsafe.Add(mBase, uint32(v107+base.B2i32(v110 == int32(0))<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v105))) = v116
	goto L32
L34:
	;
	goto L21
L35:
	;
	v131 = v8 + int32(8)
	v133 = *(*int32)(unsafe.Add(mBase, uint32(v131)))
	if v133 == int32(0) {
		goto L38
	} else {
		goto L39
	}
L36:
	;
	F_listRelease(m, v44)
	mBase = m.M
	v216 = m.ExcPending
	if v216 != 0 {
		goto L12
	} else {
		goto L61
	}
L37:
	;
	if v133 == int32(0) {
		goto L36
	} else {
		goto L40
	}
L38:
	;
	goto L37
L39:
	;
	v136 = *(*int32)(unsafe.Add(mBase, uint32(v131)+4))
	v142 = *(*int32)(unsafe.Add(mBase, uint32(v133+base.B2i32(v136 == int32(0))<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v131))) = v142
	goto L38
L40:
	;
	v149 = v133
	goto L41
L41:
	;
	v151 = *(*int32)(unsafe.Add(mBase, uint32(v149)+8))
	v152 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v151))))
	if v152&int32(8) != 0 {
		v227 = v44
		goto L2
	} else {
		goto L43
	}
L42:
	;
	goto L36
L43:
	;
	v155 = *(*int32)(unsafe.Add(mBase, uint32(v151)+144))
	v156 = *(*int32)(unsafe.Add(mBase, uint32(v155)))
	*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v8))) = v156
	goto L44
L44:
	;
	goto L46
L45:
	;
	v197 = v8 + int32(8)
	v199 = *(*int32)(unsafe.Add(mBase, uint32(v197)))
	if v199 == int32(0) {
		goto L58
	} else {
		goto L59
	}
L46:
	;
	v166 = *(*int32)(unsafe.Add(mBase, uint32(v8)))
	if v166 == int32(0) {
		goto L49
	} else {
		goto L50
	}
L47:
	;
	v183 = v8 + int32(8)
	v185 = *(*int32)(unsafe.Add(mBase, uint32(v183)))
	if v185 == int32(0) {
		goto L55
	} else {
		goto L56
	}
L48:
	;
	if v166 == int32(0) {
		goto L45
	} else {
		goto L51
	}
L49:
	;
	goto L48
L50:
	;
	v169 = *(*int32)(unsafe.Add(mBase, uint32(v8)+4))
	v175 = *(*int32)(unsafe.Add(mBase, uint32(v166+base.B2i32(v169 == int32(0))<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v8))) = v175
	goto L49
L51:
	;
	v179 = *(*int32)(unsafe.Add(mBase, uint32(v166)+8))
	v180 = F_listSearchKey(m, v44, v179)
	mBase = m.M
	v181 = m.ExcPending
	if v181 != 0 {
		goto L12
	} else {
		goto L52
	}
L52:
	;
	if v180 != 0 {
		goto L46
	} else {
		goto L53
	}
L53:
	;
	goto L47
L54:
	;
	v227 = v44
	goto L2
L55:
	;
	goto L54
L56:
	;
	v188 = *(*int32)(unsafe.Add(mBase, uint32(v183)+4))
	v194 = *(*int32)(unsafe.Add(mBase, uint32(v185+base.B2i32(v188 == int32(0))<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v183))) = v194
	goto L55
L57:
	;
	if v199 != 0 {
		v149 = v199
		goto L41
	} else {
		goto L60
	}
L58:
	;
	goto L57
L59:
	;
	v202 = *(*int32)(unsafe.Add(mBase, uint32(v197)+4))
	v208 = *(*int32)(unsafe.Add(mBase, uint32(v199+base.B2i32(v202 == int32(0))<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v197))) = v208
	goto L58
L60:
	;
	goto L42
L61:
	;
	goto L3
}
func F_getVersion(m *base.Module) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v18 int32
	_ = v18
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v63 int32
	_ = v63
	var v69 int64
	_ = v69
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	v5 = m.G0
	v7 = v5 - int32(32)
	m.G0 = v7
	v9 = F_sdsempty(m)
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		return int32(0)
	} else {
		v18 = int32(_a1908)
		for {
			v23 = v18 + int32(1)
			v24 = int32(*(*int8)(unsafe.Add(mBase, uint32(v18))))
			v25 = F___isspace_1(m, v24)
			mBase = m.M
			if v25 != 0 {
				v18 = v23
				continue
			} else {
				break
			}
			break
		}
		v26 = int32(1)
		switch v24&int32(255) + int32(-43) {
		case 0:
			v32 = v26
			v33 = int32(*(*int8)(unsafe.Add(mBase, uint32(v23))))
			v34 = v23
			v35 = v33
			v36 = v32
		default:
			v34 = v18
			v35 = v24
			v36 = v26
		case 2:
			v32 = int32(0)
			v33 = int32(*(*int8)(unsafe.Add(mBase, uint32(v23))))
			v34 = v23
			v35 = v33
			v36 = v32
		}
		v39 = v35 + int32(-48)
		if base.Ui32(int32(9)) < base.Ui32(v39) {
			v57 = int32(0)
		} else {
			v43 = int32(0)
			v44 = v34
			v45 = v39
			for {
				v47 = int32(10)
				v49 = v43*v47 - v45
				v50 = int32(*(*int8)(unsafe.Add(mBase, uint32(v44)+1)))
				v54 = v50 + int32(-48)
				if base.Ui32(v54) < base.Ui32(v47) {
					v43 = v49
					v44 = v44 + int32(1)
					v45 = v54
					continue
				} else {
					break
				}
				break
			}
			v57 = v49
		}
		if v36 != 0 {
			v63 = int32(0) - v57
		} else {
			v63 = v57
		}
		v69 = F_crc64(m, int64(0), int32(_a1910), int64(73))
		mBase = m.M
		*(*int64)(unsafe.Add(mBase, uint32(v7+int32(24)))) = v69
		*(*int32)(unsafe.Add(mBase, uint32(v7+int32(16)))) = int32(32)
		*(*int32)(unsafe.Add(mBase, uint32(v7)+12)) = int32(_a1774)
		*(*int32)(unsafe.Add(mBase, uint32(v7)+8)) = base.B2i32(int32(0) < v63)
		*(*int32)(unsafe.Add(mBase, uint32(v7)+4)) = int32(_a2283)
		*(*int32)(unsafe.Add(mBase, uint32(v7))) = int32(_a469)
		v84 = F_sdscatprintf(m, v9, int32(_a2284), v7)
		mBase = m.M
		v85 = m.ExcPending
		if v85 != 0 {
			return int32(0)
		} else {
			m.G0 = v7 + int32(32)
			return v84
		}
	}
}
func F_getexCommand(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v19 int32
	_ = v19
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
	var v41 int32
	_ = v41
	var v42 int64
	_ = v42
	var v55 int64
	_ = v55
	var v56 int64
	_ = v56
	var v57 int64
	_ = v57
	var v63 int32
	_ = v63
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
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
	var v91 int64
	_ = v91
	var v95 int32
	_ = v95
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	var v102 int64
	_ = v102
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	var v109 int32
	_ = v109
	var v113 int32
	_ = v113
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v126 int64
	_ = v126
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v129 int64
	_ = v129
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
	var v143 int32
	_ = v143
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
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
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v161 int64
	_ = v161
	var v165 int32
	_ = v165
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v185 int32
	_ = v185
	var v190 int32
	_ = v190
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v201 int64
	_ = v201
	var v209 int32
	_ = v209
	v2 = int32(0)
	v7 = m.G0
	v9 = v7 - int32(64)
	m.G0 = v9
	*(*int32)(unsafe.Add(mBase, uint32(v9)+60)) = v2
	*(*int32)(unsafe.Add(mBase, uint32(v9)+56)) = v2
	*(*int32)(unsafe.Add(mBase, uint32(v9)+52)) = v2
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v28 = F_parseExtendedCommandArgumentsOrReply(m, l0, v2, int32(2), v19, v9+int32(52), v9+int32(56), v2, v9+int32(60), v2)
	mBase = m.M
	v29 = m.ExcPending
	if v29 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v9 + int32(64)
	return
L2:
	;
	return
L3:
	;
	if v28 != 0 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v9)+40)) = int64(0)
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v9)+60))
	if v32 == int32(0) {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	v67 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v67)+4))
	v70 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+224)))
	v74 = *(*int32)(unsafe.Add(mBase, uint32(v70<<(uint(int32(2))%32))+uint32(_consts[339])))
	v75 = F_lookupKeyReadOrReply(m, l0, v68, v74)
	mBase = m.M
	v76 = m.ExcPending
	if v76 != 0 {
		goto L2
	} else {
		goto L20
	}
L6:
	;
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v9)+56))
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v9)+52))
	v40 = F_getLongLongFromObjectOrReply(m, l0, v32, v9+int32(40), int32(0))
	mBase = m.M
	v41 = m.ExcPending
	if v41 != 0 {
		goto L2
	} else {
		goto L7
	}
L7:
	;
	if v40 != 0 {
		goto L1
	} else {
		goto L8
	}
L8:
	;
	v42 = *(*int64)(unsafe.Add(mBase, uint32(v9)+40))
	if v42 < int64(1) {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	F_addReplyErrorExpireTime(m, l0)
	mBase = m.M
	v63 = m.ExcPending
	if v63 != 0 {
		goto L2
	} else {
		goto L19
	}
L10:
	;
	if v35 != 0 {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	if v35 != 0 {
		goto L14
	} else {
		goto L15
	}
L12:
	;
	if base.Ui64(int64(9223372036854775)) < base.Ui64(v42) {
		goto L9
	} else {
		goto L13
	}
L13:
	;
	goto L11
L14:
	;
	if v36&int32(12) == int32(0) {
		goto L5
	} else {
		goto L16
	}
L15:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v9)+40)) = v42 * int64(1000)
	goto L14
L16:
	;
	v55 = *(*int64)(unsafe.Add(mBase, _consts[12]))
	goto L17
L17:
	;
	v56 = *(*int64)(unsafe.Add(mBase, uint32(v9)+40))
	v57 = v55 + v56
	*(*int64)(unsafe.Add(mBase, uint32(v9)+40)) = v57
	if int64(0) < v57 {
		goto L5
	} else {
		goto L18
	}
L18:
	;
	goto L9
L19:
	;
	goto L1
L20:
	;
	if v75 == int32(0) {
		goto L1
	} else {
		goto L21
	}
L21:
	;
	v80 = F_checkType(m, l0, v75, int32(0))
	mBase = m.M
	v81 = m.ExcPending
	if v81 != 0 {
		goto L2
	} else {
		goto L22
	}
L22:
	;
	if v80 != 0 {
		goto L1
	} else {
		goto L23
	}
L23:
	;
	F_initDeferredReplyBuffer(m, l0)
	mBase = m.M
	v83 = m.ExcPending
	if v83 != 0 {
		goto L2
	} else {
		goto L24
	}
L24:
	;
	F_addReplyBulk(m, l0, v75)
	mBase = m.M
	v85 = m.ExcPending
	if v85 != 0 {
		goto L2
	} else {
		goto L25
	}
L25:
	;
	v86 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+52)))
	if v86&int32(192) == int32(0) {
		goto L27
	} else {
		goto L28
	}
L26:
	;
	F_commitDeferredReplyBuffer(m, l0, int32(1))
	mBase = m.M
	v209 = m.ExcPending
	if v209 != 0 {
		goto L2
	} else {
		goto L53
	}
L27:
	;
	v120 = *(*int32)(unsafe.Add(mBase, uint32(v9)+60))
	if v120 == int32(0) {
		goto L39
	} else {
		goto L40
	}
L28:
	;
	v91 = *(*int64)(unsafe.Add(mBase, uint32(v9)+40))
	v95 = *(*int32)(unsafe.Add(mBase, _consts[14]))
	if v95 == int32(0) {
		goto L31
	} else {
		goto L32
	}
L29:
	;
	if v113 == int32(0) {
		goto L27
	} else {
		goto L37
	}
L30:
	;
	goto L29
L31:
	;
	v101 = int32(0)
	v102 = F_commandTimeSnapshot(m)
	mBase = m.M
	if v102 < v91 {
		v113 = v101
		goto L30
	} else {
		goto L34
	}
L32:
	;
	v99 = *(*int32)(unsafe.Add(mBase, uint32(v95)+216))
	if v99 != 0 {
		v113 = int32(0)
		goto L30
	} else {
		goto L33
	}
L33:
	;
	goto L31
L34:
	;
	v104 = int32(_a20)
	v105 = *(*int32)(unsafe.Add(mBase, _consts[166]))
	v107 = *(*int32)(unsafe.Add(mBase, _consts[384]))
	if v107 != 0 {
		v113 = v101
		goto L30
	} else {
		goto L35
	}
L35:
	;
	if v105 != 0 {
		v113 = v101
		goto L30
	} else {
		goto L36
	}
L36:
	;
	v109 = *(*int32)(unsafe.Add(mBase, _consts[505]))
	v113 = base.B2i32(v109 == int32(0))
	goto L30
L37:
	;
	v116 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v117 = *(*int32)(unsafe.Add(mBase, uint32(v116)+4))
	F_deleteExpiredKeyFromOverwriteAndPropagate(m, l0, v117)
	mBase = m.M
	v119 = m.ExcPending
	if v119 != 0 {
		goto L2
	} else {
		goto L38
	}
L38:
	;
	goto L26
L39:
	;
	v165 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+53)))
	if v165&int32(1) == int32(0) {
		goto L26
	} else {
		goto L47
	}
L40:
	;
	v123 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v124 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v125 = *(*int32)(unsafe.Add(mBase, uint32(v124)+4))
	v126 = *(*int64)(unsafe.Add(mBase, uint32(v9)+40))
	v127 = F_setExpire(m, l0, v123, v125, v126)
	mBase = m.M
	v128 = m.ExcPending
	if v128 != 0 {
		goto L2
	} else {
		goto L41
	}
L41:
	;
	v129 = *(*int64)(unsafe.Add(mBase, uint32(v9)+40))
	v130 = F_createStringObjectFromLongLong(m, v129)
	mBase = m.M
	v131 = m.ExcPending
	if v131 != 0 {
		goto L2
	} else {
		goto L42
	}
L42:
	;
	v132 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v133 = *(*int32)(unsafe.Add(mBase, uint32(v132)+4))
	v135 = *(*int32)(unsafe.Add(mBase, _consts[876]))
	*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = v135
	*(*int32)(unsafe.Add(mBase, uint32(v9)+20)) = v133
	*(*int32)(unsafe.Add(mBase, uint32(v9)+24)) = v130
	F_rewriteClientCommandVector(m, l0, int32(3), v9+int32(16))
	mBase = m.M
	v143 = m.ExcPending
	if v143 != 0 {
		goto L2
	} else {
		goto L43
	}
L43:
	;
	F_decrRefCount(m, v130)
	mBase = m.M
	v145 = m.ExcPending
	if v145 != 0 {
		goto L2
	} else {
		goto L44
	}
L44:
	;
	v146 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v147 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v148 = *(*int32)(unsafe.Add(mBase, uint32(v147)+4))
	F_signalModifiedKey(m, l0, v146, v148)
	mBase = m.M
	v150 = m.ExcPending
	if v150 != 0 {
		goto L2
	} else {
		goto L45
	}
L45:
	;
	v153 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v154 = *(*int32)(unsafe.Add(mBase, uint32(v153)+4))
	v155 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v156 = *(*int32)(unsafe.Add(mBase, uint32(v155)+28))
	F_notifyKeyspaceEvent(m, int32(4), int32(_a2433), v154, v156)
	mBase = m.M
	v158 = m.ExcPending
	if v158 != 0 {
		goto L2
	} else {
		goto L46
	}
L46:
	;
	v159 = int32(_a20)
	v161 = *(*int64)(unsafe.Add(mBase, _consts[180]))
	*(*int64)(unsafe.Add(mBase, _consts[180])) = v161 + int64(1)
	goto L26
L47:
	;
	v170 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v171 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v172 = *(*int32)(unsafe.Add(mBase, uint32(v171)+4))
	v173 = F_removeExpire(m, v170, v172)
	mBase = m.M
	v174 = m.ExcPending
	if v174 != 0 {
		goto L2
	} else {
		goto L48
	}
L48:
	;
	if v173 == int32(0) {
		goto L26
	} else {
		goto L49
	}
L49:
	;
	v177 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v178 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v179 = *(*int32)(unsafe.Add(mBase, uint32(v178)+4))
	F_signalModifiedKey(m, l0, v177, v179)
	mBase = m.M
	v181 = m.ExcPending
	if v181 != 0 {
		goto L2
	} else {
		goto L50
	}
L50:
	;
	v182 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v183 = *(*int32)(unsafe.Add(mBase, uint32(v182)+4))
	v185 = *(*int32)(unsafe.Add(mBase, _consts[875]))
	*(*int32)(unsafe.Add(mBase, uint32(v9))) = v185
	*(*int32)(unsafe.Add(mBase, uint32(v9)+4)) = v183
	F_rewriteClientCommandVector(m, l0, int32(2), v9)
	mBase = m.M
	v190 = m.ExcPending
	if v190 != 0 {
		goto L2
	} else {
		goto L51
	}
L51:
	;
	v193 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v194 = *(*int32)(unsafe.Add(mBase, uint32(v193)+4))
	v195 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v196 = *(*int32)(unsafe.Add(mBase, uint32(v195)+28))
	F_notifyKeyspaceEvent(m, int32(4), int32(_a616), v194, v196)
	mBase = m.M
	v198 = m.ExcPending
	if v198 != 0 {
		goto L2
	} else {
		goto L52
	}
L52:
	;
	v199 = int32(_a20)
	v201 = *(*int64)(unsafe.Add(mBase, _consts[180]))
	*(*int64)(unsafe.Add(mBase, _consts[180])) = v201 + int64(1)
	goto L26
L53:
	;
	goto L1
}
func F_gethooktable(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v6 int32
	_ = v6
	var v10 int32
	_ = v10
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v78 int64
	_ = v78
	var v82 int32
	_ = v82
	var v98 int32
	_ = v98
	var v101 int32
	_ = v101
	var v135 int32
	_ = v135
	var v138 int32
	_ = v138
	var v141 int32
	_ = v141
	var v162 int32
	_ = v162
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v180 int32
	_ = v180
	var v184 int32
	_ = v184
	var v205 int32
	_ = v205
	var v208 int32
	_ = v208
	var v244 int32
	_ = v244
	var v245 int64
	_ = v245
	var v247 int32
	_ = v247
	var v249 int32
	_ = v249
	var v255 int32
	_ = v255
	v2 = m.G3
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v6)+8)) = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v6))) = v2 + int32(_a2700)
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v10 + int32(16)
	switch int32(2) {
	case 0:
		v67 = l0 + int32(72)
	case 1:
		v40 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
		v41 = *(*int32)(unsafe.Add(mBase, uint32(v40)+4))
		v42 = *(*int32)(unsafe.Add(mBase, uint32(v41)))
		v43 = *(*int32)(unsafe.Add(mBase, uint32(v42)+12))
		*(*int32)(unsafe.Add(mBase, uint32(l0)+96)) = int32(5)
		*(*int32)(unsafe.Add(mBase, uint32(l0)+88)) = v43
		v67 = l0 + int32(88)
	case 2:
		v37 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
		v67 = v37 + int32(96)
	default:
		v51 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
		v52 = *(*int32)(unsafe.Add(mBase, uint32(v51)+4))
		v53 = *(*int32)(unsafe.Add(mBase, uint32(v52)))
		v54 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v53)+7)))
		v55 = m.G398
		if base.Ui32(v54) < base.Ui32(int32(-2)) {
			v66 = v55
		} else {
			v66 = v53 + int32(-24)
		}
		v67 = v66
	}
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v67)))
	v71 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v72 = int32(-16)
	v74 = F_luaH_get(m, v70, v71+v72)
	mBase = m.M
	v75 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v78 = *(*int64)(unsafe.Add(mBase, uint32(v74)))
	*(*int64)(unsafe.Add(mBase, uint32(v75+v72))) = v78
	v82 = *(*int32)(unsafe.Add(mBase, uint32(v74)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v75+int32(-8)))) = v82
	v98 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v101 = v98 + int32(-16)
	v135 = m.G398
	if v101 != v135 {
		v138 = *(*int32)(unsafe.Add(mBase, uint32(v101)+8))
		v141 = v138
	} else {
		v141 = int32(-1)
	}
	if v141 == int32(5) {
		return
	} else {
		v162 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v162 + int32(-16)
		F_lua_createtable(m, l0, int32(0), int32(1))
		mBase = m.M
		v175 = m.ExcPending
		if v175 != 0 {
			return
		} else {
			v176 = m.G3
			v180 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			*(*int32)(unsafe.Add(mBase, uint32(v180)+8)) = int32(2)
			*(*int32)(unsafe.Add(mBase, uint32(v180))) = v176 + int32(_a2700)
			v184 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v184 + int32(16)
			v205 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			v208 = v205 + int32(-32)
			v244 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			v245 = *(*int64)(unsafe.Add(mBase, uint32(v208)))
			*(*int64)(unsafe.Add(mBase, uint32(v244))) = v245
			v247 = *(*int32)(unsafe.Add(mBase, uint32(v208)+8))
			*(*int32)(unsafe.Add(mBase, uint32(v244)+8)) = v247
			v249 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v249 + int32(16)
			F_lua_rawset(m, l0, int32(-10000))
			mBase = m.M
			v255 = m.ExcPending
			if v255 != 0 {
				return
			} else {
				return
			}
		}
	}
}
func F_getobjname(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v12 int32
	_ = v12
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
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
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v92 int32
	_ = v92
	var v110 int32
	_ = v110
	var v113 int32
	_ = v113
	var v120 int32
	_ = v120
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v126 int32
	_ = v126
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v136 int32
	_ = v136
	var v142 int32
	_ = v142
	var v148 int32
	_ = v148
	var v161 int32
	_ = v161
	var v164 int32
	_ = v164
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v181 int32
	_ = v181
	var v188 int32
	_ = v188
	var v191 int32
	_ = v191
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v200 int32
	_ = v200
	var v214 int32
	_ = v214
	var v220 int32
	_ = v220
	var v222 int32
	_ = v222
	var v227 int32
	_ = v227
	var v237 int32
	_ = v237
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v248 int32
	_ = v248
	var v250 int32
	_ = v250
	var v255 int32
	_ = v255
	var v260 int32
	_ = v260
	var v261 int32
	_ = v261
	var v263 int32
	_ = v263
	var v277 int32
	_ = v277
	var v293 int32
	_ = v293
	var v299 int32
	_ = v299
	var v301 int32
	_ = v301
	var v302 int32
	_ = v302
	var v323 int32
	_ = v323
	var v324 int32
	_ = v324
	var v328 int32
	_ = v328
	var v334 int32
	_ = v334
	var v344 int32
	_ = v344
	var v346 int32
	_ = v346
	var v351 int32
	_ = v351
	var v357 int32
	_ = v357
	var v363 int32
	_ = v363
	var v365 int32
	_ = v365
	var v368 int32
	_ = v368
	var v369 int32
	_ = v369
	var v372 int32
	_ = v372
	var v377 int32
	_ = v377
	var v381 int32
	_ = v381
	var v384 int32
	_ = v384
	var v388 int32
	_ = v388
	var v392 int32
	_ = v392
	var v397 int32
	_ = v397
	var v400 int32
	_ = v400
	var v401 int32
	_ = v401
	var v407 int32
	_ = v407
	var v410 int32
	_ = v410
	var v412 int32
	_ = v412
	var v419 int32
	_ = v419
	var v421 int32
	_ = v421
	var v431 int32
	_ = v431
	var v439 int32
	_ = v439
	var v447 int32
	_ = v447
	var v453 int32
	_ = v453
	var v455 int32
	_ = v455
	var v456 int32
	_ = v456
	var v458 int32
	_ = v458
	var v462 int32
	_ = v462
	var v463 int32
	_ = v463
	var v464 int32
	_ = v464
	var v487 int32
	_ = v487
	var v492 int32
	_ = v492
	var v498 int32
	_ = v498
	var v500 int32
	_ = v500
	var v503 int32
	_ = v503
	var v505 int32
	_ = v505
	var v540 int32
	_ = v540
	var v546 int32
	_ = v546
	var v549 int32
	_ = v549
	var v562 int32
	_ = v562
	var v572 int32
	_ = v572
	var v587 int32
	_ = v587
	var v591 int32
	_ = v591
	var v628 int32
	_ = v628
	var v630 int32
	_ = v630
	var v635 int32
	_ = v635
	var v638 int32
	_ = v638
	var v644 int32
	_ = v644
	var v648 int32
	_ = v648
	var v654 int32
	_ = v654
	var v661 int32
	_ = v661
	var v662 int32
	_ = v662
	var v667 int32
	_ = v667
	var v668 int32
	_ = v668
	var v675 int32
	_ = v675
	var v678 int32
	_ = v678
	var v681 int32
	_ = v681
	var v682 int32
	_ = v682
	var v687 int32
	_ = v687
	var v695 int32
	_ = v695
	var v698 int32
	_ = v698
	var v705 int32
	_ = v705
	var v706 int32
	_ = v706
	var v711 int32
	_ = v711
	var v712 int32
	_ = v712
	var v719 int32
	_ = v719
	var v722 int32
	_ = v722
	var v726 int32
	_ = v726
	var v727 int32
	_ = v727
	var v731 int32
	_ = v731
	v12 = l2
	goto L5
L1:
	;
	return v731
L2:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v727
	v731 = v726
	goto L1
L3:
	;
	v698 = m.G3
	if v628&int32(4194304) != 0 {
		goto L187
	} else {
		goto L188
	}
L4:
	;
	v681 = m.G3
	v682 = *(*int32)(unsafe.Add(mBase, uint32(v26)+28))
	if v682 != 0 {
		goto L185
	} else {
		goto L186
	}
L5:
	;
	v19 = int32(0)
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v21 = *(*int32)(unsafe.Add(mBase, uint32(v20)+8))
	if v21 != int32(6) {
		v731 = v19
		goto L1
	} else {
		goto L7
	}
L6:
	;
	v654 = m.G3
	if v628&int32(4194304) != 0 {
		goto L181
	} else {
		goto L182
	}
L7:
	;
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v20)))
	v25 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24)+6)))
	if v25 != 0 {
		v731 = v19
		goto L1
	} else {
		goto L8
	}
L8:
	;
	v26 = *(*int32)(unsafe.Add(mBase, uint32(v24)+16))
	v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if l1 == v27 {
		goto L10
	} else {
		goto L11
	}
L9:
	;
	v36 = int32(1)
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v34)+12))
	v43 = (v35-v38)>>(uint(int32(2))%32) + int32(-1)
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v26)+56))
	if v36 <= v47 {
		goto L13
	} else {
		goto L14
	}
L10:
	;
	v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	*(*int32)(unsafe.Add(mBase, uint32(l1)+12)) = v30
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v20)))
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v32)+16))
	v34 = v33
	v35 = v30
	goto L9
L11:
	;
	v29 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v34 = v26
	v35 = v29
	goto L9
L12:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v84
	v86 = m.G3
	if v84 == int32(0) {
		goto L23
	} else {
		goto L24
	}
L13:
	;
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v26)+24))
	v53 = int32(0)
	v54 = v12 + v36
	goto L16
L14:
	;
	v84 = int32(0)
	goto L12
L15:
	;
	v84 = int32(0)
	goto L12
L16:
	;
	v61 = v51 + v53*int32(12)
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v61)+4))
	if v43 < v62 {
		goto L15
	} else {
		goto L18
	}
L17:
	;
	goto L15
L18:
	;
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v61)+8))
	if v64 <= v43 {
		v71 = v54
		goto L19
	} else {
		goto L20
	}
L19:
	;
	v73 = v53 + int32(1)
	if v73 != v47 {
		v53 = v73
		v54 = v71
		goto L16
	} else {
		goto L22
	}
L20:
	;
	v67 = v54 + int32(-1)
	if v67 != 0 {
		v71 = v67
		goto L19
	} else {
		goto L21
	}
L21:
	;
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v61)))
	v84 = v68 + int32(16)
	goto L12
L22:
	;
	goto L17
L23:
	;
	v92 = int32(0)
	v110 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v26)+75)))
	if base.Ui32(int32(250)) < base.Ui32(v110) {
		v591 = v92
		goto L28
	} else {
		goto L29
	}
L24:
	;
	return v86 + int32(_a2648)
L25:
	;
	goto L6
L26:
	;
	v648 = int32(base.Ui32(v628) >> (uint(int32(23)) % 32))
	if base.Ui32(v648) < base.Ui32(int32(base.Ui32(v628)>>(uint(int32(6))%32))&int32(255)) {
		v12 = v648
		goto L5
	} else {
		goto L180
	}
L27:
	;
	v630 = v628 & int32(63)
	if v630 == int32(0) {
		goto L26
	} else {
		goto L178
	}
L28:
	;
	v628 = v591
	goto L27
L29:
	;
	v113 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v26)+74)))
	if v113&int32(5) == int32(4) {
		v591 = v92
		goto L28
	} else {
		goto L30
	}
L30:
	;
	v120 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v26)+73)))
	if base.Ui32(v110) < base.Ui32(v113&int32(1)+v120) {
		v591 = v92
		goto L28
	} else {
		goto L31
	}
L31:
	;
	v123 = *(*int32)(unsafe.Add(mBase, uint32(v26)+36))
	v124 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v26)+72)))
	if v124 < v123 {
		v591 = v92
		goto L28
	} else {
		goto L32
	}
L32:
	;
	v126 = *(*int32)(unsafe.Add(mBase, uint32(v26)+44))
	if v126 < int32(1) {
		v591 = v92
		goto L28
	} else {
		goto L33
	}
L33:
	;
	v129 = *(*int32)(unsafe.Add(mBase, uint32(v26)+48))
	v130 = int32(0)
	if base.B2i32(v129 == v130)|base.B2i32(v129 == v126) == v130 {
		v591 = v92
		goto L28
	} else {
		goto L34
	}
L34:
	;
	v136 = *(*int32)(unsafe.Add(mBase, uint32(v26)+12))
	v142 = *(*int32)(unsafe.Add(mBase, uint32(v136+v126<<(uint(int32(2))%32)+int32(-4))))
	if v142&int32(63) != int32(30) {
		v591 = v92
		goto L28
	} else {
		goto L35
	}
L35:
	;
	v148 = v126 + int32(-1)
	if int32(1) <= v43 {
		goto L37
	} else {
		goto L38
	}
L36:
	;
	v587 = *(*int32)(unsafe.Add(mBase, uint32(v136+v572<<(uint(int32(2))%32))))
	v591 = v587
	goto L28
L37:
	;
	v161 = int32(0)
	v164 = v148
	goto L39
L38:
	;
	v572 = v148
	goto L36
L39:
	;
	v178 = v136 + v161<<(uint(int32(2))%32)
	v179 = *(*int32)(unsafe.Add(mBase, uint32(v178)))
	v181 = v179 & int32(63)
	if base.Ui32(v181) <= base.Ui32(int32(37)) {
		goto L41
	} else {
		goto L42
	}
L40:
	;
	v572 = v549
	goto L36
L41:
	;
	v188 = int32(base.Ui32(v179)>>(uint(int32(6))%32)) & int32(255)
	if base.Ui32(v188) < base.Ui32(v110) {
		goto L43
	} else {
		goto L44
	}
L42:
	;
	v628 = int32(0)
	goto L27
L43:
	;
	v191 = m.G400
	v193 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v191+v181))))
	v194 = base.I32_extend8_s(v193)
	v195 = int32(0)
	switch v193 & int32(3) {
	default:
		goto L48
	case 1:
		goto L47
	case 2:
		goto L46
	case 3:
		v323 = v195
		v324 = v195
		goto L45
	}
L44:
	;
	v628 = int32(0)
	goto L27
L45:
	;
	if int32(-1) < v194 {
		goto L81
	} else {
		goto L82
	}
L46:
	;
	v255 = int32(base.Ui32(v179)>>(uint(int32(14))%32)) + int32(-131071)
	if v193&int32(48) != int32(32) {
		v323 = v195
		v324 = v255
		goto L45
	} else {
		goto L70
	}
L47:
	;
	v243 = int32(base.Ui32(v179) >> (uint(int32(14)) % 32))
	v244 = int32(48)
	if v193&v244 != v244 {
		v323 = v195
		v324 = v243
		goto L45
	} else {
		goto L68
	}
L48:
	;
	v200 = int32(base.Ui32(v179) >> (uint(int32(23)) % 32))
	switch int32(base.Ui32(v193)>>(uint(int32(4))%32)) & int32(3) {
	default:
		goto L52
	case 1:
		goto L49
	case 2:
		goto L51
	case 3:
		goto L50
	}
L49:
	;
	v220 = int32(base.Ui32(v179) >> (uint(int32(14)) % 32))
	v222 = v220 & int32(511)
	switch int32(base.Ui32(v194)>>(uint(int32(2))%32)) & int32(3) {
	default:
		goto L61
	case 1:
		v323 = v222
		v324 = v200
		goto L45
	case 2:
		goto L60
	case 3:
		goto L59
	}
L50:
	;
	if int32(-1) < v179 {
		goto L55
	} else {
		goto L56
	}
L51:
	;
	if base.Ui32(v200) < base.Ui32(v110) {
		goto L49
	} else {
		goto L54
	}
L52:
	;
	if base.Ui32(v179) < base.Ui32(int32(_a821)) {
		goto L49
	} else {
		goto L53
	}
L53:
	;
	v628 = int32(0)
	goto L27
L54:
	;
	v628 = int32(0)
	goto L27
L55:
	;
	if base.Ui32(v200) < base.Ui32(v110) {
		goto L49
	} else {
		goto L58
	}
L56:
	;
	v214 = *(*int32)(unsafe.Add(mBase, uint32(v26)+40))
	if v200&int32(255) < v214 {
		goto L49
	} else {
		goto L57
	}
L57:
	;
	v628 = int32(0)
	goto L27
L58:
	;
	v628 = int32(0)
	goto L27
L59:
	;
	if base.Ui32(v222) < base.Ui32(int32(256)) {
		goto L64
	} else {
		goto L65
	}
L60:
	;
	if base.Ui32(v222) < base.Ui32(v110) {
		v323 = v222
		v324 = v200
		goto L45
	} else {
		goto L63
	}
L61:
	;
	v227 = int32(0)
	if v222 == v227 {
		v323 = v227
		v324 = v200
		goto L45
	} else {
		goto L62
	}
L62:
	;
	v591 = v227
	goto L28
L63:
	;
	v628 = int32(0)
	goto L27
L64:
	;
	if base.Ui32(v222) < base.Ui32(v110) {
		v323 = v222
		v324 = v200
		goto L45
	} else {
		goto L67
	}
L65:
	;
	v237 = *(*int32)(unsafe.Add(mBase, uint32(v26)+40))
	if v220&int32(255) < v237 {
		v323 = v222
		v324 = v200
		goto L45
	} else {
		goto L66
	}
L66:
	;
	v628 = int32(0)
	goto L27
L67:
	;
	v628 = int32(0)
	goto L27
L68:
	;
	v248 = int32(0)
	v250 = *(*int32)(unsafe.Add(mBase, uint32(v26)+40))
	if v243 < v250 {
		v323 = v248
		v324 = v243
		goto L45
	} else {
		goto L69
	}
L69:
	;
	v591 = v248
	goto L28
L70:
	;
	v260 = int32(0)
	v261 = v161 + v255
	v263 = v261 + int32(1)
	if v263 < v260 {
		v591 = v260
		goto L28
	} else {
		goto L71
	}
L71:
	;
	if v126 <= v263 {
		v591 = v260
		goto L28
	} else {
		goto L72
	}
L72:
	;
	if v263 != 0 {
		goto L73
	} else {
		goto L74
	}
L73:
	;
	v277 = int32(0)
	goto L76
L74:
	;
	v323 = int32(0)
	v324 = v255
	goto L45
L75:
	;
	v302 = int32(0)
	if v301&int32(1) != 0 {
		v591 = v302
		goto L28
	} else {
		goto L80
	}
L76:
	;
	v293 = *(*int32)(unsafe.Add(mBase, uint32(v136+(v261-v277)<<(uint(int32(2))%32))))
	if v293&int32(8372287) != int32(34) {
		v301 = v277
		goto L75
	} else {
		goto L78
	}
L77:
	;
	v301 = v263
	goto L75
L78:
	;
	v299 = v277 + int32(1)
	if v299 != v263 {
		v277 = v299
		goto L76
	} else {
		goto L79
	}
L79:
	;
	goto L77
L80:
	;
	v323 = v302
	v324 = v255
	goto L45
L81:
	;
	if int32(base.Ui32(v194&int32(64))>>(uint(int32(6))%32)) != 0 {
		goto L85
	} else {
		goto L86
	}
L82:
	;
	v328 = int32(0)
	if v126 <= v161+int32(2) {
		v591 = v328
		goto L28
	} else {
		goto L83
	}
L83:
	;
	v334 = *(*int32)(unsafe.Add(mBase, uint32(v178+int32(4))))
	if v334&int32(63) != int32(22) {
		v591 = v328
		goto L28
	} else {
		goto L84
	}
L84:
	;
	goto L81
L85:
	;
	v344 = v161
	goto L87
L86:
	;
	v344 = v164
	goto L87
L87:
	;
	if v188 == v12 {
		goto L88
	} else {
		goto L89
	}
L88:
	;
	v346 = v344
	goto L90
L89:
	;
	v346 = v164
	goto L90
L90:
	;
	switch v181 + int32(-2) {
	case 0:
		goto L106
	case 1:
		goto L105
	case 2, 6:
		goto L104
	case 3, 5:
		goto L103
	default:
		v546 = v161
		v549 = v346
		goto L91
	case 9:
		goto L102
	case 19:
		goto L101
	case 20:
		goto L98
	case 26, 27:
		goto L97
	case 28:
		goto L96
	case 29, 30:
		goto L99
	case 31:
		goto L100
	case 32:
		goto L95
	case 34:
		goto L94
	case 35:
		goto L93
	}
L91:
	;
	v562 = v546 + int32(1)
	if v562 < v43 {
		v161 = v562
		v164 = v549
		goto L39
	} else {
		goto L177
	}
L92:
	;
	if v12 == int32(255) {
		goto L174
	} else {
		goto L175
	}
L93:
	;
	v498 = int32(0)
	if v113&int32(6) != int32(2) {
		v591 = v498
		goto L28
	} else {
		goto L165
	}
L94:
	;
	v455 = int32(0)
	v456 = *(*int32)(unsafe.Add(mBase, uint32(v26)+52))
	if v456 <= v324 {
		v591 = v455
		goto L28
	} else {
		goto L158
	}
L95:
	;
	v447 = int32(0)
	if v324 < int32(1) {
		goto L153
	} else {
		goto L154
	}
L96:
	;
	if v324 < int32(2) {
		v546 = v161
		v549 = v346
		goto L91
	} else {
		goto L151
	}
L97:
	;
	v412 = int32(0)
	if v324 == v412 {
		goto L136
	} else {
		goto L137
	}
L98:
	;
	v397 = int32(0)
	v400 = int32(1)
	v401 = v161 + v324 + v400
	if (base.B2i32(v401 <= v161)|base.B2i32(v43 < v401))&v400 != 0 {
		goto L130
	} else {
		goto L131
	}
L99:
	;
	if base.Ui32(v188+int32(3)) < base.Ui32(v110) {
		goto L98
	} else {
		goto L129
	}
L100:
	;
	v384 = int32(0)
	if v323 == v384 {
		v591 = v384
		goto L28
	} else {
		goto L124
	}
L101:
	;
	if v324 < v323 {
		v546 = v161
		v549 = v346
		goto L91
	} else {
		goto L123
	}
L102:
	;
	v377 = v188 + int32(1)
	if base.Ui32(v377) < base.Ui32(v110) {
		goto L118
	} else {
		goto L119
	}
L103:
	;
	v368 = *(*int32)(unsafe.Add(mBase, uint32(v26)+8))
	v369 = int32(4)
	v372 = *(*int32)(unsafe.Add(mBase, uint32(v368+v324<<(uint(v369)%32))+8))
	if v372 == v369 {
		v546 = v161
		v549 = v346
		goto L91
	} else {
		goto L117
	}
L104:
	;
	if v324 < v124 {
		v546 = v161
		v549 = v346
		goto L91
	} else {
		goto L116
	}
L105:
	;
	if v324 < v12 {
		goto L110
	} else {
		goto L111
	}
L106:
	;
	if v323 != int32(1) {
		v546 = v161
		v549 = v346
		goto L91
	} else {
		goto L107
	}
L107:
	;
	v351 = int32(0)
	if v126 <= v161+int32(2) {
		v591 = v351
		goto L28
	} else {
		goto L108
	}
L108:
	;
	v357 = *(*int32)(unsafe.Add(mBase, uint32(v178+int32(4))))
	if v357&int32(8372287) != int32(34) {
		v546 = v161
		v549 = v346
		goto L91
	} else {
		goto L109
	}
L109:
	;
	v591 = v351
	goto L28
L110:
	;
	v363 = v346
	goto L112
L111:
	;
	v363 = v161
	goto L112
L112:
	;
	if v12 < v188 {
		goto L113
	} else {
		goto L114
	}
L113:
	;
	v365 = v346
	goto L115
L114:
	;
	v365 = v363
	goto L115
L115:
	;
	v546 = v161
	v549 = v365
	goto L91
L116:
	;
	v628 = int32(0)
	goto L27
L117:
	;
	v628 = int32(0)
	goto L27
L118:
	;
	if v12 == v377 {
		goto L120
	} else {
		goto L121
	}
L119:
	;
	v628 = int32(0)
	goto L27
L120:
	;
	v381 = v161
	goto L122
L121:
	;
	v381 = v346
	goto L122
L122:
	;
	v546 = v161
	v549 = v381
	goto L91
L123:
	;
	v628 = int32(0)
	goto L27
L124:
	;
	v388 = v188 + int32(2)
	if base.Ui32(v110) <= base.Ui32(v323+v388) {
		v591 = v384
		goto L28
	} else {
		goto L125
	}
L125:
	;
	if v12 < v388 {
		goto L126
	} else {
		goto L127
	}
L126:
	;
	v392 = v346
	goto L128
L127:
	;
	v392 = v161
	goto L128
L128:
	;
	v546 = v161
	v549 = v392
	goto L91
L129:
	;
	v628 = int32(0)
	goto L27
L130:
	;
	v407 = v397
	goto L132
L131:
	;
	v407 = v324
	goto L132
L132:
	;
	if v12 == int32(255) {
		goto L133
	} else {
		goto L134
	}
L133:
	;
	v410 = v397
	goto L135
L134:
	;
	v410 = v407
	goto L135
L135:
	;
	v546 = v410 + v161
	v549 = v346
	goto L91
L136:
	;
	if v323 != 0 {
		goto L140
	} else {
		goto L141
	}
L137:
	;
	if v110 < v324+v188 {
		v591 = v412
		goto L28
	} else {
		goto L138
	}
L138:
	;
	goto L136
L139:
	;
	if v12 < v188 {
		goto L148
	} else {
		goto L149
	}
L140:
	;
	v431 = v323 + int32(-1)
	if v431 == int32(0) {
		goto L139
	} else {
		goto L146
	}
L141:
	;
	v419 = *(*int32)(unsafe.Add(mBase, uint32(v178+int32(4))))
	v421 = v419 & int32(63)
	if base.Ui32(v421+int32(-28)) < base.Ui32(int32(3)) {
		goto L142
	} else {
		goto L143
	}
L142:
	;
	if base.Ui32(v419) <= base.Ui32(int32(8388607)) {
		goto L139
	} else {
		goto L145
	}
L143:
	;
	if v421 != int32(34) {
		v591 = v412
		goto L28
	} else {
		goto L144
	}
L144:
	;
	goto L142
L145:
	;
	v591 = v412
	goto L28
L146:
	;
	if v110 < v431+v188 {
		v591 = v412
		goto L28
	} else {
		goto L147
	}
L147:
	;
	goto L139
L148:
	;
	v439 = v164
	goto L150
L149:
	;
	v439 = v161
	goto L150
L150:
	;
	v546 = v161
	v549 = v439
	goto L91
L151:
	;
	if v188+v324+int32(-1) <= v110 {
		v546 = v161
		v549 = v346
		goto L91
	} else {
		goto L152
	}
L152:
	;
	v628 = int32(0)
	goto L27
L153:
	;
	if v323 != 0 {
		v546 = v161
		v549 = v346
		goto L91
	} else {
		goto L156
	}
L154:
	;
	if v110 <= v324+v188 {
		v591 = v447
		goto L28
	} else {
		goto L155
	}
L155:
	;
	goto L153
L156:
	;
	v453 = v161 + int32(1)
	if v453 < v148 {
		v546 = v453
		v549 = v346
		goto L91
	} else {
		goto L157
	}
L157:
	;
	v591 = v447
	goto L28
L158:
	;
	v458 = *(*int32)(unsafe.Add(mBase, uint32(v26)+16))
	v462 = *(*int32)(unsafe.Add(mBase, uint32(v458+v324<<(uint(int32(2))%32))))
	v463 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v462)+72)))
	v464 = v161 + v463
	if v126 <= v464 {
		v591 = v455
		goto L28
	} else {
		goto L159
	}
L159:
	;
	if v463 == int32(0) {
		goto L92
	} else {
		goto L160
	}
L160:
	;
	v487 = int32(1)
	goto L161
L161:
	;
	v492 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v178+v487<<(uint(int32(2))%32)))))
	if v492&int32(59) != 0 {
		v591 = v455
		goto L28
	} else {
		goto L163
	}
L163:
	;
	if v487 == v463 {
		goto L92
	} else {
		goto L164
	}
L164:
	;
	v487 = v487 + int32(1)
	goto L161
L165:
	;
	v500 = v324 + int32(-1)
	if v324 != 0 {
		goto L166
	} else {
		goto L167
	}
L166:
	;
	if v110 < v500+v188 {
		v591 = v498
		goto L28
	} else {
		goto L173
	}
L167:
	;
	v503 = *(*int32)(unsafe.Add(mBase, uint32(v178+int32(4))))
	v505 = v503 & int32(63)
	if base.Ui32(v505+int32(-28)) < base.Ui32(int32(3)) {
		goto L168
	} else {
		goto L169
	}
L168:
	;
	if base.Ui32(int32(8388607)) < base.Ui32(v503) {
		v591 = v498
		goto L28
	} else {
		goto L171
	}
L169:
	;
	if v505 != int32(34) {
		v591 = v498
		goto L28
	} else {
		goto L170
	}
L170:
	;
	goto L168
L171:
	;
	if v500+v188 <= v110 {
		v546 = v161
		v549 = v346
		goto L91
	} else {
		goto L172
	}
L172:
	;
	v591 = v498
	goto L28
L173:
	;
	v546 = v161
	v549 = v346
	goto L91
L174:
	;
	v540 = v161
	goto L176
L175:
	;
	v540 = v464
	goto L176
L176:
	;
	v546 = v540
	v549 = v346
	goto L91
L177:
	;
	goto L40
L178:
	;
	switch v630 + int32(-4) {
	case 0:
		goto L4
	case 1:
		goto L179
	case 2:
		goto L25
	default:
		v731 = v19
		goto L1
	case 7:
		goto L3
	}
L179:
	;
	v635 = m.G3
	v638 = *(*int32)(unsafe.Add(mBase, uint32(v26)+8))
	v644 = *(*int32)(unsafe.Add(mBase, uint32(v638+int32(base.Ui32(v628)>>(uint(int32(10))%32))&int32(4194288))))
	v726 = v635 + int32(_a2649)
	v727 = v644 + int32(16)
	goto L2
L180:
	;
	v731 = v19
	goto L1
L181:
	;
	v661 = m.G3
	v662 = *(*int32)(unsafe.Add(mBase, uint32(v26)+8))
	v667 = v662 + int32(base.Ui32(v628)>>(uint(int32(10))%32))&int32(4080)
	v668 = *(*int32)(unsafe.Add(mBase, uint32(v667)+8))
	if v668 == int32(4) {
		goto L183
	} else {
		goto L184
	}
L182:
	;
	v726 = v654 + int32(_a2650)
	v727 = v654 + int32(_a319)
	goto L2
L183:
	;
	v675 = m.G3
	v678 = *(*int32)(unsafe.Add(mBase, uint32(v667)))
	v726 = v675 + int32(_a2650)
	v727 = v678 + int32(16)
	goto L2
L184:
	;
	v726 = v661 + int32(_a2650)
	v727 = v661 + int32(_a319)
	goto L2
L185:
	;
	v687 = m.G3
	v695 = *(*int32)(unsafe.Add(mBase, uint32(v682+int32(base.Ui32(v628)>>(uint(int32(21))%32))&int32(2044))))
	v726 = v687 + int32(_a2651)
	v727 = v695 + int32(16)
	goto L2
L186:
	;
	v726 = v681 + int32(_a2651)
	v727 = v681 + int32(_a319)
	goto L2
L187:
	;
	v705 = m.G3
	v706 = *(*int32)(unsafe.Add(mBase, uint32(v26)+8))
	v711 = v706 + int32(base.Ui32(v628)>>(uint(int32(10))%32))&int32(4080)
	v712 = *(*int32)(unsafe.Add(mBase, uint32(v711)+8))
	if v712 == int32(4) {
		goto L189
	} else {
		goto L190
	}
L188:
	;
	v726 = v698 + int32(_a2652)
	v727 = v698 + int32(_a319)
	goto L2
L189:
	;
	v719 = m.G3
	v722 = *(*int32)(unsafe.Add(mBase, uint32(v711)))
	v726 = v719 + int32(_a2652)
	v727 = v722 + int32(16)
	goto L2
L190:
	;
	v726 = v705 + int32(_a2652)
	v727 = v705 + int32(_a319)
	goto L2
}
func F_getpeername(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v15 int32
	_ = v15
	v4 = int32(0)
	v7 = m.Env.X__syscall_getpeername(m, l0, l1, l2, v4, v4, v4)
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
func F_getpid(m *base.Module) int32 {
	return int32(42)
}
func F_getpwnam_r(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	return int32(44)
}
func F_getsockopt(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v21 int32
	_ = v21
	var v27 int32
	_ = v27
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v52 int64
	_ = v52
	var v54 int32
	_ = v54
	var v63 int32
	_ = v63
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
	var v78 int32
	_ = v78
	v8 = m.G0
	v10 = v8 - int32(16)
	m.G0 = v10
	v13 = m.Env.X__syscall_getsockopt(m, l0, l1, l2, l3, l4, int32(0))
	mBase = m.M
	if l1 != int32(1) {
		v69 = v13
		if base.Ui32(v69) < base.Ui32(int32(-4095)) {
			v77 = v69
		} else {
			v72 = F___errno_location(m)
			mBase = m.M
			*(*int32)(unsafe.Add(mBase, uint32(v72))) = int32(0) - v69
			v77 = int32(-1)
		}
		v78 = v77
	} else {
		if v13 != int32(-50) {
			v69 = v13
			if base.Ui32(v69) < base.Ui32(int32(-4095)) {
				v77 = v69
			} else {
				v72 = F___errno_location(m)
				mBase = m.M
				*(*int32)(unsafe.Add(mBase, uint32(v72))) = int32(0) - v69
				v77 = int32(-1)
			}
			v78 = v77
		} else {
			switch l2 + int32(-63) {
			case 0, 1:
				if l2 == int32(63) {
					v63 = int32(29)
				} else {
					v63 = l2
				}
				if v63 == int32(64) {
					v66 = int32(35)
				} else {
					v66 = v63
				}
				v68 = m.Env.X__syscall_getsockopt(m, l0, int32(1), v66, l3, l4, int32(0))
				mBase = m.M
				v69 = v68
				if base.Ui32(v69) < base.Ui32(int32(-4095)) {
					v77 = v69
				} else {
					v72 = F___errno_location(m)
					mBase = m.M
					*(*int32)(unsafe.Add(mBase, uint32(v72))) = int32(0) - v69
					v77 = int32(-1)
				}
				v78 = v77
			default:
				v69 = int32(-50)
				if base.Ui32(v69) < base.Ui32(int32(-4095)) {
					v77 = v69
				} else {
					v72 = F___errno_location(m)
					mBase = m.M
					*(*int32)(unsafe.Add(mBase, uint32(v72))) = int32(0) - v69
					v77 = int32(-1)
				}
				v78 = v77
			case 3, 4:
				v21 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
				if base.Ui32(int32(15)) < base.Ui32(v21) {
					*(*int32)(unsafe.Add(mBase, uint32(v10)+4)) = int32(8)
					if l2 == int32(66) {
						v40 = int32(20)
					} else {
						v40 = l2
					}
					if v40 == int32(67) {
						v43 = int32(21)
					} else {
						v43 = v40
					}
					v48 = int32(0)
					v49 = m.Env.X__syscall_getsockopt(m, l0, int32(1), v43, v10+int32(8), v10+int32(4), v48)
					mBase = m.M
					if v49 < v48 {
						v69 = v49
					} else {
						v52 = int64(*(*int32)(unsafe.Add(mBase, uint32(v10)+8)))
						*(*int64)(unsafe.Add(mBase, uint32(l3))) = v52
						v54 = *(*int32)(unsafe.Add(mBase, uint32(v10)+12))
						*(*int32)(unsafe.Add(mBase, uint32(l3)+8)) = v54
						*(*int32)(unsafe.Add(mBase, uint32(l4))) = int32(16)
						v69 = v49
					}
					if base.Ui32(v69) < base.Ui32(int32(-4095)) {
						v77 = v69
					} else {
						v72 = F___errno_location(m)
						mBase = m.M
						*(*int32)(unsafe.Add(mBase, uint32(v72))) = int32(0) - v69
						v77 = int32(-1)
					}
					v78 = v77
				} else {
					v27 = F___errno_location(m)
					mBase = m.M
					*(*int32)(unsafe.Add(mBase, uint32(v27))) = int32(28)
					v78 = int32(-1)
				}
			}
		}
	}
	m.G0 = v10 + int32(16)
	return v78
}
func F_gfind_nodef(m *base.Module, l0 int32) int32 {
	var v2 int32
	_ = v2
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	v2 = m.G3
	v6 = F_luaL_error(m, l0, v2+int32(_a2726), int32(0))
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		return v6
	}
}
func F_gmatch(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	var v42 int32
	_ = v42
	var v47 int32
	_ = v47
	var v51 int32
	_ = v51
	var v56 int32
	_ = v56
	v4 = F_luaL_checklstring(m, l0, int32(1), int32(0))
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		v10 = F_luaL_checklstring(m, l0, int32(2), int32(0))
		mBase = m.M
		v11 = m.ExcPending
		if v11 != 0 {
			return int32(0)
		} else {
			v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
			v20 = v17 + int32(32)
			if base.Ui32(v20) <= base.Ui32(v16) {
			} else {
				v24 = v16
				for {
					*(*int32)(unsafe.Add(mBase, uint32(v24)+8)) = int32(0)
					v28 = v24 + int32(16)
					if base.Ui32(v28) < base.Ui32(v20) {
						v24 = v28
						continue
					} else {
						break
					}
					break
				}
			}
			*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v20
			v42 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			*(*int32)(unsafe.Add(mBase, uint32(v42)+8)) = int32(3)
			*(*float64)(unsafe.Add(mBase, uint32(v42))) = base.F64_convert_i32_s(int32(0))
			v47 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v47 + int32(16)
			v51 = m.G5
			F_lua_pushcclosure(m, l0, v51+int32(1328), int32(3))
			mBase = m.M
			v56 = m.ExcPending
			if v56 != 0 {
				return int32(0)
			} else {
				return int32(1)
			}
		}
	}
}
func F_gmatch_aux(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v39 int32
	_ = v39
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v53 int32
	_ = v53
	var v58 int32
	_ = v58
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v74 int32
	_ = v74
	var v82 int32
	_ = v82
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v99 int32
	_ = v99
	v8 = m.G0
	v10 = v8 - int32(288)
	m.G0 = v10
	v16 = F_lua_tolstring(m, l0, int32(-10003), v10+int32(12))
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v22 = F_lua_tolstring(m, l0, int32(-10004), int32(0))
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = v16
	*(*int32)(unsafe.Add(mBase, uint32(v10)+24)) = l0
	v26 = *(*int32)(unsafe.Add(mBase, uint32(v10)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v10)+20)) = v16 + v26
	v30 = F_lua_tointeger(m, l0, int32(-10005))
	mBase = m.M
	v31 = m.ExcPending
	if v31 != 0 {
		goto L1
	} else {
		goto L5
	}
L4:
	;
	m.G0 = v10 + int32(288)
	return v99
L5:
	;
	if base.Ui32(v26) < base.Ui32(v30) {
		v99 = int32(0)
		goto L4
	} else {
		goto L6
	}
L6:
	;
	v39 = v16 + v30
	goto L7
L7:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+28)) = int32(0)
	v45 = F_match(m, v10+int32(16), v39, v22)
	mBase = m.M
	v46 = m.ExcPending
	if v46 != 0 {
		goto L1
	} else {
		goto L10
	}
L8:
	;
	v99 = int32(0)
	goto L4
L9:
	;
	v93 = v39 + int32(1)
	v94 = *(*int32)(unsafe.Add(mBase, uint32(v10)+20))
	if base.Ui32(v93) <= base.Ui32(v94) {
		v39 = v93
		goto L7
	} else {
		goto L26
	}
L10:
	;
	if v45 == int32(0) {
		goto L9
	} else {
		goto L11
	}
L11:
	;
	v53 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v53)+8)) = int32(3)
	*(*float64)(unsafe.Add(mBase, uint32(v53))) = base.F64_convert_i32_s(v45 - v16 + base.B2i32(v45 == v39))
	v58 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v58 + int32(16)
	goto L12
L12:
	;
	F_lua_replace(m, l0, int32(-10005))
	mBase = m.M
	v64 = m.ExcPending
	if v64 != 0 {
		goto L1
	} else {
		goto L13
	}
L13:
	;
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v10)+28))
	v66 = m.G3
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v10)+24))
	if v65 != 0 {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v69 = v65
	goto L16
L15:
	;
	v69 = int32(1)
	goto L16
L16:
	;
	if v39 != 0 {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	v70 = v69
	goto L19
L18:
	;
	v70 = v65
	goto L19
L19:
	;
	F_luaL_checkstack(m, v67, v70, v66+int32(_a2727))
	mBase = m.M
	v74 = m.ExcPending
	if v74 != 0 {
		goto L1
	} else {
		goto L20
	}
L20:
	;
	if v70 < int32(1) {
		v99 = v65
		goto L4
	} else {
		goto L21
	}
L21:
	;
	v82 = int32(0)
	goto L22
L22:
	;
	F_push_onecapture(m, v10+int32(16), v82, v39, v45)
	mBase = m.M
	v88 = m.ExcPending
	if v88 != 0 {
		goto L1
	} else {
		goto L24
	}
L23:
	;
	v99 = v70
	goto L4
L24:
	;
	v90 = v82 + int32(1)
	if v90 != v70 {
		v82 = v90
		goto L22
	} else {
		goto L25
	}
L25:
	;
	goto L23
L26:
	;
	goto L8
}
