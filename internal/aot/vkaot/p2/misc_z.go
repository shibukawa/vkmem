package p2

import (
	base "github.com/shibukawa/vkmem/internal/aot/vkaot/base"
	"unsafe"
)

func F__ziplistEntryConvertAndValidate(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int64
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v33 int32
	_ = v33
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v14 = F_ziplistGet(m, l0, v8+int32(12), v8+int32(8), v8)
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		return int32(0)
	} else {
		if v14 != 0 {
			v19 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
			v20 = *(*int32)(unsafe.Add(mBase, uint32(v8)+12))
			if v20 == int32(0) {
				v26 = *(*int64)(unsafe.Add(mBase, uint32(v8)))
				v27 = F_lpAppendInteger(m, v19, v26)
				mBase = m.M
				v28 = m.ExcPending
				if v28 != 0 {
					return int32(0)
				} else {
					v29 = v27
					*(*int32)(unsafe.Add(mBase, uint32(l2))) = v29
					v33 = int32(1)
					m.G0 = v8 + int32(16)
					return v33
				}
			} else {
				v23 = *(*int32)(unsafe.Add(mBase, uint32(v8)+8))
				v24 = F_lpAppend(m, v19, v20, v23)
				mBase = m.M
				v25 = m.ExcPending
				if v25 != 0 {
					return int32(0)
				} else {
					v29 = v24
					*(*int32)(unsafe.Add(mBase, uint32(l2))) = v29
					v33 = int32(1)
					m.G0 = v8 + int32(16)
					return v33
				}
			}
		} else {
			v33 = int32(0)
			m.G0 = v8 + int32(16)
			return v33
		}
	}
}
func F_zaddGenericCommand(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v92 int32
	_ = v92
	var v95 int32
	_ = v95
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
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
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v133 int32
	_ = v133
	var v136 int32
	_ = v136
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
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v173 int32
	_ = v173
	var v176 int32
	_ = v176
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v182 int32
	_ = v182
	var v186 int32
	_ = v186
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v214 int32
	_ = v214
	var v217 int32
	_ = v217
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v223 int32
	_ = v223
	var v227 int32
	_ = v227
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v251 int32
	_ = v251
	var v255 int32
	_ = v255
	var v258 int32
	_ = v258
	var v260 int32
	_ = v260
	var v261 int32
	_ = v261
	var v262 int32
	_ = v262
	var v264 int32
	_ = v264
	var v268 int32
	_ = v268
	var v270 int32
	_ = v270
	var v271 int32
	_ = v271
	var v272 int32
	_ = v272
	var v274 int32
	_ = v274
	var v275 int32
	_ = v275
	var v276 int32
	_ = v276
	var v277 int32
	_ = v277
	var v282 int32
	_ = v282
	var v283 int32
	_ = v283
	var v286 int32
	_ = v286
	var v287 int32
	_ = v287
	var v290 int32
	_ = v290
	var v291 int32
	_ = v291
	var v292 int32
	_ = v292
	var v296 int32
	_ = v296
	var v299 int32
	_ = v299
	var v300 int32
	_ = v300
	var v302 int32
	_ = v302
	var v303 int32
	_ = v303
	var v305 int32
	_ = v305
	var v306 int32
	_ = v306
	var v307 int32
	_ = v307
	var v308 int32
	_ = v308
	var v312 int32
	_ = v312
	var v315 int32
	_ = v315
	var v316 int32
	_ = v316
	var v317 int32
	_ = v317
	var v328 int32
	_ = v328
	var v335 int32
	_ = v335
	var v337 int32
	_ = v337
	var v338 int32
	_ = v338
	var v344 int32
	_ = v344
	var v345 int32
	_ = v345
	var v349 int32
	_ = v349
	var v353 int32
	_ = v353
	var v359 int32
	_ = v359
	var v360 int32
	_ = v360
	var v361 int32
	_ = v361
	var v363 int32
	_ = v363
	var v370 int32
	_ = v370
	var v374 int32
	_ = v374
	var v375 int32
	_ = v375
	var v380 int32
	_ = v380
	var v387 int32
	_ = v387
	var v393 int32
	_ = v393
	var v398 int32
	_ = v398
	var v401 int32
	_ = v401
	var v403 int32
	_ = v403
	var v406 int32
	_ = v406
	var v407 int32
	_ = v407
	var v409 int32
	_ = v409
	var v418 int32
	_ = v418
	var v419 int32
	_ = v419
	var v422 int32
	_ = v422
	var v429 int32
	_ = v429
	var v432 int32
	_ = v432
	var v435 int32
	_ = v435
	var v438 int32
	_ = v438
	var v439 int32
	_ = v439
	var v441 int32
	_ = v441
	var v443 int32
	_ = v443
	var v456 int32
	_ = v456
	var v461 int32
	_ = v461
	var v462 int32
	_ = v462
	var v463 int32
	_ = v463
	var v466 int32
	_ = v466
	var v467 int32
	_ = v467
	var v470 int32
	_ = v470
	var v471 int32
	_ = v471
	var v473 int32
	_ = v473
	var v477 int32
	_ = v477
	var v479 int32
	_ = v479
	var v484 int32
	_ = v484
	var v489 int32
	_ = v489
	var v497 int32
	_ = v497
	var v500 int32
	_ = v500
	var v504 int32
	_ = v504
	var v505 int32
	_ = v505
	var v510 int32
	_ = v510
	var v512 float64
	_ = v512
	var v515 int32
	_ = v515
	var v520 int32
	_ = v520
	var v521 int32
	_ = v521
	var v522 int32
	_ = v522
	var v527 int32
	_ = v527
	var v528 int32
	_ = v528
	var v531 int32
	_ = v531
	var v534 int32
	_ = v534
	var v536 int32
	_ = v536
	var v541 int32
	_ = v541
	var v546 int32
	_ = v546
	var v548 int32
	_ = v548
	var v550 float64
	_ = v550
	var v558 int32
	_ = v558
	var v562 int32
	_ = v562
	var v563 int32
	_ = v563
	var v566 float64
	_ = v566
	var v567 int32
	_ = v567
	var v569 int64
	_ = v569
	var v574 int32
	_ = v574
	var v586 int32
	_ = v586
	var v590 int32
	_ = v590
	var v591 int32
	_ = v591
	var v593 int32
	_ = v593
	var v594 float64
	_ = v594
	var v595 int32
	_ = v595
	var v599 int32
	_ = v599
	var v601 int32
	_ = v601
	var v605 int32
	_ = v605
	var v606 int32
	_ = v606
	var v607 int32
	_ = v607
	var v609 int32
	_ = v609
	var v610 int32
	_ = v610
	var v615 int32
	_ = v615
	var v620 int32
	_ = v620
	var v623 int32
	_ = v623
	var v633 int32
	_ = v633
	var v637 int32
	_ = v637
	var v641 int32
	_ = v641
	var v659 int32
	_ = v659
	var v677 int32
	_ = v677
	v17 = m.G0
	v19 = v17 - int32(32)
	m.G0 = v19
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v21)+4))
	v23 = int32(2)
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if int32(3) <= v24 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	if v316 == v315 {
		goto L95
	} else {
		goto L96
	}
L2:
	;
	v30 = l1
	v33 = v23
	v36 = int32(0)
	goto L5
L3:
	;
	v312 = l1
	v315 = v23
	v316 = v24
	v317 = int32(1)
	goto L1
L4:
	;
	v312 = v305
	v315 = v306
	v316 = v307
	v317 = base.B2i32(v308 == int32(0))
	goto L1
L5:
	;
	v45 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v45+v33<<(uint(int32(2))%32))))
	v50 = F_objectGetVal(m, v49)
	mBase = m.M
	v51 = int32(_a786)
	v54 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v50))))
	if v54 != 0 {
		goto L11
	} else {
		goto L12
	}
L6:
	;
	v305 = v299
	v306 = v302
	v307 = v303
	v308 = v300
	goto L4
L7:
	;
	v302 = v33 + int32(1)
	v303 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v302 < v303 {
		v30 = v299
		v33 = v302
		v36 = v300
		goto L5
	} else {
		goto L92
	}
L8:
	;
	v92 = int32(_a787)
	v95 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v50))))
	if v95 != 0 {
		goto L25
	} else {
		goto L26
	}
L9:
	;
	if v86-v88 != 0 {
		goto L8
	} else {
		goto L21
	}
L10:
	;
	v86 = F_tolower(m, v82)
	mBase = m.M
	v87 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v83))))
	v88 = F_tolower(m, v87)
	mBase = m.M
	goto L9
L11:
	;
	v56 = v50
	v57 = v51
	v58 = v54
	goto L14
L12:
	;
	v82 = int32(0)
	v83 = v51
	goto L10
L13:
	;
	v82 = v79 & int32(255)
	v83 = v78
	goto L10
L14:
	;
	v60 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v57))))
	if v60 == int32(0) {
		v78 = v57
		v79 = v58
		goto L13
	} else {
		goto L16
	}
L15:
	;
	v78 = v72
	v79 = int32(0)
	goto L13
L16:
	;
	v64 = v58 & int32(255)
	if v64 == v60 {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	v71 = int32(1)
	v72 = v57 + v71
	v73 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v56)+1)))
	if v73 != 0 {
		v56 = v56 + v71
		v57 = v72
		v58 = v73
		goto L14
	} else {
		goto L20
	}
L18:
	;
	v66 = F_tolower(m, v64)
	mBase = m.M
	v67 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v57))))
	v68 = F_tolower(m, v67)
	mBase = m.M
	if v66 == v68 {
		goto L17
	} else {
		goto L19
	}
L19:
	;
	v70 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v56))))
	v78 = v57
	v79 = v70
	goto L13
L20:
	;
	goto L15
L21:
	;
	v299 = v30 | int32(2)
	v300 = v36
	goto L7
L22:
	;
	v133 = int32(_a1726)
	v136 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v50))))
	if v136 != 0 {
		goto L39
	} else {
		goto L40
	}
L23:
	;
	if v127-v129 != 0 {
		goto L22
	} else {
		goto L35
	}
L24:
	;
	v127 = F_tolower(m, v123)
	mBase = m.M
	v128 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v124))))
	v129 = F_tolower(m, v128)
	mBase = m.M
	goto L23
L25:
	;
	v97 = v50
	v98 = v92
	v99 = v95
	goto L28
L26:
	;
	v123 = int32(0)
	v124 = v92
	goto L24
L27:
	;
	v123 = v120 & int32(255)
	v124 = v119
	goto L24
L28:
	;
	v101 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v98))))
	if v101 == int32(0) {
		v119 = v98
		v120 = v99
		goto L27
	} else {
		goto L30
	}
L29:
	;
	v119 = v113
	v120 = int32(0)
	goto L27
L30:
	;
	v105 = v99 & int32(255)
	if v105 == v101 {
		goto L31
	} else {
		goto L32
	}
L31:
	;
	v112 = int32(1)
	v113 = v98 + v112
	v114 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v97)+1)))
	if v114 != 0 {
		v97 = v97 + v112
		v98 = v113
		v99 = v114
		goto L28
	} else {
		goto L34
	}
L32:
	;
	v107 = F_tolower(m, v105)
	mBase = m.M
	v108 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v98))))
	v109 = F_tolower(m, v108)
	mBase = m.M
	if v107 == v109 {
		goto L31
	} else {
		goto L33
	}
L33:
	;
	v111 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v97))))
	v119 = v98
	v120 = v111
	goto L27
L34:
	;
	goto L29
L35:
	;
	v299 = v30 | int32(4)
	v300 = v36
	goto L7
L36:
	;
	v173 = int32(_a1727)
	v176 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v50))))
	if v176 != 0 {
		goto L53
	} else {
		goto L54
	}
L37:
	;
	if v168-v170 != 0 {
		goto L36
	} else {
		goto L49
	}
L38:
	;
	v168 = F_tolower(m, v164)
	mBase = m.M
	v169 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v165))))
	v170 = F_tolower(m, v169)
	mBase = m.M
	goto L37
L39:
	;
	v138 = v50
	v139 = v133
	v140 = v136
	goto L42
L40:
	;
	v164 = int32(0)
	v165 = v133
	goto L38
L41:
	;
	v164 = v161 & int32(255)
	v165 = v160
	goto L38
L42:
	;
	v142 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v139))))
	if v142 == int32(0) {
		v160 = v139
		v161 = v140
		goto L41
	} else {
		goto L44
	}
L43:
	;
	v160 = v154
	v161 = int32(0)
	goto L41
L44:
	;
	v146 = v140 & int32(255)
	if v146 == v142 {
		goto L45
	} else {
		goto L46
	}
L45:
	;
	v153 = int32(1)
	v154 = v139 + v153
	v155 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v138)+1)))
	if v155 != 0 {
		v138 = v138 + v153
		v139 = v154
		v140 = v155
		goto L42
	} else {
		goto L48
	}
L46:
	;
	v148 = F_tolower(m, v146)
	mBase = m.M
	v149 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v139))))
	v150 = F_tolower(m, v149)
	mBase = m.M
	if v148 == v150 {
		goto L45
	} else {
		goto L47
	}
L47:
	;
	v152 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v138))))
	v160 = v139
	v161 = v152
	goto L41
L48:
	;
	goto L43
L49:
	;
	v299 = v30
	v300 = int32(1)
	goto L7
L50:
	;
	v214 = int32(_a788)
	v217 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v50))))
	if v217 != 0 {
		goto L67
	} else {
		goto L68
	}
L51:
	;
	if v208-v210 != 0 {
		goto L50
	} else {
		goto L63
	}
L52:
	;
	v208 = F_tolower(m, v204)
	mBase = m.M
	v209 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v205))))
	v210 = F_tolower(m, v209)
	mBase = m.M
	goto L51
L53:
	;
	v178 = v50
	v179 = v173
	v180 = v176
	goto L56
L54:
	;
	v204 = int32(0)
	v205 = v173
	goto L52
L55:
	;
	v204 = v201 & int32(255)
	v205 = v200
	goto L52
L56:
	;
	v182 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v179))))
	if v182 == int32(0) {
		v200 = v179
		v201 = v180
		goto L55
	} else {
		goto L58
	}
L57:
	;
	v200 = v194
	v201 = int32(0)
	goto L55
L58:
	;
	v186 = v180 & int32(255)
	if v186 == v182 {
		goto L59
	} else {
		goto L60
	}
L59:
	;
	v193 = int32(1)
	v194 = v179 + v193
	v195 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v178)+1)))
	if v195 != 0 {
		v178 = v178 + v193
		v179 = v194
		v180 = v195
		goto L56
	} else {
		goto L62
	}
L60:
	;
	v188 = F_tolower(m, v186)
	mBase = m.M
	v189 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v179))))
	v190 = F_tolower(m, v189)
	mBase = m.M
	if v188 == v190 {
		goto L59
	} else {
		goto L61
	}
L61:
	;
	v192 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v178))))
	v200 = v179
	v201 = v192
	goto L55
L62:
	;
	goto L57
L63:
	;
	v299 = v30 | int32(1)
	v300 = v36
	goto L7
L64:
	;
	v255 = int32(_a789)
	v258 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v50))))
	if v258 != 0 {
		goto L81
	} else {
		goto L82
	}
L65:
	;
	if v249-v251 != 0 {
		goto L64
	} else {
		goto L77
	}
L66:
	;
	v249 = F_tolower(m, v245)
	mBase = m.M
	v250 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v246))))
	v251 = F_tolower(m, v250)
	mBase = m.M
	goto L65
L67:
	;
	v219 = v50
	v220 = v214
	v221 = v217
	goto L70
L68:
	;
	v245 = int32(0)
	v246 = v214
	goto L66
L69:
	;
	v245 = v242 & int32(255)
	v246 = v241
	goto L66
L70:
	;
	v223 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v220))))
	if v223 == int32(0) {
		v241 = v220
		v242 = v221
		goto L69
	} else {
		goto L72
	}
L71:
	;
	v241 = v235
	v242 = int32(0)
	goto L69
L72:
	;
	v227 = v221 & int32(255)
	if v227 == v223 {
		goto L73
	} else {
		goto L74
	}
L73:
	;
	v234 = int32(1)
	v235 = v220 + v234
	v236 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v219)+1)))
	if v236 != 0 {
		v219 = v219 + v234
		v220 = v235
		v221 = v236
		goto L70
	} else {
		goto L76
	}
L74:
	;
	v229 = F_tolower(m, v227)
	mBase = m.M
	v230 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v220))))
	v231 = F_tolower(m, v230)
	mBase = m.M
	if v229 == v231 {
		goto L73
	} else {
		goto L75
	}
L75:
	;
	v233 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v219))))
	v241 = v220
	v242 = v233
	goto L69
L76:
	;
	goto L71
L77:
	;
	v299 = v30 | int32(8)
	v300 = v36
	goto L7
L78:
	;
	v299 = v30 | int32(16)
	v300 = v36
	goto L7
L79:
	;
	if v290-v292 == int32(0) {
		goto L78
	} else {
		goto L91
	}
L80:
	;
	v290 = F_tolower(m, v286)
	mBase = m.M
	v291 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v287))))
	v292 = F_tolower(m, v291)
	mBase = m.M
	goto L79
L81:
	;
	v260 = v50
	v261 = v255
	v262 = v258
	goto L84
L82:
	;
	v286 = int32(0)
	v287 = v255
	goto L80
L83:
	;
	v286 = v283 & int32(255)
	v287 = v282
	goto L80
L84:
	;
	v264 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v261))))
	if v264 == int32(0) {
		v282 = v261
		v283 = v262
		goto L83
	} else {
		goto L86
	}
L85:
	;
	v282 = v276
	v283 = int32(0)
	goto L83
L86:
	;
	v268 = v262 & int32(255)
	if v268 == v264 {
		goto L87
	} else {
		goto L88
	}
L87:
	;
	v275 = int32(1)
	v276 = v261 + v275
	v277 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v260)+1)))
	if v277 != 0 {
		v260 = v260 + v275
		v261 = v276
		v262 = v277
		goto L84
	} else {
		goto L90
	}
L88:
	;
	v270 = F_tolower(m, v268)
	mBase = m.M
	v271 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v261))))
	v272 = F_tolower(m, v271)
	mBase = m.M
	if v270 == v272 {
		goto L87
	} else {
		goto L89
	}
L89:
	;
	v274 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v260))))
	v282 = v261
	v283 = v274
	goto L83
L90:
	;
	goto L85
L91:
	;
	v296 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v305 = v30
	v306 = v33
	v307 = v296
	v308 = v36
	goto L4
L92:
	;
	goto L6
L93:
	;
	m.G0 = v19 + int32(32)
	return
L94:
	;
	v338 = int32(6)
	if v312&v338 != v338 {
		goto L100
	} else {
		goto L101
	}
L95:
	;
	v335 = *(*int32)(unsafe.Add(mBase, _consts[86]))
	F_addReplyErrorObject(m, l0, v335)
	mBase = m.M
	v337 = m.ExcPending
	if v337 != 0 {
		goto L98
	} else {
		goto L99
	}
L96:
	;
	v328 = v316 - v315
	if v328&int32(1) == int32(0) {
		goto L94
	} else {
		goto L97
	}
L97:
	;
	goto L95
L98:
	;
	return
L99:
	;
	goto L93
L100:
	;
	v345 = int32(10)
	if v312&v345 == v345 {
		goto L104
	} else {
		goto L105
	}
L101:
	;
	F_addReplyError(m, l0, int32(_a1728))
	mBase = m.M
	v344 = m.ExcPending
	if v344 != 0 {
		goto L98
	} else {
		goto L102
	}
L102:
	;
	goto L93
L103:
	;
	v360 = int32(1)
	v361 = v328 >> (uint(v360) % 32)
	v363 = v312 & v360
	if v363 == int32(0) {
		goto L109
	} else {
		goto L110
	}
L104:
	;
	F_addReplyError(m, l0, int32(_a1729))
	mBase = m.M
	v359 = m.ExcPending
	if v359 != 0 {
		goto L98
	} else {
		goto L108
	}
L105:
	;
	v349 = int32(18)
	if v312&v349 == v349 {
		goto L104
	} else {
		goto L106
	}
L106:
	;
	v353 = int32(24)
	if v312&v353 != v353 {
		goto L103
	} else {
		goto L107
	}
L107:
	;
	goto L104
L108:
	;
	goto L93
L109:
	;
	v374 = F_valkey_malloc(m, v328<<(uint(int32(2))%32))
	mBase = m.M
	v375 = m.ExcPending
	if v375 != 0 {
		goto L98
	} else {
		goto L113
	}
L110:
	;
	if v361 < int32(2) {
		goto L109
	} else {
		goto L111
	}
L111:
	;
	F_addReplyError(m, l0, int32(_a1730))
	mBase = m.M
	v370 = m.ExcPending
	if v370 != 0 {
		goto L98
	} else {
		goto L112
	}
L112:
	;
	goto L93
L113:
	;
	if int32(1) <= v361 {
		goto L116
	} else {
		goto L117
	}
L114:
	;
	F_valkey_free(m, v374)
	mBase = m.M
	v677 = m.ExcPending
	if v677 != 0 {
		goto L98
	} else {
		goto L175
	}
L115:
	;
	v461 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v462 = F_lookupKeyWrite(m, v461, v22)
	mBase = m.M
	v463 = m.ExcPending
	if v463 != 0 {
		goto L98
	} else {
		goto L132
	}
L116:
	;
	v380 = v315 << (uint(int32(2)) % 32)
	v387 = int32(0)
	v393 = int32(0)
	goto L118
L117:
	;
	v456 = int32(0)
	goto L115
L118:
	;
	v398 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v401 = v387 << (uint(int32(3)) % 32)
	v403 = *(*int32)(unsafe.Add(mBase, uint32(v398+v380+v401)))
	v406 = F_getDoubleFromObjectOrReply(m, l0, v403, v374+v401, int32(0))
	mBase = m.M
	v407 = m.ExcPending
	if v407 != 0 {
		goto L98
	} else {
		goto L120
	}
L119:
	;
	v456 = v441
	goto L115
L120:
	;
	if v406 != 0 {
		goto L114
	} else {
		goto L121
	}
