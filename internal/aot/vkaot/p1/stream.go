package p1

import (
	base "github.com/shibukawa/vkmem/internal/aot/vkaot/base"
	"unsafe"
)

func F_streamAppendItem(m *base.Module, l0 int32, l1 int32, l2 int64, l3 int32, l4 int32, l5 int32) int32 {
	mBase := m.M
	_ = mBase
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v30 int64
	_ = v30
	var v31 int64
	_ = v31
	var v32 int64
	_ = v32
	var v33 int64
	_ = v33
	var v34 int64
	_ = v34
	var v36 int64
	_ = v36
	var v45 int64
	_ = v45
	var v46 int64
	_ = v46
	var v48 int64
	_ = v48
	var v58 int64
	_ = v58
	var v61 int64
	_ = v61
	var v66 int64
	_ = v66
	var v67 int64
	_ = v67
	var v68 int64
	_ = v68
	var v69 int64
	_ = v69
	var v71 int32
	_ = v71
	var v73 int64
	_ = v73
	var v74 int64
	_ = v74
	var v75 int64
	_ = v75
	var v77 int32
	_ = v77
	var v79 int64
	_ = v79
	var v82 int64
	_ = v82
	var v83 int64
	_ = v83
	var v85 int32
	_ = v85
	var v89 int64
	_ = v89
	var v90 int64
	_ = v90
	var v94 int64
	_ = v94
	var v102 int32
	_ = v102
	var v104 int64
	_ = v104
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v127 int32
	_ = v127
	var v134 int32
	_ = v134
	var v137 int32
	_ = v137
	var v140 int32
	_ = v140
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v147 int64
	_ = v147
	var v160 int32
	_ = v160
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v184 int64
	_ = v184
	var v199 int32
	_ = v199
	var v201 int32
	_ = v201
	var v204 int32
	_ = v204
	var v207 int32
	_ = v207
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v221 int32
	_ = v221
	var v227 int32
	_ = v227
	var v232 int32
	_ = v232
	var v235 int64
	_ = v235
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v243 int64
	_ = v243
	var v244 int32
	_ = v244
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v259 int32
	_ = v259
	var v260 int32
	_ = v260
	var v264 int32
	_ = v264
	var v280 int32
	_ = v280
	var v282 int32
	_ = v282
	var v283 int32
	_ = v283
	var v284 int32
	_ = v284
	var v295 int64
	_ = v295
	var v301 int32
	_ = v301
	var v303 int64
	_ = v303
	var v305 int32
	_ = v305
	var v307 int32
	_ = v307
	var v317 int64
	_ = v317
	var v322 int64
	_ = v322
	var v326 int32
	_ = v326
	var v328 int64
	_ = v328
	var v330 int32
	_ = v330
	var v338 int64
	_ = v338
	var v362 int64
	_ = v362
	var v381 int32
	_ = v381
	var v390 int64
	_ = v390
	var v391 int64
	_ = v391
	var v392 int32
	_ = v392
	var v393 int32
	_ = v393
	var v395 int32
	_ = v395
	var v396 int32
	_ = v396
	var v397 int64
	_ = v397
	var v398 int32
	_ = v398
	var v400 int32
	_ = v400
	var v401 int32
	_ = v401
	var v413 int32
	_ = v413
	var v414 int32
	_ = v414
	var v418 int32
	_ = v418
	var v434 int32
	_ = v434
	var v436 int32
	_ = v436
	var v437 int32
	_ = v437
	var v438 int32
	_ = v438
	var v449 int64
	_ = v449
	var v455 int32
	_ = v455
	var v457 int64
	_ = v457
	var v459 int32
	_ = v459
	var v461 int32
	_ = v461
	var v471 int64
	_ = v471
	var v476 int64
	_ = v476
	var v480 int32
	_ = v480
	var v482 int64
	_ = v482
	var v484 int32
	_ = v484
	var v492 int64
	_ = v492
	var v516 int64
	_ = v516
	var v535 int32
	_ = v535
	var v544 int64
	_ = v544
	var v545 int64
	_ = v545
	var v548 int64
	_ = v548
	var v554 int32
	_ = v554
	var v555 int32
	_ = v555
	var v556 int32
	_ = v556
	var v558 int32
	_ = v558
	var v559 int32
	_ = v559
	var v560 int32
	_ = v560
	var v562 int32
	_ = v562
	var v563 int32
	_ = v563
	var v569 int64
	_ = v569
	var v571 int64
	_ = v571
	var v573 int64
	_ = v573
	var v576 int64
	_ = v576
	var v578 int64
	_ = v578
	var v580 int64
	_ = v580
	var v582 int64
	_ = v582
	var v643 int32
	_ = v643
	var v649 int32
	_ = v649
	var v650 int32
	_ = v650
	var v651 int32
	_ = v651
	var v653 int32
	_ = v653
	var v654 int32
	_ = v654
	var v656 int32
	_ = v656
	var v657 int32
	_ = v657
	var v658 int32
	_ = v658
	var v659 int32
	_ = v659
	var v669 int64
	_ = v669
	var v674 int32
	_ = v674
	var v688 int32
	_ = v688
	var v689 int32
	_ = v689
	var v692 int32
	_ = v692
	var v699 int32
	_ = v699
	var v702 int32
	_ = v702
	var v705 int32
	_ = v705
	var v708 int32
	_ = v708
	var v709 int32
	_ = v709
	var v710 int32
	_ = v710
	var v711 int32
	_ = v711
	var v713 int64
	_ = v713
	var v727 int32
	_ = v727
	var v737 int32
	_ = v737
	var v738 int32
	_ = v738
	var v739 int32
	_ = v739
	var v744 int32
	_ = v744
	var v745 int32
	_ = v745
	var v751 int32
	_ = v751
	var v754 int32
	_ = v754
	var v755 int64
	_ = v755
	var v759 int64
	_ = v759
	var v761 int32
	_ = v761
	var v762 int32
	_ = v762
	var v765 int32
	_ = v765
	var v766 int32
	_ = v766
	var v767 int64
	_ = v767
	var v768 int32
	_ = v768
	var v770 int32
	_ = v770
	var v771 int32
	_ = v771
	var v783 int32
	_ = v783
	var v784 int32
	_ = v784
	var v788 int32
	_ = v788
	var v804 int32
	_ = v804
	var v806 int32
	_ = v806
	var v807 int32
	_ = v807
	var v808 int32
	_ = v808
	var v819 int64
	_ = v819
	var v825 int32
	_ = v825
	var v827 int64
	_ = v827
	var v829 int32
	_ = v829
	var v831 int32
	_ = v831
	var v841 int64
	_ = v841
	var v846 int64
	_ = v846
	var v850 int32
	_ = v850
	var v852 int64
	_ = v852
	var v854 int32
	_ = v854
	var v862 int64
	_ = v862
	var v886 int64
	_ = v886
	var v905 int32
	_ = v905
	var v914 int64
	_ = v914
	var v915 int64
	_ = v915
	var v920 int32
	_ = v920
	var v921 int32
	_ = v921
	var v922 int32
	_ = v922
	var v923 int32
	_ = v923
	var v924 int32
	_ = v924
	var v926 int32
	_ = v926
	var v927 int32
	_ = v927
	var v930 int32
	_ = v930
	var v931 int32
	_ = v931
	var v932 int64
	_ = v932
	var v933 int32
	_ = v933
	var v935 int32
	_ = v935
	var v936 int32
	_ = v936
	var v948 int32
	_ = v948
	var v949 int32
	_ = v949
	var v953 int32
	_ = v953
	var v969 int32
	_ = v969
	var v971 int32
	_ = v971
	var v972 int32
	_ = v972
	var v973 int32
	_ = v973
	var v984 int64
	_ = v984
	var v990 int32
	_ = v990
	var v992 int64
	_ = v992
	var v994 int32
	_ = v994
	var v996 int32
	_ = v996
	var v1006 int64
	_ = v1006
	var v1011 int64
	_ = v1011
	var v1015 int32
	_ = v1015
	var v1017 int64
	_ = v1017
	var v1019 int32
	_ = v1019
	var v1027 int64
	_ = v1027
	var v1051 int64
	_ = v1051
	var v1070 int32
	_ = v1070
	var v1079 int64
	_ = v1079
	var v1080 int64
	_ = v1080
	var v1081 int64
	_ = v1081
	var v1083 int64
	_ = v1083
	var v1085 int64
	_ = v1085
	var v1087 int64
	_ = v1087
	var v1090 int64
	_ = v1090
	var v1092 int64
	_ = v1092
	var v1094 int64
	_ = v1094
	var v1149 int32
	_ = v1149
	var v1150 int32
	_ = v1150
	var v1151 int32
	_ = v1151
	var v1155 int64
	_ = v1155
	var v1165 int64
	_ = v1165
	var v1183 int32
	_ = v1183
	var v1184 int32
	_ = v1184
	var v1185 int32
	_ = v1185
	var v1188 int32
	_ = v1188
	var v1189 int32
	_ = v1189
	var v1193 int32
	_ = v1193
	var v1200 int32
	_ = v1200
	var v1203 int32
	_ = v1203
	var v1206 int32
	_ = v1206
	var v1209 int32
	_ = v1209
	var v1210 int32
	_ = v1210
	var v1211 int32
	_ = v1211
	var v1220 int32
	_ = v1220
	var v1221 int32
	_ = v1221
	var v1222 int32
	_ = v1222
	var v1225 int32
	_ = v1225
	var v1226 int32
	_ = v1226
	var v1228 int32
	_ = v1228
	var v1229 int32
	_ = v1229
	var v1231 int32
	_ = v1231
	var v1233 int32
	_ = v1233
	var v1236 int32
	_ = v1236
	var v1237 int32
	_ = v1237
	var v1238 int32
	_ = v1238
	var v1243 int32
	_ = v1243
	var v1244 int32
	_ = v1244
	var v1245 int32
	_ = v1245
	var v1248 int32
	_ = v1248
	var v1249 int32
	_ = v1249
	var v1250 int32
	_ = v1250
	var v1253 int32
	_ = v1253
	var v1254 int32
	_ = v1254
	var v1256 int32
	_ = v1256
	var v1261 int32
	_ = v1261
	var v1276 int32
	_ = v1276
	var v1277 int32
	_ = v1277
	var v1278 int32
	_ = v1278
	var v1279 int32
	_ = v1279
	var v1282 int64
	_ = v1282
	var v1291 int64
	_ = v1291
	var v1320 int32
	_ = v1320
	var v1335 int32
	_ = v1335
	var v1338 int64
	_ = v1338
	var v1341 int64
	_ = v1341
	var v1343 int32
	_ = v1343
	var v1353 int32
	_ = v1353
	var v1354 int32
	_ = v1354
	var v1356 int32
	_ = v1356
	var v1357 int32
	_ = v1357
	var v1359 int32
	_ = v1359
	var v1360 int32
	_ = v1360
	var v1362 int32
	_ = v1362
	var v1363 int32
	_ = v1363
	var v1364 int32
	_ = v1364
	var v1365 int32
	_ = v1365
	var v1366 int64
	_ = v1366
	var v1374 int32
	_ = v1374
	var v1376 int64
	_ = v1376
	var v1393 int32
	_ = v1393
	var v1394 int32
	_ = v1394
	var v1395 int32
	_ = v1395
	var v1398 int32
	_ = v1398
	var v1399 int32
	_ = v1399
	var v1403 int32
	_ = v1403
	var v1410 int32
	_ = v1410
	var v1413 int32
	_ = v1413
	var v1416 int32
	_ = v1416
	var v1419 int32
	_ = v1419
	var v1420 int32
	_ = v1420
	var v1421 int32
	_ = v1421
	var v1422 int32
	_ = v1422
	var v1423 int32
	_ = v1423
	var v1429 int32
	_ = v1429
	var v1436 int32
	_ = v1436
	var v1439 int32
	_ = v1439
	var v1442 int32
	_ = v1442
	var v1445 int32
	_ = v1445
	var v1446 int32
	_ = v1446
	var v1447 int32
	_ = v1447
	var v1448 int32
	_ = v1448
	var v1450 int64
	_ = v1450
	var v1457 int32
	_ = v1457
	var v1476 int64
	_ = v1476
	var v1480 int32
	_ = v1480
	var v1481 int32
	_ = v1481
	var v1482 int32
	_ = v1482
	var v1484 int32
	_ = v1484
	var v1489 int32
	_ = v1489
	var v1490 int32
	_ = v1490
	var v1492 int64
	_ = v1492
	var v1493 int64
	_ = v1493
	var v1496 int64
	_ = v1496
	var v1514 int32
	_ = v1514
	var v1520 int32
	_ = v1520
	var v1526 int32
	_ = v1526
	var v1532 int32
	_ = v1532
	var v1538 int32
	_ = v1538
	var v1553 int32
	_ = v1553
	v22 = m.G0
	v24 = v22 - int32(368)
	m.G0 = v24
	if l4 == int32(0) {
		goto L8
	} else {
		goto L9
	}
L1:
	;
	m.G0 = v24 + int32(368)
	return v1553
L2:
	;
	goto L299
L3:
	;
	if l2 < int64(1) {
		v160 = int32(0)
		goto L22
	} else {
		goto L23
	}
L4:
	;
	if base.Ui64(v75) < base.Ui64(v73) {
		goto L2
	} else {
		goto L20
	}
L5:
	;
	v71 = l0 + int32(16)
	if base.Ui64(v67) < base.Ui64(v69) {
		v82 = v68
		v83 = v69
		v85 = v71
		goto L3
	} else {
		goto L19
	}
L6:
	;
	v66 = *(*int64)(unsafe.Add(mBase, uint32(l4)+8))
	v67 = v33
	v68 = v66
	v69 = v34
	goto L5
L7:
	;
	v73 = v58
	v74 = v61 + int64(1)
	v75 = v58
	v77 = l0 + int32(16)
	goto L4
L8:
	;
	v45 = *(*int64)(unsafe.Add(mBase, _consts[12]))
	goto L15
L9:
	;
	if l5 == int32(0) {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	v33 = *(*int64)(unsafe.Add(mBase, uint32(l0)+16))
	v34 = *(*int64)(unsafe.Add(mBase, uint32(l4)))
	if v33 != v34 {
		goto L6
	} else {
		goto L12
	}
L11:
	;
	v30 = *(*int64)(unsafe.Add(mBase, uint32(l0)+16))
	v31 = *(*int64)(unsafe.Add(mBase, uint32(l4)+8))
	v32 = *(*int64)(unsafe.Add(mBase, uint32(l4)))
	v67 = v30
	v68 = v31
	v69 = v32
	goto L5
L12:
	;
	v36 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	if v36 != int64(-1) {
		v58 = v33
		v61 = v36
		goto L7
	} else {
		goto L13
	}
L13:
	;
	goto L14
L14:
	;
	*(*int32)(unsafe.Add(mBase, _consts[18])) = int32(18)
	v1553 = int32(-1)
	goto L1
L15:
	;
	v46 = *(*int64)(unsafe.Add(mBase, uint32(l0)+16))
	if base.Ui64(v46) < base.Ui64(v45) {
		v67 = v46
		v68 = int64(0)
		v69 = v45
		goto L5
	} else {
		goto L16
	}
L16:
	;
	v48 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	if v48 != int64(-1) {
		v58 = v46
		v61 = v48
		goto L7
	} else {
		goto L17
	}
L17:
	;
	if v46 == int64(-1) {
		goto L2
	} else {
		goto L18
	}
L18:
	;
	v82 = int64(0)
	v83 = v46 + int64(1)
	v85 = l0 + int32(16)
	goto L3
L19:
	;
	v73 = v67
	v74 = v68
	v75 = v69
	v77 = v71
	goto L4
L20:
	;
	v79 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	if base.Ui64(v74) <= base.Ui64(v79) {
		goto L2
	} else {
		goto L21
	}
L21:
	;
	v82 = v74
	v83 = v75
	v85 = v77
	goto L3
L22:
	;
	v177 = v24 + int32(56)
	v178 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v177)+4)) = v178
	*(*int32)(unsafe.Add(mBase, uint32(v177))) = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v177)+20)) = int32(128)
	v184 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v177)+12)) = v184
	*(*int64)(unsafe.Add(mBase, uint32(v177)+296)) = v184
	*(*int64)(unsafe.Add(mBase, uint32(v177)+160)) = int64(137438953472)
	*(*int32)(unsafe.Add(mBase, uint32(v177)+8)) = v24 + int32(80)
	*(*int32)(unsafe.Add(mBase, uint32(v177)+156)) = v24 + int32(224)
	goto L38
L23:
	;
	v89 = int64(1)
	v90 = l2 << (uint(v89) % 64)
	if v89 < v90 {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	v94 = v90
	goto L26
L25:
	;
	v94 = v89
	goto L26
L26:
	;
	v102 = int32(0)
	v104 = int64(0)
	goto L27
L27:
	;
	v123 = *(*int32)(unsafe.Add(mBase, uint32(l1+base.I32_wrap_i64(v104)<<(uint(int32(2))%32))))
	v124 = F_objectGetVal(m, v123)
	mBase = m.M
	v127 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v124+int32(-1)))))
	switch v127 & int32(7) {
	case 0:
		goto L34
	case 1:
		goto L33
	case 2:
		goto L32
	case 3:
		goto L31
	case 4:
		goto L30
	default:
		v144 = int32(0)
		goto L29
	}
L28:
	;
	if base.Ui32(v145) <= base.Ui32(int32(1073741824)) {
		v160 = v145
		goto L22
	} else {
		goto L36
	}
L29:
	;
	v145 = v144 + v102
	v147 = v104 + int64(1)
	if v147 != v94 {
		v102 = v145
		v104 = v147
		goto L27
	} else {
		goto L35
	}
L30:
	;
	v143 = *(*int32)(unsafe.Add(mBase, uint32(v124+int32(-17))))
	v144 = v143
	goto L29
L31:
	;
	v140 = *(*int32)(unsafe.Add(mBase, uint32(v124+int32(-9))))
	v144 = v140
	goto L29
L32:
	;
	v137 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v124+int32(-5)))))
	v144 = v137
	goto L29
L33:
	;
	v134 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v124+int32(-3)))))
	v144 = v134
	goto L29
L34:
	;
	v144 = int32(base.Ui32(v127) >> (uint(int32(3)) % 32))
	goto L29
L35:
	;
	goto L28
L36:
	;
	goto L37
L37:
	;
	*(*int32)(unsafe.Add(mBase, _consts[18])) = int32(68)
	v1553 = int32(-1)
	goto L1
L38:
	;
	v199 = int32(0)
	v201 = F_raxSeek(m, v24+int32(56), int32(_a1933), v199, v199)
	mBase = m.M
	v204 = m.ExcPending
	if v204 != 0 {
		goto L39
	} else {
		goto L40
	}
L39:
	;
	return int32(0)
L40:
	;
	v207 = *(*int32)(unsafe.Add(mBase, uint32(v24+int32(56))))
	goto L50
L41:
	;
	F__serverAssert(m, int32(_a2401), int32(_a2402), int32(282))
	mBase = m.M
	v1538 = m.ExcPending
	if v1538 != 0 {
		goto L39
	} else {
		goto L298
	}
L42:
	;
	F__serverAssert(m, int32(_a2401), int32(_a2402), int32(282))
	mBase = m.M
	v1532 = m.ExcPending
	if v1532 != 0 {
		goto L39
	} else {
		goto L297
	}
L43:
	;
	F__serverAssert(m, int32(_a2404), int32(_a2402), int32(578))
	mBase = m.M
	v1526 = m.ExcPending
	if v1526 != 0 {
		goto L39
	} else {
		goto L296
	}
L44:
	;
	F__serverAssert(m, int32(_a2401), int32(_a2402), int32(282))
	mBase = m.M
	v1520 = m.ExcPending
	if v1520 != 0 {
		goto L39
	} else {
		goto L295
	}
L45:
	;
	F__serverAssert(m, int32(_a2401), int32(_a2402), int32(282))
	mBase = m.M
	v1514 = m.ExcPending
	if v1514 != 0 {
		goto L39
	} else {
		goto L294
	}
L46:
	;
	v1353 = F_lpAppendInteger(m, v1335, base.I64_extend_i32_u(v1343))
	mBase = m.M
	v1354 = m.ExcPending
	if v1354 != 0 {
		goto L39
	} else {
		goto L256
	}
L47:
	;
	v751 = *(*int32)(unsafe.Add(mBase, uint32(v24)+72))
	if v751 != int32(16) {
		goto L43
	} else {
		goto L153
	}
L48:
	;
	v569 = int64(56)
	v571 = int64(65280)
	v573 = int64(40)
	v576 = int64(16711680)
	v578 = int64(24)
	v580 = int64(4278190080)
	v582 = int64(8)
	*(*int64)(unsafe.Add(mBase, uint32(v24)+40)) = v82<<(uint(v569)%64) | v82&v571<<(uint(v573)%64) | (v82&v576<<(uint(v578)%64) | v82&v580<<(uint(v582)%64)) | (int64(base.Ui64(v82)>>(uint(v582)%64))&v580 | int64(base.Ui64(v82)>>(uint(v578)%64))&v576 | (int64(base.Ui64(v82)>>(uint(v573)%64))&v571 | int64(base.Ui64(v82)>>(uint(v569)%64))))
	*(*int64)(unsafe.Add(mBase, uint32(v24)+32)) = v83<<(uint(v569)%64) | v83&v571<<(uint(v573)%64) | (v83&v576<<(uint(v578)%64) | v83&v580<<(uint(v582)%64)) | (int64(base.Ui64(v83)>>(uint(v582)%64))&v580 | int64(base.Ui64(v83)>>(uint(v578)%64))&v576 | (int64(base.Ui64(v83)>>(uint(v573)%64))&v571 | int64(base.Ui64(v83)>>(uint(v569)%64))))
	v643 = *(*int32)(unsafe.Add(mBase, _consts[864]))
	if base.Ui32(v643+int32(-1)) < base.Ui32(int32(4095)) {
		goto L132
	} else {
		goto L133
	}
L49:
	;
	v216 = *(*int32)(unsafe.Add(mBase, uint32(v24)+68))
	v217 = *(*int32)(unsafe.Add(mBase, uint32(v216)))
	goto L53
L50:
	;
	if v207&int32(2) == int32(0) {
		goto L49
	} else {
		goto L51
	}
L51:
	;
	F_raxStop(m, v24+int32(56))
	mBase = m.M
	v215 = m.ExcPending
	if v215 != 0 {
		goto L39
	} else {
		goto L52
	}
L52:
	;
	goto L48
L53:
	;
	F_raxStop(m, v24+int32(56))
	mBase = m.M
	v221 = m.ExcPending
	if v221 != 0 {
		goto L39
	} else {
		goto L54
	}
L54:
	;
	if v216 == int32(0) {
		goto L48
	} else {
		goto L55
	}
L55:
	;
	v227 = *(*int32)(unsafe.Add(mBase, _consts[864]))
	if base.Ui32(v227+int32(-1073741825)) < base.Ui32(int32(-1073741824)) {
		goto L57
	} else {
		goto L58
	}
L56:
	;
	v554 = F_lpShrinkToFit(m, v216)
	mBase = m.M
	v555 = m.ExcPending
	if v555 != 0 {
		goto L39
	} else {
		goto L129
	}
L57:
	;
	v232 = int32(1073741824)
	goto L59
L58:
	;
	v232 = v227
	goto L59
L59:
	;
	if base.Ui32(v232) <= base.Ui32(v217+v160) {
		goto L56
	} else {
		goto L60
	}
L60:
	;
	v235 = *(*int64)(unsafe.Add(mBase, _consts[865]))
	if v235 == int64(0) {
		goto L47
	} else {
		goto L61
	}
L61:
	;
	v238 = F_lpFirst(m, v216)
	mBase = m.M
	v239 = m.ExcPending
	if v239 != 0 {
		goto L39
	} else {
		goto L64
	}
L62:
	;
	v392 = F_lpNext(m, v216, v238)
	mBase = m.M
	v393 = m.ExcPending
	if v393 != 0 {
		goto L39
	} else {
		goto L97
	}
L63:
	;
	v244 = *(*int32)(unsafe.Add(mBase, uint32(v24)))
	v246 = v24 + int32(360)
	v247 = int32(0)
	if base.Ui32(v244+int32(-21)) < base.Ui32(int32(-20)) {
		v381 = v247
		goto L68
	} else {
		goto L69
	}
L64:
	;
	v241 = F_lpGet(m, v238, v24, int32(0))
	mBase = m.M
	v242 = m.ExcPending
	if v242 != 0 {
		goto L39
	} else {
		goto L65
	}
L65:
	;
	if v241 != 0 {
		goto L63
	} else {
		goto L66
	}
L66:
	;
	v243 = *(*int64)(unsafe.Add(mBase, uint32(v24)))
	v391 = v243
	goto L62
L67:
	;
	if v381 == int32(0) {
		goto L45
	} else {
		goto L94
	}
L68:
	;
	goto L67
L69:
	;
	v259 = int32(1)
	v260 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v241))))
	if v244 != v259 {
		goto L72
	} else {
		goto L73
	}
L70:
	;
	v381 = int32(1)
	goto L68
L71:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v246))) = v362
	goto L70
L72:
	;
	if v260&int32(255) == int32(45) {
		goto L77
	} else {
		goto L78
	}
L73:
	;
	v264 = v260 + int32(-48)
	if base.Ui32(int32(9)) < base.Ui32(v264&int32(255)) {
		v381 = v247
		goto L68
	} else {
		goto L74
	}
L74:
	;
	if v246 == int32(0) {
		goto L70
	} else {
		goto L75
	}
L75:
	;
	v362 = base.I64_extend_i32_u(v264) & int64(255)
	goto L71
L76:
	;
	if base.Ui32(int32(8)) < base.Ui32((v283+int32(-49))&int32(255)) {
		v381 = v247
		goto L68
	} else {
		goto L79
	}
L77:
	;
	v280 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v241)+1)))
	v282 = int32(2)
	v283 = v280
	v284 = v241 + int32(1)
	goto L76
L78:
	;
	v282 = v259
	v283 = v260
	v284 = v241
	goto L76
L79:
	;
	v295 = base.I64_extend_i32_u(v283+int32(-48)) & int64(255)
	if base.Ui32(v244) <= base.Ui32(v282) {
		v338 = v295
		goto L80
	} else {
		goto L81
	}
L80:
	;
	if v260&int32(255) != int32(45) {
		goto L88
	} else {
		goto L89
	}
L81:
	;
	v301 = v282
	v303 = v295
	v305 = v284
	goto L82
L82:
	;
	v307 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v305)+1)))
	if base.Ui32((v307+int32(-58))&int32(255)) < base.Ui32(int32(246)) {
		v381 = v247
		goto L68
	} else {
		goto L84
	}
L83:
	;
	v338 = v328
	goto L80
L84:
	;
	if base.Ui64(int64(1844674407370955161)) < base.Ui64(v303) {
		v381 = v247
		goto L68
	} else {
		goto L85
	}
