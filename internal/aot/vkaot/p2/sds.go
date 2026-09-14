package p2

import (
	base "github.com/shibukawa/valkeymem/internal/aot/vkaot/base"
	"unsafe"
)

func F__sdsMakeRoomFor(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v46 int64
	_ = v46
	var v49 int64
	_ = v49
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v87 int32
	_ = v87
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v103 int32
	_ = v103
	var v108 int32
	_ = v108
	var v111 int32
	_ = v111
	var v113 int32
	_ = v113
	var v117 int32
	_ = v117
	var v120 int32
	_ = v120
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v126 int32
	_ = v126
	var v131 int32
	_ = v131
	var v133 int32
	_ = v133
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v142 int32
	_ = v142
	var v146 int32
	_ = v146
	var v153 int32
	_ = v153
	var v156 int32
	_ = v156
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v169 int32
	_ = v169
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v174 int32
	_ = v174
	var v176 int32
	_ = v176
	var v187 int32
	_ = v187
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v196 int32
	_ = v196
	var v202 int32
	_ = v202
	var v209 int32
	_ = v209
	var v213 int32
	_ = v213
	var v216 int32
	_ = v216
	var v222 int32
	_ = v222
	var v229 int32
	_ = v229
	var v233 int32
	_ = v233
	var v236 int32
	_ = v236
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v248 int32
	_ = v248
	var v250 int32
	_ = v250
	var v252 int32
	_ = v252
	var v253 int32
	_ = v253
	var v255 int32
	_ = v255
	var v257 int32
	_ = v257
	var v261 int32
	_ = v261
	var v262 int32
	_ = v262
	var v263 int32
	_ = v263
	var v268 int32
	_ = v268
	var v269 int32
	_ = v269
	var v270 int32
	_ = v270
	var v272 int32
	_ = v272
	var v274 int32
	_ = v274
	var v279 int32
	_ = v279
	var v290 int32
	_ = v290
	var v295 int32
	_ = v295
	var v304 int32
	_ = v304
	var v305 int32
	_ = v305
	var v307 int32
	_ = v307
	var v312 int32
	_ = v312
	var v317 int32
	_ = v317
	var v319 int32
	_ = v319
	var v320 int32
	_ = v320
	var v321 int32
	_ = v321
	var v322 int32
	_ = v322
	var v323 int32
	_ = v323
	var v325 int32
	_ = v325
	var v328 int32
	_ = v328
	var v329 int32
	_ = v329
	var v331 int32
	_ = v331
	var v336 int32
	_ = v336
	var v337 int32
	_ = v337
	var v344 int32
	_ = v344
	var v349 int32
	_ = v349
	var v351 int32
	_ = v351
	var v354 int32
	_ = v354
	var v355 int32
	_ = v355
	var v357 int32
	_ = v357
	var v363 int32
	_ = v363
	var v364 int32
	_ = v364
	var v371 int32
	_ = v371
	var v380 int32
	_ = v380
	var v381 int32
	_ = v381
	var v382 int32
	_ = v382
	var v383 int32
	_ = v383
	var v384 int32
	_ = v384
	var v386 int32
	_ = v386
	var v387 int32
	_ = v387
	var v389 int32
	_ = v389
	var v396 int32
	_ = v396
	var v400 int32
	_ = v400
	var v403 int32
	_ = v403
	var v408 int32
	_ = v408
	var v409 int32
	_ = v409
	var v411 int32
	_ = v411
	var v416 int32
	_ = v416
	var v419 int32
	_ = v419
	var v423 int32
	_ = v423
	var v425 int32
	_ = v425
	var v446 int32
	_ = v446
	var v460 int32
	_ = v460
	var v466 int32
	_ = v466
	var v472 int32
	_ = v472
	v11 = m.G0
	v13 = v11 - int32(16)
	m.G0 = v13
	v16 = int32(-1)
	v18 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0+v16))))
	v20 = v18 & int32(7)
	switch v20 + v16 {
	case 0:
		goto L5
	case 1:
		goto L4
	case 2:
		goto L3
	case 3:
		goto L2
	default:
		v52 = int32(0)
		goto L1
	}
L1:
	;
	if base.Ui32(v52) < base.Ui32(l1) {
		goto L10
	} else {
		goto L11
	}
L2:
	;
	v46 = *(*int64)(unsafe.Add(mBase, uint32(l0+int32(-9))))
	v49 = *(*int64)(unsafe.Add(mBase, uint32(l0+int32(-17))))
	v52 = base.I32_wrap_i64(v46 - v49)
	goto L1
L3:
	;
	v39 = *(*int32)(unsafe.Add(mBase, uint32(l0+int32(-5))))
	v42 = *(*int32)(unsafe.Add(mBase, uint32(l0+int32(-9))))
	v52 = v39 - v42
	goto L1
L4:
	;
	v32 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0+int32(-3)))))
	v35 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0+int32(-5)))))
	v52 = v32 - v35
	goto L1
L5:
	;
	v25 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0+int32(-2)))))
	v28 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0+int32(-3)))))
	v52 = v25 - v28
	goto L1
L6:
	;
	F__serverAssert(m, int32(_a1324), int32(_a1321), int32(303))
	mBase = m.M
	v472 = m.ExcPending
	if v472 != 0 {
		goto L45
	} else {
		goto L134
	}
L7:
	;
	F__serverAssert(m, int32(_a1325), int32(_a1321), int32(277))
	mBase = m.M
	v466 = m.ExcPending
	if v466 != 0 {
		goto L45
	} else {
		goto L133
	}
L8:
	;
	F__serverAssert(m, int32(_a1326), int32(_a1321), int32(261))
	mBase = m.M
	v460 = m.ExcPending
	if v460 != 0 {
		goto L45
	} else {
		goto L132
	}
L9:
	;
	m.G0 = v13 + int32(16)
	return v446
