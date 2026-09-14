package p2

import (
	base "github.com/shibukawa/vkmem/internal/aot/vkaot/base"
	"unsafe"
)

func F___quicklistCompressNode(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v21 int32
	_ = v21
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
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
	var v67 int32
	_ = v67
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v103 int32
	_ = v103
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v122 int32
	_ = v122
	var v126 int32
	_ = v126
	var v130 int32
	_ = v130
	var v133 int32
	_ = v133
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v223 int32
	_ = v223
	var v225 int32
	_ = v225
	var v227 int32
	_ = v227
	var v239 int32
	_ = v239
	var v247 int32
	_ = v247
	var v250 int32
	_ = v250
	var v252 int32
	_ = v252
	var v254 int32
	_ = v254
	var v255 int32
	_ = v255
	var v257 int32
	_ = v257
	var v269 int32
	_ = v269
	var v270 int32
	_ = v270
	var v280 int32
	_ = v280
	var v287 int32
	_ = v287
	var v289 int32
	_ = v289
	var v290 int32
	_ = v290
	var v292 int32
	_ = v292
	var v312 int32
	_ = v312
	var v320 int32
	_ = v320
	var v325 int32
	_ = v325
	var v332 int32
	_ = v332
	var v338 int32
	_ = v338
	var v347 int32
	_ = v347
	var v348 int32
	_ = v348
	var v352 int32
	_ = v352
	var v353 int32
	_ = v353
	var v354 int32
	_ = v354
	var v357 int32
	_ = v357
	var v358 int32
	_ = v358
	var v361 int32
	_ = v361
	var v362 int32
	_ = v362
	var v363 int32
	_ = v363
	var v364 int32
	_ = v364
	var v369 int32
	_ = v369
	var v371 int32
	_ = v371
	var v377 int32
	_ = v377
	var v378 int32
	_ = v378
	var v390 int32
	_ = v390
	var v392 int32
	_ = v392
	var v398 int32
	_ = v398
	var v399 int32
	_ = v399
	var v401 int32
	_ = v401
	var v407 int32
	_ = v407
	var v409 int32
	_ = v409
	var v415 int32
	_ = v415
	var v418 int32
	_ = v418
	var v426 int32
	_ = v426
	var v428 int32
	_ = v428
	var v431 int32
	_ = v431
	var v433 int32
	_ = v433
	var v434 int32
	_ = v434
	var v436 int32
	_ = v436
	var v446 int32
	_ = v446
	var v448 int32
	_ = v448
	var v449 int32
	_ = v449
	var v450 int32
	_ = v450
	var v451 int32
	_ = v451
	var v457 int32
	_ = v457
	var v464 int32
	_ = v464
	var v466 int32
	_ = v466
	var v470 int32
	_ = v470
	var v473 int32
	_ = v473
	var v484 int32
	_ = v484
	var v485 int32
	_ = v485
	var v486 int32
	_ = v486
	var v489 int32
	_ = v489
	var v500 int32
	_ = v500
	var v501 int32
	_ = v501
	var v503 int32
	_ = v503
	var v508 int32
	_ = v508
	var v517 int32
	_ = v517
	var v521 int32
	_ = v521
	var v525 int32
	_ = v525
	var v536 int32
	_ = v536
	var v555 int32
	_ = v555
	var v558 int32
	_ = v558
	var v563 int32
	_ = v563
	var v564 int32
	_ = v564
	var v565 int32
	_ = v565
	var v567 int32
	_ = v567
	var v569 int32
	_ = v569
	var v577 int32
	_ = v577
	var v583 int32
	_ = v583
	v2 = int32(0)
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v7 = v5 | int32(2097152)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v7
	if v5&int32(4194304) != 0 {
		v577 = v2
		goto L2
	} else {
		goto L3
	}
L1:
	;
	F__serverAssert(m, int32(_a1111), int32(_a1112), int32(218))
	mBase = m.M
	v583 = m.ExcPending
	if v583 != 0 {
		goto L7
	} else {
		goto L110
	}
L2:
	;
	return v577
L3:
	;
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v12 == int32(0) {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v15 == int32(0) {
		goto L1
	} else {
		goto L5
	}
L5:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v7 & int32(-5242881)
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if base.Ui32(v21) < base.Ui32(int32(48)) {
		v577 = v2
		goto L2
	} else {
		goto L6
	}
L6:
	;
	v26 = F_valkey_malloc(m, v21+int32(4))
	mBase = m.M
	v29 = m.ExcPending
	if v29 != 0 {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	return int32(0)
L8:
	;
	v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v33 = v26 + int32(4)
	v34 = int32(0)
	v46 = m.G0
	v48 = v46 - int32(262144)
	m.G0 = v48
	if v31 == v34 {
		v536 = v34
		goto L10
	} else {
		goto L11
	}
L9:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26))) = v536
	if v536 == int32(0) {
		goto L104
	} else {
		goto L105
	}
L10:
	;
	m.G0 = v48 + int32(262144)
	goto L9
L11:
	;
	if v31 == int32(0) {
		v536 = v34
		goto L10
	} else {
		goto L12
	}
L12:
	;
	v55 = v33 + v31
	v57 = v26 + int32(5)
	v59 = v30 + v31
	v61 = v59 + int32(-2)
	if base.Ui32(v30) < base.Ui32(v61) {
		goto L14
	} else {
		goto L15
	}
L13:
	;
	if base.Ui32(v55) < base.Ui32(v409+int32(3)) {
		v536 = int32(0)
		goto L10
	} else {
		goto L84
	}
L14:
	;
	v64 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v30))))
	v67 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v30)+1)))
	v71 = v30
	v73 = v57
	v79 = int32(0)
	v80 = v64<<(uint(int32(8))%32) | v67
	goto L16
L15:
	;
	v407 = v30
	v409 = v57
	v415 = int32(0)
	v418 = v34
	goto L13
L16:
	;
	v88 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v71)+2)))
	v89 = v80<<(uint(int32(8))%32) | v88
	v97 = v48 + (v89*int32(65531)+v80)&int32(65535)<<(uint(int32(2))%32)
	v98 = *(*int32)(unsafe.Add(mBase, uint32(v97)))
	*(*int32)(unsafe.Add(mBase, uint32(v97))) = v71
	if base.Ui32(v98) <= base.Ui32(v30) {
		goto L21
	} else {
		goto L22
	}
L17:
	;
	v407 = v390
	v409 = v392
	v415 = v398
	v418 = v401
	goto L13
L18:
	;
	if base.Ui32(v390) < base.Ui32(v61) {
		v71 = v390
		v73 = v392
		v79 = v398
		v80 = v399
		goto L16
	} else {
		goto L83
	}
L19:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v332))) = uint8(v103)
	v347 = v332 + int32(2)
	v348 = v71 + v338
	if base.Ui32(v348) < base.Ui32(v61) {
		goto L81
	} else {
		goto L82
	}
L20:
	;
	v320 = v312 + int32(-9)
	*(*uint8)(unsafe.Add(mBase, uint32(v130)+1)) = uint8(v320)
	v325 = int32(base.Ui32(v103)>>(uint(int32(8))%32)) | int32(224)
	*(*uint8)(unsafe.Add(mBase, uint32(v130))) = uint8(v325)
	v332 = v130 + int32(2)
	v338 = v312
	goto L19
L21:
	;
	if base.Ui32(v73) < base.Ui32(v55) {
		goto L77
	} else {
		goto L78
	}
L22:
	;
	v103 = v98 ^ int32(-1) + v71
	if base.Ui32(int32(8191)) < base.Ui32(v103) {
		goto L21
	} else {
		goto L23
	}
L23:
	;
	v106 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v98)+2)))
	v107 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v71)+2)))
	if v106 != v107 {
		goto L21
	} else {
		goto L24
	}
L24:
	;
	v109 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v98))))
	v110 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v71))))
	if v109 != v110 {
		goto L21
	} else {
		goto L25
	}
L25:
	;
	if base.Ui32(v73+int32(4)) < base.Ui32(v55) {
		goto L26
	} else {
		goto L27
	}
L26:
	;
	v122 = int32(-1)
	v126 = v79 + v122
	*(*uint8)(unsafe.Add(mBase, uint32(v73+(v79^v122)))) = uint8(v126)
	v130 = v73 - base.B2i32(v79 == int32(0))
	v133 = v59 - v71 + int32(-2)
	if base.Ui32(int32(17)) <= base.Ui32(v133) {
		goto L31
	} else {
		goto L32
	}
L27:
	;
	if base.Ui32(v73-base.B2i32(v79 == int32(0))+int32(4)) < base.Ui32(v55) {
		goto L26
	} else {
		goto L28
	}
L28:
	;
	v536 = int32(0)
	goto L10
L29:
	;
	v280 = v270<<(uint(int32(5))%32) | int32(base.Ui32(v103)>>(uint(int32(8))%32))
	*(*uint8)(unsafe.Add(mBase, uint32(v130))) = uint8(v280)
	v332 = v130 + int32(1)
	v338 = v269
	goto L19
L30:
	;
	v220 = int32(264)
	if base.Ui32(v133) < base.Ui32(v220) {
		goto L64
	} else {
		goto L65
	}
L31:
	;
	v137 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v98)+3)))
	v138 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v71)+3)))
	if v137 == v138 {
		goto L33
	} else {
		goto L34
	}
L32:
	;
	v219 = int32(2)
	goto L30
L33:
	;
	v144 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v98)+4)))
	v145 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v71)+4)))
	if v144 == v145 {
		goto L35
	} else {
		goto L36
	}
L34:
	;
	v269 = int32(3)
	v270 = int32(1)
	goto L29
L35:
	;
	v151 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v98)+5)))
	v152 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v71)+5)))
	if v151 == v152 {
		goto L37
	} else {
		goto L38
	}
L36:
	;
	v269 = int32(4)
	v270 = int32(2)
	goto L29
L37:
	;
	v158 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v98)+6)))
	v159 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v71)+6)))
	if v158 == v159 {
		goto L39
	} else {
		goto L40
	}
L38:
	;
	v269 = int32(5)
	v270 = int32(3)
	goto L29
L39:
	;
	v165 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v98)+7)))
	v166 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v71)+7)))
	if v165 == v166 {
		goto L41
	} else {
		goto L42
	}
L40:
	;
	v269 = int32(6)
	v270 = int32(4)
	goto L29
L41:
	;
	v172 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v98)+8)))
	v173 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v71)+8)))
	if v172 == v173 {
		goto L43
	} else {
		goto L44
	}
L42:
	;
	v269 = int32(7)
	v270 = int32(5)
	goto L29
L43:
	;
	v179 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v98)+9)))
	v180 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v71)+9)))
	if v179 == v180 {
		goto L45
	} else {
		goto L46
	}
L44:
	;
	v269 = int32(8)
	v270 = int32(6)
	goto L29
L45:
	;
	v183 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v98)+10)))
	v184 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v71)+10)))
	if v183 == v184 {
		goto L47
	} else {
		goto L48
	}
L46:
	;
	v312 = int32(9)
	goto L20
L47:
	;
	v187 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v98)+11)))
	v188 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v71)+11)))
	if v187 == v188 {
		goto L49
	} else {
		goto L50
	}
L48:
	;
	v312 = int32(10)
	goto L20
L49:
	;
	v191 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v98)+12)))
	v192 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v71)+12)))
	if v191 == v192 {
		goto L51
	} else {
		goto L52
	}
L50:
	;
	v312 = int32(11)
	goto L20
L51:
	;
	v195 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v98)+13)))
	v196 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v71)+13)))
	if v195 == v196 {
		goto L53
	} else {
		goto L54
	}
L52:
	;
	v312 = int32(12)
	goto L20
L53:
	;
	v199 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v98)+14)))
	v200 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v71)+14)))
	if v199 == v200 {
		goto L55
	} else {
		goto L56
	}
L54:
	;
	v312 = int32(13)
	goto L20
L55:
	;
	v203 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v98)+15)))
	v204 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v71)+15)))
	if v203 == v204 {
		goto L57
	} else {
		goto L58
	}
L56:
	;
	v312 = int32(14)
	goto L20
L57:
	;
	v207 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v98)+16)))
	v208 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v71)+16)))
	if v207 == v208 {
		goto L59
	} else {
		goto L60
	}
L58:
	;
	v312 = int32(15)
	goto L20
L59:
	;
	v211 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v98)+17)))
	v212 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v71)+17)))
	if v211 == v212 {
		goto L61
	} else {
		goto L62
	}
L60:
	;
	v312 = int32(16)
	goto L20
L61:
	;
	v215 = int32(18)
	v216 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v98)+18)))
	v217 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v71)+18)))
	if v216 != v217 {
		v312 = v215
		goto L20
	} else {
		goto L63
	}
L62:
	;
	v312 = int32(17)
	goto L20
L63:
	;
	v219 = v215
	goto L30
L64:
	;
	v223 = v133
	goto L66
L65:
	;
	v223 = v220
	goto L66
L66:
	;
	v225 = v219 | int32(1)
	if base.Ui32(v225) < base.Ui32(v223) {
		goto L67
	} else {
		goto L68
	}
L67:
	;
	v227 = v223
	goto L69
L68:
	;
	v227 = v225
	goto L69
L69:
	;
	v239 = v219
	goto L71
L70:
	;
	v257 = v255 + int32(-1)
	if base.Ui32(int32(6)) < base.Ui32(v257) {
		v312 = v254
		goto L20
	} else {
		goto L76
	}
L71:
	;
	v247 = v239 + int32(1)
	if base.Ui32(v247) < base.Ui32(v223) {
		goto L73
	} else {
		goto L74
	}
L72:
	;
	v254 = v247
	v255 = v239
	goto L70
L73:
	;
	v250 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v98+v247))))
	v252 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v71+v247))))
	if v250 == v252 {
		v239 = v247
		goto L71
	} else {
		goto L75
	}
L74:
	;
	v254 = v227
	v255 = v227 + int32(-1)
	goto L70
L75:
	;
	goto L72
L76:
	;
	v269 = v254
	v270 = v257
	goto L29
L77:
	;
	v287 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v71))))
	*(*uint8)(unsafe.Add(mBase, uint32(v73))) = uint8(v287)
	v289 = int32(1)
	v290 = v71 + v289
	v292 = v79 + v289
	if v292 == int32(32) {
		goto L79
	} else {
		goto L80
	}
L78:
	;
	v536 = int32(0)
	goto L10
L79:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v73+int32(-32)))) = uint8(v79)
	v390 = v290
	v392 = v73 + int32(2)
	v398 = int32(0)
	v399 = v89
	v401 = v292
	goto L18
L80:
	;
	v390 = v290
	v392 = v73 + int32(1)
	v398 = v292
	v399 = v89
	v401 = v292
	goto L18
L81:
	;
	v352 = v348 + int32(-1)
	v353 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v352))))
	v354 = int32(8)
	v357 = v348 + int32(-2)
	v358 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v357))))
	v361 = v353<<(uint(v354)%32) | v358<<(uint(int32(16))%32)
	v362 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v348))))
	v363 = v361 | v362
	v364 = int32(65531)
	v369 = int32(65535)
	v371 = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v48+(v363*v364+int32(base.Ui32(v361)>>(uint(v354)%32)))&v369<<(uint(v371)%32)))) = v357
	v377 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v348)+1)))
	v378 = v363<<(uint(v354)%32) | v377
	*(*int32)(unsafe.Add(mBase, uint32(v48+(v378*v364+v363)&v369<<(uint(v371)%32)))) = v352
	v390 = v348
	v392 = v347
	v398 = int32(0)
	v399 = v378
	v401 = v352
	goto L18
L82:
	;
	v407 = v348
	v409 = v347
	v415 = int32(0)
	v418 = v103
	goto L13
L83:
	;
	goto L17
L84:
	;
	if base.Ui32(v407) < base.Ui32(v59) {
		goto L86
	} else {
		goto L87
	}
L85:
	;
	v521 = int32(-1)
	v525 = v517 + v521
	*(*uint8)(unsafe.Add(mBase, uint32(v508+(v517^v521)))) = uint8(v525)
	v536 = v508 - base.B2i32(v517 == int32(0)) - v33
	goto L10
L86:
	;
	v426 = int32(1)
	v428 = v59 - v407
	if v428&v426 != 0 {
		goto L89
	} else {
		goto L90
	}
L87:
	;
	v508 = v409
	v517 = v415
	goto L85
