package p0

import (
	base "github.com/shibukawa/vkmem/internal/aot/vkaot/base"
	"math"
	"unsafe"
)

func F___fdopen(m *base.Module, l0 int32, l1 int32) int32 {
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
	var v13 int32
	_ = v13
	var v17 int32
	_ = v17
	var v22 int32
	_ = v22
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v81 int32
	_ = v81
	var v93 int32
	_ = v93
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v107 int32
	_ = v107
	v5 = m.G0
	v7 = v5 - int32(32)
	m.G0 = v7
	v10 = int32(*(*int8)(unsafe.Add(mBase, uint32(l1))))
	v11 = F___strchrnul(m, int32(_a_F___fdopen_0), v10)
	mBase = m.M
	v13 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11))))
	if v13 == v10&int32(255) {
		v17 = v11
	} else {
		v17 = int32(0)
	}
	if v17 != 0 {
		v22 = F_emscripten_builtin_malloc(m, int32(1176))
		mBase = m.M
		if v22 != 0 {
			v28 = F__emscripten_memset_bulkmem(m, v22, base.I32_extend8_s(int32(0)), int32(144))
			mBase = m.M
			v29 = int32(43)
			v30 = F___strchrnul(m, l1, v29)
			mBase = m.M
			v32 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v30))))
			if v32 == v29 {
				v36 = v30
			} else {
				v36 = int32(0)
			}
			if v36 != 0 {
			} else {
				v39 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
				if v39 == int32(114) {
					v42 = int32(8)
				} else {
					v42 = int32(4)
				}
				*(*int32)(unsafe.Add(mBase, uint32(v22))) = v42
			}
			v44 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
			if v44 == int32(97) {
				v50 = m.Env.X__syscall_fcntl64(m, l0, int32(3), int32(0))
				mBase = m.M
				if v50&int32(1024) != 0 {
				} else {
					*(*int64)(unsafe.Add(mBase, uint32(v7)+16)) = base.I64_extend_i32_s(v50 | int32(1024))
					v60 = m.Env.X__syscall_fcntl64(m, l0, int32(4), v7+int32(16))
					mBase = m.M
				}
				v61 = *(*int32)(unsafe.Add(mBase, uint32(v22)))
				v63 = v61 | int32(128)
				*(*int32)(unsafe.Add(mBase, uint32(v22))) = v63
				v65 = v63
			} else {
				v47 = *(*int32)(unsafe.Add(mBase, uint32(v22)))
				v65 = v47
			}
			*(*int32)(unsafe.Add(mBase, uint32(v22)+80)) = int32(-1)
			*(*int32)(unsafe.Add(mBase, uint32(v22)+48)) = int32(1024)
			*(*int32)(unsafe.Add(mBase, uint32(v22)+60)) = l0
			*(*int32)(unsafe.Add(mBase, uint32(v22)+44)) = v22 + int32(152)
			if v65&int32(8) != 0 {
			} else {
				*(*int64)(unsafe.Add(mBase, uint32(v7))) = base.I64_extend_i32_u(v7 + int32(24))
				v81 = m.Env.X__syscall_ioctl(m, l0, int32(21523), v7)
				mBase = m.M
				if v81 != 0 {
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v22)+80)) = int32(10)
				}
			}
			*(*int32)(unsafe.Add(mBase, uint32(v22)+40)) = int32(1375)
			*(*int32)(unsafe.Add(mBase, uint32(v22)+36)) = int32(1376)
			*(*int32)(unsafe.Add(mBase, uint32(v22)+32)) = int32(1377)
			*(*int32)(unsafe.Add(mBase, uint32(v22)+12)) = int32(1378)
			v93 = int32(*(*uint8)(unsafe.Add(mBase, _c_F___fdopen[0])))
			if v93 != 0 {
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v22)+76)) = int32(-1)
			}
			v98 = F___ofl_lock(m)
			mBase = m.M
			v99 = *(*int32)(unsafe.Add(mBase, uint32(v98)))
			*(*int32)(unsafe.Add(mBase, uint32(v22)+56)) = v99
			if v99 == int32(0) {
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v99)+52)) = v22
			}
			*(*int32)(unsafe.Add(mBase, uint32(v98))) = v22
			F___ofl_unlock(m)
			mBase = m.M
			v107 = v22
		} else {
			v107 = int32(0)
		}
	} else {
		*(*int32)(unsafe.Add(mBase, _c_F___fdopen[1])) = int32(28)
		v107 = int32(0)
	}
	m.G0 = v7 + int32(32)
	return v107
}
func F___floatscan(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v6 int64
	_ = v6
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v90 int32
	_ = v90
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
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
	var v122 int32
	_ = v122
	var v135 int64
	_ = v135
	var v138 int32
	_ = v138
	var v157 int32
	_ = v157
	var v160 int32
	_ = v160
	var v165 int32
	_ = v165
	var v190 int32
	_ = v190
	var v192 int32
	_ = v192
	var v194 int32
	_ = v194
	var v196 int32
	_ = v196
	var v198 int32
	_ = v198
	var v200 int32
	_ = v200
	var v218 int64
	_ = v218
	var v223 int32
	_ = v223
	var v231 int64
	_ = v231
	var v234 int64
	_ = v234
	var v235 int32
	_ = v235
	var v236 int64
	_ = v236
	var v237 int64
	_ = v237
	var v255 int64
	_ = v255
	var v256 int64
	_ = v256
	var v257 int32
	_ = v257
	var v271 int32
	_ = v271
	var v276 int32
	_ = v276
	var v277 int32
	_ = v277
	var v282 int32
	_ = v282
	var v283 int32
	_ = v283
	var v284 int32
	_ = v284
	var v285 int32
	_ = v285
	var v289 int32
	_ = v289
	var v292 int32
	_ = v292
	var v296 int32
	_ = v296
	var v303 int32
	_ = v303
	var v318 int32
	_ = v318
	var v319 int32
	_ = v319
	var v324 int32
	_ = v324
	var v325 int32
	_ = v325
	var v326 int32
	_ = v326
	var v327 int32
	_ = v327
	var v331 int64
	_ = v331
	var v332 int64
	_ = v332
	var v333 int64
	_ = v333
	var v336 int32
	_ = v336
	var v349 int32
	_ = v349
	var v352 int32
	_ = v352
	var v353 int32
	_ = v353
	var v358 int32
	_ = v358
	var v359 int32
	_ = v359
	var v360 int32
	_ = v360
	var v361 int32
	_ = v361
	var v379 int64
	_ = v379
	var v382 int64
	_ = v382
	var v385 int32
	_ = v385
	var v404 int32
	_ = v404
	var v409 int32
	_ = v409
	var v414 int32
	_ = v414
	var v417 int64
	_ = v417
	var v418 int64
	_ = v418
	var v421 int32
	_ = v421
	var v433 int64
	_ = v433
	var v443 int32
	_ = v443
	var v444 int32
	_ = v444
	var v448 int32
	_ = v448
	var v456 int32
	_ = v456
	var v460 int32
	_ = v460
	var v461 int32
	_ = v461
	var v466 int32
	_ = v466
	var v467 int32
	_ = v467
	var v468 int32
	_ = v468
	var v469 int32
	_ = v469
	var v477 int32
	_ = v477
	var v480 int64
	_ = v480
	var v481 int64
	_ = v481
	var v482 int64
	_ = v482
	var v485 int32
	_ = v485
	var v493 int32
	_ = v493
	var v496 int64
	_ = v496
	var v497 int64
	_ = v497
	var v516 int64
	_ = v516
	var v529 int64
	_ = v529
	var v535 int64
	_ = v535
	v6 = int64(0)
	v13 = m.G0
	v15 = v13 - int32(48)
	m.G0 = v15
	if base.Ui32(int32(2)) < base.Ui32(l2) {
		v516 = v6
		goto L2
	} else {
		goto L3
	}
L1:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l0))) = v529
	*(*int64)(unsafe.Add(mBase, uint32(l0)+8)) = v535
	m.G0 = v15 + int32(48)
	return
L2:
	;
	v529 = v516
	v535 = int64(0)
	goto L1
L3:
	;
	v21 = l2 << (uint(int32(2)) % 32)
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v21)+uint32(_c_F___floatscan[0])))
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v21)+uint32(_c_F___floatscan[1])))
	goto L4
L4:
	;
	v40 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v41 = *(*int32)(unsafe.Add(mBase, uint32(l1)+104))
	if v40 == v41 {
		goto L7
	} else {
		goto L8
	}
L5:
	;
	switch v49 + int32(-43) {
	case 0, 2:
		goto L14
	default:
		v74 = v49
		v75 = int32(1)
		goto L13
	}
L6:
	;
	goto L11
L7:
	;
	v47 = F___shgetc(m, l1)
	mBase = m.M
	v48 = m.ExcPending
	if v48 != 0 {
		goto L9
	} else {
		goto L10
	}
L8:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v40 + int32(1)
	v46 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v40))))
	v49 = v46
	goto L6
L9:
	;
	return
L10:
	;
	v49 = v47
	goto L6
L11:
	;
	if base.B2i32(v49 == int32(32))|base.B2i32(base.Ui32(v49+int32(-9)) < base.Ui32(int32(5))) != 0 {
		goto L4
	} else {
		goto L12
	}
L12:
	;
	goto L5
L13:
	;
	v76 = int32(0)
	if v74&int32(-33) != int32(73) {
		v115 = v74
		v122 = v76
		goto L23
	} else {
		goto L24
	}
L14:
	;
	if v49 == int32(45) {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	v64 = int32(-1)
	goto L17
L16:
	;
	v64 = int32(1)
	goto L17
L17:
	;
	v65 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v66 = *(*int32)(unsafe.Add(mBase, uint32(l1)+104))
	if v65 == v66 {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	v72 = F___shgetc(m, l1)
	mBase = m.M
	v73 = m.ExcPending
	if v73 != 0 {
		goto L9
	} else {
		goto L20
	}
L19:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v65 + int32(1)
	v71 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v65))))
	v74 = v71
	v75 = v64
	goto L13
L20:
	;
	v74 = v72
	v75 = v64
	goto L13
L21:
	;
	if v122 != 0 {
		v296 = v115
		v303 = v122
		goto L61
	} else {
		goto L62
	}
L22:
	;
	v190 = m.G0
	v192 = v190 - int32(16)
	m.G0 = v192
	v194 = base.I32_reinterpret_f32(base.F32_mul(base.F32_convert_i32_s(v75), math.Float32frombits(uint32(0x7f800000))))
	v196 = v194 & int32(8388607)
	v198 = int32(base.Ui32(v194) >> (uint(int32(23)) % 32))
	v200 = v198 & int32(255)
	if v200 == int32(0) {
		goto L50
	} else {
		goto L51
	}
L23:
	;
	if v122 == int32(3) {
		goto L33
	} else {
		goto L34
	}
L24:
	;
	v90 = v76
	goto L25
L25:
	;
	if v90 == int32(7) {
		goto L22
	} else {
		goto L27
	}
L26:
	;
	v115 = v104
	v122 = v108
	goto L23
L27:
	;
	v95 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v96 = *(*int32)(unsafe.Add(mBase, uint32(l1)+104))
	if v95 == v96 {
		goto L29
	} else {
		goto L30
	}
L28:
	;
	v108 = v90 + int32(1)
	v111 = int32(*(*int8)(unsafe.Add(mBase, uint32(v90)+uint32(_c_F___floatscan[2]))))
	if v104|int32(32) == v111 {
		v90 = v108
		goto L25
	} else {
		goto L32
	}
L29:
	;
	v102 = F___shgetc(m, l1)
	mBase = m.M
	v103 = m.ExcPending
	if v103 != 0 {
		goto L9
	} else {
		goto L31
	}
L30:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v95 + int32(1)
	v101 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v95))))
	v104 = v101
	goto L28
L31:
	;
	v104 = v102
	goto L28
L32:
	;
	goto L26
L33:
	;
	v135 = *(*int64)(unsafe.Add(mBase, uint32(l1)+112))
	if v135 < int64(0) {
		goto L39
	} else {
		goto L40
	}
L34:
	;
	if v122 == int32(8) {
		goto L22
	} else {
		goto L35
	}
L35:
	;
	if l3 == int32(0) {
		goto L21
	} else {
		goto L36
	}
L36:
	;
	if base.Ui32(v122) < base.Ui32(int32(4)) {
		goto L21
	} else {
		goto L37
	}
L37:
	;
	if v122 == int32(8) {
		goto L22
	} else {
		goto L38
	}
L38:
	;
	goto L33
L39:
	;
	if l3 == int32(0) {
		goto L22
	} else {
		goto L41
	}
L40:
	;
	v138 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v138 + int32(-1)
	goto L39
L41:
	;
	if base.Ui32(v122) < base.Ui32(int32(4)) {
		goto L22
	} else {
		goto L42
	}
L42:
	;
	v157 = v122
	goto L43
L43:
	;
	if v135 < int64(0) {
		goto L45
	} else {
		goto L46
	}
L44:
	;
	goto L22
L45:
	;
	v165 = v157 + int32(-1)
	if base.Ui32(int32(3)) < base.Ui32(v165) {
		v157 = v165
		goto L43
	} else {
		goto L47
	}
L46:
	;
	v160 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v160 + int32(-1)
	goto L45
L47:
	;
	goto L44
L48:
	;
	v255 = *(*int64)(unsafe.Add(mBase, uint32(v15+int32(8))))
	v256 = *(*int64)(unsafe.Add(mBase, uint32(v15)))
	v529 = v256
	v535 = v255
	goto L1
L49:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v15))) = v237
	*(*int64)(unsafe.Add(mBase, uint32(v15)+8)) = base.I64_extend_i32_u(v235)<<(uint(int64(48))%64) | base.I64_extend_i32_u(int32(base.Ui32(v194)>>(uint(int32(31))%32)))<<(uint(int64(63))%64) | v236
	m.G0 = v192 + int32(16)
	goto L48
L50:
	;
	if v196 != 0 {
		goto L54
	} else {
		goto L55
	}
L51:
	;
	if v200 == int32(255) {
		goto L52
	} else {
		goto L53
	}
L52:
	;
	v235 = int32(32767)
	v236 = base.I64_extend_i32_u(v196) << (uint(int64(25)) % 64)
	v237 = int64(0)
	goto L49
L53:
	;
	v235 = v198&int32(255) + int32(16256)
	v236 = base.I64_extend_i32_u(v196) << (uint(int64(25)) % 64)
	v237 = int64(0)
	goto L49
L54:
	;
	v223 = base.I32_clz(v196)
	F___ashlti3(m, v192, base.I64_extend_i32_u(v196), int64(0), v223+int32(81))
	mBase = m.M
	v231 = *(*int64)(unsafe.Add(mBase, uint32(v192+int32(8))))
	v234 = *(*int64)(unsafe.Add(mBase, uint32(v192)))
	v235 = int32(16265) - v223
	v236 = v231 ^ int64(281474976710656)
	v237 = v234
	goto L49
L55:
	;
	v218 = int64(0)
	v235 = int32(0)
	v236 = v218
	v237 = v218
	goto L49
L56:
	;
	v529 = int64(0)
	v535 = v379
	goto L1
L57:
	;
	if v296 != int32(48) {
		goto L112
	} else {
		goto L113
	}
L58:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l1)+112)) = v433
	v443 = *(*int32)(unsafe.Add(mBase, uint32(l1)+44))
	v444 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	*(*int64)(unsafe.Add(mBase, uint32(l1)+120)) = base.I64_extend_i32_s(v443 - v444)
	v448 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	if v433 == int64(0) {
		v456 = v448
		goto L109
	} else {
		goto L110
	}
L59:
	;
	v417 = int64(0)
	v418 = *(*int64)(unsafe.Add(mBase, uint32(l1)+112))
	if v418 < v417 {
		goto L105
	} else {
		goto L106
	}
L60:
	;
	v318 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v319 = *(*int32)(unsafe.Add(mBase, uint32(l1)+104))
	if v318 == v319 {
		goto L73
	} else {
		goto L74
	}
L61:
	;
	switch v303 {
	case 0:
		goto L57
	default:
		goto L59
	case 3:
		goto L60
	}
L62:
	;
	v257 = int32(0)
	if v115&int32(-33) != int32(78) {
		v296 = v115
		v303 = v257
		goto L61
	} else {
		goto L63
	}
L63:
	;
	v271 = v257
	goto L64
L64:
	;
	if v271 == int32(2) {
		goto L60
	} else {
		goto L66
	}
L65:
	;
	v296 = v285
	v303 = v289
	goto L61
L66:
	;
	v276 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v277 = *(*int32)(unsafe.Add(mBase, uint32(l1)+104))
	if v276 == v277 {
		goto L68
	} else {
		goto L69
	}
L67:
	;
	v289 = v271 + int32(1)
	v292 = int32(*(*int8)(unsafe.Add(mBase, uint32(v271)+uint32(_c_F___floatscan[3]))))
	if v285|int32(32) == v292 {
		v271 = v289
		goto L64
	} else {
		goto L71
	}
L68:
	;
	v283 = F___shgetc(m, l1)
	mBase = m.M
	v284 = m.ExcPending
	if v284 != 0 {
		goto L9
	} else {
		goto L70
	}
L69:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v276 + int32(1)
	v282 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v276))))
	v285 = v282
	goto L67
L70:
	;
	v285 = v283
	goto L67
L71:
	;
	goto L65
L72:
	;
	if v327 != int32(40) {
		goto L77
	} else {
		goto L78
	}
L73:
	;
	v325 = F___shgetc(m, l1)
	mBase = m.M
	v326 = m.ExcPending
	if v326 != 0 {
		goto L9
	} else {
		goto L75
	}
L74:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v318 + int32(1)
	v324 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v318))))
	v327 = v324
	goto L72
L75:
	;
	v327 = v325
	goto L72
L76:
	;
	v349 = int32(1)
	goto L80
L77:
	;
	v331 = int64(0)
	v332 = int64(9223231299366420480)
	v333 = *(*int64)(unsafe.Add(mBase, uint32(l1)+112))
	if v333 < v331 {
		v529 = v331
		v535 = v332
		goto L1
	} else {
		goto L79
	}
L78:
	;
	goto L76
L79:
	;
	v336 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v336 + int32(-1)
	v529 = v331
	v535 = v332
	goto L1
L80:
	;
	v352 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v353 = *(*int32)(unsafe.Add(mBase, uint32(l1)+104))
	if v352 == v353 {
		goto L83
	} else {
		goto L84
	}
L81:
	;
	v379 = int64(9223231299366420480)
	if v361 == int32(41) {
		v529 = v6
		v535 = v379
		goto L1
	} else {
		goto L92
	}
L82:
	;
	if base.Ui32(v361+int32(-48)) < base.Ui32(int32(10)) {
		goto L87
	} else {
		goto L88
	}
L83:
	;
	v359 = F___shgetc(m, l1)
	mBase = m.M
	v360 = m.ExcPending
	if v360 != 0 {
		goto L9
	} else {
		goto L85
	}
L84:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v352 + int32(1)
	v358 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v352))))
	v361 = v358
	goto L82
L85:
	;
	v361 = v359
	goto L82
L86:
	;
	goto L81
L87:
	;
	v349 = v349 + int32(1)
	goto L80
L88:
	;
	if base.Ui32(v361+int32(-65)) < base.Ui32(int32(26)) {
		goto L87
	} else {
		goto L89
	}
L89:
	;
	if v361 == int32(95) {
		goto L87
	} else {
		goto L90
	}
L90:
	;
	if base.Ui32(int32(26)) <= base.Ui32(v361+int32(-97)) {
		goto L86
	} else {
		goto L91
	}
L91:
	;
	goto L87
L92:
	;
	v382 = *(*int64)(unsafe.Add(mBase, uint32(l1)+112))
	if v382 < int64(0) {
		goto L93
	} else {
		goto L94
	}
L93:
	;
	if l3 == int32(0) {
		goto L96
	} else {
		goto L97
	}
L94:
	;
	v385 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v385 + int32(-1)
	goto L93
L95:
	;
	v404 = v349
	goto L100
L96:
	;
	goto L99
L97:
	;
	if v349 != 0 {
		goto L95
	} else {
		goto L98
	}
L98:
	;
	goto L56
L99:
	;
	*(*int32)(unsafe.Add(mBase, _c_F___floatscan[4])) = int32(28)
	v433 = int64(0)
	goto L58
L100:
	;
	if v382 < int64(0) {
		goto L102
	} else {
		goto L103
	}
L102:
	;
	v414 = v404 + int32(-1)
	if v414 == int32(0) {
		goto L56
	} else {
		goto L104
	}
L103:
	;
	v409 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v409 + int32(-1)
	goto L102
L104:
	;
	v404 = v414
	goto L100
L105:
	;
	goto L107
L106:
	;
	v421 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v421 + int32(-1)
	goto L105
L107:
	;
	*(*int32)(unsafe.Add(mBase, _c_F___floatscan[4])) = int32(28)
	v433 = v417
	goto L58
L108:
	;
	v516 = v433
	goto L2
L109:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+104)) = v456
	goto L108
L110:
	;
	if base.I64_extend_i32_s(v448-v444) <= v433 {
		v456 = v448
		goto L109
	} else {
		goto L111
	}
L111:
	;
	v456 = v444 + base.I32_wrap_i64(v433)
	goto L109
L112:
	;
	F_decfloat(m, v15+int32(32), l1, v296, v27, v24, v75, l3)
	mBase = m.M
	v493 = m.ExcPending
	if v493 != 0 {
		goto L9
	} else {
		goto L122
	}
L113:
	;
	v460 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v461 = *(*int32)(unsafe.Add(mBase, uint32(l1)+104))
	if v460 == v461 {
		goto L115
	} else {
		goto L116
	}
L114:
	;
	if v469&int32(-33) != int32(88) {
		goto L118
	} else {
		goto L119
	}
L115:
	;
	v467 = F___shgetc(m, l1)
	mBase = m.M
	v468 = m.ExcPending
	if v468 != 0 {
		goto L9
	} else {
		goto L117
	}
L116:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v460 + int32(1)
	v466 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v460))))
	v469 = v466
	goto L114
L117:
	;
	v469 = v467
	goto L114
L118:
	;
	v482 = *(*int64)(unsafe.Add(mBase, uint32(l1)+112))
	if v482 < int64(0) {
		goto L112
	} else {
		goto L121
	}
L119:
	;
	F_hexfloat(m, v15+int32(16), l1, v27, v24, v75, l3)
	mBase = m.M
	v477 = m.ExcPending
	if v477 != 0 {
		goto L9
	} else {
		goto L120
	}
L120:
	;
	v480 = *(*int64)(unsafe.Add(mBase, uint32(v15+int32(24))))
	v481 = *(*int64)(unsafe.Add(mBase, uint32(v15)+16))
	v529 = v481
	v535 = v480
	goto L1
L121:
	;
	v485 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v485 + int32(-1)
	goto L112
L122:
	;
	v496 = *(*int64)(unsafe.Add(mBase, uint32(v15+int32(40))))
	v497 = *(*int64)(unsafe.Add(mBase, uint32(v15)+32))
	v529 = v497
	v535 = v496
	goto L1
}
func F___floatsitf(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int64
	_ = v11
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v17 int64
	_ = v17
	var v18 int64
	_ = v18
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v38 int64
	_ = v38
	var v42 int64
	_ = v42
	var v43 int64
	_ = v43
	var v49 int64
	_ = v49
	var v64 int64
	_ = v64
	var v65 int64
	_ = v65
	var v66 int64
	_ = v66
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	if l1 != 0 {
		v14 = l1 >> (uint(int32(31)) % 32)
		v16 = l1 ^ v14 - v14
		v17 = base.I64_extend_i32_u(v16)
		v18 = int64(0)
		v19 = base.I32_clz(v16)
		v21 = v19 + int32(81)
		if v21&int32(64) == int32(0) {
			if v21 == int32(0) {
				v42 = v17
				v43 = v18
			} else {
				v38 = base.I64_extend_i32_u(v21)
				v42 = v17 << (uint(v38) % 64)
				v43 = int64(base.Ui64(v17)>>(uint(base.I64_extend_i32_u(int32(64)-v21))%64)) | v18<<(uint(v38)%64)
			}
		} else {
			v42 = int64(0)
			v43 = v17 << (uint(base.I64_extend_i32_u(v19+int32(17))) % 64)
		}
		*(*int64)(unsafe.Add(mBase, uint32(v9))) = v42
		*(*int64)(unsafe.Add(mBase, uint32(v9)+8)) = v43
		v49 = *(*int64)(unsafe.Add(mBase, uint32(v9+int32(8))))
		v64 = *(*int64)(unsafe.Add(mBase, uint32(v9)))
		v65 = v64
		v66 = v49 ^ int64(281474976710656) + base.I64_extend_i32_u(int32(16414)-v19)<<(uint(int64(48))%64) | base.I64_extend_i32_u(l1&int32(-2147483648))<<(uint(int64(32))%64)
	} else {
		v11 = int64(0)
		v65 = v11
		v66 = v11
	}
	*(*int64)(unsafe.Add(mBase, uint32(l0))) = v65
	*(*int64)(unsafe.Add(mBase, uint32(l0)+8)) = v66
	m.G0 = v9 + int32(16)
	return
}
func F___floatunsitf(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int64
	_ = v10
	var v12 int64
	_ = v12
	var v13 int64
	_ = v13
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v35 int64
	_ = v35
	var v39 int64
	_ = v39
	var v40 int64
	_ = v40
	var v46 int64
	_ = v46
	var v55 int64
	_ = v55
	var v57 int64
	_ = v57
	var v58 int64
	_ = v58
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	if l1 != 0 {
		v12 = base.I64_extend_i32_u(l1)
		v13 = int64(0)
		v15 = base.I32_clz(l1)
		v18 = int32(112) - (v15 ^ int32(31))
		if v18&int32(64) == int32(0) {
			if v18 == int32(0) {
				v39 = v12
				v40 = v13
			} else {
				v35 = base.I64_extend_i32_u(v18)
				v39 = v12 << (uint(v35) % 64)
				v40 = int64(base.Ui64(v12)>>(uint(base.I64_extend_i32_u(int32(64)-v18))%64)) | v13<<(uint(v35)%64)
			}
		} else {
			v39 = int64(0)
			v40 = v12 << (uint(base.I64_extend_i32_u(v18+int32(-64))) % 64)
		}
		*(*int64)(unsafe.Add(mBase, uint32(v8))) = v39
		*(*int64)(unsafe.Add(mBase, uint32(v8)+8)) = v40
		v46 = *(*int64)(unsafe.Add(mBase, uint32(v8+int32(8))))
		v55 = *(*int64)(unsafe.Add(mBase, uint32(v8)))
		v57 = v55
		v58 = v46 ^ int64(281474976710656) + base.I64_extend_i32_u(int32(16414)-v15)<<(uint(int64(48))%64)
	} else {
		v10 = int64(0)
		v57 = v10
		v58 = v10
	}
	*(*int64)(unsafe.Add(mBase, uint32(l0))) = v57
	*(*int64)(unsafe.Add(mBase, uint32(l0)+8)) = v58
	m.G0 = v8 + int32(16)
	return
}
func F___fmodeflags(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v48 int32
	_ = v48
	var v53 int32
	_ = v53
	v4 = int32(43)
	v5 = F___strchrnul(m, l0, v4)
	mBase = m.M
	v7 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5))))
	if v7 == v4 {
		v11 = v5
	} else {
		v11 = int32(0)
	}
	if v11 != 0 {
		v15 = int32(2)
	} else {
		v12 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
		v15 = base.B2i32(v12 != int32(114))
	}
	v18 = int32(120)
	v19 = F___strchrnul(m, l0, v18)
	mBase = m.M
	v21 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19))))
	if v21 == v18 {
		v25 = v19
	} else {
		v25 = int32(0)
	}
	if v25 != 0 {
		v26 = v15 | int32(128)
	} else {
		v26 = v15
	}
	v29 = int32(101)
	v30 = F___strchrnul(m, l0, v29)
	mBase = m.M
	v32 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v30))))
	if v32 == v29 {
		v36 = v30
	} else {
		v36 = int32(0)
	}
	if v36 != 0 {
		v37 = v26 | int32(524288)
	} else {
		v37 = v26
	}
	v40 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v40 == int32(114) {
		v43 = v37
	} else {
		v43 = v37 | int32(64)
	}
	if v40 == int32(119) {
		v48 = v43 | int32(512)
	} else {
		v48 = v43
	}
	if v40 == int32(97) {
		v53 = v48 | int32(1024)
	} else {
		v53 = v48
	}
	return v53
}
func F___fseeko_unlocked(m *base.Module, l0 int32, l1 int64, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v19 int64
	_ = v19
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v38 int32
	_ = v38
	var v39 int64
	_ = v39
	var v40 int32
	_ = v40
	var v45 int32
	_ = v45
	if base.Ui32(l2) < base.Ui32(int32(3)) {
		if l2 != int32(1) {
			v19 = l1
		} else {
			v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			if v12 == int32(0) {
				v19 = l1
			} else {
				v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
				v19 = l1 - base.I64_extend_i32_s(v12-v15)
			}
		}
		v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
		v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		if v21 == v22 {
			*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = int32(0)
			*(*int64)(unsafe.Add(mBase, uint32(l0)+16)) = int64(0)
			v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
			v39 = m.T0[v38].(func(*base.Module, int32, int64, int32) int64)(m, l0, v19, l2)
			mBase = m.M
			v40 = m.ExcPending
			if v40 != 0 {
				return int32(0)
			} else {
				if v39 < int64(0) {
					return int32(-1)
				} else {
					*(*int64)(unsafe.Add(mBase, uint32(l0)+4)) = int64(0)
					v45 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
					*(*int32)(unsafe.Add(mBase, uint32(l0))) = v45 & int32(-17)
					return int32(0)
				}
			}
		} else {
			v24 = int32(0)
			v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
			v27 = m.T0[v26].(func(*base.Module, int32, int32, int32) int32)(m, l0, v24, v24)
			mBase = m.M
			v30 = m.ExcPending
			if v30 != 0 {
				return int32(0)
			} else {
				v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
				if v31 == int32(0) {
					return int32(-1)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = int32(0)
					*(*int64)(unsafe.Add(mBase, uint32(l0)+16)) = int64(0)
					v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
					v39 = m.T0[v38].(func(*base.Module, int32, int64, int32) int64)(m, l0, v19, l2)
					mBase = m.M
					v40 = m.ExcPending
					if v40 != 0 {
						return int32(0)
					} else {
						if v39 < int64(0) {
							return int32(-1)
						} else {
							*(*int64)(unsafe.Add(mBase, uint32(l0)+4)) = int64(0)
							v45 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
							*(*int32)(unsafe.Add(mBase, uint32(l0))) = v45 & int32(-17)
							return int32(0)
						}
					}
				}
			}
		}
	} else {
		*(*int32)(unsafe.Add(mBase, _c_F___fseeko_unlocked[0])) = int32(28)
		return int32(-1)
	}
}
func F___ftello_unlocked(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v20 int64
	_ = v20
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v30 int32
	_ = v30
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v43 int64
	_ = v43
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v7 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v7&int32(128) == int32(0) {
		v18 = int32(1)
	} else {
		v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
		v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		if v14 == v15 {
			v17 = int32(1)
		} else {
			v17 = int32(2)
		}
		v18 = v17
	}
	v20 = m.T0[v5].(func(*base.Module, int32, int64, int32) int64)(m, l0, int64(0), v18)
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		return int64(0)
	} else {
		if v20 < int64(0) {
			v43 = v20
		} else {
			v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			if v26 == int32(0) {
				v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
				if v30 == int32(0) {
					v43 = v20
				} else {
					v34 = int32(20)
					v35 = v30
					v37 = *(*int32)(unsafe.Add(mBase, uint32(l0+v34)))
					v43 = v20 + base.I64_extend_i32_s(v37-v35)
				}
			} else {
				v34 = int32(4)
				v35 = v26
				v37 = *(*int32)(unsafe.Add(mBase, uint32(l0+v34)))
				v43 = v20 + base.I64_extend_i32_s(v37-v35)
			}
		}
		return v43
	}
}
func F_f_call(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	F_luaD_call(m, l0, v3, v4)
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return
	} else {
		return
	}
}
func F_f_parser(m *base.Module, l0 int32, l1 int32) {
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
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
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
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v33 int32
	_ = v33
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v52 int32
	_ = v52
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v7 = F_luaZ_lookahead(m, v6)
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(v9)+68))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(v9)+64))
	if base.Ui32(v10) < base.Ui32(v11) {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	v19 = F_luaY_parser(m, l0, v15, l1+int32(4), v18)
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		goto L1
	} else {
		goto L6
	}
