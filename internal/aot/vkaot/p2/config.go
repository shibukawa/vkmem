package p2

import (
	base "github.com/shibukawa/vkmem/internal/aot/vkaot/base"
	"unsafe"
)

func F_configEnumGetName(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
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
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
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
	var v46 int32
	_ = v46
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v64 int32
	_ = v64
	var v68 int32
	_ = v68
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v85 int32
	_ = v85
	v9 = m.G0
	v11 = v9 - int32(16)
	m.G0 = v11
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v13 != 0 {
		goto L4
	} else {
		goto L5
	}
L1:
	;
	m.G0 = v11 + int32(16)
	return v85
L2:
	;
	v78 = F_sdsnew(m, v77)
	mBase = m.M
	v79 = m.ExcPending
	if v79 != 0 {
		goto L10
	} else {
		goto L25
	}
L3:
	;
	F_sdsfree(m, v64)
	mBase = m.M
	v68 = m.ExcPending
	if v68 != 0 {
		goto L10
	} else {
		goto L24
	}
L4:
	;
	v16 = l0
	v20 = v13
	v21 = int32(0)
	v22 = l1
	goto L6
L5:
	;
	v64 = int32(0)
	goto L3
L6:
	;
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v16)+4))
	if l1 != v24 {
		goto L8
	} else {
		goto L9
	}
L7:
	;
	if v50 == int32(0) {
		v64 = v50
		goto L3
	} else {
		goto L22
	}
L8:
	;
	if l2 == int32(0) {
		v50 = v21
		v51 = v22
		goto L12
	} else {
		goto L13
	}
L9:
	;
	F_sdsfree(m, v21)
	mBase = m.M
	v29 = m.ExcPending
	if v29 != 0 {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	return int32(0)
L11:
	;
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
	v77 = v30
	goto L2
L12:
	;
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v16)+8))
	if v54 != 0 {
		v16 = v16 + int32(8)
		v20 = v54
		v21 = v50
		v22 = v51
		goto L6
	} else {
		goto L21
	}
L13:
	;
	if v24 == int32(0) {
		v50 = v21
		v51 = v22
		goto L12
	} else {
		goto L14
	}
L14:
	;
	if v24 != v24&v22 {
		v50 = v21
		v51 = v22
		goto L12
	} else {
		goto L15
	}
L15:
	;
	if v21 == int32(0) {
		goto L17
	} else {
		goto L18
	}
L16:
	;
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v16)+4))
	v50 = v45
	v51 = v22 & (v46 ^ int32(-1))
	goto L12
L17:
	;
	v43 = F_sdsnew(m, v20)
	mBase = m.M
	v44 = m.ExcPending
	if v44 != 0 {
		goto L10
	} else {
		goto L20
	}
L18:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11))) = v20
	v41 = F_sdscatfmt(m, v21, int32(_a45), v11)
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L10
	} else {
		goto L19
	}
L19:
	;
	v45 = v41
	goto L16
L20:
	;
	v45 = v43
	goto L16
L21:
	;
	goto L7
L22:
	;
	if v51 == int32(0) {
		v85 = v50
		goto L1
	} else {
		goto L23
	}
L23:
	;
	v64 = v50
	goto L3
L24:
	;
	v77 = int32(_a288)
	goto L2
L25:
	;
	v85 = v78
	goto L1
}
func F_configGetCommand(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v19 int32
	_ = v19
	var v26 int32
	_ = v26
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v83 int32
	_ = v83
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v96 int64
	_ = v96
	var v97 int64
	_ = v97
	var v98 int64
	_ = v98
	var v99 int64
	_ = v99
	var v100 int64
	_ = v100
	var v101 int64
	_ = v101
	var v102 int64
	_ = v102
	var v104 int64
	_ = v104
	var v106 int64
	_ = v106
	var v108 int64
	_ = v108
	var v110 int64
	_ = v110
	var v112 int64
	_ = v112
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v119 int32
	_ = v119
	var v122 int32
	_ = v122
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
	var v134 int32
	_ = v134
	var v138 int32
	_ = v138
	var v140 int32
	_ = v140
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v154 int32
	_ = v154
	var v159 int32
	_ = v159
	var v163 int32
	_ = v163
	var v167 int32
	_ = v167
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
	var v180 int32
	_ = v180
	var v182 int32
	_ = v182
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v196 int32
	_ = v196
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v206 int32
	_ = v206
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v237 int32
	_ = v237
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v253 int32
	_ = v253
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v261 int32
	_ = v261
	var v265 int32
	_ = v265
	var v266 int32
	_ = v266
	var v267 int32
	_ = v267
	var v270 int32
	_ = v270
	var v272 int32
	_ = v272
	var v274 int64
	_ = v274
	var v275 int64
	_ = v275
	var v276 int64
	_ = v276
	var v277 int64
	_ = v277
	var v278 int64
	_ = v278
	var v279 int64
	_ = v279
	var v280 int64
	_ = v280
	var v282 int64
	_ = v282
	var v284 int64
	_ = v284
	var v286 int64
	_ = v286
	var v288 int64
	_ = v288
	var v290 int64
	_ = v290
	var v292 int32
	_ = v292
	var v293 int32
	_ = v293
	var v294 int32
	_ = v294
	var v297 int32
	_ = v297
	var v300 int32
	_ = v300
	var v301 int32
	_ = v301
	var v302 int32
	_ = v302
	var v304 int32
	_ = v304
	var v307 int32
	_ = v307
	var v308 int32
	_ = v308
	var v312 int32
	_ = v312
	var v316 int32
	_ = v316
	var v318 int32
	_ = v318
	var v325 int32
	_ = v325
	var v326 int32
	_ = v326
	var v327 int32
	_ = v327
	var v332 int32
	_ = v332
	var v337 int32
	_ = v337
	var v341 int32
	_ = v341
	var v345 int32
	_ = v345
	var v356 int32
	_ = v356
	var v358 int32
	_ = v358
	var v361 int32
	_ = v361
	var v364 int32
	_ = v364
	var v365 int32
	_ = v365
	var v367 int32
	_ = v367
	var v368 int32
	_ = v368
	var v369 int32
	_ = v369
	var v380 int32
	_ = v380
	var v381 int32
	_ = v381
	var v384 int32
	_ = v384
	var v387 int32
	_ = v387
	var v388 int32
	_ = v388
	var v392 int32
	_ = v392
	var v396 int32
	_ = v396
	var v397 int32
	_ = v397
	var v398 int32
	_ = v398
	var v401 int32
	_ = v401
	var v403 int32
	_ = v403
	var v405 int64
	_ = v405
	var v406 int64
	_ = v406
	var v407 int64
	_ = v407
	var v408 int64
	_ = v408
	var v409 int64
	_ = v409
	var v410 int64
	_ = v410
	var v411 int64
	_ = v411
	var v413 int64
	_ = v413
	var v415 int64
	_ = v415
	var v417 int64
	_ = v417
	var v419 int64
	_ = v419
	var v421 int64
	_ = v421
	var v423 int32
	_ = v423
	var v424 int32
	_ = v424
	var v425 int32
	_ = v425
	var v428 int32
	_ = v428
	var v431 int32
	_ = v431
	var v432 int32
	_ = v432
	var v433 int32
	_ = v433
	var v435 int32
	_ = v435
	var v438 int32
	_ = v438
	var v439 int32
	_ = v439
	var v443 int32
	_ = v443
	var v447 int32
	_ = v447
	var v449 int32
	_ = v449
	var v456 int32
	_ = v456
	var v457 int32
	_ = v457
	var v458 int32
	_ = v458
	var v463 int32
	_ = v463
	var v468 int32
	_ = v468
	var v472 int32
	_ = v472
	var v476 int32
	_ = v476
	var v490 int32
	_ = v490
	var v492 int32
	_ = v492
	var v496 int32
	_ = v496
	var v504 int32
	_ = v504
	var v511 int32
	_ = v511
	var v512 int32
	_ = v512
	var v514 int32
	_ = v514
	var v515 int32
	_ = v515
	var v517 int32
	_ = v517
	var v519 int32
	_ = v519
	var v531 int32
	_ = v531
	v11 = F_dictCreate(m, int32(_a492))
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v13 < int32(3) {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v231 = F_dictGetIterator(m, v11)
	mBase = m.M
	v232 = m.ExcPending
	if v232 != 0 {
		goto L1
	} else {
		goto L63
	}
L4:
	;
	v19 = int32(0)
	goto L5
L5:
	;
	v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v26+v19<<(uint(int32(2))%32))+8))
	v31 = F_objectGetVal(m, v30)
	mBase = m.M
	v33 = F_strcspn(m, v31, int32(_a493))
	mBase = m.M
	v34 = v31 + v33
	v36 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v34))))
	if v36 != 0 {
		goto L10
	} else {
		goto L11
	}
L6:
	;
	goto L3
L7:
	;
	v217 = v19 + int32(1)
	v218 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v217 < v218+int32(-2) {
		v19 = v217
		goto L5
	} else {
		goto L62
	}
L8:
	;
	v52 = *(*int32)(unsafe.Add(mBase, _consts[230]))
	v53 = F_dictGetIterator(m, v52)
	mBase = m.M
	v54 = m.ExcPending
	if v54 != 0 {
		goto L1
	} else {
		goto L21
	}
L9:
	;
	if v37 != 0 {
		goto L8
	} else {
		goto L13
	}
L10:
	;
	v37 = v34
	goto L12
L11:
	;
	v37 = int32(0)
	goto L12
L12:
	;
	goto L9
L13:
	;
	v38 = F_dictFind(m, v11, v31)
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L1
	} else {
		goto L14
	}
L14:
	;
	if v38 != 0 {
		goto L7
	} else {
		goto L15
	}
L15:
	;
	v41 = *(*int32)(unsafe.Add(mBase, _consts[230]))
	v42 = F_dictFind(m, v41, v31)
	mBase = m.M
	v43 = m.ExcPending
	if v43 != 0 {
		goto L1
	} else {
		goto L16
	}
L16:
	;
	if v42 == int32(0) {
		goto L7
	} else {
		goto L17
	}
L17:
	;
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v42)+8))
	goto L18
L18:
	;
	if v46 == int32(0) {
		goto L7
	} else {
		goto L19
	}
L19:
	;
	v49 = F_dictAdd(m, v11, v31, v46)
	mBase = m.M
	v50 = m.ExcPending
	if v50 != 0 {
		goto L1
	} else {
		goto L20
	}
L20:
	;
	goto L7
L21:
	;
	goto L23
L22:
	;
	F_dictReleaseIterator(m, v53)
	mBase = m.M
	v206 = m.ExcPending
	if v206 != 0 {
		goto L1
	} else {
		goto L61
	}
L23:
	;
	v71 = v53 + int32(20)
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v53)+16))
	if v72 != 0 {
		goto L27
	} else {
		goto L28
	}
L25:
	;
	if v167 == int32(0) {
		goto L22
	} else {
		goto L51
	}
L26:
	;
	v78 = v71
	v79 = v75
	goto L29
L27:
	;
	v75 = int32(1)
	goto L26
L28:
	;
	v75 = int32(0)
	goto L26
L29:
	;
	switch v79 {
	case 0:
		goto L34
	default:
		goto L33
	}
L31:
	;
	v79 = int32(0)
	goto L29
L32:
	;
	goto L25
L33:
	;
	v159 = *(*int32)(unsafe.Add(mBase, uint32(v78)))
	*(*int32)(unsafe.Add(mBase, uint32(v53)+16)) = v159
	if v159 == int32(0) {
		goto L31
	} else {
		goto L50
	}
L34:
	;
	v83 = *(*int32)(unsafe.Add(mBase, uint32(v53)+4))
	if v83 != int32(-1) {
		v122 = v83
		goto L35
	} else {
		goto L36
	}
L35:
	;
	v123 = int32(1)
	v124 = v122 + v123
	*(*int32)(unsafe.Add(mBase, uint32(v53)+4)) = v124
	v126 = int32(0)
	v129 = *(*int32)(unsafe.Add(mBase, uint32(v53)))
	v130 = *(*int32)(unsafe.Add(mBase, uint32(v53)+8))
	v134 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v129+v130+int32(26)))))
	if v134 == int32(255) {
		goto L44
	} else {
		goto L45
	}
L36:
	;
	v87 = *(*int32)(unsafe.Add(mBase, uint32(v53)+8))
	if v87 != 0 {
		v122 = int32(-1)
		goto L35
	} else {
		goto L37
	}
L37:
	;
	v88 = *(*int32)(unsafe.Add(mBase, uint32(v53)))
	v89 = *(*int32)(unsafe.Add(mBase, uint32(v53)+12))
	if v89 == int32(0) {
		goto L39
	} else {
		goto L40
	}
L38:
	;
	v116 = *(*int32)(unsafe.Add(mBase, uint32(v115)+20))
	if v116 != int32(-1) {
		goto L41
	} else {
		goto L42
	}
L39:
	;
	v96 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v88)+16)))
	v97 = int64(*(*int8)(unsafe.Add(mBase, uint32(v88)+27)))
	v98 = int64(*(*int32)(unsafe.Add(mBase, uint32(v88)+8)))
	v99 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v88)+12)))
	v100 = int64(*(*int8)(unsafe.Add(mBase, uint32(v88)+26)))
	v101 = int64(*(*int32)(unsafe.Add(mBase, uint32(v88)+4)))
	v102 = F_wangHash64(m, v101)
	mBase = m.M
	v104 = F_wangHash64(m, v100+v102)
	mBase = m.M
	v106 = F_wangHash64(m, v99+v104)
	mBase = m.M
	v108 = F_wangHash64(m, v98+v106)
	mBase = m.M
	v110 = F_wangHash64(m, v97+v108)
	mBase = m.M
	v112 = F_wangHash64(m, v96+v110)
	mBase = m.M
	*(*int64)(unsafe.Add(mBase, uint32(v53)+24)) = v112
	v114 = *(*int32)(unsafe.Add(mBase, uint32(v53)))
	v115 = v114
	goto L38
L40:
	;
	v92 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v88)+24)))
	v94 = v92 + int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v88)+24)) = uint16(v94)
	v115 = v88
	goto L38
L41:
	;
	v122 = v116 + int32(-1)
	goto L35
L42:
	;
	v119 = *(*int32)(unsafe.Add(mBase, uint32(v53)+4))
	v122 = v119
	goto L35
L43:
	;
	v149 = int32(2)
	v154 = *(*int32)(unsafe.Add(mBase, uint32(v129+v147<<(uint(v149)%32)+int32(4))))
	v78 = v154 + v148<<(uint(v149)%32)
	v79 = int32(1)
	goto L29
L44:
	;
	v138 = v126
	goto L46
L45:
	;
	v138 = v123 << (uint(v134) % 32)
	goto L46
L46:
	;
	if v124 < v138 {
		v147 = v130
		v148 = v124
		goto L43
	} else {
		goto L47
	}
L47:
	;
	if v130 != 0 {
		v167 = v126
		goto L32
	} else {
		goto L48
	}
L48:
	;
	v140 = *(*int32)(unsafe.Add(mBase, uint32(v129)+20))
	if v140 == int32(-1) {
		v167 = v126
		goto L32
	} else {
		goto L49
	}
L49:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v53)+4)) = int64(4294967296)
	v147 = int32(1)
	v148 = int32(0)
	goto L43
L50:
	;
	v163 = *(*int32)(unsafe.Add(mBase, uint32(v159)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v71))) = v163
	v167 = v159
	goto L32
L51:
	;
	v173 = *(*int32)(unsafe.Add(mBase, uint32(v167)+8))
	goto L52
L52:
	;
	v174 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v173)+8)))
	if v174&int32(16) != 0 {
		goto L23
	} else {
		goto L53
	}
L53:
	;
	v177 = *(*int32)(unsafe.Add(mBase, uint32(v173)))
	v178 = F_dictFind(m, v11, v177)
	mBase = m.M
	v179 = m.ExcPending
	if v179 != 0 {
		goto L1
	} else {
		goto L54
	}
L54:
	;
	if v178 != 0 {
		goto L23
	} else {
		goto L55
	}
L55:
	;
	v180 = *(*int32)(unsafe.Add(mBase, uint32(v167)))
	goto L56
L56:
	;
	v182 = int32(0)
	v185 = m.G0
	v186 = int32(16)
	v187 = v185 - v186
	m.G0 = v187
	v189 = F_strlen(m, v31)
	mBase = m.M
	v190 = F_strlen(m, v180)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v187)+12)) = v182
	v196 = F_stringmatchlen_impl(m, v31, v189, v180, v190, int32(1), v187+int32(12), v182)
	mBase = m.M
	m.G0 = v187 + v186
	goto L57
L57:
	;
	if v196 == int32(0) {
		goto L23
	} else {
		goto L58
	}
L58:
	;
	v202 = *(*int32)(unsafe.Add(mBase, uint32(v167)))
	goto L59
L59:
	;
	v203 = F_dictAdd(m, v11, v202, v173)
	mBase = m.M
	v204 = m.ExcPending
	if v204 != 0 {
		goto L1
	} else {
		goto L60
	}
L60:
	;
	goto L23
L61:
	;
	goto L7
L62:
	;
	goto L6
L63:
	;
	v233 = *(*int32)(unsafe.Add(mBase, uint32(v11)+16))
	v234 = *(*int32)(unsafe.Add(mBase, uint32(v11)+12))
	v235 = v233 + v234
	F_addReplyMapLen(m, l0, v235)
	mBase = m.M
	v237 = m.ExcPending
	if v237 != 0 {
		goto L1
	} else {
		goto L64
	}
L64:
	;
	v240 = F_valkey_malloc(m, v235<<(uint(int32(3))%32))
	mBase = m.M
	v241 = m.ExcPending
	if v241 != 0 {
		goto L1
	} else {
		goto L65
	}
L65:
	;
	v249 = v231 + int32(20)
	v250 = *(*int32)(unsafe.Add(mBase, uint32(v231)+16))
	if v250 != 0 {
		goto L69
	} else {
		goto L70
	}
L66:
	;
	F_dictReleaseIterator(m, v231)
	mBase = m.M
	v490 = m.ExcPending
	if v490 != 0 {
		goto L1
	} else {
		goto L126
	}
L67:
	;
	if v345 == int32(0) {
		goto L66
	} else {
		goto L93
	}
L68:
	;
	v256 = v249
	v257 = v253
	goto L71
L69:
	;
	v253 = int32(1)
	goto L68
L70:
	;
	v253 = int32(0)
	goto L68
L71:
	;
	switch v257 {
	case 0:
		goto L76
	default:
		goto L75
	}
L73:
	;
	v257 = int32(0)
	goto L71
L74:
	;
	goto L67
L75:
	;
	v337 = *(*int32)(unsafe.Add(mBase, uint32(v256)))
	*(*int32)(unsafe.Add(mBase, uint32(v231)+16)) = v337
	if v337 == int32(0) {
		goto L73
	} else {
		goto L92
	}
L76:
	;
	v261 = *(*int32)(unsafe.Add(mBase, uint32(v231)+4))
	if v261 != int32(-1) {
		v300 = v261
		goto L77
	} else {
		goto L78
	}
L77:
	;
	v301 = int32(1)
	v302 = v300 + v301
	*(*int32)(unsafe.Add(mBase, uint32(v231)+4)) = v302
	v304 = int32(0)
	v307 = *(*int32)(unsafe.Add(mBase, uint32(v231)))
	v308 = *(*int32)(unsafe.Add(mBase, uint32(v231)+8))
	v312 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v307+v308+int32(26)))))
	if v312 == int32(255) {
		goto L86
	} else {
		goto L87
	}
L78:
	;
	v265 = *(*int32)(unsafe.Add(mBase, uint32(v231)+8))
	if v265 != 0 {
		v300 = int32(-1)
		goto L77
	} else {
		goto L79
	}
L79:
	;
	v266 = *(*int32)(unsafe.Add(mBase, uint32(v231)))
	v267 = *(*int32)(unsafe.Add(mBase, uint32(v231)+12))
	if v267 == int32(0) {
		goto L81
	} else {
		goto L82
	}
L80:
	;
	v294 = *(*int32)(unsafe.Add(mBase, uint32(v293)+20))
	if v294 != int32(-1) {
		goto L83
	} else {
		goto L84
	}
L81:
	;
	v274 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v266)+16)))
	v275 = int64(*(*int8)(unsafe.Add(mBase, uint32(v266)+27)))
	v276 = int64(*(*int32)(unsafe.Add(mBase, uint32(v266)+8)))
	v277 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v266)+12)))
	v278 = int64(*(*int8)(unsafe.Add(mBase, uint32(v266)+26)))
	v279 = int64(*(*int32)(unsafe.Add(mBase, uint32(v266)+4)))
	v280 = F_wangHash64(m, v279)
	mBase = m.M
	v282 = F_wangHash64(m, v278+v280)
	mBase = m.M
	v284 = F_wangHash64(m, v277+v282)
	mBase = m.M
	v286 = F_wangHash64(m, v276+v284)
	mBase = m.M
	v288 = F_wangHash64(m, v275+v286)
	mBase = m.M
	v290 = F_wangHash64(m, v274+v288)
	mBase = m.M
	*(*int64)(unsafe.Add(mBase, uint32(v231)+24)) = v290
	v292 = *(*int32)(unsafe.Add(mBase, uint32(v231)))
	v293 = v292
	goto L80
L82:
	;
	v270 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v266)+24)))
	v272 = v270 + int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v266)+24)) = uint16(v272)
	v293 = v266
	goto L80
L83:
	;
	v300 = v294 + int32(-1)
	goto L77
L84:
	;
	v297 = *(*int32)(unsafe.Add(mBase, uint32(v231)+4))
	v300 = v297
	goto L77
L85:
	;
	v327 = int32(2)
	v332 = *(*int32)(unsafe.Add(mBase, uint32(v307+v325<<(uint(v327)%32)+int32(4))))
	v256 = v332 + v326<<(uint(v327)%32)
	v257 = int32(1)
	goto L71
L86:
	;
	v316 = v304
	goto L88
L87:
	;
	v316 = v301 << (uint(v312) % 32)
	goto L88
L88:
	;
	if v302 < v316 {
		v325 = v308
		v326 = v302
		goto L85
	} else {
		goto L89
	}
L89:
	;
	if v308 != 0 {
		v345 = v304
		goto L74
	} else {
		goto L90
	}
L90:
	;
	v318 = *(*int32)(unsafe.Add(mBase, uint32(v307)+20))
	if v318 == int32(-1) {
		v345 = v304
		goto L74
	} else {
		goto L91
	}
L91:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v231)+4)) = int64(4294967296)
	v325 = int32(1)
	v326 = int32(0)
	goto L85
L92:
	;
	v341 = *(*int32)(unsafe.Add(mBase, uint32(v337)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v249))) = v341
	v345 = v337
	goto L74
L93:
	;
	v356 = v345
	v358 = int32(0)
	goto L94
L94:
	;
	v361 = *(*int32)(unsafe.Add(mBase, uint32(v356)+8))
	goto L96
L95:
	;
	goto L66
L96:
	;
	v364 = v240 + v358<<(uint(int32(3))%32)
	v365 = *(*int32)(unsafe.Add(mBase, uint32(v356)))
	goto L97
L97:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v364))) = v365
	v367 = *(*int32)(unsafe.Add(mBase, uint32(v361)+24))
	v368 = m.T0[v367].(func(*base.Module, int32) int32)(m, v361)
	mBase = m.M
	v369 = m.ExcPending
	if v369 != 0 {
		goto L1
	} else {
		goto L98
	}
L98:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v364)+4)) = v368
	v380 = v231 + int32(20)
	v381 = *(*int32)(unsafe.Add(mBase, uint32(v231)+16))
	if v381 != 0 {
		goto L101
	} else {
		goto L102
	}
L99:
	;
	if v476 != 0 {
		v356 = v476
		v358 = v358 + int32(1)
		goto L94
	} else {
		goto L125
	}
L100:
	;
	v387 = v380
	v388 = v384
	goto L103
L101:
	;
	v384 = int32(1)
	goto L100
L102:
	;
	v384 = int32(0)
	goto L100
L103:
	;
	switch v388 {
	case 0:
		goto L108
	default:
		goto L107
	}
L105:
	;
	v388 = int32(0)
	goto L103
L106:
	;
	goto L99
L107:
	;
	v468 = *(*int32)(unsafe.Add(mBase, uint32(v387)))
	*(*int32)(unsafe.Add(mBase, uint32(v231)+16)) = v468
	if v468 == int32(0) {
		goto L105
	} else {
		goto L124
	}
L108:
	;
	v392 = *(*int32)(unsafe.Add(mBase, uint32(v231)+4))
	if v392 != int32(-1) {
		v431 = v392
		goto L109
	} else {
		goto L110
	}
L109:
	;
	v432 = int32(1)
	v433 = v431 + v432
	*(*int32)(unsafe.Add(mBase, uint32(v231)+4)) = v433
	v435 = int32(0)
	v438 = *(*int32)(unsafe.Add(mBase, uint32(v231)))
	v439 = *(*int32)(unsafe.Add(mBase, uint32(v231)+8))
	v443 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v438+v439+int32(26)))))
	if v443 == int32(255) {
		goto L118
	} else {
		goto L119
	}
L110:
	;
	v396 = *(*int32)(unsafe.Add(mBase, uint32(v231)+8))
	if v396 != 0 {
		v431 = int32(-1)
		goto L109
	} else {
		goto L111
	}
L111:
	;
	v397 = *(*int32)(unsafe.Add(mBase, uint32(v231)))
	v398 = *(*int32)(unsafe.Add(mBase, uint32(v231)+12))
	if v398 == int32(0) {
		goto L113
	} else {
		goto L114
	}
L112:
	;
	v425 = *(*int32)(unsafe.Add(mBase, uint32(v424)+20))
	if v425 != int32(-1) {
		goto L115
	} else {
		goto L116
	}
L113:
	;
	v405 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v397)+16)))
	v406 = int64(*(*int8)(unsafe.Add(mBase, uint32(v397)+27)))
	v407 = int64(*(*int32)(unsafe.Add(mBase, uint32(v397)+8)))
	v408 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v397)+12)))
	v409 = int64(*(*int8)(unsafe.Add(mBase, uint32(v397)+26)))
	v410 = int64(*(*int32)(unsafe.Add(mBase, uint32(v397)+4)))
	v411 = F_wangHash64(m, v410)
	mBase = m.M
	v413 = F_wangHash64(m, v409+v411)
	mBase = m.M
	v415 = F_wangHash64(m, v408+v413)
	mBase = m.M
	v417 = F_wangHash64(m, v407+v415)
	mBase = m.M
	v419 = F_wangHash64(m, v406+v417)
	mBase = m.M
	v421 = F_wangHash64(m, v405+v419)
	mBase = m.M
	*(*int64)(unsafe.Add(mBase, uint32(v231)+24)) = v421
	v423 = *(*int32)(unsafe.Add(mBase, uint32(v231)))
	v424 = v423
	goto L112
L114:
	;
	v401 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v397)+24)))
	v403 = v401 + int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v397)+24)) = uint16(v403)
	v424 = v397
	goto L112
L115:
	;
	v431 = v425 + int32(-1)
	goto L109
L116:
	;
	v428 = *(*int32)(unsafe.Add(mBase, uint32(v231)+4))
	v431 = v428
	goto L109
L117:
	;
	v458 = int32(2)
	v463 = *(*int32)(unsafe.Add(mBase, uint32(v438+v456<<(uint(v458)%32)+int32(4))))
	v387 = v463 + v457<<(uint(v458)%32)
	v388 = int32(1)
	goto L103
L118:
	;
	v447 = v435
	goto L120
L119:
	;
	v447 = v432 << (uint(v443) % 32)
	goto L120
L120:
	;
	if v433 < v447 {
		v456 = v439
		v457 = v433
		goto L117
	} else {
		goto L121
	}
L121:
	;
	if v439 != 0 {
		v476 = v435
		goto L106
	} else {
		goto L122
	}
L122:
	;
	v449 = *(*int32)(unsafe.Add(mBase, uint32(v438)+20))
	if v449 == int32(-1) {
		v476 = v435
		goto L106
	} else {
		goto L123
	}
L123:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v231)+4)) = int64(4294967296)
	v456 = int32(1)
	v457 = int32(0)
	goto L117
L124:
	;
	v472 = *(*int32)(unsafe.Add(mBase, uint32(v468)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v380))) = v472
	v476 = v468
	goto L106
L125:
	;
	goto L95
L126:
	;
	F_dictRelease(m, v11)
	mBase = m.M
	v492 = m.ExcPending
	if v492 != 0 {
		goto L1
	} else {
		goto L127
	}
L127:
	;
	F_qsort(m, v240, v235, int32(8), int32(415))
	mBase = m.M
	v496 = m.ExcPending
	if v496 != 0 {
		goto L1
	} else {
		goto L128
	}
L128:
	;
	if v235 < int32(1) {
		goto L129
	} else {
		goto L130
	}
L129:
	;
	F_valkey_free(m, v240)
	mBase = m.M
	v531 = m.ExcPending
	if v531 != 0 {
		goto L1
	} else {
		goto L136
	}
L130:
	;
	v504 = int32(0)
	goto L131
L131:
	;
	v511 = v240 + v504<<(uint(int32(3))%32)
	v512 = *(*int32)(unsafe.Add(mBase, uint32(v511)))
	F_addReplyBulkCString(m, l0, v512)
	mBase = m.M
	v514 = m.ExcPending
	if v514 != 0 {
		goto L1
	} else {
		goto L133
	}
L132:
	;
	goto L129
L133:
	;
	v515 = *(*int32)(unsafe.Add(mBase, uint32(v511)+4))
	F_addReplyBulkSds(m, l0, v515)
	mBase = m.M
	v517 = m.ExcPending
	if v517 != 0 {
		goto L1
	} else {
		goto L134
	}
