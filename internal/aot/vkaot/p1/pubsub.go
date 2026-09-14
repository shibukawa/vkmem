package p1

import (
	base "github.com/shibukawa/vkmem/internal/aot/vkaot/base"
	"unsafe"
)

func F_pubsubCommand(m *base.Module, l0 int32) {
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
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v60 int32
	_ = v60
	var v61 int64
	_ = v61
	var v66 int64
	_ = v66
	var v71 int64
	_ = v71
	var v76 int64
	_ = v76
	var v79 int64
	_ = v79
	var v82 int64
	_ = v82
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
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
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v102 int32
	_ = v102
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v128 int32
	_ = v128
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v141 int32
	_ = v141
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
	var v151 int32
	_ = v151
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v157 int32
	_ = v157
	var v161 int32
	_ = v161
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v203 int32
	_ = v203
	var v208 int32
	_ = v208
	var v211 int32
	_ = v211
	var v213 int32
	_ = v213
	var v215 int32
	_ = v215
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v223 int32
	_ = v223
	var v225 int32
	_ = v225
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v231 int64
	_ = v231
	var v233 int32
	_ = v233
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v246 int32
	_ = v246
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
	var v258 int32
	_ = v258
	var v259 int32
	_ = v259
	var v260 int32
	_ = v260
	var v262 int32
	_ = v262
	var v263 int32
	_ = v263
	var v264 int32
	_ = v264
	var v265 int32
	_ = v265
	var v270 int32
	_ = v270
	var v271 int32
	_ = v271
	var v274 int32
	_ = v274
	var v275 int32
	_ = v275
	var v278 int32
	_ = v278
	var v279 int32
	_ = v279
	var v280 int32
	_ = v280
	var v282 int32
	_ = v282
	var v286 int32
	_ = v286
	var v287 int32
	_ = v287
	var v288 int32
	_ = v288
	var v292 int32
	_ = v292
	var v293 int32
	_ = v293
	var v294 int32
	_ = v294
	var v295 int32
	_ = v295
	var v296 int32
	_ = v296
	var v299 int32
	_ = v299
	var v301 int32
	_ = v301
	var v302 int32
	_ = v302
	var v303 int32
	_ = v303
	var v305 int32
	_ = v305
	var v309 int32
	_ = v309
	var v311 int32
	_ = v311
	var v312 int32
	_ = v312
	var v313 int32
	_ = v313
	var v315 int32
	_ = v315
	var v316 int32
	_ = v316
	var v317 int32
	_ = v317
	var v318 int32
	_ = v318
	var v323 int32
	_ = v323
	var v324 int32
	_ = v324
	var v327 int32
	_ = v327
	var v328 int32
	_ = v328
	var v331 int32
	_ = v331
	var v332 int32
	_ = v332
	var v333 int32
	_ = v333
	var v335 int32
	_ = v335
	var v343 int32
	_ = v343
	var v344 int32
	_ = v344
	var v345 int32
	_ = v345
	var v346 int32
	_ = v346
	var v348 int32
	_ = v348
	var v350 int32
	_ = v350
	var v352 int32
	_ = v352
	var v353 int32
	_ = v353
	var v354 int32
	_ = v354
	var v355 int32
	_ = v355
	var v358 int32
	_ = v358
	var v360 int32
	_ = v360
	var v361 int32
	_ = v361
	var v362 int32
	_ = v362
	var v364 int32
	_ = v364
	var v368 int32
	_ = v368
	var v370 int32
	_ = v370
	var v371 int32
	_ = v371
	var v372 int32
	_ = v372
	var v374 int32
	_ = v374
	var v375 int32
	_ = v375
	var v376 int32
	_ = v376
	var v377 int32
	_ = v377
	var v382 int32
	_ = v382
	var v383 int32
	_ = v383
	var v386 int32
	_ = v386
	var v387 int32
	_ = v387
	var v390 int32
	_ = v390
	var v391 int32
	_ = v391
	var v392 int32
	_ = v392
	var v394 int32
	_ = v394
	var v395 int32
	_ = v395
	var v403 int32
	_ = v403
	var v404 int32
	_ = v404
	var v409 int32
	_ = v409
	var v415 int32
	_ = v415
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
	var v429 int32
	_ = v429
	var v436 int32
	_ = v436
	var v439 int32
	_ = v439
	var v442 int32
	_ = v442
	var v445 int32
	_ = v445
	var v446 int32
	_ = v446
	var v447 int32
	_ = v447
	var v455 int32
	_ = v455
	var v459 int32
	_ = v459
	var v463 int32
	_ = v463
	var v467 int32
	_ = v467
	var v476 int32
	_ = v476
	var v478 int32
	_ = v478
	var v484 int32
	_ = v484
	var v486 int32
	_ = v486
	var v490 int32
	_ = v490
	var v499 int32
	_ = v499
	var v506 int32
	_ = v506
	var v509 int32
	_ = v509
	var v514 int32
	_ = v514
	var v519 int32
	_ = v519
	var v520 int32
	_ = v520
	var v522 int32
	_ = v522
	var v523 int32
	_ = v523
	var v524 int32
	_ = v524
	var v525 int32
	_ = v525
	var v526 int32
	_ = v526
	var v528 int32
	_ = v528
	var v530 int32
	_ = v530
	var v532 int32
	_ = v532
	var v533 int32
	_ = v533
	var v536 int64
	_ = v536
	var v538 int32
	_ = v538
	var v540 int32
	_ = v540
	var v541 int32
	_ = v541
	var v546 int32
	_ = v546
	v9 = m.G0
	v11 = v9 - int32(48)
	m.G0 = v11
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v13 != int32(2) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v11 + int32(48)
	return
L2:
	;
	v86 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v87 = *(*int32)(unsafe.Add(mBase, uint32(v86)+4))
	v88 = F_objectGetVal(m, v87)
	mBase = m.M
	v89 = int32(_a_F_pubsubCommand_0)
	v92 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v88))))
	if v92 != 0 {
		goto L22
	} else {
		goto L23
	}
L3:
	;
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v17 = *(*int32)(unsafe.Add(mBase, uint32(v16)+4))
	v18 = F_objectGetVal(m, v17)
	mBase = m.M
	v19 = int32(_a_F_pubsubCommand_1)
	v22 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18))))
	if v22 != 0 {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	if v54-v56 != 0 {
		goto L2
	} else {
		goto L16
	}
L5:
	;
	v54 = F_tolower(m, v50)
	mBase = m.M
	v55 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v51))))
	v56 = F_tolower(m, v55)
	mBase = m.M
	goto L4
L6:
	;
	v24 = v18
	v25 = v19
	v26 = v22
	goto L9
L7:
	;
	v50 = int32(0)
	v51 = v19
	goto L5
L8:
	;
	v50 = v47 & int32(255)
	v51 = v46
	goto L5
L9:
	;
	v28 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v25))))
	if v28 == int32(0) {
		v46 = v25
		v47 = v26
		goto L8
	} else {
		goto L11
	}
L10:
	;
	v46 = v40
	v47 = int32(0)
	goto L8
L11:
	;
	v32 = v26 & int32(255)
	if v32 == v28 {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	v39 = int32(1)
	v40 = v25 + v39
	v41 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24)+1)))
	if v41 != 0 {
		v24 = v24 + v39
		v25 = v40
		v26 = v41
		goto L9
	} else {
		goto L15
	}
L13:
	;
	v34 = F_tolower(m, v32)
	mBase = m.M
	v35 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v25))))
	v36 = F_tolower(m, v35)
	mBase = m.M
	if v34 == v36 {
		goto L12
	} else {
		goto L14
	}
L14:
	;
	v38 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24))))
	v46 = v25
	v47 = v38
	goto L8
L15:
	;
	goto L10
