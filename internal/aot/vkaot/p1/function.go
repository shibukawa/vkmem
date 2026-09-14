package p1

import (
	base "github.com/shibukawa/vkmem/internal/aot/vkaot/base"
	"unsafe"
)

func F_functionExtractLibMetaData(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
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
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v91 int32
	_ = v91
	var v95 int32
	_ = v95
	var v98 int32
	_ = v98
	var v103 int32
	_ = v103
	var v106 int32
	_ = v106
	var v109 int32
	_ = v109
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v120 int32
	_ = v120
	var v122 int32
	_ = v122
	var v124 int32
	_ = v124
	var v128 int32
	_ = v128
	var v130 int32
	_ = v130
	var v133 int32
	_ = v133
	var v160 int32
	_ = v160
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v185 int32
	_ = v185
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v192 int32
	_ = v192
	var v196 int32
	_ = v196
	var v200 int32
	_ = v200
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v214 int32
	_ = v214
	var v216 int32
	_ = v216
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v246 int32
	_ = v246
	var v254 int32
	_ = v254
	var v255 int32
	_ = v255
	var v257 int32
	_ = v257
	var v262 int32
	_ = v262
	var v265 int32
	_ = v265
	var v268 int32
	_ = v268
	var v271 int32
	_ = v271
	var v272 int32
	_ = v272
	var v278 int32
	_ = v278
	var v282 int32
	_ = v282
	var v285 int32
	_ = v285
	var v290 int32
	_ = v290
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
	var v303 int32
	_ = v303
	var v304 int32
	_ = v304
	var v305 int32
	_ = v305
	var v307 int32
	_ = v307
	var v309 int32
	_ = v309
	var v311 int32
	_ = v311
	var v315 int32
	_ = v315
	var v317 int32
	_ = v317
	var v320 int32
	_ = v320
	var v348 int32
	_ = v348
	var v349 int32
	_ = v349
	var v363 int32
	_ = v363
	var v364 int32
	_ = v364
	var v367 int32
	_ = v367
	var v369 int32
	_ = v369
	var v373 int32
	_ = v373
	var v380 int32
	_ = v380
	var v383 int32
	_ = v383
	var v386 int32
	_ = v386
	var v389 int32
	_ = v389
	var v390 int32
	_ = v390
	var v392 int32
	_ = v392
	var v393 int32
	_ = v393
	var v396 int32
	_ = v396
	var v397 int32
	_ = v397
	var v400 int32
	_ = v400
	var v401 int32
	_ = v401
	var v407 int32
	_ = v407
	var v422 int32
	_ = v422
	var v423 int32
	_ = v423
	var v425 int32
	_ = v425
	var v440 int32
	_ = v440
	v12 = m.G0
	v14 = v12 - int32(16)
	m.G0 = v14
	v16 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v16 != int32(35) {
		goto L4
	} else {
		goto L5
	}
L1:
	;
	m.G0 = v14 + int32(16)
	return v440
L2:
	;
	v440 = int32(-1)
	goto L1
L3:
	;
	v28 = int32(10)
	v29 = F___strchrnul(m, l0, v28)
	mBase = m.M
	v31 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v29))))
	if v31 == v28 {
		goto L11
	} else {
		goto L12
	}
L4:
	;
	v23 = F_sdsnew(m, int32(_a630))
	mBase = m.M
	v26 = m.ExcPending
	if v26 != 0 {
		goto L7
	} else {
		goto L8
	}
L5:
	;
	v19 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+1)))
	if v19 == int32(33) {
		goto L3
	} else {
		goto L6
	}
L6:
	;
	goto L4
L7:
	;
	return int32(0)
L8:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v23
	goto L2
L9:
	;
	v40 = v35 - l0
	v43 = F_sdsnsplitargs(m, l0, v40, v14+int32(12))
	mBase = m.M
	v44 = m.ExcPending
	if v44 != 0 {
		goto L7
	} else {
		goto L18
	}
L10:
	;
	if v35 != 0 {
		goto L9
	} else {
		goto L14
	}
L11:
	;
	v35 = v29
	goto L13
L12:
	;
	v35 = int32(0)
	goto L13
L13:
	;
	goto L10
L14:
	;
	v37 = F_sdsnew(m, int32(_a631))
	mBase = m.M
	v38 = m.ExcPending
	if v38 != 0 {
		goto L7
	} else {
		goto L15
	}