L88:
	;
	if v59 == v407+v426 {
		v508 = v448
		v517 = v451
		goto L85
	} else {
		goto L93
	}
L89:
	;
	v431 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v407))))
	*(*uint8)(unsafe.Add(mBase, uint32(v409))) = uint8(v431)
	v433 = int32(1)
	v434 = v407 + v433
	v436 = v415 + v433
	if v436 == int32(32) {
		goto L91
	} else {
		goto L92
	}
L90:
	;
	v448 = v409
	v449 = v415
	v450 = v407
	v451 = v418
	goto L88
L91:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v409+int32(-32)))) = uint8(v415)
	v446 = int32(0)
	v448 = v409 + int32(2)
	v449 = v446
	v450 = v434
	v451 = v446
	goto L88
L92:
	;
	v448 = v409 + int32(1)
	v449 = v436
	v450 = v434
	v451 = v436
	goto L88
L93:
	;
	v457 = v448
	v464 = v450
	v466 = v449
	goto L94
L94:
	;
	v470 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v464))))
	*(*uint8)(unsafe.Add(mBase, uint32(v457))) = uint8(v470)
	v473 = v466 + int32(1)
	if v473 == int32(32) {
		goto L97
	} else {
		goto L98
	}
L95:
	;
	v508 = v500
	v517 = v501
	goto L85
L96:
	;
	v486 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v464)+1)))
	*(*uint8)(unsafe.Add(mBase, uint32(v484))) = uint8(v486)
	v489 = v485 + int32(1)
	if v489 == int32(32) {
		goto L100
	} else {
		goto L101
	}
L97:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v457+int32(-32)))) = uint8(v466)
	v484 = v457 + int32(2)
	v485 = int32(0)
	goto L96
L98:
	;
	v484 = v457 + int32(1)
	v485 = v473
	goto L96
L99:
	;
	v503 = v464 + int32(2)
	if v503 != v407+v428 {
		v457 = v500
		v464 = v503
		v466 = v501
		goto L94
	} else {
		goto L102
	}
L100:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v484+int32(-32)))) = uint8(v485)
	v500 = v484 + int32(2)
	v501 = int32(0)
	goto L99
L101:
	;
	v500 = v484 + int32(1)
	v501 = v489
	goto L99
L102:
	;
	goto L95
L103:
	;
	v563 = F_valkey_realloc(m, v26, v536+int32(4))
	mBase = m.M
	v564 = m.ExcPending
	if v564 != 0 {
		goto L7
	} else {
		goto L108
	}
L104:
	;
	F_valkey_free(m, v26)
	mBase = m.M
	v558 = m.ExcPending
	if v558 != 0 {
		goto L7
	} else {
		goto L107
	}
L105:
	;
	v555 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if base.Ui32(v536+int32(8)) < base.Ui32(v555) {
		goto L103
	} else {
		goto L106
	}
L106:
	;
	goto L104
L107:
	;
	return int32(0)
L108:
	;
	v565 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	F_valkey_free(m, v565)
	mBase = m.M
	v567 = m.ExcPending
	if v567 != 0 {
		goto L7
	} else {
		goto L109
	}
L109:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v563
	v569 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v569&int32(-196609) | int32(131072)
	v577 = int32(1)
	goto L2
L110:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
}
func F___quicklistCreateNode(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
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
	var v28 int32
	_ = v28
	v8 = F_valkey_malloc(m, int32(20))
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return int32(0)
	} else {
		*(*int64)(unsafe.Add(mBase, uint32(v8))) = int64(0)
		v14 = *(*int32)(unsafe.Add(mBase, uint32(v8)+16))
		if l0 != int32(1) {
			v24 = F_lpNew(m, int32(0))
			mBase = m.M
			v25 = m.ExcPending
			if v25 != 0 {
				return int32(0)
			} else {
				v26 = F_lpPrepend(m, v24, l1, l2)
				mBase = m.M
				v27 = m.ExcPending
				if v27 != 0 {
					return int32(0)
				} else {
					v28 = v26
					*(*int32)(unsafe.Add(mBase, uint32(v8)+12)) = l2
					*(*int32)(unsafe.Add(mBase, uint32(v8)+8)) = v28
					*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = v14&int32(-6291456) | l0<<(uint(int32(18))%32)&int32(786432) | int32(65537)
					return v8
				}
			}
		} else {
			v17 = F_valkey_malloc(m, l2)
			mBase = m.M
			v18 = m.ExcPending
			if v18 != 0 {
				return int32(0)
			} else {
				if l2 == int32(0) {
				} else {
					v21 = F__emscripten_memcpy_bulkmem(m, v17, l1, l2)
					mBase = m.M
				}
				v28 = v17
				*(*int32)(unsafe.Add(mBase, uint32(v8)+12)) = l2
				*(*int32)(unsafe.Add(mBase, uint32(v8)+8)) = v28
				*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = v14&int32(-6291456) | l0<<(uint(int32(18))%32)&int32(786432) | int32(65537)
				return v8
			}
		}
	}
}
func F__quicklistInsert(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) {
	mBase := m.M
	_ = mBase
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v85 int32
	_ = v85
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	var v95 int32
	_ = v95
	var v100 int32
	_ = v100
	var v103 int32
	_ = v103
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
	var v117 int32
	_ = v117
	var v121 int32
	_ = v121
	var v123 int32
	_ = v123
	var v134 int32
	_ = v134
	var v137 int32
	_ = v137
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v147 int32
	_ = v147
	var v150 int32
	_ = v150
	var v153 int32
	_ = v153
	var v156 int32
	_ = v156
	var v163 int32
	_ = v163
	var v168 int32
	_ = v168
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v179 int32
	_ = v179
	var v186 int32
	_ = v186
	var v189 int32
	_ = v189
	var v196 int32
	_ = v196
	var v198 int32
	_ = v198
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v206 int32
	_ = v206
	var v213 int32
	_ = v213
	var v219 int32
	_ = v219
	var v222 int32
	_ = v222
	var v225 int32
	_ = v225
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v235 int32
	_ = v235
	var v238 int32
	_ = v238
	var v241 int32
	_ = v241
	var v244 int32
	_ = v244
	var v251 int32
	_ = v251
	var v256 int32
	_ = v256
	var v260 int32
	_ = v260
	var v261 int32
	_ = v261
	var v267 int32
	_ = v267
	var v275 int32
	_ = v275
	var v278 int32
	_ = v278
	var v285 int32
	_ = v285
	var v287 int32
	_ = v287
	var v291 int32
	_ = v291
	var v292 int32
	_ = v292
	var v296 int32
	_ = v296
	var v298 int32
	_ = v298
	var v305 int32
	_ = v305
	var v309 int32
	_ = v309
	var v310 int32
	_ = v310
	var v311 int32
	_ = v311
	var v312 int32
	_ = v312
	var v320 int32
	_ = v320
	var v323 int32
	_ = v323
	var v325 int32
	_ = v325
	var v327 int32
	_ = v327
	var v328 int32
	_ = v328
	var v329 int32
	_ = v329
	var v332 int32
	_ = v332
	var v333 int32
	_ = v333
	var v334 int32
	_ = v334
	var v335 int32
	_ = v335
	var v336 int32
	_ = v336
	var v337 int32
	_ = v337
	var v338 int32
	_ = v338
	var v339 int32
	_ = v339
	var v347 int32
	_ = v347
	var v350 int32
	_ = v350
	var v357 int32
	_ = v357
	var v360 int32
	_ = v360
	var v368 int32
	_ = v368
	var v374 int32
	_ = v374
	var v375 int32
	_ = v375
	var v379 int32
	_ = v379
	var v380 int32
	_ = v380
	var v381 int32
	_ = v381
	var v382 int32
	_ = v382
	var v384 int32
	_ = v384
	var v385 int32
	_ = v385
	var v387 int32
	_ = v387
	var v389 int32
	_ = v389
	var v390 int32
	_ = v390
	var v395 int32
	_ = v395
	var v406 int32
	_ = v406
	var v407 int32
	_ = v407
	var v408 int32
	_ = v408
	var v409 int32
	_ = v409
	var v412 int32
	_ = v412
	var v413 int32
	_ = v413
	var v414 int32
	_ = v414
	var v416 int32
	_ = v416
	var v417 int32
	_ = v417
	var v419 int32
	_ = v419
	var v421 int32
	_ = v421
	var v426 int32
	_ = v426
	var v432 int32
	_ = v432
	var v433 int32
	_ = v433
	var v436 int32
	_ = v436
	var v437 int32
	_ = v437
	var v439 int32
	_ = v439
	var v441 int32
	_ = v441
	var v443 int32
	_ = v443
	var v445 int32
	_ = v445
	var v450 int32
	_ = v450
	var v451 int32
	_ = v451
	var v455 int32
	_ = v455
	var v464 int32
	_ = v464
	var v465 int32
	_ = v465
	var v466 int32
	_ = v466
	var v467 int32
	_ = v467
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
	var v477 int32
	_ = v477
	var v479 int32
	_ = v479
	var v484 int32
	_ = v484
	var v490 int32
	_ = v490
	var v491 int32
	_ = v491
	var v492 int32
	_ = v492
	var v494 int32
	_ = v494
	var v495 int32
	_ = v495
	var v497 int32
	_ = v497
	var v499 int32
	_ = v499
	var v501 int32
	_ = v501
	var v503 int32
	_ = v503
	var v508 int32
	_ = v508
	var v509 int32
	_ = v509
	var v516 int32
	_ = v516
	var v519 int32
	_ = v519
	var v527 int32
	_ = v527
	var v528 int32
	_ = v528
	var v529 int32
	_ = v529
	var v530 int32
	_ = v530
	var v533 int32
	_ = v533
	var v534 int32
	_ = v534
	var v535 int32
	_ = v535
	var v537 int32
	_ = v537
	var v538 int32
	_ = v538
	var v540 int32
	_ = v540
	var v542 int32
	_ = v542
	var v547 int32
	_ = v547
	var v553 int32
	_ = v553
	var v554 int32
	_ = v554
	var v555 int32
	_ = v555
	var v557 int32
	_ = v557
	var v559 int32
	_ = v559
	var v561 int32
	_ = v561
	var v565 int32
	_ = v565
	var v570 int32
	_ = v570
	var v571 int32
	_ = v571
	var v572 int32
	_ = v572
	var v577 int32
	_ = v577
	var v578 int32
	_ = v578
	var v583 int32
	_ = v583
	var v586 int32
	_ = v586
	var v592 int32
	_ = v592
	var v593 int32
	_ = v593
	var v597 int32
	_ = v597
	var v598 int32
	_ = v598
	var v599 int32
	_ = v599
	var v601 int32
	_ = v601
	var v603 int32
	_ = v603
	var v605 int32
	_ = v605
	var v609 int32
	_ = v609
	var v614 int32
	_ = v614
	var v615 int32
	_ = v615
	var v616 int32
	_ = v616
	var v621 int32
	_ = v621
	var v622 int32
	_ = v622
	var v631 int32
	_ = v631
	var v632 int32
	_ = v632
	var v634 int32
	_ = v634
	var v635 int32
	_ = v635
	var v636 int32
	_ = v636
	var v637 int32
	_ = v637
	var v639 int32
	_ = v639
	var v641 int32
	_ = v641
	var v643 int32
	_ = v643
	var v646 int32
	_ = v646
	var v652 int32
	_ = v652
	var v653 int32
	_ = v653
	var v657 int32
	_ = v657
	var v658 int32
	_ = v658
	var v659 int32
	_ = v659
	var v660 int32
	_ = v660
	var v661 int32
	_ = v661
	var v664 int32
	_ = v664
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
	var v672 int32
	_ = v672
	var v674 int32
	_ = v674
	var v677 int32
	_ = v677
	var v678 int32
	_ = v678
	var v679 int32
	_ = v679
	var v686 int32
	_ = v686
	var v693 int32
	_ = v693
	var v694 int32
	_ = v694
	var v695 int32
	_ = v695
	var v699 int32
	_ = v699
	var v706 int32
	_ = v706
	var v707 int32
	_ = v707
	var v708 int32
	_ = v708
	var v709 int32
	_ = v709
	var v715 int32
	_ = v715
	var v719 int32
	_ = v719
	var v721 int32
	_ = v721
	var v728 int32
	_ = v728
	var v729 int32
	_ = v729
	var v730 int32
	_ = v730
	var v731 int32
	_ = v731
	var v733 int32
	_ = v733
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v21 = *(*int32)(unsafe.Add(mBase, uint32(v20)+16))
	v22 = int32(18)
	v25 = v21 << (uint(v22) % 32) >> (uint(v22) % 32)
	v26 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v26 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	v693 = F_valkey_malloc(m, int32(20))
	mBase = m.M
	v694 = m.ExcPending
	if v694 != 0 {
		goto L15
	} else {
		goto L212
	}
L2:
	;
	v51 = int32(0)
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v26)+16))
	v54 = v52 & int32(786432)
	if v54 == int32(262144) {
		v121 = v51
		goto L17
	} else {
		goto L18
	}
L3:
	;
	v27 = int32(0)
	v28 = *(*int32)(unsafe.Add(mBase, _consts[381]))
	if v28 == v27 {
		goto L5
	} else {
		goto L6
	}
L4:
	;
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v20)+4))
	F___quicklistInsertPlainNode(m, v20, v48, l2, l3, l4)
	mBase = m.M
	v50 = m.ExcPending
	if v50 != 0 {
		goto L15
	} else {
		goto L16
	}
L5:
	;
	if v25 < int32(0) {
		goto L8
	} else {
		goto L9
	}
L6:
	;
	if base.Ui32(v28) <= base.Ui32(l3) {
		goto L4
	} else {
		goto L7
	}
L7:
	;
	goto L1
L8:
	;
	v36 = int32(-5)
	if base.Ui32(v36) < base.Ui32(v25) {
		goto L11
	} else {
		goto L12
	}
L9:
	;
	if base.Ui32(l3) <= base.Ui32(int32(8192)) {
		goto L1
	} else {
		goto L10
	}
L10:
	;
	goto L4
L11:
	;
	v39 = v25
	goto L13
L12:
	;
	v39 = v36
	goto L13
L13:
	;
	v46 = *(*int32)(unsafe.Add(mBase, uint32((v39^int32(-1))<<(uint(int32(2))%32))+uint32(_consts[628])))
	if base.Ui32(l3) <= base.Ui32(v46) {
		goto L1
	} else {
		goto L14
	}
L14:
	;
	goto L4
L15:
	;
	return
L16:
	;
	return
L17:
	;
	v123 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
	if l4 == int32(0) {
		goto L44
	} else {
		goto L45
	}
L18:
	;
	v57 = int32(0)
	v58 = *(*int32)(unsafe.Add(mBase, _consts[381]))
	if v58 == v57 {
		goto L22
	} else {
		goto L23
	}
L19:
	;
	v121 = v117 ^ int32(1)
	goto L17
L20:
	;
	v100 = int32(-5)
	if base.Ui32(v100) < base.Ui32(v25) {
		goto L34
	} else {
		goto L35
	}
L21:
	;
	v92 = int32(1)
	if base.Ui32(v92) < base.Ui32(v25) {
		goto L31
	} else {
		goto L32
	}
L22:
	;
	v80 = int32(0)
	if v25 < v80 {
		goto L20
	} else {
		goto L29
	}
L23:
	;
	if base.Ui32(v58) <= base.Ui32(l3) {
		v121 = v51
		goto L17
	} else {
		goto L24
	}
L24:
	;
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v26)+12))
	v65 = l3 + v62 + int32(8)
	if int32(-1) < v25 {
		v89 = v65
		goto L21
	} else {
		goto L25
	}
L25:
	;
	v68 = int32(-5)
	if base.Ui32(v68) < base.Ui32(v25) {
		goto L26
	} else {
		goto L27
	}
L26:
	;
	v71 = v25
	goto L28
L27:
	;
	v71 = v68
	goto L28
L28:
	;
	v78 = *(*int32)(unsafe.Add(mBase, uint32((v71^int32(-1))<<(uint(int32(2))%32))+uint32(_consts[628])))
	v117 = base.B2i32(base.Ui32(v78) < base.Ui32(v65))
	goto L19
L29:
	;
	if base.Ui32(int32(8192)) < base.Ui32(l3) {
		v121 = v80
		goto L17
	} else {
		goto L30
	}