L4:
	;
	F_luaC_step(m, l0)
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		goto L1
	} else {
		goto L5
	}
L5:
	;
	goto L3
L6:
	;
	v21 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19)+72)))
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+72))
	v23 = F_luaF_newLclosure(m, l0, v21, v22)
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		goto L1
	} else {
		goto L7
	}
L7:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+16)) = v19
	v26 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19)+72)))
	if v26 == int32(0) {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v52 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v52)+8)) = int32(6)
	*(*int32)(unsafe.Add(mBase, uint32(v52))) = v23
	v56 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v57 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if int32(16) < v56-v57 {
		v72 = v57
		goto L14
	} else {
		goto L15
	}
L9:
	;
	v33 = int32(0)
	goto L10
L10:
	;
	v40 = F_luaF_newupval(m, l0)
	mBase = m.M
	v41 = m.ExcPending
	if v41 != 0 {
		goto L1
	} else {
		goto L12
	}
L11:
	;
	goto L8
L12:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23+int32(20)+v33<<(uint(int32(2))%32)))) = v40
	v44 = v33 + int32(1)
	v45 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19)+72)))
	if base.Ui32(v44) < base.Ui32(v45) {
		v33 = v44
		goto L10
	} else {
		goto L13
	}
L13:
	;
	goto L11
L14:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v72 + int32(16)
	return
L15:
	;
	v61 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	v62 = int32(1)
	if v61 < v62 {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	v68 = v61 + v62
	goto L18
L17:
	;
	v68 = v61 << (uint(v62) % 32)
	goto L18
L18:
	;
	F_luaD_reallocstack(m, l0, v68)
	mBase = m.M
	v70 = m.ExcPending
	if v70 != 0 {
		goto L1
	} else {
		goto L19
	}
L19:
	;
	v71 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v72 = v71
	goto L14
}
func F_fabs(m *base.Module, l0 float64) float64 {
	return base.F64_abs(l0)
}
func F_fclose(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
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
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v36 int32
	_ = v36
	var v41 int32
	_ = v41
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v94 int32
	_ = v94
	var v99 int32
	_ = v99
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v110 int32
	_ = v110
	var __phi110 int32
	_ = __phi110
	var v111 int32
	_ = v111
	var __phi111 int32
	_ = __phi111
	var v117 int32
	_ = v117
	var v120 int32
	_ = v120
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v142 int32
	_ = v142
	var v151 int32
	_ = v151
	var v153 int32
	_ = v153
	var v156 int32
	_ = v156
	var v159 int32
	_ = v159
	var v161 int32
	_ = v161
	var v166 int32
	_ = v166
	var v173 int32
	_ = v173
	var v178 int32
	_ = v178
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v193 int32
	_ = v193
	var v201 int32
	_ = v201
	var v203 int32
	_ = v203
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v214 int32
	_ = v214
	var v216 int32
	_ = v216
	var v223 int32
	_ = v223
	var v225 int32
	_ = v225
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v243 int32
	_ = v243
	var v245 int32
	_ = v245
	var v247 int32
	_ = v247
	var v256 int32
	_ = v256
	var v258 int32
	_ = v258
	var v261 int32
	_ = v261
	var v266 int32
	_ = v266
	var v271 int32
	_ = v271
	var v272 int32
	_ = v272
	var v277 int32
	_ = v277
	var __phi277 int32
	_ = __phi277
	var v278 int32
	_ = v278
	var __phi278 int32
	_ = __phi278
	var v284 int32
	_ = v284
	var v287 int32
	_ = v287
	var v301 int32
	_ = v301
	var v310 int32
	_ = v310
	var v312 int32
	_ = v312
	var v315 int32
	_ = v315
	var v318 int32
	_ = v318
	var v320 int32
	_ = v320
	var v325 int32
	_ = v325
	var v332 int32
	_ = v332
	var v337 int32
	_ = v337
	var v359 int32
	_ = v359
	var v371 int32
	_ = v371
	var v373 int32
	_ = v373
	var v375 int32
	_ = v375
	var v379 int32
	_ = v379
	var v384 int32
	_ = v384
	var v385 int32
	_ = v385
	var v396 int32
	_ = v396
	var v399 int32
	_ = v399
	var v406 int32
	_ = v406
	var v411 int32
	_ = v411
	var v415 int32
	_ = v415
	var v417 int32
	_ = v417
	var v432 int32
	_ = v432
	var v434 int32
	_ = v434
	var v437 int32
	_ = v437
	var v440 int32
	_ = v440
	var v444 int32
	_ = v444
	var v454 int32
	_ = v454
	var v455 int32
	_ = v455
	var v461 int32
	_ = v461
	var v463 int32
	_ = v463
	var v466 int32
	_ = v466
	var v470 int32
	_ = v470
	var v476 int32
	_ = v476
	var v478 int32
	_ = v478
	var v479 int32
	_ = v479
	var v480 int32
	_ = v480
	var v481 int32
	_ = v481
	var v490 int32
	_ = v490
	var v492 int32
	_ = v492
	var v493 int32
	_ = v493
	var v494 int32
	_ = v494
	var v496 int32
	_ = v496
	var v526 int32
	_ = v526
	var v527 int32
	_ = v527
	var v530 int32
	_ = v530
	var v532 int32
	_ = v532
	var v533 int32
	_ = v533
	var v540 int32
	_ = v540
	var v541 int32
	_ = v541
	var v543 int32
	_ = v543
	var v545 int32
	_ = v545
	var v547 int32
	_ = v547
	var v549 int32
	_ = v549
	var v552 int32
	_ = v552
	var v554 int32
	_ = v554
	var v556 int32
	_ = v556
	var v563 int32
	_ = v563
	var v565 int32
	_ = v565
	var v568 int32
	_ = v568
	var v573 int32
	_ = v573
	var v578 int32
	_ = v578
	var v579 int32
	_ = v579
	var v584 int32
	_ = v584
	var __phi584 int32
	_ = __phi584
	var v585 int32
	_ = v585
	var __phi585 int32
	_ = __phi585
	var v591 int32
	_ = v591
	var v594 int32
	_ = v594
	var v597 int32
	_ = v597
	var v598 int32
	_ = v598
	var v616 int32
	_ = v616
	var v625 int32
	_ = v625
	var v627 int32
	_ = v627
	var v630 int32
	_ = v630
	var v633 int32
	_ = v633
	var v635 int32
	_ = v635
	var v640 int32
	_ = v640
	var v647 int32
	_ = v647
	var v652 int32
	_ = v652
	var v657 int32
	_ = v657
	var v658 int32
	_ = v658
	var v667 int32
	_ = v667
	var v675 int32
	_ = v675
	var v677 int32
	_ = v677
	var v681 int32
	_ = v681
	var v682 int32
	_ = v682
	var v688 int32
	_ = v688
	var v690 int32
	_ = v690
	var v697 int32
	_ = v697
	var v699 int32
	_ = v699
	var v703 int32
	_ = v703
	var v704 int32
	_ = v704
	var v713 int32
	_ = v713
	var v714 int32
	_ = v714
	var v717 int32
	_ = v717
	var v719 int32
	_ = v719
	var v721 int32
	_ = v721
	var v730 int32
	_ = v730
	var v732 int32
	_ = v732
	var v735 int32
	_ = v735
	var v740 int32
	_ = v740
	var v745 int32
	_ = v745
	var v746 int32
	_ = v746
	var v751 int32
	_ = v751
	var __phi751 int32
	_ = __phi751
	var v752 int32
	_ = v752
	var __phi752 int32
	_ = __phi752
	var v758 int32
	_ = v758
	var v761 int32
	_ = v761
	var v775 int32
	_ = v775
	var v784 int32
	_ = v784
	var v786 int32
	_ = v786
	var v789 int32
	_ = v789
	var v792 int32
	_ = v792
	var v794 int32
	_ = v794
	var v799 int32
	_ = v799
	var v806 int32
	_ = v806
	var v811 int32
	_ = v811
	var v833 int32
	_ = v833
	var v845 int32
	_ = v845
	var v847 int32
	_ = v847
	var v849 int32
	_ = v849
	var v853 int32
	_ = v853
	var v858 int32
	_ = v858
	var v859 int32
	_ = v859
	var v870 int32
	_ = v870
	var v873 int32
	_ = v873
	var v880 int32
	_ = v880
	var v885 int32
	_ = v885
	var v889 int32
	_ = v889
	var v891 int32
	_ = v891
	var v906 int32
	_ = v906
	var v908 int32
	_ = v908
	var v911 int32
	_ = v911
	var v914 int32
	_ = v914
	var v918 int32
	_ = v918
	var v928 int32
	_ = v928
	var v929 int32
	_ = v929
	var v935 int32
	_ = v935
	var v937 int32
	_ = v937
	var v940 int32
	_ = v940
	var v944 int32
	_ = v944
	var v950 int32
	_ = v950
	var v952 int32
	_ = v952
	var v953 int32
	_ = v953
	var v954 int32
	_ = v954
	var v955 int32
	_ = v955
	var v964 int32
	_ = v964
	var v966 int32
	_ = v966
	var v967 int32
	_ = v967
	var v968 int32
	_ = v968
	var v970 int32
	_ = v970
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
	if int32(0) <= v7 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	v15 = F_fflush(m, l0)
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		goto L5
	} else {
		goto L6
	}
L2:
	;
	goto L4
L3:
	;
	v14 = int32(1)
	goto L1
L4:
	;
	v14 = int32(0)
	goto L1
L5:
	;
	return int32(0)
L6:
	;
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v20 = m.T0[v19].(func(*base.Module, int32) int32)(m, l0)
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		goto L5
	} else {
		goto L7
	}
L7:
	;
	if v14 != 0 {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v22 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v22&int32(1) != 0 {
		goto L11
	} else {
		goto L12
	}
L9:
	;
	goto L10
L10:
	;
	goto L8
L11:
	;
	return v20 | v15
L12:
	;
	goto L13
L13:
	;
	F___lock(m, int32(9116552))
	mBase = m.M
	v27 = int32(9116556)
	goto L14
L14:
	;
	v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	if v29 == int32(0) {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	if v28 == int32(0) {
		goto L17
	} else {
		goto L18
	}
L16:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v29)+56)) = v28
	goto L15
L17:
	;
	v36 = *(*int32)(unsafe.Add(mBase, _c_F_fclose[0]))
	if v36 != l0 {
		goto L19
	} else {
		goto L20
	}
L18:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+52)) = v29
	goto L17
L19:
	;
	F___unlock(m, int32(9116552))
	mBase = m.M
	goto L21
L20:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_fclose[0])) = v28
	goto L19
L21:
	;
	v41 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	if v41 == int32(0) {
		goto L23
	} else {
		goto L24
	}
L22:
	;
	if l0 == int32(0) {
		goto L123
	} else {
		goto L124
	}
L23:
	;
	goto L22
L24:
	;
	v52 = int32(-8)
	v53 = v41 + v52
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v41+int32(-4))))
	v58 = v56 & v52
	v59 = v53 + v58
	if v56&int32(1) != 0 {
		v183 = v58
		v184 = v53
		goto L25
	} else {
		goto L26
	}
L25:
	;
	if base.Ui32(v59) <= base.Ui32(v184) {
		goto L23
	} else {
		goto L60
	}
L26:
	;
	if v56&int32(2) == int32(0) {
		goto L23
	} else {
		goto L27
	}
L27:
	;
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v53)))
	v67 = v53 - v66
	v69 = *(*int32)(unsafe.Add(mBase, _c_F_fclose[1]))
	if base.Ui32(v67) < base.Ui32(v69) {
		goto L23
	} else {
		goto L28
	}
L28:
	;
	v71 = v66 + v58
	v73 = *(*int32)(unsafe.Add(mBase, _c_F_fclose[2]))
	if v67 == v73 {
		goto L32
	} else {
		goto L33
	}
L29:
	;
	if v89 == int32(0) {
		v183 = v71
		v184 = v67
		goto L25
	} else {
		goto L48
	}
L30:
	;
	v142 = int32(0)
	goto L29
L31:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v78)+12)) = v75
	*(*int32)(unsafe.Add(mBase, uint32(v75)+8)) = v78
	v183 = v71
	v184 = v67
	goto L25
L32:
	;
	v123 = *(*int32)(unsafe.Add(mBase, uint32(v59)+4))
	v124 = int32(3)
	if v123&v124 != v124 {
		v183 = v71
		v184 = v67
		goto L25
	} else {
		goto L47
	}
L33:
	;
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v67)+12))
	if base.Ui32(int32(255)) < base.Ui32(v66) {
		goto L34
	} else {
		goto L35
	}
L34:
	;
	v89 = *(*int32)(unsafe.Add(mBase, uint32(v67)+24))
	if v75 == v67 {
		goto L37
	} else {
		goto L38
	}
L35:
	;
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v67)+8))
	if v75 != v78 {
		goto L31
	} else {
		goto L36
	}
L36:
	;
	v80 = int32(0)
	v82 = *(*int32)(unsafe.Add(mBase, _c_F_fclose[3]))
	*(*int32)(unsafe.Add(mBase, _c_F_fclose[3])) = v82 & base.I32_rotl(int32(-2), int32(base.Ui32(v66)>>(uint(int32(3))%32)))
	v183 = v71
	v184 = v67
	goto L25
L37:
	;
	v94 = *(*int32)(unsafe.Add(mBase, uint32(v67)+20))
	if v94 == int32(0) {
		goto L40
	} else {
		goto L41
	}
L38:
	;
	v91 = *(*int32)(unsafe.Add(mBase, uint32(v67)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v91)+12)) = v75
	*(*int32)(unsafe.Add(mBase, uint32(v75)+8)) = v91
	v142 = v75
	goto L29
L39:
	;
	__phi110 = v104
	__phi111 = v105
	v110 = __phi110
	v111 = __phi111
	goto L43
L40:
	;
	v99 = *(*int32)(unsafe.Add(mBase, uint32(v67)+16))
	if v99 == int32(0) {
		goto L30
	} else {
		goto L42
	}
L41:
	;
	v104 = v94
	v105 = v67 + int32(20)
	goto L39
L42:
	;
	v104 = v99
	v105 = v67 + int32(16)
	goto L39
L43:
	;
	v117 = *(*int32)(unsafe.Add(mBase, uint32(v110)+20))
	if v117 != 0 {
		__phi110 = v117
		__phi111 = v110 + int32(20)
		v110 = __phi110
		v111 = __phi111
		goto L43
	} else {
		goto L45
	}
L44:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v111))) = int32(0)
	v142 = v110
	goto L29
L45:
	;
	v120 = *(*int32)(unsafe.Add(mBase, uint32(v110)+16))
	if v120 != 0 {
		__phi110 = v120
		__phi111 = v110 + int32(16)
		v110 = __phi110
		v111 = __phi111
		goto L43
	} else {
		goto L46
	}
L46:
	;
	goto L44
L47:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_fclose[4])) = v71
	*(*int32)(unsafe.Add(mBase, uint32(v59)+4)) = v123 & int32(-2)
	*(*int32)(unsafe.Add(mBase, uint32(v67)+4)) = v71 | int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v59))) = v71
	goto L22
L48:
	;
	v151 = *(*int32)(unsafe.Add(mBase, uint32(v67)+28))
	v153 = v151 << (uint(int32(2)) % 32)
	v156 = *(*int32)(unsafe.Add(mBase, uint32(v153)+uint32(_c_F_fclose[5])))
	if v67 != v156 {
		goto L50
	} else {
		goto L51
	}
L49:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v142)+24)) = v89
	v173 = *(*int32)(unsafe.Add(mBase, uint32(v67)+16))
	if v173 == int32(0) {
		goto L57
	} else {
		goto L58
	}
L50:
	;
	v166 = *(*int32)(unsafe.Add(mBase, uint32(v89)+16))
	if v166 != v67 {
		goto L54
	} else {
		goto L55
	}
L51:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v153)+uint32(_c_F_fclose[5]))) = v142
	if v142 != 0 {
		goto L49
	} else {
		goto L52
	}
L52:
	;
	v159 = int32(0)
	v161 = *(*int32)(unsafe.Add(mBase, _c_F_fclose[6]))
	*(*int32)(unsafe.Add(mBase, _c_F_fclose[6])) = v161 & base.I32_rotl(int32(-2), v151)
	v183 = v71
	v184 = v67
	goto L25
L53:
	;
	if v142 == int32(0) {
		v183 = v71
		v184 = v67
		goto L25
	} else {
		goto L56
	}
L54:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v89)+20)) = v142
	goto L53
L55:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v89)+16)) = v142
	goto L53
L56:
	;
	goto L49
L57:
	;
	v178 = *(*int32)(unsafe.Add(mBase, uint32(v67)+20))
	if v178 == int32(0) {
		v183 = v71
		v184 = v67
		goto L25
	} else {
		goto L59
	}
L58:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v142)+16)) = v173
	*(*int32)(unsafe.Add(mBase, uint32(v173)+24)) = v142
	goto L57
L59:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v142)+20)) = v178
	*(*int32)(unsafe.Add(mBase, uint32(v178)+24)) = v142
	v183 = v71
	v184 = v67
	goto L25
L60:
	;
	v193 = *(*int32)(unsafe.Add(mBase, uint32(v59)+4))
	if v193&int32(1) == int32(0) {
		goto L23
	} else {
		goto L61
	}
L61:
	;
	if v193&int32(2) != 0 {
		goto L66
	} else {
		goto L67
	}
L62:
	;
	if base.Ui32(int32(255)) < base.Ui32(v359) {
		goto L100
	} else {
		goto L101
	}
L63:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v184)+4)) = v239 | int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v184+v239))) = v239
	if v184 != v223 {
		v359 = v239
		goto L62
	} else {
		goto L99
	}
L64:
	;
	if v256 == int32(0) {
		goto L63
	} else {
		goto L87
	}
L65:
	;
	v301 = int32(0)
	goto L64
L66:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v59)+4)) = v193 & int32(-2)
	*(*int32)(unsafe.Add(mBase, uint32(v184)+4)) = v183 | int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v184+v183))) = v183
	v359 = v183
	goto L62
L67:
	;
	v201 = *(*int32)(unsafe.Add(mBase, _c_F_fclose[7]))
	if v59 != v201 {
		goto L68
	} else {
		goto L69
	}
L68:
	;
	v223 = *(*int32)(unsafe.Add(mBase, _c_F_fclose[2]))
	if v59 != v223 {
		goto L71
	} else {
		goto L72
	}
L69:
	;
	v203 = int32(0)
	*(*int32)(unsafe.Add(mBase, _c_F_fclose[7])) = v184
	v207 = *(*int32)(unsafe.Add(mBase, _c_F_fclose[8]))
	v208 = v207 + v183
	*(*int32)(unsafe.Add(mBase, _c_F_fclose[8])) = v208
	*(*int32)(unsafe.Add(mBase, uint32(v184)+4)) = v208 | int32(1)
	v214 = *(*int32)(unsafe.Add(mBase, _c_F_fclose[2]))
	if v184 != v214 {
		goto L23
	} else {
		goto L70
	}
L70:
	;
	v216 = int32(0)
	*(*int32)(unsafe.Add(mBase, _c_F_fclose[4])) = v216
	*(*int32)(unsafe.Add(mBase, _c_F_fclose[2])) = v216
	goto L22
L71:
	;
	v239 = v193&int32(-8) + v183
	v240 = *(*int32)(unsafe.Add(mBase, uint32(v59)+12))
	if base.Ui32(int32(255)) < base.Ui32(v193) {
		goto L73
	} else {
		goto L74
	}
L72:
	;
	v225 = int32(0)
	*(*int32)(unsafe.Add(mBase, _c_F_fclose[2])) = v184
	v229 = *(*int32)(unsafe.Add(mBase, _c_F_fclose[4]))
	v230 = v229 + v183
	*(*int32)(unsafe.Add(mBase, _c_F_fclose[4])) = v230
	*(*int32)(unsafe.Add(mBase, uint32(v184)+4)) = v230 | int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v184+v230))) = v230
	goto L22
L73:
	;
	v256 = *(*int32)(unsafe.Add(mBase, uint32(v59)+24))
	if v240 == v59 {
		goto L77
	} else {
		goto L78
	}
L74:
	;
	v243 = *(*int32)(unsafe.Add(mBase, uint32(v59)+8))
	if v240 != v243 {
		goto L75
	} else {
		goto L76
	}
L75:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v243)+12)) = v240
	*(*int32)(unsafe.Add(mBase, uint32(v240)+8)) = v243
	goto L63
L76:
	;
	v245 = int32(0)
	v247 = *(*int32)(unsafe.Add(mBase, _c_F_fclose[3]))
	*(*int32)(unsafe.Add(mBase, _c_F_fclose[3])) = v247 & base.I32_rotl(int32(-2), int32(base.Ui32(v193)>>(uint(int32(3))%32)))
	goto L63
L77:
	;
	v261 = *(*int32)(unsafe.Add(mBase, uint32(v59)+20))
	if v261 == int32(0) {
		goto L80
	} else {
		goto L81
	}
L78:
	;
	v258 = *(*int32)(unsafe.Add(mBase, uint32(v59)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v258)+12)) = v240
	*(*int32)(unsafe.Add(mBase, uint32(v240)+8)) = v258
	v301 = v240
	goto L64
L79:
	;
	__phi277 = v271
	__phi278 = v272
	v277 = __phi277
	v278 = __phi278
	goto L83
L80:
	;
	v266 = *(*int32)(unsafe.Add(mBase, uint32(v59)+16))
	if v266 == int32(0) {
		goto L65
	} else {
		goto L82
	}
L81:
	;
	v271 = v261
	v272 = v59 + int32(20)
	goto L79
L82:
	;
	v271 = v266
	v272 = v59 + int32(16)
	goto L79
L83:
	;
	v284 = *(*int32)(unsafe.Add(mBase, uint32(v277)+20))
	if v284 != 0 {
		__phi277 = v284
		__phi278 = v277 + int32(20)
		v277 = __phi277
		v278 = __phi278
		goto L83
	} else {
		goto L85
	}
L84:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v278))) = int32(0)
	v301 = v277
	goto L64
L85:
	;
	v287 = *(*int32)(unsafe.Add(mBase, uint32(v277)+16))
	if v287 != 0 {
		__phi277 = v287
		__phi278 = v277 + int32(16)
		v277 = __phi277
		v278 = __phi278
		goto L83
	} else {
		goto L86
	}
L86:
	;
	goto L84
L87:
	;
	v310 = *(*int32)(unsafe.Add(mBase, uint32(v59)+28))
	v312 = v310 << (uint(int32(2)) % 32)
	v315 = *(*int32)(unsafe.Add(mBase, uint32(v312)+uint32(_c_F_fclose[5])))
	if v59 != v315 {
		goto L89
	} else {
		goto L90
	}
L88:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v301)+24)) = v256
	v332 = *(*int32)(unsafe.Add(mBase, uint32(v59)+16))
	if v332 == int32(0) {
		goto L96
	} else {
		goto L97
	}
L89:
	;
	v325 = *(*int32)(unsafe.Add(mBase, uint32(v256)+16))
	if v325 != v59 {
		goto L93
	} else {
		goto L94
	}
L90:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v312)+uint32(_c_F_fclose[5]))) = v301
	if v301 != 0 {
		goto L88
	} else {
		goto L91
	}
L91:
	;
	v318 = int32(0)
	v320 = *(*int32)(unsafe.Add(mBase, _c_F_fclose[6]))
	*(*int32)(unsafe.Add(mBase, _c_F_fclose[6])) = v320 & base.I32_rotl(int32(-2), v310)
	goto L63
L92:
	;
	if v301 == int32(0) {
		goto L63
	} else {
		goto L95
	}
L93:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v256)+20)) = v301
	goto L92
L94:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v256)+16)) = v301
	goto L92
L95:
	;
	goto L88
L96:
	;
	v337 = *(*int32)(unsafe.Add(mBase, uint32(v59)+20))
	if v337 == int32(0) {
		goto L63
	} else {
		goto L98
	}
L97:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v301)+16)) = v332
	*(*int32)(unsafe.Add(mBase, uint32(v332)+24)) = v301
	goto L96
L98:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v301)+20)) = v337
	*(*int32)(unsafe.Add(mBase, uint32(v337)+24)) = v301
	goto L63
L99:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_fclose[4])) = v239
	goto L22
L100:
	;
	if base.Ui32(int32(16777215)) < base.Ui32(v359) {
		v406 = int32(31)
		goto L105
	} else {
		goto L106
	}
L101:
	;
	v371 = v359 & int32(-8)
	v373 = v371 + int32(9128464)
	v375 = *(*int32)(unsafe.Add(mBase, _c_F_fclose[3]))
	v379 = int32(1) << (uint(int32(base.Ui32(v359)>>(uint(int32(3))%32))) % 32)
	if v375&v379 != 0 {
		goto L103
	} else {
		goto L104
	}
L102:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v371)+uint32(_c_F_fclose[9]))) = v184
	*(*int32)(unsafe.Add(mBase, uint32(v385)+12)) = v184
	*(*int32)(unsafe.Add(mBase, uint32(v184)+12)) = v373
	*(*int32)(unsafe.Add(mBase, uint32(v184)+8)) = v385
	goto L22
L103:
	;
	v384 = *(*int32)(unsafe.Add(mBase, uint32(v371)+uint32(_c_F_fclose[9])))
	v385 = v384
	goto L102
L104:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_fclose[3])) = v375 | v379
	v385 = v373
	goto L102
L105:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v184)+28)) = v406
	*(*int64)(unsafe.Add(mBase, uint32(v184)+16)) = int64(0)
	v411 = v406 << (uint(int32(2)) % 32)
	v415 = *(*int32)(unsafe.Add(mBase, _c_F_fclose[6]))
	v417 = int32(1) << (uint(v406) % 32)
	if v415&v417 != 0 {
		goto L110
	} else {
		goto L111
	}
L106:
	;
	v396 = base.I32_clz(int32(base.Ui32(v359) >> (uint(int32(8)) % 32)))
	v399 = int32(1)
	v406 = int32(base.Ui32(v359)>>(uint(int32(38)-v396)%32))&v399 - v396<<(uint(v399)%32) + int32(62)
	goto L105
L107:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v184+v478))) = v481
	*(*int32)(unsafe.Add(mBase, uint32(v184)+12)) = v480
	*(*int32)(unsafe.Add(mBase, uint32(v184+v476))) = v479
	v490 = int32(0)
	v492 = *(*int32)(unsafe.Add(mBase, _c_F_fclose[10]))
	v493 = int32(-1)
	v494 = v492 + v493
	if v494 != 0 {
		goto L119
	} else {
		goto L120
	}