L15:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v37
	goto L2
L16:
	;
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v43)))
	v56 = F_sdsdup(m, v55)
	mBase = m.M
	v57 = m.ExcPending
	if v57 != 0 {
		goto L7
	} else {
		goto L23
	}
L17:
	;
	v49 = F_sdsnew(m, int32(_a631))
	mBase = m.M
	v50 = m.ExcPending
	if v50 != 0 {
		goto L7
	} else {
		goto L21
	}
L18:
	;
	if v43 == int32(0) {
		goto L17
	} else {
		goto L19
	}
L19:
	;
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v14)+12))
	if v47 != 0 {
		goto L16
	} else {
		goto L20
	}
L20:
	;
	goto L17
L21:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v49
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v14)+12))
	F_sdsfreesplitres(m, v43, v52)
	mBase = m.M
	v54 = m.ExcPending
	if v54 != 0 {
		goto L7
	} else {
		goto L22
	}
L22:
	;
	goto L2
L23:
	;
	v59 = int32(-1)
	v67 = v56 + v59
	v68 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v67))))
	v70 = v68 & int32(7)
	switch v70 {
	case 0:
		goto L31
	case 1:
		goto L30
	case 2:
		goto L29
	case 3:
		goto L28
	case 4:
		goto L27
	default:
		goto L25
	}
L24:
	;
	v160 = *(*int32)(unsafe.Add(mBase, uint32(v14)+12))
	if v160 < int32(2) {
		goto L62
	} else {
		goto L63
	}
L25:
	;
	goto L24
L26:
	;
	if v85 == int32(0) {
		goto L25
	} else {
		goto L32
	}
L27:
	;
	v84 = *(*int32)(unsafe.Add(mBase, uint32(v56+int32(-17))))
	v85 = v84
	goto L26
L28:
	;
	v81 = *(*int32)(unsafe.Add(mBase, uint32(v56+int32(-9))))
	v85 = v81
	goto L26
L29:
	;
	v78 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v56+int32(-5)))))
	v85 = v78
	goto L26
L30:
	;
	v75 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v56+int32(-3)))))
	v85 = v75
	goto L26
L31:
	;
	v85 = int32(base.Ui32(v68) >> (uint(int32(3)) % 32))
	goto L26
L32:
	;
	v91 = int32(-1)&v85 + v59
	v95 = int32(0)&v85 + int32(2)
	v98 = v91 - v95 + int32(1)
	switch v70 {
	default:
		goto L38
	case 1:
		goto L37
	case 2:
		goto L36
	case 3:
		goto L35
	case 4:
		goto L34
	}
L33:
	;
	v114 = int32(0)
	v116 = base.B2i32(base.Ui32(v95) < base.Ui32(v113))
	if base.Ui32(v95) < base.Ui32(v113) {
		goto L40
	} else {
		goto L41
	}
L34:
	;
	v112 = *(*int32)(unsafe.Add(mBase, uint32(v56+int32(-17))))
	v113 = v112
	goto L33
L35:
	;
	v109 = *(*int32)(unsafe.Add(mBase, uint32(v56+int32(-9))))
	v113 = v109
	goto L33
L36:
	;
	v106 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v56+int32(-5)))))
	v113 = v106
	goto L33
L37:
	;
	v103 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v56+int32(-3)))))
	v113 = v103
	goto L33
L38:
	;
	v113 = int32(base.Ui32(v68) >> (uint(int32(3)) % 32))
	goto L33
L39:
	;
	v130 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v56+v124))) = uint8(v130)
	switch v70 {
	default:
		goto L57
	case 1:
		goto L56
	case 2:
		goto L55
	case 3:
		goto L54
	case 4:
		goto L53
	}
L40:
	;
	v117 = v95
	goto L42
L41:
	;
	v117 = v114
	goto L42
L42:
	;
	v118 = v113 - v117
	if base.Ui32(v98) < base.Ui32(v118) {
		goto L43
	} else {
		goto L44
	}
L43:
	;
	v120 = v98
	goto L45
L44:
	;
	v120 = v118
	goto L45
L45:
	;
	if v91 < v95 {
		goto L46
	} else {
		goto L47
	}
L46:
	;
	v122 = v114
	goto L48
L47:
	;
	v122 = v120
	goto L48
