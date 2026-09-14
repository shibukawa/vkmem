package p1

import (
	base "github.com/shibukawa/vkmem/internal/aot/vkaot/base"
	"unsafe"
)

func F__emscripten_check_timers(m *base.Module, l0 float64) {
	mBase := m.M
	_ = mBase
	var v6 float64
	_ = v6
	var v7 int32
	_ = v7
	var v11 int32
	_ = v11
	var v14 float64
	_ = v14
	var v19 float64
	_ = v19
	var v20 float64
	_ = v20
	var v21 float64
	_ = v21
	var v22 float64
	_ = v22
	var v27 int32
	_ = v27
	var v28 float64
	_ = v28
	var v31 int32
	_ = v31
	v6 = l0
	v7 = int32(0)
	goto L1
L1:
	;
	v11 = v7 << (uint(int32(3)) % 32)
	v14 = *(*float64)(unsafe.Add(mBase, uint32(v11)+uint32(_consts[1050])))
	if base.F64_eq(v14, float64(0)) != 0 {
		v28 = v6
		goto L3
	} else {
		goto L4
	}
L2:
	;
	return
L3:
	;
	v31 = v7 + int32(1)
	if v31 != int32(3) {
		v6 = v28
		v7 = v31
		goto L1
	} else {
		goto L10
	}
L4:
	;
	if base.F64_ne(v6, float64(0)) != 0 {
		v21 = v6
		v22 = v14
		goto L5
	} else {
		goto L6
	}
L5:
	;
	if base.F64_ge(v21, v22) == int32(0) {
		v28 = v21
		goto L3
	} else {
		goto L7
	}
L6:
	;
	v19 = m.Env.Emscripten_get_now(m)
	mBase = m.M
	v20 = *(*float64)(unsafe.Add(mBase, uint32(v11)+uint32(_consts[1050])))
	v21 = v19
	v22 = v20
	goto L5
L7:
	;
	F__emscripten_timeout(m, v7, v21)
	mBase = m.M
	v27 = m.ExcPending
	if v27 != 0 {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	return
L9:
	;
	v28 = v21
	goto L3
L10:
	;
	goto L2
}
func F_emscripten_builtin_free(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v59 int32
	_ = v59
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v70 int32
	_ = v70
	var __phi70 int32
	_ = __phi70
	var v71 int32
	_ = v71
	var __phi71 int32
	_ = __phi71
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v102 int32
	_ = v102
	var v111 int32
	_ = v111
	var v113 int32
	_ = v113
	var v116 int32
	_ = v116
	var v119 int32
	_ = v119
	var v121 int32
	_ = v121
	var v126 int32
	_ = v126
	var v133 int32
	_ = v133
	var v138 int32
	_ = v138
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v153 int32
	_ = v153
	var v161 int32
	_ = v161
	var v163 int32
	_ = v163
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v174 int32
	_ = v174
	var v176 int32
	_ = v176
	var v183 int32
	_ = v183
	var v185 int32
	_ = v185
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v203 int32
	_ = v203
	var v205 int32
	_ = v205
	var v207 int32
	_ = v207
	var v216 int32
	_ = v216
	var v218 int32
	_ = v218
	var v221 int32
	_ = v221
	var v226 int32
	_ = v226
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v237 int32
	_ = v237
	var __phi237 int32
	_ = __phi237
	var v238 int32
	_ = v238
	var __phi238 int32
	_ = __phi238
	var v244 int32
	_ = v244
	var v247 int32
	_ = v247
	var v261 int32
	_ = v261
	var v270 int32
	_ = v270
	var v272 int32
	_ = v272
	var v275 int32
	_ = v275
	var v278 int32
	_ = v278
	var v280 int32
	_ = v280
	var v285 int32
	_ = v285
	var v292 int32
	_ = v292
	var v297 int32
	_ = v297
	var v319 int32
	_ = v319
	var v331 int32
	_ = v331
	var v333 int32
	_ = v333
	var v335 int32
	_ = v335
	var v339 int32
	_ = v339
	var v344 int32
	_ = v344
	var v345 int32
	_ = v345
	var v356 int32
	_ = v356
	var v359 int32
	_ = v359
	var v366 int32
	_ = v366
	var v371 int32
	_ = v371
	var v375 int32
	_ = v375
	var v377 int32
	_ = v377
	var v392 int32
	_ = v392
	var v394 int32
	_ = v394
	var v397 int32
	_ = v397
	var v400 int32
	_ = v400
	var v404 int32
	_ = v404
	var v414 int32
	_ = v414
	var v415 int32
	_ = v415
	var v421 int32
	_ = v421
	var v423 int32
	_ = v423
	var v426 int32
	_ = v426
	var v430 int32
	_ = v430
	var v436 int32
	_ = v436
	var v438 int32
	_ = v438
	var v439 int32
	_ = v439
	var v440 int32
	_ = v440
	var v441 int32
	_ = v441
	var v450 int32
	_ = v450
	var v452 int32
	_ = v452
	var v453 int32
	_ = v453
	var v454 int32
	_ = v454
	var v456 int32
	_ = v456
	if l0 == int32(0) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v12 = int32(-8)
	v13 = l0 + v12
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0+int32(-4))))
	v18 = v16 & v12
	v19 = v13 + v18
	if v16&int32(1) != 0 {
		v143 = v18
		v144 = v13
		goto L3
	} else {
		goto L4
	}
