package p2

import (
	base "github.com/shibukawa/vkmem/internal/aot/vkaot/base"
	"unsafe"
)

func F___bswap_32_1(m *base.Module, l0 int32) int32 {
	var v2 int32
	_ = v2
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	v2 = int32(24)
	v4 = int32(65280)
	v6 = int32(8)
	return l0<<(uint(v2)%32) | l0&v4<<(uint(v6)%32) | (int32(base.Ui32(l0)>>(uint(v6)%32))&v4 | int32(base.Ui32(l0)>>(uint(v2)%32)))
}
func F___builtin_ctz(m *base.Module, l0 int32) int32 {
	var v4 int32
	_ = v4
	if l0 != 0 {
		v4 = base.I32_ctz(l0)
	} else {
		v4 = int32(0)
	}
	return v4
}
func F_bigNumberCallback(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v8 < int32(0) {
		v32 = v7
	} else {
		v13 = l0 + v8<<(uint(int32(2))%32)
		v14 = *(*int32)(unsafe.Add(mBase, uint32(v13)+16))
		v15 = int32(1)
		v16 = v14 + v15
		*(*int32)(unsafe.Add(mBase, uint32(v13)+16)) = v16
		v18 = *(*int32)(unsafe.Add(mBase, uint32(v13)+1040))
		if v18 != v15 {
			v32 = v7
		} else {
			v23 = *(*int32)(unsafe.Add(mBase, uint32(v7)+8))
			*(*int32)(unsafe.Add(mBase, uint32(v23)+8)) = int32(3)
			*(*float64)(unsafe.Add(mBase, uint32(v23))) = base.F64_convert_i32_u(v16)
			v27 = *(*int32)(unsafe.Add(mBase, uint32(v7)+8))
			*(*int32)(unsafe.Add(mBase, uint32(v7)+8)) = v27 + int32(16)
			v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			v32 = v31
		}
	}
	v36 = F_lua_checkstack(m, v32, int32(3))
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		return
	} else {
		if v36 != 0 {
			v41 = int32(0)
			F_lua_createtable(m, v32, v41, v41)
			mBase = m.M
			v44 = m.ExcPending
			if v44 != 0 {
				return
			} else {
				v45 = m.G3
				F_lua_pushstring(m, v32, v45+int32(_a2017))
				mBase = m.M
				v49 = m.ExcPending
				if v49 != 0 {
					return
				} else {
					F_lua_pushlstring(m, v32, l1, l2)
					mBase = m.M
					v51 = m.ExcPending
					if v51 != 0 {
						return
					} else {
						F_lua_settable(m, v32, int32(-3))
						mBase = m.M
						v54 = m.ExcPending
						if v54 != 0 {
							return
						} else {
							F_processCollectionElementEnd(m, l0)
							mBase = m.M
							v56 = m.ExcPending
							if v56 != 0 {
								return
							} else {
								return
							}
						}
					}
				}
			}
		} else {
			F__serverPanic_2(m, int32(1067))
			mBase = m.M
			v40 = m.ExcPending
			if v40 != 0 {
				return
			} else {
				base.Wasm_trap_unreachable()
				for {
				}
			}
		}
	}
}
func F_bind(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v15 int32
	_ = v15
	v4 = int32(0)
	v7 = m.Env.X__syscall_bind(m, l0, l1, l2, v4, v4, v4)
	mBase = m.M
	if base.Ui32(v7) < base.Ui32(int32(-4095)) {
		v15 = v7
	} else {
		v10 = F___errno_location(m)
		mBase = m.M
		*(*int32)(unsafe.Add(mBase, uint32(v10))) = int32(0) - v7
		v15 = int32(-1)
	}
	return v15
}
func F_bitfieldCommand(m *base.Module, l0 int32) {
	var v4 int32
	_ = v4
	F_bitfieldGeneric(m, l0, int32(0))
	v4 = m.ExcPending
	if v4 != 0 {
		return
	} else {
		return
	}
}
func F_bitfieldGeneric(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v11 int64
	_ = v11
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v34 int32
	_ = v34
	var v38 int32
	_ = v38
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v53 int64
	_ = v53
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v130 int32
	_ = v130
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v136 int32
	_ = v136
	var v140 int32
	_ = v140
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v167 int32
	_ = v167
	var v169 int32
	_ = v169
	var v172 int32
	_ = v172
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v178 int32
	_ = v178
	var v182 int32
	_ = v182
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v209 int32
	_ = v209
	var v211 int32
	_ = v211
	var v215 int32
	_ = v215
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v226 int32
	_ = v226
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v236 int32
	_ = v236
	var v239 int32
	_ = v239
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v245 int32
	_ = v245
	var v249 int32
	_ = v249
	var v251 int32
	_ = v251
	var v252 int32
	_ = v252
	var v253 int32
	_ = v253
	var v255 int32
	_ = v255
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v258 int32
	_ = v258
	var v263 int32
	_ = v263
	var v264 int32
	_ = v264
	var v267 int32
	_ = v267
	var v268 int32
	_ = v268
	var v271 int32
	_ = v271
	var v272 int32
	_ = v272
	var v273 int32
	_ = v273
	var v277 int32
	_ = v277
	var v279 int32
	_ = v279
	var v283 int32
	_ = v283
	var v284 int32
	_ = v284
	var v285 int32
	_ = v285
	var v288 int32
	_ = v288
	var v290 int32
	_ = v290
	var v291 int32
	_ = v291
	var v292 int32
	_ = v292
	var v294 int32
	_ = v294
	var v298 int32
	_ = v298
	var v300 int32
	_ = v300
	var v301 int32
	_ = v301
	var v302 int32
	_ = v302
	var v304 int32
	_ = v304
	var v305 int32
	_ = v305
	var v306 int32
	_ = v306
	var v307 int32
	_ = v307
	var v312 int32
	_ = v312
	var v313 int32
	_ = v313
	var v316 int32
	_ = v316
	var v317 int32
	_ = v317
	var v320 int32
	_ = v320
	var v321 int32
	_ = v321
	var v322 int32
	_ = v322
	var v325 int32
	_ = v325
	var v328 int32
	_ = v328
	var v330 int32
	_ = v330
	var v331 int32
	_ = v331
	var v332 int32
	_ = v332
	var v334 int32
	_ = v334
	var v338 int32
	_ = v338
	var v340 int32
	_ = v340
	var v341 int32
	_ = v341
	var v342 int32
	_ = v342
	var v344 int32
	_ = v344
	var v345 int32
	_ = v345
	var v346 int32
	_ = v346
	var v347 int32
	_ = v347
	var v352 int32
	_ = v352
	var v353 int32
	_ = v353
	var v356 int32
	_ = v356
	var v357 int32
	_ = v357
	var v360 int32
	_ = v360
	var v361 int32
	_ = v361
	var v362 int32
	_ = v362
	var v365 int32
	_ = v365
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
	var v378 int32
	_ = v378
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
	var v386 int32
	_ = v386
	var v387 int32
	_ = v387
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
	var v402 int32
	_ = v402
	var v406 int32
	_ = v406
	var v408 int32
	_ = v408
	var v409 int64
	_ = v409
	var v410 int32
	_ = v410
	var v411 int32
	_ = v411
	var v415 int32
	_ = v415
	var v417 int32
	_ = v417
	var v418 int32
	_ = v418
	var v422 int64
	_ = v422
	var v424 int64
	_ = v424
	var v426 int32
	_ = v426
	var v427 int64
	_ = v427
	var v428 int32
	_ = v428
	var v430 int32
	_ = v430
	var v433 int32
	_ = v433
	var v436 int32
	_ = v436
	var v437 int32
	_ = v437
	var v438 int32
	_ = v438
	var v439 int64
	_ = v439
	var v441 int64
	_ = v441
	var v445 int32
	_ = v445
	var v447 int32
	_ = v447
	var v450 int32
	_ = v450
	var v451 int32
	_ = v451
	var v453 int32
	_ = v453
	var v454 int32
	_ = v454
	var v455 int64
	_ = v455
	var v456 int32
	_ = v456
	var v462 int32
	_ = v462
	var v463 int32
	_ = v463
	var v471 int32
	_ = v471
	var v472 int32
	_ = v472
	var v477 int64
	_ = v477
	var v491 int32
	_ = v491
	var v492 int32
	_ = v492
	var v493 int32
	_ = v493
	var v494 int32
	_ = v494
	var v495 int32
	_ = v495
	var v499 int32
	_ = v499
	var v500 int32
	_ = v500
	var v503 int32
	_ = v503
	var v509 int32
	_ = v509
	var v512 int32
	_ = v512
	var v515 int32
	_ = v515
	var v516 int32
	_ = v516
	var v522 int32
	_ = v522
	var v523 int32
	_ = v523
	var v524 int32
	_ = v524
	var v529 int64
	_ = v529
	var v544 int32
	_ = v544
	var v546 int32
	_ = v546
	var v552 int32
	_ = v552
	var v553 int32
	_ = v553
	var v565 int64
	_ = v565
	var v568 int32
	_ = v568
	var v569 int32
	_ = v569
	var v573 int64
	_ = v573
	var v578 int64
	_ = v578
	var v581 int32
	_ = v581
	var v582 int32
	_ = v582
	var v587 int32
	_ = v587
	var v588 int32
	_ = v588
	var v589 int32
	_ = v589
	var v590 int64
	_ = v590
	var v594 int64
	_ = v594
	var v600 int64
	_ = v600
	var v612 int64
	_ = v612
	var v618 int64
	_ = v618
	var v620 int64
	_ = v620
	var v626 int64
	_ = v626
	var v627 int64
	_ = v627
	var v628 int64
	_ = v628
	var v632 int64
	_ = v632
	var v633 int64
	_ = v633
	var v635 int64
	_ = v635
	var v640 int64
	_ = v640
	var v646 int64
	_ = v646
	var v657 int64
	_ = v657
	var v659 int64
	_ = v659
	var v661 int64
	_ = v661
	var v675 int64
	_ = v675
	var v681 int64
	_ = v681
	var v683 int64
	_ = v683
	var v697 int64
	_ = v697
	var v706 int64
	_ = v706
	var v727 int64
	_ = v727
	var v761 int64
	_ = v761
	var v768 int64
	_ = v768
	var v770 int64
	_ = v770
	var v775 int32
	_ = v775
	var v776 int64
	_ = v776
	var v778 int64
	_ = v778
	var v779 int32
	_ = v779
	var v780 int64
	_ = v780
	var v781 int32
	_ = v781
	var v785 int64
	_ = v785
	var v787 int32
	_ = v787
	var v797 int64
	_ = v797
	var v808 int64
	_ = v808
	var v811 int64
	_ = v811
	var v812 int64
	_ = v812
	var v820 int64
	_ = v820
	var v821 int64
	_ = v821
	var v837 int64
	_ = v837
	var v838 int64
	_ = v838
	var v846 int64
	_ = v846
	var v849 int64
	_ = v849
	var v855 int64
	_ = v855
	var v857 int64
	_ = v857
	var v860 int32
	_ = v860
	var v862 int64
	_ = v862
	var v868 int64
	_ = v868
	var v880 int64
	_ = v880
	var v886 int64
	_ = v886
	var v888 int64
	_ = v888
	var v894 int64
	_ = v894
	var v895 int64
	_ = v895
	var v896 int64
	_ = v896
	var v900 int64
	_ = v900
	var v901 int64
	_ = v901
	var v903 int64
	_ = v903
	var v908 int64
	_ = v908
	var v914 int64
	_ = v914
	var v925 int64
	_ = v925
	var v927 int64
	_ = v927
	var v929 int64
	_ = v929
	var v943 int64
	_ = v943
	var v949 int64
	_ = v949
	var v951 int64
	_ = v951
	var v965 int64
	_ = v965
	var v992 int64
	_ = v992
	var v998 int64
	_ = v998
	var v999 int32
	_ = v999
	var v1002 int64
	_ = v1002
	var v1003 int32
	_ = v1003
	var v1004 int64
	_ = v1004
	var v1008 int64
	_ = v1008
	var v1011 int64
	_ = v1011
	var v1023 int64
	_ = v1023
	var v1029 int64
	_ = v1029
	var v1032 int64
	_ = v1032
	var v1035 int64
	_ = v1035
	var v1037 int32
	_ = v1037
	var v1039 int32
	_ = v1039
	var v1043 int64
	_ = v1043
	var v1049 int32
	_ = v1049
	var v1052 int32
	_ = v1052
	var v1056 int64
	_ = v1056
	var v1059 int32
	_ = v1059
	var v1060 int32
	_ = v1060
	var v1061 int32
	_ = v1061
	var v1065 int64
	_ = v1065
	var v1068 int64
	_ = v1068
	var v1070 int64
	_ = v1070
	var v1073 int32
	_ = v1073
	var v1074 int32
	_ = v1074
	var v1077 int64
	_ = v1077
	var v1079 int64
	_ = v1079
	var v1084 int32
	_ = v1084
	var v1090 int32
	_ = v1090
	var v1096 int32
	_ = v1096
	var v1102 int32
	_ = v1102
	var v1108 int32
	_ = v1108
	var v1114 int32
	_ = v1114
	var v1120 int32
	_ = v1120
	var v1126 int32
	_ = v1126
	var v1128 int64
	_ = v1128
	var v1130 int64
	_ = v1130
	var v1132 int32
	_ = v1132
	var v1133 int64
	_ = v1133
	var v1134 int32
	_ = v1134
	var v1138 int64
	_ = v1138
	var v1141 int64
	_ = v1141
	var v1146 int64
	_ = v1146
	var v1147 int64
	_ = v1147
	var v1159 int64
	_ = v1159
	var v1165 int64
	_ = v1165
	var v1169 int64
	_ = v1169
	var v1174 int32
	_ = v1174
	var v1175 int64
	_ = v1175
	var v1176 int64
	_ = v1176
	var v1177 int64
	_ = v1177
	var v1181 int64
	_ = v1181
	var v1182 int64
	_ = v1182
	var v1184 int64
	_ = v1184
	var v1189 int64
	_ = v1189
	var v1197 int64
	_ = v1197
	var v1208 int64
	_ = v1208
	var v1210 int64
	_ = v1210
	var v1212 int64
	_ = v1212
	var v1226 int64
	_ = v1226
	var v1232 int64
	_ = v1232
	var v1234 int64
	_ = v1234
	var v1236 int64
	_ = v1236
	var v1250 int64
	_ = v1250
	var v1259 int64
	_ = v1259
	var v1272 int64
	_ = v1272
	var v1280 int64
	_ = v1280
	var v1297 int64
	_ = v1297
	var v1300 int64
	_ = v1300
	var v1305 int64
	_ = v1305
	var v1306 int64
	_ = v1306
	var v1318 int64
	_ = v1318
	var v1324 int64
	_ = v1324
	var v1328 int64
	_ = v1328
	var v1333 int32
	_ = v1333
	var v1334 int64
	_ = v1334
	var v1335 int64
	_ = v1335
	var v1336 int64
	_ = v1336
	var v1340 int64
	_ = v1340
	var v1341 int64
	_ = v1341
	var v1343 int64
	_ = v1343
	var v1348 int64
	_ = v1348
	var v1356 int64
	_ = v1356
	var v1367 int64
	_ = v1367
	var v1369 int64
	_ = v1369
	var v1371 int64
	_ = v1371
	var v1385 int64
	_ = v1385
	var v1391 int64
	_ = v1391
	var v1393 int64
	_ = v1393
	var v1395 int64
	_ = v1395
	var v1409 int64
	_ = v1409
	var v1428 int64
	_ = v1428
	var v1436 int64
	_ = v1436
	var v1443 int32
	_ = v1443
	var v1446 int64
	_ = v1446
	var v1450 int64
	_ = v1450
	var v1452 int32
	_ = v1452
	var v1453 int32
	_ = v1453
	var v1454 int64
	_ = v1454
	var v1457 int64
	_ = v1457
	var v1469 int64
	_ = v1469
	var v1475 int64
	_ = v1475
	var v1486 int32
	_ = v1486
	var v1487 int64
	_ = v1487
	var v1491 int64
	_ = v1491
	var v1496 int64
	_ = v1496
	var v1501 int32
	_ = v1501
	var v1504 int64
	_ = v1504
	var v1509 int64
	_ = v1509
	var v1521 int64
	_ = v1521
	var v1535 int32
	_ = v1535
	var v1541 int64
	_ = v1541
	var v1542 int64
	_ = v1542
	var v1544 int64
	_ = v1544
	var v1545 int64
	_ = v1545
	var v1547 int64
	_ = v1547
	var v1548 int64
	_ = v1548
	var v1550 int32
	_ = v1550
	var v1551 int32
	_ = v1551
	var v1552 int64
	_ = v1552
	var v1555 int64
	_ = v1555
	var v1567 int64
	_ = v1567
	var v1573 int64
	_ = v1573
	var v1584 int32
	_ = v1584
	var v1585 int64
	_ = v1585
	var v1589 int64
	_ = v1589
	var v1594 int64
	_ = v1594
	var v1599 int32
	_ = v1599
	var v1602 int64
	_ = v1602
	var v1607 int64
	_ = v1607
	var v1619 int64
	_ = v1619
	var v1633 int32
	_ = v1633
	var v1649 int64
	_ = v1649
	var v1653 int32
	_ = v1653
	var v1657 int64
	_ = v1657
	var v1662 int64
	_ = v1662
	var v1664 int32
	_ = v1664
	var v1668 int32
	_ = v1668
	var v1669 int32
	_ = v1669
	var v1670 int32
	_ = v1670
	var v1672 int32
	_ = v1672
	var v1675 int32
	_ = v1675
	var v1676 int32
	_ = v1676
	var v1677 int32
	_ = v1677
	var v1678 int32
	_ = v1678
	var v1680 int32
	_ = v1680
	var v1681 int32
	_ = v1681
	var v1683 int64
	_ = v1683
	var v1713 int32
	_ = v1713
	var v1719 int32
	_ = v1719
	var v1739 int32
	_ = v1739
	v3 = int32(0)
	v11 = int64(0)
	v25 = m.G0
	v27 = v25 - int32(64)
	m.G0 = v27
	*(*int32)(unsafe.Add(mBase, uint32(v27)+52)) = v3
	v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v34 < int32(3) {
		v471 = v3
		v472 = v3
		v477 = v11
		goto L6
	} else {
		goto L7
	}
L1:
	;
	m.G0 = v27 + int32(64)
	return
L2:
	;
	F_valkey_free(m, v1719)
	mBase = m.M
	v1739 = m.ExcPending
	if v1739 != 0 {
		goto L57
	} else {
		goto L304
	}
L3:
	;
	F_initDeferredReplyBuffer(m, l0)
	mBase = m.M
	v544 = m.ExcPending
	if v544 != 0 {
		goto L57
	} else {
		goto L139
	}
L4:
	;
	if l1&int32(1) == int32(0) {
		goto L133
	} else {
		goto L134
	}
L5:
	;
	F_addReplyError(m, l0, int32(_a185))
	mBase = m.M
	v503 = m.ExcPending
	if v503 != 0 {
		goto L57
	} else {
		goto L132
	}
L6:
	;
	v491 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v492 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v493 = *(*int32)(unsafe.Add(mBase, uint32(v492)+4))
	v494 = F_lookupKeyRead(m, v491, v493)
	mBase = m.M
	v495 = m.ExcPending
	if v495 != 0 {
		goto L57
	} else {
		goto L128
	}
L7:
	;
	v38 = int32(0)
	v47 = v38
	v48 = v38
	v49 = v34
	v50 = int32(2)
	v51 = v38
	v52 = int32(1)
	v53 = int64(0)
	goto L8
L8:
	;
	v67 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v69 = v50 << (uint(int32(2)) % 32)
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v67+v69)))
	v72 = F_objectGetVal(m, v71)
	mBase = m.M
	*(*int64)(unsafe.Add(mBase, uint32(v27))) = int64(0)
	v75 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v27)+32)) = v75
	*(*int32)(unsafe.Add(mBase, uint32(v27)+48)) = v75
	v82 = v49 + (v50 ^ int32(-1))
	v85 = int32(_a186)
	v88 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v72))))
	if v88 != 0 {
		goto L17
	} else {
		goto L18
	}
