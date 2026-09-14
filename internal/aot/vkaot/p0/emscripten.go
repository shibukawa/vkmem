package p0

import (
	base "github.com/shibukawa/valkeymem/internal/aot/vkaot/base"
	"unsafe"
)

func F__emscripten_memcpy_bulkmem(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	base.MemoryCopy(m, l0, l1, l2)
	return l0
}
func F__emscripten_memset_bulkmem(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	base.MemoryFill(m, l0, l1, l2)
	return l0
}
func F__emscripten_stack_alloc(m *base.Module, l0 int32) int32 {
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	v4 = m.G0
	v7 = (v4 - l0) & int32(-16)
	m.G0 = v7
	return v7
}
func F__emscripten_stack_restore(m *base.Module, l0 int32) {
	m.G0 = l0
	return
}
func F__emscripten_timeout(m *base.Module, l0 int32, l1 float64) {
	mBase := m.M
	_ = mBase
	var v5 float64
	_ = v5
	var v11 int32
	_ = v11
	var v17 float64
	_ = v17
	var v20 float64
	_ = v20
	var v21 float64
	_ = v21
	var v22 float64
	_ = v22
	var v30 int64
	_ = v30
	var v32 int64
	_ = v32
	var v40 int64
	_ = v40
	var v42 int64
	_ = v42
	var v43 int64
	_ = v43
	var v48 float64
	_ = v48
	var v54 float64
	_ = v54
	var v59 int32
	_ = v59
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	v5 = float64(0)
	v11 = l0 << (uint(int32(3)) % 32)
	v17 = *(*float64)(unsafe.Add(mBase, uint32(v11)+uint32(_consts[1247])))
	if base.F64_eq(v17, v5) != 0 {
		*(*int64)(unsafe.Add(mBase, uint32(v11)+uint32(_consts[1246]))) = int64(0)
		v54 = v5
	} else {
		v20 = *(*float64)(unsafe.Add(mBase, uint32(v11)+uint32(_consts[1246])))
		v21 = base.F64_max(l1, v20)
		v22 = base.F64_sub(v21, v20)
		if base.F64_lt(v22, float64(1.8446744073709552e+19))&base.F64_ge(v22, float64(0)) == int32(0) {
			v32 = int64(0)
		} else {
			v30 = base.I64_trunc_f64_u(v22)
			v32 = v30
		}
		if base.F64_lt(v17, float64(1.8446744073709552e+19))&base.F64_ge(v17, float64(0)) == int32(0) {
			v42 = int64(0)
		} else {
			v40 = base.I64_trunc_f64_u(v17)
			v42 = v40
		}
		v43 = base.I64_div_u_s(v32, v42)
		v48 = base.F64_add(base.F64_mul(base.F64_convert_i64_u(v43+int64(1)), v17), v20)
		*(*float64)(unsafe.Add(mBase, uint32(v11)+uint32(_consts[1246]))) = v48
		v54 = base.F64_sub(v48, v21)
	}
	v59 = m.Env.X_setitimer_js(m, l0, v54)
	mBase = m.M
	if l0 == int32(1) {
		v65 = int32(26)
	} else {
		v65 = int32(14)
	}
	if l0 == int32(2) {
		v68 = int32(27)
	} else {
		v68 = v65
	}
	v69 = F_raise(m, v68)
	mBase = m.M
	v70 = m.ExcPending
	if v70 != 0 {
		return
	} else {
		return
	}
}
func F_emscripten_builtin_calloc(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int64
	_ = v8
	var v9 int32
	_ = v9
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v27 int32
	_ = v27
	var v34 int32
	_ = v34
	if l0 != 0 {
		v8 = base.I64_extend_i32_u(l0) * base.I64_extend_i32_u(l1)
		v9 = base.I32_wrap_i64(v8)
		if base.Ui32(l1|l0) < base.Ui32(int32(65536)) {
			v20 = v9
		} else {
			if base.I32_wrap_i64(int64(base.Ui64(v8)>>(uint(int64(32))%64))) != int32(0) {
				v19 = int32(-1)
			} else {
				v19 = v9
			}
			v20 = v19
		}
	} else {
		v20 = int32(0)
	}
	v22 = F_emscripten_builtin_malloc(m, v20)
	mBase = m.M
	if v22 == int32(0) {
	} else {
		v27 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22+int32(-4)))))
		if v27&int32(3) == int32(0) {
		} else {
			v34 = F__emscripten_memset_bulkmem(m, v22, base.I32_extend8_s(int32(0)), v20)
			mBase = m.M
		}
	}
	return v22
}
func F_emscripten_builtin_realloc(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v102 int32
	_ = v102
	var v107 int32
	_ = v107
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v118 int32
	_ = v118
	var __phi118 int32
	_ = __phi118
	var v119 int32
	_ = v119
	var __phi119 int32
	_ = __phi119
	var v125 int32
	_ = v125
	var v128 int32
	_ = v128
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v150 int32
	_ = v150
	var v159 int32
	_ = v159
	var v161 int32
	_ = v161
	var v164 int32
	_ = v164
	var v167 int32
	_ = v167
	var v169 int32
	_ = v169
	var v174 int32
	_ = v174
	var v181 int32
	_ = v181
	var v186 int32
	_ = v186
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v201 int32
	_ = v201
	var v209 int32
	_ = v209
	var v211 int32
	_ = v211
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v222 int32
	_ = v222
	var v224 int32
	_ = v224
	var v231 int32
	_ = v231
	var v233 int32
	_ = v233
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v247 int32
	_ = v247
	var v248 int32
	_ = v248
	var v251 int32
	_ = v251
	var v253 int32
	_ = v253
	var v255 int32
	_ = v255
	var v264 int32
	_ = v264
	var v266 int32
	_ = v266
	var v269 int32
	_ = v269
	var v274 int32
	_ = v274
	var v279 int32
	_ = v279
	var v280 int32
	_ = v280
	var v285 int32
	_ = v285
	var __phi285 int32
	_ = __phi285
	var v286 int32
	_ = v286
	var __phi286 int32
	_ = __phi286
	var v292 int32
	_ = v292
	var v295 int32
	_ = v295
	var v309 int32
	_ = v309
	var v318 int32
	_ = v318
	var v320 int32
	_ = v320
	var v323 int32
	_ = v323
	var v326 int32
	_ = v326
	var v328 int32
	_ = v328
	var v333 int32
	_ = v333
	var v340 int32
	_ = v340
	var v345 int32
	_ = v345
	var v367 int32
	_ = v367
	var v379 int32
	_ = v379
	var v381 int32
	_ = v381
	var v383 int32
	_ = v383
	var v387 int32
	_ = v387
	var v392 int32
	_ = v392
	var v393 int32
	_ = v393
	var v404 int32
	_ = v404
	var v407 int32
	_ = v407
	var v414 int32
	_ = v414
	var v419 int32
	_ = v419
	var v423 int32
	_ = v423
	var v425 int32
	_ = v425
	var v440 int32
	_ = v440
	var v442 int32
	_ = v442
	var v445 int32
	_ = v445
	var v448 int32
	_ = v448
	var v452 int32
	_ = v452
	var v462 int32
	_ = v462
	var v463 int32
	_ = v463
	var v469 int32
	_ = v469
	var v471 int32
	_ = v471
	var v474 int32
	_ = v474
	var v478 int32
	_ = v478
	var v484 int32
	_ = v484
	var v486 int32
	_ = v486
	var v487 int32
	_ = v487
	var v488 int32
	_ = v488
	var v489 int32
	_ = v489
	var v498 int32
	_ = v498
	var v500 int32
	_ = v500
	var v501 int32
	_ = v501
	var v502 int32
	_ = v502
	var v504 int32
	_ = v504
	if l0 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	if base.Ui32(l1) < base.Ui32(int32(-64)) {
		goto L3
	} else {
		goto L4
	}
L2:
	;
	v5 = F_emscripten_builtin_malloc(m, l1)
	mBase = m.M
	return v5
L3:
	;
	v14 = int32(-8)
	v17 = int32(11)
	if base.Ui32(l1) < base.Ui32(v17) {
		goto L7
	} else {
		goto L8
	}
L4:
	;
	goto L5
L5:
	;
	*(*int32)(unsafe.Add(mBase, _consts[9])) = int32(48)
	return int32(0)
L6:
	;
	v30 = F_emscripten_builtin_malloc(m, l1)
	mBase = m.M
	if v30 != 0 {
		goto L11
	} else {
		goto L12
	}
L7:
	;
	v23 = int32(16)
	goto L9
L8:
	;
	v23 = (l1 + v17) & v14
	goto L9
L9:
	;
	v24 = F_try_realloc_chunk(m, l0+v14, v23)
	mBase = m.M
	if v24 == int32(0) {
		goto L6
	} else {
		goto L10
	}
L10:
	;
	return v24 + int32(8)
L11:
	;
	v33 = int32(-4)
	v37 = *(*int32)(unsafe.Add(mBase, uint32(l0+v33)))
	if v37&int32(3) != 0 {
		goto L13
	} else {
		goto L14
	}
L12:
	;
	return int32(0)
L13:
	;
	v40 = v33
	goto L15
L14:
	;
	v40 = int32(-8)
	goto L15
L15:
	;
	v43 = v40 + v37&int32(-8)
	if base.Ui32(v43) < base.Ui32(l1) {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	v45 = v43
	goto L18
L17:
	;
	v45 = l1
	goto L18
L18:
	;
	if v45 == int32(0) {
		goto L20
	} else {
		goto L21
	}
L19:
	;
	if l0 == int32(0) {
		goto L23
	} else {
		goto L24
	}
L20:
	;
	goto L19
L21:
	;
	v48 = F__emscripten_memcpy_bulkmem(m, v30, l0, v45)
	mBase = m.M
	goto L20
L22:
	;
	return v30
L23:
	;
	goto L22
L24:
	;
	v60 = int32(-8)
	v61 = l0 + v60
	v64 = *(*int32)(unsafe.Add(mBase, uint32(l0+int32(-4))))
	v66 = v64 & v60
	v67 = v61 + v66
	if v64&int32(1) != 0 {
		v191 = v66
		v192 = v61
		goto L25
	} else {
		goto L26
	}
L25:
	;
	if base.Ui32(v67) <= base.Ui32(v192) {
		goto L23
	} else {
		goto L60
	}
L26:
	;
	if v64&int32(2) == int32(0) {
		goto L23
	} else {
		goto L27
	}
L27:
	;
	v74 = *(*int32)(unsafe.Add(mBase, uint32(v61)))
	v75 = v61 - v74
	v77 = *(*int32)(unsafe.Add(mBase, _consts[1207]))
	if base.Ui32(v75) < base.Ui32(v77) {
		goto L23
	} else {
		goto L28
	}
L28:
	;
	v79 = v74 + v66
	v81 = *(*int32)(unsafe.Add(mBase, _consts[1208]))
	if v75 == v81 {
		goto L32
	} else {
		goto L33
	}
L29:
	;
	if v97 == int32(0) {
		v191 = v79
		v192 = v75
		goto L25
	} else {
		goto L48
	}
L30:
	;
	v150 = int32(0)
	goto L29
L31:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v86)+12)) = v83
	*(*int32)(unsafe.Add(mBase, uint32(v83)+8)) = v86
	v191 = v79
	v192 = v75
	goto L25