L10:
	;
	v54 = int32(0)
	switch v20 {
	case 0:
		goto L17
	case 1:
		goto L16
	case 2:
		goto L15
	case 3:
		goto L14
	case 4:
		goto L13
	default:
		v75 = v54
		v76 = v54
		goto L12
	}
L11:
	;
	v446 = l0
	goto L9
L12:
	;
	v77 = v75 + l1
	if base.Ui32(v77) <= base.Ui32(v75) {
		goto L8
	} else {
		goto L18
	}
L13:
	;
	v71 = int32(-17)
	v74 = *(*int32)(unsafe.Add(mBase, uint32(l0+v71)))
	v75 = v74
	v76 = v71
	goto L12
L14:
	;
	v67 = int32(-9)
	v70 = *(*int32)(unsafe.Add(mBase, uint32(l0+v67)))
	v75 = v70
	v76 = v67
	goto L12
L15:
	;
	v63 = int32(-5)
	v66 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0+v63))))
	v75 = v66
	v76 = v63
	goto L12
L16:
	;
	v59 = int32(-3)
	v62 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0+v59))))
	v75 = v62
	v76 = v59
	goto L12
L17:
	;
	v75 = int32(base.Ui32(v18) >> (uint(int32(3)) % 32))
	v76 = int32(-1)
	goto L12
L18:
	;
	if l2 != int32(1) {
		v87 = v77
		goto L19
	} else {
		goto L20
	}
L19:
	;
	if base.Ui32(int32(32)) <= base.Ui32(v87) {
		goto L24
	} else {
		goto L25
	}
L20:
	;
	if base.Ui32(int32(1048575)) < base.Ui32(v77) {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	v87 = v77 + int32(1048576)
	goto L19
L22:
	;
	v87 = v77 << (uint(int32(1)) % 32)
	goto L19
L23:
	;
	v100 = int32(1)
	if base.Ui32(v100) < base.Ui32(v99) {
		goto L31
	} else {
		goto L32
	}
L24:
	;
	if base.Ui32(int32(253)) <= base.Ui32(v87) {
		goto L26
	} else {
		goto L27
	}
L25:
	;
	v99 = int32(0)
	goto L23
L26:
	;
	if base.Ui32(v87) < base.Ui32(int32(65531)) {
		goto L28
	} else {
		goto L29
	}
L27:
	;
	v99 = int32(1)
	goto L23
L28:
	;
	v98 = int32(2)
	goto L30
L29:
	;
	v98 = int32(3)
	goto L30
L30:
	;
	v99 = v98
	goto L23
L31:
	;
	v103 = v99
	goto L33
L32:
	;
	v103 = v100
	goto L33
L33:
	;
	v108 = *(*int32)(unsafe.Add(mBase, uint32(v103<<(uint(int32(2))%32))+uint32(_consts[722])))
	v111 = v87 + v108 + int32(1)
	if base.Ui32(v111) <= base.Ui32(v77) {
		goto L7
	} else {
		goto L34
	}
L34:
	;
	v113 = l0 + v76
	if v20 != v103 {
		goto L42
	} else {
		goto L43
	}
L35:
	;
	v423 = int32(-1)
	v425 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v419+v423))))
	switch v425&int32(7) + v423 {
	case 0:
		goto L131
	case 1:
		goto L130
	case 2:
		goto L129
	case 3:
		goto L128
	default:
		v446 = v419
		goto L9
	}
L36:
	;
	if base.Ui32(v409) < base.Ui32(v408) {
		goto L6
	} else {
		goto L127
	}
L37:
	;
	v408 = v400
	v409 = int32(65535)
	v411 = v403
	goto L36
L38:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v329+int32(-5)))) = uint16(v75)
	v396 = *(*int32)(unsafe.Add(mBase, uint32(v13)+12))
	v400 = v396 + (v322 ^ int32(-1))
	v403 = v323
	goto L37
L39:
	;
	v386 = *(*int32)(unsafe.Add(mBase, uint32(v13)+12))
	v387 = int32(-1)
	v389 = v386 + (v382 ^ v387)
	switch v384 + v387 {
	case 0:
		v408 = v389
		v409 = int32(255)
		v411 = v383
		goto L36
	case 1:
		v400 = v389
		v403 = v383
		goto L37
	default:
		v416 = v389
		v419 = v383
		goto L35
	}
L40:
	;
	v382 = v380
	v383 = v290
	v384 = v381
	goto L39
L41:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v290+int32(-9)))) = v75
	v380 = int32(9)
	v381 = int32(3)
	goto L40
L42:
	;
	v304 = F_zmalloc_usable(m, v111, v13+int32(12))
	mBase = m.M
	v305 = m.ExcPending
	if v305 != 0 {
		goto L45
	} else {
		goto L100
	}
L43:
	;
	v117 = F_zrealloc_usable(m, v113, v111, v13+int32(12))
	mBase = m.M
	v120 = m.ExcPending
	if v120 != 0 {
		goto L45
	} else {
		goto L46
	}
L44:
	;
	v122 = v117 + v108
	v123 = *(*int32)(unsafe.Add(mBase, uint32(v13)+12))
	v124 = int32(-1)
	v126 = v123 + (v108 ^ v124)
	switch v20 + v124 {
	case 0:
		v131 = int32(255)
		goto L48
	case 1:
		goto L49
	default:
		v416 = v126
		v419 = v122
		goto L35
	}
L45:
	;
	return int32(0)
L46:
	;
	if v117 != 0 {
		goto L44
	} else {
		goto L47
	}
L47:
	;
	v446 = int32(0)
	goto L9
L48:
	;
	if base.Ui32(v126) <= base.Ui32(v131) {
		v382 = v108
		v383 = v122
		v384 = v20
		goto L39
	} else {
		goto L50
	}
L49:
	;
	v131 = int32(65535)
	goto L48