L48:
	;
	if base.Ui32(v95) < base.Ui32(v113) {
		goto L49
	} else {
		goto L50
	}
L49:
	;
	v124 = v122
	goto L51
L50:
	;
	v124 = int32(0)
	goto L51
L51:
	;
	if v124 == int32(0) {
		goto L39
	} else {
		goto L52
	}
L52:
	;
	v128 = F_memmove(m, v56, v56+v117, v124)
	mBase = m.M
	goto L39
L53:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v56+int32(-17)))) = base.I64_extend_i32_u(v124)
	goto L25
L54:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v56+int32(-9)))) = v124
	goto L24
L55:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v56+int32(-5)))) = uint16(v124)
	goto L24
L56:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v56+int32(-3)))) = uint8(v124)
	goto L24
L57:
	;
	v133 = v124 << (uint(int32(3)) % 32)
	*(*uint8)(unsafe.Add(mBase, uint32(v67))) = uint8(v133)
	goto L24
L58:
	;
	if v56 == int32(0) {
		goto L136
	} else {
		goto L137
	}
L59:
	;
	F_sdsfree(m, v173)
	mBase = m.M
	v407 = m.ExcPending
	if v407 != 0 {
		goto L7
	} else {
		goto L135
	}
L60:
	;
	v396 = F_sdsempty(m)
	mBase = m.M
	v397 = m.ExcPending
	if v397 != 0 {
		goto L7
	} else {
		goto L132
	}
L61:
	;
	F_sdsfreesplitres(m, v43, v349)
	mBase = m.M
	v367 = m.ExcPending
	if v367 != 0 {
		goto L7
	} else {
		goto L124
	}
L62:
	;
	v363 = F_sdsnew(m, int32(_a632))
	mBase = m.M
	v364 = m.ExcPending
	if v364 != 0 {
		goto L7
	} else {
		goto L123
	}
L63:
	;
	v173 = int32(0)
	v174 = int32(1)
	goto L64
L64:
	;
	v179 = *(*int32)(unsafe.Add(mBase, uint32(v43+v174<<(uint(int32(2))%32))))
	v180 = int32(_a633)
	goto L67
L65:
	;
	if v243 != 0 {
		goto L61
	} else {
		goto L122
	}
L66:
	;
	if v224-v226 != 0 {
		goto L60
	} else {
		goto L81
	}
L67:
	;
	v185 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v179))))
	if v185 != 0 {
		goto L70
	} else {
		goto L71
	}
L69:
	;
	v224 = F_tolower(m, v219)
	mBase = m.M
	v225 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v220))))
	v226 = F_tolower(m, v225)
	mBase = m.M
	goto L66
L70:
	;
	v187 = v179
	v188 = v180
	v189 = int32(5)
	v190 = v185
	goto L73
L71:
	;
	v219 = int32(0)
	v220 = v180
	goto L69
L72:
	;
	v219 = v216 & int32(255)
	v220 = v214
	goto L69
L73:
	;
	v192 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v188))))
	if v192 == int32(0) {
		v214 = v188
		v216 = v190
		goto L72
	} else {
		goto L75
	}
L74:
	;
	v214 = v208
	v216 = int32(0)
	goto L72
L75:
	;
	v196 = v189 + int32(-1)
	if v196 == int32(0) {
		v214 = v188
		v216 = v190
		goto L72
	} else {
		goto L76
	}
L76:
	;
	v200 = v190 & int32(255)
	if v200 == v192 {
		goto L77
	} else {
		goto L78
	}
L77:
	;
	v207 = int32(1)
	v208 = v188 + v207
	v209 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v187)+1)))
	if v209 != 0 {
		v187 = v187 + v207
		v188 = v208
		v189 = v196
		v190 = v209
		goto L73
	} else {
		goto L80
	}
L78:
	;
	v202 = F_tolower(m, v200)
	mBase = m.M
	v203 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v188))))
	v204 = F_tolower(m, v203)
	mBase = m.M
	if v202 == v204 {
		goto L77
	} else {
		goto L79
	}
L79:
	;
	v206 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v187))))
	v214 = v188
	v216 = v206
	goto L72
L80:
	;
	goto L74
L81:
	;
	if v173 == int32(0) {
		goto L82
	} else {
		goto L83
	}