L32:
	;
	v131 = *(*int32)(unsafe.Add(mBase, uint32(v67)+4))
	v132 = int32(3)
	if v131&v132 != v132 {
		v191 = v79
		v192 = v75
		goto L25
	} else {
		goto L47
	}
L33:
	;
	v83 = *(*int32)(unsafe.Add(mBase, uint32(v75)+12))
	if base.Ui32(int32(255)) < base.Ui32(v74) {
		goto L34
	} else {
		goto L35
	}
L34:
	;
	v97 = *(*int32)(unsafe.Add(mBase, uint32(v75)+24))
	if v83 == v75 {
		goto L37
	} else {
		goto L38
	}
L35:
	;
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v75)+8))
	if v83 != v86 {
		goto L31
	} else {
		goto L36
	}
L36:
	;
	v88 = int32(0)
	v90 = *(*int32)(unsafe.Add(mBase, _consts[1209]))
	*(*int32)(unsafe.Add(mBase, _consts[1209])) = v90 & base.I32_rotl(int32(-2), int32(base.Ui32(v74)>>(uint(int32(3))%32)))
	v191 = v79
	v192 = v75
	goto L25
L37:
	;
	v102 = *(*int32)(unsafe.Add(mBase, uint32(v75)+20))
	if v102 == int32(0) {
		goto L40
	} else {
		goto L41
	}
