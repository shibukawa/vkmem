package p0

import (
	base "github.com/shibukawa/vkmem/internal/aot/vkaot/base"
	"unsafe"
)

func F_d2string(m *base.Module, l0 int32, l1 int32, l2 float64) int32 {
	mBase := m.M
	_ = mBase
	var v8 float64
	_ = v8
	var v9 int64
	_ = v9
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v59 int64
	_ = v59
	var v61 int64
	_ = v61
	var v64 int32
	_ = v64
	var v70 int32
	_ = v70
	var v74 int32
	_ = v74
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v81 int64
	_ = v81
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v87 int32
	_ = v87
	var v95 int32
	_ = v95
	var v96 int64
	_ = v96
	var v117 int32
	_ = v117
	var v122 int32
	_ = v122
	var v129 int32
	_ = v129
	var v134 int32
	_ = v134
	var v136 int32
	_ = v136
	var v140 int64
	_ = v140
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v148 int32
	_ = v148
	var v151 int32
	_ = v151
	var v154 int32
	_ = v154
	var v161 int64
	_ = v161
	var v163 int32
	_ = v163
	var v167 int64
	_ = v167
	var v168 int64
	_ = v168
	var v177 int32
	_ = v177
	var v180 int32
	_ = v180
	var v187 int32
	_ = v187
	var v188 int64
	_ = v188
	var v190 int32
	_ = v190
	var v195 int32
	_ = v195
	var v204 int32
	_ = v204
	var v208 int32
	_ = v208
	var v218 int32
	_ = v218
	var v225 int32
	_ = v225
	var v230 int32
	_ = v230
	var v232 int32
	_ = v232
	var v237 int32
	_ = v237
	v8 = base.F64_abs(l2)
	v9 = base.I64_reinterpret_f64(v8)
	if base.Ui64(v9) < base.Ui64(int64(9218868437227405313)) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	if v9 != int64(9218868437227405312) {
		goto L5
	} else {
		goto L6
	}
L2:
	;
	v14 = F_snprintf(m, l0, l1, int32(_a_F_d2string_0), int32(0))
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	return int32(0)
L4:
	;
	return v14
L5:
	;
	if base.F64_ne(l2, float64(0)) != 0 {
		goto L11
	} else {
		goto L12
	}
L6:
	;
	if base.F64_lt(l2, float64(0)) == int32(0) {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	v32 = F_snprintf(m, l0, l1, int32(_a_F_d2string_1), int32(0))
	mBase = m.M
	v33 = m.ExcPending
	if v33 != 0 {
		goto L3
	} else {
		goto L10
	}
L8:
	;
	v27 = F_snprintf(m, l0, l1, int32(_a_F_d2string_2), int32(0))
	mBase = m.M
	v28 = m.ExcPending
	if v28 != 0 {
		goto L3
	} else {
		goto L9
	}
L9:
	;
	return v27
L10:
	;
	return v32
L11:
	;
	if base.F64_gt(v8, float64(4.611686018427388e+18)) != 0 {
		goto L18
	} else {
		goto L19
	}
L12:
	;
	if base.F64_lt(base.F64_div(float64(1), l2), float64(0)) == int32(0) {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	v50 = F_snprintf(m, l0, l1, int32(_a_F_d2string_3), int32(0))
	mBase = m.M
	v51 = m.ExcPending
	if v51 != 0 {
		goto L3
	} else {
		goto L16
	}
L14:
	;
	v45 = F_snprintf(m, l0, l1, int32(_a_F_d2string_4), int32(0))
	mBase = m.M
	v46 = m.ExcPending
	if v46 != 0 {
		goto L3
	} else {
		goto L15
	}
L15:
	;
	return v45
L16:
	;
	return v50
L17:
	;
	return v237
L18:
	;
	v230 = F_fpconv_dtoa(m, l2, l0)
	mBase = m.M
	v232 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l0+v230))) = uint8(v232)
	v237 = v230
	goto L17
L19:
	;
	if base.F64_lt(v8, float64(9.223372036854776e+18)) == int32(0) {
		goto L21
	} else {
		goto L22
	}
L20:
	;
	if base.F64_ne(l2, base.F64_convert_i64_s(v61)) != 0 {
		goto L18
	} else {
		goto L23
	}
L21:
	;
	v61 = int64(-9223372036854775807 - 1)
	goto L20
L22:
	;
	v59 = base.I64_trunc_f64_s(l2)
	v61 = v59
	goto L20
L23:
	;
	v64 = int32(0)
	if v61 <= int64(-1) {
		goto L26
	} else {
		goto L27
	}
L24:
	;
	if l1 != int32(1) {
		v237 = v64
		goto L17
	} else {
		goto L74
	}
L25:
	;
	v83 = int32(0)
	v87 = int32(1)
	if base.Ui64(v81) < base.Ui64(int64(10)) {
		v144 = v87
		v145 = v83
		goto L30
	} else {
		goto L31
	}
L26:
	;
	if base.Ui32(l1) < base.Ui32(int32(2)) {
		goto L24
	} else {
		goto L28
	}
L27:
	;
	v79 = l0
	v80 = l1
	v81 = v61
	v82 = int32(0)
	goto L25
L28:
	;
	v70 = int32(45)
	*(*uint8)(unsafe.Add(mBase, uint32(l0))) = uint8(v70)
	v74 = int32(1)
	v79 = l0 + v74
	v80 = l1 + int32(-1)
	v81 = int64(0) - v61
	v82 = v74
	goto L25
L29:
	;
	if v218 == int32(0) {
		v237 = v64
		goto L17
	} else {
		goto L73
	}
L30:
	;
	v148 = v144 + v145
	if base.Ui32(v80) <= base.Ui32(v148) {
		goto L61
	} else {
		goto L62
	}
L31:
	;
	v95 = v83
	v96 = v81
	goto L32
L32:
	;
	if base.Ui64(int64(100)) <= base.Ui64(v96) {
		goto L34
	} else {
		goto L35
	}
L33:
	;
	v144 = v87
	v145 = v136
	goto L30
L34:
	;
	if base.Ui64(int64(1000)) <= base.Ui64(v96) {
		goto L36
	} else {
		goto L37
	}
L35:
	;
	v144 = int32(2)
	v145 = v95
	goto L30
L36:
	;
	if base.Ui64(int64(999999999999)) < base.Ui64(v96) {
		goto L38
	} else {
		goto L39
	}
L37:
	;
	v144 = int32(3)
	v145 = v95
	goto L30
L38:
	;
	v136 = v95 + int32(12)
	v140 = base.I64_div_u_s(v96, int64(1000000000000))
	if base.Ui64(int64(9999999999999)) < base.Ui64(v96) {
		v95 = v136
		v96 = v140
		goto L32
	} else {
		goto L60
	}
L39:
	;
	if base.Ui64(int64(99999999)) < base.Ui64(v96) {
		goto L40
	} else {
		goto L41
	}
L40:
	;
	if base.Ui64(int64(9999999999)) < base.Ui64(v96) {
		goto L52
	} else {
		goto L53
	}
L41:
	;
	if base.Ui64(int64(999999)) < base.Ui64(v96) {
		goto L42
	} else {
		goto L43
	}
L42:
	;
	if base.Ui64(int64(9999999)) < base.Ui64(v96) {
		goto L49
	} else {
		goto L50
	}
L43:
	;
	if base.Ui64(int64(10000)) <= base.Ui64(v96) {
		goto L44
	} else {
		goto L45
	}
L44:
	;
	if base.Ui64(int64(99999)) < base.Ui64(v96) {
		goto L46
	} else {
		goto L47
	}
L45:
	;
	v144 = int32(4)
	v145 = v95
	goto L30
L46:
	;
	v117 = int32(6)
	goto L48
L47:
	;
	v117 = int32(5)
	goto L48
L48:
	;
	v144 = v117
	v145 = v95
	goto L30
L49:
	;
	v122 = int32(8)
	goto L51
L50:
	;
	v122 = int32(7)
	goto L51
L51:
	;
	v144 = v122
	v145 = v95
	goto L30
L52:
	;
	if base.Ui64(int64(99999999999)) < base.Ui64(v96) {
		goto L57
	} else {
		goto L58
	}
L53:
	;
	if base.Ui64(int64(999999999)) < base.Ui64(v96) {
		goto L54
	} else {
		goto L55
	}
L54:
	;
	v129 = int32(10)
	goto L56
L55:
	;
	v129 = int32(9)
	goto L56
L56:
	;
	v144 = v129
	v145 = v95
	goto L30
L57:
	;
	v134 = int32(12)
	goto L59
L58:
	;
	v134 = int32(11)
	goto L59
L59:
	;
	v144 = v134
	v145 = v95
	goto L30
L60:
	;
	goto L33
L61:
	;
	if v80 == int32(0) {
		goto L71
	} else {
		goto L72
	}
L62:
	;
	v151 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v79+v148))) = uint8(v151)
	v154 = v148 + int32(-1)
	if base.Ui64(int64(100)) <= base.Ui64(v81) {
		goto L64
	} else {
		goto L65
	}
L63:
	;
	v190 = v79 + v187
	if base.Ui64(int64(9)) < base.Ui64(v188) {
		goto L69
	} else {
		goto L70
	}
L64:
	;
	v161 = v81
	v163 = v154
	goto L66
L65:
	;
	v187 = v154
	v188 = v81
	goto L63
L66:
	;
	v167 = int64(100)
	v168 = base.I64_div_u_s(v161, v167)
	v177 = int32(*(*uint16)(unsafe.Add(mBase, uint32(base.I32_wrap_i64(v161-v168*v167)<<(uint(int32(1))%32))+uint32(_c_F_d2string[0]))))
	*(*uint16)(unsafe.Add(mBase, uint32(v79+int32(-1)+v163))) = uint16(v177)
	v180 = v163 + int32(-2)
	if base.Ui64(int64(9999)) < base.Ui64(v161) {
		v161 = v168
		v163 = v180
		goto L66
	} else {
		goto L68
	}
L67:
	;
	v187 = v180
	v188 = v168
	goto L63
L68:
	;
	goto L67
L69:
	;
	v204 = int32(*(*uint16)(unsafe.Add(mBase, uint32(base.I32_wrap_i64(v188)<<(uint(int32(1))%32))+uint32(_c_F_d2string[0]))))
	*(*uint16)(unsafe.Add(mBase, uint32(v190+int32(-1)))) = uint16(v204)
	v218 = v148
	goto L29
L70:
	;
	v195 = base.I32_wrap_i64(v188) | int32(48)
	*(*uint8)(unsafe.Add(mBase, uint32(v190))) = uint8(v195)
	v218 = v148
	goto L29
L71:
	;
	v218 = int32(0)
	goto L29
L72:
	;
	v208 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v79))) = uint8(v208)
	goto L71
L73:
	;
	return v218 + v82
L74:
	;
	v225 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l0))) = uint8(v225)
	return v225
}
func F_dbsHaveNoKeys(m *base.Module) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v27 int64
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v34 int64
	_ = v34
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
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
	v5 = int32(1)
	v7 = *(*int32)(unsafe.Add(mBase, _c_F_dbsHaveNoKeys[0]))
	if v7 < v5 {
		v49 = v5
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return v49
L2:
	;
	v10 = int32(0)
	v11 = *(*int32)(unsafe.Add(mBase, _c_F_dbsHaveNoKeys[1]))
	v14 = v7
	v15 = v11
	v16 = v10
	goto L3
L3:
	;
	v20 = *(*int32)(unsafe.Add(mBase, uint32(v15+v16<<(uint(int32(2))%32))))
	if v20 == int32(0) {
		v43 = v14
		v44 = v15
		goto L5
	} else {
		goto L6
	}
L4:
	;
	v49 = v45
	goto L1
L5:
	;
	v45 = int32(1)
	v47 = v16 + v45
	if v47 < v43 {
		v14 = v43
		v15 = v44
		v16 = v47
		goto L3
	} else {
		goto L14
	}
L6:
	;
	v23 = *(*int32)(unsafe.Add(mBase, uint32(v20)))
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v23)+12))
	if v24 == int32(1) {
		goto L9
	} else {
		goto L10
	}
L7:
	;
	v39 = int32(0)
	v40 = *(*int32)(unsafe.Add(mBase, _c_F_dbsHaveNoKeys[0]))
	v42 = *(*int32)(unsafe.Add(mBase, _c_F_dbsHaveNoKeys[1]))
	v43 = v40
	v44 = v42
	goto L5
L8:
	;
	if v34 == int64(0) {
		goto L7
	} else {
		goto L13
	}
L9:
	;
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v23)+8))
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v28)))
	if v29 != 0 {
		goto L11
	} else {
		goto L12
	}