L16:
	;
	v60 = int32(0)
	v61 = *(*int64)(unsafe.Add(mBase, _c_F_pubsubCommand[0]))
	*(*int64)(unsafe.Add(mBase, uint32(v11+int32(40)))) = v61
	v66 = *(*int64)(unsafe.Add(mBase, _c_F_pubsubCommand[1]))
	*(*int64)(unsafe.Add(mBase, uint32(v11+int32(32)))) = v66
	v71 = *(*int64)(unsafe.Add(mBase, _c_F_pubsubCommand[2]))
	*(*int64)(unsafe.Add(mBase, uint32(v11+int32(24)))) = v71
	v76 = *(*int64)(unsafe.Add(mBase, _c_F_pubsubCommand[3]))
	*(*int64)(unsafe.Add(mBase, uint32(v11+int32(16)))) = v76
	v79 = *(*int64)(unsafe.Add(mBase, _c_F_pubsubCommand[4]))
	*(*int64)(unsafe.Add(mBase, uint32(v11)+8)) = v79
	v82 = *(*int64)(unsafe.Add(mBase, _c_F_pubsubCommand[5]))
	*(*int64)(unsafe.Add(mBase, uint32(v11))) = v82
	F_addReplyHelp(m, l0, v11)
	mBase = m.M
	v85 = m.ExcPending
	if v85 != 0 {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	return
L18:
	;
	goto L1
L19:
	;
	v145 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v146 = *(*int32)(unsafe.Add(mBase, uint32(v145)+4))
	v147 = F_objectGetVal(m, v146)
	mBase = m.M
	v148 = int32(_a_F_pubsubCommand_2)
	v151 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v147))))
	if v151 != 0 {
		goto L40
	} else {
		goto L41
	}
L20:
	;
	if v124-v126 != 0 {
		goto L19
	} else {
		goto L32
	}
L21:
	;
	v124 = F_tolower(m, v120)
	mBase = m.M
	v125 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v121))))
	v126 = F_tolower(m, v125)
	mBase = m.M
	goto L20
L22:
	;
	v94 = v88
	v95 = v89
	v96 = v92
	goto L25
L23:
	;
	v120 = int32(0)
	v121 = v89
	goto L21
L24:
	;
	v120 = v117 & int32(255)
	v121 = v116
	goto L21
L25:
	;
	v98 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v95))))
	if v98 == int32(0) {
		v116 = v95
		v117 = v96
		goto L24
	} else {
		goto L27
	}
L26:
	;
	v116 = v110
	v117 = int32(0)
	goto L24
L27:
	;
	v102 = v96 & int32(255)
	if v102 == v98 {
		goto L28
	} else {
		goto L29
	}
L28:
	;
	v109 = int32(1)
	v110 = v95 + v109
	v111 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v94)+1)))
	if v111 != 0 {
		v94 = v94 + v109
		v95 = v110
		v96 = v111
		goto L25
	} else {
		goto L31
	}
L29:
	;
	v104 = F_tolower(m, v102)
	mBase = m.M
	v105 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v95))))
	v106 = F_tolower(m, v105)
	mBase = m.M
	if v104 == v106 {
		goto L28
	} else {
		goto L30
	}
L30:
	;
	v108 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v94))))
	v116 = v95
	v117 = v108
	goto L24
L31:
	;
	goto L26
L32:
	;
	v128 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v128&int32(-2) != int32(2) {
		goto L19
	} else {
		goto L33
	}
L33:
	;
	if v128 == int32(2) {
		v139 = int32(0)
		goto L34
	} else {
		goto L35
	}
L34:
	;
	v141 = *(*int32)(unsafe.Add(mBase, _c_F_pubsubCommand[6]))
	F_channelList(m, l0, v139, v141)
	mBase = m.M
	v143 = m.ExcPending
	if v143 != 0 {
		goto L17
	} else {
		goto L36
	}
L35:
	;
	v136 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v137 = *(*int32)(unsafe.Add(mBase, uint32(v136)+8))
	v138 = F_objectGetVal(m, v137)
	mBase = m.M
	v139 = v138
	goto L34
L36:
	;
	goto L1
L37:
	;
	v240 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v241 = *(*int32)(unsafe.Add(mBase, uint32(v240)+4))
	v242 = F_objectGetVal(m, v241)
	mBase = m.M
	v243 = int32(_a_F_pubsubCommand_3)
	v246 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v242))))
	if v246 != 0 {
		goto L66
	} else {
		goto L67
	}
L38:
	;
	if v183-v185 != 0 {
		goto L37
	} else {
		goto L50
	}
L39:
	;
	v183 = F_tolower(m, v179)
	mBase = m.M
	v184 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v180))))
	v185 = F_tolower(m, v184)
	mBase = m.M
	goto L38
L40:
	;
	v153 = v147
	v154 = v148
	v155 = v151
	goto L43
L41:
	;
	v179 = int32(0)
	v180 = v148
	goto L39
L42:
	;
	v179 = v176 & int32(255)
	v180 = v175
	goto L39
L43:
	;
	v157 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v154))))
	if v157 == int32(0) {
		v175 = v154
		v176 = v155
		goto L42
	} else {
		goto L45
	}
L44:
	;
	v175 = v169
	v176 = int32(0)
	goto L42
L45:
	;
	v161 = v155 & int32(255)
	if v161 == v157 {
		goto L46
	} else {
		goto L47
	}
L46:
	;
	v168 = int32(1)
	v169 = v154 + v168
	v170 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v153)+1)))
	if v170 != 0 {
		v153 = v153 + v168
		v154 = v169
		v155 = v170
		goto L43
	} else {
		goto L49
	}
L47:
	;
	v163 = F_tolower(m, v161)
	mBase = m.M
	v164 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v154))))
	v165 = F_tolower(m, v164)
	mBase = m.M
	if v163 == v165 {
		goto L46
	} else {
		goto L48
	}
L48:
	;
	v167 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v153))))
	v175 = v154
	v176 = v167
	goto L42
L49:
	;
	goto L44
L50:
	;
	v187 = int32(2)
	v188 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v188 < v187 {
		goto L37
	} else {
		goto L51
	}
L51:
	;
	F_addReplyArrayLen(m, l0, v188<<(uint(int32(1))%32)+int32(-4))
	mBase = m.M
	v196 = m.ExcPending
	if v196 != 0 {
		goto L17
	} else {
		goto L52
	}
L52:
	;
	v197 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v197 < int32(3) {
		goto L1
	} else {
		goto L53
	}
L53:
	;
	v203 = v187
	goto L54
L54:
	;
	v208 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v11))) = v208
	v211 = *(*int32)(unsafe.Add(mBase, _c_F_pubsubCommand[6]))
	v213 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v215 = v203 << (uint(int32(2)) % 32)
	v217 = *(*int32)(unsafe.Add(mBase, uint32(v213+v215)))
	v218 = F_kvstoreHashtableFind(m, v211, v208, v217, v11)
	mBase = m.M
	v219 = m.ExcPending
	if v219 != 0 {
		goto L17
	} else {
		goto L56
	}
L56:
	;
	v220 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
	v221 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v223 = *(*int32)(unsafe.Add(mBase, uint32(v221+v215)))
	F_addReplyBulk(m, l0, v223)
	mBase = m.M
	v225 = m.ExcPending
	if v225 != 0 {
		goto L17
	} else {
		goto L57
	}
L57:
	;
	if v220 != 0 {
		goto L59
	} else {
		goto L60
	}
L58:
	;
	F_addReplyLongLong(m, l0, v231)
	mBase = m.M
	v233 = m.ExcPending
	if v233 != 0 {
		goto L17
	} else {
		goto L61
	}
L59:
	;
	v227 = *(*int32)(unsafe.Add(mBase, uint32(v220)+16))
	v228 = *(*int32)(unsafe.Add(mBase, uint32(v220)+12))
	v231 = base.I64_extend_i32_u(v227 + v228)
	goto L58
L60:
	;
	v231 = int64(0)
	goto L58
L61:
	;
	v235 = v203 + int32(1)
	v236 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v235 < v236 {
		v203 = v235
		goto L54
	} else {
		goto L62
	}
L62:
	;
	goto L1
L63:
	;
	v293 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v294 = *(*int32)(unsafe.Add(mBase, uint32(v293)+4))
	v295 = F_objectGetVal(m, v294)
	mBase = m.M
	v296 = int32(_a_F_pubsubCommand_4)
	v299 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v295))))
	if v299 != 0 {
		goto L82
	} else {
		goto L83
	}
L64:
	;
	if v278-v280 != 0 {
		goto L63
	} else {
		goto L76
	}
L65:
	;
	v278 = F_tolower(m, v274)
	mBase = m.M
	v279 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v275))))
	v280 = F_tolower(m, v279)
	mBase = m.M
	goto L64
L66:
	;
	v248 = v242
	v249 = v243
	v250 = v246
	goto L69
L67:
	;
	v274 = int32(0)
	v275 = v243
	goto L65
L68:
	;
	v274 = v271 & int32(255)
	v275 = v270
	goto L65
L69:
	;
	v252 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v249))))
	if v252 == int32(0) {
		v270 = v249
		v271 = v250
		goto L68
	} else {
		goto L71
	}
L70:
	;
	v270 = v264
	v271 = int32(0)
	goto L68
L71:
	;
	v256 = v250 & int32(255)
	if v256 == v252 {
		goto L72
	} else {
		goto L73
	}