L121:
	;
	v409 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v418 = *(*int32)(unsafe.Add(mBase, uint32(v409+v380+v387<<(uint(int32(1))%32)<<(uint(int32(2))%32)+int32(4))))
	v419 = F_objectGetVal(m, v418)
	mBase = m.M
	v422 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v419+int32(-1)))))
	switch v422 & int32(7) {
	case 0:
		goto L127
	case 1:
		goto L126
	case 2:
		goto L125
	case 3:
		goto L124
	case 4:
		goto L123
	default:
		v439 = int32(0)
		goto L122
	}
L122:
	;
	if base.Ui32(v393) < base.Ui32(v439) {
		goto L128
	} else {
		goto L129
	}
L123:
	;
	v438 = *(*int32)(unsafe.Add(mBase, uint32(v419+int32(-17))))
	v439 = v438
	goto L122
L124:
	;
	v435 = *(*int32)(unsafe.Add(mBase, uint32(v419+int32(-9))))
	v439 = v435
	goto L122
L125:
	;
	v432 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v419+int32(-5)))))
	v439 = v432
	goto L122
L126:
	;
	v429 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v419+int32(-3)))))
	v439 = v429
	goto L122
L127:
	;
	v439 = int32(base.Ui32(v422) >> (uint(int32(3)) % 32))
	goto L122
L128:
	;
	v441 = v439
	goto L130
L129:
	;
	v441 = v393
	goto L130
L130:
	;
	v443 = v387 + int32(1)
	if v443 != v361 {
		v387 = v443
		v393 = v441
		goto L118
	} else {
		goto L131
	}
L131:
	;
	goto L119
L132:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+28)) = v462
	v466 = F_checkType(m, l0, v462, int32(3))
	mBase = m.M
	v467 = m.ExcPending
	if v467 != 0 {
		goto L98
	} else {
		goto L133
	}
L133:
	;
	if v466 != 0 {
		goto L114
	} else {
		goto L134
	}
L134:
	;
	if v462 != 0 {
		goto L142
	} else {
		goto L143
	}
L135:
	;
	F_addReplyNull(m, l0)
	mBase = m.M
	v659 = m.ExcPending
	if v659 != 0 {
		goto L98
	} else {
		goto L174
	}
L136:
	;
	if v591 == int32(0) {
		goto L135
	} else {
		goto L172
	}
L137:
	;
	if v317 != 0 {
		goto L168
	} else {
		goto L169
	}
L138:
	;
	v595 = int32(0)
	if v586|v590 == v595 {
		v610 = v595
		goto L157
	} else {
		goto L158
	}
L139:
	;
	v586 = v500
	v590 = v504
	v591 = v505
	v593 = int32(1)
	v594 = v512
	goto L138
L140:
	;
	v574 = int32(0)
	if v363 == v574 {
		v620 = v574
		v623 = v574
		goto L137
	} else {
		goto L156
	}
L141:
	;
	if int32(1) <= v361 {
		goto L149
	} else {
		goto L150
	}
L142:
	;
	F_zsetTypeMaybeConvert(m, v462, v361, v456)
	mBase = m.M
	v479 = m.ExcPending
	if v479 != 0 {
		goto L98
	} else {
		goto L147
	}
L143:
	;
	if v312&int32(4) != 0 {
		goto L140
	} else {
		goto L144
	}
L144:
	;
	v470 = F_zsetTypeCreate(m, v361, v456)
	mBase = m.M
	v471 = m.ExcPending
	if v471 != 0 {
		goto L98
	} else {
		goto L145
	}
L145:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+28)) = v470
	v473 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	F_dbAdd(m, v473, v22, v19+int32(28))
	mBase = m.M
	v477 = m.ExcPending
	if v477 != 0 {
		goto L98
	} else {
		goto L146
	}
L146:
	;
	goto L141
L147:
	;
	goto L141
L148:
	;
	v567 = int32(_a44)
	v569 = *(*int64)(unsafe.Add(mBase, _consts[83]))
	*(*int64)(unsafe.Add(mBase, _consts[83])) = v569 + base.I64_extend_i32_s(v562+v558)
	v586 = v558
	v590 = v562
	v591 = v563
	v593 = int32(0)
	v594 = v566
	goto L138
L149:
	;
	v489 = int32(0)
	v497 = v489
	v500 = v489
	v504 = v489
	v505 = v489
	goto L151
L150:
	;
	v484 = int32(0)
	v558 = v484
	v562 = v484
	v563 = v484
	v566 = float64(0)
	goto L148
L151:
	;
	v510 = v497 << (uint(int32(3)) % 32)
	v512 = *(*float64)(unsafe.Add(mBase, uint32(v374+v510)))
	*(*int32)(unsafe.Add(mBase, uint32(v19)+12)) = int32(0)
	v515 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v520 = *(*int32)(unsafe.Add(mBase, uint32(v515+v315<<(uint(int32(2))%32)+v510+int32(4))))
	v521 = F_objectGetVal(m, v520)
	mBase = m.M
	v522 = *(*int32)(unsafe.Add(mBase, uint32(v19)+28))
	v527 = F_zsetAdd(m, v522, v512, v521, v312, v19+int32(12), v19+int32(16))
	mBase = m.M
	v528 = m.ExcPending
	if v528 != 0 {
		goto L98
	} else {
		goto L153
	}
L152:
	;
	v550 = *(*float64)(unsafe.Add(mBase, uint32(v19)+16))
	v558 = v546
	v562 = v541
	v563 = v536
	v566 = v550
	goto L148
L153:
	;
	if v527 == int32(0) {
		goto L139
	} else {
		goto L154
	}
L154:
	;
	v531 = *(*int32)(unsafe.Add(mBase, uint32(v19)+12))
	v534 = int32(1)
	v536 = (v531^int32(-1))&v534 + v505
	v541 = int32(base.Ui32(v531)>>(uint(int32(3))%32))&v534 + v504
	v546 = int32(base.Ui32(v531)>>(uint(int32(2))%32))&v534 + v500
	v548 = v497 + v534
	if v548 != v361 {
		v497 = v548
		v500 = v546
		v504 = v541
		v505 = v536
		goto L151
	} else {
		goto L155
	}
L155:
	;
	goto L152
L156:
	;
	goto L135
L157:
	;
	if v593 == int32(0) {
		goto L164
	} else {
		goto L165
	}
L158:
	;
	v599 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	F_signalModifiedKey(m, l0, v599, v22)
	mBase = m.M
	v601 = m.ExcPending
	if v601 != 0 {
		goto L98
	} else {
		goto L159
	}
L159:
	;
	if v363 != 0 {
		goto L160
	} else {
		goto L161
	}
L160:
	;
	v605 = int32(_a1731)
	goto L162
L161:
	;
	v605 = int32(_a1732)
	goto L162
L162:
	;
	v606 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v607 = *(*int32)(unsafe.Add(mBase, uint32(v606)+28))
	F_notifyKeyspaceEvent(m, int32(128), v605, v22, v607)
	mBase = m.M
	v609 = m.ExcPending
	if v609 != 0 {
		goto L98
	} else {
		goto L163
	}
L163:
	;
	v610 = v590
	goto L157
L164:
	;
	if v363 != 0 {
		goto L136
	} else {
		goto L167
	}
L165:
	;
	F_addReplyError(m, l0, int32(_a1733))
	mBase = m.M
	v615 = m.ExcPending
	if v615 != 0 {
		goto L98
	} else {
		goto L166
	}
L166:
	;
	goto L114
L167:
	;
	v620 = v610
	v623 = v586
	goto L137
L168:
	;
	v633 = int32(0)
	goto L170
L169:
	;
	v633 = v620
	goto L170
L170:
	;
	F_addReplyLongLong(m, l0, base.I64_extend_i32_s(v633+v623))
	mBase = m.M
	v637 = m.ExcPending
	if v637 != 0 {
		goto L98
	} else {
		goto L171
	}
L171:
	;
	goto L114
L172:
	;
	F_addReplyDouble(m, l0, v594)
	mBase = m.M
	v641 = m.ExcPending
	if v641 != 0 {
		goto L98
	} else {
		goto L173
	}
L173:
	;
	goto L114
L174:
	;
	goto L114
L175:
	;
	goto L93
}
func F_zcardCommand(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
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
	var v18 int32
	_ = v18
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v4 = *(*int32)(unsafe.Add(mBase, uint32(v3)+4))
	v6 = *(*int32)(unsafe.Add(mBase, _consts[85]))
	v7 = F_lookupKeyReadOrReply(m, l0, v4, v6)
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		return
	} else {
		if v7 == int32(0) {
			return
		} else {
			v12 = F_checkType(m, l0, v7, int32(3))
			mBase = m.M
			v13 = m.ExcPending
			if v13 != 0 {
				return
			} else {
				if v12 != 0 {
					return
				} else {
					v14 = F_zsetLength(m, v7)
					mBase = m.M
					v15 = m.ExcPending
					if v15 != 0 {
						return
					} else {
						F_addReplyLongLong(m, l0, base.I64_extend_i32_u(v14))
						mBase = m.M
						v18 = m.ExcPending
						if v18 != 0 {
							return
						} else {
							return
						}
					}
				}
			}
		}
	}
}
func F_zcountCommand(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
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
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v79 int64
	_ = v79
	var v84 int64
	_ = v84
	var v92 int32
	_ = v92
	var v99 int32
	_ = v99
	var v105 float64
	_ = v105
	var v109 int64
	_ = v109
	var v111 float64
	_ = v111
	var v112 float64
	_ = v112
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v122 int32
	_ = v122
	var v128 int32
	_ = v128
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v144 int32
	_ = v144
	var v146 int32
	_ = v146
	var v148 int32
	_ = v148
	var v154 int64
	_ = v154
	var v159 int64
	_ = v159
	var v167 int32
	_ = v167
	var v174 int32
	_ = v174
	var v180 float64
	_ = v180
	var v184 int64
	_ = v184
	var v186 float64
	_ = v186
	var v189 int32
	_ = v189
	var v195 int32
	_ = v195
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v203 int32
	_ = v203
	var v206 int32
	_ = v206
	var v224 float64
	_ = v224
	var v225 float64
	_ = v225
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v233 float64
	_ = v233
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v240 int32
	_ = v240
	var v243 float64
	_ = v243
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v252 int32
	_ = v252
	var v256 int32
	_ = v256
	var v258 int32
	_ = v258
	var v264 int32
	_ = v264
	var __phi264 int32
	_ = __phi264
	var v276 int32
	_ = v276
	var __phi276 int32
	_ = __phi276
	var v279 int32
	_ = v279
	var __phi279 int32
	_ = __phi279
	var v283 float64
	_ = v283
	var v286 int32
	_ = v286
	var v293 int32
	_ = v293
	var v294 int32
	_ = v294
	var v295 int32
	_ = v295
	var v299 int32
	_ = v299
	var v313 int32
	_ = v313
	var v314 int32
	_ = v314
	var v623 int32
	_ = v623
	var v629 int32
	_ = v629
	var v638 int32
	_ = v638
	var v643 int32
	_ = v643
	var v645 int32
	_ = v645
	var v655 int32
	_ = v655
	var __phi655 int32
	_ = __phi655
	var v664 int32
	_ = v664
	var __phi664 int32
	_ = __phi664
	var v666 int32
	_ = v666
	var __phi666 int32
	_ = __phi666
	var v668 float64
	_ = v668
	var v671 int32
	_ = v671
	var v678 int32
	_ = v678
	var v679 int32
	_ = v679
	var v680 int32
	_ = v680
	var v682 int32
	_ = v682
	var v684 int32
	_ = v684
	var v699 int32
	_ = v699
	var v708 int32
	_ = v708
	var v723 int32
	_ = v723
	var v727 int32
	_ = v727
	var v728 int32
	_ = v728
	var v729 int32
	_ = v729
	var v797 int32
	_ = v797
	var v803 int32
	_ = v803
	var v816 int32
	_ = v816
	var v818 int32
	_ = v818
	var v955 float64
	_ = v955
	var v958 int32
	_ = v958
	var v978 int32
	_ = v978
	var v1028 int32
	_ = v1028
	var v1046 int32
	_ = v1046
	var v1047 int32
	_ = v1047
	var v1050 int32
	_ = v1050
	var v1052 int32
	_ = v1052
	var v1055 int32
	_ = v1055
	var v1056 int32
	_ = v1056
	var v1073 float64
	_ = v1073
	var v1074 float64
	_ = v1074
	var v1077 int32
	_ = v1077
	var v1078 int32
	_ = v1078
	var v1079 int32
	_ = v1079
	var v1082 float64
	_ = v1082
	var v1085 int32
	_ = v1085
	var v1086 int32
	_ = v1086
	var v1089 int32
	_ = v1089
	var v1092 float64
	_ = v1092
	var v1095 int32
	_ = v1095
	var v1096 int32
	_ = v1096
	var v1101 int32
	_ = v1101
	var v1105 int32
	_ = v1105
	var v1107 int32
	_ = v1107
	var v1113 int32
	_ = v1113
	var __phi1113 int32
	_ = __phi1113
	var v1125 int32
	_ = v1125
	var __phi1125 int32
	_ = __phi1125
	var v1128 int32
	_ = v1128
	var __phi1128 int32
	_ = __phi1128
	var v1132 float64
	_ = v1132
	var v1135 int32
	_ = v1135
	var v1142 int32
	_ = v1142
	var v1143 int32
	_ = v1143
	var v1144 int32
	_ = v1144
	var v1148 int32
	_ = v1148
	var v1162 int32
	_ = v1162
	var v1163 int32
	_ = v1163
	var v1171 int32
	_ = v1171
	var v1175 int32
	_ = v1175
	var v1190 int32
	_ = v1190
	var v1191 int32
	_ = v1191
	var v1195 int32
	_ = v1195
	var v1197 int32
	_ = v1197
	var v1201 int32
	_ = v1201
	var v1211 int32
	_ = v1211
	var __phi1211 int32
	_ = __phi1211
	var v1220 int32
	_ = v1220
	var __phi1220 int32
	_ = __phi1220
	var v1222 int32
	_ = v1222
	var __phi1222 int32
	_ = __phi1222
	var v1224 float64
	_ = v1224
	var v1227 int32
	_ = v1227
	var v1230 int32
	_ = v1230
	var v1236 int32
	_ = v1236
	var v1237 int32
	_ = v1237
	var v1238 int32
	_ = v1238
	var v1242 int32
	_ = v1242
	var v1244 int32
	_ = v1244
	var v1259 int32
	_ = v1259
	var v1266 int32
	_ = v1266
	var v1281 int32
	_ = v1281
	var v1438 float64
	_ = v1438
	var v1441 int32
	_ = v1441
	var v1877 int32
	_ = v1877
	var v1896 int32
	_ = v1896
	var v1900 int32
	_ = v1900
	var v1908 int32
	_ = v1908
	var v1924 int32
	_ = v1924
	var v1930 int32
	_ = v1930
	var v1936 int32
	_ = v1936
	var v1943 int32
	_ = v1943
	v9 = m.G0
	v11 = v9 - int32(48)
	m.G0 = v11
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
	v15 = *(*int32)(unsafe.Add(mBase, uint32(v13)+8))
	v16 = *(*int32)(unsafe.Add(mBase, uint32(v13)+12))
	v19 = F_zslParseRange(m, v15, v16, v11+int32(8))
	mBase = m.M
	if v19 == int32(0) {
		goto L6
	} else {
		goto L7
	}
L1:
	;
	F__serverPanic_1(m, int32(_a1723), int32(3404), int32(_a1054), int32(0))
	mBase = m.M
	v1943 = m.ExcPending
	if v1943 != 0 {
		goto L8
	} else {
		goto L350
	}
L2:
	;
	F__serverAssert(m, int32(_a169), int32(_a1723), int32(857))
	mBase = m.M
	v1936 = m.ExcPending
	if v1936 != 0 {
		goto L8
	} else {
		goto L349
	}
L3:
	;
	F__serverAssertWithInfo(m, l0, v27, int32(_a1736), int32(_a1723), int32(3368))
	mBase = m.M
	v1930 = m.ExcPending
	if v1930 != 0 {
		goto L8
	} else {
		goto L348
	}
L4:
	;
	F__serverAssert(m, int32(_a169), int32(_a1723), int32(857))
	mBase = m.M
	v1924 = m.ExcPending
	if v1924 != 0 {
		goto L8
	} else {
		goto L347
	}
L5:
	;
	m.G0 = v11 + int32(48)
	return
L6:
	;
	v26 = *(*int32)(unsafe.Add(mBase, _consts[85]))
	v27 = F_lookupKeyReadOrReply(m, l0, v14, v26)
	mBase = m.M
	v28 = m.ExcPending
	if v28 != 0 {
		goto L8
	} else {
		goto L10
	}
L7:
	;
	F_addReplyError(m, l0, int32(_a1737))
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	return
L9:
	;
	goto L5
L10:
	;
	if v27 == int32(0) {
		goto L5
	} else {
		goto L11
	}
L11:
	;
	v32 = F_checkType(m, l0, v27, int32(3))
	mBase = m.M
	v33 = m.ExcPending
	if v33 != 0 {
		goto L8
	} else {
		goto L12
	}
L12:
	;
	if v32 != 0 {
		goto L5
	} else {
		goto L13
	}
L13:
	;
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v27)))
	switch int32(base.Ui32(v34)>>(uint(int32(4))%32))&int32(15) + int32(-7) {
	case 0:
		goto L15
	default:
		goto L1
	case 4:
		goto L16
	}
L14:
	;
	F_addReplyLongLong(m, l0, base.I64_extend_i32_u(v1900))
	mBase = m.M
	v1908 = m.ExcPending
	if v1908 != 0 {
		goto L8
	} else {
		goto L346
	}
L15:
	;
	v199 = int32(0)
	v200 = F_objectGetVal(m, v27)
	mBase = m.M
	v201 = *(*int32)(unsafe.Add(mBase, uint32(v200)+4))
	v203 = v11 + int32(8)
	v206 = v11 + int32(32)
	v224 = *(*float64)(unsafe.Add(mBase, uint32(v203)))
	v225 = *(*float64)(unsafe.Add(mBase, uint32(v203)+8))
	if base.F64_gt(v224, v225) != 0 {
		v1028 = v199
		goto L61
	} else {
		goto L62
	}
L16:
	;
	v41 = F_objectGetVal(m, v27)
	mBase = m.M
	v44 = F_zzlFirstInRange(m, v41, v11+int32(8))
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L8
	} else {
		goto L17
	}
L17:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+4)) = v44
	if v44 != 0 {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	v51 = F_lpNext(m, v41, v44)
	mBase = m.M
	v52 = m.ExcPending
	if v52 != 0 {
		goto L8
	} else {
		goto L21
	}
L19:
	;
	v48 = *(*int32)(unsafe.Add(mBase, _consts[85]))
	F_addReply(m, l0, v48)
	mBase = m.M
	v50 = m.ExcPending
	if v50 != 0 {
		goto L8
	} else {
		goto L20
	}
L20:
	;
	goto L5
L21:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11))) = v51
	if v51 == int32(0) {
		goto L4
	} else {
		goto L22
	}
L22:
	;
	v60 = F_lpGetValue(m, v51, v11+int32(44), v11+int32(32))
	mBase = m.M
	v61 = m.ExcPending
	if v61 != 0 {
		goto L8
	} else {
		goto L25
	}
L23:
	;
	v112 = *(*float64)(unsafe.Add(mBase, uint32(v11)+16))
	v115 = *(*int32)(unsafe.Add(mBase, uint32(v11)+28))
	if v115 != 0 {
		goto L35
	} else {
		goto L36
	}
L24:
	;
	v109 = *(*int64)(unsafe.Add(mBase, uint32(v11)+32))
	v111 = base.F64_convert_i64_s(v109)
	goto L23
L25:
	;
	if v60 == int32(0) {
		goto L24
	} else {
		goto L26
	}
L26:
	;
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v11)+44))
	v65 = int32(0)
	v69 = m.G0
	v71 = v69 - int32(32)
	m.G0 = v71
	v73 = F___errno_location(m)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v73))) = v65
	v79 = *(*int64)(unsafe.Add(mBase, _consts[378]))
	*(*int64)(unsafe.Add(mBase, uint32(v71+int32(8)))) = v79
	*(*int64)(unsafe.Add(mBase, uint32(v71)+24)) = int64(0)
	v84 = *(*int64)(unsafe.Add(mBase, _consts[379]))
	*(*int64)(unsafe.Add(mBase, uint32(v71))) = v84
	F_ffc_from_chars_double_options(m, v71+int32(16), v60, v60+v64, v71+int32(24), v71)
	mBase = m.M
	v92 = *(*int32)(unsafe.Add(mBase, uint32(v71)+20))
	if v92 == v65 {
		goto L28
	} else {
		goto L29
	}
L27:
	;
	v111 = v105
	goto L23
L28:
	;
	goto L33
L29:
	;
	if v92 == int32(2) {
		goto L30
	} else {
		goto L31
	}
L30:
	;
	v99 = int32(68)
	goto L32
L31:
	;
	v99 = int32(28)
	goto L32
L32:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v73))) = v99
	goto L28
L33:
	;
	v105 = *(*float64)(unsafe.Add(mBase, uint32(v71)+24))
	m.G0 = v71 + int32(32)
	goto L27
L35:
	;
	v116 = base.F64_lt(v111, v112)
	goto L37
L36:
	;
	v116 = base.F64_le(v111, v112)
	goto L37
L37:
	;
	if v116 == int32(0) {
		goto L3
	} else {
		goto L38
	}
L38:
	;
	v122 = int32(0)
	goto L39
L39:
	;
	v128 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
	if v128 == int32(0) {
		goto L2
	} else {
		goto L41
	}
L41:
	;
	v135 = F_lpGetValue(m, v128, v11+int32(44), v11+int32(32))
	mBase = m.M
	v136 = m.ExcPending
	if v136 != 0 {
		goto L8
	} else {
		goto L44
	}
L42:
	;
	if v115 != 0 {
		goto L54
	} else {
		goto L55
	}
L43:
	;
	v184 = *(*int64)(unsafe.Add(mBase, uint32(v11)+32))
	v186 = base.F64_convert_i64_s(v184)
	goto L42
L44:
	;
	if v135 == int32(0) {
		goto L43
	} else {
		goto L45
	}