L38:
	;
	v99 = *(*int32)(unsafe.Add(mBase, uint32(v75)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v99)+12)) = v83
	*(*int32)(unsafe.Add(mBase, uint32(v83)+8)) = v99
	v150 = v83
	goto L29
L39:
	;
	__phi118 = v112
	__phi119 = v113
	v118 = __phi118
	v119 = __phi119
	goto L43
L40:
	;
	v107 = *(*int32)(unsafe.Add(mBase, uint32(v75)+16))
	if v107 == int32(0) {
		goto L30
	} else {
		goto L42
	}
L41:
	;
	v112 = v102
	v113 = v75 + int32(20)
	goto L39
L42:
	;
	v112 = v107
	v113 = v75 + int32(16)
	goto L39
L43:
	;
	v125 = *(*int32)(unsafe.Add(mBase, uint32(v118)+20))
	if v125 != 0 {
		__phi118 = v125
		__phi119 = v118 + int32(20)
		v118 = __phi118
		v119 = __phi119
		goto L43
	} else {
		goto L45
	}
L44:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v119))) = int32(0)
	v150 = v118
	goto L29
L45:
	;
	v128 = *(*int32)(unsafe.Add(mBase, uint32(v118)+16))
	if v128 != 0 {
		__phi118 = v128
		__phi119 = v118 + int32(16)
		v118 = __phi118
		v119 = __phi119
		goto L43
	} else {
		goto L46
	}