L50:
	;
	v133 = int32(5)
	v138 = base.B2i32(base.Ui32(v126) < base.Ui32(int32(65531)))
	if base.Ui32(v126) < base.Ui32(int32(65531)) {
		goto L51
	} else {
		goto L52
	}
L51:
	;
	v139 = v133
	goto L53
L52:
	;
	v139 = int32(9)
	goto L53
L53:
	;
	v140 = v117 + v139
	v142 = v75 + int32(1)
	if v140 == v122 {
		goto L55
	} else {
		goto L56
	}
L54:
	;
	if base.Ui32(v126) < base.Ui32(int32(65531)) {
		goto L95
	} else {
		goto L96
	}
L55:
	;
	v290 = v140
	goto L54
L56:
	;
	v146 = v142 + v140
	if base.Ui32(int32(0)-v142<<(uint(int32(1))%32)) < base.Ui32(v122-v146) {
		goto L57
	} else {
		goto L58
	}
L57:
	;
	v156 = (v122 ^ v140) & int32(3)
	if base.Ui32(v122) <= base.Ui32(v140) {
		goto L61
	} else {
		goto L62
	}
L58:
	;
	v153 = F___memcpy(m, v140, v122, v142)
	mBase = m.M
	v290 = v153
	goto L54
L59:
	;
	if v262 == int32(0) {
		goto L55
	} else {
		goto L91
	}
L60:
	;
	if base.Ui32(v240) <= base.Ui32(int32(3)) {
		v261 = v239
		v262 = v240
		v263 = v241
		goto L59
	} else {
		goto L87
	}
L61:
	;
	if v156 != 0 {
		v222 = v142
		goto L71
	} else {
		goto L72
	}
L62:
	;
	if v156 == int32(0) {
		goto L63
	} else {
		goto L64
	}
L63:
	;
	if v140&int32(3) != 0 {
		goto L65
	} else {
		goto L66
	}
L64:
	;
	v261 = v122
	v262 = v142
	v263 = v140
	goto L59
L65:
	;
	v163 = v122
	v164 = v142
	v165 = v140
	goto L67
L66:
	;
	v239 = v122
	v240 = v142
	v241 = v140
	goto L60
L67:
	;
	if v164 == int32(0) {
		goto L55
	} else {
		goto L69
	}
L69:
	;
	v169 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v163))))
	*(*uint8)(unsafe.Add(mBase, uint32(v165))) = uint8(v169)
	v171 = int32(1)
	v172 = v163 + v171
	v174 = v164 + int32(-1)
	v176 = v165 + v171
	if v176&int32(3) == int32(0) {
		v239 = v172
		v240 = v174
		v241 = v176
		goto L60
	} else {
		goto L70
	}
L70:
	;
	v163 = v172
	v164 = v174
	v165 = v176
	goto L67
L71:
	;
	if v222 == int32(0) {
		goto L55
	} else {
		goto L83
	}
L72:
	;
	if v146&int32(3) == int32(0) {
		v202 = v142
		goto L73
	} else {
		goto L74
	}
L73:
	;
	if base.Ui32(v202) <= base.Ui32(int32(3)) {
		v222 = v202
		goto L71
	} else {
		goto L79
	}
L74:
	;
	v187 = v142
	goto L75
L75:
	;
	if v187 == int32(0) {
		goto L55
	} else {
		goto L77
	}
L76:
	;
	v202 = v193
	goto L73
L77:
	;
	v193 = v187 + int32(-1)
	v194 = v140 + v193
	v196 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v122+v193))))
	*(*uint8)(unsafe.Add(mBase, uint32(v194))) = uint8(v196)
	if v194&int32(3) != 0 {
		v187 = v193
		goto L75
	} else {
		goto L78
	}
L78:
	;
	goto L76
L79:
	;
	v209 = v202
	goto L80
L80:
	;
	v213 = v209 + int32(-4)
	v216 = *(*int32)(unsafe.Add(mBase, uint32(v122+v213)))
	*(*int32)(unsafe.Add(mBase, uint32(v140+v213))) = v216
	if base.Ui32(int32(3)) < base.Ui32(v213) {
		v209 = v213
		goto L80
	} else {
		goto L82
	}
L81:
	;
	v222 = v213
	goto L71
L82:
	;
	goto L81
L83:
	;
	v229 = v222
	goto L84
L84:
	;
	v233 = v229 + int32(-1)
	v236 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v122+v233))))
	*(*uint8)(unsafe.Add(mBase, uint32(v140+v233))) = uint8(v236)
	if v233 != 0 {
		v229 = v233
		goto L84
	} else {
		goto L86
	}
L86:
	;
	goto L55
L87:
	;
	v246 = v239
	v247 = v240
	v248 = v241
	goto L88
L88:
	;
	v250 = *(*int32)(unsafe.Add(mBase, uint32(v246)))
	*(*int32)(unsafe.Add(mBase, uint32(v248))) = v250
	v252 = int32(4)
	v253 = v246 + v252
	v255 = v248 + v252
	v257 = v247 + int32(-4)
	if base.Ui32(int32(3)) < base.Ui32(v257) {
		v246 = v253
		v247 = v257
		v248 = v255
		goto L88
	} else {
		goto L90
	}
L89:
	;
	v261 = v253
	v262 = v257
	v263 = v255
	goto L59
L90:
	;
	goto L89
L91:
	;
	v268 = v261
	v269 = v262
	v270 = v263
	goto L92
L92:
	;
	v272 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v268))))
	*(*uint8)(unsafe.Add(mBase, uint32(v270))) = uint8(v272)
	v274 = int32(1)
	v279 = v269 + int32(-1)
	if v279 != 0 {
		v268 = v268 + v274
		v269 = v279
		v270 = v270 + v274
		goto L92
	} else {
		goto L94
	}
L93:
	;
	goto L55
L94:
	;
	goto L93
L95:
	;
	v295 = int32(2)
	goto L97
L96:
	;
	v295 = int32(3)
	goto L97