L45:
	;
	v139 = *(*int32)(unsafe.Add(mBase, uint32(v11)+44))
	v140 = int32(0)
	v144 = m.G0
	v146 = v144 - int32(32)
	m.G0 = v146
	v148 = F___errno_location(m)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v148))) = v140
	v154 = *(*int64)(unsafe.Add(mBase, _consts[378]))
	*(*int64)(unsafe.Add(mBase, uint32(v146+int32(8)))) = v154
	*(*int64)(unsafe.Add(mBase, uint32(v146)+24)) = int64(0)
	v159 = *(*int64)(unsafe.Add(mBase, _consts[379]))
	*(*int64)(unsafe.Add(mBase, uint32(v146))) = v159
	F_ffc_from_chars_double_options(m, v146+int32(16), v135, v135+v139, v146+int32(24), v146)
	mBase = m.M
	v167 = *(*int32)(unsafe.Add(mBase, uint32(v146)+20))
	if v167 == v140 {
		goto L47
	} else {
		goto L48
	}
L46:
	;
	v186 = v180
	goto L42
L47:
	;
	goto L52
L48:
	;
	if v167 == int32(2) {
		goto L49
	} else {
		goto L50
	}
L49:
	;
	v174 = int32(68)
	goto L51
L50:
	;
	v174 = int32(28)
	goto L51
L51:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v148))) = v174
	goto L47
L52:
	;
	v180 = *(*float64)(unsafe.Add(mBase, uint32(v146)+24))
	m.G0 = v146 + int32(32)
	goto L46
L54:
	;
	v189 = base.F64_lt(v186, v112)
	goto L56
L55:
	;
	v189 = base.F64_le(v186, v112)
	goto L56
L56:
	;
	if v189 != int32(1) {
		v1900 = v122
		goto L14
	} else {
		goto L57
	}
L57:
	;
	F_zzlNext(m, v41, v11+int32(4), v11)
	mBase = m.M
	v195 = m.ExcPending
	if v195 != 0 {
		goto L8
	} else {
		goto L58
	}
L58:
	;
	v197 = v122 + int32(1)
	v198 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
	if v198 != 0 {
		v122 = v197
		goto L39
	} else {
		goto L59
	}
L59:
	;
	v1900 = v197
	goto L14
L60:
	;
	if v1028 == int32(0) {
		v1900 = v199
		goto L14
	} else {
		goto L202
	}
L61:
	;
	goto L60
L62:
	;
	if base.F64_ne(v224, v225) != 0 {
		goto L63
	} else {
		goto L64
	}
L63:
	;
	v230 = *(*int32)(unsafe.Add(mBase, uint32(v201)+8))
	if v230 == int32(0) {
		v1028 = v199
		goto L61
	} else {
		goto L67
	}
L64:
	;
	v228 = *(*int32)(unsafe.Add(mBase, uint32(v203)+16))
	if v228 != 0 {
		v1028 = v199
		goto L61
	} else {
		goto L65
	}
L65:
	;
	v229 = *(*int32)(unsafe.Add(mBase, uint32(v203)+20))
	if v229 != 0 {
		v1028 = v199
		goto L61
	} else {
		goto L66
	}
L66:
	;
	goto L63
L67:
	;
	v233 = *(*float64)(unsafe.Add(mBase, uint32(v230)))
	v236 = *(*int32)(unsafe.Add(mBase, uint32(v203)+16))
	if v236 != 0 {
		goto L68
	} else {
		goto L69
	}
L68:
	;
	v237 = base.F64_gt(v233, v224)
	goto L70
L69:
	;
	v237 = base.F64_ge(v233, v224)
	goto L70
L70:
	;
	if v237 != int32(1) {
		v1028 = v199
		goto L61
	} else {
		goto L71
	}
L71:
	;
	v240 = *(*int32)(unsafe.Add(mBase, uint32(v201)+12))
	if v240 == int32(0) {
		v1028 = v199
		goto L61
	} else {
		goto L72
	}
L72:
	;
	v243 = *(*float64)(unsafe.Add(mBase, uint32(v240)))
	v246 = *(*int32)(unsafe.Add(mBase, uint32(v203)+20))
	if v246 != 0 {
		goto L73
	} else {
		goto L74
	}
L73:
	;
	v247 = base.F64_lt(v243, v225)
	goto L75
L74:
	;
	v247 = base.F64_le(v243, v225)
	goto L75
L75:
	;
	if v247 != int32(1) {
		v1028 = v199
		goto L61
	} else {
		goto L76
	}
L76:
	;
	v252 = *(*int32)(unsafe.Add(mBase, uint32(v201)+16))
	v256 = (v252 + int32(-1)) << (uint(int32(3)) % 32)
	v258 = *(*int32)(unsafe.Add(mBase, uint32(v201+int32(12)+v256)))
	if v258 != 0 {
		goto L78
	} else {
		goto L79
	}
L77:
	;
	goto L92
L78:
	;
	__phi264 = v258
	__phi276 = int32(0)
	__phi279 = v201
	v264 = __phi264
	v276 = __phi276
	v279 = __phi279
	goto L80
L79:
	;
	v313 = int32(0)
	v314 = v201
	goto L77
L80:
	;
	v283 = *(*float64)(unsafe.Add(mBase, uint32(v264)))
	if v236 != 0 {
		goto L83
	} else {
		goto L84
	}
L81:
	;
	v313 = v295
	v314 = v264
	goto L77
L82:
	;
	if v252 < int32(2) {
		v294 = int32(1)
		goto L87
	} else {
		goto L88
	}
L83:
	;
	v286 = base.F64_gt(v283, v224)
	goto L85
L84:
	;
	v286 = base.F64_ge(v283, v224)
	goto L85
L85:
	;
	if v286 == int32(0) {
		goto L82
	} else {
		goto L86
	}
L86:
	;
	v313 = v276
	v314 = v279
	goto L77
L87:
	;
	v295 = v294 + v276
	v299 = *(*int32)(unsafe.Add(mBase, uint32(v264+v256+int32(12))))
	if v299 != 0 {
		__phi264 = v299
		__phi276 = v295
		__phi279 = v264
		v264 = __phi264
		v276 = __phi276
		v279 = __phi279
		goto L80
	} else {
		goto L89
	}
L88:
	;
	v293 = *(*int32)(unsafe.Add(mBase, uint32(v279+v256+int32(16))))
	v294 = v293
	goto L87
L89:
	;
	goto L81
L90:
	;
	v1028 = v978
	goto L61
L91:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v206))) = v728
	goto L90
L92:
	;
	if v252 < int32(2) {
		v708 = v314
		v723 = v313
		goto L146
	} else {
		goto L147
	}
L146:
	;
	v727 = int32(0)
	v728 = v723 + v199
	v729 = *(*int32)(unsafe.Add(mBase, uint32(v201)))
	if base.Ui32(v729) <= base.Ui32(v728) {
		v1028 = v727
		goto L61
	} else {
		goto L164
	}
L147:
	;
	v623 = v314
	v629 = v252 + int32(-2)
	v638 = v313
	goto L148
L148:
	;
	v643 = v629 << (uint(int32(3)) % 32)
	v645 = *(*int32)(unsafe.Add(mBase, uint32(v623+v643)+12))
	if v645 == int32(0) {
		v684 = v623
		v699 = v638
		goto L150
	} else {
		goto L151
	}
L149:
	;
	v708 = v684
	v723 = v699
	goto L146
L150:
	;
	if int32(0) < v629 {
		v623 = v684
		v629 = v629 + int32(-1)
		v638 = v699
		goto L148
	} else {
		goto L163
	}
L151:
	;
	__phi655 = v645
	__phi664 = v638
	__phi666 = v623
	v655 = __phi655
	v664 = __phi664
	v666 = __phi666
	goto L152
L152:
	;
	v668 = *(*float64)(unsafe.Add(mBase, uint32(v655)))
	if v236 != 0 {
		goto L155
	} else {
		goto L156
	}
L153:
	;
	v684 = v655
	v699 = v680
	goto L150
L154:
	;
	if v629 != 0 {
		goto L160
	} else {
		goto L161
	}
L155:
	;
	v671 = base.F64_gt(v668, v224)
	goto L157
L156:
	;
	v671 = base.F64_ge(v668, v224)
	goto L157
L157:
	;
	if v671 == int32(0) {
		goto L154
	} else {
		goto L158
	}
L158:
	;
	v684 = v666
	v699 = v664
	goto L150
L159:
	;
	v680 = v679 + v664
	v682 = *(*int32)(unsafe.Add(mBase, uint32(v655+v643)+12))
	if v682 != 0 {
		__phi655 = v682
		__phi664 = v680
		__phi666 = v655
		v655 = __phi655
		v664 = __phi664
		v666 = __phi666
		goto L152
	} else {
		goto L162
	}
L160:
	;
	v678 = *(*int32)(unsafe.Add(mBase, uint32(v666+v643+int32(16))))
	v679 = v678
	goto L159
L161:
	;
	v679 = int32(1)
	goto L159
L162:
	;
	goto L153
L163:
	;
	goto L149
L164:
	;
	goto L168
L165:
	;
	if v206 == int32(0) {
		v1028 = v978
		goto L61
	} else {
		goto L201
	}
L166:
	;
	v955 = *(*float64)(unsafe.Add(mBase, uint32(v816)))
	if v246 != 0 {
		goto L197
	} else {
		goto L198
	}
L168:
	;
	goto L169
L169:
	;
	goto L175
L174:
	;
	if v816 != 0 {
		goto L166
	} else {
		goto L179
	}
L175:
	;
	v797 = v708
	v803 = int32(0)
	goto L176
L176:
	;
	v816 = *(*int32)(unsafe.Add(mBase, uint32(v797)+12))
	v818 = v803 + int32(1)
	if v818 != int32(1) {
		v797 = v816
		v803 = v818
		goto L176
	} else {
		goto L178
	}
L177:
	;
	goto L174
L178:
	;
	goto L177
L179:
	;
	v978 = int32(0)
	goto L165
L197:
	;
	v958 = base.F64_lt(v955, v225)
	goto L199
L198:
	;
	v958 = base.F64_le(v955, v225)
	goto L199
L199:
	;
	if v958 != int32(1) {
		v1028 = v727
		goto L61
	} else {
		goto L200
	}
L200:
	;
	v978 = v816
	goto L165
L201:
	;
	goto L91
L202:
	;
	v1046 = *(*int32)(unsafe.Add(mBase, uint32(v201)))
	v1047 = *(*int32)(unsafe.Add(mBase, uint32(v11)+32))
	v1050 = v1046 - v1047 + int32(1)
	v1052 = v11 + int32(8)
	v1055 = v11 + int32(32)
	v1056 = int32(0)
	v1073 = *(*float64)(unsafe.Add(mBase, uint32(v1052)))
	v1074 = *(*float64)(unsafe.Add(mBase, uint32(v1052)+8))
	if base.F64_gt(v1073, v1074) != 0 {
		v1877 = v1056
		goto L204
	} else {
		goto L205
	}
L203:
	;
	if v1877 == int32(0) {
		v1900 = v1050
		goto L14
	} else {
		goto L345
	}
L204:
	;
	goto L203
L205:
	;
	if base.F64_ne(v1073, v1074) != 0 {
		goto L206
	} else {
		goto L207
	}
L206:
	;
	v1079 = *(*int32)(unsafe.Add(mBase, uint32(v201)+8))
	if v1079 == int32(0) {
		v1877 = v1056
		goto L204
	} else {
		goto L210
	}
L207:
	;
	v1077 = *(*int32)(unsafe.Add(mBase, uint32(v1052)+16))
	if v1077 != 0 {
		v1877 = v1056
		goto L204
	} else {
		goto L208
	}
L208:
	;
	v1078 = *(*int32)(unsafe.Add(mBase, uint32(v1052)+20))
	if v1078 != 0 {
		v1877 = v1056
		goto L204
	} else {
		goto L209
	}
L209:
	;
	goto L206
L210:
	;
	v1082 = *(*float64)(unsafe.Add(mBase, uint32(v1079)))
	v1085 = *(*int32)(unsafe.Add(mBase, uint32(v1052)+16))
	if v1085 != 0 {
		goto L211
	} else {
		goto L212
	}
L211:
	;
	v1086 = base.F64_gt(v1082, v1073)
	goto L213
L212:
	;
	v1086 = base.F64_ge(v1082, v1073)
	goto L213
L213:
	;
	if v1086 != int32(1) {
		v1877 = v1056
		goto L204
	} else {
		goto L214
	}
L214:
	;
	v1089 = *(*int32)(unsafe.Add(mBase, uint32(v201)+12))
	if v1089 == int32(0) {
		v1877 = v1056
		goto L204
	} else {
		goto L215
	}
L215:
	;
	v1092 = *(*float64)(unsafe.Add(mBase, uint32(v1089)))
	v1095 = *(*int32)(unsafe.Add(mBase, uint32(v1052)+20))
	if v1095 != 0 {
		goto L216
	} else {
		goto L217
	}
L216:
	;
	v1096 = base.F64_lt(v1092, v1074)
	goto L218
L217:
	;
	v1096 = base.F64_le(v1092, v1074)
	goto L218
L218:
	;
	if v1096 != int32(1) {
		v1877 = v1056
		goto L204
	} else {
		goto L219
	}
L219:
	;
	v1101 = *(*int32)(unsafe.Add(mBase, uint32(v201)+16))
	v1105 = (v1101 + int32(-1)) << (uint(int32(3)) % 32)
	v1107 = *(*int32)(unsafe.Add(mBase, uint32(v201+int32(12)+v1105)))
	if v1107 != 0 {
		goto L221
	} else {
		goto L222
	}
L220:
	;
	goto L236
L221:
	;
	__phi1113 = v1107
	__phi1125 = int32(0)
	__phi1128 = v201
	v1113 = __phi1113
	v1125 = __phi1125
	v1128 = __phi1128
	goto L223
L222:
	;
	v1162 = int32(0)
	v1163 = v201
	goto L220
L223:
	;
	v1132 = *(*float64)(unsafe.Add(mBase, uint32(v1113)))
	if v1085 != 0 {
		goto L226
	} else {
		goto L227
	}
L224:
	;
	v1162 = v1144
	v1163 = v1113
	goto L220
L225:
	;
	if v1101 < int32(2) {
		v1143 = int32(1)
		goto L230
	} else {
		goto L231
	}
L226:
	;
	v1135 = base.F64_gt(v1132, v1073)
	goto L228
L227:
	;
	v1135 = base.F64_ge(v1132, v1073)
	goto L228
L228:
	;
	if v1135 == int32(0) {
		goto L225
	} else {
		goto L229
	}
L229:
	;
	v1162 = v1125
	v1163 = v1128
	goto L220
L230:
	;
	v1144 = v1143 + v1125
	v1148 = *(*int32)(unsafe.Add(mBase, uint32(v1113+v1105+int32(12))))
	if v1148 != 0 {
		__phi1113 = v1148
		__phi1125 = v1144
		__phi1128 = v1113
		v1113 = __phi1113
		v1125 = __phi1125
		v1128 = __phi1128
		goto L223
	} else {
		goto L232
	}
L231:
	;
	v1142 = *(*int32)(unsafe.Add(mBase, uint32(v1128+v1105+int32(16))))
	v1143 = v1142
	goto L230
L232:
	;
	goto L224
L233:
	;
	v1877 = v1266
	goto L204
L234:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1055))) = v1281 + int32(-1)
	goto L233
L236:
	;
	v1171 = int32(0)
	if v1101 <= v1171 {
		v1266 = v1163
		v1281 = v1162
		goto L237
	} else {
		goto L238
	}
L237:
	;
	if v1281 < int32(1) {
		v1877 = v1171
		goto L204
	} else {
		goto L254
	}
L238:
	;
	v1175 = v1163
	v1190 = v1162
	v1191 = v1101
	goto L239
L239:
	;
	v1195 = v1191 + int32(-1)
	v1197 = v1195 << (uint(int32(3)) % 32)
	v1201 = *(*int32)(unsafe.Add(mBase, uint32(v1175+v1197+int32(12))))
	if v1201 == int32(0) {
		v1244 = v1175
		v1259 = v1190
		goto L241
	} else {
		goto L242
	}
L240:
	;
	v1266 = v1244
	v1281 = v1259
	goto L237
L241:
	;
	if int32(1) < v1191 {
		v1175 = v1244
		v1190 = v1259
		v1191 = v1195
		goto L239
	} else {
		goto L253
	}
L242:
	;
	__phi1211 = v1201
	__phi1220 = v1190
	__phi1222 = v1175
	v1211 = __phi1211
	v1220 = __phi1220
	v1222 = __phi1222
	goto L243
L243:
	;
	v1224 = *(*float64)(unsafe.Add(mBase, uint32(v1211)))
	if v1095 != 0 {
		goto L246
	} else {
		goto L247
	}
L244:
	;
	v1244 = v1211
	v1259 = v1238
	goto L241
L245:
	;
	v1230 = int32(1)
	if v1191 == v1230 {
		v1237 = v1230
		goto L250
	} else {
		goto L251
	}
L246:
	;
	v1227 = base.F64_lt(v1224, v1074)
	goto L248
L247:
	;
	v1227 = base.F64_le(v1224, v1074)
	goto L248
L248:
	;
	if v1227 == int32(1) {
		goto L245
	} else {
		goto L249
	}
L249:
	;
	v1244 = v1222
	v1259 = v1220
	goto L241
L250:
	;
	v1238 = v1237 + v1220
	v1242 = *(*int32)(unsafe.Add(mBase, uint32(v1211+v1197+int32(12))))
	if v1242 != 0 {
		__phi1211 = v1242
		__phi1220 = v1238
		__phi1222 = v1211
		v1211 = __phi1211
		v1220 = __phi1220
		v1222 = __phi1222
		goto L243
	} else {
		goto L252
	}
L251:
	;
	v1236 = *(*int32)(unsafe.Add(mBase, uint32(v1222+v1197+int32(16))))
	v1237 = v1236
	goto L250
L252:
	;
	goto L244
L253:
	;
	goto L240
L254:
	;
	goto L258
L255:
	;
	if v1055 == int32(0) {
		goto L233
	} else {
		goto L288
	}
L256:
	;
	v1438 = *(*float64)(unsafe.Add(mBase, uint32(v1266)))
	if v1085 != 0 {
		goto L284
	} else {
		goto L285
	}
L258:
	;
	goto L256
L284:
	;
	v1441 = base.F64_gt(v1438, v1073)
	goto L286
L285:
	;
	v1441 = base.F64_ge(v1438, v1073)
	goto L286
L286:
	;
	if v1441 != int32(1) {
		v1877 = v1171
		goto L204
	} else {
		goto L287
	}
L287:
	;
	goto L255
L288:
	;
	goto L234
L345:
	;
	v1896 = *(*int32)(unsafe.Add(mBase, uint32(v11)+32))
	v1900 = v1050 - v1046 + v1896
	goto L14
L346:
	;
	goto L5
L347:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L348:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L349:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L350:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_zdiffstoreCommand(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v8 int32
	_ = v8
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v3 = *(*int32)(unsafe.Add(mBase, uint32(v2)+4))
	F_zunionInterDiffGenericCommand(m, l0, v3, int32(2), int32(1), int32(0))
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		return
	} else {
		return
	}
}
func F_zfree_with_size(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v99 int32
	_ = v99
	var v104 int32
	_ = v104
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v115 int32
	_ = v115
	var __phi115 int32
	_ = __phi115
	var v116 int32
	_ = v116
	var __phi116 int32
	_ = __phi116
	var v122 int32
	_ = v122
	var v125 int32
	_ = v125
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v147 int32
	_ = v147
	var v156 int32
	_ = v156
	var v158 int32
	_ = v158
	var v161 int32
	_ = v161
	var v164 int32
	_ = v164
	var v166 int32
	_ = v166
	var v171 int32
	_ = v171
	var v178 int32
	_ = v178
	var v183 int32
	_ = v183
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v198 int32
	_ = v198
	var v206 int32
	_ = v206
	var v208 int32
	_ = v208
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v219 int32
	_ = v219
	var v221 int32
	_ = v221
	var v228 int32
	_ = v228
	var v230 int32
	_ = v230
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v248 int32
	_ = v248
	var v250 int32
	_ = v250
	var v252 int32
	_ = v252
	var v261 int32
	_ = v261
	var v263 int32
	_ = v263
	var v266 int32
	_ = v266
	var v271 int32
	_ = v271
	var v276 int32
	_ = v276
	var v277 int32
	_ = v277
	var v282 int32
	_ = v282
	var __phi282 int32
	_ = __phi282
	var v283 int32
	_ = v283
	var __phi283 int32
	_ = __phi283
	var v289 int32
	_ = v289
	var v292 int32
	_ = v292
	var v306 int32
	_ = v306
	var v315 int32
	_ = v315
	var v317 int32
	_ = v317
	var v320 int32
	_ = v320
	var v323 int32
	_ = v323
	var v325 int32
	_ = v325
	var v330 int32
	_ = v330
	var v337 int32
	_ = v337
	var v342 int32
	_ = v342
	var v364 int32
	_ = v364
	var v376 int32
	_ = v376
	var v378 int32
	_ = v378
	var v380 int32
	_ = v380
	var v384 int32
	_ = v384
	var v389 int32
	_ = v389
	var v390 int32
	_ = v390
	var v401 int32
	_ = v401
	var v404 int32
	_ = v404
	var v411 int32
	_ = v411
	var v416 int32
	_ = v416
	var v420 int32
	_ = v420
	var v422 int32
	_ = v422
	var v437 int32
	_ = v437
	var v439 int32
	_ = v439
	var v442 int32
	_ = v442
	var v445 int32
	_ = v445
	var v449 int32
	_ = v449
	var v459 int32
	_ = v459
	var v460 int32
	_ = v460
	var v466 int32
	_ = v466
	var v468 int32
	_ = v468
	var v471 int32
	_ = v471
	var v475 int32
	_ = v475
	var v481 int32
	_ = v481
	var v483 int32
	_ = v483
	var v484 int32
	_ = v484
	var v485 int32
	_ = v485
	var v486 int32
	_ = v486
	var v495 int32
	_ = v495
	var v497 int32
	_ = v497
	var v498 int32
	_ = v498
	var v499 int32
	_ = v499
	var v501 int32
	_ = v501
	var v528 int32
	_ = v528
	if l0 == int32(0) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	F__serverAssert(m, int32(_a1860), int32(_a1879), int32(463))
	mBase = m.M
	v528 = m.ExcPending
	if v528 != 0 {
		goto L112
	} else {
		goto L113
	}
L2:
	;
	return
L3:
	;
	v7 = l0 + int32(-8)
	v8 = *(*int32)(unsafe.Add(mBase, uint32(v7)))
	if int32(-1) < v8 {
		v16 = v7
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v18 = *(*int32)(unsafe.Add(mBase, _consts[411]))
	if v18 != int32(-1) {
		v29 = v18
		goto L7
	} else {
		goto L8
	}
L5:
	;
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0+int32(-12))))
	if v13 == int32(0) {
		goto L1
	} else {
		goto L6
	}