L82:
	;
	v243 = F_sdsdup(m, v179)
	mBase = m.M
	v244 = m.ExcPending
	if v244 != 0 {
		goto L7
	} else {
		goto L86
	}
L83:
	;
	v236 = F_sdsempty(m)
	mBase = m.M
	v237 = m.ExcPending
	if v237 != 0 {
		goto L7
	} else {
		goto L84
	}
L84:
	;
	v240 = F_sdscatfmt(m, v236, int32(_a634), int32(0))
	mBase = m.M
	v241 = m.ExcPending
	if v241 != 0 {
		goto L7
	} else {
		goto L85
	}
L85:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v240
	goto L59
L86:
	;
	v246 = int32(-1)
	v254 = v243 + v246
	v255 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v254))))
	v257 = v255 & int32(7)
	switch v257 {
	case 0:
		goto L94
	case 1:
		goto L93
	case 2:
		goto L92
	case 3:
		goto L91
	case 4:
		goto L90
	default:
		goto L88
	}
L87:
	;
	v348 = v174 + int32(1)
	v349 = *(*int32)(unsafe.Add(mBase, uint32(v14)+12))
	if v348 < v349 {
		v173 = v243
		v174 = v348
		goto L64
	} else {
		goto L121
	}
L88:
	;
	goto L87
L89:
	;
	if v272 == int32(0) {
		goto L88
	} else {
		goto L95
	}
L90:
	;
	v271 = *(*int32)(unsafe.Add(mBase, uint32(v243+int32(-17))))
	v272 = v271
	goto L89
L91:
	;
	v268 = *(*int32)(unsafe.Add(mBase, uint32(v243+int32(-9))))
	v272 = v268
	goto L89
L92:
	;
	v265 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v243+int32(-5)))))
	v272 = v265
	goto L89
L93:
	;
	v262 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v243+int32(-3)))))
	v272 = v262
	goto L89
L94:
	;
	v272 = int32(base.Ui32(v255) >> (uint(int32(3)) % 32))
	goto L89
L95:
	;
	v278 = int32(-1)&v272 + v246
	v282 = int32(0)&v272 + int32(5)
	v285 = v278 - v282 + int32(1)
	switch v257 {
	default:
		goto L101
	case 1:
		goto L100
	case 2:
		goto L99
	case 3:
		goto L98
	case 4:
		goto L97
	}
L96:
	;
	v301 = int32(0)
	v303 = base.B2i32(base.Ui32(v282) < base.Ui32(v300))
	if base.Ui32(v282) < base.Ui32(v300) {
		goto L103
	} else {
		goto L104
	}
L97:
	;
	v299 = *(*int32)(unsafe.Add(mBase, uint32(v243+int32(-17))))
	v300 = v299
	goto L96
L98:
	;
	v296 = *(*int32)(unsafe.Add(mBase, uint32(v243+int32(-9))))
	v300 = v296
	goto L96
L99:
	;
	v293 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v243+int32(-5)))))
	v300 = v293
	goto L96
L100:
	;
	v290 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v243+int32(-3)))))
	v300 = v290
	goto L96
L101:
	;
	v300 = int32(base.Ui32(v255) >> (uint(int32(3)) % 32))
	goto L96
L102:
	;
	v317 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v243+v311))) = uint8(v317)
	switch v257 {
	default:
		goto L120
	case 1:
		goto L119
	case 2:
		goto L118
	case 3:
		goto L117
	case 4:
		goto L116
	}
L103:
	;
	v304 = v282
	goto L105
L104:
	;
	v304 = v301
	goto L105
L105:
	;
	v305 = v300 - v304
	if base.Ui32(v285) < base.Ui32(v305) {
		goto L106
	} else {
		goto L107
	}
L106:
	;
	v307 = v285
	goto L108
L107:
	;
	v307 = v305
	goto L108
L108:
	;
	if v278 < v282 {
		goto L109
	} else {
		goto L110
	}
L109:
	;
	v309 = v301
	goto L111
L110:
	;
	v309 = v307
	goto L111
L111:
	;
	if base.Ui32(v282) < base.Ui32(v300) {
		goto L112
	} else {
		goto L113
	}
L112:
	;
	v311 = v309
	goto L114
L113:
	;
	v311 = int32(0)
	goto L114
L114:
	;
	if v311 == int32(0) {
		goto L102
	} else {
		goto L115
	}