L3:
	;
	if base.Ui32(v19) <= base.Ui32(v144) {
		goto L1
	} else {
		goto L38
	}
L4:
	;
	if v16&int32(2) == int32(0) {
		goto L1
	} else {
		goto L5
	}
L5:
	;
	v26 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
	v27 = v13 - v26
	v29 = *(*int32)(unsafe.Add(mBase, _consts[994]))
	if base.Ui32(v27) < base.Ui32(v29) {
		goto L1
	} else {
		goto L6
	}
L6:
	;
	v31 = v26 + v18
	v33 = *(*int32)(unsafe.Add(mBase, _consts[995]))
	if v27 == v33 {
		goto L10
	} else {
		goto L11
	}
L7:
	;
	if v49 == int32(0) {
		v143 = v31
		v144 = v27
		goto L3
	} else {
		goto L26
	}
L8:
	;
	v102 = int32(0)
	goto L7
L9:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v38)+12)) = v35
	*(*int32)(unsafe.Add(mBase, uint32(v35)+8)) = v38
	v143 = v31
	v144 = v27
	goto L3
L10:
	;
	v83 = *(*int32)(unsafe.Add(mBase, uint32(v19)+4))
	v84 = int32(3)
	if v83&v84 != v84 {
		v143 = v31
		v144 = v27
		goto L3
	} else {
		goto L25
	}
L11:
	;
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v27)+12))
	if base.Ui32(int32(255)) < base.Ui32(v26) {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v27)+24))
	if v35 == v27 {
		goto L15
	} else {
		goto L16
	}
L13:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v27)+8))
	if v35 != v38 {
		goto L9
	} else {
		goto L14
	}
L14:
	;
	v40 = int32(0)
	v42 = *(*int32)(unsafe.Add(mBase, _consts[996]))
	*(*int32)(unsafe.Add(mBase, _consts[996])) = v42 & base.I32_rotl(int32(-2), int32(base.Ui32(v26)>>(uint(int32(3))%32)))
	v143 = v31
	v144 = v27
	goto L3
L15:
	;
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v27)+20))
	if v54 == int32(0) {
		goto L18
	} else {
		goto L19
	}
L16:
	;
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v27)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v51)+12)) = v35
	*(*int32)(unsafe.Add(mBase, uint32(v35)+8)) = v51
	v102 = v35
	goto L7
L17:
	;
	__phi70 = v64
	__phi71 = v65
	v70 = __phi70
	v71 = __phi71
	goto L21
L18:
	;
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v27)+16))
	if v59 == int32(0) {
		goto L8
	} else {
		goto L20
	}
L19:
	;
	v64 = v54
	v65 = v27 + int32(20)
	goto L17
L20:
	;
	v64 = v59
	v65 = v27 + int32(16)
	goto L17