L134:
	;
	v519 = v504 + int32(1)
	if v519 != v235 {
		v504 = v519
		goto L131
	} else {
		goto L135
	}
L135:
	;
	goto L132
L136:
	;
	return
}
func F_configHelpCommand(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v10 int32
	_ = v10
	var v15 int64
	_ = v15
	var v20 int64
	_ = v20
	var v23 int64
	_ = v23
	var v26 int64
	_ = v26
	var v29 int32
	_ = v29
	v2 = int32(0)
	v3 = m.G0
	v5 = v3 - int32(48)
	m.G0 = v5
	v10 = *(*int32)(unsafe.Add(mBase, _consts[291]))
	*(*int32)(unsafe.Add(mBase, uint32(v5+int32(32)))) = v10
	v15 = *(*int64)(unsafe.Add(mBase, _consts[292]))
	*(*int64)(unsafe.Add(mBase, uint32(v5+int32(24)))) = v15
	v20 = *(*int64)(unsafe.Add(mBase, _consts[293]))
	*(*int64)(unsafe.Add(mBase, uint32(v5+int32(16)))) = v20
	v23 = *(*int64)(unsafe.Add(mBase, _consts[294]))
	*(*int64)(unsafe.Add(mBase, uint32(v5)+8)) = v23
	v26 = *(*int64)(unsafe.Add(mBase, _consts[295]))
	*(*int64)(unsafe.Add(mBase, uint32(v5))) = v26
	F_addReplyHelp(m, l0, v5)
	mBase = m.M
	v29 = m.ExcPending
	if v29 != 0 {
		return
	} else {
		m.G0 = v5 + int32(48)
		return
	}
}
func F_configResetStatCommand(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int64
	_ = v3
	var v14 int32
	_ = v14
	var v27 int32
	_ = v27
	var v31 int32
	_ = v31
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v90 int32
	_ = v90
	var v94 int32
	_ = v94
	var v98 int32
	_ = v98
	var v102 int32
	_ = v102
	var v106 int32
	_ = v106
	var v110 int32
	_ = v110
	var v114 int32
	_ = v114
	var v118 int32
	_ = v118
	var v128 int32
	_ = v128
	var v133 int32
	_ = v133
	var v140 int32
	_ = v140
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
	v2 = int32(0)
	v3 = int64(0)
	*(*int64)(unsafe.Add(mBase, _consts[296])) = v3
	*(*int64)(unsafe.Add(mBase, _consts[297])) = v3
	*(*int32)(unsafe.Add(mBase, _consts[43])) = v2
	v14 = F___memset(m, int32(_a525), v2, int32(96))
	mBase = m.M
	*(*int64)(unsafe.Add(mBase, _consts[298])) = v3
	*(*int64)(unsafe.Add(mBase, _consts[299])) = v3
	*(*uint8)(unsafe.Add(mBase, _consts[300])) = uint8(v2)
	v27 = F___memset(m, int32(_a526), v2, int32(176))
	mBase = m.M
	v31 = F___memset(m, int32(_a527), v2, int32(80))
	mBase = m.M
	*(*int64)(unsafe.Add(mBase, _consts[301])) = v3
	*(*int64)(unsafe.Add(mBase, _consts[302])) = v3
	*(*int64)(unsafe.Add(mBase, _consts[124])) = v3
	*(*int64)(unsafe.Add(mBase, _consts[303])) = v3
	*(*int64)(unsafe.Add(mBase, _consts[304])) = v3
	*(*int64)(unsafe.Add(mBase, _consts[305])) = v3
	*(*int64)(unsafe.Add(mBase, _consts[306])) = v3
	*(*int64)(unsafe.Add(mBase, _consts[94])) = v3
	*(*int64)(unsafe.Add(mBase, _consts[307])) = v3
	*(*int64)(unsafe.Add(mBase, _consts[308])) = v3
	*(*int64)(unsafe.Add(mBase, _consts[309])) = v3
	*(*int64)(unsafe.Add(mBase, _consts[310])) = v3
	*(*int64)(unsafe.Add(mBase, _consts[311])) = v3
	*(*int64)(unsafe.Add(mBase, _consts[312])) = v3
	*(*int64)(unsafe.Add(mBase, _consts[313])) = v3
	*(*int64)(unsafe.Add(mBase, _consts[314])) = v3
	*(*int64)(unsafe.Add(mBase, _consts[315])) = v3
	v85 = int32(148)
	v86 = F___memset(m, int32(_a528), v2, v85)
	mBase = m.M
	v90 = F___memset(m, int32(_a529), v2, v85)
	mBase = m.M
	v94 = F___memset(m, int32(_a530), v2, v85)
	mBase = m.M
	v98 = F___memset(m, int32(_a531), v2, v85)
	mBase = m.M
	v102 = F___memset(m, int32(_a532), v2, v85)
	mBase = m.M
	v106 = F___memset(m, int32(_a533), v2, v85)
	mBase = m.M
	v110 = F___memset(m, int32(_a534), v2, v85)
	mBase = m.M
	v114 = F___memset(m, int32(_a535), v2, v85)
	mBase = m.M
	v118 = F___memset(m, int32(_a536), v2, v85)
	mBase = m.M
	*(*int64)(unsafe.Add(mBase, _consts[316])) = v3
	*(*int64)(unsafe.Add(mBase, _consts[317])) = v3
	F_lazyfreeResetStats(m)
	mBase = m.M
	v128 = *(*int32)(unsafe.Add(mBase, _consts[139]))
	if v128 == int32(0) {
	} else {
		F_clusterSlotStatResetAll(m)
		mBase = m.M
		v133 = *(*int32)(unsafe.Add(mBase, _consts[136]))
		*(*int64)(unsafe.Add(mBase, uint32(v133)+uint32(_consts[200]))) = int64(0)
		v140 = F___memset(m, v133+int32(65680), int32(0), int32(224))
		mBase = m.M
	}
	v143 = *(*int32)(unsafe.Add(mBase, _consts[318]))
	F_resetCommandTableStats(m, v143)
	mBase = m.M
	v145 = m.ExcPending
	if v145 != 0 {
		return
	} else {
		F_resetErrorTableStats(m)
		mBase = m.M
		v147 = m.ExcPending
		if v147 != 0 {
			return
		} else {
			v149 = *(*int32)(unsafe.Add(mBase, _consts[84]))
			F_addReply(m, l0, v149)
			mBase = m.M
			v151 = m.ExcPending
			if v151 != 0 {
				return
			} else {
				return
			}
		}
	}
}
func F_getConfigDebugInfo(m *base.Module) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v57 int64
	_ = v57
	var v58 int64
	_ = v58
	var v59 int64
	_ = v59
	var v60 int64
	_ = v60
	var v61 int64
	_ = v61
	var v62 int64
	_ = v62
	var v63 int64
	_ = v63
	var v65 int64
	_ = v65
	var v67 int64
	_ = v67
	var v69 int64
	_ = v69
	var v71 int64
	_ = v71
	var v73 int64
	_ = v73
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v95 int32
	_ = v95
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v115 int32
	_ = v115
	var v120 int32
	_ = v120
	var v124 int32
	_ = v124
	var v128 int32
	_ = v128
	var v135 int32
	_ = v135
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v146 int32
	_ = v146
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v158 int32
	_ = v158
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v166 int32
	_ = v166
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v175 int32
	_ = v175
	var v177 int32
	_ = v177
	var v179 int64
	_ = v179
	var v180 int64
	_ = v180
	var v181 int64
	_ = v181
	var v182 int64
	_ = v182
	var v183 int64
	_ = v183
	var v184 int64
	_ = v184
	var v185 int64
	_ = v185
	var v187 int64
	_ = v187
	var v189 int64
	_ = v189
	var v191 int64
	_ = v191
	var v193 int64
	_ = v193
	var v195 int64
	_ = v195
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v202 int32
	_ = v202
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v209 int32
	_ = v209
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v217 int32
	_ = v217
	var v221 int32
	_ = v221
	var v223 int32
	_ = v223
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v237 int32
	_ = v237
	var v242 int32
	_ = v242
	var v246 int32
	_ = v246
	var v250 int32
	_ = v250
	var v258 int32
	_ = v258
	var v259 int32
	_ = v259
	var v260 int32
	_ = v260
	var v261 int32
	_ = v261
	var v262 int32
	_ = v262
	var v264 int32
	_ = v264
	var v265 int32
	_ = v265
	var v267 int32
	_ = v267
	var v268 int32
	_ = v268
	var v270 int32
	_ = v270
	var v272 int32
	_ = v272
	v5 = F_valkey_malloc(m, int32(24))
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
	v10 = F_dictCreate(m, int32(_a494))
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v5))) = v10
	v14 = F_dictCreate(m, int32(_a495))
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v5)+8)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v5)+4)) = v14
	*(*int64)(unsafe.Add(mBase, uint32(v5)+16)) = int64(4294967296)
	v22 = *(*int32)(unsafe.Add(mBase, _consts[230]))
	v23 = F_dictGetIterator(m, v22)
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		goto L1
	} else {
		goto L6
	}
L5:
	;
	F_dictReleaseIterator(m, v23)
	mBase = m.M
	v258 = m.ExcPending
	if v258 != 0 {
		goto L1
	} else {
		goto L67
	}
L6:
	;
	v32 = v23 + int32(20)
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v23)+16))
	if v33 != 0 {
		goto L9
	} else {
		goto L10
	}
L7:
	;
	if v128 == int32(0) {
		goto L5
	} else {
		goto L33
	}
L8:
	;
	v39 = v32
	v40 = v36
	goto L11
L9:
	;
	v36 = int32(1)
	goto L8
L10:
	;
	v36 = int32(0)
	goto L8
L11:
	;
	switch v40 {
	case 0:
		goto L16
	default:
		goto L15
	}
L13:
	;
	v40 = int32(0)
	goto L11
L14:
	;
	goto L7
L15:
	;
	v120 = *(*int32)(unsafe.Add(mBase, uint32(v39)))
	*(*int32)(unsafe.Add(mBase, uint32(v23)+16)) = v120
	if v120 == int32(0) {
		goto L13
	} else {
		goto L32
	}
L16:
	;
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v23)+4))
	if v44 != int32(-1) {
		v83 = v44
		goto L17
	} else {
		goto L18
	}
L17:
	;
	v84 = int32(1)
	v85 = v83 + v84
	*(*int32)(unsafe.Add(mBase, uint32(v23)+4)) = v85
	v87 = int32(0)
	v90 = *(*int32)(unsafe.Add(mBase, uint32(v23)))
	v91 = *(*int32)(unsafe.Add(mBase, uint32(v23)+8))
	v95 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v90+v91+int32(26)))))
	if v95 == int32(255) {
		goto L26
	} else {
		goto L27
	}
L18:
	;
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v23)+8))
	if v48 != 0 {
		v83 = int32(-1)
		goto L17
	} else {
		goto L19
	}
L19:
	;
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v23)))
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v23)+12))
	if v50 == int32(0) {
		goto L21
	} else {
		goto L22
	}
L20:
	;
	v77 = *(*int32)(unsafe.Add(mBase, uint32(v76)+20))
	if v77 != int32(-1) {
		goto L23
	} else {
		goto L24
	}
L21:
	;
	v57 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v49)+16)))
	v58 = int64(*(*int8)(unsafe.Add(mBase, uint32(v49)+27)))
	v59 = int64(*(*int32)(unsafe.Add(mBase, uint32(v49)+8)))
	v60 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v49)+12)))
	v61 = int64(*(*int8)(unsafe.Add(mBase, uint32(v49)+26)))
	v62 = int64(*(*int32)(unsafe.Add(mBase, uint32(v49)+4)))
	v63 = F_wangHash64(m, v62)
	mBase = m.M
	v65 = F_wangHash64(m, v61+v63)
	mBase = m.M
	v67 = F_wangHash64(m, v60+v65)
	mBase = m.M
	v69 = F_wangHash64(m, v59+v67)
	mBase = m.M
	v71 = F_wangHash64(m, v58+v69)
	mBase = m.M
	v73 = F_wangHash64(m, v57+v71)
	mBase = m.M
	*(*int64)(unsafe.Add(mBase, uint32(v23)+24)) = v73
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v23)))
	v76 = v75
	goto L20
L22:
	;
	v53 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v49)+24)))
	v55 = v53 + int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v49)+24)) = uint16(v55)
	v76 = v49
	goto L20
L23:
	;
	v83 = v77 + int32(-1)
	goto L17
L24:
	;
	v80 = *(*int32)(unsafe.Add(mBase, uint32(v23)+4))
	v83 = v80
	goto L17
L25:
	;
	v110 = int32(2)
	v115 = *(*int32)(unsafe.Add(mBase, uint32(v90+v108<<(uint(v110)%32)+int32(4))))
	v39 = v115 + v109<<(uint(v110)%32)
	v40 = int32(1)
	goto L11
L26:
	;
	v99 = v87
	goto L28
L27:
	;
	v99 = v84 << (uint(v95) % 32)
	goto L28
L28:
	;
	if v85 < v99 {
		v108 = v91
		v109 = v85
		goto L25
	} else {
		goto L29
	}
L29:
	;
	if v91 != 0 {
		v128 = v87
		goto L14
	} else {
		goto L30
	}
L30:
	;
	v101 = *(*int32)(unsafe.Add(mBase, uint32(v90)+20))
	if v101 == int32(-1) {
		v128 = v87
		goto L14
	} else {
		goto L31
	}
L31:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v23)+4)) = int64(4294967296)
	v108 = int32(1)
	v109 = int32(0)
	goto L25
L32:
	;
	v124 = *(*int32)(unsafe.Add(mBase, uint32(v120)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v32))) = v124
	v128 = v120
	goto L14
L33:
	;
	v135 = v128
	goto L34
L34:
	;
	v137 = *(*int32)(unsafe.Add(mBase, uint32(v135)+8))
	goto L37
L35:
	;
	goto L5
L36:
	;
	v154 = v23 + int32(20)
	v155 = *(*int32)(unsafe.Add(mBase, uint32(v23)+16))
	if v155 != 0 {
		goto L42
	} else {
		goto L43
	}
L37:
	;
	v138 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v137)+8)))
	if v138&int32(4) == int32(0) {
		goto L36
	} else {
		goto L38
	}
L38:
	;
	v143 = *(*int32)(unsafe.Add(mBase, uint32(v137)))
	v144 = *(*int32)(unsafe.Add(mBase, uint32(v137)+28))
	m.T0[v144].(func(*base.Module, int32, int32, int32))(m, v137, v143, v5)
	mBase = m.M
	v146 = m.ExcPending
	if v146 != 0 {
		goto L1
	} else {
		goto L39
	}
L39:
	;
	goto L36
L40:
	;
	if v250 != 0 {
		v135 = v250
		goto L34
	} else {
		goto L66
	}
L41:
	;
	v161 = v154
	v162 = v158
	goto L44
L42:
	;
	v158 = int32(1)
	goto L41
L43:
	;
	v158 = int32(0)
	goto L41
L44:
	;
	switch v162 {
	case 0:
		goto L49
	default:
		goto L48
	}
L46:
	;
	v162 = int32(0)
	goto L44
L47:
	;
	goto L40
L48:
	;
	v242 = *(*int32)(unsafe.Add(mBase, uint32(v161)))
	*(*int32)(unsafe.Add(mBase, uint32(v23)+16)) = v242
	if v242 == int32(0) {
		goto L46
	} else {
		goto L65
	}
L49:
	;
	v166 = *(*int32)(unsafe.Add(mBase, uint32(v23)+4))
	if v166 != int32(-1) {
		v205 = v166
		goto L50
	} else {
		goto L51
	}
L50:
	;
	v206 = int32(1)
	v207 = v205 + v206
	*(*int32)(unsafe.Add(mBase, uint32(v23)+4)) = v207
	v209 = int32(0)
	v212 = *(*int32)(unsafe.Add(mBase, uint32(v23)))
	v213 = *(*int32)(unsafe.Add(mBase, uint32(v23)+8))
	v217 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v212+v213+int32(26)))))
	if v217 == int32(255) {
		goto L59
	} else {
		goto L60
	}
L51:
	;
	v170 = *(*int32)(unsafe.Add(mBase, uint32(v23)+8))
	if v170 != 0 {
		v205 = int32(-1)
		goto L50
	} else {
		goto L52
	}
L52:
	;
	v171 = *(*int32)(unsafe.Add(mBase, uint32(v23)))
	v172 = *(*int32)(unsafe.Add(mBase, uint32(v23)+12))
	if v172 == int32(0) {
		goto L54
	} else {
		goto L55
	}
L53:
	;
	v199 = *(*int32)(unsafe.Add(mBase, uint32(v198)+20))
	if v199 != int32(-1) {
		goto L56
	} else {
		goto L57
	}
L54:
	;
	v179 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v171)+16)))
	v180 = int64(*(*int8)(unsafe.Add(mBase, uint32(v171)+27)))
	v181 = int64(*(*int32)(unsafe.Add(mBase, uint32(v171)+8)))
	v182 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v171)+12)))
	v183 = int64(*(*int8)(unsafe.Add(mBase, uint32(v171)+26)))
	v184 = int64(*(*int32)(unsafe.Add(mBase, uint32(v171)+4)))
	v185 = F_wangHash64(m, v184)
	mBase = m.M
	v187 = F_wangHash64(m, v183+v185)
	mBase = m.M
	v189 = F_wangHash64(m, v182+v187)
	mBase = m.M
	v191 = F_wangHash64(m, v181+v189)
	mBase = m.M
	v193 = F_wangHash64(m, v180+v191)
	mBase = m.M
	v195 = F_wangHash64(m, v179+v193)
	mBase = m.M
	*(*int64)(unsafe.Add(mBase, uint32(v23)+24)) = v195
	v197 = *(*int32)(unsafe.Add(mBase, uint32(v23)))
	v198 = v197
	goto L53
L55:
	;
	v175 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v171)+24)))
	v177 = v175 + int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v171)+24)) = uint16(v177)
	v198 = v171
	goto L53
L56:
	;
	v205 = v199 + int32(-1)
	goto L50
L57:
	;
	v202 = *(*int32)(unsafe.Add(mBase, uint32(v23)+4))
	v205 = v202
	goto L50
L58:
	;
	v232 = int32(2)
	v237 = *(*int32)(unsafe.Add(mBase, uint32(v212+v230<<(uint(v232)%32)+int32(4))))
	v161 = v237 + v231<<(uint(v232)%32)
	v162 = int32(1)
	goto L44
L59:
	;
	v221 = v209
	goto L61
L60:
	;
	v221 = v206 << (uint(v217) % 32)
	goto L61
L61:
	;
	if v207 < v221 {
		v230 = v213
		v231 = v207
		goto L58
	} else {
		goto L62
	}
L62:
	;
	if v213 != 0 {
		v250 = v209
		goto L47
	} else {
		goto L63
	}
L63:
	;
	v223 = *(*int32)(unsafe.Add(mBase, uint32(v212)+20))
	if v223 == int32(-1) {
		v250 = v209
		goto L47
	} else {
		goto L64
	}
L64:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v23)+4)) = int64(4294967296)
	v230 = int32(1)
	v231 = int32(0)
	goto L58
L65:
	;
	v246 = *(*int32)(unsafe.Add(mBase, uint32(v242)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v154))) = v246
	v250 = v242
	goto L47
L66:
	;
	goto L35
L67:
	;
	v259 = F_rewriteConfigGetContentFromState(m, v5)
	mBase = m.M
	v260 = m.ExcPending
	if v260 != 0 {
		goto L1
	} else {
		goto L68
	}
L68:
	;
	v261 = *(*int32)(unsafe.Add(mBase, uint32(v5)+12))
	v262 = *(*int32)(unsafe.Add(mBase, uint32(v5)+8))
	F_sdsfreesplitres(m, v261, v262)
	mBase = m.M
	v264 = m.ExcPending
	if v264 != 0 {
		goto L1
	} else {
		goto L69
	}
L69:
	;
	v265 = *(*int32)(unsafe.Add(mBase, uint32(v5)))
	F_dictRelease(m, v265)
	mBase = m.M
	v267 = m.ExcPending
	if v267 != 0 {
		goto L1
	} else {
		goto L70
	}
L70:
	;
	v268 = *(*int32)(unsafe.Add(mBase, uint32(v5)+4))
	F_dictRelease(m, v268)
	mBase = m.M
	v270 = m.ExcPending
	if v270 != 0 {
		goto L1
	} else {
		goto L71
	}
L71:
	;
	F_valkey_free(m, v5)
	mBase = m.M
	v272 = m.ExcPending
	if v272 != 0 {
		goto L1
	} else {
		goto L72
	}
L72:
	;
	return v259
}
func F_getConfigNotifyKeyspaceEventsOption(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	v3 = *(*int32)(unsafe.Add(mBase, _consts[289]))
	v4 = F_keyspaceEventsFlagsToString(m, v3)
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		return v4
	}
}
func F_getConfigOOMScoreAdjValuesOption(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	v4 = m.G0
	v6 = v4 - int32(48)
	m.G0 = v6
	v8 = F_sdsempty(m)
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return int32(0)
	} else {
		v13 = *(*int32)(unsafe.Add(mBase, _consts[255]))
		*(*int32)(unsafe.Add(mBase, uint32(v6)+32)) = v13
		v18 = F_sdscatprintf(m, v8, int32(_a77), v6+int32(32))
		mBase = m.M
		v19 = m.ExcPending
		if v19 != 0 {
			return int32(0)
		} else {
			v22 = F_sdscatlen(m, v18, int32(_a66), int32(1))
			mBase = m.M
			v23 = m.ExcPending
			if v23 != 0 {
				return int32(0)
			} else {
				v25 = *(*int32)(unsafe.Add(mBase, _consts[257]))
				*(*int32)(unsafe.Add(mBase, uint32(v6)+16)) = v25
				v30 = F_sdscatprintf(m, v22, int32(_a77), v6+int32(16))
				mBase = m.M
				v31 = m.ExcPending
				if v31 != 0 {
					return int32(0)
				} else {
					v34 = F_sdscatlen(m, v30, int32(_a66), int32(1))
					mBase = m.M
					v35 = m.ExcPending
					if v35 != 0 {
						return int32(0)
					} else {
						v37 = *(*int32)(unsafe.Add(mBase, _consts[259]))
						*(*int32)(unsafe.Add(mBase, uint32(v6))) = v37
						v40 = F_sdscatprintf(m, v34, int32(_a77), v6)
						mBase = m.M
						v41 = m.ExcPending
						if v41 != 0 {
							return int32(0)
						} else {
							m.G0 = v6 + int32(48)
							return v40
						}
					}
				}
			}
		}
	}
}
func F_getConfigRdmaBindOption(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	v6 = *(*int32)(unsafe.Add(mBase, _consts[276]))
	v8 = F_sdsjoin(m, int32(_a512), v6, int32(_a66))
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return int32(0)
	} else {
		return v8
	}
}
func F_rewriteConfigClientOutputBufferLimitOption(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v18 int64
	_ = v18
	var v20 int64
	_ = v20
	var v24 int64
	_ = v24
	var v26 int64
	_ = v26
	var v29 int64
	_ = v29
	var v31 int64
	_ = v31
	var v33 int32
	_ = v33
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v43 int64
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v96 int64
	_ = v96
	var v100 int32
	_ = v100
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v116 int64
	_ = v116
	var v118 int64
	_ = v118
	var v121 int64
	_ = v121
	var v123 int64
	_ = v123
	var v126 int64
	_ = v126
	var v128 int64
	_ = v128
	var v130 int32
	_ = v130
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v140 int64
	_ = v140
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v153 int32
	_ = v153
	var v155 int32
	_ = v155
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v173 int32
	_ = v173
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v194 int64
	_ = v194
	var v198 int32
	_ = v198
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v214 int64
	_ = v214
	var v216 int64
	_ = v216
	var v219 int64
	_ = v219
	var v221 int64
	_ = v221
	var v224 int64
	_ = v224
	var v226 int64
	_ = v226
	var v228 int32
	_ = v228
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v238 int64
	_ = v238
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v250 int32
	_ = v250
	var v252 int32
	_ = v252
	var v255 int32
	_ = v255
	var v256 int32
	_ = v256
	var v262 int32
	_ = v262
	var v263 int32
	_ = v263
	var v266 int32
	_ = v266
	var v267 int32
	_ = v267
	var v270 int32
	_ = v270
	var v279 int32
	_ = v279
	var v280 int32
	_ = v280
	var v284 int32
	_ = v284
	var v285 int32
	_ = v285
	var v289 int64
	_ = v289
	var v293 int32
	_ = v293
	var v302 int32
	_ = v302
	var v303 int32
	_ = v303
	var v304 int32
	_ = v304
	var v305 int32
	_ = v305
	v11 = m.G0
	v13 = v11 - int32(224)
	m.G0 = v13
	v15 = int32(1)
	v18 = *(*int64)(unsafe.Add(mBase, _consts[233]))
	v20 = *(*int64)(unsafe.Add(mBase, _consts[234]))
	if v18 != v20 {
		v33 = v15
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v37 = F_rewriteConfigFormatMemory(m, v13+int32(160), int32(64), v18)
	mBase = m.M
	v38 = m.ExcPending
	if v38 != 0 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	v24 = *(*int64)(unsafe.Add(mBase, _consts[235]))
	v26 = *(*int64)(unsafe.Add(mBase, _consts[236]))
	if v24 != v26 {
		v33 = int32(1)
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v29 = *(*int64)(unsafe.Add(mBase, _consts[237]))
	v31 = *(*int64)(unsafe.Add(mBase, _consts[238]))
	v33 = base.B2i32(v29 != v31)
	goto L1
L4:
	;
	return
L5:
	;
	v43 = *(*int64)(unsafe.Add(mBase, _consts[235]))
	v44 = F_rewriteConfigFormatMemory(m, v13+int32(96), int32(64), v43)
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L4
	} else {
		goto L6
	}
L6:
	;
	goto L9
L7:
	;
	v57 = int32(_a505)
	v60 = int32(*(*uint8)(unsafe.Add(mBase, _consts[239])))
	v61 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v55))))
	if v61 == int32(0) {
		v84 = v60
		v85 = v61
		goto L11
	} else {
		goto L12
	}
L8:
	;
	goto L7
L9:
	;
	v55 = *(*int32)(unsafe.Add(mBase, _consts[240]))
	goto L8
L10:
	;
	v89 = F_sdsempty(m)
	mBase = m.M
	v90 = m.ExcPending
	if v90 != 0 {
		goto L4
	} else {
		goto L18
	}
L11:
	;
	goto L10
L12:
	;
	if v61 != v60&int32(255) {
		v84 = v60
		v85 = v61
		goto L11
	} else {
		goto L13
	}
L13:
	;
	v67 = v55
	v68 = v57
	goto L14
L14:
	;
	v71 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v68)+1)))
	v72 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v67)+1)))
	if v72 == int32(0) {
		v84 = v71
		v85 = v72
		goto L11
	} else {
		goto L16
	}
L15:
	;
	v84 = v71
	v85 = v72
	goto L11
L16:
	;
	v75 = int32(1)
	if v72 == v71&int32(255) {
		v67 = v67 + v75
		v68 = v68 + v75
		goto L14
	} else {
		goto L17
	}
L17:
	;
	goto L15
L18:
	;
	v96 = *(*int64)(unsafe.Add(mBase, _consts[237]))
	*(*uint32)(unsafe.Add(mBase, uint32(v13+int32(80)))) = uint32(v96)
	*(*int32)(unsafe.Add(mBase, uint32(v13)+64)) = l1
	if v85-v84&int32(255) != 0 {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	v100 = v55
	goto L21
L20:
	;
	v100 = int32(_a268)
	goto L21
L21:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+68)) = v100
	*(*int32)(unsafe.Add(mBase, uint32(v13)+72)) = v13 + int32(160)
	*(*int32)(unsafe.Add(mBase, uint32(v13)+76)) = v13 + int32(96)
	v111 = F_sdscatprintf(m, v89, int32(_a506), v13+int32(64))
	mBase = m.M
	v112 = m.ExcPending
	if v112 != 0 {
		goto L4
	} else {
		goto L22
	}
L22:
	;
	v113 = F_rewriteConfigRewriteLine(m, l2, l1, v111, v33)
	mBase = m.M
	v114 = m.ExcPending
	if v114 != 0 {
		goto L4
	} else {
		goto L23
	}
L23:
	;
	v116 = *(*int64)(unsafe.Add(mBase, _consts[241]))
	v118 = *(*int64)(unsafe.Add(mBase, _consts[242]))
	if v116 != v118 {
		v130 = v15
		goto L24
	} else {
		goto L25
	}
L24:
	;
	v134 = F_rewriteConfigFormatMemory(m, v13+int32(160), int32(64), v116)
	mBase = m.M
	v135 = m.ExcPending
	if v135 != 0 {
		goto L4
	} else {
		goto L27
	}
L25:
	;
	v121 = *(*int64)(unsafe.Add(mBase, _consts[243]))
	v123 = *(*int64)(unsafe.Add(mBase, _consts[244]))
	if v121 != v123 {
		v130 = v15
		goto L24
	} else {
		goto L26
	}
L26:
	;
	v126 = *(*int64)(unsafe.Add(mBase, _consts[245]))
	v128 = *(*int64)(unsafe.Add(mBase, _consts[246]))
	v130 = base.B2i32(v126 != v128)
	goto L24
L27:
	;
	v140 = *(*int64)(unsafe.Add(mBase, _consts[243]))
	v141 = F_rewriteConfigFormatMemory(m, v13+int32(96), int32(64), v140)
	mBase = m.M
	v142 = m.ExcPending
	if v142 != 0 {
		goto L4
	} else {
		goto L28
	}