L9:
	;
	if v454 == int32(0) {
		goto L4
	} else {
		goto L127
	}
L10:
	;
	v462 = v456 + int32(1)
	v463 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v462 < v463 {
		v47 = v450
		v48 = v451
		v49 = v463
		v50 = v462
		v51 = v453
		v52 = v454
		v53 = v455
		goto L8
	} else {
		goto L126
	}
L11:
	;
	v430 = int32(5)
	v433 = v47 + int32(1)
	v436 = F_valkey_realloc(m, v48, v433<<(uint(v430)%32))
	mBase = m.M
	v437 = m.ExcPending
	if v437 != 0 {
		goto L57
	} else {
		goto L125
	}
L12:
	;
	v409 = *(*int64)(unsafe.Add(mBase, uint32(v27)+56))
	v410 = int32(0)
	v411 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v415 = *(*int32)(unsafe.Add(mBase, uint32(v411+v69+int32(12))))
	v417 = F_getLongLongFromObjectOrReply(m, l0, v415, v27, v410)
	mBase = m.M
	v418 = m.ExcPending
	if v418 != 0 {
		goto L57
	} else {
		goto L120
	}
L13:
	;
	v236 = int32(_a187)
	v239 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v72))))
	if v239 != 0 {
		goto L65
	} else {
		goto L66
	}
L14:
	;
	v211 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v215 = *(*int32)(unsafe.Add(mBase, uint32(v211+v69+int32(4))))
	v220 = F_getBitfieldTypeFromArgument(m, l0, v215, v27+int32(32), v27+int32(48))
	mBase = m.M
	v221 = m.ExcPending
	if v221 != 0 {
		goto L57
	} else {
		goto L58
	}
L15:
	;
	v126 = base.B2i32(int32(1) < v82) & base.B2i32(v120-v122 == int32(0))
	if v126 != 0 {
		v209 = v75
		goto L14
	} else {
		goto L27
	}
L16:
	;
	v120 = F_tolower(m, v116)
	mBase = m.M
	v121 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v117))))
	v122 = F_tolower(m, v121)
	mBase = m.M
	goto L15
L17:
	;
	v90 = v72
	v91 = v85
	v92 = v88
	goto L20
L18:
	;
	v116 = int32(0)
	v117 = v85
	goto L16
L19:
	;
	v116 = v113 & int32(255)
	v117 = v112
	goto L16
L20:
	;
	v94 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v91))))
	if v94 == int32(0) {
		v112 = v91
		v113 = v92
		goto L19
	} else {
		goto L22
	}
L21:
	;
	v112 = v106
	v113 = int32(0)
	goto L19
L22:
	;
	v98 = v92 & int32(255)
	if v98 == v94 {
		goto L23
	} else {
		goto L24
	}
L23:
	;
	v105 = int32(1)
	v106 = v91 + v105
	v107 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v90)+1)))
	if v107 != 0 {
		v90 = v90 + v105
		v91 = v106
		v92 = v107
		goto L20
	} else {
		goto L26
	}
L24:
	;
	v100 = F_tolower(m, v98)
	mBase = m.M
	v101 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v91))))
	v102 = F_tolower(m, v101)
	mBase = m.M
	if v100 == v102 {
		goto L23
	} else {
		goto L25
	}
L25:
	;
	v104 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v90))))
	v112 = v91
	v113 = v104
	goto L19
L26:
	;
	goto L21
L27:
	;
	v127 = int32(_a188)
	v130 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v72))))
	if v130 != 0 {
		goto L30
	} else {
		goto L31
	}
L28:
	;
	v167 = base.B2i32(v82 < int32(3))
	if v82 < int32(3) {
		goto L40
	} else {
		goto L41
	}
L29:
	;
	v162 = F_tolower(m, v158)
	mBase = m.M
	v163 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v159))))
	v164 = F_tolower(m, v163)
	mBase = m.M
	goto L28
L30:
	;
	v132 = v72
	v133 = v127
	v134 = v130
	goto L33
L31:
	;
	v158 = int32(0)
	v159 = v127
	goto L29
L32:
	;
	v158 = v155 & int32(255)
	v159 = v154
	goto L29
L33:
	;
	v136 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v133))))
	if v136 == int32(0) {
		v154 = v133
		v155 = v134
		goto L32
	} else {
		goto L35
	}
L34:
	;
	v154 = v148
	v155 = int32(0)
	goto L32
L35:
	;
	v140 = v134 & int32(255)
	if v140 == v136 {
		goto L36
	} else {
		goto L37
	}
L36:
	;
	v147 = int32(1)
	v148 = v133 + v147
	v149 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v132)+1)))
	if v149 != 0 {
		v132 = v132 + v147
		v133 = v148
		v134 = v149
		goto L33
	} else {
		goto L39
	}
L37:
	;
	v142 = F_tolower(m, v140)
	mBase = m.M
	v143 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v133))))
	v144 = F_tolower(m, v143)
	mBase = m.M
	if v142 == v144 {
		goto L36
	} else {
		goto L38
	}
L38:
	;
	v146 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v132))))
	v154 = v133
	v155 = v146
	goto L32
L39:
	;
	goto L34
L40:
	;
	v169 = int32(_a189)
	v172 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v72))))
	if v172 != 0 {
		goto L45
	} else {
		goto L46
	}
L41:
	;
	if v162-v164 != 0 {
		goto L40
	} else {
		goto L42
	}
L42:
	;
	v209 = int32(1)
	goto L14
L43:
	;
	if v82 < int32(3) {
		goto L13
	} else {
		goto L55
	}
L44:
	;
	v204 = F_tolower(m, v200)
	mBase = m.M
	v205 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v201))))
	v206 = F_tolower(m, v205)
	mBase = m.M
	goto L43
L45:
	;
	v174 = v72
	v175 = v169
	v176 = v172
	goto L48
L46:
	;
	v200 = int32(0)
	v201 = v169
	goto L44
L47:
	;
	v200 = v197 & int32(255)
	v201 = v196
	goto L44
L48:
	;
	v178 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v175))))
	if v178 == int32(0) {
		v196 = v175
		v197 = v176
		goto L47
	} else {
		goto L50
	}
L49:
	;
	v196 = v190
	v197 = int32(0)
	goto L47
L50:
	;
	v182 = v176 & int32(255)
	if v182 == v178 {
		goto L51
	} else {
		goto L52
	}
L51:
	;
	v189 = int32(1)
	v190 = v175 + v189
	v191 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v174)+1)))
	if v191 != 0 {
		v174 = v174 + v189
		v175 = v190
		v176 = v191
		goto L48
	} else {
		goto L54
	}
L52:
	;
	v184 = F_tolower(m, v182)
	mBase = m.M
	v185 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v175))))
	v186 = F_tolower(m, v185)
	mBase = m.M
	if v184 == v186 {
		goto L51
	} else {
		goto L53
	}
L53:
	;
	v188 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v174))))
	v196 = v175
	v197 = v188
	goto L47
L54:
	;
	goto L49
L55:
	;
	if v204-v206 != 0 {
		goto L13
	} else {
		goto L56
	}
L56:
	;
	v209 = int32(2)
	goto L14
L57:
	;
	return
L58:
	;
	if v220 != 0 {
		v1719 = v48
		goto L2
	} else {
		goto L59
	}
L59:
	;
	v222 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v226 = *(*int32)(unsafe.Add(mBase, uint32(v222+v69+int32(8))))
	v230 = *(*int32)(unsafe.Add(mBase, uint32(v27)+48))
	v231 = F_getBitOffsetFromArgument(m, l0, v226, v27+int32(56), int32(1), v230)
	mBase = m.M
	v232 = m.ExcPending
	if v232 != 0 {
		goto L57
	} else {
		goto L60
	}
L60:
	;
	if v231 != 0 {
		v1719 = v48
		goto L2
	} else {
		goto L61
	}
L61:
	;
	if v126 == int32(0) {
		goto L12
	} else {
		goto L62
	}
L62:
	;
	v426 = v52
	v427 = v53
	v428 = int32(2)
	goto L11
L63:
	;
	if v82 < int32(1) {
		goto L75
	} else {
		goto L76
	}
L64:
	;
	v271 = F_tolower(m, v267)
	mBase = m.M
	v272 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v268))))
	v273 = F_tolower(m, v272)
	mBase = m.M
	goto L63
L65:
	;
	v241 = v72
	v242 = v236
	v243 = v239
	goto L68
L66:
	;
	v267 = int32(0)
	v268 = v236
	goto L64
L67:
	;
	v267 = v264 & int32(255)
	v268 = v263
	goto L64
L68:
	;
	v245 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v242))))
	if v245 == int32(0) {
		v263 = v242
		v264 = v243
		goto L67
	} else {
		goto L70
	}
L69:
	;
	v263 = v257
	v264 = int32(0)
	goto L67
L70:
	;
	v249 = v243 & int32(255)
	if v249 == v245 {
		goto L71
	} else {
		goto L72
	}
L71:
	;
	v256 = int32(1)
	v257 = v242 + v256
	v258 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v241)+1)))
	if v258 != 0 {
		v241 = v241 + v256
		v242 = v257
		v243 = v258
		goto L68
	} else {
		goto L74
	}
L72:
	;
	v251 = F_tolower(m, v249)
	mBase = m.M
	v252 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v242))))
	v253 = F_tolower(m, v252)
	mBase = m.M
	if v251 == v253 {
		goto L71
	} else {
		goto L73
	}
L73:
	;
	v255 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v241))))
	v263 = v242
	v264 = v255
	goto L67
L74:
	;
	goto L69
L75:
	;
	v406 = *(*int32)(unsafe.Add(mBase, _consts[86]))
	F_addReplyErrorObject(m, l0, v406)
	mBase = m.M
	v408 = m.ExcPending
	if v408 != 0 {
		goto L57
	} else {
		goto L119
	}