L21:
	;
	v77 = *(*int32)(unsafe.Add(mBase, uint32(v70)+20))
	if v77 != 0 {
		__phi70 = v77
		__phi71 = v70 + int32(20)
		v70 = __phi70
		v71 = __phi71
		goto L21
	} else {
		goto L23
	}
L22:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v71))) = int32(0)
	v102 = v70
	goto L7
L23:
	;
	v80 = *(*int32)(unsafe.Add(mBase, uint32(v70)+16))
	if v80 != 0 {
		__phi70 = v80
		__phi71 = v70 + int32(16)
		v70 = __phi70
		v71 = __phi71
		goto L21
	} else {
		goto L24
	}
L24:
	;
	goto L22
L25:
	;
	*(*int32)(unsafe.Add(mBase, _consts[997])) = v31
	*(*int32)(unsafe.Add(mBase, uint32(v19)+4)) = v83 & int32(-2)
	*(*int32)(unsafe.Add(mBase, uint32(v27)+4)) = v31 | int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v19))) = v31
	return
L26:
	;
	v111 = *(*int32)(unsafe.Add(mBase, uint32(v27)+28))
	v113 = v111 << (uint(int32(2)) % 32)
	v116 = *(*int32)(unsafe.Add(mBase, uint32(v113)+uint32(_consts[998])))
	if v27 != v116 {
		goto L28
	} else {
		goto L29
	}
L27:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v102)+24)) = v49
	v133 = *(*int32)(unsafe.Add(mBase, uint32(v27)+16))
	if v133 == int32(0) {
		goto L35
	} else {
		goto L36
	}
L28:
	;
	v126 = *(*int32)(unsafe.Add(mBase, uint32(v49)+16))
	if v126 != v27 {
		goto L32
	} else {
		goto L33
	}
L29:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v113)+uint32(_consts[998]))) = v102
	if v102 != 0 {
		goto L27
	} else {
		goto L30
	}
L30:
	;
	v119 = int32(0)
	v121 = *(*int32)(unsafe.Add(mBase, _consts[999]))
	*(*int32)(unsafe.Add(mBase, _consts[999])) = v121 & base.I32_rotl(int32(-2), v111)
	v143 = v31
	v144 = v27
	goto L3
L31:
	;
	if v102 == int32(0) {
		v143 = v31
		v144 = v27
		goto L3
	} else {
		goto L34
	}
L32:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v49)+20)) = v102
	goto L31
L33:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v49)+16)) = v102
	goto L31
L34:
	;
	goto L27
L35:
	;
	v138 = *(*int32)(unsafe.Add(mBase, uint32(v27)+20))
	if v138 == int32(0) {
		v143 = v31
		v144 = v27
		goto L3
	} else {
		goto L37
	}
L36:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v102)+16)) = v133
	*(*int32)(unsafe.Add(mBase, uint32(v133)+24)) = v102
	goto L35
L37:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v102)+20)) = v138
	*(*int32)(unsafe.Add(mBase, uint32(v138)+24)) = v102
	v143 = v31
	v144 = v27
	goto L3
L38:
	;
	v153 = *(*int32)(unsafe.Add(mBase, uint32(v19)+4))
	if v153&int32(1) == int32(0) {
		goto L1
	} else {
		goto L39
	}
L39:
	;
	if v153&int32(2) != 0 {
		goto L44
	} else {
		goto L45
	}
L40:
	;
	if base.Ui32(int32(255)) < base.Ui32(v319) {
		goto L78
	} else {
		goto L79
	}
L41:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v144)+4)) = v199 | int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v144+v199))) = v199
	if v144 != v183 {
		v319 = v199
		goto L40
	} else {
		goto L77
	}
L42:
	;
	if v216 == int32(0) {
		goto L41
	} else {
		goto L65
	}
L43:
	;
	v261 = int32(0)
	goto L42
L44:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+4)) = v153 & int32(-2)
	*(*int32)(unsafe.Add(mBase, uint32(v144)+4)) = v143 | int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v144+v143))) = v143
	v319 = v143
	goto L40
L45:
	;
	v161 = *(*int32)(unsafe.Add(mBase, _consts[1000]))
	if v19 != v161 {
		goto L46
	} else {
		goto L47
	}