L10:
	;
	v27 = *(*int64)(unsafe.Add(mBase, uint32(v23)+40))
	v34 = v27
	goto L8
L11:
	;
	v31 = F_hashtableSize(m, v29)
	mBase = m.M
	v34 = base.I64_extend_i32_u(v31)
	goto L8
L12:
	;
	v34 = int64(0)
	goto L8
L13:
	;
	return int32(0)
L14:
	;
	goto L4
}
func F_debugPauseProcess(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v21 int32
	_ = v21
	v2 = *(*int32)(unsafe.Add(mBase, _c_F_debugPauseProcess[0]))
	if int32(2) < v2 {
		v11 = F_raise(m, int32(19))
		mBase = m.M
		v12 = m.ExcPending
		if v12 != 0 {
			return
		} else {
			v14 = *(*int32)(unsafe.Add(mBase, _c_F_debugPauseProcess[0]))
			if int32(2) < v14 {
				return
			} else {
				F__serverLog(m, int32(2), int32(_a_F_debugPauseProcess_0), int32(0))
				mBase = m.M
				v21 = m.ExcPending
				if v21 != 0 {
					return
				} else {
					return
				}
			}
		}
	} else {
		F__serverLog(m, int32(2), int32(_a_F_debugPauseProcess_1), int32(0))
		mBase = m.M
		v9 = m.ExcPending
		if v9 != 0 {
			return
		} else {
			v11 = F_raise(m, int32(19))
			mBase = m.M
			v12 = m.ExcPending
			if v12 != 0 {
				return
			} else {
				v14 = *(*int32)(unsafe.Add(mBase, _c_F_debugPauseProcess[0]))
				if int32(2) < v14 {
					return
				} else {
					F__serverLog(m, int32(2), int32(_a_F_debugPauseProcess_0), int32(0))
					mBase = m.M
					v21 = m.ExcPending
					if v21 != 0 {
						return
					} else {
						return
					}
				}
			}
		}
	}
}
func F_debugScriptRespToHuman(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v69 int32
	_ = v69
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v100 int64
	_ = v100
	var v106 int32
	_ = v106
	var v108 int64
	_ = v108
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
	var v122 int64
	_ = v122
	var v127 int64
	_ = v127
	var v131 int32
	_ = v131
	var v133 int64
	_ = v133
	var v135 int32
	_ = v135
	var v143 int64
	_ = v143
	var v167 int64
	_ = v167
	var v193 int32
	_ = v193
	var v194 int64
	_ = v194
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v202 int32
	_ = v202
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v213 int32
	_ = v213
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v219 int32
	_ = v219
	var v223 int32
	_ = v223
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v230 int32
	_ = v230
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
	var v236 int32
	_ = v236
	var v240 int32
	_ = v240
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v248 int32
	_ = v248
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v252 int32
	_ = v252
	var v256 int32
	_ = v256
	var v259 int32
	_ = v259
	var v261 int32
	_ = v261
	var v274 int32
	_ = v274
	var v275 int32
	_ = v275
	var v279 int32
	_ = v279
	var v295 int32
	_ = v295
	var v297 int32
	_ = v297
	var v298 int32
	_ = v298
	var v299 int32
	_ = v299
	var v310 int64
	_ = v310
	var v316 int32
	_ = v316
	var v318 int64
	_ = v318
	var v320 int32
	_ = v320
	var v322 int32
	_ = v322
	var v332 int64
	_ = v332
	var v337 int64
	_ = v337
	var v341 int32
	_ = v341
	var v343 int64
	_ = v343
	var v345 int32
	_ = v345
	var v353 int64
	_ = v353
	var v377 int64
	_ = v377
	var v404 int32
	_ = v404
	var v405 int32
	_ = v405
	var v406 int64
	_ = v406
	var v411 int32
	_ = v411
	var v412 int32
	_ = v412
	var v416 int32
	_ = v416
	var v417 int32
	_ = v417
	var v419 int64
	_ = v419
	var v424 int32
	_ = v424
	var v428 int64
	_ = v428
	var v431 int32
	_ = v431
	var v432 int32
	_ = v432
	var v433 int64
	_ = v433
	var v437 int32
	_ = v437
	var v440 int32
	_ = v440
	var v441 int32
	_ = v441
	var v443 int64
	_ = v443
	var v444 int64
	_ = v444
	var v446 int64
	_ = v446
	var v448 int32
	_ = v448
	var v450 int32
	_ = v450
	var v452 int32
	_ = v452
	var v459 int32
	_ = v459
	var v460 int32
	_ = v460
	var v463 int32
	_ = v463
	var v464 int32
	_ = v464
	var v465 int32
	_ = v465
	var v467 int32
	_ = v467
	var v471 int32
	_ = v471
	var v474 int32
	_ = v474
	var v476 int32
	_ = v476
	var v489 int32
	_ = v489
	var v490 int32
	_ = v490
	var v494 int32
	_ = v494
	var v510 int32
	_ = v510
	var v512 int32
	_ = v512
	var v513 int32
	_ = v513
	var v514 int32
	_ = v514
	var v525 int64
	_ = v525
	var v531 int32
	_ = v531
	var v533 int64
	_ = v533
	var v535 int32
	_ = v535
	var v537 int32
	_ = v537
	var v547 int64
	_ = v547
	var v552 int64
	_ = v552
	var v556 int32
	_ = v556
	var v558 int64
	_ = v558
	var v560 int32
	_ = v560
	var v568 int64
	_ = v568
	var v592 int64
	_ = v592
	var v618 int32
	_ = v618
	var v621 int32
	_ = v621
	var v622 int32
	_ = v622
	var v625 int32
	_ = v625
	var v626 int64
	_ = v626
	var v631 int32
	_ = v631
	var v635 int64
	_ = v635
	var v638 int32
	_ = v638
	var v639 int32
	_ = v639
	var v640 int64
	_ = v640
	var v644 int32
	_ = v644
	var v647 int32
	_ = v647
	var v648 int32
	_ = v648
	var v650 int64
	_ = v650
	var v651 int64
	_ = v651
	var v653 int64
	_ = v653
	var v655 int32
	_ = v655
	var v657 int32
	_ = v657
	var v659 int32
	_ = v659
	var v666 int32
	_ = v666
	var v667 int32
	_ = v667
	var v670 int32
	_ = v670
	var v671 int32
	_ = v671
	var v672 int32
	_ = v672
	var v674 int32
	_ = v674
	var v678 int32
	_ = v678
	var v681 int32
	_ = v681
	var v683 int32
	_ = v683
	var v696 int32
	_ = v696
	var v697 int32
	_ = v697
	var v701 int32
	_ = v701
	var v717 int32
	_ = v717
	var v719 int32
	_ = v719
	var v720 int32
	_ = v720
	var v721 int32
	_ = v721
	var v732 int64
	_ = v732
	var v738 int32
	_ = v738
	var v740 int64
	_ = v740
	var v742 int32
	_ = v742
	var v744 int32
	_ = v744
	var v754 int64
	_ = v754
	var v759 int64
	_ = v759
	var v763 int32
	_ = v763
	var v765 int64
	_ = v765
	var v767 int32
	_ = v767
	var v775 int64
	_ = v775
	var v799 int64
	_ = v799
	var v825 int32
	_ = v825
	var v828 int32
	_ = v828
	var v829 int32
	_ = v829
	var v832 int32
	_ = v832
	var v833 int64
	_ = v833
	var v838 int32
	_ = v838
	var v842 int64
	_ = v842
	var v845 int32
	_ = v845
	var v846 int32
	_ = v846
	var v847 int32
	_ = v847
	var v850 int32
	_ = v850
	var v851 int32
	_ = v851
	var v853 int32
	_ = v853
	var v854 int32
	_ = v854
	var v855 int64
	_ = v855
	var v859 int32
	_ = v859
	var v862 int32
	_ = v862
	var v863 int32
	_ = v863
	var v865 int64
	_ = v865
	var v866 int64
	_ = v866
	var v868 int64
	_ = v868
	var v870 int32
	_ = v870
	var v872 int32
	_ = v872
	var v874 int32
	_ = v874
	var v881 int32
	_ = v881
	var v882 int32
	_ = v882
	var v886 int32
	_ = v886
	var v887 int32
	_ = v887
	var v889 int32
	_ = v889
	var v893 int32
	_ = v893
	var v894 int32
	_ = v894
	var v897 int32
	_ = v897
	var v898 int32
	_ = v898
	var v904 int32
	_ = v904
	var v905 int32
	_ = v905
	var v907 int32
	_ = v907
	var v911 int32
	_ = v911
	var v912 int32
	_ = v912
	var v913 int32
	_ = v913
	var v918 int32
	_ = v918
	var v919 int32
	_ = v919
	var v922 int32
	_ = v922
	var v923 int32
	_ = v923
	var v924 int32
	_ = v924
	var v929 int32
	_ = v929
	var v930 int32
	_ = v930
	var v931 int32
	_ = v931
	var v933 int32
	_ = v933
	var v937 int32
	_ = v937
	var v938 int32
	_ = v938
	var v941 int32
	_ = v941
	var v942 int32
	_ = v942
	var v947 int32
	_ = v947
	var v948 int32
	_ = v948
	var v953 int32
	_ = v953
	var v964 int32
	_ = v964
	var v965 int32
	_ = v965
	var v966 int32
	_ = v966
	var v967 int32
	_ = v967
	v9 = m.G0
	v11 = v9 - int32(16)
	m.G0 = v11
	v13 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	switch v13 + int32(-35) {
	case 0:
		goto L3
	case 1:
		goto L10
	case 2:
		goto L5
	default:
		v953 = l1
		goto L1
	case 7:
		goto L7
	case 8:
		goto L9
	case 9:
		goto L2
	case 10:
		goto L8
	case 23:
		goto L11
	case 60:
		goto L4
	case 91:
		goto L6
	}
L1:
	;
	m.G0 = v11 + int32(16)
	return v953
L2:
	;
	v929 = l1 + int32(1)
	v930 = int32(13)
	v931 = F___strchrnul(m, v929, v930)
	mBase = m.M
	v933 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v931))))
	if v933 == v930 {
		goto L210
	} else {
		goto L211
	}
L3:
	;
	v904 = int32(13)
	v905 = F___strchrnul(m, l1+int32(1), v904)
	mBase = m.M
	v907 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v905))))
	if v907 == v904 {
		goto L201
	} else {
		goto L202
	}
L4:
	;
	v886 = int32(13)
	v887 = F___strchrnul(m, l1+int32(1), v886)
	mBase = m.M
	v889 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v887))))
	if v889 == v886 {
		goto L196
	} else {
		goto L197
	}
L5:
	;
	v670 = l1 + int32(1)
	v671 = int32(13)
	v672 = F___strchrnul(m, v670, v671)
	mBase = m.M
	v674 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v672))))
	if v674 == v671 {
		goto L152
	} else {
		goto L153
	}
L6:
	;
	v463 = l1 + int32(1)
	v464 = int32(13)
	v465 = F___strchrnul(m, v463, v464)
	mBase = m.M
	v467 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v465))))
	if v467 == v464 {
		goto L110
	} else {
		goto L111
	}
L7:
	;
	v248 = l1 + int32(1)
	v249 = int32(13)
	v250 = F___strchrnul(m, v248, v249)
	mBase = m.M
	v252 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v250))))
	if v252 == v249 {
		goto L65
	} else {
		goto L66
	}
L8:
	;
	v230 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v233 = int32(13)
	v234 = F___strchrnul(m, l1+int32(1), v233)
	mBase = m.M
	v236 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v234))))
	if v236 == v233 {
		goto L60
	} else {
		goto L61
	}
L9:
	;
	v213 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v216 = int32(13)
	v217 = F___strchrnul(m, l1+int32(1), v216)
	mBase = m.M
	v219 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v217))))
	if v219 == v216 {
		goto L55
	} else {
		goto L56
	}
L10:
	;
	v38 = l1 + int32(1)
	v39 = int32(13)
	v40 = F___strchrnul(m, v38, v39)
	mBase = m.M
	v42 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v40))))
	if v42 == v39 {
		goto L19
	} else {
		goto L20
	}
L11:
	;
	v17 = l1 + int32(1)
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v19 = int32(13)
	v20 = F___strchrnul(m, v17, v19)
	mBase = m.M
	v22 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20))))
	if v22 == v19 {
		goto L13
	} else {
		goto L14
	}
L12:
	;
	v30 = F_sdscatlen(m, v18, v17, v26+(l1^int32(-1)))
	mBase = m.M
	v33 = m.ExcPending
	if v33 != 0 {
		goto L16
	} else {
		goto L17
	}
L13:
	;
	v26 = v20
	goto L15
L14:
	;
	v26 = int32(0)
	goto L15