L28:
	;
	v143 = int32(1)
	goto L31
L29:
	;
	v155 = int32(_a505)
	v158 = int32(*(*uint8)(unsafe.Add(mBase, _consts[239])))
	v159 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v153))))
	if v159 == int32(0) {
		v182 = v158
		v183 = v159
		goto L33
	} else {
		goto L34
	}
L30:
	;
	goto L29
L31:
	;
	v153 = *(*int32)(unsafe.Add(mBase, _consts[247]))
	goto L30
L32:
	;
	v187 = F_sdsempty(m)
	mBase = m.M
	v188 = m.ExcPending
	if v188 != 0 {
		goto L4
	} else {
		goto L40
	}
L33:
	;
	goto L32
L34:
	;
	if v159 != v158&int32(255) {
		v182 = v158
		v183 = v159
		goto L33
	} else {
		goto L35
	}
L35:
	;
	v165 = v153
	v166 = v155
	goto L36
L36:
	;
	v169 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v166)+1)))
	v170 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v165)+1)))
	if v170 == int32(0) {
		v182 = v169
		v183 = v170
		goto L33
	} else {
		goto L38
	}
L37:
	;
	v182 = v169
	v183 = v170
	goto L33
L38:
	;
	v173 = int32(1)
	if v170 == v169&int32(255) {
		v165 = v165 + v173
		v166 = v166 + v173
		goto L36
	} else {
		goto L39
	}
L39:
	;
	goto L37
L40:
	;
	v194 = *(*int64)(unsafe.Add(mBase, _consts[245]))
	*(*uint32)(unsafe.Add(mBase, uint32(v13+int32(48)))) = uint32(v194)
	*(*int32)(unsafe.Add(mBase, uint32(v13)+32)) = l1
	if v183-v182&int32(255) != 0 {
		goto L41
	} else {
		goto L42
	}
L41:
	;
	v198 = v153
	goto L43
L42:
	;
	v198 = int32(_a268)
	goto L43
L43:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+36)) = v198
	*(*int32)(unsafe.Add(mBase, uint32(v13)+40)) = v13 + int32(160)
	*(*int32)(unsafe.Add(mBase, uint32(v13)+44)) = v13 + int32(96)
	v209 = F_sdscatprintf(m, v187, int32(_a506), v13+int32(32))
	mBase = m.M
	v210 = m.ExcPending
	if v210 != 0 {
		goto L4
	} else {
		goto L44
	}
L44:
	;
	v211 = F_rewriteConfigRewriteLine(m, l2, l1, v209, v130)
	mBase = m.M
	v212 = m.ExcPending
	if v212 != 0 {
		goto L4
	} else {
		goto L45
	}
L45:
	;
	v214 = *(*int64)(unsafe.Add(mBase, _consts[248]))
	v216 = *(*int64)(unsafe.Add(mBase, _consts[249]))
	if v214 != v216 {
		v228 = v143
		goto L46
	} else {
		goto L47
	}
L46:
	;
	v232 = F_rewriteConfigFormatMemory(m, v13+int32(160), int32(64), v214)
	mBase = m.M
	v233 = m.ExcPending
	if v233 != 0 {
		goto L4
	} else {
		goto L49
	}
L47:
	;
	v219 = *(*int64)(unsafe.Add(mBase, _consts[250]))
	v221 = *(*int64)(unsafe.Add(mBase, _consts[251]))
	if v219 != v221 {
		v228 = v143
		goto L46
	} else {
		goto L48
	}
L48:
	;
	v224 = *(*int64)(unsafe.Add(mBase, _consts[252]))
	v226 = *(*int64)(unsafe.Add(mBase, _consts[253]))
	v228 = base.B2i32(v224 != v226)
	goto L46
L49:
	;
	v238 = *(*int64)(unsafe.Add(mBase, _consts[250]))
	v239 = F_rewriteConfigFormatMemory(m, v13+int32(96), int32(64), v238)
	mBase = m.M
	v240 = m.ExcPending
	if v240 != 0 {
		goto L4
	} else {
		goto L50
	}
L50:
	;
	goto L53
L51:
	;
	v252 = int32(_a505)
	v255 = int32(*(*uint8)(unsafe.Add(mBase, _consts[239])))
	v256 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v250))))
	if v256 == int32(0) {
		v279 = v255
		v280 = v256
		goto L55
	} else {
		goto L56
	}
L52:
	;
	goto L51
L53:
	;
	v250 = *(*int32)(unsafe.Add(mBase, _consts[254]))
	goto L52
L54:
	;
	v284 = F_sdsempty(m)
	mBase = m.M
	v285 = m.ExcPending
	if v285 != 0 {
		goto L4
	} else {
		goto L62
	}
L55:
	;
	goto L54
L56:
	;
	if v256 != v255&int32(255) {
		v279 = v255
		v280 = v256
		goto L55
	} else {
		goto L57
	}
L57:
	;
	v262 = v250
	v263 = v252
	goto L58
L58:
	;
	v266 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v263)+1)))
	v267 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v262)+1)))
	if v267 == int32(0) {
		v279 = v266
		v280 = v267
		goto L55
	} else {
		goto L60
	}
L59:
	;
	v279 = v266
	v280 = v267
	goto L55
L60:
	;
	v270 = int32(1)
	if v267 == v266&int32(255) {
		v262 = v262 + v270
		v263 = v263 + v270
		goto L58
	} else {
		goto L61
	}
L61:
	;
	goto L59
L62:
	;
	v289 = *(*int64)(unsafe.Add(mBase, _consts[252]))
	*(*uint32)(unsafe.Add(mBase, uint32(v13+int32(16)))) = uint32(v289)
	*(*int32)(unsafe.Add(mBase, uint32(v13))) = l1
	if v280-v279&int32(255) != 0 {
		goto L63
	} else {
		goto L64
	}
L63:
	;
	v293 = v250
	goto L65
L64:
	;
	v293 = int32(_a268)
	goto L65
L65:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+4)) = v293
	*(*int32)(unsafe.Add(mBase, uint32(v13)+8)) = v13 + int32(160)
	*(*int32)(unsafe.Add(mBase, uint32(v13)+12)) = v13 + int32(96)
	v302 = F_sdscatprintf(m, v284, int32(_a506), v13)
	mBase = m.M
	v303 = m.ExcPending
	if v303 != 0 {
		goto L4
	} else {
		goto L66
	}
L66:
	;
	v304 = F_rewriteConfigRewriteLine(m, l2, l1, v302, v228)
	mBase = m.M
	v305 = m.ExcPending
	if v305 != 0 {
		goto L4
	} else {
		goto L67
	}
L67:
	;
	m.G0 = v13 + int32(224)
	return
}
func F_rewriteConfigMarkAsProcessed(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v12 int32
	_ = v12
	v3 = F_sdsnew(m, l1)
	mBase = m.M
	v4 = m.ExcPending
	if v4 != 0 {
		return
	} else {
		v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		v7 = F_dictAdd(m, v5, v3, int32(0))
		mBase = m.M
		v8 = m.ExcPending
		if v8 != 0 {
			return
		} else {
			if v7 == int32(0) {
				return
			} else {
				F_sdsfree(m, v3)
				mBase = m.M
				v12 = m.ExcPending
				if v12 != 0 {
					return
				} else {
					return
				}
			}
		}
	}
}
func F_rewriteConfigOOMScoreAdjValuesOption(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	v12 = m.G0
	v14 = v12 - int32(48)
	m.G0 = v14
	v16 = F_sdsnew(m, l1)
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		return
	} else {
		v20 = F_sdscatlen(m, v16, int32(_a66), int32(1))
		mBase = m.M
		v21 = m.ExcPending
		if v21 != 0 {
			return
		} else {
			v23 = *(*int32)(unsafe.Add(mBase, _consts[255]))
			*(*int32)(unsafe.Add(mBase, uint32(v14)+32)) = v23
			v26 = *(*int32)(unsafe.Add(mBase, _consts[256]))
			v30 = F_sdscatprintf(m, v20, int32(_a77), v14+int32(32))
			mBase = m.M
			v31 = m.ExcPending
			if v31 != 0 {
				return
			} else {
				v34 = F_sdscatlen(m, v30, int32(_a66), int32(1))
				mBase = m.M
				v35 = m.ExcPending
				if v35 != 0 {
					return
				} else {
					v37 = *(*int32)(unsafe.Add(mBase, _consts[257]))
					*(*int32)(unsafe.Add(mBase, uint32(v14)+16)) = v37
					v40 = *(*int32)(unsafe.Add(mBase, _consts[258]))
					v44 = F_sdscatprintf(m, v34, int32(_a77), v14+int32(16))
					mBase = m.M
					v45 = m.ExcPending
					if v45 != 0 {
						return
					} else {
						v48 = F_sdscatlen(m, v44, int32(_a66), int32(1))
						mBase = m.M
						v49 = m.ExcPending
						if v49 != 0 {
							return
						} else {
							v51 = *(*int32)(unsafe.Add(mBase, _consts[259]))
							*(*int32)(unsafe.Add(mBase, uint32(v14))) = v51
							v54 = *(*int32)(unsafe.Add(mBase, _consts[260]))
							v56 = F_sdscatprintf(m, v48, int32(_a77), v14)
							mBase = m.M
							v57 = m.ExcPending
							if v57 != 0 {
								return
							} else {
								v63 = F_rewriteConfigRewriteLine(m, l2, l1, v56, base.B2i32(v51 != v54)|base.B2i32(v37 != v40)|base.B2i32(v23 != v26))
								mBase = m.M
								v64 = m.ExcPending
								if v64 != 0 {
									return
								} else {
									m.G0 = v14 + int32(48)
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
func F_rewriteConfigReadOldFile(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v82 int64
	_ = v82
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
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
	var v98 int32
	_ = v98
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
	var v114 int32
	_ = v114
	var v121 int32
	_ = v121
	var v124 int32
	_ = v124
	var v127 int32
	_ = v127
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v142 int32
	_ = v142
	var v148 int32
	_ = v148
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v154 int32
	_ = v154
	var v156 int32
	_ = v156
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v178 int32
	_ = v178
	var v181 int32
	_ = v181
	var v184 int32
	_ = v184
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v191 int32
	_ = v191
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v203 int32
	_ = v203
	var v209 int32
	_ = v209
	var v213 int32
	_ = v213
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v221 int32
	_ = v221
	var v225 int32
	_ = v225
	var v230 int32
	_ = v230
	var v232 int32
	_ = v232
	var v234 int32
	_ = v234
	var v236 int32
	_ = v236
	var v240 int32
	_ = v240
	var v258 int32
	_ = v258
	var v259 int32
	_ = v259
	var v263 int32
	_ = v263
	var v266 int32
	_ = v266
	var v267 int32
	_ = v267
	var v273 int32
	_ = v273
	var v274 int32
	_ = v274
	var v277 int32
	_ = v277
	var v278 int32
	_ = v278
	var v281 int32
	_ = v281
	var v290 int32
	_ = v290
	var v291 int32
	_ = v291
	var v295 int32
	_ = v295
	var v298 int32
	_ = v298
	var v303 int32
	_ = v303
	var v304 int32
	_ = v304
	var v306 int32
	_ = v306
	var v308 int32
	_ = v308
	var v316 int32
	_ = v316
	var v317 int32
	_ = v317
	var v321 int32
	_ = v321
	var v322 int32
	_ = v322
	var v323 int32
	_ = v323
	var v324 int32
	_ = v324
	var v327 int32
	_ = v327
	var v330 int32
	_ = v330
	var v331 int32
	_ = v331
	var v332 int32
	_ = v332
	var v335 int32
	_ = v335
	var v337 int32
	_ = v337
	var v338 int32
	_ = v338
	var v339 int32
	_ = v339
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
	var v351 int32
	_ = v351
	var v352 int32
	_ = v352
	var v353 int32
	_ = v353
	var v354 int32
	_ = v354
	var v359 int32
	_ = v359
	var v360 int32
	_ = v360
	var v363 int32
	_ = v363
	var v364 int32
	_ = v364
	var v367 int32
	_ = v367
	var v368 int32
	_ = v368
	var v369 int32
	_ = v369
	var v373 int32
	_ = v373
	var v376 int32
	_ = v376
	var v378 int32
	_ = v378
	var v379 int32
	_ = v379
	var v380 int32
	_ = v380
	var v382 int32
	_ = v382
	var v386 int32
	_ = v386
	var v388 int32
	_ = v388
	var v389 int32
	_ = v389
	var v390 int32
	_ = v390
	var v392 int32
	_ = v392
	var v393 int32
	_ = v393
	var v394 int32
	_ = v394
	var v395 int32
	_ = v395
	var v400 int32
	_ = v400
	var v401 int32
	_ = v401
	var v404 int32
	_ = v404
	var v405 int32
	_ = v405
	var v408 int32
	_ = v408
	var v409 int32
	_ = v409
	var v410 int32
	_ = v410
	var v414 int32
	_ = v414
	var v417 int32
	_ = v417
	var v419 int32
	_ = v419
	var v420 int32
	_ = v420
	var v421 int32
	_ = v421
	var v423 int32
	_ = v423
	var v427 int32
	_ = v427
	var v429 int32
	_ = v429
	var v430 int32
	_ = v430
	var v431 int32
	_ = v431
	var v433 int32
	_ = v433
	var v434 int32
	_ = v434
	var v435 int32
	_ = v435
	var v436 int32
	_ = v436
	var v441 int32
	_ = v441
	var v442 int32
	_ = v442
	var v445 int32
	_ = v445
	var v446 int32
	_ = v446
	var v449 int32
	_ = v449
	var v450 int32
	_ = v450
	var v451 int32
	_ = v451
	var v455 int32
	_ = v455
	var v458 int32
	_ = v458
	var v460 int32
	_ = v460
	var v461 int32
	_ = v461
	var v462 int32
	_ = v462
	var v464 int32
	_ = v464
	var v468 int32
	_ = v468
	var v470 int32
	_ = v470
	var v471 int32
	_ = v471
	var v472 int32
	_ = v472
	var v474 int32
	_ = v474
	var v475 int32
	_ = v475
	var v476 int32
	_ = v476
	var v477 int32
	_ = v477
	var v482 int32
	_ = v482
	var v483 int32
	_ = v483
	var v486 int32
	_ = v486
	var v487 int32
	_ = v487
	var v490 int32
	_ = v490
	var v491 int32
	_ = v491
	var v492 int32
	_ = v492
	var v496 int32
	_ = v496
	var v499 int32
	_ = v499
	var v501 int32
	_ = v501
	var v502 int32
	_ = v502
	var v503 int32
	_ = v503
	var v505 int32
	_ = v505
	var v509 int32
	_ = v509
	var v511 int32
	_ = v511
	var v512 int32
	_ = v512
	var v513 int32
	_ = v513
	var v515 int32
	_ = v515
	var v516 int32
	_ = v516
	var v517 int32
	_ = v517
	var v518 int32
	_ = v518
	var v523 int32
	_ = v523
	var v524 int32
	_ = v524
	var v527 int32
	_ = v527
	var v528 int32
	_ = v528
	var v531 int32
	_ = v531
	var v532 int32
	_ = v532
	var v533 int32
	_ = v533
	var v539 int32
	_ = v539
	var v540 int32
	_ = v540
	var v541 int32
	_ = v541
	var v542 int32
	_ = v542
	var v545 int32
	_ = v545
	var v547 int32
	_ = v547
	var v549 int32
	_ = v549
	var v554 int32
	_ = v554
	var v555 int32
	_ = v555
	var v557 int32
	_ = v557
	var v559 int32
	_ = v559
	var v565 int32
	_ = v565
	var v571 int32
	_ = v571
	var v578 int32
	_ = v578
	var v581 int32
	_ = v581
	var v584 int32
	_ = v584
	var v587 int32
	_ = v587
	var v588 int32
	_ = v588
	var v593 int32
	_ = v593
	var v596 int32
	_ = v596
	var v597 int32
	_ = v597
	var v598 int32
	_ = v598
	var v601 int32
	_ = v601
	var v611 int32
	_ = v611
	var v612 int32
	_ = v612
	var v614 int32
	_ = v614
	var v616 int32
	_ = v616
	var v623 int32
	_ = v623
	var v624 int32
	_ = v624
	var v625 int32
	_ = v625
	var v626 int32
	_ = v626
	var v629 int32
	_ = v629
	var v632 int32
	_ = v632
	var v637 int32
	_ = v637
	var v639 int32
	_ = v639
	var v640 int32
	_ = v640
	var v641 int32
	_ = v641
	var v642 int32
	_ = v642
	var v645 int32
	_ = v645
	var v647 int32
	_ = v647
	var v650 int32
	_ = v650
	var v653 int32
	_ = v653
	var v656 int32
	_ = v656
	var v658 int32
	_ = v658
	var v659 int32
	_ = v659
	var v660 int32
	_ = v660
	var v662 int32
	_ = v662
	var v666 int32
	_ = v666
	var v668 int32
	_ = v668
	var v669 int32
	_ = v669
	var v670 int32
	_ = v670
	var v672 int32
	_ = v672
	var v673 int32
	_ = v673
	var v674 int32
	_ = v674
	var v675 int32
	_ = v675
	var v680 int32
	_ = v680
	var v681 int32
	_ = v681
	var v684 int32
	_ = v684
	var v685 int32
	_ = v685
	var v688 int32
	_ = v688
	var v689 int32
	_ = v689
	var v690 int32
	_ = v690
	var v692 int32
	_ = v692
	var v693 int32
	_ = v693
	var v694 int64
	_ = v694
	var v697 int32
	_ = v697
	var v698 int32
	_ = v698
	var v699 int32
	_ = v699
	var v700 int32
	_ = v700
	var v701 int32
	_ = v701
	var v702 int32
	_ = v702
	var v703 int32
	_ = v703
	var v704 int32
	_ = v704
	var v705 int32
	_ = v705
	var v706 int32
	_ = v706
	var v707 int32
	_ = v707
	var v708 int32
	_ = v708
	var v709 int32
	_ = v709
	var v710 int32
	_ = v710
	var v712 int32
	_ = v712
	var v713 int32
	_ = v713
	var v714 int32
	_ = v714
	var v715 int32
	_ = v715
	var v716 int32
	_ = v716
	var v717 int32
	_ = v717
	var v718 int32
	_ = v718
	var v719 int32
	_ = v719
	var v720 int32
	_ = v720
	var v721 int32
	_ = v721
	var v722 int32
	_ = v722
	var v723 int32
	_ = v723
	var v724 int32
	_ = v724
	var v728 int32
	_ = v728
	var v730 int32
	_ = v730
	var v731 int32
	_ = v731
	var v733 int32
	_ = v733
	var v734 int32
	_ = v734
	var v739 int32
	_ = v739
	var v740 int32
	_ = v740
	var v757 int32
	_ = v757
	var v758 int32
	_ = v758
	var v759 int32
	_ = v759
	var v761 int32
	_ = v761
	var v763 int32
	_ = v763
	var v764 int32
	_ = v764
	v16 = m.G0
	v18 = v16 - int32(112)
	m.G0 = v18
	v21 = F_fopen(m, l0, int32(_a86))
	mBase = m.M
	if v21 != 0 {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	m.G0 = v18 + int32(112)
	return v764
L2:
	;
	v66 = F_valkey_malloc(m, int32(24))
	mBase = m.M
	v67 = m.ExcPending
	if v67 != 0 {
		goto L18
	} else {
		goto L20
	}
L3:
	;
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v21)+76))
	if int32(-1) < v29 {
		goto L9
	} else {
		goto L10
	}
L4:
	;
	goto L5
L5:
	;
	v24 = *(*int32)(unsafe.Add(mBase, _consts[5]))
	if v24 != int32(44) {
		v764 = int32(0)
		goto L1
	} else {
		goto L6
	}
L6:
	;
	goto L2
L7:
	;
	if int32(-1) < v46 {
		goto L15
	} else {
		goto L16
	}
L8:
	;
	if int32(-1) < v38 {
		v46 = v38
		goto L12
	} else {
		goto L13
	}
L9:
	;
	v33 = F___lockfile(m, v21)
	mBase = m.M
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v21)+60))
	if v33 == int32(0) {
		v38 = v34
		goto L8
	} else {
		goto L11
	}
L10:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v21)+60))
	v38 = v32
	goto L8
L11:
	;
	F___unlockfile(m, v21)
	mBase = m.M
	v38 = v34
	goto L8
L12:
	;
	goto L7
L13:
	;
	v42 = F___errno_location(m)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v42))) = int32(8)
	v46 = int32(-1)
	goto L12
L14:
	;
	if v56 != int32(-1) {
		goto L2
	} else {
		goto L17
	}
L15:
	;
	v55 = F___fstatat(m, v46, int32(_a139), v18+int32(16), int32(4096))
	mBase = m.M
	v56 = v55
	goto L14
L16:
	;
	v52 = F___syscall_ret(m, int32(-8))
	mBase = m.M
	v56 = v52
	goto L14
L17:
	;
	v59 = F_fclose(m, v21)
	mBase = m.M
	v62 = m.ExcPending
	if v62 != 0 {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	return int32(0)
L19:
	;
	v764 = int32(0)
	goto L1
L20:
	;
	v69 = F_dictCreate(m, int32(_a494))
	mBase = m.M
	v70 = m.ExcPending
	if v70 != 0 {
		goto L18
	} else {
		goto L21
	}
L21:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v66))) = v69
	v73 = F_dictCreate(m, int32(_a495))
	mBase = m.M
	v74 = m.ExcPending
	if v74 != 0 {
		goto L18
	} else {
		goto L22
	}
L22:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v66)+16)) = int64(1)
	*(*int64)(unsafe.Add(mBase, uint32(v66)+8)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v66)+4)) = v73
	if v21 == int32(0) {
		v764 = v66
		goto L1
	} else {
		goto L23
	}
L23:
	;
	v82 = *(*int64)(unsafe.Add(mBase, uint32(v18)+40))
	if v82 != int64(0) {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	v87 = int32(0)
	v89 = *(*int32)(unsafe.Add(mBase, _consts[72]))
	v90 = base.I32_wrap_i64(v82)
	v91 = F_sdsnewlen(m, v89, v90)
	mBase = m.M
	v92 = m.ExcPending
	if v92 != 0 {
		goto L18
	} else {
		goto L28
	}
L25:
	;
	v85 = F_fclose(m, v21)
	mBase = m.M
	v86 = m.ExcPending
	if v86 != 0 {
		goto L18
	} else {
		goto L26
	}
L26:
	;
	v764 = v66
	goto L1
L27:
	;
	v111 = int32(-1)
	v114 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v91+v111))))
	switch v114 & int32(7) {
	case 0:
		goto L42
	case 1:
		goto L41
	case 2:
		goto L40
	case 3:
		goto L39
	case 4:
		goto L38
	default:
		v131 = v87
		goto L37
	}
L28:
	;
	v94 = F_fread(m, v91, int32(1), v90, v21)
	mBase = m.M
	v95 = m.ExcPending
	if v95 != 0 {
		goto L18
	} else {
		goto L29
	}
L29:
	;
	if v94 != 0 {
		goto L27
	} else {
		goto L30
	}
L30:
	;
	F_sdsfree(m, v91)
	mBase = m.M
	v97 = m.ExcPending
	if v97 != 0 {
		goto L18
	} else {
		goto L31
	}
L31:
	;
	v98 = int32(0)
	F_sdsfreesplitres(m, v98, v98)
	mBase = m.M
	v101 = m.ExcPending
	if v101 != 0 {
		goto L18
	} else {
		goto L32
	}
L32:
	;
	F_dictRelease(m, v69)
	mBase = m.M
	v103 = m.ExcPending
	if v103 != 0 {
		goto L18
	} else {
		goto L33
	}
L33:
	;
	F_dictRelease(m, v73)
	mBase = m.M
	v105 = m.ExcPending
	if v105 != 0 {
		goto L18
	} else {
		goto L34
	}
L34:
	;
	F_valkey_free(m, v66)
	mBase = m.M
	v107 = m.ExcPending
	if v107 != 0 {
		goto L18
	} else {
		goto L35
	}
L35:
	;
	v108 = F_fclose(m, v21)
	mBase = m.M
	v109 = m.ExcPending
	if v109 != 0 {
		goto L18
	} else {
		goto L36
	}
L36:
	;
	v764 = int32(0)
	goto L1
L37:
	;
	v132 = int32(1)
	v137 = F_sdssplitlen(m, v91, v131, int32(_a247), v132, v18+int32(12))
	mBase = m.M
	v138 = m.ExcPending
	if v138 != 0 {
		goto L18
	} else {
		goto L43
	}
L38:
	;
	v130 = *(*int32)(unsafe.Add(mBase, uint32(v91+int32(-17))))
	v131 = v130
	goto L37
L39:
	;
	v127 = *(*int32)(unsafe.Add(mBase, uint32(v91+int32(-9))))
	v131 = v127
	goto L37
L40:
	;
	v124 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v91+int32(-5)))))
	v131 = v124
	goto L37
L41:
	;
	v121 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v91+int32(-3)))))
	v131 = v121
	goto L37
L42:
	;
	v131 = int32(base.Ui32(v114) >> (uint(int32(3)) % 32))
	goto L37
L43:
	;
	v139 = *(*int32)(unsafe.Add(mBase, uint32(v18)+12))
	if v139 < int32(1) {
		goto L44
	} else {
		goto L45
	}
L44:
	;
	v757 = F_fclose(m, v21)
	mBase = m.M
	v758 = m.ExcPending
	if v758 != 0 {
		goto L18
	} else {
		goto L231
	}
L45:
	;
	v142 = int32(0)
	v148 = v142
	v151 = v142
	v152 = v111
	v154 = v132
	v156 = v142
	goto L46
L46:
	;
	v162 = v137 + v151<<(uint(int32(2))%32)
	v163 = *(*int32)(unsafe.Add(mBase, uint32(v162)))
	v164 = int32(_a496)
	v170 = v163 + int32(-1)
	v171 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v170))))
	switch v171 & int32(7) {
	case 0:
		goto L54
	case 1:
		goto L53
	case 2:
		goto L52
	case 3:
		goto L51
	case 4:
		goto L50
	default:
		v188 = int32(0)
		goto L49
	}
L47:
	;
	goto L44
L48:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v162))) = int32(0)
	v258 = v152 + int32(1)
	v259 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v163))))
	if v259 == int32(35) {
		goto L77
	} else {
		goto L78
	}
L49:
	;
	v191 = v163 + v188 + int32(-1)
	if base.Ui32(v191) < base.Ui32(v163) {
		v209 = v163
		goto L55
	} else {
		goto L56
	}
L50:
	;
	v187 = *(*int32)(unsafe.Add(mBase, uint32(v163+int32(-17))))
	v188 = v187
	goto L49
L51:
	;
	v184 = *(*int32)(unsafe.Add(mBase, uint32(v163+int32(-9))))
	v188 = v184
	goto L49
L52:
	;
	v181 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v163+int32(-5)))))
	v188 = v181
	goto L49
L53:
	;
	v178 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v163+int32(-3)))))
	v188 = v178
	goto L49
L54:
	;
	v188 = int32(base.Ui32(v171) >> (uint(int32(3)) % 32))
	goto L49
L55:
	;
	if base.Ui32(v191) <= base.Ui32(v209) {
		v225 = v191
		goto L61
	} else {
		goto L62
	}
L56:
	;
	v197 = v163
	goto L57
L57:
	;
	v198 = int32(*(*int8)(unsafe.Add(mBase, uint32(v197))))
	v199 = F_strchr(m, v164, v198)
	mBase = m.M
	if v199 == int32(0) {
		v209 = v197
		goto L55
	} else {
		goto L59
	}
L58:
	;
	v209 = v203
	goto L55
L59:
	;
	v203 = v197 + int32(1)
	if base.Ui32(v203) <= base.Ui32(v191) {
		v197 = v203
		goto L57
	} else {
		goto L60
	}
L60:
	;
	goto L58
L61:
	;
	v230 = v225 - v209 + int32(1)
	if v163 == v209 {
		goto L67
	} else {
		goto L68
	}
L62:
	;
	v213 = v191
	goto L63
L63:
	;
	v216 = int32(*(*int8)(unsafe.Add(mBase, uint32(v213))))
	v217 = F_strchr(m, v164, v216)
	mBase = m.M
	if v217 == int32(0) {
		v225 = v213
		goto L61
	} else {
		goto L65
	}
L64:
	;
	v225 = v209
	goto L61
L65:
	;
	v221 = v213 + int32(-1)
	if base.Ui32(v209) < base.Ui32(v221) {
		v213 = v221
		goto L63
	} else {
		goto L66
	}
L66:
	;
	goto L64
L67:
	;
	v234 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v163+v230))) = uint8(v234)
	v236 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v170))))
	switch v236 & int32(7) {
	case 0:
		goto L74
	case 1:
		goto L73
	case 2:
		goto L72
	case 3:
		goto L71
	case 4:
		goto L70
	default:
		goto L69
	}
L68:
	;
	v232 = F_memmove(m, v163, v209, v230)
	mBase = m.M
	goto L67
L69:
	;
	goto L48
L70:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v163+int32(-17)))) = base.I64_extend_i32_u(v230)
	goto L69
L71:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v163+int32(-9)))) = v230
	goto L48
L72:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v163+int32(-5)))) = uint16(v230)
	goto L48
L73:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v163+int32(-3)))) = uint8(v230)
	goto L48
L74:
	;
	v240 = v230 << (uint(int32(3)) % 32)
	*(*uint8)(unsafe.Add(mBase, uint32(v170))) = uint8(v240)
	goto L48