L97:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v290+int32(-1)))) = uint8(v295)
	if base.Ui32(int32(65530)) < base.Ui32(v126) {
		goto L41
	} else {
		goto L98
	}
L98:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v290+int32(-5)))) = uint16(v75)
	v380 = v133
	v381 = int32(2)
	goto L40
L99:
	;
	v307 = *(*int32)(unsafe.Add(mBase, uint32(v13)+12))
	switch v103 + int32(-1) {
	case 0:
		v312 = int32(255)
		goto L104
	case 1:
		goto L105
	default:
		goto L103
	}
L100:
	;
	if v304 != 0 {
		goto L99
	} else {
		goto L101
	}
L101:
	;
	v446 = int32(0)
	goto L9
L102:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v364+int32(-9)))) = v75
	v371 = *(*int32)(unsafe.Add(mBase, uint32(v13)+12))
	v416 = v371 + (v363 ^ int32(-1))
	v419 = v364
	goto L35
L103:
	;
	v349 = v304 + v108
	v351 = v75 + int32(1)
	if v351 == int32(0) {
		v355 = v349
		goto L124
	} else {
		goto L125
	}
L104:
	;
	v317 = v307 + (v108 ^ int32(-1))
	v319 = base.B2i32(base.Ui32(v317) < base.Ui32(int32(65531)))
	if base.Ui32(v317) < base.Ui32(int32(65531)) {
		goto L106
	} else {
		goto L107
	}
L105:
	;
	v312 = int32(65535)
	goto L104
L106:
	;
	v320 = int32(5)
	goto L108
L107:
	;
	v320 = int32(9)
	goto L108
L108:
	;
	v321 = base.B2i32(base.Ui32(v312) < base.Ui32(v317))
	if base.Ui32(v312) < base.Ui32(v317) {
		goto L109
	} else {
		goto L110
	}
L109:
	;
	v322 = v320
	goto L111
L110:
	;
	v322 = v108
	goto L111
L111:
	;
	v323 = v304 + v322
	v325 = v75 + int32(1)
	if v325 == int32(0) {
		v329 = v323
		goto L113
	} else {
		goto L114
	}
L112:
	;
	F_valkey_free(m, v113)
	mBase = m.M
	v331 = m.ExcPending
	if v331 != 0 {
		goto L45
	} else {
		goto L115
	}
L113:
	;
	goto L112
L114:
	;
	v328 = F__emscripten_memcpy_bulkmem(m, v323, l0, v325)
	mBase = m.M
	v329 = v328
	goto L113
L115:
	;
	if base.Ui32(v317) < base.Ui32(int32(65531)) {
		goto L116
	} else {
		goto L117
	}
L116:
	;
	v336 = int32(2)
	goto L118
L117:
	;
	v336 = int32(3)
	goto L118
L118:
	;
	if base.Ui32(v312) < base.Ui32(v317) {
		goto L119
	} else {
		goto L120
	}
L119:
	;
	v337 = v336
	goto L121
L120:
	;
	v337 = v103
	goto L121
L121:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v329+int32(-1)))) = uint8(v337)
	switch v337 + int32(-1) {
	default:
		goto L122
	case 1:
		goto L38
	case 2:
		v363 = v322
		v364 = v323
		goto L102
	}
L122:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v329+int32(-3)))) = uint8(v75)
	v344 = *(*int32)(unsafe.Add(mBase, uint32(v13)+12))
	v408 = v344 + (v322 ^ int32(-1))
	v409 = int32(255)
	v411 = v323
	goto L36
L123:
	;
	F_valkey_free(m, v113)
	mBase = m.M
	v357 = m.ExcPending
	if v357 != 0 {
		goto L45
	} else {
		goto L126
	}
L124:
	;
	goto L123
L125:
	;
	v354 = F__emscripten_memcpy_bulkmem(m, v349, l0, v351)
	mBase = m.M
	v355 = v354
	goto L124
L126:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v355+int32(-1)))) = uint8(v103)
	v363 = v108
	v364 = v349
	goto L102
L127:
	;
	v416 = v408
	v419 = v411
	goto L35
L128:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v419+int32(-9)))) = base.I64_extend_i32_u(v416)
	v446 = v419
	goto L9
L129:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v419+int32(-5)))) = v416
	v446 = v419
	goto L9
L130:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v419+int32(-3)))) = uint16(v416)
	v446 = v419
	goto L9
L131:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v419+int32(-2)))) = uint8(v416)
	v446 = v419
	goto L9
L132:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L133:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L134:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_isSdsRepresentableAsLongLong(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v15 int32
	_ = v15
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
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v76 int64
	_ = v76
	var v82 int32
	_ = v82
	var v84 int64
	_ = v84
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v98 int64
	_ = v98
	var v103 int64
	_ = v103
	var v107 int32
	_ = v107
	var v109 int64
	_ = v109
	var v111 int32
	_ = v111
	var v119 int64
	_ = v119
	var v143 int64
	_ = v143
	var v162 int32
	_ = v162
	var v169 int32
	_ = v169
	v8 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0+int32(-1)))))
	switch v8 & int32(7) {
	case 0:
		goto L6
	case 1:
		goto L5
	case 2:
		goto L4
	case 3:
		goto L3
	case 4:
		goto L2
	default:
		v25 = int32(0)
		goto L1
	}
L1:
	;
	v26 = int32(0)
	if base.Ui32(v25+int32(-21)) < base.Ui32(int32(-20)) {
		v162 = v26
		goto L8
	} else {
		goto L9
	}
L2:
	;
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l0+int32(-17))))
	v25 = v24
	goto L1
L3:
	;
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l0+int32(-9))))
	v25 = v21
	goto L1
L4:
	;
	v18 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0+int32(-5)))))
	v25 = v18
	goto L1
L5:
	;
	v15 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0+int32(-3)))))
	v25 = v15
	goto L1