L76:
	;
	if v271-v273 != 0 {
		goto L75
	} else {
		goto L77
	}
L77:
	;
	v277 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v279 = v50 + int32(1)
	v283 = *(*int32)(unsafe.Add(mBase, uint32(v277+v279<<(uint(int32(2))%32))))
	v284 = F_objectGetVal(m, v283)
	mBase = m.M
	v285 = int32(_a190)
	v288 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v284))))
	if v288 != 0 {
		goto L81
	} else {
		goto L82
	}
L78:
	;
	v325 = int32(_a191)
	v328 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v284))))
	if v328 != 0 {
		goto L95
	} else {
		goto L96
	}
L79:
	;
	if v320-v322 != 0 {
		goto L78
	} else {
		goto L91
	}
L80:
	;
	v320 = F_tolower(m, v316)
	mBase = m.M
	v321 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v317))))
	v322 = F_tolower(m, v321)
	mBase = m.M
	goto L79
L81:
	;
	v290 = v284
	v291 = v285
	v292 = v288
	goto L84
L82:
	;
	v316 = int32(0)
	v317 = v285
	goto L80
L83:
	;
	v316 = v313 & int32(255)
	v317 = v312
	goto L80
L84:
	;
	v294 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v291))))
	if v294 == int32(0) {
		v312 = v291
		v313 = v292
		goto L83
	} else {
		goto L86
	}
L85:
	;
	v312 = v306
	v313 = int32(0)
	goto L83
L86:
	;
	v298 = v292 & int32(255)
	if v298 == v294 {
		goto L87
	} else {
		goto L88
	}
L87:
	;
	v305 = int32(1)
	v306 = v291 + v305
	v307 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v290)+1)))
	if v307 != 0 {
		v290 = v290 + v305
		v291 = v306
		v292 = v307
		goto L84
	} else {
		goto L90
	}
L88:
	;
	v300 = F_tolower(m, v298)
	mBase = m.M
	v301 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v291))))
	v302 = F_tolower(m, v301)
	mBase = m.M
	if v300 == v302 {
		goto L87
	} else {
		goto L89
	}
L89:
	;
	v304 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v290))))
	v312 = v291
	v313 = v304
	goto L83
L90:
	;
	goto L85
L91:
	;
	v450 = v47
	v451 = v48
	v453 = int32(0)
	v454 = v52
	v455 = v53
	v456 = v279
	goto L10
L92:
	;
	v365 = int32(_a192)
	v368 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v284))))
	if v368 != 0 {
		goto L108
	} else {
		goto L109
	}
L93:
	;
	if v360-v362 != 0 {
		goto L92
	} else {
		goto L105
	}
L94:
	;
	v360 = F_tolower(m, v356)
	mBase = m.M
	v361 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v357))))
	v362 = F_tolower(m, v361)
	mBase = m.M
	goto L93
L95:
	;
	v330 = v284
	v331 = v325
	v332 = v328
	goto L98
L96:
	;
	v356 = int32(0)
	v357 = v325
	goto L94
L97:
	;
	v356 = v353 & int32(255)
	v357 = v352
	goto L94
L98:
	;
	v334 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v331))))
	if v334 == int32(0) {
		v352 = v331
		v353 = v332
		goto L97
	} else {
		goto L100
	}
L99:
	;
	v352 = v346
	v353 = int32(0)
	goto L97
L100:
	;
	v338 = v332 & int32(255)
	if v338 == v334 {
		goto L101
	} else {
		goto L102
	}
L101:
	;
	v345 = int32(1)
	v346 = v331 + v345
	v347 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v330)+1)))
	if v347 != 0 {
		v330 = v330 + v345
		v331 = v346
		v332 = v347
		goto L98
	} else {
		goto L104
	}
L102:
	;
	v340 = F_tolower(m, v338)
	mBase = m.M
	v341 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v331))))
	v342 = F_tolower(m, v341)
	mBase = m.M
	if v340 == v342 {
		goto L101
	} else {
		goto L103
	}
L103:
	;
	v344 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v330))))
	v352 = v331
	v353 = v344
	goto L97
L104:
	;
	goto L99
L105:
	;
	v450 = v47
	v451 = v48
	v453 = int32(1)
	v454 = v52
	v455 = v53
	v456 = v279
	goto L10
L106:
	;
	if v400-v402 != 0 {
		goto L5
	} else {
		goto L118
	}
L107:
	;
	v400 = F_tolower(m, v396)
	mBase = m.M
	v401 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v397))))
	v402 = F_tolower(m, v401)
	mBase = m.M
	goto L106
L108:
	;
	v370 = v284
	v371 = v365
	v372 = v368
	goto L111
L109:
	;
	v396 = int32(0)
	v397 = v365
	goto L107
L110:
	;
	v396 = v393 & int32(255)
	v397 = v392
	goto L107
L111:
	;
	v374 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v371))))
	if v374 == int32(0) {
		v392 = v371
		v393 = v372
		goto L110
	} else {
		goto L113
	}
L112:
	;
	v392 = v386
	v393 = int32(0)
	goto L110
L113:
	;
	v378 = v372 & int32(255)
	if v378 == v374 {
		goto L114
	} else {
		goto L115
	}
L114:
	;
	v385 = int32(1)
	v386 = v371 + v385
	v387 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v370)+1)))
	if v387 != 0 {
		v370 = v370 + v385
		v371 = v386
		v372 = v387
		goto L111
	} else {
		goto L117
	}
L115:
	;
	v380 = F_tolower(m, v378)
	mBase = m.M
	v381 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v371))))
	v382 = F_tolower(m, v381)
	mBase = m.M
	if v380 == v382 {
		goto L114
	} else {
		goto L116
	}
L116:
	;
	v384 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v370))))
	v392 = v371
	v393 = v384
	goto L110
L117:
	;
	goto L112
L118:
	;
	v450 = v47
	v451 = v48
	v453 = int32(2)
	v454 = v52
	v455 = v53
	v456 = v279
	goto L10
L119:
	;
	v1719 = v48
	goto L2
L120:
	;
	if v417 != 0 {
		v1719 = v48
		goto L2
	} else {
		goto L121
	}
L121:
	;
	v422 = base.I64_extend_i32_s(v230) + v409 + int64(-1)
	if base.Ui64(v422) < base.Ui64(v53) {
		goto L122
	} else {
		goto L123
	}
L122:
	;
	v424 = v53
	goto L124
L123:
	;
	v424 = v422
	goto L124
L124:
	;
	v426 = v410
	v427 = v424
	v428 = int32(3)
	goto L11
L125:
	;
	v438 = v47<<(uint(v430)%32) + v436
	v439 = *(*int64)(unsafe.Add(mBase, uint32(v27)+56))
	*(*int64)(unsafe.Add(mBase, uint32(v438))) = v439
	v441 = *(*int64)(unsafe.Add(mBase, uint32(v27)))
	*(*int32)(unsafe.Add(mBase, uint32(v438)+20)) = v51
	*(*int32)(unsafe.Add(mBase, uint32(v438)+16)) = v209
	*(*int64)(unsafe.Add(mBase, uint32(v438)+8)) = v441
	v445 = *(*int32)(unsafe.Add(mBase, uint32(v27)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v438)+24)) = v445
	v447 = *(*int32)(unsafe.Add(mBase, uint32(v27)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v438)+28)) = v447
	v450 = v433
	v451 = v436
	v453 = v51
	v454 = v426
	v455 = v427
	v456 = v428 + v50
	goto L10
L126:
	;
	goto L9
L127:
	;
	v471 = v450
	v472 = v451
	v477 = v455
	goto L6
L128:
	;
	if v494 == int32(0) {
		v522 = v3
		v523 = v471
		v524 = v472
		v529 = v477
		goto L3
	} else {
		goto L129
	}
L129:
	;
	v499 = F_checkType(m, l0, v494, int32(0))
	mBase = m.M
	v500 = m.ExcPending
	if v500 != 0 {
		goto L57
	} else {
		goto L130
	}
L130:
	;
	if v499 != 0 {
		v1719 = v472
		goto L2
	} else {
		goto L131
	}
L131:
	;
	v522 = v494
	v523 = v471
	v524 = v472
	v529 = v477
	goto L3
L132:
	;
	v1719 = v48
	goto L2
L133:
	;
	v515 = F_lookupStringForBitCommand(m, l0, v455, v27+int32(52))
	mBase = m.M
	v516 = m.ExcPending
	if v516 != 0 {
		goto L57
	} else {
		goto L137
	}
L134:
	;
	F_valkey_free(m, v451)
	mBase = m.M
	v509 = m.ExcPending
	if v509 != 0 {
		goto L57
	} else {
		goto L135
	}
L135:
	;
	F_addReplyError(m, l0, int32(_a193))
	mBase = m.M
	v512 = m.ExcPending
	if v512 != 0 {
		goto L57
	} else {
		goto L136
	}
L136:
	;
	goto L1
L137:
	;
	if v515 == int32(0) {
		v1719 = v451
		goto L2
	} else {
		goto L138
	}
L138:
	;
	v522 = v515
	v523 = v450
	v524 = v451
	v529 = v455
	goto L3
L139:
	;
	F_addReplyArrayLen(m, l0, v523)
	mBase = m.M
	v546 = m.ExcPending
	if v546 != 0 {
		goto L57
	} else {
		goto L140
	}
L140:
	;
	if v523 < int32(1) {
		goto L141
	} else {
		goto L142
	}
L141:
	;
	F_commitDeferredReplyBuffer(m, l0, int32(1))
	mBase = m.M
	v1713 = m.ExcPending
	if v1713 != 0 {
		goto L57
	} else {
		goto L303
	}
L142:
	;
	v552 = v27 + int32(40)
	v553 = int32(0)
	v565 = v529
	v568 = v553
	v569 = v553
	v573 = v11
	v578 = v11
	goto L143
L143:
	;
	v581 = v524 + v568<<(uint(int32(5))%32)
	v582 = *(*int32)(unsafe.Add(mBase, uint32(v581)+16))
	if base.Ui32(int32(1)) < base.Ui32(v582+int32(-1)) {
		goto L151
	} else {
		goto L152
	}
L144:
	;
	if v1653 == int32(0) {
		goto L141
	} else {
		goto L300
	}
L145:
	;
	v1664 = v568 + int32(1)
	if v1664 != v523 {
		v565 = v1649
		v568 = v1664
		v569 = v1653
		v573 = v1657
		v578 = v1662
		goto L143
	} else {
		goto L299
	}
L146:
	;
	F_addReplyLongLong(m, l0, v1545)
	mBase = m.M
	v1550 = m.ExcPending
	if v1550 != 0 {
		goto L57
	} else {
		goto L293
	}
L147:
	;
	v1544 = v1541
	v1545 = v1542
	v1547 = v1542
	v1548 = v1542
	goto L146
L148:
	;
	F_addReplyLongLong(m, l0, v1446)
	mBase = m.M
	v1452 = m.ExcPending
	if v1452 != 0 {
		goto L57
	} else {
		goto L287
	}
L149:
	;
	v1446 = v1002
	v1450 = v1002
	goto L148
L150:
	;
	v1446 = v992
	v1450 = v998
	goto L148
L151:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+48)) = int32(0)
	if v522 != 0 {
		goto L245
	} else {
		goto L246
	}
L152:
	;
	v587 = *(*int32)(unsafe.Add(mBase, uint32(v581)+28))
	v588 = F_objectGetVal(m, v522)
	mBase = m.M
	v589 = *(*int32)(unsafe.Add(mBase, uint32(v581)+24))
	v590 = base.I64_extend_i32_s(v589)
	if v587 == int32(0) {
		goto L153
	} else {
		goto L154
	}
L153:
	;
	if v589 != 0 {
		goto L216
	} else {
		goto L217
	}
L154:
	;
	if v589 != 0 {
		goto L157
	} else {
		goto L158
	}
L155:
	;
	v768 = int64(-1)
	v770 = int64(1) << (uint(v590+v768) % 64)
	v775 = base.B2i32(v589 != int32(64))
	if v589 != int32(64) {
		goto L169
	} else {
		goto L170
	}
L156:
	;
	if int64(base.Ui64(v727)>>(uint(v590+int64(-1))%64))&int64(1) == int64(0) {
		v761 = v727
		goto L155
	} else {
		goto L168
	}
L157:
	;
	v594 = *(*int64)(unsafe.Add(mBase, uint32(v581)))
	if v589 != int32(1) {
		goto L160
	} else {
		goto L161
	}
L158:
	;
	v727 = int64(0)
	goto L156
L159:
	;
	if v589&int32(1) == int32(0) {
		v706 = v683
		goto L165
	} else {
		goto L166
	}
L160:
	;
	v600 = int64(0)
	v612 = v600
	v618 = v594
	v620 = v600
	goto L162
L161:
	;
	v675 = int64(0)
	v681 = v594
	v683 = v573
	goto L159
L162:
	;
	v626 = int64(1)
	v627 = v618 + v626
	v628 = int64(3)
	v632 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v588+base.I32_wrap_i64(int64(base.Ui64(v627)>>(uint(v628)%64)))))))
	v633 = int64(-1)
	v635 = int64(7)
	v640 = int64(2)
	v646 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v588+base.I32_wrap_i64(int64(base.Ui64(v618)>>(uint(v628)%64)))))))
	v657 = int64(base.Ui64(v632)>>(uint((v627^v633)&v635)%64))&v626 | (v620<<(uint(v640)%64) | int64(base.Ui64(v646)>>(uint((v618^v633)&v635)%64))<<(uint(v626)%64)&v640)
	v659 = v618 + v640
	v661 = v612 + v640
	if v661 != v590&int64(-2) {
		v612 = v661
		v618 = v659
		v620 = v657
		goto L162
	} else {
		goto L164
	}
L163:
	;
	v675 = v657 << (uint(int64(1)) % 64)
	v681 = v659
	v683 = v657
	goto L159
L164:
	;
	goto L163
L165:
	;
	if base.Ui32(int32(63)) < base.Ui32(v589) {
		v761 = v706
		goto L155
	} else {
		goto L167
	}
L166:
	;
	v697 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v588+base.I32_wrap_i64(int64(base.Ui64(v681)>>(uint(int64(3))%64)))))))
	v706 = int64(base.Ui64(v697)>>(uint((v681^int64(-1))&int64(7))%64))&int64(1) | v675
	goto L165
L167:
	;
	v727 = v706
	goto L156
L168:
	;
	v761 = v727 | int64(-1)<<(uint(v590)%64)
	goto L155
L169:
	;
	v776 = v770 + v768
	goto L171
L170:
	;
	v776 = int64(9223372036854775807)
	goto L171
L171:
	;
	v778 = v776 ^ int64(-1)
	v779 = *(*int32)(unsafe.Add(mBase, uint32(v581)+20))
	v780 = *(*int64)(unsafe.Add(mBase, uint32(v581)+8))
	v781 = *(*int32)(unsafe.Add(mBase, uint32(v581)+16))
	if v781 != int32(2) {
		goto L175
	} else {
		goto L176
	}