L115:
	;
	v315 = F_memmove(m, v243, v243+v304, v311)
	mBase = m.M
	goto L102
L116:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v243+int32(-17)))) = base.I64_extend_i32_u(v311)
	goto L88
L117:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v243+int32(-9)))) = v311
	goto L87
L118:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v243+int32(-5)))) = uint16(v311)
	goto L87
L119:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v243+int32(-3)))) = uint8(v311)
	goto L87
L120:
	;
	v320 = v311 << (uint(int32(3)) % 32)
	*(*uint8)(unsafe.Add(mBase, uint32(v254))) = uint8(v320)
	goto L87
L121:
	;
	goto L65
L122:
	;
	goto L62
L123:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v363
	goto L58
L124:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v243
	v369 = int32(0)
	v373 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0+int32(-1)))))
	switch v373 & int32(7) {
	case 0:
		goto L130
	case 1:
		goto L129
	case 2:
		goto L128
	case 3:
		goto L127
	case 4:
		goto L126
	default:
		v390 = v369
		goto L125
	}
L125:
	;
	v392 = F_sdsnewlen(m, v35, v390-v40)
	mBase = m.M
	v393 = m.ExcPending
	if v393 != 0 {
		goto L7
	} else {
		goto L131
	}
L126:
	;
	v389 = *(*int32)(unsafe.Add(mBase, uint32(l0+int32(-17))))
	v390 = v389
	goto L125
L127:
	;
	v386 = *(*int32)(unsafe.Add(mBase, uint32(l0+int32(-9))))
	v390 = v386
	goto L125
L128:
	;
	v383 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0+int32(-5)))))
	v390 = v383
	goto L125
L129:
	;
	v380 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0+int32(-3)))))
	v390 = v380
	goto L125
L130:
	;
	v390 = int32(base.Ui32(v373) >> (uint(int32(3)) % 32))
	goto L125
L131:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v56
	*(*int32)(unsafe.Add(mBase, uint32(l1)+8)) = v392
	v440 = v369
	goto L1
L132:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14))) = v179
	v400 = F_sdscatfmt(m, v396, int32(_a635), v14)
	mBase = m.M
	v401 = m.ExcPending
	if v401 != 0 {
		goto L7
	} else {
		goto L133
	}
L133:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v400
	if v173 == int32(0) {
		goto L58
	} else {
		goto L134
	}
L134:
	;
	goto L59
L135:
	;
	goto L58
L136:
	;
	v423 = *(*int32)(unsafe.Add(mBase, uint32(v14)+12))
	F_sdsfreesplitres(m, v43, v423)
	mBase = m.M
	v425 = m.ExcPending
	if v425 != 0 {
		goto L7
	} else {
		goto L139
	}
L137:
	;
	F_sdsfree(m, v56)
	mBase = m.M
	v422 = m.ExcPending
	if v422 != 0 {
		goto L7
	} else {
		goto L138
	}
L138:
	;
	goto L136
L139:
	;
	goto L2
}
func F_functionFlushCommand(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
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
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
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
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v104 int32
	_ = v104
	var v107 int32
	_ = v107
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v133 int32
	_ = v133
	var v135 int64
	_ = v135
	var v140 int32
	_ = v140
	var v142 int32
	_ = v142
	var v145 int32
	_ = v145
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v3 < int32(4) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	if v3 != int32(3) {
		v100 = v3
		goto L7
	} else {
		goto L8
	}
L2:
	;
	F_addReplySubcommandSyntaxError(m, l0)
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	return
L4:
	;
	return
L5:
	;
	F_addReplyError(m, l0, int32(_a626))
	mBase = m.M
	v145 = m.ExcPending
	if v145 != 0 {
		goto L3
	} else {
		goto L46
	}
L6:
	;
	F_functionsLibCtxReleaseCurrent(m, v107, int32(0))
	mBase = m.M
	v110 = m.ExcPending
	if v110 != 0 {
		goto L3
	} else {
		goto L39
	}
L7:
	;
	if v100 != int32(2) {
		goto L5
	} else {
		goto L38
	}
L8:
	;
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(v10)+8))
	v12 = F_objectGetVal(m, v11)
	mBase = m.M
	v13 = int32(_a627)
	v16 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12))))
	if v16 != 0 {
		goto L12
	} else {
		goto L13
	}