L6:
	;
	v25 = int32(base.Ui32(v8) >> (uint(int32(3)) % 32))
	goto L1
L7:
	;
	if v162 != 0 {
		goto L34
	} else {
		goto L35
	}
L8:
	;
	goto L7
L9:
	;
	v40 = int32(1)
	v41 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v25 != v40 {
		goto L12
	} else {
		goto L13
	}
L10:
	;
	v162 = int32(1)
	goto L8
L11:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l1))) = v143
	goto L10
L12:
	;
	if v41&int32(255) == int32(45) {
		goto L17
	} else {
		goto L18
	}
L13:
	;
	v45 = v41 + int32(-48)
	if base.Ui32(int32(9)) < base.Ui32(v45&int32(255)) {
		v162 = v26
		goto L8
	} else {
		goto L14
	}
L14:
	;
	if l1 == int32(0) {
		goto L10
	} else {
		goto L15
	}
L15:
	;
	v143 = base.I64_extend_i32_u(v45) & int64(255)
	goto L11
L16:
	;
	if base.Ui32(int32(8)) < base.Ui32((v64+int32(-49))&int32(255)) {
		v162 = v26
		goto L8
	} else {
		goto L19
	}
L17:
	;
	v61 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+1)))
	v63 = int32(2)
	v64 = v61
	v65 = l0 + int32(1)
	goto L16
L18:
	;
	v63 = v40
	v64 = v41
	v65 = l0
	goto L16
L19:
	;
	v76 = base.I64_extend_i32_u(v64+int32(-48)) & int64(255)
	if base.Ui32(v25) <= base.Ui32(v63) {
		v119 = v76
		goto L20
	} else {
		goto L21
	}
L20:
	;
	if v41&int32(255) != int32(45) {
		goto L28
	} else {
		goto L29
	}
L21:
	;
	v82 = v63
	v84 = v76
	v86 = v65
	goto L22
L22:
	;
	v88 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v86)+1)))
	if base.Ui32((v88+int32(-58))&int32(255)) < base.Ui32(int32(246)) {
		v162 = v26
		goto L8
	} else {
		goto L24
	}
L23:
	;
	v119 = v109
	goto L20
L24:
	;
	if base.Ui64(int64(1844674407370955161)) < base.Ui64(v84) {
		v162 = v26
		goto L8
	} else {
		goto L25
	}
L25:
	;
	v98 = v84 * int64(10)
	v103 = base.I64_extend_i32_u(v88+int32(-48)) & int64(255)
	if base.Ui64(v103^int64(-1)) < base.Ui64(v98) {
		v162 = v26
		goto L8
	} else {
		goto L26
	}
L26:
	;
	v107 = int32(1)
	v109 = v98 + v103
	v111 = v82 + v107
	if v111 != v25 {
		v82 = v111
		v84 = v109
		v86 = v86 + v107
		goto L22
	} else {
		goto L27
	}
L27:
	;
	goto L23
L28:
	;
	if v119 < int64(0) {
		v162 = v26
		goto L8
	} else {
		goto L32
	}
L29:
	;
	if base.Ui64(int64(-9223372036854775807-1)) < base.Ui64(v119) {
		v162 = v26
		goto L8
	} else {
		goto L30
	}
L30:
	;
	if l1 == int32(0) {
		goto L10
	} else {
		goto L31
	}
L31:
	;
	v143 = int64(0) - v119
	goto L11
L32:
	;
	if l1 == int32(0) {
		goto L10
	} else {
		goto L33
	}
L33:
	;
	v143 = v119
	goto L11
L34:
	;
	v169 = v26
	goto L36
L35:
	;
	v169 = int32(-1)
	goto L36