L15:
	;
	goto L12
L16:
	;
	return int32(0)
L17:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v30
	v953 = v26 + int32(2)
	goto L1
L18:
	;
	v49 = v46 + (l1 ^ int32(-1))
	v51 = v11 + int32(8)
	if base.Ui32(v49+int32(-21)) < base.Ui32(int32(-20)) {
		goto L23
	} else {
		goto L24
	}
L19:
	;
	v46 = v40
	goto L21
L20:
	;
	v46 = int32(0)
	goto L21
L21:
	;
	goto L18
L22:
	;
	v193 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v194 = *(*int64)(unsafe.Add(mBase, uint32(v11)+8))
	if v194 != int64(-1) {
		goto L50
	} else {
		goto L51
	}
L23:
	;
	goto L22
L24:
	;
	v64 = int32(1)
	v65 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v38))))
	if v49 != v64 {
		goto L27
	} else {
		goto L28
	}
L25:
	;
	goto L23
L26:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v51))) = v167
	goto L25
L27:
	;
	if v65&int32(255) == int32(45) {
		goto L32
	} else {
		goto L33
	}
L28:
	;
	v69 = v65 + int32(-48)
	if base.Ui32(int32(9)) < base.Ui32(v69&int32(255)) {
		goto L23
	} else {
		goto L29
	}
L29:
	;
	if v51 == int32(0) {
		goto L25
	} else {
		goto L30
	}
L30:
	;
	v167 = base.I64_extend_i32_u(v69) & int64(255)
	goto L26
L31:
	;
	if base.Ui32(int32(8)) < base.Ui32((v88+int32(-49))&int32(255)) {
		goto L23
	} else {
		goto L34
	}
L32:
	;
	v964 = int32(2)
	v85 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v38)+1)))
	v87 = v964
	v88 = v85
	v89 = l1 + v964
	goto L31
L33:
	;
	v87 = v64
	v88 = v65
	v89 = v38
	goto L31
L34:
	;
	v100 = base.I64_extend_i32_u(v88+int32(-48)) & int64(255)
	if base.Ui32(v49) <= base.Ui32(v87) {
		v143 = v100
		goto L35
	} else {
		goto L36
	}
L35:
	;
	if v65&int32(255) != int32(45) {
		goto L43
	} else {
		goto L44
	}
L36:
	;
	v106 = v87
	v108 = v100
	v110 = v89
	goto L37
L37:
	;
	v112 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v110)+1)))
	if base.Ui32((v112+int32(-58))&int32(255)) < base.Ui32(int32(246)) {
		goto L23
	} else {
		goto L39
	}
L38:
	;
	v143 = v133
	goto L35
L39:
	;
	if base.Ui64(int64(1844674407370955161)) < base.Ui64(v108) {
		goto L23
	} else {
		goto L40
	}
L40:
	;
	v122 = v108 * int64(10)
	v127 = base.I64_extend_i32_u(v112+int32(-48)) & int64(255)
	if base.Ui64(v127^int64(-1)) < base.Ui64(v122) {
		goto L23
	} else {
		goto L41
	}
L41:
	;
	v131 = int32(1)
	v133 = v122 + v127
	v135 = v106 + v131
	if v135 != v49 {
		v106 = v135
		v108 = v133
		v110 = v110 + v131
		goto L37
	} else {
		goto L42
	}
L42:
	;
	goto L38
L43:
	;
	if v143 < int64(0) {
		goto L23
	} else {
		goto L47
	}
L44:
	;
	if base.Ui64(int64(-9223372036854775807-1)) < base.Ui64(v143) {
		goto L23
	} else {
		goto L45
	}
L45:
	;
	if v51 == int32(0) {
		goto L25
	} else {
		goto L46
	}
L46:
	;
	v167 = int64(0) - v143
	goto L26
L47:
	;
	if v51 == int32(0) {
		goto L25
	} else {
		goto L48
	}
L48:
	;
	v167 = v143
	goto L26
L49:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v208
	v953 = v209 + int32(2)
	goto L1
L50:
	;
	v202 = v46 + int32(2)
	v204 = F_sdscatrepr(m, v193, v202, base.I32_wrap_i64(v194))
	mBase = m.M
	v205 = m.ExcPending
	if v205 != 0 {
		goto L16
	} else {
		goto L53
	}
L51:
	;
	v199 = F_sdscatlen(m, v193, int32(_a_F_debugScriptRespToHuman_0), int32(4))
	mBase = m.M
	v200 = m.ExcPending
	if v200 != 0 {
		goto L16
	} else {
		goto L52
	}
L52:
	;
	v208 = v199
	v209 = v46
	goto L49
L53:
	;
	v206 = *(*int32)(unsafe.Add(mBase, uint32(v11)+8))
	v208 = v204
	v209 = v202 + v206
	goto L49
L54:
	;
	v225 = F_sdscatrepr(m, v213, l1, v223-l1)
	mBase = m.M
	v226 = m.ExcPending
	if v226 != 0 {
		goto L16
	} else {
		goto L58
	}
L55:
	;
	v223 = v217
	goto L57
L56:
	;
	v223 = int32(0)
	goto L57
L57:
	;
	goto L54
L58:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v225
	v953 = v223 + int32(2)
	goto L1
L59:
	;
	v242 = F_sdscatrepr(m, v230, l1, v240-l1)
	mBase = m.M
	v243 = m.ExcPending
	if v243 != 0 {
		goto L16
	} else {
		goto L63
	}
L60:
	;
	v240 = v234
	goto L62
L61:
	;
	v240 = int32(0)
	goto L62
L62:
	;
	goto L59
L63:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v242
	v953 = v240 + int32(2)
	goto L1
L64:
	;
	v259 = v256 + (l1 ^ int32(-1))
	v261 = v11 + int32(8)
	if base.Ui32(v259+int32(-21)) < base.Ui32(int32(-20)) {
		goto L69
	} else {
		goto L70
	}
L65:
	;
	v256 = v250
	goto L67
L66:
	;
	v256 = int32(0)
	goto L67
L67:
	;
	goto L64
L68:
	;
	v404 = v256 + int32(2)
	v405 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v406 = *(*int64)(unsafe.Add(mBase, uint32(v11)+8))
	if v406 != int64(-1) {
		goto L95
	} else {
		goto L96
	}
L69:
	;
	goto L68
L70:
	;
	v274 = int32(1)
	v275 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v248))))
	if v259 != v274 {
		goto L73
	} else {
		goto L74
	}
L71:
	;
	goto L69
L72:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v261))) = v377
	goto L71
L73:
	;
	if v275&int32(255) == int32(45) {
		goto L78
	} else {
		goto L79
	}
L74:
	;
	v279 = v275 + int32(-48)
	if base.Ui32(int32(9)) < base.Ui32(v279&int32(255)) {
		goto L69
	} else {
		goto L75
	}
L75:
	;
	if v261 == int32(0) {
		goto L71
	} else {
		goto L76
	}
L76:
	;
	v377 = base.I64_extend_i32_u(v279) & int64(255)
	goto L72
L77:
	;
	if base.Ui32(int32(8)) < base.Ui32((v298+int32(-49))&int32(255)) {
		goto L69
	} else {
		goto L80
	}
L78:
	;
	v965 = int32(2)
	v295 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v248)+1)))
	v297 = v965
	v298 = v295
	v299 = l1 + v965
	goto L77
L79:
	;
	v297 = v274
	v298 = v275
	v299 = v248
	goto L77
L80:
	;
	v310 = base.I64_extend_i32_u(v298+int32(-48)) & int64(255)
	if base.Ui32(v259) <= base.Ui32(v297) {
		v353 = v310
		goto L81
	} else {
		goto L82
	}
L81:
	;
	if v275&int32(255) != int32(45) {
		goto L89
	} else {
		goto L90
	}
L82:
	;
	v316 = v297
	v318 = v310
	v320 = v299
	goto L83
L83:
	;
	v322 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v320)+1)))
	if base.Ui32((v322+int32(-58))&int32(255)) < base.Ui32(int32(246)) {
		goto L69
	} else {
		goto L85
	}
L84:
	;
	v353 = v343
	goto L81
L85:
	;
	if base.Ui64(int64(1844674407370955161)) < base.Ui64(v318) {
		goto L69
	} else {
		goto L86
	}
L86:
	;
	v332 = v318 * int64(10)
	v337 = base.I64_extend_i32_u(v322+int32(-48)) & int64(255)
	if base.Ui64(v337^int64(-1)) < base.Ui64(v332) {
		goto L69
	} else {
		goto L87
	}
L87:
	;
	v341 = int32(1)
	v343 = v332 + v337
	v345 = v316 + v341
	if v345 != v259 {
		v316 = v345
		v318 = v343
		v320 = v320 + v341
		goto L83
	} else {
		goto L88
	}
L88:
	;
	goto L84
L89:
	;
	if v353 < int64(0) {
		goto L69
	} else {
		goto L93
	}
L90:
	;
	if base.Ui64(int64(-9223372036854775807-1)) < base.Ui64(v353) {
		goto L69
	} else {
		goto L91
	}
L91:
	;
	if v261 == int32(0) {
		goto L71
	} else {
		goto L92
	}
L92:
	;
	v377 = int64(0) - v353
	goto L72
L93:
	;
	if v261 == int32(0) {
		goto L71
	} else {
		goto L94
	}
L94:
	;
	v377 = v353
	goto L72
L95:
	;
	v416 = F_sdscatlen(m, v405, int32(_a_F_debugScriptRespToHuman_1), int32(1))
	mBase = m.M
	v417 = m.ExcPending
	if v417 != 0 {
		goto L16
	} else {
		goto L98
	}
L96:
	;
	v411 = F_sdscatlen(m, v405, int32(_a_F_debugScriptRespToHuman_0), int32(4))
	mBase = m.M
	v412 = m.ExcPending
	if v412 != 0 {
		goto L16
	} else {
		goto L97
	}
L97:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v411
	v953 = v404
	goto L1
L98:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v416
	v419 = *(*int64)(unsafe.Add(mBase, uint32(v11)+8))
	if v419 < int64(1) {
		v450 = v404
		v452 = v416
		goto L99
	} else {
		goto L100
	}
L99:
	;
	v459 = F_sdscatlen(m, v452, int32(_a_F_debugScriptRespToHuman_2), int32(1))
	mBase = m.M
	v460 = m.ExcPending
	if v460 != 0 {
		goto L16
	} else {
		goto L108
	}
L100:
	;
	v424 = v404
	v428 = int64(0)
	goto L101
L101:
	;
	v431 = F_debugScriptRespToHuman(m, l0, v424)
	mBase = m.M
	v432 = m.ExcPending
	if v432 != 0 {
		goto L16
	} else {
		goto L103
	}
L102:
	;
	v448 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v450 = v431
	v452 = v448
	goto L99
L103:
	;
	v433 = *(*int64)(unsafe.Add(mBase, uint32(v11)+8))
	if v433+int64(-1) == v428 {
		v444 = v433
		goto L104
	} else {
		goto L105
	}
L104:
	;
	v446 = v428 + int64(1)
	if v446 < v444 {
		v424 = v431
		v428 = v446
		goto L101
	} else {
		goto L107
	}
L105:
	;
	v437 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v440 = F_sdscatlen(m, v437, int32(_a_F_debugScriptRespToHuman_3), int32(1))
	mBase = m.M
	v441 = m.ExcPending
	if v441 != 0 {
		goto L16
	} else {
		goto L106
	}
L106:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v440
	v443 = *(*int64)(unsafe.Add(mBase, uint32(v11)+8))
	v444 = v443
	goto L104
L107:
	;
	goto L102
L108:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v459
	v953 = v450
	goto L1
L109:
	;
	v474 = v471 + (l1 ^ int32(-1))
	v476 = v11 + int32(8)
	if base.Ui32(v474+int32(-21)) < base.Ui32(int32(-20)) {
		goto L114
	} else {
		goto L115
	}
L110:
	;
	v471 = v465
	goto L112
L111:
	;
	v471 = int32(0)
	goto L112
L112:
	;
	goto L109
L113:
	;
	v618 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v621 = F_sdscatlen(m, v618, int32(_a_F_debugScriptRespToHuman_4), int32(2))
	mBase = m.M
	v622 = m.ExcPending
	if v622 != 0 {
		goto L16
	} else {
		goto L140
	}
L114:
	;
	goto L113
L115:
	;
	v489 = int32(1)
	v490 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v463))))
	if v474 != v489 {
		goto L118
	} else {
		goto L119
	}
L116:
	;
	goto L114
L117:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v476))) = v592
	goto L116
L118:
	;
	if v490&int32(255) == int32(45) {
		goto L123
	} else {
		goto L124
	}
