package p2

import (
	base "github.com/shibukawa/vkmem/internal/aot/vkaot/base"
	"unsafe"
)

func F___emscripten_environ_constructor(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v19 int32
	_ = v19
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	v12 = m.Wasi_snapshot_preview1.Environ_sizes_get(m, v6+int32(12), v6+int32(8))
	mBase = m.M
	if v12 != 0 {
	} else {
		v13 = int32(0)
		v14 = *(*int32)(unsafe.Add(mBase, uint32(v6)+12))
		v19 = F_emscripten_builtin_malloc(m, v14<<(uint(int32(2))%32)+int32(4))
		mBase = m.M
		*(*int32)(unsafe.Add(mBase, _consts[1021])) = v19
		if v19 == v13 {
		} else {
			v23 = *(*int32)(unsafe.Add(mBase, uint32(v6)+8))
			v24 = F_emscripten_builtin_malloc(m, v23)
			mBase = m.M
			if v24 == int32(0) {
				v39 = int32(0)
				*(*int32)(unsafe.Add(mBase, _consts[1021])) = v39
			} else {
				v27 = int32(0)
				v28 = *(*int32)(unsafe.Add(mBase, _consts[1021]))
				v29 = *(*int32)(unsafe.Add(mBase, uint32(v6)+12))
				*(*int32)(unsafe.Add(mBase, uint32(v28+v29<<(uint(int32(2))%32)))) = v27
				v35 = m.Wasi_snapshot_preview1.Environ_get(m, v28, v24)
				mBase = m.M
				if v35 == v27 {
				} else {
					v39 = int32(0)
					*(*int32)(unsafe.Add(mBase, _consts[1021])) = v39
				}
			}
		}
	}
	m.G0 = v6 + int32(16)
	return
}
func F___emscripten_stdout_close(m *base.Module, l0 int32) int32 {
	return int32(0)
}
func F___emscripten_stdout_seek(m *base.Module, l0 int32, l1 int64, l2 int32) int64 {
	return int64(0)
}
func F__emscripten_yield(m *base.Module, l0 float64) {
	var v3 int32
	_ = v3
	F__emscripten_check_timers(m, l0)
	v3 = m.ExcPending
	if v3 != 0 {
		return
	} else {
		return
	}
}
func F_emscripten_builtin_malloc(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v20 int32
	_ = v20
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
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v70 int32
	_ = v70
	var v76 int32
	_ = v76
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v94 int32
	_ = v94
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v102 int32
	_ = v102
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v114 int32
	_ = v114
	var v116 int32
	_ = v116
	var v118 int32
	_ = v118
	var v122 int32
	_ = v122
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v138 int32
	_ = v138
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v163 int32
	_ = v163
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v183 int32
	_ = v183
	var v186 int32
	_ = v186
	var v191 int32
	_ = v191
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v203 int32
	_ = v203
	var __phi203 int32
	_ = __phi203
	var v206 int32
	_ = v206
	var __phi206 int32
	_ = __phi206
	var v212 int32
	_ = v212
	var v215 int32
	_ = v215
	var v222 int32
	_ = v222
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v235 int32
	_ = v235
	var v238 int32
	_ = v238
	var v246 int32
	_ = v246
	var v248 int32
	_ = v248
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v256 int32
	_ = v256
	var v264 int32
	_ = v264
	var v267 int32
	_ = v267
	var v271 int32
	_ = v271
	var v272 int32
	_ = v272
	var v274 int32
	_ = v274
	var v275 int32
	_ = v275
	var v279 int32
	_ = v279
	var v282 int32
	_ = v282
	var v285 int32
	_ = v285
	var v286 int32
	_ = v286
	var v287 int32
	_ = v287
	var v293 int32
	_ = v293
	var v295 int32
	_ = v295
	var v296 int32
	_ = v296
	var v299 int32
	_ = v299
	var v303 int32
	_ = v303
	var v307 int32
	_ = v307
	var v312 int32
	_ = v312
	var v314 int32
	_ = v314
	var v318 int32
	_ = v318
	var v326 int32
	_ = v326
	var v327 int32
	_ = v327
	var v328 int32
	_ = v328
	var v331 int32
	_ = v331
	var v335 int32
	_ = v335
	var v339 int32
	_ = v339
	var v343 int32
	_ = v343
	var v347 int32
	_ = v347
	var v351 int32
	_ = v351
	var v355 int32
	_ = v355
	var v358 int32
	_ = v358
	var v359 int32
	_ = v359
	var v360 int32
	_ = v360
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
	var v373 int32
	_ = v373
	var v380 int32
	_ = v380
	var v383 int32
	_ = v383
	var v384 int32
	_ = v384
	var v386 int32
	_ = v386
	var v389 int32
	_ = v389
	var v394 int32
	_ = v394
	var v399 int32
	_ = v399
	var v400 int32
	_ = v400
	var v406 int32
	_ = v406
	var __phi406 int32
	_ = __phi406
	var v408 int32
	_ = v408
	var __phi408 int32
	_ = __phi408
	var v415 int32
	_ = v415
	var v418 int32
	_ = v418
	var v424 int32
	_ = v424
	var v434 int32
	_ = v434
	var v437 int32
	_ = v437
	var v438 int32
	_ = v438
	var v441 int32
	_ = v441
	var v453 int32
	_ = v453
	var v454 int32
	_ = v454
	var v458 int32
	_ = v458
	var v461 int32
	_ = v461
	var v462 int32
	_ = v462
	var v463 int32
	_ = v463
	var v470 int32
	_ = v470
	var v472 int32
	_ = v472
	var v473 int32
	_ = v473
	var v477 int32
	_ = v477
	var v478 int32
	_ = v478
	var v488 int32
	_ = v488
	var v489 int32
	_ = v489
	var v493 int32
	_ = v493
	var v494 int32
	_ = v494
	var v515 int32
	_ = v515
	var v516 int32
	_ = v516
	var v518 int32
	_ = v518
	var v519 int32
	_ = v519
	var v521 int32
	_ = v521
	var v522 int32
	_ = v522
	var v524 int32
	_ = v524
	var v526 int32
	_ = v526
	var v530 int32
	_ = v530
	var v531 int32
	_ = v531
	var v537 int32
	_ = v537
	var v540 int32
	_ = v540
	var v541 int32
	_ = v541
	var v545 int32
	_ = v545
	var v557 int32
	_ = v557
	var v559 int32
	_ = v559
	var v562 int32
	_ = v562
	var v576 int32
	_ = v576
	var v579 int32
	_ = v579
	var v580 int32
	_ = v580
	var v582 int32
	_ = v582
	var v592 int32
	_ = v592
	var v594 int32
	_ = v594
	var v595 int32
	_ = v595
	var v599 int32
	_ = v599
	var v600 int32
	_ = v600
	var v605 int32
	_ = v605
	var v608 int32
	_ = v608
	var v609 int32
	_ = v609
	var v610 int32
	_ = v610
	var v611 int32
	_ = v611
	var v614 int32
	_ = v614
	var v616 int32
	_ = v616
	var v632 int32
	_ = v632
	var v633 int32
	_ = v633
	var v637 int32
	_ = v637
	var v638 int32
	_ = v638
	var v656 int32
	_ = v656
	var v658 int32
	_ = v658
	var v674 int32
	_ = v674
	var v676 int32
	_ = v676
	var v682 int32
	_ = v682
	var v688 int32
	_ = v688
	var v693 int32
	_ = v693
	var v698 int32
	_ = v698
	var v700 int32
	_ = v700
	var v701 int32
	_ = v701
	var v704 int32
	_ = v704
	var v708 int32
	_ = v708
	var v709 int32
	_ = v709
	var v713 int32
	_ = v713
	var v725 int32
	_ = v725
	var v726 int32
	_ = v726
	var v729 int32
	_ = v729
	var v730 int32
	_ = v730
	var v731 int32
	_ = v731
	var v737 int32
	_ = v737
	var v747 int32
	_ = v747
	var v752 int32
	_ = v752
	var v765 int32
	_ = v765
	var v769 int32
	_ = v769
	var v775 int32
	_ = v775
	var v778 int32
	_ = v778
	var v780 int32
	_ = v780
	var v784 int32
	_ = v784
	var v785 int32
	_ = v785
	var v788 int32
	_ = v788
	var v798 int32
	_ = v798
	var v802 int32
	_ = v802
	var v807 int32
	_ = v807
	var v811 int32
	_ = v811
	var v812 int32
	_ = v812
	var v816 int32
	_ = v816
	var v817 int32
	_ = v817
	var v818 int32
	_ = v818
	var v828 int32
	_ = v828
	var v834 int32
	_ = v834
	var v840 int32
	_ = v840
	var v852 int32
	_ = v852
	var v854 int32
	_ = v854
	var v855 int32
	_ = v855
	var v862 int32
	_ = v862
	var v874 int32
	_ = v874
	var v876 int32
	_ = v876
	var v877 int32
	_ = v877
	var v880 int32
	_ = v880
	var v881 int32
	_ = v881
	var v883 int32
	_ = v883
	var v886 int32
	_ = v886
	var v887 int32
	_ = v887
	var v888 int32
	_ = v888
	var v891 int32
	_ = v891
	var v901 int32
	_ = v901
	var v909 int32
	_ = v909
	var v913 int32
	_ = v913
	var v918 int32
	_ = v918
	var v919 int64
	_ = v919
	var v922 int64
	_ = v922
	var v937 int32
	_ = v937
	var v957 int32
	_ = v957
	var v961 int32
	_ = v961
	var v969 int32
	_ = v969
	var v971 int32
	_ = v971
	var v973 int32
	_ = v973
	var v977 int32
	_ = v977
	var v982 int32
	_ = v982
	var v983 int32
	_ = v983
	var v994 int32
	_ = v994
	var v997 int32
	_ = v997
	var v1004 int32
	_ = v1004
	var v1009 int32
	_ = v1009
	var v1013 int32
	_ = v1013
	var v1015 int32
	_ = v1015
	var v1029 int32
	_ = v1029
	var v1031 int32
	_ = v1031
	var v1032 int32
	_ = v1032
	var v1040 int32
	_ = v1040
	var v1044 int32
	_ = v1044
	var v1054 int32
	_ = v1054
	var v1055 int32
	_ = v1055
	var v1074 int32
	_ = v1074
	var v1081 int32
	_ = v1081
	var v1086 int32
	_ = v1086
	var v1088 int32
	_ = v1088
	var v1089 int32
	_ = v1089
	var v1110 int32
	_ = v1110
	var v1112 int32
	_ = v1112
	var v1113 int32
	_ = v1113
	var v1117 int32
	_ = v1117
	var v1118 int32
	_ = v1118
	var v1145 int32
	_ = v1145
	var v1155 int32
	_ = v1155
	var v1157 int32
	_ = v1157
	var v1159 int32
	_ = v1159
	var v1167 int32
	_ = v1167
	var v1168 int32
	_ = v1168
	var v1169 int32
	_ = v1169
	var v1171 int32
	_ = v1171
	var v1173 int32
	_ = v1173
	var v1177 int32
	_ = v1177
	var v1178 int32
	_ = v1178
	var v1184 int32
	_ = v1184
	var v1186 int32
	_ = v1186
	var v1190 int32
	_ = v1190
	var v1191 int32
	_ = v1191
	var v1198 int32
	_ = v1198
	var v1204 int32
	_ = v1204
	var v1205 int32
	_ = v1205
	var v1208 int32
	_ = v1208
	var v1210 int32
	_ = v1210
	var v1212 int32
	_ = v1212
	var v1221 int32
	_ = v1221
	var v1223 int32
	_ = v1223
	var v1226 int32
	_ = v1226
	var v1231 int32
	_ = v1231
	var v1236 int32
	_ = v1236
	var v1237 int32
	_ = v1237
	var v1239 int32
	_ = v1239
	var __phi1239 int32
	_ = __phi1239
	var v1245 int32
	_ = v1245
	var __phi1245 int32
	_ = __phi1245
	var v1250 int32
	_ = v1250
	var v1253 int32
	_ = v1253
	var v1259 int32
	_ = v1259
	var v1269 int32
	_ = v1269
	var v1271 int32
	_ = v1271
	var v1274 int32
	_ = v1274
	var v1277 int32
	_ = v1277
	var v1279 int32
	_ = v1279
	var v1284 int32
	_ = v1284
	var v1291 int32
	_ = v1291
	var v1296 int32
	_ = v1296
	var v1312 int32
	_ = v1312
	var v1313 int32
	_ = v1313
	var v1314 int32
	_ = v1314
	var v1315 int32
	_ = v1315
	var v1318 int32
	_ = v1318
	var v1335 int32
	_ = v1335
	var v1337 int32
	_ = v1337
	var v1339 int32
	_ = v1339
	var v1343 int32
	_ = v1343
	var v1348 int32
	_ = v1348
	var v1349 int32
	_ = v1349
	var v1360 int32
	_ = v1360
	var v1363 int32
	_ = v1363
	var v1370 int32
	_ = v1370
	var v1375 int32
	_ = v1375
	var v1379 int32
	_ = v1379
	var v1381 int32
	_ = v1381
	var v1395 int32
	_ = v1395
	var v1397 int32
	_ = v1397
	var v1400 int32
	_ = v1400
	var v1405 int32
	_ = v1405
	var v1408 int32
	_ = v1408
	var v1418 int32
	_ = v1418
	var v1419 int32
	_ = v1419
	var v1436 int32
	_ = v1436
	var v1455 int32
	_ = v1455
	var v1469 int32
	_ = v1469
	var v1471 int32
	_ = v1471
	var v1474 int32
	_ = v1474
	var v1480 int32
	_ = v1480
	var v1482 int32
	_ = v1482
	var v1489 int32
	_ = v1489
	var v1494 int32
	_ = v1494
	var v1501 int32
	_ = v1501
	var v1504 int32
	_ = v1504
	var v1508 int32
	_ = v1508
	var v1509 int32
	_ = v1509
	var v1516 int32
	_ = v1516
	var v1525 int32
	_ = v1525
	var v1527 int32
	_ = v1527
	var v1529 int32
	_ = v1529
	var v1533 int32
	_ = v1533
	var v1538 int32
	_ = v1538
	var v1539 int32
	_ = v1539
	var v1550 int32
	_ = v1550
	var v1553 int32
	_ = v1553
	var v1560 int32
	_ = v1560
	var v1565 int32
	_ = v1565
	var v1569 int32
	_ = v1569
	var v1583 int32
	_ = v1583
	var v1585 int32
	_ = v1585
	var v1586 int32
	_ = v1586
	var v1591 int32
	_ = v1591
	var v1598 int32
	_ = v1598
	var v1608 int32
	_ = v1608
	var v1609 int32
	_ = v1609
	var v1628 int32
	_ = v1628
	var v1649 int32
	_ = v1649
	var v1663 int32
	_ = v1663
	var v1665 int32
	_ = v1665
	var v1668 int32
	_ = v1668
	var v1676 int32
	_ = v1676
	var v1683 int32
	_ = v1683
	var v1688 int32
	_ = v1688
	var v1697 int32
	_ = v1697
	var v1701 int32
	_ = v1701
	var v1702 int32
	_ = v1702
	var v1709 int32
	_ = v1709
	var v1718 int32
	_ = v1718
	var v1720 int32
	_ = v1720
	var v1722 int32
	_ = v1722
	var v1726 int32
	_ = v1726
	var v1731 int32
	_ = v1731
	var v1732 int32
	_ = v1732
	var v1740 int32
	_ = v1740
	var v1750 int32
	_ = v1750
	v13 = m.G0
	v15 = v13 - int32(16)
	m.G0 = v15
	if base.Ui32(int32(244)) < base.Ui32(l0) {
		goto L11
	} else {
		goto L12
	}
L1:
	;
	m.G0 = v15 + int32(16)
	return v1750
L2:
	;
	if v180 == int32(0) {
		goto L306
	} else {
		goto L307
	}
L3:
	;
	if v383 == int32(0) {
		v1501 = v226
		goto L272
	} else {
		goto L273
	}
L4:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v840))) = v693
	v1145 = *(*int32)(unsafe.Add(mBase, uint32(v840)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v840)+4)) = v1145 + v688
	v1155 = int32(-8)
	v1157 = int32(7)
	v1159 = v693 + (v1155-v693)&v1157
	*(*int32)(unsafe.Add(mBase, uint32(v1159)+4)) = v424 | int32(3)
	v1167 = v852 + (v1155-v852)&v1157
	v1168 = v1159 + v424
	v1169 = v1167 - v1168
	v1171 = *(*int32)(unsafe.Add(mBase, _consts[521]))
	if v1167 != v1171 {
		goto L219
	} else {
		goto L220
	}