L36:
	;
	return v169
}
func F_sdsAllocSize(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v22 int32
	_ = v22
	var v30 int32
	_ = v30
	var v35 int32
	_ = v35
	var v40 int32
	_ = v40
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	v9 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0+int32(-1)))))
	v11 = v9 & int32(7)
	switch v11 {
	case 0:
		v14 = *(*int32)(unsafe.Add(mBase, uint32(l0+int32(-9))))
		return v14 & int32(2147483647)
	case 1:
		v22 = int32(4)
		switch v11 {
		case 0:
			return v22 + int32(base.Ui32(v9)>>(uint(int32(3))%32))
		case 1:
			v30 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0+int32(-2)))))
			return v22 + v30
		case 2:
			v35 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0+int32(-3)))))
			return v22 + v35
		case 3:
			v40 = *(*int32)(unsafe.Add(mBase, uint32(l0+int32(-5))))
			return v22 + v40
		case 4:
			v45 = *(*int32)(unsafe.Add(mBase, uint32(l0+int32(-9))))
			v46 = v45
			return v22 + v46
		default:
			v46 = int32(0)
			return v22 + v46
		}
	case 2:
		v22 = int32(6)
		switch v11 {
		case 0:
			return v22 + int32(base.Ui32(v9)>>(uint(int32(3))%32))
		case 1:
			v30 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0+int32(-2)))))
			return v22 + v30
		case 2:
			v35 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0+int32(-3)))))
			return v22 + v35
		case 3:
			v40 = *(*int32)(unsafe.Add(mBase, uint32(l0+int32(-5))))
			return v22 + v40
		case 4:
			v45 = *(*int32)(unsafe.Add(mBase, uint32(l0+int32(-9))))
			v46 = v45
			return v22 + v46
		default:
			v46 = int32(0)
			return v22 + v46
		}
	case 3:
		v22 = int32(10)
		switch v11 {
		case 0:
			return v22 + int32(base.Ui32(v9)>>(uint(int32(3))%32))
		case 1:
			v30 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0+int32(-2)))))
			return v22 + v30
		case 2:
			v35 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0+int32(-3)))))
			return v22 + v35
		case 3:
			v40 = *(*int32)(unsafe.Add(mBase, uint32(l0+int32(-5))))
			return v22 + v40
		case 4:
			v45 = *(*int32)(unsafe.Add(mBase, uint32(l0+int32(-9))))
			v46 = v45
			return v22 + v46
		default:
			v46 = int32(0)
			return v22 + v46
		}
	case 4:
		v22 = int32(18)
		switch v11 {
		case 0:
			return v22 + int32(base.Ui32(v9)>>(uint(int32(3))%32))
		case 1:
			v30 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0+int32(-2)))))
			return v22 + v30
		case 2:
			v35 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0+int32(-3)))))
			return v22 + v35
		case 3:
			v40 = *(*int32)(unsafe.Add(mBase, uint32(l0+int32(-5))))
			return v22 + v40
		case 4:
			v45 = *(*int32)(unsafe.Add(mBase, uint32(l0+int32(-9))))
			v46 = v45
			return v22 + v46
		default:
			v46 = int32(0)
			return v22 + v46
		}
	default:
		v22 = int32(1)
		switch v11 {
		case 0:
			return v22 + int32(base.Ui32(v9)>>(uint(int32(3))%32))
		case 1:
			v30 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0+int32(-2)))))
			return v22 + v30
		case 2:
			v35 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0+int32(-3)))))
			return v22 + v35
		case 3:
			v40 = *(*int32)(unsafe.Add(mBase, uint32(l0+int32(-5))))
			return v22 + v40
		case 4:
			v45 = *(*int32)(unsafe.Add(mBase, uint32(l0+int32(-9))))
			v46 = v45
			return v22 + v46
		default:
			v46 = int32(0)
			return v22 + v46
		}
	}
}
func F_sdsConfigInit(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	if v5 == int32(0) {
		v12 = F_sdsnew(m, v4)
		mBase = m.M
		v13 = m.ExcPending
		if v13 != 0 {
			return
		} else {
			v14 = v12
			v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
			*(*int32)(unsafe.Add(mBase, uint32(v15))) = v14
			return
		}
	} else {
		v8 = int32(0)
		if v4 == v8 {
			v14 = v8
			v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
			*(*int32)(unsafe.Add(mBase, uint32(v15))) = v14
			return
		} else {
			v12 = F_sdsnew(m, v4)
			mBase = m.M
			v13 = m.ExcPending
			if v13 != 0 {
				return
			} else {
				v14 = v12
				v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
				*(*int32)(unsafe.Add(mBase, uint32(v15))) = v14
				return
			}
		}
	}
}
func F_sdsIncrLen(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v50 int32
	_ = v50
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v69 int32
	_ = v69
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v86 int32
	_ = v86
	var v91 int64
	_ = v91
	var v92 int64
	_ = v92
	var v96 int64
	_ = v96
	var v101 int64
	_ = v101
	var v103 int64
	_ = v103
	var v107 int32
	_ = v107
	var v111 int32
	_ = v111
	var v118 int32
	_ = v118
	var v125 int32
	_ = v125
	var v132 int32
	_ = v132
	var v139 int32
	_ = v139
	var v146 int32
	_ = v146
	v9 = l0 + int32(-1)
	v10 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9))))
	switch v10 & int32(7) {
	case 0:
		v14 = int32(base.Ui32(v10) >> (uint(int32(3)) % 32))
		if l1 < int32(1) {
			if l1 == int32(0) {
				F__serverAssert(m, int32(_a1327), int32(_a1321), int32(460))
				mBase = m.M
				v146 = m.ExcPending
				if v146 != 0 {
					return
				} else {
					F_abort(m)
					mBase = m.M
					base.Wasm_trap_unreachable()
					for {
					}
				}
			} else {
				if base.Ui32(v14) < base.Ui32(int32(0)-l1) {
					F__serverAssert(m, int32(_a1327), int32(_a1321), int32(460))
					mBase = m.M
					v146 = m.ExcPending
					if v146 != 0 {
						return
					} else {
						F_abort(m)
						mBase = m.M
						base.Wasm_trap_unreachable()
						for {
						}
					}
				} else {
					v26 = l1 + v14
					v28 = v26 << (uint(int32(3)) % 32)
					*(*uint8)(unsafe.Add(mBase, uint32(v9))) = uint8(v28)
					v107 = v26
					v111 = int32(0)
					*(*uint8)(unsafe.Add(mBase, uint32(l0+v107))) = uint8(v111)
					return
				}
			}
		} else {
			v17 = l1 + v14
			if base.Ui32(v17) < base.Ui32(int32(32)) {
				v26 = v17
				v28 = v26 << (uint(int32(3)) % 32)
				*(*uint8)(unsafe.Add(mBase, uint32(v9))) = uint8(v28)
				v107 = v26
				v111 = int32(0)
				*(*uint8)(unsafe.Add(mBase, uint32(l0+v107))) = uint8(v111)
				return
			} else {
				F__serverAssert(m, int32(_a1327), int32(_a1321), int32(460))
				mBase = m.M
				v146 = m.ExcPending
				if v146 != 0 {
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
	case 1:
		v31 = l0 + int32(-3)
		if l1 < int32(0) {
			v40 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v31))))
			if base.Ui32(v40) < base.Ui32(int32(0)-l1) {
				F__serverAssert(m, int32(_a1328), int32(_a1321), int32(467))
				mBase = m.M
				v139 = m.ExcPending
				if v139 != 0 {
					return
				} else {
					F_abort(m)
					mBase = m.M
					base.Wasm_trap_unreachable()
					for {
					}
				}
			} else {
				v44 = v40
				v45 = v44 + l1
				*(*uint8)(unsafe.Add(mBase, uint32(v31))) = uint8(v45)
				v107 = v45 & int32(255)
				v111 = int32(0)
				*(*uint8)(unsafe.Add(mBase, uint32(l0+v107))) = uint8(v111)
				return
			}
		} else {
			v36 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0+int32(-2)))))
			v37 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v31))))
			if l1 <= v36-v37 {
				v44 = v37
				v45 = v44 + l1
				*(*uint8)(unsafe.Add(mBase, uint32(v31))) = uint8(v45)
				v107 = v45 & int32(255)
				v111 = int32(0)
				*(*uint8)(unsafe.Add(mBase, uint32(l0+v107))) = uint8(v111)
				return
			} else {
				F__serverAssert(m, int32(_a1328), int32(_a1321), int32(467))
				mBase = m.M
				v139 = m.ExcPending
				if v139 != 0 {
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
	case 2:
		v50 = l0 + int32(-5)
		if l1 < int32(0) {
			v59 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v50))))
			if base.Ui32(v59) < base.Ui32(int32(0)-l1) {
				F__serverAssert(m, int32(_a1328), int32(_a1321), int32(473))
				mBase = m.M
				v132 = m.ExcPending
				if v132 != 0 {
					return
				} else {
					F_abort(m)
					mBase = m.M
					base.Wasm_trap_unreachable()
					for {
					}
				}
			} else {
				v63 = v59
				v64 = v63 + l1
				*(*uint16)(unsafe.Add(mBase, uint32(v50))) = uint16(v64)
				v107 = v64 & int32(65535)
				v111 = int32(0)
				*(*uint8)(unsafe.Add(mBase, uint32(l0+v107))) = uint8(v111)
				return
			}
		} else {
			v55 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0+int32(-3)))))
			v56 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v50))))
			if l1 <= v55-v56 {
				v63 = v56
				v64 = v63 + l1
				*(*uint16)(unsafe.Add(mBase, uint32(v50))) = uint16(v64)
				v107 = v64 & int32(65535)
				v111 = int32(0)
				*(*uint8)(unsafe.Add(mBase, uint32(l0+v107))) = uint8(v111)
				return
			} else {
				F__serverAssert(m, int32(_a1328), int32(_a1321), int32(473))
				mBase = m.M
				v132 = m.ExcPending
				if v132 != 0 {
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
	case 3:
		v69 = l0 + int32(-9)
		if l1 < int32(0) {
			v78 = *(*int32)(unsafe.Add(mBase, uint32(v69)))
			if base.Ui32(v78) < base.Ui32(int32(0)-l1) {
				F__serverAssert(m, int32(_a1329), int32(_a1321), int32(480))
				mBase = m.M
				v125 = m.ExcPending
				if v125 != 0 {
					return
				} else {
					F_abort(m)
					mBase = m.M
					base.Wasm_trap_unreachable()
					for {
					}
				}
			} else {
				v82 = v78
				v83 = v82 + l1
				*(*int32)(unsafe.Add(mBase, uint32(v69))) = v83
				v107 = v83
				v111 = int32(0)
				*(*uint8)(unsafe.Add(mBase, uint32(l0+v107))) = uint8(v111)
				return
			}
		} else {
			v74 = *(*int32)(unsafe.Add(mBase, uint32(l0+int32(-5))))
			v75 = *(*int32)(unsafe.Add(mBase, uint32(v69)))
			if base.Ui32(l1) <= base.Ui32(v74-v75) {
				v82 = v75
				v83 = v82 + l1
				*(*int32)(unsafe.Add(mBase, uint32(v69))) = v83
				v107 = v83
				v111 = int32(0)
				*(*uint8)(unsafe.Add(mBase, uint32(l0+v107))) = uint8(v111)
				return
			} else {
				F__serverAssert(m, int32(_a1329), int32(_a1321), int32(480))
				mBase = m.M
				v125 = m.ExcPending
				if v125 != 0 {
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
	case 4:
		v86 = l0 + int32(-17)
		if l1 < int32(0) {
			v96 = *(*int64)(unsafe.Add(mBase, uint32(v86)))
			if base.Ui64(v96) < base.Ui64(base.I64_extend_i32_u(int32(0)-l1)) {
				F__serverAssert(m, int32(_a1330), int32(_a1321), int32(486))
				mBase = m.M
				v118 = m.ExcPending
				if v118 != 0 {
					return
				} else {
					F_abort(m)
					mBase = m.M
					base.Wasm_trap_unreachable()
					for {
					}
				}
			} else {
				v101 = v96
				v103 = v101 + base.I64_extend_i32_s(l1)
				*(*int64)(unsafe.Add(mBase, uint32(v86))) = v103
				v107 = base.I32_wrap_i64(v103)
				v111 = int32(0)
				*(*uint8)(unsafe.Add(mBase, uint32(l0+v107))) = uint8(v111)
				return
			}
		} else {
			v91 = *(*int64)(unsafe.Add(mBase, uint32(l0+int32(-9))))
			v92 = *(*int64)(unsafe.Add(mBase, uint32(v86)))
			if base.Ui64(base.I64_extend_i32_u(l1)) <= base.Ui64(v91-v92) {
				v101 = v92
				v103 = v101 + base.I64_extend_i32_s(l1)
				*(*int64)(unsafe.Add(mBase, uint32(v86))) = v103
				v107 = base.I32_wrap_i64(v103)
				v111 = int32(0)
				*(*uint8)(unsafe.Add(mBase, uint32(l0+v107))) = uint8(v111)
				return
			} else {
				F__serverAssert(m, int32(_a1330), int32(_a1321), int32(486))
				mBase = m.M
				v118 = m.ExcPending
				if v118 != 0 {
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
	default:
		v107 = int32(0)
		v111 = int32(0)
		*(*uint8)(unsafe.Add(mBase, uint32(l0+v107))) = uint8(v111)
		return
	}
}
func F_sdsRemoveFreeSpace(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
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
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	v8 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0+int32(-1)))))
	switch v8 & int32(7) {
	case 0:
		v13 = F_sdsResize(m, l0, int32(base.Ui32(v8)>>(uint(int32(3))%32)), l1)
		mBase = m.M
		v16 = m.ExcPending
		if v16 != 0 {
			return int32(0)
		} else {
			return v13
		}
	case 1:
		v20 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0+int32(-3)))))
		v21 = F_sdsResize(m, l0, v20, l1)
		mBase = m.M
		v22 = m.ExcPending
		if v22 != 0 {
			return int32(0)
		} else {
			return v21
		}
	case 2:
		v26 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0+int32(-5)))))
		v27 = F_sdsResize(m, l0, v26, l1)
		mBase = m.M
		v28 = m.ExcPending
		if v28 != 0 {
			return int32(0)
		} else {
			return v27
		}
	case 3:
		v32 = *(*int32)(unsafe.Add(mBase, uint32(l0+int32(-9))))
		v33 = F_sdsResize(m, l0, v32, l1)
		mBase = m.M
		v34 = m.ExcPending
		if v34 != 0 {
			return int32(0)
		} else {
			return v33
		}
	case 4:
		v38 = *(*int32)(unsafe.Add(mBase, uint32(l0+int32(-17))))
		v39 = v38
		v40 = F_sdsResize(m, l0, v39, l1)
		mBase = m.M
		v41 = m.ExcPending
		if v41 != 0 {
			return int32(0)
		} else {
			return v40
		}
	default:
		v39 = int32(0)
		v40 = F_sdsResize(m, l0, v39, l1)
		mBase = m.M
		v41 = m.ExcPending
		if v41 != 0 {
			return int32(0)
		} else {
			return v40
		}
	}
}
func F_sdsReqType(m *base.Module, l0 int32) int32 {
	var v14 int32
	_ = v14
	if base.Ui32(int32(32)) <= base.Ui32(l0) {
		if base.Ui32(int32(253)) <= base.Ui32(l0) {
			if base.Ui32(l0) < base.Ui32(int32(65531)) {
				v14 = int32(2)
			} else {
				v14 = int32(3)
			}
			return v14
		} else {
			return int32(1)
		}
	} else {
		return int32(0)
	}
}
func F_updateSdsExtensionField(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
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
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v74 int32
	_ = v74
	var v78 int32
	_ = v78
	var v88 int32
	_ = v88
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if l1 == int32(0) {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	return
L2:
	;
	F_clearCachedClusterSlotsResponse(m)
	mBase = m.M
	v92 = m.ExcPending
	if v92 != 0 {
		goto L14
	} else {
		goto L34
	}
L3:
	;
	v45 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4+int32(-1)))))
	switch v45 & int32(7) {
	case 0:
		goto L17
	case 1:
		goto L21
	case 2:
		goto L20
	case 3:
		goto L19
	case 4:
		goto L18
	default:
		goto L1
	}