L85:
	;
	v317 = v303 * int64(10)
	v322 = base.I64_extend_i32_u(v307+int32(-48)) & int64(255)
	if base.Ui64(v322^int64(-1)) < base.Ui64(v317) {
		v381 = v247
		goto L68
	} else {
		goto L86
	}
L86:
	;
	v326 = int32(1)
	v328 = v317 + v322
	v330 = v301 + v326
	if v330 != v244 {
		v301 = v330
		v303 = v328
		v305 = v305 + v326
		goto L82
	} else {
		goto L87
	}
L87:
	;
	goto L83
L88:
	;
	if v338 < int64(0) {
		v381 = v247
		goto L68
	} else {
		goto L92
	}
L89:
	;
	if base.Ui64(int64(-9223372036854775807-1)) < base.Ui64(v338) {
		v381 = v247
		goto L68
	} else {
		goto L90
	}
L90:
	;
	if v246 == int32(0) {
		goto L70
	} else {
		goto L91
	}
L91:
	;
	v362 = int64(0) - v338
	goto L71
L92:
	;
	if v246 == int32(0) {
		goto L70
	} else {
		goto L93
	}
L93:
	;
	v362 = v338
	goto L71
L94:
	;
	v390 = *(*int64)(unsafe.Add(mBase, uint32(v24)+360))
	v391 = v390
	goto L62
L95:
	;
	v548 = *(*int64)(unsafe.Add(mBase, _consts[865]))
	if v545+v391 < v548 {
		goto L47
	} else {
		goto L128
	}
L96:
	;
	v398 = *(*int32)(unsafe.Add(mBase, uint32(v24)))
	v400 = v24 + int32(360)
	v401 = int32(0)
	if base.Ui32(v398+int32(-21)) < base.Ui32(int32(-20)) {
		v535 = v401
		goto L101
	} else {
		goto L102
	}
L97:
	;
	v395 = F_lpGet(m, v392, v24, int32(0))
	mBase = m.M
	v396 = m.ExcPending
	if v396 != 0 {
		goto L39
	} else {
		goto L98
	}
L98:
	;
	if v395 != 0 {
		goto L96
	} else {
		goto L99
	}
L99:
	;
	v397 = *(*int64)(unsafe.Add(mBase, uint32(v24)))
	v545 = v397
	goto L95
L100:
	;
	if v535 == int32(0) {
		goto L44
	} else {
		goto L127
	}
L101:
	;
	goto L100
L102:
	;
	v413 = int32(1)
	v414 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v395))))
	if v398 != v413 {
		goto L105
	} else {
		goto L106
	}
L103:
	;
	v535 = int32(1)
	goto L101
L104:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v400))) = v516
	goto L103
L105:
	;
	if v414&int32(255) == int32(45) {
		goto L110
	} else {
		goto L111
	}
L106:
	;
	v418 = v414 + int32(-48)
	if base.Ui32(int32(9)) < base.Ui32(v418&int32(255)) {
		v535 = v401
		goto L101
	} else {
		goto L107
	}
L107:
	;
	if v400 == int32(0) {
		goto L103
	} else {
		goto L108
	}
L108:
	;
	v516 = base.I64_extend_i32_u(v418) & int64(255)
	goto L104
L109:
	;
	if base.Ui32(int32(8)) < base.Ui32((v437+int32(-49))&int32(255)) {
		v535 = v401
		goto L101
	} else {
		goto L112
	}
L110:
	;
	v434 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v395)+1)))
	v436 = int32(2)
	v437 = v434
	v438 = v395 + int32(1)
	goto L109
L111:
	;
	v436 = v413
	v437 = v414
	v438 = v395
	goto L109
L112:
	;
	v449 = base.I64_extend_i32_u(v437+int32(-48)) & int64(255)
	if base.Ui32(v398) <= base.Ui32(v436) {
		v492 = v449
		goto L113
	} else {
		goto L114
	}
L113:
	;
	if v414&int32(255) != int32(45) {
		goto L121
	} else {
		goto L122
	}
L114:
	;
	v455 = v436
	v457 = v449
	v459 = v438
	goto L115
L115:
	;
	v461 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v459)+1)))
	if base.Ui32((v461+int32(-58))&int32(255)) < base.Ui32(int32(246)) {
		v535 = v401
		goto L101
	} else {
		goto L117
	}
L116:
	;
	v492 = v482
	goto L113
L117:
	;
	if base.Ui64(int64(1844674407370955161)) < base.Ui64(v457) {
		v535 = v401
		goto L101
	} else {
		goto L118
	}
L118:
	;
	v471 = v457 * int64(10)
	v476 = base.I64_extend_i32_u(v461+int32(-48)) & int64(255)
	if base.Ui64(v476^int64(-1)) < base.Ui64(v471) {
		v535 = v401
		goto L101
	} else {
		goto L119
	}
L119:
	;
	v480 = int32(1)
	v482 = v471 + v476
	v484 = v455 + v480
	if v484 != v398 {
		v455 = v484
		v457 = v482
		v459 = v459 + v480
		goto L115
	} else {
		goto L120
	}
L120:
	;
	goto L116
L121:
	;
	if v492 < int64(0) {
		v535 = v401
		goto L101
	} else {
		goto L125
	}
L122:
	;
	if base.Ui64(int64(-9223372036854775807-1)) < base.Ui64(v492) {
		v535 = v401
		goto L101
	} else {
		goto L123
	}
L123:
	;
	if v400 == int32(0) {
		goto L103
	} else {
		goto L124
	}
L124:
	;
	v516 = int64(0) - v492
	goto L104
L125:
	;
	if v400 == int32(0) {
		goto L103
	} else {
		goto L126
	}
L126:
	;
	v516 = v492
	goto L104
L127:
	;
	v544 = *(*int64)(unsafe.Add(mBase, uint32(v24)+360))
	v545 = v544
	goto L95
L128:
	;
	goto L56
L129:
	;
	v556 = *(*int32)(unsafe.Add(mBase, uint32(v24)+68))
	if v554 == v556 {
		goto L48
	} else {
		goto L130
	}
L130:
	;
	v558 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v559 = *(*int32)(unsafe.Add(mBase, uint32(v24)+64))
	v560 = *(*int32)(unsafe.Add(mBase, uint32(v24)+72))
	v562 = F_raxInsert(m, v558, v559, v560, v554, int32(0))
	mBase = m.M
	v563 = m.ExcPending
	if v563 != 0 {
		goto L39
	} else {
		goto L131
	}
L131:
	;
	goto L48
L132:
	;
	v649 = v643
	goto L134
L133:
	;
	v649 = int32(4096)
	goto L134
L134:
	;
	v650 = F_lpNew(m, v649)
	mBase = m.M
	v651 = m.ExcPending
	if v651 != 0 {
		goto L39
	} else {
		goto L135
	}
L135:
	;
	v653 = F_lpAppendInteger(m, v650, int64(1))
	mBase = m.M
	v654 = m.ExcPending
	if v654 != 0 {
		goto L39
	} else {
		goto L136
	}
L136:
	;
	v656 = F_lpAppendInteger(m, v653, int64(0))
	mBase = m.M
	v657 = m.ExcPending
	if v657 != 0 {
		goto L39
	} else {
		goto L137
	}
L137:
	;
	v658 = F_lpAppendInteger(m, v656, l2)
	mBase = m.M
	v659 = m.ExcPending
	if v659 != 0 {
		goto L39
	} else {
		goto L138
	}
L138:
	;
	if l2 <= int64(0) {
		v727 = v658
		goto L139
	} else {
		goto L140
	}
L139:
	;
	v737 = F_lpAppendInteger(m, v727, int64(0))
	mBase = m.M
	v738 = m.ExcPending
	if v738 != 0 {
		goto L39
	} else {
		goto L151
	}
L140:
	;
	v669 = int64(0)
	v674 = v658
	goto L141
L141:
	;
	v688 = *(*int32)(unsafe.Add(mBase, uint32(l1+base.I32_wrap_i64(v669)<<(uint(int32(3))%32))))
	v689 = F_objectGetVal(m, v688)
	mBase = m.M
	v692 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v689+int32(-1)))))
	switch v692 & int32(7) {
	case 0:
		goto L148
	case 1:
		goto L147
	case 2:
		goto L146
	case 3:
		goto L145
	case 4:
		goto L144
	default:
		v709 = int32(0)
		goto L143
	}
L142:
	;
	v727 = v710
	goto L139
L143:
	;
	v710 = F_lpAppend(m, v674, v689, v709)
	mBase = m.M
	v711 = m.ExcPending
	if v711 != 0 {
		goto L39
	} else {
		goto L149
	}
L144:
	;
	v708 = *(*int32)(unsafe.Add(mBase, uint32(v689+int32(-17))))
	v709 = v708
	goto L143
L145:
	;
	v705 = *(*int32)(unsafe.Add(mBase, uint32(v689+int32(-9))))
	v709 = v705
	goto L143
L146:
	;
	v702 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v689+int32(-5)))))
	v709 = v702
	goto L143
L147:
	;
	v699 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v689+int32(-3)))))
	v709 = v699
	goto L143
L148:
	;
	v709 = int32(base.Ui32(v692) >> (uint(int32(3)) % 32))
	goto L143
L149:
	;
	v713 = v669 + int64(1)
	if v713 != l2 {
		v669 = v713
		v674 = v710
		goto L141
	} else {
		goto L150
	}
L150:
	;
	goto L142
L151:
	;
	v739 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v744 = F_raxInsert(m, v739, v24+int32(32), int32(16), v737, int32(0))
	mBase = m.M
	v745 = m.ExcPending
	if v745 != 0 {
		goto L39
	} else {
		goto L152
	}
L152:
	;
	v1335 = v737
	v1338 = v82
	v1341 = v83
	v1343 = int32(2)
	goto L46
L153:
	;
	v754 = *(*int32)(unsafe.Add(mBase, uint32(v24)+64))
	v755 = *(*int64)(unsafe.Add(mBase, uint32(v754)))
	*(*int64)(unsafe.Add(mBase, uint32(v24)+32)) = v755
	v759 = *(*int64)(unsafe.Add(mBase, uint32(v754+int32(8))))
	*(*int64)(unsafe.Add(mBase, uint32(v24)+40)) = v759
	v761 = F_lpFirst(m, v216)
	mBase = m.M
	v762 = m.ExcPending
	if v762 != 0 {
		goto L39
	} else {
		goto L154
	}
L154:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+28)) = v761
	v765 = F_lpGet(m, v761, v24, int32(0))
	mBase = m.M
	v766 = m.ExcPending
	if v766 != 0 {
		goto L39
	} else {
		goto L157
	}
L155:
	;
	v920 = F_lpReplaceInteger(m, v216, v24+int32(28), v915+int64(1))
	mBase = m.M
	v921 = m.ExcPending
	if v921 != 0 {
		goto L39
	} else {
		goto L187
	}
L156:
	;
	v768 = *(*int32)(unsafe.Add(mBase, uint32(v24)))
	v770 = v24 + int32(360)
	v771 = int32(0)
	if base.Ui32(v768+int32(-21)) < base.Ui32(int32(-20)) {
		v905 = v771
		goto L160
	} else {
		goto L161
	}
L157:
	;
	if v765 != 0 {
		goto L156
	} else {
		goto L158
	}
L158:
	;
	v767 = *(*int64)(unsafe.Add(mBase, uint32(v24)))
	v915 = v767
	goto L155
L159:
	;
	if v905 == int32(0) {
		goto L42
	} else {
		goto L186
	}
L160:
	;
	goto L159
L161:
	;
	v783 = int32(1)
	v784 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v765))))
	if v768 != v783 {
		goto L164
	} else {
		goto L165
	}
L162:
	;
	v905 = int32(1)
	goto L160
L163:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v770))) = v886
	goto L162
L164:
	;
	if v784&int32(255) == int32(45) {
		goto L169
	} else {
		goto L170
	}
L165:
	;
	v788 = v784 + int32(-48)
	if base.Ui32(int32(9)) < base.Ui32(v788&int32(255)) {
		v905 = v771
		goto L160
	} else {
		goto L166
	}
L166:
	;
	if v770 == int32(0) {
		goto L162
	} else {
		goto L167
	}
L167:
	;
	v886 = base.I64_extend_i32_u(v788) & int64(255)
	goto L163
L168:
	;
	if base.Ui32(int32(8)) < base.Ui32((v807+int32(-49))&int32(255)) {
		v905 = v771
		goto L160
	} else {
		goto L171
	}
L169:
	;
	v804 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v765)+1)))
	v806 = int32(2)
	v807 = v804
	v808 = v765 + int32(1)
	goto L168
L170:
	;
	v806 = v783
	v807 = v784
	v808 = v765
	goto L168
L171:
	;
	v819 = base.I64_extend_i32_u(v807+int32(-48)) & int64(255)
	if base.Ui32(v768) <= base.Ui32(v806) {
		v862 = v819
		goto L172
	} else {
		goto L173
	}
L172:
	;
	if v784&int32(255) != int32(45) {
		goto L180
	} else {
		goto L181
	}
L173:
	;
	v825 = v806
	v827 = v819
	v829 = v808
	goto L174
L174:
	;
	v831 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v829)+1)))
	if base.Ui32((v831+int32(-58))&int32(255)) < base.Ui32(int32(246)) {
		v905 = v771
		goto L160
	} else {
		goto L176
	}
L175:
	;
	v862 = v852
	goto L172
L176:
	;
	if base.Ui64(int64(1844674407370955161)) < base.Ui64(v827) {
		v905 = v771
		goto L160
	} else {
		goto L177
	}
L177:
	;
	v841 = v827 * int64(10)
	v846 = base.I64_extend_i32_u(v831+int32(-48)) & int64(255)
	if base.Ui64(v846^int64(-1)) < base.Ui64(v841) {
		v905 = v771
		goto L160
	} else {
		goto L178
	}
L178:
	;
	v850 = int32(1)
	v852 = v841 + v846
	v854 = v825 + v850
	if v854 != v768 {
		v825 = v854
		v827 = v852
		v829 = v829 + v850
		goto L174
	} else {
		goto L179
	}
L179:
	;
	goto L175
L180:
	;
	if v862 < int64(0) {
		v905 = v771
		goto L160
	} else {
		goto L184
	}
L181:
	;
	if base.Ui64(int64(-9223372036854775807-1)) < base.Ui64(v862) {
		v905 = v771
		goto L160
	} else {
		goto L182
	}
L182:
	;
	if v770 == int32(0) {
		goto L162
	} else {
		goto L183
	}
L183:
	;
	v886 = int64(0) - v862
	goto L163
L184:
	;
	if v770 == int32(0) {
		goto L162
	} else {
		goto L185
	}
L185:
	;
	v886 = v862
	goto L163
L186:
	;
	v914 = *(*int64)(unsafe.Add(mBase, uint32(v24)+360))
	v915 = v914
	goto L155
L187:
	;
	v922 = *(*int32)(unsafe.Add(mBase, uint32(v24)+28))
	v923 = F_lpNext(m, v920, v922)
	mBase = m.M
	v924 = m.ExcPending
	if v924 != 0 {
		goto L39
	} else {
		goto L188
	}
L188:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+28)) = v923
	v926 = F_lpNext(m, v920, v923)
	mBase = m.M
	v927 = m.ExcPending
	if v927 != 0 {
		goto L39
	} else {
		goto L189
	}
L189:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+28)) = v926
	v930 = F_lpGet(m, v926, v24, int32(0))
	mBase = m.M
	v931 = m.ExcPending
	if v931 != 0 {
		goto L39
	} else {
		goto L192
	}
L190:
	;
	v1081 = int64(8)
	v1083 = int64(4278190080)
	v1085 = int64(24)
	v1087 = int64(16711680)
	v1090 = int64(40)
	v1092 = int64(65280)
	v1094 = int64(56)
	v1149 = *(*int32)(unsafe.Add(mBase, uint32(v24)+28))
	v1150 = F_lpNext(m, v920, v1149)
	mBase = m.M
	v1151 = m.ExcPending
	if v1151 != 0 {
		goto L39
	} else {
		goto L222
	}
L191:
	;
	v933 = *(*int32)(unsafe.Add(mBase, uint32(v24)))
	v935 = v24 + int32(360)
	v936 = int32(0)
	if base.Ui32(v933+int32(-21)) < base.Ui32(int32(-20)) {
		v1070 = v936
		goto L195
	} else {
		goto L196
	}
L192:
	;
	if v930 != 0 {
		goto L191
	} else {
		goto L193
	}
L193:
	;
	v932 = *(*int64)(unsafe.Add(mBase, uint32(v24)))
	v1080 = v932
	goto L190
L194:
	;
	if v1070 == int32(0) {
		goto L41
	} else {
		goto L221
	}
L195:
	;
	goto L194
L196:
	;
	v948 = int32(1)
	v949 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v930))))
	if v933 != v948 {
		goto L199
	} else {
		goto L200
	}
L197:
	;
	v1070 = int32(1)
	goto L195
L198:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v935))) = v1051
	goto L197
L199:
	;
	if v949&int32(255) == int32(45) {
		goto L204
	} else {
		goto L205
	}
L200:
	;
	v953 = v949 + int32(-48)
	if base.Ui32(int32(9)) < base.Ui32(v953&int32(255)) {
		v1070 = v936
		goto L195
	} else {
		goto L201
	}
L201:
	;
	if v935 == int32(0) {
		goto L197
	} else {
		goto L202
	}
L202:
	;
	v1051 = base.I64_extend_i32_u(v953) & int64(255)
	goto L198
L203:
	;
	if base.Ui32(int32(8)) < base.Ui32((v972+int32(-49))&int32(255)) {
		v1070 = v936
		goto L195
	} else {
		goto L206
	}
L204:
	;
	v969 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v930)+1)))
	v971 = int32(2)
	v972 = v969
	v973 = v930 + int32(1)
	goto L203
L205:
	;
	v971 = v948
	v972 = v949
	v973 = v930
	goto L203
L206:
	;
	v984 = base.I64_extend_i32_u(v972+int32(-48)) & int64(255)
	if base.Ui32(v933) <= base.Ui32(v971) {
		v1027 = v984
		goto L207
	} else {
		goto L208
	}
L207:
	;
	if v949&int32(255) != int32(45) {
		goto L215
	} else {
		goto L216
	}
L208:
	;
	v990 = v971
	v992 = v984
	v994 = v973
	goto L209
L209:
	;
	v996 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v994)+1)))
	if base.Ui32((v996+int32(-58))&int32(255)) < base.Ui32(int32(246)) {
		v1070 = v936
		goto L195
	} else {
		goto L211
	}
L210:
	;
	v1027 = v1017
	goto L207
L211:
	;
	if base.Ui64(int64(1844674407370955161)) < base.Ui64(v992) {
		v1070 = v936
		goto L195
	} else {
		goto L212
	}
L212:
	;
	v1006 = v992 * int64(10)
	v1011 = base.I64_extend_i32_u(v996+int32(-48)) & int64(255)
	if base.Ui64(v1011^int64(-1)) < base.Ui64(v1006) {
		v1070 = v936
		goto L195
	} else {
		goto L213
	}
L213:
	;
	v1015 = int32(1)
	v1017 = v1006 + v1011
	v1019 = v990 + v1015
	if v1019 != v933 {
		v990 = v1019
		v992 = v1017
		v994 = v994 + v1015
		goto L209
	} else {
		goto L214
	}
L214:
	;
	goto L210
L215:
	;
	if v1027 < int64(0) {
		v1070 = v936
		goto L195
	} else {
		goto L219
	}
L216:
	;
	if base.Ui64(int64(-9223372036854775807-1)) < base.Ui64(v1027) {
		v1070 = v936
		goto L195
	} else {
		goto L217
	}
L217:
	;
	if v935 == int32(0) {
		goto L197
	} else {
		goto L218
	}
L218:
	;
	v1051 = int64(0) - v1027
	goto L198
L219:
	;
	if v935 == int32(0) {
		goto L197
	} else {
		goto L220
	}
L220:
	;
	v1051 = v1027
	goto L198
L221:
	;
	v1079 = *(*int64)(unsafe.Add(mBase, uint32(v24)+360))
	v1080 = v1079
	goto L190
L222:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+28)) = v1150
	if l2 != v1080 {
		v1320 = int32(0)
		goto L223
	} else {
		goto L224
	}
L223:
	;
	v1335 = v920
	v1338 = v759<<(uint(v1094)%64) | v759&v1092<<(uint(v1090)%64) | (v759&v1087<<(uint(v1085)%64) | v759&v1083<<(uint(v1081)%64)) | (int64(base.Ui64(v759)>>(uint(v1081)%64))&v1083 | int64(base.Ui64(v759)>>(uint(v1085)%64))&v1087 | (int64(base.Ui64(v759)>>(uint(v1090)%64))&v1092 | int64(base.Ui64(v759)>>(uint(v1094)%64))))
	v1341 = v755<<(uint(v1094)%64) | v755&v1092<<(uint(v1090)%64) | (v755&v1087<<(uint(v1085)%64) | v755&v1083<<(uint(v1081)%64)) | (int64(base.Ui64(v755)>>(uint(v1081)%64))&v1083 | int64(base.Ui64(v755)>>(uint(v1085)%64))&v1087 | (int64(base.Ui64(v755)>>(uint(v1090)%64))&v1092 | int64(base.Ui64(v755)>>(uint(v1094)%64))))
	v1343 = v1320
	goto L46
L224:
	;
	v1155 = int64(0)
	if l2 < int64(1) {
		v1291 = v1155
		goto L225
	} else {
		goto L226
	}
L225:
	;
	v1320 = base.B2i32(v1291 == l2) << (uint(int32(1)) % 32)
	goto L223
L226:
	;
	v1165 = v1155
	goto L227
L227:
	;
	v1183 = *(*int32)(unsafe.Add(mBase, uint32(l1+base.I32_wrap_i64(v1165)<<(uint(int32(3))%32))))
	v1184 = F_objectGetVal(m, v1183)
	mBase = m.M
	v1185 = *(*int32)(unsafe.Add(mBase, uint32(v24)+28))
	v1188 = F_lpGet(m, v1185, v24+int32(360), v24)
	mBase = m.M
	v1189 = m.ExcPending
	if v1189 != 0 {
		goto L39
	} else {
		goto L229
	}
L228:
	;
	v1291 = l2
	goto L225
L229:
	;
	v1193 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1184+int32(-1)))))
	switch v1193 & int32(7) {
	case 0:
		goto L235
	case 1:
		goto L234
	case 2:
		goto L233
	case 3:
		goto L232
	case 4:
		goto L231
	default:
		v1210 = int32(0)
		goto L230
	}
L230:
	;
	v1211 = *(*int32)(unsafe.Add(mBase, uint32(v24)+360))
	if v1210 != v1211 {
		v1291 = v1165
		goto L225
	} else {
		goto L236
	}
L231:
	;
	v1209 = *(*int32)(unsafe.Add(mBase, uint32(v1184+int32(-17))))
	v1210 = v1209
	goto L230
L232:
	;
	v1206 = *(*int32)(unsafe.Add(mBase, uint32(v1184+int32(-9))))
	v1210 = v1206
	goto L230
L233:
	;
	v1203 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1184+int32(-5)))))
	v1210 = v1203
	goto L230
L234:
	;
	v1200 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1184+int32(-3)))))
	v1210 = v1200
	goto L230
L235:
	;
	v1210 = int32(base.Ui32(v1193) >> (uint(int32(3)) % 32))
	goto L230
L236:
	;
	if base.Ui32(v1210) < base.Ui32(int32(4)) {
		v1236 = v1188
		v1237 = v1184
		v1238 = v1210
		goto L240
	} else {
		goto L241
	}
L237:
	;
	if v1276 != 0 {
		v1291 = v1165
		goto L225
	} else {
		goto L253
	}
L238:
	;
	v1276 = int32(0)
	goto L237
L239:
	;
	v1248 = v1243
	v1249 = v1244
	v1250 = v1245
	goto L249
L240:
	;
	if v1238 == int32(0) {
		goto L238
	} else {
		goto L247
	}
L241:
	;
	if (v1184|v1188)&int32(3) != 0 {
		v1243 = v1188
		v1244 = v1184
		v1245 = v1210
		goto L239
	} else {
		goto L242
	}
L242:
	;
	v1220 = v1188
	v1221 = v1184
	v1222 = v1210
	goto L243
L243:
	;
	v1225 = *(*int32)(unsafe.Add(mBase, uint32(v1220)))
	v1226 = *(*int32)(unsafe.Add(mBase, uint32(v1221)))
	if v1225 != v1226 {
		v1243 = v1220
		v1244 = v1221
		v1245 = v1222
		goto L239
	} else {
		goto L245
	}
L244:
	;
	v1236 = v1231
	v1237 = v1229
	v1238 = v1233
	goto L240
L245:
	;
	v1228 = int32(4)
	v1229 = v1221 + v1228
	v1231 = v1220 + v1228
	v1233 = v1222 + int32(-4)
	if base.Ui32(int32(3)) < base.Ui32(v1233) {
		v1220 = v1231
		v1221 = v1229
		v1222 = v1233
		goto L243
	} else {
		goto L246
	}
L246:
	;
	goto L244
L247:
	;
	v1243 = v1236
	v1244 = v1237
	v1245 = v1238
	goto L239
L248:
	;
	v1276 = v1253 - v1254
	goto L237
L249:
	;
	v1253 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1248))))
	v1254 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1249))))
	if v1253 != v1254 {
		goto L248
	} else {
		goto L251
	}
L251:
	;
	v1256 = int32(1)
	v1261 = v1250 + int32(-1)
	if v1261 == int32(0) {
		goto L238
	} else {
		goto L252
	}
L252:
	;
	v1248 = v1248 + v1256
	v1249 = v1249 + v1256
	v1250 = v1261
	goto L249
L253:
	;
	v1277 = *(*int32)(unsafe.Add(mBase, uint32(v24)+28))
	v1278 = F_lpNext(m, v920, v1277)
	mBase = m.M
	v1279 = m.ExcPending
	if v1279 != 0 {
		goto L39
	} else {
		goto L254
	}
L254:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+28)) = v1278
	v1282 = v1165 + int64(1)
	if v1282 != l2 {
		v1165 = v1282
		goto L227
	} else {
		goto L255
	}
L255:
	;
	goto L228
L256:
	;
	v1356 = F_lpAppendInteger(m, v1353, v83-v1341)
	mBase = m.M
	v1357 = m.ExcPending
	if v1357 != 0 {
		goto L39
	} else {
		goto L257
	}
L257:
	;
	v1359 = F_lpAppendInteger(m, v1356, v82-v1338)
	mBase = m.M
	v1360 = m.ExcPending
	if v1360 != 0 {
		goto L39
	} else {
		goto L258
	}
L258:
	;
	v1362 = v1343 & int32(2)
	if v1362 != 0 {
		v1365 = v1359
		goto L259
	} else {
		goto L260
	}