L72:
	;
	v263 = int32(1)
	v264 = v249 + v263
	v265 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v248)+1)))
	if v265 != 0 {
		v248 = v248 + v263
		v249 = v264
		v250 = v265
		goto L69
	} else {
		goto L75
	}
L73:
	;
	v258 = F_tolower(m, v256)
	mBase = m.M
	v259 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v249))))
	v260 = F_tolower(m, v259)
	mBase = m.M
	if v258 == v260 {
		goto L72
	} else {
		goto L74
	}
L74:
	;
	v262 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v248))))
	v270 = v249
	v271 = v262
	goto L68
L75:
	;
	goto L70
L76:
	;
	v282 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v282 != int32(2) {
		goto L63
	} else {
		goto L77
	}
L77:
	;
	v286 = *(*int32)(unsafe.Add(mBase, _c_F_pubsubCommand[7]))
	v287 = *(*int32)(unsafe.Add(mBase, uint32(v286)+16))
	v288 = *(*int32)(unsafe.Add(mBase, uint32(v286)+12))
	F_addReplyLongLong(m, l0, base.I64_extend_i32_u(v287+v288))
	mBase = m.M
	v292 = m.ExcPending
	if v292 != 0 {
		goto L17
	} else {
		goto L78
	}
L78:
	;
	goto L1
L79:
	;
	v352 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v353 = *(*int32)(unsafe.Add(mBase, uint32(v352)+4))
	v354 = F_objectGetVal(m, v353)
	mBase = m.M
	v355 = int32(_a_F_pubsubCommand_5)
	v358 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v354))))
	if v358 != 0 {
		goto L100
	} else {
		goto L101
	}
L80:
	;
	if v331-v333 != 0 {
		goto L79
	} else {
		goto L92
	}
L81:
	;
	v331 = F_tolower(m, v327)
	mBase = m.M
	v332 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v328))))
	v333 = F_tolower(m, v332)
	mBase = m.M
	goto L80
L82:
	;
	v301 = v295
	v302 = v296
	v303 = v299
	goto L85
L83:
	;
	v327 = int32(0)
	v328 = v296
	goto L81
L84:
	;
	v327 = v324 & int32(255)
	v328 = v323
	goto L81
L85:
	;
	v305 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v302))))
	if v305 == int32(0) {
		v323 = v302
		v324 = v303
		goto L84
	} else {
		goto L87
	}
L86:
	;
	v323 = v317
	v324 = int32(0)
	goto L84
L87:
	;
	v309 = v303 & int32(255)
	if v309 == v305 {
		goto L88
	} else {
		goto L89
	}
L88:
	;
	v316 = int32(1)
	v317 = v302 + v316
	v318 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v301)+1)))
	if v318 != 0 {
		v301 = v301 + v316
		v302 = v317
		v303 = v318
		goto L85
	} else {
		goto L91
	}
L89:
	;
	v311 = F_tolower(m, v309)
	mBase = m.M
	v312 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v302))))
	v313 = F_tolower(m, v312)
	mBase = m.M
	if v311 == v313 {
		goto L88
	} else {
		goto L90
	}
L90:
	;
	v315 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v301))))
	v323 = v302
	v324 = v315
	goto L84
L91:
	;
	goto L86
L92:
	;
	v335 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v335&int32(-2) != int32(2) {
		goto L79
	} else {
		goto L93
	}
L93:
	;
	if v335 == int32(2) {
		v346 = int32(0)
		goto L94
	} else {
		goto L95
	}
L94:
	;
	v348 = *(*int32)(unsafe.Add(mBase, _c_F_pubsubCommand[8]))
	F_channelList(m, l0, v346, v348)
	mBase = m.M
	v350 = m.ExcPending
	if v350 != 0 {
		goto L17
	} else {
		goto L96
	}
L95:
	;
	v343 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v344 = *(*int32)(unsafe.Add(mBase, uint32(v343)+8))
	v345 = F_objectGetVal(m, v344)
	mBase = m.M
	v346 = v345
	goto L94
L96:
	;
	goto L1
L97:
	;
	F_addReplySubcommandSyntaxError(m, l0)
	mBase = m.M
	v546 = m.ExcPending
	if v546 != 0 {
		goto L17
	} else {
		goto L153
	}
L98:
	;
	if v390-v392 != 0 {
		goto L97
	} else {
		goto L110
	}
L99:
	;
	v390 = F_tolower(m, v386)
	mBase = m.M
	v391 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v387))))
	v392 = F_tolower(m, v391)
	mBase = m.M
	goto L98
L100:
	;
	v360 = v354
	v361 = v355
	v362 = v358
	goto L103
L101:
	;
	v386 = int32(0)
	v387 = v355
	goto L99
L102:
	;
	v386 = v383 & int32(255)
	v387 = v382
	goto L99
L103:
	;
	v364 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v361))))
	if v364 == int32(0) {
		v382 = v361
		v383 = v362
		goto L102
	} else {
		goto L105
	}
L104:
	;
	v382 = v376
	v383 = int32(0)
	goto L102
L105:
	;
	v368 = v362 & int32(255)
	if v368 == v364 {
		goto L106
	} else {
		goto L107
	}
L106:
	;
	v375 = int32(1)
	v376 = v361 + v375
	v377 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v360)+1)))
	if v377 != 0 {
		v360 = v360 + v375
		v361 = v376
		v362 = v377
		goto L103
	} else {
		goto L109
	}
L107:
	;
	v370 = F_tolower(m, v368)
	mBase = m.M
	v371 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v361))))
	v372 = F_tolower(m, v371)
	mBase = m.M
	if v370 == v372 {
		goto L106
	} else {
		goto L108
	}
L108:
	;
	v374 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v360))))
	v382 = v361
	v383 = v374
	goto L102
L109:
	;
	goto L104
L110:
	;
	v394 = int32(2)
	v395 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v395 < v394 {
		goto L97
	} else {
		goto L111
	}
L111:
	;
	F_addReplyArrayLen(m, l0, v395<<(uint(int32(1))%32)+int32(-4))
	mBase = m.M
	v403 = m.ExcPending
	if v403 != 0 {
		goto L17
	} else {
		goto L112
	}
L112:
	;
	v404 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v404 < int32(3) {
		goto L1
	} else {
		goto L113
	}
L113:
	;
	v409 = v394
	goto L114
L114:
	;
	v415 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v417 = v409 << (uint(int32(2)) % 32)
	v419 = *(*int32)(unsafe.Add(mBase, uint32(v415+v417)))
	v420 = F_objectGetVal(m, v419)
	mBase = m.M
	v421 = int32(0)
	v423 = *(*int32)(unsafe.Add(mBase, _c_F_pubsubCommand[9]))
	if v423 == v421 {
		v514 = v421
		goto L116
	} else {
		goto L117
	}
L116:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11))) = int32(0)
	v519 = *(*int32)(unsafe.Add(mBase, _c_F_pubsubCommand[8]))
	v520 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v522 = *(*int32)(unsafe.Add(mBase, uint32(v520+v417)))
	v523 = F_kvstoreHashtableFind(m, v519, v514, v522, v11)
	mBase = m.M
	v524 = m.ExcPending
	if v524 != 0 {
		goto L17
	} else {
		goto L145
	}
L117:
	;
	v429 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v420+int32(-1)))))
	switch v429 & int32(7) {
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
		v446 = int32(0)
		goto L118
	}
L118:
	;
	v447 = int32(0)
	if v446 < int32(1) {
		v467 = v447
		goto L128
	} else {
		goto L129
	}
L119:
	;
	v445 = *(*int32)(unsafe.Add(mBase, uint32(v420+int32(-17))))
	v446 = v445
	goto L118
L120:
	;
	v442 = *(*int32)(unsafe.Add(mBase, uint32(v420+int32(-9))))
	v446 = v442
	goto L118
L121:
	;
	v439 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v420+int32(-5)))))
	v446 = v439
	goto L118
L122:
	;
	v436 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v420+int32(-3)))))
	v446 = v436
	goto L118
L123:
	;
	v446 = int32(base.Ui32(v429) >> (uint(int32(3)) % 32))
	goto L118
L124:
	;
	v514 = v509 & int32(16383)
	goto L116
L125:
	;
	goto L124
L126:
	;
	v478 = v467 + int32(1)
	if v446 <= v478 {
		goto L136
	} else {
		goto L137
	}
L127:
	;
	v476 = F_crc16(m, v420, v446)
	mBase = m.M
	v509 = v476
	goto L125
L128:
	;
	if v467 != v446 {
		goto L126
	} else {
		goto L134
	}