L172:
	;
	F_addReplyNull(m, l0)
	mBase = m.M
	v860 = m.ExcPending
	if v860 != 0 {
		goto L57
	} else {
		goto L214
	}
L173:
	;
	v1544 = v780
	v1545 = v855
	v1547 = v857
	v1548 = v578
	goto L146
L174:
	;
	if v779 == int32(2) {
		goto L172
	} else {
		goto L213
	}
L175:
	;
	if v776 < v780 {
		goto L201
	} else {
		goto L202
	}
L176:
	;
	if v776 < v761 {
		goto L180
	} else {
		goto L181
	}
L177:
	;
	v821 = v780 + v761
	v855 = v821
	v857 = v821
	goto L173
L178:
	;
	v808 = v780 + v761
	if base.Ui32(int32(63)) < base.Ui32(v589) {
		v1541 = v780
		v1542 = v808
		goto L147
	} else {
		goto L195
	}
L179:
	;
	if v761 < v778 {
		goto L188
	} else {
		goto L189
	}
L180:
	;
	switch v779 {
	case 0:
		goto L178
	case 1:
		v1544 = v780
		v1545 = v776
		v1547 = v776
		v1548 = v776
		goto L146
	default:
		v849 = v578
		goto L174
	}
L181:
	;
	v785 = v776 - v761
	v787 = base.B2i32(v589 == int32(64))
	if v589 == int32(64) {
		goto L182
	} else {
		goto L183
	}
L182:
	;
	if v780 <= v785 {
		goto L179
	} else {
		goto L185
	}
L183:
	;
	if v785 < v780 {
		goto L180
	} else {
		goto L184
	}
L184:
	;
	goto L182
L185:
	;
	if v761 < int64(0) {
		goto L179
	} else {
		goto L186
	}
L186:
	;
	if v780 < int64(1) {
		goto L179
	} else {
		goto L187
	}
L187:
	;
	goto L180
L188:
	;
	switch v779 {
	case 0:
		goto L178
	case 1:
		v1544 = v780
		v1545 = v778
		v1547 = v778
		v1548 = v778
		goto L146
	default:
		v849 = v578
		goto L174
	}
L189:
	;
	v797 = v778 - v761
	if v589 == int32(64) {
		goto L190
	} else {
		goto L191
	}
L190:
	;
	if v797 <= v780 {
		goto L177
	} else {
		goto L193
	}
L191:
	;
	if v780 < v797 {
		goto L188
	} else {
		goto L192
	}
L192:
	;
	goto L190
L193:
	;
	if int64(-1) < v780&v761 {
		goto L177
	} else {
		goto L194
	}
L194:
	;
	goto L188
L195:
	;
	v811 = int64(-1)
	v812 = v811 << (uint(v590) % 64)
	if v808&v770 == int64(0) {
		goto L196
	} else {
		goto L197
	}
L196:
	;
	v820 = v808 & (v812 ^ v811)
	goto L198
L197:
	;
	v820 = v808 | v812
	goto L198
L198:
	;
	v1541 = v812
	v1542 = v820
	goto L147
L199:
	;
	if base.Ui32(v589) <= base.Ui32(int32(63)) {
		goto L208
	} else {
		goto L209
	}
L200:
	;
	switch v779 {
	case 0:
		goto L199
	case 1:
		v1544 = v780
		v1545 = v761
		v1547 = v778
		v1548 = v778
		goto L146
	default:
		v849 = v761
		goto L174
	}
L201:
	;
	switch v779 {
	case 0:
		goto L199
	case 1:
		v1544 = v780
		v1545 = v761
		v1547 = v776
		v1548 = v776
		goto L146
	default:
		v849 = v761
		goto L174
	}
L202:
	;
	if v589 == int32(64) {
		goto L203
	} else {
		goto L204
	}
L203:
	;
	if v780 < v778 {
		goto L200
	} else {
		goto L206
	}
L204:
	;
	if v776-v780 < int64(0) {
		goto L201
	} else {
		goto L205
	}
L205:
	;
	goto L203
L206:
	;
	if v775&base.B2i32(v780+v776 < int64(-1)) != 0 {
		goto L200
	} else {
		goto L207
	}
L207:
	;
	v855 = v761
	v857 = v780
	goto L173
L208:
	;
	v837 = int64(-1)
	v838 = v837 << (uint(v590) % 64)
	if v780&v770 == int64(0) {
		goto L210
	} else {
		goto L211
	}
L209:
	;
	v1544 = v780
	v1545 = v761
	v1547 = v780
	v1548 = v780
	goto L146
L210:
	;
	v846 = v780 & (v838 ^ v837)
	goto L212
L211:
	;
	v846 = v780 | v838
	goto L212
L212:
	;
	v1544 = v780
	v1545 = v761
	v1547 = v846
	v1548 = v846
	goto L146
L213:
	;
	v855 = v849
	v857 = v578
	goto L173
L214:
	;
	v1649 = v780
	v1653 = v569
	v1657 = v761
	v1662 = v578
	goto L145
L215:
	;
	v998 = *(*int64)(unsafe.Add(mBase, uint32(v581)+8))
	v999 = *(*int32)(unsafe.Add(mBase, uint32(v581)+16))
	if v999 != int32(2) {
		goto L226
	} else {
		goto L227
	}
L216:
	;
	v862 = *(*int64)(unsafe.Add(mBase, uint32(v581)))
	if v589 != int32(1) {
		goto L219
	} else {
		goto L220
	}
L217:
	;
	v992 = int64(0)
	goto L215
L218:
	;
	if v589&int32(1) == int32(0) {
		v992 = v951
		goto L215
	} else {
		goto L224
	}
L219:
	;
	v868 = int64(0)
	v880 = v868
	v886 = v862
	v888 = v868
	goto L221
L220:
	;
	v943 = int64(0)
	v949 = v862
	v951 = v573
	goto L218
L221:
	;
	v894 = int64(1)
	v895 = v886 + v894
	v896 = int64(3)
	v900 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v588+base.I32_wrap_i64(int64(base.Ui64(v895)>>(uint(v896)%64)))))))
	v901 = int64(-1)
	v903 = int64(7)
	v908 = int64(2)
	v914 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v588+base.I32_wrap_i64(int64(base.Ui64(v886)>>(uint(v896)%64)))))))
	v925 = int64(base.Ui64(v900)>>(uint((v895^v901)&v903)%64))&v894 | (v888<<(uint(v908)%64) | int64(base.Ui64(v914)>>(uint((v886^v901)&v903)%64))<<(uint(v894)%64)&v908)
	v927 = v886 + v908
	v929 = v880 + v908
	if v929 != v590&int64(-2) {
		v880 = v929
		v886 = v927
		v888 = v925
		goto L221
	} else {
		goto L223
	}
L222:
	;
	v943 = v925 << (uint(int64(1)) % 64)
	v949 = v927
	v951 = v925
	goto L218
L223:
	;
	goto L222
L224:
	;
	v965 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v588+base.I32_wrap_i64(int64(base.Ui64(v949)>>(uint(int64(3))%64)))))))
	v992 = int64(base.Ui64(v965)>>(uint((v949^int64(-1))&int64(7))%64))&int64(1) | v943
	goto L215
L225:
	;
	if v1039 != int32(2) {
		v1446 = v1043
		v1450 = int64(0)
		goto L148
	} else {
		goto L242
	}
L226:
	;
	if v589 == int32(64) {
		goto L150
	} else {
		goto L239
	}
L227:
	;
	v1002 = v998 + v992
	v1003 = *(*int32)(unsafe.Add(mBase, uint32(v581)+20))
	v1004 = int64(-1)
	v1008 = v1004<<(uint(v590)%64) ^ v1004
	if v589 == int32(64) {
		goto L231
	} else {
		goto L232
	}
L228:
	;
	v1029 = v1002 & v1008
	v1446 = v1029
	v1450 = v1029
	goto L148
L229:
	;
	if int64(-1) < v998 {
		goto L149
	} else {
		goto L237
	}
L230:
	;
	switch v1003 {
	case 0:
		goto L228
	case 1:
		v1446 = v1011
		v1450 = v1011
		goto L148
	default:
		v1039 = v1003
		v1043 = int64(0)
		goto L225
	}
L231:
	;
	v1011 = v1004
	goto L233
L232:
	;
	v1011 = v1008
	goto L233
L233:
	;
	if base.Ui64(v1011) < base.Ui64(v992) {
		goto L230
	} else {
		goto L234
	}
L234:
	;
	if v998 < int64(1) {
		goto L229
	} else {
		goto L235
	}
L235:
	;
	if v998 <= v1011-v992 {
		goto L229
	} else {
		goto L236
	}
L236:
	;
	goto L230
L237:
	;
	if int64(0)-v992 <= v998 {
		goto L149
	} else {
		goto L238
	}
L238:
	;
	v1023 = int64(0)
	switch v1003 {
	case 0:
		goto L228
	case 1:
		v1446 = v1023
		v1450 = v1023
		goto L148
	default:
		v1039 = v1003
		v1043 = v1023
		goto L225
	}
L239:
	;
	v1032 = int64(-1)
	v1035 = v1032<<(uint(v590)%64) ^ v1032
	if base.Ui64(v998) <= base.Ui64(v1035) {
		goto L150
	} else {
		goto L240
	}
L240:
	;
	v1037 = *(*int32)(unsafe.Add(mBase, uint32(v581)+20))
	switch v1037 {
	case 0:
		goto L241
	case 1:
		v1446 = v992
		v1450 = v1035
		goto L148
	default:
		v1039 = v1037
		v1043 = v992
		goto L225
	}
L241:
	;
	v1446 = v992
	v1450 = v998 & v1035
	goto L148
L242:
	;
	F_addReplyNull(m, l0)
	mBase = m.M
	v1049 = m.ExcPending
	if v1049 != 0 {
		goto L57
	} else {
		goto L243
	}
L243:
	;
	v1649 = v998
	v1653 = v569
	v1657 = v992
	v1662 = v578
	goto L145
L244:
	;
	v1132 = *(*int32)(unsafe.Add(mBase, uint32(v581)+24))
	v1133 = base.I64_extend_i32_s(v1132)
	v1134 = *(*int32)(unsafe.Add(mBase, uint32(v581)+28))
	if v1134 == int32(0) {
		goto L262
	} else {
		goto L263
	}
L245:
	;
	v1059 = F_getObjectReadOnlyString(m, v522, v27+int32(48), v27)
	mBase = m.M
	v1060 = m.ExcPending
	if v1060 != 0 {
		goto L57
	} else {
		goto L247
	}
L246:
	;
	v1052 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v552))) = uint8(v1052)
	*(*int64)(unsafe.Add(mBase, uint32(v27)+32)) = int64(0)
	v1056 = *(*int64)(unsafe.Add(mBase, uint32(v581)))
	v1128 = v565
	v1130 = v1056
	goto L244
L247:
	;
	v1061 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v552))) = uint8(v1061)
	*(*int64)(unsafe.Add(mBase, uint32(v27)+32)) = int64(0)
	v1065 = *(*int64)(unsafe.Add(mBase, uint32(v581)))
	if v1059 == v1061 {
		v1128 = v565
		v1130 = v1065
		goto L244
	} else {
		goto L248
	}
L248:
	;
	v1068 = int64(*(*int32)(unsafe.Add(mBase, uint32(v27)+48)))
	v1070 = int64(base.Ui64(v1065) >> (uint(int64(3)) % 64))
	if base.Ui64(v1068) <= base.Ui64(v1070) {
		v1128 = v1068
		v1130 = v1065
		goto L244
	} else {
		goto L249
	}
L249:
	;
	v1073 = v1059 + base.I32_wrap_i64(v1070)
	v1074 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1073))))
	*(*uint8)(unsafe.Add(mBase, uint32(v27)+32)) = uint8(v1074)
	v1077 = v1068 - v1070
	if base.Ui64(v1068) < base.Ui64(v1077) {
		goto L250
	} else {
		goto L251
	}
L250:
	;
	v1079 = int64(0)
	goto L252
L251:
	;
	v1079 = v1077
	goto L252
L252:
	;
	if v1079 == int64(1) {
		v1128 = v1079
		v1130 = v1065
		goto L244
	} else {
		goto L253
	}
L253:
	;
	v1084 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1073+int32(1)))))
	*(*uint8)(unsafe.Add(mBase, uint32(v27)+33)) = uint8(v1084)
	if v1079 == int64(2) {
		v1128 = v1079
		v1130 = v1065
		goto L244
	} else {
		goto L254
	}
L254:
	;
	v1090 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1073+int32(2)))))
	*(*uint8)(unsafe.Add(mBase, uint32(v27)+34)) = uint8(v1090)
	if v1079 == int64(3) {
		v1128 = v1079
		v1130 = v1065
		goto L244
	} else {
		goto L255
	}
L255:
	;
	v1096 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1073+int32(3)))))
	*(*uint8)(unsafe.Add(mBase, uint32(v27)+35)) = uint8(v1096)
	if v1079 == int64(4) {
		v1128 = v1079
		v1130 = v1065
		goto L244
	} else {
		goto L256
	}
L256:
	;
	v1102 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1073+int32(4)))))
	*(*uint8)(unsafe.Add(mBase, uint32(v27)+36)) = uint8(v1102)
	if v1079 == int64(5) {
		v1128 = v1079
		v1130 = v1065
		goto L244
	} else {
		goto L257
	}
L257:
	;
	v1108 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1073+int32(5)))))
	*(*uint8)(unsafe.Add(mBase, uint32(v27)+37)) = uint8(v1108)
	if v1079 == int64(6) {
		v1128 = v1079
		v1130 = v1065
		goto L244
	} else {
		goto L258
	}
L258:
	;
	v1114 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1073+int32(6)))))
	*(*uint8)(unsafe.Add(mBase, uint32(v27)+38)) = uint8(v1114)
	if v1079 == int64(7) {
		v1128 = v1079
		v1130 = v1065
		goto L244
	} else {
		goto L259
	}
L259:
	;
	v1120 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1073+int32(7)))))
	*(*uint8)(unsafe.Add(mBase, uint32(v27)+39)) = uint8(v1120)
	if v1079 == int64(8) {
		v1128 = v1079
		v1130 = v1065
		goto L244
	} else {
		goto L260
	}
L260:
	;
	v1126 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1073+int32(8)))))
	*(*uint8)(unsafe.Add(mBase, uint32(v27)+40)) = uint8(v1126)
	v1128 = v1079
	v1130 = v1065
	goto L244
L261:
	;
	F_addReplyLongLong(m, l0, v1428)
	mBase = m.M
	v1443 = m.ExcPending
	if v1443 != 0 {
		goto L57
	} else {
		goto L286
	}
L262:
	;
	if v1132 != 0 {
		goto L277
	} else {
		goto L278
	}
L263:
	;
	if v1132 != 0 {
		goto L265
	} else {
		goto L266
	}
L264:
	;
	if int64(base.Ui64(v1272)>>(uint(v1133+int64(-1))%64))&int64(1) == int64(0) {
		v1428 = v1272
		v1436 = v1280
		goto L261
	} else {
		goto L276
	}