L259:
	;
	v1366 = int64(0)
	if l2 <= v1366 {
		v1457 = v1365
		goto L262
	} else {
		goto L263
	}
L260:
	;
	v1363 = F_lpAppendInteger(m, v1359, l2)
	mBase = m.M
	v1364 = m.ExcPending
	if v1364 != 0 {
		goto L39
	} else {
		goto L261
	}
L261:
	;
	v1365 = v1363
	goto L259
L262:
	;
	if v1362 != 0 {
		goto L284
	} else {
		goto L285
	}
L263:
	;
	v1374 = v1365
	v1376 = v1366
	goto L264
L264:
	;
	v1393 = l1 + base.I32_wrap_i64(v1376)<<(uint(int32(3))%32)
	v1394 = *(*int32)(unsafe.Add(mBase, uint32(v1393)))
	v1395 = F_objectGetVal(m, v1394)
	mBase = m.M
	v1398 = *(*int32)(unsafe.Add(mBase, uint32(v1393+int32(4))))
	v1399 = F_objectGetVal(m, v1398)
	mBase = m.M
	if v1362 != 0 {
		v1423 = v1374
		goto L266
	} else {
		goto L267
	}
L265:
	;
	v1457 = v1447
	goto L262
L266:
	;
	v1429 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1399+int32(-1)))))
	switch v1429 & int32(7) {
	case 0:
		goto L280
	case 1:
		goto L279
	case 2:
		goto L278
	case 3:
		goto L277
	case 4:
		goto L276
	default:
		v1446 = int32(0)
		goto L275
	}
L267:
	;
	v1403 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1395+int32(-1)))))
	switch v1403 & int32(7) {
	case 0:
		goto L273
	case 1:
		goto L272
	case 2:
		goto L271
	case 3:
		goto L270
	case 4:
		goto L269
	default:
		v1420 = int32(0)
		goto L268
	}
L268:
	;
	v1421 = F_lpAppend(m, v1374, v1395, v1420)
	mBase = m.M
	v1422 = m.ExcPending
	if v1422 != 0 {
		goto L39
	} else {
		goto L274
	}
L269:
	;
	v1419 = *(*int32)(unsafe.Add(mBase, uint32(v1395+int32(-17))))
	v1420 = v1419
	goto L268
L270:
	;
	v1416 = *(*int32)(unsafe.Add(mBase, uint32(v1395+int32(-9))))
	v1420 = v1416
	goto L268
L271:
	;
	v1413 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1395+int32(-5)))))
	v1420 = v1413
	goto L268
L272:
	;
	v1410 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1395+int32(-3)))))
	v1420 = v1410
	goto L268
L273:
	;
	v1420 = int32(base.Ui32(v1403) >> (uint(int32(3)) % 32))
	goto L268
L274:
	;
	v1423 = v1421
	goto L266
L275:
	;
	v1447 = F_lpAppend(m, v1423, v1399, v1446)
	mBase = m.M
	v1448 = m.ExcPending
	if v1448 != 0 {
		goto L39
	} else {
		goto L281
	}
L276:
	;
	v1445 = *(*int32)(unsafe.Add(mBase, uint32(v1399+int32(-17))))
	v1446 = v1445
	goto L275
L277:
	;
	v1442 = *(*int32)(unsafe.Add(mBase, uint32(v1399+int32(-9))))
	v1446 = v1442
	goto L275
L278:
	;
	v1439 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1399+int32(-5)))))
	v1446 = v1439
	goto L275
L279:
	;
	v1436 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1399+int32(-3)))))
	v1446 = v1436
	goto L275
L280:
	;
	v1446 = int32(base.Ui32(v1429) >> (uint(int32(3)) % 32))
	goto L275
L281:
	;
	v1450 = v1376 + int64(1)
	if v1450 != l2 {
		v1374 = v1447
		v1376 = v1450
		goto L264
	} else {
		goto L282
	}
L282:
	;
	goto L265
L283:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l0)+16)) = v83
	v1492 = *(*int64)(unsafe.Add(mBase, uint32(l0)+8))
	v1493 = int64(1)
	*(*int64)(unsafe.Add(mBase, uint32(l0)+8)) = v1492 + v1493
	v1496 = *(*int64)(unsafe.Add(mBase, uint32(l0)+64))
	*(*int64)(unsafe.Add(mBase, uint32(l0)+64)) = v1496 + v1493
	*(*int64)(unsafe.Add(mBase, uint32(v85)+8)) = v82
	if v1492 != int64(0) {
		goto L290
	} else {
		goto L291
	}
L284:
	;
	v1476 = int64(0)
	goto L286
L285:
	;
	v1476 = l2 + int64(1)
	goto L286
L286:
	;
	v1480 = F_lpAppendInteger(m, v1457, l2+v1476+int64(3))
	mBase = m.M
	v1481 = m.ExcPending
	if v1481 != 0 {
		goto L39
	} else {
		goto L287
	}
L287:
	;
	v1482 = *(*int32)(unsafe.Add(mBase, uint32(v24)+68))
	if v1480 == v1482 {
		goto L283
	} else {
		goto L288
	}
L288:
	;
	v1484 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v1489 = F_raxInsert(m, v1484, v24+int32(32), int32(16), v1480, int32(0))
	mBase = m.M
	v1490 = m.ExcPending
	if v1490 != 0 {
		goto L39
	} else {
		goto L289
	}
L289:
	;
	goto L283
L290:
	;
	if l3 == int32(0) {
		goto L292
	} else {
		goto L293
	}
L291:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l0)+40)) = v82
	*(*int64)(unsafe.Add(mBase, uint32(l0)+32)) = v83
	goto L290
L292:
	;
	v1553 = int32(0)
	goto L1
L293:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l3)+8)) = v82
	*(*int64)(unsafe.Add(mBase, uint32(l3))) = v83
	goto L292
L294:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L295:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L296:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L297:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L298:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L299:
	;
	*(*int32)(unsafe.Add(mBase, _consts[18])) = int32(18)
	v1553 = int32(-1)
	goto L1
}
func F_streamDecodeID(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v5 int64
	_ = v5
	var v6 int64
	_ = v6
	var v7 int64
	_ = v7
	var v9 int64
	_ = v9
	var v11 int64
	_ = v11
	var v14 int64
	_ = v14
	var v16 int64
	_ = v16
	var v18 int64
	_ = v18
	var v20 int64
	_ = v20
	v5 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	v6 = *(*int64)(unsafe.Add(mBase, uint32(l0)+8))
	v7 = int64(56)
	v9 = int64(65280)
	v11 = int64(40)
	v14 = int64(16711680)
	v16 = int64(24)
	v18 = int64(4278190080)
	v20 = int64(8)
	*(*int64)(unsafe.Add(mBase, uint32(l1)+8)) = v6<<(uint(v7)%64) | v6&v9<<(uint(v11)%64) | (v6&v14<<(uint(v16)%64) | v6&v18<<(uint(v20)%64)) | (int64(base.Ui64(v6)>>(uint(v20)%64))&v18 | int64(base.Ui64(v6)>>(uint(v16)%64))&v14 | (int64(base.Ui64(v6)>>(uint(v11)%64))&v9 | int64(base.Ui64(v6)>>(uint(v7)%64))))
	*(*int64)(unsafe.Add(mBase, uint32(l1))) = v5<<(uint(v7)%64) | v5&v9<<(uint(v11)%64) | (v5&v14<<(uint(v16)%64) | v5&v18<<(uint(v20)%64)) | (int64(base.Ui64(v5)>>(uint(v20)%64))&v18 | int64(base.Ui64(v5)>>(uint(v16)%64))&v14 | (int64(base.Ui64(v5)>>(uint(v11)%64))&v9 | int64(base.Ui64(v5)>>(uint(v7)%64))))
	return
}
func F_streamDelConsumer(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v17 int64
	_ = v17
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
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
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
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	v7 = m.G0
	v9 = v7 - int32(304)
	m.G0 = v9
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v9)+4)) = v11
	*(*int32)(unsafe.Add(mBase, uint32(v9))) = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v9)+20)) = int32(128)
	v17 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v9)+12)) = v17
	*(*int64)(unsafe.Add(mBase, uint32(v9)+296)) = v17
	*(*int64)(unsafe.Add(mBase, uint32(v9)+160)) = int64(137438953472)
	*(*int32)(unsafe.Add(mBase, uint32(v9)+8)) = v9 + int32(24)
	*(*int32)(unsafe.Add(mBase, uint32(v9)+156)) = v9 + int32(168)
	goto L1
L1:
	;
	v29 = int32(0)
	v33 = F_raxSeek(m, v9, int32(_a67), v29, v29)
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
		goto L2
	} else {
		goto L3
	}
L2:
	;
	return
L3:
	;
	v35 = F_raxNext(m, v9)
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
		goto L2
	} else {
		goto L5
	}
L4:
	;
	F_raxStop(m, v9)
	mBase = m.M
	v63 = m.ExcPending
	if v63 != 0 {
		goto L2
	} else {
		goto L13
	}
L5:
	;
	if v35 == int32(0) {
		goto L4
	} else {
		goto L6
	}
L6:
	;
	goto L7
L7:
	;
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v9)+12))
	v46 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v9)+8))
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v9)+16))
	v50 = F_raxRemove(m, v46, v47, v48, int32(0))
	mBase = m.M
	v51 = m.ExcPending
	if v51 != 0 {
		goto L2
	} else {
		goto L9
	}
L8:
	;
	goto L4
L9:
	;
	F_valkey_free(m, v45)
	mBase = m.M
	v53 = m.ExcPending
	if v53 != 0 {
		goto L2
	} else {
		goto L10
	}
L10:
	;
	v54 = F_raxNext(m, v9)
	mBase = m.M
	v55 = m.ExcPending
	if v55 != 0 {
		goto L2
	} else {
		goto L11
	}
L11:
	;
	if v54 != 0 {
		goto L7
	} else {
		goto L12
	}
L12:
	;
	goto L8
L13:
	;
	v64 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v65 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	v68 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v65+int32(-1)))))
	switch v68 & int32(7) {
	case 0:
		goto L19
	case 1:
		goto L18
	case 2:
		goto L17
	case 3:
		goto L16
	case 4:
		goto L15
	default:
		v85 = v29
		goto L14
	}
L14:
	;
	v87 = F_raxRemove(m, v64, v65, v85, int32(0))
	mBase = m.M
	v88 = m.ExcPending
	if v88 != 0 {
		goto L2
	} else {
		goto L20
	}
L15:
	;
	v84 = *(*int32)(unsafe.Add(mBase, uint32(v65+int32(-17))))
	v85 = v84
	goto L14
L16:
	;
	v81 = *(*int32)(unsafe.Add(mBase, uint32(v65+int32(-9))))
	v85 = v81
	goto L14
L17:
	;
	v78 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v65+int32(-5)))))
	v85 = v78
	goto L14
L18:
	;
	v75 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v65+int32(-3)))))
	v85 = v75
	goto L14
L19:
	;
	v85 = int32(base.Ui32(v68) >> (uint(int32(3)) % 32))
	goto L14
L20:
	;
	v89 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	F_raxFree(m, v89)
	mBase = m.M
	v91 = m.ExcPending
	if v91 != 0 {
		goto L2
	} else {
		goto L21
	}
L21:
	;
	v92 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	F_sdsfree(m, v92)
	mBase = m.M
	v94 = m.ExcPending
	if v94 != 0 {
		goto L2
	} else {
		goto L22
	}
L22:
	;
	F_valkey_free(m, l1)
	mBase = m.M
	v96 = m.ExcPending
	if v96 != 0 {
		goto L2
	} else {
		goto L23
	}
L23:
	;
	m.G0 = v9 + int32(304)
	return
}
func F_streamEntryExists(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v16 int32
	_ = v16
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	var v31 int64
	_ = v31
	var v32 int64
	_ = v32
	var v35 int64
	_ = v35
	var v36 int64
	_ = v36
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	v3 = int32(0)
	v5 = m.G0
	v7 = v5 - int32(480)
	m.G0 = v7
	F_streamIteratorStart(m, v7+int32(32), l0, l1, l1, v3)
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		return int32(0)
	} else {
		v23 = F_streamIteratorGetID(m, v7+int32(32), v7+int32(16), v7+int32(8))
		mBase = m.M
		v24 = m.ExcPending
		if v24 != 0 {
			return int32(0)
		} else {
			F_raxStop(m, v7+int32(120))
			mBase = m.M
			v28 = m.ExcPending
			if v28 != 0 {
				return int32(0)
			} else {
				if v23 == int32(0) {
					v45 = v3
					m.G0 = v7 + int32(480)
					return v45
				} else {
					v31 = *(*int64)(unsafe.Add(mBase, uint32(l1)))
					v32 = *(*int64)(unsafe.Add(mBase, uint32(v7)+16))
					if v31 != v32 {
						F__serverAssert(m, int32(_a2405), int32(_a2402), int32(1324))
						mBase = m.M
						v43 = m.ExcPending
						if v43 != 0 {
							return int32(0)
						} else {
							F_abort(m)
							mBase = m.M
							base.Wasm_trap_unreachable()
							for {
							}
						}
					} else {
						v35 = *(*int64)(unsafe.Add(mBase, uint32(l1)+8))
						v36 = *(*int64)(unsafe.Add(mBase, uint32(v7)+24))
						if v35 == v36 {
							v45 = int32(1)
							m.G0 = v7 + int32(480)
							return v45
						} else {
							F__serverAssert(m, int32(_a2405), int32(_a2402), int32(1324))
							mBase = m.M
							v43 = m.ExcPending
							if v43 != 0 {
								return int32(0)
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
	}
}
func F_streamGetEdgeID(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v13 int64
	_ = v13
	var v17 int64
	_ = v17
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v29 int32
	_ = v29
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v65 int64
	_ = v65
	var v71 int64
	_ = v71
	var v78 int32
	_ = v78
	v7 = m.G0
	v9 = v7 - int32(464)
	m.G0 = v9
	v13 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v9+int32(80)))) = v13
	v17 = int64(-1)
	*(*int64)(unsafe.Add(mBase, uint32(v9+int32(96)))) = v17
	*(*int64)(unsafe.Add(mBase, uint32(v9)+72)) = v13
	*(*int64)(unsafe.Add(mBase, uint32(v9)+88)) = v17
	v24 = v9 + int32(104)
	v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v24)+4)) = v25
	*(*int32)(unsafe.Add(mBase, uint32(v24))) = int32(2)
	v29 = int32(128)
	*(*int32)(unsafe.Add(mBase, uint32(v24)+20)) = v29
	*(*int64)(unsafe.Add(mBase, uint32(v24)+12)) = v13
	*(*int64)(unsafe.Add(mBase, uint32(v24)+296)) = v13
	*(*int64)(unsafe.Add(mBase, uint32(v24)+160)) = int64(137438953472)
	*(*int32)(unsafe.Add(mBase, uint32(v24)+8)) = v9 + v29
	*(*int32)(unsafe.Add(mBase, uint32(v24)+156)) = v9 + int32(272)
	if l1 != 0 {
		v45 = int32(_a67)
	} else {
		v45 = int32(_a1933)
	}
	v46 = int32(0)
	v48 = F_raxSeek(m, v24, v45, v46, v46)
	mBase = m.M
	v49 = m.ExcPending
	if v49 != 0 {
		return
	} else {
		*(*int64)(unsafe.Add(mBase, uint32(v9)+408)) = int64(0)
		*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = l0
		*(*int32)(unsafe.Add(mBase, uint32(v9)+64)) = l2
		*(*int32)(unsafe.Add(mBase, uint32(v9)+60)) = base.B2i32(l1 == int32(0))
		v61 = F_streamIteratorGetID(m, v9+int32(16), l3, v9+int32(8))
		mBase = m.M
		v62 = m.ExcPending
		if v62 != 0 {
			return
		} else {
			if v61 != 0 {
			} else {
				if l1 == int32(0) {
					v71 = int64(0)
					*(*int64)(unsafe.Add(mBase, uint32(l3))) = v71
					*(*int64)(unsafe.Add(mBase, uint32(l3+int32(8)))) = v71
				} else {
					v65 = int64(-1)
					*(*int64)(unsafe.Add(mBase, uint32(l3))) = v65
					*(*int64)(unsafe.Add(mBase, uint32(l3+int32(8)))) = v65
				}
			}
			F_raxStop(m, v24)
			mBase = m.M
			v78 = m.ExcPending
			if v78 != 0 {
				return
			} else {
				m.G0 = v9 + int32(464)
				return
			}
		}
	}
}
func F_streamIteratorStart(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v14 int64
	_ = v14
	var v16 int64
	_ = v16
	var v17 int64
	_ = v17
	var v18 int64
	_ = v18
	var v19 int64
	_ = v19
	var v22 int64
	_ = v22
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v36 int64
	_ = v36
	var v50 int64
	_ = v50
	var v52 int64
	_ = v52
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v62 int64
	_ = v62
	var v74 int32
	_ = v74
	var v77 int64
	_ = v77
	var v78 int64
	_ = v78
	var v82 int64
	_ = v82
	var v84 int64
	_ = v84
	var v86 int64
	_ = v86
	var v89 int64
	_ = v89
	var v91 int64
	_ = v91
	var v93 int64
	_ = v93
	var v95 int64
	_ = v95
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v164 int32
	_ = v164
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v171 int32
	_ = v171
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v175 int64
	_ = v175
	var v176 int64
	_ = v176
	var v180 int64
	_ = v180
	var v182 int64
	_ = v182
	var v184 int64
	_ = v184
	var v187 int64
	_ = v187
	var v189 int64
	_ = v189
	var v191 int64
	_ = v191
	var v193 int64
	_ = v193
	var v254 int32
	_ = v254
	var v255 int32
	_ = v255
	var v256 int32
	_ = v256
	var v262 int32
	_ = v262
	var v264 int32
	_ = v264
	var v265 int32
	_ = v265
	var v268 int32
	_ = v268
	var v270 int32
	_ = v270
	var v272 int32
	_ = v272
	var v273 int32
	_ = v273
	v10 = m.G0
	v12 = v10 - int32(16)
	m.G0 = v12
	if l2 != 0 {
		v16 = *(*int64)(unsafe.Add(mBase, uint32(l2)+8))
		v17 = *(*int64)(unsafe.Add(mBase, uint32(l2)))
		v18 = v17
		v19 = v16
	} else {
		v14 = int64(0)
		v18 = v14
		v19 = v14
	}
	*(*int64)(unsafe.Add(mBase, uint32(l0)+64)) = v19
	*(*int64)(unsafe.Add(mBase, uint32(l0)+56)) = v18
	if l3 != 0 {
		v50 = *(*int64)(unsafe.Add(mBase, uint32(l3)))
		*(*int64)(unsafe.Add(mBase, uint32(l0)+72)) = v50
		v52 = *(*int64)(unsafe.Add(mBase, uint32(l3)+8))
		*(*int64)(unsafe.Add(mBase, uint32(l0)+80)) = v52
		v55 = l0 + int32(88)
		v56 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
		*(*int32)(unsafe.Add(mBase, uint32(v55)+4)) = v56
		*(*int32)(unsafe.Add(mBase, uint32(v55))) = int32(2)
		*(*int32)(unsafe.Add(mBase, uint32(v55)+20)) = int32(128)
		v62 = int64(0)
		*(*int64)(unsafe.Add(mBase, uint32(v55)+12)) = v62
		*(*int64)(unsafe.Add(mBase, uint32(v55)+296)) = v62
		*(*int64)(unsafe.Add(mBase, uint32(v55)+160)) = int64(137438953472)
		*(*int32)(unsafe.Add(mBase, uint32(v55)+8)) = l0 + int32(112)
		*(*int32)(unsafe.Add(mBase, uint32(v55)+156)) = l0 + int32(256)
		if l4 != 0 {
			v175 = *(*int64)(unsafe.Add(mBase, uint32(l3)))
			v176 = *(*int64)(unsafe.Add(mBase, uint32(l3)+8))
			if v175|v176 == int64(0) {
				v268 = v55
				v270 = int32(0)
				v272 = F_raxSeek(m, v268, int32(_a1933), v270, v270)
				mBase = m.M
				v273 = m.ExcPending
				if v273 != 0 {
					return
				} else {
					*(*int64)(unsafe.Add(mBase, uint32(l0)+392)) = int64(0)
					*(*int32)(unsafe.Add(mBase, uint32(l0))) = l1
					*(*int32)(unsafe.Add(mBase, uint32(l0)+48)) = int32(1)
					*(*int32)(unsafe.Add(mBase, uint32(l0)+44)) = l4
					m.G0 = v12 + int32(16)
					return
				}
			} else {
				v180 = int64(56)
				v182 = int64(65280)
				v184 = int64(40)
				v187 = int64(16711680)
				v189 = int64(24)
				v191 = int64(4278190080)
				v193 = int64(8)
				*(*int64)(unsafe.Add(mBase, uint32(v12)+8)) = v176<<(uint(v180)%64) | v176&v182<<(uint(v184)%64) | (v176&v187<<(uint(v189)%64) | v176&v191<<(uint(v193)%64)) | (int64(base.Ui64(v176)>>(uint(v193)%64))&v191 | int64(base.Ui64(v176)>>(uint(v189)%64))&v187 | (int64(base.Ui64(v176)>>(uint(v184)%64))&v182 | int64(base.Ui64(v176)>>(uint(v180)%64))))
				*(*int64)(unsafe.Add(mBase, uint32(v12))) = v175<<(uint(v180)%64) | v175&v182<<(uint(v184)%64) | (v175&v187<<(uint(v189)%64) | v175&v191<<(uint(v193)%64)) | (int64(base.Ui64(v175)>>(uint(v193)%64))&v191 | int64(base.Ui64(v175)>>(uint(v189)%64))&v187 | (int64(base.Ui64(v175)>>(uint(v184)%64))&v182 | int64(base.Ui64(v175)>>(uint(v180)%64))))
				v254 = F_raxSeek(m, v55, int32(_a2403), v12, int32(16))
				mBase = m.M
				v255 = m.ExcPending
				if v255 != 0 {
					return
				} else {
					v256 = *(*int32)(unsafe.Add(mBase, uint32(v55)))
					if v256&int32(2) == int32(0) {
						*(*int64)(unsafe.Add(mBase, uint32(l0)+392)) = int64(0)
						*(*int32)(unsafe.Add(mBase, uint32(l0))) = l1
						*(*int32)(unsafe.Add(mBase, uint32(l0)+48)) = int32(1)
						*(*int32)(unsafe.Add(mBase, uint32(l0)+44)) = l4
						m.G0 = v12 + int32(16)
						return
					} else {
						v262 = int32(0)
						v264 = F_raxSeek(m, v55, int32(_a1933), v262, v262)
						mBase = m.M
						v265 = m.ExcPending
						if v265 != 0 {
							return
						} else {
							*(*int64)(unsafe.Add(mBase, uint32(l0)+392)) = int64(0)
							*(*int32)(unsafe.Add(mBase, uint32(l0))) = l1
							*(*int32)(unsafe.Add(mBase, uint32(l0)+48)) = int32(1)
							*(*int32)(unsafe.Add(mBase, uint32(l0)+44)) = l4
							m.G0 = v12 + int32(16)
							return
						}
					}
				}
			}
		} else {
			v74 = v55
			if l2 == int32(0) {
				v171 = int32(0)
				v173 = F_raxSeek(m, v74, int32(_a67), v171, v171)
				mBase = m.M
				v174 = m.ExcPending
				if v174 != 0 {
					return
				} else {
					*(*int64)(unsafe.Add(mBase, uint32(l0)+392)) = int64(0)
					*(*int32)(unsafe.Add(mBase, uint32(l0))) = l1
					*(*int32)(unsafe.Add(mBase, uint32(l0)+48)) = int32(1)
					*(*int32)(unsafe.Add(mBase, uint32(l0)+44)) = l4
					m.G0 = v12 + int32(16)
					return
				}
			} else {
				v77 = *(*int64)(unsafe.Add(mBase, uint32(l2)))
				v78 = *(*int64)(unsafe.Add(mBase, uint32(l2)+8))
				if v77|v78 == int64(0) {
					v171 = int32(0)
					v173 = F_raxSeek(m, v74, int32(_a67), v171, v171)
					mBase = m.M
					v174 = m.ExcPending
					if v174 != 0 {
						return
					} else {
						*(*int64)(unsafe.Add(mBase, uint32(l0)+392)) = int64(0)
						*(*int32)(unsafe.Add(mBase, uint32(l0))) = l1
						*(*int32)(unsafe.Add(mBase, uint32(l0)+48)) = int32(1)
						*(*int32)(unsafe.Add(mBase, uint32(l0)+44)) = l4
						m.G0 = v12 + int32(16)
						return
					}
				} else {
					v82 = int64(56)
					v84 = int64(65280)
					v86 = int64(40)
					v89 = int64(16711680)
					v91 = int64(24)
					v93 = int64(4278190080)
					v95 = int64(8)
					*(*int64)(unsafe.Add(mBase, uint32(v12)+8)) = v78<<(uint(v82)%64) | v78&v84<<(uint(v86)%64) | (v78&v89<<(uint(v91)%64) | v78&v93<<(uint(v95)%64)) | (int64(base.Ui64(v78)>>(uint(v95)%64))&v93 | int64(base.Ui64(v78)>>(uint(v91)%64))&v89 | (int64(base.Ui64(v78)>>(uint(v86)%64))&v84 | int64(base.Ui64(v78)>>(uint(v82)%64))))
					*(*int64)(unsafe.Add(mBase, uint32(v12))) = v77<<(uint(v82)%64) | v77&v84<<(uint(v86)%64) | (v77&v89<<(uint(v91)%64) | v77&v93<<(uint(v95)%64)) | (int64(base.Ui64(v77)>>(uint(v95)%64))&v93 | int64(base.Ui64(v77)>>(uint(v91)%64))&v89 | (int64(base.Ui64(v77)>>(uint(v86)%64))&v84 | int64(base.Ui64(v77)>>(uint(v82)%64))))
					v156 = F_raxSeek(m, v74, int32(_a2403), v12, int32(16))
					mBase = m.M
					v157 = m.ExcPending
					if v157 != 0 {
						return
					} else {
						v158 = *(*int32)(unsafe.Add(mBase, uint32(v74)))
						if v158&int32(2) == int32(0) {
							*(*int64)(unsafe.Add(mBase, uint32(l0)+392)) = int64(0)
							*(*int32)(unsafe.Add(mBase, uint32(l0))) = l1
							*(*int32)(unsafe.Add(mBase, uint32(l0)+48)) = int32(1)
							*(*int32)(unsafe.Add(mBase, uint32(l0)+44)) = l4
							m.G0 = v12 + int32(16)
							return
						} else {
							v164 = int32(0)
							v166 = F_raxSeek(m, v74, int32(_a67), v164, v164)
							mBase = m.M
							v167 = m.ExcPending
							if v167 != 0 {
								return
							} else {
								*(*int64)(unsafe.Add(mBase, uint32(l0)+392)) = int64(0)
								*(*int32)(unsafe.Add(mBase, uint32(l0))) = l1
								*(*int32)(unsafe.Add(mBase, uint32(l0)+48)) = int32(1)
								*(*int32)(unsafe.Add(mBase, uint32(l0)+44)) = l4
								m.G0 = v12 + int32(16)
								return
							}
						}
					}
				}
			}
		}
	} else {
		v22 = int64(-1)
		*(*int64)(unsafe.Add(mBase, uint32(l0)+72)) = v22
		*(*int64)(unsafe.Add(mBase, uint32(l0+int32(80)))) = v22
		v29 = l0 + int32(88)
		v30 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
		*(*int32)(unsafe.Add(mBase, uint32(v29)+4)) = v30
		*(*int32)(unsafe.Add(mBase, uint32(v29))) = int32(2)
		*(*int32)(unsafe.Add(mBase, uint32(v29)+20)) = int32(128)
		v36 = int64(0)
		*(*int64)(unsafe.Add(mBase, uint32(v29)+12)) = v36
		*(*int64)(unsafe.Add(mBase, uint32(v29)+296)) = v36
		*(*int64)(unsafe.Add(mBase, uint32(v29)+160)) = int64(137438953472)
		*(*int32)(unsafe.Add(mBase, uint32(v29)+8)) = l0 + int32(112)
		*(*int32)(unsafe.Add(mBase, uint32(v29)+156)) = l0 + int32(256)
		if l4 == int32(0) {
			v74 = v29
			if l2 == int32(0) {
				v171 = int32(0)
				v173 = F_raxSeek(m, v74, int32(_a67), v171, v171)
				mBase = m.M
				v174 = m.ExcPending
				if v174 != 0 {
					return
				} else {
					*(*int64)(unsafe.Add(mBase, uint32(l0)+392)) = int64(0)
					*(*int32)(unsafe.Add(mBase, uint32(l0))) = l1
					*(*int32)(unsafe.Add(mBase, uint32(l0)+48)) = int32(1)
					*(*int32)(unsafe.Add(mBase, uint32(l0)+44)) = l4
					m.G0 = v12 + int32(16)
					return
				}
			} else {
				v77 = *(*int64)(unsafe.Add(mBase, uint32(l2)))
				v78 = *(*int64)(unsafe.Add(mBase, uint32(l2)+8))
				if v77|v78 == int64(0) {
					v171 = int32(0)
					v173 = F_raxSeek(m, v74, int32(_a67), v171, v171)
					mBase = m.M
					v174 = m.ExcPending
					if v174 != 0 {
						return
					} else {
						*(*int64)(unsafe.Add(mBase, uint32(l0)+392)) = int64(0)
						*(*int32)(unsafe.Add(mBase, uint32(l0))) = l1
						*(*int32)(unsafe.Add(mBase, uint32(l0)+48)) = int32(1)
						*(*int32)(unsafe.Add(mBase, uint32(l0)+44)) = l4
						m.G0 = v12 + int32(16)
						return
					}
				} else {
					v82 = int64(56)
					v84 = int64(65280)
					v86 = int64(40)
					v89 = int64(16711680)
					v91 = int64(24)
					v93 = int64(4278190080)
					v95 = int64(8)
					*(*int64)(unsafe.Add(mBase, uint32(v12)+8)) = v78<<(uint(v82)%64) | v78&v84<<(uint(v86)%64) | (v78&v89<<(uint(v91)%64) | v78&v93<<(uint(v95)%64)) | (int64(base.Ui64(v78)>>(uint(v95)%64))&v93 | int64(base.Ui64(v78)>>(uint(v91)%64))&v89 | (int64(base.Ui64(v78)>>(uint(v86)%64))&v84 | int64(base.Ui64(v78)>>(uint(v82)%64))))
					*(*int64)(unsafe.Add(mBase, uint32(v12))) = v77<<(uint(v82)%64) | v77&v84<<(uint(v86)%64) | (v77&v89<<(uint(v91)%64) | v77&v93<<(uint(v95)%64)) | (int64(base.Ui64(v77)>>(uint(v95)%64))&v93 | int64(base.Ui64(v77)>>(uint(v91)%64))&v89 | (int64(base.Ui64(v77)>>(uint(v86)%64))&v84 | int64(base.Ui64(v77)>>(uint(v82)%64))))
					v156 = F_raxSeek(m, v74, int32(_a2403), v12, int32(16))
					mBase = m.M
					v157 = m.ExcPending
					if v157 != 0 {
						return
					} else {
						v158 = *(*int32)(unsafe.Add(mBase, uint32(v74)))
						if v158&int32(2) == int32(0) {
							*(*int64)(unsafe.Add(mBase, uint32(l0)+392)) = int64(0)
							*(*int32)(unsafe.Add(mBase, uint32(l0))) = l1
							*(*int32)(unsafe.Add(mBase, uint32(l0)+48)) = int32(1)
							*(*int32)(unsafe.Add(mBase, uint32(l0)+44)) = l4
							m.G0 = v12 + int32(16)
							return
						} else {
							v164 = int32(0)
							v166 = F_raxSeek(m, v74, int32(_a67), v164, v164)
							mBase = m.M
							v167 = m.ExcPending
							if v167 != 0 {
								return
							} else {
								*(*int64)(unsafe.Add(mBase, uint32(l0)+392)) = int64(0)
								*(*int32)(unsafe.Add(mBase, uint32(l0))) = l1
								*(*int32)(unsafe.Add(mBase, uint32(l0)+48)) = int32(1)
								*(*int32)(unsafe.Add(mBase, uint32(l0)+44)) = l4
								m.G0 = v12 + int32(16)
								return
							}
						}
					}
				}
			}
		} else {
			v268 = v29
			v270 = int32(0)
			v272 = F_raxSeek(m, v268, int32(_a1933), v270, v270)
			mBase = m.M
			v273 = m.ExcPending
			if v273 != 0 {
				return
			} else {
				*(*int64)(unsafe.Add(mBase, uint32(l0)+392)) = int64(0)
				*(*int32)(unsafe.Add(mBase, uint32(l0))) = l1
				*(*int32)(unsafe.Add(mBase, uint32(l0)+48)) = int32(1)
				*(*int32)(unsafe.Add(mBase, uint32(l0)+44)) = l4
				m.G0 = v12 + int32(16)
				return
			}
		}
	}
}
func F_streamParseID(m *base.Module, l0 int32, l1 int32) int32 {
	var v3 int32
	_ = v3
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	v3 = int32(0)
	v7 = F_streamGenericParseIDOrReply(m, v3, l0, l1, int64(0), v3, v3)
	v10 = m.ExcPending
	if v10 != 0 {
		return int32(0)
	} else {
		return v7
	}
}
func F_streamPropagateXCLAIM(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v30 int32
	_ = v30
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
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v51 int64
	_ = v51
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v58 int64
	_ = v58
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v63 int64
	_ = v63
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v73 int64
	_ = v73
	var v74 int32
	_ = v74
	var v78 int32
	_ = v78
	var v86 int32
	_ = v86
	var v87 int64
	_ = v87
	var v108 int32
	_ = v108
	var v113 int32
	_ = v113
	var v120 int32
	_ = v120
	var v125 int32
	_ = v125
	var v127 int32
	_ = v127
	var v131 int64
	_ = v131
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v139 int32
	_ = v139
	var v142 int32
	_ = v142
	var v145 int32
	_ = v145
	var v152 int64
	_ = v152
	var v154 int32
	_ = v154
	var v158 int64
	_ = v158
	var v159 int64
	_ = v159
	var v168 int32
	_ = v168
	var v171 int32
	_ = v171
	var v178 int32
	_ = v178
	var v179 int64
	_ = v179
	var v181 int32
	_ = v181
	var v186 int32
	_ = v186
	var v195 int32
	_ = v195
	var v199 int32
	_ = v199
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v219 int64
	_ = v219
	var v232 int32
	_ = v232
	var v233 int64
	_ = v233
	var v254 int32
	_ = v254
	var v259 int32
	_ = v259
	var v266 int32
	_ = v266
	var v271 int32
	_ = v271
	var v273 int32
	_ = v273
	var v277 int64
	_ = v277
	var v281 int32
	_ = v281
	var v282 int32
	_ = v282
	var v285 int32
	_ = v285
	var v288 int32
	_ = v288
	var v291 int32
	_ = v291
	var v298 int64
	_ = v298
	var v300 int32
	_ = v300
	var v304 int64
	_ = v304
	var v305 int64
	_ = v305
	var v314 int32
	_ = v314
	var v317 int32
	_ = v317
	var v324 int32
	_ = v324
	var v325 int64
	_ = v325
	var v327 int32
	_ = v327
	var v332 int32
	_ = v332
	var v341 int32
	_ = v341
	var v345 int32
	_ = v345
	var v355 int32
	_ = v355
	var v360 int32
	_ = v360
	var v361 int32
	_ = v361
	var v362 int32
	_ = v362
	var v363 int32
	_ = v363
	var v365 int32
	_ = v365
	var v366 int32
	_ = v366
	var v369 int32
	_ = v369
	var v371 int32
	_ = v371
	var v372 int32
	_ = v372
	var v374 int32
	_ = v374
	var v375 int32
	_ = v375
	var v377 int32
	_ = v377
	var v378 int32
	_ = v378
	var v380 int32
	_ = v380
	var v381 int32
	_ = v381
	var v383 int32
	_ = v383
	v9 = m.G0
	v11 = v9 - int32(112)
	m.G0 = v11
	*(*int32)(unsafe.Add(mBase, uint32(v11)+8)) = l3
	*(*int32)(unsafe.Add(mBase, uint32(v11)+4)) = l1
	v16 = *(*int32)(unsafe.Add(mBase, _consts[866]))
	*(*int32)(unsafe.Add(mBase, uint32(v11))) = v16
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l5)+16))
	v20 = *(*int32)(unsafe.Add(mBase, uint32(v19)+16))
	v23 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20+int32(-1)))))
	switch v23 & int32(7) {
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
		v40 = int32(0)
		goto L1
	}
