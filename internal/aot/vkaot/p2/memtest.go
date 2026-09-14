package p2

import (
	base "github.com/shibukawa/vkmem/internal/aot/vkaot/base"
	"unsafe"
)

func F_memtest_addressing(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v19 int32
	_ = v19
	var v23 int64
	_ = v23
	var v25 int32
	_ = v25
	var v35 int64
	_ = v35
	var v36 int32
	_ = v36
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v49 int64
	_ = v49
	var v51 int64
	_ = v51
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v105 int64
	_ = v105
	var v111 int32
	_ = v111
	var v113 int32
	_ = v113
	var v121 int64
	_ = v121
	var v124 int32
	_ = v124
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v140 int32
	_ = v140
	var v142 int32
	_ = v142
	var v145 int64
	_ = v145
	var v147 int64
	_ = v147
	var v148 int32
	_ = v148
	var v151 int32
	_ = v151
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v165 int32
	_ = v165
	var v167 int32
	_ = v167
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v199 int64
	_ = v199
	var v214 int32
	_ = v214
	v12 = m.G0
	v14 = v12 - int32(16)
	m.G0 = v14
	if base.Ui32(l1) < base.Ui32(int32(4)) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v14 + int32(16)
	return v214
L2:
	;
	v214 = int32(0)
	goto L1
L3:
	;
	v19 = int32(base.Ui32(l1) >> (uint(int32(2)) % 32))
	v23 = base.I64_extend_i32_u(v19)
	v25 = *(*int32)(unsafe.Add(mBase, _c_F_memtest_addressing[0]))
	v35 = int64(0)
	v36 = l0
	goto L4
L4:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v36))) = v36
	if l2 == int32(0) {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	v111 = *(*int32)(unsafe.Add(mBase, _c_F_memtest_addressing[0]))
	v113 = l0
	v121 = int64(0)
	goto L18
L6:
	;
	v105 = v35 + int64(1)
	if v105 != v23 {
		v35 = v105
		v36 = v36 + int32(4)
		goto L4
	} else {
		goto L17
	}
L7:
	;
	if v35&int64(65535) != int64(0) {
		goto L6
	} else {
		goto L8
	}
L8:
	;
	v45 = int32(0)
	v47 = *(*int32)(unsafe.Add(mBase, _c_F_memtest_addressing[1]))
	v49 = int64(*(*uint32)(unsafe.Add(mBase, _c_F_memtest_addressing[2])))
	v51 = base.I64_div_u_s(v35*v49, base.I64_extend_i32_u(v19<<(uint(int32(1))%32)))
	v52 = base.I32_wrap_i64(v51)
	if v47 == v52 {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_memtest_addressing[1])) = v52
	v89 = F_fflush(m, v25)
	mBase = m.M
	v90 = m.ExcPending
	if v90 != 0 {
		goto L13
	} else {
		goto L16
	}
L10:
	;
	v55 = v45
	goto L11
L11:
	;
	v66 = F_putchar(m, int32(65))
	mBase = m.M
	v69 = m.ExcPending
	if v69 != 0 {
		goto L13
	} else {
		goto L14
	}
L12:
	;
	goto L9
L13:
	;
	return int32(0)
L14:
	;
	v71 = v55 + int32(1)
	v73 = *(*int32)(unsafe.Add(mBase, _c_F_memtest_addressing[1]))
	if base.Ui32(v71) < base.Ui32(v52-v73) {
		v55 = v71
		goto L11
	} else {
		goto L15
	}
L15:
	;
	goto L12
L16:
	;
	goto L6
L17:
	;
	goto L5
L18:
	;
	v124 = *(*int32)(unsafe.Add(mBase, uint32(v113)))
	if v124 == v113 {
		goto L20
	} else {
		goto L21
	}
L19:
	;
	goto L2
L20:
	;
	if l2 == int32(0) {
		goto L25
	} else {
		goto L26
	}
L21:
	;
	if l2 != 0 {
		goto L22
	} else {
		goto L23
	}
L22:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+4)) = v124
	*(*int32)(unsafe.Add(mBase, uint32(v14))) = v113
	v130 = F_iprintf(m, int32(_a_F_memtest_addressing_0), v14)
	mBase = m.M
	v131 = m.ExcPending
	if v131 != 0 {
		goto L13
	} else {
		goto L24
	}