L5:
	;
	goto L216
L6:
	;
	v1110 = *(*int32)(unsafe.Add(mBase, _consts[522]))
	if base.Ui32(v1110) <= base.Ui32(v424) {
		goto L5
	} else {
		goto L215
	}
L7:
	;
	v834 = *(*int32)(unsafe.Add(mBase, _consts[515]))
	if base.Ui32(v834) <= base.Ui32(v693) {
		goto L174
	} else {
		goto L175
	}
L8:
	;
	v1455 = int32(0)
	goto L3
L9:
	;
	v1649 = int32(0)
	goto L2
L10:
	;
	v434 = *(*int32)(unsafe.Add(mBase, _consts[518]))
	if base.Ui32(v434) < base.Ui32(v424) {
		goto L106
	} else {
		goto L107
	}
L11:
	;
	if base.Ui32(int32(-65)) < base.Ui32(l0) {
		v424 = int32(-1)
		goto L10
	} else {
		goto L55
	}
L12:
	;
	v20 = *(*int32)(unsafe.Add(mBase, _consts[517]))
	v22 = int32(11)
	if base.Ui32(l0) < base.Ui32(v22) {
		goto L14
	} else {
		goto L15
	}
L13:
	;
	v70 = *(*int32)(unsafe.Add(mBase, _consts[518]))
	if base.Ui32(v28) <= base.Ui32(v70) {
		v424 = v28
		goto L10
	} else {
		goto L21
	}
L14:
	;
	v28 = int32(16)
	goto L16
L15:
	;
	v28 = (l0 + v22) & int32(504)
	goto L16
L16:
	;
	v29 = int32(3)
	v30 = int32(base.Ui32(v28) >> (uint(v29) % 32))
	v31 = int32(base.Ui32(v20) >> (uint(v30) % 32))
	if v31&v29 == int32(0) {
		goto L13
	} else {
		goto L17
	}
L17:
	;
	v40 = (v31^int32(-1))&int32(1) + v30
	v42 = v40 << (uint(int32(3)) % 32)
	v44 = v42 + int32(9128464)
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v42)+uint32(_consts[523])))
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v47)+8))
	if v44 != v48 {
		goto L19
	} else {
		goto L20
	}
L18:
	;
	v59 = int32(3)
	v60 = v40 << (uint(v59) % 32)
	*(*int32)(unsafe.Add(mBase, uint32(v47)+4)) = v60 | v59
	v64 = v47 + v60
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v64)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v64)+4)) = v65 | int32(1)
	v1750 = v47 + int32(8)
	goto L1
L19:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v48)+12)) = v44
	*(*int32)(unsafe.Add(mBase, uint32(v42)+uint32(_consts[523]))) = v48
	goto L18
L20:
	;
	*(*int32)(unsafe.Add(mBase, _consts[517])) = v20 & base.I32_rotl(int32(-2), v40)
	goto L18
L21:
	;
	if v31 == int32(0) {
		goto L22
	} else {
		goto L23
	}