L265:
	;
	v1138 = *(*int64)(unsafe.Add(mBase, uint32(v581)))
	v1141 = v1138 - v1130&int64(-8)
	if v1132 != int32(1) {
		goto L268
	} else {
		goto L269
	}
L266:
	;
	v1272 = int64(0)
	v1280 = v573
	goto L264
L267:
	;
	if v1132&int32(1) == int32(0) {
		v1259 = v1226
		goto L273
	} else {
		goto L274
	}
L268:
	;
	v1146 = v1133 & int64(-2)
	v1147 = int64(0)
	v1159 = v1147
	v1165 = v1141
	v1169 = v1147
	goto L270
L269:
	;
	v1226 = v1128
	v1232 = v1141
	v1234 = v573
	v1236 = int64(0)
	goto L267
L270:
	;
	v1174 = v27 + int32(32)
	v1175 = int64(1)
	v1176 = v1165 + v1175
	v1177 = int64(3)
	v1181 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v1174+base.I32_wrap_i64(int64(base.Ui64(v1176)>>(uint(v1177)%64)))))))
	v1182 = int64(-1)
	v1184 = int64(7)
	v1189 = int64(2)
	v1197 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v1174+base.I32_wrap_i64(int64(base.Ui64(v1165)>>(uint(v1177)%64)))))))
	v1208 = int64(base.Ui64(v1181)>>(uint((v1176^v1182)&v1184)%64))&v1175 | (v1159<<(uint(v1189)%64) | int64(base.Ui64(v1197)>>(uint((v1165^v1182)&v1184)%64))<<(uint(v1175)%64)&v1189)
	v1210 = v1165 + v1189
	v1212 = v1169 + v1189
	if v1212 != v1146 {
		v1159 = v1208
		v1165 = v1210
		v1169 = v1212
		goto L270
	} else {
		goto L272
	}
L271:
	;
	v1226 = v1208
	v1232 = v1210
	v1234 = v1146
	v1236 = v1208 << (uint(int64(1)) % 64)
	goto L267
L272:
	;
	goto L271
L273:
	;
	if base.Ui32(int32(63)) < base.Ui32(v1132) {
		v1428 = v1259
		v1436 = v1234
		goto L261
	} else {
		goto L275
	}
L274:
	;
	v1250 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v27+int32(32)+base.I32_wrap_i64(int64(base.Ui64(v1232)>>(uint(int64(3))%64)))))))
	v1259 = int64(base.Ui64(v1250)>>(uint((v1232^int64(-1))&int64(7))%64))&int64(1) | v1236
	goto L273
L275:
	;
	v1272 = v1259
	v1280 = v1234
	goto L264
L276:
	;
	v1428 = v1272 | int64(-1)<<(uint(v1133)%64)
	v1436 = v1280
	goto L261
L277:
	;
	v1297 = *(*int64)(unsafe.Add(mBase, uint32(v581)))
	v1300 = v1297 - v1130&int64(-8)
	if v1132 != int32(1) {
		goto L280
	} else {
		goto L281
	}
L278:
	;
	v1428 = int64(0)
	v1436 = v573
	goto L261
L279:
	;
	if v1132&int32(1) == int32(0) {
		v1428 = v1385
		v1436 = v1393
		goto L261
	} else {
		goto L285
	}
L280:
	;
	v1305 = v1133 & int64(-2)
	v1306 = int64(0)
	v1318 = v1306
	v1324 = v1300
	v1328 = v1306
	goto L282
L281:
	;
	v1385 = v1128
	v1391 = v1300
	v1393 = v573
	v1395 = int64(0)
	goto L279
L282:
	;
	v1333 = v27 + int32(32)
	v1334 = int64(1)
	v1335 = v1324 + v1334
	v1336 = int64(3)
	v1340 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v1333+base.I32_wrap_i64(int64(base.Ui64(v1335)>>(uint(v1336)%64)))))))
	v1341 = int64(-1)
	v1343 = int64(7)
	v1348 = int64(2)
	v1356 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v1333+base.I32_wrap_i64(int64(base.Ui64(v1324)>>(uint(v1336)%64)))))))
	v1367 = int64(base.Ui64(v1340)>>(uint((v1335^v1341)&v1343)%64))&v1334 | (v1318<<(uint(v1348)%64) | int64(base.Ui64(v1356)>>(uint((v1324^v1341)&v1343)%64))<<(uint(v1334)%64)&v1348)
	v1369 = v1324 + v1348
	v1371 = v1328 + v1348
	if v1371 != v1305 {
		v1318 = v1367
		v1324 = v1369
		v1328 = v1371
		goto L282
	} else {
		goto L284
	}
L283:
	;
	v1385 = v1367
	v1391 = v1369
	v1393 = v1305
	v1395 = v1367 << (uint(int64(1)) % 64)
	goto L279
L284:
	;
	goto L283
L285:
	;
	v1409 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v27+int32(32)+base.I32_wrap_i64(int64(base.Ui64(v1391)>>(uint(int64(3))%64)))))))
	v1428 = int64(base.Ui64(v1409)>>(uint((v1391^int64(-1))&int64(7))%64))&int64(1) | v1395
	v1436 = v1393
	goto L261
L286:
	;
	v1649 = v1428
	v1653 = v569
	v1657 = v1436
	v1662 = v578
	goto L145
L287:
	;
	v1453 = F_objectGetVal(m, v522)
	mBase = m.M
	v1454 = int64(*(*int32)(unsafe.Add(mBase, uint32(v581)+24)))
	if v1454 == int64(0) {
		v1521 = v998
		goto L288
	} else {
		goto L289
	}
L288:
	;
	v1535 = *(*int32)(unsafe.Add(mBase, uint32(v27)+52))
	v1649 = v1521
	v1653 = v569 + (base.B2i32(v1535 != int32(0)) | base.B2i32(v992 != v1450))
	v1657 = v992
	v1662 = v578
	goto L145
L289:
	;
	v1457 = *(*int64)(unsafe.Add(mBase, uint32(v581)))
	v1469 = int64(0)
	v1475 = v1457
	goto L290
L290:
	;
	v1486 = v1453 + base.I32_wrap_i64(int64(base.Ui64(v1475)>>(uint(int64(3))%64)))
	v1487 = int64(-1)
	v1491 = int64(1)
	v1496 = (v1475 ^ v1487) & int64(7)
	v1501 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1486))))
	v1504 = int64(base.Ui64(v1450)>>(uint(v1469^v1487+v1454)%64))&v1491<<(uint(v1496)%64) | base.I64_extend_i32_u(base.I32_rotl(int32(-2), base.I32_wrap_i64(v1496))&v1501)
	*(*uint8)(unsafe.Add(mBase, uint32(v1486))) = uint8(v1504)
	v1509 = v1469 + v1491
	if v1509 != v1454 {
		v1469 = v1509
		v1475 = v1475 + v1491
		goto L290
	} else {
		goto L292
	}
L291:
	;
	v1521 = v1509
	goto L288
L292:
	;
	goto L291
L293:
	;
	v1551 = F_objectGetVal(m, v522)
	mBase = m.M
	v1552 = int64(*(*int32)(unsafe.Add(mBase, uint32(v581)+24)))
	if v1552 == int64(0) {
		v1619 = v1544
		goto L294
	} else {
		goto L295
	}
L294:
	;
	v1633 = *(*int32)(unsafe.Add(mBase, uint32(v27)+52))
	v1649 = v1619
	v1653 = v569 + (base.B2i32(v1633 != int32(0)) | base.B2i32(v761 != v1547))
	v1657 = v761
	v1662 = v1548
	goto L145
L295:
	;
	v1555 = *(*int64)(unsafe.Add(mBase, uint32(v581)))
	v1567 = int64(0)
	v1573 = v1555
	goto L296
L296:
	;
	v1584 = v1551 + base.I32_wrap_i64(int64(base.Ui64(v1573)>>(uint(int64(3))%64)))
	v1585 = int64(-1)
	v1589 = int64(1)
	v1594 = (v1573 ^ v1585) & int64(7)
	v1599 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1584))))
	v1602 = int64(base.Ui64(v1547)>>(uint(v1567^v1585+v1552)%64))&v1589<<(uint(v1594)%64) | base.I64_extend_i32_u(base.I32_rotl(int32(-2), base.I32_wrap_i64(v1594))&v1599)
	*(*uint8)(unsafe.Add(mBase, uint32(v1584))) = uint8(v1602)
	v1607 = v1567 + v1589
	if v1607 != v1552 {
		v1567 = v1607
		v1573 = v1573 + v1589
		goto L296
	} else {
		goto L298
	}
L297:
	;
	v1619 = v1607
	goto L294
L298:
	;
	goto L297
L299:
	;
	goto L144
L300:
	;
	v1668 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v1669 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v1670 = *(*int32)(unsafe.Add(mBase, uint32(v1669)+4))
	F_signalModifiedKey(m, l0, v1668, v1670)
	mBase = m.M
	v1672 = m.ExcPending
	if v1672 != 0 {
		goto L57
	} else {
		goto L301
	}
L301:
	;
	v1675 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v1676 = *(*int32)(unsafe.Add(mBase, uint32(v1675)+4))
	v1677 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v1678 = *(*int32)(unsafe.Add(mBase, uint32(v1677)+28))
	F_notifyKeyspaceEvent(m, int32(8), int32(_a184), v1676, v1678)
	mBase = m.M
	v1680 = m.ExcPending
	if v1680 != 0 {
		goto L57
	} else {
		goto L302
	}
L302:
	;
	v1681 = int32(_a44)
	v1683 = *(*int64)(unsafe.Add(mBase, _consts[83]))
	*(*int64)(unsafe.Add(mBase, _consts[83])) = v1683 + base.I64_extend_i32_s(v1653)
	goto L141
L303:
	;
	v1719 = v524
	goto L2
L304:
	;
	goto L1
}
func F_bitfieldroCommand(m *base.Module, l0 int32) {
	var v4 int32
	_ = v4
	F_bitfieldGeneric(m, l0, int32(1))
	v4 = m.ExcPending
	if v4 != 0 {
		return
	} else {
		return
	}
}
func F_blmpopCommand(m *base.Module, l0 int32) {
	var v5 int32
	_ = v5
	F_lmpopGenericCommand(m, l0, int32(2), int32(1))
	v5 = m.ExcPending
	if v5 != 0 {
		return
	} else {
		return
	}
}
func F_blockClientShutdown(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int64
	_ = v19
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v35 int32
	_ = v35
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+116))
	if v4 != 0 {
		v29 = v4
		*(*int64)(unsafe.Add(mBase, uint32(v29)+8)) = int64(0)
		F_blockClient(m, l0, int32(7))
		mBase = m.M
		v35 = m.ExcPending
		if v35 != 0 {
			return
		} else {
			return
		}
	} else {
		v6 = F_valkey_malloc(m, int32(56))
		mBase = m.M
		v7 = m.ExcPending
		if v7 != 0 {
			return
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(l0)+116)) = v6
			v9 = int32(0)
			*(*int32)(unsafe.Add(mBase, uint32(v6)+16)) = v9
			*(*int64)(unsafe.Add(mBase, uint32(v6)+8)) = int64(0)
			*(*int32)(unsafe.Add(mBase, uint32(v6))) = v9
			v16 = F_dictCreate(m, int32(_a200))
			mBase = m.M
			v17 = m.ExcPending
			if v17 != 0 {
				return
			} else {
				v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+116))
				v19 = int64(0)
				*(*int64)(unsafe.Add(mBase, uint32(v18)+40)) = v19
				*(*int64)(unsafe.Add(mBase, uint32(v18)+28)) = v19
				*(*int32)(unsafe.Add(mBase, uint32(v18)+24)) = v16
				*(*int32)(unsafe.Add(mBase, uint32(v18)+20)) = int32(0)
				v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+116))
				*(*int64)(unsafe.Add(mBase, uint32(v26)+48)) = v19
				v29 = v26
				*(*int64)(unsafe.Add(mBase, uint32(v29)+8)) = int64(0)
				F_blockClient(m, l0, int32(7))
				mBase = m.M
				v35 = m.ExcPending
				if v35 != 0 {
					return
				} else {
					return
				}
			}
		}
	}
}
func F_blockedClientMayTimeout(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
	var v30 int32
	_ = v30
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+116))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(v4)))
	if base.Ui32(int32(5)) < base.Ui32(v5) {
		v30 = int32(0)
		return v30
	} else {
		v8 = int32(1)
		if v8<<(uint(v5)%32)&int32(54) != 0 {
			v30 = v8
			return v30
		} else {
			if v5 != int32(3) {
				v30 = int32(0)
				return v30
			} else {
				v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+116))
				v16 = *(*int32)(unsafe.Add(mBase, uint32(v15)))
				if v16 == int32(3) {
					v20 = *(*int32)(unsafe.Add(mBase, uint32(v15)+48))
					if v20 != 0 {
						v22 = *(*int32)(unsafe.Add(mBase, uint32(v20)+16))
						v26 = base.B2i32(v22 != int32(0))
					} else {
						v26 = int32(0)
					}
				} else {
					v26 = int32(1)
				}
				return v26
			}
		}
	}
}
func F_blockingPopGenericCommand(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v19 int32
	_ = v19
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v32 int32
	_ = v32
	var v38 int32
	_ = v38
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
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v115 int32
	_ = v115
	var v118 int32
	_ = v118
	var v120 int32
	_ = v120
	var v122 int32
	_ = v122
	var v124 int32
	_ = v124
	var v129 int32
	_ = v129
	var v131 int32
	_ = v131
	var v135 int32
	_ = v135
	var v138 int32
	_ = v138
	var v150 int32
	_ = v150
	var v156 int32
	_ = v156
	var v158 int64
	_ = v158
	var v161 int32
	_ = v161
	var v179 int32
	_ = v179
	v11 = m.G0
	v13 = v11 - int32(48)
	m.G0 = v13
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v19 = *(*int32)(unsafe.Add(mBase, uint32(v15+l4<<(uint(int32(2))%32))))
	v23 = F_getTimeoutFromObjectOrReply(m, l0, v19, v13+int32(40), int32(0))
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	F__serverAssert(m, int32(_a1678), int32(_a1674), int32(1203))
	mBase = m.M
	v179 = m.ExcPending
	if v179 != 0 {
		goto L3
	} else {
		goto L61
	}
L2:
	;
	m.G0 = v13 + int32(48)
	return
L3:
	;
	return
L4:
	;
	if v23 != 0 {
		goto L2
	} else {
		goto L5
	}
L5:
	;
	if l2 < int32(1) {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	v150 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+205)))
	if v150&int32(16) == int32(0) {
		goto L57
	} else {
		goto L58
	}
L7:
	;
	v32 = int32(0)
	goto L8
L8:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v42 = *(*int32)(unsafe.Add(mBase, uint32(l1+v32<<(uint(int32(2))%32))))
	v43 = F_lookupKeyWrite(m, v38, v42)
	mBase = m.M
	v44 = m.ExcPending
	if v44 != 0 {
		goto L3
	} else {
		goto L11
	}
L9:
	;
	goto L6