L46:
	;
	v183 = *(*int32)(unsafe.Add(mBase, _consts[995]))
	if v19 != v183 {
		goto L49
	} else {
		goto L50
	}
L47:
	;
	v163 = int32(0)
	*(*int32)(unsafe.Add(mBase, _consts[1000])) = v144
	v167 = *(*int32)(unsafe.Add(mBase, _consts[1001]))
	v168 = v167 + v143
	*(*int32)(unsafe.Add(mBase, _consts[1001])) = v168
	*(*int32)(unsafe.Add(mBase, uint32(v144)+4)) = v168 | int32(1)
	v174 = *(*int32)(unsafe.Add(mBase, _consts[995]))
	if v144 != v174 {
		goto L1
	} else {
		goto L48
	}
L48:
	;
	v176 = int32(0)
	*(*int32)(unsafe.Add(mBase, _consts[997])) = v176
	*(*int32)(unsafe.Add(mBase, _consts[995])) = v176
	return
L49:
	;
	v199 = v153&int32(-8) + v143
	v200 = *(*int32)(unsafe.Add(mBase, uint32(v19)+12))
	if base.Ui32(int32(255)) < base.Ui32(v153) {
		goto L51
	} else {
		goto L52
	}
L50:
	;
	v185 = int32(0)
	*(*int32)(unsafe.Add(mBase, _consts[995])) = v144
	v189 = *(*int32)(unsafe.Add(mBase, _consts[997]))
	v190 = v189 + v143
	*(*int32)(unsafe.Add(mBase, _consts[997])) = v190
	*(*int32)(unsafe.Add(mBase, uint32(v144)+4)) = v190 | int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v144+v190))) = v190
	return
L51:
	;
	v216 = *(*int32)(unsafe.Add(mBase, uint32(v19)+24))
	if v200 == v19 {
		goto L55
	} else {
		goto L56
	}
L52:
	;
	v203 = *(*int32)(unsafe.Add(mBase, uint32(v19)+8))
	if v200 != v203 {
		goto L53
	} else {
		goto L54
	}
L53:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v203)+12)) = v200
	*(*int32)(unsafe.Add(mBase, uint32(v200)+8)) = v203
	goto L41
L54:
	;
	v205 = int32(0)
	v207 = *(*int32)(unsafe.Add(mBase, _consts[996]))
	*(*int32)(unsafe.Add(mBase, _consts[996])) = v207 & base.I32_rotl(int32(-2), int32(base.Ui32(v153)>>(uint(int32(3))%32)))
	goto L41
L55:
	;
	v221 = *(*int32)(unsafe.Add(mBase, uint32(v19)+20))
	if v221 == int32(0) {
		goto L58
	} else {
		goto L59
	}
L56:
	;
	v218 = *(*int32)(unsafe.Add(mBase, uint32(v19)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v218)+12)) = v200
	*(*int32)(unsafe.Add(mBase, uint32(v200)+8)) = v218
	v261 = v200
	goto L42
L57:
	;
	__phi237 = v231
	__phi238 = v232
	v237 = __phi237
	v238 = __phi238
	goto L61
L58:
	;
	v226 = *(*int32)(unsafe.Add(mBase, uint32(v19)+16))
	if v226 == int32(0) {
		goto L43
	} else {
		goto L60
	}
L59:
	;
	v231 = v221
	v232 = v19 + int32(20)
	goto L57
L60:
	;
	v231 = v226
	v232 = v19 + int32(16)
	goto L57
L61:
	;
	v244 = *(*int32)(unsafe.Add(mBase, uint32(v237)+20))
	if v244 != 0 {
		__phi237 = v244
		__phi238 = v237 + int32(20)
		v237 = __phi237
		v238 = __phi238
		goto L61
	} else {
		goto L63
	}
L62:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v238))) = int32(0)
	v261 = v237
	goto L42
L63:
	;
	v247 = *(*int32)(unsafe.Add(mBase, uint32(v237)+16))
	if v247 != 0 {
		__phi237 = v247
		__phi238 = v237 + int32(16)
		v237 = __phi237
		v238 = __phi238
		goto L61
	} else {
		goto L64
	}
L64:
	;
	goto L62