L23:
	;
	v214 = int32(1)
	goto L1
L24:
	;
	m.Env.Exit(m, int32(1))
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L25:
	;
	v199 = v121 + int64(1)
	if v199 != v23 {
		v113 = v113 + int32(4)
		v121 = v199
		goto L18
	} else {
		goto L35
	}
L26:
	;
	if v121&int64(65535) != int64(0) {
		goto L25
	} else {
		goto L27
	}
L27:
	;
	v140 = int32(0)
	v142 = *(*int32)(unsafe.Add(mBase, _c_F_memtest_addressing[1]))
	v145 = int64(*(*uint32)(unsafe.Add(mBase, _c_F_memtest_addressing[2])))
	v147 = base.I64_div_u_s((v121+v23)*v145, base.I64_extend_i32_u(v19<<(uint(int32(1))%32)))
	v148 = base.I32_wrap_i64(v147)
	if v142 == v148 {
		goto L28
	} else {
		goto L29
	}
L28:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_memtest_addressing[1])) = v148
	v183 = F_fflush(m, v111)
	mBase = m.M
	v184 = m.ExcPending
	if v184 != 0 {
		goto L13
	} else {
		goto L34
	}
L29:
	;
	v151 = v140
	goto L30
L30:
	;
	v162 = F_putchar(m, int32(65))
	mBase = m.M
	v163 = m.ExcPending
	if v163 != 0 {
		goto L13
	} else {
		goto L32
	}
L31:
	;
	goto L28
L32:
	;
	v165 = v151 + int32(1)
	v167 = *(*int32)(unsafe.Add(mBase, _c_F_memtest_addressing[1]))
	if base.Ui32(v165) < base.Ui32(v148-v167) {
		v151 = v165
		goto L30
	} else {
		goto L33
	}
L33:
	;
	goto L31
L34:
	;
	goto L25
L35:
	;
	goto L19
}
func F_memtest_alloc_and_test(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v85 int32
	_ = v85
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v96 int32
	_ = v96
	var __phi96 int32
	_ = __phi96
	var v97 int32
	_ = v97
	var __phi97 int32
	_ = __phi97
	var v103 int32
	_ = v103
	var v106 int32
	_ = v106
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v128 int32
	_ = v128
	var v137 int32
	_ = v137
	var v139 int32
	_ = v139
	var v142 int32
	_ = v142
	var v145 int32
	_ = v145
	var v147 int32
	_ = v147
	var v152 int32
	_ = v152
	var v159 int32
	_ = v159
	var v164 int32
	_ = v164
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v179 int32
	_ = v179
	var v187 int32
	_ = v187
	var v189 int32
	_ = v189
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v200 int32
	_ = v200
	var v202 int32
	_ = v202
	var v209 int32
	_ = v209
	var v211 int32
	_ = v211
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v229 int32
	_ = v229
	var v231 int32
	_ = v231
	var v233 int32
	_ = v233
	var v242 int32
	_ = v242
	var v244 int32
	_ = v244
	var v247 int32
	_ = v247
	var v252 int32
	_ = v252
	var v257 int32
	_ = v257
	var v258 int32
	_ = v258
	var v263 int32
	_ = v263
	var __phi263 int32
	_ = __phi263
	var v264 int32
	_ = v264
	var __phi264 int32
	_ = __phi264
	var v270 int32
	_ = v270
	var v273 int32
	_ = v273
	var v287 int32
	_ = v287
	var v296 int32
	_ = v296
	var v298 int32
	_ = v298
	var v301 int32
	_ = v301
	var v304 int32
	_ = v304
	var v306 int32
	_ = v306
	var v311 int32
	_ = v311
	var v318 int32
	_ = v318
	var v323 int32
	_ = v323
	var v345 int32
	_ = v345
	var v357 int32
	_ = v357
	var v359 int32
	_ = v359
	var v361 int32
	_ = v361
	var v365 int32
	_ = v365
	var v370 int32
	_ = v370
	var v371 int32
	_ = v371
	var v382 int32
	_ = v382
	var v385 int32
	_ = v385
	var v392 int32
	_ = v392
	var v397 int32
	_ = v397
	var v401 int32
	_ = v401
	var v403 int32
	_ = v403
	var v418 int32
	_ = v418
	var v420 int32
	_ = v420
	var v423 int32
	_ = v423
	var v426 int32
	_ = v426
	var v430 int32
	_ = v430
	var v440 int32
	_ = v440
	var v441 int32
	_ = v441
	var v447 int32
	_ = v447
	var v449 int32
	_ = v449
	var v452 int32
	_ = v452
	var v456 int32
	_ = v456
	var v462 int32
	_ = v462
	var v464 int32
	_ = v464
	var v465 int32
	_ = v465
	var v466 int32
	_ = v466
	var v467 int32
	_ = v467
	var v476 int32
	_ = v476
	var v478 int32
	_ = v478
	var v479 int32
	_ = v479
	var v480 int32
	_ = v480
	var v482 int32
	_ = v482
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v11 = l0 << (uint(int32(20)) % 32)
	v12 = F_emscripten_builtin_malloc(m, v11)
	mBase = m.M
	if v12 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v26 = F_memtest_test(m, v12, v11, l1, int32(1))
	mBase = m.M
	v27 = m.ExcPending
	if v27 != 0 {
		goto L5
	} else {
		goto L7
	}
L2:
	;
	goto L3
L3:
	;
	v14 = *(*int32)(unsafe.Add(mBase, _c_F_memtest_alloc_and_test[0]))
	v15 = F___strerror_l(m, v14, v14)
	mBase = m.M
	goto L4
L4:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = v15
	*(*int32)(unsafe.Add(mBase, uint32(v8))) = l0
	v19 = *(*int32)(unsafe.Add(mBase, _c_F_memtest_alloc_and_test[1]))
	v21 = F_fiprintf(m, v19, int32(_a_F_memtest_alloc_and_test_0), v8)
	mBase = m.M
	v22 = m.ExcPending
	if v22 != 0 {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	return
L6:
	;
	m.Env.Exit(m, int32(1))
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L7:
	;
	if v12 == int32(0) {
		goto L9
	} else {
		goto L10
	}
L8:
	;
	m.G0 = v8 + int32(16)
	return
L9:
	;
	goto L8
L10:
	;
	v38 = int32(-8)
	v39 = v12 + v38
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v12+int32(-4))))
	v44 = v42 & v38
	v45 = v39 + v44
	if v42&int32(1) != 0 {
		v169 = v44
		v170 = v39
		goto L11
	} else {
		goto L12
	}