L4:
	;
	v9 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4))))
	v10 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	if v10 == int32(0) {
		v33 = v9
		v34 = v10
		goto L6
	} else {
		goto L7
	}
L5:
	;
	if v34-v33&int32(255) == int32(0) {
		goto L1
	} else {
		goto L13
	}
L6:
	;
	goto L5
L7:
	;
	if v10 != v9&int32(255) {
		v33 = v9
		v34 = v10
		goto L6
	} else {
		goto L8
	}
L8:
	;
	v16 = l1
	v17 = v4
	goto L9
L9:
	;
	v20 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17)+1)))
	v21 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16)+1)))
	if v21 == int32(0) {
		v33 = v20
		v34 = v21
		goto L6
	} else {
		goto L11
	}
L10:
	;
	v33 = v20
	v34 = v21
	goto L6
L11:
	;
	v24 = int32(1)
	if v21 == v20&int32(255) {
		v16 = v16 + v24
		v17 = v17 + v24
		goto L9
	} else {
		goto L12
	}
L12:
	;
	goto L10
L13:
	;
	v40 = F_sdscpy(m, v4, l1)
	mBase = m.M
	v41 = m.ExcPending
	if v41 != 0 {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	return
L15:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v40
	goto L2
L16:
	;
	v66 = v4 + int32(-1)
	v67 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v66))))
	switch v67 & int32(7) {
	case 0:
		goto L33
	case 1:
		goto L32
	case 2:
		goto L31
	case 3:
		goto L30
	case 4:
		goto L29
	default:
		goto L28
	}