L46:
	;
	goto L44
L47:
	;
	*(*int32)(unsafe.Add(mBase, _consts[1210])) = v79
	*(*int32)(unsafe.Add(mBase, uint32(v67)+4)) = v131 & int32(-2)
	*(*int32)(unsafe.Add(mBase, uint32(v75)+4)) = v79 | int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v67))) = v79
	goto L22
L48:
	;
	v159 = *(*int32)(unsafe.Add(mBase, uint32(v75)+28))
	v161 = v159 << (uint(int32(2)) % 32)
	v164 = *(*int32)(unsafe.Add(mBase, uint32(v161)+uint32(_consts[1211])))
	if v75 != v164 {
		goto L50
	} else {
		goto L51
	}
L49:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v150)+24)) = v97
	v181 = *(*int32)(unsafe.Add(mBase, uint32(v75)+16))
	if v181 == int32(0) {
		goto L57
	} else {
		goto L58
	}
L50:
	;
	v174 = *(*int32)(unsafe.Add(mBase, uint32(v97)+16))
	if v174 != v75 {
		goto L54
	} else {
		goto L55
	}
L51:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v161)+uint32(_consts[1211]))) = v150
	if v150 != 0 {
		goto L49
	} else {
		goto L52
	}
L52:
	;
	v167 = int32(0)
	v169 = *(*int32)(unsafe.Add(mBase, _consts[1212]))
	*(*int32)(unsafe.Add(mBase, _consts[1212])) = v169 & base.I32_rotl(int32(-2), v159)
	v191 = v79
	v192 = v75
	goto L25
L53:
	;
	if v150 == int32(0) {
		v191 = v79
		v192 = v75
		goto L25
	} else {
		goto L56
	}
L54:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v97)+20)) = v150
	goto L53
L55:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v97)+16)) = v150
	goto L53
L56:
	;
	goto L49
L57:
	;
	v186 = *(*int32)(unsafe.Add(mBase, uint32(v75)+20))
	if v186 == int32(0) {
		v191 = v79
		v192 = v75
		goto L25
	} else {
		goto L59
	}
L58:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v150)+16)) = v181
	*(*int32)(unsafe.Add(mBase, uint32(v181)+24)) = v150
	goto L57