L129:
	;
	v455 = v447
	goto L130
L130:
	;
	v459 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v420+v455))))
	if v459 == int32(123) {
		v467 = v455
		goto L128
	} else {
		goto L132
	}
L132:
	;
	v463 = v455 + int32(1)
	if v463 != v446 {
		v455 = v463
		goto L130
	} else {
		goto L133
	}
L133:
	;
	goto L127
L134:
	;
	goto L127
L135:
	;
	v506 = F_crc16(m, v420+v467+int32(1), v484+(v467^int32(-1)))
	mBase = m.M
	v509 = v506
	goto L125
L136:
	;
	v499 = F_crc16(m, v420, v446)
	mBase = m.M
	v509 = v499
	goto L125
L137:
	;
	v484 = v478
	goto L139
L138:
	;
	if v484 == v446 {
		goto L136
	} else {
		goto L143
	}
L139:
	;
	v486 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v420+v484))))
	if v486 == int32(125) {
		goto L138
	} else {
		goto L141
	}
L141:
	;
	v490 = v484 + int32(1)
	if v490 != v446 {
		v484 = v490
		goto L139
	} else {
		goto L142
	}
L142:
	;
	goto L136
L143:
	;
	if v484 != v478 {
		goto L135
	} else {
		goto L144
	}
L144:
	;
	goto L136
L145:
	;
	v525 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
	v526 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v528 = *(*int32)(unsafe.Add(mBase, uint32(v526+v417)))
	F_addReplyBulk(m, l0, v528)
	mBase = m.M
	v530 = m.ExcPending
	if v530 != 0 {
		goto L17
	} else {
		goto L146
	}
L146:
	;
	if v525 != 0 {
		goto L148
	} else {
		goto L149
	}
L147:
	;
	F_addReplyLongLong(m, l0, v536)
	mBase = m.M
	v538 = m.ExcPending
	if v538 != 0 {
		goto L17
	} else {
		goto L151
	}
L148:
	;
	v532 = *(*int32)(unsafe.Add(mBase, uint32(v525)+20))
	v533 = *(*int32)(unsafe.Add(mBase, uint32(v525)+16))
	goto L150
L149:
	;
	v536 = int64(0)
	goto L147
L150:
	;
	v536 = base.I64_extend_i32_u(v532 + v533)
	goto L147
L151:
	;
	v540 = v409 + int32(1)
	v541 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v540 < v541 {
		v409 = v540
		goto L114
	} else {
		goto L152
	}
L152:
	;
	goto L1
L153:
	;
	goto L1
}
func F_pubsubPublishMessage(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v19 int64
	_ = v19
	var v24 int64
	_ = v24
	var v27 int64
	_ = v27
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v37 int64
	_ = v37
	var v42 int64
	_ = v42
	var v45 int64
	_ = v45
	var v53 int32
	_ = v53
	var v61 int64
	_ = v61
	var v69 int64
	_ = v69
	var v71 int64
	_ = v71
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	v5 = m.G0
	v7 = v5 - int32(64)
	m.G0 = v7
	if l2 == int32(0) {
		v31 = int32(0)
		v32 = *(*int32)(unsafe.Add(mBase, _c_F_pubsubPublishMessage[0]))
		*(*int32)(unsafe.Add(mBase, uint32(v7+int32(56)))) = v32
		v37 = *(*int64)(unsafe.Add(mBase, _c_F_pubsubPublishMessage[1]))
		*(*int64)(unsafe.Add(mBase, uint32(v7+int32(48)))) = v37
		v42 = *(*int64)(unsafe.Add(mBase, _c_F_pubsubPublishMessage[2]))
		*(*int64)(unsafe.Add(mBase, uint32(v7+int32(40)))) = v42
		v45 = *(*int64)(unsafe.Add(mBase, _c_F_pubsubPublishMessage[3]))
		*(*int64)(unsafe.Add(mBase, uint32(v7)+32)) = v45
	} else {
		v13 = int32(0)
		v14 = *(*int32)(unsafe.Add(mBase, _c_F_pubsubPublishMessage[4]))
		*(*int32)(unsafe.Add(mBase, uint32(v7+int32(56)))) = v14
		v19 = *(*int64)(unsafe.Add(mBase, _c_F_pubsubPublishMessage[5]))
		*(*int64)(unsafe.Add(mBase, uint32(v7+int32(48)))) = v19
		v24 = *(*int64)(unsafe.Add(mBase, _c_F_pubsubPublishMessage[6]))
		*(*int64)(unsafe.Add(mBase, uint32(v7+int32(40)))) = v24
		v27 = *(*int64)(unsafe.Add(mBase, _c_F_pubsubPublishMessage[7]))
		*(*int64)(unsafe.Add(mBase, uint32(v7)+32)) = v27
	}
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v7+int32(56))))
	*(*int32)(unsafe.Add(mBase, uint32(v7+int32(24)))) = v53
	v61 = *(*int64)(unsafe.Add(mBase, uint32(v7+int32(48))))
	*(*int64)(unsafe.Add(mBase, uint32(v7+int32(16)))) = v61
	v69 = *(*int64)(unsafe.Add(mBase, uint32(v7+int32(40))))
	*(*int64)(unsafe.Add(mBase, uint32(v7+int32(8)))) = v69
	v71 = *(*int64)(unsafe.Add(mBase, uint32(v7)+32))
	*(*int64)(unsafe.Add(mBase, uint32(v7))) = v71
	v73 = F_pubsubPublishMessageInternal(m, l0, l1, v7)
	mBase = m.M
	v76 = m.ExcPending
	if v76 != 0 {
		return int32(0)
	} else {
		m.G0 = v7 + int32(64)
		return v73
	}
}
func F_pubsubPublishMessageAndPropagateToCluster(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v20 int64
	_ = v20
	var v25 int64
	_ = v25
	var v28 int64
	_ = v28
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v38 int64
	_ = v38
	var v43 int64
	_ = v43
	var v46 int64
	_ = v46
	var v54 int32
	_ = v54
	var v62 int64
	_ = v62
	var v70 int64
	_ = v70
	var v72 int64
	_ = v72
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v83 int32
	_ = v83
	v6 = m.G0
	v8 = v6 - int32(64)
	m.G0 = v8
	if l2 == int32(0) {
		v32 = int32(0)
		v33 = *(*int32)(unsafe.Add(mBase, _c_F_pubsubPublishMessageAndPropagateToCluster[0]))
		*(*int32)(unsafe.Add(mBase, uint32(v8+int32(56)))) = v33
		v38 = *(*int64)(unsafe.Add(mBase, _c_F_pubsubPublishMessageAndPropagateToCluster[1]))
		*(*int64)(unsafe.Add(mBase, uint32(v8+int32(48)))) = v38
		v43 = *(*int64)(unsafe.Add(mBase, _c_F_pubsubPublishMessageAndPropagateToCluster[2]))
		*(*int64)(unsafe.Add(mBase, uint32(v8+int32(40)))) = v43
		v46 = *(*int64)(unsafe.Add(mBase, _c_F_pubsubPublishMessageAndPropagateToCluster[3]))
		*(*int64)(unsafe.Add(mBase, uint32(v8)+32)) = v46
	} else {
		v14 = int32(0)
		v15 = *(*int32)(unsafe.Add(mBase, _c_F_pubsubPublishMessageAndPropagateToCluster[4]))
		*(*int32)(unsafe.Add(mBase, uint32(v8+int32(56)))) = v15
		v20 = *(*int64)(unsafe.Add(mBase, _c_F_pubsubPublishMessageAndPropagateToCluster[5]))
		*(*int64)(unsafe.Add(mBase, uint32(v8+int32(48)))) = v20
		v25 = *(*int64)(unsafe.Add(mBase, _c_F_pubsubPublishMessageAndPropagateToCluster[6]))
		*(*int64)(unsafe.Add(mBase, uint32(v8+int32(40)))) = v25
		v28 = *(*int64)(unsafe.Add(mBase, _c_F_pubsubPublishMessageAndPropagateToCluster[7]))
		*(*int64)(unsafe.Add(mBase, uint32(v8)+32)) = v28
	}
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v8+int32(56))))
	*(*int32)(unsafe.Add(mBase, uint32(v8+int32(24)))) = v54
	v62 = *(*int64)(unsafe.Add(mBase, uint32(v8+int32(48))))
	*(*int64)(unsafe.Add(mBase, uint32(v8+int32(16)))) = v62
	v70 = *(*int64)(unsafe.Add(mBase, uint32(v8+int32(40))))
	*(*int64)(unsafe.Add(mBase, uint32(v8+int32(8)))) = v70
	v72 = *(*int64)(unsafe.Add(mBase, uint32(v8)+32))
	*(*int64)(unsafe.Add(mBase, uint32(v8))) = v72
	v74 = F_pubsubPublishMessageInternal(m, l0, l1, v8)
	mBase = m.M
	v77 = m.ExcPending
	if v77 != 0 {
		return int32(0)
	} else {
		v79 = *(*int32)(unsafe.Add(mBase, _c_F_pubsubPublishMessageAndPropagateToCluster[8]))
		if v79 == int32(0) {
			m.G0 = v8 + int32(64)
			return v74
		} else {
			F_clusterPropagatePublish(m, l0, l1, l2)
			mBase = m.M
			v83 = m.ExcPending
			if v83 != 0 {
				return int32(0)
			} else {
				m.G0 = v8 + int32(64)
				return v74
			}
		}
	}
}
func F_pubsubSubscribeChannel(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
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
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v100 int32
	_ = v100
	var v102 int32
	_ = v102
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v114 int32
	_ = v114
	var v118 int32
	_ = v118
	var v120 int32
	_ = v120
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v126 int32
	_ = v126
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v133 int32
	_ = v133
	var v136 int32
	_ = v136
	var v148 int32
	_ = v148
	v9 = m.G0
	v11 = v9 - int32(48)
	m.G0 = v11
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+100))
	if v13 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v42 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	v43 = m.T0[v42].(func(*base.Module, int32) int32)(m, l0)
	mBase = m.M
	v44 = m.ExcPending
	if v44 != 0 {
		goto L3
	} else {
		goto L10
	}