L22:
	;
	v142 = int32(0)
	v143 = *(*int32)(unsafe.Add(mBase, _consts[520]))
	if v143 == v142 {
		v424 = v28
		goto L10
	} else {
		goto L32
	}
L23:
	;
	v76 = int32(2) << (uint(v30) % 32)
	v81 = base.I32_ctz(v31 << (uint(v30) % 32) & (v76 | (int32(0) - v76)))
	v83 = v81 << (uint(int32(3)) % 32)
	v85 = v83 + int32(9128464)
	v88 = *(*int32)(unsafe.Add(mBase, uint32(v83)+uint32(_consts[523])))
	v89 = *(*int32)(unsafe.Add(mBase, uint32(v88)+8))
	if v85 != v89 {
		goto L25
	} else {
		goto L26
	}
L24:
	;
	v99 = int32(3)
	*(*int32)(unsafe.Add(mBase, uint32(v88)+4)) = v28 | v99
	v102 = v88 + v28
	v104 = v81 << (uint(v99) % 32)
	v105 = v104 - v28
	*(*int32)(unsafe.Add(mBase, uint32(v102)+4)) = v105 | int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v88+v104))) = v105
	if v70 == int32(0) {
		goto L27
	} else {
		goto L28
	}
L25:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v89)+12)) = v85
	*(*int32)(unsafe.Add(mBase, uint32(v83)+uint32(_consts[523]))) = v89
	v98 = v20
	goto L24
L26:
	;
	v94 = v20 & base.I32_rotl(int32(-2), v81)
	*(*int32)(unsafe.Add(mBase, _consts[517])) = v94
	v98 = v94
	goto L24
L27:
	;
	v138 = int32(0)
	*(*int32)(unsafe.Add(mBase, _consts[516])) = v102
	*(*int32)(unsafe.Add(mBase, _consts[518])) = v105
	v1750 = v88 + int32(8)
	goto L1
L28:
	;
	v114 = v70 & int32(-8)
	v116 = v114 + int32(9128464)
	v118 = *(*int32)(unsafe.Add(mBase, _consts[516]))
	v122 = int32(1) << (uint(int32(base.Ui32(v70)>>(uint(int32(3))%32))) % 32)
	if v98&v122 != 0 {
		goto L30
	} else {
		goto L31
	}
L29:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v114)+uint32(_consts[523]))) = v118
	*(*int32)(unsafe.Add(mBase, uint32(v128)+12)) = v118
	*(*int32)(unsafe.Add(mBase, uint32(v118)+12)) = v116
	*(*int32)(unsafe.Add(mBase, uint32(v118)+8)) = v128
	goto L27
L30:
	;
	v127 = *(*int32)(unsafe.Add(mBase, uint32(v114)+uint32(_consts[523])))
	v128 = v127
	goto L29
L31:
	;
	*(*int32)(unsafe.Add(mBase, _consts[517])) = v98 | v122
	v128 = v116
	goto L29
L32:
	;
	v151 = *(*int32)(unsafe.Add(mBase, uint32(base.I32_ctz(v143)<<(uint(int32(2))%32))+uint32(_consts[519])))
	v152 = *(*int32)(unsafe.Add(mBase, uint32(v151)+4))
	v160 = v152&int32(-8) - v28
	v161 = v151
	v163 = v151
	goto L34
L33:
	;
	v180 = *(*int32)(unsafe.Add(mBase, uint32(v163)+24))
	v181 = *(*int32)(unsafe.Add(mBase, uint32(v163)+12))
	if v181 == v163 {
		goto L45
	} else {
		goto L46
	}
L34:
	;
	v168 = *(*int32)(unsafe.Add(mBase, uint32(v161)+16))
	if v168 != 0 {
		v172 = v168
		goto L36
	} else {
		goto L37
	}
L36:
	;
	v173 = *(*int32)(unsafe.Add(mBase, uint32(v172)+4))
	v176 = v173&int32(-8) - v28
	v177 = base.B2i32(base.Ui32(v176) < base.Ui32(v160))
	if base.Ui32(v176) < base.Ui32(v160) {
		goto L39
	} else {
		goto L40
	}
L37:
	;
	v169 = *(*int32)(unsafe.Add(mBase, uint32(v161)+20))
	if v169 == int32(0) {
		goto L33
	} else {
		goto L38
	}
L38:
	;
	v172 = v169
	goto L36
L39:
	;
	v178 = v176
	goto L41
L40:
	;
	v178 = v160
	goto L41
L41:
	;
	if base.Ui32(v176) < base.Ui32(v160) {
		goto L42
	} else {
		goto L43
	}
L42:
	;
	v179 = v172
	goto L44
L43:
	;
	v179 = v163
	goto L44
L44:
	;
	v160 = v178
	v161 = v172
	v163 = v179
	goto L34
L45:
	;
	v186 = *(*int32)(unsafe.Add(mBase, uint32(v163)+20))
	if v186 == int32(0) {
		goto L48
	} else {
		goto L49
	}
L46:
	;
	v183 = *(*int32)(unsafe.Add(mBase, uint32(v163)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v183)+12)) = v181
	*(*int32)(unsafe.Add(mBase, uint32(v181)+8)) = v183
	v1649 = v181
	goto L2
L47:
	;
	__phi203 = v196
	__phi206 = v197
	v203 = __phi203
	v206 = __phi206
	goto L51
L48:
	;
	v191 = *(*int32)(unsafe.Add(mBase, uint32(v163)+16))
	if v191 == int32(0) {
		goto L9
	} else {
		goto L50
	}
L49:
	;
	v196 = v186
	v197 = v163 + int32(20)
	goto L47
L50:
	;
	v196 = v191
	v197 = v163 + int32(16)
	goto L47
L51:
	;
	v212 = *(*int32)(unsafe.Add(mBase, uint32(v203)+20))
	if v212 != 0 {
		__phi203 = v212
		__phi206 = v203 + int32(20)
		v203 = __phi203
		v206 = __phi206
		goto L51
	} else {
		goto L53
	}
L52:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v206))) = int32(0)
	v1649 = v203
	goto L2
L53:
	;
	v215 = *(*int32)(unsafe.Add(mBase, uint32(v203)+16))
	if v215 != 0 {
		__phi203 = v215
		__phi206 = v203 + int32(16)
		v203 = __phi203
		v206 = __phi206
		goto L51
	} else {
		goto L54
	}
L54:
	;
	goto L52
L55:
	;
	v222 = l0 + int32(11)
	v224 = v222 & int32(-8)
	v225 = int32(0)
	v226 = *(*int32)(unsafe.Add(mBase, _consts[520]))
	if v226 == v225 {
		v424 = v224
		goto L10
	} else {
		goto L56
	}
L56:
	;
	if base.Ui32(int32(16777204)) < base.Ui32(l0) {
		v246 = int32(31)
		goto L57
	} else {
		goto L58
	}
L57:
	;
	v248 = int32(0) - v224
	v253 = *(*int32)(unsafe.Add(mBase, uint32(v246<<(uint(int32(2))%32))+uint32(_consts[519])))
	if v253 != 0 {
		goto L62
	} else {
		goto L63
	}
L58:
	;
	v235 = base.I32_clz(int32(base.Ui32(v222) >> (uint(int32(8)) % 32)))
	v238 = int32(1)
	v246 = int32(base.Ui32(v224)>>(uint(int32(38)-v235)%32))&v238 - v235<<(uint(v238)%32) + int32(62)
	goto L57
L59:
	;
	if v373 == int32(0) {
		v424 = v224
		goto L10
	} else {
		goto L94
	}
L60:
	;
	v343 = v331
	v347 = v335
	v351 = v339
	goto L83
L61:
	;
	if v299|v307 != 0 {
		v327 = v299
		v328 = v307
		goto L79
	} else {
		goto L80
	}
L62:
	;
	v256 = int32(0)
	if v246 == int32(31) {
		goto L64
	} else {
		goto L65
	}
L63:
	;
	v254 = int32(0)
	v299 = v254
	v303 = v248
	v307 = v254
	goto L61
L64:
	;
	v264 = v256
	goto L66
L65:
	;
	v264 = int32(25) - int32(base.Ui32(v246)>>(uint(int32(1))%32))
	goto L66
L66:
	;
	v267 = v256
	v271 = v248
	v272 = v253
	v274 = v224 << (uint(v264) % 32)
	v275 = int32(0)
	goto L67
L67:
	;
	v279 = *(*int32)(unsafe.Add(mBase, uint32(v272)+4))
	v282 = v279&int32(-8) - v224
	if base.Ui32(v271) <= base.Ui32(v282) {
		v285 = v271
		v286 = v275
		goto L69
	} else {
		goto L70
	}
L68:
	;
	v299 = v296
	v303 = v285
	v307 = v286
	goto L61
L69:
	;
	v287 = *(*int32)(unsafe.Add(mBase, uint32(v272)+20))
	v293 = *(*int32)(unsafe.Add(mBase, uint32(v272+int32(base.Ui32(v274)>>(uint(int32(29))%32))&int32(4))+16))
	if v287 == v293 {
		goto L72
	} else {
		goto L73
	}
L70:
	;
	if v282 != 0 {
		v285 = v282
		v286 = v272
		goto L69
	} else {
		goto L71
	}
L71:
	;
	v331 = v272
	v335 = int32(0)
	v339 = v272
	goto L60
L72:
	;
	v295 = v267
	goto L74
L73:
	;
	v295 = v287
	goto L74
L74:
	;
	if v287 != 0 {
		goto L75
	} else {
		goto L76
	}
L75:
	;
	v296 = v295
	goto L77
L76:
	;
	v296 = v267
	goto L77
L77:
	;
	if v293 != 0 {
		v267 = v296
		v271 = v285
		v272 = v293
		v274 = v274 << (uint(int32(1)) % 32)
		v275 = v286
		goto L67
	} else {
		goto L78
	}
L78:
	;
	goto L68