L1:
	;
	v41 = F_createStringObject_1(m, v20, v40)
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L7
	} else {
		goto L8
	}
L2:
	;
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v20+int32(-17))))
	v40 = v39
	goto L1
L3:
	;
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v20+int32(-9))))
	v40 = v36
	goto L1
L4:
	;
	v33 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v20+int32(-5)))))
	v40 = v33
	goto L1
L5:
	;
	v30 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20+int32(-3)))))
	v40 = v30
	goto L1
L6:
	;
	v40 = int32(base.Ui32(v23) >> (uint(int32(3)) % 32))
	goto L1
L7:
	;
	return
L8:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+20)) = l4
	*(*int32)(unsafe.Add(mBase, uint32(v11)+12)) = v41
	v45 = int32(_a388)
	v46 = *(*int32)(unsafe.Add(mBase, _consts[506]))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = v46
	v49 = *(*int32)(unsafe.Add(mBase, _consts[867]))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+24)) = v49
	v51 = *(*int64)(unsafe.Add(mBase, uint32(l5)))
	v52 = F_createStringObjectFromLongLong(m, v51)
	mBase = m.M
	v53 = m.ExcPending
	if v53 != 0 {
		goto L7
	} else {
		goto L9
	}
L9:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+28)) = v52
	v56 = *(*int32)(unsafe.Add(mBase, _consts[868]))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+32)) = v56
	v58 = *(*int64)(unsafe.Add(mBase, uint32(l5)+8))
	v59 = F_createStringObjectFromLongLong(m, v58)
	mBase = m.M
	v60 = m.ExcPending
	if v60 != 0 {
		goto L7
	} else {
		goto L10
	}
L10:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+36)) = v59
	v62 = int32(_a388)
	v63 = *(*int64)(unsafe.Add(mBase, _consts[869]))
	*(*int64)(unsafe.Add(mBase, uint32(v11)+40)) = v63
	v66 = *(*int32)(unsafe.Add(mBase, _consts[870]))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+48)) = v66
	v69 = v11 + int32(64)
	v73 = *(*int64)(unsafe.Add(mBase, uint32(l2)))
	v74 = int32(0)
	v78 = int32(1)
	if base.Ui64(v73) < base.Ui64(int64(10)) {
		v135 = v78
		v136 = v74
		goto L12
	} else {
		goto L13
	}
L11:
	;
	v210 = v69 + v209
	v211 = int32(45)
	*(*uint8)(unsafe.Add(mBase, uint32(v210))) = uint8(v211)
	v213 = int32(1)
	v214 = v210 + v213
	v215 = int32(0)
	v219 = *(*int64)(unsafe.Add(mBase, uint32(l2)+8))
	if base.Ui64(v219) < base.Ui64(int64(10)) {
		v281 = v213
		v282 = v215
		goto L56
	} else {
		goto L57
	}
L12:
	;
	v139 = v135 + v136
	if base.Ui32(int32(21)) <= base.Ui32(v139) {
		goto L43
	} else {
		goto L44
	}
L13:
	;
	v86 = v74
	v87 = v73
	goto L14
L14:
	;
	if base.Ui64(int64(100)) <= base.Ui64(v87) {
		goto L16
	} else {
		goto L17
	}
L15:
	;
	v135 = v78
	v136 = v127
	goto L12
L16:
	;
	if base.Ui64(int64(1000)) <= base.Ui64(v87) {
		goto L18
	} else {
		goto L19
	}
L17:
	;
	v135 = int32(2)
	v136 = v86
	goto L12
L18:
	;
	if base.Ui64(int64(999999999999)) < base.Ui64(v87) {
		goto L20
	} else {
		goto L21
	}
L19:
	;
	v135 = int32(3)
	v136 = v86
	goto L12
L20:
	;
	v127 = v86 + int32(12)
	v131 = base.I64_div_u_s(v87, int64(1000000000000))
	if base.Ui64(int64(9999999999999)) < base.Ui64(v87) {
		v86 = v127
		v87 = v131
		goto L14
	} else {
		goto L42
	}
L21:
	;
	if base.Ui64(int64(99999999)) < base.Ui64(v87) {
		goto L22
	} else {
		goto L23
	}
L22:
	;
	if base.Ui64(int64(9999999999)) < base.Ui64(v87) {
		goto L34
	} else {
		goto L35
	}
L23:
	;
	if base.Ui64(int64(999999)) < base.Ui64(v87) {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	if base.Ui64(int64(9999999)) < base.Ui64(v87) {
		goto L31
	} else {
		goto L32
	}
L25:
	;
	if base.Ui64(int64(10000)) <= base.Ui64(v87) {
		goto L26
	} else {
		goto L27
	}
L26:
	;
	if base.Ui64(int64(99999)) < base.Ui64(v87) {
		goto L28
	} else {
		goto L29
	}
L27:
	;
	v135 = int32(4)
	v136 = v86
	goto L12
L28:
	;
	v108 = int32(6)
	goto L30
L29:
	;
	v108 = int32(5)
	goto L30
L30:
	;
	v135 = v108
	v136 = v86
	goto L12
L31:
	;
	v113 = int32(8)
	goto L33
L32:
	;
	v113 = int32(7)
	goto L33
L33:
	;
	v135 = v113
	v136 = v86
	goto L12
L34:
	;
	if base.Ui64(int64(99999999999)) < base.Ui64(v87) {
		goto L39
	} else {
		goto L40
	}
L35:
	;
	if base.Ui64(int64(999999999)) < base.Ui64(v87) {
		goto L36
	} else {
		goto L37
	}
L36:
	;
	v120 = int32(10)
	goto L38
L37:
	;
	v120 = int32(9)
	goto L38
L38:
	;
	v135 = v120
	v136 = v86
	goto L12
L39:
	;
	v125 = int32(12)
	goto L41
L40:
	;
	v125 = int32(11)
	goto L41
L41:
	;
	v135 = v125
	v136 = v86
	goto L12
L42:
	;
	goto L15
L43:
	;
	goto L54
L44:
	;
	v142 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v69+v139))) = uint8(v142)
	v145 = v139 + int32(-1)
	if base.Ui64(int64(100)) <= base.Ui64(v73) {
		goto L46
	} else {
		goto L47
	}
L45:
	;
	v181 = v69 + v178
	if base.Ui64(int64(9)) < base.Ui64(v179) {
		goto L51
	} else {
		goto L52
	}
L46:
	;
	v152 = v73
	v154 = v145
	goto L48
L47:
	;
	v178 = v145
	v179 = v73
	goto L45
L48:
	;
	v158 = int64(100)
	v159 = base.I64_div_u_s(v152, v158)
	v168 = int32(*(*uint16)(unsafe.Add(mBase, uint32(base.I32_wrap_i64(v152-v159*v158)<<(uint(int32(1))%32))+uint32(_consts[871]))))
	*(*uint16)(unsafe.Add(mBase, uint32(v11+int32(63)+v154))) = uint16(v168)
	v171 = v154 + int32(-2)
	if base.Ui64(int64(9999)) < base.Ui64(v152) {
		v152 = v159
		v154 = v171
		goto L48
	} else {
		goto L50
	}
L49:
	;
	v178 = v171
	v179 = v159
	goto L45
L50:
	;
	goto L49
L51:
	;
	v195 = int32(*(*uint16)(unsafe.Add(mBase, uint32(base.I32_wrap_i64(v179)<<(uint(int32(1))%32))+uint32(_consts[871]))))
	*(*uint16)(unsafe.Add(mBase, uint32(v181+int32(-1)))) = uint16(v195)
	v209 = v139
	goto L11
L52:
	;
	v186 = base.I32_wrap_i64(v179) | int32(48)
	*(*uint8)(unsafe.Add(mBase, uint32(v181))) = uint8(v186)
	v209 = v139
	goto L11
L53:
	;
	v209 = int32(0)
	goto L11
L54:
	;
	v199 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v69))) = uint8(v199)
	goto L53
L55:
	;
	v360 = F_sdsnewlen(m, v11+int32(64), v214+v355-(v11+int32(64)))
	mBase = m.M
	v361 = m.ExcPending
	if v361 != 0 {
		goto L7
	} else {
		goto L99
	}
L56:
	;
	v285 = v281 + v282
	if base.Ui32(int32(21)) <= base.Ui32(v285) {
		goto L87
	} else {
		goto L88
	}
L57:
	;
	v232 = v215
	v233 = v219
	goto L58
L58:
	;
	if base.Ui64(int64(100)) <= base.Ui64(v233) {
		goto L60
	} else {
		goto L61
	}
L59:
	;
	v281 = v213
	v282 = v273
	goto L56
L60:
	;
	if base.Ui64(int64(1000)) <= base.Ui64(v233) {
		goto L62
	} else {
		goto L63
	}
L61:
	;
	v281 = int32(2)
	v282 = v232
	goto L56
L62:
	;
	if base.Ui64(int64(999999999999)) < base.Ui64(v233) {
		goto L64
	} else {
		goto L65
	}
L63:
	;
	v281 = int32(3)
	v282 = v232
	goto L56
L64:
	;
	v273 = v232 + int32(12)
	v277 = base.I64_div_u_s(v233, int64(1000000000000))
	if base.Ui64(int64(9999999999999)) < base.Ui64(v233) {
		v232 = v273
		v233 = v277
		goto L58
	} else {
		goto L86
	}
L65:
	;
	if base.Ui64(int64(99999999)) < base.Ui64(v233) {
		goto L66
	} else {
		goto L67
	}
L66:
	;
	if base.Ui64(int64(9999999999)) < base.Ui64(v233) {
		goto L78
	} else {
		goto L79
	}
L67:
	;
	if base.Ui64(int64(999999)) < base.Ui64(v233) {
		goto L68
	} else {
		goto L69
	}
L68:
	;
	if base.Ui64(int64(9999999)) < base.Ui64(v233) {
		goto L75
	} else {
		goto L76
	}
L69:
	;
	if base.Ui64(int64(10000)) <= base.Ui64(v233) {
		goto L70
	} else {
		goto L71
	}
L70:
	;
	if base.Ui64(int64(99999)) < base.Ui64(v233) {
		goto L72
	} else {
		goto L73
	}
L71:
	;
	v281 = int32(4)
	v282 = v232
	goto L56
L72:
	;
	v254 = int32(6)
	goto L74
L73:
	;
	v254 = int32(5)
	goto L74
L74:
	;
	v281 = v254
	v282 = v232
	goto L56
L75:
	;
	v259 = int32(8)
	goto L77
L76:
	;
	v259 = int32(7)
	goto L77
L77:
	;
	v281 = v259
	v282 = v232
	goto L56
L78:
	;
	if base.Ui64(int64(99999999999)) < base.Ui64(v233) {
		goto L83
	} else {
		goto L84
	}
L79:
	;
	if base.Ui64(int64(999999999)) < base.Ui64(v233) {
		goto L80
	} else {
		goto L81
	}
L80:
	;
	v266 = int32(10)
	goto L82
L81:
	;
	v266 = int32(9)
	goto L82
L82:
	;
	v281 = v266
	v282 = v232
	goto L56
L83:
	;
	v271 = int32(12)
	goto L85
L84:
	;
	v271 = int32(11)
	goto L85
L85:
	;
	v281 = v271
	v282 = v232
	goto L56
L86:
	;
	goto L59
L87:
	;
	goto L98
L88:
	;
	v288 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v214+v285))) = uint8(v288)
	v291 = v285 + int32(-1)
	if base.Ui64(int64(100)) <= base.Ui64(v219) {
		goto L90
	} else {
		goto L91
	}
L89:
	;
	v327 = v214 + v324
	if base.Ui64(int64(9)) < base.Ui64(v325) {
		goto L95
	} else {
		goto L96
	}
L90:
	;
	v298 = v219
	v300 = v291
	goto L92
L91:
	;
	v324 = v291
	v325 = v219
	goto L89
L92:
	;
	v304 = int64(100)
	v305 = base.I64_div_u_s(v298, v304)
	v314 = int32(*(*uint16)(unsafe.Add(mBase, uint32(base.I32_wrap_i64(v298-v305*v304)<<(uint(int32(1))%32))+uint32(_consts[871]))))
	*(*uint16)(unsafe.Add(mBase, uint32(v210+int32(0)+v300))) = uint16(v314)
	v317 = v300 + int32(-2)
	if base.Ui64(int64(9999)) < base.Ui64(v298) {
		v298 = v305
		v300 = v317
		goto L92
	} else {
		goto L94
	}
L93:
	;
	v324 = v317
	v325 = v305
	goto L89
L94:
	;
	goto L93
L95:
	;
	v341 = int32(*(*uint16)(unsafe.Add(mBase, uint32(base.I32_wrap_i64(v325)<<(uint(int32(1))%32))+uint32(_consts[871]))))
	*(*uint16)(unsafe.Add(mBase, uint32(v327+int32(-1)))) = uint16(v341)
	v355 = v285
	goto L55
L96:
	;
	v332 = base.I32_wrap_i64(v325) | int32(48)
	*(*uint8)(unsafe.Add(mBase, uint32(v327))) = uint8(v332)
	v355 = v285
	goto L55
L97:
	;
	v355 = int32(0)
	goto L55
L98:
	;
	v345 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v214))) = uint8(v345)
	goto L97
L99:
	;
	v362 = F_createObject(m, v215, v360)
	mBase = m.M
	v363 = m.ExcPending
	if v363 != 0 {
		goto L7
	} else {
		goto L100
	}
L100:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+52)) = v362
	v365 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v366 = *(*int32)(unsafe.Add(mBase, uint32(v365)+28))
	v369 = *(*int32)(unsafe.Add(mBase, uint32(l0)+292))
	F_alsoPropagate(m, v366, v11, int32(14), int32(3), v369)
	mBase = m.M
	v371 = m.ExcPending
	if v371 != 0 {
		goto L7
	} else {
		goto L101
	}
L101:
	;
	v372 = *(*int32)(unsafe.Add(mBase, uint32(v11)+12))
	F_decrRefCount(m, v372)
	mBase = m.M
	v374 = m.ExcPending
	if v374 != 0 {
		goto L7
	} else {
		goto L102
	}
L102:
	;
	v375 = *(*int32)(unsafe.Add(mBase, uint32(v11)+28))
	F_decrRefCount(m, v375)
	mBase = m.M
	v377 = m.ExcPending
	if v377 != 0 {
		goto L7
	} else {
		goto L103
	}
L103:
	;
	v378 = *(*int32)(unsafe.Add(mBase, uint32(v11)+36))
	F_decrRefCount(m, v378)
	mBase = m.M
	v380 = m.ExcPending
	if v380 != 0 {
		goto L7
	} else {
		goto L104
	}
L104:
	;
	v381 = *(*int32)(unsafe.Add(mBase, uint32(v11)+52))
	F_decrRefCount(m, v381)
	mBase = m.M
	v383 = m.ExcPending
	if v383 != 0 {
		goto L7
	} else {
		goto L105
	}
L105:
	;
	m.G0 = v11 + int32(112)
	return
}
func F_streamReplDataBufToDb(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v14 int32
	_ = v14
	var v19 int32
	_ = v19
	var v23 int32
	_ = v23
	var v24 int64
	_ = v24
	var v28 int64
	_ = v28
	var v32 int64
	_ = v32
	var v35 int32
	_ = v35
	var v43 int64
	_ = v43
	var v47 int64
	_ = v47
	var v49 int32
	_ = v49
	var v54 int64
	_ = v54
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v75 int64
	_ = v75
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v88 int32
	_ = v88
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v98 int64
	_ = v98
	var v100 int64
	_ = v100
	var v105 int64
	_ = v105
	var v107 int64
	_ = v107
	var v111 int64
	_ = v111
	var v113 int32
	_ = v113
	var v114 int64
	_ = v114
	var v116 int64
	_ = v116
	var v118 int32
	_ = v118
	var v120 int64
	_ = v120
	var v123 int32
	_ = v123
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v134 int32
	_ = v134
	var v135 int64
	_ = v135
	var v138 int32
	_ = v138
	var v147 int32
	_ = v147
	var v150 int32
	_ = v150
	var v152 int32
	_ = v152
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v167 int32
	_ = v167
	v9 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+200)))
	if v9&int32(1) == int32(0) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	F__serverAssert(m, int32(_a1962), int32(_a1913), int32(3362))
	mBase = m.M
	v167 = m.ExcPending
	if v167 != 0 {
		goto L11
	} else {
		goto L31
	}
L2:
	;
	v14 = int32(0)
	v19 = *(*int32)(unsafe.Add(mBase, _consts[416]))
	*(*int32)(unsafe.Add(mBase, _consts[416])) = v19 + int32(1)
	if v19 != 0 {
		goto L4
	} else {
		goto L5
	}