L75:
	;
	v739 = v151 + int32(1)
	v740 = *(*int32)(unsafe.Add(mBase, uint32(v18)+12))
	if v739 < v740 {
		v148 = v731
		v151 = v739
		v152 = v258
		v154 = v733
		v156 = v734
		goto L46
	} else {
		goto L230
	}
L76:
	;
	v316 = F_sdssplitargs(m, v163, v18+int32(8))
	mBase = m.M
	v317 = m.ExcPending
	if v317 != 0 {
		goto L18
	} else {
		goto L95
	}
L77:
	;
	if v154 != 0 {
		goto L81
	} else {
		goto L82
	}
L78:
	;
	if v259 != 0 {
		goto L76
	} else {
		goto L79
	}
L79:
	;
	goto L77
L80:
	;
	v303 = F_valkey_realloc(m, v148, v156<<(uint(int32(2))%32)+int32(4))
	mBase = m.M
	v304 = m.ExcPending
	if v304 != 0 {
		goto L18
	} else {
		goto L92
	}
L81:
	;
	v263 = int32(_a497)
	v266 = int32(*(*uint8)(unsafe.Add(mBase, _consts[231])))
	v267 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v163))))
	if v267 == int32(0) {
		v290 = v266
		v291 = v267
		goto L84
	} else {
		goto L85
	}
L82:
	;
	v298 = int32(0)
	goto L80
L83:
	;
	if v291-v290&int32(255) != 0 {
		v298 = v154
		goto L80
	} else {
		goto L91
	}
L84:
	;
	goto L83
L85:
	;
	if v267 != v266&int32(255) {
		v290 = v266
		v291 = v267
		goto L84
	} else {
		goto L86
	}
L86:
	;
	v273 = v163
	v274 = v263
	goto L87
L87:
	;
	v277 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v274)+1)))
	v278 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v273)+1)))
	if v278 == int32(0) {
		v290 = v277
		v291 = v278
		goto L84
	} else {
		goto L89
	}
L88:
	;
	v290 = v277
	v291 = v278
	goto L84
L89:
	;
	v281 = int32(1)
	if v278 == v277&int32(255) {
		v273 = v273 + v281
		v274 = v274 + v281
		goto L87
	} else {
		goto L90
	}
L90:
	;
	goto L88
L91:
	;
	v295 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v66)+16)) = v295
	v298 = v295
	goto L80
L92:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v66)+12)) = v303
	v306 = *(*int32)(unsafe.Add(mBase, uint32(v66)+8))
	v308 = v306 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v66)+8)) = v308
	*(*int32)(unsafe.Add(mBase, uint32(v303+v306<<(uint(int32(2))%32)))) = v163
	v731 = v303
	v733 = v298
	v734 = v308
	goto L75
L93:
	;
	v571 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v565+int32(-1)))))
	switch v571 & int32(7) {
	case 0:
		goto L181
	case 1:
		goto L180
	case 2:
		goto L179
	case 3:
		goto L178
	case 4:
		goto L177
	default:
		goto L175
	}
L94:
	;
	v539 = F_sdsnew(m, int32(_a498))
	mBase = m.M
	v540 = m.ExcPending
	if v540 != 0 {
		goto L18
	} else {
		goto L167
	}
L95:
	;
	if v316 == int32(0) {
		goto L94
	} else {
		goto L96
	}
L96:
	;
	v321 = *(*int32)(unsafe.Add(mBase, _consts[230]))
	v322 = *(*int32)(unsafe.Add(mBase, uint32(v316)))
	v323 = F_dictFind(m, v321, v322)
	mBase = m.M
	v324 = m.ExcPending
	if v324 != 0 {
		goto L18
	} else {
		goto L98
	}
L97:
	;
	v331 = *(*int32)(unsafe.Add(mBase, uint32(v316)))
	v332 = int32(_a499)
	v335 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v331))))
	if v335 != 0 {
		goto L104
	} else {
		goto L105
	}
L98:
	;
	if v323 == int32(0) {
		goto L97
	} else {
		goto L99
	}
L99:
	;
	v327 = *(*int32)(unsafe.Add(mBase, uint32(v323)+8))
	goto L100
L100:
	;
	if v327 == int32(0) {
		goto L97
	} else {
		goto L101
	}
L101:
	;
	v330 = *(*int32)(unsafe.Add(mBase, uint32(v316)))
	v565 = v330
	goto L93
L102:
	;
	if v367-v369 == int32(0) {
		v565 = v331
		goto L93
	} else {
		goto L114
	}
L103:
	;
	v367 = F_tolower(m, v363)
	mBase = m.M
	v368 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v364))))
	v369 = F_tolower(m, v368)
	mBase = m.M
	goto L102
L104:
	;
	v337 = v331
	v338 = v332
	v339 = v335
	goto L107
L105:
	;
	v363 = int32(0)
	v364 = v332
	goto L103
L106:
	;
	v363 = v360 & int32(255)
	v364 = v359
	goto L103
L107:
	;
	v341 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v338))))
	if v341 == int32(0) {
		v359 = v338
		v360 = v339
		goto L106
	} else {
		goto L109
	}
L108:
	;
	v359 = v353
	v360 = int32(0)
	goto L106
L109:
	;
	v345 = v339 & int32(255)
	if v345 == v341 {
		goto L110
	} else {
		goto L111
	}
L110:
	;
	v352 = int32(1)
	v353 = v338 + v352
	v354 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v337)+1)))
	if v354 != 0 {
		v337 = v337 + v352
		v338 = v353
		v339 = v354
		goto L107
	} else {
		goto L113
	}
L111:
	;
	v347 = F_tolower(m, v345)
	mBase = m.M
	v348 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v338))))
	v349 = F_tolower(m, v348)
	mBase = m.M
	if v347 == v349 {
		goto L110
	} else {
		goto L112
	}
L112:
	;
	v351 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v337))))
	v359 = v338
	v360 = v351
	goto L106
L113:
	;
	goto L108
L114:
	;
	v373 = int32(_a500)
	v376 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v331))))
	if v376 != 0 {
		goto L117
	} else {
		goto L118
	}
L115:
	;
	if v408-v410 == int32(0) {
		v565 = v331
		goto L93
	} else {
		goto L127
	}
L116:
	;
	v408 = F_tolower(m, v404)
	mBase = m.M
	v409 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v405))))
	v410 = F_tolower(m, v409)
	mBase = m.M
	goto L115
L117:
	;
	v378 = v331
	v379 = v373
	v380 = v376
	goto L120
L118:
	;
	v404 = int32(0)
	v405 = v373
	goto L116
L119:
	;
	v404 = v401 & int32(255)
	v405 = v400
	goto L116
L120:
	;
	v382 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v379))))
	if v382 == int32(0) {
		v400 = v379
		v401 = v380
		goto L119
	} else {
		goto L122
	}
L121:
	;
	v400 = v394
	v401 = int32(0)
	goto L119
L122:
	;
	v386 = v380 & int32(255)
	if v386 == v382 {
		goto L123
	} else {
		goto L124
	}
L123:
	;
	v393 = int32(1)
	v394 = v379 + v393
	v395 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v378)+1)))
	if v395 != 0 {
		v378 = v378 + v393
		v379 = v394
		v380 = v395
		goto L120
	} else {
		goto L126
	}
L124:
	;
	v388 = F_tolower(m, v386)
	mBase = m.M
	v389 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v379))))
	v390 = F_tolower(m, v389)
	mBase = m.M
	if v388 == v390 {
		goto L123
	} else {
		goto L125
	}
L125:
	;
	v392 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v378))))
	v400 = v379
	v401 = v392
	goto L119
L126:
	;
	goto L121
L127:
	;
	v414 = int32(_a501)
	v417 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v331))))
	if v417 != 0 {
		goto L130
	} else {
		goto L131
	}
L128:
	;
	if v449-v451 == int32(0) {
		v565 = v331
		goto L93
	} else {
		goto L140
	}
L129:
	;
	v449 = F_tolower(m, v445)
	mBase = m.M
	v450 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v446))))
	v451 = F_tolower(m, v450)
	mBase = m.M
	goto L128
L130:
	;
	v419 = v331
	v420 = v414
	v421 = v417
	goto L133
L131:
	;
	v445 = int32(0)
	v446 = v414
	goto L129
L132:
	;
	v445 = v442 & int32(255)
	v446 = v441
	goto L129
L133:
	;
	v423 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v420))))
	if v423 == int32(0) {
		v441 = v420
		v442 = v421
		goto L132
	} else {
		goto L135
	}
L134:
	;
	v441 = v435
	v442 = int32(0)
	goto L132
L135:
	;
	v427 = v421 & int32(255)
	if v427 == v423 {
		goto L136
	} else {
		goto L137
	}
L136:
	;
	v434 = int32(1)
	v435 = v420 + v434
	v436 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v419)+1)))
	if v436 != 0 {
		v419 = v419 + v434
		v420 = v435
		v421 = v436
		goto L133
	} else {
		goto L139
	}
L137:
	;
	v429 = F_tolower(m, v427)
	mBase = m.M
	v430 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v420))))
	v431 = F_tolower(m, v430)
	mBase = m.M
	if v429 == v431 {
		goto L136
	} else {
		goto L138
	}
L138:
	;
	v433 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v419))))
	v441 = v420
	v442 = v433
	goto L132
L139:
	;
	goto L134
L140:
	;
	v455 = int32(_a502)
	v458 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v331))))
	if v458 != 0 {
		goto L143
	} else {
		goto L144
	}
L141:
	;
	if v490-v492 == int32(0) {
		v565 = v331
		goto L93
	} else {
		goto L153
	}
L142:
	;
	v490 = F_tolower(m, v486)
	mBase = m.M
	v491 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v487))))
	v492 = F_tolower(m, v491)
	mBase = m.M
	goto L141
L143:
	;
	v460 = v331
	v461 = v455
	v462 = v458
	goto L146
L144:
	;
	v486 = int32(0)
	v487 = v455
	goto L142
L145:
	;
	v486 = v483 & int32(255)
	v487 = v482
	goto L142
L146:
	;
	v464 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v461))))
	if v464 == int32(0) {
		v482 = v461
		v483 = v462
		goto L145
	} else {
		goto L148
	}
L147:
	;
	v482 = v476
	v483 = int32(0)
	goto L145
L148:
	;
	v468 = v462 & int32(255)
	if v468 == v464 {
		goto L149
	} else {
		goto L150
	}
L149:
	;
	v475 = int32(1)
	v476 = v461 + v475
	v477 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v460)+1)))
	if v477 != 0 {
		v460 = v460 + v475
		v461 = v476
		v462 = v477
		goto L146
	} else {
		goto L152
	}
L150:
	;
	v470 = F_tolower(m, v468)
	mBase = m.M
	v471 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v461))))
	v472 = F_tolower(m, v471)
	mBase = m.M
	if v470 == v472 {
		goto L149
	} else {
		goto L151
	}
L151:
	;
	v474 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v460))))
	v482 = v461
	v483 = v474
	goto L145
L152:
	;
	goto L147
L153:
	;
	v496 = int32(_a503)
	v499 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v331))))
	if v499 != 0 {
		goto L156
	} else {
		goto L157
	}
L154:
	;
	if v531-v533 == int32(0) {
		v565 = v331
		goto L93
	} else {
		goto L166
	}
L155:
	;
	v531 = F_tolower(m, v527)
	mBase = m.M
	v532 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v528))))
	v533 = F_tolower(m, v532)
	mBase = m.M
	goto L154
L156:
	;
	v501 = v331
	v502 = v496
	v503 = v499
	goto L159
L157:
	;
	v527 = int32(0)
	v528 = v496
	goto L155
L158:
	;
	v527 = v524 & int32(255)
	v528 = v523
	goto L155
L159:
	;
	v505 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v502))))
	if v505 == int32(0) {
		v523 = v502
		v524 = v503
		goto L158
	} else {
		goto L161
	}
L160:
	;
	v523 = v517
	v524 = int32(0)
	goto L158
L161:
	;
	v509 = v503 & int32(255)
	if v509 == v505 {
		goto L162
	} else {
		goto L163
	}
L162:
	;
	v516 = int32(1)
	v517 = v502 + v516
	v518 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v501)+1)))
	if v518 != 0 {
		v501 = v501 + v516
		v502 = v517
		v503 = v518
		goto L159
	} else {
		goto L165
	}
L163:
	;
	v511 = F_tolower(m, v509)
	mBase = m.M
	v512 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v502))))
	v513 = F_tolower(m, v512)
	mBase = m.M
	if v511 == v513 {
		goto L162
	} else {
		goto L164
	}
L164:
	;
	v515 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v501))))
	v523 = v502
	v524 = v515
	goto L158
L165:
	;
	goto L160
L166:
	;
	goto L94
L167:
	;
	v541 = F_sdscatsds(m, v539, v163)
	mBase = m.M
	v542 = m.ExcPending
	if v542 != 0 {
		goto L18
	} else {
		goto L168
	}
L168:
	;
	if v316 == int32(0) {
		goto L169
	} else {
		goto L170
	}
L169:
	;
	F_sdsfree(m, v163)
	mBase = m.M
	v549 = m.ExcPending
	if v549 != 0 {
		goto L18
	} else {
		goto L172
	}
L170:
	;
	v545 = *(*int32)(unsafe.Add(mBase, uint32(v18)+8))
	F_sdsfreesplitres(m, v316, v545)
	mBase = m.M
	v547 = m.ExcPending
	if v547 != 0 {
		goto L18
	} else {
		goto L171
	}
L171:
	;
	goto L169
L172:
	;
	v554 = F_valkey_realloc(m, v148, v156<<(uint(int32(2))%32)+int32(4))
	mBase = m.M
	v555 = m.ExcPending
	if v555 != 0 {
		goto L18
	} else {
		goto L173
	}
L173:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v66)+12)) = v554
	v557 = *(*int32)(unsafe.Add(mBase, uint32(v66)+8))
	v559 = v557 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v66)+8)) = v559
	*(*int32)(unsafe.Add(mBase, uint32(v554+v557<<(uint(int32(2))%32)))) = v541
	v731 = v554
	v733 = v154
	v734 = v559
	goto L75
L174:
	;
	v611 = F_valkey_realloc(m, v148, v156<<(uint(int32(2))%32)+int32(4))
	mBase = m.M
	v612 = m.ExcPending
	if v612 != 0 {
		goto L18
	} else {
		goto L186
	}
L175:
	;
	goto L174
L176:
	;
	if v588 == int32(0) {
		goto L175
	} else {
		goto L182
	}
L177:
	;
	v587 = *(*int32)(unsafe.Add(mBase, uint32(v565+int32(-17))))
	v588 = v587
	goto L176
L178:
	;
	v584 = *(*int32)(unsafe.Add(mBase, uint32(v565+int32(-9))))
	v588 = v584
	goto L176
L179:
	;
	v581 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v565+int32(-5)))))
	v588 = v581
	goto L176
L180:
	;
	v578 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v565+int32(-3)))))
	v588 = v578
	goto L176
L181:
	;
	v588 = int32(base.Ui32(v571) >> (uint(int32(3)) % 32))
	goto L176
L182:
	;
	v593 = int32(0)
	goto L183
L183:
	;
	v596 = v565 + v593
	v597 = int32(*(*int8)(unsafe.Add(mBase, uint32(v596))))
	v598 = F_tolower(m, v597)
	mBase = m.M
	*(*uint8)(unsafe.Add(mBase, uint32(v596))) = uint8(v598)
	v601 = v593 + int32(1)
	if v601 != v588 {
		v593 = v601
		goto L183
	} else {
		goto L185
	}
L184:
	;
	goto L175
L185:
	;
	goto L184
L186:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v66)+12)) = v611
	v614 = *(*int32)(unsafe.Add(mBase, uint32(v66)+8))
	v616 = v614 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v66)+8)) = v616
	*(*int32)(unsafe.Add(mBase, uint32(v611+v614<<(uint(int32(2))%32)))) = v163
	v623 = *(*int32)(unsafe.Add(mBase, _consts[230]))
	v624 = *(*int32)(unsafe.Add(mBase, uint32(v316)))
	v625 = F_dictFind(m, v623, v624)
	mBase = m.M
	v626 = m.ExcPending
	if v626 != 0 {
		goto L18
	} else {
		goto L188
	}
L187:
	;
	v645 = *(*int32)(unsafe.Add(mBase, uint32(v316)))
	v647 = *(*int32)(unsafe.Add(mBase, _consts[232]))
	if v647 == int32(0) {
		goto L196
	} else {
		goto L197
	}
L188:
	;
	if v625 == int32(0) {
		goto L187
	} else {
		goto L189
	}
L189:
	;
	v629 = *(*int32)(unsafe.Add(mBase, uint32(v625)+8))
	goto L190
L190:
	;
	if v629 == int32(0) {
		goto L187
	} else {
		goto L191
	}
L191:
	;
	v632 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v629)+8)))
	if v632&int32(128) == int32(0) {
		goto L187
	} else {
		goto L192
	}
L192:
	;
	v637 = *(*int32)(unsafe.Add(mBase, uint32(v316)))
	F_sdsfree(m, v637)
	mBase = m.M
	v639 = m.ExcPending
	if v639 != 0 {
		goto L18
	} else {
		goto L193
	}
L193:
	;
	v640 = *(*int32)(unsafe.Add(mBase, uint32(v629)+4))
	v641 = F_sdsnew(m, v640)
	mBase = m.M
	v642 = m.ExcPending
	if v642 != 0 {
		goto L18
	} else {
		goto L194
	}
L194:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v316))) = v641
	goto L187
L195:
	;
	v728 = *(*int32)(unsafe.Add(mBase, uint32(v18)+8))
	F_sdsfreesplitres(m, v316, v728)
	mBase = m.M
	v730 = m.ExcPending
	if v730 != 0 {
		goto L18
	} else {
		goto L229
	}
L196:
	;
	v713 = *(*int32)(unsafe.Add(mBase, uint32(v66)))
	v714 = F_dictFetchValue(m, v713, v645)
	mBase = m.M
	v715 = m.ExcPending
	if v715 != 0 {
		goto L18
	} else {
		goto L223
	}
L197:
	;
	v650 = *(*int32)(unsafe.Add(mBase, uint32(v18)+8))
	if v650 < int32(2) {
		goto L196
	} else {
		goto L198
	}
L198:
	;
	v653 = int32(_a503)
	v656 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v645))))
	if v656 != 0 {
		goto L201
	} else {
		goto L202
	}
L199:
	;
	if v688-v690 != 0 {
		goto L196
	} else {
		goto L211
	}
L200:
	;
	v688 = F_tolower(m, v684)
	mBase = m.M
	v689 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v685))))
	v690 = F_tolower(m, v689)
	mBase = m.M
	goto L199
L201:
	;
	v658 = v645
	v659 = v653
	v660 = v656
	goto L204
L202:
	;
	v684 = int32(0)
	v685 = v653
	goto L200
L203:
	;
	v684 = v681 & int32(255)
	v685 = v680
	goto L200
L204:
	;
	v662 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v659))))
	if v662 == int32(0) {
		v680 = v659
		v681 = v660
		goto L203
	} else {
		goto L206
	}
L205:
	;
	v680 = v674
	v681 = int32(0)
	goto L203
L206:
	;
	v666 = v660 & int32(255)
	if v666 == v662 {
		goto L207
	} else {
		goto L208
	}
L207:
	;
	v673 = int32(1)
	v674 = v659 + v673
	v675 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v658)+1)))
	if v675 != 0 {
		v658 = v658 + v673
		v659 = v674
		v660 = v675
		goto L204
	} else {
		goto L210
	}
L208:
	;
	v668 = F_tolower(m, v666)
	mBase = m.M
	v669 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v659))))
	v670 = F_tolower(m, v669)
	mBase = m.M
	if v668 == v670 {
		goto L207
	} else {
		goto L209
	}
L209:
	;
	v672 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v658))))
	v680 = v659
	v681 = v672
	goto L203
L210:
	;
	goto L205
L211:
	;
	v692 = F_sdsempty(m)
	mBase = m.M
	v693 = m.ExcPending
	if v693 != 0 {
		goto L18
	} else {
		goto L212
	}
L212:
	;
	v694 = *(*int64)(unsafe.Add(mBase, uint32(v316)))
	*(*int64)(unsafe.Add(mBase, uint32(v18))) = v694
	v697 = F_sdscatfmt(m, v692, int32(_a504), v18)
	mBase = m.M
	v698 = m.ExcPending
	if v698 != 0 {
		goto L18
	} else {
		goto L213
	}
L213:
	;
	v699 = *(*int32)(unsafe.Add(mBase, uint32(v66)))
	v700 = F_dictFetchValue(m, v699, v697)
	mBase = m.M
	v701 = m.ExcPending
	if v701 != 0 {
		goto L18
	} else {
		goto L215
	}
L214:
	;
	v709 = F_listAddNodeTail(m, v708, v258)
	mBase = m.M
	v710 = m.ExcPending
	if v710 != 0 {
		goto L18
	} else {
		goto L220
	}
L215:
	;
	if v700 != 0 {
		v708 = v700
		goto L214
	} else {
		goto L216
	}
L216:
	;
	v702 = F_listCreate(m)
	mBase = m.M
	v703 = m.ExcPending
	if v703 != 0 {
		goto L18
	} else {
		goto L217
	}
L217:
	;
	v704 = F_sdsdup(m, v697)
	mBase = m.M
	v705 = m.ExcPending
	if v705 != 0 {
		goto L18
	} else {
		goto L218
	}
L218:
	;
	v706 = F_dictAdd(m, v699, v704, v702)
	mBase = m.M
	v707 = m.ExcPending
	if v707 != 0 {
		goto L18
	} else {
		goto L219
	}
L219:
	;
	v708 = v702
	goto L214
L220:
	;
	F_sdsfree(m, v697)
	mBase = m.M
	v712 = m.ExcPending
	if v712 != 0 {
		goto L18
	} else {
		goto L221
	}
L221:
	;
	goto L195
L222:
	;
	v723 = F_listAddNodeTail(m, v722, v258)
	mBase = m.M
	v724 = m.ExcPending
	if v724 != 0 {
		goto L18
	} else {
		goto L228
	}
L223:
	;
	if v714 != 0 {
		v722 = v714
		goto L222
	} else {
		goto L224
	}
L224:
	;
	v716 = F_listCreate(m)
	mBase = m.M
	v717 = m.ExcPending
	if v717 != 0 {
		goto L18
	} else {
		goto L225
	}
L225:
	;
	v718 = F_sdsdup(m, v645)
	mBase = m.M
	v719 = m.ExcPending
	if v719 != 0 {
		goto L18
	} else {
		goto L226
	}
L226:
	;
	v720 = F_dictAdd(m, v713, v718, v716)
	mBase = m.M
	v721 = m.ExcPending
	if v721 != 0 {
		goto L18
	} else {
		goto L227
	}
L227:
	;
	v722 = v716
	goto L222
L228:
	;
	goto L195
L229:
	;
	v731 = v611
	v733 = v154
	v734 = v616
	goto L75
L230:
	;
	goto L47
L231:
	;
	v759 = *(*int32)(unsafe.Add(mBase, uint32(v18)+12))
	F_sdsfreesplitres(m, v137, v759)
	mBase = m.M
	v761 = m.ExcPending
	if v761 != 0 {
		goto L18
	} else {
		goto L232
	}
L232:
	;
	F_sdsfree(m, v91)
	mBase = m.M
	v763 = m.ExcPending
	if v763 != 0 {
		goto L18
	} else {
		goto L233
	}