L10:
	;
	v138 = v32 + int32(1)
	if v138 != l2 {
		v32 = v138
		goto L8
	} else {
		goto L56
	}
L11:
	;
	if v43 == int32(0) {
		goto L10
	} else {
		goto L12
	}
L12:
	;
	v48 = F_checkType(m, l0, v43, int32(1))
	mBase = m.M
	v49 = m.ExcPending
	if v49 != 0 {
		goto L3
	} else {
		goto L13
	}
L13:
	;
	if v48 != 0 {
		goto L2
	} else {
		goto L14
	}
L14:
	;
	v50 = F_listTypeLength(m, v43)
	mBase = m.M
	v51 = m.ExcPending
	if v51 != 0 {
		goto L3
	} else {
		goto L15
	}
L15:
	;
	if v50 == int32(0) {
		goto L10
	} else {
		goto L16
	}
L16:
	;
	if l5 == int32(-1) {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	v108 = F_listTypePop(m, v43, l3)
	mBase = m.M
	v109 = m.ExcPending
	if v109 != 0 {
		goto L3
	} else {
		goto L45
	}
L18:
	;
	v56 = F_listTypeLength(m, v43)
	mBase = m.M
	v57 = m.ExcPending
	if v57 != 0 {
		goto L3
	} else {
		goto L19
	}
L19:
	;
	F_initDeferredReplyBuffer(m, l0)
	mBase = m.M
	v59 = m.ExcPending
	if v59 != 0 {
		goto L3
	} else {
		goto L20
	}
L20:
	;
	F_addReplyArrayLen(m, l0, int32(2))
	mBase = m.M
	v62 = m.ExcPending
	if v62 != 0 {
		goto L3
	} else {
		goto L21
	}
L21:
	;
	F_addReplyBulk(m, l0, v42)
	mBase = m.M
	v64 = m.ExcPending
	if v64 != 0 {
		goto L3
	} else {
		goto L22
	}
L22:
	;
	if l5 < v56 {
		goto L23
	} else {
		goto L24
	}
L23:
	;
	v67 = l5
	goto L25
L24:
	;
	v67 = v56
	goto L25
L25:
	;
	if l3 != 0 {
		goto L26
	} else {
		goto L27
	}
L26:
	;
	v70 = int32(0) - v67
	goto L28
L27:
	;
	v70 = int32(0)
	goto L28
L28:
	;
	v71 = int32(-1)
	if l3 != 0 {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	v74 = v71
	goto L31
L30:
	;
	v74 = v67 + v71
	goto L31
L31:
	;
	F_addListRangeReply(m, l0, v43, v70, v74, base.B2i32(l3 != int32(0)))
	mBase = m.M
	v78 = m.ExcPending
	if v78 != 0 {
		goto L3
	} else {
		goto L32
	}
L32:
	;
	F_listTypeDelRange(m, v43, v70, v67)
	mBase = m.M
	v80 = m.ExcPending
	if v80 != 0 {
		goto L3
	} else {
		goto L33
	}
L33:
	;
	F_listElementsRemoved(m, l0, v42, l3, v43, v67, int32(0))
	mBase = m.M
	v83 = m.ExcPending
	if v83 != 0 {
		goto L3
	} else {
		goto L34
	}
L34:
	;
	F_commitDeferredReplyBuffer(m, l0, int32(1))
	mBase = m.M
	v86 = m.ExcPending
	if v86 != 0 {
		goto L3
	} else {
		goto L35
	}
L35:
	;
	if l5 < v50 {
		goto L36
	} else {
		goto L37
	}
L36:
	;
	v88 = l5
	goto L38
L37:
	;
	v88 = v50
	goto L38
L38:
	;
	v90 = F_createStringObjectFromLongLong(m, base.I64_extend_i32_s(v88))
	mBase = m.M
	v91 = m.ExcPending
	if v91 != 0 {
		goto L3
	} else {
		goto L39
	}
L39:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+20)) = v42
	*(*int32)(unsafe.Add(mBase, uint32(v13)+24)) = v90
	if l3 != 0 {
		goto L40
	} else {
		goto L41
	}
L40:
	;
	v97 = int32(248)
	goto L42
L41:
	;
	v97 = int32(252)
	goto L42
L42:
	;
	v99 = *(*int32)(unsafe.Add(mBase, uint32(v97)+uint32(_consts[84])))
	*(*int32)(unsafe.Add(mBase, uint32(v13)+16)) = v99
	F_rewriteClientCommandVector(m, l0, int32(3), v13+int32(16))
	mBase = m.M
	v105 = m.ExcPending
	if v105 != 0 {
		goto L3
	} else {
		goto L43
	}
L43:
	;
	F_decrRefCount(m, v90)
	mBase = m.M
	v107 = m.ExcPending
	if v107 != 0 {
		goto L3
	} else {
		goto L44
	}
L44:
	;
	goto L2
L45:
	;
	if v108 == int32(0) {
		goto L1
	} else {
		goto L46
	}
L46:
	;
	F_listElementsRemoved(m, l0, v42, l3, v43, int32(1), int32(0))
	mBase = m.M
	v115 = m.ExcPending
	if v115 != 0 {
		goto L3
	} else {
		goto L47
	}
L47:
	;
	F_addReplyArrayLen(m, l0, int32(2))
	mBase = m.M
	v118 = m.ExcPending
	if v118 != 0 {
		goto L3
	} else {
		goto L48
	}
L48:
	;
	F_addReplyBulk(m, l0, v42)
	mBase = m.M
	v120 = m.ExcPending
	if v120 != 0 {
		goto L3
	} else {
		goto L49
	}
L49:
	;
	F_addReplyBulk(m, l0, v108)
	mBase = m.M
	v122 = m.ExcPending
	if v122 != 0 {
		goto L3
	} else {
		goto L50
	}
L50:
	;
	F_decrRefCount(m, v108)
	mBase = m.M
	v124 = m.ExcPending
	if v124 != 0 {
		goto L3
	} else {
		goto L51
	}
L51:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+4)) = v42
	if l3 != 0 {
		goto L52
	} else {
		goto L53
	}
L52:
	;
	v129 = int32(248)
	goto L54
L53:
	;
	v129 = int32(252)
	goto L54
L54:
	;
	v131 = *(*int32)(unsafe.Add(mBase, uint32(v129)+uint32(_consts[84])))
	*(*int32)(unsafe.Add(mBase, uint32(v13))) = v131
	F_rewriteClientCommandVector(m, l0, int32(2), v13)
	mBase = m.M
	v135 = m.ExcPending
	if v135 != 0 {
		goto L3
	} else {
		goto L55
	}
L55:
	;
	goto L2
L56:
	;
	goto L9
L57:
	;
	v158 = *(*int64)(unsafe.Add(mBase, uint32(v13)+40))
	F_blockForKeys(m, l0, int32(1), l1, l2, v158, int32(0))
	mBase = m.M
	v161 = m.ExcPending
	if v161 != 0 {
		goto L3
	} else {
		goto L60
	}
L58:
	;
	F_addReplyNullArray(m, l0)
	mBase = m.M
	v156 = m.ExcPending
	if v156 != 0 {
		goto L3
	} else {
		goto L59
	}
L59:
	;
	goto L2
L60:
	;
	goto L2
L61:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_blpopCommand(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v10 int32
	_ = v10
	var v14 int32
	_ = v14
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v10 = int32(-1)
	F_blockingPopGenericCommand(m, l0, v3+int32(4), v6+int32(-2), int32(0), v6+v10, v10)
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		return
	} else {
		return
	}
}
func F_breakCommandHandler(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v26 int32
	_ = v26
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
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
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
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
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v80 int32
	_ = v80
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v91 int32
	_ = v91
	var v97 int32
	_ = v97
	var v100 int32
	_ = v100
	var v106 int32
	_ = v106
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v124 int32
	_ = v124
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v130 int32
	_ = v130
	var v131 int64
	_ = v131
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v148 int32
	_ = v148
	var v150 int32
	_ = v150
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v168 int32
	_ = v168
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
	var v185 int32
	_ = v185
	var v192 int32
	_ = v192
	var v196 int32
	_ = v196
	var v199 int32
	_ = v199
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v224 int32
	_ = v224
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v231 int32
	_ = v231
	var v240 int32
	_ = v240
	var v247 int32
	_ = v247
	var v250 int32
	_ = v250
	var v251 int32
	_ = v251
	var v254 int32
	_ = v254
	var v268 int32
	_ = v268
	var v269 int32
	_ = v269
	var v270 int32
	_ = v270
	var v271 int32
	_ = v271
	var v276 int32
	_ = v276
	var v277 int32
	_ = v277
	var v279 int32
	_ = v279
	var v281 int32
	_ = v281
	var v282 int32
	_ = v282
	var v284 int32
	_ = v284
	var v287 int32
	_ = v287
	var v289 int32
	_ = v289
	var v290 int32
	_ = v290
	var v294 int32
	_ = v294
	var v301 int32
	_ = v301
	var v304 int32
	_ = v304
	var v311 int32
	_ = v311
	var v312 int32
	_ = v312
	var v313 int32
	_ = v313
	var v317 int32
	_ = v317
	var v319 int32
	_ = v319
	var v320 int32
	_ = v320
	var v322 int32
	_ = v322
	var v324 int32
	_ = v324
	var v335 int32
	_ = v335
	var v341 int32
	_ = v341
	var v342 int32
	_ = v342
	var v344 int32
	_ = v344
	var v350 int32
	_ = v350
	var v357 int32
	_ = v357
	var v361 int32
	_ = v361
	var v364 int32
	_ = v364
	var v370 int32
	_ = v370
	var v377 int32
	_ = v377
	var v381 int32
	_ = v381
	var v384 int32
	_ = v384
	var v387 int32
	_ = v387
	var v388 int32
	_ = v388
	var v389 int32
	_ = v389
	var v394 int32
	_ = v394
	var v395 int32
	_ = v395
	var v396 int32
	_ = v396
	var v398 int32
	_ = v398
	var v400 int32
	_ = v400
	var v401 int32
	_ = v401
	var v403 int32
	_ = v403
	var v405 int32
	_ = v405
	var v409 int32
	_ = v409
	var v410 int32
	_ = v410
	var v411 int32
	_ = v411
	var v416 int32
	_ = v416
	var v417 int32
	_ = v417
	var v418 int32
	_ = v418
	var v420 int32
	_ = v420
	var v422 int32
	_ = v422
	var v427 int32
	_ = v427
	var v439 int32
	_ = v439
	var v440 int32
	_ = v440
	var v441 int32
	_ = v441
	var v442 int32
	_ = v442
	var v447 int32
	_ = v447
	var v448 int32
	_ = v448
	var v450 int32
	_ = v450
	var v452 int32
	_ = v452
	var v465 int32
	_ = v465
	var v466 int32
	_ = v466
	var v476 int32
	_ = v476
	var v485 int32
	_ = v485
	var v486 int32
	_ = v486
	var v491 int32
	_ = v491
	var v493 int32
	_ = v493
	var v498 int32
	_ = v498
	var v499 int32
	_ = v499
	var v500 int32
	_ = v500
	var v501 int32
	_ = v501
	var v507 int32
	_ = v507
	var v508 int32
	_ = v508
	var v520 int32
	_ = v520
	var v527 int32
	_ = v527
	var v528 int32
	_ = v528
	var v532 int32
	_ = v532
	var v534 int32
	_ = v534
	var v536 int32
	_ = v536
	var v551 int32
	_ = v551
	var v553 int32
	_ = v553
	var v560 int32
	_ = v560
	var v561 int32
	_ = v561
	var v563 int32
	_ = v563
	var v565 int32
	_ = v565
	var v566 int32
	_ = v566
	var v567 int32
	_ = v567
	var v568 int32
	_ = v568
	var v574 int32
	_ = v574
	var v575 int32
	_ = v575
	var v577 int32
	_ = v577
	var v579 int32
	_ = v579
	var v580 int32
	_ = v580
	var v588 int32
	_ = v588
	var v609 int32
	_ = v609
	var v611 int32
	_ = v611
	var v612 int32
	_ = v612
	var v613 int32
	_ = v613
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
	var v624 int32
	_ = v624
	var v626 int32
	_ = v626
	var v628 int32
	_ = v628
	var v629 int32
	_ = v629
	var v634 int32
	_ = v634
	var v635 int32
	_ = v635
	var v636 int32
	_ = v636
	var v637 int32
	_ = v637
	var v638 int32
	_ = v638
	var v647 int32
	_ = v647
	var v650 int32
	_ = v650
	var v654 int32
	_ = v654
	var v659 int32
	_ = v659
	var v661 int32
	_ = v661
	var v665 int32
	_ = v665
	var v671 int32
	_ = v671
	var v674 int32
	_ = v674
	var v680 int32
	_ = v680
	var v684 int32
	_ = v684
	var v686 int32
	_ = v686
	var v694 int32
	_ = v694
	var v695 int32
	_ = v695
	var v696 int32
	_ = v696
	var v698 int32
	_ = v698
	var v700 int32
	_ = v700
	var v701 int32
	_ = v701
	var v702 int32
	_ = v702
	var v704 int32
	_ = v704
	var v705 int32
	_ = v705
	var v706 int32
	_ = v706
	var v713 int32
	_ = v713
	var v719 int32
	_ = v719
	var v721 int32
	_ = v721
	var v723 int32
	_ = v723
	var v724 int32
	_ = v724
	var v728 int32
	_ = v728
	var v731 int32
	_ = v731
	var v733 int32
	_ = v733
	var v734 int32
	_ = v734
	var v740 int32
	_ = v740
	var v741 int32
	_ = v741
	var v742 int32
	_ = v742
	var v743 int32
	_ = v743
	var v755 int32
	_ = v755
	var v762 int32
	_ = v762
	var v763 int32
	_ = v763
	var v767 int32
	_ = v767
	var v769 int32
	_ = v769
	var v771 int32
	_ = v771
	var v785 int32
	_ = v785
	var v786 int32
	_ = v786
	var v795 int32
	_ = v795
	var v796 int32
	_ = v796
	var v798 int32
	_ = v798
	var v800 int32
	_ = v800
	var v801 int32
	_ = v801
	var v802 int32
	_ = v802
	var v803 int32
	_ = v803
	var v807 int32
	_ = v807
	var v808 int32
	_ = v808
	var v810 int32
	_ = v810
	var v812 int32
	_ = v812
	var v814 int32
	_ = v814
	var v815 int32
	_ = v815
	var v829 int32
	_ = v829
	var v830 int32
	_ = v830
	var v832 int32
	_ = v832
	v13 = m.G0
	v15 = v13 - int32(64)
	m.G0 = v15
	v17 = int32(1)
	if l1 == v17 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	v829 = m.G14
	v830 = *(*int32)(unsafe.Add(mBase, uint32(v829)))
	m.T0[v830].(func(*base.Module))(m)
	mBase = m.M
	v832 = m.ExcPending
	if v832 != 0 {
		goto L9
	} else {
		goto L170
	}
L2:
	;
	v611 = m.G6
	v612 = *(*int32)(unsafe.Add(mBase, uint32(v611)+260))
	if v612 != 0 {
		goto L126
	} else {
		goto L127
	}
L3:
	;
	if l1 <= int32(1) {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v26 = v17
	goto L5
L5:
	;
	v36 = l0 + v26<<(uint(int32(2))%32)
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v36)))
	v40 = m.G18
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v40)))
	v42 = m.T0[v41].(func(*base.Module, int32, int32) int32)(m, v37, v15+int32(56))
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L9
	} else {
		goto L10
	}