L65:
	;
	v270 = *(*int32)(unsafe.Add(mBase, uint32(v19)+28))
	v272 = v270 << (uint(int32(2)) % 32)
	v275 = *(*int32)(unsafe.Add(mBase, uint32(v272)+uint32(_consts[998])))
	if v19 != v275 {
		goto L67
	} else {
		goto L68
	}
L66:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v261)+24)) = v216
	v292 = *(*int32)(unsafe.Add(mBase, uint32(v19)+16))
	if v292 == int32(0) {
		goto L74
	} else {
		goto L75
	}
L67:
	;
	v285 = *(*int32)(unsafe.Add(mBase, uint32(v216)+16))
	if v285 != v19 {
		goto L71
	} else {
		goto L72
	}
L68:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v272)+uint32(_consts[998]))) = v261
	if v261 != 0 {
		goto L66
	} else {
		goto L69
	}
L69:
	;
	v278 = int32(0)
	v280 = *(*int32)(unsafe.Add(mBase, _consts[999]))
	*(*int32)(unsafe.Add(mBase, _consts[999])) = v280 & base.I32_rotl(int32(-2), v270)
	goto L41
L70:
	;
	if v261 == int32(0) {
		goto L41
	} else {
		goto L73
	}
L71:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v216)+20)) = v261
	goto L70
L72:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v216)+16)) = v261
	goto L70
L73:
	;
	goto L66
L74:
	;
	v297 = *(*int32)(unsafe.Add(mBase, uint32(v19)+20))
	if v297 == int32(0) {
		goto L41
	} else {
		goto L76
	}
L75:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v261)+16)) = v292
	*(*int32)(unsafe.Add(mBase, uint32(v292)+24)) = v261
	goto L74
L76:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v261)+20)) = v297
	*(*int32)(unsafe.Add(mBase, uint32(v297)+24)) = v261
	goto L41
L77:
	;
	*(*int32)(unsafe.Add(mBase, _consts[997])) = v199
	return
L78:
	;
	if base.Ui32(int32(16777215)) < base.Ui32(v319) {
		v366 = int32(31)
		goto L83
	} else {
		goto L84
	}
L79:
	;
	v331 = v319 & int32(-8)
	v333 = v331 + int32(9128464)
	v335 = *(*int32)(unsafe.Add(mBase, _consts[996]))
	v339 = int32(1) << (uint(int32(base.Ui32(v319)>>(uint(int32(3))%32))) % 32)
	if v335&v339 != 0 {
		goto L81
	} else {
		goto L82
	}
L80:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v331)+uint32(_consts[1002]))) = v144
	*(*int32)(unsafe.Add(mBase, uint32(v345)+12)) = v144
	*(*int32)(unsafe.Add(mBase, uint32(v144)+12)) = v333
	*(*int32)(unsafe.Add(mBase, uint32(v144)+8)) = v345
	return
L81:
	;
	v344 = *(*int32)(unsafe.Add(mBase, uint32(v331)+uint32(_consts[1002])))
	v345 = v344
	goto L80
L82:
	;
	*(*int32)(unsafe.Add(mBase, _consts[996])) = v335 | v339
	v345 = v333
	goto L80
L83:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v144)+28)) = v366
	*(*int64)(unsafe.Add(mBase, uint32(v144)+16)) = int64(0)
	v371 = v366 << (uint(int32(2)) % 32)
	v375 = *(*int32)(unsafe.Add(mBase, _consts[999]))
	v377 = int32(1) << (uint(v366) % 32)
	if v375&v377 != 0 {
		goto L88
	} else {
		goto L89
	}
L84:
	;
	v356 = base.I32_clz(int32(base.Ui32(v319) >> (uint(int32(8)) % 32)))
	v359 = int32(1)
	v366 = int32(base.Ui32(v319)>>(uint(int32(38)-v356)%32))&v359 - v356<<(uint(v359)%32) + int32(62)
	goto L83
L85:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v144+v438))) = v441
	*(*int32)(unsafe.Add(mBase, uint32(v144)+12)) = v440
	*(*int32)(unsafe.Add(mBase, uint32(v144+v436))) = v439
	v450 = int32(0)
	v452 = *(*int32)(unsafe.Add(mBase, _consts[1003]))
	v453 = int32(-1)
	v454 = v452 + v453
	if v454 != 0 {
		goto L97
	} else {
		goto L98
	}