L233:
	;
	v764 = v66
	goto L1
}
func F_rewriteConfigRewriteLine(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
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
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
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
	var v63 int32
	_ = v63
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v101 int32
	_ = v101
	v8 = F_sdsnew(m, l1)
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return int32(0)
	} else {
		v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		v13 = F_dictFetchValue(m, v12, v8)
		mBase = m.M
		v14 = m.ExcPending
		if v14 != 0 {
			return int32(0)
		} else {
			v15 = F_sdsnew(m, l1)
			mBase = m.M
			v16 = m.ExcPending
			if v16 != 0 {
				return int32(0)
			} else {
				v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
				v19 = F_dictAdd(m, v17, v15, int32(0))
				mBase = m.M
				v20 = m.ExcPending
				if v20 != 0 {
					return int32(0)
				} else {
					if v19 == int32(0) {
						if l3 != 0 {
							if v13 == int32(0) {
								v56 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
								if v56 != 0 {
									v60 = F_sdsnew(m, int32(_a497))
									mBase = m.M
									v61 = m.ExcPending
									if v61 != 0 {
										return int32(0)
									} else {
										v62 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
										v63 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
										v68 = F_valkey_realloc(m, v62, v63<<(uint(int32(2))%32)+int32(4))
										mBase = m.M
										v69 = m.ExcPending
										if v69 != 0 {
											return int32(0)
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v68
											v71 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
											v73 = v71 + int32(1)
											*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v73
											*(*int32)(unsafe.Add(mBase, uint32(v68+v71<<(uint(int32(2))%32)))) = v60
											*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = int32(0)
											v81 = v73
											v83 = v68
											v89 = F_valkey_realloc(m, v83, v81<<(uint(int32(2))%32)+int32(4))
											mBase = m.M
											v90 = m.ExcPending
											if v90 != 0 {
												return int32(0)
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v89
												v92 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
												*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v92 + int32(1)
												*(*int32)(unsafe.Add(mBase, uint32(v89+v92<<(uint(int32(2))%32)))) = l2
												F_sdsfree(m, v8)
												mBase = m.M
												v101 = m.ExcPending
												if v101 != 0 {
													return int32(0)
												} else {
													return int32(1)
												}
											}
										}
									}
								} else {
									v57 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
									v58 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
									v81 = v57
									v83 = v58
									v89 = F_valkey_realloc(m, v83, v81<<(uint(int32(2))%32)+int32(4))
									mBase = m.M
									v90 = m.ExcPending
									if v90 != 0 {
										return int32(0)
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v89
										v92 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
										*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v92 + int32(1)
										*(*int32)(unsafe.Add(mBase, uint32(v89+v92<<(uint(int32(2))%32)))) = l2
										F_sdsfree(m, v8)
										mBase = m.M
										v101 = m.ExcPending
										if v101 != 0 {
											return int32(0)
										} else {
											return int32(1)
										}
									}
								}
							} else {
								v34 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
								v35 = *(*int32)(unsafe.Add(mBase, uint32(v34)+8))
								F_listDelNode(m, v13, v34)
								mBase = m.M
								v37 = m.ExcPending
								if v37 != 0 {
									return int32(0)
								} else {
									v38 = *(*int32)(unsafe.Add(mBase, uint32(v13)+20))
									if v38 != 0 {
										v42 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
										v44 = v35 << (uint(int32(2)) % 32)
										v46 = *(*int32)(unsafe.Add(mBase, uint32(v42+v44)))
										F_sdsfree(m, v46)
										mBase = m.M
										v48 = m.ExcPending
										if v48 != 0 {
											return int32(0)
										} else {
											v49 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
											*(*int32)(unsafe.Add(mBase, uint32(v49+v44))) = l2
											F_sdsfree(m, v8)
											mBase = m.M
											v53 = m.ExcPending
											if v53 != 0 {
												return int32(0)
											} else {
												return int32(1)
											}
										}
									} else {
										v39 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
										v40 = F_dictDelete(m, v39, v8)
										mBase = m.M
										v41 = m.ExcPending
										if v41 != 0 {
											return int32(0)
										} else {
											v42 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
											v44 = v35 << (uint(int32(2)) % 32)
											v46 = *(*int32)(unsafe.Add(mBase, uint32(v42+v44)))
											F_sdsfree(m, v46)
											mBase = m.M
											v48 = m.ExcPending
											if v48 != 0 {
												return int32(0)
											} else {
												v49 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
												*(*int32)(unsafe.Add(mBase, uint32(v49+v44))) = l2
												F_sdsfree(m, v8)
												mBase = m.M
												v53 = m.ExcPending
												if v53 != 0 {
													return int32(0)
												} else {
													return int32(1)
												}
											}
										}
									}
								}
							}
						} else {
							if v13 != 0 {
								if v13 == int32(0) {
									v56 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
									if v56 != 0 {
										v60 = F_sdsnew(m, int32(_a497))
										mBase = m.M
										v61 = m.ExcPending
										if v61 != 0 {
											return int32(0)
										} else {
											v62 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
											v63 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
											v68 = F_valkey_realloc(m, v62, v63<<(uint(int32(2))%32)+int32(4))
											mBase = m.M
											v69 = m.ExcPending
											if v69 != 0 {
												return int32(0)
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v68
												v71 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
												v73 = v71 + int32(1)
												*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v73
												*(*int32)(unsafe.Add(mBase, uint32(v68+v71<<(uint(int32(2))%32)))) = v60
												*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = int32(0)
												v81 = v73
												v83 = v68
												v89 = F_valkey_realloc(m, v83, v81<<(uint(int32(2))%32)+int32(4))
												mBase = m.M
												v90 = m.ExcPending
												if v90 != 0 {
													return int32(0)
												} else {
													*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v89
													v92 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
													*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v92 + int32(1)
													*(*int32)(unsafe.Add(mBase, uint32(v89+v92<<(uint(int32(2))%32)))) = l2
													F_sdsfree(m, v8)
													mBase = m.M
													v101 = m.ExcPending
													if v101 != 0 {
														return int32(0)
													} else {
														return int32(1)
													}
												}
											}
										}
									} else {
										v57 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
										v58 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
										v81 = v57
										v83 = v58
										v89 = F_valkey_realloc(m, v83, v81<<(uint(int32(2))%32)+int32(4))
										mBase = m.M
										v90 = m.ExcPending
										if v90 != 0 {
											return int32(0)
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v89
											v92 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
											*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v92 + int32(1)
											*(*int32)(unsafe.Add(mBase, uint32(v89+v92<<(uint(int32(2))%32)))) = l2
											F_sdsfree(m, v8)
											mBase = m.M
											v101 = m.ExcPending
											if v101 != 0 {
												return int32(0)
											} else {
												return int32(1)
											}
										}
									}
								} else {
									v34 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
									v35 = *(*int32)(unsafe.Add(mBase, uint32(v34)+8))
									F_listDelNode(m, v13, v34)
									mBase = m.M
									v37 = m.ExcPending
									if v37 != 0 {
										return int32(0)
									} else {
										v38 = *(*int32)(unsafe.Add(mBase, uint32(v13)+20))
										if v38 != 0 {
											v42 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
											v44 = v35 << (uint(int32(2)) % 32)
											v46 = *(*int32)(unsafe.Add(mBase, uint32(v42+v44)))
											F_sdsfree(m, v46)
											mBase = m.M
											v48 = m.ExcPending
											if v48 != 0 {
												return int32(0)
											} else {
												v49 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
												*(*int32)(unsafe.Add(mBase, uint32(v49+v44))) = l2
												F_sdsfree(m, v8)
												mBase = m.M
												v53 = m.ExcPending
												if v53 != 0 {
													return int32(0)
												} else {
													return int32(1)
												}
											}
										} else {
											v39 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
											v40 = F_dictDelete(m, v39, v8)
											mBase = m.M
											v41 = m.ExcPending
											if v41 != 0 {
												return int32(0)
											} else {
												v42 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
												v44 = v35 << (uint(int32(2)) % 32)
												v46 = *(*int32)(unsafe.Add(mBase, uint32(v42+v44)))
												F_sdsfree(m, v46)
												mBase = m.M
												v48 = m.ExcPending
												if v48 != 0 {
													return int32(0)
												} else {
													v49 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
													*(*int32)(unsafe.Add(mBase, uint32(v49+v44))) = l2
													F_sdsfree(m, v8)
													mBase = m.M
													v53 = m.ExcPending
													if v53 != 0 {
														return int32(0)
													} else {
														return int32(1)
													}
												}
											}
										}
									}
								}
							} else {
								v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
								if v25 != 0 {
									v56 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
									if v56 != 0 {
										v60 = F_sdsnew(m, int32(_a497))
										mBase = m.M
										v61 = m.ExcPending
										if v61 != 0 {
											return int32(0)
										} else {
											v62 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
											v63 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
											v68 = F_valkey_realloc(m, v62, v63<<(uint(int32(2))%32)+int32(4))
											mBase = m.M
											v69 = m.ExcPending
											if v69 != 0 {
												return int32(0)
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v68
												v71 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
												v73 = v71 + int32(1)
												*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v73
												*(*int32)(unsafe.Add(mBase, uint32(v68+v71<<(uint(int32(2))%32)))) = v60
												*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = int32(0)
												v81 = v73
												v83 = v68
												v89 = F_valkey_realloc(m, v83, v81<<(uint(int32(2))%32)+int32(4))
												mBase = m.M
												v90 = m.ExcPending
												if v90 != 0 {
													return int32(0)
												} else {
													*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v89
													v92 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
													*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v92 + int32(1)
													*(*int32)(unsafe.Add(mBase, uint32(v89+v92<<(uint(int32(2))%32)))) = l2
													F_sdsfree(m, v8)
													mBase = m.M
													v101 = m.ExcPending
													if v101 != 0 {
														return int32(0)
													} else {
														return int32(1)
													}
												}
											}
										}
									} else {
										v57 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
										v58 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
										v81 = v57
										v83 = v58
										v89 = F_valkey_realloc(m, v83, v81<<(uint(int32(2))%32)+int32(4))
										mBase = m.M
										v90 = m.ExcPending
										if v90 != 0 {
											return int32(0)
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v89
											v92 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
											*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v92 + int32(1)
											*(*int32)(unsafe.Add(mBase, uint32(v89+v92<<(uint(int32(2))%32)))) = l2
											F_sdsfree(m, v8)
											mBase = m.M
											v101 = m.ExcPending
											if v101 != 0 {
												return int32(0)
											} else {
												return int32(1)
											}
										}
									}
								} else {
									F_sdsfree(m, l2)
									mBase = m.M
									v27 = m.ExcPending
									if v27 != 0 {
										return int32(0)
									} else {
										F_sdsfree(m, v8)
										mBase = m.M
										v29 = m.ExcPending
										if v29 != 0 {
											return int32(0)
										} else {
											return int32(0)
										}
									}
								}
							}
						}
					} else {
						F_sdsfree(m, v15)
						mBase = m.M
						v24 = m.ExcPending
						if v24 != 0 {
							return int32(0)
						} else {
							if l3 != 0 {
								if v13 == int32(0) {
									v56 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
									if v56 != 0 {
										v60 = F_sdsnew(m, int32(_a497))
										mBase = m.M
										v61 = m.ExcPending
										if v61 != 0 {
											return int32(0)
										} else {
											v62 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
											v63 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
											v68 = F_valkey_realloc(m, v62, v63<<(uint(int32(2))%32)+int32(4))
											mBase = m.M
											v69 = m.ExcPending
											if v69 != 0 {
												return int32(0)
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v68
												v71 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
												v73 = v71 + int32(1)
												*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v73
												*(*int32)(unsafe.Add(mBase, uint32(v68+v71<<(uint(int32(2))%32)))) = v60
												*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = int32(0)
												v81 = v73
												v83 = v68
												v89 = F_valkey_realloc(m, v83, v81<<(uint(int32(2))%32)+int32(4))
												mBase = m.M
												v90 = m.ExcPending
												if v90 != 0 {
													return int32(0)
												} else {
													*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v89
													v92 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
													*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v92 + int32(1)
													*(*int32)(unsafe.Add(mBase, uint32(v89+v92<<(uint(int32(2))%32)))) = l2
													F_sdsfree(m, v8)
													mBase = m.M
													v101 = m.ExcPending
													if v101 != 0 {
														return int32(0)
													} else {
														return int32(1)
													}
												}
											}
										}
									} else {
										v57 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
										v58 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
										v81 = v57
										v83 = v58
										v89 = F_valkey_realloc(m, v83, v81<<(uint(int32(2))%32)+int32(4))
										mBase = m.M
										v90 = m.ExcPending
										if v90 != 0 {
											return int32(0)
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v89
											v92 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
											*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v92 + int32(1)
											*(*int32)(unsafe.Add(mBase, uint32(v89+v92<<(uint(int32(2))%32)))) = l2
											F_sdsfree(m, v8)
											mBase = m.M
											v101 = m.ExcPending
											if v101 != 0 {
												return int32(0)
											} else {
												return int32(1)
											}
										}
									}
								} else {
									v34 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
									v35 = *(*int32)(unsafe.Add(mBase, uint32(v34)+8))
									F_listDelNode(m, v13, v34)
									mBase = m.M
									v37 = m.ExcPending
									if v37 != 0 {
										return int32(0)
									} else {
										v38 = *(*int32)(unsafe.Add(mBase, uint32(v13)+20))
										if v38 != 0 {
											v42 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
											v44 = v35 << (uint(int32(2)) % 32)
											v46 = *(*int32)(unsafe.Add(mBase, uint32(v42+v44)))
											F_sdsfree(m, v46)
											mBase = m.M
											v48 = m.ExcPending
											if v48 != 0 {
												return int32(0)
											} else {
												v49 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
												*(*int32)(unsafe.Add(mBase, uint32(v49+v44))) = l2
												F_sdsfree(m, v8)
												mBase = m.M
												v53 = m.ExcPending
												if v53 != 0 {
													return int32(0)
												} else {
													return int32(1)
												}
											}
										} else {
											v39 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
											v40 = F_dictDelete(m, v39, v8)
											mBase = m.M
											v41 = m.ExcPending
											if v41 != 0 {
												return int32(0)
											} else {
												v42 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
												v44 = v35 << (uint(int32(2)) % 32)
												v46 = *(*int32)(unsafe.Add(mBase, uint32(v42+v44)))
												F_sdsfree(m, v46)
												mBase = m.M
												v48 = m.ExcPending
												if v48 != 0 {
													return int32(0)
												} else {
													v49 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
													*(*int32)(unsafe.Add(mBase, uint32(v49+v44))) = l2
													F_sdsfree(m, v8)
													mBase = m.M
													v53 = m.ExcPending
													if v53 != 0 {
														return int32(0)
													} else {
														return int32(1)
													}
												}
											}
										}
									}
								}
							} else {
								if v13 != 0 {
									if v13 == int32(0) {
										v56 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
										if v56 != 0 {
											v60 = F_sdsnew(m, int32(_a497))
											mBase = m.M
											v61 = m.ExcPending
											if v61 != 0 {
												return int32(0)
											} else {
												v62 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
												v63 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
												v68 = F_valkey_realloc(m, v62, v63<<(uint(int32(2))%32)+int32(4))
												mBase = m.M
												v69 = m.ExcPending
												if v69 != 0 {
													return int32(0)
												} else {
													*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v68
													v71 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
													v73 = v71 + int32(1)
													*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v73
													*(*int32)(unsafe.Add(mBase, uint32(v68+v71<<(uint(int32(2))%32)))) = v60
													*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = int32(0)
													v81 = v73
													v83 = v68
													v89 = F_valkey_realloc(m, v83, v81<<(uint(int32(2))%32)+int32(4))
													mBase = m.M
													v90 = m.ExcPending
													if v90 != 0 {
														return int32(0)
													} else {
														*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v89
														v92 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
														*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v92 + int32(1)
														*(*int32)(unsafe.Add(mBase, uint32(v89+v92<<(uint(int32(2))%32)))) = l2
														F_sdsfree(m, v8)
														mBase = m.M
														v101 = m.ExcPending
														if v101 != 0 {
															return int32(0)
														} else {
															return int32(1)
														}
													}
												}
											}
										} else {
											v57 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
											v58 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
											v81 = v57
											v83 = v58
											v89 = F_valkey_realloc(m, v83, v81<<(uint(int32(2))%32)+int32(4))
											mBase = m.M
											v90 = m.ExcPending
											if v90 != 0 {
												return int32(0)
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v89
												v92 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
												*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v92 + int32(1)
												*(*int32)(unsafe.Add(mBase, uint32(v89+v92<<(uint(int32(2))%32)))) = l2
												F_sdsfree(m, v8)
												mBase = m.M
												v101 = m.ExcPending
												if v101 != 0 {
													return int32(0)
												} else {
													return int32(1)
												}
											}
										}
									} else {
										v34 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
										v35 = *(*int32)(unsafe.Add(mBase, uint32(v34)+8))
										F_listDelNode(m, v13, v34)
										mBase = m.M
										v37 = m.ExcPending
										if v37 != 0 {
											return int32(0)
										} else {
											v38 = *(*int32)(unsafe.Add(mBase, uint32(v13)+20))
											if v38 != 0 {
												v42 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
												v44 = v35 << (uint(int32(2)) % 32)
												v46 = *(*int32)(unsafe.Add(mBase, uint32(v42+v44)))
												F_sdsfree(m, v46)
												mBase = m.M
												v48 = m.ExcPending
												if v48 != 0 {
													return int32(0)
												} else {
													v49 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
													*(*int32)(unsafe.Add(mBase, uint32(v49+v44))) = l2
													F_sdsfree(m, v8)
													mBase = m.M
													v53 = m.ExcPending
													if v53 != 0 {
														return int32(0)
													} else {
														return int32(1)
													}
												}
											} else {
												v39 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
												v40 = F_dictDelete(m, v39, v8)
												mBase = m.M
												v41 = m.ExcPending
												if v41 != 0 {
													return int32(0)
												} else {
													v42 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
													v44 = v35 << (uint(int32(2)) % 32)
													v46 = *(*int32)(unsafe.Add(mBase, uint32(v42+v44)))
													F_sdsfree(m, v46)
													mBase = m.M
													v48 = m.ExcPending
													if v48 != 0 {
														return int32(0)
													} else {
														v49 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
														*(*int32)(unsafe.Add(mBase, uint32(v49+v44))) = l2
														F_sdsfree(m, v8)
														mBase = m.M
														v53 = m.ExcPending
														if v53 != 0 {
															return int32(0)
														} else {
															return int32(1)
														}
													}
												}
											}
										}
									}
								} else {
									v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
									if v25 != 0 {
										v56 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
										if v56 != 0 {
											v60 = F_sdsnew(m, int32(_a497))
											mBase = m.M
											v61 = m.ExcPending
											if v61 != 0 {
												return int32(0)
											} else {
												v62 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
												v63 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
												v68 = F_valkey_realloc(m, v62, v63<<(uint(int32(2))%32)+int32(4))
												mBase = m.M
												v69 = m.ExcPending
												if v69 != 0 {
													return int32(0)
												} else {
													*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v68
													v71 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
													v73 = v71 + int32(1)
													*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v73
													*(*int32)(unsafe.Add(mBase, uint32(v68+v71<<(uint(int32(2))%32)))) = v60
													*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = int32(0)
													v81 = v73
													v83 = v68
													v89 = F_valkey_realloc(m, v83, v81<<(uint(int32(2))%32)+int32(4))
													mBase = m.M
													v90 = m.ExcPending
													if v90 != 0 {
														return int32(0)
													} else {
														*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v89
														v92 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
														*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v92 + int32(1)
														*(*int32)(unsafe.Add(mBase, uint32(v89+v92<<(uint(int32(2))%32)))) = l2
														F_sdsfree(m, v8)
														mBase = m.M
														v101 = m.ExcPending
														if v101 != 0 {
															return int32(0)
														} else {
															return int32(1)
														}
													}
												}
											}
										} else {
											v57 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
											v58 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
											v81 = v57
											v83 = v58
											v89 = F_valkey_realloc(m, v83, v81<<(uint(int32(2))%32)+int32(4))
											mBase = m.M
											v90 = m.ExcPending
											if v90 != 0 {
												return int32(0)
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v89
												v92 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
												*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v92 + int32(1)
												*(*int32)(unsafe.Add(mBase, uint32(v89+v92<<(uint(int32(2))%32)))) = l2
												F_sdsfree(m, v8)
												mBase = m.M
												v101 = m.ExcPending
												if v101 != 0 {
													return int32(0)
												} else {
													return int32(1)
												}
											}
										}
									} else {
										F_sdsfree(m, l2)
										mBase = m.M
										v27 = m.ExcPending
										if v27 != 0 {
											return int32(0)
										} else {
											F_sdsfree(m, v8)
											mBase = m.M
											v29 = m.ExcPending
											if v29 != 0 {
												return int32(0)
											} else {
												return int32(0)
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
func F_rewriteConfigSdsOption(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v17 int32
	_ = v17
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v64 int32
	_ = v64
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
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
	var v85 int32
	_ = v85
	if l2 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	return
L2:
	;
	if l3 == int32(0) {
		v54 = int32(1)
		goto L9
	} else {
		goto L10
	}
L3:
	;
	v8 = F_sdsnew(m, l1)
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	return
L5:
	;
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v12 = F_dictAdd(m, v10, v8, int32(0))
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		goto L4
	} else {
		goto L6
	}
L6:
	;
	if v12 == int32(0) {
		goto L1
	} else {
		goto L7
	}
L7:
	;
	F_sdsfree(m, v8)
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		goto L4
	} else {
		goto L8
	}
L8:
	;
	return
L9:
	;
	v55 = F_sdsnew(m, l1)
	mBase = m.M
	v56 = m.ExcPending
	if v56 != 0 {
		goto L4
	} else {
		goto L19
	}
L10:
	;
	v23 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3))))
	v24 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2))))
	if v24 == int32(0) {
		v47 = v23
		v48 = v24
		goto L12
	} else {
		goto L13
	}
L11:
	;
	v54 = base.B2i32(v48-v47&int32(255) != int32(0))
	goto L9
L12:
	;
	goto L11
L13:
	;
	if v24 != v23&int32(255) {
		v47 = v23
		v48 = v24
		goto L12
	} else {
		goto L14
	}
L14:
	;
	v30 = l2
	v31 = l3
	goto L15
L15:
	;
	v34 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v31)+1)))
	v35 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v30)+1)))
	if v35 == int32(0) {
		v47 = v34
		v48 = v35
		goto L12
	} else {
		goto L17
	}
L16:
	;
	v47 = v34
	v48 = v35
	goto L12
L17:
	;
	v38 = int32(1)
	if v35 == v34&int32(255) {
		v30 = v30 + v38
		v31 = v31 + v38
		goto L15
	} else {
		goto L18
	}
L18:
	;
	goto L16
L19:
	;
	v59 = F_sdscatlen(m, v55, int32(_a66), int32(1))
	mBase = m.M
	v60 = m.ExcPending
	if v60 != 0 {
		goto L4
	} else {
		goto L20
	}
L20:
	;
	v64 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2+int32(-1)))))
	switch v64 & int32(7) {
	case 0:
		goto L26
	case 1:
		goto L25
	case 2:
		goto L24
	case 3:
		goto L23
	case 4:
		goto L22
	default:
		v81 = int32(0)
		goto L21
	}
L21:
	;
	v82 = F_sdscatrepr(m, v59, l2, v81)
	mBase = m.M
	v83 = m.ExcPending
	if v83 != 0 {
		goto L4
	} else {
		goto L27
	}
L22:
	;
	v80 = *(*int32)(unsafe.Add(mBase, uint32(l2+int32(-17))))
	v81 = v80
	goto L21
L23:
	;
	v77 = *(*int32)(unsafe.Add(mBase, uint32(l2+int32(-9))))
	v81 = v77
	goto L21
L24:
	;
	v74 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l2+int32(-5)))))
	v81 = v74
	goto L21
L25:
	;
	v71 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2+int32(-3)))))
	v81 = v71
	goto L21
L26:
	;
	v81 = int32(base.Ui32(v64) >> (uint(int32(3)) % 32))
	goto L21
L27:
	;
	v84 = F_rewriteConfigRewriteLine(m, l0, l1, v82, v54)
	mBase = m.M
	v85 = m.ExcPending
	if v85 != 0 {
		goto L4
	} else {
		goto L28
	}
L28:
	;
	goto L1
}
func F_rewriteConfigSentinelOption(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
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
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v104 int32
	_ = v104
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v112 int32
	_ = v112
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v121 int32
	_ = v121
	var v123 int32
	_ = v123
	var v125 int64
	_ = v125
	var v126 int64
	_ = v126
	var v127 int64
	_ = v127
	var v128 int64
	_ = v128
	var v129 int64
	_ = v129
	var v130 int64
	_ = v130
	var v131 int64
	_ = v131
	var v133 int64
	_ = v133
	var v135 int64
	_ = v135
	var v137 int64
	_ = v137
	var v139 int64
	_ = v139
	var v141 int64
	_ = v141
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v148 int32
	_ = v148
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v155 int32
	_ = v155
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v163 int32
	_ = v163
	var v167 int32
	_ = v167
	var v169 int32
	_ = v169
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v183 int32
	_ = v183
	var v188 int32
	_ = v188
	var v192 int32
	_ = v192
	var v196 int32
	_ = v196
	var v204 int32
	_ = v204
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v219 int32
	_ = v219
	var v222 int32
	_ = v222
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v252 int32
	_ = v252
	var v253 int32
	_ = v253
	var v254 int64
	_ = v254
	var v256 int64
	_ = v256
	var v258 int32
	_ = v258
	var v259 int32
	_ = v259
	var v260 int32
	_ = v260
	var v261 int64
	_ = v261
	var v268 int32
	_ = v268
	var v269 int32
	_ = v269
	var v271 int32
	_ = v271
	var v272 int32
	_ = v272
	var v275 int64
	_ = v275
	var v277 int64
	_ = v277
	var v279 int32
	_ = v279
	var v280 int32
	_ = v280
	var v281 int32
	_ = v281
	var v282 int64
	_ = v282
	var v289 int32
	_ = v289
	var v290 int32
	_ = v290
	var v292 int32
	_ = v292
	var v293 int32
	_ = v293
	var v296 int32
	_ = v296
	var v299 int32
	_ = v299
	var v300 int32
	_ = v300
	var v301 int32
	_ = v301
	var v302 int32
	_ = v302
	var v309 int32
	_ = v309
	var v310 int32
	_ = v310
	var v312 int32
	_ = v312
	var v313 int32
	_ = v313
	var v316 int32
	_ = v316
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
	var v329 int32
	_ = v329
	var v332 int32
	_ = v332
	var v339 int32
	_ = v339
	var v342 int32
	_ = v342
	var v345 int32
	_ = v345
	var v348 int32
	_ = v348
	var v349 int32
	_ = v349
	var v351 int32
	_ = v351
	var v352 int32
	_ = v352
	var v354 int32
	_ = v354
	var v355 int32
	_ = v355
	var v360 int32
	_ = v360
	var v363 int32
	_ = v363
	var v364 int32
	_ = v364
	var v365 int32
	_ = v365
	var v370 int32
	_ = v370
	var v371 int32
	_ = v371
	var v373 int32
	_ = v373
	var v376 int32
	_ = v376
	var v383 int32
	_ = v383
	var v386 int32
	_ = v386
	var v389 int32
	_ = v389
	var v392 int32
	_ = v392
	var v393 int32
	_ = v393
	var v395 int32
	_ = v395
	var v396 int32
	_ = v396
	var v398 int32
	_ = v398
	var v399 int32
	_ = v399
	var v404 int32
	_ = v404
	var v407 int32
	_ = v407
	var v408 int32
	_ = v408
	var v409 int32
	_ = v409
	var v414 int32
	_ = v414
	var v415 int32
	_ = v415
	var v417 int32
	_ = v417
	var v420 int32
	_ = v420
	var v427 int32
	_ = v427
	var v430 int32
	_ = v430
	var v433 int32
	_ = v433
	var v436 int32
	_ = v436
	var v437 int32
	_ = v437
	var v439 int32
	_ = v439
	var v440 int32
	_ = v440
	var v442 int32
	_ = v442
	var v443 int32
	_ = v443
	var v448 int32
	_ = v448
	var v451 int32
	_ = v451
	var v452 int32
	_ = v452
	var v453 int32
	_ = v453
	var v458 int32
	_ = v458
	var v459 int32
	_ = v459
	var v461 int32
	_ = v461
	var v464 int32
	_ = v464
	var v471 int32
	_ = v471
	var v474 int32
	_ = v474
	var v477 int32
	_ = v477
	var v480 int32
	_ = v480
	var v481 int32
	_ = v481
	var v483 int32
	_ = v483
	var v484 int32
	_ = v484
	var v486 int32
	_ = v486
	var v487 int32
	_ = v487
	var v492 int64
	_ = v492
	var v495 int32
	_ = v495
	var v496 int32
	_ = v496
	var v497 int32
	_ = v497
	var v498 int64
	_ = v498
	var v505 int32
	_ = v505
	var v506 int32
	_ = v506
	var v508 int32
	_ = v508
	var v509 int32
	_ = v509
	var v512 int32
	_ = v512
	var v513 int32
	_ = v513
	var v514 int32
	_ = v514
	var v515 int64
	_ = v515
	var v522 int32
	_ = v522
	var v523 int32
	_ = v523
	var v525 int32
	_ = v525
	var v526 int32
	_ = v526
	var v527 int32
	_ = v527
	var v528 int32
	_ = v528
	var v529 int32
	_ = v529
	var v530 int64
	_ = v530
	var v537 int32
	_ = v537
	var v538 int32
	_ = v538
	var v540 int32
	_ = v540
	var v541 int32
	_ = v541
	var v542 int32
	_ = v542
	var v543 int32
	_ = v543
	var v544 int32
	_ = v544
	var v563 int32
	_ = v563
	var v564 int32
	_ = v564
	var v567 int32
	_ = v567
	var v570 int32
	_ = v570
	var v571 int32
	_ = v571
	var v575 int32
	_ = v575
	var v579 int32
	_ = v579
	var v580 int32
	_ = v580
	var v581 int32
	_ = v581
	var v584 int32
	_ = v584
	var v586 int32
	_ = v586
	var v588 int64
	_ = v588
	var v589 int64
	_ = v589
	var v590 int64
	_ = v590
	var v591 int64
	_ = v591
	var v592 int64
	_ = v592
	var v593 int64
	_ = v593
	var v594 int64
	_ = v594
	var v596 int64
	_ = v596
	var v598 int64
	_ = v598
	var v600 int64
	_ = v600
	var v602 int64
	_ = v602
	var v604 int64
	_ = v604
	var v606 int32
	_ = v606
	var v607 int32
	_ = v607
	var v608 int32
	_ = v608
	var v611 int32
	_ = v611
	var v614 int32
	_ = v614
	var v615 int32
	_ = v615
	var v616 int32
	_ = v616
	var v618 int32
	_ = v618
	var v621 int32
	_ = v621
	var v622 int32
	_ = v622
	var v626 int32
	_ = v626
	var v630 int32
	_ = v630
	var v632 int32
	_ = v632
	var v639 int32
	_ = v639
	var v640 int32
	_ = v640
	var v641 int32
	_ = v641
	var v646 int32
	_ = v646
	var v651 int32
	_ = v651
	var v655 int32
	_ = v655
	var v659 int32
	_ = v659
	var v665 int32
	_ = v665
	var v666 int32
	_ = v666
	var v667 int32
	_ = v667
	var v668 int32
	_ = v668
	var v670 int32
	_ = v670
	var v671 int32
	_ = v671
	var v674 int32
	_ = v674
	var v675 int32
	_ = v675
	var v681 int32
	_ = v681
	var v682 int32
	_ = v682
	var v685 int32
	_ = v685
	var v686 int32
	_ = v686
	var v689 int32
	_ = v689
	var v698 int32
	_ = v698
	var v699 int32
	_ = v699
	var v705 int32
	_ = v705
	var v706 int32
	_ = v706
	var v709 int32
	_ = v709
	var v711 int32
	_ = v711
	var v712 int32
	_ = v712
	var v713 int32
	_ = v713
	var v715 int32
	_ = v715
	var v719 int32
	_ = v719
	var v721 int32
	_ = v721
	var v722 int32
	_ = v722
	var v723 int32
	_ = v723
	var v725 int32
	_ = v725
	var v726 int32
	_ = v726
	var v727 int32
	_ = v727
	var v728 int32
	_ = v728
	var v733 int32
	_ = v733
	var v734 int32
	_ = v734
	var v737 int32
	_ = v737
	var v738 int32
	_ = v738
	var v741 int32
	_ = v741
	var v742 int32
	_ = v742
	var v743 int32
	_ = v743
	var v745 int32
	_ = v745
	var v746 int32
	_ = v746
	var v747 int32
	_ = v747
	var v748 int32
	_ = v748
	var v749 int32
	_ = v749
	var v750 int32
	_ = v750
	var v756 int32
	_ = v756
	var v757 int32
	_ = v757
	var v758 int32
	_ = v758
	var v766 int32
	_ = v766
	var v767 int32
	_ = v767
	var v768 int32
	_ = v768
	var v769 int32
	_ = v769
	var v771 int32
	_ = v771
	var v772 int32
	_ = v772
	var v775 int32
	_ = v775
	var v776 int32
	_ = v776
	var v778 int32
	_ = v778
	var v780 int32
	_ = v780
	var v781 int32
	_ = v781
	var v782 int32
	_ = v782
	var v783 int32
	_ = v783
	var v802 int32
	_ = v802
	var v803 int32
	_ = v803
	var v806 int32
	_ = v806
	var v809 int32
	_ = v809
	var v810 int32
	_ = v810
	var v814 int32
	_ = v814
	var v818 int32
	_ = v818
	var v819 int32
	_ = v819
	var v820 int32
	_ = v820
	var v823 int32
	_ = v823
	var v825 int32
	_ = v825
	var v827 int64
	_ = v827
	var v828 int64
	_ = v828
	var v829 int64
	_ = v829
	var v830 int64
	_ = v830
	var v831 int64
	_ = v831
	var v832 int64
	_ = v832
	var v833 int64
	_ = v833
	var v835 int64
	_ = v835
	var v837 int64
	_ = v837
	var v839 int64
	_ = v839
	var v841 int64
	_ = v841
	var v843 int64
	_ = v843
	var v845 int32
	_ = v845
	var v846 int32
	_ = v846
	var v847 int32
	_ = v847
	var v850 int32
	_ = v850
	var v853 int32
	_ = v853
	var v854 int32
	_ = v854
	var v855 int32
	_ = v855
	var v857 int32
	_ = v857
	var v860 int32
	_ = v860
	var v861 int32
	_ = v861
	var v865 int32
	_ = v865
	var v869 int32
	_ = v869
	var v871 int32
	_ = v871
	var v878 int32
	_ = v878
	var v879 int32
	_ = v879
	var v880 int32
	_ = v880
	var v885 int32
	_ = v885
	var v890 int32
	_ = v890
	var v894 int32
	_ = v894
	var v898 int32
	_ = v898
	var v904 int32
	_ = v904
	var v905 int32
	_ = v905
	var v909 int32
	_ = v909
	var v910 int32
	_ = v910
	var v911 int32
	_ = v911
	var v912 int32
	_ = v912
	var v913 int32
	_ = v913
	var v919 int32
	_ = v919
	var v920 int32
	_ = v920
	var v921 int32
	_ = v921
	var v928 int32
	_ = v928
	var v929 int32
	_ = v929
	var v930 int32
	_ = v930
	var v933 int32
	_ = v933
	var v940 int32
	_ = v940
	var v943 int32
	_ = v943
	var v946 int32
	_ = v946
	var v949 int32
	_ = v949
	var v950 int32
	_ = v950
	var v952 int32
	_ = v952
	var v953 int32
	_ = v953
	var v955 int32
	_ = v955
	var v956 int32
	_ = v956
	var v958 int32
	_ = v958
	var v959 int32
	_ = v959
	var v960 int32
	_ = v960
	var v961 int32
	_ = v961
	var v980 int32
	_ = v980
	var v981 int32
	_ = v981
	var v984 int32
	_ = v984
	var v987 int32
	_ = v987
	var v988 int32
	_ = v988
	var v992 int32
	_ = v992
	var v996 int32
	_ = v996
	var v997 int32
	_ = v997
	var v998 int32
	_ = v998
	var v1001 int32
	_ = v1001
	var v1003 int32
	_ = v1003
	var v1005 int64
	_ = v1005
	var v1006 int64
	_ = v1006
	var v1007 int64
	_ = v1007
	var v1008 int64
	_ = v1008
	var v1009 int64
	_ = v1009
	var v1010 int64
	_ = v1010
	var v1011 int64
	_ = v1011
	var v1013 int64
	_ = v1013
	var v1015 int64
	_ = v1015
	var v1017 int64
	_ = v1017
	var v1019 int64
	_ = v1019
	var v1021 int64
	_ = v1021
	var v1023 int32
	_ = v1023
	var v1024 int32
	_ = v1024
	var v1025 int32
	_ = v1025
	var v1028 int32
	_ = v1028
	var v1031 int32
	_ = v1031
	var v1032 int32
	_ = v1032
	var v1033 int32
	_ = v1033
	var v1035 int32
	_ = v1035
	var v1038 int32
	_ = v1038
	var v1039 int32
	_ = v1039
	var v1043 int32
	_ = v1043
	var v1047 int32
	_ = v1047
	var v1049 int32
	_ = v1049
	var v1056 int32
	_ = v1056
	var v1057 int32
	_ = v1057
	var v1058 int32
	_ = v1058
	var v1063 int32
	_ = v1063
	var v1068 int32
	_ = v1068
	var v1072 int32
	_ = v1072
	var v1076 int32
	_ = v1076
	var v1082 int32
	_ = v1082
	var v1083 int32
	_ = v1083
	var v1084 int32
	_ = v1084
	var v1085 int32
	_ = v1085
	var v1086 int32
	_ = v1086
	var v1091 int32
	_ = v1091
	var v1092 int32
	_ = v1092
	var v1093 int32
	_ = v1093
	var v1097 int32
	_ = v1097
	var v1104 int32
	_ = v1104
	var v1107 int32
	_ = v1107
	var v1110 int32
	_ = v1110
	var v1113 int32
	_ = v1113
	var v1114 int32
	_ = v1114
	var v1115 int32
	_ = v1115
	var v1116 int32
	_ = v1116
	var v1119 int32
	_ = v1119
	var v1120 int32
	_ = v1120
	var v1123 int32
	_ = v1123
	var v1130 int32
	_ = v1130
	var v1133 int32
	_ = v1133
	var v1136 int32
	_ = v1136
	var v1139 int32
	_ = v1139
	var v1140 int32
	_ = v1140
	var v1142 int32
	_ = v1142
	var v1143 int32
	_ = v1143
	var v1145 int32
	_ = v1145
	var v1146 int32
	_ = v1146
	var v1148 int32
	_ = v1148
	var v1156 int32
	_ = v1156
	var v1157 int32
	_ = v1157
	var v1160 int32
	_ = v1160
	var v1163 int32
	_ = v1163
	var v1164 int32
	_ = v1164
	var v1168 int32
	_ = v1168
	var v1172 int32
	_ = v1172
	var v1173 int32
	_ = v1173
	var v1174 int32
	_ = v1174
	var v1177 int32
	_ = v1177
	var v1179 int32
	_ = v1179
	var v1181 int64
	_ = v1181
	var v1182 int64
	_ = v1182
	var v1183 int64
	_ = v1183
	var v1184 int64
	_ = v1184
	var v1185 int64
	_ = v1185
	var v1186 int64
	_ = v1186
	var v1187 int64
	_ = v1187
	var v1189 int64
	_ = v1189
	var v1191 int64
	_ = v1191
	var v1193 int64
	_ = v1193
	var v1195 int64
	_ = v1195
	var v1197 int64
	_ = v1197
	var v1199 int32
	_ = v1199
	var v1200 int32
	_ = v1200
	var v1201 int32
	_ = v1201
	var v1204 int32
	_ = v1204
	var v1207 int32
	_ = v1207
	var v1208 int32
	_ = v1208
	var v1209 int32
	_ = v1209
	var v1211 int32
	_ = v1211
	var v1214 int32
	_ = v1214
	var v1215 int32
	_ = v1215
	var v1219 int32
	_ = v1219
	var v1223 int32
	_ = v1223
	var v1225 int32
	_ = v1225
	var v1232 int32
	_ = v1232
	var v1233 int32
	_ = v1233
	var v1234 int32
	_ = v1234
	var v1239 int32
	_ = v1239
	var v1244 int32
	_ = v1244
	var v1248 int32
	_ = v1248
	var v1252 int32
	_ = v1252
	var v1267 int32
	_ = v1267
	var v1268 int32
	_ = v1268
	var v1270 int64
	_ = v1270
	var v1276 int32
	_ = v1276
	var v1277 int32
	_ = v1277
	var v1279 int32
	_ = v1279
	var v1280 int32
	_ = v1280
	var v1281 int32
	_ = v1281
	var v1282 int32
	_ = v1282
	var v1287 int32
	_ = v1287
	var v1288 int32
	_ = v1288
	var v1290 int32
	_ = v1290
	var v1293 int32
	_ = v1293
	var v1300 int32
	_ = v1300
	var v1303 int32
	_ = v1303
	var v1306 int32
	_ = v1306
	var v1309 int32
	_ = v1309
	var v1310 int32
	_ = v1310
	var v1312 int32
	_ = v1312
	var v1313 int32
	_ = v1313
	var v1315 int32
	_ = v1315
	var v1316 int32
	_ = v1316
	var v1319 int32
	_ = v1319
	var v1324 int32
	_ = v1324
	var v1325 int32
	_ = v1325
	var v1328 int32
	_ = v1328
	var v1329 int32
	_ = v1329
	var v1331 int32
	_ = v1331
	var v1335 int32
	_ = v1335
	var v1336 int32
	_ = v1336
	var v1338 int32
	_ = v1338
	var v1339 int32
	_ = v1339
	var v1342 int32
	_ = v1342
	var v1344 int32
	_ = v1344
	var v1346 int32
	_ = v1346
	var v1350 int32
	_ = v1350
	var v1351 int32
	_ = v1351
	var v1353 int32
	_ = v1353
	var v1356 int32
	_ = v1356
	var v1363 int32
	_ = v1363
	var v1366 int32
	_ = v1366
	var v1369 int32
	_ = v1369
	var v1372 int32
	_ = v1372
	var v1373 int32
	_ = v1373
	var v1375 int32
	_ = v1375
	var v1376 int32
	_ = v1376
	var v1378 int32
	_ = v1378
	var v1379 int32
	_ = v1379
	var v1382 int32
	_ = v1382
	var v1387 int32
	_ = v1387
	var v1389 int32
	_ = v1389
	var v1393 int32
	_ = v1393
	var v1394 int32
	_ = v1394
	var v1396 int32
	_ = v1396
	var v1399 int32
	_ = v1399
	var v1406 int32
	_ = v1406
	var v1409 int32
	_ = v1409
	var v1412 int32
	_ = v1412
	var v1415 int32
	_ = v1415
	var v1416 int32
	_ = v1416
	var v1418 int32
	_ = v1418
	var v1419 int32
	_ = v1419
	var v1421 int32
	_ = v1421
	var v1422 int32
	_ = v1422
	var v1425 int32
	_ = v1425
	var v1431 int32
	_ = v1431
	var v1434 int32
	_ = v1434
	var v1437 int32
	_ = v1437
	var v1440 int32
	_ = v1440
	var v1443 int32
	_ = v1443
	var v1446 int32
	_ = v1446
	var v1449 int32
	_ = v1449
	var v1452 int32
	_ = v1452
	var v1455 int32
	_ = v1455
	var v1458 int32
	_ = v1458
	var v1461 int32
	_ = v1461
	var v1464 int32
	_ = v1464
	var v1467 int32
	_ = v1467
	var v1470 int32
	_ = v1470
	var v1473 int32
	_ = v1473
	v12 = m.G0
	v14 = v12 - int32(320)
	m.G0 = v14
	v16 = F_sdsempty(m)
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+304)) = int32(_a1336)
	v24 = F_sdscatprintf(m, v16, int32(_a1364), v14+int32(304))
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v27 = F_rewriteConfigRewriteLine(m, l0, int32(_a1365), v24, int32(1))
	mBase = m.M
	v28 = m.ExcPending
	if v28 != 0 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v29 = F_sdsempty(m)
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		goto L1
	} else {
		goto L5
	}