L3:
	;
	v47 = F_mstime(m)
	mBase = m.M
	v49 = *(*int32)(unsafe.Add(mBase, _consts[606]))
	if v49 == int32(0) {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	goto L3
L5:
	;
	v23 = int32(0)
	v24 = F_ustime(m)
	mBase = m.M
	*(*int64)(unsafe.Add(mBase, _consts[178])) = v24
	v28 = base.I64_div_s(v24, int64(1000))
	*(*int64)(unsafe.Add(mBase, _consts[54])) = v28
	v32 = base.I64_div_s(v24, int64(1000000))
	*(*int64)(unsafe.Add(mBase, _consts[109])) = v32
	v35 = *(*int32)(unsafe.Add(mBase, _consts[179]))
	F_lrulfu_updateClockAndPolicy(m, v28, int32(base.Ui32(v35&int32(2))>>(uint(int32(1))%32)))
	mBase = m.M
	v43 = *(*int64)(unsafe.Add(mBase, _consts[54]))
	*(*int64)(unsafe.Add(mBase, _consts[417])) = v43
	goto L4
L6:
	;
	v147 = int32(0)
	v150 = *(*int32)(unsafe.Add(mBase, _consts[416]))
	v152 = v150 + int32(-1)
	*(*int32)(unsafe.Add(mBase, _consts[416])) = v152
	if v152 != 0 {
		goto L26
	} else {
		goto L27
	}
L7:
	;
	v54 = v47
	v55 = v49
	v56 = int32(0)
	goto L8
L8:
	;
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v55)))
	if v61 == int32(0) {
		goto L6
	} else {
		goto L10
	}
L9:
	;
	goto L6
L10:
	;
	v64 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v61)+8))
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v65)+4))
	v69 = F_sdscatlen(m, v64, v65+int32(8), v68)
	mBase = m.M
	v72 = m.ExcPending
	if v72 != 0 {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	return int32(0)
L12:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v69
	v74 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
	v75 = *(*int64)(unsafe.Add(mBase, uint32(v74)+40))
	*(*int64)(unsafe.Add(mBase, uint32(v74)+40)) = v75 + base.I64_extend_i32_u(v68)
	v79 = F_processInputBuffer(m, l0)
	mBase = m.M
	v80 = m.ExcPending
	if v80 != 0 {
		goto L11
	} else {
		goto L13
	}
L13:
	;
	v81 = int32(_a20)
	v83 = *(*int32)(unsafe.Add(mBase, _consts[607]))
	*(*int32)(unsafe.Add(mBase, _consts[607])) = v83 - v68
	v88 = *(*int32)(unsafe.Add(mBase, _consts[608]))
	*(*int32)(unsafe.Add(mBase, _consts[608])) = v88 - v68 + int32(-20)
	v94 = *(*int32)(unsafe.Add(mBase, _consts[606]))
	F_listDelNode(m, v94, v61)
	mBase = m.M
	v96 = m.ExcPending
	if v96 != 0 {
		goto L11
	} else {
		goto L14
	}
L14:
	;
	v97 = v68 + v56
	v98 = F_mstime(m)
	mBase = m.M
	v100 = *(*int64)(unsafe.Add(mBase, _consts[609]))
	if v100 == int64(0) {
		v135 = v54
		goto L15
	} else {
		goto L16
	}
L15:
	;
	v138 = *(*int32)(unsafe.Add(mBase, _consts[606]))
	if v138 != 0 {
		v54 = v135
		v55 = v138
		v56 = v97
		goto L8
	} else {
		goto L24
	}
L16:
	;
	v105 = base.I64_div_s(base.I64_extend_i32_u(v97+v68), v100)
	v107 = base.I64_div_s(base.I64_extend_i32_u(v97), v100)
	if v105 <= v107 {
		v135 = v54
		goto L15
	} else {
		goto L17
	}
L17:
	;
	v111 = *(*int64)(unsafe.Add(mBase, _consts[610]))
	if v98-v54 <= v111 {
		v135 = v54
		goto L15
	} else {
		goto L18
	}
L18:
	;
	v113 = int32(0)
	v114 = F___time(m, v113)
	mBase = m.M
	v116 = *(*int64)(unsafe.Add(mBase, _consts[598]))
	if v114 == v116 {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	F_processEventsWhileBlocked(m)
	mBase = m.M
	v134 = m.ExcPending
	if v134 != 0 {
		goto L11
	} else {
		goto L23
	}
L20:
	;
	v118 = int32(0)
	v120 = F___time(m, v118)
	mBase = m.M
	*(*int64)(unsafe.Add(mBase, _consts[598])) = v120
	v123 = *(*int32)(unsafe.Add(mBase, _consts[592]))
	if v123 == v118 {
		goto L19
	} else {
		goto L21
	}
L21:
	;
	v128 = *(*int32)(unsafe.Add(mBase, uint32(v123)))
	v129 = *(*int32)(unsafe.Add(mBase, uint32(v128)+68))
	v130 = m.T0[v129].(func(*base.Module, int32, int32, int32) int32)(m, v123, int32(_a26), int32(1))
	mBase = m.M
	v131 = m.ExcPending
	if v131 != 0 {
		goto L11
	} else {
		goto L22
	}
L22:
	;
	goto L19
L23:
	;
	v135 = v98
	goto L15
L24:
	;
	goto L9
L25:
	;
	v160 = *(*int32)(unsafe.Add(mBase, _consts[606]))
	if v160 != 0 {
		goto L28
	} else {
		goto L29
	}
L26:
	;
	goto L25
L27:
	;
	*(*int64)(unsafe.Add(mBase, _consts[417])) = int64(0)
	goto L26
L28:
	;
	v161 = int32(0)
	goto L30
L29:
	;
	v161 = int32(-1)
	goto L30
L30:
	;
	return v161
L31:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_streamReplyWithRange(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32, l9 int32) int32 {
	mBase := m.M
	_ = mBase
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v53 int32
	_ = v53
	var v63 int32
	_ = v63
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v78 int64
	_ = v78
	var v81 int64
	_ = v81
	var v84 int64
	_ = v84
	var v85 int64
	_ = v85
	var v87 int64
	_ = v87
	var v90 int64
	_ = v90
	var v93 int64
	_ = v93
	var v94 int64
	_ = v94
	var v96 int64
	_ = v96
	var v99 int64
	_ = v99
	var v102 int64
	_ = v102
	var v107 int64
	_ = v107
	var v108 int64
	_ = v108
	var v114 int64
	_ = v114
	var v118 int32
	_ = v118
	var v126 int64
	_ = v126
	var v132 int64
	_ = v132
	var v133 int64
	_ = v133
	var v136 int64
	_ = v136
	var v139 int64
	_ = v139
	var v140 int64
	_ = v140
	var v145 int64
	_ = v145
	var v146 int64
	_ = v146
	var v149 int64
	_ = v149
	var v151 int64
	_ = v151
	var v153 int64
	_ = v153
	var v155 int64
	_ = v155
	var v156 int64
	_ = v156
	var v157 int64
	_ = v157
	var v159 int64
	_ = v159
	var v160 int64
	_ = v160
	var v162 int64
	_ = v162
	var v165 int64
	_ = v165
	var v168 int64
	_ = v168
	var v169 int64
	_ = v169
	var v172 int64
	_ = v172
	var v173 int64
	_ = v173
	var v179 int32
	_ = v179
	var v181 int64
	_ = v181
	var v186 int64
	_ = v186
	var v187 int64
	_ = v187
	var v192 int32
	_ = v192
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v197 int64
	_ = v197
	var v201 int64
	_ = v201
	var v202 int64
	_ = v202
	var v208 int32
	_ = v208
	var v211 int32
	_ = v211
	var v214 int64
	_ = v214
	var v217 int32
	_ = v217
	var v220 int32
	_ = v220
	var v221 int64
	_ = v221
	var v233 int64
	_ = v233
	var v243 int64
	_ = v243
	var v244 int64
	_ = v244
	var v249 int64
	_ = v249
	var v257 int64
	_ = v257
	var v260 int32
	_ = v260
	var v266 int32
	_ = v266
	var v268 int32
	_ = v268
	var v272 int32
	_ = v272
	var v276 int32
	_ = v276
	var v284 int32
	_ = v284
	var v285 int64
	_ = v285
	var v306 int32
	_ = v306
	var v311 int32
	_ = v311
	var v318 int32
	_ = v318
	var v323 int32
	_ = v323
	var v325 int32
	_ = v325
	var v329 int64
	_ = v329
	var v333 int32
	_ = v333
	var v334 int32
	_ = v334
	var v337 int32
	_ = v337
	var v340 int32
	_ = v340
	var v343 int32
	_ = v343
	var v350 int64
	_ = v350
	var v352 int32
	_ = v352
	var v356 int64
	_ = v356
	var v357 int64
	_ = v357
	var v366 int32
	_ = v366
	var v369 int32
	_ = v369
	var v376 int32
	_ = v376
	var v377 int64
	_ = v377
	var v379 int32
	_ = v379
	var v384 int32
	_ = v384
	var v393 int32
	_ = v393
	var v397 int32
	_ = v397
	var v407 int32
	_ = v407
	var v408 int32
	_ = v408
	var v409 int32
	_ = v409
	var v411 int32
	_ = v411
	var v412 int32
	_ = v412
	var v416 int64
	_ = v416
	var v417 int32
	_ = v417
	var v429 int32
	_ = v429
	var v430 int64
	_ = v430
	var v451 int32
	_ = v451
	var v456 int32
	_ = v456
	var v463 int32
	_ = v463
	var v468 int32
	_ = v468
	var v470 int32
	_ = v470
	var v474 int64
	_ = v474
	var v478 int32
	_ = v478
	var v479 int32
	_ = v479
	var v482 int32
	_ = v482
	var v485 int32
	_ = v485
	var v488 int32
	_ = v488
	var v495 int64
	_ = v495
	var v497 int32
	_ = v497
	var v501 int64
	_ = v501
	var v502 int64
	_ = v502
	var v511 int32
	_ = v511
	var v514 int32
	_ = v514
	var v521 int32
	_ = v521
	var v522 int64
	_ = v522
	var v524 int32
	_ = v524
	var v529 int32
	_ = v529
	var v538 int32
	_ = v538
	var v542 int32
	_ = v542
	var v552 int32
	_ = v552
	var v558 int32
	_ = v558
	var v559 int64
	_ = v559
	var v564 int32
	_ = v564
	var v582 int64
	_ = v582
	var v594 int32
	_ = v594
	var v595 int32
	_ = v595
	var v596 int32
	_ = v596
	var v598 int32
	_ = v598
	var v599 int32
	_ = v599
	var v600 int32
	_ = v600
	var v602 int32
	_ = v602
	var v604 int64
	_ = v604
	var v627 int64
	_ = v627
	var v629 int64
	_ = v629
	var v631 int64
	_ = v631
	var v634 int64
	_ = v634
	var v636 int64
	_ = v636
	var v638 int64
	_ = v638
	var v640 int64
	_ = v640
	var v700 int32
	_ = v700
	var v701 int32
	_ = v701
	var v703 int64
	_ = v703
	var v708 int32
	_ = v708
	var v711 int32
	_ = v711
	var v712 int32
	_ = v712
	var v713 int32
	_ = v713
	var v716 int32
	_ = v716
	var v717 int32
	_ = v717
	var v719 int32
	_ = v719
	var v720 int32
	_ = v720
	var v721 int32
	_ = v721
	var v723 int32
	_ = v723
	var v732 int32
	_ = v732
	var v733 int32
	_ = v733
	var v744 int32
	_ = v744
	var v745 int32
	_ = v745
	var v747 int32
	_ = v747
	var v753 int32
	_ = v753
	var v754 int32
	_ = v754
	var v755 int32
	_ = v755
	var v760 int32
	_ = v760
	var v770 int32
	_ = v770
	var v773 int32
	_ = v773
	var v776 int32
	_ = v776
	var v778 int32
	_ = v778
	var v780 int32
	_ = v780
	var v781 int32
	_ = v781
	var v783 int32
	_ = v783
	var v793 int32
	_ = v793
	var v796 int32
	_ = v796
	var v801 int32
	_ = v801
	var v804 int32
	_ = v804
	var v817 int32
	_ = v817
	var v819 int32
	_ = v819
	var v821 int32
	_ = v821
	var v826 int32
	_ = v826
	var v833 int32
	_ = v833
	var v838 int32
	_ = v838
	var v848 int32
	_ = v848
	var v849 int32
	_ = v849
	var v857 int32
	_ = v857
	var v858 int32
	_ = v858
	var v860 int32
	_ = v860
	var v863 int32
	_ = v863
	var v867 int32
	_ = v867
	var v873 int32
	_ = v873
	var v874 int32
	_ = v874
	var v876 int32
	_ = v876
	var v880 int32
	_ = v880
	var v882 int32
	_ = v882
	var v888 int32
	_ = v888
	var v892 int32
	_ = v892
	var v898 int32
	_ = v898
	var v899 int32
	_ = v899
	var v909 int32
	_ = v909
	var v913 int32
	_ = v913
	var v914 int32
	_ = v914
	var v917 int32
	_ = v917
	var v921 int32
	_ = v921
	var v922 int32
	_ = v922
	var v923 int32
	_ = v923
	var v926 int32
	_ = v926
	var v927 int32
	_ = v927
	var v930 int64
	_ = v930
	var v934 int32
	_ = v934
	var v937 int32
	_ = v937
	var v938 int32
	_ = v938
	var v943 int32
	_ = v943
	var v946 int64
	_ = v946
	var v951 int32
	_ = v951
	var v955 int32
	_ = v955
	var v959 int32
	_ = v959
	var v967 int32
	_ = v967
	var v968 int64
	_ = v968
	var v989 int32
	_ = v989
	var v994 int32
	_ = v994
	var v1001 int32
	_ = v1001
	var v1006 int32
	_ = v1006
	var v1008 int32
	_ = v1008
	var v1012 int64
	_ = v1012
	var v1016 int32
	_ = v1016
	var v1017 int32
	_ = v1017
	var v1020 int32
	_ = v1020
	var v1023 int32
	_ = v1023
	var v1026 int32
	_ = v1026
	var v1033 int64
	_ = v1033
	var v1035 int32
	_ = v1035
	var v1039 int64
	_ = v1039
	var v1040 int64
	_ = v1040
	var v1049 int32
	_ = v1049
	var v1052 int32
	_ = v1052
	var v1059 int32
	_ = v1059
	var v1060 int64
	_ = v1060
	var v1062 int32
	_ = v1062
	var v1067 int32
	_ = v1067
	var v1076 int32
	_ = v1076
	var v1080 int32
	_ = v1080
	var v1090 int32
	_ = v1090
	var v1091 int32
	_ = v1091
	var v1092 int32
	_ = v1092
	var v1094 int32
	_ = v1094
	var v1095 int32
	_ = v1095
	var v1096 int32
	_ = v1096
	var v1112 int32
	_ = v1112
	var v1113 int64
	_ = v1113
	var v1134 int32
	_ = v1134
	var v1139 int32
	_ = v1139
	var v1146 int32
	_ = v1146
	var v1151 int32
	_ = v1151
	var v1153 int32
	_ = v1153
	var v1157 int64
	_ = v1157
	var v1161 int32
	_ = v1161
	var v1162 int32
	_ = v1162
	var v1165 int32
	_ = v1165
	var v1168 int32
	_ = v1168
	var v1171 int32
	_ = v1171
	var v1178 int64
	_ = v1178
	var v1180 int32
	_ = v1180
	var v1184 int64
	_ = v1184
	var v1185 int64
	_ = v1185
	var v1194 int32
	_ = v1194
	var v1197 int32
	_ = v1197
	var v1204 int32
	_ = v1204
	var v1205 int64
	_ = v1205
	var v1207 int32
	_ = v1207
	var v1212 int32
	_ = v1212
	var v1221 int32
	_ = v1221
	var v1225 int32
	_ = v1225
	var v1235 int32
	_ = v1235
	var v1240 int32
	_ = v1240
	var v1241 int32
	_ = v1241
	var v1242 int32
	_ = v1242
	var v1243 int32
	_ = v1243
	var v1244 int32
	_ = v1244
	var v1245 int32
	_ = v1245
	var v1247 int32
	_ = v1247
	var v1249 int32
	_ = v1249
	var v1255 int32
	_ = v1255
	var v1262 int32
	_ = v1262
	var v1272 int32
	_ = v1272
	var v1281 int32
	_ = v1281
	var v1282 int32
	_ = v1282
	var v1284 int32
	_ = v1284
	var v1288 int32
	_ = v1288
	var v1292 int32
	_ = v1292
	var v1296 int32
	_ = v1296
	var v1319 int32
	_ = v1319
	var v1326 int32
	_ = v1326
	v19 = m.G0
	v21 = v19 - int32(608)
	m.G0 = v21
	if l6 == int32(0) {
		goto L4
	} else {
		goto L5
	}
L1:
	;
	F__serverPanic_1(m, int32(_a2402), int32(1773), int32(_a2406), int32(0))
	mBase = m.M
	v1326 = m.ExcPending
	if v1326 != 0 {
		goto L7
	} else {
		goto L348
	}
L2:
	;
	F__serverAssert(m, int32(_a1778), int32(_a2402), int32(1763))
	mBase = m.M
	v1319 = m.ExcPending
	if v1319 != 0 {
		goto L7
	} else {
		goto L347
	}
L3:
	;
	m.G0 = v21 + int32(608)
	return v1296
L4:
	;
	if l8&int32(2) != 0 {
		v40 = int32(0)
		goto L9
	} else {
		goto L10
	}
L5:
	;
	if l8&int32(4) == int32(0) {
		goto L4
	} else {
		goto L6
	}
L6:
	;
	v29 = F_streamReplyWithRangeFromConsumerPEL(m, l0, l1, l2, l3, l4, l7)
	mBase = m.M
	v32 = m.ExcPending
	if v32 != 0 {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	return int32(0)
L8:
	;
	v1296 = v29
	goto L3
L9:
	;
	F_streamIteratorStart(m, v21+int32(112), l1, l2, l3, l5)
	mBase = m.M
	v44 = m.ExcPending
	if v44 != 0 {
		goto L7
	} else {
		goto L12
	}
L10:
	;
	v38 = F_addReplyDeferredLen(m, l0)
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L7
	} else {
		goto L11
	}
L11:
	;
	v40 = v38
	goto L9
L12:
	;
	v45 = int32(0)
	v53 = v45
	v63 = v45
	goto L14
L13:
	;
	if l9 == int32(0) {
		goto L340
	} else {
		goto L341
	}
L14:
	;
	v74 = F_streamIteratorGetID(m, v21+int32(112), v21+int32(88), v21+int32(104))
	mBase = m.M
	v75 = m.ExcPending
	if v75 != 0 {
		goto L7
	} else {
		goto L16
	}
L15:
	;
	v1262 = l4
	v1272 = v260
	goto L13
L16:
	;
	if v74 == int32(0) {
		v1262 = v53
		v1272 = v63
		goto L13
	} else {
		goto L17
	}
L17:
	;
	v78 = *(*int64)(unsafe.Add(mBase, uint32(v21)+88))
	if l6 == int32(0) {
		v260 = v63
		goto L18
	} else {
		goto L19
	}
L18:
	;
	F_addReplyArrayLen(m, l0, int32(2))
	mBase = m.M
	v266 = m.ExcPending
	if v266 != 0 {
		goto L7
	} else {
		goto L91
	}
L19:
	;
	v81 = *(*int64)(unsafe.Add(mBase, uint32(l6)))
	if base.Ui64(v81) < base.Ui64(v78) {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	v87 = *(*int64)(unsafe.Add(mBase, uint32(l6)+16))
	if v87 == int64(-1) {
		goto L26
	} else {
		goto L27
	}
L21:
	;
	if base.Ui64(v78) < base.Ui64(v81) {
		v260 = v63
		goto L18
	} else {
		goto L22
	}
L22:
	;
	v84 = *(*int64)(unsafe.Add(mBase, uint32(v21)+96))
	v85 = *(*int64)(unsafe.Add(mBase, uint32(l6)+8))
	if base.Ui64(v84) <= base.Ui64(v85) {
		v260 = v63
		goto L18
	} else {
		goto L23
	}
L23:
	;
	goto L20
L24:
	;
	v249 = *(*int64)(unsafe.Add(mBase, uint32(v21)+88))
	*(*int64)(unsafe.Add(mBase, uint32(l6))) = v249
	v257 = *(*int64)(unsafe.Add(mBase, uint32(v21+int32(96))))
	*(*int64)(unsafe.Add(mBase, uint32(l6+int32(8)))) = v257
	v260 = int32(1)
	goto L18
L25:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l6)+16)) = v244
	goto L24
L26:
	;
	v114 = *(*int64)(unsafe.Add(mBase, uint32(l1)+64))
	if v114 == int64(0) {
		goto L24
	} else {
		goto L40
	}
L27:
	;
	v90 = *(*int64)(unsafe.Add(mBase, uint32(l1)+32))
	if base.Ui64(v90) < base.Ui64(v81) {
		goto L28
	} else {
		goto L29
	}
L28:
	;
	v96 = *(*int64)(unsafe.Add(mBase, uint32(l1)+8))
	if v96 == int64(0) {
		goto L32
	} else {
		goto L33
	}
L29:
	;
	if base.Ui64(v81) < base.Ui64(v90) {
		goto L26
	} else {
		goto L30
	}
L30:
	;
	v93 = *(*int64)(unsafe.Add(mBase, uint32(l6)+8))
	v94 = *(*int64)(unsafe.Add(mBase, uint32(l1)+40))
	if base.Ui64(v93) < base.Ui64(v94) {
		goto L26
	} else {
		goto L31
	}
L31:
	;
	goto L28
L32:
	;
	v244 = v87 + int64(1)
	goto L25
L33:
	;
	v99 = *(*int64)(unsafe.Add(mBase, uint32(l1)+48))
	if v99 != int64(0) {
		goto L34
	} else {
		goto L35
	}
L34:
	;
	if base.Ui64(v99) < base.Ui64(v81) {
		goto L32
	} else {
		goto L37
	}
L35:
	;
	v102 = *(*int64)(unsafe.Add(mBase, uint32(l1)+56))
	if v102 == int64(0) {
		goto L32
	} else {
		goto L36
	}
L36:
	;
	goto L34
L37:
	;
	if base.Ui64(v81) < base.Ui64(v99) {
		goto L26
	} else {
		goto L38
	}
L38:
	;
	v107 = *(*int64)(unsafe.Add(mBase, uint32(l6)+8))
	v108 = *(*int64)(unsafe.Add(mBase, uint32(l1)+56))
	if base.Ui64(v107) <= base.Ui64(v108) {
		goto L26
	} else {
		goto L39
	}
L39:
	;
	goto L32
L40:
	;
	v118 = v21 + int32(88)
	v126 = *(*int64)(unsafe.Add(mBase, uint32(l1)+64))
	if base.B2i32(v126 == int64(0)) == int32(0) {
		goto L42
	} else {
		goto L43
	}
L41:
	;
	v244 = v243
	goto L25
L42:
	;
	v132 = *(*int64)(unsafe.Add(mBase, uint32(v118)))
	v133 = *(*int64)(unsafe.Add(mBase, uint32(l1)+8))
	if v133 != int64(0) {
		goto L49
	} else {
		goto L50
	}
L43:
	;
	v243 = int64(0)
	goto L41
L44:
	;
	v243 = v233
	goto L41
L45:
	;
	if base.Ui64(v168) < base.Ui64(v169) {
		goto L62
	} else {
		goto L63
	}
L46:
	;
	v165 = *(*int64)(unsafe.Add(mBase, uint32(l1)+16))
	if base.Ui64(v162) <= base.Ui64(v165) {
		v168 = v162
		v169 = v165
		goto L45
	} else {
		goto L61
	}
L47:
	;
	v157 = int64(-1)
	if base.Ui64(v155) < base.Ui64(v156) {
		v233 = v157
		goto L44
	} else {
		goto L59
	}
L48:
	;
	v153 = *(*int64)(unsafe.Add(mBase, uint32(l1)+48))
	if base.Ui64(v153) < base.Ui64(v132) {
		v162 = v132
		goto L46
	} else {
		goto L58
	}
L49:
	;
	if v132 != int64(0) {
		goto L48
	} else {
		goto L55
	}
L50:
	;
	v136 = *(*int64)(unsafe.Add(mBase, uint32(l1)+16))
	if base.Ui64(v136) < base.Ui64(v132) {
		goto L48
	} else {
		goto L51
	}
L51:
	;
	if base.Ui64(v136) <= base.Ui64(v132) {
		goto L52
	} else {
		goto L53
	}
L52:
	;
	v139 = *(*int64)(unsafe.Add(mBase, uint32(v118)+8))
	v140 = *(*int64)(unsafe.Add(mBase, uint32(l1)+24))
	if base.Ui64(v140) < base.Ui64(v139) {
		goto L49
	} else {
		goto L54
	}
L53:
	;
	v243 = v126
	goto L41
L54:
	;
	v243 = v126
	goto L41
L55:
	;
	v145 = int64(0)
	v146 = *(*int64)(unsafe.Add(mBase, uint32(v118)+8))
	if v146 != v145 {
		goto L56
	} else {
		goto L57
	}
L56:
	;
	v151 = *(*int64)(unsafe.Add(mBase, uint32(l1)+48))
	v155 = v145
	v156 = v151
	goto L47
L57:
	;
	v149 = *(*int64)(unsafe.Add(mBase, uint32(l1)+16))
	v168 = int64(0)
	v169 = v149
	goto L45
L58:
	;
	v155 = v132
	v156 = v153
	goto L47
L59:
	;
	v159 = *(*int64)(unsafe.Add(mBase, uint32(v118)+8))
	v160 = *(*int64)(unsafe.Add(mBase, uint32(l1)+56))
	if base.Ui64(v159) < base.Ui64(v160) {
		v233 = v157
		goto L44
	} else {
		goto L60
	}
L60:
	;
	v162 = v155
	goto L46
L61:
	;
	v243 = int64(-1)
	goto L41
L62:
	;
	v179 = int32(1)
	v181 = *(*int64)(unsafe.Add(mBase, uint32(l1)+32))
	if base.Ui64(v181) < base.Ui64(v168) {
		v195 = v179
		goto L67
	} else {
		goto L68
	}
L63:
	;
	v172 = *(*int64)(unsafe.Add(mBase, uint32(v118)+8))
	v173 = *(*int64)(unsafe.Add(mBase, uint32(l1)+24))
	if base.Ui64(v172) <= base.Ui64(v173) {
		goto L64
	} else {
		goto L65
	}
L64:
	;
	if base.Ui64(v172) < base.Ui64(v173) {
		goto L62
	} else {
		goto L66
	}
L65:
	;
	v243 = int64(-1)
	goto L41
L66:
	;
	v243 = v126
	goto L41
L67:
	;
	v196 = int32(0)
	v197 = *(*int64)(unsafe.Add(mBase, uint32(l1)+48))
	if base.Ui64(v181) < base.Ui64(v197) {
		v217 = v196
		v220 = v179
		goto L74
	} else {
		goto L75
	}
L68:
	;
	if base.Ui64(v168) < base.Ui64(v181) {
		v195 = int32(-1)
		goto L67
	} else {
		goto L69
	}
L69:
	;
	v186 = *(*int64)(unsafe.Add(mBase, uint32(v118)+8))
	v187 = *(*int64)(unsafe.Add(mBase, uint32(l1)+40))
	if base.Ui64(v187) < base.Ui64(v186) {
		v195 = int32(1)
		goto L67
	} else {
		goto L70
	}
L70:
	;
	if base.Ui64(v186) < base.Ui64(v187) {
		goto L71
	} else {
		goto L72
	}
L71:
	;
	v192 = int32(-1)
	goto L73
L72:
	;
	v192 = int32(0)
	goto L73
L73:
	;
	v195 = v192
	goto L67
L74:
	;
	v221 = int64(-1)
	if v217 != 0 {
		goto L85
	} else {
		goto L86
	}
L75:
	;
	if base.Ui64(v181) <= base.Ui64(v197) {
		goto L77
	} else {
		goto L78
	}
L76:
	;
	if v197 != int64(0) {
		v217 = v196
		v220 = v211
		goto L74
	} else {
		goto L84
	}
L77:
	;
	v201 = *(*int64)(unsafe.Add(mBase, uint32(l1)+56))
	v202 = *(*int64)(unsafe.Add(mBase, uint32(l1)+40))
	if base.Ui64(v201) <= base.Ui64(v202) {
		goto L79
	} else {
		goto L80
	}
L78:
	;
	v211 = int32(-1)
	goto L76
L79:
	;
	if base.Ui64(v201) < base.Ui64(v202) {
		goto L81
	} else {
		goto L82
	}
L80:
	;
	v211 = int32(1)
	goto L76
L81:
	;
	v208 = int32(-1)
	goto L83
L82:
	;
	v208 = int32(0)
	goto L83
L83:
	;
	v211 = v208
	goto L76
L84:
	;
	v214 = *(*int64)(unsafe.Add(mBase, uint32(l1)+56))
	v217 = base.B2i32(v214 == int64(0))
	v220 = v211
	goto L74
L85:
	;
	if int32(-1) < v195 {
		goto L88
	} else {
		goto L89
	}
L86:
	;
	if int32(-1) < v220 {
		v233 = v221
		goto L44
	} else {
		goto L87
	}
L87:
	;
	goto L85
L88:
	;
	if v195 != 0 {
		v233 = v221
		goto L44
	} else {
		goto L90
	}
L89:
	;
	v243 = v126 - v133
	goto L41
L90:
	;
	v233 = v126 - v133 + int64(1)
	goto L44
L91:
	;
	v268 = v21 + int32(32)
	v272 = int32(0)
	v276 = int32(1)
	if base.Ui64(v78) < base.Ui64(int64(10)) {
		v333 = v276
		v334 = v272
		goto L93
	} else {
		goto L94
	}
L92:
	;
	v408 = v268 + v407
	v409 = int32(45)
	*(*uint8)(unsafe.Add(mBase, uint32(v408))) = uint8(v409)
	v411 = int32(1)
	v412 = v408 + v411
	v416 = *(*int64)(unsafe.Add(mBase, uint32(v21)+96))
	v417 = int32(0)
	if base.Ui64(v416) < base.Ui64(int64(10)) {
		v478 = v411
		v479 = v417
		goto L137
	} else {
		goto L138
	}
L93:
	;
	v337 = v333 + v334
	if base.Ui32(int32(21)) <= base.Ui32(v337) {
		goto L124
	} else {
		goto L125
	}
L94:
	;
	v284 = v272
	v285 = v78
	goto L95
L95:
	;
	if base.Ui64(int64(100)) <= base.Ui64(v285) {
		goto L97
	} else {
		goto L98
	}
L96:
	;
	v333 = v276
	v334 = v325
	goto L93
L97:
	;
	if base.Ui64(int64(1000)) <= base.Ui64(v285) {
		goto L99
	} else {
		goto L100
	}
L98:
	;
	v333 = int32(2)
	v334 = v284
	goto L93
L99:
	;
	if base.Ui64(int64(999999999999)) < base.Ui64(v285) {
		goto L101
	} else {
		goto L102
	}
L100:
	;
	v333 = int32(3)
	v334 = v284
	goto L93
L101:
	;
	v325 = v284 + int32(12)
	v329 = base.I64_div_u_s(v285, int64(1000000000000))
	if base.Ui64(int64(9999999999999)) < base.Ui64(v285) {
		v284 = v325
		v285 = v329
		goto L95
	} else {
		goto L123
	}
L102:
	;
	if base.Ui64(int64(99999999)) < base.Ui64(v285) {
		goto L103
	} else {
		goto L104
	}
L103:
	;
	if base.Ui64(int64(9999999999)) < base.Ui64(v285) {
		goto L115
	} else {
		goto L116
	}
L104:
	;
	if base.Ui64(int64(999999)) < base.Ui64(v285) {
		goto L105
	} else {
		goto L106
	}
L105:
	;
	if base.Ui64(int64(9999999)) < base.Ui64(v285) {
		goto L112
	} else {
		goto L113
	}
L106:
	;
	if base.Ui64(int64(10000)) <= base.Ui64(v285) {
		goto L107
	} else {
		goto L108
	}
L107:
	;
	if base.Ui64(int64(99999)) < base.Ui64(v285) {
		goto L109
	} else {
		goto L110
	}
L108:
	;
	v333 = int32(4)
	v334 = v284
	goto L93
L109:
	;
	v306 = int32(6)
	goto L111
L110:
	;
	v306 = int32(5)
	goto L111
L111:
	;
	v333 = v306
	v334 = v284
	goto L93
L112:
	;
	v311 = int32(8)
	goto L114
L113:
	;
	v311 = int32(7)
	goto L114
L114:
	;
	v333 = v311
	v334 = v284
	goto L93
L115:
	;
	if base.Ui64(int64(99999999999)) < base.Ui64(v285) {
		goto L120
	} else {
		goto L121
	}
L116:
	;
	if base.Ui64(int64(999999999)) < base.Ui64(v285) {
		goto L117
	} else {
		goto L118
	}
L117:
	;
	v318 = int32(10)
	goto L119
L118:
	;
	v318 = int32(9)
	goto L119
L119:
	;
	v333 = v318
	v334 = v284
	goto L93
L120:
	;
	v323 = int32(12)
	goto L122
L121:
	;
	v323 = int32(11)
	goto L122
L122:
	;
	v333 = v323
	v334 = v284
	goto L93
L123:
	;
	goto L96
L124:
	;
	goto L135
L125:
	;
	v340 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v268+v337))) = uint8(v340)
	v343 = v337 + int32(-1)
	if base.Ui64(int64(100)) <= base.Ui64(v78) {
		goto L127
	} else {
		goto L128
	}