L30:
	;
	v85 = *(*int32)(unsafe.Add(mBase, uint32(v26)+12))
	v89 = l3 + v85 + int32(8)
	goto L21
L31:
	;
	v95 = v25
	goto L33
L32:
	;
	v95 = v92
	goto L33
L33:
	;
	v117 = base.B2i32(base.Ui32(v95) <= base.Ui32(v52&int32(65535))) | base.B2i32(base.Ui32(int32(8192)) < base.Ui32(v89))
	goto L19
L34:
	;
	v103 = v25
	goto L36
L35:
	;
	v103 = v100
	goto L36
L36:
	;
	v110 = *(*int32)(unsafe.Add(mBase, uint32((v103^int32(-1))<<(uint(int32(2))%32))+uint32(_consts[628])))
	if base.Ui32(v110) < base.Ui32(l3) {
		v121 = v80
		goto L17
	} else {
		goto L37
	}
L37:
	;
	v112 = *(*int32)(unsafe.Add(mBase, uint32(v26)+12))
	v117 = base.B2i32(base.Ui32(v110) < base.Ui32(l3+v112+int32(8)))
	goto L19
L38:
	;
	v338 = int32(0)
	v339 = *(*int32)(unsafe.Add(mBase, _consts[381]))
	if v339 == v338 {
		goto L103
	} else {
		goto L104
	}
L39:
	;
	v329 = int32(0)
	v332 = v325
	v333 = v329
	v334 = v327
	v335 = v328
	v336 = v329
	v337 = int32(1)
	goto L38
L40:
	;
	v325 = v320
	v327 = int32(0)
	v328 = v323
	goto L39
L41:
	;
	v320 = int32(1)
	v323 = int32(0)
	goto L40
L42:
	;
	v222 = *(*int32)(unsafe.Add(mBase, uint32(v26)))
	if v222 == int32(0) {
		goto L76
	} else {
		goto L77
	}
L43:
	;
	v320 = v219
	v323 = int32(1)
	goto L40
L44:
	;
	if v123 == int32(0) {
		goto L42
	} else {
		goto L73
	}
L45:
	;
	if v123 == int32(-1) {
		goto L46
	} else {
		goto L47
	}
L46:
	;
	v134 = *(*int32)(unsafe.Add(mBase, uint32(v26)+4))
	if v134 == int32(0) {
		goto L41
	} else {
		goto L49
	}
L47:
	;
	if v123 == v52&int32(65535)+int32(-1) {
		goto L46
	} else {
		goto L48
	}
L48:
	;
	v219 = int32(0)
	goto L43
L49:
	;
	v137 = *(*int32)(unsafe.Add(mBase, uint32(v134)+16))
	if v137&int32(786432) == int32(262144) {
		goto L41
	} else {
		goto L50
	}
L50:
	;
	v142 = int32(0)
	v143 = *(*int32)(unsafe.Add(mBase, _consts[381]))
	if v143 == v142 {
		goto L55
	} else {
		goto L56
	}
L51:
	;
	v325 = v206
	v327 = int32(1)
	v328 = int32(0)
	goto L39
L52:
	;
	if base.Ui32(v202) < base.Ui32(v203) {
		goto L41
	} else {
		goto L72
	}
L53:
	;
	v186 = int32(-5)
	if base.Ui32(v186) < base.Ui32(v25) {
		goto L68
	} else {
		goto L69
	}
L54:
	;
	v173 = int32(1)
	if base.Ui32(v173) < base.Ui32(v25) {
		goto L64
	} else {
		goto L65
	}
L55:
	;
	if v25 < int32(0) {
		goto L53
	} else {
		goto L62
	}
L56:
	;
	if base.Ui32(v143) <= base.Ui32(l3) {
		goto L41
	} else {
		goto L57
	}
L57:
	;
	v147 = *(*int32)(unsafe.Add(mBase, uint32(v134)+12))
	v150 = l3 + v147 + int32(8)
	if int32(-1) < v25 {
		v172 = v150
		goto L54
	} else {
		goto L58
	}
L58:
	;
	v153 = int32(-5)
	if base.Ui32(v153) < base.Ui32(v25) {
		goto L59
	} else {
		goto L60
	}
L59:
	;
	v156 = v25
	goto L61
L60:
	;
	v156 = v153
	goto L61
L61:
	;
	v163 = *(*int32)(unsafe.Add(mBase, uint32((v156^int32(-1))<<(uint(int32(2))%32))+uint32(_consts[628])))
	v202 = v163
	v203 = v150
	goto L52
L62:
	;
	if base.Ui32(int32(8192)) < base.Ui32(l3) {
		goto L41
	} else {
		goto L63
	}
L63:
	;
	v168 = *(*int32)(unsafe.Add(mBase, uint32(v134)+12))
	v172 = l3 + v168 + int32(8)
	goto L54
L64:
	;
	v179 = v25
	goto L66
L65:
	;
	v179 = v173
	goto L66
L66:
	;
	if base.B2i32(base.Ui32(v179) <= base.Ui32(v137&int32(65535)))|base.B2i32(base.Ui32(int32(8192)) < base.Ui32(v172)) != int32(1) {
		v206 = v173
		goto L51
	} else {
		goto L67
	}
L67:
	;
	goto L41
L68:
	;
	v189 = v25
	goto L70
L69:
	;
	v189 = v186
	goto L70
L70:
	;
	v196 = *(*int32)(unsafe.Add(mBase, uint32((v189^int32(-1))<<(uint(int32(2))%32))+uint32(_consts[628])))
	if base.Ui32(v196) < base.Ui32(l3) {
		goto L41
	} else {
		goto L71
	}
L71:
	;
	v198 = *(*int32)(unsafe.Add(mBase, uint32(v134)+12))
	v202 = v196
	v203 = l3 + v198 + int32(8)
	goto L52
L72:
	;
	v206 = int32(1)
	goto L51
L73:
	;
	v213 = int32(0)
	if v123 == v213-v52&int32(65535) {
		goto L42
	} else {
		goto L74
	}
L74:
	;
	v219 = v213
	goto L43
L75:
	;
	v332 = v309
	v333 = v310
	v334 = v311
	v335 = v312
	v336 = int32(1)
	v337 = int32(0)
	goto L38
L76:
	;
	v305 = int32(0)
	v309 = v305
	v310 = v305
	v311 = v305
	v312 = int32(1)
	goto L75
L77:
	;
	v225 = *(*int32)(unsafe.Add(mBase, uint32(v222)+16))
	if v225&int32(786432) == int32(262144) {
		goto L76
	} else {
		goto L78
	}
L78:
	;
	v230 = int32(0)
	v231 = *(*int32)(unsafe.Add(mBase, _consts[381]))
	if v231 == v230 {
		goto L83
	} else {
		goto L84
	}
L79:
	;
	v309 = v296
	v310 = int32(1)
	v311 = int32(0)
	v312 = v298
	goto L75
L80:
	;
	if base.Ui32(v292) < base.Ui32(v291) {
		goto L76
	} else {
		goto L100
	}
L81:
	;
	v275 = int32(-5)
	if base.Ui32(v275) < base.Ui32(v25) {
		goto L96
	} else {
		goto L97
	}
L82:
	;
	v261 = int32(1)
	if base.Ui32(v261) < base.Ui32(v25) {
		goto L92
	} else {
		goto L93
	}
L83:
	;
	if v25 < int32(0) {
		goto L81
	} else {
		goto L90
	}
L84:
	;
	if base.Ui32(v231) <= base.Ui32(l3) {
		goto L76
	} else {
		goto L85
	}
L85:
	;
	v235 = *(*int32)(unsafe.Add(mBase, uint32(v222)+12))
	v238 = l3 + v235 + int32(8)
	if int32(-1) < v25 {
		v260 = v238
		goto L82
	} else {
		goto L86
	}
L86:
	;
	v241 = int32(-5)
	if base.Ui32(v241) < base.Ui32(v25) {
		goto L87
	} else {
		goto L88
	}
L87:
	;
	v244 = v25
	goto L89
L88:
	;
	v244 = v241
	goto L89
L89:
	;
	v251 = *(*int32)(unsafe.Add(mBase, uint32((v244^int32(-1))<<(uint(int32(2))%32))+uint32(_consts[628])))
	v291 = v238
	v292 = v251
	goto L80
L90:
	;
	if base.Ui32(int32(8192)) < base.Ui32(l3) {
		goto L76
	} else {
		goto L91
	}
L91:
	;
	v256 = *(*int32)(unsafe.Add(mBase, uint32(v222)+12))
	v260 = l3 + v256 + int32(8)
	goto L82
L92:
	;
	v267 = v25
	goto L94
L93:
	;
	v267 = v261
	goto L94
L94:
	;
	if base.B2i32(base.Ui32(v267) <= base.Ui32(v225&int32(65535)))|base.B2i32(base.Ui32(int32(8192)) < base.Ui32(v260)) == int32(1) {
		goto L76
	} else {
		goto L95
	}
L95:
	;
	v296 = int32(0)
	v298 = v261
	goto L79
L96:
	;
	v278 = v25
	goto L98
L97:
	;
	v278 = v275
	goto L98
L98:
	;
	v285 = *(*int32)(unsafe.Add(mBase, uint32((v278^int32(-1))<<(uint(int32(2))%32))+uint32(_consts[628])))
	if base.Ui32(v285) < base.Ui32(l3) {
		goto L76
	} else {
		goto L99
	}
L99:
	;
	v287 = *(*int32)(unsafe.Add(mBase, uint32(v222)+12))
	v291 = l3 + v287 + int32(8)
	v292 = v285
	goto L80
L100:
	;
	v296 = int32(0)
	v298 = int32(1)
	goto L79
L101:
	;
	v395 = base.B2i32(l4 == int32(0))
	if v395|(v121^int32(1)) != 0 {
		goto L127
	} else {
		goto L128
	}
L102:
	;
	v360 = base.B2i32(l4 != int32(0))
	if v360|v337 != int32(1) {
		goto L114
	} else {
		goto L115
	}
L103:
	;
	if v25 < int32(0) {
		goto L106
	} else {
		goto L107
	}
L104:
	;
	if base.Ui32(v339) <= base.Ui32(l3) {
		goto L102
	} else {
		goto L105
	}
L105:
	;
	goto L101
L106:
	;
	v347 = int32(-5)
	if base.Ui32(v347) < base.Ui32(v25) {
		goto L109
	} else {
		goto L110
	}
L107:
	;
	if base.Ui32(l3) <= base.Ui32(int32(8192)) {
		goto L101
	} else {
		goto L108
	}
L108:
	;
	goto L102
L109:
	;
	v350 = v25
	goto L111
L110:
	;
	v350 = v347
	goto L111
L111:
	;
	v357 = *(*int32)(unsafe.Add(mBase, uint32((v350^int32(-1))<<(uint(int32(2))%32))+uint32(_consts[628])))
	if base.Ui32(l3) <= base.Ui32(v357) {
		goto L101
	} else {
		goto L112
	}
L112:
	;
	goto L102
L113:
	;
	if v52&int32(196608) != int32(131072) {
		v380 = v123
		goto L119
	} else {
		goto L120
	}
L114:
	;
	F___quicklistInsertPlainNode(m, v20, v26, l2, l3, l4)
	mBase = m.M
	v368 = m.ExcPending
	if v368 != 0 {
		goto L15
	} else {
		goto L118
	}
L115:
	;
	if v360&v332 != 0 {
		goto L114
	} else {
		goto L116
	}
L116:
	;
	if v54 != int32(262144) {
		goto L113
	} else {
		goto L117
	}
L117:
	;
	goto L114
L118:
	;
	return
L119:
	;
	v381 = F__quicklistSplitNode(m, v26, v380, l4)
	mBase = m.M
	v382 = m.ExcPending
	if v382 != 0 {
		goto L15
	} else {
		goto L122
	}
L120:
	;
	F___quicklistDecompressNode(m, v26)
	mBase = m.M
	v374 = m.ExcPending
	if v374 != 0 {
		goto L15
	} else {
		goto L121
	}
L121:
	;
	v375 = *(*int32)(unsafe.Add(mBase, uint32(v26)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v26)+16)) = v375 | int32(1048576)
	v379 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
	v380 = v379
	goto L119
L122:
	;
	v384 = F___quicklistCreateNode(m, int32(1), l2, l3)
	mBase = m.M
	v385 = m.ExcPending
	if v385 != 0 {
		goto L15
	} else {
		goto L123
	}
L123:
	;
	F___quicklistInsertNode(m, v20, v26, v384, l4)
	mBase = m.M
	v387 = m.ExcPending
	if v387 != 0 {
		goto L15
	} else {
		goto L124
	}
L124:
	;
	F___quicklistInsertNode(m, v20, v384, v381, l4)
	mBase = m.M
	v389 = m.ExcPending
	if v389 != 0 {
		goto L15
	} else {
		goto L125
	}
L125:
	;
	v390 = *(*int32)(unsafe.Add(mBase, uint32(v20)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v20)+8)) = v390 + int32(1)
	return
L126:
	;
	v686 = *(*int32)(unsafe.Add(mBase, uint32(v20)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v20)+8)) = v686 + int32(1)
	*(*int64)(unsafe.Add(mBase, uint32(l0)+4)) = int64(0)
	return
L127:
	;
	v455 = v121 ^ int32(1)
	if base.B2i32(l4 != int32(0))|v455 != 0 {
		goto L141
	} else {
		goto L142
	}
L128:
	;
	if v52&int32(196608) != int32(131072) {
		goto L129
	} else {
		goto L130
	}
L129:
	;
	v432 = *(*int32)(unsafe.Add(mBase, uint32(v26)+8))
	v433 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v436 = F_lpInsertString(m, v432, l2, l3, v433, int32(1), int32(0))
	mBase = m.M
	v437 = m.ExcPending
	if v437 != 0 {
		goto L15
	} else {
		goto L137
	}
L130:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26)+16)) = v52 & int32(-3211265)
	v406 = *(*int32)(unsafe.Add(mBase, uint32(v26)+12))
	v407 = F_valkey_malloc(m, v406)
	mBase = m.M
	v408 = m.ExcPending
	if v408 != 0 {
		goto L15
	} else {
		goto L131
	}
L131:
	;
	v409 = *(*int32)(unsafe.Add(mBase, uint32(v26)+8))
	v412 = *(*int32)(unsafe.Add(mBase, uint32(v409)))
	v413 = *(*int32)(unsafe.Add(mBase, uint32(v26)+12))
	v414 = F_lzf_decompress(m, v409+int32(4), v412, v407, v413)
	mBase = m.M
	if v414 != 0 {
		goto L133
	} else {
		goto L134
	}
L132:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26)+16)) = v426 | int32(1048576)
	goto L129
L133:
	;
	F_valkey_free(m, v409)
	mBase = m.M
	v419 = m.ExcPending
	if v419 != 0 {
		goto L15
	} else {
		goto L136
	}
L134:
	;
	F_valkey_free(m, v407)
	mBase = m.M
	v416 = m.ExcPending
	if v416 != 0 {
		goto L15
	} else {
		goto L135
	}
L135:
	;
	v417 = *(*int32)(unsafe.Add(mBase, uint32(v26)+16))
	v426 = v417
	goto L132
L136:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26)+8)) = v407
	v421 = *(*int32)(unsafe.Add(mBase, uint32(v26)+16))
	v426 = v421&int32(-196609) | int32(65536)
	goto L132
L137:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26)+8)) = v436
	v439 = *(*int32)(unsafe.Add(mBase, uint32(v26)+16))
	v441 = v439 + int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v26)+16)) = uint16(v441)
	v443 = *(*int32)(unsafe.Add(mBase, uint32(v436)))
	goto L138
L138:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26)+12)) = v443
	v445 = *(*int32)(unsafe.Add(mBase, uint32(v26)+16))
	if v445&int32(1245184) != int32(1114112) {
		goto L126
	} else {
		goto L139
	}
L139:
	;
	v450 = F___quicklistCompressNode(m, v26)
	mBase = m.M
	v451 = m.ExcPending
	if v451 != 0 {
		goto L15
	} else {
		goto L140
	}
L140:
	;
	goto L126
L141:
	;
	if l4 == int32(0) {
		goto L155
	} else {
		goto L156
	}
L142:
	;
	if v52&int32(196608) != int32(131072) {
		goto L143
	} else {
		goto L144
	}