L79:
	;
	if v327 == int32(0) {
		v369 = v303
		v373 = v328
		goto L59
	} else {
		goto L82
	}
L80:
	;
	v312 = int32(0)
	v314 = int32(2) << (uint(v246) % 32)
	v318 = (v314 | (v312 - v314)) & v226
	if v318 == v312 {
		v424 = v224
		goto L10
	} else {
		goto L81
	}
L81:
	;
	v326 = *(*int32)(unsafe.Add(mBase, uint32(base.I32_ctz(v318)<<(uint(int32(2))%32))+uint32(_consts[519])))
	v327 = v326
	v328 = v312
	goto L79
L82:
	;
	v331 = v327
	v335 = v303
	v339 = v328
	goto L60
L83:
	;
	v355 = *(*int32)(unsafe.Add(mBase, uint32(v343)+4))
	v358 = v355&int32(-8) - v224
	v359 = base.B2i32(base.Ui32(v358) < base.Ui32(v347))
	v360 = *(*int32)(unsafe.Add(mBase, uint32(v343)+16))
	if v360 != 0 {
		v362 = v360
		goto L85
	} else {
		goto L86
	}
L84:
	;
	v369 = v363
	v373 = v364
	goto L59
L85:
	;
	if base.Ui32(v358) < base.Ui32(v347) {
		goto L87
	} else {
		goto L88
	}
L86:
	;
	v361 = *(*int32)(unsafe.Add(mBase, uint32(v343)+20))
	v362 = v361
	goto L85
L87:
	;
	v363 = v358
	goto L89
L88:
	;
	v363 = v347
	goto L89
L89:
	;
	if base.Ui32(v358) < base.Ui32(v347) {
		goto L90
	} else {
		goto L91
	}
L90:
	;
	v364 = v343
	goto L92
L91:
	;
	v364 = v351
	goto L92
L92:
	;
	if v362 != 0 {
		v343 = v362
		v347 = v363
		v351 = v364
		goto L83
	} else {
		goto L93
	}
L93:
	;
	goto L84
L94:
	;
	v380 = *(*int32)(unsafe.Add(mBase, _consts[518]))
	if base.Ui32(v380-v224) <= base.Ui32(v369) {
		v424 = v224
		goto L10
	} else {
		goto L95
	}
L95:
	;
	v383 = *(*int32)(unsafe.Add(mBase, uint32(v373)+24))
	v384 = *(*int32)(unsafe.Add(mBase, uint32(v373)+12))
	if v384 == v373 {
		goto L96
	} else {
		goto L97
	}
L96:
	;
	v389 = *(*int32)(unsafe.Add(mBase, uint32(v373)+20))
	if v389 == int32(0) {
		goto L99
	} else {
		goto L100
	}
L97:
	;
	v386 = *(*int32)(unsafe.Add(mBase, uint32(v373)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v386)+12)) = v384
	*(*int32)(unsafe.Add(mBase, uint32(v384)+8)) = v386
	v1455 = v384
	goto L3
L98:
	;
	__phi406 = v399
	__phi408 = v400
	v406 = __phi406
	v408 = __phi408
	goto L102
L99:
	;
	v394 = *(*int32)(unsafe.Add(mBase, uint32(v373)+16))
	if v394 == int32(0) {
		goto L8
	} else {
		goto L101
	}
L100:
	;
	v399 = v389
	v400 = v373 + int32(20)
	goto L98
L101:
	;
	v399 = v394
	v400 = v373 + int32(16)
	goto L98
L102:
	;
	v415 = *(*int32)(unsafe.Add(mBase, uint32(v406)+20))
	if v415 != 0 {
		__phi406 = v415
		__phi408 = v406 + int32(20)
		v406 = __phi406
		v408 = __phi408
		goto L102
	} else {
		goto L104
	}
L103:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v408))) = int32(0)
	v1455 = v406
	goto L3
L104:
	;
	v418 = *(*int32)(unsafe.Add(mBase, uint32(v406)+16))
	if v418 != 0 {
		__phi406 = v418
		__phi408 = v406 + int32(16)
		v406 = __phi406
		v408 = __phi408
		goto L102
	} else {
		goto L105
	}
L105:
	;
	goto L103
L106:
	;
	v470 = *(*int32)(unsafe.Add(mBase, _consts[522]))
	if base.Ui32(v470) <= base.Ui32(v424) {
		goto L111
	} else {
		goto L112
	}
L107:
	;
	v437 = *(*int32)(unsafe.Add(mBase, _consts[516]))
	v438 = v434 - v424
	if base.Ui32(v438) < base.Ui32(int32(16)) {
		goto L109
	} else {
		goto L110
	}
L108:
	;
	v463 = int32(0)
	*(*int32)(unsafe.Add(mBase, _consts[518])) = v461
	*(*int32)(unsafe.Add(mBase, _consts[516])) = v462
	v1750 = v437 + int32(8)
	goto L1
L109:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v437)+4)) = v434 | int32(3)
	v453 = v437 + v434
	v454 = *(*int32)(unsafe.Add(mBase, uint32(v453)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v453)+4)) = v454 | int32(1)
	v458 = int32(0)
	v461 = v458
	v462 = v458
	goto L108
L110:
	;
	v441 = v437 + v424
	*(*int32)(unsafe.Add(mBase, uint32(v441)+4)) = v438 | int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v437+v434))) = v438
	*(*int32)(unsafe.Add(mBase, uint32(v437)+4)) = v424 | int32(3)
	v461 = v438
	v462 = v441
	goto L108
L111:
	;
	v488 = int32(0)
	v489 = *(*int32)(unsafe.Add(mBase, _consts[1089]))
	if v489 == v488 {
		goto L114
	} else {
		goto L115
	}
L112:
	;
	v472 = int32(0)
	v473 = v470 - v424
	*(*int32)(unsafe.Add(mBase, _consts[522])) = v473
	v477 = *(*int32)(unsafe.Add(mBase, _consts[521]))
	v478 = v477 + v424
	*(*int32)(unsafe.Add(mBase, _consts[521])) = v478
	*(*int32)(unsafe.Add(mBase, uint32(v478)+4)) = v473 | int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v477)+4)) = v424 | int32(3)
	v1750 = v477 + int32(8)
	goto L1
L113:
	;
	v516 = int32(0)
	v518 = v424 + int32(47)
	v519 = v515 + v518
	v521 = v516 - v515
	v522 = v519 & v521
	if base.Ui32(v522) <= base.Ui32(v424) {
		v1750 = v516
		goto L1
	} else {
		goto L116
	}
L114:
	;
	v494 = int32(0)
	*(*int64)(unsafe.Add(mBase, _consts[1090])) = int64(-1)
	*(*int64)(unsafe.Add(mBase, _consts[1091])) = int64(17592186048512)
	*(*int32)(unsafe.Add(mBase, _consts[1089])) = (v15+int32(12))&int32(-16) ^ int32(1431655768)
	*(*int32)(unsafe.Add(mBase, _consts[1092])) = v494
	*(*int32)(unsafe.Add(mBase, _consts[1093])) = v494
	v515 = int32(4096)
	goto L113
L115:
	;
	v493 = *(*int32)(unsafe.Add(mBase, _consts[1094]))
	v515 = v493
	goto L113
L116:
	;
	v524 = int32(0)
	v526 = *(*int32)(unsafe.Add(mBase, _consts[1095]))
	if v526 == v524 {
		goto L117
	} else {
		goto L118
	}
L117:
	;
	v537 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1093])))
	if v537&int32(4) != 0 {
		goto L122
	} else {
		goto L123
	}
L118:
	;
	v530 = *(*int32)(unsafe.Add(mBase, _consts[1096]))
	v531 = v530 + v522
	if base.Ui32(v531) <= base.Ui32(v530) {
		v1750 = v524
		goto L1
	} else {
		goto L119
	}
L119:
	;
	if base.Ui32(v526) < base.Ui32(v531) {
		v1750 = v524
		goto L1
	} else {
		goto L120
	}
L120:
	;
	goto L117
L121:
	;
	v698 = int32(0)
	v700 = *(*int32)(unsafe.Add(mBase, _consts[1096]))
	v701 = v700 + v688
	*(*int32)(unsafe.Add(mBase, _consts[1096])) = v701
	v704 = *(*int32)(unsafe.Add(mBase, _consts[1097]))
	if base.Ui32(v701) <= base.Ui32(v704) {
		goto L155
	} else {
		goto L156
	}
L122:
	;
	v674 = F_sbrk(m, v522)
	mBase = m.M
	v676 = F_sbrk(m, int32(0))
	mBase = m.M
	if v674 == int32(-1) {
		goto L5
	} else {
		goto L151
	}
L123:
	;
	v540 = int32(0)
	v541 = *(*int32)(unsafe.Add(mBase, _consts[521]))
	if v541 == v540 {
		goto L128
	} else {
		goto L129
	}
L124:
	;
	v656 = int32(0)
	v658 = *(*int32)(unsafe.Add(mBase, _consts[1093]))
	*(*int32)(unsafe.Add(mBase, _consts[1093])) = v658 | int32(4)
	goto L122
L125:
	;
	if v609 != int32(-1) {
		v688 = v608
		v693 = v609
		goto L121
	} else {
		goto L150
	}
L126:
	;
	if v614 == int32(-1) {
		goto L124
	} else {
		goto L146
	}
L127:
	;
	v608 = (v519 - v470) & v521
	v609 = F_sbrk(m, v608)
	mBase = m.M
	v610 = *(*int32)(unsafe.Add(mBase, uint32(v545)))
	v611 = *(*int32)(unsafe.Add(mBase, uint32(v545)+4))
	if v609 == v610+v611 {
		goto L125
	} else {
		goto L145
	}
L128:
	;
	v576 = F_sbrk(m, int32(0))
	mBase = m.M
	if v576 == int32(-1) {
		goto L124
	} else {
		goto L136
	}