L5:
	;
	v34 = *(*int32)(unsafe.Add(mBase, _consts[746]))
	if v34 != 0 {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	v35 = int32(_a483)
	goto L8
L7:
	;
	v35 = int32(_a484)
	goto L8
L8:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+288)) = v35
	v41 = F_sdscatprintf(m, v29, int32(_a1366), v14+int32(288))
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L1
	} else {
		goto L9
	}
L9:
	;
	v44 = *(*int32)(unsafe.Add(mBase, _consts[746]))
	v47 = F_rewriteConfigRewriteLine(m, l0, int32(_a1367), v41, base.B2i32(v44 != int32(1)))
	mBase = m.M
	v48 = m.ExcPending
	if v48 != 0 {
		goto L1
	} else {
		goto L10
	}
L10:
	;
	v49 = F_sdsempty(m)
	mBase = m.M
	v50 = m.ExcPending
	if v50 != 0 {
		goto L1
	} else {
		goto L11
	}
L11:
	;
	v54 = *(*int32)(unsafe.Add(mBase, _consts[729]))
	if v54 != 0 {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	v55 = int32(_a483)
	goto L14
L13:
	;
	v55 = int32(_a484)
	goto L14
L14:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+272)) = v55
	v61 = F_sdscatprintf(m, v49, int32(_a1368), v14+int32(272))
	mBase = m.M
	v62 = m.ExcPending
	if v62 != 0 {
		goto L1
	} else {
		goto L15
	}
L15:
	;
	v63 = int32(0)
	v64 = *(*int32)(unsafe.Add(mBase, _consts[729]))
	v67 = F_rewriteConfigRewriteLine(m, l0, int32(_a1369), v61, base.B2i32(v64 != v63))
	mBase = m.M
	v68 = m.ExcPending
	if v68 != 0 {
		goto L1
	} else {
		goto L16
	}
L16:
	;
	v69 = F_sdsempty(m)
	mBase = m.M
	v70 = m.ExcPending
	if v70 != 0 {
		goto L1
	} else {
		goto L17
	}
L17:
	;
	v74 = *(*int32)(unsafe.Add(mBase, _consts[734]))
	if v74 != 0 {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	v75 = int32(_a483)
	goto L20
L19:
	;
	v75 = int32(_a484)
	goto L20
L20:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+256)) = v75
	v81 = F_sdscatprintf(m, v69, int32(_a1370), v14+int32(256))
	mBase = m.M
	v82 = m.ExcPending
	if v82 != 0 {
		goto L1
	} else {
		goto L21
	}
L21:
	;
	v83 = int32(0)
	v84 = *(*int32)(unsafe.Add(mBase, _consts[734]))
	v87 = F_rewriteConfigRewriteLine(m, l0, int32(_a1371), v81, base.B2i32(v84 != v83))
	mBase = m.M
	v88 = m.ExcPending
	if v88 != 0 {
		goto L1
	} else {
		goto L22
	}
L22:
	;
	v90 = *(*int32)(unsafe.Add(mBase, _consts[733]))
	v91 = F_dictGetIterator(m, v90)
	mBase = m.M
	v92 = m.ExcPending
	if v92 != 0 {
		goto L1
	} else {
		goto L24
	}
L23:
	;
	v1267 = F_sdsempty(m)
	mBase = m.M
	v1268 = m.ExcPending
	if v1268 != 0 {
		goto L1
	} else {
		goto L326
	}
L24:
	;
	v100 = v91 + int32(20)
	v101 = *(*int32)(unsafe.Add(mBase, uint32(v91)+16))
	if v101 != 0 {
		goto L27
	} else {
		goto L28
	}
L25:
	;
	if v196 == int32(0) {
		goto L23
	} else {
		goto L51
	}
L26:
	;
	v107 = v100
	v108 = v104
	goto L29
L27:
	;
	v104 = int32(1)
	goto L26
L28:
	;
	v104 = int32(0)
	goto L26
L29:
	;
	switch v108 {
	case 0:
		goto L34
	default:
		goto L33
	}
L31:
	;
	v108 = int32(0)
	goto L29
L32:
	;
	goto L25
L33:
	;
	v188 = *(*int32)(unsafe.Add(mBase, uint32(v107)))
	*(*int32)(unsafe.Add(mBase, uint32(v91)+16)) = v188
	if v188 == int32(0) {
		goto L31
	} else {
		goto L50
	}
L34:
	;
	v112 = *(*int32)(unsafe.Add(mBase, uint32(v91)+4))
	if v112 != int32(-1) {
		v151 = v112
		goto L35
	} else {
		goto L36
	}
L35:
	;
	v152 = int32(1)
	v153 = v151 + v152
	*(*int32)(unsafe.Add(mBase, uint32(v91)+4)) = v153
	v155 = int32(0)
	v158 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
	v159 = *(*int32)(unsafe.Add(mBase, uint32(v91)+8))
	v163 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v158+v159+int32(26)))))
	if v163 == int32(255) {
		goto L44
	} else {
		goto L45
	}
L36:
	;
	v116 = *(*int32)(unsafe.Add(mBase, uint32(v91)+8))
	if v116 != 0 {
		v151 = int32(-1)
		goto L35
	} else {
		goto L37
	}
L37:
	;
	v117 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
	v118 = *(*int32)(unsafe.Add(mBase, uint32(v91)+12))
	if v118 == int32(0) {
		goto L39
	} else {
		goto L40
	}
L38:
	;
	v145 = *(*int32)(unsafe.Add(mBase, uint32(v144)+20))
	if v145 != int32(-1) {
		goto L41
	} else {
		goto L42
	}
L39:
	;
	v125 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v117)+16)))
	v126 = int64(*(*int8)(unsafe.Add(mBase, uint32(v117)+27)))
	v127 = int64(*(*int32)(unsafe.Add(mBase, uint32(v117)+8)))
	v128 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v117)+12)))
	v129 = int64(*(*int8)(unsafe.Add(mBase, uint32(v117)+26)))
	v130 = int64(*(*int32)(unsafe.Add(mBase, uint32(v117)+4)))
	v131 = F_wangHash64(m, v130)
	mBase = m.M
	v133 = F_wangHash64(m, v129+v131)
	mBase = m.M
	v135 = F_wangHash64(m, v128+v133)
	mBase = m.M
	v137 = F_wangHash64(m, v127+v135)
	mBase = m.M
	v139 = F_wangHash64(m, v126+v137)
	mBase = m.M
	v141 = F_wangHash64(m, v125+v139)
	mBase = m.M
	*(*int64)(unsafe.Add(mBase, uint32(v91)+24)) = v141
	v143 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
	v144 = v143
	goto L38
L40:
	;
	v121 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v117)+24)))
	v123 = v121 + int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v117)+24)) = uint16(v123)
	v144 = v117
	goto L38
L41:
	;
	v151 = v145 + int32(-1)
	goto L35
L42:
	;
	v148 = *(*int32)(unsafe.Add(mBase, uint32(v91)+4))
	v151 = v148
	goto L35
L43:
	;
	v178 = int32(2)
	v183 = *(*int32)(unsafe.Add(mBase, uint32(v158+v176<<(uint(v178)%32)+int32(4))))
	v107 = v183 + v177<<(uint(v178)%32)
	v108 = int32(1)
	goto L29
L44:
	;
	v167 = v155
	goto L46
L45:
	;
	v167 = v152 << (uint(v163) % 32)
	goto L46
L46:
	;
	if v153 < v167 {
		v176 = v159
		v177 = v153
		goto L43
	} else {
		goto L47
	}
L47:
	;
	if v159 != 0 {
		v196 = v155
		goto L32
	} else {
		goto L48
	}
L48:
	;
	v169 = *(*int32)(unsafe.Add(mBase, uint32(v158)+20))
	if v169 == int32(-1) {
		v196 = v155
		goto L32
	} else {
		goto L49
	}
L49:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v91)+4)) = int64(4294967296)
	v176 = int32(1)
	v177 = int32(0)
	goto L43
L50:
	;
	v192 = *(*int32)(unsafe.Add(mBase, uint32(v188)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v100))) = v192
	v196 = v188
	goto L32
L51:
	;
	v204 = v196
	goto L52
L52:
	;
	v213 = *(*int32)(unsafe.Add(mBase, uint32(v204)+8))
	goto L56
L53:
	;
	goto L23
L54:
	;
	v227 = *(*int32)(unsafe.Add(mBase, uint32(v226)+24))
	v228 = F_sdsempty(m)
	mBase = m.M
	v229 = m.ExcPending
	if v229 != 0 {
		goto L1
	} else {
		goto L60
	}
L55:
	;
	v226 = v213
	goto L54
L56:
	;
	v214 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v213))))
	if v214&int32(64) == int32(0) {
		goto L55
	} else {
		goto L57
	}
L57:
	;
	v219 = *(*int32)(unsafe.Add(mBase, uint32(v213)+280))
	if v219 == int32(0) {
		goto L55
	} else {
		goto L58
	}
L58:
	;
	v222 = *(*int32)(unsafe.Add(mBase, uint32(v213)+240))
	if int32(4) < v222 {
		v226 = v219
		goto L54
	} else {
		goto L59
	}
L59:
	;
	goto L55
L60:
	;
	v230 = int32(0)
	v231 = *(*int32)(unsafe.Add(mBase, _consts[734]))
	v237 = *(*int32)(unsafe.Add(mBase, uint32(v227+base.B2i32(v231 == v230)<<(uint(int32(2))%32))))
	v238 = *(*int32)(unsafe.Add(mBase, uint32(v213)+152))
	v239 = *(*int32)(unsafe.Add(mBase, uint32(v227)+8))
	v240 = *(*int32)(unsafe.Add(mBase, uint32(v213)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v14)+240)) = v240
	*(*int32)(unsafe.Add(mBase, uint32(v14)+244)) = v237
	*(*int32)(unsafe.Add(mBase, uint32(v14)+248)) = v239
	*(*int32)(unsafe.Add(mBase, uint32(v14)+252)) = v238
	v249 = F_sdscatprintf(m, v228, int32(_a1372), v14+int32(240))
	mBase = m.M
	v250 = m.ExcPending
	if v250 != 0 {
		goto L1
	} else {
		goto L61
	}
L61:
	;
	v252 = F_rewriteConfigRewriteLine(m, l0, int32(_a1373), v249, int32(1))
	mBase = m.M
	v253 = m.ExcPending
	if v253 != 0 {
		goto L1
	} else {
		goto L62
	}
L62:
	;
	v254 = *(*int64)(unsafe.Add(mBase, uint32(v213)+72))
	v256 = *(*int64)(unsafe.Add(mBase, _consts[747]))
	if v254 == v256 {
		goto L63
	} else {
		goto L64
	}
L63:
	;
	v275 = *(*int64)(unsafe.Add(mBase, uint32(v213)+264))
	v277 = *(*int64)(unsafe.Add(mBase, _consts[735]))
	if v275 == v277 {
		goto L68
	} else {
		goto L69
	}
L64:
	;
	v258 = F_sdsempty(m)
	mBase = m.M
	v259 = m.ExcPending
	if v259 != 0 {
		goto L1
	} else {
		goto L65
	}
L65:
	;
	v260 = *(*int32)(unsafe.Add(mBase, uint32(v213)+4))
	v261 = *(*int64)(unsafe.Add(mBase, uint32(v213)+72))
	*(*uint32)(unsafe.Add(mBase, uint32(v14)+228)) = uint32(v261)
	*(*int32)(unsafe.Add(mBase, uint32(v14)+224)) = v260
	v268 = F_sdscatprintf(m, v258, int32(_a1374), v14+int32(224))
	mBase = m.M
	v269 = m.ExcPending
	if v269 != 0 {
		goto L1
	} else {
		goto L66
	}
L66:
	;
	v271 = F_rewriteConfigRewriteLine(m, l0, int32(_a1375), v268, int32(1))
	mBase = m.M
	v272 = m.ExcPending
	if v272 != 0 {
		goto L1
	} else {
		goto L67
	}
L67:
	;
	goto L63
L68:
	;
	v296 = *(*int32)(unsafe.Add(mBase, uint32(v213)+156))
	if v296 == int32(1) {
		goto L73
	} else {
		goto L74
	}
L69:
	;
	v279 = F_sdsempty(m)
	mBase = m.M
	v280 = m.ExcPending
	if v280 != 0 {
		goto L1
	} else {
		goto L70
	}
L70:
	;
	v281 = *(*int32)(unsafe.Add(mBase, uint32(v213)+4))
	v282 = *(*int64)(unsafe.Add(mBase, uint32(v213)+264))
	*(*uint32)(unsafe.Add(mBase, uint32(v14)+212)) = uint32(v282)
	*(*int32)(unsafe.Add(mBase, uint32(v14)+208)) = v281
	v289 = F_sdscatprintf(m, v279, int32(_a1376), v14+int32(208))
	mBase = m.M
	v290 = m.ExcPending
	if v290 != 0 {
		goto L1
	} else {
		goto L71
	}
L71:
	;
	v292 = F_rewriteConfigRewriteLine(m, l0, int32(_a1377), v289, int32(1))
	mBase = m.M
	v293 = m.ExcPending
	if v293 != 0 {
		goto L1
	} else {
		goto L72
	}
L72:
	;
	goto L68
L73:
	;
	v316 = *(*int32)(unsafe.Add(mBase, uint32(v213)+284))
	if v316 == int32(0) {
		goto L78
	} else {
		goto L79
	}
L74:
	;
	v299 = F_sdsempty(m)
	mBase = m.M
	v300 = m.ExcPending
	if v300 != 0 {
		goto L1
	} else {
		goto L75
	}
L75:
	;
	v301 = *(*int32)(unsafe.Add(mBase, uint32(v213)+4))
	v302 = *(*int32)(unsafe.Add(mBase, uint32(v213)+156))
	*(*int32)(unsafe.Add(mBase, uint32(v14)+196)) = v302
	*(*int32)(unsafe.Add(mBase, uint32(v14)+192)) = v301
	v309 = F_sdscatprintf(m, v299, int32(_a1378), v14+int32(192))
	mBase = m.M
	v310 = m.ExcPending
	if v310 != 0 {
		goto L1
	} else {
		goto L76
	}
L76:
	;
	v312 = F_rewriteConfigRewriteLine(m, l0, int32(_a1379), v309, int32(1))
	mBase = m.M
	v313 = m.ExcPending
	if v313 != 0 {
		goto L1
	} else {
		goto L77
	}
L77:
	;
	goto L73
L78:
	;
	v360 = *(*int32)(unsafe.Add(mBase, uint32(v213)+288))
	if v360 == int32(0) {
		goto L90
	} else {
		goto L91
	}
L79:
	;
	v319 = F_sdsempty(m)
	mBase = m.M
	v320 = m.ExcPending
	if v320 != 0 {
		goto L1
	} else {
		goto L80
	}
L80:
	;
	v321 = *(*int32)(unsafe.Add(mBase, uint32(v213)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v14)+176)) = v321
	v326 = F_sdscatprintf(m, v319, int32(_a1380), v14+int32(176))
	mBase = m.M
	v327 = m.ExcPending
	if v327 != 0 {
		goto L1
	} else {
		goto L81
	}
L81:
	;
	v329 = *(*int32)(unsafe.Add(mBase, uint32(v213)+284))
	v332 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v329+int32(-1)))))
	switch v332 & int32(7) {
	case 0:
		goto L87
	case 1:
		goto L86
	case 2:
		goto L85
	case 3:
		goto L84
	case 4:
		goto L83
	default:
		v349 = int32(0)
		goto L82
	}
L82:
	;
	v351 = F_sdscatrepr(m, v326, v329, v349)
	mBase = m.M
	v352 = m.ExcPending
	if v352 != 0 {
		goto L1
	} else {
		goto L88
	}
L83:
	;
	v348 = *(*int32)(unsafe.Add(mBase, uint32(v329+int32(-17))))
	v349 = v348
	goto L82
L84:
	;
	v345 = *(*int32)(unsafe.Add(mBase, uint32(v329+int32(-9))))
	v349 = v345
	goto L82
L85:
	;
	v342 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v329+int32(-5)))))
	v349 = v342
	goto L82
L86:
	;
	v339 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v329+int32(-3)))))
	v349 = v339
	goto L82
L87:
	;
	v349 = int32(base.Ui32(v332) >> (uint(int32(3)) % 32))
	goto L82
L88:
	;
	v354 = F_rewriteConfigRewriteLine(m, l0, int32(_a1381), v351, int32(1))
	mBase = m.M
	v355 = m.ExcPending
	if v355 != 0 {
		goto L1
	} else {
		goto L89
	}
L89:
	;
	goto L78
L90:
	;
	v404 = *(*int32)(unsafe.Add(mBase, uint32(v213)+160))
	if v404 == int32(0) {
		goto L102
	} else {
		goto L103
	}
L91:
	;
	v363 = F_sdsempty(m)
	mBase = m.M
	v364 = m.ExcPending
	if v364 != 0 {
		goto L1
	} else {
		goto L92
	}
L92:
	;
	v365 = *(*int32)(unsafe.Add(mBase, uint32(v213)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v14)+160)) = v365
	v370 = F_sdscatprintf(m, v363, int32(_a1382), v14+int32(160))
	mBase = m.M
	v371 = m.ExcPending
	if v371 != 0 {
		goto L1
	} else {
		goto L93
	}
L93:
	;
	v373 = *(*int32)(unsafe.Add(mBase, uint32(v213)+288))
	v376 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v373+int32(-1)))))
	switch v376 & int32(7) {
	case 0:
		goto L99
	case 1:
		goto L98
	case 2:
		goto L97
	case 3:
		goto L96
	case 4:
		goto L95
	default:
		v393 = int32(0)
		goto L94
	}
L94:
	;
	v395 = F_sdscatrepr(m, v370, v373, v393)
	mBase = m.M
	v396 = m.ExcPending
	if v396 != 0 {
		goto L1
	} else {
		goto L100
	}
L95:
	;
	v392 = *(*int32)(unsafe.Add(mBase, uint32(v373+int32(-17))))
	v393 = v392
	goto L94
L96:
	;
	v389 = *(*int32)(unsafe.Add(mBase, uint32(v373+int32(-9))))
	v393 = v389
	goto L94
L97:
	;
	v386 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v373+int32(-5)))))
	v393 = v386
	goto L94
L98:
	;
	v383 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v373+int32(-3)))))
	v393 = v383
	goto L94
L99:
	;
	v393 = int32(base.Ui32(v376) >> (uint(int32(3)) % 32))
	goto L94
L100:
	;
	v398 = F_rewriteConfigRewriteLine(m, l0, int32(_a1383), v395, int32(1))
	mBase = m.M
	v399 = m.ExcPending
	if v399 != 0 {
		goto L1
	} else {
		goto L101
	}
L101:
	;
	goto L90
L102:
	;
	v448 = *(*int32)(unsafe.Add(mBase, uint32(v213)+164))
	if v448 == int32(0) {
		goto L114
	} else {
		goto L115
	}
L103:
	;
	v407 = F_sdsempty(m)
	mBase = m.M
	v408 = m.ExcPending
	if v408 != 0 {
		goto L1
	} else {
		goto L104
	}
L104:
	;
	v409 = *(*int32)(unsafe.Add(mBase, uint32(v213)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v14)+144)) = v409
	v414 = F_sdscatprintf(m, v407, int32(_a1384), v14+int32(144))
	mBase = m.M
	v415 = m.ExcPending
	if v415 != 0 {
		goto L1
	} else {
		goto L105
	}
L105:
	;
	v417 = *(*int32)(unsafe.Add(mBase, uint32(v213)+160))
	v420 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v417+int32(-1)))))
	switch v420 & int32(7) {
	case 0:
		goto L111
	case 1:
		goto L110
	case 2:
		goto L109
	case 3:
		goto L108
	case 4:
		goto L107
	default:
		v437 = int32(0)
		goto L106
	}
L106:
	;
	v439 = F_sdscatrepr(m, v414, v417, v437)
	mBase = m.M
	v440 = m.ExcPending
	if v440 != 0 {
		goto L1
	} else {
		goto L112
	}
L107:
	;
	v436 = *(*int32)(unsafe.Add(mBase, uint32(v417+int32(-17))))
	v437 = v436
	goto L106
L108:
	;
	v433 = *(*int32)(unsafe.Add(mBase, uint32(v417+int32(-9))))
	v437 = v433
	goto L106
L109:
	;
	v430 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v417+int32(-5)))))
	v437 = v430
	goto L106
L110:
	;
	v427 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v417+int32(-3)))))
	v437 = v427
	goto L106
L111:
	;
	v437 = int32(base.Ui32(v420) >> (uint(int32(3)) % 32))
	goto L106
L112:
	;
	v442 = F_rewriteConfigRewriteLine(m, l0, int32(_a1385), v439, int32(1))
	mBase = m.M
	v443 = m.ExcPending
	if v443 != 0 {
		goto L1
	} else {
		goto L113
	}
L113:
	;
	goto L102