L143:
	;
	v490 = *(*int32)(unsafe.Add(mBase, uint32(v26)+8))
	v491 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v492 = int32(0)
	v494 = F_lpInsertString(m, v490, l2, l3, v491, v492, v492)
	mBase = m.M
	v495 = m.ExcPending
	if v495 != 0 {
		goto L15
	} else {
		goto L151
	}
L144:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26)+16)) = v52 & int32(-3211265)
	v464 = *(*int32)(unsafe.Add(mBase, uint32(v26)+12))
	v465 = F_valkey_malloc(m, v464)
	mBase = m.M
	v466 = m.ExcPending
	if v466 != 0 {
		goto L15
	} else {
		goto L145
	}
L145:
	;
	v467 = *(*int32)(unsafe.Add(mBase, uint32(v26)+8))
	v470 = *(*int32)(unsafe.Add(mBase, uint32(v467)))
	v471 = *(*int32)(unsafe.Add(mBase, uint32(v26)+12))
	v472 = F_lzf_decompress(m, v467+int32(4), v470, v465, v471)
	mBase = m.M
	if v472 != 0 {
		goto L147
	} else {
		goto L148
	}
L146:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26)+16)) = v484 | int32(1048576)
	goto L143
L147:
	;
	F_valkey_free(m, v467)
	mBase = m.M
	v477 = m.ExcPending
	if v477 != 0 {
		goto L15
	} else {
		goto L150
	}
L148:
	;
	F_valkey_free(m, v465)
	mBase = m.M
	v474 = m.ExcPending
	if v474 != 0 {
		goto L15
	} else {
		goto L149
	}
L149:
	;
	v475 = *(*int32)(unsafe.Add(mBase, uint32(v26)+16))
	v484 = v475
	goto L146
L150:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26)+8)) = v465
	v479 = *(*int32)(unsafe.Add(mBase, uint32(v26)+16))
	v484 = v479&int32(-196609) | int32(65536)
	goto L146
L151:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26)+8)) = v494
	v497 = *(*int32)(unsafe.Add(mBase, uint32(v26)+16))
	v499 = v497 + int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v26)+16)) = uint16(v499)
	v501 = *(*int32)(unsafe.Add(mBase, uint32(v494)))
	goto L152
L152:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26)+12)) = v501
	v503 = *(*int32)(unsafe.Add(mBase, uint32(v26)+16))
	if v503&int32(1245184) != int32(1114112) {
		goto L126
	} else {
		goto L153
	}
L153:
	;
	v508 = F___quicklistCompressNode(m, v26)
	mBase = m.M
	v509 = m.ExcPending
	if v509 != 0 {
		goto L15
	} else {
		goto L154
	}
L154:
	;
	goto L126
L155:
	;
	if l4 != 0 {
		goto L175
	} else {
		goto L176
	}
L156:
	;
	if v332&v455&v334 == int32(0) {
		goto L155
	} else {
		goto L157
	}
L157:
	;
	v516 = *(*int32)(unsafe.Add(mBase, uint32(v26)+4))
	if v516 == int32(0) {
		goto L158
	} else {
		goto L159
	}
L158:
	;
	v553 = *(*int32)(unsafe.Add(mBase, uint32(v516)+8))
	v554 = F_lpPrepend(m, v553, l2, l3)
	mBase = m.M
	v555 = m.ExcPending
	if v555 != 0 {
		goto L15
	} else {
		goto L167
	}
L159:
	;
	v519 = *(*int32)(unsafe.Add(mBase, uint32(v516)+16))
	if v519&int32(196608) != int32(131072) {
		goto L158
	} else {
		goto L160
	}
L160:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v516)+16)) = v519 & int32(-3211265)
	v527 = *(*int32)(unsafe.Add(mBase, uint32(v516)+12))
	v528 = F_valkey_malloc(m, v527)
	mBase = m.M
	v529 = m.ExcPending
	if v529 != 0 {
		goto L15
	} else {
		goto L161
	}
L161:
	;
	v530 = *(*int32)(unsafe.Add(mBase, uint32(v516)+8))
	v533 = *(*int32)(unsafe.Add(mBase, uint32(v530)))
	v534 = *(*int32)(unsafe.Add(mBase, uint32(v516)+12))
	v535 = F_lzf_decompress(m, v530+int32(4), v533, v528, v534)
	mBase = m.M
	if v535 != 0 {
		goto L163
	} else {
		goto L164
	}
L162:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v516)+16)) = v547 | int32(1048576)
	goto L158
L163:
	;
	F_valkey_free(m, v530)
	mBase = m.M
	v540 = m.ExcPending
	if v540 != 0 {
		goto L15
	} else {
		goto L166
	}
L164:
	;
	F_valkey_free(m, v528)
	mBase = m.M
	v537 = m.ExcPending
	if v537 != 0 {
		goto L15
	} else {
		goto L165
	}
L165:
	;
	v538 = *(*int32)(unsafe.Add(mBase, uint32(v516)+16))
	v547 = v538
	goto L162
L166:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v516)+8)) = v528
	v542 = *(*int32)(unsafe.Add(mBase, uint32(v516)+16))
	v547 = v542&int32(-196609) | int32(65536)
	goto L162
L167:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v516)+8)) = v554
	v557 = *(*int32)(unsafe.Add(mBase, uint32(v516)+16))
	v559 = v557 + int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v516)+16)) = uint16(v559)
	v561 = *(*int32)(unsafe.Add(mBase, uint32(v554)))
	goto L168
L168:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v516)+12)) = v561
	if v516 == int32(0) {
		goto L169
	} else {
		goto L170
	}
L169:
	;
	v572 = *(*int32)(unsafe.Add(mBase, uint32(v26)+16))
	if v572&int32(1245184) != int32(1114112) {
		goto L126
	} else {
		goto L173
	}
L170:
	;
	v565 = *(*int32)(unsafe.Add(mBase, uint32(v516)+16))
	if v565&int32(1245184) != int32(1114112) {
		goto L169
	} else {
		goto L171
	}
L171:
	;
	v570 = F___quicklistCompressNode(m, v516)
	mBase = m.M
	v571 = m.ExcPending
	if v571 != 0 {
		goto L15
	} else {
		goto L172
	}
L172:
	;
	goto L169
L173:
	;
	v577 = F___quicklistCompressNode(m, v26)
	mBase = m.M
	v578 = m.ExcPending
	if v578 != 0 {
		goto L15
	} else {
		goto L174
	}
L174:
	;
	goto L126
L175:
	;
	if v121 != 0 {
		goto L126
	} else {
		goto L190
	}
L176:
	;
	if v333&(v336&v455) == int32(0) {
		goto L175
	} else {
		goto L177
	}
L177:
	;
	v583 = *(*int32)(unsafe.Add(mBase, uint32(v26)))
	if v583 == int32(0) {
		goto L178
	} else {
		goto L179
	}
L178:
	;
	v597 = *(*int32)(unsafe.Add(mBase, uint32(v583)+8))
	v598 = F_lpAppend(m, v597, l2, l3)
	mBase = m.M
	v599 = m.ExcPending
	if v599 != 0 {
		goto L15
	} else {
		goto L182
	}
L179:
	;
	v586 = *(*int32)(unsafe.Add(mBase, uint32(v583)+16))
	if v586&int32(196608) != int32(131072) {
		goto L178
	} else {
		goto L180
	}
L180:
	;
	F___quicklistDecompressNode(m, v583)
	mBase = m.M
	v592 = m.ExcPending
	if v592 != 0 {
		goto L15
	} else {
		goto L181
	}
L181:
	;
	v593 = *(*int32)(unsafe.Add(mBase, uint32(v583)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v583)+16)) = v593 | int32(1048576)
	goto L178
L182:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v583)+8)) = v598
	v601 = *(*int32)(unsafe.Add(mBase, uint32(v583)+16))
	v603 = v601 + int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v583)+16)) = uint16(v603)
	v605 = *(*int32)(unsafe.Add(mBase, uint32(v598)))
	goto L183
L183:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v583)+12)) = v605
	if v583 == int32(0) {
		goto L184
	} else {
		goto L185
	}
L184:
	;
	v616 = *(*int32)(unsafe.Add(mBase, uint32(v26)+16))
	if v616&int32(1245184) != int32(1114112) {
		goto L126
	} else {
		goto L188
	}
L185:
	;
	v609 = *(*int32)(unsafe.Add(mBase, uint32(v583)+16))
	if v609&int32(1245184) != int32(1114112) {
		goto L184
	} else {
		goto L186
	}
L186:
	;
	v614 = F___quicklistCompressNode(m, v583)
	mBase = m.M
	v615 = m.ExcPending
	if v615 != 0 {
		goto L15
	} else {
		goto L187
	}
L187:
	;
	goto L184
L188:
	;
	v621 = F___quicklistCompressNode(m, v26)
	mBase = m.M
	v622 = m.ExcPending
	if v622 != 0 {
		goto L15
	} else {
		goto L189
	}
L189:
	;
	goto L126
L190:
	;
	if v335|v395|v334 != int32(1) {
		goto L192
	} else {
		goto L193
	}
L191:
	;
	if v52&int32(196608) != int32(131072) {
		v658 = v123
		goto L200
	} else {
		goto L201
	}
L192:
	;
	v631 = F_quicklistCreateNode(m)
	mBase = m.M
	v632 = m.ExcPending
	if v632 != 0 {
		goto L15
	} else {
		goto L195
	}
L193:
	;
	if base.B2i32(l4 != int32(0))|(v333|v337) != 0 {
		goto L191
	} else {
		goto L194
	}
L194:
	;
	goto L192
L195:
	;
	v634 = F_lpNew(m, int32(0))
	mBase = m.M
	v635 = m.ExcPending
	if v635 != 0 {
		goto L15
	} else {
		goto L196
	}
L196:
	;
	v636 = F_lpPrepend(m, v634, l2, l3)
	mBase = m.M
	v637 = m.ExcPending
	if v637 != 0 {
		goto L15
	} else {
		goto L197
	}
L197:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v631)+8)) = v636
	v639 = *(*int32)(unsafe.Add(mBase, uint32(v631)+16))
	v641 = v639 + int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v631)+16)) = uint16(v641)
	v643 = *(*int32)(unsafe.Add(mBase, uint32(v636)))
	goto L198
L198:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v631)+12)) = v643
	F___quicklistInsertNode(m, v20, v26, v631, l4)
	mBase = m.M
	v646 = m.ExcPending
	if v646 != 0 {
		goto L15
	} else {
		goto L199
	}
L199:
	;
	goto L126
L200:
	;
	v659 = F__quicklistSplitNode(m, v26, v658, l4)
	mBase = m.M
	v660 = m.ExcPending
	if v660 != 0 {
		goto L15
	} else {
		goto L203
	}
L201:
	;
	F___quicklistDecompressNode(m, v26)
	mBase = m.M
	v652 = m.ExcPending
	if v652 != 0 {
		goto L15
	} else {
		goto L202
	}
L202:
	;
	v653 = *(*int32)(unsafe.Add(mBase, uint32(v26)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v26)+16)) = v653 | int32(1048576)
	v657 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
	v658 = v657
	goto L200
L203:
	;
	v661 = *(*int32)(unsafe.Add(mBase, uint32(v659)+8))
	if l4 == int32(0) {
		goto L205
	} else {
		goto L206
	}
L204:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v659)+8)) = v668
	v670 = *(*int32)(unsafe.Add(mBase, uint32(v659)+16))
	v672 = v670 + int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v659)+16)) = uint16(v672)
	v674 = *(*int32)(unsafe.Add(mBase, uint32(v668)))
	goto L209
L205:
	;
	v666 = F_lpAppend(m, v661, l2, l3)
	mBase = m.M
	v667 = m.ExcPending
	if v667 != 0 {
		goto L15
	} else {
		goto L208
	}
L206:
	;
	v664 = F_lpPrepend(m, v661, l2, l3)
	mBase = m.M
	v665 = m.ExcPending
	if v665 != 0 {
		goto L15
	} else {
		goto L207
	}
L207:
	;
	v668 = v664
	goto L204
L208:
	;
	v668 = v666
	goto L204
L209:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v659)+12)) = v674
	F___quicklistInsertNode(m, v20, v26, v659, l4)
	mBase = m.M
	v677 = m.ExcPending
	if v677 != 0 {
		goto L15
	} else {
		goto L210
	}
L210:
	;
	v678 = F__quicklistMergeNodes(m, v20, v26)
	mBase = m.M
	v679 = m.ExcPending
	if v679 != 0 {
		goto L15
	} else {
		goto L211
	}
L211:
	;
	goto L126
L212:
	;
	v695 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v693)+12)) = v695
	*(*int64)(unsafe.Add(mBase, uint32(v693))) = int64(0)
	v699 = *(*int32)(unsafe.Add(mBase, uint32(v693)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v693)+16)) = v699&int32(-6291456) | int32(589824)
	v706 = F_lpNew(m, v695)
	mBase = m.M
	v707 = m.ExcPending
	if v707 != 0 {
		goto L15
	} else {
		goto L213
	}
L213:
	;
	v708 = F_lpPrepend(m, v706, l2, l3)
	mBase = m.M
	v709 = m.ExcPending
	if v709 != 0 {
		goto L15
	} else {
		goto L214
	}
L214:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v693)+8)) = v708
	if l4 == int32(0) {
		goto L216
	} else {
		goto L217
	}
L215:
	;
	v721 = *(*int32)(unsafe.Add(mBase, uint32(v20)+12))
	if v721 != 0 {
		goto L220
	} else {
		goto L221
	}
L216:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v693)+4)) = int32(0)
	v719 = *(*int32)(unsafe.Add(mBase, uint32(v20)))
	if v719 != 0 {
		goto L215
	} else {
		goto L219
	}
L217:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v693))) = int32(0)
	v715 = *(*int32)(unsafe.Add(mBase, uint32(v20)+4))
	if v715 != 0 {
		goto L215
	} else {
		goto L218
	}
L218:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+4)) = v693
	goto L215
L219:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20))) = v693
	goto L215
L220:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+12)) = v721 + int32(1)
	F___quicklistCompress(m, v20, v693)
	mBase = m.M
	v728 = m.ExcPending
	if v728 != 0 {
		goto L15
	} else {
		goto L222
	}
L221:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20))) = v693
	*(*int32)(unsafe.Add(mBase, uint32(v20)+4)) = v693
	goto L220
L222:
	;
	v729 = *(*int32)(unsafe.Add(mBase, uint32(v693)+16))
	v730 = int32(1)
	v731 = v729 + v730
	*(*uint16)(unsafe.Add(mBase, uint32(v693)+16)) = uint16(v731)
	v733 = *(*int32)(unsafe.Add(mBase, uint32(v20)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v20)+8)) = v733 + v730
	return
}
func F__quicklistListpackMerge(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
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
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v42 int32
	_ = v42
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
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
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
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
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v100 int32
	_ = v100
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v119 int32
	_ = v119
	if l1 == int32(0) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	if l2 == int32(0) {
		goto L10
	} else {
		goto L11
	}
L2:
	;
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	if v8&int32(196608) != int32(131072) {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+16)) = v8 & int32(-3211265)
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v17 = F_valkey_malloc(m, v16)
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	return int32(0)
L5:
	;
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v21)))
	v25 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v26 = F_lzf_decompress(m, v21+int32(4), v24, v17, v25)
	mBase = m.M
	if v26 != 0 {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	F_valkey_free(m, v21)
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		goto L4
	} else {
		goto L9
	}
L7:
	;
	F_valkey_free(m, v17)
	mBase = m.M
	v28 = m.ExcPending
	if v28 != 0 {
		goto L4
	} else {
		goto L8
	}
L8:
	;
	goto L1
L9:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+8)) = v17
	v32 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	*(*int32)(unsafe.Add(mBase, uint32(l1)+16)) = v32&int32(-196609) | int32(65536)
	goto L1
L10:
	;
	v72 = int32(8)
	v73 = l1 + v72
	v76 = F_lpMerge(m, v73, l2+v72)
	mBase = m.M
	v77 = m.ExcPending
	if v77 != 0 {
		goto L4
	} else {
		goto L19
	}
L11:
	;
	v42 = *(*int32)(unsafe.Add(mBase, uint32(l2)+16))
	if v42&int32(196608) != int32(131072) {
		goto L10
	} else {
		goto L12
	}