L59:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v150)+20)) = v186
	*(*int32)(unsafe.Add(mBase, uint32(v186)+24)) = v150
	v191 = v79
	v192 = v75
	goto L25
L60:
	;
	v201 = *(*int32)(unsafe.Add(mBase, uint32(v67)+4))
	if v201&int32(1) == int32(0) {
		goto L23
	} else {
		goto L61
	}
L61:
	;
	if v201&int32(2) != 0 {
		goto L66
	} else {
		goto L67
	}
L62:
	;
	if base.Ui32(int32(255)) < base.Ui32(v367) {
		goto L100
	} else {
		goto L101
	}
L63:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v192)+4)) = v247 | int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v192+v247))) = v247
	if v192 != v231 {
		v367 = v247
		goto L62
	} else {
		goto L99
	}
L64:
	;
	if v264 == int32(0) {
		goto L63
	} else {
		goto L87
	}
L65:
	;
	v309 = int32(0)
	goto L64
L66:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v67)+4)) = v201 & int32(-2)
	*(*int32)(unsafe.Add(mBase, uint32(v192)+4)) = v191 | int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v192+v191))) = v191
	v367 = v191
	goto L62
L67:
	;
	v209 = *(*int32)(unsafe.Add(mBase, _consts[1213]))
	if v67 != v209 {
		goto L68
	} else {
		goto L69
	}
L68:
	;
	v231 = *(*int32)(unsafe.Add(mBase, _consts[1208]))
	if v67 != v231 {
		goto L71
	} else {
		goto L72
	}
L69:
	;
	v211 = int32(0)
	*(*int32)(unsafe.Add(mBase, _consts[1213])) = v192
	v215 = *(*int32)(unsafe.Add(mBase, _consts[1214]))
	v216 = v215 + v191
	*(*int32)(unsafe.Add(mBase, _consts[1214])) = v216
	*(*int32)(unsafe.Add(mBase, uint32(v192)+4)) = v216 | int32(1)
	v222 = *(*int32)(unsafe.Add(mBase, _consts[1208]))
	if v192 != v222 {
		goto L23
	} else {
		goto L70
	}
L70:
	;
	v224 = int32(0)
	*(*int32)(unsafe.Add(mBase, _consts[1210])) = v224
	*(*int32)(unsafe.Add(mBase, _consts[1208])) = v224
	goto L22
L71:
	;
	v247 = v201&int32(-8) + v191
	v248 = *(*int32)(unsafe.Add(mBase, uint32(v67)+12))
	if base.Ui32(int32(255)) < base.Ui32(v201) {
		goto L73
	} else {
		goto L74
	}
L72:
	;
	v233 = int32(0)
	*(*int32)(unsafe.Add(mBase, _consts[1208])) = v192
	v237 = *(*int32)(unsafe.Add(mBase, _consts[1210]))
	v238 = v237 + v191
	*(*int32)(unsafe.Add(mBase, _consts[1210])) = v238
	*(*int32)(unsafe.Add(mBase, uint32(v192)+4)) = v238 | int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v192+v238))) = v238
	goto L22
L73:
	;
	v264 = *(*int32)(unsafe.Add(mBase, uint32(v67)+24))
	if v248 == v67 {
		goto L77
	} else {
		goto L78
	}
L74:
	;
	v251 = *(*int32)(unsafe.Add(mBase, uint32(v67)+8))
	if v248 != v251 {
		goto L75
	} else {
		goto L76
	}
L75:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v251)+12)) = v248
	*(*int32)(unsafe.Add(mBase, uint32(v248)+8)) = v251
	goto L63
L76:
	;
	v253 = int32(0)
	v255 = *(*int32)(unsafe.Add(mBase, _consts[1209]))
	*(*int32)(unsafe.Add(mBase, _consts[1209])) = v255 & base.I32_rotl(int32(-2), int32(base.Ui32(v201)>>(uint(int32(3))%32)))
	goto L63
L77:
	;
	v269 = *(*int32)(unsafe.Add(mBase, uint32(v67)+20))
	if v269 == int32(0) {
		goto L80
	} else {
		goto L81
	}