L119:
	;
	v494 = v490 + int32(-48)
	if base.Ui32(int32(9)) < base.Ui32(v494&int32(255)) {
		goto L114
	} else {
		goto L120
	}
L120:
	;
	if v476 == int32(0) {
		goto L116
	} else {
		goto L121
	}
L121:
	;
	v592 = base.I64_extend_i32_u(v494) & int64(255)
	goto L117
L122:
	;
	if base.Ui32(int32(8)) < base.Ui32((v513+int32(-49))&int32(255)) {
		goto L114
	} else {
		goto L125
	}
L123:
	;
	v966 = int32(2)
	v510 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v463)+1)))
	v512 = v966
	v513 = v510
	v514 = l1 + v966
	goto L122
L124:
	;
	v512 = v489
	v513 = v490
	v514 = v463
	goto L122
L125:
	;
	v525 = base.I64_extend_i32_u(v513+int32(-48)) & int64(255)
	if base.Ui32(v474) <= base.Ui32(v512) {
		v568 = v525
		goto L126
	} else {
		goto L127
	}
L126:
	;
	if v490&int32(255) != int32(45) {
		goto L134
	} else {
		goto L135
	}
L127:
	;
	v531 = v512
	v533 = v525
	v535 = v514
	goto L128
L128:
	;
	v537 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v535)+1)))
	if base.Ui32((v537+int32(-58))&int32(255)) < base.Ui32(int32(246)) {
		goto L114
	} else {
		goto L130
	}
L129:
	;
	v568 = v558
	goto L126
L130:
	;
	if base.Ui64(int64(1844674407370955161)) < base.Ui64(v533) {
		goto L114
	} else {
		goto L131
	}
L131:
	;
	v547 = v533 * int64(10)
	v552 = base.I64_extend_i32_u(v537+int32(-48)) & int64(255)
	if base.Ui64(v552^int64(-1)) < base.Ui64(v547) {
		goto L114
	} else {
		goto L132
	}
L132:
	;
	v556 = int32(1)
	v558 = v547 + v552
	v560 = v531 + v556
	if v560 != v474 {
		v531 = v560
		v533 = v558
		v535 = v535 + v556
		goto L128
	} else {
		goto L133
	}
L133:
	;
	goto L129
L134:
	;
	if v568 < int64(0) {
		goto L114
	} else {
		goto L138
	}
L135:
	;
	if base.Ui64(int64(-9223372036854775807-1)) < base.Ui64(v568) {
		goto L114
	} else {
		goto L136
	}
L136:
	;
	if v476 == int32(0) {
		goto L116
	} else {
		goto L137
	}
L137:
	;
	v592 = int64(0) - v568
	goto L117
L138:
	;
	if v476 == int32(0) {
		goto L116
	} else {
		goto L139
	}
L139:
	;
	v592 = v568
	goto L117
L140:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v621
	v625 = v471 + int32(2)
	v626 = *(*int64)(unsafe.Add(mBase, uint32(v11)+8))
	if v626 < int64(1) {
		v657 = v625
		v659 = v621
		goto L141
	} else {
		goto L142
	}
L141:
	;
	v666 = F_sdscatlen(m, v659, int32(_a_F_debugScriptRespToHuman_5), int32(1))
	mBase = m.M
	v667 = m.ExcPending
	if v667 != 0 {
		goto L16
	} else {
		goto L150
	}
L142:
	;
	v631 = v625
	v635 = int64(0)
	goto L143
L143:
	;
	v638 = F_debugScriptRespToHuman(m, l0, v631)
	mBase = m.M
	v639 = m.ExcPending
	if v639 != 0 {
		goto L16
	} else {
		goto L145
	}
L144:
	;
	v655 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v657 = v638
	v659 = v655
	goto L141
L145:
	;
	v640 = *(*int64)(unsafe.Add(mBase, uint32(v11)+8))
	if v640+int64(-1) == v635 {
		v651 = v640
		goto L146
	} else {
		goto L147
	}
L146:
	;
	v653 = v635 + int64(1)
	if v653 < v651 {
		v631 = v638
		v635 = v653
		goto L143
	} else {
		goto L149
	}
L147:
	;
	v644 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v647 = F_sdscatlen(m, v644, int32(_a_F_debugScriptRespToHuman_3), int32(1))
	mBase = m.M
	v648 = m.ExcPending
	if v648 != 0 {
		goto L16
	} else {
		goto L148
	}
L148:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v647
	v650 = *(*int64)(unsafe.Add(mBase, uint32(v11)+8))
	v651 = v650
	goto L146
L149:
	;
	goto L144
L150:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v666
	v953 = v657
	goto L1
L151:
	;
	v681 = v678 + (l1 ^ int32(-1))
	v683 = v11 + int32(8)
	if base.Ui32(v681+int32(-21)) < base.Ui32(int32(-20)) {
		goto L156
	} else {
		goto L157
	}
L152:
	;
	v678 = v672
	goto L154
L153:
	;
	v678 = int32(0)
	goto L154
L154:
	;
	goto L151
L155:
	;
	v825 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v828 = F_sdscatlen(m, v825, int32(_a_F_debugScriptRespToHuman_6), int32(1))
	mBase = m.M
	v829 = m.ExcPending
	if v829 != 0 {
		goto L16
	} else {
		goto L182
	}
L156:
	;
	goto L155
L157:
	;
	v696 = int32(1)
	v697 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v670))))
	if v681 != v696 {
		goto L160
	} else {
		goto L161
	}
L158:
	;
	goto L156
L159:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v683))) = v799
	goto L158
L160:
	;
	if v697&int32(255) == int32(45) {
		goto L165
	} else {
		goto L166
	}
L161:
	;
	v701 = v697 + int32(-48)
	if base.Ui32(int32(9)) < base.Ui32(v701&int32(255)) {
		goto L156
	} else {
		goto L162
	}
L162:
	;
	if v683 == int32(0) {
		goto L158
	} else {
		goto L163
	}
L163:
	;
	v799 = base.I64_extend_i32_u(v701) & int64(255)
	goto L159
L164:
	;
	if base.Ui32(int32(8)) < base.Ui32((v720+int32(-49))&int32(255)) {
		goto L156
	} else {
		goto L167
	}
L165:
	;
	v967 = int32(2)
	v717 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v670)+1)))
	v719 = v967
	v720 = v717
	v721 = l1 + v967
	goto L164
L166:
	;
	v719 = v696
	v720 = v697
	v721 = v670
	goto L164
L167:
	;
	v732 = base.I64_extend_i32_u(v720+int32(-48)) & int64(255)
	if base.Ui32(v681) <= base.Ui32(v719) {
		v775 = v732
		goto L168
	} else {
		goto L169
	}
L168:
	;
	if v697&int32(255) != int32(45) {
		goto L176
	} else {
		goto L177
	}
L169:
	;
	v738 = v719
	v740 = v732
	v742 = v721
	goto L170
L170:
	;
	v744 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v742)+1)))
	if base.Ui32((v744+int32(-58))&int32(255)) < base.Ui32(int32(246)) {
		goto L156
	} else {
		goto L172
	}
L171:
	;
	v775 = v765
	goto L168
L172:
	;
	if base.Ui64(int64(1844674407370955161)) < base.Ui64(v740) {
		goto L156
	} else {
		goto L173
	}
L173:
	;
	v754 = v740 * int64(10)
	v759 = base.I64_extend_i32_u(v744+int32(-48)) & int64(255)
	if base.Ui64(v759^int64(-1)) < base.Ui64(v754) {
		goto L156
	} else {
		goto L174
	}
L174:
	;
	v763 = int32(1)
	v765 = v754 + v759
	v767 = v738 + v763
	if v767 != v681 {
		v738 = v767
		v740 = v765
		v742 = v742 + v763
		goto L170
	} else {
		goto L175
	}
L175:
	;
	goto L171
L176:
	;
	if v775 < int64(0) {
		goto L156
	} else {
		goto L180
	}
L177:
	;
	if base.Ui64(int64(-9223372036854775807-1)) < base.Ui64(v775) {
		goto L156
	} else {
		goto L178
	}
L178:
	;
	if v683 == int32(0) {
		goto L158
	} else {
		goto L179
	}
L179:
	;
	v799 = int64(0) - v775
	goto L159
L180:
	;
	if v683 == int32(0) {
		goto L158
	} else {
		goto L181
	}
L181:
	;
	v799 = v775
	goto L159
L182:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v828
	v832 = v678 + int32(2)
	v833 = *(*int64)(unsafe.Add(mBase, uint32(v11)+8))
	if v833 < int64(1) {
		v872 = v832
		v874 = v828
		goto L183
	} else {
		goto L184
	}
L183:
	;
	v881 = F_sdscatlen(m, v874, int32(_a_F_debugScriptRespToHuman_7), int32(1))
	mBase = m.M
	v882 = m.ExcPending
	if v882 != 0 {
		goto L16
	} else {
		goto L194
	}
L184:
	;
	v838 = v832
	v842 = int64(0)
	goto L185
L185:
	;
	v845 = F_debugScriptRespToHuman(m, l0, v838)
	mBase = m.M
	v846 = m.ExcPending
	if v846 != 0 {
		goto L16
	} else {
		goto L187
	}
L186:
	;
	v870 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v872 = v853
	v874 = v870
	goto L183
L187:
	;
	v847 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v850 = F_sdscatlen(m, v847, int32(_a_F_debugScriptRespToHuman_8), int32(4))
	mBase = m.M
	v851 = m.ExcPending
	if v851 != 0 {
		goto L16
	} else {
		goto L188
	}
L188:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v850
	v853 = F_debugScriptRespToHuman(m, l0, v845)
	mBase = m.M
	v854 = m.ExcPending
	if v854 != 0 {
		goto L16
	} else {
		goto L189
	}
L189:
	;
	v855 = *(*int64)(unsafe.Add(mBase, uint32(v11)+8))
	if v855+int64(-1) == v842 {
		v866 = v855
		goto L190
	} else {
		goto L191
	}
L190:
	;
	v868 = v842 + int64(1)
	if v868 < v866 {
		v838 = v853
		v842 = v868
		goto L185
	} else {
		goto L193
	}
L191:
	;
	v859 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v862 = F_sdscatlen(m, v859, int32(_a_F_debugScriptRespToHuman_3), int32(1))
	mBase = m.M
	v863 = m.ExcPending
	if v863 != 0 {
		goto L16
	} else {
		goto L192
	}
L192:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v862
	v865 = *(*int64)(unsafe.Add(mBase, uint32(v11)+8))
	v866 = v865
	goto L190
L193:
	;
	goto L186
L194:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v881
	v953 = v872
	goto L1
L195:
	;
	v894 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v897 = F_sdscatlen(m, v894, int32(_a_F_debugScriptRespToHuman_9), int32(6))
	mBase = m.M
	v898 = m.ExcPending
	if v898 != 0 {
		goto L16
	} else {
		goto L199
	}
L196:
	;
	v893 = v887
	goto L198
L197:
	;
	v893 = int32(0)
	goto L198
L198:
	;
	goto L195
L199:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v897
	v953 = v893 + int32(2)
	goto L1
L200:
	;
	v912 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v913 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+1)))
	if v913 != int32(116) {
		goto L205
	} else {
		goto L206
	}
L201:
	;
	v911 = v905
	goto L203
L202:
	;
	v911 = int32(0)
	goto L203
L203:
	;
	goto L200
L204:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v924
	v953 = v911 + int32(2)
	goto L1
L205:
	;
	v922 = F_sdscatlen(m, v912, int32(_a_F_debugScriptRespToHuman_10), int32(6))
	mBase = m.M
	v923 = m.ExcPending
	if v923 != 0 {
		goto L16
	} else {
		goto L208
	}
L206:
	;
	v918 = F_sdscatlen(m, v912, int32(_a_F_debugScriptRespToHuman_11), int32(5))
	mBase = m.M
	v919 = m.ExcPending
	if v919 != 0 {
		goto L16
	} else {
		goto L207
	}
L207:
	;
	v924 = v918
	goto L204
L208:
	;
	v924 = v922
	goto L204
L209:
	;
	v938 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v941 = F_sdscatlen(m, v938, int32(_a_F_debugScriptRespToHuman_12), int32(9))
	mBase = m.M
	v942 = m.ExcPending
	if v942 != 0 {
		goto L16
	} else {
		goto L213
	}
L210:
	;
	v937 = v931
	goto L212
L211:
	;
	v937 = int32(0)
	goto L212
L212:
	;
	goto L209
L213:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v941
	v947 = F_sdscatlen(m, v941, v929, v937+(l1^int32(-1)))
	mBase = m.M
	v948 = m.ExcPending
	if v948 != 0 {
		goto L16
	} else {
		goto L214
	}