L12:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2)+16)) = v42 & int32(-3211265)
	v50 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	v51 = F_valkey_malloc(m, v50)
	mBase = m.M
	v52 = m.ExcPending
	if v52 != 0 {
		goto L4
	} else {
		goto L13
	}
L13:
	;
	v53 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v53)))
	v57 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	v58 = F_lzf_decompress(m, v53+int32(4), v56, v51, v57)
	mBase = m.M
	if v58 != 0 {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	F_valkey_free(m, v53)
	mBase = m.M
	v62 = m.ExcPending
	if v62 != 0 {
		goto L4
	} else {
		goto L17
	}
L15:
	;
	F_valkey_free(m, v51)
	mBase = m.M
	v60 = m.ExcPending
	if v60 != 0 {
		goto L4
	} else {
		goto L16
	}
L16:
	;
	goto L10
L17:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2)+8)) = v51
	v64 = *(*int32)(unsafe.Add(mBase, uint32(l2)+16))
	*(*int32)(unsafe.Add(mBase, uint32(l2)+16)) = v64&int32(-196609) | int32(65536)
	goto L10
L18:
	;
	v80 = *(*int32)(unsafe.Add(mBase, uint32(v73)))
	if v80 != 0 {
		goto L22
	} else {
		goto L23
	}
L19:
	;
	if v76 != 0 {
		goto L18
	} else {
		goto L20
	}
L20:
	;
	return int32(0)
L21:
	;
	v89 = *(*int32)(unsafe.Add(mBase, uint32(v86)+8))
	v90 = F_lpLength(m, v89)
	mBase = m.M
	v91 = m.ExcPending
	if v91 != 0 {
		goto L4
	} else {
		goto L30
	}
L22:
	;
	v82 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	if v82 != 0 {
		goto L24
	} else {
		goto L25
	}
L23:
	;
	v86 = l2
	v87 = l1
	goto L21
L24:
	;
	v83 = int32(0)
	goto L26
L25:
	;
	v83 = l2
	goto L26
L26:
	;
	if v82 != 0 {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	v85 = int32(0)
	goto L29
L28:
	;
	v85 = l1
	goto L29
L29:
	;
	v86 = v85
	v87 = v83
	goto L21
L30:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v86)+16)) = uint16(v90)
	v93 = *(*int32)(unsafe.Add(mBase, uint32(v86)+8))
	v94 = *(*int32)(unsafe.Add(mBase, uint32(v93)))
	goto L31
L31:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v86)+12)) = v94
	v96 = *(*int32)(unsafe.Add(mBase, uint32(v86)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v86)+16)) = v96 & int32(-1048577)
	v100 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v87)+18)))
	*(*int32)(unsafe.Add(mBase, uint32(v87)+16)) = v100 << (uint(int32(16)) % 32)
	F___quicklistDelNode(m, l0, v87)
	mBase = m.M
	v105 = m.ExcPending
	if v105 != 0 {
		goto L4
	} else {
		goto L32
	}
L32:
	;
	v106 = *(*int32)(unsafe.Add(mBase, uint32(v86)+16))
	if v106&int32(1048576) == int32(0) {
		goto L34
	} else {
		goto L35
	}
L33:
	;
	return v86
L34:
	;
	F___quicklistCompress(m, l0, v86)
	mBase = m.M
	v119 = m.ExcPending
	if v119 != 0 {
		goto L4
	} else {
		goto L38
	}
L35:
	;
	if v106&int32(196608) != int32(65536) {
		goto L33
	} else {
		goto L36
	}
L36:
	;
	v115 = F___quicklistCompressNode(m, v86)
	mBase = m.M
	v116 = m.ExcPending
	if v116 != 0 {
		goto L4
	} else {
		goto L37
	}
L37:
	;
	return v86
L38:
	;
	goto L33
}
func F__quicklistSplitNode(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
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
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
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
	var v62 int32
	_ = v62
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v11 = F_valkey_malloc(m, int32(20))
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		return int32(0)
	} else {
		*(*int64)(unsafe.Add(mBase, uint32(v11))) = int64(0)
		v17 = *(*int32)(unsafe.Add(mBase, uint32(v11)+16))
		v18 = F_valkey_malloc(m, v9)
		mBase = m.M
		v19 = m.ExcPending
		if v19 != 0 {
			return int32(0)
		} else {
			v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			if v9 == int32(0) {
				v24 = v18
			} else {
				v23 = F__emscripten_memcpy_bulkmem(m, v18, v20, v9)
				mBase = m.M
				v24 = v23
			}
			if int32(-1) < l1 {
				v29 = l1
			} else {
				v27 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+16)))
				v29 = v27 + l1
			}
			v31 = v29 + int32(1)
			if l2 != 0 {
				v33 = v31
			} else {
				v33 = int32(0)
			}
			if l2 != 0 {
				v35 = int32(-1)
			} else {
				v35 = v29
			}
			v36 = F_lpDeleteRange(m, v20, v33, v35)
			mBase = m.M
			v37 = m.ExcPending
			if v37 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v36
				v39 = F_lpLength(m, v36)
				mBase = m.M
				v40 = m.ExcPending
				if v40 != 0 {
					return int32(0)
				} else {
					*(*uint16)(unsafe.Add(mBase, uint32(l0)+16)) = uint16(v39)
					v42 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
					v43 = *(*int32)(unsafe.Add(mBase, uint32(v42)))
					*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v43
					if l2 != 0 {
						v46 = int32(0)
					} else {
						v46 = v29
					}
					if l2 != 0 {
						v48 = v31
					} else {
						v48 = int32(-1)
					}
					v49 = F_lpDeleteRange(m, v24, v46, v48)
					mBase = m.M
					v50 = m.ExcPending
					if v50 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v11)+8)) = v49
						v54 = F_lpLength(m, v49)
						mBase = m.M
						v55 = m.ExcPending
						if v55 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = v17&int32(-6291456) | v54&int32(65535) | int32(589824)
							v62 = *(*int32)(unsafe.Add(mBase, uint32(v49)))
							*(*int32)(unsafe.Add(mBase, uint32(v11)+12)) = v62
							return v11
						}
					}
				}
			}
		}
	}
}
func F_quicklistCompare(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v63 int32
	_ = v63
	var v78 int32
	_ = v78
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(v5)+16))
	if v6&int32(786432) != int32(262144) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	return v87
L2:
	;
	v82 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v83 = F_lpCompare(m, v82, l1, l2)
	mBase = m.M
	v86 = m.ExcPending
	if v86 != 0 {
		goto L21
	} else {
		goto L22
	}
L3:
	;
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v12 != l2 {
		v87 = int32(0)
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if base.Ui32(l2) < base.Ui32(int32(4)) {
		v38 = v14
		v39 = l1
		v40 = l2
		goto L8
	} else {
		goto L9
	}
L5:
	;
	return base.B2i32(v78 == int32(0))
L6:
	;
	v78 = int32(0)
	goto L5
L7:
	;
	v50 = v45
	v51 = v46
	v52 = v47
	goto L17
L8:
	;
	if v40 == int32(0) {
		goto L6
	} else {
		goto L15
	}
L9:
	;
	if (l1|v14)&int32(3) != 0 {
		v45 = v14
		v46 = l1
		v47 = l2
		goto L7
	} else {
		goto L10
	}
L10:
	;
	v22 = v14
	v23 = l1
	v24 = l2
	goto L11
L11:
	;
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v22)))
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v23)))
	if v27 != v28 {
		v45 = v22
		v46 = v23
		v47 = v24
		goto L7
	} else {
		goto L13
	}
L12:
	;
	v38 = v33
	v39 = v31
	v40 = v35
	goto L8
L13:
	;
	v30 = int32(4)
	v31 = v23 + v30
	v33 = v22 + v30
	v35 = v24 + int32(-4)
	if base.Ui32(int32(3)) < base.Ui32(v35) {
		v22 = v33
		v23 = v31
		v24 = v35
		goto L11
	} else {
		goto L14
	}
L14:
	;
	goto L12
L15:
	;
	v45 = v38
	v46 = v39
	v47 = v40
	goto L7
L16:
	;
	v78 = v55 - v56
	goto L5
L17:
	;
	v55 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v50))))
	v56 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v51))))
	if v55 != v56 {
		goto L16
	} else {
		goto L19
	}
L19:
	;
	v58 = int32(1)
	v63 = v52 + int32(-1)
	if v63 == int32(0) {
		goto L6
	} else {
		goto L20
	}
L20:
	;
	v50 = v50 + v58
	v51 = v51 + v58
	v52 = v63
	goto L17
L21:
	;
	return int32(0)
L22:
	;
	v87 = v83
	goto L1
}
func F_quicklistDelIndex(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
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
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	if v5&int32(786432) != int32(262144) {
		v12 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
		v13 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
		v14 = F_lpDelete(m, v12, v13, l2)
		mBase = m.M
		v15 = m.ExcPending
		if v15 != 0 {
			return
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(l1)+8)) = v14
			v17 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
			v19 = v17 + int32(-1)
			*(*uint16)(unsafe.Add(mBase, uint32(l1)+16)) = uint16(v19)
			if v17&int32(65535) != int32(1) {
				v27 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
				*(*int32)(unsafe.Add(mBase, uint32(l1)+12)) = v27
				v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v29 + int32(-1)
				return
			} else {
				F___quicklistDelNode(m, l0, l1)
				mBase = m.M
				v26 = m.ExcPending
				if v26 != 0 {
					return
				} else {
					v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
					*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v29 + int32(-1)
					return
				}
			}
		}
	} else {
		F___quicklistDelNode(m, l0, l1)
		mBase = m.M
		v11 = m.ExcPending
		if v11 != 0 {
			return
		} else {
			return
		}
	}
}
func F_quicklistDelRange(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v35 int64
	_ = v35
	var v40 int32
	_ = v40
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v60 int64
	_ = v60
	var v63 int32
	_ = v63
	var v67 int32
	_ = v67
	var v76 int64
	_ = v76
	var v78 int32
	_ = v78
	var v80 int64
	_ = v80
	var v83 int32
	_ = v83
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v100 int64
	_ = v100
	var v107 int64
	_ = v107
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v122 int32
	_ = v122
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v133 int32
	_ = v133
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v144 int32
	_ = v144
	var v148 int32
	_ = v148
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v156 int32
	_ = v156
	var v158 int32
	_ = v158
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v169 int32
	_ = v169
	var v171 int32
	_ = v171
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
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
	var v192 int32
	_ = v192
	var v194 int32
	_ = v194
	var v199 int32
	_ = v199
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v209 int32
	_ = v209
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v214 int32
	_ = v214
	var v217 int32
	_ = v217
	var v221 int32
	_ = v221
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v231 int32
	_ = v231
	var v233 int32
	_ = v233
	var v241 int32
	_ = v241
	if int32(1) <= l2 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	if l1 < int32(0) {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	return int32(0)
L3:
	;
	v34 = l1 ^ l1>>(uint(int32(31))%32)
	v35 = base.I64_extend_i32_s(v34)
	if base.Ui64(v35) < base.Ui64(base.I64_extend_i32_u(v29)) {
		goto L12
	} else {
		goto L13
	}
L4:
	;
	v25 = int32(0) - l1
	if base.Ui32(l2) < base.Ui32(v25) {
		goto L9
	} else {
		goto L10
	}
L5:
	;
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v21 = v20 - l1
	if base.Ui32(l2) < base.Ui32(v21) {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	v23 = l2
	goto L8
L7:
	;
	v23 = v21
	goto L8
L8:
	;
	v29 = v20
	v31 = v23
	goto L3
L9:
	;
	v27 = l2
	goto L11
L10:
	;
	v27 = v25
	goto L11
L11:
	;
	v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v29 = v28
	v31 = v27
	goto L3
L12:
	;
	v40 = int32(0)
	v45 = int32(-1)
	v48 = v29 + v45
	v51 = base.B2i32(base.Ui32(int32(base.Ui32(v48)>>(uint(int32(1))%32))) < base.Ui32(v34))
	if base.Ui32(int32(base.Ui32(v48)>>(uint(int32(1))%32))) < base.Ui32(v34) {
		goto L15
	} else {
		goto L16
	}
L13:
	;
	return int32(0)
L14:
	;
	return v241
L15:
	;
	v52 = int32(base.Ui32(l1) >> (uint(int32(31)) % 32))
	goto L17
L16:
	;
	v52 = base.B2i32(v45 < l1)
	goto L17
L17:
	;
	if v52 != 0 {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	v53 = v40
	goto L20
L19:
	;
	v53 = int32(4)
	goto L20
L20:
	;
	v55 = *(*int32)(unsafe.Add(mBase, uint32(l0+v53)))
	if v55 == int32(0) {
		v241 = v40
		goto L14
	} else {
		goto L21
	}
L21:
	;
	if base.Ui32(int32(base.Ui32(v48)>>(uint(int32(1))%32))) < base.Ui32(v34) {
		goto L22
	} else {
		goto L23
	}
L22:
	;
	v60 = base.I64_extend_i32_u(v48) - v35
	goto L24
L23:
	;
	v60 = v35
	goto L24
L24:
	;
	if v52 != 0 {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	v63 = int32(4)
	goto L27
L26:
	;
	v63 = int32(0)
	goto L27
L27:
	;
	v67 = v55
	v76 = int64(0)
	goto L29
L28:
	;
	v87 = F_valkey_malloc(m, int32(20))
	mBase = m.M
	v90 = m.ExcPending
	if v90 != 0 {
		goto L33
	} else {
		goto L34
	}
L29:
	;
	v78 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v67)+16)))
	v80 = v76 + base.I64_extend_i32_u(v78)
	if base.Ui64(v60) < base.Ui64(v80) {
		goto L28
	} else {
		goto L31
	}
L30:
	;
	return int32(0)
L31:
	;
	v83 = *(*int32)(unsafe.Add(mBase, uint32(v67+v63)))
	if v83 != 0 {
		v67 = v83
		v76 = v80
		goto L29
	} else {
		goto L32
	}
L32:
	;
	goto L30
L33:
	;
	return int32(0)
L34:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v87)+8)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v87))) = l0
	*(*int32)(unsafe.Add(mBase, uint32(v87)+16)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v87)+4)) = v67
	if base.Ui32(int32(base.Ui32(v48)>>(uint(int32(1))%32))) < base.Ui32(v34) {
		goto L35
	} else {
		goto L36
	}
L35:
	;
	v100 = base.I64_extend_i32_u(v29-v78) - v76
	goto L37
L36:
	;
	v100 = v76
	goto L37
L37:
	;
	if int32(-1) < l1 {
		goto L38
	} else {
		goto L39
	}
L38:
	;
	v107 = v35 - v100
	goto L40
L39:
	;
	v107 = v100 + (v35 ^ int64(-1))
	goto L40
L40:
	;
	v108 = base.I32_wrap_i64(v107)
	*(*int32)(unsafe.Add(mBase, uint32(v87)+12)) = v108
	v110 = *(*int32)(unsafe.Add(mBase, uint32(v67)+16))
	if v110&int32(1048576) == int32(0) {
		goto L42
	} else {
		goto L43
	}
L41:
	;
	F_valkey_free(m, v87)
	mBase = m.M
	v124 = m.ExcPending
	if v124 != 0 {
		goto L33
	} else {
		goto L47
	}
L42:
	;
	F___quicklistCompress(m, l0, v67)
	mBase = m.M
	v122 = m.ExcPending
	if v122 != 0 {
		goto L33
	} else {
		goto L46
	}
L43:
	;
	if v110&int32(196608) != int32(65536) {
		goto L41
	} else {
		goto L44
	}
L44:
	;
	v119 = F___quicklistCompressNode(m, v67)
	mBase = m.M
	v120 = m.ExcPending
	if v120 != 0 {
		goto L33
	} else {
		goto L45
	}
L45:
	;
	goto L41
L46:
	;
	goto L41
L47:
	;
	v125 = int32(1)
	if v31 == int32(0) {
		v241 = v125
		goto L14
	} else {
		goto L48
	}
L48:
	;
	v129 = v108
	v130 = v67
	v133 = v31
	goto L49
L49:
	;
	v141 = *(*int32)(unsafe.Add(mBase, uint32(v130)+4))
	if v129 != 0 {
		goto L57
	} else {
		goto L58
	}
L50:
	;
	v241 = v125
	goto L14