L108:
	;
	v470 = *(*int32)(unsafe.Add(mBase, uint32(v440)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v470)+12)) = v184
	*(*int32)(unsafe.Add(mBase, uint32(v440)+8)) = v184
	v476 = int32(24)
	v478 = int32(8)
	v479 = int32(0)
	v480 = v440
	v481 = v470
	goto L107
L109:
	;
	v476 = v461
	v478 = v463
	v479 = v184
	v480 = v184
	v481 = v466
	goto L107
L110:
	;
	if v406 == int32(31) {
		goto L112
	} else {
		goto L113
	}
L111:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_fclose[6])) = v415 | v417
	*(*int32)(unsafe.Add(mBase, uint32(v411)+uint32(_c_F_fclose[5]))) = v184
	v461 = int32(8)
	v463 = int32(24)
	v466 = v411 + int32(9128728)
	goto L109
L112:
	;
	v432 = int32(0)
	goto L114
L113:
	;
	v432 = int32(25) - int32(base.Ui32(v406)>>(uint(int32(1))%32))
	goto L114
L114:
	;
	v434 = *(*int32)(unsafe.Add(mBase, uint32(v411)+uint32(_c_F_fclose[5])))
	v437 = v359 << (uint(v432) % 32)
	v440 = v434
	goto L115
L115:
	;
	v444 = *(*int32)(unsafe.Add(mBase, uint32(v440)+4))
	if v444&int32(-8) == v359 {
		goto L108
	} else {
		goto L117
	}
L116:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v454+int32(16)))) = v184
	v461 = int32(8)
	v463 = int32(24)
	v466 = v440
	goto L109
L117:
	;
	v454 = v440 + int32(base.Ui32(v437)>>(uint(int32(29))%32))&int32(4)
	v455 = *(*int32)(unsafe.Add(mBase, uint32(v454)+16))
	if v455 != 0 {
		v437 = v437 << (uint(int32(1)) % 32)
		v440 = v455
		goto L115
	} else {
		goto L118
	}
L118:
	;
	goto L116
L119:
	;
	v496 = v494
	goto L121
L120:
	;
	v496 = v493
	goto L121
L121:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_fclose[10])) = v496
	goto L23
L122:
	;
	goto L11
L123:
	;
	goto L122
L124:
	;
	v526 = int32(-8)
	v527 = l0 + v526
	v530 = *(*int32)(unsafe.Add(mBase, uint32(l0+int32(-4))))
	v532 = v530 & v526
	v533 = v527 + v532
	if v530&int32(1) != 0 {
		v657 = v532
		v658 = v527
		goto L125
	} else {
		goto L126
	}
L125:
	;
	if base.Ui32(v533) <= base.Ui32(v658) {
		goto L123
	} else {
		goto L160
	}
L126:
	;
	if v530&int32(2) == int32(0) {
		goto L123
	} else {
		goto L127
	}
L127:
	;
	v540 = *(*int32)(unsafe.Add(mBase, uint32(v527)))
	v541 = v527 - v540
	v543 = *(*int32)(unsafe.Add(mBase, _c_F_fclose[1]))
	if base.Ui32(v541) < base.Ui32(v543) {
		goto L123
	} else {
		goto L128
	}
L128:
	;
	v545 = v540 + v532
	v547 = *(*int32)(unsafe.Add(mBase, _c_F_fclose[2]))
	if v541 == v547 {
		goto L132
	} else {
		goto L133
	}
L129:
	;
	if v563 == int32(0) {
		v657 = v545
		v658 = v541
		goto L125
	} else {
		goto L148
	}
L130:
	;
	v616 = int32(0)
	goto L129
L131:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v552)+12)) = v549
	*(*int32)(unsafe.Add(mBase, uint32(v549)+8)) = v552
	v657 = v545
	v658 = v541
	goto L125
L132:
	;
	v597 = *(*int32)(unsafe.Add(mBase, uint32(v533)+4))
	v598 = int32(3)
	if v597&v598 != v598 {
		v657 = v545
		v658 = v541
		goto L125
	} else {
		goto L147
	}
L133:
	;
	v549 = *(*int32)(unsafe.Add(mBase, uint32(v541)+12))
	if base.Ui32(int32(255)) < base.Ui32(v540) {
		goto L134
	} else {
		goto L135
	}
L134:
	;
	v563 = *(*int32)(unsafe.Add(mBase, uint32(v541)+24))
	if v549 == v541 {
		goto L137
	} else {
		goto L138
	}
L135:
	;
	v552 = *(*int32)(unsafe.Add(mBase, uint32(v541)+8))
	if v549 != v552 {
		goto L131
	} else {
		goto L136
	}
L136:
	;
	v554 = int32(0)
	v556 = *(*int32)(unsafe.Add(mBase, _c_F_fclose[3]))
	*(*int32)(unsafe.Add(mBase, _c_F_fclose[3])) = v556 & base.I32_rotl(int32(-2), int32(base.Ui32(v540)>>(uint(int32(3))%32)))
	v657 = v545
	v658 = v541
	goto L125
L137:
	;
	v568 = *(*int32)(unsafe.Add(mBase, uint32(v541)+20))
	if v568 == int32(0) {
		goto L140
	} else {
		goto L141
	}
L138:
	;
	v565 = *(*int32)(unsafe.Add(mBase, uint32(v541)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v565)+12)) = v549
	*(*int32)(unsafe.Add(mBase, uint32(v549)+8)) = v565
	v616 = v549
	goto L129
L139:
	;
	__phi584 = v578
	__phi585 = v579
	v584 = __phi584
	v585 = __phi585
	goto L143
L140:
	;
	v573 = *(*int32)(unsafe.Add(mBase, uint32(v541)+16))
	if v573 == int32(0) {
		goto L130
	} else {
		goto L142
	}
L141:
	;
	v578 = v568
	v579 = v541 + int32(20)
	goto L139
L142:
	;
	v578 = v573
	v579 = v541 + int32(16)
	goto L139
L143:
	;
	v591 = *(*int32)(unsafe.Add(mBase, uint32(v584)+20))
	if v591 != 0 {
		__phi584 = v591
		__phi585 = v584 + int32(20)
		v584 = __phi584
		v585 = __phi585
		goto L143
	} else {
		goto L145
	}
L144:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v585))) = int32(0)
	v616 = v584
	goto L129
L145:
	;
	v594 = *(*int32)(unsafe.Add(mBase, uint32(v584)+16))
	if v594 != 0 {
		__phi584 = v594
		__phi585 = v584 + int32(16)
		v584 = __phi584
		v585 = __phi585
		goto L143
	} else {
		goto L146
	}
L146:
	;
	goto L144
L147:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_fclose[4])) = v545
	*(*int32)(unsafe.Add(mBase, uint32(v533)+4)) = v597 & int32(-2)
	*(*int32)(unsafe.Add(mBase, uint32(v541)+4)) = v545 | int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v533))) = v545
	goto L122
L148:
	;
	v625 = *(*int32)(unsafe.Add(mBase, uint32(v541)+28))
	v627 = v625 << (uint(int32(2)) % 32)
	v630 = *(*int32)(unsafe.Add(mBase, uint32(v627)+uint32(_c_F_fclose[5])))
	if v541 != v630 {
		goto L150
	} else {
		goto L151
	}
L149:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v616)+24)) = v563
	v647 = *(*int32)(unsafe.Add(mBase, uint32(v541)+16))
	if v647 == int32(0) {
		goto L157
	} else {
		goto L158
	}
L150:
	;
	v640 = *(*int32)(unsafe.Add(mBase, uint32(v563)+16))
	if v640 != v541 {
		goto L154
	} else {
		goto L155
	}
L151:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v627)+uint32(_c_F_fclose[5]))) = v616
	if v616 != 0 {
		goto L149
	} else {
		goto L152
	}
L152:
	;
	v633 = int32(0)
	v635 = *(*int32)(unsafe.Add(mBase, _c_F_fclose[6]))
	*(*int32)(unsafe.Add(mBase, _c_F_fclose[6])) = v635 & base.I32_rotl(int32(-2), v625)
	v657 = v545
	v658 = v541
	goto L125
L153:
	;
	if v616 == int32(0) {
		v657 = v545
		v658 = v541
		goto L125
	} else {
		goto L156
	}
L154:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v563)+20)) = v616
	goto L153
L155:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v563)+16)) = v616
	goto L153
L156:
	;
	goto L149
L157:
	;
	v652 = *(*int32)(unsafe.Add(mBase, uint32(v541)+20))
	if v652 == int32(0) {
		v657 = v545
		v658 = v541
		goto L125
	} else {
		goto L159
	}
L158:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v616)+16)) = v647
	*(*int32)(unsafe.Add(mBase, uint32(v647)+24)) = v616
	goto L157
L159:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v616)+20)) = v652
	*(*int32)(unsafe.Add(mBase, uint32(v652)+24)) = v616
	v657 = v545
	v658 = v541
	goto L125
L160:
	;
	v667 = *(*int32)(unsafe.Add(mBase, uint32(v533)+4))
	if v667&int32(1) == int32(0) {
		goto L123
	} else {
		goto L161
	}
L161:
	;
	if v667&int32(2) != 0 {
		goto L166
	} else {
		goto L167
	}
L162:
	;
	if base.Ui32(int32(255)) < base.Ui32(v833) {
		goto L200
	} else {
		goto L201
	}
L163:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v658)+4)) = v713 | int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v658+v713))) = v713
	if v658 != v697 {
		v833 = v713
		goto L162
	} else {
		goto L199
	}
L164:
	;
	if v730 == int32(0) {
		goto L163
	} else {
		goto L187
	}
L165:
	;
	v775 = int32(0)
	goto L164
L166:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v533)+4)) = v667 & int32(-2)
	*(*int32)(unsafe.Add(mBase, uint32(v658)+4)) = v657 | int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v658+v657))) = v657
	v833 = v657
	goto L162
L167:
	;
	v675 = *(*int32)(unsafe.Add(mBase, _c_F_fclose[7]))
	if v533 != v675 {
		goto L168
	} else {
		goto L169
	}
L168:
	;
	v697 = *(*int32)(unsafe.Add(mBase, _c_F_fclose[2]))
	if v533 != v697 {
		goto L171
	} else {
		goto L172
	}
L169:
	;
	v677 = int32(0)
	*(*int32)(unsafe.Add(mBase, _c_F_fclose[7])) = v658
	v681 = *(*int32)(unsafe.Add(mBase, _c_F_fclose[8]))
	v682 = v681 + v657
	*(*int32)(unsafe.Add(mBase, _c_F_fclose[8])) = v682
	*(*int32)(unsafe.Add(mBase, uint32(v658)+4)) = v682 | int32(1)
	v688 = *(*int32)(unsafe.Add(mBase, _c_F_fclose[2]))
	if v658 != v688 {
		goto L123
	} else {
		goto L170
	}
L170:
	;
	v690 = int32(0)
	*(*int32)(unsafe.Add(mBase, _c_F_fclose[4])) = v690
	*(*int32)(unsafe.Add(mBase, _c_F_fclose[2])) = v690
	goto L122
L171:
	;
	v713 = v667&int32(-8) + v657
	v714 = *(*int32)(unsafe.Add(mBase, uint32(v533)+12))
	if base.Ui32(int32(255)) < base.Ui32(v667) {
		goto L173
	} else {
		goto L174
	}
L172:
	;
	v699 = int32(0)
	*(*int32)(unsafe.Add(mBase, _c_F_fclose[2])) = v658
	v703 = *(*int32)(unsafe.Add(mBase, _c_F_fclose[4]))
	v704 = v703 + v657
	*(*int32)(unsafe.Add(mBase, _c_F_fclose[4])) = v704
	*(*int32)(unsafe.Add(mBase, uint32(v658)+4)) = v704 | int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v658+v704))) = v704
	goto L122
L173:
	;
	v730 = *(*int32)(unsafe.Add(mBase, uint32(v533)+24))
	if v714 == v533 {
		goto L177
	} else {
		goto L178
	}
L174:
	;
	v717 = *(*int32)(unsafe.Add(mBase, uint32(v533)+8))
	if v714 != v717 {
		goto L175
	} else {
		goto L176
	}
L175:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v717)+12)) = v714
	*(*int32)(unsafe.Add(mBase, uint32(v714)+8)) = v717
	goto L163
L176:
	;
	v719 = int32(0)
	v721 = *(*int32)(unsafe.Add(mBase, _c_F_fclose[3]))
	*(*int32)(unsafe.Add(mBase, _c_F_fclose[3])) = v721 & base.I32_rotl(int32(-2), int32(base.Ui32(v667)>>(uint(int32(3))%32)))
	goto L163
L177:
	;
	v735 = *(*int32)(unsafe.Add(mBase, uint32(v533)+20))
	if v735 == int32(0) {
		goto L180
	} else {
		goto L181
	}
L178:
	;
	v732 = *(*int32)(unsafe.Add(mBase, uint32(v533)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v732)+12)) = v714
	*(*int32)(unsafe.Add(mBase, uint32(v714)+8)) = v732
	v775 = v714
	goto L164
L179:
	;
	__phi751 = v745
	__phi752 = v746
	v751 = __phi751
	v752 = __phi752
	goto L183
L180:
	;
	v740 = *(*int32)(unsafe.Add(mBase, uint32(v533)+16))
	if v740 == int32(0) {
		goto L165
	} else {
		goto L182
	}
L181:
	;
	v745 = v735
	v746 = v533 + int32(20)
	goto L179
L182:
	;
	v745 = v740
	v746 = v533 + int32(16)
	goto L179
L183:
	;
	v758 = *(*int32)(unsafe.Add(mBase, uint32(v751)+20))
	if v758 != 0 {
		__phi751 = v758
		__phi752 = v751 + int32(20)
		v751 = __phi751
		v752 = __phi752
		goto L183
	} else {
		goto L185
	}
L184:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v752))) = int32(0)
	v775 = v751
	goto L164
L185:
	;
	v761 = *(*int32)(unsafe.Add(mBase, uint32(v751)+16))
	if v761 != 0 {
		__phi751 = v761
		__phi752 = v751 + int32(16)
		v751 = __phi751
		v752 = __phi752
		goto L183
	} else {
		goto L186
	}
L186:
	;
	goto L184
L187:
	;
	v784 = *(*int32)(unsafe.Add(mBase, uint32(v533)+28))
	v786 = v784 << (uint(int32(2)) % 32)
	v789 = *(*int32)(unsafe.Add(mBase, uint32(v786)+uint32(_c_F_fclose[5])))
	if v533 != v789 {
		goto L189
	} else {
		goto L190
	}
L188:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v775)+24)) = v730
	v806 = *(*int32)(unsafe.Add(mBase, uint32(v533)+16))
	if v806 == int32(0) {
		goto L196
	} else {
		goto L197
	}
L189:
	;
	v799 = *(*int32)(unsafe.Add(mBase, uint32(v730)+16))
	if v799 != v533 {
		goto L193
	} else {
		goto L194
	}
L190:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v786)+uint32(_c_F_fclose[5]))) = v775
	if v775 != 0 {
		goto L188
	} else {
		goto L191
	}
L191:
	;
	v792 = int32(0)
	v794 = *(*int32)(unsafe.Add(mBase, _c_F_fclose[6]))
	*(*int32)(unsafe.Add(mBase, _c_F_fclose[6])) = v794 & base.I32_rotl(int32(-2), v784)
	goto L163
L192:
	;
	if v775 == int32(0) {
		goto L163
	} else {
		goto L195
	}
L193:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v730)+20)) = v775
	goto L192
L194:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v730)+16)) = v775
	goto L192
L195:
	;
	goto L188
L196:
	;
	v811 = *(*int32)(unsafe.Add(mBase, uint32(v533)+20))
	if v811 == int32(0) {
		goto L163
	} else {
		goto L198
	}
L197:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v775)+16)) = v806
	*(*int32)(unsafe.Add(mBase, uint32(v806)+24)) = v775
	goto L196
L198:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v775)+20)) = v811
	*(*int32)(unsafe.Add(mBase, uint32(v811)+24)) = v775
	goto L163
L199:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_fclose[4])) = v713
	goto L122
L200:
	;
	if base.Ui32(int32(16777215)) < base.Ui32(v833) {
		v880 = int32(31)
		goto L205
	} else {
		goto L206
	}
L201:
	;
	v845 = v833 & int32(-8)
	v847 = v845 + int32(9128464)
	v849 = *(*int32)(unsafe.Add(mBase, _c_F_fclose[3]))
	v853 = int32(1) << (uint(int32(base.Ui32(v833)>>(uint(int32(3))%32))) % 32)
	if v849&v853 != 0 {
		goto L203
	} else {
		goto L204
	}
L202:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v845)+uint32(_c_F_fclose[9]))) = v658
	*(*int32)(unsafe.Add(mBase, uint32(v859)+12)) = v658
	*(*int32)(unsafe.Add(mBase, uint32(v658)+12)) = v847
	*(*int32)(unsafe.Add(mBase, uint32(v658)+8)) = v859
	goto L122
L203:
	;
	v858 = *(*int32)(unsafe.Add(mBase, uint32(v845)+uint32(_c_F_fclose[9])))
	v859 = v858
	goto L202
L204:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_fclose[3])) = v849 | v853
	v859 = v847
	goto L202
L205:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v658)+28)) = v880
	*(*int64)(unsafe.Add(mBase, uint32(v658)+16)) = int64(0)
	v885 = v880 << (uint(int32(2)) % 32)
	v889 = *(*int32)(unsafe.Add(mBase, _c_F_fclose[6]))
	v891 = int32(1) << (uint(v880) % 32)
	if v889&v891 != 0 {
		goto L210
	} else {
		goto L211
	}
L206:
	;
	v870 = base.I32_clz(int32(base.Ui32(v833) >> (uint(int32(8)) % 32)))
	v873 = int32(1)
	v880 = int32(base.Ui32(v833)>>(uint(int32(38)-v870)%32))&v873 - v870<<(uint(v873)%32) + int32(62)
	goto L205
L207:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v658+v952))) = v955
	*(*int32)(unsafe.Add(mBase, uint32(v658)+12)) = v954
	*(*int32)(unsafe.Add(mBase, uint32(v658+v950))) = v953
	v964 = int32(0)
	v966 = *(*int32)(unsafe.Add(mBase, _c_F_fclose[10]))
	v967 = int32(-1)
	v968 = v966 + v967
	if v968 != 0 {
		goto L219
	} else {
		goto L220
	}
L208:
	;
	v944 = *(*int32)(unsafe.Add(mBase, uint32(v914)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v944)+12)) = v658
	*(*int32)(unsafe.Add(mBase, uint32(v914)+8)) = v658
	v950 = int32(24)
	v952 = int32(8)
	v953 = int32(0)
	v954 = v914
	v955 = v944
	goto L207
L209:
	;
	v950 = v935
	v952 = v937
	v953 = v658
	v954 = v658
	v955 = v940
	goto L207
L210:
	;
	if v880 == int32(31) {
		goto L212
	} else {
		goto L213
	}
L211:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_fclose[6])) = v889 | v891
	*(*int32)(unsafe.Add(mBase, uint32(v885)+uint32(_c_F_fclose[5]))) = v658
	v935 = int32(8)
	v937 = int32(24)
	v940 = v885 + int32(9128728)
	goto L209
L212:
	;
	v906 = int32(0)
	goto L214
L213:
	;
	v906 = int32(25) - int32(base.Ui32(v880)>>(uint(int32(1))%32))
	goto L214
L214:
	;
	v908 = *(*int32)(unsafe.Add(mBase, uint32(v885)+uint32(_c_F_fclose[5])))
	v911 = v833 << (uint(v906) % 32)
	v914 = v908
	goto L215
L215:
	;
	v918 = *(*int32)(unsafe.Add(mBase, uint32(v914)+4))
	if v918&int32(-8) == v833 {
		goto L208
	} else {
		goto L217
	}
L216:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v928+int32(16)))) = v658
	v935 = int32(8)
	v937 = int32(24)
	v940 = v914
	goto L209
L217:
	;
	v928 = v914 + int32(base.Ui32(v911)>>(uint(int32(29))%32))&int32(4)
	v929 = *(*int32)(unsafe.Add(mBase, uint32(v928)+16))
	if v929 != 0 {
		v911 = v911 << (uint(int32(1)) % 32)
		v914 = v929
		goto L215
	} else {
		goto L218
	}
L218:
	;
	goto L216
L219:
	;
	v970 = v968
	goto L221
L220:
	;
	v970 = v967
	goto L221
L221:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_fclose[10])) = v970
	goto L123
}
func F_fcntl(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v43 int32
	_ = v43
	var v51 int32
	_ = v51
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v65 int32
	_ = v65
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
	var v90 int32
	_ = v90
	var v97 int32
	_ = v97
	var v100 int32
	_ = v100
	var v103 int32
	_ = v103
	var v108 int32
	_ = v108
	var v109 int64
	_ = v109
	var v114 int32
	_ = v114
	var v119 int32
	_ = v119
	var v124 int32
	_ = v124
	var v130 int32
	_ = v130
	var v135 int32
	_ = v135
	var v139 int32
	_ = v139
	var v149 int32
	_ = v149
	var v152 int32
	_ = v152
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	v4 = int32(0)
	v7 = m.G0
	v9 = v7 - int32(128)
	m.G0 = v9
	switch l1 + int32(-1) {
	case 0, 2:
		v19 = v4
	case 1:
		*(*int32)(unsafe.Add(mBase, uint32(v9)+120)) = l2 + int32(4)
		v18 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
		v19 = v18
	default:
		if l1 == int32(9) {
			v19 = v4
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v9)+120)) = l2 + int32(4)
			v18 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
			v19 = v18
		}
	}
	if base.Ui32(int32(16)) < base.Ui32(l1) {
		if l1 == int32(1030) {
			v109 = base.I64_extend_i32_u(v19)
			*(*int64)(unsafe.Add(mBase, uint32(v9)+96)) = v109
			v114 = m.Env.X__syscall_fcntl64(m, l0, int32(1030), v9+int32(96))
			mBase = m.M
			if v114 == int32(-28) {
				*(*int64)(unsafe.Add(mBase, uint32(v9)+80)) = int64(0)
				v130 = m.Env.X__syscall_fcntl64(m, l0, int32(1030), v9+int32(80))
				mBase = m.M
				if v130 == int32(-28) {
					*(*int64)(unsafe.Add(mBase, uint32(v9)+64)) = v109
					v149 = m.Env.X__syscall_fcntl64(m, l0, int32(0), v9+int32(64))
					mBase = m.M
					if base.Ui32(v149) < base.Ui32(int32(-4095)) {
						v157 = v149
					} else {
						v152 = F___errno_location(m)
						mBase = m.M
						*(*int32)(unsafe.Add(mBase, uint32(v152))) = int32(0) - v149
						v157 = int32(-1)
					}
					v158 = v157
				} else {
					if v130 < int32(0) {
					} else {
						v135 = m.Wasi_snapshot_preview1.Fd_close(m, v130)
						mBase = m.M
					}
					v139 = F___errno_location(m)
					mBase = m.M
					*(*int32)(unsafe.Add(mBase, uint32(v139))) = int32(28)
					v158 = int32(-1)
				}
			} else {
				if base.Ui32(v114) < base.Ui32(int32(-4095)) {
					v124 = v114
				} else {
					v119 = F___errno_location(m)
					mBase = m.M
					*(*int32)(unsafe.Add(mBase, uint32(v119))) = int32(0) - v114
					v124 = int32(-1)
				}
				v158 = v124
			}
		} else {
			if l1 == int32(4) {
				v97 = v19 | int32(32768)
			} else {
				v97 = v19
			}
			*(*int64)(unsafe.Add(mBase, uint32(v9))) = base.I64_extend_i32_u(v97)
			v100 = m.Env.X__syscall_fcntl64(m, l0, l1, v9)
			mBase = m.M
			if base.Ui32(v100) < base.Ui32(int32(-4095)) {
				v108 = v100
			} else {
				v103 = F___errno_location(m)
				mBase = m.M
				*(*int32)(unsafe.Add(mBase, uint32(v103))) = int32(0) - v100
				v108 = int32(-1)
			}
			v158 = v108
		}
	} else {
		if int32(1)<<(uint(l1)%32)&int32(110592) != 0 {
			*(*int64)(unsafe.Add(mBase, uint32(v9)+112)) = base.I64_extend_i32_u(v19)
			v82 = m.Env.X__syscall_fcntl64(m, l0, l1, v9+int32(112))
			mBase = m.M
			if base.Ui32(v82) < base.Ui32(int32(-4095)) {
				v90 = v82
			} else {
				v85 = F___errno_location(m)
				mBase = m.M
				*(*int32)(unsafe.Add(mBase, uint32(v85))) = int32(0) - v82
				v90 = int32(-1)
			}
			v158 = v90
		} else {
			if l1 == int32(9) {
				*(*int64)(unsafe.Add(mBase, uint32(v9)+48)) = base.I64_extend_i32_u(v9 + int32(120))
				v51 = m.Env.X__syscall_fcntl64(m, l0, int32(16), v9+int32(48))
				mBase = m.M
				if v51 != int32(-28) {
					v60 = v51
				} else {
					*(*int64)(unsafe.Add(mBase, uint32(v9)+32)) = base.I64_extend_i32_u(v19)
					v59 = m.Env.X__syscall_fcntl64(m, l0, int32(9), v9+int32(32))
					mBase = m.M
					v60 = v59
				}
				if v60 == int32(0) {
					v72 = *(*int32)(unsafe.Add(mBase, uint32(v9)+124))
					v74 = *(*int32)(unsafe.Add(mBase, uint32(v9)+120))
					if v74 == int32(2) {
						v77 = int32(0) - v72
					} else {
						v77 = v72
					}
					v158 = v77
				} else {
					if base.Ui32(v60) < base.Ui32(int32(-4095)) {
						v70 = v60
					} else {
						v65 = F___errno_location(m)
						mBase = m.M
						*(*int32)(unsafe.Add(mBase, uint32(v65))) = int32(0) - v60
						v70 = int32(-1)
					}
					v158 = v70
				}
			} else {
				if l1 != int32(14) {
					if l1 == int32(1030) {
						v109 = base.I64_extend_i32_u(v19)
						*(*int64)(unsafe.Add(mBase, uint32(v9)+96)) = v109
						v114 = m.Env.X__syscall_fcntl64(m, l0, int32(1030), v9+int32(96))
						mBase = m.M
						if v114 == int32(-28) {
							*(*int64)(unsafe.Add(mBase, uint32(v9)+80)) = int64(0)
							v130 = m.Env.X__syscall_fcntl64(m, l0, int32(1030), v9+int32(80))
							mBase = m.M
							if v130 == int32(-28) {
								*(*int64)(unsafe.Add(mBase, uint32(v9)+64)) = v109
								v149 = m.Env.X__syscall_fcntl64(m, l0, int32(0), v9+int32(64))
								mBase = m.M
								if base.Ui32(v149) < base.Ui32(int32(-4095)) {
									v157 = v149
								} else {
									v152 = F___errno_location(m)
									mBase = m.M
									*(*int32)(unsafe.Add(mBase, uint32(v152))) = int32(0) - v149
									v157 = int32(-1)
								}
								v158 = v157
							} else {
								if v130 < int32(0) {
								} else {
									v135 = m.Wasi_snapshot_preview1.Fd_close(m, v130)
									mBase = m.M
								}
								v139 = F___errno_location(m)
								mBase = m.M
								*(*int32)(unsafe.Add(mBase, uint32(v139))) = int32(28)
								v158 = int32(-1)
							}
						} else {
							if base.Ui32(v114) < base.Ui32(int32(-4095)) {
								v124 = v114
							} else {
								v119 = F___errno_location(m)
								mBase = m.M
								*(*int32)(unsafe.Add(mBase, uint32(v119))) = int32(0) - v114
								v124 = int32(-1)
							}
							v158 = v124
						}
					} else {
						if l1 == int32(4) {
							v97 = v19 | int32(32768)
						} else {
							v97 = v19
						}
						*(*int64)(unsafe.Add(mBase, uint32(v9))) = base.I64_extend_i32_u(v97)
						v100 = m.Env.X__syscall_fcntl64(m, l0, l1, v9)
						mBase = m.M
						if base.Ui32(v100) < base.Ui32(int32(-4095)) {
							v108 = v100
						} else {
							v103 = F___errno_location(m)
							mBase = m.M
							*(*int32)(unsafe.Add(mBase, uint32(v103))) = int32(0) - v100
							v108 = int32(-1)
						}
						v158 = v108
					}
				} else {
					*(*int64)(unsafe.Add(mBase, uint32(v9)+16)) = base.I64_extend_i32_u(v19)
					v35 = m.Env.X__syscall_fcntl64(m, l0, int32(14), v9+int32(16))
					mBase = m.M
					if base.Ui32(v35) < base.Ui32(int32(-4095)) {
						v43 = v35
					} else {
						v38 = F___errno_location(m)
						mBase = m.M
						*(*int32)(unsafe.Add(mBase, uint32(v38))) = int32(0) - v35
						v43 = int32(-1)
					}
					v158 = v43
				}
			}
		}
	}
	m.G0 = v9 + int32(128)
	return v158
}
func F_feedReplicationBufferWithObject(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v17 int64
	_ = v17
	var v26 int32
	_ = v26
	var v30 int32
	_ = v30
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int64
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v58 int32
	_ = v58
	var v62 int32
	_ = v62
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
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
	var v85 int32
	_ = v85
	v6 = m.G0
	v8 = v6 - int32(32)
	m.G0 = v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v11 = F_objectGetVal(m, l0)
	mBase = m.M
	if v10&int32(240) != int32(16) {
		v62 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11+int32(-1)))))
		switch v62 & int32(7) {
		case 0:
			v79 = int32(base.Ui32(v62) >> (uint(int32(3)) % 32))
		case 1:
			v69 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11+int32(-3)))))
			v79 = v69
		case 2:
			v72 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v11+int32(-5)))))
			v79 = v72
		case 3:
			v75 = *(*int32)(unsafe.Add(mBase, uint32(v11+int32(-9))))
			v79 = v75
		case 4:
			v78 = *(*int32)(unsafe.Add(mBase, uint32(v11+int32(-17))))
			v79 = v78
		default:
			v79 = int32(0)
		}
		v80 = F_objectGetVal(m, l0)
		mBase = m.M
		v81 = v80
		v82 = v79
	} else {
		v17 = base.I64_extend_i32_s(v11)
		if v17 <= int64(-1) {
			v26 = int32(45)
			*(*uint8)(unsafe.Add(mBase, uint32(v8))) = uint8(v26)
			v30 = int32(1)
			v35 = v8 + v30
			v36 = int32(20)
			v37 = int64(0) - v17
			v38 = v30
		} else {
			v35 = v8
			v36 = int32(21)
			v37 = v17
			v38 = int32(0)
		}
		v39 = F_ull2string(m, v35, v36, v37)
		mBase = m.M
		if v39 == int32(0) {
			v58 = int32(0)
		} else {
			v58 = v39 + v38
		}
		v81 = v8
		v82 = v58
	}
	F_feedReplicationBuffer(m, v81, v82)
	mBase = m.M
	v85 = m.ExcPending
	if v85 != 0 {
		return
	} else {
		m.G0 = v8 + int32(32)
		return
	}
}
func F_feof(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
	if int32(-1) < v4 {
		v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		v12 = v9
	} else {
		v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		v12 = v7
	}
	return int32(base.Ui32(v12)>>(uint(int32(4))%32)) & int32(1)
}
func F_fflush(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
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
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v60 int32
	_ = v60
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v86 int32
	_ = v86
	var v87 int64
	_ = v87
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v92 int64
	_ = v92
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	if l0 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v60 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
	if int32(0) <= v60 {
		goto L28
	} else {
		goto L29
	}
L2:
	;
	v5 = int32(0)
	v7 = *(*int32)(unsafe.Add(mBase, _c_F_fflush[0]))
	if v7 == v5 {
		v16 = v5
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v17 = int32(0)
	v18 = *(*int32)(unsafe.Add(mBase, _c_F_fflush[1]))
	if v18 == v17 {
		v26 = v16
		goto L7
	} else {
		goto L8
	}
L4:
	;
	v11 = *(*int32)(unsafe.Add(mBase, _c_F_fflush[0]))
	v12 = F_fflush(m, v11)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	return int32(0)
L6:
	;
	v16 = v12
	goto L3
L7:
	;
	F___lock(m, int32(9116552))
	mBase = m.M
	goto L11
L8:
	;
	v22 = *(*int32)(unsafe.Add(mBase, _c_F_fflush[1]))
	v23 = F_fflush(m, v22)
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		goto L5
	} else {
		goto L9
	}
L9:
	;
	v26 = v23 | v16
	goto L7
L10:
	;
	F___unlock(m, int32(9116552))
	mBase = m.M
	goto L26
L11:
	;
	v30 = *(*int32)(unsafe.Add(mBase, _c_F_fflush[2]))
	if v30 == int32(0) {
		v54 = v26
		goto L10
	} else {
		goto L12
	}
L12:
	;
	v33 = v30
	v34 = v26
	goto L13
L13:
	;
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v33)+76))
	if int32(0) <= v37 {
		goto L16
	} else {
		goto L17
	}