L11:
	;
	if base.Ui32(v45) <= base.Ui32(v170) {
		goto L9
	} else {
		goto L46
	}
L12:
	;
	if v42&int32(2) == int32(0) {
		goto L9
	} else {
		goto L13
	}
L13:
	;
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v39)))
	v53 = v39 - v52
	v55 = *(*int32)(unsafe.Add(mBase, _c_F_memtest_alloc_and_test[2]))
	if base.Ui32(v53) < base.Ui32(v55) {
		goto L9
	} else {
		goto L14
	}
L14:
	;
	v57 = v52 + v44
	v59 = *(*int32)(unsafe.Add(mBase, _c_F_memtest_alloc_and_test[3]))
	if v53 == v59 {
		goto L18
	} else {
		goto L19
	}
L15:
	;
	if v75 == int32(0) {
		v169 = v57
		v170 = v53
		goto L11
	} else {
		goto L34
	}
L16:
	;
	v128 = int32(0)
	goto L15
L17:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v64)+12)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v61)+8)) = v64
	v169 = v57
	v170 = v53
	goto L11
L18:
	;
	v109 = *(*int32)(unsafe.Add(mBase, uint32(v45)+4))
	v110 = int32(3)
	if v109&v110 != v110 {
		v169 = v57
		v170 = v53
		goto L11
	} else {
		goto L33
	}
L19:
	;
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v53)+12))
	if base.Ui32(int32(255)) < base.Ui32(v52) {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v53)+24))
	if v61 == v53 {
		goto L23
	} else {
		goto L24
	}
L21:
	;
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v53)+8))
	if v61 != v64 {
		goto L17
	} else {
		goto L22
	}
L22:
	;
	v66 = int32(0)
	v68 = *(*int32)(unsafe.Add(mBase, _c_F_memtest_alloc_and_test[4]))
	*(*int32)(unsafe.Add(mBase, _c_F_memtest_alloc_and_test[4])) = v68 & base.I32_rotl(int32(-2), int32(base.Ui32(v52)>>(uint(int32(3))%32)))
	v169 = v57
	v170 = v53
	goto L11