L129:
	;
	v545 = int32(9128872)
	goto L130
L130:
	;
	v557 = *(*int32)(unsafe.Add(mBase, uint32(v545)))
	if base.Ui32(v541) < base.Ui32(v557) {
		goto L132
	} else {
		goto L133
	}
L131:
	;
	goto L128
L132:
	;
	v562 = *(*int32)(unsafe.Add(mBase, uint32(v545)+8))
	if v562 != 0 {
		v545 = v562
		goto L130
	} else {
		goto L135
	}
L133:
	;
	v559 = *(*int32)(unsafe.Add(mBase, uint32(v545)+4))
	if base.Ui32(v541) < base.Ui32(v557+v559) {
		goto L127
	} else {
		goto L134
	}
L134:
	;
	goto L132
L135:
	;
	goto L131
L136:
	;
	v579 = int32(0)
	v580 = *(*int32)(unsafe.Add(mBase, _consts[1091]))
	v582 = v580 + int32(-1)
	if v582&v576 == v579 {
		v592 = v522
		goto L137
	} else {
		goto L138
	}
L137:
	;
	if base.Ui32(v592) <= base.Ui32(v424) {
		goto L124
	} else {
		goto L139
	}
L138:
	;
	v592 = v522 - v576 + (v582+v576)&(int32(0)-v580)
	goto L137
L139:
	;
	v594 = int32(0)
	v595 = *(*int32)(unsafe.Add(mBase, _consts[1095]))
	if v595 == v594 {
		goto L140
	} else {
		goto L141
	}
L140:
	;
	v605 = F_sbrk(m, v592)
	mBase = m.M
	if v605 != v576 {
		v614 = v605
		v616 = v592
		goto L126
	} else {
		goto L144
	}
L141:
	;
	v599 = *(*int32)(unsafe.Add(mBase, _consts[1096]))
	v600 = v599 + v592
	if base.Ui32(v600) <= base.Ui32(v599) {
		goto L124
	} else {
		goto L142
	}
L142:
	;
	if base.Ui32(v595) < base.Ui32(v600) {
		goto L124
	} else {
		goto L143
	}
L143:
	;
	goto L140
L144:
	;
	v688 = v592
	v693 = v576
	goto L121
L145:
	;
	v614 = v609
	v616 = v608
	goto L126
L146:
	;
	if base.Ui32(v616) < base.Ui32(v424+int32(48)) {
		goto L147
	} else {
		goto L148
	}
L147:
	;
	v632 = int32(0)
	v633 = *(*int32)(unsafe.Add(mBase, _consts[1094]))
	v637 = (v518 - v616 + v633) & (v632 - v633)
	v638 = F_sbrk(m, v637)
	mBase = m.M
	if v638 == int32(-1) {
		goto L124
	} else {
		goto L149
	}
L148:
	;
	v688 = v616
	v693 = v614
	goto L121
L149:
	;
	v688 = v637 + v616
	v693 = v614
	goto L121
L150:
	;
	goto L124
L151:
	;
	if v676 == int32(-1) {
		goto L5
	} else {
		goto L152
	}
L152:
	;
	if base.Ui32(v676) <= base.Ui32(v674) {
		goto L5
	} else {
		goto L153
	}
L153:
	;
	v682 = v676 - v674
	if base.Ui32(v682) <= base.Ui32(v424+int32(40)) {
		goto L5
	} else {
		goto L154
	}
L154:
	;
	v688 = v682
	v693 = v674
	goto L121
L155:
	;
	v708 = int32(0)
	v709 = *(*int32)(unsafe.Add(mBase, _consts[521]))
	if v709 == v708 {
		goto L158
	} else {
		goto L159
	}
L156:
	;
	*(*int32)(unsafe.Add(mBase, _consts[1097])) = v701
	goto L155
L157:
	;
	if base.Ui32(v693) <= base.Ui32(v709) {
		goto L7
	} else {
		goto L171
	}
L158:
	;
	v730 = int32(0)
	v731 = *(*int32)(unsafe.Add(mBase, _consts[515]))
	if v731 == v730 {
		goto L165
	} else {
		goto L166
	}
L159:
	;
	v713 = int32(9128872)
	goto L160
L160:
	;
	v725 = *(*int32)(unsafe.Add(mBase, uint32(v713)))
	v726 = *(*int32)(unsafe.Add(mBase, uint32(v713)+4))
	if v693 == v725+v726 {
		goto L157
	} else {
		goto L162
	}
L162:
	;
	v729 = *(*int32)(unsafe.Add(mBase, uint32(v713)+8))
	if v729 != 0 {
		v713 = v729
		goto L160
	} else {
		goto L163
	}
L163:
	;
	goto L7
L164:
	;
	v737 = int32(0)
	*(*int32)(unsafe.Add(mBase, _consts[1098])) = v688
	*(*int32)(unsafe.Add(mBase, _consts[1099])) = v693
	*(*int32)(unsafe.Add(mBase, _consts[524])) = int32(-1)
	v747 = *(*int32)(unsafe.Add(mBase, _consts[1089]))
	*(*int32)(unsafe.Add(mBase, _consts[1100])) = v747
	*(*int32)(unsafe.Add(mBase, _consts[1101])) = v737
	v752 = v737
	goto L168
L165:
	;
	*(*int32)(unsafe.Add(mBase, _consts[515])) = v693
	goto L164
L166:
	;
	if base.Ui32(v731) <= base.Ui32(v693) {
		goto L164
	} else {
		goto L167
	}
L167:
	;
	goto L165
L168:
	;
	v765 = v752 << (uint(int32(3)) % 32)
	v769 = v765 + int32(9128464)
	*(*int32)(unsafe.Add(mBase, uint32(v765)+uint32(_consts[523]))) = v769
	*(*int32)(unsafe.Add(mBase, uint32(v765)+uint32(_consts[1102]))) = v769
	v775 = v752 + int32(1)
	if v775 != int32(32) {
		v752 = v775
		goto L168
	} else {
		goto L170
	}
L169:
	;
	v778 = int32(0)
	v780 = v688 + int32(-40)
	v784 = (int32(-8) - v693) & int32(7)
	v785 = v780 - v784
	*(*int32)(unsafe.Add(mBase, _consts[522])) = v785
	v788 = v693 + v784
	*(*int32)(unsafe.Add(mBase, _consts[521])) = v788
	*(*int32)(unsafe.Add(mBase, uint32(v788)+4)) = v785 | int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v693+v780)+4)) = int32(40)
	v798 = *(*int32)(unsafe.Add(mBase, _consts[1103]))
	*(*int32)(unsafe.Add(mBase, _consts[1104])) = v798
	goto L6
L170:
	;
	goto L169
L171:
	;
	if base.Ui32(v709) < base.Ui32(v725) {
		goto L7
	} else {
		goto L172
	}
L172:
	;
	v802 = *(*int32)(unsafe.Add(mBase, uint32(v713)+12))
	if v802&int32(8) != 0 {
		goto L7
	} else {
		goto L173
	}
L173:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v713)+4)) = v726 + v688
	v807 = int32(0)
	v811 = (int32(-8) - v709) & int32(7)
	v812 = v709 + v811
	*(*int32)(unsafe.Add(mBase, _consts[521])) = v812
	v816 = *(*int32)(unsafe.Add(mBase, _consts[522]))
	v817 = v816 + v688
	v818 = v817 - v811
	*(*int32)(unsafe.Add(mBase, _consts[522])) = v818
	*(*int32)(unsafe.Add(mBase, uint32(v812)+4)) = v818 | int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v709+v817)+4)) = int32(40)
	v828 = *(*int32)(unsafe.Add(mBase, _consts[1103]))
	*(*int32)(unsafe.Add(mBase, _consts[1104])) = v828
	goto L6
L174:
	;
	v840 = int32(9128872)
	goto L178
L175:
	;
	*(*int32)(unsafe.Add(mBase, _consts[515])) = v693
	goto L174
L176:
	;
	v862 = int32(9128872)
	goto L184
L177:
	;
	v855 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v840)+12)))
	if v855&int32(8) == int32(0) {
		goto L4
	} else {
		goto L182
	}
L178:
	;
	v852 = *(*int32)(unsafe.Add(mBase, uint32(v840)))
	if v852 == v693+v688 {
		goto L177
	} else {
		goto L180
	}
L180:
	;
	v854 = *(*int32)(unsafe.Add(mBase, uint32(v840)+8))
	if v854 != 0 {
		v840 = v854
		goto L178
	} else {
		goto L181
	}
L181:
	;
	goto L176
L182:
	;
	goto L176
L183:
	;
	v881 = int32(0)
	v883 = v688 + int32(-40)
	v886 = int32(7)
	v887 = (int32(-8) - v693) & v886
	v888 = v883 - v887
	*(*int32)(unsafe.Add(mBase, _consts[522])) = v888
	v891 = v693 + v887
	*(*int32)(unsafe.Add(mBase, _consts[521])) = v891
	*(*int32)(unsafe.Add(mBase, uint32(v891)+4)) = v888 | int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v693+v883)+4)) = int32(40)
	v901 = *(*int32)(unsafe.Add(mBase, _consts[1103]))
	*(*int32)(unsafe.Add(mBase, _consts[1104])) = v901
	v909 = v877 + (int32(39)-v877)&v886 + int32(-47)
	if base.Ui32(v909) < base.Ui32(v709+int32(16)) {
		goto L189
	} else {
		goto L190
	}
L184:
	;
	v874 = *(*int32)(unsafe.Add(mBase, uint32(v862)))
	if base.Ui32(v709) < base.Ui32(v874) {
		goto L186
	} else {
		goto L187
	}
L186:
	;
	v880 = *(*int32)(unsafe.Add(mBase, uint32(v862)+8))
	v862 = v880
	goto L184