L14:
	;
	v54 = v51
	goto L10
L15:
	;
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v33)+20))
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v33)+28))
	if v45 == v46 {
		v51 = v34
		goto L19
	} else {
		goto L20
	}
L16:
	;
	goto L18
L17:
	;
	v44 = int32(1)
	goto L15
L18:
	;
	v44 = int32(0)
	goto L15
L19:
	;
	if v44 != 0 {
		goto L22
	} else {
		goto L23
	}
L20:
	;
	v48 = F_fflush(m, v33)
	mBase = m.M
	v49 = m.ExcPending
	if v49 != 0 {
		goto L5
	} else {
		goto L21
	}
L21:
	;
	v51 = v48 | v34
	goto L19
L22:
	;
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v33)+56))
	if v52 != 0 {
		v33 = v52
		v34 = v51
		goto L13
	} else {
		goto L25
	}
L23:
	;
	goto L24
L24:
	;
	goto L22
L25:
	;
	goto L14
L26:
	;
	return v54
L27:
	;
	v68 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v69 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v68 == v69 {
		goto L33
	} else {
		goto L34
	}
L28:
	;
	goto L30
L29:
	;
	v67 = int32(1)
	goto L27
L30:
	;
	v67 = int32(0)
	goto L27
L31:
	;
	return v98
L32:
	;
	goto L42
L33:
	;
	v80 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v81 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v80 == v81 {
		goto L38
	} else {
		goto L39
	}
L34:
	;
	v71 = int32(0)
	v73 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v74 = m.T0[v73].(func(*base.Module, int32, int32, int32) int32)(m, l0, v71, v71)
	mBase = m.M
	v75 = m.ExcPending
	if v75 != 0 {
		goto L5
	} else {
		goto L35
	}
L35:
	;
	v76 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v76 != 0 {
		goto L33
	} else {
		goto L36
	}
L36:
	;
	v77 = int32(-1)
	if v67 == int32(0) {
		v96 = v77
		goto L32
	} else {
		goto L37
	}
L37:
	;
	v98 = v77
	goto L31
L38:
	;
	v89 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v89
	v92 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(l0)+16)) = v92
	*(*int64)(unsafe.Add(mBase, uint32(l0)+4)) = v92
	if v67 != 0 {
		v98 = v89
		goto L31
	} else {
		goto L41
	}
L39:
	;
	v86 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v87 = m.T0[v86].(func(*base.Module, int32, int64, int32) int64)(m, l0, base.I64_extend_i32_s(v80-v81), int32(1))
	mBase = m.M
	v88 = m.ExcPending
	if v88 != 0 {
		goto L5
	} else {
		goto L40
	}
L40:
	;
	goto L38
L41:
	;
	v96 = v89
	goto L32
L42:
	;
	v98 = v96
	goto L31
}
func F_fileIsManifest(m *base.Module, l0 int32) int32 {
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
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v25 int32
	_ = v25
	var v29 int32
	_ = v29
	var v33 int32
	_ = v33
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v46 int64
	_ = v46
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v59 int32
	_ = v59
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v81 int32
	_ = v81
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v97 int32
	_ = v97
	var v101 int32
	_ = v101
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v110 int32
	_ = v110
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	v6 = m.G0
	v8 = v6 - int32(1184)
	m.G0 = v8
	v11 = F_fopen(m, l0, int32(_a_F_fileIsManifest_0))
	mBase = m.M
	if v11 == int32(0) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = l0
	v130 = F_iprintf(m, int32(_a_F_fileIsManifest_1), v8+int32(16))
	mBase = m.M
	v131 = m.ExcPending
	if v131 != 0 {
		goto L18
	} else {
		goto L42
	}
L2:
	;
	goto L39
L3:
	;
	v16 = *(*int32)(unsafe.Add(mBase, uint32(v11)+76))
	if int32(-1) < v16 {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	if int32(-1) < v33 {
		goto L12
	} else {
		goto L13
	}
L5:
	;
	if int32(-1) < v25 {
		v33 = v25
		goto L9
	} else {
		goto L10
	}
L6:
	;
	v20 = F___lockfile(m, v11)
	mBase = m.M
	v21 = *(*int32)(unsafe.Add(mBase, uint32(v11)+60))
	if v20 == int32(0) {
		v25 = v21
		goto L5
	} else {
		goto L8
	}
L7:
	;
	v19 = *(*int32)(unsafe.Add(mBase, uint32(v11)+60))
	v25 = v19
	goto L5
L8:
	;
	F___unlockfile(m, v11)
	mBase = m.M
	v25 = v21
	goto L5
L9:
	;
	goto L4
L10:
	;
	v29 = F___errno_location(m)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v29))) = int32(8)
	v33 = int32(-1)
	goto L9
L11:
	;
	if v43 == int32(-1) {
		goto L1
	} else {
		goto L14
	}
L12:
	;
	v42 = F___fstatat(m, v33, int32(_a_F_fileIsManifest_2), v8+int32(1088), int32(4096))
	mBase = m.M
	v43 = v42
	goto L11
L13:
	;
	v39 = F___syscall_ret(m, int32(-8))
	mBase = m.M
	v43 = v39
	goto L11
L14:
	;
	v46 = *(*int64)(unsafe.Add(mBase, uint32(v8)+1112))
	if v46 != int64(0) {
		goto L16
	} else {
		goto L17
	}
L15:
	;
	m.G0 = v8 + int32(1184)
	return v110
L16:
	;
	v59 = int32(0)
	goto L20
L17:
	;
	v49 = F_fclose(m, v11)
	mBase = m.M
	v52 = m.ExcPending
	if v52 != 0 {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	return int32(0)
L19:
	;
	v110 = int32(0)
	goto L15
L20:
	;
	goto L23
L21:
	;
	v105 = F_fclose(m, v11)
	mBase = m.M
	v106 = m.ExcPending
	if v106 != 0 {
		goto L18
	} else {
		goto L38
	}
L22:
	;
	goto L21
L23:
	;
	v68 = F_fgets(m, v8+int32(48), int32(1025), v11)
	mBase = m.M
	v69 = m.ExcPending
	if v69 != 0 {
		goto L18
	} else {
		goto L26
	}
L24:
	;
	v101 = *(*int32)(unsafe.Add(mBase, uint32(v8)+48))
	if v101 == int32(1701603686) {
		v59 = int32(1)
		goto L20
	} else {
		goto L37
	}
L25:
	;
	v97 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+48)))
	if v97 == int32(35) {
		goto L23
	} else {
		goto L36
	}
L26:
	;
	if v68 != 0 {
		goto L25
	} else {
		goto L27
	}
L27:
	;
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v11)+76))
	if int32(-1) < v72 {
		goto L30
	} else {
		goto L31
	}
L28:
	;
	if int32(base.Ui32(v81)>>(uint(int32(4))%32))&int32(1) != 0 {
		goto L22
	} else {
		goto L33
	}
L29:
	;
	goto L28
L30:
	;
	v76 = F___lockfile(m, v11)
	mBase = m.M
	v77 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
	if v76 == int32(0) {
		v81 = v77
		goto L29
	} else {
		goto L32
	}
L31:
	;
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
	v81 = v75
	goto L29
L32:
	;
	F___unlockfile(m, v11)
	mBase = m.M
	v81 = v77
	goto L29
L33:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8)+32)) = l0
	v91 = F_iprintf(m, int32(_a_F_fileIsManifest_3), v8+int32(32))
	mBase = m.M
	v92 = m.ExcPending
	if v92 != 0 {
		goto L18
	} else {
		goto L34
	}
L34:
	;
	v93 = F_fclose(m, v11)
	mBase = m.M
	v94 = m.ExcPending
	if v94 != 0 {
		goto L18
	} else {
		goto L35
	}
L35:
	;
	m.Env.Exit(m, int32(1))
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L36:
	;
	goto L24
L37:
	;
	goto L22
L38:
	;
	v110 = v59
	goto L15
L39:
	;
	v117 = *(*int32)(unsafe.Add(mBase, _c_F_fileIsManifest[0]))
	v118 = F___strerror_l(m, v117, v117)
	mBase = m.M
	goto L40
L40:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = v118
	*(*int32)(unsafe.Add(mBase, uint32(v8))) = l0
	v122 = F_iprintf(m, int32(_a_F_fileIsManifest_4), v8)
	mBase = m.M
	v123 = m.ExcPending
	if v123 != 0 {
		goto L18
	} else {
		goto L41
	}
L41:
	;
	m.Env.Exit(m, int32(1))
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L42:
	;
	v132 = F_fclose(m, v11)
	mBase = m.M
	v133 = m.ExcPending
	if v133 != 0 {
		goto L18
	} else {
		goto L43
	}
L43:
	;
	m.Env.Exit(m, int32(1))
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_fileno(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v20 int32
	_ = v20
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
	if int32(-1) < v4 {
		v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
		v12 = v9
	} else {
		v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
		v12 = v7
	}
	if int32(-1) < v12 {
		v20 = v12
	} else {
		*(*int32)(unsafe.Add(mBase, _c_F_fileno[0])) = int32(8)
		v20 = int32(-1)
	}
	return v20
}
func F_findBucketForInsert(m *base.Module, l0 int32, l1 int64, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v28 int32
	_ = v28
	var v34 int32
	_ = v34
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
	var v56 int32
	_ = v56
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v73 int32
	_ = v73
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
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v102 int32
	_ = v102
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
	var v115 int32
	_ = v115
	var v120 int32
	_ = v120
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v130 int32
	_ = v130
	var v139 int32
	_ = v139
	var v143 int32
	_ = v143
	var v150 int32
	_ = v150
	var v185 int32
	_ = v185
	var v195 int32
	_ = v195
	var v201 int32
	_ = v201
	var v207 int32
	_ = v207
	var v213 int32
	_ = v213
	var v219 int32
	_ = v219
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v15 = base.B2i32(v13 != int32(-1))
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l0+v15<<(uint(int32(2))%32))+8))
	if v19 == int32(0) {
		goto L5
	} else {
		goto L6
	}
L1:
	;
	F__serverAssert(m, int32(_a_F_findBucketForInsert_0), int32(_a_F_findBucketForInsert_1), int32(1083))
	mBase = m.M
	v219 = m.ExcPending
	if v219 != 0 {
		goto L18
	} else {
		goto L54
	}
L2:
	;
	F__serverAssert(m, int32(_a_F_findBucketForInsert_2), int32(_a_F_findBucketForInsert_1), int32(947))
	mBase = m.M
	v213 = m.ExcPending
	if v213 != 0 {
		goto L18
	} else {
		goto L53
	}
L3:
	;
	F__serverAssert(m, int32(_a_F_findBucketForInsert_3), int32(_a_F_findBucketForInsert_1), int32(946))
	mBase = m.M
	v207 = m.ExcPending
	if v207 != 0 {
		goto L18
	} else {
		goto L52
	}
L4:
	;
	F__serverAssert(m, int32(_a_F_findBucketForInsert_4), int32(_a_F_findBucketForInsert_1), int32(959))
	mBase = m.M
	v201 = m.ExcPending
	if v201 != 0 {
		goto L18
	} else {
		goto L51
	}
L5:
	;
	F__serverAssert(m, int32(_a_F_findBucketForInsert_5), int32(_a_F_findBucketForInsert_1), int32(1066))
	mBase = m.M
	v195 = m.ExcPending
	if v195 != 0 {
		goto L18
	} else {
		goto L50
	}
L6:
	;
	v22 = int32(0)
	v23 = int32(-1)
	v28 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0+v15)+24)))
	if v28 == int32(255) {
		goto L8
	} else {
		goto L9
	}
L7:
	;
	if v143&int32(1) == int32(0) {
		v185 = v22
		goto L26
	} else {
		goto L27
	}
L8:
	;
	v34 = v22
	goto L10
L9:
	;
	v34 = v23<<(uint(v28)%32) ^ v23
	goto L10
L10:
	;
	v39 = v19 + v34&base.I32_wrap_i64(l1)<<(uint(int32(6))%32)
	v40 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v39))))
	v41 = int32(1)
	v42 = v40 & v41
	v48 = int32(base.Ui32(v40)>>(uint(v41)%32)) & int32(4095)
	if v23<<(uint(int32(12)-v42)%32)^v48 != int32(-1) {
		v139 = v39
		v143 = v48
		goto L7
	} else {
		goto L11
	}
L11:
	;
	v56 = l0 + v15<<(uint(int32(2))%32) + int32(32)
	v62 = v39
	v64 = v40
	v65 = v42
	goto L12
L12:
	;
	if v65&int32(65535) == int32(0) {
		goto L15
	} else {
		goto L16
	}
L13:
	;
	v139 = v115
	v143 = v130
	goto L7
L14:
	;
	v120 = int32(-1)
	v122 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v115))))
	v123 = int32(1)
	v124 = v122 & v123
	v130 = int32(base.Ui32(v122)>>(uint(v123)%32)) & int32(4095)
	if v120<<(uint(int32(12)-v124)%32)^v130 == v120 {
		v62 = v115
		v64 = v122
		v65 = v124
		goto L12
	} else {
		goto L25
	}
L15:
	;
	if v64&int32(4096) == int32(0) {
		goto L4
	} else {
		goto L17
	}
L16:
	;
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v62)+60))
	v115 = v73
	goto L14
L17:
	;
	v79 = F_valkey_calloc(m, int32(64))
	mBase = m.M
	v82 = m.ExcPending
	if v82 != 0 {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	return int32(0)
L19:
	;
	v83 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v84 = *(*int32)(unsafe.Add(mBase, uint32(v83)+36))
	if v84 == int32(0) {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	v90 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v79))))
	if v90&int32(2) != 0 {
		goto L3
	} else {
		goto L23
	}
L21:
	;
	m.T0[v84].(func(*base.Module, int32, int32))(m, l0, int32(64))
	mBase = m.M
	v89 = m.ExcPending
	if v89 != 0 {
		goto L18
	} else {
		goto L22
	}
L22:
	;
	goto L20
L23:
	;
	v93 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v62))))
	if v93&int32(4096) == int32(0) {
		goto L2
	} else {
		goto L24
	}
L24:
	;
	v98 = *(*int32)(unsafe.Add(mBase, uint32(v62)+60))
	*(*int32)(unsafe.Add(mBase, uint32(v79)+16)) = v98
	v100 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v62)+13)))
	v102 = v90 | int32(2)
	*(*uint16)(unsafe.Add(mBase, uint32(v79))) = uint16(v102)
	*(*uint8)(unsafe.Add(mBase, uint32(v79)+2)) = uint8(v100)
	*(*int32)(unsafe.Add(mBase, uint32(v62)+60)) = v79
	v108 = int32(1)
	v109 = v93&int32(61438) | v108
	*(*uint16)(unsafe.Add(mBase, uint32(v62))) = uint16(v109)
	v111 = *(*int32)(unsafe.Add(mBase, uint32(v56)))
	*(*int32)(unsafe.Add(mBase, uint32(v56))) = v111 + v108
	v115 = v79
	goto L14
L25:
	;
	goto L13
L26:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v185
	if l3 == int32(0) {
		goto L48
	} else {
		goto L49
	}
L27:
	;
	v150 = int32(2)
	if v143&v150 != 0 {
		goto L28
	} else {
		goto L29
	}
L28:
	;
	if v143&int32(4) == int32(0) {
		v185 = v150
		goto L26
	} else {
		goto L30
	}
L29:
	;
	v185 = int32(1)
	goto L26
L30:
	;
	if v143&int32(8) != 0 {
		goto L31
	} else {
		goto L32
	}
L31:
	;
	if v143&int32(16) != 0 {
		goto L33
	} else {
		goto L34
	}
L32:
	;
	v185 = int32(3)
	goto L26
L33:
	;
	if v143&int32(32) != 0 {
		goto L35
	} else {
		goto L36
	}
L34:
	;
	v185 = int32(4)
	goto L26
L35:
	;
	if v143&int32(64) != 0 {
		goto L37
	} else {
		goto L38
	}
L36:
	;
	v185 = int32(5)
	goto L26
L37:
	;
	if v143&int32(128) != 0 {
		goto L39
	} else {
		goto L40
	}
L38:
	;
	v185 = int32(6)
	goto L26
L39:
	;
	if v143&int32(256) != 0 {
		goto L41
	} else {
		goto L42
	}
L40:
	;
	v185 = int32(7)
	goto L26
L41:
	;
	if v143&int32(512) != 0 {
		goto L43
	} else {
		goto L44
	}
L42:
	;
	v185 = int32(8)
	goto L26
L43:
	;
	if v143&int32(1024) != 0 {
		goto L45
	} else {
		goto L46
	}
L44:
	;
	v185 = int32(9)
	goto L26
L45:
	;
	if v143&int32(2048) != 0 {
		goto L1
	} else {
		goto L47
	}
L46:
	;
	v185 = int32(10)
	goto L26
L47:
	;
	v185 = int32(11)
	goto L26
L48:
	;
	return v139
L49:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v15
	goto L48
L50:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L51:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L52:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L53:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L54:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_fiprintf(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
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
	v10 = F_vfiprintf(m, l0, l1, l2)
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return int32(0)
	} else {
		m.G0 = v7 + int32(16)
		return v10
	}
}
func F_fireModuleSlotMigrationEvent(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v15 int64
	_ = v15
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v35 int32
	_ = v35
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v91 int64
	_ = v91
	var v97 int64
	_ = v97
	var v103 int64
	_ = v103
	var v109 int64
	_ = v109
	var v111 int64
	_ = v111
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v118 int32
	_ = v118
	v9 = m.G0
	v11 = v9 - int32(64)
	m.G0 = v11
	v14 = v11 + int32(56)
	v15 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v14))) = v15
	*(*int64)(unsafe.Add(mBase, uint32(v11)+48)) = v15
	*(*int64)(unsafe.Add(mBase, uint32(v11))) = int64(1)
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+164))
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v21)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v14))) = v22
	v26 = F_valkey_malloc(m, v22<<(uint(int32(3))%32))
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
	*(*int32)(unsafe.Add(mBase, uint32(v11)+52)) = v26
	if v22 == int32(0) {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v91 = *(*int64)(unsafe.Add(mBase, uint32(l0+int32(144))))
	*(*int64)(unsafe.Add(mBase, uint32(v11+int32(40)))) = v91
	v97 = *(*int64)(unsafe.Add(mBase, uint32(l0+int32(136))))
	*(*int64)(unsafe.Add(mBase, uint32(v11+int32(32)))) = v97
	v103 = *(*int64)(unsafe.Add(mBase, uint32(l0+int32(128))))
	*(*int64)(unsafe.Add(mBase, uint32(v11+int32(24)))) = v103
	v109 = *(*int64)(unsafe.Add(mBase, uint32(l0+int32(120))))
	*(*int64)(unsafe.Add(mBase, uint32(v11+int32(16)))) = v109
	v111 = *(*int64)(unsafe.Add(mBase, uint32(l0)+112))
	*(*int64)(unsafe.Add(mBase, uint32(v11)+8)) = v111
	F_moduleFireServerEvent(m, int64(19), l1, v11)
	mBase = m.M
	v115 = m.ExcPending
	if v115 != 0 {
		goto L1
	} else {
		goto L22
	}
L4:
	;
	v35 = int32(0)
	goto L5
L5:
	;
	v42 = v26 + v35<<(uint(int32(3))%32)
	v43 = *(*int32)(unsafe.Add(mBase, uint32(l0)+164))
	if v35 < int32(0) {
		goto L9
	} else {
		goto L10
	}
L6:
	;
	goto L3
L7:
	;
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v69)+8))
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v71)))
	*(*int32)(unsafe.Add(mBase, uint32(v42))) = v72
	v74 = *(*int32)(unsafe.Add(mBase, uint32(v71)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v42)+4)) = v74
	v77 = v35 + int32(1)
	if v77 != v22 {
		v35 = v77
		goto L5
	} else {
		goto L21
	}
L8:
	;
	goto L7
L9:
	;
	v62 = v43 + int32(4)
	v63 = v35 ^ int32(-1)
	goto L17
L10:
	;
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v43)))
	if v35 == int32(0) {
		v69 = v46
		goto L8
	} else {
		goto L11
	}
L11:
	;
	if v46 == int32(0) {
		v69 = v46
		goto L8
	} else {
		goto L12
	}
L12:
	;
	v51 = v46
	v52 = v35
	goto L13
L13:
	;
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v51)+4))
	v55 = v52 + int32(-1)
	if v55 == int32(0) {
		v69 = v53
		goto L8
	} else {
		goto L15
	}
L15:
	;
	if v53 != 0 {
		v51 = v53
		v52 = v55
		goto L13
	} else {
		goto L16
	}
L16:
	;
	v69 = v53
	goto L8
L17:
	;
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v62)))
	if v63 == int32(0) {
		v69 = v64
		goto L8
	} else {
		goto L19
	}
L18:
	;
	v69 = v64
	goto L8
L19:
	;
	if v64 != 0 {
		v62 = v64
		v63 = v63 + int32(-1)
		goto L17
	} else {
		goto L20
	}
L20:
	;
	goto L18
L21:
	;
	goto L6
L22:
	;
	v116 = *(*int32)(unsafe.Add(mBase, uint32(v11)+52))
	F_valkey_free(m, v116)
	mBase = m.M
	v118 = m.ExcPending
	if v118 != 0 {
		goto L1
	} else {
		goto L23
	}
L23:
	;
	m.G0 = v11 + int32(64)
	return
}
func F_flushPendingIOResponses(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v32 int32
	_ = v32
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v84 int32
	_ = v84
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	v2 = int32(0)
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v11 = *(*int32)(unsafe.Add(mBase, _c_F_flushPendingIOResponses[0]))
	if v11 == v2 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	m.G0 = v8 + int32(16)
	return
L2:
	;
	v15 = v8 + int32(8)
	v16 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+4)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v15))) = v16
	goto L3
L3:
	;
	v21 = v8 + int32(8)
	v23 = *(*int32)(unsafe.Add(mBase, uint32(v21)))
	if v23 == int32(0) {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	v92 = *(*int32)(unsafe.Add(mBase, _c_F_flushPendingIOResponses[0]))
	F_listRelease(m, v92)
	mBase = m.M
	v94 = m.ExcPending
	if v94 != 0 {
		goto L11
	} else {
		goto L28
	}
L5:
	;
	if v23 == int32(0) {
		goto L4
	} else {
		goto L8
	}
L6:
	;
	goto L5
L7:
	;
	v26 = *(*int32)(unsafe.Add(mBase, uint32(v21)+4))
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v23+base.B2i32(v26 == int32(0))<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v21))) = v32
	goto L6
L8:
	;
	v39 = v23
	goto L9
L9:
	;
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v39)+8))
	v44 = F_mpscEnqueue(m, int32(_a_F_flushPendingIOResponses_0), v42, int32(_a_F_flushPendingIOResponses_1))
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L11
	} else {
		goto L12
	}
L10:
	;
	goto L4
L11:
	;
	return
L12:
	;
	if l0 == int32(0) {
		goto L14
	} else {
		goto L15
	}
L13:
	;
	v69 = *(*int32)(unsafe.Add(mBase, _c_F_flushPendingIOResponses[0]))
	F_listDelNode(m, v69, v39)
	mBase = m.M
	v71 = m.ExcPending
	if v71 != 0 {
		goto L11
	} else {
		goto L23
	}
L14:
	;
	if v44 == int32(0) {
		goto L1
	} else {
		goto L22
	}
L15:
	;
	if v44 != 0 {
		goto L14
	} else {
		goto L16
	}
L16:
	;
	goto L17
L17:
	;
	v54 = *(*int32)(unsafe.Add(mBase, _c_F_flushPendingIOResponses[1]))
	if v54 != 0 {
		goto L1
	} else {
		goto L19
	}
L19:
	;
	v57 = F_mpscEnqueue(m, int32(_a_F_flushPendingIOResponses_0), v42, int32(_a_F_flushPendingIOResponses_1))
	mBase = m.M
	v58 = m.ExcPending
	if v58 != 0 {
		goto L11
	} else {
		goto L20
	}
L20:
	;
	if v57 == int32(0) {
		goto L17
	} else {
		goto L21
	}
L21:
	;
	goto L13
L22:
	;
	goto L13
L23:
	;
	v73 = v8 + int32(8)
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v73)))
	if v75 == int32(0) {
		goto L25
	} else {
		goto L26
	}
L24:
	;
	if v75 != 0 {
		v39 = v75
		goto L9
	} else {
		goto L27
	}
L25:
	;
	goto L24
L26:
	;
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v73)+4))
	v84 = *(*int32)(unsafe.Add(mBase, uint32(v75+base.B2i32(v78 == int32(0))<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v73))) = v84
	goto L25
L27:
	;
	goto L10
L28:
	;
	v95 = int32(0)
	*(*int32)(unsafe.Add(mBase, _c_F_flushPendingIOResponses[0])) = v95
	goto L1
}
func F_flushallCommand(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	v3 = m.G0
	v5 = v3 - int32(16)
	m.G0 = v5
	v9 = F_getFlushCommandFlags(m, l0, v5+int32(12))
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		return
	} else {
		if v9 == int32(-1) {
			m.G0 = v5 + int32(16)
			return
		} else {
			v13 = *(*int32)(unsafe.Add(mBase, uint32(v5)+12))
			F_flushAllDataAndResetRDB(m, v13|int32(2))
			mBase = m.M
			v17 = m.ExcPending
			if v17 != 0 {
				return
			} else {
				F_forceCommandPropagation(m, l0, int32(3))
				mBase = m.M
				v20 = m.ExcPending
				if v20 != 0 {
					return
				} else {
					v22 = *(*int32)(unsafe.Add(mBase, _c_F_flushallCommand[0]))
					F_addReply(m, l0, v22)
					mBase = m.M
					v24 = m.ExcPending
					if v24 != 0 {
						return
					} else {
						m.G0 = v5 + int32(16)
						return
					}
				}
			}
		}
	}
}
func F_fnmatch(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v34 int32
	_ = v34
	var v38 int32
	_ = v38
	var v51 int32
	_ = v51
	var v57 int32
	_ = v57
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v90 int32
	_ = v90
	var v94 int32
	_ = v94
	var v101 int32
	_ = v101
	var v114 int32
	_ = v114
	var v116 int32
	_ = v116
	var v118 int32
	_ = v118
	v10 = m.G0
	v12 = v10 - int32(16)
	m.G0 = v12
	v15 = l2 & int32(8)
	if l2&int32(1) == int32(0) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v12 + int32(16)
	return v118
L2:
	;
	if v15 == int32(0) {
		goto L21
	} else {
		goto L22
	}
L3:
	;
	v20 = l0
	v21 = l1
	goto L4
L4:
	;
	v34 = v21
	goto L7
L6:
	;
	v51 = v20
	goto L12
L7:
	;
	v38 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v34))))
	if v38 == int32(47) {
		goto L6
	} else {
		goto L9
	}