L126:
	;
	v379 = v268 + v376
	if base.Ui64(int64(9)) < base.Ui64(v377) {
		goto L132
	} else {
		goto L133
	}
L127:
	;
	v350 = v78
	v352 = v343
	goto L129
L128:
	;
	v376 = v343
	v377 = v78
	goto L126
L129:
	;
	v356 = int64(100)
	v357 = base.I64_div_u_s(v350, v356)
	v366 = int32(*(*uint16)(unsafe.Add(mBase, uint32(base.I32_wrap_i64(v350-v357*v356)<<(uint(int32(1))%32))+uint32(_consts[871]))))
	*(*uint16)(unsafe.Add(mBase, uint32(v21+int32(31)+v352))) = uint16(v366)
	v369 = v352 + int32(-2)
	if base.Ui64(int64(9999)) < base.Ui64(v350) {
		v350 = v357
		v352 = v369
		goto L129
	} else {
		goto L131
	}
L130:
	;
	v376 = v369
	v377 = v357
	goto L126
L131:
	;
	goto L130
L132:
	;
	v393 = int32(*(*uint16)(unsafe.Add(mBase, uint32(base.I32_wrap_i64(v377)<<(uint(int32(1))%32))+uint32(_consts[871]))))
	*(*uint16)(unsafe.Add(mBase, uint32(v379+int32(-1)))) = uint16(v393)
	v407 = v337
	goto L92
L133:
	;
	v384 = base.I32_wrap_i64(v377) | int32(48)
	*(*uint8)(unsafe.Add(mBase, uint32(v379))) = uint8(v384)
	v407 = v337
	goto L92
L134:
	;
	v407 = int32(0)
	goto L92
L135:
	;
	v397 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v268))) = uint8(v397)
	goto L134
L136:
	;
	F_addReplyBulkCBuffer(m, l0, v21+int32(32), v412+v552-(v21+int32(32)))
	mBase = m.M
	v558 = m.ExcPending
	if v558 != 0 {
		goto L7
	} else {
		goto L180
	}
L137:
	;
	v482 = v478 + v479
	if base.Ui32(int32(21)) <= base.Ui32(v482) {
		goto L168
	} else {
		goto L169
	}
L138:
	;
	v429 = v417
	v430 = v416
	goto L139
L139:
	;
	if base.Ui64(int64(100)) <= base.Ui64(v430) {
		goto L141
	} else {
		goto L142
	}
L140:
	;
	v478 = v411
	v479 = v470
	goto L137
L141:
	;
	if base.Ui64(int64(1000)) <= base.Ui64(v430) {
		goto L143
	} else {
		goto L144
	}
L142:
	;
	v478 = int32(2)
	v479 = v429
	goto L137
L143:
	;
	if base.Ui64(int64(999999999999)) < base.Ui64(v430) {
		goto L145
	} else {
		goto L146
	}
L144:
	;
	v478 = int32(3)
	v479 = v429
	goto L137
L145:
	;
	v470 = v429 + int32(12)
	v474 = base.I64_div_u_s(v430, int64(1000000000000))
	if base.Ui64(int64(9999999999999)) < base.Ui64(v430) {
		v429 = v470
		v430 = v474
		goto L139
	} else {
		goto L167
	}
L146:
	;
	if base.Ui64(int64(99999999)) < base.Ui64(v430) {
		goto L147
	} else {
		goto L148
	}
L147:
	;
	if base.Ui64(int64(9999999999)) < base.Ui64(v430) {
		goto L159
	} else {
		goto L160
	}
L148:
	;
	if base.Ui64(int64(999999)) < base.Ui64(v430) {
		goto L149
	} else {
		goto L150
	}
L149:
	;
	if base.Ui64(int64(9999999)) < base.Ui64(v430) {
		goto L156
	} else {
		goto L157
	}
L150:
	;
	if base.Ui64(int64(10000)) <= base.Ui64(v430) {
		goto L151
	} else {
		goto L152
	}
L151:
	;
	if base.Ui64(int64(99999)) < base.Ui64(v430) {
		goto L153
	} else {
		goto L154
	}
L152:
	;
	v478 = int32(4)
	v479 = v429
	goto L137
L153:
	;
	v451 = int32(6)
	goto L155
L154:
	;
	v451 = int32(5)
	goto L155
L155:
	;
	v478 = v451
	v479 = v429
	goto L137
L156:
	;
	v456 = int32(8)
	goto L158
L157:
	;
	v456 = int32(7)
	goto L158
L158:
	;
	v478 = v456
	v479 = v429
	goto L137
L159:
	;
	if base.Ui64(int64(99999999999)) < base.Ui64(v430) {
		goto L164
	} else {
		goto L165
	}
L160:
	;
	if base.Ui64(int64(999999999)) < base.Ui64(v430) {
		goto L161
	} else {
		goto L162
	}
L161:
	;
	v463 = int32(10)
	goto L163
L162:
	;
	v463 = int32(9)
	goto L163
L163:
	;
	v478 = v463
	v479 = v429
	goto L137
L164:
	;
	v468 = int32(12)
	goto L166
L165:
	;
	v468 = int32(11)
	goto L166
L166:
	;
	v478 = v468
	v479 = v429
	goto L137
L167:
	;
	goto L140
L168:
	;
	goto L179
L169:
	;
	v485 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v412+v482))) = uint8(v485)
	v488 = v482 + int32(-1)
	if base.Ui64(int64(100)) <= base.Ui64(v416) {
		goto L171
	} else {
		goto L172
	}
L170:
	;
	v524 = v412 + v521
	if base.Ui64(int64(9)) < base.Ui64(v522) {
		goto L176
	} else {
		goto L177
	}
L171:
	;
	v495 = v416
	v497 = v488
	goto L173
L172:
	;
	v521 = v488
	v522 = v416
	goto L170
L173:
	;
	v501 = int64(100)
	v502 = base.I64_div_u_s(v495, v501)
	v511 = int32(*(*uint16)(unsafe.Add(mBase, uint32(base.I32_wrap_i64(v495-v502*v501)<<(uint(int32(1))%32))+uint32(_consts[871]))))
	*(*uint16)(unsafe.Add(mBase, uint32(v408+int32(0)+v497))) = uint16(v511)
	v514 = v497 + int32(-2)
	if base.Ui64(int64(9999)) < base.Ui64(v495) {
		v495 = v502
		v497 = v514
		goto L173
	} else {
		goto L175
	}
L174:
	;
	v521 = v514
	v522 = v502
	goto L170
L175:
	;
	goto L174
L176:
	;
	v538 = int32(*(*uint16)(unsafe.Add(mBase, uint32(base.I32_wrap_i64(v522)<<(uint(int32(1))%32))+uint32(_consts[871]))))
	*(*uint16)(unsafe.Add(mBase, uint32(v524+int32(-1)))) = uint16(v538)
	v552 = v482
	goto L136
L177:
	;
	v529 = base.I32_wrap_i64(v522) | int32(48)
	*(*uint8)(unsafe.Add(mBase, uint32(v524))) = uint8(v529)
	v552 = v482
	goto L136
L178:
	;
	v552 = int32(0)
	goto L136
L179:
	;
	v542 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v412))) = uint8(v542)
	goto L178
L180:
	;
	v559 = *(*int64)(unsafe.Add(mBase, uint32(v21)+104))
	F_addReplyArrayLen(m, l0, base.I32_wrap_i64(v559)<<(uint(int32(1))%32))
	mBase = m.M
	v564 = m.ExcPending
	if v564 != 0 {
		goto L7
	} else {
		goto L181
	}
L181:
	;
	if v559 == int64(0) {
		goto L182
	} else {
		goto L183
	}
L182:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v21)+104)) = int64(-1)
	if base.B2i32(l6 == v45)|l8&int32(1) != 0 {
		goto L190
	} else {
		goto L191
	}
L183:
	;
	v582 = v559
	goto L184
L184:
	;
	F_streamIteratorGetField(m, v21+int32(112), v21+int32(28), v21+int32(24), v21+int32(560), v21)
	mBase = m.M
	v594 = m.ExcPending
	if v594 != 0 {
		goto L7
	} else {
		goto L186
	}
L185:
	;
	goto L182
L186:
	;
	v595 = *(*int32)(unsafe.Add(mBase, uint32(v21)+28))
	v596 = *(*int32)(unsafe.Add(mBase, uint32(v21)+560))
	F_addReplyBulkCBuffer(m, l0, v595, v596)
	mBase = m.M
	v598 = m.ExcPending
	if v598 != 0 {
		goto L7
	} else {
		goto L187
	}
L187:
	;
	v599 = *(*int32)(unsafe.Add(mBase, uint32(v21)+24))
	v600 = *(*int32)(unsafe.Add(mBase, uint32(v21)))
	F_addReplyBulkCBuffer(m, l0, v599, v600)
	mBase = m.M
	v602 = m.ExcPending
	if v602 != 0 {
		goto L7
	} else {
		goto L188
	}
L188:
	;
	v604 = v582 + int64(-1)
	if v604 != int64(0) {
		v582 = v604
		goto L184
	} else {
		goto L189
	}
L189:
	;
	goto L185
L190:
	;
	v1255 = v53 + int32(1)
	if l4 == int32(0) {
		v53 = v1255
		v63 = v260
		goto L14
	} else {
		goto L338
	}
L191:
	;
	v627 = int64(56)
	v629 = int64(65280)
	v631 = int64(40)
	v634 = int64(16711680)
	v636 = int64(24)
	v638 = int64(4278190080)
	v640 = int64(8)
	*(*int64)(unsafe.Add(mBase, uint32(v21)+8)) = v416<<(uint(v627)%64) | v416&v629<<(uint(v631)%64) | (v416&v634<<(uint(v636)%64) | v416&v638<<(uint(v640)%64)) | (int64(base.Ui64(v416)>>(uint(v640)%64))&v638 | int64(base.Ui64(v416)>>(uint(v636)%64))&v634 | (int64(base.Ui64(v416)>>(uint(v631)%64))&v629 | int64(base.Ui64(v416)>>(uint(v627)%64))))
	*(*int64)(unsafe.Add(mBase, uint32(v21))) = v78<<(uint(v627)%64) | v78&v629<<(uint(v631)%64) | (v78&v634<<(uint(v636)%64) | v78&v638<<(uint(v640)%64)) | (int64(base.Ui64(v78)>>(uint(v640)%64))&v638 | int64(base.Ui64(v78)>>(uint(v636)%64))&v634 | (int64(base.Ui64(v78)>>(uint(v631)%64))&v629 | int64(base.Ui64(v78)>>(uint(v627)%64))))
	v700 = F_valkey_malloc(m, int32(24))
	mBase = m.M
	v701 = m.ExcPending
	if v701 != 0 {
		goto L7
	} else {
		goto L192
	}
L192:
	;
	v703 = *(*int64)(unsafe.Add(mBase, _consts[12]))
	goto L193
L193:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v700)+16)) = l7
	*(*int64)(unsafe.Add(mBase, uint32(v700)+8)) = int64(1)
	*(*int64)(unsafe.Add(mBase, uint32(v700))) = v703
	v708 = *(*int32)(unsafe.Add(mBase, uint32(l6)+24))
	v711 = F_raxTryInsert(m, v708, v21, int32(16), v700, int32(0))
	mBase = m.M
	v712 = m.ExcPending
	if v712 != 0 {
		goto L7
	} else {
		goto L194
	}
L194:
	;
	v713 = *(*int32)(unsafe.Add(mBase, uint32(l7)+20))
	v716 = F_raxTryInsert(m, v713, v21, int32(16), v700, int32(0))
	mBase = m.M
	v717 = m.ExcPending
	if v717 != 0 {
		goto L7
	} else {
		goto L195
	}
L195:
	;
	if v711 != 0 {
		goto L197
	} else {
		goto L198
	}
L196:
	;
	v946 = *(*int64)(unsafe.Add(mBase, _consts[12]))
	goto L244
L197:
	;
	if v711 != int32(1) {
		v943 = v700
		goto L196
	} else {
		goto L242
	}
L198:
	;
	F_valkey_free(m, v700)
	mBase = m.M
	v719 = m.ExcPending
	if v719 != 0 {
		goto L7
	} else {
		goto L199
	}
L199:
	;
	v720 = *(*int32)(unsafe.Add(mBase, uint32(l6)+24))
	v721 = int32(16)
	v723 = v21 + int32(560)
	v732 = *(*int32)(unsafe.Add(mBase, uint32(v720)))
	v733 = *(*int32)(unsafe.Add(mBase, uint32(v732)))
	goto L203
L200:
	;
	if v917 == int32(0) {
		goto L2
	} else {
		goto L238
	}
L201:
	;
	if v876 != v721 {
		v917 = int32(0)
		goto L228
	} else {
		goto L229
	}
L202:
	;
	v867 = int32(0)
	v873 = v732
	v874 = v733
	v876 = v867
	v880 = v867
	goto L201
L203:
	;
	if base.Ui32(v733) < base.Ui32(int32(8)) {
		goto L202
	} else {
		goto L204
	}
L204:
	;
	v744 = v732
	v745 = v733
	v747 = int32(0)
	goto L206
L205:
	;
	v873 = v857
	v874 = v858
	v876 = v860
	v880 = base.B2i32(v863 != int32(0))
	goto L201
L206:
	;
	v753 = int32(base.Ui32(v745) >> (uint(int32(3)) % 32))
	v754 = int32(4)
	v755 = v744 + v754
	if v745&v754 == int32(0) {
		goto L209
	} else {
		goto L210
	}
L207:
	;
	v857 = v848
	v858 = v849
	v860 = v833
	v863 = v838
	goto L205
L208:
	;
	v838 = int32(0)
	v848 = *(*int32)(unsafe.Add(mBase, uint32(v755+v753+(v838-v753)&int32(3)+v826<<(uint(int32(2))%32))))
	v849 = *(*int32)(unsafe.Add(mBase, uint32(v848)))
	if base.Ui32(v849) < base.Ui32(int32(8)) {
		v857 = v848
		v858 = v849
		v860 = v833
		v863 = v838
		goto L205
	} else {
		goto L226
	}
L209:
	;
	v801 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21+v747))))
	v804 = int32(0)
	goto L220
L210:
	;
	v760 = int32(0)
	if base.Ui32(v721) <= base.Ui32(v747) {
		v793 = v747
		v796 = v760
		goto L211
	} else {
		goto L212
	}
L211:
	;
	if v796 == v753 {
		v826 = v760
		v833 = v793
		goto L208
	} else {
		goto L218
	}
L212:
	;
	v770 = v747
	v773 = v760
	goto L213
L213:
	;
	v776 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v755+v773))))
	v778 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21+v770))))
	if v776 != v778 {
		v793 = v770
		v796 = v773
		goto L211
	} else {
		goto L215
	}
L214:
	;
	v793 = v781
	v796 = v783
	goto L211
L215:
	;
	v780 = int32(1)
	v781 = v770 + v780
	v783 = v773 + v780
	if base.Ui32(v753) <= base.Ui32(v783) {
		v793 = v781
		v796 = v783
		goto L211
	} else {
		goto L216
	}
L216:
	;
	if base.Ui32(v781) < base.Ui32(v721) {
		v770 = v781
		v773 = v783
		goto L213
	} else {
		goto L217
	}
L217:
	;
	goto L214
L218:
	;
	v857 = v744
	v858 = v745
	v860 = v793
	v863 = v796
	goto L205
L219:
	;
	if v804 != v753 {
		goto L224
	} else {
		goto L225
	}
L220:
	;
	v817 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v755+v804))))
	if v817 == v801&int32(255) {
		goto L219
	} else {
		goto L222
	}
L222:
	;
	v819 = int32(1)
	v821 = v804 + v819
	if v821 != v753 {
		v804 = v821
		goto L220
	} else {
		goto L223
	}
L223:
	;
	v873 = v744
	v874 = v745
	v876 = v747
	v880 = v819
	goto L201
L224:
	;
	v826 = v804
	v833 = v747 + int32(1)
	goto L208
L225:
	;
	v857 = v744
	v858 = v745
	v860 = v747
	v863 = v753
	goto L205
L226:
	;
	if base.Ui32(v833) < base.Ui32(v721) {
		v744 = v848
		v745 = v849
		v747 = v833
		goto L206
	} else {
		goto L227
	}
L227:
	;
	goto L207
L228:
	;
	goto L200
L229:
	;
	v882 = int32(0)
	if v874&int32(1) == v882 {
		v917 = v882
		goto L228
	} else {
		goto L230
	}
L230:
	;
	v888 = v874 & int32(4)
	if v880&base.B2i32(v888 != int32(0)) != 0 {
		v917 = v882
		goto L228
	} else {
		goto L231
	}
L231:
	;
	v892 = int32(1)
	if v723 == int32(0) {
		v917 = v892
		goto L228
	} else {
		goto L232
	}
L232:
	;
	if v874&int32(2) != 0 {
		v914 = int32(0)
		goto L233
	} else {
		goto L234
	}
L233:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v723))) = v914
	v917 = v892
	goto L228
L234:
	;
	v898 = int32(3)
	v899 = int32(base.Ui32(v874) >> (uint(v898) % 32))
	if v888 != 0 {
		goto L235
	} else {
		goto L236
	}
L235:
	;
	v909 = int32(4)
	goto L237
L236:
	;
	v909 = v899 << (uint(int32(2)) % 32)
	goto L237
L237:
	;
	v913 = *(*int32)(unsafe.Add(mBase, uint32(v873+v899+(int32(0)-v899)&v898+v909+int32(4))))
	v914 = v913
	goto L233
L238:
	;
	v921 = *(*int32)(unsafe.Add(mBase, uint32(v21)+560))
	v922 = *(*int32)(unsafe.Add(mBase, uint32(v921)+16))
	v923 = *(*int32)(unsafe.Add(mBase, uint32(v922)+20))
	v926 = F_raxRemove(m, v923, v21, int32(16), int32(0))
	mBase = m.M
	v927 = m.ExcPending
	if v927 != 0 {
		goto L7
	} else {
		goto L239
	}
L239:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v921)+16)) = l7
	v930 = *(*int64)(unsafe.Add(mBase, _consts[12]))
	goto L240
L240:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v921)+8)) = int64(1)
	*(*int64)(unsafe.Add(mBase, uint32(v921))) = v930
	v934 = *(*int32)(unsafe.Add(mBase, uint32(l7)+20))
	v937 = F_raxInsert(m, v934, v21, int32(16), v921, int32(0))
	mBase = m.M
	v938 = m.ExcPending
	if v938 != 0 {
		goto L7
	} else {
		goto L241
	}
L241:
	;
	v943 = v921
	goto L196
L242:
	;
	if v716 == int32(0) {
		goto L1
	} else {
		goto L243
	}
L243:
	;
	v943 = v700
	goto L196
L244:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l7)+8)) = v946
	if l9 == int32(0) {
		goto L190
	} else {
		goto L245
	}
L245:
	;
	v951 = v21 + int32(560)
	v955 = int32(0)
	v959 = int32(1)
	if base.Ui64(v78) < base.Ui64(int64(10)) {
		v1016 = v959
		v1017 = v955
		goto L247
	} else {
		goto L248
	}
L246:
	;
	v1091 = v951 + v1090
	v1092 = int32(45)
	*(*uint8)(unsafe.Add(mBase, uint32(v1091))) = uint8(v1092)
	v1094 = int32(1)
	v1095 = v1091 + v1094
	v1096 = int32(0)
	if base.Ui64(v416) < base.Ui64(int64(10)) {
		v1161 = v1094
		v1162 = v1096
		goto L291
	} else {
		goto L292
	}
L247:
	;
	v1020 = v1016 + v1017
	if base.Ui32(int32(21)) <= base.Ui32(v1020) {
		goto L278
	} else {
		goto L279
	}
L248:
	;
	v967 = v955
	v968 = v78
	goto L249
L249:
	;
	if base.Ui64(int64(100)) <= base.Ui64(v968) {
		goto L251
	} else {
		goto L252
	}
L250:
	;
	v1016 = v959
	v1017 = v1008
	goto L247
L251:
	;
	if base.Ui64(int64(1000)) <= base.Ui64(v968) {
		goto L253
	} else {
		goto L254
	}
L252:
	;
	v1016 = int32(2)
	v1017 = v967
	goto L247
L253:
	;
	if base.Ui64(int64(999999999999)) < base.Ui64(v968) {
		goto L255
	} else {
		goto L256
	}