L7:
	;
	v609 = v26 + int32(1)
	if v609 != l1 {
		v26 = v609
		goto L5
	} else {
		goto L125
	}
L8:
	;
	v131 = *(*int64)(unsafe.Add(mBase, uint32(v15)+56))
	if v131 != int64(0) {
		goto L33
	} else {
		goto L34
	}
L9:
	;
	return int32(0)
L10:
	;
	if v42 == int32(0) {
		goto L8
	} else {
		goto L11
	}
L11:
	;
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v36)))
	v50 = m.G7
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v50)))
	v52 = m.T0[v51].(func(*base.Module, int32, int32) int32)(m, v48, int32(0))
	mBase = m.M
	v53 = m.ExcPending
	if v53 != 0 {
		goto L9
	} else {
		goto L12
	}
L12:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+48)) = v52
	v55 = m.G3
	v56 = m.G13
	v61 = F_lm_asprintf(m, v55+int32(_a1925), v15+int32(48))
	mBase = m.M
	v62 = m.ExcPending
	if v62 != 0 {
		goto L9
	} else {
		goto L13
	}
L13:
	;
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v56)))
	v64 = m.G12
	v65 = int32(0)
	if v61&int32(3) == v65 {
		v87 = v61
		goto L16
	} else {
		goto L17
	}
L14:
	;
	v121 = m.T0[v63].(func(*base.Module, int32, int32, int32) int32)(m, v65, v61, v120)
	mBase = m.M
	v122 = m.ExcPending
	if v122 != 0 {
		goto L9
	} else {
		goto L30
	}
L15:
	;
	v120 = v112 - v61
	goto L14
L16:
	;
	v91 = v87
	goto L24
L17:
	;
	v73 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v61))))
	if v73 != 0 {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	v76 = v61
	goto L20
L19:
	;
	v120 = v61 - v61
	goto L14
L20:
	;
	v80 = v76 + int32(1)
	if v80&int32(3) == int32(0) {
		v87 = v80
		goto L16
	} else {
		goto L22
	}
L22:
	;
	v85 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v80))))
	if v85 != 0 {
		v76 = v80
		goto L20
	} else {
		goto L23
	}
L23:
	;
	v112 = v80
	goto L15
L24:
	;
	v97 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
	v100 = int32(-2139062144)
	if (int32(16843008)-v97|v97)&v100 == v100 {
		v91 = v91 + int32(4)
		goto L24
	} else {
		goto L26
	}
L25:
	;
	v106 = v91
	goto L27
L26:
	;
	goto L25
L27:
	;
	v110 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v106))))
	if v110 != 0 {
		v106 = v106 + int32(1)
		goto L27
	} else {
		goto L29
	}
L28:
	;
	v112 = v106
	goto L15
L29:
	;
	goto L28
L30:
	;
	v124 = *(*int32)(unsafe.Add(mBase, uint32(v64)))
	m.T0[v124].(func(*base.Module, int32, int32))(m, v121, int32(0))
	mBase = m.M
	v126 = m.ExcPending
	if v126 != 0 {
		goto L9
	} else {
		goto L31
	}
L31:
	;
	v127 = m.G11
	v128 = *(*int32)(unsafe.Add(mBase, uint32(v127)))
	m.T0[v128].(func(*base.Module, int32))(m, v61)
	mBase = m.M
	v130 = m.ExcPending
	if v130 != 0 {
		goto L9
	} else {
		goto L32
	}
L32:
	;
	goto L7
L33:
	;
	if v131 < int64(1) {
		goto L38
	} else {
		goto L39
	}
L34:
	;
	v134 = m.G6
	v135 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v134)+260)) = v135
	v137 = m.G3
	v138 = m.G13
	v139 = *(*int32)(unsafe.Add(mBase, uint32(v138)))
	v140 = m.G12
	v145 = m.T0[v139].(func(*base.Module, int32, int32, int32) int32)(m, v135, v137+int32(_a1926), int32(24))
	mBase = m.M
	v146 = m.ExcPending
	if v146 != 0 {
		goto L9
	} else {
		goto L35
	}
L35:
	;
	v148 = *(*int32)(unsafe.Add(mBase, uint32(v140)))
	m.T0[v148].(func(*base.Module, int32, int32))(m, v145, int32(0))
	mBase = m.M
	v150 = m.ExcPending
	if v150 != 0 {
		goto L9
	} else {
		goto L36
	}
L36:
	;
	goto L7
L37:
	;
	v465 = int32(1)
	v466 = m.G6
	*(*int32)(unsafe.Add(mBase, uint32(v466)+260)) = v154 + v465
	*(*int32)(unsafe.Add(mBase, uint32(v466+v154<<(uint(int32(2))%32)+int32(4)))) = v172
	v476 = *(*int32)(unsafe.Add(mBase, uint32(v466)+276))
	if v476 < v465 {
		goto L7
	} else {
		goto L106
	}
L38:
	;
	v227 = m.G6
	v228 = *(*int32)(unsafe.Add(mBase, uint32(v227)+260))
	if v228 < int32(1) {
		goto L55
	} else {
		goto L56
	}
L39:
	;
	v153 = m.G6
	v154 = *(*int32)(unsafe.Add(mBase, uint32(v153)+260))
	if v154 != int32(64) {
		goto L40
	} else {
		goto L41
	}
L40:
	;
	v171 = m.G6
	v172 = base.I32_wrap_i64(v131)
	if v172 < int32(1) {
		goto L44
	} else {
		goto L45
	}
L41:
	;
	v157 = m.G3
	v158 = m.G13
	v159 = *(*int32)(unsafe.Add(mBase, uint32(v158)))
	v160 = m.G12
	v165 = m.T0[v159].(func(*base.Module, int32, int32, int32) int32)(m, int32(0), v157+int32(_a1927), int32(25))
	mBase = m.M
	v166 = m.ExcPending
	if v166 != 0 {
		goto L9
	} else {
		goto L42
	}
L42:
	;
	v168 = *(*int32)(unsafe.Add(mBase, uint32(v160)))
	m.T0[v168].(func(*base.Module, int32, int32))(m, v165, int32(0))
	mBase = m.M
	v170 = m.ExcPending
	if v170 != 0 {
		goto L9
	} else {
		goto L43
	}
L43:
	;
	goto L7
L44:
	;
	v213 = m.G3
	v214 = m.G13
	v215 = *(*int32)(unsafe.Add(mBase, uint32(v214)))
	v216 = m.G12
	v221 = m.T0[v215].(func(*base.Module, int32, int32, int32) int32)(m, int32(0), v213+int32(_a1928), int32(18))
	mBase = m.M
	v222 = m.ExcPending
	if v222 != 0 {
		goto L9
	} else {
		goto L52
	}
L45:
	;
	v175 = *(*int32)(unsafe.Add(mBase, uint32(v171)+276))
	if v175 < v172 {
		goto L44
	} else {
		goto L46
	}
L46:
	;
	v177 = int32(0)
	if v154 <= v177 {
		goto L37
	} else {
		goto L47
	}
L47:
	;
	v185 = v177
	goto L48
L48:
	;
	v192 = m.G6
	v196 = *(*int32)(unsafe.Add(mBase, uint32(v192+v185<<(uint(int32(2))%32))+4))
	if v196 == v172 {
		goto L44
	} else {
		goto L50
	}
L50:
	;
	v199 = v185 + int32(1)
	if v199 == v154 {
		goto L37
	} else {
		goto L51
	}
L51:
	;
	v185 = v199
	goto L48
L52:
	;
	v224 = *(*int32)(unsafe.Add(mBase, uint32(v216)))
	m.T0[v224].(func(*base.Module, int32, int32))(m, v221, int32(0))
	mBase = m.M
	v226 = m.ExcPending
	if v226 != 0 {
		goto L9
	} else {
		goto L53
	}
L53:
	;
	goto L7
L54:
	;
	v282 = m.G6
	v284 = v228 + int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(v282)+260)) = v284
	v287 = v250 + int32(4)
	v289 = v250 + int32(8)
	v290 = v284 - v240
	if v287 == v289 {
		goto L64
	} else {
		goto L65
	}
L55:
	;
	v268 = m.G3
	v269 = m.G13
	v270 = *(*int32)(unsafe.Add(mBase, uint32(v269)))
	v271 = m.G12
	v276 = m.T0[v270].(func(*base.Module, int32, int32, int32) int32)(m, int32(0), v268+int32(_a1929), int32(36))
	mBase = m.M
	v277 = m.ExcPending
	if v277 != 0 {
		goto L9
	} else {
		goto L61
	}
L56:
	;
	v231 = int32(0)
	v240 = v231
	goto L57
L57:
	;
	v247 = m.G6
	v250 = v247 + v240<<(uint(int32(2))%32)
	v251 = *(*int32)(unsafe.Add(mBase, uint32(v250)+4))
	if v251 == v231-base.I32_wrap_i64(v131) {
		goto L54
	} else {
		goto L59
	}
L58:
	;
	goto L55
L59:
	;
	v254 = v240 + int32(1)
	if v254 != v228 {
		v240 = v254
		goto L57
	} else {
		goto L60
	}
L60:
	;
	goto L58
L61:
	;
	v279 = *(*int32)(unsafe.Add(mBase, uint32(v271)))
	m.T0[v279].(func(*base.Module, int32, int32))(m, v276, int32(0))
	mBase = m.M
	v281 = m.ExcPending
	if v281 != 0 {
		goto L9
	} else {
		goto L62
	}
L62:
	;
	goto L7
L63:
	;
	v439 = m.G3
	v440 = m.G13
	v441 = *(*int32)(unsafe.Add(mBase, uint32(v440)))
	v442 = m.G12
	v447 = m.T0[v441].(func(*base.Module, int32, int32, int32) int32)(m, int32(0), v439+int32(_a1930), int32(19))
	mBase = m.M
	v448 = m.ExcPending
	if v448 != 0 {
		goto L9
	} else {
		goto L104
	}
L64:
	;
	goto L63
L65:
	;
	v294 = v290 + v287
	if base.Ui32(int32(0)-v290<<(uint(int32(1))%32)) < base.Ui32(v289-v294) {
		goto L66
	} else {
		goto L67
	}
L66:
	;
	v304 = (v289 ^ v287) & int32(3)
	if base.Ui32(v289) <= base.Ui32(v287) {
		goto L70
	} else {
		goto L71
	}
L67:
	;
	v301 = F___memcpy(m, v287, v289, v290)
	mBase = m.M
	goto L63
L68:
	;
	if v410 == int32(0) {
		goto L64
	} else {
		goto L100
	}
L69:
	;
	if base.Ui32(v388) <= base.Ui32(int32(3)) {
		v409 = v387
		v410 = v388
		v411 = v389
		goto L68
	} else {
		goto L96
	}
L70:
	;
	if v304 != 0 {
		v370 = v290
		goto L80
	} else {
		goto L81
	}
L71:
	;
	if v304 == int32(0) {
		goto L72
	} else {
		goto L73
	}
L72:
	;
	if v287&int32(3) != 0 {
		goto L74
	} else {
		goto L75
	}
L73:
	;
	v409 = v289
	v410 = v290
	v411 = v287
	goto L68
L74:
	;
	v311 = v289
	v312 = v290
	v313 = v287
	goto L76
L75:
	;
	v387 = v289
	v388 = v290
	v389 = v287
	goto L69
L76:
	;
	if v312 == int32(0) {
		goto L64
	} else {
		goto L78
	}
L78:
	;
	v317 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v311))))
	*(*uint8)(unsafe.Add(mBase, uint32(v313))) = uint8(v317)
	v319 = int32(1)
	v320 = v311 + v319
	v322 = v312 + int32(-1)
	v324 = v313 + v319
	if v324&int32(3) == int32(0) {
		v387 = v320
		v388 = v322
		v389 = v324
		goto L69
	} else {
		goto L79
	}
L79:
	;
	v311 = v320
	v312 = v322
	v313 = v324
	goto L76
L80:
	;
	if v370 == int32(0) {
		goto L64
	} else {
		goto L92
	}
L81:
	;
	if v294&int32(3) == int32(0) {
		v350 = v290
		goto L82
	} else {
		goto L83
	}
L82:
	;
	if base.Ui32(v350) <= base.Ui32(int32(3)) {
		v370 = v350
		goto L80
	} else {
		goto L88
	}
L83:
	;
	v335 = v290
	goto L84
L84:
	;
	if v335 == int32(0) {
		goto L64
	} else {
		goto L86
	}
L85:
	;
	v350 = v341
	goto L82
L86:
	;
	v341 = v335 + int32(-1)
	v342 = v287 + v341
	v344 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v289+v341))))
	*(*uint8)(unsafe.Add(mBase, uint32(v342))) = uint8(v344)
	if v342&int32(3) != 0 {
		v335 = v341
		goto L84
	} else {
		goto L87
	}
L87:
	;
	goto L85
L88:
	;
	v357 = v350
	goto L89
L89:
	;
	v361 = v357 + int32(-4)
	v364 = *(*int32)(unsafe.Add(mBase, uint32(v289+v361)))
	*(*int32)(unsafe.Add(mBase, uint32(v287+v361))) = v364
	if base.Ui32(int32(3)) < base.Ui32(v361) {
		v357 = v361
		goto L89
	} else {
		goto L91
	}
L90:
	;
	v370 = v361
	goto L80
L91:
	;
	goto L90
L92:
	;
	v377 = v370
	goto L93
L93:
	;
	v381 = v377 + int32(-1)
	v384 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v289+v381))))
	*(*uint8)(unsafe.Add(mBase, uint32(v287+v381))) = uint8(v384)
	if v381 != 0 {
		v377 = v381
		goto L93
	} else {
		goto L95
	}
L95:
	;
	goto L64
L96:
	;
	v394 = v387
	v395 = v388
	v396 = v389
	goto L97
L97:
	;
	v398 = *(*int32)(unsafe.Add(mBase, uint32(v394)))
	*(*int32)(unsafe.Add(mBase, uint32(v396))) = v398
	v400 = int32(4)
	v401 = v394 + v400
	v403 = v396 + v400
	v405 = v395 + int32(-4)
	if base.Ui32(int32(3)) < base.Ui32(v405) {
		v394 = v401
		v395 = v405
		v396 = v403
		goto L97
	} else {
		goto L99
	}
L98:
	;
	v409 = v401
	v410 = v405
	v411 = v403
	goto L68
L99:
	;
	goto L98
L100:
	;
	v416 = v409
	v417 = v410
	v418 = v411
	goto L101