L51:
	;
	v233 = v133 - v231
	if v233 != 0 {
		v129 = int32(0)
		v130 = v141
		v133 = v233
		goto L49
	} else {
		goto L84
	}
L52:
	;
	if v162&int32(196608) != int32(131072) {
		goto L69
	} else {
		goto L70
	}
L53:
	;
	F___quicklistDelNode(m, l0, v130)
	mBase = m.M
	v171 = m.ExcPending
	if v171 != 0 {
		goto L33
	} else {
		goto L68
	}
L54:
	;
	if v162&int32(786432) != int32(262144) {
		goto L52
	} else {
		goto L67
	}
L55:
	;
	v158 = int32(0) - v129
	if base.Ui32(v133) < base.Ui32(v158) {
		goto L64
	} else {
		goto L65
	}
L56:
	;
	if base.Ui32(v129+v133) < base.Ui32(v152) {
		goto L61
	} else {
		goto L62
	}
L57:
	;
	if v129 < int32(0) {
		goto L55
	} else {
		goto L60
	}
L58:
	;
	v142 = *(*int32)(unsafe.Add(mBase, uint32(v130)+16))
	v144 = v142 & int32(65535)
	if base.Ui32(v133) < base.Ui32(v144) {
		v151 = v142
		v152 = v144
		goto L56
	} else {
		goto L59
	}
L59:
	;
	v169 = v144
	goto L53
L60:
	;
	v148 = *(*int32)(unsafe.Add(mBase, uint32(v130)+16))
	v151 = v148
	v152 = v148 & int32(65535)
	goto L56
L61:
	;
	v156 = v133
	goto L63
L62:
	;
	v156 = v152 - v129
	goto L63
L63:
	;
	v162 = v151
	v163 = v156
	goto L54
L64:
	;
	v160 = v133
	goto L66
L65:
	;
	v160 = v158
	goto L66
L66:
	;
	v161 = *(*int32)(unsafe.Add(mBase, uint32(v130)+16))
	v162 = v161
	v163 = v160
	goto L54
L67:
	;
	v169 = v163
	goto L53
L68:
	;
	v231 = v169
	goto L51
L69:
	;
	v205 = *(*int32)(unsafe.Add(mBase, uint32(v130)+8))
	v206 = F_lpDeleteRange(m, v205, v129, v163)
	mBase = m.M
	v207 = m.ExcPending
	if v207 != 0 {
		goto L33
	} else {
		goto L77
	}
L70:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v130)+16)) = v162 & int32(-3211265)
	v179 = *(*int32)(unsafe.Add(mBase, uint32(v130)+12))
	v180 = F_valkey_malloc(m, v179)
	mBase = m.M
	v181 = m.ExcPending
	if v181 != 0 {
		goto L33
	} else {
		goto L71
	}
L71:
	;
	v182 = *(*int32)(unsafe.Add(mBase, uint32(v130)+8))
	v185 = *(*int32)(unsafe.Add(mBase, uint32(v182)))
	v186 = *(*int32)(unsafe.Add(mBase, uint32(v130)+12))
	v187 = F_lzf_decompress(m, v182+int32(4), v185, v180, v186)
	mBase = m.M
	if v187 != 0 {
		goto L73
	} else {
		goto L74
	}
L72:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v130)+16)) = v199 | int32(1048576)
	goto L69
L73:
	;
	F_valkey_free(m, v182)
	mBase = m.M
	v192 = m.ExcPending
	if v192 != 0 {
		goto L33
	} else {
		goto L76
	}
L74:
	;
	F_valkey_free(m, v180)
	mBase = m.M
	v189 = m.ExcPending
	if v189 != 0 {
		goto L33
	} else {
		goto L75
	}
L75:
	;
	v190 = *(*int32)(unsafe.Add(mBase, uint32(v130)+16))
	v199 = v190
	goto L72
L76:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v130)+8)) = v180
	v194 = *(*int32)(unsafe.Add(mBase, uint32(v130)+16))
	v199 = v194&int32(-196609) | int32(65536)
	goto L72
L77:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v130)+8)) = v206
	v209 = *(*int32)(unsafe.Add(mBase, uint32(v206)))
	goto L78
L78:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v130)+12)) = v209
	v211 = *(*int32)(unsafe.Add(mBase, uint32(v130)+16))
	v212 = v211 - v163
	*(*uint16)(unsafe.Add(mBase, uint32(v130)+16)) = uint16(v212)
	v214 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v214 - v163
	v217 = *(*int32)(unsafe.Add(mBase, uint32(v130)+16))
	if v217&int32(65535) != 0 {
		goto L79
	} else {
		goto L80
	}
L79:
	;
	if v217&int32(1245184) != int32(1114112) {
		v231 = v163
		goto L51
	} else {
		goto L82
	}
L80:
	;
	F___quicklistDelNode(m, l0, v130)
	mBase = m.M
	v221 = m.ExcPending
	if v221 != 0 {
		goto L33
	} else {
		goto L81
	}
L81:
	;
	v231 = v163
	goto L51
L82:
	;
	v226 = F___quicklistCompressNode(m, v130)
	mBase = m.M
	v227 = m.ExcPending
	if v227 != 0 {
		goto L33
	} else {
		goto L83
	}
L83:
	;
	v231 = v163
	goto L51
L84:
	;
	goto L50
}
func F_quicklistDup(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v12 int64
	_ = v12
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int64
	_ = v43
	var v47 int32
	_ = v47
	var v53 int32
	_ = v53
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
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
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v98 int32
	_ = v98
	var v102 int32
	_ = v102
	var v107 int32
	_ = v107
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v8 = F_valkey_malloc(m, int32(20))
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
	v12 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v8))) = v12
	*(*int64)(unsafe.Add(mBase, uint32(v8+int32(8)))) = v12
	v18 = int32(18)
	v21 = v6 << (uint(v18) % 32) >> (uint(v18) % 32)
	v22 = int32(-5)
	if v22 < v21 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v25 = v21
	goto L5
L4:
	;
	v25 = v22
	goto L5
L5:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = v25&int32(16383) | v6&int32(268419072)
	v32 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v32 == int32(0) {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	return v8
L7:
	;
	v36 = v32
	goto L8
L8:
	;
	v41 = F_valkey_malloc(m, int32(20))
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L1
	} else {
		goto L10
	}
L9:
	;
	goto L6
L10:
	;
	v43 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v41)+8)) = v43
	*(*int64)(unsafe.Add(mBase, uint32(v41))) = v43
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v41)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v41)+16)) = v47&int32(-6291456) | int32(589824)
	v53 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v36)+18)))
	switch v53&int32(3) + int32(-1) {
	case 0:
		goto L12
	case 1:
		goto L13
	default:
		goto L11
	}
L11:
	;
	v82 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v41)+18)))
	v85 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v36)+16)))
	v86 = v82<<(uint(int32(16))%32) | v85
	*(*int32)(unsafe.Add(mBase, uint32(v41)+16)) = v86
	v88 = *(*int32)(unsafe.Add(mBase, uint32(v8)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v8)+8)) = v85 + v88
	v91 = *(*int32)(unsafe.Add(mBase, uint32(v36)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v41)+12)) = v91
	v93 = *(*int32)(unsafe.Add(mBase, uint32(v36)+16))
	v98 = v93&int32(196608) | v86&int32(-196609)
	*(*int32)(unsafe.Add(mBase, uint32(v41)+16)) = v98
	v102 = *(*int32)(unsafe.Add(mBase, uint32(v36)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v41)+16)) = v98&int32(-786433) | v102&int32(786432)
	v107 = *(*int32)(unsafe.Add(mBase, uint32(v8)+4))
	F___quicklistInsertNode(m, v8, v107, v41, int32(1))
	mBase = m.M
	v110 = m.ExcPending
	if v110 != 0 {
		goto L1
	} else {
		goto L22
	}
L12:
	;
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v36)+12))
	v71 = F_valkey_malloc(m, v70)
	mBase = m.M
	v72 = m.ExcPending
	if v72 != 0 {
		goto L1
	} else {
		goto L18
	}
L13:
	;
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v36)+8))
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v58)))
	v61 = v59 + int32(4)
	v62 = F_valkey_malloc(m, v61)
	mBase = m.M
	v63 = m.ExcPending
	if v63 != 0 {
		goto L1
	} else {
		goto L14
	}
L14:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v41)+8)) = v62
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v36)+8))
	if v61 == int32(0) {
		goto L16
	} else {
		goto L17
	}
L15:
	;
	goto L11
L16:
	;
	goto L15
L17:
	;
	v68 = F__emscripten_memcpy_bulkmem(m, v62, v65, v61)
	mBase = m.M
	goto L16
L18:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v41)+8)) = v71
	v74 = *(*int32)(unsafe.Add(mBase, uint32(v36)+8))
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v36)+12))
	if v75 == int32(0) {
		goto L20
	} else {
		goto L21
	}
L19:
	;
	goto L11
L20:
	;
	goto L19
L21:
	;
	v78 = F__emscripten_memcpy_bulkmem(m, v71, v74, v75)
	mBase = m.M
	goto L20
L22:
	;
	v111 = *(*int32)(unsafe.Add(mBase, uint32(v36)+4))
	if v111 != 0 {
		v36 = v111
		goto L8
	} else {
		goto L23
	}
L23:
	;
	goto L9
}
func F_quicklistGetIteratorEntryAtIdx(m *base.Module, l0 int32, l1 int64, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v16 int64
	_ = v16
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v29 int32
	_ = v29
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v42 int64
	_ = v42
	var v45 int32
	_ = v45
	var v56 int32
	_ = v56
	var v58 int64
	_ = v58
	var v60 int32
	_ = v60
	var v62 int64
	_ = v62
	var v65 int32
	_ = v65
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v94 int64
	_ = v94
	var v96 int64
	_ = v96
	var v101 int64
	_ = v101
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v109 int32
	_ = v109
	var v116 int32
	_ = v116
	v16 = l1 ^ l1>>(uint(int64(63))%64)
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if base.Ui64(base.I64_extend_i32_u(v17)) <= base.Ui64(v16) {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	return v116
L2:
	;
	v81 = F_valkey_malloc(m, int32(20))
	mBase = m.M
	v84 = m.ExcPending
	if v84 != 0 {
		goto L22
	} else {
		goto L23
	}
L3:
	;
	v116 = int32(0)
	goto L1
L4:
	;
	v20 = int32(0)
	v29 = v17 + int32(-1)
	v33 = base.B2i32(base.Ui64(base.I64_extend_i32_u(int32(base.Ui32(v29)>>(uint(int32(1))%32)))) < base.Ui64(v16))
	if base.Ui64(base.I64_extend_i32_u(int32(base.Ui32(v29)>>(uint(int32(1))%32)))) < base.Ui64(v16) {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	v34 = base.I32_wrap_i64(int64(base.Ui64(l1) >> (uint(int64(63)) % 64)))
	goto L7
L6:
	;
	v34 = base.B2i32(int64(-1) < l1)
	goto L7
L7:
	;
	if v34 != 0 {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v35 = v20
	goto L10
L9:
	;
	v35 = int32(4)
	goto L10
L10:
	;
	v37 = *(*int32)(unsafe.Add(mBase, uint32(l0+v35)))
	if v37 == int32(0) {
		v116 = v20
		goto L1
	} else {
		goto L11
	}
L11:
	;
	if base.Ui64(base.I64_extend_i32_u(int32(base.Ui32(v29)>>(uint(int32(1))%32)))) < base.Ui64(v16) {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	v42 = base.I64_extend_i32_u(v29) - v16
	goto L14
L13:
	;
	v42 = v16
	goto L14
L14:
	;
	if v34 != 0 {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	v45 = int32(4)
	goto L17
L16:
	;
	v45 = int32(0)
	goto L17
L17:
	;
	v56 = v37
	v58 = int64(0)
	goto L18
L18:
	;
	v60 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v56)+16)))
	v62 = v58 + base.I64_extend_i32_u(v60)
	if base.Ui64(v42) < base.Ui64(v62) {
		goto L2
	} else {
		goto L20
	}
L19:
	;
	goto L3
L20:
	;
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v56+v45)))
	if v65 != 0 {
		v56 = v65
		v58 = v62
		goto L18
	} else {
		goto L21
	}
L21:
	;
	goto L19
L22:
	;
	return int32(0)
L23:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v81)+8)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v81))) = l0
	*(*int32)(unsafe.Add(mBase, uint32(v81)+16)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v81)+4)) = v56
	if base.Ui64(base.I64_extend_i32_u(int32(base.Ui32(v29)>>(uint(int32(1))%32)))) < base.Ui64(v16) {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	v94 = base.I64_extend_i32_u(v17-v60) - v58
	goto L26
L25:
	;
	v94 = v58
	goto L26
L26:
	;
	v96 = int64(-1)
	if v96 < l1 {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	v101 = v16 - v94
	goto L29
L28:
	;
	v101 = v94 + (v16 ^ v96)
	goto L29
L29:
	;
	*(*uint32)(unsafe.Add(mBase, uint32(v81)+12)) = uint32(v101)
	v103 = F_quicklistNext(m, v81, l2)
	mBase = m.M
	v104 = m.ExcPending
	if v104 != 0 {
		goto L22
	} else {
		goto L30
	}
L30:
	;
	if v103 != 0 {
		v116 = v81
		goto L1
	} else {
		goto L31
	}
L31:
	;
	F__serverAssert(m, int32(_a1113), int32(_a1112), int32(1420))
	mBase = m.M
	v109 = m.ExcPending
	if v109 != 0 {
		goto L22
	} else {
		goto L32
	}
L32:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_quicklistGetLzf(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v7 int32
	_ = v7
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v3 + int32(4)
	v7 = *(*int32)(unsafe.Add(mBase, uint32(v3)))
	return v7
}
func F_quicklistNext(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v15 int64
	_ = v15
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v36 int32
	_ = v36
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
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
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v106 int32
	_ = v106
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v118 int32
	_ = v118
	var v121 int32
	_ = v121
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v128 int32
	_ = v128
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v138 int32
	_ = v138
	var v142 int32
	_ = v142
	var v145 int32
	_ = v145
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
	var v160 int32
	_ = v160
	var v164 int32
	_ = v164
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v173 int32
	_ = v173
	var v178 int64
	_ = v178
	var v184 int32
	_ = v184
	var v186 int32
	_ = v186
	var v191 int32
	_ = v191
	v3 = int32(0)
	v9 = m.G0
	v11 = v9 - int32(16)
	m.G0 = v11
	*(*int64)(unsafe.Add(mBase, uint32(l1)+16)) = int64(-123456789)
	v15 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(l1)+8)) = v15
	*(*int64)(unsafe.Add(mBase, uint32(l1))) = v15
	*(*int64)(unsafe.Add(mBase, uint32(l1)+24)) = int64(530242871224172544)
	if l0 == v3 {
		v191 = v3
		goto L1
	} else {
		goto L2
	}
L1:
	;
	m.G0 = v11 + int32(16)
	return v191
L2:
	;
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v24
	v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v26
	if v26 == int32(0) {
		v191 = v3
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v36 = v26
	goto L4
L4:
	;
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v36)+16))
	v42 = v40 & int32(786432)
	v43 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v43 != 0 {
		goto L10
	} else {
		goto L11
	}
L5:
	;
	v191 = v173
	goto L1
L6:
	;
	v145 = *(*int32)(unsafe.Add(mBase, uint32(v142)+16))
	if v145&int32(1048576) == int32(0) {
		goto L36
	} else {
		goto L37
	}
L7:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+12)) = int32(0)
	v135 = F_lpGetValue(m, v113, v11+int32(12), l1+int32(16))
	mBase = m.M
	v136 = m.ExcPending
	if v136 != 0 {
		goto L14
	} else {
		goto L34
	}
L8:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v82
	*(*int32)(unsafe.Add(mBase, uint32(l1)+8)) = v82
	v121 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l1)+28)) = v121
	if v82 == int32(0) {
		v142 = v81
		goto L6
	} else {
		goto L33
	}
L9:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+28)) = v112
	*(*int32)(unsafe.Add(mBase, uint32(l1)+8)) = v113
	if v113 != 0 {
		goto L7
	} else {
		goto L32
	}
L10:
	;
	if v42 != int32(262144) {
		goto L23
	} else {
		goto L24
	}