L9:
	;
	v53 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v53 != int32(3) {
		v100 = v53
		goto L7
	} else {
		goto L23
	}
L10:
	;
	if v48-v50 != 0 {
		goto L9
	} else {
		goto L22
	}
L11:
	;
	v48 = F_tolower(m, v44)
	mBase = m.M
	v49 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v45))))
	v50 = F_tolower(m, v49)
	mBase = m.M
	goto L10
L12:
	;
	v18 = v12
	v19 = v13
	v20 = v16
	goto L15
L13:
	;
	v44 = int32(0)
	v45 = v13
	goto L11
L14:
	;
	v44 = v41 & int32(255)
	v45 = v40
	goto L11
L15:
	;
	v22 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19))))
	if v22 == int32(0) {
		v40 = v19
		v41 = v20
		goto L14
	} else {
		goto L17
	}
L16:
	;
	v40 = v34
	v41 = int32(0)
	goto L14
L17:
	;
	v26 = v20 & int32(255)
	if v26 == v22 {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	v33 = int32(1)
	v34 = v19 + v33
	v35 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18)+1)))
	if v35 != 0 {
		v18 = v18 + v33
		v19 = v34
		v20 = v35
		goto L15
	} else {
		goto L21
	}
L19:
	;
	v28 = F_tolower(m, v26)
	mBase = m.M
	v29 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19))))
	v30 = F_tolower(m, v29)
	mBase = m.M
	if v28 == v30 {
		goto L18
	} else {
		goto L20
	}
L20:
	;
	v32 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18))))
	v40 = v19
	v41 = v32
	goto L14
L21:
	;
	goto L16
L22:
	;
	v107 = int32(0)
	goto L6
L23:
	;
	v56 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v56)+8))
	v58 = F_objectGetVal(m, v57)
	mBase = m.M
	v59 = int32(_a628)
	v62 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v58))))
	if v62 != 0 {
		goto L27
	} else {
		goto L28
	}
L24:
	;
	v99 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v100 = v99
	goto L7
L25:
	;
	if v94-v96 != 0 {
		goto L24
	} else {
		goto L37
	}
L26:
	;
	v94 = F_tolower(m, v90)
	mBase = m.M
	v95 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v91))))
	v96 = F_tolower(m, v95)
	mBase = m.M
	goto L25
L27:
	;
	v64 = v58
	v65 = v59
	v66 = v62
	goto L30
L28:
	;
	v90 = int32(0)
	v91 = v59
	goto L26
L29:
	;
	v90 = v87 & int32(255)
	v91 = v86
	goto L26
L30:
	;
	v68 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v65))))
	if v68 == int32(0) {
		v86 = v65
		v87 = v66
		goto L29
	} else {
		goto L32
	}
L31:
	;
	v86 = v80
	v87 = int32(0)
	goto L29
L32:
	;
	v72 = v66 & int32(255)
	if v72 == v68 {
		goto L33
	} else {
		goto L34
	}
L33:
	;
	v79 = int32(1)
	v80 = v65 + v79
	v81 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v64)+1)))
	if v81 != 0 {
		v64 = v64 + v79
		v65 = v80
		v66 = v81
		goto L30
	} else {
		goto L36
	}
L34:
	;
	v74 = F_tolower(m, v72)
	mBase = m.M
	v75 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v65))))
	v76 = F_tolower(m, v75)
	mBase = m.M
	if v74 == v76 {
		goto L33
	} else {
		goto L35
	}
L35:
	;
	v78 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v64))))
	v86 = v65
	v87 = v78
	goto L29
L36:
	;
	goto L31
L37:
	;
	v107 = int32(1)
	goto L6
L38:
	;
	v104 = *(*int32)(unsafe.Add(mBase, _consts[337]))
	v107 = base.B2i32(v104 != int32(0))
	goto L6
L39:
	;
	v112 = F_valkey_malloc(m, int32(16))
	mBase = m.M
	v113 = m.ExcPending
	if v113 != 0 {
		goto L3
	} else {
		goto L40
	}
L40:
	;
	v115 = F_dictCreate(m, int32(_a617))
	mBase = m.M
	v116 = m.ExcPending
	if v116 != 0 {
		goto L3
	} else {
		goto L41
	}
L41:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v112))) = v115
	v119 = F_dictCreate(m, int32(_a618))
	mBase = m.M
	v120 = m.ExcPending
	if v120 != 0 {
		goto L3
	} else {
		goto L42
	}