L101:
	;
	v420 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v416))))
	*(*uint8)(unsafe.Add(mBase, uint32(v418))) = uint8(v420)
	v422 = int32(1)
	v427 = v417 + int32(-1)
	if v427 != 0 {
		v416 = v416 + v422
		v417 = v427
		v418 = v418 + v422
		goto L101
	} else {
		goto L103
	}
L102:
	;
	goto L64
L103:
	;
	goto L102
L104:
	;
	v450 = *(*int32)(unsafe.Add(mBase, uint32(v442)))
	m.T0[v450].(func(*base.Module, int32, int32))(m, v447, int32(0))
	mBase = m.M
	v452 = m.ExcPending
	if v452 != 0 {
		goto L9
	} else {
		goto L105
	}
L105:
	;
	goto L7
L106:
	;
	v485 = v465
	v486 = v476
	goto L107
L107:
	;
	v491 = v172 - v485
	v493 = v491 >> (uint(int32(31)) % 32)
	if base.Ui32(int32(1)) < base.Ui32(v491^v493-v493) {
		v588 = v486
		goto L109
	} else {
		goto L110
	}
L108:
	;
	goto L7
L109:
	;
	if v485 < v588 {
		v485 = v485 + int32(1)
		v486 = v588
		goto L107
	} else {
		goto L124
	}
L110:
	;
	v498 = m.G3
	v499 = int32(0)
	v500 = m.G6
	v501 = *(*int32)(unsafe.Add(mBase, uint32(v500)+272))
	v507 = *(*int32)(unsafe.Add(mBase, uint32(v501+v485<<(uint(int32(2))%32)+int32(-4))))
	v508 = *(*int32)(unsafe.Add(mBase, uint32(v500)+260))
	if v499 < v508 {
		goto L112
	} else {
		goto L113
	}
L111:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+40)) = v507
	*(*int32)(unsafe.Add(mBase, uint32(v15)+36)) = v485
	v560 = m.G6
	v561 = *(*int32)(unsafe.Add(mBase, uint32(v560)+280))
	if v561 == v485 {
		goto L119
	} else {
		goto L120
	}
L112:
	;
	v520 = v499
	goto L115
L113:
	;
	v551 = v498 + int32(_a1931)
	v553 = v498 + int32(_a1932)
	goto L111
L114:
	;
	v551 = v527 + int32(_a1933)
	v553 = v527 + int32(_a1934)
	goto L111
L115:
	;
	v527 = m.G3
	v528 = m.G6
	v532 = *(*int32)(unsafe.Add(mBase, uint32(v528+v520<<(uint(int32(2))%32))+4))
	if v532 == v485 {
		goto L114
	} else {
		goto L117
	}
L116:
	;
	v551 = v534 + int32(_a1931)
	v553 = v534 + int32(_a1932)
	goto L111
L117:
	;
	v534 = m.G3
	v536 = v520 + int32(1)
	if v536 != v508 {
		v520 = v536
		goto L115
	} else {
		goto L118
	}
L118:
	;
	goto L116
L119:
	;
	v563 = v551
	goto L121
L120:
	;
	v563 = v553
	goto L121
L121:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+32)) = v563
	v565 = m.G3
	v566 = m.G15
	v567 = *(*int32)(unsafe.Add(mBase, uint32(v566)))
	v568 = m.G12
	v574 = m.T0[v567].(func(*base.Module, int32, int32, int32) int32)(m, int32(0), v565+int32(_a1935), v15+int32(32))
	mBase = m.M
	v575 = m.ExcPending
	if v575 != 0 {
		goto L9
	} else {
		goto L122
	}
L122:
	;
	v577 = *(*int32)(unsafe.Add(mBase, uint32(v568)))
	m.T0[v577].(func(*base.Module, int32, int32))(m, v574, int32(0))
	mBase = m.M
	v579 = m.ExcPending
	if v579 != 0 {
		goto L9
	} else {
		goto L123
	}
L123:
	;
	v580 = *(*int32)(unsafe.Add(mBase, uint32(v560)+276))
	v588 = v580
	goto L109
L124:
	;
	goto L108
L125:
	;
	goto L1
L126:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+16)) = v612
	v628 = m.G3
	v629 = m.G13
	v634 = F_lm_asprintf(m, v628+int32(_a1936), v15+int32(16))
	mBase = m.M
	v635 = m.ExcPending
	if v635 != 0 {
		goto L9
	} else {
		goto L130
	}
L127:
	;
	v613 = m.G3
	v614 = m.G13
	v615 = *(*int32)(unsafe.Add(mBase, uint32(v614)))
	v616 = m.G12
	v621 = m.T0[v615].(func(*base.Module, int32, int32, int32) int32)(m, int32(0), v613+int32(_a1937), int32(46))
	mBase = m.M
	v622 = m.ExcPending
	if v622 != 0 {
		goto L9
	} else {
		goto L128
	}
L128:
	;
	v624 = *(*int32)(unsafe.Add(mBase, uint32(v616)))
	m.T0[v624].(func(*base.Module, int32, int32))(m, v621, int32(0))
	mBase = m.M
	v626 = m.ExcPending
	if v626 != 0 {
		goto L9
	} else {
		goto L129
	}
L129:
	;
	goto L1
L130:
	;
	v636 = *(*int32)(unsafe.Add(mBase, uint32(v629)))
	v637 = m.G12
	v638 = int32(0)
	if v634&int32(3) == v638 {
		v661 = v634
		goto L133
	} else {
		goto L134
	}
L131:
	;
	v695 = m.T0[v636].(func(*base.Module, int32, int32, int32) int32)(m, v638, v634, v694)
	mBase = m.M
	v696 = m.ExcPending
	if v696 != 0 {
		goto L9
	} else {
		goto L147
	}
L132:
	;
	v694 = v686 - v634
	goto L131
L133:
	;
	v665 = v661
	goto L141
L134:
	;
	v647 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v634))))
	if v647 != 0 {
		goto L135
	} else {
		goto L136
	}
L135:
	;
	v650 = v634
	goto L137
L136:
	;
	v694 = v634 - v634
	goto L131
L137:
	;
	v654 = v650 + int32(1)
	if v654&int32(3) == int32(0) {
		v661 = v654
		goto L133
	} else {
		goto L139
	}
L139:
	;
	v659 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v654))))
	if v659 != 0 {
		v650 = v654
		goto L137
	} else {
		goto L140
	}
L140:
	;
	v686 = v654
	goto L132
L141:
	;
	v671 = *(*int32)(unsafe.Add(mBase, uint32(v665)))
	v674 = int32(-2139062144)
	if (int32(16843008)-v671|v671)&v674 == v674 {
		v665 = v665 + int32(4)
		goto L141
	} else {
		goto L143
	}
L142:
	;
	v680 = v665
	goto L144
L143:
	;
	goto L142
L144:
	;
	v684 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v680))))
	if v684 != 0 {
		v680 = v680 + int32(1)
		goto L144
	} else {
		goto L146
	}
L145:
	;
	v686 = v680
	goto L132
L146:
	;
	goto L145
L147:
	;
	v698 = *(*int32)(unsafe.Add(mBase, uint32(v637)))
	m.T0[v698].(func(*base.Module, int32, int32))(m, v695, int32(0))
	mBase = m.M
	v700 = m.ExcPending
	if v700 != 0 {
		goto L9
	} else {
		goto L148
	}
L148:
	;
	v701 = m.G11
	v702 = *(*int32)(unsafe.Add(mBase, uint32(v701)))
	m.T0[v702].(func(*base.Module, int32))(m, v634)
	mBase = m.M
	v704 = m.ExcPending
	if v704 != 0 {
		goto L9
	} else {
		goto L149
	}
L149:
	;
	v705 = m.G6
	v706 = *(*int32)(unsafe.Add(mBase, uint32(v705)+260))
	if v706 < int32(1) {
		goto L1
	} else {
		goto L150
	}
L150:
	;
	v713 = v706
	v719 = v638
	goto L151
L151:
	;
	v721 = m.G3
	v723 = v721 + int32(_a1938)
	v724 = m.G6
	v728 = *(*int32)(unsafe.Add(mBase, uint32(v724+v719<<(uint(int32(2))%32))+4))
	if v728 < int32(1) {
		v741 = v723
		goto L153
	} else {
		goto L154
	}
L152:
	;
	goto L1
L153:
	;
	v742 = m.G3
	v743 = int32(0)
	if v743 < v713 {
		goto L157
	} else {
		goto L158
	}
L154:
	;
	v731 = *(*int32)(unsafe.Add(mBase, uint32(v724)+276))
	if v731 < v728 {
		v741 = v723
		goto L153
	} else {
		goto L155
	}
L155:
	;
	v733 = m.G6
	v734 = *(*int32)(unsafe.Add(mBase, uint32(v733)+272))
	v740 = *(*int32)(unsafe.Add(mBase, uint32(v734+v728<<(uint(int32(2))%32)+int32(-4))))
	v741 = v740
	goto L153
L156:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+8)) = v741
	*(*int32)(unsafe.Add(mBase, uint32(v15)+4)) = v728
	v795 = m.G6
	v796 = *(*int32)(unsafe.Add(mBase, uint32(v795)+280))
	if v796 == v728 {
		goto L164
	} else {
		goto L165
	}
L157:
	;
	v755 = v743
	goto L160
L158:
	;
	v785 = v742 + int32(_a1932)
	v786 = v742 + int32(_a1931)
	goto L156
L159:
	;
	v785 = v762 + int32(_a1934)
	v786 = v762 + int32(_a1933)
	goto L156
L160:
	;
	v762 = m.G3
	v763 = m.G6
	v767 = *(*int32)(unsafe.Add(mBase, uint32(v763+v755<<(uint(int32(2))%32))+4))
	if v767 == v728 {
		goto L159
	} else {
		goto L162
	}
L161:
	;
	v785 = v769 + int32(_a1932)
	v786 = v769 + int32(_a1931)
	goto L156
L162:
	;
	v769 = m.G3
	v771 = v755 + int32(1)
	if v771 != v713 {
		v755 = v771
		goto L160
	} else {
		goto L163
	}
L163:
	;
	goto L161
L164:
	;
	v798 = v786
	goto L166
L165:
	;
	v798 = v785
	goto L166
L166:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15))) = v798
	v800 = m.G3
	v801 = m.G15
	v802 = *(*int32)(unsafe.Add(mBase, uint32(v801)))
	v803 = m.G12
	v807 = m.T0[v802].(func(*base.Module, int32, int32, int32) int32)(m, int32(0), v800+int32(_a1935), v15)
	mBase = m.M
	v808 = m.ExcPending
	if v808 != 0 {
		goto L9
	} else {
		goto L167
	}
L167:
	;
	v810 = *(*int32)(unsafe.Add(mBase, uint32(v803)))
	m.T0[v810].(func(*base.Module, int32, int32))(m, v807, int32(0))
	mBase = m.M
	v812 = m.ExcPending
	if v812 != 0 {
		goto L9
	} else {
		goto L168
	}
L168:
	;
	v814 = v719 + int32(1)
	v815 = *(*int32)(unsafe.Add(mBase, uint32(v795)+260))
	if v814 < v815 {
		v713 = v815
		v719 = v814
		goto L151
	} else {
		goto L169
	}
L169:
	;
	goto L152
L170:
	;
	m.G0 = v15 + int32(64)
	return int32(1)
}
func F_brpopCommand(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v10 int32
	_ = v10
	var v14 int32
	_ = v14
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v10 = int32(-1)
	F_blockingPopGenericCommand(m, l0, v3+int32(4), v6+int32(-2), int32(1), v6+v10, v10)
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		return
	} else {
		return
	}
}
func F_bugReportEnd(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v24 int32
	_ = v24
	var v30 int32
	_ = v30
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v49 int32
	_ = v49
	var v54 int32
	_ = v54
	var v58 int32
	_ = v58
	var v72 int32
	_ = v72
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	v5 = m.G0
	v7 = v5 - int32(144)
	m.G0 = v7
	v12 = *(*int32)(unsafe.Add(mBase, _consts[261]))
	if v12 != 0 {
		v13 = int32(_a614)
	} else {
		v13 = int32(_a141)
	}
	*(*int32)(unsafe.Add(mBase, uint32(v7))) = v13
	v18 = m.G0
	v19 = int32(1040)
	v20 = v18 - v19
	m.G0 = v20
	*(*int32)(unsafe.Add(mBase, uint32(v20)+1036)) = v7
	v24 = F_vsnprintf_async_signal_safe(m, v20, int32(1024), int32(_a615), v7)
	mBase = m.M
	F_serverLogRawFromHandler(m, int32(1027), v20)
	mBase = m.M
	m.G0 = v20 + v19
	v30 = *(*int32)(unsafe.Add(mBase, _consts[355]))
	if v30 == int32(0) {
	} else {
		v34 = *(*int32)(unsafe.Add(mBase, _consts[356]))
		if v34 != 0 {
		} else {
			v36 = *(*int32)(unsafe.Add(mBase, _consts[357]))
			if v36 == int32(0) {
			} else {
				v39 = F_unlink(m, v36)
				mBase = m.M
			}
		}
	}
	if l0 != 0 {
		*(*int64)(unsafe.Add(mBase, uint32(v7+int32(8)))) = int64(0)
		v49 = int32(0)
		*(*int32)(unsafe.Add(mBase, uint32(v7)+4)) = v49
		*(*int32)(unsafe.Add(mBase, uint32(v7)+136)) = v49
		v54 = v7 + int32(4)
		if base.Ui32(l1) < base.Ui32(int32(65)) {
			if v54 == int32(0) {
			} else {
				v72 = int32(140)
				v77 = F___memcpy(m, l1*v72+int32(9117184), v54, v72)
				mBase = m.M
			}
		} else {
			v58 = F___errno_location(m)
			mBase = m.M
			*(*int32)(unsafe.Add(mBase, uint32(v58))) = int32(28)
		}
		v80 = F___syscall_getpid(m)
		mBase = m.M
		v81 = F_kill(m, v80, l1)
		mBase = m.M
		v82 = m.ExcPending
		if v82 != 0 {
			return
		} else {
			m.G0 = v7 + int32(144)
			return
		}
	} else {
		v42 = *(*int32)(unsafe.Add(mBase, _consts[358]))
		if v42 != 0 {
			v87 = *(*int32)(unsafe.Add(mBase, _consts[359]))
			v88 = F_fflush(m, v87)
			mBase = m.M
			v89 = m.ExcPending
			if v89 != 0 {
				return
			} else {
				F__Exit(m, int32(1))
				mBase = m.M
				base.Wasm_trap_unreachable()
				for {
				}
			}
		} else {
			F_abort(m)
			mBase = m.M
			base.Wasm_trap_unreachable()
			for {
			}
		}
	}
}
func F_bzpopmaxCommand(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v10 int32
	_ = v10
	var v16 int32
	_ = v16
	v2 = int32(0)
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v10 = int32(-1)
	F_blockingGenericZpopCommand(m, l0, v3+int32(4), v6+int32(-2), int32(1), v6+v10, v10, v2, v2)
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		return
	} else {
		return
	}
}