L78:
	;
	v266 = *(*int32)(unsafe.Add(mBase, uint32(v67)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v266)+12)) = v248
	*(*int32)(unsafe.Add(mBase, uint32(v248)+8)) = v266
	v309 = v248
	goto L64
L79:
	;
	__phi285 = v279
	__phi286 = v280
	v285 = __phi285
	v286 = __phi286
	goto L83
L80:
	;
	v274 = *(*int32)(unsafe.Add(mBase, uint32(v67)+16))
	if v274 == int32(0) {
		goto L65
	} else {
		goto L82
	}
L81:
	;
	v279 = v269
	v280 = v67 + int32(20)
	goto L79
L82:
	;
	v279 = v274
	v280 = v67 + int32(16)
	goto L79
L83:
	;
	v292 = *(*int32)(unsafe.Add(mBase, uint32(v285)+20))
	if v292 != 0 {
		__phi285 = v292
		__phi286 = v285 + int32(20)
		v285 = __phi285
		v286 = __phi286
		goto L83
	} else {
		goto L85
	}
L84:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v286))) = int32(0)
	v309 = v285
	goto L64
L85:
	;
	v295 = *(*int32)(unsafe.Add(mBase, uint32(v285)+16))
	if v295 != 0 {
		__phi285 = v295
		__phi286 = v285 + int32(16)
		v285 = __phi285
		v286 = __phi286
		goto L83
	} else {
		goto L86
	}
L86:
	;
	goto L84
L87:
	;
	v318 = *(*int32)(unsafe.Add(mBase, uint32(v67)+28))
	v320 = v318 << (uint(int32(2)) % 32)
	v323 = *(*int32)(unsafe.Add(mBase, uint32(v320)+uint32(_consts[1211])))
	if v67 != v323 {
		goto L89
	} else {
		goto L90
	}
L88:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v309)+24)) = v264
	v340 = *(*int32)(unsafe.Add(mBase, uint32(v67)+16))
	if v340 == int32(0) {
		goto L96
	} else {
		goto L97
	}
L89:
	;
	v333 = *(*int32)(unsafe.Add(mBase, uint32(v264)+16))
	if v333 != v67 {
		goto L93
	} else {
		goto L94
	}
L90:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v320)+uint32(_consts[1211]))) = v309
	if v309 != 0 {
		goto L88
	} else {
		goto L91
	}
L91:
	;
	v326 = int32(0)
	v328 = *(*int32)(unsafe.Add(mBase, _consts[1212]))
	*(*int32)(unsafe.Add(mBase, _consts[1212])) = v328 & base.I32_rotl(int32(-2), v318)
	goto L63
L92:
	;
	if v309 == int32(0) {
		goto L63
	} else {
		goto L95
	}
L93:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v264)+20)) = v309
	goto L92
L94:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v264)+16)) = v309
	goto L92
L95:
	;
	goto L88
L96:
	;
	v345 = *(*int32)(unsafe.Add(mBase, uint32(v67)+20))
	if v345 == int32(0) {
		goto L63
	} else {
		goto L98
	}
L97:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v309)+16)) = v340
	*(*int32)(unsafe.Add(mBase, uint32(v340)+24)) = v309
	goto L96
L98:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v309)+20)) = v345
	*(*int32)(unsafe.Add(mBase, uint32(v345)+24)) = v309
	goto L63
L99:
	;
	*(*int32)(unsafe.Add(mBase, _consts[1210])) = v247
	goto L22
L100:
	;
	if base.Ui32(int32(16777215)) < base.Ui32(v367) {
		v414 = int32(31)
		goto L105
	} else {
		goto L106
	}
L101:
	;
	v379 = v367 & int32(-8)
	v381 = v379 + int32(9128464)
	v383 = *(*int32)(unsafe.Add(mBase, _consts[1209]))
	v387 = int32(1) << (uint(int32(base.Ui32(v367)>>(uint(int32(3))%32))) % 32)
	if v383&v387 != 0 {
		goto L103
	} else {
		goto L104
	}