L114:
	;
	v492 = *(*int64)(unsafe.Add(mBase, uint32(v213)+80))
	if v492 == int64(0) {
		goto L126
	} else {
		goto L127
	}
L115:
	;
	v451 = F_sdsempty(m)
	mBase = m.M
	v452 = m.ExcPending
	if v452 != 0 {
		goto L1
	} else {
		goto L116
	}
L116:
	;
	v453 = *(*int32)(unsafe.Add(mBase, uint32(v213)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v14)+128)) = v453
	v458 = F_sdscatprintf(m, v451, int32(_a1386), v14+int32(128))
	mBase = m.M
	v459 = m.ExcPending
	if v459 != 0 {
		goto L1
	} else {
		goto L117
	}
L117:
	;
	v461 = *(*int32)(unsafe.Add(mBase, uint32(v213)+164))
	v464 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v461+int32(-1)))))
	switch v464 & int32(7) {
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
		v481 = int32(0)
		goto L118
	}
L118:
	;
	v483 = F_sdscatrepr(m, v458, v461, v481)
	mBase = m.M
	v484 = m.ExcPending
	if v484 != 0 {
		goto L1
	} else {
		goto L124
	}
L119:
	;
	v480 = *(*int32)(unsafe.Add(mBase, uint32(v461+int32(-17))))
	v481 = v480
	goto L118
L120:
	;
	v477 = *(*int32)(unsafe.Add(mBase, uint32(v461+int32(-9))))
	v481 = v477
	goto L118
L121:
	;
	v474 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v461+int32(-5)))))
	v481 = v474
	goto L118
L122:
	;
	v471 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v461+int32(-3)))))
	v481 = v471
	goto L118
L123:
	;
	v481 = int32(base.Ui32(v464) >> (uint(int32(3)) % 32))
	goto L118
L124:
	;
	v486 = F_rewriteConfigRewriteLine(m, l0, int32(_a1387), v483, int32(1))
	mBase = m.M
	v487 = m.ExcPending
	if v487 != 0 {
		goto L1
	} else {
		goto L125
	}
L125:
	;
	goto L114
L126:
	;
	v512 = F_sdsempty(m)
	mBase = m.M
	v513 = m.ExcPending
	if v513 != 0 {
		goto L1
	} else {
		goto L131
	}
L127:
	;
	v495 = F_sdsempty(m)
	mBase = m.M
	v496 = m.ExcPending
	if v496 != 0 {
		goto L1
	} else {
		goto L128
	}
L128:
	;
	v497 = *(*int32)(unsafe.Add(mBase, uint32(v213)+4))
	v498 = *(*int64)(unsafe.Add(mBase, uint32(v213)+80))
	*(*uint32)(unsafe.Add(mBase, uint32(v14)+116)) = uint32(v498)
	*(*int32)(unsafe.Add(mBase, uint32(v14)+112)) = v497
	v505 = F_sdscatprintf(m, v495, int32(_a1388), v14+int32(112))
	mBase = m.M
	v506 = m.ExcPending
	if v506 != 0 {
		goto L1
	} else {
		goto L129
	}
L129:
	;
	v508 = F_rewriteConfigRewriteLine(m, l0, int32(_a1389), v505, int32(1))
	mBase = m.M
	v509 = m.ExcPending
	if v509 != 0 {
		goto L1
	} else {
		goto L130
	}
L130:
	;
	goto L126
L131:
	;
	v514 = *(*int32)(unsafe.Add(mBase, uint32(v213)+4))
	v515 = *(*int64)(unsafe.Add(mBase, uint32(v213)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v14)+104)) = v515
	*(*int32)(unsafe.Add(mBase, uint32(v14)+96)) = v514
	v522 = F_sdscatprintf(m, v512, int32(_a1390), v14+int32(96))
	mBase = m.M
	v523 = m.ExcPending
	if v523 != 0 {
		goto L1
	} else {
		goto L132
	}
L132:
	;
	v525 = F_rewriteConfigRewriteLine(m, l0, int32(_a1391), v522, int32(1))
	mBase = m.M
	v526 = m.ExcPending
	if v526 != 0 {
		goto L1
	} else {
		goto L133
	}
L133:
	;
	v527 = F_sdsempty(m)
	mBase = m.M
	v528 = m.ExcPending
	if v528 != 0 {
		goto L1
	} else {
		goto L134
	}
L134:
	;
	v529 = *(*int32)(unsafe.Add(mBase, uint32(v213)+4))
	v530 = *(*int64)(unsafe.Add(mBase, uint32(v213)+224))
	*(*int64)(unsafe.Add(mBase, uint32(v14)+88)) = v530
	*(*int32)(unsafe.Add(mBase, uint32(v14)+80)) = v529
	v537 = F_sdscatprintf(m, v527, int32(_a1392), v14+int32(80))
	mBase = m.M
	v538 = m.ExcPending
	if v538 != 0 {
		goto L1
	} else {
		goto L135
	}
L135:
	;
	v540 = F_rewriteConfigRewriteLine(m, l0, int32(_a1393), v537, int32(1))
	mBase = m.M
	v541 = m.ExcPending
	if v541 != 0 {
		goto L1
	} else {
		goto L136
	}
L136:
	;
	v542 = *(*int32)(unsafe.Add(mBase, uint32(v213)+148))
	v543 = F_dictGetIterator(m, v542)
	mBase = m.M
	v544 = m.ExcPending
	if v544 != 0 {
		goto L1
	} else {
		goto L137
	}
L137:
	;
	goto L139
L138:
	;
	F_dictReleaseIterator(m, v543)
	mBase = m.M
	v780 = m.ExcPending
	if v780 != 0 {
		goto L1
	} else {
		goto L202
	}
L139:
	;
	v563 = v543 + int32(20)
	v564 = *(*int32)(unsafe.Add(mBase, uint32(v543)+16))
	if v564 != 0 {
		goto L143
	} else {
		goto L144
	}
L141:
	;
	if v659 == int32(0) {
		goto L138
	} else {
		goto L167
	}
L142:
	;
	v570 = v563
	v571 = v567
	goto L145
L143:
	;
	v567 = int32(1)
	goto L142
L144:
	;
	v567 = int32(0)
	goto L142
L145:
	;
	switch v571 {
	case 0:
		goto L150
	default:
		goto L149
	}
L147:
	;
	v571 = int32(0)
	goto L145
L148:
	;
	goto L141
L149:
	;
	v651 = *(*int32)(unsafe.Add(mBase, uint32(v570)))
	*(*int32)(unsafe.Add(mBase, uint32(v543)+16)) = v651
	if v651 == int32(0) {
		goto L147
	} else {
		goto L166
	}
L150:
	;
	v575 = *(*int32)(unsafe.Add(mBase, uint32(v543)+4))
	if v575 != int32(-1) {
		v614 = v575
		goto L151
	} else {
		goto L152
	}
L151:
	;
	v615 = int32(1)
	v616 = v614 + v615
	*(*int32)(unsafe.Add(mBase, uint32(v543)+4)) = v616
	v618 = int32(0)
	v621 = *(*int32)(unsafe.Add(mBase, uint32(v543)))
	v622 = *(*int32)(unsafe.Add(mBase, uint32(v543)+8))
	v626 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v621+v622+int32(26)))))
	if v626 == int32(255) {
		goto L160
	} else {
		goto L161
	}
L152:
	;
	v579 = *(*int32)(unsafe.Add(mBase, uint32(v543)+8))
	if v579 != 0 {
		v614 = int32(-1)
		goto L151
	} else {
		goto L153
	}
L153:
	;
	v580 = *(*int32)(unsafe.Add(mBase, uint32(v543)))
	v581 = *(*int32)(unsafe.Add(mBase, uint32(v543)+12))
	if v581 == int32(0) {
		goto L155
	} else {
		goto L156
	}
L154:
	;
	v608 = *(*int32)(unsafe.Add(mBase, uint32(v607)+20))
	if v608 != int32(-1) {
		goto L157
	} else {
		goto L158
	}
L155:
	;
	v588 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v580)+16)))
	v589 = int64(*(*int8)(unsafe.Add(mBase, uint32(v580)+27)))
	v590 = int64(*(*int32)(unsafe.Add(mBase, uint32(v580)+8)))
	v591 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v580)+12)))
	v592 = int64(*(*int8)(unsafe.Add(mBase, uint32(v580)+26)))
	v593 = int64(*(*int32)(unsafe.Add(mBase, uint32(v580)+4)))
	v594 = F_wangHash64(m, v593)
	mBase = m.M
	v596 = F_wangHash64(m, v592+v594)
	mBase = m.M
	v598 = F_wangHash64(m, v591+v596)
	mBase = m.M
	v600 = F_wangHash64(m, v590+v598)
	mBase = m.M
	v602 = F_wangHash64(m, v589+v600)
	mBase = m.M
	v604 = F_wangHash64(m, v588+v602)
	mBase = m.M
	*(*int64)(unsafe.Add(mBase, uint32(v543)+24)) = v604
	v606 = *(*int32)(unsafe.Add(mBase, uint32(v543)))
	v607 = v606
	goto L154
L156:
	;
	v584 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v580)+24)))
	v586 = v584 + int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v580)+24)) = uint16(v586)
	v607 = v580
	goto L154
L157:
	;
	v614 = v608 + int32(-1)
	goto L151
L158:
	;
	v611 = *(*int32)(unsafe.Add(mBase, uint32(v543)+4))
	v614 = v611
	goto L151
L159:
	;
	v641 = int32(2)
	v646 = *(*int32)(unsafe.Add(mBase, uint32(v621+v639<<(uint(v641)%32)+int32(4))))
	v570 = v646 + v640<<(uint(v641)%32)
	v571 = int32(1)
	goto L145
L160:
	;
	v630 = v618
	goto L162
L161:
	;
	v630 = v615 << (uint(v626) % 32)
	goto L162
L162:
	;
	if v616 < v630 {
		v639 = v622
		v640 = v616
		goto L159
	} else {
		goto L163
	}
L163:
	;
	if v622 != 0 {
		v659 = v618
		goto L148
	} else {
		goto L164
	}
L164:
	;
	v632 = *(*int32)(unsafe.Add(mBase, uint32(v621)+20))
	if v632 == int32(-1) {
		v659 = v618
		goto L148
	} else {
		goto L165
	}
L165:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v543)+4)) = int64(4294967296)
	v639 = int32(1)
	v640 = int32(0)
	goto L159
L166:
	;
	v655 = *(*int32)(unsafe.Add(mBase, uint32(v651)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v563))) = v655
	v659 = v651
	goto L148
L167:
	;
	v665 = *(*int32)(unsafe.Add(mBase, uint32(v659)+8))
	goto L169
L168:
	;
	v747 = F_sdsempty(m)
	mBase = m.M
	v748 = m.ExcPending
	if v748 != 0 {
		goto L1
	} else {
		goto L194
	}
L169:
	;
	v666 = *(*int32)(unsafe.Add(mBase, uint32(v665)+24))
	v667 = *(*int32)(unsafe.Add(mBase, uint32(v666)+8))
	v668 = *(*int32)(unsafe.Add(mBase, uint32(v227)+8))
	if v667 != v668 {
		v746 = v666
		goto L168
	} else {
		goto L170
	}
L170:
	;
	v670 = *(*int32)(unsafe.Add(mBase, uint32(v666)+4))
	v671 = *(*int32)(unsafe.Add(mBase, uint32(v227)+4))
	v674 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v671))))
	v675 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v670))))
	if v675 == int32(0) {
		v698 = v674
		v699 = v675
		goto L173
	} else {
		goto L174
	}
L171:
	;
	v745 = *(*int32)(unsafe.Add(mBase, uint32(v213)+24))
	v746 = v745
	goto L168
L172:
	;
	if v699-v698&int32(255) == int32(0) {
		goto L171
	} else {
		goto L180
	}
L173:
	;
	goto L172
L174:
	;
	if v675 != v674&int32(255) {
		v698 = v674
		v699 = v675
		goto L173
	} else {
		goto L175
	}
L175:
	;
	v681 = v670
	v682 = v671
	goto L176
L176:
	;
	v685 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v682)+1)))
	v686 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v681)+1)))
	if v686 == int32(0) {
		v698 = v685
		v699 = v686
		goto L173
	} else {
		goto L178
	}
L177:
	;
	v698 = v685
	v699 = v686
	goto L173
L178:
	;
	v689 = int32(1)
	if v686 == v685&int32(255) {
		v681 = v681 + v689
		v682 = v682 + v689
		goto L176
	} else {
		goto L179
	}
L179:
	;
	goto L177
L180:
	;
	v705 = *(*int32)(unsafe.Add(mBase, uint32(v666)))
	v706 = *(*int32)(unsafe.Add(mBase, uint32(v227)))
	v709 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v705))))
	if v709 != 0 {
		goto L183
	} else {
		goto L184
	}
L181:
	;
	if v741-v743 != 0 {
		v746 = v666
		goto L168
	} else {
		goto L193
	}
L182:
	;
	v741 = F_tolower(m, v737)
	mBase = m.M
	v742 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v738))))
	v743 = F_tolower(m, v742)
	mBase = m.M
	goto L181
L183:
	;
	v711 = v705
	v712 = v706
	v713 = v709
	goto L186
L184:
	;
	v737 = int32(0)
	v738 = v706
	goto L182
L185:
	;
	v737 = v734 & int32(255)
	v738 = v733
	goto L182
L186:
	;
	v715 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v712))))
	if v715 == int32(0) {
		v733 = v712
		v734 = v713
		goto L185
	} else {
		goto L188
	}
L187:
	;
	v733 = v727
	v734 = int32(0)
	goto L185
L188:
	;
	v719 = v713 & int32(255)
	if v719 == v715 {
		goto L189
	} else {
		goto L190
	}
L189:
	;
	v726 = int32(1)
	v727 = v712 + v726
	v728 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v711)+1)))
	if v728 != 0 {
		v711 = v711 + v726
		v712 = v727
		v713 = v728
		goto L186
	} else {
		goto L192
	}
L190:
	;
	v721 = F_tolower(m, v719)
	mBase = m.M
	v722 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v712))))
	v723 = F_tolower(m, v722)
	mBase = m.M
	if v721 == v723 {
		goto L189
	} else {
		goto L191
	}
L191:
	;
	v725 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v711))))
	v733 = v712
	v734 = v725
	goto L185
L192:
	;
	goto L187
L193:
	;
	goto L171
L194:
	;
	v749 = int32(0)
	v750 = *(*int32)(unsafe.Add(mBase, _consts[734]))
	v756 = *(*int32)(unsafe.Add(mBase, uint32(v746+base.B2i32(v750 == v749)<<(uint(int32(2))%32))))
	v757 = *(*int32)(unsafe.Add(mBase, uint32(v746)+8))
	v758 = *(*int32)(unsafe.Add(mBase, uint32(v213)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v14)+64)) = v758
	*(*int32)(unsafe.Add(mBase, uint32(v14)+68)) = v756
	*(*int32)(unsafe.Add(mBase, uint32(v14)+72)) = v757
	v766 = F_sdscatprintf(m, v747, int32(_a1394), v14+int32(64))
	mBase = m.M
	v767 = m.ExcPending
	if v767 != 0 {
		goto L1
	} else {
		goto L196
	}
L195:
	;
	F_sdsfree(m, v766)
	mBase = m.M
	v778 = m.ExcPending
	if v778 != 0 {
		goto L1
	} else {
		goto L201
	}
L196:
	;
	v768 = F_sdsdup(m, v766)
	mBase = m.M
	v769 = m.ExcPending
	if v769 != 0 {
		goto L1
	} else {
		goto L197
	}
L197:
	;
	v771 = F_rewriteConfigRewriteLine(m, l0, int32(_a1395), v768, int32(0))
	mBase = m.M
	v772 = m.ExcPending
	if v772 != 0 {
		goto L1
	} else {
		goto L198
	}
L198:
	;
	if v771 != 0 {
		goto L195
	} else {
		goto L199
	}
L199:
	;
	v775 = F_rewriteConfigRewriteLine(m, l0, int32(_a1396), v766, int32(1))
	mBase = m.M
	v776 = m.ExcPending
	if v776 != 0 {
		goto L1
	} else {
		goto L200
	}
L200:
	;
	goto L139
L201:
	;
	goto L139
L202:
	;
	v781 = *(*int32)(unsafe.Add(mBase, uint32(v213)+144))
	v782 = F_dictGetIterator(m, v781)
	mBase = m.M
	v783 = m.ExcPending
	if v783 != 0 {
		goto L1
	} else {
		goto L203
	}
L203:
	;
	goto L205
L204:
	;
	F_dictReleaseIterator(m, v782)
	mBase = m.M
	v958 = m.ExcPending
	if v958 != 0 {
		goto L1
	} else {
		goto L246
	}
L205:
	;
	v802 = v782 + int32(20)
	v803 = *(*int32)(unsafe.Add(mBase, uint32(v782)+16))
	if v803 != 0 {
		goto L209
	} else {
		goto L210
	}
L207:
	;
	if v898 == int32(0) {
		goto L204
	} else {
		goto L233
	}
L208:
	;
	v809 = v802
	v810 = v806
	goto L211
L209:
	;
	v806 = int32(1)
	goto L208
L210:
	;
	v806 = int32(0)
	goto L208
L211:
	;
	switch v810 {
	case 0:
		goto L216
	default:
		goto L215
	}
L213:
	;
	v810 = int32(0)
	goto L211
L214:
	;
	goto L207
L215:
	;
	v890 = *(*int32)(unsafe.Add(mBase, uint32(v809)))
	*(*int32)(unsafe.Add(mBase, uint32(v782)+16)) = v890
	if v890 == int32(0) {
		goto L213
	} else {
		goto L232
	}
L216:
	;
	v814 = *(*int32)(unsafe.Add(mBase, uint32(v782)+4))
	if v814 != int32(-1) {
		v853 = v814
		goto L217
	} else {
		goto L218
	}
L217:
	;
	v854 = int32(1)
	v855 = v853 + v854
	*(*int32)(unsafe.Add(mBase, uint32(v782)+4)) = v855
	v857 = int32(0)
	v860 = *(*int32)(unsafe.Add(mBase, uint32(v782)))
	v861 = *(*int32)(unsafe.Add(mBase, uint32(v782)+8))
	v865 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v860+v861+int32(26)))))
	if v865 == int32(255) {
		goto L226
	} else {
		goto L227
	}
L218:
	;
	v818 = *(*int32)(unsafe.Add(mBase, uint32(v782)+8))
	if v818 != 0 {
		v853 = int32(-1)
		goto L217
	} else {
		goto L219
	}
L219:
	;
	v819 = *(*int32)(unsafe.Add(mBase, uint32(v782)))
	v820 = *(*int32)(unsafe.Add(mBase, uint32(v782)+12))
	if v820 == int32(0) {
		goto L221
	} else {
		goto L222
	}
L220:
	;
	v847 = *(*int32)(unsafe.Add(mBase, uint32(v846)+20))
	if v847 != int32(-1) {
		goto L223
	} else {
		goto L224
	}
L221:
	;
	v827 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v819)+16)))
	v828 = int64(*(*int8)(unsafe.Add(mBase, uint32(v819)+27)))
	v829 = int64(*(*int32)(unsafe.Add(mBase, uint32(v819)+8)))
	v830 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v819)+12)))
	v831 = int64(*(*int8)(unsafe.Add(mBase, uint32(v819)+26)))
	v832 = int64(*(*int32)(unsafe.Add(mBase, uint32(v819)+4)))
	v833 = F_wangHash64(m, v832)
	mBase = m.M
	v835 = F_wangHash64(m, v831+v833)
	mBase = m.M
	v837 = F_wangHash64(m, v830+v835)
	mBase = m.M
	v839 = F_wangHash64(m, v829+v837)
	mBase = m.M
	v841 = F_wangHash64(m, v828+v839)
	mBase = m.M
	v843 = F_wangHash64(m, v827+v841)
	mBase = m.M
	*(*int64)(unsafe.Add(mBase, uint32(v782)+24)) = v843
	v845 = *(*int32)(unsafe.Add(mBase, uint32(v782)))
	v846 = v845
	goto L220
L222:
	;
	v823 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v819)+24)))
	v825 = v823 + int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v819)+24)) = uint16(v825)
	v846 = v819
	goto L220
L223:
	;
	v853 = v847 + int32(-1)
	goto L217
L224:
	;
	v850 = *(*int32)(unsafe.Add(mBase, uint32(v782)+4))
	v853 = v850
	goto L217
L225:
	;
	v880 = int32(2)
	v885 = *(*int32)(unsafe.Add(mBase, uint32(v860+v878<<(uint(v880)%32)+int32(4))))
	v809 = v885 + v879<<(uint(v880)%32)
	v810 = int32(1)
	goto L211
L226:
	;
	v869 = v857
	goto L228
L227:
	;
	v869 = v854 << (uint(v865) % 32)
	goto L228
L228:
	;
	if v855 < v869 {
		v878 = v861
		v879 = v855
		goto L225
	} else {
		goto L229
	}
L229:
	;
	if v861 != 0 {
		v898 = v857
		goto L214
	} else {
		goto L230
	}
L230:
	;
	v871 = *(*int32)(unsafe.Add(mBase, uint32(v860)+20))
	if v871 == int32(-1) {
		v898 = v857
		goto L214
	} else {
		goto L231
	}
L231:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v782)+4)) = int64(4294967296)
	v878 = int32(1)
	v879 = int32(0)
	goto L225
L232:
	;
	v894 = *(*int32)(unsafe.Add(mBase, uint32(v890)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v802))) = v894
	v898 = v890
	goto L214
L233:
	;
	v904 = *(*int32)(unsafe.Add(mBase, uint32(v898)+8))
	goto L234
L234:
	;
	v905 = *(*int32)(unsafe.Add(mBase, uint32(v904)+8))
	if v905 == int32(0) {
		goto L205
	} else {
		goto L235
	}
L235:
	;
	v909 = F_sdsempty(m)
	mBase = m.M
	v910 = m.ExcPending
	if v910 != 0 {
		goto L1
	} else {
		goto L236
	}
L236:
	;
	v911 = *(*int32)(unsafe.Add(mBase, uint32(v904)+24))
	v912 = int32(0)
	v913 = *(*int32)(unsafe.Add(mBase, _consts[734]))
	v919 = *(*int32)(unsafe.Add(mBase, uint32(v911+base.B2i32(v913 == v912)<<(uint(int32(2))%32))))
	v920 = *(*int32)(unsafe.Add(mBase, uint32(v911)+8))
	v921 = *(*int32)(unsafe.Add(mBase, uint32(v213)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v14)+48)) = v921
	*(*int32)(unsafe.Add(mBase, uint32(v14)+52)) = v919
	*(*int32)(unsafe.Add(mBase, uint32(v14)+56)) = v920
	v928 = F_sdscatprintf(m, v909, int32(_a1397), v14+int32(48))
	mBase = m.M
	v929 = m.ExcPending
	if v929 != 0 {
		goto L1
	} else {
		goto L237
	}
L237:
	;
	v930 = *(*int32)(unsafe.Add(mBase, uint32(v904)+8))
	v933 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v930+int32(-1)))))
	switch v933 & int32(7) {
	case 0:
		goto L243
	case 1:
		goto L242
	case 2:
		goto L241
	case 3:
		goto L240
	case 4:
		goto L239
	default:
		v950 = int32(0)
		goto L238
	}
L238:
	;
	v952 = F_sdscatrepr(m, v928, v930, v950)
	mBase = m.M
	v953 = m.ExcPending
	if v953 != 0 {
		goto L1
	} else {
		goto L244
	}
L239:
	;
	v949 = *(*int32)(unsafe.Add(mBase, uint32(v930+int32(-17))))
	v950 = v949
	goto L238
L240:
	;
	v946 = *(*int32)(unsafe.Add(mBase, uint32(v930+int32(-9))))
	v950 = v946
	goto L238
L241:
	;
	v943 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v930+int32(-5)))))
	v950 = v943
	goto L238
L242:
	;
	v940 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v930+int32(-3)))))
	v950 = v940
	goto L238
L243:
	;
	v950 = int32(base.Ui32(v933) >> (uint(int32(3)) % 32))
	goto L238
L244:
	;
	v955 = F_rewriteConfigRewriteLine(m, l0, int32(_a1398), v952, int32(1))
	mBase = m.M
	v956 = m.ExcPending
	if v956 != 0 {
		goto L1
	} else {
		goto L245
	}
L245:
	;
	goto L205
L246:
	;
	v959 = *(*int32)(unsafe.Add(mBase, uint32(v213)+104))
	v960 = F_dictGetIterator(m, v959)
	mBase = m.M
	v961 = m.ExcPending
	if v961 != 0 {
		goto L1
	} else {
		goto L247
	}
L247:
	;
	goto L249
L248:
	;
	F_dictReleaseIterator(m, v960)
	mBase = m.M
	v1148 = m.ExcPending
	if v1148 != 0 {
		goto L1
	} else {
		goto L298
	}
L249:
	;
	v980 = v960 + int32(20)
	v981 = *(*int32)(unsafe.Add(mBase, uint32(v960)+16))
	if v981 != 0 {
		goto L253
	} else {
		goto L254
	}
L251:
	;
	if v1076 == int32(0) {
		goto L248
	} else {
		goto L277
	}
L252:
	;
	v987 = v980
	v988 = v984
	goto L255
L253:
	;
	v984 = int32(1)
	goto L252
L254:
	;
	v984 = int32(0)
	goto L252
L255:
	;
	switch v988 {
	case 0:
		goto L260
	default:
		goto L259
	}
L257:
	;
	v988 = int32(0)
	goto L255
L258:
	;
	goto L251
L259:
	;
	v1068 = *(*int32)(unsafe.Add(mBase, uint32(v987)))
	*(*int32)(unsafe.Add(mBase, uint32(v960)+16)) = v1068
	if v1068 == int32(0) {
		goto L257
	} else {
		goto L276
	}
L260:
	;
	v992 = *(*int32)(unsafe.Add(mBase, uint32(v960)+4))
	if v992 != int32(-1) {
		v1031 = v992
		goto L261
	} else {
		goto L262
	}
L261:
	;
	v1032 = int32(1)
	v1033 = v1031 + v1032
	*(*int32)(unsafe.Add(mBase, uint32(v960)+4)) = v1033
	v1035 = int32(0)
	v1038 = *(*int32)(unsafe.Add(mBase, uint32(v960)))
	v1039 = *(*int32)(unsafe.Add(mBase, uint32(v960)+8))
	v1043 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1038+v1039+int32(26)))))
	if v1043 == int32(255) {
		goto L270
	} else {
		goto L271
	}
L262:
	;
	v996 = *(*int32)(unsafe.Add(mBase, uint32(v960)+8))
	if v996 != 0 {
		v1031 = int32(-1)
		goto L261
	} else {
		goto L263
	}
L263:
	;
	v997 = *(*int32)(unsafe.Add(mBase, uint32(v960)))
	v998 = *(*int32)(unsafe.Add(mBase, uint32(v960)+12))
	if v998 == int32(0) {
		goto L265
	} else {
		goto L266
	}
L264:
	;
	v1025 = *(*int32)(unsafe.Add(mBase, uint32(v1024)+20))
	if v1025 != int32(-1) {
		goto L267
	} else {
		goto L268
	}
L265:
	;
	v1005 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v997)+16)))
	v1006 = int64(*(*int8)(unsafe.Add(mBase, uint32(v997)+27)))
	v1007 = int64(*(*int32)(unsafe.Add(mBase, uint32(v997)+8)))
	v1008 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v997)+12)))
	v1009 = int64(*(*int8)(unsafe.Add(mBase, uint32(v997)+26)))
	v1010 = int64(*(*int32)(unsafe.Add(mBase, uint32(v997)+4)))
	v1011 = F_wangHash64(m, v1010)
	mBase = m.M
	v1013 = F_wangHash64(m, v1009+v1011)
	mBase = m.M
	v1015 = F_wangHash64(m, v1008+v1013)
	mBase = m.M
	v1017 = F_wangHash64(m, v1007+v1015)
	mBase = m.M
	v1019 = F_wangHash64(m, v1006+v1017)
	mBase = m.M
	v1021 = F_wangHash64(m, v1005+v1019)
	mBase = m.M
	*(*int64)(unsafe.Add(mBase, uint32(v960)+24)) = v1021
	v1023 = *(*int32)(unsafe.Add(mBase, uint32(v960)))
	v1024 = v1023
	goto L264
L266:
	;
	v1001 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v997)+24)))
	v1003 = v1001 + int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v997)+24)) = uint16(v1003)
	v1024 = v997
	goto L264
L267:
	;
	v1031 = v1025 + int32(-1)
	goto L261
L268:
	;
	v1028 = *(*int32)(unsafe.Add(mBase, uint32(v960)+4))
	v1031 = v1028
	goto L261
L269:
	;
	v1058 = int32(2)
	v1063 = *(*int32)(unsafe.Add(mBase, uint32(v1038+v1056<<(uint(v1058)%32)+int32(4))))
	v987 = v1063 + v1057<<(uint(v1058)%32)
	v988 = int32(1)
	goto L255
L270:
	;
	v1047 = v1035
	goto L272
L271:
	;
	v1047 = v1032 << (uint(v1043) % 32)
	goto L272
L272:
	;
	if v1033 < v1047 {
		v1056 = v1039
		v1057 = v1033
		goto L269
	} else {
		goto L273
	}
L273:
	;
	if v1039 != 0 {
		v1076 = v1035
		goto L258
	} else {
		goto L274
	}
L274:
	;
	v1049 = *(*int32)(unsafe.Add(mBase, uint32(v1038)+20))
	if v1049 == int32(-1) {
		v1076 = v1035
		goto L258
	} else {
		goto L275
	}