L2:
	;
	v15 = F_valkey_malloc(m, int32(32))
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	return int32(0)
L4:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+100)) = v15
	v21 = F_hashtableCreate(m, int32(_a_F_pubsubSubscribeChannel_3))
	mBase = m.M
	v22 = m.ExcPending
	if v22 != 0 {
		goto L3
	} else {
		goto L5
	}
L5:
	;
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+100))
	*(*int32)(unsafe.Add(mBase, uint32(v23))) = v21
	v26 = F_hashtableCreate(m, int32(_a_F_pubsubSubscribeChannel_3))
	mBase = m.M
	v27 = m.ExcPending
	if v27 != 0 {
		goto L3
	} else {
		goto L6
	}
L6:
	;
	v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+100))
	*(*int32)(unsafe.Add(mBase, uint32(v28)+4)) = v26
	v31 = F_hashtableCreate(m, int32(_a_F_pubsubSubscribeChannel_3))
	mBase = m.M
	v32 = m.ExcPending
	if v32 != 0 {
		goto L3
	} else {
		goto L7
	}
L7:
	;
	v33 = *(*int32)(unsafe.Add(mBase, uint32(l0)+100))
	*(*int32)(unsafe.Add(mBase, uint32(v33)+24)) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v33)+16)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v33)+8)) = v31
	goto L1
L8:
	;
	F__serverAssert(m, int32(_a_F_pubsubSubscribeChannel_1), int32(_a_F_pubsubSubscribeChannel_2), int32(321))
	mBase = m.M
	v148 = m.ExcPending
	if v148 != 0 {
		goto L3
	} else {
		goto L42
	}
L9:
	;
	v108 = *(*int32)(unsafe.Add(mBase, uint32(l2)+16))
	v109 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	v110 = *(*int32)(unsafe.Add(mBase, uint32(l0)+204))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+204)) = v110 | int32(131072)
	v114 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+224)))
	if v114 != int32(2) {
		goto L32
	} else {
		goto L33
	}
L10:
	;
	v48 = F_hashtableFindPositionForInsert(m, v43, l1, v11+int32(32), int32(0))
	mBase = m.M
	v49 = m.ExcPending
	if v49 != 0 {
		goto L3
	} else {
		goto L11
	}
L11:
	;
	if v48 == int32(0) {
		v104 = l1
		v105 = int32(0)
		goto L9
	} else {
		goto L12
	}
L12:
	;
	v52 = int32(0)
	v54 = *(*int32)(unsafe.Add(mBase, _c_F_pubsubSubscribeChannel[1]))
	if v54 == v52 {
		v63 = v52
		goto L13
	} else {
		goto L14
	}
L13:
	;
	v64 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v64)))
	v70 = F_kvstoreHashtableFindPositionForInsert(m, v65, v63, l1, v11+int32(16), v11+int32(12))
	mBase = m.M
	v71 = m.ExcPending
	if v71 != 0 {
		goto L3
	} else {
		goto L19
	}
L14:
	;
	v57 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	if v57 == int32(0) {
		v63 = v52
		goto L13
	} else {
		goto L15
	}
L15:
	;
	v60 = F_objectGetVal(m, l1)
	mBase = m.M
	v61 = F_getKeySlot(m, v60)
	mBase = m.M
	v62 = m.ExcPending
	if v62 != 0 {
		goto L3
	} else {
		goto L16
	}
L16:
	;
	v63 = v61
	goto L13
L17:
	;
	v91 = F_hashtableAdd(m, v90, l0)
	mBase = m.M
	v92 = m.ExcPending
	if v92 != 0 {
		goto L3
	} else {
		goto L26
	}
L18:
	;
	v77 = F_hashtableCreate(m, int32(_a_F_pubsubSubscribeChannel_0))
	mBase = m.M
	v78 = m.ExcPending
	if v78 != 0 {
		goto L3
	} else {
		goto L22
	}
L19:
	;
	if v70 != 0 {
		goto L18
	} else {
		goto L20
	}
L20:
	;
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v11)+12))
	goto L21
L21:
	;
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v72+int32(44))))
	v89 = v75
	v90 = v72
	goto L17
L22:
	;
	goto L23
L23:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v77+int32(44)))) = l1
	F_incrRefCount(m, l1)
	mBase = m.M
	v83 = m.ExcPending
	if v83 != 0 {
		goto L3
	} else {
		goto L24
	}
L24:
	;
	v84 = *(*int32)(unsafe.Add(mBase, uint32(v64)))
	F_kvstoreHashtableInsertAtPosition(m, v84, v63, v77, v11+int32(16))
	mBase = m.M
	v88 = m.ExcPending
	if v88 != 0 {
		goto L3
	} else {
		goto L25
	}
L25:
	;
	v89 = l1
	v90 = v77
	goto L17
L26:
	;
	if v91 == int32(0) {
		goto L8
	} else {
		goto L27
	}
L27:
	;
	v95 = m.T0[v42].(func(*base.Module, int32) int32)(m, l0)
	mBase = m.M
	v96 = m.ExcPending
	if v96 != 0 {
		goto L3
	} else {
		goto L28
	}
L28:
	;
	F_hashtableInsertAtPosition(m, v95, v89, v11+int32(32))
	mBase = m.M
	v100 = m.ExcPending
	if v100 != 0 {
		goto L3
	} else {
		goto L29
	}
L29:
	;
	F_incrRefCount(m, v89)
	mBase = m.M
	v102 = m.ExcPending
	if v102 != 0 {
		goto L3
	} else {
		goto L30
	}
L30:
	;
	v104 = v89
	v105 = int32(1)
	goto L9
L31:
	;
	v124 = *(*int32)(unsafe.Add(mBase, uint32(v108)))
	F_addReply(m, l0, v124)
	mBase = m.M
	v126 = m.ExcPending
	if v126 != 0 {
		goto L3
	} else {
		goto L36
	}
L32:
	;
	F_addReplyPushLen(m, l0, int32(3))
	mBase = m.M
	v123 = m.ExcPending
	if v123 != 0 {
		goto L3
	} else {
		goto L35
	}
L33:
	;
	v118 = *(*int32)(unsafe.Add(mBase, _c_F_pubsubSubscribeChannel[0]))
	F_addReply(m, l0, v118)
	mBase = m.M
	v120 = m.ExcPending
	if v120 != 0 {
		goto L3
	} else {
		goto L34
	}
L34:
	;
	goto L31
L35:
	;
	goto L31
L36:
	;
	F_addReplyBulk(m, l0, v104)
	mBase = m.M
	v128 = m.ExcPending
	if v128 != 0 {
		goto L3
	} else {
		goto L37
	}
L37:
	;
	v129 = m.T0[v109].(func(*base.Module, int32) int32)(m, l0)
	mBase = m.M
	v130 = m.ExcPending
	if v130 != 0 {
		goto L3
	} else {
		goto L38
	}
L38:
	;
	F_addReplyLongLong(m, l0, base.I64_extend_i32_s(v129))
	mBase = m.M
	v133 = m.ExcPending
	if v133 != 0 {
		goto L3
	} else {
		goto L39
	}