L214:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v947
	v953 = v937 + int32(2)
	goto L1
}
func F_decrRefCount(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
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
	var v52 int32
	_ = v52
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v76 int32
	_ = v76
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v119 int32
	_ = v119
	var v126 int32
	_ = v126
	var v129 int32
	_ = v129
	var v132 int32
	_ = v132
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v141 int32
	_ = v141
	var v155 int32
	_ = v155
	var v159 int32
	_ = v159
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v164 int32
	_ = v164
	var v166 int32
	_ = v166
	var v172 int32
	_ = v172
	var v175 int32
	_ = v175
	var v177 int32
	_ = v177
	var v190 int32
	_ = v190
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v8 = int32(base.Ui32(v6) >> (uint(int32(3)) % 32))
	if v8 != int32(1) {
		if base.Ui32(v6) <= base.Ui32(int32(7)) {
			F__serverPanic_1(m, int32(_a_F_decrRefCount_0), int32(643), int32(_a_F_decrRefCount_1), int32(0))
			mBase = m.M
			v190 = m.ExcPending
			if v190 != 0 {
				return
			} else {
				F_abort(m)
				mBase = m.M
				base.Wasm_trap_unreachable()
				for {
				}
			}
		} else {
			if v8 == int32(536870911) {
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v6 + int32(-8)
			}
			return
		}
	} else {
		if v6&int32(4) == int32(0) {
			v69 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			v70 = v69
		} else {
			if v6&int32(1) != 0 {
				v19 = int32(16)
			} else {
				v19 = int32(8)
			}
			v20 = l0 + v19
			if v6&int32(2) == int32(0) {
				v52 = v20
			} else {
				v25 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20))))
				v26 = v20 + v25
				v30 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v26))))
				switch v30 & int32(7) {
				case 0:
					v47 = int32(base.Ui32(v30) >> (uint(int32(3)) % 32))
				case 1:
					v37 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v26+int32(-2)))))
					v47 = v37
				case 2:
					v40 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v26+int32(-4)))))
					v47 = v40
				case 3:
					v43 = *(*int32)(unsafe.Add(mBase, uint32(v26+int32(-8))))
					v47 = v43
				case 4:
					v46 = *(*int32)(unsafe.Add(mBase, uint32(v26+int32(-16))))
					v47 = v46
				default:
					v47 = int32(0)
				}
				v52 = v26 + int32(1) + v47 + int32(1)
			}
			v66 = *(*int32)(unsafe.Add(mBase, _c_F_decrRefCount[0]))
			v70 = v52 + v66
		}
		if v70 == int32(0) {
			F_valkey_free(m, l0)
			mBase = m.M
			v177 = m.ExcPending
			if v177 != 0 {
				return
			} else {
				return
			}
		} else {
			v76 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			switch v76 & int32(15) {
			case 0:
				F_freeStringObject(m, l0)
				mBase = m.M
				v175 = m.ExcPending
				if v175 != 0 {
					return
				} else {
					F_valkey_free(m, l0)
					mBase = m.M
					v177 = m.ExcPending
					if v177 != 0 {
						return
					} else {
						return
					}
				}
			case 1:
				F_freeListObject(m, l0)
				mBase = m.M
				v80 = m.ExcPending
				if v80 != 0 {
					return
				} else {
					F_valkey_free(m, l0)
					mBase = m.M
					v82 = m.ExcPending
					if v82 != 0 {
						return
					} else {
						return
					}
				}
			case 2:
				F_freeSetObject(m, l0)
				mBase = m.M
				v84 = m.ExcPending
				if v84 != 0 {
					return
				} else {
					F_valkey_free(m, l0)
					mBase = m.M
					v86 = m.ExcPending
					if v86 != 0 {
						return
					} else {
						return
					}
				}
			case 3:
				F_freeZsetObject(m, l0)
				mBase = m.M
				v88 = m.ExcPending
				if v88 != 0 {
					return
				} else {
					F_valkey_free(m, l0)
					mBase = m.M
					v90 = m.ExcPending
					if v90 != 0 {
						return
					} else {
						return
					}
				}
			case 4:
				F_freeHashObject(m, l0)
				mBase = m.M
				v92 = m.ExcPending
				if v92 != 0 {
					return
				} else {
					F_valkey_free(m, l0)
					mBase = m.M
					v94 = m.ExcPending
					if v94 != 0 {
						return
					} else {
						return
					}
				}
			case 5:
				F_freeModuleObject(m, l0)
				mBase = m.M
				v96 = m.ExcPending
				if v96 != 0 {
					return
				} else {
					F_valkey_free(m, l0)
					mBase = m.M
					v98 = m.ExcPending
					if v98 != 0 {
						return
					} else {
						return
					}
				}
			case 6:
				v99 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
				if v99&int32(4) == int32(0) {
					v162 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
					F_freeStream(m, v162)
					mBase = m.M
					v164 = m.ExcPending
					if v164 != 0 {
						return
					} else {
						F_valkey_free(m, l0)
						mBase = m.M
						v166 = m.ExcPending
						if v166 != 0 {
							return
						} else {
							return
						}
					}
				} else {
					if v99&int32(1) != 0 {
						v108 = int32(16)
					} else {
						v108 = int32(8)
					}
					v109 = l0 + v108
					if v99&int32(2) == int32(0) {
						v141 = v109
					} else {
						v114 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v109))))
						v115 = v109 + v114
						v119 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v115))))
						switch v119 & int32(7) {
						case 0:
							v136 = int32(base.Ui32(v119) >> (uint(int32(3)) % 32))
						case 1:
							v126 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v115+int32(-2)))))
							v136 = v126
						case 2:
							v129 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v115+int32(-4)))))
							v136 = v129
						case 3:
							v132 = *(*int32)(unsafe.Add(mBase, uint32(v115+int32(-8))))
							v136 = v132
						case 4:
							v135 = *(*int32)(unsafe.Add(mBase, uint32(v115+int32(-16))))
							v136 = v135
						default:
							v136 = int32(0)
						}
						v141 = v115 + int32(1) + v136 + int32(1)
					}
					v155 = *(*int32)(unsafe.Add(mBase, _c_F_decrRefCount[0]))
					F_freeStream(m, v141+v155)
					mBase = m.M
					v159 = m.ExcPending
					if v159 != 0 {
						return
					} else {
						F_valkey_free(m, l0)
						mBase = m.M
						v161 = m.ExcPending
						if v161 != 0 {
							return
						} else {
							return
						}
					}
				}
			default:
				F__serverPanic_1(m, int32(_a_F_decrRefCount_0), int32(638), int32(_a_F_decrRefCount_2), int32(0))
				mBase = m.M
				v172 = m.ExcPending
				if v172 != 0 {
					return
				} else {
					F_abort(m)
					mBase = m.M
					base.Wasm_trap_unreachable()
					for {
					}
				}
			}
		}
	}
}
func F_defragWhileBlocked(m *base.Module) {
	return
}
func F_delKeysNotOwnedByMyself(m *base.Module, l0 int32) {
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
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v29 int32
	_ = v29
	var v33 int32
	_ = v33
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
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
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v90 int32
	_ = v90
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	v12 = v9 + int32(8)
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(v12)+4)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v12))) = v13
	goto L1
L1:
	;
	v18 = v9 + int32(8)
	v20 = *(*int32)(unsafe.Add(mBase, uint32(v18)))
	if v20 == int32(0) {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	m.G0 = v9 + int32(16)
	return
L3:
	;
	if v20 == int32(0) {
		goto L2
	} else {
		goto L6
	}
L4:
	;
	goto L3
L5:
	;
	v23 = *(*int32)(unsafe.Add(mBase, uint32(v18)+4))
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v20+base.B2i32(v23 == int32(0))<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v18))) = v29
	goto L4
L6:
	;
	v33 = v20
	goto L7
L7:
	;
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v33)+8))
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v39)))
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v39)+4))
	if v41 < v40 {
		goto L9
	} else {
		goto L10
	}
L8:
	;
	goto L2
L9:
	;
	v79 = v9 + int32(8)
	v81 = *(*int32)(unsafe.Add(mBase, uint32(v79)))
	if v81 == int32(0) {
		goto L19
	} else {
		goto L20
	}
L10:
	;
	v44 = *(*int32)(unsafe.Add(mBase, _c_F_delKeysNotOwnedByMyself[0]))
	v45 = v40
	v48 = v41
	v49 = v44
	goto L11
L11:
	;
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v49+v45<<(uint(int32(2))%32)+int32(52))))
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v49)))
	if v56 == v57 {
		v67 = v48
		v68 = v49
		goto L13
	} else {
		goto L14
	}
L12:
	;
	goto L9
L13:
	;
	if v45 < v67 {
		v45 = v45 + int32(1)
		v48 = v67
		v49 = v68
		goto L11
	} else {
		goto L17
	}
L14:
	;
	v59 = int32(1)
	v62 = F_delKeysInSlot(m, v45, v59, v59, int32(0))
	mBase = m.M
	v63 = m.ExcPending
	if v63 != 0 {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	return
L16:
	;
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v39)+4))
	v66 = *(*int32)(unsafe.Add(mBase, _c_F_delKeysNotOwnedByMyself[0]))
	v67 = v64
	v68 = v66
	goto L13
L17:
	;
	goto L12
L18:
	;
	if v81 != 0 {
		v33 = v81
		goto L7
	} else {
		goto L21
	}
L19:
	;
	goto L18
L20:
	;
	v84 = *(*int32)(unsafe.Add(mBase, uint32(v79)+4))
	v90 = *(*int32)(unsafe.Add(mBase, uint32(v81+base.B2i32(v84 == int32(0))<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v79))) = v90
	goto L19
L21:
	;
	goto L8
}
func F_delifeqCommand(m *base.Module, l0 int32) {
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
	var v20 int32
	_ = v20
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
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v59 int64
	_ = v59
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v78 int32
	_ = v78
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v9 = *(*int32)(unsafe.Add(mBase, uint32(v8)+4))
	v11 = *(*int32)(unsafe.Add(mBase, _c_F_delifeqCommand[0]))
	v12 = F_lookupKeyWriteOrReply(m, l0, v9, v11)
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return
	} else {
		if v12 == int32(0) {
			m.G0 = v6 + int32(16)
			return
		} else {
			v17 = F_checkType(m, l0, v12, int32(0))
			mBase = m.M
			v18 = m.ExcPending
			if v18 != 0 {
				return
			} else {
				if v17 != 0 {
					m.G0 = v6 + int32(16)
					return
				} else {
					v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
					v20 = *(*int32)(unsafe.Add(mBase, uint32(v19)+8))
					v21 = F_compareStringObjects(m, v12, v20)
					mBase = m.M
					v22 = m.ExcPending
					if v22 != 0 {
						return
					} else {
						if v21 == int32(0) {
							v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
							v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
							v30 = *(*int32)(unsafe.Add(mBase, uint32(v29)+4))
							v31 = F_dbSyncDelete(m, v28, v30)
							mBase = m.M
							v32 = m.ExcPending
							if v32 != 0 {
								return
							} else {
								if v31 == int32(0) {
									F__serverAssert(m, int32(_a_F_delifeqCommand_0), int32(_a_F_delifeqCommand_1), int32(292))
									mBase = m.M
									v78 = m.ExcPending
									if v78 != 0 {
										return
									} else {
										F_abort(m)
										mBase = m.M
										base.Wasm_trap_unreachable()
										for {
										}
									}
								} else {
									v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
									v36 = *(*int32)(unsafe.Add(mBase, uint32(v35)+4))
									v38 = *(*int32)(unsafe.Add(mBase, _c_F_delifeqCommand[1]))
									*(*int32)(unsafe.Add(mBase, uint32(v6))) = v38
									*(*int32)(unsafe.Add(mBase, uint32(v6)+4)) = v36
									F_rewriteClientCommandVector(m, l0, int32(2), v6)
									mBase = m.M
									v43 = m.ExcPending
									if v43 != 0 {
										return
									} else {
										v44 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
										v45 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
										v46 = *(*int32)(unsafe.Add(mBase, uint32(v45)+4))
										F_signalModifiedKey(m, l0, v44, v46)
										mBase = m.M
										v48 = m.ExcPending
										if v48 != 0 {
											return
										} else {
											v51 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
											v52 = *(*int32)(unsafe.Add(mBase, uint32(v51)+4))
											v53 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
											v54 = *(*int32)(unsafe.Add(mBase, uint32(v53)+28))
											F_notifyKeyspaceEvent(m, int32(4), int32(_a_F_delifeqCommand_2), v52, v54)
											mBase = m.M
											v56 = m.ExcPending
											if v56 != 0 {
												return
											} else {
												v57 = int32(_a_F_delifeqCommand_3)
												v59 = *(*int64)(unsafe.Add(mBase, _c_F_delifeqCommand[2]))
												*(*int64)(unsafe.Add(mBase, _c_F_delifeqCommand[2])) = v59 + int64(1)
												v66 = int32(_a_F_delifeqCommand_4)
												v67 = *(*int32)(unsafe.Add(mBase, uint32(v66)))
												F_addReply(m, l0, v67)
												mBase = m.M
												v69 = m.ExcPending
												if v69 != 0 {
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
						} else {
							v66 = int32(_a_F_delifeqCommand_5)
							v67 = *(*int32)(unsafe.Add(mBase, uint32(v66)))
							F_addReply(m, l0, v67)
							mBase = m.M
							v69 = m.ExcPending
							if v69 != 0 {
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
func F_dirExists(m *base.Module, l0 int32) int32 {
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
	v4 = m.G0
	v6 = v4 - int32(96)
	m.G0 = v6
	v10 = F___fstatat(m, int32(-100), l0, v6, int32(0))
	mBase = m.M
	v11 = *(*int32)(unsafe.Add(mBase, uint32(v6)+4))
	m.G0 = v6 + int32(96)
	return base.B2i32(v10 == int32(0)) & base.B2i32(v11&int32(61440) == int32(16384))
}
func F_disconnectOrRedirectAllBlockedClients(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v28 int32
	_ = v28
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v68 int32
	_ = v68
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v105 int64
	_ = v105
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v127 int32
	_ = v127
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v153 int32
	_ = v153
	var v155 int32
	_ = v155
	var v158 int32
	_ = v158
	var v164 int32
	_ = v164
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	v9 = *(*int32)(unsafe.Add(mBase, _c_F_disconnectOrRedirectAllBlockedClients[0]))
	v11 = v6 + int32(8)
	v12 = *(*int32)(unsafe.Add(mBase, uint32(v9)))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+4)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v11))) = v12
	goto L1
L1:
	;
	v17 = v6 + int32(8)
	v19 = *(*int32)(unsafe.Add(mBase, uint32(v17)))
	if v19 == int32(0) {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	m.G0 = v6 + int32(16)
	return
L3:
	;
	if v19 == int32(0) {
		goto L2
	} else {
		goto L6
	}
L4:
	;
	goto L3
L5:
	;
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v17)+4))
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v19+base.B2i32(v22 == int32(0))<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v17))) = v28
	goto L4
L6:
	;
	v33 = v19
	goto L7
L7:
	;
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v33)+8))
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v35)+200))
	if v36&int32(16) == int32(0) {
		goto L9
	} else {
		goto L10
	}