L275:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v960)+4)) = int64(4294967296)
	v1056 = int32(1)
	v1057 = int32(0)
	goto L269
L276:
	;
	v1072 = *(*int32)(unsafe.Add(mBase, uint32(v1068)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v980))) = v1072
	v1076 = v1068
	goto L258
L277:
	;
	v1082 = *(*int32)(unsafe.Add(mBase, uint32(v1076)))
	goto L278
L278:
	;
	v1083 = *(*int32)(unsafe.Add(mBase, uint32(v1076)+8))
	goto L279
L279:
	;
	v1084 = F_sdsempty(m)
	mBase = m.M
	v1085 = m.ExcPending
	if v1085 != 0 {
		goto L1
	} else {
		goto L280
	}
L280:
	;
	v1086 = *(*int32)(unsafe.Add(mBase, uint32(v213)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v14)+32)) = v1086
	v1091 = F_sdscatprintf(m, v1084, int32(_a1399), v14+int32(32))
	mBase = m.M
	v1092 = m.ExcPending
	if v1092 != 0 {
		goto L1
	} else {
		goto L281
	}
L281:
	;
	v1093 = int32(0)
	v1097 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1082+int32(-1)))))
	switch v1097 & int32(7) {
	case 0:
		goto L287
	case 1:
		goto L286
	case 2:
		goto L285
	case 3:
		goto L284
	case 4:
		goto L283
	default:
		v1114 = v1093
		goto L282
	}
L282:
	;
	v1115 = F_sdscatrepr(m, v1091, v1082, v1114)
	mBase = m.M
	v1116 = m.ExcPending
	if v1116 != 0 {
		goto L1
	} else {
		goto L288
	}
L283:
	;
	v1113 = *(*int32)(unsafe.Add(mBase, uint32(v1082+int32(-17))))
	v1114 = v1113
	goto L282
L284:
	;
	v1110 = *(*int32)(unsafe.Add(mBase, uint32(v1082+int32(-9))))
	v1114 = v1110
	goto L282
L285:
	;
	v1107 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1082+int32(-5)))))
	v1114 = v1107
	goto L282
L286:
	;
	v1104 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1082+int32(-3)))))
	v1114 = v1104
	goto L282
L287:
	;
	v1114 = int32(base.Ui32(v1097) >> (uint(int32(3)) % 32))
	goto L282
L288:
	;
	v1119 = F_sdscatlen(m, v1115, int32(_a66), int32(1))
	mBase = m.M
	v1120 = m.ExcPending
	if v1120 != 0 {
		goto L1
	} else {
		goto L289
	}
L289:
	;
	v1123 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1083+int32(-1)))))
	switch v1123 & int32(7) {
	case 0:
		goto L295
	case 1:
		goto L294
	case 2:
		goto L293
	case 3:
		goto L292
	case 4:
		goto L291
	default:
		v1140 = v1093
		goto L290
	}
L290:
	;
	v1142 = F_sdscatrepr(m, v1119, v1083, v1140)
	mBase = m.M
	v1143 = m.ExcPending
	if v1143 != 0 {
		goto L1
	} else {
		goto L296
	}
L291:
	;
	v1139 = *(*int32)(unsafe.Add(mBase, uint32(v1083+int32(-17))))
	v1140 = v1139
	goto L290
L292:
	;
	v1136 = *(*int32)(unsafe.Add(mBase, uint32(v1083+int32(-9))))
	v1140 = v1136
	goto L290
L293:
	;
	v1133 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1083+int32(-5)))))
	v1140 = v1133
	goto L290
L294:
	;
	v1130 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1083+int32(-3)))))
	v1140 = v1130
	goto L290
L295:
	;
	v1140 = int32(base.Ui32(v1123) >> (uint(int32(3)) % 32))
	goto L290
L296:
	;
	v1145 = F_rewriteConfigRewriteLine(m, l0, int32(_a1400), v1142, int32(1))
	mBase = m.M
	v1146 = m.ExcPending
	if v1146 != 0 {
		goto L1
	} else {
		goto L297
	}
L297:
	;
	goto L249
L298:
	;
	v1156 = v91 + int32(20)
	v1157 = *(*int32)(unsafe.Add(mBase, uint32(v91)+16))
	if v1157 != 0 {
		goto L301
	} else {
		goto L302
	}
L299:
	;
	if v1252 != 0 {
		v204 = v1252
		goto L52
	} else {
		goto L325
	}
L300:
	;
	v1163 = v1156
	v1164 = v1160
	goto L303
L301:
	;
	v1160 = int32(1)
	goto L300
L302:
	;
	v1160 = int32(0)
	goto L300
L303:
	;
	switch v1164 {
	case 0:
		goto L308
	default:
		goto L307
	}
L305:
	;
	v1164 = int32(0)
	goto L303
L306:
	;
	goto L299
L307:
	;
	v1244 = *(*int32)(unsafe.Add(mBase, uint32(v1163)))
	*(*int32)(unsafe.Add(mBase, uint32(v91)+16)) = v1244
	if v1244 == int32(0) {
		goto L305
	} else {
		goto L324
	}
L308:
	;
	v1168 = *(*int32)(unsafe.Add(mBase, uint32(v91)+4))
	if v1168 != int32(-1) {
		v1207 = v1168
		goto L309
	} else {
		goto L310
	}
L309:
	;
	v1208 = int32(1)
	v1209 = v1207 + v1208
	*(*int32)(unsafe.Add(mBase, uint32(v91)+4)) = v1209
	v1211 = int32(0)
	v1214 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
	v1215 = *(*int32)(unsafe.Add(mBase, uint32(v91)+8))
	v1219 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1214+v1215+int32(26)))))
	if v1219 == int32(255) {
		goto L318
	} else {
		goto L319
	}
L310:
	;
	v1172 = *(*int32)(unsafe.Add(mBase, uint32(v91)+8))
	if v1172 != 0 {
		v1207 = int32(-1)
		goto L309
	} else {
		goto L311
	}
L311:
	;
	v1173 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
	v1174 = *(*int32)(unsafe.Add(mBase, uint32(v91)+12))
	if v1174 == int32(0) {
		goto L313
	} else {
		goto L314
	}
L312:
	;
	v1201 = *(*int32)(unsafe.Add(mBase, uint32(v1200)+20))
	if v1201 != int32(-1) {
		goto L315
	} else {
		goto L316
	}
L313:
	;
	v1181 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v1173)+16)))
	v1182 = int64(*(*int8)(unsafe.Add(mBase, uint32(v1173)+27)))
	v1183 = int64(*(*int32)(unsafe.Add(mBase, uint32(v1173)+8)))
	v1184 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v1173)+12)))
	v1185 = int64(*(*int8)(unsafe.Add(mBase, uint32(v1173)+26)))
	v1186 = int64(*(*int32)(unsafe.Add(mBase, uint32(v1173)+4)))
	v1187 = F_wangHash64(m, v1186)
	mBase = m.M
	v1189 = F_wangHash64(m, v1185+v1187)
	mBase = m.M
	v1191 = F_wangHash64(m, v1184+v1189)
	mBase = m.M
	v1193 = F_wangHash64(m, v1183+v1191)
	mBase = m.M
	v1195 = F_wangHash64(m, v1182+v1193)
	mBase = m.M
	v1197 = F_wangHash64(m, v1181+v1195)
	mBase = m.M
	*(*int64)(unsafe.Add(mBase, uint32(v91)+24)) = v1197
	v1199 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
	v1200 = v1199
	goto L312
L314:
	;
	v1177 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1173)+24)))
	v1179 = v1177 + int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v1173)+24)) = uint16(v1179)
	v1200 = v1173
	goto L312
L315:
	;
	v1207 = v1201 + int32(-1)
	goto L309
L316:
	;
	v1204 = *(*int32)(unsafe.Add(mBase, uint32(v91)+4))
	v1207 = v1204
	goto L309
L317:
	;
	v1234 = int32(2)
	v1239 = *(*int32)(unsafe.Add(mBase, uint32(v1214+v1232<<(uint(v1234)%32)+int32(4))))
	v1163 = v1239 + v1233<<(uint(v1234)%32)
	v1164 = int32(1)
	goto L303
L318:
	;
	v1223 = v1211
	goto L320
L319:
	;
	v1223 = v1208 << (uint(v1219) % 32)
	goto L320
L320:
	;
	if v1209 < v1223 {
		v1232 = v1215
		v1233 = v1209
		goto L317
	} else {
		goto L321
	}
L321:
	;
	if v1215 != 0 {
		v1252 = v1211
		goto L306
	} else {
		goto L322
	}
L322:
	;
	v1225 = *(*int32)(unsafe.Add(mBase, uint32(v1214)+20))
	if v1225 == int32(-1) {
		v1252 = v1211
		goto L306
	} else {
		goto L323
	}
L323:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v91)+4)) = int64(4294967296)
	v1232 = int32(1)
	v1233 = int32(0)
	goto L317
L324:
	;
	v1248 = *(*int32)(unsafe.Add(mBase, uint32(v1244)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v1156))) = v1248
	v1252 = v1244
	goto L306
L325:
	;
	goto L53
L326:
	;
	v1270 = *(*int64)(unsafe.Add(mBase, _consts[748]))
	*(*int64)(unsafe.Add(mBase, uint32(v14)+16)) = v1270
	v1276 = F_sdscatprintf(m, v1267, int32(_a1401), v14+int32(16))
	mBase = m.M
	v1277 = m.ExcPending
	if v1277 != 0 {
		goto L1
	} else {
		goto L327
	}
L327:
	;
	v1279 = F_rewriteConfigRewriteLine(m, l0, int32(_a1402), v1276, int32(1))
	mBase = m.M
	v1280 = m.ExcPending
	if v1280 != 0 {
		goto L1
	} else {
		goto L328
	}
L328:
	;
	v1281 = int32(0)
	v1282 = *(*int32)(unsafe.Add(mBase, _consts[749]))
	if v1282 == v1281 {
		goto L330
	} else {
		goto L331
	}
L329:
	;
	v1324 = int32(0)
	v1325 = *(*int32)(unsafe.Add(mBase, _consts[750]))
	if v1325 == v1324 {
		goto L343
	} else {
		goto L344
	}
L330:
	;
	F_rewriteConfigMarkAsProcessed(m, l0, int32(_a1403))
	mBase = m.M
	v1319 = m.ExcPending
	if v1319 != 0 {
		goto L1
	} else {
		goto L341
	}
L331:
	;
	v1287 = F_sdsnew(m, int32(_a1404))
	mBase = m.M
	v1288 = m.ExcPending
	if v1288 != 0 {
		goto L1
	} else {
		goto L332
	}
L332:
	;
	v1290 = *(*int32)(unsafe.Add(mBase, _consts[749]))
	v1293 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1290+int32(-1)))))
	switch v1293 & int32(7) {
	case 0:
		goto L338
	case 1:
		goto L337
	case 2:
		goto L336
	case 3:
		goto L335
	case 4:
		goto L334
	default:
		v1310 = int32(0)
		goto L333
	}
L333:
	;
	v1312 = F_sdscatrepr(m, v1287, v1290, v1310)
	mBase = m.M
	v1313 = m.ExcPending
	if v1313 != 0 {
		goto L1
	} else {
		goto L339
	}
L334:
	;
	v1309 = *(*int32)(unsafe.Add(mBase, uint32(v1290+int32(-17))))
	v1310 = v1309
	goto L333
L335:
	;
	v1306 = *(*int32)(unsafe.Add(mBase, uint32(v1290+int32(-9))))
	v1310 = v1306
	goto L333
L336:
	;
	v1303 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1290+int32(-5)))))
	v1310 = v1303
	goto L333
L337:
	;
	v1300 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1290+int32(-3)))))
	v1310 = v1300
	goto L333
L338:
	;
	v1310 = int32(base.Ui32(v1293) >> (uint(int32(3)) % 32))
	goto L333
L339:
	;
	v1315 = F_rewriteConfigRewriteLine(m, l0, int32(_a1403), v1312, int32(1))
	mBase = m.M
	v1316 = m.ExcPending
	if v1316 != 0 {
		goto L1
	} else {
		goto L340
	}
L340:
	;
	goto L329
L341:
	;
	goto L329
L342:
	;
	v1344 = int32(0)
	v1346 = *(*int32)(unsafe.Add(mBase, _consts[751]))
	if v1346 == v1344 {
		goto L350
	} else {
		goto L351
	}
L343:
	;
	F_rewriteConfigMarkAsProcessed(m, l0, int32(_a1405))
	mBase = m.M
	v1342 = m.ExcPending
	if v1342 != 0 {
		goto L1
	} else {
		goto L348
	}
L344:
	;
	v1328 = F_sdsempty(m)
	mBase = m.M
	v1329 = m.ExcPending
	if v1329 != 0 {
		goto L1
	} else {
		goto L345
	}
L345:
	;
	v1331 = *(*int32)(unsafe.Add(mBase, _consts[750]))
	*(*int32)(unsafe.Add(mBase, uint32(v14))) = v1331
	v1335 = F_sdscatprintf(m, v1328, int32(_a1406), v14)
	mBase = m.M
	v1336 = m.ExcPending
	if v1336 != 0 {
		goto L1
	} else {
		goto L346
	}
L346:
	;
	v1338 = F_rewriteConfigRewriteLine(m, l0, int32(_a1405), v1335, int32(1))
	mBase = m.M
	v1339 = m.ExcPending
	if v1339 != 0 {
		goto L1
	} else {
		goto L347
	}
L347:
	;
	goto L342
L348:
	;
	goto L342
L349:
	;
	v1387 = int32(0)
	v1389 = *(*int32)(unsafe.Add(mBase, _consts[752]))
	if v1389 == v1387 {
		goto L363
	} else {
		goto L364
	}
L350:
	;
	F_rewriteConfigMarkAsProcessed(m, l0, int32(_a1407))
	mBase = m.M
	v1382 = m.ExcPending
	if v1382 != 0 {
		goto L1
	} else {
		goto L361
	}
L351:
	;
	v1350 = F_sdsnew(m, int32(_a1408))
	mBase = m.M
	v1351 = m.ExcPending
	if v1351 != 0 {
		goto L1
	} else {
		goto L352
	}
L352:
	;
	v1353 = *(*int32)(unsafe.Add(mBase, _consts[751]))
	v1356 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1353+int32(-1)))))
	switch v1356 & int32(7) {
	case 0:
		goto L358
	case 1:
		goto L357
	case 2:
		goto L356
	case 3:
		goto L355
	case 4:
		goto L354
	default:
		v1373 = v1344
		goto L353
	}
L353:
	;
	v1375 = F_sdscatrepr(m, v1350, v1353, v1373)
	mBase = m.M
	v1376 = m.ExcPending
	if v1376 != 0 {
		goto L1
	} else {
		goto L359
	}
L354:
	;
	v1372 = *(*int32)(unsafe.Add(mBase, uint32(v1353+int32(-17))))
	v1373 = v1372
	goto L353
L355:
	;
	v1369 = *(*int32)(unsafe.Add(mBase, uint32(v1353+int32(-9))))
	v1373 = v1369
	goto L353
L356:
	;
	v1366 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1353+int32(-5)))))
	v1373 = v1366
	goto L353
L357:
	;
	v1363 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1353+int32(-3)))))
	v1373 = v1363
	goto L353
L358:
	;
	v1373 = int32(base.Ui32(v1356) >> (uint(int32(3)) % 32))
	goto L353
L359:
	;
	v1378 = F_rewriteConfigRewriteLine(m, l0, int32(_a1407), v1375, int32(1))
	mBase = m.M
	v1379 = m.ExcPending
	if v1379 != 0 {
		goto L1
	} else {
		goto L360
	}
L360:
	;
	goto L349
L361:
	;
	goto L349
L362:
	;
	F_dictReleaseIterator(m, v91)
	mBase = m.M
	v1431 = m.ExcPending
	if v1431 != 0 {
		goto L1
	} else {
		goto L375
	}
L363:
	;
	F_rewriteConfigMarkAsProcessed(m, l0, int32(_a1409))
	mBase = m.M
	v1425 = m.ExcPending
	if v1425 != 0 {
		goto L1
	} else {
		goto L374
	}
L364:
	;
	v1393 = F_sdsnew(m, int32(_a1410))
	mBase = m.M
	v1394 = m.ExcPending
	if v1394 != 0 {
		goto L1
	} else {
		goto L365
	}
L365:
	;
	v1396 = *(*int32)(unsafe.Add(mBase, _consts[752]))
	v1399 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1396+int32(-1)))))
	switch v1399 & int32(7) {
	case 0:
		goto L371
	case 1:
		goto L370
	case 2:
		goto L369
	case 3:
		goto L368
	case 4:
		goto L367
	default:
		v1416 = v1387
		goto L366
	}
L366:
	;
	v1418 = F_sdscatrepr(m, v1393, v1396, v1416)
	mBase = m.M
	v1419 = m.ExcPending
	if v1419 != 0 {
		goto L1
	} else {
		goto L372
	}
L367:
	;
	v1415 = *(*int32)(unsafe.Add(mBase, uint32(v1396+int32(-17))))
	v1416 = v1415
	goto L366
L368:
	;
	v1412 = *(*int32)(unsafe.Add(mBase, uint32(v1396+int32(-9))))
	v1416 = v1412
	goto L366
L369:
	;
	v1409 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1396+int32(-5)))))
	v1416 = v1409
	goto L366
L370:
	;
	v1406 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1396+int32(-3)))))
	v1416 = v1406
	goto L366
L371:
	;
	v1416 = int32(base.Ui32(v1399) >> (uint(int32(3)) % 32))
	goto L366
L372:
	;
	v1421 = F_rewriteConfigRewriteLine(m, l0, int32(_a1409), v1418, int32(1))
	mBase = m.M
	v1422 = m.ExcPending
	if v1422 != 0 {
		goto L1
	} else {
		goto L373
	}
L373:
	;
	goto L362
L374:
	;
	goto L362
L375:
	;
	F_rewriteConfigMarkAsProcessed(m, l0, int32(_a1373))
	mBase = m.M
	v1434 = m.ExcPending
	if v1434 != 0 {
		goto L1
	} else {
		goto L376
	}
L376:
	;
	F_rewriteConfigMarkAsProcessed(m, l0, int32(_a1375))
	mBase = m.M
	v1437 = m.ExcPending
	if v1437 != 0 {
		goto L1
	} else {
		goto L377
	}
L377:
	;
	F_rewriteConfigMarkAsProcessed(m, l0, int32(_a1377))
	mBase = m.M
	v1440 = m.ExcPending
	if v1440 != 0 {
		goto L1
	} else {
		goto L378
	}
L378:
	;
	F_rewriteConfigMarkAsProcessed(m, l0, int32(_a1379))
	mBase = m.M
	v1443 = m.ExcPending
	if v1443 != 0 {
		goto L1
	} else {
		goto L379
	}
L379:
	;
	F_rewriteConfigMarkAsProcessed(m, l0, int32(_a1381))
	mBase = m.M
	v1446 = m.ExcPending
	if v1446 != 0 {
		goto L1
	} else {
		goto L380
	}
L380:
	;
	F_rewriteConfigMarkAsProcessed(m, l0, int32(_a1383))
	mBase = m.M
	v1449 = m.ExcPending
	if v1449 != 0 {
		goto L1
	} else {
		goto L381
	}
L381:
	;
	F_rewriteConfigMarkAsProcessed(m, l0, int32(_a1385))
	mBase = m.M
	v1452 = m.ExcPending
	if v1452 != 0 {
		goto L1
	} else {
		goto L382
	}
L382:
	;
	F_rewriteConfigMarkAsProcessed(m, l0, int32(_a1387))
	mBase = m.M
	v1455 = m.ExcPending
	if v1455 != 0 {
		goto L1
	} else {
		goto L383
	}
L383:
	;
	F_rewriteConfigMarkAsProcessed(m, l0, int32(_a1391))
	mBase = m.M
	v1458 = m.ExcPending
	if v1458 != 0 {
		goto L1
	} else {
		goto L384
	}
L384:
	;
	F_rewriteConfigMarkAsProcessed(m, l0, int32(_a1393))
	mBase = m.M
	v1461 = m.ExcPending
	if v1461 != 0 {
		goto L1
	} else {
		goto L385
	}
L385:
	;
	F_rewriteConfigMarkAsProcessed(m, l0, int32(_a1396))
	mBase = m.M
	v1464 = m.ExcPending
	if v1464 != 0 {
		goto L1
	} else {
		goto L386
	}
L386:
	;
	F_rewriteConfigMarkAsProcessed(m, l0, int32(_a1398))
	mBase = m.M
	v1467 = m.ExcPending
	if v1467 != 0 {
		goto L1
	} else {
		goto L387
	}
L387:
	;
	F_rewriteConfigMarkAsProcessed(m, l0, int32(_a1400))
	mBase = m.M
	v1470 = m.ExcPending
	if v1470 != 0 {
		goto L1
	} else {
		goto L388
	}
L388:
	;
	F_rewriteConfigMarkAsProcessed(m, l0, int32(_a1389))
	mBase = m.M
	v1473 = m.ExcPending
	if v1473 != 0 {
		goto L1
	} else {
		goto L389
	}
L389:
	;
	m.G0 = v14 + int32(320)
	return
}
func F_setConfigDirOption(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v23 int32
	_ = v23
	if l2 == int32(1) {
		v8 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
		v9 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8))))
		if v9 != 0 {
			v12 = F_chdir(m, v8)
			mBase = m.M
			if v12 != int32(-1) {
				v23 = int32(1)
			} else {
				v16 = *(*int32)(unsafe.Add(mBase, _consts[5]))
				v17 = F___strerror_l(m, v16, v16)
				mBase = m.M
				v19 = v17
				*(*int32)(unsafe.Add(mBase, uint32(l3))) = v19
				v23 = int32(0)
			}
		} else {
			v19 = int32(_a519)
			*(*int32)(unsafe.Add(mBase, uint32(l3))) = v19
			v23 = int32(0)
		}
	} else {
		v19 = int32(_a520)
		*(*int32)(unsafe.Add(mBase, uint32(l3))) = v19
		v23 = int32(0)
	}
	return v23
}
func F_setConfigNotifyKeyspaceEventsOption(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v34 int32
	_ = v34
	var v39 int32
	_ = v39
	if l2 == int32(1) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v15 = v11
	v16 = int32(0)
	goto L7
L2:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = int32(_a520)
	return int32(0)
L3:
	;
	*(*int32)(unsafe.Add(mBase, _consts[289])) = v39
	return int32(1)
L4:
	;
	if v39 != int32(-1) {
		goto L3
	} else {
		goto L24
	}
L5:
	;
	goto L4
L6:
	;
	v39 = int32(-1)
	goto L5
L7:
	;
	v19 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15))))
	switch v19 {
	case 0:
		v39 = v16
		goto L5
	default:
		goto L6
	case 36:
		goto L22
	case 65:
		v34 = int32(10236)
		goto L9
	case 69:
		goto L14
	case 75:
		goto L15
	case 100:
		goto L11
	case 101:
		goto L16
	case 103:
		goto L23
	case 104:
		goto L19
	case 108:
		goto L21
	case 109:
		goto L12
	case 110:
		goto L10
	case 115:
		goto L20
	case 116:
		goto L13
	case 120:
		goto L17
	case 122:
		goto L18
	}
L9:
	;
	v15 = v15 + int32(1)
	v16 = v16 | v34
	goto L7
L10:
	;
	v34 = int32(16384)
	goto L9
L11:
	;
	v34 = int32(8192)
	goto L9
L12:
	;
	v34 = int32(2048)
	goto L9
L13:
	;
	v34 = int32(1024)
	goto L9
L14:
	;
	v34 = int32(2)
	goto L9
L15:
	;
	v34 = int32(1)
	goto L9
L16:
	;
	v34 = int32(512)
	goto L9
L17:
	;
	v34 = int32(256)
	goto L9
L18:
	;
	v34 = int32(128)
	goto L9
L19:
	;
	v34 = int32(64)
	goto L9
L20:
	;
	v34 = int32(32)
	goto L9
L21:
	;
	v34 = int32(16)
	goto L9
L22:
	;
	v34 = int32(8)
	goto L9
L23:
	;
	v34 = int32(4)
	goto L9
L24:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = int32(_a521)
	return int32(0)
}
func F_setConfigRdmaBindOption(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	v11 = F_setConfigBindOption(m, l1, l2, l3, int32(_a512), int32(_a522))
	v14 = m.ExcPending
	if v14 != 0 {
		return int32(0)
	} else {
		return v11
	}
}
func F_setConfigReplicaOfOption(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v16 int32
	_ = v16
	var v20 int32
	_ = v20
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
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
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v110 int64
	_ = v110
	var v111 int32
	_ = v111
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v132 int32
	_ = v132
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	if l2 == int32(2) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v8 + int32(16)
	return v132
L2:
	;
	v16 = *(*int32)(unsafe.Add(mBase, _consts[130]))
	F_sdsfree(m, v16)
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		goto L4
	} else {
		goto L5
	}
L3:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = int32(_a520)
	v132 = int32(0)
	goto L1
L4:
	;
	return int32(0)
L5:
	;
	*(*int32)(unsafe.Add(mBase, _consts[130])) = int32(0)
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v25 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v26 = int32(_a484)
	v29 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v25))))
	if v29 != 0 {
		goto L9
	} else {
		goto L10
	}
L6:
	;
	v110 = F_strtox_2(m, v24, v8+int32(12), int32(10), int64(2147483648))
	mBase = m.M
	v111 = base.I32_wrap_i64(v110)
	goto L33
L7:
	;
	if v61-v63 != 0 {
		goto L6
	} else {
		goto L19
	}
L8:
	;
	v61 = F_tolower(m, v57)
	mBase = m.M
	v62 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v58))))
	v63 = F_tolower(m, v62)
	mBase = m.M
	goto L7
L9:
	;
	v31 = v25
	v32 = v26
	v33 = v29
	goto L12
L10:
	;
	v57 = int32(0)
	v58 = v26
	goto L8
L11:
	;
	v57 = v54 & int32(255)
	v58 = v53
	goto L8
L12:
	;
	v35 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v32))))
	if v35 == int32(0) {
		v53 = v32
		v54 = v33
		goto L11
	} else {
		goto L14
	}
L13:
	;
	v53 = v47
	v54 = int32(0)
	goto L11
L14:
	;
	v39 = v33 & int32(255)
	if v39 == v35 {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	v46 = int32(1)
	v47 = v32 + v46
	v48 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v31)+1)))
	if v48 != 0 {
		v31 = v31 + v46
		v32 = v47
		v33 = v48
		goto L12
	} else {
		goto L18
	}
L16:
	;
	v41 = F_tolower(m, v39)
	mBase = m.M
	v42 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v32))))
	v43 = F_tolower(m, v42)
	mBase = m.M
	if v41 == v43 {
		goto L15
	} else {
		goto L17
	}
L17:
	;
	v45 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v31))))
	v53 = v32
	v54 = v45
	goto L11
L18:
	;
	goto L13
L19:
	;
	v65 = int32(_a523)
	v68 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24))))
	if v68 != 0 {
		goto L22
	} else {
		goto L23
	}
L20:
	;
	if v100-v102 != 0 {
		goto L6
	} else {
		goto L32
	}
L21:
	;
	v100 = F_tolower(m, v96)
	mBase = m.M
	v101 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v97))))
	v102 = F_tolower(m, v101)
	mBase = m.M
	goto L20
L22:
	;
	v70 = v24
	v71 = v65
	v72 = v68
	goto L25
L23:
	;
	v96 = int32(0)
	v97 = v65
	goto L21
L24:
	;
	v96 = v93 & int32(255)
	v97 = v92
	goto L21
L25:
	;
	v74 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v71))))
	if v74 == int32(0) {
		v92 = v71
		v93 = v72
		goto L24
	} else {
		goto L27
	}
L26:
	;
	v92 = v86
	v93 = int32(0)
	goto L24
L27:
	;
	v78 = v72 & int32(255)
	if v78 == v74 {
		goto L28
	} else {
		goto L29
	}
L28:
	;
	v85 = int32(1)
	v86 = v71 + v85
	v87 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v70)+1)))
	if v87 != 0 {
		v70 = v70 + v85
		v71 = v86
		v72 = v87
		goto L25
	} else {
		goto L31
	}
L29:
	;
	v80 = F_tolower(m, v78)
	mBase = m.M
	v81 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v71))))
	v82 = F_tolower(m, v81)
	mBase = m.M
	if v80 == v82 {
		goto L28
	} else {
		goto L30
	}
L30:
	;
	v84 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v70))))
	v92 = v71
	v93 = v84
	goto L24
L31:
	;
	goto L26
L32:
	;
	v132 = int32(1)
	goto L1
L33:
	;
	*(*int32)(unsafe.Add(mBase, _consts[290])) = v111
	if base.Ui32(int32(65535)) < base.Ui32(v111) {
		goto L35
	} else {
		goto L36
	}
L34:
	;
	v123 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v124 = F_sdsnew(m, v123)
	mBase = m.M
	v125 = m.ExcPending
	if v125 != 0 {
		goto L4
	} else {
		goto L38
	}
L35:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = int32(_a524)
	v132 = int32(0)
	goto L1
L36:
	;
	v115 = *(*int32)(unsafe.Add(mBase, uint32(v8)+12))
	v116 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v115))))
	if v116 == int32(0) {
		goto L34
	} else {
		goto L37
	}
L37:
	;
	goto L35
L38:
	;
	v126 = int32(_a44)
	*(*int32)(unsafe.Add(mBase, _consts[193])) = int32(1)
	*(*int32)(unsafe.Add(mBase, _consts[130])) = v124
	v132 = int32(1)
	goto L1
}