L23:
	;
	v80 = *(*int32)(unsafe.Add(mBase, uint32(v53)+20))
	if v80 == int32(0) {
		goto L26
	} else {
		goto L27
	}
L24:
	;
	v77 = *(*int32)(unsafe.Add(mBase, uint32(v53)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v77)+12)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v61)+8)) = v77
	v128 = v61
	goto L15
L25:
	;
	__phi96 = v90
	__phi97 = v91
	v96 = __phi96
	v97 = __phi97
	goto L29
L26:
	;
	v85 = *(*int32)(unsafe.Add(mBase, uint32(v53)+16))
	if v85 == int32(0) {
		goto L16
	} else {
		goto L28
	}
L27:
	;
	v90 = v80
	v91 = v53 + int32(20)
	goto L25
L28:
	;
	v90 = v85
	v91 = v53 + int32(16)
	goto L25
L29:
	;
	v103 = *(*int32)(unsafe.Add(mBase, uint32(v96)+20))
	if v103 != 0 {
		__phi96 = v103
		__phi97 = v96 + int32(20)
		v96 = __phi96
		v97 = __phi97
		goto L29
	} else {
		goto L31
	}
L30:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v97))) = int32(0)
	v128 = v96
	goto L15
L31:
	;
	v106 = *(*int32)(unsafe.Add(mBase, uint32(v96)+16))
	if v106 != 0 {
		__phi96 = v106
		__phi97 = v96 + int32(16)
		v96 = __phi96
		v97 = __phi97
		goto L29
	} else {
		goto L32
	}
L32:
	;
	goto L30
L33:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_memtest_alloc_and_test[5])) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v45)+4)) = v109 & int32(-2)
	*(*int32)(unsafe.Add(mBase, uint32(v53)+4)) = v57 | int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v45))) = v57
	goto L8
L34:
	;
	v137 = *(*int32)(unsafe.Add(mBase, uint32(v53)+28))
	v139 = v137 << (uint(int32(2)) % 32)
	v142 = *(*int32)(unsafe.Add(mBase, uint32(v139)+uint32(_c_F_memtest_alloc_and_test[6])))
	if v53 != v142 {
		goto L36
	} else {
		goto L37
	}
L35:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v128)+24)) = v75
	v159 = *(*int32)(unsafe.Add(mBase, uint32(v53)+16))
	if v159 == int32(0) {
		goto L43
	} else {
		goto L44
	}
L36:
	;
	v152 = *(*int32)(unsafe.Add(mBase, uint32(v75)+16))
	if v152 != v53 {
		goto L40
	} else {
		goto L41
	}
L37:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v139)+uint32(_c_F_memtest_alloc_and_test[6]))) = v128
	if v128 != 0 {
		goto L35
	} else {
		goto L38
	}
L38:
	;
	v145 = int32(0)
	v147 = *(*int32)(unsafe.Add(mBase, _c_F_memtest_alloc_and_test[7]))
	*(*int32)(unsafe.Add(mBase, _c_F_memtest_alloc_and_test[7])) = v147 & base.I32_rotl(int32(-2), v137)
	v169 = v57
	v170 = v53
	goto L11
L39:
	;
	if v128 == int32(0) {
		v169 = v57
		v170 = v53
		goto L11
	} else {
		goto L42
	}
L40:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v75)+20)) = v128
	goto L39
L41:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v75)+16)) = v128
	goto L39
L42:
	;
	goto L35
L43:
	;
	v164 = *(*int32)(unsafe.Add(mBase, uint32(v53)+20))
	if v164 == int32(0) {
		v169 = v57
		v170 = v53
		goto L11
	} else {
		goto L45
	}
L44:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v128)+16)) = v159
	*(*int32)(unsafe.Add(mBase, uint32(v159)+24)) = v128
	goto L43
L45:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v128)+20)) = v164
	*(*int32)(unsafe.Add(mBase, uint32(v164)+24)) = v128
	v169 = v57
	v170 = v53
	goto L11
L46:
	;
	v179 = *(*int32)(unsafe.Add(mBase, uint32(v45)+4))
	if v179&int32(1) == int32(0) {
		goto L9
	} else {
		goto L47
	}