L6:
	;
	v16 = v13
	goto L4
L7:
	;
	v31 = l1 + int32(8)
	if v29 < int32(260) {
		goto L10
	} else {
		goto L11
	}
L8:
	;
	v21 = int32(0)
	v23 = *(*int32)(unsafe.Add(mBase, _consts[281]))
	*(*int32)(unsafe.Add(mBase, _consts[411])) = v23
	*(*int32)(unsafe.Add(mBase, _consts[281])) = v23 + int32(1)
	v29 = v23
	goto L7
L9:
	;
	if v16 == int32(0) {
		goto L13
	} else {
		goto L14
	}
L10:
	;
	v40 = v29 << (uint(int32(2)) % 32)
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v40)+uint32(_consts[285])))
	*(*int32)(unsafe.Add(mBase, uint32(v40)+uint32(_consts[285]))) = v43 - v31
	goto L9
L11:
	;
	v34 = int32(0)
	v36 = *(*int32)(unsafe.Add(mBase, _consts[286]))
	*(*int32)(unsafe.Add(mBase, _consts[286])) = v36 - v31
	goto L9
L12:
	;
	goto L2
L13:
	;
	goto L12
L14:
	;
	v57 = int32(-8)
	v58 = v16 + v57
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v16+int32(-4))))
	v63 = v61 & v57
	v64 = v58 + v63
	if v61&int32(1) != 0 {
		v188 = v63
		v189 = v58
		goto L15
	} else {
		goto L16
	}
L15:
	;
	if base.Ui32(v64) <= base.Ui32(v189) {
		goto L13
	} else {
		goto L50
	}
L16:
	;
	if v61&int32(2) == int32(0) {
		goto L13
	} else {
		goto L17
	}
L17:
	;
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v58)))
	v72 = v58 - v71
	v74 = *(*int32)(unsafe.Add(mBase, _consts[515]))
	if base.Ui32(v72) < base.Ui32(v74) {
		goto L13
	} else {
		goto L18
	}
L18:
	;
	v76 = v71 + v63
	v78 = *(*int32)(unsafe.Add(mBase, _consts[516]))
	if v72 == v78 {
		goto L22
	} else {
		goto L23
	}
L19:
	;
	if v94 == int32(0) {
		v188 = v76
		v189 = v72
		goto L15
	} else {
		goto L38
	}
L20:
	;
	v147 = int32(0)
	goto L19
L21:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v83)+12)) = v80
	*(*int32)(unsafe.Add(mBase, uint32(v80)+8)) = v83
	v188 = v76
	v189 = v72
	goto L15
L22:
	;
	v128 = *(*int32)(unsafe.Add(mBase, uint32(v64)+4))
	v129 = int32(3)
	if v128&v129 != v129 {
		v188 = v76
		v189 = v72
		goto L15
	} else {
		goto L37
	}
L23:
	;
	v80 = *(*int32)(unsafe.Add(mBase, uint32(v72)+12))
	if base.Ui32(int32(255)) < base.Ui32(v71) {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	v94 = *(*int32)(unsafe.Add(mBase, uint32(v72)+24))
	if v80 == v72 {
		goto L27
	} else {
		goto L28
	}
L25:
	;
	v83 = *(*int32)(unsafe.Add(mBase, uint32(v72)+8))
	if v80 != v83 {
		goto L21
	} else {
		goto L26
	}
L26:
	;
	v85 = int32(0)
	v87 = *(*int32)(unsafe.Add(mBase, _consts[517]))
	*(*int32)(unsafe.Add(mBase, _consts[517])) = v87 & base.I32_rotl(int32(-2), int32(base.Ui32(v71)>>(uint(int32(3))%32)))
	v188 = v76
	v189 = v72
	goto L15
L27:
	;
	v99 = *(*int32)(unsafe.Add(mBase, uint32(v72)+20))
	if v99 == int32(0) {
		goto L30
	} else {
		goto L31
	}
L28:
	;
	v96 = *(*int32)(unsafe.Add(mBase, uint32(v72)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v96)+12)) = v80
	*(*int32)(unsafe.Add(mBase, uint32(v80)+8)) = v96
	v147 = v80
	goto L19
L29:
	;
	__phi115 = v109
	__phi116 = v110
	v115 = __phi115
	v116 = __phi116
	goto L33
L30:
	;
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v72)+16))
	if v104 == int32(0) {
		goto L20
	} else {
		goto L32
	}
L31:
	;
	v109 = v99
	v110 = v72 + int32(20)
	goto L29
L32:
	;
	v109 = v104
	v110 = v72 + int32(16)
	goto L29
L33:
	;
	v122 = *(*int32)(unsafe.Add(mBase, uint32(v115)+20))
	if v122 != 0 {
		__phi115 = v122
		__phi116 = v115 + int32(20)
		v115 = __phi115
		v116 = __phi116
		goto L33
	} else {
		goto L35
	}
L34:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v116))) = int32(0)
	v147 = v115
	goto L19
L35:
	;
	v125 = *(*int32)(unsafe.Add(mBase, uint32(v115)+16))
	if v125 != 0 {
		__phi115 = v125
		__phi116 = v115 + int32(16)
		v115 = __phi115
		v116 = __phi116
		goto L33
	} else {
		goto L36
	}
L36:
	;
	goto L34
L37:
	;
	*(*int32)(unsafe.Add(mBase, _consts[518])) = v76
	*(*int32)(unsafe.Add(mBase, uint32(v64)+4)) = v128 & int32(-2)
	*(*int32)(unsafe.Add(mBase, uint32(v72)+4)) = v76 | int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v64))) = v76
	goto L12
L38:
	;
	v156 = *(*int32)(unsafe.Add(mBase, uint32(v72)+28))
	v158 = v156 << (uint(int32(2)) % 32)
	v161 = *(*int32)(unsafe.Add(mBase, uint32(v158)+uint32(_consts[519])))
	if v72 != v161 {
		goto L40
	} else {
		goto L41
	}
L39:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v147)+24)) = v94
	v178 = *(*int32)(unsafe.Add(mBase, uint32(v72)+16))
	if v178 == int32(0) {
		goto L47
	} else {
		goto L48
	}
L40:
	;
	v171 = *(*int32)(unsafe.Add(mBase, uint32(v94)+16))
	if v171 != v72 {
		goto L44
	} else {
		goto L45
	}
L41:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v158)+uint32(_consts[519]))) = v147
	if v147 != 0 {
		goto L39
	} else {
		goto L42
	}
L42:
	;
	v164 = int32(0)
	v166 = *(*int32)(unsafe.Add(mBase, _consts[520]))
	*(*int32)(unsafe.Add(mBase, _consts[520])) = v166 & base.I32_rotl(int32(-2), v156)
	v188 = v76
	v189 = v72
	goto L15
L43:
	;
	if v147 == int32(0) {
		v188 = v76
		v189 = v72
		goto L15
	} else {
		goto L46
	}
L44:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v94)+20)) = v147
	goto L43
L45:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v94)+16)) = v147
	goto L43
L46:
	;
	goto L39
L47:
	;
	v183 = *(*int32)(unsafe.Add(mBase, uint32(v72)+20))
	if v183 == int32(0) {
		v188 = v76
		v189 = v72
		goto L15
	} else {
		goto L49
	}
L48:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v147)+16)) = v178
	*(*int32)(unsafe.Add(mBase, uint32(v178)+24)) = v147
	goto L47
L49:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v147)+20)) = v183
	*(*int32)(unsafe.Add(mBase, uint32(v183)+24)) = v147
	v188 = v76
	v189 = v72
	goto L15
L50:
	;
	v198 = *(*int32)(unsafe.Add(mBase, uint32(v64)+4))
	if v198&int32(1) == int32(0) {
		goto L13
	} else {
		goto L51
	}
L51:
	;
	if v198&int32(2) != 0 {
		goto L56
	} else {
		goto L57
	}
L52:
	;
	if base.Ui32(int32(255)) < base.Ui32(v364) {
		goto L90
	} else {
		goto L91
	}
L53:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v189)+4)) = v244 | int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v189+v244))) = v244
	if v189 != v228 {
		v364 = v244
		goto L52
	} else {
		goto L89
	}
L54:
	;
	if v261 == int32(0) {
		goto L53
	} else {
		goto L77
	}
L55:
	;
	v306 = int32(0)
	goto L54
L56:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v64)+4)) = v198 & int32(-2)
	*(*int32)(unsafe.Add(mBase, uint32(v189)+4)) = v188 | int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v189+v188))) = v188
	v364 = v188
	goto L52
L57:
	;
	v206 = *(*int32)(unsafe.Add(mBase, _consts[521]))
	if v64 != v206 {
		goto L58
	} else {
		goto L59
	}
L58:
	;
	v228 = *(*int32)(unsafe.Add(mBase, _consts[516]))
	if v64 != v228 {
		goto L61
	} else {
		goto L62
	}
L59:
	;
	v208 = int32(0)
	*(*int32)(unsafe.Add(mBase, _consts[521])) = v189
	v212 = *(*int32)(unsafe.Add(mBase, _consts[522]))
	v213 = v212 + v188
	*(*int32)(unsafe.Add(mBase, _consts[522])) = v213
	*(*int32)(unsafe.Add(mBase, uint32(v189)+4)) = v213 | int32(1)
	v219 = *(*int32)(unsafe.Add(mBase, _consts[516]))
	if v189 != v219 {
		goto L13
	} else {
		goto L60
	}
L60:
	;
	v221 = int32(0)
	*(*int32)(unsafe.Add(mBase, _consts[518])) = v221
	*(*int32)(unsafe.Add(mBase, _consts[516])) = v221
	goto L12
L61:
	;
	v244 = v198&int32(-8) + v188
	v245 = *(*int32)(unsafe.Add(mBase, uint32(v64)+12))
	if base.Ui32(int32(255)) < base.Ui32(v198) {
		goto L63
	} else {
		goto L64
	}
L62:
	;
	v230 = int32(0)
	*(*int32)(unsafe.Add(mBase, _consts[516])) = v189
	v234 = *(*int32)(unsafe.Add(mBase, _consts[518]))
	v235 = v234 + v188
	*(*int32)(unsafe.Add(mBase, _consts[518])) = v235
	*(*int32)(unsafe.Add(mBase, uint32(v189)+4)) = v235 | int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v189+v235))) = v235
	goto L12
L63:
	;
	v261 = *(*int32)(unsafe.Add(mBase, uint32(v64)+24))
	if v245 == v64 {
		goto L67
	} else {
		goto L68
	}
L64:
	;
	v248 = *(*int32)(unsafe.Add(mBase, uint32(v64)+8))
	if v245 != v248 {
		goto L65
	} else {
		goto L66
	}
L65:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v248)+12)) = v245
	*(*int32)(unsafe.Add(mBase, uint32(v245)+8)) = v248
	goto L53
L66:
	;
	v250 = int32(0)
	v252 = *(*int32)(unsafe.Add(mBase, _consts[517]))
	*(*int32)(unsafe.Add(mBase, _consts[517])) = v252 & base.I32_rotl(int32(-2), int32(base.Ui32(v198)>>(uint(int32(3))%32)))
	goto L53
L67:
	;
	v266 = *(*int32)(unsafe.Add(mBase, uint32(v64)+20))
	if v266 == int32(0) {
		goto L70
	} else {
		goto L71
	}
L68:
	;
	v263 = *(*int32)(unsafe.Add(mBase, uint32(v64)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v263)+12)) = v245
	*(*int32)(unsafe.Add(mBase, uint32(v245)+8)) = v263
	v306 = v245
	goto L54
L69:
	;
	__phi282 = v276
	__phi283 = v277
	v282 = __phi282
	v283 = __phi283
	goto L73
L70:
	;
	v271 = *(*int32)(unsafe.Add(mBase, uint32(v64)+16))
	if v271 == int32(0) {
		goto L55
	} else {
		goto L72
	}
L71:
	;
	v276 = v266
	v277 = v64 + int32(20)
	goto L69
L72:
	;
	v276 = v271
	v277 = v64 + int32(16)
	goto L69
L73:
	;
	v289 = *(*int32)(unsafe.Add(mBase, uint32(v282)+20))
	if v289 != 0 {
		__phi282 = v289
		__phi283 = v282 + int32(20)
		v282 = __phi282
		v283 = __phi283
		goto L73
	} else {
		goto L75
	}
L74:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v283))) = int32(0)
	v306 = v282
	goto L54
L75:
	;
	v292 = *(*int32)(unsafe.Add(mBase, uint32(v282)+16))
	if v292 != 0 {
		__phi282 = v292
		__phi283 = v282 + int32(16)
		v282 = __phi282
		v283 = __phi283
		goto L73
	} else {
		goto L76
	}
L76:
	;
	goto L74
L77:
	;
	v315 = *(*int32)(unsafe.Add(mBase, uint32(v64)+28))
	v317 = v315 << (uint(int32(2)) % 32)
	v320 = *(*int32)(unsafe.Add(mBase, uint32(v317)+uint32(_consts[519])))
	if v64 != v320 {
		goto L79
	} else {
		goto L80
	}
L78:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v306)+24)) = v261
	v337 = *(*int32)(unsafe.Add(mBase, uint32(v64)+16))
	if v337 == int32(0) {
		goto L86
	} else {
		goto L87
	}
L79:
	;
	v330 = *(*int32)(unsafe.Add(mBase, uint32(v261)+16))
	if v330 != v64 {
		goto L83
	} else {
		goto L84
	}
L80:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v317)+uint32(_consts[519]))) = v306
	if v306 != 0 {
		goto L78
	} else {
		goto L81
	}
L81:
	;
	v323 = int32(0)
	v325 = *(*int32)(unsafe.Add(mBase, _consts[520]))
	*(*int32)(unsafe.Add(mBase, _consts[520])) = v325 & base.I32_rotl(int32(-2), v315)
	goto L53
L82:
	;
	if v306 == int32(0) {
		goto L53
	} else {
		goto L85
	}
L83:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v261)+20)) = v306
	goto L82
L84:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v261)+16)) = v306
	goto L82
L85:
	;
	goto L78
L86:
	;
	v342 = *(*int32)(unsafe.Add(mBase, uint32(v64)+20))
	if v342 == int32(0) {
		goto L53
	} else {
		goto L88
	}
L87:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v306)+16)) = v337
	*(*int32)(unsafe.Add(mBase, uint32(v337)+24)) = v306
	goto L86
L88:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v306)+20)) = v342
	*(*int32)(unsafe.Add(mBase, uint32(v342)+24)) = v306
	goto L53
L89:
	;
	*(*int32)(unsafe.Add(mBase, _consts[518])) = v244
	goto L12
L90:
	;
	if base.Ui32(int32(16777215)) < base.Ui32(v364) {
		v411 = int32(31)
		goto L95
	} else {
		goto L96
	}
L91:
	;
	v376 = v364 & int32(-8)
	v378 = v376 + int32(9128464)
	v380 = *(*int32)(unsafe.Add(mBase, _consts[517]))
	v384 = int32(1) << (uint(int32(base.Ui32(v364)>>(uint(int32(3))%32))) % 32)
	if v380&v384 != 0 {
		goto L93
	} else {
		goto L94
	}
L92:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v376)+uint32(_consts[523]))) = v189
	*(*int32)(unsafe.Add(mBase, uint32(v390)+12)) = v189
	*(*int32)(unsafe.Add(mBase, uint32(v189)+12)) = v378
	*(*int32)(unsafe.Add(mBase, uint32(v189)+8)) = v390
	goto L12
L93:
	;
	v389 = *(*int32)(unsafe.Add(mBase, uint32(v376)+uint32(_consts[523])))
	v390 = v389
	goto L92
L94:
	;
	*(*int32)(unsafe.Add(mBase, _consts[517])) = v380 | v384
	v390 = v378
	goto L92
L95:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v189)+28)) = v411
	*(*int64)(unsafe.Add(mBase, uint32(v189)+16)) = int64(0)
	v416 = v411 << (uint(int32(2)) % 32)
	v420 = *(*int32)(unsafe.Add(mBase, _consts[520]))
	v422 = int32(1) << (uint(v411) % 32)
	if v420&v422 != 0 {
		goto L100
	} else {
		goto L101
	}
L96:
	;
	v401 = base.I32_clz(int32(base.Ui32(v364) >> (uint(int32(8)) % 32)))
	v404 = int32(1)
	v411 = int32(base.Ui32(v364)>>(uint(int32(38)-v401)%32))&v404 - v401<<(uint(v404)%32) + int32(62)
	goto L95
L97:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v189+v483))) = v486
	*(*int32)(unsafe.Add(mBase, uint32(v189)+12)) = v485
	*(*int32)(unsafe.Add(mBase, uint32(v189+v481))) = v484
	v495 = int32(0)
	v497 = *(*int32)(unsafe.Add(mBase, _consts[524]))
	v498 = int32(-1)
	v499 = v497 + v498
	if v499 != 0 {
		goto L109
	} else {
		goto L110
	}
L98:
	;
	v475 = *(*int32)(unsafe.Add(mBase, uint32(v445)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v475)+12)) = v189
	*(*int32)(unsafe.Add(mBase, uint32(v445)+8)) = v189
	v481 = int32(24)
	v483 = int32(8)
	v484 = int32(0)
	v485 = v445
	v486 = v475
	goto L97
L99:
	;
	v481 = v466
	v483 = v468
	v484 = v189
	v485 = v189
	v486 = v471
	goto L97
L100:
	;
	if v411 == int32(31) {
		goto L102
	} else {
		goto L103
	}
L101:
	;
	*(*int32)(unsafe.Add(mBase, _consts[520])) = v420 | v422
	*(*int32)(unsafe.Add(mBase, uint32(v416)+uint32(_consts[519]))) = v189
	v466 = int32(8)
	v468 = int32(24)
	v471 = v416 + int32(9128728)
	goto L99
L102:
	;
	v437 = int32(0)
	goto L104
L103:
	;
	v437 = int32(25) - int32(base.Ui32(v411)>>(uint(int32(1))%32))
	goto L104
L104:
	;
	v439 = *(*int32)(unsafe.Add(mBase, uint32(v416)+uint32(_consts[519])))
	v442 = v364 << (uint(v437) % 32)
	v445 = v439
	goto L105
L105:
	;
	v449 = *(*int32)(unsafe.Add(mBase, uint32(v445)+4))
	if v449&int32(-8) == v364 {
		goto L98
	} else {
		goto L107
	}
L106:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v459+int32(16)))) = v189
	v466 = int32(8)
	v468 = int32(24)
	v471 = v445
	goto L99
L107:
	;
	v459 = v445 + int32(base.Ui32(v442)>>(uint(int32(29))%32))&int32(4)
	v460 = *(*int32)(unsafe.Add(mBase, uint32(v459)+16))
	if v460 != 0 {
		v442 = v442 << (uint(int32(1)) % 32)
		v445 = v460
		goto L105
	} else {
		goto L108
	}
L108:
	;
	goto L106
L109:
	;
	v501 = v499
	goto L111
L110:
	;
	v501 = v498
	goto L111
L111:
	;
	*(*int32)(unsafe.Add(mBase, _consts[524])) = v501
	goto L13
L112:
	;
	return
L113:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_ziplistPairsConvertAndValidateIntegrity(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
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
	var v32 int32
	_ = v32
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	*(*int32)(unsafe.Add(mBase, uint32(v7)+12)) = l2
	*(*int64)(unsafe.Add(mBase, uint32(v7)+4)) = int64(0)
	v16 = F_ziplistValidateIntegrity(m, l0, l1, int32(1), int32(963), v7+int32(4))
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		return int32(0)
	} else {
		v20 = *(*int32)(unsafe.Add(mBase, uint32(v7)+4))
		v21 = *(*int32)(unsafe.Add(mBase, uint32(v7)+8))
		if v21 == int32(0) {
			m.G0 = v7 + int32(16)
			if v20&int32(1) != 0 {
				v32 = int32(0)
			} else {
				v32 = v16
			}
			return v32
		} else {
			F_hashtableRelease(m, v21)
			mBase = m.M
			v25 = m.ExcPending
			if v25 != 0 {
				return int32(0)
			} else {
				m.G0 = v7 + int32(16)
				if v20&int32(1) != 0 {
					v32 = int32(0)
				} else {
					v32 = v16
				}
				return v32
			}
		}
	}
}
func F_zipmapRewind(m *base.Module, l0 int32) int32 {
	return l0 + int32(1)
}
func F_zmpopGenericCommand(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v18 int32
	_ = v18
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
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
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
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v103 int32
	_ = v103
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v130 int32
	_ = v130
	var v135 int32
	_ = v135
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v144 int32
	_ = v144
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v157 int32
	_ = v157
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v163 int32
	_ = v163
	var v167 int32
	_ = v167
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v194 int32
	_ = v194
	var v199 int32
	_ = v199
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v206 int32
	_ = v206
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v216 int32
	_ = v216
	var v219 int32
	_ = v219
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v232 int32
	_ = v232
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v242 int32
	_ = v242
	var v244 int32
	_ = v244
	var v246 int32
	_ = v246
	v10 = m.G0
	v12 = v10 - int32(16)
	m.G0 = v12
	*(*int32)(unsafe.Add(mBase, uint32(v12)+12)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v12)+8)) = int32(-1)
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v18+l1<<(uint(int32(2))%32))))
	v28 = F_getRangeLongFromObjectOrReply(m, l0, v22, int32(1), int32(2147483647), v12+int32(12), int32(_a1740))
	mBase = m.M
	v29 = m.ExcPending
	if v29 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v12 + int32(16)
	return
L2:
	;
	return