L8:
	;
	goto L2
L9:
	;
	v153 = v6 + int32(8)
	v155 = *(*int32)(unsafe.Add(mBase, uint32(v153)))
	if v155 == int32(0) {
		goto L47
	} else {
		goto L48
	}
L10:
	;
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v35)+116))
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v41)))
	if v42 == int32(6) {
		goto L9
	} else {
		goto L11
	}
L11:
	;
	v46 = *(*int32)(unsafe.Add(mBase, _c_F_disconnectOrRedirectAllBlockedClients[1]))
	if v46 == int32(0) {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	if v36&int32(131072) == int32(0) {
		goto L21
	} else {
		goto L22
	}
L13:
	;
	v49 = F_clusterRedirectBlockedClientIfNeeded(m, v35)
	mBase = m.M
	v50 = m.ExcPending
	if v50 != 0 {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	return
L15:
	;
	if v49 == int32(0) {
		goto L9
	} else {
		goto L16
	}
L16:
	;
	v53 = int32(0)
	F_updateStatsOnUnblock(m, v35, v53, v53, int32(1))
	mBase = m.M
	v57 = m.ExcPending
	if v57 != 0 {
		goto L14
	} else {
		goto L17
	}
L17:
	;
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v35)+204))
	if v58&int32(2) == int32(0) {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	F_unblockClient(m, v35, int32(1))
	mBase = m.M
	v68 = m.ExcPending
	if v68 != 0 {
		goto L14
	} else {
		goto L20
	}
L19:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v35)+204)) = v58 & int32(-3)
	goto L18
L20:
	;
	goto L9
L21:
	;
	v79 = int32(0)
	v82 = *(*int32)(unsafe.Add(mBase, _c_F_disconnectOrRedirectAllBlockedClients[1]))
	if v82 != 0 {
		v90 = v79
		goto L26
	} else {
		goto L27
	}
L22:
	;
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v35)+72))
	v74 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v73)+56)))
	if v74&int32(1) == int32(0) {
		goto L9
	} else {
		goto L23
	}
L23:
	;
	goto L21
L24:
	;
	F_addReplyError(m, v35, int32(_a_F_disconnectOrRedirectAllBlockedClients_0))
	mBase = m.M
	v130 = m.ExcPending
	if v130 != 0 {
		goto L14
	} else {
		goto L41
	}
L25:
	;
	if v90 == int32(0) {
		goto L24
	} else {
		goto L29
	}
L26:
	;
	goto L25
L27:
	;
	v83 = int32(0)
	v84 = *(*int32)(unsafe.Add(mBase, _c_F_disconnectOrRedirectAllBlockedClients[2]))
	if v84 == v83 {
		v90 = v79
		goto L26
	} else {
		goto L28
	}
L28:
	;
	v87 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v35)+227)))
	v90 = v87 & int32(1)
	goto L26
L29:
	;
	v93 = *(*int32)(unsafe.Add(mBase, uint32(v35)+116))
	v94 = *(*int32)(unsafe.Add(mBase, uint32(v93)))
	switch v94 + int32(-1) {
	case 0, 3, 4:
		goto L30
	default:
		goto L24
	case 2:
		goto L31
	}
L30:
	;
	v102 = F_sdsempty(m)
	mBase = m.M
	v103 = m.ExcPending
	if v103 != 0 {
		goto L14
	} else {
		goto L34
	}
L31:
	;
	v97 = *(*int32)(unsafe.Add(mBase, uint32(v35)+116))
	v98 = *(*int32)(unsafe.Add(mBase, uint32(v97)+48))
	v99 = *(*int32)(unsafe.Add(mBase, uint32(v98)+44))
	goto L32
L32:
	;
	if v99 == int32(0) {
		goto L9
	} else {
		goto L33
	}
L33:
	;
	goto L30
L34:
	;
	v105 = *(*int64)(unsafe.Add(mBase, _c_F_disconnectOrRedirectAllBlockedClients[2]))
	*(*int64)(unsafe.Add(mBase, uint32(v6))) = v105
	v108 = F_sdscatprintf(m, v102, int32(_a_F_disconnectOrRedirectAllBlockedClients_1), v6)
	mBase = m.M
	v109 = m.ExcPending
	if v109 != 0 {
		goto L14
	} else {
		goto L35
	}
L35:
	;
	F_addReplyErrorSds(m, v35, v108)
	mBase = m.M
	v111 = m.ExcPending
	if v111 != 0 {
		goto L14
	} else {
		goto L36
	}
L36:
	;
	v112 = int32(0)
	F_updateStatsOnUnblock(m, v35, v112, v112, int32(1))
	mBase = m.M
	v116 = m.ExcPending
	if v116 != 0 {
		goto L14
	} else {
		goto L37
	}
L37:
	;
	v117 = *(*int32)(unsafe.Add(mBase, uint32(v35)+204))
	if v117&int32(2) == int32(0) {
		goto L38
	} else {
		goto L39
	}
L38:
	;
	F_unblockClient(m, v35, int32(1))
	mBase = m.M
	v127 = m.ExcPending
	if v127 != 0 {
		goto L14
	} else {
		goto L40
	}
L39:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v35)+204)) = v117 & int32(-3)
	goto L38
L40:
	;
	goto L9
L41:
	;
	v131 = int32(0)
	F_updateStatsOnUnblock(m, v35, v131, v131, int32(1))
	mBase = m.M
	v135 = m.ExcPending
	if v135 != 0 {
		goto L14
	} else {
		goto L42
	}
L42:
	;
	v136 = *(*int32)(unsafe.Add(mBase, uint32(v35)+204))
	if v136&int32(2) == int32(0) {
		goto L43
	} else {
		goto L44
	}
L43:
	;
	F_unblockClient(m, v35, int32(1))
	mBase = m.M
	v146 = m.ExcPending
	if v146 != 0 {
		goto L14
	} else {
		goto L45
	}
L44:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v35)+204)) = v136 & int32(-3)
	goto L43
L45:
	;
	v147 = *(*int32)(unsafe.Add(mBase, uint32(v35)+200))
	*(*int32)(unsafe.Add(mBase, uint32(v35)+200)) = v147 | int32(64)
	goto L9
L46:
	;
	if v155 != 0 {
		v33 = v155
		goto L7
	} else {
		goto L49
	}
L47:
	;
	goto L46
L48:
	;
	v158 = *(*int32)(unsafe.Add(mBase, uint32(v153)+4))
	v164 = *(*int32)(unsafe.Add(mBase, uint32(v155+base.B2i32(v158 == int32(0))<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v153))) = v164
	goto L47
L49:
	;
	goto L8
}
func F_dismissObject(m *base.Module, l0 int32, l1 int32) {
	return
}
func F_do_putc_2(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v1 int32
	_ = v1
	var v2 int32
	_ = v2
	var v5 int32
	_ = v5
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	v1 = l0
	v2 = int32(0)
	v5 = *(*int32)(unsafe.Add(mBase, _c_F_do_putc_2[0]))
	if v5 < v2 {
		v16 = v1 & int32(255)
		v18 = *(*int32)(unsafe.Add(mBase, _c_F_do_putc_2[1]))
		if v16 == v18 {
			v33 = F___overflow(m, int32(_a_F_do_putc_2_0), v16)
			mBase = m.M
			v36 = m.ExcPending
			if v36 != 0 {
				return int32(0)
			} else {
				return v33
			}
		} else {
			v20 = int32(0)
			v21 = *(*int32)(unsafe.Add(mBase, _c_F_do_putc_2[2]))
			v23 = *(*int32)(unsafe.Add(mBase, _c_F_do_putc_2[3]))
			if v21 == v23 {
				v33 = F___overflow(m, int32(_a_F_do_putc_2_0), v16)
				mBase = m.M
				v36 = m.ExcPending
				if v36 != 0 {
					return int32(0)
				} else {
					return v33
				}
			} else {
				*(*int32)(unsafe.Add(mBase, _c_F_do_putc_2[2])) = v21 + int32(1)
				*(*uint8)(unsafe.Add(mBase, uint32(v21))) = uint8(v1)
				return v16
			}
		}
	} else {
		if v5 == int32(0) {
			v38 = F_locking_putc_2(m, v1)
			mBase = m.M
			v39 = m.ExcPending
			if v39 != 0 {
				return int32(0)
			} else {
				return v38
			}
		} else {
			v13 = *(*int32)(unsafe.Add(mBase, _c_F_do_putc_2[4]))
			if v5&int32(1073741823) != v13 {
				v38 = F_locking_putc_2(m, v1)
				mBase = m.M
				v39 = m.ExcPending
				if v39 != 0 {
					return int32(0)
				} else {
					return v38
				}
			} else {
				v16 = v1 & int32(255)
				v18 = *(*int32)(unsafe.Add(mBase, _c_F_do_putc_2[1]))
				if v16 == v18 {
					v33 = F___overflow(m, int32(_a_F_do_putc_2_0), v16)
					mBase = m.M
					v36 = m.ExcPending
					if v36 != 0 {
						return int32(0)
					} else {
						return v33
					}
				} else {
					v20 = int32(0)
					v21 = *(*int32)(unsafe.Add(mBase, _c_F_do_putc_2[2]))
					v23 = *(*int32)(unsafe.Add(mBase, _c_F_do_putc_2[3]))
					if v21 == v23 {
						v33 = F___overflow(m, int32(_a_F_do_putc_2_0), v16)
						mBase = m.M
						v36 = m.ExcPending
						if v36 != 0 {
							return int32(0)
						} else {
							return v33
						}
					} else {
						*(*int32)(unsafe.Add(mBase, _c_F_do_putc_2[2])) = v21 + int32(1)
						*(*uint8)(unsafe.Add(mBase, uint32(v21))) = uint8(v1)
						return v16
					}
				}
			}
		}
	}
}
func F_do_setrlimit(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if int32(0) < v2 {
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = int32(0)
	}
	return
}
func F_do_tzset(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v8 int32
	_ = v8
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v23 int32
	_ = v23
	v2 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_do_tzset[0])))
	if v2&int32(1) != 0 {
	} else {
		v8 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_do_tzset[0])))
		if v8&int32(1) != 0 {
		} else {
			v13 = int32(9116432)
			v14 = int32(9116464)
			m.Env.X_tzset_js(m, int32(9116380), int32(9116384), v13, v14)
			mBase = m.M
			v16 = int32(0)
			*(*int32)(unsafe.Add(mBase, _c_F_do_tzset[1])) = v14
			*(*int32)(unsafe.Add(mBase, _c_F_do_tzset[2])) = v13
			v23 = int32(1)
			*(*uint8)(unsafe.Add(mBase, _c_F_do_tzset[0])) = uint8(v23)
		}
	}
	return
}
func F_doesCommandHaveChannelsWithFlags(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v23 int64
	_ = v23
	var v29 int32
	_ = v29
	var v34 int32
	_ = v34
	v7 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+59)))
	if v7&int32(8) != 0 {
		v34 = int32(1)
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return v34
L2:
	;
	v10 = int32(0)
	v12 = *(*int32)(unsafe.Add(mBase, _c_F_doesCommandHaveChannelsWithFlags[0]))
	if v12 == v10 {
		v34 = v10
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v17 = int32(_a_F_doesCommandHaveChannelsWithFlags_0)
	v20 = v12
	goto L4
L4:
	;
	if v15 != v20 {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	v34 = v10
	goto L1
L6:
	;
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v17)+24))
	if v29 != 0 {
		v17 = v17 + int32(24)
		v20 = v29
		goto L4
	} else {
		goto L8
	}