L254:
	;
	v1016 = int32(3)
	v1017 = v967
	goto L247
L255:
	;
	v1008 = v967 + int32(12)
	v1012 = base.I64_div_u_s(v968, int64(1000000000000))
	if base.Ui64(int64(9999999999999)) < base.Ui64(v968) {
		v967 = v1008
		v968 = v1012
		goto L249
	} else {
		goto L277
	}
L256:
	;
	if base.Ui64(int64(99999999)) < base.Ui64(v968) {
		goto L257
	} else {
		goto L258
	}
L257:
	;
	if base.Ui64(int64(9999999999)) < base.Ui64(v968) {
		goto L269
	} else {
		goto L270
	}
L258:
	;
	if base.Ui64(int64(999999)) < base.Ui64(v968) {
		goto L259
	} else {
		goto L260
	}
L259:
	;
	if base.Ui64(int64(9999999)) < base.Ui64(v968) {
		goto L266
	} else {
		goto L267
	}
L260:
	;
	if base.Ui64(int64(10000)) <= base.Ui64(v968) {
		goto L261
	} else {
		goto L262
	}
L261:
	;
	if base.Ui64(int64(99999)) < base.Ui64(v968) {
		goto L263
	} else {
		goto L264
	}
L262:
	;
	v1016 = int32(4)
	v1017 = v967
	goto L247
L263:
	;
	v989 = int32(6)
	goto L265
L264:
	;
	v989 = int32(5)
	goto L265
L265:
	;
	v1016 = v989
	v1017 = v967
	goto L247
L266:
	;
	v994 = int32(8)
	goto L268
L267:
	;
	v994 = int32(7)
	goto L268
L268:
	;
	v1016 = v994
	v1017 = v967
	goto L247
L269:
	;
	if base.Ui64(int64(99999999999)) < base.Ui64(v968) {
		goto L274
	} else {
		goto L275
	}
L270:
	;
	if base.Ui64(int64(999999999)) < base.Ui64(v968) {
		goto L271
	} else {
		goto L272
	}
L271:
	;
	v1001 = int32(10)
	goto L273
L272:
	;
	v1001 = int32(9)
	goto L273
L273:
	;
	v1016 = v1001
	v1017 = v967
	goto L247
L274:
	;
	v1006 = int32(12)
	goto L276
L275:
	;
	v1006 = int32(11)
	goto L276
L276:
	;
	v1016 = v1006
	v1017 = v967
	goto L247
L277:
	;
	goto L250
L278:
	;
	goto L289
L279:
	;
	v1023 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v951+v1020))) = uint8(v1023)
	v1026 = v1020 + int32(-1)
	if base.Ui64(int64(100)) <= base.Ui64(v78) {
		goto L281
	} else {
		goto L282
	}
L280:
	;
	v1062 = v951 + v1059
	if base.Ui64(int64(9)) < base.Ui64(v1060) {
		goto L286
	} else {
		goto L287
	}
L281:
	;
	v1033 = v78
	v1035 = v1026
	goto L283
L282:
	;
	v1059 = v1026
	v1060 = v78
	goto L280
L283:
	;
	v1039 = int64(100)
	v1040 = base.I64_div_u_s(v1033, v1039)
	v1049 = int32(*(*uint16)(unsafe.Add(mBase, uint32(base.I32_wrap_i64(v1033-v1040*v1039)<<(uint(int32(1))%32))+uint32(_consts[871]))))
	*(*uint16)(unsafe.Add(mBase, uint32(v21+int32(559)+v1035))) = uint16(v1049)
	v1052 = v1035 + int32(-2)
	if base.Ui64(int64(9999)) < base.Ui64(v1033) {
		v1033 = v1040
		v1035 = v1052
		goto L283
	} else {
		goto L285
	}
L284:
	;
	v1059 = v1052
	v1060 = v1040
	goto L280
L285:
	;
	goto L284
L286:
	;
	v1076 = int32(*(*uint16)(unsafe.Add(mBase, uint32(base.I32_wrap_i64(v1060)<<(uint(int32(1))%32))+uint32(_consts[871]))))
	*(*uint16)(unsafe.Add(mBase, uint32(v1062+int32(-1)))) = uint16(v1076)
	v1090 = v1020
	goto L246
L287:
	;
	v1067 = base.I32_wrap_i64(v1060) | int32(48)
	*(*uint8)(unsafe.Add(mBase, uint32(v1062))) = uint8(v1067)
	v1090 = v1020
	goto L246
L288:
	;
	v1090 = int32(0)
	goto L246
L289:
	;
	v1080 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v951))) = uint8(v1080)
	goto L288
L290:
	;
	v1240 = F_sdsnewlen(m, v21+int32(560), v1095+v1235-(v21+int32(560)))
	mBase = m.M
	v1241 = m.ExcPending
	if v1241 != 0 {
		goto L7
	} else {
		goto L334
	}
L291:
	;
	v1165 = v1161 + v1162
	if base.Ui32(int32(21)) <= base.Ui32(v1165) {
		goto L322
	} else {
		goto L323
	}
L292:
	;
	v1112 = v1096
	v1113 = v416
	goto L293
L293:
	;
	if base.Ui64(int64(100)) <= base.Ui64(v1113) {
		goto L295
	} else {
		goto L296
	}
L294:
	;
	v1161 = v1094
	v1162 = v1153
	goto L291
L295:
	;
	if base.Ui64(int64(1000)) <= base.Ui64(v1113) {
		goto L297
	} else {
		goto L298
	}
L296:
	;
	v1161 = int32(2)
	v1162 = v1112
	goto L291
L297:
	;
	if base.Ui64(int64(999999999999)) < base.Ui64(v1113) {
		goto L299
	} else {
		goto L300
	}
L298:
	;
	v1161 = int32(3)
	v1162 = v1112
	goto L291
L299:
	;
	v1153 = v1112 + int32(12)
	v1157 = base.I64_div_u_s(v1113, int64(1000000000000))
	if base.Ui64(int64(9999999999999)) < base.Ui64(v1113) {
		v1112 = v1153
		v1113 = v1157
		goto L293
	} else {
		goto L321
	}
L300:
	;
	if base.Ui64(int64(99999999)) < base.Ui64(v1113) {
		goto L301
	} else {
		goto L302
	}
L301:
	;
	if base.Ui64(int64(9999999999)) < base.Ui64(v1113) {
		goto L313
	} else {
		goto L314
	}
L302:
	;
	if base.Ui64(int64(999999)) < base.Ui64(v1113) {
		goto L303
	} else {
		goto L304
	}
L303:
	;
	if base.Ui64(int64(9999999)) < base.Ui64(v1113) {
		goto L310
	} else {
		goto L311
	}
L304:
	;
	if base.Ui64(int64(10000)) <= base.Ui64(v1113) {
		goto L305
	} else {
		goto L306
	}
L305:
	;
	if base.Ui64(int64(99999)) < base.Ui64(v1113) {
		goto L307
	} else {
		goto L308
	}
L306:
	;
	v1161 = int32(4)
	v1162 = v1112
	goto L291
L307:
	;
	v1134 = int32(6)
	goto L309
L308:
	;
	v1134 = int32(5)
	goto L309
L309:
	;
	v1161 = v1134
	v1162 = v1112
	goto L291
L310:
	;
	v1139 = int32(8)
	goto L312
L311:
	;
	v1139 = int32(7)
	goto L312
L312:
	;
	v1161 = v1139
	v1162 = v1112
	goto L291
L313:
	;
	if base.Ui64(int64(99999999999)) < base.Ui64(v1113) {
		goto L318
	} else {
		goto L319
	}
L314:
	;
	if base.Ui64(int64(999999999)) < base.Ui64(v1113) {
		goto L315
	} else {
		goto L316
	}
L315:
	;
	v1146 = int32(10)
	goto L317
L316:
	;
	v1146 = int32(9)
	goto L317
L317:
	;
	v1161 = v1146
	v1162 = v1112
	goto L291
L318:
	;
	v1151 = int32(12)
	goto L320
L319:
	;
	v1151 = int32(11)
	goto L320
L320:
	;
	v1161 = v1151
	v1162 = v1112
	goto L291
L321:
	;
	goto L294
L322:
	;
	goto L333
L323:
	;
	v1168 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v1095+v1165))) = uint8(v1168)
	v1171 = v1165 + int32(-1)
	if base.Ui64(int64(100)) <= base.Ui64(v416) {
		goto L325
	} else {
		goto L326
	}
L324:
	;
	v1207 = v1095 + v1204
	if base.Ui64(int64(9)) < base.Ui64(v1205) {
		goto L330
	} else {
		goto L331
	}
L325:
	;
	v1178 = v416
	v1180 = v1171
	goto L327
L326:
	;
	v1204 = v1171
	v1205 = v416
	goto L324
L327:
	;
	v1184 = int64(100)
	v1185 = base.I64_div_u_s(v1178, v1184)
	v1194 = int32(*(*uint16)(unsafe.Add(mBase, uint32(base.I32_wrap_i64(v1178-v1185*v1184)<<(uint(int32(1))%32))+uint32(_consts[871]))))
	*(*uint16)(unsafe.Add(mBase, uint32(v1091+int32(0)+v1180))) = uint16(v1194)
	v1197 = v1180 + int32(-2)
	if base.Ui64(int64(9999)) < base.Ui64(v1178) {
		v1178 = v1185
		v1180 = v1197
		goto L327
	} else {
		goto L329
	}
L328:
	;
	v1204 = v1197
	v1205 = v1185
	goto L324
L329:
	;
	goto L328
L330:
	;
	v1221 = int32(*(*uint16)(unsafe.Add(mBase, uint32(base.I32_wrap_i64(v1205)<<(uint(int32(1))%32))+uint32(_consts[871]))))
	*(*uint16)(unsafe.Add(mBase, uint32(v1207+int32(-1)))) = uint16(v1221)
	v1235 = v1165
	goto L290
L331:
	;
	v1212 = base.I32_wrap_i64(v1205) | int32(48)
	*(*uint8)(unsafe.Add(mBase, uint32(v1207))) = uint8(v1212)
	v1235 = v1165
	goto L290
L332:
	;
	v1235 = int32(0)
	goto L290
L333:
	;
	v1225 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v1095))) = uint8(v1225)
	goto L332
L334:
	;
	v1242 = F_createObject(m, v1096, v1240)
	mBase = m.M
	v1243 = m.ExcPending
	if v1243 != 0 {
		goto L7
	} else {
		goto L335
	}
L335:
	;
	v1244 = *(*int32)(unsafe.Add(mBase, uint32(l9)))
	v1245 = *(*int32)(unsafe.Add(mBase, uint32(l9)+4))
	F_streamPropagateXCLAIM(m, l0, v1244, l6, v1245, v1242, v943)
	mBase = m.M
	v1247 = m.ExcPending
	if v1247 != 0 {
		goto L7
	} else {
		goto L336
	}
L336:
	;
	F_decrRefCount(m, v1242)
	mBase = m.M
	v1249 = m.ExcPending
	if v1249 != 0 {
		goto L7
	} else {
		goto L337
	}
L337:
	;
	goto L190
L338:
	;
	if l4 != v1255 {
		v53 = v1255
		v63 = v260
		goto L14
	} else {
		goto L339
	}
L339:
	;
	goto L15
L340:
	;
	F_raxStop(m, v21+int32(200))
	mBase = m.M
	v1288 = m.ExcPending
	if v1288 != 0 {
		goto L7
	} else {
		goto L344
	}
L341:
	;
	if v1272 == int32(0) {
		goto L340
	} else {
		goto L342
	}
L342:
	;
	v1281 = *(*int32)(unsafe.Add(mBase, uint32(l9)))
	v1282 = *(*int32)(unsafe.Add(mBase, uint32(l9)+4))
	F_streamPropagateGroupID(m, l0, v1281, l6, v1282)
	mBase = m.M
	v1284 = m.ExcPending
	if v1284 != 0 {
		goto L7
	} else {
		goto L343
	}
L343:
	;
	goto L340
L344:
	;
	if v40 == int32(0) {
		v1296 = v1262
		goto L3
	} else {
		goto L345
	}
L345:
	;
	F_setDeferredArrayLen(m, l0, v40, v1262)
	mBase = m.M
	v1292 = m.ExcPending
	if v1292 != 0 {
		goto L7
	} else {
		goto L346
	}
L346:
	;
	v1296 = v1262
	goto L3
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
}
func F_streamReplyWithRangeFromConsumerPEL(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32 {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v15 int64
	_ = v15
	var v16 int64
	_ = v16
	var v17 int64
	_ = v17
	var v19 int64
	_ = v19
	var v21 int64
	_ = v21
	var v24 int64
	_ = v24
	var v26 int64
	_ = v26
	var v28 int64
	_ = v28
	var v30 int64
	_ = v30
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v101 int64
	_ = v101
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
	var v134 int32
	_ = v134
	var v139 int32
	_ = v139
	var v142 int32
	_ = v142
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v161 int32
	_ = v161
	var v163 int32
	_ = v163
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v186 int32
	_ = v186
	var v191 int32
	_ = v191
	var v206 int32
	_ = v206
	var v209 int64
	_ = v209
	var v210 int64
	_ = v210
	var v211 int64
	_ = v211
	var v213 int64
	_ = v213
	var v215 int64
	_ = v215
	var v218 int64
	_ = v218
	var v220 int64
	_ = v220
	var v222 int64
	_ = v222
	var v224 int64
	_ = v224
	var v284 int32
	_ = v284
	var v289 int32
	_ = v289
	var v290 int32
	_ = v290
	var v293 int32
	_ = v293
	var v295 int32
	_ = v295
	var v299 int64
	_ = v299
	var v300 int32
	_ = v300
	var v304 int32
	_ = v304
	var v312 int32
	_ = v312
	var v313 int64
	_ = v313
	var v334 int32
	_ = v334
	var v339 int32
	_ = v339
	var v346 int32
	_ = v346
	var v351 int32
	_ = v351
	var v353 int32
	_ = v353
	var v357 int64
	_ = v357
	var v361 int32
	_ = v361
	var v362 int32
	_ = v362
	var v365 int32
	_ = v365
	var v368 int32
	_ = v368
	var v371 int32
	_ = v371
	var v378 int64
	_ = v378
	var v380 int32
	_ = v380
	var v384 int64
	_ = v384
	var v385 int64
	_ = v385
	var v394 int32
	_ = v394
	var v397 int32
	_ = v397
	var v404 int32
	_ = v404
	var v405 int64
	_ = v405
	var v407 int32
	_ = v407
	var v412 int32
	_ = v412
	var v421 int32
	_ = v421
	var v425 int32
	_ = v425
	var v435 int32
	_ = v435
	var v436 int32
	_ = v436
	var v437 int32
	_ = v437
	var v439 int32
	_ = v439
	var v440 int32
	_ = v440
	var v444 int64
	_ = v444
	var v445 int32
	_ = v445
	var v457 int32
	_ = v457
	var v458 int64
	_ = v458
	var v479 int32
	_ = v479
	var v484 int32
	_ = v484
	var v491 int32
	_ = v491
	var v496 int32
	_ = v496
	var v498 int32
	_ = v498
	var v502 int64
	_ = v502
	var v506 int32
	_ = v506
	var v507 int32
	_ = v507
	var v510 int32
	_ = v510
	var v513 int32
	_ = v513
	var v516 int32
	_ = v516
	var v523 int64
	_ = v523
	var v525 int32
	_ = v525
	var v529 int64
	_ = v529
	var v530 int64
	_ = v530
	var v539 int32
	_ = v539
	var v542 int32
	_ = v542
	var v549 int32
	_ = v549
	var v550 int64
	_ = v550
	var v552 int32
	_ = v552
	var v557 int32
	_ = v557
	var v566 int32
	_ = v566
	var v570 int32
	_ = v570
	var v580 int32
	_ = v580
	var v585 int32
	_ = v585
	var v586 int32
	_ = v586
	var v588 int32
	_ = v588
	var v590 int32
	_ = v590
	var v591 int32
	_ = v591
	var v593 int64
	_ = v593
	var v595 int64
	_ = v595
	var v601 int32
	_ = v601
	var v604 int32
	_ = v604
	var v605 int32
	_ = v605
	var v614 int32
	_ = v614
	var v622 int32
	_ = v622
	var v624 int32
	_ = v624
	v11 = m.G0
	v13 = v11 - int32(384)
	m.G0 = v13
	v15 = *(*int64)(unsafe.Add(mBase, uint32(l2)))
	v16 = *(*int64)(unsafe.Add(mBase, uint32(l2)+8))
	v17 = int64(56)
	v19 = int64(65280)
	v21 = int64(40)
	v24 = int64(16711680)
	v26 = int64(24)
	v28 = int64(4278190080)
	v30 = int64(8)
	*(*int64)(unsafe.Add(mBase, uint32(v13)+24)) = v16<<(uint(v17)%64) | v16&v19<<(uint(v21)%64) | (v16&v24<<(uint(v26)%64) | v16&v28<<(uint(v30)%64)) | (int64(base.Ui64(v16)>>(uint(v30)%64))&v28 | int64(base.Ui64(v16)>>(uint(v26)%64))&v24 | (int64(base.Ui64(v16)>>(uint(v21)%64))&v19 | int64(base.Ui64(v16)>>(uint(v17)%64))))
	*(*int64)(unsafe.Add(mBase, uint32(v13)+16)) = v15<<(uint(v17)%64) | v15&v19<<(uint(v21)%64) | (v15&v24<<(uint(v26)%64) | v15&v28<<(uint(v30)%64)) | (int64(base.Ui64(v15)>>(uint(v30)%64))&v28 | int64(base.Ui64(v15)>>(uint(v26)%64))&v24 | (int64(base.Ui64(v15)>>(uint(v21)%64))&v19 | int64(base.Ui64(v15)>>(uint(v17)%64))))
	v89 = F_addReplyDeferredLen(m, l0)
	mBase = m.M
	v92 = m.ExcPending
	if v92 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v94 = v13 + int32(32)
	v95 = *(*int32)(unsafe.Add(mBase, uint32(l5)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v94)+4)) = v95
	*(*int32)(unsafe.Add(mBase, uint32(v94))) = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v94)+20)) = int32(128)
	v101 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v94)+12)) = v101
	*(*int64)(unsafe.Add(mBase, uint32(v94)+296)) = v101
	*(*int64)(unsafe.Add(mBase, uint32(v94)+160)) = int64(137438953472)
	*(*int32)(unsafe.Add(mBase, uint32(v94)+8)) = v13 + int32(56)
	*(*int32)(unsafe.Add(mBase, uint32(v94)+156)) = v13 + int32(200)
	goto L3
L3:
	;
	v116 = int32(16)
	v119 = F_raxSeek(m, v13+int32(32), int32(_a2407), v13+v116, v116)
	mBase = m.M
	v120 = m.ExcPending
	if v120 != 0 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v123 = F_raxNext(m, v13+int32(32))
	mBase = m.M
	v124 = m.ExcPending
	if v124 != 0 {
		goto L1
	} else {
		goto L7
	}
L5:
	;
	F_raxStop(m, v13+int32(32))
	mBase = m.M
	v622 = m.ExcPending
	if v622 != 0 {
		goto L1
	} else {
		goto L130
	}
L6:
	;
	v134 = int32(0)
	goto L9
L7:
	;
	if v123 != 0 {
		goto L6
	} else {
		goto L8
	}
L8:
	;
	v614 = int32(0)
	goto L5
L9:
	;
	v139 = *(*int32)(unsafe.Add(mBase, uint32(v13)+40))
	if l3 == int32(0) {
		goto L11
	} else {
		goto L12
	}
L10:
	;
	v614 = v601
	goto L5
L11:
	;
	v209 = *(*int64)(unsafe.Add(mBase, uint32(v139)))
	v210 = *(*int64)(unsafe.Add(mBase, uint32(v139)+8))
	v211 = int64(56)
	v213 = int64(65280)
	v215 = int64(40)
	v218 = int64(16711680)
	v220 = int64(24)
	v222 = int64(4278190080)
	v224 = int64(8)
	*(*int64)(unsafe.Add(mBase, uint32(v13)+8)) = v210<<(uint(v211)%64) | v210&v213<<(uint(v215)%64) | (v210&v218<<(uint(v220)%64) | v210&v222<<(uint(v224)%64)) | (int64(base.Ui64(v210)>>(uint(v224)%64))&v222 | int64(base.Ui64(v210)>>(uint(v220)%64))&v218 | (int64(base.Ui64(v210)>>(uint(v215)%64))&v213 | int64(base.Ui64(v210)>>(uint(v211)%64))))
	*(*int64)(unsafe.Add(mBase, uint32(v13))) = v209<<(uint(v211)%64) | v209&v213<<(uint(v215)%64) | (v209&v218<<(uint(v220)%64) | v209&v222<<(uint(v224)%64)) | (int64(base.Ui64(v209)>>(uint(v224)%64))&v222 | int64(base.Ui64(v209)>>(uint(v220)%64))&v218 | (int64(base.Ui64(v209)>>(uint(v215)%64))&v213 | int64(base.Ui64(v209)>>(uint(v211)%64))))
	v284 = int32(0)
	v289 = F_streamReplyWithRange(m, l0, l1, v13, v13, int32(1), v284, v284, v284, int32(2), v284)
	mBase = m.M
	v290 = m.ExcPending
	if v290 != 0 {
		goto L1
	} else {
		goto L32
	}
L12:
	;
	v142 = *(*int32)(unsafe.Add(mBase, uint32(v13)+48))
	if base.Ui32(v142) < base.Ui32(int32(4)) {
		v166 = v139
		v167 = l3
		v168 = v142
		goto L16
	} else {
		goto L17
	}
L13:
	;
	if int32(0) < v206 {
		v614 = v134
		goto L5
	} else {
		goto L29
	}
L14:
	;
	v206 = int32(0)
	goto L13
L15:
	;
	v178 = v173
	v179 = v174
	v180 = v175
	goto L25
L16:
	;
	if v168 == int32(0) {
		goto L14
	} else {
		goto L23
	}
L17:
	;
	if (l3|v139)&int32(3) != 0 {
		v173 = v139
		v174 = l3
		v175 = v142
		goto L15
	} else {
		goto L18
	}
L18:
	;
	v150 = v139
	v151 = l3
	v152 = v142
	goto L19
L19:
	;
	v155 = *(*int32)(unsafe.Add(mBase, uint32(v150)))
	v156 = *(*int32)(unsafe.Add(mBase, uint32(v151)))
	if v155 != v156 {
		v173 = v150
		v174 = v151
		v175 = v152
		goto L15
	} else {
		goto L21
	}
L20:
	;
	v166 = v161
	v167 = v159
	v168 = v163
	goto L16
L21:
	;
	v158 = int32(4)
	v159 = v151 + v158
	v161 = v150 + v158
	v163 = v152 + int32(-4)
	if base.Ui32(int32(3)) < base.Ui32(v163) {
		v150 = v161
		v151 = v159
		v152 = v163
		goto L19
	} else {
		goto L22
	}
L22:
	;
	goto L20
L23:
	;
	v173 = v166
	v174 = v167
	v175 = v168
	goto L15
L24:
	;
	v206 = v183 - v184
	goto L13
L25:
	;
	v183 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v178))))
	v184 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v179))))
	if v183 != v184 {
		goto L24
	} else {
		goto L27
	}
L27:
	;
	v186 = int32(1)
	v191 = v180 + int32(-1)
	if v191 == int32(0) {
		goto L14
	} else {
		goto L28
	}
L28:
	;
	v178 = v178 + v186
	v179 = v179 + v186
	v180 = v191
	goto L25
L29:
	;
	goto L11
L30:
	;
	v601 = v134 + int32(1)
	v604 = F_raxNext(m, v13+int32(32))
	mBase = m.M
	v605 = m.ExcPending
	if v605 != 0 {
		goto L1
	} else {
		goto L127
	}
L31:
	;
	v591 = *(*int32)(unsafe.Add(mBase, uint32(v13)+44))
	v593 = *(*int64)(unsafe.Add(mBase, _consts[12]))
	goto L126
L32:
	;
	if v289 != 0 {
		goto L31
	} else {
		goto L33
	}
L33:
	;
	F_addReplyArrayLen(m, l0, int32(2))
	mBase = m.M
	v293 = m.ExcPending
	if v293 != 0 {
		goto L1
	} else {
		goto L34
	}
L34:
	;
	v295 = v13 + int32(336)
	v299 = *(*int64)(unsafe.Add(mBase, uint32(v13)))
	v300 = int32(0)
	v304 = int32(1)
	if base.Ui64(v299) < base.Ui64(int64(10)) {
		v361 = v304
		v362 = v300
		goto L36
	} else {
		goto L37
	}
L35:
	;
	v436 = v295 + v435
	v437 = int32(45)
	*(*uint8)(unsafe.Add(mBase, uint32(v436))) = uint8(v437)
	v439 = int32(1)
	v440 = v436 + v439
	v444 = *(*int64)(unsafe.Add(mBase, uint32(v13)+8))
	v445 = int32(0)
	if base.Ui64(v444) < base.Ui64(int64(10)) {
		v506 = v439
		v507 = v445
		goto L80
	} else {
		goto L81
	}
L36:
	;
	v365 = v361 + v362
	if base.Ui32(int32(21)) <= base.Ui32(v365) {
		goto L67
	} else {
		goto L68
	}
L37:
	;
	v312 = v300
	v313 = v299
	goto L38
L38:
	;
	if base.Ui64(int64(100)) <= base.Ui64(v313) {
		goto L40
	} else {
		goto L41
	}
L39:
	;
	v361 = v304
	v362 = v353
	goto L36
L40:
	;
	if base.Ui64(int64(1000)) <= base.Ui64(v313) {
		goto L42
	} else {
		goto L43
	}
L41:
	;
	v361 = int32(2)
	v362 = v312
	goto L36
L42:
	;
	if base.Ui64(int64(999999999999)) < base.Ui64(v313) {
		goto L44
	} else {
		goto L45
	}
L43:
	;
	v361 = int32(3)
	v362 = v312
	goto L36
L44:
	;
	v353 = v312 + int32(12)
	v357 = base.I64_div_u_s(v313, int64(1000000000000))
	if base.Ui64(int64(9999999999999)) < base.Ui64(v313) {
		v312 = v353
		v313 = v357
		goto L38
	} else {
		goto L66
	}
L45:
	;
	if base.Ui64(int64(99999999)) < base.Ui64(v313) {
		goto L46
	} else {
		goto L47
	}
L46:
	;
	if base.Ui64(int64(9999999999)) < base.Ui64(v313) {
		goto L58
	} else {
		goto L59
	}
L47:
	;
	if base.Ui64(int64(999999)) < base.Ui64(v313) {
		goto L48
	} else {
		goto L49
	}
L48:
	;
	if base.Ui64(int64(9999999)) < base.Ui64(v313) {
		goto L55
	} else {
		goto L56
	}