L9:
	;
	if v38 == int32(0) {
		goto L6
	} else {
		goto L10
	}
L10:
	;
	v34 = v34 + int32(1)
	goto L7
L11:
	;
	v64 = int32(*(*int8)(unsafe.Add(mBase, uint32(v34))))
	if v57 == v64 {
		goto L16
	} else {
		goto L17
	}
L12:
	;
	v57 = F_pat_next(m, v51, int32(-1), v12+int32(12), l2)
	mBase = m.M
	if v57 == int32(0) {
		goto L11
	} else {
		goto L14
	}
L14:
	;
	if v57 == int32(47) {
		goto L11
	} else {
		goto L15
	}
L15:
	;
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v12)+12))
	v51 = v51 + v62
	goto L12
L16:
	;
	v76 = F_fnmatch_internal(m, v20, v51-v20, v21, v34-v21, l2)
	mBase = m.M
	if v57 == int32(0) {
		v118 = v76
		goto L1
	} else {
		goto L19
	}
L17:
	;
	v66 = int32(0)
	if base.B2i32(v15 == v66)|base.B2i32(v64 == v66) == v66 {
		goto L16
	} else {
		goto L18
	}
L18:
	;
	v118 = int32(1)
	goto L1
L19:
	;
	if v76 != 0 {
		v118 = v76
		goto L1
	} else {
		goto L20
	}
L20:
	;
	v79 = *(*int32)(unsafe.Add(mBase, uint32(v12)+12))
	v20 = v51 + v79
	v21 = v34 + int32(1)
	goto L4
L21:
	;
	v114 = int32(-1)
	v116 = F_fnmatch_internal(m, l0, v114, l1, v114, l2)
	mBase = m.M
	v118 = v116
	goto L1
L22:
	;
	v90 = l1
	goto L23
L23:
	;
	v94 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v90))))
	if v94 == int32(47) {
		goto L26
	} else {
		goto L27
	}
L25:
	;
	v90 = v90 + int32(1)
	goto L23
L26:
	;
	v101 = F_fnmatch_internal(m, l0, int32(-1), l1, v90-l1, l2)
	mBase = m.M
	if v101 != 0 {
		goto L25
	} else {
		goto L29
	}
L27:
	;
	if v94 == int32(0) {
		goto L21
	} else {
		goto L28
	}
L28:
	;
	goto L25
L29:
	;
	v118 = int32(0)
	goto L1
}
func F_fnmatch_internal(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v95 int32
	_ = v95
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v109 int32
	_ = v109
	var v113 int32
	_ = v113
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v122 int32
	_ = v122
	var v127 int32
	_ = v127
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v143 int32
	_ = v143
	var v147 int32
	_ = v147
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v180 int32
	_ = v180
	var v182 int32
	_ = v182
	var v184 int32
	_ = v184
	var v186 int32
	_ = v186
	var v190 int32
	_ = v190
	var v200 int32
	_ = v200
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v217 int32
	_ = v217
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v242 int32
	_ = v242
	var v257 int32
	_ = v257
	var v259 int32
	_ = v259
	var v264 int32
	_ = v264
	var v269 int32
	_ = v269
	var v279 int32
	_ = v279
	var v283 int32
	_ = v283
	var v284 int32
	_ = v284
	var v286 int32
	_ = v286
	var v288 int32
	_ = v288
	var v291 int32
	_ = v291
	var v293 int32
	_ = v293
	var v295 int32
	_ = v295
	var v298 int32
	_ = v298
	var v303 int32
	_ = v303
	var v304 int32
	_ = v304
	var v307 int32
	_ = v307
	var v308 int32
	_ = v308
	var v311 int32
	_ = v311
	var v314 int32
	_ = v314
	var v315 int32
	_ = v315
	var v324 int32
	_ = v324
	var v325 int32
	_ = v325
	var v332 int32
	_ = v332
	var v335 int32
	_ = v335
	var v341 int32
	_ = v341
	var v345 int32
	_ = v345
	var v347 int32
	_ = v347
	var v348 int32
	_ = v348
	var v349 int32
	_ = v349
	var v352 int32
	_ = v352
	var v354 int32
	_ = v354
	var v362 int32
	_ = v362
	var v363 int32
	_ = v363
	var v365 int32
	_ = v365
	var v366 int32
	_ = v366
	var v368 int32
	_ = v368
	var v371 int32
	_ = v371
	var v373 int32
	_ = v373
	var v375 int32
	_ = v375
	var v378 int32
	_ = v378
	var v383 int32
	_ = v383
	var v384 int32
	_ = v384
	var v387 int32
	_ = v387
	var v388 int32
	_ = v388
	var v391 int32
	_ = v391
	var v394 int32
	_ = v394
	var v395 int32
	_ = v395
	var v403 int32
	_ = v403
	var v412 int32
	_ = v412
	var v425 int32
	_ = v425
	var v432 int32
	_ = v432
	var v440 int32
	_ = v440
	var v443 int32
	_ = v443
	var v444 int32
	_ = v444
	var v445 int32
	_ = v445
	var v449 int32
	_ = v449
	var v450 int32
	_ = v450
	var v457 int32
	_ = v457
	var v466 int32
	_ = v466
	var v467 int32
	_ = v467
	var v474 int32
	_ = v474
	var v482 int32
	_ = v482
	var v484 int32
	_ = v484
	var v487 int32
	_ = v487
	var v489 int32
	_ = v489
	var v491 int32
	_ = v491
	var v494 int32
	_ = v494
	var v499 int32
	_ = v499
	var v500 int32
	_ = v500
	var v503 int32
	_ = v503
	var v504 int32
	_ = v504
	var v507 int32
	_ = v507
	var v510 int32
	_ = v510
	var v511 int32
	_ = v511
	var v522 int32
	_ = v522
	var v524 int32
	_ = v524
	var v525 int32
	_ = v525
	var v526 int32
	_ = v526
	var v531 int32
	_ = v531
	var v533 int32
	_ = v533
	var v534 int32
	_ = v534
	var v536 int32
	_ = v536
	var v539 int32
	_ = v539
	var v541 int32
	_ = v541
	var v543 int32
	_ = v543
	var v546 int32
	_ = v546
	var v551 int32
	_ = v551
	var v552 int32
	_ = v552
	var v555 int32
	_ = v555
	var v556 int32
	_ = v556
	var v559 int32
	_ = v559
	var v562 int32
	_ = v562
	var v563 int32
	_ = v563
	var v573 int32
	_ = v573
	var v589 int32
	_ = v589
	var v590 int32
	_ = v590
	var v592 int32
	_ = v592
	var v595 int32
	_ = v595
	var v597 int32
	_ = v597
	var v599 int32
	_ = v599
	var v602 int32
	_ = v602
	var v607 int32
	_ = v607
	var v608 int32
	_ = v608
	var v611 int32
	_ = v611
	var v612 int32
	_ = v612
	var v615 int32
	_ = v615
	var v618 int32
	_ = v618
	var v619 int32
	_ = v619
	var v627 int32
	_ = v627
	var v631 int32
	_ = v631
	var v648 int32
	_ = v648
	var v649 int32
	_ = v649
	var v650 int32
	_ = v650
	var v653 int32
	_ = v653
	var v654 int32
	_ = v654
	var v658 int32
	_ = v658
	var v659 int32
	_ = v659
	var v660 int32
	_ = v660
	var v665 int32
	_ = v665
	var v672 int32
	_ = v672
	var v701 int32
	_ = v701
	var v729 int32
	_ = v729
	v18 = m.G0
	v20 = v18 - int32(16)
	m.G0 = v20
	if l4&int32(4) == int32(0) {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	m.G0 = v20 + int32(16)
	return v729
L2:
	;
	v729 = int32(1)
	goto L1
L3:
	;
	v33 = l4 & int32(16)
	v34 = l0
	v35 = l1
	v36 = l2
	v37 = l3
	goto L7
L4:
	;
	v26 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2))))
	if v26 != int32(46) {
		goto L3
	} else {
		goto L5
	}
L5:
	;
	v29 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v29 != int32(46) {
		goto L2
	} else {
		goto L6
	}
L6:
	;
	goto L3
L7:
	;
	v51 = int32(1)
	v54 = F_pat_next(m, v34, v35, v20+int32(12), l4)
	mBase = m.M
	switch v54 + int32(5) {
	case 0:
		goto L15
	default:
		goto L14
	case 3:
		v729 = v51
		goto L1
	}
L9:
	;
	v701 = *(*int32)(unsafe.Add(mBase, uint32(v20)+12))
	v34 = v34 + v701
	v35 = v35 - v701
	v36 = v36 + v109
	v37 = v37 - v109
	goto L7
L10:
	;
	if v412 != 0 {
		v729 = int32(1)
		goto L1
	} else {
		goto L112
	}
L11:
	;
	v403 = F_match_bracket(m, v34, v99, v117)
	mBase = m.M
	if v403 != 0 {
		goto L9
	} else {
		goto L111
	}
L12:
	;
	v180 = F_memchr(m, v36, int32(0), v37)
	mBase = m.M
	if v180 != 0 {
		goto L50
	} else {
		goto L51
	}
L13:
	;
	v122 = int32(1)
	v127 = v122
	v137 = int32(0)
	v138 = v122
	goto L43
L14:
	;
	v72 = v20 + int32(8)
	v75 = m.G0
	v77 = v75 - int32(16)
	m.G0 = v77
	if v37 != 0 {
		goto L24
	} else {
		goto L25
	}
L15:
	;
	v58 = v34 + int32(1)
	v60 = v35 + int32(-1)
	v63 = F_memchr(m, v58, int32(0), v60)
	mBase = m.M
	if v63 != 0 {
		goto L17
	} else {
		goto L18
	}
L16:
	;
	v66 = v58 + v65
	if int32(1) <= v65 {
		goto L13
	} else {
		goto L20
	}
L17:
	;
	v65 = v63 - v58
	goto L19
L18:
	;
	v65 = v60
	goto L19
L19:
	;
	goto L16
L20:
	;
	v172 = int32(0)
	v173 = int32(1)
	goto L12
L21:
	;
	v109 = *(*int32)(unsafe.Add(mBase, uint32(v20)+8))
	if v33 == int32(0) {
		v117 = v99
		goto L35
	} else {
		goto L36
	}
L22:
	;
	if int32(0) < v99 {
		goto L21
	} else {
		goto L34
	}
L23:
	;
	m.G0 = v77 + int32(16)
	goto L22
L24:
	;
	v82 = int32(*(*int8)(unsafe.Add(mBase, uint32(v36))))
	if int32(-1) < v82 {
		goto L26
	} else {
		goto L27
	}
L25:
	;
	v79 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v72))) = v79
	v99 = v79
	goto L23
L26:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v72))) = int32(1)
	v98 = int32(*(*int8)(unsafe.Add(mBase, uint32(v36))))
	v99 = v98
	goto L23
L27:
	;
	v87 = F_mbtowc(m, v77+int32(12), v36, v37)
	mBase = m.M
	v88 = *(*int32)(unsafe.Add(mBase, uint32(v77)+12))
	v91 = base.B2i32(v87 < int32(0))
	if v87 < int32(0) {
		goto L28
	} else {
		goto L29
	}
L28:
	;
	v92 = int32(1)
	goto L30
L29:
	;
	v92 = v87
	goto L30
L30:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v72))) = v92
	if v87 < int32(0) {
		goto L31
	} else {
		goto L32
	}
L31:
	;
	v95 = int32(-1)
	goto L33
L32:
	;
	v95 = v88
	goto L33
L33:
	;
	v99 = v95
	goto L23
L34:
	;
	v729 = base.B2i32(v54 != int32(0))
	goto L1
L35:
	;
	switch v54 + int32(4) {
	case 0:
		goto L9
	case 1:
		goto L11
	default:
		goto L40
	}
L36:
	;
	v113 = F_towupper(m, v99)
	mBase = m.M
	if v113 != v99 {
		v116 = v113
		goto L38
	} else {
		goto L39
	}
L37:
	;
	v117 = v116
	goto L35
L38:
	;
	goto L37
L39:
	;
	v115 = F_towlower(m, v99)
	mBase = m.M
	v116 = v115
	goto L38
L40:
	;
	if v99 == v54 {
		goto L9
	} else {
		goto L41
	}
L41:
	;
	if v117 != v54 {
		v729 = v51
		goto L1
	} else {
		goto L42
	}
L42:
	;
	goto L9
L43:
	;
	v143 = v34 + v127
	v147 = F_pat_next(m, v143, v66-v143, v20+int32(12), l4)
	mBase = m.M
	switch v147 + int32(5) {
	case 0:
		goto L47
	default:
		goto L46
	case 3:
		v729 = v122
		goto L1
	}
L44:
	;
	v172 = v155
	v173 = v156
	goto L12
L45:
	;
	v157 = *(*int32)(unsafe.Add(mBase, uint32(v20)+12))
	v158 = v127 + v157
	if base.Ui32(v34+v158) < base.Ui32(v66) {
		v127 = v158
		v137 = v155
		v138 = v156
		goto L43
	} else {
		goto L48
	}
L46:
	;
	v155 = v137 + int32(1)
	v156 = v138
	goto L45
L47:
	;
	v155 = int32(0)
	v156 = v127 + int32(1)
	goto L45
L48:
	;
	goto L44
L49:
	;
	if base.Ui32(v182) < base.Ui32(v172) {
		goto L2
	} else {
		goto L53
	}
L50:
	;
	v182 = v180 - v36
	goto L52
L51:
	;
	v182 = v37
	goto L52
L52:
	;
	goto L49
L53:
	;
	v184 = v36 + v182
	v186 = base.B2i32(v172 != int32(0))
	if v182 != 0 {
		goto L55
	} else {
		goto L56
	}
L54:
	;
	if v269 != 0 {
		v729 = int32(1)
		goto L1
	} else {
		goto L71
	}
L55:
	;
	if v172 == int32(0) {
		v264 = v184
		v269 = v186
		goto L54
	} else {
		goto L57
	}
L56:
	;
	v264 = v184
	v269 = v186
	goto L54
L57:
	;
	v190 = v184
	v200 = v172
	goto L58
L58:
	;
	v206 = int32(-1)
	v207 = v190 + v206
	v208 = int32(*(*int8)(unsafe.Add(mBase, uint32(v207))))
	if v206 < v208 {
		v242 = v207
		goto L60
	} else {
		goto L61
	}
L59:
	;
	v264 = v242
	v269 = v259
	goto L54
L60:
	;
	v257 = v200 + int32(-1)
	v259 = base.B2i32(v257 != int32(0))
	if base.Ui32(v242) <= base.Ui32(v36) {
		v264 = v242
		v269 = v259
		goto L54
	} else {
		goto L69
	}
L61:
	;
	goto L62
L62:
	;
	v212 = *(*int32)(unsafe.Add(mBase, _c_F_fnmatch_internal[0]))
	v213 = *(*int32)(unsafe.Add(mBase, uint32(v212)))
	if v213 == int32(0) {
		v242 = v207
		goto L60
	} else {
		goto L63
	}
L63:
	;
	v217 = v190
	goto L64
L64:
	;
	v234 = v217 + int32(-1)
	v235 = int32(*(*int8)(unsafe.Add(mBase, uint32(v234))))
	if v235 <= int32(-65) {
		goto L66
	} else {
		goto L67
	}
L65:
	;
	v242 = v234
	goto L60
L66:
	;
	if base.Ui32(v36) < base.Ui32(v234) {
		v217 = v234
		goto L64
	} else {
		goto L68
	}
L67:
	;
	v242 = v234
	goto L60
L68:
	;
	goto L65
L69:
	;
	if v257 != 0 {
		v190 = v242
		v200 = v257
		goto L58
	} else {
		goto L70
	}
L70:
	;
	goto L59
L71:
	;
	v279 = v34 + v173
	v283 = F_pat_next(m, v279, v66-v279, v20+int32(12), l4)
	mBase = m.M
	v284 = *(*int32)(unsafe.Add(mBase, uint32(v20)+12))
	v286 = v184 - v264
	v288 = v20 + int32(8)
	v291 = m.G0
	v293 = v291 - int32(16)
	m.G0 = v293
	if v286 != 0 {
		goto L74
	} else {
		goto L75
	}
L72:
	;
	if v315 < int32(1) {
		v412 = v283
		goto L10
	} else {
		goto L84
	}
L73:
	;
	m.G0 = v293 + int32(16)
	goto L72
L74:
	;
	v298 = int32(*(*int8)(unsafe.Add(mBase, uint32(v264))))
	if int32(-1) < v298 {
		goto L76
	} else {
		goto L77
	}
L75:
	;
	v295 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v288))) = v295
	v315 = v295
	goto L73
L76:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v288))) = int32(1)
	v314 = int32(*(*int8)(unsafe.Add(mBase, uint32(v264))))
	v315 = v314
	goto L73
L77:
	;
	v303 = F_mbtowc(m, v293+int32(12), v264, v286)
	mBase = m.M
	v304 = *(*int32)(unsafe.Add(mBase, uint32(v293)+12))
	v307 = base.B2i32(v303 < int32(0))
	if v303 < int32(0) {
		goto L78
	} else {
		goto L79
	}
L78:
	;
	v308 = int32(1)
	goto L80
L79:
	;
	v308 = v303
	goto L80
L80:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v288))) = v308
	if v303 < int32(0) {
		goto L81
	} else {
		goto L82
	}
L81:
	;
	v311 = int32(-1)
	goto L83
L82:
	;
	v311 = v304
	goto L83
L83:
	;
	v315 = v311
	goto L73
L84:
	;
	v324 = v315
	v325 = v279 + v284
	v332 = v283
	v335 = v264
	goto L85
L85:
	;
	v341 = *(*int32)(unsafe.Add(mBase, uint32(v20)+8))
	if v33 == int32(0) {
		v349 = v324
		goto L87
	} else {
		goto L88
	}
L87:
	;
	switch v332 + int32(4) {
	case 0:
		goto L92
	case 1:
		goto L94
	default:
		goto L93
	}
L88:
	;
	v345 = F_towupper(m, v324)
	mBase = m.M
	if v345 != v324 {
		v348 = v345
		goto L90
	} else {
		goto L91
	}
L89:
	;
	v349 = v348
	goto L87
L90:
	;
	goto L89
L91:
	;
	v347 = F_towlower(m, v324)
	mBase = m.M
	v348 = v347
	goto L90
L92:
	;
	v362 = F_pat_next(m, v325, v66-v325, v20+int32(12), l4)
	mBase = m.M
	v363 = *(*int32)(unsafe.Add(mBase, uint32(v20)+12))
	v365 = v335 + v341
	v366 = v184 - v365
	v368 = v20 + int32(8)
	v371 = m.G0
	v373 = v371 - int32(16)
	m.G0 = v373
	if v366 != 0 {
		goto L100
	} else {
		goto L101
	}
L93:
	;
	if v324 == v332 {
		goto L92
	} else {
		goto L96
	}
L94:
	;
	v352 = *(*int32)(unsafe.Add(mBase, uint32(v20)+12))
	v354 = F_match_bracket(m, v325-v352, v324, v349)
	mBase = m.M
	if v354 == int32(0) {
		goto L2
	} else {
		goto L95
	}
L95:
	;
	goto L92
L96:
	;
	if v349 != v332 {
		goto L2
	} else {
		goto L97
	}
L97:
	;
	goto L92
L98:
	;
	if int32(1) <= v395 {
		v324 = v395
		v325 = v325 + v363
		v332 = v362
		v335 = v365
		goto L85
	} else {
		goto L110
	}
L99:
	;
	m.G0 = v373 + int32(16)
	goto L98
L100:
	;
	v378 = int32(*(*int8)(unsafe.Add(mBase, uint32(v365))))
	if int32(-1) < v378 {
		goto L102
	} else {
		goto L103
	}
L101:
	;
	v375 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v368))) = v375
	v395 = v375
	goto L99
L102:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v368))) = int32(1)
	v394 = int32(*(*int8)(unsafe.Add(mBase, uint32(v365))))
	v395 = v394
	goto L99
L103:
	;
	v383 = F_mbtowc(m, v373+int32(12), v365, v366)
	mBase = m.M
	v384 = *(*int32)(unsafe.Add(mBase, uint32(v373)+12))
	v387 = base.B2i32(v383 < int32(0))
	if v383 < int32(0) {
		goto L104
	} else {
		goto L105
	}
L104:
	;
	v388 = int32(1)
	goto L106
L105:
	;
	v388 = v383
	goto L106
L106:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v368))) = v388
	if v383 < int32(0) {
		goto L107
	} else {
		goto L108
	}
L107:
	;
	v391 = int32(-1)
	goto L109
L108:
	;
	v391 = v384
	goto L109
L109:
	;
	v395 = v391
	goto L99
L110:
	;
	v412 = v362
	goto L10
L111:
	;
	v729 = v51
	goto L1
L112:
	;
	if v173 <= int32(1) {
		goto L113
	} else {
		goto L114
	}
L113:
	;
	v729 = int32(0)
	goto L1
L114:
	;
	v425 = v36
	v432 = v58
	goto L115
L115:
	;
	v440 = v279 - v432
	v443 = F_pat_next(m, v432, v440, v20+int32(12), l4)
	mBase = m.M
	v444 = *(*int32)(unsafe.Add(mBase, uint32(v20)+12))
	v445 = v432 + v444
	if v443 == int32(-5) {
		v665 = v425
		v672 = v445
		goto L117
	} else {
		goto L118
	}
L116:
	;
	goto L113
L117:
	;
	if base.Ui32(v672) < base.Ui32(v279) {
		v425 = v665
		v432 = v672
		goto L115
	} else {
		goto L180
	}
L118:
	;
	v449 = v443
	v450 = v425
	v457 = v445
	goto L119
L119:
	;
	v466 = v449
	v467 = v450
	v474 = v457
	goto L121
L120:
	;
	v665 = v654
	v672 = v660
	goto L117
L121:
	;
	v482 = v264 - v467
	v484 = v20 + int32(8)
	v487 = m.G0
	v489 = v487 - int32(16)
	m.G0 = v489
	if v482 != 0 {
		goto L125
	} else {
		goto L126
	}
L122:
	;
	goto L120
L123:
	;
	if v511 == int32(0) {
		goto L2
	} else {
		goto L135
	}
L124:
	;
	m.G0 = v489 + int32(16)
	goto L123
L125:
	;
	v494 = int32(*(*int8)(unsafe.Add(mBase, uint32(v467))))
	if int32(-1) < v494 {
		goto L127
	} else {
		goto L128
	}
L126:
	;
	v491 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v484))) = v491
	v511 = v491
	goto L124
L127:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v484))) = int32(1)
	v510 = int32(*(*int8)(unsafe.Add(mBase, uint32(v467))))
	v511 = v510
	goto L124
L128:
	;
	v499 = F_mbtowc(m, v489+int32(12), v467, v482)
	mBase = m.M
	v500 = *(*int32)(unsafe.Add(mBase, uint32(v489)+12))
	v503 = base.B2i32(v499 < int32(0))
	if v499 < int32(0) {
		goto L129
	} else {
		goto L130
	}
L129:
	;
	v504 = int32(1)
	goto L131
L130:
	;
	v504 = v499
	goto L131
L131:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v484))) = v504
	if v499 < int32(0) {
		goto L132
	} else {
		goto L133
	}
L132:
	;
	v507 = int32(-1)
	goto L134
L133:
	;
	v507 = v500
	goto L134
L134:
	;
	v511 = v507
	goto L124
L135:
	;
	if v33 == int32(0) {
		v526 = v511
		goto L136
	} else {
		goto L137
	}
L136:
	;
	switch v466 + int32(4) {
	case 0:
		goto L141
	case 1:
		goto L143
	default:
		goto L144
	}
L137:
	;
	v522 = F_towupper(m, v511)
	mBase = m.M
	if v522 != v511 {
		v525 = v522
		goto L139
	} else {
		goto L140
	}
L138:
	;
	v526 = v525
	goto L136
L139:
	;
	goto L138
L140:
	;
	v524 = F_towlower(m, v511)
	mBase = m.M
	v525 = v524
	goto L139
L141:
	;
	v653 = *(*int32)(unsafe.Add(mBase, uint32(v20)+8))
	v654 = v467 + v653
	v658 = F_pat_next(m, v474, v279-v474, v20+int32(12), l4)
	mBase = m.M
	v659 = *(*int32)(unsafe.Add(mBase, uint32(v20)+12))
	v660 = v474 + v659
	if v658 != int32(-5) {
		v466 = v658
		v467 = v654
		v474 = v660
		goto L121
	} else {
		goto L179
	}
L142:
	;
	v534 = v264 - v450
	v536 = v20 + int32(8)
	v539 = m.G0
	v541 = v539 - int32(16)
	m.G0 = v541
	if v534 != 0 {
		goto L152
	} else {
		goto L153
	}
L143:
	;
	v531 = *(*int32)(unsafe.Add(mBase, uint32(v20)+12))
	v533 = F_match_bracket(m, v474-v531, v511, v526)
	mBase = m.M
	if v533 != 0 {
		goto L141
	} else {
		goto L147
	}
L144:
	;
	if v511 == v466 {
		goto L141
	} else {
		goto L145
	}
L145:
	;
	if v526 != v466 {
		goto L142
	} else {
		goto L146
	}
L146:
	;
	goto L141
L147:
	;
	goto L142
L148:
	;
	v648 = F_pat_next(m, v432, v440, v20+int32(12), l4)
	mBase = m.M
	v649 = *(*int32)(unsafe.Add(mBase, uint32(v20)+12))
	v650 = v432 + v649
	if v648 != int32(-5) {
		v449 = v648
		v450 = v631
		v457 = v650
		goto L119
	} else {
		goto L178
	}
L149:
	;
	v627 = *(*int32)(unsafe.Add(mBase, uint32(v20)+8))
	v631 = v450 + v627
	goto L148
L150:
	;
	if int32(1) <= v563 {
		goto L149
	} else {
		goto L162
	}
L151:
	;
	m.G0 = v541 + int32(16)
	goto L150
L152:
	;
	v546 = int32(*(*int8)(unsafe.Add(mBase, uint32(v450))))
	if int32(-1) < v546 {
		goto L154
	} else {
		goto L155
	}
L153:
	;
	v543 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v536))) = v543
	v563 = v543
	goto L151
L154:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v536))) = int32(1)
	v562 = int32(*(*int8)(unsafe.Add(mBase, uint32(v450))))
	v563 = v562
	goto L151
L155:
	;
	v551 = F_mbtowc(m, v541+int32(12), v450, v534)
	mBase = m.M
	v552 = *(*int32)(unsafe.Add(mBase, uint32(v541)+12))
	v555 = base.B2i32(v551 < int32(0))
	if v551 < int32(0) {
		goto L156
	} else {
		goto L157
	}
L156:
	;
	v556 = int32(1)
	goto L158
L157:
	;
	v556 = v551
	goto L158
L158:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v536))) = v556
	if v551 < int32(0) {
		goto L159
	} else {
		goto L160
	}
L159:
	;
	v559 = int32(-1)
	goto L161
L160:
	;
	v559 = v552
	goto L161
L161:
	;
	v563 = v559
	goto L151
L162:
	;
	v573 = v450
	goto L163
L163:
	;
	v589 = v573 + int32(1)
	v590 = v264 - v589
	v592 = v20 + int32(8)
	v595 = m.G0
	v597 = v595 - int32(16)
	m.G0 = v597
	if v590 != 0 {
		goto L167
	} else {
		goto L168
	}
L165:
	;
	if v619 < int32(0) {
		v573 = v589
		goto L163
	} else {
		goto L177
	}
L166:
	;
	m.G0 = v597 + int32(16)
	goto L165
L167:
	;
	v602 = int32(*(*int8)(unsafe.Add(mBase, uint32(v589))))
	if int32(-1) < v602 {
		goto L169
	} else {
		goto L170
	}
L168:
	;
	v599 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v592))) = v599
	v619 = v599
	goto L166
L169:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v592))) = int32(1)
	v618 = int32(*(*int8)(unsafe.Add(mBase, uint32(v589))))
	v619 = v618
	goto L166
L170:
	;
	v607 = F_mbtowc(m, v597+int32(12), v589, v590)
	mBase = m.M
	v608 = *(*int32)(unsafe.Add(mBase, uint32(v597)+12))
	v611 = base.B2i32(v607 < int32(0))
	if v607 < int32(0) {
		goto L171
	} else {
		goto L172
	}
L171:
	;
	v612 = int32(1)
	goto L173
L172:
	;
	v612 = v607
	goto L173
L173:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v592))) = v612
	if v607 < int32(0) {
		goto L174
	} else {
		goto L175
	}
L174:
	;
	v615 = int32(-1)
	goto L176
L175:
	;
	v615 = v608
	goto L176
L176:
	;
	v619 = v615
	goto L166
L177:
	;
	v631 = v589
	goto L148
L178:
	;
	v665 = v631
	v672 = v650
	goto L117