L3:
	;
	if v28 != 0 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v12)+12))
	v31 = v30 + l1
	v33 = v31 + int32(1)
	v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v33 < v34 {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	v40 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v42 = v33 << (uint(int32(2)) % 32)
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v40+v42)))
	v45 = F_objectGetVal(m, v44)
	mBase = m.M
	v46 = int32(_a1741)
	v49 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v45))))
	if v49 != 0 {
		goto L13
	} else {
		goto L14
	}
L6:
	;
	v37 = *(*int32)(unsafe.Add(mBase, _consts[86]))
	F_addReplyErrorObject(m, l0, v37)
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L2
	} else {
		goto L7
	}
L7:
	;
	goto L1
L8:
	;
	v244 = *(*int32)(unsafe.Add(mBase, _consts[86]))
	F_addReplyErrorObject(m, l0, v244)
	mBase = m.M
	v246 = m.ExcPending
	if v246 != 0 {
		goto L2
	} else {
		goto L67
	}
L9:
	;
	v135 = v31
	goto L38
L10:
	;
	v86 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v88 = *(*int32)(unsafe.Add(mBase, uint32(v86+v42)))
	v89 = F_objectGetVal(m, v88)
	mBase = m.M
	v90 = int32(_a1742)
	v93 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v89))))
	if v93 != 0 {
		goto L26
	} else {
		goto L27
	}
L11:
	;
	if v81-v83 != 0 {
		goto L10
	} else {
		goto L23
	}
L12:
	;
	v81 = F_tolower(m, v77)
	mBase = m.M
	v82 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v78))))
	v83 = F_tolower(m, v82)
	mBase = m.M
	goto L11
L13:
	;
	v51 = v45
	v52 = v46
	v53 = v49
	goto L16
L14:
	;
	v77 = int32(0)
	v78 = v46
	goto L12
L15:
	;
	v77 = v74 & int32(255)
	v78 = v73
	goto L12
L16:
	;
	v55 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v52))))
	if v55 == int32(0) {
		v73 = v52
		v74 = v53
		goto L15
	} else {
		goto L18
	}
L17:
	;
	v73 = v67
	v74 = int32(0)
	goto L15
L18:
	;
	v59 = v53 & int32(255)
	if v59 == v55 {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	v66 = int32(1)
	v67 = v52 + v66
	v68 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v51)+1)))
	if v68 != 0 {
		v51 = v51 + v66
		v52 = v67
		v53 = v68
		goto L16
	} else {
		goto L22
	}
L20:
	;
	v61 = F_tolower(m, v59)
	mBase = m.M
	v62 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v52))))
	v63 = F_tolower(m, v62)
	mBase = m.M
	if v61 == v63 {
		goto L19
	} else {
		goto L21
	}
L21:
	;
	v65 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v51))))
	v73 = v52
	v74 = v65
	goto L15
L22:
	;
	goto L17
L23:
	;
	v130 = int32(0)
	goto L9
L24:
	;
	if v125-v127 != 0 {
		goto L8
	} else {
		goto L36
	}
L25:
	;
	v125 = F_tolower(m, v121)
	mBase = m.M
	v126 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v122))))
	v127 = F_tolower(m, v126)
	mBase = m.M
	goto L24
L26:
	;
	v95 = v89
	v96 = v90
	v97 = v93
	goto L29
L27:
	;
	v121 = int32(0)
	v122 = v90
	goto L25
L28:
	;
	v121 = v118 & int32(255)
	v122 = v117
	goto L25
L29:
	;
	v99 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v96))))
	if v99 == int32(0) {
		v117 = v96
		v118 = v97
		goto L28
	} else {
		goto L31
	}
L30:
	;
	v117 = v111
	v118 = int32(0)
	goto L28
L31:
	;
	v103 = v97 & int32(255)
	if v103 == v99 {
		goto L32
	} else {
		goto L33
	}
L32:
	;
	v110 = int32(1)
	v111 = v96 + v110
	v112 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v95)+1)))
	if v112 != 0 {
		v95 = v95 + v110
		v96 = v111
		v97 = v112
		goto L29
	} else {
		goto L35
	}
L33:
	;
	v105 = F_tolower(m, v103)
	mBase = m.M
	v106 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v96))))
	v107 = F_tolower(m, v106)
	mBase = m.M
	if v105 == v107 {
		goto L32
	} else {
		goto L34
	}
L34:
	;
	v109 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v95))))
	v117 = v96
	v118 = v109
	goto L28
L35:
	;
	goto L30
L36:
	;
	v130 = int32(1)
	goto L9
L37:
	;
	v216 = *(*int32)(unsafe.Add(mBase, uint32(v12)+8))
	if v216 != int32(-1) {
		v222 = v216
		goto L61
	} else {
		goto L62
	}
L38:
	;
	v141 = v135 + int32(2)
	v142 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v142 <= v141 {
		goto L37
	} else {
		goto L40
	}
L40:
	;
	v144 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v148 = *(*int32)(unsafe.Add(mBase, uint32(v144+v141<<(uint(int32(2))%32))))
	v149 = F_objectGetVal(m, v148)
	mBase = m.M
	v150 = *(*int32)(unsafe.Add(mBase, uint32(v12)+8))
	if v150 != int32(-1) {
		goto L42
	} else {
		goto L43
	}
L41:
	;
	v202 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v206 = *(*int32)(unsafe.Add(mBase, uint32(v202+v194<<(uint(int32(2))%32))))
	v212 = F_getRangeLongFromObjectOrReply(m, l0, v206, int32(1), int32(2147483647), v12+int32(8), int32(_a1743))
	mBase = m.M
	v213 = m.ExcPending
	if v213 != 0 {
		goto L2
	} else {
		goto L59
	}
L42:
	;
	v199 = *(*int32)(unsafe.Add(mBase, _consts[86]))
	F_addReplyErrorObject(m, l0, v199)
	mBase = m.M
	v201 = m.ExcPending
	if v201 != 0 {
		goto L2
	} else {
		goto L58
	}
L43:
	;
	v153 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v154 = int32(_a1700)
	v157 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v149))))
	if v157 != 0 {
		goto L46
	} else {
		goto L47
	}
L44:
	;
	if v189-v191 != 0 {
		goto L42
	} else {
		goto L56
	}
L45:
	;
	v189 = F_tolower(m, v185)
	mBase = m.M
	v190 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v186))))
	v191 = F_tolower(m, v190)
	mBase = m.M
	goto L44
L46:
	;
	v159 = v149
	v160 = v154
	v161 = v157
	goto L49
L47:
	;
	v185 = int32(0)
	v186 = v154
	goto L45
L48:
	;
	v185 = v182 & int32(255)
	v186 = v181
	goto L45
L49:
	;
	v163 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v160))))
	if v163 == int32(0) {
		v181 = v160
		v182 = v161
		goto L48
	} else {
		goto L51
	}
L50:
	;
	v181 = v175
	v182 = int32(0)
	goto L48
L51:
	;
	v167 = v161 & int32(255)
	if v167 == v163 {
		goto L52
	} else {
		goto L53
	}
L52:
	;
	v174 = int32(1)
	v175 = v160 + v174
	v176 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v159)+1)))
	if v176 != 0 {
		v159 = v159 + v174
		v160 = v175
		v161 = v176
		goto L49
	} else {
		goto L55
	}
L53:
	;
	v169 = F_tolower(m, v167)
	mBase = m.M
	v170 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v160))))
	v171 = F_tolower(m, v170)
	mBase = m.M
	if v169 == v171 {
		goto L52
	} else {
		goto L54
	}
L54:
	;
	v173 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v159))))
	v181 = v160
	v182 = v173
	goto L48
L55:
	;
	goto L50
L56:
	;
	v194 = v135 + int32(3)
	if v153 != v194 {
		goto L41
	} else {
		goto L57
	}
L57:
	;
	goto L42
L58:
	;
	goto L1
L59:
	;
	if v212 == int32(0) {
		v135 = v141
		goto L38
	} else {
		goto L60
	}
L60:
	;
	goto L1
L61:
	;
	v223 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v228 = v223 + l1<<(uint(int32(2))%32) + int32(4)
	v229 = *(*int32)(unsafe.Add(mBase, uint32(v12)+12))
	if l2 == int32(0) {
		goto L63
	} else {
		goto L64
	}
L62:
	;
	v219 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v12)+8)) = v219
	v222 = v219
	goto L61
L63:
	;
	v237 = int32(1)
	F_genericZpopCommand(m, l0, v228, v229, v130, v237, v222, v237, v237, int32(0))
	mBase = m.M
	v242 = m.ExcPending
	if v242 != 0 {
		goto L2
	} else {
		goto L66
	}
L64:
	;
	v232 = int32(1)
	F_blockingGenericZpopCommand(m, l0, v228, v229, v130, v232, v222, v232, v232)
	mBase = m.M
	v236 = m.ExcPending
	if v236 != 0 {
		goto L2
	} else {
		goto L65
	}
L65:
	;
	goto L1
L66:
	;
	goto L1
L67:
	;
	goto L1
}
func F_zmpopGetKeys(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	v6 = int32(1)
	v9 = F_genericGetKeys(m, int32(0), v6, int32(2), v6, l1, l2, l3)
	v12 = m.ExcPending
	if v12 != 0 {
		return int32(0)
	} else {
		return v9
	}
}
func F_zrangebylexCommand(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v13 int64
	_ = v13
	var v30 int32
	_ = v30
	var v35 int32
	_ = v35
	v2 = int32(0)
	v3 = m.G0
	v5 = v3 - int32(48)
	m.G0 = v5
	*(*int32)(unsafe.Add(mBase, uint32(v5+int32(24)))) = v2
	v13 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v5+int32(16)))) = v13
	*(*int64)(unsafe.Add(mBase, uint32(v5+int32(8)))) = v13
	*(*int64)(unsafe.Add(mBase, uint32(v5))) = v13
	*(*int32)(unsafe.Add(mBase, uint32(v5)+40)) = int32(1101)
	*(*int32)(unsafe.Add(mBase, uint32(v5)+36)) = int32(1102)
	*(*int32)(unsafe.Add(mBase, uint32(v5)+32)) = int32(1103)
	*(*int32)(unsafe.Add(mBase, uint32(v5)+28)) = int32(1104)
	*(*int32)(unsafe.Add(mBase, uint32(v5)+4)) = l0
	v30 = int32(1)
	F_zrangeGenericCommand(m, v5, v30, v2, int32(3), v30)
	mBase = m.M
	v35 = m.ExcPending
	if v35 != 0 {
		return
	} else {
		m.G0 = v5 + int32(48)
		return
	}
}
func F_zrangestoreCommand(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v13 int64
	_ = v13
	var v40 int32
	_ = v40
	v2 = int32(0)
	v5 = m.G0
	v7 = v5 - int32(48)
	m.G0 = v7
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(v9)+4))
	v12 = v7 + int32(8)
	v13 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v12))) = v13
	*(*int32)(unsafe.Add(mBase, uint32(v7+int32(24)))) = v2
	*(*int64)(unsafe.Add(mBase, uint32(v7+int32(16)))) = v13
	*(*int32)(unsafe.Add(mBase, uint32(v12))) = v10
	*(*int64)(unsafe.Add(mBase, uint32(v7))) = v13
	*(*int32)(unsafe.Add(mBase, uint32(v7)+40)) = int32(1097)
	*(*int32)(unsafe.Add(mBase, uint32(v7)+36)) = int32(1098)
	*(*int32)(unsafe.Add(mBase, uint32(v7)+32)) = int32(1099)
	*(*int32)(unsafe.Add(mBase, uint32(v7)+28)) = int32(1100)
	*(*int32)(unsafe.Add(mBase, uint32(v7)+4)) = l0
	F_zrangeGenericCommand(m, v7, int32(2), int32(1), v2, v2)
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
		return
	} else {
		m.G0 = v7 + int32(48)
		return
	}
}
func F_zrealloc_usable(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	*(*int32)(unsafe.Add(mBase, uint32(v8)+12)) = int32(0)
	v14 = F_ztryrealloc_usable_internal(m, l0, l1, v8+int32(12))
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		return int32(0)
	} else {
		v18 = *(*int32)(unsafe.Add(mBase, uint32(v8)+12))
		if l1 == int32(0) {
			if l2 == int32(0) {
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(l2))) = v18
			}
			m.G0 = v8 + int32(16)
			return v14
		} else {
			if v14 != 0 {
				if l2 == int32(0) {
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(l2))) = v18
				}
				m.G0 = v8 + int32(16)
				return v14
			} else {
				v22 = *(*int32)(unsafe.Add(mBase, _consts[1001]))
				m.T0[v22].(func(*base.Module, int32))(m, l1)
				mBase = m.M
				v24 = m.ExcPending
				if v24 != 0 {
					return int32(0)
				} else {
					if l2 == int32(0) {
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(l2))) = v18
					}
					m.G0 = v8 + int32(16)
					return v14
				}
			}
		}
	}
}
func F_zremrangebyrankCommand(m *base.Module, l0 int32) {
	var v4 int32
	_ = v4
	F_zremrangeGenericCommand(m, l0, int32(1))
	v4 = m.ExcPending
	if v4 != 0 {
		return
	} else {
		return
	}
}
func F_zrevrangeCommand(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v13 int64
	_ = v13
	var v30 int32
	_ = v30
	var v35 int32
	_ = v35
	v2 = int32(0)
	v3 = m.G0
	v5 = v3 - int32(48)
	m.G0 = v5
	*(*int32)(unsafe.Add(mBase, uint32(v5+int32(24)))) = v2
	v13 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v5+int32(16)))) = v13
	*(*int64)(unsafe.Add(mBase, uint32(v5+int32(8)))) = v13
	*(*int64)(unsafe.Add(mBase, uint32(v5))) = v13
	*(*int32)(unsafe.Add(mBase, uint32(v5)+40)) = int32(1101)
	*(*int32)(unsafe.Add(mBase, uint32(v5)+36)) = int32(1102)
	*(*int32)(unsafe.Add(mBase, uint32(v5)+32)) = int32(1103)
	*(*int32)(unsafe.Add(mBase, uint32(v5)+28)) = int32(1104)
	*(*int32)(unsafe.Add(mBase, uint32(v5)+4)) = l0
	v30 = int32(1)
	F_zrangeGenericCommand(m, v5, v30, v2, v30, int32(2))
	mBase = m.M
	v35 = m.ExcPending
	if v35 != 0 {
		return
	} else {
		m.G0 = v5 + int32(48)
		return
	}
}
func F_zscanCommand(m *base.Module, l0 int32) {
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
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
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
	var v28 int64
	_ = v28
	var v30 int32
	_ = v30
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v9 = *(*int32)(unsafe.Add(mBase, uint32(v8)+8))
	v10 = F_objectGetVal(m, v9)
	mBase = m.M
	v13 = F_parseScanCursorOrReply(m, l0, v10, v6+int32(8))
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		return
	} else {
		if v13 == int32(-1) {
			m.G0 = v6 + int32(16)
			return
		} else {
			v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
			v18 = *(*int32)(unsafe.Add(mBase, uint32(v17)+4))
			v20 = *(*int32)(unsafe.Add(mBase, _consts[931]))
			v21 = F_lookupKeyReadOrReply(m, l0, v18, v20)
			mBase = m.M
			v22 = m.ExcPending
			if v22 != 0 {
				return
			} else {
				if v21 == int32(0) {
					m.G0 = v6 + int32(16)
					return
				} else {
					v26 = F_checkType(m, l0, v21, int32(3))
					mBase = m.M
					v27 = m.ExcPending
					if v27 != 0 {
						return
					} else {
						if v26 != 0 {
							m.G0 = v6 + int32(16)
							return
						} else {
							v28 = *(*int64)(unsafe.Add(mBase, uint32(v6)+8))
							F_scanGenericCommand(m, l0, v21, v28)
							mBase = m.M
							v30 = m.ExcPending
							if v30 != 0 {
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
func F_zscoreCommand(m *base.Module, l0 int32) {
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
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v33 int32
	_ = v33
	var v34 float64
	_ = v34
	var v36 int32
	_ = v36
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v9 = *(*int32)(unsafe.Add(mBase, uint32(v8)+4))
	v11 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+224)))
	v15 = *(*int32)(unsafe.Add(mBase, uint32(v11<<(uint(int32(2))%32))+uint32(_consts[927])))
	v16 = F_lookupKeyReadOrReply(m, l0, v9, v15)
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		return
	} else {
		if v16 == int32(0) {
			m.G0 = v6 + int32(16)
			return
		} else {
			v21 = F_checkType(m, l0, v16, int32(3))
			mBase = m.M
			v22 = m.ExcPending
			if v22 != 0 {
				return
			} else {
				if v21 != 0 {
					m.G0 = v6 + int32(16)
					return
				} else {
					v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
					v24 = *(*int32)(unsafe.Add(mBase, uint32(v23)+8))
					v25 = F_objectGetVal(m, v24)
					mBase = m.M
					v28 = F_zsetScore(m, v16, v25, v6+int32(8))
					mBase = m.M
					v29 = m.ExcPending
					if v29 != 0 {
						return
					} else {
						if v28 != int32(-1) {
							v34 = *(*float64)(unsafe.Add(mBase, uint32(v6)+8))
							F_addReplyDouble(m, l0, v34)
							mBase = m.M
							v36 = m.ExcPending
							if v36 != 0 {
								return
							} else {
								m.G0 = v6 + int32(16)
								return
							}
						} else {
							F_addReplyNull(m, l0)
							mBase = m.M
							v33 = m.ExcPending
							if v33 != 0 {
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
func F_ztryrealloc_usable(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	*(*int32)(unsafe.Add(mBase, uint32(v7)+12)) = int32(0)
	v13 = F_ztryrealloc_usable_internal(m, l0, l1, v7+int32(12))
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		return int32(0)
	} else {
		if l2 == int32(0) {
		} else {
			v19 = *(*int32)(unsafe.Add(mBase, uint32(v7)+12))
			*(*int32)(unsafe.Add(mBase, uint32(l2))) = v19
		}
		m.G0 = v7 + int32(16)
		return v13
	}
}
func F_ztryrealloc_usable_internal(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v15 int32
	_ = v15
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v79 int32
	_ = v79
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v95 int32
	_ = v95
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
	var v114 int32
	_ = v114
	var v117 int32
	_ = v117
	var v119 int32
	_ = v119
	var v121 int32
	_ = v121
	var v128 int32
	_ = v128
	var v130 int32
	_ = v130
	var v133 int32
	_ = v133
	var v138 int32
	_ = v138
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v149 int32
	_ = v149
	var __phi149 int32
	_ = __phi149
	var v150 int32
	_ = v150
	var __phi150 int32
	_ = __phi150
	var v156 int32
	_ = v156
	var v159 int32
	_ = v159
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v181 int32
	_ = v181
	var v190 int32
	_ = v190
	var v192 int32
	_ = v192
	var v195 int32
	_ = v195
	var v198 int32
	_ = v198
	var v200 int32
	_ = v200
	var v205 int32
	_ = v205
	var v212 int32
	_ = v212
	var v217 int32
	_ = v217
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v232 int32
	_ = v232
	var v240 int32
	_ = v240
	var v242 int32
	_ = v242
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v253 int32
	_ = v253
	var v255 int32
	_ = v255
	var v262 int32
	_ = v262
	var v264 int32
	_ = v264
	var v268 int32
	_ = v268
	var v269 int32
	_ = v269
	var v278 int32
	_ = v278
	var v279 int32
	_ = v279
	var v282 int32
	_ = v282
	var v284 int32
	_ = v284
	var v286 int32
	_ = v286
	var v295 int32
	_ = v295
	var v297 int32
	_ = v297
	var v300 int32
	_ = v300
	var v305 int32
	_ = v305
	var v310 int32
	_ = v310
	var v311 int32
	_ = v311
	var v316 int32
	_ = v316
	var __phi316 int32
	_ = __phi316
	var v317 int32
	_ = v317
	var __phi317 int32
	_ = __phi317
	var v323 int32
	_ = v323
	var v326 int32
	_ = v326
	var v340 int32
	_ = v340
	var v349 int32
	_ = v349
	var v351 int32
	_ = v351
	var v354 int32
	_ = v354
	var v357 int32
	_ = v357
	var v359 int32
	_ = v359
	var v364 int32
	_ = v364
	var v371 int32
	_ = v371
	var v376 int32
	_ = v376
	var v398 int32
	_ = v398
	var v410 int32
	_ = v410
	var v412 int32
	_ = v412
	var v414 int32
	_ = v414
	var v418 int32
	_ = v418
	var v423 int32
	_ = v423
	var v424 int32
	_ = v424
	var v435 int32
	_ = v435
	var v438 int32
	_ = v438
	var v445 int32
	_ = v445
	var v450 int32
	_ = v450
	var v454 int32
	_ = v454
	var v456 int32
	_ = v456
	var v471 int32
	_ = v471
	var v473 int32
	_ = v473
	var v476 int32
	_ = v476
	var v479 int32
	_ = v479
	var v483 int32
	_ = v483
	var v493 int32
	_ = v493
	var v494 int32
	_ = v494
	var v500 int32
	_ = v500
	var v502 int32
	_ = v502
	var v505 int32
	_ = v505
	var v509 int32
	_ = v509
	var v515 int32
	_ = v515
	var v517 int32
	_ = v517
	var v518 int32
	_ = v518
	var v519 int32
	_ = v519
	var v520 int32
	_ = v520
	var v529 int32
	_ = v529
	var v531 int32
	_ = v531
	var v532 int32
	_ = v532
	var v533 int32
	_ = v533
	var v535 int32
	_ = v535
	var v557 int32
	_ = v557
	var v562 int32
	_ = v562
	var v564 int32
	_ = v564
	var v565 int32
	_ = v565
	var v566 int32
	_ = v566
	var v570 int32
	_ = v570
	var v573 int32
	_ = v573
	var v575 int32
	_ = v575
	var v581 int32
	_ = v581
	var v584 int32
	_ = v584
	var v586 int32
	_ = v586
	var v590 int32
	_ = v590
	var v593 int32
	_ = v593
	var v599 int32
	_ = v599
	var v601 int32
	_ = v601
	var v608 int32
	_ = v608
	var v609 int32
	_ = v609
	var v616 int32
	_ = v616
	var v622 int32
	_ = v622
	var v625 int32
	_ = v625
	var v627 int32
	_ = v627
	var v633 int32
	_ = v633
	var v635 int32
	_ = v635
	var v638 int32
	_ = v638
	var v640 int32
	_ = v640
	var v644 int32
	_ = v644
	var v647 int32
	_ = v647
	var v651 int32
	_ = v651
	var v654 int32
	_ = v654
	var v656 int32
	_ = v656
	var v662 int32
	_ = v662
	var v664 int32
	_ = v664
	var v667 int32
	_ = v667
	var v669 int32
	_ = v669
	var v673 int32
	_ = v673
	var v676 int32
	_ = v676
	var v681 int32
	_ = v681
	var v692 int32
	_ = v692
	var v693 int32
	_ = v693
	var v696 int32
	_ = v696
	var v698 int32
	_ = v698
	var v699 int32
	_ = v699
	var v706 int32
	_ = v706
	var v707 int32
	_ = v707
	var v709 int32
	_ = v709
	var v711 int32
	_ = v711
	var v713 int32
	_ = v713
	var v715 int32
	_ = v715
	var v718 int32
	_ = v718
	var v720 int32
	_ = v720
	var v722 int32
	_ = v722
	var v729 int32
	_ = v729
	var v731 int32
	_ = v731
	var v734 int32
	_ = v734
	var v739 int32
	_ = v739
	var v744 int32
	_ = v744
	var v745 int32
	_ = v745
	var v750 int32
	_ = v750
	var __phi750 int32
	_ = __phi750
	var v751 int32
	_ = v751
	var __phi751 int32
	_ = __phi751
	var v757 int32
	_ = v757
	var v760 int32
	_ = v760
	var v763 int32
	_ = v763
	var v764 int32
	_ = v764
	var v782 int32
	_ = v782
	var v791 int32
	_ = v791
	var v793 int32
	_ = v793
	var v796 int32
	_ = v796
	var v799 int32
	_ = v799
	var v801 int32
	_ = v801
	var v806 int32
	_ = v806
	var v813 int32
	_ = v813
	var v818 int32
	_ = v818
	var v823 int32
	_ = v823
	var v824 int32
	_ = v824
	var v833 int32
	_ = v833
	var v841 int32
	_ = v841
	var v843 int32
	_ = v843
	var v847 int32
	_ = v847
	var v848 int32
	_ = v848
	var v854 int32
	_ = v854
	var v856 int32
	_ = v856
	var v863 int32
	_ = v863
	var v865 int32
	_ = v865
	var v869 int32
	_ = v869
	var v870 int32
	_ = v870
	var v879 int32
	_ = v879
	var v880 int32
	_ = v880
	var v883 int32
	_ = v883
	var v885 int32
	_ = v885
	var v887 int32
	_ = v887
	var v896 int32
	_ = v896
	var v898 int32
	_ = v898
	var v901 int32
	_ = v901
	var v906 int32
	_ = v906
	var v911 int32
	_ = v911
	var v912 int32
	_ = v912
	var v917 int32
	_ = v917
	var __phi917 int32
	_ = __phi917
	var v918 int32
	_ = v918
	var __phi918 int32
	_ = __phi918
	var v924 int32
	_ = v924
	var v927 int32
	_ = v927
	var v941 int32
	_ = v941
	var v950 int32
	_ = v950
	var v952 int32
	_ = v952
	var v955 int32
	_ = v955
	var v958 int32
	_ = v958
	var v960 int32
	_ = v960
	var v965 int32
	_ = v965
	var v972 int32
	_ = v972
	var v977 int32
	_ = v977
	var v999 int32
	_ = v999
	var v1011 int32
	_ = v1011
	var v1013 int32
	_ = v1013
	var v1015 int32
	_ = v1015
	var v1019 int32
	_ = v1019
	var v1024 int32
	_ = v1024
	var v1025 int32
	_ = v1025
	var v1036 int32
	_ = v1036
	var v1039 int32
	_ = v1039
	var v1046 int32
	_ = v1046
	var v1051 int32
	_ = v1051
	var v1055 int32
	_ = v1055
	var v1057 int32
	_ = v1057
	var v1072 int32
	_ = v1072
	var v1074 int32
	_ = v1074
	var v1077 int32
	_ = v1077
	var v1080 int32
	_ = v1080
	var v1084 int32
	_ = v1084
	var v1094 int32
	_ = v1094
	var v1095 int32
	_ = v1095
	var v1101 int32
	_ = v1101
	var v1103 int32
	_ = v1103
	var v1106 int32
	_ = v1106
	var v1110 int32
	_ = v1110
	var v1116 int32
	_ = v1116
	var v1118 int32
	_ = v1118
	var v1119 int32
	_ = v1119
	var v1120 int32
	_ = v1120
	var v1121 int32
	_ = v1121
	var v1130 int32
	_ = v1130
	var v1132 int32
	_ = v1132
	var v1133 int32
	_ = v1133
	var v1134 int32
	_ = v1134
	var v1136 int32
	_ = v1136
	var v1159 int32
	_ = v1159
	var v1162 int32
	_ = v1162
	var v1165 int32
	_ = v1165
	var v1178 int32
	_ = v1178
	var v1179 int32
	_ = v1179
	var v1184 int32
	_ = v1184
	var v1190 int32
	_ = v1190
	var v1193 int32
	_ = v1193
	var v1196 int32
	_ = v1196
	var v1198 int32
	_ = v1198
	var v1199 int32
	_ = v1199
	var v1203 int32
	_ = v1203
	var v1210 int32
	_ = v1210
	var v1216 int32
	_ = v1216
	var v1219 int32
	_ = v1219
	var v1221 int32
	_ = v1221
	var v1227 int32
	_ = v1227
	var v1230 int32
	_ = v1230
	var v1232 int32
	_ = v1232
	var v1236 int32
	_ = v1236
	var v1239 int32
	_ = v1239
	var v1243 int32
	_ = v1243
	var v1246 int32
	_ = v1246
	var v1248 int32
	_ = v1248
	var v1254 int32
	_ = v1254
	var v1259 int32
	_ = v1259
	var v1261 int32
	_ = v1261
	var v1265 int32
	_ = v1265
	var v1268 int32
	_ = v1268
	var v1278 int32
	_ = v1278
	var v1289 int32
	_ = v1289
	var v1295 int32
	_ = v1295
	if l0 == int32(0) {
		goto L6
	} else {
		goto L7
	}
L1:
	;
	F__serverAssert(m, int32(_a1860), int32(_a1879), int32(463))
	mBase = m.M
	v1295 = m.ExcPending
	if v1295 != 0 {
		goto L284
	} else {
		goto L286
	}
L2:
	;
	F__serverAssert(m, int32(_a1860), int32(_a1879), int32(463))
	mBase = m.M
	v1289 = m.ExcPending
	if v1289 != 0 {
		goto L284
	} else {
		goto L285
	}
L3:
	;
	return v1278
L4:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1203))) = l1
	v1216 = *(*int32)(unsafe.Add(mBase, _consts[411]))
	if v1216 != int32(-1) {
		v1227 = v1216
		goto L272
	} else {
		goto L273
	}
L5:
	;
	v1210 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v1210
	return v1210
L6:
	;
	if l0 != 0 {
		goto L122
	} else {
		goto L123
	}
L7:
	;
	if l1 != 0 {
		goto L6
	} else {
		goto L8
	}
L8:
	;
	v9 = l0 + int32(-8)
	v10 = *(*int32)(unsafe.Add(mBase, uint32(v9)))
	if int32(-1) < v10 {
		goto L10
	} else {
		goto L11
	}
L9:
	;
	if v79 == int32(0) {
		goto L22
	} else {
		goto L23
	}
L10:
	;
	v50 = *(*int32)(unsafe.Add(mBase, _consts[411]))
	if v50 != int32(-1) {
		v61 = v50
		goto L17
	} else {
		goto L18
	}
L11:
	;
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0+int32(-12))))
	if v15 == int32(0) {
		goto L2
	} else {
		goto L12
	}
L12:
	;
	v21 = *(*int32)(unsafe.Add(mBase, _consts[411]))
	if v21 != int32(-1) {
		v32 = v21
		goto L13
	} else {
		goto L14
	}
L13:
	;
	v34 = v10&int32(2147483647) + int32(8)
	if v32 < int32(260) {
		goto L15
	} else {
		goto L16
	}
L14:
	;
	v24 = int32(0)
	v26 = *(*int32)(unsafe.Add(mBase, _consts[281]))
	*(*int32)(unsafe.Add(mBase, _consts[411])) = v26
	*(*int32)(unsafe.Add(mBase, _consts[281])) = v26 + int32(1)
	v32 = v26
	goto L13
L15:
	;
	v43 = v32 << (uint(int32(2)) % 32)
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v43)+uint32(_consts[285])))
	*(*int32)(unsafe.Add(mBase, uint32(v43)+uint32(_consts[285]))) = v46 - v34
	v79 = v15
	goto L9
L16:
	;
	v37 = int32(0)
	v39 = *(*int32)(unsafe.Add(mBase, _consts[286]))
	*(*int32)(unsafe.Add(mBase, _consts[286])) = v39 - v34
	v79 = v15
	goto L9
L17:
	;
	v63 = v10 + int32(8)
	if v61 < int32(260) {
		goto L19
	} else {
		goto L20
	}
L18:
	;
	v53 = int32(0)
	v55 = *(*int32)(unsafe.Add(mBase, _consts[281]))
	*(*int32)(unsafe.Add(mBase, _consts[411])) = v55
	*(*int32)(unsafe.Add(mBase, _consts[281])) = v55 + int32(1)
	v61 = v55
	goto L17
L19:
	;
	v72 = v61 << (uint(int32(2)) % 32)
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v72)+uint32(_consts[285])))
	*(*int32)(unsafe.Add(mBase, uint32(v72)+uint32(_consts[285]))) = v75 - v63
	v79 = v9
	goto L9