L47:
	;
	if v179&int32(2) != 0 {
		goto L52
	} else {
		goto L53
	}
L48:
	;
	if base.Ui32(int32(255)) < base.Ui32(v345) {
		goto L86
	} else {
		goto L87
	}
L49:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v170)+4)) = v225 | int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v170+v225))) = v225
	if v170 != v209 {
		v345 = v225
		goto L48
	} else {
		goto L85
	}
L50:
	;
	if v242 == int32(0) {
		goto L49
	} else {
		goto L73
	}
L51:
	;
	v287 = int32(0)
	goto L50
L52:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v45)+4)) = v179 & int32(-2)
	*(*int32)(unsafe.Add(mBase, uint32(v170)+4)) = v169 | int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v170+v169))) = v169
	v345 = v169
	goto L48
L53:
	;
	v187 = *(*int32)(unsafe.Add(mBase, _c_F_memtest_alloc_and_test[8]))
	if v45 != v187 {
		goto L54
	} else {
		goto L55
	}
L54:
	;
	v209 = *(*int32)(unsafe.Add(mBase, _c_F_memtest_alloc_and_test[3]))
	if v45 != v209 {
		goto L57
	} else {
		goto L58
	}
L55:
	;
	v189 = int32(0)
	*(*int32)(unsafe.Add(mBase, _c_F_memtest_alloc_and_test[8])) = v170
	v193 = *(*int32)(unsafe.Add(mBase, _c_F_memtest_alloc_and_test[9]))
	v194 = v193 + v169
	*(*int32)(unsafe.Add(mBase, _c_F_memtest_alloc_and_test[9])) = v194
	*(*int32)(unsafe.Add(mBase, uint32(v170)+4)) = v194 | int32(1)
	v200 = *(*int32)(unsafe.Add(mBase, _c_F_memtest_alloc_and_test[3]))
	if v170 != v200 {
		goto L9
	} else {
		goto L56
	}
L56:
	;
	v202 = int32(0)
	*(*int32)(unsafe.Add(mBase, _c_F_memtest_alloc_and_test[5])) = v202
	*(*int32)(unsafe.Add(mBase, _c_F_memtest_alloc_and_test[3])) = v202
	goto L8
L57:
	;
	v225 = v179&int32(-8) + v169
	v226 = *(*int32)(unsafe.Add(mBase, uint32(v45)+12))
	if base.Ui32(int32(255)) < base.Ui32(v179) {
		goto L59
	} else {
		goto L60
	}
L58:
	;
	v211 = int32(0)
	*(*int32)(unsafe.Add(mBase, _c_F_memtest_alloc_and_test[3])) = v170
	v215 = *(*int32)(unsafe.Add(mBase, _c_F_memtest_alloc_and_test[5]))
	v216 = v215 + v169
	*(*int32)(unsafe.Add(mBase, _c_F_memtest_alloc_and_test[5])) = v216
	*(*int32)(unsafe.Add(mBase, uint32(v170)+4)) = v216 | int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v170+v216))) = v216
	goto L8
L59:
	;
	v242 = *(*int32)(unsafe.Add(mBase, uint32(v45)+24))
	if v226 == v45 {
		goto L63
	} else {
		goto L64
	}
L60:
	;
	v229 = *(*int32)(unsafe.Add(mBase, uint32(v45)+8))
	if v226 != v229 {
		goto L61
	} else {
		goto L62
	}
L61:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v229)+12)) = v226
	*(*int32)(unsafe.Add(mBase, uint32(v226)+8)) = v229
	goto L49
L62:
	;
	v231 = int32(0)
	v233 = *(*int32)(unsafe.Add(mBase, _c_F_memtest_alloc_and_test[4]))
	*(*int32)(unsafe.Add(mBase, _c_F_memtest_alloc_and_test[4])) = v233 & base.I32_rotl(int32(-2), int32(base.Ui32(v179)>>(uint(int32(3))%32)))
	goto L49
L63:
	;
	v247 = *(*int32)(unsafe.Add(mBase, uint32(v45)+20))
	if v247 == int32(0) {
		goto L66
	} else {
		goto L67
	}