L179:
	;
	goto L122
L180:
	;
	goto L116
}
func F_forbody(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) {
	mBase := m.M
	_ = mBase
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	var v52 int32
	_ = v52
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
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
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v89 int32
	_ = v89
	var v93 int32
	_ = v93
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
	var v123 int32
	_ = v123
	var v126 int32
	_ = v126
	var v134 int32
	_ = v134
	var v137 int32
	_ = v137
	var v143 int32
	_ = v143
	var v145 int32
	_ = v145
	var v155 int32
	_ = v155
	var v175 int32
	_ = v175
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v196 int32
	_ = v196
	var v203 int32
	_ = v203
	var v210 int32
	_ = v210
	var v216 int32
	_ = v216
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v243 int32
	_ = v243
	var v247 int32
	_ = v247
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
	var v259 int32
	_ = v259
	var v260 int32
	_ = v260
	var v263 int32
	_ = v263
	var v264 int32
	_ = v264
	var v265 int32
	_ = v265
	var v266 int32
	_ = v266
	var v269 int32
	_ = v269
	var v278 int32
	_ = v278
	var v282 int32
	_ = v282
	var v289 int32
	_ = v289
	var v290 int32
	_ = v290
	var v293 int32
	_ = v293
	var v299 int32
	_ = v299
	var v309 int32
	_ = v309
	var v330 int32
	_ = v330
	var v340 int32
	_ = v340
	var v342 int32
	_ = v342
	var v345 int32
	_ = v345
	var v346 int32
	_ = v346
	var v350 int32
	_ = v350
	var v352 int32
	_ = v352
	var v359 int32
	_ = v359
	var v365 int32
	_ = v365
	var v369 int32
	_ = v369
	var v383 int32
	_ = v383
	var v412 int32
	_ = v412
	var v416 int32
	_ = v416
	var v418 int32
	_ = v418
	var v419 int32
	_ = v419
	var v420 int32
	_ = v420
	var v422 int32
	_ = v422
	var v424 int32
	_ = v424
	var v425 int32
	_ = v425
	var v426 int32
	_ = v426
	var v428 int32
	_ = v428
	var v429 int32
	_ = v429
	var v430 int32
	_ = v430
	var v431 int32
	_ = v431
	var v434 int32
	_ = v434
	var v435 int32
	_ = v435
	var v436 int32
	_ = v436
	var v437 int32
	_ = v437
	var v440 int32
	_ = v440
	var v450 int32
	_ = v450
	var v454 int32
	_ = v454
	var v461 int32
	_ = v461
	var v462 int32
	_ = v462
	var v465 int32
	_ = v465
	var v471 int32
	_ = v471
	var v481 int32
	_ = v481
	var v502 int32
	_ = v502
	var v512 int32
	_ = v512
	var v514 int32
	_ = v514
	var v517 int32
	_ = v517
	var v518 int32
	_ = v518
	var v522 int32
	_ = v522
	var v524 int32
	_ = v524
	var v531 int32
	_ = v531
	var v537 int32
	_ = v537
	var v541 int32
	_ = v541
	var v555 int32
	_ = v555
	var v584 int32
	_ = v584
	var v588 int32
	_ = v588
	var v590 int32
	_ = v590
	var v591 int32
	_ = v591
	var v592 int32
	_ = v592
	var v594 int32
	_ = v594
	var v596 int32
	_ = v596
	var v598 int32
	_ = v598
	var v603 int32
	_ = v603
	var v604 int32
	_ = v604
	var v605 int32
	_ = v605
	var v606 int32
	_ = v606
	var v607 int32
	_ = v607
	var v616 int32
	_ = v616
	var v617 int32
	_ = v617
	var v618 int32
	_ = v618
	var v619 int32
	_ = v619
	var v620 int32
	_ = v620
	var v627 int32
	_ = v627
	var v628 int32
	_ = v628
	var v629 int32
	_ = v629
	var v633 int32
	_ = v633
	v19 = m.G0
	v21 = v19 - int32(32)
	m.G0 = v21
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v24 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23)+50)))
	v26 = v24 + int32(3)
	*(*uint8)(unsafe.Add(mBase, uint32(v23)+50)) = uint8(v26)
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v23)))
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v28)+24))
	v34 = v23 + v26&int32(255)<<(uint(int32(1))%32)
	v37 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v34+int32(166)))))
	v38 = int32(12)
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v23)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v29+v37*v38)+4)) = v41
	v45 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v34+int32(168)))))
	*(*int32)(unsafe.Add(mBase, uint32(v29+v45*v38)+4)) = v41
	v52 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v34+int32(170)))))
	*(*int32)(unsafe.Add(mBase, uint32(v29+v52*v38)+4)) = v41
	v57 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v57 == int32(259) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	F_luaX_next(m, l0)
	mBase = m.M
	v74 = m.ExcPending
	if v74 != 0 {
		goto L3
	} else {
		goto L7
	}
L2:
	;
	v60 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v62 = F_luaX_token2str(m, l0, int32(259))
	mBase = m.M
	v63 = m.ExcPending
	if v63 != 0 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	return
L4:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21))) = v62
	v65 = m.G3
	v68 = F_luaO_pushfstring(m, v60, v65+int32(_a_F_forbody_0), v21)
	mBase = m.M
	v69 = m.ExcPending
	if v69 != 0 {
		goto L3
	} else {
		goto L5
	}
L5:
	;
	F_luaX_syntaxerror(m, l0, v68)
	mBase = m.M
	v71 = m.ExcPending
	if v71 != 0 {
		goto L3
	} else {
		goto L6
	}
L6:
	;
	goto L1
L7:
	;
	if l4 == int32(0) {
		goto L9
	} else {
		goto L10
	}
L8:
	;
	v84 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v21)+18)) = uint8(v84)
	*(*int32)(unsafe.Add(mBase, uint32(v21)+12)) = int32(-1)
	v89 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23)+50)))
	*(*uint8)(unsafe.Add(mBase, uint32(v21)+17)) = uint8(v84)
	*(*uint8)(unsafe.Add(mBase, uint32(v21)+16)) = uint8(v89)
	v93 = *(*int32)(unsafe.Add(mBase, uint32(v23)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v21)+8)) = v93
	*(*int32)(unsafe.Add(mBase, uint32(v23)+20)) = v21 + int32(8)
	v98 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v99 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v98)+50)))
	v100 = v99 + l3
	*(*uint8)(unsafe.Add(mBase, uint32(v98)+50)) = uint8(v100)
	if l3 == v84 {
		goto L13
	} else {
		goto L14
	}
L9:
	;
	v81 = F_luaK_jump(m, v23)
	mBase = m.M
	v82 = m.ExcPending
	if v82 != 0 {
		goto L3
	} else {
		goto L12
	}
L10:
	;
	v79 = F_luaK_codeABx(m, v23, int32(32), l1, int32(131070))
	mBase = m.M
	v80 = m.ExcPending
	if v80 != 0 {
		goto L3
	} else {
		goto L11
	}
L11:
	;
	v83 = v79
	goto L8
L12:
	;
	v83 = v81
	goto L8
L13:
	;
	F_luaK_reserveregs(m, v23, l3)
	mBase = m.M
	v236 = m.ExcPending
	if v236 != 0 {
		goto L3
	} else {
		goto L24
	}
L14:
	;
	v105 = v100 & int32(255)
	v107 = v98 + int32(172)
	v108 = *(*int32)(unsafe.Add(mBase, uint32(v98)+24))
	v109 = *(*int32)(unsafe.Add(mBase, uint32(v98)))
	v110 = *(*int32)(unsafe.Add(mBase, uint32(v109)+24))
	v112 = l3 & int32(3)
	if v112 == int32(0) {
		v155 = l3
		goto L15
	} else {
		goto L16
	}
L15:
	;
	if base.Ui32(l3) < base.Ui32(int32(4)) {
		goto L13
	} else {
		goto L20
	}
L16:
	;
	v123 = l3
	v126 = v84
	goto L17
L17:
	;
	v134 = int32(1)
	v137 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v107+(v105-v123)<<(uint(v134)%32)))))
	*(*int32)(unsafe.Add(mBase, uint32(v110+v137*int32(12))+4)) = v108
	v143 = v123 + int32(-1)
	v145 = v126 + v134
	if v145 != v112 {
		v123 = v143
		v126 = v145
		goto L17
	} else {
		goto L19
	}
L18:
	;
	v155 = v143
	goto L15
L19:
	;
	goto L18
L20:
	;
	v175 = v155
	goto L21
L21:
	;
	v188 = v107 + (v105-v175)<<(uint(int32(1))%32)
	v189 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v188))))
	v190 = int32(12)
	*(*int32)(unsafe.Add(mBase, uint32(v110+v189*v190)+4)) = v108
	v196 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v188+int32(2)))))
	*(*int32)(unsafe.Add(mBase, uint32(v110+v196*v190)+4)) = v108
	v203 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v188+int32(4)))))
	*(*int32)(unsafe.Add(mBase, uint32(v110+v203*v190)+4)) = v108
	v210 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v188+int32(6)))))
	*(*int32)(unsafe.Add(mBase, uint32(v110+v210*v190)+4)) = v108
	v216 = v175 + int32(-4)
	if v216 != 0 {
		v175 = v216
		goto L21
	} else {
		goto L23
	}
L22:
	;
	goto L13
L23:
	;
	goto L22
L24:
	;
	v237 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v238 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v21)+30)) = uint8(v238)
	*(*int32)(unsafe.Add(mBase, uint32(v21)+24)) = int32(-1)
	v243 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v237)+50)))
	*(*uint8)(unsafe.Add(mBase, uint32(v21)+29)) = uint8(v238)
	*(*uint8)(unsafe.Add(mBase, uint32(v21)+28)) = uint8(v243)
	v247 = *(*int32)(unsafe.Add(mBase, uint32(v237)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v21)+20)) = v247
	*(*int32)(unsafe.Add(mBase, uint32(v237)+20)) = v21 + int32(20)
	F_chunk(m, l0)
	mBase = m.M
	v253 = m.ExcPending
	if v253 != 0 {
		goto L3
	} else {
		goto L25
	}
L25:
	;
	v254 = *(*int32)(unsafe.Add(mBase, uint32(v237)+20))
	v255 = *(*int32)(unsafe.Add(mBase, uint32(v254)))
	*(*int32)(unsafe.Add(mBase, uint32(v237)+20)) = v255
	v257 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v254)+8)))
	v258 = *(*int32)(unsafe.Add(mBase, uint32(v237)+12))
	v259 = *(*int32)(unsafe.Add(mBase, uint32(v258)+48))
	v260 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v259)+50)))
	if base.Ui32(v260) <= base.Ui32(v257) {
		goto L26
	} else {
		goto L27
	}
L26:
	;
	v412 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v254)+9)))
	if v412 == int32(0) {
		goto L39
	} else {
		goto L40
	}
L27:
	;
	v263 = v259 + int32(172)
	v264 = *(*int32)(unsafe.Add(mBase, uint32(v259)+24))
	v265 = *(*int32)(unsafe.Add(mBase, uint32(v259)))
	v266 = *(*int32)(unsafe.Add(mBase, uint32(v265)+24))
	v269 = (v260 - v257) & int32(3)
	if v269 != 0 {
		goto L29
	} else {
		goto L30
	}
L28:
	;
	if base.Ui32(int32(-4)) < base.Ui32(v257-v260) {
		v383 = v309
		goto L34
	} else {
		goto L35
	}
L29:
	;
	v278 = v260
	v282 = v238
	goto L31
L30:
	;
	v309 = v260
	goto L28
L31:
	;
	v289 = v278 + int32(-1)
	v290 = int32(1)
	v293 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v263+v289<<(uint(v290)%32)))))
	*(*int32)(unsafe.Add(mBase, uint32(v266+v293*int32(12))+8)) = v264
	v299 = v282 + v290
	if v299 != v269 {
		v278 = v289
		v282 = v299
		goto L31
	} else {
		goto L33
	}
L32:
	;
	v309 = v289
	goto L28
L33:
	;
	goto L32
L34:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v259)+50)) = uint8(v383)
	goto L26
L35:
	;
	v330 = v309
	goto L36
L36:
	;
	v340 = int32(1)
	v342 = v330<<(uint(v340)%32) + v263
	v345 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v342+int32(-2)))))
	v346 = int32(12)
	*(*int32)(unsafe.Add(mBase, uint32(v266+v345*v346)+8)) = v264
	v350 = int32(-4)
	v352 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v342+v350))))
	*(*int32)(unsafe.Add(mBase, uint32(v266+v352*v346)+8)) = v264
	v359 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v342+int32(-6)))))
	*(*int32)(unsafe.Add(mBase, uint32(v266+v359*v346)+8)) = v264
	v365 = v330 + v350
	v369 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v263+v365<<(uint(v340)%32)))))
	*(*int32)(unsafe.Add(mBase, uint32(v266+v369*v346)+8)) = v264
	if base.Ui32(v257) < base.Ui32(v365) {
		v330 = v365
		goto L36
	} else {
		goto L38
	}
L37:
	;
	v383 = v365
	goto L34
L38:
	;
	goto L37
L39:
	;
	v420 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v237)+50)))
	*(*int32)(unsafe.Add(mBase, uint32(v237)+36)) = v420
	v422 = *(*int32)(unsafe.Add(mBase, uint32(v254)+4))
	F_luaK_patchtohere(m, v237, v422)
	mBase = m.M
	v424 = m.ExcPending
	if v424 != 0 {
		goto L3
	} else {
		goto L42
	}
L40:
	;
	v416 = int32(0)
	v418 = F_luaK_codeABC(m, v237, int32(35), v257, v416, v416)
	mBase = m.M
	v419 = m.ExcPending
	if v419 != 0 {
		goto L3
	} else {
		goto L41
	}
L41:
	;
	goto L39
L42:
	;
	v425 = *(*int32)(unsafe.Add(mBase, uint32(v23)+20))
	v426 = *(*int32)(unsafe.Add(mBase, uint32(v425)))
	*(*int32)(unsafe.Add(mBase, uint32(v23)+20)) = v426
	v428 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v425)+8)))
	v429 = *(*int32)(unsafe.Add(mBase, uint32(v23)+12))
	v430 = *(*int32)(unsafe.Add(mBase, uint32(v429)+48))
	v431 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v430)+50)))
	if base.Ui32(v431) <= base.Ui32(v428) {
		goto L43
	} else {
		goto L44
	}
L43:
	;
	v584 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v425)+9)))
	if v584 == int32(0) {
		goto L56
	} else {
		goto L57
	}
L44:
	;
	v434 = v430 + int32(172)
	v435 = *(*int32)(unsafe.Add(mBase, uint32(v430)+24))
	v436 = *(*int32)(unsafe.Add(mBase, uint32(v430)))
	v437 = *(*int32)(unsafe.Add(mBase, uint32(v436)+24))
	v440 = (v431 - v428) & int32(3)
	if v440 != 0 {
		goto L46
	} else {
		goto L47
	}
L45:
	;
	if base.Ui32(int32(-4)) < base.Ui32(v428-v431) {
		v555 = v481
		goto L51
	} else {
		goto L52
	}
L46:
	;
	v450 = v431
	v454 = int32(0)
	goto L48
L47:
	;
	v481 = v431
	goto L45
L48:
	;
	v461 = v450 + int32(-1)
	v462 = int32(1)
	v465 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v434+v461<<(uint(v462)%32)))))
	*(*int32)(unsafe.Add(mBase, uint32(v437+v465*int32(12))+8)) = v435
	v471 = v454 + v462
	if v471 != v440 {
		v450 = v461
		v454 = v471
		goto L48
	} else {
		goto L50
	}
L49:
	;
	v481 = v461
	goto L45
L50:
	;
	goto L49
L51:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v430)+50)) = uint8(v555)
	goto L43
L52:
	;
	v502 = v481
	goto L53
L53:
	;
	v512 = int32(1)
	v514 = v502<<(uint(v512)%32) + v434
	v517 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v514+int32(-2)))))
	v518 = int32(12)
	*(*int32)(unsafe.Add(mBase, uint32(v437+v517*v518)+8)) = v435
	v522 = int32(-4)
	v524 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v514+v522))))
	*(*int32)(unsafe.Add(mBase, uint32(v437+v524*v518)+8)) = v435
	v531 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v514+int32(-6)))))
	*(*int32)(unsafe.Add(mBase, uint32(v437+v531*v518)+8)) = v435
	v537 = v502 + v522
	v541 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v434+v537<<(uint(v512)%32)))))
	*(*int32)(unsafe.Add(mBase, uint32(v437+v541*v518)+8)) = v435
	if base.Ui32(v428) < base.Ui32(v537) {
		v502 = v537
		goto L53
	} else {
		goto L55
	}
L54:
	;
	v555 = v537
	goto L51
L55:
	;
	goto L54
L56:
	;
	v592 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23)+50)))
	*(*int32)(unsafe.Add(mBase, uint32(v23)+36)) = v592
	v594 = *(*int32)(unsafe.Add(mBase, uint32(v425)+4))
	F_luaK_patchtohere(m, v23, v594)
	mBase = m.M
	v596 = m.ExcPending
	if v596 != 0 {
		goto L3
	} else {
		goto L59
	}
L57:
	;
	v588 = int32(0)
	v590 = F_luaK_codeABC(m, v23, int32(35), v428, v588, v588)
	mBase = m.M
	v591 = m.ExcPending
	if v591 != 0 {
		goto L3
	} else {
		goto L58
	}
L58:
	;
	goto L56
L59:
	;
	F_luaK_patchtohere(m, v23, v83)
	mBase = m.M
	v598 = m.ExcPending
	if v598 != 0 {
		goto L3
	} else {
		goto L60
	}
L60:
	;
	if l4 == int32(0) {
		goto L62
	} else {
		goto L63
	}
L61:
	;
	F_luaK_patchlist(m, v23, v629, v83+int32(1))
	mBase = m.M
	v633 = m.ExcPending
	if v633 != 0 {
		goto L3
	} else {
		goto L69
	}
L62:
	;
	v616 = F_luaK_codeABC(m, v23, int32(33), l1, int32(0), l3)
	mBase = m.M
	v617 = m.ExcPending
	if v617 != 0 {
		goto L3
	} else {
		goto L66
	}
L63:
	;
	v603 = F_luaK_codeABx(m, v23, int32(31), l1, int32(131070))
	mBase = m.M
	v604 = m.ExcPending
	if v604 != 0 {
		goto L3
	} else {
		goto L64
	}
L64:
	;
	v605 = *(*int32)(unsafe.Add(mBase, uint32(v23)))
	v606 = *(*int32)(unsafe.Add(mBase, uint32(v605)+20))
	v607 = *(*int32)(unsafe.Add(mBase, uint32(v23)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v606+v607<<(uint(int32(2))%32)+int32(-4)))) = l2
	goto L65
L65:
	;
	v629 = v603
	goto L61
L66:
	;
	v618 = *(*int32)(unsafe.Add(mBase, uint32(v23)))
	v619 = *(*int32)(unsafe.Add(mBase, uint32(v618)+20))
	v620 = *(*int32)(unsafe.Add(mBase, uint32(v23)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v619+v620<<(uint(int32(2))%32)+int32(-4)))) = l2
	goto L67
L67:
	;
	v627 = F_luaK_jump(m, v23)
	mBase = m.M
	v628 = m.ExcPending
	if v628 != 0 {
		goto L3
	} else {
		goto L68
	}
L68:
	;
	v629 = v627
	goto L61
L69:
	;
	m.G0 = v21 + int32(32)
	return
}
func F_forceCommandPropagation(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v13 int32
	_ = v13
	var v21 int32
	_ = v21
	var v29 int32
	_ = v29
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
	v4 = *(*int32)(unsafe.Add(mBase, uint32(v3)+56))
	if v4&int32(65537) == int32(0) {
		F__serverAssert(m, int32(_a_F_forceCommandPropagation_0), int32(_a_F_forceCommandPropagation_1), int32(3664))
		mBase = m.M
		v29 = m.ExcPending
		if v29 != 0 {
			return
		} else {
			F_abort(m)
			mBase = m.M
			base.Wasm_trap_unreachable()
			for {
			}
		}
	} else {
		if l1&int32(2) == int32(0) {
		} else {
			v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+200))
			*(*int32)(unsafe.Add(mBase, uint32(l0)+200)) = v13 | int32(32768)
		}
		if l1&int32(1) == int32(0) {
		} else {
			v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+200))
			*(*int32)(unsafe.Add(mBase, uint32(l0)+200)) = v21 | int32(16384)
		}
		return
	}
}
func F_foreach(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v83 int32
	_ = v83
	var v84 int64
	_ = v84
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v109 int32
	_ = v109
	var v112 int32
	_ = v112
	var v148 int32
	_ = v148
	var v149 int64
	_ = v149
	var v151 int32
	_ = v151
	var v153 int32
	_ = v153
	var v174 int32
	_ = v174
	var v177 int32
	_ = v177
	var v213 int32
	_ = v213
	var v214 int64
	_ = v214
	var v216 int32
	_ = v216
	var v218 int32
	_ = v218
	var v222 int32
	_ = v222
	var v226 int32
	_ = v226
	var v241 int32
	_ = v241
	var v244 int32
	_ = v244
	var v278 int32
	_ = v278
	var v281 int32
	_ = v281
	var v284 int32
	_ = v284
	var v303 int32
	_ = v303
	var v314 int32
	_ = v314
	var v315 int32
	_ = v315
	var v320 int32
	_ = v320
	F_luaL_checktype(m, l0, int32(1), int32(5))
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	F_luaL_checktype(m, l0, int32(2), int32(6))
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v14)+8)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v14 + int32(16)
	goto L4
L4:
	;
	v21 = F_lua_next(m, l0, int32(1))
	mBase = m.M
	v22 = m.ExcPending
	if v22 != 0 {
		goto L1
	} else {
		goto L7
	}
L5:
	;
	return v320
L6:
	;
	v320 = int32(0)
	goto L5
L7:
	;
	if v21 == int32(0) {
		goto L6
	} else {
		goto L8
	}
L8:
	;
	goto L9
L9:
	;
	goto L14
L10:
	;
	goto L6
L11:
	;
	goto L29
L12:
	;
	v83 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v84 = *(*int64)(unsafe.Add(mBase, uint32(v41)))
	*(*int64)(unsafe.Add(mBase, uint32(v83))) = v84
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v41)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v83)+8)) = v86
	v88 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v88 + int32(16)
	goto L11
L14:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v33 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v38 = v33 + int32(16)
	v39 = m.G398
	if base.Ui32(v38) < base.Ui32(v32) {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	v41 = v38
	goto L17
L16:
	;
	v41 = v39
	goto L17
L17:
	;
	goto L12
L27:
	;
	goto L45
L28:
	;
	v148 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v149 = *(*int64)(unsafe.Add(mBase, uint32(v112)))
	*(*int64)(unsafe.Add(mBase, uint32(v148))) = v149
	v151 = *(*int32)(unsafe.Add(mBase, uint32(v112)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v148)+8)) = v151
	v153 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v153 + int32(16)
	goto L27
L29:
	;
	goto L35
L35:
	;
	v109 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v112 = v109 + int32(-48)
	goto L28
L43:
	;
	v222 = int32(1)
	F_lua_call(m, l0, int32(2), v222)
	mBase = m.M
	v226 = m.ExcPending
	if v226 != 0 {
		goto L1
	} else {
		goto L59
	}
L44:
	;
	v213 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v214 = *(*int64)(unsafe.Add(mBase, uint32(v177)))
	*(*int64)(unsafe.Add(mBase, uint32(v213))) = v214
	v216 = *(*int32)(unsafe.Add(mBase, uint32(v177)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v213)+8)) = v216
	v218 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v218 + int32(16)
	goto L43
L45:
	;
	goto L51
L51:
	;
	v174 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v177 = v174 + int32(-48)
	goto L44
L59:
	;
	goto L62
L60:
	;
	if v284 != 0 {
		v320 = v222
		goto L5
	} else {
		goto L75
	}
L61:
	;
	v278 = m.G398
	if v244 != v278 {
		goto L73
	} else {
		goto L74
	}
L62:
	;
	goto L66
L66:
	;
	v241 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v244 = v241 + int32(-16)
	goto L61
L73:
	;
	v281 = *(*int32)(unsafe.Add(mBase, uint32(v244)+8))
	v284 = v281
	goto L60
L74:
	;
	v284 = int32(-1)
	goto L60
L75:
	;
	goto L78
L76:
	;
	v314 = F_lua_next(m, l0, int32(1))
	mBase = m.M
	v315 = m.ExcPending
	if v315 != 0 {
		goto L1
	} else {
		goto L84
	}
L77:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v303 + int32(-32)
	goto L76
L78:
	;
	v303 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	goto L77
L84:
	;
	if v314 != 0 {
		goto L9
	} else {
		goto L85
	}
L85:
	;
	goto L10
}
func F_foreachi(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v18 int32
	_ = v18
	var v22 int32
	_ = v22
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v81 int32
	_ = v81
	var v82 int64
	_ = v82
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v91 int32
	_ = v91
	var v96 int32
	_ = v96
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v114 int32
	_ = v114
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v159 int64
	_ = v159
	var v161 int32
	_ = v161
	var v163 int32
	_ = v163
	var v170 int32
	_ = v170
	var v185 int32
	_ = v185
	var v188 int32
	_ = v188
	var v222 int32
	_ = v222
	var v225 int32
	_ = v225
	var v228 int32
	_ = v228
	var v251 int32
	_ = v251
	v5 = int32(1)
	F_luaL_checktype(m, l0, v5, int32(5))
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
	v13 = F_lua_objlen(m, l0, int32(1))
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	F_luaL_checktype(m, l0, int32(2), int32(6))
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	if v13 < int32(1) {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	return int32(0)
L6:
	;
	v22 = v5
	goto L7
L7:
	;
	goto L12
L8:
	;
	goto L5
L9:
	;
	v91 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v91)+8)) = int32(3)
	*(*float64)(unsafe.Add(mBase, uint32(v91))) = base.F64_convert_i32_s(v22)
	v96 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v96 + int32(16)
	goto L25
L10:
	;
	v81 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v82 = *(*int64)(unsafe.Add(mBase, uint32(v39)))
	*(*int64)(unsafe.Add(mBase, uint32(v81))) = v82
	v84 = *(*int32)(unsafe.Add(mBase, uint32(v39)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v81)+8)) = v84
	v86 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v86 + int32(16)
	goto L9
L12:
	;
	v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v36 = v31 + int32(16)
	v37 = m.G398
	if base.Ui32(v36) < base.Ui32(v30) {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	v39 = v36
	goto L15
L14:
	;
	v39 = v37
	goto L15
L15:
	;
	goto L10
L25:
	;
	goto L29
L26:
	;
	F_lua_call(m, l0, int32(2), int32(1))
	mBase = m.M
	v170 = m.ExcPending
	if v170 != 0 {
		goto L1
	} else {
		goto L42
	}
L27:
	;
	v156 = *(*int32)(unsafe.Add(mBase, uint32(v114)))
	v157 = F_luaH_getnum(m, v156, v22)
	mBase = m.M
	v158 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v159 = *(*int64)(unsafe.Add(mBase, uint32(v157)))
	*(*int64)(unsafe.Add(mBase, uint32(v158))) = v159
	v161 = *(*int32)(unsafe.Add(mBase, uint32(v157)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v158)+8)) = v161
	v163 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v163 + int32(16)
	goto L26
L29:
	;
	v105 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v106 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v111 = v106 + int32(0)
	v112 = m.G398
	if base.Ui32(v111) < base.Ui32(v105) {
		goto L30
	} else {
		goto L31
	}
L30:
	;
	v114 = v111
	goto L32
L31:
	;
	v114 = v112
	goto L32
L32:
	;
	goto L27
L42:
	;
	goto L46
L43:
	;
	goto L62
L44:
	;
	if v228 == int32(0) {
		goto L43
	} else {
		goto L59
	}
L45:
	;
	v222 = m.G398
	if v188 != v222 {
		goto L57
	} else {
		goto L58
	}
L46:
	;
	goto L50
L50:
	;
	v185 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v188 = v185 + int32(-16)
	goto L45
L57:
	;
	v225 = *(*int32)(unsafe.Add(mBase, uint32(v188)+8))
	v228 = v225
	goto L44
L58:
	;
	v228 = int32(-1)
	goto L44
L59:
	;
	return int32(1)
L60:
	;
	if v22 != v13 {
		v22 = v22 + int32(1)
		goto L7
	} else {
		goto L68
	}
L61:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v251 + int32(-16)
	goto L60
L62:
	;
	v251 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	goto L61