L20:
	;
	v66 = int32(0)
	v68 = *(*int32)(unsafe.Add(mBase, _consts[286]))
	*(*int32)(unsafe.Add(mBase, _consts[286])) = v68 - v63
	v79 = v9
	goto L9
L21:
	;
	if l2 != 0 {
		goto L5
	} else {
		goto L121
	}
L22:
	;
	goto L21
L23:
	;
	v91 = int32(-8)
	v92 = v79 + v91
	v95 = *(*int32)(unsafe.Add(mBase, uint32(v79+int32(-4))))
	v97 = v95 & v91
	v98 = v92 + v97
	if v95&int32(1) != 0 {
		v222 = v97
		v223 = v92
		goto L24
	} else {
		goto L25
	}
L24:
	;
	if base.Ui32(v98) <= base.Ui32(v223) {
		goto L22
	} else {
		goto L59
	}
L25:
	;
	if v95&int32(2) == int32(0) {
		goto L22
	} else {
		goto L26
	}
L26:
	;
	v105 = *(*int32)(unsafe.Add(mBase, uint32(v92)))
	v106 = v92 - v105
	v108 = *(*int32)(unsafe.Add(mBase, _consts[515]))
	if base.Ui32(v106) < base.Ui32(v108) {
		goto L22
	} else {
		goto L27
	}
L27:
	;
	v110 = v105 + v97
	v112 = *(*int32)(unsafe.Add(mBase, _consts[516]))
	if v106 == v112 {
		goto L31
	} else {
		goto L32
	}
L28:
	;
	if v128 == int32(0) {
		v222 = v110
		v223 = v106
		goto L24
	} else {
		goto L47
	}
L29:
	;
	v181 = int32(0)
	goto L28
L30:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v117)+12)) = v114
	*(*int32)(unsafe.Add(mBase, uint32(v114)+8)) = v117
	v222 = v110
	v223 = v106
	goto L24
L31:
	;
	v162 = *(*int32)(unsafe.Add(mBase, uint32(v98)+4))
	v163 = int32(3)
	if v162&v163 != v163 {
		v222 = v110
		v223 = v106
		goto L24
	} else {
		goto L46
	}
L32:
	;
	v114 = *(*int32)(unsafe.Add(mBase, uint32(v106)+12))
	if base.Ui32(int32(255)) < base.Ui32(v105) {
		goto L33
	} else {
		goto L34
	}
L33:
	;
	v128 = *(*int32)(unsafe.Add(mBase, uint32(v106)+24))
	if v114 == v106 {
		goto L36
	} else {
		goto L37
	}
L34:
	;
	v117 = *(*int32)(unsafe.Add(mBase, uint32(v106)+8))
	if v114 != v117 {
		goto L30
	} else {
		goto L35
	}
L35:
	;
	v119 = int32(0)
	v121 = *(*int32)(unsafe.Add(mBase, _consts[517]))
	*(*int32)(unsafe.Add(mBase, _consts[517])) = v121 & base.I32_rotl(int32(-2), int32(base.Ui32(v105)>>(uint(int32(3))%32)))
	v222 = v110
	v223 = v106
	goto L24
L36:
	;
	v133 = *(*int32)(unsafe.Add(mBase, uint32(v106)+20))
	if v133 == int32(0) {
		goto L39
	} else {
		goto L40
	}
L37:
	;
	v130 = *(*int32)(unsafe.Add(mBase, uint32(v106)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v130)+12)) = v114
	*(*int32)(unsafe.Add(mBase, uint32(v114)+8)) = v130
	v181 = v114
	goto L28
L38:
	;
	__phi149 = v143
	__phi150 = v144
	v149 = __phi149
	v150 = __phi150
	goto L42
L39:
	;
	v138 = *(*int32)(unsafe.Add(mBase, uint32(v106)+16))
	if v138 == int32(0) {
		goto L29
	} else {
		goto L41
	}
L40:
	;
	v143 = v133
	v144 = v106 + int32(20)
	goto L38
L41:
	;
	v143 = v138
	v144 = v106 + int32(16)
	goto L38
L42:
	;
	v156 = *(*int32)(unsafe.Add(mBase, uint32(v149)+20))
	if v156 != 0 {
		__phi149 = v156
		__phi150 = v149 + int32(20)
		v149 = __phi149
		v150 = __phi150
		goto L42
	} else {
		goto L44
	}
L43:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v150))) = int32(0)
	v181 = v149
	goto L28
L44:
	;
	v159 = *(*int32)(unsafe.Add(mBase, uint32(v149)+16))
	if v159 != 0 {
		__phi149 = v159
		__phi150 = v149 + int32(16)
		v149 = __phi149
		v150 = __phi150
		goto L42
	} else {
		goto L45
	}
L45:
	;
	goto L43
L46:
	;
	*(*int32)(unsafe.Add(mBase, _consts[518])) = v110
	*(*int32)(unsafe.Add(mBase, uint32(v98)+4)) = v162 & int32(-2)
	*(*int32)(unsafe.Add(mBase, uint32(v106)+4)) = v110 | int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v98))) = v110
	goto L21
L47:
	;
	v190 = *(*int32)(unsafe.Add(mBase, uint32(v106)+28))
	v192 = v190 << (uint(int32(2)) % 32)
	v195 = *(*int32)(unsafe.Add(mBase, uint32(v192)+uint32(_consts[519])))
	if v106 != v195 {
		goto L49
	} else {
		goto L50
	}
L48:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v181)+24)) = v128
	v212 = *(*int32)(unsafe.Add(mBase, uint32(v106)+16))
	if v212 == int32(0) {
		goto L56
	} else {
		goto L57
	}
L49:
	;
	v205 = *(*int32)(unsafe.Add(mBase, uint32(v128)+16))
	if v205 != v106 {
		goto L53
	} else {
		goto L54
	}
L50:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v192)+uint32(_consts[519]))) = v181
	if v181 != 0 {
		goto L48
	} else {
		goto L51
	}
L51:
	;
	v198 = int32(0)
	v200 = *(*int32)(unsafe.Add(mBase, _consts[520]))
	*(*int32)(unsafe.Add(mBase, _consts[520])) = v200 & base.I32_rotl(int32(-2), v190)
	v222 = v110
	v223 = v106
	goto L24
L52:
	;
	if v181 == int32(0) {
		v222 = v110
		v223 = v106
		goto L24
	} else {
		goto L55
	}
L53:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v128)+20)) = v181
	goto L52
L54:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v128)+16)) = v181
	goto L52
L55:
	;
	goto L48
L56:
	;
	v217 = *(*int32)(unsafe.Add(mBase, uint32(v106)+20))
	if v217 == int32(0) {
		v222 = v110
		v223 = v106
		goto L24
	} else {
		goto L58
	}
L57:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v181)+16)) = v212
	*(*int32)(unsafe.Add(mBase, uint32(v212)+24)) = v181
	goto L56
L58:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v181)+20)) = v217
	*(*int32)(unsafe.Add(mBase, uint32(v217)+24)) = v181
	v222 = v110
	v223 = v106
	goto L24
L59:
	;
	v232 = *(*int32)(unsafe.Add(mBase, uint32(v98)+4))
	if v232&int32(1) == int32(0) {
		goto L22
	} else {
		goto L60
	}
L60:
	;
	if v232&int32(2) != 0 {
		goto L65
	} else {
		goto L66
	}
L61:
	;
	if base.Ui32(int32(255)) < base.Ui32(v398) {
		goto L99
	} else {
		goto L100
	}
L62:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v223)+4)) = v278 | int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v223+v278))) = v278
	if v223 != v262 {
		v398 = v278
		goto L61
	} else {
		goto L98
	}
L63:
	;
	if v295 == int32(0) {
		goto L62
	} else {
		goto L86
	}
L64:
	;
	v340 = int32(0)
	goto L63
L65:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v98)+4)) = v232 & int32(-2)
	*(*int32)(unsafe.Add(mBase, uint32(v223)+4)) = v222 | int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v223+v222))) = v222
	v398 = v222
	goto L61
L66:
	;
	v240 = *(*int32)(unsafe.Add(mBase, _consts[521]))
	if v98 != v240 {
		goto L67
	} else {
		goto L68
	}
L67:
	;
	v262 = *(*int32)(unsafe.Add(mBase, _consts[516]))
	if v98 != v262 {
		goto L70
	} else {
		goto L71
	}
L68:
	;
	v242 = int32(0)
	*(*int32)(unsafe.Add(mBase, _consts[521])) = v223
	v246 = *(*int32)(unsafe.Add(mBase, _consts[522]))
	v247 = v246 + v222
	*(*int32)(unsafe.Add(mBase, _consts[522])) = v247
	*(*int32)(unsafe.Add(mBase, uint32(v223)+4)) = v247 | int32(1)
	v253 = *(*int32)(unsafe.Add(mBase, _consts[516]))
	if v223 != v253 {
		goto L22
	} else {
		goto L69
	}
L69:
	;
	v255 = int32(0)
	*(*int32)(unsafe.Add(mBase, _consts[518])) = v255
	*(*int32)(unsafe.Add(mBase, _consts[516])) = v255
	goto L21
L70:
	;
	v278 = v232&int32(-8) + v222
	v279 = *(*int32)(unsafe.Add(mBase, uint32(v98)+12))
	if base.Ui32(int32(255)) < base.Ui32(v232) {
		goto L72
	} else {
		goto L73
	}
L71:
	;
	v264 = int32(0)
	*(*int32)(unsafe.Add(mBase, _consts[516])) = v223
	v268 = *(*int32)(unsafe.Add(mBase, _consts[518]))
	v269 = v268 + v222
	*(*int32)(unsafe.Add(mBase, _consts[518])) = v269
	*(*int32)(unsafe.Add(mBase, uint32(v223)+4)) = v269 | int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v223+v269))) = v269
	goto L21
L72:
	;
	v295 = *(*int32)(unsafe.Add(mBase, uint32(v98)+24))
	if v279 == v98 {
		goto L76
	} else {
		goto L77
	}
L73:
	;
	v282 = *(*int32)(unsafe.Add(mBase, uint32(v98)+8))
	if v279 != v282 {
		goto L74
	} else {
		goto L75
	}
L74:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v282)+12)) = v279
	*(*int32)(unsafe.Add(mBase, uint32(v279)+8)) = v282
	goto L62
L75:
	;
	v284 = int32(0)
	v286 = *(*int32)(unsafe.Add(mBase, _consts[517]))
	*(*int32)(unsafe.Add(mBase, _consts[517])) = v286 & base.I32_rotl(int32(-2), int32(base.Ui32(v232)>>(uint(int32(3))%32)))
	goto L62
L76:
	;
	v300 = *(*int32)(unsafe.Add(mBase, uint32(v98)+20))
	if v300 == int32(0) {
		goto L79
	} else {
		goto L80
	}
L77:
	;
	v297 = *(*int32)(unsafe.Add(mBase, uint32(v98)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v297)+12)) = v279
	*(*int32)(unsafe.Add(mBase, uint32(v279)+8)) = v297
	v340 = v279
	goto L63
L78:
	;
	__phi316 = v310
	__phi317 = v311
	v316 = __phi316
	v317 = __phi317
	goto L82
L79:
	;
	v305 = *(*int32)(unsafe.Add(mBase, uint32(v98)+16))
	if v305 == int32(0) {
		goto L64
	} else {
		goto L81
	}
L80:
	;
	v310 = v300
	v311 = v98 + int32(20)
	goto L78
L81:
	;
	v310 = v305
	v311 = v98 + int32(16)
	goto L78
L82:
	;
	v323 = *(*int32)(unsafe.Add(mBase, uint32(v316)+20))
	if v323 != 0 {
		__phi316 = v323
		__phi317 = v316 + int32(20)
		v316 = __phi316
		v317 = __phi317
		goto L82
	} else {
		goto L84
	}
L83:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v317))) = int32(0)
	v340 = v316
	goto L63
L84:
	;
	v326 = *(*int32)(unsafe.Add(mBase, uint32(v316)+16))
	if v326 != 0 {
		__phi316 = v326
		__phi317 = v316 + int32(16)
		v316 = __phi316
		v317 = __phi317
		goto L82
	} else {
		goto L85
	}
L85:
	;
	goto L83
L86:
	;
	v349 = *(*int32)(unsafe.Add(mBase, uint32(v98)+28))
	v351 = v349 << (uint(int32(2)) % 32)
	v354 = *(*int32)(unsafe.Add(mBase, uint32(v351)+uint32(_consts[519])))
	if v98 != v354 {
		goto L88
	} else {
		goto L89
	}
L87:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v340)+24)) = v295
	v371 = *(*int32)(unsafe.Add(mBase, uint32(v98)+16))
	if v371 == int32(0) {
		goto L95
	} else {
		goto L96
	}