L64:
	;
	v244 = *(*int32)(unsafe.Add(mBase, uint32(v45)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v244)+12)) = v226
	*(*int32)(unsafe.Add(mBase, uint32(v226)+8)) = v244
	v287 = v226
	goto L50
L65:
	;
	__phi263 = v257
	__phi264 = v258
	v263 = __phi263
	v264 = __phi264
	goto L69
L66:
	;
	v252 = *(*int32)(unsafe.Add(mBase, uint32(v45)+16))
	if v252 == int32(0) {
		goto L51
	} else {
		goto L68
	}
L67:
	;
	v257 = v247
	v258 = v45 + int32(20)
	goto L65
L68:
	;
	v257 = v252
	v258 = v45 + int32(16)
	goto L65
L69:
	;
	v270 = *(*int32)(unsafe.Add(mBase, uint32(v263)+20))
	if v270 != 0 {
		__phi263 = v270
		__phi264 = v263 + int32(20)
		v263 = __phi263
		v264 = __phi264
		goto L69
	} else {
		goto L71
	}
L70:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v264))) = int32(0)
	v287 = v263
	goto L50
L71:
	;
	v273 = *(*int32)(unsafe.Add(mBase, uint32(v263)+16))
	if v273 != 0 {
		__phi263 = v273
		__phi264 = v263 + int32(16)
		v263 = __phi263
		v264 = __phi264
		goto L69
	} else {
		goto L72
	}
L72:
	;
	goto L70
L73:
	;
	v296 = *(*int32)(unsafe.Add(mBase, uint32(v45)+28))
	v298 = v296 << (uint(int32(2)) % 32)
	v301 = *(*int32)(unsafe.Add(mBase, uint32(v298)+uint32(_c_F_memtest_alloc_and_test[6])))
	if v45 != v301 {
		goto L75
	} else {
		goto L76
	}
L74:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v287)+24)) = v242
	v318 = *(*int32)(unsafe.Add(mBase, uint32(v45)+16))
	if v318 == int32(0) {
		goto L82
	} else {
		goto L83
	}
L75:
	;
	v311 = *(*int32)(unsafe.Add(mBase, uint32(v242)+16))
	if v311 != v45 {
		goto L79
	} else {
		goto L80
	}
L76:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v298)+uint32(_c_F_memtest_alloc_and_test[6]))) = v287
	if v287 != 0 {
		goto L74
	} else {
		goto L77
	}
L77:
	;
	v304 = int32(0)
	v306 = *(*int32)(unsafe.Add(mBase, _c_F_memtest_alloc_and_test[7]))
	*(*int32)(unsafe.Add(mBase, _c_F_memtest_alloc_and_test[7])) = v306 & base.I32_rotl(int32(-2), v296)
	goto L49
L78:
	;
	if v287 == int32(0) {
		goto L49
	} else {
		goto L81
	}
L79:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v242)+20)) = v287
	goto L78
L80:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v242)+16)) = v287
	goto L78
L81:
	;
	goto L74
L82:
	;
	v323 = *(*int32)(unsafe.Add(mBase, uint32(v45)+20))
	if v323 == int32(0) {
		goto L49
	} else {
		goto L84
	}
L83:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v287)+16)) = v318
	*(*int32)(unsafe.Add(mBase, uint32(v318)+24)) = v287
	goto L82
L84:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v287)+20)) = v323
	*(*int32)(unsafe.Add(mBase, uint32(v323)+24)) = v287
	goto L49
L85:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_memtest_alloc_and_test[5])) = v225
	goto L8
L86:
	;
	if base.Ui32(int32(16777215)) < base.Ui32(v345) {
		v392 = int32(31)
		goto L91
	} else {
		goto L92
	}
L87:
	;
	v357 = v345 & int32(-8)
	v359 = v357 + int32(9128464)
	v361 = *(*int32)(unsafe.Add(mBase, _c_F_memtest_alloc_and_test[4]))
	v365 = int32(1) << (uint(int32(base.Ui32(v345)>>(uint(int32(3))%32))) % 32)
	if v361&v365 != 0 {
		goto L89
	} else {
		goto L90
	}
L88:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v357)+uint32(_c_F_memtest_alloc_and_test[10]))) = v170
	*(*int32)(unsafe.Add(mBase, uint32(v371)+12)) = v170
	*(*int32)(unsafe.Add(mBase, uint32(v170)+12)) = v359
	*(*int32)(unsafe.Add(mBase, uint32(v170)+8)) = v371
	goto L8