L7:
	;
	v23 = *(*int64)(unsafe.Add(mBase, uint32(v17)+8))
	return base.B2i32(v23&base.I64_extend_i32_s(l1) != int64(0))
L8:
	;
	goto L5
}
func F_dprintf(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
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
	v10 = F_vdprintf(m, l0, l1, l2)
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return int32(0)
	} else {
		m.G0 = v7 + int32(16)
		return v10
	}
}
func F_dualChannelReplMainConnRecvCapaReply(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v13 int32
	_ = v13
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v42 int64
	_ = v42
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v52 int32
	_ = v52
	var v56 int32
	_ = v56
	var v65 int32
	_ = v65
	var v69 int32
	_ = v69
	v5 = m.G0
	v7 = v5 - int32(288)
	m.G0 = v7
	v13 = *(*int32)(unsafe.Add(mBase, _c_F_dualChannelReplMainConnRecvCapaReply[0]))
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v18 = *(*int32)(unsafe.Add(mBase, uint32(v17)+100))
	v19 = m.T0[v18].(func(*base.Module, int32, int32, int32, int64) int32)(m, l0, v7+int32(32), int32(256), base.I64_extend_i32_s(v13*int32(1000)))
	mBase = m.M
	v22 = m.ExcPending
	if v22 != 0 {
		return int32(0)
	} else {
		if v19 != int32(-1) {
			v40 = int32(_a_F_dualChannelReplMainConnRecvCapaReply_0)
			v42 = *(*int64)(unsafe.Add(mBase, _c_F_dualChannelReplMainConnRecvCapaReply[1]))
			*(*int64)(unsafe.Add(mBase, _c_F_dualChannelReplMainConnRecvCapaReply[2])) = v42
			v46 = F_sdsnew(m, v7+int32(32))
			mBase = m.M
			v47 = m.ExcPending
			if v47 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(l1))) = v46
				if v46 == int32(0) {
					v69 = int32(-1)
					m.G0 = v7 + int32(288)
					return v69
				} else {
					v52 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v46))))
					if v52 != int32(45) {
						v69 = int32(0)
						m.G0 = v7 + int32(288)
						return v69
					} else {
						v56 = *(*int32)(unsafe.Add(mBase, _c_F_dualChannelReplMainConnRecvCapaReply[3]))
						if int32(2) < v56 {
							v69 = int32(-1)
							m.G0 = v7 + int32(288)
							return v69
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v7)+16)) = v46
							F__serverLog(m, int32(2), int32(_a_F_dualChannelReplMainConnRecvCapaReply_1), v7+int32(16))
							mBase = m.M
							v65 = m.ExcPending
							if v65 != 0 {
								return int32(0)
							} else {
								v69 = int32(-1)
								m.G0 = v7 + int32(288)
								return v69
							}
						}
					}
				}
			}
		} else {
			v26 = *(*int32)(unsafe.Add(mBase, _c_F_dualChannelReplMainConnRecvCapaReply[3]))
			if int32(3) < v26 {
				*(*int32)(unsafe.Add(mBase, uint32(l1))) = int32(0)
				v69 = int32(-1)
				m.G0 = v7 + int32(288)
				return v69
			} else {
				v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
				v30 = *(*int32)(unsafe.Add(mBase, uint32(v29)+88))
				v31 = m.T0[v30].(func(*base.Module, int32) int32)(m, l0)
				mBase = m.M
				v32 = m.ExcPending
				if v32 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v7))) = v31
					F__serverLog(m, int32(3), int32(_a_F_dualChannelReplMainConnRecvCapaReply_2), v7)
					mBase = m.M
					v37 = m.ExcPending
					if v37 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(l1))) = int32(0)
						v69 = int32(-1)
						m.G0 = v7 + int32(288)
						return v69
					}
				}
			}
		}
	}
}
func F_dualChannelSetupMainConnForPsync(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v11 int32
	_ = v11
	var v15 int32
	_ = v15
	var v18 int64
	_ = v18
	var v19 int32
	_ = v19
	var v23 int32
	_ = v23
	var v31 int32
	_ = v31
	var v32 int64
	_ = v32
	var v53 int32
	_ = v53
	var v58 int32
	_ = v58
	var v65 int32
	_ = v65
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v76 int64
	_ = v76
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v97 int64
	_ = v97
	var v99 int32
	_ = v99
	var v103 int64
	_ = v103
	var v104 int64
	_ = v104
	var v113 int32
	_ = v113
	var v116 int32
	_ = v116
	var v123 int32
	_ = v123
	var v124 int64
	_ = v124
	var v126 int32
	_ = v126
	var v131 int32
	_ = v131
	var v140 int32
	_ = v140
	var v144 int32
	_ = v144
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v181 int32
	_ = v181
	var v183 int32
	_ = v183
	var v187 int32
	_ = v187
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v195 int32
	_ = v195
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v218 int32
	_ = v218
	var v229 int32
	_ = v229
	var v235 int32
	_ = v235
	var v240 int32
	_ = v240
	var v242 int32
	_ = v242
	var v249 int32
	_ = v249
	var v252 int32
	_ = v252
	var v253 int32
	_ = v253
	var v255 int32
	_ = v255
	var v257 int32
	_ = v257
	v4 = m.G0
	v6 = v4 - int32(96)
	m.G0 = v6
	*(*int32)(unsafe.Add(mBase, uint32(v6)+60)) = int32(0)
	v11 = *(*int32)(unsafe.Add(mBase, _c_F_dualChannelSetupMainConnForPsync[0]))
	switch v11 + int32(-4) {
	case 0:
		goto L8
	default:
		goto L4
	case 4:
		goto L7
	case 7:
		goto L6
	case 8:
		goto L5
	}
L1:
	;
	v255 = *(*int32)(unsafe.Add(mBase, uint32(v6)+60))
	F_sdsfree(m, v255)
	mBase = m.M
	v257 = m.ExcPending
	if v257 != 0 {
		goto L53
	} else {
		goto L82
	}
L2:
	;
	v235 = *(*int32)(unsafe.Add(mBase, _c_F_dualChannelSetupMainConnForPsync[1]))
	if int32(3) < v235 {
		goto L75
	} else {
		goto L76
	}
L3:
	;
	if v215 != int32(-1) {
		goto L1
	} else {
		goto L74
	}
L4:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v6))) = v11
	F__serverPanic_1(m, int32(_a_F_dualChannelSetupMainConnForPsync_0), int32(3769), int32(_a_F_dualChannelSetupMainConnForPsync_1), v6)
	mBase = m.M
	v229 = m.ExcPending
	if v229 != 0 {
		goto L53
	} else {
		goto L73
	}
L5:
	;
	v215 = F_dualChannelReplMainConnRecvPsyncReply(m, l0, v6+int32(60))
	mBase = m.M
	v216 = m.ExcPending
	if v216 != 0 {
		goto L53
	} else {
		goto L70
	}
L6:
	;
	v187 = *(*int32)(unsafe.Add(mBase, _c_F_dualChannelSetupMainConnForPsync[2]))
	if v187 == int32(0) {
		goto L59
	} else {
		goto L60
	}
L7:
	;
	v174 = F_dualChannelReplMainConnRecvCapaReply(m, l0, v6+int32(60))
	mBase = m.M
	v175 = m.ExcPending
	if v175 != 0 {
		goto L53
	} else {
		goto L56
	}
L8:
	;
	v15 = v6 + int32(64)
	v18 = *(*int64)(unsafe.Add(mBase, _c_F_dualChannelSetupMainConnForPsync[3]))
	v19 = int32(0)
	v23 = int32(1)
	if base.Ui64(v18) < base.Ui64(int64(10)) {
		v80 = v23
		v81 = v19
		goto L10
	} else {
		goto L11
	}
L9:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v6)+44)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v6)+36)) = int32(_a_F_dualChannelSetupMainConnForPsync_2)
	*(*int32)(unsafe.Add(mBase, uint32(v6)+32)) = int32(_a_F_dualChannelSetupMainConnForPsync_3)
	*(*int32)(unsafe.Add(mBase, uint32(v6)+40)) = v6 + int32(64)
	v166 = F_sendCommand(m, l0, v6+int32(32))
	mBase = m.M
	v167 = m.ExcPending
	if v167 != 0 {
		goto L53
	} else {
		goto L54
	}
L10:
	;
	v84 = v80 + v81
	if base.Ui32(int32(21)) <= base.Ui32(v84) {
		goto L41
	} else {
		goto L42
	}
L11:
	;
	v31 = v19
	v32 = v18
	goto L12
L12:
	;
	if base.Ui64(int64(100)) <= base.Ui64(v32) {
		goto L14
	} else {
		goto L15
	}
L13:
	;
	v80 = v23
	v81 = v72
	goto L10
L14:
	;
	if base.Ui64(int64(1000)) <= base.Ui64(v32) {
		goto L16
	} else {
		goto L17
	}
L15:
	;
	v80 = int32(2)
	v81 = v31
	goto L10
L16:
	;
	if base.Ui64(int64(999999999999)) < base.Ui64(v32) {
		goto L18
	} else {
		goto L19
	}
L17:
	;
	v80 = int32(3)
	v81 = v31
	goto L10
L18:
	;
	v72 = v31 + int32(12)
	v76 = base.I64_div_u_s(v32, int64(1000000000000))
	if base.Ui64(int64(9999999999999)) < base.Ui64(v32) {
		v31 = v72
		v32 = v76
		goto L12
	} else {
		goto L40
	}
L19:
	;
	if base.Ui64(int64(99999999)) < base.Ui64(v32) {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	if base.Ui64(int64(9999999999)) < base.Ui64(v32) {
		goto L32
	} else {
		goto L33
	}
L21:
	;
	if base.Ui64(int64(999999)) < base.Ui64(v32) {
		goto L22
	} else {
		goto L23
	}
L22:
	;
	if base.Ui64(int64(9999999)) < base.Ui64(v32) {
		goto L29
	} else {
		goto L30
	}
L23:
	;
	if base.Ui64(int64(10000)) <= base.Ui64(v32) {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	if base.Ui64(int64(99999)) < base.Ui64(v32) {
		goto L26
	} else {
		goto L27
	}
L25:
	;
	v80 = int32(4)
	v81 = v31
	goto L10
L26:
	;
	v53 = int32(6)
	goto L28
L27:
	;
	v53 = int32(5)
	goto L28
L28:
	;
	v80 = v53
	v81 = v31
	goto L10
L29:
	;
	v58 = int32(8)
	goto L31
L30:
	;
	v58 = int32(7)
	goto L31
L31:
	;
	v80 = v58
	v81 = v31
	goto L10
L32:
	;
	if base.Ui64(int64(99999999999)) < base.Ui64(v32) {
		goto L37
	} else {
		goto L38
	}
L33:
	;
	if base.Ui64(int64(999999999)) < base.Ui64(v32) {
		goto L34
	} else {
		goto L35
	}
L34:
	;
	v65 = int32(10)
	goto L36
L35:
	;
	v65 = int32(9)
	goto L36
L36:
	;
	v80 = v65
	v81 = v31
	goto L10
L37:
	;
	v70 = int32(12)
	goto L39
L38:
	;
	v70 = int32(11)
	goto L39
L39:
	;
	v80 = v70
	v81 = v31
	goto L10
L40:
	;
	goto L13
L41:
	;
	goto L52
L42:
	;
	v87 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v15+v84))) = uint8(v87)
	v90 = v84 + int32(-1)
	if base.Ui64(int64(100)) <= base.Ui64(v18) {
		goto L44
	} else {
		goto L45
	}