L102:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v379)+uint32(_consts[1215]))) = v192
	*(*int32)(unsafe.Add(mBase, uint32(v393)+12)) = v192
	*(*int32)(unsafe.Add(mBase, uint32(v192)+12)) = v381
	*(*int32)(unsafe.Add(mBase, uint32(v192)+8)) = v393
	goto L22
L103:
	;
	v392 = *(*int32)(unsafe.Add(mBase, uint32(v379)+uint32(_consts[1215])))
	v393 = v392
	goto L102
L104:
	;
	*(*int32)(unsafe.Add(mBase, _consts[1209])) = v383 | v387
	v393 = v381
	goto L102
L105:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v192)+28)) = v414
	*(*int64)(unsafe.Add(mBase, uint32(v192)+16)) = int64(0)
	v419 = v414 << (uint(int32(2)) % 32)
	v423 = *(*int32)(unsafe.Add(mBase, _consts[1212]))
	v425 = int32(1) << (uint(v414) % 32)
	if v423&v425 != 0 {
		goto L110
	} else {
		goto L111
	}
L106:
	;
	v404 = base.I32_clz(int32(base.Ui32(v367) >> (uint(int32(8)) % 32)))
	v407 = int32(1)
	v414 = int32(base.Ui32(v367)>>(uint(int32(38)-v404)%32))&v407 - v404<<(uint(v407)%32) + int32(62)
	goto L105
L107:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v192+v486))) = v489
	*(*int32)(unsafe.Add(mBase, uint32(v192)+12)) = v488
	*(*int32)(unsafe.Add(mBase, uint32(v192+v484))) = v487
	v498 = int32(0)
	v500 = *(*int32)(unsafe.Add(mBase, _consts[1216]))
	v501 = int32(-1)
	v502 = v500 + v501
	if v502 != 0 {
		goto L119
	} else {
		goto L120
	}
L108:
	;
	v478 = *(*int32)(unsafe.Add(mBase, uint32(v448)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v478)+12)) = v192
	*(*int32)(unsafe.Add(mBase, uint32(v448)+8)) = v192
	v484 = int32(24)
	v486 = int32(8)
	v487 = int32(0)
	v488 = v448
	v489 = v478
	goto L107
L109:
	;
	v484 = v469
	v486 = v471
	v487 = v192
	v488 = v192
	v489 = v474
	goto L107
L110:
	;
	if v414 == int32(31) {
		goto L112
	} else {
		goto L113
	}
L111:
	;
	*(*int32)(unsafe.Add(mBase, _consts[1212])) = v423 | v425
	*(*int32)(unsafe.Add(mBase, uint32(v419)+uint32(_consts[1211]))) = v192
	v469 = int32(8)
	v471 = int32(24)
	v474 = v419 + int32(9128728)
	goto L109
L112:
	;
	v440 = int32(0)
	goto L114
L113:
	;
	v440 = int32(25) - int32(base.Ui32(v414)>>(uint(int32(1))%32))
	goto L114
L114:
	;
	v442 = *(*int32)(unsafe.Add(mBase, uint32(v419)+uint32(_consts[1211])))
	v445 = v367 << (uint(v440) % 32)
	v448 = v442
	goto L115
L115:
	;
	v452 = *(*int32)(unsafe.Add(mBase, uint32(v448)+4))
	if v452&int32(-8) == v367 {
		goto L108
	} else {
		goto L117
	}
L116:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v462+int32(16)))) = v192
	v469 = int32(8)
	v471 = int32(24)
	v474 = v448
	goto L109
L117:
	;
	v462 = v448 + int32(base.Ui32(v445)>>(uint(int32(29))%32))&int32(4)
	v463 = *(*int32)(unsafe.Add(mBase, uint32(v462)+16))
	if v463 != 0 {
		v445 = v445 << (uint(int32(1)) % 32)
		v448 = v463
		goto L115
	} else {
		goto L118
	}
L118:
	;
	goto L116
L119:
	;
	v504 = v502
	goto L121
L120:
	;
	v504 = v501
	goto L121
L121:
	;
	*(*int32)(unsafe.Add(mBase, _consts[1216])) = v504
	goto L23
}
func F_emscripten_stack_get_base(m *base.Module) int32 {
	var v1 int32
	_ = v1
	v1 = m.G2
	return v1
}