L86:
	;
	v430 = *(*int32)(unsafe.Add(mBase, uint32(v400)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v430)+12)) = v144
	*(*int32)(unsafe.Add(mBase, uint32(v400)+8)) = v144
	v436 = int32(24)
	v438 = int32(8)
	v439 = int32(0)
	v440 = v400
	v441 = v430
	goto L85
L87:
	;
	v436 = v421
	v438 = v423
	v439 = v144
	v440 = v144
	v441 = v426
	goto L85
L88:
	;
	if v366 == int32(31) {
		goto L90
	} else {
		goto L91
	}
L89:
	;
	*(*int32)(unsafe.Add(mBase, _consts[999])) = v375 | v377
	*(*int32)(unsafe.Add(mBase, uint32(v371)+uint32(_consts[998]))) = v144
	v421 = int32(8)
	v423 = int32(24)
	v426 = v371 + int32(9128728)
	goto L87
L90:
	;
	v392 = int32(0)
	goto L92
L91:
	;
	v392 = int32(25) - int32(base.Ui32(v366)>>(uint(int32(1))%32))
	goto L92
L92:
	;
	v394 = *(*int32)(unsafe.Add(mBase, uint32(v371)+uint32(_consts[998])))
	v397 = v319 << (uint(v392) % 32)
	v400 = v394
	goto L93
L93:
	;
	v404 = *(*int32)(unsafe.Add(mBase, uint32(v400)+4))
	if v404&int32(-8) == v319 {
		goto L86
	} else {
		goto L95
	}
L94:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v414+int32(16)))) = v144
	v421 = int32(8)
	v423 = int32(24)
	v426 = v400
	goto L87
L95:
	;
	v414 = v400 + int32(base.Ui32(v397)>>(uint(int32(29))%32))&int32(4)
	v415 = *(*int32)(unsafe.Add(mBase, uint32(v414)+16))
	if v415 != 0 {
		v397 = v397 << (uint(int32(1)) % 32)
		v400 = v415
		goto L93
	} else {
		goto L96
	}
L96:
	;
	goto L94
L97:
	;
	v456 = v454
	goto L99
L98:
	;
	v456 = v453
	goto L99