L43:
	;
	v126 = v15 + v123
	if base.Ui64(int64(9)) < base.Ui64(v124) {
		goto L49
	} else {
		goto L50
	}
L44:
	;
	v97 = v18
	v99 = v90
	goto L46
L45:
	;
	v123 = v90
	v124 = v18
	goto L43
L46:
	;
	v103 = int64(100)
	v104 = base.I64_div_u_s(v97, v103)
	v113 = int32(*(*uint16)(unsafe.Add(mBase, uint32(base.I32_wrap_i64(v97-v104*v103)<<(uint(int32(1))%32))+uint32(_c_F_dualChannelSetupMainConnForPsync[4]))))
	*(*uint16)(unsafe.Add(mBase, uint32(v6+int32(63)+v99))) = uint16(v113)
	v116 = v99 + int32(-2)
	if base.Ui64(int64(9999)) < base.Ui64(v97) {
		v97 = v104
		v99 = v116
		goto L46
	} else {
		goto L48
	}
L47:
	;
	v123 = v116
	v124 = v104
	goto L43
L48:
	;
	goto L47
L49:
	;
	v140 = int32(*(*uint16)(unsafe.Add(mBase, uint32(base.I32_wrap_i64(v124)<<(uint(int32(1))%32))+uint32(_c_F_dualChannelSetupMainConnForPsync[4]))))
	*(*uint16)(unsafe.Add(mBase, uint32(v126+int32(-1)))) = uint16(v140)
	goto L9
L50:
	;
	v131 = base.I32_wrap_i64(v124) | int32(48)
	*(*uint8)(unsafe.Add(mBase, uint32(v126))) = uint8(v131)
	goto L9
L51:
	;
	goto L9
L52:
	;
	v144 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v15))) = uint8(v144)
	goto L51
L53:
	;
	return
L54:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v6)+60)) = v166
	if v166 != 0 {
		goto L2
	} else {
		goto L55
	}
L55:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_dualChannelSetupMainConnForPsync[0])) = int32(8)
	goto L1
L56:
	;
	if v174 == int32(-1) {
		goto L2
	} else {
		goto L57
	}
L57:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_dualChannelSetupMainConnForPsync[0])) = int32(11)
	v181 = *(*int32)(unsafe.Add(mBase, uint32(v6)+60))
	F_sdsfree(m, v181)
	mBase = m.M
	v183 = m.ExcPending
	if v183 != 0 {
		goto L53
	} else {
		goto L58
	}
L58:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v6)+60)) = int32(0)
	goto L6
L59:
	;
	v192 = F_replicaSendPsyncCommand(m, l0)
	mBase = m.M
	v193 = m.ExcPending
	if v193 != 0 {
		goto L53
	} else {
		goto L63
	}
L60:
	;
	F_debugPauseProcess(m)
	mBase = m.M
	v191 = m.ExcPending
	if v191 != 0 {
		goto L53
	} else {
		goto L61
	}
L61:
	;
	goto L59
L62:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_dualChannelSetupMainConnForPsync[0])) = int32(12)
	goto L1
L63:
	;
	if v192 != 0 {
		goto L62
	} else {
		goto L64
	}
L64:
	;
	v195 = *(*int32)(unsafe.Add(mBase, _c_F_dualChannelSetupMainConnForPsync[1]))
	if int32(3) < v195 {
		goto L65
	} else {
		goto L66
	}
L65:
	;
	v203 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v204 = *(*int32)(unsafe.Add(mBase, uint32(v203)+88))
	v205 = m.T0[v204].(func(*base.Module, int32) int32)(m, l0)
	mBase = m.M
	v206 = m.ExcPending
	if v206 != 0 {
		goto L53
	} else {
		goto L68
	}
L66:
	;
	F__serverLog(m, int32(3), int32(_a_F_dualChannelSetupMainConnForPsync_4), int32(0))
	mBase = m.M
	v202 = m.ExcPending
	if v202 != 0 {
		goto L53
	} else {
		goto L67
	}
L67:
	;
	goto L65
L68:
	;
	v207 = F_sdsnew(m, v205)
	mBase = m.M
	v208 = m.ExcPending
	if v208 != 0 {
		goto L53
	} else {
		goto L69
	}
L69:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v6)+60)) = v207
	goto L2
L70:
	;
	if v215 != 0 {
		goto L3
	} else {
		goto L71
	}
L71:
	;
	v218 = *(*int32)(unsafe.Add(mBase, _c_F_dualChannelSetupMainConnForPsync[5]))
	if v218 == int32(0) {
		goto L3
	} else {
		goto L72
	}
L72:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_dualChannelSetupMainConnForPsync[0])) = int32(13)
	goto L1
L73:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L74:
	;
	goto L2
L75:
	;
	v252 = F_cancelReplicationHandshake(m, int32(1))
	mBase = m.M
	v253 = m.ExcPending
	if v253 != 0 {
		goto L53
	} else {
		goto L81
	}
L76:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v6)+16)) = int32(-1)
	v240 = *(*int32)(unsafe.Add(mBase, uint32(v6)+60))
	if v240 != 0 {
		goto L77
	} else {
		goto L78
	}
L77:
	;
	v242 = v240
	goto L79
L78:
	;
	v242 = int32(_a_F_dualChannelSetupMainConnForPsync_5)
	goto L79
L79:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v6)+20)) = v242
	F__serverLog(m, int32(3), int32(_a_F_dualChannelSetupMainConnForPsync_6), v6+int32(16))
	mBase = m.M
	v249 = m.ExcPending
	if v249 != 0 {
		goto L53
	} else {
		goto L80
	}
L80:
	;
	goto L75
L81:
	;
	goto L1
L82:
	;
	m.G0 = v6 + int32(96)
	return
}
func F_dualChannelSyncHandlePsync(m *base.Module) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
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
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v82 int32
	_ = v82
	var v88 int32
	_ = v88
	var v94 int32
	_ = v94
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	v9 = *(*int32)(unsafe.Add(mBase, _c_F_dualChannelSyncHandlePsync[0]))
	if v9 != int32(12) {
		F__serverAssert(m, int32(_a_F_dualChannelSyncHandlePsync_0), int32(_a_F_dualChannelSyncHandlePsync_1), int32(3420))
		mBase = m.M
		v82 = m.ExcPending
		if v82 != 0 {
			return int32(0)
		} else {
			F_abort(m)
			mBase = m.M
			base.Wasm_trap_unreachable()
			for {
			}
		}
	} else {
		v13 = *(*int32)(unsafe.Add(mBase, _c_F_dualChannelSyncHandlePsync[1]))
		if int32(5) < v13 {
			if v13 != int32(6) {
				F__serverAssert(m, int32(_a_F_dualChannelSyncHandlePsync_2), int32(_a_F_dualChannelSyncHandlePsync_1), int32(3431))
				mBase = m.M
				v94 = m.ExcPending
				if v94 != 0 {
					return int32(0)
				} else {
					F_abort(m)
					mBase = m.M
					base.Wasm_trap_unreachable()
					for {
					}
				}
			} else {
				v60 = int32(0)
				v62 = *(*int32)(unsafe.Add(mBase, _c_F_dualChannelSyncHandlePsync[2]))
				if v60 < v62 {
					F_dualChannelSyncSuccess(m)
					mBase = m.M
					v71 = m.ExcPending
					if v71 != 0 {
						return int32(0)
					} else {
						v72 = v60
						m.G0 = v6 + int32(16)
						return v72
					}
				} else {
					v65 = int32(0)
					F__serverLog(m, v65, int32(_a_F_dualChannelSyncHandlePsync_3), v65)
					mBase = m.M
					v69 = m.ExcPending
					if v69 != 0 {
						return int32(0)
					} else {
						F_dualChannelSyncSuccess(m)
						mBase = m.M
						v71 = m.ExcPending
						if v71 != 0 {
							return int32(0)
						} else {
							v72 = v60
							m.G0 = v6 + int32(16)
							return v72
						}
					}
				}
			}
		} else {
			v17 = *(*int32)(unsafe.Add(mBase, _c_F_dualChannelSyncHandlePsync[3]))
			v19 = *(*int32)(unsafe.Add(mBase, uint32(v17)))
			v20 = *(*int32)(unsafe.Add(mBase, uint32(v19)+84))
			v21 = m.T0[v20].(func(*base.Module, int32, int32) int32)(m, v17, int32(974))
			mBase = m.M
			v24 = m.ExcPending
			if v24 != 0 {
				return int32(0)
			} else {
				if v21 != int32(-1) {
					v44 = *(*int32)(unsafe.Add(mBase, _c_F_dualChannelSyncHandlePsync[4]))
					if v44 != 0 {
						F__serverAssert(m, int32(_a_F_dualChannelSyncHandlePsync_4), int32(_a_F_dualChannelSyncHandlePsync_1), int32(3254))
						mBase = m.M
						v88 = m.ExcPending
						if v88 != 0 {
							return int32(0)
						} else {
							F_abort(m)
							mBase = m.M
							base.Wasm_trap_unreachable()
							for {
							}
						}
					} else {
						v45 = int32(0)
						v46 = int32(_a_F_dualChannelSyncHandlePsync_5)
						*(*int32)(unsafe.Add(mBase, _c_F_dualChannelSyncHandlePsync[5])) = v45
						*(*int64)(unsafe.Add(mBase, _c_F_dualChannelSyncHandlePsync[6])) = int64(0)
						v53 = F_listCreate(m)
						mBase = m.M
						v54 = m.ExcPending
						if v54 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, _c_F_dualChannelSyncHandlePsync[4])) = v53
							*(*int32)(unsafe.Add(mBase, uint32(v53)+12)) = int32(102)
							v72 = v45
							m.G0 = v6 + int32(16)
							return v72
						}
					}
				} else {
					v28 = *(*int32)(unsafe.Add(mBase, _c_F_dualChannelSyncHandlePsync[2]))
					if int32(3) < v28 {
						v40 = F_cancelReplicationHandshake(m, int32(1))
						mBase = m.M
						v41 = m.ExcPending
						if v41 != 0 {
							return int32(0)
						} else {
							v72 = int32(-1)
							m.G0 = v6 + int32(16)
							return v72
						}
					} else {
						v32 = *(*int32)(unsafe.Add(mBase, _c_F_dualChannelSyncHandlePsync[7]))
						v33 = F___strerror_l(m, v32, v32)
						mBase = m.M
						*(*int32)(unsafe.Add(mBase, uint32(v6))) = v33
						F__serverLog(m, int32(3), int32(_a_F_dualChannelSyncHandlePsync_6), v6)
						mBase = m.M
						v38 = m.ExcPending
						if v38 != 0 {
							return int32(0)
						} else {
							v40 = F_cancelReplicationHandshake(m, int32(1))
							mBase = m.M
							v41 = m.ExcPending
							if v41 != 0 {
								return int32(0)
							} else {
								v72 = int32(-1)
								m.G0 = v6 + int32(16)
								return v72
							}
						}
					}
				}
			}
		}
	}
}
func F_dummy_3(m *base.Module) {
	return
}
func F_dup2(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	if l0 == l1 {
		v12 = F___wasi_fd_is_valid(m, l0)
		mBase = m.M
		if v12 != 0 {
			v25 = l0
		} else {
			v16 = int32(-8)
			if base.Ui32(v16) < base.Ui32(int32(-4095)) {
				v24 = v16
			} else {
				v19 = F___errno_location(m)
				mBase = m.M
				*(*int32)(unsafe.Add(mBase, uint32(v19))) = int32(0) - v16
				v24 = int32(-1)
			}
			v25 = v24
		}
	} else {
		for {
			v9 = m.Env.X__syscall_dup3(m, l0, l1, int32(0))
			mBase = m.M
			if v9 == int32(-10) {
				continue
			} else {
				break
			}
			break
		}
		v16 = v9
		if base.Ui32(v16) < base.Ui32(int32(-4095)) {
			v24 = v16
		} else {
			v19 = F___errno_location(m)
			mBase = m.M
			*(*int32)(unsafe.Add(mBase, uint32(v19))) = int32(0) - v16
			v24 = int32(-1)
		}
		v25 = v24
	}
	return v25
}
func F_dupClientReplyValue(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v5 = v3 + int32(16)
	v6 = F_valkey_malloc(m, v5)
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		if v5 == int32(0) {
			v13 = v6
		} else {
			v12 = F__emscripten_memcpy_bulkmem(m, v6, l0, v5)
			mBase = m.M
			v13 = v12
		}
		return v13
	}
}