L187:
	;
	v876 = *(*int32)(unsafe.Add(mBase, uint32(v862)+4))
	v877 = v874 + v876
	if base.Ui32(v709) < base.Ui32(v877) {
		goto L183
	} else {
		goto L188
	}
L188:
	;
	goto L186
L189:
	;
	v913 = v709
	goto L191
L190:
	;
	v913 = v909
	goto L191
L191:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v913)+4)) = int32(27)
	v918 = int32(0)
	v919 = *(*int64)(unsafe.Add(mBase, _consts[1105]))
	*(*int64)(unsafe.Add(mBase, uint32(v913+int32(16)))) = v919
	v922 = *(*int64)(unsafe.Add(mBase, _consts[1099]))
	*(*int64)(unsafe.Add(mBase, uint32(v913)+8)) = v922
	*(*int32)(unsafe.Add(mBase, _consts[1105])) = v913 + int32(8)
	*(*int32)(unsafe.Add(mBase, _consts[1098])) = v688
	*(*int32)(unsafe.Add(mBase, _consts[1099])) = v693
	*(*int32)(unsafe.Add(mBase, _consts[1101])) = v918
	v937 = v913 + int32(24)
	goto L192
L192:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v937)+4)) = int32(7)
	if base.Ui32(v937+int32(8)) < base.Ui32(v877) {
		v937 = v937 + int32(4)
		goto L192
	} else {
		goto L194
	}
L193:
	;
	if v913 == v709 {
		goto L6
	} else {
		goto L195
	}
L194:
	;
	goto L193
L195:
	;
	v957 = *(*int32)(unsafe.Add(mBase, uint32(v913)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v913)+4)) = v957 & int32(-2)
	v961 = v913 - v709
	*(*int32)(unsafe.Add(mBase, uint32(v709)+4)) = v961 | int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v913))) = v961
	if base.Ui32(int32(255)) < base.Ui32(v961) {
		goto L197
	} else {
		goto L198
	}
L196:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v709+v1089))) = v1086
	*(*int32)(unsafe.Add(mBase, uint32(v709+v1088))) = v1081
	goto L6
L197:
	;
	if base.Ui32(int32(16777215)) < base.Ui32(v961) {
		v1004 = int32(31)
		goto L202
	} else {
		goto L203
	}
L198:
	;
	v969 = v961 & int32(-8)
	v971 = v969 + int32(9128464)
	v973 = *(*int32)(unsafe.Add(mBase, _consts[517]))
	v977 = int32(1) << (uint(int32(base.Ui32(v961)>>(uint(int32(3))%32))) % 32)
	if v973&v977 != 0 {
		goto L200
	} else {
		goto L201
	}
L199:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v969)+uint32(_consts[523]))) = v709
	*(*int32)(unsafe.Add(mBase, uint32(v983)+12)) = v709
	v1081 = v971
	v1086 = v983
	v1088 = int32(12)
	v1089 = int32(8)
	goto L196
L200:
	;
	v982 = *(*int32)(unsafe.Add(mBase, uint32(v969)+uint32(_consts[523])))
	v983 = v982
	goto L199
L201:
	;
	*(*int32)(unsafe.Add(mBase, _consts[517])) = v973 | v977
	v983 = v971
	goto L199
L202:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v709)+28)) = v1004
	*(*int64)(unsafe.Add(mBase, uint32(v709)+16)) = int64(0)
	v1009 = v1004 << (uint(int32(2)) % 32)
	v1013 = *(*int32)(unsafe.Add(mBase, _consts[520]))
	v1015 = int32(1) << (uint(v1004) % 32)
	if v1013&v1015 != 0 {
		goto L206
	} else {
		goto L207
	}
L203:
	;
	v994 = base.I32_clz(int32(base.Ui32(v961) >> (uint(int32(8)) % 32)))
	v997 = int32(1)
	v1004 = int32(base.Ui32(v961)>>(uint(int32(38)-v994)%32))&v997 - v994<<(uint(v997)%32) + int32(62)
	goto L202
L204:
	;
	v1074 = *(*int32)(unsafe.Add(mBase, uint32(v1040)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v1074)+12)) = v709
	*(*int32)(unsafe.Add(mBase, uint32(v1040)+8)) = v709
	*(*int32)(unsafe.Add(mBase, uint32(v709)+8)) = v1074
	v1081 = int32(0)
	v1086 = v1040
	v1088 = int32(24)
	v1089 = int32(12)
	goto L196
L205:
	;
	v1081 = v709
	v1086 = v709
	v1088 = int32(8)
	v1089 = int32(12)
	goto L196
L206:
	;
	if v1004 == int32(31) {
		goto L208
	} else {
		goto L209
	}
L207:
	;
	*(*int32)(unsafe.Add(mBase, _consts[520])) = v1013 | v1015
	*(*int32)(unsafe.Add(mBase, uint32(v1009)+uint32(_consts[519]))) = v709
	*(*int32)(unsafe.Add(mBase, uint32(v709)+24)) = v1009 + int32(9128728)
	goto L205
L208:
	;
	v1029 = int32(0)
	goto L210
L209:
	;
	v1029 = int32(25) - int32(base.Ui32(v1004)>>(uint(int32(1))%32))
	goto L210
L210:
	;
	v1031 = *(*int32)(unsafe.Add(mBase, uint32(v1009)+uint32(_consts[519])))
	v1032 = v961 << (uint(v1029) % 32)
	v1040 = v1031
	goto L211
L211:
	;
	v1044 = *(*int32)(unsafe.Add(mBase, uint32(v1040)+4))
	if v1044&int32(-8) == v961 {
		goto L204
	} else {
		goto L213
	}
L212:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1054+int32(16)))) = v709
	*(*int32)(unsafe.Add(mBase, uint32(v709)+24)) = v1040
	goto L205
L213:
	;
	v1054 = v1040 + int32(base.Ui32(v1032)>>(uint(int32(29))%32))&int32(4)
	v1055 = *(*int32)(unsafe.Add(mBase, uint32(v1054)+16))
	if v1055 != 0 {
		v1032 = v1032 << (uint(int32(1)) % 32)
		v1040 = v1055
		goto L211
	} else {
		goto L214
	}
L214:
	;
	goto L212
L215:
	;
	v1112 = int32(0)
	v1113 = v1110 - v424
	*(*int32)(unsafe.Add(mBase, _consts[522])) = v1113
	v1117 = *(*int32)(unsafe.Add(mBase, _consts[521]))
	v1118 = v1117 + v424
	*(*int32)(unsafe.Add(mBase, _consts[521])) = v1118
	*(*int32)(unsafe.Add(mBase, uint32(v1118)+4)) = v1113 | int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v1117)+4)) = v424 | int32(3)
	v1750 = v1117 + int32(8)
	goto L1
L216:
	;
	*(*int32)(unsafe.Add(mBase, _consts[5])) = int32(48)
	v1750 = int32(0)
	goto L1
L217:
	;
	v1750 = v1159 + int32(8)
	goto L1
L218:
	;
	goto L217
L219:
	;
	v1184 = *(*int32)(unsafe.Add(mBase, _consts[516]))
	if v1167 != v1184 {
		goto L221
	} else {
		goto L222
	}
L220:
	;
	v1173 = int32(0)
	*(*int32)(unsafe.Add(mBase, _consts[521])) = v1168
	v1177 = *(*int32)(unsafe.Add(mBase, _consts[522]))
	v1178 = v1177 + v1169
	*(*int32)(unsafe.Add(mBase, _consts[522])) = v1178
	*(*int32)(unsafe.Add(mBase, uint32(v1168)+4)) = v1178 | int32(1)
	goto L218
L221:
	;
	v1198 = *(*int32)(unsafe.Add(mBase, uint32(v1167)+4))
	if v1198&int32(3) != int32(1) {
		v1314 = v1169
		v1315 = v1198
		v1318 = v1167
		goto L223
	} else {
		goto L224
	}
L222:
	;
	v1186 = int32(0)
	*(*int32)(unsafe.Add(mBase, _consts[516])) = v1168
	v1190 = *(*int32)(unsafe.Add(mBase, _consts[518]))
	v1191 = v1190 + v1169
	*(*int32)(unsafe.Add(mBase, _consts[518])) = v1191
	*(*int32)(unsafe.Add(mBase, uint32(v1168)+4)) = v1191 | int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v1168+v1191))) = v1191
	goto L218
L223:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1318)+4)) = v1315 & int32(-2)
	*(*int32)(unsafe.Add(mBase, uint32(v1168)+4)) = v1314 | int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v1168+v1314))) = v1314
	if base.Ui32(int32(255)) < base.Ui32(v1314) {
		goto L254
	} else {
		goto L255
	}
L224:
	;
	v1204 = v1198 & int32(-8)
	v1205 = *(*int32)(unsafe.Add(mBase, uint32(v1167)+12))
	if base.Ui32(int32(255)) < base.Ui32(v1198) {
		goto L226
	} else {
		goto L227
	}
L225:
	;
	v1312 = v1167 + v1204
	v1313 = *(*int32)(unsafe.Add(mBase, uint32(v1312)+4))
	v1314 = v1204 + v1169
	v1315 = v1313
	v1318 = v1312
	goto L223
L226:
	;
	v1221 = *(*int32)(unsafe.Add(mBase, uint32(v1167)+24))
	if v1205 == v1167 {
		goto L231
	} else {
		goto L232
	}
L227:
	;
	v1208 = *(*int32)(unsafe.Add(mBase, uint32(v1167)+8))
	if v1205 != v1208 {
		goto L228
	} else {
		goto L229
	}
L228:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1208)+12)) = v1205
	*(*int32)(unsafe.Add(mBase, uint32(v1205)+8)) = v1208
	goto L225
L229:
	;
	v1210 = int32(0)
	v1212 = *(*int32)(unsafe.Add(mBase, _consts[517]))
	*(*int32)(unsafe.Add(mBase, _consts[517])) = v1212 & base.I32_rotl(int32(-2), int32(base.Ui32(v1198)>>(uint(int32(3))%32)))
	goto L225