L88:
	;
	v364 = *(*int32)(unsafe.Add(mBase, uint32(v295)+16))
	if v364 != v98 {
		goto L92
	} else {
		goto L93
	}
L89:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v351)+uint32(_consts[519]))) = v340
	if v340 != 0 {
		goto L87
	} else {
		goto L90
	}
L90:
	;
	v357 = int32(0)
	v359 = *(*int32)(unsafe.Add(mBase, _consts[520]))
	*(*int32)(unsafe.Add(mBase, _consts[520])) = v359 & base.I32_rotl(int32(-2), v349)
	goto L62
L91:
	;
	if v340 == int32(0) {
		goto L62
	} else {
		goto L94
	}
L92:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v295)+20)) = v340
	goto L91
L93:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v295)+16)) = v340
	goto L91
L94:
	;
	goto L87
L95:
	;
	v376 = *(*int32)(unsafe.Add(mBase, uint32(v98)+20))
	if v376 == int32(0) {
		goto L62
	} else {
		goto L97
	}
L96:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v340)+16)) = v371
	*(*int32)(unsafe.Add(mBase, uint32(v371)+24)) = v340
	goto L95
L97:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v340)+20)) = v376
	*(*int32)(unsafe.Add(mBase, uint32(v376)+24)) = v340
	goto L62
L98:
	;
	*(*int32)(unsafe.Add(mBase, _consts[518])) = v278
	goto L21
L99:
	;
	if base.Ui32(int32(16777215)) < base.Ui32(v398) {
		v445 = int32(31)
		goto L104
	} else {
		goto L105
	}
L100:
	;
	v410 = v398 & int32(-8)
	v412 = v410 + int32(9128464)
	v414 = *(*int32)(unsafe.Add(mBase, _consts[517]))
	v418 = int32(1) << (uint(int32(base.Ui32(v398)>>(uint(int32(3))%32))) % 32)
	if v414&v418 != 0 {
		goto L102
	} else {
		goto L103
	}
L101:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v410)+uint32(_consts[523]))) = v223
	*(*int32)(unsafe.Add(mBase, uint32(v424)+12)) = v223
	*(*int32)(unsafe.Add(mBase, uint32(v223)+12)) = v412
	*(*int32)(unsafe.Add(mBase, uint32(v223)+8)) = v424
	goto L21
L102:
	;
	v423 = *(*int32)(unsafe.Add(mBase, uint32(v410)+uint32(_consts[523])))
	v424 = v423
	goto L101
L103:
	;
	*(*int32)(unsafe.Add(mBase, _consts[517])) = v414 | v418
	v424 = v412
	goto L101
L104:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v223)+28)) = v445
	*(*int64)(unsafe.Add(mBase, uint32(v223)+16)) = int64(0)
	v450 = v445 << (uint(int32(2)) % 32)
	v454 = *(*int32)(unsafe.Add(mBase, _consts[520]))
	v456 = int32(1) << (uint(v445) % 32)
	if v454&v456 != 0 {
		goto L109
	} else {
		goto L110
	}
L105:
	;
	v435 = base.I32_clz(int32(base.Ui32(v398) >> (uint(int32(8)) % 32)))
	v438 = int32(1)
	v445 = int32(base.Ui32(v398)>>(uint(int32(38)-v435)%32))&v438 - v435<<(uint(v438)%32) + int32(62)
	goto L104
L106:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v223+v517))) = v520
	*(*int32)(unsafe.Add(mBase, uint32(v223)+12)) = v519
	*(*int32)(unsafe.Add(mBase, uint32(v223+v515))) = v518
	v529 = int32(0)
	v531 = *(*int32)(unsafe.Add(mBase, _consts[524]))
	v532 = int32(-1)
	v533 = v531 + v532
	if v533 != 0 {
		goto L118
	} else {
		goto L119
	}
L107:
	;
	v509 = *(*int32)(unsafe.Add(mBase, uint32(v479)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v509)+12)) = v223
	*(*int32)(unsafe.Add(mBase, uint32(v479)+8)) = v223
	v515 = int32(24)
	v517 = int32(8)
	v518 = int32(0)
	v519 = v479
	v520 = v509
	goto L106
L108:
	;
	v515 = v500
	v517 = v502
	v518 = v223
	v519 = v223
	v520 = v505
	goto L106
L109:
	;
	if v445 == int32(31) {
		goto L111
	} else {
		goto L112
	}
L110:
	;
	*(*int32)(unsafe.Add(mBase, _consts[520])) = v454 | v456
	*(*int32)(unsafe.Add(mBase, uint32(v450)+uint32(_consts[519]))) = v223
	v500 = int32(8)
	v502 = int32(24)
	v505 = v450 + int32(9128728)
	goto L108
L111:
	;
	v471 = int32(0)
	goto L113
L112:
	;
	v471 = int32(25) - int32(base.Ui32(v445)>>(uint(int32(1))%32))
	goto L113
L113:
	;
	v473 = *(*int32)(unsafe.Add(mBase, uint32(v450)+uint32(_consts[519])))
	v476 = v398 << (uint(v471) % 32)
	v479 = v473
	goto L114
L114:
	;
	v483 = *(*int32)(unsafe.Add(mBase, uint32(v479)+4))
	if v483&int32(-8) == v398 {
		goto L107
	} else {
		goto L116
	}
L115:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v493+int32(16)))) = v223
	v500 = int32(8)
	v502 = int32(24)
	v505 = v479
	goto L108
L116:
	;
	v493 = v479 + int32(base.Ui32(v476)>>(uint(int32(29))%32))&int32(4)
	v494 = *(*int32)(unsafe.Add(mBase, uint32(v493)+16))
	if v494 != 0 {
		v476 = v476 << (uint(int32(1)) % 32)
		v479 = v494
		goto L114
	} else {
		goto L117
	}
L117:
	;
	goto L115
L118:
	;
	v535 = v533
	goto L120
L119:
	;
	v535 = v532
	goto L120
L120:
	;
	*(*int32)(unsafe.Add(mBase, _consts[524])) = v535
	goto L22
L121:
	;
	return int32(0)
L122:
	;
	v608 = l0 + int32(-8)
	v609 = *(*int32)(unsafe.Add(mBase, uint32(v608)))
	if base.Ui32(l1) < base.Ui32(int32(2147483647)) {
		goto L137
	} else {
		goto L138
	}
L123:
	;
	v557 = int32(0)
	if base.Ui32(int32(2147483646)) < base.Ui32(l1) {
		v599 = v557
		v601 = v557
		goto L124
	} else {
		goto L125
	}
L124:
	;
	if l2 == int32(0) {
		v1278 = v599
		goto L3
	} else {
		goto L136
	}
L125:
	;
	if l1 != 0 {
		goto L127
	} else {
		goto L128
	}
L126:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v565))) = v562
	v570 = *(*int32)(unsafe.Add(mBase, _consts[411]))
	if v570 != int32(-1) {
		v581 = v570
		goto L131
	} else {
		goto L132
	}
L127:
	;
	v562 = l1
	goto L129
L128:
	;
	v562 = int32(4)
	goto L129
L129:
	;
	v564 = v562 + int32(8)
	v565 = F_emscripten_builtin_malloc(m, v564)
	mBase = m.M
	if v565 != 0 {
		goto L126
	} else {
		goto L130
	}
L130:
	;
	v566 = int32(0)
	v599 = v566
	v601 = v566
	goto L124
L131:
	;
	if v581 < int32(260) {
		goto L134
	} else {
		goto L135
	}
L132:
	;
	v573 = int32(0)
	v575 = *(*int32)(unsafe.Add(mBase, _consts[281]))
	*(*int32)(unsafe.Add(mBase, _consts[411])) = v575
	*(*int32)(unsafe.Add(mBase, _consts[281])) = v575 + int32(1)
	v581 = v575
	goto L131
L133:
	;
	v599 = v565 + int32(8)
	v601 = v562
	goto L124
L134:
	;
	v590 = v581 << (uint(int32(2)) % 32)
	v593 = *(*int32)(unsafe.Add(mBase, uint32(v590)+uint32(_consts[285])))
	*(*int32)(unsafe.Add(mBase, uint32(v590)+uint32(_consts[285]))) = v593 + v564
	goto L133
L135:
	;
	v584 = int32(0)
	v586 = *(*int32)(unsafe.Add(mBase, _consts[286]))
	*(*int32)(unsafe.Add(mBase, _consts[286])) = v586 + v564
	goto L133
L136:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v601
	return v599
L137:
	;
	v1159 = l1 + int32(8)
	if v608 != 0 {
		goto L253
	} else {
		goto L254
	}
L138:
	;
	if int32(-1) < v609 {
		goto L140
	} else {
		goto L141
	}
L139:
	;
	if v681 == int32(0) {
		goto L152
	} else {
		goto L153
	}
L140:
	;
	v651 = *(*int32)(unsafe.Add(mBase, _consts[411]))
	if v651 != int32(-1) {
		v662 = v651
		goto L147
	} else {
		goto L148
	}
L141:
	;
	v616 = *(*int32)(unsafe.Add(mBase, uint32(l0+int32(-12))))
	if v616 == int32(0) {
		goto L1
	} else {
		goto L142
	}
L142:
	;
	v622 = *(*int32)(unsafe.Add(mBase, _consts[411]))
	if v622 != int32(-1) {
		v633 = v622
		goto L143
	} else {
		goto L144
	}
L143:
	;
	v635 = v609&int32(2147483647) + int32(8)
	if v633 < int32(260) {
		goto L145
	} else {
		goto L146
	}
L144:
	;
	v625 = int32(0)
	v627 = *(*int32)(unsafe.Add(mBase, _consts[281]))
	*(*int32)(unsafe.Add(mBase, _consts[411])) = v627
	*(*int32)(unsafe.Add(mBase, _consts[281])) = v627 + int32(1)
	v633 = v627
	goto L143
L145:
	;
	v644 = v633 << (uint(int32(2)) % 32)
	v647 = *(*int32)(unsafe.Add(mBase, uint32(v644)+uint32(_consts[285])))
	*(*int32)(unsafe.Add(mBase, uint32(v644)+uint32(_consts[285]))) = v647 - v635
	v681 = v616
	goto L139
L146:
	;
	v638 = int32(0)
	v640 = *(*int32)(unsafe.Add(mBase, _consts[286]))
	*(*int32)(unsafe.Add(mBase, _consts[286])) = v640 - v635
	v681 = v616
	goto L139
L147:
	;
	v664 = v609 + int32(8)
	if v662 < int32(260) {
		goto L149
	} else {
		goto L150
	}
L148:
	;
	v654 = int32(0)
	v656 = *(*int32)(unsafe.Add(mBase, _consts[281]))
	*(*int32)(unsafe.Add(mBase, _consts[411])) = v656
	*(*int32)(unsafe.Add(mBase, _consts[281])) = v656 + int32(1)
	v662 = v656
	goto L147
L149:
	;
	v673 = v662 << (uint(int32(2)) % 32)
	v676 = *(*int32)(unsafe.Add(mBase, uint32(v673)+uint32(_consts[285])))
	*(*int32)(unsafe.Add(mBase, uint32(v673)+uint32(_consts[285]))) = v676 - v664
	v681 = v608
	goto L139
L150:
	;
	v667 = int32(0)
	v669 = *(*int32)(unsafe.Add(mBase, _consts[286]))
	*(*int32)(unsafe.Add(mBase, _consts[286])) = v669 - v664
	v681 = v608
	goto L139
L151:
	;
	if l2 != 0 {
		goto L5
	} else {
		goto L251
	}
L152:
	;
	goto L151
L153:
	;
	v692 = int32(-8)
	v693 = v681 + v692
	v696 = *(*int32)(unsafe.Add(mBase, uint32(v681+int32(-4))))
	v698 = v696 & v692
	v699 = v693 + v698
	if v696&int32(1) != 0 {
		v823 = v698
		v824 = v693
		goto L154
	} else {
		goto L155
	}
L154:
	;
	if base.Ui32(v699) <= base.Ui32(v824) {
		goto L152
	} else {
		goto L189
	}
L155:
	;
	if v696&int32(2) == int32(0) {
		goto L152
	} else {
		goto L156
	}
L156:
	;
	v706 = *(*int32)(unsafe.Add(mBase, uint32(v693)))
	v707 = v693 - v706
	v709 = *(*int32)(unsafe.Add(mBase, _consts[515]))
	if base.Ui32(v707) < base.Ui32(v709) {
		goto L152
	} else {
		goto L157
	}
L157:
	;
	v711 = v706 + v698
	v713 = *(*int32)(unsafe.Add(mBase, _consts[516]))
	if v707 == v713 {
		goto L161
	} else {
		goto L162
	}
L158:
	;
	if v729 == int32(0) {
		v823 = v711
		v824 = v707
		goto L154
	} else {
		goto L177
	}
L159:
	;
	v782 = int32(0)
	goto L158
L160:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v718)+12)) = v715
	*(*int32)(unsafe.Add(mBase, uint32(v715)+8)) = v718
	v823 = v711
	v824 = v707
	goto L154
L161:
	;
	v763 = *(*int32)(unsafe.Add(mBase, uint32(v699)+4))
	v764 = int32(3)
	if v763&v764 != v764 {
		v823 = v711
		v824 = v707
		goto L154
	} else {
		goto L176
	}
L162:
	;
	v715 = *(*int32)(unsafe.Add(mBase, uint32(v707)+12))
	if base.Ui32(int32(255)) < base.Ui32(v706) {
		goto L163
	} else {
		goto L164
	}
L163:
	;
	v729 = *(*int32)(unsafe.Add(mBase, uint32(v707)+24))
	if v715 == v707 {
		goto L166
	} else {
		goto L167
	}
L164:
	;
	v718 = *(*int32)(unsafe.Add(mBase, uint32(v707)+8))
	if v715 != v718 {
		goto L160
	} else {
		goto L165
	}
L165:
	;
	v720 = int32(0)
	v722 = *(*int32)(unsafe.Add(mBase, _consts[517]))
	*(*int32)(unsafe.Add(mBase, _consts[517])) = v722 & base.I32_rotl(int32(-2), int32(base.Ui32(v706)>>(uint(int32(3))%32)))
	v823 = v711
	v824 = v707
	goto L154
L166:
	;
	v734 = *(*int32)(unsafe.Add(mBase, uint32(v707)+20))
	if v734 == int32(0) {
		goto L169
	} else {
		goto L170
	}
L167:
	;
	v731 = *(*int32)(unsafe.Add(mBase, uint32(v707)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v731)+12)) = v715
	*(*int32)(unsafe.Add(mBase, uint32(v715)+8)) = v731
	v782 = v715
	goto L158
L168:
	;
	__phi750 = v744
	__phi751 = v745
	v750 = __phi750
	v751 = __phi751
	goto L172
L169:
	;
	v739 = *(*int32)(unsafe.Add(mBase, uint32(v707)+16))
	if v739 == int32(0) {
		goto L159
	} else {
		goto L171
	}
L170:
	;
	v744 = v734
	v745 = v707 + int32(20)
	goto L168
L171:
	;
	v744 = v739
	v745 = v707 + int32(16)
	goto L168
L172:
	;
	v757 = *(*int32)(unsafe.Add(mBase, uint32(v750)+20))
	if v757 != 0 {
		__phi750 = v757
		__phi751 = v750 + int32(20)
		v750 = __phi750
		v751 = __phi751
		goto L172
	} else {
		goto L174
	}
L173:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v751))) = int32(0)
	v782 = v750
	goto L158
L174:
	;
	v760 = *(*int32)(unsafe.Add(mBase, uint32(v750)+16))
	if v760 != 0 {
		__phi750 = v760
		__phi751 = v750 + int32(16)
		v750 = __phi750
		v751 = __phi751
		goto L172
	} else {
		goto L175
	}
L175:
	;
	goto L173
L176:
	;
	*(*int32)(unsafe.Add(mBase, _consts[518])) = v711
	*(*int32)(unsafe.Add(mBase, uint32(v699)+4)) = v763 & int32(-2)
	*(*int32)(unsafe.Add(mBase, uint32(v707)+4)) = v711 | int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v699))) = v711
	goto L151
L177:
	;
	v791 = *(*int32)(unsafe.Add(mBase, uint32(v707)+28))
	v793 = v791 << (uint(int32(2)) % 32)
	v796 = *(*int32)(unsafe.Add(mBase, uint32(v793)+uint32(_consts[519])))
	if v707 != v796 {
		goto L179
	} else {
		goto L180
	}
L178:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v782)+24)) = v729
	v813 = *(*int32)(unsafe.Add(mBase, uint32(v707)+16))
	if v813 == int32(0) {
		goto L186
	} else {
		goto L187
	}
L179:
	;
	v806 = *(*int32)(unsafe.Add(mBase, uint32(v729)+16))
	if v806 != v707 {
		goto L183
	} else {
		goto L184
	}
L180:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v793)+uint32(_consts[519]))) = v782
	if v782 != 0 {
		goto L178
	} else {
		goto L181
	}
L181:
	;
	v799 = int32(0)
	v801 = *(*int32)(unsafe.Add(mBase, _consts[520]))
	*(*int32)(unsafe.Add(mBase, _consts[520])) = v801 & base.I32_rotl(int32(-2), v791)
	v823 = v711
	v824 = v707
	goto L154
L182:
	;
	if v782 == int32(0) {
		v823 = v711
		v824 = v707
		goto L154
	} else {
		goto L185
	}
L183:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v729)+20)) = v782
	goto L182
L184:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v729)+16)) = v782
	goto L182
L185:
	;
	goto L178
L186:
	;
	v818 = *(*int32)(unsafe.Add(mBase, uint32(v707)+20))
	if v818 == int32(0) {
		v823 = v711
		v824 = v707
		goto L154
	} else {
		goto L188
	}
L187:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v782)+16)) = v813
	*(*int32)(unsafe.Add(mBase, uint32(v813)+24)) = v782
	goto L186
L188:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v782)+20)) = v818
	*(*int32)(unsafe.Add(mBase, uint32(v818)+24)) = v782
	v823 = v711
	v824 = v707
	goto L154
L189:
	;
	v833 = *(*int32)(unsafe.Add(mBase, uint32(v699)+4))
	if v833&int32(1) == int32(0) {
		goto L152
	} else {
		goto L190
	}
L190:
	;
	if v833&int32(2) != 0 {
		goto L195
	} else {
		goto L196
	}
L191:
	;
	if base.Ui32(int32(255)) < base.Ui32(v999) {
		goto L229
	} else {
		goto L230
	}
L192:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v824)+4)) = v879 | int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v824+v879))) = v879
	if v824 != v863 {
		v999 = v879
		goto L191
	} else {
		goto L228
	}
L193:
	;
	if v896 == int32(0) {
		goto L192
	} else {
		goto L216
	}
L194:
	;
	v941 = int32(0)
	goto L193
L195:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v699)+4)) = v833 & int32(-2)
	*(*int32)(unsafe.Add(mBase, uint32(v824)+4)) = v823 | int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v824+v823))) = v823
	v999 = v823
	goto L191
L196:
	;
	v841 = *(*int32)(unsafe.Add(mBase, _consts[521]))
	if v699 != v841 {
		goto L197
	} else {
		goto L198
	}
L197:
	;
	v863 = *(*int32)(unsafe.Add(mBase, _consts[516]))
	if v699 != v863 {
		goto L200
	} else {
		goto L201
	}
L198:
	;
	v843 = int32(0)
	*(*int32)(unsafe.Add(mBase, _consts[521])) = v824
	v847 = *(*int32)(unsafe.Add(mBase, _consts[522]))
	v848 = v847 + v823
	*(*int32)(unsafe.Add(mBase, _consts[522])) = v848
	*(*int32)(unsafe.Add(mBase, uint32(v824)+4)) = v848 | int32(1)
	v854 = *(*int32)(unsafe.Add(mBase, _consts[516]))
	if v824 != v854 {
		goto L152
	} else {
		goto L199
	}
L199:
	;
	v856 = int32(0)
	*(*int32)(unsafe.Add(mBase, _consts[518])) = v856
	*(*int32)(unsafe.Add(mBase, _consts[516])) = v856
	goto L151
L200:
	;
	v879 = v833&int32(-8) + v823
	v880 = *(*int32)(unsafe.Add(mBase, uint32(v699)+12))
	if base.Ui32(int32(255)) < base.Ui32(v833) {
		goto L202
	} else {
		goto L203
	}
L201:
	;
	v865 = int32(0)
	*(*int32)(unsafe.Add(mBase, _consts[516])) = v824
	v869 = *(*int32)(unsafe.Add(mBase, _consts[518]))
	v870 = v869 + v823
	*(*int32)(unsafe.Add(mBase, _consts[518])) = v870
	*(*int32)(unsafe.Add(mBase, uint32(v824)+4)) = v870 | int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v824+v870))) = v870
	goto L151
L202:
	;
	v896 = *(*int32)(unsafe.Add(mBase, uint32(v699)+24))
	if v880 == v699 {
		goto L206
	} else {
		goto L207
	}
L203:
	;
	v883 = *(*int32)(unsafe.Add(mBase, uint32(v699)+8))
	if v880 != v883 {
		goto L204
	} else {
		goto L205
	}
L204:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v883)+12)) = v880
	*(*int32)(unsafe.Add(mBase, uint32(v880)+8)) = v883
	goto L192
L205:
	;
	v885 = int32(0)
	v887 = *(*int32)(unsafe.Add(mBase, _consts[517]))
	*(*int32)(unsafe.Add(mBase, _consts[517])) = v887 & base.I32_rotl(int32(-2), int32(base.Ui32(v833)>>(uint(int32(3))%32)))
	goto L192
L206:
	;
	v901 = *(*int32)(unsafe.Add(mBase, uint32(v699)+20))
	if v901 == int32(0) {
		goto L209
	} else {
		goto L210
	}
L207:
	;
	v898 = *(*int32)(unsafe.Add(mBase, uint32(v699)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v898)+12)) = v880
	*(*int32)(unsafe.Add(mBase, uint32(v880)+8)) = v898
	v941 = v880
	goto L193
L208:
	;
	__phi917 = v911
	__phi918 = v912
	v917 = __phi917
	v918 = __phi918
	goto L212