L99:
	;
	*(*int32)(unsafe.Add(mBase, _consts[1003])) = v456
	goto L1
}
func F_emscripten_builtin_memalign(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v12 int32
	_ = v12
	var v16 int32
	_ = v16
	var v22 int32
	_ = v22
	var v30 int32
	_ = v30
	var v40 int32
	_ = v40
	var v45 int32
	_ = v45
	var v51 int32
	_ = v51
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v70 int32
	_ = v70
	var v74 int32
	_ = v74
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v101 int32
	_ = v101
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v114 int32
	_ = v114
	var v120 int32
	_ = v120
	var v126 int32
	_ = v126
	var v130 int32
	_ = v130
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v159 int32
	_ = v159
	if base.Ui32(int32(8)) < base.Ui32(l0) {
		v12 = int32(16)
		if base.Ui32(v12) < base.Ui32(l0) {
			v16 = l0
		} else {
			v16 = v12
		}
		if v16&(v16+int32(-1)) != 0 {
			v22 = v12
			for {
				if base.Ui32(v22) < base.Ui32(v16) {
					v22 = v22 << (uint(int32(1)) % 32)
					continue
				} else {
					break
				}
				break
			}
			v30 = v22
		} else {
			v30 = v16
		}
		if base.Ui32(l1) < base.Ui32(int32(-64)-v30) {
			v45 = int32(11)
			if base.Ui32(l1) < base.Ui32(v45) {
				v51 = int32(16)
			} else {
				v51 = (l1 + v45) & int32(-8)
			}
			v55 = F_emscripten_builtin_malloc(m, v51+v30+int32(12))
			mBase = m.M
			if v55 != 0 {
				v58 = v55 + int32(-8)
				if (v30+int32(-1))&v55 != 0 {
					v63 = v55 + int32(-4)
					v64 = *(*int32)(unsafe.Add(mBase, uint32(v63)))
					v65 = int32(-8)
					v70 = int32(0)
					v74 = (v55+v30+int32(-1))&(v70-v30) + v65
					if base.Ui32(int32(15)) < base.Ui32(v74-v58) {
						v79 = v70
					} else {
						v79 = v30
					}
					v80 = v74 + v79
					v81 = v80 - v58
					v82 = v64&v65 - v81
					if v64&int32(3) != 0 {
						v89 = *(*int32)(unsafe.Add(mBase, uint32(v80)+4))
						v90 = int32(1)
						v93 = int32(2)
						*(*int32)(unsafe.Add(mBase, uint32(v80)+4)) = v82 | v89&v90 | v93
						v96 = v80 + v82
						v97 = *(*int32)(unsafe.Add(mBase, uint32(v96)+4))
						*(*int32)(unsafe.Add(mBase, uint32(v96)+4)) = v97 | v90
						v101 = *(*int32)(unsafe.Add(mBase, uint32(v63)))
						*(*int32)(unsafe.Add(mBase, uint32(v63))) = v81 | v101&v90 | v93
						v108 = v58 + v81
						v109 = *(*int32)(unsafe.Add(mBase, uint32(v108)+4))
						*(*int32)(unsafe.Add(mBase, uint32(v108)+4)) = v109 | v90
						F_dispose_chunk(m, v58, v81)
						mBase = m.M
						v114 = v80
					} else {
						v85 = *(*int32)(unsafe.Add(mBase, uint32(v58)))
						*(*int32)(unsafe.Add(mBase, uint32(v80)+4)) = v82
						*(*int32)(unsafe.Add(mBase, uint32(v80))) = v85 + v81
						v114 = v80
					}
				} else {
					v114 = v58
				}
				v120 = *(*int32)(unsafe.Add(mBase, uint32(v114)+4))
				if v120&int32(3) == int32(0) {
				} else {
					v126 = v120 & int32(-8)
					if base.Ui32(v126) <= base.Ui32(v51+int32(16)) {
					} else {
						v130 = int32(1)
						*(*int32)(unsafe.Add(mBase, uint32(v114)+4)) = v51 | v120&v130 | int32(2)
						v136 = v114 + v51
						v137 = v126 - v51
						*(*int32)(unsafe.Add(mBase, uint32(v136)+4)) = v137 | int32(3)
						v141 = v114 + v126
						v142 = *(*int32)(unsafe.Add(mBase, uint32(v141)+4))
						*(*int32)(unsafe.Add(mBase, uint32(v141)+4)) = v142 | v130
						F_dispose_chunk(m, v136, v137)
						mBase = m.M
					}
				}
				v159 = v114 + int32(8)
			} else {
				v159 = int32(0)
			}
		} else {
			v40 = F___errno_location(m)
			mBase = m.M
			*(*int32)(unsafe.Add(mBase, uint32(v40))) = int32(48)
			v159 = int32(0)
		}
		return v159
	} else {
		v5 = F_emscripten_builtin_malloc(m, l1)
		mBase = m.M
		return v5
	}
}
func F_emscripten_futex_wake(m *base.Module, l0 int32, l1 int32) int32 {
	return int32(0)
}
func F_emscripten_stack_get_end(m *base.Module) int32 {
	var v1 int32
	_ = v1
	v1 = m.G1
	return v1
}
func F_emscripten_thread_sleep(m *base.Module, l0 float64) {
	var v4 float64
	_ = v4
	var v7 float64
	_ = v7
	var v9 int32
	_ = v9
	var v10 float64
	_ = v10
	v4 = m.Env.Emscripten_get_now(m)
	v7 = v4
	goto L1
L1:
	;
	F__emscripten_yield(m, v7)
	v9 = m.ExcPending
	if v9 != 0 {
		goto L3
	} else {
		goto L4
	}
L2:
	;
	return
L3:
	;
	return
L4:
	;
	v10 = m.Env.Emscripten_get_now(m)
	if base.F64_lt(base.F64_sub(v10, v4), l0) != 0 {
		v7 = v10
		goto L1
	} else {
		goto L5
	}
L5:
	;
	goto L2
}