L230:
	;
	if v1221 == int32(0) {
		goto L225
	} else {
		goto L242
	}
L231:
	;
	v1226 = *(*int32)(unsafe.Add(mBase, uint32(v1167)+20))
	if v1226 == int32(0) {
		goto L235
	} else {
		goto L236
	}
L232:
	;
	v1223 = *(*int32)(unsafe.Add(mBase, uint32(v1167)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v1223)+12)) = v1205
	*(*int32)(unsafe.Add(mBase, uint32(v1205)+8)) = v1223
	v1259 = v1205
	goto L230
L233:
	;
	v1259 = int32(0)
	goto L230
L234:
	;
	__phi1239 = v1236
	__phi1245 = v1237
	v1239 = __phi1239
	v1245 = __phi1245
	goto L238
L235:
	;
	v1231 = *(*int32)(unsafe.Add(mBase, uint32(v1167)+16))
	if v1231 == int32(0) {
		goto L233
	} else {
		goto L237
	}
L236:
	;
	v1236 = v1226
	v1237 = v1167 + int32(20)
	goto L234
L237:
	;
	v1236 = v1231
	v1237 = v1167 + int32(16)
	goto L234
L238:
	;
	v1250 = *(*int32)(unsafe.Add(mBase, uint32(v1239)+20))
	if v1250 != 0 {
		__phi1239 = v1250
		__phi1245 = v1239 + int32(20)
		v1239 = __phi1239
		v1245 = __phi1245
		goto L238
	} else {
		goto L240
	}
L239:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1245))) = int32(0)
	v1259 = v1239
	goto L230
L240:
	;
	v1253 = *(*int32)(unsafe.Add(mBase, uint32(v1239)+16))
	if v1253 != 0 {
		__phi1239 = v1253
		__phi1245 = v1239 + int32(16)
		v1239 = __phi1239
		v1245 = __phi1245
		goto L238
	} else {
		goto L241
	}
L241:
	;
	goto L239
L242:
	;
	v1269 = *(*int32)(unsafe.Add(mBase, uint32(v1167)+28))
	v1271 = v1269 << (uint(int32(2)) % 32)
	v1274 = *(*int32)(unsafe.Add(mBase, uint32(v1271)+uint32(_consts[519])))
	if v1167 != v1274 {
		goto L244
	} else {
		goto L245
	}
L243:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1259)+24)) = v1221
	v1291 = *(*int32)(unsafe.Add(mBase, uint32(v1167)+16))
	if v1291 == int32(0) {
		goto L251
	} else {
		goto L252
	}
L244:
	;
	v1284 = *(*int32)(unsafe.Add(mBase, uint32(v1221)+16))
	if v1284 != v1167 {
		goto L248
	} else {
		goto L249
	}
L245:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1271)+uint32(_consts[519]))) = v1259
	if v1259 != 0 {
		goto L243
	} else {
		goto L246
	}
L246:
	;
	v1277 = int32(0)
	v1279 = *(*int32)(unsafe.Add(mBase, _consts[520]))
	*(*int32)(unsafe.Add(mBase, _consts[520])) = v1279 & base.I32_rotl(int32(-2), v1269)
	goto L225
L247:
	;
	if v1259 == int32(0) {
		goto L225
	} else {
		goto L250
	}
L248:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1221)+20)) = v1259
	goto L247
L249:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1221)+16)) = v1259
	goto L247
L250:
	;
	goto L243
L251:
	;
	v1296 = *(*int32)(unsafe.Add(mBase, uint32(v1167)+20))
	if v1296 == int32(0) {
		goto L225
	} else {
		goto L253
	}
L252:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1259)+16)) = v1291
	*(*int32)(unsafe.Add(mBase, uint32(v1291)+24)) = v1259
	goto L251
L253:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1259)+20)) = v1296
	*(*int32)(unsafe.Add(mBase, uint32(v1296)+24)) = v1259
	goto L225
L254:
	;
	if base.Ui32(int32(16777215)) < base.Ui32(v1314) {
		v1370 = int32(31)
		goto L259
	} else {
		goto L260
	}
L255:
	;
	v1335 = v1314 & int32(-8)
	v1337 = v1335 + int32(9128464)
	v1339 = *(*int32)(unsafe.Add(mBase, _consts[517]))
	v1343 = int32(1) << (uint(int32(base.Ui32(v1314)>>(uint(int32(3))%32))) % 32)
	if v1339&v1343 != 0 {
		goto L257
	} else {
		goto L258
	}
L256:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1335)+uint32(_consts[523]))) = v1168
	*(*int32)(unsafe.Add(mBase, uint32(v1349)+12)) = v1168
	*(*int32)(unsafe.Add(mBase, uint32(v1168)+12)) = v1337
	*(*int32)(unsafe.Add(mBase, uint32(v1168)+8)) = v1349
	goto L218
L257:
	;
	v1348 = *(*int32)(unsafe.Add(mBase, uint32(v1335)+uint32(_consts[523])))
	v1349 = v1348
	goto L256
L258:
	;
	*(*int32)(unsafe.Add(mBase, _consts[517])) = v1339 | v1343
	v1349 = v1337
	goto L256
L259:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1168)+28)) = v1370
	*(*int64)(unsafe.Add(mBase, uint32(v1168)+16)) = int64(0)
	v1375 = v1370 << (uint(int32(2)) % 32)
	v1379 = *(*int32)(unsafe.Add(mBase, _consts[520]))
	v1381 = int32(1) << (uint(v1370) % 32)
	if v1379&v1381 != 0 {
		goto L263
	} else {
		goto L264
	}
L260:
	;
	v1360 = base.I32_clz(int32(base.Ui32(v1314) >> (uint(int32(8)) % 32)))
	v1363 = int32(1)
	v1370 = int32(base.Ui32(v1314)>>(uint(int32(38)-v1360)%32))&v1363 - v1360<<(uint(v1363)%32) + int32(62)
	goto L259
L261:
	;
	v1436 = *(*int32)(unsafe.Add(mBase, uint32(v1405)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v1436)+12)) = v1168
	*(*int32)(unsafe.Add(mBase, uint32(v1405)+8)) = v1168
	*(*int32)(unsafe.Add(mBase, uint32(v1168)+24)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v1168)+12)) = v1405
	*(*int32)(unsafe.Add(mBase, uint32(v1168)+8)) = v1436
	goto L218
L262:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1168)+12)) = v1168
	*(*int32)(unsafe.Add(mBase, uint32(v1168)+8)) = v1168
	goto L218
L263:
	;
	if v1370 == int32(31) {
		goto L265
	} else {
		goto L266
	}
L264:
	;
	*(*int32)(unsafe.Add(mBase, _consts[520])) = v1379 | v1381
	*(*int32)(unsafe.Add(mBase, uint32(v1375)+uint32(_consts[519]))) = v1168
	*(*int32)(unsafe.Add(mBase, uint32(v1168)+24)) = v1375 + int32(9128728)
	goto L262
L265:
	;
	v1395 = int32(0)
	goto L267
L266:
	;
	v1395 = int32(25) - int32(base.Ui32(v1370)>>(uint(int32(1))%32))
	goto L267
L267:
	;
	v1397 = *(*int32)(unsafe.Add(mBase, uint32(v1375)+uint32(_consts[519])))
	v1400 = v1314 << (uint(v1395) % 32)
	v1405 = v1397
	goto L268
L268:
	;
	v1408 = *(*int32)(unsafe.Add(mBase, uint32(v1405)+4))
	if v1408&int32(-8) == v1314 {
		goto L261
	} else {
		goto L270
	}
L269:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1418+int32(16)))) = v1168
	*(*int32)(unsafe.Add(mBase, uint32(v1168)+24)) = v1405
	goto L262
L270:
	;
	v1418 = v1405 + int32(base.Ui32(v1400)>>(uint(int32(29))%32))&int32(4)
	v1419 = *(*int32)(unsafe.Add(mBase, uint32(v1418)+16))
	if v1419 != 0 {
		v1400 = v1400 << (uint(int32(1)) % 32)
		v1405 = v1419
		goto L268
	} else {
		goto L271
	}
L271:
	;
	goto L269
L272:
	;
	if base.Ui32(int32(15)) < base.Ui32(v369) {
		goto L286
	} else {
		goto L287
	}
L273:
	;
	v1469 = *(*int32)(unsafe.Add(mBase, uint32(v373)+28))
	v1471 = v1469 << (uint(int32(2)) % 32)
	v1474 = *(*int32)(unsafe.Add(mBase, uint32(v1471)+uint32(_consts[519])))
	if v373 != v1474 {
		goto L275
	} else {
		goto L276
	}
L274:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1455)+24)) = v383
	v1489 = *(*int32)(unsafe.Add(mBase, uint32(v373)+16))
	if v1489 == int32(0) {
		goto L282
	} else {
		goto L283
	}
L275:
	;
	v1482 = *(*int32)(unsafe.Add(mBase, uint32(v383)+16))
	if v1482 != v373 {
		goto L279
	} else {
		goto L280
	}
L276:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1471)+uint32(_consts[519]))) = v1455
	if v1455 != 0 {
		goto L274
	} else {
		goto L277
	}
L277:
	;
	v1480 = v226 & base.I32_rotl(int32(-2), v1469)
	*(*int32)(unsafe.Add(mBase, _consts[520])) = v1480
	v1501 = v1480
	goto L272
L278:
	;
	if v1455 == int32(0) {
		v1501 = v226
		goto L272
	} else {
		goto L281
	}
L279:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v383)+20)) = v1455
	goto L278
L280:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v383)+16)) = v1455
	goto L278
L281:
	;
	goto L274
L282:
	;
	v1494 = *(*int32)(unsafe.Add(mBase, uint32(v373)+20))
	if v1494 == int32(0) {
		v1501 = v226
		goto L272
	} else {
		goto L284
	}