L39:
	;
	if v110&int32(131072) != 0 {
		goto L40
	} else {
		goto L41
	}
L40:
	;
	m.G0 = v11 + int32(48)
	return v105
L41:
	;
	v136 = *(*int32)(unsafe.Add(mBase, uint32(l0)+204))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+204)) = v136 & int32(-131073)
	goto L40
L42:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_pubsubUnsubscribeAllChannelsInternal(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
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
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v61 int64
	_ = v61
	var v63 int32
	_ = v63
	var v67 int64
	_ = v67
	var v69 int32
	_ = v69
	var v73 int32
	_ = v73
	var v75 int64
	_ = v75
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v91 int32
	_ = v91
	var v96 int32
	_ = v96
	var v101 int32
	_ = v101
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v111 int32
	_ = v111
	var v115 int32
	_ = v115
	var v117 int32
	_ = v117
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v123 int32
	_ = v123
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v130 int32
	_ = v130
	var v133 int32
	_ = v133
	v7 = m.G0
	v9 = v7 - int32(80)
	m.G0 = v9
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	v12 = m.T0[v11].(func(*base.Module, int32) int32)(m, l0)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	if l1 == int32(0) {
		goto L21
	} else {
		goto L22
	}
L2:
	;
	v21 = v9 + int32(32)
	v22 = m.T0[v11].(func(*base.Module, int32) int32)(m, l0)
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		goto L3
	} else {
		goto L7
	}
L3:
	;
	return int32(0)
L4:
	;
	v16 = *(*int32)(unsafe.Add(mBase, uint32(v12)+20))
	v17 = *(*int32)(unsafe.Add(mBase, uint32(v12)+16))
	goto L5
L5:
	;
	if v16+v17 != 0 {
		goto L2
	} else {
		goto L6
	}
L6:
	;
	v101 = int32(0)
	goto L1
L7:
	;
	v24 = int32(1)
	v25 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v21)+14)) = uint8(v25)
	*(*int32)(unsafe.Add(mBase, uint32(v21))) = v22
	*(*int32)(unsafe.Add(mBase, uint32(v21)+24)) = v25
	*(*uint8)(unsafe.Add(mBase, uint32(v21)+15)) = uint8(v24)
	*(*int32)(unsafe.Add(mBase, uint32(v21)+8)) = int32(-1)
	if v22 == v25 {
		goto L9
	} else {
		goto L10
	}
L8:
	;
	v42 = int32(0)
	v47 = F_hashtableNext(m, v9+int32(32), v9+int32(28))
	mBase = m.M
	v48 = m.ExcPending
	if v48 != 0 {
		goto L3
	} else {
		goto L13
	}
L9:
	;
	goto L8
L10:
	;
	goto L11
L11:
	;
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v22)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v21)+24)) = v39
	*(*int32)(unsafe.Add(mBase, uint32(v22)+40)) = v21
	goto L9
L12:
	;
	F_hashtableCleanupIterator(m, v9+int32(32))
	mBase = m.M
	v96 = m.ExcPending
	if v96 != 0 {
		goto L3
	} else {
		goto L20
	}
L13:
	;
	if v47 == int32(0) {
		v91 = v42
		goto L12
	} else {
		goto L14
	}
L14:
	;
	v55 = v42
	goto L15
L15:
	;
	v57 = int32(8)
	v61 = *(*int64)(unsafe.Add(mBase, uint32(l2+v57)))
	*(*int64)(unsafe.Add(mBase, uint32(v9+v57))) = v61
	v63 = int32(16)
	v67 = *(*int64)(unsafe.Add(mBase, uint32(l2+v63)))
	*(*int64)(unsafe.Add(mBase, uint32(v9+v63))) = v67
	v69 = int32(24)
	v73 = *(*int32)(unsafe.Add(mBase, uint32(l2+v69)))
	*(*int32)(unsafe.Add(mBase, uint32(v9+v69))) = v73
	v75 = *(*int64)(unsafe.Add(mBase, uint32(l2)))
	*(*int64)(unsafe.Add(mBase, uint32(v9))) = v75
	v77 = *(*int32)(unsafe.Add(mBase, uint32(v9)+28))
	v78 = F_pubsubUnsubscribeChannel(m, l0, v77, l1, v9)
	mBase = m.M
	v79 = m.ExcPending
	if v79 != 0 {
		goto L3
	} else {
		goto L17
	}
L16:
	;
	v91 = v80
	goto L12
L17:
	;
	v80 = v78 + v55
	v85 = F_hashtableNext(m, v9+int32(32), v9+int32(28))
	mBase = m.M
	v86 = m.ExcPending
	if v86 != 0 {
		goto L3
	} else {
		goto L18
	}
L18:
	;
	if v85 != 0 {
		v55 = v80
		goto L15
	} else {
		goto L19
	}
L19:
	;
	goto L16
L20:
	;
	v101 = v91
	goto L1
L21:
	;
	m.G0 = v9 + int32(80)
	return v101
L22:
	;
	if v101 != 0 {
		goto L21
	} else {
		goto L23
	}
L23:
	;
	v105 = *(*int32)(unsafe.Add(mBase, uint32(l2)+20))
	v106 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	v107 = *(*int32)(unsafe.Add(mBase, uint32(l0)+204))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+204)) = v107 | int32(131072)
	v111 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+224)))
	if v111 != int32(2) {
		goto L25
	} else {
		goto L26
	}
L24:
	;
	v121 = *(*int32)(unsafe.Add(mBase, uint32(v105)))
	F_addReply(m, l0, v121)
	mBase = m.M
	v123 = m.ExcPending
	if v123 != 0 {
		goto L3
	} else {
		goto L29
	}
L25:
	;
	F_addReplyPushLen(m, l0, int32(3))
	mBase = m.M
	v120 = m.ExcPending
	if v120 != 0 {
		goto L3
	} else {
		goto L28
	}
L26:
	;
	v115 = *(*int32)(unsafe.Add(mBase, _c_F_pubsubUnsubscribeAllChannelsInternal[0]))
	F_addReply(m, l0, v115)
	mBase = m.M
	v117 = m.ExcPending
	if v117 != 0 {
		goto L3
	} else {
		goto L27
	}
L27:
	;
	goto L24
L28:
	;
	goto L24
L29:
	;
	F_addReplyNull(m, l0)
	mBase = m.M
	v125 = m.ExcPending
	if v125 != 0 {
		goto L3
	} else {
		goto L30
	}
L30:
	;
	v126 = m.T0[v106].(func(*base.Module, int32) int32)(m, l0)
	mBase = m.M
	v127 = m.ExcPending
	if v127 != 0 {
		goto L3
	} else {
		goto L31
	}
L31:
	;
	F_addReplyLongLong(m, l0, base.I64_extend_i32_s(v126))
	mBase = m.M
	v130 = m.ExcPending
	if v130 != 0 {
		goto L3
	} else {
		goto L32
	}
L32:
	;
	if v107&int32(131072) != 0 {
		goto L21
	} else {
		goto L33
	}
L33:
	;
	v133 = *(*int32)(unsafe.Add(mBase, uint32(l0)+204))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+204)) = v133 & int32(-131073)
	goto L21
}
func F_pubsubUnsubscribeChannel(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v17 int32
	_ = v17
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
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v47 int32
	_ = v47
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
	var v66 int32
	_ = v66
	var v70 int32
	_ = v70
	var v74 int32
	_ = v74
	var v78 int32
	_ = v78
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v95 int32
	_ = v95
	var v97 int32
	_ = v97
	var v101 int32
	_ = v101
	var v110 int32
	_ = v110
	var v117 int32
	_ = v117
	var v120 int32
	_ = v120
	var v125 int32
	_ = v125
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v151 int32
	_ = v151
	var v157 int32
	_ = v157
	var v161 int32
	_ = v161
	var v163 int32
	_ = v163
	var v167 int64
	_ = v167
	var v169 int32
	_ = v169
	var v173 int64
	_ = v173
	var v175 int64
	_ = v175
	var v178 int32
	_ = v178
	var v180 int32
	_ = v180
	var v190 int32
	_ = v190
	var v197 int32
	_ = v197
	v10 = m.G0
	v12 = v10 - int32(32)
	m.G0 = v12
	F_incrRefCount(m, l1)
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l3)+4))
	v20 = m.T0[v19].(func(*base.Module, int32) int32)(m, l0)
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		goto L1
	} else {
		goto L6
	}