L89:
	;
	v370 = *(*int32)(unsafe.Add(mBase, uint32(v357)+uint32(_c_F_memtest_alloc_and_test[10])))
	v371 = v370
	goto L88
L90:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_memtest_alloc_and_test[4])) = v361 | v365
	v371 = v359
	goto L88
L91:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v170)+28)) = v392
	*(*int64)(unsafe.Add(mBase, uint32(v170)+16)) = int64(0)
	v397 = v392 << (uint(int32(2)) % 32)
	v401 = *(*int32)(unsafe.Add(mBase, _c_F_memtest_alloc_and_test[7]))
	v403 = int32(1) << (uint(v392) % 32)
	if v401&v403 != 0 {
		goto L96
	} else {
		goto L97
	}
L92:
	;
	v382 = base.I32_clz(int32(base.Ui32(v345) >> (uint(int32(8)) % 32)))
	v385 = int32(1)
	v392 = int32(base.Ui32(v345)>>(uint(int32(38)-v382)%32))&v385 - v382<<(uint(v385)%32) + int32(62)
	goto L91
L93:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v170+v464))) = v467
	*(*int32)(unsafe.Add(mBase, uint32(v170)+12)) = v466
	*(*int32)(unsafe.Add(mBase, uint32(v170+v462))) = v465
	v476 = int32(0)
	v478 = *(*int32)(unsafe.Add(mBase, _c_F_memtest_alloc_and_test[11]))
	v479 = int32(-1)
	v480 = v478 + v479
	if v480 != 0 {
		goto L105
	} else {
		goto L106
	}
L94:
	;
	v456 = *(*int32)(unsafe.Add(mBase, uint32(v426)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v456)+12)) = v170
	*(*int32)(unsafe.Add(mBase, uint32(v426)+8)) = v170
	v462 = int32(24)
	v464 = int32(8)
	v465 = int32(0)
	v466 = v426
	v467 = v456
	goto L93
L95:
	;
	v462 = v447
	v464 = v449
	v465 = v170
	v466 = v170
	v467 = v452
	goto L93
L96:
	;
	if v392 == int32(31) {
		goto L98
	} else {
		goto L99
	}
L97:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_memtest_alloc_and_test[7])) = v401 | v403
	*(*int32)(unsafe.Add(mBase, uint32(v397)+uint32(_c_F_memtest_alloc_and_test[6]))) = v170
	v447 = int32(8)
	v449 = int32(24)
	v452 = v397 + int32(9128728)
	goto L95
L98:
	;
	v418 = int32(0)
	goto L100
L99:
	;
	v418 = int32(25) - int32(base.Ui32(v392)>>(uint(int32(1))%32))
	goto L100
L100:
	;
	v420 = *(*int32)(unsafe.Add(mBase, uint32(v397)+uint32(_c_F_memtest_alloc_and_test[6])))
	v423 = v345 << (uint(v418) % 32)
	v426 = v420
	goto L101
L101:
	;
	v430 = *(*int32)(unsafe.Add(mBase, uint32(v426)+4))
	if v430&int32(-8) == v345 {
		goto L94
	} else {
		goto L103
	}
L102:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v440+int32(16)))) = v170
	v447 = int32(8)
	v449 = int32(24)
	v452 = v426
	goto L95
L103:
	;
	v440 = v426 + int32(base.Ui32(v423)>>(uint(int32(29))%32))&int32(4)
	v441 = *(*int32)(unsafe.Add(mBase, uint32(v440)+16))
	if v441 != 0 {
		v423 = v423 << (uint(int32(1)) % 32)
		v426 = v441
		goto L101
	} else {
		goto L104
	}
L104:
	;
	goto L102
L105:
	;
	v482 = v480
	goto L107
L106:
	;
	v482 = v479
	goto L107