L283:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1455)+16)) = v1489
	*(*int32)(unsafe.Add(mBase, uint32(v1489)+24)) = v1455
	goto L282
L284:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1455)+20)) = v1494
	*(*int32)(unsafe.Add(mBase, uint32(v1494)+24)) = v1455
	v1501 = v226
	goto L272
L285:
	;
	v1750 = v373 + int32(8)
	goto L1
L286:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v373)+4)) = v224 | int32(3)
	v1516 = v373 + v224
	*(*int32)(unsafe.Add(mBase, uint32(v1516)+4)) = v369 | int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v1516+v369))) = v369
	if base.Ui32(int32(255)) < base.Ui32(v369) {
		goto L288
	} else {
		goto L289
	}
L287:
	;
	v1504 = v369 + v224
	*(*int32)(unsafe.Add(mBase, uint32(v373)+4)) = v1504 | int32(3)
	v1508 = v373 + v1504
	v1509 = *(*int32)(unsafe.Add(mBase, uint32(v1508)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v1508)+4)) = v1509 | int32(1)
	goto L285
L288:
	;
	if base.Ui32(int32(16777215)) < base.Ui32(v369) {
		v1560 = int32(31)
		goto L293
	} else {
		goto L294
	}
L289:
	;
	v1525 = v369 & int32(-8)
	v1527 = v1525 + int32(9128464)
	v1529 = *(*int32)(unsafe.Add(mBase, _consts[517]))
	v1533 = int32(1) << (uint(int32(base.Ui32(v369)>>(uint(int32(3))%32))) % 32)
	if v1529&v1533 != 0 {
		goto L291
	} else {
		goto L292
	}
L290:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1525)+uint32(_consts[523]))) = v1516
	*(*int32)(unsafe.Add(mBase, uint32(v1539)+12)) = v1516
	*(*int32)(unsafe.Add(mBase, uint32(v1516)+12)) = v1527
	*(*int32)(unsafe.Add(mBase, uint32(v1516)+8)) = v1539
	goto L285
L291:
	;
	v1538 = *(*int32)(unsafe.Add(mBase, uint32(v1525)+uint32(_consts[523])))
	v1539 = v1538
	goto L290
L292:
	;
	*(*int32)(unsafe.Add(mBase, _consts[517])) = v1529 | v1533
	v1539 = v1527
	goto L290
L293:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1516)+28)) = v1560
	*(*int64)(unsafe.Add(mBase, uint32(v1516)+16)) = int64(0)
	v1565 = v1560 << (uint(int32(2)) % 32)
	v1569 = int32(1) << (uint(v1560) % 32)
	if v1501&v1569 != 0 {
		goto L297
	} else {
		goto L298
	}
L294:
	;
	v1550 = base.I32_clz(int32(base.Ui32(v369) >> (uint(int32(8)) % 32)))
	v1553 = int32(1)
	v1560 = int32(base.Ui32(v369)>>(uint(int32(38)-v1550)%32))&v1553 - v1550<<(uint(v1553)%32) + int32(62)
	goto L293
L295:
	;
	v1628 = *(*int32)(unsafe.Add(mBase, uint32(v1591)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v1628)+12)) = v1516
	*(*int32)(unsafe.Add(mBase, uint32(v1591)+8)) = v1516
	*(*int32)(unsafe.Add(mBase, uint32(v1516)+24)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v1516)+12)) = v1591
	*(*int32)(unsafe.Add(mBase, uint32(v1516)+8)) = v1628
	goto L285
L296:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1516)+12)) = v1516
	*(*int32)(unsafe.Add(mBase, uint32(v1516)+8)) = v1516
	goto L285
L297:
	;
	if v1560 == int32(31) {
		goto L299
	} else {
		goto L300
	}
L298:
	;
	*(*int32)(unsafe.Add(mBase, _consts[520])) = v1501 | v1569
	*(*int32)(unsafe.Add(mBase, uint32(v1565)+uint32(_consts[519]))) = v1516
	*(*int32)(unsafe.Add(mBase, uint32(v1516)+24)) = v1565 + int32(9128728)
	goto L296
L299:
	;
	v1583 = int32(0)
	goto L301
L300:
	;
	v1583 = int32(25) - int32(base.Ui32(v1560)>>(uint(int32(1))%32))
	goto L301
L301:
	;
	v1585 = *(*int32)(unsafe.Add(mBase, uint32(v1565)+uint32(_consts[519])))
	v1586 = v369 << (uint(v1583) % 32)
	v1591 = v1585
	goto L302
L302:
	;
	v1598 = *(*int32)(unsafe.Add(mBase, uint32(v1591)+4))
	if v1598&int32(-8) == v369 {
		goto L295
	} else {
		goto L304
	}
L303:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1608+int32(16)))) = v1516
	*(*int32)(unsafe.Add(mBase, uint32(v1516)+24)) = v1591
	goto L296
L304:
	;
	v1608 = v1591 + int32(base.Ui32(v1586)>>(uint(int32(29))%32))&int32(4)
	v1609 = *(*int32)(unsafe.Add(mBase, uint32(v1608)+16))
	if v1609 != 0 {
		v1586 = v1586 << (uint(int32(1)) % 32)
		v1591 = v1609
		goto L302
	} else {
		goto L305
	}
L305:
	;
	goto L303
L306:
	;
	if base.Ui32(int32(15)) < base.Ui32(v160) {
		goto L320
	} else {
		goto L321
	}
L307:
	;
	v1663 = *(*int32)(unsafe.Add(mBase, uint32(v163)+28))
	v1665 = v1663 << (uint(int32(2)) % 32)
	v1668 = *(*int32)(unsafe.Add(mBase, uint32(v1665)+uint32(_consts[519])))
	if v163 != v1668 {
		goto L309
	} else {
		goto L310
	}
L308:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1649)+24)) = v180
	v1683 = *(*int32)(unsafe.Add(mBase, uint32(v163)+16))
	if v1683 == int32(0) {
		goto L316
	} else {
		goto L317
	}
L309:
	;
	v1676 = *(*int32)(unsafe.Add(mBase, uint32(v180)+16))
	if v1676 != v163 {
		goto L313
	} else {
		goto L314
	}
L310:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1665)+uint32(_consts[519]))) = v1649
	if v1649 != 0 {
		goto L308
	} else {
		goto L311
	}
L311:
	;
	*(*int32)(unsafe.Add(mBase, _consts[520])) = v143 & base.I32_rotl(int32(-2), v1663)
	goto L306
L312:
	;
	if v1649 == int32(0) {
		goto L306
	} else {
		goto L315
	}
L313:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v180)+20)) = v1649
	goto L312
L314:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v180)+16)) = v1649
	goto L312
L315:
	;
	goto L308
L316:
	;
	v1688 = *(*int32)(unsafe.Add(mBase, uint32(v163)+20))
	if v1688 == int32(0) {
		goto L306
	} else {
		goto L318
	}
L317:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1649)+16)) = v1683
	*(*int32)(unsafe.Add(mBase, uint32(v1683)+24)) = v1649
	goto L316
L318:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1649)+20)) = v1688
	*(*int32)(unsafe.Add(mBase, uint32(v1688)+24)) = v1649
	goto L306
L319:
	;
	v1750 = v163 + int32(8)
	goto L1
L320:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v163)+4)) = v28 | int32(3)
	v1709 = v163 + v28
	*(*int32)(unsafe.Add(mBase, uint32(v1709)+4)) = v160 | int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v1709+v160))) = v160
	if v70 == int32(0) {
		goto L322
	} else {
		goto L323
	}
L321:
	;
	v1697 = v160 + v28
	*(*int32)(unsafe.Add(mBase, uint32(v163)+4)) = v1697 | int32(3)
	v1701 = v163 + v1697
	v1702 = *(*int32)(unsafe.Add(mBase, uint32(v1701)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v1701)+4)) = v1702 | int32(1)
	goto L319
L322:
	;
	v1740 = int32(0)
	*(*int32)(unsafe.Add(mBase, _consts[516])) = v1709
	*(*int32)(unsafe.Add(mBase, _consts[518])) = v160
	goto L319
L323:
	;
	v1718 = v70 & int32(-8)
	v1720 = v1718 + int32(9128464)
	v1722 = *(*int32)(unsafe.Add(mBase, _consts[516]))
	v1726 = int32(1) << (uint(int32(base.Ui32(v70)>>(uint(int32(3))%32))) % 32)
	if v1726&v20 != 0 {
		goto L325
	} else {
		goto L326
	}
L324:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1718)+uint32(_consts[523]))) = v1722
	*(*int32)(unsafe.Add(mBase, uint32(v1732)+12)) = v1722
	*(*int32)(unsafe.Add(mBase, uint32(v1722)+12)) = v1720
	*(*int32)(unsafe.Add(mBase, uint32(v1722)+8)) = v1732
	goto L322
L325:
	;
	v1731 = *(*int32)(unsafe.Add(mBase, uint32(v1718)+uint32(_consts[523])))
	v1732 = v1731
	goto L324
L326:
	;
	*(*int32)(unsafe.Add(mBase, _consts[517])) = v1726 | v20
	v1732 = v1720
	goto L324
}
func F_emscripten_get_heap_size(m *base.Module) int32 {
	return base.MemorySize(m) << (uint(int32(16)) % 32)
}
func F_emscripten_num_logical_cores(m *base.Module) int32 {
	return int32(1)
}
func F_emscripten_stack_get_current(m *base.Module) int32 {
	var v1 int32
	_ = v1
	v1 = m.G0
	return v1
}
func F_emscripten_stack_get_free(m *base.Module) int32 {
	var v1 int32
	_ = v1
	var v2 int32
	_ = v2
	v1 = m.G0
	v2 = m.G1
	return v1 - v2
}
func F_emscripten_stack_init(m *base.Module) {
	m.G2 = int32(_a0)
	m.G1 = int32(0)
	return
}