L3:
	;
	F__serverAssertWithInfo(m, l0, int32(0), int32(_a_F_pubsubUnsubscribeChannel_0), int32(_a_F_pubsubUnsubscribeChannel_1), int32(352))
	mBase = m.M
	v197 = m.ExcPending
	if v197 != 0 {
		goto L1
	} else {
		goto L52
	}
L4:
	;
	F__serverAssertWithInfo(m, l0, int32(0), int32(_a_F_pubsubUnsubscribeChannel_2), int32(_a_F_pubsubUnsubscribeChannel_1), int32(350))
	mBase = m.M
	v190 = m.ExcPending
	if v190 != 0 {
		goto L1
	} else {
		goto L51
	}
L5:
	;
	if l2 == int32(0) {
		goto L47
	} else {
		goto L48
	}
L6:
	;
	v22 = F_hashtableDelete(m, v20, l1)
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		goto L1
	} else {
		goto L7
	}
L7:
	;
	if v22 == int32(0) {
		v151 = int32(0)
		goto L5
	} else {
		goto L8
	}
L8:
	;
	v26 = int32(0)
	v28 = *(*int32)(unsafe.Add(mBase, _c_F_pubsubUnsubscribeChannel[0]))
	if v28 == v26 {
		v125 = v26
		goto L9
	} else {
		goto L10
	}
L9:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+28)) = int32(0)
	v131 = *(*int32)(unsafe.Add(mBase, uint32(l3)+12))
	v132 = *(*int32)(unsafe.Add(mBase, uint32(v131)))
	v135 = F_kvstoreHashtableFind(m, v132, v125, l1, v12+int32(28))
	mBase = m.M
	v136 = m.ExcPending
	if v136 != 0 {
		goto L1
	} else {
		goto L39
	}
L10:
	;
	v31 = int32(0)
	v32 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	if v32 == v31 {
		v125 = v31
		goto L9
	} else {
		goto L11
	}
L11:
	;
	v35 = F_objectGetVal(m, l1)
	mBase = m.M
	v37 = F_objectGetVal(m, l1)
	mBase = m.M
	v40 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v37+int32(-1)))))
	switch v40 & int32(7) {
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
		v57 = int32(0)
		goto L12
	}
L12:
	;
	v58 = int32(0)
	if v57 < int32(1) {
		v78 = v58
		goto L22
	} else {
		goto L23
	}
L13:
	;
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v37+int32(-17))))
	v57 = v56
	goto L12
L14:
	;
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v37+int32(-9))))
	v57 = v53
	goto L12
L15:
	;
	v50 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v37+int32(-5)))))
	v57 = v50
	goto L12
L16:
	;
	v47 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v37+int32(-3)))))
	v57 = v47
	goto L12
L17:
	;
	v57 = int32(base.Ui32(v40) >> (uint(int32(3)) % 32))
	goto L12
L18:
	;
	v125 = v120 & int32(16383)
	goto L9
L19:
	;
	goto L18
L20:
	;
	v89 = v78 + int32(1)
	if v57 <= v89 {
		goto L30
	} else {
		goto L31
	}
L21:
	;
	v87 = F_crc16(m, v35, v57)
	mBase = m.M
	v120 = v87
	goto L19
L22:
	;
	if v78 != v57 {
		goto L20
	} else {
		goto L28
	}
L23:
	;
	v66 = v58
	goto L24
L24:
	;
	v70 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v35+v66))))
	if v70 == int32(123) {
		v78 = v66
		goto L22
	} else {
		goto L26
	}
L26:
	;
	v74 = v66 + int32(1)
	if v74 != v57 {
		v66 = v74
		goto L24
	} else {
		goto L27
	}
L27:
	;
	goto L21
L28:
	;
	goto L21
L29:
	;
	v117 = F_crc16(m, v35+v78+int32(1), v95+(v78^int32(-1)))
	mBase = m.M
	v120 = v117
	goto L19
L30:
	;
	v110 = F_crc16(m, v35, v57)
	mBase = m.M
	v120 = v110
	goto L19
L31:
	;
	v95 = v89
	goto L33
L32:
	;
	if v95 == v57 {
		goto L30
	} else {
		goto L37
	}
L33:
	;
	v97 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v35+v95))))
	if v97 == int32(125) {
		goto L32
	} else {
		goto L35
	}
L35:
	;
	v101 = v95 + int32(1)
	if v101 != v57 {
		v95 = v101
		goto L33
	} else {
		goto L36
	}
L36:
	;
	goto L30
L37:
	;
	if v95 != v89 {
		goto L29
	} else {
		goto L38
	}
L38:
	;
	goto L30
L39:
	;
	v137 = *(*int32)(unsafe.Add(mBase, uint32(v12)+28))
	if v137 == int32(0) {
		goto L4
	} else {
		goto L40
	}
L40:
	;
	v140 = F_hashtableDelete(m, v137, l0)
	mBase = m.M
	v141 = m.ExcPending
	if v141 != 0 {
		goto L1
	} else {
		goto L41
	}
L41:
	;
	if v140 == int32(0) {
		goto L3
	} else {
		goto L42
	}
L42:
	;
	v144 = *(*int32)(unsafe.Add(mBase, uint32(v137)+20))
	v145 = *(*int32)(unsafe.Add(mBase, uint32(v137)+16))
	goto L44
L43:
	;
	v151 = int32(1)
	goto L5
L44:
	;
	if v144+v145 != 0 {
		goto L43
	} else {
		goto L45
	}
L45:
	;
	v147 = *(*int32)(unsafe.Add(mBase, uint32(v131)))
	v148 = F_kvstoreHashtableDelete(m, v147, v125, l1)
	mBase = m.M
	v149 = m.ExcPending
	if v149 != 0 {
		goto L1
	} else {
		goto L46
	}
L46:
	;
	goto L43
L47:
	;
	F_decrRefCount(m, l1)
	mBase = m.M
	v180 = m.ExcPending
	if v180 != 0 {
		goto L1
	} else {
		goto L50
	}
L48:
	;
	v157 = int32(24)
	v161 = *(*int32)(unsafe.Add(mBase, uint32(l3+v157)))
	*(*int32)(unsafe.Add(mBase, uint32(v12+v157))) = v161
	v163 = int32(16)
	v167 = *(*int64)(unsafe.Add(mBase, uint32(l3+v163)))
	*(*int64)(unsafe.Add(mBase, uint32(v12+v163))) = v167
	v169 = int32(8)
	v173 = *(*int64)(unsafe.Add(mBase, uint32(l3+v169)))
	*(*int64)(unsafe.Add(mBase, uint32(v12+v169))) = v173
	v175 = *(*int64)(unsafe.Add(mBase, uint32(l3)))
	*(*int64)(unsafe.Add(mBase, uint32(v12))) = v175
	F_addReplyPubsubUnsubscribed(m, l0, l1, v12)
	mBase = m.M
	v178 = m.ExcPending
	if v178 != 0 {
		goto L1
	} else {
		goto L49
	}
L49:
	;
	goto L47