L11:
	;
	if v40&int32(196608) != int32(131072) {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	v81 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v82 = *(*int32)(unsafe.Add(mBase, uint32(v81)+8))
	if v42 == int32(262144) {
		goto L8
	} else {
		goto L21
	}
L13:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v36)+16)) = v40 & int32(-3211265)
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v36)+12))
	v52 = F_valkey_malloc(m, v51)
	mBase = m.M
	v55 = m.ExcPending
	if v55 != 0 {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	return int32(0)
L15:
	;
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v36)+8))
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v56)))
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v36)+12))
	v61 = F_lzf_decompress(m, v56+int32(4), v59, v52, v60)
	mBase = m.M
	if v61 != 0 {
		goto L17
	} else {
		goto L18
	}
L16:
	;
	v73 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v74 = *(*int32)(unsafe.Add(mBase, uint32(v73)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v73)+16)) = v74 | int32(1048576)
	goto L12
L17:
	;
	F_valkey_free(m, v56)
	mBase = m.M
	v65 = m.ExcPending
	if v65 != 0 {
		goto L14
	} else {
		goto L20
	}
L18:
	;
	F_valkey_free(m, v52)
	mBase = m.M
	v63 = m.ExcPending
	if v63 != 0 {
		goto L14
	} else {
		goto L19
	}
L19:
	;
	goto L16
L20:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v36)+8)) = v52
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v36)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v36)+16)) = v67&int32(-196609) | int32(65536)
	goto L16
L21:
	;
	v85 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v86 = F_lpSeek(m, v82, v85)
	mBase = m.M
	v87 = m.ExcPending
	if v87 != 0 {
		goto L14
	} else {
		goto L22
	}
L22:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v86
	v89 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v112 = v89
	v113 = v86
	goto L9
L23:
	;
	v98 = *(*int32)(unsafe.Add(mBase, uint32(v36)+8))
	v101 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v101 != 0 {
		goto L25
	} else {
		goto L26
	}
L24:
	;
	v92 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v92
	*(*int32)(unsafe.Add(mBase, uint32(l1)+8)) = v92
	v96 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l1)+28)) = v96
	v142 = v36
	goto L6
L25:
	;
	v102 = int32(960)
	goto L27
L26:
	;
	v102 = int32(961)
	goto L27
L27:
	;
	v103 = m.T0[v102].(func(*base.Module, int32, int32) int32)(m, v98, v43)
	mBase = m.M
	v104 = m.ExcPending
	if v104 != 0 {
		goto L14
	} else {
		goto L28
	}
L28:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v103
	v106 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v101 != 0 {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	v109 = int32(-1)
	goto L31
L30:
	;
	v109 = int32(1)
	goto L31
L31:
	;
	v110 = v106 + v109
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v110
	v112 = v110
	v113 = v103
	goto L9
L32:
	;
	v118 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v142 = v118
	goto L6
L33:
	;
	v125 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v126 = *(*int32)(unsafe.Add(mBase, uint32(v125)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l1)+12)) = v126
	v128 = *(*int32)(unsafe.Add(mBase, uint32(v125)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l1)+24)) = v128
	v191 = int32(1)
	goto L1
L34:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+12)) = v135
	v138 = *(*int32)(unsafe.Add(mBase, uint32(v11)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l1)+24)) = v138
	v191 = int32(1)
	goto L1
L35:
	;
	v159 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	switch v159 {
	case 0:
		goto L44
	case 1:
		goto L43
	default:
		goto L41
	}
L36:
	;
	v156 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	F___quicklistCompress(m, v156, v142)
	mBase = m.M
	v158 = m.ExcPending
	if v158 != 0 {
		goto L14
	} else {
		goto L40
	}
L37:
	;
	if v145&int32(196608) != int32(65536) {
		goto L35
	} else {
		goto L38
	}
L38:
	;
	v154 = F___quicklistCompressNode(m, v142)
	mBase = m.M
	v155 = m.ExcPending
	if v155 != 0 {
		goto L14
	} else {
		goto L39
	}
L39:
	;
	goto L35
L40:
	;
	goto L35
L41:
	;
	v173 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v173
	*(*int64)(unsafe.Add(mBase, uint32(l1)+16)) = int64(-123456789)
	v178 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(l1)+8)) = v178
	*(*int64)(unsafe.Add(mBase, uint32(l1))) = v178
	*(*int64)(unsafe.Add(mBase, uint32(l1)+24)) = int64(530242871224172544)
	v184 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v184
	v186 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v186
	if v186 != 0 {
		v36 = v186
		goto L4
	} else {
		goto L45
	}
L42:
	;
	v168 = *(*int32)(unsafe.Add(mBase, uint32(v167)))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v166
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v168
	goto L41
L43:
	;
	v164 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v166 = int32(-1)
	v167 = v164
	goto L42
L44:
	;
	v160 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v166 = int32(0)
	v167 = v160 + int32(4)
	goto L42
L45:
	;
	goto L5
}
func F_quicklistNodeExceedsLimit(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v16 int32
	_ = v16
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	if int32(-1) < l0 {
		v21 = int32(1)
		if base.Ui32(v21) < base.Ui32(l0) {
			v24 = l0
		} else {
			v24 = v21
		}
		return base.B2i32(base.Ui32(int32(8192)) < base.Ui32(l1)) | base.B2i32(base.Ui32(v24) < base.Ui32(l2))
	} else {
		v6 = int32(-5)
		if base.Ui32(v6) < base.Ui32(l0) {
			v9 = l0
		} else {
			v9 = v6
		}
		v16 = *(*int32)(unsafe.Add(mBase, uint32((v9^int32(-1))<<(uint(int32(2))%32))+uint32(_consts[628])))
		return base.B2i32(base.Ui32(v16) < base.Ui32(l1))
	}
}
func F_quicklistPopCustom(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32 {
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
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v57 int32
	_ = v57
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
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
	var v83 int32
	_ = v83
	var v91 int64
	_ = v91
	var v93 int32
	_ = v93
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v118 int32
	_ = v118
	var v130 int32
	_ = v130
	var v140 int32
	_ = v140
	v7 = int32(0)
	v10 = m.G0
	v12 = v10 - int32(16)
	m.G0 = v12
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v15 == v7 {
		v130 = v7
		m.G0 = v12 + int32(16)
		return v130
	} else {
		if l2 == int32(0) {
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(l2))) = int32(0)
		}
		if l3 == int32(0) {
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(l3))) = int32(0)
		}
		if l4 == int32(0) {
		} else {
			*(*int64)(unsafe.Add(mBase, uint32(l4))) = int64(-123456789)
		}
		if l1 != 0 {
			if l1 != int32(-1) {
				v130 = v7
				m.G0 = v12 + int32(16)
				return v130
			} else {
				v33 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
				if v33 == int32(0) {
					v130 = v7
					m.G0 = v12 + int32(16)
					return v130
				} else {
					v36 = v33
					v37 = *(*int32)(unsafe.Add(mBase, uint32(v36)+16))
					if v37&int32(196608) == int32(131072) {
						F__serverAssert(m, int32(_a1114), int32(_a1112), int32(1518))
						mBase = m.M
						v140 = m.ExcPending
						if v140 != 0 {
							return int32(0)
						} else {
							F_abort(m)
							mBase = m.M
							base.Wasm_trap_unreachable()
							for {
							}
						}
					} else {
						if v37&int32(786432) != int32(262144) {
							v62 = *(*int32)(unsafe.Add(mBase, uint32(v36)+8))
							if l1 != 0 {
								v65 = int32(-1)
							} else {
								v65 = int32(0)
							}
							v66 = F_lpSeek(m, v62, v65)
							mBase = m.M
							v67 = m.ExcPending
							if v67 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v12)+12)) = v66
								v71 = F_lpGetValue(m, v66, v12+int32(8), v12)
								mBase = m.M
								v72 = m.ExcPending
								if v72 != 0 {
									return int32(0)
								} else {
									if v71 == int32(0) {
										if l2 == int32(0) {
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(l2))) = int32(0)
										}
										if l4 == int32(0) {
										} else {
											v91 = *(*int64)(unsafe.Add(mBase, uint32(v12)))
											*(*int64)(unsafe.Add(mBase, uint32(l4))) = v91
										}
										v93 = *(*int32)(unsafe.Add(mBase, uint32(v36)+16))
										if v93&int32(786432) != int32(262144) {
											v100 = *(*int32)(unsafe.Add(mBase, uint32(v36)+8))
											v103 = F_lpDelete(m, v100, v66, v12+int32(12))
											mBase = m.M
											v104 = m.ExcPending
											if v104 != 0 {
												return int32(0)
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(v36)+8)) = v103
												v106 = *(*int32)(unsafe.Add(mBase, uint32(v36)+16))
												v108 = v106 + int32(-1)
												*(*uint16)(unsafe.Add(mBase, uint32(v36)+16)) = uint16(v108)
												if v106&int32(65535) != int32(1) {
													v116 = *(*int32)(unsafe.Add(mBase, uint32(v103)))
													*(*int32)(unsafe.Add(mBase, uint32(v36)+12)) = v116
													v118 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
													*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v118 + int32(-1)
													v130 = int32(1)
													m.G0 = v12 + int32(16)
													return v130
												} else {
													F___quicklistDelNode(m, l0, v36)
													mBase = m.M
													v115 = m.ExcPending
													if v115 != 0 {
														return int32(0)
													} else {
														v118 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
														*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v118 + int32(-1)
														v130 = int32(1)
														m.G0 = v12 + int32(16)
														return v130
													}
												}
											}
										} else {
											F___quicklistDelNode(m, l0, v36)
											mBase = m.M
											v99 = m.ExcPending
											if v99 != 0 {
												return int32(0)
											} else {
												v130 = int32(1)
												m.G0 = v12 + int32(16)
												return v130
											}
										}
									} else {
										if l2 == int32(0) {
											if l3 == int32(0) {
											} else {
												v83 = *(*int32)(unsafe.Add(mBase, uint32(v12)+8))
												*(*int32)(unsafe.Add(mBase, uint32(l3))) = v83
											}
											v93 = *(*int32)(unsafe.Add(mBase, uint32(v36)+16))
											if v93&int32(786432) != int32(262144) {
												v100 = *(*int32)(unsafe.Add(mBase, uint32(v36)+8))
												v103 = F_lpDelete(m, v100, v66, v12+int32(12))
												mBase = m.M
												v104 = m.ExcPending
												if v104 != 0 {
													return int32(0)
												} else {
													*(*int32)(unsafe.Add(mBase, uint32(v36)+8)) = v103
													v106 = *(*int32)(unsafe.Add(mBase, uint32(v36)+16))
													v108 = v106 + int32(-1)
													*(*uint16)(unsafe.Add(mBase, uint32(v36)+16)) = uint16(v108)
													if v106&int32(65535) != int32(1) {
														v116 = *(*int32)(unsafe.Add(mBase, uint32(v103)))
														*(*int32)(unsafe.Add(mBase, uint32(v36)+12)) = v116
														v118 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
														*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v118 + int32(-1)
														v130 = int32(1)
														m.G0 = v12 + int32(16)
														return v130
													} else {
														F___quicklistDelNode(m, l0, v36)
														mBase = m.M
														v115 = m.ExcPending
														if v115 != 0 {
															return int32(0)
														} else {
															v118 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
															*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v118 + int32(-1)
															v130 = int32(1)
															m.G0 = v12 + int32(16)
															return v130
														}
													}
												}
											} else {
												F___quicklistDelNode(m, l0, v36)
												mBase = m.M
												v99 = m.ExcPending
												if v99 != 0 {
													return int32(0)
												} else {
													v130 = int32(1)
													m.G0 = v12 + int32(16)
													return v130
												}
											}
										} else {
											v77 = *(*int32)(unsafe.Add(mBase, uint32(v12)+8))
											v78 = m.T0[l5].(func(*base.Module, int32, int32) int32)(m, v71, v77)
											mBase = m.M
											v79 = m.ExcPending
											if v79 != 0 {
												return int32(0)
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(l2))) = v78
												if l3 == int32(0) {
												} else {
													v83 = *(*int32)(unsafe.Add(mBase, uint32(v12)+8))
													*(*int32)(unsafe.Add(mBase, uint32(l3))) = v83
												}
												v93 = *(*int32)(unsafe.Add(mBase, uint32(v36)+16))
												if v93&int32(786432) != int32(262144) {
													v100 = *(*int32)(unsafe.Add(mBase, uint32(v36)+8))
													v103 = F_lpDelete(m, v100, v66, v12+int32(12))
													mBase = m.M
													v104 = m.ExcPending
													if v104 != 0 {
														return int32(0)
													} else {
														*(*int32)(unsafe.Add(mBase, uint32(v36)+8)) = v103
														v106 = *(*int32)(unsafe.Add(mBase, uint32(v36)+16))
														v108 = v106 + int32(-1)
														*(*uint16)(unsafe.Add(mBase, uint32(v36)+16)) = uint16(v108)
														if v106&int32(65535) != int32(1) {
															v116 = *(*int32)(unsafe.Add(mBase, uint32(v103)))
															*(*int32)(unsafe.Add(mBase, uint32(v36)+12)) = v116
															v118 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
															*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v118 + int32(-1)
															v130 = int32(1)
															m.G0 = v12 + int32(16)
															return v130
														} else {
															F___quicklistDelNode(m, l0, v36)
															mBase = m.M
															v115 = m.ExcPending
															if v115 != 0 {
																return int32(0)
															} else {
																v118 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
																*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v118 + int32(-1)
																v130 = int32(1)
																m.G0 = v12 + int32(16)
																return v130
															}
														}
													}
												} else {
													F___quicklistDelNode(m, l0, v36)
													mBase = m.M
													v99 = m.ExcPending
													if v99 != 0 {
														return int32(0)
													} else {
														v130 = int32(1)
														m.G0 = v12 + int32(16)
														return v130
													}
												}
											}
										}
									}
								}
							}
						} else {
							if l2 == int32(0) {
								if l3 == int32(0) {
								} else {
									v57 = *(*int32)(unsafe.Add(mBase, uint32(v36)+12))
									*(*int32)(unsafe.Add(mBase, uint32(l3))) = v57
								}
								F_quicklistDelIndex(m, l0, v36, int32(0))
								mBase = m.M
								v61 = m.ExcPending
								if v61 != 0 {
									return int32(0)
								} else {
									v130 = int32(1)
									m.G0 = v12 + int32(16)
									return v130
								}
							} else {
								v48 = *(*int32)(unsafe.Add(mBase, uint32(v36)+8))
								v49 = *(*int32)(unsafe.Add(mBase, uint32(v36)+12))
								v50 = m.T0[l5].(func(*base.Module, int32, int32) int32)(m, v48, v49)
								mBase = m.M
								v53 = m.ExcPending
								if v53 != 0 {
									return int32(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(l2))) = v50
									if l3 == int32(0) {
									} else {
										v57 = *(*int32)(unsafe.Add(mBase, uint32(v36)+12))
										*(*int32)(unsafe.Add(mBase, uint32(l3))) = v57
									}
									F_quicklistDelIndex(m, l0, v36, int32(0))
									mBase = m.M
									v61 = m.ExcPending
									if v61 != 0 {
										return int32(0)
									} else {
										v130 = int32(1)
										m.G0 = v12 + int32(16)
										return v130
									}
								}
							}
						}
					}
				}
			}
		} else {
			v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			if v30 != 0 {
				v36 = v30
				v37 = *(*int32)(unsafe.Add(mBase, uint32(v36)+16))
				if v37&int32(196608) == int32(131072) {
					F__serverAssert(m, int32(_a1114), int32(_a1112), int32(1518))
					mBase = m.M
					v140 = m.ExcPending
					if v140 != 0 {
						return int32(0)
					} else {
						F_abort(m)
						mBase = m.M
						base.Wasm_trap_unreachable()
						for {
						}
					}
				} else {
					if v37&int32(786432) != int32(262144) {
						v62 = *(*int32)(unsafe.Add(mBase, uint32(v36)+8))
						if l1 != 0 {
							v65 = int32(-1)
						} else {
							v65 = int32(0)
						}
						v66 = F_lpSeek(m, v62, v65)
						mBase = m.M
						v67 = m.ExcPending
						if v67 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v12)+12)) = v66
							v71 = F_lpGetValue(m, v66, v12+int32(8), v12)
							mBase = m.M
							v72 = m.ExcPending
							if v72 != 0 {
								return int32(0)
							} else {
								if v71 == int32(0) {
									if l2 == int32(0) {
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(l2))) = int32(0)
									}
									if l4 == int32(0) {
									} else {
										v91 = *(*int64)(unsafe.Add(mBase, uint32(v12)))
										*(*int64)(unsafe.Add(mBase, uint32(l4))) = v91
									}
									v93 = *(*int32)(unsafe.Add(mBase, uint32(v36)+16))
									if v93&int32(786432) != int32(262144) {
										v100 = *(*int32)(unsafe.Add(mBase, uint32(v36)+8))
										v103 = F_lpDelete(m, v100, v66, v12+int32(12))
										mBase = m.M
										v104 = m.ExcPending
										if v104 != 0 {
											return int32(0)
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(v36)+8)) = v103
											v106 = *(*int32)(unsafe.Add(mBase, uint32(v36)+16))
											v108 = v106 + int32(-1)
											*(*uint16)(unsafe.Add(mBase, uint32(v36)+16)) = uint16(v108)
											if v106&int32(65535) != int32(1) {
												v116 = *(*int32)(unsafe.Add(mBase, uint32(v103)))
												*(*int32)(unsafe.Add(mBase, uint32(v36)+12)) = v116
												v118 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
												*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v118 + int32(-1)
												v130 = int32(1)
												m.G0 = v12 + int32(16)
												return v130
											} else {
												F___quicklistDelNode(m, l0, v36)
												mBase = m.M
												v115 = m.ExcPending
												if v115 != 0 {
													return int32(0)
												} else {
													v118 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
													*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v118 + int32(-1)
													v130 = int32(1)
													m.G0 = v12 + int32(16)
													return v130
												}
											}
										}
									} else {
										F___quicklistDelNode(m, l0, v36)
										mBase = m.M
										v99 = m.ExcPending
										if v99 != 0 {
											return int32(0)
										} else {
											v130 = int32(1)
											m.G0 = v12 + int32(16)
											return v130
										}
									}
								} else {
									if l2 == int32(0) {
										if l3 == int32(0) {
										} else {
											v83 = *(*int32)(unsafe.Add(mBase, uint32(v12)+8))
											*(*int32)(unsafe.Add(mBase, uint32(l3))) = v83
										}
										v93 = *(*int32)(unsafe.Add(mBase, uint32(v36)+16))
										if v93&int32(786432) != int32(262144) {
											v100 = *(*int32)(unsafe.Add(mBase, uint32(v36)+8))
											v103 = F_lpDelete(m, v100, v66, v12+int32(12))
											mBase = m.M
											v104 = m.ExcPending
											if v104 != 0 {
												return int32(0)
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(v36)+8)) = v103
												v106 = *(*int32)(unsafe.Add(mBase, uint32(v36)+16))
												v108 = v106 + int32(-1)
												*(*uint16)(unsafe.Add(mBase, uint32(v36)+16)) = uint16(v108)
												if v106&int32(65535) != int32(1) {
													v116 = *(*int32)(unsafe.Add(mBase, uint32(v103)))
													*(*int32)(unsafe.Add(mBase, uint32(v36)+12)) = v116
													v118 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
													*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v118 + int32(-1)
													v130 = int32(1)
													m.G0 = v12 + int32(16)
													return v130
												} else {
													F___quicklistDelNode(m, l0, v36)
													mBase = m.M
													v115 = m.ExcPending
													if v115 != 0 {
														return int32(0)
													} else {
														v118 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
														*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v118 + int32(-1)
														v130 = int32(1)
														m.G0 = v12 + int32(16)
														return v130
													}
												}
											}
										} else {
											F___quicklistDelNode(m, l0, v36)
											mBase = m.M
											v99 = m.ExcPending
											if v99 != 0 {
												return int32(0)
											} else {
												v130 = int32(1)
												m.G0 = v12 + int32(16)
												return v130
											}
										}
									} else {
										v77 = *(*int32)(unsafe.Add(mBase, uint32(v12)+8))
										v78 = m.T0[l5].(func(*base.Module, int32, int32) int32)(m, v71, v77)
										mBase = m.M
										v79 = m.ExcPending
										if v79 != 0 {
											return int32(0)
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(l2))) = v78
											if l3 == int32(0) {
											} else {
												v83 = *(*int32)(unsafe.Add(mBase, uint32(v12)+8))
												*(*int32)(unsafe.Add(mBase, uint32(l3))) = v83
											}
											v93 = *(*int32)(unsafe.Add(mBase, uint32(v36)+16))
											if v93&int32(786432) != int32(262144) {
												v100 = *(*int32)(unsafe.Add(mBase, uint32(v36)+8))
												v103 = F_lpDelete(m, v100, v66, v12+int32(12))
												mBase = m.M
												v104 = m.ExcPending
												if v104 != 0 {
													return int32(0)
												} else {
													*(*int32)(unsafe.Add(mBase, uint32(v36)+8)) = v103
													v106 = *(*int32)(unsafe.Add(mBase, uint32(v36)+16))
													v108 = v106 + int32(-1)
													*(*uint16)(unsafe.Add(mBase, uint32(v36)+16)) = uint16(v108)
													if v106&int32(65535) != int32(1) {
														v116 = *(*int32)(unsafe.Add(mBase, uint32(v103)))
														*(*int32)(unsafe.Add(mBase, uint32(v36)+12)) = v116
														v118 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
														*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v118 + int32(-1)
														v130 = int32(1)
														m.G0 = v12 + int32(16)
														return v130
													} else {
														F___quicklistDelNode(m, l0, v36)
														mBase = m.M
														v115 = m.ExcPending
														if v115 != 0 {
															return int32(0)
														} else {
															v118 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
															*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v118 + int32(-1)
															v130 = int32(1)
															m.G0 = v12 + int32(16)
															return v130
														}
													}
												}
											} else {
												F___quicklistDelNode(m, l0, v36)
												mBase = m.M
												v99 = m.ExcPending
												if v99 != 0 {
													return int32(0)
												} else {
													v130 = int32(1)
													m.G0 = v12 + int32(16)
													return v130
												}
											}
										}
									}
								}
							}
						}
					} else {
						if l2 == int32(0) {
							if l3 == int32(0) {
							} else {
								v57 = *(*int32)(unsafe.Add(mBase, uint32(v36)+12))
								*(*int32)(unsafe.Add(mBase, uint32(l3))) = v57
							}
							F_quicklistDelIndex(m, l0, v36, int32(0))
							mBase = m.M
							v61 = m.ExcPending
							if v61 != 0 {
								return int32(0)
							} else {
								v130 = int32(1)
								m.G0 = v12 + int32(16)
								return v130
							}
						} else {
							v48 = *(*int32)(unsafe.Add(mBase, uint32(v36)+8))
							v49 = *(*int32)(unsafe.Add(mBase, uint32(v36)+12))
							v50 = m.T0[l5].(func(*base.Module, int32, int32) int32)(m, v48, v49)
							mBase = m.M
							v53 = m.ExcPending
							if v53 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(l2))) = v50
								if l3 == int32(0) {
								} else {
									v57 = *(*int32)(unsafe.Add(mBase, uint32(v36)+12))
									*(*int32)(unsafe.Add(mBase, uint32(l3))) = v57
								}
								F_quicklistDelIndex(m, l0, v36, int32(0))
								mBase = m.M
								v61 = m.ExcPending
								if v61 != 0 {
									return int32(0)
								} else {
									v130 = int32(1)
									m.G0 = v12 + int32(16)
									return v130
								}
							}
						}
					}
				}
			} else {
				v130 = v7
				m.G0 = v12 + int32(16)
				return v130
			}
		}
	}
}
func F_quicklistPushTail(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v33 int32
	_ = v33
	var v39 int32
	_ = v39
	var v44 int32
	_ = v44
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v68 int32
	_ = v68
	var v73 int32
	_ = v73
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v127 int32
	_ = v127
	var v129 int32
	_ = v129
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v135 int32
	_ = v135
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v142 int32
	_ = v142
	var v144 int32
	_ = v144
	var v146 int32
	_ = v146
	v4 = int32(0)
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v9 = int32(18)
	v12 = v8 << (uint(v9) % 32) >> (uint(v9) % 32)
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v15 = *(*int32)(unsafe.Add(mBase, _consts[381]))
	if v15 == v4 {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	if v13 == int32(0) {
		goto L16
	} else {
		goto L17
	}
L2:
	;
	F___quicklistInsertPlainNode(m, l0, v13, l1, l2, int32(1))
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L13
	} else {
		goto L14
	}
L3:
	;
	if v12 < int32(0) {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	if base.Ui32(v15) <= base.Ui32(l2) {
		goto L2
	} else {
		goto L5
	}
L5:
	;
	goto L1
L6:
	;
	v23 = int32(-5)
	if base.Ui32(v23) < base.Ui32(v12) {
		goto L9
	} else {
		goto L10
	}
L7:
	;
	if base.Ui32(l2) <= base.Ui32(int32(8192)) {
		goto L1
	} else {
		goto L8
	}
L8:
	;
	goto L2
L9:
	;
	v26 = v12
	goto L11
L10:
	;
	v26 = v23
	goto L11
L11:
	;
	v33 = *(*int32)(unsafe.Add(mBase, uint32((v26^int32(-1))<<(uint(int32(2))%32))+uint32(_consts[628])))
	if base.Ui32(l2) <= base.Ui32(v33) {
		goto L1
	} else {
		goto L12
	}
L12:
	;
	goto L2
L13:
	;
	return int32(0)
L14:
	;
	return int32(1)
L15:
	;
	v138 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v139 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v138 + v139
	v142 = *(*int32)(unsafe.Add(mBase, uint32(v135)+16))
	v144 = v142 + v139
	*(*uint16)(unsafe.Add(mBase, uint32(v135)+16)) = uint16(v144)
	v146 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	return base.B2i32(v13 != v146)
L16:
	;
	v119 = F_quicklistCreateNode(m)
	mBase = m.M
	v120 = m.ExcPending
	if v120 != 0 {
		goto L13
	} else {
		goto L44
	}
L17:
	;
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v13)+16))
	if v44&int32(786432) == int32(262144) {
		goto L16
	} else {
		goto L18
	}