L42:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v112)+4)) = v119
	v123 = F_dictCreate(m, int32(_a619))
	mBase = m.M
	v124 = m.ExcPending
	if v124 != 0 {
		goto L3
	} else {
		goto L43
	}
L43:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v112)+12)) = v123
	F_scriptingEngineManagerForEachEngine(m, int32(529), v112)
	mBase = m.M
	v128 = m.ExcPending
	if v128 != 0 {
		goto L3
	} else {
		goto L44
	}
L44:
	;
	v129 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v112)+8)) = v129
	*(*int32)(unsafe.Add(mBase, _consts[332])) = v112
	v133 = int32(_a20)
	v135 = *(*int64)(unsafe.Add(mBase, _consts[180]))
	*(*int64)(unsafe.Add(mBase, _consts[180])) = v135 + int64(1)
	v140 = *(*int32)(unsafe.Add(mBase, _consts[27]))
	F_addReply(m, l0, v140)
	mBase = m.M
	v142 = m.ExcPending
	if v142 != 0 {
		goto L3
	} else {
		goto L45
	}
L45:
	;
	return
L46:
	;
	return
}
func F_functionGetKeys(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	v9 = F_genericGetKeys(m, int32(0), int32(2), int32(3), int32(1), l1, l2, l3)
	v12 = m.ExcPending
	if v12 != 0 {
		return int32(0)
	} else {
		return v9
	}
}
func F_functionHelpCommand(m *base.Module, l0 int32) {
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	v3 = m.G0
	v5 = v3 - int32(160)
	m.G0 = v5
	v11 = F__emscripten_memcpy_bulkmem(m, v5, int32(_a629), int32(148))
	F_addReplyHelp(m, l0, v11)
	v14 = m.ExcPending
	if v14 != 0 {
		return
	} else {
		m.G0 = v11 + int32(160)
		return
	}
}
func F_functionStatsCommand(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v9 int32
	_ = v9
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v58 int32
	_ = v58
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v83 int32
	_ = v83
	var v86 int32
	_ = v86
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v109 int32
	_ = v109
	var v110 int64
	_ = v110
	var v111 int32
	_ = v111
	var v113 int32
	_ = v113
	var v123 int32
	_ = v123
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v131 int32
	_ = v131
	var v134 int32
	_ = v134
	v2 = int32(0)
	v9 = *(*int32)(unsafe.Add(mBase, _consts[333]))
	goto L2
L1:
	;
	F_addReplyMapLen(m, l0, int32(2))
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		goto L4
	} else {
		goto L8
	}
L2:
	;
	if base.B2i32(v9 != v2) == int32(0) {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v14 = F_scriptIsEval(m)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	return
L5:
	;
	if v14 == int32(0) {
		goto L1
	} else {
		goto L6
	}
L6:
	;
	v19 = *(*int32)(unsafe.Add(mBase, _consts[334]))
	F_addReplyErrorObject(m, l0, v19)
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		goto L4
	} else {
		goto L7
	}
L7:
	;
	return
L8:
	;
	F_addReplyBulkCString(m, l0, int32(_a620))
	mBase = m.M
	v27 = m.ExcPending
	if v27 != 0 {
		goto L4
	} else {
		goto L9
	}
L9:
	;
	v28 = int32(0)
	v29 = *(*int32)(unsafe.Add(mBase, _consts[333]))
	goto L12
L10:
	;
	F_addReplyBulkCString(m, l0, int32(_a621))
	mBase = m.M
	v123 = m.ExcPending
	if v123 != 0 {
		goto L4
	} else {
		goto L37
	}
L11:
	;
	F_addReplyMapLen(m, l0, int32(3))
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
		goto L4
	} else {
		goto L15
	}
L12:
	;
	if v29 != v28 {
		goto L11
	} else {
		goto L13
	}
L13:
	;
	F_addReplyNull(m, l0)
	mBase = m.M
	v33 = m.ExcPending
	if v33 != 0 {
		goto L4
	} else {
		goto L14
	}
L14:
	;
	goto L10
L15:
	;
	F_addReplyBulkCString(m, l0, int32(_a373))
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L4
	} else {
		goto L16
	}