L50:
	;
	m.G0 = v12 + int32(32)
	return v151
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
}
func F_pubsubUnsubscribePattern(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
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
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
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
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
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
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v73 int32
	_ = v73
	var v80 int32
	_ = v80
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+100))
	if v6 != 0 {
		F_incrRefCount(m, l1)
		mBase = m.M
		v35 = m.ExcPending
		if v35 != 0 {
			return int32(0)
		} else {
			v36 = *(*int32)(unsafe.Add(mBase, uint32(l0)+100))
			v37 = *(*int32)(unsafe.Add(mBase, uint32(v36)+4))
			v38 = F_hashtableDelete(m, v37, l1)
			mBase = m.M
			v39 = m.ExcPending
			if v39 != 0 {
				return int32(0)
			} else {
				if v38 == int32(0) {
					if l2 == int32(0) {
						F_decrRefCount(m, l1)
						mBase = m.M
						v66 = m.ExcPending
						if v66 != 0 {
							return int32(0)
						} else {
							return v38
						}
					} else {
						F_addReplyPubsubPatUnsubscribed(m, l0, l1)
						mBase = m.M
						v64 = m.ExcPending
						if v64 != 0 {
							return int32(0)
						} else {
							F_decrRefCount(m, l1)
							mBase = m.M
							v66 = m.ExcPending
							if v66 != 0 {
								return int32(0)
							} else {
								return v38
							}
						}
					}
				} else {
					v43 = *(*int32)(unsafe.Add(mBase, _c_F_pubsubUnsubscribePattern[0]))
					v44 = F_dictFind(m, v43, l1)
					mBase = m.M
					v45 = m.ExcPending
					if v45 != 0 {
						return int32(0)
					} else {
						if v44 == int32(0) {
							F__serverAssertWithInfo(m, l0, int32(0), int32(_a_F_pubsubUnsubscribePattern_0), int32(_a_F_pubsubUnsubscribePattern_1), int32(432))
							mBase = m.M
							v73 = m.ExcPending
							if v73 != 0 {
								return int32(0)
							} else {
								F_abort(m)
								mBase = m.M
								base.Wasm_trap_unreachable()
								for {
								}
							}
						} else {
							v48 = *(*int32)(unsafe.Add(mBase, uint32(v44)+8))
							v49 = F_hashtableDelete(m, v48, l0)
							mBase = m.M
							v50 = m.ExcPending
							if v50 != 0 {
								return int32(0)
							} else {
								if v49 == int32(0) {
									F__serverAssertWithInfo(m, l0, int32(0), int32(_a_F_pubsubUnsubscribePattern_2), int32(_a_F_pubsubUnsubscribePattern_1), int32(434))
									mBase = m.M
									v80 = m.ExcPending
									if v80 != 0 {
										return int32(0)
									} else {
										F_abort(m)
										mBase = m.M
										base.Wasm_trap_unreachable()
										for {
										}
									}
								} else {
									v53 = *(*int32)(unsafe.Add(mBase, uint32(v48)+20))
									v54 = *(*int32)(unsafe.Add(mBase, uint32(v48)+16))
									if v53+v54 != 0 {
										if l2 == int32(0) {
											F_decrRefCount(m, l1)
											mBase = m.M
											v66 = m.ExcPending
											if v66 != 0 {
												return int32(0)
											} else {
												return v38
											}
										} else {
											F_addReplyPubsubPatUnsubscribed(m, l0, l1)
											mBase = m.M
											v64 = m.ExcPending
											if v64 != 0 {
												return int32(0)
											} else {
												F_decrRefCount(m, l1)
												mBase = m.M
												v66 = m.ExcPending
												if v66 != 0 {
													return int32(0)
												} else {
													return v38
												}
											}
										}
									} else {
										v57 = *(*int32)(unsafe.Add(mBase, _c_F_pubsubUnsubscribePattern[0]))
										v58 = F_dictDelete(m, v57, l1)
										mBase = m.M
										v59 = m.ExcPending
										if v59 != 0 {
											return int32(0)
										} else {
											if l2 == int32(0) {
												F_decrRefCount(m, l1)
												mBase = m.M
												v66 = m.ExcPending
												if v66 != 0 {
													return int32(0)
												} else {
													return v38
												}
											} else {
												F_addReplyPubsubPatUnsubscribed(m, l0, l1)
												mBase = m.M
												v64 = m.ExcPending
												if v64 != 0 {
													return int32(0)
												} else {
													F_decrRefCount(m, l1)
													mBase = m.M
													v66 = m.ExcPending
													if v66 != 0 {
														return int32(0)
													} else {
														return v38
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
	} else {
		v8 = F_valkey_malloc(m, int32(32))
		mBase = m.M
		v11 = m.ExcPending
		if v11 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(l0)+100)) = v8
			v14 = F_hashtableCreate(m, int32(_a_F_pubsubUnsubscribePattern_3))
			mBase = m.M
			v15 = m.ExcPending
			if v15 != 0 {
				return int32(0)
			} else {
				v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+100))
				*(*int32)(unsafe.Add(mBase, uint32(v16))) = v14
				v19 = F_hashtableCreate(m, int32(_a_F_pubsubUnsubscribePattern_3))
				mBase = m.M
				v20 = m.ExcPending
				if v20 != 0 {
					return int32(0)
				} else {
					v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+100))
					*(*int32)(unsafe.Add(mBase, uint32(v21)+4)) = v19
					v24 = F_hashtableCreate(m, int32(_a_F_pubsubUnsubscribePattern_3))
					mBase = m.M
					v25 = m.ExcPending
					if v25 != 0 {
						return int32(0)
					} else {
						v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+100))
						*(*int32)(unsafe.Add(mBase, uint32(v26)+24)) = int32(0)
						*(*int64)(unsafe.Add(mBase, uint32(v26)+16)) = int64(0)
						*(*int32)(unsafe.Add(mBase, uint32(v26)+8)) = v24
						F_incrRefCount(m, l1)
						mBase = m.M
						v35 = m.ExcPending
						if v35 != 0 {
							return int32(0)
						} else {
							v36 = *(*int32)(unsafe.Add(mBase, uint32(l0)+100))
							v37 = *(*int32)(unsafe.Add(mBase, uint32(v36)+4))
							v38 = F_hashtableDelete(m, v37, l1)
							mBase = m.M
							v39 = m.ExcPending
							if v39 != 0 {
								return int32(0)
							} else {
								if v38 == int32(0) {
									if l2 == int32(0) {
										F_decrRefCount(m, l1)
										mBase = m.M
										v66 = m.ExcPending
										if v66 != 0 {
											return int32(0)
										} else {
											return v38
										}
									} else {
										F_addReplyPubsubPatUnsubscribed(m, l0, l1)
										mBase = m.M
										v64 = m.ExcPending
										if v64 != 0 {
											return int32(0)
										} else {
											F_decrRefCount(m, l1)
											mBase = m.M
											v66 = m.ExcPending
											if v66 != 0 {
												return int32(0)
											} else {
												return v38
											}
										}
									}
								} else {
									v43 = *(*int32)(unsafe.Add(mBase, _c_F_pubsubUnsubscribePattern[0]))
									v44 = F_dictFind(m, v43, l1)
									mBase = m.M
									v45 = m.ExcPending
									if v45 != 0 {
										return int32(0)
									} else {
										if v44 == int32(0) {
											F__serverAssertWithInfo(m, l0, int32(0), int32(_a_F_pubsubUnsubscribePattern_0), int32(_a_F_pubsubUnsubscribePattern_1), int32(432))
											mBase = m.M
											v73 = m.ExcPending
											if v73 != 0 {
												return int32(0)
											} else {
												F_abort(m)
												mBase = m.M
												base.Wasm_trap_unreachable()
												for {
												}
											}
										} else {
											v48 = *(*int32)(unsafe.Add(mBase, uint32(v44)+8))
											v49 = F_hashtableDelete(m, v48, l0)
											mBase = m.M
											v50 = m.ExcPending
											if v50 != 0 {
												return int32(0)
											} else {
												if v49 == int32(0) {
													F__serverAssertWithInfo(m, l0, int32(0), int32(_a_F_pubsubUnsubscribePattern_2), int32(_a_F_pubsubUnsubscribePattern_1), int32(434))
													mBase = m.M
													v80 = m.ExcPending
													if v80 != 0 {
														return int32(0)
													} else {
														F_abort(m)
														mBase = m.M
														base.Wasm_trap_unreachable()
														for {
														}
													}
												} else {
													v53 = *(*int32)(unsafe.Add(mBase, uint32(v48)+20))
													v54 = *(*int32)(unsafe.Add(mBase, uint32(v48)+16))
													if v53+v54 != 0 {
														if l2 == int32(0) {
															F_decrRefCount(m, l1)
															mBase = m.M
															v66 = m.ExcPending
															if v66 != 0 {
																return int32(0)
															} else {
																return v38
															}
														} else {
															F_addReplyPubsubPatUnsubscribed(m, l0, l1)
															mBase = m.M
															v64 = m.ExcPending
															if v64 != 0 {
																return int32(0)
															} else {
																F_decrRefCount(m, l1)
																mBase = m.M
																v66 = m.ExcPending
																if v66 != 0 {
																	return int32(0)
																} else {
																	return v38
																}
															}
														}
													} else {
														v57 = *(*int32)(unsafe.Add(mBase, _c_F_pubsubUnsubscribePattern[0]))
														v58 = F_dictDelete(m, v57, l1)
														mBase = m.M
														v59 = m.ExcPending
														if v59 != 0 {
															return int32(0)
														} else {
															if l2 == int32(0) {
																F_decrRefCount(m, l1)
																mBase = m.M
																v66 = m.ExcPending
																if v66 != 0 {
																	return int32(0)
																} else {
																	return v38
																}
															} else {
																F_addReplyPubsubPatUnsubscribed(m, l0, l1)
																mBase = m.M
																v64 = m.ExcPending
																if v64 != 0 {
																	return int32(0)
																} else {
																	F_decrRefCount(m, l1)
																	mBase = m.M
																	v66 = m.ExcPending
																	if v66 != 0 {
																		return int32(0)
																	} else {
																		return v38
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
				}
			}
		}
	}
}