L68:
	;
	goto L8
}
func F_fork(m *base.Module) int32 {
	mBase := m.M
	_ = mBase
	*(*int32)(unsafe.Add(mBase, _c_F_fork[0])) = int32(52)
	return int32(-1)
}
func F_fp_barrier_1(m *base.Module, l0 float64) float64 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	v3 = m.G0
	*(*float64)(unsafe.Add(mBase, uint32(v3-int32(16))+8)) = l0
	return l0
}
func F_fp_barrier_2(m *base.Module) float64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v4 int32
	_ = v4
	var v7 float64
	_ = v7
	v2 = m.G0
	v4 = v2 - int32(16)
	*(*int64)(unsafe.Add(mBase, uint32(v4)+8)) = int64(4503599627370496)
	v7 = *(*float64)(unsafe.Add(mBase, uint32(v4)+8))
	return v7
}
func F_fp_force_eval_1(m *base.Module, l0 float64) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	v2 = m.G0
	*(*float64)(unsafe.Add(mBase, uint32(v2-int32(16))+8)) = l0
	return
}
func F_fpconv_init(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v4 int32
	_ = v4
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v24 int32
	_ = v24
	var v29 int32
	_ = v29
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	v2 = m.G0
	v4 = v2 - int32(16)
	m.G0 = v4
	*(*int64)(unsafe.Add(mBase, uint32(v4))) = int64(4602678819172646912)
	v8 = int32(8)
	v11 = m.G3
	v14 = F_snprintf(m, v4+v8, v8, v11+int32(_a_F_fpconv_init_0), v4)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		return
	} else {
		v16 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4)+8)))
		if v16 != int32(48) {
			v29 = m.G3
			v34 = m.G397
			v35 = *(*int32)(unsafe.Add(mBase, uint32(v34)))
			v36 = F_fwrite(m, v29+int32(_a_F_fpconv_init_1), int32(45), int32(1), v35)
			mBase = m.M
			v37 = m.ExcPending
			if v37 != 0 {
				return
			} else {
				F_abort(m)
				mBase = m.M
				base.Wasm_trap_unreachable()
				for {
				}
			}
		} else {
			v19 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4)+10)))
			if v19&int32(255) != int32(53) {
				v29 = m.G3
				v34 = m.G397
				v35 = *(*int32)(unsafe.Add(mBase, uint32(v34)))
				v36 = F_fwrite(m, v29+int32(_a_F_fpconv_init_1), int32(45), int32(1), v35)
				mBase = m.M
				v37 = m.ExcPending
				if v37 != 0 {
					return
				} else {
					F_abort(m)
					mBase = m.M
					base.Wasm_trap_unreachable()
					for {
					}
				}
			} else {
				v24 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4)+11)))
				if v24&int32(255) == int32(0) {
					v39 = m.G3
					v42 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4)+9)))
					*(*uint8)(unsafe.Add(mBase, uint32(v39)+uint32(_c_F_fpconv_init[0]))) = uint8(v42)
					m.G0 = v4 + int32(16)
					return
				} else {
					v29 = m.G3
					v34 = m.G397
					v35 = *(*int32)(unsafe.Add(mBase, uint32(v34)))
					v36 = F_fwrite(m, v29+int32(_a_F_fpconv_init_1), int32(45), int32(1), v35)
					mBase = m.M
					v37 = m.ExcPending
					if v37 != 0 {
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
}
func F_freeErrorsRadixTreeAsync(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v7 int64
	_ = v7
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	v3 = m.G0
	v5 = v3 - int32(16)
	m.G0 = v5
	v7 = *(*int64)(unsafe.Add(mBase, uint32(l0)+16))
	if base.Ui64(v7) < base.Ui64(int64(65)) {
		F_raxFreeWithCallback(m, l0, int32(102))
		mBase = m.M
		v23 = m.ExcPending
		if v23 != 0 {
			return
		} else {
			m.G0 = v5 + int32(16)
			return
		}
	} else {
		v10 = int32(0)
		v12 = *(*int32)(unsafe.Add(mBase, _c_F_freeErrorsRadixTreeAsync[0]))
		v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		*(*int32)(unsafe.Add(mBase, _c_F_freeErrorsRadixTreeAsync[0])) = v12 + v13
		*(*int32)(unsafe.Add(mBase, uint32(v5))) = l0
		F_bioCreateLazyFreeJob(m, int32(554), int32(1), v5)
		mBase = m.M
		v20 = m.ExcPending
		if v20 != 0 {
			return
		} else {
			m.G0 = v5 + int32(16)
			return
		}
	}
}
func F_freeObjAsync(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v9 = F_lazyfreeGetFreeEffort(m, l0, l1, l2)
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		return
	} else {
		if base.Ui32(v9) < base.Ui32(int32(65)) {
			F_decrRefCount(m, l1)
			mBase = m.M
			v30 = m.ExcPending
			if v30 != 0 {
				return
			} else {
				m.G0 = v7 + int32(16)
				return
			}
		} else {
			v13 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
			if v13&int32(-8) != int32(8) {
				F_decrRefCount(m, l1)
				mBase = m.M
				v30 = m.ExcPending
				if v30 != 0 {
					return
				} else {
					m.G0 = v7 + int32(16)
					return
				}
			} else {
				v18 = int32(0)
				v20 = *(*int32)(unsafe.Add(mBase, _c_F_freeObjAsync[0]))
				v21 = int32(1)
				*(*int32)(unsafe.Add(mBase, _c_F_freeObjAsync[0])) = v20 + v21
				*(*int32)(unsafe.Add(mBase, uint32(v7))) = l1
				F_bioCreateLazyFreeJob(m, int32(551), v21, v7)
				mBase = m.M
				v28 = m.ExcPending
				if v28 != 0 {
					return
				} else {
					m.G0 = v7 + int32(16)
					return
				}
			}
		}
	}
}
func F_freePrefetchCommandsBatch(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v1 int32
	_ = v1
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
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
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	v1 = int32(0)
	v3 = *(*int32)(unsafe.Add(mBase, _c_F_freePrefetchCommandsBatch[0]))
	if v3 == v1 {
		return
	} else {
		v6 = *(*int32)(unsafe.Add(mBase, uint32(v3)+32))
		F_valkey_free(m, v6)
		mBase = m.M
		v8 = m.ExcPending
		if v8 != 0 {
			return
		} else {
			v10 = *(*int32)(unsafe.Add(mBase, _c_F_freePrefetchCommandsBatch[0]))
			v11 = *(*int32)(unsafe.Add(mBase, uint32(v10)+28))
			F_valkey_free(m, v11)
			mBase = m.M
			v13 = m.ExcPending
			if v13 != 0 {
				return
			} else {
				v15 = *(*int32)(unsafe.Add(mBase, _c_F_freePrefetchCommandsBatch[0]))
				v16 = *(*int32)(unsafe.Add(mBase, uint32(v15)+36))
				F_valkey_free(m, v16)
				mBase = m.M
				v18 = m.ExcPending
				if v18 != 0 {
					return
				} else {
					v20 = *(*int32)(unsafe.Add(mBase, _c_F_freePrefetchCommandsBatch[0]))
					v21 = *(*int32)(unsafe.Add(mBase, uint32(v20)+24))
					F_valkey_free(m, v21)
					mBase = m.M
					v23 = m.ExcPending
					if v23 != 0 {
						return
					} else {
						v25 = *(*int32)(unsafe.Add(mBase, _c_F_freePrefetchCommandsBatch[0]))
						v26 = *(*int32)(unsafe.Add(mBase, uint32(v25)+40))
						F_valkey_free(m, v26)
						mBase = m.M
						v28 = m.ExcPending
						if v28 != 0 {
							return
						} else {
							v30 = *(*int32)(unsafe.Add(mBase, _c_F_freePrefetchCommandsBatch[0]))
							F_valkey_free(m, v30)
							mBase = m.M
							v32 = m.ExcPending
							if v32 != 0 {
								return
							} else {
								v33 = int32(0)
								*(*int32)(unsafe.Add(mBase, _c_F_freePrefetchCommandsBatch[0])) = v33
								return
							}
						}
					}
				}
			}
		}
	}
}
func F_freeSharedQueryBuf(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	v2 = *(*int32)(unsafe.Add(mBase, _c_F_freeSharedQueryBuf[0]))
	F_sdsfree(m, v2)
	mBase = m.M
	v4 = m.ExcPending
	if v4 != 0 {
		return
	} else {
		v5 = int32(0)
		*(*int32)(unsafe.Add(mBase, _c_F_freeSharedQueryBuf[0])) = v5
		return
	}
}
func F_freeaddrinfo(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v60 int32
	_ = v60
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v71 int32
	_ = v71
	var __phi71 int32
	_ = __phi71
	var v72 int32
	_ = v72
	var __phi72 int32
	_ = __phi72
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v103 int32
	_ = v103
	var v112 int32
	_ = v112
	var v114 int32
	_ = v114
	var v117 int32
	_ = v117
	var v120 int32
	_ = v120
	var v122 int32
	_ = v122
	var v127 int32
	_ = v127
	var v134 int32
	_ = v134
	var v139 int32
	_ = v139
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v154 int32
	_ = v154
	var v162 int32
	_ = v162
	var v164 int32
	_ = v164
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v175 int32
	_ = v175
	var v177 int32
	_ = v177
	var v184 int32
	_ = v184
	var v186 int32
	_ = v186
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v204 int32
	_ = v204
	var v206 int32
	_ = v206
	var v208 int32
	_ = v208
	var v217 int32
	_ = v217
	var v219 int32
	_ = v219
	var v222 int32
	_ = v222
	var v227 int32
	_ = v227
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v238 int32
	_ = v238
	var __phi238 int32
	_ = __phi238
	var v239 int32
	_ = v239
	var __phi239 int32
	_ = __phi239
	var v245 int32
	_ = v245
	var v248 int32
	_ = v248
	var v262 int32
	_ = v262
	var v271 int32
	_ = v271
	var v273 int32
	_ = v273
	var v276 int32
	_ = v276
	var v279 int32
	_ = v279
	var v281 int32
	_ = v281
	var v286 int32
	_ = v286
	var v293 int32
	_ = v293
	var v298 int32
	_ = v298
	var v320 int32
	_ = v320
	var v332 int32
	_ = v332
	var v334 int32
	_ = v334
	var v336 int32
	_ = v336
	var v340 int32
	_ = v340
	var v345 int32
	_ = v345
	var v346 int32
	_ = v346
	var v357 int32
	_ = v357
	var v360 int32
	_ = v360
	var v367 int32
	_ = v367
	var v372 int32
	_ = v372
	var v376 int32
	_ = v376
	var v378 int32
	_ = v378
	var v393 int32
	_ = v393
	var v395 int32
	_ = v395
	var v398 int32
	_ = v398
	var v401 int32
	_ = v401
	var v405 int32
	_ = v405
	var v415 int32
	_ = v415
	var v416 int32
	_ = v416
	var v422 int32
	_ = v422
	var v424 int32
	_ = v424
	var v427 int32
	_ = v427
	var v431 int32
	_ = v431
	var v437 int32
	_ = v437
	var v439 int32
	_ = v439
	var v440 int32
	_ = v440
	var v441 int32
	_ = v441
	var v442 int32
	_ = v442
	var v451 int32
	_ = v451
	var v453 int32
	_ = v453
	var v454 int32
	_ = v454
	var v455 int32
	_ = v455
	var v457 int32
	_ = v457
	var v487 int32
	_ = v487
	var v488 int32
	_ = v488
	var v491 int32
	_ = v491
	var v493 int32
	_ = v493
	var v494 int32
	_ = v494
	var v501 int32
	_ = v501
	var v502 int32
	_ = v502
	var v504 int32
	_ = v504
	var v506 int32
	_ = v506
	var v508 int32
	_ = v508
	var v510 int32
	_ = v510
	var v513 int32
	_ = v513
	var v515 int32
	_ = v515
	var v517 int32
	_ = v517
	var v524 int32
	_ = v524
	var v526 int32
	_ = v526
	var v529 int32
	_ = v529
	var v534 int32
	_ = v534
	var v539 int32
	_ = v539
	var v540 int32
	_ = v540
	var v545 int32
	_ = v545
	var __phi545 int32
	_ = __phi545
	var v546 int32
	_ = v546
	var __phi546 int32
	_ = __phi546
	var v552 int32
	_ = v552
	var v555 int32
	_ = v555
	var v558 int32
	_ = v558
	var v559 int32
	_ = v559
	var v577 int32
	_ = v577
	var v586 int32
	_ = v586
	var v588 int32
	_ = v588
	var v591 int32
	_ = v591
	var v594 int32
	_ = v594
	var v596 int32
	_ = v596
	var v601 int32
	_ = v601
	var v608 int32
	_ = v608
	var v613 int32
	_ = v613
	var v618 int32
	_ = v618
	var v619 int32
	_ = v619
	var v628 int32
	_ = v628
	var v636 int32
	_ = v636
	var v638 int32
	_ = v638
	var v642 int32
	_ = v642
	var v643 int32
	_ = v643
	var v649 int32
	_ = v649
	var v651 int32
	_ = v651
	var v658 int32
	_ = v658
	var v660 int32
	_ = v660
	var v664 int32
	_ = v664
	var v665 int32
	_ = v665
	var v674 int32
	_ = v674
	var v675 int32
	_ = v675
	var v678 int32
	_ = v678
	var v680 int32
	_ = v680
	var v682 int32
	_ = v682
	var v691 int32
	_ = v691
	var v693 int32
	_ = v693
	var v696 int32
	_ = v696
	var v701 int32
	_ = v701
	var v706 int32
	_ = v706
	var v707 int32
	_ = v707
	var v712 int32
	_ = v712
	var __phi712 int32
	_ = __phi712
	var v713 int32
	_ = v713
	var __phi713 int32
	_ = __phi713
	var v719 int32
	_ = v719
	var v722 int32
	_ = v722
	var v736 int32
	_ = v736
	var v745 int32
	_ = v745
	var v747 int32
	_ = v747
	var v750 int32
	_ = v750
	var v753 int32
	_ = v753
	var v755 int32
	_ = v755
	var v760 int32
	_ = v760
	var v767 int32
	_ = v767
	var v772 int32
	_ = v772
	var v794 int32
	_ = v794
	var v806 int32
	_ = v806
	var v808 int32
	_ = v808
	var v810 int32
	_ = v810
	var v814 int32
	_ = v814
	var v819 int32
	_ = v819
	var v820 int32
	_ = v820
	var v831 int32
	_ = v831
	var v834 int32
	_ = v834
	var v841 int32
	_ = v841
	var v846 int32
	_ = v846
	var v850 int32
	_ = v850
	var v852 int32
	_ = v852
	var v867 int32
	_ = v867
	var v869 int32
	_ = v869
	var v872 int32
	_ = v872
	var v875 int32
	_ = v875
	var v879 int32
	_ = v879
	var v889 int32
	_ = v889
	var v890 int32
	_ = v890
	var v896 int32
	_ = v896
	var v898 int32
	_ = v898
	var v901 int32
	_ = v901
	var v905 int32
	_ = v905
	var v911 int32
	_ = v911
	var v913 int32
	_ = v913
	var v914 int32
	_ = v914
	var v915 int32
	_ = v915
	var v916 int32
	_ = v916
	var v925 int32
	_ = v925
	var v927 int32
	_ = v927
	var v928 int32
	_ = v928
	var v929 int32
	_ = v929
	var v931 int32
	_ = v931
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v2 == int32(0) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	if l0 == int32(0) {
		goto L102
	} else {
		goto L103
	}
L2:
	;
	goto L1
L3:
	;
	v13 = int32(-8)
	v14 = v2 + v13
	v17 = *(*int32)(unsafe.Add(mBase, uint32(v2+int32(-4))))
	v19 = v17 & v13
	v20 = v14 + v19
	if v17&int32(1) != 0 {
		v144 = v19
		v145 = v14
		goto L4
	} else {
		goto L5
	}
L4:
	;
	if base.Ui32(v20) <= base.Ui32(v145) {
		goto L2
	} else {
		goto L39
	}
L5:
	;
	if v17&int32(2) == int32(0) {
		goto L2
	} else {
		goto L6
	}
L6:
	;
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
	v28 = v14 - v27
	v30 = *(*int32)(unsafe.Add(mBase, _c_F_freeaddrinfo[0]))
	if base.Ui32(v28) < base.Ui32(v30) {
		goto L2
	} else {
		goto L7
	}
L7:
	;
	v32 = v27 + v19
	v34 = *(*int32)(unsafe.Add(mBase, _c_F_freeaddrinfo[1]))
	if v28 == v34 {
		goto L11
	} else {
		goto L12
	}
L8:
	;
	if v50 == int32(0) {
		v144 = v32
		v145 = v28
		goto L4
	} else {
		goto L27
	}
L9:
	;
	v103 = int32(0)
	goto L8
L10:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v39)+12)) = v36
	*(*int32)(unsafe.Add(mBase, uint32(v36)+8)) = v39
	v144 = v32
	v145 = v28
	goto L4
L11:
	;
	v84 = *(*int32)(unsafe.Add(mBase, uint32(v20)+4))
	v85 = int32(3)
	if v84&v85 != v85 {
		v144 = v32
		v145 = v28
		goto L4
	} else {
		goto L26
	}
L12:
	;
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v28)+12))
	if base.Ui32(int32(255)) < base.Ui32(v27) {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v28)+24))
	if v36 == v28 {
		goto L16
	} else {
		goto L17
	}
L14:
	;
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v28)+8))
	if v36 != v39 {
		goto L10
	} else {
		goto L15
	}
L15:
	;
	v41 = int32(0)
	v43 = *(*int32)(unsafe.Add(mBase, _c_F_freeaddrinfo[2]))
	*(*int32)(unsafe.Add(mBase, _c_F_freeaddrinfo[2])) = v43 & base.I32_rotl(int32(-2), int32(base.Ui32(v27)>>(uint(int32(3))%32)))
	v144 = v32
	v145 = v28
	goto L4
L16:
	;
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v28)+20))
	if v55 == int32(0) {
		goto L19
	} else {
		goto L20
	}
L17:
	;
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v28)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v52)+12)) = v36
	*(*int32)(unsafe.Add(mBase, uint32(v36)+8)) = v52
	v103 = v36
	goto L8
L18:
	;
	__phi71 = v65
	__phi72 = v66
	v71 = __phi71
	v72 = __phi72
	goto L22
L19:
	;
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v28)+16))
	if v60 == int32(0) {
		goto L9
	} else {
		goto L21
	}
L20:
	;
	v65 = v55
	v66 = v28 + int32(20)
	goto L18
L21:
	;
	v65 = v60
	v66 = v28 + int32(16)
	goto L18
L22:
	;
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v71)+20))
	if v78 != 0 {
		__phi71 = v78
		__phi72 = v71 + int32(20)
		v71 = __phi71
		v72 = __phi72
		goto L22
	} else {
		goto L24
	}
L23:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v72))) = int32(0)
	v103 = v71
	goto L8
L24:
	;
	v81 = *(*int32)(unsafe.Add(mBase, uint32(v71)+16))
	if v81 != 0 {
		__phi71 = v81
		__phi72 = v71 + int32(16)
		v71 = __phi71
		v72 = __phi72
		goto L22
	} else {
		goto L25
	}
L25:
	;
	goto L23
L26:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_freeaddrinfo[3])) = v32
	*(*int32)(unsafe.Add(mBase, uint32(v20)+4)) = v84 & int32(-2)
	*(*int32)(unsafe.Add(mBase, uint32(v28)+4)) = v32 | int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v20))) = v32
	goto L1
L27:
	;
	v112 = *(*int32)(unsafe.Add(mBase, uint32(v28)+28))
	v114 = v112 << (uint(int32(2)) % 32)
	v117 = *(*int32)(unsafe.Add(mBase, uint32(v114)+uint32(_c_F_freeaddrinfo[4])))
	if v28 != v117 {
		goto L29
	} else {
		goto L30
	}
L28:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v103)+24)) = v50
	v134 = *(*int32)(unsafe.Add(mBase, uint32(v28)+16))
	if v134 == int32(0) {
		goto L36
	} else {
		goto L37
	}
L29:
	;
	v127 = *(*int32)(unsafe.Add(mBase, uint32(v50)+16))
	if v127 != v28 {
		goto L33
	} else {
		goto L34
	}
L30:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v114)+uint32(_c_F_freeaddrinfo[4]))) = v103
	if v103 != 0 {
		goto L28
	} else {
		goto L31
	}
L31:
	;
	v120 = int32(0)
	v122 = *(*int32)(unsafe.Add(mBase, _c_F_freeaddrinfo[5]))
	*(*int32)(unsafe.Add(mBase, _c_F_freeaddrinfo[5])) = v122 & base.I32_rotl(int32(-2), v112)
	v144 = v32
	v145 = v28
	goto L4
L32:
	;
	if v103 == int32(0) {
		v144 = v32
		v145 = v28
		goto L4
	} else {
		goto L35
	}
L33:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v50)+20)) = v103
	goto L32
L34:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v50)+16)) = v103
	goto L32
L35:
	;
	goto L28
L36:
	;
	v139 = *(*int32)(unsafe.Add(mBase, uint32(v28)+20))
	if v139 == int32(0) {
		v144 = v32
		v145 = v28
		goto L4
	} else {
		goto L38
	}
L37:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v103)+16)) = v134
	*(*int32)(unsafe.Add(mBase, uint32(v134)+24)) = v103
	goto L36
L38:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v103)+20)) = v139
	*(*int32)(unsafe.Add(mBase, uint32(v139)+24)) = v103
	v144 = v32
	v145 = v28
	goto L4
L39:
	;
	v154 = *(*int32)(unsafe.Add(mBase, uint32(v20)+4))
	if v154&int32(1) == int32(0) {
		goto L2
	} else {
		goto L40
	}
L40:
	;
	if v154&int32(2) != 0 {
		goto L45
	} else {
		goto L46
	}
L41:
	;
	if base.Ui32(int32(255)) < base.Ui32(v320) {
		goto L79
	} else {
		goto L80
	}
L42:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v145)+4)) = v200 | int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v145+v200))) = v200
	if v145 != v184 {
		v320 = v200
		goto L41
	} else {
		goto L78
	}
L43:
	;
	if v217 == int32(0) {
		goto L42
	} else {
		goto L66
	}
L44:
	;
	v262 = int32(0)
	goto L43
L45:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+4)) = v154 & int32(-2)
	*(*int32)(unsafe.Add(mBase, uint32(v145)+4)) = v144 | int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v145+v144))) = v144
	v320 = v144
	goto L41
L46:
	;
	v162 = *(*int32)(unsafe.Add(mBase, _c_F_freeaddrinfo[6]))
	if v20 != v162 {
		goto L47
	} else {
		goto L48
	}
L47:
	;
	v184 = *(*int32)(unsafe.Add(mBase, _c_F_freeaddrinfo[1]))
	if v20 != v184 {
		goto L50
	} else {
		goto L51
	}
L48:
	;
	v164 = int32(0)
	*(*int32)(unsafe.Add(mBase, _c_F_freeaddrinfo[6])) = v145
	v168 = *(*int32)(unsafe.Add(mBase, _c_F_freeaddrinfo[7]))
	v169 = v168 + v144
	*(*int32)(unsafe.Add(mBase, _c_F_freeaddrinfo[7])) = v169
	*(*int32)(unsafe.Add(mBase, uint32(v145)+4)) = v169 | int32(1)
	v175 = *(*int32)(unsafe.Add(mBase, _c_F_freeaddrinfo[1]))
	if v145 != v175 {
		goto L2
	} else {
		goto L49
	}
L49:
	;
	v177 = int32(0)
	*(*int32)(unsafe.Add(mBase, _c_F_freeaddrinfo[3])) = v177
	*(*int32)(unsafe.Add(mBase, _c_F_freeaddrinfo[1])) = v177
	goto L1
L50:
	;
	v200 = v154&int32(-8) + v144
	v201 = *(*int32)(unsafe.Add(mBase, uint32(v20)+12))
	if base.Ui32(int32(255)) < base.Ui32(v154) {
		goto L52
	} else {
		goto L53
	}
L51:
	;
	v186 = int32(0)
	*(*int32)(unsafe.Add(mBase, _c_F_freeaddrinfo[1])) = v145
	v190 = *(*int32)(unsafe.Add(mBase, _c_F_freeaddrinfo[3]))
	v191 = v190 + v144
	*(*int32)(unsafe.Add(mBase, _c_F_freeaddrinfo[3])) = v191
	*(*int32)(unsafe.Add(mBase, uint32(v145)+4)) = v191 | int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v145+v191))) = v191
	goto L1
L52:
	;
	v217 = *(*int32)(unsafe.Add(mBase, uint32(v20)+24))
	if v201 == v20 {
		goto L56
	} else {
		goto L57
	}
L53:
	;
	v204 = *(*int32)(unsafe.Add(mBase, uint32(v20)+8))
	if v201 != v204 {
		goto L54
	} else {
		goto L55
	}
L54:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v204)+12)) = v201
	*(*int32)(unsafe.Add(mBase, uint32(v201)+8)) = v204
	goto L42
L55:
	;
	v206 = int32(0)
	v208 = *(*int32)(unsafe.Add(mBase, _c_F_freeaddrinfo[2]))
	*(*int32)(unsafe.Add(mBase, _c_F_freeaddrinfo[2])) = v208 & base.I32_rotl(int32(-2), int32(base.Ui32(v154)>>(uint(int32(3))%32)))
	goto L42
L56:
	;
	v222 = *(*int32)(unsafe.Add(mBase, uint32(v20)+20))
	if v222 == int32(0) {
		goto L59
	} else {
		goto L60
	}
L57:
	;
	v219 = *(*int32)(unsafe.Add(mBase, uint32(v20)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v219)+12)) = v201
	*(*int32)(unsafe.Add(mBase, uint32(v201)+8)) = v219
	v262 = v201
	goto L43
L58:
	;
	__phi238 = v232
	__phi239 = v233
	v238 = __phi238
	v239 = __phi239
	goto L62
L59:
	;
	v227 = *(*int32)(unsafe.Add(mBase, uint32(v20)+16))
	if v227 == int32(0) {
		goto L44
	} else {
		goto L61
	}
L60:
	;
	v232 = v222
	v233 = v20 + int32(20)
	goto L58
L61:
	;
	v232 = v227
	v233 = v20 + int32(16)
	goto L58
L62:
	;
	v245 = *(*int32)(unsafe.Add(mBase, uint32(v238)+20))
	if v245 != 0 {
		__phi238 = v245
		__phi239 = v238 + int32(20)
		v238 = __phi238
		v239 = __phi239
		goto L62
	} else {
		goto L64
	}
L63:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v239))) = int32(0)
	v262 = v238
	goto L43
L64:
	;
	v248 = *(*int32)(unsafe.Add(mBase, uint32(v238)+16))
	if v248 != 0 {
		__phi238 = v248
		__phi239 = v238 + int32(16)
		v238 = __phi238
		v239 = __phi239
		goto L62
	} else {
		goto L65
	}
L65:
	;
	goto L63
L66:
	;
	v271 = *(*int32)(unsafe.Add(mBase, uint32(v20)+28))
	v273 = v271 << (uint(int32(2)) % 32)
	v276 = *(*int32)(unsafe.Add(mBase, uint32(v273)+uint32(_c_F_freeaddrinfo[4])))
	if v20 != v276 {
		goto L68
	} else {
		goto L69
	}
L67:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v262)+24)) = v217
	v293 = *(*int32)(unsafe.Add(mBase, uint32(v20)+16))
	if v293 == int32(0) {
		goto L75
	} else {
		goto L76
	}
L68:
	;
	v286 = *(*int32)(unsafe.Add(mBase, uint32(v217)+16))
	if v286 != v20 {
		goto L72
	} else {
		goto L73
	}
L69:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v273)+uint32(_c_F_freeaddrinfo[4]))) = v262
	if v262 != 0 {
		goto L67
	} else {
		goto L70
	}
L70:
	;
	v279 = int32(0)
	v281 = *(*int32)(unsafe.Add(mBase, _c_F_freeaddrinfo[5]))
	*(*int32)(unsafe.Add(mBase, _c_F_freeaddrinfo[5])) = v281 & base.I32_rotl(int32(-2), v271)
	goto L42
L71:
	;
	if v262 == int32(0) {
		goto L42
	} else {
		goto L74
	}
L72:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v217)+20)) = v262
	goto L71
L73:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v217)+16)) = v262
	goto L71
L74:
	;
	goto L67
L75:
	;
	v298 = *(*int32)(unsafe.Add(mBase, uint32(v20)+20))
	if v298 == int32(0) {
		goto L42
	} else {
		goto L77
	}
L76:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v262)+16)) = v293
	*(*int32)(unsafe.Add(mBase, uint32(v293)+24)) = v262
	goto L75
L77:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v262)+20)) = v298
	*(*int32)(unsafe.Add(mBase, uint32(v298)+24)) = v262
	goto L42
L78:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_freeaddrinfo[3])) = v200
	goto L1
L79:
	;
	if base.Ui32(int32(16777215)) < base.Ui32(v320) {
		v367 = int32(31)
		goto L84
	} else {
		goto L85
	}
L80:
	;
	v332 = v320 & int32(-8)
	v334 = v332 + int32(9128464)
	v336 = *(*int32)(unsafe.Add(mBase, _c_F_freeaddrinfo[2]))
	v340 = int32(1) << (uint(int32(base.Ui32(v320)>>(uint(int32(3))%32))) % 32)
	if v336&v340 != 0 {
		goto L82
	} else {
		goto L83
	}
L81:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v332)+uint32(_c_F_freeaddrinfo[8]))) = v145
	*(*int32)(unsafe.Add(mBase, uint32(v346)+12)) = v145
	*(*int32)(unsafe.Add(mBase, uint32(v145)+12)) = v334
	*(*int32)(unsafe.Add(mBase, uint32(v145)+8)) = v346
	goto L1
L82:
	;
	v345 = *(*int32)(unsafe.Add(mBase, uint32(v332)+uint32(_c_F_freeaddrinfo[8])))
	v346 = v345
	goto L81
L83:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_freeaddrinfo[2])) = v336 | v340
	v346 = v334
	goto L81
L84:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v145)+28)) = v367
	*(*int64)(unsafe.Add(mBase, uint32(v145)+16)) = int64(0)
	v372 = v367 << (uint(int32(2)) % 32)
	v376 = *(*int32)(unsafe.Add(mBase, _c_F_freeaddrinfo[5]))
	v378 = int32(1) << (uint(v367) % 32)
	if v376&v378 != 0 {
		goto L89
	} else {
		goto L90
	}