L49:
	;
	if base.Ui64(int64(10000)) <= base.Ui64(v313) {
		goto L50
	} else {
		goto L51
	}
L50:
	;
	if base.Ui64(int64(99999)) < base.Ui64(v313) {
		goto L52
	} else {
		goto L53
	}
L51:
	;
	v361 = int32(4)
	v362 = v312
	goto L36
L52:
	;
	v334 = int32(6)
	goto L54
L53:
	;
	v334 = int32(5)
	goto L54
L54:
	;
	v361 = v334
	v362 = v312
	goto L36
L55:
	;
	v339 = int32(8)
	goto L57
L56:
	;
	v339 = int32(7)
	goto L57
L57:
	;
	v361 = v339
	v362 = v312
	goto L36
L58:
	;
	if base.Ui64(int64(99999999999)) < base.Ui64(v313) {
		goto L63
	} else {
		goto L64
	}
L59:
	;
	if base.Ui64(int64(999999999)) < base.Ui64(v313) {
		goto L60
	} else {
		goto L61
	}
L60:
	;
	v346 = int32(10)
	goto L62
L61:
	;
	v346 = int32(9)
	goto L62
L62:
	;
	v361 = v346
	v362 = v312
	goto L36
L63:
	;
	v351 = int32(12)
	goto L65
L64:
	;
	v351 = int32(11)
	goto L65
L65:
	;
	v361 = v351
	v362 = v312
	goto L36
L66:
	;
	goto L39
L67:
	;
	goto L78
L68:
	;
	v368 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v295+v365))) = uint8(v368)
	v371 = v365 + int32(-1)
	if base.Ui64(int64(100)) <= base.Ui64(v299) {
		goto L70
	} else {
		goto L71
	}
L69:
	;
	v407 = v295 + v404
	if base.Ui64(int64(9)) < base.Ui64(v405) {
		goto L75
	} else {
		goto L76
	}
L70:
	;
	v378 = v299
	v380 = v371
	goto L72
L71:
	;
	v404 = v371
	v405 = v299
	goto L69
L72:
	;
	v384 = int64(100)
	v385 = base.I64_div_u_s(v378, v384)
	v394 = int32(*(*uint16)(unsafe.Add(mBase, uint32(base.I32_wrap_i64(v378-v385*v384)<<(uint(int32(1))%32))+uint32(_consts[871]))))
	*(*uint16)(unsafe.Add(mBase, uint32(v13+int32(335)+v380))) = uint16(v394)
	v397 = v380 + int32(-2)
	if base.Ui64(int64(9999)) < base.Ui64(v378) {
		v378 = v385
		v380 = v397
		goto L72
	} else {
		goto L74
	}
L73:
	;
	v404 = v397
	v405 = v385
	goto L69
L74:
	;
	goto L73
L75:
	;
	v421 = int32(*(*uint16)(unsafe.Add(mBase, uint32(base.I32_wrap_i64(v405)<<(uint(int32(1))%32))+uint32(_consts[871]))))
	*(*uint16)(unsafe.Add(mBase, uint32(v407+int32(-1)))) = uint16(v421)
	v435 = v365
	goto L35
L76:
	;
	v412 = base.I32_wrap_i64(v405) | int32(48)
	*(*uint8)(unsafe.Add(mBase, uint32(v407))) = uint8(v412)
	v435 = v365
	goto L35
L77:
	;
	v435 = int32(0)
	goto L35
L78:
	;
	v425 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v295))) = uint8(v425)
	goto L77
L79:
	;
	v585 = F_sdsnewlen(m, v13+int32(336), v440+v580-(v13+int32(336)))
	mBase = m.M
	v586 = m.ExcPending
	if v586 != 0 {
		goto L1
	} else {
		goto L123
	}
L80:
	;
	v510 = v506 + v507
	if base.Ui32(int32(21)) <= base.Ui32(v510) {
		goto L111
	} else {
		goto L112
	}
L81:
	;
	v457 = v445
	v458 = v444
	goto L82
L82:
	;
	if base.Ui64(int64(100)) <= base.Ui64(v458) {
		goto L84
	} else {
		goto L85
	}
L83:
	;
	v506 = v439
	v507 = v498
	goto L80
L84:
	;
	if base.Ui64(int64(1000)) <= base.Ui64(v458) {
		goto L86
	} else {
		goto L87
	}
L85:
	;
	v506 = int32(2)
	v507 = v457
	goto L80
L86:
	;
	if base.Ui64(int64(999999999999)) < base.Ui64(v458) {
		goto L88
	} else {
		goto L89
	}
L87:
	;
	v506 = int32(3)
	v507 = v457
	goto L80
L88:
	;
	v498 = v457 + int32(12)
	v502 = base.I64_div_u_s(v458, int64(1000000000000))
	if base.Ui64(int64(9999999999999)) < base.Ui64(v458) {
		v457 = v498
		v458 = v502
		goto L82
	} else {
		goto L110
	}
L89:
	;
	if base.Ui64(int64(99999999)) < base.Ui64(v458) {
		goto L90
	} else {
		goto L91
	}
L90:
	;
	if base.Ui64(int64(9999999999)) < base.Ui64(v458) {
		goto L102
	} else {
		goto L103
	}
L91:
	;
	if base.Ui64(int64(999999)) < base.Ui64(v458) {
		goto L92
	} else {
		goto L93
	}
L92:
	;
	if base.Ui64(int64(9999999)) < base.Ui64(v458) {
		goto L99
	} else {
		goto L100
	}
L93:
	;
	if base.Ui64(int64(10000)) <= base.Ui64(v458) {
		goto L94
	} else {
		goto L95
	}
L94:
	;
	if base.Ui64(int64(99999)) < base.Ui64(v458) {
		goto L96
	} else {
		goto L97
	}
L95:
	;
	v506 = int32(4)
	v507 = v457
	goto L80
L96:
	;
	v479 = int32(6)
	goto L98
L97:
	;
	v479 = int32(5)
	goto L98
L98:
	;
	v506 = v479
	v507 = v457
	goto L80
L99:
	;
	v484 = int32(8)
	goto L101
L100:
	;
	v484 = int32(7)
	goto L101
L101:
	;
	v506 = v484
	v507 = v457
	goto L80
L102:
	;
	if base.Ui64(int64(99999999999)) < base.Ui64(v458) {
		goto L107
	} else {
		goto L108
	}
L103:
	;
	if base.Ui64(int64(999999999)) < base.Ui64(v458) {
		goto L104
	} else {
		goto L105
	}
L104:
	;
	v491 = int32(10)
	goto L106
L105:
	;
	v491 = int32(9)
	goto L106
L106:
	;
	v506 = v491
	v507 = v457
	goto L80
L107:
	;
	v496 = int32(12)
	goto L109
L108:
	;
	v496 = int32(11)
	goto L109
L109:
	;
	v506 = v496
	v507 = v457
	goto L80
L110:
	;
	goto L83
L111:
	;
	goto L122
L112:
	;
	v513 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v440+v510))) = uint8(v513)
	v516 = v510 + int32(-1)
	if base.Ui64(int64(100)) <= base.Ui64(v444) {
		goto L114
	} else {
		goto L115
	}
L113:
	;
	v552 = v440 + v549
	if base.Ui64(int64(9)) < base.Ui64(v550) {
		goto L119
	} else {
		goto L120
	}
L114:
	;
	v523 = v444
	v525 = v516
	goto L116
L115:
	;
	v549 = v516
	v550 = v444
	goto L113
L116:
	;
	v529 = int64(100)
	v530 = base.I64_div_u_s(v523, v529)
	v539 = int32(*(*uint16)(unsafe.Add(mBase, uint32(base.I32_wrap_i64(v523-v530*v529)<<(uint(int32(1))%32))+uint32(_consts[871]))))
	*(*uint16)(unsafe.Add(mBase, uint32(v436+int32(0)+v525))) = uint16(v539)
	v542 = v525 + int32(-2)
	if base.Ui64(int64(9999)) < base.Ui64(v523) {
		v523 = v530
		v525 = v542
		goto L116
	} else {
		goto L118
	}
L117:
	;
	v549 = v542
	v550 = v530
	goto L113
L118:
	;
	goto L117
L119:
	;
	v566 = int32(*(*uint16)(unsafe.Add(mBase, uint32(base.I32_wrap_i64(v550)<<(uint(int32(1))%32))+uint32(_consts[871]))))
	*(*uint16)(unsafe.Add(mBase, uint32(v552+int32(-1)))) = uint16(v566)
	v580 = v510
	goto L79
L120:
	;
	v557 = base.I32_wrap_i64(v550) | int32(48)
	*(*uint8)(unsafe.Add(mBase, uint32(v552))) = uint8(v557)
	v580 = v510
	goto L79
L121:
	;
	v580 = int32(0)
	goto L79
L122:
	;
	v570 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v440))) = uint8(v570)
	goto L121
L123:
	;
	F_addReplyBulkSds(m, l0, v585)
	mBase = m.M
	v588 = m.ExcPending
	if v588 != 0 {
		goto L1
	} else {
		goto L124
	}
L124:
	;
	F_addReplyNullArray(m, l0)
	mBase = m.M
	v590 = m.ExcPending
	if v590 != 0 {
		goto L1
	} else {
		goto L125
	}
L125:
	;
	goto L30
L126:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v591))) = v593
	v595 = *(*int64)(unsafe.Add(mBase, uint32(v591)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v591)+8)) = v595 + int64(1)
	goto L30
L127:
	;
	if v604 == int32(0) {
		v614 = v601
		goto L5
	} else {
		goto L128
	}
L128:
	;
	if base.Ui32(v601) <= base.Ui32(l4+int32(-1)) {
		v134 = v601
		goto L9
	} else {
		goto L129
	}
L129:
	;
	goto L10
L130:
	;
	F_setDeferredArrayLen(m, l0, v89, v614)
	mBase = m.M
	v624 = m.ExcPending
	if v624 != 0 {
		goto L1
	} else {
		goto L131
	}
L131:
	;
	m.G0 = v13 + int32(384)
	return v614
}
func F_streamRewriteTrimArgument(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v14 int64
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v19 int64
	_ = v19
	var v23 int64
	_ = v23
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v54 int64
	_ = v54
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v67 int64
	_ = v67
	var v68 int64
	_ = v68
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v77 int64
	_ = v77
	var v78 int32
	_ = v78
	var v82 int32
	_ = v82
	var v90 int32
	_ = v90
	var v91 int64
	_ = v91
	var v112 int32
	_ = v112
	var v117 int32
	_ = v117
	var v124 int32
	_ = v124
	var v129 int32
	_ = v129
	var v131 int32
	_ = v131
	var v135 int64
	_ = v135
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v143 int32
	_ = v143
	var v146 int32
	_ = v146
	var v149 int32
	_ = v149
	var v156 int64
	_ = v156
	var v158 int32
	_ = v158
	var v162 int64
	_ = v162
	var v163 int64
	_ = v163
	var v172 int32
	_ = v172
	var v175 int32
	_ = v175
	var v182 int32
	_ = v182
	var v183 int64
	_ = v183
	var v185 int32
	_ = v185
	var v190 int32
	_ = v190
	var v199 int32
	_ = v199
	var v203 int32
	_ = v203
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v218 int32
	_ = v218
	var v224 int64
	_ = v224
	var v225 int32
	_ = v225
	var v229 int32
	_ = v229
	var v237 int32
	_ = v237
	var v238 int64
	_ = v238
	var v259 int32
	_ = v259
	var v264 int32
	_ = v264
	var v271 int32
	_ = v271
	var v276 int32
	_ = v276
	var v278 int32
	_ = v278
	var v282 int64
	_ = v282
	var v286 int32
	_ = v286
	var v287 int32
	_ = v287
	var v290 int32
	_ = v290
	var v293 int32
	_ = v293
	var v296 int32
	_ = v296
	var v303 int64
	_ = v303
	var v305 int32
	_ = v305
	var v309 int64
	_ = v309
	var v310 int64
	_ = v310
	var v319 int32
	_ = v319
	var v322 int32
	_ = v322
	var v329 int32
	_ = v329
	var v330 int64
	_ = v330
	var v332 int32
	_ = v332
	var v337 int32
	_ = v337
	var v346 int32
	_ = v346
	var v350 int32
	_ = v350
	var v360 int32
	_ = v360
	var v365 int32
	_ = v365
	var v366 int32
	_ = v366
	var v367 int32
	_ = v367
	var v368 int32
	_ = v368
	var v369 int32
	_ = v369
	var v374 int32
	_ = v374
	var v376 int32
	_ = v376
	v8 = m.G0
	v10 = v8 - int32(480)
	m.G0 = v10
	if l2 != int32(1) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	F_rewriteClientCommandArgument(m, l0, l3, v369)
	mBase = m.M
	v374 = m.ExcPending
	if v374 != 0 {
		goto L4
	} else {
		goto L106
	}
L2:
	;
	v19 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v10+int32(96)))) = v19
	v23 = int64(-1)
	*(*int64)(unsafe.Add(mBase, uint32(v10+int32(112)))) = v23
	*(*int64)(unsafe.Add(mBase, uint32(v10)+88)) = v19
	*(*int64)(unsafe.Add(mBase, uint32(v10)+104)) = v23
	v30 = v10 + int32(120)
	v31 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v30)+4)) = v31
	*(*int32)(unsafe.Add(mBase, uint32(v30))) = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v30)+20)) = int32(128)
	*(*int64)(unsafe.Add(mBase, uint32(v30)+12)) = v19
	*(*int64)(unsafe.Add(mBase, uint32(v30)+296)) = v19
	*(*int64)(unsafe.Add(mBase, uint32(v30)+160)) = int64(137438953472)
	*(*int32)(unsafe.Add(mBase, uint32(v30)+8)) = v10 + int32(144)
	*(*int32)(unsafe.Add(mBase, uint32(v30)+156)) = v10 + int32(288)
	goto L6
L3:
	;
	v14 = *(*int64)(unsafe.Add(mBase, uint32(l1)+8))
	v15 = F_createStringObjectFromLongLong(m, v14)
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	return
L5:
	;
	v369 = v15
	goto L1
L6:
	;
	v50 = int32(0)
	v52 = F_raxSeek(m, v30, int32(_a67), v50, v50)
	mBase = m.M
	v53 = m.ExcPending
	if v53 != 0 {
		goto L4
	} else {
		goto L7
	}
L7:
	;
	v54 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v10)+424)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v10)+32)) = l1
	*(*int64)(unsafe.Add(mBase, uint32(v10)+76)) = v54
	v65 = F_streamIteratorGetID(m, v10+int32(32), v10+int32(8), v10+int32(24))
	mBase = m.M
	v66 = m.ExcPending
	if v66 != 0 {
		goto L4
	} else {
		goto L8
	}
L8:
	;
	v67 = *(*int64)(unsafe.Add(mBase, uint32(v10)+16))
	v68 = *(*int64)(unsafe.Add(mBase, uint32(v10)+8))
	F_raxStop(m, v30)
	mBase = m.M
	v70 = m.ExcPending
	if v70 != 0 {
		goto L4
	} else {
		goto L9
	}
L9:
	;
	v72 = v10 + int32(32)
	if v65 != 0 {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	v77 = v68
	goto L12
L11:
	;
	v77 = int64(-1)
	goto L12
L12:
	;
	v78 = int32(0)
	v82 = int32(1)
	if base.Ui64(v77) < base.Ui64(int64(10)) {
		v139 = v82
		v140 = v78
		goto L14
	} else {
		goto L15
	}
L13:
	;
	v214 = v72 + v213
	v215 = int32(45)
	*(*uint8)(unsafe.Add(mBase, uint32(v214))) = uint8(v215)
	v218 = v214 + int32(1)
	if v65 != 0 {
		goto L57
	} else {
		goto L58
	}
L14:
	;
	v143 = v139 + v140
	if base.Ui32(int32(21)) <= base.Ui32(v143) {
		goto L45
	} else {
		goto L46
	}
L15:
	;
	v90 = v78
	v91 = v77
	goto L16
L16:
	;
	if base.Ui64(int64(100)) <= base.Ui64(v91) {
		goto L18
	} else {
		goto L19
	}
L17:
	;
	v139 = v82
	v140 = v131
	goto L14
L18:
	;
	if base.Ui64(int64(1000)) <= base.Ui64(v91) {
		goto L20
	} else {
		goto L21
	}
L19:
	;
	v139 = int32(2)
	v140 = v90
	goto L14
L20:
	;
	if base.Ui64(int64(999999999999)) < base.Ui64(v91) {
		goto L22
	} else {
		goto L23
	}
L21:
	;
	v139 = int32(3)
	v140 = v90
	goto L14
L22:
	;
	v131 = v90 + int32(12)
	v135 = base.I64_div_u_s(v91, int64(1000000000000))
	if base.Ui64(int64(9999999999999)) < base.Ui64(v91) {
		v90 = v131
		v91 = v135
		goto L16
	} else {
		goto L44
	}
L23:
	;
	if base.Ui64(int64(99999999)) < base.Ui64(v91) {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	if base.Ui64(int64(9999999999)) < base.Ui64(v91) {
		goto L36
	} else {
		goto L37
	}
L25:
	;
	if base.Ui64(int64(999999)) < base.Ui64(v91) {
		goto L26
	} else {
		goto L27
	}
L26:
	;
	if base.Ui64(int64(9999999)) < base.Ui64(v91) {
		goto L33
	} else {
		goto L34
	}
L27:
	;
	if base.Ui64(int64(10000)) <= base.Ui64(v91) {
		goto L28
	} else {
		goto L29
	}
L28:
	;
	if base.Ui64(int64(99999)) < base.Ui64(v91) {
		goto L30
	} else {
		goto L31
	}
L29:
	;
	v139 = int32(4)
	v140 = v90
	goto L14
L30:
	;
	v112 = int32(6)
	goto L32
L31:
	;
	v112 = int32(5)
	goto L32
L32:
	;
	v139 = v112
	v140 = v90
	goto L14
L33:
	;
	v117 = int32(8)
	goto L35
L34:
	;
	v117 = int32(7)
	goto L35
L35:
	;
	v139 = v117
	v140 = v90
	goto L14
L36:
	;
	if base.Ui64(int64(99999999999)) < base.Ui64(v91) {
		goto L41
	} else {
		goto L42
	}
L37:
	;
	if base.Ui64(int64(999999999)) < base.Ui64(v91) {
		goto L38
	} else {
		goto L39
	}
L38:
	;
	v124 = int32(10)
	goto L40
L39:
	;
	v124 = int32(9)
	goto L40
L40:
	;
	v139 = v124
	v140 = v90
	goto L14
L41:
	;
	v129 = int32(12)
	goto L43
L42:
	;
	v129 = int32(11)
	goto L43
L43:
	;
	v139 = v129
	v140 = v90
	goto L14
L44:
	;
	goto L17
L45:
	;
	goto L56
L46:
	;
	v146 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v72+v143))) = uint8(v146)
	v149 = v143 + int32(-1)
	if base.Ui64(int64(100)) <= base.Ui64(v77) {
		goto L48
	} else {
		goto L49
	}
L47:
	;
	v185 = v72 + v182
	if base.Ui64(int64(9)) < base.Ui64(v183) {
		goto L53
	} else {
		goto L54
	}
L48:
	;
	v156 = v77
	v158 = v149
	goto L50
L49:
	;
	v182 = v149
	v183 = v77
	goto L47
L50:
	;
	v162 = int64(100)
	v163 = base.I64_div_u_s(v156, v162)
	v172 = int32(*(*uint16)(unsafe.Add(mBase, uint32(base.I32_wrap_i64(v156-v163*v162)<<(uint(int32(1))%32))+uint32(_consts[871]))))
	*(*uint16)(unsafe.Add(mBase, uint32(v10+int32(31)+v158))) = uint16(v172)
	v175 = v158 + int32(-2)
	if base.Ui64(int64(9999)) < base.Ui64(v156) {
		v156 = v163
		v158 = v175
		goto L50
	} else {
		goto L52
	}
L51:
	;
	v182 = v175
	v183 = v163
	goto L47
L52:
	;
	goto L51
L53:
	;
	v199 = int32(*(*uint16)(unsafe.Add(mBase, uint32(base.I32_wrap_i64(v183)<<(uint(int32(1))%32))+uint32(_consts[871]))))
	*(*uint16)(unsafe.Add(mBase, uint32(v185+int32(-1)))) = uint16(v199)
	v213 = v143
	goto L13
L54:
	;
	v190 = base.I32_wrap_i64(v183) | int32(48)
	*(*uint8)(unsafe.Add(mBase, uint32(v185))) = uint8(v190)
	v213 = v143
	goto L13
L55:
	;
	v213 = int32(0)
	goto L13
L56:
	;
	v203 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v72))) = uint8(v203)
	goto L55
L57:
	;
	v224 = v67
	goto L59
L58:
	;
	v224 = int64(-1)
	goto L59
L59:
	;
	v225 = int32(0)
	v229 = int32(1)
	if base.Ui64(v224) < base.Ui64(int64(10)) {
		v286 = v229
		v287 = v225
		goto L61
	} else {
		goto L62
	}
L60:
	;
	v365 = F_sdsnewlen(m, v10+int32(32), v218+v360-(v10+int32(32)))
	mBase = m.M
	v366 = m.ExcPending
	if v366 != 0 {
		goto L4
	} else {
		goto L104
	}
L61:
	;
	v290 = v286 + v287
	if base.Ui32(int32(21)) <= base.Ui32(v290) {
		goto L92
	} else {
		goto L93
	}
L62:
	;
	v237 = v225
	v238 = v224
	goto L63
L63:
	;
	if base.Ui64(int64(100)) <= base.Ui64(v238) {
		goto L65
	} else {
		goto L66
	}
L64:
	;
	v286 = v229
	v287 = v278
	goto L61
L65:
	;
	if base.Ui64(int64(1000)) <= base.Ui64(v238) {
		goto L67
	} else {
		goto L68
	}
L66:
	;
	v286 = int32(2)
	v287 = v237
	goto L61
L67:
	;
	if base.Ui64(int64(999999999999)) < base.Ui64(v238) {
		goto L69
	} else {
		goto L70
	}
L68:
	;
	v286 = int32(3)
	v287 = v237
	goto L61
L69:
	;
	v278 = v237 + int32(12)
	v282 = base.I64_div_u_s(v238, int64(1000000000000))
	if base.Ui64(int64(9999999999999)) < base.Ui64(v238) {
		v237 = v278
		v238 = v282
		goto L63
	} else {
		goto L91
	}
L70:
	;
	if base.Ui64(int64(99999999)) < base.Ui64(v238) {
		goto L71
	} else {
		goto L72
	}
L71:
	;
	if base.Ui64(int64(9999999999)) < base.Ui64(v238) {
		goto L83
	} else {
		goto L84
	}
L72:
	;
	if base.Ui64(int64(999999)) < base.Ui64(v238) {
		goto L73
	} else {
		goto L74
	}
L73:
	;
	if base.Ui64(int64(9999999)) < base.Ui64(v238) {
		goto L80
	} else {
		goto L81
	}
L74:
	;
	if base.Ui64(int64(10000)) <= base.Ui64(v238) {
		goto L75
	} else {
		goto L76
	}
L75:
	;
	if base.Ui64(int64(99999)) < base.Ui64(v238) {
		goto L77
	} else {
		goto L78
	}
L76:
	;
	v286 = int32(4)
	v287 = v237
	goto L61
L77:
	;
	v259 = int32(6)
	goto L79
L78:
	;
	v259 = int32(5)
	goto L79
L79:
	;
	v286 = v259
	v287 = v237
	goto L61
L80:
	;
	v264 = int32(8)
	goto L82
L81:
	;
	v264 = int32(7)
	goto L82
L82:
	;
	v286 = v264
	v287 = v237
	goto L61
L83:
	;
	if base.Ui64(int64(99999999999)) < base.Ui64(v238) {
		goto L88
	} else {
		goto L89
	}
L84:
	;
	if base.Ui64(int64(999999999)) < base.Ui64(v238) {
		goto L85
	} else {
		goto L86
	}
L85:
	;
	v271 = int32(10)
	goto L87
L86:
	;
	v271 = int32(9)
	goto L87
L87:
	;
	v286 = v271
	v287 = v237
	goto L61
L88:
	;
	v276 = int32(12)
	goto L90
L89:
	;
	v276 = int32(11)
	goto L90
L90:
	;
	v286 = v276
	v287 = v237
	goto L61
L91:
	;
	goto L64
L92:
	;
	goto L103
L93:
	;
	v293 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v218+v290))) = uint8(v293)
	v296 = v290 + int32(-1)
	if base.Ui64(int64(100)) <= base.Ui64(v224) {
		goto L95
	} else {
		goto L96
	}
L94:
	;
	v332 = v218 + v329
	if base.Ui64(int64(9)) < base.Ui64(v330) {
		goto L100
	} else {
		goto L101
	}
L95:
	;
	v303 = v224
	v305 = v296
	goto L97
L96:
	;
	v329 = v296
	v330 = v224
	goto L94
L97:
	;
	v309 = int64(100)
	v310 = base.I64_div_u_s(v303, v309)
	v319 = int32(*(*uint16)(unsafe.Add(mBase, uint32(base.I32_wrap_i64(v303-v310*v309)<<(uint(int32(1))%32))+uint32(_consts[871]))))
	*(*uint16)(unsafe.Add(mBase, uint32(v214+int32(0)+v305))) = uint16(v319)
	v322 = v305 + int32(-2)
	if base.Ui64(int64(9999)) < base.Ui64(v303) {
		v303 = v310
		v305 = v322
		goto L97
	} else {
		goto L99
	}
L98:
	;
	v329 = v322
	v330 = v310
	goto L94
L99:
	;
	goto L98
L100:
	;
	v346 = int32(*(*uint16)(unsafe.Add(mBase, uint32(base.I32_wrap_i64(v330)<<(uint(int32(1))%32))+uint32(_consts[871]))))
	*(*uint16)(unsafe.Add(mBase, uint32(v332+int32(-1)))) = uint16(v346)
	v360 = v290
	goto L60
L101:
	;
	v337 = base.I32_wrap_i64(v330) | int32(48)
	*(*uint8)(unsafe.Add(mBase, uint32(v332))) = uint8(v337)
	v360 = v290
	goto L60
L102:
	;
	v360 = int32(0)
	goto L60
L103:
	;
	v350 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v218))) = uint8(v350)
	goto L102
L104:
	;
	v367 = F_createObject(m, int32(0), v365)
	mBase = m.M
	v368 = m.ExcPending
	if v368 != 0 {
		goto L4
	} else {
		goto L105
	}
L105:
	;
	v369 = v367
	goto L1
L106:
	;
	F_decrRefCount(m, v369)
	mBase = m.M
	v376 = m.ExcPending
	if v376 != 0 {
		goto L4
	} else {
		goto L107
	}
L107:
	;
	m.G0 = v10 + int32(480)
	return
}