L107:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_memtest_alloc_and_test[11])) = v482
	goto L9
}
func F_memtest_fill_random(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v20 int64
	_ = v20
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v37 int64
	_ = v37
	var v39 int64
	_ = v39
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v61 int64
	_ = v61
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v67 int64
	_ = v67
	var v71 int64
	_ = v71
	var v74 int64
	_ = v74
	var v77 int64
	_ = v77
	var v80 int32
	_ = v80
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v94 int64
	_ = v94
	var v96 int64
	_ = v96
	var v97 int32
	_ = v97
	var v100 int32
	_ = v100
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v118 int32
	_ = v118
	var v120 int32
	_ = v120
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v157 int32
	_ = v157
	var v162 int64
	_ = v162
	var v171 int64
	_ = v171
	var v180 int64
	_ = v180
	var v187 int32
	_ = v187
	if l1&int32(4095) != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	F__serverAssert(m, int32(_a_F_memtest_fill_random_0), int32(_a_F_memtest_fill_random_1), int32(154))
	mBase = m.M
	v187 = m.ExcPending
	if v187 != 0 {
		goto L16
	} else {
		goto L22
	}
L2:
	;
	v20 = base.I64_extend_i32_u(int32(base.Ui32(l1) >> (uint(int32(13)) % 32)))
	v22 = *(*int32)(unsafe.Add(mBase, _c_F_memtest_fill_random[0]))
	v24 = int32(base.Ui32(l1) >> (uint(int32(3)) % 32))
	v37 = int64(-3372857614747716250)
	v39 = int64(0)
	goto L3
L3:
	;
	if base.Ui32(l1) < base.Ui32(int32(8192)) {
		v171 = v37
		goto L5
	} else {
		goto L6
	}
L4:
	;
	return
L5:
	;
	v180 = v39 + int64(1)
	if v180 != int64(1024) {
		v37 = v171
		v39 = v180
		goto L3
	} else {
		goto L21
	}
L6:
	;
	v46 = int32(2)
	v48 = l0 + base.I32_wrap_i64(v39)<<(uint(v46)%32)
	v61 = v37
	v64 = v48
	v65 = v48 + v24<<(uint(v46)%32)
	v67 = int64(0)
	goto L7
L7:
	;
	v71 = int64(base.Ui64(v61)>>(uint(int64(12))%64)) ^ v61
	v74 = v71<<(uint(int64(25))%64) ^ v71
	v77 = int64(base.Ui64(v74)>>(uint(int64(27))%64)) ^ v74
	v80 = base.I32_wrap_i64(v77) * int32(1332534557)
	*(*int32)(unsafe.Add(mBase, uint32(v65))) = v80
	*(*int32)(unsafe.Add(mBase, uint32(v64))) = v80
	if l2 == int32(0) {
		goto L9
	} else {
		goto L10
	}
L8:
	;
	v171 = v77
	goto L5
L9:
	;
	v157 = int32(4096)
	v162 = v67 + int64(1)
	if v162 != v20 {
		v61 = v77
		v64 = v64 + v157
		v65 = v65 + v157
		v67 = v162
		goto L7
	} else {
		goto L20
	}
L10:
	;
	if v67&int64(65535) != int64(0) {
		goto L9
	} else {
		goto L11
	}
L11:
	;
	v89 = int32(0)
	v91 = *(*int32)(unsafe.Add(mBase, _c_F_memtest_fill_random[1]))
	v94 = int64(*(*uint32)(unsafe.Add(mBase, _c_F_memtest_fill_random[2])))
	v96 = base.I64_div_u_s((v67+v39*v20)*v94, base.I64_extend_i32_u(v24))
	v97 = base.I32_wrap_i64(v96)
	if v91 == v97 {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_memtest_fill_random[1])) = v97
	v140 = F_fflush(m, v22)
	mBase = m.M
	v141 = m.ExcPending
	if v141 != 0 {
		goto L16
	} else {
		goto L19
	}
L13:
	;
	v100 = v89
	goto L14
L14:
	;
	v115 = F_putchar(m, int32(82))
	mBase = m.M
	v116 = m.ExcPending
	if v116 != 0 {
		goto L16
	} else {
		goto L17
	}
L15:
	;
	goto L12
L16:
	;
	return
L17:
	;
	v118 = v100 + int32(1)
	v120 = *(*int32)(unsafe.Add(mBase, _c_F_memtest_fill_random[1]))
	if base.Ui32(v118) < base.Ui32(v97-v120) {
		v100 = v118
		goto L14
	} else {
		goto L18
	}
L18:
	;
	goto L15
L19:
	;
	goto L9
L20:
	;
	goto L8
L21:
	;
	goto L4
L22:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
}