L16:
	;
	v40 = F_scriptCurrFunction(m)
	mBase = m.M
	v41 = m.ExcPending
	if v41 != 0 {
		goto L4
	} else {
		goto L17
	}
L17:
	;
	F_addReplyBulkCString(m, l0, v40)
	mBase = m.M
	v43 = m.ExcPending
	if v43 != 0 {
		goto L4
	} else {
		goto L18
	}
L18:
	;
	F_addReplyBulkCString(m, l0, int32(_a622))
	mBase = m.M
	v46 = m.ExcPending
	if v46 != 0 {
		goto L4
	} else {
		goto L19
	}
L19:
	;
	v47 = F_scriptGetCaller(m)
	mBase = m.M
	v48 = m.ExcPending
	if v48 != 0 {
		goto L4
	} else {
		goto L20
	}
L20:
	;
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v47)+24))
	F_addReplyArrayLen(m, l0, v49)
	mBase = m.M
	v51 = m.ExcPending
	if v51 != 0 {
		goto L4
	} else {
		goto L21
	}
L21:
	;
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v47)+24))
	if v52 < int32(1) {
		goto L22
	} else {
		goto L23
	}
L22:
	;
	F_addReplyBulkCString(m, l0, int32(_a623))
	mBase = m.M
	v109 = m.ExcPending
	if v109 != 0 {
		goto L4
	} else {
		goto L34
	}
L23:
	;
	v58 = int32(0)
	goto L24
L24:
	;
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v47)+20))
	v65 = v58 << (uint(int32(2)) % 32)
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v63+v65)))
	v68 = F_objectGetVal(m, v67)
	mBase = m.M
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v47)+20))
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v70+v65)))
	v73 = F_objectGetVal(m, v72)
	mBase = m.M
	v76 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v73+int32(-1)))))
	switch v76 & int32(7) {
	case 0:
		goto L31
	case 1:
		goto L30
	case 2:
		goto L29
	case 3:
		goto L28
	case 4:
		goto L27
	default:
		v93 = int32(0)
		goto L26
	}
L25:
	;
	goto L22
L26:
	;
	F_addReplyBulkCBuffer(m, l0, v68, v93)
	mBase = m.M
	v95 = m.ExcPending
	if v95 != 0 {
		goto L4
	} else {
		goto L32
	}
L27:
	;
	v92 = *(*int32)(unsafe.Add(mBase, uint32(v73+int32(-17))))
	v93 = v92
	goto L26
L28:
	;
	v89 = *(*int32)(unsafe.Add(mBase, uint32(v73+int32(-9))))
	v93 = v89
	goto L26
L29:
	;
	v86 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v73+int32(-5)))))
	v93 = v86
	goto L26
L30:
	;
	v83 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v73+int32(-3)))))
	v93 = v83
	goto L26
L31:
	;
	v93 = int32(base.Ui32(v76) >> (uint(int32(3)) % 32))
	goto L26
L32:
	;
	v97 = v58 + int32(1)
	v98 = *(*int32)(unsafe.Add(mBase, uint32(v47)+24))
	if v97 < v98 {
		v58 = v97
		goto L24
	} else {
		goto L33
	}
L33:
	;
	goto L25
L34:
	;
	v110 = F_scriptRunDuration(m)
	mBase = m.M
	v111 = m.ExcPending
	if v111 != 0 {
		goto L4
	} else {
		goto L35
	}
L35:
	;
	F_addReplyLongLong(m, l0, v110)
	mBase = m.M
	v113 = m.ExcPending
	if v113 != 0 {
		goto L4
	} else {
		goto L36
	}
L36:
	;
	goto L10
L37:
	;
	v126 = *(*int32)(unsafe.Add(mBase, _consts[335]))
	v127 = *(*int32)(unsafe.Add(mBase, uint32(v126)+16))
	v128 = *(*int32)(unsafe.Add(mBase, uint32(v126)+12))
	goto L38
L38:
	;
	F_addReplyMapLen(m, l0, v127+v128)
	mBase = m.M
	v131 = m.ExcPending
	if v131 != 0 {
		goto L4
	} else {
		goto L39
	}
L39:
	;
	F_scriptingEngineManagerForEachEngine(m, int32(530), l0)
	mBase = m.M
	v134 = m.ExcPending
	if v134 != 0 {
		goto L4
	} else {
		goto L40
	}
L40:
	;
	return
}