L18:
	;
	if v15 == int32(0) {
		goto L23
	} else {
		goto L24
	}
L19:
	;
	v108 = *(*int32)(unsafe.Add(mBase, uint32(v13)+8))
	v109 = F_lpAppend(m, v108, l1, l2)
	mBase = m.M
	v110 = m.ExcPending
	if v110 != 0 {
		goto L13
	} else {
		goto L42
	}
L20:
	;
	if base.Ui32(v103) < base.Ui32(v104) {
		goto L16
	} else {
		goto L41
	}
L21:
	;
	v87 = int32(-5)
	if base.Ui32(v87) < base.Ui32(v12) {
		goto L37
	} else {
		goto L38
	}
L22:
	;
	v80 = int32(1)
	if base.Ui32(v80) < base.Ui32(v12) {
		goto L32
	} else {
		goto L33
	}
L23:
	;
	if v12 < int32(0) {
		goto L21
	} else {
		goto L30
	}
L24:
	;
	if base.Ui32(v15) <= base.Ui32(l2) {
		goto L16
	} else {
		goto L25
	}
L25:
	;
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v13)+12))
	v55 = l2 + v52 + int32(8)
	if int32(-1) < v12 {
		v77 = v55
		goto L22
	} else {
		goto L26
	}
L26:
	;
	v58 = int32(-5)
	if base.Ui32(v58) < base.Ui32(v12) {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	v61 = v12
	goto L29
L28:
	;
	v61 = v58
	goto L29
L29:
	;
	v68 = *(*int32)(unsafe.Add(mBase, uint32((v61^int32(-1))<<(uint(int32(2))%32))+uint32(_consts[628])))
	v103 = v68
	v104 = v55
	goto L20
L30:
	;
	if base.Ui32(int32(8192)) < base.Ui32(l2) {
		goto L16
	} else {
		goto L31
	}
L31:
	;
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v13)+12))
	v77 = l2 + v73 + int32(8)
	goto L22
L32:
	;
	v83 = v12
	goto L34
L33:
	;
	v83 = v80
	goto L34
L34:
	;
	if base.Ui32(v83) <= base.Ui32(v44&int32(65535)) {
		goto L16
	} else {
		goto L35
	}
L35:
	;
	if base.Ui32(int32(8192)) < base.Ui32(v77) {
		goto L16
	} else {
		goto L36
	}
L36:
	;
	goto L19
L37:
	;
	v90 = v12
	goto L39
L38:
	;
	v90 = v87
	goto L39
L39:
	;
	v97 = *(*int32)(unsafe.Add(mBase, uint32((v90^int32(-1))<<(uint(int32(2))%32))+uint32(_consts[628])))
	if base.Ui32(v97) < base.Ui32(l2) {
		goto L16
	} else {
		goto L40
	}
L40:
	;
	v99 = *(*int32)(unsafe.Add(mBase, uint32(v13)+12))
	v103 = v97
	v104 = l2 + v99 + int32(8)
	goto L20
L41:
	;
	goto L19
L42:
	;
	v111 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v111)+8)) = v109
	v113 = *(*int32)(unsafe.Add(mBase, uint32(v109)))
	goto L43
L43:
	;
	v114 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v114)+12)) = v113
	v135 = v114
	goto L15
L44:
	;
	v122 = F_lpNew(m, int32(0))
	mBase = m.M
	v123 = m.ExcPending
	if v123 != 0 {
		goto L13
	} else {
		goto L45
	}
L45:
	;
	v124 = F_lpAppend(m, v122, l1, l2)
	mBase = m.M
	v125 = m.ExcPending
	if v125 != 0 {
		goto L13
	} else {
		goto L46
	}
L46:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v119)+8)) = v124
	v127 = *(*int32)(unsafe.Add(mBase, uint32(v124)))
	goto L47
L47:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v119)+12)) = v127
	v129 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	F___quicklistInsertNode(m, l0, v129, v119, int32(1))
	mBase = m.M
	v132 = m.ExcPending
	if v132 != 0 {
		goto L13
	} else {
		goto L48
	}
L48:
	;
	v133 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v135 = v133
	goto L15
}
func F_quicklistReplaceAtIndex(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
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
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	v6 = m.G0
	v8 = v6 - int32(32)
	m.G0 = v8
	v11 = F_quicklistGetIteratorEntryAtIdx(m, l0, base.I64_extend_i32_s(l1), v8)
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		return int32(0)
	} else {
		if v11 != 0 {
			F_quicklistReplaceEntry(m, v11, v8, l2, l3)
			mBase = m.M
			v17 = m.ExcPending
			if v17 != 0 {
				return int32(0)
			} else {
				v18 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
				if v18 == int32(0) {
					F_valkey_free(m, v11)
					mBase = m.M
					v37 = m.ExcPending
					if v37 != 0 {
						return int32(0)
					} else {
						v40 = int32(1)
						m.G0 = v8 + int32(32)
						return v40
					}
				} else {
					v21 = *(*int32)(unsafe.Add(mBase, uint32(v18)+16))
					if v21&int32(1048576) == int32(0) {
						v32 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
						F___quicklistCompress(m, v32, v18)
						mBase = m.M
						v34 = m.ExcPending
						if v34 != 0 {
							return int32(0)
						} else {
							F_valkey_free(m, v11)
							mBase = m.M
							v37 = m.ExcPending
							if v37 != 0 {
								return int32(0)
							} else {
								v40 = int32(1)
								m.G0 = v8 + int32(32)
								return v40
							}
						}
					} else {
						if v21&int32(196608) != int32(65536) {
							F_valkey_free(m, v11)
							mBase = m.M
							v37 = m.ExcPending
							if v37 != 0 {
								return int32(0)
							} else {
								v40 = int32(1)
								m.G0 = v8 + int32(32)
								return v40
							}
						} else {
							v30 = F___quicklistCompressNode(m, v18)
							mBase = m.M
							v31 = m.ExcPending
							if v31 != 0 {
								return int32(0)
							} else {
								F_valkey_free(m, v11)
								mBase = m.M
								v37 = m.ExcPending
								if v37 != 0 {
									return int32(0)
								} else {
									v40 = int32(1)
									m.G0 = v8 + int32(32)
									return v40
								}
							}
						}
					}
				}
			}
		} else {
			v40 = int32(0)
			m.G0 = v8 + int32(32)
			return v40
		}
	}
}