L17:
	;
	if int32(base.Ui32(v45)>>(uint(int32(3))%32)) == int32(0) {
		goto L1
	} else {
		goto L26
	}
L18:
	;
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v4+int32(-17))))
	if v59 != 0 {
		goto L16
	} else {
		goto L25
	}
L19:
	;
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v4+int32(-9))))
	if v56 != 0 {
		goto L16
	} else {
		goto L24
	}
L20:
	;
	v53 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v4+int32(-5)))))
	if v53 != 0 {
		goto L16
	} else {
		goto L23
	}
L21:
	;
	v50 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4+int32(-3)))))
	if v50 != 0 {
		goto L16
	} else {
		goto L22
	}
L22:
	;
	goto L1
L23:
	;
	goto L1
L24:
	;
	goto L1
L25:
	;
	goto L1
L26:
	;
	goto L16
L27:
	;
	goto L2
L28:
	;
	v88 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v4))) = uint8(v88)
	goto L27
L29:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v4+int32(-17)))) = int64(0)
	goto L28
L30:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v4+int32(-9)))) = int32(0)
	goto L28
L31:
	;
	v78 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v4+int32(-5)))) = uint16(v78)
	goto L28
L32:
	;
	v74 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v4+int32(-3)))) = uint8(v74)
	goto L28
L33:
	;
	v70 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v66))) = uint8(v70)
	goto L28
L34:
	;
	v94 = *(*int32)(unsafe.Add(mBase, _consts[136]))
	v95 = *(*int32)(unsafe.Add(mBase, uint32(v94)+uint32(_consts[150])))
	*(*int32)(unsafe.Add(mBase, uint32(v94)+uint32(_consts[150]))) = v95 | int32(4)
	goto L1
}