L85:
	;
	v357 = base.I32_clz(int32(base.Ui32(v320) >> (uint(int32(8)) % 32)))
	v360 = int32(1)
	v367 = int32(base.Ui32(v320)>>(uint(int32(38)-v357)%32))&v360 - v357<<(uint(v360)%32) + int32(62)
	goto L84
L86:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v145+v439))) = v442
	*(*int32)(unsafe.Add(mBase, uint32(v145)+12)) = v441
	*(*int32)(unsafe.Add(mBase, uint32(v145+v437))) = v440
	v451 = int32(0)
	v453 = *(*int32)(unsafe.Add(mBase, _c_F_freeaddrinfo[9]))
	v454 = int32(-1)
	v455 = v453 + v454
	if v455 != 0 {
		goto L98
	} else {
		goto L99
	}
L87:
	;
	v431 = *(*int32)(unsafe.Add(mBase, uint32(v401)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v431)+12)) = v145
	*(*int32)(unsafe.Add(mBase, uint32(v401)+8)) = v145
	v437 = int32(24)
	v439 = int32(8)
	v440 = int32(0)
	v441 = v401
	v442 = v431
	goto L86
L88:
	;
	v437 = v422
	v439 = v424
	v440 = v145
	v441 = v145
	v442 = v427
	goto L86
L89:
	;
	if v367 == int32(31) {
		goto L91
	} else {
		goto L92
	}
L90:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_freeaddrinfo[5])) = v376 | v378
	*(*int32)(unsafe.Add(mBase, uint32(v372)+uint32(_c_F_freeaddrinfo[4]))) = v145
	v422 = int32(8)
	v424 = int32(24)
	v427 = v372 + int32(9128728)
	goto L88
L91:
	;
	v393 = int32(0)
	goto L93
L92:
	;
	v393 = int32(25) - int32(base.Ui32(v367)>>(uint(int32(1))%32))
	goto L93
L93:
	;
	v395 = *(*int32)(unsafe.Add(mBase, uint32(v372)+uint32(_c_F_freeaddrinfo[4])))
	v398 = v320 << (uint(v393) % 32)
	v401 = v395
	goto L94
L94:
	;
	v405 = *(*int32)(unsafe.Add(mBase, uint32(v401)+4))
	if v405&int32(-8) == v320 {
		goto L87
	} else {
		goto L96
	}
L95:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v415+int32(16)))) = v145
	v422 = int32(8)
	v424 = int32(24)
	v427 = v401
	goto L88
L96:
	;
	v415 = v401 + int32(base.Ui32(v398)>>(uint(int32(29))%32))&int32(4)
	v416 = *(*int32)(unsafe.Add(mBase, uint32(v415)+16))
	if v416 != 0 {
		v398 = v398 << (uint(int32(1)) % 32)
		v401 = v416
		goto L94
	} else {
		goto L97
	}
L97:
	;
	goto L95
L98:
	;
	v457 = v455
	goto L100
L99:
	;
	v457 = v454
	goto L100
L100:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_freeaddrinfo[9])) = v457
	goto L2
L101:
	;
	return
L102:
	;
	goto L101
L103:
	;
	v487 = int32(-8)
	v488 = l0 + v487
	v491 = *(*int32)(unsafe.Add(mBase, uint32(l0+int32(-4))))
	v493 = v491 & v487
	v494 = v488 + v493
	if v491&int32(1) != 0 {
		v618 = v493
		v619 = v488
		goto L104
	} else {
		goto L105
	}
L104:
	;
	if base.Ui32(v494) <= base.Ui32(v619) {
		goto L102
	} else {
		goto L139
	}
L105:
	;
	if v491&int32(2) == int32(0) {
		goto L102
	} else {
		goto L106
	}
L106:
	;
	v501 = *(*int32)(unsafe.Add(mBase, uint32(v488)))
	v502 = v488 - v501
	v504 = *(*int32)(unsafe.Add(mBase, _c_F_freeaddrinfo[0]))
	if base.Ui32(v502) < base.Ui32(v504) {
		goto L102
	} else {
		goto L107
	}
L107:
	;
	v506 = v501 + v493
	v508 = *(*int32)(unsafe.Add(mBase, _c_F_freeaddrinfo[1]))
	if v502 == v508 {
		goto L111
	} else {
		goto L112
	}
L108:
	;
	if v524 == int32(0) {
		v618 = v506
		v619 = v502
		goto L104
	} else {
		goto L127
	}
L109:
	;
	v577 = int32(0)
	goto L108
L110:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v513)+12)) = v510
	*(*int32)(unsafe.Add(mBase, uint32(v510)+8)) = v513
	v618 = v506
	v619 = v502
	goto L104
L111:
	;
	v558 = *(*int32)(unsafe.Add(mBase, uint32(v494)+4))
	v559 = int32(3)
	if v558&v559 != v559 {
		v618 = v506
		v619 = v502
		goto L104
	} else {
		goto L126
	}
L112:
	;
	v510 = *(*int32)(unsafe.Add(mBase, uint32(v502)+12))
	if base.Ui32(int32(255)) < base.Ui32(v501) {
		goto L113
	} else {
		goto L114
	}
L113:
	;
	v524 = *(*int32)(unsafe.Add(mBase, uint32(v502)+24))
	if v510 == v502 {
		goto L116
	} else {
		goto L117
	}
L114:
	;
	v513 = *(*int32)(unsafe.Add(mBase, uint32(v502)+8))
	if v510 != v513 {
		goto L110
	} else {
		goto L115
	}
L115:
	;
	v515 = int32(0)
	v517 = *(*int32)(unsafe.Add(mBase, _c_F_freeaddrinfo[2]))
	*(*int32)(unsafe.Add(mBase, _c_F_freeaddrinfo[2])) = v517 & base.I32_rotl(int32(-2), int32(base.Ui32(v501)>>(uint(int32(3))%32)))
	v618 = v506
	v619 = v502
	goto L104
L116:
	;
	v529 = *(*int32)(unsafe.Add(mBase, uint32(v502)+20))
	if v529 == int32(0) {
		goto L119
	} else {
		goto L120
	}
L117:
	;
	v526 = *(*int32)(unsafe.Add(mBase, uint32(v502)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v526)+12)) = v510
	*(*int32)(unsafe.Add(mBase, uint32(v510)+8)) = v526
	v577 = v510
	goto L108
L118:
	;
	__phi545 = v539
	__phi546 = v540
	v545 = __phi545
	v546 = __phi546
	goto L122
L119:
	;
	v534 = *(*int32)(unsafe.Add(mBase, uint32(v502)+16))
	if v534 == int32(0) {
		goto L109
	} else {
		goto L121
	}
L120:
	;
	v539 = v529
	v540 = v502 + int32(20)
	goto L118
L121:
	;
	v539 = v534
	v540 = v502 + int32(16)
	goto L118
L122:
	;
	v552 = *(*int32)(unsafe.Add(mBase, uint32(v545)+20))
	if v552 != 0 {
		__phi545 = v552
		__phi546 = v545 + int32(20)
		v545 = __phi545
		v546 = __phi546
		goto L122
	} else {
		goto L124
	}
L123:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v546))) = int32(0)
	v577 = v545
	goto L108
L124:
	;
	v555 = *(*int32)(unsafe.Add(mBase, uint32(v545)+16))
	if v555 != 0 {
		__phi545 = v555
		__phi546 = v545 + int32(16)
		v545 = __phi545
		v546 = __phi546
		goto L122
	} else {
		goto L125
	}
L125:
	;
	goto L123
L126:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_freeaddrinfo[3])) = v506
	*(*int32)(unsafe.Add(mBase, uint32(v494)+4)) = v558 & int32(-2)
	*(*int32)(unsafe.Add(mBase, uint32(v502)+4)) = v506 | int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v494))) = v506
	goto L101
L127:
	;
	v586 = *(*int32)(unsafe.Add(mBase, uint32(v502)+28))
	v588 = v586 << (uint(int32(2)) % 32)
	v591 = *(*int32)(unsafe.Add(mBase, uint32(v588)+uint32(_c_F_freeaddrinfo[4])))
	if v502 != v591 {
		goto L129
	} else {
		goto L130
	}
L128:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v577)+24)) = v524
	v608 = *(*int32)(unsafe.Add(mBase, uint32(v502)+16))
	if v608 == int32(0) {
		goto L136
	} else {
		goto L137
	}
L129:
	;
	v601 = *(*int32)(unsafe.Add(mBase, uint32(v524)+16))
	if v601 != v502 {
		goto L133
	} else {
		goto L134
	}
L130:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v588)+uint32(_c_F_freeaddrinfo[4]))) = v577
	if v577 != 0 {
		goto L128
	} else {
		goto L131
	}
L131:
	;
	v594 = int32(0)
	v596 = *(*int32)(unsafe.Add(mBase, _c_F_freeaddrinfo[5]))
	*(*int32)(unsafe.Add(mBase, _c_F_freeaddrinfo[5])) = v596 & base.I32_rotl(int32(-2), v586)
	v618 = v506
	v619 = v502
	goto L104
L132:
	;
	if v577 == int32(0) {
		v618 = v506
		v619 = v502
		goto L104
	} else {
		goto L135
	}
L133:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v524)+20)) = v577
	goto L132
L134:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v524)+16)) = v577
	goto L132
L135:
	;
	goto L128
L136:
	;
	v613 = *(*int32)(unsafe.Add(mBase, uint32(v502)+20))
	if v613 == int32(0) {
		v618 = v506
		v619 = v502
		goto L104
	} else {
		goto L138
	}
L137:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v577)+16)) = v608
	*(*int32)(unsafe.Add(mBase, uint32(v608)+24)) = v577
	goto L136
L138:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v577)+20)) = v613
	*(*int32)(unsafe.Add(mBase, uint32(v613)+24)) = v577
	v618 = v506
	v619 = v502
	goto L104
L139:
	;
	v628 = *(*int32)(unsafe.Add(mBase, uint32(v494)+4))
	if v628&int32(1) == int32(0) {
		goto L102
	} else {
		goto L140
	}
L140:
	;
	if v628&int32(2) != 0 {
		goto L145
	} else {
		goto L146
	}
L141:
	;
	if base.Ui32(int32(255)) < base.Ui32(v794) {
		goto L179
	} else {
		goto L180
	}
L142:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v619)+4)) = v674 | int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v619+v674))) = v674
	if v619 != v658 {
		v794 = v674
		goto L141
	} else {
		goto L178
	}
L143:
	;
	if v691 == int32(0) {
		goto L142
	} else {
		goto L166
	}
L144:
	;
	v736 = int32(0)
	goto L143
L145:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v494)+4)) = v628 & int32(-2)
	*(*int32)(unsafe.Add(mBase, uint32(v619)+4)) = v618 | int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v619+v618))) = v618
	v794 = v618
	goto L141
L146:
	;
	v636 = *(*int32)(unsafe.Add(mBase, _c_F_freeaddrinfo[6]))
	if v494 != v636 {
		goto L147
	} else {
		goto L148
	}
L147:
	;
	v658 = *(*int32)(unsafe.Add(mBase, _c_F_freeaddrinfo[1]))
	if v494 != v658 {
		goto L150
	} else {
		goto L151
	}
L148:
	;
	v638 = int32(0)
	*(*int32)(unsafe.Add(mBase, _c_F_freeaddrinfo[6])) = v619
	v642 = *(*int32)(unsafe.Add(mBase, _c_F_freeaddrinfo[7]))
	v643 = v642 + v618
	*(*int32)(unsafe.Add(mBase, _c_F_freeaddrinfo[7])) = v643
	*(*int32)(unsafe.Add(mBase, uint32(v619)+4)) = v643 | int32(1)
	v649 = *(*int32)(unsafe.Add(mBase, _c_F_freeaddrinfo[1]))
	if v619 != v649 {
		goto L102
	} else {
		goto L149
	}
L149:
	;
	v651 = int32(0)
	*(*int32)(unsafe.Add(mBase, _c_F_freeaddrinfo[3])) = v651
	*(*int32)(unsafe.Add(mBase, _c_F_freeaddrinfo[1])) = v651
	goto L101
L150:
	;
	v674 = v628&int32(-8) + v618
	v675 = *(*int32)(unsafe.Add(mBase, uint32(v494)+12))
	if base.Ui32(int32(255)) < base.Ui32(v628) {
		goto L152
	} else {
		goto L153
	}
L151:
	;
	v660 = int32(0)
	*(*int32)(unsafe.Add(mBase, _c_F_freeaddrinfo[1])) = v619
	v664 = *(*int32)(unsafe.Add(mBase, _c_F_freeaddrinfo[3]))
	v665 = v664 + v618
	*(*int32)(unsafe.Add(mBase, _c_F_freeaddrinfo[3])) = v665
	*(*int32)(unsafe.Add(mBase, uint32(v619)+4)) = v665 | int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v619+v665))) = v665
	goto L101
L152:
	;
	v691 = *(*int32)(unsafe.Add(mBase, uint32(v494)+24))
	if v675 == v494 {
		goto L156
	} else {
		goto L157
	}
L153:
	;
	v678 = *(*int32)(unsafe.Add(mBase, uint32(v494)+8))
	if v675 != v678 {
		goto L154
	} else {
		goto L155
	}
L154:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v678)+12)) = v675
	*(*int32)(unsafe.Add(mBase, uint32(v675)+8)) = v678
	goto L142
L155:
	;
	v680 = int32(0)
	v682 = *(*int32)(unsafe.Add(mBase, _c_F_freeaddrinfo[2]))
	*(*int32)(unsafe.Add(mBase, _c_F_freeaddrinfo[2])) = v682 & base.I32_rotl(int32(-2), int32(base.Ui32(v628)>>(uint(int32(3))%32)))
	goto L142
L156:
	;
	v696 = *(*int32)(unsafe.Add(mBase, uint32(v494)+20))
	if v696 == int32(0) {
		goto L159
	} else {
		goto L160
	}
L157:
	;
	v693 = *(*int32)(unsafe.Add(mBase, uint32(v494)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v693)+12)) = v675
	*(*int32)(unsafe.Add(mBase, uint32(v675)+8)) = v693
	v736 = v675
	goto L143
L158:
	;
	__phi712 = v706
	__phi713 = v707
	v712 = __phi712
	v713 = __phi713
	goto L162
L159:
	;
	v701 = *(*int32)(unsafe.Add(mBase, uint32(v494)+16))
	if v701 == int32(0) {
		goto L144
	} else {
		goto L161
	}
L160:
	;
	v706 = v696
	v707 = v494 + int32(20)
	goto L158
L161:
	;
	v706 = v701
	v707 = v494 + int32(16)
	goto L158
L162:
	;
	v719 = *(*int32)(unsafe.Add(mBase, uint32(v712)+20))
	if v719 != 0 {
		__phi712 = v719
		__phi713 = v712 + int32(20)
		v712 = __phi712
		v713 = __phi713
		goto L162
	} else {
		goto L164
	}
L163:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v713))) = int32(0)
	v736 = v712
	goto L143
L164:
	;
	v722 = *(*int32)(unsafe.Add(mBase, uint32(v712)+16))
	if v722 != 0 {
		__phi712 = v722
		__phi713 = v712 + int32(16)
		v712 = __phi712
		v713 = __phi713
		goto L162
	} else {
		goto L165
	}
L165:
	;
	goto L163
L166:
	;
	v745 = *(*int32)(unsafe.Add(mBase, uint32(v494)+28))
	v747 = v745 << (uint(int32(2)) % 32)
	v750 = *(*int32)(unsafe.Add(mBase, uint32(v747)+uint32(_c_F_freeaddrinfo[4])))
	if v494 != v750 {
		goto L168
	} else {
		goto L169
	}
L167:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v736)+24)) = v691
	v767 = *(*int32)(unsafe.Add(mBase, uint32(v494)+16))
	if v767 == int32(0) {
		goto L175
	} else {
		goto L176
	}
L168:
	;
	v760 = *(*int32)(unsafe.Add(mBase, uint32(v691)+16))
	if v760 != v494 {
		goto L172
	} else {
		goto L173
	}
L169:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v747)+uint32(_c_F_freeaddrinfo[4]))) = v736
	if v736 != 0 {
		goto L167
	} else {
		goto L170
	}
L170:
	;
	v753 = int32(0)
	v755 = *(*int32)(unsafe.Add(mBase, _c_F_freeaddrinfo[5]))
	*(*int32)(unsafe.Add(mBase, _c_F_freeaddrinfo[5])) = v755 & base.I32_rotl(int32(-2), v745)
	goto L142
L171:
	;
	if v736 == int32(0) {
		goto L142
	} else {
		goto L174
	}
L172:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v691)+20)) = v736
	goto L171
L173:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v691)+16)) = v736
	goto L171
L174:
	;
	goto L167
L175:
	;
	v772 = *(*int32)(unsafe.Add(mBase, uint32(v494)+20))
	if v772 == int32(0) {
		goto L142
	} else {
		goto L177
	}
L176:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v736)+16)) = v767
	*(*int32)(unsafe.Add(mBase, uint32(v767)+24)) = v736
	goto L175
L177:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v736)+20)) = v772
	*(*int32)(unsafe.Add(mBase, uint32(v772)+24)) = v736
	goto L142
L178:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_freeaddrinfo[3])) = v674
	goto L101
L179:
	;
	if base.Ui32(int32(16777215)) < base.Ui32(v794) {
		v841 = int32(31)
		goto L184
	} else {
		goto L185
	}
L180:
	;
	v806 = v794 & int32(-8)
	v808 = v806 + int32(9128464)
	v810 = *(*int32)(unsafe.Add(mBase, _c_F_freeaddrinfo[2]))
	v814 = int32(1) << (uint(int32(base.Ui32(v794)>>(uint(int32(3))%32))) % 32)
	if v810&v814 != 0 {
		goto L182
	} else {
		goto L183
	}
L181:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v806)+uint32(_c_F_freeaddrinfo[8]))) = v619
	*(*int32)(unsafe.Add(mBase, uint32(v820)+12)) = v619
	*(*int32)(unsafe.Add(mBase, uint32(v619)+12)) = v808
	*(*int32)(unsafe.Add(mBase, uint32(v619)+8)) = v820
	goto L101
L182:
	;
	v819 = *(*int32)(unsafe.Add(mBase, uint32(v806)+uint32(_c_F_freeaddrinfo[8])))
	v820 = v819
	goto L181
L183:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_freeaddrinfo[2])) = v810 | v814
	v820 = v808
	goto L181
L184:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v619)+28)) = v841
	*(*int64)(unsafe.Add(mBase, uint32(v619)+16)) = int64(0)
	v846 = v841 << (uint(int32(2)) % 32)
	v850 = *(*int32)(unsafe.Add(mBase, _c_F_freeaddrinfo[5]))
	v852 = int32(1) << (uint(v841) % 32)
	if v850&v852 != 0 {
		goto L189
	} else {
		goto L190
	}
L185:
	;
	v831 = base.I32_clz(int32(base.Ui32(v794) >> (uint(int32(8)) % 32)))
	v834 = int32(1)
	v841 = int32(base.Ui32(v794)>>(uint(int32(38)-v831)%32))&v834 - v831<<(uint(v834)%32) + int32(62)
	goto L184
L186:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v619+v913))) = v916
	*(*int32)(unsafe.Add(mBase, uint32(v619)+12)) = v915
	*(*int32)(unsafe.Add(mBase, uint32(v619+v911))) = v914
	v925 = int32(0)
	v927 = *(*int32)(unsafe.Add(mBase, _c_F_freeaddrinfo[9]))
	v928 = int32(-1)
	v929 = v927 + v928
	if v929 != 0 {
		goto L198
	} else {
		goto L199
	}
L187:
	;
	v905 = *(*int32)(unsafe.Add(mBase, uint32(v875)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v905)+12)) = v619
	*(*int32)(unsafe.Add(mBase, uint32(v875)+8)) = v619
	v911 = int32(24)
	v913 = int32(8)
	v914 = int32(0)
	v915 = v875
	v916 = v905
	goto L186
L188:
	;
	v911 = v896
	v913 = v898
	v914 = v619
	v915 = v619
	v916 = v901
	goto L186
L189:
	;
	if v841 == int32(31) {
		goto L191
	} else {
		goto L192
	}
L190:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_freeaddrinfo[5])) = v850 | v852
	*(*int32)(unsafe.Add(mBase, uint32(v846)+uint32(_c_F_freeaddrinfo[4]))) = v619
	v896 = int32(8)
	v898 = int32(24)
	v901 = v846 + int32(9128728)
	goto L188
L191:
	;
	v867 = int32(0)
	goto L193
L192:
	;
	v867 = int32(25) - int32(base.Ui32(v841)>>(uint(int32(1))%32))
	goto L193
L193:
	;
	v869 = *(*int32)(unsafe.Add(mBase, uint32(v846)+uint32(_c_F_freeaddrinfo[4])))
	v872 = v794 << (uint(v867) % 32)
	v875 = v869
	goto L194
L194:
	;
	v879 = *(*int32)(unsafe.Add(mBase, uint32(v875)+4))
	if v879&int32(-8) == v794 {
		goto L187
	} else {
		goto L196
	}
L195:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v889+int32(16)))) = v619
	v896 = int32(8)
	v898 = int32(24)
	v901 = v875
	goto L188
L196:
	;
	v889 = v875 + int32(base.Ui32(v872)>>(uint(int32(29))%32))&int32(4)
	v890 = *(*int32)(unsafe.Add(mBase, uint32(v889)+16))
	if v890 != 0 {
		v872 = v872 << (uint(int32(1)) % 32)
		v875 = v890
		goto L194
	} else {
		goto L197
	}
L197:
	;
	goto L195
L198:
	;
	v931 = v929
	goto L200
L199:
	;
	v931 = v928
	goto L200
L200:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_freeaddrinfo[9])) = v931
	goto L102
}
func F_freopen(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v40 int32
	_ = v40
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v85 int32
	_ = v85
	var v88 int32
	_ = v88
	var v91 int32
	_ = v91
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
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
	var v121 int32
	_ = v121
	v8 = m.G0
	v10 = v8 - int32(16)
	m.G0 = v10
	v15 = F_strchr(m, l1, int32(43))
	mBase = m.M
	if v15 != 0 {
		v19 = int32(2)
	} else {
		v16 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
		v19 = base.B2i32(v16 != int32(114))
	}
	v23 = F_strchr(m, l1, int32(120))
	mBase = m.M
	if v23 != 0 {
		v24 = v19 | int32(128)
	} else {
		v24 = v19
	}
	v28 = F_strchr(m, l1, int32(101))
	mBase = m.M
	if v28 != 0 {
		v29 = v24 | int32(524288)
	} else {
		v29 = v24
	}
	v32 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	if v32 == int32(114) {
		v35 = v29
	} else {
		v35 = v29 | int32(64)
	}
	if v32 == int32(119) {
		v40 = v35 | int32(512)
	} else {
		v40 = v35
	}
	if v32 == int32(97) {
		v45 = v40 | int32(1024)
	} else {
		v45 = v40
	}
	v46 = *(*int32)(unsafe.Add(mBase, uint32(l2)+76))
	if int32(0) <= v46 {
		v53 = int32(0)
	} else {
		v53 = int32(1)
	}
	v54 = F_fflush(m, l2)
	mBase = m.M
	v57 = m.ExcPending
	if v57 != 0 {
		return int32(0)
	} else {
		if l0 != 0 {
			v75 = F_fopen(m, l0, l1)
			mBase = m.M
			if v75 == int32(0) {
				v116 = F_fclose(m, l2)
				mBase = m.M
				v117 = m.ExcPending
				if v117 != 0 {
					return int32(0)
				} else {
					v121 = int32(0)
					m.G0 = v10 + int32(16)
					return v121
				}
			} else {
				v78 = *(*int32)(unsafe.Add(mBase, uint32(v75)+60))
				v79 = *(*int32)(unsafe.Add(mBase, uint32(l2)+60))
				if v78 != v79 {
					v85 = F___dup3(m, v78, v79, v45&int32(524288))
					mBase = m.M
					if v85 < int32(0) {
						v111 = F_fclose(m, v75)
						mBase = m.M
						v112 = m.ExcPending
						if v112 != 0 {
							return int32(0)
						} else {
							v116 = F_fclose(m, l2)
							mBase = m.M
							v117 = m.ExcPending
							if v117 != 0 {
								return int32(0)
							} else {
								v121 = int32(0)
								m.G0 = v10 + int32(16)
								return v121
							}
						}
					} else {
						v88 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
						v91 = *(*int32)(unsafe.Add(mBase, uint32(v75)))
						*(*int32)(unsafe.Add(mBase, uint32(l2))) = v88&int32(1) | v91
						v94 = *(*int32)(unsafe.Add(mBase, uint32(v75)+32))
						*(*int32)(unsafe.Add(mBase, uint32(l2)+32)) = v94
						v96 = *(*int32)(unsafe.Add(mBase, uint32(v75)+36))
						*(*int32)(unsafe.Add(mBase, uint32(l2)+36)) = v96
						v98 = *(*int32)(unsafe.Add(mBase, uint32(v75)+40))
						*(*int32)(unsafe.Add(mBase, uint32(l2)+40)) = v98
						v100 = *(*int32)(unsafe.Add(mBase, uint32(v75)+12))
						*(*int32)(unsafe.Add(mBase, uint32(l2)+12)) = v100
						v102 = F_fclose(m, v75)
						mBase = m.M
						v103 = m.ExcPending
						if v103 != 0 {
							return int32(0)
						} else {
							v107 = int32(0)
							*(*int32)(unsafe.Add(mBase, uint32(l2)+136)) = v107
							*(*int32)(unsafe.Add(mBase, uint32(l2)+72)) = v107
							if v53 != 0 {
								v121 = l2
							} else {
								v121 = l2
							}
							m.G0 = v10 + int32(16)
							return v121
						}
					}
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v75)+60)) = int32(-1)
					v88 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
					v91 = *(*int32)(unsafe.Add(mBase, uint32(v75)))
					*(*int32)(unsafe.Add(mBase, uint32(l2))) = v88&int32(1) | v91
					v94 = *(*int32)(unsafe.Add(mBase, uint32(v75)+32))
					*(*int32)(unsafe.Add(mBase, uint32(l2)+32)) = v94
					v96 = *(*int32)(unsafe.Add(mBase, uint32(v75)+36))
					*(*int32)(unsafe.Add(mBase, uint32(l2)+36)) = v96
					v98 = *(*int32)(unsafe.Add(mBase, uint32(v75)+40))
					*(*int32)(unsafe.Add(mBase, uint32(l2)+40)) = v98
					v100 = *(*int32)(unsafe.Add(mBase, uint32(v75)+12))
					*(*int32)(unsafe.Add(mBase, uint32(l2)+12)) = v100
					v102 = F_fclose(m, v75)
					mBase = m.M
					v103 = m.ExcPending
					if v103 != 0 {
						return int32(0)
					} else {
						v107 = int32(0)
						*(*int32)(unsafe.Add(mBase, uint32(l2)+136)) = v107
						*(*int32)(unsafe.Add(mBase, uint32(l2)+72)) = v107
						if v53 != 0 {
							v121 = l2
						} else {
							v121 = l2
						}
						m.G0 = v10 + int32(16)
						return v121
					}
				}
			}
		} else {
			v58 = *(*int32)(unsafe.Add(mBase, uint32(l2)+60))
			*(*int64)(unsafe.Add(mBase, uint32(v10))) = base.I64_extend_i32_s(v45 & int32(-524481))
			v64 = m.Env.X__syscall_fcntl64(m, v58, int32(4), v10)
			mBase = m.M
			if base.Ui32(v64) < base.Ui32(int32(-4095)) {
				v72 = v64
			} else {
				v67 = F___errno_location(m)
				mBase = m.M
				*(*int32)(unsafe.Add(mBase, uint32(v67))) = int32(0) - v64
				v72 = int32(-1)
			}
			if int32(0) <= v72 {
				v107 = int32(0)
				*(*int32)(unsafe.Add(mBase, uint32(l2)+136)) = v107
				*(*int32)(unsafe.Add(mBase, uint32(l2)+72)) = v107
				if v53 != 0 {
					v121 = l2
				} else {
					v121 = l2
				}
				m.G0 = v10 + int32(16)
				return v121
			} else {
				v116 = F_fclose(m, l2)
				mBase = m.M
				v117 = m.ExcPending
				if v117 != 0 {
					return int32(0)
				} else {
					v121 = int32(0)
					m.G0 = v10 + int32(16)
					return v121
				}
			}
		}
	}
}
func F_fseek(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	v5 = F___fseeko(m, l0, base.I64_extend_i32_s(l1), l2)
	v8 = m.ExcPending
	if v8 != 0 {
		return int32(0)
	} else {
		return v5
	}
}
func F_fwrite(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	v7 = l2 * l1
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l3)+76))
	if int32(-1) < v8 {
		v16 = F___fwritex(m, l0, v7, l3)
		mBase = m.M
		v17 = m.ExcPending
		if v17 != 0 {
			return int32(0)
		} else {
			v20 = v16
			if v20 != v7 {
				v26 = base.I32_div_u_s(v20, l1)
				return v26
			} else {
				if l1 != 0 {
					v24 = l2
				} else {
					v24 = int32(0)
				}
				return v24
			}
		}
	} else {
		v11 = F___fwritex(m, l0, v7, l3)
		mBase = m.M
		v14 = m.ExcPending
		if v14 != 0 {
			return int32(0)
		} else {
			v20 = v11
			if v20 != v7 {
				v26 = base.I32_div_u_s(v20, l1)
				return v26
			} else {
				if l1 != 0 {
					v24 = l2
				} else {
					v24 = int32(0)
				}
				return v24
			}
		}
	}
}