L209:
	;
	v906 = *(*int32)(unsafe.Add(mBase, uint32(v699)+16))
	if v906 == int32(0) {
		goto L194
	} else {
		goto L211
	}
L210:
	;
	v911 = v901
	v912 = v699 + int32(20)
	goto L208
L211:
	;
	v911 = v906
	v912 = v699 + int32(16)
	goto L208
L212:
	;
	v924 = *(*int32)(unsafe.Add(mBase, uint32(v917)+20))
	if v924 != 0 {
		__phi917 = v924
		__phi918 = v917 + int32(20)
		v917 = __phi917
		v918 = __phi918
		goto L212
	} else {
		goto L214
	}
L213:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v918))) = int32(0)
	v941 = v917
	goto L193
L214:
	;
	v927 = *(*int32)(unsafe.Add(mBase, uint32(v917)+16))
	if v927 != 0 {
		__phi917 = v927
		__phi918 = v917 + int32(16)
		v917 = __phi917
		v918 = __phi918
		goto L212
	} else {
		goto L215
	}
L215:
	;
	goto L213
L216:
	;
	v950 = *(*int32)(unsafe.Add(mBase, uint32(v699)+28))
	v952 = v950 << (uint(int32(2)) % 32)
	v955 = *(*int32)(unsafe.Add(mBase, uint32(v952)+uint32(_consts[519])))
	if v699 != v955 {
		goto L218
	} else {
		goto L219
	}
L217:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v941)+24)) = v896
	v972 = *(*int32)(unsafe.Add(mBase, uint32(v699)+16))
	if v972 == int32(0) {
		goto L225
	} else {
		goto L226
	}
L218:
	;
	v965 = *(*int32)(unsafe.Add(mBase, uint32(v896)+16))
	if v965 != v699 {
		goto L222
	} else {
		goto L223
	}
L219:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v952)+uint32(_consts[519]))) = v941
	if v941 != 0 {
		goto L217
	} else {
		goto L220
	}
L220:
	;
	v958 = int32(0)
	v960 = *(*int32)(unsafe.Add(mBase, _consts[520]))
	*(*int32)(unsafe.Add(mBase, _consts[520])) = v960 & base.I32_rotl(int32(-2), v950)
	goto L192
L221:
	;
	if v941 == int32(0) {
		goto L192
	} else {
		goto L224
	}
L222:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v896)+20)) = v941
	goto L221
L223:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v896)+16)) = v941
	goto L221
L224:
	;
	goto L217
L225:
	;
	v977 = *(*int32)(unsafe.Add(mBase, uint32(v699)+20))
	if v977 == int32(0) {
		goto L192
	} else {
		goto L227
	}
L226:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v941)+16)) = v972
	*(*int32)(unsafe.Add(mBase, uint32(v972)+24)) = v941
	goto L225
L227:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v941)+20)) = v977
	*(*int32)(unsafe.Add(mBase, uint32(v977)+24)) = v941
	goto L192
L228:
	;
	*(*int32)(unsafe.Add(mBase, _consts[518])) = v879
	goto L151
L229:
	;
	if base.Ui32(int32(16777215)) < base.Ui32(v999) {
		v1046 = int32(31)
		goto L234
	} else {
		goto L235
	}
L230:
	;
	v1011 = v999 & int32(-8)
	v1013 = v1011 + int32(9128464)
	v1015 = *(*int32)(unsafe.Add(mBase, _consts[517]))
	v1019 = int32(1) << (uint(int32(base.Ui32(v999)>>(uint(int32(3))%32))) % 32)
	if v1015&v1019 != 0 {
		goto L232
	} else {
		goto L233
	}
L231:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1011)+uint32(_consts[523]))) = v824
	*(*int32)(unsafe.Add(mBase, uint32(v1025)+12)) = v824
	*(*int32)(unsafe.Add(mBase, uint32(v824)+12)) = v1013
	*(*int32)(unsafe.Add(mBase, uint32(v824)+8)) = v1025
	goto L151
L232:
	;
	v1024 = *(*int32)(unsafe.Add(mBase, uint32(v1011)+uint32(_consts[523])))
	v1025 = v1024
	goto L231
L233:
	;
	*(*int32)(unsafe.Add(mBase, _consts[517])) = v1015 | v1019
	v1025 = v1013
	goto L231
L234:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v824)+28)) = v1046
	*(*int64)(unsafe.Add(mBase, uint32(v824)+16)) = int64(0)
	v1051 = v1046 << (uint(int32(2)) % 32)
	v1055 = *(*int32)(unsafe.Add(mBase, _consts[520]))
	v1057 = int32(1) << (uint(v1046) % 32)
	if v1055&v1057 != 0 {
		goto L239
	} else {
		goto L240
	}
L235:
	;
	v1036 = base.I32_clz(int32(base.Ui32(v999) >> (uint(int32(8)) % 32)))
	v1039 = int32(1)
	v1046 = int32(base.Ui32(v999)>>(uint(int32(38)-v1036)%32))&v1039 - v1036<<(uint(v1039)%32) + int32(62)
	goto L234
L236:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v824+v1118))) = v1121
	*(*int32)(unsafe.Add(mBase, uint32(v824)+12)) = v1120
	*(*int32)(unsafe.Add(mBase, uint32(v824+v1116))) = v1119
	v1130 = int32(0)
	v1132 = *(*int32)(unsafe.Add(mBase, _consts[524]))
	v1133 = int32(-1)
	v1134 = v1132 + v1133
	if v1134 != 0 {
		goto L248
	} else {
		goto L249
	}
L237:
	;
	v1110 = *(*int32)(unsafe.Add(mBase, uint32(v1080)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v1110)+12)) = v824
	*(*int32)(unsafe.Add(mBase, uint32(v1080)+8)) = v824
	v1116 = int32(24)
	v1118 = int32(8)
	v1119 = int32(0)
	v1120 = v1080
	v1121 = v1110
	goto L236
L238:
	;
	v1116 = v1101
	v1118 = v1103
	v1119 = v824
	v1120 = v824
	v1121 = v1106
	goto L236
L239:
	;
	if v1046 == int32(31) {
		goto L241
	} else {
		goto L242
	}
L240:
	;
	*(*int32)(unsafe.Add(mBase, _consts[520])) = v1055 | v1057
	*(*int32)(unsafe.Add(mBase, uint32(v1051)+uint32(_consts[519]))) = v824
	v1101 = int32(8)
	v1103 = int32(24)
	v1106 = v1051 + int32(9128728)
	goto L238
L241:
	;
	v1072 = int32(0)
	goto L243
L242:
	;
	v1072 = int32(25) - int32(base.Ui32(v1046)>>(uint(int32(1))%32))
	goto L243
L243:
	;
	v1074 = *(*int32)(unsafe.Add(mBase, uint32(v1051)+uint32(_consts[519])))
	v1077 = v999 << (uint(v1072) % 32)
	v1080 = v1074
	goto L244
L244:
	;
	v1084 = *(*int32)(unsafe.Add(mBase, uint32(v1080)+4))
	if v1084&int32(-8) == v999 {
		goto L237
	} else {
		goto L246
	}
L245:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1094+int32(16)))) = v824
	v1101 = int32(8)
	v1103 = int32(24)
	v1106 = v1080
	goto L238
L246:
	;
	v1094 = v1080 + int32(base.Ui32(v1077)>>(uint(int32(29))%32))&int32(4)
	v1095 = *(*int32)(unsafe.Add(mBase, uint32(v1094)+16))
	if v1095 != 0 {
		v1077 = v1077 << (uint(int32(1)) % 32)
		v1080 = v1095
		goto L244
	} else {
		goto L247
	}
L247:
	;
	goto L245
L248:
	;
	v1136 = v1134
	goto L250
L249:
	;
	v1136 = v1133
	goto L250
L250:
	;
	*(*int32)(unsafe.Add(mBase, _consts[524])) = v1136
	goto L152
L251:
	;
	return int32(0)
L252:
	;
	if v1203 != 0 {
		goto L4
	} else {
		goto L270
	}
L253:
	;
	if base.Ui32(v1159) < base.Ui32(int32(-64)) {
		goto L255
	} else {
		goto L256
	}
L254:
	;
	v1162 = F_emscripten_builtin_malloc(m, v1159)
	mBase = m.M
	v1203 = v1162
	goto L252
L255:
	;
	if base.Ui32(v1159) < base.Ui32(int32(11)) {
		goto L258
	} else {
		goto L259
	}
L256:
	;
	v1165 = F___errno_location(m)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v1165))) = int32(48)
	v1203 = int32(0)
	goto L252
L257:
	;
	v1184 = F_emscripten_builtin_malloc(m, v1159)
	mBase = m.M
	if v1184 != 0 {
		goto L262
	} else {
		goto L263
	}
L258:
	;
	v1178 = int32(16)
	goto L260
L259:
	;
	v1178 = (l1 + int32(19)) & int32(-8)
	goto L260
L260:
	;
	v1179 = F_try_realloc_chunk(m, l0+int32(-16), v1178)
	mBase = m.M
	if v1179 == int32(0) {
		goto L257
	} else {
		goto L261
	}
L261:
	;
	v1203 = v1179 + int32(8)
	goto L252
L262:
	;
	v1190 = *(*int32)(unsafe.Add(mBase, uint32(l0+int32(-12))))
	if v1190&int32(3) != 0 {
		goto L264
	} else {
		goto L265
	}
L263:
	;
	v1203 = int32(0)
	goto L252
L264:
	;
	v1193 = int32(-4)
	goto L266
L265:
	;
	v1193 = int32(-8)
	goto L266
L266:
	;
	v1196 = v1193 + v1190&int32(-8)
	if base.Ui32(v1196) < base.Ui32(v1159) {
		goto L267
	} else {
		goto L268
	}
L267:
	;
	v1198 = v1196
	goto L269
L268:
	;
	v1198 = v1159
	goto L269
L269:
	;
	v1199 = F___memcpy(m, v1184, v608, v1198)
	mBase = m.M
	F_emscripten_builtin_free(m, v608)
	mBase = m.M
	v1203 = v1184
	goto L252
L270:
	;
	if l2 != 0 {
		goto L5
	} else {
		goto L271
	}
L271:
	;
	return int32(0)
L272:
	;
	if v1227 < int32(260) {
		goto L277
	} else {
		goto L278
	}
L273:
	;
	v1219 = int32(0)
	v1221 = *(*int32)(unsafe.Add(mBase, _consts[281]))
	*(*int32)(unsafe.Add(mBase, _consts[411])) = v1221
	*(*int32)(unsafe.Add(mBase, _consts[281])) = v1221 + int32(1)
	v1227 = v1221
	goto L272
L274:
	;
	if l2 == int32(0) {
		goto L282
	} else {
		goto L283
	}
L275:
	;
	v1265 = v1254 << (uint(int32(2)) % 32)
	v1268 = *(*int32)(unsafe.Add(mBase, uint32(v1265)+uint32(_consts[285])))
	*(*int32)(unsafe.Add(mBase, uint32(v1265)+uint32(_consts[285]))) = v1268 + l1
	goto L274
L276:
	;
	v1259 = int32(0)
	v1261 = *(*int32)(unsafe.Add(mBase, _consts[286]))
	*(*int32)(unsafe.Add(mBase, _consts[286])) = v1261 + l1
	goto L274
L277:
	;
	v1236 = v1227 << (uint(int32(2)) % 32)
	v1239 = *(*int32)(unsafe.Add(mBase, uint32(v1236)+uint32(_consts[285])))
	*(*int32)(unsafe.Add(mBase, uint32(v1236)+uint32(_consts[285]))) = v1239 - v609
	v1243 = *(*int32)(unsafe.Add(mBase, _consts[411]))
	if v1243 != int32(-1) {
		v1254 = v1243
		goto L279
	} else {
		goto L280
	}
L278:
	;
	v1230 = int32(0)
	v1232 = *(*int32)(unsafe.Add(mBase, _consts[286]))
	*(*int32)(unsafe.Add(mBase, _consts[286])) = v1232 - v609
	goto L276
L279:
	;
	if v1254 < int32(260) {
		goto L275
	} else {
		goto L281
	}
L280:
	;
	v1246 = int32(0)
	v1248 = *(*int32)(unsafe.Add(mBase, _consts[281]))
	*(*int32)(unsafe.Add(mBase, _consts[411])) = v1248
	*(*int32)(unsafe.Add(mBase, _consts[281])) = v1248 + int32(1)
	v1254 = v1248
	goto L279
L281:
	;
	goto L276
L282:
	;
	v1278 = v1203 + int32(8)
	goto L3
L283:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = l1
	goto L282
L284:
	;
	return int32(0)
L285:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L286:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_zuiFind(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
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
	var v24 int32
	_ = v24
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int64
	_ = v47
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int64
	_ = v66
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v76 int32
	_ = v76
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v91 float64
	_ = v91
	var v99 int32
	_ = v99
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
	var v112 int32
	_ = v112
	var v116 int32
	_ = v116
	v9 = m.G0
	v11 = v9 - int32(16)
	m.G0 = v11
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v13 != 0 {
		v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		switch v15 + int32(-2) {
		case 0:
			v18 = *(*int32)(unsafe.Add(mBase, uint32(l1)+36))
			if v18 == int32(0) {
				v41 = *(*int32)(unsafe.Add(mBase, uint32(l1)+44))
				v42 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
				v44 = v41
				v45 = v42
			} else {
				v24 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18+int32(-1)))))
				switch v24 & int32(7) {
				case 0:
					v44 = int32(base.Ui32(v24) >> (uint(int32(3)) % 32))
					v45 = v18
				case 1:
					v31 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18+int32(-3)))))
					v44 = v31
					v45 = v18
				case 2:
					v34 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v18+int32(-5)))))
					v44 = v34
					v45 = v18
				case 3:
					v37 = *(*int32)(unsafe.Add(mBase, uint32(v18+int32(-9))))
					v44 = v37
					v45 = v18
				case 4:
					v40 = *(*int32)(unsafe.Add(mBase, uint32(v18+int32(-17))))
					v44 = v40
					v45 = v18
				default:
					v44 = int32(0)
					v45 = v18
				}
			}
			v46 = int32(0)
			v47 = *(*int64)(unsafe.Add(mBase, uint32(l1)+48))
			v50 = F_setTypeIsMemberAux(m, v13, v45, v44, v47, base.B2i32(v18 != v46))
			mBase = m.M
			v53 = m.ExcPending
			if v53 != 0 {
				return int32(0)
			} else {
				if v50 == int32(0) {
					v116 = v46
				} else {
					*(*int64)(unsafe.Add(mBase, uint32(l2))) = int64(4607182418800017408)
					v116 = int32(1)
				}
				m.G0 = v11 + int32(16)
				return v116
			}
		case 1:
			v59 = *(*int32)(unsafe.Add(mBase, uint32(l1)+36))
			if v59 != 0 {
				v76 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				switch v76 + int32(-7) {
				case 0:
					v80 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
					v81 = F_objectGetVal(m, v80)
					mBase = m.M
					v82 = *(*int32)(unsafe.Add(mBase, uint32(v81)))
					v83 = *(*int32)(unsafe.Add(mBase, uint32(l1)+36))
					v86 = F_hashtableFind(m, v82, v83, v11+int32(12))
					mBase = m.M
					v87 = m.ExcPending
					if v87 != 0 {
						return int32(0)
					} else {
						if v86 == int32(0) {
							v116 = int32(0)
						} else {
							v90 = *(*int32)(unsafe.Add(mBase, uint32(v11)+12))
							v91 = *(*float64)(unsafe.Add(mBase, uint32(v90)))
							*(*float64)(unsafe.Add(mBase, uint32(l2))) = v91
							v116 = int32(1)
						}
						m.G0 = v11 + int32(16)
						return v116
					}
				default:
					F__serverPanic_1(m, int32(_a1723), int32(2371), int32(_a1054), int32(0))
					mBase = m.M
					v99 = m.ExcPending
					if v99 != 0 {
						return int32(0)
					} else {
						F_abort(m)
						mBase = m.M
						base.Wasm_trap_unreachable()
						for {
						}
					}
				case 4:
					v108 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
					v109 = F_objectGetVal(m, v108)
					mBase = m.M
					v110 = *(*int32)(unsafe.Add(mBase, uint32(l1)+36))
					v111 = F_zzlFind(m, v109, v110, l2)
					mBase = m.M
					v112 = m.ExcPending
					if v112 != 0 {
						return int32(0)
					} else {
						v116 = base.B2i32(v111 != int32(0))
						m.G0 = v11 + int32(16)
						return v116
					}
				}
			} else {
				v60 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
				if v60 == int32(0) {
					v66 = *(*int64)(unsafe.Add(mBase, uint32(l1)+48))
					v67 = F_sdsfromlonglong(m, v66)
					mBase = m.M
					v68 = m.ExcPending
					if v68 != 0 {
						return int32(0)
					} else {
						v69 = v67
						*(*int32)(unsafe.Add(mBase, uint32(l1)+36)) = v69
						v71 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
						*(*int32)(unsafe.Add(mBase, uint32(l1))) = v71 | int32(1)
						v76 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
						switch v76 + int32(-7) {
						case 0:
							v80 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
							v81 = F_objectGetVal(m, v80)
							mBase = m.M
							v82 = *(*int32)(unsafe.Add(mBase, uint32(v81)))
							v83 = *(*int32)(unsafe.Add(mBase, uint32(l1)+36))
							v86 = F_hashtableFind(m, v82, v83, v11+int32(12))
							mBase = m.M
							v87 = m.ExcPending
							if v87 != 0 {
								return int32(0)
							} else {
								if v86 == int32(0) {
									v116 = int32(0)
								} else {
									v90 = *(*int32)(unsafe.Add(mBase, uint32(v11)+12))
									v91 = *(*float64)(unsafe.Add(mBase, uint32(v90)))
									*(*float64)(unsafe.Add(mBase, uint32(l2))) = v91
									v116 = int32(1)
								}
								m.G0 = v11 + int32(16)
								return v116
							}
						default:
							F__serverPanic_1(m, int32(_a1723), int32(2371), int32(_a1054), int32(0))
							mBase = m.M
							v99 = m.ExcPending
							if v99 != 0 {
								return int32(0)
							} else {
								F_abort(m)
								mBase = m.M
								base.Wasm_trap_unreachable()
								for {
								}
							}
						case 4:
							v108 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
							v109 = F_objectGetVal(m, v108)
							mBase = m.M
							v110 = *(*int32)(unsafe.Add(mBase, uint32(l1)+36))
							v111 = F_zzlFind(m, v109, v110, l2)
							mBase = m.M
							v112 = m.ExcPending
							if v112 != 0 {
								return int32(0)
							} else {
								v116 = base.B2i32(v111 != int32(0))
								m.G0 = v11 + int32(16)
								return v116
							}
						}
					}
				} else {
					v63 = *(*int32)(unsafe.Add(mBase, uint32(l1)+44))
					v64 = F_sdsnewlen(m, v60, v63)
					mBase = m.M
					v65 = m.ExcPending
					if v65 != 0 {
						return int32(0)
					} else {
						v69 = v64
						*(*int32)(unsafe.Add(mBase, uint32(l1)+36)) = v69
						v71 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
						*(*int32)(unsafe.Add(mBase, uint32(l1))) = v71 | int32(1)
						v76 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
						switch v76 + int32(-7) {
						case 0:
							v80 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
							v81 = F_objectGetVal(m, v80)
							mBase = m.M
							v82 = *(*int32)(unsafe.Add(mBase, uint32(v81)))
							v83 = *(*int32)(unsafe.Add(mBase, uint32(l1)+36))
							v86 = F_hashtableFind(m, v82, v83, v11+int32(12))
							mBase = m.M
							v87 = m.ExcPending
							if v87 != 0 {
								return int32(0)
							} else {
								if v86 == int32(0) {
									v116 = int32(0)
								} else {
									v90 = *(*int32)(unsafe.Add(mBase, uint32(v11)+12))
									v91 = *(*float64)(unsafe.Add(mBase, uint32(v90)))
									*(*float64)(unsafe.Add(mBase, uint32(l2))) = v91
									v116 = int32(1)
								}
								m.G0 = v11 + int32(16)
								return v116
							}
						default:
							F__serverPanic_1(m, int32(_a1723), int32(2371), int32(_a1054), int32(0))
							mBase = m.M
							v99 = m.ExcPending
							if v99 != 0 {
								return int32(0)
							} else {
								F_abort(m)
								mBase = m.M
								base.Wasm_trap_unreachable()
								for {
								}
							}
						case 4:
							v108 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
							v109 = F_objectGetVal(m, v108)
							mBase = m.M
							v110 = *(*int32)(unsafe.Add(mBase, uint32(l1)+36))
							v111 = F_zzlFind(m, v109, v110, l2)
							mBase = m.M
							v112 = m.ExcPending
							if v112 != 0 {
								return int32(0)
							} else {
								v116 = base.B2i32(v111 != int32(0))
								m.G0 = v11 + int32(16)
								return v116
							}
						}
					}
				}
			}
		default:
			F__serverPanic_1(m, int32(_a1723), int32(2374), int32(_a1734), int32(0))
			mBase = m.M
			v106 = m.ExcPending
			if v106 != 0 {
				return int32(0)
			} else {
				F_abort(m)
				mBase = m.M
				base.Wasm_trap_unreachable()
				for {
				}
			}
		}
	} else {
		v116 = int32(0)
		m.G0 = v11 + int32(16)
		return v116
	}
}
func F_zunionInterDiffGetKeys(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	v6 = int32(1)
	v9 = F_genericGetKeys(m, int32(0), v6, int32(2), v6, l1, l2, l3)
	v12 = m.ExcPending
	if v12 != 0 {
		return int32(0)
	} else {
		return v9
	}
}
func F_zunionInterDiffStoreGetKeys(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	var v5 int32
	_ = v5
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	v5 = int32(1)
	v9 = F_genericGetKeys(m, v5, int32(2), int32(3), v5, l1, l2, l3)
	v12 = m.ExcPending
	if v12 != 0 {
		return int32(0)
	} else {
		return v9
	}
}
func F_zunionstoreCommand(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v3 = *(*int32)(unsafe.Add(mBase, uint32(v2)+4))
	v5 = int32(0)
	F_zunionInterDiffGenericCommand(m, l0, v3, int32(2), v5, v5)
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		return
	} else {
		return
	}
}
